/***** CLAIRE Compilation of file /Users/ycaseau/Dropbox/src/clairev4.12/src/meta/control.cl 
         [version 4.1.4 / safety 5] Wednesday 01-01-2025 17:30:04 *****/

package Language
import (_ "fmt"
	. "Kernel"
	"Core"
)

//-------- dumb function to prevent import errors --------
func import_g0128() { 
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
// v4.12 conditional comment
/* The go function for: self_print(self:If) [status=1] */
func (self *If) SelfPrint () EID { 
  var Result EID
  PRINC("(")
  Core.C_pretty.Index = (Core.C_pretty.Index+1)
  Result = self.Printstat()
  if !ErrorIn(Result) {
  Core.C_pretty.Index = (Core.C_pretty.Index-1)
  PRINC(")")
  Result = EVOID
  }
  return Result} 

// The EID go function for: self_print @ If (throw: true) 
func E_self_print_If_Language (self EID) EID { 
  return To_If(OBJ(self)).SelfPrint( )} 

/* The go function for: printstat(self:If) [status=1] */
func (self *If) Printstat () EID { 
  var Result EID
  PRINC("if ")
  Result = F_printexp_any(self.Test,CFALSE)
  if !ErrorIn(Result) {
  Result = F_checkfar_void()
  }
  if !ErrorIn(Result) {
  PRINC(" ")
  Result = F_printif_any(self.Arg)
  if !ErrorIn(Result) {
  Core.C_pretty.Index = (Core.C_pretty.Index-3)
  Result = self.Printelse()
  }}
  return Result} 

// The EID go function for: printstat @ If (throw: true) 
func E_printstat_If (self EID) EID { 
  return To_If(OBJ(self)).Printstat( )} 

/* The go function for: printif(self:any) [status=1] */
func F_printif_any (self *ClaireAny) EID { 
  var Result EID
  Core.C_pretty.Index = (Core.C_pretty.Index+3)
  if (Core.C_pretty.Pbreak == CTRUE) { 
    { var b_index int = Core.F_buffer_length_void()
      { var _Zl int = Core.C_pretty.Index
        Core.C_pretty.Pbreak = CFALSE
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
        if !ErrorIn(Result) {
        Core.C_pretty.Pbreak = CTRUE
        if (Core.F_short_enough_integer(Core.F_buffer_length_void()) == CTRUE) { 
          Result = EID{CEMPTY.Id(),0}
          } else {
          Core.F_buffer_set_length_integer(b_index)
          Core.C_pretty.Index = _Zl
          Result = F_lbreak_void()
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

/* The go function for: printelse(self:If) [status=1] */
func (self *If) Printelse () EID { 
  var Result EID
  { var e *ClaireAny = self.Other
    if (e.Isa.IsIn(C_If) == CTRUE) { 
      { var g0134 *If = To_If(e)
        PRINC(" ")
        Result = F_lbreak_void()
        if !ErrorIn(Result) {
        PRINC("else if ")
        Result = F_printexp_any(g0134.Test,CFALSE)
        if !ErrorIn(Result) {
        PRINC(" ")
        Result = F_printif_any(g0134.Arg)
        if !ErrorIn(Result) {
        Core.C_pretty.Index = (Core.C_pretty.Index-3)
        Result = g0134.Printelse()
        }}}
        } 
      }  else if (Equal(e,CNIL.Id()) != CTRUE) { 
      { var _Zl int = Core.C_pretty.Index
        PRINC(" ")
        Result = F_lbreak_void()
        if !ErrorIn(Result) {
        PRINC("else ")
        F_set_level_integer(1)
        Result = Core.F_CALL(C_print,ARGS(e.ToEID()))
        }
        if !ErrorIn(Result) {
        { 
          var va_arg1 *Core.PrettyPrinter
          var va_arg2 int
          va_arg1 = Core.C_pretty
          va_arg2 = _Zl
          va_arg1.Index = va_arg2
          Result = EID{C__INT,IVAL(va_arg2)}
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
/* The go function for: self_eval(self:If) [status=1] */
func (self *If) SelfEval () EID { 
  var Result EID
  { var x *ClaireAny
    var try_1 EID
    try_1 = EVAL(self.Test)
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
/* The go function for: self_print(self:Do) [status=1] */
func (self *Do) SelfPrint () EID { 
  var Result EID
  { var _Zl int = Core.C_pretty.Index
    PRINC("(")
    F_set_level_integer(1)
    Result = F_printdo_list(self.Args,CTRUE)
    if !ErrorIn(Result) {
    { 
      var va_arg1 *Core.PrettyPrinter
      var va_arg2 int
      va_arg1 = Core.C_pretty
      va_arg2 = _Zl
      va_arg1.Index = va_arg2
      Result = EID{C__INT,IVAL(va_arg2)}
      } 
    }
    } 
  return Result} 

// The EID go function for: self_print @ Do (throw: true) 
func E_self_print_Do_Language (self EID) EID { 
  return To_Do(OBJ(self)).SelfPrint( )} 

/* The go function for: printdo(l:list,clo:boolean) [status=1] */
func F_printdo_list (l *ClaireList,clo *ClaireBoolean) EID { 
  var Result EID
  { var n int = l.Length()
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
        if (x.Isa.IsIn(C_If) == CTRUE) { 
          { var g0136 *If = To_If(x)
            loop_1 = g0136.Printstat()
            } 
          } else {
          loop_1 = Core.F_CALL(C_print,ARGS(x.ToEID()))
          } 
        if ErrorIn(loop_1) {Result = loop_1
        break
        } else {
        n = (n-1)
        if (n == 0) { 
          if (clo == CTRUE) { 
            PRINC(")")
            loop_1 = EVOID
            } else {
            loop_1 = EID{CFALSE.Id(),0}
            } 
          } else {
          PRINC(", ")
          loop_1 = F_lbreak_void()
          if ErrorIn(loop_1) {Result = loop_1
          break
          } else {
          }
          } 
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

/* The go function for: printblock(x:any) [status=1] */
func F_printblock_any (x *ClaireAny) EID { 
  var Result EID
  if (x.Isa.IsIn(C_Do) == CTRUE) { 
    { var g0138 *Do = To_Do(x)
      Result = F_printdo_list(g0138.Args,CFALSE)
      } 
    }  else if (x.Isa.IsIn(C_If) == CTRUE) { 
    { var g0139 *If = To_If(x)
      Result = g0139.Printstat()
      } 
    } else {
    Result = Core.F_CALL(C_print,ARGS(x.ToEID()))
    } 
  return Result} 

// The EID go function for: printblock @ any (throw: true) 
func E_printblock_any (x EID) EID { 
  return F_printblock_any(ANY(x) )} 

// use res:EID pragma when compiled with CLAIRE4, res:any for CLAIRE3
/* The go function for: self_eval(self:Do) [status=1] */
func (self *Do) SelfEval () EID { 
  var Result EID
  { var res *ClaireAny = CEMPTY.Id()
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
        try_2 = EVAL(_Zx)
        if ErrorIn(try_2) {Result = try_2
        break
        } else {
        res = ANY(try_2)
        loop_1 = res.ToEID()
        }
        } 
      } 
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
/* The go function for: self_print(self:Let) [status=1] */
func (self *Let) SelfPrint () EID { 
  var Result EID
  { var _Zl int = Core.C_pretty.Index
    F_set_level_integer(1)
    PRINC("let ")
    Result = F_ppvariable_Variable(self.ClaireVar)
    if !ErrorIn(Result) {
    PRINC(" := ")
    Result = F_printexp_any(self.Value,CFALSE)
    if !ErrorIn(Result) {
    Result = self.Printbody()
    }}
    if !ErrorIn(Result) {
    { 
      var va_arg1 *Core.PrettyPrinter
      var va_arg2 int
      va_arg1 = Core.C_pretty
      va_arg2 = _Zl
      va_arg1.Index = va_arg2
      Result = EID{C__INT,IVAL(va_arg2)}
      } 
    }
    } 
  return Result} 

// The EID go function for: self_print @ Let (throw: true) 
func E_self_print_Let_Language (self EID) EID { 
  return To_Let(OBJ(self)).SelfPrint( )} 

/* The go function for: printbody(self:Let) [status=1] */
func (self *Let) Printbody () EID { 
  var Result EID
  { var a *ClaireAny = self.Arg
    if (a.Isa.IsIn(C_Let) == CTRUE) { 
      { var g0142 *Let = To_Let(a)
        PRINC(",")
        Result = F_lbreak_integer(4)
        if !ErrorIn(Result) {
        Result = F_ppvariable_Variable(g0142.ClaireVar)
        if !ErrorIn(Result) {
        PRINC(" := ")
        Result = F_printexp_any(g0142.Value,CFALSE)
        if !ErrorIn(Result) {
        Core.C_pretty.Index = (Core.C_pretty.Index-4)
        Result = g0142.Printbody()
        }}}
        } 
      } else {
      PRINC(" in ")
      Result = F_lbreak_integer(2)
      if !ErrorIn(Result) {
      Result = Core.F_CALL(C_print,ARGS(a.ToEID()))
      }
      } 
    } 
  return Result} 

// The EID go function for: printbody @ Let (throw: true) 
func E_printbody_Let (self EID) EID { 
  return To_Let(OBJ(self)).Printbody( )} 

/* The go function for: self_eval(self:Let) [status=1] */
func (self *Let) SelfEval () EID { 
  var Result EID
  { var val *ClaireAny
    var try_1 EID
    try_1 = EVAL(self.Value)
    if ErrorIn(try_1) {Result = try_1
    } else {
    val = ANY(try_1)
    Result = F_write_value_Variable(self.ClaireVar,val)
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
/* The go function for: self_print(self:When) [status=1] */
func (self *When) SelfPrint () EID { 
  var Result EID
  { var _Zl int = Core.C_pretty.Index
    F_set_level_integer(1)
    PRINC("when ")
    Result = F_ppvariable_Variable(self.ClaireVar)
    if !ErrorIn(Result) {
    PRINC(" := ")
    Result = F_printexp_any(self.Value,CFALSE)
    if !ErrorIn(Result) {
    PRINC(" in ")
    Result = F_lbreak_integer(2)
    if !ErrorIn(Result) {
    Result = Core.F_CALL(C_print,ARGS(self.Arg.ToEID()))
    }}}
    if !ErrorIn(Result) {
    if (self.Other != CNULL) { 
      PRINC(" ")
      Result = F_lbreak_void()
      if !ErrorIn(Result) {
      PRINC("else ")
      F_set_level_integer(1)
      Result = Core.F_CALL(C_print,ARGS(self.Other.ToEID()))
      }
      } else {
      Result = EID{CFALSE.Id(),0}
      } 
    if !ErrorIn(Result) {
    { 
      var va_arg1 *Core.PrettyPrinter
      var va_arg2 int
      va_arg1 = Core.C_pretty
      va_arg2 = _Zl
      va_arg1.Index = va_arg2
      Result = EID{C__INT,IVAL(va_arg2)}
      } 
    }}
    } 
  return Result} 

// The EID go function for: self_print @ When (throw: true) 
func E_self_print_When_Language (self EID) EID { 
  return To_When(OBJ(self)).SelfPrint( )} 

/* The go function for: self_eval(self:When) [status=1] */
func (self *When) SelfEval () EID { 
  var Result EID
  { var val *ClaireAny
    var try_1 EID
    try_1 = EVAL(self.Value)
    if ErrorIn(try_1) {Result = try_1
    } else {
    val = ANY(try_1)
    { var n int = ClEnv.Trace_I
      _ = n
      if (val != CNULL) { 
        Result = F_write_value_Variable(self.ClaireVar,val)
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
/* The go function for: self_print(self:Let+) [status=1] */
func (self *Let_plus) SelfPrint () EID { 
  var Result EID
  { var _Zl int = Core.C_pretty.Index
    { var l *ClaireList = To_Do(self.Arg).Args
      F_set_level_integer(1)
      PRINC("let ")
      Result = F_printexp_any(self.Value,CFALSE)
      if !ErrorIn(Result) {
      PRINC(" := ")
      Result = F_printexp_any(ToList(OBJ(Core.F_CALL(C_args,ARGS(l.At(0).ToEID())))).At(2),CFALSE)
      if !ErrorIn(Result) {
      PRINC(" in ")
      Result = F_lbreak_integer(2)
      if !ErrorIn(Result) {
      Result = Core.F_CALL(C_print,ARGS(To_Let(l.At(1)).Value.ToEID()))
      }}}
      if !ErrorIn(Result) {
      { 
        var va_arg1 *Core.PrettyPrinter
        var va_arg2 int
        va_arg1 = Core.C_pretty
        va_arg2 = _Zl
        va_arg1.Index = va_arg2
        Result = EID{C__INT,IVAL(va_arg2)}
        } 
      }
      } 
    } 
  return Result} 

// The EID go function for: self_print @ Let+ (throw: true) 
func E_self_print_Let_plus_Language (self EID) EID { 
  return To_Let_plus(OBJ(self)).SelfPrint( )} 

/* The go function for: self_print(self:Let*) [status=1] */
func (self *Let_star) SelfPrint () EID { 
  var Result EID
  { var _Zl int = Core.C_pretty.Index
    { var l *ClaireAny = self.Arg
      F_set_level_integer(1)
      if (l.Isa.IsIn(C_Let) == CTRUE) { 
        PRINC("let (")
        Result= EID{CFALSE.Id(),0}
        for true { 
          var loop_1 EID
          _ = loop_1
          { 
          loop_1 = Core.F_CALL(C_Language_ppvariable,ARGS(Core.F_CALL(C_var,ARGS(l.ToEID()))))
          if ErrorIn(loop_1) {Result = loop_1
          break
          } else {
          { var lnext *ClaireAny = ANY(Core.F_CALL(C_arg,ARGS(l.ToEID())))
            var g0145I *ClaireBoolean
            if (lnext.Isa.IsIn(C_Let) == CTRUE) { 
              { var g0144 *Let = To_Let(lnext)
                g0145I = MakeBoolean((g0144.Value.Isa.IsIn(C_Call) == CTRUE) && (ToList(OBJ(Core.F_CALL(C_args,ARGS(g0144.Value.ToEID())))).At(0) == self.ClaireVar.Id()))
                } 
              } else {
              g0145I = CFALSE
              } 
            if (g0145I == CTRUE) { 
              PRINC(",")
              l = lnext
              } else {
              Result = EID{CTRUE.Id(),0}
              break
              } 
            } 
          }
          } 
        }
        if !ErrorIn(Result) {
        PRINC(") := ")
        Result = F_printexp_any(self.Value,CFALSE)
        if !ErrorIn(Result) {
        Result = To_Let(l).Printbody()
        }}
        } else {
        PRINC("(")
        { var _Zf *ClaireBoolean = CTRUE
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
              loop_2 = Core.F_CALL(C_Language_ppvariable,ARGS(Core.F_CALL(C_var,ARGS(_Za.ToEID()))))
              if ErrorIn(loop_2) {Result = loop_2
              break
              } else {
              }
              }
              } 
            } 
          } 
        if !ErrorIn(Result) {
        PRINC(") := ")
        Result = F_printexp_any(self.Value,CFALSE)
        }
        } 
      if !ErrorIn(Result) {
      { 
        var va_arg1 *Core.PrettyPrinter
        var va_arg2 int
        va_arg1 = Core.C_pretty
        va_arg2 = _Zl
        va_arg1.Index = va_arg2
        Result = EID{C__INT,IVAL(va_arg2)}
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
/* The go function for: self_print(self:For) [status=1] */
func (self *For) SelfPrint () EID { 
  var Result EID
  PRINC("for ")
  Result = F_ppvariable_Variable(self.ClaireVar)
  if !ErrorIn(Result) {
  PRINC(" in ")
  { var _Zl int = Core.C_pretty.Index
    F_set_level_void()
    Result = F_printexp_any(self.SetArg,CFALSE)
    if !ErrorIn(Result) {
    { 
      var va_arg1 *Core.PrettyPrinter
      var va_arg2 int
      va_arg1 = Core.C_pretty
      va_arg2 = _Zl
      va_arg1.Index = va_arg2
      Result = EID{C__INT,IVAL(va_arg2)}
      } 
    }
    } 
  if !ErrorIn(Result) {
  PRINC(" ")
  Core.C_pretty.Index = (Core.C_pretty.Index+2)
  Result = F_lbreak_void()
  if !ErrorIn(Result) {
  Result = Core.F_CALL(C_print,ARGS(self.Arg.ToEID()))
  if !ErrorIn(Result) {
  { 
    var va_arg1 *Core.PrettyPrinter
    var va_arg2 int
    va_arg1 = Core.C_pretty
    va_arg2 = (Core.C_pretty.Index-2)
    va_arg1.Index = va_arg2
    Result = EID{C__INT,IVAL(va_arg2)}
    } 
  }}
  }}
  return Result} 

// The EID go function for: self_print @ For (throw: true) 
func E_self_print_For_Language (self EID) EID { 
  return To_For(OBJ(self)).SelfPrint( )} 

/* The go function for: self_eval(self:For) [status=1] */
func (self *For) SelfEval () EID { 
  var Result EID
  { var x *ClaireAny
    var try_1 EID
    try_1 = EVAL(self.SetArg)
    if ErrorIn(try_1) {Result = try_1
    } else {
    x = ANY(try_1)
    { 
      h_index := ClEnv.Index
      h_base := ClEnv.Base
      if (C_class.Id() == x.Isa.Id()) { 
        { var g0148 *ClaireClass = ToClass(x)
          { 
            var y *ClaireClass
            _ = y
            var y_iter *ClaireAny
            Result= EID{CFALSE.Id(),0}
            var y_support *ClaireSet
            y_support = g0148.Descendants
            for i_it := 0; i_it < y_support.Count; i_it++ { 
              y_iter = y_support.At(i_it)
              y = ToClass(y_iter)
              var loop_2 EID
              _ = loop_2
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
                  loop_3 = F_write_value_Variable(self.ClaireVar,z)
                  if ErrorIn(loop_3) {loop_2 = loop_3
                  break
                  } else {
                  loop_3 = EVAL(self.Arg)
                  if ErrorIn(loop_3) {loop_2 = loop_3
                  break
                  } else {
                  }}
                  }
                  } 
                } 
              if ErrorIn(loop_2) {Result = loop_2
              break
              } else {
              }
              } 
            } 
          } 
        }  else if (x.Isa.IsIn(C_list) == CTRUE) { 
        { var g0149 *ClaireList = ToList(x)
          { 
            var z *ClaireAny
            _ = z
            Result= EID{CFALSE.Id(),0}
            var z_support *ClaireList
            z_support = g0149
            z_len := z_support.Length()
            for i_it := 0; i_it < z_len; i_it++ { 
              z = z_support.At(i_it)
              var loop_4 EID
              _ = loop_4
              { 
              loop_4 = F_write_value_Variable(self.ClaireVar,z)
              if ErrorIn(loop_4) {Result = loop_4
              break
              } else {
              loop_4 = EVAL(self.Arg)
              if ErrorIn(loop_4) {Result = loop_4
              break
              } else {
              }}
              }
              } 
            } 
          } 
        }  else if (x.Isa.IsIn(C_array) == CTRUE) { 
        { var g0150 *ClaireList = ToArray(x)
          { var n int = g0150.Length()
            { var g0151 int = 1
              { var g0152 int = n
                Result= EID{CFALSE.Id(),0}
                for (g0151 <= g0152) { 
                  var loop_5 EID
                  _ = loop_5
                  { 
                  { var z *ClaireAny = ToList(g0150.Id()).At(g0151-1)
                    loop_5 = F_write_value_Variable(self.ClaireVar,z)
                    if ErrorIn(loop_5) {Result = loop_5
                    break
                    } else {
                    loop_5 = EVAL(self.Arg)
                    if ErrorIn(loop_5) {Result = loop_5
                    break
                    } else {
                    }}
                    } 
                  if ErrorIn(loop_5) {Result = loop_5
                  break
                  } else {
                  g0151 = (g0151+1)
                  }
                  } 
                }
                } 
              } 
            } 
          } 
        }  else if (x.Isa.IsIn(C_Interval) == CTRUE) { 
        { var g0153 *ClaireInterval = To_Interval(x)
          { var y int = g0153.Arg1
            { var g0154 int = g0153.Arg2
              Result= EID{CFALSE.Id(),0}
              for (y <= g0154) { 
                var loop_6 EID
                _ = loop_6
                { 
                loop_6 = self.ClaireVar.WriteEID(EID{C__INT,IVAL(y)})
                if ErrorIn(loop_6) {Result = loop_6
                break
                } else {
                loop_6 = EVAL(self.Arg)
                if ErrorIn(loop_6) {Result = loop_6
                break
                } else {
                }}
                if ErrorIn(loop_6) {Result = loop_6
                break
                } else {
                y = (y+1)
                }
                } 
              }
              } 
            } 
          } 
        }  else if (x.Isa.IsIn(C_collection) == CTRUE) { 
        { var g0155 *ClaireCollection = ToCollection(x)
          { 
            var y *ClaireAny
            _ = y
            Result= EID{CFALSE.Id(),0}
            var y_support *ClaireList
            var try_7 EID
            try_7 = Core.F_enumerate_any(g0155.Id())
            if ErrorIn(try_7) {Result = try_7
            } else {
            y_support = ToList(OBJ(try_7))
            y_len := y_support.Length()
            for i_it := 0; i_it < y_len; i_it++ { 
              y = y_support.At(i_it)
              var loop_8 EID
              _ = loop_8
              { 
              loop_8 = F_write_value_Variable(self.ClaireVar,y)
              if ErrorIn(loop_8) {Result = loop_8
              break
              } else {
              loop_8 = EVAL(self.Arg)
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
/* The go function for: self_print(self:Collect) [status=1] */
func (self *Collect) SelfPrint () EID { 
  var Result EID
  PRINC("list{ ")
  Core.C_pretty.Index = (Core.C_pretty.Index+2)
  Result = F_printexp_any(self.Arg,CFALSE)
  if !ErrorIn(Result) {
  PRINC(" | ")
  Result = F_lbreak_void()
  if !ErrorIn(Result) {
  Result = F_ppvariable_Variable(self.ClaireVar)
  if !ErrorIn(Result) {
  PRINC(" in ")
  { var _Zl int = Core.C_pretty.Index
    F_set_level_void()
    Result = F_printexp_any(self.SetArg,CFALSE)
    if !ErrorIn(Result) {
    { 
      var va_arg1 *Core.PrettyPrinter
      var va_arg2 int
      va_arg1 = Core.C_pretty
      va_arg2 = (_Zl-2)
      va_arg1.Index = va_arg2
      Result = EID{C__INT,IVAL(va_arg2)}
      } 
    }
    } 
  if !ErrorIn(Result) {
  PRINC("}")
  Result = EVOID
  }}}}
  return Result} 

// The EID go function for: self_print @ Collect (throw: true) 
func E_self_print_Collect_Language (self EID) EID { 
  return To_Collect(OBJ(self)).SelfPrint( )} 

// list image : preserve the order for lists and intervals (v4)
/* The go function for: self_eval(self:Collect) [status=1] */
func (self *Collect) SelfEval () EID { 
  var Result EID
  { var x *ClaireAny
    var try_1 EID
    try_1 = EVAL(self.SetArg)
    if ErrorIn(try_1) {Result = try_1
    } else {
    x = ANY(try_1)
    { var res *ClaireList
      { var arg_2 *ClaireType
        if (self.Of.Id() != CNULL) { 
          arg_2 = self.Of
          } else {
          arg_2 = ToType(CEMPTY.Id())
          } 
        res = arg_2.EmptyList()
        } 
      if (C_class.Id() == x.Isa.Id()) { 
        { var g0158 *ClaireClass = ToClass(x)
          { 
            var y *ClaireClass
            _ = y
            var y_iter *ClaireAny
            Result= EID{CFALSE.Id(),0}
            var y_support *ClaireSet
            y_support = g0158.Descendants
            for i_it := 0; i_it < y_support.Count; i_it++ { 
              y_iter = y_support.At(i_it)
              y = ToClass(y_iter)
              var loop_3 EID
              _ = loop_3
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
                  loop_4 = F_write_value_Variable(self.ClaireVar,z)
                  if ErrorIn(loop_4) {loop_3 = loop_4
                  break
                  } else {
                  var try_5 EID
                  { var arg_6 *ClaireAny
                    var try_7 EID
                    try_7 = EVAL(self.Arg)
                    if ErrorIn(try_7) {try_5 = try_7
                    } else {
                    arg_6 = ANY(try_7)
                    try_5 = EID{res.AddFast(arg_6).Id(),0}
                    }
                    } 
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
              if ErrorIn(loop_3) {Result = loop_3
              break
              } else {
              }
              } 
            } 
          } 
        }  else if (x.Isa.IsIn(C_list) == CTRUE) { 
        { var g0159 *ClaireList = ToList(x)
          { 
            var y *ClaireAny
            _ = y
            Result= EID{CFALSE.Id(),0}
            var y_support *ClaireList
            y_support = g0159
            y_len := y_support.Length()
            for i_it := 0; i_it < y_len; i_it++ { 
              y = y_support.At(i_it)
              var loop_8 EID
              _ = loop_8
              { 
              loop_8 = F_write_value_Variable(self.ClaireVar,y)
              if ErrorIn(loop_8) {Result = loop_8
              break
              } else {
              var try_9 EID
              { var arg_10 *ClaireAny
                var try_11 EID
                try_11 = EVAL(self.Arg)
                if ErrorIn(try_11) {try_9 = try_11
                } else {
                arg_10 = ANY(try_11)
                try_9 = EID{res.AddFast(arg_10).Id(),0}
                }
                } 
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
        { var g0160 *ClaireInterval = To_Interval(x)
          { var y int = g0160.Arg1
            { var g0161 int = g0160.Arg2
              Result= EID{CFALSE.Id(),0}
              for (y <= g0161) { 
                var loop_12 EID
                _ = loop_12
                { 
                loop_12 = self.ClaireVar.WriteEID(EID{C__INT,IVAL(y)})
                if ErrorIn(loop_12) {Result = loop_12
                break
                } else {
                var try_13 EID
                { var arg_14 *ClaireAny
                  var try_15 EID
                  try_15 = EVAL(self.Arg)
                  if ErrorIn(try_15) {try_13 = try_15
                  } else {
                  arg_14 = ANY(try_15)
                  try_13 = EID{res.AddFast(arg_14).Id(),0}
                  }
                  } 
                if ErrorIn(try_13) {loop_12 = try_13
                Result = try_13
                break
                } else {
                res = ToList(OBJ(try_13))
                loop_12 = EID{res.Id(),0}
                }}
                if ErrorIn(loop_12) {Result = loop_12
                break
                } else {
                y = (y+1)
                }
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
          try_16 = Core.F_enumerate_any(x)
          if ErrorIn(try_16) {Result = try_16
          } else {
          y_support = ToList(OBJ(try_16))
          y_len := y_support.Length()
          for i_it := 0; i_it < y_len; i_it++ { 
            y = y_support.At(i_it)
            var loop_17 EID
            _ = loop_17
            { 
            loop_17 = F_write_value_Variable(self.ClaireVar,y)
            if ErrorIn(loop_17) {Result = loop_17
            break
            } else {
            var try_18 EID
            { var arg_19 *ClaireAny
              var try_20 EID
              try_20 = EVAL(self.Arg)
              if ErrorIn(try_20) {try_18 = try_20
              } else {
              arg_19 = ANY(try_20)
              try_18 = EID{res.AddFast(arg_19).Id(),0}
              }
              } 
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
/* The go function for: self_print(self:Image) [status=1] */
func (self *Image) SelfPrint () EID { 
  var Result EID
  PRINC("{ ")
  Core.C_pretty.Index = (Core.C_pretty.Index+2)
  Result = F_printexp_any(self.Arg,CFALSE)
  if !ErrorIn(Result) {
  PRINC(" | ")
  Result = F_lbreak_void()
  if !ErrorIn(Result) {
  Result = F_ppvariable_Variable(self.ClaireVar)
  if !ErrorIn(Result) {
  PRINC(" in ")
  { var _Zl int = Core.C_pretty.Index
    F_set_level_void()
    Result = F_printexp_any(self.SetArg,CFALSE)
    if !ErrorIn(Result) {
    { 
      var va_arg1 *Core.PrettyPrinter
      var va_arg2 int
      va_arg1 = Core.C_pretty
      va_arg2 = (_Zl-2)
      va_arg1.Index = va_arg2
      Result = EID{C__INT,IVAL(va_arg2)}
      } 
    }
    } 
  if !ErrorIn(Result) {
  PRINC("}")
  Result = EVOID
  }}}}
  return Result} 

// The EID go function for: self_print @ Image (throw: true) 
func E_self_print_Image_Language (self EID) EID { 
  return To_Image(OBJ(self)).SelfPrint( )} 

/* The go function for: self_eval(self:Image) [status=1] */
func (self *Image) SelfEval () EID { 
  var Result EID
  { var x *ClaireAny
    var try_1 EID
    try_1 = EVAL(self.SetArg)
    if ErrorIn(try_1) {Result = try_1
    } else {
    x = ANY(try_1)
    { var res *ClaireSet
      { var arg_2 *ClaireType
        if (self.Of.Id() != CNULL) { 
          arg_2 = self.Of
          } else {
          arg_2 = ToType(CEMPTY.Id())
          } 
        res = arg_2.EmptySet()
        } 
      { 
        var y *ClaireAny
        _ = y
        Result= EID{CFALSE.Id(),0}
        var y_support *ClaireList
        var try_3 EID
        try_3 = Core.F_enumerate_any(x)
        if ErrorIn(try_3) {Result = try_3
        } else {
        y_support = ToList(OBJ(try_3))
        y_len := y_support.Length()
        for i_it := 0; i_it < y_len; i_it++ { 
          y = y_support.At(i_it)
          var loop_4 EID
          _ = loop_4
          { 
          loop_4 = F_write_value_Variable(self.ClaireVar,y)
          if ErrorIn(loop_4) {Result = loop_4
          break
          } else {
          var try_5 EID
          { var arg_6 *ClaireAny
            var try_7 EID
            try_7 = EVAL(self.Arg)
            if ErrorIn(try_7) {try_5 = try_7
            } else {
            arg_6 = ANY(try_7)
            try_5 = EID{res.AddFast(arg_6).Id(),0}
            }
            } 
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
/* The go function for: self_print(self:Select) [status=1] */
func (self *Select) SelfPrint () EID { 
  var Result EID
  PRINC("{ ")
  Result = F_ppvariable_Variable(self.ClaireVar)
  if !ErrorIn(Result) {
  PRINC(" in ")
  { var _Zl int = Core.C_pretty.Index
    F_set_level_void()
    Result = F_printexp_any(self.SetArg,CFALSE)
    if !ErrorIn(Result) {
    { 
      var va_arg1 *Core.PrettyPrinter
      var va_arg2 int
      va_arg1 = Core.C_pretty
      va_arg2 = _Zl
      va_arg1.Index = va_arg2
      Result = EID{C__INT,IVAL(va_arg2)}
      } 
    }
    } 
  if !ErrorIn(Result) {
  PRINC(" | ")
  Result = F_lbreak_integer(2)
  if !ErrorIn(Result) {
  Result = Core.F_CALL(C_print,ARGS(self.Arg.ToEID()))
  if !ErrorIn(Result) {
  { 
    var va_arg1 *Core.PrettyPrinter
    var va_arg2 int
    va_arg1 = Core.C_pretty
    va_arg2 = (Core.C_pretty.Index-2)
    va_arg1.Index = va_arg2
    Result = EID{C__INT,IVAL(va_arg2)}
    } 
  }}
  if !ErrorIn(Result) {
  PRINC("}")
  Result = EVOID
  }}}
  return Result} 

// The EID go function for: self_print @ Select (throw: true) 
func E_self_print_Select_Language (self EID) EID { 
  return To_Select(OBJ(self)).SelfPrint( )} 

/* The go function for: self_eval(self:Select) [status=1] */
func (self *Select) SelfEval () EID { 
  var Result EID
  { var x *ClaireAny
    var try_1 EID
    try_1 = EVAL(self.SetArg)
    if ErrorIn(try_1) {Result = try_1
    } else {
    x = ANY(try_1)
    { var res *ClaireSet
      { var arg_2 *ClaireType
        if (self.Of.Id() != CNULL) { 
          arg_2 = self.Of
          } else {
          arg_2 = ToType(CEMPTY.Id())
          } 
        res = arg_2.EmptySet()
        } 
      if (C_class.Id() == x.Isa.Id()) { 
        { var g0165 *ClaireClass = ToClass(x)
          { 
            var y *ClaireClass
            _ = y
            var y_iter *ClaireAny
            Result= EID{CFALSE.Id(),0}
            var y_support *ClaireSet
            y_support = g0165.Descendants
            for i_it := 0; i_it < y_support.Count; i_it++ { 
              y_iter = y_support.At(i_it)
              y = ToClass(y_iter)
              var loop_3 EID
              _ = loop_3
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
                  loop_4 = F_write_value_Variable(self.ClaireVar,z)
                  if ErrorIn(loop_4) {loop_3 = loop_4
                  break
                  } else {
                  var g0169I *ClaireBoolean
                  var try_5 EID
                  { var arg_6 *ClaireAny
                    var try_7 EID
                    try_7 = EVAL(self.Arg)
                    if ErrorIn(try_7) {try_5 = try_7
                    } else {
                    arg_6 = ANY(try_7)
                    try_5 = EID{Core.F__I_equal_any(arg_6,CFALSE.Id()).Id(),0}
                    }
                    } 
                  if ErrorIn(try_5) {loop_4 = try_5
                  } else {
                  g0169I = ToBoolean(OBJ(try_5))
                  if (g0169I == CTRUE) { 
                    res = res.AddFast(z)
                    loop_4 = EID{res.Id(),0}
                    } else {
                    loop_4 = EID{CFALSE.Id(),0}
                    } 
                  }
                  if ErrorIn(loop_4) {loop_3 = loop_4
                  break
                  } else {
                  }}
                  }
                  } 
                } 
              if ErrorIn(loop_3) {Result = loop_3
              break
              } else {
              }
              } 
            } 
          } 
        }  else if (x.Isa.IsIn(C_Interval) == CTRUE) { 
        { var g0166 *ClaireInterval = To_Interval(x)
          { var y int = g0166.Arg1
            { var g0167 int = g0166.Arg2
              Result= EID{CFALSE.Id(),0}
              for (y <= g0167) { 
                var loop_8 EID
                _ = loop_8
                { 
                loop_8 = self.ClaireVar.WriteEID(EID{C__INT,IVAL(y)})
                if ErrorIn(loop_8) {Result = loop_8
                break
                } else {
                var g0170I *ClaireBoolean
                var try_9 EID
                { var arg_10 *ClaireAny
                  var try_11 EID
                  try_11 = EVAL(self.Arg)
                  if ErrorIn(try_11) {try_9 = try_11
                  } else {
                  arg_10 = ANY(try_11)
                  try_9 = EID{Core.F__I_equal_any(arg_10,CFALSE.Id()).Id(),0}
                  }
                  } 
                if ErrorIn(try_9) {loop_8 = try_9
                } else {
                g0170I = ToBoolean(OBJ(try_9))
                if (g0170I == CTRUE) { 
                  res = res.AddFast(MakeInteger(y).Id())
                  loop_8 = EID{res.Id(),0}
                  } else {
                  loop_8 = EID{CFALSE.Id(),0}
                  } 
                }
                if ErrorIn(loop_8) {Result = loop_8
                break
                } else {
                }}
                if ErrorIn(loop_8) {Result = loop_8
                break
                } else {
                y = (y+1)
                }
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
          try_12 = Core.F_enumerate_any(x)
          if ErrorIn(try_12) {Result = try_12
          } else {
          y_support = ToList(OBJ(try_12))
          y_len := y_support.Length()
          for i_it := 0; i_it < y_len; i_it++ { 
            y = y_support.At(i_it)
            var loop_13 EID
            _ = loop_13
            { 
            loop_13 = F_write_value_Variable(self.ClaireVar,y)
            if ErrorIn(loop_13) {Result = loop_13
            break
            } else {
            var g0171I *ClaireBoolean
            var try_14 EID
            { var arg_15 *ClaireAny
              var try_16 EID
              try_16 = EVAL(self.Arg)
              if ErrorIn(try_16) {try_14 = try_16
              } else {
              arg_15 = ANY(try_16)
              try_14 = EID{Core.F__I_equal_any(arg_15,CFALSE.Id()).Id(),0}
              }
              } 
            if ErrorIn(try_14) {loop_13 = try_14
            } else {
            g0171I = ToBoolean(OBJ(try_14))
            if (g0171I == CTRUE) { 
              res = res.AddFast(y)
              loop_13 = EID{res.Id(),0}
              } else {
              loop_13 = EID{CFALSE.Id(),0}
              } 
            }
            if ErrorIn(loop_13) {Result = loop_13
            break
            } else {
            }}
            }}
            } 
          } 
        } 
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
/* The go function for: self_print(self:Lselect) [status=1] */
func (self *Lselect) SelfPrint () EID { 
  var Result EID
  PRINC("list{ ")
  Result = F_ppvariable_Variable(self.ClaireVar)
  if !ErrorIn(Result) {
  PRINC(" in ")
  { var _Zl int = Core.C_pretty.Index
    F_set_level_void()
    Result = F_printexp_any(self.SetArg,CFALSE)
    if !ErrorIn(Result) {
    { 
      var va_arg1 *Core.PrettyPrinter
      var va_arg2 int
      va_arg1 = Core.C_pretty
      va_arg2 = _Zl
      va_arg1.Index = va_arg2
      Result = EID{C__INT,IVAL(va_arg2)}
      } 
    }
    } 
  if !ErrorIn(Result) {
  PRINC(" | ")
  Result = F_lbreak_integer(2)
  if !ErrorIn(Result) {
  Result = Core.F_CALL(C_print,ARGS(self.Arg.ToEID()))
  if !ErrorIn(Result) {
  { 
    var va_arg1 *Core.PrettyPrinter
    var va_arg2 int
    va_arg1 = Core.C_pretty
    va_arg2 = (Core.C_pretty.Index-2)
    va_arg1.Index = va_arg2
    Result = EID{C__INT,IVAL(va_arg2)}
    } 
  }}
  if !ErrorIn(Result) {
  PRINC("}")
  Result = EVOID
  }}}
  return Result} 

// The EID go function for: self_print @ Lselect (throw: true) 
func E_self_print_Lselect_Language (self EID) EID { 
  return To_Lselect(OBJ(self)).SelfPrint( )} 

/* The go function for: self_eval(self:Lselect) [status=1] */
func (self *Lselect) SelfEval () EID { 
  var Result EID
  { var x *ClaireAny
    var try_1 EID
    try_1 = EVAL(self.SetArg)
    if ErrorIn(try_1) {Result = try_1
    } else {
    x = ANY(try_1)
    { var res *ClaireList
      if (x.Isa.IsIn(C_list) == CTRUE) { 
        { var g0173 *ClaireList = ToList(x)
          res = g0173.Empty()
          } 
        } else {
        res = ToType(CEMPTY.Id()).EmptyList()
        } 
      if (C_class.Id() == x.Isa.Id()) { 
        { var g0175 *ClaireClass = ToClass(x)
          { 
            var y *ClaireClass
            _ = y
            var y_iter *ClaireAny
            Result= EID{CFALSE.Id(),0}
            var y_support *ClaireSet
            y_support = g0175.Descendants
            for i_it := 0; i_it < y_support.Count; i_it++ { 
              y_iter = y_support.At(i_it)
              y = ToClass(y_iter)
              var loop_2 EID
              _ = loop_2
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
                  loop_3 = F_write_value_Variable(self.ClaireVar,z)
                  if ErrorIn(loop_3) {loop_2 = loop_3
                  break
                  } else {
                  var g0177I *ClaireBoolean
                  var try_4 EID
                  { var arg_5 *ClaireAny
                    var try_6 EID
                    try_6 = EVAL(self.Arg)
                    if ErrorIn(try_6) {try_4 = try_6
                    } else {
                    arg_5 = ANY(try_6)
                    try_4 = EID{Core.F__I_equal_any(arg_5,CFALSE.Id()).Id(),0}
                    }
                    } 
                  if ErrorIn(try_4) {loop_3 = try_4
                  } else {
                  g0177I = ToBoolean(OBJ(try_4))
                  if (g0177I == CTRUE) { 
                    res = res.AddFast(z)
                    loop_3 = EID{res.Id(),0}
                    } else {
                    loop_3 = EID{CFALSE.Id(),0}
                    } 
                  }
                  if ErrorIn(loop_3) {loop_2 = loop_3
                  break
                  } else {
                  }}
                  }
                  } 
                } 
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
          try_7 = Core.F_enumerate_any(x)
          if ErrorIn(try_7) {Result = try_7
          } else {
          y_support = ToList(OBJ(try_7))
          y_len := y_support.Length()
          for i_it := 0; i_it < y_len; i_it++ { 
            y = y_support.At(i_it)
            var loop_8 EID
            _ = loop_8
            { 
            loop_8 = F_write_value_Variable(self.ClaireVar,y)
            if ErrorIn(loop_8) {Result = loop_8
            break
            } else {
            var g0178I *ClaireBoolean
            var try_9 EID
            { var arg_10 *ClaireAny
              var try_11 EID
              try_11 = EVAL(self.Arg)
              if ErrorIn(try_11) {try_9 = try_11
              } else {
              arg_10 = ANY(try_11)
              try_9 = EID{Core.F__I_equal_any(arg_10,CFALSE.Id()).Id(),0}
              }
              } 
            if ErrorIn(try_9) {loop_8 = try_9
            } else {
            g0178I = ToBoolean(OBJ(try_9))
            if (g0178I == CTRUE) { 
              res = res.AddFast(y)
              loop_8 = EID{res.Id(),0}
              } else {
              loop_8 = EID{CFALSE.Id(),0}
              } 
            }
            if ErrorIn(loop_8) {Result = loop_8
            break
            } else {
            }}
            }}
            } 
          } 
        } 
      if !ErrorIn(Result) {
      if (self.Of.Id() != CNULL) { 
        { var x *ClaireAny
          { var x_some *ClaireAny = CNULL
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
            Result = ToException(Core.C_range_error.Make(self.Id(),x,self.Of.Id())).Close()
            } else {
            Result = EID{CNULL,0}
            } 
          } 
        if !ErrorIn(Result) {
        Result = EID{res.Cast_I(self.Of).Id(),0}
        }
        } else {
        Result = EID{CFALSE.Id(),0}
        } 
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
/* The go function for: self_print(self:Exists) [status=1] */
func (self *Exists) SelfPrint () EID { 
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
    Result = F_ppvariable_Variable(self.ClaireVar)
    if !ErrorIn(Result) {
    PRINC(",")
    Result = Core.F_CALL(C_print,ARGS(self.Arg.ToEID()))
    if !ErrorIn(Result) {
    PRINC(")")
    Result = EVOID
    }}
    } else {
    PRINC("(")
    Result = F_ppvariable_Variable(self.ClaireVar)
    if !ErrorIn(Result) {
    PRINC(" in ")
    { var _Zl int = Core.C_pretty.Index
      F_set_level_void()
      Result = F_printexp_any(self.SetArg,CFALSE)
      if !ErrorIn(Result) {
      { 
        var va_arg1 *Core.PrettyPrinter
        var va_arg2 int
        va_arg1 = Core.C_pretty
        va_arg2 = _Zl
        va_arg1.Index = va_arg2
        Result = EID{C__INT,IVAL(va_arg2)}
        } 
      }
      } 
    if !ErrorIn(Result) {
    PRINC(" | ")
    Result = F_lbreak_integer(2)
    if !ErrorIn(Result) {
    Result = Core.F_CALL(C_print,ARGS(self.Arg.ToEID()))
    if !ErrorIn(Result) {
    { 
      var va_arg1 *Core.PrettyPrinter
      var va_arg2 int
      va_arg1 = Core.C_pretty
      va_arg2 = (Core.C_pretty.Index-2)
      va_arg1.Index = va_arg2
      Result = EID{C__INT,IVAL(va_arg2)}
      } 
    }}
    if !ErrorIn(Result) {
    PRINC(")")
    Result = EVOID
    }}}
    } 
  return Result} 

// The EID go function for: self_print @ Exists (throw: true) 
func E_self_print_Exists_Language (self EID) EID { 
  return To_Exists(OBJ(self)).SelfPrint( )} 

/* The go function for: self_eval(self:Exists) [status=1] */
func (self *Exists) SelfEval () EID { 
  var Result EID
  { var x *ClaireAny
    var try_1 EID
    try_1 = EVAL(self.SetArg)
    if ErrorIn(try_1) {Result = try_1
    } else {
    x = ANY(try_1)
    { var b *ClaireAny = self.Other
      { var res *ClaireAny = b
        if (C_class.Id() == x.Isa.Id()) { 
          { var g0180 *ClaireClass = ToClass(x)
            { 
              var y *ClaireClass
              _ = y
              var y_iter *ClaireAny
              Result= EID{CFALSE.Id(),0}
              var y_support *ClaireSet
              y_support = g0180.Descendants
              for i_it := 0; i_it < y_support.Count; i_it++ { 
                y_iter = y_support.At(i_it)
                y = ToClass(y_iter)
                var loop_2 EID
                _ = loop_2
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
                    loop_3 = F_write_value_Variable(self.ClaireVar,z)
                    if ErrorIn(loop_3) {loop_2 = loop_3
                    break
                    } else {
                    var g0182I *ClaireBoolean
                    var try_4 EID
                    { var arg_5 *ClaireAny
                      var try_6 EID
                      try_6 = EVAL(self.Arg)
                      if ErrorIn(try_6) {try_4 = try_6
                      } else {
                      arg_5 = ANY(try_6)
                      try_4 = EID{Core.F__I_equal_any(arg_5,CFALSE.Id()).Id(),0}
                      }
                      } 
                    if ErrorIn(try_4) {loop_3 = try_4
                    } else {
                    g0182I = ToBoolean(OBJ(try_4))
                    if (g0182I == CTRUE) { 
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
                    if ErrorIn(loop_3) {loop_2 = loop_3
                    break
                    } else {
                    }}
                    }
                    } 
                  } 
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
            try_7 = Core.F_enumerate_any(x)
            if ErrorIn(try_7) {Result = try_7
            } else {
            y_support = ToList(OBJ(try_7))
            y_len := y_support.Length()
            for i_it := 0; i_it < y_len; i_it++ { 
              y = y_support.At(i_it)
              var loop_8 EID
              _ = loop_8
              { 
              loop_8 = F_write_value_Variable(self.ClaireVar,y)
              if ErrorIn(loop_8) {Result = loop_8
              break
              } else {
              var g0183I *ClaireBoolean
              var try_9 EID
              { var arg_10 *ClaireAny
                var try_11 EID
                try_11 = EVAL(self.Arg)
                if ErrorIn(try_11) {try_9 = try_11
                } else {
                arg_10 = ANY(try_11)
                try_9 = EID{Core.F__I_equal_any(arg_10,CFALSE.Id()).Id(),0}
                }
                } 
              if ErrorIn(try_9) {loop_8 = try_9
              } else {
              g0183I = ToBoolean(OBJ(try_9))
              if (g0183I == CTRUE) { 
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
              if ErrorIn(loop_8) {Result = loop_8
              break
              } else {
              }}
              }}
              } 
            } 
          } 
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
/* The go function for: self_print(self:Case) [status=1] */
func (self *Case) SelfPrint () EID { 
  var Result EID
  PRINC("case ")
  Result = Core.F_CALL(C_print,ARGS(self.ClaireVar.ToEID()))
  if !ErrorIn(Result) {
  PRINC(" ")
  Result = F_lbreak_integer(1)
  if !ErrorIn(Result) {
  PRINC("(")
  Result = EVOID
  }}
  if !ErrorIn(Result) {
  { var n int = 1
    { var m int = self.Args.Length()
      Core.C_pretty.Index = (Core.C_pretty.Index+1)
      Result= EID{CFALSE.Id(),0}
      for (n <= m) { 
        var loop_1 EID
        _ = loop_1
        { var _Zl int = Core.C_pretty.Index
          loop_1 = F_printexp_any(self.Args.At(n-1),CFALSE)
          if ErrorIn(loop_1) {Result = loop_1
          break
          } else {
          PRINC(" ")
          if (Core.F_buffer_length_void() > (Core.C_pretty.Width-50)) { 
            loop_1 = F_lbreak_integer(2)
            } else {
            F_set_level_void()
            loop_1 = EVOID
            } 
          if ErrorIn(loop_1) {Result = loop_1
          break
          } else {
          loop_1 = Core.F_CALL(C_print,ARGS(self.Args.At((n+1)-1).ToEID()))
          if ErrorIn(loop_1) {Result = loop_1
          break
          } else {
          }}
          if ErrorIn(loop_1) {Result = loop_1
          break
          } else {
          Core.C_pretty.Index = _Zl
          if ((n+1) != m) { 
            PRINC(", ")
            loop_1 = F_lbreak_void()
            if ErrorIn(loop_1) {Result = loop_1
            break
            } else {
            }
            } else {
            loop_1 = EID{CFALSE.Id(),0}
            } 
          if ErrorIn(loop_1) {Result = loop_1
          break
          } else {
          }
          if ErrorIn(loop_1) {Result = loop_1
          break
          } else {
          }}}
          if ErrorIn(loop_1) {Result = loop_1
          break
          } else {
          n = (n+2)
          loop_1 = EID{C__INT,IVAL(n)}
          }
          } 
        if ErrorIn(loop_1) {Result = loop_1
        break
        } else {
        } 
      }
      if !ErrorIn(Result) {
      PRINC(")")
      { 
        var va_arg1 *Core.PrettyPrinter
        var va_arg2 int
        va_arg1 = Core.C_pretty
        va_arg2 = (Core.C_pretty.Index-2)
        va_arg1.Index = va_arg2
        Result = EID{C__INT,IVAL(va_arg2)}
        } 
      }
      } 
    } 
  }
  return Result} 

// The EID go function for: self_print @ Case (throw: true) 
func E_self_print_Case_Language (self EID) EID { 
  return To_Case(OBJ(self)).SelfPrint( )} 

/* The go function for: self_eval(self:Case) [status=1] */
func (self *Case) SelfEval () EID { 
  var Result EID
  { var truc *ClaireAny
    var try_1 EID
    try_1 = EVAL(self.ClaireVar)
    if ErrorIn(try_1) {Result = try_1
    } else {
    truc = ANY(try_1)
    { var flip *ClaireBoolean = CTRUE
      { var previous *ClaireAny = CFALSE.Id()
        var g0186I *ClaireBoolean
        var try_2 EID
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
            if (flip == CTRUE) { 
              flip = CFALSE
              var try_4 EID
              try_4 = EVAL(x)
              if ErrorIn(try_4) {loop_3 = try_4
              try_2 = try_4
              break
              } else {
              previous = ANY(try_4)
              loop_3 = previous.ToEID()
              }
              } else {
              var g0187I *ClaireBoolean
              var try_5 EID
              try_5 = Core.F_BELONG(truc,previous)
              if ErrorIn(try_5) {loop_3 = try_5
              } else {
              g0187I = ToBoolean(OBJ(try_5))
              if (g0187I == CTRUE) { 
                var try_6 EID
                try_6 = EVAL(x)
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
            if ErrorIn(loop_3) {try_2 = loop_3
            break
            } else {
            }
            } 
          } 
        if ErrorIn(try_2) {Result = try_2
        } else {
        g0186I = ToBoolean(OBJ(try_2))
        if (g0186I == CTRUE) { 
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
/* The go function for: self_print(self:While) [status=1] */
func (self *While) SelfPrint () EID { 
  var Result EID
  F_princ_string(ToString(IfThenElse((self.Other == CTRUE),
    MakeString("until").Id(),
    MakeString("while").Id())))
  PRINC(" ")
  Result = F_printexp_any(self.Test,CFALSE)
  if !ErrorIn(Result) {
  PRINC(" ")
  Result = F_lbreak_integer(2)
  if !ErrorIn(Result) {
  Result = Core.F_CALL(C_print,ARGS(self.Arg.ToEID()))
  }}
  if !ErrorIn(Result) {
  { 
    var va_arg1 *Core.PrettyPrinter
    var va_arg2 int
    va_arg1 = Core.C_pretty
    va_arg2 = (Core.C_pretty.Index-2)
    va_arg1.Index = va_arg2
    Result = EID{C__INT,IVAL(va_arg2)}
    } 
  }
  return Result} 

// The EID go function for: self_print @ While (throw: true) 
func E_self_print_While_Language (self EID) EID { 
  return To_While(OBJ(self)).SelfPrint( )} 

// other = true => self means  repeat self.arg until self.test = true
/* The go function for: self_eval(self:While) [status=1] */
func (self *While) SelfEval () EID { 
  var Result EID
  { var a *ClaireBoolean = self.Other
    { var b *ClaireBoolean = a
      { 
        h_index := ClEnv.Index
        h_base := ClEnv.Base
        var v_while4 *ClaireBoolean
        Result= EID{CFALSE.Id(),0}
        var try_1 EID
        { 
          var v_or4 *ClaireBoolean
          
          v_or4 = b
          if (v_or4 == CTRUE) {try_1 = EID{CTRUE.Id(),0}
          } else { 
            var try_2 EID
            { var arg_3 *ClaireBoolean
              var try_4 EID
              { var arg_5 *ClaireAny
                var try_6 EID
                try_6 = EVAL(self.Test)
                if ErrorIn(try_6) {try_4 = try_6
                } else {
                arg_5 = ANY(try_6)
                try_4 = EID{Core.F_not_any(arg_5).Id(),0}
                }
                } 
              if ErrorIn(try_4) {try_2 = try_4
              } else {
              arg_3 = ToBoolean(OBJ(try_4))
              try_2 = EID{Equal(arg_3.Id(),a.Id()).Id(),0}
              }
              } 
            if ErrorIn(try_2) {try_1 = try_2
            } else {
            v_or4 = ToBoolean(OBJ(try_2))
            if (v_or4 == CTRUE) {try_1 = EID{CTRUE.Id(),0}
            } else { 
              try_1 = EID{CFALSE.Id(),0}} 
            } 
          }
          } 
        if ErrorIn(try_1) {Result = try_1
        } else {
        v_while4 = ToBoolean(OBJ(try_1))
        
        for v_while4 == CTRUE { 
          var loop_7 EID
          _ = loop_7
          { 
          b = CFALSE
          loop_7 = EVAL(self.Arg)
          if ErrorIn(loop_7) {Result = loop_7
          break
          } else {
          }
          var try_8 EID
          { 
            var v_or5 *ClaireBoolean
            
            v_or5 = b
            if (v_or5 == CTRUE) {try_8 = EID{CTRUE.Id(),0}
            } else { 
              var try_9 EID
              { var arg_10 *ClaireBoolean
                var try_11 EID
                { var arg_12 *ClaireAny
                  var try_13 EID
                  try_13 = EVAL(self.Test)
                  if ErrorIn(try_13) {try_11 = try_13
                  } else {
                  arg_12 = ANY(try_13)
                  try_11 = EID{Core.F_not_any(arg_12).Id(),0}
                  }
                  } 
                if ErrorIn(try_11) {try_9 = try_11
                } else {
                arg_10 = ToBoolean(OBJ(try_11))
                try_9 = EID{Equal(arg_10.Id(),a.Id()).Id(),0}
                }
                } 
              if ErrorIn(try_9) {try_8 = try_9
              Result = try_9
              break
              } else {
              v_or5 = ToBoolean(OBJ(try_9))
              if (v_or5 == CTRUE) {try_8 = EID{CTRUE.Id(),0}
              } else { 
                try_8 = EID{CFALSE.Id(),0}} 
              } 
            }
            } 
          if ErrorIn(try_8) {Result = try_8
          break
          } else {
          v_while4 = ToBoolean(OBJ(try_8))
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
/* The go function for: self_print(self:Handle) [status=1] */
func (self *ClaireHandle) SelfPrint () EID { 
  var Result EID
  PRINC("try ")
  Result = Core.F_CALL(C_print,ARGS(self.Arg.ToEID()))
  if !ErrorIn(Result) {
  PRINC(" ")
  Result = F_lbreak_integer(0)
  if !ErrorIn(Result) {
  PRINC("catch ")
  Result = Core.F_CALL(C_print,ARGS(self.Test.ToEID()))
  if !ErrorIn(Result) {
  PRINC(" ")
  Result = Core.F_CALL(C_print,ARGS(self.Other.ToEID()))
  }}}
  if !ErrorIn(Result) {
  { 
    var va_arg1 *Core.PrettyPrinter
    var va_arg2 int
    va_arg1 = Core.C_pretty
    va_arg2 = (Core.C_pretty.Index-2)
    va_arg1.Index = va_arg2
    Result = EID{C__INT,IVAL(va_arg2)}
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
// notice that return_error should be called return_exception since they travel through interpreted
// not a problem at compile time since return_exceptions are handled with break(x)
/* The go function for: self_eval(self:Handle) [status=1] */
func (self *ClaireHandle) SelfEval () EID { 
  var Result EID
  { 
    h_index := ClEnv.Index
    h_base := ClEnv.Base
    Result = EVAL(self.Arg)
    if ErrorIn(Result){ 
      ClEnv.Index = h_index
      ClEnv.Base = h_base
      { var e *ClaireException = ClEnv.Exception_I
        { var x *ClaireClass = ToClass(self.Test)
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
/* The go function for: self_print(self:Construct) [status=1] */
func (self *Construct) SelfPrint () EID { 
  var Result EID
  { var _Zl int = Core.C_pretty.Index
    { var arg_1 *ClaireString
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
    if ((self.Isa.IsIn(C_List) == CTRUE) || 
        (self.Isa.IsIn(C_Set) == CTRUE)) { 
      { var g0200 *Construct = self
        { var _Zt *ClaireAny = Core.F_get_property(C_of,ToObject(g0200.Id()))
          if (_Zt != CNULL) { 
            if (Equal(_Zt,CEMPTY.Id()) != CTRUE) { 
              PRINC("<")
              Result = Core.F_CALL(C_print,ARGS(_Zt.ToEID()))
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
      { var g0201 *Map = To_Map(self.Id())
        PRINC("<")
        Result = Core.F_print_any(g0201.Domain.Id())
        if !ErrorIn(Result) {
        PRINC(",")
        Result = Core.F_print_any(g0201.Of.Id())
        if !ErrorIn(Result) {
        PRINC(">")
        Result = EVOID
        }}
        } 
      } else {
      Result = EID{CFALSE.Id(),0}
      } 
    if !ErrorIn(Result) {
    PRINC("(")
    F_set_level_void()
    Result = F_Language_printbox_list2(self.Args)
    if !ErrorIn(Result) {
    PRINC(")")
    Result = EVOID
    }}
    if !ErrorIn(Result) {
    { 
      var va_arg1 *Core.PrettyPrinter
      var va_arg2 int
      va_arg1 = Core.C_pretty
      va_arg2 = _Zl
      va_arg1.Index = va_arg2
      Result = EID{C__INT,IVAL(va_arg2)}
      } 
    }
    } 
  return Result} 

// The EID go function for: self_print @ Construct (throw: true) 
func E_self_print_Construct_Language (self EID) EID { 
  return To_Construct(OBJ(self)).SelfPrint( )} 

// constructors: how to create a list, a set, a tuple or an array
// note that the constructor is typed
// CLAIRE4: must build the list with the proper type from the beginning, so that Srange is correct
/* The go function for: self_eval(self:List) [status=1] */
func (self *List) SelfEval () EID { 
  var Result EID
  { var type_ask *ClaireBoolean = MakeBoolean((self.Of.Id() == CNULL)).Not
    { var n int = self.Args.Length()
      if (type_ask == CTRUE) { 
        { var l *ClaireList = CreateList(self.Of,n)
          { var i int = 1
            { var g0202 int = n
              Result= EID{CFALSE.Id(),0}
              for (i <= g0202) { 
                var loop_1 EID
                _ = loop_1
                { 
                { 
                  var arg_2 EID
                  arg_2 = EVAL(self.Args.At(i-1))
                  if ErrorIn(arg_2) {loop_1 = arg_2
                  } else {
                  loop_1 = l.WriteEID(i,arg_2)}
                  } 
                if ErrorIn(loop_1) {Result = loop_1
                break
                } else {
                i = (i+1)
                }
                } 
              }
              } 
            } 
          if !ErrorIn(Result) {
          Result = EID{l.Id(),0}
          }
          } 
        } else {
        { var l *ClaireList = CreateList(ToType(CEMPTY.Id()),n)
          { var i int = 1
            { var g0203 int = n
              Result= EID{CFALSE.Id(),0}
              for (i <= g0203) { 
                var loop_3 EID
                _ = loop_3
                { 
                { var arg_4 *ClaireAny
                  var try_5 EID
                  try_5 = EVAL(self.Args.At(i-1))
                  if ErrorIn(try_5) {loop_3 = try_5
                  } else {
                  arg_4 = ANY(try_5)
                  loop_3 = ToArray(l.Id()).NthPut(i,arg_4).ToEID()
                  }
                  } 
                if ErrorIn(loop_3) {Result = loop_3
                break
                } else {
                i = (i+1)
                }
                } 
              }
              } 
            } 
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
/* The go function for: self_eval(self:Set) [status=1] */
func (self *Set) SelfEval () EID { 
  var Result EID
  { var type_ask *ClaireBoolean = MakeBoolean((self.Of.Id() == CNULL)).Not
    { var n int = self.Args.Length()
      if (type_ask == CTRUE) { 
        { var l *ClaireSet = self.Of.EmptySet()
          { var i int = 1
            { var g0204 int = n
              Result= EID{CFALSE.Id(),0}
              for (i <= g0204) { 
                var loop_1 EID
                _ = loop_1
                { 
                { var arg_2 *ClaireAny
                  var try_3 EID
                  try_3 = EVAL(self.Args.At(i-1))
                  if ErrorIn(try_3) {loop_1 = try_3
                  } else {
                  arg_2 = ANY(try_3)
                  loop_1 = EID{l.AddFast(arg_2).Id(),0}
                  }
                  } 
                if ErrorIn(loop_1) {Result = loop_1
                break
                } else {
                i = (i+1)
                }
                } 
              }
              } 
            } 
          if !ErrorIn(Result) {
          Result = EID{l.Id(),0}
          }
          } 
        } else {
        { var l *ClaireSet = CEMPTY.EmptySet()
          { var i int = 1
            { var g0205 int = n
              Result= EID{CFALSE.Id(),0}
              for (i <= g0205) { 
                var loop_4 EID
                _ = loop_4
                { 
                { var arg_5 *ClaireAny
                  var try_6 EID
                  try_6 = EVAL(self.Args.At(i-1))
                  if ErrorIn(try_6) {loop_4 = try_6
                  } else {
                  arg_5 = ANY(try_6)
                  loop_4 = EID{l.AddFast(arg_5).Id(),0}
                  }
                  } 
                if ErrorIn(loop_4) {Result = loop_4
                break
                } else {
                i = (i+1)
                }
                } 
              }
              } 
            } 
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
/* The go function for: self_eval(self:Tuple) [status=1] */
func (self *Tuple) SelfEval () EID { 
  var Result EID
  { var arg_1 *ClaireList
    var try_2 EID
    { 
      var v_list2 *ClaireList
      var x *ClaireAny
      var v_local2 *ClaireAny
      v_list2 = self.Args
      try_2 = EID{CreateList(ToType(CEMPTY.Id()),v_list2.Length()).Id(),0}
      for CLcount := 0; CLcount < v_list2.Length(); CLcount++{ 
        x = v_list2.At(CLcount)
        var try_3 EID
        try_3 = EVAL(x)
        if ErrorIn(try_3) {try_2 = try_3
        break
        } else {
        v_local2 = ANY(try_3)
        ToList(OBJ(try_2)).PutAt(CLcount,v_local2)
        } 
      }
      } 
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
/* The go function for: self_eval(self:Array) [status=1] */
func (self *Array) SelfEval () EID { 
  var Result EID
  { var type_ask *ClaireBoolean = MakeBoolean((self.Of.Id() == CNULL)).Not
    { var n int = self.Args.Length()
      if (type_ask == CTRUE) { 
        { var l *ClaireList = CreateList(self.Of,n)
          { var i int = 1
            { var g0206 int = n
              Result= EID{CFALSE.Id(),0}
              for (i <= g0206) { 
                var loop_1 EID
                _ = loop_1
                { 
                { 
                  var arg_2 EID
                  arg_2 = EVAL(self.Args.At(i-1))
                  if ErrorIn(arg_2) {loop_1 = arg_2
                  } else {
                  loop_1 = l.WriteEID(i,arg_2)}
                  } 
                if ErrorIn(loop_1) {Result = loop_1
                break
                } else {
                i = (i+1)
                }
                } 
              }
              } 
            } 
          if !ErrorIn(Result) {
          Result = EID{l.Array_I().Id(),0}
          }
          } 
        } else {
        { var l *ClaireList = CreateList(ToType(CEMPTY.Id()),n)
          { var i int = 1
            { var g0207 int = n
              Result= EID{CFALSE.Id(),0}
              for (i <= g0207) { 
                var loop_3 EID
                _ = loop_3
                { 
                { var arg_4 *ClaireAny
                  var try_5 EID
                  try_5 = EVAL(self.Args.At(i-1))
                  if ErrorIn(try_5) {loop_3 = try_5
                  } else {
                  arg_4 = ANY(try_5)
                  loop_3 = ToArray(l.Id()).NthPut(i,arg_4).ToEID()
                  }
                  } 
                if ErrorIn(loop_3) {Result = loop_3
                break
                } else {
                i = (i+1)
                }
                } 
              }
              } 
            } 
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
/* The go function for: self_eval(self:Map) [status=1] */
func (self *Map) SelfEval () EID { 
  var Result EID
  { var m *ClaireMapSet = self.Domain.Map_I(self.Of)
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
        if (x.Isa.IsIn(C_pair) == CTRUE) { 
          { var g0208 *ClairePair = ToPair(x)
            { var arg_2 *ClaireAny
              var try_4 EID
              try_4 = EVAL(g0208.First)
              if ErrorIn(try_4) {loop_1 = try_4
              } else {
              arg_2 = ANY(try_4)
              { var arg_3 *ClaireAny
                var try_5 EID
                try_5 = EVAL(g0208.Second)
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
        if ErrorIn(loop_1) {Result = loop_1
        break
        } else {
        }
        } 
      } 
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
/* The go function for: self_eval(self:Macro) [status=1] */
func (self *Macro) SelfEval () EID { 
  var Result EID
  { var arg_1 *ClaireAny
    var try_2 EID
    try_2 = Core.F_CALL(C_macroexpand,ARGS(EID{self.Id(),0}))
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
/* The go function for: self_eval(self:Error) [status=1] */
func (self *Error) SelfEval () EID { 
  var Result EID
  if ((F_boolean_I_any(self.Args.Id()).Id() != CTRUE.Id()) || 
      (C_string.Id() != self.Args.At(0).Isa.Id())) { 
    Result = ToException(Core.C_general_error.Make(MakeString("Syntax error: ~S").Id(),MakeConstantList(self.Id()).Id())).Close()
    } else {
    Result = EID{CFALSE.Id(),0}
    } 
  if !ErrorIn(Result) {
  { var x *Core.GeneralError = Core.ToGeneralError(new(Core.GeneralError).Is(Core.C_general_error))
    { 
      var va_arg1 *Core.GeneralError
      var va_arg2 *ClaireAny
      va_arg1 = x
      var try_1 EID
      try_1 = Core.F_car_list(self.Args)
      if ErrorIn(try_1) {Result = try_1
      } else {
      va_arg2 = ANY(try_1)
      va_arg1.Cause = va_arg2
      Result = va_arg2.ToEID()
      }
      } 
    if !ErrorIn(Result) {
    { 
      var va_arg1 *Core.GeneralError
      var va_arg2 *ClaireAny
      va_arg1 = x
      var try_2 EID
      { 
        var v_list3 *ClaireList
        var x *ClaireAny
        var v_local3 *ClaireAny
        var try_3 EID
        try_3 = self.Args.Cdr()
        if ErrorIn(try_3) {try_2 = try_3
        } else {
        v_list3 = ToList(OBJ(try_3))
        try_2 = EID{CreateList(ToType(CEMPTY.Id()),v_list3.Length()).Id(),0}
        for CLcount := 0; CLcount < v_list3.Length(); CLcount++{ 
          x = v_list3.At(CLcount)
          var try_4 EID
          try_4 = EVAL(x)
          if ErrorIn(try_4) {try_2 = try_4
          break
          } else {
          v_local3 = ANY(try_4)
          ToList(OBJ(try_2)).PutAt(CLcount,v_local3)
          } 
        }}
        } 
      if ErrorIn(try_2) {Result = try_2
      } else {
      va_arg2 = ANY(try_2)
      va_arg1.Arg = va_arg2
      Result = va_arg2.ToEID()
      }
      } 
    if !ErrorIn(Result) {
    Result = ToException(x.Id()).Close()
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
/* The go function for: self_eval(self:Printf) [status=1] */
func (self *Printf) SelfEval () EID { 
  var Result EID
  { var l *ClaireList = self.Args
    { var s *ClaireAny = l.At(0)
      if (C_string.Id() != s.Isa.Id()) { 
        Result = ToException(Core.C_general_error.Make(MakeString("[102] the first argument in ~S must be a string").Id(),MakeConstantList(self.Id()).Id())).Close()
        } else {
        { var i int = 2
          { var n int = F_get_string(ToString(s),'~')
            Result= EID{CFALSE.Id(),0}
            for (n != 0) { 
              var loop_1 EID
              _ = loop_1
              { var m *ClaireAny
                var try_2 EID
                try_2 = Core.F_CALL(C_nth,ARGS(s.ToEID(),EID{C__INT,IVAL((n+1))}))
                if ErrorIn(try_2) {loop_1 = try_2
                } else {
                m = ANY(try_2)
                if (i > l.Length()) { 
                  loop_1 = ToException(Core.C_general_error.Make(MakeString("[103] not enough arguments in ~S").Id(),MakeConstantList(self.Id()).Id())).Close()
                  } else {
                  loop_1 = EID{CFALSE.Id(),0}
                  } 
                if ErrorIn(loop_1) {Result = loop_1
                break
                } else {
                if (n > 1) { 
                  F_princ_string(F_substring_string(ToString(s),1,(n-1)))
                  } 
                if ('A' == ToChar(m).Value) { 
                  { var arg_3 *ClaireAny
                    var try_4 EID
                    try_4 = EVAL(l.At(i-1))
                    if ErrorIn(try_4) {loop_1 = try_4
                    } else {
                    arg_3 = ANY(try_4)
                    loop_1 = Core.F_CALL(C_princ,ARGS(arg_3.ToEID()))
                    }
                    } 
                  }  else if ('S' == ToChar(m).Value) { 
                  { var arg_5 *ClaireAny
                    var try_6 EID
                    try_6 = EVAL(l.At(i-1))
                    if ErrorIn(try_6) {loop_1 = try_6
                    } else {
                    arg_5 = ANY(try_6)
                    loop_1 = Core.F_CALL(C_print,ARGS(arg_5.ToEID()))
                    }
                    } 
                  }  else if ('F' == ToChar(m).Value) { 
                  { var fv *ClaireAny
                    var try_7 EID
                    try_7 = EVAL(l.At(i-1))
                    if ErrorIn(try_7) {loop_1 = try_7
                    } else {
                    fv = ANY(try_7)
                    { var p_Z *ClaireBoolean = CFALSE
                      { var j int
                        var try_8 EID
                        { var arg_9 int
                          var try_10 EID
                          { var arg_11 rune
                            var try_12 EID
                            try_12 = Core.F_nth_get_string(ToString(s),(n+2),(n+2))
                            if ErrorIn(try_12) {try_10 = try_12
                            } else {
                            arg_11 = CHAR(try_12)
                            try_10 = EID{C__INT,IVAL(int(arg_11))}
                            }
                            } 
                          if ErrorIn(try_10) {try_8 = try_10
                          } else {
                          arg_9 = INT(try_10)
                          try_8 = EID{C__INT,IVAL((arg_9-48))}
                          }
                          } 
                        if ErrorIn(try_8) {loop_1 = try_8
                        } else {
                        j = INT(try_8)
                        var g0210I *ClaireBoolean
                        var try_13 EID
                        { var arg_14 *ClaireAny
                          var try_15 EID
                          try_15 = Core.F_CALL(C_nth,ARGS(s.ToEID(),EID{C__INT,IVAL((n+2))}))
                          if ErrorIn(try_15) {try_13 = try_15
                          } else {
                          arg_14 = ANY(try_15)
                          try_13 = EID{Equal(MakeChar('%').Id(),arg_14).Id(),0}
                          }
                          } 
                        if ErrorIn(try_13) {loop_1 = try_13
                        } else {
                        g0210I = ToBoolean(OBJ(try_13))
                        if (g0210I == CTRUE) { 
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
                        if ErrorIn(loop_1) {Result = loop_1
                        break
                        } else {
                        var g0211I *ClaireBoolean
                        var try_16 EID
                        { 
                          var v_and12 *ClaireBoolean
                          
                          v_and12 = p_Z.Not
                          if (v_and12 == CFALSE) {try_16 = EID{CFALSE.Id(),0}
                          } else { 
                            var try_17 EID
                            { var arg_18 *ClaireAny
                              var try_19 EID
                              try_19 = Core.F_CALL(C_nth,ARGS(s.ToEID(),EID{C__INT,IVAL((n+3))}))
                              if ErrorIn(try_19) {try_17 = try_19
                              } else {
                              arg_18 = ANY(try_19)
                              try_17 = EID{Equal(MakeChar('%').Id(),arg_18).Id(),0}
                              }
                              } 
                            if ErrorIn(try_17) {try_16 = try_17
                            } else {
                            v_and12 = ToBoolean(OBJ(try_17))
                            if (v_and12 == CFALSE) {try_16 = EID{CFALSE.Id(),0}
                            } else { 
                              try_16 = EID{CTRUE.Id(),0}} 
                            } 
                          }
                          } 
                        if ErrorIn(try_16) {loop_1 = try_16
                        } else {
                        g0211I = ToBoolean(OBJ(try_16))
                        if (g0211I == CTRUE) { 
                          p_Z = CTRUE
                          fv = ANY(Core.F_CALL(ToProperty(C__star.Id()),ARGS(fv.ToEID(),EID{C__FLOAT,FVAL(100)})))
                          n = (n+1)
                          loop_1 = EID{C__INT,IVAL(n)}
                          } else {
                          loop_1 = EID{CFALSE.Id(),0}
                          } 
                        }
                        if ErrorIn(loop_1) {Result = loop_1
                        break
                        } else {
                        loop_1 = Core.F_CALL(Core.C_mClaire_printFDigit,ARGS(fv.ToEID(),EID{C__INT,IVAL(j)}))
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
                if ErrorIn(loop_1) {Result = loop_1
                break
                } else {
                i = (i+1)
                s = (F_substring_string(ToString(s),(n+2),1000)).Id()
                var try_20 EID
                try_20 = Core.F_CALL(C_get,ARGS(s.ToEID(),EID{C__CHAR,CVAL('~')}))
                if ErrorIn(try_20) {loop_1 = try_20
                Result = try_20
                break
                } else {
                n = INT(try_20)
                loop_1 = EID{C__INT,IVAL(n)}
                }}}
                }
                } 
              if ErrorIn(loop_1) {Result = loop_1
              break
              } else {
              } 
            }
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
// CLAIRE4: self_eval is defined once for all, hence extended
/* The go function for: self_eval(self:Trace) [status=1] */
func (self *Trace) SelfEval () EID { 
  var Result EID
  if (self.Args.Length() == 0) { 
    Result = IfThenElse((ClEnv.Trace_I == 0),
      MakeString("inactive").Id(),
      MakeString("active").Id()).ToEID()
    } else {
    { var a *ClaireList = self.Args
      { var l *ClaireList
        var try_1 EID
        { 
          var v_list4 *ClaireList
          var x *ClaireAny
          var v_local4 *ClaireAny
          v_list4 = a
          try_1 = EID{CreateList(ToType(CEMPTY.Id()),v_list4.Length()).Id(),0}
          for CLcount := 0; CLcount < v_list4.Length(); CLcount++{ 
            x = v_list4.At(CLcount)
            var try_2 EID
            try_2 = EVAL(x)
            if ErrorIn(try_2) {try_1 = try_2
            break
            } else {
            v_local4 = ANY(try_2)
            ToList(OBJ(try_1)).PutAt(CLcount,v_local4)
            } 
          }
          } 
        if ErrorIn(try_1) {Result = try_1
        } else {
        l = ToList(OBJ(try_1))
        { var i *ClaireAny = l.At(0)
          { var a2 *ClaireAny
            if (a.Length() > 1) { 
              a2 = a.At(1)
              } else {
              a2 = CFALSE.Id()
              } 
            if (a.Length() == 1) { 
              { var a1 *ClaireAny
                var try_3 EID
                try_3 = EVAL(a.At(0))
                if ErrorIn(try_3) {Result = try_3
                } else {
                a1 = ANY(try_3)
                { var p *ClaireProperty = C_iClaire_trace_on
                  if (p.Restrictions.Length() != 0) { 
                    if (ClEnv.Trace_I == 0) { 
                      ClEnv.Trace_I = 1
                      } 
                    Result = Core.F_call_property(p,ToList(MakeConstantList(a1).Id()))
                    } else {
                    Result = EID{CFALSE.Id(),0}
                    } 
                  } 
                }
                } 
              } else {
              var g0213I *ClaireBoolean
              { 
                var v_and7 *ClaireBoolean
                
                v_and7 = Equal(C_string.Id(),a2.Isa.Id())
                if (v_and7 == CFALSE) {g0213I = CFALSE
                } else { 
                  if (C_integer.Id() == i.Isa.Id()) { 
                    { var g0212 int = ToInteger(i).Value
                      v_and7 = Core.F__inf_equal_integer(g0212,ClEnv.Verbose)
                      } 
                    } else {
                    v_and7 = CFALSE
                    } 
                  if (v_and7 == CFALSE) {g0213I = CFALSE
                  } else { 
                    g0213I = CTRUE} 
                  } 
                } 
              if (g0213I == CTRUE) { 
                { var p *ClaireAny = Core.F_get_property(C_ctrace,ToObject(ClEnv.Id()))
                  if (p != CNULL) { 
                    p = ToPort(p).UseAsOutput().Id()
                    } 
                  Result = Core.F_format_string(ToString(a2),l.Skip(2))
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
/* The go function for: self_eval(self:Assert) [status=1] */
func (self *Assert) SelfEval () EID { 
  var Result EID
  { var a *ClaireList = self.Args
    var g0214I *ClaireBoolean
    var try_1 EID
    { 
      var v_and2 *ClaireBoolean
      
      v_and2 = Core.F__sup_integer(a.Length(),0)
      if (v_and2 == CFALSE) {try_1 = EID{CFALSE.Id(),0}
      } else { 
        v_and2 = Core.F_known_ask_any(Core.F_get_property(C_ctrace,ToObject(ClEnv.Id())))
        if (v_and2 == CFALSE) {try_1 = EID{CFALSE.Id(),0}
        } else { 
          var try_2 EID
          { var arg_3 *ClaireBoolean
            var try_4 EID
            { var arg_5 *ClaireAny
              var try_6 EID
              try_6 = EVAL(a.At(0))
              if ErrorIn(try_6) {try_4 = try_6
              } else {
              arg_5 = ANY(try_6)
              try_4 = EID{F_boolean_I_any(arg_5).Id(),0}
              }
              } 
            if ErrorIn(try_4) {try_2 = try_4
            } else {
            arg_3 = ToBoolean(OBJ(try_4))
            try_2 = EID{Core.F__I_equal_any(arg_3.Id(),CTRUE.Id()).Id(),0}
            }
            } 
          if ErrorIn(try_2) {try_1 = try_2
          } else {
          v_and2 = ToBoolean(OBJ(try_2))
          if (v_and2 == CFALSE) {try_1 = EID{CFALSE.Id(),0}
          } else { 
            try_1 = EID{CTRUE.Id(),0}} 
          } 
        } 
      }
      } 
    if ErrorIn(try_1) {Result = try_1
    } else {
    g0214I = ToBoolean(OBJ(try_1))
    if (g0214I == CTRUE) { 
      { var p *ClairePort = ClEnv.Ctrace.UseAsOutput()
        Result = Core.F_print_any((self.External).Id())
        if !ErrorIn(Result) {
        PRINC(",line=")
        F_princ_integer(self.Index)
        PRINC(": (ASSERT) ")
        Result = Core.F_CALL(C_print,ARGS(a.At(0).ToEID()))
        if !ErrorIn(Result) {
        PRINC("\n")
        Result = EVOID
        }}
        if !ErrorIn(Result) {
        p.UseAsOutput()
        if (ClEnv.Debug_I >= 0) { 
          Result = ToException(Core.C_general_error.Make(MakeString("Assertion Violation").Id(),CNIL.Id())).Close()
          } else {
          Result = EID{CFALSE.Id(),0}
          } 
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

/* The go function for: self_eval(self:Branch) [status=1] */
func (self *Branch) SelfEval () EID { 
  var Result EID
  if (self.Args.Length() != 1) { 
    Result = ToException(Core.C_general_error.Make(MakeString("[104] Syntax error with ~S (one arg. expected)").Id(),MakeConstantList(self.Id()).Id())).Close()
    } else {
    Result = EID{CFALSE.Id(),0}
    } 
  if !ErrorIn(Result) {
  { 
    h_index := ClEnv.Index
    h_base := ClEnv.Base
    F_world_push()
    var g0215I *ClaireBoolean
    var try_1 EID
    { var arg_2 *ClaireAny
      var try_3 EID
      try_3 = EVAL(self.Args.At(0))
      if ErrorIn(try_3) {try_1 = try_3
      } else {
      arg_2 = ANY(try_3)
      try_1 = EID{Core.F__I_equal_any(arg_2,CFALSE.Id()).Id(),0}
      }
      } 
    if ErrorIn(try_1) {Result = try_1
    } else {
    g0215I = ToBoolean(OBJ(try_1))
    if (g0215I == CTRUE) { 
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