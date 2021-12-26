/***** CLAIRE Compilation of file /Users/ycaseau/Dropbox/src/clairev4.03/src/compile/ocontrol.cl 
         [version 4.0.04 / safety 5] Sunday 12-26-2021 17:16:12 *****/

package Optimize
import (_ "fmt"
	. "Kernel"
	"Core"
	"Language"
	"Reader"
)

//-------- dumb function to prevent import errors --------
func import_g0233() { 
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
/* {1} The go function for: c_type(self:Assign) [status=1] */
func F_c_type_Assign (self *Language.Assign ) EID { 
    var Result EID 
    Result = Core.F_CALL(C_c_type,ARGS(self.Arg.ToEID()))
    return Result} 
  
// The EID go function for: c_type @ Assign (throw: true) 
func E_c_type_Assign (self EID) EID { 
    return F_c_type_Assign(Language.To_Assign(OBJ(self)) )} 
  
// we must include the type checking if needed
/* {1} The go function for: c_code(self:Assign) [status=1] */
func F_c_code_Assign (self *Language.Assign ) EID { 
    var Result EID 
    { var v *ClaireAny   = self.ClaireVar
      { var x *ClaireAny   = self.Arg
        { var _Ztype *ClaireType  
          _ = _Ztype
          var try_1 EID 
          /*g_try(v2:"try_1",loop:false) */
          { var arg_2 *ClaireType  
            _ = arg_2
            var try_3 EID 
            /*g_try(v2:"try_3",loop:false) */
            try_3 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
            /* ERROR PROTECTION INSERTED (arg_2-try_1) */
            if ErrorIn(try_3) {try_1 = try_3
            } else {
            arg_2 = ToType(OBJ(try_3))
            try_1 = EID{F_Optimize_ptype_type(arg_2).Id(),0}
            }
            } 
          /* ERROR PROTECTION INSERTED (_Ztype-Result) */
          if ErrorIn(try_1) {Result = try_1
          } else {
          _Ztype = ToType(OBJ(try_1))
          /*g_try(v2:"Result",loop:true) */
          if (v.Isa.IsIn(C_Variable) != CTRUE) { 
            Result = Core.F_CALL(C_Compile_Cerror,ARGS(EID{MakeString("[213] ~S is not a variable").Id(),0},v.ToEID()))
            } else {
            Result = EID{CFALSE.Id(),0}
            } 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          /*g_try(v2:"Result",loop:true) */
          if (_Ztype.Included(ToType(OBJ(Core.F_CALL(C_range,ARGS(v.ToEID()))))) != CTRUE) { 
            var try_4 EID 
            /*g_try(v2:"try_4",loop:false) */
            try_4 = Core.F_CALL(C_Optimize_c_warn,ARGS(self.ClaireVar.ToEID(),x.ToEID(),EID{_Ztype.Id(),0}))
            /* ERROR PROTECTION INSERTED (x-Result) */
            if ErrorIn(try_4) {Result = try_4
            } else {
            x = ANY(try_4)
            Result = x.ToEID()
            }
            } else {
            Result = EID{CFALSE.Id(),0}
            } 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          { var _Zarg *ClaireAny  
            _ = _Zarg
            var try_5 EID 
            /*g_try(v2:"try_5",loop:false) */
            try_5 = F_Compile_c_strict_code_any(x,F_Compile_psort_any(ANY(Core.F_CALL(C_range,ARGS(v.ToEID())))))
            /* ERROR PROTECTION INSERTED (_Zarg-Result) */
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
/* {1} The go function for: c_type(self:Gassign) [status=1] */
func F_c_type_Gassign (self *Language.Gassign ) EID { 
    var Result EID 
    Result = Core.F_CALL(C_c_type,ARGS(self.Arg.ToEID()))
    return Result} 
  
// The EID go function for: c_type @ Gassign (throw: true) 
func E_c_type_Gassign (self EID) EID { 
    return F_c_type_Gassign(Language.To_Gassign(OBJ(self)) )} 
  
/* {1} The go function for: c_code(self:Gassign) [status=1] */
func F_c_code_Gassign (self *Language.Gassign ) EID { 
    var Result EID 
    { var _Zv *ClaireAny   = self.Arg
      { var _Ztype *ClaireType  
        _ = _Ztype
        var try_1 EID 
        /*g_try(v2:"try_1",loop:false) */
        { var arg_2 *ClaireType  
          _ = arg_2
          var try_3 EID 
          /*g_try(v2:"try_3",loop:false) */
          try_3 = Core.F_CALL(C_c_type,ARGS(_Zv.ToEID()))
          /* ERROR PROTECTION INSERTED (arg_2-try_1) */
          if ErrorIn(try_3) {try_1 = try_3
          } else {
          arg_2 = ToType(OBJ(try_3))
          try_1 = EID{F_Optimize_ptype_type(arg_2).Id(),0}
          }
          } 
        /* ERROR PROTECTION INSERTED (_Ztype-Result) */
        if ErrorIn(try_1) {Result = try_1
        } else {
        _Ztype = ToType(OBJ(try_1))
        /*g_try(v2:"Result",loop:true) */
        if (F_boolean_I_any(self.ClaireVar.Range.Id()).Id() != CTRUE.Id()) { 
          Result = Core.F_CALL(C_Compile_Cerror,ARGS(EID{MakeString("[214] cannot assign ~S").Id(),0},EID{self.Id(),0}))
          } else {
          Result = EID{CFALSE.Id(),0}
          } 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        /*g_try(v2:"Result",loop:true) */
        if (_Ztype.Included(self.ClaireVar.Range) != CTRUE) { 
          var try_4 EID 
          /*g_try(v2:"try_4",loop:false) */
          try_4 = Core.F_CALL(C_c_code,ARGS(_Zv.ToEID(),EID{C_any.Id(),0}))
          /* ERROR PROTECTION INSERTED (_Zv-Result) */
          if ErrorIn(try_4) {Result = try_4
          } else {
          _Zv = ANY(try_4)
          Result = _Zv.ToEID()
          }
          } else {
          Result = EID{CFALSE.Id(),0}
          } 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        { var _CL_obj *Language.Gassign   = Language.To_Gassign(new(Language.Gassign).Is(Language.C_Gassign))
          /*g_try(v2:"Result",loop:true) */
          { 
            var va_arg1 *Language.Gassign  
            var va_arg2 *Core.GlobalVariable  
            va_arg1 = _CL_obj
            var try_5 EID 
            /*g_try(v2:"try_5",loop:false) */
            try_5 = Core.F_CALL(C_c_code,ARGS(EID{self.ClaireVar.Id(),0}))
            /* ERROR PROTECTION INSERTED (va_arg2-Result) */
            if ErrorIn(try_5) {Result = try_5
            } else {
            va_arg2 = Core.ToGlobalVariable(OBJ(try_5))
            va_arg1.ClaireVar = va_arg2
            /*global_variable->global_variable*/Result = EID{va_arg2.Id(),0}
            }
            } 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          /*g_try(v2:"Result",loop:true) */
          { 
            var va_arg1 *Language.Gassign  
            var va_arg2 *ClaireAny  
            va_arg1 = _CL_obj
            var try_6 EID 
            /*g_try(v2:"try_6",loop:false) */
            if (F_Compile_nativeVar_ask_global_variable(self.ClaireVar) == CTRUE) { 
              try_6 = F_Compile_c_strict_code_any(_Zv,F_Compile_psort_any(self.ClaireVar.Range.Id()))
              } else {
              try_6 = Core.F_CALL(C_c_code,ARGS(_Zv.ToEID(),EID{C_any.Id(),0}))
              } 
            /* ERROR PROTECTION INSERTED (va_arg2-Result) */
            if ErrorIn(try_6) {Result = try_6
            } else {
            va_arg2 = ANY(try_6)
            va_arg1.Arg = va_arg2
            /*any->any*/Result = va_arg2.ToEID()
            }
            } 
          /* ERROR PROTECTION INSERTED (Result-Result) */
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
/* {1} The go function for: c_type(self:And) [status=0] */
func F_c_type_And (self *Language.And ) *ClaireType  { 
    return  ToType(C_boolean.Id())
    } 
  
// The EID go function for: c_type @ And (throw: false) 
func E_c_type_And (self EID) EID { 
    return EID{F_c_type_And(Language.To_And(OBJ(self)) ).Id(),0}} 
  
/* {1} The go function for: c_code(self:And) [status=1] */
func F_c_code_And (self *Language.And ) EID { 
    var Result EID 
    { var _CL_obj *Language.And   = Language.To_And(new(Language.And).Is(Language.C_And))
      /*g_try(v2:"Result",loop:true) */
      { 
        var va_arg1 *Language.And  
        var va_arg2 *ClaireList  
        va_arg1 = _CL_obj
        var try_1 EID 
        /*g_try(v2:"try_1",loop:false) */
        { 
          var v_list4 *ClaireList  
          var x *ClaireAny  
          var v_local4 *ClaireAny  
          v_list4 = self.Args
          try_1 = EID{CreateList(ToType(CEMPTY.Id()),v_list4.Length()).Id(),0}
          for CLcount := 0; CLcount < v_list4.Length(); CLcount++{ 
            x = v_list4.At(CLcount)
            var try_2 EID 
            /*g_try(v2:"try_2",loop:tuple("try_1", EID)) */
            /*g_try(v2:"try_2",loop:tuple("try_1", EID)) */
            var g0234I *ClaireBoolean  
            var try_3 EID 
            /*g_try(v2:"try_3",loop:false) */
            { var arg_4 *ClaireType  
              _ = arg_4
              var try_5 EID 
              /*g_try(v2:"try_5",loop:false) */
              try_5 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
              /* ERROR PROTECTION INSERTED (arg_4-try_3) */
              if ErrorIn(try_5) {try_3 = try_5
              } else {
              arg_4 = ToType(OBJ(try_5))
              try_3 = EID{Equal(arg_4.Id(),C_void.Id()).Id(),0}
              }
              } 
            /* ERROR PROTECTION INSERTED (g0234I-try_2) */
            if ErrorIn(try_3) {try_2 = try_3
            } else {
            g0234I = ToBoolean(OBJ(try_3))
            if (g0234I == CTRUE) { 
              try_2 = Core.F_CALL(C_Compile_Cerror,ARGS(EID{MakeString("[206] void ~S in ~S").Id(),0},x.ToEID(),EID{self.Id(),0}))
              } else {
              try_2 = EID{CFALSE.Id(),0}
              } 
            }
            /* ERROR PROTECTION INSERTED (try_2-try_2) */
            if ErrorIn(try_2) {try_1 = try_2
            break
            } else {
            /*g_try(v2:"try_2",loop:tuple("try_1", EID)) */
            try_2 = F_Optimize_c_boolean_any(x)
            /* ERROR PROTECTION INSERTED (try_2-try_2) */
            if ErrorIn(try_2) {try_1 = try_2
            break
            } else {
            }}
            /* ERROR PROTECTION INSERTED (v_local4-try_1) */
            if ErrorIn(try_2) {try_1 = try_2
            break
            } else {
            v_local4 = ANY(try_2)
            ToList(OBJ(try_1)).PutAt(CLcount,v_local4)
            } 
          }
          } 
        /* ERROR PROTECTION INSERTED (va_arg2-Result) */
        if ErrorIn(try_1) {Result = try_1
        } else {
        va_arg2 = ToList(OBJ(try_1))
        va_arg1.Args = va_arg2
        /*list->list*/Result = EID{va_arg2.Id(),0}
        }
        } 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = EID{_CL_obj.Id(),0}
      }
      } 
    return Result} 
  
// The EID go function for: c_code @ And (throw: true) 
func E_c_code_And (self EID) EID { 
    return F_c_code_And(Language.To_And(OBJ(self)) )} 
  
/* {1} The go function for: c_type(self:Or) [status=0] */
func F_c_type_Or (self *Language.Or ) *ClaireType  { 
    return  ToType(C_boolean.Id())
    } 
  
// The EID go function for: c_type @ Or (throw: false) 
func E_c_type_Or (self EID) EID { 
    return EID{F_c_type_Or(Language.To_Or(OBJ(self)) ).Id(),0}} 
  
/* {1} The go function for: c_code(self:Or) [status=1] */
func F_c_code_Or (self *Language.Or ) EID { 
    var Result EID 
    { var _CL_obj *Language.Or   = Language.To_Or(new(Language.Or).Is(Language.C_Or))
      /*g_try(v2:"Result",loop:true) */
      { 
        var va_arg1 *Language.Or  
        var va_arg2 *ClaireList  
        va_arg1 = _CL_obj
        var try_1 EID 
        /*g_try(v2:"try_1",loop:false) */
        { 
          var v_list4 *ClaireList  
          var x *ClaireAny  
          var v_local4 *ClaireAny  
          v_list4 = self.Args
          try_1 = EID{CreateList(ToType(CEMPTY.Id()),v_list4.Length()).Id(),0}
          for CLcount := 0; CLcount < v_list4.Length(); CLcount++{ 
            x = v_list4.At(CLcount)
            var try_2 EID 
            /*g_try(v2:"try_2",loop:tuple("try_1", EID)) */
            /*g_try(v2:"try_2",loop:tuple("try_1", EID)) */
            var g0235I *ClaireBoolean  
            var try_3 EID 
            /*g_try(v2:"try_3",loop:false) */
            { var arg_4 *ClaireType  
              _ = arg_4
              var try_5 EID 
              /*g_try(v2:"try_5",loop:false) */
              try_5 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
              /* ERROR PROTECTION INSERTED (arg_4-try_3) */
              if ErrorIn(try_5) {try_3 = try_5
              } else {
              arg_4 = ToType(OBJ(try_5))
              try_3 = EID{Equal(arg_4.Id(),C_void.Id()).Id(),0}
              }
              } 
            /* ERROR PROTECTION INSERTED (g0235I-try_2) */
            if ErrorIn(try_3) {try_2 = try_3
            } else {
            g0235I = ToBoolean(OBJ(try_3))
            if (g0235I == CTRUE) { 
              try_2 = Core.F_CALL(C_Compile_Cerror,ARGS(EID{MakeString("[206] void ~S in ~S").Id(),0},x.ToEID(),EID{self.Id(),0}))
              } else {
              try_2 = EID{CFALSE.Id(),0}
              } 
            }
            /* ERROR PROTECTION INSERTED (try_2-try_2) */
            if ErrorIn(try_2) {try_1 = try_2
            break
            } else {
            /*g_try(v2:"try_2",loop:tuple("try_1", EID)) */
            try_2 = F_Optimize_c_boolean_any(x)
            /* ERROR PROTECTION INSERTED (try_2-try_2) */
            if ErrorIn(try_2) {try_1 = try_2
            break
            } else {
            }}
            /* ERROR PROTECTION INSERTED (v_local4-try_1) */
            if ErrorIn(try_2) {try_1 = try_2
            break
            } else {
            v_local4 = ANY(try_2)
            ToList(OBJ(try_1)).PutAt(CLcount,v_local4)
            } 
          }
          } 
        /* ERROR PROTECTION INSERTED (va_arg2-Result) */
        if ErrorIn(try_1) {Result = try_1
        } else {
        va_arg2 = ToList(OBJ(try_1))
        va_arg1.Args = va_arg2
        /*list->list*/Result = EID{va_arg2.Id(),0}
        }
        } 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = EID{_CL_obj.Id(),0}
      }
      } 
    return Result} 
  
// The EID go function for: c_code @ Or (throw: true) 
func E_c_code_Or (self EID) EID { 
    return F_c_code_Or(Language.To_Or(OBJ(self)) )} 
  
//---------------- quote and return -------------------------------------
/* {1} The go function for: c_type(self:Quote) [status=0] */
func F_c_type_Quote (self *Language.Quote ) *ClaireType  { 
    return  ToType(self.Arg.Isa.Id())
    } 
  
// The EID go function for: c_type @ Quote (throw: false) 
func E_c_type_Quote (self EID) EID { 
    return EID{F_c_type_Quote(Language.To_Quote(OBJ(self)) ).Id(),0}} 
  
/* {1} The go function for: c_code(self:Quote) [status=1] */
func F_c_code_Quote (self *Language.Quote ) EID { 
    var Result EID 
    Result = Core.F_CALL(C_Compile_Cerror,ARGS(EID{MakeString("[internal] optimization of quote not implemented yet! ~S").Id(),0},EID{self.Id(),0}))
    return Result} 
  
// The EID go function for: c_code @ Quote (throw: true) 
func E_c_code_Quote (self EID) EID { 
    return F_c_code_Quote(Language.To_Quote(OBJ(self)) )} 
  
/* {1} The go function for: c_type(self:Return) [status=0] */
func F_c_type_Return (self *Language.Return ) *ClaireType  { 
    return  ToType(C_any.Id())
    } 
  
// The EID go function for: c_type @ Return (throw: false) 
func E_c_type_Return (self EID) EID { 
    return EID{F_c_type_Return(Language.To_Return(OBJ(self)) ).Id(),0}} 
  
/* {1} The go function for: c_code(self:Return) [status=1] */
func F_c_code_Return (self *Language.Return ) EID { 
    var Result EID 
    { var arg_1 *ClaireAny  
      _ = arg_1
      var try_2 EID 
      /*g_try(v2:"try_2",loop:false) */
      try_2 = Core.F_CALL(C_c_code,ARGS(self.Arg.ToEID(),EID{C_any.Id(),0}))
      /* ERROR PROTECTION INSERTED (arg_1-Result) */
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
/* {1} The go function for: c_type(self:Handle) [status=1] */
func F_c_type_Handle (self *Language.ClaireHandle ) EID { 
    var Result EID 
    { var arg_1 *ClaireType  
      _ = arg_1
      var try_3 EID 
      /*g_try(v2:"try_3",loop:false) */
      try_3 = Core.F_CALL(C_c_type,ARGS(self.Arg.ToEID()))
      /* ERROR PROTECTION INSERTED (arg_1-Result) */
      if ErrorIn(try_3) {Result = try_3
      } else {
      arg_1 = ToType(OBJ(try_3))
      { var arg_2 *ClaireType  
        _ = arg_2
        var try_4 EID 
        /*g_try(v2:"try_4",loop:false) */
        try_4 = Core.F_CALL(C_c_type,ARGS(self.Other.ToEID()))
        /* ERROR PROTECTION INSERTED (arg_2-Result) */
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
  
/* {1} The go function for: c_code(self:Handle,s:class) [status=1] */
func F_c_code_Handle (self *Language.ClaireHandle ,s *ClaireClass ) EID { 
    var Result EID 
    { var x *Language.ClaireHandle  
      var try_1 EID 
      /*g_try(v2:"try_1",loop:false) */
      { var arg_2 *ClaireAny  
        _ = arg_2
        var try_4 EID 
        /*g_try(v2:"try_4",loop:false) */
        try_4 = Core.F_CALL(C_c_code,ARGS(self.Arg.ToEID(),EID{s.Id(),0}))
        /* ERROR PROTECTION INSERTED (arg_2-try_1) */
        if ErrorIn(try_4) {try_1 = try_4
        } else {
        arg_2 = ANY(try_4)
        { var arg_3 *ClaireAny  
          _ = arg_3
          var try_5 EID 
          /*g_try(v2:"try_5",loop:false) */
          try_5 = Core.F_CALL(C_c_code,ARGS(self.Other.ToEID(),EID{s.Id(),0}))
          /* ERROR PROTECTION INSERTED (arg_3-try_1) */
          if ErrorIn(try_5) {try_1 = try_5
          } else {
          arg_3 = ANY(try_5)
          try_1 = Language.C_Handle.Make(C_any.Id(),arg_2,arg_3).ToEID()
          }
          } 
        }
        } 
      /* ERROR PROTECTION INSERTED (x-Result) */
      if ErrorIn(try_1) {Result = try_1
      } else {
      x = Language.To_ClaireHandle(OBJ(try_1))
      x.Test = self.Test
      /*any->any*/Result = EID{x.Id(),0}
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
/* {1} The go function for: c_type(self:Cast) [status=0] */
func F_c_type_Cast (self *Language.Cast ) *ClaireType  { 
    return  self.SetArg
    } 
  
// The EID go function for: c_type @ Cast (throw: false) 
func E_c_type_Cast (self EID) EID { 
    return EID{F_c_type_Cast(Language.To_Cast(OBJ(self)) ).Id(),0}} 
  
// insert dynamic types (check_in) when we see a claire cast
// CLAIRE 4 : when we decide to drop the cast (safety), we generate a C_cast
/* {1} The go function for: c_code(self:Cast) [status=1] */
func F_c_code_Cast (self *Language.Cast ) EID { 
    var Result EID 
    { var y *ClaireType   = self.SetArg
      { var ftype *ClaireClass   = F_Compile_psort_any(y.Id())
        var g0237I *ClaireBoolean  
        if (y.Isa.IsIn(C_Param) == CTRUE) { 
          { var g0236 *ClaireParam   = To_Param(y.Id())
            g0237I = MakeBoolean(((g0236.Arg.Id() == C_list.Id()) || 
                (g0236.Arg.Id() == C_set.Id())) && (C_set.Id() == g0236.Args.At(1-1).Isa.Id()))
            } 
          } else {
          g0237I = CFALSE
          } 
        if (g0237I == CTRUE) { 
          { var utype *ClaireAny  
            var try_1 EID 
            /*g_try(v2:"try_1",loop:false) */
            try_1 = Core.F_the_type(ToType(To_Param(y.Id()).Args.At(1-1)))
            /* ERROR PROTECTION INSERTED (utype-Result) */
            if ErrorIn(try_1) {Result = try_1
            } else {
            utype = ANY(try_1)
            var g0238I *ClaireBoolean  
            var try_2 EID 
            /*g_try(v2:"try_2",loop:false) */
            { 
              /* Or stat: v="try_2", loop=false */
              var v_or6 *ClaireBoolean  
              
              /* Or stat: try = @ any(@ @ type(c_type((arg @ Cast(self))),of),utype) with try:true, v="try_2", loop=false */
              var try_3 EID 
              /*g_try(v2:"try_3",loop:false) */
              { var arg_4 *ClaireType  
                _ = arg_4
                var try_5 EID 
                /*g_try(v2:"try_5",loop:false) */
                { var arg_6 *ClaireType  
                  _ = arg_6
                  var try_7 EID 
                  /*g_try(v2:"try_7",loop:false) */
                  try_7 = Core.F_CALL(C_c_type,ARGS(self.Arg.ToEID()))
                  /* ERROR PROTECTION INSERTED (arg_6-try_5) */
                  if ErrorIn(try_7) {try_5 = try_7
                  } else {
                  arg_6 = ToType(OBJ(try_7))
                  try_5 = EID{arg_6.At(C_of).Id(),0}
                  }
                  } 
                /* ERROR PROTECTION INSERTED (arg_4-try_3) */
                if ErrorIn(try_5) {try_3 = try_5
                } else {
                arg_4 = ToType(OBJ(try_5))
                try_3 = EID{Equal(arg_4.Id(),utype).Id(),0}
                }
                } 
              /* ERROR PROTECTION INSERTED (v_or6-try_2) */
              if ErrorIn(try_3) {try_2 = try_3
              } else {
              v_or6 = ToBoolean(OBJ(try_3))
              if (v_or6 == CTRUE) {try_2 = EID{CTRUE.Id(),0}
              } else { 
                /* Or stat: try >= @ integer(safety @ meta_compiler(compiler),2) with try:false, v="try_2", loop=false */
                v_or6 = F__sup_equal_integer(C_compiler.Safety,2)
                if (v_or6 == CTRUE) {try_2 = EID{CTRUE.Id(),0}
                } else { 
                  try_2 = EID{CFALSE.Id(),0}} 
                } 
              }
              } 
            /* ERROR PROTECTION INSERTED (g0238I-Result) */
            if ErrorIn(try_2) {Result = try_2
            } else {
            g0238I = ToBoolean(OBJ(try_2))
            if (g0238I == CTRUE) { 
              Result = Core.F_CALL(C_c_code,ARGS(self.Arg.ToEID(),EID{ftype.Id(),0}))
              } else {
              { var arg_8 *Language.Call  
                _ = arg_8
                { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                  _CL_obj.Selector = Core.C_check_in
                  /*property->property*/_CL_obj.Args = MakeConstantList(self.Arg,To_Param(y.Id()).Arg.Id(),utype)
                  /*list->list*/arg_8 = _CL_obj
                  } 
                Result = Core.F_CALL(C_c_code,ARGS(EID{arg_8.Id(),0},EID{ftype.Id(),0}))
                } 
              } 
            }
            }
            } 
          } else {
          var g0239I *ClaireBoolean  
          var try_9 EID 
          /*g_try(v2:"try_9",loop:false) */
          { var arg_10 *ClaireType  
            _ = arg_10
            var try_11 EID 
            /*g_try(v2:"try_11",loop:false) */
            try_11 = Core.F_CALL(C_c_type,ARGS(self.Arg.ToEID()))
            /* ERROR PROTECTION INSERTED (arg_10-try_9) */
            if ErrorIn(try_11) {try_9 = try_11
            } else {
            arg_10 = ToType(OBJ(try_11))
            try_9 = EID{arg_10.Included(y).Id(),0}
            }
            } 
          /* ERROR PROTECTION INSERTED (g0239I-Result) */
          if ErrorIn(try_9) {Result = try_9
          } else {
          g0239I = ToBoolean(OBJ(try_9))
          if (g0239I == CTRUE) { 
            Result = Core.F_CALL(C_c_code,ARGS(self.Arg.ToEID(),EID{ftype.Id(),0}))
            }  else if (C_compiler.Safety >= 2) { 
            { var _CL_obj *Compile_CCast   = To_Compile_CCast(new(Compile_CCast).Is(C_Compile_C_cast))
              /*g_try(v2:"Result",loop:true) */
              { 
                var va_arg1 *Compile_CCast  
                var va_arg2 *ClaireAny  
                va_arg1 = _CL_obj
                var try_12 EID 
                /*g_try(v2:"try_12",loop:false) */
                try_12 = Core.F_CALL(C_c_code,ARGS(self.Arg.ToEID(),EID{ftype.Id(),0}))
                /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                if ErrorIn(try_12) {Result = try_12
                } else {
                va_arg2 = ANY(try_12)
                va_arg1.Arg = va_arg2
                /*any->any*/Result = va_arg2.ToEID()
                }
                } 
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              _CL_obj.SetArg = ToClass(y.Id())
              /*class->class*/Result = EID{_CL_obj.Id(),0}
              }
              } 
            } else {
            { var arg_13 *Language.Call  
              _ = arg_13
              { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                _CL_obj.Selector = Core.C_check_in
                /*property->property*/_CL_obj.Args = MakeConstantList(self.Arg,y.Id())
                /*list->list*/arg_13 = _CL_obj
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
/* {1} The go function for: c_type(self:Super) [status=1] */
func F_c_type_Super (self *Language.Super ) EID { 
    var Result EID 
    { var _Ztype *ClaireList  
      var try_1 EID 
      /*g_try(v2:"try_1",loop:false) */
      { 
        var v_list3 *ClaireList  
        var x *ClaireAny  
        var v_local3 *ClaireAny  
        v_list3 = self.Args
        try_1 = EID{CreateList(ToType(CEMPTY.Id()),v_list3.Length()).Id(),0}
        for CLcount := 0; CLcount < v_list3.Length(); CLcount++{ 
          x = v_list3.At(CLcount)
          var try_2 EID 
          /*g_try(v2:"try_2",loop:tuple("try_1", EID)) */
          try_2 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
          /* ERROR PROTECTION INSERTED (v_local3-try_1) */
          if ErrorIn(try_2) {try_1 = try_2
          break
          } else {
          v_local3 = ANY(try_2)
          ToList(OBJ(try_1)).PutAt(CLcount,v_local3)
          } 
        }
        } 
      /* ERROR PROTECTION INSERTED (_Ztype-Result) */
      if ErrorIn(try_1) {Result = try_1
      } else {
      _Ztype = ToList(OBJ(try_1))
      { var s *ClaireProperty   = self.Selector
        ToArray(_Ztype.Id()).NthPut(1,self.CastTo.Id())
        { var prop *ClaireAny  
          if (s.Open == 3) { 
            prop = CNIL.Id()
            } else {
            prop = F_Optimize_restriction_I_class(self.CastTo.Class_I(),s.Definition,_Ztype)
            } 
          if (C_slot.Id() == prop.Isa.Id()) { 
            { var g0240 *ClaireSlot   = ToSlot(prop)
              _ = g0240
              Result = EID{g0240.Range.Id(),0}
              } 
            }  else if (C_method.Id() == prop.Isa.Id()) { 
            { var g0241 *ClaireMethod   = ToMethod(prop)
              _ = g0241
              Result = F_Optimize_use_range_method(g0241,_Ztype)
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
/* {1} The go function for: c_code(self:Super) [status=1] */
func F_c_code_Super (self *Language.Super ) EID { 
    var Result EID 
    { var s *ClaireProperty   = self.Selector
      { var l *ClaireList   = self.Args
        { var _Ztype *ClaireList  
          var try_1 EID 
          /*g_try(v2:"try_1",loop:false) */
          { 
            var v_list5 *ClaireList  
            var x *ClaireAny  
            var v_local5 *ClaireAny  
            v_list5 = self.Args
            try_1 = EID{CreateList(ToType(CEMPTY.Id()),v_list5.Length()).Id(),0}
            for CLcount := 0; CLcount < v_list5.Length(); CLcount++{ 
              x = v_list5.At(CLcount)
              var try_2 EID 
              /*g_try(v2:"try_2",loop:tuple("try_1", EID)) */
              try_2 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
              /* ERROR PROTECTION INSERTED (v_local5-try_1) */
              if ErrorIn(try_2) {try_1 = try_2
              break
              } else {
              v_local5 = ANY(try_2)
              ToList(OBJ(try_1)).PutAt(CLcount,v_local5)
              } 
            }
            } 
          /* ERROR PROTECTION INSERTED (_Ztype-Result) */
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
              { var g0243 *ClaireSlot   = ToSlot(prop)
                { var _CL_obj *Language.CallSlot   = Language.To_CallSlot(new(Language.CallSlot).Is(Language.C_Call_slot))
                  _CL_obj.Selector = g0243
                  /*slot->slot*//*g_try(v2:"Result",loop:true) */
                  { 
                    var va_arg1 *Language.CallSlot  
                    var va_arg2 *ClaireAny  
                    va_arg1 = _CL_obj
                    var try_3 EID 
                    /*g_try(v2:"try_3",loop:false) */
                    try_3 = Core.F_CALL(C_c_code,ARGS(l.At(1-1).ToEID(),EID{F_Compile_psort_any(Core.F_domain_I_restriction(ToRestriction(g0243.Id())).Id()).Id(),0}))
                    /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                    if ErrorIn(try_3) {Result = try_3
                    } else {
                    va_arg2 = ANY(try_3)
                    va_arg1.Arg = va_arg2
                    /*any->any*/Result = va_arg2.ToEID()
                    }
                    } 
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  _CL_obj.Test = MakeBoolean((g0243.Range.Contains(g0243.Default) != CTRUE) && (C_compiler.Safety < 5))
                  /*boolean->boolean*/Result = EID{_CL_obj.Id(),0}
                  }
                  } 
                } 
              }  else if (C_method.Id() == prop.Isa.Id()) { 
              { var g0244 *ClaireMethod   = ToMethod(prop)
                _ = g0244
                Result = F_Optimize_c_code_method_method1(g0244,l,_Ztype)
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
/* {1} The go function for: self_print(self:Call_function2) [status=1] */
func (self *Optimize_CallFunction2 ) SelfPrint () EID { 
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
  
// The EID go function for: self_print @ Call_function2 (throw: true) 
func E_self_print_Call_function2 (self EID) EID { 
    return To_Optimize_CallFunction2(OBJ(self)).SelfPrint( )} 
  
/* {1} The go function for: c_type(self:Call_function2) [status=0] */
func (self *Optimize_CallFunction2 ) CType () *ClaireType  { 
    return  ToType(C_any.Id())
    } 
  
// The EID go function for: c_type @ Call_function2 (throw: false) 
func E_c_type_Call_function2 (self EID) EID { 
    return EID{To_Optimize_CallFunction2(OBJ(self)).CType( ).Id(),0}} 
  
/* {1} The go function for: c_code(self:Call_function2) [status=1] */
func F_c_code_Call_function2 (self *Optimize_CallFunction2 ) EID { 
    var Result EID 
    { var _CL_obj *Optimize_CallFunction2   = To_Optimize_CallFunction2(new(Optimize_CallFunction2).Is(C_Optimize_Call_function2))
      _CL_obj.Arg = self.Arg
      /*function->function*//*g_try(v2:"Result",loop:true) */
      { 
        var va_arg1 *Optimize_CallFunction2  
        var va_arg2 *ClaireList  
        va_arg1 = _CL_obj
        var try_1 EID 
        /*g_try(v2:"try_1",loop:false) */
        { 
          var v_list4 *ClaireList  
          var x *ClaireAny  
          var v_local4 *ClaireAny  
          v_list4 = self.Args
          try_1 = EID{CreateList(ToType(CEMPTY.Id()),v_list4.Length()).Id(),0}
          for CLcount := 0; CLcount < v_list4.Length(); CLcount++{ 
            x = v_list4.At(CLcount)
            var try_2 EID 
            /*g_try(v2:"try_2",loop:tuple("try_1", EID)) */
            try_2 = Core.F_CALL(C_c_code,ARGS(x.ToEID(),EID{C_any.Id(),0}))
            /* ERROR PROTECTION INSERTED (v_local4-try_1) */
            if ErrorIn(try_2) {try_1 = try_2
            break
            } else {
            v_local4 = ANY(try_2)
            ToList(OBJ(try_1)).PutAt(CLcount,v_local4)
            } 
          }
          } 
        /* ERROR PROTECTION INSERTED (va_arg2-Result) */
        if ErrorIn(try_1) {Result = try_1
        } else {
        va_arg2 = ToList(OBJ(try_1))
        va_arg1.Args = va_arg2
        /*list->list*/Result = EID{va_arg2.Id(),0}
        }
        } 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = EID{_CL_obj.Id(),0}
      }
      } 
    return Result} 
  
// The EID go function for: c_code @ Call_function2 (throw: true) 
func E_c_code_Call_function2 (self EID) EID { 
    return F_c_code_Call_function2(To_Optimize_CallFunction2(OBJ(self)) )} 
  
// ASSERT & trace
/* {1} The go function for: c_code(self:Assert) [status=1] */
func F_c_code_Assert (self *Language.Assert ) EID { 
    var Result EID 
    if ((C_compiler.Safety == 0) || 
        (C_compiler.Debug_ask.Length() != 0)) { 
      { var arg_1 *ClaireObject  
        _ = arg_1
        { var arg_2 *Language.Call  
          _ = arg_2
          { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
            _CL_obj.Selector = Core.C_not
            /*property->property*/_CL_obj.Args = MakeConstantList(self.Args.At(1-1))
            /*list->list*/arg_2 = _CL_obj
            } 
          { var arg_3 *Language.Call  
            _ = arg_3
            { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
              _CL_obj.Selector = Core.C_Core_tformat
              /*property->property*/{ 
                var va_arg1 *Language.Call  
                var va_arg2 *ClaireList  
                va_arg1 = _CL_obj
                { 
                  var v_bag_arg *ClaireAny  
                  va_arg2= ToType(CEMPTY.Id()).EmptyList()
                  va_arg2.AddFast(MakeString("Assertion violation in ~A line ~A\n").Id())
                  va_arg2.AddFast(MakeInteger(0).Id())
                  { var _CL_obj *Language.List   = Language.To_List(new(Language.List).Is(Language.C_List))
                    _CL_obj.Args = MakeConstantList((self.External).Id(),MakeInteger(self.Index).Id())
                    /*list->list*/v_bag_arg = _CL_obj.Id()
                    } 
                  va_arg2.AddFast(v_bag_arg)} 
                va_arg1.Args = va_arg2
                /*list->list*/} 
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
/* {1} The go function for: c_code(self:Trace) [status=1] */
func F_c_code_Trace (self *Language.Trace ) EID { 
    var Result EID 
    { var a *ClaireList   = self.Args
      var g0246I *ClaireBoolean  
      var try_1 EID 
      /*g_try(v2:"try_1",loop:false) */
      { 
        var v_and3 *ClaireBoolean  
        
        v_and3 = Equal(MakeInteger(a.Length()).Id(),MakeInteger(1).Id())
        if (v_and3 == CFALSE) {try_1 = EID{CFALSE.Id(),0}
        } else { 
          var try_2 EID 
          /*g_try(v2:"try_2",loop:false) */
          { var arg_3 *ClaireType  
            _ = arg_3
            var try_4 EID 
            /*g_try(v2:"try_4",loop:false) */
            try_4 = Core.F_CALL(C_c_type,ARGS(a.At(1-1).ToEID()))
            /* ERROR PROTECTION INSERTED (arg_3-try_2) */
            if ErrorIn(try_4) {try_2 = try_4
            } else {
            arg_3 = ToType(OBJ(try_4))
            try_2 = EID{arg_3.Included(ToType(C_integer.Id())).Id(),0}
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
      /* ERROR PROTECTION INSERTED (g0246I-Result) */
      if ErrorIn(try_1) {Result = try_1
      } else {
      g0246I = ToBoolean(OBJ(try_1))
      if (g0246I == CTRUE) { 
        { var arg_5 *Language.Call  
          _ = arg_5
          { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
            _CL_obj.Selector = Core.C_write
            /*property->property*/_CL_obj.Args = MakeConstantList(C_verbose.Id(),ClEnv.Id(),a.At(1-1))
            /*list->list*/arg_5 = _CL_obj
            } 
          Result = Core.F_CALL(C_c_code,ARGS(EID{arg_5.Id(),0}))
          } 
        } else {
        var g0247I *ClaireBoolean  
        { 
          var v_and4 *ClaireBoolean  
          
          v_and4 = Core.F__sup_integer(a.Length(),1)
          if (v_and4 == CFALSE) {g0247I = CFALSE
          } else { 
            v_and4 = Equal(C_string.Id(),a.At(2-1).Isa.Id())
            if (v_and4 == CFALSE) {g0247I = CFALSE
            } else { 
              { 
                /* Or stat: v="v_and4", loop=true */
                var v_or7 *ClaireBoolean  
                
                /* Or stat: try != @ any(length @ list(debug? @ meta_compiler(compiler)),0) with try:false, v="v_and4", loop=true */
                v_or7 = Core.F__I_equal_any(MakeInteger(C_compiler.Debug_ask.Length()).Id(),MakeInteger(0).Id())
                if (v_or7 == CTRUE) {v_and4 = CTRUE
                } else { 
                  /* Or stat: try <try <= @ integer(eval @ list<type_expression>(any)(nth @ list(a,1)),max @ integer(2,verbose @ environment(<environment>))) catch any true:boolean> with try:false, v="v_and4", loop=true */
                  { 
                    var v_or7_H EID 
                    h_index := ClEnv.Index
                    h_base := ClEnv.Base
                    { var arg_6 *ClaireAny  
                      _ = arg_6
                      var try_7 EID 
                      /*g_try(v2:"try_7",loop:false) */
                      try_7 = EVAL(a.At(1-1))
                      /* ERROR PROTECTION INSERTED (arg_6-v_or7_H) */
                      if ErrorIn(try_7) {v_or7_H = try_7
                      } else {
                      arg_6 = ANY(try_7)
                      v_or7_H = EID{Core.F__inf_equal_integer(ToInteger(arg_6).Value,Reader.F_max_integer(2,ClEnv.Verbose)).Id(),0}
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
              if (v_and4 == CFALSE) {g0247I = CFALSE
              } else { 
                g0247I = CTRUE} 
              } 
            } 
          } 
        if (g0247I == CTRUE) { 
          { var _Zc *Language.Call  
            { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
              _CL_obj.Selector = Core.C_Core_tformat
              /*property->property*/{ 
                var va_arg1 *Language.Call  
                var va_arg2 *ClaireList  
                va_arg1 = _CL_obj
                { 
                  var v_bag_arg *ClaireAny  
                  va_arg2= ToType(CEMPTY.Id()).EmptyList()
                  va_arg2.AddFast(a.At(2-1))
                  va_arg2.AddFast(a.At(1-1))
                  { var _CL_obj *Language.List   = Language.To_List(new(Language.List).Is(Language.C_List))
                    _CL_obj.Args = a.Copy().Skip(2)
                    /*list->list*/v_bag_arg = _CL_obj.Id()
                    } 
                  va_arg2.AddFast(v_bag_arg)} 
                va_arg1.Args = va_arg2
                /*list->list*/} 
              _Zc = _CL_obj
              } 
            { var arg_8 *ClaireObject  
              _ = arg_8
              if (C_integer.Id() != a.At(1-1).Isa.Id()) { 
                { var arg_9 *Language.Call  
                  _ = arg_9
                  { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                    _CL_obj.Selector = ToProperty(C__inf_equal.Id())
                    /*property->property*/{ 
                      var va_arg1 *Language.Call  
                      var va_arg2 *ClaireList  
                      va_arg1 = _CL_obj
                      { 
                        var v_bag_arg *ClaireAny  
                        va_arg2= ToType(CEMPTY.Id()).EmptyList()
                        va_arg2.AddFast(a.At(1-1))
                        { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                          _CL_obj.Selector = C_verbose
                          /*property->property*/_CL_obj.Args = MakeConstantList(ClEnv.Id())
                          /*list->list*/v_bag_arg = _CL_obj.Id()
                          } 
                        va_arg2.AddFast(v_bag_arg)} 
                      va_arg1.Args = va_arg2
                      /*list->list*/} 
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
  
/* {1} The go function for: c_type(self:Assert) [status=0] */
func F_c_type_Assert (self *Language.Assert ) *ClaireType  { 
    return  ToType(C_any.Id())
    } 
  
// The EID go function for: c_type @ Assert (throw: false) 
func E_c_type_Assert (self EID) EID { 
    return EID{F_c_type_Assert(Language.To_Assert(OBJ(self)) ).Id(),0}} 
  
/* {1} The go function for: c_type(self:Trace) [status=0] */
func F_c_type_Trace (self *Language.Trace ) *ClaireType  { 
    return  ToType(C_any.Id())
    } 
  
// The EID go function for: c_type @ Trace (throw: false) 
func E_c_type_Trace (self EID) EID { 
    return EID{F_c_type_Trace(Language.To_Trace(OBJ(self)) ).Id(),0}} 
  
/* {1} The go function for: c_type(self:Branch) [status=0] */
func F_c_type_Branch (self *Language.Branch ) *ClaireType  { 
    return  ToType(C_boolean.Id())
    } 
  
// The EID go function for: c_type @ Branch (throw: false) 
func E_c_type_Branch (self EID) EID { 
    return EID{F_c_type_Branch(Language.To_Branch(OBJ(self)) ).Id(),0}} 
  
/* {1} The go function for: c_code(self:Branch) [status=1] */
func F_c_code_Branch (self *Language.Branch ) EID { 
    var Result EID 
    { var arg_1 *ClaireObject  
      _ = arg_1
      { var arg_2 *Language.Do  
        _ = arg_2
        { var _CL_obj *Language.Do   = Language.To_Do(new(Language.Do).Is(Language.C_Do))
          { 
            var va_arg1 *Language.Do  
            var va_arg2 *ClaireList  
            va_arg1 = _CL_obj
            { 
              var v_bag_arg *ClaireAny  
              va_arg2= ToType(CEMPTY.Id()).EmptyList()
              { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                _CL_obj.Selector = C_choice
                /*property->property*/_CL_obj.Args = MakeConstantList(ClEnv.Id())
                /*list->list*/v_bag_arg = _CL_obj.Id()
                } 
              va_arg2.AddFast(v_bag_arg)
              { var arg_4 *Language.Do  
                _ = arg_4
                { var _CL_obj *Language.Do   = Language.To_Do(new(Language.Do).Is(Language.C_Do))
                  { 
                    var va_arg1 *Language.Do  
                    var va_arg2 *ClaireList  
                    va_arg1 = _CL_obj
                    { 
                      var v_bag_arg *ClaireAny  
                      va_arg2= ToType(CEMPTY.Id()).EmptyList()
                      { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                        _CL_obj.Selector = C_backtrack
                        /*property->property*/_CL_obj.Args = MakeConstantList(ClEnv.Id())
                        /*list->list*/v_bag_arg = _CL_obj.Id()
                        } 
                      va_arg2.AddFast(v_bag_arg)
                      va_arg2.AddFast(CFALSE.Id())} 
                    va_arg1.Args = va_arg2
                    /*list->list*/} 
                  arg_4 = _CL_obj
                  } 
                v_bag_arg = Language.C_If.Make(self.Args.At(1-1),CTRUE.Id(),arg_4.Id())
                } 
              va_arg2.AddFast(v_bag_arg)} 
            va_arg1.Args = va_arg2
            /*list->list*/} 
          arg_2 = _CL_obj
          } 
        { var arg_3 *Language.Do  
          _ = arg_3
          { var _CL_obj *Language.Do   = Language.To_Do(new(Language.Do).Is(Language.C_Do))
            { 
              var va_arg1 *Language.Do  
              var va_arg2 *ClaireList  
              va_arg1 = _CL_obj
              { 
                var v_bag_arg *ClaireAny  
                va_arg2= ToType(CEMPTY.Id()).EmptyList()
                { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                  _CL_obj.Selector = C_backtrack
                  /*property->property*/_CL_obj.Args = MakeConstantList(ClEnv.Id())
                  /*list->list*/v_bag_arg = _CL_obj.Id()
                  } 
                va_arg2.AddFast(v_bag_arg)
                va_arg2.AddFast(CFALSE.Id())} 
              va_arg1.Args = va_arg2
              /*list->list*/} 
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
  
/* {1} The go function for: c_code(self:Macro,s:class) [status=1] */
func F_c_code_Macro (self *Language.Macro ,s *ClaireClass ) EID { 
    var Result EID 
    { var arg_1 *ClaireAny  
      _ = arg_1
      var try_2 EID 
      /*g_try(v2:"try_2",loop:false) */
      try_2 = Core.F_CALL(Language.C_macroexpand,ARGS(EID{self.Id(),0}))
      /* ERROR PROTECTION INSERTED (arg_1-Result) */
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
  
/* {1} The go function for: c_type(self:Macro) [status=1] */
func F_c_type_Macro (self *Language.Macro ) EID { 
    var Result EID 
    { var arg_1 *ClaireAny  
      _ = arg_1
      var try_2 EID 
      /*g_try(v2:"try_2",loop:false) */
      try_2 = Core.F_CALL(Language.C_macroexpand,ARGS(EID{self.Id(),0}))
      /* ERROR PROTECTION INSERTED (arg_1-Result) */
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
  
/* {1} The go function for: c_type(self:Printf) [status=0] */
func F_c_type_Printf (self *Language.Printf ) *ClaireType  { 
    return  ToType(C_any.Id())
    } 
  
// The EID go function for: c_type @ Printf (throw: false) 
func E_c_type_Printf (self EID) EID { 
    return EID{F_c_type_Printf(Language.To_Printf(OBJ(self)) ).Id(),0}} 
  
/* {1} The go function for: c_code(self:Printf) [status=1] */
func F_c_code_Printf (self *Language.Printf ) EID { 
    var Result EID 
    { var l *ClaireList   = self.Args
      if (C_string.Id() != l.At(1-1).Isa.Id()) { 
        Result = Core.F_CALL(C_Compile_Cerror,ARGS(EID{MakeString("[209] the first argument in ~S must be a string").Id(),0},EID{self.Id(),0}))
        } else {
        { var s *ClaireString   = ToString(l.At(1-1))
          { var i int  = 1
            _ = i
            { var n int  = F_get_string(s,'~')
              { var r *ClaireList   = ToType(C_any.Id()).EmptyList()
                _ = r
                /*g_try(v2:"Result",loop:true) */
                Result= EID{CFALSE.Id(),0}
                for (n != 0) { 
                  /* While stat, v:"Result" loop:true */
                  var loop_1 EID 
                  _ = loop_1
                  /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
                  { var m rune  = s.At((n+1))
                    _ = m
                    /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
                    if (i < l.Length()) { 
                      i = (i+1)
                      loop_1 = EID{C__INT,IVAL(i)}
                      } else {
                      loop_1 = Core.F_CALL(C_Compile_Cerror,ARGS(EID{MakeString("[210] not enough arguments in ~S").Id(),0},EID{self.Id(),0}))
                      } 
                    /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
                    if ErrorIn(loop_1) {Result = loop_1
                    break
                    } else {
                    if (n > 1) { 
                      { var arg_2 *Language.Call  
                        _ = arg_2
                        { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                          _CL_obj.Selector = C_princ
                          /*property->property*/_CL_obj.Args = MakeConstantList((F_substring_string(s,1,(n-1))).Id())
                          /*list->list*/arg_2 = _CL_obj
                          } 
                        r = r.AddFast(arg_2.Id())/*t=any,s=list*/
                        } 
                      } 
                    var try_3 EID 
                    /*g_try(v2:"try_3",loop:tuple("Result", EID)) */
                    { var arg_4 *ClaireAny  
                      _ = arg_4
                      var try_5 EID 
                      /*g_try(v2:"try_5",loop:false) */
                      if ('A' == m) { 
                        { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                          _CL_obj.Selector = C_princ
                          /*property->property*/_CL_obj.Args = MakeConstantList(l.At(i-1))
                          /*list->list*/try_5 = EID{_CL_obj.Id(),0}
                          } 
                        }  else if ('S' == m) { 
                        { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                          _CL_obj.Selector = C_print
                          /*property->property*/_CL_obj.Args = MakeConstantList(l.At(i-1))
                          /*list->list*/try_5 = EID{_CL_obj.Id(),0}
                          } 
                        }  else if ('F' == m) { 
                        { var p_Z *ClaireBoolean   = CFALSE
                          { var j int 
                            var try_6 EID 
                            /*g_try(v2:"try_6",loop:false) */
                            { var arg_7 int 
                              _ = arg_7
                              var try_8 EID 
                              /*g_try(v2:"try_8",loop:false) */
                              { var arg_9 rune 
                                _ = arg_9
                                var try_10 EID 
                                /*g_try(v2:"try_10",loop:false) */
                                try_10 = Core.F_nth_get_string(s,(n+2),(n+2))
                                /* ERROR PROTECTION INSERTED (arg_9-try_8) */
                                if ErrorIn(try_10) {try_8 = try_10
                                } else {
                                arg_9 = CHAR(try_10)
                                try_8 = EID{C__INT,IVAL(int(arg_9))}
                                }
                                } 
                              /* ERROR PROTECTION INSERTED (arg_7-try_6) */
                              if ErrorIn(try_8) {try_6 = try_8
                              } else {
                              arg_7 = INT(try_8)
                              try_6 = EID{C__INT,IVAL((arg_7-48))}
                              }
                              } 
                            /* ERROR PROTECTION INSERTED (j-try_5) */
                            if ErrorIn(try_6) {try_5 = try_6
                            } else {
                            j = INT(try_6)
                            /*g_try(v2:"try_5",loop:false) */
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
                            /* ERROR PROTECTION INSERTED (try_5-try_5) */
                            if !ErrorIn(try_5) {
                            if ((p_Z != CTRUE) && 
                                ('%' == s.At((n+3)))) { 
                              p_Z = CTRUE
                              n = (n+1)
                              } 
                            n = (n+1)
                            { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                              _CL_obj.Selector = Core.C_mClaire_printFDigit
                              /*property->property*/{ 
                                var va_arg1 *Language.Call  
                                var va_arg2 *ClaireList  
                                va_arg1 = _CL_obj
                                { 
                                  var v_bag_arg *ClaireAny  
                                  va_arg2= ToType(CEMPTY.Id()).EmptyList()
                                  if (p_Z == CTRUE) { 
                                    { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                      _CL_obj.Selector = ToProperty(C__star.Id())
                                      /*property->property*/_CL_obj.Args = MakeConstantList(l.At(i-1),MakeFloat(100).Id())
                                      /*list->list*/v_bag_arg = _CL_obj.Id()
                                      } 
                                    } else {
                                    v_bag_arg = l.At(i-1)
                                    } 
                                  va_arg2.AddFast(v_bag_arg)
                                  va_arg2.AddFast(MakeInteger(j).Id())} 
                                va_arg1.Args = va_arg2
                                /*list->list*/} 
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
                      /* ERROR PROTECTION INSERTED (arg_4-try_3) */
                      if ErrorIn(try_5) {try_3 = try_5
                      } else {
                      arg_4 = ANY(try_5)
                      try_3 = EID{r.AddFast(arg_4).Id(),0}/*t=any,s=EID*/
                      }
                      } 
                    /* ERROR PROTECTION INSERTED (r-loop_1) */
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
                  /* ERROR PROTECTION INSERTED (loop_1-Result) */
                  if ErrorIn(loop_1) {Result = loop_1
                  break
                  } else {
                  /* try?:false, v2:"v_while8" loop will be:tuple("Result", EID) */
                  } 
                }
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                if (F_boolean_I_any((s).Id()) == CTRUE) { 
                  { var arg_11 *Language.Call  
                    _ = arg_11
                    { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                      _CL_obj.Selector = C_princ
                      /*property->property*/_CL_obj.Args = MakeConstantList((s).Id())
                      /*list->list*/arg_11 = _CL_obj
                      } 
                    r = r.AddFast(arg_11.Id())/*t=any,s=list*/
                    } 
                  } 
                { var arg_12 *Language.Do  
                  _ = arg_12
                  { var _CL_obj *Language.Do   = Language.To_Do(new(Language.Do).Is(Language.C_Do))
                    _CL_obj.Args = r
                    /*list->list*/arg_12 = _CL_obj
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
  
/* {1} The go function for: c_type(self:Error) [status=0] */
func F_c_type_Error (self *Language.Error ) *ClaireType  { 
    return  ToType(CEMPTY.Id())
    } 
  
// The EID go function for: c_type @ Error (throw: false) 
func E_c_type_Error (self EID) EID { 
    return EID{F_c_type_Error(Language.To_Error(OBJ(self)) ).Id(),0}} 
  
/* {1} The go function for: c_code(self:Error) [status=1] */
func F_c_code_Error (self *Language.Error ) EID { 
    var Result EID 
    { var arg_1 *Language.Call  
      _ = arg_1
      var try_2 EID 
      /*g_try(v2:"try_2",loop:false) */
      { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
        _CL_obj.Selector = C_close
        /*property->property*//*g_try(v2:"try_2",loop:false) */
        { 
          var va_arg1 *Language.Call  
          var va_arg2 *ClaireList  
          va_arg1 = _CL_obj
          var try_3 EID 
          /*g_try(v2:"try_3",loop:false) */
          { 
            var v_bag_arg *ClaireAny  
            try_3= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
            var try_4 EID 
            /*g_try(v2:"try_4",loop:false) */
            { var _CL_obj *Language.Cast   = Language.To_Cast(new(Language.Cast).Is(Language.C_Cast))
              /*g_try(v2:"try_4",loop:false) */
              { 
                var va_arg1 *Language.Cast  
                var va_arg2 *ClaireAny  
                va_arg1 = _CL_obj
                var try_5 EID 
                /*g_try(v2:"try_5",loop:false) */
                { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                  _CL_obj.Selector = C_Compile_anyObject_I
                  /*property->property*//*g_try(v2:"try_5",loop:false) */
                  { 
                    var va_arg1 *Language.Call  
                    var va_arg2 *ClaireList  
                    va_arg1 = _CL_obj
                    var try_6 EID 
                    /*g_try(v2:"try_6",loop:false) */
                    { 
                      var v_bag_arg *ClaireAny  
                      try_6= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                      ToList(OBJ(try_6)).AddFast(Core.C_general_error.Id())
                      var try_7 EID 
                      /*g_try(v2:"try_7",loop:false) */
                      { var arg_8 *ClaireAny  
                        _ = arg_8
                        var try_9 EID 
                        /*g_try(v2:"try_9",loop:false) */
                        try_9 = Core.F_car_list(self.Args)
                        /* ERROR PROTECTION INSERTED (arg_8-try_7) */
                        if ErrorIn(try_9) {try_7 = try_9
                        } else {
                        arg_8 = ANY(try_9)
                        try_7 = Core.F_CALL(C_c_code,ARGS(arg_8.ToEID(),EID{C_any.Id(),0}))
                        }
                        } 
                      /* ERROR PROTECTION INSERTED (v_bag_arg-try_6) */
                      if ErrorIn(try_7) {try_6 = try_7
                      } else {
                      v_bag_arg = ANY(try_7)
                      ToList(OBJ(try_6)).AddFast(v_bag_arg)
                      var try_10 EID 
                      /*g_try(v2:"try_10",loop:false) */
                      { var arg_11 *ClaireObject  
                        _ = arg_11
                        var try_12 EID 
                        /*g_try(v2:"try_12",loop:false) */
                        if (self.Args.Length() != 1) { 
                          { var _CL_obj *Language.List   = Language.To_List(new(Language.List).Is(Language.C_List))
                            /*g_try(v2:"try_12",loop:false) */
                            { 
                              var va_arg1 *Language.Construct  
                              var va_arg2 *ClaireList  
                              va_arg1 = Language.To_Construct(_CL_obj.Id())
                              var try_13 EID 
                              /*g_try(v2:"try_13",loop:false) */
                              try_13 = self.Args.Cdr()
                              /* ERROR PROTECTION INSERTED (va_arg2-try_12) */
                              if ErrorIn(try_13) {try_12 = try_13
                              } else {
                              va_arg2 = ToList(OBJ(try_13))
                              va_arg1.Args = va_arg2
                              /*list->list*/try_12 = EID{va_arg2.Id(),0}
                              }
                              } 
                            /* ERROR PROTECTION INSERTED (try_12-try_12) */
                            if !ErrorIn(try_12) {
                            try_12 = EID{_CL_obj.Id(),0}
                            }
                            } 
                          } else {
                          try_12 = EID{CNIL.Id(),0}
                          } 
                        /* ERROR PROTECTION INSERTED (arg_11-try_10) */
                        if ErrorIn(try_12) {try_10 = try_12
                        } else {
                        arg_11 = ToObject(OBJ(try_12))
                        try_10 = Core.F_CALL(C_c_code,ARGS(EID{arg_11.Id(),0},EID{C_any.Id(),0}))
                        }
                        } 
                      /* ERROR PROTECTION INSERTED (v_bag_arg-try_6) */
                      if ErrorIn(try_10) {try_6 = try_10
                      } else {
                      v_bag_arg = ANY(try_10)
                      ToList(OBJ(try_6)).AddFast(v_bag_arg)}}
                      } 
                    /* ERROR PROTECTION INSERTED (va_arg2-try_5) */
                    if ErrorIn(try_6) {try_5 = try_6
                    } else {
                    va_arg2 = ToList(OBJ(try_6))
                    va_arg1.Args = va_arg2
                    /*list->list*/try_5 = EID{va_arg2.Id(),0}
                    }
                    } 
                  /* ERROR PROTECTION INSERTED (try_5-try_5) */
                  if !ErrorIn(try_5) {
                  try_5 = EID{_CL_obj.Id(),0}
                  }
                  } 
                /* ERROR PROTECTION INSERTED (va_arg2-try_4) */
                if ErrorIn(try_5) {try_4 = try_5
                } else {
                va_arg2 = ANY(try_5)
                va_arg1.Arg = va_arg2
                /*any->any*/try_4 = va_arg2.ToEID()
                }
                } 
              /* ERROR PROTECTION INSERTED (try_4-try_4) */
              if !ErrorIn(try_4) {
              _CL_obj.SetArg = ToType(C_exception.Id())
              /*type->type*/try_4 = EID{_CL_obj.Id(),0}
              }
              } 
            /* ERROR PROTECTION INSERTED (v_bag_arg-try_3) */
            if ErrorIn(try_4) {try_3 = try_4
            } else {
            v_bag_arg = ANY(try_4)
            ToList(OBJ(try_3)).AddFast(v_bag_arg)}
            } 
          /* ERROR PROTECTION INSERTED (va_arg2-try_2) */
          if ErrorIn(try_3) {try_2 = try_3
          } else {
          va_arg2 = ToList(OBJ(try_3))
          va_arg1.Args = va_arg2
          /*list->list*/try_2 = EID{va_arg2.Id(),0}
          }
          } 
        /* ERROR PROTECTION INSERTED (try_2-try_2) */
        if !ErrorIn(try_2) {
        try_2 = EID{_CL_obj.Id(),0}
        }
        } 
      /* ERROR PROTECTION INSERTED (arg_1-Result) */
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
/* {1} The go function for: extendedTest?(self:If) [status=0] */
func F_Optimize_extendedTest_ask_If (self *Language.If ) *ClaireType  { 
    // procedure body with s = type 
var Result *ClaireType  
    { var _Zt *ClaireAny   = self.Test
      if (_Zt.Isa.IsIn(Language.C_Call) == CTRUE) { 
        { var g0248 *Language.Call   = Language.To_Call(_Zt)
          if ((g0248.Args.At(1-1).Isa.IsIn(C_Variable) == CTRUE) && 
              (g0248.Selector.Id() == Core.C_known_ask.Id())) { 
            Result = ToType(OBJ(Core.F_CALL(C_range,ARGS(g0248.Args.At(1-1).ToEID()))))
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
/* {1} The go function for: c_type(self:If) [status=1] */
func F_c_type_If (self *Language.If ) EID { 
    var Result EID 
    { var _Zr *ClaireType   = F_Optimize_extendedTest_ask_If(self)
      /*g_try(v2:"Result",loop:true) */
      var g0250I *ClaireBoolean  
      var try_1 EID 
      /*g_try(v2:"try_1",loop:false) */
      try_1 = F_Optimize_extended_ask_type(_Zr)
      /* ERROR PROTECTION INSERTED (g0250I-Result) */
      if ErrorIn(try_1) {Result = try_1
      } else {
      g0250I = ToBoolean(OBJ(try_1))
      if (g0250I == CTRUE) { 
        F_Optimize_range_sets_any(ToList(OBJ(Core.F_CALL(C_args,ARGS(self.Test.ToEID())))).At(1-1),F_Optimize_sort_abstract_I_type(ToType(OBJ(Core.F_CALL(C_mClaire_t1,ARGS(EID{_Zr.Id(),0}))))))
        Result = EVOID
        } else {
        Result = EID{CFALSE.Id(),0}
        } 
      }
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      { var result *ClaireType  
        _ = result
        var try_2 EID 
        /*g_try(v2:"try_2",loop:false) */
        { var arg_3 *ClaireType  
          _ = arg_3
          var try_5 EID 
          /*g_try(v2:"try_5",loop:false) */
          try_5 = Core.F_CALL(C_c_type,ARGS(self.Arg.ToEID()))
          /* ERROR PROTECTION INSERTED (arg_3-try_2) */
          if ErrorIn(try_5) {try_2 = try_5
          } else {
          arg_3 = ToType(OBJ(try_5))
          { var arg_4 *ClaireType  
            _ = arg_4
            var try_6 EID 
            /*g_try(v2:"try_6",loop:false) */
            try_6 = Core.F_CALL(C_c_type,ARGS(self.Other.ToEID()))
            /* ERROR PROTECTION INSERTED (arg_4-try_2) */
            if ErrorIn(try_6) {try_2 = try_6
            } else {
            arg_4 = ToType(OBJ(try_6))
            try_2 = EID{Core.F_U_type(arg_3,arg_4).Id(),0}
            }
            } 
          }
          } 
        /* ERROR PROTECTION INSERTED (result-Result) */
        if ErrorIn(try_2) {Result = try_2
        } else {
        result = ToType(OBJ(try_2))
        /*g_try(v2:"Result",loop:true) */
        var g0251I *ClaireBoolean  
        var try_7 EID 
        /*g_try(v2:"try_7",loop:false) */
        try_7 = F_Optimize_extended_ask_type(_Zr)
        /* ERROR PROTECTION INSERTED (g0251I-Result) */
        if ErrorIn(try_7) {Result = try_7
        } else {
        g0251I = ToBoolean(OBJ(try_7))
        if (g0251I == CTRUE) { 
          Result = Core.F_put_property2(C_range,ToObject(ToList(OBJ(Core.F_CALL(C_args,ARGS(self.Test.ToEID())))).At(1-1)),_Zr.Id())
          } else {
          Result = EID{CFALSE.Id(),0}
          } 
        }
        /* ERROR PROTECTION INSERTED (Result-Result) */
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
/* {1} The go function for: c_code(self:If,s:class) [status=1] */
func F_c_code_If (self *Language.If ,s *ClaireClass ) EID { 
    var Result EID 
    { var _Zr *ClaireType   = F_Optimize_extendedTest_ask_If(self)
      /*g_try(v2:"Result",loop:true) */
      var g0252I *ClaireBoolean  
      var try_1 EID 
      /*g_try(v2:"try_1",loop:false) */
      try_1 = F_Optimize_extended_ask_type(_Zr)
      /* ERROR PROTECTION INSERTED (g0252I-Result) */
      if ErrorIn(try_1) {Result = try_1
      } else {
      g0252I = ToBoolean(OBJ(try_1))
      if (g0252I == CTRUE) { 
        F_Optimize_range_sets_any(ToList(OBJ(Core.F_CALL(C_args,ARGS(self.Test.ToEID())))).At(1-1),F_Optimize_sort_abstract_I_type(ToType(OBJ(Core.F_CALL(C_mClaire_t1,ARGS(EID{_Zr.Id(),0}))))))
        Result = EVOID
        } else {
        Result = EID{CFALSE.Id(),0}
        } 
      }
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      /*g_try(v2:"Result",loop:true) */
      var g0253I *ClaireBoolean  
      var try_2 EID 
      /*g_try(v2:"try_2",loop:false) */
      { 
        var v_and3 *ClaireBoolean  
        
        var try_3 EID 
        /*g_try(v2:"try_3",loop:false) */
        { var arg_4 *ClaireBoolean  
          _ = arg_4
          var try_5 EID 
          /*g_try(v2:"try_5",loop:false) */
          { var arg_6 *ClaireType  
            _ = arg_6
            var try_7 EID 
            /*g_try(v2:"try_7",loop:false) */
            { var arg_8 *ClaireType  
              _ = arg_8
              var try_9 EID 
              /*g_try(v2:"try_9",loop:false) */
              try_9 = Core.F_CALL(C_c_type,ARGS(self.Test.ToEID()))
              /* ERROR PROTECTION INSERTED (arg_8-try_7) */
              if ErrorIn(try_9) {try_7 = try_9
              } else {
              arg_8 = ToType(OBJ(try_9))
              try_7 = EID{F_Optimize_ptype_type(arg_8).Id(),0}
              }
              } 
            /* ERROR PROTECTION INSERTED (arg_6-try_5) */
            if ErrorIn(try_7) {try_5 = try_7
            } else {
            arg_6 = ToType(OBJ(try_7))
            try_5 = EID{arg_6.Included(ToType(C_boolean.Id())).Id(),0}
            }
            } 
          /* ERROR PROTECTION INSERTED (arg_4-try_3) */
          if ErrorIn(try_5) {try_3 = try_5
          } else {
          arg_4 = ToBoolean(OBJ(try_5))
          try_3 = EID{arg_4.Not.Id(),0}
          }
          } 
        /* ERROR PROTECTION INSERTED (v_and3-try_2) */
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
      /* ERROR PROTECTION INSERTED (g0253I-Result) */
      if ErrorIn(try_2) {Result = try_2
      } else {
      g0253I = ToBoolean(OBJ(try_2))
      if (g0253I == CTRUE) { 
        F_Compile_warn_void()
        Result = Core.F_tformat_string(MakeString("CLAIRE 3.3 SYNTAX - Test in ~S should be a boolean [260]\n"),1,MakeConstantList(self.Id()))
        } else {
        Result = EID{CFALSE.Id(),0}
        } 
      }
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      { var result *Language.If  
        _ = result
        var try_10 EID 
        /*g_try(v2:"try_10",loop:false) */
        { var arg_11 *ClaireAny  
          _ = arg_11
          var try_14 EID 
          /*g_try(v2:"try_14",loop:false) */
          try_14 = F_Optimize_c_boolean_any(self.Test)
          /* ERROR PROTECTION INSERTED (arg_11-try_10) */
          if ErrorIn(try_14) {try_10 = try_14
          } else {
          arg_11 = ANY(try_14)
          { var arg_12 *ClaireAny  
            _ = arg_12
            var try_15 EID 
            /*g_try(v2:"try_15",loop:false) */
            try_15 = Core.F_CALL(C_c_code,ARGS(self.Arg.ToEID(),EID{s.Id(),0}))
            /* ERROR PROTECTION INSERTED (arg_12-try_10) */
            if ErrorIn(try_15) {try_10 = try_15
            } else {
            arg_12 = ANY(try_15)
            { var arg_13 *ClaireAny  
              _ = arg_13
              var try_16 EID 
              /*g_try(v2:"try_16",loop:false) */
              try_16 = Core.F_CALL(C_c_code,ARGS(self.Other.ToEID(),EID{s.Id(),0}))
              /* ERROR PROTECTION INSERTED (arg_13-try_10) */
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
        /* ERROR PROTECTION INSERTED (result-Result) */
        if ErrorIn(try_10) {Result = try_10
        } else {
        result = Language.To_If(OBJ(try_10))
        /*g_try(v2:"Result",loop:true) */
        var g0254I *ClaireBoolean  
        var try_17 EID 
        /*g_try(v2:"try_17",loop:false) */
        try_17 = F_Optimize_extended_ask_type(_Zr)
        /* ERROR PROTECTION INSERTED (g0254I-Result) */
        if ErrorIn(try_17) {Result = try_17
        } else {
        g0254I = ToBoolean(OBJ(try_17))
        if (g0254I == CTRUE) { 
          Result = Core.F_put_property2(C_range,ToObject(ToList(OBJ(Core.F_CALL(C_args,ARGS(self.Test.ToEID())))).At(1-1)),_Zr.Id())
          } else {
          Result = EID{CFALSE.Id(),0}
          } 
        }
        /* ERROR PROTECTION INSERTED (Result-Result) */
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
/* {1} The go function for: c_type(self:Case) [status=1] */
func F_c_type_Case (self *Language.Case ) EID { 
    var Result EID 
    { var _Zvar *ClaireAny   = self.ClaireVar
      { var _Ztype *ClaireAny  
        if (_Zvar.Isa.IsIn(C_Variable) == CTRUE) { 
          { var g0255 *ClaireVariable   = To_Variable(_Zvar)
            _ = g0255
            _Ztype = Core.F_get_property(C_range,ToObject(g0255.Id()))
            } 
          } else {
          _Ztype = C_any.Id()
          } 
        { var l *ClaireList   = self.Args.Copy()
          { var rtype *ClaireType   = ToType(CEMPTY.Id())
            { var utype *ClaireType   = ToType(CEMPTY.Id())
              _ = utype
              /*g_try(v2:"Result",loop:true) */
              Result= EID{CFALSE.Id(),0}
              for (l.Length() > 0) { 
                /* While stat, v:"Result" loop:true */
                var loop_1 EID 
                _ = loop_1
                { 
                /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
                if (l.At(1-1).Isa.IsIn(C_type) == CTRUE) { 
                  utype = Core.F_U_type(utype,ToType(l.At(1-1)))
                  if (F_Compile_osort_any(_Ztype).Id() == F_Compile_osort_any(l.At(1-1)).Id()) { 
                    F_Optimize_range_sets_any(_Zvar,ToType(l.At(1-1)))
                    loop_1 = EVOID
                    }  else if (F_Compile_osort_any(_Ztype).Id() == C_any.Id()) { 
                    F_Optimize_range_sets_any(_Zvar,F_Optimize_sort_abstract_I_type(ToType(l.At(1-1))))
                    loop_1 = EVOID
                    } else {
                    loop_1 = EID{CFALSE.Id(),0}
                    } 
                  } else {
                  { var arg_2 *ClaireAny  
                    _ = arg_2
                    var try_3 EID 
                    /*g_try(v2:"try_3",loop:false) */
                    try_3 = Core.F_car_list(l)
                    /* ERROR PROTECTION INSERTED (arg_2-loop_1) */
                    if ErrorIn(try_3) {loop_1 = try_3
                    } else {
                    arg_2 = ANY(try_3)
                    loop_1 = Core.F_CALL(C_Compile_Cerror,ARGS(EID{MakeString("[208] wrong type declaration for case: ~S in ~S").Id(),0},arg_2.ToEID(),EID{self.Id(),0}))
                    }
                    } 
                  } 
                /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
                if ErrorIn(loop_1) {Result = loop_1
                break
                } else {
                var try_4 EID 
                /*g_try(v2:"try_4",loop:tuple("Result", EID)) */
                { var arg_5 *ClaireType  
                  _ = arg_5
                  var try_6 EID 
                  /*g_try(v2:"try_6",loop:false) */
                  try_6 = Core.F_CALL(C_c_type,ARGS(l.At(2-1).ToEID()))
                  /* ERROR PROTECTION INSERTED (arg_5-try_4) */
                  if ErrorIn(try_6) {try_4 = try_6
                  } else {
                  arg_5 = ToType(OBJ(try_6))
                  try_4 = EID{Core.F_U_type(rtype,arg_5).Id(),0}
                  }
                  } 
                /* ERROR PROTECTION INSERTED (rtype-loop_1) */
                if ErrorIn(try_4) {loop_1 = try_4
                Result = try_4
                break
                } else {
                rtype = ToType(OBJ(try_4))
                loop_1 = EID{rtype.Id(),0}
                
                if (_Zvar.Isa.IsIn(C_Variable) == CTRUE) { 
                  { var g0257 *ClaireVariable   = To_Variable(_Zvar)
                    _ = g0257
                    g0257.Range = ToType(_Ztype)
                    /*type->type*/} 
                  } 
                l = l.Skip(2)
                }}
                /* try?:false, v2:"v_while7" loop will be:tuple("Result", EID) */
                } 
              }
              /* ERROR PROTECTION INSERTED (Result-Result) */
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
/* {1} The go function for: case_branch(x:any,%var:any,%type:type) [status=0] */
func F_Optimize_case_branch_any (x *ClaireAny ,_Zvar *ClaireAny ,_Ztype *ClaireType ) *ClaireAny  { 
    // procedure body with s = any 
var Result *ClaireAny  
    if (_Zvar.Isa.IsIn(C_Variable) == CTRUE) { 
      { var g0258 *ClaireVariable   = To_Variable(_Zvar)
        { var vsub *ClaireVariable   = F_Compile_Variable_I_symbol(Core.F_gensym_void(),-1,_Ztype.Id())
          if ((Equal(_Ztype.Id(),g0258.Range.Id()) != CTRUE) && 
              ((_Ztype.Id() != C_any.Id()) && 
                (Language.F_occurrence_any(x,g0258) > 0))) { 
            { var _CL_obj *Language.Let   = Language.To_Let(new(Language.Let).Is(Language.C_Let))
              _CL_obj.ClaireVar = vsub
              /*Variable->Variable*/_CL_obj.Value = g0258.Id()
              /*any->any*/_CL_obj.Arg = F_Optimize_case_substitution_any(x,g0258,vsub)
              /*any->any*/Result = _CL_obj.Id()
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
/* {1} The go function for: case_substitution(x:any,%var:Variable,vsub:Variable) [status=0] */
func F_Optimize_case_substitution_any (x *ClaireAny ,_Zvar *ClaireVariable ,vsub *ClaireVariable ) *ClaireAny  { 
    // procedure body with s = any 
var Result *ClaireAny  
    { var y *ClaireAny   = Language.F_substitution_any(Language.F_instruction_copy_any(x),_Zvar,vsub.Id())
      if (Language.F_occurchange_any(y,vsub) == CTRUE) { 
        { var _CL_obj *Language.Do   = Language.To_Do(new(Language.Do).Is(Language.C_Do))
          _CL_obj.Args = MakeConstantList(y,Language.C_Assign.Make(_Zvar.Id(),vsub.Id()))
          /*list->list*/Result = _CL_obj.Id()
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
/* {1} The go function for: c_code(self:Case,s:class) [status=1] */
func F_c_code_Case (self *Language.Case ,s *ClaireClass ) EID { 
    var Result EID 
    { var _Zvar *ClaireAny   = self.ClaireVar
      { var _Ztype *ClaireAny  
        _ = _Ztype
        if (_Zvar.Isa.IsIn(C_Variable) == CTRUE) { 
          { var g0260 *ClaireVariable   = To_Variable(_Zvar)
            _ = g0260
            _Ztype = Core.F_get_property(C_range,ToObject(g0260.Id()))
            } 
          } else {
          _Ztype = C_any.Id()
          } 
        { var l *ClaireList   = self.Args.Copy()
          { var utype *ClaireAny   = CEMPTY.Id()
            _ = utype
            { var ctest1 *ClaireAny  
              _ = ctest1
              var try_1 EID 
              /*g_try(v2:"try_1",loop:false) */
              { var arg_2 *Language.Call  
                _ = arg_2
                { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                  _CL_obj.Selector = ToProperty(C__Z.Id())
                  /*property->property*/_CL_obj.Args = MakeConstantList(_Zvar,l.At(1-1))
                  /*list->list*/arg_2 = _CL_obj
                  } 
                try_1 = F_Optimize_c_boolean_any(arg_2.Id())
                } 
              /* ERROR PROTECTION INSERTED (ctest1-Result) */
              if ErrorIn(try_1) {Result = try_1
              } else {
              ctest1 = ANY(try_1)
              { var rep *Language.If  
                var try_3 EID 
                /*g_try(v2:"try_3",loop:false) */
                { var arg_4 *ClaireAny  
                  _ = arg_4
                  var try_6 EID 
                  /*g_try(v2:"try_6",loop:false) */
                  try_6 = Core.F_CALL(C_c_code,ARGS(F_Optimize_case_branch_any(l.At(2-1),_Zvar,ToType(l.At(1-1))).ToEID(),EID{s.Id(),0}))
                  /* ERROR PROTECTION INSERTED (arg_4-try_3) */
                  if ErrorIn(try_6) {try_3 = try_6
                  } else {
                  arg_4 = ANY(try_6)
                  { var arg_5 *ClaireAny  
                    _ = arg_5
                    var try_7 EID 
                    /*g_try(v2:"try_7",loop:false) */
                    try_7 = Core.F_CALL(C_c_code,ARGS(EID{CFALSE.Id(),0},EID{s.Id(),0}))
                    /* ERROR PROTECTION INSERTED (arg_5-try_3) */
                    if ErrorIn(try_7) {try_3 = try_7
                    } else {
                    arg_5 = ANY(try_7)
                    try_3 = Language.C_If.Make(ctest1,arg_4,arg_5).ToEID()
                    }
                    } 
                  }
                  } 
                /* ERROR PROTECTION INSERTED (rep-Result) */
                if ErrorIn(try_3) {Result = try_3
                } else {
                rep = Language.To_If(OBJ(try_3))
                { var pointer *Language.If   = rep
                  l = l.Skip(2)
                  /*g_try(v2:"Result",loop:true) */
                  Result= EID{CFALSE.Id(),0}
                  for (l.Length() > 0) { 
                    /* While stat, v:"Result" loop:true */
                    var loop_8 EID 
                    _ = loop_8
                    { 
                    utype = Core.F_U_type(ToType(utype),ToType(l.At(1-1))).Id()
                    /*g_try(v2:"loop_8",loop:tuple("Result", EID)) */
                    if (ToType(_Ztype).Included(ToType(utype)) == CTRUE) { 
                      /*g_try(v2:"loop_8",loop:tuple("Result", EID)) */
                      { 
                        var va_arg1 *Language.If  
                        var va_arg2 *ClaireAny  
                        va_arg1 = pointer
                        var try_9 EID 
                        /*g_try(v2:"try_9",loop:false) */
                        try_9 = Core.F_CALL(C_c_code,ARGS(F_Optimize_case_branch_any(l.At(2-1),_Zvar,ToType(l.At(1-1))).ToEID(),EID{s.Id(),0}))
                        /* ERROR PROTECTION INSERTED (va_arg2-loop_8) */
                        if ErrorIn(try_9) {loop_8 = try_9
                        } else {
                        va_arg2 = ANY(try_9)
                        va_arg1.Other = va_arg2
                        /*any->any*/loop_8 = va_arg2.ToEID()
                        }
                        } 
                      /* ERROR PROTECTION INSERTED (loop_8-loop_8) */
                      if ErrorIn(loop_8) {Result = loop_8
                      break
                      } else {
                      Result = EID{CTRUE.Id(),0}
                      break
                      }
                      } else {
                      { var ctest *ClaireAny  
                        _ = ctest
                        var try_10 EID 
                        /*g_try(v2:"try_10",loop:false) */
                        { var arg_11 *Language.Call  
                          _ = arg_11
                          { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                            _CL_obj.Selector = ToProperty(C__Z.Id())
                            /*property->property*/_CL_obj.Args = MakeConstantList(_Zvar,l.At(1-1))
                            /*list->list*/arg_11 = _CL_obj
                            } 
                          try_10 = F_Optimize_c_boolean_any(arg_11.Id())
                          } 
                        /* ERROR PROTECTION INSERTED (ctest-loop_8) */
                        if ErrorIn(try_10) {loop_8 = try_10
                        } else {
                        ctest = ANY(try_10)
                        /*g_try(v2:"loop_8",loop:tuple("Result", EID)) */
                        { 
                          var va_arg1 *Language.If  
                          var va_arg2 *ClaireAny  
                          va_arg1 = pointer
                          var try_12 EID 
                          /*g_try(v2:"try_12",loop:false) */
                          { var arg_13 *ClaireAny  
                            _ = arg_13
                            var try_15 EID 
                            /*g_try(v2:"try_15",loop:false) */
                            try_15 = Core.F_CALL(C_c_code,ARGS(F_Optimize_case_branch_any(l.At(2-1),_Zvar,ToType(l.At(1-1))).ToEID(),EID{s.Id(),0}))
                            /* ERROR PROTECTION INSERTED (arg_13-try_12) */
                            if ErrorIn(try_15) {try_12 = try_15
                            } else {
                            arg_13 = ANY(try_15)
                            { var arg_14 *ClaireAny  
                              _ = arg_14
                              var try_16 EID 
                              /*g_try(v2:"try_16",loop:false) */
                              try_16 = Core.F_CALL(C_c_code,ARGS(EID{CFALSE.Id(),0},EID{s.Id(),0}))
                              /* ERROR PROTECTION INSERTED (arg_14-try_12) */
                              if ErrorIn(try_16) {try_12 = try_16
                              } else {
                              arg_14 = ANY(try_16)
                              try_12 = Language.C_If.Make(ctest,arg_13,arg_14).ToEID()
                              }
                              } 
                            }
                            } 
                          /* ERROR PROTECTION INSERTED (va_arg2-loop_8) */
                          if ErrorIn(try_12) {loop_8 = try_12
                          } else {
                          va_arg2 = ANY(try_12)
                          va_arg1.Other = va_arg2
                          /*any->any*/loop_8 = va_arg2.ToEID()
                          }
                          } 
                        /* ERROR PROTECTION INSERTED (loop_8-loop_8) */
                        if ErrorIn(loop_8) {Result = loop_8
                        break
                        } else {
                        pointer = Language.To_If(pointer.Other)
                        loop_8 = EID{pointer.Id(),0}
                        }
                        }
                        } 
                      } 
                    /* ERROR PROTECTION INSERTED (loop_8-loop_8) */
                    if ErrorIn(loop_8) {Result = loop_8
                    break
                    } else {
                    l = l.Skip(2)
                    }
                    /* try?:false, v2:"v_while9" loop will be:tuple("Result", EID) */
                    } 
                  }
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  var g0263I *ClaireBoolean  
                  if (_Zvar.Isa.IsIn(Language.C_Definition) == CTRUE) { 
                    { var g0262 *Language.Definition   = Language.To_Definition(_Zvar)
                      _ = g0262
                      g0263I = g0262.Arg.Isa.IsIn(C_exception)
                      } 
                    } else {
                    g0263I = CFALSE
                    } 
                  if (g0263I == CTRUE) { 
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
/* {1} The go function for: c_type(self:Do) [status=1] */
func F_c_type_Do (self *Language.Do ) EID { 
    var Result EID 
    { var arg_1 *ClaireAny  
      _ = arg_1
      var try_2 EID 
      /*g_try(v2:"try_2",loop:false) */
      try_2 = Core.F_last_list(self.Args)
      /* ERROR PROTECTION INSERTED (arg_1-Result) */
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
  
/* {1} The go function for: c_code(self:Do,s:class) [status=1] */
func F_c_code_Do (self *Language.Do ,s *ClaireClass ) EID { 
    var Result EID 
    { var _CL_obj *Language.Do   = Language.To_Do(new(Language.Do).Is(Language.C_Do))
      /*g_try(v2:"Result",loop:true) */
      { 
        var va_arg1 *Language.Do  
        var va_arg2 *ClaireList  
        va_arg1 = _CL_obj
        var try_1 EID 
        /*g_try(v2:"try_1",loop:false) */
        { var m int  = self.Args.Length()
          _ = m
          { var n int  = 0
            _ = n
            { 
              var v_list6 *ClaireList  
              var x *ClaireAny  
              var v_local6 *ClaireAny  
              v_list6 = self.Args
              try_1 = EID{CreateList(ToType(CEMPTY.Id()),v_list6.Length()).Id(),0}
              for CLcount := 0; CLcount < v_list6.Length(); CLcount++{ 
                x = v_list6.At(CLcount)
                var try_2 EID 
                /*g_try(v2:"try_2",loop:tuple("try_1", EID)) */
                n = (n+1)
                /*g_try(v2:"try_2",loop:tuple("try_1", EID)) */
                { var arg_3 *ClaireClass  
                  _ = arg_3
                  if (n == m) { 
                    arg_3 = s
                    } else {
                    arg_3 = C_void
                    } 
                  try_2 = Core.F_CALL(C_c_code,ARGS(x.ToEID(),EID{arg_3.Id(),0}))
                  } 
                /* ERROR PROTECTION INSERTED (try_2-try_2) */
                if ErrorIn(try_2) {try_1 = try_2
                break
                } else {
                }
                /* ERROR PROTECTION INSERTED (v_local6-try_1) */
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
        /* ERROR PROTECTION INSERTED (va_arg2-Result) */
        if ErrorIn(try_1) {Result = try_1
        } else {
        va_arg2 = ToList(OBJ(try_1))
        va_arg1.Args = va_arg2
        /*list->list*/Result = EID{va_arg2.Id(),0}
        }
        } 
      /* ERROR PROTECTION INSERTED (Result-Result) */
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
/* {1} The go function for: c_type(self:Let) [status=1] */
func F_c_type_Let (self *Language.Let ) EID { 
    var Result EID 
    /*g_try(v2:"Result",loop:true) */
    { var arg_1 *ClaireType  
      _ = arg_1
      var try_2 EID 
      /*g_try(v2:"try_2",loop:false) */
      try_2 = Core.F_CALL(C_c_type,ARGS(self.Value.ToEID()))
      /* ERROR PROTECTION INSERTED (arg_1-Result) */
      if ErrorIn(try_2) {Result = try_2
      } else {
      arg_1 = ToType(OBJ(try_2))
      Result = F_Optimize_range_infers_Variable(self.ClaireVar,arg_1)
      }
      } 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    Result = Core.F_CALL(C_c_type,ARGS(self.Arg.ToEID()))
    }
    return Result} 
  
// The EID go function for: c_type @ Let (throw: true) 
func E_c_type_Let (self EID) EID { 
    return F_c_type_Let(Language.To_Let(OBJ(self)) )} 
  
// works also for Let+ / Let*
/* {1} The go function for: c_code(self:Let,s:class) [status=1] */
func F_c_code_Let (self *Language.Let ,s *ClaireClass ) EID { 
    var Result EID 
    { var _Zv *ClaireAny   = self.Value
      { var _Ztype *ClaireType  
        var try_1 EID 
        /*g_try(v2:"try_1",loop:false) */
        { var arg_2 *ClaireType  
          _ = arg_2
          var try_3 EID 
          /*g_try(v2:"try_3",loop:false) */
          try_3 = Core.F_CALL(C_c_type,ARGS(_Zv.ToEID()))
          /* ERROR PROTECTION INSERTED (arg_2-try_1) */
          if ErrorIn(try_3) {try_1 = try_3
          } else {
          arg_2 = ToType(OBJ(try_3))
          try_1 = EID{F_Optimize_ptype_type(arg_2).Id(),0}
          }
          } 
        /* ERROR PROTECTION INSERTED (_Ztype-Result) */
        if ErrorIn(try_1) {Result = try_1
        } else {
        _Ztype = ToType(OBJ(try_1))
        /*g_try(v2:"Result",loop:true) */
        Result = F_Optimize_range_infers_Variable(self.ClaireVar,_Ztype)
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        /*g_try(v2:"Result",loop:true) */
        if (_Ztype.Included(self.ClaireVar.Range) != CTRUE) { 
          var try_4 EID 
          /*g_try(v2:"try_4",loop:false) */
          try_4 = F_Optimize_c_warn_Variable(self.ClaireVar,_Zv,_Ztype)
          /* ERROR PROTECTION INSERTED (_Zv-Result) */
          if ErrorIn(try_4) {Result = try_4
          } else {
          _Zv = ANY(try_4)
          Result = _Zv.ToEID()
          }
          } else {
          Result = EID{CFALSE.Id(),0}
          } 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        { var x *Language.Let  
          var try_5 EID 
          /*g_try(v2:"try_5",loop:false) */
          { var _CL_obj *Language.Let   = Language.To_Let(new(Language.Let).Is(Language.C_Let))
            _CL_obj.ClaireVar = self.ClaireVar
            /*Variable->Variable*//*g_try(v2:"try_5",loop:false) */
            { 
              var va_arg1 *Language.Let  
              var va_arg2 *ClaireAny  
              va_arg1 = _CL_obj
              var try_6 EID 
              /*g_try(v2:"try_6",loop:false) */
              try_6 = F_Compile_c_strict_code_any(_Zv,F_Compile_psort_any(self.ClaireVar.Range.Id()))
              /* ERROR PROTECTION INSERTED (va_arg2-try_5) */
              if ErrorIn(try_6) {try_5 = try_6
              } else {
              va_arg2 = ANY(try_6)
              va_arg1.Value = va_arg2
              /*any->any*/try_5 = va_arg2.ToEID()
              }
              } 
            /* ERROR PROTECTION INSERTED (try_5-try_5) */
            if !ErrorIn(try_5) {
            /*g_try(v2:"try_5",loop:false) */
            { 
              var va_arg1 *Language.Let  
              var va_arg2 *ClaireAny  
              va_arg1 = _CL_obj
              var try_7 EID 
              /*g_try(v2:"try_7",loop:false) */
              try_7 = Core.F_CALL(C_c_code,ARGS(self.Arg.ToEID(),EID{s.Id(),0}))
              /* ERROR PROTECTION INSERTED (va_arg2-try_5) */
              if ErrorIn(try_7) {try_5 = try_7
              } else {
              va_arg2 = ANY(try_7)
              va_arg1.Arg = va_arg2
              /*any->any*/try_5 = va_arg2.ToEID()
              }
              } 
            /* ERROR PROTECTION INSERTED (try_5-try_5) */
            if !ErrorIn(try_5) {
            try_5 = EID{_CL_obj.Id(),0}
            }}
            } 
          /* ERROR PROTECTION INSERTED (x-Result) */
          if ErrorIn(try_5) {Result = try_5
          } else {
          x = Language.To_Let(OBJ(try_5))
          x.Isa = self.Isa
          /*class->class*/Result = EID{x.Id(),0}
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
/* {1} The go function for: c_type(self:When) [status=1] */
func F_c_type_When (self *Language.When ) EID { 
    var Result EID 
    { var _Zv *ClaireAny   = self.Value
      { var v *ClaireVariable   = self.ClaireVar
        _ = v
        { var d *ClaireAny  
          var try_1 EID 
          /*g_try(v2:"try_1",loop:false) */
          try_1 = F_Optimize_daccess_any(_Zv,CTRUE)
          /* ERROR PROTECTION INSERTED (d-Result) */
          if ErrorIn(try_1) {Result = try_1
          } else {
          d = ANY(try_1)
          { var _Ztype *ClaireType  
            var try_2 EID 
            /*g_try(v2:"try_2",loop:false) */
            if (d != CNULL) { 
              try_2 = Core.F_CALL(C_c_type,ARGS(d.ToEID()))
              } else {
              try_2 = Core.F_CALL(C_c_type,ARGS(_Zv.ToEID()))
              } 
            /* ERROR PROTECTION INSERTED (_Ztype-Result) */
            if ErrorIn(try_2) {Result = try_2
            } else {
            _Ztype = ToType(OBJ(try_2))
            /*g_try(v2:"Result",loop:true) */
            var g0264I *ClaireBoolean  
            var try_3 EID 
            /*g_try(v2:"try_3",loop:false) */
            try_3 = F_Optimize_extended_ask_type(_Ztype)
            /* ERROR PROTECTION INSERTED (g0264I-Result) */
            if ErrorIn(try_3) {Result = try_3
            } else {
            g0264I = ToBoolean(OBJ(try_3))
            if (g0264I == CTRUE) { 
              _Ztype = ToType(OBJ(Core.F_CALL(C_mClaire_t1,ARGS(EID{_Ztype.Id(),0}))))
              Result = EID{_Ztype.Id(),0}
              } else {
              Result = EID{CFALSE.Id(),0}
              } 
            }
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            /*g_try(v2:"Result",loop:true) */
            Result = F_Optimize_range_infers_Variable(v,_Ztype)
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            { var arg_4 *ClaireType  
              _ = arg_4
              var try_6 EID 
              /*g_try(v2:"try_6",loop:false) */
              try_6 = Core.F_CALL(C_c_type,ARGS(self.Arg.ToEID()))
              /* ERROR PROTECTION INSERTED (arg_4-Result) */
              if ErrorIn(try_6) {Result = try_6
              } else {
              arg_4 = ToType(OBJ(try_6))
              { var arg_5 *ClaireType  
                _ = arg_5
                var try_7 EID 
                /*g_try(v2:"try_7",loop:false) */
                try_7 = Core.F_CALL(C_c_type,ARGS(self.Other.ToEID()))
                /* ERROR PROTECTION INSERTED (arg_5-Result) */
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
/* {1} The go function for: c_code(self:When,s:class) [status=1] */
func F_c_code_When (self *Language.When ,s *ClaireClass ) EID { 
    var Result EID 
    { var _Zv *ClaireAny   = self.Value
      { var v *ClaireVariable   = self.ClaireVar
        { var d *ClaireAny  
          var try_1 EID 
          /*g_try(v2:"try_1",loop:false) */
          try_1 = F_Optimize_daccess_any(_Zv,CTRUE)
          /* ERROR PROTECTION INSERTED (d-Result) */
          if ErrorIn(try_1) {Result = try_1
          } else {
          d = ANY(try_1)
          { var v2 *ClaireVariable   = F_Compile_Variable_I_symbol(F_append_symbol(self.ClaireVar.Pname,MakeString("test").Id()),self.ClaireVar.Index,C_any.Id())
            { var _Ztype *ClaireType  
              var try_2 EID 
              /*g_try(v2:"try_2",loop:false) */
              if (d != CNULL) { 
                try_2 = Core.F_CALL(C_c_type,ARGS(d.ToEID()))
                } else {
                try_2 = Core.F_CALL(C_c_type,ARGS(_Zv.ToEID()))
                } 
              /* ERROR PROTECTION INSERTED (_Ztype-Result) */
              if ErrorIn(try_2) {Result = try_2
              } else {
              _Ztype = ToType(OBJ(try_2))
              /*g_try(v2:"Result",loop:true) */
              var g0265I *ClaireBoolean  
              var try_3 EID 
              /*g_try(v2:"try_3",loop:false) */
              try_3 = F_Optimize_extended_ask_type(_Ztype)
              /* ERROR PROTECTION INSERTED (g0265I-Result) */
              if ErrorIn(try_3) {Result = try_3
              } else {
              g0265I = ToBoolean(OBJ(try_3))
              if (g0265I == CTRUE) { 
                _Ztype = ToType(OBJ(Core.F_CALL(C_mClaire_t1,ARGS(EID{_Ztype.Id(),0}))))
                Result = EID{_Ztype.Id(),0}
                } else {
                Result = EID{CFALSE.Id(),0}
                } 
              }
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              /*g_try(v2:"Result",loop:true) */
              Result = F_Optimize_range_infers_Variable(v,_Ztype)
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              var g0266I *ClaireBoolean  
              var try_4 EID 
              /*g_try(v2:"try_4",loop:false) */
              { 
                var v_and7 *ClaireBoolean  
                
                v_and7 = Core.F_known_ask_any(d)
                if (v_and7 == CFALSE) {try_4 = EID{CFALSE.Id(),0}
                } else { 
                  var try_5 EID 
                  /*g_try(v2:"try_5",loop:false) */
                  { var arg_6 *ClaireBoolean  
                    _ = arg_6
                    var try_7 EID 
                    /*g_try(v2:"try_7",loop:false) */
                    try_7 = F_Optimize_extended_ask_type(ToType(OBJ(Core.F_CALL(C_range,ARGS(Core.F_CALL(C_selector,ARGS(d.ToEID())))))))
                    /* ERROR PROTECTION INSERTED (arg_6-try_5) */
                    if ErrorIn(try_7) {try_5 = try_7
                    } else {
                    arg_6 = ToBoolean(OBJ(try_7))
                    try_5 = EID{arg_6.Not.Id(),0}
                    }
                    } 
                  /* ERROR PROTECTION INSERTED (v_and7-try_4) */
                  if ErrorIn(try_5) {try_4 = try_5
                  } else {
                  v_and7 = ToBoolean(OBJ(try_5))
                  if (v_and7 == CFALSE) {try_4 = EID{CFALSE.Id(),0}
                  } else { 
                    try_4 = EID{CTRUE.Id(),0}} 
                  } 
                }
                } 
              /* ERROR PROTECTION INSERTED (g0266I-Result) */
              if ErrorIn(try_4) {Result = try_4
              } else {
              g0266I = ToBoolean(OBJ(try_4))
              if (g0266I == CTRUE) { 
                { var _CL_obj *Language.Let   = Language.To_Let(new(Language.Let).Is(Language.C_Let))
                  _CL_obj.ClaireVar = v
                  /*Variable->Variable*/_CL_obj.Value = d
                  /*any->any*//*g_try(v2:"Result",loop:true) */
                  { 
                    var va_arg1 *Language.Let  
                    var va_arg2 *ClaireAny  
                    va_arg1 = _CL_obj
                    var try_8 EID 
                    /*g_try(v2:"try_8",loop:false) */
                    { var arg_9 *Language.CallMethod2  
                      _ = arg_9
                      var try_12 EID 
                      /*g_try(v2:"try_12",loop:false) */
                      { var _CL_obj *Language.CallMethod2   = Language.To_CallMethod2(new(Language.CallMethod2).Is(Language.C_Call_method2))
                        _CL_obj.Arg = ToMethod(Core.F__at_property1(Core.C_identical_ask,C_any).Id())
                        /*method->method*//*g_try(v2:"try_12",loop:false) */
                        { 
                          var va_arg1 *Language.CallMethod  
                          var va_arg2 *ClaireList  
                          va_arg1 = Language.To_CallMethod(_CL_obj.Id())
                          var try_13 EID 
                          /*g_try(v2:"try_13",loop:false) */
                          { 
                            var v_bag_arg *ClaireAny  
                            try_13= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                            ToList(OBJ(try_13)).AddFast(v.Id())
                            var try_14 EID 
                            /*g_try(v2:"try_14",loop:false) */
                            { var arg_15 *ClaireAny  
                              _ = arg_15
                              var try_16 EID 
                              /*g_try(v2:"try_16",loop:false) */
                              try_16 = Core.F_CALL(C_Compile_c_sort,ARGS(EID{v.Id(),0}))
                              /* ERROR PROTECTION INSERTED (arg_15-try_14) */
                              if ErrorIn(try_16) {try_14 = try_16
                              } else {
                              arg_15 = ANY(try_16)
                              try_14 = Core.F_CALL(C_c_code,ARGS(EID{CNULL,0},arg_15.ToEID()))
                              }
                              } 
                            /* ERROR PROTECTION INSERTED (v_bag_arg-try_13) */
                            if ErrorIn(try_14) {try_13 = try_14
                            } else {
                            v_bag_arg = ANY(try_14)
                            ToList(OBJ(try_13)).AddFast(v_bag_arg)}
                            } 
                          /* ERROR PROTECTION INSERTED (va_arg2-try_12) */
                          if ErrorIn(try_13) {try_12 = try_13
                          } else {
                          va_arg2 = ToList(OBJ(try_13))
                          va_arg1.Args = va_arg2
                          /*list->list*/try_12 = EID{va_arg2.Id(),0}
                          }
                          } 
                        /* ERROR PROTECTION INSERTED (try_12-try_12) */
                        if !ErrorIn(try_12) {
                        try_12 = EID{_CL_obj.Id(),0}
                        }
                        } 
                      /* ERROR PROTECTION INSERTED (arg_9-try_8) */
                      if ErrorIn(try_12) {try_8 = try_12
                      } else {
                      arg_9 = Language.To_CallMethod2(OBJ(try_12))
                      { var arg_10 *ClaireAny  
                        _ = arg_10
                        var try_17 EID 
                        /*g_try(v2:"try_17",loop:false) */
                        try_17 = Core.F_CALL(C_c_code,ARGS(self.Other.ToEID(),EID{s.Id(),0}))
                        /* ERROR PROTECTION INSERTED (arg_10-try_8) */
                        if ErrorIn(try_17) {try_8 = try_17
                        } else {
                        arg_10 = ANY(try_17)
                        { var arg_11 *ClaireAny  
                          _ = arg_11
                          var try_18 EID 
                          /*g_try(v2:"try_18",loop:false) */
                          try_18 = Core.F_CALL(C_c_code,ARGS(self.Arg.ToEID(),EID{s.Id(),0}))
                          /* ERROR PROTECTION INSERTED (arg_11-try_8) */
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
                    /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                    if ErrorIn(try_8) {Result = try_8
                    } else {
                    va_arg2 = ANY(try_8)
                    va_arg1.Arg = va_arg2
                    /*any->any*/Result = va_arg2.ToEID()
                    }
                    } 
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  Result = EID{_CL_obj.Id(),0}
                  }
                  } 
                } else {
                var g0267I *ClaireBoolean  
                var try_19 EID 
                /*g_try(v2:"try_19",loop:false) */
                { 
                  var v_and8 *ClaireBoolean  
                  
                  var try_20 EID 
                  /*g_try(v2:"try_20",loop:false) */
                  { var arg_21 *ClaireAny  
                    _ = arg_21
                    var try_22 EID 
                    /*g_try(v2:"try_22",loop:false) */
                    try_22 = Core.F_CALL(C_Compile_c_sort,ARGS(EID{v.Id(),0}))
                    /* ERROR PROTECTION INSERTED (arg_21-try_20) */
                    if ErrorIn(try_22) {try_20 = try_22
                    } else {
                    arg_21 = ANY(try_22)
                    try_20 = EID{Equal(arg_21,C_any.Id()).Id(),0}
                    }
                    } 
                  /* ERROR PROTECTION INSERTED (v_and8-try_19) */
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
                /* ERROR PROTECTION INSERTED (g0267I-Result) */
                if ErrorIn(try_19) {Result = try_19
                } else {
                g0267I = ToBoolean(OBJ(try_19))
                if (g0267I == CTRUE) { 
                  { var arg_23 *Language.Let  
                    _ = arg_23
                    { var _CL_obj *Language.Let   = Language.To_Let(new(Language.Let).Is(Language.C_Let))
                      _CL_obj.ClaireVar = v
                      /*Variable->Variable*/_CL_obj.Value = _Zv
                      /*any->any*/{ 
                        var va_arg1 *Language.Let  
                        var va_arg2 *ClaireAny  
                        va_arg1 = _CL_obj
                        { var arg_24 *Language.Call  
                          _ = arg_24
                          { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                            _CL_obj.Selector = ToProperty(Core.C__I_equal.Id())
                            /*property->property*/_CL_obj.Args = MakeConstantList(v.Id(),CNULL)
                            /*list->list*/arg_24 = _CL_obj
                            } 
                          va_arg2 = Language.C_If.Make(arg_24.Id(),self.Arg,self.Other)
                          } 
                        va_arg1.Arg = va_arg2
                        /*any->any*/} 
                      arg_23 = _CL_obj
                      } 
                    Result = Core.F_CALL(C_c_code,ARGS(EID{arg_23.Id(),0},EID{s.Id(),0}))
                    } 
                  } else {
                  { var arg_25 *Language.Let  
                    _ = arg_25
                    var try_26 EID 
                    /*g_try(v2:"try_26",loop:false) */
                    { var _CL_obj *Language.Let   = Language.To_Let(new(Language.Let).Is(Language.C_Let))
                      _CL_obj.ClaireVar = v2
                      /*Variable->Variable*/_CL_obj.Value = _Zv
                      /*any->any*//*g_try(v2:"try_26",loop:false) */
                      { 
                        var va_arg1 *Language.Let  
                        var va_arg2 *ClaireAny  
                        va_arg1 = _CL_obj
                        var try_27 EID 
                        /*g_try(v2:"try_27",loop:false) */
                        { var arg_28 *Language.Call  
                          _ = arg_28
                          { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                            _CL_obj.Selector = ToProperty(Core.C__I_equal.Id())
                            /*property->property*/_CL_obj.Args = MakeConstantList(v2.Id(),CNULL)
                            /*list->list*/arg_28 = _CL_obj
                            } 
                          { var arg_29 *Language.Let  
                            _ = arg_29
                            { var _CL_obj *Language.Let   = Language.To_Let(new(Language.Let).Is(Language.C_Let))
                              _CL_obj.ClaireVar = v
                              /*Variable->Variable*/{ 
                                var va_arg1 *Language.Let  
                                var va_arg2 *ClaireAny  
                                va_arg1 = _CL_obj
                                { var _CL_obj *Language.Cast   = Language.To_Cast(new(Language.Cast).Is(Language.C_Cast))
                                  _CL_obj.Arg = v2.Id()
                                  /*any->any*/_CL_obj.SetArg = _Ztype
                                  /*type->type*/va_arg2 = _CL_obj.Id()
                                  } 
                                va_arg1.Value = va_arg2
                                /*any->any*/} 
                              _CL_obj.Arg = self.Arg
                              /*any->any*/arg_29 = _CL_obj
                              } 
                            { var arg_30 *ClaireAny  
                              _ = arg_30
                              var try_31 EID 
                              /*g_try(v2:"try_31",loop:false) */
                              try_31 = Core.F_CALL(C_c_code,ARGS(self.Other.ToEID(),EID{s.Id(),0}))
                              /* ERROR PROTECTION INSERTED (arg_30-try_27) */
                              if ErrorIn(try_31) {try_27 = try_31
                              } else {
                              arg_30 = ANY(try_31)
                              try_27 = Language.C_If.Make(arg_28.Id(),arg_29.Id(),arg_30).ToEID()
                              }
                              } 
                            } 
                          } 
                        /* ERROR PROTECTION INSERTED (va_arg2-try_26) */
                        if ErrorIn(try_27) {try_26 = try_27
                        } else {
                        va_arg2 = ANY(try_27)
                        va_arg1.Arg = va_arg2
                        /*any->any*/try_26 = va_arg2.ToEID()
                        }
                        } 
                      /* ERROR PROTECTION INSERTED (try_26-try_26) */
                      if !ErrorIn(try_26) {
                      try_26 = EID{_CL_obj.Id(),0}
                      }
                      } 
                    /* ERROR PROTECTION INSERTED (arg_25-Result) */
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
/* {1} The go function for: c_type(self:For) [status=1] */
func F_c_type_For (self *Language.For ) EID { 
    var Result EID 
    { var arg_1 *ClaireType  
      _ = arg_1
      var try_2 EID 
      /*g_try(v2:"try_2",loop:false) */
      try_2 = F_Compile_return_type_any(self.Arg)
      /* ERROR PROTECTION INSERTED (arg_1-Result) */
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
  
/* {1} The go function for: infers_from(t:type,self:any) [status=0] */
func F_Optimize_infers_from_type (t *ClaireType ,self *ClaireAny ) *ClaireType  { 
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
/* {1} The go function for: c_code(self:For,s:class) [status=1] */
func F_c_code_For (self *Language.For ,s *ClaireClass ) EID { 
    var Result EID 
    { var sx *ClaireAny   = self.SetArg
      { var ns int  = C_compiler.Safety
        _ = ns
        { var vold *ClaireVariable   = self.ClaireVar
          { var v *ClaireVariable   = F_Compile_Variable_I_symbol(vold.Pname,vold.Index,Core.F_get_property(C_range,ToObject(vold.Id())))
            { var narg *ClaireAny   = Language.F_substitution_any(self.Arg,vold,v.Id())
              { var scs *ClaireAny  
                var try_1 EID 
                /*g_try(v2:"try_1",loop:false) */
                try_1 = F_Optimize_c_inline_arg_ask_any(sx)
                /* ERROR PROTECTION INSERTED (scs-Result) */
                if ErrorIn(try_1) {Result = try_1
                } else {
                scs = ANY(try_1)
                if (sx.Isa.IsIn(Core.C_global_variable) == CTRUE) { 
                  { var g0268 *Core.GlobalVariable   = Core.ToGlobalVariable(sx)
                    if (F_boolean_I_any(g0268.Range.Id()).Id() != CTRUE.Id()) { 
                      self.SetArg = g0268.Value
                      /*any->any*/g0268 = Core.ToGlobalVariable(g0268.Value)
                      } 
                    sx = g0268.Id()
                    } 
                  }  else if (C_class.Id() == sx.Isa.Id()) { 
                  { var g0269 *ClaireClass   = ToClass(sx)
                    if ((g0269.Open <= 1) && 
                        (F_boolean_I_any(g0269.Subclass.Id()).Id() != CTRUE.Id())) { 
                      { 
                        var va_arg1 *Language.Iteration  
                        var va_arg2 *ClaireAny  
                        va_arg1 = Language.To_Iteration(self.Id())
                        { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                          _CL_obj.Selector = C_instances
                          /*property->property*/_CL_obj.Args = MakeConstantList(g0269.Id())
                          /*list->list*/va_arg2 = _CL_obj.Id()
                          } 
                        va_arg1.SetArg = va_arg2
                        /*any->any*/} 
                      } 
                    } 
                  } 
                { var _Zt *ClaireType  
                  var try_2 EID 
                  /*g_try(v2:"try_2",loop:false) */
                  try_2 = Core.F_CALL(C_c_type,ARGS(self.SetArg.ToEID()))
                  /* ERROR PROTECTION INSERTED (_Zt-Result) */
                  if ErrorIn(try_2) {Result = try_2
                  } else {
                  _Zt = ToType(OBJ(try_2))
                  { var _Zt2 *ClaireType   = F_Optimize_pmember_type(_Zt)
                    { var _Zt3 *ClaireType  
                      _ = _Zt3
                      var try_3 EID 
                      /*g_try(v2:"try_3",loop:false) */
                      { var arg_4 *ClaireType  
                        _ = arg_4
                        var try_5 EID 
                        /*g_try(v2:"try_5",loop:false) */
                        { var arg_6 *Language.Call  
                          _ = arg_6
                          { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                            _CL_obj.Selector = C_set_I
                            /*property->property*/_CL_obj.Args = MakeConstantList(sx)
                            /*list->list*/arg_6 = _CL_obj
                            } 
                          try_5 = Core.F_CALL(C_c_type,ARGS(EID{arg_6.Id(),0}))
                          } 
                        /* ERROR PROTECTION INSERTED (arg_4-try_3) */
                        if ErrorIn(try_5) {try_3 = try_5
                        } else {
                        arg_4 = ToType(OBJ(try_5))
                        try_3 = EID{F_Optimize_pmember_type(arg_4).Id(),0}
                        }
                        } 
                      /* ERROR PROTECTION INSERTED (_Zt3-Result) */
                      if ErrorIn(try_3) {Result = try_3
                      } else {
                      _Zt3 = ToType(OBJ(try_3))
                      if (_Zt3.Included(_Zt2) == CTRUE) { 
                        _Zt2 = _Zt3
                        } 
                      F_Optimize_range_infers_for_Variable(v,_Zt2,_Zt)
                      C_compiler.Safety = 1
                      /*integer->integer*/_Zt2 = v.Range
                      v.Range = F_Optimize_ptype_type(_Zt2)
                      /*type->type*/{ var m *ClaireAny   = F_Optimize_Iterate_I_Iteration(Language.To_Iteration(self.Id()))
                        if (C_method.Id() != m.Isa.Id()) { 
                          { var m2 *ClaireAny   = F_Optimize_restriction_I_property(Language.C_iterate,MakeConstantList(_Zt.Id(),MakeConstantSet(v.Id()).Id(),C_any.Id()),CTRUE)
                            if (C_method.Id() == m2.Isa.Id()) { 
                              { var g0270 *ClaireMethod   = ToMethod(m2)
                                _ = g0270
                                m = g0270.Id()
                                } 
                              } 
                            } 
                          } 
                        C_compiler.Safety = ns
                        /*integer->integer*/v.Range = _Zt2
                        /*type->type*/var g0273I *ClaireBoolean  
                        if (C_method.Id() == m.Isa.Id()) { 
                          { var g0271 *ClaireMethod   = ToMethod(m)
                            _ = g0271
                            g0273I = g0271.Inline_ask
                            } 
                          } else {
                          g0273I = CFALSE
                          } 
                        if (g0273I == CTRUE) { 
                          
                          if (F_Optimize_sort_abstract_ask_type(v.Range) == CTRUE) { 
                            v.Range = To_Union(v.Range.Id()).T2
                            /*type->type*/} 
                          Result = F_Optimize_c_inline_method1(ToMethod(m),MakeConstantList(Language.F_instruction_copy_any(self.SetArg),v.Id(),narg),s)
                          }  else if (F_boolean_I_any(scs) == CTRUE) { 
                          { var arg_7 *Language.For  
                            _ = arg_7
                            { var _CL_obj *Language.For   = Language.To_For(new(Language.For).Is(Language.C_For))
                              _CL_obj.ClaireVar = v
                              /*Variable->Variable*/_CL_obj.SetArg = scs
                              /*any->any*/_CL_obj.Arg = narg
                              /*any->any*/arg_7 = _CL_obj
                              } 
                            Result = Core.F_CALL(C_c_code,ARGS(EID{arg_7.Id(),0},EID{s.Id(),0}))
                            } 
                          } else {
                          var g0274I *ClaireBoolean  
                          if (sx.Isa.IsIn(Language.C_Call) == CTRUE) { 
                            { var g0272 *Language.Call   = Language.To_Call(sx)
                              _ = g0272
                              g0274I = Equal(g0272.Selector.Id(),Core.C_Id.Id())
                              } 
                            } else {
                            g0274I = CFALSE
                            } 
                          if (g0274I == CTRUE) { 
                            Result = F_Optimize_c_code_multiple_For(self,_Zt,s)
                            } else {
                            { var _CL_obj *Language.For   = Language.To_For(new(Language.For).Is(Language.C_For))
                              _CL_obj.ClaireVar = v
                              /*Variable->Variable*//*g_try(v2:"Result",loop:true) */
                              { 
                                var va_arg1 *Language.Iteration  
                                var va_arg2 *ClaireAny  
                                va_arg1 = Language.To_Iteration(_CL_obj.Id())
                                var try_8 EID 
                                /*g_try(v2:"try_8",loop:false) */
                                try_8 = F_Optimize_enumerate_code_any(self.SetArg,_Zt)
                                /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                                if ErrorIn(try_8) {Result = try_8
                                } else {
                                va_arg2 = ANY(try_8)
                                va_arg1.SetArg = va_arg2
                                /*any->any*/Result = va_arg2.ToEID()
                                }
                                } 
                              /* ERROR PROTECTION INSERTED (Result-Result) */
                              if !ErrorIn(Result) {
                              /*g_try(v2:"Result",loop:true) */
                              { 
                                var va_arg1 *Language.Iteration  
                                var va_arg2 *ClaireAny  
                                va_arg1 = Language.To_Iteration(_CL_obj.Id())
                                var try_9 EID 
                                /*g_try(v2:"try_9",loop:false) */
                                try_9 = Core.F_CALL(C_c_code,ARGS(narg.ToEID(),EID{C_void.Id(),0}))
                                /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                                if ErrorIn(try_9) {Result = try_9
                                } else {
                                va_arg2 = ANY(try_9)
                                va_arg1.Arg = va_arg2
                                /*any->any*/Result = va_arg2.ToEID()
                                }
                                } 
                              /* ERROR PROTECTION INSERTED (Result-Result) */
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
/* {1} The go function for: c_code_multiple(self:For,%t:type,s:class) [status=1] */
func F_Optimize_c_code_multiple_For (self *Language.For ,_Zt *ClaireType ,s *ClaireClass ) EID { 
    var Result EID 
    { var v *ClaireVariable   = self.ClaireVar
      { var sx *ClaireAny   = Language.To_Call(self.SetArg).Args.At(1-1)
        { var v2 *ClaireVariable   = F_Compile_Variable_I_symbol(F_append_symbol(v.Pname,MakeString("test").Id()),self.ClaireVar.Index,_Zt.Id())
          { var n *Language.Let  
            var try_1 EID 
            /*g_try(v2:"try_1",loop:false) */
            { var _CL_obj *Language.Let   = Language.To_Let(new(Language.Let).Is(Language.C_Let))
              _CL_obj.ClaireVar = v2
              /*Variable->Variable*/_CL_obj.Value = sx
              /*any->any*//*g_try(v2:"try_1",loop:false) */
              { 
                var va_arg1 *Language.Let  
                var va_arg2 *ClaireAny  
                va_arg1 = _CL_obj
                var try_2 EID 
                /*g_try(v2:"try_2",loop:false) */
                { var _CL_obj *Language.For   = Language.To_For(new(Language.For).Is(Language.C_For))
                  _CL_obj.ClaireVar = self.ClaireVar
                  /*Variable->Variable*//*g_try(v2:"try_2",loop:false) */
                  { 
                    var va_arg1 *Language.Iteration  
                    var va_arg2 *ClaireAny  
                    va_arg1 = Language.To_Iteration(_CL_obj.Id())
                    var try_3 EID 
                    /*g_try(v2:"try_3",loop:false) */
                    try_3 = F_Optimize_enumerate_code_any(sx,_Zt)
                    /* ERROR PROTECTION INSERTED (va_arg2-try_2) */
                    if ErrorIn(try_3) {try_2 = try_3
                    } else {
                    va_arg2 = ANY(try_3)
                    va_arg1.SetArg = va_arg2
                    /*any->any*/try_2 = va_arg2.ToEID()
                    }
                    } 
                  /* ERROR PROTECTION INSERTED (try_2-try_2) */
                  if !ErrorIn(try_2) {
                  _CL_obj.Arg = self.Arg
                  /*any->any*/try_2 = EID{_CL_obj.Id(),0}
                  }
                  } 
                /* ERROR PROTECTION INSERTED (va_arg2-try_1) */
                if ErrorIn(try_2) {try_1 = try_2
                } else {
                va_arg2 = ANY(try_2)
                va_arg1.Arg = va_arg2
                /*any->any*/try_1 = va_arg2.ToEID()
                }
                } 
              /* ERROR PROTECTION INSERTED (try_1-try_1) */
              if !ErrorIn(try_1) {
              try_1 = EID{_CL_obj.Id(),0}
              }
              } 
            /* ERROR PROTECTION INSERTED (n-Result) */
            if ErrorIn(try_1) {Result = try_1
            } else {
            n = Language.To_Let(OBJ(try_1))
            Core.F_tformat_string(MakeString("---- note: use an expended iteration for {~S} \n"),0,MakeConstantList(self.Id()))
            /*g_try(v2:"Result",loop:true) */
            { 
              var r *ClaireRestriction  
              _ = r
              var r_iter *ClaireAny  
              Result= EID{CFALSE.Id(),0}
              for _,r_iter = range(Language.C_iterate.Restrictions.ValuesO()){ 
                r = ToRestriction(r_iter)
                var loop_4 EID 
                _ = loop_4
                /*g_try(v2:"loop_4",loop:tuple("Result", EID)) */
                var g0275I *ClaireBoolean  
                var try_5 EID 
                /*g_try(v2:"try_5",loop:false) */
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
                        /*g_try(v2:"try_6",loop:false) */
                        try_6 = Core.F_BELONG(v.Id(),r.Domain.ValuesO()[2-1])
                        /* ERROR PROTECTION INSERTED (v_and8-try_5) */
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
                /* ERROR PROTECTION INSERTED (g0275I-loop_4) */
                if ErrorIn(try_5) {loop_4 = try_5
                } else {
                g0275I = ToBoolean(OBJ(try_5))
                if (g0275I == CTRUE) { 
                  { var vnew *ClaireVariable  
                    { var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
                      _CL_obj.Pname = v.Pname
                      /*symbol->symbol*/_CL_obj.Range = v.Range
                      /*type->type*/_CL_obj.Index = v.Index
                      /*integer->integer*/vnew = _CL_obj
                      } 
                    { var narg *ClaireAny   = Language.F_substitution_any(self.Arg,v,vnew.Id())
                      _ = narg
                      { 
                        var va_arg1 *Language.Let  
                        var va_arg2 *ClaireAny  
                        va_arg1 = n
                        var try_7 EID 
                        /*g_try(v2:"try_7",loop:false) */
                        { var arg_8 *Language.Call  
                          _ = arg_8
                          { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                            _CL_obj.Selector = ToProperty(C__Z.Id())
                            /*property->property*/_CL_obj.Args = MakeConstantList(v2.Id(),r.Domain.ValuesO()[1-1])
                            /*list->list*/arg_8 = _CL_obj
                            } 
                          { var arg_9 *ClaireAny  
                            _ = arg_9
                            var try_10 EID 
                            /*g_try(v2:"try_10",loop:false) */
                            if (F_Optimize_sort_abstract_ask_type(vnew.Range) == CTRUE) { 
                              vnew.Range = To_Union(v.Range.Id()).T2
                              /*type->type*/} 
                            try_10 = F_Optimize_c_inline_method1(ToMethod(r.Id()),MakeConstantList(v2.Id(),vnew.Id(),narg),s)
                            /* ERROR PROTECTION INSERTED (arg_9-try_7) */
                            if ErrorIn(try_10) {try_7 = try_10
                            } else {
                            arg_9 = ANY(try_10)
                            try_7 = Language.C_If.Make(arg_8.Id(),arg_9,n.Arg).ToEID()
                            }
                            } 
                          } 
                        /* ERROR PROTECTION INSERTED (va_arg2-loop_4) */
                        if ErrorIn(try_7) {loop_4 = try_7
                        } else {
                        va_arg2 = ANY(try_7)
                        va_arg1.Arg = va_arg2
                        /*any->any*/loop_4 = va_arg2.ToEID()
                        }
                        } 
                      } 
                    } 
                  } else {
                  loop_4 = EID{CFALSE.Id(),0}
                  } 
                }
                /* ERROR PROTECTION INSERTED (loop_4-Result) */
                if ErrorIn(loop_4) {Result = loop_4
                break
                } else {
                }
                } 
              } 
            /* ERROR PROTECTION INSERTED (Result-Result) */
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
/* {1} The go function for: c_type(self:Iteration) [status=1] */
func F_c_type_Iteration (self *Language.Iteration ) EID { 
    var Result EID 
    { var _Zt *ClaireType  
      var try_1 EID 
      /*g_try(v2:"try_1",loop:false) */
      try_1 = Core.F_CALL(C_c_type,ARGS(self.SetArg.ToEID()))
      /* ERROR PROTECTION INSERTED (_Zt-Result) */
      if ErrorIn(try_1) {Result = try_1
      } else {
      _Zt = ToType(OBJ(try_1))
      F_Optimize_range_infers_for_Variable(self.ClaireVar,F_Optimize_pmember_type(_Zt),_Zt)
      if (Core.F_get_property(C_of,ToObject(self.Id())) != CNULL) { 
        { var arg_2 *ClaireClass  
          _ = arg_2
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
          _ = arg_3
          if ((self.Isa.IsIn(Language.C_Select) == CTRUE) || 
              (self.Isa.IsIn(Language.C_Image) == CTRUE)) { 
            arg_3 = C_set
            } else {
            arg_3 = C_list
            } 
          { var arg_4 *ClaireType  
            _ = arg_4
            var try_5 EID 
            /*g_try(v2:"try_5",loop:false) */
            if ((self.Isa.IsIn(Language.C_Select) == CTRUE) || 
                (self.Isa.IsIn(Language.C_Lselect) == CTRUE)) { 
              { var arg_6 *ClaireType  
                _ = arg_6
                var try_7 EID 
                /*g_try(v2:"try_7",loop:false) */
                try_7 = Core.F_CALL(C_c_type,ARGS(self.SetArg.ToEID()))
                /* ERROR PROTECTION INSERTED (arg_6-try_5) */
                if ErrorIn(try_7) {try_5 = try_7
                } else {
                arg_6 = ToType(OBJ(try_7))
                try_5 = EID{F_Optimize_pmember_type(arg_6).Id(),0}
                }
                } 
              } else {
              { var arg_8 *ClaireType  
                _ = arg_8
                var try_9 EID 
                /*g_try(v2:"try_9",loop:false) */
                try_9 = Core.F_CALL(C_c_type,ARGS(self.Arg.ToEID()))
                /* ERROR PROTECTION INSERTED (arg_8-try_5) */
                if ErrorIn(try_9) {try_5 = try_9
                } else {
                arg_8 = ToType(OBJ(try_9))
                try_5 = EID{F_Optimize_ptype_type(arg_8).Id(),0}
                }
                } 
              } 
            /* ERROR PROTECTION INSERTED (arg_4-Result) */
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
/* {1} The go function for: c_code(self:Iteration) [status=1] */
func F_c_code_Iteration (self *Language.Iteration ) EID { 
    var Result EID 
    { var sx *ClaireAny   = self.SetArg
      { var _Zt *ClaireType  
        var try_1 EID 
        /*g_try(v2:"try_1",loop:false) */
        try_1 = Core.F_CALL(C_c_type,ARGS(sx.ToEID()))
        /* ERROR PROTECTION INSERTED (_Zt-Result) */
        if ErrorIn(try_1) {Result = try_1
        } else {
        _Zt = ToType(OBJ(try_1))
        if (self.Isa.IsIn(Language.C_For) == CTRUE) { 
          Result = Core.F_CALL(C_c_code,ARGS(EID{self.Id(),0},EID{C_any.Id(),0}))
          } else {
          var g0278I *ClaireBoolean  
          if (self.Isa.IsIn(Language.C_Collect) == CTRUE) { 
            g0278I = MakeBoolean((_Zt.Included(ToType(C_list.Id())) == CTRUE) || (_Zt.Included(ToType(C_set.Id())) == CTRUE))
            } else {
            g0278I = CFALSE
            } 
          if (g0278I == CTRUE) { 
            F_Optimize_range_infers_for_Variable(self.ClaireVar,F_Optimize_pmember_type(_Zt),_Zt)
            { var ty *ClaireType  
              var try_2 EID 
              /*g_try(v2:"try_2",loop:false) */
              { var arg_3 *ClaireType  
                _ = arg_3
                var try_4 EID 
                /*g_try(v2:"try_4",loop:false) */
                try_4 = Core.F_CALL(C_c_type,ARGS(self.Arg.ToEID()))
                /* ERROR PROTECTION INSERTED (arg_3-try_2) */
                if ErrorIn(try_4) {try_2 = try_4
                } else {
                arg_3 = ToType(OBJ(try_4))
                try_2 = EID{F_Optimize_ptype_type(arg_3).Id(),0}
                }
                } 
              /* ERROR PROTECTION INSERTED (ty-Result) */
              if ErrorIn(try_2) {Result = try_2
              } else {
              ty = ToType(OBJ(try_2))
              { var x *Language.Collect  
                var try_5 EID 
                /*g_try(v2:"try_5",loop:false) */
                { var _CL_obj *Language.Collect   = Language.To_Collect(new(Language.Collect).Is(Language.C_Collect))
                  _CL_obj.ClaireVar = self.ClaireVar
                  /*Variable->Variable*//*g_try(v2:"try_5",loop:false) */
                  { 
                    var va_arg1 *Language.Iteration  
                    var va_arg2 *ClaireAny  
                    va_arg1 = Language.To_Iteration(_CL_obj.Id())
                    var try_6 EID 
                    /*g_try(v2:"try_6",loop:false) */
                    try_6 = F_Compile_c_strict_code_any(sx,_Zt.Class_I())
                    /* ERROR PROTECTION INSERTED (va_arg2-try_5) */
                    if ErrorIn(try_6) {try_5 = try_6
                    } else {
                    va_arg2 = ANY(try_6)
                    va_arg1.SetArg = va_arg2
                    /*any->any*/try_5 = va_arg2.ToEID()
                    }
                    } 
                  /* ERROR PROTECTION INSERTED (try_5-try_5) */
                  if !ErrorIn(try_5) {
                  /*g_try(v2:"try_5",loop:false) */
                  { 
                    var va_arg1 *Language.Iteration  
                    var va_arg2 *ClaireAny  
                    va_arg1 = Language.To_Iteration(_CL_obj.Id())
                    var try_7 EID 
                    /*g_try(v2:"try_7",loop:false) */
                    try_7 = Core.F_CALL(C_c_code,ARGS(self.Arg.ToEID(),EID{C_any.Id(),0}))
                    /* ERROR PROTECTION INSERTED (va_arg2-try_5) */
                    if ErrorIn(try_7) {try_5 = try_7
                    } else {
                    va_arg2 = ANY(try_7)
                    va_arg1.Arg = va_arg2
                    /*any->any*/try_5 = va_arg2.ToEID()
                    }
                    } 
                  /* ERROR PROTECTION INSERTED (try_5-try_5) */
                  if !ErrorIn(try_5) {
                  try_5 = EID{_CL_obj.Id(),0}
                  }}
                  } 
                /* ERROR PROTECTION INSERTED (x-Result) */
                if ErrorIn(try_5) {Result = try_5
                } else {
                x = Language.To_Collect(OBJ(try_5))
                /*g_try(v2:"Result",loop:true) */
                if (ty.Id() == C_void.Id()) { 
                  Result = Core.F_CALL(C_Compile_Cerror,ARGS(EID{MakeString("[205] use of void expression ~S in ~S").Id(),0},self.Arg.ToEID(),EID{self.Id(),0}))
                  } else {
                  Result = EID{CFALSE.Id(),0}
                  } 
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                if (Core.F_get_property(C_of,ToObject(self.Id())) != CNULL) { 
                  if ((C_compiler.Safety >= 2) || 
                      (ty.Included(ToType(OBJ(Core.F_CALL(C_of,ARGS(EID{self.Id(),0}))))) == CTRUE)) { 
                    x.Of = ToType(OBJ(Core.F_CALL(C_of,ARGS(EID{self.Id(),0}))))
                    /*type->type*/Result = EID{x.Id(),0}
                    } else {
                    F_Compile_warn_void()
                    /*g_try(v2:"Result",loop:true) */
                    { var arg_8 *ClaireList  
                      _ = arg_8
                      var try_9 EID 
                      /*g_try(v2:"try_9",loop:false) */
                      { 
                        var v_bag_arg *ClaireAny  
                        try_9= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                        ToList(OBJ(try_9)).AddFast(self.Id())
                        var try_10 EID 
                        /*g_try(v2:"try_10",loop:false) */
                        try_10 = Core.F_CALL(C_c_type,ARGS(self.Arg.ToEID()))
                        /* ERROR PROTECTION INSERTED (v_bag_arg-try_9) */
                        if ErrorIn(try_10) {try_9 = try_10
                        } else {
                        v_bag_arg = ANY(try_10)
                        ToList(OBJ(try_9)).AddFast(v_bag_arg)
                        ToList(OBJ(try_9)).AddFast(ANY(Core.F_CALL(C_of,ARGS(EID{self.Id(),0}))))}
                        } 
                      /* ERROR PROTECTION INSERTED (arg_8-Result) */
                      if ErrorIn(try_9) {Result = try_9
                      } else {
                      arg_8 = ToList(OBJ(try_9))
                      Result = Core.F_tformat_string(MakeString("unsafe typed collect (~S): ~S not in ~S [261]\n"),1,arg_8)
                      }
                      } 
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    { var arg_11 *Language.Call  
                      _ = arg_11
                      { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                        _CL_obj.Selector = Core.C_check_in
                        /*property->property*/_CL_obj.Args = MakeConstantList(x.Id(),C_list.Id(),ANY(Core.F_CALL(C_of,ARGS(EID{self.Id(),0}))))
                        /*list->list*/arg_11 = _CL_obj
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
                  _ = arg_12
                  C_OPT.MaxVars = (C_OPT.MaxVars+1)
                  /*integer->integer*/arg_12 = 0
                  { var arg_13 *ClaireClass  
                    _ = arg_13
                    if (self.Isa.IsIn(Language.C_Image) == CTRUE) { 
                      arg_13 = C_set
                      } else {
                      arg_13 = C_list
                      } 
                    v = F_Compile_Variable_I_symbol(F_append_symbol(self.ClaireVar.Pname,MakeString("_bag").Id()),arg_12,arg_13.Id())
                    } 
                  } 
                
                /*g_try(v2:"Result",loop:true) */
                if (Core.F_get_property(C_of,ToObject(self.Id())) != CNULL) { 
                  { var _ZtypeIn *ClaireType  
                    var try_14 EID 
                    /*g_try(v2:"try_14",loop:false) */
                    try_14 = Core.F_CALL(C_c_type,ARGS(self.Arg.ToEID()))
                    /* ERROR PROTECTION INSERTED (_ZtypeIn-Result) */
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
                      /*type->type*/Result = EID{va_arg2.Id(),0}
                      } 
                    }
                    } 
                  } else {
                  if (C_set.Id() == val.Isa.Id()) { 
                    { var _CL_obj *Language.Set   = Language.To_Set(new(Language.Set).Is(Language.C_Set))
                      _CL_obj.Of = ToType(CEMPTY.Id())
                      /*type->type*/val = _CL_obj.Id()
                      } 
                    } else {
                    { var _CL_obj *Language.List   = Language.To_List(new(Language.List).Is(Language.C_List))
                      _CL_obj.Of = ToType(CEMPTY.Id())
                      /*type->type*/val = _CL_obj.Id()
                      } 
                    } 
                  Result = val.ToEID()
                  } 
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                { var _CL_obj *Language.Let   = Language.To_Let(new(Language.Let).Is(Language.C_Let))
                  _CL_obj.ClaireVar = v
                  /*Variable->Variable*/_CL_obj.Value = val
                  /*any->any*//*g_try(v2:"Result",loop:true) */
                  { 
                    var va_arg1 *Language.Let  
                    var va_arg2 *ClaireAny  
                    va_arg1 = _CL_obj
                    var try_15 EID 
                    /*g_try(v2:"try_15",loop:false) */
                    { var _CL_obj *Language.Do   = Language.To_Do(new(Language.Do).Is(Language.C_Do))
                      /*g_try(v2:"try_15",loop:false) */
                      { 
                        var va_arg1 *Language.Do  
                        var va_arg2 *ClaireList  
                        va_arg1 = _CL_obj
                        var try_16 EID 
                        /*g_try(v2:"try_16",loop:false) */
                        { 
                          var v_bag_arg *ClaireAny  
                          try_16= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                          var try_17 EID 
                          /*g_try(v2:"try_17",loop:false) */
                          { var arg_18 *Language.For  
                            _ = arg_18
                            { var _CL_obj *Language.For   = Language.To_For(new(Language.For).Is(Language.C_For))
                              _CL_obj.ClaireVar = self.ClaireVar
                              /*Variable->Variable*/_CL_obj.SetArg = sx
                              /*any->any*/{ 
                                var va_arg1 *Language.Iteration  
                                var va_arg2 *ClaireAny  
                                va_arg1 = Language.To_Iteration(_CL_obj.Id())
                                { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                  _CL_obj.Selector = ToProperty(C_add_I.Id())
                                  /*property->property*/_CL_obj.Args = MakeConstantList(v.Id(),self.Arg)
                                  /*list->list*/va_arg2 = _CL_obj.Id()
                                  } 
                                va_arg1.Arg = va_arg2
                                /*any->any*/} 
                              arg_18 = _CL_obj
                              } 
                            try_17 = Core.F_CALL(C_c_code,ARGS(EID{arg_18.Id(),0},EID{C_any.Id(),0}))
                            } 
                          /* ERROR PROTECTION INSERTED (v_bag_arg-try_16) */
                          if ErrorIn(try_17) {try_16 = try_17
                          } else {
                          v_bag_arg = ANY(try_17)
                          ToList(OBJ(try_16)).AddFast(v_bag_arg)
                          ToList(OBJ(try_16)).AddFast(v.Id())}
                          } 
                        /* ERROR PROTECTION INSERTED (va_arg2-try_15) */
                        if ErrorIn(try_16) {try_15 = try_16
                        } else {
                        va_arg2 = ToList(OBJ(try_16))
                        va_arg1.Args = va_arg2
                        /*list->list*/try_15 = EID{va_arg2.Id(),0}
                        }
                        } 
                      /* ERROR PROTECTION INSERTED (try_15-try_15) */
                      if !ErrorIn(try_15) {
                      try_15 = EID{_CL_obj.Id(),0}
                      }
                      } 
                    /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                    if ErrorIn(try_15) {Result = try_15
                    } else {
                    va_arg2 = ANY(try_15)
                    va_arg1.Arg = va_arg2
                    /*any->any*/Result = va_arg2.ToEID()
                    }
                    } 
                  /* ERROR PROTECTION INSERTED (Result-Result) */
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
/* {1} The go function for: c_code(self:Select) [status=1] */
func F_c_code_Select (self *Language.Select ) EID { 
    var Result EID 
    Result = F_Optimize_c_code_select_Iteration(Language.To_Iteration(self.Id()),C_set)
    return Result} 
  
// The EID go function for: c_code @ Select (throw: true) 
func E_c_code_Select (self EID) EID { 
    return F_c_code_Select(Language.To_Select(OBJ(self)) )} 
  
/* {1} The go function for: c_code(self:Lselect) [status=1] */
func F_c_code_Lselect (self *Language.Lselect ) EID { 
    var Result EID 
    Result = F_Optimize_c_code_select_Iteration(Language.To_Iteration(self.Id()),C_list)
    return Result} 
  
// The EID go function for: c_code @ Lselect (throw: true) 
func E_c_code_Lselect (self EID) EID { 
    return F_c_code_Lselect(Language.To_Lselect(OBJ(self)) )} 
  
// changed in CLAIRE 4 (cf trans -> init.cl)
// x is set or list, tells what we want as output
/* {1} The go function for: c_code_select(self:Iteration,x:class) [status=1] */
func F_Optimize_c_code_select_Iteration (self *Language.Iteration ,x *ClaireClass ) EID { 
    var Result EID 
    { var sx *ClaireAny   = self.SetArg
      { var _Zt *ClaireType  
        var try_1 EID 
        /*g_try(v2:"try_1",loop:false) */
        try_1 = Core.F_CALL(C_c_type,ARGS(sx.ToEID()))
        /* ERROR PROTECTION INSERTED (_Zt-Result) */
        if ErrorIn(try_1) {Result = try_1
        } else {
        _Zt = ToType(OBJ(try_1))
        { var st *ClaireAny  
          _ = st
          var try_2 EID 
          /*g_try(v2:"try_2",loop:false) */
          try_2 = F_Optimize_enumerate_code_any(sx,_Zt)
          /* ERROR PROTECTION INSERTED (st-Result) */
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
                _ = arg_3
                C_OPT.MaxVars = (C_OPT.MaxVars+1)
                /*integer->integer*/arg_3 = 0
                v1 = F_Compile_Variable_I_symbol(F_append_symbol(self.ClaireVar.Pname,MakeString("_in").Id()),arg_3,_Zt.Id())
                } 
              { var v2 *ClaireVariable  
                { var arg_4 int 
                  _ = arg_4
                  C_OPT.MaxVars = (C_OPT.MaxVars+1)
                  /*integer->integer*/arg_4 = 0
                  v2 = F_Compile_Variable_I_symbol(F_append_symbol(self.ClaireVar.Pname,MakeString("_out").Id()),arg_4,x.Id())
                  } 
                if (Core.F_get_property(C_of,ToObject(self.Id())) != CNULL) { 
                  { var _ZtypeIn *ClaireType   = F_Optimize_pmember_type(_Zt)
                    if ((F_Optimize_ptype_type(_ZtypeIn).Included(ToType(OBJ(Core.F_CALL(C_of,ARGS(EID{self.Id(),0}))))) != CTRUE) && 
                        (C_compiler.Safety <= 1)) { 
                      F_Compile_warn_void()
                      Core.F_tformat_string(MakeString("unsafe bag construction (~S) : a ~S is not a ~S [262]\n"),1,MakeConstantList(self.ClaireVar.Id(),_ZtypeIn.Id(),ANY(Core.F_CALL(C_of,ARGS(EID{self.Id(),0})))))
                      } 
                    val.Cast_I(ToType(OBJ(Core.F_CALL(C_of,ARGS(EID{self.Id(),0})))))
                    v2.Range = Core.F_param_I_class(x,ToType(OBJ(Core.F_CALL(C_of,ARGS(EID{self.Id(),0})))))
                    /*type->type*/Result = F_Optimize_inner_select_Iteration(self,v2.Id(),sx,val.Id())
                    } 
                  }  else if (_Zt.Included(ToType(x.Id())) == CTRUE) { 
                  { var _CL_obj *Language.Let   = Language.To_Let(new(Language.Let).Is(Language.C_Let))
                    _CL_obj.ClaireVar = v1
                    /*Variable->Variable*/_CL_obj.Value = st
                    /*any->any*//*g_try(v2:"Result",loop:true) */
                    { 
                      var va_arg1 *Language.Let  
                      var va_arg2 *ClaireAny  
                      va_arg1 = _CL_obj
                      var try_5 EID 
                      /*g_try(v2:"try_5",loop:false) */
                      { var arg_6 *Compile_CCast  
                        _ = arg_6
                        var try_7 EID 
                        /*g_try(v2:"try_7",loop:false) */
                        { var _CL_obj *Compile_CCast   = To_Compile_CCast(new(Compile_CCast).Is(C_Compile_C_cast))
                          /*g_try(v2:"try_7",loop:false) */
                          { 
                            var va_arg1 *Compile_CCast  
                            var va_arg2 *ClaireAny  
                            va_arg1 = _CL_obj
                            var try_8 EID 
                            /*g_try(v2:"try_8",loop:false) */
                            { var arg_9 *Language.Call  
                              _ = arg_9
                              { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                _CL_obj.Selector = C_empty
                                /*property->property*/_CL_obj.Args = MakeConstantList(v1.Id())
                                /*list->list*/arg_9 = _CL_obj
                                } 
                              try_8 = Core.F_CALL(C_c_code,ARGS(EID{arg_9.Id(),0},EID{x.Id(),0}))
                              } 
                            /* ERROR PROTECTION INSERTED (va_arg2-try_7) */
                            if ErrorIn(try_8) {try_7 = try_8
                            } else {
                            va_arg2 = ANY(try_8)
                            va_arg1.Arg = va_arg2
                            /*any->any*/try_7 = va_arg2.ToEID()
                            }
                            } 
                          /* ERROR PROTECTION INSERTED (try_7-try_7) */
                          if !ErrorIn(try_7) {
                          _CL_obj.SetArg = x
                          /*class->class*/try_7 = EID{_CL_obj.Id(),0}
                          }
                          } 
                        /* ERROR PROTECTION INSERTED (arg_6-try_5) */
                        if ErrorIn(try_7) {try_5 = try_7
                        } else {
                        arg_6 = To_Compile_CCast(OBJ(try_7))
                        try_5 = F_Optimize_inner_select_Iteration(self,v2.Id(),v1.Id(),arg_6.Id())
                        }
                        } 
                      /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                      if ErrorIn(try_5) {Result = try_5
                      } else {
                      va_arg2 = ANY(try_5)
                      va_arg1.Arg = va_arg2
                      /*any->any*/Result = va_arg2.ToEID()
                      }
                      } 
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    Result = EID{_CL_obj.Id(),0}
                    }
                    } 
                  } else {
                  { var arg_10 *Language.Construct  
                    _ = arg_10
                    if (x.Id() == C_set.Id()) { 
                      { var _CL_obj *Language.Set   = Language.To_Set(new(Language.Set).Is(Language.C_Set))
                        _CL_obj.Of = ToType(CEMPTY.Id())
                        /*type->type*/arg_10 = Language.To_Construct(_CL_obj.Id())
                        } 
                      } else {
                      { var _CL_obj *Language.List   = Language.To_List(new(Language.List).Is(Language.C_List))
                        _CL_obj.Of = ToType(CEMPTY.Id())
                        /*type->type*/arg_10 = Language.To_Construct(_CL_obj.Id())
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
/* {1} The go function for: inner_select(self:Iteration,v2:any,sx:any,val:any) [status=1] */
func F_Optimize_inner_select_Iteration (self *Language.Iteration ,v2 *ClaireAny ,sx *ClaireAny ,val *ClaireAny ) EID { 
    var Result EID 
    { var _CL_obj *Language.Let   = Language.To_Let(new(Language.Let).Is(Language.C_Let))
      _CL_obj.ClaireVar = To_Variable(v2)
      /*Variable->Variable*/_CL_obj.Value = val
      /*any->any*//*g_try(v2:"Result",loop:true) */
      { 
        var va_arg1 *Language.Let  
        var va_arg2 *ClaireAny  
        va_arg1 = _CL_obj
        var try_1 EID 
        /*g_try(v2:"try_1",loop:false) */
        { var _CL_obj *Language.Do   = Language.To_Do(new(Language.Do).Is(Language.C_Do))
          /*g_try(v2:"try_1",loop:false) */
          { 
            var va_arg1 *Language.Do  
            var va_arg2 *ClaireList  
            va_arg1 = _CL_obj
            var try_2 EID 
            /*g_try(v2:"try_2",loop:false) */
            { 
              var v_bag_arg *ClaireAny  
              try_2= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
              var try_3 EID 
              /*g_try(v2:"try_3",loop:false) */
              { var arg_4 *Language.For  
                _ = arg_4
                { var _CL_obj *Language.For   = Language.To_For(new(Language.For).Is(Language.C_For))
                  _CL_obj.ClaireVar = self.ClaireVar
                  /*Variable->Variable*/_CL_obj.SetArg = sx
                  /*any->any*/{ 
                    var va_arg1 *Language.Iteration  
                    var va_arg2 *ClaireAny  
                    va_arg1 = Language.To_Iteration(_CL_obj.Id())
                    { var _CL_obj *Language.If   = Language.To_If(new(Language.If).Is(Language.C_If))
                      _CL_obj.Test = self.Arg
                      /*any->any*/{ 
                        var va_arg1 *Language.If  
                        var va_arg2 *ClaireAny  
                        va_arg1 = _CL_obj
                        { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                          _CL_obj.Selector = ToProperty(C_add_I.Id())
                          /*property->property*/_CL_obj.Args = MakeConstantList(v2,self.ClaireVar.Id())
                          /*list->list*/va_arg2 = _CL_obj.Id()
                          } 
                        va_arg1.Arg = va_arg2
                        /*any->any*/} 
                      va_arg2 = _CL_obj.Id()
                      } 
                    va_arg1.Arg = va_arg2
                    /*any->any*/} 
                  arg_4 = _CL_obj
                  } 
                try_3 = Core.F_CALL(C_c_code,ARGS(EID{arg_4.Id(),0},EID{C_any.Id(),0}))
                } 
              /* ERROR PROTECTION INSERTED (v_bag_arg-try_2) */
              if ErrorIn(try_3) {try_2 = try_3
              } else {
              v_bag_arg = ANY(try_3)
              ToList(OBJ(try_2)).AddFast(v_bag_arg)
              ToList(OBJ(try_2)).AddFast(v2)}
              } 
            /* ERROR PROTECTION INSERTED (va_arg2-try_1) */
            if ErrorIn(try_2) {try_1 = try_2
            } else {
            va_arg2 = ToList(OBJ(try_2))
            va_arg1.Args = va_arg2
            /*list->list*/try_1 = EID{va_arg2.Id(),0}
            }
            } 
          /* ERROR PROTECTION INSERTED (try_1-try_1) */
          if !ErrorIn(try_1) {
          try_1 = EID{_CL_obj.Id(),0}
          }
          } 
        /* ERROR PROTECTION INSERTED (va_arg2-Result) */
        if ErrorIn(try_1) {Result = try_1
        } else {
        va_arg2 = ANY(try_1)
        va_arg1.Arg = va_arg2
        /*any->any*/Result = va_arg2.ToEID()
        }
        } 
      /* ERROR PROTECTION INSERTED (Result-Result) */
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
/* {1} The go function for: c_type(self:Exists) [status=1] */
func F_c_type_Exists (self *Language.Exists ) EID { 
    var Result EID 
    { var _Zt *ClaireType  
      var try_1 EID 
      /*g_try(v2:"try_1",loop:false) */
      try_1 = Core.F_CALL(C_c_type,ARGS(self.SetArg.ToEID()))
      /* ERROR PROTECTION INSERTED (_Zt-Result) */
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
/* {1} The go function for: c_code(self:Exists,s:class) [status=1] */
func F_c_code_Exists (self *Language.Exists ,s *ClaireClass ) EID { 
    var Result EID 
    { var _Zt *ClaireType  
      var try_1 EID 
      /*g_try(v2:"try_1",loop:false) */
      try_1 = Core.F_CALL(C_c_type,ARGS(self.SetArg.ToEID()))
      /* ERROR PROTECTION INSERTED (_Zt-Result) */
      if ErrorIn(try_1) {Result = try_1
      } else {
      _Zt = ToType(OBJ(try_1))
      F_Optimize_range_infers_for_Variable(self.ClaireVar,F_Optimize_pmember_type(_Zt),_Zt)
      if (self.Other == CTRUE.Id()) { 
        { var arg_2 *Language.Call  
          _ = arg_2
          { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
            _CL_obj.Selector = Core.C_not
            /*property->property*/{ 
              var va_arg1 *Language.Call  
              var va_arg2 *ClaireList  
              va_arg1 = _CL_obj
              { 
                var v_bag_arg *ClaireAny  
                va_arg2= ToType(CEMPTY.Id()).EmptyList()
                { var _CL_obj *Language.For   = Language.To_For(new(Language.For).Is(Language.C_For))
                  _CL_obj.ClaireVar = self.ClaireVar
                  /*Variable->Variable*/_CL_obj.SetArg = self.SetArg
                  /*any->any*/{ 
                    var va_arg1 *Language.Iteration  
                    var va_arg2 *ClaireAny  
                    va_arg1 = Language.To_Iteration(_CL_obj.Id())
                    { var _CL_obj *Language.If   = Language.To_If(new(Language.If).Is(Language.C_If))
                      { 
                        var va_arg1 *Language.If  
                        var va_arg2 *ClaireAny  
                        va_arg1 = _CL_obj
                        { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                          _CL_obj.Selector = Core.C_not
                          /*property->property*/_CL_obj.Args = MakeConstantList(self.Arg)
                          /*list->list*/va_arg2 = _CL_obj.Id()
                          } 
                        va_arg1.Test = va_arg2
                        /*any->any*/} 
                      _CL_obj.Arg = Language.C_Return.Make(CTRUE.Id())
                      /*any->any*/va_arg2 = _CL_obj.Id()
                      } 
                    va_arg1.Arg = va_arg2
                    /*any->any*/} 
                  v_bag_arg = _CL_obj.Id()
                  } 
                va_arg2.AddFast(v_bag_arg)} 
              va_arg1.Args = va_arg2
              /*list->list*/} 
            arg_2 = _CL_obj
            } 
          Result = Core.F_CALL(C_c_code,ARGS(EID{arg_2.Id(),0},EID{s.Id(),0}))
          } 
        }  else if (self.Other == CNULL) { 
        { var v *ClaireVariable  
          { var arg_3 int 
            _ = arg_3
            C_OPT.MaxVars = (C_OPT.MaxVars+1)
            /*integer->integer*/arg_3 = 0
            v = F_Compile_Variable_I_symbol(F_append_symbol(self.ClaireVar.Pname,MakeString("_some").Id()),arg_3,F_Optimize_extends_type(self.ClaireVar.Range).Id())
            } 
          { var arg_4 *Language.Let  
            _ = arg_4
            { var _CL_obj *Language.Let   = Language.To_Let(new(Language.Let).Is(Language.C_Let))
              _CL_obj.ClaireVar = v
              /*Variable->Variable*/_CL_obj.Value = CNULL
              /*any->any*/{ 
                var va_arg1 *Language.Let  
                var va_arg2 *ClaireAny  
                va_arg1 = _CL_obj
                { var _CL_obj *Language.Do   = Language.To_Do(new(Language.Do).Is(Language.C_Do))
                  { 
                    var va_arg1 *Language.Do  
                    var va_arg2 *ClaireList  
                    va_arg1 = _CL_obj
                    { 
                      var v_bag_arg *ClaireAny  
                      va_arg2= ToType(CEMPTY.Id()).EmptyList()
                      { var _CL_obj *Language.For   = Language.To_For(new(Language.For).Is(Language.C_For))
                        _CL_obj.ClaireVar = self.ClaireVar
                        /*Variable->Variable*/_CL_obj.SetArg = self.SetArg
                        /*any->any*/{ 
                          var va_arg1 *Language.Iteration  
                          var va_arg2 *ClaireAny  
                          va_arg1 = Language.To_Iteration(_CL_obj.Id())
                          { var _CL_obj *Language.If   = Language.To_If(new(Language.If).Is(Language.C_If))
                            _CL_obj.Test = self.Arg
                            /*any->any*/_CL_obj.Arg = Language.C_Return.Make(Language.C_Assign.Make(v.Id(),self.ClaireVar.Id()))
                            /*any->any*/va_arg2 = _CL_obj.Id()
                            } 
                          va_arg1.Arg = va_arg2
                          /*any->any*/} 
                        v_bag_arg = _CL_obj.Id()
                        } 
                      va_arg2.AddFast(v_bag_arg)
                      va_arg2.AddFast(v.Id())} 
                    va_arg1.Args = va_arg2
                    /*list->list*/} 
                  va_arg2 = _CL_obj.Id()
                  } 
                va_arg1.Arg = va_arg2
                /*any->any*/} 
              arg_4 = _CL_obj
              } 
            Result = Core.F_CALL(C_c_code,ARGS(EID{arg_4.Id(),0},EID{s.Id(),0}))
            } 
          } 
        } else {
        { var arg_5 *Language.Call  
          _ = arg_5
          { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
            _CL_obj.Selector = C_boolean_I
            /*property->property*/{ 
              var va_arg1 *Language.Call  
              var va_arg2 *ClaireList  
              va_arg1 = _CL_obj
              { 
                var v_bag_arg *ClaireAny  
                va_arg2= ToType(CEMPTY.Id()).EmptyList()
                { var _CL_obj *Language.For   = Language.To_For(new(Language.For).Is(Language.C_For))
                  _CL_obj.ClaireVar = self.ClaireVar
                  /*Variable->Variable*/_CL_obj.SetArg = self.SetArg
                  /*any->any*/{ 
                    var va_arg1 *Language.Iteration  
                    var va_arg2 *ClaireAny  
                    va_arg1 = Language.To_Iteration(_CL_obj.Id())
                    { var _CL_obj *Language.If   = Language.To_If(new(Language.If).Is(Language.C_If))
                      _CL_obj.Test = self.Arg
                      /*any->any*/_CL_obj.Arg = Language.C_Return.Make(CTRUE.Id())
                      /*any->any*/va_arg2 = _CL_obj.Id()
                      } 
                    va_arg1.Arg = va_arg2
                    /*any->any*/} 
                  v_bag_arg = _CL_obj.Id()
                  } 
                va_arg2.AddFast(v_bag_arg)} 
              va_arg1.Args = va_arg2
              /*list->list*/} 
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
/* {1} The go function for: c_type(self:Image) [status=1] */
func F_c_type_Image (self *Language.Image ) EID { 
    var Result EID 
    { var _Zt *ClaireType  
      var try_1 EID 
      /*g_try(v2:"try_1",loop:false) */
      { var arg_2 *ClaireType  
        _ = arg_2
        var try_3 EID 
        /*g_try(v2:"try_3",loop:false) */
        try_3 = Core.F_CALL(C_c_type,ARGS(self.SetArg.ToEID()))
        /* ERROR PROTECTION INSERTED (arg_2-try_1) */
        if ErrorIn(try_3) {try_1 = try_3
        } else {
        arg_2 = ToType(OBJ(try_3))
        try_1 = EID{F_Optimize_ptype_type(arg_2).Id(),0}
        }
        } 
      /* ERROR PROTECTION INSERTED (_Zt-Result) */
      if ErrorIn(try_1) {Result = try_1
      } else {
      _Zt = ToType(OBJ(try_1))
      F_Optimize_range_infers_for_Variable(self.ClaireVar,F_Optimize_pmember_type(_Zt),_Zt)
      if (self.Of.Id() != CNULL) { 
        Result = EID{Core.F_param_I_class(C_set,self.Of).Id(),0}
        } else {
        { var arg_4 *ClaireType  
          _ = arg_4
          var try_5 EID 
          /*g_try(v2:"try_5",loop:false) */
          try_5 = Core.F_CALL(C_c_type,ARGS(self.Arg.ToEID()))
          /* ERROR PROTECTION INSERTED (arg_4-Result) */
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
  
/* {1} The go function for: c_type(self:Select) [status=1] */
func F_c_type_Select (self *Language.Select ) EID { 
    var Result EID 
    { var _Zt *ClaireType  
      var try_1 EID 
      /*g_try(v2:"try_1",loop:false) */
      try_1 = Core.F_CALL(C_c_type,ARGS(self.SetArg.ToEID()))
      /* ERROR PROTECTION INSERTED (_Zt-Result) */
      if ErrorIn(try_1) {Result = try_1
      } else {
      _Zt = ToType(OBJ(try_1))
      F_Optimize_range_infers_for_Variable(self.ClaireVar,F_Optimize_pmember_type(_Zt),_Zt)
      if (self.Of.Id() != CNULL) { 
        Result = EID{Core.F_param_I_class(C_set,self.Of).Id(),0}
        } else {
        { var arg_2 *ClaireType  
          _ = arg_2
          var try_3 EID 
          /*g_try(v2:"try_3",loop:false) */
          { var arg_4 *ClaireType  
            _ = arg_4
            var try_5 EID 
            /*g_try(v2:"try_5",loop:false) */
            try_5 = Core.F_CALL(C_c_type,ARGS(self.SetArg.ToEID()))
            /* ERROR PROTECTION INSERTED (arg_4-try_3) */
            if ErrorIn(try_5) {try_3 = try_5
            } else {
            arg_4 = ToType(OBJ(try_5))
            try_3 = EID{F_Optimize_pmember_type(arg_4).Id(),0}
            }
            } 
          /* ERROR PROTECTION INSERTED (arg_2-Result) */
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
/* {1} The go function for: c_type(self:Lselect) [status=1] */
func F_c_type_Lselect (self *Language.Lselect ) EID { 
    var Result EID 
    { var _Zt *ClaireType  
      var try_1 EID 
      /*g_try(v2:"try_1",loop:false) */
      try_1 = Core.F_CALL(C_c_type,ARGS(self.SetArg.ToEID()))
      /* ERROR PROTECTION INSERTED (_Zt-Result) */
      if ErrorIn(try_1) {Result = try_1
      } else {
      _Zt = ToType(OBJ(try_1))
      F_Optimize_range_infers_for_Variable(self.ClaireVar,F_Optimize_pmember_type(_Zt),_Zt)
      if (self.Of.Id() != CNULL) { 
        Result = EID{Core.F_param_I_class(C_list,self.Of).Id(),0}
        } else {
        { var arg_2 *ClaireType  
          _ = arg_2
          var try_3 EID 
          /*g_try(v2:"try_3",loop:false) */
          { var arg_4 *ClaireType  
            _ = arg_4
            var try_5 EID 
            /*g_try(v2:"try_5",loop:false) */
            try_5 = Core.F_CALL(C_c_type,ARGS(self.SetArg.ToEID()))
            /* ERROR PROTECTION INSERTED (arg_4-try_3) */
            if ErrorIn(try_5) {try_3 = try_5
            } else {
            arg_4 = ToType(OBJ(try_5))
            try_3 = EID{F_Optimize_pmember_type(arg_4).Id(),0}
            }
            } 
          /* ERROR PROTECTION INSERTED (arg_2-Result) */
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
/* {1} The go function for: c_type(self:While) [status=1] */
func F_c_type_While (self *Language.While ) EID { 
    var Result EID 
    { var arg_1 *ClaireType  
      _ = arg_1
      var try_2 EID 
      /*g_try(v2:"try_2",loop:false) */
      try_2 = F_Compile_return_type_any(self.Arg)
      /* ERROR PROTECTION INSERTED (arg_1-Result) */
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
  
/* {1} The go function for: c_code(self:While,s:class) [status=1] */
func F_c_code_While (self *Language.While ,s *ClaireClass ) EID { 
    var Result EID 
    { var _CL_obj *Language.While   = Language.To_While(new(Language.While).Is(Language.C_While))
      /*g_try(v2:"Result",loop:true) */
      { 
        var va_arg1 *Language.While  
        var va_arg2 *ClaireAny  
        va_arg1 = _CL_obj
        var try_1 EID 
        /*g_try(v2:"try_1",loop:false) */
        try_1 = F_Optimize_c_boolean_any(self.Test)
        /* ERROR PROTECTION INSERTED (va_arg2-Result) */
        if ErrorIn(try_1) {Result = try_1
        } else {
        va_arg2 = ANY(try_1)
        va_arg1.Test = va_arg2
        /*any->any*/Result = va_arg2.ToEID()
        }
        } 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      /*g_try(v2:"Result",loop:true) */
      { 
        var va_arg1 *Language.While  
        var va_arg2 *ClaireAny  
        va_arg1 = _CL_obj
        var try_2 EID 
        /*g_try(v2:"try_2",loop:false) */
        try_2 = Core.F_CALL(C_c_code,ARGS(self.Arg.ToEID(),EID{C_void.Id(),0}))
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
      _CL_obj.Other = self.Other
      /*boolean->boolean*/Result = EID{_CL_obj.Id(),0}
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
/* {1} The go function for: Iterate!(self:Iteration) [status=0] */
func F_Optimize_Iterate_I_Iteration (self *Language.Iteration ) *ClaireAny  { 
    return  F_Optimize_restriction_I_property(Language.C_Iterate,MakeConstantList(MakeConstantSet(self.SetArg).Id(),MakeConstantSet(self.ClaireVar.Id()).Id(),C_any.Id()),CTRUE)
    } 
  
// The EID go function for: Iterate! @ Iteration (throw: false) 
func E_Optimize_Iterate_I_Iteration (self EID) EID { 
    return F_Optimize_Iterate_I_Iteration(Language.To_Iteration(OBJ(self)) ).ToEID()} 
  
// iteration methods
// note the beauty of this: we only apply the code transformation if
// we actually get a constant Interval
/* {1} The go function for: iterate(x:Interval,v:Variable[range:(subtype[integer])],e:any) [status=1] */
func F_iterate_Interval (x *ClaireInterval ,v *ClaireVariable ,e *ClaireAny ) EID { 
    var Result EID 
    { var v *ClaireAny  
      _ = v
      var try_1 EID 
      /*g_try(v2:"try_1",loop:false) */
      try_1 = F_eval_any2(MakeInteger(x.Arg1).Id(),C_Interval)
      /* ERROR PROTECTION INSERTED (v-Result) */
      if ErrorIn(try_1) {Result = try_1
      } else {
      v = ANY(try_1)
      { var _Zmax int 
        _ = _Zmax
        var try_2 EID 
        /*g_try(v2:"try_2",loop:false) */
        try_2 = F_eval_any2(MakeInteger(x.Arg2).Id(),C_Interval)
        /* ERROR PROTECTION INSERTED (_Zmax-Result) */
        if ErrorIn(try_2) {Result = try_2
        } else {
        _Zmax = INT(try_2)
        Result= EID{CFALSE.Id(),0}
        for (ToInteger(v).Value <= _Zmax) { 
          /* While stat, v:"Result" loop:true */
          
          v = ANY(Core.F_CALL(ToProperty(Core.C__plus.Id()),ARGS(v.ToEID(),EID{C__INT,IVAL(1)})))
          /* try?:false, v2:"v_while4" loop will be:tuple("Result", EID) */
          } 
        }
        } 
      }
      } 
    return Result} 
  
// The EID go function for: iterate @ Interval (throw: true) 
func E_iterate_Interval (x EID,v EID,e EID) EID { 
    return F_iterate_Interval(To_Interval(OBJ(x)),To_Variable(OBJ(v)),ANY(e) )} 
  
/* {1} The go function for: iterate(x:array,v:Variable,e:any) [status=1] */
func F_iterate_array (x *ClaireList ,v *ClaireVariable ,e *ClaireAny ) EID { 
    var Result EID 
    { var _Zi int  = 1
      { var _Za *ClaireList   = x
        { var _Zmax int  = _Za.Length()
          _ = _Zmax
          Result= EID{CFALSE.Id(),0}
          for (_Zi <= _Zmax) { 
            /* While stat, v:"Result" loop:true */
            var loop_1 EID 
            _ = loop_1
            /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
            { var v *ClaireAny  
              _ = v
              var try_2 EID 
              /*g_try(v2:"try_2",loop:false) */
              try_2 = Core.F_nth_array(_Za,_Zi)
              /* ERROR PROTECTION INSERTED (v-loop_1) */
              if ErrorIn(try_2) {loop_1 = try_2
              } else {
              v = ANY(try_2)
              
              _Zi = (_Zi+1)
              loop_1 = EID{C__INT,IVAL(_Zi)}
              }
              } 
            /* ERROR PROTECTION INSERTED (loop_1-Result) */
            if ErrorIn(loop_1) {Result = loop_1
            break
            } else {
            /* try?:false, v2:"v_while5" loop will be:tuple("Result", EID) */
            } 
          }
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: iterate @ array (throw: true) 
func E_iterate_array (x EID,v EID,e EID) EID { 
    return F_iterate_array(ToArray(OBJ(x)),To_Variable(OBJ(v)),ANY(e) )} 
  
/* {1} The go function for: Iterate(x:class,v:Variable,e:any) [status=0] */
func F_Iterate_class (x *ClaireClass ,v *ClaireVariable ,e *ClaireAny ) *ClaireAny  { 
    // procedure body with s = any 
var Result *ClaireAny  
    { 
      var _Zv_1 *ClaireClass  
      _ = _Zv_1
      var _Zv_1_iter *ClaireAny  
      Result= CFALSE.Id()
      var _Zv_1_support *ClaireSet  
      _Zv_1_support = x.Descendents
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
  
/* {1} The go function for: Iterate(x:any,v:Variable,e:any) [status=1] */
func F_Iterate_any1 (x *ClaireAny ,v *ClaireVariable ,e *ClaireAny ) EID { 
    var Result EID 
    { var v *ClaireAny  
      _ = v
      var try_1 EID 
      /*g_try(v2:"try_1",loop:false) */
      try_1 = EVAL(ToList(OBJ(Core.F_CALL(C_args,ARGS(x.ToEID())))).At(1-1))
      /* ERROR PROTECTION INSERTED (v-Result) */
      if ErrorIn(try_1) {Result = try_1
      } else {
      v = ANY(try_1)
      { var _Zmax *ClaireAny  
        _ = _Zmax
        var try_2 EID 
        /*g_try(v2:"try_2",loop:false) */
        try_2 = EVAL(ToList(OBJ(Core.F_CALL(C_args,ARGS(x.ToEID())))).At(2-1))
        /* ERROR PROTECTION INSERTED (_Zmax-Result) */
        if ErrorIn(try_2) {Result = try_2
        } else {
        _Zmax = ANY(try_2)
        Result= EID{CFALSE.Id(),0}
        for (ToBoolean(OBJ(Core.F_CALL(ToProperty(C__inf_equal.Id()),ARGS(v.ToEID(),_Zmax.ToEID())))) == CTRUE) { 
          /* While stat, v:"Result" loop:true */
          
          v = ANY(Core.F_CALL(ToProperty(Core.C__plus.Id()),ARGS(v.ToEID(),EID{C__INT,IVAL(1)})))
          /* try?:false, v2:"v_while4" loop will be:tuple("Result", EID) */
          } 
        }
        } 
      }
      } 
    return Result} 
  
// The EID go function for: Iterate @ list<type_expression>(..[tuple(integer,integer)], Variable, any) (throw: true) 
func E_Iterate_any1 (x EID,v EID,e EID) EID { 
    return F_Iterate_any1(ANY(x),To_Variable(OBJ(v)),ANY(e) )} 
  
/* {1} The go function for: Iterate(x:Lselect,v:Variable,e:any) [status=1] */
func F_Iterate_Lselect (x *Language.Lselect ,v *ClaireVariable ,e *ClaireAny ) EID { 
    var Result EID 
    { 
      var v *ClaireAny  
      _ = v
      Result= EID{CFALSE.Id(),0}
      var v_support *ClaireList  
      var try_1 EID 
      /*g_try(v2:"try_1",loop:false) */
      { var arg_2 *ClaireAny  
        _ = arg_2
        var try_3 EID 
        /*g_try(v2:"try_3",loop:false) */
        try_3 = EVAL(x.SetArg)
        /* ERROR PROTECTION INSERTED (arg_2-try_1) */
        if ErrorIn(try_3) {try_1 = try_3
        } else {
        arg_2 = ANY(try_3)
        try_1 = Core.F_enumerate_any(arg_2)
        }
        } 
      /* ERROR PROTECTION INSERTED (v_support-Result) */
      if ErrorIn(try_1) {Result = try_1
      } else {
      v_support = ToList(OBJ(try_1))
      v_len := v_support.Length()
      for i_it := 0; i_it < v_len; i_it++ { 
        v = v_support.At(i_it)
        var loop_4 EID 
        _ = loop_4
        /*g_try(v2:"loop_4",loop:tuple("Result", EID)) */
        var g0282I *ClaireBoolean  
        var try_5 EID 
        /*g_try(v2:"try_5",loop:false) */
        { var arg_6 *ClaireAny  
          _ = arg_6
          var try_7 EID 
          /*g_try(v2:"try_7",loop:false) */
          try_7 = EVAL(Language.F_substitution_any(x.Arg,x.ClaireVar,v))
          /* ERROR PROTECTION INSERTED (arg_6-try_5) */
          if ErrorIn(try_7) {try_5 = try_7
          } else {
          arg_6 = ANY(try_7)
          try_5 = EID{F_boolean_I_any(arg_6).Id(),0}
          }
          } 
        /* ERROR PROTECTION INSERTED (g0282I-loop_4) */
        if ErrorIn(try_5) {loop_4 = try_5
        } else {
        g0282I = ToBoolean(OBJ(try_5))
        if (g0282I == CTRUE) { 
          loop_4 = e.ToEID()
          } else {
          loop_4 = EID{CFALSE.Id(),0}
          } 
        }
        /* ERROR PROTECTION INSERTED (loop_4-Result) */
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
  
/* {1} The go function for: Iterate(x:Select,v:Variable,e:any) [status=1] */
func F_Iterate_Select (x *Language.Select ,v *ClaireVariable ,e *ClaireAny ) EID { 
    var Result EID 
    { 
      var v *ClaireAny  
      _ = v
      Result= EID{CFALSE.Id(),0}
      var v_support *ClaireList  
      var try_1 EID 
      /*g_try(v2:"try_1",loop:false) */
      { var arg_2 *ClaireAny  
        _ = arg_2
        var try_3 EID 
        /*g_try(v2:"try_3",loop:false) */
        try_3 = EVAL(x.SetArg)
        /* ERROR PROTECTION INSERTED (arg_2-try_1) */
        if ErrorIn(try_3) {try_1 = try_3
        } else {
        arg_2 = ANY(try_3)
        try_1 = Core.F_enumerate_any(arg_2)
        }
        } 
      /* ERROR PROTECTION INSERTED (v_support-Result) */
      if ErrorIn(try_1) {Result = try_1
      } else {
      v_support = ToList(OBJ(try_1))
      v_len := v_support.Length()
      for i_it := 0; i_it < v_len; i_it++ { 
        v = v_support.At(i_it)
        var loop_4 EID 
        _ = loop_4
        /*g_try(v2:"loop_4",loop:tuple("Result", EID)) */
        var g0283I *ClaireBoolean  
        var try_5 EID 
        /*g_try(v2:"try_5",loop:false) */
        { var arg_6 *ClaireAny  
          _ = arg_6
          var try_7 EID 
          /*g_try(v2:"try_7",loop:false) */
          try_7 = EVAL(Language.F_substitution_any(x.Arg,x.ClaireVar,v))
          /* ERROR PROTECTION INSERTED (arg_6-try_5) */
          if ErrorIn(try_7) {try_5 = try_7
          } else {
          arg_6 = ANY(try_7)
          try_5 = EID{F_boolean_I_any(arg_6).Id(),0}
          }
          } 
        /* ERROR PROTECTION INSERTED (g0283I-loop_4) */
        if ErrorIn(try_5) {loop_4 = try_5
        } else {
        g0283I = ToBoolean(OBJ(try_5))
        if (g0283I == CTRUE) { 
          loop_4 = e.ToEID()
          } else {
          loop_4 = EID{CFALSE.Id(),0}
          } 
        }
        /* ERROR PROTECTION INSERTED (loop_4-Result) */
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
  
/* {1} The go function for: Iterate(x:Collect,v:Variable,e:any) [status=1] */
func F_Iterate_Collect (x *Language.Collect ,v *ClaireVariable ,e *ClaireAny ) EID { 
    var Result EID 
    { 
      var C_Zv *ClaireAny  
      _ = C_Zv
      Result= EID{CFALSE.Id(),0}
      var C_Zv_support *ClaireList  
      var try_1 EID 
      /*g_try(v2:"try_1",loop:false) */
      { var arg_2 *ClaireAny  
        _ = arg_2
        var try_3 EID 
        /*g_try(v2:"try_3",loop:false) */
        try_3 = EVAL(x.SetArg)
        /* ERROR PROTECTION INSERTED (arg_2-try_1) */
        if ErrorIn(try_3) {try_1 = try_3
        } else {
        arg_2 = ANY(try_3)
        try_1 = Core.F_enumerate_any(arg_2)
        }
        } 
      /* ERROR PROTECTION INSERTED (C_Zv_support-Result) */
      if ErrorIn(try_1) {Result = try_1
      } else {
      C_Zv_support = ToList(OBJ(try_1))
      C_Zv_len := C_Zv_support.Length()
      for i_it := 0; i_it < C_Zv_len; i_it++ { 
        C_Zv = C_Zv_support.At(i_it)
        var loop_4 EID 
        _ = loop_4
        /*g_try(v2:"loop_4",loop:tuple("Result", EID)) */
        { var v *ClaireAny  
          _ = v
          var try_5 EID 
          /*g_try(v2:"try_5",loop:false) */
          try_5 = EVAL(Language.F_substitution_any(x.Arg,x.ClaireVar,C_Zv))
          /* ERROR PROTECTION INSERTED (v-loop_4) */
          if ErrorIn(try_5) {loop_4 = try_5
          } else {
          v = ANY(try_5)
          loop_4 = e.ToEID()
          }
          } 
        /* ERROR PROTECTION INSERTED (loop_4-Result) */
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
  
/* {1} The go function for: Iterate(x:any,v:Variable,e:any) [status=1] */
func F_Iterate_any2 (x *ClaireAny ,v *ClaireVariable ,e *ClaireAny ) EID { 
    var Result EID 
    { 
      var v *ClaireAny  
      _ = v
      Result= EID{CFALSE.Id(),0}
      var v_support *ClaireList  
      var try_1 EID 
      /*g_try(v2:"try_1",loop:false) */
      { var arg_2 *ClaireAny  
        _ = arg_2
        var try_3 EID 
        /*g_try(v2:"try_3",loop:false) */
        try_3 = EVAL(ToList(OBJ(Core.F_CALL(C_args,ARGS(x.ToEID())))).At(1-1))
        /* ERROR PROTECTION INSERTED (arg_2-try_1) */
        if ErrorIn(try_3) {try_1 = try_3
        } else {
        arg_2 = ANY(try_3)
        try_1 = Core.F_enumerate_any(arg_2)
        }
        } 
      /* ERROR PROTECTION INSERTED (v_support-Result) */
      if ErrorIn(try_1) {Result = try_1
      } else {
      v_support = ToList(OBJ(try_1))
      v_len := v_support.Length()
      for i_it := 0; i_it < v_len; i_it++ { 
        v = v_support.At(i_it)
        var loop_4 EID 
        _ = loop_4
        /*g_try(v2:"loop_4",loop:tuple("Result", EID)) */
        var g0284I *ClaireBoolean  
        var try_5 EID 
        /*g_try(v2:"try_5",loop:false) */
        { var arg_6 *ClaireAny  
          _ = arg_6
          var try_7 EID 
          /*g_try(v2:"try_7",loop:false) */
          try_7 = EVAL(ToList(OBJ(Core.F_CALL(C_args,ARGS(x.ToEID())))).At(2-1))
          /* ERROR PROTECTION INSERTED (arg_6-try_5) */
          if ErrorIn(try_7) {try_5 = try_7
          } else {
          arg_6 = ANY(try_7)
          try_5 = EID{Core.F__I_equal_any(v,arg_6).Id(),0}
          }
          } 
        /* ERROR PROTECTION INSERTED (g0284I-loop_4) */
        if ErrorIn(try_5) {loop_4 = try_5
        } else {
        g0284I = ToBoolean(OBJ(try_5))
        if (g0284I == CTRUE) { 
          loop_4 = e.ToEID()
          } else {
          loop_4 = EID{CFALSE.Id(),0}
          } 
        }
        /* ERROR PROTECTION INSERTED (loop_4-Result) */
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
  
/* {1} The go function for: Iterate(x:any,v:Variable,e:any) [status=1] */
func F_Iterate_any3 (x *ClaireAny ,v *ClaireVariable ,e *ClaireAny ) EID { 
    var Result EID 
    /*g_try(v2:"Result",loop:true) */
    { 
      var v *ClaireAny  
      _ = v
      Result= EID{CFALSE.Id(),0}
      var v_support *ClaireList  
      var try_1 EID 
      /*g_try(v2:"try_1",loop:false) */
      { var arg_2 *ClaireAny  
        _ = arg_2
        var try_3 EID 
        /*g_try(v2:"try_3",loop:false) */
        try_3 = EVAL(ToList(OBJ(Core.F_CALL(C_args,ARGS(x.ToEID())))).At(1-1))
        /* ERROR PROTECTION INSERTED (arg_2-try_1) */
        if ErrorIn(try_3) {try_1 = try_3
        } else {
        arg_2 = ANY(try_3)
        try_1 = Core.F_enumerate_any(arg_2)
        }
        } 
      /* ERROR PROTECTION INSERTED (v_support-Result) */
      if ErrorIn(try_1) {Result = try_1
      } else {
      v_support = ToList(OBJ(try_1))
      v_len := v_support.Length()
      for i_it := 0; i_it < v_len; i_it++ { 
        v = v_support.At(i_it)
        
        }
        } 
      } 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    { 
      var v *ClaireAny  
      _ = v
      Result= EID{CFALSE.Id(),0}
      var v_support *ClaireList  
      var try_4 EID 
      /*g_try(v2:"try_4",loop:false) */
      { var arg_5 *ClaireAny  
        _ = arg_5
        var try_6 EID 
        /*g_try(v2:"try_6",loop:false) */
        try_6 = EVAL(ToList(OBJ(Core.F_CALL(C_args,ARGS(x.ToEID())))).At(2-1))
        /* ERROR PROTECTION INSERTED (arg_5-try_4) */
        if ErrorIn(try_6) {try_4 = try_6
        } else {
        arg_5 = ANY(try_6)
        try_4 = Core.F_enumerate_any(arg_5)
        }
        } 
      /* ERROR PROTECTION INSERTED (v_support-Result) */
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
  