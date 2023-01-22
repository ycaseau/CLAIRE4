/***** CLAIRE Compilation of file /Users/ycaseau/Dropbox/src/clairev4.07/src/meta/call.cl 
         [version 4.0.07 / safety 5] Sunday 01-01-2023 08:56:16 *****/

package Language
import (_ "fmt"
	. "Kernel"
	"Core"
)

//-------- dumb function to prevent import errors --------
func import_g0064() { 
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
/* The go function for: self_print(self:Call) [status=1] */
func (self *Call) SelfPrint () EID { 
    var Result EID
    { var _Zl int = Core.C_pretty.Index
      { var _Zs *ClaireProperty = self.Selector
        { var _Za *ClaireList = self.Args
          if ((_Zs.Isa.IsIn(C_operation) == CTRUE) && 
              (_Za.Length() == 2)) { 
            Core.C_pretty.Index = (Core.C_pretty.Index+2)
            Result = F_printe_any(_Za.At(0),_Zs)
            if !ErrorIn(Result) {
            PRINC(" ")
            Result = Core.F_print_any(_Zs.Id())
            if !ErrorIn(Result) {
            PRINC(" ")
            Result = F_lbreak_void()
            if !ErrorIn(Result) {
            Result = F_printe_any(_Za.At(1),_Zs)
            }}}
            }  else if (_Zs.Id() == C_nth.Id()) { 
            if (_Za.Length() == 3) { 
              Result = F_printexp_any(_Za.At(0),CFALSE)
              if !ErrorIn(Result) {
              PRINC("[")
              Result = Core.F_CALL(C_print,ARGS(_Za.At(1).ToEID()))
              if !ErrorIn(Result) {
              PRINC(",")
              Result = Core.F_CALL(C_print,ARGS(_Za.At(2).ToEID()))
              if !ErrorIn(Result) {
              PRINC("]")
              Result = EVOID
              }}}
              }  else if (_Za.Length() == 1) { 
              Result = F_printexp_any(_Za.At(0),CFALSE)
              if !ErrorIn(Result) {
              PRINC("[]")
              Result = EVOID
              }
              } else {
              Result = F_printexp_any(_Za.At(0),CFALSE)
              if !ErrorIn(Result) {
              PRINC("[")
              if (_Za.Length() == 2) { 
                Result = Core.F_CALL(C_print,ARGS(_Za.At(1).ToEID()))
                } else {
                Result = EID{CFALSE.Id(),0}
                } 
              if !ErrorIn(Result) {
              PRINC("]")
              Result = EVOID
              }}
              } 
            }  else if ((_Zs.Id() == C_nth_equal.Id()) && 
              (_Za.Length() >= 3)) { 
            { var a *ClaireAny = _Za.At(2)
              { var o *ClaireAny
                if (a.Isa.IsIn(C_Call) == CTRUE) { 
                  { var g0066 *Call = To_Call(a)
                    o = g0066.Selector.Id()
                    } 
                  } else {
                  o = CFALSE.Id()
                  } 
                if (_Za.Length() == 4) { 
                  Result = F_printexp_any(_Za.At(0),CFALSE)
                  if !ErrorIn(Result) {
                  PRINC("[")
                  Result = Core.F_CALL(C_print,ARGS(_Za.At(1).ToEID()))
                  if !ErrorIn(Result) {
                  PRINC(",")
                  Result = Core.F_CALL(C_print,ARGS(a.ToEID()))
                  if !ErrorIn(Result) {
                  PRINC("] := ")
                  Result = F_lbreak_integer(2)
                  if !ErrorIn(Result) {
                  Result = Core.F_CALL(C_print,ARGS(_Za.At(3).ToEID()))
                  }}}}
                  } else {
                  var g0070I *ClaireBoolean
                  { var arg_1 *ClaireAny
                    if (a.Isa.IsIn(C_Call) == CTRUE) { 
                      { var g0067 *Call = To_Call(a)
                        arg_1 = g0067.Args.At(0)
                        } 
                      } else {
                      arg_1 = CFALSE.Id()
                      } 
                    g0070I = F_sugar_ask_any(_Za.At(0),_Za.At(1),o,arg_1)
                    } 
                  if (g0070I == CTRUE) { 
                    Result = Core.F_CALL(C_print,ARGS(_Za.At(0).ToEID()))
                    if !ErrorIn(Result) {
                    PRINC("[")
                    Result = Core.F_CALL(C_print,ARGS(_Za.At(1).ToEID()))
                    if !ErrorIn(Result) {
                    PRINC("] :")
                    Result = Core.F_CALL(C_print,ARGS(o.ToEID()))
                    if !ErrorIn(Result) {
                    PRINC(" ")
                    Result = F_lbreak_integer(2)
                    if !ErrorIn(Result) {
                    Result = Core.F_CALL(C_print,ARGS(ToList(OBJ(Core.F_CALL(C_args,ARGS(a.ToEID())))).At(1).ToEID()))
                    }}}}
                    } else {
                    Result = Core.F_CALL(C_print,ARGS(_Za.At(0).ToEID()))
                    if !ErrorIn(Result) {
                    PRINC("[")
                    Result = Core.F_CALL(C_print,ARGS(_Za.At(1).ToEID()))
                    if !ErrorIn(Result) {
                    PRINC("] := ")
                    Result = F_lbreak_integer(2)
                    if !ErrorIn(Result) {
                    Result = Core.F_CALL(C_print,ARGS(a.ToEID()))
                    }}}
                    } 
                  } 
                } 
              } 
            }  else if ((_Zs.Id() == C_Language_assign.Id()) && 
              (_Za.At(0).Isa.IsIn(C_property) == CTRUE)) { 
            { var a *ClaireAny = _Za.At(2)
              { var o *ClaireAny
                if (a.Isa.IsIn(C_Call) == CTRUE) { 
                  { var g0068 *Call = To_Call(a)
                    o = g0068.Selector.Id()
                    } 
                  } else {
                  o = CFALSE.Id()
                  } 
                var g0071I *ClaireBoolean
                { var arg_2 *ClaireAny
                  if (a.Isa.IsIn(C_Call) == CTRUE) { 
                    { var g0069 *Call = To_Call(a)
                      arg_2 = g0069.Args.At(0)
                      } 
                    } else {
                    arg_2 = CFALSE.Id()
                    } 
                  g0071I = F_sugar_ask_any(_Za.At(0),_Za.At(1),o,arg_2)
                  } 
                if (g0071I == CTRUE) { 
                  Result = Core.F_CALL(C_print,ARGS(_Za.At(0).ToEID()))
                  if !ErrorIn(Result) {
                  PRINC("(")
                  Result = Core.F_CALL(C_print,ARGS(_Za.At(1).ToEID()))
                  if !ErrorIn(Result) {
                  PRINC(") :")
                  Result = Core.F_CALL(C_print,ARGS(o.ToEID()))
                  if !ErrorIn(Result) {
                  PRINC(" ")
                  Result = F_lbreak_integer(2)
                  if !ErrorIn(Result) {
                  Result = Core.F_CALL(C_print,ARGS(ToList(OBJ(Core.F_CALL(C_args,ARGS(a.ToEID())))).At(1).ToEID()))
                  }}}}
                  } else {
                  Result = Core.F_CALL(C_print,ARGS(_Za.At(0).ToEID()))
                  if !ErrorIn(Result) {
                  PRINC("(")
                  Result = Core.F_CALL(C_print,ARGS(_Za.At(1).ToEID()))
                  if !ErrorIn(Result) {
                  PRINC(") := ")
                  Result = F_lbreak_integer(2)
                  if !ErrorIn(Result) {
                  Result = Core.F_CALL(C_print,ARGS(_Za.At(2).ToEID()))
                  }}}
                  } 
                } 
              } 
            }  else if ((_Zs.Id() == C_add.Id()) && 
              (_Za.At(0).Isa.IsIn(C_property) == CTRUE)) { 
            Result = Core.F_CALL(C_print,ARGS(_Za.At(0).ToEID()))
            if !ErrorIn(Result) {
            PRINC("(")
            Result = Core.F_CALL(C_print,ARGS(_Za.At(1).ToEID()))
            if !ErrorIn(Result) {
            PRINC(") :add ")
            Result = F_lbreak_integer(2)
            if !ErrorIn(Result) {
            Result = Core.F_CALL(C_print,ARGS(_Za.At(2).ToEID()))
            }}}
            }  else if ((_Zs.Id() == C_delete.Id()) && 
              (_Za.At(0).Isa.IsIn(C_property) == CTRUE)) { 
            Result = Core.F_CALL(C_print,ARGS(_Za.At(0).ToEID()))
            if !ErrorIn(Result) {
            PRINC("(")
            Result = Core.F_CALL(C_print,ARGS(_Za.At(1).ToEID()))
            if !ErrorIn(Result) {
            PRINC(") :delete ")
            Result = F_lbreak_integer(2)
            if !ErrorIn(Result) {
            Result = Core.F_CALL(C_print,ARGS(_Za.At(2).ToEID()))
            }}}
            }  else if ((_Za.At(0) == ClEnv.Id()) && 
              (_Za.Length() == 1)) { 
            Result = Core.F_print_any(_Zs.Id())
            if !ErrorIn(Result) {
            PRINC("()")
            Result = EVOID
            }
            } else {
            Result = Core.F_print_any(_Zs.Id())
            if !ErrorIn(Result) {
            PRINC("(")
            F_set_level_void()
            Result = F_Language_printbox_list2(_Za)
            if !ErrorIn(Result) {
            PRINC(")")
            Result = EVOID
            }}
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
      } 
    return Result} 
  
// The EID go function for: self_print @ Call (throw: true) 
func E_self_print_Call_Language (self EID) EID { 
    return To_Call(OBJ(self)).SelfPrint( )} 
  
/* The go function for: self_print(self:Call+) [status=1] */
func (self *Call_plus) SelfPrint () EID { 
    var Result EID
    Result = F_printexp_any(self.Args.At(0),CTRUE)
    if !ErrorIn(Result) {
    PRINC(".")
    Result = Core.F_print_any(self.Selector.Id())
    }
    return Result} 
  
// The EID go function for: self_print @ Call+ (throw: true) 
func E_self_print_Call_plus_Language (self EID) EID { 
    return To_Call_plus(OBJ(self)).SelfPrint( )} 
  
/* The go function for: self_eval(self:Call) [status=1] */
func (self *Call) SelfEval () EID { 
    var Result EID
    { var start int = ClEnv.Index
      { var p *ClaireProperty = self.Selector
        if (ClEnv.Debug_I >= 0) { 
          C_iClaire_LastCall.Value = self.Id()
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
              { 
                var arg_2 EID
                arg_2 = EVAL(x)
                if ErrorIn(arg_2) {loop_1 = arg_2
                } else {
                loop_1 = ClEnv.Push(arg_2)}
                } 
              if ErrorIn(loop_1) {Result = loop_1
              break
              } else {
              }
              } 
            } 
          if !ErrorIn(Result) {
          { 
            var rx EID
            rx = Core.F_eval_message_property(p,Core.F_find_which_property(p,start,OWNER(ClEnv.EvalStack[start])),start,CTRUE)
            if ErrorIn(rx) {Result = rx
            } else {
            C_iClaire_LastCall.Value = self.Id()
            Result = rx}
            } 
          }
          } else {
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
              { 
                var arg_4 EID
                arg_4 = EVAL(x)
                if ErrorIn(arg_4) {loop_3 = arg_4
                } else {
                loop_3 = ClEnv.Push(arg_4)}
                } 
              if ErrorIn(loop_3) {Result = loop_3
              break
              } else {
              }
              } 
            } 
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
/* The go function for: self_eval(self:Call+) [status=1] */
func (self *Call_plus) SelfEval () EID { 
    var Result EID
    { 
      var arg_1 EID
      arg_1 = EVAL(self.Args.At(0))
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
/* The go function for: self_eval(self:Call*) [status=1] */
func (self *Call_star) SelfEval () EID { 
    var Result EID
    { var start int = ClEnv.Index
      { var p *ClaireProperty = self.Selector
        if (ClEnv.Debug_I >= 0) { 
          C_iClaire_LastCall.Value = self.Id()
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
              { 
                var arg_2 EID
                arg_2 = EVAL(x)
                if ErrorIn(arg_2) {loop_1 = arg_2
                } else {
                loop_1 = ClEnv.Push(arg_2)}
                } 
              if ErrorIn(loop_1) {Result = loop_1
              break
              } else {
              }
              } 
            } 
          if !ErrorIn(Result) {
          { 
            var rx EID
            rx = Core.F_eval_message_property(p,Core.F_find_which_property(p,start,OWNER(ClEnv.EvalStack[start])),start,CTRUE)
            if ErrorIn(rx) {Result = rx
            } else {
            C_iClaire_LastCall.Value = self.Id()
            Result = rx}
            } 
          }
          } else {
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
              { 
                var arg_4 EID
                arg_4 = EVAL(x)
                if ErrorIn(arg_4) {loop_3 = arg_4
                } else {
                loop_3 = ClEnv.Push(arg_4)}
                } 
              if ErrorIn(loop_3) {Result = loop_3
              break
              } else {
              }
              } 
            } 
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
/* The go function for: printe(self:any,s:property) [status=1] */
func F_printe_any (self *ClaireAny,s *ClaireProperty) EID { 
    var Result EID
    var g0073I *ClaireBoolean
    if (self.Isa.IsIn(C_Call) == CTRUE) { 
      { var g0072 *Call = To_Call(self)
        g0073I = MakeBoolean((g0072.Selector.Isa.IsIn(C_operation) == CTRUE) && (g0072.Args.Length() == 2))
        } 
      } else {
      g0073I = CFALSE
      } 
    if (g0073I == CTRUE) { 
      PRINC("(")
      Result = Core.F_CALL(C_print,ARGS(self.ToEID()))
      if !ErrorIn(Result) {
      PRINC(")")
      Result = EVOID
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
/* The go function for: sugar?(x:any,x2:any,o:any,a:any) [status=0] */
func F_sugar_ask_any (x *ClaireAny,x2 *ClaireAny,o *ClaireAny,a *ClaireAny) *ClaireBoolean { 
    var Result *ClaireBoolean
    if (o.Isa.IsIn(C_operation) == CTRUE) { 
      if (x.Isa.IsIn(C_property) == CTRUE) { 
        { var g0075 *ClaireProperty = ToProperty(x)
          if (a.Isa.IsIn(C_Call) == CTRUE) { 
            { var g0076 *Call = To_Call(a)
              Result = MakeBoolean((g0075.Id() == g0076.Selector.Id()) && (Equal(g0076.Args.At(0),x2) == CTRUE))
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
/* The go function for: self_print(self:Assign) [status=1] */
func (self *Assign) SelfPrint () EID { 
    var Result EID
    { var a *ClaireAny = self.Arg
      { var o *ClaireAny
        if (a.Isa.IsIn(C_Call) == CTRUE) { 
          { var g0080 *Call = To_Call(a)
            o = g0080.Selector.Id()
            } 
          } else {
          o = CFALSE.Id()
          } 
        var g0082I *ClaireBoolean
        { var arg_1 *ClaireAny
          if (a.Isa.IsIn(C_Call) == CTRUE) { 
            { var g0081 *Call = To_Call(a)
              arg_1 = g0081.Args.At(0)
              } 
            } else {
            arg_1 = CFALSE.Id()
            } 
          g0082I = F_sugar_ask_any(self.ClaireVar,CEMPTY.Id(),o,arg_1)
          } 
        if (g0082I == CTRUE) { 
          Result = Core.F_CALL(C_print,ARGS(self.ClaireVar.ToEID()))
          if !ErrorIn(Result) {
          PRINC(" :")
          Result = Core.F_CALL(C_print,ARGS(o.ToEID()))
          if !ErrorIn(Result) {
          PRINC(" ")
          Result = F_lbreak_integer(2)
          if !ErrorIn(Result) {
          Result = F_printexp_any(ToList(OBJ(Core.F_CALL(C_args,ARGS(a.ToEID())))).At(1),CTRUE)
          }}}
          } else {
          Result = Core.F_CALL(C_print,ARGS(self.ClaireVar.ToEID()))
          if !ErrorIn(Result) {
          PRINC(" := ")
          Result = F_lbreak_integer(2)
          if !ErrorIn(Result) {
          Result = F_printexp_any(a,CTRUE)
          }}
          } 
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
        } 
      } 
    return Result} 
  
// The EID go function for: self_print @ Assign (throw: true) 
func E_self_print_Assign_Language (self EID) EID { 
    return To_Assign(OBJ(self)).SelfPrint( )} 
  
// assumes that Assign is well-formed (self.var must be a variable)
/* The go function for: self_eval(self:Assign) [status=1] */
func (self *Assign) SelfEval () EID { 
    var Result EID
    { 
      var arg_1 EID
      arg_1 = EVAL(self.Arg)
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
/* The go function for: self_print(self:Gassign) [status=1] */
func (self *Gassign) SelfPrint () EID { 
    var Result EID
    { var a *ClaireAny = self.Arg
      { var o *ClaireAny
        if (a.Isa.IsIn(C_Call) == CTRUE) { 
          { var g0084 *Call = To_Call(a)
            o = g0084.Selector.Id()
            } 
          } else {
          o = CFALSE.Id()
          } 
        var g0086I *ClaireBoolean
        { var arg_1 *ClaireAny
          if (a.Isa.IsIn(C_Call) == CTRUE) { 
            { var g0085 *Call = To_Call(a)
              arg_1 = g0085.Args.At(0)
              } 
            } else {
            arg_1 = CFALSE.Id()
            } 
          g0086I = F_sugar_ask_any(self.ClaireVar.Id(),CEMPTY.Id(),o,arg_1)
          } 
        if (g0086I == CTRUE) { 
          Result = Core.F_print_any(self.ClaireVar.Id())
          if !ErrorIn(Result) {
          PRINC(" :")
          Result = Core.F_CALL(C_print,ARGS(o.ToEID()))
          if !ErrorIn(Result) {
          PRINC(" ")
          Result = F_lbreak_integer(2)
          if !ErrorIn(Result) {
          Result = Core.F_CALL(C_print,ARGS(ToList(OBJ(Core.F_CALL(C_args,ARGS(a.ToEID())))).At(1).ToEID()))
          }}}
          } else {
          Result = Core.F_print_any(self.ClaireVar.Id())
          if !ErrorIn(Result) {
          PRINC(" := ")
          Result = F_lbreak_integer(2)
          if !ErrorIn(Result) {
          Result = Core.F_CALL(C_print,ARGS(a.ToEID()))
          }}
          } 
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
        } 
      } 
    return Result} 
  
// The EID go function for: self_print @ Gassign (throw: true) 
func E_self_print_Gassign_Language (self EID) EID { 
    return To_Gassign(OBJ(self)).SelfPrint( )} 
  
/* The go function for: self_eval(self:Gassign) [status=1] */
func (self *Gassign) SelfEval () EID { 
    var Result EID
    { var v *Core.GlobalVariable = self.ClaireVar
      { var arg_1 *ClaireAny
        var try_2 EID
        try_2 = EVAL(self.Arg)
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
/* The go function for: self_print(self:And) [status=1] */
func (self *And) SelfPrint () EID { 
    var Result EID
    PRINC("(")
    Result = F_Language_printbox_list3(self.Args,MakeString(" & "))
    if !ErrorIn(Result) {
    PRINC(")")
    Result = EVOID
    }
    return Result} 
  
// The EID go function for: self_print @ And (throw: true) 
func E_self_print_And_Language (self EID) EID { 
    return To_And(OBJ(self)).SelfPrint( )} 
  
/* The go function for: self_eval(self:And) [status=1] */
func (self *And) SelfEval () EID { 
    var Result EID
    { var arg_1 *ClaireAny
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
          var g0087I *ClaireBoolean
          var try_4 EID
          { var arg_5 *ClaireBoolean
            var try_6 EID
            { var arg_7 *ClaireAny
              var try_8 EID
              try_8 = EVAL(x)
              if ErrorIn(try_8) {try_6 = try_8
              } else {
              arg_7 = ANY(try_8)
              try_6 = EID{F_boolean_I_any(arg_7).Id(),0}
              }
              } 
            if ErrorIn(try_6) {try_4 = try_6
            } else {
            arg_5 = ToBoolean(OBJ(try_6))
            try_4 = EID{Core.F__I_equal_any(arg_5.Id(),CTRUE.Id()).Id(),0}
            }
            } 
          if ErrorIn(try_4) {loop_3 = try_4
          } else {
          g0087I = ToBoolean(OBJ(try_4))
          if (g0087I == CTRUE) { 
            try_2 = EID{CTRUE.Id(),0}
            break
            } else {
            loop_3 = EID{CFALSE.Id(),0}
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
/* The go function for: self_print(self:Or) [status=1] */
func (self *Or) SelfPrint () EID { 
    var Result EID
    PRINC("(")
    Result = F_Language_printbox_list3(self.Args,MakeString(" | "))
    if !ErrorIn(Result) {
    PRINC(")")
    Result = EVOID
    }
    return Result} 
  
// The EID go function for: self_print @ Or (throw: true) 
func E_self_print_Or_Language (self EID) EID { 
    return To_Or(OBJ(self)).SelfPrint( )} 
  
/* The go function for: self_eval(self:Or) [status=1] */
func (self *Or) SelfEval () EID { 
    var Result EID
    var g0088I *ClaireBoolean
    var try_1 EID
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
        var g0089I *ClaireBoolean
        var try_3 EID
        { var arg_4 *ClaireAny
          var try_5 EID
          try_5 = EVAL(x)
          if ErrorIn(try_5) {try_3 = try_5
          } else {
          arg_4 = ANY(try_5)
          try_3 = EID{F_boolean_I_any(arg_4).Id(),0}
          }
          } 
        if ErrorIn(try_3) {loop_2 = try_3
        } else {
        g0089I = ToBoolean(OBJ(try_3))
        if (g0089I == CTRUE) { 
          try_1 = EID{CTRUE.Id(),0}
          break
          } else {
          loop_2 = EID{CFALSE.Id(),0}
          } 
        }
        if ErrorIn(loop_2) {try_1 = loop_2
        break
        } else {
        }
        } 
      } 
    if ErrorIn(try_1) {Result = try_1
    } else {
    g0088I = ToBoolean(OBJ(try_1))
    if (g0088I == CTRUE) { 
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
/* The go function for: self_print(self:Quote) [status=1] */
func (self *Quote) SelfPrint () EID { 
    var Result EID
    PRINC("quote(")
    Result = Core.F_CALL(C_print,ARGS(self.Arg.ToEID()))
    if !ErrorIn(Result) {
    PRINC(")")
    Result = EVOID
    }
    return Result} 
  
// The EID go function for: self_print @ Quote (throw: true) 
func E_self_print_Quote_Language (self EID) EID { 
    return To_Quote(OBJ(self)).SelfPrint( )} 
  
/* The go function for: self_eval(self:Quote) [status=0] */
func (self *Quote) SelfEval () EID { 
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
/* The go function for: self_print(self:Call_method) [status=1] */
func (self *CallMethod) SelfPrint () EID { 
    var Result EID
    Result = Core.F_print_any(self.Arg.Id())
    if !ErrorIn(Result) {
    PRINC("(")
    Result = Core.F_princ_list(self.Args)
    if !ErrorIn(Result) {
    PRINC(")")
    Result = EVOID
    }}
    return Result} 
  
// The EID go function for: self_print @ Call_method (throw: true) 
func E_self_print_Call_method_Language (self EID) EID { 
    return To_CallMethod(OBJ(self)).SelfPrint( )} 
  
/* The go function for: self_eval(self:Call_method) [status=1] */
func (self *CallMethod) SelfEval () EID { 
    var Result EID
    { var start int = ClEnv.Index
      { var Cprop *ClaireMethod = self.Arg
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
            { 
              var arg_2 EID
              arg_2 = EVAL(x)
              if ErrorIn(arg_2) {loop_1 = arg_2
              } else {
              loop_1 = ClEnv.Push(arg_2)}
              } 
            if ErrorIn(loop_1) {Result = loop_1
            break
            } else {
            }
            } 
          } 
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
/* The go function for: self_eval(self:Call_method1) [status=1] */
func (self *CallMethod1) SelfEval () EID { 
    var Result EID
    { var f *ClaireMethod = self.Arg
      { var l *ClaireList = self.Args
        { 
          var arg_1 EID
          arg_1 = EVAL(l.At(0))
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
/* The go function for: self_eval(self:Call_method2) [status=1] */
func (self *CallMethod2) SelfEval () EID { 
    var Result EID
    { var f *ClaireMethod = self.Arg
      { var l *ClaireList = self.Args
        { 
          var arg_1 EID
          arg_1 = EVAL(l.At(0))
          if ErrorIn(arg_1) {Result = arg_1
          } else {
          var arg_2 EID
          arg_2 = EVAL(l.At(1))
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
/* The go function for: self_eval(self:Call_method3) [status=1] */
func (self *Language_CallMethod3) SelfEval () EID { 
    var Result EID
    { var f *ClaireMethod = self.Arg
      { var l *ClaireList = self.Args
        { 
          var arg_1 EID
          arg_1 = EVAL(l.At(0))
          if ErrorIn(arg_1) {Result = arg_1
          } else {
          var arg_2 EID
          arg_2 = EVAL(l.At(1))
          if ErrorIn(arg_2) {Result = arg_2
          } else {
          var arg_3 EID
          arg_3 = EVAL(l.At(2))
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
/* The go function for: self_print(self:Call_slot) [status=1] */
func (self *CallSlot) SelfPrint () EID { 
    var Result EID
    Result = Core.F_print_any(self.Selector.Id())
    if !ErrorIn(Result) {
    PRINC("(")
    Result = Core.F_CALL(C_print,ARGS(self.Arg.ToEID()))
    if !ErrorIn(Result) {
    PRINC(")")
    Result = EVOID
    }}
    return Result} 
  
// The EID go function for: self_print @ Call_slot (throw: true) 
func E_self_print_Call_slot_Language (self EID) EID { 
    return To_CallSlot(OBJ(self)).SelfPrint( )} 
  
/* The go function for: self_eval(self:Call_slot) [status=1] */
func (self *CallSlot) SelfEval () EID { 
    var Result EID
    { var arg_1 *ClaireAny
      var try_2 EID
      try_2 = EVAL(self.Arg)
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
/* The go function for: self_print(self:Call_array) [status=1] */
func (self *CallArray) SelfPrint () EID { 
    var Result EID
    Result = Core.F_CALL(C_print,ARGS(self.Selector.ToEID()))
    if !ErrorIn(Result) {
    PRINC("[")
    Result = Core.F_CALL(C_print,ARGS(self.Arg.ToEID()))
    if !ErrorIn(Result) {
    PRINC("]")
    Result = EVOID
    }}
    return Result} 
  
// The EID go function for: self_print @ Call_array (throw: true) 
func E_self_print_Call_array_Language (self EID) EID { 
    return To_CallArray(OBJ(self)).SelfPrint( )} 
  
/* The go function for: self_eval(self:Call_array) [status=1] */
func (self *CallArray) SelfEval () EID { 
    var Result EID
    { var arg_1 *ClaireAny
      var try_3 EID
      try_3 = EVAL(self.Selector)
      if ErrorIn(try_3) {Result = try_3
      } else {
      arg_1 = ANY(try_3)
      { var arg_2 *ClaireAny
        var try_4 EID
        try_4 = EVAL(self.Arg)
        if ErrorIn(try_4) {Result = try_4
        } else {
        arg_2 = ANY(try_4)
        Result = ToArray(arg_1).Nth(ToInteger(arg_2).Value)
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
/* The go function for: self_print(self:Call_table) [status=1] */
func (self *CallTable) SelfPrint () EID { 
    var Result EID
    Result = Core.F_print_any(self.Selector.Id())
    if !ErrorIn(Result) {
    PRINC("[")
    Result = Core.F_CALL(C_print,ARGS(self.Arg.ToEID()))
    if !ErrorIn(Result) {
    PRINC("]")
    Result = EVOID
    }}
    return Result} 
  
// The EID go function for: self_print @ Call_table (throw: true) 
func E_self_print_Call_table_Language (self EID) EID { 
    return To_CallTable(OBJ(self)).SelfPrint( )} 
  
/* The go function for: self_eval(self:Call_table) [status=1] */
func (self *CallTable) SelfEval () EID { 
    var Result EID
    if (self.Test == CTRUE) { 
      { var arg_1 *ClaireAny
        var try_2 EID
        try_2 = EVAL(self.Arg)
        if ErrorIn(try_2) {Result = try_2
        } else {
        arg_1 = ANY(try_2)
        Result = Core.F_nth_table1(self.Selector,arg_1)
        }
        } 
      } else {
      { var arg_3 *ClaireAny
        var try_4 EID
        try_4 = EVAL(self.Arg)
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
/* The go function for: self_print(self:Update) [status=1] */
func (self *Update) SelfPrint () EID { 
    var Result EID
    Result = Core.F_CALL(C_print,ARGS(self.Selector.ToEID()))
    if !ErrorIn(Result) {
    PRINC("(")
    Result = Core.F_CALL(C_print,ARGS(Core.F_CALL(C_arg,ARGS(self.ClaireVar.ToEID()))))
    if !ErrorIn(Result) {
    PRINC(") := ")
    Result = Core.F_CALL(C_print,ARGS(self.Value.ToEID()))
    }}
    return Result} 
  
// The EID go function for: self_print @ Update (throw: true) 
func E_self_print_Update_Language (self EID) EID { 
    return To_Update(OBJ(self)).SelfPrint( )} 
  
/* The go function for: self_eval(self:Update) [status=1] */
func (self *Update) SelfEval () EID { 
    var Result EID
    { var s *ClaireAny = self.Selector
      if (s.Isa.IsIn(C_property) == CTRUE) { 
        { var g0090 *ClaireProperty = ToProperty(s)
          { var arg_1 *ClaireAny
            var try_3 EID
            try_3 = EVAL(ANY(Core.F_CALL(C_arg,ARGS(self.ClaireVar.ToEID()))))
            if ErrorIn(try_3) {Result = try_3
            } else {
            arg_1 = ANY(try_3)
            { var arg_2 *ClaireAny
              var try_4 EID
              try_4 = EVAL(self.Value)
              if ErrorIn(try_4) {Result = try_4
              } else {
              arg_2 = ANY(try_4)
              Result = Core.F_put_property2(g0090,ToObject(arg_1),arg_2)
              }
              } 
            }
            } 
          } 
        }  else if (C_table.Id() == s.Isa.Id()) { 
        { var g0091 *ClaireTable = ToTable(s)
          { var arg_5 *ClaireAny
            var try_7 EID
            try_7 = EVAL(ANY(Core.F_CALL(C_arg,ARGS(self.ClaireVar.ToEID()))))
            if ErrorIn(try_7) {Result = try_7
            } else {
            arg_5 = ANY(try_7)
            { var arg_6 *ClaireAny
              var try_8 EID
              try_8 = EVAL(self.Value)
              if ErrorIn(try_8) {Result = try_8
              } else {
              arg_6 = ANY(try_8)
              Result = Core.F_nth_equal_table1(g0091,arg_5,arg_6)
              }
              } 
            }
            } 
          } 
        } else {
        Result = EID{CFALSE.Id(),0}
        } 
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
/* The go function for: self_print(self:Super) [status=1] */
func (self *Super) SelfPrint () EID { 
    var Result EID
    { var _Zl int = Core.C_pretty.Index
      { var _Zs *ClaireProperty = self.Selector
        { var _Za *ClaireList = self.Args
          Result = Core.F_print_any(_Zs.Id())
          if !ErrorIn(Result) {
          PRINC("@")
          Result = Core.F_print_any(self.CastTo.Id())
          if !ErrorIn(Result) {
          PRINC("(")
          F_set_level_void()
          Result = F_Language_printbox_list2(_Za)
          if !ErrorIn(Result) {
          PRINC(")")
          Result = EVOID
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
      } 
    return Result} 
  
// The EID go function for: self_print @ Super (throw: true) 
func E_self_print_Super_Language (self EID) EID { 
    return To_Super(OBJ(self)).SelfPrint( )} 
  
/* The go function for: self_eval(self:Super) [status=1] */
func (self *Super) SelfEval () EID { 
    var Result EID
    { var start int = ClEnv.Index
      { var t *ClaireType = self.CastTo
        { var c *ClaireClass = t.Class_I()
          { var p *ClaireProperty = self.Selector
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
                { 
                  var arg_2 EID
                  arg_2 = EVAL(x)
                  if ErrorIn(arg_2) {loop_1 = arg_2
                  } else {
                  loop_1 = ClEnv.Push(arg_2)}
                  } 
                if ErrorIn(loop_1) {Result = loop_1
                break
                } else {
                }
                } 
              } 
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
/* The go function for: self_print(x:Cast) [status=1] */
func (x *Cast) SelfPrint () EID { 
    var Result EID
    Result = F_printexp_any(x.Arg,CFALSE)
    if !ErrorIn(Result) {
    PRINC(" as ")
    Result = F_printexp_any(x.SetArg.Id(),CFALSE)
    }
    return Result} 
  
// The EID go function for: self_print @ Cast (throw: true) 
func E_self_print_Cast_Language (x EID) EID { 
    return To_Cast(OBJ(x)).SelfPrint( )} 
  
/* The go function for: self_eval(self:Cast) [status=1] */
func (self *Cast) SelfEval () EID { 
    var Result EID
    { var x *ClaireAny
      var try_1 EID
      try_1 = EVAL(self.Arg)
      if ErrorIn(try_1) {Result = try_1
      } else {
      x = ANY(try_1)
      { var y *ClaireType = self.SetArg
        var g0093I *ClaireBoolean
        if (y.Isa.IsIn(C_Param) == CTRUE) { 
          { var g0092 *ClaireParam = To_Param(y.Id())
            g0093I = MakeBoolean(((g0092.Arg.Id() == C_list.Id()) || 
                (g0092.Arg.Id() == C_set.Id())) && (C_set.Id() == g0092.Args.At(0).Isa.Id()) && (Core.F__Z_any1(x,g0092.Arg) == CTRUE))
            } 
          } else {
          g0093I = CFALSE
          } 
        if (g0093I == CTRUE) { 
          { var arg_2 *ClaireAny
            var try_3 EID
            try_3 = Core.F_the_type(ToType(To_Param(y.Id()).Args.At(0)))
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
/* The go function for: self_print(self:Return) [status=1] */
func (self *Return) SelfPrint () EID { 
    var Result EID
    PRINC("break(")
    Core.C_pretty.Index = (Core.C_pretty.Index+2)
    Result = Core.F_CALL(C_print,ARGS(self.Arg.ToEID()))
    if !ErrorIn(Result) {
    Core.C_pretty.Index = (Core.C_pretty.Index-2)
    PRINC(")")
    Result = EVOID
    }
    return Result} 
  
// The EID go function for: self_print @ Return (throw: true) 
func E_self_print_Return_Language (self EID) EID { 
    return To_Return(OBJ(self)).SelfPrint( )} 
  
/* The go function for: self_eval(self:Return) [status=1] */
func (self *Return) SelfEval () EID { 
    var Result EID
    { var arg_1 *ClaireAny
      var try_2 EID
      { var arg_3 *ClaireAny
        var try_4 EID
        try_4 = EVAL(self.Arg)
        if ErrorIn(try_4) {try_2 = try_4
        } else {
        arg_3 = ANY(try_4)
        try_2 = Core.C_return_error.Make(arg_3).ToEID()
        }
        } 
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
/* The go function for: substitution(self:any,x:Variable,val:any) [status=0] */
func F_substitution_any (self *ClaireAny,x *ClaireVariable,val *ClaireAny) *ClaireAny { 
    var Result *ClaireAny
    if (self.Isa.IsIn(C_Variable) == CTRUE) { 
      { var g0096 *ClaireVariable = To_Variable(self)
        Result = IfThenElse((g0096.Pname.Id() == x.Pname.Id()),
          val,
          g0096.Id())
        } 
      }  else if (self.Isa.IsIn(C_list) == CTRUE) { 
      { var g0097 *ClaireList = ToList(self)
        { var i int = 1
          { var g0098 int = g0097.Length()
            for (i <= g0098) { 
              if ((g0097.At(i-1).Isa.IsIn(C_Variable) == CTRUE) || 
                  (g0097.At(i-1).Isa.IsIn(C_unbound_symbol) == CTRUE)) { 
                ToArray(g0097.Id()).NthPut(i,F_substitution_any(g0097.At(i-1),x,val))
                } else {
                F_substitution_any(g0097.At(i-1),x,val)
                } 
              i = (i+1)
              } 
            } 
          } 
        Result = g0097.Id()
        } 
      }  else if (self.Isa.IsIn(C_unbound_symbol) == CTRUE) { 
      { var g0099 *ClaireUnboundSymbol = ToUnboundSymbol(self)
        Result = IfThenElse((g0099.Name.Id() == x.Pname.Id()),
          val,
          g0099.Id())
        } 
      }  else if (self.Isa.IsIn(C_Instruction) == CTRUE) { 
      { var g0100 *ClaireInstruction = To_Instruction(self)
        { 
          var s *ClaireSlot
          _ = s
          var s_iter *ClaireAny
          var s_support *ClaireList
          s_support = g0100.Id().Isa.Slots
          for _,s_iter = range(s_support.ValuesO()){ 
            s = ToSlot(s_iter)
            { var y *ClaireAny = Core.F_get_slot(s,ToObject(g0100.Id()))
              if ((y.Isa.IsIn(C_Variable) == CTRUE) || 
                  (y.Isa.IsIn(C_unbound_symbol) == CTRUE)) { 
                Core.F_put_slot(s,ToObject(g0100.Id()),F_substitution_any(y,x,val))
                } else {
                F_substitution_any(y,x,val)
                } 
              } 
            } 
          } 
        Result = g0100.Id()
        } 
      } else {
      Result = self
      } 
    return Result} 
  
// The EID go function for: substitution @ any (throw: false) 
func E_substitution_any (self EID,x EID,val EID) EID { 
    return F_substitution_any(ANY(self),To_Variable(OBJ(x)),ANY(val) ).ToEID()} 
  
// count the number of occurrences of x
/* The go function for: occurrence(self:any,x:Variable) [status=0] */
func F_occurrence_any (self *ClaireAny,x *ClaireVariable) int { 
    var Result int
    if (self.Isa.IsIn(C_Variable) == CTRUE) { 
      { var g0102 *ClaireVariable = To_Variable(self)
        if (g0102.Pname.Id() == x.Pname.Id()) { 
          Result = 1
          } else {
          Result = 0
          } 
        } 
      }  else if (self.Isa.IsIn(C_list) == CTRUE) { 
      { var g0103 *ClaireList = ToList(self)
        { var n int = 0
          { var i int = 1
            { var g0104 int = g0103.Length()
              for (i <= g0104) { 
                n = (n+F_occurrence_any(g0103.At(i-1),x))
                i = (i+1)
                } 
              } 
            } 
          Result = n
          } 
        } 
      }  else if (self.Isa.IsIn(C_unbound_symbol) == CTRUE) { 
      { var g0105 *ClaireUnboundSymbol = ToUnboundSymbol(self)
        if (g0105.Name.Id() == x.Pname.Id()) { 
          Result = 1
          } else {
          Result = 0
          } 
        } 
      }  else if (self.Isa.IsIn(C_Instruction) == CTRUE) { 
      { var g0106 *ClaireInstruction = To_Instruction(self)
        { var n int = 0
          { 
            var s *ClaireSlot
            _ = s
            var s_iter *ClaireAny
            var s_support *ClaireList
            s_support = g0106.Id().Isa.Slots
            for _,s_iter = range(s_support.ValuesO()){ 
              s = ToSlot(s_iter)
              n = (n+F_occurrence_any(Core.F_get_slot(s,ToObject(g0106.Id())),x))
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
/* The go function for: occurchange(self:any,x:Variable) [status=0] */
func F_occurchange_any (self *ClaireAny,x *ClaireVariable) *ClaireBoolean { 
    var Result *ClaireBoolean
    if (self.Isa.IsIn(C_Assign) == CTRUE) { 
      { var g0108 *Assign = To_Assign(self)
        Result = Equal(ANY(Core.F_CALL(C_mClaire_pname,ARGS(g0108.ClaireVar.ToEID()))),x.Pname.Id())
        } 
      }  else if (self.Isa.IsIn(C_list) == CTRUE) { 
      { var g0109 *ClaireList = ToList(self)
        { var arg_1 *ClaireAny
          { 
            var y *ClaireAny
            _ = y
            arg_1= CFALSE.Id()
            var y_support *ClaireList
            y_support = g0109
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
      { var g0110 *ClaireInstruction = To_Instruction(self)
        { var arg_2 *ClaireAny
          { 
            var s *ClaireSlot
            _ = s
            var s_iter *ClaireAny
            arg_2= CFALSE.Id()
            var s_support *ClaireList
            s_support = g0110.Id().Isa.Slots
            for _,s_iter = range(s_support.ValuesO()){ 
              s = ToSlot(s_iter)
              if (F_occurchange_any(Core.F_get_slot(s,ToObject(g0110.Id())),x) == CTRUE) { 
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
/* The go function for: occurexact(self:any,x:Variable) [status=0] */
func F_Language_occurexact_any (self *ClaireAny,x *ClaireVariable) int { 
    var Result int
    if (self.Isa.IsIn(C_Variable) == CTRUE) { 
      { var g0112 *ClaireVariable = To_Variable(self)
        if ((g0112.Pname.Id() == x.Pname.Id()) && 
            (g0112.Index == x.Index)) { 
          Result = 1
          } else {
          Result = 0
          } 
        } 
      }  else if (self.Isa.IsIn(C_list) == CTRUE) { 
      { var g0113 *ClaireList = ToList(self)
        { var n int = 0
          { var i int = 1
            { var g0114 int = g0113.Length()
              for (i <= g0114) { 
                n = (n+F_Language_occurexact_any(g0113.At(i-1),x))
                i = (i+1)
                } 
              } 
            } 
          Result = n
          } 
        } 
      }  else if (self.Isa.IsIn(C_unbound_symbol) == CTRUE) { 
      Result = 0
      }  else if (self.Isa.IsIn(C_Assign) == CTRUE) { 
      { var g0116 *Assign = To_Assign(self)
        Result = F_Language_occurexact_any(g0116.Arg,x)
        } 
      }  else if (self.Isa.IsIn(C_Instruction) == CTRUE) { 
      { var g0117 *ClaireInstruction = To_Instruction(self)
        { var n int = 0
          { 
            var s *ClaireSlot
            _ = s
            var s_iter *ClaireAny
            var s_support *ClaireList
            s_support = g0117.Id().Isa.Slots
            for _,s_iter = range(s_support.ValuesO()){ 
              s = ToSlot(s_iter)
              n = (n+F_Language_occurexact_any(Core.F_get_slot(s,ToObject(g0117.Id())),x))
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
/* The go function for: instruction_copy(self:any) [status=0] */
func F_instruction_copy_any (self *ClaireAny) *ClaireAny { 
    var Result *ClaireAny
    if (self.Isa.IsIn(C_list) == CTRUE) { 
      { var g0119 *ClaireList = ToList(self)
        { var l *ClaireList = g0119.Copy()
          { var i int = 1
            { var g0120 int = g0119.Length()
              for (i <= g0120) { 
                ToArray(l.Id()).NthPut(i,F_instruction_copy_any(g0119.At(i-1)))
                i = (i+1)
                } 
              } 
            } 
          Result = l.Id()
          } 
        } 
      }  else if (self.Isa.IsIn(C_Variable) == CTRUE) { 
      { var g0121 *ClaireVariable = To_Variable(self)
        Result = g0121.Id()
        } 
      }  else if (self.Isa.IsIn(C_Instruction) == CTRUE) { 
      { var g0122 *ClaireInstruction = To_Instruction(self)
        { var o *ClaireInstruction = To_Instruction(g0122.Copy().Id())
          { 
            var s *ClaireSlot
            _ = s
            var s_iter *ClaireAny
            var s_support *ClaireList
            s_support = g0122.Id().Isa.Slots
            for _,s_iter = range(s_support.ValuesO()){ 
              s = ToSlot(s_iter)
              Core.F_put_slot(s,ToObject(o.Id()),F_instruction_copy_any(Core.F_get_slot(s,ToObject(g0122.Id()))))
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
  