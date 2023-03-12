/***** CLAIRE Compilation of file /Users/ycaseau/Dropbox/src/clairev4.07/src/compile/ocontrol.cl 
         [version 4.0.08 / safety 5] Sunday 03-12-2023 14:47:37 *****/

package Optimize
import (_ "fmt"
	. "Kernel"
	"Core"
	"Language"
	"Reader"
)

//-------- dumb function to prevent import errors --------
func import_g0236() { 
    _ = Core.It
    _ = Language.It
    _ = Reader.It
    } 
  
  
//+-------------------------------------------------------------+
//| CLAIRE                                                      |
//| ocontrol.cl                                                 |
//| Copyright (C) 1994 - 2021 Yves Caseau. All Rights Reserved  |
//| cf. copyright info in file object.cl: about()               |
//+-------------------------------------------------------------+
// *********************************************************************
// * Contents                                                          *
// *     Part 1: Basic Instructions                                    *
// *     Part 2: other control structures                              *
// *     Part 3: If, Case, Do, Let                                     *
// *     Part 4: Loops                                                 *
// *     Part 5: Iterate                                               *
// *********************************************************************
// *********************************************************************
// *      Part 1: Basic Instructions                                   *
// *********************************************************************
// the type of an assignment is the type of the result
/* The go function for: c_type(self:Assign) [status=1] */
func F_c_type_Assign (self *Language.Assign) EID { 
    var Result EID
    Result = Core.F_CALL(C_c_type,ARGS(self.Arg.ToEID()))
    return Result} 
  
// The EID go function for: c_type @ Assign (throw: true) 
func E_c_type_Assign (self EID) EID { 
    return F_c_type_Assign(Language.To_Assign(OBJ(self)) )} 
  
// we must include the type checking if needed
/* The go function for: c_code(self:Assign) [status=1] */
func F_c_code_Assign (self *Language.Assign) EID { 
    var Result EID
    { var v *ClaireAny = self.ClaireVar
      { var x *ClaireAny = self.Arg
        { var _Ztype *ClaireType
          var try_1 EID
          { var arg_2 *ClaireType
            var try_3 EID
            try_3 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
            if ErrorIn(try_3) {try_1 = try_3
            } else {
            arg_2 = ToType(OBJ(try_3))
            try_1 = EID{F_Optimize_ptype_type(arg_2).Id(),0}
            }
            } 
          if ErrorIn(try_1) {Result = try_1
          } else {
          _Ztype = ToType(OBJ(try_1))
          if (v.Isa.IsIn(C_Variable) != CTRUE) { 
            Result = Core.F_CALL(C_Compile_Cerror,ARGS(EID{MakeString("[213] ~S is not a variable").Id(),0},v.ToEID()))
            } else {
            Result = EID{CFALSE.Id(),0}
            } 
          if !ErrorIn(Result) {
          if (_Ztype.Included(ToType(OBJ(Core.F_CALL(C_range,ARGS(v.ToEID()))))) != CTRUE) { 
            var try_4 EID
            try_4 = Core.F_CALL(C_Optimize_c_warn,ARGS(self.ClaireVar.ToEID(),x.ToEID(),EID{_Ztype.Id(),0}))
            if ErrorIn(try_4) {Result = try_4
            } else {
            x = ANY(try_4)
            Result = x.ToEID()
            }
            } else {
            Result = EID{CFALSE.Id(),0}
            } 
          if !ErrorIn(Result) {
          { var _Zarg *ClaireAny
            var try_5 EID
            try_5 = F_Compile_c_strict_code_any(x,F_Compile_psort_any(ANY(Core.F_CALL(C_range,ARGS(v.ToEID())))))
            if ErrorIn(try_5) {Result = try_5
            } else {
            _Zarg = ANY(try_5)
            Result = Language.C_Assign.Make(v,_Zarg).ToEID()
            }
            } 
          }}
          }
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: c_code @ Assign (throw: true) 
func E_c_code_Assign (self EID) EID { 
    return F_c_code_Assign(Language.To_Assign(OBJ(self)) )} 
  
// assignment to a global variable
/* The go function for: c_type(self:Gassign) [status=1] */
func F_c_type_Gassign (self *Language.Gassign) EID { 
    var Result EID
    Result = Core.F_CALL(C_c_type,ARGS(self.Arg.ToEID()))
    return Result} 
  
// The EID go function for: c_type @ Gassign (throw: true) 
func E_c_type_Gassign (self EID) EID { 
    return F_c_type_Gassign(Language.To_Gassign(OBJ(self)) )} 
  
/* The go function for: c_code(self:Gassign) [status=1] */
func F_c_code_Gassign (self *Language.Gassign) EID { 
    var Result EID
    { var _Zv *ClaireAny = self.Arg
      { var _Ztype *ClaireType
        var try_1 EID
        { var arg_2 *ClaireType
          var try_3 EID
          try_3 = Core.F_CALL(C_c_type,ARGS(_Zv.ToEID()))
          if ErrorIn(try_3) {try_1 = try_3
          } else {
          arg_2 = ToType(OBJ(try_3))
          try_1 = EID{F_Optimize_ptype_type(arg_2).Id(),0}
          }
          } 
        if ErrorIn(try_1) {Result = try_1
        } else {
        _Ztype = ToType(OBJ(try_1))
        if (F_boolean_I_any(self.ClaireVar.Range.Id()).Id() != CTRUE.Id()) { 
          Result = Core.F_CALL(C_Compile_Cerror,ARGS(EID{MakeString("[214] cannot assign ~S").Id(),0},EID{self.Id(),0}))
          } else {
          Result = EID{CFALSE.Id(),0}
          } 
        if !ErrorIn(Result) {
        if (_Ztype.Included(self.ClaireVar.Range) != CTRUE) { 
          var try_4 EID
          try_4 = Core.F_CALL(C_c_code,ARGS(_Zv.ToEID(),EID{C_any.Id(),0}))
          if ErrorIn(try_4) {Result = try_4
          } else {
          _Zv = ANY(try_4)
          Result = _Zv.ToEID()
          }
          } else {
          Result = EID{CFALSE.Id(),0}
          } 
        if !ErrorIn(Result) {
        { var _CL_obj *Language.Gassign = Language.To_Gassign(new(Language.Gassign).Is(Language.C_Gassign))
          { 
            var va_arg1 *Language.Gassign
            var va_arg2 *Core.GlobalVariable
            va_arg1 = _CL_obj
            var try_5 EID
            try_5 = Core.F_CALL(C_c_code,ARGS(EID{self.ClaireVar.Id(),0}))
            if ErrorIn(try_5) {Result = try_5
            } else {
            va_arg2 = Core.ToGlobalVariable(OBJ(try_5))
            va_arg1.ClaireVar = va_arg2
            Result = EID{va_arg2.Id(),0}
            }
            } 
          if !ErrorIn(Result) {
          { 
            var va_arg1 *Language.Gassign
            var va_arg2 *ClaireAny
            va_arg1 = _CL_obj
            var try_6 EID
            if (F_Compile_nativeVar_ask_global_variable(self.ClaireVar) == CTRUE) { 
              try_6 = F_Compile_c_strict_code_any(_Zv,F_Compile_psort_any(self.ClaireVar.Range.Id()))
              } else {
              try_6 = Core.F_CALL(C_c_code,ARGS(_Zv.ToEID(),EID{C_any.Id(),0}))
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
          }}
          } 
        }}
        }
        } 
      } 
    return Result} 
  
// The EID go function for: c_code @ Gassign (throw: true) 
func E_c_code_Gassign (self EID) EID { 
    return F_c_code_Gassign(Language.To_Gassign(OBJ(self)) )} 
  
// v3.3 !
// _______________ l AND/OR     ____________________________________
/* The go function for: c_type(self:And) [status=0] */
func F_c_type_And (self *Language.And) *ClaireType { 
    return  ToType(C_boolean.Id())
    } 
  
// The EID go function for: c_type @ And (throw: false) 
func E_c_type_And (self EID) EID { 
    return EID{F_c_type_And(Language.To_And(OBJ(self)) ).Id(),0}} 
  
/* The go function for: c_code(self:And) [status=1] */
func F_c_code_And (self *Language.And) EID { 
    var Result EID
    { var _CL_obj *Language.And = Language.To_And(new(Language.And).Is(Language.C_And))
      { 
        var va_arg1 *Language.And
        var va_arg2 *ClaireList
        va_arg1 = _CL_obj
        var try_1 EID
        { 
          var v_list4 *ClaireList
          var x *ClaireAny
          var v_local4 *ClaireAny
          v_list4 = self.Args
          try_1 = EID{CreateList(ToType(CEMPTY.Id()),v_list4.Length()).Id(),0}
          for CLcount := 0; CLcount < v_list4.Length(); CLcount++{ 
            x = v_list4.At(CLcount)
            var try_2 EID
            var g0237I *ClaireBoolean
            var try_3 EID
            { var arg_4 *ClaireType
              var try_5 EID
              try_5 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
              if ErrorIn(try_5) {try_3 = try_5
              } else {
              arg_4 = ToType(OBJ(try_5))
              try_3 = EID{Equal(arg_4.Id(),C_void.Id()).Id(),0}
              }
              } 
            if ErrorIn(try_3) {try_2 = try_3
            } else {
            g0237I = ToBoolean(OBJ(try_3))
            if (g0237I == CTRUE) { 
              try_2 = Core.F_CALL(C_Compile_Cerror,ARGS(EID{MakeString("[206] void ~S in ~S").Id(),0},x.ToEID(),EID{self.Id(),0}))
              } else {
              try_2 = EID{CFALSE.Id(),0}
              } 
            }
            if ErrorIn(try_2) {try_1 = try_2
            break
            } else {
            try_2 = F_Optimize_c_boolean_any(x)
            if ErrorIn(try_2) {try_1 = try_2
            break
            } else {
            }}
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
        va_arg2 = ToList(OBJ(try_1))
        va_arg1.Args = va_arg2
        Result = EID{va_arg2.Id(),0}
        }
        } 
      if !ErrorIn(Result) {
      Result = EID{_CL_obj.Id(),0}
      }
      } 
    return Result} 
  
// The EID go function for: c_code @ And (throw: true) 
func E_c_code_And (self EID) EID { 
    return F_c_code_And(Language.To_And(OBJ(self)) )} 
  
/* The go function for: c_type(self:Or) [status=0] */
func F_c_type_Or (self *Language.Or) *ClaireType { 
    return  ToType(C_boolean.Id())
    } 
  
// The EID go function for: c_type @ Or (throw: false) 
func E_c_type_Or (self EID) EID { 
    return EID{F_c_type_Or(Language.To_Or(OBJ(self)) ).Id(),0}} 
  
/* The go function for: c_code(self:Or) [status=1] */
func F_c_code_Or (self *Language.Or) EID { 
    var Result EID
    { var _CL_obj *Language.Or = Language.To_Or(new(Language.Or).Is(Language.C_Or))
      { 
        var va_arg1 *Language.Or
        var va_arg2 *ClaireList
        va_arg1 = _CL_obj
        var try_1 EID
        { 
          var v_list4 *ClaireList
          var x *ClaireAny
          var v_local4 *ClaireAny
          v_list4 = self.Args
          try_1 = EID{CreateList(ToType(CEMPTY.Id()),v_list4.Length()).Id(),0}
          for CLcount := 0; CLcount < v_list4.Length(); CLcount++{ 
            x = v_list4.At(CLcount)
            var try_2 EID
            var g0238I *ClaireBoolean
            var try_3 EID
            { var arg_4 *ClaireType
              var try_5 EID
              try_5 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
              if ErrorIn(try_5) {try_3 = try_5
              } else {
              arg_4 = ToType(OBJ(try_5))
              try_3 = EID{Equal(arg_4.Id(),C_void.Id()).Id(),0}
              }
              } 
            if ErrorIn(try_3) {try_2 = try_3
            } else {
            g0238I = ToBoolean(OBJ(try_3))
            if (g0238I == CTRUE) { 
              try_2 = Core.F_CALL(C_Compile_Cerror,ARGS(EID{MakeString("[206] void ~S in ~S").Id(),0},x.ToEID(),EID{self.Id(),0}))
              } else {
              try_2 = EID{CFALSE.Id(),0}
              } 
            }
            if ErrorIn(try_2) {try_1 = try_2
            break
            } else {
            try_2 = F_Optimize_c_boolean_any(x)
            if ErrorIn(try_2) {try_1 = try_2
            break
            } else {
            }}
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
        va_arg2 = ToList(OBJ(try_1))
        va_arg1.Args = va_arg2
        Result = EID{va_arg2.Id(),0}
        }
        } 
      if !ErrorIn(Result) {
      Result = EID{_CL_obj.Id(),0}
      }
      } 
    return Result} 
  
// The EID go function for: c_code @ Or (throw: true) 
func E_c_code_Or (self EID) EID { 
    return F_c_code_Or(Language.To_Or(OBJ(self)) )} 
  
//---------------- quote and return -------------------------------------
/* The go function for: c_type(self:Quote) [status=0] */
func F_c_type_Quote (self *Language.Quote) *ClaireType { 
    return  ToType(self.Arg.Isa.Id())
    } 
  
// The EID go function for: c_type @ Quote (throw: false) 
func E_c_type_Quote (self EID) EID { 
    return EID{F_c_type_Quote(Language.To_Quote(OBJ(self)) ).Id(),0}} 
  
/* The go function for: c_code(self:Quote) [status=1] */
func F_c_code_Quote (self *Language.Quote) EID { 
    var Result EID
    Result = Core.F_CALL(C_Compile_Cerror,ARGS(EID{MakeString("[internal] optimization of quote not implemented yet! ~S").Id(),0},EID{self.Id(),0}))
    return Result} 
  
// The EID go function for: c_code @ Quote (throw: true) 
func E_c_code_Quote (self EID) EID { 
    return F_c_code_Quote(Language.To_Quote(OBJ(self)) )} 
  
/* The go function for: c_type(self:Return) [status=0] */
func F_c_type_Return (self *Language.Return) *ClaireType { 
    return  ToType(C_any.Id())
    } 
  
// The EID go function for: c_type @ Return (throw: false) 
func E_c_type_Return (self EID) EID { 
    return EID{F_c_type_Return(Language.To_Return(OBJ(self)) ).Id(),0}} 
  
/* The go function for: c_code(self:Return) [status=1] */
func F_c_code_Return (self *Language.Return) EID { 
    var Result EID
    { var arg_1 *ClaireAny
      var try_2 EID
      try_2 = Core.F_CALL(C_c_code,ARGS(self.Arg.ToEID(),EID{C_any.Id(),0}))
      if ErrorIn(try_2) {Result = try_2
      } else {
      arg_1 = ANY(try_2)
      Result = Language.C_Return.Make(arg_1).ToEID()
      }
      } 
    return Result} 
  
// The EID go function for: c_code @ Return (throw: true) 
func E_c_code_Return (self EID) EID { 
    return F_c_code_Return(Language.To_Return(OBJ(self)) )} 
  
// optimisation of exception handlers
/* The go function for: c_type(self:Handle) [status=1] */
func F_c_type_Handle (self *Language.ClaireHandle) EID { 
    var Result EID
    { var arg_1 *ClaireType
      var try_3 EID
      try_3 = Core.F_CALL(C_c_type,ARGS(self.Arg.ToEID()))
      if ErrorIn(try_3) {Result = try_3
      } else {
      arg_1 = ToType(OBJ(try_3))
      { var arg_2 *ClaireType
        var try_4 EID
        try_4 = Core.F_CALL(C_c_type,ARGS(self.Other.ToEID()))
        if ErrorIn(try_4) {Result = try_4
        } else {
        arg_2 = ToType(OBJ(try_4))
        Result = EID{Core.F_U_type(arg_1,arg_2).Id(),0}
        }
        } 
      }
      } 
    return Result} 
  
// The EID go function for: c_type @ Handle (throw: true) 
func E_c_type_Handle (self EID) EID { 
    return F_c_type_Handle(Language.To_ClaireHandle(OBJ(self)) )} 
  
/* The go function for: c_code(self:Handle,s:class) [status=1] */
func F_c_code_Handle (self *Language.ClaireHandle,s *ClaireClass) EID { 
    var Result EID
    { var x *Language.ClaireHandle
      var try_1 EID
      { var arg_2 *ClaireAny
        var try_4 EID
        try_4 = Core.F_CALL(C_c_code,ARGS(self.Arg.ToEID(),EID{s.Id(),0}))
        if ErrorIn(try_4) {try_1 = try_4
        } else {
        arg_2 = ANY(try_4)
        { var arg_3 *ClaireAny
          var try_5 EID
          try_5 = Core.F_CALL(C_c_code,ARGS(self.Other.ToEID(),EID{s.Id(),0}))
          if ErrorIn(try_5) {try_1 = try_5
          } else {
          arg_3 = ANY(try_5)
          try_1 = Language.C_Handle.Make(C_any.Id(),arg_2,arg_3).ToEID()
          }
          } 
        }
        } 
      if ErrorIn(try_1) {Result = try_1
      } else {
      x = Language.To_ClaireHandle(OBJ(try_1))
      x.Test = self.Test
      Result = EID{x.Id(),0}
      }
      } 
    return Result} 
  
// The EID go function for: c_code @ Handle (throw: true) 
func E_c_code_Handle (self EID,s EID) EID { 
    return F_c_code_Handle(Language.To_ClaireHandle(OBJ(self)),ToClass(OBJ(s)) )} 
  
// ****************************************************************
// *      Part 2: Specific structures                             *
// ****************************************************************
// __________________ CAST ________________________________________
// cast is now more subtle since we introduce coercion for list<t> types
/* The go function for: c_type(self:Cast) [status=0] */
func F_c_type_Cast (self *Language.Cast) *ClaireType { 
    return  self.SetArg
    } 
  
// The EID go function for: c_type @ Cast (throw: false) 
func E_c_type_Cast (self EID) EID { 
    return EID{F_c_type_Cast(Language.To_Cast(OBJ(self)) ).Id(),0}} 
  
// insert dynamic types (check_in) when we see a claire cast
// CLAIRE 4 : when we decide to drop the cast (safety), we generate a C_cast
/* The go function for: c_code(self:Cast) [status=1] */
func F_c_code_Cast (self *Language.Cast) EID { 
    var Result EID
    { var y *ClaireType = self.SetArg
      { var ftype *ClaireClass = F_Compile_psort_any(y.Id())
        var g0240I *ClaireBoolean
        if (y.Isa.IsIn(C_Param) == CTRUE) { 
          { var g0239 *ClaireParam = To_Param(y.Id())
            g0240I = MakeBoolean(((g0239.Arg.Id() == C_list.Id()) || 
                (g0239.Arg.Id() == C_set.Id())) && (C_set.Id() == g0239.Args.At(0).Isa.Id()))
            } 
          } else {
          g0240I = CFALSE
          } 
        if (g0240I == CTRUE) { 
          { var utype *ClaireAny
            var try_1 EID
            try_1 = Core.F_the_type(ToType(To_Param(y.Id()).Args.At(0)))
            if ErrorIn(try_1) {Result = try_1
            } else {
            utype = ANY(try_1)
            var g0241I *ClaireBoolean
            var try_2 EID
            { 
              var v_or6 *ClaireBoolean
              
              var try_3 EID
              { var arg_4 *ClaireType
                var try_5 EID
                { var arg_6 *ClaireType
                  var try_7 EID
                  try_7 = Core.F_CALL(C_c_type,ARGS(self.Arg.ToEID()))
                  if ErrorIn(try_7) {try_5 = try_7
                  } else {
                  arg_6 = ToType(OBJ(try_7))
                  try_5 = EID{arg_6.At(C_of).Id(),0}
                  }
                  } 
                if ErrorIn(try_5) {try_3 = try_5
                } else {
                arg_4 = ToType(OBJ(try_5))
                try_3 = EID{Equal(arg_4.Id(),utype).Id(),0}
                }
                } 
              if ErrorIn(try_3) {try_2 = try_3
              } else {
              v_or6 = ToBoolean(OBJ(try_3))
              if (v_or6 == CTRUE) {try_2 = EID{CTRUE.Id(),0}
              } else { 
                v_or6 = F__sup_equal_integer(C_compiler.Safety,2)
                if (v_or6 == CTRUE) {try_2 = EID{CTRUE.Id(),0}
                } else { 
                  try_2 = EID{CFALSE.Id(),0}} 
                } 
              }
              } 
            if ErrorIn(try_2) {Result = try_2
            } else {
            g0241I = ToBoolean(OBJ(try_2))
            if (g0241I == CTRUE) { 
              Result = Core.F_CALL(C_c_code,ARGS(self.Arg.ToEID(),EID{ftype.Id(),0}))
              } else {
              { var arg_8 *Language.Call
                { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                  _CL_obj.Selector = Core.C_check_in
                  _CL_obj.Args = MakeConstantList(self.Arg,To_Param(y.Id()).Arg.Id(),utype)
                  arg_8 = _CL_obj
                  } 
                Result = Core.F_CALL(C_c_code,ARGS(EID{arg_8.Id(),0},EID{ftype.Id(),0}))
                } 
              } 
            }
            }
            } 
          } else {
          var g0242I *ClaireBoolean
          var try_9 EID
          { var arg_10 *ClaireType
            var try_11 EID
            try_11 = Core.F_CALL(C_c_type,ARGS(self.Arg.ToEID()))
            if ErrorIn(try_11) {try_9 = try_11
            } else {
            arg_10 = ToType(OBJ(try_11))
            try_9 = EID{arg_10.Included(y).Id(),0}
            }
            } 
          if ErrorIn(try_9) {Result = try_9
          } else {
          g0242I = ToBoolean(OBJ(try_9))
          if (g0242I == CTRUE) { 
            Result = Core.F_CALL(C_c_code,ARGS(self.Arg.ToEID(),EID{ftype.Id(),0}))
            }  else if (C_compiler.Safety >= 2) { 
            { var _CL_obj *Compile_CCast = To_Compile_CCast(new(Compile_CCast).Is(C_Compile_C_cast))
              { 
                var va_arg1 *Compile_CCast
                var va_arg2 *ClaireAny
                va_arg1 = _CL_obj
                var try_12 EID
                try_12 = Core.F_CALL(C_c_code,ARGS(self.Arg.ToEID(),EID{ftype.Id(),0}))
                if ErrorIn(try_12) {Result = try_12
                } else {
                va_arg2 = ANY(try_12)
                va_arg1.Arg = va_arg2
                Result = va_arg2.ToEID()
                }
                } 
              if !ErrorIn(Result) {
              _CL_obj.SetArg = ToClass(y.Id())
              Result = EID{_CL_obj.Id(),0}
              }
              } 
            } else {
            { var arg_13 *Language.Call
              { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                _CL_obj.Selector = Core.C_check_in
                _CL_obj.Args = MakeConstantList(self.Arg,y.Id())
                arg_13 = _CL_obj
                } 
              Result = Core.F_CALL(C_c_code,ARGS(EID{arg_13.Id(),0},EID{ftype.Id(),0}))
              } 
            } 
          }
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: c_code @ Cast (throw: true) 
func E_c_code_Cast (self EID) EID { 
    return F_c_code_Cast(Language.To_Cast(OBJ(self)) )} 
  
// _________________ SUPER _________________________________________
/* The go function for: c_type(self:Super) [status=1] */
func F_c_type_Super (self *Language.Super) EID { 
    var Result EID
    { var _Ztype *ClaireList
      var try_1 EID
      { 
        var v_list3 *ClaireList
        var x *ClaireAny
        var v_local3 *ClaireAny
        v_list3 = self.Args
        try_1 = EID{CreateList(ToType(CEMPTY.Id()),v_list3.Length()).Id(),0}
        for CLcount := 0; CLcount < v_list3.Length(); CLcount++{ 
          x = v_list3.At(CLcount)
          var try_2 EID
          try_2 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
          if ErrorIn(try_2) {try_1 = try_2
          break
          } else {
          v_local3 = ANY(try_2)
          ToList(OBJ(try_1)).PutAt(CLcount,v_local3)
          } 
        }
        } 
      if ErrorIn(try_1) {Result = try_1
      } else {
      _Ztype = ToList(OBJ(try_1))
      { var s *ClaireProperty = self.Selector
        ToArray(_Ztype.Id()).NthPut(1,self.CastTo.Id())
        { var prop *ClaireAny
          if (s.Open == 3) { 
            prop = CNIL.Id()
            } else {
            prop = F_Optimize_restriction_I_class(self.CastTo.Class_I(),s.Definition,_Ztype)
            } 
          if (C_slot.Id() == prop.Isa.Id()) { 
            { var g0243 *ClaireSlot = ToSlot(prop)
              Result = EID{g0243.Range.Id(),0}
              } 
            }  else if (C_method.Id() == prop.Isa.Id()) { 
            { var g0244 *ClaireMethod = ToMethod(prop)
              Result = F_Optimize_use_range_method(g0244,_Ztype)
              } 
            } else {
            Result = EID{s.Range.Id(),0}
            } 
          } 
        } 
      }
      } 
    return Result} 
  
// The EID go function for: c_type @ Super (throw: true) 
func E_c_type_Super (self EID) EID { 
    return F_c_type_Super(Language.To_Super(OBJ(self)) )} 
  
// this is the optimizer for messages
/* The go function for: c_code(self:Super) [status=1] */
func F_c_code_Super (self *Language.Super) EID { 
    var Result EID
    { var s *ClaireProperty = self.Selector
      { var l *ClaireList = self.Args
        { var _Ztype *ClaireList
          var try_1 EID
          { 
            var v_list5 *ClaireList
            var x *ClaireAny
            var v_local5 *ClaireAny
            v_list5 = self.Args
            try_1 = EID{CreateList(ToType(CEMPTY.Id()),v_list5.Length()).Id(),0}
            for CLcount := 0; CLcount < v_list5.Length(); CLcount++{ 
              x = v_list5.At(CLcount)
              var try_2 EID
              try_2 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
              if ErrorIn(try_2) {try_1 = try_2
              break
              } else {
              v_local5 = ANY(try_2)
              ToList(OBJ(try_1)).PutAt(CLcount,v_local5)
              } 
            }
            } 
          if ErrorIn(try_1) {Result = try_1
          } else {
          _Ztype = ToList(OBJ(try_1))
          { var prop *ClaireAny
            if (s.Open == 3) { 
              prop = CNIL.Id()
              } else {
              prop = F_Optimize_restriction_I_class(self.CastTo.Class_I(),s.Definition,_Ztype)
              } 
            if (C_slot.Id() == prop.Isa.Id()) { 
              { var g0246 *ClaireSlot = ToSlot(prop)
                { var _CL_obj *Language.CallSlot = Language.To_CallSlot(new(Language.CallSlot).Is(Language.C_Call_slot))
                  _CL_obj.Selector = g0246
                  { 
                    var va_arg1 *Language.CallSlot
                    var va_arg2 *ClaireAny
                    va_arg1 = _CL_obj
                    var try_3 EID
                    try_3 = Core.F_CALL(C_c_code,ARGS(l.At(0).ToEID(),EID{F_Compile_psort_any(Core.F_domain_I_restriction(ToRestriction(g0246.Id())).Id()).Id(),0}))
                    if ErrorIn(try_3) {Result = try_3
                    } else {
                    va_arg2 = ANY(try_3)
                    va_arg1.Arg = va_arg2
                    Result = va_arg2.ToEID()
                    }
                    } 
                  if !ErrorIn(Result) {
                  _CL_obj.Test = MakeBoolean((g0246.Range.Contains(g0246.Default) != CTRUE) && (C_compiler.Safety < 5))
                  Result = EID{_CL_obj.Id(),0}
                  }
                  } 
                } 
              }  else if (C_method.Id() == prop.Isa.Id()) { 
              { var g0247 *ClaireMethod = ToMethod(prop)
                Result = F_Optimize_c_code_method_method1(g0247,l,_Ztype)
                } 
              } else {
              Result = F_Optimize_c_warn_Super(self,_Ztype.Id())
              } 
            } 
          }
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: c_code @ Super (throw: true) 
func E_c_code_Super (self EID) EID { 
    return F_c_code_Super(Language.To_Super(OBJ(self)) )} 
  
// we will need this direct call for compiling call to CLAIRE_demons
/* The go function for: self_print(self:Call_function2) [status=1] */
func (self *Optimize_CallFunction2) SelfPrint () EID { 
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
  
// The EID go function for: self_print @ Call_function2 (throw: true) 
func E_self_print_Call_function2 (self EID) EID { 
    return To_Optimize_CallFunction2(OBJ(self)).SelfPrint( )} 
  
/* The go function for: c_type(self:Call_function2) [status=0] */
func (self *Optimize_CallFunction2) CType () *ClaireType { 
    return  ToType(C_any.Id())
    } 
  
// The EID go function for: c_type @ Call_function2 (throw: false) 
func E_c_type_Call_function2 (self EID) EID { 
    return EID{To_Optimize_CallFunction2(OBJ(self)).CType( ).Id(),0}} 
  
/* The go function for: c_code(self:Call_function2) [status=1] */
func F_c_code_Call_function2 (self *Optimize_CallFunction2) EID { 
    var Result EID
    { var _CL_obj *Optimize_CallFunction2 = To_Optimize_CallFunction2(new(Optimize_CallFunction2).Is(C_Optimize_Call_function2))
      _CL_obj.Arg = self.Arg
      { 
        var va_arg1 *Optimize_CallFunction2
        var va_arg2 *ClaireList
        va_arg1 = _CL_obj
        var try_1 EID
        { 
          var v_list4 *ClaireList
          var x *ClaireAny
          var v_local4 *ClaireAny
          v_list4 = self.Args
          try_1 = EID{CreateList(ToType(CEMPTY.Id()),v_list4.Length()).Id(),0}
          for CLcount := 0; CLcount < v_list4.Length(); CLcount++{ 
            x = v_list4.At(CLcount)
            var try_2 EID
            try_2 = Core.F_CALL(C_c_code,ARGS(x.ToEID(),EID{C_any.Id(),0}))
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
        va_arg2 = ToList(OBJ(try_1))
        va_arg1.Args = va_arg2
        Result = EID{va_arg2.Id(),0}
        }
        } 
      if !ErrorIn(Result) {
      Result = EID{_CL_obj.Id(),0}
      }
      } 
    return Result} 
  
// The EID go function for: c_code @ Call_function2 (throw: true) 
func E_c_code_Call_function2 (self EID) EID { 
    return F_c_code_Call_function2(To_Optimize_CallFunction2(OBJ(self)) )} 
  
// ASSERT & trace
/* The go function for: c_code(self:Assert) [status=1] */
func F_c_code_Assert (self *Language.Assert) EID { 
    var Result EID
    if ((C_compiler.Safety == 0) || 
        (C_compiler.Debug_ask.Length() != 0)) { 
      { var arg_1 *ClaireObject
        { var arg_2 *Language.Call
          { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
            _CL_obj.Selector = Core.C_not
            _CL_obj.Args = MakeConstantList(self.Args.At(0))
            arg_2 = _CL_obj
            } 
          { var arg_3 *Language.Call
            { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
              _CL_obj.Selector = Core.C_Core_tformat
              { 
                var va_arg1 *Language.Call
                var va_arg2 *ClaireList
                va_arg1 = _CL_obj
                { 
                  var v_bag_arg *ClaireAny
                  va_arg2= ToType(CEMPTY.Id()).EmptyList()
                  va_arg2.AddFast(MakeString("Assertion violation in ~A line ~A\n").Id())
                  va_arg2.AddFast(MakeInteger(0).Id())
                  { var _CL_obj *Language.List = Language.To_List(new(Language.List).Is(Language.C_List))
                    _CL_obj.Args = MakeConstantList((self.External).Id(),MakeInteger(self.Index).Id())
                    v_bag_arg = _CL_obj.Id()
                    } 
                  va_arg2.AddFast(v_bag_arg)} 
                va_arg1.Args = va_arg2
                } 
              arg_3 = _CL_obj
              } 
            arg_1 = ToObject(Language.C_If.Make(arg_2.Id(),arg_3.Id(),CFALSE.Id()))
            } 
          } 
        Result = Core.F_CALL(C_c_code,ARGS(EID{arg_1.Id(),0},EID{C_any.Id(),0}))
        } 
      } else {
      Result = EID{CNIL.Id(),0}
      } 
    return Result} 
  
// The EID go function for: c_code @ Assert (throw: true) 
func E_c_code_Assert (self EID) EID { 
    return F_c_code_Assert(Language.To_Assert(OBJ(self)) )} 
  
// ignore assertion
// CLAIRE4 : we keep traces up to levels 2
/* The go function for: c_code(self:Trace) [status=1] */
func F_c_code_Trace (self *Language.Trace) EID { 
    var Result EID
    { var a *ClaireList = self.Args
      var g0249I *ClaireBoolean
      var try_1 EID
      { 
        var v_and3 *ClaireBoolean
        
        v_and3 = Equal(MakeInteger(a.Length()).Id(),MakeInteger(1).Id())
        if (v_and3 == CFALSE) {try_1 = EID{CFALSE.Id(),0}
        } else { 
          var try_2 EID
          { var arg_3 *ClaireType
            var try_4 EID
            try_4 = Core.F_CALL(C_c_type,ARGS(a.At(0).ToEID()))
            if ErrorIn(try_4) {try_2 = try_4
            } else {
            arg_3 = ToType(OBJ(try_4))
            try_2 = EID{arg_3.Included(ToType(C_integer.Id())).Id(),0}
            }
            } 
          if ErrorIn(try_2) {try_1 = try_2
          } else {
          v_and3 = ToBoolean(OBJ(try_2))
          if (v_and3 == CFALSE) {try_1 = EID{CFALSE.Id(),0}
          } else { 
            try_1 = EID{CTRUE.Id(),0}} 
          } 
        }
        } 
      if ErrorIn(try_1) {Result = try_1
      } else {
      g0249I = ToBoolean(OBJ(try_1))
      if (g0249I == CTRUE) { 
        { var arg_5 *Language.Call
          { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
            _CL_obj.Selector = Core.C_write
            _CL_obj.Args = MakeConstantList(C_verbose.Id(),ClEnv.Id(),a.At(0))
            arg_5 = _CL_obj
            } 
          Result = Core.F_CALL(C_c_code,ARGS(EID{arg_5.Id(),0}))
          } 
        } else {
        var g0250I *ClaireBoolean
        { 
          var v_and4 *ClaireBoolean
          
          v_and4 = Core.F__sup_integer(a.Length(),1)
          if (v_and4 == CFALSE) {g0250I = CFALSE
          } else { 
            v_and4 = Equal(C_string.Id(),a.At(1).Isa.Id())
            if (v_and4 == CFALSE) {g0250I = CFALSE
            } else { 
              { 
                var v_or7 *ClaireBoolean
                
                v_or7 = Core.F__I_equal_any(MakeInteger(C_compiler.Debug_ask.Length()).Id(),MakeInteger(0).Id())
                if (v_or7 == CTRUE) {v_and4 = CTRUE
                } else { 
                  { 
                    var v_or7_H EID
                    h_index := ClEnv.Index
                    h_base := ClEnv.Base
                    { var arg_6 *ClaireAny
                      var try_7 EID
                      try_7 = EVAL(a.At(0))
                      if ErrorIn(try_7) {v_or7_H = try_7
                      } else {
                      arg_6 = ANY(try_7)
                      v_or7_H = EID{Core.F__inf_equal_integer(ToInteger(arg_6).Value,Reader.F_max_integer(3,ClEnv.Verbose)).Id(),0}
                      }
                      } 
                    if ErrorIn(v_or7_H){ 
                      ClEnv.Index = h_index
                      ClEnv.Base = h_base
                      v_or7 = CTRUE
                      } else {
                      v_or7 = ToBoolean(OBJ(v_or7_H))
                      } 
                    } 
                  if (v_or7 == CTRUE) {v_and4 = CTRUE
                  } else { 
                    v_and4 = CFALSE} 
                  } 
                } 
              if (v_and4 == CFALSE) {g0250I = CFALSE
              } else { 
                g0250I = CTRUE} 
              } 
            } 
          } 
        if (g0250I == CTRUE) { 
          { var _Zc *Language.Call
            { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
              _CL_obj.Selector = Core.C_Core_tformat
              { 
                var va_arg1 *Language.Call
                var va_arg2 *ClaireList
                va_arg1 = _CL_obj
                { 
                  var v_bag_arg *ClaireAny
                  va_arg2= ToType(CEMPTY.Id()).EmptyList()
                  va_arg2.AddFast(a.At(1))
                  va_arg2.AddFast(a.At(0))
                  { var _CL_obj *Language.List = Language.To_List(new(Language.List).Is(Language.C_List))
                    _CL_obj.Args = a.Copy().Skip(2)
                    v_bag_arg = _CL_obj.Id()
                    } 
                  va_arg2.AddFast(v_bag_arg)} 
                va_arg1.Args = va_arg2
                } 
              _Zc = _CL_obj
              } 
            { var arg_8 *ClaireObject
              if (C_integer.Id() != a.At(0).Isa.Id()) { 
                { var arg_9 *Language.Call
                  { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                    _CL_obj.Selector = ToProperty(C__inf_equal.Id())
                    { 
                      var va_arg1 *Language.Call
                      var va_arg2 *ClaireList
                      va_arg1 = _CL_obj
                      { 
                        var v_bag_arg *ClaireAny
                        va_arg2= ToType(CEMPTY.Id()).EmptyList()
                        va_arg2.AddFast(a.At(0))
                        { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                          _CL_obj.Selector = C_verbose
                          _CL_obj.Args = MakeConstantList(ClEnv.Id())
                          v_bag_arg = _CL_obj.Id()
                          } 
                        va_arg2.AddFast(v_bag_arg)} 
                      va_arg1.Args = va_arg2
                      } 
                    arg_9 = _CL_obj
                    } 
                  arg_8 = ToObject(Language.C_If.Make(arg_9.Id(),_Zc.Id(),CFALSE.Id()))
                  } 
                } else {
                arg_8 = ToObject(_Zc.Id())
                } 
              Result = Core.F_CALL(C_c_code,ARGS(EID{arg_8.Id(),0},EID{C_any.Id(),0}))
              } 
            } 
          } else {
          Result = EID{CFALSE.Id(),0}
          } 
        } 
      }
      } 
    return Result} 
  
// The EID go function for: c_code @ Trace (throw: true) 
func E_c_code_Trace (self EID) EID { 
    return F_c_code_Trace(Language.To_Trace(OBJ(self)) )} 
  
/* The go function for: c_type(self:Assert) [status=0] */
func F_c_type_Assert (self *Language.Assert) *ClaireType { 
    return  ToType(C_any.Id())
    } 
  
// The EID go function for: c_type @ Assert (throw: false) 
func E_c_type_Assert (self EID) EID { 
    return EID{F_c_type_Assert(Language.To_Assert(OBJ(self)) ).Id(),0}} 
  
/* The go function for: c_type(self:Trace) [status=0] */
func F_c_type_Trace (self *Language.Trace) *ClaireType { 
    return  ToType(C_any.Id())
    } 
  
// The EID go function for: c_type @ Trace (throw: false) 
func E_c_type_Trace (self EID) EID { 
    return EID{F_c_type_Trace(Language.To_Trace(OBJ(self)) ).Id(),0}} 
  
/* The go function for: c_type(self:Branch) [status=0] */
func F_c_type_Branch (self *Language.Branch) *ClaireType { 
    return  ToType(C_boolean.Id())
    } 
  
// The EID go function for: c_type @ Branch (throw: false) 
func E_c_type_Branch (self EID) EID { 
    return EID{F_c_type_Branch(Language.To_Branch(OBJ(self)) ).Id(),0}} 
  
/* The go function for: c_code(self:Branch) [status=1] */
func F_c_code_Branch (self *Language.Branch) EID { 
    var Result EID
    { var arg_1 *ClaireObject
      { var arg_2 *Language.Do
        { var _CL_obj *Language.Do = Language.To_Do(new(Language.Do).Is(Language.C_Do))
          { 
            var va_arg1 *Language.Do
            var va_arg2 *ClaireList
            va_arg1 = _CL_obj
            { 
              var v_bag_arg *ClaireAny
              va_arg2= ToType(CEMPTY.Id()).EmptyList()
              { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                _CL_obj.Selector = C_choice
                _CL_obj.Args = MakeConstantList(ClEnv.Id())
                v_bag_arg = _CL_obj.Id()
                } 
              va_arg2.AddFast(v_bag_arg)
              { var arg_4 *Language.Do
                { var _CL_obj *Language.Do = Language.To_Do(new(Language.Do).Is(Language.C_Do))
                  { 
                    var va_arg1 *Language.Do
                    var va_arg2 *ClaireList
                    va_arg1 = _CL_obj
                    { 
                      var v_bag_arg *ClaireAny
                      va_arg2= ToType(CEMPTY.Id()).EmptyList()
                      { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                        _CL_obj.Selector = C_backtrack
                        _CL_obj.Args = MakeConstantList(ClEnv.Id())
                        v_bag_arg = _CL_obj.Id()
                        } 
                      va_arg2.AddFast(v_bag_arg)
                      va_arg2.AddFast(CFALSE.Id())} 
                    va_arg1.Args = va_arg2
                    } 
                  arg_4 = _CL_obj
                  } 
                v_bag_arg = Language.C_If.Make(self.Args.At(0),CTRUE.Id(),arg_4.Id())
                } 
              va_arg2.AddFast(v_bag_arg)} 
            va_arg1.Args = va_arg2
            } 
          arg_2 = _CL_obj
          } 
        { var arg_3 *Language.Do
          { var _CL_obj *Language.Do = Language.To_Do(new(Language.Do).Is(Language.C_Do))
            { 
              var va_arg1 *Language.Do
              var va_arg2 *ClaireList
              va_arg1 = _CL_obj
              { 
                var v_bag_arg *ClaireAny
                va_arg2= ToType(CEMPTY.Id()).EmptyList()
                { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                  _CL_obj.Selector = C_backtrack
                  _CL_obj.Args = MakeConstantList(ClEnv.Id())
                  v_bag_arg = _CL_obj.Id()
                  } 
                va_arg2.AddFast(v_bag_arg)
                va_arg2.AddFast(CFALSE.Id())} 
              va_arg1.Args = va_arg2
              } 
            arg_3 = _CL_obj
            } 
          arg_1 = ToObject(Language.C_Handle.Make(Core.C_contradiction.Id(),arg_2.Id(),arg_3.Id()))
          } 
        } 
      Result = Core.F_CALL(C_c_code,ARGS(EID{arg_1.Id(),0},EID{C_any.Id(),0}))
      } 
    return Result} 
  
// The EID go function for: c_code @ Branch (throw: true) 
func E_c_code_Branch (self EID) EID { 
    return F_c_code_Branch(Language.To_Branch(OBJ(self)) )} 
  
/* The go function for: c_code(self:Macro,s:class) [status=1] */
func F_c_code_Macro (self *Language.Macro,s *ClaireClass) EID { 
    var Result EID
    { var arg_1 *ClaireAny
      var try_2 EID
      try_2 = Core.F_CALL(Language.C_macroexpand,ARGS(EID{self.Id(),0}))
      if ErrorIn(try_2) {Result = try_2
      } else {
      arg_1 = ANY(try_2)
      Result = Core.F_CALL(C_c_code,ARGS(arg_1.ToEID(),EID{s.Id(),0}))
      }
      } 
    return Result} 
  
// The EID go function for: c_code @ Macro (throw: true) 
func E_c_code_Macro (self EID,s EID) EID { 
    return F_c_code_Macro(Language.To_Macro(OBJ(self)),ToClass(OBJ(s)) )} 
  
/* The go function for: c_type(self:Macro) [status=1] */
func F_c_type_Macro (self *Language.Macro) EID { 
    var Result EID
    { var arg_1 *ClaireAny
      var try_2 EID
      try_2 = Core.F_CALL(Language.C_macroexpand,ARGS(EID{self.Id(),0}))
      if ErrorIn(try_2) {Result = try_2
      } else {
      arg_1 = ANY(try_2)
      Result = Core.F_CALL(C_c_type,ARGS(arg_1.ToEID()))
      }
      } 
    return Result} 
  
// The EID go function for: c_type @ Macro (throw: true) 
func E_c_type_Macro (self EID) EID { 
    return F_c_type_Macro(Language.To_Macro(OBJ(self)) )} 
  
/* The go function for: c_type(self:Printf) [status=0] */
func F_c_type_Printf (self *Language.Printf) *ClaireType { 
    return  ToType(C_any.Id())
    } 
  
// The EID go function for: c_type @ Printf (throw: false) 
func E_c_type_Printf (self EID) EID { 
    return EID{F_c_type_Printf(Language.To_Printf(OBJ(self)) ).Id(),0}} 
  
/* The go function for: c_code(self:Printf) [status=1] */
func F_c_code_Printf (self *Language.Printf) EID { 
    var Result EID
    { var l *ClaireList = self.Args
      if (C_string.Id() != l.At(0).Isa.Id()) { 
        Result = Core.F_CALL(C_Compile_Cerror,ARGS(EID{MakeString("[209] the first argument in ~S must be a string").Id(),0},EID{self.Id(),0}))
        } else {
        { var s *ClaireString = ToString(l.At(0))
          { var i int = 1
            { var n int = F_get_string(s,'~')
              { var r *ClaireList = ToType(C_any.Id()).EmptyList()
                Result= EID{CFALSE.Id(),0}
                for (n != 0) { 
                  var loop_1 EID
                  _ = loop_1
                  { var m rune = s.At((n+1))
                    if (i < l.Length()) { 
                      i = (i+1)
                      loop_1 = EID{C__INT,IVAL(i)}
                      } else {
                      loop_1 = Core.F_CALL(C_Compile_Cerror,ARGS(EID{MakeString("[210] not enough arguments in ~S").Id(),0},EID{self.Id(),0}))
                      } 
                    if ErrorIn(loop_1) {Result = loop_1
                    break
                    } else {
                    if (n > 1) { 
                      { var arg_2 *Language.Call
                        { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                          _CL_obj.Selector = C_princ
                          _CL_obj.Args = MakeConstantList((F_substring_string(s,1,(n-1))).Id())
                          arg_2 = _CL_obj
                          } 
                        r = r.AddFast(arg_2.Id())
                        } 
                      } 
                    var try_3 EID
                    { var arg_4 *ClaireAny
                      var try_5 EID
                      if ('A' == m) { 
                        { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                          _CL_obj.Selector = C_princ
                          _CL_obj.Args = MakeConstantList(l.At(i-1))
                          try_5 = EID{_CL_obj.Id(),0}
                          } 
                        }  else if ('S' == m) { 
                        { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                          _CL_obj.Selector = C_print
                          _CL_obj.Args = MakeConstantList(l.At(i-1))
                          try_5 = EID{_CL_obj.Id(),0}
                          } 
                        }  else if ('F' == m) { 
                        { var p_Z *ClaireBoolean = CFALSE
                          { var j int
                            var try_6 EID
                            { var arg_7 int
                              var try_8 EID
                              { var arg_9 rune
                                var try_10 EID
                                try_10 = Core.F_nth_get_string(s,(n+2),(n+2))
                                if ErrorIn(try_10) {try_8 = try_10
                                } else {
                                arg_9 = CHAR(try_10)
                                try_8 = EID{C__INT,IVAL(int(arg_9))}
                                }
                                } 
                              if ErrorIn(try_8) {try_6 = try_8
                              } else {
                              arg_7 = INT(try_8)
                              try_6 = EID{C__INT,IVAL((arg_7-48))}
                              }
                              } 
                            if ErrorIn(try_6) {try_5 = try_6
                            } else {
                            j = INT(try_6)
                            if ('%' == s.At((n+2))) { 
                              p_Z = CTRUE
                              j = 1
                              try_5 = EID{C__INT,IVAL(j)}
                              }  else if ((j < 0) || 
                                (j > 9)) { 
                              try_5 = ToException(Core.C_general_error.Make(MakeString("[189] F requires a single digit integer in ~S").Id(),MakeConstantList(self.Id()).Id())).Close()
                              } else {
                              try_5 = EID{CFALSE.Id(),0}
                              } 
                            if !ErrorIn(try_5) {
                            if ((p_Z != CTRUE) && 
                                ('%' == s.At((n+3)))) { 
                              p_Z = CTRUE
                              n = (n+1)
                              } 
                            n = (n+1)
                            { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                              _CL_obj.Selector = Core.C_mClaire_printFDigit
                              { 
                                var va_arg1 *Language.Call
                                var va_arg2 *ClaireList
                                va_arg1 = _CL_obj
                                { 
                                  var v_bag_arg *ClaireAny
                                  va_arg2= ToType(CEMPTY.Id()).EmptyList()
                                  if (p_Z == CTRUE) { 
                                    { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                      _CL_obj.Selector = ToProperty(C__star.Id())
                                      _CL_obj.Args = MakeConstantList(l.At(i-1),MakeFloat(100).Id())
                                      v_bag_arg = _CL_obj.Id()
                                      } 
                                    } else {
                                    v_bag_arg = l.At(i-1)
                                    } 
                                  va_arg2.AddFast(v_bag_arg)
                                  va_arg2.AddFast(MakeInteger(j).Id())} 
                                va_arg1.Args = va_arg2
                                } 
                              try_5 = EID{_CL_obj.Id(),0}
                              } 
                            }
                            }
                            } 
                          } 
                        }  else if ('I' == m) { 
                        try_5 = l.At(i-1).ToEID()
                        } else {
                        try_5 = EID{CFALSE.Id(),0}
                        } 
                      if ErrorIn(try_5) {try_3 = try_5
                      } else {
                      arg_4 = ANY(try_5)
                      try_3 = EID{r.AddFast(arg_4).Id(),0}
                      }
                      } 
                    if ErrorIn(try_3) {loop_1 = try_3
                    Result = try_3
                    break
                    } else {
                    r = ToList(OBJ(try_3))
                    loop_1 = EID{r.Id(),0}
                    s = F_substring_string(s,(n+2),1000)
                    n = F_get_string(s,'~')
                    loop_1 = EID{C__INT,IVAL(n)}
                    }}
                    } 
                  if ErrorIn(loop_1) {Result = loop_1
                  break
                  } else {
                  } 
                }
                if !ErrorIn(Result) {
                if (F_length_string(s) > 0) { 
                  { var arg_11 *Language.Call
                    { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                      _CL_obj.Selector = C_princ
                      _CL_obj.Args = MakeConstantList((s).Id())
                      arg_11 = _CL_obj
                      } 
                    r = r.AddFast(arg_11.Id())
                    } 
                  } 
                { var arg_12 *Language.Do
                  { var _CL_obj *Language.Do = Language.To_Do(new(Language.Do).Is(Language.C_Do))
                    _CL_obj.Args = r
                    arg_12 = _CL_obj
                    } 
                  Result = Core.F_CALL(C_c_code,ARGS(EID{arg_12.Id(),0},EID{C_any.Id(),0}))
                  } 
                }
                } 
              } 
            } 
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: c_code @ Printf (throw: true) 
func E_c_code_Printf (self EID) EID { 
    return F_c_code_Printf(Language.To_Printf(OBJ(self)) )} 
  
/* The go function for: c_type(self:Error) [status=0] */
func F_c_type_Error (self *Language.Error) *ClaireType { 
    return  ToType(CEMPTY.Id())
    } 
  
// The EID go function for: c_type @ Error (throw: false) 
func E_c_type_Error (self EID) EID { 
    return EID{F_c_type_Error(Language.To_Error(OBJ(self)) ).Id(),0}} 
  
/* The go function for: c_code(self:Error) [status=1] */
func F_c_code_Error (self *Language.Error) EID { 
    var Result EID
    { var arg_1 *Language.Call
      var try_2 EID
      { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
        _CL_obj.Selector = C_close
        { 
          var va_arg1 *Language.Call
          var va_arg2 *ClaireList
          va_arg1 = _CL_obj
          var try_3 EID
          { 
            var v_bag_arg *ClaireAny
            try_3= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
            var try_4 EID
            { var _CL_obj *Language.Cast = Language.To_Cast(new(Language.Cast).Is(Language.C_Cast))
              { 
                var va_arg1 *Language.Cast
                var va_arg2 *ClaireAny
                va_arg1 = _CL_obj
                var try_5 EID
                { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                  _CL_obj.Selector = C_Compile_anyObject_I
                  { 
                    var va_arg1 *Language.Call
                    var va_arg2 *ClaireList
                    va_arg1 = _CL_obj
                    var try_6 EID
                    { 
                      var v_bag_arg *ClaireAny
                      try_6= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                      ToList(OBJ(try_6)).AddFast(Core.C_general_error.Id())
                      var try_7 EID
                      { var arg_8 *ClaireAny
                        var try_9 EID
                        try_9 = Core.F_car_list(self.Args)
                        if ErrorIn(try_9) {try_7 = try_9
                        } else {
                        arg_8 = ANY(try_9)
                        try_7 = Core.F_CALL(C_c_code,ARGS(arg_8.ToEID(),EID{C_any.Id(),0}))
                        }
                        } 
                      if ErrorIn(try_7) {try_6 = try_7
                      } else {
                      v_bag_arg = ANY(try_7)
                      ToList(OBJ(try_6)).AddFast(v_bag_arg)
                      var try_10 EID
                      { var arg_11 *ClaireObject
                        var try_12 EID
                        if (self.Args.Length() != 1) { 
                          { var _CL_obj *Language.List = Language.To_List(new(Language.List).Is(Language.C_List))
                            { 
                              var va_arg1 *Language.Construct
                              var va_arg2 *ClaireList
                              va_arg1 = Language.To_Construct(_CL_obj.Id())
                              var try_13 EID
                              try_13 = self.Args.Cdr()
                              if ErrorIn(try_13) {try_12 = try_13
                              } else {
                              va_arg2 = ToList(OBJ(try_13))
                              va_arg1.Args = va_arg2
                              try_12 = EID{va_arg2.Id(),0}
                              }
                              } 
                            if !ErrorIn(try_12) {
                            try_12 = EID{_CL_obj.Id(),0}
                            }
                            } 
                          } else {
                          try_12 = EID{CNIL.Id(),0}
                          } 
                        if ErrorIn(try_12) {try_10 = try_12
                        } else {
                        arg_11 = ToObject(OBJ(try_12))
                        try_10 = Core.F_CALL(C_c_code,ARGS(EID{arg_11.Id(),0},EID{C_any.Id(),0}))
                        }
                        } 
                      if ErrorIn(try_10) {try_6 = try_10
                      } else {
                      v_bag_arg = ANY(try_10)
                      ToList(OBJ(try_6)).AddFast(v_bag_arg)}}
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
                if ErrorIn(try_5) {try_4 = try_5
                } else {
                va_arg2 = ANY(try_5)
                va_arg1.Arg = va_arg2
                try_4 = va_arg2.ToEID()
                }
                } 
              if !ErrorIn(try_4) {
              _CL_obj.SetArg = ToType(C_exception.Id())
              try_4 = EID{_CL_obj.Id(),0}
              }
              } 
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
      if ErrorIn(try_2) {Result = try_2
      } else {
      arg_1 = Language.To_Call(OBJ(try_2))
      Result = Core.F_CALL(C_c_code,ARGS(EID{arg_1.Id(),0},EID{C_void.Id(),0}))
      }
      } 
    return Result} 
  
// The EID go function for: c_code @ Error (throw: true) 
func E_c_code_Error (self EID) EID { 
    return F_c_code_Error(Language.To_Error(OBJ(self)) )} 
  
// *********************************************************************
// *     Part 3: If, Case, Do, Let                                     *
// *********************************************************************
//_______________ IF __________________________________________
// check if the test is of the form known?(v) so that the type (result) can be reduced
/* The go function for: extendedTest?(self:If) [status=0] */
func F_Optimize_extendedTest_ask_If (self *Language.If) *ClaireType { 
    var Result *ClaireType
    { var _Zt *ClaireAny = self.Test
      if (_Zt.Isa.IsIn(Language.C_Call) == CTRUE) { 
        { var g0251 *Language.Call = Language.To_Call(_Zt)
          if ((g0251.Args.At(0).Isa.IsIn(C_Variable) == CTRUE) && 
              (g0251.Selector.Id() == Core.C_known_ask.Id())) { 
            Result = ToType(OBJ(Core.F_CALL(C_range,ARGS(g0251.Args.At(0).ToEID()))))
            } else {
            Result = ToType(C_any.Id())
            } 
          } 
        } else {
        Result = ToType(C_any.Id())
        } 
      } 
    return Result} 
  
// The EID go function for: extendedTest? @ If (throw: false) 
func E_Optimize_extendedTest_ask_If (self EID) EID { 
    return EID{F_Optimize_extendedTest_ask_If(Language.To_If(OBJ(self)) ).Id(),0}} 
  
// notice that we analyze the test to detect the know? filter
/* The go function for: c_type(self:If) [status=1] */
func F_c_type_If (self *Language.If) EID { 
    var Result EID
    { var _Zr *ClaireType = F_Optimize_extendedTest_ask_If(self)
      var g0253I *ClaireBoolean
      var try_1 EID
      try_1 = F_Optimize_extended_ask_type(_Zr)
      if ErrorIn(try_1) {Result = try_1
      } else {
      g0253I = ToBoolean(OBJ(try_1))
      if (g0253I == CTRUE) { 
        F_Optimize_range_sets_any(ToList(OBJ(Core.F_CALL(C_args,ARGS(self.Test.ToEID())))).At(0),F_Optimize_sort_abstract_I_type(ToType(OBJ(Core.F_CALL(C_mClaire_t1,ARGS(EID{_Zr.Id(),0}))))))
        Result = EVOID
        } else {
        Result = EID{CFALSE.Id(),0}
        } 
      }
      if !ErrorIn(Result) {
      { var result *ClaireType
        var try_2 EID
        { var arg_3 *ClaireType
          var try_5 EID
          try_5 = Core.F_CALL(C_c_type,ARGS(self.Arg.ToEID()))
          if ErrorIn(try_5) {try_2 = try_5
          } else {
          arg_3 = ToType(OBJ(try_5))
          { var arg_4 *ClaireType
            var try_6 EID
            try_6 = Core.F_CALL(C_c_type,ARGS(self.Other.ToEID()))
            if ErrorIn(try_6) {try_2 = try_6
            } else {
            arg_4 = ToType(OBJ(try_6))
            try_2 = EID{Core.F_U_type(arg_3,arg_4).Id(),0}
            }
            } 
          }
          } 
        if ErrorIn(try_2) {Result = try_2
        } else {
        result = ToType(OBJ(try_2))
        var g0254I *ClaireBoolean
        var try_7 EID
        try_7 = F_Optimize_extended_ask_type(_Zr)
        if ErrorIn(try_7) {Result = try_7
        } else {
        g0254I = ToBoolean(OBJ(try_7))
        if (g0254I == CTRUE) { 
          Result = Core.F_put_property2(C_range,ToObject(ToList(OBJ(Core.F_CALL(C_args,ARGS(self.Test.ToEID())))).At(0)),_Zr.Id())
          } else {
          Result = EID{CFALSE.Id(),0}
          } 
        }
        if !ErrorIn(Result) {
        Result = EID{result.Id(),0}
        }
        }
        } 
      }
      } 
    return Result} 
  
// The EID go function for: c_type @ If (throw: true) 
func E_c_type_If (self EID) EID { 
    return F_c_type_If(Language.To_If(OBJ(self)) )} 
  
// debug boolean variable to flag the use of extented X U {unknown} + test : if known?(x)
/* The go function for: c_code(self:If,s:class) [status=1] */
func F_c_code_If (self *Language.If,s *ClaireClass) EID { 
    var Result EID
    { var _Zr *ClaireType = F_Optimize_extendedTest_ask_If(self)
      var g0255I *ClaireBoolean
      var try_1 EID
      try_1 = F_Optimize_extended_ask_type(_Zr)
      if ErrorIn(try_1) {Result = try_1
      } else {
      g0255I = ToBoolean(OBJ(try_1))
      if (g0255I == CTRUE) { 
        F_Optimize_range_sets_any(ToList(OBJ(Core.F_CALL(C_args,ARGS(self.Test.ToEID())))).At(0),F_Optimize_sort_abstract_I_type(ToType(OBJ(Core.F_CALL(C_mClaire_t1,ARGS(EID{_Zr.Id(),0}))))))
        Result = EVOID
        } else {
        Result = EID{CFALSE.Id(),0}
        } 
      }
      if !ErrorIn(Result) {
      var g0256I *ClaireBoolean
      var try_2 EID
      { 
        var v_and3 *ClaireBoolean
        
        var try_3 EID
        { var arg_4 *ClaireBoolean
          var try_5 EID
          { var arg_6 *ClaireType
            var try_7 EID
            { var arg_8 *ClaireType
              var try_9 EID
              try_9 = Core.F_CALL(C_c_type,ARGS(self.Test.ToEID()))
              if ErrorIn(try_9) {try_7 = try_9
              } else {
              arg_8 = ToType(OBJ(try_9))
              try_7 = EID{F_Optimize_ptype_type(arg_8).Id(),0}
              }
              } 
            if ErrorIn(try_7) {try_5 = try_7
            } else {
            arg_6 = ToType(OBJ(try_7))
            try_5 = EID{arg_6.Included(ToType(C_boolean.Id())).Id(),0}
            }
            } 
          if ErrorIn(try_5) {try_3 = try_5
          } else {
          arg_4 = ToBoolean(OBJ(try_5))
          try_3 = EID{arg_4.Not.Id(),0}
          }
          } 
        if ErrorIn(try_3) {try_2 = try_3
        } else {
        v_and3 = ToBoolean(OBJ(try_3))
        if (v_and3 == CFALSE) {try_2 = EID{CFALSE.Id(),0}
        } else { 
          v_and3 = Equal(C_PENIBLE.Value,CTRUE.Id())
          if (v_and3 == CFALSE) {try_2 = EID{CFALSE.Id(),0}
          } else { 
            try_2 = EID{CTRUE.Id(),0}} 
          } 
        }
        } 
      if ErrorIn(try_2) {Result = try_2
      } else {
      g0256I = ToBoolean(OBJ(try_2))
      if (g0256I == CTRUE) { 
        F_Compile_warn_void()
        Result = Core.F_tformat_string(MakeString("CLAIRE 3.3 SYNTAX - Test in ~S should be a boolean [260]\n"),1,MakeConstantList(self.Id()))
        } else {
        Result = EID{CFALSE.Id(),0}
        } 
      }
      if !ErrorIn(Result) {
      { var result *Language.If
        var try_10 EID
        { var arg_11 *ClaireAny
          var try_14 EID
          try_14 = F_Optimize_c_boolean_any(self.Test)
          if ErrorIn(try_14) {try_10 = try_14
          } else {
          arg_11 = ANY(try_14)
          { var arg_12 *ClaireAny
            var try_15 EID
            try_15 = Core.F_CALL(C_c_code,ARGS(self.Arg.ToEID(),EID{s.Id(),0}))
            if ErrorIn(try_15) {try_10 = try_15
            } else {
            arg_12 = ANY(try_15)
            { var arg_13 *ClaireAny
              var try_16 EID
              try_16 = Core.F_CALL(C_c_code,ARGS(self.Other.ToEID(),EID{s.Id(),0}))
              if ErrorIn(try_16) {try_10 = try_16
              } else {
              arg_13 = ANY(try_16)
              try_10 = Language.C_If.Make(arg_11,arg_12,arg_13).ToEID()
              }
              } 
            }
            } 
          }
          } 
        if ErrorIn(try_10) {Result = try_10
        } else {
        result = Language.To_If(OBJ(try_10))
        var g0257I *ClaireBoolean
        var try_17 EID
        try_17 = F_Optimize_extended_ask_type(_Zr)
        if ErrorIn(try_17) {Result = try_17
        } else {
        g0257I = ToBoolean(OBJ(try_17))
        if (g0257I == CTRUE) { 
          Result = Core.F_put_property2(C_range,ToObject(ToList(OBJ(Core.F_CALL(C_args,ARGS(self.Test.ToEID())))).At(0)),_Zr.Id())
          } else {
          Result = EID{CFALSE.Id(),0}
          } 
        }
        if !ErrorIn(Result) {
        Result = EID{result.Id(),0}
        }
        }
        } 
      }}
      } 
    return Result} 
  
// The EID go function for: c_code @ If (throw: true) 
func E_c_code_If (self EID,s EID) EID { 
    return F_c_code_If(Language.To_If(OBJ(self)),ToClass(OBJ(s)) )} 
  
// ------------------ CASE -------------------------------------------
// a member-of is a CLAIRE case. [yc 1/29/98]
// note that type inference supposes that the case is "closed" (all types are delt with)
// but only with safety >= 5
/* The go function for: c_type(self:Case) [status=1] */
func F_c_type_Case (self *Language.Case) EID { 
    var Result EID
    { var _Zvar *ClaireAny = self.ClaireVar
      { var _Ztype *ClaireAny
        if (_Zvar.Isa.IsIn(C_Variable) == CTRUE) { 
          { var g0258 *ClaireVariable = To_Variable(_Zvar)
            _Ztype = Core.F_get_property(C_range,ToObject(g0258.Id()))
            } 
          } else {
          _Ztype = C_any.Id()
          } 
        { var l *ClaireList = self.Args.Copy()
          { var rtype *ClaireType = ToType(CEMPTY.Id())
            { var utype *ClaireType = ToType(CEMPTY.Id())
              Result= EID{CFALSE.Id(),0}
              for (l.Length() > 0) { 
                var loop_1 EID
                _ = loop_1
                { 
                if (l.At(0).Isa.IsIn(C_type) == CTRUE) { 
                  utype = Core.F_U_type(utype,ToType(l.At(0)))
                  if (F_Compile_osort_any(_Ztype).Id() == F_Compile_osort_any(l.At(0)).Id()) { 
                    F_Optimize_range_sets_any(_Zvar,ToType(l.At(0)))
                    loop_1 = EVOID
                    }  else if (F_Compile_osort_any(_Ztype).Id() == C_any.Id()) { 
                    F_Optimize_range_sets_any(_Zvar,F_Optimize_sort_abstract_I_type(ToType(l.At(0))))
                    loop_1 = EVOID
                    } else {
                    loop_1 = EID{CFALSE.Id(),0}
                    } 
                  } else {
                  { var arg_2 *ClaireAny
                    var try_3 EID
                    try_3 = Core.F_car_list(l)
                    if ErrorIn(try_3) {loop_1 = try_3
                    } else {
                    arg_2 = ANY(try_3)
                    loop_1 = Core.F_CALL(C_Compile_Cerror,ARGS(EID{MakeString("[208] wrong type declaration for case: ~S in ~S").Id(),0},arg_2.ToEID(),EID{self.Id(),0}))
                    }
                    } 
                  } 
                if ErrorIn(loop_1) {Result = loop_1
                break
                } else {
                var try_4 EID
                { var arg_5 *ClaireType
                  var try_6 EID
                  try_6 = Core.F_CALL(C_c_type,ARGS(l.At(1).ToEID()))
                  if ErrorIn(try_6) {try_4 = try_6
                  } else {
                  arg_5 = ToType(OBJ(try_6))
                  try_4 = EID{Core.F_U_type(rtype,arg_5).Id(),0}
                  }
                  } 
                if ErrorIn(try_4) {loop_1 = try_4
                Result = try_4
                break
                } else {
                rtype = ToType(OBJ(try_4))
                loop_1 = EID{rtype.Id(),0}
                
                if (_Zvar.Isa.IsIn(C_Variable) == CTRUE) { 
                  { var g0260 *ClaireVariable = To_Variable(_Zvar)
                    g0260.Range = ToType(_Ztype)
                    } 
                  } 
                l = l.Skip(2)
                }}
                } 
              }
              if !ErrorIn(Result) {
              if (ToType(_Ztype).Included(utype) == CTRUE) { 
                Result = EID{rtype.Id(),0}
                }  else if (rtype.Included(ToType(C_boolean.Id())) == CTRUE) { 
                Result = EID{C_boolean.Id(),0}
                } else {
                Result = EID{C_any.Id(),0}
                } 
              }
              } 
            } 
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: c_type @ Case (throw: true) 
func E_c_type_Case (self EID) EID { 
    return F_c_type_Case(Language.To_Case(OBJ(self)) )} 
  
// safety
// utility : create a branch with substituted variable
// notice the use of occurence: create a let only if necessary :)
// ugly fix : use -1 in inded to indicate no type inference complain ....
/* The go function for: case_branch(x:any,%var:any,%type:type) [status=0] */
func F_Optimize_case_branch_any (x *ClaireAny,_Zvar *ClaireAny,_Ztype *ClaireType) *ClaireAny { 
    var Result *ClaireAny
    if (_Zvar.Isa.IsIn(C_Variable) == CTRUE) { 
      { var g0261 *ClaireVariable = To_Variable(_Zvar)
        { var vsub *ClaireVariable = F_Compile_Variable_I_symbol(Core.F_gensym_void(),-1,_Ztype.Id())
          if ((Equal(_Ztype.Id(),g0261.Range.Id()) != CTRUE) && 
              ((_Ztype.Id() != C_any.Id()) && 
                (Language.F_occurrence_any(x,g0261) > 0))) { 
            { var _CL_obj *Language.Let = Language.To_Let(new(Language.Let).Is(Language.C_Let))
              _CL_obj.ClaireVar = vsub
              _CL_obj.Value = g0261.Id()
              _CL_obj.Arg = F_Optimize_case_substitution_any(x,g0261,vsub)
              Result = _CL_obj.Id()
              } 
            } else {
            Result = x
            } 
          } 
        } 
      } else {
      Result = x
      } 
    return Result} 
  
// The EID go function for: case_branch @ any (throw: false) 
func E_Optimize_case_branch_any (x EID,_Zvar EID,_Ztype EID) EID { 
    return F_Optimize_case_branch_any(ANY(x),ANY(_Zvar),ToType(OBJ(_Ztype)) ).ToEID()} 
  
// this gets tricky, if the variable is changed in the case_branch, we need to update the original variable
// we add a copy to be able to compile the compiler (do not change the original code)
/* The go function for: case_substitution(x:any,%var:Variable,vsub:Variable) [status=0] */
func F_Optimize_case_substitution_any (x *ClaireAny,_Zvar *ClaireVariable,vsub *ClaireVariable) *ClaireAny { 
    var Result *ClaireAny
    { var y *ClaireAny = Language.F_substitution_any(Language.F_instruction_copy_any(x),_Zvar,vsub.Id())
      if (Language.F_occurchange_any(y,vsub) == CTRUE) { 
        { var _CL_obj *Language.Do = Language.To_Do(new(Language.Do).Is(Language.C_Do))
          _CL_obj.Args = MakeConstantList(y,Language.C_Assign.Make(_Zvar.Id(),vsub.Id()))
          Result = _CL_obj.Id()
          } 
        } else {
        Result = y
        } 
      } 
    return Result} 
  
// The EID go function for: case_substitution @ any (throw: false) 
func E_Optimize_case_substitution_any (x EID,_Zvar EID,vsub EID) EID { 
    return F_Optimize_case_substitution_any(ANY(x),To_Variable(OBJ(_Zvar)),To_Variable(OBJ(vsub)) ).ToEID()} 
  
// case is treated like a macro and vanishes into a large if.
// the last line is a trap for code generated by the logic compiler.
// in CLAIRE 4 we substitute the variables in the branches with a properly typed variable (borrowed from c_code@For)
// note: the range sets are now useless and have been removed
/* The go function for: c_code(self:Case,s:class) [status=1] */
func F_c_code_Case (self *Language.Case,s *ClaireClass) EID { 
    var Result EID
    { var _Zvar *ClaireAny = self.ClaireVar
      { var _Ztype *ClaireAny
        if (_Zvar.Isa.IsIn(C_Variable) == CTRUE) { 
          { var g0263 *ClaireVariable = To_Variable(_Zvar)
            _Ztype = Core.F_get_property(C_range,ToObject(g0263.Id()))
            } 
          } else {
          _Ztype = C_any.Id()
          } 
        { var l *ClaireList = self.Args.Copy()
          { var utype *ClaireAny = CEMPTY.Id()
            { var ctest1 *ClaireAny
              var try_1 EID
              { var arg_2 *Language.Call
                { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                  _CL_obj.Selector = ToProperty(C__Z.Id())
                  _CL_obj.Args = MakeConstantList(_Zvar,l.At(0))
                  arg_2 = _CL_obj
                  } 
                try_1 = F_Optimize_c_boolean_any(arg_2.Id())
                } 
              if ErrorIn(try_1) {Result = try_1
              } else {
              ctest1 = ANY(try_1)
              { var rep *Language.If
                var try_3 EID
                { var arg_4 *ClaireAny
                  var try_6 EID
                  try_6 = Core.F_CALL(C_c_code,ARGS(F_Optimize_case_branch_any(l.At(1),_Zvar,ToType(l.At(0))).ToEID(),EID{s.Id(),0}))
                  if ErrorIn(try_6) {try_3 = try_6
                  } else {
                  arg_4 = ANY(try_6)
                  { var arg_5 *ClaireAny
                    var try_7 EID
                    try_7 = Core.F_CALL(C_c_code,ARGS(EID{CFALSE.Id(),0},EID{s.Id(),0}))
                    if ErrorIn(try_7) {try_3 = try_7
                    } else {
                    arg_5 = ANY(try_7)
                    try_3 = Language.C_If.Make(ctest1,arg_4,arg_5).ToEID()
                    }
                    } 
                  }
                  } 
                if ErrorIn(try_3) {Result = try_3
                } else {
                rep = Language.To_If(OBJ(try_3))
                { var pointer *Language.If = rep
                  l = l.Skip(2)
                  Result= EID{CFALSE.Id(),0}
                  for (l.Length() > 0) { 
                    var loop_8 EID
                    _ = loop_8
                    { 
                    utype = Core.F_U_type(ToType(utype),ToType(l.At(0))).Id()
                    if (ToType(_Ztype).Included(ToType(utype)) == CTRUE) { 
                      { 
                        var va_arg1 *Language.If
                        var va_arg2 *ClaireAny
                        va_arg1 = pointer
                        var try_9 EID
                        try_9 = Core.F_CALL(C_c_code,ARGS(F_Optimize_case_branch_any(l.At(1),_Zvar,ToType(l.At(0))).ToEID(),EID{s.Id(),0}))
                        if ErrorIn(try_9) {loop_8 = try_9
                        } else {
                        va_arg2 = ANY(try_9)
                        va_arg1.Other = va_arg2
                        loop_8 = va_arg2.ToEID()
                        }
                        } 
                      if ErrorIn(loop_8) {Result = loop_8
                      break
                      } else {
                      Result = EID{CTRUE.Id(),0}
                      break
                      }
                      } else {
                      { var ctest *ClaireAny
                        var try_10 EID
                        { var arg_11 *Language.Call
                          { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                            _CL_obj.Selector = ToProperty(C__Z.Id())
                            _CL_obj.Args = MakeConstantList(_Zvar,l.At(0))
                            arg_11 = _CL_obj
                            } 
                          try_10 = F_Optimize_c_boolean_any(arg_11.Id())
                          } 
                        if ErrorIn(try_10) {loop_8 = try_10
                        } else {
                        ctest = ANY(try_10)
                        { 
                          var va_arg1 *Language.If
                          var va_arg2 *ClaireAny
                          va_arg1 = pointer
                          var try_12 EID
                          { var arg_13 *ClaireAny
                            var try_15 EID
                            try_15 = Core.F_CALL(C_c_code,ARGS(F_Optimize_case_branch_any(l.At(1),_Zvar,ToType(l.At(0))).ToEID(),EID{s.Id(),0}))
                            if ErrorIn(try_15) {try_12 = try_15
                            } else {
                            arg_13 = ANY(try_15)
                            { var arg_14 *ClaireAny
                              var try_16 EID
                              try_16 = Core.F_CALL(C_c_code,ARGS(EID{CFALSE.Id(),0},EID{s.Id(),0}))
                              if ErrorIn(try_16) {try_12 = try_16
                              } else {
                              arg_14 = ANY(try_16)
                              try_12 = Language.C_If.Make(ctest,arg_13,arg_14).ToEID()
                              }
                              } 
                            }
                            } 
                          if ErrorIn(try_12) {loop_8 = try_12
                          } else {
                          va_arg2 = ANY(try_12)
                          va_arg1.Other = va_arg2
                          loop_8 = va_arg2.ToEID()
                          }
                          } 
                        if ErrorIn(loop_8) {Result = loop_8
                        break
                        } else {
                        pointer = Language.To_If(pointer.Other)
                        loop_8 = EID{pointer.Id(),0}
                        }
                        }
                        } 
                      } 
                    if ErrorIn(loop_8) {Result = loop_8
                    break
                    } else {
                    l = l.Skip(2)
                    }
                    } 
                  }
                  if !ErrorIn(Result) {
                  var g0266I *ClaireBoolean
                  if (_Zvar.Isa.IsIn(Language.C_Definition) == CTRUE) { 
                    { var g0265 *Language.Definition = Language.To_Definition(_Zvar)
                      g0266I = g0265.Arg.Isa.IsIn(C_exception)
                      } 
                    } else {
                    g0266I = CFALSE
                    } 
                  if (g0266I == CTRUE) { 
                    Result = _Zvar.ToEID()
                    } else {
                    Result = EID{rep.Id(),0}
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
  
// The EID go function for: c_code @ Case (throw: true) 
func E_c_code_Case (self EID,s EID) EID { 
    return F_c_code_Case(Language.To_Case(OBJ(self)),ToClass(OBJ(s)) )} 
  
// member_of is treated like a macro and vanishes into a large if.
//_____________________ Block structure________________________
/* The go function for: c_type(self:Do) [status=1] */
func F_c_type_Do (self *Language.Do) EID { 
    var Result EID
    { var arg_1 *ClaireAny
      var try_2 EID
      try_2 = Core.F_last_list(self.Args)
      if ErrorIn(try_2) {Result = try_2
      } else {
      arg_1 = ANY(try_2)
      Result = Core.F_CALL(C_c_type,ARGS(arg_1.ToEID()))
      }
      } 
    return Result} 
  
// The EID go function for: c_type @ Do (throw: true) 
func E_c_type_Do (self EID) EID { 
    return F_c_type_Do(Language.To_Do(OBJ(self)) )} 
  
/* The go function for: c_code(self:Do,s:class) [status=1] */
func F_c_code_Do (self *Language.Do,s *ClaireClass) EID { 
    var Result EID
    { var _CL_obj *Language.Do = Language.To_Do(new(Language.Do).Is(Language.C_Do))
      { 
        var va_arg1 *Language.Do
        var va_arg2 *ClaireList
        va_arg1 = _CL_obj
        var try_1 EID
        { var m int = self.Args.Length()
          { var n int = 0
            { 
              var v_list6 *ClaireList
              var x *ClaireAny
              var v_local6 *ClaireAny
              v_list6 = self.Args
              try_1 = EID{CreateList(ToType(CEMPTY.Id()),v_list6.Length()).Id(),0}
              for CLcount := 0; CLcount < v_list6.Length(); CLcount++{ 
                x = v_list6.At(CLcount)
                var try_2 EID
                n = (n+1)
                { var arg_3 *ClaireClass
                  if (n == m) { 
                    arg_3 = s
                    } else {
                    arg_3 = C_void
                    } 
                  try_2 = Core.F_CALL(C_c_code,ARGS(x.ToEID(),EID{arg_3.Id(),0}))
                  } 
                if ErrorIn(try_2) {try_1 = try_2
                break
                } else {
                }
                if ErrorIn(try_2) {try_1 = try_2
                break
                } else {
                v_local6 = ANY(try_2)
                ToList(OBJ(try_1)).PutAt(CLcount,v_local6)
                } 
              }
              } 
            } 
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
    return Result} 
  
// The EID go function for: c_code @ Do (throw: true) 
func E_c_code_Do (self EID,s EID) EID { 
    return F_c_code_Do(Language.To_Do(OBJ(self)),ToClass(OBJ(s)) )} 
  
// ----------------------- LET -----------------------------------
// we make a range inference
//
/* The go function for: c_type(self:Let) [status=1] */
func F_c_type_Let (self *Language.Let) EID { 
    var Result EID
    { var arg_1 *ClaireType
      var try_2 EID
      try_2 = Core.F_CALL(C_c_type,ARGS(self.Value.ToEID()))
      if ErrorIn(try_2) {Result = try_2
      } else {
      arg_1 = ToType(OBJ(try_2))
      Result = F_Optimize_range_infers_Variable(self.ClaireVar,arg_1)
      }
      } 
    if !ErrorIn(Result) {
    Result = Core.F_CALL(C_c_type,ARGS(self.Arg.ToEID()))
    }
    return Result} 
  
// The EID go function for: c_type @ Let (throw: true) 
func E_c_type_Let (self EID) EID { 
    return F_c_type_Let(Language.To_Let(OBJ(self)) )} 
  
// works also for Let+ / Let*
/* The go function for: c_code(self:Let,s:class) [status=1] */
func F_c_code_Let (self *Language.Let,s *ClaireClass) EID { 
    var Result EID
    { var _Zv *ClaireAny = self.Value
      { var _Ztype *ClaireType
        var try_1 EID
        { var arg_2 *ClaireType
          var try_3 EID
          try_3 = Core.F_CALL(C_c_type,ARGS(_Zv.ToEID()))
          if ErrorIn(try_3) {try_1 = try_3
          } else {
          arg_2 = ToType(OBJ(try_3))
          try_1 = EID{F_Optimize_ptype_type(arg_2).Id(),0}
          }
          } 
        if ErrorIn(try_1) {Result = try_1
        } else {
        _Ztype = ToType(OBJ(try_1))
        Result = F_Optimize_range_infers_Variable(self.ClaireVar,_Ztype)
        if !ErrorIn(Result) {
        if (_Ztype.Included(self.ClaireVar.Range) != CTRUE) { 
          var try_4 EID
          try_4 = F_Optimize_c_warn_Variable(self.ClaireVar,_Zv,_Ztype)
          if ErrorIn(try_4) {Result = try_4
          } else {
          _Zv = ANY(try_4)
          Result = _Zv.ToEID()
          }
          } else {
          Result = EID{CFALSE.Id(),0}
          } 
        if !ErrorIn(Result) {
        { var x *Language.Let
          var try_5 EID
          { var _CL_obj *Language.Let = Language.To_Let(new(Language.Let).Is(Language.C_Let))
            _CL_obj.ClaireVar = self.ClaireVar
            { 
              var va_arg1 *Language.Let
              var va_arg2 *ClaireAny
              va_arg1 = _CL_obj
              var try_6 EID
              try_6 = F_Compile_c_strict_code_any(_Zv,F_Compile_psort_any(self.ClaireVar.Range.Id()))
              if ErrorIn(try_6) {try_5 = try_6
              } else {
              va_arg2 = ANY(try_6)
              va_arg1.Value = va_arg2
              try_5 = va_arg2.ToEID()
              }
              } 
            if !ErrorIn(try_5) {
            { 
              var va_arg1 *Language.Let
              var va_arg2 *ClaireAny
              va_arg1 = _CL_obj
              var try_7 EID
              try_7 = Core.F_CALL(C_c_code,ARGS(self.Arg.ToEID(),EID{s.Id(),0}))
              if ErrorIn(try_7) {try_5 = try_7
              } else {
              va_arg2 = ANY(try_7)
              va_arg1.Arg = va_arg2
              try_5 = va_arg2.ToEID()
              }
              } 
            if !ErrorIn(try_5) {
            try_5 = EID{_CL_obj.Id(),0}
            }}
            } 
          if ErrorIn(try_5) {Result = try_5
          } else {
          x = Language.To_Let(OBJ(try_5))
          x.Isa = self.Isa
          Result = EID{x.Id(),0}
          }
          } 
        }}
        }
        } 
      } 
    return Result} 
  
// The EID go function for: c_code @ Let (throw: true) 
func E_c_code_Let (self EID,s EID) EID { 
    return F_c_code_Let(Language.To_Let(OBJ(self)),ToClass(OBJ(s)) )} 
  
// type inference for When is more subtle
/* The go function for: c_type(self:When) [status=1] */
func F_c_type_When (self *Language.When) EID { 
    var Result EID
    { var _Zv *ClaireAny = self.Value
      { var v *ClaireVariable = self.ClaireVar
        { var d *ClaireAny
          var try_1 EID
          try_1 = F_Optimize_daccess_any(_Zv,CTRUE)
          if ErrorIn(try_1) {Result = try_1
          } else {
          d = ANY(try_1)
          { var _Ztype *ClaireType
            var try_2 EID
            if (d != CNULL) { 
              try_2 = Core.F_CALL(C_c_type,ARGS(d.ToEID()))
              } else {
              try_2 = Core.F_CALL(C_c_type,ARGS(_Zv.ToEID()))
              } 
            if ErrorIn(try_2) {Result = try_2
            } else {
            _Ztype = ToType(OBJ(try_2))
            var g0267I *ClaireBoolean
            var try_3 EID
            try_3 = F_Optimize_extended_ask_type(_Ztype)
            if ErrorIn(try_3) {Result = try_3
            } else {
            g0267I = ToBoolean(OBJ(try_3))
            if (g0267I == CTRUE) { 
              _Ztype = ToType(OBJ(Core.F_CALL(C_mClaire_t1,ARGS(EID{_Ztype.Id(),0}))))
              Result = EID{_Ztype.Id(),0}
              } else {
              Result = EID{CFALSE.Id(),0}
              } 
            }
            if !ErrorIn(Result) {
            Result = F_Optimize_range_infers_Variable(v,_Ztype)
            if !ErrorIn(Result) {
            { var arg_4 *ClaireType
              var try_6 EID
              try_6 = Core.F_CALL(C_c_type,ARGS(self.Arg.ToEID()))
              if ErrorIn(try_6) {Result = try_6
              } else {
              arg_4 = ToType(OBJ(try_6))
              { var arg_5 *ClaireType
                var try_7 EID
                try_7 = Core.F_CALL(C_c_type,ARGS(self.Other.ToEID()))
                if ErrorIn(try_7) {Result = try_7
                } else {
                arg_5 = ToType(OBJ(try_7))
                Result = EID{Core.F_U_type(arg_4,arg_5).Id(),0}
                }
                } 
              }
              } 
            }}
            }
            } 
          }
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: c_type @ When (throw: true) 
func E_c_type_When (self EID) EID { 
    return F_c_type_When(Language.To_When(OBJ(self)) )} 
  
// A When is macroexpanded into one/two Let
/* The go function for: c_code(self:When,s:class) [status=1] */
func F_c_code_When (self *Language.When,s *ClaireClass) EID { 
    var Result EID
    { var _Zv *ClaireAny = self.Value
      { var v *ClaireVariable = self.ClaireVar
        { var d *ClaireAny
          var try_1 EID
          try_1 = F_Optimize_daccess_any(_Zv,CTRUE)
          if ErrorIn(try_1) {Result = try_1
          } else {
          d = ANY(try_1)
          { var v2 *ClaireVariable = F_Compile_Variable_I_symbol(F_append_symbol(self.ClaireVar.Pname,MakeString("test").Id()),self.ClaireVar.Index,C_any.Id())
            { var _Ztype *ClaireType
              var try_2 EID
              if (d != CNULL) { 
                try_2 = Core.F_CALL(C_c_type,ARGS(d.ToEID()))
                } else {
                try_2 = Core.F_CALL(C_c_type,ARGS(_Zv.ToEID()))
                } 
              if ErrorIn(try_2) {Result = try_2
              } else {
              _Ztype = ToType(OBJ(try_2))
              var g0268I *ClaireBoolean
              var try_3 EID
              try_3 = F_Optimize_extended_ask_type(_Ztype)
              if ErrorIn(try_3) {Result = try_3
              } else {
              g0268I = ToBoolean(OBJ(try_3))
              if (g0268I == CTRUE) { 
                _Ztype = ToType(OBJ(Core.F_CALL(C_mClaire_t1,ARGS(EID{_Ztype.Id(),0}))))
                Result = EID{_Ztype.Id(),0}
                } else {
                Result = EID{CFALSE.Id(),0}
                } 
              }
              if !ErrorIn(Result) {
              Result = F_Optimize_range_infers_Variable(v,_Ztype)
              if !ErrorIn(Result) {
              var g0269I *ClaireBoolean
              var try_4 EID
              { 
                var v_and7 *ClaireBoolean
                
                v_and7 = Core.F_known_ask_any(d)
                if (v_and7 == CFALSE) {try_4 = EID{CFALSE.Id(),0}
                } else { 
                  var try_5 EID
                  { var arg_6 *ClaireBoolean
                    var try_7 EID
                    try_7 = F_Optimize_extended_ask_type(ToType(OBJ(Core.F_CALL(C_range,ARGS(Core.F_CALL(C_selector,ARGS(d.ToEID())))))))
                    if ErrorIn(try_7) {try_5 = try_7
                    } else {
                    arg_6 = ToBoolean(OBJ(try_7))
                    try_5 = EID{arg_6.Not.Id(),0}
                    }
                    } 
                  if ErrorIn(try_5) {try_4 = try_5
                  } else {
                  v_and7 = ToBoolean(OBJ(try_5))
                  if (v_and7 == CFALSE) {try_4 = EID{CFALSE.Id(),0}
                  } else { 
                    try_4 = EID{CTRUE.Id(),0}} 
                  } 
                }
                } 
              if ErrorIn(try_4) {Result = try_4
              } else {
              g0269I = ToBoolean(OBJ(try_4))
              if (g0269I == CTRUE) { 
                { var _CL_obj *Language.Let = Language.To_Let(new(Language.Let).Is(Language.C_Let))
                  _CL_obj.ClaireVar = v
                  _CL_obj.Value = d
                  { 
                    var va_arg1 *Language.Let
                    var va_arg2 *ClaireAny
                    va_arg1 = _CL_obj
                    var try_8 EID
                    { var arg_9 *Language.CallMethod2
                      var try_12 EID
                      { var _CL_obj *Language.CallMethod2 = Language.To_CallMethod2(new(Language.CallMethod2).Is(Language.C_Call_method2))
                        _CL_obj.Arg = ToMethod(Core.F__at_property1(Core.C_identical_ask,C_any).Id())
                        { 
                          var va_arg1 *Language.CallMethod
                          var va_arg2 *ClaireList
                          va_arg1 = Language.To_CallMethod(_CL_obj.Id())
                          var try_13 EID
                          { 
                            var v_bag_arg *ClaireAny
                            try_13= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                            ToList(OBJ(try_13)).AddFast(v.Id())
                            var try_14 EID
                            { var arg_15 *ClaireAny
                              var try_16 EID
                              try_16 = Core.F_CALL(C_Compile_c_sort,ARGS(EID{v.Id(),0}))
                              if ErrorIn(try_16) {try_14 = try_16
                              } else {
                              arg_15 = ANY(try_16)
                              try_14 = Core.F_CALL(C_c_code,ARGS(EID{CNULL,0},arg_15.ToEID()))
                              }
                              } 
                            if ErrorIn(try_14) {try_13 = try_14
                            } else {
                            v_bag_arg = ANY(try_14)
                            ToList(OBJ(try_13)).AddFast(v_bag_arg)}
                            } 
                          if ErrorIn(try_13) {try_12 = try_13
                          } else {
                          va_arg2 = ToList(OBJ(try_13))
                          va_arg1.Args = va_arg2
                          try_12 = EID{va_arg2.Id(),0}
                          }
                          } 
                        if !ErrorIn(try_12) {
                        try_12 = EID{_CL_obj.Id(),0}
                        }
                        } 
                      if ErrorIn(try_12) {try_8 = try_12
                      } else {
                      arg_9 = Language.To_CallMethod2(OBJ(try_12))
                      { var arg_10 *ClaireAny
                        var try_17 EID
                        try_17 = Core.F_CALL(C_c_code,ARGS(self.Other.ToEID(),EID{s.Id(),0}))
                        if ErrorIn(try_17) {try_8 = try_17
                        } else {
                        arg_10 = ANY(try_17)
                        { var arg_11 *ClaireAny
                          var try_18 EID
                          try_18 = Core.F_CALL(C_c_code,ARGS(self.Arg.ToEID(),EID{s.Id(),0}))
                          if ErrorIn(try_18) {try_8 = try_18
                          } else {
                          arg_11 = ANY(try_18)
                          try_8 = Language.C_If.Make(arg_9.Id(),arg_10,arg_11).ToEID()
                          }
                          } 
                        }
                        } 
                      }
                      } 
                    if ErrorIn(try_8) {Result = try_8
                    } else {
                    va_arg2 = ANY(try_8)
                    va_arg1.Arg = va_arg2
                    Result = va_arg2.ToEID()
                    }
                    } 
                  if !ErrorIn(Result) {
                  Result = EID{_CL_obj.Id(),0}
                  }
                  } 
                } else {
                var g0270I *ClaireBoolean
                var try_19 EID
                { 
                  var v_and8 *ClaireBoolean
                  
                  var try_20 EID
                  { var arg_21 *ClaireAny
                    var try_22 EID
                    try_22 = Core.F_CALL(C_Compile_c_sort,ARGS(EID{v.Id(),0}))
                    if ErrorIn(try_22) {try_20 = try_22
                    } else {
                    arg_21 = ANY(try_22)
                    try_20 = EID{Equal(arg_21,C_any.Id()).Id(),0}
                    }
                    } 
                  if ErrorIn(try_20) {try_19 = try_20
                  } else {
                  v_and8 = ToBoolean(OBJ(try_20))
                  if (v_and8 == CFALSE) {try_19 = EID{CFALSE.Id(),0}
                  } else { 
                    v_and8 = _Ztype.Included(v.Range)
                    if (v_and8 == CFALSE) {try_19 = EID{CFALSE.Id(),0}
                    } else { 
                      v_and8 = F__sup_equal_integer(C_compiler.Safety,2)
                      if (v_and8 == CFALSE) {try_19 = EID{CFALSE.Id(),0}
                      } else { 
                        try_19 = EID{CTRUE.Id(),0}} 
                      } 
                    } 
                  }
                  } 
                if ErrorIn(try_19) {Result = try_19
                } else {
                g0270I = ToBoolean(OBJ(try_19))
                if (g0270I == CTRUE) { 
                  { var arg_23 *Language.Let
                    { var _CL_obj *Language.Let = Language.To_Let(new(Language.Let).Is(Language.C_Let))
                      _CL_obj.ClaireVar = v
                      _CL_obj.Value = _Zv
                      { 
                        var va_arg1 *Language.Let
                        var va_arg2 *ClaireAny
                        va_arg1 = _CL_obj
                        { var arg_24 *Language.Call
                          { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                            _CL_obj.Selector = ToProperty(Core.C__I_equal.Id())
                            _CL_obj.Args = MakeConstantList(v.Id(),CNULL)
                            arg_24 = _CL_obj
                            } 
                          va_arg2 = Language.C_If.Make(arg_24.Id(),self.Arg,self.Other)
                          } 
                        va_arg1.Arg = va_arg2
                        } 
                      arg_23 = _CL_obj
                      } 
                    Result = Core.F_CALL(C_c_code,ARGS(EID{arg_23.Id(),0},EID{s.Id(),0}))
                    } 
                  } else {
                  { var arg_25 *Language.Let
                    var try_26 EID
                    { var _CL_obj *Language.Let = Language.To_Let(new(Language.Let).Is(Language.C_Let))
                      _CL_obj.ClaireVar = v2
                      _CL_obj.Value = _Zv
                      { 
                        var va_arg1 *Language.Let
                        var va_arg2 *ClaireAny
                        va_arg1 = _CL_obj
                        var try_27 EID
                        { var arg_28 *Language.Call
                          { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                            _CL_obj.Selector = ToProperty(Core.C__I_equal.Id())
                            _CL_obj.Args = MakeConstantList(v2.Id(),CNULL)
                            arg_28 = _CL_obj
                            } 
                          { var arg_29 *Language.Let
                            { var _CL_obj *Language.Let = Language.To_Let(new(Language.Let).Is(Language.C_Let))
                              _CL_obj.ClaireVar = v
                              { 
                                var va_arg1 *Language.Let
                                var va_arg2 *ClaireAny
                                va_arg1 = _CL_obj
                                { var _CL_obj *Language.Cast = Language.To_Cast(new(Language.Cast).Is(Language.C_Cast))
                                  _CL_obj.Arg = v2.Id()
                                  _CL_obj.SetArg = _Ztype
                                  va_arg2 = _CL_obj.Id()
                                  } 
                                va_arg1.Value = va_arg2
                                } 
                              _CL_obj.Arg = self.Arg
                              arg_29 = _CL_obj
                              } 
                            { var arg_30 *ClaireAny
                              var try_31 EID
                              try_31 = Core.F_CALL(C_c_code,ARGS(self.Other.ToEID(),EID{s.Id(),0}))
                              if ErrorIn(try_31) {try_27 = try_31
                              } else {
                              arg_30 = ANY(try_31)
                              try_27 = Language.C_If.Make(arg_28.Id(),arg_29.Id(),arg_30).ToEID()
                              }
                              } 
                            } 
                          } 
                        if ErrorIn(try_27) {try_26 = try_27
                        } else {
                        va_arg2 = ANY(try_27)
                        va_arg1.Arg = va_arg2
                        try_26 = va_arg2.ToEID()
                        }
                        } 
                      if !ErrorIn(try_26) {
                      try_26 = EID{_CL_obj.Id(),0}
                      }
                      } 
                    if ErrorIn(try_26) {Result = try_26
                    } else {
                    arg_25 = Language.To_Let(OBJ(try_26))
                    Result = Core.F_CALL(C_c_code,ARGS(EID{arg_25.Id(),0},EID{s.Id(),0}))
                    }
                    } 
                  } 
                }
                } 
              }
              }}
              }
              } 
            } 
          }
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: c_code @ When (throw: true) 
func E_c_code_When (self EID,s EID) EID { 
    return F_c_code_When(Language.To_When(OBJ(self)),ToClass(OBJ(s)) )} 
  
// *********************************************************************
// *     Part 4: Loops                                                 *
// *********************************************************************
// here we could do a return extraction
/* The go function for: c_type(self:For) [status=1] */
func F_c_type_For (self *Language.For) EID { 
    var Result EID
    { var arg_1 *ClaireType
      var try_2 EID
      try_2 = F_Compile_return_type_any(self.Arg)
      if ErrorIn(try_2) {Result = try_2
      } else {
      arg_1 = ToType(OBJ(try_2))
      Result = EID{F_Optimize_infers_from_type(arg_1,self.Id()).Id(),0}
      }
      } 
    return Result} 
  
// The EID go function for: c_type @ For (throw: true) 
func E_c_type_For (self EID) EID { 
    return F_c_type_For(Language.To_For(OBJ(self)) )} 
  
/* The go function for: infers_from(t:type,self:any) [status=0] */
func F_Optimize_infers_from_type (t *ClaireType,self *ClaireAny) *ClaireType { 
    if (Equal(t.Id(),CEMPTY.Id()) == CTRUE) { 
      return  F_Optimize_sort_abstract_I_type(ToType(C_boolean.Id()))
      }  else if (C_compiler.Safety >= 2) { 
      Core.F_tformat_string(MakeString("... c_type(~S) -> ~S - ~S \n"),2,MakeConstantList(self,t.Id(),F_Optimize_sort_abstract_I_type(t).Id()))
      return  F_Optimize_sort_abstract_I_type(t)
      } else {
      return  ToType(C_any.Id())
      } 
    } 
  
// The EID go function for: infers_from @ type (throw: false) 
func E_Optimize_infers_from_type (t EID,self EID) EID { 
    return EID{F_Optimize_infers_from_type(ToType(OBJ(t)),ANY(self) ).Id(),0}} 
  
// false or the return value
// notice that for is of sort any and may require a cast ..
/* The go function for: c_code(self:For,s:class) [status=1] */
func F_c_code_For (self *Language.For,s *ClaireClass) EID { 
    var Result EID
    { var sx *ClaireAny = self.SetArg
      { var ns int = C_compiler.Safety
        { var vold *ClaireVariable = self.ClaireVar
          { var v *ClaireVariable = F_Compile_Variable_I_symbol(vold.Pname,vold.Index,Core.F_get_property(C_range,ToObject(vold.Id())))
            { var narg *ClaireAny = Language.F_substitution_any(self.Arg,vold,v.Id())
              { var scs *ClaireAny
                var try_1 EID
                try_1 = F_Optimize_c_inline_arg_ask_any(sx)
                if ErrorIn(try_1) {Result = try_1
                } else {
                scs = ANY(try_1)
                if (sx.Isa.IsIn(Core.C_global_variable) == CTRUE) { 
                  { var g0271 *Core.GlobalVariable = Core.ToGlobalVariable(sx)
                    if (F_boolean_I_any(g0271.Range.Id()).Id() != CTRUE.Id()) { 
                      self.SetArg = g0271.Value
                      g0271 = Core.ToGlobalVariable(g0271.Value)
                      } 
                    sx = g0271.Id()
                    } 
                  }  else if (C_class.Id() == sx.Isa.Id()) { 
                  { var g0272 *ClaireClass = ToClass(sx)
                    if ((g0272.Open <= 1) && 
                        (F_boolean_I_any(g0272.Subclass.Id()).Id() != CTRUE.Id())) { 
                      { 
                        var va_arg1 *Language.Iteration
                        var va_arg2 *ClaireAny
                        va_arg1 = Language.To_Iteration(self.Id())
                        { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                          _CL_obj.Selector = C_instances
                          _CL_obj.Args = MakeConstantList(g0272.Id())
                          va_arg2 = _CL_obj.Id()
                          } 
                        va_arg1.SetArg = va_arg2
                        } 
                      } 
                    } 
                  } 
                { var _Zt *ClaireType
                  var try_2 EID
                  try_2 = Core.F_CALL(C_c_type,ARGS(self.SetArg.ToEID()))
                  if ErrorIn(try_2) {Result = try_2
                  } else {
                  _Zt = ToType(OBJ(try_2))
                  { var _Zt2 *ClaireType = F_Optimize_pmember_type(_Zt)
                    { var _Zt3 *ClaireType
                      var try_3 EID
                      { var arg_4 *ClaireType
                        var try_5 EID
                        { var arg_6 *Language.Call
                          { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                            _CL_obj.Selector = C_set_I
                            _CL_obj.Args = MakeConstantList(sx)
                            arg_6 = _CL_obj
                            } 
                          try_5 = Core.F_CALL(C_c_type,ARGS(EID{arg_6.Id(),0}))
                          } 
                        if ErrorIn(try_5) {try_3 = try_5
                        } else {
                        arg_4 = ToType(OBJ(try_5))
                        try_3 = EID{F_Optimize_pmember_type(arg_4).Id(),0}
                        }
                        } 
                      if ErrorIn(try_3) {Result = try_3
                      } else {
                      _Zt3 = ToType(OBJ(try_3))
                      if (_Zt3.Included(_Zt2) == CTRUE) { 
                        _Zt2 = _Zt3
                        } 
                      F_Optimize_range_infers_for_Variable(v,_Zt2,_Zt)
                      C_compiler.Safety = 1
                      _Zt2 = v.Range
                      v.Range = F_Optimize_ptype_type(_Zt2)
                      { var m *ClaireAny = F_Optimize_Iterate_I_Iteration(Language.To_Iteration(self.Id()))
                        if (C_method.Id() != m.Isa.Id()) { 
                          { var m2 *ClaireAny = F_Optimize_restriction_I_property(Language.C_iterate,MakeConstantList(_Zt.Id(),MakeConstantSet(v.Id()).Id(),C_any.Id()),CTRUE)
                            if (C_method.Id() == m2.Isa.Id()) { 
                              { var g0273 *ClaireMethod = ToMethod(m2)
                                m = g0273.Id()
                                } 
                              } 
                            } 
                          } 
                        C_compiler.Safety = ns
                        v.Range = _Zt2
                        var g0276I *ClaireBoolean
                        if (C_method.Id() == m.Isa.Id()) { 
                          { var g0274 *ClaireMethod = ToMethod(m)
                            g0276I = g0274.Inline_ask
                            } 
                          } else {
                          g0276I = CFALSE
                          } 
                        if (g0276I == CTRUE) { 
                          
                          if (F_Optimize_sort_abstract_ask_type(v.Range) == CTRUE) { 
                            v.Range = To_Union(v.Range.Id()).T2
                            } 
                          Result = F_Optimize_c_inline_method1(ToMethod(m),MakeConstantList(Language.F_instruction_copy_any(self.SetArg),v.Id(),narg),s)
                          }  else if (F_boolean_I_any(scs) == CTRUE) { 
                          { var arg_7 *Language.For
                            { var _CL_obj *Language.For = Language.To_For(new(Language.For).Is(Language.C_For))
                              _CL_obj.ClaireVar = v
                              _CL_obj.SetArg = scs
                              _CL_obj.Arg = narg
                              arg_7 = _CL_obj
                              } 
                            Result = Core.F_CALL(C_c_code,ARGS(EID{arg_7.Id(),0},EID{s.Id(),0}))
                            } 
                          } else {
                          var g0277I *ClaireBoolean
                          if (sx.Isa.IsIn(Language.C_Call) == CTRUE) { 
                            { var g0275 *Language.Call = Language.To_Call(sx)
                              g0277I = Equal(g0275.Selector.Id(),Core.C_Id.Id())
                              } 
                            } else {
                            g0277I = CFALSE
                            } 
                          if (g0277I == CTRUE) { 
                            Result = F_Optimize_c_code_multiple_For(self,_Zt,s)
                            } else {
                            { var _CL_obj *Language.For = Language.To_For(new(Language.For).Is(Language.C_For))
                              _CL_obj.ClaireVar = v
                              { 
                                var va_arg1 *Language.Iteration
                                var va_arg2 *ClaireAny
                                va_arg1 = Language.To_Iteration(_CL_obj.Id())
                                var try_8 EID
                                try_8 = F_Optimize_enumerate_code_any(self.SetArg,_Zt)
                                if ErrorIn(try_8) {Result = try_8
                                } else {
                                va_arg2 = ANY(try_8)
                                va_arg1.SetArg = va_arg2
                                Result = va_arg2.ToEID()
                                }
                                } 
                              if !ErrorIn(Result) {
                              { 
                                var va_arg1 *Language.Iteration
                                var va_arg2 *ClaireAny
                                va_arg1 = Language.To_Iteration(_CL_obj.Id())
                                var try_9 EID
                                try_9 = Core.F_CALL(C_c_code,ARGS(narg.ToEID(),EID{C_void.Id(),0}))
                                if ErrorIn(try_9) {Result = try_9
                                } else {
                                va_arg2 = ANY(try_9)
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
        } 
      } 
    return Result} 
  
// The EID go function for: c_code @ For (throw: true) 
func E_c_code_For (self EID,s EID) EID { 
    return F_c_code_For(Language.To_For(OBJ(self)),ToClass(OBJ(s)) )} 
  
//             (if (s = any) r2 else to_C(arg = r2, set_arg = s)))))]
// new: we macroexpand the iteration  based on the type
// this is only called if the set is wrapped in an Id
/* The go function for: c_code_multiple(self:For,%t:type,s:class) [status=1] */
func F_Optimize_c_code_multiple_For (self *Language.For,_Zt *ClaireType,s *ClaireClass) EID { 
    var Result EID
    { var v *ClaireVariable = self.ClaireVar
      { var sx *ClaireAny = Language.To_Call(self.SetArg).Args.At(0)
        { var v2 *ClaireVariable = F_Compile_Variable_I_symbol(F_append_symbol(v.Pname,MakeString("test").Id()),self.ClaireVar.Index,_Zt.Id())
          { var n *Language.Let
            var try_1 EID
            { var _CL_obj *Language.Let = Language.To_Let(new(Language.Let).Is(Language.C_Let))
              _CL_obj.ClaireVar = v2
              _CL_obj.Value = sx
              { 
                var va_arg1 *Language.Let
                var va_arg2 *ClaireAny
                va_arg1 = _CL_obj
                var try_2 EID
                { var _CL_obj *Language.For = Language.To_For(new(Language.For).Is(Language.C_For))
                  _CL_obj.ClaireVar = self.ClaireVar
                  { 
                    var va_arg1 *Language.Iteration
                    var va_arg2 *ClaireAny
                    va_arg1 = Language.To_Iteration(_CL_obj.Id())
                    var try_3 EID
                    try_3 = F_Optimize_enumerate_code_any(sx,_Zt)
                    if ErrorIn(try_3) {try_2 = try_3
                    } else {
                    va_arg2 = ANY(try_3)
                    va_arg1.SetArg = va_arg2
                    try_2 = va_arg2.ToEID()
                    }
                    } 
                  if !ErrorIn(try_2) {
                  _CL_obj.Arg = self.Arg
                  try_2 = EID{_CL_obj.Id(),0}
                  }
                  } 
                if ErrorIn(try_2) {try_1 = try_2
                } else {
                va_arg2 = ANY(try_2)
                va_arg1.Arg = va_arg2
                try_1 = va_arg2.ToEID()
                }
                } 
              if !ErrorIn(try_1) {
              try_1 = EID{_CL_obj.Id(),0}
              }
              } 
            if ErrorIn(try_1) {Result = try_1
            } else {
            n = Language.To_Let(OBJ(try_1))
            Core.F_tformat_string(MakeString("---- note: use an expended iteration for {~S} \n"),0,MakeConstantList(self.Id()))
            { 
              var r *ClaireRestriction
              _ = r
              var r_iter *ClaireAny
              Result= EID{CFALSE.Id(),0}
              for _,r_iter = range(Language.C_iterate.Restrictions.ValuesO()){ 
                r = ToRestriction(r_iter)
                var loop_4 EID
                _ = loop_4
                var g0278I *ClaireBoolean
                var try_5 EID
                { 
                  var v_and8 *ClaireBoolean
                  
                  v_and8 = ToType(Core.F_domain_I_restriction(r).Id()).Included(_Zt)
                  if (v_and8 == CFALSE) {try_5 = EID{CFALSE.Id(),0}
                  } else { 
                    v_and8 = ToType(Core.F_domain_I_restriction(r).Id()).Included(ToType(C_collection.Id()))
                    if (v_and8 == CFALSE) {try_5 = EID{CFALSE.Id(),0}
                    } else { 
                      v_and8 = ToMethod(r.Id()).Inline_ask
                      if (v_and8 == CFALSE) {try_5 = EID{CFALSE.Id(),0}
                      } else { 
                        var try_6 EID
                        try_6 = Core.F_BELONG(v.Id(),r.Domain.ValuesO()[1])
                        if ErrorIn(try_6) {try_5 = try_6
                        } else {
                        v_and8 = ToBoolean(OBJ(try_6))
                        if (v_and8 == CFALSE) {try_5 = EID{CFALSE.Id(),0}
                        } else { 
                          try_5 = EID{CTRUE.Id(),0}} 
                        } 
                      } 
                    } 
                  }
                  } 
                if ErrorIn(try_5) {loop_4 = try_5
                } else {
                g0278I = ToBoolean(OBJ(try_5))
                if (g0278I == CTRUE) { 
                  { var vnew *ClaireVariable
                    { var _CL_obj *ClaireVariable = To_Variable(new(ClaireVariable).Is(C_Variable))
                      _CL_obj.Pname = v.Pname
                      _CL_obj.Range = v.Range
                      _CL_obj.Index = v.Index
                      vnew = _CL_obj
                      } 
                    { var narg *ClaireAny = Language.F_substitution_any(self.Arg,v,vnew.Id())
                      { 
                        var va_arg1 *Language.Let
                        var va_arg2 *ClaireAny
                        va_arg1 = n
                        var try_7 EID
                        { var arg_8 *Language.Call
                          { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                            _CL_obj.Selector = ToProperty(C__Z.Id())
                            _CL_obj.Args = MakeConstantList(v2.Id(),r.Domain.ValuesO()[0])
                            arg_8 = _CL_obj
                            } 
                          { var arg_9 *ClaireAny
                            var try_10 EID
                            if (F_Optimize_sort_abstract_ask_type(vnew.Range) == CTRUE) { 
                              vnew.Range = To_Union(v.Range.Id()).T2
                              } 
                            try_10 = F_Optimize_c_inline_method1(ToMethod(r.Id()),MakeConstantList(v2.Id(),vnew.Id(),narg),s)
                            if ErrorIn(try_10) {try_7 = try_10
                            } else {
                            arg_9 = ANY(try_10)
                            try_7 = Language.C_If.Make(arg_8.Id(),arg_9,n.Arg).ToEID()
                            }
                            } 
                          } 
                        if ErrorIn(try_7) {loop_4 = try_7
                        } else {
                        va_arg2 = ANY(try_7)
                        va_arg1.Arg = va_arg2
                        loop_4 = va_arg2.ToEID()
                        }
                        } 
                      } 
                    } 
                  } else {
                  loop_4 = EID{CFALSE.Id(),0}
                  } 
                }
                if ErrorIn(loop_4) {Result = loop_4
                break
                } else {
                }
                } 
              } 
            if !ErrorIn(Result) {
            Result = Core.F_CALL(C_c_code,ARGS(EID{n.Id(),0},EID{s.Id(),0}))
            }
            }
            } 
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: c_code_multiple @ For (throw: true) 
func E_Optimize_c_code_multiple_For (self EID,_Zt EID,s EID) EID { 
    return F_Optimize_c_code_multiple_For(Language.To_For(OBJ(self)),ToType(OBJ(_Zt)),ToClass(OBJ(s)) )} 
  
// ------------------------ Collect/ Image / Select / Lselect ------------------------------
// an Iteration builds a set
/* The go function for: c_type(self:Iteration) [status=1] */
func F_c_type_Iteration (self *Language.Iteration) EID { 
    var Result EID
    { var _Zt *ClaireType
      var try_1 EID
      try_1 = Core.F_CALL(C_c_type,ARGS(self.SetArg.ToEID()))
      if ErrorIn(try_1) {Result = try_1
      } else {
      _Zt = ToType(OBJ(try_1))
      F_Optimize_range_infers_for_Variable(self.ClaireVar,F_Optimize_pmember_type(_Zt),_Zt)
      if (Core.F_get_property(C_of,ToObject(self.Id())) != CNULL) { 
        { var arg_2 *ClaireClass
          if ((self.Isa.IsIn(Language.C_Select) == CTRUE) || 
              (self.Isa.IsIn(Language.C_Image) == CTRUE)) { 
            arg_2 = C_set
            } else {
            arg_2 = C_list
            } 
          Result = EID{Core.F_param_I_class(arg_2,ToType(OBJ(Core.F_CALL(C_of,ARGS(EID{self.Id(),0}))))).Id(),0}
          } 
        } else {
        { var arg_3 *ClaireClass
          if ((self.Isa.IsIn(Language.C_Select) == CTRUE) || 
              (self.Isa.IsIn(Language.C_Image) == CTRUE)) { 
            arg_3 = C_set
            } else {
            arg_3 = C_list
            } 
          { var arg_4 *ClaireType
            var try_5 EID
            if ((self.Isa.IsIn(Language.C_Select) == CTRUE) || 
                (self.Isa.IsIn(Language.C_Lselect) == CTRUE)) { 
              { var arg_6 *ClaireType
                var try_7 EID
                try_7 = Core.F_CALL(C_c_type,ARGS(self.SetArg.ToEID()))
                if ErrorIn(try_7) {try_5 = try_7
                } else {
                arg_6 = ToType(OBJ(try_7))
                try_5 = EID{F_Optimize_pmember_type(arg_6).Id(),0}
                }
                } 
              } else {
              { var arg_8 *ClaireType
                var try_9 EID
                try_9 = Core.F_CALL(C_c_type,ARGS(self.Arg.ToEID()))
                if ErrorIn(try_9) {try_5 = try_9
                } else {
                arg_8 = ToType(OBJ(try_9))
                try_5 = EID{F_Optimize_ptype_type(arg_8).Id(),0}
                }
                } 
              } 
            if ErrorIn(try_5) {Result = try_5
            } else {
            arg_4 = ToType(OBJ(try_5))
            Result = EID{Core.F_nth_class1(arg_3,arg_4).Id(),0}
            }
            } 
          } 
        } 
      }
      } 
    return Result} 
  
// The EID go function for: c_type @ Iteration (throw: true) 
func E_c_type_Iteration (self EID) EID { 
    return F_c_type_Iteration(Language.To_Iteration(OBJ(self)) )} 
  
// They are all expended into a For except for Collect(bag : list or set)
/* The go function for: c_code(self:Iteration) [status=1] */
func F_c_code_Iteration (self *Language.Iteration) EID { 
    var Result EID
    { var sx *ClaireAny = self.SetArg
      { var _Zt *ClaireType
        var try_1 EID
        try_1 = Core.F_CALL(C_c_type,ARGS(sx.ToEID()))
        if ErrorIn(try_1) {Result = try_1
        } else {
        _Zt = ToType(OBJ(try_1))
        if (self.Isa.IsIn(Language.C_For) == CTRUE) { 
          Result = Core.F_CALL(C_c_code,ARGS(EID{self.Id(),0},EID{C_any.Id(),0}))
          } else {
          var g0281I *ClaireBoolean
          if (self.Isa.IsIn(Language.C_Collect) == CTRUE) { 
            g0281I = MakeBoolean((_Zt.Included(ToType(C_list.Id())) == CTRUE) || (_Zt.Included(ToType(C_set.Id())) == CTRUE))
            } else {
            g0281I = CFALSE
            } 
          if (g0281I == CTRUE) { 
            F_Optimize_range_infers_for_Variable(self.ClaireVar,F_Optimize_pmember_type(_Zt),_Zt)
            { var ty *ClaireType
              var try_2 EID
              { var arg_3 *ClaireType
                var try_4 EID
                try_4 = Core.F_CALL(C_c_type,ARGS(self.Arg.ToEID()))
                if ErrorIn(try_4) {try_2 = try_4
                } else {
                arg_3 = ToType(OBJ(try_4))
                try_2 = EID{F_Optimize_ptype_type(arg_3).Id(),0}
                }
                } 
              if ErrorIn(try_2) {Result = try_2
              } else {
              ty = ToType(OBJ(try_2))
              { var x *Language.Collect
                var try_5 EID
                { var _CL_obj *Language.Collect = Language.To_Collect(new(Language.Collect).Is(Language.C_Collect))
                  _CL_obj.ClaireVar = self.ClaireVar
                  { 
                    var va_arg1 *Language.Iteration
                    var va_arg2 *ClaireAny
                    va_arg1 = Language.To_Iteration(_CL_obj.Id())
                    var try_6 EID
                    try_6 = F_Compile_c_strict_code_any(sx,_Zt.Class_I())
                    if ErrorIn(try_6) {try_5 = try_6
                    } else {
                    va_arg2 = ANY(try_6)
                    va_arg1.SetArg = va_arg2
                    try_5 = va_arg2.ToEID()
                    }
                    } 
                  if !ErrorIn(try_5) {
                  { 
                    var va_arg1 *Language.Iteration
                    var va_arg2 *ClaireAny
                    va_arg1 = Language.To_Iteration(_CL_obj.Id())
                    var try_7 EID
                    try_7 = Core.F_CALL(C_c_code,ARGS(self.Arg.ToEID(),EID{C_any.Id(),0}))
                    if ErrorIn(try_7) {try_5 = try_7
                    } else {
                    va_arg2 = ANY(try_7)
                    va_arg1.Arg = va_arg2
                    try_5 = va_arg2.ToEID()
                    }
                    } 
                  if !ErrorIn(try_5) {
                  try_5 = EID{_CL_obj.Id(),0}
                  }}
                  } 
                if ErrorIn(try_5) {Result = try_5
                } else {
                x = Language.To_Collect(OBJ(try_5))
                if (ty.Id() == C_void.Id()) { 
                  Result = Core.F_CALL(C_Compile_Cerror,ARGS(EID{MakeString("[205] use of void expression ~S in ~S").Id(),0},self.Arg.ToEID(),EID{self.Id(),0}))
                  } else {
                  Result = EID{CFALSE.Id(),0}
                  } 
                if !ErrorIn(Result) {
                if (Core.F_get_property(C_of,ToObject(self.Id())) != CNULL) { 
                  if ((C_compiler.Safety >= 2) || 
                      (ty.Included(ToType(OBJ(Core.F_CALL(C_of,ARGS(EID{self.Id(),0}))))) == CTRUE)) { 
                    x.Of = ToType(OBJ(Core.F_CALL(C_of,ARGS(EID{self.Id(),0}))))
                    Result = EID{x.Id(),0}
                    } else {
                    F_Compile_warn_void()
                    { var arg_8 *ClaireList
                      var try_9 EID
                      { 
                        var v_bag_arg *ClaireAny
                        try_9= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                        ToList(OBJ(try_9)).AddFast(self.Id())
                        var try_10 EID
                        try_10 = Core.F_CALL(C_c_type,ARGS(self.Arg.ToEID()))
                        if ErrorIn(try_10) {try_9 = try_10
                        } else {
                        v_bag_arg = ANY(try_10)
                        ToList(OBJ(try_9)).AddFast(v_bag_arg)
                        ToList(OBJ(try_9)).AddFast(ANY(Core.F_CALL(C_of,ARGS(EID{self.Id(),0}))))}
                        } 
                      if ErrorIn(try_9) {Result = try_9
                      } else {
                      arg_8 = ToList(OBJ(try_9))
                      Result = Core.F_tformat_string(MakeString("unsafe typed collect (~S): ~S not in ~S [261]\n"),1,arg_8)
                      }
                      } 
                    if !ErrorIn(Result) {
                    { var arg_11 *Language.Call
                      { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                        _CL_obj.Selector = Core.C_check_in
                        _CL_obj.Args = MakeConstantList(x.Id(),C_list.Id(),ANY(Core.F_CALL(C_of,ARGS(EID{self.Id(),0}))))
                        arg_11 = _CL_obj
                        } 
                      Result = Core.F_CALL(C_c_code,ARGS(EID{arg_11.Id(),0},EID{C_list.Id(),0}))
                      } 
                    }
                    } 
                  } else {
                  Result = EID{x.Id(),0}
                  } 
                }
                }
                } 
              }
              } 
            } else {
            { var val *ClaireAny
              if (self.Isa.IsIn(Language.C_Image) == CTRUE) { 
                val = ToType(CEMPTY.Id()).EmptySet().Id()
                } else {
                val = ToType(CEMPTY.Id()).EmptyList().Id()
                } 
              { var v *ClaireVariable
                { var arg_12 int
                  C_OPT.MaxVars = (C_OPT.MaxVars+1)
                  arg_12 = 0
                  { var arg_13 *ClaireClass
                    if (self.Isa.IsIn(Language.C_Image) == CTRUE) { 
                      arg_13 = C_set
                      } else {
                      arg_13 = C_list
                      } 
                    v = F_Compile_Variable_I_symbol(F_append_symbol(self.ClaireVar.Pname,MakeString("_bag").Id()),arg_12,arg_13.Id())
                    } 
                  } 
                
                if (Core.F_get_property(C_of,ToObject(self.Id())) != CNULL) { 
                  { var _ZtypeIn *ClaireType
                    var try_14 EID
                    try_14 = Core.F_CALL(C_c_type,ARGS(self.Arg.ToEID()))
                    if ErrorIn(try_14) {Result = try_14
                    } else {
                    _ZtypeIn = ToType(OBJ(try_14))
                    if ((F_Optimize_ptype_type(_ZtypeIn).Included(ToType(OBJ(Core.F_CALL(C_of,ARGS(EID{self.Id(),0}))))) != CTRUE) && 
                        (C_compiler.Safety <= 2)) { 
                      F_Compile_warn_void()
                      Core.F_tformat_string(MakeString("unsafe bag construction (~S) : a ~S is not a ~S [262]\n"),1,MakeConstantList(self.ClaireVar.Id(),_ZtypeIn.Id(),ANY(Core.F_CALL(C_of,ARGS(EID{self.Id(),0})))))
                      } 
                    ToBag(val).Cast_I(ToType(OBJ(Core.F_CALL(C_of,ARGS(EID{self.Id(),0})))))
                    { 
                      var va_arg1 *ClaireVariable
                      var va_arg2 *ClaireType
                      va_arg1 = v
                      va_arg2 = Core.F_param_I_class(ToClass(v.Range.Id()),ToType(OBJ(Core.F_CALL(C_of,ARGS(EID{self.Id(),0})))))
                      va_arg1.Range = va_arg2
                      Result = EID{va_arg2.Id(),0}
                      } 
                    }
                    } 
                  } else {
                  if (C_set.Id() == val.Isa.Id()) { 
                    { var _CL_obj *Language.Set = Language.To_Set(new(Language.Set).Is(Language.C_Set))
                      _CL_obj.Of = ToType(CEMPTY.Id())
                      val = _CL_obj.Id()
                      } 
                    } else {
                    { var _CL_obj *Language.List = Language.To_List(new(Language.List).Is(Language.C_List))
                      _CL_obj.Of = ToType(CEMPTY.Id())
                      val = _CL_obj.Id()
                      } 
                    } 
                  Result = val.ToEID()
                  } 
                if !ErrorIn(Result) {
                { var _CL_obj *Language.Let = Language.To_Let(new(Language.Let).Is(Language.C_Let))
                  _CL_obj.ClaireVar = v
                  _CL_obj.Value = val
                  { 
                    var va_arg1 *Language.Let
                    var va_arg2 *ClaireAny
                    va_arg1 = _CL_obj
                    var try_15 EID
                    { var _CL_obj *Language.Do = Language.To_Do(new(Language.Do).Is(Language.C_Do))
                      { 
                        var va_arg1 *Language.Do
                        var va_arg2 *ClaireList
                        va_arg1 = _CL_obj
                        var try_16 EID
                        { 
                          var v_bag_arg *ClaireAny
                          try_16= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                          var try_17 EID
                          { var arg_18 *Language.For
                            { var _CL_obj *Language.For = Language.To_For(new(Language.For).Is(Language.C_For))
                              _CL_obj.ClaireVar = self.ClaireVar
                              _CL_obj.SetArg = sx
                              { 
                                var va_arg1 *Language.Iteration
                                var va_arg2 *ClaireAny
                                va_arg1 = Language.To_Iteration(_CL_obj.Id())
                                { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                  _CL_obj.Selector = ToProperty(C_add_I.Id())
                                  _CL_obj.Args = MakeConstantList(v.Id(),self.Arg)
                                  va_arg2 = _CL_obj.Id()
                                  } 
                                va_arg1.Arg = va_arg2
                                } 
                              arg_18 = _CL_obj
                              } 
                            try_17 = Core.F_CALL(C_c_code,ARGS(EID{arg_18.Id(),0},EID{C_any.Id(),0}))
                            } 
                          if ErrorIn(try_17) {try_16 = try_17
                          } else {
                          v_bag_arg = ANY(try_17)
                          ToList(OBJ(try_16)).AddFast(v_bag_arg)
                          ToList(OBJ(try_16)).AddFast(v.Id())}
                          } 
                        if ErrorIn(try_16) {try_15 = try_16
                        } else {
                        va_arg2 = ToList(OBJ(try_16))
                        va_arg1.Args = va_arg2
                        try_15 = EID{va_arg2.Id(),0}
                        }
                        } 
                      if !ErrorIn(try_15) {
                      try_15 = EID{_CL_obj.Id(),0}
                      }
                      } 
                    if ErrorIn(try_15) {Result = try_15
                    } else {
                    va_arg2 = ANY(try_15)
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
        } 
      } 
    return Result} 
  
// The EID go function for: c_code @ Iteration (throw: true) 
func E_c_code_Iteration (self EID) EID { 
    return F_c_code_Iteration(Language.To_Iteration(OBJ(self)) )} 
  
// new in v3.1.16
// selection has its own optimization method that takes care of the polymorphism
/* The go function for: c_code(self:Select) [status=1] */
func F_c_code_Select (self *Language.Select) EID { 
    var Result EID
    Result = F_Optimize_c_code_select_Iteration(Language.To_Iteration(self.Id()),C_set)
    return Result} 
  
// The EID go function for: c_code @ Select (throw: true) 
func E_c_code_Select (self EID) EID { 
    return F_c_code_Select(Language.To_Select(OBJ(self)) )} 
  
/* The go function for: c_code(self:Lselect) [status=1] */
func F_c_code_Lselect (self *Language.Lselect) EID { 
    var Result EID
    Result = F_Optimize_c_code_select_Iteration(Language.To_Iteration(self.Id()),C_list)
    return Result} 
  
// The EID go function for: c_code @ Lselect (throw: true) 
func E_c_code_Lselect (self EID) EID { 
    return F_c_code_Lselect(Language.To_Lselect(OBJ(self)) )} 
  
// changed in CLAIRE 4 (cf trans -> init.cl)
// x is set or list, tells what we want as output
/* The go function for: c_code_select(self:Iteration,x:class) [status=1] */
func F_Optimize_c_code_select_Iteration (self *Language.Iteration,x *ClaireClass) EID { 
    var Result EID
    { var sx *ClaireAny = self.SetArg
      { var _Zt *ClaireType
        var try_1 EID
        try_1 = Core.F_CALL(C_c_type,ARGS(sx.ToEID()))
        if ErrorIn(try_1) {Result = try_1
        } else {
        _Zt = ToType(OBJ(try_1))
        { var st *ClaireAny
          var try_2 EID
          try_2 = F_Optimize_enumerate_code_any(sx,_Zt)
          if ErrorIn(try_2) {Result = try_2
          } else {
          st = ANY(try_2)
          { var val *ClaireBag
            if (x.Id() == C_set.Id()) { 
              val = ToBag(ToType(CEMPTY.Id()).EmptySet().Id())
              } else {
              val = ToBag(ToType(CEMPTY.Id()).EmptyList().Id())
              } 
            { var v1 *ClaireVariable
              { var arg_3 int
                C_OPT.MaxVars = (C_OPT.MaxVars+1)
                arg_3 = 0
                v1 = F_Compile_Variable_I_symbol(F_append_symbol(self.ClaireVar.Pname,MakeString("_in").Id()),arg_3,_Zt.Id())
                } 
              { var v2 *ClaireVariable
                { var arg_4 int
                  C_OPT.MaxVars = (C_OPT.MaxVars+1)
                  arg_4 = 0
                  v2 = F_Compile_Variable_I_symbol(F_append_symbol(self.ClaireVar.Pname,MakeString("_out").Id()),arg_4,x.Id())
                  } 
                if (Core.F_get_property(C_of,ToObject(self.Id())) != CNULL) { 
                  { var _ZtypeIn *ClaireType = F_Optimize_pmember_type(_Zt)
                    if ((F_Optimize_ptype_type(_ZtypeIn).Included(ToType(OBJ(Core.F_CALL(C_of,ARGS(EID{self.Id(),0}))))) != CTRUE) && 
                        (C_compiler.Safety <= 1)) { 
                      F_Compile_warn_void()
                      Core.F_tformat_string(MakeString("unsafe bag construction (~S) : a ~S is not a ~S [262]\n"),1,MakeConstantList(self.ClaireVar.Id(),_ZtypeIn.Id(),ANY(Core.F_CALL(C_of,ARGS(EID{self.Id(),0})))))
                      } 
                    val.Cast_I(ToType(OBJ(Core.F_CALL(C_of,ARGS(EID{self.Id(),0})))))
                    v2.Range = Core.F_param_I_class(x,ToType(OBJ(Core.F_CALL(C_of,ARGS(EID{self.Id(),0})))))
                    Result = F_Optimize_inner_select_Iteration(self,v2.Id(),sx,val.Id())
                    } 
                  }  else if (_Zt.Included(ToType(x.Id())) == CTRUE) { 
                  { var _CL_obj *Language.Let = Language.To_Let(new(Language.Let).Is(Language.C_Let))
                    _CL_obj.ClaireVar = v1
                    _CL_obj.Value = st
                    { 
                      var va_arg1 *Language.Let
                      var va_arg2 *ClaireAny
                      va_arg1 = _CL_obj
                      var try_5 EID
                      { var arg_6 *Compile_CCast
                        var try_7 EID
                        { var _CL_obj *Compile_CCast = To_Compile_CCast(new(Compile_CCast).Is(C_Compile_C_cast))
                          { 
                            var va_arg1 *Compile_CCast
                            var va_arg2 *ClaireAny
                            va_arg1 = _CL_obj
                            var try_8 EID
                            { var arg_9 *Language.Call
                              { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                _CL_obj.Selector = C_empty
                                _CL_obj.Args = MakeConstantList(v1.Id())
                                arg_9 = _CL_obj
                                } 
                              try_8 = Core.F_CALL(C_c_code,ARGS(EID{arg_9.Id(),0},EID{x.Id(),0}))
                              } 
                            if ErrorIn(try_8) {try_7 = try_8
                            } else {
                            va_arg2 = ANY(try_8)
                            va_arg1.Arg = va_arg2
                            try_7 = va_arg2.ToEID()
                            }
                            } 
                          if !ErrorIn(try_7) {
                          _CL_obj.SetArg = x
                          try_7 = EID{_CL_obj.Id(),0}
                          }
                          } 
                        if ErrorIn(try_7) {try_5 = try_7
                        } else {
                        arg_6 = To_Compile_CCast(OBJ(try_7))
                        try_5 = F_Optimize_inner_select_Iteration(self,v2.Id(),v1.Id(),arg_6.Id())
                        }
                        } 
                      if ErrorIn(try_5) {Result = try_5
                      } else {
                      va_arg2 = ANY(try_5)
                      va_arg1.Arg = va_arg2
                      Result = va_arg2.ToEID()
                      }
                      } 
                    if !ErrorIn(Result) {
                    Result = EID{_CL_obj.Id(),0}
                    }
                    } 
                  } else {
                  { var arg_10 *Language.Construct
                    if (x.Id() == C_set.Id()) { 
                      { var _CL_obj *Language.Set = Language.To_Set(new(Language.Set).Is(Language.C_Set))
                        _CL_obj.Of = ToType(CEMPTY.Id())
                        arg_10 = Language.To_Construct(_CL_obj.Id())
                        } 
                      } else {
                      { var _CL_obj *Language.List = Language.To_List(new(Language.List).Is(Language.C_List))
                        _CL_obj.Of = ToType(CEMPTY.Id())
                        arg_10 = Language.To_Construct(_CL_obj.Id())
                        } 
                      } 
                    Result = F_Optimize_inner_select_Iteration(self,v2.Id(),sx,arg_10.Id())
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
  
// The EID go function for: c_code_select @ Iteration (throw: true) 
func E_Optimize_c_code_select_Iteration (self EID,x EID) EID { 
    return F_Optimize_c_code_select_Iteration(Language.To_Iteration(OBJ(self)),ToClass(OBJ(x)) )} 
  
// v3.2.01
// sub-procedure : creates the iteration over sx, adds the selected value (var) into l2 (set or list)
/* The go function for: inner_select(self:Iteration,v2:any,sx:any,val:any) [status=1] */
func F_Optimize_inner_select_Iteration (self *Language.Iteration,v2 *ClaireAny,sx *ClaireAny,val *ClaireAny) EID { 
    var Result EID
    { var _CL_obj *Language.Let = Language.To_Let(new(Language.Let).Is(Language.C_Let))
      _CL_obj.ClaireVar = To_Variable(v2)
      _CL_obj.Value = val
      { 
        var va_arg1 *Language.Let
        var va_arg2 *ClaireAny
        va_arg1 = _CL_obj
        var try_1 EID
        { var _CL_obj *Language.Do = Language.To_Do(new(Language.Do).Is(Language.C_Do))
          { 
            var va_arg1 *Language.Do
            var va_arg2 *ClaireList
            va_arg1 = _CL_obj
            var try_2 EID
            { 
              var v_bag_arg *ClaireAny
              try_2= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
              var try_3 EID
              { var arg_4 *Language.For
                { var _CL_obj *Language.For = Language.To_For(new(Language.For).Is(Language.C_For))
                  _CL_obj.ClaireVar = self.ClaireVar
                  _CL_obj.SetArg = sx
                  { 
                    var va_arg1 *Language.Iteration
                    var va_arg2 *ClaireAny
                    va_arg1 = Language.To_Iteration(_CL_obj.Id())
                    { var _CL_obj *Language.If = Language.To_If(new(Language.If).Is(Language.C_If))
                      _CL_obj.Test = self.Arg
                      { 
                        var va_arg1 *Language.If
                        var va_arg2 *ClaireAny
                        va_arg1 = _CL_obj
                        { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                          _CL_obj.Selector = ToProperty(C_add_I.Id())
                          _CL_obj.Args = MakeConstantList(v2,self.ClaireVar.Id())
                          va_arg2 = _CL_obj.Id()
                          } 
                        va_arg1.Arg = va_arg2
                        } 
                      va_arg2 = _CL_obj.Id()
                      } 
                    va_arg1.Arg = va_arg2
                    } 
                  arg_4 = _CL_obj
                  } 
                try_3 = Core.F_CALL(C_c_code,ARGS(EID{arg_4.Id(),0},EID{C_any.Id(),0}))
                } 
              if ErrorIn(try_3) {try_2 = try_3
              } else {
              v_bag_arg = ANY(try_3)
              ToList(OBJ(try_2)).AddFast(v_bag_arg)
              ToList(OBJ(try_2)).AddFast(v2)}
              } 
            if ErrorIn(try_2) {try_1 = try_2
            } else {
            va_arg2 = ToList(OBJ(try_2))
            va_arg1.Args = va_arg2
            try_1 = EID{va_arg2.Id(),0}
            }
            } 
          if !ErrorIn(try_1) {
          try_1 = EID{_CL_obj.Id(),0}
          }
          } 
        if ErrorIn(try_1) {Result = try_1
        } else {
        va_arg2 = ANY(try_1)
        va_arg1.Arg = va_arg2
        Result = va_arg2.ToEID()
        }
        } 
      if !ErrorIn(Result) {
      Result = EID{_CL_obj.Id(),0}
      }
      } 
    return Result} 
  
// The EID go function for: inner_select @ Iteration (throw: true) 
func E_Optimize_inner_select_Iteration (self EID,v2 EID,sx EID,val EID) EID { 
    return F_Optimize_inner_select_Iteration(Language.To_Iteration(OBJ(self)),
      ANY(v2),
      ANY(sx),
      ANY(val) )} 
  
// if (other = unknown : some) the result is either the variable or unknown
/* The go function for: c_type(self:Exists) [status=1] */
func F_c_type_Exists (self *Language.Exists) EID { 
    var Result EID
    { var _Zt *ClaireType
      var try_1 EID
      try_1 = Core.F_CALL(C_c_type,ARGS(self.SetArg.ToEID()))
      if ErrorIn(try_1) {Result = try_1
      } else {
      _Zt = ToType(OBJ(try_1))
      F_Optimize_range_infers_for_Variable(self.ClaireVar,F_Optimize_pmember_type(_Zt),_Zt)
      if (self.Other == CNULL) { 
        Result = EID{F_Optimize_extends_type(F_Optimize_pmember_type(_Zt)).Id(),0}
        } else {
        Result = EID{C_boolean.Id(),0}
        } 
      }
      } 
    return Result} 
  
// The EID go function for: c_type @ Exists (throw: true) 
func E_c_type_Exists (self EID) EID { 
    return F_c_type_Exists(Language.To_Exists(OBJ(self)) )} 
  
// boolean, or any U boolean ?
/* The go function for: c_code(self:Exists,s:class) [status=1] */
func F_c_code_Exists (self *Language.Exists,s *ClaireClass) EID { 
    var Result EID
    { var _Zt *ClaireType
      var try_1 EID
      try_1 = Core.F_CALL(C_c_type,ARGS(self.SetArg.ToEID()))
      if ErrorIn(try_1) {Result = try_1
      } else {
      _Zt = ToType(OBJ(try_1))
      F_Optimize_range_infers_for_Variable(self.ClaireVar,F_Optimize_pmember_type(_Zt),_Zt)
      if (self.Other == CTRUE.Id()) { 
        { var arg_2 *Language.Call
          { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
            _CL_obj.Selector = Core.C_not
            { 
              var va_arg1 *Language.Call
              var va_arg2 *ClaireList
              va_arg1 = _CL_obj
              { 
                var v_bag_arg *ClaireAny
                va_arg2= ToType(CEMPTY.Id()).EmptyList()
                { var _CL_obj *Language.For = Language.To_For(new(Language.For).Is(Language.C_For))
                  _CL_obj.ClaireVar = self.ClaireVar
                  _CL_obj.SetArg = self.SetArg
                  { 
                    var va_arg1 *Language.Iteration
                    var va_arg2 *ClaireAny
                    va_arg1 = Language.To_Iteration(_CL_obj.Id())
                    { var _CL_obj *Language.If = Language.To_If(new(Language.If).Is(Language.C_If))
                      { 
                        var va_arg1 *Language.If
                        var va_arg2 *ClaireAny
                        va_arg1 = _CL_obj
                        { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                          _CL_obj.Selector = Core.C_not
                          _CL_obj.Args = MakeConstantList(self.Arg)
                          va_arg2 = _CL_obj.Id()
                          } 
                        va_arg1.Test = va_arg2
                        } 
                      _CL_obj.Arg = Language.C_Return.Make(CTRUE.Id())
                      va_arg2 = _CL_obj.Id()
                      } 
                    va_arg1.Arg = va_arg2
                    } 
                  v_bag_arg = _CL_obj.Id()
                  } 
                va_arg2.AddFast(v_bag_arg)} 
              va_arg1.Args = va_arg2
              } 
            arg_2 = _CL_obj
            } 
          Result = Core.F_CALL(C_c_code,ARGS(EID{arg_2.Id(),0},EID{s.Id(),0}))
          } 
        }  else if (self.Other == CNULL) { 
        { var v *ClaireVariable
          { var arg_3 int
            C_OPT.MaxVars = (C_OPT.MaxVars+1)
            arg_3 = 0
            v = F_Compile_Variable_I_symbol(F_append_symbol(self.ClaireVar.Pname,MakeString("_some").Id()),arg_3,F_Optimize_extends_type(self.ClaireVar.Range).Id())
            } 
          { var arg_4 *Language.Let
            { var _CL_obj *Language.Let = Language.To_Let(new(Language.Let).Is(Language.C_Let))
              _CL_obj.ClaireVar = v
              _CL_obj.Value = CNULL
              { 
                var va_arg1 *Language.Let
                var va_arg2 *ClaireAny
                va_arg1 = _CL_obj
                { var _CL_obj *Language.Do = Language.To_Do(new(Language.Do).Is(Language.C_Do))
                  { 
                    var va_arg1 *Language.Do
                    var va_arg2 *ClaireList
                    va_arg1 = _CL_obj
                    { 
                      var v_bag_arg *ClaireAny
                      va_arg2= ToType(CEMPTY.Id()).EmptyList()
                      { var _CL_obj *Language.For = Language.To_For(new(Language.For).Is(Language.C_For))
                        _CL_obj.ClaireVar = self.ClaireVar
                        _CL_obj.SetArg = self.SetArg
                        { 
                          var va_arg1 *Language.Iteration
                          var va_arg2 *ClaireAny
                          va_arg1 = Language.To_Iteration(_CL_obj.Id())
                          { var _CL_obj *Language.If = Language.To_If(new(Language.If).Is(Language.C_If))
                            _CL_obj.Test = self.Arg
                            _CL_obj.Arg = Language.C_Return.Make(Language.C_Assign.Make(v.Id(),self.ClaireVar.Id()))
                            va_arg2 = _CL_obj.Id()
                            } 
                          va_arg1.Arg = va_arg2
                          } 
                        v_bag_arg = _CL_obj.Id()
                        } 
                      va_arg2.AddFast(v_bag_arg)
                      va_arg2.AddFast(v.Id())} 
                    va_arg1.Args = va_arg2
                    } 
                  va_arg2 = _CL_obj.Id()
                  } 
                va_arg1.Arg = va_arg2
                } 
              arg_4 = _CL_obj
              } 
            Result = Core.F_CALL(C_c_code,ARGS(EID{arg_4.Id(),0},EID{s.Id(),0}))
            } 
          } 
        } else {
        { var arg_5 *Language.Call
          { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
            _CL_obj.Selector = C_boolean_I
            { 
              var va_arg1 *Language.Call
              var va_arg2 *ClaireList
              va_arg1 = _CL_obj
              { 
                var v_bag_arg *ClaireAny
                va_arg2= ToType(CEMPTY.Id()).EmptyList()
                { var _CL_obj *Language.For = Language.To_For(new(Language.For).Is(Language.C_For))
                  _CL_obj.ClaireVar = self.ClaireVar
                  _CL_obj.SetArg = self.SetArg
                  { 
                    var va_arg1 *Language.Iteration
                    var va_arg2 *ClaireAny
                    va_arg1 = Language.To_Iteration(_CL_obj.Id())
                    { var _CL_obj *Language.If = Language.To_If(new(Language.If).Is(Language.C_If))
                      _CL_obj.Test = self.Arg
                      _CL_obj.Arg = Language.C_Return.Make(CTRUE.Id())
                      va_arg2 = _CL_obj.Id()
                      } 
                    va_arg1.Arg = va_arg2
                    } 
                  v_bag_arg = _CL_obj.Id()
                  } 
                va_arg2.AddFast(v_bag_arg)} 
              va_arg1.Args = va_arg2
              } 
            arg_5 = _CL_obj
            } 
          Result = Core.F_CALL(C_c_code,ARGS(EID{arg_5.Id(),0},EID{s.Id(),0}))
          } 
        } 
      }
      } 
    return Result} 
  
// The EID go function for: c_code @ Exists (throw: true) 
func E_c_code_Exists (self EID,s EID) EID { 
    return F_c_code_Exists(Language.To_Exists(OBJ(self)),ToClass(OBJ(s)) )} 
  
// exists
// am Image builds a set
/* The go function for: c_type(self:Image) [status=1] */
func F_c_type_Image (self *Language.Image) EID { 
    var Result EID
    { var _Zt *ClaireType
      var try_1 EID
      { var arg_2 *ClaireType
        var try_3 EID
        try_3 = Core.F_CALL(C_c_type,ARGS(self.SetArg.ToEID()))
        if ErrorIn(try_3) {try_1 = try_3
        } else {
        arg_2 = ToType(OBJ(try_3))
        try_1 = EID{F_Optimize_ptype_type(arg_2).Id(),0}
        }
        } 
      if ErrorIn(try_1) {Result = try_1
      } else {
      _Zt = ToType(OBJ(try_1))
      F_Optimize_range_infers_for_Variable(self.ClaireVar,F_Optimize_pmember_type(_Zt),_Zt)
      if (self.Of.Id() != CNULL) { 
        Result = EID{Core.F_param_I_class(C_set,self.Of).Id(),0}
        } else {
        { var arg_4 *ClaireType
          var try_5 EID
          try_5 = Core.F_CALL(C_c_type,ARGS(self.Arg.ToEID()))
          if ErrorIn(try_5) {Result = try_5
          } else {
          arg_4 = ToType(OBJ(try_5))
          Result = EID{Core.F_nth_class1(C_set,arg_4).Id(),0}
          }
          } 
        } 
      }
      } 
    return Result} 
  
// The EID go function for: c_type @ Image (throw: true) 
func E_c_type_Image (self EID) EID { 
    return F_c_type_Image(Language.To_Image(OBJ(self)) )} 
  
/* The go function for: c_type(self:Select) [status=1] */
func F_c_type_Select (self *Language.Select) EID { 
    var Result EID
    { var _Zt *ClaireType
      var try_1 EID
      try_1 = Core.F_CALL(C_c_type,ARGS(self.SetArg.ToEID()))
      if ErrorIn(try_1) {Result = try_1
      } else {
      _Zt = ToType(OBJ(try_1))
      F_Optimize_range_infers_for_Variable(self.ClaireVar,F_Optimize_pmember_type(_Zt),_Zt)
      if (self.Of.Id() != CNULL) { 
        Result = EID{Core.F_param_I_class(C_set,self.Of).Id(),0}
        } else {
        { var arg_2 *ClaireType
          var try_3 EID
          { var arg_4 *ClaireType
            var try_5 EID
            try_5 = Core.F_CALL(C_c_type,ARGS(self.SetArg.ToEID()))
            if ErrorIn(try_5) {try_3 = try_5
            } else {
            arg_4 = ToType(OBJ(try_5))
            try_3 = EID{F_Optimize_pmember_type(arg_4).Id(),0}
            }
            } 
          if ErrorIn(try_3) {Result = try_3
          } else {
          arg_2 = ToType(OBJ(try_3))
          Result = EID{Core.F_nth_class1(C_set,arg_2).Id(),0}
          }
          } 
        } 
      }
      } 
    return Result} 
  
// The EID go function for: c_type @ Select (throw: true) 
func E_c_type_Select (self EID) EID { 
    return F_c_type_Select(Language.To_Select(OBJ(self)) )} 
  
// new in v3.1.06 : proper type inference !
/* The go function for: c_type(self:Lselect) [status=1] */
func F_c_type_Lselect (self *Language.Lselect) EID { 
    var Result EID
    { var _Zt *ClaireType
      var try_1 EID
      try_1 = Core.F_CALL(C_c_type,ARGS(self.SetArg.ToEID()))
      if ErrorIn(try_1) {Result = try_1
      } else {
      _Zt = ToType(OBJ(try_1))
      F_Optimize_range_infers_for_Variable(self.ClaireVar,F_Optimize_pmember_type(_Zt),_Zt)
      if (self.Of.Id() != CNULL) { 
        Result = EID{Core.F_param_I_class(C_list,self.Of).Id(),0}
        } else {
        { var arg_2 *ClaireType
          var try_3 EID
          { var arg_4 *ClaireType
            var try_5 EID
            try_5 = Core.F_CALL(C_c_type,ARGS(self.SetArg.ToEID()))
            if ErrorIn(try_5) {try_3 = try_5
            } else {
            arg_4 = ToType(OBJ(try_5))
            try_3 = EID{F_Optimize_pmember_type(arg_4).Id(),0}
            }
            } 
          if ErrorIn(try_3) {Result = try_3
          } else {
          arg_2 = ToType(OBJ(try_3))
          Result = EID{Core.F_nth_class1(C_list,arg_2).Id(),0}
          }
          } 
        } 
      }
      } 
    return Result} 
  
// The EID go function for: c_type @ Lselect (throw: true) 
func E_c_type_Lselect (self EID) EID { 
    return F_c_type_Lselect(Language.To_Lselect(OBJ(self)) )} 
  
//______________________  while/until  __________________________________
// similar to a For
/* The go function for: c_type(self:While) [status=1] */
func F_c_type_While (self *Language.While) EID { 
    var Result EID
    { var arg_1 *ClaireType
      var try_2 EID
      try_2 = F_Compile_return_type_any(self.Arg)
      if ErrorIn(try_2) {Result = try_2
      } else {
      arg_1 = ToType(OBJ(try_2))
      Result = EID{F_Optimize_infers_from_type(arg_1,self.Id()).Id(),0}
      }
      } 
    return Result} 
  
// The EID go function for: c_type @ While (throw: true) 
func E_c_type_While (self EID) EID { 
    return F_c_type_While(Language.To_While(OBJ(self)) )} 
  
/* The go function for: c_code(self:While,s:class) [status=1] */
func F_c_code_While (self *Language.While,s *ClaireClass) EID { 
    var Result EID
    { var _CL_obj *Language.While = Language.To_While(new(Language.While).Is(Language.C_While))
      { 
        var va_arg1 *Language.While
        var va_arg2 *ClaireAny
        va_arg1 = _CL_obj
        var try_1 EID
        try_1 = F_Optimize_c_boolean_any(self.Test)
        if ErrorIn(try_1) {Result = try_1
        } else {
        va_arg2 = ANY(try_1)
        va_arg1.Test = va_arg2
        Result = va_arg2.ToEID()
        }
        } 
      if !ErrorIn(Result) {
      { 
        var va_arg1 *Language.While
        var va_arg2 *ClaireAny
        va_arg1 = _CL_obj
        var try_2 EID
        try_2 = Core.F_CALL(C_c_code,ARGS(self.Arg.ToEID(),EID{C_void.Id(),0}))
        if ErrorIn(try_2) {Result = try_2
        } else {
        va_arg2 = ANY(try_2)
        va_arg1.Arg = va_arg2
        Result = va_arg2.ToEID()
        }
        } 
      if !ErrorIn(Result) {
      _CL_obj.Other = self.Other
      Result = EID{_CL_obj.Id(),0}
      }}
      } 
    return Result} 
  
// The EID go function for: c_code @ While (throw: true) 
func E_c_code_While (self EID,s EID) EID { 
    return F_c_code_While(Language.To_While(OBJ(self)),ToClass(OBJ(s)) )} 
  
//      if (s != void & s != any)
//        (//[5] ... insert a to_C with s = ~S for ~S // s,self,
//         to_C(arg = r, set_arg = s)) // v3.3
//      else r)))
// *********************************************************************
// *     Part 6: Iterate                                               *
// *********************************************************************
// finds the right restriction of Iterate
// Iterate applies to the non-evaluated types (meta level)
/* The go function for: Iterate!(self:Iteration) [status=0] */
func F_Optimize_Iterate_I_Iteration (self *Language.Iteration) *ClaireAny { 
    return  F_Optimize_restriction_I_property(Language.C_Iterate,MakeConstantList(MakeConstantSet(self.SetArg).Id(),MakeConstantSet(self.ClaireVar.Id()).Id(),C_any.Id()),CTRUE)
    } 
  
// The EID go function for: Iterate! @ Iteration (throw: false) 
func E_Optimize_Iterate_I_Iteration (self EID) EID { 
    return F_Optimize_Iterate_I_Iteration(Language.To_Iteration(OBJ(self)) ).ToEID()} 
  
// iteration methods
// note the beauty of this: we only apply the code transformation if
// we actually get a constant Interval
/* The go function for: iterate(x:Interval,v:Variable[range:(subtype[integer])],e:any) [status=1] */
func F_iterate_Interval (x *ClaireInterval,v *ClaireVariable,e *ClaireAny) EID { 
    var Result EID
    { var v *ClaireAny
      var try_1 EID
      try_1 = F_eval_any2(MakeInteger(x.Arg1).Id(),C_Interval)
      if ErrorIn(try_1) {Result = try_1
      } else {
      v = ANY(try_1)
      { var _Zmax int
        var try_2 EID
        try_2 = F_eval_any2(MakeInteger(x.Arg2).Id(),C_Interval)
        if ErrorIn(try_2) {Result = try_2
        } else {
        _Zmax = INT(try_2)
        Result= EID{CFALSE.Id(),0}
        for (ToInteger(v).Value <= _Zmax) { 
          
          v = ANY(Core.F_CALL(ToProperty(Core.C__plus.Id()),ARGS(v.ToEID(),EID{C__INT,IVAL(1)})))
          } 
        }
        } 
      }
      } 
    return Result} 
  
// The EID go function for: iterate @ Interval (throw: true) 
func E_iterate_Interval (x EID,v EID,e EID) EID { 
    return F_iterate_Interval(To_Interval(OBJ(x)),To_Variable(OBJ(v)),ANY(e) )} 
  
/* The go function for: iterate(x:array,v:Variable,e:any) [status=1] */
func F_iterate_array (x *ClaireList,v *ClaireVariable,e *ClaireAny) EID { 
    var Result EID
    { var _Zi int = 1
      { var _Za *ClaireList = x
        { var _Zmax int = _Za.Length()
          Result= EID{CFALSE.Id(),0}
          for (_Zi <= _Zmax) { 
            var loop_1 EID
            _ = loop_1
            { var v *ClaireAny
              _ = v
              var try_2 EID
              try_2 = _Za.Nth(_Zi)
              if ErrorIn(try_2) {loop_1 = try_2
              } else {
              v = ANY(try_2)
              
              _Zi = (_Zi+1)
              loop_1 = EID{C__INT,IVAL(_Zi)}
              }
              } 
            if ErrorIn(loop_1) {Result = loop_1
            break
            } else {
            } 
          }
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: iterate @ array (throw: true) 
func E_iterate_array (x EID,v EID,e EID) EID { 
    return F_iterate_array(ToArray(OBJ(x)),To_Variable(OBJ(v)),ANY(e) )} 
  
/* The go function for: Iterate(x:class,v:Variable,e:any) [status=0] */
func F_Iterate_class (x *ClaireClass,v *ClaireVariable,e *ClaireAny) *ClaireAny { 
    var Result *ClaireAny
    { 
      var _Zv_1 *ClaireClass
      _ = _Zv_1
      var _Zv_1_iter *ClaireAny
      Result= CFALSE.Id()
      var _Zv_1_support *ClaireSet
      _Zv_1_support = x.Descendants
      for i_it := 0; i_it < _Zv_1_support.Count; i_it++ { 
        _Zv_1_iter = _Zv_1_support.At(i_it)
        _Zv_1 = ToClass(_Zv_1_iter)
        { var _Zv_2 *ClaireBoolean
          { 
            var v *ClaireAny
            _ = v
            _Zv_2= CFALSE
            var v_support *ClaireList
            v_support = _Zv_1.Instances
            v_len := v_support.Length()
            for i_it := 0; i_it < v_len; i_it++ { 
              v = v_support.At(i_it)
              
              } 
            } 
          if (_Zv_2 == CTRUE) { 
            Result = _Zv_2.Id()
            break
            } 
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: Iterate @ class (throw: false) 
func E_Iterate_class (x EID,v EID,e EID) EID { 
    return F_Iterate_class(ToClass(OBJ(x)),To_Variable(OBJ(v)),ANY(e) ).ToEID()} 
  
/* The go function for: Iterate(x:any,v:Variable,e:any) [status=1] */
func F_Iterate_any1 (x *ClaireAny,v *ClaireVariable,e *ClaireAny) EID { 
    var Result EID
    { var v *ClaireAny
      var try_1 EID
      try_1 = EVAL(ToList(OBJ(Core.F_CALL(C_args,ARGS(x.ToEID())))).At(0))
      if ErrorIn(try_1) {Result = try_1
      } else {
      v = ANY(try_1)
      { var _Zmax *ClaireAny
        var try_2 EID
        try_2 = EVAL(ToList(OBJ(Core.F_CALL(C_args,ARGS(x.ToEID())))).At(1))
        if ErrorIn(try_2) {Result = try_2
        } else {
        _Zmax = ANY(try_2)
        Result= EID{CFALSE.Id(),0}
        for (ToBoolean(OBJ(Core.F_CALL(ToProperty(C__inf_equal.Id()),ARGS(v.ToEID(),_Zmax.ToEID())))) == CTRUE) { 
          
          v = ANY(Core.F_CALL(ToProperty(Core.C__plus.Id()),ARGS(v.ToEID(),EID{C__INT,IVAL(1)})))
          } 
        }
        } 
      }
      } 
    return Result} 
  
// The EID go function for: Iterate @ list<type_expression>(..[tuple(integer,integer)], Variable, any) (throw: true) 
func E_Iterate_any1 (x EID,v EID,e EID) EID { 
    return F_Iterate_any1(ANY(x),To_Variable(OBJ(v)),ANY(e) )} 
  
/* The go function for: Iterate(x:Lselect,v:Variable,e:any) [status=1] */
func F_Iterate_Lselect (x *Language.Lselect,v *ClaireVariable,e *ClaireAny) EID { 
    var Result EID
    { 
      var v *ClaireAny
      _ = v
      Result= EID{CFALSE.Id(),0}
      var v_support *ClaireList
      var try_1 EID
      { var arg_2 *ClaireAny
        var try_3 EID
        try_3 = EVAL(x.SetArg)
        if ErrorIn(try_3) {try_1 = try_3
        } else {
        arg_2 = ANY(try_3)
        try_1 = Core.F_enumerate_any(arg_2)
        }
        } 
      if ErrorIn(try_1) {Result = try_1
      } else {
      v_support = ToList(OBJ(try_1))
      v_len := v_support.Length()
      for i_it := 0; i_it < v_len; i_it++ { 
        v = v_support.At(i_it)
        var loop_4 EID
        _ = loop_4
        var g0285I *ClaireBoolean
        var try_5 EID
        { var arg_6 *ClaireAny
          var try_7 EID
          try_7 = EVAL(Language.F_substitution_any(x.Arg,x.ClaireVar,v))
          if ErrorIn(try_7) {try_5 = try_7
          } else {
          arg_6 = ANY(try_7)
          try_5 = EID{F_boolean_I_any(arg_6).Id(),0}
          }
          } 
        if ErrorIn(try_5) {loop_4 = try_5
        } else {
        g0285I = ToBoolean(OBJ(try_5))
        if (g0285I == CTRUE) { 
          loop_4 = e.ToEID()
          } else {
          loop_4 = EID{CFALSE.Id(),0}
          } 
        }
        if ErrorIn(loop_4) {Result = loop_4
        break
        } else {
        }}
        } 
      } 
    return Result} 
  
// The EID go function for: Iterate @ Lselect (throw: true) 
func E_Iterate_Lselect (x EID,v EID,e EID) EID { 
    return F_Iterate_Lselect(Language.To_Lselect(OBJ(x)),To_Variable(OBJ(v)),ANY(e) )} 
  
/* The go function for: Iterate(x:Select,v:Variable,e:any) [status=1] */
func F_Iterate_Select (x *Language.Select,v *ClaireVariable,e *ClaireAny) EID { 
    var Result EID
    { 
      var v *ClaireAny
      _ = v
      Result= EID{CFALSE.Id(),0}
      var v_support *ClaireList
      var try_1 EID
      { var arg_2 *ClaireAny
        var try_3 EID
        try_3 = EVAL(x.SetArg)
        if ErrorIn(try_3) {try_1 = try_3
        } else {
        arg_2 = ANY(try_3)
        try_1 = Core.F_enumerate_any(arg_2)
        }
        } 
      if ErrorIn(try_1) {Result = try_1
      } else {
      v_support = ToList(OBJ(try_1))
      v_len := v_support.Length()
      for i_it := 0; i_it < v_len; i_it++ { 
        v = v_support.At(i_it)
        var loop_4 EID
        _ = loop_4
        var g0286I *ClaireBoolean
        var try_5 EID
        { var arg_6 *ClaireAny
          var try_7 EID
          try_7 = EVAL(Language.F_substitution_any(x.Arg,x.ClaireVar,v))
          if ErrorIn(try_7) {try_5 = try_7
          } else {
          arg_6 = ANY(try_7)
          try_5 = EID{F_boolean_I_any(arg_6).Id(),0}
          }
          } 
        if ErrorIn(try_5) {loop_4 = try_5
        } else {
        g0286I = ToBoolean(OBJ(try_5))
        if (g0286I == CTRUE) { 
          loop_4 = e.ToEID()
          } else {
          loop_4 = EID{CFALSE.Id(),0}
          } 
        }
        if ErrorIn(loop_4) {Result = loop_4
        break
        } else {
        }}
        } 
      } 
    return Result} 
  
// The EID go function for: Iterate @ Select (throw: true) 
func E_Iterate_Select (x EID,v EID,e EID) EID { 
    return F_Iterate_Select(Language.To_Select(OBJ(x)),To_Variable(OBJ(v)),ANY(e) )} 
  
/* The go function for: Iterate(x:Collect,v:Variable,e:any) [status=1] */
func F_Iterate_Collect (x *Language.Collect,v *ClaireVariable,e *ClaireAny) EID { 
    var Result EID
    { 
      var C_Zv *ClaireAny
      _ = C_Zv
      Result= EID{CFALSE.Id(),0}
      var C_Zv_support *ClaireList
      var try_1 EID
      { var arg_2 *ClaireAny
        var try_3 EID
        try_3 = EVAL(x.SetArg)
        if ErrorIn(try_3) {try_1 = try_3
        } else {
        arg_2 = ANY(try_3)
        try_1 = Core.F_enumerate_any(arg_2)
        }
        } 
      if ErrorIn(try_1) {Result = try_1
      } else {
      C_Zv_support = ToList(OBJ(try_1))
      C_Zv_len := C_Zv_support.Length()
      for i_it := 0; i_it < C_Zv_len; i_it++ { 
        C_Zv = C_Zv_support.At(i_it)
        var loop_4 EID
        _ = loop_4
        { var v *ClaireAny
          _ = v
          var try_5 EID
          try_5 = EVAL(Language.F_substitution_any(x.Arg,x.ClaireVar,C_Zv))
          if ErrorIn(try_5) {loop_4 = try_5
          } else {
          v = ANY(try_5)
          loop_4 = e.ToEID()
          }
          } 
        if ErrorIn(loop_4) {Result = loop_4
        break
        } else {
        }}
        } 
      } 
    return Result} 
  
// The EID go function for: Iterate @ Collect (throw: true) 
func E_Iterate_Collect (x EID,v EID,e EID) EID { 
    return F_Iterate_Collect(Language.To_Collect(OBJ(x)),To_Variable(OBJ(v)),ANY(e) )} 
  
/* The go function for: Iterate(x:any,v:Variable,e:any) [status=1] */
func F_Iterate_any2 (x *ClaireAny,v *ClaireVariable,e *ClaireAny) EID { 
    var Result EID
    { 
      var v *ClaireAny
      _ = v
      Result= EID{CFALSE.Id(),0}
      var v_support *ClaireList
      var try_1 EID
      { var arg_2 *ClaireAny
        var try_3 EID
        try_3 = EVAL(ToList(OBJ(Core.F_CALL(C_args,ARGS(x.ToEID())))).At(0))
        if ErrorIn(try_3) {try_1 = try_3
        } else {
        arg_2 = ANY(try_3)
        try_1 = Core.F_enumerate_any(arg_2)
        }
        } 
      if ErrorIn(try_1) {Result = try_1
      } else {
      v_support = ToList(OBJ(try_1))
      v_len := v_support.Length()
      for i_it := 0; i_it < v_len; i_it++ { 
        v = v_support.At(i_it)
        var loop_4 EID
        _ = loop_4
        var g0287I *ClaireBoolean
        var try_5 EID
        { var arg_6 *ClaireAny
          var try_7 EID
          try_7 = EVAL(ToList(OBJ(Core.F_CALL(C_args,ARGS(x.ToEID())))).At(1))
          if ErrorIn(try_7) {try_5 = try_7
          } else {
          arg_6 = ANY(try_7)
          try_5 = EID{Core.F__I_equal_any(v,arg_6).Id(),0}
          }
          } 
        if ErrorIn(try_5) {loop_4 = try_5
        } else {
        g0287I = ToBoolean(OBJ(try_5))
        if (g0287I == CTRUE) { 
          loop_4 = e.ToEID()
          } else {
          loop_4 = EID{CFALSE.Id(),0}
          } 
        }
        if ErrorIn(loop_4) {Result = loop_4
        break
        } else {
        }}
        } 
      } 
    return Result} 
  
// The EID go function for: Iterate @ list<type_expression>(but[tuple(any,any)], Variable, any) (throw: true) 
func E_Iterate_any2 (x EID,v EID,e EID) EID { 
    return F_Iterate_any2(ANY(x),To_Variable(OBJ(v)),ANY(e) )} 
  
/* The go function for: Iterate(x:any,v:Variable,e:any) [status=1] */
func F_Iterate_any3 (x *ClaireAny,v *ClaireVariable,e *ClaireAny) EID { 
    var Result EID
    { 
      var v *ClaireAny
      _ = v
      Result= EID{CFALSE.Id(),0}
      var v_support *ClaireList
      var try_1 EID
      { var arg_2 *ClaireAny
        var try_3 EID
        try_3 = EVAL(ToList(OBJ(Core.F_CALL(C_args,ARGS(x.ToEID())))).At(0))
        if ErrorIn(try_3) {try_1 = try_3
        } else {
        arg_2 = ANY(try_3)
        try_1 = Core.F_enumerate_any(arg_2)
        }
        } 
      if ErrorIn(try_1) {Result = try_1
      } else {
      v_support = ToList(OBJ(try_1))
      v_len := v_support.Length()
      for i_it := 0; i_it < v_len; i_it++ { 
        v = v_support.At(i_it)
        
        }
        } 
      } 
    if !ErrorIn(Result) {
    { 
      var v *ClaireAny
      _ = v
      Result= EID{CFALSE.Id(),0}
      var v_support *ClaireList
      var try_4 EID
      { var arg_5 *ClaireAny
        var try_6 EID
        try_6 = EVAL(ToList(OBJ(Core.F_CALL(C_args,ARGS(x.ToEID())))).At(1))
        if ErrorIn(try_6) {try_4 = try_6
        } else {
        arg_5 = ANY(try_6)
        try_4 = Core.F_enumerate_any(arg_5)
        }
        } 
      if ErrorIn(try_4) {Result = try_4
      } else {
      v_support = ToList(OBJ(try_4))
      v_len := v_support.Length()
      for i_it := 0; i_it < v_len; i_it++ { 
        v = v_support.At(i_it)
        
        }
        } 
      } 
    }
    return Result} 
  
// The EID go function for: Iterate @ list<type_expression>(/+[tuple(any,any)], Variable, any) (throw: true) 
func E_Iterate_any3 (x EID,v EID,e EID) EID { 
    return F_Iterate_any3(ANY(x),To_Variable(OBJ(v)),ANY(e) )} 
  