/***** CLAIRE Compilation of file /Users/ycaseau/Dropbox/src/clairev4.0/src/meta/types.cl 
         [version 4.0.02 / safety 5] Saturday 10-30-2021 *****/

package Core
import (_ "fmt"
	. "Kernel"
)

//+-------------------------------------------------------------+
//| CLAIRE                                                      |
//| types.cl                                                    |
//| Copyright (C) 1994 - 2021 Yves Caseau. All Rights Reserved  |
//| cf. copyright info in file object.cl: about()               |
//+-------------------------------------------------------------+
// --------------------------------------------------------------------
// This file contains the definition of the CLAIRE type system (a true lattice).
// that is used both at compile- and at run-time.
// --------------------------------------------------------------------
// ******************************************************************
// *  Table of contents                                             *
// *    Part 1: Common Set Methods                                  *
// *    Part 2: definition of the type operators                    *
// *    Part 3: Interface methods                                   *
// *    Part 4: Lattice methods                                     *
// *    Part 5: Type methods                                        *
// ******************************************************************
// *********************************************************************
// *   Part 1: Common Set Methods                                      *
// *********************************************************************
// ----------------------- useful methods ------------------------------
/* {1} OPT.The go function for: finite?(self:type) [] */
func F_finite_ask_type (self *ClaireType ) *ClaireBoolean  { 
    // procedure body with s =  
var Result *ClaireBoolean  
    if (C_set.Id() == self.Isa.Id()) /* If:2 */{ 
      Result = CTRUE
      /* If!2 */}  else if (self.Isa.IsIn(C_list) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0234 *ClaireList   = ToList(self.Id())
        /* noccur = 1 */
        /* Let:4 */{ 
          var g0237UU *ClaireAny  
          /* noccur = 1 */
          /* For:5 */{ 
            var t *ClaireAny  
            _ = t
            g0237UU= CFALSE.Id()
            var t_support *ClaireList  
            t_support = g0234
            t_len := t_support.Length()
            for i_it := 0; i_it < t_len; i_it++ { 
              t = t_support.At(i_it)
              if (ToBoolean(OBJ(F_CALL(C_finite_ask,ARGS(t.ToEID())))) != CTRUE) /* If:7 */{ 
                 /*v = g0237UU, s =any*/
g0237UU = CTRUE.Id()
                break
                } else {
                
                /* If-7 */} 
              /* loop-6 */} 
            /* For-5 */} 
          Result = F_not_any(g0237UU)
          /* Let-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (C_class.Id() == self.Isa.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var g0235 *ClaireClass   = ToClass(self.Id())
        /* noccur = 2 */
        /* Let:4 */{ 
          var n int  = g0235.Open
          /* noccur = 4 */
          /* or:5 */{ 
            var v_or5 *ClaireBoolean  
            
            v_or5 = Equal(MakeInteger(n).Id(),MakeInteger(ClEnv.Open).Id())
            if (v_or5 == CTRUE) {Result = CTRUE
            } else /* or:6 */{ 
              v_or5 = Equal(MakeInteger(n).Id(),MakeInteger(ClEnv.Final).Id())
              if (v_or5 == CTRUE) {Result = CTRUE
              } else /* or:7 */{ 
                v_or5 = Equal(MakeInteger(n).Id(),ANY(F_CALL(C_Core_closed,ARGS(EID{ClEnv.Id(),0}))))
                if (v_or5 == CTRUE) {Result = CTRUE
                } else /* or:8 */{ 
                  /* and:9 */{ 
                    var v_and9 *ClaireBoolean  
                    
                    v_and9 = Equal(MakeInteger(n).Id(),MakeInteger(ClEnv.ABSTRACT).Id())
                    if (v_and9 == CFALSE) {v_or5 = CFALSE
                    } else /* arg:10 */{ 
                      /* Let:11 */{ 
                        var g0238UU *ClaireAny  
                        /* noccur = 1 */
                        /* For:12 */{ 
                          var c *ClaireAny  
                          _ = c
                          g0238UU= CFALSE.Id()
                          for _,c = range(g0235.Subclass.Values)/* loop:13 */{ 
                            if (F_finite_ask_type(ToType(c)) != CTRUE) /* If:14 */{ 
                               /*v = g0238UU, s =any*/
g0238UU = CTRUE.Id()
                              break
                              /* If-14 */} 
                            /* loop-13 */} 
                          /* For-12 */} 
                        v_and9 = F_not_any(g0238UU)
                        /* Let-11 */} 
                      if (v_and9 == CFALSE) {v_or5 = CFALSE
                      } else /* arg:11 */{ 
                        v_or5 = CTRUE/* arg-11 */} 
                      /* arg-10 */} 
                    /* and-9 */} 
                  if (v_or5 == CTRUE) {Result = CTRUE
                  } else /* or:9 */{ 
                    Result = CFALSE/* org-9 */} 
                  /* org-8 */} 
                /* org-7 */} 
              /* org-6 */} 
            /* or-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      } else {
      Result = CFALSE
      /* If-2 */} 
    return Result} 
  
// The EID go function for: finite? @ type (throw: false) 
func E_finite_ask_type (self EID) EID { 
    return EID{/*(sm for finite? @ type= boolean)*/ F_finite_ask_type(ToType(OBJ(self)) ).Id(),0}} 
  
// making a set from an abstract_set  (CLAIRE 4 : bag is not longer a concrete type)
// this is a list since order matters in enumeration
/* {1} OPT.The go function for: enumerate(self:any) [] */
func F_enumerate_any (self *ClaireAny ) EID { 
    var Result EID 
    if (self.Isa.IsIn(C_list) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0239 *ClaireList   = ToList(self)
        /* noccur = 1 */
        Result = EID{g0239.Id(),0}
        /* Let-3 */} 
      /* If!2 */}  else if (C_set.Id() == self.Isa.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var g0240 *ClaireSet   = ToSet(self)
        /* noccur = 1 */
        Result = EID{g0240.List_I().Id(),0}
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(C_array) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0241 *ClaireList   = ToArray(self)
        /* noccur = 1 */
        Result = EID{F_list_I_array(g0241).Id(),0}
        /* Let-3 */} 
      /* If!2 */}  else if (C_class.Id() == self.Isa.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var g0242 *ClaireClass   = ToClass(self)
        /* noccur = 1 */
        /* Let:4 */{ 
          var l *ClaireList   = ToType(C_object.Id()).EmptyList()
          /* noccur = 3 */
          /* For:5 */{ 
            var c *ClaireAny  
            _ = c
            for _,c = range(g0242.Descendents.Values)/* loop:6 */{ 
              l = l.Append(ToClass(c).Instances)
              /* loop-6 */} 
            /* For-5 */} 
          Result = EID{l.Id(),0}
          /* Let-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(C_Interval) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0243 *ClaireInterval   = To_Interval(self)
        /* noccur = 2 */
        Result = EID{F_list_integer(g0243.Arg1,g0243.Arg2).Id(),0}
        /* Let-3 */} 
      /* If!2 */}  else if (C_integer.Id() == self.Isa.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var g0244 int  = ToInteger(self).Value
        /* noccur = 1 */
        Result = EID{F_make_set_integer(g0244).List_I().Id(),0}
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(C_collection) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0245 *ClaireCollection   = ToCollection(self)
        /* noccur = 1 */
        /* Let:4 */{ 
          var g0247UU *ClaireAny  
          /* noccur = 1 */
          var g0247UU_try02485 EID 
          g0247UU_try02485 = F_CALL(C_set_I,ARGS(EID{g0245.Id(),0}))
          /* ERROR PROTECTION INSERTED (g0247UU-Result) */
          if ErrorIn(g0247UU_try02485) {Result = g0247UU_try02485
          } else {
          g0247UU = ANY(g0247UU_try02485)
          Result = EID{ToSet(g0247UU).List_I().Id(),0}
          }
          /* Let-4 */} 
        /* Let-3 */} 
      } else {
      Result = ToException(C_general_error.Make(MakeString("[178] cannot enumerate ~S").Id(),MakeConstantList(self).Id())).Close()
      /* If-2 */} 
    return Result} 
  
// The EID go function for: enumerate @ any (throw: true) 
func E_enumerate_any (self EID) EID { 
    return /*(sm for enumerate @ any= EID)*/ F_enumerate_any(ANY(self) )} 
  
// =type? is an operation (equality on types)
/* {1} OPT.The go function for: =type?(self:type,ens:type) [] */
func F__equaltype_ask_any (self *ClaireType ,ens *ClaireType ) *ClaireBoolean  { 
    // use function body compiling 
return  MakeBoolean((self.Included(ens) == CTRUE) && (ens.Included(self) == CTRUE))
    } 
  
// The EID go function for: =type? @ type (throw: false) 
func E__equaltype_ask_any (self EID,ens EID) EID { 
    return EID{/*(sm for =type? @ type= boolean)*/ F__equaltype_ask_any(ToType(OBJ(self)),ToType(OBJ(ens)) ).Id(),0}} 
  
// finds the sort associated to a type
/* {1} OPT.The go function for: sort!(x:type) [] */
func F_sort_I_type (x *ClaireType ) *ClaireClass  { 
    // procedure body with s =  
var Result *ClaireClass  
    if (C_class.Id() == x.Isa.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var g0249 *ClaireClass   = ToClass(x.Id())
        /* noccur = 1 */
        Result = g0249.Sort_I()
        /* Let-3 */} 
      } else {
      Result = x.Class_I().Sort_I()
      /* If-2 */} 
    return Result} 
  
// The EID go function for: sort! @ type (throw: false) 
func E_sort_I_type (x EID) EID { 
    return EID{/*(sm for sort! @ type= class)*/ F_sort_I_type(ToType(OBJ(x)) ).Id(),0}} 
  
// the membership for classes
/* {1} OPT.The go function for: %(self:any,ens:class) [] */
func F__Z_any1 (self *ClaireAny ,ens *ClaireClass ) *ClaireBoolean  { 
    // use function body compiling 
if (self.Isa.IsIn(ens) == CTRUE) /* body If:2 */{ 
      return  CTRUE
      } else {
      return  CFALSE
      /* body If-2 */} 
    } 
  
// The EID go function for: % @ list<type_expression>(any, class) (throw: false) 
func E__Z_any1 (self EID,ens EID) EID { 
    return EID{/*(sm for % @ list<type_expression>(any, class)= boolean)*/ F__Z_any1(ANY(self),ToClass(OBJ(ens)) ).Id(),0}} 
  
//
// v4.0 : belong is the unique method (static call for any) for membership
// replaces belong_to + member? in claire 3 => works on everything, collections and integer as well :)
// see belong_exp in gexp.cl to see how it is used + open-conding patterns
// note that belong may create an error => heavier => optimize with %t when possible 
/* {1} OPT.The go function for: belong(x:any,y:any) [] */
func F_BELONG (x *ClaireAny ,y *ClaireAny ) EID { 
    var Result EID 
    if (C_class.Id() == y.Isa.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var g0251 *ClaireClass   = ToClass(y)
        /* noccur = 1 */
        Result = EID{F__Z_any1(x,g0251).Id(),0}
        /* Let-3 */} 
      /* If!2 */}  else if (y.Isa.IsIn(C_list) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0252 *ClaireList   = ToList(y)
        /* noccur = 1 */
        Result = EID{g0252.Contain_ask(x).Id(),0}
        /* Let-3 */} 
      /* If!2 */}  else if (C_set.Id() == y.Isa.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var g0253 *ClaireSet   = ToSet(y)
        /* noccur = 1 */
        Result = EID{ToType(g0253.Id()).Contains(x).Id(),0}
        /* Let-3 */} 
      /* If!2 */}  else if (y.Isa.IsIn(C_array) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0254 *ClaireList   = ToArray(y)
        /* noccur = 1 */
        Result = EID{ToList(g0254.Id()).Contain_ask(x).Id(),0}
        /* Let-3 */} 
      /* If!2 */}  else if (y.Isa.IsIn(C_type_operator) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0255 *ClaireTypeOperator   = ToTypeOperator(y)
        /* noccur = 1 */
        Result = EID{ToType(g0255.Id()).Contains(x).Id(),0}
        /* Let-3 */} 
      /* If!2 */}  else if (C_integer.Id() == y.Isa.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var g0256 int  = ToInteger(y).Value
        /* noccur = 1 */
        if (C_integer.Id() == x.Isa.Id()) /* If:4 */{ 
          /* Let:5 */{ 
            var g0257 int  = ToInteger(x).Value
            /* noccur = 1 */
            Result = EID{BitVectorContains(g0256,g0257).Id(),0}
            /* Let-5 */} 
          } else {
          Result = EID{CFALSE.Id(),0}
          /* If-4 */} 
        /* Let-3 */} 
      } else {
      /* Let:3 */{ 
        var start int  = ClEnv.Index
        /* noccur = 2 */
        ClEnv.Push(x.ToEID())
        ClEnv.Push(y.ToEID())
        /* Let:4 */{ 
          var m *ClaireObject   = F_find_which_property(ToProperty(C__Z.Id()),start,x.Isa)
          /* noccur = 3 */
          var g0261I *ClaireBoolean  
          if (C_method.Id() == m.Isa.Id()) /* If:5 */{ 
            /* Let:6 */{ 
              var g0260 *ClaireMethod   = ToMethod(m.Id())
              /* noccur = 2 */
              g0261I = MakeBoolean((g0260.Domain.Length() == 2) && (g0260.Domain.ValuesO()[2-1] != C_any.Id()))
              /* Let-6 */} 
            } else {
            g0261I = CFALSE
            /* If-5 */} 
          if (g0261I == CTRUE) /* If:5 */{ 
            Result = F_eval_message_property(ToProperty(C__Z.Id()),m,start,CTRUE)
            } else {
            Result = ToException(C_general_error.Make(MakeString("[179] (~S % ~S): not implemented!").Id(),MakeConstantList(x,y).Id())).Close()
            /* If-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* If-2 */} 
    return Result} 
  
// The EID go function for: belong @ any (throw: true) 
func E_BELONG (x EID,y EID) EID { 
    return /*(sm for belong @ any= EID)*/ F_BELONG(ANY(x),ANY(y) )} 
  
// x % y is a short cut 
// CLAIRE4 : cannot be a macro (too early)
/* {1} OPT.The go function for: %(x:any,y:any) [] */
func F_belong_to (x *ClaireAny ,y *ClaireAny ) EID { 
    var Result EID 
    Result = F_BELONG(x,y)
    return Result} 
  
// The EID go function for: % @ list<type_expression>(any, any) (throw: true) 
func E_belong_to (x EID,y EID) EID { 
    return /*(sm for % @ list<type_expression>(any, any)= EID)*/ F_belong_to(ANY(x),ANY(y) )} 
  
// ****************************************************************
// *         Part 2: definition of the type operators             *
// ****************************************************************
// in CLAIRE4, types are defined in the Kernel go module
// type_operator <: type()
// union of two types ---------------------------------------------
// Disjonctive Union Axiom (DU): Each union (A U B) is stricly disjunctive:
//       (1) A ^B = 0
//       (2) x < A U B <=> x < A or x < B
// Producing disjunction union is a form of normalization (the previous notion
// of diustributivity was a lousy bug)
// DU Axiom is necessary to make <= and ^ easier to define
// This is achieved in the U method
/* {1} OPT.The go function for: self_print(self:Union) [] */
func F_self_print_Union_Core (self *ClaireUnion ) EID { 
    var Result EID 
    PRINC("(")
    Result = F_print_any(self.T1.Id())
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(" U ")
    Result = F_print_any(self.T2.Id())
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(")")
    Result = EVOID
    }}
    return Result} 
  
// The EID go function for: self_print @ Union (throw: true) 
func E_self_print_Union_Core (self EID) EID { 
    return /*(sm for self_print @ Union= EID)*/ F_self_print_Union_Core(To_Union(OBJ(self)) )} 
  
/* {1} OPT.The go function for: finite?(self:Union) [] */
func F_finite_ask_Union (self *ClaireUnion ) *ClaireBoolean  { 
    // use function body compiling 
return  MakeBoolean((ToBoolean(OBJ(F_CALL(C_finite_ask,ARGS(EID{self.T1.Id(),0})))) == CTRUE) && (ToBoolean(OBJ(F_CALL(C_finite_ask,ARGS(EID{self.T2.Id(),0})))) == CTRUE))
    } 
  
// The EID go function for: finite? @ Union (throw: false) 
func E_finite_ask_Union (self EID) EID { 
    return EID{/*(sm for finite? @ Union= boolean)*/ F_finite_ask_Union(To_Union(OBJ(self)) ).Id(),0}} 
  
// Intervals of integers ----------
/* {1} OPT.The go function for: self_print(self:Interval) [] */
func F_self_print_Interval_Core (self *ClaireInterval ) EID { 
    var Result EID 
    PRINC("(")
    Result = F_print_any(MakeInteger(self.Arg1).Id())
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(" .. ")
    Result = F_print_any(MakeInteger(self.Arg2).Id())
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(")")
    Result = EVOID
    }}
    return Result} 
  
// The EID go function for: self_print @ Interval (throw: true) 
func E_self_print_Interval_Core (self EID) EID { 
    return /*(sm for self_print @ Interval= EID)*/ F_self_print_Interval_Core(To_Interval(OBJ(self)) )} 
  
/* {1} OPT.The go function for: finite?(self:Interval) [] */
func F_finite_ask_Interval (self *ClaireInterval ) *ClaireBoolean  { 
    // use function body compiling 
return  CTRUE
    } 
  
// The EID go function for: finite? @ Interval (throw: false) 
func E_finite_ask_Interval (self EID) EID { 
    return EID{/*(sm for finite? @ Interval= boolean)*/ F_finite_ask_Interval(To_Interval(OBJ(self)) ).Id(),0}} 
  
// true constructor
/* {1} OPT.The go function for: --(x:integer,y:integer) [] */
func F__dash_dash_integer (x int,y int) EID { 
    var Result EID 
    if (x <= y) /* If:2 */{ 
      Result = EID{F__dot_dot_integer(x,y).Id(),0}
      } else {
      Result = ToException(C_general_error.Make(MakeString("[182] the interval (~S -- ~S) is empty").Id(),MakeConstantList(MakeInteger(x).Id(),MakeInteger(y).Id()).Id())).Close()
      /* If-2 */} 
    return Result} 
  
// The EID go function for: -- @ integer (throw: true) 
func E__dash_dash_integer (x EID,y EID) EID { 
    return /*(sm for -- @ integer= EID)*/ F__dash_dash_integer(INT(x),INT(y) )} 
  
// Parameterized class. -------------------------------------------
/* {1} OPT.The go function for: self_print(self:Param) [] */
func F_self_print_Param_Core (self *ClaireParam ) EID { 
    var Result EID 
    if ((self.Params.Length() == 1) && 
        ((self.Params.At(1-1) == C_of.Id()) && 
          (C_set.Id() == self.Args.At(1-1).Isa.Id()))) /* If:2 */{ 
      Result = F_print_any(self.Arg.Id())
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC("<")
      /* Let:3 */{ 
        var g0263UU *ClaireAny  
        /* noccur = 1 */
        var g0263UU_try02644 EID 
        g0263UU_try02644 = F_the_type(ToType(self.Args.At(1-1)))
        /* ERROR PROTECTION INSERTED (g0263UU-Result) */
        if ErrorIn(g0263UU_try02644) {Result = g0263UU_try02644
        } else {
        g0263UU = ANY(g0263UU_try02644)
        Result = F_CALL(C_print,ARGS(g0263UU.ToEID()))
        }
        /* Let-3 */} 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(">")
      Result = EVOID
      }}
      } else {
      Result = F_print_any(self.Arg.Id())
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC("[")
      /* Let:3 */{ 
        var i int  = 1
        /* noccur = 6 */
        /* Let:4 */{ 
          var g0262 int  = self.Args.Length()
          /* noccur = 1 */
          Result= EID{CFALSE.Id(),0}
          for (i <= g0262) /* while:5 */{ 
            var void_try6 EID 
            _ = void_try6
            { 
            if (i != 1) /* If:6 */{ 
              PRINC(", ")
              /* If-6 */} 
            void_try6 = F_CALL(C_print,ARGS(self.Params.At(i-1).ToEID()))
            /* ERROR PROTECTION INSERTED (void_try6-void_try6) */
            if ErrorIn(void_try6) {Result = void_try6
            break
            } else {
            PRINC(":(")
            void_try6 = F_CALL(C_print,ARGS(self.Args.At(i-1).ToEID()))
            /* ERROR PROTECTION INSERTED (void_try6-void_try6) */
            if ErrorIn(void_try6) {Result = void_try6
            break
            } else {
            PRINC(")")
            void_try6 = EVOID
            }}
            {
            }
            {
            i = (i+1)
            }
            /* while-5 */} 
          }
          /* Let-4 */} 
        /* Let-3 */} 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC("]")
      Result = EVOID
      }}
      /* If-2 */} 
    return Result} 
  
// The EID go function for: self_print @ Param (throw: true) 
func E_self_print_Param_Core (self EID) EID { 
    return /*(sm for self_print @ Param= EID)*/ F_self_print_Param_Core(To_Param(OBJ(self)) )} 
  
/* {1} OPT.The go function for: finite?(self:Param) [] */
func F_finite_ask_Param (self *ClaireParam ) *ClaireBoolean  { 
    // use function body compiling 
return  F_finite_ask_type(ToType(self.Arg.Id()))
    } 
  
// The EID go function for: finite? @ Param (throw: false) 
func E_finite_ask_Param (self EID) EID { 
    return EID{/*(sm for finite? @ Param= boolean)*/ F_finite_ask_Param(To_Param(OBJ(self)) ).Id(),0}} 
  
// subtype[X] ----------------------------------------------
// subtype[X] = {u in type | u <= t}
// for closure purposes, we add an arg Y -> Y inter st[X]
// Y can be any type class, but we forbid parametrisation on such classes !
// thus we can ensure that Y is a class
/* {1} OPT.The go function for: self_print(self:subtype) [] */
func F_self_print_subtype_Core (self *ClaireSubtype ) EID { 
    var Result EID 
    if (self.Arg.Id() == C_type.Id()) /* If:2 */{ 
      PRINC("subtype[")
      Result = F_print_any(self.T1.Id())
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC("]")
      Result = EVOID
      }
      } else {
      Result = F_print_any(self.Arg.Id())
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC("[")
      Result = F_print_any(self.T1.Id())
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC("]")
      Result = EVOID
      }}
      /* If-2 */} 
    return Result} 
  
// The EID go function for: self_print @ subtype (throw: true) 
func E_self_print_subtype_Core (self EID) EID { 
    return /*(sm for self_print @ subtype= EID)*/ F_self_print_subtype_Core(ToSubtype(OBJ(self)) )} 
  
// v3.2
/* {1} OPT.The go function for: finite?(self:subtype) [] */
func F_finite_ask_subtype (self *ClaireSubtype ) *ClaireBoolean  { 
    // use function body compiling 
return  MakeBoolean((self.Arg.Id() == C_set.Id()) && (ToBoolean(OBJ(F_CALL(C_finite_ask,ARGS(EID{self.T1.Id(),0})))) == CTRUE))
    } 
  
// The EID go function for: finite? @ subtype (throw: false) 
func E_finite_ask_subtype (self EID) EID { 
    return EID{/*(sm for finite? @ subtype= boolean)*/ F_finite_ask_subtype(ToSubtype(OBJ(self)) ).Id(),0}} 
  
// creates a subtype, with some normalization
// v3.2 list[t] -> subtype 
// v4.0 => no error
/* {1} OPT.The go function for: nth(self:class,x:type) [] */
func F_nth_class1 (self *ClaireClass ,x *ClaireType ) *ClaireType  { 
    // procedure body with s =  
var Result *ClaireType  
    if ((self.Id() == C_set.Id()) || 
        (self.Id() == C_list.Id())) /* If:2 */{ 
      /* Let:3 */{ 
        var _CL_obj *ClaireSubtype   = ToSubtype(new(ClaireSubtype).Is(C_subtype))
        /* noccur = 5 */
        _CL_obj.Arg = self
        _CL_obj.T1 = x
        Result = ToType(_CL_obj.Id())
        /* Let-3 */} 
      /* If!2 */}  else if (self.IsIn(C_type) != CTRUE) /* If:2 */{ 
      Result = ToType(CEMPTY.Id())
      } else {
      /* Let:3 */{ 
        var _CL_obj *ClaireSubtype   = ToSubtype(new(ClaireSubtype).Is(C_subtype))
        /* noccur = 5 */
        /* update:4 */{ 
          var va_arg1 *ClaireSubtype  
          var va_arg2 *ClaireClass  
          va_arg1 = _CL_obj
          if (self.Id() == C_subtype.Id()) /* If:5 */{ 
            va_arg2 = C_type
            } else {
            va_arg2 = self
            /* If-5 */} 
          /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
          va_arg1.Arg = va_arg2
          /* update-4 */} 
        _CL_obj.T1 = x
        Result = ToType(_CL_obj.Id())
        /* Let-3 */} 
      /* If-2 */} 
    return Result} 
  
// The EID go function for: nth @ list<type_expression>(class, type) (throw: false) 
func E_nth_class1 (self EID,x EID) EID { 
    return EID{/*(sm for nth @ list<type_expression>(class, type)= type)*/ F_nth_class1(ToClass(OBJ(self)),ToType(OBJ(x)) ).Id(),0}} 
  
// create a Param with a list of parameters (constant properties) l1 and a list
// of types l2
// v4.0 => no error
/* {1} OPT.The go function for: nth(self:class,l1:list,l2:list) [] */
func F_nth_class2 (self *ClaireClass ,l1 *ClaireList ,l2 *ClaireList ) *ClaireType  { 
    // procedure body with s =  
var Result *ClaireType  
    if (((self.Id() == C_list.Id()) || 
          (self.Id() == C_set.Id())) && 
        (l2.At(1-1).Isa.IsIn(C_subtype) == CTRUE)) /* If:2 */{ 
      Result = F_nth_class1(self,ToSubtype(l2.At(1-1)).T1)
      /* If!2 */}  else if (((self.Id() == C_list.Id()) || 
          (self.Id() == C_set.Id())) && 
        (l1.At(1-1) != C_of.Id())) /* If:2 */{ 
      Result = ToType(CEMPTY.Id())
      } else {
      /* Let:3 */{ 
        var _CL_obj *ClaireParam   = To_Param(new(ClaireParam).Is(C_Param))
        /* noccur = 7 */
        _CL_obj.Arg = self
        _CL_obj.Params = l1
        _CL_obj.Args = l2
        Result = ToType(_CL_obj.Id())
        /* Let-3 */} 
      /* If-2 */} 
    return Result} 
  
// The EID go function for: nth @ list<type_expression>(class, list, list) (throw: false) 
func E_nth_class2 (self EID,l1 EID,l2 EID) EID { 
    return EID{/*(sm for nth @ list<type_expression>(class, list, list)= type)*/ F_nth_class2(ToClass(OBJ(self)),ToList(OBJ(l1)),ToList(OBJ(l2)) ).Id(),0}} 
  
// create a Param of the stack[X] kind
/* {1} OPT.The go function for: param!(self:class,tx:type) [] */
func F_param_I_class (self *ClaireClass ,tx *ClaireType ) *ClaireType  { 
    // procedure body with s =  
var Result *ClaireType  
    /* Let:2 */{ 
      var _CL_obj *ClaireParam   = To_Param(new(ClaireParam).Is(C_Param))
      /* noccur = 7 */
      _CL_obj.Arg = self
      _CL_obj.Params = MakeConstantList(C_of.Id())
      _CL_obj.Args = MakeConstantList(MakeConstantSet(tx.Id()).Id())
      Result = ToType(_CL_obj.Id())
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: param! @ class (throw: false) 
func E_param_I_class (self EID,tx EID) EID { 
    return EID{/*(sm for param! @ class= type)*/ F_param_I_class(ToClass(OBJ(self)),ToType(OBJ(tx)) ).Id(),0}} 
  
// create the t[] param
/* {1} OPT.The go function for: nth(self:type) [] */
func F_nth_type (self *ClaireType ) *ClaireType  { 
    // procedure body with s =  
var Result *ClaireType  
    /* Let:2 */{ 
      var _CL_obj *ClaireParam   = To_Param(new(ClaireParam).Is(C_Param))
      /* noccur = 7 */
      _CL_obj.Arg = C_array
      _CL_obj.Params = MakeConstantList(C_of.Id())
      _CL_obj.Args = MakeConstantList(MakeConstantSet(self.Id()).Id())
      Result = ToType(_CL_obj.Id())
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: nth @ type (throw: false) 
func E_nth_type (self EID) EID { 
    return EID{/*(sm for nth @ type= type)*/ F_nth_type(ToType(OBJ(self)) ).Id(),0}} 
  
// tuple are types
/* {1} OPT.The go function for: finite?(self:tuple) [] */
func F_finite_ask_tuple (self *ClaireTuple ) *ClaireBoolean  { 
    // procedure body with s =  
var Result *ClaireBoolean  
    /* Let:2 */{ 
      var g0265UU *ClaireAny  
      /* noccur = 1 */
      /* For:3 */{ 
        var x *ClaireAny  
        _ = x
        g0265UU= CFALSE.Id()
        var x_support *ClaireList  
        x_support = ToList(self.Id())
        x_len := x_support.Length()
        for i_it := 0; i_it < x_len; i_it++ { 
          x = x_support.At(i_it)
          if (ToBoolean(OBJ(F_CALL(C_finite_ask,ARGS(x.ToEID())))) != CTRUE) /* If:5 */{ 
             /*v = g0265UU, s =any*/
g0265UU = CTRUE.Id()
            break
            /* If-5 */} 
          /* loop-4 */} 
        /* For-3 */} 
      Result = F_not_any(g0265UU)
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: finite? @ tuple (throw: false) 
func E_finite_ask_tuple (self EID) EID { 
    return EID{/*(sm for finite? @ tuple= boolean)*/ F_finite_ask_tuple(ToTuple(OBJ(self)) ).Id(),0}} 
  
// reference to a previous variable, not a type but a pattern -------
// index is the position of the stack of the referred type
// args is a list representing the path (a sequence of properties (parameters))
// a property is applied to the referred type
// if arg = true, the reference is the singleton containing the ref. value
// TODO check that arg is still used !
/* {1} OPT.The go function for: self_print(self:Reference) [] */
func F_self_print_Reference_Core (self *ClaireReference ) EID { 
    var Result EID 
    PRINC("<ref:")
    Result = F_print_any(self.Args.Id())
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("(ltype[")
    F_princ_integer(self.Index)
    PRINC("])>")
    Result = EVOID
    }
    return Result} 
  
// The EID go function for: self_print @ Reference (throw: true) 
func E_self_print_Reference_Core (self EID) EID { 
    return /*(sm for self_print @ Reference= EID)*/ F_self_print_Reference_Core(To_Reference(OBJ(self)) )} 
  
/* {1} OPT.The go function for: get(self:Reference,y:any) [] */
func F_get_Reference (self *ClaireReference ,y *ClaireAny ) *ClaireAny  { 
    // procedure body with s =  
var Result *ClaireAny  
    /* Let:2 */{ 
      var l *ClaireList   = self.Args
      /* noccur = 2 */
      /* Let:3 */{ 
        var i int  = 1
        /* noccur = 4 */
        /* Let:4 */{ 
          var g0266 int  = l.Length()
          /* noccur = 1 */
          for (i <= g0266) /* while:5 */{ 
            y = ANY(F_funcall_property(ToProperty(l.At(i-1)),y))
            i = (i+1)
            /* while-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      Result = y
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: get @ Reference (throw: false) 
func E_get_Reference (self EID,y EID) EID { 
    return /*(sm for get @ Reference= any)*/ F_get_Reference(To_Reference(OBJ(self)),ANY(y) ).ToEID()} 
  
// apply a reference to a type (l is args(self), passed for disambiguation)
/* {1} OPT.The go function for: @(self:Reference,l:list,y:any) [] */
func F__at_Reference (self *ClaireReference ,l *ClaireList ,y *ClaireAny ) *ClaireAny  { 
    // use function body compiling 

    /* Let:2 */{ 
      var i int  = 1
      /* noccur = 4 */
      /* Let:3 */{ 
        var g0267 int  = l.Length()
        /* noccur = 1 */
        for (i <= g0267) /* while:4 */{ 
          y = ToType(y).At(ToProperty(l.At(i-1))).Id()
          i = (i+1)
          /* while-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return  y
    } 
  
// The EID go function for: @ @ Reference (throw: false) 
func E__at_Reference (self EID,l EID,y EID) EID { 
    return /*(sm for @ @ Reference= any)*/ F__at_Reference(To_Reference(OBJ(self)),ToList(OBJ(l)),ANY(y) ).ToEID()} 
  
// type to set coercion  -------------------------------------------------
// new in v3.0.5 = use an interface method for type enumeration
// the default strategy is extensible: we look if there exists
// a proper definition that could be interpreted !
/* {1} OPT.The go function for: set!(x:collection) [] */
func F_set_I_collection (x *ClaireCollection ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var m *ClaireAny   = F__at_property1(C_set_I,x.Isa).Id()
      /* noccur = 2 */
      if (F_domain_I_restriction(ToRestriction(m)).Id() != C_collection.Id()) /* If:3 */{ 
        Result = F_CALL(C_funcall,ARGS(m.ToEID(),EID{x.Id(),0}))
        } else {
        Result = ToException(C_general_error.Make(MakeString("[178] cannot enumerate ~S").Id(),MakeConstantList(x.Id()).Id())).Close()
        /* If-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: set! @ collection (throw: true) 
func E_set_I_collection (x EID) EID { 
    return /*(sm for set! @ collection= EID)*/ F_set_I_collection(ToCollection(OBJ(x)) )} 
  
/* {1} OPT.The go function for: size(x:collection) [] */
func F_size_collection (x *ClaireCollection ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var m *ClaireAny   = F__at_property1(C_size,x.Isa).Id()
      /* noccur = 2 */
      if (F_domain_I_restriction(ToRestriction(m)).Id() != C_collection.Id()) /* If:3 */{ 
        Result = F_CALL(C_funcall,ARGS(m.ToEID(),EID{x.Id(),0}))
        } else {
        /* Let:4 */{ 
          var g0268UU *ClaireAny  
          /* noccur = 1 */
          var g0268UU_try02695 EID 
          g0268UU_try02695 = F_CALL(C_set_I,ARGS(EID{x.Id(),0}))
          /* ERROR PROTECTION INSERTED (g0268UU-Result) */
          if ErrorIn(g0268UU_try02695) {Result = g0268UU_try02695
          } else {
          g0268UU = ANY(g0268UU_try02695)
          Result = EID{C__INT,IVAL(ToSet(g0268UU).Size())}
          }
          /* Let-4 */} 
        /* If-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: size @ collection (throw: true) 
func E_size_collection (x EID) EID { 
    return /*(sm for size @ collection= EID)*/ F_size_collection(ToCollection(OBJ(x)) )} 
  
// v3.2.34  -> makes the API simpler
// (interface(size))
// set is needed for recursive def
/* {1} OPT.The go function for: set!(x:set) [] */
func F_set_I_set (x *ClaireSet ) *ClaireSet  { 
    // use function body compiling 
return  x
    } 
  
// The EID go function for: set! @ set (throw: false) 
func E_set_I_set (x EID) EID { 
    return EID{/*(sm for set! @ set= set)*/ F_set_I_set(ToSet(OBJ(x)) ).Id(),0}} 
  
// set is needed for recursive def
/* {1} OPT.The go function for: size(x:list) [] */
func F_size_list2_Core (x *ClaireList ) int { 
    // use function body compiling 
return  x.Set_I().Size()
    } 
  
// The EID go function for: size @ list (throw: false) 
func E_size_list2_Core (x EID) EID { 
    return EID{C__INT,IVAL(/*(sm for size @ list= integer)*/ F_size_list2_Core(ToList(OBJ(x)) ))}} 
  
// class  -> return a read-only list  (v3.2)
/* {1} OPT.The go function for: set!(x:class) [] */
func F_set_I_class (x *ClaireClass ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var rep *ClaireList   = ToType(CEMPTY.Id()).EmptyList()
      /* noccur = 3 */
      /* For:3 */{ 
        var c *ClaireAny  
        _ = c
        Result= EID{CFALSE.Id(),0}
        for _,c = range(x.Descendents.Values)/* loop:4 */{ 
          var void_try5 EID 
          _ = void_try5
          if ((ToClass(c).IsIn(C_primitive) == CTRUE) && 
              (c != C_boolean.Id())) /* If:5 */{ 
            void_try5 = ToException(C_general_error.Make(MakeString("[178] cannot enumerate ~S").Id(),MakeConstantList(c).Id())).Close()
            } else {
            rep = rep.Append(ToClass(c).Instances)
            void_try5 = EID{rep.Id(),0}
            /* If-5 */} 
          /* ERROR PROTECTION INSERTED (void_try5-Result) */
          if ErrorIn(void_try5) {Result = void_try5
          Result = void_try5
          break
          } else {
          }
          /* loop-4 */} 
        /* For-3 */} 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = EID{rep.Set_I().Id(),0}
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: set! @ class (throw: true) 
func E_set_I_class (x EID) EID { 
    return /*(sm for set! @ class= EID)*/ F_set_I_class(ToClass(OBJ(x)) )} 
  
/* {1} OPT.The go function for: size(self:class) [] */
func F_size_class (self *ClaireClass ) int { 
    // procedure body with s =  
var Result int 
    /* Let:2 */{ 
      var n int  = 0
      /* noccur = 3 */
      /* For:3 */{ 
        var x *ClaireAny  
        _ = x
        for _,x = range(self.Descendents.Values)/* loop:4 */{ 
          n = (n+ToClass(x).Instances.Length())
          /* loop-4 */} 
        /* For-3 */} 
      Result = n
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: size @ class (throw: false) 
func E_size_class (self EID) EID { 
    return EID{C__INT,IVAL(/*(sm for size @ class= integer)*/ F_size_class(ToClass(OBJ(self)) ))}} 
  
// Union
/* {1} OPT.The go function for: set!(x:Union) [] */
func F_set_I_Union (x *ClaireUnion ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var g0270UU *ClaireAny  
      /* noccur = 1 */
      var g0270UU_try02723 EID 
      g0270UU_try02723 = F_CALL(C_set_I,ARGS(EID{x.T1.Id(),0}))
      /* ERROR PROTECTION INSERTED (g0270UU-Result) */
      if ErrorIn(g0270UU_try02723) {Result = g0270UU_try02723
      } else {
      g0270UU = ANY(g0270UU_try02723)
      /* Let:3 */{ 
        var g0271UU *ClaireAny  
        /* noccur = 1 */
        var g0271UU_try02734 EID 
        g0271UU_try02734 = F_CALL(C_set_I,ARGS(EID{x.T2.Id(),0}))
        /* ERROR PROTECTION INSERTED (g0271UU-Result) */
        if ErrorIn(g0271UU_try02734) {Result = g0271UU_try02734
        } else {
        g0271UU = ANY(g0271UU_try02734)
        Result = EID{F_append_set(ToSet(g0270UU),ToSet(g0271UU)).Id(),0}
        }
        /* Let-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: set! @ Union (throw: true) 
func E_set_I_Union (x EID) EID { 
    return /*(sm for set! @ Union= EID)*/ F_set_I_Union(To_Union(OBJ(x)) )} 
  
/* {1} OPT.The go function for: size(x:Union) [] */
func F_size_Union (x *ClaireUnion ) EID { 
    var Result EID 
    if ((x.T1.Isa.IsIn(C_Interval) == CTRUE) || 
        (C_set.Id() == x.T1.Isa.Id())) /* If:2 */{ 
      /* Let:3 */{ 
        var g0274UU *ClaireAny  
        /* noccur = 1 */
        var g0274UU_try02764 EID 
        g0274UU_try02764 = F_CALL(C_size,ARGS(EID{x.T1.Id(),0}))
        /* ERROR PROTECTION INSERTED (g0274UU-Result) */
        if ErrorIn(g0274UU_try02764) {Result = g0274UU_try02764
        } else {
        g0274UU = ANY(g0274UU_try02764)
        /* Let:4 */{ 
          var g0275UU *ClaireAny  
          /* noccur = 1 */
          var g0275UU_try02775 EID 
          g0275UU_try02775 = F_CALL(C_size,ARGS(EID{x.T2.Id(),0}))
          /* ERROR PROTECTION INSERTED (g0275UU-Result) */
          if ErrorIn(g0275UU_try02775) {Result = g0275UU_try02775
          } else {
          g0275UU = ANY(g0275UU_try02775)
          Result = EID{C__INT,IVAL(F__plus_integer(ToInteger(g0274UU).Value,ToInteger(g0275UU).Value))}
          }
          /* Let-4 */} 
        }
        /* Let-3 */} 
      } else {
      /* Let:3 */{ 
        var g0278UU *ClaireSet  
        /* noccur = 1 */
        var g0278UU_try02794 EID 
        g0278UU_try02794 = F_set_I_Union(x)
        /* ERROR PROTECTION INSERTED (g0278UU-Result) */
        if ErrorIn(g0278UU_try02794) {Result = g0278UU_try02794
        } else {
        g0278UU = ToSet(OBJ(g0278UU_try02794))
        Result = F_CALL(C_length,ARGS(EID{g0278UU.Id(),0}))
        }
        /* Let-3 */} 
      /* If-2 */} 
    return RangeCheck(ToType(C_integer.Id()),Result)} 
  
// The EID go function for: size @ Union (throw: true) 
func E_size_Union (x EID) EID { 
    return /*(sm for size @ Union= EID)*/ F_size_Union(To_Union(OBJ(x)) )} 
  
// interval
/* {1} OPT.The go function for: set!(x:Interval) [] */
func F_set_I_Interval (x *ClaireInterval ) *ClaireSet  { 
    // use function body compiling 
return  ToSet(F_sequence_integer(x.Arg1,x.Arg2).Id())
    } 
  
// The EID go function for: set! @ Interval (throw: false) 
func E_set_I_Interval (x EID) EID { 
    return EID{/*(sm for set! @ Interval= set)*/ F_set_I_Interval(To_Interval(OBJ(x)) ).Id(),0}} 
  
/* {1} OPT.The go function for: size(self:Interval) [] */
func F_size_Interval (self *ClaireInterval ) int { 
    // use function body compiling 
return  ((self.Arg2+1)-self.Arg1)
    } 
  
// The EID go function for: size @ Interval (throw: false) 
func E_size_Interval (self EID) EID { 
    return EID{C__INT,IVAL(/*(sm for size @ Interval= integer)*/ F_size_Interval(To_Interval(OBJ(self)) ))}} 
  
// param
/* {1} OPT.The go function for: set!(x:Param) [] */
func F_set_I_Param (x *ClaireParam ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var y_in *ClaireSet  
      /* noccur = 2 */
      var y_in_try02803 EID 
      y_in_try02803 = F_set_I_class(x.Arg)
      /* ERROR PROTECTION INSERTED (y_in-Result) */
      if ErrorIn(y_in_try02803) {Result = y_in_try02803
      } else {
      y_in = ToSet(OBJ(y_in_try02803))
      /* Let:3 */{ 
        var y_out *ClaireSet   = y_in.Empty()
        /* noccur = 2 */
        /* For:4 */{ 
          var y *ClaireAny  
          _ = y
          for _,y = range(y_in.Values)/* loop:5 */{ 
            if (ToType(x.Id()).Contains(y) == CTRUE) /* If:6 */{ 
              y_out.AddFast(y)
              /* If-6 */} 
            /* loop-5 */} 
          /* For-4 */} 
        Result = EID{y_out.Id(),0}
        /* Let-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: set! @ Param (throw: true) 
func E_set_I_Param (x EID) EID { 
    return /*(sm for set! @ Param= EID)*/ F_set_I_Param(To_Param(OBJ(x)) )} 
  
/* {1} OPT.The go function for: size(x:Param) [] */
func F_size_Param (x *ClaireParam ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var g0281UU *ClaireSet  
      /* noccur = 1 */
      var g0281UU_try02823 EID 
      g0281UU_try02823 = F_set_I_Param(x)
      /* ERROR PROTECTION INSERTED (g0281UU-Result) */
      if ErrorIn(g0281UU_try02823) {Result = g0281UU_try02823
      } else {
      g0281UU = ToSet(OBJ(g0281UU_try02823))
      Result = F_CALL(C_length,ARGS(EID{g0281UU.Id(),0}))
      }
      /* Let-2 */} 
    return RangeCheck(ToType(C_integer.Id()),Result)} 
  
// The EID go function for: size @ Param (throw: true) 
func E_size_Param (x EID) EID { 
    return /*(sm for size @ Param= EID)*/ F_size_Param(To_Param(OBJ(x)) )} 
  
// subtype
/* {1} OPT.The go function for: set!(x:subtype) [] */
func F_set_I_subtype (x *ClaireSubtype ) EID { 
    var Result EID 
    if (x.Arg.Id() == C_set.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var g0283UU *ClaireList  
        /* noccur = 1 */
        var g0283UU_try02844 EID 
        /* Let:4 */{ 
          var g0285UU *ClaireAny  
          /* noccur = 1 */
          var g0285UU_try02865 EID 
          g0285UU_try02865 = F_CALL(C_set_I,ARGS(EID{x.T1.Id(),0}))
          /* ERROR PROTECTION INSERTED (g0285UU-g0283UU_try02844) */
          if ErrorIn(g0285UU_try02865) {g0283UU_try02844 = g0285UU_try02865
          } else {
          g0285UU = ANY(g0285UU_try02865)
          g0283UU_try02844 = EID{ToSet(g0285UU).List_I().Id(),0}
          }
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (g0283UU-Result) */
        if ErrorIn(g0283UU_try02844) {Result = g0283UU_try02844
        } else {
        g0283UU = ToList(OBJ(g0283UU_try02844))
        Result = EID{F_build_powerset_list(g0283UU).Id(),0}
        }
        /* Let-3 */} 
      } else {
      Result = ToException(C_general_error.Make(MakeString("[178] cannot enumerate ~S").Id(),MakeConstantList(x.Id()).Id())).Close()
      /* If-2 */} 
    return Result} 
  
// The EID go function for: set! @ subtype (throw: true) 
func E_set_I_subtype (x EID) EID { 
    return /*(sm for set! @ subtype= EID)*/ F_set_I_subtype(ToSubtype(OBJ(x)) )} 
  
/* {1} OPT.The go function for: size(x:subtype) [] */
func F_size_subtype (x *ClaireSubtype ) EID { 
    var Result EID 
    if (x.Arg.Id() == C_set.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var g0287UU *ClaireAny  
        /* noccur = 1 */
        var g0287UU_try02884 EID 
        g0287UU_try02884 = F_CALL(C_size,ARGS(EID{x.T1.Id(),0}))
        /* ERROR PROTECTION INSERTED (g0287UU-Result) */
        if ErrorIn(g0287UU_try02884) {Result = g0287UU_try02884
        } else {
        g0287UU = ANY(g0287UU_try02884)
        Result = F__exp2_integer(ToInteger(g0287UU).Value)
        }
        /* Let-3 */} 
      } else {
      Result = ToException(C_general_error.Make(MakeString("[178] cannot enumerate ~S").Id(),MakeConstantList(x.Id()).Id())).Close()
      /* If-2 */} 
    return Result} 
  
// The EID go function for: size @ subtype (throw: true) 
func E_size_subtype (x EID) EID { 
    return /*(sm for size @ subtype= EID)*/ F_size_subtype(ToSubtype(OBJ(x)) )} 
  
// tuple
/* {1} OPT.The go function for: set!(x:tuple) [] */
func F_set_I_tuple (x *ClaireTuple ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var l *ClaireList   = ToList(x.Id())
      /* noccur = 4 */
      if (F_boolean_I_any(l.Id()).Id() != CTRUE.Id()) /* If:3 */{ 
        Result = EID{MakeConstantSet(CEMPTY.Id()).Id(),0}
        } else {
        /* Let:4 */{ 
          var l1 *ClaireSet  
          /* noccur = 3 */
          var l1_try02905 EID 
          /* Let:5 */{ 
            var y_bag *ClaireSet   = ToType(CEMPTY.Id()).EmptySet()
            /* noccur = 2 */
            /* For:6 */{ 
              var y *ClaireAny  
              _ = y
              l1_try02905= EID{CFALSE.Id(),0}
              var y_support *ClaireSet  
              var y_support_try02917 EID 
              y_support_try02917 = F_CALL(C_set_I,ARGS(l.At(1-1).ToEID()))
              /* ERROR PROTECTION INSERTED (y_support-l1_try02905) */
              if ErrorIn(y_support_try02917) {l1_try02905 = y_support_try02917
              } else {
              y_support = ToSet(OBJ(y_support_try02917))
              for _,y = range(y_support.Values)/* loop2:7 */{ 
                y_bag.AddFast(MakeConstantList(y).Id())
                }
                /* loop-7 */} 
              /* For-6 */} 
            /* ERROR PROTECTION INSERTED (l1_try02905-l1_try02905) */
            if !ErrorIn(l1_try02905) {
            l1_try02905 = EID{y_bag.Id(),0}
            }
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (l1-Result) */
          if ErrorIn(l1_try02905) {Result = l1_try02905
          } else {
          l1 = ToSet(OBJ(l1_try02905))
          /* Let:5 */{ 
            var n int  = 2
            /* noccur = 4 */
            /* Let:6 */{ 
              var g0289 int  = l.Length()
              /* noccur = 1 */
              Result= EID{CFALSE.Id(),0}
              for (n <= g0289) /* while:7 */{ 
                var void_try8 EID 
                _ = void_try8
                { 
                /* Let:8 */{ 
                  var l2 *ClaireSet   = ToType(C_any.Id()).EmptySet()
                  /* noccur = 3 */
                  /* For:9 */{ 
                    var z *ClaireAny  
                    _ = z
                    void_try8= EID{CFALSE.Id(),0}
                    var z_support *ClaireSet  
                    var z_support_try029210 EID 
                    z_support_try029210 = F_CALL(C_set_I,ARGS(l.At(n-1).ToEID()))
                    /* ERROR PROTECTION INSERTED (z_support-void_try8) */
                    if ErrorIn(z_support_try029210) {void_try8 = z_support_try029210
                    } else {
                    z_support = ToSet(OBJ(z_support_try029210))
                    for _,z = range(z_support.Values)/* loop2:10 */{ 
                      /* For:11 */{ 
                        var l3 *ClaireAny  
                        _ = l3
                        for _,l3 = range(l1.Values)/* loop:12 */{ 
                          l2 = l2.AddFast(ToList(l3).Copy().AddFast(z).Id())
                          /* loop-12 */} 
                        /* For-11 */} 
                      }
                      /* loop-10 */} 
                    /* For-9 */} 
                  /* ERROR PROTECTION INSERTED (void_try8-void_try8) */
                  if ErrorIn(void_try8) {Result = void_try8
                  break
                  } else {
                  l1 = l2
                  void_try8 = EID{l1.Id(),0}
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (void_try8-void_try8) */
                if ErrorIn(void_try8) {Result = void_try8
                break
                } else {
                n = (n+1)
                }
                /* while-7 */} 
              }
              /* Let-6 */} 
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          Result = EID{l1.Id(),0}
          }
          }
          /* Let-4 */} 
        /* If-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: set! @ tuple (throw: true) 
func E_set_I_tuple (x EID) EID { 
    return /*(sm for set! @ tuple= EID)*/ F_set_I_tuple(ToTuple(OBJ(x)) )} 
  
/* {1} OPT.The go function for: size(l:tuple) [] */
func F_size_tuple (l *ClaireTuple ) EID { 
    var Result EID 
    if (F_boolean_I_any(l.Id()).Id() != CTRUE.Id()) /* If:2 */{ 
      Result = EID{C__INT,IVAL(1)}
      } else {
      /* Let:3 */{ 
        var m int 
        /* noccur = 3 */
        var m_try02944 EID 
        m_try02944 = F_CALL(C_size,ARGS(ToList(l.Id()).At(1-1).ToEID()))
        /* ERROR PROTECTION INSERTED (m-Result) */
        if ErrorIn(m_try02944) {Result = m_try02944
        } else {
        m = INT(m_try02944)
        /* Let:4 */{ 
          var n int  = 2
          /* noccur = 4 */
          /* Let:5 */{ 
            var g0293 int  = l.Length()
            /* noccur = 1 */
            Result= EID{CFALSE.Id(),0}
            for (n <= g0293) /* while:6 */{ 
              var void_try7 EID 
              _ = void_try7
              { 
              var m_try02957 EID 
              /* Let:7 */{ 
                var g0296UU *ClaireAny  
                /* noccur = 1 */
                var g0296UU_try02978 EID 
                g0296UU_try02978 = F_CALL(C_size,ARGS(ToList(l.Id()).At(n-1).ToEID()))
                /* ERROR PROTECTION INSERTED (g0296UU-m_try02957) */
                if ErrorIn(g0296UU_try02978) {m_try02957 = g0296UU_try02978
                } else {
                g0296UU = ANY(g0296UU_try02978)
                m_try02957 = EID{C__INT,IVAL((m*ToInteger(g0296UU).Value))}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (m-void_try7) */
              if ErrorIn(m_try02957) {void_try7 = m_try02957
              Result = m_try02957
              break
              } else {
              m = INT(m_try02957)
              void_try7 = EID{C__INT,IVAL(m)}
              n = (n+1)
              }
              /* while-6 */} 
            }
            /* Let-5 */} 
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        Result = EID{C__INT,IVAL(m)}
        }
        }
        /* Let-3 */} 
      /* If-2 */} 
    return Result} 
  
// The EID go function for: size @ tuple (throw: true) 
func E_size_tuple (l EID) EID { 
    return /*(sm for size @ tuple= EID)*/ F_size_tuple(ToTuple(OBJ(l)) )} 
  
// declarations (now useless in CLAIRE4)
// ********************************************************************
// *                Part 3: Interface Methods                         *
// ********************************************************************
// there is a special restriction for + to specify the way the inheritance
// conflict should be solved
//U(self:set,ens:type) : type -> (case ens (set self /+ ens, any ens U self))
// the union makes a partial reduction to the normal form. The complete
// reduction is done by enumeration if needed during the type subsumption
// union is left-associative: A U B U C is represented by (A U B) U C  => never(t2(x:Union) % union)
// a union of intervals is ALWAYS disjoint
/* {1} OPT.The go function for: U(x:type,y:type) [] */
func F_U_type (x *ClaireType ,y *ClaireType ) *ClaireType  { 
    // procedure body with s =  
var Result *ClaireType  
    if (C_set.Id() == x.Isa.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var g0298 *ClaireSet   = ToSet(x.Id())
        /* noccur = 2 */
        if (C_set.Id() == y.Isa.Id()) /* If:4 */{ 
          /* Let:5 */{ 
            var g0299 *ClaireSet   = ToSet(y.Id())
            /* noccur = 1 */
            Result = ToType(F_append_set(g0298,g0299).Id())
            /* Let-5 */} 
          } else {
          Result = F_U_type(y,ToType(g0298.Id()))
          /* If-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (y.Included(x) == CTRUE) /* If:2 */{ 
      Result = x
      /* If!2 */}  else if (x.Included(y) == CTRUE) /* If:2 */{ 
      Result = y
      /* If!2 */}  else if (y.Isa.IsIn(C_Union) == CTRUE) /* If:2 */{ 
      Result = F_U_type(F_U_type(x,ToType(OBJ(F_CALL(C_mClaire_t1,ARGS(EID{y.Id(),0}))))),To_Union(y.Id()).T2)
      } else {
      var g0307I *ClaireBoolean  
      if (x.Isa.IsIn(C_Interval) == CTRUE) /* If:3 */{ 
        g0307I = y.Isa.IsIn(C_Interval)
        } else {
        g0307I = CFALSE
        /* If-3 */} 
      if (g0307I == CTRUE) /* If:3 */{ 
        if (((To_Interval(y.Id()).Arg1-1) <= To_Interval(x.Id()).Arg2) && 
            (To_Interval(x.Id()).Arg1 <= To_Interval(y.Id()).Arg1)) /* If:4 */{ 
          Result = F__dot_dot_integer(To_Interval(x.Id()).Arg1,To_Interval(y.Id()).Arg2)
          /* If!4 */}  else if (((To_Interval(x.Id()).Arg1-1) <= To_Interval(y.Id()).Arg2) && 
            (To_Interval(y.Id()).Arg1 <= To_Interval(x.Id()).Arg1)) /* If:4 */{ 
          Result = F__dot_dot_integer(To_Interval(y.Id()).Arg1,To_Interval(x.Id()).Arg2)
          } else {
          /* Let:5 */{ 
            var _CL_obj *ClaireUnion   = To_Union(new(ClaireUnion).Is(C_Union))
            /* noccur = 5 */
            _CL_obj.T1 = x
            _CL_obj.T2 = y
            Result = ToType(_CL_obj.Id())
            /* Let-5 */} 
          /* If-4 */} 
        } else {
        var g0308I *ClaireBoolean  
        if (x.Isa.IsIn(C_Union) == CTRUE) /* If:4 */{ 
          g0308I = y.Isa.IsIn(C_Interval)
          } else {
          g0308I = CFALSE
          /* If-4 */} 
        if (g0308I == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var z *ClaireType   = F_U_type(To_Union(x.Id()).T2,y)
            /* noccur = 2 */
            if (z.Isa.IsIn(C_Union) == CTRUE) /* If:6 */{ 
              /* Let:7 */{ 
                var _CL_obj *ClaireUnion   = To_Union(new(ClaireUnion).Is(C_Union))
                /* noccur = 5 */
                _CL_obj.T1 = F_U_type(ToType(OBJ(F_CALL(C_mClaire_t1,ARGS(EID{x.Id(),0})))),y)
                _CL_obj.T2 = To_Union(x.Id()).T2
                Result = ToType(_CL_obj.Id())
                /* Let-7 */} 
              } else {
              Result = F_U_type(ToType(OBJ(F_CALL(C_mClaire_t1,ARGS(EID{x.Id(),0})))),z)
              /* If-6 */} 
            /* Let-5 */} 
          } else {
          var g0309I *ClaireBoolean  
          if (x.Isa.IsIn(C_Interval) == CTRUE) /* If:5 */{ 
            /* Let:6 */{ 
              var g0306 *ClaireInterval   = To_Interval(x.Id())
              /* noccur = 2 */
              g0309I = MakeBoolean((C_set.Id() == y.Isa.Id()) && ((y.Contains(MakeInteger((g0306.Arg1-1)).Id()) == CTRUE) || 
                  (y.Contains(MakeInteger((g0306.Arg2+1)).Id()) == CTRUE)))
              /* Let-6 */} 
            } else {
            g0309I = CFALSE
            /* If-5 */} 
          if (g0309I == CTRUE) /* If:5 */{ 
            /* Let:6 */{ 
              var a int  = To_Interval(x.Id()).Arg1
              /* noccur = 4 */
              /* Let:7 */{ 
                var b int  = To_Interval(x.Id()).Arg2
                /* noccur = 4 */
                if (y.Contains(MakeInteger((a-1)).Id()) == CTRUE) /* If:8 */{ 
                  a = (a-1)
                  /* If-8 */} 
                if (y.Contains(MakeInteger((b+1)).Id()) == CTRUE) /* If:8 */{ 
                  b = (b+1)
                  /* If-8 */} 
                Result = F_U_type(F__dot_dot_integer(a,b),y)
                /* Let-7 */} 
              /* Let-6 */} 
            } else {
            if (C_set.Id() == y.Isa.Id()) /* If:6 */{ 
              /* Let:7 */{ 
                var z_in *ClaireSet   = ToSet(y.Id())
                /* noccur = 2 */
                /* Let:8 */{ 
                  var z_out *ClaireSet   = z_in.Empty()
                  /* noccur = 2 */
                  /* For:9 */{ 
                    var z *ClaireAny  
                    _ = z
                    for _,z = range(z_in.Values)/* loop:10 */{ 
                      if (x.Contains(z) != CTRUE) /* If:11 */{ 
                        z_out.AddFast(z)
                        /* If-11 */} 
                      /* loop-10 */} 
                    /* For-9 */} 
                  y = ToType(z_out.Id())
                  /* Let-8 */} 
                /* Let-7 */} 
              /* If-6 */} 
            /* Let:6 */{ 
              var _CL_obj *ClaireUnion   = To_Union(new(ClaireUnion).Is(C_Union))
              /* noccur = 5 */
              _CL_obj.T1 = x
              _CL_obj.T2 = y
              Result = ToType(_CL_obj.Id())
              /* Let-6 */} 
            /* If-5 */} 
          /* If-4 */} 
        /* If-3 */} 
      /* If-2 */} 
    return Result} 
  
// The EID go function for: U @ type (throw: false) 
func E_U_type (x EID,y EID) EID { 
    return EID{/*(sm for U @ type= type)*/ F_U_type(ToType(OBJ(x)),ToType(OBJ(y)) ).Id(),0}} 
  
// the Interval construction method has a smart second-order type  - fix on v3.1.06
/* {1} OPT.The go function for: ..(x:integer,y:integer) [] */
func F__dot_dot_integer (x int,y int) *ClaireType  { 
    // procedure body with s =  
var Result *ClaireType  
    if (x <= y) /* If:2 */{ 
      Result = ToType(C_Interval.MakeInts(x,y))
      } else {
      Result = ToType(CEMPTY.Id())
      /* If-2 */} 
    return Result} 
  
// The EID go function for: .. @ list<type_expression>(integer, integer) (throw: false) 
func E__dot_dot_integer (x EID,y EID) EID { 
    return EID{/*(sm for .. @ list<type_expression>(integer, integer)= type)*/ F__dot_dot_integer(INT(x),INT(y) ).Id(),0}} 
  
/* {1} OPT.The go function for: _dot_dot_integer_type */
func F__dot_dot_integer_type (x *ClaireType ,y *ClaireType ) EID { 
    /* eid body: (if (unique?(x) & unique?(y) & the(x) <= the(y)) set(the(x) .. the(y)) else subtype[integer]) */
    var Result EID 
    var g0310I *ClaireBoolean  
    var g0310I_try03112 EID 
    /* and:2 */{ 
      var v_and2 *ClaireBoolean  
      
      v_and2 = ToBoolean(OBJ(F_CALL(C_unique_ask,ARGS(EID{x.Id(),0}))))
      if (v_and2 == CFALSE) {g0310I_try03112 = EID{CFALSE.Id(),0}
      } else /* arg:3 */{ 
        v_and2 = ToBoolean(OBJ(F_CALL(C_unique_ask,ARGS(EID{y.Id(),0}))))
        if (v_and2 == CFALSE) {g0310I_try03112 = EID{CFALSE.Id(),0}
        } else /* arg:4 */{ 
          var v_and2_try03125 EID 
          /* Let:5 */{ 
            var g0313UU *ClaireAny  
            /* noccur = 1 */
            var g0313UU_try03156 EID 
            g0313UU_try03156 = F_CALL(C_the,ARGS(EID{x.Id(),0}))
            /* ERROR PROTECTION INSERTED (g0313UU-v_and2_try03125) */
            if ErrorIn(g0313UU_try03156) {v_and2_try03125 = g0313UU_try03156
            } else {
            g0313UU = ANY(g0313UU_try03156)
            /* Let:6 */{ 
              var g0314UU *ClaireAny  
              /* noccur = 1 */
              var g0314UU_try03167 EID 
              g0314UU_try03167 = F_CALL(C_the,ARGS(EID{y.Id(),0}))
              /* ERROR PROTECTION INSERTED (g0314UU-v_and2_try03125) */
              if ErrorIn(g0314UU_try03167) {v_and2_try03125 = g0314UU_try03167
              } else {
              g0314UU = ANY(g0314UU_try03167)
              v_and2_try03125 = F_CALL(ToProperty(C__inf_equal.Id()),ARGS(g0313UU.ToEID(),g0314UU.ToEID()))
              }
              /* Let-6 */} 
            }
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (v_and2-g0310I_try03112) */
          if ErrorIn(v_and2_try03125) {g0310I_try03112 = v_and2_try03125
          } else {
          v_and2 = ToBoolean(OBJ(v_and2_try03125))
          if (v_and2 == CFALSE) {g0310I_try03112 = EID{CFALSE.Id(),0}
          } else /* arg:5 */{ 
            g0310I_try03112 = EID{CTRUE.Id(),0}/* arg-5 */} 
          /* arg-4 */} 
        /* arg-3 */} 
      }
      /* and-2 */} 
    /* ERROR PROTECTION INSERTED (g0310I-Result) */
    if ErrorIn(g0310I_try03112) {Result = g0310I_try03112
    } else {
    g0310I = ToBoolean(OBJ(g0310I_try03112))
    if (g0310I == CTRUE) /* If:2 */{ 
      /* Construct:3 */{ 
        var v_bag_arg *ClaireAny  
        Result= EID{ToType(CEMPTY.Id()).EmptySet().Id(),0}
        var v_bag_arg_try03174 EID 
        /* Let:4 */{ 
          var g0318UU *ClaireAny  
          /* noccur = 1 */
          var g0318UU_try03205 EID 
          g0318UU_try03205 = F_CALL(C_the,ARGS(EID{x.Id(),0}))
          /* ERROR PROTECTION INSERTED (g0318UU-v_bag_arg_try03174) */
          if ErrorIn(g0318UU_try03205) {v_bag_arg_try03174 = g0318UU_try03205
          } else {
          g0318UU = ANY(g0318UU_try03205)
          /* Let:5 */{ 
            var g0319UU *ClaireAny  
            /* noccur = 1 */
            var g0319UU_try03216 EID 
            g0319UU_try03216 = F_CALL(C_the,ARGS(EID{y.Id(),0}))
            /* ERROR PROTECTION INSERTED (g0319UU-v_bag_arg_try03174) */
            if ErrorIn(g0319UU_try03216) {v_bag_arg_try03174 = g0319UU_try03216
            } else {
            g0319UU = ANY(g0319UU_try03216)
            v_bag_arg_try03174 = F_CALL(ToProperty(C__dot_dot.Id()),ARGS(g0318UU.ToEID(),g0319UU.ToEID()))
            }
            /* Let-5 */} 
          }
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (v_bag_arg-Result) */
        if ErrorIn(v_bag_arg_try03174) {Result = v_bag_arg_try03174
        } else {
        v_bag_arg = ANY(v_bag_arg_try03174)
        ToSet(OBJ(Result)).AddFast(v_bag_arg)}
        /* Construct-3 */} 
      } else {
      Result = F_CALL(C_nth,ARGS(EID{C_subtype.Id(),0},EID{C_integer.Id(),0}))
      /* If-2 */} 
    }
    return Result} 
  
  
// The dual EID go function for: "_dot_dot_integer_type" 
func E__dot_dot_integer_type (x EID,y EID) EID { 
    return F__dot_dot_integer_type(ToType(OBJ(x)),ToType(OBJ(y)))} 
  
// exception
/* {1} OPT.The go function for: but(s:any,x:any) [] */
func F_but_any (s *ClaireAny ,x *ClaireAny ) EID { 
    var Result EID 
    if (s.Isa.IsIn(C_list) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0322 *ClaireList   = ToList(s)
        /* noccur = 1 */
        /* Let:4 */{ 
          var y_in *ClaireList   = g0322
          /* noccur = 2 */
          /* Let:5 */{ 
            var y_out *ClaireList   = y_in.Empty()
            /* noccur = 2 */
            /* For:6 */{ 
              var y *ClaireAny  
              _ = y
              var y_support *ClaireList  
              y_support = y_in
              y_len := y_support.Length()
              for i_it := 0; i_it < y_len; i_it++ { 
                y = y_support.At(i_it)
                if (Equal(y,x) != CTRUE) /* If:8 */{ 
                  y_out.AddFast(y)
                  /* If-8 */} 
                /* loop-7 */} 
              /* For-6 */} 
            Result = EID{y_out.Id(),0}
            /* Let-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (C_set.Id() == s.Isa.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var g0323 *ClaireSet   = ToSet(s)
        /* noccur = 1 */
        Result = EID{g0323.Copy().Delete(x).Id(),0}
        /* Let-3 */} 
      } else {
      /* Let:3 */{ 
        var g0325UU *ClaireList  
        /* noccur = 1 */
        var g0325UU_try03264 EID 
        g0325UU_try03264 = F_enumerate_any(s)
        /* ERROR PROTECTION INSERTED (g0325UU-Result) */
        if ErrorIn(g0325UU_try03264) {Result = g0325UU_try03264
        } else {
        g0325UU = ToList(OBJ(g0325UU_try03264))
        Result = EID{g0325UU.Delete(x).Id(),0}
        }
        /* Let-3 */} 
      /* If-2 */} 
    return Result} 
  
// The EID go function for: but @ any (throw: true) 
func E_but_any (s EID,x EID) EID { 
    return /*(sm for but @ any= EID)*/ F_but_any(ANY(s),ANY(x) )} 
  
/* {1} OPT.The go function for: but_any_type */
func F_but_any_type (s *ClaireType ,x *ClaireType ) EID { 
    /* eid body: (if (x <= list) list[member(s)] else if (x <= set) set[member(s)] else any) */
    var Result EID 
    if (ToBoolean(OBJ(F_CALL(ToProperty(C__inf_equal.Id()),ARGS(EID{x.Id(),0},EID{C_list.Id(),0})))) == CTRUE) /* If:2 */{ 
      Result = F_CALL(C_nth,ARGS(EID{C_list.Id(),0},F_CALL(C_member,ARGS(EID{s.Id(),0}))))
      /* If!2 */}  else if (ToBoolean(OBJ(F_CALL(ToProperty(C__inf_equal.Id()),ARGS(EID{x.Id(),0},EID{C_set.Id(),0})))) == CTRUE) /* If:2 */{ 
      Result = F_CALL(C_nth,ARGS(EID{C_set.Id(),0},F_CALL(C_member,ARGS(EID{s.Id(),0}))))
      } else {
      Result = EID{C_any.Id(),0}
      /* If-2 */} 
    return Result} 
  
  
// The dual EID go function for: "but_any_type" 
func E_but_any_type (s EID,x EID) EID { 
    return F_but_any_type(ToType(OBJ(s)),ToType(OBJ(x)))} 
  
// a set difference (extended to types, with implicit enumeration)
/* {1} OPT.The go function for: \(x:type,y:type) [] */
func F__backslash_type (x *ClaireType ,y *ClaireType ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var z_out *ClaireSet   = ToType(CEMPTY.Id()).EmptySet()
      /* noccur = 2 */
      /* For:3 */{ 
        var z *ClaireAny  
        _ = z
        Result= EID{CFALSE.Id(),0}
        var z_support *ClaireList  
        var z_support_try03274 EID 
        z_support_try03274 = F_enumerate_any(x.Id())
        /* ERROR PROTECTION INSERTED (z_support-Result) */
        if ErrorIn(z_support_try03274) {Result = z_support_try03274
        } else {
        z_support = ToList(OBJ(z_support_try03274))
        z_len := z_support.Length()
        for i_it := 0; i_it < z_len; i_it++ { 
          z = z_support.At(i_it)
          if (y.Contains(z) != CTRUE) /* If:5 */{ 
            z_out.AddFast(z)
            /* If-5 */} 
          }
          /* loop-4 */} 
        /* For-3 */} 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = EID{z_out.Id(),0}
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: \ @ type (throw: true) 
func E__backslash_type (x EID,y EID) EID { 
    return /*(sm for \ @ type= EID)*/ F__backslash_type(ToType(OBJ(x)),ToType(OBJ(y)) )} 
  
// ******************************************************************
// *    Part 4: Lattice methods                                     *
// ******************************************************************
// glb operation ---------------------------------------------------
// gbl is the extension of the lattice operator ^ for types to type_expressions
// new in v3.0.60: we reintroduce a glb method
/* {1} OPT.The go function for: glb(x:set,y:type) [] */
func F_glb_set (x *ClaireSet ,y *ClaireType ) *ClaireSet  { 
    // procedure body with s =  
var Result *ClaireSet  
    /* Let:2 */{ 
      var z_in *ClaireSet   = x
      /* noccur = 2 */
      /* Let:3 */{ 
        var z_out *ClaireSet   = z_in.Empty()
        /* noccur = 2 */
        /* For:4 */{ 
          var z *ClaireAny  
          _ = z
          for _,z = range(z_in.Values)/* loop:5 */{ 
            if (y.Contains(z) == CTRUE) /* If:6 */{ 
              z_out.AddFast(z)
              /* If-6 */} 
            /* loop-5 */} 
          /* For-4 */} 
        Result = z_out
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: glb @ set (throw: false) 
func E_glb_set (x EID,y EID) EID { 
    return EID{/*(sm for glb @ set= set)*/ F_glb_set(ToSet(OBJ(x)),ToType(OBJ(y)) ).Id(),0}} 
  
/* {1} OPT.The go function for: glb(x:Union,y:type) [] */
func F_glb_Union (x *ClaireUnion ,y *ClaireType ) *ClaireType  { 
    // use function body compiling 
return  F_U_type(ToType(OBJ(F_CALL(ToProperty(C_glb.Id()),ARGS(EID{x.T1.Id(),0},EID{y.Id(),0})))),ToType(OBJ(F_CALL(ToProperty(C_glb.Id()),ARGS(EID{x.T2.Id(),0},EID{y.Id(),0})))))
    } 
  
// The EID go function for: glb @ Union (throw: false) 
func E_glb_Union (x EID,y EID) EID { 
    return EID{/*(sm for glb @ Union= type)*/ F_glb_Union(To_Union(OBJ(x)),ToType(OBJ(y)) ).Id(),0}} 
  
/* {1} OPT.The go function for: glb(x:Interval,y:type) [] */
func F_glb_Interval (x *ClaireInterval ,y *ClaireType ) *ClaireType  { 
    // procedure body with s =  
var Result *ClaireType  
    if (C_class.Id() == y.Isa.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var g0328 *ClaireClass   = ToClass(y.Id())
        /* noccur = 1 */
        Result = ToType(IfThenElse((ToType(C_integer.Id()).Included(ToType(g0328.Id())) == CTRUE),
          x.Id(),
          CEMPTY.Id()))
        /* Let-3 */} 
      /* If!2 */}  else if (C_set.Id() == y.Isa.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var g0329 *ClaireSet   = ToSet(y.Id())
        /* noccur = 1 */
        Result = ToType(F_glb_set(g0329,ToType(x.Id())).Id())
        /* Let-3 */} 
      /* If!2 */}  else if (y.Isa.IsIn(C_Interval) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0330 *ClaireInterval   = To_Interval(y.Id())
        /* noccur = 6 */
        if (x.Arg1 <= g0330.Arg1) /* If:4 */{ 
          if (g0330.Arg1 <= x.Arg2) /* If:5 */{ 
            if (x.Arg2 <= g0330.Arg2) /* If:6 */{ 
              Result = F__dot_dot_integer(g0330.Arg1,x.Arg2)
              } else {
              Result = ToType(g0330.Id())
              /* If-6 */} 
            } else {
            Result = ToType(CEMPTY.Id())
            /* If-5 */} 
          } else {
          Result = F_glb_Interval(g0330,ToType(x.Id()))
          /* If-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (y.Isa.IsIn(C_Union) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0331 *ClaireUnion   = To_Union(y.Id())
        /* noccur = 2 */
        Result = F_U_type(F_glb_Interval(x,g0331.T1),F_glb_Interval(x,g0331.T2))
        /* Let-3 */} 
      } else {
      Result = ToType(CEMPTY.Id())
      /* If-2 */} 
    return Result} 
  
// The EID go function for: glb @ Interval (throw: false) 
func E_glb_Interval (x EID,y EID) EID { 
    return EID{/*(sm for glb @ Interval= type)*/ F_glb_Interval(To_Interval(OBJ(x)),ToType(OBJ(y)) ).Id(),0}} 
  
/* {1} OPT.The go function for: glb(x:class,y:type) [] */
func F_glb_class (x *ClaireClass ,y *ClaireType ) *ClaireType  { 
    // procedure body with s =  
var Result *ClaireType  
    if ((x.Open == ClEnv.ABSTRACT) && 
        (F_boolean_I_any(x.Subclass.Id()).Id() != CTRUE.Id())) /* If:2 */{ 
      /* Let:3 */{ 
        var z_out *ClaireSet   = ToType(CEMPTY.Id()).EmptySet()
        /* noccur = 2 */
        /* For:4 */{ 
          var z *ClaireAny  
          _ = z
          var z_support *ClaireList  
          z_support = x.Instances
          z_len := z_support.Length()
          for i_it := 0; i_it < z_len; i_it++ { 
            z = z_support.At(i_it)
            if (y.Contains(z) == CTRUE) /* If:6 */{ 
              z_out.AddFast(z)
              /* If-6 */} 
            /* loop-5 */} 
          /* For-4 */} 
        Result = ToType(z_out.Id())
        /* Let-3 */} 
      /* If!2 */}  else if ((x.Open == ClEnv.ABSTRACT) && 
        (F_boolean_I_any(x.Instances.Id()).Id() != CTRUE.Id())) /* If:2 */{ 
      /* Let:3 */{ 
        var g0335UU *ClaireList  
        /* noccur = 1 */
        /* Iteration:4 */{ 
          var v_list4 *ClaireSet  
          var v_local4 *ClaireAny  
          v_list4 = x.Subclass
          g0335UU = CreateList(ToType(CEMPTY.Id()),v_list4.Length())
          var CLcount = -1
          for _,z := range(v_list4.Values) { 
            CLcount++
            v_local4 = ANY(F_CALL(ToProperty(C_glb.Id()),ARGS(z.ToEID(),EID{y.Id(),0})))
            g0335UU.PutAt(CLcount,v_local4)
            } 
          /* Iteration-4 */} 
        Result = F_Uall_list(g0335UU)
        /* Let-3 */} 
      /* If!2 */}  else if (C_class.Id() == y.Isa.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var g0333 *ClaireClass   = ToClass(y.Id())
        /* noccur = 1 */
        Result = F_join_class(x,g0333)
        /* Let-3 */} 
      } else {
      Result = ToType(OBJ(F_CALL(ToProperty(C_glb.Id()),ARGS(EID{y.Id(),0},EID{x.Id(),0}))))
      /* If-2 */} 
    return Result} 
  
// The EID go function for: glb @ class (throw: false) 
func E_glb_class (x EID,y EID) EID { 
    return EID{/*(sm for glb @ class= type)*/ F_glb_class(ToClass(OBJ(x)),ToType(OBJ(y)) ).Id(),0}} 
  
/* {1} OPT.The go function for: glb(x:Param,y:type) [] */
func F_glb_Param (x *ClaireParam ,y *ClaireType ) *ClaireType  { 
    // procedure body with s =  
var Result *ClaireType  
    if (y.Isa.IsIn(C_Param) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0336 *ClaireParam   = To_Param(y.Id())
        /* noccur = 3 */
        /* Let:4 */{ 
          var c *ClaireType   = F_join_class(x.Arg,g0336.Arg)
          /* noccur = 3 */
          /* Let:5 */{ 
            var lp *ClaireList   = x.Params.Append(g0336.Params).Set_I().List_I()
            /* noccur = 2 */
            /* Let:6 */{ 
              var l *ClaireList   = ToType(C_any.Id()).EmptyList()
              /* noccur = 3 */
              /* For:7 */{ 
                var p *ClaireAny  
                _ = p
                var p_support *ClaireList  
                p_support = lp
                p_len := p_support.Length()
                for i_it := 0; i_it < p_len; i_it++ { 
                  p = p_support.At(i_it)
                  /* Let:9 */{ 
                    var t *ClaireType   = ToType(OBJ(F_CALL(ToProperty(C_glb.Id()),ARGS(EID{x.At(ToProperty(p)).Id(),0},EID{g0336.At(ToProperty(p)).Id(),0}))))
                    /* noccur = 2 */
                    if (Equal(t.Id(),CEMPTY.Id()) != CTRUE) /* If:10 */{ 
                      l = l.AddFast(t.Id())
                      } else {
                      c = ToType(CEMPTY.Id())
                       /*v = Result, s =void*/

                      break
                      /* If-10 */} 
                    /* Let-9 */} 
                  /* loop-8 */} 
                /* For-7 */} 
              if (Equal(c.Id(),CEMPTY.Id()) != CTRUE) /* If:7 */{ 
                /* Let:8 */{ 
                  var _CL_obj *ClaireParam   = To_Param(new(ClaireParam).Is(C_Param))
                  /* noccur = 7 */
                  _CL_obj.Arg = ToClass(c.Id())
                  _CL_obj.Params = lp
                  _CL_obj.Args = l
                  Result = ToType(_CL_obj.Id())
                  /* Let-8 */} 
                } else {
                Result = ToType(CEMPTY.Id())
                /* If-7 */} 
              /* Let-6 */} 
            /* Let-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (C_class.Id() == y.Isa.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var g0337 *ClaireClass   = ToClass(y.Id())
        /* noccur = 1 */
        /* Let:4 */{ 
          var c *ClaireType   = F_join_class(x.Arg,g0337)
          /* noccur = 2 */
          if (Equal(c.Id(),CEMPTY.Id()) != CTRUE) /* If:5 */{ 
            /* Let:6 */{ 
              var _CL_obj *ClaireParam   = To_Param(new(ClaireParam).Is(C_Param))
              /* noccur = 7 */
              _CL_obj.Arg = ToClass(c.Id())
              _CL_obj.Params = x.Params
              _CL_obj.Args = x.Args
              Result = ToType(_CL_obj.Id())
              /* Let-6 */} 
            } else {
            Result = ToType(CEMPTY.Id())
            /* If-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      } else {
      Result = ToType(OBJ(F_CALL(ToProperty(C_glb.Id()),ARGS(EID{y.Id(),0},EID{x.Id(),0}))))
      /* If-2 */} 
    return Result} 
  
// The EID go function for: glb @ Param (throw: false) 
func E_glb_Param (x EID,y EID) EID { 
    return EID{/*(sm for glb @ Param= type)*/ F_glb_Param(To_Param(OBJ(x)),ToType(OBJ(y)) ).Id(),0}} 
  
// notice that a param whose class is a type must use of (only parameter allowed!)
// the result is a subtype
/* {1} OPT.The go function for: glb(x:subtype,y:type) [] */
func F_glb_subtype (x *ClaireSubtype ,y *ClaireType ) *ClaireType  { 
    // procedure body with s =  
var Result *ClaireType  
    if (C_class.Id() == y.Isa.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var g0339 *ClaireClass   = ToClass(y.Id())
        /* noccur = 2 */
        if (Equal(F_join_class(x.Arg,g0339).Id(),CEMPTY.Id()) != CTRUE) /* If:4 */{ 
          Result = F_nth_class1(ToClass(F_join_class(x.Arg,g0339).Id()),x.T1)
          } else {
          Result = ToType(CEMPTY.Id())
          /* If-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (y.Isa.IsIn(C_Param) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0340 *ClaireParam   = To_Param(y.Id())
        /* noccur = 3 */
        if (Equal(F_join_class(x.Arg,g0340.Arg).Id(),CEMPTY.Id()) != CTRUE) /* If:4 */{ 
          Result = F_param_I_class(ToClass(F_join_class(x.Arg,g0340.Arg).Id()),ToType(OBJ(F_CALL(ToProperty(C_glb.Id()),ARGS(EID{F_member_type(ToType(x.Id())).Id(),0},EID{F_member_type(ToType(g0340.Id())).Id(),0})))))
          } else {
          Result = ToType(CEMPTY.Id())
          /* If-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (y.Isa.IsIn(C_subtype) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0341 *ClaireSubtype   = ToSubtype(y.Id())
        /* noccur = 3 */
        if (Equal(F_join_class(x.Arg,g0341.Arg).Id(),CEMPTY.Id()) != CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var t *ClaireAny   = ANY(F_CALL(ToProperty(C_glb.Id()),ARGS(EID{x.T1.Id(),0},EID{g0341.T1.Id(),0})))
            /* noccur = 2 */
            if (Equal(t,CEMPTY.Id()) != CTRUE) /* If:6 */{ 
              Result = F_nth_class1(ToClass(F_join_class(x.Arg,g0341.Arg).Id()),ToType(t))
              } else {
              Result = ToType(CEMPTY.Id())
              /* If-6 */} 
            /* Let-5 */} 
          } else {
          Result = ToType(CEMPTY.Id())
          /* If-4 */} 
        /* Let-3 */} 
      } else {
      Result = ToType(OBJ(F_CALL(ToProperty(C_glb.Id()),ARGS(EID{y.Id(),0},EID{x.Id(),0}))))
      /* If-2 */} 
    return Result} 
  
// The EID go function for: glb @ subtype (throw: false) 
func E_glb_subtype (x EID,y EID) EID { 
    return EID{/*(sm for glb @ subtype= type)*/ F_glb_subtype(ToSubtype(OBJ(x)),ToType(OBJ(y)) ).Id(),0}} 
  
// set, Interval, list
/* {1} OPT.The go function for: glb(x:tuple,y:type) [] */
func F_glb_tuple (x *ClaireTuple ,y *ClaireType ) *ClaireType  { 
    // procedure body with s =  
var Result *ClaireType  
    if (C_class.Id() == y.Isa.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var g0343 *ClaireClass   = ToClass(y.Id())
        /* noccur = 1 */
        Result = ToType(IfThenElse((C_tuple.IsIn(g0343) == CTRUE),
          x.Id(),
          CEMPTY.Id()))
        /* Let-3 */} 
      /* If!2 */}  else if (y.Isa.IsIn(C_Param) == CTRUE) /* If:2 */{ 
      Result = ToType(CEMPTY.Id())
      /* If!2 */}  else if (C_tuple.Id() == y.Isa.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var g0345 *ClaireTuple   = ToTuple(y.Id())
        /* noccur = 1 */
        Result = ToType(F__exp_list(ToList(x.Id()),ToList(g0345.Id())).Tuple_I().Id())
        /* Let-3 */} 
      /* If!2 */}  else if (y.Isa.IsIn(C_subtype) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0346 *ClaireSubtype   = ToSubtype(y.Id())
        /* noccur = 2 */
        if (g0346.Arg.Id() == C_tuple.Id()) /* If:4 */{ 
          /* Let:5 */{ 
            var g0348UU *ClaireList  
            /* noccur = 1 */
            /* Let:6 */{ 
              var z_bag *ClaireList   = ToType(CEMPTY.Id()).EmptyList()
              /* noccur = 2 */
              /* For:7 */{ 
                var z *ClaireAny  
                _ = z
                var z_support *ClaireList  
                z_support = ToList(x.Id())
                z_len := z_support.Length()
                for i_it := 0; i_it < z_len; i_it++ { 
                  z = z_support.At(i_it)
                  z_bag.AddFast(ANY(F_CALL(ToProperty(C_glb.Id()),ARGS(z.ToEID(),EID{g0346.T1.Id(),0}))))
                  /* loop-8 */} 
                /* For-7 */} 
              g0348UU = z_bag
              /* Let-6 */} 
            Result = ToType(g0348UU.Tuple_I().Id())
            /* Let-5 */} 
          } else {
          Result = ToType(CEMPTY.Id())
          /* If-4 */} 
        /* Let-3 */} 
      } else {
      Result = ToType(OBJ(F_CALL(ToProperty(C_glb.Id()),ARGS(EID{y.Id(),0},EID{x.Id(),0}))))
      /* If-2 */} 
    return Result} 
  
// The EID go function for: glb @ tuple (throw: false) 
func E_glb_tuple (x EID,y EID) EID { 
    return EID{/*(sm for glb @ tuple= type)*/ F_glb_tuple(ToTuple(OBJ(x)),ToType(OBJ(y)) ).Id(),0}} 
  
// a reference is seen as "any"
/* {1} OPT.The go function for: glb(x:Reference,y:type) [] */
func F_glb_Reference (x *ClaireReference ,y *ClaireType ) *ClaireType  { 
    // use function body compiling 
return  y
    } 
  
// The EID go function for: glb @ Reference (throw: false) 
func E_glb_Reference (x EID,y EID) EID { 
    return EID{/*(sm for glb @ Reference= type)*/ F_glb_Reference(To_Reference(OBJ(x)),ToType(OBJ(y)) ).Id(),0}} 
  
// this will be greatly simplified in a few minutes !
/* {1} OPT.The go function for: ^(x:type,y:type) [] */
func F__exp_type (x *ClaireType ,y *ClaireType ) *ClaireType  { 
    // use function body compiling 
return  ToType(OBJ(F_CALL(ToProperty(C_glb.Id()),ARGS(EID{x.Id(),0},EID{y.Id(),0}))))
    } 
  
// The EID go function for: ^ @ type (throw: false) 
func E__exp_type (x EID,y EID) EID { 
    return EID{/*(sm for ^ @ type= type)*/ F__exp_type(ToType(OBJ(x)),ToType(OBJ(y)) ).Id(),0}} 
  
// the old lattice_glb
/* {1} OPT.The go function for: join(x:class,y:class) [] */
func F_join_class (x *ClaireClass ,y *ClaireClass ) *ClaireType  { 
    // procedure body with s =  
var Result *ClaireType  
    /* Let:2 */{ 
      var l1 *ClaireList   = x.Ancestors
      /* noccur = 2 */
      /* Let:3 */{ 
        var n1 int  = l1.Length()
        /* noccur = 2 */
        /* Let:4 */{ 
          var l2 *ClaireList   = y.Ancestors
          /* noccur = 2 */
          /* Let:5 */{ 
            var n2 int  = l2.Length()
            /* noccur = 2 */
            if (n1 < n2) /* If:6 */{ 
              Result = ToType(IfThenElse((l2.ValuesO()[n1-1] == x.Id()),
                y.Id(),
                CEMPTY.Id()))
              /* If!6 */}  else if (l1.ValuesO()[n2-1] == y.Id()) /* If:6 */{ 
              Result = ToType(x.Id())
              } else {
              Result = ToType(CEMPTY.Id())
              /* If-6 */} 
            /* Let-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: join @ class (throw: false) 
func E_join_class (x EID,y EID) EID { 
    return EID{/*(sm for join @ class= type)*/ F_join_class(ToClass(OBJ(x)),ToClass(OBJ(y)) ).Id(),0}} 
  
// for lists
/* {1} OPT.The go function for: ^(x:list,y:list) [] */
func F__exp_list (x *ClaireList ,y *ClaireList ) *ClaireList  { 
    // procedure body with s =  
var Result *ClaireList  
    /* Let:2 */{ 
      var n int  = x.Length()
      /* noccur = 2 */
      /* Let:3 */{ 
        var r *ClaireList   = ToType(CEMPTY.Id()).EmptyList()
        /* noccur = 4 */
        if (n == y.Length()) /* If:4 */{ 
          /* Let:5 */{ 
            var i int  = 1
            /* noccur = 5 */
            /* Let:6 */{ 
              var g0349 int  = n
              /* noccur = 1 */
              for (i <= g0349) /* while:7 */{ 
                /* Let:8 */{ 
                  var z *ClaireTypeExpression   = ToTypeExpression(OBJ(F_CALL(ToProperty(C_glb.Id()),ARGS(x.At(i-1).ToEID(),y.At(i-1).ToEID()))))
                  /* noccur = 2 */
                  if (Equal(z.Id(),CEMPTY.Id()) != CTRUE) /* If:9 */{ 
                    r = r.AddFast(z.Id())
                    } else {
                    r = CNIL
                     /*v = Result, s =void*/

                    break
                    /* If-9 */} 
                  /* Let-8 */} 
                i = (i+1)
                /* while-7 */} 
              /* Let-6 */} 
            /* Let-5 */} 
          /* If-4 */} 
        Result = r
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: ^ @ list (throw: false) 
func E__exp_list (x EID,y EID) EID { 
    return EID{/*(sm for ^ @ list= list)*/ F__exp_list(ToList(OBJ(x)),ToList(OBJ(y)) ).Id(),0}} 
  
// a combined union
/* {1} OPT.The go function for: Uall(l:list) [] */
func F_Uall_list (l *ClaireList ) *ClaireType  { 
    // procedure body with s =  
var Result *ClaireType  
    /* Let:2 */{ 
      var rep *ClaireSet   = CEMPTY
      /* noccur = 3 */
      /* For:3 */{ 
        var x *ClaireAny  
        _ = x
        var x_support *ClaireList  
        x_support = l
        x_len := x_support.Length()
        for i_it := 0; i_it < x_len; i_it++ { 
          x = x_support.At(i_it)
          rep = ToSet(F_U_type(ToType(rep.Id()),ToType(x)).Id())
          /* loop-4 */} 
        /* For-3 */} 
      Result = ToType(rep.Id())
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: Uall @ list (throw: false) 
func E_Uall_list (l EID) EID { 
    return EID{/*(sm for Uall @ list= type)*/ F_Uall_list(ToList(OBJ(l)) ).Id(),0}} 
  
// ------------------- The inclusion operation ------------------------
// the specialized versions %t and <=t are hard coded in Kernel, hence not extensible.
// if we create new types they will be used as patterns, not concrete types.
// hand-made
// v4 open coded (link to Included kernel method)
/* {1} OPT.The go function for: <=t(s:type,y:type) [] */
func F__inf_equalt_type (s *ClaireType ,y *ClaireType ) *ClaireBoolean  { 
    // use function body compiling 
return  s.Included(y)
    } 
  
// The EID go function for: <=t @ type (throw: false) 
func E__inf_equalt_type (s EID,y EID) EID { 
    return EID{/*(sm for <=t @ type= boolean)*/ F__inf_equalt_type(ToType(OBJ(s)),ToType(OBJ(y)) ).Id(),0}} 
  
// default order for types
/* {1} OPT.The go function for: <=(x:type_expression,y:type_expression) [] */
func F__inf_equal_type_expression (x *ClaireTypeExpression ,y *ClaireTypeExpression ) EID { 
    var Result EID 
    if (C_set.Id() == x.Isa.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var g0350 *ClaireSet   = ToSet(x.Id())
        /* noccur = 1 */
        /* Let:4 */{ 
          var g0355UU *ClaireAny  
          /* noccur = 1 */
          var g0355UU_try03565 EID 
          /* For:5 */{ 
            var z *ClaireAny  
            _ = z
            g0355UU_try03565= EID{CFALSE.Id(),0}
            for _,z = range(g0350.Values)/* loop:6 */{ 
              var void_try7 EID 
              _ = void_try7
              var g0357I *ClaireBoolean  
              var g0357I_try03587 EID 
              /* Let:7 */{ 
                var g0359UU *ClaireBoolean  
                /* noccur = 1 */
                var g0359UU_try03608 EID 
                g0359UU_try03608 = F_BELONG(z,y.Id())
                /* ERROR PROTECTION INSERTED (g0359UU-g0357I_try03587) */
                if ErrorIn(g0359UU_try03608) {g0357I_try03587 = g0359UU_try03608
                } else {
                g0359UU = ToBoolean(OBJ(g0359UU_try03608))
                g0357I_try03587 = EID{g0359UU.Not.Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (g0357I-void_try7) */
              if ErrorIn(g0357I_try03587) {void_try7 = g0357I_try03587
              } else {
              g0357I = ToBoolean(OBJ(g0357I_try03587))
              if (g0357I == CTRUE) /* If:7 */{ 
                 /*v = g0355UU_try03565, s =EID*/
g0355UU_try03565 = EID{CTRUE.Id(),0}
                break
                } else {
                void_try7 = EID{CFALSE.Id(),0}
                /* If-7 */} 
              }
              /* ERROR PROTECTION INSERTED (void_try7-g0355UU_try03565) */
              if ErrorIn(void_try7) {g0355UU_try03565 = void_try7
              g0355UU_try03565 = void_try7
              break
              } else {
              }
              /* loop-6 */} 
            /* For-5 */} 
          /* ERROR PROTECTION INSERTED (g0355UU-Result) */
          if ErrorIn(g0355UU_try03565) {Result = g0355UU_try03565
          } else {
          g0355UU = ANY(g0355UU_try03565)
          Result = EID{F_not_any(g0355UU).Id(),0}
          }
          /* Let-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (x.Isa.IsIn(C_type) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0351 *ClaireType   = ToType(x.Id())
        /* noccur = 2 */
        if (y.Isa.IsIn(C_type) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var g0352 *ClaireType   = ToType(y.Id())
            /* noccur = 1 */
            Result = EID{g0351.Included(g0352).Id(),0}
            /* Let-5 */} 
          } else {
          /* Let:5 */{ 
            var z *ClaireAny   = g0351.Id()
            /* noccur = 1 */
            Result = F_CALL(ToProperty(C_less_ask.Id()),ARGS(z.ToEID(),EID{y.Id(),0}))
            /* Let-5 */} 
          /* If-4 */} 
        /* Let-3 */} 
      } else {
      Result = F_CALL(ToProperty(C_less_ask.Id()),ARGS(EID{x.Id(),0},EID{y.Id(),0}))
      /* If-2 */} 
    return RangeCheck(ToType(C_boolean.Id()),Result)} 
  
// The EID go function for: <= @ type_expression (throw: true) 
func E__inf_equal_type_expression (x EID,y EID) EID { 
    return /*(sm for <= @ type_expression= EID)*/ F__inf_equal_type_expression(ToTypeExpression(OBJ(x)),ToTypeExpression(OBJ(y)) )} 
  
// membership for types
// hand-made
// v4 open coded (link to Contains kernel method)
/* {1} OPT.The go function for: %t(x:any,y:type) [] */
func F_Core__Zt_any (x *ClaireAny ,y *ClaireType ) *ClaireBoolean  { 
    // use function body compiling 
return  y.Contains(x)
    } 
  
// The EID go function for: %t @ any (throw: false) 
func E_Core__Zt_any (x EID,y EID) EID { 
    return EID{/*(sm for %t @ any= boolean)*/ F_Core__Zt_any(ANY(x),ToType(OBJ(y)) ).Id(),0}} 
  
// extensibility for type_expression is through less?, that always returns a value (hence no error returned)
/* {1} OPT.The go function for: less?(x:type_expression,y:type_expression) [] */
func F_less_ask_type_expression (x *ClaireTypeExpression ,y *ClaireTypeExpression ) *ClaireBoolean  { 
    // use function body compiling 
return  CFALSE
    } 
  
// The EID go function for: less? @ list<type_expression>(type_expression, type_expression) (throw: false) 
func E_less_ask_type_expression (x EID,y EID) EID { 
    return EID{/*(sm for less? @ list<type_expression>(type_expression, type_expression)= boolean)*/ F_less_ask_type_expression(ToTypeExpression(OBJ(x)),ToTypeExpression(OBJ(y)) ).Id(),0}} 
  
// ******************************************************************
// *    Part 5: type methods                                        *
// ******************************************************************
// --------------------- extract tuple type information -------------
// extract a member type, that is a valid type for all members (z) of instances of
// the type x.This is much simpler in v3.0
/* {1} OPT.The go function for: member(x:type) [] */
func F_member_type (x *ClaireType ) *ClaireType  { 
    // procedure body with s =  
var Result *ClaireType  
    if (C_class.Id() == x.Isa.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var g0361 *ClaireClass   = ToClass(x.Id())
        /* noccur = 1 */
        if (g0361.Id() == C_Interval.Id()) /* If:4 */{ 
          Result = ToType(C_integer.Id())
          } else {
          Result = ToType(C_any.Id())
          /* If-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (x.Isa.IsIn(C_Union) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0362 *ClaireUnion   = To_Union(x.Id())
        /* noccur = 2 */
        Result = F_U_type(F_member_type(g0362.T1),F_member_type(g0362.T2))
        /* Let-3 */} 
      /* If!2 */}  else if (x.Isa.IsIn(C_Interval) == CTRUE) /* If:2 */{ 
      Result = ToType(CEMPTY.Id())
      /* If!2 */}  else if (x.Isa.IsIn(C_Param) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0364 *ClaireParam   = To_Param(x.Id())
        /* noccur = 1 */
        Result = F_member_type(g0364.At(C_of))
        /* Let-3 */} 
      /* If!2 */}  else if (C_tuple.Id() == x.Isa.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var g0365 *ClaireTuple   = ToTuple(x.Id())
        /* noccur = 1 */
        Result = F_Uall_list(ToList(g0365.Id()))
        /* Let-3 */} 
      /* If!2 */}  else if (x.Isa.IsIn(C_subtype) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0366 *ClaireSubtype   = ToSubtype(x.Id())
        /* noccur = 1 */
        Result = g0366.T1
        /* Let-3 */} 
      /* If!2 */}  else if (C_set.Id() == x.Isa.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var g0367 *ClaireSet   = ToSet(x.Id())
        /* noccur = 1 */
        /* Let:4 */{ 
          var g0372UU *ClaireList  
          /* noccur = 1 */
          /* Iteration:5 */{ 
            var v_list5 *ClaireSet  
            var v_local5 *ClaireAny  
            v_list5 = g0367
            g0372UU = CreateList(ToType(CEMPTY.Id()),v_list5.Length())
            var CLcount = -1
            for _,y := range(v_list5.Values) { 
              CLcount++
              if (y.Isa.IsIn(C_list) == CTRUE) /* If:7 */{ 
                /* Let:8 */{ 
                  var g0368 *ClaireList   = ToList(y)
                  /* noccur = 1 */
                  v_local5 = g0368.Set_I().Id()
                  /* Let-8 */} 
                /* If!7 */}  else if (y.Isa.IsIn(C_type) == CTRUE) /* If:7 */{ 
                /* Let:8 */{ 
                  var g0369 *ClaireType   = ToType(y)
                  /* noccur = 1 */
                  v_local5 = g0369.Id()
                  /* Let-8 */} 
                } else {
                v_local5 = CEMPTY.Id()
                /* If-7 */} 
              g0372UU.PutAt(CLcount,v_local5)
              } 
            /* Iteration-5 */} 
          Result = F_Uall_list(g0372UU)
          /* Let-4 */} 
        /* Let-3 */} 
      } else {
      Result = ToType(CEMPTY.Id())
      /* If-2 */} 
    return Result} 
  
// The EID go function for: member @ type (throw: false) 
func E_member_type (x EID) EID { 
    return EID{/*(sm for member @ type= type)*/ F_member_type(ToType(OBJ(x)) ).Id(),0}} 
  
// a simpler version (projection on bag subtypes)
// dumb code because it is used early in the bootstrap
/* {1} OPT.The go function for: of_extract(x:type) [] */
func F_of_extract_type (x *ClaireType ) *ClaireType  { 
    // procedure body with s =  
var Result *ClaireType  
    /* Let:2 */{ 
      var c *ClaireClass   = x.Isa
      /* noccur = 2 */
      if (c.Id() == C_subtype.Id()) /* If:3 */{ 
        Result = ToSubtype(x.Id()).T1
        /* If!3 */}  else if (c.Id() == C_Param.Id()) /* If:3 */{ 
        if (To_Param(x.Id()).Params.At(1-1) == C_of.Id()) /* If:4 */{ 
          /* Let:5 */{ 
            var y *ClaireType   = ToType(To_Param(x.Id()).Args.At(1-1))
            /* noccur = 4 */
            if (C_set.Id() == y.Isa.Id()) /* If:6 */{ 
              /* Let:7 */{ 
                var g0373 *ClaireSet   = ToSet(y.Id())
                /* noccur = 1 */
                Result = ToType(g0373.List_I().At(1-1))
                /* Let-7 */} 
              /* If!6 */}  else if (y.Isa.IsIn(C_subtype) == CTRUE) /* If:6 */{ 
              /* Let:7 */{ 
                var g0374 *ClaireSubtype   = ToSubtype(y.Id())
                /* noccur = 1 */
                Result = g0374.T1
                /* Let-7 */} 
              } else {
              Result = ToType(C_any.Id())
              /* If-6 */} 
            /* Let-5 */} 
          } else {
          Result = ToType(C_any.Id())
          /* If-4 */} 
        } else {
        Result = ToType(C_any.Id())
        /* If-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: of_extract @ type (throw: false) 
func E_of_extract_type (x EID) EID { 
    return EID{/*(sm for of_extract @ type= type)*/ F_of_extract_type(ToType(OBJ(x)) ).Id(),0}} 
  
// useful type functions for the compiler
/* {1} OPT.The go function for: unique?(x:type) [] */
func F_unique_ask_type (x *ClaireType ) *ClaireBoolean  { 
    // procedure body with s =  
var Result *ClaireBoolean  
    if (C_set.Id() == x.Isa.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var g0376 *ClaireSet   = ToSet(x.Id())
        /* noccur = 1 */
        Result = Equal(MakeInteger(g0376.Size()).Id(),MakeInteger(1).Id())
        /* Let-3 */} 
      /* If!2 */}  else if (C_class.Id() == x.Isa.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var g0377 *ClaireClass   = ToClass(x.Id())
        /* noccur = 2 */
        Result = MakeBoolean((g0377.Open == 0) && (F_size_class(g0377) == 1))
        /* Let-3 */} 
      } else {
      Result = CFALSE
      /* If-2 */} 
    return Result} 
  
// The EID go function for: unique? @ type (throw: false) 
func E_unique_ask_type (x EID) EID { 
    return EID{/*(sm for unique? @ type= boolean)*/ F_unique_ask_type(ToType(OBJ(x)) ).Id(),0}} 
  
// returns the unique element of the type
/* {1} OPT.The go function for: the(x:type) [] */
func F_the_type (x *ClaireType ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var g0379UU *ClaireList  
      /* noccur = 1 */
      var g0379UU_try03803 EID 
      /* Let:3 */{ 
        var g0381UU *ClaireAny  
        /* noccur = 1 */
        var g0381UU_try03824 EID 
        g0381UU_try03824 = F_CALL(C_set_I,ARGS(EID{x.Id(),0}))
        /* ERROR PROTECTION INSERTED (g0381UU-g0379UU_try03803) */
        if ErrorIn(g0381UU_try03824) {g0379UU_try03803 = g0381UU_try03824
        } else {
        g0381UU = ANY(g0381UU_try03824)
        g0379UU_try03803 = EID{ToSet(g0381UU).List_I().Id(),0}
        }
        /* Let-3 */} 
      /* ERROR PROTECTION INSERTED (g0379UU-Result) */
      if ErrorIn(g0379UU_try03803) {Result = g0379UU_try03803
      } else {
      g0379UU = ToList(OBJ(g0379UU_try03803))
      Result = g0379UU.At(1-1).ToEID()
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: the @ type (throw: true) 
func E_the_type (x EID) EID { 
    return /*(sm for the @ type= EID)*/ F_the_type(ToType(OBJ(x)) )} 
  
// bitvector made easy
// v0.01: should not use set[0 .. 29] => burden on caller is too heavy
/* {1} OPT.The go function for: integer!(s:set[integer]) [] */
func F_integer_I_set (s *ClaireSet ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var n int  = 0
      /* noccur = 3 */
      /* For:3 */{ 
        var y int 
        _ = y
        var y_iter *ClaireAny  
        Result= EID{CFALSE.Id(),0}
        for _,y_iter = range(s.Values)/* loop:4 */{ 
          y = ToInteger(y_iter).Value
          var void_try5 EID 
          _ = void_try5
          if ((y >= 0) && 
              (y <= 29)) /* If:5 */{ 
            var n_try03836 EID 
            /* Let:6 */{ 
              var g0384UU int 
              /* noccur = 1 */
              var g0384UU_try03857 EID 
              g0384UU_try03857 = F__exp2_integer(y)
              /* ERROR PROTECTION INSERTED (g0384UU-n_try03836) */
              if ErrorIn(g0384UU_try03857) {n_try03836 = g0384UU_try03857
              } else {
              g0384UU = INT(g0384UU_try03857)
              n_try03836 = EID{C__INT,IVAL((n+g0384UU))}
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (n-void_try5) */
            if ErrorIn(n_try03836) {void_try5 = n_try03836
            } else {
            n = INT(n_try03836)
            void_try5 = EID{C__INT,IVAL(n)}
            }
            } else {
            void_try5 = EID{CFALSE.Id(),0}
            /* If-5 */} 
          /* ERROR PROTECTION INSERTED (void_try5-Result) */
          if ErrorIn(void_try5) {Result = void_try5
          Result = void_try5
          break
          } else {
          }
          /* loop-4 */} 
        /* For-3 */} 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = EID{C__INT,IVAL(n)}
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: integer! @ set (throw: true) 
func E_integer_I_set (s EID) EID { 
    return /*(sm for integer! @ set= EID)*/ F_integer_I_set(ToSet(OBJ(s)) )} 
  
/* {1} OPT.The go function for: make_set(x:integer) [] */
func F_make_set_integer (x int) *ClaireSet  { 
    // procedure body with s =  
var Result *ClaireSet  
    /* Let:2 */{ 
      var i_out *ClaireSet   = ToType(CEMPTY.Id()).EmptySet()
      /* noccur = 2 */
      /* Let:3 */{ 
        var i int  = 0
        /* noccur = 5 */
        /* Let:4 */{ 
          var g0386 int  = 29
          /* noccur = 1 */
          for (i <= g0386) /* while:5 */{ 
            if (F_nth_integer(x,i) == CTRUE) /* If:6 */{ 
              i_out.AddFast(MakeInteger(i).Id())
              /* If-6 */} 
            i = (i+1)
            /* while-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      Result = i_out
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: make_set @ integer (throw: false) 
func E_make_set_integer (x EID) EID { 
    return EID{/*(sm for make_set @ integer= set)*/ F_make_set_integer(INT(x) ).Id(),0}} 
  
// asbtract coercion of a set into an interval
/* {1} OPT.The go function for: abstract_type(xt1:set) [] */
func F_abstract_type_set (xt1 *ClaireSet ) *ClaireType  { 
    // procedure body with s =  
var Result *ClaireType  
    /* Let:2 */{ 
      var m1 int  = 1
      /* noccur = 6 */
      /* Let:3 */{ 
        var m2 int  = 0
        /* noccur = 6 */
        /* For:4 */{ 
          var x *ClaireAny  
          _ = x
          for _,x = range(xt1.Values)/* loop:5 */{ 
            if (C_integer.Id() == x.Isa.Id()) /* If:6 */{ 
              /* Let:7 */{ 
                var g0387 int  = ToInteger(x).Value
                /* noccur = 6 */
                if (m1 > m2) /* If:8 */{ 
                  m1 = g0387
                  m2 = g0387
                  /* If!8 */}  else if (g0387 > m2) /* If:8 */{ 
                  m2 = g0387
                  /* If!8 */}  else if (g0387 < m1) /* If:8 */{ 
                  m1 = g0387
                  /* If-8 */} 
                /* Let-7 */} 
              } else {
              m1 = 1
              m2 = 0
               /*v = Result, s =void*/

              break
              /* If-6 */} 
            /* loop-5 */} 
          /* For-4 */} 
        Result = F__dot_dot_integer(m1,m2)
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: abstract_type @ set (throw: false) 
func E_abstract_type_set (xt1 EID) EID { 
    return EID{/*(sm for abstract_type @ set= type)*/ F_abstract_type_set(ToSet(OBJ(xt1)) ).Id(),0}} 
  
// abstract interpretation of integer arithmetique
/* {1} OPT.The go function for: abstract_type(p:operation,xt1:type,xt2:type) [] */
func F_abstract_type_operation (p *ClaireOperation ,xt1 *ClaireType ,xt2 *ClaireType ) *ClaireType  { 
    // procedure body with s =  
var Result *ClaireType  
    if (C_set.Id() == xt1.Isa.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var g0389 *ClaireSet   = ToSet(xt1.Id())
        /* noccur = 3 */
        if (Equal(g0389.Id(),CEMPTY.Id()) != CTRUE) /* If:4 */{ 
          Result = F_abstract_type_operation(p,F_abstract_type_set(g0389),xt2)
          } else {
          Result = ToType(g0389.Id())
          /* If-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (xt1.Isa.IsIn(C_Interval) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0390 *ClaireInterval   = To_Interval(xt1.Id())
        /* noccur = 7 */
        if (xt2.Isa.IsIn(C_Interval) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var g0391 *ClaireInterval   = To_Interval(xt2.Id())
            /* noccur = 4 */
            if (p.Id() == C__plus.Id()) /* If:6 */{ 
              Result = F__dot_dot_integer((g0390.Arg1+g0391.Arg1),(g0390.Arg2+g0391.Arg2))
              /* If!6 */}  else if (p.Id() == C__dash.Id()) /* If:6 */{ 
              Result = F__dot_dot_integer((g0390.Arg1-g0391.Arg2),(g0390.Arg2-g0391.Arg1))
              } else {
              Result = ToType(C_integer.Id())
              /* If-6 */} 
            /* Let-5 */} 
          /* If!4 */}  else if (C_set.Id() == xt2.Isa.Id()) /* If:4 */{ 
          /* Let:5 */{ 
            var g0392 *ClaireSet   = ToSet(xt2.Id())
            /* noccur = 3 */
            if (Equal(g0392.Id(),CEMPTY.Id()) != CTRUE) /* If:6 */{ 
              Result = F_abstract_type_operation(p,ToType(g0390.Id()),F_abstract_type_set(g0392))
              } else {
              Result = ToType(g0392.Id())
              /* If-6 */} 
            /* Let-5 */} 
          /* If!4 */}  else if (xt2.Isa.IsIn(C_Union) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var g0393 *ClaireUnion   = To_Union(xt2.Id())
            /* noccur = 2 */
            Result = F_U_type(F_abstract_type_operation(p,ToType(g0390.Id()),g0393.T1),F_abstract_type_operation(p,ToType(g0390.Id()),g0393.T2))
            /* Let-5 */} 
          } else {
          Result = ToType(C_integer.Id())
          /* If-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (xt1.Isa.IsIn(C_Union) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0395 *ClaireUnion   = To_Union(xt1.Id())
        /* noccur = 2 */
        Result = F_U_type(F_abstract_type_operation(p,g0395.T1,xt2),F_abstract_type_operation(p,g0395.T2,xt2))
        /* Let-3 */} 
      } else {
      Result = ToType(C_integer.Id())
      /* If-2 */} 
    return Result} 
  
// The EID go function for: abstract_type @ operation (throw: false) 
func E_abstract_type_operation (p EID,xt1 EID,xt2 EID) EID { 
    return EID{/*(sm for abstract_type @ operation= type)*/ F_abstract_type_operation(ToOperation(OBJ(p)),ToType(OBJ(xt1)),ToType(OBJ(xt2)) ).Id(),0}} 
  
// we create some types that we need
// a useful second ortder type
/* {1} OPT.The go function for: first_arg_type(x:type,y:type) [] */
func F_first_arg_type_type (x *ClaireType ,y *ClaireType ) *ClaireType  { 
    // use function body compiling 
return  x
    } 
  
// The EID go function for: first_arg_type @ list<type_expression>(type, type) (throw: false) 
func E_first_arg_type_type (x EID,y EID) EID { 
    return EID{/*(sm for first_arg_type @ list<type_expression>(type, type)= type)*/ F_first_arg_type_type(ToType(OBJ(x)),ToType(OBJ(y)) ).Id(),0}} 
  
/* {1} OPT.The go function for: first_arg_type(x:type,y:type,z:type) [] */
func F_first_arg_type_type2 (x *ClaireType ,y *ClaireType ,z *ClaireType ) *ClaireType  { 
    // use function body compiling 
return  x
    } 
  
// The EID go function for: first_arg_type @ list<type_expression>(type, type, type) (throw: false) 
func E_first_arg_type_type2 (x EID,y EID,z EID) EID { 
    return EID{/*(sm for first_arg_type @ list<type_expression>(type, type, type)= type)*/ F_first_arg_type_type2(ToType(OBJ(x)),ToType(OBJ(y)),ToType(OBJ(z)) ).Id(),0}} 
  
/* {1} OPT.The go function for: second_arg_type(x:type,y:type) [] */
func F_second_arg_type_type (x *ClaireType ,y *ClaireType ) *ClaireType  { 
    // use function body compiling 
return  y
    } 
  
// The EID go function for: second_arg_type @ type (throw: false) 
func E_second_arg_type_type (x EID,y EID) EID { 
    return EID{/*(sm for second_arg_type @ type= type)*/ F_second_arg_type_type(ToType(OBJ(x)),ToType(OBJ(y)) ).Id(),0}} 
  
/* {1} OPT.The go function for: meet_arg_types(x:type,y:type) [] */
func F_meet_arg_types_type (x *ClaireType ,y *ClaireType ) *ClaireType  { 
    // use function body compiling 
return  F_U_type(x,y)
    } 
  
// The EID go function for: meet_arg_types @ type (throw: false) 
func E_meet_arg_types_type (x EID,y EID) EID { 
    return EID{/*(sm for meet_arg_types @ type= type)*/ F_meet_arg_types_type(ToType(OBJ(x)),ToType(OBJ(y)) ).Id(),0}} 
  
/* {1} OPT.The go function for: first_member_type(x:type,y:type) [] */
func F_first_member_type_type (x *ClaireType ,y *ClaireType ) *ClaireType  { 
    // use function body compiling 
return  F_member_type(x)
    } 
  
// The EID go function for: first_member_type @ type (throw: false) 
func E_first_member_type_type (x EID,y EID) EID { 
    return EID{/*(sm for first_member_type @ type= type)*/ F_first_member_type_type(ToType(OBJ(x)),ToType(OBJ(y)) ).Id(),0}} 
  
// v3.3.10
// nth@bag (list / set) is now in Kernel (CLAIRE4)
/* {1} OPT.The go function for: nth_arg_type(x:type,y:type) [] */
func F_Core_nth_arg_type_type (x *ClaireType ,y *ClaireType ) EID { 
    var Result EID 
    if ((C_tuple.Id() == x.Isa.Id()) && 
        (F_unique_ask_type(y) == CTRUE)) /* If:2 */{ 
      /* Let:3 */{ 
        var g0397UU *ClaireAny  
        /* noccur = 1 */
        var g0397UU_try03984 EID 
        g0397UU_try03984 = F_the_type(y)
        /* ERROR PROTECTION INSERTED (g0397UU-Result) */
        if ErrorIn(g0397UU_try03984) {Result = g0397UU_try03984
        } else {
        g0397UU = ANY(g0397UU_try03984)
        Result = F_CALL(C_nth,ARGS(EID{x.Id(),0},g0397UU.ToEID()))
        }
        /* Let-3 */} 
      } else {
      Result = EID{F_member_type(x).Id(),0}
      /* If-2 */} 
    return RangeCheck(ToType(C_type.Id()),Result)} 
  
// The EID go function for: nth_arg_type @ type (throw: true) 
func E_Core_nth_arg_type_type (x EID,y EID) EID { 
    return /*(sm for nth_arg_type @ type= EID)*/ F_Core_nth_arg_type_type(ToType(OBJ(x)),ToType(OBJ(y)) )} 
  
// we place here all methods that require second order types !!!!
/* {1} OPT.The go function for: nth_get(a:array,n:integer) [] */
func F_nth_get_array (a *ClaireList ,n int) *ClaireAny  { 
    // use function body compiling 
return  ToList(a.Id()).At(n-1)
    } 
  
// The EID go function for: nth_get @ array (throw: false) 
func E_nth_get_array (a EID,n EID) EID { 
    return /*(sm for nth_get @ array= any)*/ F_nth_get_array(ToArray(OBJ(a)),INT(n) ).ToEID()} 
  
/* {1} OPT.The go function for: nth_get_array_type */
func F_nth_get_array_type (a *ClaireType ,n *ClaireType ) EID { 
    /* eid body: member(a) */
    var Result EID 
    Result = F_CALL(C_member,ARGS(EID{a.Id(),0}))
    return Result} 
  
  
// The dual EID go function for: "nth_get_array_type" 
func E_nth_get_array_type (a EID,n EID) EID { 
    return F_nth_get_array_type(ToType(OBJ(a)),ToType(OBJ(n)))} 
  
// managed by cross-compiler ?
/* {1} OPT.The go function for: nth(self:array,x:integer) [] */
func F_nth_array (self *ClaireList ,x int) EID { 
    var Result EID 
    if ((x > 0) && 
        (x <= self.Length())) /* If:2 */{ 
      Result = ToList(self.Id()).At(x-1).ToEID()
      } else {
      Result = ToException(C_general_error.Make(MakeString("[180] nth[~S] out of scope for ~S").Id(),MakeConstantList(MakeInteger(x).Id(),self.Id()).Id())).Close()
      /* If-2 */} 
    return Result} 
  
// The EID go function for: nth @ array (throw: true) 
func E_nth_array (self EID,x EID) EID { 
    return /*(sm for nth @ array= EID)*/ F_nth_array(ToArray(OBJ(self)),INT(x) )} 
  
/* {1} OPT.The go function for: nth_array_type */
func F_nth_array_type (self *ClaireType ,x *ClaireType ) EID { 
    /* eid body: member(self) */
    var Result EID 
    Result = F_CALL(C_member,ARGS(EID{self.Id(),0}))
    return Result} 
  
  
// The dual EID go function for: "nth_array_type" 
func E_nth_array_type (self EID,x EID) EID { 
    return F_nth_array_type(ToType(OBJ(self)),ToType(OBJ(x)))} 
  
/* {1} OPT.The go function for: make_array_integer_type */
func F_make_array_integer_type (i *ClaireType ,t *ClaireType ,v *ClaireType ) EID { 
    /* eid body: (if unique?(t) the(t)[] else array) */
    var Result EID 
    if (ToBoolean(OBJ(F_CALL(C_unique_ask,ARGS(EID{t.Id(),0})))) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0399UU *ClaireAny  
        /* noccur = 1 */
        var g0399UU_try04004 EID 
        g0399UU_try04004 = F_CALL(C_the,ARGS(EID{t.Id(),0}))
        /* ERROR PROTECTION INSERTED (g0399UU-Result) */
        if ErrorIn(g0399UU_try04004) {Result = g0399UU_try04004
        } else {
        g0399UU = ANY(g0399UU_try04004)
        Result = F_CALL(C_nth,ARGS(g0399UU.ToEID()))
        }
        /* Let-3 */} 
      } else {
      Result = EID{C_array.Id(),0}
      /* If-2 */} 
    return Result} 
  
  
// The dual EID go function for: "make_array_integer_type" 
func E_make_array_integer_type (i EID,t EID,v EID) EID { 
    return F_make_array_integer_type(ToType(OBJ(i)),ToType(OBJ(t)),ToType(OBJ(v)))} 
  
/* {1} OPT.The go function for: make_list(n:integer,t:type,x:any) [] */
func F_make_list_integer2 (n int,t *ClaireType ,x *ClaireAny ) *ClaireList  { 
    // use function body compiling 
return  ToList(F_make_list_integer(n,x).Cast_I(t).Id())
    } 
  
// The EID go function for: make_list @ list<type_expression>(integer, type, any) (throw: false) 
func E_make_list_integer2 (n EID,t EID,x EID) EID { 
    return EID{/*(sm for make_list @ list<type_expression>(integer, type, any)= list)*/ F_make_list_integer2(INT(n),ToType(OBJ(t)),ANY(x) ).Id(),0}} 
  
/* {1} OPT.The go function for: make_list_integer2_type */
func F_make_list_integer2_type (n *ClaireType ,t *ClaireType ,x *ClaireType ) EID { 
    /* eid body: (if unique?(t) list[the(t)] else list) */
    var Result EID 
    if (ToBoolean(OBJ(F_CALL(C_unique_ask,ARGS(EID{t.Id(),0})))) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0401UU *ClaireAny  
        /* noccur = 1 */
        var g0401UU_try04024 EID 
        g0401UU_try04024 = F_CALL(C_the,ARGS(EID{t.Id(),0}))
        /* ERROR PROTECTION INSERTED (g0401UU-Result) */
        if ErrorIn(g0401UU_try04024) {Result = g0401UU_try04024
        } else {
        g0401UU = ANY(g0401UU_try04024)
        Result = F_CALL(C_nth,ARGS(EID{C_list.Id(),0},g0401UU.ToEID()))
        }
        /* Let-3 */} 
      } else {
      Result = EID{C_list.Id(),0}
      /* If-2 */} 
    return Result} 
  
  
// The dual EID go function for: "make_list_integer2_type" 
func E_make_list_integer2_type (n EID,t EID,x EID) EID { 
    return F_make_list_integer2_type(ToType(OBJ(n)),ToType(OBJ(t)),ToType(OBJ(x)))} 
  
/* {1} OPT.The go function for: make_set(self:array[of:(any)]) [] */
func F_make_set_array (self *ClaireList ) *ClaireSet  { 
    // use function body compiling 
return  F_list_I_array(self).Set_I()
    } 
  
// The EID go function for: make_set @ array (throw: false) 
func E_make_set_array (self EID) EID { 
    return EID{/*(sm for make_set @ array= set)*/ F_make_set_array(ToArray(OBJ(self)) ).Id(),0}} 
  
/* {1} OPT.The go function for: make_set_array_type */
func F_make_set_array_type (self *ClaireType ) EID { 
    /* eid body: (if (member(self @ of) = any) set else set[list<any>(of),list(set(member(self @ of)))]) */
    var Result EID 
    if (ToBoolean(OBJ(F_CALL(ToProperty(C__equal.Id()),ARGS(F_CALL(C_member,ARGS(F_CALL(ToProperty(C__at.Id()),ARGS(EID{self.Id(),0},EID{C_of.Id(),0})))),EID{C_any.Id(),0})))) == CTRUE) /* If:2 */{ 
      Result = EID{C_set.Id(),0}
      } else {
      Result = F_CALL(C_nth,ARGS(EID{C_set.Id(),0},EID{MakeList(ToType(C_any.Id()),C_of.Id()).Id(),0},EID{MakeConstantList(MakeConstantSet(ANY(F_CALL(C_member,ARGS(F_CALL(ToProperty(C__at.Id()),ARGS(EID{self.Id(),0},EID{C_of.Id(),0})))))).Id()).Id(),0}))
      /* If-2 */} 
    return Result} 
  
  
// The dual EID go function for: "make_set_array_type" 
func E_make_set_array_type (self EID) EID { 
    return F_make_set_array_type(ToType(OBJ(self)))} 
  
// these four functions are defined in Core with Kernel functions because we want to
// add second order types
/* {1} OPT.The go function for: list_I_array_type */
func F_list_I_array_type (a *ClaireType ) EID { 
    /* eid body: (if (member(a @ of) = any) list else list[list<any>(of),list(set(member(a @ of)))]) */
    var Result EID 
    if (ToBoolean(OBJ(F_CALL(ToProperty(C__equal.Id()),ARGS(F_CALL(C_member,ARGS(F_CALL(ToProperty(C__at.Id()),ARGS(EID{a.Id(),0},EID{C_of.Id(),0})))),EID{C_any.Id(),0})))) == CTRUE) /* If:2 */{ 
      Result = EID{C_list.Id(),0}
      } else {
      Result = F_CALL(C_nth,ARGS(EID{C_list.Id(),0},EID{MakeList(ToType(C_any.Id()),C_of.Id()).Id(),0},EID{MakeConstantList(MakeConstantSet(ANY(F_CALL(C_member,ARGS(F_CALL(ToProperty(C__at.Id()),ARGS(EID{a.Id(),0},EID{C_of.Id(),0})))))).Id()).Id(),0}))
      /* If-2 */} 
    return Result} 
  
  
// The dual EID go function for: "list_I_array_type" 
func E_list_I_array_type (a EID) EID { 
    return F_list_I_array_type(ToType(OBJ(a)))} 
  
/* {1} OPT.The go function for: array_I_list_type */
func F_array_I_list_type (a *ClaireType ) EID { 
    /* eid body: (if (member(a @ of) = any) array else array[list<any>(of),list(set(member(a @ of)))]) */
    var Result EID 
    if (ToBoolean(OBJ(F_CALL(ToProperty(C__equal.Id()),ARGS(F_CALL(C_member,ARGS(F_CALL(ToProperty(C__at.Id()),ARGS(EID{a.Id(),0},EID{C_of.Id(),0})))),EID{C_any.Id(),0})))) == CTRUE) /* If:2 */{ 
      Result = EID{C_array.Id(),0}
      } else {
      Result = F_CALL(C_nth,ARGS(EID{C_array.Id(),0},EID{MakeList(ToType(C_any.Id()),C_of.Id()).Id(),0},EID{MakeConstantList(MakeConstantSet(ANY(F_CALL(C_member,ARGS(F_CALL(ToProperty(C__at.Id()),ARGS(EID{a.Id(),0},EID{C_of.Id(),0})))))).Id()).Id(),0}))
      /* If-2 */} 
    return Result} 
  
  
// The dual EID go function for: "array_I_list_type" 
func E_array_I_list_type (a EID) EID { 
    return F_array_I_list_type(ToType(OBJ(a)))} 
  
// v3.0.72
/* {1} OPT.The go function for: set_I_list_type */
func F_set_I_list_type (l *ClaireType ) EID { 
    /* eid body: (if (member(l @ of) = any) set else set[list<any>(of),list(set(member(l @ of)))]) */
    var Result EID 
    if (ToBoolean(OBJ(F_CALL(ToProperty(C__equal.Id()),ARGS(F_CALL(C_member,ARGS(F_CALL(ToProperty(C__at.Id()),ARGS(EID{l.Id(),0},EID{C_of.Id(),0})))),EID{C_any.Id(),0})))) == CTRUE) /* If:2 */{ 
      Result = EID{C_set.Id(),0}
      } else {
      Result = F_CALL(C_nth,ARGS(EID{C_set.Id(),0},EID{MakeList(ToType(C_any.Id()),C_of.Id()).Id(),0},EID{MakeConstantList(MakeConstantSet(ANY(F_CALL(C_member,ARGS(F_CALL(ToProperty(C__at.Id()),ARGS(EID{l.Id(),0},EID{C_of.Id(),0})))))).Id()).Id(),0}))
      /* If-2 */} 
    return Result} 
  
  
// The dual EID go function for: "set_I_list_type" 
func E_set_I_list_type (l EID) EID { 
    return F_set_I_list_type(ToType(OBJ(l)))} 
  
/* {1} OPT.The go function for: list_I_set_type */
func F_list_I_set_type (l *ClaireType ) EID { 
    /* eid body: (if (member(l @ of) = any) list else list[list<any>(of),list(set(member(l @ of)))]) */
    var Result EID 
    if (ToBoolean(OBJ(F_CALL(ToProperty(C__equal.Id()),ARGS(F_CALL(C_member,ARGS(F_CALL(ToProperty(C__at.Id()),ARGS(EID{l.Id(),0},EID{C_of.Id(),0})))),EID{C_any.Id(),0})))) == CTRUE) /* If:2 */{ 
      Result = EID{C_list.Id(),0}
      } else {
      Result = F_CALL(C_nth,ARGS(EID{C_list.Id(),0},EID{MakeList(ToType(C_any.Id()),C_of.Id()).Id(),0},EID{MakeConstantList(MakeConstantSet(ANY(F_CALL(C_member,ARGS(F_CALL(ToProperty(C__at.Id()),ARGS(EID{l.Id(),0},EID{C_of.Id(),0})))))).Id()).Id(),0}))
      /* If-2 */} 
    return Result} 
  
  
// The dual EID go function for: "list_I_set_type" 
func E_list_I_set_type (l EID) EID { 
    return F_list_I_set_type(ToType(OBJ(l)))} 
  
// get the type from class if a constant
/* {1} OPT.The go function for: thing_type_class(x:type) [] */
func F_Core_thing_type_class_type (x *ClaireType ) *ClaireType  { 
    // use function body compiling 
return  F_glb_class(C_thing,F_member_type(x))
    } 
  
// The EID go function for: thing_type_class @ type (throw: false) 
func E_Core_thing_type_class_type (x EID) EID { 
    return EID{/*(sm for thing_type_class @ type= type)*/ F_Core_thing_type_class_type(ToType(OBJ(x)) ).Id(),0}} 
  
/* {1} OPT.The go function for: object_type_class(x:type) [] */
func F_Core_object_type_class_type (x *ClaireType ) *ClaireType  { 
    // use function body compiling 
return  F_glb_class(C_object,F_member_type(x))
    } 
  
// The EID go function for: object_type_class @ type (throw: false) 
func E_Core_object_type_class_type (x EID) EID { 
    return EID{/*(sm for object_type_class @ type= type)*/ F_Core_object_type_class_type(ToType(OBJ(x)) ).Id(),0}} 
  
// new in v3.0.60 : second-order type for copy