/***** CLAIRE Compilation of file /Users/ycaseau/Dropbox/src/clairev4.0/src/meta/call.cl 
         [version 4.0.02 / safety 5] Monday 12-13-2021 *****/

package Language
import (_ "fmt"
	. "Kernel"
	"Core"
)

//-------- dumb function to prevent import errors --------
func import_g0062() { 
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
/* {1} The go function for: self_print(self:Call) [status=1] */
func (self *Call ) SelfPrint () EID { 
    var Result EID 
    { var _Zl int  = Core.C_pretty.Index
      _ = _Zl
      { var _Zs *ClaireProperty   = self.Selector
        { var _Za *ClaireList   = self.Args
          /*g_try(v2:"Result",loop:true) */
          if ((_Zs.Isa.IsIn(C_operation) == CTRUE) && 
              (_Za.Length() == 2)) { 
            Core.C_pretty.Index = (Core.C_pretty.Index+2)
            /*integer->integer*//*g_try(v2:"Result",loop:true) */
            Result = F_printe_any(_Za.At(1-1),_Zs)
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC(" ")
            /*g_try(v2:"Result",loop:true) */
            Result = Core.F_print_any(_Zs.Id())
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC(" ")
            /*g_try(v2:"Result",loop:true) */
            Result = F_lbreak_void()
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            /*g_try(v2:"Result",loop:true) */
            Result = F_printe_any(_Za.At(2-1),_Zs)
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC("")
            Result = EVOID
            }}}}
            }  else if (_Zs.Id() == C_nth.Id()) { 
            if (_Za.Length() == 3) { 
              /*g_try(v2:"Result",loop:true) */
              Result = F_printexp_any(_Za.At(1-1),CFALSE)
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              PRINC("[")
              /*g_try(v2:"Result",loop:true) */
              Result = Core.F_CALL(C_print,ARGS(_Za.At(2-1).ToEID()))
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              PRINC(",")
              /*g_try(v2:"Result",loop:true) */
              Result = Core.F_CALL(C_print,ARGS(_Za.At(3-1).ToEID()))
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              PRINC("]")
              Result = EVOID
              }}}
              }  else if (_Za.Length() == 1) { 
              /*g_try(v2:"Result",loop:true) */
              Result = F_printexp_any(_Za.At(1-1),CFALSE)
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              PRINC("[]")
              Result = EVOID
              }
              } else {
              /*g_try(v2:"Result",loop:true) */
              Result = F_printexp_any(_Za.At(1-1),CFALSE)
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              PRINC("[")
              /*g_try(v2:"Result",loop:true) */
              if (_Za.Length() == 2) { 
                Result = Core.F_CALL(C_print,ARGS(_Za.At(2-1).ToEID()))
                } else {
                Result = EID{CFALSE.Id(),0}
                } 
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              PRINC("]")
              Result = EVOID
              }}
              } 
            }  else if ((_Zs.Id() == C_nth_equal.Id()) && 
              (_Za.Length() >= 3)) { 
            { var a *ClaireAny   = _Za.At(3-1)
              { var o *ClaireAny  
                if (a.Isa.IsIn(C_Call) == CTRUE) { 
                  { var g0064 *Call   = To_Call(a)
                    _ = g0064
                    o = g0064.Selector.Id()
                    } 
                  } else {
                  o = CFALSE.Id()
                  } 
                if (_Za.Length() == 4) { 
                  /*g_try(v2:"Result",loop:true) */
                  Result = F_printexp_any(_Za.At(1-1),CFALSE)
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  PRINC("[")
                  /*g_try(v2:"Result",loop:true) */
                  Result = Core.F_CALL(C_print,ARGS(_Za.At(2-1).ToEID()))
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  PRINC(",")
                  /*g_try(v2:"Result",loop:true) */
                  Result = Core.F_CALL(C_print,ARGS(a.ToEID()))
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  PRINC("] := ")
                  /*g_try(v2:"Result",loop:true) */
                  Result = F_lbreak_integer(2)
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  /*g_try(v2:"Result",loop:true) */
                  Result = Core.F_CALL(C_print,ARGS(_Za.At(4-1).ToEID()))
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  PRINC("")
                  Result = EVOID
                  }}}}}
                  } else {
                  var g0068I *ClaireBoolean  
                  { var arg_1 *ClaireAny  
                    _ = arg_1
                    if (a.Isa.IsIn(C_Call) == CTRUE) { 
                      { var g0065 *Call   = To_Call(a)
                        _ = g0065
                        arg_1 = g0065.Args.At(1-1)
                        } 
                      } else {
                      arg_1 = CFALSE.Id()
                      } 
                    g0068I = F_sugar_ask_any(_Za.At(1-1),_Za.At(2-1),o,arg_1)
                    } 
                  if (g0068I == CTRUE) { 
                    /*g_try(v2:"Result",loop:true) */
                    Result = Core.F_CALL(C_print,ARGS(_Za.At(1-1).ToEID()))
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    PRINC("[")
                    /*g_try(v2:"Result",loop:true) */
                    Result = Core.F_CALL(C_print,ARGS(_Za.At(2-1).ToEID()))
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    PRINC("] :")
                    /*g_try(v2:"Result",loop:true) */
                    Result = Core.F_CALL(C_print,ARGS(o.ToEID()))
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    PRINC(" ")
                    /*g_try(v2:"Result",loop:true) */
                    Result = F_lbreak_integer(2)
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    /*g_try(v2:"Result",loop:true) */
                    Result = Core.F_CALL(C_print,ARGS(ToList(OBJ(Core.F_CALL(C_args,ARGS(a.ToEID())))).At(2-1).ToEID()))
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    PRINC("")
                    Result = EVOID
                    }}}}}
                    } else {
                    /*g_try(v2:"Result",loop:true) */
                    Result = Core.F_CALL(C_print,ARGS(_Za.At(1-1).ToEID()))
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    PRINC("[")
                    /*g_try(v2:"Result",loop:true) */
                    Result = Core.F_CALL(C_print,ARGS(_Za.At(2-1).ToEID()))
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    PRINC("] := ")
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
                    }}}}
                    } 
                  } 
                } 
              } 
            }  else if ((_Zs.Id() == C_Language_assign.Id()) && 
              (_Za.At(1-1).Isa.IsIn(C_property) == CTRUE)) { 
            { var a *ClaireAny   = _Za.At(3-1)
              { var o *ClaireAny  
                if (a.Isa.IsIn(C_Call) == CTRUE) { 
                  { var g0066 *Call   = To_Call(a)
                    _ = g0066
                    o = g0066.Selector.Id()
                    } 
                  } else {
                  o = CFALSE.Id()
                  } 
                var g0069I *ClaireBoolean  
                { var arg_2 *ClaireAny  
                  _ = arg_2
                  if (a.Isa.IsIn(C_Call) == CTRUE) { 
                    { var g0067 *Call   = To_Call(a)
                      _ = g0067
                      arg_2 = g0067.Args.At(1-1)
                      } 
                    } else {
                    arg_2 = CFALSE.Id()
                    } 
                  g0069I = F_sugar_ask_any(_Za.At(1-1),_Za.At(2-1),o,arg_2)
                  } 
                if (g0069I == CTRUE) { 
                  /*g_try(v2:"Result",loop:true) */
                  Result = Core.F_CALL(C_print,ARGS(_Za.At(1-1).ToEID()))
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  PRINC("(")
                  /*g_try(v2:"Result",loop:true) */
                  Result = Core.F_CALL(C_print,ARGS(_Za.At(2-1).ToEID()))
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  PRINC(") :")
                  /*g_try(v2:"Result",loop:true) */
                  Result = Core.F_CALL(C_print,ARGS(o.ToEID()))
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  PRINC(" ")
                  /*g_try(v2:"Result",loop:true) */
                  Result = F_lbreak_integer(2)
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  /*g_try(v2:"Result",loop:true) */
                  Result = Core.F_CALL(C_print,ARGS(ToList(OBJ(Core.F_CALL(C_args,ARGS(a.ToEID())))).At(2-1).ToEID()))
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  PRINC("")
                  Result = EVOID
                  }}}}}
                  } else {
                  /*g_try(v2:"Result",loop:true) */
                  Result = Core.F_CALL(C_print,ARGS(_Za.At(1-1).ToEID()))
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  PRINC("(")
                  /*g_try(v2:"Result",loop:true) */
                  Result = Core.F_CALL(C_print,ARGS(_Za.At(2-1).ToEID()))
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  PRINC(") := ")
                  /*g_try(v2:"Result",loop:true) */
                  Result = F_lbreak_integer(2)
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  /*g_try(v2:"Result",loop:true) */
                  Result = Core.F_CALL(C_print,ARGS(_Za.At(3-1).ToEID()))
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  PRINC("")
                  Result = EVOID
                  }}}}
                  } 
                } 
              } 
            }  else if ((_Zs.Id() == C_add.Id()) && 
              (_Za.At(1-1).Isa.IsIn(C_property) == CTRUE)) { 
            /*g_try(v2:"Result",loop:true) */
            Result = Core.F_CALL(C_print,ARGS(_Za.At(1-1).ToEID()))
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC("(")
            /*g_try(v2:"Result",loop:true) */
            Result = Core.F_CALL(C_print,ARGS(_Za.At(2-1).ToEID()))
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC(") :add ")
            /*g_try(v2:"Result",loop:true) */
            Result = F_lbreak_integer(2)
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            /*g_try(v2:"Result",loop:true) */
            Result = Core.F_CALL(C_print,ARGS(_Za.At(3-1).ToEID()))
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC("")
            Result = EVOID
            }}}}
            }  else if ((_Zs.Id() == C_delete.Id()) && 
              (_Za.At(1-1).Isa.IsIn(C_property) == CTRUE)) { 
            /*g_try(v2:"Result",loop:true) */
            Result = Core.F_CALL(C_print,ARGS(_Za.At(1-1).ToEID()))
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC("(")
            /*g_try(v2:"Result",loop:true) */
            Result = Core.F_CALL(C_print,ARGS(_Za.At(2-1).ToEID()))
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC(") :delete ")
            /*g_try(v2:"Result",loop:true) */
            Result = F_lbreak_integer(2)
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            /*g_try(v2:"Result",loop:true) */
            Result = Core.F_CALL(C_print,ARGS(_Za.At(3-1).ToEID()))
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC("")
            Result = EVOID
            }}}}
            }  else if ((_Za.At(1-1) == ClEnv.Id()) && 
              (_Za.Length() == 1)) { 
            /*g_try(v2:"Result",loop:true) */
            Result = Core.F_print_any(_Zs.Id())
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC("()")
            Result = EVOID
            }
            } else {
            /*g_try(v2:"Result",loop:true) */
            Result = Core.F_print_any(_Zs.Id())
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC("(")
            F_set_level_void()
            /*g_try(v2:"Result",loop:true) */
            Result = F_Language_printbox_list2(_Za)
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC(")")
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
      } 
    return Result} 
  
// The EID go function for: self_print @ Call (throw: true) 
func E_self_print_Call_Language (self EID) EID { 
    return To_Call(OBJ(self)).SelfPrint( )} 
  
/* {1} The go function for: self_print(self:Call+) [status=1] */
func (self *Call_plus ) SelfPrint () EID { 
    var Result EID 
    /*g_try(v2:"Result",loop:true) */
    Result = F_printexp_any(self.Args.At(1-1),CTRUE)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(".")
    /*g_try(v2:"Result",loop:true) */
    Result = Core.F_print_any(self.Selector.Id())
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("")
    Result = EVOID
    }}
    return Result} 
  
// The EID go function for: self_print @ Call+ (throw: true) 
func E_self_print_Call_plus_Language (self EID) EID { 
    return To_Call_plus(OBJ(self)).SelfPrint( )} 
  
/* {1} The go function for: self_eval(self:Call) [status=1] */
func (self *Call ) SelfEval () EID { 
    var Result EID 
    { var start int  = ClEnv.Index
      { var p *ClaireProperty   = self.Selector
        if (ClEnv.Debug_I >= 0) { 
          C_iClaire_LastCall.Value = self.Id()
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
              { 
                var arg_2 EID 
                /*g_try(v2:"arg_2",loop:false) */
                arg_2 = EVAL(x)
                /* ERROR PROTECTION INSERTED (arg_2-loop_1) */
                if ErrorIn(arg_2) {loop_1 = arg_2
                } else {
                loop_1 = ClEnv.Push(arg_2)}
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
          { 
            var rx EID 
            /*g_try(v2:"rx",loop:false) */
            rx = Core.F_eval_message_property(p,Core.F_find_which_property(p,start,OWNER(ClEnv.EvalStack[start])),start,CTRUE)
            /* ERROR PROTECTION INSERTED (rx-Result) */
            if ErrorIn(rx) {Result = rx
            } else {
            C_iClaire_LastCall.Value = self.Id()
            Result = rx}
            } 
          }
          } else {
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
              var loop_3 EID 
              _ = loop_3
              /*g_try(v2:"loop_3",loop:tuple("Result", EID)) */
              { 
                var arg_4 EID 
                /*g_try(v2:"arg_4",loop:false) */
                arg_4 = EVAL(x)
                /* ERROR PROTECTION INSERTED (arg_4-loop_3) */
                if ErrorIn(arg_4) {loop_3 = arg_4
                } else {
                loop_3 = ClEnv.Push(arg_4)}
                } 
              /* ERROR PROTECTION INSERTED (loop_3-Result) */
              if ErrorIn(loop_3) {Result = loop_3
              break
              } else {
              }
              } 
            } 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          Result = Core.F_eval_message_property(p,Core.F_find_which_property(p,start,OWNER(ClEnv.EvalStack[start])),start,CTRUE)
          }
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: self_eval @ Call (throw: true) 
func E_self_eval_Call (self EID) EID { 
    return To_Call(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Call 
func EVAL_Call (x *ClaireAny) EID { 
     return To_Call(x).SelfEval()} 
  
// read slot : 
/* {1} The go function for: self_eval(self:Call+) [status=1] */
func (self *Call_plus ) SelfEval () EID { 
    var Result EID 
    { 
      var arg_1 EID 
      /*g_try(v2:"arg_1",loop:false) */
      arg_1 = EVAL(self.Args.At(1-1))
      /* ERROR PROTECTION INSERTED (arg_1-Result) */
      if ErrorIn(arg_1) {Result = arg_1
      } else {
      Result = self.Selector.ReadEID(arg_1)}
      } 
    return Result} 
  
// The EID go function for: self_eval @ Call+ (throw: true) 
func E_self_eval_Call_plus (self EID) EID { 
    return To_Call_plus(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Call+ 
func EVAL_Call_plus (x *ClaireAny) EID { 
     return To_Call_plus(x).SelfEval()} 
  
// dumb : we need to repeat
/* {1} The go function for: self_eval(self:Call*) [status=1] */
func (self *Call_star ) SelfEval () EID { 
    var Result EID 
    { var start int  = ClEnv.Index
      { var p *ClaireProperty   = self.Selector
        if (ClEnv.Debug_I >= 0) { 
          C_iClaire_LastCall.Value = self.Id()
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
              { 
                var arg_2 EID 
                /*g_try(v2:"arg_2",loop:false) */
                arg_2 = EVAL(x)
                /* ERROR PROTECTION INSERTED (arg_2-loop_1) */
                if ErrorIn(arg_2) {loop_1 = arg_2
                } else {
                loop_1 = ClEnv.Push(arg_2)}
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
          { 
            var rx EID 
            /*g_try(v2:"rx",loop:false) */
            rx = Core.F_eval_message_property(p,Core.F_find_which_property(p,start,OWNER(ClEnv.EvalStack[start])),start,CTRUE)
            /* ERROR PROTECTION INSERTED (rx-Result) */
            if ErrorIn(rx) {Result = rx
            } else {
            C_iClaire_LastCall.Value = self.Id()
            Result = rx}
            } 
          }
          } else {
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
              var loop_3 EID 
              _ = loop_3
              /*g_try(v2:"loop_3",loop:tuple("Result", EID)) */
              { 
                var arg_4 EID 
                /*g_try(v2:"arg_4",loop:false) */
                arg_4 = EVAL(x)
                /* ERROR PROTECTION INSERTED (arg_4-loop_3) */
                if ErrorIn(arg_4) {loop_3 = arg_4
                } else {
                loop_3 = ClEnv.Push(arg_4)}
                } 
              /* ERROR PROTECTION INSERTED (loop_3-Result) */
              if ErrorIn(loop_3) {Result = loop_3
              break
              } else {
              }
              } 
            } 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          Result = Core.F_eval_message_property(p,Core.F_find_which_property(p,start,OWNER(ClEnv.EvalStack[start])),start,CTRUE)
          }
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: self_eval @ Call* (throw: true) 
func E_self_eval_Call_star (self EID) EID { 
    return To_Call_star(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Call* 
func EVAL_Call_star (x *ClaireAny) EID { 
     return To_Call_star(x).SelfEval()} 
  
// recursive printing of bicall
/* {1} The go function for: printe(self:any,s:property) [status=1] */
func F_printe_any (self *ClaireAny ,s *ClaireProperty ) EID { 
    var Result EID 
    var g0071I *ClaireBoolean  
    if (self.Isa.IsIn(C_Call) == CTRUE) { 
      { var g0070 *Call   = To_Call(self)
        g0071I = MakeBoolean((g0070.Selector.Isa.IsIn(C_operation) == CTRUE) && (g0070.Args.Length() == 2))
        } 
      } else {
      g0071I = CFALSE
      } 
    if (g0071I == CTRUE) { 
      if (CTRUE == CTRUE) { 
        PRINC("(")
        /*g_try(v2:"Result",loop:true) */
        Result = Core.F_CALL(C_print,ARGS(self.ToEID()))
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC(")")
        Result = EVOID
        }
        } else {
        Result = F_printexp_any(self,CTRUE)
        } 
      } else {
      Result = F_printexp_any(self,CTRUE)
      } 
    return Result} 
  
// The EID go function for: printe @ any (throw: true) 
func E_printe_any (self EID,s EID) EID { 
    return F_printe_any(ANY(self),ToProperty(OBJ(s)) )} 
  
// tells if the sugar :op can be used
// x(x2) = o(a:x(x2), y) =>  x(x2) :o y
/* {1} The go function for: sugar?(x:any,x2:any,o:any,a:any) [status=0] */
func F_sugar_ask_any (x *ClaireAny ,x2 *ClaireAny ,o *ClaireAny ,a *ClaireAny ) *ClaireBoolean  { 
    // procedure body with s = boolean 
var Result *ClaireBoolean  
    if (o.Isa.IsIn(C_operation) == CTRUE) { 
      if (x.Isa.IsIn(C_property) == CTRUE) { 
        { var g0073 *ClaireProperty   = ToProperty(x)
          _ = g0073
          if (a.Isa.IsIn(C_Call) == CTRUE) { 
            { var g0074 *Call   = To_Call(a)
              Result = MakeBoolean((g0073.Id() == g0074.Selector.Id()) && (Equal(g0074.Args.At(1-1),x2) == CTRUE))
              } 
            } else {
            Result = CFALSE
            } 
          } 
        } else {
        Result = CFALSE
        } 
      } else {
      Result = CFALSE
      } 
    return Result} 
  
// The EID go function for: sugar? @ any (throw: false) 
func E_sugar_ask_any (x EID,x2 EID,o EID,a EID) EID { 
    return EID{F_sugar_ask_any(ANY(x),
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
/* {1} The go function for: self_print(self:Assign) [status=1] */
func (self *Assign ) SelfPrint () EID { 
    var Result EID 
    { var a *ClaireAny   = self.Arg
      { var o *ClaireAny  
        if (a.Isa.IsIn(C_Call) == CTRUE) { 
          { var g0078 *Call   = To_Call(a)
            _ = g0078
            o = g0078.Selector.Id()
            } 
          } else {
          o = CFALSE.Id()
          } 
        /*g_try(v2:"Result",loop:true) */
        var g0080I *ClaireBoolean  
        { var arg_1 *ClaireAny  
          _ = arg_1
          if (a.Isa.IsIn(C_Call) == CTRUE) { 
            { var g0079 *Call   = To_Call(a)
              _ = g0079
              arg_1 = g0079.Args.At(1-1)
              } 
            } else {
            arg_1 = CFALSE.Id()
            } 
          g0080I = F_sugar_ask_any(self.ClaireVar,CEMPTY.Id(),o,arg_1)
          } 
        if (g0080I == CTRUE) { 
          /*g_try(v2:"Result",loop:true) */
          Result = Core.F_CALL(C_print,ARGS(self.ClaireVar.ToEID()))
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(" :")
          /*g_try(v2:"Result",loop:true) */
          Result = Core.F_CALL(C_print,ARGS(o.ToEID()))
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(" ")
          /*g_try(v2:"Result",loop:true) */
          Result = F_lbreak_integer(2)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          /*g_try(v2:"Result",loop:true) */
          Result = F_printexp_any(ToList(OBJ(Core.F_CALL(C_args,ARGS(a.ToEID())))).At(2-1),CTRUE)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC("")
          Result = EVOID
          }}}}
          } else {
          /*g_try(v2:"Result",loop:true) */
          Result = Core.F_CALL(C_print,ARGS(self.ClaireVar.ToEID()))
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(" := ")
          /*g_try(v2:"Result",loop:true) */
          Result = F_lbreak_integer(2)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          /*g_try(v2:"Result",loop:true) */
          Result = F_printexp_any(a,CTRUE)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC("")
          Result = EVOID
          }}}
          } 
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
        } 
      } 
    return Result} 
  
// The EID go function for: self_print @ Assign (throw: true) 
func E_self_print_Assign_Language (self EID) EID { 
    return To_Assign(OBJ(self)).SelfPrint( )} 
  
/* {1} The go function for: self_eval(self:Assign) [status=1] */
func (self *Assign ) SelfEval () EID { 
    var Result EID 
    { 
      var arg_1 EID 
      /*g_try(v2:"arg_1",loop:false) */
      arg_1 = EVAL(self.Arg)
      /* ERROR PROTECTION INSERTED (arg_1-Result) */
      if ErrorIn(arg_1) {Result = arg_1
      } else {
      Result = To_Variable(self.ClaireVar).WriteEID(arg_1)}
      } 
    return Result} 
  
// The EID go function for: self_eval @ Assign (throw: true) 
func E_self_eval_Assign (self EID) EID { 
    return To_Assign(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Assign 
func EVAL_Assign (x *ClaireAny) EID { 
     return To_Assign(x).SelfEval()} 
  
// global variables
//
/* {1} The go function for: self_print(self:Gassign) [status=1] */
func (self *Gassign ) SelfPrint () EID { 
    var Result EID 
    { var a *ClaireAny   = self.Arg
      { var o *ClaireAny  
        if (a.Isa.IsIn(C_Call) == CTRUE) { 
          { var g0082 *Call   = To_Call(a)
            _ = g0082
            o = g0082.Selector.Id()
            } 
          } else {
          o = CFALSE.Id()
          } 
        /*g_try(v2:"Result",loop:true) */
        var g0084I *ClaireBoolean  
        { var arg_1 *ClaireAny  
          _ = arg_1
          if (a.Isa.IsIn(C_Call) == CTRUE) { 
            { var g0083 *Call   = To_Call(a)
              _ = g0083
              arg_1 = g0083.Args.At(1-1)
              } 
            } else {
            arg_1 = CFALSE.Id()
            } 
          g0084I = F_sugar_ask_any(self.ClaireVar.Id(),CEMPTY.Id(),o,arg_1)
          } 
        if (g0084I == CTRUE) { 
          /*g_try(v2:"Result",loop:true) */
          Result = Core.F_print_any(self.ClaireVar.Id())
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(" :")
          /*g_try(v2:"Result",loop:true) */
          Result = Core.F_CALL(C_print,ARGS(o.ToEID()))
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(" ")
          /*g_try(v2:"Result",loop:true) */
          Result = F_lbreak_integer(2)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          /*g_try(v2:"Result",loop:true) */
          Result = Core.F_CALL(C_print,ARGS(ToList(OBJ(Core.F_CALL(C_args,ARGS(a.ToEID())))).At(2-1).ToEID()))
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC("")
          Result = EVOID
          }}}}
          } else {
          /*g_try(v2:"Result",loop:true) */
          Result = Core.F_print_any(self.ClaireVar.Id())
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(" := ")
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
          }}}
          } 
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
        } 
      } 
    return Result} 
  
// The EID go function for: self_print @ Gassign (throw: true) 
func E_self_print_Gassign_Language (self EID) EID { 
    return To_Gassign(OBJ(self)).SelfPrint( )} 
  
/* {1} The go function for: self_eval(self:Gassign) [status=1] */
func (self *Gassign ) SelfEval () EID { 
    var Result EID 
    { var v *Core.GlobalVariable   = self.ClaireVar
      _ = v
      { var arg_1 *ClaireAny  
        _ = arg_1
        var try_2 EID 
        /*g_try(v2:"try_2",loop:false) */
        try_2 = EVAL(self.Arg)
        /* ERROR PROTECTION INSERTED (arg_1-Result) */
        if ErrorIn(try_2) {Result = try_2
        } else {
        arg_1 = ANY(try_2)
        Result = F_write_value_global_variable(v,arg_1)
        }
        } 
      } 
    return Result} 
  
// The EID go function for: self_eval @ Gassign (throw: true) 
func E_self_eval_Gassign (self EID) EID { 
    return To_Gassign(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Gassign 
func EVAL_Gassign (x *ClaireAny) EID { 
     return To_Gassign(x).SelfEval()} 
  
//--------------- BOOLEAN OPERATIONS ---------------------------------
// "and" is strictly boolean and is based on short-circuit evaluation.
//
/* {1} The go function for: self_print(self:And) [status=1] */
func (self *And ) SelfPrint () EID { 
    var Result EID 
    PRINC("(")
    /*g_try(v2:"Result",loop:true) */
    Result = F_Language_printbox_list3(self.Args,MakeString(" & "))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(")")
    Result = EVOID
    }
    return Result} 
  
// The EID go function for: self_print @ And (throw: true) 
func E_self_print_And_Language (self EID) EID { 
    return To_And(OBJ(self)).SelfPrint( )} 
  
/* {1} The go function for: self_eval(self:And) [status=1] */
func (self *And ) SelfEval () EID { 
    var Result EID 
    { var arg_1 *ClaireAny  
      _ = arg_1
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
          var g0085I *ClaireBoolean  
          var try_4 EID 
          /*g_try(v2:"try_4",loop:false) */
          { var arg_5 *ClaireBoolean  
            _ = arg_5
            var try_6 EID 
            /*g_try(v2:"try_6",loop:false) */
            { var arg_7 *ClaireAny  
              _ = arg_7
              var try_8 EID 
              /*g_try(v2:"try_8",loop:false) */
              try_8 = EVAL(x)
              /* ERROR PROTECTION INSERTED (arg_7-try_6) */
              if ErrorIn(try_8) {try_6 = try_8
              } else {
              arg_7 = ANY(try_8)
              try_6 = EID{F_boolean_I_any(arg_7).Id(),0}
              }
              } 
            /* ERROR PROTECTION INSERTED (arg_5-try_4) */
            if ErrorIn(try_6) {try_4 = try_6
            } else {
            arg_5 = ToBoolean(OBJ(try_6))
            try_4 = EID{Core.F__I_equal_any(arg_5.Id(),CTRUE.Id()).Id(),0}
            }
            } 
          /* ERROR PROTECTION INSERTED (g0085I-loop_3) */
          if ErrorIn(try_4) {loop_3 = try_4
          } else {
          g0085I = ToBoolean(OBJ(try_4))
          if (g0085I == CTRUE) { 
            try_2 = EID{CTRUE.Id(),0}
            break
            } else {
            loop_3 = EID{CFALSE.Id(),0}
            } 
          }
          /* ERROR PROTECTION INSERTED (loop_3-try_2) */
          if ErrorIn(loop_3) {try_2 = loop_3
          break
          } else {
          }
          } 
        } 
      /* ERROR PROTECTION INSERTED (arg_1-Result) */
      if ErrorIn(try_2) {Result = try_2
      } else {
      arg_1 = ANY(try_2)
      Result = EID{Core.F_not_any(arg_1).Id(),0}
      }
      } 
    return Result} 
  
// The EID go function for: self_eval @ And (throw: true) 
func E_self_eval_And (self EID) EID { 
    return To_And(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: And 
func EVAL_And (x *ClaireAny) EID { 
     return To_And(x).SelfEval()} 
  
// or expression
//
/* {1} The go function for: self_print(self:Or) [status=1] */
func (self *Or ) SelfPrint () EID { 
    var Result EID 
    PRINC("(")
    /*g_try(v2:"Result",loop:true) */
    Result = F_Language_printbox_list3(self.Args,MakeString(" | "))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(")")
    Result = EVOID
    }
    return Result} 
  
// The EID go function for: self_print @ Or (throw: true) 
func E_self_print_Or_Language (self EID) EID { 
    return To_Or(OBJ(self)).SelfPrint( )} 
  
/* {1} The go function for: self_eval(self:Or) [status=1] */
func (self *Or ) SelfEval () EID { 
    var Result EID 
    var g0086I *ClaireBoolean  
    var try_1 EID 
    /*g_try(v2:"try_1",loop:false) */
    { 
      var x *ClaireAny  
      _ = x
      try_1= EID{CFALSE.Id(),0}
      var x_support *ClaireList  
      x_support = self.Args
      x_len := x_support.Length()
      for i_it := 0; i_it < x_len; i_it++ { 
        x = x_support.At(i_it)
        var loop_2 EID 
        _ = loop_2
        /*g_try(v2:"loop_2",loop:tuple("try_1", EID)) */
        var g0087I *ClaireBoolean  
        var try_3 EID 
        /*g_try(v2:"try_3",loop:false) */
        { var arg_4 *ClaireAny  
          _ = arg_4
          var try_5 EID 
          /*g_try(v2:"try_5",loop:false) */
          try_5 = EVAL(x)
          /* ERROR PROTECTION INSERTED (arg_4-try_3) */
          if ErrorIn(try_5) {try_3 = try_5
          } else {
          arg_4 = ANY(try_5)
          try_3 = EID{F_boolean_I_any(arg_4).Id(),0}
          }
          } 
        /* ERROR PROTECTION INSERTED (g0087I-loop_2) */
        if ErrorIn(try_3) {loop_2 = try_3
        } else {
        g0087I = ToBoolean(OBJ(try_3))
        if (g0087I == CTRUE) { 
          try_1 = EID{CTRUE.Id(),0}
          break
          } else {
          loop_2 = EID{CFALSE.Id(),0}
          } 
        }
        /* ERROR PROTECTION INSERTED (loop_2-try_1) */
        if ErrorIn(loop_2) {try_1 = loop_2
        break
        } else {
        }
        } 
      } 
    /* ERROR PROTECTION INSERTED (g0086I-Result) */
    if ErrorIn(try_1) {Result = try_1
    } else {
    g0086I = ToBoolean(OBJ(try_1))
    if (g0086I == CTRUE) { 
      Result = EID{CTRUE.Id(),0}
      } else {
      Result = EID{CFALSE.Id(),0}
      } 
    }
    return Result} 
  
// The EID go function for: self_eval @ Or (throw: true) 
func E_self_eval_Or (self EID) EID { 
    return To_Or(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Or 
func EVAL_Or (x *ClaireAny) EID { 
     return To_Or(x).SelfEval()} 
  
// ----------------- an anti-evaluator ---------------------------------
//
/* {1} The go function for: self_print(self:Quote) [status=1] */
func (self *Quote ) SelfPrint () EID { 
    var Result EID 
    PRINC("quote(")
    /*g_try(v2:"Result",loop:true) */
    Result = Core.F_CALL(C_print,ARGS(self.Arg.ToEID()))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(")")
    Result = EVOID
    }
    return Result} 
  
// The EID go function for: self_print @ Quote (throw: true) 
func E_self_print_Quote_Language (self EID) EID { 
    return To_Quote(OBJ(self)).SelfPrint( )} 
  
/* {1} The go function for: self_eval(self:Quote) [status=0] */
func (self *Quote ) SelfEval () EID { 
    var Result EID 
    Result = self.Arg.ToEID()
    return Result} 
  
// The EID go function for: self_eval @ Quote (throw: true) 
func E_self_eval_Quote (self EID) EID { 
    return To_Quote(OBJ(self)).SelfEval( )} 
  
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
/* {1} The go function for: self_print(self:Call_method) [status=1] */
func (self *CallMethod ) SelfPrint () EID { 
    var Result EID 
    /*g_try(v2:"Result",loop:true) */
    Result = Core.F_print_any(self.Arg.Id())
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("(")
    /*g_try(v2:"Result",loop:true) */
    Result = Core.F_princ_list(self.Args)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(")")
    Result = EVOID
    }}
    return Result} 
  
// The EID go function for: self_print @ Call_method (throw: true) 
func E_self_print_Call_method_Language (self EID) EID { 
    return To_CallMethod(OBJ(self)).SelfPrint( )} 
  
/* {1} The go function for: self_eval(self:Call_method) [status=1] */
func (self *CallMethod ) SelfEval () EID { 
    var Result EID 
    { var start int  = ClEnv.Index
      _ = start
      { var Cprop *ClaireMethod   = self.Arg
        _ = Cprop
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
            { 
              var arg_2 EID 
              /*g_try(v2:"arg_2",loop:false) */
              arg_2 = EVAL(x)
              /* ERROR PROTECTION INSERTED (arg_2-loop_1) */
              if ErrorIn(arg_2) {loop_1 = arg_2
              } else {
              loop_1 = ClEnv.Push(arg_2)}
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
        Result = Core.F_execute_method(Cprop,start,CTRUE)
        }
        } 
      } 
    return Result} 
  
// The EID go function for: self_eval @ Call_method (throw: true) 
func E_self_eval_Call_method (self EID) EID { 
    return To_CallMethod(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Call_method 
func EVAL_Call_method (x *ClaireAny) EID { 
     return To_CallMethod(x).SelfEval()} 
  
// same thing with one only argument: we do not use the stack
/* {1} The go function for: self_eval(self:Call_method1) [status=1] */
func (self *CallMethod1 ) SelfEval () EID { 
    var Result EID 
    { var f *ClaireMethod   = self.Arg
      _ = f
      { var l *ClaireList   = self.Args
        _ = l
        { 
          var arg_1 EID 
          /*g_try(v2:"arg_1",loop:false) */
          arg_1 = EVAL(l.At(1-1))
          /* ERROR PROTECTION INSERTED (arg_1-Result) */
          if ErrorIn(arg_1) {Result = arg_1
          } else {
          Result = FASTCALL1(f,arg_1)}
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: self_eval @ Call_method1 (throw: true) 
func E_self_eval_Call_method1 (self EID) EID { 
    return To_CallMethod1(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Call_method1 
func EVAL_Call_method1 (x *ClaireAny) EID { 
     return To_CallMethod1(x).SelfEval()} 
  
// same thing with two arguments
/* {1} The go function for: self_eval(self:Call_method2) [status=1] */
func (self *CallMethod2 ) SelfEval () EID { 
    var Result EID 
    { var f *ClaireMethod   = self.Arg
      _ = f
      { var l *ClaireList   = self.Args
        { 
          var arg_1 EID 
          /*g_try(v2:"arg_1",loop:false) */
          arg_1 = EVAL(l.At(1-1))
          /* ERROR PROTECTION INSERTED (arg_1-Result) */
          if ErrorIn(arg_1) {Result = arg_1
          } else {
          var arg_2 EID 
          /*g_try(v2:"arg_2",loop:false) */
          arg_2 = EVAL(l.At(2-1))
          /* ERROR PROTECTION INSERTED (arg_2-Result) */
          if ErrorIn(arg_2) {Result = arg_2
          } else {
          Result = FASTCALL2(f,arg_1,arg_2)}}
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: self_eval @ Call_method2 (throw: true) 
func E_self_eval_Call_method2 (self EID) EID { 
    return To_CallMethod2(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Call_method2 
func EVAL_Call_method2 (x *ClaireAny) EID { 
     return To_CallMethod2(x).SelfEval()} 
  
// same thing with two arguments
/* {1} The go function for: self_eval(self:Call_method3) [status=1] */
func (self *Language_CallMethod3 ) SelfEval () EID { 
    var Result EID 
    { var f *ClaireMethod   = self.Arg
      _ = f
      { var l *ClaireList   = self.Args
        { 
          var arg_1 EID 
          /*g_try(v2:"arg_1",loop:false) */
          arg_1 = EVAL(l.At(1-1))
          /* ERROR PROTECTION INSERTED (arg_1-Result) */
          if ErrorIn(arg_1) {Result = arg_1
          } else {
          var arg_2 EID 
          /*g_try(v2:"arg_2",loop:false) */
          arg_2 = EVAL(l.At(2-1))
          /* ERROR PROTECTION INSERTED (arg_2-Result) */
          if ErrorIn(arg_2) {Result = arg_2
          } else {
          var arg_3 EID 
          /*g_try(v2:"arg_3",loop:false) */
          arg_3 = EVAL(l.At(3-1))
          /* ERROR PROTECTION INSERTED (arg_3-Result) */
          if ErrorIn(arg_3) {Result = arg_3
          } else {
          Result = FASTCALL3(f,arg_1,arg_2,arg_3)}}}
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: self_eval @ Call_method3 (throw: true) 
func E_self_eval_Call_method3 (self EID) EID { 
    return To_Language_CallMethod3(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Call_method3 
func EVAL_Language_Call_method3 (x *ClaireAny) EID { 
     return To_Language_CallMethod3(x).SelfEval()} 
  
// an instruction to read a slot
//
/* {1} The go function for: self_print(self:Call_slot) [status=1] */
func (self *CallSlot ) SelfPrint () EID { 
    var Result EID 
    /*g_try(v2:"Result",loop:true) */
    Result = Core.F_print_any(self.Selector.Id())
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("(")
    /*g_try(v2:"Result",loop:true) */
    Result = Core.F_CALL(C_print,ARGS(self.Arg.ToEID()))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(")")
    Result = EVOID
    }}
    return Result} 
  
// The EID go function for: self_print @ Call_slot (throw: true) 
func E_self_print_Call_slot_Language (self EID) EID { 
    return To_CallSlot(OBJ(self)).SelfPrint( )} 
  
/* {1} The go function for: self_eval(self:Call_slot) [status=1] */
func (self *CallSlot ) SelfEval () EID { 
    var Result EID 
    { var arg_1 *ClaireAny  
      _ = arg_1
      var try_2 EID 
      /*g_try(v2:"try_2",loop:false) */
      try_2 = EVAL(self.Arg)
      /* ERROR PROTECTION INSERTED (arg_1-Result) */
      if ErrorIn(try_2) {Result = try_2
      } else {
      arg_1 = ANY(try_2)
      Result = Core.F_get_slot(self.Selector,ToObject(arg_1)).ToEID()
      }
      } 
    return Result} 
  
// The EID go function for: self_eval @ Call_slot (throw: true) 
func E_self_eval_Call_slot (self EID) EID { 
    return To_CallSlot(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Call_slot 
func EVAL_Call_slot (x *ClaireAny) EID { 
     return To_CallSlot(x).SelfEval()} 
  
// an instruction to read an array
// selector is an exp with type array, arg is an exp with type integer, and test
// contains the inferred member_type of the array
//
/* {1} The go function for: self_print(self:Call_array) [status=1] */
func (self *CallArray ) SelfPrint () EID { 
    var Result EID 
    /*g_try(v2:"Result",loop:true) */
    Result = Core.F_CALL(C_print,ARGS(self.Selector.ToEID()))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("[")
    /*g_try(v2:"Result",loop:true) */
    Result = Core.F_CALL(C_print,ARGS(self.Arg.ToEID()))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("]")
    Result = EVOID
    }}
    return Result} 
  
// The EID go function for: self_print @ Call_array (throw: true) 
func E_self_print_Call_array_Language (self EID) EID { 
    return To_CallArray(OBJ(self)).SelfPrint( )} 
  
/* {1} The go function for: self_eval(self:Call_array) [status=1] */
func (self *CallArray ) SelfEval () EID { 
    var Result EID 
    { var arg_1 *ClaireAny  
      _ = arg_1
      var try_3 EID 
      /*g_try(v2:"try_3",loop:false) */
      try_3 = EVAL(self.Selector)
      /* ERROR PROTECTION INSERTED (arg_1-Result) */
      if ErrorIn(try_3) {Result = try_3
      } else {
      arg_1 = ANY(try_3)
      { var arg_2 *ClaireAny  
        _ = arg_2
        var try_4 EID 
        /*g_try(v2:"try_4",loop:false) */
        try_4 = EVAL(self.Arg)
        /* ERROR PROTECTION INSERTED (arg_2-Result) */
        if ErrorIn(try_4) {Result = try_4
        } else {
        arg_2 = ANY(try_4)
        Result = Core.F_nth_array(ToArray(arg_1),ToInteger(arg_2).Value)
        }
        } 
      }
      } 
    return Result} 
  
// The EID go function for: self_eval @ Call_array (throw: true) 
func E_self_eval_Call_array (self EID) EID { 
    return To_CallArray(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Call_array 
func EVAL_Call_array (x *ClaireAny) EID { 
     return To_CallArray(x).SelfEval()} 
  
// an instruction to read a table
//
/* {1} The go function for: self_print(self:Call_table) [status=1] */
func (self *CallTable ) SelfPrint () EID { 
    var Result EID 
    /*g_try(v2:"Result",loop:true) */
    Result = Core.F_print_any(self.Selector.Id())
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("[")
    /*g_try(v2:"Result",loop:true) */
    Result = Core.F_CALL(C_print,ARGS(self.Arg.ToEID()))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("]")
    Result = EVOID
    }}
    return Result} 
  
// The EID go function for: self_print @ Call_table (throw: true) 
func E_self_print_Call_table_Language (self EID) EID { 
    return To_CallTable(OBJ(self)).SelfPrint( )} 
  
/* {1} The go function for: self_eval(self:Call_table) [status=1] */
func (self *CallTable ) SelfEval () EID { 
    var Result EID 
    if (self.Test == CTRUE) { 
      { var arg_1 *ClaireAny  
        _ = arg_1
        var try_2 EID 
        /*g_try(v2:"try_2",loop:false) */
        try_2 = EVAL(self.Arg)
        /* ERROR PROTECTION INSERTED (arg_1-Result) */
        if ErrorIn(try_2) {Result = try_2
        } else {
        arg_1 = ANY(try_2)
        Result = Core.F_nth_table1(self.Selector,arg_1)
        }
        } 
      } else {
      { var arg_3 *ClaireAny  
        _ = arg_3
        var try_4 EID 
        /*g_try(v2:"try_4",loop:false) */
        try_4 = EVAL(self.Arg)
        /* ERROR PROTECTION INSERTED (arg_3-Result) */
        if ErrorIn(try_4) {Result = try_4
        } else {
        arg_3 = ANY(try_4)
        Result = Core.F_get_table(self.Selector,arg_3).ToEID()
        }
        } 
      } 
    return Result} 
  
// The EID go function for: self_eval @ Call_table (throw: true) 
func E_self_eval_Call_table (self EID) EID { 
    return To_CallTable(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Call_table 
func EVAL_Call_table (x *ClaireAny) EID { 
     return To_CallTable(x).SelfEval()} 
  
// Update = [R(x) := y] where R(x) is a Call_slot, a call_array or a call_table 
// the actual "meta" property that represents := is stored in self.arg
// the structure is complex: see ocall.cl
// self.var is the writable container (call_slot, call_array, call_table)
// R(x)
/* {1} The go function for: self_print(self:Update) [status=1] */
func (self *Update ) SelfPrint () EID { 
    var Result EID 
    /*g_try(v2:"Result",loop:true) */
    Result = Core.F_CALL(C_print,ARGS(self.Selector.ToEID()))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("(")
    /*g_try(v2:"Result",loop:true) */
    Result = Core.F_CALL(C_print,ARGS(Core.F_CALL(C_arg,ARGS(self.ClaireVar.ToEID()))))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(") := ")
    /*g_try(v2:"Result",loop:true) */
    Result = Core.F_CALL(C_print,ARGS(self.Value.ToEID()))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("")
    Result = EVOID
    }}}
    return Result} 
  
// The EID go function for: self_print @ Update (throw: true) 
func E_self_print_Update_Language (self EID) EID { 
    return To_Update(OBJ(self)).SelfPrint( )} 
  
/* {1} The go function for: self_eval(self:Update) [status=1] */
func (self *Update ) SelfEval () EID { 
    var Result EID 
    { var s *ClaireAny   = self.Selector
      /*g_try(v2:"Result",loop:true) */
      if (s.Isa.IsIn(C_property) == CTRUE) { 
        { var g0088 *ClaireProperty   = ToProperty(s)
          _ = g0088
          { var arg_1 *ClaireAny  
            _ = arg_1
            var try_3 EID 
            /*g_try(v2:"try_3",loop:false) */
            try_3 = EVAL(ANY(Core.F_CALL(C_arg,ARGS(self.ClaireVar.ToEID()))))
            /* ERROR PROTECTION INSERTED (arg_1-Result) */
            if ErrorIn(try_3) {Result = try_3
            } else {
            arg_1 = ANY(try_3)
            { var arg_2 *ClaireAny  
              _ = arg_2
              var try_4 EID 
              /*g_try(v2:"try_4",loop:false) */
              try_4 = EVAL(self.Value)
              /* ERROR PROTECTION INSERTED (arg_2-Result) */
              if ErrorIn(try_4) {Result = try_4
              } else {
              arg_2 = ANY(try_4)
              Result = Core.F_put_property2(g0088,ToObject(arg_1),arg_2)
              }
              } 
            }
            } 
          } 
        }  else if (C_table.Id() == s.Isa.Id()) { 
        { var g0089 *ClaireTable   = ToTable(s)
          _ = g0089
          { var arg_5 *ClaireAny  
            _ = arg_5
            var try_7 EID 
            /*g_try(v2:"try_7",loop:false) */
            try_7 = EVAL(ANY(Core.F_CALL(C_arg,ARGS(self.ClaireVar.ToEID()))))
            /* ERROR PROTECTION INSERTED (arg_5-Result) */
            if ErrorIn(try_7) {Result = try_7
            } else {
            arg_5 = ANY(try_7)
            { var arg_6 *ClaireAny  
              _ = arg_6
              var try_8 EID 
              /*g_try(v2:"try_8",loop:false) */
              try_8 = EVAL(self.Value)
              /* ERROR PROTECTION INSERTED (arg_6-Result) */
              if ErrorIn(try_8) {Result = try_8
              } else {
              arg_6 = ANY(try_8)
              Result = Core.F_nth_equal_table1(g0089,arg_5,arg_6)
              }
              } 
            }
            } 
          } 
        } else {
        Result = EID{CFALSE.Id(),0}
        } 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = EID{CNULL,0}
      }
      } 
    return Result} 
  
// The EID go function for: self_eval @ Update (throw: true) 
func E_self_eval_Update (self EID) EID { 
    return To_Update(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Update 
func EVAL_Update (x *ClaireAny) EID { 
     return To_Update(x).SelfEval()} 
  
// ------------------ SUPER: a jump in the set lattice ---------------
// A "super" allows one to execute a message as if the type of the receiver
// was a given abstract_class.
// However we require that the receiver be in the specified abstract_class.
// The form of the super is: SELECTOR@ABSTRACT_CLASS(RECEIVER , ...)
//
/* {1} The go function for: self_print(self:Super) [status=1] */
func (self *Super ) SelfPrint () EID { 
    var Result EID 
    { var _Zl int  = Core.C_pretty.Index
      _ = _Zl
      { var _Zs *ClaireProperty   = self.Selector
        _ = _Zs
        { var _Za *ClaireList   = self.Args
          _ = _Za
          /*g_try(v2:"Result",loop:true) */
          /*g_try(v2:"Result",loop:true) */
          Result = Core.F_print_any(_Zs.Id())
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC("@")
          /*g_try(v2:"Result",loop:true) */
          Result = Core.F_print_any(self.CastTo.Id())
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC("(")
          F_set_level_void()
          /*g_try(v2:"Result",loop:true) */
          Result = F_Language_printbox_list2(_Za)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(")")
          Result = EVOID
          }}}
          {
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
      } 
    return Result} 
  
// The EID go function for: self_print @ Super (throw: true) 
func E_self_print_Super_Language (self EID) EID { 
    return To_Super(OBJ(self)).SelfPrint( )} 
  
/* {1} The go function for: self_eval(self:Super) [status=1] */
func (self *Super ) SelfEval () EID { 
    var Result EID 
    { var start int  = ClEnv.Index
      { var t *ClaireType   = self.CastTo
        _ = t
        { var c *ClaireClass   = t.Class_I()
          _ = c
          { var p *ClaireProperty   = self.Selector
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
                { 
                  var arg_2 EID 
                  /*g_try(v2:"arg_2",loop:false) */
                  arg_2 = EVAL(x)
                  /* ERROR PROTECTION INSERTED (arg_2-loop_1) */
                  if ErrorIn(arg_2) {loop_1 = arg_2
                  } else {
                  loop_1 = ClEnv.Push(arg_2)}
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
            Result = Core.F_eval_message_property(p,Core.F_find_which_class(c,p.Definition,start,ClEnv.Index),start,CTRUE)
            }
            } 
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: self_eval @ Super (throw: true) 
func E_self_eval_Super (self EID) EID { 
    return To_Super(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Super 
func EVAL_Super (x *ClaireAny) EID { 
     return To_Super(x).SelfEval()} 
  
//--------------- comments ------------------------------------------
// the cast is the new form of simple super
//
/* {1} The go function for: self_print(x:Cast) [status=1] */
func (x *Cast ) SelfPrint () EID { 
    var Result EID 
    /*g_try(v2:"Result",loop:true) */
    Result = F_printexp_any(x.Arg,CFALSE)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(" as ")
    /*g_try(v2:"Result",loop:true) */
    Result = F_printexp_any(x.SetArg.Id(),CFALSE)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("")
    Result = EVOID
    }}
    return Result} 
  
// The EID go function for: self_print @ Cast (throw: true) 
func E_self_print_Cast_Language (x EID) EID { 
    return To_Cast(OBJ(x)).SelfPrint( )} 
  
/* {1} The go function for: self_eval(self:Cast) [status=1] */
func (self *Cast ) SelfEval () EID { 
    var Result EID 
    { var x *ClaireAny  
      var try_1 EID 
      /*g_try(v2:"try_1",loop:false) */
      try_1 = EVAL(self.Arg)
      /* ERROR PROTECTION INSERTED (x-Result) */
      if ErrorIn(try_1) {Result = try_1
      } else {
      x = ANY(try_1)
      { var y *ClaireType   = self.SetArg
        var g0091I *ClaireBoolean  
        if (y.Isa.IsIn(C_Param) == CTRUE) { 
          { var g0090 *ClaireParam   = To_Param(y.Id())
            g0091I = MakeBoolean(((g0090.Arg.Id() == C_list.Id()) || 
                (g0090.Arg.Id() == C_set.Id())) && (C_set.Id() == g0090.Args.At(1-1).Isa.Id()) && (Core.F__Z_any1(x,g0090.Arg) == CTRUE))
            } 
          } else {
          g0091I = CFALSE
          } 
        if (g0091I == CTRUE) { 
          { var arg_2 *ClaireAny  
            _ = arg_2
            var try_3 EID 
            /*g_try(v2:"try_3",loop:false) */
            try_3 = Core.F_the_type(ToType(To_Param(y.Id()).Args.At(1-1)))
            /* ERROR PROTECTION INSERTED (arg_2-Result) */
            if ErrorIn(try_3) {Result = try_3
            } else {
            arg_2 = ANY(try_3)
            Result = Core.F_check_in_bag(ToBag(x),C_bag,ToType(arg_2))
            }
            } 
          } else {
          Result = Core.F_check_in_any(x,y)
          } 
        } 
      }
      } 
    return Result} 
  
// The EID go function for: self_eval @ Cast (throw: true) 
func E_self_eval_Cast (self EID) EID { 
    return To_Cast(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Cast 
func EVAL_Cast (x *ClaireAny) EID { 
     return To_Cast(x).SelfEval()} 
  
// v3.3.16 */
// ----------------- return from a loop --------------------------------
//
// return_error is an exception that is handled by the "for" family
// of structures
//
/* {1} The go function for: self_print(self:Return) [status=1] */
func (self *Return ) SelfPrint () EID { 
    var Result EID 
    PRINC("break(")
    Core.C_pretty.Index = (Core.C_pretty.Index+2)
    /*integer->integer*//*g_try(v2:"Result",loop:true) */
    Result = Core.F_CALL(C_print,ARGS(self.Arg.ToEID()))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    Core.C_pretty.Index = (Core.C_pretty.Index-2)
    /*integer->integer*/PRINC(")")
    Result = EVOID
    }
    return Result} 
  
// The EID go function for: self_print @ Return (throw: true) 
func E_self_print_Return_Language (self EID) EID { 
    return To_Return(OBJ(self)).SelfPrint( )} 
  
/* {1} The go function for: self_eval(self:Return) [status=1] */
func (self *Return ) SelfEval () EID { 
    var Result EID 
    { var arg_1 *ClaireAny  
      _ = arg_1
      var try_2 EID 
      /*g_try(v2:"try_2",loop:false) */
      { var arg_3 *ClaireAny  
        _ = arg_3
        var try_4 EID 
        /*g_try(v2:"try_4",loop:false) */
        try_4 = EVAL(self.Arg)
        /* ERROR PROTECTION INSERTED (arg_3-try_2) */
        if ErrorIn(try_4) {try_2 = try_4
        } else {
        arg_3 = ANY(try_4)
        try_2 = Core.C_return_error.Make(arg_3).ToEID()
        }
        } 
      /* ERROR PROTECTION INSERTED (arg_1-Result) */
      if ErrorIn(try_2) {Result = try_2
      } else {
      arg_1 = ANY(try_2)
      Result = ToException(arg_1).Close()
      }
      } 
    return Result} 
  
// The EID go function for: self_eval @ Return (throw: true) 
func E_self_eval_Return (self EID) EID { 
    return To_Return(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Return 
func EVAL_Return (x *ClaireAny) EID { 
     return To_Return(x).SelfEval()} 
  
// ****************************************************************
// *       Part 4: Miscellaneous on instructions                  *
// ****************************************************************
// substitute any variable with same name as x with the value val
/* {1} The go function for: substitution(self:any,x:Variable,val:any) [status=0] */
func F_substitution_any (self *ClaireAny ,x *ClaireVariable ,val *ClaireAny ) *ClaireAny  { 
    // procedure body with s = any 
var Result *ClaireAny  
    if (self.Isa.IsIn(C_Variable) == CTRUE) { 
      { var g0094 *ClaireVariable   = To_Variable(self)
        Result = IfThenElse((g0094.Pname.Id() == x.Pname.Id()),
          val,
          g0094.Id())
        } 
      }  else if (self.Isa.IsIn(C_list) == CTRUE) { 
      { var g0095 *ClaireList   = ToList(self)
        { var i int  = 1
          { var g0096 int  = g0095.Length()
            _ = g0096
            for (i <= g0096) { 
              /* While stat, v:"Result" loop:false */
              if ((g0095.At(i-1).Isa.IsIn(C_Variable) == CTRUE) || 
                  (g0095.At(i-1).Isa.IsIn(C_unbound_symbol) == CTRUE)) { 
                ToArray(g0095.Id()).NthPut(i,F_substitution_any(g0095.At(i-1),x,val))
                } else {
                F_substitution_any(g0095.At(i-1),x,val)
                } 
              i = (i+1)
              /* try?:false, v2:"v_while6" loop will be:tuple("Result", void) */
              } 
            } 
          } 
        Result = g0095.Id()
        } 
      }  else if (self.Isa.IsIn(C_unbound_symbol) == CTRUE) { 
      { var g0097 *ClaireUnboundSymbol   = ToUnboundSymbol(self)
        Result = IfThenElse((g0097.Name.Id() == x.Pname.Id()),
          val,
          g0097.Id())
        } 
      }  else if (self.Isa.IsIn(C_Instruction) == CTRUE) { 
      { var g0098 *ClaireInstruction   = To_Instruction(self)
        { 
          var s *ClaireSlot  
          _ = s
          var s_iter *ClaireAny  
          var s_support *ClaireList  
          s_support = g0098.Id().Isa.Slots
          for _,s_iter = range(s_support.ValuesO()){ 
            s = ToSlot(s_iter)
            { var y *ClaireAny   = Core.F_get_slot(s,ToObject(g0098.Id()))
              if ((y.Isa.IsIn(C_Variable) == CTRUE) || 
                  (y.Isa.IsIn(C_unbound_symbol) == CTRUE)) { 
                Core.F_put_slot(s,ToObject(g0098.Id()),F_substitution_any(y,x,val))
                } else {
                F_substitution_any(y,x,val)
                } 
              } 
            } 
          } 
        Result = g0098.Id()
        } 
      } else {
      Result = self
      } 
    return Result} 
  
// The EID go function for: substitution @ any (throw: false) 
func E_substitution_any (self EID,x EID,val EID) EID { 
    return F_substitution_any(ANY(self),To_Variable(OBJ(x)),ANY(val) ).ToEID()} 
  
// count the number of occurrences of x
/* {1} The go function for: occurrence(self:any,x:Variable) [status=0] */
func F_occurrence_any (self *ClaireAny ,x *ClaireVariable ) int { 
    // procedure body with s = integer 
var Result int 
    if (self.Isa.IsIn(C_Variable) == CTRUE) { 
      { var g0100 *ClaireVariable   = To_Variable(self)
        _ = g0100
        if (g0100.Pname.Id() == x.Pname.Id()) { 
          Result = 1
          } else {
          Result = 0
          } 
        } 
      }  else if (self.Isa.IsIn(C_list) == CTRUE) { 
      { var g0101 *ClaireList   = ToList(self)
        _ = g0101
        { var n int  = 0
          _ = n
          { var i int  = 1
            _ = i
            { var g0102 int  = g0101.Length()
              _ = g0102
              for (i <= g0102) { 
                /* While stat, v:"Result" loop:false */
                n = (n+F_occurrence_any(g0101.At(i-1),x))
                i = (i+1)
                /* try?:false, v2:"v_while7" loop will be:tuple("Result", void) */
                } 
              } 
            } 
          Result = n
          } 
        } 
      }  else if (self.Isa.IsIn(C_unbound_symbol) == CTRUE) { 
      { var g0103 *ClaireUnboundSymbol   = ToUnboundSymbol(self)
        _ = g0103
        if (g0103.Name.Id() == x.Pname.Id()) { 
          Result = 1
          } else {
          Result = 0
          } 
        } 
      }  else if (self.Isa.IsIn(C_Instruction) == CTRUE) { 
      { var g0104 *ClaireInstruction   = To_Instruction(self)
        _ = g0104
        { var n int  = 0
          _ = n
          { 
            var s *ClaireSlot  
            _ = s
            var s_iter *ClaireAny  
            var s_support *ClaireList  
            s_support = g0104.Id().Isa.Slots
            for _,s_iter = range(s_support.ValuesO()){ 
              s = ToSlot(s_iter)
              n = (n+F_occurrence_any(Core.F_get_slot(s,ToObject(g0104.Id())),x))
              } 
            } 
          Result = n
          } 
        } 
      } else {
      Result = 0
      } 
    return Result} 
  
// The EID go function for: occurrence @ any (throw: false) 
func E_occurrence_any (self EID,x EID) EID { 
    return EID{C__INT,IVAL(F_occurrence_any(ANY(self),To_Variable(OBJ(x)) ))}} 
  
// new version in CLAIRE4 : see if the variable is changed
/* {1} The go function for: occurchange(self:any,x:Variable) [status=0] */
func F_occurchange_any (self *ClaireAny ,x *ClaireVariable ) *ClaireBoolean  { 
    // procedure body with s = boolean 
var Result *ClaireBoolean  
    if (self.Isa.IsIn(C_Assign) == CTRUE) { 
      { var g0106 *Assign   = To_Assign(self)
        _ = g0106
        Result = Equal(ANY(Core.F_CALL(C_mClaire_pname,ARGS(g0106.ClaireVar.ToEID()))),x.Pname.Id())
        } 
      }  else if (self.Isa.IsIn(C_list) == CTRUE) { 
      { var g0107 *ClaireList   = ToList(self)
        _ = g0107
        { var arg_1 *ClaireAny  
          _ = arg_1
          { 
            var y *ClaireAny  
            _ = y
            arg_1= CFALSE.Id()
            var y_support *ClaireList  
            y_support = g0107
            y_len := y_support.Length()
            for i_it := 0; i_it < y_len; i_it++ { 
              y = y_support.At(i_it)
              if (F_occurchange_any(y,x) == CTRUE) { 
                arg_1 = CTRUE.Id()
                break
                } 
              } 
            } 
          Result = F_boolean_I_any(arg_1)
          } 
        } 
      }  else if (self.Isa.IsIn(C_Instruction) == CTRUE) { 
      { var g0108 *ClaireInstruction   = To_Instruction(self)
        { var arg_2 *ClaireAny  
          _ = arg_2
          { 
            var s *ClaireSlot  
            _ = s
            var s_iter *ClaireAny  
            arg_2= CFALSE.Id()
            var s_support *ClaireList  
            s_support = g0108.Id().Isa.Slots
            for _,s_iter = range(s_support.ValuesO()){ 
              s = ToSlot(s_iter)
              if (F_occurchange_any(Core.F_get_slot(s,ToObject(g0108.Id())),x) == CTRUE) { 
                arg_2 = CTRUE.Id()
                break
                } 
              } 
            } 
          Result = F_boolean_I_any(arg_2)
          } 
        } 
      } else {
      Result = CFALSE
      } 
    return Result} 
  
// The EID go function for: occurchange @ any (throw: false) 
func E_occurchange_any (self EID,x EID) EID { 
    return EID{F_occurchange_any(ANY(self),To_Variable(OBJ(x)) ).Id(),0}} 
  
// a variant in CLAIRE 4 that assumes that variable have reveived their lexical bind (index)
// it also does not count the variable itself in an assign
/* {1} The go function for: occurexact(self:any,x:Variable) [status=0] */
func F_Language_occurexact_any (self *ClaireAny ,x *ClaireVariable ) int { 
    // procedure body with s = integer 
var Result int 
    if (self.Isa.IsIn(C_Variable) == CTRUE) { 
      { var g0110 *ClaireVariable   = To_Variable(self)
        if ((g0110.Pname.Id() == x.Pname.Id()) && 
            (g0110.Index == x.Index)) { 
          Result = 1
          } else {
          Result = 0
          } 
        } 
      }  else if (self.Isa.IsIn(C_list) == CTRUE) { 
      { var g0111 *ClaireList   = ToList(self)
        _ = g0111
        { var n int  = 0
          _ = n
          { var i int  = 1
            _ = i
            { var g0112 int  = g0111.Length()
              _ = g0112
              for (i <= g0112) { 
                /* While stat, v:"Result" loop:false */
                n = (n+F_Language_occurexact_any(g0111.At(i-1),x))
                i = (i+1)
                /* try?:false, v2:"v_while7" loop will be:tuple("Result", void) */
                } 
              } 
            } 
          Result = n
          } 
        } 
      }  else if (self.Isa.IsIn(C_unbound_symbol) == CTRUE) { 
      Result = 0
      }  else if (self.Isa.IsIn(C_Assign) == CTRUE) { 
      { var g0114 *Assign   = To_Assign(self)
        _ = g0114
        Result = F_Language_occurexact_any(ANY(Core.F_CALL(C_value,ARGS(EID{g0114.Id(),0}))),x)
        } 
      }  else if (self.Isa.IsIn(C_Instruction) == CTRUE) { 
      { var g0115 *ClaireInstruction   = To_Instruction(self)
        _ = g0115
        { var n int  = 0
          _ = n
          { 
            var s *ClaireSlot  
            _ = s
            var s_iter *ClaireAny  
            var s_support *ClaireList  
            s_support = g0115.Id().Isa.Slots
            for _,s_iter = range(s_support.ValuesO()){ 
              s = ToSlot(s_iter)
              n = (n+F_Language_occurexact_any(Core.F_get_slot(s,ToObject(g0115.Id())),x))
              } 
            } 
          Result = n
          } 
        } 
      } else {
      Result = 0
      } 
    return Result} 
  
// The EID go function for: occurexact @ any (throw: false) 
func E_Language_occurexact_any (self EID,x EID) EID { 
    return EID{C__INT,IVAL(F_Language_occurexact_any(ANY(self),To_Variable(OBJ(x)) ))}} 
  
// makes a (deep) copy of the instruction self
//
/* {1} The go function for: instruction_copy(self:any) [status=0] */
func F_instruction_copy_any (self *ClaireAny ) *ClaireAny  { 
    // procedure body with s = any 
var Result *ClaireAny  
    if (self.Isa.IsIn(C_list) == CTRUE) { 
      { var g0117 *ClaireList   = ToList(self)
        { var l *ClaireList   = g0117.Copy()
          { var i int  = 1
            { var g0118 int  = g0117.Length()
              _ = g0118
              for (i <= g0118) { 
                /* While stat, v:"Result" loop:false */
                ToArray(l.Id()).NthPut(i,F_instruction_copy_any(g0117.At(i-1)))
                i = (i+1)
                /* try?:false, v2:"v_while7" loop will be:tuple("Result", void) */
                } 
              } 
            } 
          Result = l.Id()
          } 
        } 
      }  else if (self.Isa.IsIn(C_Variable) == CTRUE) { 
      { var g0119 *ClaireVariable   = To_Variable(self)
        _ = g0119
        Result = g0119.Id()
        } 
      }  else if (self.Isa.IsIn(C_Instruction) == CTRUE) { 
      { var g0120 *ClaireInstruction   = To_Instruction(self)
        { var o *ClaireInstruction   = To_Instruction(g0120.Copy().Id())
          { 
            var s *ClaireSlot  
            _ = s
            var s_iter *ClaireAny  
            var s_support *ClaireList  
            s_support = g0120.Id().Isa.Slots
            for _,s_iter = range(s_support.ValuesO()){ 
              s = ToSlot(s_iter)
              Core.F_put_slot(s,ToObject(o.Id()),F_instruction_copy_any(Core.F_get_slot(s,ToObject(g0120.Id()))))
              } 
            } 
          Result = o.Id()
          } 
        } 
      } else {
      Result = self
      } 
    return Result} 
  
// The EID go function for: instruction_copy @ any (throw: false) 
func E_instruction_copy_any (self EID) EID { 
    return F_instruction_copy_any(ANY(self) ).ToEID()} 
  