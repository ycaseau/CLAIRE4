/***** CLAIRE Compilation of file /Users/ycaseau/Dropbox/src/clairev4.0/src/meta/control.cl 
         [version 4.0.01 / safety 5] Saturday 10-30-2021 *****/

package Language
import (_ "fmt"
	. "Kernel"
	"Core"
)

//-------- dumb function to prevent import errors --------
func import_g0236() { 
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
          var g0242 *If   = To_If(e)
          /* noccur = 3 */
          PRINC(" ")
          Result = F_lbreak_void()
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC("else if ")
          Result = F_printexp_any(g0242.Test,CFALSE)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(" ")
          Result = F_printif_any(g0242.Arg)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          Core.C_pretty.Index = (Core.C_pretty.Index-3)
          Result = g0242.Printelse()
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
      var x_try02443 EID 
      x_try02443 = EVAL(self.Test)
      /* ERROR PROTECTION INSERTED (x-Result) */
      if ErrorIn(x_try02443) {Result = x_try02443
      } else {
      x = ANY(x_try02443)
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
              var g0245 *If   = To_If(x)
              /* noccur = 1 */
              void_try5 = g0245.Printstat()
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
        var g0247 *Do   = To_Do(x)
        /* noccur = 1 */
        Result = F_printdo_list(g0247.Args,CFALSE)
        /* Let-3 */} 
      /* If!2 */}  else if (x.Isa.IsIn(C_If) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0248 *If   = To_If(x)
        /* noccur = 1 */
        Result = g0248.Printstat()
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
          var res_try02505 EID 
          res_try02505 = EVAL(_Zx)
          /* ERROR PROTECTION INSERTED (res-Result) */
          if ErrorIn(res_try02505) {Result = res_try02505
          Result = res_try02505
          break
          } else {
          res = ANY(res_try02505)
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
          var g0252 *Let   = To_Let(a)
          /* noccur = 3 */
          PRINC(",")
          Result = F_lbreak_integer(4)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          Result = F_ppvariable_Variable(g0252.ClaireVar)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(" := ")
          Result = F_printexp_any(g0252.Value,CFALSE)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          Core.C_pretty.Index = (Core.C_pretty.Index-4)
          Result = g0252.Printbody()
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
      var val_try02543 EID 
      val_try02543 = EVAL(self.Value)
      /* ERROR PROTECTION INSERTED (val-Result) */
      if ErrorIn(val_try02543) {Result = val_try02543
      } else {
      val = ANY(val_try02543)
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
      var val_try02553 EID 
      val_try02553 = EVAL(self.Value)
      /* ERROR PROTECTION INSERTED (val-Result) */
      if ErrorIn(val_try02553) {Result = val_try02553
      } else {
      val = ANY(val_try02553)
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
              var g0257I *ClaireBoolean  
              if (lnext.Isa.IsIn(C_Let) == CTRUE) /* If:7 */{ 
                /* Let:8 */{ 
                  var g0256 *Let   = To_Let(lnext)
                  /* noccur = 2 */
                  g0257I = MakeBoolean((g0256.Value.Isa.IsIn(C_Call) == CTRUE) && (ToList(OBJ(Core.F_CALL(C_args,ARGS(g0256.Value.ToEID())))).At(1-1) == self.ClaireVar.Id()))
                  /* Let-8 */} 
                } else {
                g0257I = CFALSE
                /* If-7 */} 
              if (g0257I == CTRUE) /* If:7 */{ 
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
      var x_try02693 EID 
      x_try02693 = EVAL(self.SetArg)
      /* ERROR PROTECTION INSERTED (x-Result) */
      if ErrorIn(x_try02693) {Result = x_try02693
      } else {
      x = ANY(x_try02693)
      h_index := ClEnv.Index /* Handle */
      h_base := ClEnv.Base
      if (C_class.Id() == x.Isa.Id()) /* If:3 */{ 
        /* Let:4 */{ 
          var g0260 *ClaireClass   = ToClass(x)
          /* noccur = 1 */
          /* For:5 */{ 
            var y *ClaireAny  
            _ = y
            Result= EID{CFALSE.Id(),0}
            for _,y = range(g0260.Descendents.Values)/* loop:6 */{ 
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
          var g0261 *ClaireList   = ToList(x)
          /* noccur = 1 */
          /* For:5 */{ 
            var z *ClaireAny  
            _ = z
            Result= EID{CFALSE.Id(),0}
            var z_support *ClaireList  
            z_support = g0261
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
          var g0262 *ClaireList   = ToArray(x)
          /* noccur = 2 */
          /* Let:5 */{ 
            var n int  = g0262.Length()
            /* noccur = 1 */
            /* Let:6 */{ 
              var g0263 int  = 1
              /* noccur = 4 */
              /* Let:7 */{ 
                var g0264 int  = n
                /* noccur = 1 */
                Result= EID{CFALSE.Id(),0}
                for (g0263 <= g0264) /* while:8 */{ 
                  var void_try9 EID 
                  _ = void_try9
                  { 
                  /* Let:9 */{ 
                    var z *ClaireAny   = ToList(g0262.Id()).At(g0263-1)
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
                  g0263 = (g0263+1)
                  }
                  /* while-8 */} 
                }
                /* Let-7 */} 
              /* Let-6 */} 
            /* Let-5 */} 
          /* Let-4 */} 
        /* If!3 */}  else if (x.Isa.IsIn(C_Interval) == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var g0265 *ClaireInterval   = To_Interval(x)
          /* noccur = 2 */
          /* Let:5 */{ 
            var y int  = g0265.Arg1
            /* noccur = 4 */
            /* Let:6 */{ 
              var g0266 int  = g0265.Arg2
              /* noccur = 1 */
              Result= EID{CFALSE.Id(),0}
              for (y <= g0266) /* while:7 */{ 
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
          var g0267 *ClaireCollection   = ToCollection(x)
          /* noccur = 1 */
          /* For:5 */{ 
            var y *ClaireAny  
            _ = y
            Result= EID{CFALSE.Id(),0}
            var y_support *ClaireList  
            var y_support_try02706 EID 
            y_support_try02706 = Core.F_enumerate_any(g0267.Id())
            /* ERROR PROTECTION INSERTED (y_support-Result) */
            if ErrorIn(y_support_try02706) {Result = y_support_try02706
            } else {
            y_support = ToList(OBJ(y_support_try02706))
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
      var x_try02773 EID 
      x_try02773 = EVAL(self.SetArg)
      /* ERROR PROTECTION INSERTED (x-Result) */
      if ErrorIn(x_try02773) {Result = x_try02773
      } else {
      x = ANY(x_try02773)
      /* Let:3 */{ 
        var res *ClaireList  
        /* noccur = 9 */
        /* Let:4 */{ 
          var g0278UU *ClaireType  
          /* noccur = 1 */
          if (self.Of.Id() != CNULL) /* If:5 */{ 
            g0278UU = self.Of
            } else {
            g0278UU = ToType(CEMPTY.Id())
            /* If-5 */} 
          res = g0278UU.EmptyList()
          /* Let-4 */} 
        if (C_class.Id() == x.Isa.Id()) /* If:4 */{ 
          /* Let:5 */{ 
            var g0272 *ClaireClass   = ToClass(x)
            /* noccur = 1 */
            /* For:6 */{ 
              var y *ClaireAny  
              _ = y
              Result= EID{CFALSE.Id(),0}
              for _,y = range(g0272.Descendents.Values)/* loop:7 */{ 
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
                    var res_try027910 EID 
                    /* Let:10 */{ 
                      var g0280UU *ClaireAny  
                      /* noccur = 1 */
                      var g0280UU_try028111 EID 
                      g0280UU_try028111 = EVAL(self.Arg)
                      /* ERROR PROTECTION INSERTED (g0280UU-res_try027910) */
                      if ErrorIn(g0280UU_try028111) {res_try027910 = g0280UU_try028111
                      } else {
                      g0280UU = ANY(g0280UU_try028111)
                      res_try027910 = EID{res.AddFast(g0280UU).Id(),0}
                      }
                      /* Let-10 */} 
                    /* ERROR PROTECTION INSERTED (res-void_try10) */
                    if ErrorIn(res_try027910) {void_try10 = res_try027910
                    void_try8 = res_try027910
                    break
                    } else {
                    res = ToList(OBJ(res_try027910))
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
            var g0273 *ClaireList   = ToList(x)
            /* noccur = 1 */
            /* For:6 */{ 
              var y *ClaireAny  
              _ = y
              Result= EID{CFALSE.Id(),0}
              var y_support *ClaireList  
              y_support = g0273
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
                var res_try02828 EID 
                /* Let:8 */{ 
                  var g0283UU *ClaireAny  
                  /* noccur = 1 */
                  var g0283UU_try02849 EID 
                  g0283UU_try02849 = EVAL(self.Arg)
                  /* ERROR PROTECTION INSERTED (g0283UU-res_try02828) */
                  if ErrorIn(g0283UU_try02849) {res_try02828 = g0283UU_try02849
                  } else {
                  g0283UU = ANY(g0283UU_try02849)
                  res_try02828 = EID{res.AddFast(g0283UU).Id(),0}
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (res-void_try8) */
                if ErrorIn(res_try02828) {void_try8 = res_try02828
                Result = res_try02828
                break
                } else {
                res = ToList(OBJ(res_try02828))
                void_try8 = EID{res.Id(),0}
                }}
                }
                /* loop-7 */} 
              /* For-6 */} 
            /* Let-5 */} 
          /* If!4 */}  else if (x.Isa.IsIn(C_Interval) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var g0274 *ClaireInterval   = To_Interval(x)
            /* noccur = 2 */
            /* Let:6 */{ 
              var y int  = g0274.Arg1
              /* noccur = 4 */
              /* Let:7 */{ 
                var g0275 int  = g0274.Arg2
                /* noccur = 1 */
                Result= EID{CFALSE.Id(),0}
                for (y <= g0275) /* while:8 */{ 
                  var void_try9 EID 
                  _ = void_try9
                  { 
                  void_try9 = self.ClaireVar.WriteEID(EID{C__INT,IVAL(y)})
                  /* ERROR PROTECTION INSERTED (void_try9-void_try9) */
                  if ErrorIn(void_try9) {Result = void_try9
                  break
                  } else {
                  var res_try02859 EID 
                  /* Let:9 */{ 
                    var g0286UU *ClaireAny  
                    /* noccur = 1 */
                    var g0286UU_try028710 EID 
                    g0286UU_try028710 = EVAL(self.Arg)
                    /* ERROR PROTECTION INSERTED (g0286UU-res_try02859) */
                    if ErrorIn(g0286UU_try028710) {res_try02859 = g0286UU_try028710
                    } else {
                    g0286UU = ANY(g0286UU_try028710)
                    res_try02859 = EID{res.AddFast(g0286UU).Id(),0}
                    }
                    /* Let-9 */} 
                  /* ERROR PROTECTION INSERTED (res-void_try9) */
                  if ErrorIn(res_try02859) {void_try9 = res_try02859
                  Result = res_try02859
                  break
                  } else {
                  res = ToList(OBJ(res_try02859))
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
            var y_support_try02886 EID 
            y_support_try02886 = Core.F_enumerate_any(x)
            /* ERROR PROTECTION INSERTED (y_support-Result) */
            if ErrorIn(y_support_try02886) {Result = y_support_try02886
            } else {
            y_support = ToList(OBJ(y_support_try02886))
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
              var res_try02897 EID 
              /* Let:7 */{ 
                var g0290UU *ClaireAny  
                /* noccur = 1 */
                var g0290UU_try02918 EID 
                g0290UU_try02918 = EVAL(self.Arg)
                /* ERROR PROTECTION INSERTED (g0290UU-res_try02897) */
                if ErrorIn(g0290UU_try02918) {res_try02897 = g0290UU_try02918
                } else {
                g0290UU = ANY(g0290UU_try02918)
                res_try02897 = EID{res.AddFast(g0290UU).Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (res-void_try7) */
              if ErrorIn(res_try02897) {void_try7 = res_try02897
              Result = res_try02897
              break
              } else {
              res = ToList(OBJ(res_try02897))
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
      var x_try02933 EID 
      x_try02933 = EVAL(self.SetArg)
      /* ERROR PROTECTION INSERTED (x-Result) */
      if ErrorIn(x_try02933) {Result = x_try02933
      } else {
      x = ANY(x_try02933)
      /* Let:3 */{ 
        var res *ClaireSet  
        /* noccur = 3 */
        /* Let:4 */{ 
          var g0294UU *ClaireType  
          /* noccur = 1 */
          if (self.Of.Id() != CNULL) /* If:5 */{ 
            g0294UU = self.Of
            } else {
            g0294UU = ToType(CEMPTY.Id())
            /* If-5 */} 
          res = g0294UU.EmptySet()
          /* Let-4 */} 
        /* For:4 */{ 
          var y *ClaireAny  
          _ = y
          Result= EID{CFALSE.Id(),0}
          var y_support *ClaireList  
          var y_support_try02955 EID 
          y_support_try02955 = Core.F_enumerate_any(x)
          /* ERROR PROTECTION INSERTED (y_support-Result) */
          if ErrorIn(y_support_try02955) {Result = y_support_try02955
          } else {
          y_support = ToList(OBJ(y_support_try02955))
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
            var res_try02966 EID 
            /* Let:6 */{ 
              var g0297UU *ClaireAny  
              /* noccur = 1 */
              var g0297UU_try02987 EID 
              g0297UU_try02987 = EVAL(self.Arg)
              /* ERROR PROTECTION INSERTED (g0297UU-res_try02966) */
              if ErrorIn(g0297UU_try02987) {res_try02966 = g0297UU_try02987
              } else {
              g0297UU = ANY(g0297UU_try02987)
              res_try02966 = EID{res.AddFast(g0297UU).Id(),0}
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (res-void_try6) */
            if ErrorIn(res_try02966) {void_try6 = res_try02966
            Result = res_try02966
            break
            } else {
            res = ToSet(OBJ(res_try02966))
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
      var x_try03043 EID 
      x_try03043 = EVAL(self.SetArg)
      /* ERROR PROTECTION INSERTED (x-Result) */
      if ErrorIn(x_try03043) {Result = x_try03043
      } else {
      x = ANY(x_try03043)
      /* Let:3 */{ 
        var res *ClaireSet  
        /* noccur = 7 */
        /* Let:4 */{ 
          var g0305UU *ClaireType  
          /* noccur = 1 */
          if (self.Of.Id() != CNULL) /* If:5 */{ 
            g0305UU = self.Of
            } else {
            g0305UU = ToType(CEMPTY.Id())
            /* If-5 */} 
          res = g0305UU.EmptySet()
          /* Let-4 */} 
        if (C_class.Id() == x.Isa.Id()) /* If:4 */{ 
          /* Let:5 */{ 
            var g0300 *ClaireClass   = ToClass(x)
            /* noccur = 1 */
            /* For:6 */{ 
              var y *ClaireAny  
              _ = y
              Result= EID{CFALSE.Id(),0}
              for _,y = range(g0300.Descendents.Values)/* loop:7 */{ 
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
                    var g0306I *ClaireBoolean  
                    var g0306I_try030710 EID 
                    /* Let:10 */{ 
                      var g0308UU *ClaireAny  
                      /* noccur = 1 */
                      var g0308UU_try030911 EID 
                      g0308UU_try030911 = EVAL(self.Arg)
                      /* ERROR PROTECTION INSERTED (g0308UU-g0306I_try030710) */
                      if ErrorIn(g0308UU_try030911) {g0306I_try030710 = g0308UU_try030911
                      } else {
                      g0308UU = ANY(g0308UU_try030911)
                      g0306I_try030710 = EID{Core.F__I_equal_any(g0308UU,CFALSE.Id()).Id(),0}
                      }
                      /* Let-10 */} 
                    /* ERROR PROTECTION INSERTED (g0306I-void_try10) */
                    if ErrorIn(g0306I_try030710) {void_try10 = g0306I_try030710
                    } else {
                    g0306I = ToBoolean(OBJ(g0306I_try030710))
                    if (g0306I == CTRUE) /* If:10 */{ 
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
            var g0301 *ClaireInterval   = To_Interval(x)
            /* noccur = 2 */
            /* Let:6 */{ 
              var y int  = g0301.Arg1
              /* noccur = 5 */
              /* Let:7 */{ 
                var g0302 int  = g0301.Arg2
                /* noccur = 1 */
                Result= EID{CFALSE.Id(),0}
                for (y <= g0302) /* while:8 */{ 
                  var void_try9 EID 
                  _ = void_try9
                  { 
                  void_try9 = self.ClaireVar.WriteEID(EID{C__INT,IVAL(y)})
                  /* ERROR PROTECTION INSERTED (void_try9-void_try9) */
                  if ErrorIn(void_try9) {Result = void_try9
                  break
                  } else {
                  var g0310I *ClaireBoolean  
                  var g0310I_try03119 EID 
                  /* Let:9 */{ 
                    var g0312UU *ClaireAny  
                    /* noccur = 1 */
                    var g0312UU_try031310 EID 
                    g0312UU_try031310 = EVAL(self.Arg)
                    /* ERROR PROTECTION INSERTED (g0312UU-g0310I_try03119) */
                    if ErrorIn(g0312UU_try031310) {g0310I_try03119 = g0312UU_try031310
                    } else {
                    g0312UU = ANY(g0312UU_try031310)
                    g0310I_try03119 = EID{Core.F__I_equal_any(g0312UU,CFALSE.Id()).Id(),0}
                    }
                    /* Let-9 */} 
                  /* ERROR PROTECTION INSERTED (g0310I-void_try9) */
                  if ErrorIn(g0310I_try03119) {void_try9 = g0310I_try03119
                  } else {
                  g0310I = ToBoolean(OBJ(g0310I_try03119))
                  if (g0310I == CTRUE) /* If:9 */{ 
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
            var y_support_try03146 EID 
            y_support_try03146 = Core.F_enumerate_any(x)
            /* ERROR PROTECTION INSERTED (y_support-Result) */
            if ErrorIn(y_support_try03146) {Result = y_support_try03146
            } else {
            y_support = ToList(OBJ(y_support_try03146))
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
              var g0315I *ClaireBoolean  
              var g0315I_try03167 EID 
              /* Let:7 */{ 
                var g0317UU *ClaireAny  
                /* noccur = 1 */
                var g0317UU_try03188 EID 
                g0317UU_try03188 = EVAL(self.Arg)
                /* ERROR PROTECTION INSERTED (g0317UU-g0315I_try03167) */
                if ErrorIn(g0317UU_try03188) {g0315I_try03167 = g0317UU_try03188
                } else {
                g0317UU = ANY(g0317UU_try03188)
                g0315I_try03167 = EID{Core.F__I_equal_any(g0317UU,CFALSE.Id()).Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (g0315I-void_try7) */
              if ErrorIn(g0315I_try03167) {void_try7 = g0315I_try03167
              } else {
              g0315I = ToBoolean(OBJ(g0315I_try03167))
              if (g0315I == CTRUE) /* If:7 */{ 
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
      var x_try03243 EID 
      x_try03243 = EVAL(self.SetArg)
      /* ERROR PROTECTION INSERTED (x-Result) */
      if ErrorIn(x_try03243) {Result = x_try03243
      } else {
      x = ANY(x_try03243)
      /* Let:3 */{ 
        var res *ClaireList  
        /* noccur = 7 */
        if (x.Isa.IsIn(C_list) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var g0320 *ClaireList   = ToList(x)
            /* noccur = 1 */
            res = g0320.Empty()
            /* Let-5 */} 
          } else {
          res = ToType(CEMPTY.Id()).EmptyList()
          /* If-4 */} 
        if (C_class.Id() == x.Isa.Id()) /* If:4 */{ 
          /* Let:5 */{ 
            var g0322 *ClaireClass   = ToClass(x)
            /* noccur = 1 */
            /* For:6 */{ 
              var y *ClaireAny  
              _ = y
              Result= EID{CFALSE.Id(),0}
              for _,y = range(g0322.Descendents.Values)/* loop:7 */{ 
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
                    var g0325I *ClaireBoolean  
                    var g0325I_try032610 EID 
                    /* Let:10 */{ 
                      var g0327UU *ClaireAny  
                      /* noccur = 1 */
                      var g0327UU_try032811 EID 
                      g0327UU_try032811 = EVAL(self.Arg)
                      /* ERROR PROTECTION INSERTED (g0327UU-g0325I_try032610) */
                      if ErrorIn(g0327UU_try032811) {g0325I_try032610 = g0327UU_try032811
                      } else {
                      g0327UU = ANY(g0327UU_try032811)
                      g0325I_try032610 = EID{Core.F__I_equal_any(g0327UU,CFALSE.Id()).Id(),0}
                      }
                      /* Let-10 */} 
                    /* ERROR PROTECTION INSERTED (g0325I-void_try10) */
                    if ErrorIn(g0325I_try032610) {void_try10 = g0325I_try032610
                    } else {
                    g0325I = ToBoolean(OBJ(g0325I_try032610))
                    if (g0325I == CTRUE) /* If:10 */{ 
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
            var y_support_try03296 EID 
            y_support_try03296 = Core.F_enumerate_any(x)
            /* ERROR PROTECTION INSERTED (y_support-Result) */
            if ErrorIn(y_support_try03296) {Result = y_support_try03296
            } else {
            y_support = ToList(OBJ(y_support_try03296))
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
              var g0330I *ClaireBoolean  
              var g0330I_try03317 EID 
              /* Let:7 */{ 
                var g0332UU *ClaireAny  
                /* noccur = 1 */
                var g0332UU_try03338 EID 
                g0332UU_try03338 = EVAL(self.Arg)
                /* ERROR PROTECTION INSERTED (g0332UU-g0330I_try03317) */
                if ErrorIn(g0332UU_try03338) {g0330I_try03317 = g0332UU_try03338
                } else {
                g0332UU = ANY(g0332UU_try03338)
                g0330I_try03317 = EID{Core.F__I_equal_any(g0332UU,CFALSE.Id()).Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (g0330I-void_try7) */
              if ErrorIn(g0330I_try03317) {void_try7 = g0330I_try03317
              } else {
              g0330I = ToBoolean(OBJ(g0330I_try03317))
              if (g0330I == CTRUE) /* If:7 */{ 
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
      var x_try03373 EID 
      x_try03373 = EVAL(self.SetArg)
      /* ERROR PROTECTION INSERTED (x-Result) */
      if ErrorIn(x_try03373) {Result = x_try03373
      } else {
      x = ANY(x_try03373)
      /* Let:3 */{ 
        var b *ClaireAny   = self.Other
        /* noccur = 7 */
        /* Let:4 */{ 
          var res *ClaireAny   = b
          /* noccur = 5 */
          if (C_class.Id() == x.Isa.Id()) /* If:5 */{ 
            /* Let:6 */{ 
              var g0335 *ClaireClass   = ToClass(x)
              /* noccur = 1 */
              /* For:7 */{ 
                var y *ClaireAny  
                _ = y
                Result= EID{CFALSE.Id(),0}
                for _,y = range(g0335.Descendents.Values)/* loop:8 */{ 
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
                      var g0338I *ClaireBoolean  
                      var g0338I_try033911 EID 
                      /* Let:11 */{ 
                        var g0340UU *ClaireAny  
                        /* noccur = 1 */
                        var g0340UU_try034112 EID 
                        g0340UU_try034112 = EVAL(self.Arg)
                        /* ERROR PROTECTION INSERTED (g0340UU-g0338I_try033911) */
                        if ErrorIn(g0340UU_try034112) {g0338I_try033911 = g0340UU_try034112
                        } else {
                        g0340UU = ANY(g0340UU_try034112)
                        g0338I_try033911 = EID{Core.F__I_equal_any(g0340UU,CFALSE.Id()).Id(),0}
                        }
                        /* Let-11 */} 
                      /* ERROR PROTECTION INSERTED (g0338I-void_try11) */
                      if ErrorIn(g0338I_try033911) {void_try11 = g0338I_try033911
                      } else {
                      g0338I = ToBoolean(OBJ(g0338I_try033911))
                      if (g0338I == CTRUE) /* If:11 */{ 
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
              var y_support_try03427 EID 
              y_support_try03427 = Core.F_enumerate_any(x)
              /* ERROR PROTECTION INSERTED (y_support-Result) */
              if ErrorIn(y_support_try03427) {Result = y_support_try03427
              } else {
              y_support = ToList(OBJ(y_support_try03427))
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
                var g0343I *ClaireBoolean  
                var g0343I_try03448 EID 
                /* Let:8 */{ 
                  var g0345UU *ClaireAny  
                  /* noccur = 1 */
                  var g0345UU_try03469 EID 
                  g0345UU_try03469 = EVAL(self.Arg)
                  /* ERROR PROTECTION INSERTED (g0345UU-g0343I_try03448) */
                  if ErrorIn(g0345UU_try03469) {g0343I_try03448 = g0345UU_try03469
                  } else {
                  g0345UU = ANY(g0345UU_try03469)
                  g0343I_try03448 = EID{Core.F__I_equal_any(g0345UU,CFALSE.Id()).Id(),0}
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (g0343I-void_try8) */
                if ErrorIn(g0343I_try03448) {void_try8 = g0343I_try03448
                } else {
                g0343I = ToBoolean(OBJ(g0343I_try03448))
                if (g0343I == CTRUE) /* If:8 */{ 
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
      var truc_try03493 EID 
      truc_try03493 = EVAL(self.ClaireVar)
      /* ERROR PROTECTION INSERTED (truc-Result) */
      if ErrorIn(truc_try03493) {Result = truc_try03493
      } else {
      truc = ANY(truc_try03493)
      /* Let:3 */{ 
        var flip *ClaireBoolean   = CTRUE
        /* noccur = 3 */
        /* Let:4 */{ 
          var previous *ClaireAny   = CFALSE.Id()
          /* noccur = 4 */
          var g0350I *ClaireBoolean  
          var g0350I_try03515 EID 
          /* For:5 */{ 
            var x *ClaireAny  
            _ = x
            g0350I_try03515= EID{CFALSE.Id(),0}
            var x_support *ClaireList  
            x_support = self.Args
            x_len := x_support.Length()
            for i_it := 0; i_it < x_len; i_it++ { 
              x = x_support.At(i_it)
              var void_try7 EID 
              _ = void_try7
              if (flip == CTRUE) /* If:7 */{ 
                flip = CFALSE
                var previous_try03528 EID 
                previous_try03528 = EVAL(x)
                /* ERROR PROTECTION INSERTED (previous-void_try7) */
                if ErrorIn(previous_try03528) {void_try7 = previous_try03528
                g0350I_try03515 = previous_try03528
                break
                } else {
                previous = ANY(previous_try03528)
                void_try7 = previous.ToEID()
                }
                } else {
                var g0353I *ClaireBoolean  
                var g0353I_try03548 EID 
                g0353I_try03548 = Core.F_BELONG(truc,previous)
                /* ERROR PROTECTION INSERTED (g0353I-void_try7) */
                if ErrorIn(g0353I_try03548) {void_try7 = g0353I_try03548
                } else {
                g0353I = ToBoolean(OBJ(g0353I_try03548))
                if (g0353I == CTRUE) /* If:8 */{ 
                  var previous_try03559 EID 
                  previous_try03559 = EVAL(x)
                  /* ERROR PROTECTION INSERTED (previous-void_try7) */
                  if ErrorIn(previous_try03559) {void_try7 = previous_try03559
                  g0350I_try03515 = previous_try03559
                  break
                  } else {
                  previous = ANY(previous_try03559)
                  void_try7 = previous.ToEID()
                   /*v = g0350I_try03515, s =EID*/
g0350I_try03515 = EID{CTRUE.Id(),0}
                  break
                  }
                  } else {
                  flip = CTRUE
                  void_try7 = EID{flip.Id(),0}
                  /* If-8 */} 
                }
                /* If-7 */} 
              /* ERROR PROTECTION INSERTED (void_try7-g0350I_try03515) */
              if ErrorIn(void_try7) {g0350I_try03515 = void_try7
              g0350I_try03515 = void_try7
              break
              } else {
              }
              /* loop-6 */} 
            /* For-5 */} 
          /* ERROR PROTECTION INSERTED (g0350I-Result) */
          if ErrorIn(g0350I_try03515) {Result = g0350I_try03515
          } else {
          g0350I = ToBoolean(OBJ(g0350I_try03515))
          if (g0350I == CTRUE) /* If:5 */{ 
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
        var v_while4_try03574 EID 
        /* or:4 */{ 
          var v_or4 *ClaireBoolean  
          
          v_or4 = b
          if (v_or4 == CTRUE) {v_while4_try03574 = EID{CTRUE.Id(),0}
          } else /* or:5 */{ 
            var v_or4_try03586 EID 
            /* Let:6 */{ 
              var g0359UU *ClaireBoolean  
              /* noccur = 1 */
              var g0359UU_try03607 EID 
              /* Let:7 */{ 
                var g0361UU *ClaireAny  
                /* noccur = 1 */
                var g0361UU_try03628 EID 
                g0361UU_try03628 = EVAL(self.Test)
                /* ERROR PROTECTION INSERTED (g0361UU-g0359UU_try03607) */
                if ErrorIn(g0361UU_try03628) {g0359UU_try03607 = g0361UU_try03628
                } else {
                g0361UU = ANY(g0361UU_try03628)
                g0359UU_try03607 = EID{Core.F_not_any(g0361UU).Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (g0359UU-v_or4_try03586) */
              if ErrorIn(g0359UU_try03607) {v_or4_try03586 = g0359UU_try03607
              } else {
              g0359UU = ToBoolean(OBJ(g0359UU_try03607))
              v_or4_try03586 = EID{Equal(g0359UU.Id(),a.Id()).Id(),0}
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (v_or4-v_while4_try03574) */
            if ErrorIn(v_or4_try03586) {v_while4_try03574 = v_or4_try03586
            } else {
            v_or4 = ToBoolean(OBJ(v_or4_try03586))
            if (v_or4 == CTRUE) {v_while4_try03574 = EID{CTRUE.Id(),0}
            } else /* or:6 */{ 
              v_while4_try03574 = EID{CFALSE.Id(),0}/* org-6 */} 
            /* org-5 */} 
          }
          /* or-4 */} 
        /* ERROR PROTECTION INSERTED (v_while4-Result) */
        if ErrorIn(v_while4_try03574) {Result = v_while4_try03574
        } else {
        v_while4 = ToBoolean(OBJ(v_while4_try03574))
        
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
          var v_while4_try03635 EID 
          /* or:5 */{ 
            var v_or5 *ClaireBoolean  
            
            v_or5 = b
            if (v_or5 == CTRUE) {v_while4_try03635 = EID{CTRUE.Id(),0}
            } else /* or:6 */{ 
              var v_or5_try03647 EID 
              /* Let:7 */{ 
                var g0365UU *ClaireBoolean  
                /* noccur = 1 */
                var g0365UU_try03668 EID 
                /* Let:8 */{ 
                  var g0367UU *ClaireAny  
                  /* noccur = 1 */
                  var g0367UU_try03689 EID 
                  g0367UU_try03689 = EVAL(self.Test)
                  /* ERROR PROTECTION INSERTED (g0367UU-g0365UU_try03668) */
                  if ErrorIn(g0367UU_try03689) {g0365UU_try03668 = g0367UU_try03689
                  } else {
                  g0367UU = ANY(g0367UU_try03689)
                  g0365UU_try03668 = EID{Core.F_not_any(g0367UU).Id(),0}
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (g0365UU-v_or5_try03647) */
                if ErrorIn(g0365UU_try03668) {v_or5_try03647 = g0365UU_try03668
                } else {
                g0365UU = ToBoolean(OBJ(g0365UU_try03668))
                v_or5_try03647 = EID{Equal(g0365UU.Id(),a.Id()).Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (v_or5-v_while4_try03635) */
              if ErrorIn(v_or5_try03647) {v_while4_try03635 = v_or5_try03647
              } else {
              v_or5 = ToBoolean(OBJ(v_or5_try03647))
              if (v_or5 == CTRUE) {v_while4_try03635 = EID{CTRUE.Id(),0}
              } else /* or:7 */{ 
                v_while4_try03635 = EID{CFALSE.Id(),0}/* org-7 */} 
              /* org-6 */} 
            }
            /* or-5 */} 
          /* ERROR PROTECTION INSERTED (v_while4-Result) */
          if ErrorIn(v_while4_try03635) {Result = v_while4_try03635
          Result = v_while4_try03635
          break
          } else {
          v_while4 = ToBoolean(OBJ(v_while4_try03635))
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
        var g0380UU *ClaireString  
        /* noccur = 1 */
        if (self.Isa.IsIn(C_List) == CTRUE) /* If:4 */{ 
          g0380UU = MakeString("list")
          /* If!4 */}  else if (self.Isa.IsIn(C_Set) == CTRUE) /* If:4 */{ 
          g0380UU = MakeString("set")
          /* If!4 */}  else if (self.Isa.IsIn(C_Tuple) == CTRUE) /* If:4 */{ 
          g0380UU = MakeString("tuple")
          /* If!4 */}  else if (self.Isa.IsIn(C_Printf) == CTRUE) /* If:4 */{ 
          g0380UU = MakeString("printf")
          /* If!4 */}  else if (self.Isa.IsIn(C_Error) == CTRUE) /* If:4 */{ 
          g0380UU = MakeString("error")
          /* If!4 */}  else if (self.Isa.IsIn(C_Trace) == CTRUE) /* If:4 */{ 
          g0380UU = MakeString("trace")
          /* If!4 */}  else if (self.Isa.IsIn(C_Assert) == CTRUE) /* If:4 */{ 
          g0380UU = MakeString("assert")
          /* If!4 */}  else if (self.Isa.IsIn(C_Branch) == CTRUE) /* If:4 */{ 
          g0380UU = MakeString("branch")
          } else {
          g0380UU = self.Isa.Name.String_I()
          /* If-4 */} 
        F_princ_string(g0380UU)
        /* Let-3 */} 
      if ((self.Isa.IsIn(C_List) == CTRUE) || 
          (self.Isa.IsIn(C_Set) == CTRUE)) /* If:3 */{ 
        /* Let:4 */{ 
          var g0379 *Construct   = self
          /* noccur = 1 */
          /* Let:5 */{ 
            var _Zt *ClaireAny   = Core.F_get_property(C_of,ToObject(g0379.Id()))
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
                var g0381 int  = n
                /* noccur = 1 */
                Result= EID{CFALSE.Id(),0}
                for (i <= g0381) /* while:8 */{ 
                  var void_try9 EID 
                  _ = void_try9
                  { 
                  /* LetEID:9 */{ 
                    var g0383UU EID 
                    g0383UU = EVAL(self.Args.At(i-1))
                    /* ERROR PROTECTION INSERTED (g0383UU-void_try9) */
                    if ErrorIn(g0383UU) {void_try9 = g0383UU
                    } else {
                    void_try9 = l.WriteEID(i,g0383UU)}
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
                var g0382 int  = n
                /* noccur = 1 */
                Result= EID{CFALSE.Id(),0}
                for (i <= g0382) /* while:8 */{ 
                  var void_try9 EID 
                  _ = void_try9
                  { 
                  /* Let:9 */{ 
                    var g0384UU *ClaireAny  
                    /* noccur = 1 */
                    var g0384UU_try038510 EID 
                    g0384UU_try038510 = EVAL(self.Args.At(i-1))
                    /* ERROR PROTECTION INSERTED (g0384UU-void_try9) */
                    if ErrorIn(g0384UU_try038510) {void_try9 = g0384UU_try038510
                    } else {
                    g0384UU = ANY(g0384UU_try038510)
                    void_try9 = ToArray(l.Id()).NthPut(i,g0384UU).ToEID()
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
      var s_try03863 EID 
      /* Let:3 */{ 
        var x_bag *ClaireSet   = ToType(CEMPTY.Id()).EmptySet()
        /* noccur = 2 */
        /* For:4 */{ 
          var x *ClaireAny  
          _ = x
          s_try03863= EID{CFALSE.Id(),0}
          var x_support *ClaireList  
          x_support = self.Args
          x_len := x_support.Length()
          for i_it := 0; i_it < x_len; i_it++ { 
            x = x_support.At(i_it)
            var void_try6 EID 
            _ = void_try6
            /* Let:6 */{ 
              var g0387UU *ClaireAny  
              /* noccur = 1 */
              var g0387UU_try03887 EID 
              g0387UU_try03887 = EVAL(x)
              /* ERROR PROTECTION INSERTED (g0387UU-void_try6) */
              if ErrorIn(g0387UU_try03887) {void_try6 = g0387UU_try03887
              } else {
              g0387UU = ANY(g0387UU_try03887)
              void_try6 = EID{x_bag.AddFast(g0387UU).Id(),0}
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (void_try6-s_try03863) */
            if ErrorIn(void_try6) {s_try03863 = void_try6
            s_try03863 = void_try6
            break
            } else {
            }
            /* loop-5 */} 
          /* For-4 */} 
        /* ERROR PROTECTION INSERTED (s_try03863-s_try03863) */
        if !ErrorIn(s_try03863) {
        s_try03863 = EID{x_bag.Id(),0}
        }
        /* Let-3 */} 
      /* ERROR PROTECTION INSERTED (s-Result) */
      if ErrorIn(s_try03863) {Result = s_try03863
      } else {
      s = ToSet(OBJ(s_try03863))
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
      var g0389UU *ClaireList  
      /* noccur = 1 */
      var g0389UU_try03903 EID 
      /* Iteration:3 */{ 
        var v_list3 *ClaireList  
        var x *ClaireAny  
        var v_local3 *ClaireAny  
        v_list3 = self.Args
        g0389UU_try03903 = EID{CreateList(ToType(CEMPTY.Id()),v_list3.Length()).Id(),0}
        for CLcount := 0; CLcount < v_list3.Length(); CLcount++{ 
          x = v_list3.At(CLcount)
          var v_local3_try03915 EID 
          v_local3_try03915 = EVAL(x)
          /* ERROR PROTECTION INSERTED (v_local3-g0389UU_try03903) */
          if ErrorIn(v_local3_try03915) {g0389UU_try03903 = v_local3_try03915
          g0389UU_try03903 = v_local3_try03915
          break
          } else {
          v_local3 = ANY(v_local3_try03915)
          ToList(OBJ(g0389UU_try03903)).PutAt(CLcount,v_local3)
          } 
        }
        /* Iteration-3 */} 
      /* ERROR PROTECTION INSERTED (g0389UU-Result) */
      if ErrorIn(g0389UU_try03903) {Result = g0389UU_try03903
      } else {
      g0389UU = ToList(OBJ(g0389UU_try03903))
      Result = EID{g0389UU.Tuple_I().Id(),0}
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
                var g0392 int  = n
                /* noccur = 1 */
                Result= EID{CFALSE.Id(),0}
                for (i <= g0392) /* while:8 */{ 
                  var void_try9 EID 
                  _ = void_try9
                  { 
                  /* LetEID:9 */{ 
                    var g0394UU EID 
                    g0394UU = EVAL(self.Args.At(i-1))
                    /* ERROR PROTECTION INSERTED (g0394UU-void_try9) */
                    if ErrorIn(g0394UU) {void_try9 = g0394UU
                    } else {
                    void_try9 = l.WriteEID(i,g0394UU)}
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
                var g0393 int  = n
                /* noccur = 1 */
                Result= EID{CFALSE.Id(),0}
                for (i <= g0393) /* while:8 */{ 
                  var void_try9 EID 
                  _ = void_try9
                  { 
                  /* Let:9 */{ 
                    var g0395UU *ClaireAny  
                    /* noccur = 1 */
                    var g0395UU_try039610 EID 
                    g0395UU_try039610 = EVAL(self.Args.At(i-1))
                    /* ERROR PROTECTION INSERTED (g0395UU-void_try9) */
                    if ErrorIn(g0395UU_try039610) {void_try9 = g0395UU_try039610
                    } else {
                    g0395UU = ANY(g0395UU_try039610)
                    void_try9 = ToArray(l.Id()).NthPut(i,g0395UU).ToEID()
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
      var g0397UU *ClaireAny  
      /* noccur = 1 */
      var g0397UU_try03983 EID 
      g0397UU_try03983 = Core.F_CALL(C_macroexpand,ARGS(EID{self.Id(),0}))
      /* ERROR PROTECTION INSERTED (g0397UU-Result) */
      if ErrorIn(g0397UU_try03983) {Result = g0397UU_try03983
      } else {
      g0397UU = ANY(g0397UU_try03983)
      Result = EVAL(g0397UU)
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
        var va_arg2_try03994 EID 
        va_arg2_try03994 = Core.F_car_list(self.Args)
        /* ERROR PROTECTION INSERTED (va_arg2-Result) */
        if ErrorIn(va_arg2_try03994) {Result = va_arg2_try03994
        } else {
        va_arg2 = ANY(va_arg2_try03994)
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
        var va_arg2_try04004 EID 
        /* Iteration:4 */{ 
          var v_list4 *ClaireList  
          var x *ClaireAny  
          var v_local4 *ClaireAny  
          var v_list4_try04015 EID 
          v_list4_try04015 = self.Args.Cdr()
          /* ERROR PROTECTION INSERTED (v_list4-va_arg2_try04004) */
          if ErrorIn(v_list4_try04015) {va_arg2_try04004 = v_list4_try04015
          } else {
          v_list4 = ToList(OBJ(v_list4_try04015))
          va_arg2_try04004 = EID{CreateList(ToType(CEMPTY.Id()),v_list4.Length()).Id(),0}
          for CLcount := 0; CLcount < v_list4.Length(); CLcount++{ 
            x = v_list4.At(CLcount)
            var v_local4_try04026 EID 
            v_local4_try04026 = EVAL(x)
            /* ERROR PROTECTION INSERTED (v_local4-va_arg2_try04004) */
            if ErrorIn(v_local4_try04026) {va_arg2_try04004 = v_local4_try04026
            va_arg2_try04004 = v_local4_try04026
            break
            } else {
            v_local4 = ANY(v_local4_try04026)
            ToList(OBJ(va_arg2_try04004)).PutAt(CLcount,v_local4)
            } 
          }}
          /* Iteration-4 */} 
        /* ERROR PROTECTION INSERTED (va_arg2-Result) */
        if ErrorIn(va_arg2_try04004) {Result = va_arg2_try04004
        } else {
        va_arg2 = ANY(va_arg2_try04004)
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
                  var m_try04039 EID 
                  m_try04039 = Core.F_CALL(C_nth,ARGS(s.ToEID(),EID{C__INT,IVAL((n+1))}))
                  /* ERROR PROTECTION INSERTED (m-void_try8) */
                  if ErrorIn(m_try04039) {void_try8 = m_try04039
                  } else {
                  m = ANY(m_try04039)
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
                      var g0404UU *ClaireAny  
                      /* noccur = 1 */
                      var g0404UU_try040511 EID 
                      g0404UU_try040511 = EVAL(l.At(i-1))
                      /* ERROR PROTECTION INSERTED (g0404UU-void_try8) */
                      if ErrorIn(g0404UU_try040511) {void_try8 = g0404UU_try040511
                      } else {
                      g0404UU = ANY(g0404UU_try040511)
                      void_try8 = Core.F_CALL(C_princ,ARGS(g0404UU.ToEID()))
                      }
                      /* Let-10 */} 
                    /* If!9 */}  else if ('S' == ToChar(m).Value) /* If:9 */{ 
                    /* Let:10 */{ 
                      var g0406UU *ClaireAny  
                      /* noccur = 1 */
                      var g0406UU_try040711 EID 
                      g0406UU_try040711 = EVAL(l.At(i-1))
                      /* ERROR PROTECTION INSERTED (g0406UU-void_try8) */
                      if ErrorIn(g0406UU_try040711) {void_try8 = g0406UU_try040711
                      } else {
                      g0406UU = ANY(g0406UU_try040711)
                      void_try8 = Core.F_CALL(C_print,ARGS(g0406UU.ToEID()))
                      }
                      /* Let-10 */} 
                    /* If!9 */}  else if ('F' == ToChar(m).Value) /* If:9 */{ 
                    /* Let:10 */{ 
                      var fv *ClaireAny  
                      /* noccur = 5 */
                      var fv_try040811 EID 
                      fv_try040811 = EVAL(l.At(i-1))
                      /* ERROR PROTECTION INSERTED (fv-void_try8) */
                      if ErrorIn(fv_try040811) {void_try8 = fv_try040811
                      } else {
                      fv = ANY(fv_try040811)
                      /* Let:11 */{ 
                        var p_Z *ClaireBoolean   = CFALSE
                        /* noccur = 4 */
                        /* Let:12 */{ 
                          var j int 
                          /* noccur = 4 */
                          var j_try040913 EID 
                          /* Let:13 */{ 
                            var g0410UU int 
                            /* noccur = 1 */
                            var g0410UU_try041114 EID 
                            /* Let:14 */{ 
                              var g0412UU rune 
                              /* noccur = 1 */
                              var g0412UU_try041315 EID 
                              g0412UU_try041315 = Core.F_nth_get_string(ToString(s),(n+2),(n+2))
                              /* ERROR PROTECTION INSERTED (g0412UU-g0410UU_try041114) */
                              if ErrorIn(g0412UU_try041315) {g0410UU_try041114 = g0412UU_try041315
                              } else {
                              g0412UU = CHAR(g0412UU_try041315)
                              g0410UU_try041114 = EID{C__INT,IVAL(int(g0412UU))}
                              }
                              /* Let-14 */} 
                            /* ERROR PROTECTION INSERTED (g0410UU-j_try040913) */
                            if ErrorIn(g0410UU_try041114) {j_try040913 = g0410UU_try041114
                            } else {
                            g0410UU = INT(g0410UU_try041114)
                            j_try040913 = EID{C__INT,IVAL((g0410UU-48))}
                            }
                            /* Let-13 */} 
                          /* ERROR PROTECTION INSERTED (j-void_try8) */
                          if ErrorIn(j_try040913) {void_try8 = j_try040913
                          } else {
                          j = INT(j_try040913)
                          var g0414I *ClaireBoolean  
                          var g0414I_try041513 EID 
                          /* Let:13 */{ 
                            var g0416UU *ClaireAny  
                            /* noccur = 1 */
                            var g0416UU_try041714 EID 
                            g0416UU_try041714 = Core.F_CALL(C_nth,ARGS(s.ToEID(),EID{C__INT,IVAL((n+2))}))
                            /* ERROR PROTECTION INSERTED (g0416UU-g0414I_try041513) */
                            if ErrorIn(g0416UU_try041714) {g0414I_try041513 = g0416UU_try041714
                            } else {
                            g0416UU = ANY(g0416UU_try041714)
                            g0414I_try041513 = EID{Equal(MakeChar('%').Id(),g0416UU).Id(),0}
                            }
                            /* Let-13 */} 
                          /* ERROR PROTECTION INSERTED (g0414I-void_try8) */
                          if ErrorIn(g0414I_try041513) {void_try8 = g0414I_try041513
                          } else {
                          g0414I = ToBoolean(OBJ(g0414I_try041513))
                          if (g0414I == CTRUE) /* If:13 */{ 
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
                          var g0418I *ClaireBoolean  
                          var g0418I_try041913 EID 
                          /* and:13 */{ 
                            var v_and13 *ClaireBoolean  
                            
                            v_and13 = p_Z.Not
                            if (v_and13 == CFALSE) {g0418I_try041913 = EID{CFALSE.Id(),0}
                            } else /* arg:14 */{ 
                              var v_and13_try042015 EID 
                              /* Let:15 */{ 
                                var g0421UU *ClaireAny  
                                /* noccur = 1 */
                                var g0421UU_try042216 EID 
                                g0421UU_try042216 = Core.F_CALL(C_nth,ARGS(s.ToEID(),EID{C__INT,IVAL((n+3))}))
                                /* ERROR PROTECTION INSERTED (g0421UU-v_and13_try042015) */
                                if ErrorIn(g0421UU_try042216) {v_and13_try042015 = g0421UU_try042216
                                } else {
                                g0421UU = ANY(g0421UU_try042216)
                                v_and13_try042015 = EID{Equal(MakeChar('%').Id(),g0421UU).Id(),0}
                                }
                                /* Let-15 */} 
                              /* ERROR PROTECTION INSERTED (v_and13-g0418I_try041913) */
                              if ErrorIn(v_and13_try042015) {g0418I_try041913 = v_and13_try042015
                              } else {
                              v_and13 = ToBoolean(OBJ(v_and13_try042015))
                              if (v_and13 == CFALSE) {g0418I_try041913 = EID{CFALSE.Id(),0}
                              } else /* arg:15 */{ 
                                g0418I_try041913 = EID{CTRUE.Id(),0}/* arg-15 */} 
                              /* arg-14 */} 
                            }
                            /* and-13 */} 
                          /* ERROR PROTECTION INSERTED (g0418I-void_try8) */
                          if ErrorIn(g0418I_try041913) {void_try8 = g0418I_try041913
                          } else {
                          g0418I = ToBoolean(OBJ(g0418I_try041913))
                          if (g0418I == CTRUE) /* If:13 */{ 
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
                  var n_try04239 EID 
                  n_try04239 = Core.F_CALL(C_get,ARGS(s.ToEID(),EID{C__CHAR,CVAL('~')}))
                  /* ERROR PROTECTION INSERTED (n-void_try8) */
                  if ErrorIn(n_try04239) {void_try8 = n_try04239
                  Result = n_try04239
                  break
                  } else {
                  n = INT(n_try04239)
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
        var l_try04254 EID 
        /* Iteration:4 */{ 
          var v_list4 *ClaireList  
          var x *ClaireAny  
          var v_local4 *ClaireAny  
          v_list4 = a
          l_try04254 = EID{CreateList(ToType(CEMPTY.Id()),v_list4.Length()).Id(),0}
          for CLcount := 0; CLcount < v_list4.Length(); CLcount++{ 
            x = v_list4.At(CLcount)
            var v_local4_try04266 EID 
            v_local4_try04266 = EVAL(x)
            /* ERROR PROTECTION INSERTED (v_local4-l_try04254) */
            if ErrorIn(v_local4_try04266) {l_try04254 = v_local4_try04266
            l_try04254 = v_local4_try04266
            break
            } else {
            v_local4 = ANY(v_local4_try04266)
            ToList(OBJ(l_try04254)).PutAt(CLcount,v_local4)
            } 
          }
          /* Iteration-4 */} 
        /* ERROR PROTECTION INSERTED (l-Result) */
        if ErrorIn(l_try04254) {Result = l_try04254
        } else {
        l = ToList(OBJ(l_try04254))
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
                var a1_try04278 EID 
                a1_try04278 = EVAL(a.At(1-1))
                /* ERROR PROTECTION INSERTED (a1-Result) */
                if ErrorIn(a1_try04278) {Result = a1_try04278
                } else {
                a1 = ANY(a1_try04278)
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
              var g0428I *ClaireBoolean  
              /* and:7 */{ 
                var v_and7 *ClaireBoolean  
                
                v_and7 = Equal(C_string.Id(),a2.Isa.Id())
                if (v_and7 == CFALSE) {g0428I = CFALSE
                } else /* arg:8 */{ 
                  if (C_integer.Id() == i.Isa.Id()) /* If:9 */{ 
                    /* Let:10 */{ 
                      var g0424 int  = ToInteger(i).Value
                      /* noccur = 1 */
                      v_and7 = Core.F__inf_equal_integer(g0424,ClEnv.Verbose)
                      /* Let-10 */} 
                    } else {
                    v_and7 = CFALSE
                    /* If-9 */} 
                  if (v_and7 == CFALSE) {g0428I = CFALSE
                  } else /* arg:9 */{ 
                    g0428I = CTRUE/* arg-9 */} 
                  /* arg-8 */} 
                /* and-7 */} 
              if (g0428I == CTRUE) /* If:7 */{ 
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
      var g0429I *ClaireBoolean  
      var g0429I_try04303 EID 
      /* and:3 */{ 
        var v_and3 *ClaireBoolean  
        
        v_and3 = Core.F__sup_integer(a.Length(),0)
        if (v_and3 == CFALSE) {g0429I_try04303 = EID{CFALSE.Id(),0}
        } else /* arg:4 */{ 
          v_and3 = Core.F_known_ask_any(Core.F_get_property(C_ctrace,ToObject(ClEnv.Id())))
          if (v_and3 == CFALSE) {g0429I_try04303 = EID{CFALSE.Id(),0}
          } else /* arg:5 */{ 
            var v_and3_try04316 EID 
            /* Let:6 */{ 
              var g0432UU *ClaireBoolean  
              /* noccur = 1 */
              var g0432UU_try04337 EID 
              /* Let:7 */{ 
                var g0434UU *ClaireAny  
                /* noccur = 1 */
                var g0434UU_try04358 EID 
                g0434UU_try04358 = EVAL(a.At(1-1))
                /* ERROR PROTECTION INSERTED (g0434UU-g0432UU_try04337) */
                if ErrorIn(g0434UU_try04358) {g0432UU_try04337 = g0434UU_try04358
                } else {
                g0434UU = ANY(g0434UU_try04358)
                g0432UU_try04337 = EID{F_boolean_I_any(g0434UU).Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (g0432UU-v_and3_try04316) */
              if ErrorIn(g0432UU_try04337) {v_and3_try04316 = g0432UU_try04337
              } else {
              g0432UU = ToBoolean(OBJ(g0432UU_try04337))
              v_and3_try04316 = EID{Core.F__I_equal_any(g0432UU.Id(),CTRUE.Id()).Id(),0}
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (v_and3-g0429I_try04303) */
            if ErrorIn(v_and3_try04316) {g0429I_try04303 = v_and3_try04316
            } else {
            v_and3 = ToBoolean(OBJ(v_and3_try04316))
            if (v_and3 == CFALSE) {g0429I_try04303 = EID{CFALSE.Id(),0}
            } else /* arg:6 */{ 
              g0429I_try04303 = EID{CTRUE.Id(),0}/* arg-6 */} 
            /* arg-5 */} 
          /* arg-4 */} 
        }
        /* and-3 */} 
      /* ERROR PROTECTION INSERTED (g0429I-Result) */
      if ErrorIn(g0429I_try04303) {Result = g0429I_try04303
      } else {
      g0429I = ToBoolean(OBJ(g0429I_try04303))
      if (g0429I == CTRUE) /* If:3 */{ 
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
    var g0436I *ClaireBoolean  
    var g0436I_try04372 EID 
    /* Let:2 */{ 
      var g0438UU *ClaireAny  
      /* noccur = 1 */
      var g0438UU_try04393 EID 
      g0438UU_try04393 = EVAL(self.Args.At(1-1))
      /* ERROR PROTECTION INSERTED (g0438UU-g0436I_try04372) */
      if ErrorIn(g0438UU_try04393) {g0436I_try04372 = g0438UU_try04393
      } else {
      g0438UU = ANY(g0438UU_try04393)
      g0436I_try04372 = EID{Core.F__I_equal_any(g0438UU,CFALSE.Id()).Id(),0}
      }
      /* Let-2 */} 
    /* ERROR PROTECTION INSERTED (g0436I-Result) */
    if ErrorIn(g0436I_try04372) {Result = g0436I_try04372
    } else {
    g0436I = ToBoolean(OBJ(g0436I_try04372))
    if (g0436I == CTRUE) /* If:2 */{ 
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