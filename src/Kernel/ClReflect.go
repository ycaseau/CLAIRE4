// ==================================================================
// microCLAIRE                                              CLAIRE 4
// golang microClaire Kernel - started on June 21st, 2020
//
//  Copyright (C) 2020-2023 Yves Caseau. All Rights Reserved.
//  Redistribution and use in source and binary forms are permitted
//  provided that source distribution retains this entire copyright
//  notice and comments.
//  THIS SOFTWARE IS PROVIDED AS IS AND WITHOUT ANY WARRANTY,
//  INCLUDING, WITHOUT LIMITATION, THE IMPLIED WARRANTIES OF
//  MERCHANTABILTY AND FITNESS FOR A PARTICULAR PURPOSE
//
// clReflect.go
// ==================================================================

/* This is the reflective description of the class kernel
 */

package Kernel

import (
	"fmt"
	"os"
//	"time"                // uncomment if perf measures are used
	"unsafe"
)

// +---------------------------------------------------------------------------+
// |  Table of contents                                                        |
// |  Part 1: reflexive description of Claire Objects & Entities  (bootstrap)  |
// |  Part 2: Property methods                                                 |
// |  Part 3: Class & Types methods                                            |
// |  Part 4: Slots and Methods for classes                                    |
// |  Part 5: Read & Write                                                     |
// +---------------------------------------------------------------------------+

// +---------------------------------------------------------------------------+
// |  Part 1: reflexive description of Claire Objects & Entities  (bootstrap)  |
// +---------------------------------------------------------------------------+

// hardcore : just create the objects (classes)
func BootCore() {
	if ClEnv.Verbose > 10 {
		print("--- Start Bootcore ------------- \n")
	}
	C_class = new(ClaireClass)   // C_class is needed for makeClass1
	C_object = new(ClaireClass)   // C_object is needed for Srange (list or set)
	C_list = new(ClaireClass)   // C_list is needed for MakeClass1
	C_set = new(ClaireClass)   // C_set is needed for MakeClass1
	C_class.Isa = C_class
	// create the boot classes
	C_void = makeClass1(new(ClaireClass))
	C_object = makeClass1(C_object)  // needed for c.Instances
	C_symbol = makeClass1(new(ClaireClass))  // needed for names
	C_claire = makeModule1() // needed to create
	C_slot = makeClass1(new(ClaireClass))    // need for slots list
	C_list = makeClass1(C_list)    // needed to create proper core lists
	C_set = makeClass1(C_set)     // descendant is a set
	// two special values : NIL and EMPTY
	CEMPTY = makeNilSet()
	CNIL = makeBootList()
	// fix boots lists
	CEMPTY.of = ToType(CEMPTY.Id())
	CNIL.of = ToType(CEMPTY.Id())
	C_void.Slots.of = ToType(CEMPTY.Id())
	C_object.Slots.of = ToType(CEMPTY.Id())
	C_symbol.Slots.of = ToType(CEMPTY.Id())
	C_slot.Slots.of = ToType(CEMPTY.Id())
	C_list.Slots.of = ToType(CEMPTY.Id())
	C_set.Slots.of = ToType(CEMPTY.Id())
	// hand fix some critical slots
	C_void.Superclass = C_void
	C_void.Ancestors = coreList(ToType(C_class.Id()), C_void.Id())
	C_void.Name = makeSymbol("void")
	C_class.Subclass = ToType(C_class.Id()).EmptySetObject()
	C_class.Instances = coreList(ToType(C_class.Id()))
	// void, being the root, will never gop through MakeClass2 hence we apply the fixes here
	C_void.Subclass.Isa = C_set
	C_void.Instances.Isa = C_list
	C_void.evaluate = EVAL_object // propagate default through inheritance
	C_class.Instances.AddFast(C_void.Id())
	C_void.Descendants = ToType(C_class.Id()).EmptySetObject()    // empty descendant set
	// create other constant objects
	CNULL = new(ClaireAny).Is(C_void)                        // create the unknown object
	C_void.IfWrite = CNULL
	C_void.Dictionary = ToMapSet(CNULL)
	EVOID = EID{CNULL.Id(), 0}
	CTRUE = new(ClaireBoolean)
	CFALSE = new(ClaireBoolean)
	if ClEnv.Verbose > 10 {
		fmt.Printf("--- End Bootcore ------------- \n")
	}
}

// bootstrap : create all the named object
// this is the heart of reflective CLAIRE
func Bootstrap() {
	// start := time.Now()
	InitClEnv()
	ClEnv.Verbose = 0 // debug: show the makeClass & AddSlot
	// two steps to create Class to break the circularity --------
	BootCore()
	// now we can start building the class hierarchy from the top ------------
	C_any = MakeClass("any", C_void, C_claire)
	C_any.Open = 2                        // by default, in CLAIRE 4, we do not keep extensions
	CNULL.Isa = C_any				      // unknown is allowed as an any value
	makeClass2("object", C_object, C_any, C_claire)
	C_exception = MakeClass("exception", C_object, C_claire)
	C_error = MakeClass("error", C_exception, C_claire)
	C_system_error = MakeClass("system_error", C_error, C_claire)
	C_collection = MakeClass("collection", C_object, C_claire)
	C_type_expression = MakeClass("type_expression", C_collection, C_claire)
	C_type = MakeClass("type", C_type_expression, C_claire)

	// second step
	makeClass2("class", C_class, C_type, C_claire)
	C_thing = MakeClass("thing", C_object, C_claire)
	C_thing.Open = 3                   // we keep the extension for all named classes
	C_primitive = MakeClass("primitive", C_any, C_claire)
	C_primitive.Open = -1              // closed
	C_string = MakeClass("string", C_primitive, C_claire)
	C_integer = MakeClass("integer", C_primitive, C_claire)
	C_integer.evaluate = EVAL_integer
	C_float = MakeClass("float", C_primitive, C_claire)
	C_float.evaluate = EVAL_float
	C_system_object = MakeClass("system_object", C_object, C_claire)
    C_unbound_symbol = MakeClass("unbound_symbol", C_system_object, C_claire)
	
    C_environment = MakeClass("environment", C_system_object, C_claire)
	C_Instruction = MakeClass("Instruction", C_system_object, C_claire)
	C_Variable = MakeClass("Variable", C_Instruction, C_claire)
	
	// ClEnv.Isa = C_environment
	// C_claire.register(makeSymbol("system"), ClEnv.Id()) // link symbol to object
	makeClass2("symbol", C_symbol, C_system_object, C_claire)
	C_char = MakeClass("char", C_primitive, C_claire)
	C_char.evaluate = EVAL_char
	C_boolean = MakeClass("boolean", C_thing, C_claire)
	C_system_thing = MakeClass("system_thing", C_thing, C_claire)
	C_keyword = MakeClass("keyword", C_system_thing, C_claire)
	C_relation = MakeClass("relation", C_system_thing, C_claire)
	C_property = MakeClass("property", C_relation, C_claire)
	C_operation = MakeClass("operation", C_property, C_claire)
	C_table = MakeClass("table", C_relation, C_claire)
	C_restriction = MakeClass("restriction", C_object, C_claire)
	C_restriction.Open = 3           // keep instances
	C_method = MakeClass("method", C_restriction, C_claire)
	makeClass2("slot", C_slot, C_restriction, C_claire)
	C_void.Slots = ToType(C_slot.Id()).EmptyList()
	
	// special instances
	ToBoolean(CTRUE.IsNamed(C_boolean, MakeSymbol("true", C_claire)).Id())
	ToBoolean(CFALSE.IsNamed(C_boolean, MakeSymbol("false", C_claire)).Id())
	CTRUE.Not = CFALSE // hack to make not(b) fast for ClaireBoolean
	CFALSE.Not = CTRUE
	CERROR = ToError(C_error.makeObject().Id())
	CEOF = (rune)(-1) // see if that works
	C__INT = new(ClaireThing).Id()
	C__FLOAT = new(ClaireThing).Id()
	C__CHAR = new(ClaireThing).Id()

	// lists, sets, maps
	C_bag = MakeClass("bag", C_type, C_claire)
	C_bag.Open = -1                      // closed to subclassing or instanciation
	makeClass2("list", C_list, C_bag, C_claire)
	C_listargs = MakeClass("listargs", C_list, C_claire)
	// a tuple is physically a list, but different semantics {example: set!(tuple(X,Y))}
	C_tuple = MakeClass("tuple", C_bag, C_claire)
	makeClass2("set", C_set, C_bag, C_claire)
	C_map_set = MakeClass("map_set", C_collection, C_claire)
	C_module = MakeClass("module", C_system_thing, C_claire)
	C_port = MakeClass("port", C_primitive, C_claire)
	C_function = MakeClass("function", C_primitive, C_claire)
	C_array = MakeClass("array", C_collection, C_claire)
	// ports
	claireStdout = MakeOutPort(os.Stdout)
	claireStdin = MakeInPort(os.Stdin)
	ClEnv.Cout = claireStdout
	ClEnv.Ctrace = claireStdout
	ClEnv.Cin = claireStdin

	// Types
	C_type_operator = MakeClass("type_operator", C_type, C_claire)
	C_Union = MakeClass("Union", C_type_operator, C_claire)
	C_Interval = MakeClass("Interval", C_type_operator, C_claire)
	C_Param = MakeClass("Param", C_type_operator, C_claire)
	C_subtype = MakeClass("subtype", C_type_operator, C_claire)
	C_Reference = MakeClass("Reference", C_type_expression, C_claire)
	// end
	C_claire.Name = makeSymbol("claire")
	makeModule2(C_claire, ToModule(CNULL)) // finishes to create Claire module
	// was sup = C_claire >>>>>
	C_mClaire = MakeModule("mClaire", C_claire)
	C_Kernel = MakeModule("Kernel", C_mClaire) // create this package/module :)
	it = C_Kernel                              // check if useful
	ClEnv.Module_I = C_Kernel                  // this is where we are
	C_lambda = MakeClass("lambda", C_object, C_claire)
	C_pair = MakeClass("pair", C_object, C_claire)


	// fixes a few pieces that are missing
	for _, c := range C_class.Instances.ValuesO() {
		ToClass(c).Comment = MakeString(ToClass(c).Name.key) // because C_string does not exist first
		ToClass(c).Ident_ask = CTRUE
		ToClass(c).Params = CNIL
		ToClass(c).Name.definition = C_Kernel
		C_claire.register(ToClass(c).Name, c)
	}
	for _, m := range C_module.Instances.ValuesO() {
		ToModule(m).Name.definition = C_Kernel
	}
	finishClEnv()

	// these classes require equal
	C_set.Ident_ask = CFALSE
	C_list.Ident_ask = CFALSE
	C_tuple.Ident_ask = CFALSE
	C_string.Ident_ask = CFALSE
	C_float.Ident_ask = CFALSE
	C_integer.Ident_ask = CFALSE
	C_port.Ident_ask = CFALSE
	C_char.Ident_ask = CFALSE
	InitChar() // now that C_char exists :)
	// reflective descriptions of slots and methods
	BootSlot()
	BootMethod()
	unknownName = makeSymbol("unknown")            // we need to mark this symbol for lookup
	C_claire.register(unknownName, CNULL)          // it contains CNULL but the "value is known"
	PRIVATE = makeSymbol("private")                // private is a special symbol ...
	C_claire.register(PRIVATE,PRIVATE.Id())        // self-referenced
	// fmt.Println("=== end of bootstrap ====")
		
}

// +---------------------------------------------------------------------------+
// |  Part 2: Property methods                                                 |
// +---------------------------------------------------------------------------+

// create a property - this code could use a cleanup (1) fewer methods (2) use new(C).Isa(C) pattern
func makeProperty(name string) *ClaireProperty {
	return makeNewProperty(makeSymbol(name))
}

// two short cuts to create Kernel and microClaire (mClaire) properties
func makeKernelProperty(name string) *ClaireProperty {
	return makeNewProperty(MakeSymbol(name,C_Kernel))
}

func makeMicroProperty(name string) *ClaireProperty {
	return makeNewProperty(MakeSymbol(name,C_mClaire))
}

// internal function used by Kernel (without instantiate => sets all default values)	
func makeNewProperty(name *ClaireSymbol) *ClaireProperty {
	var o *ClaireProperty = new(ClaireProperty)
	o.Isa = C_property
	o.Name = name
	name.module_I.register(o.Name, o.Id()) // link symbol to object
	o.Restrictions = MakeList(ToType(C_restriction.Id()))
	o.Definition = MakeList(ToType(C_restriction.Id()))
	//if ClEnv.Verbose > 10 {
	//	fmt.Printf("C_restriction : %s = %x\n", C_restriction.Prt(), C_restriction)
	//	fmt.Printf("MakeProperty -> definition -> of:%x = %s\n", o.Definition.of, o.Definition.Of().Prt())
	// }
	o.Multivalued_ask = CFALSE
	o.Inverse = ToRelation(CNULL)
	o.Domain = ToType(C_any.Id())
	o.Range = ToType(C_any.Id())
	o.Reified = CFALSE
	o.Dictionary = CFALSE
	o.Comment = ToString(CNULL)
	o.IfWrite = CNULL
	o.Store_ask = CFALSE
	o.Open = 1
	C_property.Instances.AddFast(o.ToAny())
	return o
}

// version that is is used by the compiler ... 
// default values are put by instantiate (in IsNamed)
func MakeProperty(name string, op int, m *ClaireModule) *ClaireProperty {
	var o *ClaireProperty = ToProperty(new(ClaireProperty).IsNamed(C_property, MakeSymbol(name, m)))
	o.Open = op
	o.Comment = MakeString(name)
	o.Restrictions = MakeList(ToType(C_restriction.Id()))
	if (o.Definition == nil) {panic("Instantiate failed on " + name)}
	if ClEnv.Verbose > 10 {
		fmt.Printf("MakeProprerty(%s) -> %s - definition.Of:%s\n", name, o.Prt(), o.Definition.Of().Prt())
	}
	return o
}

// creates an operation
// todo : cleanup like MakeProperty
func MakeOperation(name string, op int, m *ClaireModule, prec int) *ClaireOperation {
	var o *ClaireOperation = new(ClaireOperation)
	o.Isa = C_operation
	o.Name = makeSymbol(name)
	C_claire.register(o.Name, o.Id()) // link symbol to object
	o.Open = op
	o.Precedence = prec
	o.Restrictions = MakeList(ToType(C_restriction.Id()))
	o.Definition = MakeList(ToType(C_restriction.Id()))
	o.Multivalued_ask = CFALSE
	o.Inverse = ToRelation(CNULL)
	o.Domain = ToType(C_any.Id())
	o.Range = ToType(C_any.Id())
	o.Reified = CFALSE
	o.Dictionary = CFALSE
	o.Comment = ToString(CNULL)
	o.IfWrite = CNULL
	o.Store_ask = CFALSE
	o.Open = 1
	// fmt.Printf("MakeOperation %s -> restriction:%s\n", name, o.Restrictions.Prt())
	C_operation.Instances.AddFast(o.ToAny())
	return o
}

// AddMethod : this creates the method (private)
func (p *ClaireProperty) makeMethod(ltype []*ClaireAny) *ClaireMethod {
	// work on signature
	n := len(ltype)
	ldom := makeListObject(ToType(C_type_expression.Id()), ltype[0:n-1])
	r := ToType(ltype[n-1])
	// test c.ToType is actually c.ClaireType (already in go !)
	lsort := listClass(ltype) // should be a sort extraction from the types
	if ClEnv.Verbose > 10 {
		fmt.Printf("--- create method %s@%s [%d] \n", p.Name.key, ToClass(lsort.At(0)).Name.key, n)
	}
	// create the method
	var o *ClaireMethod = new(ClaireMethod)
	o.Isa = C_method
	o.Selector = p
	if ClEnv.Verbose > 10 {
		fmt.Printf("add method to restrictions: %s\n", p.Restrictions.Prt())
	}
	p.Restrictions.AddFast(o.Id())
	C_method.Instances.AddFast(o.ToAny())
	o.Domain = ToList(ldom.Id())
	o.Range = r
	o.Formula = ToLambda(CNULL) // unknown
	o.Module_I = ClEnv.Module_I // current module
	o.Srange = ToList(lsort.Id())
	o.Status = 0
	o.Comment = ToString(CNULL)
	o.Inline_ask = CFALSE
	o.Typing = CNULL
	return o
}

// define a method with a lambda
func (p *ClaireProperty) DefMethod(l *ClaireLambda, ld []*ClaireAny) {
	m := p.makeMethod(ld)
	m.Formula = l
}

// add a method with a f(x) golang function - special form for compiler
func (p *ClaireProperty) AddMethod(ld []*ClaireAny, status int, f *ClaireFunction) *ClaireMethod {
	if ClEnv.Verbose == 10 {fmt.Printf("--- start creating a method for %s\n", p.Name.key)}
	m := p.makeMethod(ld)
	m.Functional = f
	m.Status = status
	return m
}

// API function - used by the interpreter
func F_add_method_property(p *ClaireProperty, dom *ClaireList, r *ClaireType, status int, f *ClaireFunction) *ClaireMethod {
	l2 := dom.AddFast(r.Id())
	// fmt.Printf("call AddMethod on %s via API\n", p.Prt())
	return p.AddMethod(l2.ValuesO(), status, f)
}

func E_add_method_property(p EID, dom EID, r EID, status EID, f EID) EID {
	return EID{F_add_method_property(ToProperty(OBJ(p)), ToList(OBJ(dom)), ToType(OBJ(r)), INT(status), ToFunction(OBJ(f))).Id(), 0}
}

// version that adds the evaluator function for self_eval
func (p *ClaireProperty) AddEvalMethod(ld []*ClaireAny, status int, f *ClaireFunction, feval eFunc) *ClaireMethod {
	// fmt.Printf("--- start creating a method for %s\n", p.Name.key)
	m := p.makeMethod(ld)
	m.Functional = f
	m.Status = status
	ToClass(ld[0]).evaluate = feval
	return m
}

// how to create a new slot (restriction of p, with range r) in position ix with default def
// This is MUCH simpler than C++ : we have dropped the multivalued based on list
// it may not return an error (produced by the compiler only when safe)
// the error code should move to the definition (Cerror28 when p multivalued)
// note that the default is now the right value computed in define.cl for special cases (int, float, list, set)
func (c *ClaireClass) AddSlot(p *ClaireProperty, r *ClaireType, def *ClaireAny) *ClaireSlot {
	s := new(ClaireSlot)
	c1 := r.Class_I()
	s.Isa = C_slot
	C_slot.Instances.AddFast(s.ToAny())
	// if ClEnv.Verbose > 10 {
	//	fmt.Printf("--- create slot %s@%s position %d and default %x\n", p.Name.key, c.Name.key, s.Index, def)
	//	fmt.Printf(" -> srange c1 will be %s from range = %s\n", c1.Prt(), r.Prt())
	//  }
	// puts s in the list of slots at the right positions and sets the index
	ls := c.Slots.ValuesO()
	ix := len(ls)
	for ix > 0 && ToSlot(ls[ix - 1]).Selector != p {ix--}     // look for slots from
	if ix > 0  {   // slot co-variant override
		// fmt.Printf("------ slot override for p=%s ---------\n",p.Prt())
		s.Index = ix
		s.Srange = ToSlot(ls[ix - 1]).Srange                  // srange cannot change ! 
		ls[ix - 1] = s.Id()
	} else {
		// new in v4.10: check that the class does not have more than 50 slots
		if len(ls) >= 49 {panic("too many slots (more than 50) in class " + c.Name.key + "(fatal error)")}
		s.Index = len(ls) + 1
		s.Srange = c1
	    // CLAIRE 4 : propagate slots down - used only during the bootstrap
	    s2 := c.Descendants
		for k := 0; k < s2.Count; k++  {
			   ToClass(s2.At(k)).Slots.AddFast(s.ToAny())
			   }
	}
	s.Domain = MakeList(ToType(C_type.Id()), c.ToAny())
	s.Selector = p
	if c1 == C_set && p.Restrictions.Length() == 0 {
		p.Multivalued_ask = CTRUE
	}
	s.Default = def
	p.Restrictions.AddFast(s.ToAny())
	s.Module_I = ClEnv.Module_I
	s.Range = r
	s.Comment = ToString(CNULL) // v3.3.42  - from Sylvain ... may cause inspect crash
	if ClEnv.Verbose > 10 {
		fmt.Printf("--- create slot %s@%s srange is %s\n", p.Name.key, c.Name.key, s.Srange.Prt())
	}
	return s
}

func E_add_slot_class(c EID, p EID, r EID, def EID) EID {
	return EID{ToClass(OBJ(c)).AddSlot(ToProperty(OBJ(p)), ToType(OBJ(r)), OBJ(def)).Id(), 0}
}

// +---------------------------------------------------------------------------+
// |  Part 3: Class & Types methods                                            |
// +---------------------------------------------------------------------------+

// create a module
func makeModule1() *ClaireModule {
	m := new(ClaireModule)
	m.table = make(map[string]*ClaireSymbol)
	return m
}

// second step
func makeModule2(m *ClaireModule, sup *ClaireModule) {
	m.Isa = C_module
	C_module.Instances.AddFast(m.Id())
	//m.register(m.Name, m.Id())
	C_claire.register(m.Name, m.Id())
	m.PartOf = sup
	if m != C_claire {sup.Parts.AddFast(m.Id())}
	m.Parts = ToType(C_module.Id()).EmptyList()
	m.Status = 0
	m.Comment = ToString(CNULL)
	m.MadeOf = ToType(C_string.Id()).EmptyList()
	m.Imports = ToType(C_string.Id()).EmptySet()
	m.Resources = ToType(C_string.Id()).EmptyList()
	m.Uses = ToType(C_module.Id()).EmptyList()
	m.Source = ToString(CNULL)
	m.Evaluate = ToFunction(CNULL)        // unused ? 
	m.External = ToString(CNULL)
}

// constructor: create a module from the name and the father in the hierarchy
// this is low level function : idempotent (module may already exist) but not robust
func MakeModule(s string, sup *ClaireModule) *ClaireModule {
	m2 := C_claire.table[s]
	if m2 != nil {
		if m2.value.Isa == C_module { return ToModule(m2.value)
		} else {panic("unsupported conflict on module names with: " + s)}
	} else {
		m := makeModule1()
		m.Name = C_claire.createSymbol(s)
		makeModule2(m, sup)
		return m
	}
}

// register a symbol into a module + assign the value to the symbol
// unclear that we should do both
func (m *ClaireModule) register(s *ClaireSymbol, o *ClaireAny) {
	m.table[s.key] = s
	s.value = o
	if ClEnv.Verbose > 0 {
		// fmt.Printf("--- register %s in module %s\n", s.key, m.Name.key)
		if s != m.Lookup(s.key) {
			panic("register error")
		}
	}
}

// we need this special function to create class in two steps (break the recursion)
// makeClass1 is incomplete because C_set and C_list do not necessarily exist
func makeClass1(c *ClaireClass) *ClaireClass {
	c.Isa = C_class
	// these are incomplete list creation patterns (to be fixed in step2)
	c.Subclass = ToType(C_class.Id()).EmptySetObject()
	c.Instances = ToType(c.Id()).emptyListObject()
	               // makeNilList()   // ToList(makeListObject(ToType(C_class.Id()), []*ClaireAny{}).Id())
	c.Slots = makeBootList()
	return c
}

// second step - further defines c1 that inherits from c2
func makeClass2(name string, c1 *ClaireClass, c2 *ClaireClass, m *ClaireModule) {
	//fmt.Printf("---- MakeClass2(%s)\n", name)
	c1.Name = MakeSymbol(name, m)
	instantiateClass(name,c1,c2)  /// now we define c1 as a subclass of c2
}

// Debug : print a set of classes (temporary : could be removed)
func SCS(s *ClaireSet) string {
	var res string = "{"
	for k := 0; k < s.Count; k++ {res = res + ToClass(s.At(k)).Name.key + " "}
	return res + "}"
}


// this is the bulk of class instantiation
func instantiateClass(name string, c1 *ClaireClass, c2 *ClaireClass) {
	c1.Comment = MakeString(name)
	c1.Superclass = c2
	// complete the classes
	c2.Subclass.AddFast(c1.ToAny())
	c1.Ancestors = ToList(makeListObject(ToType(C_class.Id()), append(copySlice(c2.Ancestors.ValuesO()),
		c1.ToAny())).Id())
	// c1.descendants = transitive closure of subclass / inverse of Ancestors
	c1.Descendants = ToType(C_class.Id()).EmptySetObject() // emptylist
	// fmt.Printf("makeclass(%s): ancestors: %d\n", name, c1.Ancestors.Length())
	for _, y := range c1.Ancestors.ValuesO() {
		ToClass(y).Descendants.AddFast(c1.Id())
    }
	C_class.Instances.AddFast(c1.ToAny())
	// if (c2->open == ClEnv->ephemeral) c->open = ClEnv->ephemeral;
	if c2.Slots.Length() > 0 {
		c1.Slots = c2.Slots.Copy() // slots inheritance
	} else {
		c1.Slots = ToType(c1.Id()).emptyListObject()
		// c1.Slots = makeNilList() // empty tuple (no _expression / no update)
	} // empty list
	c1.Open = c2.Open
	c1.Params = CNIL
	c1.Ident_ask = CTRUE
	c1.IfWrite = CNULL
	c1.evaluate = EVAL_object
	// fmt.Printf("c1:%s, c2:%s, c2.dic = %x\n", c1.Prt(), c2.Prt(), (uintptr)(unsafe.Pointer(c2.Dictionary)))
	if c2.Dictionary.Id() == CNULL {
		c1.Dictionary = ToMapSet(CNULL)
	} else {
		// fmt.Printf("copy dictionary from %s\n", c2.Name.key)
		c1.Dictionary = c2.Dictionary.Copy()
	}

}

// this is the regular function to create a new class
// this is used both in Kernel and with compiled code when we known that no error will occur
func MakeClass(name string, c *ClaireClass, m *ClaireModule) *ClaireClass {
	if ClEnv.Verbose > 1 {
		fmt.Printf("--- make compiled class %s\n", name)
		// fmt.Printf("--- super = %x\n", c)
	}
	o := makeClass1(new(ClaireClass))
	makeClass2(name, o, c, m)
	m.register(o.Name, o.Id())
	return o
}

// this is safe (compiler) version to create a class unless it already exists (then reuse)
// this is only used by the compiler (assumes no mistakes) ... to support forward declaration
func NewClass(name string, c *ClaireClass, m *ClaireModule) *ClaireClass {
	s := m.createSymbol(name)
	if s.value != nil && s.value.Isa == C_class {
		 return ToClass(s.value)
	} else { o := makeClass1(new(ClaireClass))
		     o.Name = s
			 instantiateClass(name,o,c)
			 m.register(s, o.Id())
			 return o }
}

// this is how it is called in define.cl - c is the super class
// this call may return an error if symbol is already used
func (s *ClaireSymbol) Class_I(c *ClaireClass) EID {
	if ClEnv.Verbose > 10 {
		fmt.Printf("--- make interpreted class %s\n", s.key)
	}
	// x := F_new_thing_class(C_class,s)  will not work because of the extra slot (evaluate)
	var o *ClaireClass
	if s.value != nil {
		if s.value.Isa != C_class {return Cerror(18,s.Id(),s.value.Isa.Id())
		} else {o = ToClass(s.value)}
	} else {
		o = new(ClaireClass)
		o.Isa = C_class
		o.Name = s
		s.module_I.register(s, o.Id())
		C_class.Instances.AddFast(o.Id())
	}
	o.Subclass = ToType(C_class.Id()).EmptySet()
	o.Instances = ToType(o.Id()).EmptyList() //  makeNilList()
	instantiateClass(s.key, o, c)
	return EID{o.Id(),0}
}

func E_class_I_symbol(s EID, c EID) EID {
	return ToSymbol(OBJ(s)).Class_I(ToClass(OBJ(s))) 
}


// make an object - used by compiler => range is *ClaireAny
func (c *ClaireClass) New() *ClaireAny {
	o := c.makeObject()
	o.Isa = c
	if c.Open == 3  {
		c.Instances.AddFast(&o.ClaireAny)
	}
	return o.ToAny()
}

// compiler pattern : new(C).Is(c)
func (o *ClaireAny) Is(c *ClaireClass) *ClaireAny {
	o.Isa = c
	c.instantiate(ToObject(o))
	if c.Open == 3  {
		c.Instances.AddFast(o)
	}
	return o
}

// compiler pattern : ToC(new(C).IsNamed(c,s))
// notice that we cannot use the .Is :)
// for the time being, assumes no name conflict ....
func (o *ClaireThing) IsNamed(c *ClaireClass, s *ClaireSymbol) *ClaireAny {
	o.Isa = c
	c.instantiate(ToObject(o.Id()))
	c.Instances.AddFast(o.Id())
	o.Name = s
	s.module_I.register(s, o.Id())
	if ClEnv.Verbose > 10 {
		fmt.Printf("instantiate %s with _expression %s\n", o.Name.key, o.Isa.Name.key)
		fmt.Printf("Prt gives %s\n", o.Prt())
	}
	return o.Id()
}

// CLAIRE4 pattern : places all default values
// we do not keep the prototype slot any more
// we need to distinguish easy default (no inverse, no if_write) that are dealt with here .. and complex ones
// that are managed in define.cl (instantiation) or odefine.cl(compilation)
func (c *ClaireClass) instantiate(o *ClaireObject) {
	if ClEnv.Verbose == -1 {
		fmt.Printf("instantiate %s with %d slots\n", c.Name.key, c.Slots.Length())
		if c.Slots.Length() > 0 {fmt.Printf("slots = %s\n",c.Slots.Prt())}
	}
	for i, s := range c.Slots.ValuesO() {
		if i > 0 {
			r := ToSlot(s).Srange
			if ClEnv.Verbose == -1 {
				fmt.Printf("look at slot %s:%d with default %s\n", ToSlot(s).Selector.Name.key, ToSlot(s).Index, ToSlot(s).Default.Prt())
				fmt.Printf("Srange is %s, owner(def) is %s\n", r.Prt(), ToSlot(s).Default.Isa.Prt())
			}
			if ToSlot(s).Default == CNULL {
				o.SetObj(i+1, CNULL)
			} else if r == C_integer {
				o.SetInt(i+1, ToInteger(ToSlot(s).Default).Value)
			} else if r == C_float {
				o.SetFloat(i+1, ToFloat(ToSlot(s).Default).Value)
			} else if ToSlot(s).Default.Isa == C_list { // notice that the copy is based on the value
				o.SetObj(i+1, ToList(ToSlot(s).Default).Copy().Id())
			} else if ToSlot(s).Default.Isa == C_set { // same for sets
				o.SetObj(i+1, ToSet(ToSlot(s).Default).Copy().Id())
				if ClEnv.Verbose == -1 {
					fmt.Printf(">> we have a set at position %d: %s\n",i+1,o.GetObj(i+1).Prt())}
			} else {
				o.SetObj(i+1, ToSlot(s).Default)
			}
		}
	}
}

// instanciation of a class - compiler pattern for anyObject! - exact number of args + right types
func (c *ClaireClass) Make(args ...*ClaireAny) *ClaireAny {
	o := c.New()
	n := len(args)
	for i := 0; i < n; i++ { // n args : n+1 slot (Isa not included)
		r := c.Slots.ValuesO()[i+1]
		if ToSlot(r).Srange.Id() == C_integer.Id() {
			// fmt.Printf("Make : use arg1 = %d\n", ToInteger(args[i]).Value)
			ToObject(o).SetInt(i+2, ToInteger(args[i]).Value)
		} else if ToSlot(r).Srange.Id() == C_float.Id() {
			ToObject(o).SetFloat(i+2, ToFloat(args[i]).Value)
		} else {
			ToObject(o).SetObj(i+2, args[i])
		}
	}
	return o
}

// special form for integer args (cool for intervals or any object with integer slots)
func (c *ClaireClass) MakeInts(args ...int) *ClaireAny {
	o := c.New()
	n := len(args)
	for i := 0; i < n; i++ { // n args : n+1 slot (Isa not included)
		ToObject(o).SetInt(i+2, args[i])
		} 
	return o
}


// the functions that are imported in Core: (1) mClaire/new!(self:class)
func F_new_object_class(c *ClaireClass) *ClaireObject {
	o := ToObject(c.New())
	c.instantiate(o)
	return o
}

func E_new_object_class(e EID) EID {
	return EID{F_new_object_class(ToClass(OBJ(e))).Id(), 0}
}

// same for (2)  mClaire/new!(self:class,%nom:symbol)
// this one is different since it can produce an error
func F_new_thing_class(c *ClaireClass, s *ClaireSymbol) EID {
	if ClEnv.Verbose > 10 {
		fmt.Printf("create new_thing[%s] %s\n", c.Prt(), s.key)
	}
	if s.value != nil && s.value != CNULL {
		// fmt.Printf("thing %s already exists\n",s.key)
		// fmt.Printf("value is  %s already exists\n",s.value.Prt())
		if s.value.Isa != c {return Cerror(18,s.Id(),s.value.Isa.Id())
		} else {return EID{s.value,0}}
	} else {
		o := ToThing(c.New())
		c.instantiate(ToObject(o.Id()))
		o.Name = s
		s.module_I.register(s, o.Id())
		return EID{o.Id(),0}}
}

func E_new_thing_class(e EID, s EID) EID {
	return F_new_thing_class(ToClass(OBJ(e)), ToSymbol(OBJ(s)))
}

// returns the sort associated to class : any (represents EID), object, char, integer or float
// for the time being, we allow string & functions as sorts (they are not objects) 
func (c *ClaireClass) Sort_I() *ClaireClass {
	if c == C_integer || c == C_float || c == C_char || c == C_string || c == C_function { return c
    } else if c == C_any || c == C_primitive || c == C_void  { return C_any
	} else {return C_object}
	}

func E_sort_I_class (c EID) EID {return EID{ToClass(OBJ(c)).Sort_I().Id(),0}}

// create a dumb function
func F_make_function_string(name *ClaireString) *ClaireFunction {
	o := new(ClaireFunction1)
	o.Isa = C_function
	o.call = ESELF
	o.name = name.Value
	return (*ClaireFunction)(unsafe.Pointer(o))
}

func ESELF(x EID) EID { panic("don't do this !") }

func E_make_function_string(s EID) EID {
	return EID{F_make_function_string(ToString(OBJ(s))).Id(), 0}
}

// shallow copy of object
func (o *ClaireObject) Copy() *ClaireObject {
	x := ToObject(o.Isa.New())
	for i, s := range o.Isa.Slots.ValuesO() {
		if i > 0 {
			r := ToSlot(s).Srange
			if r == C_integer {
				x.SetInt(i+1, o.GetInt(i+1))
			} else if r == C_float {
				x.SetFloat(i+1, o.GetFloat(i+1))
			} else {
				x.SetObj(i+1, o.GetObj(i+1))
			}
		}
	}
	return x
}

func E_copy_object(o EID) EID { return EID{ToObject(OBJ(o)).Copy().Id(), 0} }

// +---------------------------------------------------------------------------+
// |  Part 4: Properties & Methods reflection                                  |
// +---------------------------------------------------------------------------+

// create all the slots
func BootSlot() {
	// create properties
	C_isa = makeProperty("isa")
	C_mClaire_index = makeMicroProperty("index")
	C_value = makeProperty("value")
	C_arg = makeProperty("arg")
	C_name = makeProperty("name")
	C_mClaire_pname = makeMicroProperty("pname")
	C_comment = makeProperty("comment")
	C_slots = makeProperty("slots")
	C_superclass = makeProperty("superclass")
	C_ancestors = makeProperty("ancestors")
	C_subclass = makeProperty("subclass")
	C_descendants = makeProperty("descendants")
	C_open = makeProperty("open")
	C_instances = makeProperty("instances")
	C_params = makeProperty("params")	
	C_mClaire_graph = makeMicroProperty("graph")
	C_if_write = makeProperty("if_write")
	C_ident_ask = makeProperty("ident?")
	C_domain = makeProperty("domain")
	C_range = makeProperty("range")
	C_store_ask = makeProperty("store?")
	C_inverse = makeProperty("inverse")
	C_multivalued_ask = makeProperty("multivalued?")
	C_trace_I = makeProperty("trace!")
	C_restrictions = makeProperty("restrictions")
	C_mClaire_definition = makeMicroProperty("definition")
	C_reified = makeProperty("reified")
	C_module_I = makeProperty("module!")
	C_selector = makeProperty("selector")
	C_mClaire_srange = makeMicroProperty("srange")
	C_default = makeProperty("default")
	C_formula = makeProperty("formula")
	C_vars = makeProperty("vars")
	C_body = makeProperty("body")
	C_dimension = makeProperty("dimension")
	C_parts = makeProperty("parts")
	C_part_of = makeProperty("part_of")
	C_uses = makeProperty("uses")
	C_source = makeProperty("source")
	C_made_of = makeProperty("made_of")
	C_resources = makeProperty("resources")
	C_mClaire_status = makeMicroProperty("status")
	C_external = makeProperty("external")
	C_put = makeProperty("put")
	C_get = makeProperty("get")
	C_precedence = makeProperty("precedence")
	C_dictionary = makeProperty("dictionary")
	C_mClaire_t1 = makeMicroProperty("t1")
	C_mClaire_t2 = makeMicroProperty("t2")
	C_arg1 = makeProperty("arg1")
	C_arg2 = makeProperty("arg2")
	C_args = makeProperty("args")
	C_inline_ask = makeProperty("inline?")
	C_Kernel_typing = makeKernelProperty("typing")
	C_functional = makeProperty("functional")
	C_verbose = makeProperty("verbose")
	C_exception_I = makeProperty("exception!")
	C_version = makeProperty("version")
	C_ctrace = makeProperty("ctrace")
	C_cout = makeProperty("cout")
	C_cin = makeProperty("cin")
	C_base = makeProperty("base")
	C_debug_I = makeProperty("debug!")
//	C_step_I = makeProperty("step!")
	C_last_debug = makeProperty("last_debug")
	C_last_index = makeProperty("last_index")
	C_spy_I = makeProperty("spy!")
	C_count_call = makeMicroProperty("count_call")
	C_count_level = makeMicroProperty("count_level")
	C_count_trigger = makeMicroProperty("count_trigger")
	C_close = makeProperty("close")
	C_final = makeProperty("final")
	C_abstract = makeProperty("abstract")
	C_ephemeral = makeProperty("ephemeral")
	C_jito_ask = makeProperty("jito?")	
	C_n_line = makeProperty("n_line")
	C_imports = makeProperty("imports")
	C_first = makeProperty("first")
	C_second = makeProperty("second")
	C_osname = makeProperty("osname")

	// slots (same order as Kernel)
	C_any.AddSlot(C_isa, ToType(C_class.Id()), CNULL)
	C_system_error.AddSlot(C_mClaire_index, ToType(C_integer.Id()), AnyInteger(0))
	C_system_error.AddSlot(C_value, ToType(C_any.Id()), CNULL)
	C_system_error.AddSlot(C_arg, ToType(C_any.Id()), CNULL)
	C_thing.AddSlot(C_name, ToType(C_symbol.Id()), MakeSymbol("unamed", C_claire).Id())
	// a class
	C_class.AddSlot(C_name, ToType(C_symbol.Id()), CNULL)
	C_class.AddSlot(C_comment, ToType(C_string.Id()), CNULL)
	C_class.AddSlot(C_slots, ToType(C_list.Id()), CNIL.Id())
	C_class.AddSlot(C_superclass, ToType(C_class.Id()), CNULL)
	C_class.AddSlot(C_subclass, ToType(C_list.Id()), CNIL.Id())
	C_class.AddSlot(C_ancestors, ToType(C_list.Id()), CNIL.Id())
	C_class.AddSlot(C_descendants, ToType(C_set.Id()), CNIL.Id())
	C_class.AddSlot(C_open, ToType(C_integer.Id()), AnyInteger(1))
	C_class.AddSlot(C_instances, ToType(C_list.Id()), CNULL)
	C_class.AddSlot(C_params, ToType(C_list.Id()), CNULL)
	C_class.AddSlot(C_dictionary, ToType(C_map_set.Id()), CNULL)
	C_class.AddSlot(C_ident_ask, ToType(C_boolean.Id()), CTRUE.ToAny())
	C_class.AddSlot(C_if_write, ToType(C_any.Id()), CNULL)
	// properties, restrictions, methods, slots
	C_relation.AddSlot(C_comment, ToType(C_string.Id()), CNULL)
	C_relation.AddSlot(C_domain, ToType(C_type.Id()), C_any.Id())
	C_relation.AddSlot(C_range, ToType(C_type.Id()), C_any.Id())
	C_relation.AddSlot(C_if_write, ToType(C_any.Id()), CNULL)
	C_relation.AddSlot(C_store_ask, ToType(C_boolean.Id()), CFALSE.ToAny())
	C_relation.AddSlot(C_inverse, ToType(C_relation.Id()), CNULL)
	C_relation.AddSlot(C_open, ToType(C_integer.Id()), AnyInteger(2))
	C_relation.AddSlot(C_multivalued_ask, ToType(C_boolean.Id()), CFALSE.ToAny())
	C_table.AddSlot(C_mClaire_graph, ToType(C_any.Id()), CNULL)
	C_table.AddSlot(C_params, ToType(C_any.Id()), CNULL)
	C_table.AddSlot(C_default, ToType(C_any.Id()), CNULL)
	C_property.AddSlot(C_trace_I, ToType(C_integer.Id()), AnyInteger(0))
	C_property.AddSlot(C_restrictions, ToType(C_list.Id()), MakeList(ToType(C_restriction.Id())).Id()) // TODO : create types directly ??
	C_property.AddSlot(C_mClaire_definition, ToType(C_list.Id()), MakeList(ToType(C_restriction.Id())).Id())   // same : try to create a _expression
	C_property.AddSlot(C_dictionary, ToType(C_boolean.Id()), CFALSE.ToAny())
	C_property.AddSlot(C_reified, ToType(C_boolean.Id()), CFALSE.ToAny())
	C_operation.AddSlot(C_precedence, ToType(C_integer.Id()), AnyInteger(0))
	C_restriction.AddSlot(C_module_I, ToType(C_module.Id()), CNULL)
	C_restriction.AddSlot(C_comment, ToType(C_string.Id()), CNULL)
	C_restriction.AddSlot(C_domain, ToType(C_list.Id()), CNULL) // TODO : same
	C_restriction.AddSlot(C_range, ToType(C_type.Id()), C_any.Id())
	C_restriction.AddSlot(C_selector, ToType(C_property.Id()), CNULL)
	C_slot.AddSlot(C_mClaire_srange, ToType(C_class.Id()), CNULL)
	C_slot.AddSlot(C_default, ToType(C_any.Id()), CNULL)
	C_slot.AddSlot(C_mClaire_index, ToType(C_integer.Id()), AnyInteger(0))
	C_method.AddSlot(C_mClaire_srange, ToType(C_list.Id()), CNULL)
	C_method.AddSlot(C_formula, ToType(C_lambda.Id()), CNULL)
	C_method.AddSlot(C_functional, ToType(C_function.Id()), CNULL)
	C_method.AddSlot(C_Kernel_typing, ToType(C_any.Id()), CNULL)
	C_method.AddSlot(C_mClaire_status, ToType(C_integer.Id()), AnyInteger(0))
	C_method.AddSlot(C_inline_ask, ToType(C_boolean.Id()), CFALSE.Id())
	C_lambda.AddSlot(C_vars, ToType(C_list.Id()), CNIL.Id())
	C_lambda.AddSlot(C_body, ToType(C_any.Id()), CNULL)
	C_lambda.AddSlot(C_dimension, ToType(C_integer.Id()), AnyInteger(0))
	C_unbound_symbol.AddSlot(C_name, ToType(C_symbol.Id()), CNULL)
	C_pair.AddSlot(C_first, ToType(C_any.Id()), CNULL)
	C_pair.AddSlot(C_second, ToType(C_any.Id()), CNULL)
	
	// modules
	C_module.AddSlot(C_comment, ToType(C_string.Id()), CNULL)
	C_module.AddSlot(C_parts, ToType(C_list.Id()), ToType(C_module.Id()).EmptyList().Id())
	C_module.AddSlot(C_part_of, ToType(C_module.Id()), C_claire.Id())
	C_module.AddSlot(C_uses, ToType(C_list.Id()), ToType(C_module.Id()).EmptyList().Id())
	C_module.AddSlot(C_source, ToType(C_string.Id()), CNULL)
	C_module.AddSlot(C_made_of, ToType(C_list.Id()), ToType(C_string.Id()).EmptyList().Id()) // CNIL.Id())
	C_module.AddSlot(C_mClaire_status, ToType(C_integer.Id()), AnyInteger(0))
	C_module.AddSlot(C_external, ToType(C_string.Id()), CNULL)
	C_module.AddSlot(C_imports, ToType(C_map_set.Id()), CNULL)
	// this does not work ...
	C_module.AddSlot(C_resources, ToType(C_list.Id()), ToType(C_string.Id()).EmptyList().Id()) 
	// Types
	C_Union.AddSlot(C_mClaire_t1, ToType(C_type.Id()), CNULL)
	C_Union.AddSlot(C_mClaire_t2, ToType(C_type.Id()), CNULL)
	C_Interval.AddSlot(C_arg1, ToType(C_integer.Id()), CNULL)
	C_Interval.AddSlot(C_arg2, ToType(C_integer.Id()), CNULL)
	C_Param.AddSlot(C_arg, ToType(C_class.Id()), CNULL)
	C_Param.AddSlot(C_params, ToType(C_list.Id()), CNULL)
	C_Param.AddSlot(C_args, ToType(C_list.Id()), CNULL)
	C_subtype.AddSlot(C_arg, ToType(C_class.Id()), CNULL)
	C_subtype.AddSlot(C_mClaire_t1, ToType(C_type.Id()), CNULL)
	C_Reference.AddSlot(C_args, ToType(C_list.Id()), CNULL)
	C_Reference.AddSlot(C_mClaire_index, ToType(C_integer.Id()), AnyInteger(0))
	C_Reference.AddSlot(C_arg, ToType(C_boolean.Id()), CFALSE.Id())
	// meta description of the system object
	C_environment.AddSlot(C_verbose, ToType(C_integer.Id()), AnyInteger(0))
	C_environment.AddSlot(C_exception_I, ToType(C_exception.Id()), CNULL)
	C_environment.AddSlot(C_module_I, ToType(C_module.Id()), CNULL)
	C_environment.AddSlot(C_name, ToType(C_string.Id()), CNULL)
	C_environment.AddSlot(C_osname, ToType(C_string.Id()), CNULL)
	C_environment.AddSlot(C_version, ToType(C_float.Id()), AnyFloat(0.0))
	C_environment.AddSlot(C_ctrace, ToType(C_port.Id()), CNULL)
	C_environment.AddSlot(C_cout, ToType(C_port.Id()), CNULL)
	C_environment.AddSlot(C_cin, ToType(C_port.Id()), CNULL)
	C_environment.AddSlot(C_mClaire_index, ToType(C_integer.Id()), AnyInteger(0))
	C_environment.AddSlot(C_base, ToType(C_integer.Id()), AnyInteger(0))
	C_environment.AddSlot(C_debug_I, ToType(C_integer.Id()), AnyInteger(0))
	C_environment.AddSlot(C_trace_I, ToType(C_integer.Id()), AnyInteger(0))
	C_environment.AddSlot(C_last_debug, ToType(C_integer.Id()), AnyInteger(0))
	C_environment.AddSlot(C_last_index, ToType(C_integer.Id()), AnyInteger(0))
	C_environment.AddSlot(C_spy_I, ToType(C_object.Id()), CNULL)
	C_environment.AddSlot(C_count_call, ToType(C_integer.Id()), AnyInteger(0))
	C_environment.AddSlot(C_count_level, ToType(C_integer.Id()), AnyInteger(0))
	C_environment.AddSlot(C_count_trigger, ToType(C_any.Id()), CNULL)
	C_environment.AddSlot(C_params, ToType(C_list.Id()), CNULL)
	C_environment.AddSlot(C_close, ToType(C_integer.Id()), AnyInteger(-1))
	C_environment.AddSlot(C_abstract, ToType(C_integer.Id()), AnyInteger(0))
	C_environment.AddSlot(C_final, ToType(C_integer.Id()), AnyInteger(1))
	C_environment.AddSlot(C_default, ToType(C_integer.Id()), AnyInteger(2))
	C_environment.AddSlot(C_open, ToType(C_integer.Id()), AnyInteger(3))
	C_environment.AddSlot(C_ephemeral, ToType(C_integer.Id()), AnyInteger(4))
	C_environment.AddSlot(C_jito_ask, ToType(C_boolean.Id()), CTRUE.Id())
	C_environment.AddSlot(C_n_line, ToType(C_integer.Id()), AnyInteger(0))
	// variables
	C_Variable.AddSlot(C_mClaire_pname, ToType(C_symbol.Id()), CNULL)
	C_Variable.AddSlot(C_range, ToType(C_type.Id()), CNULL)
	C_Variable.AddSlot(C_mClaire_index, ToType(C_integer.Id()), AnyInteger(-1))       //  -1 is the new default = no Index
	C_Variable.Params = MakeList(ToType(C_any.Id()), C_mClaire_pname.Id(), C_range.Id())
}

// add the bootstrap part for lists, sets, tuples and array
func BootMethod() {
	// fmt.Println("-------------- start Boot Method ------------------------")
	C_copy = makeProperty("copy")
	C_empty = makeProperty("empty")
	C_length = makeProperty("length")
	C_contain_ask = makeProperty("contain?")
	C_integer_I = makeProperty("integer!")
	C_float_I = makeProperty("float!")
	C_make_list = makeProperty("make_list")
	C_array_I = makeProperty("array!")
	C_list_I = makeProperty("list!")
	C_funcall = makeProperty("funcall")
	C_fastcall = makeProperty("fastcall")
	C_nth = makeProperty("nth")
	C_nth_equal = makeProperty("nth=")
	C_nth_plus = makeProperty("nth+")
	C_nth_dash = makeProperty("nth-")
	C_nth_put = makeProperty("nth_put")
	C_nth_get = makeProperty("nth_get")
	C_self_print = makeProperty("self_print")
	C_self_print.Open = 3
	C_princ = makeProperty("princ")
	C_gensym = makeProperty("gensym")
	C_store = makeProperty("store")
	C_commit = makeProperty("commit")
	C_backtrack = makeProperty("backtrack")
	C_choice = makeProperty("choice")
	C_symbol_I = makeProperty("symbol!")
	C_make_string = makeProperty("make_string")
	C_make_array = makeProperty("make_array")
	C_random = makeProperty("random")
	C_random_I = makeProperty("random!")
	C_string_I = makeProperty("string!")
	C_class_I = makeProperty("class!")
	C_new = makeProperty("new")
	C_mClaire_new_I = makeMicroProperty("new!")
	C_make_function = makeProperty("make_function")
	C_cdr = makeProperty("cdr")
	C_skip = makeProperty("skip")
	C_shrink = makeProperty("shrink")
	C_size = makeProperty("size")
	C_cast_I = makeProperty("cast!")
	C_tuple_I = makeProperty("tuple!")
	C_c_princ = makeProperty("c_princ")
	C_substring = makeProperty("substring")
	C_included = makeProperty("included")
	C_begin = makeProperty("begin")
	C_end = makeProperty("end")
	C_port_I = makeProperty("port!")
	C_use_as_output = makeProperty("use_as_output")
	//	C_time_set = makeProperty("time_set")
	//	C_time_get = makeProperty("time_get")
	//	C_time_read = makeProperty("time_read")
	//	C_time_show = makeProperty("time_show")
	C_shell = makeProperty("shell")
	C_fclose = makeProperty("fclose")
	// C_getenv = makeProperty("getenv")
	C_world_ask = makeProperty("world?")
	C_world_id = makeProperty("world_id")
	C_set_length = makeProperty("set_length")
	C_exit = makeProperty("exit")
	C_graph_get = makeProperty("graph_get")
	C_graph_put = makeProperty("graph_put")
	C_graph_init = makeProperty("graph_init")
	C_map_I = makeProperty("map!")
	C_boolean_I = makeProperty("boolean!")
	C_dict_get = makeProperty("dict_get")
	C_dict_put = makeProperty("dict_put")
	C_read = makeProperty("read")
	C_read_ident = makeProperty("read_ident")
	C_read_number = makeProperty("read_number")
	C_read_thing = makeProperty("read_thing")
	C_read_string = makeProperty("read_string")
	C_print = makeProperty("print")
	C_log = makeProperty("log")
	C_cos = makeProperty("cos")
	C_sin = makeProperty("sin")
	C_atan = makeProperty("atan")
	C_sqrt = makeProperty("sqrt")
	C_sqrt = makeProperty("sqrt")
	C__exp2 = makeProperty("^2")
	C_set_I = makeProperty("set!")
	C_char_I = makeProperty("char!")
	C_apply = makeProperty("apply")
	C_stack_apply = makeProperty("stack_apply")
	C_of = makeProperty("of")
	C_mClaire_restore_state = makeMicroProperty("restore_state")
	C_abort = makeProperty("abort")
	C_reboot = makeProperty("reboot")
	C_defined = makeProperty("defined")
	C_fopen = makeProperty("fopen")
	C_empty_list = makeProperty("empty_list")
	C_empty_set = makeProperty("empty_set")
	C_slot_get = makeProperty("slot_get")
	C_write_fast = makeProperty("write_fast")
	C_getc = makeProperty("getc")
	C_putc = makeProperty("putc")
	C_namespace = makeProperty("namespace")
	C_date_I = makeProperty("date!")
	C_sort_I = makeProperty("sort!")
	C_add_slot = makeProperty("add_slot")
	C_flush = makeProperty("flush")
	C_add_method = makeProperty("add_method")
	C_arity = makeKernelProperty("arity")
	C_set_arity = makeKernelProperty("set_arity")
	C_slice = makeProperty("slice")
	C_statistics = makeProperty("statistics")	
	C_hash = makeProperty("hash")
	C_file_separator = makeProperty("file_separator")
	// C_getenv = makeProperty("getenv")

	// operation
	C_add = MakeOperation("add",0, C_claire, 10)
	C_add_I = MakeOperation("add!",0, C_claire, 10)
	C_add_star = MakeOperation("add*",0, C_claire, 10)
	C_delete = MakeOperation("delete",0, C_claire, 10)
	C__equal = MakeOperation("=", 0, C_claire, 60)
	C__inf = MakeOperation("<", 0, C_claire, 60)
	C__inf_equal = MakeOperation("<=", 0, C_claire, 60)
	C__sup = MakeOperation(">", 0, C_claire, 60)
	C__sup_equal = MakeOperation(">=", 0, C_claire, 60)
	C__dot_dot = MakeOperation("..", 0, C_claire, 30)
	C__7 = MakeOperation("/", 0, C_claire, 10)
	C__7_plus = MakeOperation("/+", 0, C_claire, 10)
	C__dash = MakeOperation("-", 0, C_claire, 20)
	C__star = MakeOperation("*", 0, C_claire, 10)
	C__Z = MakeOperation("€", 0, C_claire, 50)
	// v4.12 : create an alias
	C_claire.register(makeSymbol("%"), C__Z.Id())        // link symbol to object
	C__exp = MakeOperation("^", 0, C_claire, 5)
	C_min = MakeOperation("min", 0, C_claire, 20)
	C_max = MakeOperation("max", 0, C_claire, 20)
	C_mod = MakeOperation("mod", 0, C_claire, 10)
	C_cons = MakeOperation("cons", 0, C_claire, 10)
	C_max.Open = 2
	C_min.Open = 2

	// fmt.Println("---------------------- start defining methods -----------------------------")
	// methods that are defined in Kernel
	C_apply.AddMethod(Signature(C_function.Id(), C_list.Id(), C_any.Id()), 1, MakeFunction2(E_apply_function, "apply_function"))
	C_stack_apply.AddMethod(Signature(C_integer.Id(), C_void.Id()), 0, MakeFunction1(E_stack_add, "stack_add"))
	C_stack_apply.AddMethod(Signature(C_function.Id(), C_integer.Id(), C_integer.Id(), C_any.Id()), 1, MakeFunction3(E_stack_apply_function, "stack_apply_function"))
	C__equal.AddMethod(Signature(C_any.Id(), C_any.Id(), C_boolean.Id()), 0, MakeFunction2(E_equal_any, "equal_any"))
    C_string_I.AddMethod(Signature(C_function.Id(), C_string.Id()), 0, MakeFunction1(E_string_I_function, "string_I_function"))
    C_arity.AddMethod(Signature(C_function.Id(), C_integer.Id()), 0, MakeFunction1(E_arity_function, "arity_function"))
	C_set_arity.AddMethod(Signature(C_function.Id(), C_integer.Id(),C_void.Id()), 0, MakeFunction2(E_set_arity_function, "set_arity_function"))
	C_funcall.AddMethod(Signature(C_function.Id(), C_any.Id(),C_any.Id()), 1, MakeFunction2(E_funcall1, "funcall1"))
	C_funcall.AddMethod(Signature(C_function.Id(), C_any.Id(),C_any.Id(),C_any.Id()), 2, MakeFunction3(E_funcall2, "funcall2"))
	C_funcall.AddMethod(Signature(C_function.Id(), C_any.Id(),C_any.Id(),C_any.Id(),C_any.Id()), 3, MakeFunction4(E_funcall3, "funcall3"))

	// ClReflect
	C_class_I.AddMethod(Signature(C_symbol.Id(), C_class.Id(), C_class.Id()), 1, MakeFunction2(E_class_I_symbol, "class_I_symbol"))
	C_mClaire_new_I.AddMethod(Signature(C_class.Id(), C_object.Id()), 0, MakeFunction1(E_new_object_class, "new_object_class"))
	C_mClaire_new_I.AddMethod(Signature(C_class.Id(), C_symbol.Id(), C_thing.Id()), 1, MakeFunction2(E_new_thing_class, "new_thing_class"))
	C_sort_I.AddMethod(Signature(C_class.Id(), C_class.Id()), 0, MakeFunction1(E_sort_I_class, "sort_I_class"))
	C_make_function.AddMethod(Signature(C_string.Id(), C_function.Id()), 0, MakeFunction1(E_make_function_string, "make_function_string"))
	C_add_slot.AddMethod(Signature(C_class.Id(), C_property.Id(),C_type.Id(),C_any.Id(),C_slot.Id()), 0, MakeFunction4(E_add_slot_class, "add_slot_class"))
	C_add_method.AddMethod(Signature(C_property.Id(),C_list.Id(),C_type.Id(),C_integer.Id(),C_any.Id(),C_method.Id()), 0, MakeFunction5(E_add_method_property, "add_method_property"))
	C_copy.AddMethod(Signature(C_object.Id(), C_object.Id()), 0, MakeFunction1(E_copy_object, "copy_object"))
    C_write_fast.AddMethod(Signature(C_property.Id(), C_object.Id(), C_any.Id(), C_any.Id()), 1, MakeFunction3(E_write_fast_property, "write_fast_property"))
	C_read.AddMethod(Signature(C_property.Id(), C_object.Id(), C_any.Id()), 1, MakeFunction2(E_read_property, "read_property"))
	C_slot_get.AddMethod(Signature(C_object.Id(), C_integer.Id(), C_class.Id(), C_any.Id()), 0, MakeFunction3(E_slot_get_object, "slot_get_object"))

	// ClBag
	C_of.AddMethod(Signature(C_bag.Id(), C_type.Id()), 0, MakeFunction1(E_of_bag, "of_bag"))
	C_of.AddMethod(Signature(C_array.Id(), C_type.Id()), 0, MakeFunction1(E_of_bag, "of_bag"))
	C_cast_I.AddMethod(Signature(C_bag.Id(), C_type.Id(), C_type.Id()), 0, MakeFunction2(E_cast_I_bag, "cast_I_bag"))
	C_copy.AddMethod(Signature(C_list.Id(), C_list.Id()), 0, MakeFunction1(E_copy_list, "copy_list"))
	C_nth_get.AddMethod(Signature(C_list.Id(), C_integer.Id(), C_any.Id()), 0, MakeFunction2(E_nth_get_list, "nth_get_list"))
	// CLAIRE4 : duplication since tuple methods are the same as lists, but cannot be inherited from bags
	C_nth_get.AddMethod(Signature(C_tuple.Id(), C_integer.Id(), C_any.Id()), 0, MakeFunction2(E_nth_get_list, "nth_get_list"))
	C_nth_put.AddMethod(Signature(C_list.Id(), C_integer.Id(), C_any.Id(), C_any.Id()), 0, MakeFunction3(E_nth_put_list, "nth_put_list"))
	C_nth_put.AddMethod(Signature(C_array.Id(), C_integer.Id(),C_any.Id(), C_any.Id()), 0, MakeFunction3(E_nth_put_list, "nth_put_list"))
	C_empty.AddMethod(Signature(C_list.Id(), C_list.Id()), 0, MakeFunction1(E_empty_list, "empty_list"))
	C_delete.AddMethod(Signature(C_list.Id(), C_any.Id(), C_list.Id()), 0, MakeFunction2(E_delete_list, "delete_list"))
	C_length.AddMethod(Signature(C_list.Id(), C_integer.Id()), 0, MakeFunction1(E_length_list, "length_list"))
	C_length.AddMethod(Signature(C_tuple.Id(), C_integer.Id()), 0, MakeFunction1(E_length_list, "length_list"))
	C_length.AddMethod(Signature(C_array.Id(), C_integer.Id()), 0, MakeFunction1(E_length_list, "length_list"))
	C_contain_ask.AddMethod(Signature(C_list.Id(), C_any.Id(), C_list.Id()), 0, MakeFunction2(E_contain_ask_list, "contain_ask_list"))
	C_add.AddMethod(Signature(C_list.Id(), C_any.Id(), C_list.Id()), 1, MakeFunction2(E_add_list, "add_list"))
	C_contain_ask.AddMethod(Signature(C_set.Id(), C_any.Id(), C_set.Id()), 0, MakeFunction2(E_contain_ask_set, "contain_ask_set"))
	C_add.AddMethod(Signature(C_set.Id(), C_any.Id(), C_set.Id()), 0, MakeFunction2(E_add_set, "add_set"))
	C_cons.AddMethod(Signature(C_any.Id(), C_list.Id(), C_list.Id()), 0, MakeFunction2(E_cons_any, "cons_any"))
	C_cdr.AddMethod(Signature(C_list.Id(), C_list.Id()), 1, MakeFunction1(E_cdr_list, "cdr_list"))
	C_make_list.AddMethod(Signature(C_integer.Id(), C_any.Id(),C_list.Id()), 0, MakeFunction2(E_make_list_integer, "make_list_integer"))
	C_get.AddMethod(Signature(C_list.Id(), C_any.Id(), C_integer.Id()), 0, MakeFunction2(E_index_list, "#index_list"))
	C_add_star.AddMethod(Signature(C_list.Id(), C_list.Id(), C_list.Id()), 1, MakeFunction2(E_add_star_list, "add_star_list"))
	C_add_I.AddMethod(Signature(C_list.Id(), C_any.Id(), C_list.Id()), 0, MakeFunction2(E_add_I_list, "add_I_list"))
	C__7_plus.AddMethod(Signature(C_list.Id(), C_list.Id(), C_list.Id()), 0, MakeFunction2(E_append_list, "append_list"))
	C_nth.AddMethod(Signature(C_list.Id(), C_integer.Id(), C_any.Id()), 1, MakeFunction2(E_nth_list, "nth_list"))
	C_nth.AddMethod(Signature(C_tuple.Id(), C_integer.Id(), C_any.Id()), 1, MakeFunction2(E_nth_list, "nth_list"))
	C_nth.AddMethod(Signature(C_array.Id(), C_integer.Id(), C_any.Id()), 1, MakeFunction2(E_nth_list, "nth_list"))   // v4.0.5
	C_nth_plus.AddMethod(Signature(C_list.Id(), C_integer.Id(), C_any.Id(), C_list.Id()), 1, MakeFunction3(E_nth_plus_list, "nth_plus_list"))
	C_nth_dash.AddMethod(Signature(C_list.Id(), C_integer.Id(), C_list.Id()), 1, MakeFunction2(E_nth_dash_list, "nth_dash_list"))
	C_nth_equal.AddMethod(Signature(C_list.Id(), C_integer.Id(), C_any.Id(),C_any.Id()), 1, MakeFunction3(E_nth_equal_list, "nth_equal_list"))
	C_nth_equal.AddMethod(Signature(C_array.Id(), C_integer.Id(), C_any.Id(),C_any.Id()), 1, MakeFunction3(E_nth_equal_list, "nth_equal_list"))  // v4.0.5
	C_skip.AddMethod(Signature(C_list.Id(), C_integer.Id(), C_list.Id()), 0, MakeFunction2(E_skip_list, "skip_list"))
	C_shrink.AddMethod(Signature(C_list.Id(), C_integer.Id(), C_list.Id()), 0, MakeFunction2(E_shrink_list, "shrink_list"))
	C_slice.AddMethod(Signature(C_list.Id(), C_integer.Id(), C_integer.Id(), C_list.Id()), 0, MakeFunction3(E_slice_list, "slice_list"))
	C_size.AddMethod(Signature(C_set.Id(), C_integer.Id()), 0, MakeFunction1(E_size_set, "size_set"))
	C_contain_ask.AddMethod(Signature(C_set.Id(), C_any.Id(), C_boolean.Id()), 0, MakeFunction2(E_contain_ask_set, "contain_ask_set"))
	C_add_I.AddMethod(Signature(C_set.Id(), C_any.Id(), C_set.Id()), 0, MakeFunction2(E_add_I_set, "add_I_set"))
	C_copy.AddMethod(Signature(C_set.Id(), C_set.Id()), 0, MakeFunction1(E_copy_set, "copy_set"))
	C_empty.AddMethod(Signature(C_set.Id(), C_set.Id()), 0, MakeFunction1(E_empty_set, "empty_set"))
	C_delete.AddMethod(Signature(C_set.Id(), C_any.Id(), C_set.Id()), 0, MakeFunction2(E_delete_set, "delete_set"))
	C__exp.AddMethod(Signature(C_set.Id(), C_set.Id(), C_set.Id()), 0, MakeFunction2(E__exp_set, "_exp_set"))
	C__7_plus.AddMethod(Signature(C_set.Id(), C_set.Id(), C_set.Id()), 0, MakeFunction2(E_append_set, "append_set"))
	// C_set_I.AddMethod(Signature(C_set.Id(), C_list.Id(), C_set.Id()), 0, MakeFunction1(E_set_I_list, "set_I_list"))
	// C_list_I.AddMethod(Signature(C_set.Id(), C_list.Id()), 0, MakeFunction1(E_list_I_set, "list_I_set"))
	C__dot_dot.AddMethod(Signature(C_integer.Id(), C_integer.Id(), C_set.Id()), 0, MakeFunction2(E_sequence_integer, "sequence_integer"))
	C_copy.AddMethod(Signature(C_tuple.Id(), C_tuple.Id()), 0, MakeFunction1(E_copy_tuple, "copy_tuple"))
	C_tuple_I.AddMethod(Signature(C_list.Id(), C_tuple.Id()), 0, MakeFunction1(E_tuple_I_list, "tuple_I_list"))
	C_list_I.AddMethod(Signature(C_tuple.Id(), C_list.Id()), 0, MakeFunction1(E_list_I_tuple, "list_I_tuple"))
	// C_addFast.AddMethod(Signature(func E_addFast_tuple(l EID, x EID) EID {
	// C_make_array.AddMethod(...) => in types.cl, to get a 2nd order _expression
	C_list_I.AddMethod(Signature(C_array.Id(), C_list.Id()), 0, MakeFunction1(E_list_I_array, "list_I_array"))
	// C_array_I.AddMethod(Signature(C_list.Id(), C_array.Id()), 0, MakeFunction1(E_array_I_list, "array_I_list"))
	C_empty_list.AddMethod(Signature(C_type.Id(), C_list.Id()), 0, MakeFunction1(E_empty_list_type, "empty_list_type"))
	C_empty_set.AddMethod(Signature(C_type.Id(), C_set.Id()), 0, MakeFunction1(E_empty_set_type, "empty_set_type"))

	// ClString
	C_char_I.AddMethod(Signature(C_integer.Id(), C_char.Id()), 0, MakeFunction1(E_char_I_integer, "char_I_integer"))
	C_integer_I.AddMethod(Signature(C_char.Id(), C_integer.Id()), 0, MakeFunction1(E_integer_I_char, "integer_I_char"))
	C_princ.AddMethod(Signature(C_char.Id(), C_void.Id()), 0, MakeFunction1(E_princ_char, "princ_char"))
	C_c_princ.AddMethod(Signature(C_char.Id(), C_void.Id()), 0, MakeFunction1(E_c_princ_char, "c_princ_char"))
	C_length.AddMethod(Signature(C_string.Id(), C_integer.Id()), 0, MakeFunction1(E_length_string, "length_string"))
	C_copy.AddMethod(Signature(C_string.Id(), C_string.Id()), 0, MakeFunction1(E_copy_string, "copy_string"))
	C_princ.AddMethod(Signature(C_string.Id(), C_void.Id()), 0, MakeFunction1(E_princ_string, "princ_string"))
	C_self_print.AddMethod(Signature(C_string.Id(), C_void.Id()), 0, MakeFunction1(E_self_print_string, "self_print_string"))
	C__7_plus.AddMethod(Signature(C_string.Id(), C_string.Id(), C_string.Id()), 0, MakeFunction2(E_append_string, "append_string"))
	C_integer_I.AddMethod(Signature(C_string.Id(), C_integer.Id()), 0, MakeFunction1(E_integer_I_string, "integer_I_string"))
	C_slice.AddMethod(Signature(C_string.Id(), C_integer.Id(), C_integer.Id(), C_string.Id()), 0, MakeFunction3(E_slice_string, "substring_string"))
	C_substring.AddMethod(Signature(C_string.Id(), C_integer.Id(), C_integer.Id(), C_string.Id()), 0, MakeFunction3(E_substring_string, "substring_string"))
	C_substring.AddMethod(Signature(C_string.Id(), C_string.Id(), C_boolean.Id(), C_integer.Id()), 0, MakeFunction3(E_included_string, "included_string"))
	C_get.AddMethod(Signature(C_string.Id(), C_char.Id(), C_integer.Id()), 0, MakeFunction2(E_get_string, "get_string"))
	C__inf_equal.AddMethod(Signature(C_string.Id(), C_string.Id(), C_boolean.Id()), 0, MakeFunction2(E_less_string, "_less_string"))
	C_included.AddMethod(Signature(C_string.Id(), C_string.Id(), C_boolean.Id(), C_integer.Id()), 0, MakeFunction3(E_included_string, "included_string"))
	C_nth.AddMethod(Signature(C_string.Id(), C_integer.Id(), C_char.Id()), 1, MakeFunction2(E_nth_string, "nth_string"))
	C_nth_equal.AddMethod(Signature(C_string.Id(), C_integer.Id(), C_char.Id(), C_void.Id()), 0, MakeFunction3(E_nth_set_string, "nth_set_string"))
	C_string_I.AddMethod(Signature(C_integer.Id(), C_string.Id()), 0, MakeFunction1(E_string_I_integer, "string_I_integer"))
	C_make_string.AddMethod(Signature(C_integer.Id(), C_char.Id(), C_string.Id()), 0, MakeFunction2(E_make_string_integer, "make_string_integer"))
	C_make_string.AddMethod(Signature(C_list.Id(), C_string.Id()), 1, MakeFunction1(E_make_string_list, "make_string_list"))
	C_list_I.AddMethod(Signature(C_string.Id(), C_list.Id()), 0, MakeFunction1(E_list_I_string, "list_I_string"))
	C_c_princ.AddMethod(Signature(C_string.Id(), C_void.Id()), 0, MakeFunction1(E_c_princ_string, "c_princ_string"))
	C_symbol_I.AddMethod(Signature(C_string.Id(), C_module.Id(), C_symbol.Id()), 0, MakeFunction2(E_symbol_I_string, "symbol_I_string"))
	C_put.AddMethod(Signature(C_symbol.Id(), C_any.Id(), C_any.Id()), 0, MakeFunction2(E_put_symbol, "put_symbol"))
	//C_get.AddMethod(Signature(C_symbol.Id(), C_any.Id()), 0, MakeFunction1(E_get_symbol, "get_symbol"))
	C__7_plus.AddMethod(Signature(C_symbol.Id(), C_any.Id(), C_symbol.Id()), 0, MakeFunction2(E_append_symbol, "append_symbol"))
	C_princ.AddMethod(Signature(C_symbol.Id(), C_void.Id()), 0, MakeFunction1(E_princ_symbol, "princ_symbol"))
	// in CLAIRE4 : these 4 methods replace 4 slots access (from CLAIRE3.5) 
	C_string_I.AddMethod(Signature(C_symbol.Id(), C_string.Id()), 0, MakeFunction1(E_string_I_symbol, "string_I_symbol"))
	C_module_I.AddMethod(Signature(C_symbol.Id(), C_module.Id()), 0, MakeFunction1(E_module_I_symbol, "module_I_symbol"))
	C_defined.AddMethod(Signature(C_symbol.Id(), C_module.Id()), 0, MakeFunction1(E_defined_symbol, "defined_symbol"))
	C_value.AddMethod(Signature(C_symbol.Id(), C_any.Id()), 0, MakeFunction1(E_value_symbol, "value_symbol"))
    //
	C_gensym.AddMethod(Signature(C_string.Id(), C_symbol.Id()), 0, MakeFunction1(E_gensym_string, "gensym_string"))
	C_c_princ.AddMethod(Signature(C_symbol.Id(), C_void.Id()), 0, MakeFunction1(E_c_princ_symbol, "c_princ_symbol"))
	C_begin.AddMethod(Signature(C_module.Id(), C_void.Id()), 0, MakeFunction1(E_begin_module, "begin_module"))
	C_end.AddMethod(Signature(C_module.Id(), C_void.Id()), 0, MakeFunction1(E_end_module, "end_module"))
	C_namespace.AddMethod(Signature(C_module.Id(), C_void.Id()), 0, MakeFunction1(E_namespace_module, "namespace"))
	// <deprecated> C_value.AddMethod(Signature(C_string.Id(), C_any.Id()), 0, MakeFunction1(E_value_string, "value_string"))
	// C_value.AddMethod(Signature(C_module.Id(), C_string.Id(), C_any.Id()), 0, MakeFunction2(E_value_module, "value_module"))
	C_get.AddMethod(Signature(C_module.Id(), C_string.Id(), C_any.Id()), 0, MakeFunction2(E_get_symbol_module, "get_symbol_module"))
	// C_getenv.AddMethod(Signature(C_string.Id(), C_string.Id()), 0, MakeFunction1(E_getenv_string, "getenv_string"))
	C_port_I.AddMethod(Signature(C_string.Id(), C_port.Id()), 0, MakeFunction1(E_port_I_string, "port_I_string"))
	C_port_I.AddMethod(Signature(C_void.Id(), C_port.Id()), 0, MakeFunction1(E_port_I_void, "port_I_void"))
	C_use_as_output.AddMethod(Signature(C_port.Id(), C_port.Id()), 0, MakeFunction1(E_use_as_output, "use_as_output"))
	C_fclose.AddMethod(Signature(C_port.Id(), C_void.Id()), 0, MakeFunction1(E_fclose_port, "fclose_port"))
	C_string_I.AddMethod(Signature(C_port.Id(), C_string.Id()), 0, MakeFunction1(E_string_I_port, "string_I_port"))
	C_length.AddMethod(Signature(C_port.Id(), C_integer.Id()), 0, MakeFunction1(E_length_port, "length_port"))
	C_set_length.AddMethod(Signature(C_port.Id(), C_integer.Id(), C_void.Id()), 0, MakeFunction2(E_set_length_port, "set_length_port"))
	C_getc.AddMethod(Signature(C_port.Id(), C_char.Id()), 0, MakeFunction1(E_getc_port, "Getc"))
	C_putc.AddMethod(Signature(C_char.Id(), C_port.Id(),C_void.Id()), 0, MakeFunction2(E_putc_char, "Putc"))
	C_fopen.AddMethod(Signature(C_string.Id(), C_string.Id(), C_port.Id()), 1, MakeFunction2(E_fopen_string, "fopen_string"))
	C_flush.AddMethod(Signature(C_port.Id(), C_void.Id()), 0, MakeFunction1(E_flush_port, "flush_port"))
	C_flush.AddMethod(Signature(C_port.Id(), C_integer.Id(),C_void.Id()), 0, MakeFunction2(E_pushback_port, "pushback_port"))


	// ClEnv
	//C_time_set.AddMethod(Signature(C_void.Id(), C_void.Id()), 0, MakeFunction1(E_time_set_void, "time_set_void"))
	//C_time_get.AddMethod(Signature(C_void.Id(), C_void.Id()), 0, MakeFunction1(E_time_get_void, "time_get_void"))
	//C_time_read.AddMethod(Signature(C_void.Id(), C_void.Id()), 0, MakeFunction1(E_time_read_void, "time_read_void"))
	//C_time_show.AddMethod(Signature(C_void.Id(), C_void.Id()), 0, MakeFunction1(E_time_show_void, "time_show_void"))
	C_close.AddMethod(Signature(C_exception.Id(), CEMPTY.Id()), 1, MakeFunction1(E_close_exception, "close_exception"))
	C_shell.AddMethod(Signature(C_string.Id(), C_void.Id()), 0, MakeFunction1(E_claire_shell, "claire_shell"))
	C_exit.AddMethod(Signature(C_integer.Id(), C_void.Id()), 0, MakeFunction1(E_CL_exit, "CL_exit"))
	C_reboot.AddMethod(Signature(C_void.Id(), C_void.Id()), 0, MakeFunction1(E_claire_stat, "claire_stat"))
	C_abort.AddMethod(Signature(C_environment.Id(), C_void.Id()), 0, MakeFunction1(E_abort_system, "abort_system"))
	C_file_separator.AddMethod(Signature(C_environment.Id(), C_string.Id()), 0, MakeFunction1(E_file_separator, "file_separator"))
	C_statistics.AddMethod(Signature(C_void.Id(), C_void.Id()), 0, MakeFunction1(E_claire_stat, "claire_stat"))
	C_hash.AddMethod(Signature(C_list.Id(), C_any.Id(), C_integer.Id()), 0, MakeFunction2(E_hash_list, "hash_list"))
	C_mClaire_restore_state.AddMethod(Signature(C_void.Id(), C_void.Id()), 0, MakeFunction1(E_restore_state_void, "restore_state_void"))
	C_store.AddMethod(Signature(C_list.Id(), C_integer.Id(), C_any.Id(), C_boolean.Id(), C_any.Id()), 0, MakeFunction4(E_store_list, "store_list"))
	C_store.AddMethod(Signature(C_array.Id(), C_integer.Id(), C_any.Id(), C_boolean.Id(), C_any.Id()), 0, MakeFunction4(E_store_list, "store_list"))
	C_store.AddMethod(Signature(C_object.Id(), C_integer.Id(), C_class.Id(), C_any.Id(), C_boolean.Id(), C_any.Id()), 0, MakeFunction5(E_store_object, "store_object"))
	C_choice.AddMethod(Signature(C_void.Id(), C_void.Id()), 0, MakeFunction1(E_world_push, "world_push"))
	C_backtrack.AddMethod(Signature(C_void.Id(), C_void.Id()), 0, MakeFunction1(E_world_pop, "world_pop"))
	C_commit.AddMethod(Signature(C_void.Id(), C_void.Id()), 0, MakeFunction1(E_world_remove, "world_remove"))
	C_world_ask.AddMethod(Signature(C_void.Id(), C_integer.Id()), 0, MakeFunction1(E_world_number, "world_number"))
	C_world_id.AddMethod(Signature(C_void.Id(), C_integer.Id()), 0, MakeFunction1(E_world_get_id, "world_get_id"))
	C_map_I.AddMethod(Signature(C_type.Id(), C_type.Id(), C_map_set.Id()), 0, MakeFunction2(E_map_I_type, "map_I_type"))
	C_get.AddMethod(Signature(C_map_set.Id(), C_any.Id(), C_any.Id()), 0, MakeFunction2(E_get_map, "get_map"))
	C_put.AddMethod(Signature(C_map_set.Id(), C_any.Id(), C_any.Id(), C_any.Id()), 1, MakeFunction3(E_put_map, "put_map"))
	C_dict_get.AddMethod(Signature(C_any.Id(), C_any.Id(), C_any.Id()), 0, MakeFunction2(E_dict_get_any, "dict_get_any"))
	C_dict_put.AddMethod(Signature(C_any.Id(), C_any.Id(), C_any.Id(), C_void.Id()), 0, MakeFunction3(E_dict_put_any, "dict_put_any"))
	C_graph_get.AddMethod(Signature(C_table.Id(), C_any.Id(), C_any.Id()), 0, MakeFunction2(E_graph_get_table, "graph_get_table"))
	C_graph_put.AddMethod(Signature(C_table.Id(), C_any.Id(), C_any.Id(), C_any.Id()), 0, MakeFunction3(E_graph_put_table, "graph_put_table"))
	C_graph_init.AddMethod(Signature(C_table.Id(), C_void.Id()), 0, MakeFunction1(E_graph_init_table, "graph_init_table"))
	C_read_string.AddMethod(Signature(C_port.Id(), C_string.Id()), 0, MakeFunction1(E_read_string_port, "read_string_port"))
	C_read_ident.AddMethod(Signature(C_port.Id(), C_any.Id()), 1, MakeFunction1(E_read_ident_port, "read_ident_port"))
	C_read_number.AddMethod(Signature(C_port.Id(), C_any.Id()), 0, MakeFunction1(E_read_number_port, "read_number_port"))
	C_read_thing.AddMethod(Signature(C_port.Id(), C_module.Id(), C_char.Id(), C_module.Id(), C_any.Id()), 1, MakeFunction4(E_read_thing_port, "read_thing_port"))
    C_set_I.AddMethod(Signature(C_map_set.Id(), C_set.Id()), 0, MakeFunction1(E_set_I_map_set, "set_I_map_set"))
	C_domain.AddMethod(Signature(C_map_set.Id(), C_type.Id()), 0, MakeFunction1(E_domain_map_set, "domain_map_set"))
	C_range.AddMethod(Signature(C_map_set.Id(), C_type.Id()), 0, MakeFunction1(E_range_map_set, "range_map_set"))
	
	// ClUtil
	C_class_I.AddMethod(Signature(C_type_expression.Id(), C_class.Id()), 0, MakeFunction1(E_class_I_type, "class_I_type"))
	C_boolean_I.AddMethod(Signature(C_any.Id(), C_boolean.Id()), 0, MakeFunction1(E_boolean_I_any, "boolean_I_any"))
	C_princ.AddMethod(Signature(C_integer.Id(), C_void.Id()), 0, MakeFunction1(E_princ_integer, "princ_integer"))
	C_random.AddMethod(Signature(C_integer.Id(), C_integer.Id()), 0, MakeFunction1(E_random_integer, "random_integer"))
	C_random_I.AddMethod(Signature(C_integer.Id(), C_void.Id()), 0, MakeFunction1(E_random_I_integer, "random_I_integer"))
	// C_min.AddMethod(Signature(C_integer.Id(), C_integer.Id(), C_integer.Id()), 0, MakeFunction2(E_min_integer, "min_integer"))
	// C_max.AddMethod(Signature(C_integer.Id(), C_integer.Id(), C_integer.Id()), 0, MakeFunction2(E_max_integer, "max_integer"))
	C__dash.AddMethod(Signature(C_integer.Id(), C_integer.Id()), 0, MakeFunction1(E_ch_sign, "ch_sign"))
	C_date_I.AddMethod(Signature(C_integer.Id(), C_string.Id()), 0, MakeFunction1(E_date_I_integer, "date_I_integer"))
	C__7.AddMethod(Signature(C_integer.Id(), C_integer.Id(), C_integer.Id()), 1, MakeFunction2(E__7_integer, "_7_integer"))
	C_mod.AddMethod(Signature(C_integer.Id(), C_integer.Id(), C_integer.Id()), 1, MakeFunction2(E_mod_integer, "mod_integer"))
	C__exp.AddMethod(Signature(C_integer.Id(), C_integer.Id(), C_integer.Id()), 1, MakeFunction2(E__exp_integer, "_exp_integer"))
	C__exp2.AddMethod(Signature(C_integer.Id(), C_integer.Id()), 1, MakeFunction1(E__exp2_integer, "_exp2_integer"))
	// C__inf.AddMethod(Signature(C_integer.Id(), C_integer.Id(), C_boolean.Id()), 0, MakeFunction2(E_inf_integer, "_inf_integer"))
	// C__inf_equal.AddMethod(Signature(C_integer.Id(), C_integer.Id(), C_boolean.Id()), 0, MakeFunction2(E_inf_equal_integer, "_inf_equal_integer"))
	// C__sup.AddMethod(Signature(C_integer.Id(), C_integer.Id(), C_boolean.Id()), 0, MakeFunction2(E_sup_integer, "_sup_integer"))
	C__sup_equal.AddMethod(Signature(C_integer.Id(), C_integer.Id(), C_boolean.Id()), 0, MakeFunction2(E__sup_equal_integer, "_sup_equal_integer"))
	C__star.AddMethod(Signature(C_integer.Id(), C_integer.Id(), C_integer.Id()), 1, MakeFunction2(E_times_integer, "times_integer"))
	C_princ.AddMethod(Signature(C_float.Id(), C_void.Id()), 0, MakeFunction1(E_princ_float, "princ_float"))
	C_print.AddMethod(Signature(C_float.Id(), C_void.Id()), 0, MakeFunction1(E_print_float, "print_float"))
	C_print.AddMethod(Signature(C_float.Id(), C_integer.Id(), C_void.Id()), 0, MakeFunction2(E_print_format_float, "print_format_float"))
	C__exp.AddMethod(Signature(C_float.Id(), C_float.Id(), C_float.Id()), 0, MakeFunction2(E__exp_float, "_exp_float"))
	C_sqrt.AddMethod(Signature(C_float.Id(), C_float.Id()), 0, MakeFunction1(E_sqrt_float, "sqrt_float"))
	C_log.AddMethod(Signature(C_float.Id(), C_float.Id()), 0, MakeFunction1(E_log_float, "log_float"))
	C_sin.AddMethod(Signature(C_float.Id(), C_float.Id()), 0, MakeFunction1(E_sin_float, "sin_float"))
	C_cos.AddMethod(Signature(C_float.Id(), C_float.Id()), 0, MakeFunction1(E_cos_float, "cos_float"))
	C_atan.AddMethod(Signature(C_float.Id(), C_float.Id()), 0, MakeFunction1(E_atan_float, "atan_float"))
	C_float_I.AddMethod(Signature(C_integer.Id(), C_float.Id()), 0, MakeFunction1(E_to_float, "to_float"))
	C_integer_I.AddMethod(Signature(C_float.Id(), C_integer.Id()), 1, MakeFunction1(E_integer_I_float, "integer_I_float"))
	C__inf.AddMethod(Signature(C_float.Id(), C_float.Id(), C_boolean.Id()), 0, MakeFunction2(E__inf_float, "_inf_float"))
	C__inf_equal.AddMethod(Signature(C_float.Id(), C_float.Id(), C_boolean.Id()), 0, MakeFunction2(E__inf_equal_float, "_inf_equal_float"))
	C__sup.AddMethod(Signature(C_float.Id(), C_float.Id(), C_boolean.Id()), 0, MakeFunction2(E__sup_float, "_sup_float"))
	C__sup_equal.AddMethod(Signature(C_float.Id(), C_float.Id(), C_boolean.Id()), 0, MakeFunction2(E__sup_equal_float, "_sup_equal_float"))
	// fmt.Println("============= End Boot Method ===========================")

}

// +---------------------------------------------------------------------------+
// |  Part 5: Read & Write                                                     |
// +---------------------------------------------------------------------------+

// reflective slot access (use n: index and s : srange)
func (p *ClaireObject) SlotGet(n int, s *ClaireClass) *ClaireAny {
	if s == C_integer {
		return MakeInteger(p.GetInt(n)).Id()
	} else if s == C_float {
		return MakeFloat(p.GetFloat(n)).Id()
	} else {
		return p.GetObj(n)
	}
}

// this method is accessible from CLAIRE
func E_slot_get_object(x EID, n EID, s EID) EID {
	return ToObject(OBJ(x)).SlotGet(INT(n),ToClass(OBJ(s))).ToEID()
}

// --------------- read: property method --------------------------------------------------------
// reflective slot access 
// we use to the EID version since this is not really used much
// func (p *ClaireProperty) Read(x *ClaireObject) *ClaireAny {
//	return ANY(p.ReadEID(x.ToEID()))
func (p *ClaireProperty) Read(x *ClaireObject) EID {    // must return EID since errors may occur
	return p.ReadEID(x.ToEID())
}

// optimized version that writes and returns an EID (no allocation)
func (p *ClaireProperty) ReadEID(x EID) EID {
	// fmt.Printf("read property p : %s\n",p.Prt)
	r := p.findRestriction(OWNER(x))
	if r == nil || r.Isa != C_slot { return Cerror(6, ANY(x), p.Id())
    } else { 
		s := ToSlot(r.Id()).Srange
		i := ToSlot(r.Id()).Index
		z := ToObject(OBJ(x))
		if s == C_integer { return EID{C__INT,IVAL(z.GetInt(i))}
		} else if s == C_float { return EID{C__FLOAT,FVAL(z.GetFloat(i))}
		} else { 
			    y := z.GetObj(i) 
				if y == CNULL && ToSlot(r.Id()).Range.Contains(y) == CFALSE {
					return Cerror(1,p.Id(),ANY(x))
				} else {return y.ToEID()}
				}}
}

// this method is accessible from CLAIRE
func E_read_property(p EID,y EID) EID {
	return ToProperty(OBJ(p)).ReadEID(y)
}


// --------------- write : property method -------------------------------------------------------

// reflective slot access 
// we use to the EID version since this is not really used much
func (p *ClaireProperty) WriteFast(x *ClaireObject, y *ClaireAny) *ClaireAny {
	return ANY(p.WriteEID(x,y.ToEID()))
}

// optimized version that writes and returns an EID (no allocation)
func (p *ClaireProperty) WriteEID(x *ClaireObject, y EID) EID {
	r := p.findRestriction(x.Isa)
	if r == nil || r.Isa != C_slot { return Cerror(6, x.Id(), p.Id())
    } else { 
		s := ToSlot(r.Id()).Srange
		i := ToSlot(r.Id()).Index
		if ToSlot(r.Id()).Range.CONTAINS(y) == CTRUE {
			if s == C_integer { x.SetInt(i,INT(y))
			} else if s == C_float { x.SetFloat(i,FLOAT(y))
			} else { x.SetObj(i,ANY(y)) }                       // convert to ANY (can be an int)
			return y
        } else {fmt.Printf("WriteEID fails because %s does not belong to %s\n",ToSlot(r.Id()).Range.Prt(),PEID(y))
			    return Cerror(17,ANY(y),ToSlot(r.Id()).Range.Id())} }
}

// this method is accessible from CLAIRE interpreter
func E_write_fast_property(p EID, x EID, y EID) EID {
	return ToProperty(OBJ(p)).WriteEID(ToObject(OBJ(x)), y)
}