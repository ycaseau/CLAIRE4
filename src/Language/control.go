/***** CLAIRE Compilation of file /Users/ycaseau/Dropbox/src/clairev4.0/src/meta/control.cl 
         [version 4.0.02 / safety 5] Saturday 10-30-2021 *****/

package Language
import (_ "fmt"
	. "Kernel"
	"Core"
)

//-------- dumb function to prevent import errors --------
func import_g0201() { 
    _ = Core.It
    } 
  
  
//+-------------------------------------------------------------+
//| CLAIRE                                                      |
//| control.cl                                                  |
//| Copyright (C) 1994 - 2021 Yves Caseau. All Rights Reserved  |
//| cf. copyright info in file object.cl: about()               |
//+-------------------------------------------------------------+
// *********************************************************************
// *     Part 1: If, Do, Let                                           *
// *     Part 2: set control structures                                *
// *     Part 3: other control structures                              *
// *     Part 4: the constructs                                        *
// *********************************************************************
// *********************************************************************
// *     Part 1: If, Do, Let                                           *
// *********************************************************************
//--------------- the IF --------------------------------------------
/* {1} OPT.The go function for: self_print(self:If) [] */
func (self *If ) SelfPrint () EID { 
    var Result EID 
    PRINC("(")
    Core.C_pretty.Index = (Core.C_pretty.Index+1)
    Result = self.Printstat()
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    Core.C_pretty.Index = (Core.C_pretty.Index-1)
    PRINC(")")
    Result = EVOID
    }
    return Result} 
  
// The EID go function for: self_print @ If (throw: true) 
func E_self_print_If_Language (self EID) EID { 
    return /*(sm for self_print @ If= EID)*/ To_If(OBJ(self)).SelfPrint( )} 
  
/* {1} OPT.The go function for: printstat(self:If) [] */
func (self *If ) Printstat () EID { 
    var Result EID 
    PRINC("if ")
    Result = F_printexp_any(self.Test,CFALSE)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    Result = F_checkfar_void()
    }
    {
    PRINC(" ")
    Result = F_printif_any(self.Arg)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    Core.C_pretty.Index = (Core.C_pretty.Index-3)
    Result = self.Printelse()
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("")
    Result = EVOID
    }}}
    return Result} 
  
// The EID go function for: printstat @ If (throw: true) 
func E_printstat_If (self EID) EID { 
    return /*(sm for printstat @ If= EID)*/ To_If(OBJ(self)).Printstat( )} 
  
/* {1} OPT.The go function for: printif(self:any) [] */
func F_printif_any (self *ClaireAny ) EID { 
    var Result EID 
    Core.C_pretty.Index = (Core.C_pretty.Index+3)
    if (Core.C_pretty.Pbreak == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var b_index int  = Core.F_buffer_length_void()
        /* noccur = 1 */
        /* Let:4 */{ 
          var _Zl int  = Core.C_pretty.Index
          /* noccur = 1 */
          Core.C_pretty.Pbreak = CFALSE
          h_index := ClEnv.Index /* Handle */
          h_base := ClEnv.Base
          Result = Core.F_CALL(C_print,ARGS(self.ToEID()))
          if ErrorIn(Result) && ToType(Core.C_much_too_far.Id()).Contains(ANY(Result)) == CTRUE { 
            /* s=EID */ClEnv.Index = h_index
            ClEnv.Base = h_base
            Result = EID{CEMPTY.Id(),0}
            } 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          Core.C_pretty.Pbreak = CTRUE
          if (Core.F_short_enough_integer(Core.F_buffer_length_void()) == CTRUE) /* If:5 */{ 
            Result = EID{CEMPTY.Id(),0}
            } else {
            Core.F_buffer_set_length_integer(b_index)
            Core.C_pretty.Index = _Zl
            Result = F_lbreak_void()
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            Result = Core.F_CALL(C_print,ARGS(self.ToEID()))
            }
            /* If-5 */} 
          }
          /* Let-4 */} 
        /* Let-3 */} 
      } else {
      Result = Core.F_CALL(C_print,ARGS(self.ToEID()))
      /* If-2 */} 
    return Result} 
  
// The EID go function for: printif @ any (throw: true) 
func E_printif_any (self EID) EID { 
    return /*(sm for printif @ any= EID)*/ F_printif_any(ANY(self) )} 
  
/* {1} OPT.The go function for: printelse(self:If) [] */
func (self *If ) Printelse () EID { 
    var Result EID 
    /* Let:2 */{ 
      var e *ClaireAny   = self.Other
      /* noccur = 4 */
      if (e.Isa.IsIn(C_If) == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var g0207 *If   = To_If(e)
          /* noccur = 3 */
          PRINC(" ")
          Result = F_lbreak_void()
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC("else if ")
          Result = F_printexp_any(g0207.Test,CFALSE)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(" ")
          Result = F_printif_any(g0207.Arg)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          Core.C_pretty.Index = (Core.C_pretty.Index-3)
          Result = g0207.Printelse()
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC("")
          Result = EVOID
          }}}}
          /* Let-4 */} 
        /* If!3 */}  else if (Equal(e,CNIL.Id()) != CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var _Zl int  = Core.C_pretty.Index
          /* noccur = 1 */
          PRINC(" ")
          Result = F_lbreak_void()
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC("else ")
          F_set_level_integer(1)
          Result = Core.F_CALL(C_print,ARGS(e.ToEID()))
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC("")
          Result = EVOID
          }}
          {
          /* update:5 */{ 
            var va_arg1 *Core.PrettyPrinter  
            var va_arg2 int 
            va_arg1 = Core.C_pretty
            va_arg2 = _Zl
            /* ---------- now we compile update mClaire/index(va_arg1) := va_arg2 ------- */
            va_arg1.Index = va_arg2
            Result = EID{C__INT,IVAL(va_arg2)}
            /* update-5 */} 
          }
          /* Let-4 */} 
        } else {
        Result = EID{CFALSE.Id(),0}
        /* If-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: printelse @ If (throw: true) 
func E_printelse_If (self EID) EID { 
    return /*(sm for printelse @ If= EID)*/ To_If(OBJ(self)).Printelse( )} 
  
// notice that the eval(test) is not a boolean thus the compiler will add
// something
// TODO: check that is is not too slow (may use a constant for _oid_(true))
/* {1} OPT.The go function for: self_eval(self:If) [] */
func (self *If ) SelfEval () EID { 
    var Result EID 
    /* Let:2 */{ 
      var x *ClaireAny  
      /* noccur = 3 */
      var x_try02093 EID 
      x_try02093 = EVAL(self.Test)
      /* ERROR PROTECTION INSERTED (x-Result) */
      if ErrorIn(x_try02093) {Result = x_try02093
      } else {
      x = ANY(x_try02093)
      if (x == CTRUE.Id()) /* If:3 */{ 
        Result = EVAL(self.Arg)
        /* If!3 */}  else if (x == CFALSE.Id()) /* If:3 */{ 
        Result = EVAL(self.Other)
        /* If!3 */}  else if (F_boolean_I_any(x) == CTRUE) /* If:3 */{ 
        Result = EVAL(self.Arg)
        } else {
        Result = EVAL(self.Other)
        /* If-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: self_eval @ If (throw: true) 
func E_self_eval_If (self EID) EID { 
    return /*(sm for self_eval @ If= EID)*/ To_If(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: If 
func EVAL_If (x *ClaireAny) EID { 
     return To_If(x).SelfEval()} 
  
//--------------------- block structure------------------------------
/* {1} OPT.The go function for: self_print(self:Do) [] */
func (self *Do ) SelfPrint () EID { 
    var Result EID 
    /* Let:2 */{ 
      var _Zl int  = Core.C_pretty.Index
      /* noccur = 1 */
      PRINC("(")
      F_set_level_integer(1)
      Result = F_printdo_list(self.Args,CTRUE)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      /* update:3 */{ 
        var va_arg1 *Core.PrettyPrinter  
        var va_arg2 int 
        va_arg1 = Core.C_pretty
        va_arg2 = _Zl
        /* ---------- now we compile update mClaire/index(va_arg1) := va_arg2 ------- */
        va_arg1.Index = va_arg2
        Result = EID{C__INT,IVAL(va_arg2)}
        /* update-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: self_print @ Do (throw: true) 
func E_self_print_Do_Language (self EID) EID { 
    return /*(sm for self_print @ Do= EID)*/ To_Do(OBJ(self)).SelfPrint( )} 
  
/* {1} OPT.The go function for: printdo(l:list,clo:boolean) [] */
func F_printdo_list (l *ClaireList ,clo *ClaireBoolean ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var n int  = l.Length()
      /* noccur = 3 */
      /* For:3 */{ 
        var x *ClaireAny  
        _ = x
        Result= EID{CFALSE.Id(),0}
        var x_support *ClaireList  
        x_support = l
        x_len := x_support.Length()
        for i_it := 0; i_it < x_len; i_it++ { 
          x = x_support.At(i_it)
          var void_try5 EID 
          _ = void_try5
          { 
          if (x.Isa.IsIn(C_If) == CTRUE) /* If:5 */{ 
            /* Let:6 */{ 
              var g0210 *If   = To_If(x)
              /* noccur = 1 */
              void_try5 = g0210.Printstat()
              /* Let-6 */} 
            } else {
            void_try5 = Core.F_CALL(C_print,ARGS(x.ToEID()))
            /* If-5 */} 
          /* ERROR PROTECTION INSERTED (void_try5-void_try5) */
          if ErrorIn(void_try5) {Result = void_try5
          break
          } else {
          n = (n-1)
          if (n == 0) /* If:5 */{ 
            if (clo == CTRUE) /* If:6 */{ 
              PRINC(")")
              void_try5 = EVOID
              } else {
              void_try5 = EID{CFALSE.Id(),0}
              /* If-6 */} 
            } else {
            PRINC(", ")
            void_try5 = F_lbreak_void()
            /* ERROR PROTECTION INSERTED (void_try5-void_try5) */
            if ErrorIn(void_try5) {Result = void_try5
            break
            } else {
            }
            /* If-5 */} 
          /* ERROR PROTECTION INSERTED (void_try5-void_try5) */
          if ErrorIn(void_try5) {Result = void_try5
          break
          } else {
          }}
          }
          /* loop-4 */} 
        /* For-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: printdo @ list (throw: true) 
func E_printdo_list (l EID,clo EID) EID { 
    return /*(sm for printdo @ list= EID)*/ F_printdo_list(ToList(OBJ(l)),ToBoolean(OBJ(clo)) )} 
  
/* {1} OPT.The go function for: printblock(x:any) [] */
func F_printblock_any (x *ClaireAny ) EID { 
    var Result EID 
    if (x.Isa.IsIn(C_Do) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0212 *Do   = To_Do(x)
        /* noccur = 1 */
        Result = F_printdo_list(g0212.Args,CFALSE)
        /* Let-3 */} 
      /* If!2 */}  else if (x.Isa.IsIn(C_If) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0213 *If   = To_If(x)
        /* noccur = 1 */
        Result = g0213.Printstat()
        /* Let-3 */} 
      } else {
      Result = Core.F_CALL(C_print,ARGS(x.ToEID()))
      /* If-2 */} 
    return Result} 
  
// The EID go function for: printblock @ any (throw: true) 
func E_printblock_any (x EID) EID { 
    return /*(sm for printblock @ any= EID)*/ F_printblock_any(ANY(x) )} 
  
// use res:EID pragma when compiled with CLAIRE4, res:any for CLAIRE3
/* {1} OPT.The go function for: self_eval(self:Do) [] */
func (self *Do ) SelfEval () EID { 
    var Result EID 
    /* Let:2 */{ 
      var res *ClaireAny   = CEMPTY.Id()
      /* noccur = 2 */
      /* For:3 */{ 
        var _Zx *ClaireAny  
        _ = _Zx
        Result= EID{CFALSE.Id(),0}
        var _Zx_support *ClaireList  
        _Zx_support = self.Args
        _Zx_len := _Zx_support.Length()
        for i_it := 0; i_it < _Zx_len; i_it++ { 
          _Zx = _Zx_support.At(i_it)
          var void_try5 EID 
          _ = void_try5
          var res_try02155 EID 
          res_try02155 = EVAL(_Zx)
          /* ERROR PROTECTION INSERTED (res-Result) */
          if ErrorIn(res_try02155) {Result = res_try02155
          Result = res_try02155
          break
          } else {
          res = ANY(res_try02155)
          void_try5 = res.ToEID()
          }
          /* loop-4 */} 
        /* For-3 */} 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = res.ToEID()
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: self_eval @ Do (throw: true) 
func E_self_eval_Do (self EID) EID { 
    return /*(sm for self_eval @ Do= EID)*/ To_Do(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Do 
func EVAL_Do (x *ClaireAny) EID { 
     return To_Do(x).SelfEval()} 
  
// ----------------- lexical variable definition -----------------------
/* {1} OPT.The go function for: self_print(self:Let) [] */
func (self *Let ) SelfPrint () EID { 
    var Result EID 
    /* Let:2 */{ 
      var _Zl int  = Core.C_pretty.Index
      /* noccur = 1 */
      F_set_level_integer(1)
      PRINC("let ")
      Result = F_ppvariable_Variable(self.ClaireVar)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(" := ")
      Result = F_printexp_any(self.Value,CFALSE)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = self.Printbody()
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC("")
      Result = EVOID
      }}}
      {
      /* update:3 */{ 
        var va_arg1 *Core.PrettyPrinter  
        var va_arg2 int 
        va_arg1 = Core.C_pretty
        va_arg2 = _Zl
        /* ---------- now we compile update mClaire/index(va_arg1) := va_arg2 ------- */
        va_arg1.Index = va_arg2
        Result = EID{C__INT,IVAL(va_arg2)}
        /* update-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: self_print @ Let (throw: true) 
func E_self_print_Let_Language (self EID) EID { 
    return /*(sm for self_print @ Let= EID)*/ To_Let(OBJ(self)).SelfPrint( )} 
  
/* {1} OPT.The go function for: printbody(self:Let) [] */
func (self *Let ) Printbody () EID { 
    var Result EID 
    /* Let:2 */{ 
      var a *ClaireAny   = self.Arg
      /* noccur = 3 */
      if (a.Isa.IsIn(C_Let) == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var g0217 *Let   = To_Let(a)
          /* noccur = 3 */
          PRINC(",")
          Result = F_lbreak_integer(4)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          Result = F_ppvariable_Variable(g0217.ClaireVar)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(" := ")
          Result = F_printexp_any(g0217.Value,CFALSE)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          Core.C_pretty.Index = (Core.C_pretty.Index-4)
          Result = g0217.Printbody()
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC("")
          Result = EVOID
          }}}}
          /* Let-4 */} 
        } else {
        PRINC(" in ")
        Result = F_lbreak_integer(2)
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        Result = Core.F_CALL(C_print,ARGS(a.ToEID()))
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC("")
        Result = EVOID
        }}
        /* If-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: printbody @ Let (throw: true) 
func E_printbody_Let (self EID) EID { 
    return /*(sm for printbody @ Let= EID)*/ To_Let(OBJ(self)).Printbody( )} 
  
/* {1} OPT.The go function for: self_eval(self:Let) [] */
func (self *Let ) SelfEval () EID { 
    var Result EID 
    /* Let:2 */{ 
      var val *ClaireAny  
      /* noccur = 1 */
      var val_try02193 EID 
      val_try02193 = EVAL(self.Value)
      /* ERROR PROTECTION INSERTED (val-Result) */
      if ErrorIn(val_try02193) {Result = val_try02193
      } else {
      val = ANY(val_try02193)
      Result = F_write_value_Variable(self.ClaireVar,val)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = EVAL(self.Arg)
      }
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: self_eval @ Let (throw: true) 
func E_self_eval_Let (self EID) EID { 
    return /*(sm for self_eval @ Let= EID)*/ To_Let(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Let 
func EVAL_Let (x *ClaireAny) EID { 
     return To_Let(x).SelfEval()} 
  
// a when is a special Let that filters out the unknown value !
//
/* {1} OPT.The go function for: self_print(self:When) [] */
func (self *When ) SelfPrint () EID { 
    var Result EID 
    /* Let:2 */{ 
      var _Zl int  = Core.C_pretty.Index
      /* noccur = 1 */
      F_set_level_integer(1)
      PRINC("when ")
      Result = F_ppvariable_Variable(self.ClaireVar)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(" := ")
      Result = F_printexp_any(self.Value,CFALSE)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(" in ")
      Result = F_lbreak_integer(2)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = Core.F_CALL(C_print,ARGS(self.Arg.ToEID()))
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC("")
      Result = EVOID
      }}}}
      {
      if (self.Other != CNULL) /* If:3 */{ 
        PRINC(" ")
        Result = F_lbreak_void()
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC("else ")
        F_set_level_integer(1)
        Result = Core.F_CALL(C_print,ARGS(self.Other.ToEID()))
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC("")
        Result = EVOID
        }}
        } else {
        Result = EID{CFALSE.Id(),0}
        /* If-3 */} 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      /* update:3 */{ 
        var va_arg1 *Core.PrettyPrinter  
        var va_arg2 int 
        va_arg1 = Core.C_pretty
        va_arg2 = _Zl
        /* ---------- now we compile update mClaire/index(va_arg1) := va_arg2 ------- */
        va_arg1.Index = va_arg2
        Result = EID{C__INT,IVAL(va_arg2)}
        /* update-3 */} 
      }}
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: self_print @ When (throw: true) 
func E_self_print_When_Language (self EID) EID { 
    return /*(sm for self_print @ When= EID)*/ To_When(OBJ(self)).SelfPrint( )} 
  
/* {1} OPT.The go function for: self_eval(self:When) [] */
func (self *When ) SelfEval () EID { 
    var Result EID 
    /* Let:2 */{ 
      var val *ClaireAny  
      /* noccur = 2 */
      var val_try02203 EID 
      val_try02203 = EVAL(self.Value)
      /* ERROR PROTECTION INSERTED (val-Result) */
      if ErrorIn(val_try02203) {Result = val_try02203
      } else {
      val = ANY(val_try02203)
      /* Let:3 */{ 
        var n int  = ClEnv.Trace_I
        /* noccur = 0 */
        _ = n
        if (val != CNULL) /* If:4 */{ 
          Result = F_write_value_Variable(self.ClaireVar,val)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          Result = EVAL(self.Arg)
          }
          } else {
          Result = EVAL(self.Other)
          /* If-4 */} 
        /* Let-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: self_eval @ When (throw: true) 
func E_self_eval_When (self EID) EID { 
    return /*(sm for self_eval @ When= EID)*/ To_When(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: When 
func EVAL_When (x *ClaireAny) EID { 
     return To_When(x).SelfEval()} 
  
// two special forms of Let:
// Let+(v,r(x),(r(x) := y),Let(v2,e,(r(x) := v,v2)))    <=>  let r(x) = y in e
// Let*(v,f(),Let(v1,v[1],...(Let(vn,v[n],e))   <=> let (v1,v2,...vn) := f() in e
//
//note: the Let* is also used for multi-assignments
// Let*(v,f(),(v1 := v[1], v2 := v[2], ...))   <=>  (v1,v2,...vn) := f()
//
/* {1} OPT.The go function for: self_print(self:Let+) [] */
func (self *Let_plus ) SelfPrint () EID { 
    var Result EID 
    /* Let:2 */{ 
      var _Zl int  = Core.C_pretty.Index
      /* noccur = 1 */
      /* Let:3 */{ 
        var l *ClaireList   = To_Do(self.Arg).Args
        /* noccur = 2 */
        F_set_level_integer(1)
        PRINC("let ")
        Result = F_printexp_any(self.Value,CFALSE)
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC(" := ")
        Result = F_printexp_any(ToList(OBJ(Core.F_CALL(C_args,ARGS(l.At(1-1).ToEID())))).At(3-1),CFALSE)
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC(" in ")
        Result = F_lbreak_integer(2)
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        Result = Core.F_CALL(C_print,ARGS(To_Let(l.At(2-1)).Value.ToEID()))
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC("")
        Result = EVOID
        }}}}
        {
        /* update:4 */{ 
          var va_arg1 *Core.PrettyPrinter  
          var va_arg2 int 
          va_arg1 = Core.C_pretty
          va_arg2 = _Zl
          /* ---------- now we compile update mClaire/index(va_arg1) := va_arg2 ------- */
          va_arg1.Index = va_arg2
          Result = EID{C__INT,IVAL(va_arg2)}
          /* update-4 */} 
        }
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: self_print @ Let+ (throw: true) 
func E_self_print_Let_plus_Language (self EID) EID { 
    return /*(sm for self_print @ Let+= EID)*/ To_Let_plus(OBJ(self)).SelfPrint( )} 
  
/* {1} OPT.The go function for: self_print(self:Let*) [] */
func (self *Let_star ) SelfPrint () EID { 
    var Result EID 
    /* Let:2 */{ 
      var _Zl int  = Core.C_pretty.Index
      /* noccur = 1 */
      /* Let:3 */{ 
        var l *ClaireAny   = self.Arg
        /* noccur = 6 */
        F_set_level_integer(1)
        if (l.Isa.IsIn(C_Let) == CTRUE) /* If:4 */{ 
          PRINC("let (")
          Result= EID{CFALSE.Id(),0}
          for (CTRUE == CTRUE) /* while:5 */{ 
            var void_try6 EID 
            _ = void_try6
            { 
            void_try6 = Core.F_CALL(C_Language_ppvariable,ARGS(Core.F_CALL(C_var,ARGS(l.ToEID()))))
            /* ERROR PROTECTION INSERTED (void_try6-void_try6) */
            if ErrorIn(void_try6) {Result = void_try6
            break
            } else {
            /* Let:6 */{ 
              var lnext *ClaireAny   = ANY(Core.F_CALL(C_arg,ARGS(l.ToEID())))
              /* noccur = 3 */
              var g0222I *ClaireBoolean  
              if (lnext.Isa.IsIn(C_Let) == CTRUE) /* If:7 */{ 
                /* Let:8 */{ 
                  var g0221 *Let   = To_Let(lnext)
                  /* noccur = 2 */
                  g0222I = MakeBoolean((g0221.Value.Isa.IsIn(C_Call) == CTRUE) && (ToList(OBJ(Core.F_CALL(C_args,ARGS(g0221.Value.ToEID())))).At(1-1) == self.ClaireVar.Id()))
                  /* Let-8 */} 
                } else {
                g0222I = CFALSE
                /* If-7 */} 
              if (g0222I == CTRUE) /* If:7 */{ 
                PRINC(",")
                l = lnext
                } else {
                 /*v = Result, s =EID*/
Result = EID{CTRUE.Id(),0}
                break
                /* If-7 */} 
              /* Let-6 */} 
            }
            /* while-5 */} 
          }
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(") := ")
          Result = F_printexp_any(self.Value,CFALSE)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          Result = To_Let(l).Printbody()
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC("")
          Result = EVOID
          }}}
          } else {
          PRINC("(")
          /* Let:5 */{ 
            var _Zf *ClaireBoolean   = CTRUE
            /* noccur = 2 */
            /* For:6 */{ 
              var _Za *ClaireAny  
              _ = _Za
              Result= EID{CFALSE.Id(),0}
              var _Za_support *ClaireList  
              _Za_support = ToList(OBJ(Core.F_CALL(C_args,ARGS(l.ToEID()))))
              _Za_len := _Za_support.Length()
              for i_it := 0; i_it < _Za_len; i_it++ { 
                _Za = _Za_support.At(i_it)
                var void_try8 EID 
                _ = void_try8
                { 
                if (_Zf == CTRUE) /* If:8 */{ 
                  _Zf = CFALSE
                  } else {
                  PRINC(",")
                  /* If-8 */} 
                void_try8 = Core.F_CALL(C_Language_ppvariable,ARGS(Core.F_CALL(C_var,ARGS(_Za.ToEID()))))
                /* ERROR PROTECTION INSERTED (void_try8-void_try8) */
                if ErrorIn(void_try8) {Result = void_try8
                break
                } else {
                }
                }
                /* loop-7 */} 
              /* For-6 */} 
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(") := ")
          Result = F_printexp_any(self.Value,CFALSE)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC("")
          Result = EVOID
          }}
          /* If-4 */} 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        /* update:4 */{ 
          var va_arg1 *Core.PrettyPrinter  
          var va_arg2 int 
          va_arg1 = Core.C_pretty
          va_arg2 = _Zl
          /* ---------- now we compile update mClaire/index(va_arg1) := va_arg2 ------- */
          va_arg1.Index = va_arg2
          Result = EID{C__INT,IVAL(va_arg2)}
          /* update-4 */} 
        }
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: self_print @ Let* (throw: true) 
func E_self_print_Let_star_Language (self EID) EID { 
    return /*(sm for self_print @ Let*= EID)*/ To_Let_star(OBJ(self)).SelfPrint( )} 
  
// *********************************************************************
// *     Part 2: set control structures                                *
// *********************************************************************
// for is the simplest evaluation loop
//
/* {1} OPT.The go function for: self_print(self:For) [] */
func (self *For ) SelfPrint () EID { 
    var Result EID 
    PRINC("for ")
    Result = F_ppvariable_Variable(self.ClaireVar)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(" in ")
    /* Let:2 */{ 
      var _Zl int  = Core.C_pretty.Index
      /* noccur = 1 */
      F_set_level_void()
      Result = F_printexp_any(self.SetArg,CFALSE)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      /* update:3 */{ 
        var va_arg1 *Core.PrettyPrinter  
        var va_arg2 int 
        va_arg1 = Core.C_pretty
        va_arg2 = _Zl
        /* ---------- now we compile update mClaire/index(va_arg1) := va_arg2 ------- */
        va_arg1.Index = va_arg2
        Result = EID{C__INT,IVAL(va_arg2)}
        /* update-3 */} 
      }
      /* Let-2 */} 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(" ")
    Core.C_pretty.Index = (Core.C_pretty.Index+2)
    Result = F_lbreak_void()
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    Result = Core.F_CALL(C_print,ARGS(self.Arg.ToEID()))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    /* update:2 */{ 
      var va_arg1 *Core.PrettyPrinter  
      var va_arg2 int 
      va_arg1 = Core.C_pretty
      va_arg2 = (Core.C_pretty.Index-2)
      /* ---------- now we compile update mClaire/index(va_arg1) := va_arg2 ------- */
      va_arg1.Index = va_arg2
      Result = EID{C__INT,IVAL(va_arg2)}
      /* update-2 */} 
    }}
    {
    PRINC("")
    Result = EVOID
    }}}
    return Result} 
  
// The EID go function for: self_print @ For (throw: true) 
func E_self_print_For_Language (self EID) EID { 
    return /*(sm for self_print @ For= EID)*/ To_For(OBJ(self)).SelfPrint( )} 
  
/* {1} OPT.The go function for: self_eval(self:For) [] */
func (self *For ) SelfEval () EID { 
    var Result EID 
    /* Let:2 */{ 
      var x *ClaireAny  
      /* noccur = 11 */
      var x_try02343 EID 
      x_try02343 = EVAL(self.SetArg)
      /* ERROR PROTECTION INSERTED (x-Result) */
      if ErrorIn(x_try02343) {Result = x_try02343
      } else {
      x = ANY(x_try02343)
      h_index := ClEnv.Index /* Handle */
      h_base := ClEnv.Base
      if (C_class.Id() == x.Isa.Id()) /* If:3 */{ 
        /* Let:4 */{ 
          var g0225 *ClaireClass   = ToClass(x)
          /* noccur = 1 */
          /* For:5 */{ 
            var y *ClaireAny  
            _ = y
            Result= EID{CFALSE.Id(),0}
            for _,y = range(g0225.Descendents.Values)/* loop:6 */{ 
              var void_try7 EID 
              _ = void_try7
              /* For:7 */{ 
                var z *ClaireAny  
                _ = z
                void_try7= EID{CFALSE.Id(),0}
                var z_support *ClaireList  
                z_support = ToClass(y).Instances
                z_len := z_support.Length()
                for i_it := 0; i_it < z_len; i_it++ { 
                  z = z_support.At(i_it)
                  var void_try9 EID 
                  _ = void_try9
                  { 
                  void_try9 = F_write_value_Variable(self.ClaireVar,z)
                  /* ERROR PROTECTION INSERTED (void_try9-void_try9) */
                  if ErrorIn(void_try9) {void_try7 = void_try9
                  break
                  } else {
                  void_try9 = EVAL(self.Arg)
                  /* ERROR PROTECTION INSERTED (void_try9-void_try9) */
                  if ErrorIn(void_try9) {void_try7 = void_try9
                  break
                  } else {
                  }}
                  }
                  /* loop-8 */} 
                /* For-7 */} 
              /* ERROR PROTECTION INSERTED (void_try7-Result) */
              if ErrorIn(void_try7) {Result = void_try7
              Result = void_try7
              break
              } else {
              }
              /* loop-6 */} 
            /* For-5 */} 
          /* Let-4 */} 
        /* If!3 */}  else if (x.Isa.IsIn(C_list) == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var g0226 *ClaireList   = ToList(x)
          /* noccur = 1 */
          /* For:5 */{ 
            var z *ClaireAny  
            _ = z
            Result= EID{CFALSE.Id(),0}
            var z_support *ClaireList  
            z_support = g0226
            z_len := z_support.Length()
            for i_it := 0; i_it < z_len; i_it++ { 
              z = z_support.At(i_it)
              var void_try7 EID 
              _ = void_try7
              { 
              void_try7 = F_write_value_Variable(self.ClaireVar,z)
              /* ERROR PROTECTION INSERTED (void_try7-void_try7) */
              if ErrorIn(void_try7) {Result = void_try7
              break
              } else {
              void_try7 = EVAL(self.Arg)
              /* ERROR PROTECTION INSERTED (void_try7-void_try7) */
              if ErrorIn(void_try7) {Result = void_try7
              break
              } else {
              }}
              }
              /* loop-6 */} 
            /* For-5 */} 
          /* Let-4 */} 
        /* If!3 */}  else if (x.Isa.IsIn(C_array) == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var g0227 *ClaireList   = ToArray(x)
          /* noccur = 2 */
          /* Let:5 */{ 
            var n int  = g0227.Length()
            /* noccur = 1 */
            /* Let:6 */{ 
              var g0228 int  = 1
              /* noccur = 4 */
              /* Let:7 */{ 
                var g0229 int  = n
                /* noccur = 1 */
                Result= EID{CFALSE.Id(),0}
                for (g0228 <= g0229) /* while:8 */{ 
                  var void_try9 EID 
                  _ = void_try9
                  { 
                  /* Let:9 */{ 
                    var z *ClaireAny   = ToList(g0227.Id()).At(g0228-1)
                    /* noccur = 1 */
                    void_try9 = F_write_value_Variable(self.ClaireVar,z)
                    /* ERROR PROTECTION INSERTED (void_try9-void_try9) */
                    if ErrorIn(void_try9) {Result = void_try9
                    break
                    } else {
                    void_try9 = EVAL(self.Arg)
                    /* ERROR PROTECTION INSERTED (void_try9-void_try9) */
                    if ErrorIn(void_try9) {Result = void_try9
                    break
                    } else {
                    }}
                    /* Let-9 */} 
                  /* ERROR PROTECTION INSERTED (void_try9-void_try9) */
                  if ErrorIn(void_try9) {Result = void_try9
                  break
                  } else {
                  g0228 = (g0228+1)
                  }
                  /* while-8 */} 
                }
                /* Let-7 */} 
              /* Let-6 */} 
            /* Let-5 */} 
          /* Let-4 */} 
        /* If!3 */}  else if (x.Isa.IsIn(C_Interval) == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var g0230 *ClaireInterval   = To_Interval(x)
          /* noccur = 2 */
          /* Let:5 */{ 
            var y int  = g0230.Arg1
            /* noccur = 4 */
            /* Let:6 */{ 
              var g0231 int  = g0230.Arg2
              /* noccur = 1 */
              Result= EID{CFALSE.Id(),0}
              for (y <= g0231) /* while:7 */{ 
                var void_try8 EID 
                _ = void_try8
                { 
                void_try8 = self.ClaireVar.WriteEID(EID{C__INT,IVAL(y)})
                /* ERROR PROTECTION INSERTED (void_try8-void_try8) */
                if ErrorIn(void_try8) {Result = void_try8
                break
                } else {
                void_try8 = EVAL(self.Arg)
                /* ERROR PROTECTION INSERTED (void_try8-void_try8) */
                if ErrorIn(void_try8) {Result = void_try8
                break
                } else {
                }}
                {
                y = (y+1)
                }
                /* while-7 */} 
              }
              /* Let-6 */} 
            /* Let-5 */} 
          /* Let-4 */} 
        /* If!3 */}  else if (x.Isa.IsIn(C_collection) == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var g0232 *ClaireCollection   = ToCollection(x)
          /* noccur = 1 */
          /* For:5 */{ 
            var y *ClaireAny  
            _ = y
            Result= EID{CFALSE.Id(),0}
            var y_support *ClaireList  
            var y_support_try02356 EID 
            y_support_try02356 = Core.F_enumerate_any(g0232.Id())
            /* ERROR PROTECTION INSERTED (y_support-Result) */
            if ErrorIn(y_support_try02356) {Result = y_support_try02356
            } else {
            y_support = ToList(OBJ(y_support_try02356))
            y_len := y_support.Length()
            for i_it := 0; i_it < y_len; i_it++ { 
              y = y_support.At(i_it)
              var void_try7 EID 
              _ = void_try7
              { 
              void_try7 = F_write_value_Variable(self.ClaireVar,y)
              /* ERROR PROTECTION INSERTED (void_try7-void_try7) */
              if ErrorIn(void_try7) {Result = void_try7
              break
              } else {
              void_try7 = EVAL(self.Arg)
              /* ERROR PROTECTION INSERTED (void_try7-void_try7) */
              if ErrorIn(void_try7) {Result = void_try7
              break
              } else {
              }}
              }}
              /* loop-6 */} 
            /* For-5 */} 
          /* Let-4 */} 
        } else {
        Result = ToException(Core.C_general_error.Make(MakeString("[136] ~S is not a collection !").Id(),MakeConstantList(x).Id())).Close()
        /* If-3 */} 
      if ErrorIn(Result) && ToType(Core.C_return_error.Id()).Contains(ANY(Result)) == CTRUE { 
        /* s=EID */ClEnv.Index = h_index
        ClEnv.Base = h_base
        Result = Core.F_CALL(C_arg,ARGS(EID{ClEnv.Exception_I.Id(),0}))
        } 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: self_eval @ For (throw: true) 
func E_self_eval_For (self EID) EID { 
    return /*(sm for self_eval @ For= EID)*/ To_For(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: For 
func EVAL_For (x *ClaireAny) EID { 
     return To_For(x).SelfEval()} 
  
// [collect VAR in SET_EXPR, ...] is the same as a "for", but returns the list of values
//
/* {1} OPT.The go function for: self_print(self:Collect) [] */
func (self *Collect ) SelfPrint () EID { 
    var Result EID 
    PRINC("list{ ")
    Core.C_pretty.Index = (Core.C_pretty.Index+2)
    Result = F_printexp_any(self.Arg,CFALSE)
    {
    PRINC(" | ")
    Result = F_lbreak_void()
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    Result = F_ppvariable_Variable(self.ClaireVar)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(" in ")
    /* Let:2 */{ 
      var _Zl int  = Core.C_pretty.Index
      /* noccur = 1 */
      F_set_level_void()
      Result = F_printexp_any(self.SetArg,CFALSE)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      /* update:3 */{ 
        var va_arg1 *Core.PrettyPrinter  
        var va_arg2 int 
        va_arg1 = Core.C_pretty
        va_arg2 = (_Zl-2)
        /* ---------- now we compile update mClaire/index(va_arg1) := va_arg2 ------- */
        va_arg1.Index = va_arg2
        Result = EID{C__INT,IVAL(va_arg2)}
        /* update-3 */} 
      }
      /* Let-2 */} 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("}")
    Result = EVOID
    }}}}
    return Result} 
  
// The EID go function for: self_print @ Collect (throw: true) 
func E_self_print_Collect_Language (self EID) EID { 
    return /*(sm for self_print @ Collect= EID)*/ To_Collect(OBJ(self)).SelfPrint( )} 
  
// list image : preserve the order for lists and intervals (v4)
/* {1} OPT.The go function for: self_eval(self:Collect) [] */
func (self *Collect ) SelfEval () EID { 
    var Result EID 
    /* Let:2 */{ 
      var x *ClaireAny  
      /* noccur = 7 */
      var x_try02423 EID 
      x_try02423 = EVAL(self.SetArg)
      /* ERROR PROTECTION INSERTED (x-Result) */
      if ErrorIn(x_try02423) {Result = x_try02423
      } else {
      x = ANY(x_try02423)
      /* Let:3 */{ 
        var res *ClaireList  
        /* noccur = 9 */
        /* Let:4 */{ 
          var g0243UU *ClaireType  
          /* noccur = 1 */
          if (self.Of.Id() != CNULL) /* If:5 */{ 
            g0243UU = self.Of
            } else {
            g0243UU = ToType(CEMPTY.Id())
            /* If-5 */} 
          res = g0243UU.EmptyList()
          /* Let-4 */} 
        if (C_class.Id() == x.Isa.Id()) /* If:4 */{ 
          /* Let:5 */{ 
            var g0237 *ClaireClass   = ToClass(x)
            /* noccur = 1 */
            /* For:6 */{ 
              var y *ClaireAny  
              _ = y
              Result= EID{CFALSE.Id(),0}
              for _,y = range(g0237.Descendents.Values)/* loop:7 */{ 
                var void_try8 EID 
                _ = void_try8
                /* For:8 */{ 
                  var z *ClaireAny  
                  _ = z
                  void_try8= EID{CFALSE.Id(),0}
                  var z_support *ClaireList  
                  z_support = ToClass(y).Instances
                  z_len := z_support.Length()
                  for i_it := 0; i_it < z_len; i_it++ { 
                    z = z_support.At(i_it)
                    var void_try10 EID 
                    _ = void_try10
                    { 
                    void_try10 = F_write_value_Variable(self.ClaireVar,z)
                    /* ERROR PROTECTION INSERTED (void_try10-void_try10) */
                    if ErrorIn(void_try10) {void_try8 = void_try10
                    break
                    } else {
                    var res_try024410 EID 
                    /* Let:10 */{ 
                      var g0245UU *ClaireAny  
                      /* noccur = 1 */
                      var g0245UU_try024611 EID 
                      g0245UU_try024611 = EVAL(self.Arg)
                      /* ERROR PROTECTION INSERTED (g0245UU-res_try024410) */
                      if ErrorIn(g0245UU_try024611) {res_try024410 = g0245UU_try024611
                      } else {
                      g0245UU = ANY(g0245UU_try024611)
                      res_try024410 = EID{res.AddFast(g0245UU).Id(),0}
                      }
                      /* Let-10 */} 
                    /* ERROR PROTECTION INSERTED (res-void_try10) */
                    if ErrorIn(res_try024410) {void_try10 = res_try024410
                    void_try8 = res_try024410
                    break
                    } else {
                    res = ToList(OBJ(res_try024410))
                    void_try10 = EID{res.Id(),0}
                    }}
                    }
                    /* loop-9 */} 
                  /* For-8 */} 
                /* ERROR PROTECTION INSERTED (void_try8-Result) */
                if ErrorIn(void_try8) {Result = void_try8
                Result = void_try8
                break
                } else {
                }
                /* loop-7 */} 
              /* For-6 */} 
            /* Let-5 */} 
          /* If!4 */}  else if (x.Isa.IsIn(C_list) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var g0238 *ClaireList   = ToList(x)
            /* noccur = 1 */
            /* For:6 */{ 
              var y *ClaireAny  
              _ = y
              Result= EID{CFALSE.Id(),0}
              var y_support *ClaireList  
              y_support = g0238
              y_len := y_support.Length()
              for i_it := 0; i_it < y_len; i_it++ { 
                y = y_support.At(i_it)
                var void_try8 EID 
                _ = void_try8
                { 
                void_try8 = F_write_value_Variable(self.ClaireVar,y)
                /* ERROR PROTECTION INSERTED (void_try8-void_try8) */
                if ErrorIn(void_try8) {Result = void_try8
                break
                } else {
                var res_try02478 EID 
                /* Let:8 */{ 
                  var g0248UU *ClaireAny  
                  /* noccur = 1 */
                  var g0248UU_try02499 EID 
                  g0248UU_try02499 = EVAL(self.Arg)
                  /* ERROR PROTECTION INSERTED (g0248UU-res_try02478) */
                  if ErrorIn(g0248UU_try02499) {res_try02478 = g0248UU_try02499
                  } else {
                  g0248UU = ANY(g0248UU_try02499)
                  res_try02478 = EID{res.AddFast(g0248UU).Id(),0}
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (res-void_try8) */
                if ErrorIn(res_try02478) {void_try8 = res_try02478
                Result = res_try02478
                break
                } else {
                res = ToList(OBJ(res_try02478))
                void_try8 = EID{res.Id(),0}
                }}
                }
                /* loop-7 */} 
              /* For-6 */} 
            /* Let-5 */} 
          /* If!4 */}  else if (x.Isa.IsIn(C_Interval) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var g0239 *ClaireInterval   = To_Interval(x)
            /* noccur = 2 */
            /* Let:6 */{ 
              var y int  = g0239.Arg1
              /* noccur = 4 */
              /* Let:7 */{ 
                var g0240 int  = g0239.Arg2
                /* noccur = 1 */
                Result= EID{CFALSE.Id(),0}
                for (y <= g0240) /* while:8 */{ 
                  var void_try9 EID 
                  _ = void_try9
                  { 
                  void_try9 = self.ClaireVar.WriteEID(EID{C__INT,IVAL(y)})
                  /* ERROR PROTECTION INSERTED (void_try9-void_try9) */
                  if ErrorIn(void_try9) {Result = void_try9
                  break
                  } else {
                  var res_try02509 EID 
                  /* Let:9 */{ 
                    var g0251UU *ClaireAny  
                    /* noccur = 1 */
                    var g0251UU_try025210 EID 
                    g0251UU_try025210 = EVAL(self.Arg)
                    /* ERROR PROTECTION INSERTED (g0251UU-res_try02509) */
                    if ErrorIn(g0251UU_try025210) {res_try02509 = g0251UU_try025210
                    } else {
                    g0251UU = ANY(g0251UU_try025210)
                    res_try02509 = EID{res.AddFast(g0251UU).Id(),0}
                    }
                    /* Let-9 */} 
                  /* ERROR PROTECTION INSERTED (res-void_try9) */
                  if ErrorIn(res_try02509) {void_try9 = res_try02509
                  Result = res_try02509
                  break
                  } else {
                  res = ToList(OBJ(res_try02509))
                  void_try9 = EID{res.Id(),0}
                  }}
                  {
                  y = (y+1)
                  }
                  /* while-8 */} 
                }
                /* Let-7 */} 
              /* Let-6 */} 
            /* Let-5 */} 
          } else {
          /* For:5 */{ 
            var y *ClaireAny  
            _ = y
            Result= EID{CFALSE.Id(),0}
            var y_support *ClaireList  
            var y_support_try02536 EID 
            y_support_try02536 = Core.F_enumerate_any(x)
            /* ERROR PROTECTION INSERTED (y_support-Result) */
            if ErrorIn(y_support_try02536) {Result = y_support_try02536
            } else {
            y_support = ToList(OBJ(y_support_try02536))
            y_len := y_support.Length()
            for i_it := 0; i_it < y_len; i_it++ { 
              y = y_support.At(i_it)
              var void_try7 EID 
              _ = void_try7
              { 
              void_try7 = F_write_value_Variable(self.ClaireVar,y)
              /* ERROR PROTECTION INSERTED (void_try7-void_try7) */
              if ErrorIn(void_try7) {Result = void_try7
              break
              } else {
              var res_try02547 EID 
              /* Let:7 */{ 
                var g0255UU *ClaireAny  
                /* noccur = 1 */
                var g0255UU_try02568 EID 
                g0255UU_try02568 = EVAL(self.Arg)
                /* ERROR PROTECTION INSERTED (g0255UU-res_try02547) */
                if ErrorIn(g0255UU_try02568) {res_try02547 = g0255UU_try02568
                } else {
                g0255UU = ANY(g0255UU_try02568)
                res_try02547 = EID{res.AddFast(g0255UU).Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (res-void_try7) */
              if ErrorIn(res_try02547) {void_try7 = res_try02547
              Result = res_try02547
              break
              } else {
              res = ToList(OBJ(res_try02547))
              void_try7 = EID{res.Id(),0}
              }}
              }}
              /* loop-6 */} 
            /* For-5 */} 
          /* If-4 */} 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        Result = EID{res.Id(),0}
        }
        /* Let-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: self_eval @ Collect (throw: true) 
func E_self_eval_Collect (self EID) EID { 
    return /*(sm for self_eval @ Collect= EID)*/ To_Collect(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Collect 
func EVAL_Collect (x *ClaireAny) EID { 
     return To_Collect(x).SelfEval()} 
  
// this is a set image version, that produces a set
//
/* {1} OPT.The go function for: self_print(self:Image) [] */
func (self *Image ) SelfPrint () EID { 
    var Result EID 
    PRINC("{ ")
    Core.C_pretty.Index = (Core.C_pretty.Index+2)
    Result = F_printexp_any(self.Arg,CFALSE)
    {
    PRINC(" | ")
    Result = F_lbreak_void()
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    Result = F_ppvariable_Variable(self.ClaireVar)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(" in ")
    /* Let:2 */{ 
      var _Zl int  = Core.C_pretty.Index
      /* noccur = 1 */
      F_set_level_void()
      Result = F_printexp_any(self.SetArg,CFALSE)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      /* update:3 */{ 
        var va_arg1 *Core.PrettyPrinter  
        var va_arg2 int 
        va_arg1 = Core.C_pretty
        va_arg2 = (_Zl-2)
        /* ---------- now we compile update mClaire/index(va_arg1) := va_arg2 ------- */
        va_arg1.Index = va_arg2
        Result = EID{C__INT,IVAL(va_arg2)}
        /* update-3 */} 
      }
      /* Let-2 */} 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("}")
    Result = EVOID
    }}}}
    return Result} 
  
// The EID go function for: self_print @ Image (throw: true) 
func E_self_print_Image_Language (self EID) EID { 
    return /*(sm for self_print @ Image= EID)*/ To_Image(OBJ(self)).SelfPrint( )} 
  
/* {1} OPT.The go function for: self_eval(self:Image) [] */
func (self *Image ) SelfEval () EID { 
    var Result EID 
    /* Let:2 */{ 
      var x *ClaireAny  
      /* noccur = 1 */
      var x_try02583 EID 
      x_try02583 = EVAL(self.SetArg)
      /* ERROR PROTECTION INSERTED (x-Result) */
      if ErrorIn(x_try02583) {Result = x_try02583
      } else {
      x = ANY(x_try02583)
      /* Let:3 */{ 
        var res *ClaireSet  
        /* noccur = 3 */
        /* Let:4 */{ 
          var g0259UU *ClaireType  
          /* noccur = 1 */
          if (self.Of.Id() != CNULL) /* If:5 */{ 
            g0259UU = self.Of
            } else {
            g0259UU = ToType(CEMPTY.Id())
            /* If-5 */} 
          res = g0259UU.EmptySet()
          /* Let-4 */} 
        /* For:4 */{ 
          var y *ClaireAny  
          _ = y
          Result= EID{CFALSE.Id(),0}
          var y_support *ClaireList  
          var y_support_try02605 EID 
          y_support_try02605 = Core.F_enumerate_any(x)
          /* ERROR PROTECTION INSERTED (y_support-Result) */
          if ErrorIn(y_support_try02605) {Result = y_support_try02605
          } else {
          y_support = ToList(OBJ(y_support_try02605))
          y_len := y_support.Length()
          for i_it := 0; i_it < y_len; i_it++ { 
            y = y_support.At(i_it)
            var void_try6 EID 
            _ = void_try6
            { 
            void_try6 = F_write_value_Variable(self.ClaireVar,y)
            /* ERROR PROTECTION INSERTED (void_try6-void_try6) */
            if ErrorIn(void_try6) {Result = void_try6
            break
            } else {
            var res_try02616 EID 
            /* Let:6 */{ 
              var g0262UU *ClaireAny  
              /* noccur = 1 */
              var g0262UU_try02637 EID 
              g0262UU_try02637 = EVAL(self.Arg)
              /* ERROR PROTECTION INSERTED (g0262UU-res_try02616) */
              if ErrorIn(g0262UU_try02637) {res_try02616 = g0262UU_try02637
              } else {
              g0262UU = ANY(g0262UU_try02637)
              res_try02616 = EID{res.AddFast(g0262UU).Id(),0}
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (res-void_try6) */
            if ErrorIn(res_try02616) {void_try6 = res_try02616
            Result = res_try02616
            break
            } else {
            res = ToSet(OBJ(res_try02616))
            void_try6 = EID{res.Id(),0}
            }}
            }}
            /* loop-5 */} 
          /* For-4 */} 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        Result = EID{res.Id(),0}
        }
        /* Let-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: self_eval @ Image (throw: true) 
func E_self_eval_Image (self EID) EID { 
    return /*(sm for self_eval @ Image= EID)*/ To_Image(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Image 
func EVAL_Image (x *ClaireAny) EID { 
     return To_Image(x).SelfEval()} 
  
// [select VAR in SET_EXPR, ...] is the same as a "for" but returns the subset of
//  members that produce a true value
//
/* {1} OPT.The go function for: self_print(self:Select) [] */
func (self *Select ) SelfPrint () EID { 
    var Result EID 
    PRINC("{ ")
    Result = F_ppvariable_Variable(self.ClaireVar)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(" in ")
    /* Let:2 */{ 
      var _Zl int  = Core.C_pretty.Index
      /* noccur = 1 */
      F_set_level_void()
      Result = F_printexp_any(self.SetArg,CFALSE)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      /* update:3 */{ 
        var va_arg1 *Core.PrettyPrinter  
        var va_arg2 int 
        va_arg1 = Core.C_pretty
        va_arg2 = _Zl
        /* ---------- now we compile update mClaire/index(va_arg1) := va_arg2 ------- */
        va_arg1.Index = va_arg2
        Result = EID{C__INT,IVAL(va_arg2)}
        /* update-3 */} 
      }
      /* Let-2 */} 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(" | ")
    Result = F_lbreak_integer(2)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    Result = Core.F_CALL(C_print,ARGS(self.Arg.ToEID()))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    /* update:2 */{ 
      var va_arg1 *Core.PrettyPrinter  
      var va_arg2 int 
      va_arg1 = Core.C_pretty
      va_arg2 = (Core.C_pretty.Index-2)
      /* ---------- now we compile update mClaire/index(va_arg1) := va_arg2 ------- */
      va_arg1.Index = va_arg2
      Result = EID{C__INT,IVAL(va_arg2)}
      /* update-2 */} 
    }}
    {
    PRINC("}")
    Result = EVOID
    }}}
    return Result} 
  
// The EID go function for: self_print @ Select (throw: true) 
func E_self_print_Select_Language (self EID) EID { 
    return /*(sm for self_print @ Select= EID)*/ To_Select(OBJ(self)).SelfPrint( )} 
  
/* {1} OPT.The go function for: self_eval(self:Select) [] */
func (self *Select ) SelfEval () EID { 
    var Result EID 
    /* Let:2 */{ 
      var x *ClaireAny  
      /* noccur = 5 */
      var x_try02693 EID 
      x_try02693 = EVAL(self.SetArg)
      /* ERROR PROTECTION INSERTED (x-Result) */
      if ErrorIn(x_try02693) {Result = x_try02693
      } else {
      x = ANY(x_try02693)
      /* Let:3 */{ 
        var res *ClaireSet  
        /* noccur = 7 */
        /* Let:4 */{ 
          var g0270UU *ClaireType  
          /* noccur = 1 */
          if (self.Of.Id() != CNULL) /* If:5 */{ 
            g0270UU = self.Of
            } else {
            g0270UU = ToType(CEMPTY.Id())
            /* If-5 */} 
          res = g0270UU.EmptySet()
          /* Let-4 */} 
        if (C_class.Id() == x.Isa.Id()) /* If:4 */{ 
          /* Let:5 */{ 
            var g0265 *ClaireClass   = ToClass(x)
            /* noccur = 1 */
            /* For:6 */{ 
              var y *ClaireAny  
              _ = y
              Result= EID{CFALSE.Id(),0}
              for _,y = range(g0265.Descendents.Values)/* loop:7 */{ 
                var void_try8 EID 
                _ = void_try8
                /* For:8 */{ 
                  var z *ClaireAny  
                  _ = z
                  void_try8= EID{CFALSE.Id(),0}
                  var z_support *ClaireList  
                  z_support = ToClass(y).Instances
                  z_len := z_support.Length()
                  for i_it := 0; i_it < z_len; i_it++ { 
                    z = z_support.At(i_it)
                    var void_try10 EID 
                    _ = void_try10
                    { 
                    void_try10 = F_write_value_Variable(self.ClaireVar,z)
                    /* ERROR PROTECTION INSERTED (void_try10-void_try10) */
                    if ErrorIn(void_try10) {void_try8 = void_try10
                    break
                    } else {
                    var g0271I *ClaireBoolean  
                    var g0271I_try027210 EID 
                    /* Let:10 */{ 
                      var g0273UU *ClaireAny  
                      /* noccur = 1 */
                      var g0273UU_try027411 EID 
                      g0273UU_try027411 = EVAL(self.Arg)
                      /* ERROR PROTECTION INSERTED (g0273UU-g0271I_try027210) */
                      if ErrorIn(g0273UU_try027411) {g0271I_try027210 = g0273UU_try027411
                      } else {
                      g0273UU = ANY(g0273UU_try027411)
                      g0271I_try027210 = EID{Core.F__I_equal_any(g0273UU,CFALSE.Id()).Id(),0}
                      }
                      /* Let-10 */} 
                    /* ERROR PROTECTION INSERTED (g0271I-void_try10) */
                    if ErrorIn(g0271I_try027210) {void_try10 = g0271I_try027210
                    } else {
                    g0271I = ToBoolean(OBJ(g0271I_try027210))
                    if (g0271I == CTRUE) /* If:10 */{ 
                      res = res.AddFast(z)
                      void_try10 = EID{res.Id(),0}
                      } else {
                      void_try10 = EID{CFALSE.Id(),0}
                      /* If-10 */} 
                    }
                    /* ERROR PROTECTION INSERTED (void_try10-void_try10) */
                    if ErrorIn(void_try10) {void_try8 = void_try10
                    break
                    } else {
                    }}
                    }
                    /* loop-9 */} 
                  /* For-8 */} 
                /* ERROR PROTECTION INSERTED (void_try8-Result) */
                if ErrorIn(void_try8) {Result = void_try8
                Result = void_try8
                break
                } else {
                }
                /* loop-7 */} 
              /* For-6 */} 
            /* Let-5 */} 
          /* If!4 */}  else if (x.Isa.IsIn(C_Interval) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var g0266 *ClaireInterval   = To_Interval(x)
            /* noccur = 2 */
            /* Let:6 */{ 
              var y int  = g0266.Arg1
              /* noccur = 5 */
              /* Let:7 */{ 
                var g0267 int  = g0266.Arg2
                /* noccur = 1 */
                Result= EID{CFALSE.Id(),0}
                for (y <= g0267) /* while:8 */{ 
                  var void_try9 EID 
                  _ = void_try9
                  { 
                  void_try9 = self.ClaireVar.WriteEID(EID{C__INT,IVAL(y)})
                  /* ERROR PROTECTION INSERTED (void_try9-void_try9) */
                  if ErrorIn(void_try9) {Result = void_try9
                  break
                  } else {
                  var g0275I *ClaireBoolean  
                  var g0275I_try02769 EID 
                  /* Let:9 */{ 
                    var g0277UU *ClaireAny  
                    /* noccur = 1 */
                    var g0277UU_try027810 EID 
                    g0277UU_try027810 = EVAL(self.Arg)
                    /* ERROR PROTECTION INSERTED (g0277UU-g0275I_try02769) */
                    if ErrorIn(g0277UU_try027810) {g0275I_try02769 = g0277UU_try027810
                    } else {
                    g0277UU = ANY(g0277UU_try027810)
                    g0275I_try02769 = EID{Core.F__I_equal_any(g0277UU,CFALSE.Id()).Id(),0}
                    }
                    /* Let-9 */} 
                  /* ERROR PROTECTION INSERTED (g0275I-void_try9) */
                  if ErrorIn(g0275I_try02769) {void_try9 = g0275I_try02769
                  } else {
                  g0275I = ToBoolean(OBJ(g0275I_try02769))
                  if (g0275I == CTRUE) /* If:9 */{ 
                    res = res.AddFast(MakeInteger(y).Id())
                    void_try9 = EID{res.Id(),0}
                    } else {
                    void_try9 = EID{CFALSE.Id(),0}
                    /* If-9 */} 
                  }
                  /* ERROR PROTECTION INSERTED (void_try9-void_try9) */
                  if ErrorIn(void_try9) {Result = void_try9
                  break
                  } else {
                  }}
                  {
                  y = (y+1)
                  }
                  /* while-8 */} 
                }
                /* Let-7 */} 
              /* Let-6 */} 
            /* Let-5 */} 
          } else {
          /* For:5 */{ 
            var y *ClaireAny  
            _ = y
            Result= EID{CFALSE.Id(),0}
            var y_support *ClaireList  
            var y_support_try02796 EID 
            y_support_try02796 = Core.F_enumerate_any(x)
            /* ERROR PROTECTION INSERTED (y_support-Result) */
            if ErrorIn(y_support_try02796) {Result = y_support_try02796
            } else {
            y_support = ToList(OBJ(y_support_try02796))
            y_len := y_support.Length()
            for i_it := 0; i_it < y_len; i_it++ { 
              y = y_support.At(i_it)
              var void_try7 EID 
              _ = void_try7
              { 
              void_try7 = F_write_value_Variable(self.ClaireVar,y)
              /* ERROR PROTECTION INSERTED (void_try7-void_try7) */
              if ErrorIn(void_try7) {Result = void_try7
              break
              } else {
              var g0280I *ClaireBoolean  
              var g0280I_try02817 EID 
              /* Let:7 */{ 
                var g0282UU *ClaireAny  
                /* noccur = 1 */
                var g0282UU_try02838 EID 
                g0282UU_try02838 = EVAL(self.Arg)
                /* ERROR PROTECTION INSERTED (g0282UU-g0280I_try02817) */
                if ErrorIn(g0282UU_try02838) {g0280I_try02817 = g0282UU_try02838
                } else {
                g0282UU = ANY(g0282UU_try02838)
                g0280I_try02817 = EID{Core.F__I_equal_any(g0282UU,CFALSE.Id()).Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (g0280I-void_try7) */
              if ErrorIn(g0280I_try02817) {void_try7 = g0280I_try02817
              } else {
              g0280I = ToBoolean(OBJ(g0280I_try02817))
              if (g0280I == CTRUE) /* If:7 */{ 
                res = res.AddFast(y)
                void_try7 = EID{res.Id(),0}
                } else {
                void_try7 = EID{CFALSE.Id(),0}
                /* If-7 */} 
              }
              /* ERROR PROTECTION INSERTED (void_try7-void_try7) */
              if ErrorIn(void_try7) {Result = void_try7
              break
              } else {
              }}
              }}
              /* loop-6 */} 
            /* For-5 */} 
          /* If-4 */} 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        Result = EID{res.Id(),0}
        }
        /* Let-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: self_eval @ Select (throw: true) 
func E_self_eval_Select (self EID) EID { 
    return /*(sm for self_eval @ Select= EID)*/ To_Select(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Select 
func EVAL_Select (x *ClaireAny) EID { 
     return To_Select(x).SelfEval()} 
  
// [select VAR in SET_EXPR, ...] is the same as a "for" but returns the subset of
//  members that produce a true value
//
/* {1} OPT.The go function for: self_print(self:Lselect) [] */
func (self *Lselect ) SelfPrint () EID { 
    var Result EID 
    PRINC("list{ ")
    Result = F_ppvariable_Variable(self.ClaireVar)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(" in ")
    /* Let:2 */{ 
      var _Zl int  = Core.C_pretty.Index
      /* noccur = 1 */
      F_set_level_void()
      Result = F_printexp_any(self.SetArg,CFALSE)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      /* update:3 */{ 
        var va_arg1 *Core.PrettyPrinter  
        var va_arg2 int 
        va_arg1 = Core.C_pretty
        va_arg2 = _Zl
        /* ---------- now we compile update mClaire/index(va_arg1) := va_arg2 ------- */
        va_arg1.Index = va_arg2
        Result = EID{C__INT,IVAL(va_arg2)}
        /* update-3 */} 
      }
      /* Let-2 */} 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(" | ")
    Result = F_lbreak_integer(2)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    Result = Core.F_CALL(C_print,ARGS(self.Arg.ToEID()))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    /* update:2 */{ 
      var va_arg1 *Core.PrettyPrinter  
      var va_arg2 int 
      va_arg1 = Core.C_pretty
      va_arg2 = (Core.C_pretty.Index-2)
      /* ---------- now we compile update mClaire/index(va_arg1) := va_arg2 ------- */
      va_arg1.Index = va_arg2
      Result = EID{C__INT,IVAL(va_arg2)}
      /* update-2 */} 
    }}
    {
    PRINC("}")
    Result = EVOID
    }}}
    return Result} 
  
// The EID go function for: self_print @ Lselect (throw: true) 
func E_self_print_Lselect_Language (self EID) EID { 
    return /*(sm for self_print @ Lselect= EID)*/ To_Lselect(OBJ(self)).SelfPrint( )} 
  
/* {1} OPT.The go function for: self_eval(self:Lselect) [] */
func (self *Lselect ) SelfEval () EID { 
    var Result EID 
    /* Let:2 */{ 
      var x *ClaireAny  
      /* noccur = 5 */
      var x_try02893 EID 
      x_try02893 = EVAL(self.SetArg)
      /* ERROR PROTECTION INSERTED (x-Result) */
      if ErrorIn(x_try02893) {Result = x_try02893
      } else {
      x = ANY(x_try02893)
      /* Let:3 */{ 
        var res *ClaireList  
        /* noccur = 7 */
        if (x.Isa.IsIn(C_list) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var g0285 *ClaireList   = ToList(x)
            /* noccur = 1 */
            res = g0285.Empty()
            /* Let-5 */} 
          } else {
          res = ToType(CEMPTY.Id()).EmptyList()
          /* If-4 */} 
        if (C_class.Id() == x.Isa.Id()) /* If:4 */{ 
          /* Let:5 */{ 
            var g0287 *ClaireClass   = ToClass(x)
            /* noccur = 1 */
            /* For:6 */{ 
              var y *ClaireAny  
              _ = y
              Result= EID{CFALSE.Id(),0}
              for _,y = range(g0287.Descendents.Values)/* loop:7 */{ 
                var void_try8 EID 
                _ = void_try8
                /* For:8 */{ 
                  var z *ClaireAny  
                  _ = z
                  void_try8= EID{CFALSE.Id(),0}
                  var z_support *ClaireList  
                  z_support = ToClass(y).Instances
                  z_len := z_support.Length()
                  for i_it := 0; i_it < z_len; i_it++ { 
                    z = z_support.At(i_it)
                    var void_try10 EID 
                    _ = void_try10
                    { 
                    void_try10 = F_write_value_Variable(self.ClaireVar,z)
                    /* ERROR PROTECTION INSERTED (void_try10-void_try10) */
                    if ErrorIn(void_try10) {void_try8 = void_try10
                    break
                    } else {
                    var g0290I *ClaireBoolean  
                    var g0290I_try029110 EID 
                    /* Let:10 */{ 
                      var g0292UU *ClaireAny  
                      /* noccur = 1 */
                      var g0292UU_try029311 EID 
                      g0292UU_try029311 = EVAL(self.Arg)
                      /* ERROR PROTECTION INSERTED (g0292UU-g0290I_try029110) */
                      if ErrorIn(g0292UU_try029311) {g0290I_try029110 = g0292UU_try029311
                      } else {
                      g0292UU = ANY(g0292UU_try029311)
                      g0290I_try029110 = EID{Core.F__I_equal_any(g0292UU,CFALSE.Id()).Id(),0}
                      }
                      /* Let-10 */} 
                    /* ERROR PROTECTION INSERTED (g0290I-void_try10) */
                    if ErrorIn(g0290I_try029110) {void_try10 = g0290I_try029110
                    } else {
                    g0290I = ToBoolean(OBJ(g0290I_try029110))
                    if (g0290I == CTRUE) /* If:10 */{ 
                      res = res.AddFast(z)
                      void_try10 = EID{res.Id(),0}
                      } else {
                      void_try10 = EID{CFALSE.Id(),0}
                      /* If-10 */} 
                    }
                    /* ERROR PROTECTION INSERTED (void_try10-void_try10) */
                    if ErrorIn(void_try10) {void_try8 = void_try10
                    break
                    } else {
                    }}
                    }
                    /* loop-9 */} 
                  /* For-8 */} 
                /* ERROR PROTECTION INSERTED (void_try8-Result) */
                if ErrorIn(void_try8) {Result = void_try8
                Result = void_try8
                break
                } else {
                }
                /* loop-7 */} 
              /* For-6 */} 
            /* Let-5 */} 
          } else {
          /* For:5 */{ 
            var y *ClaireAny  
            _ = y
            Result= EID{CFALSE.Id(),0}
            var y_support *ClaireList  
            var y_support_try02946 EID 
            y_support_try02946 = Core.F_enumerate_any(x)
            /* ERROR PROTECTION INSERTED (y_support-Result) */
            if ErrorIn(y_support_try02946) {Result = y_support_try02946
            } else {
            y_support = ToList(OBJ(y_support_try02946))
            y_len := y_support.Length()
            for i_it := 0; i_it < y_len; i_it++ { 
              y = y_support.At(i_it)
              var void_try7 EID 
              _ = void_try7
              { 
              void_try7 = F_write_value_Variable(self.ClaireVar,y)
              /* ERROR PROTECTION INSERTED (void_try7-void_try7) */
              if ErrorIn(void_try7) {Result = void_try7
              break
              } else {
              var g0295I *ClaireBoolean  
              var g0295I_try02967 EID 
              /* Let:7 */{ 
                var g0297UU *ClaireAny  
                /* noccur = 1 */
                var g0297UU_try02988 EID 
                g0297UU_try02988 = EVAL(self.Arg)
                /* ERROR PROTECTION INSERTED (g0297UU-g0295I_try02967) */
                if ErrorIn(g0297UU_try02988) {g0295I_try02967 = g0297UU_try02988
                } else {
                g0297UU = ANY(g0297UU_try02988)
                g0295I_try02967 = EID{Core.F__I_equal_any(g0297UU,CFALSE.Id()).Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (g0295I-void_try7) */
              if ErrorIn(g0295I_try02967) {void_try7 = g0295I_try02967
              } else {
              g0295I = ToBoolean(OBJ(g0295I_try02967))
              if (g0295I == CTRUE) /* If:7 */{ 
                res = res.AddFast(y)
                void_try7 = EID{res.Id(),0}
                } else {
                void_try7 = EID{CFALSE.Id(),0}
                /* If-7 */} 
              }
              /* ERROR PROTECTION INSERTED (void_try7-void_try7) */
              if ErrorIn(void_try7) {Result = void_try7
              break
              } else {
              }}
              }}
              /* loop-6 */} 
            /* For-5 */} 
          /* If-4 */} 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        if (self.Of.Id() != CNULL) /* If:4 */{ 
          /* Let:5 */{ 
            var x *ClaireAny  
            /* noccur = 2 */
            /* Let:6 */{ 
              var x_some *ClaireAny   = CNULL
              /* noccur = 2 */
              /* For:7 */{ 
                var x *ClaireAny  
                _ = x
                var x_support *ClaireList  
                x_support = res
                x_len := x_support.Length()
                for i_it := 0; i_it < x_len; i_it++ { 
                  x = x_support.At(i_it)
                  if (self.Of.Contains(x) != CTRUE) /* If:9 */{ 
                     /*v = x, s =void*/
x_some = x
                    break
                    /* If-9 */} 
                  /* loop-8 */} 
                /* For-7 */} 
              x = x_some
              /* Let-6 */} 
            if (x != CNULL) /* If:6 */{ 
              Result = ToException(Core.C_range_error.Make(self.Id(),x,self.Of.Id())).Close()
              } else {
              Result = EID{CNULL,0}
              /* If-6 */} 
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          Result = EID{res.Cast_I(self.Of).Id(),0}
          }
          } else {
          Result = EID{CFALSE.Id(),0}
          /* If-4 */} 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        Result = EID{res.Id(),0}
        }}
        /* Let-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: self_eval @ Lselect (throw: true) 
func E_self_eval_Lselect (self EID) EID { 
    return /*(sm for self_eval @ Lselect= EID)*/ To_Lselect(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Lselect 
func EVAL_Lselect (x *ClaireAny) EID { 
     return To_Lselect(x).SelfEval()} 
  
// Exists is an iteration that checks a condition
// other = true => forall,  other = false => exists, other = unknown => some
/* {1} OPT.The go function for: self_print(self:Exists) [] */
func (self *Exists ) SelfPrint () EID { 
    var Result EID 
    if (self.Other == CTRUE.Id()) /* If:2 */{ 
      PRINC("forall")
      /* If!2 */}  else if (self.Other == CFALSE.Id()) /* If:2 */{ 
      PRINC("exists")
      } else {
      PRINC("some")
      /* If-2 */} 
    if (self.SetArg == C_any.Id()) /* If:2 */{ 
      PRINC("(")
      Result = F_ppvariable_Variable(self.ClaireVar)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(",")
      Result = Core.F_CALL(C_print,ARGS(self.Arg.ToEID()))
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(")")
      Result = EVOID
      }}
      } else {
      PRINC("(")
      Result = F_ppvariable_Variable(self.ClaireVar)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(" in ")
      /* Let:3 */{ 
        var _Zl int  = Core.C_pretty.Index
        /* noccur = 1 */
        F_set_level_void()
        Result = F_printexp_any(self.SetArg,CFALSE)
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        /* update:4 */{ 
          var va_arg1 *Core.PrettyPrinter  
          var va_arg2 int 
          va_arg1 = Core.C_pretty
          va_arg2 = _Zl
          /* ---------- now we compile update mClaire/index(va_arg1) := va_arg2 ------- */
          va_arg1.Index = va_arg2
          Result = EID{C__INT,IVAL(va_arg2)}
          /* update-4 */} 
        }
        /* Let-3 */} 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(" | ")
      Result = F_lbreak_integer(2)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = Core.F_CALL(C_print,ARGS(self.Arg.ToEID()))
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      /* update:3 */{ 
        var va_arg1 *Core.PrettyPrinter  
        var va_arg2 int 
        va_arg1 = Core.C_pretty
        va_arg2 = (Core.C_pretty.Index-2)
        /* ---------- now we compile update mClaire/index(va_arg1) := va_arg2 ------- */
        va_arg1.Index = va_arg2
        Result = EID{C__INT,IVAL(va_arg2)}
        /* update-3 */} 
      }}
      {
      PRINC(")")
      Result = EVOID
      }}}
      /* If-2 */} 
    return Result} 
  
// The EID go function for: self_print @ Exists (throw: true) 
func E_self_print_Exists_Language (self EID) EID { 
    return /*(sm for self_print @ Exists= EID)*/ To_Exists(OBJ(self)).SelfPrint( )} 
  
/* {1} OPT.The go function for: self_eval(self:Exists) [] */
func (self *Exists ) SelfEval () EID { 
    var Result EID 
    /* Let:2 */{ 
      var x *ClaireAny  
      /* noccur = 3 */
      var x_try03023 EID 
      x_try03023 = EVAL(self.SetArg)
      /* ERROR PROTECTION INSERTED (x-Result) */
      if ErrorIn(x_try03023) {Result = x_try03023
      } else {
      x = ANY(x_try03023)
      /* Let:3 */{ 
        var b *ClaireAny   = self.Other
        /* noccur = 7 */
        /* Let:4 */{ 
          var res *ClaireAny   = b
          /* noccur = 5 */
          if (C_class.Id() == x.Isa.Id()) /* If:5 */{ 
            /* Let:6 */{ 
              var g0300 *ClaireClass   = ToClass(x)
              /* noccur = 1 */
              /* For:7 */{ 
                var y *ClaireAny  
                _ = y
                Result= EID{CFALSE.Id(),0}
                for _,y = range(g0300.Descendents.Values)/* loop:8 */{ 
                  var void_try9 EID 
                  _ = void_try9
                  /* For:9 */{ 
                    var z *ClaireAny  
                    _ = z
                    void_try9= EID{CFALSE.Id(),0}
                    var z_support *ClaireList  
                    z_support = ToClass(y).Instances
                    z_len := z_support.Length()
                    for i_it := 0; i_it < z_len; i_it++ { 
                      z = z_support.At(i_it)
                      var void_try11 EID 
                      _ = void_try11
                      { 
                      void_try11 = F_write_value_Variable(self.ClaireVar,z)
                      /* ERROR PROTECTION INSERTED (void_try11-void_try11) */
                      if ErrorIn(void_try11) {void_try9 = void_try11
                      break
                      } else {
                      var g0303I *ClaireBoolean  
                      var g0303I_try030411 EID 
                      /* Let:11 */{ 
                        var g0305UU *ClaireAny  
                        /* noccur = 1 */
                        var g0305UU_try030612 EID 
                        g0305UU_try030612 = EVAL(self.Arg)
                        /* ERROR PROTECTION INSERTED (g0305UU-g0303I_try030411) */
                        if ErrorIn(g0305UU_try030612) {g0303I_try030411 = g0305UU_try030612
                        } else {
                        g0305UU = ANY(g0305UU_try030612)
                        g0303I_try030411 = EID{Core.F__I_equal_any(g0305UU,CFALSE.Id()).Id(),0}
                        }
                        /* Let-11 */} 
                      /* ERROR PROTECTION INSERTED (g0303I-void_try11) */
                      if ErrorIn(g0303I_try030411) {void_try11 = g0303I_try030411
                      } else {
                      g0303I = ToBoolean(OBJ(g0303I_try030411))
                      if (g0303I == CTRUE) /* If:11 */{ 
                        if (b != CTRUE.Id()) /* If:12 */{ 
                           /*v = void_try9, s =EID*/
res = IfThenElse((F_boolean_I_any(b) == CTRUE),
                            z,
                            CTRUE.Id())
                          void_try9 = res.ToEID()
                          break
                          } else {
                          void_try11 = EID{CFALSE.Id(),0}
                          /* If-12 */} 
                        /* If!11 */}  else if (b == CTRUE.Id()) /* If:11 */{ 
                         /*v = void_try9, s =EID*/
res = CFALSE.Id()
                        void_try9 = res.ToEID()
                        break
                        } else {
                        void_try11 = EID{CFALSE.Id(),0}
                        /* If-11 */} 
                      }
                      /* ERROR PROTECTION INSERTED (void_try11-void_try11) */
                      if ErrorIn(void_try11) {void_try9 = void_try11
                      break
                      } else {
                      }}
                      }
                      /* loop-10 */} 
                    /* For-9 */} 
                  /* ERROR PROTECTION INSERTED (void_try9-Result) */
                  if ErrorIn(void_try9) {Result = void_try9
                  Result = void_try9
                  break
                  } else {
                  }
                  /* loop-8 */} 
                /* For-7 */} 
              /* Let-6 */} 
            } else {
            /* For:6 */{ 
              var y *ClaireAny  
              _ = y
              Result= EID{CFALSE.Id(),0}
              var y_support *ClaireList  
              var y_support_try03077 EID 
              y_support_try03077 = Core.F_enumerate_any(x)
              /* ERROR PROTECTION INSERTED (y_support-Result) */
              if ErrorIn(y_support_try03077) {Result = y_support_try03077
              } else {
              y_support = ToList(OBJ(y_support_try03077))
              y_len := y_support.Length()
              for i_it := 0; i_it < y_len; i_it++ { 
                y = y_support.At(i_it)
                var void_try8 EID 
                _ = void_try8
                { 
                void_try8 = F_write_value_Variable(self.ClaireVar,y)
                /* ERROR PROTECTION INSERTED (void_try8-void_try8) */
                if ErrorIn(void_try8) {Result = void_try8
                break
                } else {
                var g0308I *ClaireBoolean  
                var g0308I_try03098 EID 
                /* Let:8 */{ 
                  var g0310UU *ClaireAny  
                  /* noccur = 1 */
                  var g0310UU_try03119 EID 
                  g0310UU_try03119 = EVAL(self.Arg)
                  /* ERROR PROTECTION INSERTED (g0310UU-g0308I_try03098) */
                  if ErrorIn(g0310UU_try03119) {g0308I_try03098 = g0310UU_try03119
                  } else {
                  g0310UU = ANY(g0310UU_try03119)
                  g0308I_try03098 = EID{Core.F__I_equal_any(g0310UU,CFALSE.Id()).Id(),0}
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (g0308I-void_try8) */
                if ErrorIn(g0308I_try03098) {void_try8 = g0308I_try03098
                } else {
                g0308I = ToBoolean(OBJ(g0308I_try03098))
                if (g0308I == CTRUE) /* If:8 */{ 
                  if (b != CTRUE.Id()) /* If:9 */{ 
                     /*v = Result, s =EID*/
res = IfThenElse((F_boolean_I_any(b) == CTRUE),
                      y,
                      CTRUE.Id())
                    Result = res.ToEID()
                    break
                    } else {
                    void_try8 = EID{CFALSE.Id(),0}
                    /* If-9 */} 
                  /* If!8 */}  else if (b == CTRUE.Id()) /* If:8 */{ 
                   /*v = Result, s =EID*/
res = CFALSE.Id()
                  Result = res.ToEID()
                  break
                  } else {
                  void_try8 = EID{CFALSE.Id(),0}
                  /* If-8 */} 
                }
                /* ERROR PROTECTION INSERTED (void_try8-void_try8) */
                if ErrorIn(void_try8) {Result = void_try8
                break
                } else {
                }}
                }}
                /* loop-7 */} 
              /* For-6 */} 
            /* If-5 */} 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          Result = res.ToEID()
          }
          /* Let-4 */} 
        /* Let-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: self_eval @ Exists (throw: true) 
func E_self_eval_Exists (self EID) EID { 
    return /*(sm for self_eval @ Exists= EID)*/ To_Exists(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Exists 
func EVAL_Exists (x *ClaireAny) EID { 
     return To_Exists(x).SelfEval()} 
  
// *********************************************************************
// *     Part 3: other control structures                              *
// *********************************************************************
// ----------------- case  --------------------------------------
/* {1} OPT.The go function for: self_print(self:Case) [] */
func (self *Case ) SelfPrint () EID { 
    var Result EID 
    PRINC("case ")
    Result = Core.F_CALL(C_print,ARGS(self.ClaireVar.ToEID()))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(" ")
    Result = F_lbreak_integer(1)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("(")
    Result = EVOID
    }}
    {
    /* Let:2 */{ 
      var n int  = 1
      /* noccur = 6 */
      /* Let:3 */{ 
        var m int  = self.Args.Length()
        /* noccur = 2 */
        Core.C_pretty.Index = (Core.C_pretty.Index+1)
        Result= EID{CFALSE.Id(),0}
        for (n <= m) /* while:4 */{ 
          var void_try5 EID 
          _ = void_try5
          /* Let:5 */{ 
            var _Zl int  = Core.C_pretty.Index
            /* noccur = 1 */
            void_try5 = F_printexp_any(self.Args.At(n-1),CFALSE)
            /* ERROR PROTECTION INSERTED (void_try5-void_try5) */
            if ErrorIn(void_try5) {Result = void_try5
            break
            } else {
            PRINC(" ")
            if (Core.F_buffer_length_void() > (Core.C_pretty.Width-50)) /* If:6 */{ 
              void_try5 = F_lbreak_integer(2)
              } else {
              F_set_level_void()
              void_try5 = EVOID
              /* If-6 */} 
            /* ERROR PROTECTION INSERTED (void_try5-void_try5) */
            if ErrorIn(void_try5) {Result = void_try5
            break
            } else {
            void_try5 = Core.F_CALL(C_print,ARGS(self.Args.At((n+1)-1).ToEID()))
            /* ERROR PROTECTION INSERTED (void_try5-void_try5) */
            if ErrorIn(void_try5) {Result = void_try5
            break
            } else {
            }}
            {
            Core.C_pretty.Index = _Zl
            if ((n+1) != m) /* If:6 */{ 
              PRINC(", ")
              void_try5 = F_lbreak_void()
              /* ERROR PROTECTION INSERTED (void_try5-void_try5) */
              if ErrorIn(void_try5) {Result = void_try5
              break
              } else {
              PRINC("")
              void_try5 = EVOID
              }
              } else {
              void_try5 = EID{CFALSE.Id(),0}
              /* If-6 */} 
            /* ERROR PROTECTION INSERTED (void_try5-void_try5) */
            if ErrorIn(void_try5) {Result = void_try5
            break
            } else {
            }
            {
            PRINC("")
            void_try5 = EVOID
            }}}
            {
            n = (n+2)
            void_try5 = EID{C__INT,IVAL(n)}
            }
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (void_try5-Result) */
          if ErrorIn(void_try5) {Result = void_try5
          Result = void_try5
          break
          } else {
          /* while-4 */} 
        }
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC(")")
        /* update:4 */{ 
          var va_arg1 *Core.PrettyPrinter  
          var va_arg2 int 
          va_arg1 = Core.C_pretty
          va_arg2 = (Core.C_pretty.Index-2)
          /* ---------- now we compile update mClaire/index(va_arg1) := va_arg2 ------- */
          va_arg1.Index = va_arg2
          Result = EID{C__INT,IVAL(va_arg2)}
          /* update-4 */} 
        }
        /* Let-3 */} 
      /* Let-2 */} 
    }
    return Result} 
  
// The EID go function for: self_print @ Case (throw: true) 
func E_self_print_Case_Language (self EID) EID { 
    return /*(sm for self_print @ Case= EID)*/ To_Case(OBJ(self)).SelfPrint( )} 
  
/* {1} OPT.The go function for: self_eval(self:Case) [] */
func (self *Case ) SelfEval () EID { 
    var Result EID 
    /* Let:2 */{ 
      var truc *ClaireAny  
      /* noccur = 1 */
      var truc_try03143 EID 
      truc_try03143 = EVAL(self.ClaireVar)
      /* ERROR PROTECTION INSERTED (truc-Result) */
      if ErrorIn(truc_try03143) {Result = truc_try03143
      } else {
      truc = ANY(truc_try03143)
      /* Let:3 */{ 
        var flip *ClaireBoolean   = CTRUE
        /* noccur = 3 */
        /* Let:4 */{ 
          var previous *ClaireAny   = CFALSE.Id()
          /* noccur = 4 */
          var g0315I *ClaireBoolean  
          var g0315I_try03165 EID 
          /* For:5 */{ 
            var x *ClaireAny  
            _ = x
            g0315I_try03165= EID{CFALSE.Id(),0}
            var x_support *ClaireList  
            x_support = self.Args
            x_len := x_support.Length()
            for i_it := 0; i_it < x_len; i_it++ { 
              x = x_support.At(i_it)
              var void_try7 EID 
              _ = void_try7
              if (flip == CTRUE) /* If:7 */{ 
                flip = CFALSE
                var previous_try03178 EID 
                previous_try03178 = EVAL(x)
                /* ERROR PROTECTION INSERTED (previous-void_try7) */
                if ErrorIn(previous_try03178) {void_try7 = previous_try03178
                g0315I_try03165 = previous_try03178
                break
                } else {
                previous = ANY(previous_try03178)
                void_try7 = previous.ToEID()
                }
                } else {
                var g0318I *ClaireBoolean  
                var g0318I_try03198 EID 
                g0318I_try03198 = Core.F_BELONG(truc,previous)
                /* ERROR PROTECTION INSERTED (g0318I-void_try7) */
                if ErrorIn(g0318I_try03198) {void_try7 = g0318I_try03198
                } else {
                g0318I = ToBoolean(OBJ(g0318I_try03198))
                if (g0318I == CTRUE) /* If:8 */{ 
                  var previous_try03209 EID 
                  previous_try03209 = EVAL(x)
                  /* ERROR PROTECTION INSERTED (previous-void_try7) */
                  if ErrorIn(previous_try03209) {void_try7 = previous_try03209
                  g0315I_try03165 = previous_try03209
                  break
                  } else {
                  previous = ANY(previous_try03209)
                  void_try7 = previous.ToEID()
                   /*v = g0315I_try03165, s =EID*/
g0315I_try03165 = EID{CTRUE.Id(),0}
                  break
                  }
                  } else {
                  flip = CTRUE
                  void_try7 = EID{flip.Id(),0}
                  /* If-8 */} 
                }
                /* If-7 */} 
              /* ERROR PROTECTION INSERTED (void_try7-g0315I_try03165) */
              if ErrorIn(void_try7) {g0315I_try03165 = void_try7
              g0315I_try03165 = void_try7
              break
              } else {
              }
              /* loop-6 */} 
            /* For-5 */} 
          /* ERROR PROTECTION INSERTED (g0315I-Result) */
          if ErrorIn(g0315I_try03165) {Result = g0315I_try03165
          } else {
          g0315I = ToBoolean(OBJ(g0315I_try03165))
          if (g0315I == CTRUE) /* If:5 */{ 
            Result = previous.ToEID()
            } else {
            Result = EID{CFALSE.Id(),0}
            /* If-5 */} 
          }
          /* Let-4 */} 
        /* Let-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: self_eval @ Case (throw: true) 
func E_self_eval_Case (self EID) EID { 
    return /*(sm for self_eval @ Case= EID)*/ To_Case(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Case 
func EVAL_Case (x *ClaireAny) EID { 
     return To_Case(x).SelfEval()} 
  
// ------------------ WHILE  and UNTIL  -----------------------------
// the "other" while is until, where the first test is skipped
/* {1} OPT.The go function for: self_print(self:While) [] */
func (self *While ) SelfPrint () EID { 
    var Result EID 
    F_princ_string(ToString(IfThenElse((self.Other == CTRUE),
      MakeString("until").Id(),
      MakeString("while").Id())))
    PRINC(" ")
    Result = F_printexp_any(self.Test,CFALSE)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(" ")
    Result = F_lbreak_integer(2)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    Result = Core.F_CALL(C_print,ARGS(self.Arg.ToEID()))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("")
    Result = EVOID
    }}}
    {
    /* update:2 */{ 
      var va_arg1 *Core.PrettyPrinter  
      var va_arg2 int 
      va_arg1 = Core.C_pretty
      va_arg2 = (Core.C_pretty.Index-2)
      /* ---------- now we compile update mClaire/index(va_arg1) := va_arg2 ------- */
      va_arg1.Index = va_arg2
      Result = EID{C__INT,IVAL(va_arg2)}
      /* update-2 */} 
    }
    return Result} 
  
// The EID go function for: self_print @ While (throw: true) 
func E_self_print_While_Language (self EID) EID { 
    return /*(sm for self_print @ While= EID)*/ To_While(OBJ(self)).SelfPrint( )} 
  
/* {1} OPT.The go function for: self_eval(self:While) [] */
func (self *While ) SelfEval () EID { 
    var Result EID 
    /* Let:2 */{ 
      var a *ClaireBoolean   = self.Other
      /* noccur = 2 */
      /* Let:3 */{ 
        var b *ClaireBoolean   = a
        /* noccur = 2 */
        h_index := ClEnv.Index /* Handle */
        h_base := ClEnv.Base
        var v_while4 *ClaireBoolean  
        Result= EID{CFALSE.Id(),0}
        var v_while4_try03224 EID 
        /* or:4 */{ 
          var v_or4 *ClaireBoolean  
          
          v_or4 = b
          if (v_or4 == CTRUE) {v_while4_try03224 = EID{CTRUE.Id(),0}
          } else /* or:5 */{ 
            var v_or4_try03236 EID 
            /* Let:6 */{ 
              var g0324UU *ClaireBoolean  
              /* noccur = 1 */
              var g0324UU_try03257 EID 
              /* Let:7 */{ 
                var g0326UU *ClaireAny  
                /* noccur = 1 */
                var g0326UU_try03278 EID 
                g0326UU_try03278 = EVAL(self.Test)
                /* ERROR PROTECTION INSERTED (g0326UU-g0324UU_try03257) */
                if ErrorIn(g0326UU_try03278) {g0324UU_try03257 = g0326UU_try03278
                } else {
                g0326UU = ANY(g0326UU_try03278)
                g0324UU_try03257 = EID{Core.F_not_any(g0326UU).Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (g0324UU-v_or4_try03236) */
              if ErrorIn(g0324UU_try03257) {v_or4_try03236 = g0324UU_try03257
              } else {
              g0324UU = ToBoolean(OBJ(g0324UU_try03257))
              v_or4_try03236 = EID{Equal(g0324UU.Id(),a.Id()).Id(),0}
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (v_or4-v_while4_try03224) */
            if ErrorIn(v_or4_try03236) {v_while4_try03224 = v_or4_try03236
            } else {
            v_or4 = ToBoolean(OBJ(v_or4_try03236))
            if (v_or4 == CTRUE) {v_while4_try03224 = EID{CTRUE.Id(),0}
            } else /* or:6 */{ 
              v_while4_try03224 = EID{CFALSE.Id(),0}/* org-6 */} 
            /* org-5 */} 
          }
          /* or-4 */} 
        /* ERROR PROTECTION INSERTED (v_while4-Result) */
        if ErrorIn(v_while4_try03224) {Result = v_while4_try03224
        } else {
        v_while4 = ToBoolean(OBJ(v_while4_try03224))
        
        for v_while4 == CTRUE /* while:4 */{ 
          var void_try5 EID 
          _ = void_try5
          { 
          b = CFALSE
          void_try5 = EVAL(self.Arg)
          /* ERROR PROTECTION INSERTED (void_try5-void_try5) */
          if ErrorIn(void_try5) {Result = void_try5
          break
          } else {
          }
          var v_while4_try03285 EID 
          /* or:5 */{ 
            var v_or5 *ClaireBoolean  
            
            v_or5 = b
            if (v_or5 == CTRUE) {v_while4_try03285 = EID{CTRUE.Id(),0}
            } else /* or:6 */{ 
              var v_or5_try03297 EID 
              /* Let:7 */{ 
                var g0330UU *ClaireBoolean  
                /* noccur = 1 */
                var g0330UU_try03318 EID 
                /* Let:8 */{ 
                  var g0332UU *ClaireAny  
                  /* noccur = 1 */
                  var g0332UU_try03339 EID 
                  g0332UU_try03339 = EVAL(self.Test)
                  /* ERROR PROTECTION INSERTED (g0332UU-g0330UU_try03318) */
                  if ErrorIn(g0332UU_try03339) {g0330UU_try03318 = g0332UU_try03339
                  } else {
                  g0332UU = ANY(g0332UU_try03339)
                  g0330UU_try03318 = EID{Core.F_not_any(g0332UU).Id(),0}
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (g0330UU-v_or5_try03297) */
                if ErrorIn(g0330UU_try03318) {v_or5_try03297 = g0330UU_try03318
                } else {
                g0330UU = ToBoolean(OBJ(g0330UU_try03318))
                v_or5_try03297 = EID{Equal(g0330UU.Id(),a.Id()).Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (v_or5-v_while4_try03285) */
              if ErrorIn(v_or5_try03297) {v_while4_try03285 = v_or5_try03297
              } else {
              v_or5 = ToBoolean(OBJ(v_or5_try03297))
              if (v_or5 == CTRUE) {v_while4_try03285 = EID{CTRUE.Id(),0}
              } else /* or:7 */{ 
                v_while4_try03285 = EID{CFALSE.Id(),0}/* org-7 */} 
              /* org-6 */} 
            }
            /* or-5 */} 
          /* ERROR PROTECTION INSERTED (v_while4-Result) */
          if ErrorIn(v_while4_try03285) {Result = v_while4_try03285
          Result = v_while4_try03285
          break
          } else {
          v_while4 = ToBoolean(OBJ(v_while4_try03285))
          /* while-4 */} 
        }}
        }
        if ErrorIn(Result) && ToType(Core.C_return_error.Id()).Contains(ANY(Result)) == CTRUE { 
          /* s=EID */ClEnv.Index = h_index
          ClEnv.Base = h_base
          Result = Core.F_CALL(C_arg,ARGS(EID{ClEnv.Exception_I.Id(),0}))
          } 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: self_eval @ While (throw: true) 
func E_self_eval_While (self EID) EID { 
    return /*(sm for self_eval @ While= EID)*/ To_While(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: While 
func EVAL_While (x *ClaireAny) EID { 
     return To_While(x).SelfEval()} 
  
//-------------- handling errors -----------------------------------
// This is the control structure associated with these errors. Its real
// semantics is defined in the C compiler file
//
/* {1} OPT.The go function for: self_print(self:Handle) [] */
func (self *ClaireHandle ) SelfPrint () EID { 
    var Result EID 
    PRINC("try ")
    Result = Core.F_CALL(C_print,ARGS(self.Arg.ToEID()))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(" ")
    Result = F_lbreak_integer(0)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("catch ")
    Result = Core.F_CALL(C_print,ARGS(self.Test.ToEID()))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(" ")
    Result = Core.F_CALL(C_print,ARGS(self.Other.ToEID()))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("")
    Result = EVOID
    }}}}
    {
    /* update:2 */{ 
      var va_arg1 *Core.PrettyPrinter  
      var va_arg2 int 
      va_arg1 = Core.C_pretty
      va_arg2 = (Core.C_pretty.Index-2)
      /* ---------- now we compile update mClaire/index(va_arg1) := va_arg2 ------- */
      va_arg1.Index = va_arg2
      Result = EID{C__INT,IVAL(va_arg2)}
      /* update-2 */} 
    }
    return Result} 
  
// The EID go function for: self_print @ Handle (throw: true) 
func E_self_print_Handle_Language (self EID) EID { 
    return /*(sm for self_print @ Handle= EID)*/ To_ClaireHandle(OBJ(self)).SelfPrint( )} 
  
// original code
// self_eval(self:Handle) : any
//  -> (let x := (self.test as class) in
//       try eval(self.arg)
//       catch x (if (exception!() % return_error) close(exception!())
//                else eval(self.other)))     // <yc> 6/98
// CLAIRE 4 VERSION, because catch x => x is a constant class
// notice that return_error should be called return_exception since they travel through intepreted
// not a problem at compile time since return_exceptions are handled with break(x)
/* {1} OPT.The go function for: self_eval(self:Handle) [] */
func (self *ClaireHandle ) SelfEval () EID { 
    var Result EID 
    h_index := ClEnv.Index /* Handle */
    h_base := ClEnv.Base
    Result = EVAL(self.Arg)
    if ErrorIn(Result){ 
      /* s=EID */ClEnv.Index = h_index
      ClEnv.Base = h_base
      /* Let:3 */{ 
        var e *ClaireException   = ClEnv.Exception_I
        /* noccur = 3 */
        /* Let:4 */{ 
          var x *ClaireClass   = ToClass(self.Test)
          /* noccur = 1 */
          if ((e.Isa.IsIn(Core.C_return_error) == CTRUE) || 
              (Core.F__Z_any1(e.Id(),x) != CTRUE)) /* If:5 */{ 
            Result = e.Close()
            } else {
            Result = EVAL(self.Other)
            /* If-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      } 
    return Result} 
  
// The EID go function for: self_eval @ Handle (throw: true) 
func E_self_eval_Handle (self EID) EID { 
    return /*(sm for self_eval @ Handle= EID)*/ To_ClaireHandle(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Handle 
func EVAL_Handle (x *ClaireAny) EID { 
     return To_ClaireHandle(x).SelfEval()} 
  
// <yc> 6/98
// *********************************************************************
// *     Part 4: the constructs                                         *
// *********************************************************************
// v3.2.16   constructor for arrays
/* {1} OPT.The go function for: self_print(self:Construct) [] */
func (self *Construct ) SelfPrint () EID { 
    var Result EID 
    /* Let:2 */{ 
      var _Zl int  = Core.C_pretty.Index
      /* noccur = 1 */
      /* Let:3 */{ 
        var g0345UU *ClaireString  
        /* noccur = 1 */
        if (self.Isa.IsIn(C_List) == CTRUE) /* If:4 */{ 
          g0345UU = MakeString("list")
          /* If!4 */}  else if (self.Isa.IsIn(C_Set) == CTRUE) /* If:4 */{ 
          g0345UU = MakeString("set")
          /* If!4 */}  else if (self.Isa.IsIn(C_Tuple) == CTRUE) /* If:4 */{ 
          g0345UU = MakeString("tuple")
          /* If!4 */}  else if (self.Isa.IsIn(C_Printf) == CTRUE) /* If:4 */{ 
          g0345UU = MakeString("printf")
          /* If!4 */}  else if (self.Isa.IsIn(C_Error) == CTRUE) /* If:4 */{ 
          g0345UU = MakeString("error")
          /* If!4 */}  else if (self.Isa.IsIn(C_Trace) == CTRUE) /* If:4 */{ 
          g0345UU = MakeString("trace")
          /* If!4 */}  else if (self.Isa.IsIn(C_Assert) == CTRUE) /* If:4 */{ 
          g0345UU = MakeString("assert")
          /* If!4 */}  else if (self.Isa.IsIn(C_Branch) == CTRUE) /* If:4 */{ 
          g0345UU = MakeString("branch")
          } else {
          g0345UU = self.Isa.Name.String_I()
          /* If-4 */} 
        F_princ_string(g0345UU)
        /* Let-3 */} 
      if ((self.Isa.IsIn(C_List) == CTRUE) || 
          (self.Isa.IsIn(C_Set) == CTRUE)) /* If:3 */{ 
        /* Let:4 */{ 
          var g0344 *Construct   = self
          /* noccur = 1 */
          /* Let:5 */{ 
            var _Zt *ClaireAny   = Core.F_get_property(C_of,ToObject(g0344.Id()))
            /* noccur = 3 */
            if (_Zt != CNULL) /* If:6 */{ 
              if (Equal(_Zt,CEMPTY.Id()) != CTRUE) /* If:7 */{ 
                PRINC("<")
                Result = Core.F_CALL(C_print,ARGS(_Zt.ToEID()))
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                PRINC(">")
                Result = EVOID
                }
                } else {
                Result = EID{CFALSE.Id(),0}
                /* If-7 */} 
              } else {
              Result = EID{CNULL,0}
              /* If-6 */} 
            /* Let-5 */} 
          /* Let-4 */} 
        } else {
        Result = EID{CFALSE.Id(),0}
        /* If-3 */} 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC("(")
      F_set_level_void()
      Result = F_Language_printbox_list2(self.Args)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(")")
      Result = EVOID
      }}
      {
      /* update:3 */{ 
        var va_arg1 *Core.PrettyPrinter  
        var va_arg2 int 
        va_arg1 = Core.C_pretty
        va_arg2 = _Zl
        /* ---------- now we compile update mClaire/index(va_arg1) := va_arg2 ------- */
        va_arg1.Index = va_arg2
        Result = EID{C__INT,IVAL(va_arg2)}
        /* update-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: self_print @ Construct (throw: true) 
func E_self_print_Construct_Language (self EID) EID { 
    return /*(sm for self_print @ Construct= EID)*/ To_Construct(OBJ(self)).SelfPrint( )} 
  
// constructors: how to create a list, a set, a tuple or an array
// note that the constructor is typed
// CLAIRE4: must build the list with the proper type from the begining, so that Srange is correct
/* {1} OPT.The go function for: self_eval(self:List) [] */
func (self *List ) SelfEval () EID { 
    var Result EID 
    /* Let:2 */{ 
      var type_ask *ClaireBoolean   = MakeBoolean((self.Of.Id() == CNULL)).Not
      /* noccur = 1 */
      /* Let:3 */{ 
        var n int  = self.Args.Length()
        /* noccur = 4 */
        if (type_ask == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var l *ClaireList   = CreateList(self.Of,n)
            /* noccur = 2 */
            /* Let:6 */{ 
              var i int  = 1
              /* noccur = 5 */
              /* Let:7 */{ 
                var g0346 int  = n
                /* noccur = 1 */
                Result= EID{CFALSE.Id(),0}
                for (i <= g0346) /* while:8 */{ 
                  var void_try9 EID 
                  _ = void_try9
                  { 
                  /* LetEID:9 */{ 
                    var g0348UU EID 
                    g0348UU = EVAL(self.Args.At(i-1))
                    /* ERROR PROTECTION INSERTED (g0348UU-void_try9) */
                    if ErrorIn(g0348UU) {void_try9 = g0348UU
                    } else {
                    void_try9 = l.WriteEID(i,g0348UU)}
                    /* LetEID-9 */} 
                  /* ERROR PROTECTION INSERTED (void_try9-void_try9) */
                  if ErrorIn(void_try9) {Result = void_try9
                  break
                  } else {
                  i = (i+1)
                  }
                  /* while-8 */} 
                }
                /* Let-7 */} 
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            Result = EID{l.Id(),0}
            }
            /* Let-5 */} 
          } else {
          /* Let:5 */{ 
            var l *ClaireList   = CreateList(ToType(CEMPTY.Id()),n)
            /* noccur = 2 */
            /* Let:6 */{ 
              var i int  = 1
              /* noccur = 5 */
              /* Let:7 */{ 
                var g0347 int  = n
                /* noccur = 1 */
                Result= EID{CFALSE.Id(),0}
                for (i <= g0347) /* while:8 */{ 
                  var void_try9 EID 
                  _ = void_try9
                  { 
                  /* Let:9 */{ 
                    var g0349UU *ClaireAny  
                    /* noccur = 1 */
                    var g0349UU_try035010 EID 
                    g0349UU_try035010 = EVAL(self.Args.At(i-1))
                    /* ERROR PROTECTION INSERTED (g0349UU-void_try9) */
                    if ErrorIn(g0349UU_try035010) {void_try9 = g0349UU_try035010
                    } else {
                    g0349UU = ANY(g0349UU_try035010)
                    void_try9 = ToArray(l.Id()).NthPut(i,g0349UU).ToEID()
                    }
                    /* Let-9 */} 
                  /* ERROR PROTECTION INSERTED (void_try9-void_try9) */
                  if ErrorIn(void_try9) {Result = void_try9
                  break
                  } else {
                  i = (i+1)
                  }
                  /* while-8 */} 
                }
                /* Let-7 */} 
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            Result = EID{l.Id(),0}
            }
            /* Let-5 */} 
          /* If-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: self_eval @ List (throw: true) 
func E_self_eval_List (self EID) EID { 
    return /*(sm for self_eval @ List= EID)*/ To_List(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: List 
func EVAL_List (x *ClaireAny) EID { 
     return To_List(x).SelfEval()} 
  
// here we use the CLAIRE 3 style of post-typing with a cast! 
/* {1} OPT.The go function for: self_eval(self:Set) [] */
func (self *Set ) SelfEval () EID { 
    var Result EID 
    /* Let:2 */{ 
      var s *ClaireSet  
      /* noccur = 3 */
      var s_try03513 EID 
      /* Let:3 */{ 
        var x_bag *ClaireSet   = ToType(CEMPTY.Id()).EmptySet()
        /* noccur = 2 */
        /* For:4 */{ 
          var x *ClaireAny  
          _ = x
          s_try03513= EID{CFALSE.Id(),0}
          var x_support *ClaireList  
          x_support = self.Args
          x_len := x_support.Length()
          for i_it := 0; i_it < x_len; i_it++ { 
            x = x_support.At(i_it)
            var void_try6 EID 
            _ = void_try6
            /* Let:6 */{ 
              var g0352UU *ClaireAny  
              /* noccur = 1 */
              var g0352UU_try03537 EID 
              g0352UU_try03537 = EVAL(x)
              /* ERROR PROTECTION INSERTED (g0352UU-void_try6) */
              if ErrorIn(g0352UU_try03537) {void_try6 = g0352UU_try03537
              } else {
              g0352UU = ANY(g0352UU_try03537)
              void_try6 = EID{x_bag.AddFast(g0352UU).Id(),0}
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (void_try6-s_try03513) */
            if ErrorIn(void_try6) {s_try03513 = void_try6
            s_try03513 = void_try6
            break
            } else {
            }
            /* loop-5 */} 
          /* For-4 */} 
        /* ERROR PROTECTION INSERTED (s_try03513-s_try03513) */
        if !ErrorIn(s_try03513) {
        s_try03513 = EID{x_bag.Id(),0}
        }
        /* Let-3 */} 
      /* ERROR PROTECTION INSERTED (s-Result) */
      if ErrorIn(s_try03513) {Result = s_try03513
      } else {
      s = ToSet(OBJ(s_try03513))
      if (self.Of.Id() != CNULL) /* If:3 */{ 
        /* Let:4 */{ 
          var x *ClaireAny  
          /* noccur = 2 */
          /* Let:5 */{ 
            var x_some *ClaireAny   = CNULL
            /* noccur = 2 */
            /* For:6 */{ 
              var x *ClaireAny  
              _ = x
              for _,x = range(s.Values)/* loop:7 */{ 
                if (self.Of.Contains(x) != CTRUE) /* If:8 */{ 
                   /*v = x, s =void*/
x_some = x
                  break
                  /* If-8 */} 
                /* loop-7 */} 
              /* For-6 */} 
            x = x_some
            /* Let-5 */} 
          if (x != CNULL) /* If:5 */{ 
            Result = ToException(Core.C_range_error.Make(self.Id(),x,self.Of.Id())).Close()
            } else {
            Result = EID{CNULL,0}
            /* If-5 */} 
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        Result = EID{s.Cast_I(self.Of).Id(),0}
        }
        } else {
        Result = EID{s.Cast_I(ToType(CEMPTY.Id())).Id(),0}
        /* If-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: self_eval @ Set (throw: true) 
func E_self_eval_Set (self EID) EID { 
    return /*(sm for self_eval @ Set= EID)*/ To_Set(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Set 
func EVAL_Set (x *ClaireAny) EID { 
     return To_Set(x).SelfEval()} 
  
/* {1} OPT.The go function for: self_eval(self:Tuple) [] */
func (self *Tuple ) SelfEval () EID { 
    var Result EID 
    /* Let:2 */{ 
      var g0354UU *ClaireList  
      /* noccur = 1 */
      var g0354UU_try03553 EID 
      /* Iteration:3 */{ 
        var v_list3 *ClaireList  
        var x *ClaireAny  
        var v_local3 *ClaireAny  
        v_list3 = self.Args
        g0354UU_try03553 = EID{CreateList(ToType(CEMPTY.Id()),v_list3.Length()).Id(),0}
        for CLcount := 0; CLcount < v_list3.Length(); CLcount++{ 
          x = v_list3.At(CLcount)
          var v_local3_try03565 EID 
          v_local3_try03565 = EVAL(x)
          /* ERROR PROTECTION INSERTED (v_local3-g0354UU_try03553) */
          if ErrorIn(v_local3_try03565) {g0354UU_try03553 = v_local3_try03565
          g0354UU_try03553 = v_local3_try03565
          break
          } else {
          v_local3 = ANY(v_local3_try03565)
          ToList(OBJ(g0354UU_try03553)).PutAt(CLcount,v_local3)
          } 
        }
        /* Iteration-3 */} 
      /* ERROR PROTECTION INSERTED (g0354UU-Result) */
      if ErrorIn(g0354UU_try03553) {Result = g0354UU_try03553
      } else {
      g0354UU = ToList(OBJ(g0354UU_try03553))
      Result = EID{g0354UU.Tuple_I().Id(),0}
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: self_eval @ Tuple (throw: true) 
func E_self_eval_Tuple (self EID) EID { 
    return /*(sm for self_eval @ Tuple= EID)*/ To_Tuple(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Tuple 
func EVAL_Tuple (x *ClaireAny) EID { 
     return To_Tuple(x).SelfEval()} 
  
// same as creating a list (same constraints since same underlying structure)
/* {1} OPT.The go function for: self_eval(self:Array) [] */
func (self *Array ) SelfEval () EID { 
    var Result EID 
    /* Let:2 */{ 
      var type_ask *ClaireBoolean   = MakeBoolean((self.Of.Id() == CNULL)).Not
      /* noccur = 1 */
      /* Let:3 */{ 
        var n int  = self.Args.Length()
        /* noccur = 4 */
        if (type_ask == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var l *ClaireList   = CreateList(self.Of,n)
            /* noccur = 2 */
            /* Let:6 */{ 
              var i int  = 1
              /* noccur = 5 */
              /* Let:7 */{ 
                var g0357 int  = n
                /* noccur = 1 */
                Result= EID{CFALSE.Id(),0}
                for (i <= g0357) /* while:8 */{ 
                  var void_try9 EID 
                  _ = void_try9
                  { 
                  /* LetEID:9 */{ 
                    var g0359UU EID 
                    g0359UU = EVAL(self.Args.At(i-1))
                    /* ERROR PROTECTION INSERTED (g0359UU-void_try9) */
                    if ErrorIn(g0359UU) {void_try9 = g0359UU
                    } else {
                    void_try9 = l.WriteEID(i,g0359UU)}
                    /* LetEID-9 */} 
                  /* ERROR PROTECTION INSERTED (void_try9-void_try9) */
                  if ErrorIn(void_try9) {Result = void_try9
                  break
                  } else {
                  i = (i+1)
                  }
                  /* while-8 */} 
                }
                /* Let-7 */} 
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            Result = EID{l.Array_I().Id(),0}
            }
            /* Let-5 */} 
          } else {
          /* Let:5 */{ 
            var l *ClaireList   = CreateList(ToType(CEMPTY.Id()),n)
            /* noccur = 2 */
            /* Let:6 */{ 
              var i int  = 1
              /* noccur = 5 */
              /* Let:7 */{ 
                var g0358 int  = n
                /* noccur = 1 */
                Result= EID{CFALSE.Id(),0}
                for (i <= g0358) /* while:8 */{ 
                  var void_try9 EID 
                  _ = void_try9
                  { 
                  /* Let:9 */{ 
                    var g0360UU *ClaireAny  
                    /* noccur = 1 */
                    var g0360UU_try036110 EID 
                    g0360UU_try036110 = EVAL(self.Args.At(i-1))
                    /* ERROR PROTECTION INSERTED (g0360UU-void_try9) */
                    if ErrorIn(g0360UU_try036110) {void_try9 = g0360UU_try036110
                    } else {
                    g0360UU = ANY(g0360UU_try036110)
                    void_try9 = ToArray(l.Id()).NthPut(i,g0360UU).ToEID()
                    }
                    /* Let-9 */} 
                  /* ERROR PROTECTION INSERTED (void_try9-void_try9) */
                  if ErrorIn(void_try9) {Result = void_try9
                  break
                  } else {
                  i = (i+1)
                  }
                  /* while-8 */} 
                }
                /* Let-7 */} 
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            Result = EID{l.Array_I().Id(),0}
            }
            /* Let-5 */} 
          /* If-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: self_eval @ Array (throw: true) 
func E_self_eval_Array2 (self EID) EID { 
    return /*(sm for self_eval @ Array= EID)*/ To_Array(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Array 
func EVAL_Array (x *ClaireAny) EID { 
     return To_Array(x).SelfEval()} 
  
// Macros are a nice but undocumented feature of CLAIRE. This is deliberate :)
// it is an advanced feature for those who want to expand the language. This
// makes CLAIRE a nice framework for DSL
//
/* {1} OPT.The go function for: self_eval(self:Macro) [] */
func (self *Macro ) SelfEval () EID { 
    var Result EID 
    /* Let:2 */{ 
      var g0362UU *ClaireAny  
      /* noccur = 1 */
      var g0362UU_try03633 EID 
      g0362UU_try03633 = Core.F_CALL(C_macroexpand,ARGS(EID{self.Id(),0}))
      /* ERROR PROTECTION INSERTED (g0362UU-Result) */
      if ErrorIn(g0362UU_try03633) {Result = g0362UU_try03633
      } else {
      g0362UU = ANY(g0362UU_try03633)
      Result = EVAL(g0362UU)
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: self_eval @ Macro (throw: true) 
func E_self_eval_Macro2 (self EID) EID { 
    return /*(sm for self_eval @ Macro= EID)*/ To_Macro(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Macro 
func EVAL_Macro (x *ClaireAny) EID { 
     return To_Macro(x).SelfEval()} 
  
// error produces an exception of type general_error
/* {1} OPT.The go function for: self_eval(self:Error) [] */
func (self *Error ) SelfEval () EID { 
    var Result EID 
    if ((F_boolean_I_any(self.Args.Id()).Id() != CTRUE.Id()) || 
        (C_string.Id() != self.Args.At(1-1).Isa.Id())) /* If:2 */{ 
      Result = ToException(Core.C_general_error.Make(MakeString("Syntax error: ~S").Id(),MakeConstantList(self.Id()).Id())).Close()
      } else {
      Result = EID{CFALSE.Id(),0}
      /* If-2 */} 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    /* Let:2 */{ 
      var x *Core.GeneralError   = Core.ToGeneralError(new(Core.GeneralError).Is(Core.C_general_error))
      /* noccur = 5 */
      /* update:3 */{ 
        var va_arg1 *Core.GeneralError  
        var va_arg2 *ClaireAny  
        va_arg1 = x
        var va_arg2_try03644 EID 
        va_arg2_try03644 = Core.F_car_list(self.Args)
        /* ERROR PROTECTION INSERTED (va_arg2-Result) */
        if ErrorIn(va_arg2_try03644) {Result = va_arg2_try03644
        } else {
        va_arg2 = ANY(va_arg2_try03644)
        /* ---------- now we compile update mClaire/cause(va_arg1) := va_arg2 ------- */
        va_arg1.Cause = va_arg2
        Result = va_arg2.ToEID()
        }
        /* update-3 */} 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      /* update:3 */{ 
        var va_arg1 *Core.GeneralError  
        var va_arg2 *ClaireAny  
        va_arg1 = x
        var va_arg2_try03654 EID 
        /* Iteration:4 */{ 
          var v_list4 *ClaireList  
          var x *ClaireAny  
          var v_local4 *ClaireAny  
          var v_list4_try03665 EID 
          v_list4_try03665 = self.Args.Cdr()
          /* ERROR PROTECTION INSERTED (v_list4-va_arg2_try03654) */
          if ErrorIn(v_list4_try03665) {va_arg2_try03654 = v_list4_try03665
          } else {
          v_list4 = ToList(OBJ(v_list4_try03665))
          va_arg2_try03654 = EID{CreateList(ToType(CEMPTY.Id()),v_list4.Length()).Id(),0}
          for CLcount := 0; CLcount < v_list4.Length(); CLcount++{ 
            x = v_list4.At(CLcount)
            var v_local4_try03676 EID 
            v_local4_try03676 = EVAL(x)
            /* ERROR PROTECTION INSERTED (v_local4-va_arg2_try03654) */
            if ErrorIn(v_local4_try03676) {va_arg2_try03654 = v_local4_try03676
            va_arg2_try03654 = v_local4_try03676
            break
            } else {
            v_local4 = ANY(v_local4_try03676)
            ToList(OBJ(va_arg2_try03654)).PutAt(CLcount,v_local4)
            } 
          }}
          /* Iteration-4 */} 
        /* ERROR PROTECTION INSERTED (va_arg2-Result) */
        if ErrorIn(va_arg2_try03654) {Result = va_arg2_try03654
        } else {
        va_arg2 = ANY(va_arg2_try03654)
        /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
        va_arg1.Arg = va_arg2
        Result = va_arg2.ToEID()
        }
        /* update-3 */} 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = x.Close()
      }}
      /* Let-2 */} 
    }
    return RangeCheck(ToType(C_error.Id()),Result)} 
  
// The EID go function for: self_eval @ Error (throw: true) 
func E_self_eval_Error (self EID) EID { 
    return /*(sm for self_eval @ Error= EID)*/ To_Error(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Error 
func EVAL_Error (x *ClaireAny) EID { 
     return To_Error(x).SelfEval()} 
  
// this is the basic tool for printing in CLAIRE. A complex statement
// is macroexpanded into basic printing instructions
//
/* {1} OPT.The go function for: self_eval(self:Printf) [] */
func (self *Printf ) SelfEval () EID { 
    var Result EID 
    /* Let:2 */{ 
      var l *ClaireList   = self.Args
      /* noccur = 6 */
      /* Let:3 */{ 
        var s *ClaireAny   = l.At(1-1)
        /* noccur = 12 */
        if (C_string.Id() != s.Isa.Id()) /* If:4 */{ 
          Result = ToException(Core.C_general_error.Make(MakeString("[102] the first argument in ~S must be a string").Id(),MakeConstantList(self.Id()).Id())).Close()
          } else {
          /* Let:5 */{ 
            var i int  = 2
            /* noccur = 7 */
            /* Let:6 */{ 
              var n int  = F_get_string(ToString(s),'~')
              /* noccur = 14 */
              Result= EID{CFALSE.Id(),0}
              for (n != 0) /* while:7 */{ 
                var void_try8 EID 
                _ = void_try8
                /* Let:8 */{ 
                  var m *ClaireAny  
                  /* noccur = 4 */
                  var m_try03689 EID 
                  m_try03689 = Core.F_CALL(C_nth,ARGS(s.ToEID(),EID{C__INT,IVAL((n+1))}))
                  /* ERROR PROTECTION INSERTED (m-void_try8) */
                  if ErrorIn(m_try03689) {void_try8 = m_try03689
                  } else {
                  m = ANY(m_try03689)
                  if (i > l.Length()) /* If:9 */{ 
                    void_try8 = ToException(Core.C_general_error.Make(MakeString("[103] not enough arguments in ~S").Id(),MakeConstantList(self.Id()).Id())).Close()
                    } else {
                    void_try8 = EID{CFALSE.Id(),0}
                    /* If-9 */} 
                  /* ERROR PROTECTION INSERTED (void_try8-void_try8) */
                  if ErrorIn(void_try8) {Result = void_try8
                  break
                  } else {
                  if (n > 1) /* If:9 */{ 
                    F_princ_string(F_substring_string(ToString(s),1,(n-1)))
                    /* If-9 */} 
                  if ('A' == ToChar(m).Value) /* If:9 */{ 
                    /* Let:10 */{ 
                      var g0369UU *ClaireAny  
                      /* noccur = 1 */
                      var g0369UU_try037011 EID 
                      g0369UU_try037011 = EVAL(l.At(i-1))
                      /* ERROR PROTECTION INSERTED (g0369UU-void_try8) */
                      if ErrorIn(g0369UU_try037011) {void_try8 = g0369UU_try037011
                      } else {
                      g0369UU = ANY(g0369UU_try037011)
                      void_try8 = Core.F_CALL(C_princ,ARGS(g0369UU.ToEID()))
                      }
                      /* Let-10 */} 
                    /* If!9 */}  else if ('S' == ToChar(m).Value) /* If:9 */{ 
                    /* Let:10 */{ 
                      var g0371UU *ClaireAny  
                      /* noccur = 1 */
                      var g0371UU_try037211 EID 
                      g0371UU_try037211 = EVAL(l.At(i-1))
                      /* ERROR PROTECTION INSERTED (g0371UU-void_try8) */
                      if ErrorIn(g0371UU_try037211) {void_try8 = g0371UU_try037211
                      } else {
                      g0371UU = ANY(g0371UU_try037211)
                      void_try8 = Core.F_CALL(C_print,ARGS(g0371UU.ToEID()))
                      }
                      /* Let-10 */} 
                    /* If!9 */}  else if ('F' == ToChar(m).Value) /* If:9 */{ 
                    /* Let:10 */{ 
                      var fv *ClaireAny  
                      /* noccur = 5 */
                      var fv_try037311 EID 
                      fv_try037311 = EVAL(l.At(i-1))
                      /* ERROR PROTECTION INSERTED (fv-void_try8) */
                      if ErrorIn(fv_try037311) {void_try8 = fv_try037311
                      } else {
                      fv = ANY(fv_try037311)
                      /* Let:11 */{ 
                        var p_Z *ClaireBoolean   = CFALSE
                        /* noccur = 4 */
                        /* Let:12 */{ 
                          var j int 
                          /* noccur = 4 */
                          var j_try037413 EID 
                          /* Let:13 */{ 
                            var g0375UU int 
                            /* noccur = 1 */
                            var g0375UU_try037614 EID 
                            /* Let:14 */{ 
                              var g0377UU rune 
                              /* noccur = 1 */
                              var g0377UU_try037815 EID 
                              g0377UU_try037815 = Core.F_nth_get_string(ToString(s),(n+2),(n+2))
                              /* ERROR PROTECTION INSERTED (g0377UU-g0375UU_try037614) */
                              if ErrorIn(g0377UU_try037815) {g0375UU_try037614 = g0377UU_try037815
                              } else {
                              g0377UU = CHAR(g0377UU_try037815)
                              g0375UU_try037614 = EID{C__INT,IVAL(int(g0377UU))}
                              }
                              /* Let-14 */} 
                            /* ERROR PROTECTION INSERTED (g0375UU-j_try037413) */
                            if ErrorIn(g0375UU_try037614) {j_try037413 = g0375UU_try037614
                            } else {
                            g0375UU = INT(g0375UU_try037614)
                            j_try037413 = EID{C__INT,IVAL((g0375UU-48))}
                            }
                            /* Let-13 */} 
                          /* ERROR PROTECTION INSERTED (j-void_try8) */
                          if ErrorIn(j_try037413) {void_try8 = j_try037413
                          } else {
                          j = INT(j_try037413)
                          var g0379I *ClaireBoolean  
                          var g0379I_try038013 EID 
                          /* Let:13 */{ 
                            var g0381UU *ClaireAny  
                            /* noccur = 1 */
                            var g0381UU_try038214 EID 
                            g0381UU_try038214 = Core.F_CALL(C_nth,ARGS(s.ToEID(),EID{C__INT,IVAL((n+2))}))
                            /* ERROR PROTECTION INSERTED (g0381UU-g0379I_try038013) */
                            if ErrorIn(g0381UU_try038214) {g0379I_try038013 = g0381UU_try038214
                            } else {
                            g0381UU = ANY(g0381UU_try038214)
                            g0379I_try038013 = EID{Equal(MakeChar('%').Id(),g0381UU).Id(),0}
                            }
                            /* Let-13 */} 
                          /* ERROR PROTECTION INSERTED (g0379I-void_try8) */
                          if ErrorIn(g0379I_try038013) {void_try8 = g0379I_try038013
                          } else {
                          g0379I = ToBoolean(OBJ(g0379I_try038013))
                          if (g0379I == CTRUE) /* If:13 */{ 
                            p_Z = CTRUE
                            j = 1
                            fv = ANY(Core.F_CALL(ToProperty(C__star.Id()),ARGS(fv.ToEID(),EID{C__FLOAT,FVAL(100)})))
                            void_try8 = fv.ToEID()
                            /* If!13 */}  else if ((j < 0) || 
                              (j > 9)) /* If:13 */{ 
                            void_try8 = ToException(Core.C_general_error.Make(MakeString("[189] F requires a single digit integer in ~S").Id(),MakeConstantList(self.Id()).Id())).Close()
                            } else {
                            void_try8 = EID{CFALSE.Id(),0}
                            /* If-13 */} 
                          }
                          /* ERROR PROTECTION INSERTED (void_try8-void_try8) */
                          if ErrorIn(void_try8) {Result = void_try8
                          break
                          } else {
                          var g0383I *ClaireBoolean  
                          var g0383I_try038413 EID 
                          /* and:13 */{ 
                            var v_and13 *ClaireBoolean  
                            
                            v_and13 = p_Z.Not
                            if (v_and13 == CFALSE) {g0383I_try038413 = EID{CFALSE.Id(),0}
                            } else /* arg:14 */{ 
                              var v_and13_try038515 EID 
                              /* Let:15 */{ 
                                var g0386UU *ClaireAny  
                                /* noccur = 1 */
                                var g0386UU_try038716 EID 
                                g0386UU_try038716 = Core.F_CALL(C_nth,ARGS(s.ToEID(),EID{C__INT,IVAL((n+3))}))
                                /* ERROR PROTECTION INSERTED (g0386UU-v_and13_try038515) */
                                if ErrorIn(g0386UU_try038716) {v_and13_try038515 = g0386UU_try038716
                                } else {
                                g0386UU = ANY(g0386UU_try038716)
                                v_and13_try038515 = EID{Equal(MakeChar('%').Id(),g0386UU).Id(),0}
                                }
                                /* Let-15 */} 
                              /* ERROR PROTECTION INSERTED (v_and13-g0383I_try038413) */
                              if ErrorIn(v_and13_try038515) {g0383I_try038413 = v_and13_try038515
                              } else {
                              v_and13 = ToBoolean(OBJ(v_and13_try038515))
                              if (v_and13 == CFALSE) {g0383I_try038413 = EID{CFALSE.Id(),0}
                              } else /* arg:15 */{ 
                                g0383I_try038413 = EID{CTRUE.Id(),0}/* arg-15 */} 
                              /* arg-14 */} 
                            }
                            /* and-13 */} 
                          /* ERROR PROTECTION INSERTED (g0383I-void_try8) */
                          if ErrorIn(g0383I_try038413) {void_try8 = g0383I_try038413
                          } else {
                          g0383I = ToBoolean(OBJ(g0383I_try038413))
                          if (g0383I == CTRUE) /* If:13 */{ 
                            p_Z = CTRUE
                            fv = ANY(Core.F_CALL(ToProperty(C__star.Id()),ARGS(fv.ToEID(),EID{C__FLOAT,FVAL(100)})))
                            n = (n+1)
                            void_try8 = EID{C__INT,IVAL(n)}
                            } else {
                            void_try8 = EID{CFALSE.Id(),0}
                            /* If-13 */} 
                          }
                          /* ERROR PROTECTION INSERTED (void_try8-void_try8) */
                          if ErrorIn(void_try8) {Result = void_try8
                          break
                          } else {
                          void_try8 = Core.F_CALL(Core.C_mClaire_printFDigit,ARGS(fv.ToEID(),EID{C__INT,IVAL(j)}))
                          /* ERROR PROTECTION INSERTED (void_try8-void_try8) */
                          if ErrorIn(void_try8) {Result = void_try8
                          break
                          } else {
                          if (p_Z == CTRUE) /* If:13 */{ 
                            PRINC("%")
                            /* If-13 */} 
                          n = (n+1)
                          void_try8 = EID{C__INT,IVAL(n)}
                          }}}
                          }
                          /* Let-12 */} 
                        /* Let-11 */} 
                      }
                      /* Let-10 */} 
                    /* If!9 */}  else if ('I' == ToChar(m).Value) /* If:9 */{ 
                    void_try8 = EVAL(l.At(i-1))
                    } else {
                    void_try8 = EID{CFALSE.Id(),0}
                    /* If-9 */} 
                  /* ERROR PROTECTION INSERTED (void_try8-void_try8) */
                  if ErrorIn(void_try8) {Result = void_try8
                  break
                  } else {
                  i = (i+1)
                  s = (F_substring_string(ToString(s),(n+2),1000)).Id()
                  var n_try03889 EID 
                  n_try03889 = Core.F_CALL(C_get,ARGS(s.ToEID(),EID{C__CHAR,CVAL('~')}))
                  /* ERROR PROTECTION INSERTED (n-void_try8) */
                  if ErrorIn(n_try03889) {void_try8 = n_try03889
                  Result = n_try03889
                  break
                  } else {
                  n = INT(n_try03889)
                  void_try8 = EID{C__INT,IVAL(n)}
                  }}}
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (void_try8-Result) */
                if ErrorIn(void_try8) {Result = void_try8
                Result = void_try8
                break
                } else {
                /* while-7 */} 
              }
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              if (F_boolean_I_any(s) == CTRUE) /* If:7 */{ 
                Result = Core.F_CALL(C_princ,ARGS(s.ToEID()))
                } else {
                Result = EID{CFALSE.Id(),0}
                /* If-7 */} 
              }
              /* Let-6 */} 
            /* Let-5 */} 
          /* If-4 */} 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        Result = EID{CNULL,0}
        }
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: self_eval @ Printf (throw: true) 
func E_self_eval_Printf (self EID) EID { 
    return /*(sm for self_eval @ Printf= EID)*/ To_Printf(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Printf 
func EVAL_Printf (x *ClaireAny) EID { 
     return To_Printf(x).SelfEval()} 
  
// trace is refined in inspect.cl
// If trace_output() is known, use it, else use current output.
// defined in inspect.cl
// CLAIRE4: self_eval is defined once for all, hence exteneded
/* {1} OPT.The go function for: self_eval(self:Trace) [] */
func (self *Trace ) SelfEval () EID { 
    var Result EID 
    /* Let:2 */{ 
      var a *ClaireList   = self.Args
      /* noccur = 5 */
      /* Let:3 */{ 
        var l *ClaireList  
        /* noccur = 2 */
        var l_try03904 EID 
        /* Iteration:4 */{ 
          var v_list4 *ClaireList  
          var x *ClaireAny  
          var v_local4 *ClaireAny  
          v_list4 = a
          l_try03904 = EID{CreateList(ToType(CEMPTY.Id()),v_list4.Length()).Id(),0}
          for CLcount := 0; CLcount < v_list4.Length(); CLcount++{ 
            x = v_list4.At(CLcount)
            var v_local4_try03916 EID 
            v_local4_try03916 = EVAL(x)
            /* ERROR PROTECTION INSERTED (v_local4-l_try03904) */
            if ErrorIn(v_local4_try03916) {l_try03904 = v_local4_try03916
            l_try03904 = v_local4_try03916
            break
            } else {
            v_local4 = ANY(v_local4_try03916)
            ToList(OBJ(l_try03904)).PutAt(CLcount,v_local4)
            } 
          }
          /* Iteration-4 */} 
        /* ERROR PROTECTION INSERTED (l-Result) */
        if ErrorIn(l_try03904) {Result = l_try03904
        } else {
        l = ToList(OBJ(l_try03904))
        /* Let:4 */{ 
          var i *ClaireAny   = l.At(1-1)
          /* noccur = 2 */
          /* Let:5 */{ 
            var a2 *ClaireAny  
            /* noccur = 2 */
            if (a.Length() > 1) /* If:6 */{ 
              a2 = a.At(2-1)
              } else {
              a2 = CFALSE.Id()
              /* If-6 */} 
            if (a.Length() == 1) /* If:6 */{ 
              /* Let:7 */{ 
                var a1 *ClaireAny  
                /* noccur = 1 */
                var a1_try03928 EID 
                a1_try03928 = EVAL(a.At(1-1))
                /* ERROR PROTECTION INSERTED (a1-Result) */
                if ErrorIn(a1_try03928) {Result = a1_try03928
                } else {
                a1 = ANY(a1_try03928)
                /* Let:8 */{ 
                  var p *ClaireProperty   = C_iClaire_trace_on
                  /* noccur = 2 */
                  if (p.Restrictions.Length() != 0) /* If:9 */{ 
                    if (ClEnv.Trace_I == 0) /* If:10 */{ 
                      ClEnv.Trace_I = 1
                      /* If-10 */} 
                    Result = Core.F_CALL(Core.C_call,ARGS(EID{p.Id(),0},a1.ToEID()))
                    } else {
                    Result = EID{CFALSE.Id(),0}
                    /* If-9 */} 
                  /* Let-8 */} 
                }
                /* Let-7 */} 
              } else {
              var g0393I *ClaireBoolean  
              /* and:7 */{ 
                var v_and7 *ClaireBoolean  
                
                v_and7 = Equal(C_string.Id(),a2.Isa.Id())
                if (v_and7 == CFALSE) {g0393I = CFALSE
                } else /* arg:8 */{ 
                  if (C_integer.Id() == i.Isa.Id()) /* If:9 */{ 
                    /* Let:10 */{ 
                      var g0389 int  = ToInteger(i).Value
                      /* noccur = 1 */
                      v_and7 = Core.F__inf_equal_integer(g0389,ClEnv.Verbose)
                      /* Let-10 */} 
                    } else {
                    v_and7 = CFALSE
                    /* If-9 */} 
                  if (v_and7 == CFALSE) {g0393I = CFALSE
                  } else /* arg:9 */{ 
                    g0393I = CTRUE/* arg-9 */} 
                  /* arg-8 */} 
                /* and-7 */} 
              if (g0393I == CTRUE) /* If:7 */{ 
                /* Let:8 */{ 
                  var p *ClaireAny   = Core.F_get_property(C_ctrace,ToObject(ClEnv.Id()))
                  /* noccur = 5 */
                  if (p != CNULL) /* If:9 */{ 
                    p = ToPort(p).UseAsOutput().Id()
                    /* If-9 */} 
                  Result = Core.F_format_string(ToString(a2),l.Skip(2))
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  if (p != CNULL) /* If:9 */{ 
                    ToPort(p).UseAsOutput()
                    /* If-9 */} 
                  Result = EID{CEMPTY.Id(),0}
                  }
                  /* Let-8 */} 
                } else {
                Result = EID{CFALSE.Id(),0}
                /* If-7 */} 
              /* If-6 */} 
            /* Let-5 */} 
          /* Let-4 */} 
        }
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: self_eval @ Trace (throw: true) 
func E_self_eval_Trace (self EID) EID { 
    return /*(sm for self_eval @ Trace= EID)*/ To_Trace(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Trace 
func EVAL_Trace (x *ClaireAny) EID { 
     return To_Trace(x).SelfEval()} 
  
// assert is refined in trace.la
//
/* {1} OPT.The go function for: self_eval(self:Assert) [] */
func (self *Assert ) SelfEval () EID { 
    var Result EID 
    /* Let:2 */{ 
      var a *ClaireList   = self.Args
      /* noccur = 3 */
      var g0394I *ClaireBoolean  
      var g0394I_try03953 EID 
      /* and:3 */{ 
        var v_and3 *ClaireBoolean  
        
        v_and3 = Core.F__sup_integer(a.Length(),0)
        if (v_and3 == CFALSE) {g0394I_try03953 = EID{CFALSE.Id(),0}
        } else /* arg:4 */{ 
          v_and3 = Core.F_known_ask_any(Core.F_get_property(C_ctrace,ToObject(ClEnv.Id())))
          if (v_and3 == CFALSE) {g0394I_try03953 = EID{CFALSE.Id(),0}
          } else /* arg:5 */{ 
            var v_and3_try03966 EID 
            /* Let:6 */{ 
              var g0397UU *ClaireBoolean  
              /* noccur = 1 */
              var g0397UU_try03987 EID 
              /* Let:7 */{ 
                var g0399UU *ClaireAny  
                /* noccur = 1 */
                var g0399UU_try04008 EID 
                g0399UU_try04008 = EVAL(a.At(1-1))
                /* ERROR PROTECTION INSERTED (g0399UU-g0397UU_try03987) */
                if ErrorIn(g0399UU_try04008) {g0397UU_try03987 = g0399UU_try04008
                } else {
                g0399UU = ANY(g0399UU_try04008)
                g0397UU_try03987 = EID{F_boolean_I_any(g0399UU).Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (g0397UU-v_and3_try03966) */
              if ErrorIn(g0397UU_try03987) {v_and3_try03966 = g0397UU_try03987
              } else {
              g0397UU = ToBoolean(OBJ(g0397UU_try03987))
              v_and3_try03966 = EID{Core.F__I_equal_any(g0397UU.Id(),CTRUE.Id()).Id(),0}
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (v_and3-g0394I_try03953) */
            if ErrorIn(v_and3_try03966) {g0394I_try03953 = v_and3_try03966
            } else {
            v_and3 = ToBoolean(OBJ(v_and3_try03966))
            if (v_and3 == CFALSE) {g0394I_try03953 = EID{CFALSE.Id(),0}
            } else /* arg:6 */{ 
              g0394I_try03953 = EID{CTRUE.Id(),0}/* arg-6 */} 
            /* arg-5 */} 
          /* arg-4 */} 
        }
        /* and-3 */} 
      /* ERROR PROTECTION INSERTED (g0394I-Result) */
      if ErrorIn(g0394I_try03953) {Result = g0394I_try03953
      } else {
      g0394I = ToBoolean(OBJ(g0394I_try03953))
      if (g0394I == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var p *ClairePort   = ClEnv.Ctrace.UseAsOutput()
          /* noccur = 1 */
          Result = Core.F_print_any((self.External).Id())
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(",line=")
          F_princ_integer(self.Index)
          PRINC(": (ASSERT) ")
          Result = Core.F_CALL(C_print,ARGS(a.At(1-1).ToEID()))
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC("\n")
          Result = EVOID
          }}
          {
          p.UseAsOutput()
          if (ClEnv.Debug_I >= 0) /* If:5 */{ 
            Result = ToException(Core.C_general_error.Make(MakeString("Assertion Violation").Id(),CNIL.Id())).Close()
            } else {
            Result = EID{CFALSE.Id(),0}
            /* If-5 */} 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          Result = EID{CEMPTY.Id(),0}
          }}
          /* Let-4 */} 
        } else {
        Result = EID{CFALSE.Id(),0}
        /* If-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: self_eval @ Assert (throw: true) 
func E_self_eval_Assert (self EID) EID { 
    return /*(sm for self_eval @ Assert= EID)*/ To_Assert(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Assert 
func EVAL_Assert (x *ClaireAny) EID { 
     return To_Assert(x).SelfEval()} 
  
/* {1} OPT.The go function for: self_eval(self:Branch) [] */
func (self *Branch ) SelfEval () EID { 
    var Result EID 
    if (self.Args.Length() != 1) /* If:2 */{ 
      Result = ToException(Core.C_general_error.Make(MakeString("[104] Syntax error with ~S (one arg. expected)").Id(),MakeConstantList(self.Id()).Id())).Close()
      } else {
      Result = EID{CFALSE.Id(),0}
      /* If-2 */} 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    h_index := ClEnv.Index /* Handle */
    h_base := ClEnv.Base
    F_world_push()
    var g0401I *ClaireBoolean  
    var g0401I_try04022 EID 
    /* Let:2 */{ 
      var g0403UU *ClaireAny  
      /* noccur = 1 */
      var g0403UU_try04043 EID 
      g0403UU_try04043 = EVAL(self.Args.At(1-1))
      /* ERROR PROTECTION INSERTED (g0403UU-g0401I_try04022) */
      if ErrorIn(g0403UU_try04043) {g0401I_try04022 = g0403UU_try04043
      } else {
      g0403UU = ANY(g0403UU_try04043)
      g0401I_try04022 = EID{Core.F__I_equal_any(g0403UU,CFALSE.Id()).Id(),0}
      }
      /* Let-2 */} 
    /* ERROR PROTECTION INSERTED (g0401I-Result) */
    if ErrorIn(g0401I_try04022) {Result = g0401I_try04022
    } else {
    g0401I = ToBoolean(OBJ(g0401I_try04022))
    if (g0401I == CTRUE) /* If:2 */{ 
      Result = EID{CTRUE.Id(),0}
      } else {
      F_world_pop()
      Result = EID{CFALSE.Id(),0}
      /* If-2 */} 
    }
    if ErrorIn(Result) && ToType(Core.C_contradiction.Id()).Contains(ANY(Result)) == CTRUE { 
      /* s=EID */ClEnv.Index = h_index
      ClEnv.Base = h_base
      F_world_pop()
      Result = EID{CFALSE.Id(),0}
      } 
    }
    return Result} 
  
// The EID go function for: self_eval @ Branch (throw: true) 
func E_self_eval_Branch (self EID) EID { 
    return /*(sm for self_eval @ Branch= EID)*/ To_Branch(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Branch 
func EVAL_Branch (x *ClaireAny) EID { 
     return To_Branch(x).SelfEval()} 
  
// end of file