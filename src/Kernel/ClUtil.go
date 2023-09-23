// ==================================================================
// microCLAIRE                                              CLAIRE 4
// golang microClaire Kernel - started on June 21st, 2020
//
//  Copyright (C) 2020-2022 Yves Caseau. All Rights Reserved.
//  Redistribution and use in source and binary forms are permitted
//  provided that source distribution retains this entire copyright
//  notice and comments.
//  THIS SOFTWARE IS PROVIDED AS IS AND WITHOUT ANY WARRANTY,
//  INCLUDING, WITHOUT LIMITATION, THE IMPLIED WARRANTIES OF
//  MERCHANTABILTY AND FITNESS FOR A PARTICULAR PURPOSE
//
// clUtil.go
// ==================================================================

/* This is the reflective description of the class kernel
 */

package Kernel

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
	"unicode/utf8"
)

// +---------------------------------------------------------------------------+
// |  Table of contents                                                        |
// |  Part 1: Types                                                            |
// |  Part 2: Utilities (print,p...)                                           |
// |  Part 3: Integer                                                          |
// |  Part 4: Float                                                            |
// +---------------------------------------------------------------------------+

// +---------------------------------------------------------------------------+
// |  Part 1: Types                                                            |
// +---------------------------------------------------------------------------+

// ================================================================================================
// new in CLAIRE 4 : the main two type methods (contains and includes) are in Kernel
// type :: Class | Interval | Set | subtype[t] | C[p:T] | t U t |

// these are the kernel methods  : Contains, Included, Member, At, Class_I, Meet

// this is the membership function (x % t) defined for types only (made extensible in CLAIRE)
func (t *ClaireType) Contains(x *ClaireAny) *ClaireBoolean {
	if (ClEnv.Verbose > 10){
		fmt.Printf("--- call good old contains on %s\n",t.Prt())
	}
	switch t.Isa {
	case C_class:
		if x.Isa.IsIn(ToClass(t.Id())) == CTRUE {
			return CTRUE
		} else {
			return CFALSE
		}
	case C_Interval:
		if x.Isa.IsIn(C_integer) == CTRUE &&
			ToInteger(x).Value >= To_Interval(t.Id()).Arg1 &&
			ToInteger(x).Value <= To_Interval(t.Id()).Arg2 {
			return CTRUE
		} else {
			return CFALSE
		}
	case C_set:
		return ToSet(t.Id()).Contain_ask(x)
	case C_list:                                       // since v3.2, list is accepted as type (although not advertised)
		return ToList(t.Id()).Contain_ask(x)
	case C_tuple:                                      // tuple as a type is not a list ! it is a template
		if x.Isa == t.Isa {
			if ToList(x).Length() == ToList(t.Id()).Length() {
				for i,v := range(ToList(t.Id()).ValuesO()) {
                   if ToType(v).Contains(ToList(x).ValuesO()[i]) == CFALSE {return CFALSE}
				}
                return CTRUE
			} else {return CFALSE}
		} else {return CFALSE}
	case C_subtype:
		if ToSubtype(t.Id()).Arg.Contains(x) == CFALSE {
			return CFALSE
		} else {
			return ToType(x).Included(ToSubtype(t.Id()).T1)
		}
	case C_Param:
		if x.Isa.IsIn(To_Param(t.Id()).Arg) == CTRUE {
			l := To_Param(t.Id()).Args
			for i, p := range To_Param(t.Id()).Params.ValuesO() {
				if ToType(l.ValuesO()[i]).containsType(ToProperty(p).Of(ToObject(x))) == CFALSE {
					// fmt.Printf("Param member failed because %s not in %s\n", ToProperty(p).Of(ToObject(x)).Prt(),l.ValuesO()[i].Prt())
					return CFALSE
				}
			}
			return CTRUE
		} else {
			return CFALSE
		}
	case C_Union:
		if To_Union(t.Id()).T1.Contains(x) == CTRUE {
			return CTRUE
		} else {
			return To_Union(t.Id()).T2.Contains(x)
		}
	// case C_Reference:
	//	return CTRUE
	default:           // this path exists since Types are extensible, but this method should not be called on such types
	    fmt.Printf("===================== Contains: VERY CHELOU ==================== \n")
	    fmt.Printf("===================== Contains:%s in %s ==================== \n",x.Prt(),t.Prt())
		// if ClEnv.Verbose == -1 {panic("stop and see why")}
		return CFALSE
	}
}

// equivalent of %type that was defined in types.cl in CLAIRE 3.5 
// special memberhip for t:set when y is a type (deep equality)
func (t *ClaireType) containsType (x *ClaireAny) *ClaireBoolean {
	if x.Isa.IsIn(C_type) == CTRUE && t.Isa == C_set {
		for k := 0; k < ToSet(t.Id()).Count; k++ {
	   		z := ToSet(t.Id()).At(k)
		   	if z.Isa.IsIn(C_type) == CTRUE && ToType(z).Included(ToType(x)) == CTRUE && ToType(x).Included(ToType(z)) == CTRUE {
			   return CTRUE  // deep type equality
		   }}
        return CFALSE
	} else {return t.Contains(x)}
} 

// special optimization for EID arg
func (t *ClaireType) CONTAINS(x EID) *ClaireBoolean {
	switch t.Isa {
	case C_class:
		if x.PTR == C__INT { return C_integer.isIn(ToClass(t.Id()))
		} else if x.PTR == C__FLOAT { return C_float.isIn(ToClass(t.Id()))
		} else if x.PTR == C__CHAR { return C_char.isIn(ToClass(t.Id()))
		} else {
			return x.PTR.Isa.IsIn(ToClass(t.Id()))
		}
	default:
		return t.Contains(ANY(x))
	}
}

// range check : verify that x (EID) belongs to t (Type)
func RangeCheck(t *ClaireType, x EID) EID {
	if ErrorIn(x) {return x
	} else if t.Contains(ANY(x)) == CTRUE {
		return x
	} else {
		return Cerror(12,ANY(x),t.Id())
	} 
}

// Included is subtyping based on set inclusion
// this is a major piece :)
// is is broken into cases (sub functions) by type, following the CLAIRE 3.5 definition
func (x *ClaireType) Included(y *ClaireType) *ClaireBoolean {
	switch x.Isa {
	case C_class:
		return ToClass(x.Id()).Included(y)
	case C_Interval:
		return To_Interval(x.Id()).Included(y)
	case C_subtype:
		return ToSubtype(x.Id()).Included(y)
	case C_Param:
		return To_Param(x.Id()).Included(y)
	case C_Union:
		return To_Union(x.Id()).Included(y)
	case C_set:
		return ToSet(x.Id()).Included(y)
	case C_list:
		return ToList(x.Id()).listIncluded(y)    // beware of tuple vs list
	case C_tuple:
		return ToTuple(x.Id()).tupleIncluded(y)   // one underlying GO structure for 2 Claire classes => 2 methods
//	case C_Reference:
//		return CTRUE
	default:
		return CFALSE
	}
}

// class version with a switch
func (x *ClaireClass) Included(y *ClaireType) *ClaireBoolean {
	switch y.Isa {
	case C_class:
		return x.IsIn(ToClass(y.Id()))
	case C_tuple:
		return CFALSE
	case C_Union:
		if x.Open == 0 { // abstract class x < y <=> xi < y
			for k := 0; k <= x.Subclass.Count; k++ {
				if ToClass(x.Subclass.At(k)).Included(y) == CFALSE {
					return CFALSE
				}
				return CTRUE
			}
		} else if x.Included(To_Union(y.Id()).T1) == CTRUE {
			return CTRUE
		} else {
			return x.Included(To_Union(y.Id()).T2)
		}
	case C_set:
		if x.Open == 0 && x.Subclass.Count == 0 { // enumerate: forall(u in x | u % y)
			for _, u := range x.Instances.ValuesO() {
				if ToSet(y.Id()).Contain_ask(u) == CFALSE {
					return CFALSE
				}
			}
			return CTRUE
		} else {
			return CFALSE
		}
	default:
		return x.Included2(y)
	}
	return CFALSE // stupid => bug with compiler (check periodically if this will re-occur)
}

// Union
func (x *ClaireUnion) Included(y *ClaireType) *ClaireBoolean {
	if x.T1.Included(y) == CFALSE {
		return CFALSE
	} else {
		return x.T2.Included(y)
	}
}

// Set : there is a difference between constant sets (of = {}) and others (that may grow)
func (x *ClaireSet) Included(y *ClaireType) *ClaireBoolean {
	t := x.of
	if t.Id() != CEMPTY.Id() {
		return t.Included(y)
	} else {
	for k := 0; k < x.Count; k++ {
		z := x.At(k)
		if y.containsType(z) == CFALSE {
			return CFALSE
		}
	}
	return CTRUE
   }
}

// Set : there is a difference between constant sets (of = {}) and others (that may grow)
func (x *ClaireList) listIncluded(y *ClaireType) *ClaireBoolean {
	t := x.of
	if t.Id() != CEMPTY.Id() {return t.Included(y)
	} else {
	n := x.Length()
	for i := 0; i < n ; i++  {
		if y.containsType(x.At(i)) == CFALSE {
			return CFALSE
		}
	}
	return CTRUE
   }
}

// Interval
func (x *ClaireInterval) Included(y *ClaireType) *ClaireBoolean {
	switch y.Isa {
	case C_Interval:
		if x.Arg1 >= To_Interval(y.Id()).Arg1 && x.Arg2 <= To_Interval(y.Id()).Arg2 {
			return CTRUE
		} else {
			return CFALSE
		}
	case C_set:
		for k := x.Arg1; k <= x.Arg2; k++ {
			if y.Contains(AnyInteger(k)) == CFALSE {
				return CFALSE
			}
		}
		return CTRUE
	case C_Union:
		if x.Included(To_Union(y.Id()).T1) == CTRUE {
			return CTRUE
		} else {
			return x.Included(To_Union(y.Id()).T2)
		}
	case C_class:
		return C_integer.IsIn(ToClass(y.Id()))
	default:
		{
			return x.Included2(y)
		}
	}
}

// subtype
func (x *ClaireSubtype) Included(y *ClaireType) *ClaireBoolean {
	if y.Isa == C_Param || y.Isa == C_subtype {
		return x.Included2(y)
	} else if y.Isa == C_Union {
		if x.Included(To_Union(y.Id()).T1) == CTRUE {
			return CTRUE
		} else {
			return x.Included(To_Union(y.Id()).T2)
		}
	} else {
		return x.Arg.Included(y)
	}
}

// Param is similar !
func (x *ClaireParam) Included(y *ClaireType) *ClaireBoolean {
	if y.Isa == C_Param || y.Isa == C_subtype {
		return x.Included2(y)
	} else if y.Isa == C_Union {
		if x.Included(To_Union(y.Id()).T1) == CTRUE {
			return CTRUE
		} else {
			return x.Included(To_Union(y.Id()).T2)
		}
	} else {
		return x.Arg.Included(y)
	}
}

// tuple : the only subtlety is the de-normalization of U within a tuple type
func (x *ClaireTuple) tupleIncluded(y *ClaireType) *ClaireBoolean {
	for i, t := range x.ValuesO() {
		if t.Isa == C_Union { // denormalize => create two new tuples
			x1 := x.Copy()
			x2 := x.Copy()
			x1.toObject().Values[i] = To_Union(t).T1.ToAny()
			x2.toObject().Values[i] = To_Union(t).T2.ToAny()
			if x1.Included(y) == CFALSE {
				return CFALSE // and
			} else {
				return x2.Included(y)
			}
		}
	}
	if y.Isa == C_tuple {
		if x.Length() != ToTuple(y.Id()).Length() {
			return CFALSE
		} else {
			for i, x1 := range x.ValuesO() {
				if ToType(x1.Id()).Included(ToType(ToTuple(y.Id()).ValuesO()[i])) == CFALSE {
					return CFALSE
				}
			}
			return CTRUE
		}
	} else if y.Isa == C_Union {
		if x.Included(To_Union(y.Id()).T1) == CTRUE {
			return CTRUE // or
		} else {
			return x.Included(To_Union(y.Id()).T2)
		}
	} else {
		return C_tuple.Included(y)
	}
}

// this is a generic ordering when y is a type Interval, a subtype or a Param
// x <= one such type is actually easy
func (x *ClaireType) Included2(y *ClaireType) *ClaireBoolean {
	switch y.Isa {
	case C_Param:
		if x.Included(ToType(To_Param(y.Id()).Arg.Id())) == CFALSE {
			return CFALSE // and
		} else {
			for i, p := range To_Param(y.Id()).Params.ValuesO() {
				if x.At(ToProperty(p)).Included(ToType(To_Param(y.Id()).Args.ValuesO()[i])) == CFALSE {
					return CFALSE
				}
			}
			return CTRUE
		}
	// case C_Reference:
	//	return CTRUE
	case C_subtype:
		if x.Included(ToType(ToSubtype(y.Id()).Arg.Id())) == CFALSE {
			return CFALSE // and
		} else {
			return x.Member().Included(ToType(ToSubtype(y.Id()).T1.Id()))
		}
	case C_Interval:
		return CFALSE
	default:
		return CFALSE
	}
}

// we need to type extractor functions which are auxiliary to subtyping
// t.Member  : a type for u, where u belongs to C and C is a collection of type t
// t.GetAt(p):   a type for p(x) if p is a of type t
// extract a member type, that is a valid type for all members (z) of instances of
// the type x.This is much simpler in v3.0
func (x *ClaireType) Member() *ClaireType {
	switch x.Isa {
	case C_class:
		if ToClass(x.Id()) == C_Interval {
			return ToType(C_integer.Id())
		} else {
			return ToType(C_any.Id())
		}
	case C_Union:
		return To_Union(x.Id()).T1.Member().Union(To_Union(x.Id()).T2.Member())
	case C_Interval:
		return ToType(CEMPTY.Id())
	case C_Param:
		return x.At(C_of).Member()
	case C_tuple:
		return x.toObject().Uall()
	case C_subtype:
		return ToSubtype(x.Id()).T1
	case C_set:
		{
			var rep *ClaireType = ToType(CEMPTY.Id())
			for k := 0; k < ToSet(x.Id()).Count; k++ {
				y := ToSet(x.Id()).At(k)
				if C_type.Contains(y) == CTRUE {
					rep = ToType(rep.Id()).Union(ToType(y.Id()))
				}
			}
			return rep
		}
	default:
		return ToType(CEMPTY.Id())
	}
}

// the method @ is used to extract the range information contained
// in a type. This method returns a type and is crucial for compiling !
func (x *ClaireType) At(p *ClaireProperty) *ClaireType {
	if ClEnv.Verbose > 10 {
		fmt.Printf("Call Type At on t:%s and p:%s\n", x.Prt(), p.Prt())
	}
	switch x.Isa {
	case C_class:
		r := p.findRestriction(ToClass(x.Id()))
		if r == nil {
			return ToType(C_any.Id())
		} else {
			return r.Range
		}
	case C_Param:
		i := F_index_list(To_Param(x.Id()).Params,p.ToAny())
		if i > 0 {
			return ToType(To_Param(x.Id()).Args.ValuesO()[i-1])
		} else {
			return To_Param(x.Id()).Arg.At(p)
		}
	case C_Union:
		return To_Union(x.Id()).T1.At(p).Union(To_Union(x.Id()).T2.At(p))
	case C_set:
		{
			var l *ClaireListObject = ToType(C_type.Id()).EmptyList().toObject()
			for k := 0; k < ToSet(x.Id()).Count; k++ {
				y := ToSet(x.Id()).At(k)
				l.AddFast(p.Of(ToObject(y)))
			}
			return l.Uall()
		}
	default:
		return x.Class_I().At(p)
	}
}

// exported to claire
func E__at_type(t EID, p EID) EID {
	return EID{ToType(OBJ(t)).At(ToProperty(OBJ(p))).Id(), 0}
}

// last but not least : Union constructor
// the union makes a partial reduction to the normal form. The complete
// reduction is done by enumeration if needed during the type subsumption
// union is left-associative: A U B U C is represented by (A U B) U C  => never(t2(x:Union) % union)
// a union of intervals is ALWAYS disjoint
func (x *ClaireType) Union(y *ClaireType) *ClaireType {
	if x.Isa == C_set {
		if y.Isa == C_set {
			return ToType(F_append_set(ToSet(x.Id()), ToSet(y.Id())).Id())
		} else {
			return y.Union(x)
		}
	} else if y.Included(x) == CTRUE {
		return x
	} else if x.Included(y) == CTRUE {
		return y
	} else if y.Isa == C_Union {
		return x.Union(To_Union(y.Id()).T1).Union(To_Union(y.Id()).T2)
	} else if x.Isa == C_Interval && y.Isa == C_Interval { // two intervals
		if To_Interval(y.Id()).Arg1-1 <= To_Interval(x.Id()).Arg2 &&
			To_Interval(x.Id()).Arg1 <= To_Interval(y.Id()).Arg1 {
			return ToType(MakeInterval(To_Interval(x.Id()).Arg1, To_Interval(y.Id()).Arg2).Id())
		} else if To_Interval(x.Id()).Arg1-1 <= To_Interval(y.Id()).Arg2 &&
			To_Interval(y.Id()).Arg1 <= To_Interval(x.Id()).Arg1 {
			return ToType(MakeInterval(To_Interval(y.Id()).Arg1, To_Interval(x.Id()).Arg2).Id())
		} else {
			return ToType(MakeUnion(x, y).Id())
		}
	} else if x.Isa == C_Union && y.Isa == C_Interval {
		z := To_Union(x.Id()).T2.Union(y)
		if z.Isa == C_Union {
			return ToType(MakeUnion(To_Union(x.Id()).T1.Union(y), To_Union(x.Id()).T2).Id())
		} else {
			return To_Union(x.Id()).T1.Union(z)
		}
	} else if x.Isa == C_Interval && y.Isa == C_set { // this is tricky : see if we can grow the interval because of y
		a := To_Interval(x.Id()).Arg1
		b := To_Interval(x.Id()).Arg2
		expand := false
		for ToSet(y.Id()).Contain_ask(AnyInteger(a-1)) == CTRUE {
			a = a - 1
			expand = true
		} // maximal extension
		for ToSet(y.Id()).Contain_ask(AnyInteger(b+1)) == CTRUE {
			b = b + 1
			expand = true
		}
		if expand {
			return ToType(MakeDUnion(ToType(MakeInterval(a, b).Id()), ToSet(y.Id())).Id()) // faster in claire 4: avoid recursion
		} else {
			return MakeDUnion(x, ToSet(y.Id()))
		}
	} else {
		return ToType(MakeUnion(x, y).Id())
	}
}

// make a union with a set / DU hypothesis implies that we remove duplicate values
func MakeDUnion(x *ClaireType, y *ClaireSet) *ClaireType {
	var y2 *ClaireSet = ToType(CEMPTY.Id()).EmptySet()
	for k := 0; k < y.Count; k ++ {
		z := y.At(k)
		if x.Contains(z) == CFALSE {
			y2.AddFast(z)
		}
	}
	return ToType(MakeUnion(x, ToType(y2.Id())).Id())
}

// Union constructor
func MakeUnion(x *ClaireType, y *ClaireType) *ClaireUnion {
	var o = new(ClaireUnion)
	o.Isa = C_Union
	o.T1 = x
	o.T2 = y
	return o
}

// Interval constructor
func MakeInterval(x int, y int) *ClaireInterval {
	var o = new(ClaireInterval)
	o.Isa = C_Interval
	o.Arg1 = x
	o.Arg2 = y
	return o
}

// a combined union - l is a list of types
func (l *ClaireListObject) Uall() *ClaireType {
	var rep *ClaireType = nil
	for i, x := range l.Values {
		if i == 0 {
			rep = ToType(x)
		} else {
			rep = rep.Union(ToType(x))
		}
	}
	return rep
}

// the best class approximation
func (x *ClaireTypeExpression) Class_I() *ClaireClass {
	switch x.Isa {
	case C_class:
		return ToClass(x.Id())
	case C_set:
		if ToSet(x.Id()).Count == 0 {
			return C_void
		} else {
			var rep *ClaireClass = nil
			for k := 0; k < ToSet(x.Id()).Count; k++ {
				y := ToSet(x.Id()).At(k)
				if rep == nil {
					rep = y.Isa
				} else {
					rep = rep.Meet(y.Isa)
				}
			}
			return rep
		}
	case C_Union:
		return To_Union(x.Id()).T1.Class_I().Meet(To_Union(x.Id()).T2.Class_I())
	case C_Interval:
		return C_integer
	case C_subtype: // (if (x.arg = subtype) any else x.arg),
		return ToSubtype(x.Id()).Arg
	// case C_Reference:
	//	return C_any
	case C_Param:
		return To_Param(x.Id()).Arg
	case C_tuple:
		return C_tuple
	default:
		return C_any
	}
}

func E_class_I_type(x EID) EID {
	return EID{ToTypeExpression(OBJ(x)).Class_I().Id(),0}
}

// extract a list of class (useful to get a signature)
func listClass(ltype []*ClaireAny) *ClaireList {
	o := makeEmptyObjectList(len(ltype))
	o.of = ToType(C_class.Id()) // list<class>
	for i, y := range ltype {
		o.Values[i] = ToType(y).Class_I().Id()
		// fmt.Printf("-- listClass adds %s \n", ToClass(y).Name.key)
	}
	return ToList(o.Id())
}

// class meet is next to trivial
func (c1 *ClaireClass) Meet(c2 *ClaireClass) *ClaireClass {
	if c1.IsIn(c2) == CTRUE {
		return c2
	} else if c2.IsIn(c1) == CTRUE {
		return c1
	} else {
		return c1.Superclass.Meet(c2)
	}
}

// +---------------------------------------------------------------------------+
// |  Part 2: Utilities (print,p...)                                           |
// +---------------------------------------------------------------------------+

// convert anything into a Boolean
func F_boolean_I_any(x *ClaireAny) *ClaireBoolean {
	if x.Isa == C_boolean {
		return ToBoolean(x)
	} else if x.Isa == C_list {
		if ToList(x).Length() == 0 {
			return CFALSE
		} else {
			return CTRUE
		}
	} else if x.Isa == C_set {
		if ToSet(x).Count == 0 {
			return CFALSE
		} else {
			return CTRUE
		}
	} else if x == CFALSE.Id() {
		return CFALSE // to remove later on
	} else {
		return CTRUE
	}
}

func E_boolean_I_any(x EID) EID { return EID{F_boolean_I_any(ANY(x)).Id(), 0} }

// (0) fun: print & show
// ========================

// Print for an EID
func PEID(x EID) string {
	if x.PTR == C__INT {
		return "IID:" + strconv.Itoa(INT(x))
	} else if x.PTR == C__FLOAT {
		return "FID:" + fmt.Sprintf("%f", FLOAT(x))
	} else if x.PTR == C__CHAR {
		return "CID:" + string(CHAR(x))
	} else if x.VAL == 1 {
		return "ERROR:" + OBJ(x).Prt()
	} else {
		return "OID:" + OBJ(x).Prt()
	}
}

// ========  this is a debug print function that returns a string  ==============================
// keep this even if not in use + enrich it progressively, key for debugging generated code =====
func (p *ClaireAny) Prt() string {
	// fmt.Printf("call Prt() on %p with class %p \n", p, p.Isa)
	if p.Isa.Isa != C_class {
		fmt.Printf("Prt(%p) -> class is %p\n", p, p.Isa)
		panic("cannot Prt an object vérolé")
	}
	if p.Isa == C_integer {
		return strconv.Itoa(ToInteger(p).Value)
	} else if p.Isa == C_float {
		return fmt.Sprintf("%f", ToFloat(p).Value)
	} else if p.Isa == C_char {
		return fmt.Sprintf("'%c'", ToChar(p).Value)
	} else if p.Isa == C_class {
		return ToClass(p).Name.key                         // debug + fmt.Sprintf(":%p", p)
	} else if p.Isa == C_symbol {
		return ToSymbol(p).module_I.Name.key + "/" + ToSymbol(p).key
	} else if p.Isa == C_unbound_symbol {
		return "unbound@" + ToUnboundSymbol(p).Name.key
	} else if p.Isa.isIn(C_thing) == CTRUE {
		if utf8.RuneCountInString(ToThing(p).Name.key) > 1000 {
			panic("garbled name")
		}
		return ToThing(p).Name.key
	} else if p.Isa.isIn(C_list) == CTRUE {
		if p == CNIL.Id() {
			return "CNIL"
		} else {
			return ToList(p).PrtList()
		}
	} else if p.Isa == C_set {
		return ToSet(p).PrtSet()
	} else if p.Isa == C_method {
		return ToMethod(p).Selector.Prt() + "@" + ToMethod(p).Domain.prtIn()
	} else if p.Isa == C_slot {
		return ToSlot(p).Selector.Prt() + "#" + ToSlot(p).Domain.At(0).Prt()
	} else if p.Isa == C_subtype {
		return ToSubtype(p).Arg.Prt() + "[" + ToSubtype(p).T1.Prt() + "]"
	} else if p.Isa == C_Param {
		return To_Param(p).Arg.Prt() + "<" + To_Param(p).Params.prtIn() + ":" + To_Param(p).Args.prtIn() + ">"
	} else if p.Isa == C_string {
		return "\"" + ToString(p).Value + "\""
	} else if p == CTRUE.Id() {
		return "CTRUE"
	} else if p == CFALSE.Id() {
		return "CFALSE"
	} else if p == CNULL {
		return "CNULL"
	} else {
		return "<" + p.Isa.Name.key + ":" + fmt.Sprintf("%p", p) + ">"
	}
}

// printList : native version, with and without trimmings
func (l *ClaireList) PrtList() string { 
	if l.of == nil || l.of.Id() == CEMPTY.Id()  {return "list" + l.prtIn()
	} else {return "list<" + l.of.Prt() + ">" + l.prtIn()}
}

func (l *ClaireList) prtIn() string {
	res := "("
	n := l.Length()
	for i := 0; i < n; i++ {
		if i > 0 {
			res = res + ","
		}
		res = res + l.At(i).Prt()
	}
	return res + ")"
}


// Print Set
func (s *ClaireSet) PrtSet() string {
	if s.of == nil || s.of.Id() == CEMPTY.Id()  {return s.prtIn()
	} else {return "set<" + s.of.Prt() + ">" + s.prtIn()}
}

func (s *ClaireSet) prtIn() string {
	res := "{"
	first := true
	for k := 0; k < s.Count; k++ {
		if first {
			first = false
		} else {
			res = res + ","
		}
		res = res + s.At(k).Prt()
	}
	return res + "}"
}

// display a test
func (x *ClaireAny) Test(s string) {
	fmt.Printf("%s => %s.\n", s, x.Prt())
}

// Look() is like Prt() but uses a default "indepth" for objects C{s:v,*}
// beware of circular structure
func (p *ClaireAny) Look() string {
	if p.Isa.IsIn(C_object) == CFALSE ||  p.Isa.IsIn(C_bag) == CTRUE || 
	   p.Isa.IsIn(C_thing) == CTRUE ||  p.Isa.IsIn(C_class) == CTRUE  {return p.Prt()
	} else {
		c := p.Isa
		n := c.Slots.Length()
		var m int = n
		var s string = c.Name.key + "{"
		if n > 4 {m = 4}
		for i := 1; i < n; i++ {
			r := ToSlot(c.Slots.ValuesO()[i]).Srange
			s = s + ToObject(p).Get(i+1,r).Look()
			if i < (n - 1) {s = s +","
		    } else if (n > m) {s = s +"..."}
		}
		return s + "}"
	}}

//   reads a string on the keyboad - used to manage breakpoints in inspect.cl [read module]
func F_CommandLoopVoid() *ClaireString {
	i := 0
	c := make([]byte, 1)
	buff := make([]byte, 100)
	// os.Stdin.Flush()
	F_flush_port(ClEnv.Cout)
	os.Stdin.Read(c)
	for (c[0] != 10 && c[0] != 114) || i == 0 { // \n or \r
		buff[i] = c[0]
		i++
		os.Stdin.Read(c)
	}
	return MakeString(string(buff[0:i]))
}

func E_CommandLoopVoid(s EID) EID {
	return EID{F_CommandLoopVoid().Id(), 0}
}

// +---------------------------------------------------------------------------+
// |  Part 3: Integer                                                          |
// +---------------------------------------------------------------------------+

// usefull to check overflow
const (
	CLMAXFLOAT = 2.30583e+18
	CLMINFLOAT = -2.30583e+18
)

// test integer equality without MakeInteger allocation (compiler sweetener)
func (x *ClaireAny) IsInt(i int) bool {	return x.Isa == C_integer && ToInteger(x).Value == i}

// new: return the current date
// i is meant to give a few options (TBD)
func F_date_I_integer(i int) *ClaireString {
	dt := time.Now()
	if i == 1 {
		return MakeString(dt.Format("Monday 01-02-2006"))
	} else if i == 2  {
		return MakeString(dt.Format("15:04:05"))
	} else {
	return MakeString(dt.Format("Monday 01-02-2006 15:04:05"))}
}

func E_date_I_integer(i EID) EID { return EID{F_date_I_integer(INT(i)).Id(), 0} }

// princ a  int / float
func F_princ_integer(n int) { 
	ClEnv.Cout.PutInteger(n) }
func E_princ_integer(c EID) EID {
	BadI(c,"princ_integer")
	F_princ_integer(INT(c))
	return EVOID
}

// random  (notice the extension to 0, compared to Go, for upward compatibility)
func F_random_integer(n int) int {
	if n == 0 {return 0
	}  else {return rand.Intn(n)}
}
func E_random_integer(c EID) EID { return EID{C__INT, IVAL(F_random_integer(INT(c)))} }

// random_I_integer resets the seed
func F_random_I_integer(n int) {
	rand.Seed(int64(n))
}

func E_random_I_integer(c EID) EID { 
	F_random_I_integer(INT(c))
	return EVOID}

// used locally
func minInt(x int, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}

// arithmetic functions
func F_ch_sign(n int) int {
	return -n
}
func E_ch_sign(x EID) EID { return EID{C__INT, IVAL(F_ch_sign(INT(x)))} }

// safe division
func F__7_integer(n int, m int) EID {
	if m == 0 {
		return Cerror(20, MakeInteger(n).Id(), MakeInteger(0).Id())
	} else {
		return EID{C__INT, IVAL(n / m)}
	}
}
func E__7_integer(n EID, m EID) EID { return F__7_integer(INT(n), INT(m)) }

func F_mod_integer(n int, m int) EID {
	if m == 0 {
		return Cerror(20, MakeInteger(n).Id(), MakeInteger(0).Id())
	} else {
		return EID{C__INT, IVAL(n % m)}
	}
}
func E_mod_integer(n EID, m EID) EID { return F_mod_integer(INT(n), INT(m)) }

// v3.3.16: use float exponentiation and check overflow
// this is not enough for 64 bits integer !!
func F__exp_integer(n int, m int) EID {
	a := float64(n)
	b := float64(m)
	c := math.Pow(a, b)
	if c < CLMINFLOAT || c > CLMAXFLOAT {
		return Cerror(40, MakeInteger(n).Id(), MakeInteger(m).Id())
	} else {
		return EID{C__INT, IVAL((int)(math.Round(c)))}
	}
}
func E__exp_integer(n EID, m EID) EID { return F__exp_integer(INT(n), INT(m)) }

// return a power of 2
func F__exp2_integer(n int) EID {
	if (n >= 0) && (n <= 63) {
		return EID{C__INT, IVAL(1 << n)}
	} else {
		return Cerror(19, MakeInteger(n).Id(), MakeInteger(0).Id())
	}
}
func E__exp2_integer(n EID) EID { return F__exp2_integer(INT(n)) }

// gives a safe multiplication
func F_times_integer(n int, m int) EID {
	a := float64(n)
	b := float64(m)
	c := a * b
	if c < CLMINFLOAT || c > CLMAXFLOAT {
		return Cerror(40, MakeInteger(n).Id(), MakeInteger(m).Id())
	} else {
		return EID{C__INT, IVAL((int)(math.Round(c)))}
	}
}

func E_times_integer(n EID, m EID) EID { return F_times_integer(INT(n), INT(m)) }

func F__sup_equal_integer(n int, m int) *ClaireBoolean {
	if n >= m {
		return CTRUE
	} else {
		return CFALSE
	}
}

func E__sup_equal_integer(self EID, x EID) EID {
	return EID{F__sup_equal_integer(INT(self), INT(x)).Id(), 0}
}


// used by the compiler
func BitVectorContains(s int, x int) *ClaireBoolean {
	if ((s >> x) & 1) == 1 {
		return CTRUE
	} else {
		return CFALSE
	}
}

// +---------------------------------------------------------------------------+
// |  Part 4: Float                                                            |
// +---------------------------------------------------------------------------+

// regular princ with the proper precision
func F_princ_float(n float64) { ClEnv.Cout.PutFloat(n) }
func E_princ_float(c EID) EID {
	F_princ_float(FLOAT(c))
	return EVOID
}

// print is supposed to be nicer (check)
func F_print_float(n float64) { ClEnv.Cout.PrettyFloat(n) }
func E_print_float(c EID) EID {
	F_print_float(FLOAT(c))
	return EVOID
}

// there are two forms for float
func F_print_format_float(d float64, n int) {
	ClEnv.Cout.PutFormat(d, n)
}
func E_print_format_float(c EID, n EID) EID {
	F_print_format_float(FLOAT(c), INT(n))
	return EVOID
}

// ^(self:float,x:float) : float
func F__exp_float(self float64, x float64) float64 {
	return math.Pow(self, x)
}
func E__exp_float(self EID, x EID) EID {
	return EID{C__FLOAT, FVAL(F__exp_float(FLOAT(self), FLOAT(x)))}
}

// square root
func F_sqrt_float(self float64) float64 {
	return math.Sqrt(self)
}
func E_sqrt_float(self EID) EID { return EID{C__FLOAT, FVAL(F_sqrt_float(FLOAT(self)))} }

// log(self:float) : float
func F_log_float(self float64) float64 {
	return math.Log(self)
}
func E_log_float(self EID) EID { return EID{C__FLOAT, FVAL(F_log_float(FLOAT(self)))} }

// sin(self:float) : float
func F_sin_float(self float64) float64 {
	return math.Sin(self)
}
func E_sin_float(self EID) EID { return EID{C__FLOAT, FVAL(F_sin_float(FLOAT(self)))} }

// cos(self:float) : float
func F_cos_float(self float64) float64 {
	return math.Cos(self)
}
func E_cos_float(self EID) EID { return EID{C__FLOAT, FVAL(F_cos_float(FLOAT(self)))} }

// atan(self:float) : float
func F_atan_float(self float64) float64 {
	return math.Atan(self)
}
func E_atan_float(self EID) EID { return EID{C__FLOAT, FVAL(F_atan_float(FLOAT(self)))} }

// makes an integer into a float */
func F_to_float(n int) float64 { return float64(n) }
func E_to_float(self EID) EID  { return EID{C__FLOAT, FVAL(F_to_float(INT(self)))} }

// create a  claire integer from a claire float */
func F_integer_I_float(n float64) EID {
	if n < CLMINFLOAT || n > CLMAXFLOAT {
		return Cerror(39, MakeFloat(n).Id(), MakeInteger(0).Id())
	} else {
		return EID{C__INT, IVAL((int)(math.Floor(n)))}        // v4.0.6
	}
} // v3.3

func E_integer_I_float(self EID) EID { return F_integer_I_float(FLOAT(self)) }

// the classical order comparisons for two float
func F__inf_float(n float64, m float64) *ClaireBoolean {
	if n < m {
		return CTRUE
	} else {
		return CFALSE
	}
}
func E__inf_float(self EID, x EID) EID { return EID{F__inf_float(FLOAT(self), FLOAT(x)).Id(), 0} }

func F__inf_equal_float(n float64, m float64) *ClaireBoolean {
	if n <= m {
		return CTRUE
	} else {
		return CFALSE
	}
}
func E__inf_equal_float(self EID, x EID) EID {
	return EID{F__inf_equal_float(FLOAT(self), FLOAT(x)).Id(), 0}
}

func F__sup_float(n float64, m float64) *ClaireBoolean {
	if n > m {
		return CTRUE
	} else {
		return CFALSE
	}
}
func E__sup_float(self EID, x EID) EID { return EID{F__sup_float(FLOAT(self), FLOAT(x)).Id(), 0} }

func F__sup_equal_float(n float64, m float64) *ClaireBoolean {
	if n >= m {
		return CTRUE
	} else {
		return CFALSE
	}
}

func E__sup_equal_float(self EID, x EID) EID {
	return EID{F__sup_equal_float(FLOAT(self), FLOAT(x)).Id(), 0}
}
