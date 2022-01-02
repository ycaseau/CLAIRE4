// ==================================================================
// golang experiments Phase 2 - Copyright (C) Yves Caseau
// started on June 21st, 2020-2021
// clBag.go
// ==================================================================

/* This is the golang version of the clBag.cpp kernel file
   It contains the definition of lists, sets, tuples and arrays
*/

package Kernel

import (
	"fmt"
	// "strconv"
	// "unsafe"
)

// +---------------------------------------------------------------------------+
// |  Table of contents                                                        |
// |  Part 1: List Objects  (3 kinds)                                          |
// |  Part 2: Sets (using sorted bags - 3 kinds)                               |
// |  Part 3: Tuples (constant list)                                           |
// |  Part 4: Arays (fixed size lists)                                         |
// +---------------------------------------------------------------------------+

/// bag is the common root - we could reintroduce common methods ... but the data stuctures are
// completely different in go
// NOTE: the fact that nil is allowed and transformed into void is strange and may be deprecated
// we could fix this by ensuring that *.of is never nil  (CEMPTY is defined late in the boot)
func (l *ClaireBag) Of() *ClaireType {
	if l.of == nil {
		// fmt.Printf("=== conversion nil to EMPTY ==== \n")
		// only remove this once we have made sure that the test is not needed
		return ToType(CEMPTY.Id())        // changed in claire 4
	} else {
		return l.of
	}
}

func E_of_bag(l EID) EID { 
	return EID{ToBag(OBJ(l)).Of().Id(), 0} }

// this is the associated write method
func (l *ClaireBag) Cast_I(x *ClaireType) *ClaireBag {
	l.of = x
	return l
}
func E_cast_I_bag(l EID, x EID) EID {
	return EID{ToBag(OBJ(l)).Cast_I(ToType(OBJ(x))).Id(), 0}
}



// +---------------------------------------------------------------------------+
// |  Part 1: List Objects  (3 kinds)                                          |
// +---------------------------------------------------------------------------+

// list is implemented with four kinds: List (root), ListObject, ListInteger, ListFloat
// Note : the only one visible in CLAIRE are List and ListObject (typed)
// however slot generation is smart and uses ClaireList* as needed

// NOTE : the proper way to iterate a slice is with range
// for i := 0; i < elems; i++ {       -> correct but less idiomatic
// for i, v := range slice {          -> better

// constructors --------------------------------------------------------

// there are four kinds of lists : list (untyped) => implemented with listAny, object list => listAny, listInt, listFloat

// create a  list of size n and type t - not meant to be used outside this package
func createList(t *ClaireType, n int) *ClaireList {
	var l *ClaireList
	if t == ToType(C_integer.Id()) {
		l = ToList(makeEmptyIntegerList(n).Id())
	} else if t == ToType(C_float.Id()) {
		l = ToList(makeEmptyFloatList(n).Id())
	} else {
		l = ToList(makeEmptyObjectList(n).Id())
	}
	l.of = t
	return l
}

func makeNilList() *ClaireList {
	o := new(ClaireListObject)
	o.Isa = C_list
	o.Srange = C_object
	o.of = nil // constant empty list
	o.Values = []*ClaireAny{}
	return ToList(o.Id())
}

// create a typed empty list
func (t *ClaireType) EmptyList() *ClaireList {
	var l *ClaireList
	if t == ToType(C_integer.Id()) {
		l = ToList(makeEmptyIntegerList(0).Id())
	} else if t == ToType(C_float.Id()) {
		l = ToList(makeEmptyFloatList(0).Id())
	} else {
		l = ToList(makeEmptyObjectList(0).Id())
	}
	l.of = t
	return l
}

// in claire 4.0 this is accessible through empty_list(t)
func E_empty_list_type (t EID) EID { return EID{ToType(OBJ(t)).EmptyList().Id(),0}}

// three constructors for the compiler - preallocate the memory zone but no initialization (hence private methods)  
func makeEmptyIntegerList(n int) *ClaireListInteger {
	o := new(ClaireListInteger)
	o.Isa = C_list
	o.Srange = C_integer
	o.Values = make([]int, n)
	return o
}

func makeEmptyFloatList(n int) *ClaireListFloat {
	o := new(ClaireListFloat)
	o.Isa = C_list
	o.Srange = C_float
	o.Values = make([]float64, n)
	return o
}

func makeEmptyObjectList(n int) *ClaireListObject {
	o := new(ClaireListObject)
	o.Isa = C_list
	o.Srange = C_object
	o.Values = make([]*ClaireAny, n)
	return o
}

// special version for the compiler : created a typed list of size n
func CreateList(t *ClaireType, n int) *ClaireList {
	var l *ClaireList
	if t.Id() == CEMPTY.Id() {
		l = ToList(makeEmptyObjectList(n).Id())
	} else if t.Included(ToType(C_integer.Id())) == CTRUE {
		l = ToList(makeEmptyIntegerList(n).Id())
	} else if t.Included(ToType(C_float.Id())) == CTRUE {
		l = ToList(makeEmptyFloatList(n).Id())
	} else {
		l = ToList(makeEmptyObjectList(n).Id())
	}
	l.of = t
	return l
}

// ==== full constructor withough type checks   =====
func MakeList(t *ClaireType, args ...*ClaireAny) *ClaireList {
	n := len(args)
	l := t.EmptyList()
	for i := 0; i < n; i++ {
		l.AddFast(args[i])
	}
	return l
}

// create a non mutable list (of = EMPTY)
func MakeConstantList(args ...*ClaireAny) *ClaireList {
	n := len(args)
	l := ToType(CEMPTY.Id()).EmptyList()
	for i := 0; i < n; i++ {
		l.AddFast(args[i])
	}
	return l
}

// core version (used in Bootcore)
func coreList(t *ClaireType, args ...*ClaireAny) *ClaireList {
	n := len(args)
	l := makeEmptyObjectList(n)
	l.of = t
	for i := 0; i < n; i++ {
		l.Values[i] = args[i]
	}
	return ToList(l.Id())
}

// full constructor for int list
func MakeListInteger(args ...int) *ClaireList {
	n := len(args)
	ls := make([]int, n)
	for i := 0; i < n; i++ {
		ls[i] = args[i]
	}
	o := new(ClaireListInteger)
	o.Isa = C_list
	o.Srange = C_integer
	o.of = ToType(C_integer.Id())
	o.Values = ls
	return ToList(o.Id())
}

// full constructor for float list
func MakeListFloat(args ...float64) *ClaireList {
	n := len(args)
	ls := make([]float64, n)
	for i := 0; i < n; i++ {
		ls[i] = args[i]
	}
	o := new(ClaireListFloat)
	o.Isa = C_list
	o.Srange = C_float
	o.of = ToType(C_float.Id())
	o.Values = ls
	return ToList(o.Id())
}

// create an object list, using *.Id() pattern
func objectSlice(args ...*ClaireAny) []*ClaireAny {
	n := len(args)
	l := make([]*ClaireAny, n)
	for i := 0; i < n; i++ {
		l[i] = args[i]
	}
	return l
}

// another shorthand to create a list of class (equivalent of list::domain)
// takes *ClaireAny to simplify the code (works for classes and types alike)
// could become more generic ! like Slice or Many or List
func Signature(args ...*ClaireAny) []*ClaireAny {
	n := len(args)
	l := make([]*ClaireAny, n)
	for i := 0; i < n; i++ {
		l[i] = args[i].Id()
	}
	return l
}

// syntactic sugar : untyped list of *ClaireAny
func AnyList(args ...*ClaireAny) *ClaireListObject {
	n := len(args)
	l := make([]*ClaireAny, n)
	for i := 0; i < n; i++ {
		l[i] = args[i]
	}
	o := new(ClaireListObject)
	o.Isa = C_list
	o.Srange = C_object
	o.of = ToType(C_any.Id())
	o.Values = l
	return o
}

// copy a slice (3 flavors)
func copySlice(l []*ClaireAny) []*ClaireAny {
	l2 := make([]*ClaireAny, len(l)) // create a list of a given size
	copy(l2, l)                      // golang copy
	return l2
}

func copySliceInteger(l []int) []int {
	l2 := make([]int, len(l)) // create a list of a given size
	copy(l2, l)               // golang copy
	return l2
}

func copySliceFloat(l []float64) []float64 {
	l2 := make([]float64, len(l)) // create a list of a given size
	copy(l2, l)                   // golang copy
	return l2
}

// create a list from z slice
func makeListObject(t *ClaireType, l []*ClaireAny) *ClaireListObject {
	o := new(ClaireListObject)
	o.Isa = C_list
	o.Srange = C_object
	// fmt.Printf("create listObject %x isa:%x\n",o,o.Isa)
	o.of = t
	o.Values = l
	return o
}

// create a listInteger
func makeListInteger(l []int) *ClaireListInteger {
	o := new(ClaireListInteger)
	o.Isa = C_list
	o.Srange = C_integer
	o.of = &C_integer.ClaireType // alternate to .ToType
	o.Values = l
	return o
}

// create a listFloat
func makeListFloat(l []float64) *ClaireListFloat {
	o := new(ClaireListFloat)
	o.Isa = C_list
	o.Srange = C_float
	o.of = &C_float.ClaireType // ???
	o.Values = l
	return o
}

// equality on lists
func (l1 *ClaireList) equalList(l2 *ClaireList) *ClaireBoolean {
	if l1.Srange == C_object { // deep equality required
		n := len(l1.ValuesO())
		if l1.Srange != l2.Srange || n != len(l2.ValuesO())  {
			return CFALSE
		} else {
			for i := 0; i < n; i++ {
				if Equal(l1.ValuesO()[i], l2.ValuesO()[i]) == CFALSE {
					return CFALSE
				}
			}
			return CTRUE
		}
	} else { // float or int, shallow is enough
		n := len(l1.ValuesI())
		if l1.Srange != l2.Srange || n != len(l2.ValuesI()) {
			return CFALSE
		} else {
			for i := 0; i < n; i++ {
				if l1.ValuesI()[i] != l2.ValuesI()[i] {
					return CFALSE
				}
			}
			return CTRUE
		}
	}
}

// fast membership (we know that val is identified)
func (l1 *ClaireList) Memq(val *ClaireAny) *ClaireBoolean {
	if l1.Srange == C_object {
		for _, v := range l1.ValuesO() {
			if v == val {
				return CTRUE
			}
		}
		return CFALSE
	} else if l1.Srange == C_integer {
		if val.Isa != C_integer {
			return CFALSE
		}
		for _, v := range l1.ValuesI() {
			if v == ToInteger(val).Value {
				return CTRUE
			}
		}
		return CFALSE
	} else {
		if val.Isa != C_float {
			return CFALSE
		}
		for _, v := range l1.ValuesF() {
			if v == ToFloat(val).Value {
				return CTRUE
			}
		}
		return CFALSE
	}
}

// -----------------------------------------------------------------------------------------------
// API functions : each API function has its golang method version (for the compiler) and the f_ interpreted
// version that gets and returns EIDs
// from bag + from list

// copy a list ---------------------------------------------
func (l *ClaireList) Copy() *ClaireList {
	o := new(ClaireListObject)
	if l.of != nil {
		o.of = l.of
	} // should worlf with tuple = list with nil as l.of
	if l.Length() > 1000 {
		panic("bug - too long list - stop and look")
	}
	o.Isa = l.Isa // copy of an array is an array
	o.Srange = l.Srange
	if l.Srange == C_integer {
		o.toInteger().Values = copySliceInteger(l.ValuesI())
	} else if l.Srange == C_float {
		o.toFloat().Values = copySliceFloat(l.ValuesF())
	} else {
		o.toObject().Values = copySlice(l.ValuesO())
	}
	return ToList(o.Id())
}

// EID function
func E_copy_list(l EID) EID { return EID{ToList(OBJ(l)).Copy().Id(), 0} }

// acess to l[i] --------------------------------------------------------------
// needed by the compiler : reads an object value at position i, independently of the type representation
func (l *ClaireList) At(i int) *ClaireAny {
	if l.Srange == C_integer {
		return MakeInteger(l.ValuesI()[i]).Id()
	} else if l.Srange == C_float {
		return MakeFloat(l.ValuesF()[i]).Id()
	} else {
		return l.ValuesO()[i]
	}
}

// used to define nth_get ???
func (l *ClaireList) NthGet(i int) *ClaireAny { return l.At(i - 1) }

// EID function
func E_nth_get_list(l EID, i EID) EID { return ToList(OBJ(l)).NthGet(INT(i)).ToEID() }

// this is the proper method with bound checks
func (l *ClaireList) Nth(i int) EID { 
	if i <= 0 || i > l.Length() {
		return Cerror(41, MakeInteger(i).Id(), l.Id())
	}
	if l.Srange == C_integer {
		return EID{C__INT, IVAL(l.ValuesI()[i-1])}
	} else if l.Srange == C_float {
		return EID{C__FLOAT, FVAL(l.ValuesF()[i-1])}
	} else {
		return l.ValuesO()[i-1].ToEID()
	}
}


// EID function
func E_nth_list(x EID, i EID) EID { return ToList(OBJ(x)).Nth(INT(i)) }
	

// write onto l[i] --------------------------------------------------------------
// needed by the compiler : reawrites ds an object value at position i, independently of the type representation
func (l *ClaireList) PutAt(i int, y *ClaireAny) {
	if l.Srange == C_integer {
		l.ValuesI()[i] = ToInteger(y).Value
	} else if l.Srange == C_float {
		l.ValuesF()[i] = ToFloat(y).Value
	} else {
		l.ValuesO()[i] = y
	}
}

// EID form : same operation, avoids the EID to *Any conversion (with alloc)
func (l *ClaireList) PutAtEID(i int, y EID) {
	if l.Srange == C_integer {
		if y.PTR == C__INT {
			l.ValuesI()[i] = INT(y)
		} else {
			l.ValuesI()[i] = ToInteger(OBJ(y)).Value
		}
	} else if l.Srange == C_float {
		if y.PTR == C__FLOAT {
			l.ValuesF()[i] = FLOAT(y)
		} else {
			l.ValuesF()[i] = ToFloat(OBJ(y)).Value
		}
	} else {
		l.ValuesO()[i] = ANY(y)
	}
}

// used to define nth_put ???
func (l *ClaireList) NthPut(i int, y *ClaireAny) *ClaireAny {
	l.PutAt(i-1, y)
	return y
}

// EID function
func E_nth_put_list(l EID, i EID, y EID) EID {
	return ToList(OBJ(l)).NthPut(INT(i), ANY(y)).ToEID()
}

// this is the version that performs check
// moved from Core to Kernel
func (l *ClaireList) NthEqual(i int, y *ClaireAny) EID {
	if i <= 0 || i > l.Length() {
		return Cerror(41, MakeInteger(i).Id(), l.Id())
	}
	l.PutAt(i-1, y)
	return y.ToEID()
}

// EID uses the better version below
func E_nth_equal_list(x EID, i EID, y EID) EID {
	l := ToList(OBJ(x))
	j := INT(i)
	if j <= 0 || j > l.Length() {
		return Cerror(41, MakeInteger(j).Id(), l.Id())
	}
	return l.WriteEID(j, y)
}

// new in CLAIRE 4 : put with a type check : gets an EID and performs the test with l.Of
// Notice that the error (17) uses l.of and not l which may be ill-formed (being built)
func (l *ClaireList) WriteEID(i int, y EID) EID {
	if l.Srange == C_integer {
		if y.PTR == C__INT {
			if l.of.Id() == C_integer.Id() || l.of.Id() == CEMPTY.Id() || l.of.Contains(ANY(y)) == CTRUE {
				l.ValuesI()[i-1] = INT(y)
			} else {
				return Cerror(17, ANY(y), l.of.Id())
			}
		} else if l.of.Id() == CEMPTY.Id() || l.of.Contains(ANY(y)) == CTRUE {
			l.ValuesI()[i-1] = ToInteger(OBJ(y)).Value
		} else {
			return Cerror(17, ANY(y), l.of.Id())
		}
	} else if l.Srange == C_float {
		if y.PTR == C__FLOAT {
			if l.of.Id() == C_float.Id() || l.of.Id() == CEMPTY.Id() || l.of.Contains(ANY(y)) == CTRUE {
				l.ValuesF()[i-1] = FLOAT(y)
			} else {
				return Cerror(17, ANY(y), l.of.Id())
			}
		} else if l.of.Id() == CEMPTY.Id() || l.of.Contains(ANY(y)) == CTRUE {
			l.ValuesF()[i-1] = ToFloat(OBJ(y)).Value
		} else {
			return Cerror(17, ANY(y), l.of.Id())
		}
	} else {
		if l.of.Id() == CEMPTY.Id() || l.of.Contains(ANY(y)) == CTRUE {
			l.ValuesO()[i-1] = ANY(y)
		} else {
			return Cerror(17, ANY(y), l.of.Id())
		}
	}
	return y
}

// make an empty list with same type ------------------------
func (l *ClaireList) Empty() *ClaireList {
	o := new(ClaireListObject)
	o.of = l.of // check with nil
	o.Isa = C_list
	o.Srange = l.Srange
	if l.Srange == C_integer {
		o.toInteger().Values = []int{}
	} else if l.Srange == C_float {
		o.toFloat().Values = []float64{}
	} else {
		o.toObject().Values = []*ClaireAny{}
	}
	return ToList(o.Id())
}

// EID function
func E_empty_list(l EID) EID { return EID{ToList(OBJ(l)).Empty().Id(), 0} }

// delete the first occurence of an object in a list ------------------
//
func (l *ClaireList) Delete(x *ClaireAny) *ClaireList {
	// find a position
	n := len(l.ValuesI())
	i := 0
	if l.Srange == C_integer {
		ls := l.ValuesI()
		if x.Isa != C_integer {
			return l
		}
		for i = 0; i < n; i++ {
			if ls[i] == ToInteger(x).Value {
				break
			}
		}
		if i < n-1 {
			copy(ls[i:n-1], ls[i+1:n])
		}
		if i < n {
			l.toInteger().Values = ls[:n-1]
		}
	} else if l.Srange == C_float {
		ls := l.ValuesF()
		if x.Isa != C_float {
			return l
		}
		for i = 0; i < n; i++ {
			if ls[i] == ToFloat(x).Value {
				break
			}
		}
		if i < n-1 {
			copy(ls[i:n-1], ls[i+1:n])
		}
		if i < n {
			l.toFloat().Values = ls[:n-1]
		}
	} else {
		ls := l.ValuesO()
		for i = 0; i < n; i++ {
			if Equal(ls[i], x) == CTRUE {
				break
			}
		}
		if i < n-1 {
			copy(ls[i:n-1], ls[i+1:n])
		}
		if i < n {
			l.toObject().Values = ls[:n-1]
		}
	}
	return l
}

// EID function
func E_delete_list(l EID, x EID) EID { return EID{ToList(OBJ(l)).Delete(ANY(x)).Id(), 0} }

// length of a list (should be a macro ? len(l.Value)  ---------------------
func (l *ClaireList) Length() int { return len(l.ValuesO()) }
func E_length_list(l EID) EID     { return EID{C__INT, IVAL(len(ToList(OBJ(l)).ValuesO()))} }

// generic contains function ------------------------------
func (l *ClaireList) Contain_ask(val *ClaireAny) *ClaireBoolean {
	if val.Isa.Ident_ask == CTRUE || l.Srange == C_integer || l.Srange == C_float {
		return l.Memq(val)
	} else {
		for _, x := range l.ValuesO() {
			if Equal(x, val) == CTRUE {
				return CTRUE
			}
		}
		return CFALSE
	}
}

// EID function
func E_contain_ask_list(l EID, val EID) EID {
	return EID{ToList(OBJ(l)).Contain_ask(ANY(val)).Id(), 0}
}

// add an object to a list --------------------------------------------
// this function may return an error
// uses type member contain_ask_any
func (l *ClaireList) Add(x *ClaireAny) EID {
	if  l.of.Contains(x) == CFALSE {            // removed the test l.of != nil
		return Cerror(17, x, l.of.Id())
	} else {
		return l.AddFast(x).Id().ToEID()         // Any EID
	}
}

// this is optimized for the CLAIRE interpreter to avoid (EID -> ANY) translation
// this is a little bit long because of two optim
// (1) sort according to Srange to go directly to the low level version (avoids ANY)
// (2) perform the direct membership test for list<integer> and list<float>
func E_add_list(x EID, y EID) EID {
	l := ToList(OBJ(x))
	if l.Srange == C_integer {
		if y.PTR == C__INT {
			if l.of.Id() == C_integer.Id() || l.of.Id() == CEMPTY.Id() || l.of.Contains(ANY(y)) == CTRUE {
				F_add_listInteger(l.toInteger(), INT(y))
			} else {
				return Cerror(17, ANY(y), l.of.Id())
			}
		} else if l.of.Contains(ANY(y)) == CTRUE {
			F_add_listInteger(l.toInteger(), ToInteger(OBJ(y)).Value)
		} else {
			return Cerror(17, ANY(y), l.of.Id())
		}
	} else if l.Srange == C_float {
		if y.PTR == C__FLOAT {
			if l.of.Id() == C_float.Id() || l.of.Id() == CEMPTY.Id() || l.of.Contains(ANY(y)) == CTRUE {
				F_add_listFloat(l.toFloat(), FLOAT(y))
			} else {
				return Cerror(17, ANY(y), l.of.Id())
			}
		} else if l.of.Contains(ANY(y)) == CTRUE {
			F_add_listFloat(l.toFloat(), ToFloat(OBJ(y)).Value)
		} else {
			return Cerror(17, ANY(y), l.of.Id())
		}
	} else {
		if l.of.Contains(ANY(y)) == CTRUE {
			F_add_listObject(l.toObject(), ANY(y))
		} else {
			return Cerror(17, ANY(y), l.of.Id())
		}
	}
	return x
}

// this is the version without the check when we know that addition is type safe
func (l *ClaireList) AddFast(x *ClaireAny) *ClaireList {
	if l == CNIL {
		panic("trying to fuck CNIL")        // debug : to remove later !!!!
	}
	if l.Srange == C_integer {
		l.AddFastInteger(ToInteger(x).Value)
	} else if l.Srange == C_float {
		l.AddFastFloat(ToFloat(x).Value)
	} else {
		l.AddFastObject(x)
	}
	return l
}

// optimized forms used by the compiler
func (l *ClaireList) AddFastObject(x *ClaireAny)  {	l.toObject().Values = append(l.toObject().Values,x)}
func (l *ClaireList) AddFastInteger(x int)  {	l.toInteger().Values = append(l.toInteger().Values,x)}
func (l *ClaireList) AddFastFloat(x float64)  {	l.toFloat().Values = append(l.toFloat().Values,x)}

//deprecated
func F_add_listInteger(l *ClaireListInteger, x int) { l.Values = append(l.Values, x) }
func F_add_listFloat(l *ClaireListFloat, x float64) { l.Values = append(l.Values, x) }
func F_add_listObject(l *ClaireListObject, x *ClaireAny) {
	l.Values = append(l.Values, x)
}

// -------- exported as add! @ list --------------------------
// same without type check
func (l *ClaireList) Add_I(x *ClaireAny) *ClaireList {
	return l.AddFast(x)
}

// could be EID optimized to avoid ANY(val)
func E_add_I_list(l EID, val EID) EID { return EID{ToList(OBJ(l)).AddFast(ANY(val)).Id(), 0} }


// good old cons from lisp --------------------------------------------------
// returns a list of any (untyped)
func F_cons_any(x *ClaireAny, l *ClaireList) *ClaireList {
	n := len(l.ValuesO())
	ls := make([]*ClaireAny, n+1)
	ls[0] = x
	if l.Srange == C_integer {
		for j, v := range l.ValuesI() {
			ls[j+1] = AnyInteger(v)
		}
	} else if l.Srange == C_float {
		for j, v := range l.ValuesF() {
			ls[j+1] = AnyFloat(v)
		}
	} else {
		for j, v := range l.ValuesO() {
			ls[j+1] = v
		}
	}
	return ToList(makeListObject(ToType(CEMPTY.Id()), ls).Id())
}

// EID function
func E_cons_any(x EID, l EID) EID { return EID{F_cons_any(ANY(x), ToList(OBJ(l))).Id(), 0} }

// good old cdr from lisp - may return an error --------------------------------
func (l *ClaireList) Cdr() EID {
	n := len(l.ValuesO())
	if n == 0 {
		return Cerror(8, l.Id(), AnyInteger(0))
	} else if l.Srange == C_integer {
		ls := make([]int, n-1)
		copy(ls, l.ValuesI()[1:n])
		return EID{makeListInteger(ls).Id(), 0}
	} else if l.Srange == C_float {
		ls := make([]float64, n-1)
		copy(ls, l.ValuesF()[1:n])
		return EID{makeListFloat(ls).Id(), 0}
	} else {
		ls := make([]*ClaireAny, n-1)
		copy(ls, l.ValuesO()[1:n])
		return EID{makeListObject(l.of, ls).Id(), 0}
	}
}

func E_cdr_list(l EID) EID { return ToList(OBJ(l)).Cdr() }

// create a list of n things --------------------------------------------------------
func F_make_list_integer(n int, x *ClaireAny) *ClaireList {
	if x.Isa == C_integer {
		ls := make([]int, n)
		for i := 0; i < n; i++ {
			ls[i] = ToInteger(x).Value
		}
		return ToList(makeListInteger(ls).Id())
	} else if x.Isa == C_float {
		ls := make([]float64, n)
		for i := 0; i < n; i++ {
			ls[i] = ToFloat(x).Value
		}
		return ToList(makeListFloat(ls).Id())
	} else {
		ls := make([]*ClaireAny, n)
		for i := 0; i < n; i++ {
			ls[i] = x
		}
		return ToList(makeListObject(ToType(C_any.Id()), ls).Id())
	}
}

func E_make_list_integer(n EID, x EID) EID {
	return EID{F_make_list_integer(INT(n), ANY(x)).Id(), 0}
}

// position in a list
func F_index_list (l *ClaireList, x *ClaireAny) int {
	if l.Srange == C_integer {
		for i, y := range l.ValuesI() {
			if y == ToInteger(x).Value {
				return i + 1
			}
		}
	} else if l.Srange == C_float {
		for i, y := range l.ValuesF() {
			if y == ToFloat(x).Value {
				return i + 1
			}
		}
	} else {
		for i, y := range l.ValuesO() {
			if Equal(y, x) == CTRUE {
				return i + 1
			}
		}
	}
	return 0
}

func E_index_list(l EID, x EID) EID { 
	return EID{C__INT, uint64(F_index_list(ToList(OBJ(l)),ANY(x)))}
}

// nconc : destructive append
// add*(l1,l2) adds members of l2 => may return an error (hence EID as range)
func (l1 *ClaireList) Add_star(l2 *ClaireList) EID {
	if l2.Srange == C_integer {
		for _, x := range l2.ValuesI() {
			if l1.Srange == C_integer {
				F_add_listInteger(l1.toInteger(), x)
			} else {
				y := l1.Add(AnyInteger(x))
				if ErrorIn(y) {
					return y
				}
			}
		}
	} else if l2.Srange == C_float {
		for _, x := range l2.ValuesF() {
			if l1.Srange == C_float {
				F_add_listFloat(l1.toFloat(), x)
			} else {
				y := l1.Add(AnyFloat(x))
				if ErrorIn(y) {
					return y
				}
			}
		}
	} else {
		for _, x := range l2.ValuesO() {
			y := l1.Add(x)
			if ErrorIn(y) {
				fmt.Printf("adding %s to %s fails\n",x.Prt(),l1.Prt())
				return y
			}
		}
	}
	return EID{l1.Id(), 0}
}

func E_add_star_list(l1 EID, l2 EID) EID { return ToList(OBJ(l1)).Add_star(ToList(OBJ(l2))) }

// same code without the type checks
func (l1 *ClaireList) AddStarFast(l2 *ClaireList) *ClaireList {
	if l2.Srange == C_integer {
		for _, x := range l2.ValuesI() {
			if l1.Srange == C_integer {	F_add_listInteger(l1.toInteger(), x)
			} else { l1.AddFast(AnyInteger(x))}}
	} else if l2.Srange == C_float {
		for _, x := range l2.ValuesF() {
			if l1.Srange == C_float { F_add_listFloat(l1.toFloat(), x)
			} else { l1.AddFast(AnyFloat(x))}}
	} else {
		for _, x := range l2.ValuesO() {l1.AddFast(x)}}
	return l1
}

// non-destructive append : nconc to copy
// create a copy with either the same type or any
// does not produce an error
func (l1 *ClaireList) Append(l2 *ClaireList) *ClaireList {
	var l3 *ClaireList
	if l1.Srange == l2.Srange && (l1.Srange == C_float || l1.Srange == C_integer || l1.of == l2.of) {
		l3 = l1.Copy()
	} else {
		ls := make([]*ClaireAny, len(l1.ValuesO()))
		if l2.Srange == C_integer {
			for i, x := range l1.ValuesI() {
				ls[i] = AnyInteger(x)
			}
		} else if l2.Srange == C_float {
			for i, x := range l1.ValuesF() {
				ls[i] = AnyFloat(x)
			}
		} else {
			for i, x := range l1.ValuesO() {
				ls[i] = x
			}
			l3 = ToList(makeListObject(ToType(C_any.Id()), ls).Id())
		}
	}
	return l3.AddStarFast(l2)           // fast add* (nconc) no type error by construction
}

func E_append_list(l1 EID, l2 EID) EID {
	return EID{ToList(OBJ(l1)).Append(ToList(OBJ(l2))).Id(), 0}
}

// insert after a member at a given position, works only for a list
// may return an error
func (l *ClaireList) Nth_plus(n int, val *ClaireAny) EID {
	m := l.Length()
	if l.of == nil || l.of.Contains(val) == CFALSE {
		fmt.Printf("NTH+ fails with %s in %s\n", val.Prt(), l.Of().Prt())
		return Cerror(17, val, l.of.Id())
	}
	if n <= 0 || n > m+1 {
		return Cerror(5, AnyInteger(n), l.Id())
	}
	if l.Srange == C_integer {
		l.AddFastInteger(l.ValuesI()[m-1])
		ls := l.toInteger().Values
		for i := m - 1; i >= n; i-- {
			ls[i] = ls[i-1]
		}
		ls[n-1] = ToInteger(val).Value
	} else if l.Srange == C_float {
		l.AddFastFloat(l.ValuesF()[m-1])
		ls := l.toFloat().Values
		for i := m - 1; i >= n; i-- {
			ls[i] = ls[i-1]
		}
		ls[n-1] = ToFloat(val).Value
	} else {
		F_add_listObject(l.toObject(), l.ValuesO()[m-1])
		ls := l.toObject().Values // change !
		for i := m - 1; i >= n; i-- {
			ls[i] = ls[i-1]
		}
		ls[n-1] = val
	}
	return EID{l.Id(), 0}
}

func E_nth_plus_list(l EID, n EID, val EID) EID { return ToList(OBJ(l)).Nth_plus(INT(n), ANY(val)) }

// removes the nth member of a list
func (l *ClaireList) Nth_dash(n int) EID {
	m := l.Length()
	if n <= 0 || n > m+1 {
		return Cerror(5, AnyInteger(n), l.Id())
	}
	if l.Srange == C_integer {
		ls := l.ValuesI()
		if n < m-1 {
			copy(ls[n:m-1], ls[n+1:m])
		}
		l.toInteger().Values = ls[:m-1]
	} else if l.Srange == C_float {
		ls := l.ValuesF()
		if n < m-1 {
			copy(ls[n:m-1], ls[n+1:m])
		}
		l.toFloat().Values = ls[:m-1]
	} else {
		ls := l.ValuesO()
		if n < m-1 {
			copy(ls[n:m-1], ls[n+1:m])
		}
		l.toObject().Values = ls[:m-1]
	}
	return EID{l.Id(), 0}
}

func E_nth_dash_list(l EID, n EID) EID { return ToList(OBJ(l)).Nth_dash(INT(n)) }

// remove the n first elements of a list (skip_list)
func (l *ClaireList) Skip(n int) *ClaireList {
	m := l.Length()
	if n <= 0 {
		return l
	} else if l.Srange == C_integer {
		l.toInteger().Values = l.ValuesI()[n:m]
	} else if l.Srange == C_float {
		l.toFloat().Values = l.ValuesF()[n:m]
	} else {
		l.toObject().Values = l.ValuesO()[n:m]
	}
	return l
}

func E_skip_list(l EID, n EID) EID { return EID{ToList(OBJ(l)).Skip(INT(n)).Id(), 0} }

// shrinks to the n first element  -------------------------------------------------
// golang: if n < 0, n = 0 (avoid errors)
func (l *ClaireList) Shrink(n int) *ClaireList {
	m := l.Length()
	if n >= m {
		return l
	} else if n <= 0 {
		if l.Srange == C_integer {
			l.toInteger().Values = []int{}
		} else if l.Srange == C_float {
			l.toFloat().Values = []float64{}
		} else {
			l.toObject().Values = []*ClaireAny{}
		}
	} else if l.Srange == C_integer {
		l.toInteger().Values = l.toInteger().Values[0 : n]
	} else if l.Srange == C_float {
		l.toFloat().Values = l.toFloat().Values[0 : n]
	} else {
		l.toObject().Values = l.toObject().Values[0 : n]
	}
	return l
}

func E_shrink_list(l EID, n EID) EID { return EID{ToList(OBJ(l)).Shrink(INT(n)).Id(), 0} }

// a slice method in the honor of golang without errors ()
func (l *ClaireList) Slice(i int,j int) *ClaireList {
	m := l.Length()
    if i <= 0 {i = 1}
	if j >= m {j = m}
    if j < i {        // empty slice
		if l.Srange == C_integer {
			l.toInteger().Values = []int{}
		} else if l.Srange == C_float {
			l.toFloat().Values = []float64{}
		} else {
			l.toObject().Values = []*ClaireAny{}
		}
	} else if l.Srange == C_integer {
		l.toInteger().Values = l.toInteger().Values[i - 1 : j]
	} else if l.Srange == C_float {
		l.toFloat().Values = l.toFloat().Values[i - 1 : j]
	} else {
		l.toObject().Values = l.toObject().Values[i - 1 : j]
	}
	return l
}

func E_slice_list(l EID, i EID, j EID) EID { return EID{ToList(OBJ(l)).Slice(INT(i),INT(j)).Id(), 0} }

// used by the reader to change this information
func (l *ClaireList) Cast_I(x *ClaireType) *ClaireList {
	l.of = x
	return l
}

// +---------------------------------------------------------------------------+
// |  Part 2: Sets (sorted bags)                                               |
// +---------------------------------------------------------------------------+

// history note: we tried map[string]*any but is was slower (see the go bench results)
// We implement sets with the CLAIRE3.5 pattern, using ordered lists
// the key is the value for ints & floats, the adress for identified object and the class for others


// create a set - generic
func MakeSet(t *ClaireType, args ...*ClaireAny) *ClaireSet {
	n := len(args)
	l := t.EmptySet()
	for i := 0; i < n; i++ {
		l.AddFast(args[i])
	}
	return l
}

// specific version for integers 
func MakeSetInteger(args ...int) *ClaireSet {
	n := len(args)
	ls := make([]int, 2 * n)
	for i := 0; i < n; i++ {	ls[i] = args[i]}
	o := new(ClaireSetInteger)
	o.Isa = C_set
	o.Srange = C_integer	
	o.of = ToType(C_integer.Id())
	o.Values = ls
	o.Count = n
	return ToSet(o.Id())
}

// specific version for floats 
func MakeSetFloat(args ...float64) *ClaireSet {
	n := len(args)
	ls := make([]float64, 2 * n)
	for i := 0; i < n; i++ {	ls[i] = args[i]}
	o := new(ClaireSetFloat)
	o.Isa = C_set
	o.Srange = C_float
	o.of = ToType(C_float.Id())
	o.Values = ls
	o.Count = n
	return ToSet(o.Id())
}

// create a typed empty list
func (t *ClaireType) EmptySet() *ClaireSet {
	var l *ClaireSet
	if t == ToType(C_integer.Id()) {
		l = ToSet(makeEmptyIntegerSet(2).Id())
	} else if t == ToType(C_float.Id()) {
		l = ToSet(makeEmptyFloatSet(2).Id())
	} else {
		l = ToSet(makeEmptyObjectSet(2).Id())
	}
	l.of = t
	return l
}

// create a typed empty object list (needed in bootcore)
func (t *ClaireType) EmptySetObject() *ClaireSet {
	var l *ClaireSet
	l = ToSet(makeEmptyObjectSet(2).Id())
	l.of = t
	return l
}

// in claire 4.0 this is accessible through empty_list(t)
func E_empty_set_type (t EID) EID { return EID{ToType(OBJ(t)).EmptySet().Id(),0}}

// three constructors for the compiler - preallocate the memory zone but no initialization (hence private methods)  
func makeEmptyIntegerSet(n int) *ClaireSetInteger {
	o := new(ClaireSetInteger)
	o.Isa = C_set
	o.Srange = C_integer
	o.Values = make([]int, 2 * n)
	o.Count = 0
	return o
}

func makeEmptyFloatSet(n int) *ClaireSetFloat {
	o := new(ClaireSetFloat)
	o.Isa = C_set
	o.Srange = C_float
	o.Values = make([]float64, 2* n)
	o.Count = 0
	return o
}

func makeEmptyObjectSet(n int) *ClaireSetObject {
	o := new(ClaireSetObject)
	o.Isa = C_set
	o.Srange = C_object
	o.Values = make([]*ClaireAny, 2 * n)
	o.Count = 0
	return o
}

// we import the At method from ClaireList (same code but different location)
func (l *ClaireSet) At(i int) *ClaireAny {
	if l.Srange == C_integer {
		return MakeInteger(l.ValuesI()[i]).Id()
	} else if l.Srange == C_float {
		return MakeFloat(l.ValuesF()[i]).Id()
	} else {
		return l.ValuesO()[i]
	}
}


// ----------- other constructors

// create an typed empty set  => this should be called only once => unicity of empty set (CEMPTY)
func makeNilSet() *ClaireSet {
	o := ToSet(makeEmptyObjectSet(0).Id())
	o.of = ToType(o.Id()) // o.of = empty <=> constant set
	return o
}


// constructor (use the same pattern for all bags)
func MakeConstantSet(args ...*ClaireAny) *ClaireSet {
	n := len(args)
	l := ToType(CEMPTY.Id()).EmptySet()
	for i := 0; i < n; i++ {
		l.AddFast(args[i])
	}
	return l
}


// ----------------- Insertion : named AddFast  ---------

// add a new element to a set (without checking the type)
// tricky since it may return another set (as in the C++ implementation) to avoid changing empty
// this could be improved ... probably
func (s *ClaireSet) AddFast(x *ClaireAny) *ClaireSet {
	if ClEnv.Verbose == 12 {fmt.Printf("--- %p to %p\n",x,s)}
	if s.Srange == C_integer {return s.AddSetInteger(ToInteger(x).Value)
	} else if s.Srange == C_float {return s.AddSetFloat(ToFloat(x).Value)
	} else {
		var s2 *ClaireSet = s
		if s == CEMPTY { s2 = ToType(CEMPTY.Id()).EmptySet()}
		return s2.AddSetObject(x)}
 }


// Integer version : easy  (we switch the name to avoid confusion with AddFastInteger from lists)
// clone of CLAIRE 3.5 set insertion : dichotomic search of insertion point
func (s *ClaireSet) AddSetInteger(val int) *ClaireSet {
	var ls []int = ToSetInteger(s.Id()).Values
	n := ToSetInteger(s.Id()).Count
	// fmt.Printf(">>> add %d to set %s with size %d\n",val,s.Debug(),n)
	var i int = 0
	var j int = n - 1 
	for ((i + 1) < j)  {    // dichotomic seach - 30 years old fragment :)
		k := ((i + j) >> 1);       	// k is neither l or j
	    v := ls[k]                  // cashing ls[k] helps
		if v == val {return s
		} else if v < val {i = k
		} else {j = k}}
	if n == 0 { i = -1
	} else if ls[i] == val || ls[j] == val {return s
	} else if ls[i] > val { i = i - 1
	} else if ls[j] < val { i = j }
	if len(ls) == n {      // full => we allocate a double size
		ls2 := make([]int, 2 * n)
		for k := 0; k < n; k++ {ls2[k] = ls[k]}  // copy(ls2[:],ls[:]) is slower :)
		ToSetInteger(s.Id()).Values = ls2
		ls = ls2
	}
	for k := n; k > i+1; k-- {ls[k] = ls[k-1]}  // copy(ls[i+2:n+1],ls[i+1:n])
	ls[i+1] = val 
	ToSetInteger(s.Id()).Count = n + 1
	return s 
}

// same for Floats
func (s *ClaireSet) AddSetFloat(val float64) *ClaireSet {
	var ls []float64 = ToSetFloat(s.Id()).Values
	n := ToSetFloat(s.Id()).Count
	var i int = 0
	var j int = n - 1 
	for ((i + 1) < j)  {    // dichotomic seach - 30 years old fragment :)
		k := ((i + j) >> 1);       	// k is neither l or j
	    v := ls[k]                  // cashing ls[k] helps
		if v == val {return s
		} else if v < val {i = k
		} else {j = k}}
	if n == 0 { i = -1
	} else if ls[i] == val || ls[j] == val {return s
	} else if ls[i] > val { i = i - 1
	} else if ls[j] < val { i = j }
	if len(ls) == n {      // full => we allocate a double size
		ls2 := make([]float64, 2 * n)
		for k := 0; k < n; k++ {ls2[k] = ls[k]}  // copy(ls2[:],ls[:]) is slower :)
		ToSetFloat(s.Id()).Values = ls2
		ls = ls2
	}
	for k := n; k > i+1; k-- {ls[k] = ls[k-1]}  // copy(ls[i+2:n+1],ls[i+1:n])
	ls[i+1] = val 
	ToSetFloat(s.Id()).Count = n + 1
	return s 
}

// Object version  : more sophisticated
// SKEY(x) = adress(x) if x is identified, adress(x.isa) otherwise
// we look for the iso-KEY zone, make sure that x is not there using Equal
// then do an insertion at the end of the zone
func (s *ClaireSet) AddSetObject(val *ClaireAny) *ClaireSet {
	var ls []*ClaireAny = ToSetObject(s.Id()).Values
	n := ToSetObject(s.Id()).Count
	kval := SKEY(val)
	if ClEnv.Verbose == 13 { fmt.Printf("==> addSetObject %p into set:%d, key=%d\n",val,n,kval)}
	var i int = 0
	var j int = n - 1 
	for ((i + 1) < j)  {    // dichotomic seach - 30 years old fragment :)
		k := ((i + j) >> 1);       	// k is neither l or j
	    v := SKEY(ls[k])                  // cashing ls[k] helps
		if v == kval {               // we have found the key
			// fmt.Printf("==> addSetObject has found key=%d at position %d and content %p\n",val,k,ls[k])
			if val.Isa.Ident_ask == CTRUE {return s                  // hence we have found the object
			} else {
				var u int
				for u = k; u >= i && SKEY(ls[u]) == kval; u-- {
					if Equal(ls[u],val) == CTRUE {return s}}
				for u = k + 1; u <= j && SKEY(ls[u]) == kval; u++ {
					if Equal(ls[u],val) == CTRUE {return s}}
			    // fmt.Printf("found %p at position %d => u = %d",val,k,u)
				return s.AddInsertObject(val,u-1,n)    // j = upped bound of iso-zone + 1
				}
		} else if v < kval {i = k
		} else {j = k}}
	// the search stopped since i and j are adjacent
	if n == 0 { i = -1          // empty set
	} else if (val.Isa.Ident_ask == CTRUE && (ls[i] == val || ls[j] == val)) ||
	          (val.Isa.Ident_ask == CFALSE && (Equal(ls[i],val) == CTRUE || Equal(ls[j],val) == CTRUE)) {
	    // fmt.Printf("==> EQUAL MATCH with %p or %p versus val =%p\n",ls[i],ls[i],val)
		return s
	} else if SKEY(ls[i]) > kval { i = i - 1
	} else if SKEY(ls[j]) < kval { i = j }
	// fmt.Printf("==> addSet has not found the key and i = %d (size %d)\n",i,n)
	return s.AddInsertObject(val,i,n)              // insert after position i
}

// insertion at position i+1
// shift content from i+1 to n and update ls[i+1]
func (s *ClaireSet) AddInsertObject(val *ClaireAny, i int, n int) *ClaireSet {
	var ls []*ClaireAny = ToSetObject(s.Id()).Values
	// fmt.Printf("--> insert into set:%d at position %d\n",n,i)
	if len(ls) == n {                            // full => we allocate a double size
		ls2 := make([]*ClaireAny, 2 * n)
		for k := 0; k < n; k++ {ls2[k] = ls[k]}  // copy(ls2[:],ls[:]) is slower :)
		ToSetObject(s.Id()).Values = ls2
		ls = ls2
	}
	for k := n; k > i+1; k-- {ls[k] = ls[k-1]}  // copy(ls[i+2:n+1],ls[i+1:n])
	ls[i+1] = val 
	ToSetObject(s.Id()).Count = n + 1
	return s 
}

// the sort KEY
// Note : this could be made faster for string at the expense of other objects but hashing(adress) strings is tricky
func SKEY(x *ClaireAny) uint64 {
	if x.Isa.Ident_ask == CTRUE {return x.ui64()
	} else {return x.Isa.ui64()}
}

// ----------------- Membership  ---------

// the generic  contains function
func (s *ClaireSet) Contain_ask(x *ClaireAny) *ClaireBoolean {
     if s.Srange == C_integer {return s.ContainSetInteger(ToInteger(x).Value)
	 } else if s.Srange == C_float {return s.ContainSetFloat(ToFloat(x).Value)
	 } else {return s.ContainSetObject(x)}
}

// EID version uses EID switching to avoid int/float allocation
func E_contain_ask_set(s EID, val EID) EID {
	s2 := ToSet(OBJ(s))
	if s2.Srange == C_integer { return EID{s2.ContainSetInteger(INT(val)).Id(),0}
    } else if s2.Srange == C_float { return EID{s2.ContainSetFloat(FLOAT(val)).Id(),0}
	} else {return EID{s2.ContainSetObject(ANY(val)).Id(),0}}
}

// use the specific getIndex* method that returns -1 if not here or i if at position i (needed for deletion)
// simple dichotomic search
func (l *ClaireSet) ContainSetInteger(x int) *ClaireBoolean { 
	if l.getInteger(x) == -1 {return CFALSE
	} else {return CTRUE}}
func (l *ClaireSet) ContainSetFloat(x float64) *ClaireBoolean {
	if l.getFloat(x) == -1 {return CFALSE
		} else {return CTRUE}}	
func (l *ClaireSet) ContainSetObject(x *ClaireAny) *ClaireBoolean {
	if l.getObject(x) == -1 {return CFALSE
		} else {return CTRUE}}
	

// simple dichotomic search
// returns the go position : -1 if not here, or k such l[k]=x
func (l *ClaireSet) getInteger(x int) int {
	s := ToSetInteger(l.Id())
	if s.Count == 0 { return -1} 
	var i int = 0
	var j int = s.Count - 1
	for ((i + 1) < j)  {    			// dichotomic seach - 30 years old fragment :)
		k := ((i + j) >> 1)       		// k is neither l or j
		if s.Values[k] == x {return k
		} else if s.Values[k] < x {i = k
		} else {j = k}}
	if s.Values[i] == x {return i 
	} else if s.Values[j] == x {return j
	} else {return -1}
}

// simple dichotomic search : same for float !
func (l *ClaireSet) getFloat(x float64) int {
	s := ToSetFloat(l.Id())
	if s.Count == 0 { return -1} 
	var i int = 0
	var j int = s.Count - 1
	for ((i + 1) < j)  {    
		k := ((i + j) >> 1)       	
		if s.Values[k] == x {return k
		} else if s.Values[k] < x {i = k
		} else {j = k}}
	if s.Values[i] == x {return i 
	} else if s.Values[j] == x {return j
	} else {return -1}
}

// This is more sophisticated since the KEY for object changes according to x.Isa.Identified
// non identified => find the iso-Key zone and then do a linear search with Equal
// this code is very similar to AddFast
func (l *ClaireSet) getObject(val *ClaireAny) int {
	s := ToSetObject(l.Id())
	kval := SKEY(val)
	if ClEnv.Verbose == 13 { fmt.Printf("==>  getObject %p into set:%d, key=%d\n",val,s.Count,kval)}
	ls := s.Values
	if s.Count == 0 { return -1} 
	var i int = 0
	var j int = s.Count - 1
	for ((i + 1) < j)  {            
		k := ((i + j) >> 1)       	
		if SKEY(ls[k]) == kval {       // found the KEY !  -> sweep for the zone if not identified
			if ClEnv.Verbose == 13 { fmt.Printf("==>  found the key at position %d\n",k)}
			if val.Isa.Ident_ask == CTRUE {return k
			} else {
				for u := k; u >= i && SKEY(ls[u]) == kval; u-- {
					if Equal(ls[u],val) == CTRUE {return u}}
				for u := k + 1; u <= j && SKEY(ls[u]) == kval; u++ {
						if Equal(ls[u],val) == CTRUE {return u}}
				return -1 }			                 // not found in the iso-KEY zone		
		} else if SKEY(ls[k]) < kval {i = k
		} else {j = k}}
		if ClEnv.Verbose == 13 { fmt.Printf("==>  look at two bounds %s, %s\n",ls[i].Prt(),ls[j].Prt())}
	if Equal(ls[i],val) == CTRUE {return i
	} else if  Equal(ls[j],val) == CTRUE {return j
	} else {return -1}
}

// generic access to index: -1 if not here
func (s *ClaireSet) getIndex(x *ClaireAny) int {
	if s.Srange == C_integer {return s.getInteger(ToInteger(x).Value)
	} else if s.Srange == C_float {return s.getFloat(ToFloat(x).Value)
	} else {return s.getObject(x)}
}

// member methods ------------------------------------------------------


// set equality : same size + inclusion :)
func (s1 *ClaireSet) equalSet(s2 *ClaireSet) *ClaireBoolean {
	if ToSetObject(s1.Id()).Count != ToSetObject(s2.Id()).Count {
		return CFALSE
	} else {
		for k := 0; k < ToSetObject(s1.Id()).Count; k++ {
			if s2.Contain_ask(s1.At(k)) == CFALSE {
				return CFALSE
			}
		}
		return CTRUE
	}
}

// API functions ----------------------------------------------------------------

// generic contains function
func (s *ClaireSet) Size() int { return ToSetObject(s.Id()).Count}
func (s *ClaireSet) Length() int { return ToSetObject(s.Id()).Count}

func E_size_set(s EID) EID { return EID{C__INT, IVAL(ToSet(OBJ(s)).Size())} }



// adds a value into a set
func (l *ClaireSet) Add(x *ClaireAny) EID {
	if l.of.Contains(x) == CFALSE {     // removed l.of == nil || 
		return Cerror(17, x, l.of.Id())
	} else {
		return EID{l.AddFast(x).Id(), 0}
	}
}

// EID optimization added to avoid allocating floats or integers
func E_add_set(x EID, val EID) EID { 
	l := ToSet(OBJ(x))
	if l.Srange == C_integer { 
		 if val.PTR != C__INT { return Cerror(17, ANY(val), l.of.Id())
		 } else {return EID{ToSet(l.Id()).AddSetInteger(INT(val)).Id(), 0}}
	} else if l.Srange == C_float { 
		if val.PTR != C__FLOAT { return Cerror(17, ANY(val), l.of.Id())
		} else {return EID{ToSet(l.Id()).AddSetFloat(FLOAT(val)).Id(), 0}}
	} else {
		if l.of == nil || l.of.CONTAINS(val) == CFALSE {
			return Cerror(17, ANY(val), l.of.Id())
		} else {return EID{l.AddSetObject(ANY(val)).Id(), 0}}
		}}


// same without type check
func (l *ClaireSet) Add_I(x *ClaireAny) *ClaireSet {
	return l.AddFast(x)
}

func E_add_I_set(x EID, val EID) EID { 
	l := ToSet(OBJ(x))
    if l.Srange == C_integer { return EID{ToSet(l.Id()).AddSetInteger(INT(val)).Id(), 0}
	} else if l.Srange == C_float { return EID{ToSet(l.Id()).AddSetFloat(FLOAT(val)).Id(), 0}
	} else { return EID{l.AddSetObject(ANY(val)).Id(), 0}}
}

// copy a bag - in v0.01, we cannot opy nil or {} we return a generic empty
func (l *ClaireSet) Copy() *ClaireSet {
	s := l.Empty()
	for k := 0; k < l.Length(); k++ {
		s.AddFast(l.At(k))
	}
	return s
}

func E_copy_set(s EID) EID { return EID{ToSet(OBJ(s)).Copy().Id(), 0} }

// new in v3.1.16: create an empty copy with same type
func (l *ClaireSet) Empty() *ClaireSet {
	if l == CEMPTY {
		return CEMPTY
	} else {
		return l.of.EmptySet()
	}
}

func E_empty_set(s EID) EID { return EID{ToSet(OBJ(s)).Empty().Id(), 0} }

// ----------------  delete an object from a set is another complex method ----------------------------------

// two steps : find the index (similar to contains) + deleteAt
func (s *ClaireSet) Delete(x *ClaireAny) *ClaireSet {
	i := s.getIndex(x)
	if i >= 0 {s.deleteAt(i)}
	return s
}

func E_delete_set(l EID, val EID) EID { return EID{ToSet(OBJ(l)).Delete(ANY(val)).Id(), 0} }

// deletition is the inverse property of addition / we need to split into three 
func (s *ClaireSet) deleteAt(i int)  {
     if s.Srange == C_integer {s.deleteAtInteger(i)
	 } else if s.Srange == C_float {s.deleteAtFloat(i)
	 } else {s.deleteAtObject(i)}
}

// deletion at position i : shift content from i+1 to n 
func (s *ClaireSet) deleteAtInteger(i int)  {
	ls := ToSetInteger(s.Id()).Values
	n := ToSetInteger(s.Id()).Count
	for k := i + 1; k < n ; k++ {ls[k - 1] = ls[k]}  // 
	ToSetInteger(s.Id()).Count = n - 1
}
func (s *ClaireSet) deleteAtFloat(i int)  {
	ls := ToSetFloat(s.Id()).Values
	n := ToSetFloat(s.Id()).Count
	for k := i + 1; k < n ; k++ {ls[k - 1] = ls[k]}  // 
	ToSetFloat(s.Id()).Count = n - 1
}
func (s *ClaireSet) deleteAtObject(i int) {
	ls := ToSetObject(s.Id()).Values
	n := ToSetObject(s.Id()).Count
	for k := i + 1; k < n ; k++ {ls[k - 1] = ls[k]}  // 
	ToSetObject(s.Id()).Count = n - 1
}


// intersection of two sets --------------------------------------------------------------------------------
func F__exp_set(s1 *ClaireSet, s2 *ClaireSet) *ClaireSet {
	s := ToType(CEMPTY.Id()).EmptySet()
	for k := 0; k < s1.Length(); k++ {
		  x := s1.At(k)
		  if s2.Contain_ask(x) == CTRUE {s.AddFast(x)}
		}
	return s
}

func E__exp_set(s1 EID, s2 EID) EID { return EID{F__exp_set(ToSet(OBJ(s1)),ToSet(OBJ(s2))).Id(), 0} }

// union of two sets: merge of sorted lists (sort_of) */
func F_append_set(s1 *ClaireSet, s2 *ClaireSet) *ClaireSet {
	s := ToType(CEMPTY.Id()).EmptySet()
	for k := 0; k < s1.Length(); k++ {
		s.AddFast(s1.At(k))
	}
	for k := 0; k < s2.Length(); k++ {
		s.AddFast(s2.At(k))
	}
	return s
}

func E_append_set(s1 EID, s2 EID) EID {
	return EID{F_append_set(ToSet(OBJ(s1)), ToSet(OBJ(s2))).Id(), 0}
}

// create a set from a list => remove duplicates: very useful - function because Core method (vs Kernel)
func (l *ClaireList) Set_I() *ClaireSet {
	var s *ClaireSet
	if l.Srange == C_object {
		s = l.of.EmptySet()
		for _, v := range l.ValuesO() {
			s.AddSetObject(v)
		}
	} else if l.Srange == C_integer {
		s = ToType(C_integer.Id()).EmptySet()
		for _, v := range l.ValuesI() {
			s.AddSetInteger(v)
		}
	} else {
		s := ToType(C_float.Id()).EmptySet()
		for _, v := range l.ValuesF() {
			s.AddSetFloat(v)
		}
	}
	return s
}

func E_set_I_list(l EID) EID { return EID{ToList(OBJ(l)).Set_I().Id(), 0} }

// reciprocate : create a list from a set - function because Core method (vs Kernel)
func (s *ClaireSet) List_I() *ClaireList {
	l := s.of.EmptyList()
	for k := 0; k < s.Length(); k++  { l.AddFast(s.At(k))}
	return l
}

func E_list_I_set(s EID) EID { return EID{ToSet(OBJ(s)).List_I().Id(), 0} }

// returns a nice sequence of consecutive numbers
func F_sequence_integer(n int, m int) *ClaireSet {
	if m < n {
		return ToType(CEMPTY.Id()).EmptySet()
	} else {
		s := ToType(C_integer.Id()).EmptySet()
		for i := n; i <= m; i++ { s.AddSetInteger(i)}
		return s
	}
}

func E_sequence_integer(n EID, m EID) EID {
	return EID{F_sequence_integer(INT(n), INT(m)).Id(), 0}
}

// returns a nice sequence of consecutive numbers
func F_list_integer(n int, m int) *ClaireList {
	if m < n {
		return ToType(CEMPTY.Id()).EmptyList()
	} else {
		s := ToType(C_integer.Id()).EmptyList()
		for i := n; i <= m; i++ {
			s.AddFastInteger(i)
		}
		return s
	}
}

func E_list_integer(n EID, m EID) EID {
	return EID{F_list_integer(INT(n), INT(m)).Id(), 0}
}


// +---------------------------------------------------------------------------+
// |  Part 3: Tuples (constant list)                                           |
// +---------------------------------------------------------------------------+

// in Claire 3 & 4, tuples are list (without the ability to modify) without
// parameter types (they are values)
// tuples are also types: tuple(1,2) % tuple(integer,integer)
// with the Go implementation, we use ListObject as the support for tuples

// create a tuple skeleton
func MakeEmptyTuple() *ClaireTuple {
	l := new(ClaireListObject)
	l.Isa = C_tuple
	l.Srange = C_object
	l.of = nil
	l.Values = make([]*ClaireAny, 0)
	return ToTuple(l.Id())
}

// creates a tuple directly from the content
func MakeTuple(args ...*ClaireAny) *ClaireTuple {
	n := len(args)
	l := MakeEmptyTuple()
	for i := 0; i < n; i++ {
		l.AddFast(args[i])
	}
	return l
}

// API functions -----------------------------------------------------------------------------

/* copy a tuple ---------------------------------------------
func (l *ClaireTuple) Copy() *ClaireTuple {
	o := new(ClaireTuple)
	o.of = nil
	o.Isa = l.Isa
	o.Srange = l.Srange
	o.Values = copySlice(l.Values)
	return o
} */

func E_copy_tuple(l EID) EID { return EID{ToList(OBJ(l)).Copy().Id(), 0} }

// create a tuple from a list
func (l *ClaireList) Tuple_I() *ClaireTuple {
	x := MakeEmptyTuple()
	if l.Srange == C_integer {
		for _, y := range l.ValuesI() {
			x.AddFast(AnyInteger(y))
		}
	} else if l.Srange == C_float {
		for _, y := range l.ValuesF() {
			x.AddFast(AnyFloat(y))
		}
	} else {
		for _, y := range l.ValuesO() {
			x.AddFast(y)
		}
	}
	return x
}

func E_tuple_I_list(l EID) EID { return EID{ToList(OBJ(l)).Tuple_I().Id(), 0} }

// create a list<any> (updatable) from a tuple
func (l *ClaireTuple) List_I() *ClaireList {
	x := ToType(C_any.Id()).EmptyList()
	for _, y := range l.ValuesO() {
		x.AddFast(y)
	}
	return x
}

func E_list_I_tuple(l EID) EID { return EID{ToTuple(OBJ(l)).List_I().Id(), 0} }

/* add is reused from list.
func (l *ClaireTuple) AddFast(x *ClaireAny) {
	ToListObject(l).Values = append(l.Values, x)
}
*/

func E_addFast_tuple(l EID, x EID) EID {
	ToList(OBJ(l)).AddFast(ANY(x))
	return EVOID
}

// +---------------------------------------------------------------------------+
// |  Part 4: Arrays (fixed size lists)                                         |
// +---------------------------------------------------------------------------+

// think about it : arrays existed in CLAIRE 3 because of the native implementation
// which is no longer needed. An array is just a typed list whose type says that it does
// not support addition.

// because GO is a simple language without OOP, we use casts to (1) create the proper list then
// (2) cast to array

// creates a typed empty array (internal function)
func makeEmptyArray(n int, t *ClaireType) *ClaireList {
	var a *ClaireList
	if t.Included(ToType(C_integer.Id())) == CTRUE {
		o := new(ClaireListInteger)
		o.Srange = C_integer
		o.Values = make([]int, n)
		a = ToArray(o.Id())
	} else if t.Included(ToType(C_float.Id())) == CTRUE {
		o := new(ClaireListFloat)
		o.Srange = C_float
		o.Values = make([]float64, n)
		a = ToArray(o.Id())
	} else {
		o := new(ClaireListObject)
		o.Srange = C_object
		o.Values = make([]*ClaireAny, n)
		a = ToArray(o.Id())
	}
	a.Isa = C_array
	a.of = t
	return a
}

// claire-exported function: creates an array with a default value x
func F_make_array_integer(n int, t *ClaireType, x *ClaireAny) *ClaireList {
	l := makeEmptyArray(n, t)
	for i := 0; i < n; i++ {
		l.PutAt(i, x)
	}
	return l
}

func E_make_array_integer(n EID, t EID, x EID) EID {
	return EID{F_make_array_integer(INT(n), ToType(OBJ(t)), ANY(x)).Id(), 0}
}

// convert a list to an arry and reciprocately
func F_list_I_array(l *ClaireList) *ClaireList {
	l2 := l.Copy()
	l2.Isa = C_array
	return l2
}

func E_list_I_array(l EID) EID { return EID{F_list_I_array(ToList(OBJ(l))).Id(), 0} }

func (l *ClaireList) Array_I() *ClaireList {
	l2 := l.Copy()
	l2.Isa = C_list
	return l2
}

func E_array_I_list(l EID) EID { return EID{ToList(OBJ(l)).Array_I().Id(), 0} }

// ==== full constructor withough type checks : used by compiler   =====
func MakeArray(t *ClaireType, args ...*ClaireAny) *ClaireList {
	n := len(args)
	l := makeEmptyArray(n, t)
	for i := 0; i < n; i++ {
		l.PutAt(i, args[i])
	}
	return l
}
