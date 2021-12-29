/***** CLAIRE Compilation of file /Users/ycaseau/Dropbox/src/clairev4.03/src/meta/syntax.cl 
         [version 4.0.03 / safety 5] Wednesday 12-29-2021 08:34:14 *****/

package Reader
import (_ "fmt"
	. "Kernel"
	"Core"
	"Language"
)

//-------- dumb function to prevent import errors --------
func import_g0039() { 
    _ = Core.It
    _ = Language.It
    } 
  
  
//+-------------------------------------------------------------+
//| CLAIRE                                                      |
//| syntax.cl                                                   |
//| Copyright (C) 1994 - 2021 Yves Caseau. All Rights Reserved  |
//| cf. copyright info in file object.cl: about()               |
//+-------------------------------------------------------------+
// --------------------------------------------------------------
// this file contains specialized reading methods
// --------------------------------------------------------------
// **********************************************************************
// *  Content:                                                          *
// *   Part 1: read operation expressions (<exp> <op> <exp>)            *
// *   Part 2: read control structures                                  *
// *   Part 3: read functional calls                                    *
// *   Part 4: read definitions                                         *
// **********************************************************************
// **********************************************************************
// *   Part 1: read operation expressions (<exp> <op> <exp>)            *
// **********************************************************************
// who is an operation?
//
/* {1} The go function for: operation?(y:any) [status=0] */
func F_operation_ask_any (y *ClaireAny ) *ClaireBoolean  { 
    if ((y == C_as.Id()) || 
        ((y == C_L__equal.Id()) || 
          ((y.Isa.IsIn(C_operation) == CTRUE) || 
            ((Equal(y,C_OR.Value) == CTRUE) || 
              ((y == C__Z.Id()) || 
                (y == C_add.Id())))))) {return CTRUE
    } else {return CFALSE}} 
  
// The EID go function for: operation? @ any (throw: false) 
func E_operation_ask_any (y EID) EID { 
    return EID{F_operation_ask_any(ANY(y) ).Id(),0}} 
  
// produce an expression from an operation
// apply precedence rules ((x1 op x2) y  z) -> x1 op (x2 y z)
/* {1} The go function for: combine(x:any,y:any,z:any) [status=1] */
func F_combine_any (x *ClaireAny ,y *ClaireAny ,z *ClaireAny ) EID { 
    // eid body s = any
    var Result EID 
    { var p *ClaireAny   = F_operation_I_any(x)
      if ((F_boolean_I_any(p) == CTRUE) && 
          (F_precedence_I_any(y) < F_precedence_I_any(p))) { 
        { var arg_1 *ClaireAny  
          var try_3 EID 
          try_3 = F_operand_I_any(x,1)
          if ErrorIn(try_3) {Result = try_3
          } else {
          arg_1 = ANY(try_3)
          { var arg_2 *ClaireAny  
            var try_4 EID 
            { var arg_5 *ClaireAny  
              var try_6 EID 
              try_6 = F_operand_I_any(x,2)
              if ErrorIn(try_6) {try_4 = try_6
              } else {
              arg_5 = ANY(try_6)
              try_4 = F_combine_any(arg_5,y,z)
              }
              } 
            if ErrorIn(try_4) {Result = try_4
            } else {
            arg_2 = ANY(try_4)
            Result = F_combine_I_any(arg_1,p,arg_2)
            }
            } 
          }
          } 
        } else {
        Result = F_combine_I_any(x,y,z)
        } 
      } 
    return Result} 
  
// The EID go function for: combine @ any (throw: true) 
func E_combine_any (x EID,y EID,z EID) EID { 
    return F_combine_any(ANY(x),ANY(y),ANY(z) )} 
  
// produces x op=y z
// replace r(x) :add y with add(r,x,y) for multivalued or defeasible .. also with delete
/* {1} The go function for: combine!(x:any,y:any,z:any) [status=1] */
func F_combine_I_any (x *ClaireAny ,y *ClaireAny ,z *ClaireAny ) EID { 
    // eid body s = any
    var Result EID 
    if (y == C_as.Id()) { 
      { var _CL_obj *Language.Cast   = Language.To_Cast(new(Language.Cast).Is(Language.C_Cast))
        _CL_obj.Arg = x
        { 
          var va_arg1 *Language.Cast  
          var va_arg2 *ClaireType  
          va_arg1 = _CL_obj
          var try_1 EID 
          try_1 = Language.F_extract_type_any(z)
          if ErrorIn(try_1) {Result = try_1
          } else {
          va_arg2 = ToType(OBJ(try_1))
          va_arg1.SetArg = va_arg2
          Result = EID{va_arg2.Id(),0}
          }
          } 
        if !ErrorIn(Result) {
        Result = EID{_CL_obj.Id(),0}
        }
        } 
      }  else if (y == C_L__equal.Id()) { 
      if (x.Isa.IsIn(Core.C_global_variable) == CTRUE) { 
        { var g0040 *Core.GlobalVariable   = Core.ToGlobalVariable(x)
          { var _CL_obj *Language.Gassign   = Language.To_Gassign(new(Language.Gassign).Is(Language.C_Gassign))
            _CL_obj.ClaireVar = g0040
            _CL_obj.Arg = z
            Result = EID{_CL_obj.Id(),0}
            } 
          } 
        }  else if (x.Isa.IsIn(Language.C_Call) == CTRUE) { 
        { var g0041 *Language.Call   = Language.To_Call(x)
          var g0053I *ClaireBoolean  
          if (z.Isa.IsIn(Language.C_Call) == CTRUE) { 
            { var g0042 *Language.Call   = Language.To_Call(z)
              { 
                var v_and7 *ClaireBoolean  
                
                v_and7 = MakeBoolean((g0042.Selector.Id() == C_add.Id()) || (g0042.Selector.Id() == C_delete.Id()))
                if (v_and7 == CFALSE) {g0053I = CFALSE
                } else { 
                  v_and7 = Equal(g0042.Args.At(1-1),g0041.Id())
                  if (v_and7 == CFALSE) {g0053I = CFALSE
                  } else { 
                    if (g0041.Args.Length() == 1) { 
                      v_and7 = g0041.Selector.Multivalued_ask
                      }  else if ((g0041.Selector.Id() == C_nth.Id()) && 
                        (g0041.Args.Length() == 2)) { 
                      { var p *ClaireAny   = g0041.Args.At(1-1)
                        if (p.Isa.IsIn(C_relation) == CTRUE) { 
                          { var g0043 *ClaireRelation   = ToRelation(p)
                            v_and7 = g0043.Multivalued_ask
                            } 
                          } else {
                          v_and7 = CFALSE
                          } 
                        } 
                      } else {
                      v_and7 = CFALSE
                      } 
                    if (v_and7 == CFALSE) {g0053I = CFALSE
                    } else { 
                      g0053I = CTRUE} 
                    } 
                  } 
                } 
              } 
            } else {
            g0053I = CFALSE
            } 
          if (g0053I == CTRUE) { 
            { var arg_2 *ClaireList  
              { 
                var v_bag_arg *ClaireAny  
                arg_2= ToType(CEMPTY.Id()).EmptyList()
                if (g0041.Args.Length() == 1) { 
                  v_bag_arg = g0041.Selector.Id()
                  } else {
                  v_bag_arg = g0041.Args.At(1-1)
                  } 
                arg_2.AddFast(v_bag_arg)
                if (g0041.Args.Length() == 1) { 
                  v_bag_arg = g0041.Args.At(1-1)
                  } else {
                  v_bag_arg = g0041.Args.At(2-1)
                  } 
                arg_2.AddFast(v_bag_arg)
                arg_2.AddFast(ToList(OBJ(Core.F_CALL(C_args,ARGS(z.ToEID())))).At(2-1))} 
              Result = F_Call_I_property(ToProperty(OBJ(Core.F_CALL(C_selector,ARGS(z.ToEID())))),arg_2)
              } 
            }  else if (g0041.Selector.Id() == C_nth.Id()) { 
            Result = F_Call_I_property(C_nth_equal,g0041.Args.Copy().AddFast(z))
            }  else if (g0041.Args.Length() == 1) { 
            { var p *ClaireProperty  
              var try_3 EID 
              try_3 = Language.F_make_a_property_any(g0041.Selector.Id())
              if ErrorIn(try_3) {Result = try_3
              } else {
              p = ToProperty(OBJ(try_3))
              { var y *ClaireAny   = g0041.Args.At(1-1)
                var g0054I *ClaireBoolean  
                { 
                  var v_and8 *ClaireBoolean  
                  
                  v_and8 = Equal(p.Id(),C_read.Id())
                  if (v_and8 == CFALSE) {g0054I = CFALSE
                  } else { 
                    if (y.Isa.IsIn(Language.C_Call_plus) == CTRUE) { 
                      { var g0045 *Language.Call_plus   = Language.To_Call_plus(y)
                        v_and8 = Equal(g0045.Selector.Reified.Id(),CTRUE.Id())
                        } 
                      } else {
                      v_and8 = CFALSE
                      } 
                    if (v_and8 == CFALSE) {g0054I = CFALSE
                    } else { 
                      g0054I = CTRUE} 
                    } 
                  } 
                if (g0054I == CTRUE) { 
                  Result = F_Call_I_property(Core.C_write,MakeConstantList(y,z))
                  } else {
                  Result = F_Call_I_property(Core.C_write,MakeConstantList(p.Id(),y,z))
                  } 
                } 
              }
              } 
            } else {
            Result = F_Serror_string(MakeString("[164] ~S cannot be assigned with :="),MakeConstantList(g0041.Id()))
            } 
          } 
        }  else if (x.Isa.IsIn(Language.C_Do) == CTRUE) { 
        { var g0046 *Language.Do   = Language.To_Do(x)
          { var l *ClaireList   = g0046.Args
            { var m int  = l.Length()
              { var v *ClaireVariable  
                { var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
                  _CL_obj.Pname = Core.F_gensym_void()
                  v = _CL_obj
                  } 
                { var _CL_obj *Language.Let_star   = Language.To_Let_star(new(Language.Let_star).Is(Language.C_Let_star))
                  _CL_obj.ClaireVar = v
                  _CL_obj.Value = z
                  { 
                    var va_arg1 *Language.Let  
                    var va_arg2 *ClaireAny  
                    va_arg1 = Language.To_Let(_CL_obj.Id())
                    var try_4 EID 
                    { var _CL_obj *Language.Do   = Language.To_Do(new(Language.Do).Is(Language.C_Do))
                      { 
                        var va_arg1 *Language.Do  
                        var va_arg2 *ClaireList  
                        va_arg1 = _CL_obj
                        var try_5 EID 
                        { var i_bag *ClaireList   = ToType(C_any.Id()).EmptyList()
                          { var i int  = 1
                            { var g0047 int  = m
                              try_5= EID{CFALSE.Id(),0}
                              for (i <= g0047) { 
                                var loop_6 EID 
                                _ = loop_6
                                { 
                                { var arg_7 *ClaireObject  
                                  var try_8 EID 
                                  { var arg_9 *Language.Call  
                                    var try_10 EID 
                                    try_10 = F_Call_I_property(C_nth,MakeConstantList(v.Id(),MakeInteger(i).Id()))
                                    if ErrorIn(try_10) {try_8 = try_10
                                    } else {
                                    arg_9 = Language.To_Call(OBJ(try_10))
                                    try_8 = Language.C_Assign.Make(l.At(i-1),arg_9.Id()).ToEID()
                                    }
                                    } 
                                  if ErrorIn(try_8) {loop_6 = try_8
                                  } else {
                                  arg_7 = ToObject(OBJ(try_8))
                                  loop_6 = EID{i_bag.AddFast(arg_7.Id()).Id(),0}
                                  }
                                  } 
                                if ErrorIn(loop_6) {try_5 = loop_6
                                break
                                } else {
                                i = (i+1)
                                }
                                } 
                              }
                              } 
                            } 
                          if !ErrorIn(try_5) {
                          try_5 = EID{i_bag.Id(),0}
                          }
                          } 
                        if ErrorIn(try_5) {try_4 = try_5
                        } else {
                        va_arg2 = ToList(OBJ(try_5))
                        va_arg1.Args = va_arg2
                        try_4 = EID{va_arg2.Id(),0}
                        }
                        } 
                      if !ErrorIn(try_4) {
                      try_4 = EID{_CL_obj.Id(),0}
                      }
                      } 
                    if ErrorIn(try_4) {Result = try_4
                    } else {
                    va_arg2 = ANY(try_4)
                    va_arg1.Arg = va_arg2
                    Result = va_arg2.ToEID()
                    }
                    } 
                  if !ErrorIn(Result) {
                  Result = EID{_CL_obj.Id(),0}
                  }
                  } 
                } 
              } 
            } 
          } 
        } else {
        Result = Language.C_Assign.Make(x,z).ToEID()
        } 
      }  else if (Equal(y,C_OR.Value) == CTRUE) { 
      if (x.Isa.IsIn(Language.C_Or) == CTRUE) { 
        { var g0049 *Language.Or   = Language.To_Or(x)
          g0049.Args.AddFast(z)
          Result = EID{g0049.Id(),0}
          } 
        } else {
        { var _CL_obj *Language.Or   = Language.To_Or(new(Language.Or).Is(Language.C_Or))
          _CL_obj.Args = MakeConstantList(x,z)
          Result = EID{_CL_obj.Id(),0}
          } 
        } 
      }  else if (Equal(y,C_AND.Value) == CTRUE) { 
      if (x.Isa.IsIn(Language.C_And) == CTRUE) { 
        { var g0051 *Language.And   = Language.To_And(x)
          g0051.Args.AddFast(z)
          Result = EID{g0051.Id(),0}
          } 
        } else {
        { var _CL_obj *Language.And   = Language.To_And(new(Language.And).Is(Language.C_And))
          _CL_obj.Args = MakeConstantList(x,z)
          Result = EID{_CL_obj.Id(),0}
          } 
        } 
      }  else if (y == C__Z.Id()) { 
      Result = F_Call_I_property(ToProperty(C__Z.Id()),MakeConstantList(x,z))
      } else {
      { var arg_11 *Language.Call_star  
        { var _CL_obj *Language.Call_star   = Language.To_Call_star(new(Language.Call_star).Is(Language.C_Call_star))
          _CL_obj.Selector = ToProperty(y)
          _CL_obj.Args = MakeConstantList(x,z)
          arg_11 = _CL_obj
          } 
        Result = F_DBregister_Call(Language.To_Call(arg_11.Id()))
        } 
      } 
    return Result} 
  
// The EID go function for: combine! @ any (throw: true) 
func E_combine_I_any (x EID,y EID,z EID) EID { 
    return F_combine_I_any(ANY(x),ANY(y),ANY(z) )} 
  
// Call* says that combining is OK
// allows to treats Calls, Assigns, Gassign in an homogeneous way
// return false if the pattern is not (x OP y) and OP otherwise
/* {1} The go function for: operation!(x:any) [status=0] */
func F_operation_I_any (x *ClaireAny ) *ClaireAny  { 
    // procedure body with s = any
    var Result *ClaireAny  
    if (x.Isa.IsIn(Language.C_Or) == CTRUE) { 
      Result = C_OR.Value
      }  else if (x.Isa.IsIn(Language.C_And) == CTRUE) { 
      Result = C_AND.Value
      }  else if (x.Isa.IsIn(Language.C_Assign) == CTRUE) { 
      Result = C_L__equal.Id()
      }  else if (x.Isa.IsIn(Language.C_Gassign) == CTRUE) { 
      Result = C_L__equal.Id()
      }  else if (x.Isa.IsIn(Language.C_Call) == CTRUE) { 
      { var g0059 *Language.Call   = Language.To_Call(x)
        { var p *ClaireProperty   = g0059.Selector
          if ((g0059.Isa.IsIn(Language.C_Call_star) == CTRUE) && 
              (F_operation_ask_any(p.Id()) == CTRUE)) { 
            Result = p.Id()
            }  else if (p.Id() == C_nth_equal.Id()) { 
            Result = C_L__equal.Id()
            }  else if (p.Id() == Core.C_write.Id()) { 
            Result = C_L__equal.Id()
            } else {
            Result = CFALSE.Id()
            } 
          } 
        } 
      } else {
      Result = CFALSE.Id()
      } 
    return Result} 
  
// The EID go function for: operation! @ any (throw: false) 
func E_operation_I_any (x EID) EID { 
    return F_operation_I_any(ANY(x) ).ToEID()} 
  
// extract the two operands from an expression x such that operation!(x) != false
/* {1} The go function for: operand!(x:any,n:integer) [status=1] */
func F_operand_I_any (x *ClaireAny ,n int) EID { 
    // eid body s = any
    var Result EID 
    if (x.Isa.IsIn(Language.C_Or) == CTRUE) { 
      { var g0061 *Language.Or   = Language.To_Or(x)
        if (n == 1) { 
          { var _CL_obj *Language.Or   = Language.To_Or(new(Language.Or).Is(Language.C_Or))
            { 
              var va_arg1 *Language.Or  
              var va_arg2 *ClaireList  
              va_arg1 = _CL_obj
              var try_1 EID 
              try_1 = Core.F_rmlast_list(g0061.Args.Copy())
              if ErrorIn(try_1) {Result = try_1
              } else {
              va_arg2 = ToList(OBJ(try_1))
              va_arg1.Args = va_arg2
              Result = EID{va_arg2.Id(),0}
              }
              } 
            if !ErrorIn(Result) {
            Result = EID{_CL_obj.Id(),0}
            }
            } 
          } else {
          Result = Core.F_last_list(g0061.Args)
          } 
        } 
      }  else if (x.Isa.IsIn(Language.C_And) == CTRUE) { 
      { var g0062 *Language.And   = Language.To_And(x)
        if (n == 1) { 
          { var _CL_obj *Language.And   = Language.To_And(new(Language.And).Is(Language.C_And))
            { 
              var va_arg1 *Language.And  
              var va_arg2 *ClaireList  
              va_arg1 = _CL_obj
              var try_2 EID 
              try_2 = Core.F_rmlast_list(g0062.Args.Copy())
              if ErrorIn(try_2) {Result = try_2
              } else {
              va_arg2 = ToList(OBJ(try_2))
              va_arg1.Args = va_arg2
              Result = EID{va_arg2.Id(),0}
              }
              } 
            if !ErrorIn(Result) {
            Result = EID{_CL_obj.Id(),0}
            }
            } 
          } else {
          Result = Core.F_last_list(g0062.Args)
          } 
        } 
      }  else if (x.Isa.IsIn(Language.C_Assign) == CTRUE) { 
      { var g0063 *Language.Assign   = Language.To_Assign(x)
        if (n == 1) { 
          Result = g0063.ClaireVar.ToEID()
          } else {
          Result = g0063.Arg.ToEID()
          } 
        } 
      }  else if (x.Isa.IsIn(Language.C_Gassign) == CTRUE) { 
      { var g0064 *Language.Gassign   = Language.To_Gassign(x)
        if (n == 1) { 
          Result = EID{g0064.ClaireVar.Id(),0}
          } else {
          Result = g0064.Arg.ToEID()
          } 
        } 
      }  else if (x.Isa.IsIn(Language.C_Call) == CTRUE) { 
      { var g0065 *Language.Call   = Language.To_Call(x)
        if (g0065.Selector.Id() == Core.C_write.Id()) { 
          if (n == 2) { 
            Result = g0065.Args.At(3-1).ToEID()
            } else {
            Result = F_Call_I_property(ToProperty(g0065.Args.At(1-1)),MakeConstantList(g0065.Args.At(2-1)))
            } 
          }  else if (g0065.Selector.Id() == C_nth_equal.Id()) { 
          if (n == 2) { 
            Result = Core.F_last_list(g0065.Args)
            } else {
            { var arg_3 *ClaireList  
              var try_4 EID 
              try_4 = Core.F_rmlast_list(g0065.Args.Copy())
              if ErrorIn(try_4) {Result = try_4
              } else {
              arg_3 = ToList(OBJ(try_4))
              Result = F_Call_I_property(C_nth,arg_3)
              }
              } 
            } 
          } else {
          Result = g0065.Args.At(n-1).ToEID()
          } 
        } 
      } else {
      Result = EID{CFALSE.Id(),0}
      } 
    return Result} 
  
// The EID go function for: operand! @ any (throw: true) 
func E_operand_I_any (x EID,n EID) EID { 
    return F_operand_I_any(ANY(x),INT(n) )} 
  
// precedence
//
/* {1} The go function for: precedence!(y:any) [status=0] */
func F_precedence_I_any (y *ClaireAny ) int { 
    if (y == C_as.Id()) { 
      return  0
      }  else if (y == C_L__equal.Id()) { 
      return  100
      }  else if (Equal(y,C_AND.Value) == CTRUE) { 
      return  1000
      }  else if (Equal(y,C_OR.Value) == CTRUE) { 
      return  1010
      } else {
      return  ToOperation(y).Precedence
      } 
    } 
  
// The EID go function for: precedence! @ any (throw: false) 
func E_precedence_I_any (y EID) EID { 
    return EID{C__INT,IVAL(F_precedence_I_any(ANY(y) ))}} 
  
// **********************************************************************
// *   Part 2: read control structures                                  *
// **********************************************************************
/* {1} The go function for: nextstruct(r:meta_reader,%first:keyword,e:keyword) [status=1] */
func (r *MetaReader ) Nextstruct (_Zfirst *ClaireKeyword ,e *ClaireKeyword ) EID { 
    // eid body s = any
    var Result EID 
    if (_Zfirst.Id() == C_let.Id()) { 
      Result = r.Readlet(e)
      }  else if (_Zfirst.Id() == C_when.Id()) { 
      Result = r.Readwhen(e)
      }  else if (_Zfirst.Id() == C_case.Id()) { 
      Result = r.Readcase(e)
      }  else if (_Zfirst.Id() == C_for.Id()) { 
      { var _Zvar *ClaireVariable  
        var try_1 EID 
        { var arg_2 *ClaireAny  
          var try_3 EID 
          try_3 = F_nexts_I_meta_reader1(r,C_in)
          if ErrorIn(try_3) {try_1 = try_3
          } else {
          arg_2 = ANY(try_3)
          try_1 = F_extract_variable_any(arg_2)
          }
          } 
        if ErrorIn(try_1) {Result = try_1
        } else {
        _Zvar = To_Variable(OBJ(try_1))
        { var _Zset *ClaireAny  
          var try_4 EID 
          try_4 = r.Nexte()
          if ErrorIn(try_4) {Result = try_4
          } else {
          _Zset = ANY(try_4)
          { var _Zbind *ClaireList   = r.Bind_I(_Zvar)
            { 
              var x EID 
              if (r.Firstc() == 44) { 
                r.Next()
                } 
              { var _CL_obj *Language.For   = Language.To_For(new(Language.For).Is(Language.C_For))
                _CL_obj.ClaireVar = _Zvar
                _CL_obj.SetArg = _Zset
                { 
                  var va_arg1 *Language.Iteration  
                  var va_arg2 *ClaireAny  
                  va_arg1 = Language.To_Iteration(_CL_obj.Id())
                  var try_5 EID 
                  try_5 = r.Nexts(e)
                  if ErrorIn(try_5) {x = try_5
                  } else {
                  va_arg2 = ANY(try_5)
                  va_arg1.Arg = va_arg2
                  x = va_arg2.ToEID()
                  }
                  } 
                if !ErrorIn(x) {
                x = EID{_CL_obj.Id(),0}
                }
                } 
              if ErrorIn(x) {Result = x
              } else {
              r.Unbind_I(_Zbind)
              Result = x}
              } 
            } 
          }
          } 
        }
        } 
      }  else if (_Zfirst.Id() == C_while.Id()) { 
      { var _CL_obj *Language.While   = Language.To_While(new(Language.While).Is(Language.C_While))
        { 
          var va_arg1 *Language.While  
          var va_arg2 *ClaireAny  
          va_arg1 = _CL_obj
          var try_6 EID 
          try_6 = r.Nexte()
          if ErrorIn(try_6) {Result = try_6
          } else {
          va_arg2 = ANY(try_6)
          va_arg1.Test = va_arg2
          Result = va_arg2.ToEID()
          }
          } 
        if !ErrorIn(Result) {
        { 
          var va_arg1 *Language.While  
          var va_arg2 *ClaireAny  
          va_arg1 = _CL_obj
          var try_7 EID 
          try_7 = r.Nexts(e)
          if ErrorIn(try_7) {Result = try_7
          } else {
          va_arg2 = ANY(try_7)
          va_arg1.Arg = va_arg2
          Result = va_arg2.ToEID()
          }
          } 
        if !ErrorIn(Result) {
        _CL_obj.Other = CFALSE
        Result = EID{_CL_obj.Id(),0}
        }}
        } 
      }  else if (_Zfirst.Id() == C_until.Id()) { 
      { var _CL_obj *Language.While   = Language.To_While(new(Language.While).Is(Language.C_While))
        { 
          var va_arg1 *Language.While  
          var va_arg2 *ClaireAny  
          va_arg1 = _CL_obj
          var try_8 EID 
          try_8 = r.Nexte()
          if ErrorIn(try_8) {Result = try_8
          } else {
          va_arg2 = ANY(try_8)
          va_arg1.Test = va_arg2
          Result = va_arg2.ToEID()
          }
          } 
        if !ErrorIn(Result) {
        { 
          var va_arg1 *Language.While  
          var va_arg2 *ClaireAny  
          va_arg1 = _CL_obj
          var try_9 EID 
          try_9 = r.Nexts(e)
          if ErrorIn(try_9) {Result = try_9
          } else {
          va_arg2 = ANY(try_9)
          va_arg1.Arg = va_arg2
          Result = va_arg2.ToEID()
          }
          } 
        if !ErrorIn(Result) {
        _CL_obj.Other = CTRUE
        Result = EID{_CL_obj.Id(),0}
        }}
        } 
      }  else if (_Zfirst.Id() == C_try.Id()) { 
      { var _Za *ClaireAny  
        var try_10 EID 
        try_10 = F_nexts_I_meta_reader1(r,C_catch)
        if ErrorIn(try_10) {Result = try_10
        } else {
        _Za = ANY(try_10)
        { var _Zt *ClaireAny  
          var try_11 EID 
          try_11 = r.Nexte()
          if ErrorIn(try_11) {Result = try_11
          } else {
          _Zt = ANY(try_11)
          if (C_class.Id() == _Zt.Isa.Id()) { 
            { var arg_12 *ClaireAny  
              var try_13 EID 
              try_13 = r.Nexts(e)
              if ErrorIn(try_13) {Result = try_13
              } else {
              arg_12 = ANY(try_13)
              Result = Language.C_Handle.Make(_Zt,_Za,arg_12).ToEID()
              }
              } 
            } else {
            Result = F_Serror_string(MakeString("[00] in try/catch, ~S is not a class"),MakeConstantList(_Zt))
            } 
          }
          } 
        }
        } 
      } else {
      Result = EID{_Zfirst.Id(),0}
      } 
    return Result} 
  
// The EID go function for: nextstruct @ meta_reader (throw: true) 
func E_nextstruct_meta_reader (r EID,_Zfirst EID,e EID) EID { 
    return ToMetaReader(OBJ(r)).Nextstruct(ToKeyword(OBJ(_Zfirst)),ToKeyword(OBJ(e)) )} 
  
// reads a let expression
//
/* {1} The go function for: readlet(r:meta_reader,e:keyword) [status=1] */
func (r *MetaReader ) Readlet (e *ClaireKeyword ) EID { 
    // eid body s = any
    var Result EID 
    { var _Zdef *ClaireAny  
      var try_1 EID 
      try_1 = F_nexts_I_meta_reader3(r,C_in,44)
      if ErrorIn(try_1) {Result = try_1
      } else {
      _Zdef = ANY(try_1)
      if (_Zdef.Isa.IsIn(Language.C_Assign) == CTRUE) { 
        { var g0067 *Language.Assign   = Language.To_Assign(_Zdef)
          { var v *ClaireVariable  
            var try_2 EID 
            try_2 = F_extract_variable_any(g0067.ClaireVar)
            if ErrorIn(try_2) {Result = try_2
            } else {
            v = To_Variable(OBJ(try_2))
            { var _Zbind *ClaireList   = r.Bind_I(v)
              { 
                var x EID 
                { var _CL_obj *Language.Let   = Language.To_Let(new(Language.Let).Is(Language.C_Let))
                  _CL_obj.ClaireVar = v
                  _CL_obj.Value = g0067.Arg
                  { 
                    var va_arg1 *Language.Let  
                    var va_arg2 *ClaireAny  
                    va_arg1 = _CL_obj
                    var try_3 EID 
                    if (r.Firstc() == 44) { 
                      try_3 = r.Cnext().Readlet(e)
                      } else {
                      try_3 = r.Nexts(e)
                      } 
                    if ErrorIn(try_3) {x = try_3
                    } else {
                    va_arg2 = ANY(try_3)
                    va_arg1.Arg = va_arg2
                    x = va_arg2.ToEID()
                    }
                    } 
                  if !ErrorIn(x) {
                  x = EID{_CL_obj.Id(),0}
                  }
                  } 
                if ErrorIn(x) {Result = x
                } else {
                r.Unbind_I(_Zbind)
                Result = x}
                } 
              } 
            }
            } 
          } 
        }  else if (_Zdef.Isa.IsIn(Language.C_Let_star) == CTRUE) { 
        { var g0068 *Language.Let_star   = Language.To_Let_star(_Zdef)
          { 
            var va_arg1 *Language.Let  
            var va_arg2 *ClaireAny  
            va_arg1 = Language.To_Let(g0068.Id())
            var try_4 EID 
            try_4 = r.Readlet_star(Language.To_Do(g0068.Arg).Args,1,e)
            if ErrorIn(try_4) {Result = try_4
            } else {
            va_arg2 = ANY(try_4)
            va_arg1.Arg = va_arg2
            Result = va_arg2.ToEID()
            }
            } 
          if !ErrorIn(Result) {
          Result = EID{g0068.Id(),0}
          }
          } 
        }  else if ((_Zdef.Isa.IsIn(Language.C_Call) == CTRUE) && 
          (ANY(Core.F_CALL(C_selector,ARGS(_Zdef.ToEID()))) == Core.C_write.Id())) { 
        { var g0069 *Language.Call   = Language.To_Call(_Zdef)
          { var v1 *ClaireVariable  
            { var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
              _CL_obj.Pname = Core.F_gensym_void()
              _CL_obj.Range = ToType(C_any.Id())
              v1 = _CL_obj
              } 
            { var v2 *ClaireVariable  
              { var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
                _CL_obj.Pname = Core.F_gensym_void()
                _CL_obj.Range = ToType(C_any.Id())
                v2 = _CL_obj
                } 
              { var _Za *ClaireList   = g0069.Args
                { var _Ze *ClaireAny  
                  var try_5 EID 
                  try_5 = r.Nexts(e)
                  if ErrorIn(try_5) {Result = try_5
                  } else {
                  _Ze = ANY(try_5)
                  { var _CL_obj *Language.Let_plus   = Language.To_Let_plus(new(Language.Let_plus).Is(Language.C_Let_plus))
                    _CL_obj.ClaireVar = v1
                    { 
                      var va_arg1 *Language.Let  
                      var va_arg2 *ClaireAny  
                      va_arg1 = Language.To_Let(_CL_obj.Id())
                      { var _CL_obj *Language.Call_plus   = Language.To_Call_plus(new(Language.Call_plus).Is(Language.C_Call_plus))
                        _CL_obj.Selector = ToProperty(_Za.At(1-1))
                        _CL_obj.Args = MakeConstantList(_Za.At(2-1))
                        va_arg2 = _CL_obj.Id()
                        } 
                      va_arg1.Value = va_arg2
                      } 
                    { 
                      var va_arg1 *Language.Let  
                      var va_arg2 *ClaireAny  
                      va_arg1 = Language.To_Let(_CL_obj.Id())
                      var try_6 EID 
                      { var _CL_obj *Language.Do   = Language.To_Do(new(Language.Do).Is(Language.C_Do))
                        { 
                          var va_arg1 *Language.Do  
                          var va_arg2 *ClaireList  
                          va_arg1 = _CL_obj
                          var try_7 EID 
                          { 
                            var v_bag_arg *ClaireAny  
                            try_7= EID{ToType(C_any.Id()).EmptyList().Id(),0}
                            ToList(OBJ(try_7)).AddFast(g0069.Id())
                            var try_8 EID 
                            { var _CL_obj *Language.Let   = Language.To_Let(new(Language.Let).Is(Language.C_Let))
                              _CL_obj.ClaireVar = v2
                              _CL_obj.Value = _Ze
                              { 
                                var va_arg1 *Language.Let  
                                var va_arg2 *ClaireAny  
                                va_arg1 = _CL_obj
                                var try_9 EID 
                                { var _CL_obj *Language.Do   = Language.To_Do(new(Language.Do).Is(Language.C_Do))
                                  { 
                                    var va_arg1 *Language.Do  
                                    var va_arg2 *ClaireList  
                                    va_arg1 = _CL_obj
                                    var try_10 EID 
                                    { 
                                      var v_bag_arg *ClaireAny  
                                      try_10= EID{ToType(C_any.Id()).EmptyList().Id(),0}
                                      var try_11 EID 
                                      try_11 = F_Call_I_property(Core.C_write,MakeConstantList(_Za.At(1-1),_Za.At(2-1),v1.Id()))
                                      if ErrorIn(try_11) {try_10 = try_11
                                      } else {
                                      v_bag_arg = ANY(try_11)
                                      ToList(OBJ(try_10)).AddFast(v_bag_arg)
                                      ToList(OBJ(try_10)).AddFast(v2.Id())}
                                      } 
                                    if ErrorIn(try_10) {try_9 = try_10
                                    } else {
                                    va_arg2 = ToList(OBJ(try_10))
                                    va_arg1.Args = va_arg2
                                    try_9 = EID{va_arg2.Id(),0}
                                    }
                                    } 
                                  if !ErrorIn(try_9) {
                                  try_9 = EID{_CL_obj.Id(),0}
                                  }
                                  } 
                                if ErrorIn(try_9) {try_8 = try_9
                                } else {
                                va_arg2 = ANY(try_9)
                                va_arg1.Arg = va_arg2
                                try_8 = va_arg2.ToEID()
                                }
                                } 
                              if !ErrorIn(try_8) {
                              try_8 = EID{_CL_obj.Id(),0}
                              }
                              } 
                            if ErrorIn(try_8) {try_7 = try_8
                            } else {
                            v_bag_arg = ANY(try_8)
                            ToList(OBJ(try_7)).AddFast(v_bag_arg)}
                            } 
                          if ErrorIn(try_7) {try_6 = try_7
                          } else {
                          va_arg2 = ToList(OBJ(try_7))
                          va_arg1.Args = va_arg2
                          try_6 = EID{va_arg2.Id(),0}
                          }
                          } 
                        if !ErrorIn(try_6) {
                        try_6 = EID{_CL_obj.Id(),0}
                        }
                        } 
                      if ErrorIn(try_6) {Result = try_6
                      } else {
                      va_arg2 = ANY(try_6)
                      va_arg1.Arg = va_arg2
                      Result = va_arg2.ToEID()
                      }
                      } 
                    if !ErrorIn(Result) {
                    Result = EID{_CL_obj.Id(),0}
                    }
                    } 
                  }
                  } 
                } 
              } 
            } 
          } 
        } else {
        Result = F_Serror_string(MakeString("[165] ~S is illegal after a let"),MakeConstantList(_Zdef))
        } 
      }
      } 
    return Result} 
  
// The EID go function for: readlet @ meta_reader (throw: true) 
func E_readlet_meta_reader (r EID,e EID) EID { 
    return ToMetaReader(OBJ(r)).Readlet(ToKeyword(OBJ(e)) )} 
  
// recursive construction of the tail of a Let*
//
/* {1} The go function for: readlet*(r:meta_reader,l:list,n:integer,e:keyword) [status=1] */
func (r *MetaReader ) Readlet_star (l *ClaireList ,n int,e *ClaireKeyword ) EID { 
    // eid body s = any
    var Result EID 
    if (n > l.Length()) { 
      Result = r.Nexts(e)
      } else {
      { var v *ClaireVariable  
        var try_1 EID 
        try_1 = F_extract_variable_any(ANY(Core.F_CALL(Language.C_var,ARGS(l.At(n-1).ToEID()))))
        if ErrorIn(try_1) {Result = try_1
        } else {
        v = To_Variable(OBJ(try_1))
        { var _Zbind *ClaireList   = r.Bind_I(v)
          { 
            var x EID 
            { var _CL_obj *Language.Let   = Language.To_Let(new(Language.Let).Is(Language.C_Let))
              _CL_obj.ClaireVar = v
              _CL_obj.Value = ANY(Core.F_CALL(C_arg,ARGS(l.At(n-1).ToEID())))
              { 
                var va_arg1 *Language.Let  
                var va_arg2 *ClaireAny  
                va_arg1 = _CL_obj
                var try_2 EID 
                try_2 = r.Readlet_star(l,(n+1),e)
                if ErrorIn(try_2) {x = try_2
                } else {
                va_arg2 = ANY(try_2)
                va_arg1.Arg = va_arg2
                x = va_arg2.ToEID()
                }
                } 
              if !ErrorIn(x) {
              x = EID{_CL_obj.Id(),0}
              }
              } 
            if ErrorIn(x) {Result = x
            } else {
            r.Unbind_I(_Zbind)
            Result = x}
            } 
          } 
        }
        } 
      } 
    return Result} 
  
// The EID go function for: readlet* @ meta_reader (throw: true) 
func E_readlet_star_meta_reader (r EID,l EID,n EID,e EID) EID { 
    return ToMetaReader(OBJ(r)).Readlet_star(ToList(OBJ(l)),
      INT(n),
      ToKeyword(OBJ(e)) )} 
  
// reads a when expression
//
/* {1} The go function for: readwhen(r:meta_reader,e:keyword) [status=1] */
func (r *MetaReader ) Readwhen (e *ClaireKeyword ) EID { 
    // eid body s = any
    var Result EID 
    { var _Zdef *ClaireAny  
      var try_1 EID 
      try_1 = F_nexts_I_meta_reader3(r,C_in,44)
      if ErrorIn(try_1) {Result = try_1
      } else {
      _Zdef = ANY(try_1)
      if (_Zdef.Isa.IsIn(Language.C_Assign) == CTRUE) { 
        { var g0071 *Language.Assign   = Language.To_Assign(_Zdef)
          { var v *ClaireVariable  
            var try_2 EID 
            try_2 = F_extract_variable_any(g0071.ClaireVar)
            if ErrorIn(try_2) {Result = try_2
            } else {
            v = To_Variable(OBJ(try_2))
            { var _Zbind *ClaireList   = r.Bind_I(v)
              { var _Za *ClaireAny  
                var try_3 EID 
                try_3 = r.Nexts(ToKeyword(C_else.Id()))
                if ErrorIn(try_3) {Result = try_3
                } else {
                _Za = ANY(try_3)
                { 
                  var x EID 
                  { var _CL_obj *Language.When   = Language.To_When(new(Language.When).Is(Language.C_When))
                    _CL_obj.ClaireVar = v
                    _CL_obj.Value = g0071.Arg
                    _CL_obj.Arg = _Za
                    { 
                      var va_arg1 *Language.When  
                      var va_arg2 *ClaireAny  
                      va_arg1 = _CL_obj
                      var try_4 EID 
                      if (F_boolean_I_any(F_stop_ask_integer(r.Firstc())) == CTRUE) { 
                        try_4 = EID{CNULL,0}
                        } else {
                        try_4 = r.Nexts(e)
                        } 
                      if ErrorIn(try_4) {x = try_4
                      } else {
                      va_arg2 = ANY(try_4)
                      va_arg1.Other = va_arg2
                      x = va_arg2.ToEID()
                      }
                      } 
                    if !ErrorIn(x) {
                    x = EID{_CL_obj.Id(),0}
                    }
                    } 
                  if ErrorIn(x) {Result = x
                  } else {
                  r.Unbind_I(_Zbind)
                  Result = x}
                  } 
                }
                } 
              } 
            }
            } 
          } 
        } else {
        Result = F_Serror_string(MakeString("[165] ~S is illegal after a when"),MakeConstantList(_Zdef))
        } 
      }
      } 
    return Result} 
  
// The EID go function for: readwhen @ meta_reader (throw: true) 
func E_readwhen_meta_reader (r EID,e EID) EID { 
    return ToMetaReader(OBJ(r)).Readwhen(ToKeyword(OBJ(e)) )} 
  
// read an if
//
/* {1} The go function for: readif(r:meta_reader,e:integer) [status=1] */
func (r *MetaReader ) Readif (e int) EID { 
    // eid body s = any
    var Result EID 
    { var _Za1 *ClaireAny  
      var try_1 EID 
      try_1 = r.Nexte()
      if ErrorIn(try_1) {Result = try_1
      } else {
      _Za1 = ANY(try_1)
      { var _Za2 *ClaireAny  
        var try_2 EID 
        try_2 = r.Nexts(ToKeyword(C_else.Id()))
        if ErrorIn(try_2) {Result = try_2
        } else {
        _Za2 = ANY(try_2)
        { var arg_3 *ClaireAny  
          var try_4 EID 
          if ((r.Firstc() == 44) || 
              (r.Firstc() == e)) { 
            try_4 = EID{CFALSE.Id(),0}
            } else {
            { var x *ClaireAny  
              var try_5 EID 
              try_5 = r.Nexte()
              if ErrorIn(try_5) {try_4 = try_5
              } else {
              x = ANY(try_5)
              if (x == C_if.Id()) { 
                try_4 = r.Readif(e)
                }  else if (F_keyword_ask_any(x) == CTRUE) { 
                try_4 = r.Nextstruct(ToKeyword(x),C_none)
                } else {
                try_4 = r.Loopexp(x,C_none,CFALSE)
                } 
              }
              } 
            } 
          if ErrorIn(try_4) {Result = try_4
          } else {
          arg_3 = ANY(try_4)
          Result = Language.C_If.Make(_Za1,_Za2,arg_3).ToEID()
          }
          } 
        }
        } 
      }
      } 
    return Result} 
  
// The EID go function for: readif @ meta_reader (throw: true) 
func E_readif_meta_reader (r EID,e EID) EID { 
    return ToMetaReader(OBJ(r)).Readif(INT(e) )} 
  
// reads a member_of
//
/* {1} The go function for: readcase(r:meta_reader,e:keyword) [status=1] */
func (r *MetaReader ) Readcase (e *ClaireKeyword ) EID { 
    // eid body s = any
    var Result EID 
    { var _Zv *ClaireAny  
      var try_1 EID 
      try_1 = r.Nexte()
      if ErrorIn(try_1) {Result = try_1
      } else {
      _Zv = ANY(try_1)
      var g0073I *ClaireBoolean  
      var try_2 EID 
      { var arg_3 int 
        var try_4 EID 
        try_4 = r.Skipc_I()
        if ErrorIn(try_4) {try_2 = try_4
        } else {
        arg_3 = INT(try_4)
        try_2 = EID{Core.F__I_equal_any(MakeInteger(arg_3).Id(),MakeInteger(40).Id()).Id(),0}
        }
        } 
      if ErrorIn(try_2) {Result = try_2
      } else {
      g0073I = ToBoolean(OBJ(try_2))
      if (g0073I == CTRUE) { 
        Result = F_Serror_string(MakeString("[166] Missing ( after case ~S"),MakeConstantList(_Zv))
        } else {
        Result = EID{CFALSE.Id(),0}
        } 
      }
      if !ErrorIn(Result) {
      { var _Zx *Language.Case  
        { var _CL_obj *Language.Case   = Language.To_Case(new(Language.Case).Is(Language.C_Case))
          _CL_obj.ClaireVar = _Zv
          _CL_obj.Args = ToType(CEMPTY.Id()).EmptyList()
          _Zx = _CL_obj
          } 
        { var _Zt *ClaireAny   = C_any.Id()
          Result= EID{CFALSE.Id(),0}
          for (r.Firstc() != 41) { 
            var loop_5 EID 
            _ = loop_5
            { 
            r.Next()
            var try_6 EID 
            { var arg_7 *ClaireAny  
              var try_8 EID 
              try_8 = r.Nexte()
              if ErrorIn(try_8) {try_6 = try_8
              } else {
              arg_7 = ANY(try_8)
              try_6 = Language.F_extract_type_any(arg_7)
              }
              } 
            if ErrorIn(try_6) {loop_5 = try_6
            Result = try_6
            break
            } else {
            _Zt = ANY(try_6)
            loop_5 = _Zt.ToEID()
            { 
              var va_arg1 *Language.Case  
              var va_arg2 *ClaireList  
              va_arg1 = _Zx
              var try_9 EID 
              { var arg_10 *ClaireAny  
                var try_11 EID 
                try_11 = r.Nexts(C_none)
                if ErrorIn(try_11) {try_9 = try_11
                } else {
                arg_10 = ANY(try_11)
                try_9 = EID{_Zx.Args.AddFast(_Zt).AddFast(arg_10).Id(),0}
                }
                } 
              if ErrorIn(try_9) {loop_5 = try_9
              } else {
              va_arg2 = ToList(OBJ(try_9))
              va_arg1.Args = va_arg2
              loop_5 = EID{va_arg2.Id(),0}
              }
              } 
            if ErrorIn(loop_5) {Result = loop_5
            break
            } else {
            if ((F_boolean_I_any(F_stop_ask_integer(r.Firstc())).Id() != CTRUE.Id()) && 
                (F_boolean_I_any(F_stop_ask_integer(r.Skipc())).Id() != CTRUE.Id())) { 
              loop_5 = F_Serror_string(MakeString("[167] missing ) or , after ~S"),MakeConstantList(_Zx.Id()))
              } else {
              loop_5 = EID{CFALSE.Id(),0}
              } 
            if ErrorIn(loop_5) {Result = loop_5
            break
            } else {
            }}}
            } 
          }
          if !ErrorIn(Result) {
          r.Next()
          var g0074I *ClaireBoolean  
          var try_12 EID 
          { 
            var v_and5 *ClaireBoolean  
            
            v_and5 = Core.F__I_equal_any(e.Id(),C_none.Id())
            if (v_and5 == CFALSE) {try_12 = EID{CFALSE.Id(),0}
            } else { 
              v_and5 = Core.F__I_equal_any(F_boolean_I_any(F_stop_ask_integer(r.Skipc())).Id(),CTRUE.Id())
              if (v_and5 == CFALSE) {try_12 = EID{CFALSE.Id(),0}
              } else { 
                var try_13 EID 
                { var arg_14 *ClaireAny  
                  var try_15 EID 
                  try_15 = r.Nexte()
                  if ErrorIn(try_15) {try_13 = try_15
                  } else {
                  arg_14 = ANY(try_15)
                  try_13 = EID{Core.F__I_equal_any(arg_14,e.Id()).Id(),0}
                  }
                  } 
                if ErrorIn(try_13) {try_12 = try_13
                } else {
                v_and5 = ToBoolean(OBJ(try_13))
                if (v_and5 == CFALSE) {try_12 = EID{CFALSE.Id(),0}
                } else { 
                  try_12 = EID{CTRUE.Id(),0}} 
                } 
              } 
            }
            } 
          if ErrorIn(try_12) {Result = try_12
          } else {
          g0074I = ToBoolean(OBJ(try_12))
          if (g0074I == CTRUE) { 
            Result = F_Serror_string(MakeString("[161] missing ~S after ~S"),MakeConstantList(e.Id(),_Zx.Id()))
            } else {
            Result = EID{CFALSE.Id(),0}
            } 
          }
          if !ErrorIn(Result) {
          Result = EID{_Zx.Id(),0}
          }}
          } 
        } 
      }
      }
      } 
    return Result} 
  
// The EID go function for: readcase @ meta_reader (throw: true) 
func E_readcase_meta_reader (r EID,e EID) EID { 
    return ToMetaReader(OBJ(r)).Readcase(ToKeyword(OBJ(e)) )} 
  
// if the expression begins with "{"
//
/* {1} The go function for: readset(r:meta_reader,%a1:any) [status=1] */
func (r *MetaReader ) Readset (_Za1 *ClaireAny ) EID { 
    // eid body s = any
    var Result EID 
    if (Equal(_Za1,r.Curly) == CTRUE) { 
      r.Next()
      Result = EID{CEMPTY.Id(),0}
      } else {
      if (F_keyword_ask_any(_Za1) == CTRUE) { 
        var try_1 EID 
        try_1 = r.Nextstruct(ToKeyword(_Za1),C_none)
        if ErrorIn(try_1) {Result = try_1
        } else {
        _Za1 = ANY(try_1)
        Result = _Za1.ToEID()
        }
        } else {
        Result = EID{CFALSE.Id(),0}
        } 
      if !ErrorIn(Result) {
      { var _Za2 *ClaireAny  
        var try_2 EID 
        try_2 = r.Nexte()
        if ErrorIn(try_2) {Result = try_2
        } else {
        _Za2 = ANY(try_2)
        if (Equal(_Za2,r.Comma) == CTRUE) { 
          { var arg_3 *ClaireSet  
            var try_4 EID 
            { var u_bag *ClaireSet   = ToType(CEMPTY.Id()).EmptySet()
              { 
                var u *ClaireAny  
                _ = u
                try_4= EID{CFALSE.Id(),0}
                var u_support *ClaireList  
                var try_5 EID 
                { var arg_6 *ClaireAny  
                  var try_7 EID 
                  try_7 = r.Cnext().Nextseq(125)
                  if ErrorIn(try_7) {try_5 = try_7
                  } else {
                  arg_6 = ANY(try_7)
                  try_5 = EID{F_cons_any(_Za1,ToList(arg_6)).Id(),0}
                  }
                  } 
                if ErrorIn(try_5) {try_4 = try_5
                } else {
                u_support = ToList(OBJ(try_5))
                u_len := u_support.Length()
                for i_it := 0; i_it < u_len; i_it++ { 
                  u = u_support.At(i_it)
                  var loop_8 EID 
                  _ = loop_8
                  { var arg_9 *ClaireAny  
                    var try_10 EID 
                    try_10 = F_dereference_any(u)
                    if ErrorIn(try_10) {loop_8 = try_10
                    } else {
                    arg_9 = ANY(try_10)
                    loop_8 = EID{u_bag.AddFast(arg_9).Id(),0}
                    }
                    } 
                  if ErrorIn(loop_8) {try_4 = loop_8
                  break
                  } else {
                  }}
                  } 
                } 
              if !ErrorIn(try_4) {
              try_4 = EID{u_bag.Id(),0}
              }
              } 
            if ErrorIn(try_4) {Result = try_4
            } else {
            arg_3 = ToSet(OBJ(try_4))
            Result = EID{arg_3.Cast_I(ToType(CEMPTY.Id())).Id(),0}
            }
            } 
          }  else if (Equal(_Za2,r.Curly) == CTRUE) { 
          r.Next()
          { var arg_11 *ClaireSet  
            var try_12 EID 
            { 
              var v_bag_arg *ClaireAny  
              try_12= EID{ToType(CEMPTY.Id()).EmptySet().Id(),0}
              var try_13 EID 
              try_13 = F_dereference_any(_Za1)
              if ErrorIn(try_13) {try_12 = try_13
              } else {
              v_bag_arg = ANY(try_13)
              ToSet(OBJ(try_12)).AddFast(v_bag_arg)}
              } 
            if ErrorIn(try_12) {Result = try_12
            } else {
            arg_11 = ToSet(OBJ(try_12))
            Result = EID{arg_11.Cast_I(ToType(CEMPTY.Id())).Id(),0}
            }
            } 
          }  else if (_Za2 == C_in.Id()) { 
          { var v *ClaireVariable  
            var try_14 EID 
            try_14 = F_extract_variable_any(_Za1)
            if ErrorIn(try_14) {Result = try_14
            } else {
            v = To_Variable(OBJ(try_14))
            { var _CL_obj *Language.Select   = Language.To_Select(new(Language.Select).Is(Language.C_Select))
              _CL_obj.ClaireVar = v
              { 
                var va_arg1 *Language.Iteration  
                var va_arg2 *ClaireAny  
                va_arg1 = Language.To_Iteration(_CL_obj.Id())
                var try_15 EID 
                try_15 = r.Nexte()
                if ErrorIn(try_15) {Result = try_15
                } else {
                va_arg2 = ANY(try_15)
                va_arg1.SetArg = va_arg2
                Result = va_arg2.ToEID()
                }
                } 
              if !ErrorIn(Result) {
              { 
                var va_arg1 *Language.Iteration  
                var va_arg2 *ClaireAny  
                va_arg1 = Language.To_Iteration(_CL_obj.Id())
                var try_16 EID 
                { var _Zbind *ClaireList   = r.Bind_I(v)
                  { 
                    var x EID 
                    var g0075I *ClaireBoolean  
                    var try_17 EID 
                    { var arg_18 *ClaireAny  
                      var try_19 EID 
                      try_19 = r.Nexte()
                      if ErrorIn(try_19) {try_17 = try_19
                      } else {
                      arg_18 = ANY(try_19)
                      try_17 = EID{Core.F__I_equal_any(arg_18,C_OR.Value).Id(),0}
                      }
                      } 
                    if ErrorIn(try_17) {x = try_17
                    } else {
                    g0075I = ToBoolean(OBJ(try_17))
                    if (g0075I == CTRUE) { 
                      x = F_Serror_string(MakeString("[168] missing | in selection"),CNIL)
                      } else {
                      x = F_nexts_I_meta_reader2(r,125)
                      } 
                    }
                    if ErrorIn(x) {try_16 = x
                    } else {
                    r.Unbind_I(_Zbind)
                    try_16 = x}
                    } 
                  } 
                if ErrorIn(try_16) {Result = try_16
                } else {
                va_arg2 = ANY(try_16)
                va_arg1.Arg = va_arg2
                Result = va_arg2.ToEID()
                }
                } 
              if !ErrorIn(Result) {
              Result = EID{_CL_obj.Id(),0}
              }}
              } 
            }
            } 
          }  else if (Equal(_Za2,C_OR.Value) == CTRUE) { 
          { var v *ClaireVariable  
            var try_20 EID 
            { var arg_21 *ClaireAny  
              var try_22 EID 
              try_22 = F_nexts_I_meta_reader1(r,C_in)
              if ErrorIn(try_22) {try_20 = try_22
              } else {
              arg_21 = ANY(try_22)
              try_20 = F_extract_variable_any(arg_21)
              }
              } 
            if ErrorIn(try_20) {Result = try_20
            } else {
            v = To_Variable(OBJ(try_20))
            { var arg_23 *Language.Image  
              var try_24 EID 
              { var _CL_obj *Language.Image   = Language.To_Image(new(Language.Image).Is(Language.C_Image))
                _CL_obj.ClaireVar = v
                { 
                  var va_arg1 *Language.Iteration  
                  var va_arg2 *ClaireAny  
                  va_arg1 = Language.To_Iteration(_CL_obj.Id())
                  var try_25 EID 
                  try_25 = F_nexts_I_meta_reader2(r,125)
                  if ErrorIn(try_25) {try_24 = try_25
                  } else {
                  va_arg2 = ANY(try_25)
                  va_arg1.SetArg = va_arg2
                  try_24 = va_arg2.ToEID()
                  }
                  } 
                if !ErrorIn(try_24) {
                _CL_obj.Arg = Language.F_substitution_any(_Za1,v,v.Id())
                try_24 = EID{_CL_obj.Id(),0}
                }
                } 
              if ErrorIn(try_24) {Result = try_24
              } else {
              arg_23 = Language.To_Image(OBJ(try_24))
              Result = Language.F_lexical_build_any(arg_23.Id(),MakeConstantList(v.Id()),0)
              }
              } 
            }
            } 
          }  else if (F_operation_ask_any(_Za2) == CTRUE) { 
          { var arg_26 *ClaireAny  
            var try_27 EID 
            { var arg_28 *ClaireAny  
              var try_29 EID 
              { var arg_30 *ClaireAny  
                var try_31 EID 
                try_31 = r.Nexte()
                if ErrorIn(try_31) {try_29 = try_31
                } else {
                arg_30 = ANY(try_31)
                try_29 = F_combine_any(_Za1,_Za2,arg_30)
                }
                } 
              if ErrorIn(try_29) {try_27 = try_29
              } else {
              arg_28 = ANY(try_29)
              try_27 = r.Loopexp(arg_28,C_none,CFALSE)
              }
              } 
            if ErrorIn(try_27) {Result = try_27
            } else {
            arg_26 = ANY(try_27)
            Result = r.Readset(arg_26)
            }
            } 
          } else {
          Result = F_Serror_string(MakeString("[169] missing separation between ~S and ~S"),MakeConstantList(_Za1,_Za2))
          } 
        }
        } 
      }
      } 
    return Result} 
  
// The EID go function for: readset @ meta_reader (throw: true) 
func E_readset_meta_reader (r EID,_Za1 EID) EID { 
    return ToMetaReader(OBJ(r)).Readset(ANY(_Za1) )} 
  
/* {1} The go function for: dereference(x:any) [status=1] */
func F_dereference_any (x *ClaireAny ) EID { 
    // eid body s = any
    var Result EID 
    if (x.Isa.IsIn(C_unbound_symbol) == CTRUE) { 
      { var g0076 *ClaireUnboundSymbol   = ToUnboundSymbol(x)
        Result = ToException(Core.C_general_error.Make(MakeString("[170] cannot use ~S in a set constant").Id(),MakeConstantList(g0076.Id()).Id())).Close()
        } 
      }  else if (x.Isa.IsIn(C_Variable) == CTRUE) { 
      { var g0077 *ClaireVariable   = To_Variable(x)
        Result = ToException(Core.C_general_error.Make(MakeString("[170] cannot use a variable (~S) in a set constant").Id(),MakeConstantList(g0077.Id()).Id())).Close()
        } 
      } else {
      Result = EVAL(x)
      } 
    return Result} 
  
// The EID go function for: dereference @ any (throw: true) 
func E_dereference_any (x EID) EID { 
    return F_dereference_any(ANY(x) )} 
  
// reads a sequence of exp. Must end with a e = ) | ] | }
//
/* {1} The go function for: nextseq(r:meta_reader,e:integer) [status=1] */
func (r *MetaReader ) Nextseq (e int) EID { 
    // eid body s = any
    var Result EID 
    
    if (r.Firstc() == e) { 
      r.Next()
      Result = EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
      } else {
      { var x *ClaireAny  
        var try_1 EID 
        try_1 = r.Nexts(ToKeyword(IfThenElse((e == 62),
          C_None.Id(),
          C_none.Id())))
        if ErrorIn(try_1) {Result = try_1
        } else {
        x = ANY(try_1)
        if ((r.Firstc() == 10) && 
            (r.Toplevel == CTRUE)) { 
          r.Skipc()
          } 
        if (r.Firstc() == e) { 
          r.Next()
          Result = EID{MakeConstantList(x).Id(),0}
          }  else if (r.Firstc() == 44) { 
          { var arg_2 *ClaireAny  
            var try_3 EID 
            try_3 = r.Cnext().Nextseq(e)
            if ErrorIn(try_3) {Result = try_3
            } else {
            arg_2 = ANY(try_3)
            Result = EID{F_cons_any(x,ToList(arg_2)).Id(),0}
            }
            } 
          } else {
          Result = F_Serror_string(MakeString("[171] Read the character ~S inside a sequence"),MakeConstantList(MakeChar(F_char_I_integer(r.Firstc())).Id()))
          } 
        }
        } 
      } 
    return Result} 
  
// The EID go function for: nextseq @ meta_reader (throw: true) 
func E_nextseq_meta_reader (r EID,e EID) EID { 
    return ToMetaReader(OBJ(r)).Nextseq(INT(e) )} 
  
// read the next block: a sequence of exp. Must end with a e = ) | ] | }
//
/* {1} The go function for: readblock(r:meta_reader,x:any,e:integer) [status=1] */
func (r *MetaReader ) Readblock (x *ClaireAny ,e int) EID { 
    // eid body s = any
    var Result EID 
    r.Skipc()
    if (Equal(x,r.Paren) == CTRUE) { 
      Result = EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
      }  else if (r.Firstc() == 44) { 
      { var arg_1 *ClaireAny  
        var try_2 EID 
        { var arg_3 *ClaireAny  
          var try_4 EID 
          try_4 = r.Cnext().Nexte()
          if ErrorIn(try_4) {try_2 = try_4
          } else {
          arg_3 = ANY(try_4)
          try_2 = r.Readblock(arg_3,e)
          }
          } 
        if ErrorIn(try_2) {Result = try_2
        } else {
        arg_1 = ANY(try_2)
        Result = F_Do_I_any(x,arg_1)
        }
        } 
      }  else if (r.Firstc() == e) { 
      r.Cnext()
      Result = x.ToEID()
      }  else if (F_boolean_I_any(F_stop_ask_integer(r.Firstc())) == CTRUE) { 
      Result = F_Serror_string(MakeString("[172] the sequence ...~S must end with ~A"),MakeConstantList(x,MakeChar(F_char_I_integer(e)).Id()))
      }  else if (x == C_if.Id()) { 
      { var arg_5 *ClaireAny  
        var try_6 EID 
        try_6 = r.Readif(e)
        if ErrorIn(try_6) {Result = try_6
        } else {
        arg_5 = ANY(try_6)
        Result = r.Readblock(arg_5,e)
        }
        } 
      }  else if (x == C_Zif.Id()) { 
      { var _Zi *Language.If  
        var try_7 EID 
        try_7 = r.Readif(e)
        if ErrorIn(try_7) {Result = try_7
        } else {
        _Zi = Language.To_If(OBJ(try_7))
        { var arg_8 *ClaireAny  
          var try_9 EID 
          var g0080I *ClaireBoolean  
          var try_10 EID 
          { var arg_11 *ClaireAny  
            var try_12 EID 
            try_12 = EVAL(_Zi.Test)
            if ErrorIn(try_12) {try_10 = try_12
            } else {
            arg_11 = ANY(try_12)
            try_10 = EID{F_boolean_I_any(arg_11).Id(),0}
            }
            } 
          if ErrorIn(try_10) {try_9 = try_10
          } else {
          g0080I = ToBoolean(OBJ(try_10))
          if (g0080I == CTRUE) { 
            try_9 = _Zi.Arg.ToEID()
            } else {
            try_9 = _Zi.Other.ToEID()
            } 
          }
          if ErrorIn(try_9) {Result = try_9
          } else {
          arg_8 = ANY(try_9)
          Result = r.Readblock(arg_8,e)
          }
          } 
        }
        } 
      }  else if (x == C_else.Id()) { 
      Result = F_Serror_string(MakeString("[173] Expression starting with else"),CNIL)
      }  else if (F_keyword_ask_any(x) == CTRUE) { 
      { var arg_13 *ClaireAny  
        var try_14 EID 
        try_14 = r.Nextstruct(ToKeyword(x),C_none)
        if ErrorIn(try_14) {Result = try_14
        } else {
        arg_13 = ANY(try_14)
        Result = r.Readblock(arg_13,e)
        }
        } 
      } else {
      { var y *ClaireAny  
        var try_15 EID 
        try_15 = r.Loopexp(x,C_none,CFALSE)
        if ErrorIn(try_15) {Result = try_15
        } else {
        y = ANY(try_15)
        if (y.Isa.IsIn(Language.C_Call_star) == CTRUE) { 
          { var g0079 *Language.Call_star   = Language.To_Call_star(y)
            g0079.Isa = Language.C_Call
            } 
          } 
        Result = r.Readblock(y,e)
        }
        } 
      } 
    return Result} 
  
// The EID go function for: readblock @ meta_reader (throw: true) 
func E_readblock_meta_reader (r EID,x EID,e EID) EID { 
    return ToMetaReader(OBJ(r)).Readblock(ANY(x),INT(e) )} 
  
// variant in CLAIRE4 when e = ), which can also read a lambda
/* {1} The go function for: readList(r:meta_reader,x:any) [status=1] */
func (r *MetaReader ) ReadList (x *ClaireAny ) EID { 
    // eid body s = any
    var Result EID 
    { var y *ClaireAny  
      var try_1 EID 
      try_1 = r.Readblock(x,41)
      if ErrorIn(try_1) {Result = try_1
      } else {
      y = ANY(try_1)
      
      if (r.Firstc() == 123) { 
        Result = r.Readlambda(y)
        } else {
        Result = y.ToEID()
        } 
      }
      } 
    return Result} 
  
// The EID go function for: readList @ meta_reader (throw: true) 
func E_Reader_readList_meta_reader (r EID,x EID) EID { 
    return ToMetaReader(OBJ(r)).ReadList(ANY(x) )} 
  
// create the lambda
/* {1} The go function for: readlambda(r:meta_reader,l:any) [status=1] */
func (r *MetaReader ) Readlambda (l *ClaireAny ) EID { 
    // eid body s = any
    var Result EID 
    { var lbody *ClaireAny  
      var try_1 EID 
      try_1 = r.Cnext().Nextseq(125)
      if ErrorIn(try_1) {Result = try_1
      } else {
      lbody = ANY(try_1)
      { var lvar *ClaireList   = ToType(CEMPTY.Id()).EmptyList()
        
        if (l.Isa.IsIn(Language.C_Vardef) == CTRUE) { 
          { var g0081 *Language.Vardef   = Language.To_Vardef(l)
            lvar = lvar.AddFast(g0081.Id())
            Result = EID{lvar.Id(),0}
            } 
          }  else if (l.Isa.IsIn(Language.C_Do) == CTRUE) { 
          { var g0082 *Language.Do   = Language.To_Do(l)
            { 
              var y *ClaireAny  
              _ = y
              Result= EID{CFALSE.Id(),0}
              var y_support *ClaireList  
              y_support = g0082.Args
              y_len := y_support.Length()
              for i_it := 0; i_it < y_len; i_it++ { 
                y = y_support.At(i_it)
                var loop_2 EID 
                _ = loop_2
                if (y.Isa.IsIn(Language.C_Vardef) == CTRUE) { 
                  { var g0083 *Language.Vardef   = Language.To_Vardef(y)
                    lvar = lvar.AddFast(g0083.Id())
                    loop_2 = EID{lvar.Id(),0}
                    } 
                  } else {
                  var try_3 EID 
                  { var arg_4 *ClaireVariable  
                    var try_5 EID 
                    { var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
                      { 
                        var va_arg1 *ClaireVariable  
                        var va_arg2 *ClaireSymbol  
                        va_arg1 = _CL_obj
                        var try_6 EID 
                        try_6 = Language.F_extract_symbol_any(y)
                        if ErrorIn(try_6) {try_5 = try_6
                        } else {
                        va_arg2 = ToSymbol(OBJ(try_6))
                        va_arg1.Pname = va_arg2
                        try_5 = EID{va_arg2.Id(),0}
                        }
                        } 
                      if !ErrorIn(try_5) {
                      _CL_obj.Range = ToType(C_any.Id())
                      try_5 = EID{_CL_obj.Id(),0}
                      }
                      } 
                    if ErrorIn(try_5) {try_3 = try_5
                    } else {
                    arg_4 = To_Variable(OBJ(try_5))
                    try_3 = EID{lvar.AddFast(arg_4.Id()).Id(),0}
                    }
                    } 
                  if ErrorIn(try_3) {loop_2 = try_3
                  } else {
                  lvar = ToList(OBJ(try_3))
                  loop_2 = EID{lvar.Id(),0}
                  }
                  } 
                if ErrorIn(loop_2) {Result = loop_2
                break
                } else {
                }
                } 
              } 
            } 
          }  else if (l.Isa.IsIn(C_list) == CTRUE) { 
          { var g0085 *ClaireList   = ToList(l)
            { 
              var y *ClaireAny  
              _ = y
              Result= EID{CFALSE.Id(),0}
              var y_support *ClaireList  
              y_support = g0085
              y_len := y_support.Length()
              for i_it := 0; i_it < y_len; i_it++ { 
                y = y_support.At(i_it)
                var loop_7 EID 
                _ = loop_7
                if (y.Isa.IsIn(Language.C_Vardef) == CTRUE) { 
                  { var g0086 *Language.Vardef   = Language.To_Vardef(y)
                    lvar = lvar.AddFast(g0086.Id())
                    loop_7 = EID{lvar.Id(),0}
                    } 
                  } else {
                  var try_8 EID 
                  { var arg_9 *ClaireVariable  
                    var try_10 EID 
                    { var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
                      { 
                        var va_arg1 *ClaireVariable  
                        var va_arg2 *ClaireSymbol  
                        va_arg1 = _CL_obj
                        var try_11 EID 
                        try_11 = Language.F_extract_symbol_any(y)
                        if ErrorIn(try_11) {try_10 = try_11
                        } else {
                        va_arg2 = ToSymbol(OBJ(try_11))
                        va_arg1.Pname = va_arg2
                        try_10 = EID{va_arg2.Id(),0}
                        }
                        } 
                      if !ErrorIn(try_10) {
                      _CL_obj.Range = ToType(C_any.Id())
                      try_10 = EID{_CL_obj.Id(),0}
                      }
                      } 
                    if ErrorIn(try_10) {try_8 = try_10
                    } else {
                    arg_9 = To_Variable(OBJ(try_10))
                    try_8 = EID{lvar.AddFast(arg_9.Id()).Id(),0}
                    }
                    } 
                  if ErrorIn(try_8) {loop_7 = try_8
                  } else {
                  lvar = ToList(OBJ(try_8))
                  loop_7 = EID{lvar.Id(),0}
                  }
                  } 
                if ErrorIn(loop_7) {Result = loop_7
                break
                } else {
                }
                } 
              } 
            } 
          } else {
          var try_12 EID 
          { var arg_13 *ClaireVariable  
            var try_14 EID 
            { var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
              { 
                var va_arg1 *ClaireVariable  
                var va_arg2 *ClaireSymbol  
                va_arg1 = _CL_obj
                var try_15 EID 
                try_15 = Language.F_extract_symbol_any(l)
                if ErrorIn(try_15) {try_14 = try_15
                } else {
                va_arg2 = ToSymbol(OBJ(try_15))
                va_arg1.Pname = va_arg2
                try_14 = EID{va_arg2.Id(),0}
                }
                } 
              if !ErrorIn(try_14) {
              _CL_obj.Range = ToType(C_any.Id())
              try_14 = EID{_CL_obj.Id(),0}
              }
              } 
            if ErrorIn(try_14) {try_12 = try_14
            } else {
            arg_13 = To_Variable(OBJ(try_14))
            try_12 = EID{lvar.AddFast(arg_13.Id()).Id(),0}
            }
            } 
          if ErrorIn(try_12) {Result = try_12
          } else {
          lvar = ToList(OBJ(try_12))
          Result = EID{lvar.Id(),0}
          }
          } 
        if !ErrorIn(Result) {
        { var arg_16 *ClaireAny  
          var try_17 EID 
          if ANY(Core.F_CALL(C_length,ARGS(lbody.ToEID()))).IsInt(1) { 
            try_17 = Core.F_CALL(C_nth,ARGS(lbody.ToEID(),EID{C__INT,IVAL(1)}))
            } else {
            { var _CL_obj *Language.Do   = Language.To_Do(new(Language.Do).Is(Language.C_Do))
              _CL_obj.Args = ToList(lbody)
              try_17 = EID{_CL_obj.Id(),0}
              } 
            } 
          if ErrorIn(try_17) {Result = try_17
          } else {
          arg_16 = ANY(try_17)
          Result = Language.F_lambda_I_list(lvar,arg_16)
          }
          } 
        }
        } 
      }
      } 
    return Result} 
  
// The EID go function for: readlambda @ meta_reader (throw: true) 
func E_Reader_readlambda_meta_reader (r EID,l EID) EID { 
    return ToMetaReader(OBJ(r)).Readlambda(ANY(l) )} 
  
/* {1} The go function for: Do!(x:any,y:any) [status=1] */
func F_Do_I_any (x *ClaireAny ,y *ClaireAny ) EID { 
    // eid body s = any
    var Result EID 
    if (y.Isa.IsIn(Language.C_Do) == CTRUE) { 
      { var g0089 *Language.Do   = Language.To_Do(y)
        { 
          var va_arg1 *Language.Do  
          var va_arg2 *ClaireList  
          va_arg1 = g0089
          var try_1 EID 
          try_1 = g0089.Args.Nth_plus(1,x)
          if ErrorIn(try_1) {Result = try_1
          } else {
          va_arg2 = ToList(OBJ(try_1))
          va_arg1.Args = va_arg2
          Result = EID{va_arg2.Id(),0}
          }
          } 
        if !ErrorIn(Result) {
        Result = EID{g0089.Id(),0}
        }
        } 
      } else {
      { var _CL_obj *Language.Do   = Language.To_Do(new(Language.Do).Is(Language.C_Do))
        _CL_obj.Args = MakeList(ToType(C_any.Id()),x,y)
        Result = EID{_CL_obj.Id(),0}
        } 
      } 
    return Result} 
  
// The EID go function for: Do! @ any (throw: true) 
func E_Do_I_any (x EID,y EID) EID { 
    return F_Do_I_any(ANY(x),ANY(y) )} 
  
// extract the type from a list<X> expression
/* {1} The go function for: extract_of_type(x:Call) [status=1] */
func F_extract_of_type_Call (x *Language.Call ) EID { 
    // eid body s = type
    var Result EID 
    { var l *ClaireList   = x.Args
      if (l.Length() > 2) { 
        { var y *ClaireAny   = l.At(3-1)
          if (y.Isa.IsIn(Language.C_List) == CTRUE) { 
            { var g0091 *Language.List   = Language.To_List(y)
              { var z *ClaireAny   = g0091.Args.At(1-1)
                if (z.Isa.IsIn(Language.C_Set) == CTRUE) { 
                  { var g0092 *Language.Set   = Language.To_Set(z)
                    Result = Language.F_extract_type_any(g0092.Args.At(1-1))
                    } 
                  } else {
                  Result = EID{C_any.Id(),0}
                  } 
                } 
              } 
            } else {
            Result = EID{C_any.Id(),0}
            } 
          } 
        } else {
        Result = EID{C_any.Id(),0}
        } 
      } 
    return Result} 
  
// The EID go function for: extract_of_type @ Call (throw: true) 
func E_extract_of_type_Call (x EID) EID { 
    return F_extract_of_type_Call(Language.To_Call(OBJ(x)) )} 
  
// **********************************************************************
// *   Part 3: read functional calls                                    *
// **********************************************************************
// store the line number in debug mode
// in v4.0 we will not do this for JITO calls :)
// this is a cool trick when operating in debug mode: we store the last evaluated
// call so we can tell very simply which last call triggered the error
//
/* {1} The go function for: DBregister(c:Call) [status=1] */
func F_DBregister_Call (c *Language.Call ) EID { 
    // eid body s = Call
    var Result EID 
    if (ClEnv.Debug_I >= 0) { 
      Language.C_iClaire_LastCall.Value = c.Id()
      Core.F_put_table(C_Reader_DBline,c.Id(),MakeInteger(ClEnv.NLine).Id())
      } 
    if ((c.Selector.Id() == C_store.Id()) && 
        (c.Args.Length() == 1)) { 
      { var l *ClaireList   = c.Args
        if (l.At(1-1).Isa.IsIn(Core.C_global_variable) == CTRUE) { 
          { var arg_1 *ClaireAny  
            var try_2 EID 
            try_2 = Core.F_CALL(C_make_string,ARGS(Core.F_CALL(C_name,ARGS(l.At(1-1).ToEID()))))
            if ErrorIn(try_2) {Result = try_2
            } else {
            arg_1 = ANY(try_2)
            Result = ToArray(l.Id()).NthPut(1,arg_1).ToEID()
            }
            } 
          } else {
          Result = EID{CFALSE.Id(),0}
          } 
        } 
      } else {
      Result = EID{CFALSE.Id(),0}
      } 
    if !ErrorIn(Result) {
    Result = EID{c.Id(),0}
    }
    return Result} 
  
// The EID go function for: DBregister @ Call (throw: true) 
func E_DBregister_Call (c EID) EID { 
    return F_DBregister_Call(Language.To_Call(OBJ(c)) )} 
  
/* {1} The go function for: Call!(p:property,l:list) [status=1] */
func F_Call_I_property (p *ClaireProperty ,l *ClaireList ) EID { 
    // eid body s = Call
    var Result EID 
    { var arg_1 *Language.Call  
      { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
        _CL_obj.Selector = p
        _CL_obj.Args = l
        arg_1 = _CL_obj
        } 
      Result = F_DBregister_Call(arg_1)
      } 
    return Result} 
  
// The EID go function for: Call! @ property (throw: true) 
func E_Call_I_property (p EID,l EID) EID { 
    return F_Call_I_property(ToProperty(OBJ(p)),ToList(OBJ(l)) )} 
  
// if the expression is a call -------------------------------------------
// x is the first token that we have read x(...)
// there are many special case (x is not a propery) then the regular case
// t is a type when x was read as (p@t)
/* {1} The go function for: readcall(r:meta_reader,x:any,t:any) [status=1] */
func (r *MetaReader ) Readcall (x *ClaireAny ,t *ClaireAny ) EID { 
    // eid body s = any
    var Result EID 
    { var l *ClaireAny  
      var try_1 EID 
      try_1 = r.Cnext().Nextseq(41)
      if ErrorIn(try_1) {Result = try_1
      } else {
      l = ANY(try_1)
      if (x == C_printf.Id()) { 
        { var _CL_obj *Language.Printf   = Language.To_Printf(new(Language.Printf).Is(Language.C_Printf))
          _CL_obj.Args = ToList(l)
          Result = EID{_CL_obj.Id(),0}
          } 
        }  else if (x == C_error.Id()) { 
        { var _CL_obj *Language.Error   = Language.To_Error(new(Language.Error).Is(Language.C_Error))
          _CL_obj.Args = ToList(l)
          Result = EID{_CL_obj.Id(),0}
          } 
        }  else if (x == C_assert.Id()) { 
        { var _CL_obj *Language.Assert   = Language.To_Assert(new(Language.Assert).Is(Language.C_Assert))
          _CL_obj.Args = ToList(l)
          _CL_obj.Index = ClEnv.NLine
          _CL_obj.External = r.External
          Result = EID{_CL_obj.Id(),0}
          } 
        }  else if (x == C_trace.Id()) { 
        { var _CL_obj *Language.Trace   = Language.To_Trace(new(Language.Trace).Is(Language.C_Trace))
          _CL_obj.Args = ToList(l)
          Result = EID{_CL_obj.Id(),0}
          } 
        }  else if (x == C_branch.Id()) { 
        { var _CL_obj *Language.Branch   = Language.To_Branch(new(Language.Branch).Is(Language.C_Branch))
          _CL_obj.Args = ToList(l)
          Result = EID{_CL_obj.Id(),0}
          } 
        }  else if (x == C_quote.Id()) { 
        { var arg_2 *ClaireAny  
          var try_3 EID 
          if (F_boolean_I_any(l) == CTRUE) { 
            try_3 = Core.F_CALL(C_nth,ARGS(l.ToEID(),EID{C__INT,IVAL(1)}))
            } else {
            try_3 = EID{CFALSE.Id(),0}
            } 
          if ErrorIn(try_3) {Result = try_3
          } else {
          arg_2 = ANY(try_3)
          Result = Language.C_Quote.Make(arg_2).ToEID()
          }
          } 
        }  else if (x == C_tuple.Id()) { 
        { var _CL_obj *Language.Tuple   = Language.To_Tuple(new(Language.Tuple).Is(Language.C_Tuple))
          _CL_obj.Args = ToList(l)
          Result = EID{_CL_obj.Id(),0}
          } 
        }  else if (x == C_list.Id()) { 
        { var _CL_obj *Language.List   = Language.To_List(new(Language.List).Is(Language.C_List))
          _CL_obj.Args = ToList(l)
          Result = EID{_CL_obj.Id(),0}
          } 
        } else {
        var g0103I *ClaireBoolean  
        if (x.Isa.IsIn(Language.C_Call) == CTRUE) { 
          { var g0095 *Language.Call   = Language.To_Call(x)
            g0103I = Equal(g0095.Args.At(1-1),C_list.Id())
            } 
          } else {
          g0103I = CFALSE
          } 
        if (g0103I == CTRUE) { 
          { var _CL_obj *Language.List   = Language.To_List(new(Language.List).Is(Language.C_List))
            { 
              var va_arg1 *Language.List  
              var va_arg2 *ClaireType  
              va_arg1 = _CL_obj
              var try_4 EID 
              try_4 = F_extract_of_type_Call(Language.To_Call(x))
              if ErrorIn(try_4) {Result = try_4
              } else {
              va_arg2 = ToType(OBJ(try_4))
              va_arg1.Of = va_arg2
              Result = EID{va_arg2.Id(),0}
              }
              } 
            if !ErrorIn(Result) {
            _CL_obj.Args = ToList(l)
            Result = EID{_CL_obj.Id(),0}
            }
            } 
          } else {
          var g0104I *ClaireBoolean  
          if (x.Isa.IsIn(Language.C_Call) == CTRUE) { 
            { var g0096 *Language.Call   = Language.To_Call(x)
              g0104I = Equal(g0096.Args.At(1-1),C_array.Id())
              } 
            } else {
            g0104I = CFALSE
            } 
          if (g0104I == CTRUE) { 
            { var _CL_obj *Language.Array   = Language.To_Array(new(Language.Array).Is(Language.C_Array))
              { 
                var va_arg1 *Language.Array  
                var va_arg2 *ClaireType  
                va_arg1 = _CL_obj
                var try_5 EID 
                try_5 = F_extract_of_type_Call(Language.To_Call(x))
                if ErrorIn(try_5) {Result = try_5
                } else {
                va_arg2 = ToType(OBJ(try_5))
                va_arg1.Of = va_arg2
                Result = EID{va_arg2.Id(),0}
                }
                } 
              if !ErrorIn(Result) {
              _CL_obj.Args = ToList(l)
              Result = EID{_CL_obj.Id(),0}
              }
              } 
            } else {
            var g0105I *ClaireBoolean  
            if (x.Isa.IsIn(Language.C_Call) == CTRUE) { 
              { var g0097 *Language.Call   = Language.To_Call(x)
                g0105I = Equal(g0097.Args.At(1-1),C_set.Id())
                } 
              } else {
              g0105I = CFALSE
              } 
            if (g0105I == CTRUE) { 
              { var _CL_obj *Language.Set   = Language.To_Set(new(Language.Set).Is(Language.C_Set))
                { 
                  var va_arg1 *Language.Set  
                  var va_arg2 *ClaireType  
                  va_arg1 = _CL_obj
                  var try_6 EID 
                  try_6 = F_extract_of_type_Call(Language.To_Call(x))
                  if ErrorIn(try_6) {Result = try_6
                  } else {
                  va_arg2 = ToType(OBJ(try_6))
                  va_arg1.Of = va_arg2
                  Result = EID{va_arg2.Id(),0}
                  }
                  } 
                if !ErrorIn(Result) {
                _CL_obj.Args = ToList(l)
                Result = EID{_CL_obj.Id(),0}
                }
                } 
              }  else if ((C_class.Id() == x.Isa.Id()) && 
                (ToClass(x).IsIn(Language.C_Macro) == CTRUE)) { 
              { var o *ClaireObject   = F_new_object_class(ToClass(x))
                Result = Core.F_put_property2(C_args,o,l)
                if !ErrorIn(Result) {
                Result = EID{o.Id(),0}
                }
                } 
              }  else if (x == C_set.Id()) { 
              { var _CL_obj *Language.Set   = Language.To_Set(new(Language.Set).Is(Language.C_Set))
                _CL_obj.Args = ToList(l)
                Result = EID{_CL_obj.Id(),0}
                } 
              }  else if ((x == C_return.Id()) || 
                (x == C_break.Id())) { 
              { var arg_7 *ClaireAny  
                var try_8 EID 
                if (F_boolean_I_any(l) == CTRUE) { 
                  try_8 = Core.F_CALL(C_nth,ARGS(l.ToEID(),EID{C__INT,IVAL(1)}))
                  } else {
                  try_8 = EID{CTRUE.Id(),0}
                  } 
                if ErrorIn(try_8) {Result = try_8
                } else {
                arg_7 = ANY(try_8)
                Result = Language.C_Return.Make(arg_7).ToEID()
                }
                } 
              }  else if (C_class.Id() == x.Isa.Id()) { 
              var g0106I *ClaireBoolean  
              var try_9 EID 
              { var arg_10 *ClaireBoolean  
                var try_11 EID 
                { var arg_12 *ClaireAny  
                  var try_13 EID 
                  { 
                    var y *ClaireAny  
                    _ = y
                    try_13= EID{CFALSE.Id(),0}
                    var y_support *ClaireList  
                    var try_14 EID 
                    try_14 = Core.F_enumerate_any(l)
                    if ErrorIn(try_14) {try_13 = try_14
                    } else {
                    y_support = ToList(OBJ(try_14))
                    y_len := y_support.Length()
                    for i_it := 0; i_it < y_len; i_it++ { 
                      y = y_support.At(i_it)
                      var loop_15 EID 
                      _ = loop_15
                      var g0107I *ClaireBoolean  
                      var try_16 EID 
                      { var arg_17 *ClaireBoolean  
                        var try_18 EID 
                        if (y.Isa.IsIn(Language.C_Call) == CTRUE) { 
                          { var g0098 *Language.Call   = Language.To_Call(y)
                            if (g0098.Selector.Id() == C__equal.Id()) { 
                              { var arg_19 *ClaireProperty  
                                var try_20 EID 
                                try_20 = Language.F_make_a_property_any(g0098.Args.At(1-1))
                                if ErrorIn(try_20) {try_18 = try_20
                                } else {
                                arg_19 = ToProperty(OBJ(try_20))
                                try_18 = ToArray(g0098.Args.Id()).NthPut(1,arg_19.Id()).ToEID()
                                }
                                } 
                              if !ErrorIn(try_18) {
                              try_18 = EID{CTRUE.Id(),0}
                              }
                              } else {
                              try_18 = EID{CFALSE.Id(),0}
                              } 
                            } 
                          } else {
                          try_18 = EID{CFALSE.Id(),0}
                          } 
                        if ErrorIn(try_18) {try_16 = try_18
                        } else {
                        arg_17 = ToBoolean(OBJ(try_18))
                        try_16 = EID{arg_17.Not.Id(),0}
                        }
                        } 
                      if ErrorIn(try_16) {loop_15 = try_16
                      } else {
                      g0107I = ToBoolean(OBJ(try_16))
                      if (g0107I == CTRUE) { 
                        try_13 = EID{CTRUE.Id(),0}
                        break
                        } else {
                        loop_15 = EID{CFALSE.Id(),0}
                        } 
                      }
                      if ErrorIn(loop_15) {try_13 = loop_15
                      break
                      } else {
                      }}
                      } 
                    } 
                  if ErrorIn(try_13) {try_11 = try_13
                  } else {
                  arg_12 = ANY(try_13)
                  try_11 = EID{Core.F_not_any(arg_12).Id(),0}
                  }
                  } 
                if ErrorIn(try_11) {try_9 = try_11
                } else {
                arg_10 = ToBoolean(OBJ(try_11))
                try_9 = EID{arg_10.Not.Id(),0}
                }
                } 
              if ErrorIn(try_9) {Result = try_9
              } else {
              g0106I = ToBoolean(OBJ(try_9))
              if (g0106I == CTRUE) { 
                { var l2 *ClaireList   = ToClass(x).Params
                  { var n int 
                    if (l2.Isa.IsIn(C_list) == CTRUE) { 
                      n = l2.Length()
                      } else {
                      n = 0
                      } 
                    if ANY(Core.F_CALL(C_length,ARGS(l.ToEID()))).IsInt(n) { 
                      var try_21 EID 
                      { var i_bag *ClaireList   = ToType(CEMPTY.Id()).EmptyList()
                        { var i int  = 1
                          { var g0101 int  = n
                            try_21= EID{CFALSE.Id(),0}
                            for (i <= g0101) { 
                              var loop_22 EID 
                              _ = loop_22
                              { 
                              { var arg_23 *Language.Call  
                                var try_24 EID 
                                { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                  _CL_obj.Selector = ToProperty(C__equal.Id())
                                  { 
                                    var va_arg1 *Language.Call  
                                    var va_arg2 *ClaireList  
                                    va_arg1 = _CL_obj
                                    var try_25 EID 
                                    { 
                                      var v_bag_arg *ClaireAny  
                                      try_25= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                                      ToList(OBJ(try_25)).AddFast(l2.At(i-1))
                                      var try_26 EID 
                                      try_26 = Core.F_CALL(C_nth,ARGS(l.ToEID(),EID{C__INT,IVAL(i)}))
                                      if ErrorIn(try_26) {try_25 = try_26
                                      } else {
                                      v_bag_arg = ANY(try_26)
                                      ToList(OBJ(try_25)).AddFast(v_bag_arg)}
                                      } 
                                    if ErrorIn(try_25) {try_24 = try_25
                                    } else {
                                    va_arg2 = ToList(OBJ(try_25))
                                    va_arg1.Args = va_arg2
                                    try_24 = EID{va_arg2.Id(),0}
                                    }
                                    } 
                                  if !ErrorIn(try_24) {
                                  try_24 = EID{_CL_obj.Id(),0}
                                  }
                                  } 
                                if ErrorIn(try_24) {loop_22 = try_24
                                } else {
                                arg_23 = Language.To_Call(OBJ(try_24))
                                loop_22 = EID{i_bag.AddFast(arg_23.Id()).Id(),0}
                                }
                                } 
                              if ErrorIn(loop_22) {try_21 = loop_22
                              break
                              } else {
                              i = (i+1)
                              }
                              } 
                            }
                            } 
                          } 
                        if !ErrorIn(try_21) {
                        try_21 = EID{i_bag.Id(),0}
                        }
                        } 
                      if ErrorIn(try_21) {Result = try_21
                      } else {
                      l = ANY(try_21)
                      Result = l.ToEID()
                      }
                      } else {
                      Result = F_Serror_string(MakeString("[174] Wrong instantiation list ~S(~S..."),MakeConstantList(x,MakeConstantList(l).Id()))
                      } 
                    } 
                  } 
                } else {
                Result = EID{CFALSE.Id(),0}
                } 
              }
              if !ErrorIn(Result) {
              { var _CL_obj *Language.Definition   = Language.To_Definition(new(Language.Definition).Is(Language.C_Definition))
                _CL_obj.Arg = ToClass(x)
                _CL_obj.Args = ToList(l)
                Result = EID{_CL_obj.Id(),0}
                } 
              }
              } else {
              var g0108I *ClaireBoolean  
              { 
                var v_or7 *ClaireBoolean  
                
                v_or7 = x.Isa.IsIn(C_Variable)
                if (v_or7 == CTRUE) {g0108I = CTRUE
                } else { 
                  { var arg_27 *ClaireObject  
                    if (x.Isa.IsIn(Core.C_global_variable) == CTRUE) { 
                      { var g0102 *Core.GlobalVariable   = Core.ToGlobalVariable(x)
                        arg_27 = ToObject(g0102.Range.Id())
                        } 
                      } else {
                      arg_27 = ToObject(CFALSE.Id())
                      } 
                    v_or7 = F_boolean_I_any(arg_27.Id())
                    } 
                  if (v_or7 == CTRUE) {g0108I = CTRUE
                  } else { 
                    g0108I = CFALSE} 
                  } 
                } 
              if (g0108I == CTRUE) { 
                { var arg_28 *ClaireList  
                  { var arg_29 *ClaireAny  
                    if (F_boolean_I_any(l) == CTRUE) { 
                      arg_29 = l
                      } else {
                      arg_29 = MakeConstantList(ClEnv.Id()).Id()
                      } 
                    arg_28 = F_cons_any(x,ToList(arg_29))
                    } 
                  Result = F_Call_I_property(Core.C_call,arg_28)
                  } 
                } else {
                { var p *ClaireProperty  
                  var try_30 EID 
                  try_30 = Language.F_make_a_property_any(x)
                  if ErrorIn(try_30) {Result = try_30
                  } else {
                  p = ToProperty(OBJ(try_30))
                  { var l2 *ClaireAny  
                    if (F_boolean_I_any(l) == CTRUE) { 
                      l2 = l
                      } else {
                      l2 = MakeConstantList(ClEnv.Id()).Id()
                      } 
                    if (t != CNULL) { 
                      { var _CL_obj *Language.Super   = Language.To_Super(new(Language.Super).Is(Language.C_Super))
                        _CL_obj.Selector = p
                        _CL_obj.CastTo = ToType(t)
                        _CL_obj.Args = ToList(l2)
                        Result = EID{_CL_obj.Id(),0}
                        } 
                      } else {
                      Result = F_Call_I_property(p,ToList(l2))
                      } 
                    } 
                  }
                  } 
                } 
              } 
            } 
          } 
        } 
      }
      } 
    return Result} 
  
// The EID go function for: readcall @ meta_reader (throw: true) 
func E_readcall_meta_reader (r EID,x EID,t EID) EID { 
    return ToMetaReader(OBJ(r)).Readcall(ANY(x),ANY(t) )} 
  
// **********************************************************************
// *   Part 4: read definitions                                         *
// **********************************************************************
// reads a definition (CLAIRE2 syntax)   - x and y are two expressions that have been read
//
/* {1} The go function for: nextdefinition(r:meta_reader,x:any,y:any,old?:boolean) [status=1] */
func (r *MetaReader ) Nextdefinition (x *ClaireAny ,y *ClaireAny ,old_ask *ClaireBoolean ) EID { 
    // eid body s = any
    var Result EID 
    r.LastArrow = CFALSE
    if (Equal(y,C_triangle.Value) == CTRUE) { 
      Result = r.Cnext().NextDefclass(x,old_ask)
      }  else if (y == C_L_.Id()) { 
      { var table_ask *ClaireBoolean  
        if (x.Isa.IsIn(Language.C_Call) == CTRUE) { 
          { var g0109 *Language.Call   = Language.To_Call(x)
            table_ask = MakeBoolean((g0109.Selector.Id() == C_nth.Id()) && ((g0109.Args.At(1-1).Isa.IsIn(C_unbound_symbol) == CTRUE) || 
                (C_table.Id() == g0109.Args.At(1-1).Isa.Id())))
            } 
          } else {
          table_ask = CFALSE
          } 
        { var z *ClaireAny  
          var try_1 EID 
          try_1 = r.Nexte()
          if ErrorIn(try_1) {Result = try_1
          } else {
          z = ANY(try_1)
          { var w *ClaireAny  
            var try_2 EID 
            try_2 = r.Nexte()
            if ErrorIn(try_2) {Result = try_2
            } else {
            w = ANY(try_2)
            var g0112I *ClaireBoolean  
            if (table_ask == CTRUE) { 
              g0112I = Equal(w,C_L__equal.Id())
              } else {
              g0112I = MakeBoolean((Equal(w,C_arrow.Value) == CTRUE) || (w == C__equal_sup.Id()))
              } 
            if (g0112I == CTRUE) { 
              Result = EID{CNIL.Id(),0}
              } else {
              Result = F_Serror_string(MakeString("[149] wrong keyword (~S) after ~S"),MakeConstantList(w,z))
              } 
            if !ErrorIn(Result) {
            Result = r.Nextmethod(x,
              z,
              table_ask,
              old_ask,
              Equal(w,C__equal_sup.Id()))
            }
            }
            } 
          }
          } 
        } 
      }  else if (y == C_L_L_.Id()) { 
      if (x.Isa.IsIn(Language.C_Call) == CTRUE) { 
        { var g0110 *Language.Call   = Language.To_Call(x)
          { var ru *ClaireAny  
            _ = ru
            var try_3 EID 
            try_3 = r.Nexte()
            if ErrorIn(try_3) {Result = try_3
            } else {
            ru = ANY(try_3)
            { var z *ClaireAny  
              var try_4 EID 
              try_4 = r.Nexts(C__equal_sup)
              if ErrorIn(try_4) {Result = try_4
              } else {
              z = ANY(try_4)
              { var _CL_obj *Language.Defrule   = Language.To_Defrule(new(Language.Defrule).Is(Language.C_Defrule))
                _CL_obj.Ident = g0110.Selector.Name
                _CL_obj.Args = g0110.Args
                _CL_obj.Arg = z
                { 
                  var va_arg1 *Language.Defrule  
                  var va_arg2 *ClaireAny  
                  va_arg1 = _CL_obj
                  var try_5 EID 
                  if (r.Firstc() == 41) { 
                    r.Next()
                    try_5 = EID{CNIL.Id(),0}
                    } else {
                    { var arg_6 *ClaireAny  
                      var try_7 EID 
                      try_7 = r.Nexte()
                      if ErrorIn(try_7) {try_5 = try_7
                      } else {
                      arg_6 = ANY(try_7)
                      try_5 = r.Readblock(arg_6,41)
                      }
                      } 
                    } 
                  if ErrorIn(try_5) {Result = try_5
                  } else {
                  va_arg2 = ANY(try_5)
                  va_arg1.Body = va_arg2
                  Result = va_arg2.ToEID()
                  }
                  } 
                if !ErrorIn(Result) {
                Result = EID{_CL_obj.Id(),0}
                }
                } 
              }
              } 
            }
            } 
          } 
        } else {
        Result = r.Nextinst(x)
        } 
      }  else if ((Equal(y,C_arrow.Value) == CTRUE) || 
        (y == C__equal_sup.Id())) { 
      r.LastArrow = Equal(y,C__equal_sup.Id())
      
      Result = r.Nextmethod(x,
        C_void.Id(),
        CFALSE,
        old_ask,
        Equal(y,C__equal_sup.Id()))
      }  else if ((y == C_L__equal.Id()) && 
        (x.Isa.IsIn(Language.C_Vardef) == CTRUE)) { 
      { var _CL_obj *Language.Defobj   = Language.To_Defobj(new(Language.Defobj).Is(Language.C_Defobj))
        _CL_obj.Ident = ToSymbol(OBJ(Core.F_CALL(C_mClaire_pname,ARGS(x.ToEID()))))
        _CL_obj.Arg = Core.C_global_variable
        { 
          var va_arg1 *Language.Definition  
          var va_arg2 *ClaireList  
          va_arg1 = Language.To_Definition(_CL_obj.Id())
          var try_8 EID 
          { 
            var v_bag_arg *ClaireAny  
            try_8= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
            var try_9 EID 
            { var arg_10 *ClaireList  
              var try_11 EID 
              { 
                var v_bag_arg *ClaireAny  
                try_11= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                ToList(OBJ(try_11)).AddFast(C_range.Id())
                var try_12 EID 
                try_12 = Language.F_extract_type_any(ANY(Core.F_CALL(C_range,ARGS(x.ToEID()))))
                if ErrorIn(try_12) {try_11 = try_12
                } else {
                v_bag_arg = ANY(try_12)
                ToList(OBJ(try_11)).AddFast(v_bag_arg)}
                } 
              if ErrorIn(try_11) {try_9 = try_11
              } else {
              arg_10 = ToList(OBJ(try_11))
              try_9 = F_Call_I_property(ToProperty(C__equal.Id()),arg_10)
              }
              } 
            if ErrorIn(try_9) {try_8 = try_9
            } else {
            v_bag_arg = ANY(try_9)
            ToList(OBJ(try_8)).AddFast(v_bag_arg)
            var try_13 EID 
            { var arg_14 *ClaireList  
              var try_15 EID 
              { 
                var v_bag_arg *ClaireAny  
                try_15= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                ToList(OBJ(try_15)).AddFast(C_value.Id())
                var try_16 EID 
                try_16 = r.Nexte()
                if ErrorIn(try_16) {try_15 = try_16
                } else {
                v_bag_arg = ANY(try_16)
                ToList(OBJ(try_15)).AddFast(v_bag_arg)}
                } 
              if ErrorIn(try_15) {try_13 = try_15
              } else {
              arg_14 = ToList(OBJ(try_15))
              try_13 = F_Call_I_property(ToProperty(C__equal.Id()),arg_14)
              }
              } 
            if ErrorIn(try_13) {try_8 = try_13
            } else {
            v_bag_arg = ANY(try_13)
            ToList(OBJ(try_8)).AddFast(v_bag_arg)}}
            } 
          if ErrorIn(try_8) {Result = try_8
          } else {
          va_arg2 = ToList(OBJ(try_8))
          va_arg1.Args = va_arg2
          Result = EID{va_arg2.Id(),0}
          }
          } 
        if !ErrorIn(Result) {
        Result = EID{_CL_obj.Id(),0}
        }
        } 
      } else {
      { var _CL_obj *Language.Do   = Language.To_Do(new(Language.Do).Is(Language.C_Do))
        _CL_obj.Args = MakeList(ToType(C_any.Id()),x,y)
        Result = EID{_CL_obj.Id(),0}
        } 
      } 
    return Result} 
  
// The EID go function for: nextdefinition @ meta_reader (throw: true) 
func E_nextdefinition_meta_reader (r EID,x EID,y EID,old_ask EID) EID { 
    return ToMetaReader(OBJ(r)).Nextdefinition(ANY(x),
      ANY(y),
      ToBoolean(OBJ(old_ask)) )} 
  
/* {1} The go function for: nextmethod(r:meta_reader,x:any,y:any,table?:boolean,old?:boolean,inl?:boolean) [status=1] */
func (r *MetaReader ) Nextmethod (x *ClaireAny ,y *ClaireAny ,table_ask *ClaireBoolean ,old_ask *ClaireBoolean ,inl_ask *ClaireBoolean ) EID { 
    // eid body s = any
    var Result EID 
    { var n int  = r.Skipc()
      { var z *ClaireAny  
        var try_1 EID 
        if (old_ask == CTRUE) { 
          { var arg_2 *ClaireAny  
            var try_3 EID 
            try_3 = r.Nexte()
            if ErrorIn(try_3) {try_1 = try_3
            } else {
            arg_2 = ANY(try_3)
            try_1 = r.Readblock(arg_2,93)
            }
            } 
          }  else if (n == 40) { 
          if (r.Toplevel == CTRUE) { 
            try_1 = r.Nexts(C_none)
            } else {
            { var arg_4 *ClaireAny  
              var try_5 EID 
              try_5 = r.Cnext().Nexte()
              if ErrorIn(try_5) {try_1 = try_5
              } else {
              arg_4 = ANY(try_5)
              try_1 = r.Readblock(arg_4,41)
              }
              } 
            } 
          } else {
          try_1 = r.Nexte()
          } 
        if ErrorIn(try_1) {Result = try_1
        } else {
        z = ANY(try_1)
        { var rs *Language.Defmethod  
          var try_6 EID 
          { var _CL_obj *Language.Defmethod   = Language.To_Defmethod(new(Language.Defmethod).Is(Language.C_Defmethod))
            _CL_obj.Arg = Language.To_Call(x)
            _CL_obj.SetArg = y
            { 
              var va_arg1 *Language.Defmethod  
              var va_arg2 *ClaireAny  
              va_arg1 = _CL_obj
              var try_7 EID 
              if (z == C_let.Id()) { 
                try_7 = r.Readlet(C_None)
                } else {
                try_7 = z.ToEID()
                } 
              if ErrorIn(try_7) {try_6 = try_7
              } else {
              va_arg2 = ANY(try_7)
              va_arg1.Body = va_arg2
              try_6 = va_arg2.ToEID()
              }
              } 
            if !ErrorIn(try_6) {
            _CL_obj.Inline_ask = inl_ask
            try_6 = EID{_CL_obj.Id(),0}
            }
            } 
          if ErrorIn(try_6) {Result = try_6
          } else {
          rs = Language.To_Defmethod(OBJ(try_6))
          if (table_ask == CTRUE) { 
            rs.Isa = Language.C_Defarray
            } 
          Result = EID{rs.Id(),0}
          }
          } 
        }
        } 
      } 
    return Result} 
  
// The EID go function for: nextmethod @ meta_reader (throw: true) 
func E_nextmethod_meta_reader (r EID,x EID,y EID,table_ask EID,old_ask EID,inl_ask EID) EID { 
    return ToMetaReader(OBJ(r)).Nextmethod(ANY(x),
      ANY(y),
      ToBoolean(OBJ(table_ask)),
      ToBoolean(OBJ(old_ask)),
      ToBoolean(OBJ(inl_ask)) )} 
  
// reads an instantiation
//
/* {1} The go function for: nextinst(r:meta_reader,x:any) [status=1] */
func (r *MetaReader ) Nextinst (x *ClaireAny ) EID { 
    // eid body s = any
    var Result EID 
    if (x.Isa.IsIn(C_Variable) == CTRUE) { 
      { var g0113 *ClaireVariable   = To_Variable(x)
        { var _CL_obj *Language.Defobj   = Language.To_Defobj(new(Language.Defobj).Is(Language.C_Defobj))
          _CL_obj.Ident = g0113.Pname
          _CL_obj.Arg = Core.C_global_variable
          { 
            var va_arg1 *Language.Definition  
            var va_arg2 *ClaireList  
            va_arg1 = Language.To_Definition(_CL_obj.Id())
            var try_1 EID 
            { 
              var v_bag_arg *ClaireAny  
              try_1= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
              var try_2 EID 
              { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                _CL_obj.Selector = ToProperty(C__equal.Id())
                { 
                  var va_arg1 *Language.Call  
                  var va_arg2 *ClaireList  
                  va_arg1 = _CL_obj
                  var try_3 EID 
                  { 
                    var v_bag_arg *ClaireAny  
                    try_3= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                    ToList(OBJ(try_3)).AddFast(C_range.Id())
                    var try_4 EID 
                    try_4 = Language.F_extract_type_any(g0113.Range.Id())
                    if ErrorIn(try_4) {try_3 = try_4
                    } else {
                    v_bag_arg = ANY(try_4)
                    ToList(OBJ(try_3)).AddFast(v_bag_arg)}
                    } 
                  if ErrorIn(try_3) {try_2 = try_3
                  } else {
                  va_arg2 = ToList(OBJ(try_3))
                  va_arg1.Args = va_arg2
                  try_2 = EID{va_arg2.Id(),0}
                  }
                  } 
                if !ErrorIn(try_2) {
                try_2 = EID{_CL_obj.Id(),0}
                }
                } 
              if ErrorIn(try_2) {try_1 = try_2
              } else {
              v_bag_arg = ANY(try_2)
              ToList(OBJ(try_1)).AddFast(v_bag_arg)
              var try_5 EID 
              { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                _CL_obj.Selector = ToProperty(C__equal.Id())
                { 
                  var va_arg1 *Language.Call  
                  var va_arg2 *ClaireList  
                  va_arg1 = _CL_obj
                  var try_6 EID 
                  { 
                    var v_bag_arg *ClaireAny  
                    try_6= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                    ToList(OBJ(try_6)).AddFast(C_value.Id())
                    var try_7 EID 
                    try_7 = r.Nexte()
                    if ErrorIn(try_7) {try_6 = try_7
                    } else {
                    v_bag_arg = ANY(try_7)
                    ToList(OBJ(try_6)).AddFast(v_bag_arg)}
                    } 
                  if ErrorIn(try_6) {try_5 = try_6
                  } else {
                  va_arg2 = ToList(OBJ(try_6))
                  va_arg1.Args = va_arg2
                  try_5 = EID{va_arg2.Id(),0}
                  }
                  } 
                if !ErrorIn(try_5) {
                try_5 = EID{_CL_obj.Id(),0}
                }
                } 
              if ErrorIn(try_5) {try_1 = try_5
              } else {
              v_bag_arg = ANY(try_5)
              ToList(OBJ(try_1)).AddFast(v_bag_arg)}}
              } 
            if ErrorIn(try_1) {Result = try_1
            } else {
            va_arg2 = ToList(OBJ(try_1))
            va_arg1.Args = va_arg2
            Result = EID{va_arg2.Id(),0}
            }
            } 
          if !ErrorIn(Result) {
          Result = EID{_CL_obj.Id(),0}
          }
          } 
        } 
      }  else if (x.Isa.IsIn(Language.C_Call) == CTRUE) { 
      { var g0114 *Language.Call   = Language.To_Call(x)
        { var ru *ClaireAny  
          _ = ru
          var try_8 EID 
          try_8 = r.Nexte()
          if ErrorIn(try_8) {Result = try_8
          } else {
          ru = ANY(try_8)
          { var z *ClaireAny  
            var try_9 EID 
            try_9 = r.Nexts(C__equal_sup)
            if ErrorIn(try_9) {Result = try_9
            } else {
            z = ANY(try_9)
            { var _CL_obj *Language.Defrule   = Language.To_Defrule(new(Language.Defrule).Is(Language.C_Defrule))
              _CL_obj.Ident = g0114.Selector.Name
              _CL_obj.Args = g0114.Args
              _CL_obj.Arg = z
              { 
                var va_arg1 *Language.Defrule  
                var va_arg2 *ClaireAny  
                va_arg1 = _CL_obj
                var try_10 EID 
                if (r.Firstc() == 41) { 
                  r.Next()
                  try_10 = EID{CNIL.Id(),0}
                  } else {
                  { var arg_11 *ClaireAny  
                    var try_12 EID 
                    try_12 = r.Nexte()
                    if ErrorIn(try_12) {try_10 = try_12
                    } else {
                    arg_11 = ANY(try_12)
                    try_10 = r.Readblock(arg_11,41)
                    }
                    } 
                  } 
                if ErrorIn(try_10) {Result = try_10
                } else {
                va_arg2 = ANY(try_10)
                va_arg1.Body = va_arg2
                Result = va_arg2.ToEID()
                }
                } 
              if !ErrorIn(Result) {
              Result = EID{_CL_obj.Id(),0}
              }
              } 
            }
            } 
          }
          } 
        } 
      } else {
      { var y *ClaireAny  
        var try_13 EID 
        try_13 = r.Nexte()
        if ErrorIn(try_13) {Result = try_13
        } else {
        y = ANY(try_13)
        var g0118I *ClaireBoolean  
        if (x.Isa.IsIn(Core.C_global_variable) == CTRUE) { 
          g0118I = Core.F_unknown_ask_any(y)
          } else {
          g0118I = CFALSE
          } 
        if (g0118I == CTRUE) { 
          Result = y.ToEID()
          } else {
          var g0119I *ClaireBoolean  
          if (y.Isa.IsIn(Language.C_Definition) == CTRUE) { 
            { var g0117 *Language.Definition   = Language.To_Definition(y)
              g0119I = g0117.Arg.IsIn(C_thing)
              } 
            } else {
            g0119I = CFALSE
            } 
          if (g0119I == CTRUE) { 
            { var _CL_obj *Language.Defobj   = Language.To_Defobj(new(Language.Defobj).Is(Language.C_Defobj))
              { 
                var va_arg1 *Language.Defobj  
                var va_arg2 *ClaireSymbol  
                va_arg1 = _CL_obj
                var try_14 EID 
                try_14 = Language.F_extract_symbol_any(x)
                if ErrorIn(try_14) {Result = try_14
                } else {
                va_arg2 = ToSymbol(OBJ(try_14))
                va_arg1.Ident = va_arg2
                Result = EID{va_arg2.Id(),0}
                }
                } 
              if !ErrorIn(Result) {
              _CL_obj.Arg = ToClass(OBJ(Core.F_CALL(C_arg,ARGS(y.ToEID()))))
              _CL_obj.Args = ToList(OBJ(Core.F_CALL(C_args,ARGS(y.ToEID()))))
              Result = EID{_CL_obj.Id(),0}
              }
              } 
            } else {
            { var _CL_obj *Language.Defobj   = Language.To_Defobj(new(Language.Defobj).Is(Language.C_Defobj))
              { 
                var va_arg1 *Language.Defobj  
                var va_arg2 *ClaireSymbol  
                va_arg1 = _CL_obj
                var try_15 EID 
                try_15 = Language.F_extract_symbol_any(x)
                if ErrorIn(try_15) {Result = try_15
                } else {
                va_arg2 = ToSymbol(OBJ(try_15))
                va_arg1.Ident = va_arg2
                Result = EID{va_arg2.Id(),0}
                }
                } 
              if !ErrorIn(Result) {
              _CL_obj.Arg = Core.C_global_variable
              { 
                var va_arg1 *Language.Definition  
                var va_arg2 *ClaireList  
                va_arg1 = Language.To_Definition(_CL_obj.Id())
                { 
                  var v_bag_arg *ClaireAny  
                  va_arg2= ToType(CEMPTY.Id()).EmptyList()
                  { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                    _CL_obj.Selector = ToProperty(C__equal.Id())
                    _CL_obj.Args = MakeConstantList(C_range.Id(),CEMPTY.Id())
                    v_bag_arg = _CL_obj.Id()
                    } 
                  va_arg2.AddFast(v_bag_arg)
                  { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                    _CL_obj.Selector = ToProperty(C__equal.Id())
                    _CL_obj.Args = MakeConstantList(C_value.Id(),y)
                    v_bag_arg = _CL_obj.Id()
                    } 
                  va_arg2.AddFast(v_bag_arg)} 
                va_arg1.Args = va_arg2
                } 
              Result = EID{_CL_obj.Id(),0}
              }
              } 
            } 
          } 
        }
        } 
      } 
    return Result} 
  
// The EID go function for: nextinst @ meta_reader (throw: true) 
func E_nextinst_meta_reader (r EID,x EID) EID { 
    return ToMetaReader(OBJ(r)).Nextinst(ANY(x) )} 
  
// reads a class Definition of the form C(p:t | p:t = v *)
// new in v2.5
/* {1} The go function for: nextDefclass(r:meta_reader,x:any,old?:boolean) [status=1] */
func (r *MetaReader ) NextDefclass (x *ClaireAny ,old_ask *ClaireBoolean ) EID { 
    // eid body s = Defclass
    var Result EID 
    r.Skipc()
    { var c *ClaireAny  
      var try_1 EID 
      { var arg_2 *ClaireAny  
        var try_3 EID 
        try_3 = r.Fromp.ReadIdent()
        if ErrorIn(try_3) {try_1 = try_3
        } else {
        arg_2 = ANY(try_3)
        try_1 = F_verify_any(C_class.Id(),arg_2,Language.C_Defclass.Id())
        }
        } 
      if ErrorIn(try_1) {Result = try_1
      } else {
      c = ANY(try_1)
      { var y *Language.Defclass  
        var try_4 EID 
        if (r.Firstc() != 40) { 
          { var _CL_obj *Language.Defclass   = Language.To_Defclass(new(Language.Defclass).Is(Language.C_Defclass))
            _CL_obj.Arg = ToClass(c)
            _CL_obj.Args = CNIL
            _CL_obj.Forward_ask = CTRUE
            try_4 = EID{_CL_obj.Id(),0}
            } 
          } else {
          { var l *ClaireAny  
            var try_5 EID 
            try_5 = r.Cnext().Nextseq(41)
            if ErrorIn(try_5) {try_4 = try_5
            } else {
            l = ANY(try_5)
            { 
              var y1 *ClaireAny  
              _ = y1
              try_4= EID{CFALSE.Id(),0}
              var y1_support *ClaireList  
              var try_6 EID 
              try_6 = Core.F_enumerate_any(l)
              if ErrorIn(try_6) {try_4 = try_6
              } else {
              y1_support = ToList(OBJ(try_6))
              y1_len := y1_support.Length()
              for i_it := 0; i_it < y1_len; i_it++ { 
                y1 = y1_support.At(i_it)
                var loop_7 EID 
                _ = loop_7
                var g0124I *ClaireBoolean  
                { var arg_8 *ClaireBoolean  
                  if (y1.Isa.IsIn(Language.C_Call) == CTRUE) { 
                    { var g0120 *Language.Call   = Language.To_Call(y1)
                      arg_8 = MakeBoolean((g0120.Selector.Id() == C__equal.Id()) && (g0120.Args.At(1-1).Isa.IsIn(Language.C_Vardef) == CTRUE))
                      } 
                    }  else if (y1.Isa.IsIn(Language.C_Vardef) == CTRUE) { 
                    arg_8 = CTRUE
                    } else {
                    arg_8 = CFALSE
                    } 
                  g0124I = arg_8.Not
                  } 
                if (g0124I == CTRUE) { 
                  loop_7 = F_Serror_string(MakeString("[175] Wrong form ~S in ~S(~S)"),MakeConstantList(y1,c,l))
                  } else {
                  loop_7 = EID{CFALSE.Id(),0}
                  } 
                if ErrorIn(loop_7) {try_4 = loop_7
                break
                } else {
                }}
                } 
              } 
            if !ErrorIn(try_4) {
            { var _CL_obj *Language.Defclass   = Language.To_Defclass(new(Language.Defclass).Is(Language.C_Defclass))
              _CL_obj.Arg = ToClass(c)
              _CL_obj.Args = ToList(l)
              _CL_obj.Forward_ask = CFALSE
              try_4 = EID{_CL_obj.Id(),0}
              } 
            }
            }
            } 
          } 
        if ErrorIn(try_4) {Result = try_4
        } else {
        y = Language.To_Defclass(OBJ(try_4))
        { var lp *ClaireList   = CNIL
          { var idt *ClaireSymbol  
            var try_9 EID 
            var g0125I *ClaireBoolean  
            if (x.Isa.IsIn(Language.C_Call) == CTRUE) { 
              { var g0123 *Language.Call   = Language.To_Call(x)
                g0125I = Equal(g0123.Selector.Id(),C_nth.Id())
                } 
              } else {
              g0125I = CFALSE
              } 
            if (g0125I == CTRUE) { 
              { var l *ClaireList   = Language.To_Call(x).Args
                if (C_class.Id() == l.At(1-1).Isa.Id()) { 
                  lp = ToList(l.At(2-1))
                  try_9 = EID{lp.Id(),0}
                  } else {
                  var try_10 EID 
                  { 
                    var v_list9 *ClaireList  
                    var y2 *ClaireAny  
                    var v_local9 *ClaireAny  
                    var try_11 EID 
                    try_11 = l.Cdr()
                    if ErrorIn(try_11) {try_10 = try_11
                    } else {
                    v_list9 = ToList(OBJ(try_11))
                    try_10 = EID{CreateList(ToType(CEMPTY.Id()),v_list9.Length()).Id(),0}
                    for CLcount := 0; CLcount < v_list9.Length(); CLcount++{ 
                      y2 = v_list9.At(CLcount)
                      var try_12 EID 
                      try_12 = Language.F_make_a_property_any(y2)
                      if ErrorIn(try_12) {try_10 = try_12
                      break
                      } else {
                      v_local9 = ANY(try_12)
                      ToList(OBJ(try_10)).PutAt(CLcount,v_local9)
                      } 
                    }}
                    } 
                  if ErrorIn(try_10) {try_9 = try_10
                  } else {
                  lp = ToList(OBJ(try_10))
                  try_9 = EID{lp.Id(),0}
                  }
                  } 
                if !ErrorIn(try_9) {
                try_9 = Language.F_extract_symbol_any(l.At(1-1))
                }
                } 
              } else {
              try_9 = Language.F_extract_symbol_any(x)
              } 
            if ErrorIn(try_9) {Result = try_9
            } else {
            idt = ToSymbol(OBJ(try_9))
            if ((old_ask == CTRUE) && 
                (r.Skipc() != 93)) { 
              Result = F_Serror_string(MakeString("[176] Missing ] after ~S "),MakeConstantList(y.Id()))
              }  else if (old_ask == CTRUE) { 
              r.Next()
              Result = EVOID
              } else {
              Result = EID{CFALSE.Id(),0}
              } 
            if !ErrorIn(Result) {
            y.Ident = idt
            y.Params = lp
            Result = EID{y.Id(),0}
            }
            }
            } 
          } 
        }
        } 
      }
      } 
    return Result} 
  
// The EID go function for: nextDefclass @ meta_reader (throw: true) 
func E_nextDefclass_meta_reader (r EID,x EID,old_ask EID) EID { 
    return ToMetaReader(OBJ(r)).NextDefclass(ANY(x),ToBoolean(OBJ(old_ask)) )} 
  
// end of file