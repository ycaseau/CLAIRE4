/***** CLAIRE Compilation of file /Users/ycaseau/Dropbox/src/clairev4.0/src/meta/call.cl 
         [version 4.0.01 / safety 5] Saturday 10-30-2021 *****/

package Language
import (_ "fmt"
	. "Kernel"
	"Core"
)

//-------- dumb function to prevent import errors --------
func import_g0119() { 
    _ = Core.It
    } 
  
  
//+-------------------------------------------------------------+
//| CLAIRE                                                      |
//| call.cl                                                     |
//| Copyright (C) 1994 - 2021 Yves Caseau. All Rights Reserved  |
//| cf. copyright info in file object.cl: about()               |
//+-------------------------------------------------------------+
// -----------------------------------------------------------------
// This file holds the definition of functional calls in CLAIRE
// -----------------------------------------------------------------`
// *********************************************************************
// * Contents                                                          *
// *      Part 1: the basic object messages                            *
// *      Part 2: Basic structures                                     *
// *      Part 3: Specialized structures                               *
// *      Part 4: Functions on instructions                            *
// *********************************************************************
// *********************************************************************
// *      Part 1: the basic object messages                            *
// *********************************************************************
// contains the last message that was evaluated
// messages in CLAIRE are called calls --------------------------------
// Note that a Call* is a Call, but the * is a reader's note that allows precedence combining
// These are special (x op y) Calls that can be combined !
// syntactic mark: slot call (x.s)
/* {1} OPT.The go function for: self_print(self:Call) [] */
func (self *Call ) SelfPrint () EID { 
    var Result EID 
    /* Let:2 */{ 
      var _Zl int  = Core.C_pretty.Index
      /* noccur = 1 */
      /* Let:3 */{ 
        var _Zs *ClaireProperty   = self.Selector
        /* noccur = 11 */
        /* Let:4 */{ 
          var _Za *ClaireList   = self.Args
          /* noccur = 44 */
          if ((_Zs.Isa.IsIn(C_operation) == CTRUE) && 
              (_Za.Length() == 2)) /* If:5 */{ 
            Core.C_pretty.Index = (Core.C_pretty.Index+2)
            Result = F_printe_any(_Za.At(1-1),_Zs)
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC(" ")
            Result = Core.F_print_any(_Zs.Id())
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC(" ")
            Result = F_lbreak_void()
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            Result = F_printe_any(_Za.At(2-1),_Zs)
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC("")
            Result = EVOID
            }}}}
            /* If!5 */}  else if (_Zs.Id() == C_nth.Id()) /* If:5 */{ 
            if (_Za.Length() == 3) /* If:6 */{ 
              Result = F_printexp_any(_Za.At(1-1),CFALSE)
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              PRINC("[")
              Result = Core.F_CALL(C_print,ARGS(_Za.At(2-1).ToEID()))
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              PRINC(",")
              Result = Core.F_CALL(C_print,ARGS(_Za.At(3-1).ToEID()))
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              PRINC("]")
              Result = EVOID
              }}}
              /* If!6 */}  else if (_Za.Length() == 1) /* If:6 */{ 
              Result = F_printexp_any(_Za.At(1-1),CFALSE)
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              PRINC("[]")
              Result = EVOID
              }
              } else {
              Result = F_printexp_any(_Za.At(1-1),CFALSE)
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              PRINC("[")
              if (_Za.Length() == 2) /* If:7 */{ 
                Result = Core.F_CALL(C_print,ARGS(_Za.At(2-1).ToEID()))
                } else {
                Result = EID{CFALSE.Id(),0}
                /* If-7 */} 
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              PRINC("]")
              Result = EVOID
              }}
              /* If-6 */} 
            /* If!5 */}  else if ((_Zs.Id() == C_nth_equal.Id()) && 
              (_Za.Length() >= 3)) /* If:5 */{ 
            /* Let:6 */{ 
              var a *ClaireAny   = _Za.At(3-1)
              /* noccur = 7 */
              /* Let:7 */{ 
                var o *ClaireAny  
                /* noccur = 2 */
                if (a.Isa.IsIn(C_Call) == CTRUE) /* If:8 */{ 
                  /* Let:9 */{ 
                    var g0121 *Call   = To_Call(a)
                    /* noccur = 1 */
                    o = g0121.Selector.Id()
                    /* Let-9 */} 
                  } else {
                  o = CFALSE.Id()
                  /* If-8 */} 
                if (_Za.Length() == 4) /* If:8 */{ 
                  Result = F_printexp_any(_Za.At(1-1),CFALSE)
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  PRINC("[")
                  Result = Core.F_CALL(C_print,ARGS(_Za.At(2-1).ToEID()))
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  PRINC(",")
                  Result = Core.F_CALL(C_print,ARGS(a.ToEID()))
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  PRINC("] := ")
                  Result = F_lbreak_integer(2)
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  Result = Core.F_CALL(C_print,ARGS(_Za.At(4-1).ToEID()))
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  PRINC("")
                  Result = EVOID
                  }}}}}
                  } else {
                  var g0125I *ClaireBoolean  
                  /* Let:9 */{ 
                    var g0126UU *ClaireAny  
                    /* noccur = 1 */
                    if (a.Isa.IsIn(C_Call) == CTRUE) /* If:10 */{ 
                      /* Let:11 */{ 
                        var g0122 *Call   = To_Call(a)
                        /* noccur = 1 */
                        g0126UU = g0122.Args.At(1-1)
                        /* Let-11 */} 
                      } else {
                      g0126UU = CFALSE.Id()
                      /* If-10 */} 
                    g0125I = F_sugar_ask_any(_Za.At(1-1),_Za.At(2-1),o,g0126UU)
                    /* Let-9 */} 
                  if (g0125I == CTRUE) /* If:9 */{ 
                    Result = Core.F_CALL(C_print,ARGS(_Za.At(1-1).ToEID()))
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    PRINC("[")
                    Result = Core.F_CALL(C_print,ARGS(_Za.At(2-1).ToEID()))
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    PRINC("] :")
                    Result = Core.F_CALL(C_print,ARGS(o.ToEID()))
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    PRINC(" ")
                    Result = F_lbreak_integer(2)
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    Result = Core.F_CALL(C_print,ARGS(ToList(OBJ(Core.F_CALL(C_args,ARGS(a.ToEID())))).At(2-1).ToEID()))
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    PRINC("")
                    Result = EVOID
                    }}}}}
                    } else {
                    Result = Core.F_CALL(C_print,ARGS(_Za.At(1-1).ToEID()))
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    PRINC("[")
                    Result = Core.F_CALL(C_print,ARGS(_Za.At(2-1).ToEID()))
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    PRINC("] := ")
                    Result = F_lbreak_integer(2)
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    Result = Core.F_CALL(C_print,ARGS(a.ToEID()))
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    PRINC("")
                    Result = EVOID
                    }}}}
                    /* If-9 */} 
                  /* If-8 */} 
                /* Let-7 */} 
              /* Let-6 */} 
            /* If!5 */}  else if ((_Zs.Id() == C_Language_assign.Id()) && 
              (_Za.At(1-1).Isa.IsIn(C_property) == CTRUE)) /* If:5 */{ 
            /* Let:6 */{ 
              var a *ClaireAny   = _Za.At(3-1)
              /* noccur = 5 */
              /* Let:7 */{ 
                var o *ClaireAny  
                /* noccur = 2 */
                if (a.Isa.IsIn(C_Call) == CTRUE) /* If:8 */{ 
                  /* Let:9 */{ 
                    var g0123 *Call   = To_Call(a)
                    /* noccur = 1 */
                    o = g0123.Selector.Id()
                    /* Let-9 */} 
                  } else {
                  o = CFALSE.Id()
                  /* If-8 */} 
                var g0127I *ClaireBoolean  
                /* Let:8 */{ 
                  var g0128UU *ClaireAny  
                  /* noccur = 1 */
                  if (a.Isa.IsIn(C_Call) == CTRUE) /* If:9 */{ 
                    /* Let:10 */{ 
                      var g0124 *Call   = To_Call(a)
                      /* noccur = 1 */
                      g0128UU = g0124.Args.At(1-1)
                      /* Let-10 */} 
                    } else {
                    g0128UU = CFALSE.Id()
                    /* If-9 */} 
                  g0127I = F_sugar_ask_any(_Za.At(1-1),_Za.At(2-1),o,g0128UU)
                  /* Let-8 */} 
                if (g0127I == CTRUE) /* If:8 */{ 
                  Result = Core.F_CALL(C_print,ARGS(_Za.At(1-1).ToEID()))
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  PRINC("(")
                  Result = Core.F_CALL(C_print,ARGS(_Za.At(2-1).ToEID()))
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  PRINC(") :")
                  Result = Core.F_CALL(C_print,ARGS(o.ToEID()))
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  PRINC(" ")
                  Result = F_lbreak_integer(2)
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  Result = Core.F_CALL(C_print,ARGS(ToList(OBJ(Core.F_CALL(C_args,ARGS(a.ToEID())))).At(2-1).ToEID()))
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  PRINC("")
                  Result = EVOID
                  }}}}}
                  } else {
                  Result = Core.F_CALL(C_print,ARGS(_Za.At(1-1).ToEID()))
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  PRINC("(")
                  Result = Core.F_CALL(C_print,ARGS(_Za.At(2-1).ToEID()))
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  PRINC(") := ")
                  Result = F_lbreak_integer(2)
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  Result = Core.F_CALL(C_print,ARGS(_Za.At(3-1).ToEID()))
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  PRINC("")
                  Result = EVOID
                  }}}}
                  /* If-8 */} 
                /* Let-7 */} 
              /* Let-6 */} 
            /* If!5 */}  else if ((_Zs.Id() == C_add.Id()) && 
              (_Za.At(1-1).Isa.IsIn(C_property) == CTRUE)) /* If:5 */{ 
            Result = Core.F_CALL(C_print,ARGS(_Za.At(1-1).ToEID()))
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC("(")
            Result = Core.F_CALL(C_print,ARGS(_Za.At(2-1).ToEID()))
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC(") :add ")
            Result = F_lbreak_integer(2)
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            Result = Core.F_CALL(C_print,ARGS(_Za.At(3-1).ToEID()))
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC("")
            Result = EVOID
            }}}}
            /* If!5 */}  else if ((_Zs.Id() == C_delete.Id()) && 
              (_Za.At(1-1).Isa.IsIn(C_property) == CTRUE)) /* If:5 */{ 
            Result = Core.F_CALL(C_print,ARGS(_Za.At(1-1).ToEID()))
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC("(")
            Result = Core.F_CALL(C_print,ARGS(_Za.At(2-1).ToEID()))
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC(") :delete ")
            Result = F_lbreak_integer(2)
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            Result = Core.F_CALL(C_print,ARGS(_Za.At(3-1).ToEID()))
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC("")
            Result = EVOID
            }}}}
            /* If!5 */}  else if ((_Za.At(1-1) == ClEnv.Id()) && 
              (_Za.Length() == 1)) /* If:5 */{ 
            Result = Core.F_print_any(_Zs.Id())
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC("()")
            Result = EVOID
            }
            } else {
            Result = Core.F_print_any(_Zs.Id())
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC("(")
            F_set_level_void()
            Result = F_Language_printbox_list2(_Za)
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC(")")
            Result = EVOID
            }}
            /* If-5 */} 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
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
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: self_print @ Call (throw: true) 
func E_self_print_Call_Language (self EID) EID { 
    return /*(sm for self_print @ Call= EID)*/ To_Call(OBJ(self)).SelfPrint( )} 
  
/* {1} OPT.The go function for: self_print(self:Call+) [] */
func (self *Call_plus ) SelfPrint () EID { 
    var Result EID 
    Result = F_printexp_any(self.Args.At(1-1),CTRUE)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(".")
    Result = Core.F_print_any(self.Selector.Id())
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("")
    Result = EVOID
    }}
    return Result} 
  
// The EID go function for: self_print @ Call+ (throw: true) 
func E_self_print_Call_plus_Language (self EID) EID { 
    return /*(sm for self_print @ Call+= EID)*/ To_Call_plus(OBJ(self)).SelfPrint( )} 
  
/* {1} OPT.The go function for: self_eval(self:Call) [] */
func (self *Call ) SelfEval () EID { 
    var Result EID 
    /* Let:2 */{ 
      var start int  = ClEnv.Index
      /* noccur = 6 */
      /* Let:3 */{ 
        var p *ClaireProperty   = self.Selector
        /* noccur = 4 */
        if (ClEnv.Debug_I >= 0) /* If:4 */{ 
          C_iClaire_LastCall.Value = self.Id()
          /* For:5 */{ 
            var x *ClaireAny  
            _ = x
            Result= EID{CFALSE.Id(),0}
            var x_support *ClaireList  
            x_support = self.Args
            x_len := x_support.Length()
            for i_it := 0; i_it < x_len; i_it++ { 
              x = x_support.At(i_it)
              var void_try7 EID 
              _ = void_try7
              /* LetEID:7 */{ 
                var g0129UU EID 
                g0129UU = EVAL(x)
                /* ERROR PROTECTION INSERTED (g0129UU-void_try7) */
                if ErrorIn(g0129UU) {void_try7 = g0129UU
                } else {
                void_try7 = ClEnv.Push(g0129UU)}
                /* LetEID-7 */} 
              /* ERROR PROTECTION INSERTED (void_try7-Result) */
              if ErrorIn(void_try7) {Result = void_try7
              Result = void_try7
              break
              } else {
              }
              /* loop-6 */} 
            /* For-5 */} 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          /* LetE:5 */{ 
            var rx EID 
            rx = Core.F_eval_message_property(p,Core.F_find_which_property(p,start,OWNER(ClEnv.EvalStack[start])),start,CTRUE)
            /* ERROR PROTECTION INSERTED (rx-Result) */
            if ErrorIn(rx) {Result = rx
            } else {
            C_iClaire_LastCall.Value = self.Id()
            Result = rx}
            /* LetE-5 */} 
          }
          } else {
          /* For:5 */{ 
            var x *ClaireAny  
            _ = x
            Result= EID{CFALSE.Id(),0}
            var x_support *ClaireList  
            x_support = self.Args
            x_len := x_support.Length()
            for i_it := 0; i_it < x_len; i_it++ { 
              x = x_support.At(i_it)
              var void_try7 EID 
              _ = void_try7
              /* LetEID:7 */{ 
                var g0130UU EID 
                g0130UU = EVAL(x)
                /* ERROR PROTECTION INSERTED (g0130UU-void_try7) */
                if ErrorIn(g0130UU) {void_try7 = g0130UU
                } else {
                void_try7 = ClEnv.Push(g0130UU)}
                /* LetEID-7 */} 
              /* ERROR PROTECTION INSERTED (void_try7-Result) */
              if ErrorIn(void_try7) {Result = void_try7
              Result = void_try7
              break
              } else {
              }
              /* loop-6 */} 
            /* For-5 */} 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          Result = Core.F_eval_message_property(p,Core.F_find_which_property(p,start,OWNER(ClEnv.EvalStack[start])),start,CTRUE)
          }
          /* If-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: self_eval @ Call (throw: true) 
func E_self_eval_Call (self EID) EID { 
    return /*(sm for self_eval @ Call= EID)*/ To_Call(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Call 
func EVAL_Call (x *ClaireAny) EID { 
     return To_Call(x).SelfEval()} 
  
// read slot : 
/* {1} OPT.The go function for: self_eval(self:Call+) [] */
func (self *Call_plus ) SelfEval () EID { 
    var Result EID 
    /* LetEID:2 */{ 
      var g0131UU EID 
      g0131UU = EVAL(self.Args.At(1-1))
      /* ERROR PROTECTION INSERTED (g0131UU-Result) */
      if ErrorIn(g0131UU) {Result = g0131UU
      } else {
      Result = self.Selector.ReadEID(g0131UU)}
      /* LetEID-2 */} 
    return Result} 
  
// The EID go function for: self_eval @ Call+ (throw: true) 
func E_self_eval_Call_plus (self EID) EID { 
    return /*(sm for self_eval @ Call+= EID)*/ To_Call_plus(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Call+ 
func EVAL_Call_plus (x *ClaireAny) EID { 
     return To_Call_plus(x).SelfEval()} 
  
// dumb : we need to repeat
/* {1} OPT.The go function for: self_eval(self:Call*) [] */
func (self *Call_star ) SelfEval () EID { 
    var Result EID 
    /* Let:2 */{ 
      var start int  = ClEnv.Index
      /* noccur = 6 */
      /* Let:3 */{ 
        var p *ClaireProperty   = self.Selector
        /* noccur = 4 */
        if (ClEnv.Debug_I >= 0) /* If:4 */{ 
          C_iClaire_LastCall.Value = self.Id()
          /* For:5 */{ 
            var x *ClaireAny  
            _ = x
            Result= EID{CFALSE.Id(),0}
            var x_support *ClaireList  
            x_support = self.Args
            x_len := x_support.Length()
            for i_it := 0; i_it < x_len; i_it++ { 
              x = x_support.At(i_it)
              var void_try7 EID 
              _ = void_try7
              /* LetEID:7 */{ 
                var g0132UU EID 
                g0132UU = EVAL(x)
                /* ERROR PROTECTION INSERTED (g0132UU-void_try7) */
                if ErrorIn(g0132UU) {void_try7 = g0132UU
                } else {
                void_try7 = ClEnv.Push(g0132UU)}
                /* LetEID-7 */} 
              /* ERROR PROTECTION INSERTED (void_try7-Result) */
              if ErrorIn(void_try7) {Result = void_try7
              Result = void_try7
              break
              } else {
              }
              /* loop-6 */} 
            /* For-5 */} 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          /* LetE:5 */{ 
            var rx EID 
            rx = Core.F_eval_message_property(p,Core.F_find_which_property(p,start,OWNER(ClEnv.EvalStack[start])),start,CTRUE)
            /* ERROR PROTECTION INSERTED (rx-Result) */
            if ErrorIn(rx) {Result = rx
            } else {
            C_iClaire_LastCall.Value = self.Id()
            Result = rx}
            /* LetE-5 */} 
          }
          } else {
          /* For:5 */{ 
            var x *ClaireAny  
            _ = x
            Result= EID{CFALSE.Id(),0}
            var x_support *ClaireList  
            x_support = self.Args
            x_len := x_support.Length()
            for i_it := 0; i_it < x_len; i_it++ { 
              x = x_support.At(i_it)
              var void_try7 EID 
              _ = void_try7
              /* LetEID:7 */{ 
                var g0133UU EID 
                g0133UU = EVAL(x)
                /* ERROR PROTECTION INSERTED (g0133UU-void_try7) */
                if ErrorIn(g0133UU) {void_try7 = g0133UU
                } else {
                void_try7 = ClEnv.Push(g0133UU)}
                /* LetEID-7 */} 
              /* ERROR PROTECTION INSERTED (void_try7-Result) */
              if ErrorIn(void_try7) {Result = void_try7
              Result = void_try7
              break
              } else {
              }
              /* loop-6 */} 
            /* For-5 */} 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          Result = Core.F_eval_message_property(p,Core.F_find_which_property(p,start,OWNER(ClEnv.EvalStack[start])),start,CTRUE)
          }
          /* If-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: self_eval @ Call* (throw: true) 
func E_self_eval_Call_star (self EID) EID { 
    return /*(sm for self_eval @ Call*= EID)*/ To_Call_star(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Call* 
func EVAL_Call_star (x *ClaireAny) EID { 
     return To_Call_star(x).SelfEval()} 
  
// recursive printing of bicall
/* {1} OPT.The go function for: printe(self:any,s:property) [] */
func F_printe_any (self *ClaireAny ,s *ClaireProperty ) EID { 
    var Result EID 
    var g0135I *ClaireBoolean  
    if (self.Isa.IsIn(C_Call) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0134 *Call   = To_Call(self)
        /* noccur = 2 */
        g0135I = MakeBoolean((g0134.Selector.Isa.IsIn(C_operation) == CTRUE) && (g0134.Args.Length() == 2))
        /* Let-3 */} 
      } else {
      g0135I = CFALSE
      /* If-2 */} 
    if (g0135I == CTRUE) /* If:2 */{ 
      if (CTRUE == CTRUE) /* If:3 */{ 
        PRINC("(")
        Result = Core.F_CALL(C_print,ARGS(self.ToEID()))
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC(")")
        Result = EVOID
        }
        } else {
        Result = F_printexp_any(self,CTRUE)
        /* If-3 */} 
      } else {
      Result = F_printexp_any(self,CTRUE)
      /* If-2 */} 
    return Result} 
  
// The EID go function for: printe @ any (throw: true) 
func E_printe_any (self EID,s EID) EID { 
    return /*(sm for printe @ any= EID)*/ F_printe_any(ANY(self),ToProperty(OBJ(s)) )} 
  
// tells if the sugar :op can be used
// x(x2) = o(a:x(x2), y) =>  x(x2) :o y
/* {1} OPT.The go function for: sugar?(x:any,x2:any,o:any,a:any) [] */
func F_sugar_ask_any (x *ClaireAny ,x2 *ClaireAny ,o *ClaireAny ,a *ClaireAny ) *ClaireBoolean  { 
    // procedure body with s =  
var Result *ClaireBoolean  
    if (o.Isa.IsIn(C_operation) == CTRUE) /* If:2 */{ 
      if (x.Isa.IsIn(C_property) == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var g0137 *ClaireProperty   = ToProperty(x)
          /* noccur = 1 */
          if (a.Isa.IsIn(C_Call) == CTRUE) /* If:5 */{ 
            /* Let:6 */{ 
              var g0138 *Call   = To_Call(a)
              /* noccur = 2 */
              Result = MakeBoolean((g0137.Id() == g0138.Selector.Id()) && (Equal(g0138.Args.At(1-1),x2) == CTRUE))
              /* Let-6 */} 
            } else {
            Result = CFALSE
            /* If-5 */} 
          /* Let-4 */} 
        } else {
        Result = CFALSE
        /* If-3 */} 
      } else {
      Result = CFALSE
      /* If-2 */} 
    return Result} 
  
// The EID go function for: sugar? @ any (throw: false) 
func E_sugar_ask_any (x EID,x2 EID,o EID,a EID) EID { 
    return EID{/*(sm for sugar? @ any= boolean)*/ F_sugar_ask_any(ANY(x),
      ANY(x2),
      ANY(o),
      ANY(a) ).Id(),0}} 
  
// CLAIRE 3.5 code : strange  
//   any (case a (Call  (a.selector = nth & a.args[1] = x & a.args[2] = x2),   any false)))
// *********************************************************************
// *      Part 2: Basic structures                                     *
// *********************************************************************
// ------------------ assignment ---------------------------------------
// <-(var V, arg E) where V is a variable (and therefore NOT a global_variable)
//
// the var slot is filled with a real variable later.
/* {1} OPT.The go function for: self_print(self:Assign) [] */
func (self *Assign ) SelfPrint () EID { 
    var Result EID 
    /* Let:2 */{ 
      var a *ClaireAny   = self.Arg
      /* noccur = 6 */
      /* Let:3 */{ 
        var o *ClaireAny  
        /* noccur = 2 */
        if (a.Isa.IsIn(C_Call) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var g0142 *Call   = To_Call(a)
            /* noccur = 1 */
            o = g0142.Selector.Id()
            /* Let-5 */} 
          } else {
          o = CFALSE.Id()
          /* If-4 */} 
        var g0144I *ClaireBoolean  
        /* Let:4 */{ 
          var g0145UU *ClaireAny  
          /* noccur = 1 */
          if (a.Isa.IsIn(C_Call) == CTRUE) /* If:5 */{ 
            /* Let:6 */{ 
              var g0143 *Call   = To_Call(a)
              /* noccur = 1 */
              g0145UU = g0143.Args.At(1-1)
              /* Let-6 */} 
            } else {
            g0145UU = CFALSE.Id()
            /* If-5 */} 
          g0144I = F_sugar_ask_any(self.ClaireVar,CEMPTY.Id(),o,g0145UU)
          /* Let-4 */} 
        if (g0144I == CTRUE) /* If:4 */{ 
          Result = Core.F_CALL(C_print,ARGS(self.ClaireVar.ToEID()))
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(" :")
          Result = Core.F_CALL(C_print,ARGS(o.ToEID()))
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(" ")
          Result = F_lbreak_integer(2)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          Result = F_printexp_any(ToList(OBJ(Core.F_CALL(C_args,ARGS(a.ToEID())))).At(2-1),CTRUE)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC("")
          Result = EVOID
          }}}}
          } else {
          Result = Core.F_CALL(C_print,ARGS(self.ClaireVar.ToEID()))
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(" := ")
          Result = F_lbreak_integer(2)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          Result = F_printexp_any(a,CTRUE)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC("")
          Result = EVOID
          }}}
          /* If-4 */} 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
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
    return Result} 
  
// The EID go function for: self_print @ Assign (throw: true) 
func E_self_print_Assign_Language (self EID) EID { 
    return /*(sm for self_print @ Assign= EID)*/ To_Assign(OBJ(self)).SelfPrint( )} 
  
/* {1} OPT.The go function for: self_eval(self:Assign) [] */
func (self *Assign ) SelfEval () EID { 
    var Result EID 
    /* LetEID:2 */{ 
      var g0146UU EID 
      g0146UU = EVAL(self.Arg)
      /* ERROR PROTECTION INSERTED (g0146UU-Result) */
      if ErrorIn(g0146UU) {Result = g0146UU
      } else {
      Result = To_Variable(self.ClaireVar).WriteEID(g0146UU)}
      /* LetEID-2 */} 
    return Result} 
  
// The EID go function for: self_eval @ Assign (throw: true) 
func E_self_eval_Assign (self EID) EID { 
    return /*(sm for self_eval @ Assign= EID)*/ To_Assign(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Assign 
func EVAL_Assign (x *ClaireAny) EID { 
     return To_Assign(x).SelfEval()} 
  
// global variables
//
/* {1} OPT.The go function for: self_print(self:Gassign) [] */
func (self *Gassign ) SelfPrint () EID { 
    var Result EID 
    /* Let:2 */{ 
      var a *ClaireAny   = self.Arg
      /* noccur = 6 */
      /* Let:3 */{ 
        var o *ClaireAny  
        /* noccur = 2 */
        if (a.Isa.IsIn(C_Call) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var g0148 *Call   = To_Call(a)
            /* noccur = 1 */
            o = g0148.Selector.Id()
            /* Let-5 */} 
          } else {
          o = CFALSE.Id()
          /* If-4 */} 
        var g0150I *ClaireBoolean  
        /* Let:4 */{ 
          var g0151UU *ClaireAny  
          /* noccur = 1 */
          if (a.Isa.IsIn(C_Call) == CTRUE) /* If:5 */{ 
            /* Let:6 */{ 
              var g0149 *Call   = To_Call(a)
              /* noccur = 1 */
              g0151UU = g0149.Args.At(1-1)
              /* Let-6 */} 
            } else {
            g0151UU = CFALSE.Id()
            /* If-5 */} 
          g0150I = F_sugar_ask_any(self.ClaireVar.Id(),CEMPTY.Id(),o,g0151UU)
          /* Let-4 */} 
        if (g0150I == CTRUE) /* If:4 */{ 
          Result = Core.F_print_any(self.ClaireVar.Id())
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(" :")
          Result = Core.F_CALL(C_print,ARGS(o.ToEID()))
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(" ")
          Result = F_lbreak_integer(2)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          Result = Core.F_CALL(C_print,ARGS(ToList(OBJ(Core.F_CALL(C_args,ARGS(a.ToEID())))).At(2-1).ToEID()))
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC("")
          Result = EVOID
          }}}}
          } else {
          Result = Core.F_print_any(self.ClaireVar.Id())
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(" := ")
          Result = F_lbreak_integer(2)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          Result = Core.F_CALL(C_print,ARGS(a.ToEID()))
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC("")
          Result = EVOID
          }}}
          /* If-4 */} 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
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
    return Result} 
  
// The EID go function for: self_print @ Gassign (throw: true) 
func E_self_print_Gassign_Language (self EID) EID { 
    return /*(sm for self_print @ Gassign= EID)*/ To_Gassign(OBJ(self)).SelfPrint( )} 
  
/* {1} OPT.The go function for: self_eval(self:Gassign) [] */
func (self *Gassign ) SelfEval () EID { 
    var Result EID 
    /* Let:2 */{ 
      var v *Core.GlobalVariable   = self.ClaireVar
      /* noccur = 1 */
      /* Let:3 */{ 
        var g0152UU *ClaireAny  
        /* noccur = 1 */
        var g0152UU_try01534 EID 
        g0152UU_try01534 = EVAL(self.Arg)
        /* ERROR PROTECTION INSERTED (g0152UU-Result) */
        if ErrorIn(g0152UU_try01534) {Result = g0152UU_try01534
        } else {
        g0152UU = ANY(g0152UU_try01534)
        Result = F_write_value_global_variable(v,g0152UU)
        }
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: self_eval @ Gassign (throw: true) 
func E_self_eval_Gassign (self EID) EID { 
    return /*(sm for self_eval @ Gassign= EID)*/ To_Gassign(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Gassign 
func EVAL_Gassign (x *ClaireAny) EID { 
     return To_Gassign(x).SelfEval()} 
  
//--------------- BOOLEAN OPERATIONS ---------------------------------
// "and" is strictly boolean and is based on short-circuit evaluation.
//
/* {1} OPT.The go function for: self_print(self:And) [] */
func (self *And ) SelfPrint () EID { 
    var Result EID 
    PRINC("(")
    Result = F_Language_printbox_list3(self.Args,MakeString(" & "))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(")")
    Result = EVOID
    }
    return Result} 
  
// The EID go function for: self_print @ And (throw: true) 
func E_self_print_And_Language (self EID) EID { 
    return /*(sm for self_print @ And= EID)*/ To_And(OBJ(self)).SelfPrint( )} 
  
/* {1} OPT.The go function for: self_eval(self:And) [] */
func (self *And ) SelfEval () EID { 
    var Result EID 
    /* Let:2 */{ 
      var g0154UU *ClaireAny  
      /* noccur = 1 */
      var g0154UU_try01553 EID 
      /* For:3 */{ 
        var x *ClaireAny  
        _ = x
        g0154UU_try01553= EID{CFALSE.Id(),0}
        var x_support *ClaireList  
        x_support = self.Args
        x_len := x_support.Length()
        for i_it := 0; i_it < x_len; i_it++ { 
          x = x_support.At(i_it)
          var void_try5 EID 
          _ = void_try5
          var g0156I *ClaireBoolean  
          var g0156I_try01575 EID 
          /* Let:5 */{ 
            var g0158UU *ClaireBoolean  
            /* noccur = 1 */
            var g0158UU_try01596 EID 
            /* Let:6 */{ 
              var g0160UU *ClaireAny  
              /* noccur = 1 */
              var g0160UU_try01617 EID 
              g0160UU_try01617 = EVAL(x)
              /* ERROR PROTECTION INSERTED (g0160UU-g0158UU_try01596) */
              if ErrorIn(g0160UU_try01617) {g0158UU_try01596 = g0160UU_try01617
              } else {
              g0160UU = ANY(g0160UU_try01617)
              g0158UU_try01596 = EID{F_boolean_I_any(g0160UU).Id(),0}
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (g0158UU-g0156I_try01575) */
            if ErrorIn(g0158UU_try01596) {g0156I_try01575 = g0158UU_try01596
            } else {
            g0158UU = ToBoolean(OBJ(g0158UU_try01596))
            g0156I_try01575 = EID{Core.F__I_equal_any(g0158UU.Id(),CTRUE.Id()).Id(),0}
            }
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (g0156I-void_try5) */
          if ErrorIn(g0156I_try01575) {void_try5 = g0156I_try01575
          } else {
          g0156I = ToBoolean(OBJ(g0156I_try01575))
          if (g0156I == CTRUE) /* If:5 */{ 
             /*v = g0154UU_try01553, s =EID*/
g0154UU_try01553 = EID{CTRUE.Id(),0}
            break
            } else {
            void_try5 = EID{CFALSE.Id(),0}
            /* If-5 */} 
          }
          /* ERROR PROTECTION INSERTED (void_try5-g0154UU_try01553) */
          if ErrorIn(void_try5) {g0154UU_try01553 = void_try5
          g0154UU_try01553 = void_try5
          break
          } else {
          }
          /* loop-4 */} 
        /* For-3 */} 
      /* ERROR PROTECTION INSERTED (g0154UU-Result) */
      if ErrorIn(g0154UU_try01553) {Result = g0154UU_try01553
      } else {
      g0154UU = ANY(g0154UU_try01553)
      Result = EID{Core.F_not_any(g0154UU).Id(),0}
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: self_eval @ And (throw: true) 
func E_self_eval_And (self EID) EID { 
    return /*(sm for self_eval @ And= EID)*/ To_And(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: And 
func EVAL_And (x *ClaireAny) EID { 
     return To_And(x).SelfEval()} 
  
// or expression
//
/* {1} OPT.The go function for: self_print(self:Or) [] */
func (self *Or ) SelfPrint () EID { 
    var Result EID 
    PRINC("(")
    Result = F_Language_printbox_list3(self.Args,MakeString(" | "))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(")")
    Result = EVOID
    }
    return Result} 
  
// The EID go function for: self_print @ Or (throw: true) 
func E_self_print_Or_Language (self EID) EID { 
    return /*(sm for self_print @ Or= EID)*/ To_Or(OBJ(self)).SelfPrint( )} 
  
/* {1} OPT.The go function for: self_eval(self:Or) [] */
func (self *Or ) SelfEval () EID { 
    var Result EID 
    var g0162I *ClaireBoolean  
    var g0162I_try01632 EID 
    /* For:2 */{ 
      var x *ClaireAny  
      _ = x
      g0162I_try01632= EID{CFALSE.Id(),0}
      var x_support *ClaireList  
      x_support = self.Args
      x_len := x_support.Length()
      for i_it := 0; i_it < x_len; i_it++ { 
        x = x_support.At(i_it)
        var void_try4 EID 
        _ = void_try4
        var g0164I *ClaireBoolean  
        var g0164I_try01654 EID 
        /* Let:4 */{ 
          var g0166UU *ClaireAny  
          /* noccur = 1 */
          var g0166UU_try01675 EID 
          g0166UU_try01675 = EVAL(x)
          /* ERROR PROTECTION INSERTED (g0166UU-g0164I_try01654) */
          if ErrorIn(g0166UU_try01675) {g0164I_try01654 = g0166UU_try01675
          } else {
          g0166UU = ANY(g0166UU_try01675)
          g0164I_try01654 = EID{F_boolean_I_any(g0166UU).Id(),0}
          }
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (g0164I-void_try4) */
        if ErrorIn(g0164I_try01654) {void_try4 = g0164I_try01654
        } else {
        g0164I = ToBoolean(OBJ(g0164I_try01654))
        if (g0164I == CTRUE) /* If:4 */{ 
           /*v = g0162I_try01632, s =EID*/
g0162I_try01632 = EID{CTRUE.Id(),0}
          break
          } else {
          void_try4 = EID{CFALSE.Id(),0}
          /* If-4 */} 
        }
        /* ERROR PROTECTION INSERTED (void_try4-g0162I_try01632) */
        if ErrorIn(void_try4) {g0162I_try01632 = void_try4
        g0162I_try01632 = void_try4
        break
        } else {
        }
        /* loop-3 */} 
      /* For-2 */} 
    /* ERROR PROTECTION INSERTED (g0162I-Result) */
    if ErrorIn(g0162I_try01632) {Result = g0162I_try01632
    } else {
    g0162I = ToBoolean(OBJ(g0162I_try01632))
    if (g0162I == CTRUE) /* If:2 */{ 
      Result = EID{CTRUE.Id(),0}
      } else {
      Result = EID{CFALSE.Id(),0}
      /* If-2 */} 
    }
    return Result} 
  
// The EID go function for: self_eval @ Or (throw: true) 
func E_self_eval_Or (self EID) EID { 
    return /*(sm for self_eval @ Or= EID)*/ To_Or(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Or 
func EVAL_Or (x *ClaireAny) EID { 
     return To_Or(x).SelfEval()} 
  
// ----------------- an anti-evaluator ---------------------------------
//
/* {1} OPT.The go function for: self_print(self:Quote) [] */
func (self *Quote ) SelfPrint () EID { 
    var Result EID 
    PRINC("quote(")
    Result = Core.F_CALL(C_print,ARGS(self.Arg.ToEID()))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(")")
    Result = EVOID
    }
    return Result} 
  
// The EID go function for: self_print @ Quote (throw: true) 
func E_self_print_Quote_Language (self EID) EID { 
    return /*(sm for self_print @ Quote= EID)*/ To_Quote(OBJ(self)).SelfPrint( )} 
  
/* {1} OPT.The go function for: self_eval(self:Quote) [] */
func (self *Quote ) SelfEval () EID { 
    var Result EID 
    Result = self.Arg.ToEID()
    return Result} 
  
// The EID go function for: self_eval @ Quote (throw: true) 
func E_self_eval_Quote (self EID) EID { 
    return /*(sm for self_eval @ Quote= EID)*/ To_Quote(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Quote 
func EVAL_Quote (x *ClaireAny) EID { 
     return To_Quote(x).SelfEval()} 
  
// *********************************************************************
// *      Part 3: Specialized structures                               *
// *********************************************************************
// optimized_instruction is the set of optimized messages.
// These are the forms produced by the optimizer. They correspond to basic
// kinds of evaluation.
//
// This is how a call to a compiled method can be compiled.
// We use the C external function
//
/* {1} OPT.The go function for: self_print(self:Call_method) [] */
func (self *CallMethod ) SelfPrint () EID { 
    var Result EID 
    Result = Core.F_print_any(self.Arg.Id())
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("(")
    Result = Core.F_princ_list(self.Args)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(")")
    Result = EVOID
    }}
    return Result} 
  
// The EID go function for: self_print @ Call_method (throw: true) 
func E_self_print_Call_method_Language (self EID) EID { 
    return /*(sm for self_print @ Call_method= EID)*/ To_CallMethod(OBJ(self)).SelfPrint( )} 
  
/* {1} OPT.The go function for: self_eval(self:Call_method) [] */
func (self *CallMethod ) SelfEval () EID { 
    var Result EID 
    /* Let:2 */{ 
      var start int  = ClEnv.Index
      /* noccur = 1 */
      /* Let:3 */{ 
        var Cprop *ClaireMethod   = self.Arg
        /* noccur = 1 */
        /* For:4 */{ 
          var x *ClaireAny  
          _ = x
          Result= EID{CFALSE.Id(),0}
          var x_support *ClaireList  
          x_support = self.Args
          x_len := x_support.Length()
          for i_it := 0; i_it < x_len; i_it++ { 
            x = x_support.At(i_it)
            var void_try6 EID 
            _ = void_try6
            /* LetEID:6 */{ 
              var g0168UU EID 
              g0168UU = EVAL(x)
              /* ERROR PROTECTION INSERTED (g0168UU-void_try6) */
              if ErrorIn(g0168UU) {void_try6 = g0168UU
              } else {
              void_try6 = ClEnv.Push(g0168UU)}
              /* LetEID-6 */} 
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
        Result = Core.F_execute_method(Cprop,start,CTRUE)
        }
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: self_eval @ Call_method (throw: true) 
func E_self_eval_Call_method (self EID) EID { 
    return /*(sm for self_eval @ Call_method= EID)*/ To_CallMethod(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Call_method 
func EVAL_Call_method (x *ClaireAny) EID { 
     return To_CallMethod(x).SelfEval()} 
  
// same thing with one only argument: we do not use the stack
/* {1} OPT.The go function for: self_eval(self:Call_method1) [] */
func (self *CallMethod1 ) SelfEval () EID { 
    var Result EID 
    /* Let:2 */{ 
      var f *ClaireMethod   = self.Arg
      /* noccur = 1 */
      /* Let:3 */{ 
        var l *ClaireList   = self.Args
        /* noccur = 1 */
        /* LetEID:4 */{ 
          var g0169UU EID 
          g0169UU = EVAL(l.At(1-1))
          /* ERROR PROTECTION INSERTED (g0169UU-Result) */
          if ErrorIn(g0169UU) {Result = g0169UU
          } else {
          Result = FASTCALL1(f,g0169UU)}
          /* LetEID-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: self_eval @ Call_method1 (throw: true) 
func E_self_eval_Call_method1 (self EID) EID { 
    return /*(sm for self_eval @ Call_method1= EID)*/ To_CallMethod1(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Call_method1 
func EVAL_Call_method1 (x *ClaireAny) EID { 
     return To_CallMethod1(x).SelfEval()} 
  
// same thing with two arguments
/* {1} OPT.The go function for: self_eval(self:Call_method2) [] */
func (self *CallMethod2 ) SelfEval () EID { 
    var Result EID 
    /* Let:2 */{ 
      var f *ClaireMethod   = self.Arg
      /* noccur = 1 */
      /* Let:3 */{ 
        var l *ClaireList   = self.Args
        /* noccur = 2 */
        /* LetEID:4 */{ 
          var g0170UU EID 
          g0170UU = EVAL(l.At(1-1))
          /* ERROR PROTECTION INSERTED (g0170UU-Result) */
          if ErrorIn(g0170UU) {Result = g0170UU
          } else {
          var g0171UU EID 
          g0171UU = EVAL(l.At(2-1))
          /* ERROR PROTECTION INSERTED (g0171UU-Result) */
          if ErrorIn(g0171UU) {Result = g0171UU
          } else {
          Result = FASTCALL2(f,g0170UU,g0171UU)}}
          /* LetEID-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: self_eval @ Call_method2 (throw: true) 
func E_self_eval_Call_method2 (self EID) EID { 
    return /*(sm for self_eval @ Call_method2= EID)*/ To_CallMethod2(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Call_method2 
func EVAL_Call_method2 (x *ClaireAny) EID { 
     return To_CallMethod2(x).SelfEval()} 
  
// same thing with two arguments
/* {1} OPT.The go function for: self_eval(self:Call_method3) [] */
func (self *LanguageCallMethod3 ) SelfEval () EID { 
    var Result EID 
    /* Let:2 */{ 
      var f *ClaireMethod   = self.Arg
      /* noccur = 1 */
      /* Let:3 */{ 
        var l *ClaireList   = self.Args
        /* noccur = 3 */
        /* LetEID:4 */{ 
          var g0172UU EID 
          g0172UU = EVAL(l.At(1-1))
          /* ERROR PROTECTION INSERTED (g0172UU-Result) */
          if ErrorIn(g0172UU) {Result = g0172UU
          } else {
          var g0173UU EID 
          g0173UU = EVAL(l.At(2-1))
          /* ERROR PROTECTION INSERTED (g0173UU-Result) */
          if ErrorIn(g0173UU) {Result = g0173UU
          } else {
          var g0174UU EID 
          g0174UU = EVAL(l.At(3-1))
          /* ERROR PROTECTION INSERTED (g0174UU-Result) */
          if ErrorIn(g0174UU) {Result = g0174UU
          } else {
          Result = FASTCALL3(f,g0172UU,g0173UU,g0174UU)}}}
          /* LetEID-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: self_eval @ Call_method3 (throw: true) 
func E_self_eval_Call_method3 (self EID) EID { 
    return /*(sm for self_eval @ Call_method3= EID)*/ To_LanguageCallMethod3(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Call_method3 
func EVAL_Language_Call_method3 (x *ClaireAny) EID { 
     return To_LanguageCallMethod3(x).SelfEval()} 
  
// an instruction to read a slot
//
/* {1} OPT.The go function for: self_print(self:Call_slot) [] */
func (self *CallSlot ) SelfPrint () EID { 
    var Result EID 
    Result = Core.F_print_any(self.Selector.Id())
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("(")
    Result = Core.F_CALL(C_print,ARGS(self.Arg.ToEID()))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(")")
    Result = EVOID
    }}
    return Result} 
  
// The EID go function for: self_print @ Call_slot (throw: true) 
func E_self_print_Call_slot_Language (self EID) EID { 
    return /*(sm for self_print @ Call_slot= EID)*/ To_CallSlot(OBJ(self)).SelfPrint( )} 
  
/* {1} OPT.The go function for: self_eval(self:Call_slot) [] */
func (self *CallSlot ) SelfEval () EID { 
    var Result EID 
    /* Let:2 */{ 
      var g0175UU *ClaireAny  
      /* noccur = 1 */
      var g0175UU_try01763 EID 
      g0175UU_try01763 = EVAL(self.Arg)
      /* ERROR PROTECTION INSERTED (g0175UU-Result) */
      if ErrorIn(g0175UU_try01763) {Result = g0175UU_try01763
      } else {
      g0175UU = ANY(g0175UU_try01763)
      Result = Core.F_get_slot(self.Selector,ToObject(g0175UU)).ToEID()
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: self_eval @ Call_slot (throw: true) 
func E_self_eval_Call_slot (self EID) EID { 
    return /*(sm for self_eval @ Call_slot= EID)*/ To_CallSlot(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Call_slot 
func EVAL_Call_slot (x *ClaireAny) EID { 
     return To_CallSlot(x).SelfEval()} 
  
// an instruction to read an array
// selector is an exp with type array, arg is an exp with type integer, and test
// contains the inferred member_type of the array
//
/* {1} OPT.The go function for: self_print(self:Call_array) [] */
func (self *CallArray ) SelfPrint () EID { 
    var Result EID 
    Result = Core.F_CALL(C_print,ARGS(self.Selector.ToEID()))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("[")
    Result = Core.F_CALL(C_print,ARGS(self.Arg.ToEID()))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("]")
    Result = EVOID
    }}
    return Result} 
  
// The EID go function for: self_print @ Call_array (throw: true) 
func E_self_print_Call_array_Language (self EID) EID { 
    return /*(sm for self_print @ Call_array= EID)*/ To_CallArray(OBJ(self)).SelfPrint( )} 
  
/* {1} OPT.The go function for: self_eval(self:Call_array) [] */
func (self *CallArray ) SelfEval () EID { 
    var Result EID 
    /* Let:2 */{ 
      var g0177UU *ClaireAny  
      /* noccur = 1 */
      var g0177UU_try01793 EID 
      g0177UU_try01793 = EVAL(self.Selector)
      /* ERROR PROTECTION INSERTED (g0177UU-Result) */
      if ErrorIn(g0177UU_try01793) {Result = g0177UU_try01793
      } else {
      g0177UU = ANY(g0177UU_try01793)
      /* Let:3 */{ 
        var g0178UU *ClaireAny  
        /* noccur = 1 */
        var g0178UU_try01804 EID 
        g0178UU_try01804 = EVAL(self.Arg)
        /* ERROR PROTECTION INSERTED (g0178UU-Result) */
        if ErrorIn(g0178UU_try01804) {Result = g0178UU_try01804
        } else {
        g0178UU = ANY(g0178UU_try01804)
        Result = Core.F_nth_array(ToArray(g0177UU),ToInteger(g0178UU).Value)
        }
        /* Let-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: self_eval @ Call_array (throw: true) 
func E_self_eval_Call_array (self EID) EID { 
    return /*(sm for self_eval @ Call_array= EID)*/ To_CallArray(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Call_array 
func EVAL_Call_array (x *ClaireAny) EID { 
     return To_CallArray(x).SelfEval()} 
  
// an instruction to read a table
//
/* {1} OPT.The go function for: self_print(self:Call_table) [] */
func (self *CallTable ) SelfPrint () EID { 
    var Result EID 
    Result = Core.F_print_any(self.Selector.Id())
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("[")
    Result = Core.F_CALL(C_print,ARGS(self.Arg.ToEID()))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("]")
    Result = EVOID
    }}
    return Result} 
  
// The EID go function for: self_print @ Call_table (throw: true) 
func E_self_print_Call_table_Language (self EID) EID { 
    return /*(sm for self_print @ Call_table= EID)*/ To_CallTable(OBJ(self)).SelfPrint( )} 
  
/* {1} OPT.The go function for: self_eval(self:Call_table) [] */
func (self *CallTable ) SelfEval () EID { 
    var Result EID 
    if (self.Test == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0181UU *ClaireAny  
        /* noccur = 1 */
        var g0181UU_try01824 EID 
        g0181UU_try01824 = EVAL(self.Arg)
        /* ERROR PROTECTION INSERTED (g0181UU-Result) */
        if ErrorIn(g0181UU_try01824) {Result = g0181UU_try01824
        } else {
        g0181UU = ANY(g0181UU_try01824)
        Result = Core.F_nth_table1(self.Selector,g0181UU)
        }
        /* Let-3 */} 
      } else {
      /* Let:3 */{ 
        var g0183UU *ClaireAny  
        /* noccur = 1 */
        var g0183UU_try01844 EID 
        g0183UU_try01844 = EVAL(self.Arg)
        /* ERROR PROTECTION INSERTED (g0183UU-Result) */
        if ErrorIn(g0183UU_try01844) {Result = g0183UU_try01844
        } else {
        g0183UU = ANY(g0183UU_try01844)
        Result = Core.F_get_table(self.Selector,g0183UU).ToEID()
        }
        /* Let-3 */} 
      /* If-2 */} 
    return Result} 
  
// The EID go function for: self_eval @ Call_table (throw: true) 
func E_self_eval_Call_table (self EID) EID { 
    return /*(sm for self_eval @ Call_table= EID)*/ To_CallTable(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Call_table 
func EVAL_Call_table (x *ClaireAny) EID { 
     return To_CallTable(x).SelfEval()} 
  
// an instruction to write a slot
// the structure is complex: see ocall.cl
//
/* {1} OPT.The go function for: self_print(self:Update) [] */
func (self *Update ) SelfPrint () EID { 
    var Result EID 
    Result = Core.F_CALL(C_print,ARGS(self.Selector.ToEID()))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("(")
    Result = Core.F_CALL(C_print,ARGS(Core.F_CALL(C_arg,ARGS(self.ClaireVar.ToEID()))))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(") := ")
    Result = Core.F_CALL(C_print,ARGS(self.Value.ToEID()))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("")
    Result = EVOID
    }}}
    return Result} 
  
// The EID go function for: self_print @ Update (throw: true) 
func E_self_print_Update_Language (self EID) EID { 
    return /*(sm for self_print @ Update= EID)*/ To_Update(OBJ(self)).SelfPrint( )} 
  
/* {1} OPT.The go function for: self_eval(self:Update) [] */
func (self *Update ) SelfEval () EID { 
    var Result EID 
    /* Let:2 */{ 
      var s *ClaireAny   = self.Selector
      /* noccur = 4 */
      if (s.Isa.IsIn(C_property) == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var g0185 *ClaireProperty   = ToProperty(s)
          /* noccur = 1 */
          /* Let:5 */{ 
            var g0187UU *ClaireAny  
            /* noccur = 1 */
            var g0187UU_try01896 EID 
            g0187UU_try01896 = EVAL(ANY(Core.F_CALL(C_arg,ARGS(self.ClaireVar.ToEID()))))
            /* ERROR PROTECTION INSERTED (g0187UU-Result) */
            if ErrorIn(g0187UU_try01896) {Result = g0187UU_try01896
            } else {
            g0187UU = ANY(g0187UU_try01896)
            /* Let:6 */{ 
              var g0188UU *ClaireAny  
              /* noccur = 1 */
              var g0188UU_try01907 EID 
              g0188UU_try01907 = EVAL(self.Value)
              /* ERROR PROTECTION INSERTED (g0188UU-Result) */
              if ErrorIn(g0188UU_try01907) {Result = g0188UU_try01907
              } else {
              g0188UU = ANY(g0188UU_try01907)
              Result = Core.F_put_property2(g0185,ToObject(g0187UU),g0188UU)
              }
              /* Let-6 */} 
            }
            /* Let-5 */} 
          /* Let-4 */} 
        /* If!3 */}  else if (C_table.Id() == s.Isa.Id()) /* If:3 */{ 
        /* Let:4 */{ 
          var g0186 *ClaireTable   = ToTable(s)
          /* noccur = 1 */
          /* Let:5 */{ 
            var g0191UU *ClaireAny  
            /* noccur = 1 */
            var g0191UU_try01936 EID 
            g0191UU_try01936 = EVAL(ANY(Core.F_CALL(C_arg,ARGS(self.ClaireVar.ToEID()))))
            /* ERROR PROTECTION INSERTED (g0191UU-Result) */
            if ErrorIn(g0191UU_try01936) {Result = g0191UU_try01936
            } else {
            g0191UU = ANY(g0191UU_try01936)
            /* Let:6 */{ 
              var g0192UU *ClaireAny  
              /* noccur = 1 */
              var g0192UU_try01947 EID 
              g0192UU_try01947 = EVAL(self.Value)
              /* ERROR PROTECTION INSERTED (g0192UU-Result) */
              if ErrorIn(g0192UU_try01947) {Result = g0192UU_try01947
              } else {
              g0192UU = ANY(g0192UU_try01947)
              Result = Core.F_nth_equal_table1(g0186,g0191UU,g0192UU)
              }
              /* Let-6 */} 
            }
            /* Let-5 */} 
          /* Let-4 */} 
        } else {
        Result = EID{CFALSE.Id(),0}
        /* If-3 */} 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = EID{CNULL,0}
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: self_eval @ Update (throw: true) 
func E_self_eval_Update (self EID) EID { 
    return /*(sm for self_eval @ Update= EID)*/ To_Update(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Update 
func EVAL_Update (x *ClaireAny) EID { 
     return To_Update(x).SelfEval()} 
  
// ------------------ SUPER: a jump in the set lattice ---------------
// A "super" allows one to execute a message as if the type of the receiver
// was a given abstract_class.
// However we require that the receiver be in the specified abstract_class.
// The form of the super is: SELECTOR@ABSTRACT_CLASS(RECEIVER , ...)
//
/* {1} OPT.The go function for: self_print(self:Super) [] */
func (self *Super ) SelfPrint () EID { 
    var Result EID 
    /* Let:2 */{ 
      var _Zl int  = Core.C_pretty.Index
      /* noccur = 1 */
      /* Let:3 */{ 
        var _Zs *ClaireProperty   = self.Selector
        /* noccur = 1 */
        /* Let:4 */{ 
          var _Za *ClaireList   = self.Args
          /* noccur = 1 */
          Result = Core.F_print_any(_Zs.Id())
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC("@")
          Result = Core.F_print_any(self.CastTo.Id())
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC("(")
          F_set_level_void()
          Result = F_Language_printbox_list2(_Za)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(")")
          Result = EVOID
          }}}
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
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: self_print @ Super (throw: true) 
func E_self_print_Super_Language (self EID) EID { 
    return /*(sm for self_print @ Super= EID)*/ To_Super(OBJ(self)).SelfPrint( )} 
  
/* {1} OPT.The go function for: self_eval(self:Super) [] */
func (self *Super ) SelfEval () EID { 
    var Result EID 
    /* Let:2 */{ 
      var start int  = ClEnv.Index
      /* noccur = 2 */
      /* Let:3 */{ 
        var t *ClaireType   = self.CastTo
        /* noccur = 1 */
        /* Let:4 */{ 
          var c *ClaireClass   = t.Class_I()
          /* noccur = 1 */
          /* Let:5 */{ 
            var p *ClaireProperty   = self.Selector
            /* noccur = 2 */
            /* For:6 */{ 
              var x *ClaireAny  
              _ = x
              Result= EID{CFALSE.Id(),0}
              var x_support *ClaireList  
              x_support = self.Args
              x_len := x_support.Length()
              for i_it := 0; i_it < x_len; i_it++ { 
                x = x_support.At(i_it)
                var void_try8 EID 
                _ = void_try8
                /* LetEID:8 */{ 
                  var g0195UU EID 
                  g0195UU = EVAL(x)
                  /* ERROR PROTECTION INSERTED (g0195UU-void_try8) */
                  if ErrorIn(g0195UU) {void_try8 = g0195UU
                  } else {
                  void_try8 = ClEnv.Push(g0195UU)}
                  /* LetEID-8 */} 
                /* ERROR PROTECTION INSERTED (void_try8-Result) */
                if ErrorIn(void_try8) {Result = void_try8
                Result = void_try8
                break
                } else {
                }
                /* loop-7 */} 
              /* For-6 */} 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            Result = Core.F_eval_message_property(p,Core.F_find_which_class(c,p.Definition,start,ClEnv.Index),start,CTRUE)
            }
            /* Let-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: self_eval @ Super (throw: true) 
func E_self_eval_Super (self EID) EID { 
    return /*(sm for self_eval @ Super= EID)*/ To_Super(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Super 
func EVAL_Super (x *ClaireAny) EID { 
     return To_Super(x).SelfEval()} 
  
//--------------- comments ------------------------------------------
// the cast is the new form of simple super
//
/* {1} OPT.The go function for: self_print(x:Cast) [] */
func (x *Cast ) SelfPrint () EID { 
    var Result EID 
    Result = F_printexp_any(x.Arg,CFALSE)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(" as ")
    Result = F_printexp_any(x.SetArg.Id(),CFALSE)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("")
    Result = EVOID
    }}
    return Result} 
  
// The EID go function for: self_print @ Cast (throw: true) 
func E_self_print_Cast_Language (x EID) EID { 
    return /*(sm for self_print @ Cast= EID)*/ To_Cast(OBJ(x)).SelfPrint( )} 
  
/* {1} OPT.The go function for: self_eval(self:Cast) [] */
func (self *Cast ) SelfEval () EID { 
    var Result EID 
    /* Let:2 */{ 
      var x *ClaireAny  
      /* noccur = 3 */
      var x_try01973 EID 
      x_try01973 = EVAL(self.Arg)
      /* ERROR PROTECTION INSERTED (x-Result) */
      if ErrorIn(x_try01973) {Result = x_try01973
      } else {
      x = ANY(x_try01973)
      /* Let:3 */{ 
        var y *ClaireType   = self.SetArg
        /* noccur = 4 */
        var g0198I *ClaireBoolean  
        if (y.Isa.IsIn(C_Param) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var g0196 *ClaireParam   = To_Param(y.Id())
            /* noccur = 4 */
            g0198I = MakeBoolean(((g0196.Arg.Id() == C_list.Id()) || 
                (g0196.Arg.Id() == C_set.Id())) && (C_set.Id() == g0196.Args.At(1-1).Isa.Id()) && (Core.F__Z_any1(x,g0196.Arg) == CTRUE))
            /* Let-5 */} 
          } else {
          g0198I = CFALSE
          /* If-4 */} 
        if (g0198I == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var g0199UU *ClaireAny  
            /* noccur = 1 */
            var g0199UU_try02006 EID 
            g0199UU_try02006 = Core.F_the_type(ToType(To_Param(y.Id()).Args.At(1-1)))
            /* ERROR PROTECTION INSERTED (g0199UU-Result) */
            if ErrorIn(g0199UU_try02006) {Result = g0199UU_try02006
            } else {
            g0199UU = ANY(g0199UU_try02006)
            Result = Core.F_check_in_bag(ToBag(x),C_bag,ToType(g0199UU))
            }
            /* Let-5 */} 
          } else {
          Result = Core.F_check_in_any(x,y)
          /* If-4 */} 
        /* Let-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: self_eval @ Cast (throw: true) 
func E_self_eval_Cast (self EID) EID { 
    return /*(sm for self_eval @ Cast= EID)*/ To_Cast(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Cast 
func EVAL_Cast (x *ClaireAny) EID { 
     return To_Cast(x).SelfEval()} 
  
// v3.3.16 */
// ----------------- return from a loop --------------------------------
//
// return_error is an exception that is handled by the "for" family
// of structures
//
/* {1} OPT.The go function for: self_print(self:Return) [] */
func (self *Return ) SelfPrint () EID { 
    var Result EID 
    PRINC("break(")
    Core.C_pretty.Index = (Core.C_pretty.Index+2)
    Result = Core.F_CALL(C_print,ARGS(self.Arg.ToEID()))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    Core.C_pretty.Index = (Core.C_pretty.Index-2)
    PRINC(")")
    Result = EVOID
    }
    return Result} 
  
// The EID go function for: self_print @ Return (throw: true) 
func E_self_print_Return_Language (self EID) EID { 
    return /*(sm for self_print @ Return= EID)*/ To_Return(OBJ(self)).SelfPrint( )} 
  
/* {1} OPT.The go function for: self_eval(self:Return) [] */
func (self *Return ) SelfEval () EID { 
    var Result EID 
    /* Let:2 */{ 
      var g0203UU *ClaireObject  
      /* noccur = 1 */
      var g0203UU_try02043 EID 
      /* Let:3 */{ 
        var g0205UU *ClaireAny  
        /* noccur = 1 */
        var g0205UU_try02064 EID 
        g0205UU_try02064 = EVAL(self.Arg)
        /* ERROR PROTECTION INSERTED (g0205UU-g0203UU_try02043) */
        if ErrorIn(g0205UU_try02064) {g0203UU_try02043 = g0205UU_try02064
        } else {
        g0205UU = ANY(g0205UU_try02064)
        g0203UU_try02043 = Core.C_return_error.Make(g0205UU).ToEID()
        }
        /* Let-3 */} 
      /* ERROR PROTECTION INSERTED (g0203UU-Result) */
      if ErrorIn(g0203UU_try02043) {Result = g0203UU_try02043
      } else {
      g0203UU = ToObject(OBJ(g0203UU_try02043))
      Result = ToException(g0203UU.Id()).Close()
      }
      /* Let-2 */} 
    return RangeCheck(ToType(C_error.Id()),Result)} 
  
// The EID go function for: self_eval @ Return (throw: true) 
func E_self_eval_Return (self EID) EID { 
    return /*(sm for self_eval @ Return= EID)*/ To_Return(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Return 
func EVAL_Return (x *ClaireAny) EID { 
     return To_Return(x).SelfEval()} 
  
// ****************************************************************
// *       Part 4: Miscellaneous on instructions                  *
// ****************************************************************
// substitute any variable with same name as x with the value val
/* {1} OPT.The go function for: substitution(self:any,x:Variable,val:any) [] */
func F_substitution_any (self *ClaireAny ,x *ClaireVariable ,val *ClaireAny ) *ClaireAny  { 
    // procedure body with s =  
var Result *ClaireAny  
    if (self.Isa.IsIn(C_Variable) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0207 *ClaireVariable   = To_Variable(self)
        /* noccur = 2 */
        Result = IfThenElse((g0207.Pname.Id() == x.Pname.Id()),
          val,
          g0207.Id())
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(C_list) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0208 *ClaireList   = ToList(self)
        /* noccur = 7 */
        /* Let:4 */{ 
          var i int  = 1
          /* noccur = 8 */
          /* Let:5 */{ 
            var g0209 int  = g0208.Length()
            /* noccur = 1 */
            for (i <= g0209) /* while:6 */{ 
              if ((g0208.At(i-1).Isa.IsIn(C_Variable) == CTRUE) || 
                  (g0208.At(i-1).Isa.IsIn(C_unbound_symbol) == CTRUE)) /* If:7 */{ 
                ToArray(g0208.Id()).NthPut(i,F_substitution_any(g0208.At(i-1),x,val))
                } else {
                F_substitution_any(g0208.At(i-1),x,val)
                /* If-7 */} 
              i = (i+1)
              /* while-6 */} 
            /* Let-5 */} 
          /* Let-4 */} 
        Result = g0208.Id()
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(C_unbound_symbol) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0210 *ClaireUnboundSymbol   = ToUnboundSymbol(self)
        /* noccur = 2 */
        Result = IfThenElse((g0210.Name.Id() == x.Pname.Id()),
          val,
          g0210.Id())
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(C_Instruction) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0211 *ClaireInstruction   = To_Instruction(self)
        /* noccur = 4 */
        /* For:4 */{ 
          var s *ClaireAny  
          _ = s
          var s_support *ClaireList  
          s_support = g0211.Id().Isa.Slots
          for _,s = range(s_support.ValuesO())/* loop2:5 */{ 
            /* Let:6 */{ 
              var y *ClaireAny   = Core.F_get_slot(ToSlot(s),ToObject(g0211.Id()))
              /* noccur = 4 */
              if ((y.Isa.IsIn(C_Variable) == CTRUE) || 
                  (y.Isa.IsIn(C_unbound_symbol) == CTRUE)) /* If:7 */{ 
                Core.F_put_slot(ToSlot(s),ToObject(g0211.Id()),F_substitution_any(y,x,val))
                } else {
                F_substitution_any(y,x,val)
                /* If-7 */} 
              /* Let-6 */} 
            /* loop-5 */} 
          /* For-4 */} 
        Result = g0211.Id()
        /* Let-3 */} 
      } else {
      Result = self
      /* If-2 */} 
    return Result} 
  
// The EID go function for: substitution @ any (throw: false) 
func E_substitution_any (self EID,x EID,val EID) EID { 
    return /*(sm for substitution @ any= any)*/ F_substitution_any(ANY(self),To_Variable(OBJ(x)),ANY(val) ).ToEID()} 
  
// count the number of occurrences of x
/* {1} OPT.The go function for: occurrence(self:any,x:Variable) [] */
func F_occurrence_any (self *ClaireAny ,x *ClaireVariable ) int { 
    // procedure body with s =  
var Result int 
    if (self.Isa.IsIn(C_Variable) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0213 *ClaireVariable   = To_Variable(self)
        /* noccur = 1 */
        if (g0213.Pname.Id() == x.Pname.Id()) /* If:4 */{ 
          Result = 1
          } else {
          Result = 0
          /* If-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(C_list) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0214 *ClaireList   = ToList(self)
        /* noccur = 2 */
        /* Let:4 */{ 
          var n int  = 0
          /* noccur = 3 */
          /* Let:5 */{ 
            var i int  = 1
            /* noccur = 4 */
            /* Let:6 */{ 
              var g0215 int  = g0214.Length()
              /* noccur = 1 */
              for (i <= g0215) /* while:7 */{ 
                n = (n+F_occurrence_any(g0214.At(i-1),x))
                i = (i+1)
                /* while-7 */} 
              /* Let-6 */} 
            /* Let-5 */} 
          Result = n
          /* Let-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(C_unbound_symbol) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0216 *ClaireUnboundSymbol   = ToUnboundSymbol(self)
        /* noccur = 1 */
        if (g0216.Name.Id() == x.Pname.Id()) /* If:4 */{ 
          Result = 1
          } else {
          Result = 0
          /* If-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(C_Instruction) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0217 *ClaireInstruction   = To_Instruction(self)
        /* noccur = 2 */
        /* Let:4 */{ 
          var n int  = 0
          /* noccur = 3 */
          /* For:5 */{ 
            var s *ClaireAny  
            _ = s
            var s_support *ClaireList  
            s_support = g0217.Id().Isa.Slots
            for _,s = range(s_support.ValuesO())/* loop2:6 */{ 
              n = (n+F_occurrence_any(Core.F_get_slot(ToSlot(s),ToObject(g0217.Id())),x))
              /* loop-6 */} 
            /* For-5 */} 
          Result = n
          /* Let-4 */} 
        /* Let-3 */} 
      } else {
      Result = 0
      /* If-2 */} 
    return Result} 
  
// The EID go function for: occurrence @ any (throw: false) 
func E_occurrence_any (self EID,x EID) EID { 
    return EID{C__INT,IVAL(/*(sm for occurrence @ any= integer)*/ F_occurrence_any(ANY(self),To_Variable(OBJ(x)) ))}} 
  
// new version in CLAIRE4 : see if the variable is changed
/* {1} OPT.The go function for: occurchange(self:any,x:Variable) [] */
func F_occurchange_any (self *ClaireAny ,x *ClaireVariable ) *ClaireBoolean  { 
    // procedure body with s =  
var Result *ClaireBoolean  
    if (self.Isa.IsIn(C_Assign) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0219 *Assign   = To_Assign(self)
        /* noccur = 1 */
        Result = Equal(ANY(Core.F_CALL(C_mClaire_pname,ARGS(g0219.ClaireVar.ToEID()))),x.Pname.Id())
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(C_list) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0220 *ClaireList   = ToList(self)
        /* noccur = 1 */
        /* Let:4 */{ 
          var g0223UU *ClaireAny  
          /* noccur = 1 */
          /* For:5 */{ 
            var y *ClaireAny  
            _ = y
            g0223UU= CFALSE.Id()
            var y_support *ClaireList  
            y_support = g0220
            y_len := y_support.Length()
            for i_it := 0; i_it < y_len; i_it++ { 
              y = y_support.At(i_it)
              if (F_occurchange_any(y,x) == CTRUE) /* If:7 */{ 
                 /*v = g0223UU, s =any*/
g0223UU = CTRUE.Id()
                break
                /* If-7 */} 
              /* loop-6 */} 
            /* For-5 */} 
          Result = F_boolean_I_any(g0223UU)
          /* Let-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(C_Instruction) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0221 *ClaireInstruction   = To_Instruction(self)
        /* noccur = 2 */
        /* Let:4 */{ 
          var g0224UU *ClaireAny  
          /* noccur = 1 */
          /* For:5 */{ 
            var s *ClaireAny  
            _ = s
            g0224UU= CFALSE.Id()
            var s_support *ClaireList  
            s_support = g0221.Id().Isa.Slots
            for _,s = range(s_support.ValuesO())/* loop2:6 */{ 
              if (F_occurchange_any(Core.F_get_slot(ToSlot(s),ToObject(g0221.Id())),x) == CTRUE) /* If:7 */{ 
                 /*v = g0224UU, s =any*/
g0224UU = CTRUE.Id()
                break
                /* If-7 */} 
              /* loop-6 */} 
            /* For-5 */} 
          Result = F_boolean_I_any(g0224UU)
          /* Let-4 */} 
        /* Let-3 */} 
      } else {
      Result = CFALSE
      /* If-2 */} 
    return Result} 
  
// The EID go function for: occurchange @ any (throw: false) 
func E_occurchange_any (self EID,x EID) EID { 
    return EID{/*(sm for occurchange @ any= boolean)*/ F_occurchange_any(ANY(self),To_Variable(OBJ(x)) ).Id(),0}} 
  
// a variant in CLAIRE 4 that assumes that variable have reveived their lexical bind (index)
/* {1} OPT.The go function for: occurexact(self:any,x:Variable) [] */
func F_Language_occurexact_any (self *ClaireAny ,x *ClaireVariable ) int { 
    // procedure body with s =  
var Result int 
    if (self.Isa.IsIn(C_Variable) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0225 *ClaireVariable   = To_Variable(self)
        /* noccur = 2 */
        if ((g0225.Pname.Id() == x.Pname.Id()) && 
            (g0225.Index == x.Index)) /* If:4 */{ 
          Result = 1
          } else {
          Result = 0
          /* If-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(C_list) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0226 *ClaireList   = ToList(self)
        /* noccur = 2 */
        /* Let:4 */{ 
          var n int  = 0
          /* noccur = 3 */
          /* Let:5 */{ 
            var i int  = 1
            /* noccur = 4 */
            /* Let:6 */{ 
              var g0227 int  = g0226.Length()
              /* noccur = 1 */
              for (i <= g0227) /* while:7 */{ 
                n = (n+F_Language_occurexact_any(g0226.At(i-1),x))
                i = (i+1)
                /* while-7 */} 
              /* Let-6 */} 
            /* Let-5 */} 
          Result = n
          /* Let-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(C_unbound_symbol) == CTRUE) /* If:2 */{ 
      Result = 0
      /* If!2 */}  else if (self.Isa.IsIn(C_Instruction) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0229 *ClaireInstruction   = To_Instruction(self)
        /* noccur = 2 */
        /* Let:4 */{ 
          var n int  = 0
          /* noccur = 3 */
          /* For:5 */{ 
            var s *ClaireAny  
            _ = s
            var s_support *ClaireList  
            s_support = g0229.Id().Isa.Slots
            for _,s = range(s_support.ValuesO())/* loop2:6 */{ 
              n = (n+F_Language_occurexact_any(Core.F_get_slot(ToSlot(s),ToObject(g0229.Id())),x))
              /* loop-6 */} 
            /* For-5 */} 
          Result = n
          /* Let-4 */} 
        /* Let-3 */} 
      } else {
      Result = 0
      /* If-2 */} 
    return Result} 
  
// The EID go function for: occurexact @ any (throw: false) 
func E_Language_occurexact_any (self EID,x EID) EID { 
    return EID{C__INT,IVAL(/*(sm for occurexact @ any= integer)*/ F_Language_occurexact_any(ANY(self),To_Variable(OBJ(x)) ))}} 
  
// makes a (deep) copy of the instruction self
//
/* {1} OPT.The go function for: instruction_copy(self:any) [] */
func F_instruction_copy_any (self *ClaireAny ) *ClaireAny  { 
    // procedure body with s =  
var Result *ClaireAny  
    if (self.Isa.IsIn(C_list) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0231 *ClaireList   = ToList(self)
        /* noccur = 3 */
        /* Let:4 */{ 
          var l *ClaireList   = g0231.Copy()
          /* noccur = 2 */
          /* Let:5 */{ 
            var i int  = 1
            /* noccur = 5 */
            /* Let:6 */{ 
              var g0232 int  = g0231.Length()
              /* noccur = 1 */
              for (i <= g0232) /* while:7 */{ 
                ToArray(l.Id()).NthPut(i,F_instruction_copy_any(g0231.At(i-1)))
                i = (i+1)
                /* while-7 */} 
              /* Let-6 */} 
            /* Let-5 */} 
          Result = l.Id()
          /* Let-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(C_Variable) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0233 *ClaireVariable   = To_Variable(self)
        /* noccur = 1 */
        Result = g0233.Id()
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(C_Instruction) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0234 *ClaireInstruction   = To_Instruction(self)
        /* noccur = 3 */
        /* Let:4 */{ 
          var o *ClaireInstruction   = To_Instruction(g0234.Copy().Id())
          /* noccur = 2 */
          /* For:5 */{ 
            var s *ClaireAny  
            _ = s
            var s_support *ClaireList  
            s_support = g0234.Id().Isa.Slots
            for _,s = range(s_support.ValuesO())/* loop2:6 */{ 
              Core.F_put_slot(ToSlot(s),ToObject(o.Id()),F_instruction_copy_any(Core.F_get_slot(ToSlot(s),ToObject(g0234.Id()))))
              /* loop-6 */} 
            /* For-5 */} 
          Result = o.Id()
          /* Let-4 */} 
        /* Let-3 */} 
      } else {
      Result = self
      /* If-2 */} 
    return Result} 
  
// The EID go function for: instruction_copy @ any (throw: false) 
func E_instruction_copy_any (self EID) EID { 
    return /*(sm for instruction_copy @ any= any)*/ F_instruction_copy_any(ANY(self) ).ToEID()} 
  