/***** CLAIRE Compilation of file /Users/ycaseau/claire/v4.0/meta/control.cl 
         [version 4.0.02 / safety 5] Friday 12-24-2021 *****/

package Language
import (_ "fmt"
	. "Kernel"
	"Core"
)

//-------- dumb function to prevent import errors --------
func import_g0125() { 
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
/* {1} The go function for: self_print(self:If) [status=1] */
func (self *If ) SelfPrint () EID { 
    var Result EID 
    PRINC("(")
    Core.C_pretty.Index = (Core.C_pretty.Index+1)
    /*integer->integer*//*g_try(v2:"Result",loop:true) */
    Result = self.Printstat()
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    Core.C_pretty.Index = (Core.C_pretty.Index-1)
    /*integer->integer*/PRINC(")")
    Result = EVOID
    }
    return Result} 
  
// The EID go function for: self_print @ If (throw: true) 
func E_self_print_If_Language (self EID) EID { 
    return To_If(OBJ(self)).SelfPrint( )} 
  
/* {1} The go function for: printstat(self:If) [status=1] */
func (self *If ) Printstat () EID { 
    var Result EID 
    PRINC("if ")
    /*g_try(v2:"Result",loop:true) */
    /*g_try(v2:"Result",loop:true) */
    Result = F_printexp_any(self.Test,CFALSE)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    Result = F_checkfar_void()
    }
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(" ")
    /*g_try(v2:"Result",loop:true) */
    Result = F_printif_any(self.Arg)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    Core.C_pretty.Index = (Core.C_pretty.Index-3)
    /*integer->integer*//*g_try(v2:"Result",loop:true) */
    Result = self.Printelse()
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("")
    Result = EVOID
    }}}
    return Result} 
  
// The EID go function for: printstat @ If (throw: true) 
func E_printstat_If (self EID) EID { 
    return To_If(OBJ(self)).Printstat( )} 
  
/* {1} The go function for: printif(self:any) [status=1] */
func F_printif_any (self *ClaireAny ) EID { 
    var Result EID 
    Core.C_pretty.Index = (Core.C_pretty.Index+3)
    /*integer->integer*/if (Core.C_pretty.Pbreak == CTRUE) { 
      { var b_index int  = Core.F_buffer_length_void()
        _ = b_index
        { var _Zl int  = Core.C_pretty.Index
          _ = _Zl
          Core.C_pretty.Pbreak = CFALSE
          /*boolean->boolean*//*g_try(v2:"Result",loop:true) */
          { 
            h_index := ClEnv.Index
            h_base := ClEnv.Base
            Result = Core.F_CALL(C_print,ARGS(self.ToEID()))
            if ErrorIn(Result) && ToType(Core.C_much_too_far.Id()).Contains(ANY(Result)) == CTRUE { 
              ClEnv.Index = h_index
              ClEnv.Base = h_base
              Result = EID{CEMPTY.Id(),0}
              } 
            } 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          Core.C_pretty.Pbreak = CTRUE
          /*boolean->boolean*/if (Core.F_short_enough_integer(Core.F_buffer_length_void()) == CTRUE) { 
            Result = EID{CEMPTY.Id(),0}
            } else {
            Core.F_buffer_set_length_integer(b_index)
            Core.C_pretty.Index = _Zl
            /*integer->integer*//*g_try(v2:"Result",loop:true) */
            Result = F_lbreak_void()
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            Result = Core.F_CALL(C_print,ARGS(self.ToEID()))
            }
            } 
          }
          } 
        } 
      } else {
      Result = Core.F_CALL(C_print,ARGS(self.ToEID()))
      } 
    return Result} 
  
// The EID go function for: printif @ any (throw: true) 
func E_printif_any (self EID) EID { 
    return F_printif_any(ANY(self) )} 
  
/* {1} The go function for: printelse(self:If) [status=1] */
func (self *If ) Printelse () EID { 
    var Result EID 
    { var e *ClaireAny   = self.Other
      if (e.Isa.IsIn(C_If) == CTRUE) { 
        { var g0131 *If   = To_If(e)
          PRINC(" ")
          /*g_try(v2:"Result",loop:true) */
          Result = F_lbreak_void()
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC("else if ")
          /*g_try(v2:"Result",loop:true) */
          Result = F_printexp_any(g0131.Test,CFALSE)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(" ")
          /*g_try(v2:"Result",loop:true) */
          Result = F_printif_any(g0131.Arg)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          Core.C_pretty.Index = (Core.C_pretty.Index-3)
          /*integer->integer*//*g_try(v2:"Result",loop:true) */
          Result = g0131.Printelse()
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC("")
          Result = EVOID
          }}}}
          } 
        }  else if (Equal(e,CNIL.Id()) != CTRUE) { 
        { var _Zl int  = Core.C_pretty.Index
          _ = _Zl
          /*g_try(v2:"Result",loop:true) */
          PRINC(" ")
          /*g_try(v2:"Result",loop:true) */
          Result = F_lbreak_void()
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC("else ")
          F_set_level_integer(1)
          /*g_try(v2:"Result",loop:true) */
          Result = Core.F_CALL(C_print,ARGS(e.ToEID()))
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC("")
          Result = EVOID
          }}
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          { 
            var va_arg1 *Core.PrettyPrinter  
            var va_arg2 int 
            va_arg1 = Core.C_pretty
            va_arg2 = _Zl
            va_arg1.Index = va_arg2
            /*integer->integer*/Result = EID{C__INT,IVAL(va_arg2)}
            } 
          }
          } 
        } else {
        Result = EID{CFALSE.Id(),0}
        } 
      } 
    return Result} 
  
// The EID go function for: printelse @ If (throw: true) 
func E_printelse_If (self EID) EID { 
    return To_If(OBJ(self)).Printelse( )} 
  
// notice that the eval(test) is not a boolean thus the compiler will add
// something
// TODO: check that is is not too slow (may use a constant for _oid_(true))
/* {1} The go function for: self_eval(self:If) [status=1] */
func (self *If ) SelfEval () EID { 
    var Result EID 
    { var x *ClaireAny  
      var try_1 EID 
      /*g_try(v2:"try_1",loop:false) */
      try_1 = EVAL(self.Test)
      /* ERROR PROTECTION INSERTED (x-Result) */
      if ErrorIn(try_1) {Result = try_1
      } else {
      x = ANY(try_1)
      if (x == CTRUE.Id()) { 
        Result = EVAL(self.Arg)
        }  else if (x == CFALSE.Id()) { 
        Result = EVAL(self.Other)
        }  else if (F_boolean_I_any(x) == CTRUE) { 
        Result = EVAL(self.Arg)
        } else {
        Result = EVAL(self.Other)
        } 
      }
      } 
    return Result} 
  
// The EID go function for: self_eval @ If (throw: true) 
func E_self_eval_If (self EID) EID { 
    return To_If(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: If 
func EVAL_If (x *ClaireAny) EID { 
     return To_If(x).SelfEval()} 
  
//--------------------- block structure------------------------------
/* {1} The go function for: self_print(self:Do) [status=1] */
func (self *Do ) SelfPrint () EID { 
    var Result EID 
    { var _Zl int  = Core.C_pretty.Index
      _ = _Zl
      PRINC("(")
      F_set_level_integer(1)
      /*g_try(v2:"Result",loop:true) */
      Result = F_printdo_list(self.Args,CTRUE)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      { 
        var va_arg1 *Core.PrettyPrinter  
        var va_arg2 int 
        va_arg1 = Core.C_pretty
        va_arg2 = _Zl
        va_arg1.Index = va_arg2
        /*integer->integer*/Result = EID{C__INT,IVAL(va_arg2)}
        } 
      }
      } 
    return Result} 
  
// The EID go function for: self_print @ Do (throw: true) 
func E_self_print_Do_Language (self EID) EID { 
    return To_Do(OBJ(self)).SelfPrint( )} 
  
/* {1} The go function for: printdo(l:list,clo:boolean) [status=1] */
func F_printdo_list (l *ClaireList ,clo *ClaireBoolean ) EID { 
    var Result EID 
    { var n int  = l.Length()
      _ = n
      { 
        var x *ClaireAny  
        _ = x
        Result= EID{CFALSE.Id(),0}
        var x_support *ClaireList  
        x_support = l
        x_len := x_support.Length()
        for i_it := 0; i_it < x_len; i_it++ { 
          x = x_support.At(i_it)
          var loop_1 EID 
          _ = loop_1
          { 
          /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
          if (x.Isa.IsIn(C_If) == CTRUE) { 
            { var g0133 *If   = To_If(x)
              _ = g0133
              loop_1 = g0133.Printstat()
              } 
            } else {
            loop_1 = Core.F_CALL(C_print,ARGS(x.ToEID()))
            } 
          /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
          if ErrorIn(loop_1) {Result = loop_1
          break
          } else {
          n = (n-1)
          /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
          if (n == 0) { 
            if (clo == CTRUE) { 
              PRINC(")")
              loop_1 = EVOID
              } else {
              loop_1 = EID{CFALSE.Id(),0}
              } 
            } else {
            PRINC(", ")
            /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
            loop_1 = F_lbreak_void()
            /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
            if ErrorIn(loop_1) {Result = loop_1
            break
            } else {
            }
            } 
          /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
          if ErrorIn(loop_1) {Result = loop_1
          break
          } else {
          }}
          }
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: printdo @ list (throw: true) 
func E_printdo_list (l EID,clo EID) EID { 
    return F_printdo_list(ToList(OBJ(l)),ToBoolean(OBJ(clo)) )} 
  
/* {1} The go function for: printblock(x:any) [status=1] */
func F_printblock_any (x *ClaireAny ) EID { 
    var Result EID 
    if (x.Isa.IsIn(C_Do) == CTRUE) { 
      { var g0135 *Do   = To_Do(x)
        _ = g0135
        Result = F_printdo_list(g0135.Args,CFALSE)
        } 
      }  else if (x.Isa.IsIn(C_If) == CTRUE) { 
      { var g0136 *If   = To_If(x)
        _ = g0136
        Result = g0136.Printstat()
        } 
      } else {
      Result = Core.F_CALL(C_print,ARGS(x.ToEID()))
      } 
    return Result} 
  
// The EID go function for: printblock @ any (throw: true) 
func E_printblock_any (x EID) EID { 
    return F_printblock_any(ANY(x) )} 
  
// use res:EID pragma when compiled with CLAIRE4, res:any for CLAIRE3
/* {1} The go function for: self_eval(self:Do) [status=1] */
func (self *Do ) SelfEval () EID { 
    var Result EID 
    { var res *ClaireAny   = CEMPTY.Id()
      _ = res
      /*g_try(v2:"Result",loop:true) */
      { 
        var _Zx *ClaireAny  
        _ = _Zx
        Result= EID{CFALSE.Id(),0}
        var _Zx_support *ClaireList  
        _Zx_support = self.Args
        _Zx_len := _Zx_support.Length()
        for i_it := 0; i_it < _Zx_len; i_it++ { 
          _Zx = _Zx_support.At(i_it)
          var loop_1 EID 
          _ = loop_1
          var try_2 EID 
          /*g_try(v2:"try_2",loop:tuple("Result", EID)) */
          try_2 = EVAL(_Zx)
          /* ERROR PROTECTION INSERTED (res-Result) */
          if ErrorIn(try_2) {Result = try_2
          break
          } else {
          res = ANY(try_2)
          loop_1 = res.ToEID()
          }
          } 
        } 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = res.ToEID()
      }
      } 
    return Result} 
  
// The EID go function for: self_eval @ Do (throw: true) 
func E_self_eval_Do (self EID) EID { 
    return To_Do(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Do 
func EVAL_Do (x *ClaireAny) EID { 
     return To_Do(x).SelfEval()} 
  
// ----------------- lexical variable definition -----------------------
/* {1} The go function for: self_print(self:Let) [status=1] */
func (self *Let ) SelfPrint () EID { 
    var Result EID 
    { var _Zl int  = Core.C_pretty.Index
      _ = _Zl
      F_set_level_integer(1)
      /*g_try(v2:"Result",loop:true) */
      PRINC("let ")
      /*g_try(v2:"Result",loop:true) */
      Result = F_ppvariable_Variable(self.ClaireVar)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(" := ")
      /*g_try(v2:"Result",loop:true) */
      Result = F_printexp_any(self.Value,CFALSE)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      /*g_try(v2:"Result",loop:true) */
      Result = self.Printbody()
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC("")
      Result = EVOID
      }}}
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      { 
        var va_arg1 *Core.PrettyPrinter  
        var va_arg2 int 
        va_arg1 = Core.C_pretty
        va_arg2 = _Zl
        va_arg1.Index = va_arg2
        /*integer->integer*/Result = EID{C__INT,IVAL(va_arg2)}
        } 
      }
      } 
    return Result} 
  
// The EID go function for: self_print @ Let (throw: true) 
func E_self_print_Let_Language (self EID) EID { 
    return To_Let(OBJ(self)).SelfPrint( )} 
  
/* {1} The go function for: printbody(self:Let) [status=1] */
func (self *Let ) Printbody () EID { 
    var Result EID 
    { var a *ClaireAny   = self.Arg
      if (a.Isa.IsIn(C_Let) == CTRUE) { 
        { var g0139 *Let   = To_Let(a)
          PRINC(",")
          /*g_try(v2:"Result",loop:true) */
          Result = F_lbreak_integer(4)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          /*g_try(v2:"Result",loop:true) */
          Result = F_ppvariable_Variable(g0139.ClaireVar)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(" := ")
          /*g_try(v2:"Result",loop:true) */
          Result = F_printexp_any(g0139.Value,CFALSE)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          Core.C_pretty.Index = (Core.C_pretty.Index-4)
          /*integer->integer*//*g_try(v2:"Result",loop:true) */
          Result = g0139.Printbody()
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC("")
          Result = EVOID
          }}}}
          } 
        } else {
        PRINC(" in ")
        /*g_try(v2:"Result",loop:true) */
        Result = F_lbreak_integer(2)
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        /*g_try(v2:"Result",loop:true) */
        Result = Core.F_CALL(C_print,ARGS(a.ToEID()))
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC("")
        Result = EVOID
        }}
        } 
      } 
    return Result} 
  
// The EID go function for: printbody @ Let (throw: true) 
func E_printbody_Let (self EID) EID { 
    return To_Let(OBJ(self)).Printbody( )} 
  
/* {1} The go function for: self_eval(self:Let) [status=1] */
func (self *Let ) SelfEval () EID { 
    var Result EID 
    { var val *ClaireAny  
      _ = val
      var try_1 EID 
      /*g_try(v2:"try_1",loop:false) */
      try_1 = EVAL(self.Value)
      /* ERROR PROTECTION INSERTED (val-Result) */
      if ErrorIn(try_1) {Result = try_1
      } else {
      val = ANY(try_1)
      /*g_try(v2:"Result",loop:true) */
      Result = F_write_value_Variable(self.ClaireVar,val)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = EVAL(self.Arg)
      }
      }
      } 
    return Result} 
  
// The EID go function for: self_eval @ Let (throw: true) 
func E_self_eval_Let (self EID) EID { 
    return To_Let(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Let 
func EVAL_Let (x *ClaireAny) EID { 
     return To_Let(x).SelfEval()} 
  
// a when is a special Let that filters out the unknown value !
//
/* {1} The go function for: self_print(self:When) [status=1] */
func (self *When ) SelfPrint () EID { 
    var Result EID 
    { var _Zl int  = Core.C_pretty.Index
      _ = _Zl
      F_set_level_integer(1)
      /*g_try(v2:"Result",loop:true) */
      PRINC("when ")
      /*g_try(v2:"Result",loop:true) */
      Result = F_ppvariable_Variable(self.ClaireVar)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(" := ")
      /*g_try(v2:"Result",loop:true) */
      Result = F_printexp_any(self.Value,CFALSE)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(" in ")
      /*g_try(v2:"Result",loop:true) */
      Result = F_lbreak_integer(2)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      /*g_try(v2:"Result",loop:true) */
      Result = Core.F_CALL(C_print,ARGS(self.Arg.ToEID()))
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC("")
      Result = EVOID
      }}}}
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      /*g_try(v2:"Result",loop:true) */
      if (self.Other != CNULL) { 
        PRINC(" ")
        /*g_try(v2:"Result",loop:true) */
        Result = F_lbreak_void()
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC("else ")
        F_set_level_integer(1)
        /*g_try(v2:"Result",loop:true) */
        Result = Core.F_CALL(C_print,ARGS(self.Other.ToEID()))
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC("")
        Result = EVOID
        }}
        } else {
        Result = EID{CFALSE.Id(),0}
        } 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      { 
        var va_arg1 *Core.PrettyPrinter  
        var va_arg2 int 
        va_arg1 = Core.C_pretty
        va_arg2 = _Zl
        va_arg1.Index = va_arg2
        /*integer->integer*/Result = EID{C__INT,IVAL(va_arg2)}
        } 
      }}
      } 
    return Result} 
  
// The EID go function for: self_print @ When (throw: true) 
func E_self_print_When_Language (self EID) EID { 
    return To_When(OBJ(self)).SelfPrint( )} 
  
/* {1} The go function for: self_eval(self:When) [status=1] */
func (self *When ) SelfEval () EID { 
    var Result EID 
    { var val *ClaireAny  
      var try_1 EID 
      /*g_try(v2:"try_1",loop:false) */
      try_1 = EVAL(self.Value)
      /* ERROR PROTECTION INSERTED (val-Result) */
      if ErrorIn(try_1) {Result = try_1
      } else {
      val = ANY(try_1)
      { var n int  = ClEnv.Trace_I
        _ = n
        if (val != CNULL) { 
          /*g_try(v2:"Result",loop:true) */
          Result = F_write_value_Variable(self.ClaireVar,val)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          Result = EVAL(self.Arg)
          }
          } else {
          Result = EVAL(self.Other)
          } 
        } 
      }
      } 
    return Result} 
  
// The EID go function for: self_eval @ When (throw: true) 
func E_self_eval_When (self EID) EID { 
    return To_When(OBJ(self)).SelfEval( )} 
  
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
/* {1} The go function for: self_print(self:Let+) [status=1] */
func (self *Let_plus ) SelfPrint () EID { 
    var Result EID 
    { var _Zl int  = Core.C_pretty.Index
      _ = _Zl
      { var l *ClaireList   = To_Do(self.Arg).Args
        F_set_level_integer(1)
        /*g_try(v2:"Result",loop:true) */
        PRINC("let ")
        /*g_try(v2:"Result",loop:true) */
        Result = F_printexp_any(self.Value,CFALSE)
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC(" := ")
        /*g_try(v2:"Result",loop:true) */
        Result = F_printexp_any(ToList(OBJ(Core.F_CALL(C_args,ARGS(l.At(1-1).ToEID())))).At(3-1),CFALSE)
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC(" in ")
        /*g_try(v2:"Result",loop:true) */
        Result = F_lbreak_integer(2)
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        /*g_try(v2:"Result",loop:true) */
        Result = Core.F_CALL(C_print,ARGS(To_Let(l.At(2-1)).Value.ToEID()))
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC("")
        Result = EVOID
        }}}}
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        { 
          var va_arg1 *Core.PrettyPrinter  
          var va_arg2 int 
          va_arg1 = Core.C_pretty
          va_arg2 = _Zl
          va_arg1.Index = va_arg2
          /*integer->integer*/Result = EID{C__INT,IVAL(va_arg2)}
          } 
        }
        } 
      } 
    return Result} 
  
// The EID go function for: self_print @ Let+ (throw: true) 
func E_self_print_Let_plus_Language (self EID) EID { 
    return To_Let_plus(OBJ(self)).SelfPrint( )} 
  
/* {1} The go function for: self_print(self:Let*) [status=1] */
func (self *Let_star ) SelfPrint () EID { 
    var Result EID 
    { var _Zl int  = Core.C_pretty.Index
      _ = _Zl
      { var l *ClaireAny   = self.Arg
        F_set_level_integer(1)
        /*g_try(v2:"Result",loop:true) */
        if (l.Isa.IsIn(C_Let) == CTRUE) { 
          PRINC("let (")
          /*g_try(v2:"Result",loop:true) */
          Result= EID{CFALSE.Id(),0}
          for (CTRUE == CTRUE) { 
            /* While stat, v:"Result" loop:true */
            var loop_1 EID 
            _ = loop_1
            { 
            /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
            loop_1 = Core.F_CALL(C_Language_ppvariable,ARGS(Core.F_CALL(C_var,ARGS(l.ToEID()))))
            /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
            if ErrorIn(loop_1) {Result = loop_1
            break
            } else {
            { var lnext *ClaireAny   = ANY(Core.F_CALL(C_arg,ARGS(l.ToEID())))
              var g0142I *ClaireBoolean  
              if (lnext.Isa.IsIn(C_Let) == CTRUE) { 
                { var g0141 *Let   = To_Let(lnext)
                  g0142I = MakeBoolean((g0141.Value.Isa.IsIn(C_Call) == CTRUE) && (ToList(OBJ(Core.F_CALL(C_args,ARGS(g0141.Value.ToEID())))).At(1-1) == self.ClaireVar.Id()))
                  } 
                } else {
                g0142I = CFALSE
                } 
              if (g0142I == CTRUE) { 
                PRINC(",")
                l = lnext
                } else {
                Result = EID{CTRUE.Id(),0}
                break
                } 
              } 
            }
            /* try?:false, v2:"v_while5" loop will be:tuple("Result", EID) */
            } 
          }
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(") := ")
          /*g_try(v2:"Result",loop:true) */
          Result = F_printexp_any(self.Value,CFALSE)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          /*g_try(v2:"Result",loop:true) */
          Result = To_Let(l).Printbody()
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC("")
          Result = EVOID
          }}}
          } else {
          PRINC("(")
          /*g_try(v2:"Result",loop:true) */
          { var _Zf *ClaireBoolean   = CTRUE
            _ = _Zf
            { 
              var _Za *ClaireAny  
              _ = _Za
              Result= EID{CFALSE.Id(),0}
              var _Za_support *ClaireList  
              _Za_support = ToList(OBJ(Core.F_CALL(C_args,ARGS(l.ToEID()))))
              _Za_len := _Za_support.Length()
              for i_it := 0; i_it < _Za_len; i_it++ { 
                _Za = _Za_support.At(i_it)
                var loop_2 EID 
                _ = loop_2
                { 
                if (_Zf == CTRUE) { 
                  _Zf = CFALSE
                  } else {
                  PRINC(",")
                  } 
                /*g_try(v2:"loop_2",loop:tuple("Result", EID)) */
                loop_2 = Core.F_CALL(C_Language_ppvariable,ARGS(Core.F_CALL(C_var,ARGS(_Za.ToEID()))))
                /* ERROR PROTECTION INSERTED (loop_2-loop_2) */
                if ErrorIn(loop_2) {Result = loop_2
                break
                } else {
                }
                }
                } 
              } 
            } 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(") := ")
          /*g_try(v2:"Result",loop:true) */
          Result = F_printexp_any(self.Value,CFALSE)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC("")
          Result = EVOID
          }}
          } 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        { 
          var va_arg1 *Core.PrettyPrinter  
          var va_arg2 int 
          va_arg1 = Core.C_pretty
          va_arg2 = _Zl
          va_arg1.Index = va_arg2
          /*integer->integer*/Result = EID{C__INT,IVAL(va_arg2)}
          } 
        }
        } 
      } 
    return Result} 
  
// The EID go function for: self_print @ Let* (throw: true) 
func E_self_print_Let_star_Language (self EID) EID { 
    return To_Let_star(OBJ(self)).SelfPrint( )} 
  
// *********************************************************************
// *     Part 2: set control structures                                *
// *********************************************************************
// for is the simplest evaluation loop
//
/* {1} The go function for: self_print(self:For) [status=1] */
func (self *For ) SelfPrint () EID { 
    var Result EID 
    PRINC("for ")
    /*g_try(v2:"Result",loop:true) */
    Result = F_ppvariable_Variable(self.ClaireVar)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(" in ")
    /*g_try(v2:"Result",loop:true) */
    { var _Zl int  = Core.C_pretty.Index
      _ = _Zl
      F_set_level_void()
      /*g_try(v2:"Result",loop:true) */
      Result = F_printexp_any(self.SetArg,CFALSE)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      { 
        var va_arg1 *Core.PrettyPrinter  
        var va_arg2 int 
        va_arg1 = Core.C_pretty
        va_arg2 = _Zl
        va_arg1.Index = va_arg2
        /*integer->integer*/Result = EID{C__INT,IVAL(va_arg2)}
        } 
      }
      } 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(" ")
    /*g_try(v2:"Result",loop:true) */
    Core.C_pretty.Index = (Core.C_pretty.Index+2)
    /*integer->integer*//*g_try(v2:"Result",loop:true) */
    Result = F_lbreak_void()
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    /*g_try(v2:"Result",loop:true) */
    Result = Core.F_CALL(C_print,ARGS(self.Arg.ToEID()))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    { 
      var va_arg1 *Core.PrettyPrinter  
      var va_arg2 int 
      va_arg1 = Core.C_pretty
      va_arg2 = (Core.C_pretty.Index-2)
      va_arg1.Index = va_arg2
      /*integer->integer*/Result = EID{C__INT,IVAL(va_arg2)}
      } 
    }}
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("")
    Result = EVOID
    }}}
    return Result} 
  
// The EID go function for: self_print @ For (throw: true) 
func E_self_print_For_Language (self EID) EID { 
    return To_For(OBJ(self)).SelfPrint( )} 
  
/* {1} The go function for: self_eval(self:For) [status=1] */
func (self *For ) SelfEval () EID { 
    var Result EID 
    { var x *ClaireAny  
      var try_1 EID 
      /*g_try(v2:"try_1",loop:false) */
      try_1 = EVAL(self.SetArg)
      /* ERROR PROTECTION INSERTED (x-Result) */
      if ErrorIn(try_1) {Result = try_1
      } else {
      x = ANY(try_1)
      { 
        h_index := ClEnv.Index
        h_base := ClEnv.Base
        if (C_class.Id() == x.Isa.Id()) { 
          { var g0145 *ClaireClass   = ToClass(x)
            _ = g0145
            { 
              var y *ClaireClass  
              _ = y
              var y_iter *ClaireAny  
              Result= EID{CFALSE.Id(),0}
              var y_support *ClaireSet  
              y_support = g0145.Descendents
              for i_it := 0; i_it < y_support.Count; i_it++ { 
                y_iter = y_support.At(i_it)
                y = ToClass(y_iter)
                var loop_2 EID 
                _ = loop_2
                /*g_try(v2:"loop_2",loop:tuple("Result", EID)) */
                { 
                  var z *ClaireAny  
                  _ = z
                  loop_2= EID{CFALSE.Id(),0}
                  var z_support *ClaireList  
                  z_support = y.Instances
                  z_len := z_support.Length()
                  for i_it := 0; i_it < z_len; i_it++ { 
                    z = z_support.At(i_it)
                    var loop_3 EID 
                    _ = loop_3
                    { 
                    /*g_try(v2:"loop_3",loop:tuple("loop_2", EID)) */
                    loop_3 = F_write_value_Variable(self.ClaireVar,z)
                    /* ERROR PROTECTION INSERTED (loop_3-loop_3) */
                    if ErrorIn(loop_3) {loop_2 = loop_3
                    break
                    } else {
                    /*g_try(v2:"loop_3",loop:tuple("loop_2", EID)) */
                    loop_3 = EVAL(self.Arg)
                    /* ERROR PROTECTION INSERTED (loop_3-loop_3) */
                    if ErrorIn(loop_3) {loop_2 = loop_3
                    break
                    } else {
                    }}
                    }
                    } 
                  } 
                /* ERROR PROTECTION INSERTED (loop_2-Result) */
                if ErrorIn(loop_2) {Result = loop_2
                break
                } else {
                }
                } 
              } 
            } 
          }  else if (x.Isa.IsIn(C_list) == CTRUE) { 
          { var g0146 *ClaireList   = ToList(x)
            _ = g0146
            { 
              var z *ClaireAny  
              _ = z
              Result= EID{CFALSE.Id(),0}
              var z_support *ClaireList  
              z_support = g0146
              z_len := z_support.Length()
              for i_it := 0; i_it < z_len; i_it++ { 
                z = z_support.At(i_it)
                var loop_4 EID 
                _ = loop_4
                { 
                /*g_try(v2:"loop_4",loop:tuple("Result", EID)) */
                loop_4 = F_write_value_Variable(self.ClaireVar,z)
                /* ERROR PROTECTION INSERTED (loop_4-loop_4) */
                if ErrorIn(loop_4) {Result = loop_4
                break
                } else {
                /*g_try(v2:"loop_4",loop:tuple("Result", EID)) */
                loop_4 = EVAL(self.Arg)
                /* ERROR PROTECTION INSERTED (loop_4-loop_4) */
                if ErrorIn(loop_4) {Result = loop_4
                break
                } else {
                }}
                }
                } 
              } 
            } 
          }  else if (x.Isa.IsIn(C_array) == CTRUE) { 
          { var g0147 *ClaireList   = ToArray(x)
            { var n int  = g0147.Length()
              _ = n
              { var g0148 int  = 1
                { var g0149 int  = n
                  _ = g0149
                  Result= EID{CFALSE.Id(),0}
                  for (g0148 <= g0149) { 
                    /* While stat, v:"Result" loop:false */
                    var loop_5 EID 
                    _ = loop_5
                    { 
                    /*g_try(v2:"loop_5",loop:tuple("Result", EID)) */
                    { var z *ClaireAny   = ToList(g0147.Id()).At(g0148-1)
                      _ = z
                      /*g_try(v2:"loop_5",loop:tuple("Result", EID)) */
                      loop_5 = F_write_value_Variable(self.ClaireVar,z)
                      /* ERROR PROTECTION INSERTED (loop_5-loop_5) */
                      if ErrorIn(loop_5) {Result = loop_5
                      break
                      } else {
                      /*g_try(v2:"loop_5",loop:tuple("Result", EID)) */
                      loop_5 = EVAL(self.Arg)
                      /* ERROR PROTECTION INSERTED (loop_5-loop_5) */
                      if ErrorIn(loop_5) {Result = loop_5
                      break
                      } else {
                      }}
                      } 
                    /* ERROR PROTECTION INSERTED (loop_5-loop_5) */
                    if ErrorIn(loop_5) {Result = loop_5
                    break
                    } else {
                    g0148 = (g0148+1)
                    }
                    /* try?:false, v2:"v_while9" loop will be:tuple("Result", EID) */
                    } 
                  }
                  } 
                } 
              } 
            } 
          }  else if (x.Isa.IsIn(C_Interval) == CTRUE) { 
          { var g0150 *ClaireInterval   = To_Interval(x)
            { var y int  = g0150.Arg1
              { var g0151 int  = g0150.Arg2
                _ = g0151
                Result= EID{CFALSE.Id(),0}
                for (y <= g0151) { 
                  /* While stat, v:"Result" loop:false */
                  var loop_6 EID 
                  _ = loop_6
                  { 
                  /*g_try(v2:"loop_6",loop:tuple("Result", EID)) */
                  /*g_try(v2:"loop_6",loop:tuple("Result", EID)) */
                  loop_6 = self.ClaireVar.WriteEID(EID{C__INT,IVAL(y)})
                  /* ERROR PROTECTION INSERTED (loop_6-loop_6) */
                  if ErrorIn(loop_6) {Result = loop_6
                  break
                  } else {
                  /*g_try(v2:"loop_6",loop:tuple("Result", EID)) */
                  loop_6 = EVAL(self.Arg)
                  /* ERROR PROTECTION INSERTED (loop_6-loop_6) */
                  if ErrorIn(loop_6) {Result = loop_6
                  break
                  } else {
                  }}
                  /* ERROR PROTECTION INSERTED (loop_6-loop_6) */
                  if ErrorIn(loop_6) {Result = loop_6
                  break
                  } else {
                  y = (y+1)
                  }
                  /* try?:false, v2:"v_while8" loop will be:tuple("Result", EID) */
                  } 
                }
                } 
              } 
            } 
          }  else if (x.Isa.IsIn(C_collection) == CTRUE) { 
          { var g0152 *ClaireCollection   = ToCollection(x)
            _ = g0152
            { 
              var y *ClaireAny  
              _ = y
              Result= EID{CFALSE.Id(),0}
              var y_support *ClaireList  
              var try_7 EID 
              /*g_try(v2:"try_7",loop:false) */
              try_7 = Core.F_enumerate_any(g0152.Id())
              /* ERROR PROTECTION INSERTED (y_support-Result) */
              if ErrorIn(try_7) {Result = try_7
              } else {
              y_support = ToList(OBJ(try_7))
              y_len := y_support.Length()
              for i_it := 0; i_it < y_len; i_it++ { 
                y = y_support.At(i_it)
                var loop_8 EID 
                _ = loop_8
                { 
                /*g_try(v2:"loop_8",loop:tuple("Result", EID)) */
                loop_8 = F_write_value_Variable(self.ClaireVar,y)
                /* ERROR PROTECTION INSERTED (loop_8-loop_8) */
                if ErrorIn(loop_8) {Result = loop_8
                break
                } else {
                /*g_try(v2:"loop_8",loop:tuple("Result", EID)) */
                loop_8 = EVAL(self.Arg)
                /* ERROR PROTECTION INSERTED (loop_8-loop_8) */
                if ErrorIn(loop_8) {Result = loop_8
                break
                } else {
                }}
                }}
                } 
              } 
            } 
          } else {
          Result = ToException(Core.C_general_error.Make(MakeString("[136] ~S is not a collection !").Id(),MakeConstantList(x).Id())).Close()
          } 
        if ErrorIn(Result) && ToType(Core.C_return_error.Id()).Contains(ANY(Result)) == CTRUE { 
          ClEnv.Index = h_index
          ClEnv.Base = h_base
          Result = Core.F_CALL(C_arg,ARGS(EID{ClEnv.Exception_I.Id(),0}))
          } 
        } 
      }
      } 
    return Result} 
  
// The EID go function for: self_eval @ For (throw: true) 
func E_self_eval_For (self EID) EID { 
    return To_For(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: For 
func EVAL_For (x *ClaireAny) EID { 
     return To_For(x).SelfEval()} 
  
// [collect VAR in SET_EXPR, ...] is the same as a "for", but returns the list of values
//
/* {1} The go function for: self_print(self:Collect) [status=1] */
func (self *Collect ) SelfPrint () EID { 
    var Result EID 
    PRINC("list{ ")
    /*g_try(v2:"Result",loop:true) */
    Core.C_pretty.Index = (Core.C_pretty.Index+2)
    /*integer->integer*/Result = F_printexp_any(self.Arg,CFALSE)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(" | ")
    /*g_try(v2:"Result",loop:true) */
    Result = F_lbreak_void()
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    /*g_try(v2:"Result",loop:true) */
    Result = F_ppvariable_Variable(self.ClaireVar)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(" in ")
    /*g_try(v2:"Result",loop:true) */
    { var _Zl int  = Core.C_pretty.Index
      _ = _Zl
      F_set_level_void()
      /*g_try(v2:"Result",loop:true) */
      Result = F_printexp_any(self.SetArg,CFALSE)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      { 
        var va_arg1 *Core.PrettyPrinter  
        var va_arg2 int 
        va_arg1 = Core.C_pretty
        va_arg2 = (_Zl-2)
        va_arg1.Index = va_arg2
        /*integer->integer*/Result = EID{C__INT,IVAL(va_arg2)}
        } 
      }
      } 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("}")
    Result = EVOID
    }}}}
    return Result} 
  
// The EID go function for: self_print @ Collect (throw: true) 
func E_self_print_Collect_Language (self EID) EID { 
    return To_Collect(OBJ(self)).SelfPrint( )} 
  
// list image : preserve the order for lists and intervals (v4)
/* {1} The go function for: self_eval(self:Collect) [status=1] */
func (self *Collect ) SelfEval () EID { 
    var Result EID 
    { var x *ClaireAny  
      var try_1 EID 
      /*g_try(v2:"try_1",loop:false) */
      try_1 = EVAL(self.SetArg)
      /* ERROR PROTECTION INSERTED (x-Result) */
      if ErrorIn(try_1) {Result = try_1
      } else {
      x = ANY(try_1)
      { var res *ClaireList  
        _ = res
        { var arg_2 *ClaireType  
          _ = arg_2
          if (self.Of.Id() != CNULL) { 
            arg_2 = self.Of
            } else {
            arg_2 = ToType(CEMPTY.Id())
            } 
          res = arg_2.EmptyList()
          } 
        /*g_try(v2:"Result",loop:true) */
        if (C_class.Id() == x.Isa.Id()) { 
          { var g0155 *ClaireClass   = ToClass(x)
            _ = g0155
            { 
              var y *ClaireClass  
              _ = y
              var y_iter *ClaireAny  
              Result= EID{CFALSE.Id(),0}
              var y_support *ClaireSet  
              y_support = g0155.Descendents
              for i_it := 0; i_it < y_support.Count; i_it++ { 
                y_iter = y_support.At(i_it)
                y = ToClass(y_iter)
                var loop_3 EID 
                _ = loop_3
                /*g_try(v2:"loop_3",loop:tuple("Result", EID)) */
                { 
                  var z *ClaireAny  
                  _ = z
                  loop_3= EID{CFALSE.Id(),0}
                  var z_support *ClaireList  
                  z_support = y.Instances
                  z_len := z_support.Length()
                  for i_it := 0; i_it < z_len; i_it++ { 
                    z = z_support.At(i_it)
                    var loop_4 EID 
                    _ = loop_4
                    { 
                    /*g_try(v2:"loop_4",loop:tuple("loop_3", EID)) */
                    loop_4 = F_write_value_Variable(self.ClaireVar,z)
                    /* ERROR PROTECTION INSERTED (loop_4-loop_4) */
                    if ErrorIn(loop_4) {loop_3 = loop_4
                    break
                    } else {
                    var try_5 EID 
                    /*g_try(v2:"try_5",loop:tuple("loop_3", EID)) */
                    { var arg_6 *ClaireAny  
                      _ = arg_6
                      var try_7 EID 
                      /*g_try(v2:"try_7",loop:false) */
                      try_7 = EVAL(self.Arg)
                      /* ERROR PROTECTION INSERTED (arg_6-try_5) */
                      if ErrorIn(try_7) {try_5 = try_7
                      } else {
                      arg_6 = ANY(try_7)
                      try_5 = EID{res.AddFast(arg_6).Id(),0}/*t=any,s=EID*/
                      }
                      } 
                    /* ERROR PROTECTION INSERTED (res-loop_4) */
                    if ErrorIn(try_5) {loop_4 = try_5
                    loop_3 = try_5
                    break
                    } else {
                    res = ToList(OBJ(try_5))
                    loop_4 = EID{res.Id(),0}
                    }}
                    }
                    } 
                  } 
                /* ERROR PROTECTION INSERTED (loop_3-Result) */
                if ErrorIn(loop_3) {Result = loop_3
                break
                } else {
                }
                } 
              } 
            } 
          }  else if (x.Isa.IsIn(C_list) == CTRUE) { 
          { var g0156 *ClaireList   = ToList(x)
            _ = g0156
            { 
              var y *ClaireAny  
              _ = y
              Result= EID{CFALSE.Id(),0}
              var y_support *ClaireList  
              y_support = g0156
              y_len := y_support.Length()
              for i_it := 0; i_it < y_len; i_it++ { 
                y = y_support.At(i_it)
                var loop_8 EID 
                _ = loop_8
                { 
                /*g_try(v2:"loop_8",loop:tuple("Result", EID)) */
                loop_8 = F_write_value_Variable(self.ClaireVar,y)
                /* ERROR PROTECTION INSERTED (loop_8-loop_8) */
                if ErrorIn(loop_8) {Result = loop_8
                break
                } else {
                var try_9 EID 
                /*g_try(v2:"try_9",loop:tuple("Result", EID)) */
                { var arg_10 *ClaireAny  
                  _ = arg_10
                  var try_11 EID 
                  /*g_try(v2:"try_11",loop:false) */
                  try_11 = EVAL(self.Arg)
                  /* ERROR PROTECTION INSERTED (arg_10-try_9) */
                  if ErrorIn(try_11) {try_9 = try_11
                  } else {
                  arg_10 = ANY(try_11)
                  try_9 = EID{res.AddFast(arg_10).Id(),0}/*t=any,s=EID*/
                  }
                  } 
                /* ERROR PROTECTION INSERTED (res-loop_8) */
                if ErrorIn(try_9) {loop_8 = try_9
                Result = try_9
                break
                } else {
                res = ToList(OBJ(try_9))
                loop_8 = EID{res.Id(),0}
                }}
                }
                } 
              } 
            } 
          }  else if (x.Isa.IsIn(C_Interval) == CTRUE) { 
          { var g0157 *ClaireInterval   = To_Interval(x)
            { var y int  = g0157.Arg1
              { var g0158 int  = g0157.Arg2
                _ = g0158
                Result= EID{CFALSE.Id(),0}
                for (y <= g0158) { 
                  /* While stat, v:"Result" loop:true */
                  var loop_12 EID 
                  _ = loop_12
                  { 
                  /*g_try(v2:"loop_12",loop:tuple("Result", EID)) */
                  /*g_try(v2:"loop_12",loop:tuple("Result", EID)) */
                  loop_12 = self.ClaireVar.WriteEID(EID{C__INT,IVAL(y)})
                  /* ERROR PROTECTION INSERTED (loop_12-loop_12) */
                  if ErrorIn(loop_12) {Result = loop_12
                  break
                  } else {
                  var try_13 EID 
                  /*g_try(v2:"try_13",loop:tuple("Result", EID)) */
                  { var arg_14 *ClaireAny  
                    _ = arg_14
                    var try_15 EID 
                    /*g_try(v2:"try_15",loop:false) */
                    try_15 = EVAL(self.Arg)
                    /* ERROR PROTECTION INSERTED (arg_14-try_13) */
                    if ErrorIn(try_15) {try_13 = try_15
                    } else {
                    arg_14 = ANY(try_15)
                    try_13 = EID{res.AddFast(arg_14).Id(),0}/*t=any,s=EID*/
                    }
                    } 
                  /* ERROR PROTECTION INSERTED (res-loop_12) */
                  if ErrorIn(try_13) {loop_12 = try_13
                  Result = try_13
                  break
                  } else {
                  res = ToList(OBJ(try_13))
                  loop_12 = EID{res.Id(),0}
                  }}
                  /* ERROR PROTECTION INSERTED (loop_12-loop_12) */
                  if ErrorIn(loop_12) {Result = loop_12
                  break
                  } else {
                  y = (y+1)
                  }
                  /* try?:false, v2:"v_while8" loop will be:tuple("Result", EID) */
                  } 
                }
                } 
              } 
            } 
          } else {
          { 
            var y *ClaireAny  
            _ = y
            Result= EID{CFALSE.Id(),0}
            var y_support *ClaireList  
            var try_16 EID 
            /*g_try(v2:"try_16",loop:false) */
            try_16 = Core.F_enumerate_any(x)
            /* ERROR PROTECTION INSERTED (y_support-Result) */
            if ErrorIn(try_16) {Result = try_16
            } else {
            y_support = ToList(OBJ(try_16))
            y_len := y_support.Length()
            for i_it := 0; i_it < y_len; i_it++ { 
              y = y_support.At(i_it)
              var loop_17 EID 
              _ = loop_17
              { 
              /*g_try(v2:"loop_17",loop:tuple("Result", EID)) */
              loop_17 = F_write_value_Variable(self.ClaireVar,y)
              /* ERROR PROTECTION INSERTED (loop_17-loop_17) */
              if ErrorIn(loop_17) {Result = loop_17
              break
              } else {
              var try_18 EID 
              /*g_try(v2:"try_18",loop:tuple("Result", EID)) */
              { var arg_19 *ClaireAny  
                _ = arg_19
                var try_20 EID 
                /*g_try(v2:"try_20",loop:false) */
                try_20 = EVAL(self.Arg)
                /* ERROR PROTECTION INSERTED (arg_19-try_18) */
                if ErrorIn(try_20) {try_18 = try_20
                } else {
                arg_19 = ANY(try_20)
                try_18 = EID{res.AddFast(arg_19).Id(),0}/*t=any,s=EID*/
                }
                } 
              /* ERROR PROTECTION INSERTED (res-loop_17) */
              if ErrorIn(try_18) {loop_17 = try_18
              Result = try_18
              break
              } else {
              res = ToList(OBJ(try_18))
              loop_17 = EID{res.Id(),0}
              }}
              }}
              } 
            } 
          } 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        Result = EID{res.Id(),0}
        }
        } 
      }
      } 
    return Result} 
  
// The EID go function for: self_eval @ Collect (throw: true) 
func E_self_eval_Collect (self EID) EID { 
    return To_Collect(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Collect 
func EVAL_Collect (x *ClaireAny) EID { 
     return To_Collect(x).SelfEval()} 
  
// this is a set image version, that produces a set
//
/* {1} The go function for: self_print(self:Image) [status=1] */
func (self *Image ) SelfPrint () EID { 
    var Result EID 
    PRINC("{ ")
    /*g_try(v2:"Result",loop:true) */
    Core.C_pretty.Index = (Core.C_pretty.Index+2)
    /*integer->integer*/Result = F_printexp_any(self.Arg,CFALSE)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(" | ")
    /*g_try(v2:"Result",loop:true) */
    Result = F_lbreak_void()
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    /*g_try(v2:"Result",loop:true) */
    Result = F_ppvariable_Variable(self.ClaireVar)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(" in ")
    /*g_try(v2:"Result",loop:true) */
    { var _Zl int  = Core.C_pretty.Index
      _ = _Zl
      F_set_level_void()
      /*g_try(v2:"Result",loop:true) */
      Result = F_printexp_any(self.SetArg,CFALSE)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      { 
        var va_arg1 *Core.PrettyPrinter  
        var va_arg2 int 
        va_arg1 = Core.C_pretty
        va_arg2 = (_Zl-2)
        va_arg1.Index = va_arg2
        /*integer->integer*/Result = EID{C__INT,IVAL(va_arg2)}
        } 
      }
      } 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("}")
    Result = EVOID
    }}}}
    return Result} 
  
// The EID go function for: self_print @ Image (throw: true) 
func E_self_print_Image_Language (self EID) EID { 
    return To_Image(OBJ(self)).SelfPrint( )} 
  
/* {1} The go function for: self_eval(self:Image) [status=1] */
func (self *Image ) SelfEval () EID { 
    var Result EID 
    { var x *ClaireAny  
      _ = x
      var try_1 EID 
      /*g_try(v2:"try_1",loop:false) */
      try_1 = EVAL(self.SetArg)
      /* ERROR PROTECTION INSERTED (x-Result) */
      if ErrorIn(try_1) {Result = try_1
      } else {
      x = ANY(try_1)
      { var res *ClaireSet  
        _ = res
        { var arg_2 *ClaireType  
          _ = arg_2
          if (self.Of.Id() != CNULL) { 
            arg_2 = self.Of
            } else {
            arg_2 = ToType(CEMPTY.Id())
            } 
          res = arg_2.EmptySet()
          } 
        /*g_try(v2:"Result",loop:true) */
        { 
          var y *ClaireAny  
          _ = y
          Result= EID{CFALSE.Id(),0}
          var y_support *ClaireList  
          var try_3 EID 
          /*g_try(v2:"try_3",loop:false) */
          try_3 = Core.F_enumerate_any(x)
          /* ERROR PROTECTION INSERTED (y_support-Result) */
          if ErrorIn(try_3) {Result = try_3
          } else {
          y_support = ToList(OBJ(try_3))
          y_len := y_support.Length()
          for i_it := 0; i_it < y_len; i_it++ { 
            y = y_support.At(i_it)
            var loop_4 EID 
            _ = loop_4
            { 
            /*g_try(v2:"loop_4",loop:tuple("Result", EID)) */
            loop_4 = F_write_value_Variable(self.ClaireVar,y)
            /* ERROR PROTECTION INSERTED (loop_4-loop_4) */
            if ErrorIn(loop_4) {Result = loop_4
            break
            } else {
            var try_5 EID 
            /*g_try(v2:"try_5",loop:tuple("Result", EID)) */
            { var arg_6 *ClaireAny  
              _ = arg_6
              var try_7 EID 
              /*g_try(v2:"try_7",loop:false) */
              try_7 = EVAL(self.Arg)
              /* ERROR PROTECTION INSERTED (arg_6-try_5) */
              if ErrorIn(try_7) {try_5 = try_7
              } else {
              arg_6 = ANY(try_7)
              try_5 = EID{res.AddFast(arg_6).Id(),0}/*t=any,s=EID*/
              }
              } 
            /* ERROR PROTECTION INSERTED (res-loop_4) */
            if ErrorIn(try_5) {loop_4 = try_5
            Result = try_5
            break
            } else {
            res = ToSet(OBJ(try_5))
            loop_4 = EID{res.Id(),0}
            }}
            }}
            } 
          } 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        Result = EID{res.Id(),0}
        }
        } 
      }
      } 
    return Result} 
  
// The EID go function for: self_eval @ Image (throw: true) 
func E_self_eval_Image (self EID) EID { 
    return To_Image(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Image 
func EVAL_Image (x *ClaireAny) EID { 
     return To_Image(x).SelfEval()} 
  
// [select VAR in SET_EXPR, ...] is the same as a "for" but returns the subset of
//  members that produce a true value
//
/* {1} The go function for: self_print(self:Select) [status=1] */
func (self *Select ) SelfPrint () EID { 
    var Result EID 
    PRINC("{ ")
    /*g_try(v2:"Result",loop:true) */
    Result = F_ppvariable_Variable(self.ClaireVar)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(" in ")
    /*g_try(v2:"Result",loop:true) */
    { var _Zl int  = Core.C_pretty.Index
      _ = _Zl
      F_set_level_void()
      /*g_try(v2:"Result",loop:true) */
      Result = F_printexp_any(self.SetArg,CFALSE)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      { 
        var va_arg1 *Core.PrettyPrinter  
        var va_arg2 int 
        va_arg1 = Core.C_pretty
        va_arg2 = _Zl
        va_arg1.Index = va_arg2
        /*integer->integer*/Result = EID{C__INT,IVAL(va_arg2)}
        } 
      }
      } 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(" | ")
    /*g_try(v2:"Result",loop:true) */
    /*g_try(v2:"Result",loop:true) */
    Result = F_lbreak_integer(2)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    /*g_try(v2:"Result",loop:true) */
    Result = Core.F_CALL(C_print,ARGS(self.Arg.ToEID()))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    { 
      var va_arg1 *Core.PrettyPrinter  
      var va_arg2 int 
      va_arg1 = Core.C_pretty
      va_arg2 = (Core.C_pretty.Index-2)
      va_arg1.Index = va_arg2
      /*integer->integer*/Result = EID{C__INT,IVAL(va_arg2)}
      } 
    }}
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("}")
    Result = EVOID
    }}}
    return Result} 
  
// The EID go function for: self_print @ Select (throw: true) 
func E_self_print_Select_Language (self EID) EID { 
    return To_Select(OBJ(self)).SelfPrint( )} 
  
/* {1} The go function for: self_eval(self:Select) [status=1] */
func (self *Select ) SelfEval () EID { 
    var Result EID 
    { var x *ClaireAny  
      var try_1 EID 
      /*g_try(v2:"try_1",loop:false) */
      try_1 = EVAL(self.SetArg)
      /* ERROR PROTECTION INSERTED (x-Result) */
      if ErrorIn(try_1) {Result = try_1
      } else {
      x = ANY(try_1)
      { var res *ClaireSet  
        _ = res
        { var arg_2 *ClaireType  
          _ = arg_2
          if (self.Of.Id() != CNULL) { 
            arg_2 = self.Of
            } else {
            arg_2 = ToType(CEMPTY.Id())
            } 
          res = arg_2.EmptySet()
          } 
        /*g_try(v2:"Result",loop:true) */
        if (C_class.Id() == x.Isa.Id()) { 
          { var g0162 *ClaireClass   = ToClass(x)
            _ = g0162
            { 
              var y *ClaireClass  
              _ = y
              var y_iter *ClaireAny  
              Result= EID{CFALSE.Id(),0}
              var y_support *ClaireSet  
              y_support = g0162.Descendents
              for i_it := 0; i_it < y_support.Count; i_it++ { 
                y_iter = y_support.At(i_it)
                y = ToClass(y_iter)
                var loop_3 EID 
                _ = loop_3
                /*g_try(v2:"loop_3",loop:tuple("Result", EID)) */
                { 
                  var z *ClaireAny  
                  _ = z
                  loop_3= EID{CFALSE.Id(),0}
                  var z_support *ClaireList  
                  z_support = y.Instances
                  z_len := z_support.Length()
                  for i_it := 0; i_it < z_len; i_it++ { 
                    z = z_support.At(i_it)
                    var loop_4 EID 
                    _ = loop_4
                    { 
                    /*g_try(v2:"loop_4",loop:tuple("loop_3", EID)) */
                    loop_4 = F_write_value_Variable(self.ClaireVar,z)
                    /* ERROR PROTECTION INSERTED (loop_4-loop_4) */
                    if ErrorIn(loop_4) {loop_3 = loop_4
                    break
                    } else {
                    /*g_try(v2:"loop_4",loop:tuple("loop_3", EID)) */
                    var g0166I *ClaireBoolean  
                    var try_5 EID 
                    /*g_try(v2:"try_5",loop:false) */
                    { var arg_6 *ClaireAny  
                      _ = arg_6
                      var try_7 EID 
                      /*g_try(v2:"try_7",loop:false) */
                      try_7 = EVAL(self.Arg)
                      /* ERROR PROTECTION INSERTED (arg_6-try_5) */
                      if ErrorIn(try_7) {try_5 = try_7
                      } else {
                      arg_6 = ANY(try_7)
                      try_5 = EID{Core.F__I_equal_any(arg_6,CFALSE.Id()).Id(),0}
                      }
                      } 
                    /* ERROR PROTECTION INSERTED (g0166I-loop_4) */
                    if ErrorIn(try_5) {loop_4 = try_5
                    } else {
                    g0166I = ToBoolean(OBJ(try_5))
                    if (g0166I == CTRUE) { 
                      res = res.AddFast(z)/*t=any,s=set*/
                      loop_4 = EID{res.Id(),0}
                      } else {
                      loop_4 = EID{CFALSE.Id(),0}
                      } 
                    }
                    /* ERROR PROTECTION INSERTED (loop_4-loop_4) */
                    if ErrorIn(loop_4) {loop_3 = loop_4
                    break
                    } else {
                    }}
                    }
                    } 
                  } 
                /* ERROR PROTECTION INSERTED (loop_3-Result) */
                if ErrorIn(loop_3) {Result = loop_3
                break
                } else {
                }
                } 
              } 
            } 
          }  else if (x.Isa.IsIn(C_Interval) == CTRUE) { 
          { var g0163 *ClaireInterval   = To_Interval(x)
            { var y int  = g0163.Arg1
              { var g0164 int  = g0163.Arg2
                _ = g0164
                Result= EID{CFALSE.Id(),0}
                for (y <= g0164) { 
                  /* While stat, v:"Result" loop:true */
                  var loop_8 EID 
                  _ = loop_8
                  { 
                  /*g_try(v2:"loop_8",loop:tuple("Result", EID)) */
                  /*g_try(v2:"loop_8",loop:tuple("Result", EID)) */
                  loop_8 = self.ClaireVar.WriteEID(EID{C__INT,IVAL(y)})
                  /* ERROR PROTECTION INSERTED (loop_8-loop_8) */
                  if ErrorIn(loop_8) {Result = loop_8
                  break
                  } else {
                  /*g_try(v2:"loop_8",loop:tuple("Result", EID)) */
                  var g0167I *ClaireBoolean  
                  var try_9 EID 
                  /*g_try(v2:"try_9",loop:false) */
                  { var arg_10 *ClaireAny  
                    _ = arg_10
                    var try_11 EID 
                    /*g_try(v2:"try_11",loop:false) */
                    try_11 = EVAL(self.Arg)
                    /* ERROR PROTECTION INSERTED (arg_10-try_9) */
                    if ErrorIn(try_11) {try_9 = try_11
                    } else {
                    arg_10 = ANY(try_11)
                    try_9 = EID{Core.F__I_equal_any(arg_10,CFALSE.Id()).Id(),0}
                    }
                    } 
                  /* ERROR PROTECTION INSERTED (g0167I-loop_8) */
                  if ErrorIn(try_9) {loop_8 = try_9
                  } else {
                  g0167I = ToBoolean(OBJ(try_9))
                  if (g0167I == CTRUE) { 
                    res = res.AddFast(MakeInteger(y).Id())/*t=any,s=set*/
                    loop_8 = EID{res.Id(),0}
                    } else {
                    loop_8 = EID{CFALSE.Id(),0}
                    } 
                  }
                  /* ERROR PROTECTION INSERTED (loop_8-loop_8) */
                  if ErrorIn(loop_8) {Result = loop_8
                  break
                  } else {
                  }}
                  /* ERROR PROTECTION INSERTED (loop_8-loop_8) */
                  if ErrorIn(loop_8) {Result = loop_8
                  break
                  } else {
                  y = (y+1)
                  }
                  /* try?:false, v2:"v_while8" loop will be:tuple("Result", EID) */
                  } 
                }
                } 
              } 
            } 
          } else {
          { 
            var y *ClaireAny  
            _ = y
            Result= EID{CFALSE.Id(),0}
            var y_support *ClaireList  
            var try_12 EID 
            /*g_try(v2:"try_12",loop:false) */
            try_12 = Core.F_enumerate_any(x)
            /* ERROR PROTECTION INSERTED (y_support-Result) */
            if ErrorIn(try_12) {Result = try_12
            } else {
            y_support = ToList(OBJ(try_12))
            y_len := y_support.Length()
            for i_it := 0; i_it < y_len; i_it++ { 
              y = y_support.At(i_it)
              var loop_13 EID 
              _ = loop_13
              { 
              /*g_try(v2:"loop_13",loop:tuple("Result", EID)) */
              loop_13 = F_write_value_Variable(self.ClaireVar,y)
              /* ERROR PROTECTION INSERTED (loop_13-loop_13) */
              if ErrorIn(loop_13) {Result = loop_13
              break
              } else {
              /*g_try(v2:"loop_13",loop:tuple("Result", EID)) */
              var g0168I *ClaireBoolean  
              var try_14 EID 
              /*g_try(v2:"try_14",loop:false) */
              { var arg_15 *ClaireAny  
                _ = arg_15
                var try_16 EID 
                /*g_try(v2:"try_16",loop:false) */
                try_16 = EVAL(self.Arg)
                /* ERROR PROTECTION INSERTED (arg_15-try_14) */
                if ErrorIn(try_16) {try_14 = try_16
                } else {
                arg_15 = ANY(try_16)
                try_14 = EID{Core.F__I_equal_any(arg_15,CFALSE.Id()).Id(),0}
                }
                } 
              /* ERROR PROTECTION INSERTED (g0168I-loop_13) */
              if ErrorIn(try_14) {loop_13 = try_14
              } else {
              g0168I = ToBoolean(OBJ(try_14))
              if (g0168I == CTRUE) { 
                res = res.AddFast(y)/*t=any,s=set*/
                loop_13 = EID{res.Id(),0}
                } else {
                loop_13 = EID{CFALSE.Id(),0}
                } 
              }
              /* ERROR PROTECTION INSERTED (loop_13-loop_13) */
              if ErrorIn(loop_13) {Result = loop_13
              break
              } else {
              }}
              }}
              } 
            } 
          } 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        Result = EID{res.Id(),0}
        }
        } 
      }
      } 
    return Result} 
  
// The EID go function for: self_eval @ Select (throw: true) 
func E_self_eval_Select (self EID) EID { 
    return To_Select(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Select 
func EVAL_Select (x *ClaireAny) EID { 
     return To_Select(x).SelfEval()} 
  
// [select VAR in SET_EXPR, ...] is the same as a "for" but returns the subset of
//  members that produce a true value
//
/* {1} The go function for: self_print(self:Lselect) [status=1] */
func (self *Lselect ) SelfPrint () EID { 
    var Result EID 
    PRINC("list{ ")
    /*g_try(v2:"Result",loop:true) */
    Result = F_ppvariable_Variable(self.ClaireVar)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(" in ")
    /*g_try(v2:"Result",loop:true) */
    { var _Zl int  = Core.C_pretty.Index
      _ = _Zl
      F_set_level_void()
      /*g_try(v2:"Result",loop:true) */
      Result = F_printexp_any(self.SetArg,CFALSE)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      { 
        var va_arg1 *Core.PrettyPrinter  
        var va_arg2 int 
        va_arg1 = Core.C_pretty
        va_arg2 = _Zl
        va_arg1.Index = va_arg2
        /*integer->integer*/Result = EID{C__INT,IVAL(va_arg2)}
        } 
      }
      } 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(" | ")
    /*g_try(v2:"Result",loop:true) */
    /*g_try(v2:"Result",loop:true) */
    Result = F_lbreak_integer(2)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    /*g_try(v2:"Result",loop:true) */
    Result = Core.F_CALL(C_print,ARGS(self.Arg.ToEID()))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    { 
      var va_arg1 *Core.PrettyPrinter  
      var va_arg2 int 
      va_arg1 = Core.C_pretty
      va_arg2 = (Core.C_pretty.Index-2)
      va_arg1.Index = va_arg2
      /*integer->integer*/Result = EID{C__INT,IVAL(va_arg2)}
      } 
    }}
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("}")
    Result = EVOID
    }}}
    return Result} 
  
// The EID go function for: self_print @ Lselect (throw: true) 
func E_self_print_Lselect_Language (self EID) EID { 
    return To_Lselect(OBJ(self)).SelfPrint( )} 
  
/* {1} The go function for: self_eval(self:Lselect) [status=1] */
func (self *Lselect ) SelfEval () EID { 
    var Result EID 
    { var x *ClaireAny  
      var try_1 EID 
      /*g_try(v2:"try_1",loop:false) */
      try_1 = EVAL(self.SetArg)
      /* ERROR PROTECTION INSERTED (x-Result) */
      if ErrorIn(try_1) {Result = try_1
      } else {
      x = ANY(try_1)
      { var res *ClaireList  
        if (x.Isa.IsIn(C_list) == CTRUE) { 
          { var g0170 *ClaireList   = ToList(x)
            _ = g0170
            res = g0170.Empty()
            } 
          } else {
          res = ToType(CEMPTY.Id()).EmptyList()
          } 
        /*g_try(v2:"Result",loop:true) */
        if (C_class.Id() == x.Isa.Id()) { 
          { var g0172 *ClaireClass   = ToClass(x)
            _ = g0172
            { 
              var y *ClaireClass  
              _ = y
              var y_iter *ClaireAny  
              Result= EID{CFALSE.Id(),0}
              var y_support *ClaireSet  
              y_support = g0172.Descendents
              for i_it := 0; i_it < y_support.Count; i_it++ { 
                y_iter = y_support.At(i_it)
                y = ToClass(y_iter)
                var loop_2 EID 
                _ = loop_2
                /*g_try(v2:"loop_2",loop:tuple("Result", EID)) */
                { 
                  var z *ClaireAny  
                  _ = z
                  loop_2= EID{CFALSE.Id(),0}
                  var z_support *ClaireList  
                  z_support = y.Instances
                  z_len := z_support.Length()
                  for i_it := 0; i_it < z_len; i_it++ { 
                    z = z_support.At(i_it)
                    var loop_3 EID 
                    _ = loop_3
                    { 
                    /*g_try(v2:"loop_3",loop:tuple("loop_2", EID)) */
                    loop_3 = F_write_value_Variable(self.ClaireVar,z)
                    /* ERROR PROTECTION INSERTED (loop_3-loop_3) */
                    if ErrorIn(loop_3) {loop_2 = loop_3
                    break
                    } else {
                    /*g_try(v2:"loop_3",loop:tuple("loop_2", EID)) */
                    var g0174I *ClaireBoolean  
                    var try_4 EID 
                    /*g_try(v2:"try_4",loop:false) */
                    { var arg_5 *ClaireAny  
                      _ = arg_5
                      var try_6 EID 
                      /*g_try(v2:"try_6",loop:false) */
                      try_6 = EVAL(self.Arg)
                      /* ERROR PROTECTION INSERTED (arg_5-try_4) */
                      if ErrorIn(try_6) {try_4 = try_6
                      } else {
                      arg_5 = ANY(try_6)
                      try_4 = EID{Core.F__I_equal_any(arg_5,CFALSE.Id()).Id(),0}
                      }
                      } 
                    /* ERROR PROTECTION INSERTED (g0174I-loop_3) */
                    if ErrorIn(try_4) {loop_3 = try_4
                    } else {
                    g0174I = ToBoolean(OBJ(try_4))
                    if (g0174I == CTRUE) { 
                      res = res.AddFast(z)/*t=any,s=list*/
                      loop_3 = EID{res.Id(),0}
                      } else {
                      loop_3 = EID{CFALSE.Id(),0}
                      } 
                    }
                    /* ERROR PROTECTION INSERTED (loop_3-loop_3) */
                    if ErrorIn(loop_3) {loop_2 = loop_3
                    break
                    } else {
                    }}
                    }
                    } 
                  } 
                /* ERROR PROTECTION INSERTED (loop_2-Result) */
                if ErrorIn(loop_2) {Result = loop_2
                break
                } else {
                }
                } 
              } 
            } 
          } else {
          { 
            var y *ClaireAny  
            _ = y
            Result= EID{CFALSE.Id(),0}
            var y_support *ClaireList  
            var try_7 EID 
            /*g_try(v2:"try_7",loop:false) */
            try_7 = Core.F_enumerate_any(x)
            /* ERROR PROTECTION INSERTED (y_support-Result) */
            if ErrorIn(try_7) {Result = try_7
            } else {
            y_support = ToList(OBJ(try_7))
            y_len := y_support.Length()
            for i_it := 0; i_it < y_len; i_it++ { 
              y = y_support.At(i_it)
              var loop_8 EID 
              _ = loop_8
              { 
              /*g_try(v2:"loop_8",loop:tuple("Result", EID)) */
              loop_8 = F_write_value_Variable(self.ClaireVar,y)
              /* ERROR PROTECTION INSERTED (loop_8-loop_8) */
              if ErrorIn(loop_8) {Result = loop_8
              break
              } else {
              /*g_try(v2:"loop_8",loop:tuple("Result", EID)) */
              var g0175I *ClaireBoolean  
              var try_9 EID 
              /*g_try(v2:"try_9",loop:false) */
              { var arg_10 *ClaireAny  
                _ = arg_10
                var try_11 EID 
                /*g_try(v2:"try_11",loop:false) */
                try_11 = EVAL(self.Arg)
                /* ERROR PROTECTION INSERTED (arg_10-try_9) */
                if ErrorIn(try_11) {try_9 = try_11
                } else {
                arg_10 = ANY(try_11)
                try_9 = EID{Core.F__I_equal_any(arg_10,CFALSE.Id()).Id(),0}
                }
                } 
              /* ERROR PROTECTION INSERTED (g0175I-loop_8) */
              if ErrorIn(try_9) {loop_8 = try_9
              } else {
              g0175I = ToBoolean(OBJ(try_9))
              if (g0175I == CTRUE) { 
                res = res.AddFast(y)/*t=any,s=list*/
                loop_8 = EID{res.Id(),0}
                } else {
                loop_8 = EID{CFALSE.Id(),0}
                } 
              }
              /* ERROR PROTECTION INSERTED (loop_8-loop_8) */
              if ErrorIn(loop_8) {Result = loop_8
              break
              } else {
              }}
              }}
              } 
            } 
          } 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        /*g_try(v2:"Result",loop:true) */
        if (self.Of.Id() != CNULL) { 
          /*g_try(v2:"Result",loop:true) */
          { var x *ClaireAny  
            { var x_some *ClaireAny   = CNULL
              _ = x_some
              { 
                var x *ClaireAny  
                _ = x
                var x_support *ClaireList  
                x_support = res
                x_len := x_support.Length()
                for i_it := 0; i_it < x_len; i_it++ { 
                  x = x_support.At(i_it)
                  if (self.Of.Contains(x) != CTRUE) { 
                    x_some = x
                    break
                    } 
                  } 
                } 
              x = x_some
              } 
            if (x != CNULL) { 
              { var _CL_obj *Core.RangeError   = Core.ToRangeError(new(Core.RangeError).Is(Core.C_range_error))
                _CL_obj.Cause = self.Id()
                /*any->any*/_CL_obj.Arg = x
                /*any->any*//*g_try(v2:"Result",loop:true) */
                Result = Core.F_write_property(C_Language_wrong,ToObject(_CL_obj.Id()),self.Of.Id())
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                Result = _CL_obj.Close()
                }
                } 
              } else {
              Result = EID{CNULL,0}
              } 
            } 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          Result = EID{res.Cast_I(self.Of).Id(),0}
          }
          } else {
          Result = EID{CFALSE.Id(),0}
          } 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        Result = EID{res.Id(),0}
        }}
        } 
      }
      } 
    return Result} 
  
// The EID go function for: self_eval @ Lselect (throw: true) 
func E_self_eval_Lselect (self EID) EID { 
    return To_Lselect(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Lselect 
func EVAL_Lselect (x *ClaireAny) EID { 
     return To_Lselect(x).SelfEval()} 
  
// Exists is an iteration that checks a condition
// other = true => forall,  other = false => exists, other = unknown => some
/* {1} The go function for: self_print(self:Exists) [status=1] */
func (self *Exists ) SelfPrint () EID { 
    var Result EID 
    if (self.Other == CTRUE.Id()) { 
      PRINC("forall")
      }  else if (self.Other == CFALSE.Id()) { 
      PRINC("exists")
      } else {
      PRINC("some")
      } 
    if (self.SetArg == C_any.Id()) { 
      PRINC("(")
      /*g_try(v2:"Result",loop:true) */
      Result = F_ppvariable_Variable(self.ClaireVar)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(",")
      /*g_try(v2:"Result",loop:true) */
      Result = Core.F_CALL(C_print,ARGS(self.Arg.ToEID()))
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(")")
      Result = EVOID
      }}
      } else {
      PRINC("(")
      /*g_try(v2:"Result",loop:true) */
      Result = F_ppvariable_Variable(self.ClaireVar)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(" in ")
      /*g_try(v2:"Result",loop:true) */
      { var _Zl int  = Core.C_pretty.Index
        _ = _Zl
        F_set_level_void()
        /*g_try(v2:"Result",loop:true) */
        Result = F_printexp_any(self.SetArg,CFALSE)
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        { 
          var va_arg1 *Core.PrettyPrinter  
          var va_arg2 int 
          va_arg1 = Core.C_pretty
          va_arg2 = _Zl
          va_arg1.Index = va_arg2
          /*integer->integer*/Result = EID{C__INT,IVAL(va_arg2)}
          } 
        }
        } 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(" | ")
      /*g_try(v2:"Result",loop:true) */
      /*g_try(v2:"Result",loop:true) */
      Result = F_lbreak_integer(2)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      /*g_try(v2:"Result",loop:true) */
      Result = Core.F_CALL(C_print,ARGS(self.Arg.ToEID()))
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      { 
        var va_arg1 *Core.PrettyPrinter  
        var va_arg2 int 
        va_arg1 = Core.C_pretty
        va_arg2 = (Core.C_pretty.Index-2)
        va_arg1.Index = va_arg2
        /*integer->integer*/Result = EID{C__INT,IVAL(va_arg2)}
        } 
      }}
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(")")
      Result = EVOID
      }}}
      } 
    return Result} 
  
// The EID go function for: self_print @ Exists (throw: true) 
func E_self_print_Exists_Language (self EID) EID { 
    return To_Exists(OBJ(self)).SelfPrint( )} 
  
/* {1} The go function for: self_eval(self:Exists) [status=1] */
func (self *Exists ) SelfEval () EID { 
    var Result EID 
    { var x *ClaireAny  
      var try_1 EID 
      /*g_try(v2:"try_1",loop:false) */
      try_1 = EVAL(self.SetArg)
      /* ERROR PROTECTION INSERTED (x-Result) */
      if ErrorIn(try_1) {Result = try_1
      } else {
      x = ANY(try_1)
      { var b *ClaireAny   = self.Other
        { var res *ClaireAny   = b
          _ = res
          /*g_try(v2:"Result",loop:true) */
          if (C_class.Id() == x.Isa.Id()) { 
            { var g0177 *ClaireClass   = ToClass(x)
              _ = g0177
              { 
                var y *ClaireClass  
                _ = y
                var y_iter *ClaireAny  
                Result= EID{CFALSE.Id(),0}
                var y_support *ClaireSet  
                y_support = g0177.Descendents
                for i_it := 0; i_it < y_support.Count; i_it++ { 
                  y_iter = y_support.At(i_it)
                  y = ToClass(y_iter)
                  var loop_2 EID 
                  _ = loop_2
                  /*g_try(v2:"loop_2",loop:tuple("Result", EID)) */
                  { 
                    var z *ClaireAny  
                    _ = z
                    loop_2= EID{CFALSE.Id(),0}
                    var z_support *ClaireList  
                    z_support = y.Instances
                    z_len := z_support.Length()
                    for i_it := 0; i_it < z_len; i_it++ { 
                      z = z_support.At(i_it)
                      var loop_3 EID 
                      _ = loop_3
                      { 
                      /*g_try(v2:"loop_3",loop:tuple("loop_2", EID)) */
                      loop_3 = F_write_value_Variable(self.ClaireVar,z)
                      /* ERROR PROTECTION INSERTED (loop_3-loop_3) */
                      if ErrorIn(loop_3) {loop_2 = loop_3
                      break
                      } else {
                      /*g_try(v2:"loop_3",loop:tuple("loop_2", EID)) */
                      var g0179I *ClaireBoolean  
                      var try_4 EID 
                      /*g_try(v2:"try_4",loop:false) */
                      { var arg_5 *ClaireAny  
                        _ = arg_5
                        var try_6 EID 
                        /*g_try(v2:"try_6",loop:false) */
                        try_6 = EVAL(self.Arg)
                        /* ERROR PROTECTION INSERTED (arg_5-try_4) */
                        if ErrorIn(try_6) {try_4 = try_6
                        } else {
                        arg_5 = ANY(try_6)
                        try_4 = EID{Core.F__I_equal_any(arg_5,CFALSE.Id()).Id(),0}
                        }
                        } 
                      /* ERROR PROTECTION INSERTED (g0179I-loop_3) */
                      if ErrorIn(try_4) {loop_3 = try_4
                      } else {
                      g0179I = ToBoolean(OBJ(try_4))
                      if (g0179I == CTRUE) { 
                        if (b != CTRUE.Id()) { 
                          res = IfThenElse((F_boolean_I_any(b) == CTRUE),
                            z,
                            CTRUE.Id())
                          loop_2 = res.ToEID()
                          break
                          } else {
                          loop_3 = EID{CFALSE.Id(),0}
                          } 
                        }  else if (b == CTRUE.Id()) { 
                        res = CFALSE.Id()
                        loop_2 = res.ToEID()
                        break
                        } else {
                        loop_3 = EID{CFALSE.Id(),0}
                        } 
                      }
                      /* ERROR PROTECTION INSERTED (loop_3-loop_3) */
                      if ErrorIn(loop_3) {loop_2 = loop_3
                      break
                      } else {
                      }}
                      }
                      } 
                    } 
                  /* ERROR PROTECTION INSERTED (loop_2-Result) */
                  if ErrorIn(loop_2) {Result = loop_2
                  break
                  } else {
                  }
                  } 
                } 
              } 
            } else {
            { 
              var y *ClaireAny  
              _ = y
              Result= EID{CFALSE.Id(),0}
              var y_support *ClaireList  
              var try_7 EID 
              /*g_try(v2:"try_7",loop:false) */
              try_7 = Core.F_enumerate_any(x)
              /* ERROR PROTECTION INSERTED (y_support-Result) */
              if ErrorIn(try_7) {Result = try_7
              } else {
              y_support = ToList(OBJ(try_7))
              y_len := y_support.Length()
              for i_it := 0; i_it < y_len; i_it++ { 
                y = y_support.At(i_it)
                var loop_8 EID 
                _ = loop_8
                { 
                /*g_try(v2:"loop_8",loop:tuple("Result", EID)) */
                loop_8 = F_write_value_Variable(self.ClaireVar,y)
                /* ERROR PROTECTION INSERTED (loop_8-loop_8) */
                if ErrorIn(loop_8) {Result = loop_8
                break
                } else {
                /*g_try(v2:"loop_8",loop:tuple("Result", EID)) */
                var g0180I *ClaireBoolean  
                var try_9 EID 
                /*g_try(v2:"try_9",loop:false) */
                { var arg_10 *ClaireAny  
                  _ = arg_10
                  var try_11 EID 
                  /*g_try(v2:"try_11",loop:false) */
                  try_11 = EVAL(self.Arg)
                  /* ERROR PROTECTION INSERTED (arg_10-try_9) */
                  if ErrorIn(try_11) {try_9 = try_11
                  } else {
                  arg_10 = ANY(try_11)
                  try_9 = EID{Core.F__I_equal_any(arg_10,CFALSE.Id()).Id(),0}
                  }
                  } 
                /* ERROR PROTECTION INSERTED (g0180I-loop_8) */
                if ErrorIn(try_9) {loop_8 = try_9
                } else {
                g0180I = ToBoolean(OBJ(try_9))
                if (g0180I == CTRUE) { 
                  if (b != CTRUE.Id()) { 
                    res = IfThenElse((F_boolean_I_any(b) == CTRUE),
                      y,
                      CTRUE.Id())
                    Result = res.ToEID()
                    break
                    } else {
                    loop_8 = EID{CFALSE.Id(),0}
                    } 
                  }  else if (b == CTRUE.Id()) { 
                  res = CFALSE.Id()
                  Result = res.ToEID()
                  break
                  } else {
                  loop_8 = EID{CFALSE.Id(),0}
                  } 
                }
                /* ERROR PROTECTION INSERTED (loop_8-loop_8) */
                if ErrorIn(loop_8) {Result = loop_8
                break
                } else {
                }}
                }}
                } 
              } 
            } 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          Result = res.ToEID()
          }
          } 
        } 
      }
      } 
    return Result} 
  
// The EID go function for: self_eval @ Exists (throw: true) 
func E_self_eval_Exists (self EID) EID { 
    return To_Exists(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Exists 
func EVAL_Exists (x *ClaireAny) EID { 
     return To_Exists(x).SelfEval()} 
  
// *********************************************************************
// *     Part 3: other control structures                              *
// *********************************************************************
// ----------------- case  --------------------------------------
/* {1} The go function for: self_print(self:Case) [status=1] */
func (self *Case ) SelfPrint () EID { 
    var Result EID 
    /*g_try(v2:"Result",loop:true) */
    PRINC("case ")
    /*g_try(v2:"Result",loop:true) */
    Result = Core.F_CALL(C_print,ARGS(self.ClaireVar.ToEID()))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(" ")
    /*g_try(v2:"Result",loop:true) */
    Result = F_lbreak_integer(1)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("(")
    Result = EVOID
    }}
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    { var n int  = 1
      { var m int  = self.Args.Length()
        Core.C_pretty.Index = (Core.C_pretty.Index+1)
        /*integer->integer*//*g_try(v2:"Result",loop:true) */
        Result= EID{CFALSE.Id(),0}
        for (n <= m) { 
          /* While stat, v:"Result" loop:true */
          var loop_1 EID 
          _ = loop_1
          /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
          { var _Zl int  = Core.C_pretty.Index
            _ = _Zl
            /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
            /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
            loop_1 = F_printexp_any(self.Args.At(n-1),CFALSE)
            /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
            if ErrorIn(loop_1) {Result = loop_1
            break
            } else {
            PRINC(" ")
            /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
            /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
            if (Core.F_buffer_length_void() > (Core.C_pretty.Width-50)) { 
              loop_1 = F_lbreak_integer(2)
              } else {
              F_set_level_void()
              loop_1 = EVOID
              } 
            /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
            if ErrorIn(loop_1) {Result = loop_1
            break
            } else {
            /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
            loop_1 = Core.F_CALL(C_print,ARGS(self.Args.At((n+1)-1).ToEID()))
            /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
            if ErrorIn(loop_1) {Result = loop_1
            break
            } else {
            }}
            /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
            if ErrorIn(loop_1) {Result = loop_1
            break
            } else {
            /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
            Core.C_pretty.Index = _Zl
            /*integer->integer*//*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
            if ((n+1) != m) { 
              PRINC(", ")
              /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
              loop_1 = F_lbreak_void()
              /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
              if ErrorIn(loop_1) {Result = loop_1
              break
              } else {
              PRINC("")
              loop_1 = EVOID
              }
              } else {
              loop_1 = EID{CFALSE.Id(),0}
              } 
            /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
            if ErrorIn(loop_1) {Result = loop_1
            break
            } else {
            }
            /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
            if ErrorIn(loop_1) {Result = loop_1
            break
            } else {
            PRINC("")
            loop_1 = EVOID
            }}}
            /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
            if ErrorIn(loop_1) {Result = loop_1
            break
            } else {
            n = (n+2)
            loop_1 = EID{C__INT,IVAL(n)}
            }
            } 
          /* ERROR PROTECTION INSERTED (loop_1-Result) */
          if ErrorIn(loop_1) {Result = loop_1
          break
          } else {
          /* try?:false, v2:"v_while4" loop will be:tuple("Result", EID) */
          } 
        }
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC(")")
        { 
          var va_arg1 *Core.PrettyPrinter  
          var va_arg2 int 
          va_arg1 = Core.C_pretty
          va_arg2 = (Core.C_pretty.Index-2)
          va_arg1.Index = va_arg2
          /*integer->integer*/Result = EID{C__INT,IVAL(va_arg2)}
          } 
        }
        } 
      } 
    }
    return Result} 
  
// The EID go function for: self_print @ Case (throw: true) 
func E_self_print_Case_Language (self EID) EID { 
    return To_Case(OBJ(self)).SelfPrint( )} 
  
/* {1} The go function for: self_eval(self:Case) [status=1] */
func (self *Case ) SelfEval () EID { 
    var Result EID 
    { var truc *ClaireAny  
      _ = truc
      var try_1 EID 
      /*g_try(v2:"try_1",loop:false) */
      try_1 = EVAL(self.ClaireVar)
      /* ERROR PROTECTION INSERTED (truc-Result) */
      if ErrorIn(try_1) {Result = try_1
      } else {
      truc = ANY(try_1)
      { var flip *ClaireBoolean   = CTRUE
        _ = flip
        { var previous *ClaireAny   = CFALSE.Id()
          var g0183I *ClaireBoolean  
          var try_2 EID 
          /*g_try(v2:"try_2",loop:false) */
          { 
            var x *ClaireAny  
            _ = x
            try_2= EID{CFALSE.Id(),0}
            var x_support *ClaireList  
            x_support = self.Args
            x_len := x_support.Length()
            for i_it := 0; i_it < x_len; i_it++ { 
              x = x_support.At(i_it)
              var loop_3 EID 
              _ = loop_3
              /*g_try(v2:"loop_3",loop:tuple("try_2", EID)) */
              if (flip == CTRUE) { 
                flip = CFALSE
                var try_4 EID 
                /*g_try(v2:"try_4",loop:tuple("try_2", EID)) */
                try_4 = EVAL(x)
                /* ERROR PROTECTION INSERTED (previous-loop_3) */
                if ErrorIn(try_4) {loop_3 = try_4
                try_2 = try_4
                break
                } else {
                previous = ANY(try_4)
                loop_3 = previous.ToEID()
                }
                } else {
                var g0184I *ClaireBoolean  
                var try_5 EID 
                /*g_try(v2:"try_5",loop:false) */
                try_5 = Core.F_BELONG(truc,previous)
                /* ERROR PROTECTION INSERTED (g0184I-loop_3) */
                if ErrorIn(try_5) {loop_3 = try_5
                } else {
                g0184I = ToBoolean(OBJ(try_5))
                if (g0184I == CTRUE) { 
                  var try_6 EID 
                  /*g_try(v2:"try_6",loop:tuple("try_2", EID)) */
                  try_6 = EVAL(x)
                  /* ERROR PROTECTION INSERTED (previous-loop_3) */
                  if ErrorIn(try_6) {loop_3 = try_6
                  try_2 = try_6
                  break
                  } else {
                  previous = ANY(try_6)
                  loop_3 = previous.ToEID()
                  try_2 = EID{CTRUE.Id(),0}
                  break
                  }
                  } else {
                  flip = CTRUE
                  loop_3 = EID{flip.Id(),0}
                  } 
                }
                } 
              /* ERROR PROTECTION INSERTED (loop_3-try_2) */
              if ErrorIn(loop_3) {try_2 = loop_3
              break
              } else {
              }
              } 
            } 
          /* ERROR PROTECTION INSERTED (g0183I-Result) */
          if ErrorIn(try_2) {Result = try_2
          } else {
          g0183I = ToBoolean(OBJ(try_2))
          if (g0183I == CTRUE) { 
            Result = previous.ToEID()
            } else {
            Result = EID{CFALSE.Id(),0}
            } 
          }
          } 
        } 
      }
      } 
    return Result} 
  
// The EID go function for: self_eval @ Case (throw: true) 
func E_self_eval_Case (self EID) EID { 
    return To_Case(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Case 
func EVAL_Case (x *ClaireAny) EID { 
     return To_Case(x).SelfEval()} 
  
// ------------------ WHILE  and UNTIL  -----------------------------
// the "other" while is until, where the first test is skipped
/* {1} The go function for: self_print(self:While) [status=1] */
func (self *While ) SelfPrint () EID { 
    var Result EID 
    /*g_try(v2:"Result",loop:true) */
    F_princ_string(ToString(IfThenElse((self.Other == CTRUE),
      MakeString("until").Id(),
      MakeString("while").Id())))
    PRINC(" ")
    /*g_try(v2:"Result",loop:true) */
    Result = F_printexp_any(self.Test,CFALSE)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(" ")
    /*g_try(v2:"Result",loop:true) */
    Result = F_lbreak_integer(2)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    /*g_try(v2:"Result",loop:true) */
    Result = Core.F_CALL(C_print,ARGS(self.Arg.ToEID()))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("")
    Result = EVOID
    }}}
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    { 
      var va_arg1 *Core.PrettyPrinter  
      var va_arg2 int 
      va_arg1 = Core.C_pretty
      va_arg2 = (Core.C_pretty.Index-2)
      va_arg1.Index = va_arg2
      /*integer->integer*/Result = EID{C__INT,IVAL(va_arg2)}
      } 
    }
    return Result} 
  
// The EID go function for: self_print @ While (throw: true) 
func E_self_print_While_Language (self EID) EID { 
    return To_While(OBJ(self)).SelfPrint( )} 
  
// other = true => self means  repeat self.arg until self.test = true
/* {1} The go function for: self_eval(self:While) [status=1] */
func (self *While ) SelfEval () EID { 
    var Result EID 
    { var a *ClaireBoolean   = self.Other
      { var b *ClaireBoolean   = a
        _ = b
        { 
          h_index := ClEnv.Index
          h_base := ClEnv.Base
          var v_while5 *ClaireBoolean  
          Result= EID{CFALSE.Id(),0}
          var try_1 EID 
          /*g_try(v2:"try_1",loop:false) */
          { 
            /* Or stat: v="try_1", loop=false */
            var v_or5 *ClaireBoolean  
            
            /* Or stat: try b with try:false, v="try_1", loop=false */
            v_or5 = b
            if (v_or5 == CTRUE) {try_1 = EID{CTRUE.Id(),0}
            } else { 
              /* Or stat: try = @ any(not @ any(eval @ list<type_expression>(any)(iClaire/test @ While(self))),a) with try:true, v="try_1", loop=false */
              var try_2 EID 
              /*g_try(v2:"try_2",loop:false) */
              { var arg_3 *ClaireBoolean  
                _ = arg_3
                var try_4 EID 
                /*g_try(v2:"try_4",loop:false) */
                { var arg_5 *ClaireAny  
                  _ = arg_5
                  var try_6 EID 
                  /*g_try(v2:"try_6",loop:false) */
                  try_6 = EVAL(self.Test)
                  /* ERROR PROTECTION INSERTED (arg_5-try_4) */
                  if ErrorIn(try_6) {try_4 = try_6
                  } else {
                  arg_5 = ANY(try_6)
                  try_4 = EID{Core.F_not_any(arg_5).Id(),0}
                  }
                  } 
                /* ERROR PROTECTION INSERTED (arg_3-try_2) */
                if ErrorIn(try_4) {try_2 = try_4
                } else {
                arg_3 = ToBoolean(OBJ(try_4))
                try_2 = EID{Equal(arg_3.Id(),a.Id()).Id(),0}
                }
                } 
              /* ERROR PROTECTION INSERTED (v_or5-try_1) */
              if ErrorIn(try_2) {try_1 = try_2
              } else {
              v_or5 = ToBoolean(OBJ(try_2))
              if (v_or5 == CTRUE) {try_1 = EID{CTRUE.Id(),0}
              } else { 
                try_1 = EID{CFALSE.Id(),0}} 
              } 
            }
            } 
          /* ERROR PROTECTION INSERTED (v_while5-Result) */
          if ErrorIn(try_1) {Result = try_1
          } else {
          v_while5 = ToBoolean(OBJ(try_1))
          
          for v_while5 == CTRUE { 
            /* While stat, v:"Result" loop:false */
            var loop_7 EID 
            _ = loop_7
            { 
            b = CFALSE
            /*g_try(v2:"loop_7",loop:tuple("Result", EID)) */
            loop_7 = EVAL(self.Arg)
            /* ERROR PROTECTION INSERTED (loop_7-loop_7) */
            if ErrorIn(loop_7) {Result = loop_7
            break
            } else {
            }
            /* try?:true, v2:"v_while5" loop will be:tuple("Result", EID) */
            var try_8 EID 
            /*g_try(v2:"try_8",loop:tuple("Result", EID)) */
            { 
              /* Or stat: v="try_8", loop=tuple("Result", EID) */
              var v_or6 *ClaireBoolean  
              
              /* Or stat: try b with try:false, v="try_8", loop=tuple("Result", EID) */
              v_or6 = b
              if (v_or6 == CTRUE) {try_8 = EID{CTRUE.Id(),0}
              } else { 
                /* Or stat: try = @ any(not @ any(eval @ list<type_expression>(any)(iClaire/test @ While(self))),a) with try:true, v="try_8", loop=tuple("Result", EID) */
                var try_9 EID 
                /*g_try(v2:"try_9",loop:tuple("Result", EID)) */
                { var arg_10 *ClaireBoolean  
                  _ = arg_10
                  var try_11 EID 
                  /*g_try(v2:"try_11",loop:false) */
                  { var arg_12 *ClaireAny  
                    _ = arg_12
                    var try_13 EID 
                    /*g_try(v2:"try_13",loop:false) */
                    try_13 = EVAL(self.Test)
                    /* ERROR PROTECTION INSERTED (arg_12-try_11) */
                    if ErrorIn(try_13) {try_11 = try_13
                    } else {
                    arg_12 = ANY(try_13)
                    try_11 = EID{Core.F_not_any(arg_12).Id(),0}
                    }
                    } 
                  /* ERROR PROTECTION INSERTED (arg_10-try_9) */
                  if ErrorIn(try_11) {try_9 = try_11
                  } else {
                  arg_10 = ToBoolean(OBJ(try_11))
                  try_9 = EID{Equal(arg_10.Id(),a.Id()).Id(),0}
                  }
                  } 
                /* ERROR PROTECTION INSERTED (v_or6-try_8) */
                if ErrorIn(try_9) {try_8 = try_9
                Result = try_9
                break
                } else {
                v_or6 = ToBoolean(OBJ(try_9))
                if (v_or6 == CTRUE) {try_8 = EID{CTRUE.Id(),0}
                } else { 
                  try_8 = EID{CFALSE.Id(),0}} 
                } 
              }
              } 
            /* ERROR PROTECTION INSERTED (v_while5-Result) */
            if ErrorIn(try_8) {Result = try_8
            break
            } else {
            v_while5 = ToBoolean(OBJ(try_8))
            } 
          }}
          }
          if ErrorIn(Result) && ToType(Core.C_return_error.Id()).Contains(ANY(Result)) == CTRUE { 
            ClEnv.Index = h_index
            ClEnv.Base = h_base
            Result = Core.F_CALL(C_arg,ARGS(EID{ClEnv.Exception_I.Id(),0}))
            } 
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: self_eval @ While (throw: true) 
func E_self_eval_While (self EID) EID { 
    return To_While(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: While 
func EVAL_While (x *ClaireAny) EID { 
     return To_While(x).SelfEval()} 
  
//-------------- handling errors -----------------------------------
// This is the control structure associated with these errors. Its real
// semantics is defined in the C compiler file
//
/* {1} The go function for: self_print(self:Handle) [status=1] */
func (self *ClaireHandle ) SelfPrint () EID { 
    var Result EID 
    /*g_try(v2:"Result",loop:true) */
    PRINC("try ")
    /*g_try(v2:"Result",loop:true) */
    Result = Core.F_CALL(C_print,ARGS(self.Arg.ToEID()))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(" ")
    /*g_try(v2:"Result",loop:true) */
    Result = F_lbreak_integer(0)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("catch ")
    /*g_try(v2:"Result",loop:true) */
    Result = Core.F_CALL(C_print,ARGS(self.Test.ToEID()))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(" ")
    /*g_try(v2:"Result",loop:true) */
    Result = Core.F_CALL(C_print,ARGS(self.Other.ToEID()))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("")
    Result = EVOID
    }}}}
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    { 
      var va_arg1 *Core.PrettyPrinter  
      var va_arg2 int 
      va_arg1 = Core.C_pretty
      va_arg2 = (Core.C_pretty.Index-2)
      va_arg1.Index = va_arg2
      /*integer->integer*/Result = EID{C__INT,IVAL(va_arg2)}
      } 
    }
    return Result} 
  
// The EID go function for: self_print @ Handle (throw: true) 
func E_self_print_Handle_Language (self EID) EID { 
    return To_ClaireHandle(OBJ(self)).SelfPrint( )} 
  
// original code
// self_eval(self:Handle) : any
//  -> (let x := (self.test as class) in
//       try eval(self.arg)
//       catch x (if (exception!() % return_error) close(exception!())
//                else eval(self.other)))     // <yc> 6/98
// CLAIRE 4 VERSION, because catch x => x is a constant class
// notice that return_error should be called return_exception since they travel through intepreted
// not a problem at compile time since return_exceptions are handled with break(x)
/* {1} The go function for: self_eval(self:Handle) [status=1] */
func (self *ClaireHandle ) SelfEval () EID { 
    var Result EID 
    { 
      h_index := ClEnv.Index
      h_base := ClEnv.Base
      Result = EVAL(self.Arg)
      if ErrorIn(Result){ 
        ClEnv.Index = h_index
        ClEnv.Base = h_base
        { var e *ClaireException   = ClEnv.Exception_I
          { var x *ClaireClass   = ToClass(self.Test)
            _ = x
            if ((e.Isa.IsIn(Core.C_return_error) == CTRUE) || 
                (Core.F__Z_any1(e.Id(),x) != CTRUE)) { 
              Result = e.Close()
              } else {
              Result = EVAL(self.Other)
              } 
            } 
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: self_eval @ Handle (throw: true) 
func E_self_eval_Handle (self EID) EID { 
    return To_ClaireHandle(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Handle 
func EVAL_Handle (x *ClaireAny) EID { 
     return To_ClaireHandle(x).SelfEval()} 
  
// <yc> 6/98
// *********************************************************************
// *     Part 4: the constructs                                         *
// *********************************************************************
// v3.2.16   constructor for arrays
/* {1} The go function for: self_print(self:Construct) [status=1] */
func (self *Construct ) SelfPrint () EID { 
    var Result EID 
    { var _Zl int  = Core.C_pretty.Index
      _ = _Zl
      /*g_try(v2:"Result",loop:true) */
      { var arg_1 *ClaireString  
        _ = arg_1
        if (self.Isa.IsIn(C_List) == CTRUE) { 
          arg_1 = MakeString("list")
          }  else if (self.Isa.IsIn(C_Set) == CTRUE) { 
          arg_1 = MakeString("set")
          }  else if (self.Isa.IsIn(C_Tuple) == CTRUE) { 
          arg_1 = MakeString("tuple")
          }  else if (self.Isa.IsIn(C_Printf) == CTRUE) { 
          arg_1 = MakeString("printf")
          }  else if (self.Isa.IsIn(C_Error) == CTRUE) { 
          arg_1 = MakeString("error")
          }  else if (self.Isa.IsIn(C_Trace) == CTRUE) { 
          arg_1 = MakeString("trace")
          }  else if (self.Isa.IsIn(C_Assert) == CTRUE) { 
          arg_1 = MakeString("assert")
          }  else if (self.Isa.IsIn(C_Branch) == CTRUE) { 
          arg_1 = MakeString("branch")
          }  else if (self.Isa.IsIn(C_Map) == CTRUE) { 
          arg_1 = MakeString("map")
          } else {
          arg_1 = self.Isa.Name.String_I()
          } 
        F_princ_string(arg_1)
        } 
      /*g_try(v2:"Result",loop:true) */
      if ((self.Isa.IsIn(C_List) == CTRUE) || 
          (self.Isa.IsIn(C_Set) == CTRUE)) { 
        { var g0197 *Construct   = self
          _ = g0197
          { var _Zt *ClaireAny   = Core.F_get_property(C_of,ToObject(g0197.Id()))
            if (_Zt != CNULL) { 
              if (Equal(_Zt,CEMPTY.Id()) != CTRUE) { 
                PRINC("<")
                /*g_try(v2:"Result",loop:true) */
                Result = Core.F_CALL(C_print,ARGS(_Zt.ToEID()))
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                PRINC(">")
                Result = EVOID
                }
                } else {
                Result = EID{CFALSE.Id(),0}
                } 
              } else {
              Result = EID{CNULL,0}
              } 
            } 
          } 
        }  else if (self.Isa.IsIn(C_Map) == CTRUE) { 
        { var g0198 *Map   = To_Map(self.Id())
          PRINC("<")
          /*g_try(v2:"Result",loop:true) */
          Result = Core.F_print_any(g0198.Domain.Id())
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(",")
          /*g_try(v2:"Result",loop:true) */
          Result = Core.F_print_any(g0198.Of.Id())
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(">")
          Result = EVOID
          }}
          } 
        } else {
        Result = EID{CFALSE.Id(),0}
        } 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC("(")
      F_set_level_void()
      /*g_try(v2:"Result",loop:true) */
      Result = F_Language_printbox_list2(self.Args)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(")")
      Result = EVOID
      }}
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      { 
        var va_arg1 *Core.PrettyPrinter  
        var va_arg2 int 
        va_arg1 = Core.C_pretty
        va_arg2 = _Zl
        va_arg1.Index = va_arg2
        /*integer->integer*/Result = EID{C__INT,IVAL(va_arg2)}
        } 
      }
      } 
    return Result} 
  
// The EID go function for: self_print @ Construct (throw: true) 
func E_self_print_Construct_Language (self EID) EID { 
    return To_Construct(OBJ(self)).SelfPrint( )} 
  
// constructors: how to create a list, a set, a tuple or an array
// note that the constructor is typed
// CLAIRE4: must build the list with the proper type from the begining, so that Srange is correct
/* {1} The go function for: self_eval(self:List) [status=1] */
func (self *List ) SelfEval () EID { 
    var Result EID 
    { var type_ask *ClaireBoolean   = MakeBoolean((self.Of.Id() == CNULL)).Not
      _ = type_ask
      { var n int  = self.Args.Length()
        if (type_ask == CTRUE) { 
          { var l *ClaireList   = CreateList(self.Of,n)
            /*g_try(v2:"Result",loop:true) */
            { var i int  = 1
              { var g0199 int  = n
                _ = g0199
                Result= EID{CFALSE.Id(),0}
                for (i <= g0199) { 
                  /* While stat, v:"Result" loop:true */
                  var loop_1 EID 
                  _ = loop_1
                  { 
                  /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
                  { 
                    var arg_2 EID 
                    /*g_try(v2:"arg_2",loop:false) */
                    arg_2 = EVAL(self.Args.At(i-1))
                    /* ERROR PROTECTION INSERTED (arg_2-loop_1) */
                    if ErrorIn(arg_2) {loop_1 = arg_2
                    } else {
                    loop_1 = l.WriteEID(i,arg_2)}
                    } 
                  /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
                  if ErrorIn(loop_1) {Result = loop_1
                  break
                  } else {
                  i = (i+1)
                  }
                  /* try?:false, v2:"v_while8" loop will be:tuple("Result", EID) */
                  } 
                }
                } 
              } 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            Result = EID{l.Id(),0}
            }
            } 
          } else {
          { var l *ClaireList   = CreateList(ToType(CEMPTY.Id()),n)
            /*g_try(v2:"Result",loop:true) */
            { var i int  = 1
              { var g0200 int  = n
                _ = g0200
                Result= EID{CFALSE.Id(),0}
                for (i <= g0200) { 
                  /* While stat, v:"Result" loop:true */
                  var loop_3 EID 
                  _ = loop_3
                  { 
                  /*g_try(v2:"loop_3",loop:tuple("Result", EID)) */
                  { var arg_4 *ClaireAny  
                    _ = arg_4
                    var try_5 EID 
                    /*g_try(v2:"try_5",loop:false) */
                    try_5 = EVAL(self.Args.At(i-1))
                    /* ERROR PROTECTION INSERTED (arg_4-loop_3) */
                    if ErrorIn(try_5) {loop_3 = try_5
                    } else {
                    arg_4 = ANY(try_5)
                    loop_3 = ToArray(l.Id()).NthPut(i,arg_4).ToEID()
                    }
                    } 
                  /* ERROR PROTECTION INSERTED (loop_3-loop_3) */
                  if ErrorIn(loop_3) {Result = loop_3
                  break
                  } else {
                  i = (i+1)
                  }
                  /* try?:false, v2:"v_while8" loop will be:tuple("Result", EID) */
                  } 
                }
                } 
              } 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            Result = EID{l.Id(),0}
            }
            } 
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: self_eval @ List (throw: true) 
func E_self_eval_List (self EID) EID { 
    return To_List(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: List 
func EVAL_List (x *ClaireAny) EID { 
     return To_List(x).SelfEval()} 
  
// here we use the CLAIRE 3 style of post-typing with a cast! 
/* {1} The go function for: self_eval(self:Set) [status=1] */
func (self *Set ) SelfEval () EID { 
    var Result EID 
    { var type_ask *ClaireBoolean   = MakeBoolean((self.Of.Id() == CNULL)).Not
      _ = type_ask
      { var n int  = self.Args.Length()
        if (type_ask == CTRUE) { 
          { var l *ClaireSet   = self.Of.EmptySet()
            /*g_try(v2:"Result",loop:true) */
            { var i int  = 1
              { var g0201 int  = n
                _ = g0201
                Result= EID{CFALSE.Id(),0}
                for (i <= g0201) { 
                  /* While stat, v:"Result" loop:true */
                  var loop_1 EID 
                  _ = loop_1
                  { 
                  /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
                  { var arg_2 *ClaireAny  
                    _ = arg_2
                    var try_3 EID 
                    /*g_try(v2:"try_3",loop:false) */
                    try_3 = EVAL(self.Args.At(i-1))
                    /* ERROR PROTECTION INSERTED (arg_2-loop_1) */
                    if ErrorIn(try_3) {loop_1 = try_3
                    } else {
                    arg_2 = ANY(try_3)
                    loop_1 = EID{l.AddFast(arg_2).Id(),0}/*t=any,s=EID*/
                    }
                    } 
                  /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
                  if ErrorIn(loop_1) {Result = loop_1
                  break
                  } else {
                  i = (i+1)
                  }
                  /* try?:false, v2:"v_while8" loop will be:tuple("Result", EID) */
                  } 
                }
                } 
              } 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            Result = EID{l.Id(),0}
            }
            } 
          } else {
          { var l *ClaireSet   = CEMPTY.EmptySet()
            /*g_try(v2:"Result",loop:true) */
            { var i int  = 1
              { var g0202 int  = n
                _ = g0202
                Result= EID{CFALSE.Id(),0}
                for (i <= g0202) { 
                  /* While stat, v:"Result" loop:true */
                  var loop_4 EID 
                  _ = loop_4
                  { 
                  /*g_try(v2:"loop_4",loop:tuple("Result", EID)) */
                  { var arg_5 *ClaireAny  
                    _ = arg_5
                    var try_6 EID 
                    /*g_try(v2:"try_6",loop:false) */
                    try_6 = EVAL(self.Args.At(i-1))
                    /* ERROR PROTECTION INSERTED (arg_5-loop_4) */
                    if ErrorIn(try_6) {loop_4 = try_6
                    } else {
                    arg_5 = ANY(try_6)
                    loop_4 = EID{l.AddFast(arg_5).Id(),0}/*t=any,s=EID*/
                    }
                    } 
                  /* ERROR PROTECTION INSERTED (loop_4-loop_4) */
                  if ErrorIn(loop_4) {Result = loop_4
                  break
                  } else {
                  i = (i+1)
                  }
                  /* try?:false, v2:"v_while8" loop will be:tuple("Result", EID) */
                  } 
                }
                } 
              } 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            Result = EID{l.Id(),0}
            }
            } 
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: self_eval @ Set (throw: true) 
func E_self_eval_Set (self EID) EID { 
    return To_Set(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Set 
func EVAL_Set (x *ClaireAny) EID { 
     return To_Set(x).SelfEval()} 
  
//
/* {1} The go function for: self_eval(self:Tuple) [status=1] */
func (self *Tuple ) SelfEval () EID { 
    var Result EID 
    { var arg_1 *ClaireList  
      _ = arg_1
      var try_2 EID 
      /*g_try(v2:"try_2",loop:false) */
      { 
        var v_list3 *ClaireList  
        var x *ClaireAny  
        var v_local3 *ClaireAny  
        v_list3 = self.Args
        try_2 = EID{CreateList(ToType(CEMPTY.Id()),v_list3.Length()).Id(),0}
        for CLcount := 0; CLcount < v_list3.Length(); CLcount++{ 
          x = v_list3.At(CLcount)
          var try_3 EID 
          /*g_try(v2:"try_3",loop:tuple("try_2", EID)) */
          try_3 = EVAL(x)
          /* ERROR PROTECTION INSERTED (v_local3-try_2) */
          if ErrorIn(try_3) {try_2 = try_3
          break
          } else {
          v_local3 = ANY(try_3)
          ToList(OBJ(try_2)).PutAt(CLcount,v_local3)
          } 
        }
        } 
      /* ERROR PROTECTION INSERTED (arg_1-Result) */
      if ErrorIn(try_2) {Result = try_2
      } else {
      arg_1 = ToList(OBJ(try_2))
      Result = EID{arg_1.Tuple_I().Id(),0}
      }
      } 
    return Result} 
  
// The EID go function for: self_eval @ Tuple (throw: true) 
func E_self_eval_Tuple (self EID) EID { 
    return To_Tuple(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Tuple 
func EVAL_Tuple (x *ClaireAny) EID { 
     return To_Tuple(x).SelfEval()} 
  
// same as creating a list (same constraints since same underlying structure)
/* {1} The go function for: self_eval(self:Array) [status=1] */
func (self *Array ) SelfEval () EID { 
    var Result EID 
    { var type_ask *ClaireBoolean   = MakeBoolean((self.Of.Id() == CNULL)).Not
      _ = type_ask
      { var n int  = self.Args.Length()
        if (type_ask == CTRUE) { 
          { var l *ClaireList   = CreateList(self.Of,n)
            /*g_try(v2:"Result",loop:true) */
            { var i int  = 1
              { var g0203 int  = n
                _ = g0203
                Result= EID{CFALSE.Id(),0}
                for (i <= g0203) { 
                  /* While stat, v:"Result" loop:true */
                  var loop_1 EID 
                  _ = loop_1
                  { 
                  /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
                  { 
                    var arg_2 EID 
                    /*g_try(v2:"arg_2",loop:false) */
                    arg_2 = EVAL(self.Args.At(i-1))
                    /* ERROR PROTECTION INSERTED (arg_2-loop_1) */
                    if ErrorIn(arg_2) {loop_1 = arg_2
                    } else {
                    loop_1 = l.WriteEID(i,arg_2)}
                    } 
                  /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
                  if ErrorIn(loop_1) {Result = loop_1
                  break
                  } else {
                  i = (i+1)
                  }
                  /* try?:false, v2:"v_while8" loop will be:tuple("Result", EID) */
                  } 
                }
                } 
              } 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            Result = EID{l.Array_I().Id(),0}
            }
            } 
          } else {
          { var l *ClaireList   = CreateList(ToType(CEMPTY.Id()),n)
            /*g_try(v2:"Result",loop:true) */
            { var i int  = 1
              { var g0204 int  = n
                _ = g0204
                Result= EID{CFALSE.Id(),0}
                for (i <= g0204) { 
                  /* While stat, v:"Result" loop:true */
                  var loop_3 EID 
                  _ = loop_3
                  { 
                  /*g_try(v2:"loop_3",loop:tuple("Result", EID)) */
                  { var arg_4 *ClaireAny  
                    _ = arg_4
                    var try_5 EID 
                    /*g_try(v2:"try_5",loop:false) */
                    try_5 = EVAL(self.Args.At(i-1))
                    /* ERROR PROTECTION INSERTED (arg_4-loop_3) */
                    if ErrorIn(try_5) {loop_3 = try_5
                    } else {
                    arg_4 = ANY(try_5)
                    loop_3 = ToArray(l.Id()).NthPut(i,arg_4).ToEID()
                    }
                    } 
                  /* ERROR PROTECTION INSERTED (loop_3-loop_3) */
                  if ErrorIn(loop_3) {Result = loop_3
                  break
                  } else {
                  i = (i+1)
                  }
                  /* try?:false, v2:"v_while8" loop will be:tuple("Result", EID) */
                  } 
                }
                } 
              } 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            Result = EID{l.Array_I().Id(),0}
            }
            } 
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: self_eval @ Array (throw: true) 
func E_self_eval_Array2 (self EID) EID { 
    return To_Array(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Array 
func EVAL_Array (x *ClaireAny) EID { 
     return To_Array(x).SelfEval()} 
  
// create a map from a list of pairs
/* {1} The go function for: self_eval(self:Map) [status=1] */
func (self *Map ) SelfEval () EID { 
    var Result EID 
    { var m *ClaireMapSet   = self.Domain.Map_I(self.Of)
      Core.F_tformat_string(MakeString("self_eval(~S) \n"),0,MakeConstantList(m.Id()))
      /*g_try(v2:"Result",loop:true) */
      { 
        var x *ClaireAny  
        _ = x
        Result= EID{CFALSE.Id(),0}
        var x_support *ClaireList  
        x_support = self.Args
        x_len := x_support.Length()
        for i_it := 0; i_it < x_len; i_it++ { 
          x = x_support.At(i_it)
          var loop_1 EID 
          _ = loop_1
          /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
          if (x.Isa.IsIn(C_pair) == CTRUE) { 
            { var g0205 *ClairePair   = ToPair(x)
              { var arg_2 *ClaireAny  
                _ = arg_2
                var try_4 EID 
                /*g_try(v2:"try_4",loop:false) */
                try_4 = EVAL(g0205.First)
                /* ERROR PROTECTION INSERTED (arg_2-loop_1) */
                if ErrorIn(try_4) {loop_1 = try_4
                } else {
                arg_2 = ANY(try_4)
                { var arg_3 *ClaireAny  
                  _ = arg_3
                  var try_5 EID 
                  /*g_try(v2:"try_5",loop:false) */
                  try_5 = EVAL(g0205.Second)
                  /* ERROR PROTECTION INSERTED (arg_3-loop_1) */
                  if ErrorIn(try_5) {loop_1 = try_5
                  } else {
                  arg_3 = ANY(try_5)
                  loop_1 = m.Put(arg_2,arg_3)
                  }
                  } 
                }
                } 
              } 
            } else {
            loop_1 = ToException(Core.C_general_error.Make(MakeString("~S is not a pair, cannot be inserted in map ~S").Id(),MakeConstantList(x,m.Id()).Id())).Close()
            } 
          /* ERROR PROTECTION INSERTED (loop_1-Result) */
          if ErrorIn(loop_1) {Result = loop_1
          break
          } else {
          }
          } 
        } 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = EID{m.Id(),0}
      }
      } 
    return Result} 
  
// The EID go function for: self_eval @ Map (throw: true) 
func E_self_eval_Map (self EID) EID { 
    return To_Map(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Map 
func EVAL_Map (x *ClaireAny) EID { 
     return To_Map(x).SelfEval()} 
  
// Macros are a nice but undocumented feature of CLAIRE. This is deliberate :)
// it is an advanced feature for those who want to expand the language. This
// makes CLAIRE a nice framework for DSL
//
/* {1} The go function for: self_eval(self:Macro) [status=1] */
func (self *Macro ) SelfEval () EID { 
    var Result EID 
    { var arg_1 *ClaireAny  
      _ = arg_1
      var try_2 EID 
      /*g_try(v2:"try_2",loop:false) */
      try_2 = Core.F_CALL(C_macroexpand,ARGS(EID{self.Id(),0}))
      /* ERROR PROTECTION INSERTED (arg_1-Result) */
      if ErrorIn(try_2) {Result = try_2
      } else {
      arg_1 = ANY(try_2)
      Result = EVAL(arg_1)
      }
      } 
    return Result} 
  
// The EID go function for: self_eval @ Macro (throw: true) 
func E_self_eval_Macro2 (self EID) EID { 
    return To_Macro(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Macro 
func EVAL_Macro (x *ClaireAny) EID { 
     return To_Macro(x).SelfEval()} 
  
// error produces an exception of type general_error
/* {1} The go function for: self_eval(self:Error) [status=1] */
func (self *Error ) SelfEval () EID { 
    var Result EID 
    /*g_try(v2:"Result",loop:true) */
    if ((F_boolean_I_any(self.Args.Id()).Id() != CTRUE.Id()) || 
        (C_string.Id() != self.Args.At(1-1).Isa.Id())) { 
      Result = ToException(Core.C_general_error.Make(MakeString("Syntax error: ~S").Id(),MakeConstantList(self.Id()).Id())).Close()
      } else {
      Result = EID{CFALSE.Id(),0}
      } 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    { var x *Core.GeneralError   = Core.ToGeneralError(new(Core.GeneralError).Is(Core.C_general_error))
      /*g_try(v2:"Result",loop:true) */
      { 
        var va_arg1 *Core.GeneralError  
        var va_arg2 *ClaireAny  
        va_arg1 = x
        var try_1 EID 
        /*g_try(v2:"try_1",loop:false) */
        try_1 = Core.F_car_list(self.Args)
        /* ERROR PROTECTION INSERTED (va_arg2-Result) */
        if ErrorIn(try_1) {Result = try_1
        } else {
        va_arg2 = ANY(try_1)
        va_arg1.Cause = va_arg2
        /*any->any*/Result = va_arg2.ToEID()
        }
        } 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      /*g_try(v2:"Result",loop:true) */
      { 
        var va_arg1 *Core.GeneralError  
        var va_arg2 *ClaireAny  
        va_arg1 = x
        var try_2 EID 
        /*g_try(v2:"try_2",loop:false) */
        { 
          var v_list4 *ClaireList  
          var x *ClaireAny  
          var v_local4 *ClaireAny  
          var try_3 EID 
          /*g_try(v2:"try_3",loop:false) */
          try_3 = self.Args.Cdr()
          /* ERROR PROTECTION INSERTED (v_list4-try_2) */
          if ErrorIn(try_3) {try_2 = try_3
          } else {
          v_list4 = ToList(OBJ(try_3))
          try_2 = EID{CreateList(ToType(CEMPTY.Id()),v_list4.Length()).Id(),0}
          for CLcount := 0; CLcount < v_list4.Length(); CLcount++{ 
            x = v_list4.At(CLcount)
            var try_4 EID 
            /*g_try(v2:"try_4",loop:tuple("try_2", EID)) */
            try_4 = EVAL(x)
            /* ERROR PROTECTION INSERTED (v_local4-try_2) */
            if ErrorIn(try_4) {try_2 = try_4
            break
            } else {
            v_local4 = ANY(try_4)
            ToList(OBJ(try_2)).PutAt(CLcount,v_local4)
            } 
          }}
          } 
        /* ERROR PROTECTION INSERTED (va_arg2-Result) */
        if ErrorIn(try_2) {Result = try_2
        } else {
        va_arg2 = ANY(try_2)
        va_arg1.Arg = va_arg2
        /*any->any*/Result = va_arg2.ToEID()
        }
        } 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = x.Close()
      }}
      } 
    }
    return Result} 
  
// The EID go function for: self_eval @ Error (throw: true) 
func E_self_eval_Error (self EID) EID { 
    return To_Error(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Error 
func EVAL_Error (x *ClaireAny) EID { 
     return To_Error(x).SelfEval()} 
  
// this is the basic tool for printing in CLAIRE. A complex statement
// is macroexpanded into basic printing instructions
//
/* {1} The go function for: self_eval(self:Printf) [status=1] */
func (self *Printf ) SelfEval () EID { 
    var Result EID 
    { var l *ClaireList   = self.Args
      { var s *ClaireAny   = l.At(1-1)
        /*g_try(v2:"Result",loop:true) */
        if (C_string.Id() != s.Isa.Id()) { 
          Result = ToException(Core.C_general_error.Make(MakeString("[102] the first argument in ~S must be a string").Id(),MakeConstantList(self.Id()).Id())).Close()
          } else {
          { var i int  = 2
            { var n int  = F_get_string(ToString(s),'~')
              /*g_try(v2:"Result",loop:true) */
              Result= EID{CFALSE.Id(),0}
              for (n != 0) { 
                /* While stat, v:"Result" loop:true */
                var loop_1 EID 
                _ = loop_1
                /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
                { var m *ClaireAny  
                  var try_2 EID 
                  /*g_try(v2:"try_2",loop:false) */
                  try_2 = Core.F_CALL(C_nth,ARGS(s.ToEID(),EID{C__INT,IVAL((n+1))}))
                  /* ERROR PROTECTION INSERTED (m-loop_1) */
                  if ErrorIn(try_2) {loop_1 = try_2
                  } else {
                  m = ANY(try_2)
                  /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
                  if (i > l.Length()) { 
                    loop_1 = ToException(Core.C_general_error.Make(MakeString("[103] not enough arguments in ~S").Id(),MakeConstantList(self.Id()).Id())).Close()
                    } else {
                    loop_1 = EID{CFALSE.Id(),0}
                    } 
                  /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
                  if ErrorIn(loop_1) {Result = loop_1
                  break
                  } else {
                  if (n > 1) { 
                    F_princ_string(F_substring_string(ToString(s),1,(n-1)))
                    } 
                  /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
                  if ('A' == ToChar(m).Value) { 
                    { var arg_3 *ClaireAny  
                      _ = arg_3
                      var try_4 EID 
                      /*g_try(v2:"try_4",loop:false) */
                      try_4 = EVAL(l.At(i-1))
                      /* ERROR PROTECTION INSERTED (arg_3-loop_1) */
                      if ErrorIn(try_4) {loop_1 = try_4
                      } else {
                      arg_3 = ANY(try_4)
                      loop_1 = Core.F_CALL(C_princ,ARGS(arg_3.ToEID()))
                      }
                      } 
                    }  else if ('S' == ToChar(m).Value) { 
                    { var arg_5 *ClaireAny  
                      _ = arg_5
                      var try_6 EID 
                      /*g_try(v2:"try_6",loop:false) */
                      try_6 = EVAL(l.At(i-1))
                      /* ERROR PROTECTION INSERTED (arg_5-loop_1) */
                      if ErrorIn(try_6) {loop_1 = try_6
                      } else {
                      arg_5 = ANY(try_6)
                      loop_1 = Core.F_CALL(C_print,ARGS(arg_5.ToEID()))
                      }
                      } 
                    }  else if ('F' == ToChar(m).Value) { 
                    { var fv *ClaireAny  
                      _ = fv
                      var try_7 EID 
                      /*g_try(v2:"try_7",loop:false) */
                      try_7 = EVAL(l.At(i-1))
                      /* ERROR PROTECTION INSERTED (fv-loop_1) */
                      if ErrorIn(try_7) {loop_1 = try_7
                      } else {
                      fv = ANY(try_7)
                      { var p_Z *ClaireBoolean   = CFALSE
                        { var j int 
                          var try_8 EID 
                          /*g_try(v2:"try_8",loop:false) */
                          { var arg_9 int 
                            _ = arg_9
                            var try_10 EID 
                            /*g_try(v2:"try_10",loop:false) */
                            { var arg_11 rune 
                              _ = arg_11
                              var try_12 EID 
                              /*g_try(v2:"try_12",loop:false) */
                              try_12 = Core.F_nth_get_string(ToString(s),(n+2),(n+2))
                              /* ERROR PROTECTION INSERTED (arg_11-try_10) */
                              if ErrorIn(try_12) {try_10 = try_12
                              } else {
                              arg_11 = CHAR(try_12)
                              try_10 = EID{C__INT,IVAL(int(arg_11))}
                              }
                              } 
                            /* ERROR PROTECTION INSERTED (arg_9-try_8) */
                            if ErrorIn(try_10) {try_8 = try_10
                            } else {
                            arg_9 = INT(try_10)
                            try_8 = EID{C__INT,IVAL((arg_9-48))}
                            }
                            } 
                          /* ERROR PROTECTION INSERTED (j-loop_1) */
                          if ErrorIn(try_8) {loop_1 = try_8
                          } else {
                          j = INT(try_8)
                          /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
                          var g0207I *ClaireBoolean  
                          var try_13 EID 
                          /*g_try(v2:"try_13",loop:false) */
                          { var arg_14 *ClaireAny  
                            _ = arg_14
                            var try_15 EID 
                            /*g_try(v2:"try_15",loop:false) */
                            try_15 = Core.F_CALL(C_nth,ARGS(s.ToEID(),EID{C__INT,IVAL((n+2))}))
                            /* ERROR PROTECTION INSERTED (arg_14-try_13) */
                            if ErrorIn(try_15) {try_13 = try_15
                            } else {
                            arg_14 = ANY(try_15)
                            try_13 = EID{Equal(MakeChar('%').Id(),arg_14).Id(),0}
                            }
                            } 
                          /* ERROR PROTECTION INSERTED (g0207I-loop_1) */
                          if ErrorIn(try_13) {loop_1 = try_13
                          } else {
                          g0207I = ToBoolean(OBJ(try_13))
                          if (g0207I == CTRUE) { 
                            p_Z = CTRUE
                            j = 1
                            fv = ANY(Core.F_CALL(ToProperty(C__star.Id()),ARGS(fv.ToEID(),EID{C__FLOAT,FVAL(100)})))
                            loop_1 = fv.ToEID()
                            }  else if ((j < 0) || 
                              (j > 9)) { 
                            loop_1 = ToException(Core.C_general_error.Make(MakeString("[189] F requires a single digit integer in ~S").Id(),MakeConstantList(self.Id()).Id())).Close()
                            } else {
                            loop_1 = EID{CFALSE.Id(),0}
                            } 
                          }
                          /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
                          if ErrorIn(loop_1) {Result = loop_1
                          break
                          } else {
                          /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
                          var g0208I *ClaireBoolean  
                          var try_16 EID 
                          /*g_try(v2:"try_16",loop:false) */
                          { 
                            var v_and13 *ClaireBoolean  
                            
                            v_and13 = p_Z.Not
                            if (v_and13 == CFALSE) {try_16 = EID{CFALSE.Id(),0}
                            } else { 
                              var try_17 EID 
                              /*g_try(v2:"try_17",loop:false) */
                              { var arg_18 *ClaireAny  
                                _ = arg_18
                                var try_19 EID 
                                /*g_try(v2:"try_19",loop:false) */
                                try_19 = Core.F_CALL(C_nth,ARGS(s.ToEID(),EID{C__INT,IVAL((n+3))}))
                                /* ERROR PROTECTION INSERTED (arg_18-try_17) */
                                if ErrorIn(try_19) {try_17 = try_19
                                } else {
                                arg_18 = ANY(try_19)
                                try_17 = EID{Equal(MakeChar('%').Id(),arg_18).Id(),0}
                                }
                                } 
                              /* ERROR PROTECTION INSERTED (v_and13-try_16) */
                              if ErrorIn(try_17) {try_16 = try_17
                              } else {
                              v_and13 = ToBoolean(OBJ(try_17))
                              if (v_and13 == CFALSE) {try_16 = EID{CFALSE.Id(),0}
                              } else { 
                                try_16 = EID{CTRUE.Id(),0}} 
                              } 
                            }
                            } 
                          /* ERROR PROTECTION INSERTED (g0208I-loop_1) */
                          if ErrorIn(try_16) {loop_1 = try_16
                          } else {
                          g0208I = ToBoolean(OBJ(try_16))
                          if (g0208I == CTRUE) { 
                            p_Z = CTRUE
                            fv = ANY(Core.F_CALL(ToProperty(C__star.Id()),ARGS(fv.ToEID(),EID{C__FLOAT,FVAL(100)})))
                            n = (n+1)
                            loop_1 = EID{C__INT,IVAL(n)}
                            } else {
                            loop_1 = EID{CFALSE.Id(),0}
                            } 
                          }
                          /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
                          if ErrorIn(loop_1) {Result = loop_1
                          break
                          } else {
                          /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
                          loop_1 = Core.F_CALL(Core.C_mClaire_printFDigit,ARGS(fv.ToEID(),EID{C__INT,IVAL(j)}))
                          /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
                          if ErrorIn(loop_1) {Result = loop_1
                          break
                          } else {
                          if (p_Z == CTRUE) { 
                            PRINC("%")
                            } 
                          n = (n+1)
                          loop_1 = EID{C__INT,IVAL(n)}
                          }}}
                          }
                          } 
                        } 
                      }
                      } 
                    }  else if ('I' == ToChar(m).Value) { 
                    loop_1 = EVAL(l.At(i-1))
                    } else {
                    loop_1 = EID{CFALSE.Id(),0}
                    } 
                  /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
                  if ErrorIn(loop_1) {Result = loop_1
                  break
                  } else {
                  i = (i+1)
                  s = (F_substring_string(ToString(s),(n+2),1000)).Id()
                  var try_20 EID 
                  /*g_try(v2:"try_20",loop:tuple("Result", EID)) */
                  try_20 = Core.F_CALL(C_get,ARGS(s.ToEID(),EID{C__CHAR,CVAL('~')}))
                  /* ERROR PROTECTION INSERTED (n-loop_1) */
                  if ErrorIn(try_20) {loop_1 = try_20
                  Result = try_20
                  break
                  } else {
                  n = INT(try_20)
                  loop_1 = EID{C__INT,IVAL(n)}
                  }}}
                  }
                  } 
                /* ERROR PROTECTION INSERTED (loop_1-Result) */
                if ErrorIn(loop_1) {Result = loop_1
                break
                } else {
                /* try?:false, v2:"v_while7" loop will be:tuple("Result", EID) */
                } 
              }
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              if (F_boolean_I_any(s) == CTRUE) { 
                Result = Core.F_CALL(C_princ,ARGS(s.ToEID()))
                } else {
                Result = EID{CFALSE.Id(),0}
                } 
              }
              } 
            } 
          } 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        Result = EID{CNULL,0}
        }
        } 
      } 
    return Result} 
  
// The EID go function for: self_eval @ Printf (throw: true) 
func E_self_eval_Printf (self EID) EID { 
    return To_Printf(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Printf 
func EVAL_Printf (x *ClaireAny) EID { 
     return To_Printf(x).SelfEval()} 
  
// trace is refined in inspect.cl
// If trace_output() is known, use it, else use current output.
// defined in inspect.cl
// CLAIRE4: self_eval is defined once for all, hence exteneded
/* {1} The go function for: self_eval(self:Trace) [status=1] */
func (self *Trace ) SelfEval () EID { 
    var Result EID 
    { var a *ClaireList   = self.Args
      { var l *ClaireList  
        var try_1 EID 
        /*g_try(v2:"try_1",loop:false) */
        { 
          var v_list4 *ClaireList  
          var x *ClaireAny  
          var v_local4 *ClaireAny  
          v_list4 = a
          try_1 = EID{CreateList(ToType(CEMPTY.Id()),v_list4.Length()).Id(),0}
          for CLcount := 0; CLcount < v_list4.Length(); CLcount++{ 
            x = v_list4.At(CLcount)
            var try_2 EID 
            /*g_try(v2:"try_2",loop:tuple("try_1", EID)) */
            try_2 = EVAL(x)
            /* ERROR PROTECTION INSERTED (v_local4-try_1) */
            if ErrorIn(try_2) {try_1 = try_2
            break
            } else {
            v_local4 = ANY(try_2)
            ToList(OBJ(try_1)).PutAt(CLcount,v_local4)
            } 
          }
          } 
        /* ERROR PROTECTION INSERTED (l-Result) */
        if ErrorIn(try_1) {Result = try_1
        } else {
        l = ToList(OBJ(try_1))
        { var i *ClaireAny   = l.At(1-1)
          { var a2 *ClaireAny  
            if (a.Length() > 1) { 
              a2 = a.At(2-1)
              } else {
              a2 = CFALSE.Id()
              } 
            if (a.Length() == 1) { 
              { var a1 *ClaireAny  
                _ = a1
                var try_3 EID 
                /*g_try(v2:"try_3",loop:false) */
                try_3 = EVAL(a.At(1-1))
                /* ERROR PROTECTION INSERTED (a1-Result) */
                if ErrorIn(try_3) {Result = try_3
                } else {
                a1 = ANY(try_3)
                { var p *ClaireProperty   = C_iClaire_trace_on
                  if (p.Restrictions.Length() != 0) { 
                    if (ClEnv.Trace_I == 0) { 
                      ClEnv.Trace_I = 1
                      /*integer->integer*/} 
                    Result = Core.F_CALL(Core.C_call,ARGS(EID{p.Id(),0},a1.ToEID()))
                    } else {
                    Result = EID{CFALSE.Id(),0}
                    } 
                  } 
                }
                } 
              } else {
              var g0210I *ClaireBoolean  
              { 
                var v_and7 *ClaireBoolean  
                
                v_and7 = Equal(C_string.Id(),a2.Isa.Id())
                if (v_and7 == CFALSE) {g0210I = CFALSE
                } else { 
                  if (C_integer.Id() == i.Isa.Id()) { 
                    { var g0209 int  = ToInteger(i).Value
                      _ = g0209
                      v_and7 = Core.F__inf_equal_integer(g0209,ClEnv.Verbose)
                      } 
                    } else {
                    v_and7 = CFALSE
                    } 
                  if (v_and7 == CFALSE) {g0210I = CFALSE
                  } else { 
                    g0210I = CTRUE} 
                  } 
                } 
              if (g0210I == CTRUE) { 
                { var p *ClaireAny   = Core.F_get_property(C_ctrace,ToObject(ClEnv.Id()))
                  if (p != CNULL) { 
                    p = ToPort(p).UseAsOutput().Id()
                    } 
                  /*g_try(v2:"Result",loop:true) */
                  Result = Core.F_format_string(ToString(a2),l.Skip(2))
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  if (p != CNULL) { 
                    ToPort(p).UseAsOutput()
                    } 
                  Result = EID{CEMPTY.Id(),0}
                  }
                  } 
                } else {
                Result = EID{CFALSE.Id(),0}
                } 
              } 
            } 
          } 
        }
        } 
      } 
    return Result} 
  
// The EID go function for: self_eval @ Trace (throw: true) 
func E_self_eval_Trace (self EID) EID { 
    return To_Trace(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Trace 
func EVAL_Trace (x *ClaireAny) EID { 
     return To_Trace(x).SelfEval()} 
  
// assert is refined in trace.la
//
/* {1} The go function for: self_eval(self:Assert) [status=1] */
func (self *Assert ) SelfEval () EID { 
    var Result EID 
    { var a *ClaireList   = self.Args
      var g0211I *ClaireBoolean  
      var try_1 EID 
      /*g_try(v2:"try_1",loop:false) */
      { 
        var v_and3 *ClaireBoolean  
        
        v_and3 = Core.F__sup_integer(a.Length(),0)
        if (v_and3 == CFALSE) {try_1 = EID{CFALSE.Id(),0}
        } else { 
          v_and3 = Core.F_known_ask_any(Core.F_get_property(C_ctrace,ToObject(ClEnv.Id())))
          if (v_and3 == CFALSE) {try_1 = EID{CFALSE.Id(),0}
          } else { 
            var try_2 EID 
            /*g_try(v2:"try_2",loop:false) */
            { var arg_3 *ClaireBoolean  
              _ = arg_3
              var try_4 EID 
              /*g_try(v2:"try_4",loop:false) */
              { var arg_5 *ClaireAny  
                _ = arg_5
                var try_6 EID 
                /*g_try(v2:"try_6",loop:false) */
                try_6 = EVAL(a.At(1-1))
                /* ERROR PROTECTION INSERTED (arg_5-try_4) */
                if ErrorIn(try_6) {try_4 = try_6
                } else {
                arg_5 = ANY(try_6)
                try_4 = EID{F_boolean_I_any(arg_5).Id(),0}
                }
                } 
              /* ERROR PROTECTION INSERTED (arg_3-try_2) */
              if ErrorIn(try_4) {try_2 = try_4
              } else {
              arg_3 = ToBoolean(OBJ(try_4))
              try_2 = EID{Core.F__I_equal_any(arg_3.Id(),CTRUE.Id()).Id(),0}
              }
              } 
            /* ERROR PROTECTION INSERTED (v_and3-try_1) */
            if ErrorIn(try_2) {try_1 = try_2
            } else {
            v_and3 = ToBoolean(OBJ(try_2))
            if (v_and3 == CFALSE) {try_1 = EID{CFALSE.Id(),0}
            } else { 
              try_1 = EID{CTRUE.Id(),0}} 
            } 
          } 
        }
        } 
      /* ERROR PROTECTION INSERTED (g0211I-Result) */
      if ErrorIn(try_1) {Result = try_1
      } else {
      g0211I = ToBoolean(OBJ(try_1))
      if (g0211I == CTRUE) { 
        { var p *ClairePort   = ClEnv.Ctrace.UseAsOutput()
          _ = p
          /*g_try(v2:"Result",loop:true) */
          /*g_try(v2:"Result",loop:true) */
          Result = Core.F_print_any((self.External).Id())
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(",line=")
          F_princ_integer(self.Index)
          PRINC(": (ASSERT) ")
          /*g_try(v2:"Result",loop:true) */
          Result = Core.F_CALL(C_print,ARGS(a.At(1-1).ToEID()))
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC("\n")
          Result = EVOID
          }}
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          p.UseAsOutput()
          /*g_try(v2:"Result",loop:true) */
          if (ClEnv.Debug_I >= 0) { 
            Result = ToException(Core.C_general_error.Make(MakeString("Assertion Violation").Id(),CNIL.Id())).Close()
            } else {
            Result = EID{CFALSE.Id(),0}
            } 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          Result = EID{CEMPTY.Id(),0}
          }}
          } 
        } else {
        Result = EID{CFALSE.Id(),0}
        } 
      }
      } 
    return Result} 
  
// The EID go function for: self_eval @ Assert (throw: true) 
func E_self_eval_Assert (self EID) EID { 
    return To_Assert(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Assert 
func EVAL_Assert (x *ClaireAny) EID { 
     return To_Assert(x).SelfEval()} 
  
/* {1} The go function for: self_eval(self:Branch) [status=1] */
func (self *Branch ) SelfEval () EID { 
    var Result EID 
    /*g_try(v2:"Result",loop:true) */
    if (self.Args.Length() != 1) { 
      Result = ToException(Core.C_general_error.Make(MakeString("[104] Syntax error with ~S (one arg. expected)").Id(),MakeConstantList(self.Id()).Id())).Close()
      } else {
      Result = EID{CFALSE.Id(),0}
      } 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    { 
      h_index := ClEnv.Index
      h_base := ClEnv.Base
      F_world_push()
      var g0212I *ClaireBoolean  
      var try_1 EID 
      /*g_try(v2:"try_1",loop:false) */
      { var arg_2 *ClaireAny  
        _ = arg_2
        var try_3 EID 
        /*g_try(v2:"try_3",loop:false) */
        try_3 = EVAL(self.Args.At(1-1))
        /* ERROR PROTECTION INSERTED (arg_2-try_1) */
        if ErrorIn(try_3) {try_1 = try_3
        } else {
        arg_2 = ANY(try_3)
        try_1 = EID{Core.F__I_equal_any(arg_2,CFALSE.Id()).Id(),0}
        }
        } 
      /* ERROR PROTECTION INSERTED (g0212I-Result) */
      if ErrorIn(try_1) {Result = try_1
      } else {
      g0212I = ToBoolean(OBJ(try_1))
      if (g0212I == CTRUE) { 
        Result = EID{CTRUE.Id(),0}
        } else {
        F_world_pop()
        Result = EID{CFALSE.Id(),0}
        } 
      }
      if ErrorIn(Result) && ToType(Core.C_contradiction.Id()).Contains(ANY(Result)) == CTRUE { 
        ClEnv.Index = h_index
        ClEnv.Base = h_base
        F_world_pop()
        Result = EID{CFALSE.Id(),0}
        } 
      } 
    }
    return Result} 
  
// The EID go function for: self_eval @ Branch (throw: true) 
func E_self_eval_Branch (self EID) EID { 
    return To_Branch(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Branch 
func EVAL_Branch (x *ClaireAny) EID { 
     return To_Branch(x).SelfEval()} 
  
// end of file