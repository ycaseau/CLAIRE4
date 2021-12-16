// ==================================================================
// microCLAIRE                                              CLAIRE 4
// golang microClaire Kernel - started on June 21st, 2020
//
//  Copyright (C) 2020 Yves Caseau. All Rights Reserved.
//  Redistribution and use in source and binary forms are permitted
//  provided that source distribution retains this entire copyright
//  notice and comments.
//  THIS SOFTWARE IS PROVIDED AS IS AND WITHOUT ANY WARRANTY,
//  INCLUDING, WITHOUT LIMITATION, THE IMPLIED WARRANTIES OF
//  MERCHANTABILTY AND FITNESS FOR A PARTICULAR PURPOSE
//
// clKernel.go
// ==================================================================

// This is the go Kernel for microClaire.
// It contains
//  - all the struct definitions
//  - all the low level utility functions
//  - all uses of unsafe.Pointer

package Kernel

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"unsafe"
)

// +---------------------------------------------------------------------------+
// |  Table of contents                                                        |
// |  Part 1: Description of EID (Entity IDs)                                  |
// |  Part 2: description of Claire Objects & Imported (Struct)                |
// |  Part 3: Dummy Classes for instanciation / get/ set                       |
// |  Part 4: unsafe utilities (optimized code for speed)                      |
/// +---------------------------------------------------------------------------+

// Entities only exist in the OID forms (128 bits)
// Any = Int + Float + *ClaireAny + list (*Vector<OID>) + set(*Set<OID>)
// ClaireAny is anything represented by a struct whose first slot is type
// ClaireAny = ClaireObject + ClaireBool + ClaireString + ClaireSymbol +
//             Tlist + Tset + TDict ???<what do we do with types>
type NativeInteger int
type NativeFloat float64
type NativeString *string

// note : string, list, set, etc... are all ClaireAny
// anything that can contain an any (slot or list) contains a *ClaireAny
// however, we optimize as soon as the sort is known

// hypothesis there are five sorts = five native types
// int, float64, string, rune and *ClaireAny

// Kernel imports Type functions that are defined elsewhere
// F_class_I_type (t) = best class
// T_contain_type (t,x)  = type membership

// +---------------------------------------------------------------------------+
// |  Part 1: description of EID                                               |
// +---------------------------------------------------------------------------+

// this is the new EID with three special markers
var C__INT *ClaireAny
var C__CHAR *ClaireAny
var C__FLOAT *ClaireAny

// EID is a struct used by the interpreter
type EID struct {
	PTR *ClaireAny // object, integer, char, float or exception
	VAL uint64
}

// three functions that would be macros in C
func OBJ(x EID) *ClaireAny { return x.PTR }
func INT(x EID) int        { return (int)(x.VAL) }
func FLOAT(x EID) float64  { return math.Float64frombits((uint64)(x.VAL)) }
func CHAR(x EID) rune      { return (rune)(x.VAL) }

// transforms a FLOAT/INTEGER/CHAR into a uint64 value
func IVAL(x int) uint64     { return uint64(x) }
func FVAL(x float64) uint64 { return math.Float64bits(x) }
func CVAL(x rune) uint64    { return uint64((int)(x)) }

func ANY(x EID) *ClaireAny {
	if x.PTR == C__INT {
		return MakeInteger(INT(x)).Id()
	} else if x.PTR == C__FLOAT {
		return MakeFloat(FLOAT(x)).Id()
	} else if x.PTR == C__CHAR {
		return MakeChar(CHAR(x)).Id()
	} else {
		return x.PTR
	}
}

//debug - to remove later - ensure that integer EID are not OID
// key for native integer method to work properly
func BAD(x EID) bool {
 return (x.PTR != C__INT && x.PTR != C__FLOAT && x.PTR.Isa == C_integer) }
func BadI(x EID, tag string) {
		if BAD(x) {
			panic("Bad Integer represented as OID in " + tag)} }

// owner class (x.Isa) for an EID			
func OWNER(x EID) *ClaireClass {
	if x.PTR == C__INT {
		return C_integer
	} else if x.PTR == C__FLOAT {
		return C_float
	} else if x.PTR == C__CHAR {
		return C_char
	} else {
		return x.PTR.Isa
	}
}

// the function types associated
type eFunc1 func(EID) EID
type eFunc2 func(EID, EID) EID
type eFunc3 func(EID, EID, EID) EID
type eFunc4 func(EID, EID, EID, EID) EID
type eFunc5 func(EID, EID, EID, EID, EID) EID
type eFunc6 func(EID, EID, EID, EID, EID, EID) EID
type eFunc7 func(EID, EID, EID, EID, EID, EID, EID) EID
type eFunc8 func(EID, EID, EID, EID, EID, EID, EID, EID) EID
type eFunc9 func(EID, EID, EID, EID, EID, EID, EID, EID, EID) EID
type eFunc10 func(EID, EID, EID, EID, EID, EID, EID, EID, EID, EID) EID

// used for evaluator function
type eFunc func(*ClaireAny) EID

// unsafe utilities
// let see if this works
func (p *ClaireAny) ui64() uint64    { return (uint64)(uintptr(unsafe.Pointer(p))) }
// func (p *ClaireAny) any() *ClaireAny { return p } // deprecated => Id()

// usefull utility get the pointer as a uintptr
// func (x *ClaireAny) Uip() uintptr { return uintptr(unsafe.Pointer(x)) }

// +---------------------------------------------------------------------------+
// |  Part 2: description of Claire Objects & Entities                         |
// +---------------------------------------------------------------------------+

// this is the root (object form)
type ClaireAny struct {
	Isa *ClaireClass
} // the heart of reflection in CLAIRE

// generic cast to Any : works for all
func (p *ClaireAny) ToAny() *ClaireAny { return p } // this is the new syntax
func (p *ClaireAny) Id() *ClaireAny    { return p } // this is the new syntax

// generic convert to proper EID  (the conversion into one of 3 EID type is done dynamically)
func (x *ClaireAny) ToEID() EID {
	c := x.Isa
	if c == C_integer {
		return EID{C__INT, IVAL(ToInteger(x).Value)}
	} else if c == C_float {
		return EID{C__FLOAT, FVAL(ToFloat(x).Value)}
	} else if c == C_char {
		return EID{C__CHAR, CVAL(ToChar(x).Value)}
	} else {
		return EID{x, 0}
	}
}

// ======================================= IMPORTED ==================================================
// imported is a root for objects that are defined with golang code (but with isa slot) ---------------
// integer, string, float, function, char, bool
type ClairePrimitive struct {
	ClaireAny
}
func ToPrimitive(x *ClaireAny) *ClairePrimitive { return (*ClairePrimitive)(unsafe.Pointer(x)) }

// integers have both an object form and a native form (int)
type ClaireInteger struct {
	ClairePrimitive
	Value int
}

// constructor for Claire Integer
func MakeInteger(n int) *ClaireInteger {
	var o *ClaireInteger = new(ClaireInteger)
	o.Isa = C_integer
	o.Value = n
	return o
}

// WARNING : cast ToX(..) must be functions because we cannot add methods
// cast (generic added for each class)
func ToInteger(x *ClaireAny) *ClaireInteger { return (*ClaireInteger)(unsafe.Pointer(x)) }

// AnyInteger : return a Claire Any (syntactic sugar)
func AnyInteger(n int) *ClaireAny { return (*ClaireAny)(unsafe.Pointer(MakeInteger(n))) }

func EVAL_integer(x *ClaireAny) EID { return EID{C__INT, IVAL(ToInteger(x).Value)} }

// floats have both an object form and a native form (int)
type ClaireFloat struct {
	ClairePrimitive
	Value float64
}

// constructor for Claire Float
func MakeFloat(x float64) *ClaireFloat {
	var o *ClaireFloat = new(ClaireFloat)
	o.Isa = C_float
	o.Value = x
	return o
}

func AnyFloat(x float64) *ClaireAny { return (*ClaireAny)(unsafe.Pointer(MakeFloat(x))) }

// cast (generic added for each class)
func ToFloat(x *ClaireAny) *ClaireFloat { return (*ClaireFloat)(unsafe.Pointer(x)) }

// eval a ClaireFloat
func EVAL_float(x *ClaireAny) EID { return EID{C__FLOAT, FVAL(ToFloat(x).Value)} }

// chars are imported runes (native import : functions)
type ClaireChar struct {
	ClairePrimitive
	Value rune
}

func ToChar(x *ClaireAny) *ClaireChar { return (*ClaireChar)(unsafe.Pointer(x)) }

// eval a ClaireChar
func EVAL_char(x *ClaireAny) EID { return EID{C__CHAR, CVAL(ToChar(x).Value)} }

// strings are Claire Object (hence they are mutable :))
// whereas symbols are native *string since they are immutable
type ClaireString struct {
	ClairePrimitive
	Value string
}

func ToString(x *ClaireAny) *ClaireString { return (*ClaireString)(unsafe.Pointer(x)) }

// Ports
// since class inheritance does not exists, a ClairePort is an "enum" : all the field
type ClairePort struct {
	ClairePrimitive
	file   *os.File // a ClairePort is a wrapper around a file
	buffer []byte   // that uses a buffer
	nChar  int      // number of character read in buffer  (used to be called index)
	total  int      // total read (multiple buffers)
	size   int      // size of file (> 0 => read, -1 => write)
	nEof   int      // sets at the end
	status int      // see ClString: 0,1 (buffered file), 2,3(write / read to string), 4 (read stream)
	firstc rune     // buffered first (read) char
}

func ToPort(x *ClaireAny) *ClairePort { return (*ClairePort)(unsafe.Pointer(x)) }

// extension with a GO reader (necessary to read from a stream like stdin)
type ClaireGoPort struct {
	ClairePort
	reader *bufio.Reader
}

func ToGoPort(x *ClaireAny) *ClaireGoPort { return (*ClaireGoPort)(unsafe.Pointer(x)) }

// EID functions are pointers with a name - foundation for interpreter
type ClaireFunction struct {
	ClairePrimitive
	name string
	arity int
}

func ToFunction(x *ClaireAny) *ClaireFunction { return (*ClaireFunction)(unsafe.Pointer(x)) }

// 10 variants with arity 1 to 10
type ClaireFunction1 struct {
	ClaireFunction
	call eFunc1
}

func toFunction1(f *ClaireFunction) *ClaireFunction1 { return ((*ClaireFunction1)(unsafe.Pointer(f))) }
func MakeFunction1(f eFunc1, name string) *ClaireFunction {
	o := new(ClaireFunction1)
	o.Isa = C_function
	o.call = f
	o.name = name
	o.arity = 1
	return (*ClaireFunction)(unsafe.Pointer(o))
}

type ClaireFunction2 struct {
	ClaireFunction
	call eFunc2
}

func toFunction2(f *ClaireFunction) *ClaireFunction2 { return ((*ClaireFunction2)(unsafe.Pointer(f))) }
func MakeFunction2(f eFunc2, name string) *ClaireFunction {
	o := new(ClaireFunction2)
	o.Isa = C_function
	o.call = f
	o.name = name
	o.arity = 2
	return (*ClaireFunction)(unsafe.Pointer(o))
}

type ClaireFunction3 struct {
	ClaireFunction
	call eFunc3
}

func toFunction3(f *ClaireFunction) *ClaireFunction3 { return ((*ClaireFunction3)(unsafe.Pointer(f))) }

func MakeFunction3(f eFunc3, name string) *ClaireFunction {
	o := new(ClaireFunction3)
	o.Isa = C_function
	o.call = f
	o.name = name
	o.arity = 3
	return (*ClaireFunction)(unsafe.Pointer(o))
}

type ClaireFunction4 struct {
	ClaireFunction
	call eFunc4
}

func MakeFunction4(f eFunc4, name string) *ClaireFunction {
	o := new(ClaireFunction4)
	o.Isa = C_function
	o.call = f
	o.name = name
	o.arity = 4
	return (*ClaireFunction)(unsafe.Pointer(o))
}

type ClaireFunction5 struct {
	ClaireFunction
	call eFunc5
}

func MakeFunction5(f eFunc5, name string) *ClaireFunction {
	o := new(ClaireFunction5)
	o.Isa = C_function
	o.call = f
	o.name = name
	o.arity = 5
	return (*ClaireFunction)(unsafe.Pointer(o))
}

type ClaireFunction6 struct {
	ClaireFunction
	call eFunc6
}

func MakeFunction6(f eFunc6, name string) *ClaireFunction {
	o := new(ClaireFunction6)
	o.Isa = C_function
	o.call = f
	o.name = name
	o.arity = 6
	return (*ClaireFunction)(unsafe.Pointer(o))
}

type ClaireFunction7 struct {
	ClaireFunction
	call eFunc7
}

func MakeFunction7(f eFunc7, name string) *ClaireFunction {
	o := new(ClaireFunction7)
	o.Isa = C_function
	o.call = f
	o.name = name
	o.arity = 7
	return (*ClaireFunction)(unsafe.Pointer(o))
}

// ==================================== OBJECTS  =============================================
// note: in CLAIRE 4 we merge ephemeral_object and (ephemeral) things
// we will use c.open to make sure that extension is kept
type ClaireObject struct {
	ClaireAny
}

func ToObject(x *ClaireAny) *ClaireObject { return (*ClaireObject)(unsafe.Pointer(x)) }

// claire system objects are objects managed by claire ---------------------------------------
type ClaireSystemObject struct {
	ClaireObject
}

func ToSystemObject(x *ClaireAny) *ClaireSystemObject {
	return (*ClaireSystemObject)(unsafe.Pointer(x))
}

// we create instructions and variables in Kernel for CLAIRE4 to be allowed to optimize
type ClaireInstruction struct {
	ClaireSystemObject
}

// notice the appartion on of _, when a class C starts with a capital letter, the cast function is To_C (separate ToList and To_List)
func To_Instruction(x *ClaireAny) *ClaireInstruction {
	return (*ClaireInstruction)(unsafe.Pointer(x))
}

// a variable is a low level object in the interpreter => put it in Kernel to support low-level optimization
type ClaireVariable struct {
	ClaireInstruction
	Pname *ClaireSymbol // name of the variable
	Range *ClaireType   // type of the variable content
	Index int           // position in the stack
}

func To_Variable(x *ClaireAny) *ClaireVariable {
	return (*ClaireVariable)(unsafe.Pointer(x))
}

// a ClaireSymbol is a name(string) in a namespace (module)
// we keep in memory where the module was defined (NULL means private)
type ClaireSymbol struct {
	ClaireSystemObject
	module_I   *ClaireModule // module m that owns the symbol m/name
	key        string        // name that defines the symbol in the m namepace
	value      *ClaireAny    // private -> accessed through Value()
	definition *ClaireModule // private : where the symbol was defined
}

func ToSymbol(x *ClaireAny) *ClaireSymbol { return (*ClaireSymbol)(unsafe.Pointer(x)) }

// constructor : local for bootstrap and global
func makeSymbol(s string) *ClaireSymbol { return MakeSymbol(s, C_claire) }
func MakeSymbol(s string, m *ClaireModule) *ClaireSymbol {
	var o *ClaireSymbol = new(ClaireSymbol)
	o.Isa = C_symbol
	o.module_I = m
	o.key = s
	o.definition = ClEnv.Module_I // where the symbol is defined
	return o
}

// unbound symbols are useful : wrapper around symbol
type ClaireUnboundSymbol struct {
	ClaireSystemObject
	Name *ClaireSymbol
}

func ToUnboundSymbol(x *ClaireAny) *ClaireUnboundSymbol {
	return (*ClaireUnboundSymbol)(unsafe.Pointer(x))
}

// constructor
func MakeUnboundSymbol(s *ClaireSymbol) *ClaireUnboundSymbol {
	o := new(ClaireUnboundSymbol)
	o.Isa = C_unbound_symbol
	o.Name = s
	return o
}

// Exception are objects - since there are no try / catch, we will make sure that
// errors answers are caught + passed straight to the top
type ClaireException struct {
	ClaireSystemObject
}

func ToException(x *ClaireAny) *ClaireException { return (*ClaireException)(unsafe.Pointer(x)) }

// errors are exceptions (assumption EID => EID(ClaireException,e) )
type ClaireError struct {
	ClaireException
}

func ToError(x *ClaireAny) *ClaireError { return (*ClaireError)(unsafe.Pointer(x)) }

// system error are produced in the go microClaire kernel
type ClaireSystemError struct {
	ClaireError
	Index int
	Value *ClaireAny
	Arg   *ClaireAny
}

func ToSystemError(x *ClaireAny) *ClaireSystemError { return (*ClaireSystemError)(unsafe.Pointer(x)) }

// things are objects with a name  (that will evolve when ClaireSymbol are introduced)
type ClaireThing struct {
	ClaireObject
	Name *ClaireSymbol // symbol is a string + a module
}

func ToThing(x *ClaireAny) *ClaireThing { return (*ClaireThing)(unsafe.Pointer(x)) }

// conversion function for ClaireEnv (defined in clEnv)
func ToEnvironment(x *ClaireAny) *ClaireEnvironment { return (*ClaireEnvironment)(unsafe.Pointer(x)) }

// Boolean are things
type ClaireBoolean struct {
	ClaireThing
	Not *ClaireBoolean
}

// constructor
func MakeBoolean(b bool) *ClaireBoolean {
	if b {
		return CTRUE
	} else {
		return CFALSE
	}
}
func ToBoolean(x *ClaireAny) *ClaireBoolean { return (*ClaireBoolean)(unsafe.Pointer(x)) }

// should be a standard in go !
func IfThenElse(b bool, a1 *ClaireAny, a2 *ClaireAny) *ClaireAny {
	if b {
		return a1
	} else {
		return a2
	}
}

// test boolean on EID (strict bool) - no conversion necessary
func ETRUE(x EID) bool { return x.PTR == CTRUE.Id() }

// a collection can be iterated : root of types (including classes and sets) and lists, tuple
type ClaireCollection struct {
	ClaireObject
}

func ToCollection(x *ClaireAny) *ClaireCollection { return (*ClaireCollection)(unsafe.Pointer(x)) }


// in CLAIRE 4, a type_expression is a "pseudo type" -> can be used in method definition
type ClaireTypeExpression struct {
	ClaireCollection
}

func ToTypeExpression(x *ClaireAny) *ClaireTypeExpression {
	return (*ClaireTypeExpression)(unsafe.Pointer(x))
}

// a root for all types : Class | Union | Interval | Finterval | ...
type ClaireType struct {
	ClaireTypeExpression
}

func ToType(x *ClaireAny) *ClaireType { return (*ClaireType)(unsafe.Pointer(x)) }

// a class
// v4: note that ClaireList is the only type that use / specialization (Obj/Int/Float) is done dynamically
type ClaireClass struct {
	ClaireType                 // a class is a type and a thing (with a name)
	Name        *ClaireSymbol  // name is a symbol
	Comment     *ClaireString  // comment that tells what the class is
	Slots       *ClaireList    // list of slots (no prototype, use slots)
	Superclass  *ClaireClass   // super class
	Subclass    *ClaireSet    // subclasses
	Ancestors   *ClaireList    // list of Ancestors
	Descendents *ClaireSet     // descendants
	Open        int            // open status for class
	Instances   *ClaireList    // instance list is only kept when open
	Params      *ClaireList    // for parameterized classes (subset of slots)
	Dictionary  *ClaireMap     // map of functions (when p.dictionary is true => faster than using the list)
	Ident_ask   *ClaireBoolean // equal = eq (identified objects)
	IfWrite     *ClaireAny
	evaluate    eFunc // evaluate function (private)
}

func ToClass(x *ClaireAny) *ClaireClass { return (*ClaireClass)(unsafe.Pointer(x)) }

// evaluation (identity) for objects
func EVAL_object(x *ClaireAny) EID { return EID{x, 0} }

// a class is a thing and a type
// heart of evaluator
func EVAL(x *ClaireAny) EID { 
	/*y := x.Isa.evaluate(x)
	if BAD(y) {fmt.Printf("=== eval(%s) produces a bad Int ==\n",x.Prt())}
    return y} */
	return x.Isa.evaluate(x) }

// intermediate
type ClaireSystemThing struct {
	ClaireThing
}

// keywords
type ClaireKeyword struct {
	ClaireSystemThing
}

func ToKeyword(x *ClaireAny) *ClaireKeyword { return (*ClaireKeyword)(unsafe.Pointer(x)) }

// relation = table or property
type ClaireRelation struct {
	ClaireSystemThing
	Comment         *ClaireString // comment that tells what the class is
	Domain          *ClaireType   // union of all classes with a restriction from p
	Range           *ClaireType   // range type tn
	IfWrite         *ClaireAny
	Store_ask       *ClaireBoolean  // stored (defeasible) updates for slots
	Inverse         *ClaireRelation // for binary relations (slots)
	Open            int
	Multivalued_ask *ClaireBoolean // set-based relation
}

func ToRelation(x *ClaireAny) *ClaireRelation { return (*ClaireRelation)(unsafe.Pointer(x)) }

// properties are selector (sets of methods)
type ClaireProperty struct {
	ClaireRelation
	Trace_I      int            // do we want to trace
	Restrictions *ClaireList    // list of methods and slots
	Definition   *ClaireList    // list of methods, in the proper order
	Dictionary   *ClaireBoolean // do we us a dictionary (stored at the class level) when the property is uniform ?
	Reified      *ClaireBoolean // o.x is an object with a value slot
}

func ToProperty(x *ClaireAny) *ClaireProperty { return (*ClaireProperty)(unsafe.Pointer(x)) }

// a table is a relation
type ClaireOperation struct {
	ClaireProperty
	Precedence int // order
}

func ToOperation(x *ClaireAny) *ClaireOperation { return (*ClaireOperation)(unsafe.Pointer(x)) }

// a table is a relation
type ClaireTable struct {
	ClaireRelation
	Graph   *ClaireAny // list or Dictionary
	Params  *ClaireAny // int, list(int,int) or any
	Default *ClaireAny // default value
}

func ToTable(x *ClaireAny) *ClaireTable { return (*ClaireTable)(unsafe.Pointer(x)) }

// a restriction is a slot or a method
type ClaireRestriction struct {
	ClaireAny
	Module_I *ClaireModule   // module where the restriction is defined
	Comment  *ClaireString   // comment that tells what the restriction is
	Domain   *ClaireList     // list of types (multiple args : domain is cartesian product)
	Range    *ClaireType     // range of method / slot
	Selector *ClaireProperty // property that is defined by this restriction
}

func ToRestriction(x *ClaireAny) *ClaireRestriction { return (*ClaireRestriction)(unsafe.Pointer(x)) }

// a slot has an index and default value
type ClaireSlot struct {
	ClaireRestriction
	Srange  *ClaireClass // sort range (class) for fast handling
	Default *ClaireAny
	Index   int
}

func ToSlot(x *ClaireAny) *ClaireSlot { return (*ClaireSlot)(unsafe.Pointer(x)) }

// Methods - need the arity because of golang
// todo: add module
type ClaireMethod struct {
	ClaireRestriction
	Srange     *ClaireList // list of class types (ex-sorts), where Domain is a list of types
	Formula    *ClaireLambda
	Functional *ClaireFunction
	Typing     *ClaireAny // can be a property, a function or a lambda !
	Status     int        // used to store if the method may raise an error
	Inline_ask *ClaireBoolean
}

func ToMethod(x *ClaireAny) *ClaireMethod { return (*ClaireMethod)(unsafe.Pointer(x)) }

// ------------------------- Claire Lists ---------------------------------------------------

// we recreate bags as a common root (bags are types so that set are types because of single inheritance)
// CLAIRE combines read-only dynamically typed list/sets and mutable statically typed list/sets [of = {} <=> unmutable]
type ClaireBag struct {
	ClaireType
	of     *ClaireType          // type that contains all list members if mutable
}

func ToBag(x *ClaireAny) *ClaireBag { return (*ClaireBag)(unsafe.Pointer(x)) }

// ClaireList is just a common root (what is seen in claire)
// actual implementation is one of the three next types
type ClaireList struct {
	ClaireBag
	Srange *ClaireClass // sort : object, float or integer
}

func ToList(x *ClaireAny) *ClaireList { return (*ClaireList)(unsafe.Pointer(x)) }

// 3 subtypes (internal to golang) - we
type ClaireListObject struct {
	ClaireList
	Values []*ClaireAny
}

// how to access directy the values field
func (x *ClaireList) ValuesO() []*ClaireAny { return ((*ClaireListObject)(unsafe.Pointer(x))).Values }
func (x *ClaireList) ValuesI() []int        { return ((*ClaireListInteger)(unsafe.Pointer(x))).Values }
func (x *ClaireList) ValuesF() []float64    { return ((*ClaireListFloat)(unsafe.Pointer(x))).Values }

// these cast are specific -> defined as methods (private to Kernel)
func (x *ClaireAny) toObject() *ClaireListObject { return (*ClaireListObject)(unsafe.Pointer(x)) }

type ClaireListInteger struct {
	ClaireList
	Values []int
}

func (x *ClaireAny) toInteger() *ClaireListInteger {
	return (*ClaireListInteger)(unsafe.Pointer(x))
}

type ClaireListFloat struct {
	ClaireList
	Values []float64
}

func (x *ClaireAny) toFloat() *ClaireListFloat { return (*ClaireListFloat)(unsafe.Pointer(x)) }

// listargs is actually an alias
type ClaireListargs ClaireList

// key decision : Arrays are just fixed size lists => implemented with a *ClaireList whose isa is C_array.
// only the instanciation changes. In the future, add a new type <type>[n] to capture the length
func ToArray(x *ClaireAny) *ClaireList { return (*ClaireList)(unsafe.Pointer(x)) }

// key decision : ClaireTuple is ClaireList (from go point of view)
// the difference is that Isa contains C_tuple
type ClaireTuple = ClaireList

func ToTuple(x *ClaireAny) *ClaireTuple { return (*ClaireList)(unsafe.Pointer(x)) }

// sets are implemented with sorted maps - hence they are clone from lists
type ClaireSet struct {
		ClaireBag
		Srange *ClaireClass // sort : object, float or integer
		Count int
}

func ToSet(x *ClaireAny) *ClaireSet { return (*ClaireSet)(unsafe.Pointer(x)) }

// we actually need 3 subclasses from ClaireListNative ....
// count is accessed with Size()
type ClaireSetObject struct {
	ClaireSet
	Values []*ClaireAny
}
func ToSetObject(x *ClaireAny) *ClaireSetObject { return (*ClaireSetObject)(unsafe.Pointer(x)) }

type ClaireSetInteger struct {
	ClaireSet
	Values []int
}
func ToSetInteger(x *ClaireAny) *ClaireSetInteger { return (*ClaireSetInteger)(unsafe.Pointer(x)) }

type ClaireSetFloat struct {
	ClaireSet
	Values []float64
}
func ToSetFloat(x *ClaireAny) *ClaireSetFloat { return (*ClaireSetFloat)(unsafe.Pointer(x)) }

// how to access directy the values field
func (x *ClaireSet) ValuesO() []*ClaireAny { return ((*ClaireSetObject)(unsafe.Pointer(x))).Values }
func (x *ClaireSet) ValuesI() []int        { return ((*ClaireSetInteger)(unsafe.Pointer(x))).Values }
func (x *ClaireSet) ValuesF() []float64    { return ((*ClaireSetFloat)(unsafe.Pointer(x))).Values }

// lambda
type ClaireLambda struct {
	ClaireSystemObject
	Vars      *ClaireList
	Body      *ClaireAny
	Dimension int // needed in claire (actually length of Vars ?)
}

func ToLambda(x *ClaireAny) *ClaireLambda { return (*ClaireLambda)(unsafe.Pointer(x)) }

// constructor
func MakeLambda(l *ClaireList, a *ClaireAny) *ClaireLambda {
	var o = new(ClaireLambda)
	o.Isa = C_lambda
	o.Vars = l
	o.Body = a
	return o
}

// modules are things with a special structure
// the symbol table is the last slot (unvisible to claire)
type ClaireModule struct {
	ClaireSystemThing
	Comment  *ClaireString            // each module has a comment !
	Parts    *ClaireList              // list of submodules
	PartOf   *ClaireModule            // sub-module of
	Uses     *ClaireList              // other modules that are used
	Source   *ClaireString            // directory where the sources can be found
	MadeOf   *ClaireList              // listf of file names
	Status   int                      // new (0:default, 1:loaded, 2 compiled, 3:c+loaded, 4:c+l+traced)
	External *ClaireString            // external of the module: .lib library
	Evaluate *ClaireFunction          // load function produced by compiler
	table    map[string]*ClaireSymbol // each mod has a symbol table (private)
}

func ToModule(x *ClaireAny) *ClaireModule { return (*ClaireModule)(unsafe.Pointer(x)) }

// dictionaries are called map in reference to go
type ClaireMap struct {
	ClaireCollection
	of    *ClaireType           // range for entry
	Range *ClaireType           // range for content
	Value map[string]*ClaireAny // leverage go maps + same Key functions as sets
}

func ToMap(x *ClaireAny) *ClaireMap { return (*ClaireMap)(unsafe.Pointer(x)) }

// new in CLAIRE 4: Types move to Kernel for functional closure (contains and includes) -----------------------------------------------------

// root for type operator (used to be Type -> type_operatore )
// this is the type lattice (not extensible), defined in Kernel  
// extensions such as Patterns or References are type_expressions
type ClaireTypeOperator struct {
	ClaireType
}

func ToTypeOperator(x *ClaireAny) *ClaireTypeOperator {
	return (*ClaireTypeOperator)(unsafe.Pointer(x))
}

// Union of two types
type ClaireUnion struct {
	ClaireTypeOperator
	T1 *ClaireType
	T2 *ClaireType
}

func To_Union(x *ClaireAny) *ClaireUnion { return (*ClaireUnion)(unsafe.Pointer(x)) }

// interval (only int because we want enumerability + this is useful for array/ list range inference for indices)
type ClaireInterval struct {
	ClaireTypeOperator
	Arg1 int
	Arg2 int
}

func To_Interval(x *ClaireAny) *ClaireInterval { return (*ClaireInterval)(unsafe.Pointer(x)) }

// a paramerized subclass
type ClaireParam struct {
	ClaireTypeOperator
	Arg    *ClaireClass
	Params *ClaireList
	Args   *ClaireList
}

func To_Param(x *ClaireAny) *ClaireParam { return (*ClaireParam)(unsafe.Pointer(x)) }

// a generic subtype C[B] : subtype of class C whose members belong to B, e.g., list[integer]
type ClaireSubtype struct {
	ClaireTypeOperator
	Arg *ClaireClass
	T1  *ClaireType
}

func ToSubtype(x *ClaireAny) *ClaireSubtype { return (*ClaireSubtype)(unsafe.Pointer(x)) }

// reference to a previous variable, not a type but a type_expression, like a pattern -------
// index is the position of the stack of the referred type
// args is a list representing the path (a sequence of properties (parameters))
// a property is applied to the referred type
// if arg = true, the reference is the singleton containing the ref. value
type ClaireReference struct {
	ClaireTypeExpression
	Args  *ClaireList              // 
	Index int
	Arg   *ClaireBoolean           // usually false
}

// arg:boolean = false)
func To_Reference(x *ClaireAny) *ClaireReference { return (*ClaireReference)(unsafe.Pointer(x)) }


// some global variables
var claireStdout *ClairePort
var claireStdin *ClairePort
var it *ClaireModule = nil
var C_class *ClaireClass = nil
var C_void *ClaireClass = nil
var C_any *ClaireClass = nil
var C_object *ClaireClass = nil
var C_primitive *ClaireClass = nil
var C_system_object *ClaireClass = nil
var C_Instruction *ClaireClass = nil
var C_Variable *ClaireClass = nil
var C_environnement *ClaireClass = nil
var C_exception *ClaireClass = nil
var C_error *ClaireClass = nil
var CERROR *ClaireError = nil
var C_system_error *ClaireClass = nil
var C_environment *ClaireClass = nil
var C_function *ClaireClass = nil
var C_collection *ClaireClass = nil
var C_type *ClaireClass = nil
var C_thing *ClaireClass = nil
var C_string *ClaireClass = nil
var C_integer *ClaireClass = nil
var C_float *ClaireClass = nil
var C_char *ClaireClass = nil
var C_symbol *ClaireClass = nil
var C_unbound_symbol *ClaireClass = nil
var C_boolean *ClaireClass = nil
var CTRUE *ClaireBoolean = nil
var CFALSE *ClaireBoolean = nil
var CNULL *ClaireAny = nil  // unknown object
var unknownName *ClaireSymbol = nil
var PRIVATE *ClaireSymbol = nil
var CEMPTY *ClaireSet = nil // empty set
var CNIL *ClaireList = nil  // empty list
var EVOID EID               // EID handler on void, CNULL
var Niet EID                // a dumb variable that we can write onto :)
var CEOF rune
var C_system_thing *ClaireClass = nil
var C_keyword *ClaireClass = nil
var C_relation *ClaireClass = nil
var C_property *ClaireClass = nil
var C_operation *ClaireClass = nil
var C_table *ClaireClass = nil
var C_restriction *ClaireClass = nil
var C_method *ClaireClass = nil
var C_slot *ClaireClass = nil
var C_bag *ClaireClass = nil
var C_list *ClaireClass = nil
var C_listargs *ClaireClass = nil
var C_lambda *ClaireClass = nil
var C_set *ClaireClass = nil
var C_tuple *ClaireClass = nil
var C_array *ClaireClass = nil
var C_module *ClaireClass = nil
var C_claire *ClaireModule = nil
var C_mClaire *ClaireModule = nil
var C_Kernel *ClaireModule = nil
var C_port *ClaireClass = nil
var C_map *ClaireClass = nil
var C_type_expression *ClaireClass = nil
var C_type_operator *ClaireClass = nil
var C_Union *ClaireClass = nil
var C_Interval *ClaireClass = nil
var C_Param *ClaireClass = nil
var C_subtype *ClaireClass = nil
var C_Reference *ClaireClass = nil

// property vars
var C_copy *ClaireProperty
var C_empty *ClaireProperty
var C_length *ClaireProperty
var C_contain_ask *ClaireProperty
var C__in *ClaireProperty          // % in claire is _in in go
var C_included_ask *ClaireProperty // <= @ type : extensible through included?
var C_of *ClaireProperty
var C_isa *ClaireProperty
var C_mClaire_index *ClaireProperty
var C_value *ClaireProperty
var C_arg *ClaireProperty
var C_name *ClaireProperty
var C_mClaire_pname *ClaireProperty
var C_comment *ClaireProperty
var C_slots *ClaireProperty
var C_superclass *ClaireProperty
var C_ancestors *ClaireProperty
var C_subclass *ClaireProperty
var C_descendents *ClaireProperty
var C_open *ClaireProperty
var C_instances *ClaireProperty
var C_params *ClaireProperty
var C_mClaire_graph *ClaireProperty
var C_if_write *ClaireProperty
var C_dictionary *ClaireProperty
var C_ident_ask *ClaireProperty
var C_domain *ClaireProperty
var C_range *ClaireProperty
var C_store_ask *ClaireProperty
var C_inverse *ClaireProperty
var C_multivalued_ask *ClaireProperty
var C_restrictions *ClaireProperty
var C_mClaire_definition *ClaireProperty
var C_reified *ClaireProperty
var C_module_I *ClaireProperty
var C_trace_I *ClaireProperty
var C_selector *ClaireProperty
var C_mClaire_srange *ClaireProperty
var C_Kernel_typing *ClaireProperty
var C_default *ClaireProperty
var C_functional *ClaireProperty
var C_formula *ClaireProperty
var C_vars *ClaireProperty
var C_body *ClaireProperty
var C_dimension *ClaireProperty
var C_parts *ClaireProperty
var C_part_of *ClaireProperty
var C_uses *ClaireProperty
var C_source *ClaireProperty
var C_made_of *ClaireProperty
var C_mClaire_status *ClaireProperty
var C_external *ClaireProperty
var C_get *ClaireProperty
// var C_getenv *ClaireProperty   in Core
var C_put *ClaireProperty
var C_funcall *ClaireProperty
var C_fastcall *ClaireProperty
var C_nth *ClaireProperty
var C_nth_equal *ClaireProperty
var C_nth_put *ClaireProperty
var C_nth_plus *ClaireProperty
var C_nth_dash *ClaireProperty
var C_nth_get *ClaireProperty
var C_self_print *ClaireProperty
var C_princ *ClaireProperty
var C_close *ClaireProperty
var C_ephemeral *ClaireProperty
var C_final *ClaireProperty
var C_abstract *ClaireProperty
var C_jito_ask *ClaireProperty
var C_n_line *ClaireProperty
var C_gensym *ClaireProperty
var C_store *ClaireProperty
var C_commit *ClaireProperty
var C_backtrack *ClaireProperty
var C_choice *ClaireProperty
var C_symbol_I *ClaireProperty
var C_make_string *ClaireProperty
var C_make_array *ClaireProperty
var C_random *ClaireProperty
var C_string_I *ClaireProperty
var C_set_I *ClaireProperty
var C_integer_I *ClaireProperty
var C_float_I *ClaireProperty
var C_make_list *ClaireProperty
var C_array_I *ClaireProperty
var C_list_I *ClaireProperty
var C_class_I *ClaireProperty
var C_new *ClaireProperty
var C_mClaire_new_I *ClaireProperty
var C_make_function *ClaireProperty
var C_cdr *ClaireProperty
var C_skip *ClaireProperty
var C_shrink *ClaireProperty
var C_size *ClaireProperty
var C_cast_I *ClaireProperty
var C_tuple_I *ClaireProperty
var C_c_princ *ClaireProperty
var C_substring *ClaireProperty
var C_included *ClaireProperty
var C_begin *ClaireProperty
var C_end *ClaireProperty
var C_port_I *ClaireProperty
var C_use_as_output *ClaireProperty
var C_precedence *ClaireProperty
var C_shell *ClaireProperty
// var C_getenv *ClaireProperty
var C_fclose *ClaireProperty
var C_world_ask *ClaireProperty
var C_world_id *ClaireProperty
var C_set_length *ClaireProperty
var C_exit *ClaireProperty
var C_graph_get *ClaireProperty
var C_graph_put *ClaireProperty
var C_graph_init *ClaireProperty
var C_map_I *ClaireProperty
var C_boolean_I *ClaireProperty
var C_dict_get *ClaireProperty
var C_dict_put *ClaireProperty
var C_read *ClaireProperty
var C_read_ident *ClaireProperty
var C_read_number *ClaireProperty
var C_read_thing *ClaireProperty
var C_read_string *ClaireProperty
var C_print *ClaireProperty
var C_log *ClaireProperty
var C_cos *ClaireProperty
var C_sin *ClaireProperty
var C_atan *ClaireProperty
var C_sqrt *ClaireProperty
var C__exp2 *ClaireProperty
var C_stack_apply *ClaireProperty
var C_mClaire_t1 *ClaireProperty
var C_mClaire_t2 *ClaireProperty
var C_arg1 *ClaireProperty
var C_arg2 *ClaireProperty
var C_args *ClaireProperty
var C_inline_ask *ClaireProperty
var C_verbose *ClaireProperty
var C_exception_I *ClaireProperty
var C_version *ClaireProperty
var C_ctrace *ClaireProperty
var C_cout *ClaireProperty
var C_cin *ClaireProperty
var C_base *ClaireProperty
var C_mClaire_restore_state *ClaireProperty
var C_abort *ClaireProperty
var C_debug_I *ClaireProperty
// var C_step_I *ClaireProperty
var C_last_debug *ClaireProperty
var C_last_index *ClaireProperty
var C_spy_I *ClaireProperty
var C_defined *ClaireProperty
var C_fopen *ClaireProperty
var C_empty_list *ClaireProperty
var C_empty_set *ClaireProperty
var C_write_fast *ClaireProperty
var C_slot_get *ClaireProperty
var C_putc *ClaireProperty
var C_getc *ClaireProperty
var C_namespace *ClaireProperty
var C_date_I *ClaireProperty
var C_sort_I *ClaireProperty
var C_apply *ClaireProperty
var C_count_call *ClaireProperty
var C_count_level *ClaireProperty
var C_count_trigger *ClaireProperty
var C_add_slot *ClaireProperty
var C_add_method *ClaireProperty
var C_arity *ClaireProperty
var C_set_arity *ClaireProperty
var C_flush *ClaireProperty
var C_imports *ClaireProperty      // new in CLAIRE 4: pragma for modules
	
// operations
var C_add *ClaireOperation
var C_add_I *ClaireOperation
var C_add_star *ClaireOperation
var C_delete *ClaireOperation
var C__equal *ClaireOperation
var C__dash *ClaireOperation
var C__star *ClaireOperation
var C__7 *ClaireOperation
var C__7_plus *ClaireOperation
var C__Z *ClaireOperation
var C__exp *ClaireOperation
var C__inf *ClaireOperation
var C__inf_equal *ClaireOperation
var C__sup *ClaireOperation
var C__sup_equal *ClaireOperation
var C__dot_dot *ClaireOperation
var C_min *ClaireOperation
var C_max *ClaireOperation
var C_mod *ClaireOperation
var C_cons *ClaireOperation

// +---------------------------------------------------------------------------+
// |  Part 3: Dummy Classes for instanciation / get/ set                       |
// +---------------------------------------------------------------------------+

// THIS IS CRITICAL : WE CANNOT FOOL GOLAND COMPILER WITH POINTERS WHEN DOING A SET
// OTHERWISE WE CREATE GC ERRORS (OF THE UGLY KIND)
// HENCE WE HAVE A DUMMY OBJECT CLASS THAT WE USE FOR INDIRECT SLOT READ/WRITE
// THE POINTER STUFF WORKS FOR INTEGER (FASTER + AVOID COMPILER GC MARKS)

// consequence: Claire Object slots are either *Clany or direct 64bits values (float or int)

// this is a dummy hierarchy to support object up to size 30
type ClaireDummy1 struct {
	ClaireAny
	a2 *ClaireAny
	a3 *ClaireAny
	a4 *ClaireAny
}

type ClaireDummy2 struct {
	ClaireDummy1
	a5 *ClaireAny
	a6 *ClaireAny
	a7 *ClaireAny
	a8 *ClaireAny
	a9 *ClaireAny
}

type ClaireDummy3 struct {
	ClaireDummy2
	a10 *ClaireAny
	a11 *ClaireAny
	a12 *ClaireAny
	a13 *ClaireAny
	a14 *ClaireAny
}

type ClaireDummy4 struct {
	ClaireDummy3
	a15 *ClaireAny
	a16 *ClaireAny
	a17 *ClaireAny
	a18 *ClaireAny
	a19 *ClaireAny
}

type ClaireDummy5 struct {
	ClaireDummy4
	a20 *ClaireAny
	a21 *ClaireAny
	a22 *ClaireAny
	a23 *ClaireAny
	a24 *ClaireAny
}

// for the time being we stop at 30 slots.
type ClaireDummy6 struct {
	ClaireDummy5
	a25 *ClaireAny
	a26 *ClaireAny
	a27 *ClaireAny
	a28 *ClaireAny
	a29 *ClaireAny
}


// constructors : generic "make object" using c.slots to know the size 
// HENCE IT DOES NOT WORK WITH CLASSES or objects with hidden go slots
func (c *ClaireClass) makeObject() *ClaireObject {
	n := c.Slots.Length()
	var o *ClaireObject
	if n < 5 {
		o = (*ClaireObject)(unsafe.Pointer(new(ClaireDummy1)))
	} else if n < 10 {
		o = (*ClaireObject)(unsafe.Pointer(new(ClaireDummy2)))
	} else if n < 15 {
		o = (*ClaireObject)(unsafe.Pointer(new(ClaireDummy3)))
	} else if n < 20 {
		o = (*ClaireObject)(unsafe.Pointer(new(ClaireDummy4)))
	} else if n < 25 {
		o = (*ClaireObject)(unsafe.Pointer(new(ClaireDummy5)))
	} else if n < 30 {
		o = (*ClaireObject)(unsafe.Pointer(new(ClaireDummy6)))
	} else {
		panic(fmt.Sprintf("object of size %d is too big for claire ", n))
	}
	return o
}

// generic access to the n-th slot with the srange (s) info
func (x *ClaireObject) Get(n int, s *ClaireClass) *ClaireAny {
	if s == C_integer {
		// fmt.Printf("Get Int %d -> %d\n", n, x.GetInt(n))
		return MakeInteger(x.GetInt(n)).Id()
	} else if s == C_float {
		return MakeFloat(x.GetFloat(n)).Id()
	} else {
		return x.GetObj(n)
	}
}

// specific function when the srange is an object - this needs to be extended to 30
// since it looks that direct acess is dangerous we have this crazy code for interpreter
func (x *ClaireObject) GetObj(i int) *ClaireAny {
	if i == 1 {
		return ((*ClaireDummy1)(unsafe.Pointer(x))).Isa.Id()
	} else if i == 2 {
		return ((*ClaireDummy1)(unsafe.Pointer(x))).a2
	} else if i == 3 {
		return ((*ClaireDummy1)(unsafe.Pointer(x))).a3
	} else if i == 4 {
		return ((*ClaireDummy1)(unsafe.Pointer(x))).a4
	} else if i == 5 {
		return ((*ClaireDummy2)(unsafe.Pointer(x))).a5
	} else if i == 6 {
		return ((*ClaireDummy2)(unsafe.Pointer(x))).a6
	} else if i == 7 {
		return ((*ClaireDummy2)(unsafe.Pointer(x))).a7
	} else if i == 8 {
		return ((*ClaireDummy2)(unsafe.Pointer(x))).a8
	} else if i == 9 {
		return ((*ClaireDummy2)(unsafe.Pointer(x))).a9
	} else if i == 10 {
		return ((*ClaireDummy3)(unsafe.Pointer(x))).a10
	} else if i == 11 {
		return ((*ClaireDummy3)(unsafe.Pointer(x))).a11
	} else if i == 12 {
		return ((*ClaireDummy3)(unsafe.Pointer(x))).a12
	} else if i == 13 {
		return ((*ClaireDummy3)(unsafe.Pointer(x))).a13
	} else if i == 14 {
		return ((*ClaireDummy3)(unsafe.Pointer(x))).a14
	} else if i == 15 {
		return ((*ClaireDummy4)(unsafe.Pointer(x))).a15
	} else if i == 16 {
		return ((*ClaireDummy4)(unsafe.Pointer(x))).a16
	} else if i == 17 {
		return ((*ClaireDummy4)(unsafe.Pointer(x))).a17
	} else if i == 18 {
		return ((*ClaireDummy4)(unsafe.Pointer(x))).a18
	} else if i == 19 {
		return ((*ClaireDummy4)(unsafe.Pointer(x))).a19
	} else if i == 20 {
		return ((*ClaireDummy5)(unsafe.Pointer(x))).a20
	} else if i == 21 {
		return ((*ClaireDummy5)(unsafe.Pointer(x))).a21
	} else if i == 22 {
		return ((*ClaireDummy5)(unsafe.Pointer(x))).a22
	} else if i == 23 {
		return ((*ClaireDummy5)(unsafe.Pointer(x))).a23
	} else if i == 24 {
		return ((*ClaireDummy5)(unsafe.Pointer(x))).a24
	} else if i == 25 {
		return ((*ClaireDummy6)(unsafe.Pointer(x))).a25
	} else if i == 26 {
		return ((*ClaireDummy6)(unsafe.Pointer(x))).a26
	} else if i == 27 {
		return ((*ClaireDummy6)(unsafe.Pointer(x))).a27
	} else if i == 28 {
		return ((*ClaireDummy6)(unsafe.Pointer(x))).a28
	} else if i == 29 {
		return ((*ClaireDummy6)(unsafe.Pointer(x))).a29
	} else {
		panic("Fatal error with getObj (i too big >= 30)")
		return C_class.Id()
	} // need a Cerror !
}

// symetric: generic set
func (x *ClaireObject) Set(n int, s *ClaireClass, y *ClaireAny) {
	if s == C_integer {
		s.SetInt(n, ToInteger(y).Value)
	} else if s == C_float {
		s.SetFloat(n, ToFloat(y).Value)
	} else {
		x.SetObj(n, y)
	}
}

// This is the key method. set needs to use a golang object assign for GC to do its
// wonder. direct pointer handling (cf. integer code) fails ! (badly, I must add)
func (x *ClaireObject) SetObj(i int, y *ClaireAny) {
	if i == 1 {
		((*ClaireDummy1)(unsafe.Pointer(x))).Isa = ToClass(y)
	} else if i == 2 {
		((*ClaireDummy1)(unsafe.Pointer(x))).a2 = y
	} else if i == 3 {
		((*ClaireDummy1)(unsafe.Pointer(x))).a3 = y
	} else if i == 4 {
		((*ClaireDummy1)(unsafe.Pointer(x))).a4 = y
	} else if i == 5 {
		((*ClaireDummy2)(unsafe.Pointer(x))).a5 = y
	} else if i == 6 {
		((*ClaireDummy2)(unsafe.Pointer(x))).a6 = y
	} else if i == 7 {
		((*ClaireDummy2)(unsafe.Pointer(x))).a7 = y
	} else if i == 8 {
		((*ClaireDummy2)(unsafe.Pointer(x))).a8 = y
	} else if i == 9 {
		((*ClaireDummy2)(unsafe.Pointer(x))).a9 = y
	} else if i == 10 {
		((*ClaireDummy3)(unsafe.Pointer(x))).a10 = y
	} else if i == 11 {
		((*ClaireDummy3)(unsafe.Pointer(x))).a11 = y
	} else if i == 12 {
		((*ClaireDummy3)(unsafe.Pointer(x))).a12 = y
	} else if i == 13 {
		((*ClaireDummy3)(unsafe.Pointer(x))).a13 = y
	} else if i == 14 {
		((*ClaireDummy3)(unsafe.Pointer(x))).a14 = y
	} else if i == 15 {
		((*ClaireDummy4)(unsafe.Pointer(x))).a15 = y
	} else if i == 16 {
		((*ClaireDummy4)(unsafe.Pointer(x))).a16 = y
	} else if i == 17 {
		((*ClaireDummy4)(unsafe.Pointer(x))).a17 = y
	} else if i == 18 {
		((*ClaireDummy4)(unsafe.Pointer(x))).a18 = y
	} else if i == 19 {
		((*ClaireDummy4)(unsafe.Pointer(x))).a19 = y
	} else if i == 20 {
		((*ClaireDummy5)(unsafe.Pointer(x))).a20 = y
	} else if i == 21 {
		((*ClaireDummy5)(unsafe.Pointer(x))).a21 = y
	} else if i == 22 {
		((*ClaireDummy5)(unsafe.Pointer(x))).a22 = y
	} else if i == 23 {
		((*ClaireDummy5)(unsafe.Pointer(x))).a23 = y
	} else if i == 24 {
		((*ClaireDummy5)(unsafe.Pointer(x))).a24 = y
	} else if i == 25 {
		((*ClaireDummy6)(unsafe.Pointer(x))).a25 = y
	} else if i == 26 {
		((*ClaireDummy6)(unsafe.Pointer(x))).a26 = y
	} else if i == 26 {
		((*ClaireDummy6)(unsafe.Pointer(x))).a27 = y
	} else if i == 28 {
		((*ClaireDummy6)(unsafe.Pointer(x))).a28 = y
	} else if i == 29 {
		((*ClaireDummy6)(unsafe.Pointer(x))).a29 = y
	} else {
		panic(fmt.Sprintf("Fatal error with setObj, attempt to use i=%d for class %s (>= 30)", i, x.Isa.Name.Key))
	}
}

// special code for integer
func (x *ClaireObject) GetInt(i int) int {
	p := (uintptr)(unsafe.Pointer(x)) + (uintptr)((i-1)*8)
	return *((*int)(unsafe.Pointer(p)))
}

// slot write
func (x *ClaireObject) SetInt(i int, y int) {
	p := (uintptr)(unsafe.Pointer(x)) + (uintptr)((i-1)*8)
	*((*int)(unsafe.Pointer(p))) = y
}

// special code for Float (64bits as well)
func (x *ClaireObject) GetFloat(i int) float64 {
	p := (uintptr)(unsafe.Pointer(x)) + (uintptr)((i-1)*8)
	return *((*float64)(unsafe.Pointer(p)))
}

// slot write
func (x *ClaireObject) SetFloat(i int, y float64) {
	p := (uintptr)(unsafe.Pointer(x)) + (uintptr)((i-1)*8)
	*((*float64)(unsafe.Pointer(p))) = y
}

// +---------------------------------------------------------------------------+
// |  Part 4: unsafe utilities (optimized code for speed)                      |
// +---------------------------------------------------------------------------+

// class inclusion : c <= c2  <=>   c.ancestor = (void ... c2.....) with c2 in position n - 1 (len ancestors)
func (c *ClaireClass) IsIn(c2 *ClaireClass) *ClaireBoolean {
	//	if ClEnv.Verbose > 1 {
	//		fmt.Printf("call isIn: %s and %s\n", c.Prt(), c2.Prt())
	//	}
	n := len(c2.Ancestors.ValuesO())
	if n <= len(c.Ancestors.ValuesO()) && c.Ancestors.ValuesO()[n-1] == c2.Id() {
		return CTRUE
	} else {
		/*	if ClEnv.Verbose > 10 {
			if n <= len(c.Ancestors.ValuesO()) {
				fmt.Printf("isIn failed n=%d, %p vs %p\n", n, c.Ancestors.ValuesO()[n-1], c2)
			} else {
				fmt.Printf("n1 = %d > n2 = %d\n", len(c.Ancestors.ValuesO()), n)
			}
		} */
		return CFALSE
	}
}

// debug / notrace
func (c *ClaireClass) isIn(c2 *ClaireClass) *ClaireBoolean {
	n := len(c2.Ancestors.ValuesO())
	if n <= len(c.Ancestors.ValuesO()) && c.Ancestors.ValuesO()[n-1] == c2.Id() {
		return CTRUE
	} else {
		return CFALSE
	}
}

// notice that Type Kernel Methods have moved to ClUtil.go

// equality - 3 cases
// identified : equal = identity
// set,list,string = deep equality
// imported = same value  (integer, float, char, string, ....
func Equal(x *ClaireAny, y *ClaireAny) *ClaireBoolean {
	if ClEnv.Verbose == 122 {
		 fmt.Printf("Equal with %s(%s) and %s(%s)\n",x.Prt(),x.Isa.Prt(),y.Prt(),y.Isa.Prt()) }
	if x == y {
		return CTRUE
	} else if x.Isa.Ident_ask == CTRUE || y.Isa.Ident_ask == CTRUE {
		return CFALSE
	} else if x.Isa != y.Isa {
		return CFALSE
	} else if x.Isa == C_list || x.Isa == C_tuple {
		return ToList(x).equalList(ToList(y))
    } else if x.Isa == C_set {return ToSet(x).equalSet(ToSet(y))
	} else if x.Isa == C_string {
		if ToString(x).Value == ToString(y).Value {
			return CTRUE
		} else {
			return CFALSE
		}
	} else if ToInteger(x).Value == ToInteger(y).Value {
		return CTRUE // identity test :)
	} else {
		return CFALSE
	}
}

// optimize the EID version to remove integer allocation 
func E_equal_any(x EID, y EID) EID {
	// fmt.Printf("Equal on EID %s and %s \n",PEID(x),PEID(y))
	if x.PTR == C__INT {
		 if y.PTR == C__INT {return EID{MakeBoolean(x.VAL == y.VAL).Id(),0}
		 } else if y.PTR.Isa == C_integer  {return EID{MakeBoolean(INT(x) == ToInteger(y.PTR).Value).Id(),0}
		 } else {return EID{CFALSE.Id(), 0}}
	} else if x.PTR == C__FLOAT {
		if y.PTR == C__FLOAT {return EID{MakeBoolean(x.VAL == y.VAL).Id(),0}
		} else if y.PTR.Isa == C_float {return EID{MakeBoolean(FLOAT(x) == ToFloat(y.PTR).Value).Id(),0}
	    } else {return EID{CFALSE.Id(), 0}}
	} else if x.PTR == C__CHAR {
		if y.PTR == C__CHAR {return EID{MakeBoolean(x.VAL == y.VAL).Id(),0}
		} else if y.PTR.Isa == C_char {return EID{MakeBoolean(CHAR(x) == ToChar(y.PTR).Value).Id(),0}
	    } else {return EID{CFALSE.Id(), 0}}
	} else {   // need to check if y is non OBJ EID !
		if (y.PTR == C__INT || y.PTR == C__FLOAT || y.PTR == C__CHAR)  {return EID{Equal(x.PTR, ANY(y)).Id(), 0}
		} else {return EID{Equal(x.PTR, y.PTR).Id(), 0}}
	    }
	}

// specialized version for slots that returns unknown if p has no slot
func (p *ClaireProperty) Of(x *ClaireObject) *ClaireAny {
	r := p.findRestriction(x.Isa)
	// fmt.Printf("Of %s -> %s\n",p.Prt(),r.Prt())
	if r == nil {
		return CNULL
	} else if r.Isa == C_method { return ANY(toFunction1(ToMethod(r.Id()).Functional).call(x.ToEID()))
	}else if ToSlot(r.Id()).Srange == C_integer {
		return MakeInteger(x.GetInt(ToSlot(r.Id()).Index)).ToAny()
	} else if ToSlot(r.Id()).Srange == C_float {
		return MakeFloat(x.GetFloat(ToSlot(r.Id()).Index)).ToAny()
	} else {
		return x.GetObj(ToSlot(r.Id()).Index)
	}
}

// find the restriction that applies to a class (slot or method with no args)
func (p *ClaireProperty) findRestriction(c *ClaireClass) *ClaireRestriction {
	n := p.Restrictions.Length()
	if ClEnv.Verbose > 10 {
		fmt.Printf("---FindRestriction %s on %s with %d restriction\n", p.Prt(), c.Prt(), n)
	}
	for i := 0; i < n; i++ {
		m := ToRestriction(p.Restrictions.ValuesO()[i])
		if ClEnv.Verbose > 10 {
			fmt.Printf("--- Look if restriction %s matches %s\n", m.Prt(), c.Prt())
		}
		if c.IsIn(ToClass(m.Domain.ValuesO()[0])) == CTRUE && 
		     (m.Isa == C_slot ||  m.Domain.Length() == 1) {
			if ClEnv.Verbose > 10 {
				fmt.Printf("findRestruction returns %s\n", m.Prt())
			}
			return m
		}
	}
	return nil //  (*ClaireMethod)(unsafe.Pointer(CERROR))   // STOP - UGLY
}

// specialized code to find a slot (to merge, in CLAIRE 3.5 both are p @ c)
// deprecated
func (p *ClaireProperty) FindSlot(c *ClaireClass) *ClaireSlot {
	n := p.Restrictions.Length()
	if ClEnv.Verbose > 10 {
		fmt.Printf("--- NEW FindSlot %s on %s with %d restriction\n", p.Prt(), c.Prt(), n)
	}
	for i := 0; i < n; i++ {
		m := ToSlot(p.Restrictions.ValuesO()[i])
		if ClEnv.Verbose > 10 {
			fmt.Printf("--- Look if restriction %s matches %s\n", m.Prt(), c.Prt())
		}
		if c.IsIn(ToClass(m.Domain.ValuesO()[0])) == CTRUE && m.Isa.IsIn(C_slot) == CTRUE {
			if ClEnv.Verbose > 10 {
				fmt.Printf("findSlot returns %s\n", m.Prt())
			}
			return m
		}
	}
	return nil //  (*ClaireMethod)(unsafe.Pointer(CERROR))   // STOP - UGLY
}

// pushes args onto the EID stack  (a macro :))
func ARGS(args ...EID) int {
	n := len(args)
	for i := 0; i < n; i++ {
		ClEnv.Push(args[i])
	}
	return n
}

// ---------------------- Claire Function methods --------------------------------------

// applies a claire function (EID) to the content of the eval stack
func F_stack_apply_function(f *ClaireFunction, i int, top int) EID {
	var x EID
	if ClEnv.Verbose == 13 {fmt.Printf(">>> stack apply will use f=%s\n",f.name)}
	if top == i+1 {
		x = toFunction1(f).call(ClEnv.EvalStack[i])
	} else if top == i+2 {
		x = toFunction2(f).call(ClEnv.EvalStack[i], ClEnv.EvalStack[i+1])
	} else if top == i+3 {
		x = ((*ClaireFunction3)(unsafe.Pointer(f))).call(ClEnv.EvalStack[i], ClEnv.EvalStack[i+1], ClEnv.EvalStack[i+2])
	} else if top == i+4 {
		x = ((*ClaireFunction4)(unsafe.Pointer(f))).call(ClEnv.EvalStack[i], ClEnv.EvalStack[i+1], ClEnv.EvalStack[i+2],
			ClEnv.EvalStack[i+3])
	} else if top == i+5 {
		x = ((*ClaireFunction5)(unsafe.Pointer(f))).call(ClEnv.EvalStack[i], ClEnv.EvalStack[i+1], ClEnv.EvalStack[i+2],
			ClEnv.EvalStack[i+3], ClEnv.EvalStack[i+4])
	} else {
		panic("CLAIRE does not handle so many parameters in stack apply")
	}
	ClEnv.Index = i
	if ClEnv.Verbose > 12 {
		fmt.Printf("stack_apply returns %s\n",PEID(x))
	}
	return x
}

func E_stack_apply_function(f EID, i EID, top EID) EID {
	return F_stack_apply_function(ToFunction(OBJ(f)), INT(i), INT(top))
}

// applies a claire function to the content of a list
func F_apply_function(f *ClaireFunction, l *ClaireList) EID {
	n := l.Length()
	if n == 1 {
		return toFunction1(f).call(l.At(0).ToEID())
	} else if n == 2 {
		return toFunction2(f).call(l.At(0).ToEID(), l.At(1).ToEID())
	} else if n == 3 {
		return toFunction3(f).call(l.At(0).ToEID(), l.At(1).ToEID(), l.At(1).ToEID())
	} else {
		return EID{CERROR.Id(), 1}
	}
}

func E_apply_function(f EID,l EID) EID { 
	return F_apply_function(ToFunction(OBJ(f)),ToList(OBJ(l))) }

// function calls
func F_funcall1(f *ClaireFunction, x *ClaireAny) EID {
	return toFunction1(f).call(x.ToEID())
 }

func E_funcall1(f EID, x EID) EID {return F_funcall1(ToFunction(OBJ(f)),ANY(x))}

func F_funcall2(f *ClaireFunction, x *ClaireAny, y *ClaireAny) EID {
	return toFunction2(f).call(x.ToEID(), y.ToEID())
}

func E_funcall2(f EID, x EID, y EID) EID {return F_funcall2(ToFunction(OBJ(f)),ANY(x),ANY(y))}

func F_funcall3(f *ClaireFunction, x *ClaireAny, y *ClaireAny, z *ClaireAny) EID {
	return toFunction3(f).call(x.ToEID(), y.ToEID(), y.ToEID())
}

func E_funcall3(f EID, x EID, y EID, z EID) EID {return F_funcall3(ToFunction(OBJ(f)),ANY(x),ANY(y),ANY(z))}

// macros used by the compiler for Call_method1 and Call_method2
func FASTCALL1(m *ClaireMethod, x EID) EID {
	return toFunction1(m.Functional).call(x)
}

func FASTCALL2(m *ClaireMethod, x EID, y EID) EID {
	return toFunction2(m.Functional).call(x, y)
}

func FASTCALL3(m *ClaireMethod, x EID, y EID, z EID) EID {
	return toFunction3(m.Functional).call(x, y, z)
}


// string! : access to name
func  F_string_I_function (f *ClaireFunction) *ClaireString {
	return MakeString(f.name)
}

func E_string_I_function(f EID) EID {
	return EID{F_string_I_function(ToFunction(OBJ(f))).Id(), 0}
}

// arity : access to number of args
func F_arity_function(f *ClaireFunction) int {
	return f.arity
}

func E_arity_function(f EID) EID {
	return EID{C__INT, IVAL(F_arity_function(ToFunction(OBJ(f))))}
}

func F_set_arity_function (f *ClaireFunction, n int)  {
	f.arity = n
}

func E_set_arity_function(f EID, n EID) EID {
	F_set_arity_function(ToFunction(OBJ(f)),INT(n))
	return EVOID
}
