/***** CLAIRE Compilation of file /Users/ycaseau/Dropbox/src/clairev4.0/src/meta/function.cl 
         [version 4.0.02 / safety 5] Saturday 10-30-2021 *****/

package Core
import (_ "fmt"
	. "Kernel"
)

//+-------------------------------------------------------------+
//| CLAIRE                                                      |
//| function.cl                                                 |
//| Copyright (C) 1994 - 2021 Yves Caseau. All Rights Reserved  |
//| cf. copyright info in file object.cl: about()               |
//+-------------------------------------------------------------+
// --------------------------------------------------------------------
// This file hold the methods that are defined by an external function
// and those that apply to a primitive type
// --------------------------------------------------------------------
// *********************************************************************
// *  Table of contents                                                *
// *   Part 1: Basics of pretty printing                               *
// *   Part 2: Methods for CLAIRE objects                              *
// *   Part 3: System Methods                                          *
// *   Part 4: Methods for Native entities                             *
// *********************************************************************
// we find here what is necessary for the minimal kernel of CLAIRE
// ==============================================================
// min :: operation(precedence = 20)
// max :: operation(precedence = 20)
// *********************************************************************
// *   Part 1: Basics of pretty printing                               *
// *********************************************************************
// we use a nice object
// support reccursive print-in-string 
// print :: property() - defined in Kernel
// buffered print
// new in v3.3.26: unbounded recursion is supported :-)
/* {1} OPT.The go function for: print_in_string(_CL_obj:void) [] */
func F_print_in_string_void ()  { 
    // procedure body with s =  
/* Let:2 */{ 
      var n int  = (C_pretty.Cprevious+1)
      /* noccur = 5 */
      /* Let:3 */{ 
        var p1 *ClairePort  
        /* noccur = 4 */
        if (n < C_pretty.Cpstack.Length()) /* If:4 */{ 
          p1 = ToPort(C_pretty.Cpstack.At((n+1)-1))
          } else {
          p1 = F_port_I_void()
          /* If-4 */} 
        /* Let:4 */{ 
          var p2 *ClairePort   = p1.UseAsOutput()
          /* noccur = 2 */
          C_pretty.Cprevious = n
          C_pretty.Cpretty = p1
          if (Equal(C_pretty.Cpstack.Id(),CNIL.Id()) == CTRUE) /* If:5 */{ 
            C_pretty.Cpstack = MakeList(ToType(C_port.Id()),p2.Id(),p1.Id())
            } else {
            ToArray(C_pretty.Cpstack.Id()).NthPut(n,p2.Id())
            if (n == C_pretty.Cpstack.Length()) /* If:6 */{ 
              C_pretty.Cpstack = C_pretty.Cpstack.AddFast(p1.Id())
              /* If-6 */} 
            /* If-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    } 
  
// The EID go function for: print_in_string @ void (throw: false) 
func E_print_in_string_void (_CL_obj EID) EID { 
    /*(sm for print_in_string @ void= void)*/ F_print_in_string_void( )
    return EVOID} 
  
/* {1} OPT.The go function for: end_of_string(_CL_obj:void) [] */
func F_end_of_string_void () EID { 
    var Result EID 
    if (C_pretty.Cprevious == 0) /* If:2 */{ 
      Result = ToException(C_general_error.Make(MakeString("[123] unbalanced use of print-in-string").Id(),CNIL.Id())).Close()
      } else {
      Result = EID{CFALSE.Id(),0}
      /* If-2 */} 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    /* Let:2 */{ 
      var n int  = C_pretty.Cprevious
      /* noccur = 1 */
      /* Let:3 */{ 
        var s *ClaireString   = C_pretty.Cpretty.String_I()
        /* noccur = 1 */
        /* Let:4 */{ 
          var p *ClairePort   = ToPort(C_pretty.Cpstack.At(n-1))
          /* noccur = 2 */
          C_pretty.Cpretty.SetLength(0)
          p.UseAsOutput()
          C_pretty.Cpretty = p
          C_pretty.Cprevious = (C_pretty.Cprevious-1)
          Result = EID{(s).Id(),0}
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    }
    return Result} 
  
// The EID go function for: end_of_string @ void (throw: true) 
func E_end_of_string_void (_CL_obj EID) EID { 
    return /*(sm for end_of_string @ void= EID)*/ F_end_of_string_void( )} 
  
/* {1} OPT.The go function for: mClaire/buffer_length(_CL_obj:void) [] */
func F_buffer_length_void () int { 
    // use function body compiling 
return  C_pretty.Cpretty.Length()
    } 
  
// The EID go function for: mClaire/buffer_length @ void (throw: false) 
func E_buffer_length_void (_CL_obj EID) EID { 
    return EID{C__INT,IVAL(/*(sm for mClaire/buffer_length @ void= integer)*/ F_buffer_length_void( ))}} 
  
/* {1} OPT.The go function for: mClaire/buffer_set_length(i:integer) [] */
func F_buffer_set_length_integer (i int)  { 
    // procedure body with s =  
C_pretty.Cpretty.SetLength(i)
    } 
  
// The EID go function for: mClaire/buffer_set_length @ integer (throw: false) 
func E_buffer_set_length_integer (i EID) EID { 
    /*(sm for mClaire/buffer_set_length @ integer= void)*/ F_buffer_set_length_integer(INT(i) )
    return EVOID} 
  
// a method for calling the printer without issuing a message
// here we assume that self_print is always defined as a function
/* {1} OPT.The go function for: apply_self_print(self:any) [] */
func F_apply_self_print_any (self *ClaireAny ) EID { 
    var Result EID 
    if (self == CNULL) /* If:2 */{ 
      PRINC("unknown")
      Result = EVOID
      /* If!2 */}  else if (self.Isa.IsIn(C_thing) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0161 *ClaireThing   = ToThing(self)
        /* noccur = 1 */
        g0161.Name.Princ()
        Result = EVOID
        /* Let-3 */} 
      /* If!2 */}  else if (C_class.Id() == self.Isa.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var g0162 *ClaireClass   = ToClass(self)
        /* noccur = 1 */
        g0162.Name.Princ()
        Result = EVOID
        /* Let-3 */} 
      /* If!2 */}  else if (C_integer.Id() == self.Isa.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var g0163 int  = ToInteger(self).Value
        /* noccur = 1 */
        F_princ_integer(g0163)
        Result = EVOID
        /* Let-3 */} 
      /* If!2 */}  else if (C_string.Id() == self.Isa.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var g0164 *ClaireString   = ToString(self)
        /* noccur = 1 */
        Result = F_CALL(C_self_print,ARGS(EID{(g0164).Id(),0}))
        /* Let-3 */} 
      } else {
      /* Let:3 */{ 
        var _Zprop *ClaireMethod   = ToMethod(F__at_property1(C_self_print,self.Isa).Id())
        /* noccur = 3 */
        if ((F_boolean_I_any(_Zprop.Id()) == CTRUE) && 
            (_Zprop.Functional.Id() != CNULL)) /* If:4 */{ 
          Result = F_funcall1(_Zprop.Functional,self)
          } else {
          h_index := ClEnv.Index /* Handle */
          h_base := ClEnv.Base
          Result = F_CALL(C_self_print,ARGS(self.ToEID()))
          if ErrorIn(Result){ 
            /* s=EID */ClEnv.Index = h_index
            ClEnv.Base = h_base
            PRINC("<unprintable:")
            Result = F_print_any(self.Isa.Id())
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC(">")
            Result = EVOID
            }
            } 
          /* If-4 */} 
        /* Let-3 */} 
      /* If-2 */} 
    return Result} 
  
// The EID go function for: apply_self_print @ any (throw: true) 
func E_apply_self_print_any (self EID) EID { 
    return /*(sm for apply_self_print @ any= EID)*/ F_apply_self_print_any(ANY(self) )} 
  
// some basic definitions
/* {1} OPT.The go function for: self_print(self:any) [] */
func F_self_print_any_Core (self *ClaireAny ) EID { 
    var Result EID 
    if (self == CNULL) /* If:2 */{ 
      PRINC("unknown")
      Result = EVOID
      } else {
      /* Let:3 */{ 
        var c *ClaireClass   = self.Isa
        /* noccur = 4 */
        /* Let:4 */{ 
          var n int  = c.Params.Length()
          /* noccur = 3 */
          if (n > 0) /* If:5 */{ 
            Result = F_print_any(c.Id())
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC("(")
            /* Let:6 */{ 
              var i int  = 1
              /* noccur = 5 */
              /* Let:7 */{ 
                var g0166 int  = n
                /* noccur = 1 */
                Result= EID{CFALSE.Id(),0}
                for (i <= g0166) /* while:8 */{ 
                  var void_try9 EID 
                  _ = void_try9
                  { 
                  /* Let:9 */{ 
                    var g0167UU *ClaireAny  
                    /* noccur = 1 */
                    var g0167UU_try016810 EID 
                    g0167UU_try016810 = F_CALL(C_get,ARGS(c.Params.At(i-1).ToEID(),self.ToEID()))
                    /* ERROR PROTECTION INSERTED (g0167UU-void_try9) */
                    if ErrorIn(g0167UU_try016810) {void_try9 = g0167UU_try016810
                    } else {
                    g0167UU = ANY(g0167UU_try016810)
                    void_try9 = F_CALL(C_print,ARGS(g0167UU.ToEID()))
                    }
                    /* Let-9 */} 
                  /* ERROR PROTECTION INSERTED (void_try9-void_try9) */
                  if ErrorIn(void_try9) {Result = void_try9
                  break
                  } else {
                  if (i < n) /* If:9 */{ 
                    PRINC(",")
                    void_try9 = EVOID
                    } else {
                    void_try9 = EID{CFALSE.Id(),0}
                    /* If-9 */} 
                  }
                  {
                  i = (i+1)
                  }
                  /* while-8 */} 
                }
                /* Let-7 */} 
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC(")")
            Result = EVOID
            }}
            } else {
            PRINC("<")
            Result = F_print_any(c.Id())
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC(">")
            Result = EVOID
            }
            /* If-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* If-2 */} 
    return Result} 
  
// The EID go function for: self_print @ any (throw: true) 
func E_self_print_any_Core (self EID) EID { 
    return /*(sm for self_print @ any= EID)*/ F_self_print_any_Core(ANY(self) )} 
  
/* {1} OPT.The go function for: self_print(self:boolean) [] */
func F_self_print_boolean_Core (self *ClaireBoolean )  { 
    // procedure body with s =  
if (self == CTRUE) /* If:2 */{ 
      PRINC("true")
      } else {
      PRINC("false")
      /* If-2 */} 
    } 
  
// The EID go function for: self_print @ boolean (throw: false) 
func E_self_print_boolean_Core (self EID) EID { 
    /*(sm for self_print @ boolean= void)*/ F_self_print_boolean_Core(ToBoolean(OBJ(self)) )
    return EVOID} 
  
/* {1} OPT.The go function for: self_print(self:function) [] */
func F_self_print_function_Core (self *ClaireFunction )  { 
    // procedure body with s =  
PRINC("#'")
    F_princ_string(F_string_I_function(self))
    PRINC("")
    } 
  
// The EID go function for: self_print @ function (throw: false) 
func E_self_print_function_Core (self EID) EID { 
    /*(sm for self_print @ function= void)*/ F_self_print_function_Core(ToFunction(OBJ(self)) )
    return EVOID} 
  
// prints the name of a restriction. If we have a close property and if a
// short-cut is possible, we use it.
/* {1} OPT.The go function for: self_print(self:restriction) [] */
func F_self_print_restriction_Core (self *ClaireRestriction ) EID { 
    var Result EID 
    if ((self.Selector.Id() == CNULL) || 
        (self.Domain.Id() == CNULL)) /* If:2 */{ 
      PRINC("<")
      Result = F_print_any(self.Id().Isa.Id())
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(">")
      Result = EVOID
      }
      } else {
      /* Let:3 */{ 
        var p *ClaireProperty   = self.Selector
        /* noccur = 2 */
        /* Let:4 */{ 
          var n int  = 0
          /* noccur = 3 */
          /* Let:5 */{ 
            var c *ClaireClass   = F_domain_I_restriction(self)
            /* noccur = 2 */
            /* For:6 */{ 
              var r *ClaireAny  
              _ = r
              for _,r = range(p.Restrictions.ValuesO())/* loop:7 */{ 
                if (F_domain_I_restriction(ToRestriction(r)).Id() == c.Id()) /* If:8 */{ 
                  n = (n+1)
                  /* If-8 */} 
                /* loop-7 */} 
              /* For-6 */} 
            p.Name.Princ()
            PRINC(" @ ")
            /* Let:6 */{ 
              var g0169UU *ClaireType  
              /* noccur = 1 */
              if (n == 1) /* If:7 */{ 
                g0169UU = ToType(c.Id())
                } else {
                g0169UU = ToType(self.Domain.Id())
                /* If-7 */} 
              Result = F_print_any(g0169UU.Id())
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC("")
            Result = EVOID
            }
            /* Let-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* If-2 */} 
    return Result} 
  
// The EID go function for: self_print @ restriction (throw: true) 
func E_self_print_restriction_Core (self EID) EID { 
    return /*(sm for self_print @ restriction= EID)*/ F_self_print_restriction_Core(ToRestriction(OBJ(self)) )} 
  
// we are too far
/* {1} OPT.The go function for: print(x:any) [] */
func F_print_any (x *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var _Zl int  = C_pretty.Index
      /* noccur = 2 */
      if ((C_pretty.Pbreak == CTRUE) && 
          (C_pretty.Pprint == CTRUE)) /* If:3 */{ 
        /* Let:4 */{ 
          var b_index int  = F_buffer_length_void()
          /* noccur = 2 */
          /* Let:5 */{ 
            var missed *ClaireBoolean   = CFALSE
            /* noccur = 2 */
            if (F_short_enough_integer((b_index+10)) != CTRUE) /* If:6 */{ 
              C_pretty.Pprint = CFALSE
              C_pretty.Pbreak = CFALSE
              Result = F_CALL(C_print,ARGS(x.ToEID()))
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              C_pretty.Pbreak = CTRUE
              /* update:7 */{ 
                var va_arg1 *PrettyPrinter  
                var va_arg2 *ClaireBoolean  
                va_arg1 = C_pretty
                va_arg2 = CTRUE
                /* ---------- now we compile update mClaire/pprint(va_arg1) := va_arg2 ------- */
                va_arg1.Pprint = va_arg2
                Result = EID{va_arg2.Id(),0}
                /* update-7 */} 
              }
              } else {
              h_index := ClEnv.Index /* Handle */
              h_base := ClEnv.Base
              C_pretty.Pbreak = CFALSE
              Result = F_apply_self_print_any(x)
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              /* update:7 */{ 
                var va_arg1 *PrettyPrinter  
                var va_arg2 *ClaireBoolean  
                va_arg1 = C_pretty
                va_arg2 = CTRUE
                /* ---------- now we compile update mClaire/pbreak(va_arg1) := va_arg2 ------- */
                va_arg1.Pbreak = va_arg2
                Result = EID{va_arg2.Id(),0}
                /* update-7 */} 
              }
              if ErrorIn(Result) && ToType(C_much_too_far.Id()).Contains(ANY(Result)) == CTRUE { 
                /* s=EID */ClEnv.Index = h_index
                ClEnv.Base = h_base
                missed = CTRUE
                Result = EID{missed.Id(),0}
                } 
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              if (missed == CTRUE) /* If:7 */{ 
                C_pretty.Pprint = CTRUE
                C_pretty.Pbreak = CTRUE
                F_buffer_set_length_integer(b_index)
                C_pretty.Index = _Zl
                Result = F_apply_self_print_any(x)
                } else {
                Result = EID{CFALSE.Id(),0}
                /* If-7 */} 
              }
              /* If-6 */} 
            /* Let-5 */} 
          /* Let-4 */} 
        } else {
        Result = F_apply_self_print_any(x)
        /* If-3 */} 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      C_pretty.Index = _Zl
      Result = EID{CNULL,0}
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: print @ any (throw: true) 
func E_print_any (x EID) EID { 
    return /*(sm for print @ any= EID)*/ F_print_any(ANY(x) )} 
  
// short_enough = we expect that what we want to print is short enough (more that 10 chars to the width)
/* {1} OPT.The go function for: short_enough(self:integer) [] */
func F_short_enough_integer (self int) *ClaireBoolean  { 
    // use function body compiling 
return  F__inf_integer(self,C_pretty.Width)
    } 
  
// The EID go function for: short_enough @ integer (throw: false) 
func E_short_enough_integer (self EID) EID { 
    return EID{/*(sm for short_enough @ integer= boolean)*/ F_short_enough_integer(INT(self) ).Id(),0}} 
  
// *********************************************************************
// *   Part 2: Methods for CLAIRE objects                              *
// *********************************************************************
// the instantiation body is a sequence of words from which the initialization
// of the object must be built.
//  copied_def = object (for object) + float (for float) + integer (for all)
//               + NULL for objects
// v3.2.12: use a condition that is coherent with ClReflect.cl : a slot defaut value must be placed
// unless it is a copied_def
// CLAIRE4 (fix OLD bug): lp is list of p such that o.p := unknown is requested !
/* {1} OPT.The go function for: new_defaults(self:object,lp:list) [] */
func F_Core_new_defaults_object (self *ClaireObject ,lp *ClaireList ) EID { 
    var Result EID 
    /* For:2 */{ 
      var s *ClaireAny  
      _ = s
      Result= EID{CFALSE.Id(),0}
      for _,s = range(self.Isa.Slots.ValuesO())/* loop:3 */{ 
        var void_try4 EID 
        _ = void_try4
        /* Let:4 */{ 
          var p *ClaireProperty   = ToRestriction(s).Selector
          /* noccur = 5 */
          /* Let:5 */{ 
            var s2 *ClaireClass   = ToSlot(s).Srange
            /* noccur = 2 */
            /* Let:6 */{ 
              var d *ClaireAny   = ToSlot(s).Default
              /* noccur = 6 */
              if (d != CNULL) /* If:7 */{ 
                /* Let:8 */{ 
                  var v *ClaireAny   = self.SlotGet(ToSlot(s).Index,ToSlot(s).Srange)
                  /* noccur = 3 */
                  if ((v == CNULL) && 
                      (lp.Memq(p.Id()) == CTRUE)) /* If:9 */{ 
                    void_try4 = EID{CNIL.Id(),0}
                    /* If!9 */}  else if ((v == CNULL) && 
                      ((s2.Id() != C_object.Id()) && 
                          ((C_integer.Id() != d.Isa.Id()) && 
                            (s2.Id() != C_float.Id())))) /* If:9 */{ 
                    void_try4 = F_update_property(p,
                      self,
                      ToSlot(s).Index,
                      ToSlot(s).Srange,
                      d)
                    /* If!9 */}  else if (Equal(d,v) == CTRUE) /* If:9 */{ 
                    if (p.Multivalued_ask == CTRUE) /* If:10 */{ 
                      /* For:11 */{ 
                        var y *ClaireAny  
                        _ = y
                        void_try4= EID{CFALSE.Id(),0}
                        var y_support *ClaireList  
                        var y_support_try017212 EID 
                        y_support_try017212 = F_enumerate_any(d)
                        /* ERROR PROTECTION INSERTED (y_support-void_try4) */
                        if ErrorIn(y_support_try017212) {void_try4 = y_support_try017212
                        } else {
                        y_support = ToList(OBJ(y_support_try017212))
                        y_len := y_support.Length()
                        for i_it := 0; i_it < y_len; i_it++ { 
                          y = y_support.At(i_it)
                          var void_try13 EID 
                          _ = void_try13
                          void_try13 = F_update_plus_relation(ToRelation(p.Id()),self.Id(),y)
                          /* ERROR PROTECTION INSERTED (void_try13-void_try4) */
                          if ErrorIn(void_try13) {void_try4 = void_try13
                          void_try4 = void_try13
                          break
                          } else {
                          }}
                          /* loop-12 */} 
                        /* For-11 */} 
                      } else {
                      void_try4 = F_update_plus_relation(ToRelation(p.Id()),self.Id(),d)
                      /* If-10 */} 
                    } else {
                    void_try4 = EID{CFALSE.Id(),0}
                    /* If-9 */} 
                  /* Let-8 */} 
                } else {
                void_try4 = EID{CFALSE.Id(),0}
                /* If-7 */} 
              /* Let-6 */} 
            /* Let-5 */} 
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (void_try4-Result) */
        if ErrorIn(void_try4) {Result = void_try4
        Result = void_try4
        break
        } else {
        }
        /* loop-3 */} 
      /* For-2 */} 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    /* Let:2 */{ 
      var m *ClaireAny   = F__at_property1(C_close,self.Id().Isa).Id()
      /* noccur = 2 */
      if (C_method.Id() == m.Isa.Id()) /* If:3 */{ 
        /* Let:4 */{ 
          var g0170 *ClaireMethod   = ToMethod(m)
          /* noccur = 1 */
          Result = F_funcall_method1(g0170,self.Id())
          /* Let-4 */} 
        } else {
        Result = EID{self.Id(),0}
        /* If-3 */} 
      /* Let-2 */} 
    }
    return Result} 
  
// The EID go function for: new_defaults @ object (throw: true) 
func E_Core_new_defaults_object (self EID,lp EID) EID { 
    return /*(sm for new_defaults @ object= EID)*/ F_Core_new_defaults_object(ToObject(OBJ(self)),ToList(OBJ(lp)) )} 
  
// v3.0.41  obviously
//-------------------------- ENTITY   --------------------------------------
/* {1} OPT.The go function for: not(self:any) [] */
func F_not_any (self *ClaireAny ) *ClaireBoolean  { 
    // use function body compiling 
if (self == CTRUE.Id()) /* body If:2 */{ 
      return  CFALSE
      }  else if (self == CFALSE.Id()) /* body If:2 */{ 
      return  CTRUE
      }  else if (F_boolean_I_any(self).Id() != CTRUE.Id()) /* body If:2 */{ 
      return  CTRUE
      } else {
      return  CFALSE
      /* body If-2 */} 
    } 
  
// The EID go function for: not @ any (throw: false) 
func E_not_any (self EID) EID { 
    return EID{/*(sm for not @ any= boolean)*/ F_not_any(ANY(self) ).Id(),0}} 
  
/* {1} OPT.The go function for: !=(self:any,x:any) [] */
func F__I_equal_any (self *ClaireAny ,x *ClaireAny ) *ClaireBoolean  { 
    // use function body compiling 
if (Equal(self,x) == CTRUE) /* body If:2 */{ 
      return  CFALSE
      } else {
      return  CTRUE
      /* body If-2 */} 
    } 
  
// The EID go function for: != @ any (throw: false) 
func E__I_equal_any (self EID,x EID) EID { 
    return EID{/*(sm for != @ any= boolean)*/ F__I_equal_any(ANY(self),ANY(x) ).Id(),0}} 
  
// gives the type of any object. This is open_coded.
/* {1} OPT.The go function for: owner(self:any) [] */
func F_owner_any (self *ClaireAny ) *ClaireClass  { 
    // use function body compiling 
return  self.Isa
    } 
  
// The EID go function for: owner @ any (throw: false) 
func E_owner_any (self EID) EID { 
    return EID{/*(sm for owner @ any= class)*/ F_owner_any(ANY(self) ).Id(),0}} 
  
// some useful methods
/* {1} OPT.The go function for: known?(self:any) [] */
func F_known_ask_any (self *ClaireAny ) *ClaireBoolean  { 
    // use function body compiling 
return  F__I_equal_any(CNULL,self)
    } 
  
// The EID go function for: known? @ any (throw: false) 
func E_known_ask_any (self EID) EID { 
    return EID{/*(sm for known? @ any= boolean)*/ F_known_ask_any(ANY(self) ).Id(),0}} 
  
/* {1} OPT.The go function for: unknown?(self:any) [] */
func F_unknown_ask_any (self *ClaireAny ) *ClaireBoolean  { 
    // use function body compiling 
return  Equal(CNULL,self)
    } 
  
// The EID go function for: unknown? @ any (throw: false) 
func E_unknown_ask_any (self EID) EID { 
    return EID{/*(sm for unknown? @ any= boolean)*/ F_unknown_ask_any(ANY(self) ).Id(),0}} 
  
// needed by the compiled code for casts (inserted by ocontrol for dynamic type checks)
// Claire 4: TODO - add a second order type
/* {1} OPT.The go function for: check_in(self:any,y:type) [] */
func F_check_in_any (self *ClaireAny ,y *ClaireType ) EID { 
    var Result EID 
    if (y.Contains(self) == CTRUE) /* If:2 */{ 
      Result = self.ToEID()
      } else {
      Result = ToException(C_general_error.Make(MakeString("[124] the value ~S does not belong to the range ~S").Id(),MakeConstantList(self,y.Id()).Id())).Close()
      /* If-2 */} 
    return Result} 
  
// The EID go function for: check_in @ any (throw: true) 
func E_check_in_any (self EID,y EID) EID { 
    return /*(sm for check_in @ any= EID)*/ F_check_in_any(ANY(self),ToType(OBJ(y)) )} 
  
// used to cast dynamically a non-mutable bag to a typed mutable bag
// claire 4 : aplied to list and sets (bags)
/* {1} OPT.The go function for: check_in(self:bag,c:class,y:type) [] */
func F_check_in_bag (self *ClaireBag ,c *ClaireClass ,y *ClaireType ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var OK *ClaireBoolean  
      /* noccur = 1 */
      if (self.Isa.IsIn(C_list) == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var g0173 *ClaireList   = ToList(self.Id())
          /* noccur = 1 */
          /* Let:5 */{ 
            var g0175UU *ClaireAny  
            /* noccur = 1 */
            /* For:6 */{ 
              var z *ClaireAny  
              _ = z
              g0175UU= CFALSE.Id()
              var z_support *ClaireList  
              z_support = g0173
              z_len := z_support.Length()
              for i_it := 0; i_it < z_len; i_it++ { 
                z = z_support.At(i_it)
                if (y.Contains(z) != CTRUE) /* If:8 */{ 
                   /*v = g0175UU, s =any*/
g0175UU = CTRUE.Id()
                  break
                  /* If-8 */} 
                /* loop-7 */} 
              /* For-6 */} 
            OK = F_not_any(g0175UU)
            /* Let-5 */} 
          /* Let-4 */} 
        } else {
        /* Let:4 */{ 
          var g0176UU *ClaireAny  
          /* noccur = 1 */
          /* For:5 */{ 
            var z *ClaireAny  
            _ = z
            g0176UU= CFALSE.Id()
            var z_support *ClaireSet  
            z_support = ToSet(self.Id())
            for _,z = range(z_support.Values)/* loop2:6 */{ 
              if (y.Contains(z) != CTRUE) /* If:7 */{ 
                 /*v = g0176UU, s =any*/
g0176UU = CTRUE.Id()
                break
                /* If-7 */} 
              /* loop-6 */} 
            /* For-5 */} 
          OK = F_not_any(g0176UU)
          /* Let-4 */} 
        /* If-3 */} 
      if (OK == CTRUE) /* If:3 */{ 
        Result = EID{self.Cast_I(y).Id(),0}
        } else {
        Result = ToException(C_general_error.Make(MakeString("[124] the value ~S does not belong to subtype[~S]").Id(),MakeConstantList(self.Id(),y.Id()).Id())).Close()
        /* If-3 */} 
      /* Let-2 */} 
    return RangeCheck(ToType(C_bag.Id()),Result)} 
  
// The EID go function for: check_in @ bag (throw: true) 
func E_check_in_bag (self EID,c EID,y EID) EID { 
    return /*(sm for check_in @ bag= EID)*/ F_check_in_bag(ToBag(OBJ(self)),ToClass(OBJ(c)),ToType(OBJ(y)) )} 
  
// new in v3.00.48
/* {1} OPT.The go function for: <(self:any,x:any) [] */
func F__inf_any (self *ClaireAny ,x *ClaireAny ) *ClaireBoolean  { 
    // procedure body with s =  
var Result *ClaireBoolean  
    if (Equal(self,x) == CTRUE) /* If:2 */{ 
      Result = CFALSE
      } else {
      Result = ToBoolean(OBJ(F_CALL(ToProperty(C__inf_equal.Id()),ARGS(self.ToEID(),x.ToEID()))))
      /* If-2 */} 
    return Result} 
  
// The EID go function for: < @ any (throw: false) 
func E__inf_any (self EID,x EID) EID { 
    return EID{/*(sm for < @ any= boolean)*/ F__inf_any(ANY(self),ANY(x) ).Id(),0}} 
  
/* {1} OPT.The go function for: >(self:any,x:any) [] */
func F__sup_any (self *ClaireAny ,x *ClaireAny ) *ClaireBoolean  { 
    // procedure body with s =  
var Result *ClaireBoolean  
    if (Equal(self,x) == CTRUE) /* If:2 */{ 
      Result = CFALSE
      } else {
      Result = ToBoolean(OBJ(F_CALL(ToProperty(C__inf_equal.Id()),ARGS(x.ToEID(),self.ToEID()))))
      /* If-2 */} 
    return Result} 
  
// The EID go function for: > @ any (throw: false) 
func E__sup_any (self EID,x EID) EID { 
    return EID{/*(sm for > @ any= boolean)*/ F__sup_any(ANY(self),ANY(x) ).Id(),0}} 
  
// >= is defined as a macro in file.cl
// unsafe is a pragma : it tells the compiler not to catch a possible error 
// to add in the documentation
/* {1} OPT.The go function for: unsafe(x:any) [] */
func F_unsafe_any (x *ClaireAny ) *ClaireAny  { 
    // use function body compiling 
return  x
    } 
  
// The EID go function for: unsafe @ any (throw: false) 
func E_unsafe_any (x EID) EID { 
    return /*(sm for unsafe @ any= any)*/ F_unsafe_any(ANY(x) ).ToEID()} 
  
// ----------------------- CLASS ---------------------------------------------
// classes have a "open" status with three states
// forward : -2          - use open as a marker for forward definition
// close : -1           ????
// abstract :  open = 0  - cannot create an instance
// final : open = 1  - no more subclasses or instance
// default : open : 2 - default for ephemeral - no instance
// open :  open = 3  - default for things (named things can be queried)
// declares a class as ephemeral: the member set is not maintained
// v3.2.14 recusively applies to subclasses
/* {1} OPT.The go function for: ephemeral(self:class) [] */
func F_ephemeral_class (self *ClaireClass ) EID { 
    var Result EID 
    /* For:2 */{ 
      var c *ClaireAny  
      _ = c
      Result= EID{CFALSE.Id(),0}
      for _,c = range(self.Descendents.Values)/* loop:3 */{ 
        var void_try4 EID 
        _ = void_try4
        if (ToClass(c).Instances.Length() != 0) /* If:4 */{ 
          void_try4 = ToException(C_general_error.Make(MakeString("[187] cannot declare ~S as ephemeral because of ~S has instances").Id(),MakeConstantList(self.Id(),c).Id())).Close()
          } else {
          /* update:5 */{ 
            var va_arg1 *ClaireClass  
            var va_arg2 int 
            va_arg1 = ToClass(c)
            va_arg2 = ClEnv.Default
            /* ---------- now we compile update open(va_arg1) := va_arg2 ------- */
            va_arg1.Open = va_arg2
            void_try4 = EID{C__INT,IVAL(va_arg2)}
            /* update-5 */} 
          /* If-4 */} 
        /* ERROR PROTECTION INSERTED (void_try4-Result) */
        if ErrorIn(void_try4) {Result = void_try4
        Result = void_try4
        break
        } else {
        }
        /* loop-3 */} 
      /* For-2 */} 
    return Result} 
  
// The EID go function for: ephemeral @ class (throw: true) 
func E_ephemeral_class (self EID) EID { 
    return /*(sm for ephemeral @ class= EID)*/ F_ephemeral_class(ToClass(OBJ(self)) )} 
  
// claire4 : opposite of ephemeral(c)
/* {1} OPT.The go function for: instanced(c:class) [] */
func F_instanced_class (c *ClaireClass ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var n int  = c.Open
      /* noccur = 2 */
      if ((n == ClEnv.Default) || 
          (n == ClEnv.Open)) /* If:3 */{ 
        /* update:4 */{ 
          var va_arg1 *ClaireClass  
          var va_arg2 int 
          va_arg1 = c
          va_arg2 = ClEnv.Open
          /* ---------- now we compile update open(va_arg1) := va_arg2 ------- */
          va_arg1.Open = va_arg2
          Result = EID{C__INT,IVAL(va_arg2)}
          /* update-4 */} 
        } else {
        Result = ToException(C_general_error.Make(MakeString("[125] abstract classes cannot be instanced").Id(),CNIL.Id())).Close()
        /* If-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: instanced @ class (throw: true) 
func E_instanced_class (c EID) EID { 
    return /*(sm for instanced @ class= EID)*/ F_instanced_class(ToClass(OBJ(c)) )} 
  
// declares a class as an abtract class (without instances)
/* {1} OPT.The go function for: abstract(c:class) [] */
func F_abstract_class (c *ClaireClass ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var n int  = c.Open
      /* noccur = 1 */
      if (Equal(MakeInteger(n).Id(),ANY(F_CALL(C_Core_closed,ARGS(EID{ClEnv.Id(),0})))) == CTRUE) /* If:3 */{ 
        Result = ToException(C_general_error.Make(MakeString("[125] closed classes cannot be abstract").Id(),CNIL.Id())).Close()
        /* If!3 */}  else if (c.Instances.Length() != 0) /* If:3 */{ 
        Result = ToException(C_general_error.Make(MakeString("[125] instanced classes cannot be abstract").Id(),CNIL.Id())).Close()
        } else {
        /* update:4 */{ 
          var va_arg1 *ClaireClass  
          var va_arg2 int 
          va_arg1 = c
          va_arg2 = ClEnv.ABSTRACT
          /* ---------- now we compile update open(va_arg1) := va_arg2 ------- */
          va_arg1.Open = va_arg2
          Result = EID{C__INT,IVAL(va_arg2)}
          /* update-4 */} 
        /* If-3 */} 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = EID{c.Id(),0}
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: abstract @ class (throw: true) 
func E_abstract_class (c EID) EID { 
    return /*(sm for abstract @ class= EID)*/ F_abstract_class(ToClass(OBJ(c)) )} 
  
// declares a class with no subclasses (apply to things)
/* {1} OPT.The go function for: final(c:class) [] */
func F_final_class (c *ClaireClass ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var n int  = c.Open
      /* noccur = 1 */
      if (n == ClEnv.Default) /* If:3 */{ 
        Result = ToException(C_general_error.Make(MakeString("[125] ephemetral classes cannot be final").Id(),CNIL.Id())).Close()
        /* If!3 */}  else if (Equal(ANY(F_CALL(C_length,ARGS(EID{c.Subclass.Id(),0}))),MakeInteger(0).Id()) != CTRUE) /* If:3 */{ 
        Result = ToException(C_general_error.Make(MakeString("[125] a class with wsubclasses cannot be final").Id(),CNIL.Id())).Close()
        } else {
        /* update:4 */{ 
          var va_arg1 *ClaireClass  
          var va_arg2 int 
          va_arg1 = c
          va_arg2 = ClEnv.Final
          /* ---------- now we compile update open(va_arg1) := va_arg2 ------- */
          va_arg1.Open = va_arg2
          Result = EID{C__INT,IVAL(va_arg2)}
          /* update-4 */} 
        /* If-3 */} 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = EID{c.Id(),0}
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: final @ class (throw: true) 
func E_final_class (c EID) EID { 
    return /*(sm for final @ class= EID)*/ F_final_class(ToClass(OBJ(c)) )} 
  
// instantiation with and without a name
// new! is a method-less property that is managed by the compiler
/* {1} OPT.The go function for: new(self:class) [] */
func F_new_class1 (self *ClaireClass ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var o *ClaireObject  
      /* noccur = 1 */
      var o_try01773 EID 
      if (self.Open <= 0) /* If:3 */{ 
        o_try01773 = ToException(C_general_error.Make(MakeString("[105] cannot instantiate ~S").Id(),MakeConstantList(self.Id()).Id())).Close()
        } else {
        o_try01773 = EID{CFALSE.Id(),0}
        /* If-3 */} 
      /* ERROR PROTECTION INSERTED (o_try01773-o_try01773) */
      if !ErrorIn(o_try01773) {
      o_try01773 = EID{F_new_object_class(self).Id(),0}
      }
      /* ERROR PROTECTION INSERTED (o-Result) */
      if ErrorIn(o_try01773) {Result = o_try01773
      } else {
      o = ToObject(OBJ(o_try01773))
      Result = F_Core_new_defaults_object(o,CNIL)
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: new @ list<type_expression>(class) (throw: true) 
func E_new_class1 (self EID) EID { 
    return /*(sm for new @ list<type_expression>(class)= EID)*/ F_new_class1(ToClass(OBJ(self)) )} 
  
/* {1} OPT.The go function for: new_class1_type */
func F_new_class1_type (self *ClaireType ) EID { 
    /* eid body: object glb member(self) */
    var Result EID 
    Result = F_CALL(ToProperty(C_glb.Id()),ARGS(EID{C_object.Id(),0},F_CALL(C_member,ARGS(EID{self.Id(),0}))))
    return Result} 
  
  
// The dual EID go function for: "new_class1_type" 
func E_new_class1_type (self EID) EID { 
    return F_new_class1_type(ToType(OBJ(self)))} 
  
// v3.2.26
/* {1} OPT.The go function for: new(self:class,%nom:symbol) [] */
func F_new_class2 (self *ClaireClass ,_Znom *ClaireSymbol ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var o *ClaireThing  
      /* noccur = 1 */
      var o_try01783 EID 
      if (self.Open <= 0) /* If:3 */{ 
        o_try01783 = ToException(C_general_error.Make(MakeString("[105] cannot instantiate ~S").Id(),MakeConstantList(self.Id()).Id())).Close()
        } else {
        o_try01783 = EID{CFALSE.Id(),0}
        /* If-3 */} 
      /* ERROR PROTECTION INSERTED (o_try01783-o_try01783) */
      if !ErrorIn(o_try01783) {
      o_try01783 = F_new_thing_class(self,_Znom)
      }
      /* ERROR PROTECTION INSERTED (o-Result) */
      if ErrorIn(o_try01783) {Result = o_try01783
      } else {
      o = ToThing(OBJ(o_try01783))
      Result = F_Core_new_defaults_object(ToObject(o.Id()),CNIL)
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: new @ list<type_expression>(class, symbol) (throw: true) 
func E_new_class2 (self EID,_Znom EID) EID { 
    return /*(sm for new @ list<type_expression>(class, symbol)= EID)*/ F_new_class2(ToClass(OBJ(self)),ToSymbol(OBJ(_Znom)) )} 
  
/* {1} OPT.The go function for: new_class2_type */
func F_new_class2_type (self *ClaireType ,_Znom *ClaireType ) EID { 
    /* eid body: thing glb member(self) */
    var Result EID 
    Result = F_CALL(ToProperty(C_glb.Id()),ARGS(EID{C_thing.Id(),0},F_CALL(C_member,ARGS(EID{self.Id(),0}))))
    return Result} 
  
  
// The dual EID go function for: "new_class2_type" 
func E_new_class2_type (self EID,_Znom EID) EID { 
    return F_new_class2_type(ToType(OBJ(self)),ToType(OBJ(_Znom)))} 
  
// v3.2.26
// internal version => moved to Kernel
// mClaire/new!(self:class) :  type[object glb member(self)]
//  -> function!(new_object_class)
// mClaire/new!(self:class,%nom:symbol) :  type[thing glb member(self)]
//   -> function!(new_thing_class)
// the smallest super_set of two sets
// there is always any, so it always returns a class
/* {1} OPT.The go function for: meet(self:class,ens:class) [] */
func F_meet_class (self *ClaireClass ,ens *ClaireClass ) *ClaireClass  { 
    // procedure body with s =  
var Result *ClaireClass  
    /* Let:2 */{ 
      var l1 *ClaireList   = self.Ancestors
      /* noccur = 4 */
      /* Let:3 */{ 
        var l2 *ClaireList   = ens.Ancestors
        /* noccur = 3 */
        /* Let:4 */{ 
          var m int 
          /* noccur = 5 */
          if (l1.Length() < l2.Length()) /* If:5 */{ 
            m = l1.Length()
            } else {
            m = l2.Length()
            /* If-5 */} 
          for (Equal(l1.ValuesO()[m-1],l2.ValuesO()[m-1]) != CTRUE) /* while:5 */{ 
            m = (m-1)
            /* while-5 */} 
          Result = ToClass(l1.ValuesO()[m-1])
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: meet @ class (throw: false) 
func E_meet_class (self EID,ens EID) EID { 
    return EID{/*(sm for meet @ class= class)*/ F_meet_class(ToClass(OBJ(self)),ToClass(OBJ(ens)) ).Id(),0}} 
  
// fast inclusion method for lattice_sets (lattice order). The argument is
// either a lattice_set or {}
/* {1} OPT.The go function for: inherit?(self:class,ens:class) [] */
func F_inherit_ask_class (self *ClaireClass ,ens *ClaireClass ) *ClaireBoolean  { 
    // procedure body with s =  
var Result *ClaireBoolean  
    /* Let:2 */{ 
      var l *ClaireList   = self.Ancestors
      /* noccur = 2 */
      /* Let:3 */{ 
        var n int  = ens.Ancestors.Length()
        /* noccur = 2 */
        Result = MakeBoolean((n <= l.Length()) && (l.ValuesO()[n-1] == ens.Id()))
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: inherit? @ class (throw: false) 
func E_inherit_ask_class (self EID,ens EID) EID { 
    return EID{/*(sm for inherit? @ class= boolean)*/ F_inherit_ask_class(ToClass(OBJ(self)),ToClass(OBJ(ens)) ).Id(),0}} 
  
//------------- PROPERTY ---------------------------------------------------
// the two methods to access open(r)
// an abstract property is extensible and can receive new restrictions
/* {1} OPT.The go function for: abstract(p:property) [] */
func F_abstract_property (p *ClaireProperty ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var n int  = p.Open
      /* noccur = 1 */
      if (n < 2) /* If:3 */{ 
        Result = ToException(C_general_error.Make(MakeString("[127] ~S can no longer become abstract").Id(),MakeConstantList(p.Id()).Id())).Close()
        } else {
        /* update:4 */{ 
          var va_arg1 *ClaireRelation  
          var va_arg2 int 
          va_arg1 = ToRelation(p.Id())
          va_arg2 = 3
          /* ---------- now we compile update open(va_arg1) := va_arg2 ------- */
          va_arg1.Open = va_arg2
          Result = EID{C__INT,IVAL(va_arg2)}
          /* update-4 */} 
        /* If-3 */} 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = EID{p.Id(),0}
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: abstract @ property (throw: true) 
func E_abstract_property (p EID) EID { 
    return /*(sm for abstract @ property= EID)*/ F_abstract_property(ToProperty(OBJ(p)) )} 
  
// a final property is completely defined and cannot receive a new restriction
// v3.2.04: the new value 4 will be used to represent (compiled but open)
/* {1} OPT.The go function for: final(r:relation) [] */
func F_final_relation (r *ClaireRelation )  { 
    // procedure body with s =  
if (r.Isa.IsIn(C_property) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0179 *ClaireProperty   = ToProperty(r.Id())
        /* noccur = 7 */
        if (g0179.Open <= 2) /* If:4 */{ 
          g0179.Open = 1
          /* update:5 */{ 
            var va_arg1 *ClaireRelation  
            var va_arg2 *ClaireType  
            va_arg1 = ToRelation(g0179.Id())
            /* Let:6 */{ 
              var g0180UU *ClaireList  
              /* noccur = 1 */
              /* Iteration:7 */{ 
                var v_list7 *ClaireList  
                var x *ClaireAny  
                var v_local7 *ClaireAny  
                v_list7 = g0179.Restrictions
                g0180UU = CreateList(ToType(CEMPTY.Id()),v_list7.Length())
                for CLcount := 0; CLcount < v_list7.Length(); CLcount++{ 
                  x = v_list7.At(CLcount)
                  v_local7 = ToRestriction(x).Domain.ValuesO()[1-1]
                  g0180UU.PutAt(CLcount,v_local7)
                  } 
                /* Iteration-7 */} 
              va_arg2 = F_Uall_list(g0180UU)
              /* Let-6 */} 
            /* ---------- now we compile update domain(va_arg1) := va_arg2 ------- */
            va_arg1.Domain = va_arg2
            /* update-5 */} 
          /* update:5 */{ 
            var va_arg1 *ClaireRelation  
            var va_arg2 *ClaireType  
            va_arg1 = ToRelation(g0179.Id())
            /* Let:6 */{ 
              var g0181UU *ClaireList  
              /* noccur = 1 */
              /* Iteration:7 */{ 
                var v_list7 *ClaireList  
                var x *ClaireAny  
                var v_local7 *ClaireAny  
                v_list7 = g0179.Restrictions
                g0181UU = CreateList(ToType(CEMPTY.Id()),v_list7.Length())
                for CLcount := 0; CLcount < v_list7.Length(); CLcount++{ 
                  x = v_list7.At(CLcount)
                  v_local7 = ToRestriction(x).Range.Id()
                  g0181UU.PutAt(CLcount,v_local7)
                  } 
                /* Iteration-7 */} 
              va_arg2 = F_Uall_list(g0181UU)
              /* Let-6 */} 
            /* ---------- now we compile update range(va_arg1) := va_arg2 ------- */
            va_arg1.Range = va_arg2
            /* update-5 */} 
          /* If-4 */} 
        /* Let-3 */} 
      /* If-2 */} 
    } 
  
// The EID go function for: final @ relation (throw: false) 
func E_final_relation (r EID) EID { 
    /*(sm for final @ relation= void)*/ F_final_relation(ToRelation(OBJ(r)) )
    return EVOID} 
  
//------------- MODULES   --------------------------------------------------
// book-keeping for a module : based on parts/part_of hierarchy
// propagates uses declaration + register a new associated namespace.
/* {1} OPT.The go function for: close(self:module) [] */
func F_close_module (self *ClaireModule ) *ClaireModule  { 
    // use function body compiling 
if (self.Id() != C_claire.Id()) /* If:2 */{ 
      if (self.PartOf.Id() != CNULL) /* If:3 */{ 
        /* Let:4 */{ 
          var sup *ClaireModule   = self.PartOf
          /* noccur = 4 */
          sup.Parts = sup.Parts.AddFast(self.Id())
          /* For:5 */{ 
            var x *ClaireAny  
            _ = x
            var x_support *ClaireList  
            x_support = sup.Uses
            x_len := x_support.Length()
            for i_it := 0; i_it < x_len; i_it++ { 
              x = x_support.At(i_it)
              if ((ToBoolean(self.Uses.Contain_ask(x).Id()) != CTRUE) && 
                  (x.Isa.IsIn(C_module) == CTRUE)) /* If:7 */{ 
                self.Uses = self.Uses.AddFast(x)
                /* If-7 */} 
              /* loop-6 */} 
            /* For-5 */} 
          /* Let-4 */} 
        /* If-3 */} 
      /* If-2 */} 
    self.Namespace()
    return  self
    } 
  
// The EID go function for: close @ module (throw: false) 
func E_close_module (self EID) EID { 
    return EID{/*(sm for close @ module= module)*/ F_close_module(ToModule(OBJ(self)) ).Id(),0}} 
  
// note: dynamic modules are no longer supported
/* {1} OPT.The go function for: get_symbol(self:string) [] */
func F_get_symbol_string (self *ClaireString ) *ClaireAny  { 
    // use function body compiling 
return  F_get_symbol_module(C_claire,self)
    } 
  
// The EID go function for: get_symbol @ string (throw: false) 
func E_get_symbol_string (self EID) EID { 
    return /*(sm for get_symbol @ string= any)*/ F_get_symbol_string(ToString(OBJ(self)) ).ToEID()} 
  
// *********************************************************************
// *   Part 3: System Methods                                          *
// *********************************************************************
// all these methods will be open-coded by the compiler
//get_stack(self:integer) : any -> get_stack(self)
//put_stack(self:integer,x:any) : any -> put_stack(self, x)
//push!(self:meta_system,x:any) : void -> push!(self, x)
//stat() : void -> function!(claire_stat)
/* {1} OPT.The go function for: gensym(self:void) [] */
func F_gensym_void () *ClaireSymbol  { 
    // use function body compiling 
return  F_gensym_string(MakeString("g"))
    } 
  
// The EID go function for: gensym @ void (throw: false) 
func E_gensym_void (self EID) EID { 
    return EID{/*(sm for gensym @ void= symbol)*/ F_gensym_void( ).Id(),0}} 
  
// world management
/* {1} OPT.The go function for: store(l:list,n:integer,y:any) [] */
func F_store_list4 (l *ClaireList ,n int,y *ClaireAny ) *ClaireAny  { 
    // use function body compiling 
return  F_store_list(l,n,y,CTRUE)
    } 
  
// The EID go function for: store @ list<type_expression>(list, integer, any) (throw: false) 
func E_store_list4 (l EID,n EID,y EID) EID { 
    return /*(sm for store @ list<type_expression>(list, integer, any)= any)*/ F_store_list4(ToList(OBJ(l)),INT(n),ANY(y) ).ToEID()} 
  
/* {1} OPT.The go function for: store(l:array,n:integer,y:any) [] */
func F_store_array1 (l *ClaireList ,n int,y *ClaireAny ) *ClaireAny  { 
    // use function body compiling 
return  F_store_list(ToList(l.Id()),n,y,CTRUE)
    } 
  
// The EID go function for: store @ list<type_expression>(array, integer, any) (throw: false) 
func E_store_array1 (l EID,n EID,y EID) EID { 
    return /*(sm for store @ list<type_expression>(array, integer, any)= any)*/ F_store_array1(ToArray(OBJ(l)),INT(n),ANY(y) ).ToEID()} 
  
/* {1} OPT.The go function for: commit(n:integer) [] */
func F_commit_integer (n int)  { 
    // procedure body with s =  
for (n < F_world_number()) /* while:2 */{ 
      F_world_remove()
      /* while-2 */} 
    } 
  
// The EID go function for: commit @ integer (throw: false) 
func E_commit_integer (n EID) EID { 
    /*(sm for commit @ integer= void)*/ F_commit_integer(INT(n) )
    return EVOID} 
  
/* {1} OPT.The go function for: backtrack(n:integer) [] */
func F_backtrack_integer (n int)  { 
    // procedure body with s =  
for (n < F_world_number()) /* while:2 */{ 
      F_world_pop()
      /* while-2 */} 
    } 
  
// The EID go function for: backtrack @ integer (throw: false) 
func E_backtrack_integer (n EID) EID { 
    /*(sm for backtrack @ integer= void)*/ F_backtrack_integer(INT(n) )
    return EVOID} 
  
// allows to change the storage class
/* {1} OPT.The go function for: store(l:listargs) [] */
func F_store_listargs (l *ClaireList ) *ClaireAny  { 
    // procedure body with s =  
var Result *ClaireAny  
    /* For:2 */{ 
      var r *ClaireAny  
      _ = r
      Result= CFALSE.Id()
      var r_support *ClaireList  
      r_support = ToList(l.Id())
      r_len := r_support.Length()
      for i_it := 0; i_it < r_len; i_it++ { 
        r = r_support.At(i_it)
        if (r.Isa.IsIn(C_relation) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var g0184 *ClaireRelation   = ToRelation(r)
            /* noccur = 2 */
            g0184.Store_ask = CTRUE
            /* Let-5 */} 
          /* If!4 */}  else if (C_string.Id() == r.Isa.Id()) /* If:4 */{ 
          /* Let:5 */{ 
            var g0185 *ClaireString   = ToString(r)
            /* noccur = 1 */
            /* Let:6 */{ 
              var v *ClaireAny   = F_value_string(g0185)
              /* noccur = 2 */
              if (v.Isa.IsIn(C_global_variable) == CTRUE) /* If:7 */{ 
                /* Let:8 */{ 
                  var g0186 *GlobalVariable   = ToGlobalVariable(v)
                  /* noccur = 2 */
                  g0186.Store_ask = CTRUE
                  /* Let-8 */} 
                /* If-7 */} 
              /* Let-6 */} 
            /* Let-5 */} 
          /* If-4 */} 
        /* loop-3 */} 
      /* For-2 */} 
    return Result} 
  
// The EID go function for: store @ listargs (throw: false) 
func E_store_listargs (l EID) EID { 
    return /*(sm for store @ listargs= any)*/ F_store_listargs(ToList(OBJ(l)) ).ToEID()} 
  
// *********************************************************************
// *   Part 4: Methods for Native entities                             *
// *********************************************************************
//------------------- STRING -----------------------------------------------
// length(self:string) : integer -> function!(length_string)  -> in Kernel
// make_function(self:string) : function -> function!(make_function_string)
/* {1} OPT.The go function for: symbol!(self:string) [] */
func F_symbol_I_string2 (self *ClaireString ) *ClaireSymbol  { 
    // use function body compiling 
return  F_symbol_I_string(self,C_claire)
    } 
  
// The EID go function for: symbol! @ list<type_expression>(string) (throw: false) 
func E_symbol_I_string2 (self EID) EID { 
    return EID{/*(sm for symbol! @ list<type_expression>(string)= symbol)*/ F_symbol_I_string2(ToString(OBJ(self)) ).Id(),0}} 
  
// deprecated in claire 4 - do not use a string as a byte buffer
/* {1} OPT.The go function for: nth_get(s:string,n:integer,max:integer) [] */
func F_nth_get_string (s *ClaireString ,n int,max int) EID { 
    var Result EID 
    if (n <= max) /* If:2 */{ 
      Result = EID{C__CHAR,CVAL(s.At(n))}
      } else {
      Result = ToException(C_general_error.Make(MakeString("Buffer string access").Id(),CNIL.Id())).Close()
      /* If-2 */} 
    return Result} 
  
// The EID go function for: nth_get @ string (throw: true) 
func E_nth_get_string (s EID,n EID,max EID) EID { 
    return /*(sm for nth_get @ string= EID)*/ F_nth_get_string(ToString(OBJ(s)),INT(n),INT(max) )} 
  
/* {1} OPT.The go function for: nth_put(s:string,n:integer,c:char,max:integer) [] */
func F_nth_put_string (s *ClaireString ,n int,c rune,max int) EID { 
    var Result EID 
    if (n <= max) /* If:2 */{ 
      F_nth_set_string(s,n,c)
      Result = EVOID
      } else {
      Result = ToException(C_general_error.Make(MakeString("Buffer string access").Id(),CNIL.Id())).Close()
      /* If-2 */} 
    return Result} 
  
// The EID go function for: nth_put @ string (throw: true) 
func E_nth_put_string (s EID,n EID,c EID,max EID) EID { 
    return /*(sm for nth_put @ string= EID)*/ F_nth_put_string(ToString(OBJ(s)),
      INT(n),
      CHAR(c),
      INT(max) )} 
  
//  v3.2.14
// we keep the externC method name even if it now support go code
/* {1} OPT.The go function for: externC(s:string) [] */
func F_externC_string (s *ClaireString ) EID { 
    var Result EID 
    Result = ToException(C_general_error.Make(MakeString("cannot execute Go code: ~A").Id(),MakeConstantList((s).Id()).Id())).Close()
    return Result} 
  
// The EID go function for: externC @ list<type_expression>(string) (throw: true) 
func E_externC_string (s EID) EID { 
    return /*(sm for externC @ list<type_expression>(string)= EID)*/ F_externC_string(ToString(OBJ(s)) )} 
  
/* {1} OPT.The go function for: externC(s:string,c:class) [] */
func F_externC_string2 (s *ClaireString ,c *ClaireClass ) EID { 
    var Result EID 
    Result = ToException(C_general_error.Make(MakeString("cannot execute ~A").Id(),MakeConstantList((s).Id()).Id())).Close()
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    Result = EID{CNULL,0}
    }
    return Result} 
  
// The EID go function for: externC @ list<type_expression>(string, class) (throw: true) 
func E_externC_string2 (s EID,c EID) EID { 
    return /*(sm for externC @ list<type_expression>(string, class)= EID)*/ F_externC_string2(ToString(OBJ(s)),ToClass(OBJ(c)) )} 
  
/* {1} OPT.The go function for: externC_string2_type */
func F_externC_string2_type (s *ClaireType ,c *ClaireType ) EID { 
    /* eid body: member(c) */
    var Result EID 
    Result = F_CALL(C_member,ARGS(EID{c.Id(),0}))
    return Result} 
  
  
// The dual EID go function for: "externC_string2_type" 
func E_externC_string2_type (s EID,c EID) EID { 
    return F_externC_string2_type(ToType(OBJ(s)),ToType(OBJ(c)))} 
  
//------------------- SYMBOL -----------------------------------------------
/* {1} OPT.The go function for: make_string(self:symbol) [] */
func F_make_string_symbol (self *ClaireSymbol ) EID { 
    var Result EID 
    F_print_in_string_void()
    self.Princ()
    Result = F_end_of_string_void()
    return Result} 
  
// The EID go function for: make_string @ symbol (throw: true) 
func E_make_string_symbol (self EID) EID { 
    return /*(sm for make_string @ symbol= EID)*/ F_make_string_symbol(ToSymbol(OBJ(self)) )} 
  
//princ(self:symbol) : any -> function!(princ_symbol)
/* {1} OPT.The go function for: self_print(self:symbol) [] */
func F_self_print_symbol_Core (self *ClaireSymbol ) EID { 
    var Result EID 
    self.Module_I().Name.Princ()
    PRINC("/")
    Result = F_print_any((self.String_I()).Id())
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("")
    Result = EVOID
    }
    return Result} 
  
// The EID go function for: self_print @ symbol (throw: true) 
func E_self_print_symbol_Core (self EID) EID { 
    return /*(sm for self_print @ symbol= EID)*/ F_self_print_symbol_Core(ToSymbol(OBJ(self)) )} 
  
//c_princ(self:symbol) : any -> function!(c_princ_symbol)
//gensym(self:string) : symbol -> function!(gensym_string, NEW_ALLOC)
//--------------------- INTEGER -----------------------------------------
/* {1} OPT.The go function for: +(self:integer,x:integer) [] */
func F__plus_integer (self int,x int) int { 
    // use function body compiling 
return  (self+x)
    } 
  
// The EID go function for: + @ list<type_expression>(integer, integer) (throw: false) 
func E__plus_integer (self EID,x EID) EID { 
    return EID{C__INT,IVAL(/*(sm for + @ list<type_expression>(integer, integer)= integer)*/ F__plus_integer(INT(self),INT(x) ))}} 
  
/* {1} OPT.The go function for: _plus_integer_type */
func F__plus_integer_type (self *ClaireType ,x *ClaireType ) EID { 
    /* eid body: abstract_type(+, self, x) */
    var Result EID 
    Result = F_CALL(C_Core_abstract_type,ARGS(EID{C__plus.Id(),0},EID{self.Id(),0},EID{x.Id(),0}))
    return Result} 
  
  
// The dual EID go function for: "_plus_integer_type" 
func E__plus_integer_type (self EID,x EID) EID { 
    return F__plus_integer_type(ToType(OBJ(self)),ToType(OBJ(x)))} 
  
/* {1} OPT.The go function for: -(self:integer,x:integer) [] */
func F__dash_integer1 (self int,x int) int { 
    // use function body compiling 
return  (self-x)
    } 
  
// The EID go function for: - @ list<type_expression>(integer, integer) (throw: false) 
func E__dash_integer1 (self EID,x EID) EID { 
    return EID{C__INT,IVAL(/*(sm for - @ list<type_expression>(integer, integer)= integer)*/ F__dash_integer1(INT(self),INT(x) ))}} 
  
/* {1} OPT.The go function for: _dash_integer1_type */
func F__dash_integer1_type (self *ClaireType ,x *ClaireType ) EID { 
    /* eid body: abstract_type(-, self, x) */
    var Result EID 
    Result = F_CALL(C_Core_abstract_type,ARGS(EID{C__dash.Id(),0},EID{self.Id(),0},EID{x.Id(),0}))
    return Result} 
  
  
// The dual EID go function for: "_dash_integer1_type" 
func E__dash_integer1_type (self EID,x EID) EID { 
    return F__dash_integer1_type(ToType(OBJ(self)),ToType(OBJ(x)))} 
  
//-(self:integer) : integer -> function!(ch_sign)
//float!(self:integer) : float -> function!(to_float)
//mod(self:integer,x:integer) : integer -> function!(mod_integer)
//less_code(n:integer,i:integer) : boolean -> function!(less_code_integer)
/* {1} OPT.The go function for: <<(x:integer,y:integer) [] */
func F__inf_inf_integer (x int,y int) int { 
    // use function body compiling 
return  (x << y)
    } 
  
// The EID go function for: << @ integer (throw: false) 
func E__inf_inf_integer (x EID,y EID) EID { 
    return EID{C__INT,IVAL(/*(sm for << @ integer= integer)*/ F__inf_inf_integer(INT(x),INT(y) ))}} 
  
/* {1} OPT.The go function for: >>(x:integer,y:integer) [] */
func F__sup_sup_integer (x int,y int) int { 
    // use function body compiling 
return  (x >> y)
    } 
  
// The EID go function for: >> @ integer (throw: false) 
func E__sup_sup_integer (x EID,y EID) EID { 
    return EID{C__INT,IVAL(/*(sm for >> @ integer= integer)*/ F__sup_sup_integer(INT(x),INT(y) ))}} 
  
/* {1} OPT.The go function for: and(x:integer,y:integer) [] */
func F_and_integer (x int,y int) int { 
    // use function body compiling 
return  (x & y)
    } 
  
// The EID go function for: and @ integer (throw: false) 
func E_and_integer (x EID,y EID) EID { 
    return EID{C__INT,IVAL(/*(sm for and @ integer= integer)*/ F_and_integer(INT(x),INT(y) ))}} 
  
/* {1} OPT.The go function for: or(x:integer,y:integer) [] */
func F_or_integer (x int,y int) int { 
    // use function body compiling 
return  (x | y)
    } 
  
// The EID go function for: or @ integer (throw: false) 
func E_or_integer (x EID,y EID) EID { 
    return EID{C__INT,IVAL(/*(sm for or @ integer= integer)*/ F_or_integer(INT(x),INT(y) ))}} 
  
// open-coded
/* {1} OPT.The go function for: <(self:integer,x:integer) [] */
func F__inf_integer (self int,x int) *ClaireBoolean  { 
    // use function body compiling 
if (self < x) /* body If:2 */{ 
      return  CTRUE
      } else {
      return  CFALSE
      /* body If-2 */} 
    } 
  
// The EID go function for: < @ integer (throw: false) 
func E__inf_integer (self EID,x EID) EID { 
    return EID{/*(sm for < @ integer= boolean)*/ F__inf_integer(INT(self),INT(x) ).Id(),0}} 
  
/* {1} OPT.The go function for: <=(self:integer,x:integer) [] */
func F__inf_equal_integer (self int,x int) *ClaireBoolean  { 
    // use function body compiling 
if (self <= x) /* body If:2 */{ 
      return  CTRUE
      } else {
      return  CFALSE
      /* body If-2 */} 
    } 
  
// The EID go function for: <= @ integer (throw: false) 
func E__inf_equal_integer (self EID,x EID) EID { 
    return EID{/*(sm for <= @ integer= boolean)*/ F__inf_equal_integer(INT(self),INT(x) ).Id(),0}} 
  
/* {1} OPT.The go function for: >(self:integer,x:integer) [] */
func F__sup_integer (self int,x int) *ClaireBoolean  { 
    // use function body compiling 
if (self > x) /* body If:2 */{ 
      return  CTRUE
      } else {
      return  CFALSE
      /* body If-2 */} 
    } 
  
// The EID go function for: > @ integer (throw: false) 
func E__sup_integer (self EID,x EID) EID { 
    return EID{/*(sm for > @ integer= boolean)*/ F__sup_integer(INT(self),INT(x) ).Id(),0}} 
  
/* {1} OPT.The go function for: nth(self:integer,y:integer) [] */
func F_nth_integer (self int,y int) *ClaireBoolean  { 
    // use function body compiling 
if (F_nth_integer(self,y) == CTRUE) /* body If:2 */{ 
      return  CTRUE
      } else {
      return  CFALSE
      /* body If-2 */} 
    } 
  
// The EID go function for: nth @ integer (throw: false) 
func E_nth_integer (self EID,y EID) EID { 
    return EID{/*(sm for nth @ integer= boolean)*/ F_nth_integer(INT(self),INT(y) ).Id(),0}} 
  
/* {1} OPT.The go function for: abs(x:integer) [] */
func F_abs_integer (x int) int { 
    // use function body compiling 
if (x >= 0) /* body If:2 */{ 
      return  x
      } else {
      return  (-x)
      /* body If-2 */} 
    } 
  
// The EID go function for: abs @ integer (throw: false) 
func E_abs_integer (x EID) EID { 
    return EID{C__INT,IVAL(/*(sm for abs @ integer= integer)*/ F_abs_integer(INT(x) ))}} 
  
/* {1} OPT.The go function for: random(a:integer,b:integer) [] */
func F_random_integer2 (a int,b int) int { 
    // use function body compiling 
return  (a+F_random_integer(((b+1)-a)))
    } 
  
// The EID go function for: random @ list<type_expression>(integer, integer) (throw: false) 
func E_random_integer2 (a EID,b EID) EID { 
    return EID{C__INT,IVAL(/*(sm for random @ list<type_expression>(integer, integer)= integer)*/ F_random_integer2(INT(a),INT(b) ))}} 
  
// used by the logic
/* {1} OPT.The go function for: factor?(x:integer,y:integer) [] */
func F_factor_ask_integer (x int,y int) EID { 
    var Result EID 
    /* Let:2 */{ 
      var g0187UU int 
      /* noccur = 1 */
      var g0187UU_try01883 EID 
      g0187UU_try01883 = F_mod_integer(x,y)
      /* ERROR PROTECTION INSERTED (g0187UU-Result) */
      if ErrorIn(g0187UU_try01883) {Result = g0187UU_try01883
      } else {
      g0187UU = INT(g0187UU_try01883)
      Result = EID{Equal(MakeInteger(g0187UU).Id(),MakeInteger(0).Id()).Id(),0}
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: factor? @ integer (throw: true) 
func E_factor_ask_integer (x EID,y EID) EID { 
    return /*(sm for factor? @ integer= EID)*/ F_factor_ask_integer(INT(x),INT(y) )} 
  
/* {1} OPT.The go function for: divide?(x:integer,y:integer) [] */
func F_divide_ask_integer (x int,y int) EID { 
    var Result EID 
    /* Let:2 */{ 
      var g0189UU int 
      /* noccur = 1 */
      var g0189UU_try01903 EID 
      g0189UU_try01903 = F_mod_integer(y,x)
      /* ERROR PROTECTION INSERTED (g0189UU-Result) */
      if ErrorIn(g0189UU_try01903) {Result = g0189UU_try01903
      } else {
      g0189UU = INT(g0189UU_try01903)
      Result = EID{Equal(MakeInteger(g0189UU).Id(),MakeInteger(0).Id()).Id(),0}
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: divide? @ integer (throw: true) 
func E_divide_ask_integer (x EID,y EID) EID { 
    return /*(sm for divide? @ integer= EID)*/ F_divide_ask_integer(INT(x),INT(y) )} 
  
/* {1} OPT.The go function for: Id(x:any) [] */
func F_Id_any (x *ClaireAny ) *ClaireAny  { 
    // use function body compiling 
return  x
    } 
  
// The EID go function for: Id @ any (throw: false) 
func E_Id_any (x EID) EID { 
    return /*(sm for Id @ any= any)*/ F_Id_any(ANY(x) ).ToEID()} 
  
/* {1} OPT.The go function for: Id_any_type */
func F_Id_any_type (x *ClaireType ) EID { 
    /* eid body: x */
    var Result EID 
    Result = EID{x.Id(),0}
    return Result} 
  
  
// The dual EID go function for: "Id_any_type" 
func E_Id_any_type (x EID) EID { 
    return F_Id_any_type(ToType(OBJ(x)))} 
  
/* {1} OPT.The go function for: pair(x:any,y:any) [] */
func F_pair_any (x *ClaireAny ,y *ClaireAny ) *ClaireList  { 
    // use function body compiling 
return  MakeConstantList(x,y)
    } 
  
// The EID go function for: pair @ any (throw: false) 
func E_pair_any (x EID,y EID) EID { 
    return EID{/*(sm for pair @ any= list)*/ F_pair_any(ANY(x),ANY(y) ).Id(),0}} 
  
/* {1} OPT.The go function for: pair_1(x:list) [] */
func F_pair_1_list (x *ClaireList ) *ClaireAny  { 
    // use function body compiling 
return  x.At(1-1)
    } 
  
// The EID go function for: pair_1 @ list (throw: false) 
func E_pair_1_list (x EID) EID { 
    return /*(sm for pair_1 @ list= any)*/ F_pair_1_list(ToList(OBJ(x)) ).ToEID()} 
  
/* {1} OPT.The go function for: pair_1_list_type */
func F_pair_1_list_type (x *ClaireType ) EID { 
    /* eid body: member(x) */
    var Result EID 
    Result = F_CALL(C_member,ARGS(EID{x.Id(),0}))
    return Result} 
  
  
// The dual EID go function for: "pair_1_list_type" 
func E_pair_1_list_type (x EID) EID { 
    return F_pair_1_list_type(ToType(OBJ(x)))} 
  
/* {1} OPT.The go function for: pair_2(x:list) [] */
func F_pair_2_list (x *ClaireList ) *ClaireAny  { 
    // use function body compiling 
return  x.At(2-1)
    } 
  
// The EID go function for: pair_2 @ list (throw: false) 
func E_pair_2_list (x EID) EID { 
    return /*(sm for pair_2 @ list= any)*/ F_pair_2_list(ToList(OBJ(x)) ).ToEID()} 
  
/* {1} OPT.The go function for: pair_2_list_type */
func F_pair_2_list_type (x *ClaireType ) EID { 
    /* eid body: member(x) */
    var Result EID 
    Result = F_CALL(C_member,ARGS(EID{x.Id(),0}))
    return Result} 
  
  
// The dual EID go function for: "pair_2_list_type" 
func E_pair_2_list_type (x EID) EID { 
    return F_pair_2_list_type(ToType(OBJ(x)))} 
  
//------------------------ FLOAT ---------------------------------------------
/* {1} OPT.The go function for: +(self:float,x:float) [] */
func F__plus_float (self float64,x float64) float64 { 
    // procedure body with s =  
var Result float64 
    /* Let:2 */{ 
      var y float64  = (self+x)
      /* noccur = 1 */
      Result = y
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: + @ list<type_expression>(float, float) (throw: false) 
func E__plus_float (self EID,x EID) EID { 
    return EID{C__FLOAT,FVAL(/*(sm for + @ list<type_expression>(float, float)= float)*/ F__plus_float(FLOAT(self),FLOAT(x) ))}} 
  
/* {1} OPT.The go function for: -(self:float,x:float) [] */
func F__dash_float (self float64,x float64) float64 { 
    // procedure body with s =  
var Result float64 
    /* Let:2 */{ 
      var y float64  = (self-x)
      /* noccur = 1 */
      Result = y
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: - @ list<type_expression>(float, float) (throw: false) 
func E__dash_float (self EID,x EID) EID { 
    return EID{C__FLOAT,FVAL(/*(sm for - @ list<type_expression>(float, float)= float)*/ F__dash_float(FLOAT(self),FLOAT(x) ))}} 
  
/* {1} OPT.The go function for: *(self:float,x:float) [] */
func F__star_float (self float64,x float64) float64 { 
    // procedure body with s =  
var Result float64 
    /* Let:2 */{ 
      var y float64  = (self*x)
      /* noccur = 1 */
      Result = y
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: * @ list<type_expression>(float, float) (throw: false) 
func E__star_float (self EID,x EID) EID { 
    return EID{C__FLOAT,FVAL(/*(sm for * @ list<type_expression>(float, float)= float)*/ F__star_float(FLOAT(self),FLOAT(x) ))}} 
  
/* {1} OPT.The go function for: /(self:float,x:float) [] */
func F__7_float (self float64,x float64) float64 { 
    // procedure body with s =  
var Result float64 
    /* Let:2 */{ 
      var y float64  = (self/x)
      /* noccur = 1 */
      Result = y
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: / @ list<type_expression>(float, float) (throw: false) 
func E__7_float (self EID,x EID) EID { 
    return EID{C__FLOAT,FVAL(/*(sm for / @ list<type_expression>(float, float)= float)*/ F__7_float(FLOAT(self),FLOAT(x) ))}} 
  
/* {1} OPT.The go function for: -(self:float) [] */
func F__dash_float2 (self float64) float64 { 
    // use function body compiling 
return  ((-1)*self)
    } 
  
// The EID go function for: - @ list<type_expression>(float) (throw: false) 
func E__dash_float2 (self EID) EID { 
    return EID{C__FLOAT,FVAL(/*(sm for - @ list<type_expression>(float)= float)*/ F__dash_float2(FLOAT(self) ))}} 
  
/* {1} OPT.The go function for: string!(self:float) [] */
func F_string_I_float (self float64) EID { 
    var Result EID 
    F_print_in_string_void()
    F_princ_float(self)
    Result = F_end_of_string_void()
    return Result} 
  
// The EID go function for: string! @ float (throw: true) 
func E_string_I_float (self EID) EID { 
    return /*(sm for string! @ float= EID)*/ F_string_I_float(FLOAT(self) )} 
  
// v3.3.42
/* {1} OPT.The go function for: abs(x:float) [] */
func F_abs_float (x float64) float64 { 
    // use function body compiling 
if (x >= 0) /* body If:2 */{ 
      return  x
      } else {
      return  (-x)
      /* body If-2 */} 
    } 
  
// The EID go function for: abs @ float (throw: false) 
func E_abs_float (x EID) EID { 
    return EID{C__FLOAT,FVAL(/*(sm for abs @ float= float)*/ F_abs_float(FLOAT(x) ))}} 
  
// the pF is my ugly duckling :) -------------------------------------------
// float print is now standard in v3.4.42 (princ(float_integer)  but this is still a cuter print ...
/* {1} OPT.The go function for: mClaire/printFDigit(x:float,i:integer) [] */
func F_printFDigit_float (x float64,i int) EID { 
    var Result EID 
    if (x < 0) /* If:2 */{ 
      PRINC("-")
      Result = F_printFDigit_float((-x),i)
      } else {
      /* Let:3 */{ 
        var frac float64 
        /* noccur = 1 */
        var frac_try01914 EID 
        /* Let:4 */{ 
          var g0192UU float64 
          /* noccur = 1 */
          var g0192UU_try01935 EID 
          /* Let:5 */{ 
            var g0194UU float64 
            /* noccur = 1 */
            var g0194UU_try01956 EID 
            /* Let:6 */{ 
              var g0196UU int 
              /* noccur = 1 */
              var g0196UU_try01977 EID 
              g0196UU_try01977 = F_integer_I_float((x+1e-10))
              /* ERROR PROTECTION INSERTED (g0196UU-g0194UU_try01956) */
              if ErrorIn(g0196UU_try01977) {g0194UU_try01956 = g0196UU_try01977
              } else {
              g0196UU = INT(g0196UU_try01977)
              g0194UU_try01956 = EID{C__FLOAT,FVAL(F_to_float(g0196UU))}
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (g0194UU-g0192UU_try01935) */
            if ErrorIn(g0194UU_try01956) {g0192UU_try01935 = g0194UU_try01956
            } else {
            g0194UU = FLOAT(g0194UU_try01956)
            g0192UU_try01935 = EID{C__FLOAT,FVAL((x-g0194UU))}
            }
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (g0192UU-frac_try01914) */
          if ErrorIn(g0192UU_try01935) {frac_try01914 = g0192UU_try01935
          } else {
          g0192UU = FLOAT(g0192UU_try01935)
          frac_try01914 = EID{C__FLOAT,FVAL((g0192UU+1e-10))}
          }
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (frac-Result) */
        if ErrorIn(frac_try01914) {Result = frac_try01914
        } else {
        frac = FLOAT(frac_try01914)
        /* Let:4 */{ 
          var g0198UU int 
          /* noccur = 1 */
          var g0198UU_try01995 EID 
          g0198UU_try01995 = F_integer_I_float((x+1e-10))
          /* ERROR PROTECTION INSERTED (g0198UU-Result) */
          if ErrorIn(g0198UU_try01995) {Result = g0198UU_try01995
          } else {
          g0198UU = INT(g0198UU_try01995)
          F_princ_integer(g0198UU)
          Result = EVOID
          }
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC(".")
        /* Let:4 */{ 
          var g0200UU int 
          /* noccur = 1 */
          var g0200UU_try02015 EID 
          g0200UU_try02015 = F_integer_I_float((frac*F__exp_float(10,F_to_float(i))))
          /* ERROR PROTECTION INSERTED (g0200UU-Result) */
          if ErrorIn(g0200UU_try02015) {Result = g0200UU_try02015
          } else {
          g0200UU = INT(g0200UU_try02015)
          Result = F_printFDigit_integer(g0200UU,i)
          }
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC("")
        Result = EVOID
        }}
        }
        /* Let-3 */} 
      /* If-2 */} 
    return Result} 
  
// The EID go function for: mClaire/printFDigit @ float (throw: true) 
func E_printFDigit_float (x EID,i EID) EID { 
    return /*(sm for mClaire/printFDigit @ float= EID)*/ F_printFDigit_float(FLOAT(x),INT(i) )} 
  
// print the first i digits of an integer
/* {1} OPT.The go function for: mClaire/printFDigit(x:integer,i:integer) [] */
func F_printFDigit_integer (x int,i int) EID { 
    var Result EID 
    if (i > 0) /* If:2 */{ 
      /* Let:3 */{ 
        var f int 
        /* noccur = 2 */
        var f_try02024 EID 
        f_try02024 = F__exp_integer(10,(i-1))
        /* ERROR PROTECTION INSERTED (f-Result) */
        if ErrorIn(f_try02024) {Result = f_try02024
        } else {
        f = INT(f_try02024)
        /* Let:4 */{ 
          var d int 
          /* noccur = 1 */
          var d_try02035 EID 
          d_try02035 = EID{C__INT,IVAL((x/f))}
          /* ERROR PROTECTION INSERTED (d-Result) */
          if ErrorIn(d_try02035) {Result = d_try02035
          } else {
          d = INT(d_try02035)
          F_princ_integer(d)
          if (i > 1) /* If:5 */{ 
            /* Let:6 */{ 
              var g0204UU int 
              /* noccur = 1 */
              var g0204UU_try02057 EID 
              g0204UU_try02057 = F_mod_integer(x,f)
              /* ERROR PROTECTION INSERTED (g0204UU-Result) */
              if ErrorIn(g0204UU_try02057) {Result = g0204UU_try02057
              } else {
              g0204UU = INT(g0204UU_try02057)
              Result = F_printFDigit_integer(g0204UU,(i-1))
              }
              /* Let-6 */} 
            } else {
            Result = EID{CFALSE.Id(),0}
            /* If-5 */} 
          }
          /* Let-4 */} 
        }
        /* Let-3 */} 
      } else {
      Result = EID{CFALSE.Id(),0}
      /* If-2 */} 
    return Result} 
  
// The EID go function for: mClaire/printFDigit @ integer (throw: true) 
func E_printFDigit_integer (x EID,i EID) EID { 
    return /*(sm for mClaire/printFDigit @ integer= EID)*/ F_printFDigit_integer(INT(x),INT(i) )} 
  
//--------- BAG --------------------------------------------------------
// in CLAIRE 4, bag is no longer a concrete data type, hence bag methods must be divided between sets and lists
// length(self:bag ) : integer -> length(self)  => becomes part of go
// in CLAIRE 4, we can tell the compiler that the list is a list<object> (optimization purpose)
/* {1} OPT.The go function for: mClaire/nth_object(self:list,n:integer) [] */
func F_mClaire_nth_object_list (self *ClaireList ,n int) *ClaireAny  { 
    // use function body compiling 
return  self.ValuesO()[n-1].Id()
    } 
  
// The EID go function for: mClaire/nth_object @ list (throw: false) 
func E_mClaire_nth_object_list (self EID,n EID) EID { 
    return /*(sm for mClaire/nth_object @ list= any)*/ F_mClaire_nth_object_list(ToList(OBJ(self)),INT(n) ).ToEID()} 
  
//
// nth_get(self:list,x:integer) : any -> nth_get(self, x)
// new in claire 4: tells the compiler that range check is required + EID optimized
/* {1} OPT.The go function for: nth_write(self:list,i:integer,v:any) [] */
func F_nth_write_list (self *ClaireList ,i int,v *ClaireAny ) EID { 
    var Result EID 
    if (self.Of().Contains(v) == CTRUE) /* If:2 */{ 
      Result = ToArray(self.Id()).NthPut(i,v).ToEID()
      } else {
      Result = ToException(C_system_error.Make(MakeInteger(17).Id(),v,self.Of().Id())).Close()
      /* If-2 */} 
    return Result} 
  
// The EID go function for: nth_write @ list (throw: true) 
func E_nth_write_list (self EID,i EID,v EID) EID { 
    return /*(sm for nth_write @ list= EID)*/ F_nth_write_list(ToList(OBJ(self)),INT(i),ANY(v) )} 
  
// CLAIRE 4 duplication: define min/max for sets first
/* {1} OPT.The go function for: min(f:method,self:set) [] */
func F_min_method2 (f *ClaireMethod ,self *ClaireSet ) EID { 
    var Result EID 
    if (Equal(ANY(F_CALL(C_length,ARGS(EID{self.Id(),0}))),MakeInteger(0).Id()) != CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var x *ClaireAny   = CNULL
        /* noccur = 4 */
        /* For:4 */{ 
          var y *ClaireAny  
          _ = y
          Result= EID{CFALSE.Id(),0}
          for _,y = range(self.Values)/* loop:5 */{ 
            var void_try6 EID 
            _ = void_try6
            var g0206I *ClaireBoolean  
            var g0206I_try02076 EID 
            /* or:6 */{ 
              var v_or6 *ClaireBoolean  
              
              v_or6 = Equal(x,CNULL)
              if (v_or6 == CTRUE) {g0206I_try02076 = EID{CTRUE.Id(),0}
              } else /* or:7 */{ 
                var v_or6_try02088 EID 
                v_or6_try02088 = F_funcall_method2(f,y,x)
                /* ERROR PROTECTION INSERTED (v_or6-g0206I_try02076) */
                if ErrorIn(v_or6_try02088) {g0206I_try02076 = v_or6_try02088
                } else {
                v_or6 = ToBoolean(OBJ(v_or6_try02088))
                if (v_or6 == CTRUE) {g0206I_try02076 = EID{CTRUE.Id(),0}
                } else /* or:8 */{ 
                  g0206I_try02076 = EID{CFALSE.Id(),0}/* org-8 */} 
                /* org-7 */} 
              }
              /* or-6 */} 
            /* ERROR PROTECTION INSERTED (g0206I-void_try6) */
            if ErrorIn(g0206I_try02076) {void_try6 = g0206I_try02076
            } else {
            g0206I = ToBoolean(OBJ(g0206I_try02076))
            if (g0206I == CTRUE) /* If:6 */{ 
              x = y
              void_try6 = x.ToEID()
              } else {
              void_try6 = EID{CFALSE.Id(),0}
              /* If-6 */} 
            }
            /* ERROR PROTECTION INSERTED (void_try6-Result) */
            if ErrorIn(void_try6) {Result = void_try6
            Result = void_try6
            break
            } else {
            }
            /* loop-5 */} 
          /* For-4 */} 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        Result = x.ToEID()
        }
        /* Let-3 */} 
      } else {
      Result = ToException(C_general_error.Make(MakeString("[183] min of empty set is undefined").Id(),CNIL.Id())).Close()
      /* If-2 */} 
    return Result} 
  
// The EID go function for: min @ list<type_expression>(method, set) (throw: true) 
func E_min_method2 (f EID,self EID) EID { 
    return /*(sm for min @ list<type_expression>(method, set)= EID)*/ F_min_method2(ToMethod(OBJ(f)),ToSet(OBJ(self)) )} 
  
/* {1} OPT.The go function for: min_method2_type */
func F_min_method2_type (f *ClaireType ,self *ClaireType ) EID { 
    /* eid body: member(self) */
    var Result EID 
    Result = F_CALL(C_member,ARGS(EID{self.Id(),0}))
    return Result} 
  
  
// The dual EID go function for: "min_method2_type" 
func E_min_method2_type (f EID,self EID) EID { 
    return F_min_method2_type(ToType(OBJ(f)),ToType(OBJ(self)))} 
  
/* {1} OPT.The go function for: max(f:method,self:set) [] */
func F_max_method2 (f *ClaireMethod ,self *ClaireSet ) EID { 
    var Result EID 
    if (Equal(ANY(F_CALL(C_length,ARGS(EID{self.Id(),0}))),MakeInteger(0).Id()) != CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var x *ClaireAny   = CNULL
        /* noccur = 4 */
        /* For:4 */{ 
          var y *ClaireAny  
          _ = y
          Result= EID{CFALSE.Id(),0}
          for _,y = range(self.Values)/* loop:5 */{ 
            var void_try6 EID 
            _ = void_try6
            var g0209I *ClaireBoolean  
            var g0209I_try02106 EID 
            /* or:6 */{ 
              var v_or6 *ClaireBoolean  
              
              v_or6 = Equal(x,CNULL)
              if (v_or6 == CTRUE) {g0209I_try02106 = EID{CTRUE.Id(),0}
              } else /* or:7 */{ 
                var v_or6_try02118 EID 
                /* Let:8 */{ 
                  var g0212UU *ClaireBoolean  
                  /* noccur = 1 */
                  var g0212UU_try02139 EID 
                  /* Let:9 */{ 
                    var g0214UU *ClaireAny  
                    /* noccur = 1 */
                    var g0214UU_try021510 EID 
                    g0214UU_try021510 = F_funcall_method2(f,y,x)
                    /* ERROR PROTECTION INSERTED (g0214UU-g0212UU_try02139) */
                    if ErrorIn(g0214UU_try021510) {g0212UU_try02139 = g0214UU_try021510
                    } else {
                    g0214UU = ANY(g0214UU_try021510)
                    g0212UU_try02139 = EID{F_boolean_I_any(g0214UU).Id(),0}
                    }
                    /* Let-9 */} 
                  /* ERROR PROTECTION INSERTED (g0212UU-v_or6_try02118) */
                  if ErrorIn(g0212UU_try02139) {v_or6_try02118 = g0212UU_try02139
                  } else {
                  g0212UU = ToBoolean(OBJ(g0212UU_try02139))
                  v_or6_try02118 = EID{F__I_equal_any(g0212UU.Id(),CTRUE.Id()).Id(),0}
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (v_or6-g0209I_try02106) */
                if ErrorIn(v_or6_try02118) {g0209I_try02106 = v_or6_try02118
                } else {
                v_or6 = ToBoolean(OBJ(v_or6_try02118))
                if (v_or6 == CTRUE) {g0209I_try02106 = EID{CTRUE.Id(),0}
                } else /* or:8 */{ 
                  g0209I_try02106 = EID{CFALSE.Id(),0}/* org-8 */} 
                /* org-7 */} 
              }
              /* or-6 */} 
            /* ERROR PROTECTION INSERTED (g0209I-void_try6) */
            if ErrorIn(g0209I_try02106) {void_try6 = g0209I_try02106
            } else {
            g0209I = ToBoolean(OBJ(g0209I_try02106))
            if (g0209I == CTRUE) /* If:6 */{ 
              x = y
              void_try6 = x.ToEID()
              } else {
              void_try6 = EID{CFALSE.Id(),0}
              /* If-6 */} 
            }
            /* ERROR PROTECTION INSERTED (void_try6-Result) */
            if ErrorIn(void_try6) {Result = void_try6
            Result = void_try6
            break
            } else {
            }
            /* loop-5 */} 
          /* For-4 */} 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        Result = x.ToEID()
        }
        /* Let-3 */} 
      } else {
      Result = ToException(C_general_error.Make(MakeString("[183] max of empty set is undefined").Id(),CNIL.Id())).Close()
      /* If-2 */} 
    return Result} 
  
// The EID go function for: max @ list<type_expression>(method, set) (throw: true) 
func E_max_method2 (f EID,self EID) EID { 
    return /*(sm for max @ list<type_expression>(method, set)= EID)*/ F_max_method2(ToMethod(OBJ(f)),ToSet(OBJ(self)) )} 
  
/* {1} OPT.The go function for: max_method2_type */
func F_max_method2_type (f *ClaireType ,self *ClaireType ) EID { 
    /* eid body: member(self) */
    var Result EID 
    Result = F_CALL(C_member,ARGS(EID{self.Id(),0}))
    return Result} 
  
  
// The dual EID go function for: "max_method2_type" 
func E_max_method2_type (f EID,self EID) EID { 
    return F_max_method2_type(ToType(OBJ(f)),ToType(OBJ(self)))} 
  
// CLAIRE 4 : optimize for lists
/* {1} OPT.The go function for: min(f:method,self:list) [] */
func F_min_method3 (f *ClaireMethod ,self *ClaireList ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var n int  = self.Length()
      /* noccur = 2 */
      if (n != 0) /* If:3 */{ 
        /* Let:4 */{ 
          var x *ClaireAny   = self.At(1-1)
          /* noccur = 3 */
          /* Let:5 */{ 
            var i int  = 2
            /* noccur = 5 */
            /* Let:6 */{ 
              var g0216 int  = n
              /* noccur = 1 */
              Result= EID{CFALSE.Id(),0}
              for (i <= g0216) /* while:7 */{ 
                var void_try8 EID 
                _ = void_try8
                { 
                var g0217I *ClaireBoolean  
                var g0217I_try02188 EID 
                g0217I_try02188 = F_funcall_method2(f,self.At(i-1),x)
                /* ERROR PROTECTION INSERTED (g0217I-void_try8) */
                if ErrorIn(g0217I_try02188) {void_try8 = g0217I_try02188
                } else {
                g0217I = ToBoolean(OBJ(g0217I_try02188))
                if (g0217I == CTRUE) /* If:8 */{ 
                  x = self.At(i-1)
                  void_try8 = x.ToEID()
                  } else {
                  void_try8 = EID{CFALSE.Id(),0}
                  /* If-8 */} 
                }
                /* ERROR PROTECTION INSERTED (void_try8-void_try8) */
                if ErrorIn(void_try8) {Result = void_try8
                break
                } else {
                i = (i+1)
                }
                /* while-7 */} 
              }
              /* Let-6 */} 
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          Result = x.ToEID()
          }
          /* Let-4 */} 
        } else {
        Result = ToException(C_general_error.Make(MakeString("[183] min of empty list is undefined").Id(),CNIL.Id())).Close()
        /* If-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: min @ list<type_expression>(method, list) (throw: true) 
func E_min_method3 (f EID,self EID) EID { 
    return /*(sm for min @ list<type_expression>(method, list)= EID)*/ F_min_method3(ToMethod(OBJ(f)),ToList(OBJ(self)) )} 
  
/* {1} OPT.The go function for: min_method3_type */
func F_min_method3_type (f *ClaireType ,self *ClaireType ) EID { 
    /* eid body: member(self) */
    var Result EID 
    Result = F_CALL(C_member,ARGS(EID{self.Id(),0}))
    return Result} 
  
  
// The dual EID go function for: "min_method3_type" 
func E_min_method3_type (f EID,self EID) EID { 
    return F_min_method3_type(ToType(OBJ(f)),ToType(OBJ(self)))} 
  
/* {1} OPT.The go function for: max(f:method,self:list) [] */
func F_max_method3 (f *ClaireMethod ,self *ClaireList ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var n int  = self.Length()
      /* noccur = 2 */
      if (n != 0) /* If:3 */{ 
        /* Let:4 */{ 
          var x *ClaireAny   = self.At(1-1)
          /* noccur = 3 */
          /* Let:5 */{ 
            var i int  = 2
            /* noccur = 5 */
            /* Let:6 */{ 
              var g0219 int  = n
              /* noccur = 1 */
              Result= EID{CFALSE.Id(),0}
              for (i <= g0219) /* while:7 */{ 
                var void_try8 EID 
                _ = void_try8
                { 
                var g0220I *ClaireBoolean  
                var g0220I_try02218 EID 
                /* Let:8 */{ 
                  var g0222UU *ClaireBoolean  
                  /* noccur = 1 */
                  var g0222UU_try02239 EID 
                  /* Let:9 */{ 
                    var g0224UU *ClaireAny  
                    /* noccur = 1 */
                    var g0224UU_try022510 EID 
                    g0224UU_try022510 = F_funcall_method2(f,self.At(i-1),x)
                    /* ERROR PROTECTION INSERTED (g0224UU-g0222UU_try02239) */
                    if ErrorIn(g0224UU_try022510) {g0222UU_try02239 = g0224UU_try022510
                    } else {
                    g0224UU = ANY(g0224UU_try022510)
                    g0222UU_try02239 = EID{F_boolean_I_any(g0224UU).Id(),0}
                    }
                    /* Let-9 */} 
                  /* ERROR PROTECTION INSERTED (g0222UU-g0220I_try02218) */
                  if ErrorIn(g0222UU_try02239) {g0220I_try02218 = g0222UU_try02239
                  } else {
                  g0222UU = ToBoolean(OBJ(g0222UU_try02239))
                  g0220I_try02218 = EID{F__I_equal_any(g0222UU.Id(),CTRUE.Id()).Id(),0}
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (g0220I-void_try8) */
                if ErrorIn(g0220I_try02218) {void_try8 = g0220I_try02218
                } else {
                g0220I = ToBoolean(OBJ(g0220I_try02218))
                if (g0220I == CTRUE) /* If:8 */{ 
                  x = self.At(i-1)
                  void_try8 = x.ToEID()
                  } else {
                  void_try8 = EID{CFALSE.Id(),0}
                  /* If-8 */} 
                }
                /* ERROR PROTECTION INSERTED (void_try8-void_try8) */
                if ErrorIn(void_try8) {Result = void_try8
                break
                } else {
                i = (i+1)
                }
                /* while-7 */} 
              }
              /* Let-6 */} 
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          Result = x.ToEID()
          }
          /* Let-4 */} 
        } else {
        Result = ToException(C_general_error.Make(MakeString("[183] max of empty list is undefined").Id(),CNIL.Id())).Close()
        /* If-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: max @ list<type_expression>(method, list) (throw: true) 
func E_max_method3 (f EID,self EID) EID { 
    return /*(sm for max @ list<type_expression>(method, list)= EID)*/ F_max_method3(ToMethod(OBJ(f)),ToList(OBJ(self)) )} 
  
/* {1} OPT.The go function for: max_method3_type */
func F_max_method3_type (f *ClaireType ,self *ClaireType ) EID { 
    /* eid body: member(self) */
    var Result EID 
    Result = F_CALL(C_member,ARGS(EID{self.Id(),0}))
    return Result} 
  
  
// The dual EID go function for: "max_method3_type" 
func E_max_method3_type (f EID,self EID) EID { 
    return F_max_method3_type(ToType(OBJ(f)),ToType(OBJ(self)))} 
  
// CLAIRE4 : /+ is native for list
// new for claire 3.4
/* {1} OPT.The go function for: random(self:list) [] */
func F_random_list (self *ClaireList ) *ClaireAny  { 
    // procedure body with s =  
var Result *ClaireAny  
    /* Let:2 */{ 
      var n int  = self.Length()
      /* noccur = 1 */
      Result = self.At((1+F_random_integer(n))-1)
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: random @ list (throw: false) 
func E_random_list (self EID) EID { 
    return /*(sm for random @ list= any)*/ F_random_list(ToList(OBJ(self)) ).ToEID()} 
  
/* {1} OPT.The go function for: random_list_type */
func F_random_list_type (self *ClaireType ) EID { 
    /* eid body: member(self) */
    var Result EID 
    Result = F_CALL(C_member,ARGS(EID{self.Id(),0}))
    return Result} 
  
  
// The dual EID go function for: "random_list_type" 
func E_random_list_type (self EID) EID { 
    return F_random_list_type(ToType(OBJ(self)))} 
  
//--------- LIST --------------------------------------------------------
// last element of a list
/* {1} OPT.The go function for: last(self:list) [] */
func F_last_list (self *ClaireList ) EID { 
    var Result EID 
    if (self.Length() > 0) /* If:2 */{ 
      Result = self.At(self.Length()-1).ToEID()
      } else {
      Result = ToException(C_general_error.Make(MakeString("[41] car(nil) is undefined").Id(),CNIL.Id())).Close()
      /* If-2 */} 
    return Result} 
  
// The EID go function for: last @ list (throw: true) 
func E_last_list (self EID) EID { 
    return /*(sm for last @ list= EID)*/ F_last_list(ToList(OBJ(self)) )} 
  
/* {1} OPT.The go function for: last_list_type */
func F_last_list_type (self *ClaireType ) EID { 
    /* eid body: member(self) */
    var Result EID 
    Result = F_CALL(C_member,ARGS(EID{self.Id(),0}))
    return Result} 
  
  
// The dual EID go function for: "last_list_type" 
func E_last_list_type (self EID) EID { 
    return F_last_list_type(ToType(OBJ(self)))} 
  
// remove the last element
/* {1} OPT.The go function for: rmlast(self:list) [] */
func F_rmlast_list (self *ClaireList ) EID { 
    var Result EID 
    Result = self.Nth_dash(self.Length())
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    Result = EID{self.Id(),0}
    }
    return Result} 
  
// The EID go function for: rmlast @ list (throw: true) 
func E_rmlast_list (self EID) EID { 
    return /*(sm for rmlast @ list= EID)*/ F_rmlast_list(ToList(OBJ(self)) )} 
  
// the old LISP method
/* {1} OPT.The go function for: car(self:list) [] */
func F_car_list (self *ClaireList ) EID { 
    var Result EID 
    if (self.Length() > 0) /* If:2 */{ 
      Result = self.At(1-1).ToEID()
      } else {
      Result = ToException(C_general_error.Make(MakeString("[41] car(nil) is undefined").Id(),CNIL.Id())).Close()
      /* If-2 */} 
    return Result} 
  
// The EID go function for: car @ list (throw: true) 
func E_car_list (self EID) EID { 
    return /*(sm for car @ list= EID)*/ F_car_list(ToList(OBJ(self)) )} 
  
// this method sorts a list according to an order
/* {1} OPT.The go function for: sort(f:method,self:list) [] */
func F_sort_method (f *ClaireMethod ,self *ClaireList ) EID { 
    var Result EID 
    Result = F_quicksort_list(self,f,1,self.Length())
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    Result = EID{self.Id(),0}
    }
    return Result} 
  
// The EID go function for: sort @ method (throw: true) 
func E_sort_method (f EID,self EID) EID { 
    return /*(sm for sort @ method= EID)*/ F_sort_method(ToMethod(OBJ(f)),ToList(OBJ(self)) )} 
  
// v3.0.38: upgrade the quicksort algorithm with a better pivot selection cf.bag.cpp
// this is also proposed as a macro: cf. file.cl
/* {1} OPT.The go function for: quicksort(self:list,f:method,n:integer,m:integer) [] */
func F_quicksort_list (self *ClaireList ,f *ClaireMethod ,n int,m int) EID { 
    var Result EID 
    if (m > n) /* If:2 */{ 
      /* Let:3 */{ 
        var x *ClaireAny   = self.At(n-1)
        /* noccur = 5 */
        if (m == (n+1)) /* If:4 */{ 
          var g0227I *ClaireBoolean  
          var g0227I_try02285 EID 
          g0227I_try02285 = F_funcall_method2(f,self.At(m-1),x)
          /* ERROR PROTECTION INSERTED (g0227I-Result) */
          if ErrorIn(g0227I_try02285) {Result = g0227I_try02285
          } else {
          g0227I = ToBoolean(OBJ(g0227I_try02285))
          if (g0227I == CTRUE) /* If:5 */{ 
            ToArray(self.Id()).NthPut(n,self.At(m-1))
            Result = ToArray(self.Id()).NthPut(m,x).ToEID()
            } else {
            Result = EID{CFALSE.Id(),0}
            /* If-5 */} 
          }
          } else {
          /* Let:5 */{ 
            var p int  = ((m+n)>>1)
            /* noccur = 3 */
            /* Let:6 */{ 
              var q int  = n
              /* noccur = 1 */
              x = self.At(p-1)
              if (p != n) /* If:7 */{ 
                ToArray(self.Id()).NthPut(p,self.At(n-1))
                /* If-7 */} 
              /* Let:7 */{ 
                var p int  = (n+1)
                /* noccur = 7 */
                /* Let:8 */{ 
                  var g0226 int  = m
                  /* noccur = 1 */
                  Result= EID{CFALSE.Id(),0}
                  for (p <= g0226) /* while:9 */{ 
                    var void_try10 EID 
                    _ = void_try10
                    { 
                    var g0229I *ClaireBoolean  
                    var g0229I_try023010 EID 
                    g0229I_try023010 = F_funcall_method2(f,self.At(p-1),x)
                    /* ERROR PROTECTION INSERTED (g0229I-void_try10) */
                    if ErrorIn(g0229I_try023010) {void_try10 = g0229I_try023010
                    } else {
                    g0229I = ToBoolean(OBJ(g0229I_try023010))
                    if (g0229I == CTRUE) /* If:10 */{ 
                      ToArray(self.Id()).NthPut(n,self.At(p-1))
                      n = (n+1)
                      if (p > n) /* If:11 */{ 
                        void_try10 = ToArray(self.Id()).NthPut(p,self.At(n-1)).ToEID()
                        } else {
                        void_try10 = EID{CFALSE.Id(),0}
                        /* If-11 */} 
                      } else {
                      void_try10 = EID{CFALSE.Id(),0}
                      /* If-10 */} 
                    }
                    /* ERROR PROTECTION INSERTED (void_try10-void_try10) */
                    if ErrorIn(void_try10) {Result = void_try10
                    break
                    } else {
                    p = (p+1)
                    }
                    /* while-9 */} 
                  }
                  /* Let-8 */} 
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              ToArray(self.Id()).NthPut(n,x)
              Result = F_quicksort_list(self,f,q,(n-1))
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              Result = F_quicksort_list(self,f,(n+1),m)
              }}
              /* Let-6 */} 
            /* Let-5 */} 
          /* If-4 */} 
        /* Let-3 */} 
      } else {
      Result = EID{CFALSE.Id(),0}
      /* If-2 */} 
    return Result} 
  
// The EID go function for: quicksort @ list (throw: true) 
func E_quicksort_list (self EID,f EID,n EID,m EID) EID { 
    return /*(sm for quicksort @ list= EID)*/ F_quicksort_list(ToList(OBJ(self)),
      ToMethod(OBJ(f)),
      INT(n),
      INT(m) )} 
  
// destructive method that build the powerset
/* {1} OPT.The go function for: build_powerset(self:list) [] */
func F_build_powerset_list (self *ClaireList ) *ClaireSet  { 
    // procedure body with s =  
var Result *ClaireSet  
    if (self.Length() != 0) /* If:2 */{ 
      /* Let:3 */{ 
        var x *ClaireAny   = self.At(1-1)
        /* noccur = 1 */
        /* Let:4 */{ 
          var l1 *ClaireSet   = F_build_powerset_list(self.Skip(1))
          /* noccur = 2 */
          /* Let:5 */{ 
            var l2 *ClaireSet   = l1
            /* noccur = 3 */
            /* For:6 */{ 
              var y *ClaireAny  
              _ = y
              for _,y = range(l1.Values)/* loop:7 */{ 
                l2 = l2.AddFast(F_append_set(MakeConstantSet(x),ToSet(y)).Id())
                /* loop-7 */} 
              /* For-6 */} 
            Result = l2
            /* Let-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      } else {
      Result = MakeConstantSet(CEMPTY.Id())
      /* If-2 */} 
    return Result} 
  
// The EID go function for: build_powerset @ list (throw: false) 
func E_build_powerset_list (self EID) EID { 
    return EID{/*(sm for build_powerset @ list= set)*/ F_build_powerset_list(ToList(OBJ(self)) ).Id(),0}} 
  
// skip 
// new and useful (v3.1.06) - create a list with n replication of the default value
/* {1} OPT.The go function for: make_copy_list(n:integer,d:any) [] */
func F_make_copy_list_integer (n int,d *ClaireAny ) *ClaireList  { 
    // procedure body with s =  
var Result *ClaireList  
    /* Let:2 */{ 
      var l *ClaireList   = F_make_list_integer(n,d)
      /* noccur = 2 */
      if (d.Isa.IsIn(C_list) == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var g0231 *ClaireList   = ToList(d)
          /* noccur = 1 */
          /* Let:5 */{ 
            var i int  = 1
            /* noccur = 4 */
            /* Let:6 */{ 
              var g0232 int  = n
              /* noccur = 1 */
              for (i <= g0232) /* while:7 */{ 
                ToArray(l.Id()).NthPut(i,g0231.Copy().Id())
                i = (i+1)
                /* while-7 */} 
              /* Let-6 */} 
            /* Let-5 */} 
          /* Let-4 */} 
        /* If-3 */} 
      Result = l
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: make_copy_list @ integer (throw: false) 
func E_make_copy_list_integer (n EID,d EID) EID { 
    return EID{/*(sm for make_copy_list @ integer= list)*/ F_make_copy_list_integer(INT(n),ANY(d) ).Id(),0}} 
  
//----------------------  SET  ---------------------------------------------
/* {1} OPT.The go function for: difference(self:set,x:set) [] */
func F_difference_set (self *ClaireSet ,x *ClaireSet ) *ClaireSet  { 
    // procedure body with s =  
var Result *ClaireSet  
    /* Let:2 */{ 
      var y_in *ClaireSet   = self
      /* noccur = 2 */
      /* Let:3 */{ 
        var y_out *ClaireSet   = y_in.Empty()
        /* noccur = 2 */
        /* For:4 */{ 
          var y *ClaireAny  
          _ = y
          for _,y = range(y_in.Values)/* loop:5 */{ 
            if (x.Contain_ask(y) != CTRUE) /* If:6 */{ 
              y_out.AddFast(y)
              /* If-6 */} 
            /* loop-5 */} 
          /* For-4 */} 
        Result = y_out
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: difference @ set (throw: false) 
func E_difference_set (self EID,x EID) EID { 
    return EID{/*(sm for difference @ set= set)*/ F_difference_set(ToSet(OBJ(self)),ToSet(OBJ(x)) ).Id(),0}} 
  
//----------------------  TYPE ---------------------------------------------
//--------- ARRAY --------------------------------------------------------
/* {1} OPT.The go function for: nth=(self:array,x:integer,y:any) [] */
func F_nth_equal_array (self *ClaireList ,x int,y *ClaireAny ) EID { 
    var Result EID 
    if (ToList(self.Id()).Of().Contains(y) != CTRUE) /* If:2 */{ 
      Result = ToException(C_general_error.Make(MakeString("type mismatch for array update ~S, ~S").Id(),MakeConstantList(y,self.Id()).Id())).Close()
      /* If!2 */}  else if ((x > 0) && 
        (x <= self.Length())) /* If:2 */{ 
      Result = self.NthPut(x,y).ToEID()
      } else {
      Result = ToException(C_general_error.Make(MakeString("nth[~S] out of scope for ~S").Id(),MakeConstantList(MakeInteger(x).Id(),self.Id()).Id())).Close()
      /* If-2 */} 
    return Result} 
  
// The EID go function for: nth= @ array (throw: true) 
func E_nth_equal_array (self EID,x EID,y EID) EID { 
    return /*(sm for nth= @ array= EID)*/ F_nth_equal_array(ToArray(OBJ(self)),INT(x),ANY(y) )} 
  
/* {1} OPT.The go function for: self_print(self:array) [] */
func F_self_print_array_Core (self *ClaireList ) EID { 
    var Result EID 
    PRINC("array<")
    Result = F_print_any(ToList(self.Id()).Of().Id())
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(">[")
    F_princ_integer(self.Length())
    PRINC("]")
    Result = EVOID
    }
    return Result} 
  
// The EID go function for: self_print @ array (throw: true) 
func E_self_print_array_Core (self EID) EID { 
    return /*(sm for self_print @ array= EID)*/ F_self_print_array_Core(ToArray(OBJ(self)) )} 
  
//---------------------- CHAR --------------------------------------------
/* {1} OPT.The go function for: self_print(self:char) [] */
func F_self_print_char_Core (self rune)  { 
    // procedure body with s =  
PRINC("'")
    F_princ_char(self)
    PRINC("'")
    } 
  
// The EID go function for: self_print @ char (throw: false) 
func E_self_print_char_Core (self EID) EID { 
    /*(sm for self_print @ char= void)*/ F_self_print_char_Core(CHAR(self) )
    return EVOID} 
  
/* {1} OPT.The go function for: <=(c1:char,c2:char) [] */
func F__inf_equal_char (c1 rune,c2 rune) *ClaireBoolean  { 
    // use function body compiling 
return  F__inf_equal_integer(int(c1),int(c2))
    } 
  
// The EID go function for: <= @ char (throw: false) 
func E__inf_equal_char (c1 EID,c2 EID) EID { 
    return EID{/*(sm for <= @ char= boolean)*/ F__inf_equal_char(CHAR(c1),CHAR(c2) ).Id(),0}} 
  
// --------------------- BOOL -----------------------------------------------
/* {1} OPT.The go function for: random(b:boolean) [] */
func F_random_boolean (b *ClaireBoolean ) *ClaireBoolean  { 
    // use function body compiling 
if (b == CTRUE) /* body If:2 */{ 
      return  F__sup_equal_integer(F_random_integer(10000),5000)
      } else {
      return  CFALSE
      /* body If-2 */} 
    } 
  
// The EID go function for: random @ boolean (throw: false) 
func E_random_boolean (b EID) EID { 
    return EID{/*(sm for random @ boolean= boolean)*/ F_random_boolean(ToBoolean(OBJ(b)) ).Id(),0}} 
  