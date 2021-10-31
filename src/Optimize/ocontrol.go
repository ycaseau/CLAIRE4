/***** CLAIRE Compilation of file /Users/ycaseau/Dropbox/src/clairev4.0/src/compile/ocontrol.cl 
         [version 4.0.02 / safety 5] Saturday 10-30-2021 *****/

package Optimize
import (_ "fmt"
	. "Kernel"
	"Core"
	"Language"
	"Reader"
)

//-------- dumb function to prevent import errors --------
func import_g0706() { 
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
/* {1} OPT.The go function for: c_type(self:Assign) [] */
func F_c_type_Assign (self *Language.Assign ) EID { 
    var Result EID 
    Result = Core.F_CALL(C_c_type,ARGS(self.Arg.ToEID()))
    return RangeCheck(ToType(C_type.Id()),Result)} 
  
// The EID go function for: c_type @ Assign (throw: true) 
func E_c_type_Assign (self EID) EID { 
    return /*(sm for c_type @ Assign= EID)*/ F_c_type_Assign(Language.To_Assign(OBJ(self)) )} 
  
// we must include the type checking if needed
/* {1} OPT.The go function for: c_code(self:Assign) [] */
func F_c_code_Assign (self *Language.Assign ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var v *ClaireAny   = self.ClaireVar
      /* noccur = 5 */
      /* Let:3 */{ 
        var x *ClaireAny   = self.Arg
        /* noccur = 4 */
        /* Let:4 */{ 
          var _Ztype *ClaireType  
          /* noccur = 2 */
          var _Ztype_try07075 EID 
          /* Let:5 */{ 
            var g0708UU *ClaireType  
            /* noccur = 1 */
            var g0708UU_try07096 EID 
            g0708UU_try07096 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
            /* ERROR PROTECTION INSERTED (g0708UU-_Ztype_try07075) */
            if ErrorIn(g0708UU_try07096) {_Ztype_try07075 = g0708UU_try07096
            } else {
            g0708UU = ToType(OBJ(g0708UU_try07096))
            _Ztype_try07075 = EID{F_Optimize_ptype_type(g0708UU).Id(),0}
            }
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (_Ztype-Result) */
          if ErrorIn(_Ztype_try07075) {Result = _Ztype_try07075
          } else {
          _Ztype = ToType(OBJ(_Ztype_try07075))
          if (v.Isa.IsIn(C_Variable) != CTRUE) /* If:5 */{ 
            Result = Core.F_CALL(C_Compile_Cerror,ARGS(EID{MakeString("[213] ~S is not a variable").Id(),0},v.ToEID()))
            } else {
            Result = EID{CFALSE.Id(),0}
            /* If-5 */} 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          if (_Ztype.Included(ToType(OBJ(Core.F_CALL(C_range,ARGS(v.ToEID()))))) != CTRUE) /* If:5 */{ 
            var x_try07106 EID 
            x_try07106 = Core.F_CALL(C_Optimize_c_warn,ARGS(self.ClaireVar.ToEID(),x.ToEID(),EID{_Ztype.Id(),0}))
            /* ERROR PROTECTION INSERTED (x-Result) */
            if ErrorIn(x_try07106) {Result = x_try07106
            } else {
            x = ANY(x_try07106)
            Result = x.ToEID()
            }
            } else {
            Result = EID{CFALSE.Id(),0}
            /* If-5 */} 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          /* Let:5 */{ 
            var _Zarg *ClaireAny  
            /* noccur = 1 */
            var _Zarg_try07116 EID 
            _Zarg_try07116 = F_Compile_c_strict_code_any(x,F_Compile_psort_any(ANY(Core.F_CALL(C_range,ARGS(v.ToEID())))))
            /* ERROR PROTECTION INSERTED (_Zarg-Result) */
            if ErrorIn(_Zarg_try07116) {Result = _Zarg_try07116
            } else {
            _Zarg = ANY(_Zarg_try07116)
            Result = Language.C_Assign.Make(v,_Zarg).ToEID()
            }
            /* Let-5 */} 
          }}
          }
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_code @ Assign (throw: true) 
func E_c_code_Assign (self EID) EID { 
    return /*(sm for c_code @ Assign= EID)*/ F_c_code_Assign(Language.To_Assign(OBJ(self)) )} 
  
// assignment to a global variable
/* {1} OPT.The go function for: c_type(self:Gassign) [] */
func F_c_type_Gassign (self *Language.Gassign ) EID { 
    var Result EID 
    Result = Core.F_CALL(C_c_type,ARGS(self.Arg.ToEID()))
    return RangeCheck(ToType(C_type.Id()),Result)} 
  
// The EID go function for: c_type @ Gassign (throw: true) 
func E_c_type_Gassign (self EID) EID { 
    return /*(sm for c_type @ Gassign= EID)*/ F_c_type_Gassign(Language.To_Gassign(OBJ(self)) )} 
  
/* {1} OPT.The go function for: c_code(self:Gassign) [] */
func F_c_code_Gassign (self *Language.Gassign ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var _Zv *ClaireAny   = self.Arg
      /* noccur = 5 */
      /* Let:3 */{ 
        var _Ztype *ClaireType  
        /* noccur = 1 */
        var _Ztype_try07124 EID 
        /* Let:4 */{ 
          var g0713UU *ClaireType  
          /* noccur = 1 */
          var g0713UU_try07145 EID 
          g0713UU_try07145 = Core.F_CALL(C_c_type,ARGS(_Zv.ToEID()))
          /* ERROR PROTECTION INSERTED (g0713UU-_Ztype_try07124) */
          if ErrorIn(g0713UU_try07145) {_Ztype_try07124 = g0713UU_try07145
          } else {
          g0713UU = ToType(OBJ(g0713UU_try07145))
          _Ztype_try07124 = EID{F_Optimize_ptype_type(g0713UU).Id(),0}
          }
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (_Ztype-Result) */
        if ErrorIn(_Ztype_try07124) {Result = _Ztype_try07124
        } else {
        _Ztype = ToType(OBJ(_Ztype_try07124))
        if (F_boolean_I_any(self.ClaireVar.Range.Id()).Id() != CTRUE.Id()) /* If:4 */{ 
          Result = Core.F_CALL(C_Compile_Cerror,ARGS(EID{MakeString("[214] cannot assign ~S").Id(),0},EID{self.Id(),0}))
          } else {
          Result = EID{CFALSE.Id(),0}
          /* If-4 */} 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        if (_Ztype.Included(self.ClaireVar.Range) != CTRUE) /* If:4 */{ 
          var _Zv_try07155 EID 
          _Zv_try07155 = Core.F_CALL(C_c_code,ARGS(_Zv.ToEID(),EID{C_any.Id(),0}))
          /* ERROR PROTECTION INSERTED (_Zv-Result) */
          if ErrorIn(_Zv_try07155) {Result = _Zv_try07155
          } else {
          _Zv = ANY(_Zv_try07155)
          Result = _Zv.ToEID()
          }
          } else {
          Result = EID{CFALSE.Id(),0}
          /* If-4 */} 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        /* Let:4 */{ 
          var _CL_obj *Language.Gassign   = Language.To_Gassign(new(Language.Gassign).Is(Language.C_Gassign))
          /* noccur = 4 */
          /* update:5 */{ 
            var va_arg1 *Language.Gassign  
            var va_arg2 *Core.GlobalVariable  
            va_arg1 = _CL_obj
            var va_arg2_try07166 EID 
            va_arg2_try07166 = Core.F_CALL(C_c_code,ARGS(EID{self.ClaireVar.Id(),0}))
            /* ERROR PROTECTION INSERTED (va_arg2-Result) */
            if ErrorIn(va_arg2_try07166) {Result = va_arg2_try07166
            } else {
            va_arg2 = Core.ToGlobalVariable(OBJ(va_arg2_try07166))
            /* ---------- now we compile update var(va_arg1) := va_arg2 ------- */
            va_arg1.ClaireVar = va_arg2
            Result = EID{va_arg2.Id(),0}
            }
            /* update-5 */} 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          /* update:5 */{ 
            var va_arg1 *Language.Gassign  
            var va_arg2 *ClaireAny  
            va_arg1 = _CL_obj
            var va_arg2_try07176 EID 
            if (F_Compile_nativeVar_ask_global_variable(self.ClaireVar) == CTRUE) /* If:6 */{ 
              va_arg2_try07176 = F_Compile_c_strict_code_any(_Zv,F_Compile_psort_any(self.ClaireVar.Range.Id()))
              } else {
              va_arg2_try07176 = Core.F_CALL(C_c_code,ARGS(_Zv.ToEID(),EID{C_any.Id(),0}))
              /* If-6 */} 
            /* ERROR PROTECTION INSERTED (va_arg2-Result) */
            if ErrorIn(va_arg2_try07176) {Result = va_arg2_try07176
            } else {
            va_arg2 = ANY(va_arg2_try07176)
            /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
            va_arg1.Arg = va_arg2
            Result = va_arg2.ToEID()
            }
            /* update-5 */} 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          Result = EID{_CL_obj.Id(),0}
          }}
          /* Let-4 */} 
        }}
        }
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_code @ Gassign (throw: true) 
func E_c_code_Gassign (self EID) EID { 
    return /*(sm for c_code @ Gassign= EID)*/ F_c_code_Gassign(Language.To_Gassign(OBJ(self)) )} 
  
// v3.3 !
// _______________ l AND/OR     ____________________________________
/* {1} OPT.The go function for: c_type(self:And) [] */
func F_c_type_And (self *Language.And ) *ClaireType  { 
    // use function body compiling 
return  ToType(C_boolean.Id())
    } 
  
// The EID go function for: c_type @ And (throw: false) 
func E_c_type_And (self EID) EID { 
    return EID{/*(sm for c_type @ And= type)*/ F_c_type_And(Language.To_And(OBJ(self)) ).Id(),0}} 
  
/* {1} OPT.The go function for: c_code(self:And) [] */
func F_c_code_And (self *Language.And ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var _CL_obj *Language.And   = Language.To_And(new(Language.And).Is(Language.C_And))
      /* noccur = 3 */
      /* update:3 */{ 
        var va_arg1 *Language.And  
        var va_arg2 *ClaireList  
        va_arg1 = _CL_obj
        var va_arg2_try07184 EID 
        /* Iteration:4 */{ 
          var v_list4 *ClaireList  
          var x *ClaireAny  
          var v_local4 *ClaireAny  
          v_list4 = self.Args
          va_arg2_try07184 = EID{CreateList(ToType(CEMPTY.Id()),v_list4.Length()).Id(),0}
          for CLcount := 0; CLcount < v_list4.Length(); CLcount++{ 
            x = v_list4.At(CLcount)
            var v_local4_try07196 EID 
            var g0720I *ClaireBoolean  
            var g0720I_try07216 EID 
            /* Let:6 */{ 
              var g0722UU *ClaireType  
              /* noccur = 1 */
              var g0722UU_try07237 EID 
              g0722UU_try07237 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
              /* ERROR PROTECTION INSERTED (g0722UU-g0720I_try07216) */
              if ErrorIn(g0722UU_try07237) {g0720I_try07216 = g0722UU_try07237
              } else {
              g0722UU = ToType(OBJ(g0722UU_try07237))
              g0720I_try07216 = EID{Equal(g0722UU.Id(),C_void.Id()).Id(),0}
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (g0720I-v_local4_try07196) */
            if ErrorIn(g0720I_try07216) {v_local4_try07196 = g0720I_try07216
            } else {
            g0720I = ToBoolean(OBJ(g0720I_try07216))
            if (g0720I == CTRUE) /* If:6 */{ 
              v_local4_try07196 = Core.F_CALL(C_Compile_Cerror,ARGS(EID{MakeString("[206] void ~S in ~S").Id(),0},x.ToEID(),EID{self.Id(),0}))
              } else {
              v_local4_try07196 = EID{CFALSE.Id(),0}
              /* If-6 */} 
            }
            /* ERROR PROTECTION INSERTED (v_local4_try07196-v_local4_try07196) */
            if ErrorIn(v_local4_try07196) {va_arg2_try07184 = v_local4_try07196
            break
            } else {
            v_local4_try07196 = F_Optimize_c_boolean_any(x)
            /* ERROR PROTECTION INSERTED (v_local4_try07196-v_local4_try07196) */
            if ErrorIn(v_local4_try07196) {va_arg2_try07184 = v_local4_try07196
            break
            } else {
            }}
            {
            v_local4 = ANY(v_local4_try07196)
            ToList(OBJ(va_arg2_try07184)).PutAt(CLcount,v_local4)
            } 
          }
          /* Iteration-4 */} 
        /* ERROR PROTECTION INSERTED (va_arg2-Result) */
        if ErrorIn(va_arg2_try07184) {Result = va_arg2_try07184
        } else {
        va_arg2 = ToList(OBJ(va_arg2_try07184))
        /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
        va_arg1.Args = va_arg2
        Result = EID{va_arg2.Id(),0}
        }
        /* update-3 */} 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = EID{_CL_obj.Id(),0}
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_code @ And (throw: true) 
func E_c_code_And (self EID) EID { 
    return /*(sm for c_code @ And= EID)*/ F_c_code_And(Language.To_And(OBJ(self)) )} 
  
/* {1} OPT.The go function for: c_type(self:Or) [] */
func F_c_type_Or (self *Language.Or ) *ClaireType  { 
    // use function body compiling 
return  ToType(C_boolean.Id())
    } 
  
// The EID go function for: c_type @ Or (throw: false) 
func E_c_type_Or (self EID) EID { 
    return EID{/*(sm for c_type @ Or= type)*/ F_c_type_Or(Language.To_Or(OBJ(self)) ).Id(),0}} 
  
/* {1} OPT.The go function for: c_code(self:Or) [] */
func F_c_code_Or (self *Language.Or ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var _CL_obj *Language.Or   = Language.To_Or(new(Language.Or).Is(Language.C_Or))
      /* noccur = 3 */
      /* update:3 */{ 
        var va_arg1 *Language.Or  
        var va_arg2 *ClaireList  
        va_arg1 = _CL_obj
        var va_arg2_try07244 EID 
        /* Iteration:4 */{ 
          var v_list4 *ClaireList  
          var x *ClaireAny  
          var v_local4 *ClaireAny  
          v_list4 = self.Args
          va_arg2_try07244 = EID{CreateList(ToType(CEMPTY.Id()),v_list4.Length()).Id(),0}
          for CLcount := 0; CLcount < v_list4.Length(); CLcount++{ 
            x = v_list4.At(CLcount)
            var v_local4_try07256 EID 
            var g0726I *ClaireBoolean  
            var g0726I_try07276 EID 
            /* Let:6 */{ 
              var g0728UU *ClaireType  
              /* noccur = 1 */
              var g0728UU_try07297 EID 
              g0728UU_try07297 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
              /* ERROR PROTECTION INSERTED (g0728UU-g0726I_try07276) */
              if ErrorIn(g0728UU_try07297) {g0726I_try07276 = g0728UU_try07297
              } else {
              g0728UU = ToType(OBJ(g0728UU_try07297))
              g0726I_try07276 = EID{Equal(g0728UU.Id(),C_void.Id()).Id(),0}
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (g0726I-v_local4_try07256) */
            if ErrorIn(g0726I_try07276) {v_local4_try07256 = g0726I_try07276
            } else {
            g0726I = ToBoolean(OBJ(g0726I_try07276))
            if (g0726I == CTRUE) /* If:6 */{ 
              v_local4_try07256 = Core.F_CALL(C_Compile_Cerror,ARGS(EID{MakeString("[206] void ~S in ~S").Id(),0},x.ToEID(),EID{self.Id(),0}))
              } else {
              v_local4_try07256 = EID{CFALSE.Id(),0}
              /* If-6 */} 
            }
            /* ERROR PROTECTION INSERTED (v_local4_try07256-v_local4_try07256) */
            if ErrorIn(v_local4_try07256) {va_arg2_try07244 = v_local4_try07256
            break
            } else {
            v_local4_try07256 = F_Optimize_c_boolean_any(x)
            /* ERROR PROTECTION INSERTED (v_local4_try07256-v_local4_try07256) */
            if ErrorIn(v_local4_try07256) {va_arg2_try07244 = v_local4_try07256
            break
            } else {
            }}
            {
            v_local4 = ANY(v_local4_try07256)
            ToList(OBJ(va_arg2_try07244)).PutAt(CLcount,v_local4)
            } 
          }
          /* Iteration-4 */} 
        /* ERROR PROTECTION INSERTED (va_arg2-Result) */
        if ErrorIn(va_arg2_try07244) {Result = va_arg2_try07244
        } else {
        va_arg2 = ToList(OBJ(va_arg2_try07244))
        /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
        va_arg1.Args = va_arg2
        Result = EID{va_arg2.Id(),0}
        }
        /* update-3 */} 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = EID{_CL_obj.Id(),0}
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_code @ Or (throw: true) 
func E_c_code_Or (self EID) EID { 
    return /*(sm for c_code @ Or= EID)*/ F_c_code_Or(Language.To_Or(OBJ(self)) )} 
  
//---------------- quote and return -------------------------------------
/* {1} OPT.The go function for: c_type(self:Quote) [] */
func F_c_type_Quote (self *Language.Quote ) *ClaireType  { 
    // use function body compiling 
return  ToType(self.Arg.Isa.Id())
    } 
  
// The EID go function for: c_type @ Quote (throw: false) 
func E_c_type_Quote (self EID) EID { 
    return EID{/*(sm for c_type @ Quote= type)*/ F_c_type_Quote(Language.To_Quote(OBJ(self)) ).Id(),0}} 
  
/* {1} OPT.The go function for: c_code(self:Quote) [] */
func F_c_code_Quote (self *Language.Quote ) EID { 
    var Result EID 
    Result = Core.F_CALL(C_Compile_Cerror,ARGS(EID{MakeString("[internal] optimization of quote not implemented yet! ~S").Id(),0},EID{self.Id(),0}))
    return Result} 
  
// The EID go function for: c_code @ Quote (throw: true) 
func E_c_code_Quote (self EID) EID { 
    return /*(sm for c_code @ Quote= EID)*/ F_c_code_Quote(Language.To_Quote(OBJ(self)) )} 
  
/* {1} OPT.The go function for: c_type(self:Return) [] */
func F_c_type_Return (self *Language.Return ) *ClaireType  { 
    // use function body compiling 
return  ToType(C_any.Id())
    } 
  
// The EID go function for: c_type @ Return (throw: false) 
func E_c_type_Return (self EID) EID { 
    return EID{/*(sm for c_type @ Return= type)*/ F_c_type_Return(Language.To_Return(OBJ(self)) ).Id(),0}} 
  
/* {1} OPT.The go function for: c_code(self:Return) [] */
func F_c_code_Return (self *Language.Return ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var g0730UU *ClaireAny  
      /* noccur = 1 */
      var g0730UU_try07313 EID 
      g0730UU_try07313 = Core.F_CALL(C_c_code,ARGS(self.Arg.ToEID(),EID{C_any.Id(),0}))
      /* ERROR PROTECTION INSERTED (g0730UU-Result) */
      if ErrorIn(g0730UU_try07313) {Result = g0730UU_try07313
      } else {
      g0730UU = ANY(g0730UU_try07313)
      Result = Language.C_Return.Make(g0730UU).ToEID()
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_code @ Return (throw: true) 
func E_c_code_Return (self EID) EID { 
    return /*(sm for c_code @ Return= EID)*/ F_c_code_Return(Language.To_Return(OBJ(self)) )} 
  
// optimisation of exception handlers
/* {1} OPT.The go function for: c_type(self:Handle) [] */
func F_c_type_Handle (self *Language.ClaireHandle ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var g0732UU *ClaireType  
      /* noccur = 1 */
      var g0732UU_try07343 EID 
      g0732UU_try07343 = Core.F_CALL(C_c_type,ARGS(self.Arg.ToEID()))
      /* ERROR PROTECTION INSERTED (g0732UU-Result) */
      if ErrorIn(g0732UU_try07343) {Result = g0732UU_try07343
      } else {
      g0732UU = ToType(OBJ(g0732UU_try07343))
      /* Let:3 */{ 
        var g0733UU *ClaireType  
        /* noccur = 1 */
        var g0733UU_try07354 EID 
        g0733UU_try07354 = Core.F_CALL(C_c_type,ARGS(self.Other.ToEID()))
        /* ERROR PROTECTION INSERTED (g0733UU-Result) */
        if ErrorIn(g0733UU_try07354) {Result = g0733UU_try07354
        } else {
        g0733UU = ToType(OBJ(g0733UU_try07354))
        Result = EID{Core.F_U_type(g0732UU,g0733UU).Id(),0}
        }
        /* Let-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_type @ Handle (throw: true) 
func E_c_type_Handle (self EID) EID { 
    return /*(sm for c_type @ Handle= EID)*/ F_c_type_Handle(Language.To_ClaireHandle(OBJ(self)) )} 
  
/* {1} OPT.The go function for: c_code(self:Handle,s:class) [] */
func F_c_code_Handle (self *Language.ClaireHandle ,s *ClaireClass ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var x *Language.ClaireHandle  
      /* noccur = 2 */
      var x_try07363 EID 
      /* Let:3 */{ 
        var g0737UU *ClaireAny  
        /* noccur = 1 */
        var g0737UU_try07394 EID 
        g0737UU_try07394 = Core.F_CALL(C_c_code,ARGS(self.Arg.ToEID(),EID{s.Id(),0}))
        /* ERROR PROTECTION INSERTED (g0737UU-x_try07363) */
        if ErrorIn(g0737UU_try07394) {x_try07363 = g0737UU_try07394
        } else {
        g0737UU = ANY(g0737UU_try07394)
        /* Let:4 */{ 
          var g0738UU *ClaireAny  
          /* noccur = 1 */
          var g0738UU_try07405 EID 
          g0738UU_try07405 = Core.F_CALL(C_c_code,ARGS(self.Other.ToEID(),EID{s.Id(),0}))
          /* ERROR PROTECTION INSERTED (g0738UU-x_try07363) */
          if ErrorIn(g0738UU_try07405) {x_try07363 = g0738UU_try07405
          } else {
          g0738UU = ANY(g0738UU_try07405)
          x_try07363 = Language.C_Handle.Make(C_any.Id(),g0737UU,g0738UU).ToEID()
          }
          /* Let-4 */} 
        }
        /* Let-3 */} 
      /* ERROR PROTECTION INSERTED (x-Result) */
      if ErrorIn(x_try07363) {Result = x_try07363
      } else {
      x = Language.To_ClaireHandle(OBJ(x_try07363))
      x.Test = self.Test
      Result = EID{x.Id(),0}
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_code @ Handle (throw: true) 
func E_c_code_Handle (self EID,s EID) EID { 
    return /*(sm for c_code @ Handle= EID)*/ F_c_code_Handle(Language.To_ClaireHandle(OBJ(self)),ToClass(OBJ(s)) )} 
  
// ****************************************************************
// *      Part 2: Specific structures                             *
// ****************************************************************
// __________________ CAST ________________________________________
// cast is now more subtle since we introduce coercion for list<t> types
/* {1} OPT.The go function for: c_type(self:Cast) [] */
func F_c_type_Cast (self *Language.Cast ) *ClaireType  { 
    // use function body compiling 
return  self.SetArg
    } 
  
// The EID go function for: c_type @ Cast (throw: false) 
func E_c_type_Cast (self EID) EID { 
    return EID{/*(sm for c_type @ Cast= type)*/ F_c_type_Cast(Language.To_Cast(OBJ(self)) ).Id(),0}} 
  
// insert dynamic types (check_in) when we see a claire cast
// CLAIRE 4 : when we decide to drop the cast (safety), we generate a C_cast
/* {1} OPT.The go function for: c_code(self:Cast) [] */
func F_c_code_Cast (self *Language.Cast ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var y *ClaireType   = self.SetArg
      /* noccur = 8 */
      /* Let:3 */{ 
        var ftype *ClaireClass   = F_Compile_psort_any(y.Id())
        /* noccur = 5 */
        var g0742I *ClaireBoolean  
        if (y.Isa.IsIn(C_Param) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var g0741 *ClaireParam   = To_Param(y.Id())
            /* noccur = 3 */
            g0742I = MakeBoolean(((g0741.Arg.Id() == C_list.Id()) || 
                (g0741.Arg.Id() == C_set.Id())) && (C_set.Id() == g0741.Args.At(1-1).Isa.Id()))
            /* Let-5 */} 
          } else {
          g0742I = CFALSE
          /* If-4 */} 
        if (g0742I == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var utype *ClaireAny  
            /* noccur = 2 */
            var utype_try07436 EID 
            utype_try07436 = Core.F_the_type(ToType(To_Param(y.Id()).Args.At(1-1)))
            /* ERROR PROTECTION INSERTED (utype-Result) */
            if ErrorIn(utype_try07436) {Result = utype_try07436
            } else {
            utype = ANY(utype_try07436)
            var g0744I *ClaireBoolean  
            var g0744I_try07456 EID 
            /* or:6 */{ 
              var v_or6 *ClaireBoolean  
              
              var v_or6_try07467 EID 
              /* Let:7 */{ 
                var g0747UU *ClaireType  
                /* noccur = 1 */
                var g0747UU_try07488 EID 
                /* Let:8 */{ 
                  var g0749UU *ClaireType  
                  /* noccur = 1 */
                  var g0749UU_try07509 EID 
                  g0749UU_try07509 = Core.F_CALL(C_c_type,ARGS(self.Arg.ToEID()))
                  /* ERROR PROTECTION INSERTED (g0749UU-g0747UU_try07488) */
                  if ErrorIn(g0749UU_try07509) {g0747UU_try07488 = g0749UU_try07509
                  } else {
                  g0749UU = ToType(OBJ(g0749UU_try07509))
                  g0747UU_try07488 = EID{g0749UU.At(C_of).Id(),0}
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (g0747UU-v_or6_try07467) */
                if ErrorIn(g0747UU_try07488) {v_or6_try07467 = g0747UU_try07488
                } else {
                g0747UU = ToType(OBJ(g0747UU_try07488))
                v_or6_try07467 = EID{Equal(g0747UU.Id(),utype).Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (v_or6-g0744I_try07456) */
              if ErrorIn(v_or6_try07467) {g0744I_try07456 = v_or6_try07467
              } else {
              v_or6 = ToBoolean(OBJ(v_or6_try07467))
              if (v_or6 == CTRUE) {g0744I_try07456 = EID{CTRUE.Id(),0}
              } else /* or:7 */{ 
                v_or6 = Core.F__sup_integer(C_compiler.Safety,4)
                if (v_or6 == CTRUE) {g0744I_try07456 = EID{CTRUE.Id(),0}
                } else /* or:8 */{ 
                  g0744I_try07456 = EID{CFALSE.Id(),0}/* org-8 */} 
                /* org-7 */} 
              }
              /* or-6 */} 
            /* ERROR PROTECTION INSERTED (g0744I-Result) */
            if ErrorIn(g0744I_try07456) {Result = g0744I_try07456
            } else {
            g0744I = ToBoolean(OBJ(g0744I_try07456))
            if (g0744I == CTRUE) /* If:6 */{ 
              Result = Core.F_CALL(C_c_code,ARGS(self.Arg.ToEID(),EID{ftype.Id(),0}))
              } else {
              /* Let:7 */{ 
                var g0751UU *Language.Call  
                /* noccur = 1 */
                /* Let:8 */{ 
                  var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                  /* noccur = 5 */
                  _CL_obj.Selector = Core.C_check_in
                  _CL_obj.Args = MakeConstantList(self.Arg,To_Param(y.Id()).Arg.Id(),utype)
                  g0751UU = _CL_obj
                  /* Let-8 */} 
                Result = Core.F_CALL(C_c_code,ARGS(EID{g0751UU.Id(),0},EID{ftype.Id(),0}))
                /* Let-7 */} 
              /* If-6 */} 
            }
            }
            /* Let-5 */} 
          } else {
          var g0752I *ClaireBoolean  
          var g0752I_try07535 EID 
          /* Let:5 */{ 
            var g0754UU *ClaireType  
            /* noccur = 1 */
            var g0754UU_try07556 EID 
            g0754UU_try07556 = Core.F_CALL(C_c_type,ARGS(self.Arg.ToEID()))
            /* ERROR PROTECTION INSERTED (g0754UU-g0752I_try07535) */
            if ErrorIn(g0754UU_try07556) {g0752I_try07535 = g0754UU_try07556
            } else {
            g0754UU = ToType(OBJ(g0754UU_try07556))
            g0752I_try07535 = EID{g0754UU.Included(y).Id(),0}
            }
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (g0752I-Result) */
          if ErrorIn(g0752I_try07535) {Result = g0752I_try07535
          } else {
          g0752I = ToBoolean(OBJ(g0752I_try07535))
          if (g0752I == CTRUE) /* If:5 */{ 
            Result = Core.F_CALL(C_c_code,ARGS(self.Arg.ToEID(),EID{ftype.Id(),0}))
            /* If!5 */}  else if (C_compiler.Safety > 1) /* If:5 */{ 
            /* Let:6 */{ 
              var _CL_obj *CompileCCast   = To_CompileCCast(new(CompileCCast).Is(C_Compile_C_cast))
              /* noccur = 4 */
              /* update:7 */{ 
                var va_arg1 *CompileCCast  
                var va_arg2 *ClaireAny  
                va_arg1 = _CL_obj
                var va_arg2_try07568 EID 
                va_arg2_try07568 = Core.F_CALL(C_c_code,ARGS(self.Arg.ToEID(),EID{ftype.Id(),0}))
                /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                if ErrorIn(va_arg2_try07568) {Result = va_arg2_try07568
                } else {
                va_arg2 = ANY(va_arg2_try07568)
                /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                va_arg1.Arg = va_arg2
                Result = va_arg2.ToEID()
                }
                /* update-7 */} 
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              _CL_obj.SetArg = ToClass(y.Id())
              Result = EID{_CL_obj.Id(),0}
              }
              /* Let-6 */} 
            } else {
            /* Let:6 */{ 
              var g0757UU *Language.Call  
              /* noccur = 1 */
              /* Let:7 */{ 
                var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                /* noccur = 5 */
                _CL_obj.Selector = Core.C_check_in
                _CL_obj.Args = MakeConstantList(self.Arg,y.Id())
                g0757UU = _CL_obj
                /* Let-7 */} 
              Result = Core.F_CALL(C_c_code,ARGS(EID{g0757UU.Id(),0},EID{ftype.Id(),0}))
              /* Let-6 */} 
            /* If-5 */} 
          }
          /* If-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_code @ Cast (throw: true) 
func E_c_code_Cast (self EID) EID { 
    return /*(sm for c_code @ Cast= EID)*/ F_c_code_Cast(Language.To_Cast(OBJ(self)) )} 
  
// _________________ SUPER _________________________________________
/* {1} OPT.The go function for: c_type(self:Super) [] */
func F_c_type_Super (self *Language.Super ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var _Ztype *ClaireList  
      /* noccur = 3 */
      var _Ztype_try07613 EID 
      /* Iteration:3 */{ 
        var v_list3 *ClaireList  
        var x *ClaireAny  
        var v_local3 *ClaireAny  
        v_list3 = self.Args
        _Ztype_try07613 = EID{CreateList(ToType(CEMPTY.Id()),v_list3.Length()).Id(),0}
        for CLcount := 0; CLcount < v_list3.Length(); CLcount++{ 
          x = v_list3.At(CLcount)
          var v_local3_try07625 EID 
          v_local3_try07625 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
          /* ERROR PROTECTION INSERTED (v_local3-_Ztype_try07613) */
          if ErrorIn(v_local3_try07625) {_Ztype_try07613 = v_local3_try07625
          _Ztype_try07613 = v_local3_try07625
          break
          } else {
          v_local3 = ANY(v_local3_try07625)
          ToList(OBJ(_Ztype_try07613)).PutAt(CLcount,v_local3)
          } 
        }
        /* Iteration-3 */} 
      /* ERROR PROTECTION INSERTED (_Ztype-Result) */
      if ErrorIn(_Ztype_try07613) {Result = _Ztype_try07613
      } else {
      _Ztype = ToList(OBJ(_Ztype_try07613))
      /* Let:3 */{ 
        var s *ClaireProperty   = self.Selector
        /* noccur = 3 */
        ToArray(_Ztype.Id()).NthPut(1,self.CastTo.Id())
        /* Let:4 */{ 
          var prop *ClaireAny  
          /* noccur = 4 */
          if (s.Open == 3) /* If:5 */{ 
            prop = CNIL.Id()
            } else {
            prop = F_Optimize_restriction_I_class(self.CastTo.Class_I(),s.Definition,_Ztype)
            /* If-5 */} 
          if (C_slot.Id() == prop.Isa.Id()) /* If:5 */{ 
            /* Let:6 */{ 
              var g0758 *ClaireSlot   = ToSlot(prop)
              /* noccur = 1 */
              Result = EID{g0758.Range.Id(),0}
              /* Let-6 */} 
            /* If!5 */}  else if (C_method.Id() == prop.Isa.Id()) /* If:5 */{ 
            /* Let:6 */{ 
              var g0759 *ClaireMethod   = ToMethod(prop)
              /* noccur = 1 */
              Result = F_Optimize_use_range_method(g0759,_Ztype)
              /* Let-6 */} 
            } else {
            Result = EID{s.Range.Id(),0}
            /* If-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_type @ Super (throw: true) 
func E_c_type_Super (self EID) EID { 
    return /*(sm for c_type @ Super= EID)*/ F_c_type_Super(Language.To_Super(OBJ(self)) )} 
  
// this is the optimizer for messages
/* {1} OPT.The go function for: c_code(self:Super) [] */
func F_c_code_Super (self *Language.Super ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var s *ClaireProperty   = self.Selector
      /* noccur = 2 */
      /* Let:3 */{ 
        var l *ClaireList   = self.Args
        /* noccur = 2 */
        /* Let:4 */{ 
          var _Ztype *ClaireList  
          /* noccur = 3 */
          var _Ztype_try07665 EID 
          /* Iteration:5 */{ 
            var v_list5 *ClaireList  
            var x *ClaireAny  
            var v_local5 *ClaireAny  
            v_list5 = self.Args
            _Ztype_try07665 = EID{CreateList(ToType(CEMPTY.Id()),v_list5.Length()).Id(),0}
            for CLcount := 0; CLcount < v_list5.Length(); CLcount++{ 
              x = v_list5.At(CLcount)
              var v_local5_try07677 EID 
              v_local5_try07677 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
              /* ERROR PROTECTION INSERTED (v_local5-_Ztype_try07665) */
              if ErrorIn(v_local5_try07677) {_Ztype_try07665 = v_local5_try07677
              _Ztype_try07665 = v_local5_try07677
              break
              } else {
              v_local5 = ANY(v_local5_try07677)
              ToList(OBJ(_Ztype_try07665)).PutAt(CLcount,v_local5)
              } 
            }
            /* Iteration-5 */} 
          /* ERROR PROTECTION INSERTED (_Ztype-Result) */
          if ErrorIn(_Ztype_try07665) {Result = _Ztype_try07665
          } else {
          _Ztype = ToList(OBJ(_Ztype_try07665))
          /* Let:5 */{ 
            var prop *ClaireAny  
            /* noccur = 4 */
            if (s.Open == 3) /* If:6 */{ 
              prop = CNIL.Id()
              } else {
              prop = F_Optimize_restriction_I_class(self.CastTo.Class_I(),s.Definition,_Ztype)
              /* If-6 */} 
            if (C_slot.Id() == prop.Isa.Id()) /* If:6 */{ 
              /* Let:7 */{ 
                var g0763 *ClaireSlot   = ToSlot(prop)
                /* noccur = 4 */
                /* Let:8 */{ 
                  var _CL_obj *Language.CallSlot   = Language.To_CallSlot(new(Language.CallSlot).Is(Language.C_Call_slot))
                  /* noccur = 7 */
                  _CL_obj.Selector = g0763
                  /* update:9 */{ 
                    var va_arg1 *Language.CallSlot  
                    var va_arg2 *ClaireAny  
                    va_arg1 = _CL_obj
                    var va_arg2_try076810 EID 
                    va_arg2_try076810 = Core.F_CALL(C_c_code,ARGS(l.At(1-1).ToEID(),EID{F_Compile_psort_any(Core.F_domain_I_restriction(ToRestriction(g0763.Id())).Id()).Id(),0}))
                    /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                    if ErrorIn(va_arg2_try076810) {Result = va_arg2_try076810
                    } else {
                    va_arg2 = ANY(va_arg2_try076810)
                    /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                    va_arg1.Arg = va_arg2
                    Result = va_arg2.ToEID()
                    }
                    /* update-9 */} 
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  _CL_obj.Test = MakeBoolean((g0763.Range.Contains(g0763.Default) != CTRUE) && (C_compiler.Safety < 5))
                  Result = EID{_CL_obj.Id(),0}
                  }
                  /* Let-8 */} 
                /* Let-7 */} 
              /* If!6 */}  else if (C_method.Id() == prop.Isa.Id()) /* If:6 */{ 
              /* Let:7 */{ 
                var g0764 *ClaireMethod   = ToMethod(prop)
                /* noccur = 1 */
                Result = F_Optimize_c_code_method_method1(g0764,l,_Ztype)
                /* Let-7 */} 
              } else {
              Result = F_Optimize_c_warn_Super(self,_Ztype.Id())
              /* If-6 */} 
            /* Let-5 */} 
          }
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_code @ Super (throw: true) 
func E_c_code_Super (self EID) EID { 
    return /*(sm for c_code @ Super= EID)*/ F_c_code_Super(Language.To_Super(OBJ(self)) )} 
  
// we will need this direct call for compiling call to CLAIRE_demons
/* {1} OPT.The go function for: self_print(self:Call_function2) [] */
func (self *OptimizeCallFunction2 ) SelfPrint () EID { 
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
  
// The EID go function for: self_print @ Call_function2 (throw: true) 
func E_self_print_Call_function2 (self EID) EID { 
    return /*(sm for self_print @ Call_function2= EID)*/ To_OptimizeCallFunction2(OBJ(self)).SelfPrint( )} 
  
/* {1} OPT.The go function for: c_type(self:Call_function2) [] */
func (self *OptimizeCallFunction2 ) CType () *ClaireType  { 
    // use function body compiling 
return  ToType(C_any.Id())
    } 
  
// The EID go function for: c_type @ Call_function2 (throw: false) 
func E_c_type_Call_function2 (self EID) EID { 
    return EID{/*(sm for c_type @ Call_function2= type)*/ To_OptimizeCallFunction2(OBJ(self)).CType( ).Id(),0}} 
  
/* {1} OPT.The go function for: c_code(self:Call_function2) [] */
func F_c_code_Call_function2 (self *OptimizeCallFunction2 ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var _CL_obj *OptimizeCallFunction2   = To_OptimizeCallFunction2(new(OptimizeCallFunction2).Is(C_Optimize_Call_function2))
      /* noccur = 5 */
      _CL_obj.Arg = self.Arg
      /* update:3 */{ 
        var va_arg1 *OptimizeCallFunction2  
        var va_arg2 *ClaireList  
        va_arg1 = _CL_obj
        var va_arg2_try07694 EID 
        /* Iteration:4 */{ 
          var v_list4 *ClaireList  
          var x *ClaireAny  
          var v_local4 *ClaireAny  
          v_list4 = self.Args
          va_arg2_try07694 = EID{CreateList(ToType(CEMPTY.Id()),v_list4.Length()).Id(),0}
          for CLcount := 0; CLcount < v_list4.Length(); CLcount++{ 
            x = v_list4.At(CLcount)
            var v_local4_try07706 EID 
            v_local4_try07706 = Core.F_CALL(C_c_code,ARGS(x.ToEID(),EID{C_any.Id(),0}))
            /* ERROR PROTECTION INSERTED (v_local4-va_arg2_try07694) */
            if ErrorIn(v_local4_try07706) {va_arg2_try07694 = v_local4_try07706
            va_arg2_try07694 = v_local4_try07706
            break
            } else {
            v_local4 = ANY(v_local4_try07706)
            ToList(OBJ(va_arg2_try07694)).PutAt(CLcount,v_local4)
            } 
          }
          /* Iteration-4 */} 
        /* ERROR PROTECTION INSERTED (va_arg2-Result) */
        if ErrorIn(va_arg2_try07694) {Result = va_arg2_try07694
        } else {
        va_arg2 = ToList(OBJ(va_arg2_try07694))
        /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
        va_arg1.Args = va_arg2
        Result = EID{va_arg2.Id(),0}
        }
        /* update-3 */} 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = EID{_CL_obj.Id(),0}
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_code @ Call_function2 (throw: true) 
func E_c_code_Call_function2 (self EID) EID { 
    return /*(sm for c_code @ Call_function2= EID)*/ F_c_code_Call_function2(To_OptimizeCallFunction2(OBJ(self)) )} 
  
// ASSERT & trace
/* {1} OPT.The go function for: c_code(self:Assert) [] */
func F_c_code_Assert (self *Language.Assert ) EID { 
    var Result EID 
    if ((C_compiler.Safety == 0) || 
        (C_compiler.Debug_ask.Length() != 0)) /* If:2 */{ 
      /* Let:3 */{ 
        var g0771UU *ClaireObject  
        /* noccur = 1 */
        /* Let:4 */{ 
          var g0772UU *Language.Call  
          /* noccur = 1 */
          /* Let:5 */{ 
            var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
            /* noccur = 5 */
            _CL_obj.Selector = Core.C_not
            _CL_obj.Args = MakeConstantList(self.Args.At(1-1))
            g0772UU = _CL_obj
            /* Let-5 */} 
          /* Let:5 */{ 
            var g0773UU *Language.Call  
            /* noccur = 1 */
            /* Let:6 */{ 
              var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
              /* noccur = 9 */
              _CL_obj.Selector = Core.C_Core_tformat
              /* update:7 */{ 
                var va_arg1 *Language.Call  
                var va_arg2 *ClaireList  
                va_arg1 = _CL_obj
                /* Construct:8 */{ 
                  var v_bag_arg *ClaireAny  
                  va_arg2= ToType(CEMPTY.Id()).EmptyList()
                  va_arg2.AddFast(MakeString("Assertion violation in ~A line ~A\n").Id())
                  va_arg2.AddFast(MakeInteger(0).Id())
                  /* Let:9 */{ 
                    var _CL_obj *Language.List   = Language.To_List(new(Language.List).Is(Language.C_List))
                    /* noccur = 3 */
                    _CL_obj.Args = MakeConstantList((self.External).Id(),MakeInteger(self.Index).Id())
                    v_bag_arg = _CL_obj.Id()
                    /* Let-9 */} 
                  va_arg2.AddFast(v_bag_arg)/* Construct-8 */} 
                /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                va_arg1.Args = va_arg2
                /* update-7 */} 
              g0773UU = _CL_obj
              /* Let-6 */} 
            g0771UU = ToObject(Language.C_If.Make(g0772UU.Id(),g0773UU.Id(),CFALSE.Id()))
            /* Let-5 */} 
          /* Let-4 */} 
        Result = Core.F_CALL(C_c_code,ARGS(EID{g0771UU.Id(),0},EID{C_any.Id(),0}))
        /* Let-3 */} 
      } else {
      Result = EID{CNIL.Id(),0}
      /* If-2 */} 
    return Result} 
  
// The EID go function for: c_code @ Assert (throw: true) 
func E_c_code_Assert (self EID) EID { 
    return /*(sm for c_code @ Assert= EID)*/ F_c_code_Assert(Language.To_Assert(OBJ(self)) )} 
  
// ignore assertion
/* {1} OPT.The go function for: c_code(self:Trace) [] */
func F_c_code_Trace (self *Language.Trace ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var a *ClaireList   = self.Args
      /* noccur = 11 */
      var g0774I *ClaireBoolean  
      var g0774I_try07753 EID 
      /* and:3 */{ 
        var v_and3 *ClaireBoolean  
        
        v_and3 = Equal(MakeInteger(a.Length()).Id(),MakeInteger(1).Id())
        if (v_and3 == CFALSE) {g0774I_try07753 = EID{CFALSE.Id(),0}
        } else /* arg:4 */{ 
          var v_and3_try07765 EID 
          /* Let:5 */{ 
            var g0777UU *ClaireType  
            /* noccur = 1 */
            var g0777UU_try07786 EID 
            g0777UU_try07786 = Core.F_CALL(C_c_type,ARGS(a.At(1-1).ToEID()))
            /* ERROR PROTECTION INSERTED (g0777UU-v_and3_try07765) */
            if ErrorIn(g0777UU_try07786) {v_and3_try07765 = g0777UU_try07786
            } else {
            g0777UU = ToType(OBJ(g0777UU_try07786))
            v_and3_try07765 = EID{g0777UU.Included(ToType(C_integer.Id())).Id(),0}
            }
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (v_and3-g0774I_try07753) */
          if ErrorIn(v_and3_try07765) {g0774I_try07753 = v_and3_try07765
          } else {
          v_and3 = ToBoolean(OBJ(v_and3_try07765))
          if (v_and3 == CFALSE) {g0774I_try07753 = EID{CFALSE.Id(),0}
          } else /* arg:5 */{ 
            g0774I_try07753 = EID{CTRUE.Id(),0}/* arg-5 */} 
          /* arg-4 */} 
        }
        /* and-3 */} 
      /* ERROR PROTECTION INSERTED (g0774I-Result) */
      if ErrorIn(g0774I_try07753) {Result = g0774I_try07753
      } else {
      g0774I = ToBoolean(OBJ(g0774I_try07753))
      if (g0774I == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var g0779UU *Language.Call  
          /* noccur = 1 */
          /* Let:5 */{ 
            var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
            /* noccur = 5 */
            _CL_obj.Selector = Core.C_write
            _CL_obj.Args = MakeConstantList(C_verbose.Id(),ClEnv.Id(),a.At(1-1))
            g0779UU = _CL_obj
            /* Let-5 */} 
          Result = Core.F_CALL(C_c_code,ARGS(EID{g0779UU.Id(),0}))
          /* Let-4 */} 
        } else {
        var g0780I *ClaireBoolean  
        /* and:4 */{ 
          var v_and4 *ClaireBoolean  
          
          v_and4 = Core.F__sup_integer(a.Length(),1)
          if (v_and4 == CFALSE) {g0780I = CFALSE
          } else /* arg:5 */{ 
            v_and4 = Equal(C_string.Id(),a.At(2-1).Isa.Id())
            if (v_and4 == CFALSE) {g0780I = CFALSE
            } else /* arg:6 */{ 
              /* or:7 */{ 
                var v_or7 *ClaireBoolean  
                
                v_or7 = Core.F__I_equal_any(MakeInteger(C_compiler.Debug_ask.Length()).Id(),MakeInteger(0).Id())
                if (v_or7 == CTRUE) {v_and4 = CTRUE
                } else /* or:8 */{ 
                  var v_or7_try9 EID 
                  h_index := ClEnv.Index /* Handle */
                  h_base := ClEnv.Base
                  /* Let:9 */{ 
                    var g0781UU *ClaireAny  
                    /* noccur = 1 */
                    var g0781UU_try078210 EID 
                    g0781UU_try078210 = EVAL(a.At(1-1))
                    /* ERROR PROTECTION INSERTED (g0781UU-v_or7_try9) */
                    if ErrorIn(g0781UU_try078210) {v_or7_try9 = g0781UU_try078210
                    } else {
                    g0781UU = ANY(g0781UU_try078210)
                    v_or7_try9 = EID{Core.F__inf_equal_integer(ToInteger(g0781UU).Value,ClEnv.Verbose).Id(),0}
                    }
                    /* Let-9 */} 
                  if ErrorIn(v_or7_try9){ 
                    /* s=boolean */ClEnv.Index = h_index
                    ClEnv.Base = h_base
                    v_or7 = CTRUE
                    } else {
                    v_or7 = ToBoolean(OBJ(v_or7_try9))
                    } 
                  if (v_or7 == CTRUE) {v_and4 = CTRUE
                  } else /* or:9 */{ 
                    v_and4 = CFALSE/* org-9 */} 
                  /* org-8 */} 
                /* or-7 */} 
              if (v_and4 == CFALSE) {g0780I = CFALSE
              } else /* arg:7 */{ 
                g0780I = CTRUE/* arg-7 */} 
              /* arg-6 */} 
            /* arg-5 */} 
          /* and-4 */} 
        if (g0780I == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var _Zc *Language.Call  
            /* noccur = 2 */
            /* Let:6 */{ 
              var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
              /* noccur = 9 */
              _CL_obj.Selector = Core.C_Core_tformat
              /* update:7 */{ 
                var va_arg1 *Language.Call  
                var va_arg2 *ClaireList  
                va_arg1 = _CL_obj
                /* Construct:8 */{ 
                  var v_bag_arg *ClaireAny  
                  va_arg2= ToType(CEMPTY.Id()).EmptyList()
                  va_arg2.AddFast(a.At(2-1))
                  va_arg2.AddFast(a.At(1-1))
                  /* Let:9 */{ 
                    var _CL_obj *Language.List   = Language.To_List(new(Language.List).Is(Language.C_List))
                    /* noccur = 3 */
                    _CL_obj.Args = a.Copy().Skip(2)
                    v_bag_arg = _CL_obj.Id()
                    /* Let-9 */} 
                  va_arg2.AddFast(v_bag_arg)/* Construct-8 */} 
                /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                va_arg1.Args = va_arg2
                /* update-7 */} 
              _Zc = _CL_obj
              /* Let-6 */} 
            /* Let:6 */{ 
              var g0783UU *ClaireObject  
              /* noccur = 1 */
              if (C_integer.Id() != a.At(1-1).Isa.Id()) /* If:7 */{ 
                /* Let:8 */{ 
                  var g0784UU *Language.Call  
                  /* noccur = 1 */
                  /* Let:9 */{ 
                    var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                    /* noccur = 11 */
                    _CL_obj.Selector = ToProperty(C__inf_equal.Id())
                    /* update:10 */{ 
                      var va_arg1 *Language.Call  
                      var va_arg2 *ClaireList  
                      va_arg1 = _CL_obj
                      /* Construct:11 */{ 
                        var v_bag_arg *ClaireAny  
                        va_arg2= ToType(CEMPTY.Id()).EmptyList()
                        va_arg2.AddFast(a.At(1-1))
                        /* Let:12 */{ 
                          var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                          /* noccur = 5 */
                          _CL_obj.Selector = C_verbose
                          _CL_obj.Args = MakeConstantList(ClEnv.Id())
                          v_bag_arg = _CL_obj.Id()
                          /* Let-12 */} 
                        va_arg2.AddFast(v_bag_arg)/* Construct-11 */} 
                      /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                      va_arg1.Args = va_arg2
                      /* update-10 */} 
                    g0784UU = _CL_obj
                    /* Let-9 */} 
                  g0783UU = ToObject(Language.C_If.Make(g0784UU.Id(),_Zc.Id(),CFALSE.Id()))
                  /* Let-8 */} 
                } else {
                g0783UU = ToObject(_Zc.Id())
                /* If-7 */} 
              Result = Core.F_CALL(C_c_code,ARGS(EID{g0783UU.Id(),0},EID{C_any.Id(),0}))
              /* Let-6 */} 
            /* Let-5 */} 
          } else {
          Result = EID{CFALSE.Id(),0}
          /* If-4 */} 
        /* If-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_code @ Trace (throw: true) 
func E_c_code_Trace (self EID) EID { 
    return /*(sm for c_code @ Trace= EID)*/ F_c_code_Trace(Language.To_Trace(OBJ(self)) )} 
  
/* {1} OPT.The go function for: c_type(self:Assert) [] */
func F_c_type_Assert (self *Language.Assert ) *ClaireType  { 
    // use function body compiling 
return  ToType(C_any.Id())
    } 
  
// The EID go function for: c_type @ Assert (throw: false) 
func E_c_type_Assert (self EID) EID { 
    return EID{/*(sm for c_type @ Assert= type)*/ F_c_type_Assert(Language.To_Assert(OBJ(self)) ).Id(),0}} 
  
/* {1} OPT.The go function for: c_type(self:Trace) [] */
func F_c_type_Trace (self *Language.Trace ) *ClaireType  { 
    // use function body compiling 
return  ToType(C_any.Id())
    } 
  
// The EID go function for: c_type @ Trace (throw: false) 
func E_c_type_Trace (self EID) EID { 
    return EID{/*(sm for c_type @ Trace= type)*/ F_c_type_Trace(Language.To_Trace(OBJ(self)) ).Id(),0}} 
  
/* {1} OPT.The go function for: c_type(self:Branch) [] */
func F_c_type_Branch (self *Language.Branch ) *ClaireType  { 
    // use function body compiling 
return  ToType(C_boolean.Id())
    } 
  
// The EID go function for: c_type @ Branch (throw: false) 
func E_c_type_Branch (self EID) EID { 
    return EID{/*(sm for c_type @ Branch= type)*/ F_c_type_Branch(Language.To_Branch(OBJ(self)) ).Id(),0}} 
  
/* {1} OPT.The go function for: c_code(self:Branch) [] */
func F_c_code_Branch (self *Language.Branch ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var g0785UU *ClaireObject  
      /* noccur = 1 */
      /* Let:3 */{ 
        var g0786UU *Language.Do  
        /* noccur = 1 */
        /* Let:4 */{ 
          var _CL_obj *Language.Do   = Language.To_Do(new(Language.Do).Is(Language.C_Do))
          /* noccur = 19 */
          /* update:5 */{ 
            var va_arg1 *Language.Do  
            var va_arg2 *ClaireList  
            va_arg1 = _CL_obj
            /* Construct:6 */{ 
              var v_bag_arg *ClaireAny  
              va_arg2= ToType(CEMPTY.Id()).EmptyList()
              /* Let:7 */{ 
                var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                /* noccur = 5 */
                _CL_obj.Selector = C_choice
                _CL_obj.Args = MakeConstantList(ClEnv.Id())
                v_bag_arg = _CL_obj.Id()
                /* Let-7 */} 
              va_arg2.AddFast(v_bag_arg)
              /* Let:7 */{ 
                var g0788UU *Language.Do  
                /* noccur = 1 */
                /* Let:8 */{ 
                  var _CL_obj *Language.Do   = Language.To_Do(new(Language.Do).Is(Language.C_Do))
                  /* noccur = 9 */
                  /* update:9 */{ 
                    var va_arg1 *Language.Do  
                    var va_arg2 *ClaireList  
                    va_arg1 = _CL_obj
                    /* Construct:10 */{ 
                      var v_bag_arg *ClaireAny  
                      va_arg2= ToType(CEMPTY.Id()).EmptyList()
                      /* Let:11 */{ 
                        var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                        /* noccur = 5 */
                        _CL_obj.Selector = C_backtrack
                        _CL_obj.Args = MakeConstantList(ClEnv.Id())
                        v_bag_arg = _CL_obj.Id()
                        /* Let-11 */} 
                      va_arg2.AddFast(v_bag_arg)
                      va_arg2.AddFast(CFALSE.Id())/* Construct-10 */} 
                    /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                    va_arg1.Args = va_arg2
                    /* update-9 */} 
                  g0788UU = _CL_obj
                  /* Let-8 */} 
                v_bag_arg = Language.C_If.Make(self.Args.At(1-1),CTRUE.Id(),g0788UU.Id())
                /* Let-7 */} 
              va_arg2.AddFast(v_bag_arg)/* Construct-6 */} 
            /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
            va_arg1.Args = va_arg2
            /* update-5 */} 
          g0786UU = _CL_obj
          /* Let-4 */} 
        /* Let:4 */{ 
          var g0787UU *Language.Do  
          /* noccur = 1 */
          /* Let:5 */{ 
            var _CL_obj *Language.Do   = Language.To_Do(new(Language.Do).Is(Language.C_Do))
            /* noccur = 9 */
            /* update:6 */{ 
              var va_arg1 *Language.Do  
              var va_arg2 *ClaireList  
              va_arg1 = _CL_obj
              /* Construct:7 */{ 
                var v_bag_arg *ClaireAny  
                va_arg2= ToType(CEMPTY.Id()).EmptyList()
                /* Let:8 */{ 
                  var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                  /* noccur = 5 */
                  _CL_obj.Selector = C_backtrack
                  _CL_obj.Args = MakeConstantList(ClEnv.Id())
                  v_bag_arg = _CL_obj.Id()
                  /* Let-8 */} 
                va_arg2.AddFast(v_bag_arg)
                va_arg2.AddFast(CFALSE.Id())/* Construct-7 */} 
              /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
              va_arg1.Args = va_arg2
              /* update-6 */} 
            g0787UU = _CL_obj
            /* Let-5 */} 
          g0785UU = ToObject(Language.C_Handle.Make(g0786UU.Id(),Core.C_contradiction.Id(),g0787UU.Id()))
          /* Let-4 */} 
        /* Let-3 */} 
      Result = Core.F_CALL(C_c_code,ARGS(EID{g0785UU.Id(),0},EID{C_any.Id(),0}))
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_code @ Branch (throw: true) 
func E_c_code_Branch (self EID) EID { 
    return /*(sm for c_code @ Branch= EID)*/ F_c_code_Branch(Language.To_Branch(OBJ(self)) )} 
  
/* {1} OPT.The go function for: c_code(self:Macro,s:class) [] */
func F_c_code_Macro (self *Language.Macro ,s *ClaireClass ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var g0789UU *ClaireAny  
      /* noccur = 1 */
      var g0789UU_try07903 EID 
      g0789UU_try07903 = Core.F_CALL(Language.C_macroexpand,ARGS(EID{self.Id(),0}))
      /* ERROR PROTECTION INSERTED (g0789UU-Result) */
      if ErrorIn(g0789UU_try07903) {Result = g0789UU_try07903
      } else {
      g0789UU = ANY(g0789UU_try07903)
      Result = Core.F_CALL(C_c_code,ARGS(g0789UU.ToEID(),EID{s.Id(),0}))
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_code @ Macro (throw: true) 
func E_c_code_Macro (self EID,s EID) EID { 
    return /*(sm for c_code @ Macro= EID)*/ F_c_code_Macro(Language.To_Macro(OBJ(self)),ToClass(OBJ(s)) )} 
  
/* {1} OPT.The go function for: c_type(self:Macro) [] */
func F_c_type_Macro (self *Language.Macro ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var g0791UU *ClaireAny  
      /* noccur = 1 */
      var g0791UU_try07923 EID 
      g0791UU_try07923 = Core.F_CALL(Language.C_macroexpand,ARGS(EID{self.Id(),0}))
      /* ERROR PROTECTION INSERTED (g0791UU-Result) */
      if ErrorIn(g0791UU_try07923) {Result = g0791UU_try07923
      } else {
      g0791UU = ANY(g0791UU_try07923)
      Result = Core.F_CALL(C_c_type,ARGS(g0791UU.ToEID()))
      }
      /* Let-2 */} 
    return RangeCheck(ToType(C_type.Id()),Result)} 
  
// The EID go function for: c_type @ Macro (throw: true) 
func E_c_type_Macro (self EID) EID { 
    return /*(sm for c_type @ Macro= EID)*/ F_c_type_Macro(Language.To_Macro(OBJ(self)) )} 
  
/* {1} OPT.The go function for: c_type(self:Printf) [] */
func F_c_type_Printf (self *Language.Printf ) *ClaireType  { 
    // use function body compiling 
return  ToType(C_any.Id())
    } 
  
// The EID go function for: c_type @ Printf (throw: false) 
func E_c_type_Printf (self EID) EID { 
    return EID{/*(sm for c_type @ Printf= type)*/ F_c_type_Printf(Language.To_Printf(OBJ(self)) ).Id(),0}} 
  
/* {1} OPT.The go function for: c_code(self:Printf) [] */
func F_c_code_Printf (self *Language.Printf ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var l *ClaireList   = self.Args
      /* noccur = 8 */
      if (C_string.Id() != l.At(1-1).Isa.Id()) /* If:3 */{ 
        Result = Core.F_CALL(C_Compile_Cerror,ARGS(EID{MakeString("[209] the first argument in ~S must be a string").Id(),0},EID{self.Id(),0}))
        } else {
        /* Let:4 */{ 
          var s *ClaireString   = ToString(l.At(1-1))
          /* noccur = 11 */
          /* Let:5 */{ 
            var i int  = 1
            /* noccur = 8 */
            /* Let:6 */{ 
              var n int  = F_get_string(s,'~')
              /* noccur = 14 */
              /* Let:7 */{ 
                var r *ClaireList   = ToType(C_any.Id()).EmptyList()
                /* noccur = 7 */
                Result= EID{CFALSE.Id(),0}
                for (n != 0) /* while:8 */{ 
                  var void_try9 EID 
                  _ = void_try9
                  /* Let:9 */{ 
                    var m rune  = s.At((n+1))
                    /* noccur = 4 */
                    if (i < l.Length()) /* If:10 */{ 
                      i = (i+1)
                      void_try9 = EID{C__INT,IVAL(i)}
                      } else {
                      void_try9 = Core.F_CALL(C_Compile_Cerror,ARGS(EID{MakeString("[210] not enough arguments in ~S").Id(),0},EID{self.Id(),0}))
                      /* If-10 */} 
                    /* ERROR PROTECTION INSERTED (void_try9-void_try9) */
                    if ErrorIn(void_try9) {Result = void_try9
                    break
                    } else {
                    if (n > 1) /* If:10 */{ 
                      /* Let:11 */{ 
                        var g0793UU *Language.Call  
                        /* noccur = 1 */
                        /* Let:12 */{ 
                          var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                          /* noccur = 5 */
                          _CL_obj.Selector = C_princ
                          _CL_obj.Args = MakeConstantList((F_substring_string(s,1,(n-1))).Id())
                          g0793UU = _CL_obj
                          /* Let-12 */} 
                        r = r.AddFast(g0793UU.Id())
                        /* Let-11 */} 
                      /* If-10 */} 
                    var r_try079410 EID 
                    /* Let:10 */{ 
                      var g0795UU *ClaireAny  
                      /* noccur = 1 */
                      var g0795UU_try079611 EID 
                      if ('A' == m) /* If:11 */{ 
                        /* Let:12 */{ 
                          var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                          /* noccur = 5 */
                          _CL_obj.Selector = C_princ
                          _CL_obj.Args = MakeConstantList(l.At(i-1))
                          g0795UU_try079611 = EID{_CL_obj.Id(),0}
                          /* Let-12 */} 
                        /* If!11 */}  else if ('S' == m) /* If:11 */{ 
                        /* Let:12 */{ 
                          var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                          /* noccur = 5 */
                          _CL_obj.Selector = C_print
                          _CL_obj.Args = MakeConstantList(l.At(i-1))
                          g0795UU_try079611 = EID{_CL_obj.Id(),0}
                          /* Let-12 */} 
                        /* If!11 */}  else if ('F' == m) /* If:11 */{ 
                        /* Let:12 */{ 
                          var p_Z *ClaireBoolean   = CFALSE
                          /* noccur = 4 */
                          /* Let:13 */{ 
                            var j int 
                            /* noccur = 4 */
                            var j_try079714 EID 
                            /* Let:14 */{ 
                              var g0798UU int 
                              /* noccur = 1 */
                              var g0798UU_try079915 EID 
                              /* Let:15 */{ 
                                var g0800UU rune 
                                /* noccur = 1 */
                                var g0800UU_try080116 EID 
                                g0800UU_try080116 = Core.F_nth_get_string(s,(n+2),(n+2))
                                /* ERROR PROTECTION INSERTED (g0800UU-g0798UU_try079915) */
                                if ErrorIn(g0800UU_try080116) {g0798UU_try079915 = g0800UU_try080116
                                } else {
                                g0800UU = CHAR(g0800UU_try080116)
                                g0798UU_try079915 = EID{C__INT,IVAL(int(g0800UU))}
                                }
                                /* Let-15 */} 
                              /* ERROR PROTECTION INSERTED (g0798UU-j_try079714) */
                              if ErrorIn(g0798UU_try079915) {j_try079714 = g0798UU_try079915
                              } else {
                              g0798UU = INT(g0798UU_try079915)
                              j_try079714 = EID{C__INT,IVAL((g0798UU-48))}
                              }
                              /* Let-14 */} 
                            /* ERROR PROTECTION INSERTED (j-g0795UU_try079611) */
                            if ErrorIn(j_try079714) {g0795UU_try079611 = j_try079714
                            } else {
                            j = INT(j_try079714)
                            if ('%' == s.At((n+2))) /* If:14 */{ 
                              p_Z = CTRUE
                              j = 1
                              g0795UU_try079611 = EID{C__INT,IVAL(j)}
                              /* If!14 */}  else if ((j < 0) || 
                                (j > 9)) /* If:14 */{ 
                              g0795UU_try079611 = ToException(Core.C_general_error.Make(MakeString("[189] F requires a single digit integer in ~S").Id(),MakeConstantList(self.Id()).Id())).Close()
                              } else {
                              g0795UU_try079611 = EID{CFALSE.Id(),0}
                              /* If-14 */} 
                            /* ERROR PROTECTION INSERTED (g0795UU_try079611-g0795UU_try079611) */
                            if !ErrorIn(g0795UU_try079611) {
                            if ((p_Z != CTRUE) && 
                                ('%' == s.At((n+3)))) /* If:14 */{ 
                              p_Z = CTRUE
                              n = (n+1)
                              /* If-14 */} 
                            n = (n+1)
                            /* Let:14 */{ 
                              var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                              /* noccur = 11 */
                              _CL_obj.Selector = Core.C_mClaire_printFDigit
                              /* update:15 */{ 
                                var va_arg1 *Language.Call  
                                var va_arg2 *ClaireList  
                                va_arg1 = _CL_obj
                                /* Construct:16 */{ 
                                  var v_bag_arg *ClaireAny  
                                  va_arg2= ToType(CEMPTY.Id()).EmptyList()
                                  if (p_Z == CTRUE) /* If:17 */{ 
                                    /* Let:18 */{ 
                                      var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                      /* noccur = 5 */
                                      _CL_obj.Selector = ToProperty(C__star.Id())
                                      _CL_obj.Args = MakeConstantList(l.At(i-1),MakeFloat(100).Id())
                                      v_bag_arg = _CL_obj.Id()
                                      /* Let-18 */} 
                                    } else {
                                    v_bag_arg = l.At(i-1)
                                    /* If-17 */} 
                                  va_arg2.AddFast(v_bag_arg)
                                  va_arg2.AddFast(MakeInteger(j).Id())/* Construct-16 */} 
                                /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                                va_arg1.Args = va_arg2
                                /* update-15 */} 
                              g0795UU_try079611 = EID{_CL_obj.Id(),0}
                              /* Let-14 */} 
                            }
                            }
                            /* Let-13 */} 
                          /* Let-12 */} 
                        /* If!11 */}  else if ('I' == m) /* If:11 */{ 
                        g0795UU_try079611 = l.At(i-1).ToEID()
                        } else {
                        g0795UU_try079611 = EID{CFALSE.Id(),0}
                        /* If-11 */} 
                      /* ERROR PROTECTION INSERTED (g0795UU-r_try079410) */
                      if ErrorIn(g0795UU_try079611) {r_try079410 = g0795UU_try079611
                      } else {
                      g0795UU = ANY(g0795UU_try079611)
                      r_try079410 = EID{r.AddFast(g0795UU).Id(),0}
                      }
                      /* Let-10 */} 
                    /* ERROR PROTECTION INSERTED (r-void_try9) */
                    if ErrorIn(r_try079410) {void_try9 = r_try079410
                    Result = r_try079410
                    break
                    } else {
                    r = ToList(OBJ(r_try079410))
                    void_try9 = EID{r.Id(),0}
                    s = F_substring_string(s,(n+2),1000)
                    n = F_get_string(s,'~')
                    void_try9 = EID{C__INT,IVAL(n)}
                    }}
                    /* Let-9 */} 
                  /* ERROR PROTECTION INSERTED (void_try9-Result) */
                  if ErrorIn(void_try9) {Result = void_try9
                  Result = void_try9
                  break
                  } else {
                  /* while-8 */} 
                }
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                if (F_boolean_I_any((s).Id()) == CTRUE) /* If:8 */{ 
                  /* Let:9 */{ 
                    var g0802UU *Language.Call  
                    /* noccur = 1 */
                    /* Let:10 */{ 
                      var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                      /* noccur = 5 */
                      _CL_obj.Selector = C_princ
                      _CL_obj.Args = MakeConstantList((s).Id())
                      g0802UU = _CL_obj
                      /* Let-10 */} 
                    r = r.AddFast(g0802UU.Id())
                    /* Let-9 */} 
                  /* If-8 */} 
                /* Let:8 */{ 
                  var g0803UU *Language.Do  
                  /* noccur = 1 */
                  /* Let:9 */{ 
                    var _CL_obj *Language.Do   = Language.To_Do(new(Language.Do).Is(Language.C_Do))
                    /* noccur = 3 */
                    _CL_obj.Args = r
                    g0803UU = _CL_obj
                    /* Let-9 */} 
                  Result = Core.F_CALL(C_c_code,ARGS(EID{g0803UU.Id(),0},EID{C_any.Id(),0}))
                  /* Let-8 */} 
                }
                /* Let-7 */} 
              /* Let-6 */} 
            /* Let-5 */} 
          /* Let-4 */} 
        /* If-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_code @ Printf (throw: true) 
func E_c_code_Printf (self EID) EID { 
    return /*(sm for c_code @ Printf= EID)*/ F_c_code_Printf(Language.To_Printf(OBJ(self)) )} 
  
/* {1} OPT.The go function for: c_type(self:Error) [] */
func F_c_type_Error (self *Language.Error ) *ClaireType  { 
    // use function body compiling 
return  ToType(CEMPTY.Id())
    } 
  
// The EID go function for: c_type @ Error (throw: false) 
func E_c_type_Error (self EID) EID { 
    return EID{/*(sm for c_type @ Error= type)*/ F_c_type_Error(Language.To_Error(OBJ(self)) ).Id(),0}} 
  
/* {1} OPT.The go function for: c_code(self:Error) [] */
func F_c_code_Error (self *Language.Error ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var g0804UU *Language.Call  
      /* noccur = 1 */
      var g0804UU_try08053 EID 
      /* Let:3 */{ 
        var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
        /* noccur = 20 */
        _CL_obj.Selector = C_close
        /* update:4 */{ 
          var va_arg1 *Language.Call  
          var va_arg2 *ClaireList  
          va_arg1 = _CL_obj
          var va_arg2_try08065 EID 
          /* Construct:5 */{ 
            var v_bag_arg *ClaireAny  
            va_arg2_try08065= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
            var v_bag_arg_try08076 EID 
            /* Let:6 */{ 
              var _CL_obj *Language.Cast   = Language.To_Cast(new(Language.Cast).Is(Language.C_Cast))
              /* noccur = 14 */
              /* update:7 */{ 
                var va_arg1 *Language.Cast  
                var va_arg2 *ClaireAny  
                va_arg1 = _CL_obj
                var va_arg2_try08088 EID 
                /* Let:8 */{ 
                  var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                  /* noccur = 9 */
                  _CL_obj.Selector = C_Compile_anyObject_I
                  /* update:9 */{ 
                    var va_arg1 *Language.Call  
                    var va_arg2 *ClaireList  
                    va_arg1 = _CL_obj
                    var va_arg2_try080910 EID 
                    /* Construct:10 */{ 
                      var v_bag_arg *ClaireAny  
                      va_arg2_try080910= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                      ToList(OBJ(va_arg2_try080910)).AddFast(Core.C_general_error.Id())
                      var v_bag_arg_try081011 EID 
                      /* Let:11 */{ 
                        var g0811UU *ClaireAny  
                        /* noccur = 1 */
                        var g0811UU_try081212 EID 
                        g0811UU_try081212 = Core.F_car_list(self.Args)
                        /* ERROR PROTECTION INSERTED (g0811UU-v_bag_arg_try081011) */
                        if ErrorIn(g0811UU_try081212) {v_bag_arg_try081011 = g0811UU_try081212
                        } else {
                        g0811UU = ANY(g0811UU_try081212)
                        v_bag_arg_try081011 = Core.F_CALL(C_c_code,ARGS(g0811UU.ToEID(),EID{C_any.Id(),0}))
                        }
                        /* Let-11 */} 
                      /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try080910) */
                      if ErrorIn(v_bag_arg_try081011) {va_arg2_try080910 = v_bag_arg_try081011
                      } else {
                      v_bag_arg = ANY(v_bag_arg_try081011)
                      ToList(OBJ(va_arg2_try080910)).AddFast(v_bag_arg)
                      var v_bag_arg_try081311 EID 
                      /* Let:11 */{ 
                        var g0814UU *ClaireObject  
                        /* noccur = 1 */
                        var g0814UU_try081512 EID 
                        if (self.Args.Length() != 1) /* If:12 */{ 
                          /* Let:13 */{ 
                            var _CL_obj *Language.List   = Language.To_List(new(Language.List).Is(Language.C_List))
                            /* noccur = 3 */
                            /* update:14 */{ 
                              var va_arg1 *Language.Construct  
                              var va_arg2 *ClaireList  
                              va_arg1 = Language.To_Construct(_CL_obj.Id())
                              var va_arg2_try081615 EID 
                              va_arg2_try081615 = self.Args.Cdr()
                              /* ERROR PROTECTION INSERTED (va_arg2-g0814UU_try081512) */
                              if ErrorIn(va_arg2_try081615) {g0814UU_try081512 = va_arg2_try081615
                              } else {
                              va_arg2 = ToList(OBJ(va_arg2_try081615))
                              /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                              va_arg1.Args = va_arg2
                              g0814UU_try081512 = EID{va_arg2.Id(),0}
                              }
                              /* update-14 */} 
                            /* ERROR PROTECTION INSERTED (g0814UU_try081512-g0814UU_try081512) */
                            if !ErrorIn(g0814UU_try081512) {
                            g0814UU_try081512 = EID{_CL_obj.Id(),0}
                            }
                            /* Let-13 */} 
                          } else {
                          g0814UU_try081512 = EID{CNIL.Id(),0}
                          /* If-12 */} 
                        /* ERROR PROTECTION INSERTED (g0814UU-v_bag_arg_try081311) */
                        if ErrorIn(g0814UU_try081512) {v_bag_arg_try081311 = g0814UU_try081512
                        } else {
                        g0814UU = ToObject(OBJ(g0814UU_try081512))
                        v_bag_arg_try081311 = Core.F_CALL(C_c_code,ARGS(EID{g0814UU.Id(),0},EID{C_any.Id(),0}))
                        }
                        /* Let-11 */} 
                      /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try080910) */
                      if ErrorIn(v_bag_arg_try081311) {va_arg2_try080910 = v_bag_arg_try081311
                      } else {
                      v_bag_arg = ANY(v_bag_arg_try081311)
                      ToList(OBJ(va_arg2_try080910)).AddFast(v_bag_arg)}}
                      /* Construct-10 */} 
                    /* ERROR PROTECTION INSERTED (va_arg2-va_arg2_try08088) */
                    if ErrorIn(va_arg2_try080910) {va_arg2_try08088 = va_arg2_try080910
                    } else {
                    va_arg2 = ToList(OBJ(va_arg2_try080910))
                    /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                    va_arg1.Args = va_arg2
                    va_arg2_try08088 = EID{va_arg2.Id(),0}
                    }
                    /* update-9 */} 
                  /* ERROR PROTECTION INSERTED (va_arg2_try08088-va_arg2_try08088) */
                  if !ErrorIn(va_arg2_try08088) {
                  va_arg2_try08088 = EID{_CL_obj.Id(),0}
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (va_arg2-v_bag_arg_try08076) */
                if ErrorIn(va_arg2_try08088) {v_bag_arg_try08076 = va_arg2_try08088
                } else {
                va_arg2 = ANY(va_arg2_try08088)
                /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                va_arg1.Arg = va_arg2
                v_bag_arg_try08076 = va_arg2.ToEID()
                }
                /* update-7 */} 
              /* ERROR PROTECTION INSERTED (v_bag_arg_try08076-v_bag_arg_try08076) */
              if !ErrorIn(v_bag_arg_try08076) {
              _CL_obj.SetArg = ToType(C_exception.Id())
              v_bag_arg_try08076 = EID{_CL_obj.Id(),0}
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try08065) */
            if ErrorIn(v_bag_arg_try08076) {va_arg2_try08065 = v_bag_arg_try08076
            } else {
            v_bag_arg = ANY(v_bag_arg_try08076)
            ToList(OBJ(va_arg2_try08065)).AddFast(v_bag_arg)}
            /* Construct-5 */} 
          /* ERROR PROTECTION INSERTED (va_arg2-g0804UU_try08053) */
          if ErrorIn(va_arg2_try08065) {g0804UU_try08053 = va_arg2_try08065
          } else {
          va_arg2 = ToList(OBJ(va_arg2_try08065))
          /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
          va_arg1.Args = va_arg2
          g0804UU_try08053 = EID{va_arg2.Id(),0}
          }
          /* update-4 */} 
        /* ERROR PROTECTION INSERTED (g0804UU_try08053-g0804UU_try08053) */
        if !ErrorIn(g0804UU_try08053) {
        g0804UU_try08053 = EID{_CL_obj.Id(),0}
        }
        /* Let-3 */} 
      /* ERROR PROTECTION INSERTED (g0804UU-Result) */
      if ErrorIn(g0804UU_try08053) {Result = g0804UU_try08053
      } else {
      g0804UU = Language.To_Call(OBJ(g0804UU_try08053))
      Result = Core.F_CALL(C_c_code,ARGS(EID{g0804UU.Id(),0},EID{C_void.Id(),0}))
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_code @ Error (throw: true) 
func E_c_code_Error (self EID) EID { 
    return /*(sm for c_code @ Error= EID)*/ F_c_code_Error(Language.To_Error(OBJ(self)) )} 
  
// *********************************************************************
// *     Part 3: If, Case, Do, Let                                     *
// *********************************************************************
//_______________ IF __________________________________________
// check if the test is of the form known?(v) so that the type (result) can be reduced
/* {1} OPT.The go function for: extendedTest?(self:If) [] */
func F_Optimize_extendedTest_ask_If (self *Language.If ) *ClaireType  { 
    // procedure body with s =  
var Result *ClaireType  
    /* Let:2 */{ 
      var _Zt *ClaireAny   = self.Test
      /* noccur = 2 */
      if (_Zt.Isa.IsIn(Language.C_Call) == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var g0817 *Language.Call   = Language.To_Call(_Zt)
          /* noccur = 3 */
          if ((g0817.Args.At(1-1).Isa.IsIn(C_Variable) == CTRUE) && 
              (g0817.Selector.Id() == Core.C_known_ask.Id())) /* If:5 */{ 
            Result = ToType(OBJ(Core.F_CALL(C_range,ARGS(g0817.Args.At(1-1).ToEID()))))
            } else {
            Result = ToType(C_any.Id())
            /* If-5 */} 
          /* Let-4 */} 
        } else {
        Result = ToType(C_any.Id())
        /* If-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: extendedTest? @ If (throw: false) 
func E_Optimize_extendedTest_ask_If (self EID) EID { 
    return EID{/*(sm for extendedTest? @ If= type)*/ F_Optimize_extendedTest_ask_If(Language.To_If(OBJ(self)) ).Id(),0}} 
  
// notice that we analyze the test to detect the know? filter
/* {1} OPT.The go function for: c_type(self:If) [] */
func F_c_type_If (self *Language.If ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var _Zr *ClaireType   = F_Optimize_extendedTest_ask_If(self)
      /* noccur = 4 */
      var g0819I *ClaireBoolean  
      var g0819I_try08203 EID 
      g0819I_try08203 = F_Optimize_extended_ask_type(_Zr)
      /* ERROR PROTECTION INSERTED (g0819I-Result) */
      if ErrorIn(g0819I_try08203) {Result = g0819I_try08203
      } else {
      g0819I = ToBoolean(OBJ(g0819I_try08203))
      if (g0819I == CTRUE) /* If:3 */{ 
        F_Optimize_range_sets_any(ToList(OBJ(Core.F_CALL(C_args,ARGS(self.Test.ToEID())))).At(1-1),F_Optimize_sort_abstract_I_type(ToType(OBJ(Core.F_CALL(C_mClaire_t1,ARGS(EID{_Zr.Id(),0}))))))
        Result = EVOID
        } else {
        Result = EID{CFALSE.Id(),0}
        /* If-3 */} 
      }
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      /* Let:3 */{ 
        var result *ClaireType  
        /* noccur = 1 */
        var result_try08214 EID 
        /* Let:4 */{ 
          var g0822UU *ClaireType  
          /* noccur = 1 */
          var g0822UU_try08245 EID 
          g0822UU_try08245 = Core.F_CALL(C_c_type,ARGS(self.Arg.ToEID()))
          /* ERROR PROTECTION INSERTED (g0822UU-result_try08214) */
          if ErrorIn(g0822UU_try08245) {result_try08214 = g0822UU_try08245
          } else {
          g0822UU = ToType(OBJ(g0822UU_try08245))
          /* Let:5 */{ 
            var g0823UU *ClaireType  
            /* noccur = 1 */
            var g0823UU_try08256 EID 
            g0823UU_try08256 = Core.F_CALL(C_c_type,ARGS(self.Other.ToEID()))
            /* ERROR PROTECTION INSERTED (g0823UU-result_try08214) */
            if ErrorIn(g0823UU_try08256) {result_try08214 = g0823UU_try08256
            } else {
            g0823UU = ToType(OBJ(g0823UU_try08256))
            result_try08214 = EID{Core.F_U_type(g0822UU,g0823UU).Id(),0}
            }
            /* Let-5 */} 
          }
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (result-Result) */
        if ErrorIn(result_try08214) {Result = result_try08214
        } else {
        result = ToType(OBJ(result_try08214))
        var g0826I *ClaireBoolean  
        var g0826I_try08274 EID 
        g0826I_try08274 = F_Optimize_extended_ask_type(_Zr)
        /* ERROR PROTECTION INSERTED (g0826I-Result) */
        if ErrorIn(g0826I_try08274) {Result = g0826I_try08274
        } else {
        g0826I = ToBoolean(OBJ(g0826I_try08274))
        if (g0826I == CTRUE) /* If:4 */{ 
          Result = Core.F_put_property2(C_range,ToObject(ToList(OBJ(Core.F_CALL(C_args,ARGS(self.Test.ToEID())))).At(1-1)),_Zr.Id())
          } else {
          Result = EID{CFALSE.Id(),0}
          /* If-4 */} 
        }
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        Result = EID{result.Id(),0}
        }
        }
        /* Let-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_type @ If (throw: true) 
func E_c_type_If (self EID) EID { 
    return /*(sm for c_type @ If= EID)*/ F_c_type_If(Language.To_If(OBJ(self)) )} 
  
// debug boolean variable to flag the use of extented X U {unknown} + test : if known?(x)
/* {1} OPT.The go function for: c_code(self:If,s:class) [] */
func F_c_code_If (self *Language.If ,s *ClaireClass ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var _Zr *ClaireType   = F_Optimize_extendedTest_ask_If(self)
      /* noccur = 4 */
      var g0828I *ClaireBoolean  
      var g0828I_try08293 EID 
      g0828I_try08293 = F_Optimize_extended_ask_type(_Zr)
      /* ERROR PROTECTION INSERTED (g0828I-Result) */
      if ErrorIn(g0828I_try08293) {Result = g0828I_try08293
      } else {
      g0828I = ToBoolean(OBJ(g0828I_try08293))
      if (g0828I == CTRUE) /* If:3 */{ 
        F_Optimize_range_sets_any(ToList(OBJ(Core.F_CALL(C_args,ARGS(self.Test.ToEID())))).At(1-1),F_Optimize_sort_abstract_I_type(ToType(OBJ(Core.F_CALL(C_mClaire_t1,ARGS(EID{_Zr.Id(),0}))))))
        Result = EVOID
        } else {
        Result = EID{CFALSE.Id(),0}
        /* If-3 */} 
      }
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      var g0830I *ClaireBoolean  
      var g0830I_try08313 EID 
      /* and:3 */{ 
        var v_and3 *ClaireBoolean  
        
        var v_and3_try08324 EID 
        /* Let:4 */{ 
          var g0833UU *ClaireBoolean  
          /* noccur = 1 */
          var g0833UU_try08345 EID 
          /* Let:5 */{ 
            var g0835UU *ClaireType  
            /* noccur = 1 */
            var g0835UU_try08366 EID 
            /* Let:6 */{ 
              var g0837UU *ClaireType  
              /* noccur = 1 */
              var g0837UU_try08387 EID 
              g0837UU_try08387 = Core.F_CALL(C_c_type,ARGS(self.Test.ToEID()))
              /* ERROR PROTECTION INSERTED (g0837UU-g0835UU_try08366) */
              if ErrorIn(g0837UU_try08387) {g0835UU_try08366 = g0837UU_try08387
              } else {
              g0837UU = ToType(OBJ(g0837UU_try08387))
              g0835UU_try08366 = EID{F_Optimize_ptype_type(g0837UU).Id(),0}
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (g0835UU-g0833UU_try08345) */
            if ErrorIn(g0835UU_try08366) {g0833UU_try08345 = g0835UU_try08366
            } else {
            g0835UU = ToType(OBJ(g0835UU_try08366))
            g0833UU_try08345 = EID{g0835UU.Included(ToType(C_boolean.Id())).Id(),0}
            }
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (g0833UU-v_and3_try08324) */
          if ErrorIn(g0833UU_try08345) {v_and3_try08324 = g0833UU_try08345
          } else {
          g0833UU = ToBoolean(OBJ(g0833UU_try08345))
          v_and3_try08324 = EID{g0833UU.Not.Id(),0}
          }
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (v_and3-g0830I_try08313) */
        if ErrorIn(v_and3_try08324) {g0830I_try08313 = v_and3_try08324
        } else {
        v_and3 = ToBoolean(OBJ(v_and3_try08324))
        if (v_and3 == CFALSE) {g0830I_try08313 = EID{CFALSE.Id(),0}
        } else /* arg:4 */{ 
          v_and3 = Equal(C_PENIBLE.Value,CTRUE.Id())
          if (v_and3 == CFALSE) {g0830I_try08313 = EID{CFALSE.Id(),0}
          } else /* arg:5 */{ 
            g0830I_try08313 = EID{CTRUE.Id(),0}/* arg-5 */} 
          /* arg-4 */} 
        }
        /* and-3 */} 
      /* ERROR PROTECTION INSERTED (g0830I-Result) */
      if ErrorIn(g0830I_try08313) {Result = g0830I_try08313
      } else {
      g0830I = ToBoolean(OBJ(g0830I_try08313))
      if (g0830I == CTRUE) /* If:3 */{ 
        F_Compile_warn_void()
        Result = Core.F_tformat_string(MakeString("CLAIRE 3.3 SYNTAX - Test in ~S should be a boolean [260]\n"),2,MakeConstantList(self.Id()))
        } else {
        Result = EID{CFALSE.Id(),0}
        /* If-3 */} 
      }
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      /* Let:3 */{ 
        var result *Language.If  
        /* noccur = 1 */
        var result_try08394 EID 
        /* Let:4 */{ 
          var g0840UU *ClaireAny  
          /* noccur = 1 */
          var g0840UU_try08435 EID 
          g0840UU_try08435 = F_Optimize_c_boolean_any(self.Test)
          /* ERROR PROTECTION INSERTED (g0840UU-result_try08394) */
          if ErrorIn(g0840UU_try08435) {result_try08394 = g0840UU_try08435
          } else {
          g0840UU = ANY(g0840UU_try08435)
          /* Let:5 */{ 
            var g0841UU *ClaireAny  
            /* noccur = 1 */
            var g0841UU_try08446 EID 
            g0841UU_try08446 = Core.F_CALL(C_c_code,ARGS(self.Arg.ToEID(),EID{s.Id(),0}))
            /* ERROR PROTECTION INSERTED (g0841UU-result_try08394) */
            if ErrorIn(g0841UU_try08446) {result_try08394 = g0841UU_try08446
            } else {
            g0841UU = ANY(g0841UU_try08446)
            /* Let:6 */{ 
              var g0842UU *ClaireAny  
              /* noccur = 1 */
              var g0842UU_try08457 EID 
              g0842UU_try08457 = Core.F_CALL(C_c_code,ARGS(self.Other.ToEID(),EID{s.Id(),0}))
              /* ERROR PROTECTION INSERTED (g0842UU-result_try08394) */
              if ErrorIn(g0842UU_try08457) {result_try08394 = g0842UU_try08457
              } else {
              g0842UU = ANY(g0842UU_try08457)
              result_try08394 = Language.C_If.Make(g0840UU,g0841UU,g0842UU).ToEID()
              }
              /* Let-6 */} 
            }
            /* Let-5 */} 
          }
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (result-Result) */
        if ErrorIn(result_try08394) {Result = result_try08394
        } else {
        result = Language.To_If(OBJ(result_try08394))
        var g0846I *ClaireBoolean  
        var g0846I_try08474 EID 
        g0846I_try08474 = F_Optimize_extended_ask_type(_Zr)
        /* ERROR PROTECTION INSERTED (g0846I-Result) */
        if ErrorIn(g0846I_try08474) {Result = g0846I_try08474
        } else {
        g0846I = ToBoolean(OBJ(g0846I_try08474))
        if (g0846I == CTRUE) /* If:4 */{ 
          Result = Core.F_put_property2(C_range,ToObject(ToList(OBJ(Core.F_CALL(C_args,ARGS(self.Test.ToEID())))).At(1-1)),_Zr.Id())
          } else {
          Result = EID{CFALSE.Id(),0}
          /* If-4 */} 
        }
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        Result = EID{result.Id(),0}
        }
        }
        /* Let-3 */} 
      }}
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_code @ If (throw: true) 
func E_c_code_If (self EID,s EID) EID { 
    return /*(sm for c_code @ If= EID)*/ F_c_code_If(Language.To_If(OBJ(self)),ToClass(OBJ(s)) )} 
  
// ------------------ CASE -------------------------------------------
// a member-of is a CLAIRE case. [yc 1/29/98]
// note that type inference supposes that the case is "closed" (all types are delt with)
// but only with safety >= 5
/* {1} OPT.The go function for: c_type(self:Case) [] */
func F_c_type_Case (self *Language.Case ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var _Zvar *ClaireAny   = self.ClaireVar
      /* noccur = 6 */
      /* Let:3 */{ 
        var _Ztype *ClaireAny  
        /* noccur = 4 */
        if (_Zvar.Isa.IsIn(C_Variable) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var g0848 *ClaireVariable   = To_Variable(_Zvar)
            /* noccur = 1 */
            _Ztype = Core.F_get_property(C_range,ToObject(g0848.Id()))
            /* Let-5 */} 
          } else {
          _Ztype = C_any.Id()
          /* If-4 */} 
        /* Let:4 */{ 
          var l *ClaireList   = self.Args.Copy()
          /* noccur = 10 */
          /* Let:5 */{ 
            var rtype *ClaireType   = ToType(CEMPTY.Id())
            /* noccur = 4 */
            /* Let:6 */{ 
              var utype *ClaireType   = ToType(CEMPTY.Id())
              /* noccur = 3 */
              Result= EID{CFALSE.Id(),0}
              for (l.Length() > 0) /* while:7 */{ 
                var void_try8 EID 
                _ = void_try8
                { 
                if (l.At(1-1).Isa.IsIn(C_type) == CTRUE) /* If:8 */{ 
                  utype = Core.F_U_type(utype,ToType(l.At(1-1)))
                  if (F_Compile_osort_any(_Ztype).Id() == F_Compile_osort_any(l.At(1-1)).Id()) /* If:9 */{ 
                    F_Optimize_range_sets_any(_Zvar,ToType(l.At(1-1)))
                    void_try8 = EVOID
                    /* If!9 */}  else if (F_Compile_osort_any(_Ztype).Id() == C_any.Id()) /* If:9 */{ 
                    F_Optimize_range_sets_any(_Zvar,F_Optimize_sort_abstract_I_type(ToType(l.At(1-1))))
                    void_try8 = EVOID
                    } else {
                    void_try8 = EID{CFALSE.Id(),0}
                    /* If-9 */} 
                  } else {
                  /* Let:9 */{ 
                    var g0851UU *ClaireAny  
                    /* noccur = 1 */
                    var g0851UU_try085210 EID 
                    g0851UU_try085210 = Core.F_car_list(l)
                    /* ERROR PROTECTION INSERTED (g0851UU-void_try8) */
                    if ErrorIn(g0851UU_try085210) {void_try8 = g0851UU_try085210
                    } else {
                    g0851UU = ANY(g0851UU_try085210)
                    void_try8 = Core.F_CALL(C_Compile_Cerror,ARGS(EID{MakeString("[208] wrong type declaration for case: ~S in ~S").Id(),0},g0851UU.ToEID(),EID{self.Id(),0}))
                    }
                    /* Let-9 */} 
                  /* If-8 */} 
                /* ERROR PROTECTION INSERTED (void_try8-void_try8) */
                if ErrorIn(void_try8) {Result = void_try8
                break
                } else {
                var rtype_try08538 EID 
                /* Let:8 */{ 
                  var g0854UU *ClaireType  
                  /* noccur = 1 */
                  var g0854UU_try08559 EID 
                  g0854UU_try08559 = Core.F_CALL(C_c_type,ARGS(l.At(2-1).ToEID()))
                  /* ERROR PROTECTION INSERTED (g0854UU-rtype_try08538) */
                  if ErrorIn(g0854UU_try08559) {rtype_try08538 = g0854UU_try08559
                  } else {
                  g0854UU = ToType(OBJ(g0854UU_try08559))
                  rtype_try08538 = EID{Core.F_U_type(rtype,g0854UU).Id(),0}
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (rtype-void_try8) */
                if ErrorIn(rtype_try08538) {void_try8 = rtype_try08538
                Result = rtype_try08538
                break
                } else {
                rtype = ToType(OBJ(rtype_try08538))
                void_try8 = EID{rtype.Id(),0}
                
                if (_Zvar.Isa.IsIn(C_Variable) == CTRUE) /* If:8 */{ 
                  /* Let:9 */{ 
                    var g0850 *ClaireVariable   = To_Variable(_Zvar)
                    /* noccur = 1 */
                    g0850.Range = ToType(_Ztype)
                    /* Let-9 */} 
                  /* If-8 */} 
                l = l.Skip(2)
                }}
                /* while-7 */} 
              }
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              if (ToType(_Ztype).Included(utype) == CTRUE) /* If:7 */{ 
                Result = EID{rtype.Id(),0}
                /* If!7 */}  else if (rtype.Included(ToType(C_boolean.Id())) == CTRUE) /* If:7 */{ 
                Result = EID{C_boolean.Id(),0}
                } else {
                Result = EID{C_any.Id(),0}
                /* If-7 */} 
              }
              /* Let-6 */} 
            /* Let-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_type @ Case (throw: true) 
func E_c_type_Case (self EID) EID { 
    return /*(sm for c_type @ Case= EID)*/ F_c_type_Case(Language.To_Case(OBJ(self)) )} 
  
// safety
// utility : create a branch with substituted variable
// notice the use of occurence: create a let only if necessary :)
/* {1} OPT.The go function for: case_branch(x:any,%var:any,%type:type) [] */
func F_Optimize_case_branch_any (x *ClaireAny ,_Zvar *ClaireAny ,_Ztype *ClaireType ) *ClaireAny  { 
    // procedure body with s =  
var Result *ClaireAny  
    if (_Zvar.Isa.IsIn(C_Variable) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0856 *ClaireVariable   = To_Variable(_Zvar)
        /* noccur = 4 */
        /* Let:4 */{ 
          var vsub *ClaireVariable   = F_Compile_Variable_I_symbol(Core.F_gensym_void(),0,_Ztype.Id())
          /* noccur = 2 */
          if ((Equal(_Ztype.Id(),g0856.Range.Id()) != CTRUE) && 
              ((_Ztype.Id() != C_any.Id()) && 
                (Language.F_occurrence_any(x,g0856) > 0))) /* If:5 */{ 
            /* Let:6 */{ 
              var _CL_obj *Language.Let   = Language.To_Let(new(Language.Let).Is(Language.C_Let))
              /* noccur = 6 */
              _CL_obj.ClaireVar = vsub
              _CL_obj.Value = g0856.Id()
              _CL_obj.Arg = F_Optimize_case_substitution_any(x,g0856,vsub)
              Result = _CL_obj.Id()
              /* Let-6 */} 
            } else {
            Result = x
            /* If-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      } else {
      Result = x
      /* If-2 */} 
    return Result} 
  
// The EID go function for: case_branch @ any (throw: false) 
func E_Optimize_case_branch_any (x EID,_Zvar EID,_Ztype EID) EID { 
    return /*(sm for case_branch @ any= any)*/ F_Optimize_case_branch_any(ANY(x),ANY(_Zvar),ToType(OBJ(_Ztype)) ).ToEID()} 
  
// this gets tricky, if the variable is changed in the case_branch, we need to update the original variable
// we add a copy to be able to compile the compiler (do not change the original code)
/* {1} OPT.The go function for: case_substitution(x:any,%var:Variable,vsub:Variable) [] */
func F_Optimize_case_substitution_any (x *ClaireAny ,_Zvar *ClaireVariable ,vsub *ClaireVariable ) *ClaireAny  { 
    // procedure body with s =  
var Result *ClaireAny  
    /* Let:2 */{ 
      var y *ClaireAny   = Language.F_substitution_any(Language.F_instruction_copy_any(x),_Zvar,vsub.Id())
      /* noccur = 3 */
      if (Language.F_occurchange_any(y,vsub) == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var _CL_obj *Language.Do   = Language.To_Do(new(Language.Do).Is(Language.C_Do))
          /* noccur = 3 */
          _CL_obj.Args = MakeConstantList(y,Language.C_Assign.Make(_Zvar.Id(),vsub.Id()))
          Result = _CL_obj.Id()
          /* Let-4 */} 
        } else {
        Result = y
        /* If-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: case_substitution @ any (throw: false) 
func E_Optimize_case_substitution_any (x EID,_Zvar EID,vsub EID) EID { 
    return /*(sm for case_substitution @ any= any)*/ F_Optimize_case_substitution_any(ANY(x),To_Variable(OBJ(_Zvar)),To_Variable(OBJ(vsub)) ).ToEID()} 
  
// case is treated like a macro and vanishes into a large if.
// the last line is a trap for code generated by the logic compiler.
// in CLAIRE 4 we substitute the variables in the branches with a properly typed variable (borrowed from c_code@For)
// note: the range sets are now useless and have been removed
/* {1} OPT.The go function for: c_code(self:Case,s:class) [] */
func F_c_code_Case (self *Language.Case ,s *ClaireClass ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var _Zvar *ClaireAny   = self.ClaireVar
      /* noccur = 10 */
      /* Let:3 */{ 
        var _Ztype *ClaireAny  
        /* noccur = 1 */
        if (_Zvar.Isa.IsIn(C_Variable) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var g0858 *ClaireVariable   = To_Variable(_Zvar)
            /* noccur = 1 */
            _Ztype = Core.F_get_property(C_range,ToObject(g0858.Id()))
            /* Let-5 */} 
          } else {
          _Ztype = C_any.Id()
          /* If-4 */} 
        /* Let:4 */{ 
          var l *ClaireList   = self.Args.Copy()
          /* noccur = 14 */
          /* Let:5 */{ 
            var utype *ClaireAny   = CEMPTY.Id()
            /* noccur = 3 */
            /* Let:6 */{ 
              var ctest1 *ClaireAny  
              /* noccur = 1 */
              var ctest1_try08617 EID 
              /* Let:7 */{ 
                var g0862UU *Language.Call  
                /* noccur = 1 */
                /* Let:8 */{ 
                  var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                  /* noccur = 5 */
                  _CL_obj.Selector = ToProperty(C__Z.Id())
                  _CL_obj.Args = MakeConstantList(_Zvar,l.At(1-1))
                  g0862UU = _CL_obj
                  /* Let-8 */} 
                ctest1_try08617 = F_Optimize_c_boolean_any(g0862UU.Id())
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (ctest1-Result) */
              if ErrorIn(ctest1_try08617) {Result = ctest1_try08617
              } else {
              ctest1 = ANY(ctest1_try08617)
              /* Let:7 */{ 
                var rep *Language.If  
                /* noccur = 2 */
                var rep_try08638 EID 
                /* Let:8 */{ 
                  var g0864UU *ClaireAny  
                  /* noccur = 1 */
                  var g0864UU_try08669 EID 
                  g0864UU_try08669 = Core.F_CALL(C_c_code,ARGS(F_Optimize_case_branch_any(l.At(2-1),_Zvar,ToType(l.At(1-1))).ToEID(),EID{s.Id(),0}))
                  /* ERROR PROTECTION INSERTED (g0864UU-rep_try08638) */
                  if ErrorIn(g0864UU_try08669) {rep_try08638 = g0864UU_try08669
                  } else {
                  g0864UU = ANY(g0864UU_try08669)
                  /* Let:9 */{ 
                    var g0865UU *ClaireAny  
                    /* noccur = 1 */
                    var g0865UU_try086710 EID 
                    g0865UU_try086710 = Core.F_CALL(C_c_code,ARGS(EID{CFALSE.Id(),0},EID{s.Id(),0}))
                    /* ERROR PROTECTION INSERTED (g0865UU-rep_try08638) */
                    if ErrorIn(g0865UU_try086710) {rep_try08638 = g0865UU_try086710
                    } else {
                    g0865UU = ANY(g0865UU_try086710)
                    rep_try08638 = Language.C_If.Make(ctest1,g0864UU,g0865UU).ToEID()
                    }
                    /* Let-9 */} 
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (rep-Result) */
                if ErrorIn(rep_try08638) {Result = rep_try08638
                } else {
                rep = Language.To_If(OBJ(rep_try08638))
                /* Let:8 */{ 
                  var pointer *Language.If   = rep
                  /* noccur = 6 */
                  l = l.Skip(2)
                  Result= EID{CFALSE.Id(),0}
                  for (l.Length() > 0) /* while:9 */{ 
                    var void_try10 EID 
                    _ = void_try10
                    { 
                    utype = Core.F_U_type(ToType(utype),ToType(l.At(1-1))).Id()
                    if (ToType(_Ztype).Included(ToType(utype)) == CTRUE) /* If:10 */{ 
                      /* update:11 */{ 
                        var va_arg1 *Language.If  
                        var va_arg2 *ClaireAny  
                        va_arg1 = pointer
                        var va_arg2_try086812 EID 
                        va_arg2_try086812 = Core.F_CALL(C_c_code,ARGS(F_Optimize_case_branch_any(l.At(2-1),_Zvar,ToType(l.At(1-1))).ToEID(),EID{s.Id(),0}))
                        /* ERROR PROTECTION INSERTED (va_arg2-void_try10) */
                        if ErrorIn(va_arg2_try086812) {void_try10 = va_arg2_try086812
                        } else {
                        va_arg2 = ANY(va_arg2_try086812)
                        /* ---------- now we compile update iClaire/other(va_arg1) := va_arg2 ------- */
                        va_arg1.Other = va_arg2
                        void_try10 = va_arg2.ToEID()
                        }
                        /* update-11 */} 
                      /* ERROR PROTECTION INSERTED (void_try10-void_try10) */
                      if ErrorIn(void_try10) {Result = void_try10
                      break
                      } else {
                       /*v = Result, s =EID*/
Result = EID{CTRUE.Id(),0}
                      break
                      }
                      } else {
                      /* Let:11 */{ 
                        var ctest *ClaireAny  
                        /* noccur = 1 */
                        var ctest_try086912 EID 
                        /* Let:12 */{ 
                          var g0870UU *Language.Call  
                          /* noccur = 1 */
                          /* Let:13 */{ 
                            var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                            /* noccur = 5 */
                            _CL_obj.Selector = ToProperty(C__Z.Id())
                            _CL_obj.Args = MakeConstantList(_Zvar,l.At(1-1))
                            g0870UU = _CL_obj
                            /* Let-13 */} 
                          ctest_try086912 = F_Optimize_c_boolean_any(g0870UU.Id())
                          /* Let-12 */} 
                        /* ERROR PROTECTION INSERTED (ctest-void_try10) */
                        if ErrorIn(ctest_try086912) {void_try10 = ctest_try086912
                        } else {
                        ctest = ANY(ctest_try086912)
                        /* update:12 */{ 
                          var va_arg1 *Language.If  
                          var va_arg2 *ClaireAny  
                          va_arg1 = pointer
                          var va_arg2_try087113 EID 
                          /* Let:13 */{ 
                            var g0872UU *ClaireAny  
                            /* noccur = 1 */
                            var g0872UU_try087414 EID 
                            g0872UU_try087414 = Core.F_CALL(C_c_code,ARGS(F_Optimize_case_branch_any(l.At(2-1),_Zvar,ToType(l.At(1-1))).ToEID(),EID{s.Id(),0}))
                            /* ERROR PROTECTION INSERTED (g0872UU-va_arg2_try087113) */
                            if ErrorIn(g0872UU_try087414) {va_arg2_try087113 = g0872UU_try087414
                            } else {
                            g0872UU = ANY(g0872UU_try087414)
                            /* Let:14 */{ 
                              var g0873UU *ClaireAny  
                              /* noccur = 1 */
                              var g0873UU_try087515 EID 
                              g0873UU_try087515 = Core.F_CALL(C_c_code,ARGS(EID{CFALSE.Id(),0},EID{s.Id(),0}))
                              /* ERROR PROTECTION INSERTED (g0873UU-va_arg2_try087113) */
                              if ErrorIn(g0873UU_try087515) {va_arg2_try087113 = g0873UU_try087515
                              } else {
                              g0873UU = ANY(g0873UU_try087515)
                              va_arg2_try087113 = Language.C_If.Make(ctest,g0872UU,g0873UU).ToEID()
                              }
                              /* Let-14 */} 
                            }
                            /* Let-13 */} 
                          /* ERROR PROTECTION INSERTED (va_arg2-void_try10) */
                          if ErrorIn(va_arg2_try087113) {void_try10 = va_arg2_try087113
                          } else {
                          va_arg2 = ANY(va_arg2_try087113)
                          /* ---------- now we compile update iClaire/other(va_arg1) := va_arg2 ------- */
                          va_arg1.Other = va_arg2
                          void_try10 = va_arg2.ToEID()
                          }
                          /* update-12 */} 
                        /* ERROR PROTECTION INSERTED (void_try10-void_try10) */
                        if ErrorIn(void_try10) {Result = void_try10
                        break
                        } else {
                        pointer = Language.To_If(pointer.Other)
                        void_try10 = EID{pointer.Id(),0}
                        }
                        }
                        /* Let-11 */} 
                      /* If-10 */} 
                    /* ERROR PROTECTION INSERTED (void_try10-void_try10) */
                    if ErrorIn(void_try10) {Result = void_try10
                    break
                    } else {
                    l = l.Skip(2)
                    }
                    /* while-9 */} 
                  }
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  var g0876I *ClaireBoolean  
                  if (_Zvar.Isa.IsIn(Language.C_Definition) == CTRUE) /* If:9 */{ 
                    /* Let:10 */{ 
                      var g0860 *Language.Definition   = Language.To_Definition(_Zvar)
                      /* noccur = 1 */
                      g0876I = g0860.Arg.Isa.IsIn(C_exception)
                      /* Let-10 */} 
                    } else {
                    g0876I = CFALSE
                    /* If-9 */} 
                  if (g0876I == CTRUE) /* If:9 */{ 
                    Result = _Zvar.ToEID()
                    } else {
                    Result = EID{rep.Id(),0}
                    /* If-9 */} 
                  }
                  /* Let-8 */} 
                }
                /* Let-7 */} 
              }
              /* Let-6 */} 
            /* Let-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_code @ Case (throw: true) 
func E_c_code_Case (self EID,s EID) EID { 
    return /*(sm for c_code @ Case= EID)*/ F_c_code_Case(Language.To_Case(OBJ(self)),ToClass(OBJ(s)) )} 
  
// member_of is treated like a macro and vanishes into a large if.
//_____________________ Block structure________________________
/* {1} OPT.The go function for: c_type(self:Do) [] */
func F_c_type_Do (self *Language.Do ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var g0877UU *ClaireAny  
      /* noccur = 1 */
      var g0877UU_try08783 EID 
      g0877UU_try08783 = Core.F_last_list(self.Args)
      /* ERROR PROTECTION INSERTED (g0877UU-Result) */
      if ErrorIn(g0877UU_try08783) {Result = g0877UU_try08783
      } else {
      g0877UU = ANY(g0877UU_try08783)
      Result = Core.F_CALL(C_c_type,ARGS(g0877UU.ToEID()))
      }
      /* Let-2 */} 
    return RangeCheck(ToType(C_type.Id()),Result)} 
  
// The EID go function for: c_type @ Do (throw: true) 
func E_c_type_Do (self EID) EID { 
    return /*(sm for c_type @ Do= EID)*/ F_c_type_Do(Language.To_Do(OBJ(self)) )} 
  
/* {1} OPT.The go function for: c_code(self:Do,s:class) [] */
func F_c_code_Do (self *Language.Do ,s *ClaireClass ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var _CL_obj *Language.Do   = Language.To_Do(new(Language.Do).Is(Language.C_Do))
      /* noccur = 3 */
      /* update:3 */{ 
        var va_arg1 *Language.Do  
        var va_arg2 *ClaireList  
        va_arg1 = _CL_obj
        var va_arg2_try08794 EID 
        /* Let:4 */{ 
          var m int  = self.Args.Length()
          /* noccur = 1 */
          /* Let:5 */{ 
            var n int  = 0
            /* noccur = 3 */
            /* Iteration:6 */{ 
              var v_list6 *ClaireList  
              var x *ClaireAny  
              var v_local6 *ClaireAny  
              v_list6 = self.Args
              va_arg2_try08794 = EID{CreateList(ToType(CEMPTY.Id()),v_list6.Length()).Id(),0}
              for CLcount := 0; CLcount < v_list6.Length(); CLcount++{ 
                x = v_list6.At(CLcount)
                var v_local6_try08808 EID 
                n = (n+1)
                /* Let:8 */{ 
                  var g0881UU *ClaireClass  
                  /* noccur = 1 */
                  if (n == m) /* If:9 */{ 
                    g0881UU = s
                    } else {
                    g0881UU = C_void
                    /* If-9 */} 
                  v_local6_try08808 = Core.F_CALL(C_c_code,ARGS(x.ToEID(),EID{g0881UU.Id(),0}))
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (v_local6_try08808-v_local6_try08808) */
                if ErrorIn(v_local6_try08808) {va_arg2_try08794 = v_local6_try08808
                break
                } else {
                }
                {
                v_local6 = ANY(v_local6_try08808)
                ToList(OBJ(va_arg2_try08794)).PutAt(CLcount,v_local6)
                } 
              }
              /* Iteration-6 */} 
            /* Let-5 */} 
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (va_arg2-Result) */
        if ErrorIn(va_arg2_try08794) {Result = va_arg2_try08794
        } else {
        va_arg2 = ToList(OBJ(va_arg2_try08794))
        /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
        va_arg1.Args = va_arg2
        Result = EID{va_arg2.Id(),0}
        }
        /* update-3 */} 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = EID{_CL_obj.Id(),0}
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_code @ Do (throw: true) 
func E_c_code_Do (self EID,s EID) EID { 
    return /*(sm for c_code @ Do= EID)*/ F_c_code_Do(Language.To_Do(OBJ(self)),ToClass(OBJ(s)) )} 
  
// ----------------------- LET -----------------------------------
// we make a range inference
//
/* {1} OPT.The go function for: c_type(self:Let) [] */
func F_c_type_Let (self *Language.Let ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var g0882UU *ClaireType  
      /* noccur = 1 */
      var g0882UU_try08833 EID 
      g0882UU_try08833 = Core.F_CALL(C_c_type,ARGS(self.Value.ToEID()))
      /* ERROR PROTECTION INSERTED (g0882UU-Result) */
      if ErrorIn(g0882UU_try08833) {Result = g0882UU_try08833
      } else {
      g0882UU = ToType(OBJ(g0882UU_try08833))
      Result = F_Optimize_range_infers_Variable(self.ClaireVar,g0882UU)
      }
      /* Let-2 */} 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    Result = Core.F_CALL(C_c_type,ARGS(self.Arg.ToEID()))
    }
    return RangeCheck(ToType(C_type.Id()),Result)} 
  
// The EID go function for: c_type @ Let (throw: true) 
func E_c_type_Let (self EID) EID { 
    return /*(sm for c_type @ Let= EID)*/ F_c_type_Let(Language.To_Let(OBJ(self)) )} 
  
// works also for Let+ / Let*
/* {1} OPT.The go function for: c_code(self:Let,s:class) [] */
func F_c_code_Let (self *Language.Let ,s *ClaireClass ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var _Zv *ClaireAny   = self.Value
      /* noccur = 4 */
      /* Let:3 */{ 
        var _Ztype *ClaireType  
        /* noccur = 3 */
        var _Ztype_try08844 EID 
        /* Let:4 */{ 
          var g0885UU *ClaireType  
          /* noccur = 1 */
          var g0885UU_try08865 EID 
          g0885UU_try08865 = Core.F_CALL(C_c_type,ARGS(_Zv.ToEID()))
          /* ERROR PROTECTION INSERTED (g0885UU-_Ztype_try08844) */
          if ErrorIn(g0885UU_try08865) {_Ztype_try08844 = g0885UU_try08865
          } else {
          g0885UU = ToType(OBJ(g0885UU_try08865))
          _Ztype_try08844 = EID{F_Optimize_ptype_type(g0885UU).Id(),0}
          }
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (_Ztype-Result) */
        if ErrorIn(_Ztype_try08844) {Result = _Ztype_try08844
        } else {
        _Ztype = ToType(OBJ(_Ztype_try08844))
        Result = F_Optimize_range_infers_Variable(self.ClaireVar,_Ztype)
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        if (_Ztype.Included(self.ClaireVar.Range) != CTRUE) /* If:4 */{ 
          var _Zv_try08875 EID 
          _Zv_try08875 = F_Optimize_c_warn_Variable(self.ClaireVar,_Zv,_Ztype)
          /* ERROR PROTECTION INSERTED (_Zv-Result) */
          if ErrorIn(_Zv_try08875) {Result = _Zv_try08875
          } else {
          _Zv = ANY(_Zv_try08875)
          Result = _Zv.ToEID()
          }
          } else {
          Result = EID{CFALSE.Id(),0}
          /* If-4 */} 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        /* Let:4 */{ 
          var x *Language.Let  
          /* noccur = 2 */
          var x_try08885 EID 
          /* Let:5 */{ 
            var _CL_obj *Language.Let   = Language.To_Let(new(Language.Let).Is(Language.C_Let))
            /* noccur = 6 */
            _CL_obj.ClaireVar = self.ClaireVar
            /* update:6 */{ 
              var va_arg1 *Language.Let  
              var va_arg2 *ClaireAny  
              va_arg1 = _CL_obj
              var va_arg2_try08897 EID 
              va_arg2_try08897 = F_Compile_c_strict_code_any(_Zv,F_Compile_psort_any(self.ClaireVar.Range.Id()))
              /* ERROR PROTECTION INSERTED (va_arg2-x_try08885) */
              if ErrorIn(va_arg2_try08897) {x_try08885 = va_arg2_try08897
              } else {
              va_arg2 = ANY(va_arg2_try08897)
              /* ---------- now we compile update value(va_arg1) := va_arg2 ------- */
              va_arg1.Value = va_arg2
              x_try08885 = va_arg2.ToEID()
              }
              /* update-6 */} 
            /* ERROR PROTECTION INSERTED (x_try08885-x_try08885) */
            if !ErrorIn(x_try08885) {
            /* update:6 */{ 
              var va_arg1 *Language.Let  
              var va_arg2 *ClaireAny  
              va_arg1 = _CL_obj
              var va_arg2_try08907 EID 
              va_arg2_try08907 = Core.F_CALL(C_c_code,ARGS(self.Arg.ToEID(),EID{s.Id(),0}))
              /* ERROR PROTECTION INSERTED (va_arg2-x_try08885) */
              if ErrorIn(va_arg2_try08907) {x_try08885 = va_arg2_try08907
              } else {
              va_arg2 = ANY(va_arg2_try08907)
              /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
              va_arg1.Arg = va_arg2
              x_try08885 = va_arg2.ToEID()
              }
              /* update-6 */} 
            /* ERROR PROTECTION INSERTED (x_try08885-x_try08885) */
            if !ErrorIn(x_try08885) {
            x_try08885 = EID{_CL_obj.Id(),0}
            }}
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (x-Result) */
          if ErrorIn(x_try08885) {Result = x_try08885
          } else {
          x = Language.To_Let(OBJ(x_try08885))
          x.Isa = self.Isa
          Result = EID{x.Id(),0}
          }
          /* Let-4 */} 
        }}
        }
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_code @ Let (throw: true) 
func E_c_code_Let (self EID,s EID) EID { 
    return /*(sm for c_code @ Let= EID)*/ F_c_code_Let(Language.To_Let(OBJ(self)),ToClass(OBJ(s)) )} 
  
// type inference for When is more subtle
/* {1} OPT.The go function for: c_type(self:When) [] */
func F_c_type_When (self *Language.When ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var _Zv *ClaireAny   = self.Value
      /* noccur = 2 */
      /* Let:3 */{ 
        var v *ClaireVariable   = self.ClaireVar
        /* noccur = 1 */
        /* Let:4 */{ 
          var d *ClaireAny  
          /* noccur = 2 */
          var d_try08915 EID 
          d_try08915 = F_Optimize_daccess_any(_Zv,CTRUE)
          /* ERROR PROTECTION INSERTED (d-Result) */
          if ErrorIn(d_try08915) {Result = d_try08915
          } else {
          d = ANY(d_try08915)
          /* Let:5 */{ 
            var _Ztype *ClaireAny  
            /* noccur = 4 */
            var _Ztype_try08926 EID 
            if (d != CNULL) /* If:6 */{ 
              _Ztype_try08926 = Core.F_CALL(C_c_type,ARGS(d.ToEID()))
              } else {
              _Ztype_try08926 = Core.F_CALL(C_c_type,ARGS(_Zv.ToEID()))
              /* If-6 */} 
            /* ERROR PROTECTION INSERTED (_Ztype-Result) */
            if ErrorIn(_Ztype_try08926) {Result = _Ztype_try08926
            } else {
            _Ztype = ANY(_Ztype_try08926)
            var g0893I *ClaireBoolean  
            var g0893I_try08946 EID 
            g0893I_try08946 = F_Optimize_extended_ask_type(ToType(_Ztype))
            /* ERROR PROTECTION INSERTED (g0893I-Result) */
            if ErrorIn(g0893I_try08946) {Result = g0893I_try08946
            } else {
            g0893I = ToBoolean(OBJ(g0893I_try08946))
            if (g0893I == CTRUE) /* If:6 */{ 
              _Ztype = ANY(Core.F_CALL(C_mClaire_t1,ARGS(_Ztype.ToEID())))
              Result = _Ztype.ToEID()
              } else {
              Result = EID{CFALSE.Id(),0}
              /* If-6 */} 
            }
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            Result = F_Optimize_range_infers_Variable(v,ToType(_Ztype))
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            /* Let:6 */{ 
              var g0895UU *ClaireType  
              /* noccur = 1 */
              var g0895UU_try08977 EID 
              g0895UU_try08977 = Core.F_CALL(C_c_type,ARGS(self.Arg.ToEID()))
              /* ERROR PROTECTION INSERTED (g0895UU-Result) */
              if ErrorIn(g0895UU_try08977) {Result = g0895UU_try08977
              } else {
              g0895UU = ToType(OBJ(g0895UU_try08977))
              /* Let:7 */{ 
                var g0896UU *ClaireType  
                /* noccur = 1 */
                var g0896UU_try08988 EID 
                g0896UU_try08988 = Core.F_CALL(C_c_type,ARGS(self.Other.ToEID()))
                /* ERROR PROTECTION INSERTED (g0896UU-Result) */
                if ErrorIn(g0896UU_try08988) {Result = g0896UU_try08988
                } else {
                g0896UU = ToType(OBJ(g0896UU_try08988))
                Result = EID{Core.F_U_type(g0895UU,g0896UU).Id(),0}
                }
                /* Let-7 */} 
              }
              /* Let-6 */} 
            }}
            }
            /* Let-5 */} 
          }
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_type @ When (throw: true) 
func E_c_type_When (self EID) EID { 
    return /*(sm for c_type @ When= EID)*/ F_c_type_When(Language.To_When(OBJ(self)) )} 
  
// A When is macroexpanded into one/two Let
/* {1} OPT.The go function for: c_code(self:When,s:class) [] */
func F_c_code_When (self *Language.When ,s *ClaireClass ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var _Zv *ClaireAny   = self.Value
      /* noccur = 4 */
      /* Let:3 */{ 
        var v *ClaireVariable   = self.ClaireVar
        /* noccur = 9 */
        /* Let:4 */{ 
          var d *ClaireAny  
          /* noccur = 5 */
          var d_try08995 EID 
          d_try08995 = F_Optimize_daccess_any(_Zv,CTRUE)
          /* ERROR PROTECTION INSERTED (d-Result) */
          if ErrorIn(d_try08995) {Result = d_try08995
          } else {
          d = ANY(d_try08995)
          /* Let:5 */{ 
            var v2 *ClaireVariable   = F_Compile_Variable_I_symbol(F_append_symbol(self.ClaireVar.Pname,MakeString("test").Id()),self.ClaireVar.Index,C_any.Id())
            /* noccur = 3 */
            /* Let:6 */{ 
              var _Ztype *ClaireAny  
              /* noccur = 6 */
              var _Ztype_try09007 EID 
              if (d != CNULL) /* If:7 */{ 
                _Ztype_try09007 = Core.F_CALL(C_c_type,ARGS(d.ToEID()))
                } else {
                _Ztype_try09007 = Core.F_CALL(C_c_type,ARGS(_Zv.ToEID()))
                /* If-7 */} 
              /* ERROR PROTECTION INSERTED (_Ztype-Result) */
              if ErrorIn(_Ztype_try09007) {Result = _Ztype_try09007
              } else {
              _Ztype = ANY(_Ztype_try09007)
              var g0901I *ClaireBoolean  
              var g0901I_try09027 EID 
              g0901I_try09027 = F_Optimize_extended_ask_type(ToType(_Ztype))
              /* ERROR PROTECTION INSERTED (g0901I-Result) */
              if ErrorIn(g0901I_try09027) {Result = g0901I_try09027
              } else {
              g0901I = ToBoolean(OBJ(g0901I_try09027))
              if (g0901I == CTRUE) /* If:7 */{ 
                _Ztype = ANY(Core.F_CALL(C_mClaire_t1,ARGS(_Ztype.ToEID())))
                Result = _Ztype.ToEID()
                } else {
                Result = EID{CFALSE.Id(),0}
                /* If-7 */} 
              }
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              Result = F_Optimize_range_infers_Variable(v,ToType(_Ztype))
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              var g0903I *ClaireBoolean  
              var g0903I_try09047 EID 
              /* and:7 */{ 
                var v_and7 *ClaireBoolean  
                
                v_and7 = Core.F_known_ask_any(d)
                if (v_and7 == CFALSE) {g0903I_try09047 = EID{CFALSE.Id(),0}
                } else /* arg:8 */{ 
                  var v_and7_try09059 EID 
                  /* Let:9 */{ 
                    var g0906UU *ClaireBoolean  
                    /* noccur = 1 */
                    var g0906UU_try090710 EID 
                    g0906UU_try090710 = F_Optimize_extended_ask_type(ToType(OBJ(Core.F_CALL(C_range,ARGS(Core.F_CALL(C_selector,ARGS(d.ToEID())))))))
                    /* ERROR PROTECTION INSERTED (g0906UU-v_and7_try09059) */
                    if ErrorIn(g0906UU_try090710) {v_and7_try09059 = g0906UU_try090710
                    } else {
                    g0906UU = ToBoolean(OBJ(g0906UU_try090710))
                    v_and7_try09059 = EID{g0906UU.Not.Id(),0}
                    }
                    /* Let-9 */} 
                  /* ERROR PROTECTION INSERTED (v_and7-g0903I_try09047) */
                  if ErrorIn(v_and7_try09059) {g0903I_try09047 = v_and7_try09059
                  } else {
                  v_and7 = ToBoolean(OBJ(v_and7_try09059))
                  if (v_and7 == CFALSE) {g0903I_try09047 = EID{CFALSE.Id(),0}
                  } else /* arg:9 */{ 
                    g0903I_try09047 = EID{CTRUE.Id(),0}/* arg-9 */} 
                  /* arg-8 */} 
                }
                /* and-7 */} 
              /* ERROR PROTECTION INSERTED (g0903I-Result) */
              if ErrorIn(g0903I_try09047) {Result = g0903I_try09047
              } else {
              g0903I = ToBoolean(OBJ(g0903I_try09047))
              if (g0903I == CTRUE) /* If:7 */{ 
                /* Let:8 */{ 
                  var _CL_obj *Language.Let   = Language.To_Let(new(Language.Let).Is(Language.C_Let))
                  /* noccur = 12 */
                  _CL_obj.ClaireVar = v
                  _CL_obj.Value = d
                  /* update:9 */{ 
                    var va_arg1 *Language.Let  
                    var va_arg2 *ClaireAny  
                    va_arg1 = _CL_obj
                    var va_arg2_try090810 EID 
                    /* Let:10 */{ 
                      var g0909UU *Language.CallMethod2  
                      /* noccur = 1 */
                      var g0909UU_try091211 EID 
                      /* Let:11 */{ 
                        var _CL_obj *Language.CallMethod2   = Language.To_CallMethod2(new(Language.CallMethod2).Is(Language.C_Call_method2))
                        /* noccur = 5 */
                        _CL_obj.Arg = ToMethod(Core.F__at_property1(Core.C_identical_ask,C_any).Id())
                        /* update:12 */{ 
                          var va_arg1 *Language.CallMethod  
                          var va_arg2 *ClaireList  
                          va_arg1 = Language.To_CallMethod(_CL_obj.Id())
                          var va_arg2_try091313 EID 
                          /* Construct:13 */{ 
                            var v_bag_arg *ClaireAny  
                            va_arg2_try091313= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                            ToList(OBJ(va_arg2_try091313)).AddFast(v.Id())
                            var v_bag_arg_try091414 EID 
                            /* Let:14 */{ 
                              var g0915UU *ClaireAny  
                              /* noccur = 1 */
                              var g0915UU_try091615 EID 
                              g0915UU_try091615 = Core.F_CALL(C_Compile_c_sort,ARGS(EID{v.Id(),0}))
                              /* ERROR PROTECTION INSERTED (g0915UU-v_bag_arg_try091414) */
                              if ErrorIn(g0915UU_try091615) {v_bag_arg_try091414 = g0915UU_try091615
                              } else {
                              g0915UU = ANY(g0915UU_try091615)
                              v_bag_arg_try091414 = Core.F_CALL(C_c_code,ARGS(EID{CNULL,0},g0915UU.ToEID()))
                              }
                              /* Let-14 */} 
                            /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try091313) */
                            if ErrorIn(v_bag_arg_try091414) {va_arg2_try091313 = v_bag_arg_try091414
                            } else {
                            v_bag_arg = ANY(v_bag_arg_try091414)
                            ToList(OBJ(va_arg2_try091313)).AddFast(v_bag_arg)}
                            /* Construct-13 */} 
                          /* ERROR PROTECTION INSERTED (va_arg2-g0909UU_try091211) */
                          if ErrorIn(va_arg2_try091313) {g0909UU_try091211 = va_arg2_try091313
                          } else {
                          va_arg2 = ToList(OBJ(va_arg2_try091313))
                          /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                          va_arg1.Args = va_arg2
                          g0909UU_try091211 = EID{va_arg2.Id(),0}
                          }
                          /* update-12 */} 
                        /* ERROR PROTECTION INSERTED (g0909UU_try091211-g0909UU_try091211) */
                        if !ErrorIn(g0909UU_try091211) {
                        g0909UU_try091211 = EID{_CL_obj.Id(),0}
                        }
                        /* Let-11 */} 
                      /* ERROR PROTECTION INSERTED (g0909UU-va_arg2_try090810) */
                      if ErrorIn(g0909UU_try091211) {va_arg2_try090810 = g0909UU_try091211
                      } else {
                      g0909UU = Language.To_CallMethod2(OBJ(g0909UU_try091211))
                      /* Let:11 */{ 
                        var g0910UU *ClaireAny  
                        /* noccur = 1 */
                        var g0910UU_try091712 EID 
                        g0910UU_try091712 = Core.F_CALL(C_c_code,ARGS(self.Other.ToEID(),EID{s.Id(),0}))
                        /* ERROR PROTECTION INSERTED (g0910UU-va_arg2_try090810) */
                        if ErrorIn(g0910UU_try091712) {va_arg2_try090810 = g0910UU_try091712
                        } else {
                        g0910UU = ANY(g0910UU_try091712)
                        /* Let:12 */{ 
                          var g0911UU *ClaireAny  
                          /* noccur = 1 */
                          var g0911UU_try091813 EID 
                          g0911UU_try091813 = Core.F_CALL(C_c_code,ARGS(self.Arg.ToEID(),EID{s.Id(),0}))
                          /* ERROR PROTECTION INSERTED (g0911UU-va_arg2_try090810) */
                          if ErrorIn(g0911UU_try091813) {va_arg2_try090810 = g0911UU_try091813
                          } else {
                          g0911UU = ANY(g0911UU_try091813)
                          va_arg2_try090810 = Language.C_If.Make(g0909UU.Id(),g0910UU,g0911UU).ToEID()
                          }
                          /* Let-12 */} 
                        }
                        /* Let-11 */} 
                      }
                      /* Let-10 */} 
                    /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                    if ErrorIn(va_arg2_try090810) {Result = va_arg2_try090810
                    } else {
                    va_arg2 = ANY(va_arg2_try090810)
                    /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                    va_arg1.Arg = va_arg2
                    Result = va_arg2.ToEID()
                    }
                    /* update-9 */} 
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  Result = EID{_CL_obj.Id(),0}
                  }
                  /* Let-8 */} 
                } else {
                var g0919I *ClaireBoolean  
                var g0919I_try09208 EID 
                /* and:8 */{ 
                  var v_and8 *ClaireBoolean  
                  
                  var v_and8_try09219 EID 
                  /* Let:9 */{ 
                    var g0922UU *ClaireAny  
                    /* noccur = 1 */
                    var g0922UU_try092310 EID 
                    g0922UU_try092310 = Core.F_CALL(C_Compile_c_sort,ARGS(EID{v.Id(),0}))
                    /* ERROR PROTECTION INSERTED (g0922UU-v_and8_try09219) */
                    if ErrorIn(g0922UU_try092310) {v_and8_try09219 = g0922UU_try092310
                    } else {
                    g0922UU = ANY(g0922UU_try092310)
                    v_and8_try09219 = EID{Equal(g0922UU,C_any.Id()).Id(),0}
                    }
                    /* Let-9 */} 
                  /* ERROR PROTECTION INSERTED (v_and8-g0919I_try09208) */
                  if ErrorIn(v_and8_try09219) {g0919I_try09208 = v_and8_try09219
                  } else {
                  v_and8 = ToBoolean(OBJ(v_and8_try09219))
                  if (v_and8 == CFALSE) {g0919I_try09208 = EID{CFALSE.Id(),0}
                  } else /* arg:9 */{ 
                    v_and8 = ToType(_Ztype).Included(v.Range)
                    if (v_and8 == CFALSE) {g0919I_try09208 = EID{CFALSE.Id(),0}
                    } else /* arg:10 */{ 
                      v_and8 = F__sup_equal_integer(C_compiler.Safety,3)
                      if (v_and8 == CFALSE) {g0919I_try09208 = EID{CFALSE.Id(),0}
                      } else /* arg:11 */{ 
                        g0919I_try09208 = EID{CTRUE.Id(),0}/* arg-11 */} 
                      /* arg-10 */} 
                    /* arg-9 */} 
                  }
                  /* and-8 */} 
                /* ERROR PROTECTION INSERTED (g0919I-Result) */
                if ErrorIn(g0919I_try09208) {Result = g0919I_try09208
                } else {
                g0919I = ToBoolean(OBJ(g0919I_try09208))
                if (g0919I == CTRUE) /* If:8 */{ 
                  /* Let:9 */{ 
                    var g0924UU *Language.Let  
                    /* noccur = 1 */
                    /* Let:10 */{ 
                      var _CL_obj *Language.Let   = Language.To_Let(new(Language.Let).Is(Language.C_Let))
                      /* noccur = 12 */
                      _CL_obj.ClaireVar = v
                      _CL_obj.Value = _Zv
                      /* update:11 */{ 
                        var va_arg1 *Language.Let  
                        var va_arg2 *ClaireAny  
                        va_arg1 = _CL_obj
                        /* Let:12 */{ 
                          var g0925UU *Language.Call  
                          /* noccur = 1 */
                          /* Let:13 */{ 
                            var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                            /* noccur = 5 */
                            _CL_obj.Selector = ToProperty(Core.C__I_equal.Id())
                            _CL_obj.Args = MakeConstantList(v.Id(),CNULL)
                            g0925UU = _CL_obj
                            /* Let-13 */} 
                          va_arg2 = Language.C_If.Make(g0925UU.Id(),self.Arg,self.Other)
                          /* Let-12 */} 
                        /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                        va_arg1.Arg = va_arg2
                        /* update-11 */} 
                      g0924UU = _CL_obj
                      /* Let-10 */} 
                    Result = Core.F_CALL(C_c_code,ARGS(EID{g0924UU.Id(),0},EID{s.Id(),0}))
                    /* Let-9 */} 
                  } else {
                  /* Let:9 */{ 
                    var g0926UU *Language.Let  
                    /* noccur = 1 */
                    var g0926UU_try092710 EID 
                    /* Let:10 */{ 
                      var _CL_obj *Language.Let   = Language.To_Let(new(Language.Let).Is(Language.C_Let))
                      /* noccur = 24 */
                      _CL_obj.ClaireVar = v2
                      _CL_obj.Value = _Zv
                      /* update:11 */{ 
                        var va_arg1 *Language.Let  
                        var va_arg2 *ClaireAny  
                        va_arg1 = _CL_obj
                        var va_arg2_try092812 EID 
                        /* Let:12 */{ 
                          var g0929UU *Language.Call  
                          /* noccur = 1 */
                          /* Let:13 */{ 
                            var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                            /* noccur = 5 */
                            _CL_obj.Selector = ToProperty(Core.C__I_equal.Id())
                            _CL_obj.Args = MakeConstantList(v2.Id(),CNULL)
                            g0929UU = _CL_obj
                            /* Let-13 */} 
                          /* Let:13 */{ 
                            var g0930UU *Language.Let  
                            /* noccur = 1 */
                            /* Let:14 */{ 
                              var _CL_obj *Language.Let   = Language.To_Let(new(Language.Let).Is(Language.C_Let))
                              /* noccur = 11 */
                              _CL_obj.ClaireVar = v
                              /* update:15 */{ 
                                var va_arg1 *Language.Let  
                                var va_arg2 *ClaireAny  
                                va_arg1 = _CL_obj
                                /* Let:16 */{ 
                                  var _CL_obj *Language.Cast   = Language.To_Cast(new(Language.Cast).Is(Language.C_Cast))
                                  /* noccur = 4 */
                                  _CL_obj.Arg = v2.Id()
                                  _CL_obj.SetArg = ToType(_Ztype)
                                  va_arg2 = _CL_obj.Id()
                                  /* Let-16 */} 
                                /* ---------- now we compile update value(va_arg1) := va_arg2 ------- */
                                va_arg1.Value = va_arg2
                                /* update-15 */} 
                              _CL_obj.Arg = self.Arg
                              g0930UU = _CL_obj
                              /* Let-14 */} 
                            /* Let:14 */{ 
                              var g0931UU *ClaireAny  
                              /* noccur = 1 */
                              var g0931UU_try093215 EID 
                              g0931UU_try093215 = Core.F_CALL(C_c_code,ARGS(self.Other.ToEID(),EID{s.Id(),0}))
                              /* ERROR PROTECTION INSERTED (g0931UU-va_arg2_try092812) */
                              if ErrorIn(g0931UU_try093215) {va_arg2_try092812 = g0931UU_try093215
                              } else {
                              g0931UU = ANY(g0931UU_try093215)
                              va_arg2_try092812 = Language.C_If.Make(g0929UU.Id(),g0930UU.Id(),g0931UU).ToEID()
                              }
                              /* Let-14 */} 
                            /* Let-13 */} 
                          /* Let-12 */} 
                        /* ERROR PROTECTION INSERTED (va_arg2-g0926UU_try092710) */
                        if ErrorIn(va_arg2_try092812) {g0926UU_try092710 = va_arg2_try092812
                        } else {
                        va_arg2 = ANY(va_arg2_try092812)
                        /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                        va_arg1.Arg = va_arg2
                        g0926UU_try092710 = va_arg2.ToEID()
                        }
                        /* update-11 */} 
                      /* ERROR PROTECTION INSERTED (g0926UU_try092710-g0926UU_try092710) */
                      if !ErrorIn(g0926UU_try092710) {
                      g0926UU_try092710 = EID{_CL_obj.Id(),0}
                      }
                      /* Let-10 */} 
                    /* ERROR PROTECTION INSERTED (g0926UU-Result) */
                    if ErrorIn(g0926UU_try092710) {Result = g0926UU_try092710
                    } else {
                    g0926UU = Language.To_Let(OBJ(g0926UU_try092710))
                    Result = Core.F_CALL(C_c_code,ARGS(EID{g0926UU.Id(),0},EID{s.Id(),0}))
                    }
                    /* Let-9 */} 
                  /* If-8 */} 
                }
                /* If-7 */} 
              }
              }}
              }
              /* Let-6 */} 
            /* Let-5 */} 
          }
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_code @ When (throw: true) 
func E_c_code_When (self EID,s EID) EID { 
    return /*(sm for c_code @ When= EID)*/ F_c_code_When(Language.To_When(OBJ(self)),ToClass(OBJ(s)) )} 
  
// *********************************************************************
// *     Part 4: Loops                                                 *
// *********************************************************************
// here we could do a return extraction
/* {1} OPT.The go function for: c_type(self:For) [] */
func F_c_type_For (self *Language.For ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var g0933UU *ClaireType  
      /* noccur = 1 */
      var g0933UU_try09343 EID 
      g0933UU_try09343 = F_Compile_return_type_any(self.Arg)
      /* ERROR PROTECTION INSERTED (g0933UU-Result) */
      if ErrorIn(g0933UU_try09343) {Result = g0933UU_try09343
      } else {
      g0933UU = ToType(OBJ(g0933UU_try09343))
      Result = EID{F_Optimize_infers_from_type(g0933UU,self.Id()).Id(),0}
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_type @ For (throw: true) 
func E_c_type_For (self EID) EID { 
    return /*(sm for c_type @ For= EID)*/ F_c_type_For(Language.To_For(OBJ(self)) )} 
  
/* {1} OPT.The go function for: infers_from(t:type,self:any) [] */
func F_Optimize_infers_from_type (t *ClaireType ,self *ClaireAny ) *ClaireType  { 
    // use function body compiling 
if (Equal(t.Id(),CEMPTY.Id()) == CTRUE) /* body If:2 */{ 
      return  F_Optimize_sort_abstract_I_type(ToType(C_boolean.Id()))
      }  else if (C_compiler.Safety > 3) /* body If:2 */{ 
      Core.F_tformat_string(MakeString("... c_type(~S) -> ~S - ~S \n"),2,MakeConstantList(self,t.Id(),F_Optimize_sort_abstract_I_type(t).Id()))
      return  F_Optimize_sort_abstract_I_type(t)
      } else {
      return  ToType(C_any.Id())
      /* body If-2 */} 
    } 
  
// The EID go function for: infers_from @ type (throw: false) 
func E_Optimize_infers_from_type (t EID,self EID) EID { 
    return EID{/*(sm for infers_from @ type= type)*/ F_Optimize_infers_from_type(ToType(OBJ(t)),ANY(self) ).Id(),0}} 
  
// false or the return value
// notice that for is of sort any and may require a cast ..
/* {1} OPT.The go function for: c_code(self:For,s:class) [] */
func F_c_code_For (self *Language.For ,s *ClaireClass ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var sx *ClaireAny   = self.SetArg
      /* noccur = 11 */
      /* Let:3 */{ 
        var ns int  = C_compiler.Safety
        /* noccur = 1 */
        /* Let:4 */{ 
          var vold *ClaireVariable   = self.ClaireVar
          /* noccur = 4 */
          /* Let:5 */{ 
            var v *ClaireVariable   = F_Compile_Variable_I_symbol(vold.Pname,vold.Index,Core.F_get_property(C_range,ToObject(vold.Id())))
            /* noccur = 12 */
            /* Let:6 */{ 
              var narg *ClaireAny   = Language.F_substitution_any(self.Arg,vold,v.Id())
              /* noccur = 3 */
              /* Let:7 */{ 
                var scs *ClaireAny  
                /* noccur = 2 */
                var scs_try09418 EID 
                scs_try09418 = F_Optimize_c_inline_arg_ask_any(sx)
                /* ERROR PROTECTION INSERTED (scs-Result) */
                if ErrorIn(scs_try09418) {Result = scs_try09418
                } else {
                scs = ANY(scs_try09418)
                if (sx.Isa.IsIn(Core.C_global_variable) == CTRUE) /* If:8 */{ 
                  /* Let:9 */{ 
                    var g0935 *Core.GlobalVariable   = Core.ToGlobalVariable(sx)
                    /* noccur = 5 */
                    if (F_boolean_I_any(g0935.Range.Id()).Id() != CTRUE.Id()) /* If:10 */{ 
                      self.SetArg = g0935.Value
                      g0935 = Core.ToGlobalVariable(g0935.Value)
                      /* If-10 */} 
                    sx = g0935.Id()
                    Result = sx.ToEID()
                    /* Let-9 */} 
                  /* If!8 */}  else if (sx.Isa.IsIn(Language.C_Select) == CTRUE) /* If:8 */{ 
                  /* Let:9 */{ 
                    var g0936 *Language.Select   = Language.To_Select(sx)
                    /* noccur = 1 */
                    /* Let:10 */{ 
                      var _Zt *ClaireType  
                      /* noccur = 2 */
                      var _Zt_try094211 EID 
                      _Zt_try094211 = Core.F_CALL(C_c_type,ARGS(EID{g0936.Id(),0}))
                      /* ERROR PROTECTION INSERTED (_Zt-Result) */
                      if ErrorIn(_Zt_try094211) {Result = _Zt_try094211
                      } else {
                      _Zt = ToType(OBJ(_Zt_try094211))
                      if ((_Zt.Included(ToType(C_list.Id())) != CTRUE) || 
                          (_Zt.Included(ToType(C_set.Id())) == CTRUE)) /* If:11 */{ 
                        Core.F_tformat_string(MakeString("STRANGE : transform ~S into a select ... \n"),0,MakeConstantList(self.Id()))
                        self = Language.To_For(self.Copy().Id())
                        /* update:12 */{ 
                          var va_arg1 *ClaireAny  
                          var va_arg2 *ClaireClass  
                          va_arg1 = self.Id()
                          va_arg2 = Language.C_Select
                          /* ---------- now we compile update isa(va_arg1) := va_arg2 ------- */
                          va_arg1.Isa = va_arg2
                          Result = EID{va_arg2.Id(),0}
                          /* update-12 */} 
                        } else {
                        Result = EID{CFALSE.Id(),0}
                        /* If-11 */} 
                      }
                      /* Let-10 */} 
                    /* Let-9 */} 
                  /* If!8 */}  else if (C_class.Id() == sx.Isa.Id()) /* If:8 */{ 
                  /* Let:9 */{ 
                    var g0937 *ClaireClass   = ToClass(sx)
                    /* noccur = 3 */
                    if ((g0937.Open <= 1) && 
                        (F_boolean_I_any(g0937.Subclass.Id()).Id() != CTRUE.Id())) /* If:10 */{ 
                      /* update:11 */{ 
                        var va_arg1 *Language.Iteration  
                        var va_arg2 *ClaireAny  
                        va_arg1 = Language.To_Iteration(self.Id())
                        /* Let:12 */{ 
                          var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                          /* noccur = 5 */
                          _CL_obj.Selector = C_instances
                          _CL_obj.Args = MakeConstantList(g0937.Id())
                          va_arg2 = _CL_obj.Id()
                          /* Let-12 */} 
                        /* ---------- now we compile update iClaire/set_arg(va_arg1) := va_arg2 ------- */
                        va_arg1.SetArg = va_arg2
                        Result = va_arg2.ToEID()
                        /* update-11 */} 
                      } else {
                      Result = EID{CFALSE.Id(),0}
                      /* If-10 */} 
                    /* Let-9 */} 
                  } else {
                  Result = EID{CFALSE.Id(),0}
                  /* If-8 */} 
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                /* Let:8 */{ 
                  var _Zt *ClaireAny  
                  /* noccur = 5 */
                  var _Zt_try09439 EID 
                  _Zt_try09439 = Core.F_CALL(C_c_type,ARGS(self.SetArg.ToEID()))
                  /* ERROR PROTECTION INSERTED (_Zt-Result) */
                  if ErrorIn(_Zt_try09439) {Result = _Zt_try09439
                  } else {
                  _Zt = ANY(_Zt_try09439)
                  /* Let:9 */{ 
                    var _Zt2 *ClaireType   = F_Optimize_pmember_type(ToType(_Zt))
                    /* noccur = 6 */
                    /* Let:10 */{ 
                      var _Zt3 *ClaireType  
                      /* noccur = 2 */
                      var _Zt3_try094411 EID 
                      /* Let:11 */{ 
                        var g0945UU *ClaireType  
                        /* noccur = 1 */
                        var g0945UU_try094612 EID 
                        /* Let:12 */{ 
                          var g0947UU *Language.Call  
                          /* noccur = 1 */
                          /* Let:13 */{ 
                            var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                            /* noccur = 5 */
                            _CL_obj.Selector = C_set_I
                            _CL_obj.Args = MakeConstantList(sx)
                            g0947UU = _CL_obj
                            /* Let-13 */} 
                          g0945UU_try094612 = Core.F_CALL(C_c_type,ARGS(EID{g0947UU.Id(),0}))
                          /* Let-12 */} 
                        /* ERROR PROTECTION INSERTED (g0945UU-_Zt3_try094411) */
                        if ErrorIn(g0945UU_try094612) {_Zt3_try094411 = g0945UU_try094612
                        } else {
                        g0945UU = ToType(OBJ(g0945UU_try094612))
                        _Zt3_try094411 = EID{F_Optimize_pmember_type(g0945UU).Id(),0}
                        }
                        /* Let-11 */} 
                      /* ERROR PROTECTION INSERTED (_Zt3-Result) */
                      if ErrorIn(_Zt3_try094411) {Result = _Zt3_try094411
                      } else {
                      _Zt3 = ToType(OBJ(_Zt3_try094411))
                      if (_Zt3.Included(_Zt2) == CTRUE) /* If:11 */{ 
                        _Zt2 = _Zt3
                        /* If-11 */} 
                      F_Optimize_range_infers_for_Variable(v,_Zt2,ToType(_Zt))
                      C_compiler.Safety = 1
                      _Zt2 = v.Range
                      v.Range = F_Optimize_ptype_type(_Zt2)
                      /* Let:11 */{ 
                        var m *ClaireAny   = F_Optimize_Iterate_I_Iteration(Language.To_Iteration(self.Id()))
                        /* noccur = 5 */
                        if (C_method.Id() != m.Isa.Id()) /* If:12 */{ 
                          /* Let:13 */{ 
                            var m2 *ClaireAny   = F_Optimize_restriction_I_property(Language.C_iterate,MakeConstantList(_Zt,MakeConstantSet(v.Id()).Id(),C_any.Id()),CTRUE)
                            /* noccur = 2 */
                            if (C_method.Id() == m2.Isa.Id()) /* If:14 */{ 
                              /* Let:15 */{ 
                                var g0938 *ClaireMethod   = ToMethod(m2)
                                /* noccur = 1 */
                                m = g0938.Id()
                                /* Let-15 */} 
                              /* If-14 */} 
                            /* Let-13 */} 
                          /* If-12 */} 
                        C_compiler.Safety = ns
                        v.Range = _Zt2
                        var g0948I *ClaireBoolean  
                        if (C_method.Id() == m.Isa.Id()) /* If:12 */{ 
                          /* Let:13 */{ 
                            var g0939 *ClaireMethod   = ToMethod(m)
                            /* noccur = 1 */
                            g0948I = g0939.Inline_ask
                            /* Let-13 */} 
                          } else {
                          g0948I = CFALSE
                          /* If-12 */} 
                        if (g0948I == CTRUE) /* If:12 */{ 
                          
                          if (F_Optimize_sort_abstract_ask_type(v.Range) == CTRUE) /* If:13 */{ 
                            v.Range = To_Union(v.Range.Id()).T2
                            /* If-13 */} 
                          Result = F_Optimize_c_inline_method1(ToMethod(m),MakeConstantList(Language.F_instruction_copy_any(self.SetArg),v.Id(),narg),s)
                          /* If!12 */}  else if (F_boolean_I_any(scs) == CTRUE) /* If:12 */{ 
                          /* Let:13 */{ 
                            var g0949UU *Language.For  
                            /* noccur = 1 */
                            /* Let:14 */{ 
                              var _CL_obj *Language.For   = Language.To_For(new(Language.For).Is(Language.C_For))
                              /* noccur = 5 */
                              _CL_obj.ClaireVar = v
                              _CL_obj.SetArg = scs
                              _CL_obj.Arg = narg
                              g0949UU = _CL_obj
                              /* Let-14 */} 
                            Result = Core.F_CALL(C_c_code,ARGS(EID{g0949UU.Id(),0},EID{s.Id(),0}))
                            /* Let-13 */} 
                          } else {
                          var g0950I *ClaireBoolean  
                          if (sx.Isa.IsIn(Language.C_Call) == CTRUE) /* If:13 */{ 
                            /* Let:14 */{ 
                              var g0940 *Language.Call   = Language.To_Call(sx)
                              /* noccur = 1 */
                              g0950I = Equal(g0940.Selector.Id(),Core.C_Id.Id())
                              /* Let-14 */} 
                            } else {
                            g0950I = CFALSE
                            /* If-13 */} 
                          if (g0950I == CTRUE) /* If:13 */{ 
                            Result = F_Optimize_c_code_multiple_For(self,ToType(_Zt),s)
                            } else {
                            /* Let:14 */{ 
                              var _CL_obj *Language.For   = Language.To_For(new(Language.For).Is(Language.C_For))
                              /* noccur = 5 */
                              _CL_obj.ClaireVar = v
                              /* update:15 */{ 
                                var va_arg1 *Language.Iteration  
                                var va_arg2 *ClaireAny  
                                va_arg1 = Language.To_Iteration(_CL_obj.Id())
                                var va_arg2_try095116 EID 
                                va_arg2_try095116 = F_Optimize_enumerate_code_any(self.SetArg,ToType(_Zt))
                                /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                                if ErrorIn(va_arg2_try095116) {Result = va_arg2_try095116
                                } else {
                                va_arg2 = ANY(va_arg2_try095116)
                                /* ---------- now we compile update iClaire/set_arg(va_arg1) := va_arg2 ------- */
                                va_arg1.SetArg = va_arg2
                                Result = va_arg2.ToEID()
                                }
                                /* update-15 */} 
                              /* ERROR PROTECTION INSERTED (Result-Result) */
                              if !ErrorIn(Result) {
                              /* update:15 */{ 
                                var va_arg1 *Language.Iteration  
                                var va_arg2 *ClaireAny  
                                va_arg1 = Language.To_Iteration(_CL_obj.Id())
                                var va_arg2_try095216 EID 
                                va_arg2_try095216 = Core.F_CALL(C_c_code,ARGS(narg.ToEID(),EID{C_void.Id(),0}))
                                /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                                if ErrorIn(va_arg2_try095216) {Result = va_arg2_try095216
                                } else {
                                va_arg2 = ANY(va_arg2_try095216)
                                /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                                va_arg1.Arg = va_arg2
                                Result = va_arg2.ToEID()
                                }
                                /* update-15 */} 
                              /* ERROR PROTECTION INSERTED (Result-Result) */
                              if !ErrorIn(Result) {
                              Result = EID{_CL_obj.Id(),0}
                              }}
                              /* Let-14 */} 
                            /* If-13 */} 
                          /* If-12 */} 
                        /* Let-11 */} 
                      }
                      /* Let-10 */} 
                    /* Let-9 */} 
                  }
                  /* Let-8 */} 
                }
                }
                /* Let-7 */} 
              /* Let-6 */} 
            /* Let-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_code @ For (throw: true) 
func E_c_code_For (self EID,s EID) EID { 
    return /*(sm for c_code @ For= EID)*/ F_c_code_For(Language.To_For(OBJ(self)),ToClass(OBJ(s)) )} 
  
//             (if (s = any) r2 else to_C(arg = r2, set_arg = s)))))]
// new: we macroexpand the iteration  based on the type
// this is only called if the set is wrapped in an Id
/* {1} OPT.The go function for: c_code_multiple(self:For,%t:type,s:class) [] */
func F_Optimize_c_code_multiple_For (self *Language.For ,_Zt *ClaireType ,s *ClaireClass ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var v *ClaireVariable   = self.ClaireVar
      /* noccur = 7 */
      /* Let:3 */{ 
        var sx *ClaireAny   = Language.To_Call(self.SetArg).Args.At(1-1)
        /* noccur = 2 */
        /* Let:4 */{ 
          var v2 *ClaireVariable   = F_Compile_Variable_I_symbol(F_append_symbol(v.Pname,MakeString("test").Id()),self.ClaireVar.Index,_Zt.Id())
          /* noccur = 3 */
          /* Let:5 */{ 
            var n *Language.Let  
            /* noccur = 4 */
            var n_try09536 EID 
            /* Let:6 */{ 
              var _CL_obj *Language.Let   = Language.To_Let(new(Language.Let).Is(Language.C_Let))
              /* noccur = 12 */
              _CL_obj.ClaireVar = v2
              _CL_obj.Value = sx
              /* update:7 */{ 
                var va_arg1 *Language.Let  
                var va_arg2 *ClaireAny  
                va_arg1 = _CL_obj
                var va_arg2_try09548 EID 
                /* Let:8 */{ 
                  var _CL_obj *Language.For   = Language.To_For(new(Language.For).Is(Language.C_For))
                  /* noccur = 5 */
                  _CL_obj.ClaireVar = self.ClaireVar
                  /* update:9 */{ 
                    var va_arg1 *Language.Iteration  
                    var va_arg2 *ClaireAny  
                    va_arg1 = Language.To_Iteration(_CL_obj.Id())
                    var va_arg2_try095510 EID 
                    va_arg2_try095510 = F_Optimize_enumerate_code_any(sx,_Zt)
                    /* ERROR PROTECTION INSERTED (va_arg2-va_arg2_try09548) */
                    if ErrorIn(va_arg2_try095510) {va_arg2_try09548 = va_arg2_try095510
                    } else {
                    va_arg2 = ANY(va_arg2_try095510)
                    /* ---------- now we compile update iClaire/set_arg(va_arg1) := va_arg2 ------- */
                    va_arg1.SetArg = va_arg2
                    va_arg2_try09548 = va_arg2.ToEID()
                    }
                    /* update-9 */} 
                  /* ERROR PROTECTION INSERTED (va_arg2_try09548-va_arg2_try09548) */
                  if !ErrorIn(va_arg2_try09548) {
                  _CL_obj.Arg = self.Arg
                  va_arg2_try09548 = EID{_CL_obj.Id(),0}
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (va_arg2-n_try09536) */
                if ErrorIn(va_arg2_try09548) {n_try09536 = va_arg2_try09548
                } else {
                va_arg2 = ANY(va_arg2_try09548)
                /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                va_arg1.Arg = va_arg2
                n_try09536 = va_arg2.ToEID()
                }
                /* update-7 */} 
              /* ERROR PROTECTION INSERTED (n_try09536-n_try09536) */
              if !ErrorIn(n_try09536) {
              n_try09536 = EID{_CL_obj.Id(),0}
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (n-Result) */
            if ErrorIn(n_try09536) {Result = n_try09536
            } else {
            n = Language.To_Let(OBJ(n_try09536))
            Core.F_tformat_string(MakeString("---- note: use an expended iteration for {~S} \n"),0,MakeConstantList(self.Id()))
            /* For:6 */{ 
              var r *ClaireAny  
              _ = r
              Result= EID{CFALSE.Id(),0}
              for _,r = range(Language.C_iterate.Restrictions.ValuesO())/* loop:7 */{ 
                var void_try8 EID 
                _ = void_try8
                var g0956I *ClaireBoolean  
                var g0956I_try09578 EID 
                /* and:8 */{ 
                  var v_and8 *ClaireBoolean  
                  
                  v_and8 = ToType(Core.F_domain_I_restriction(ToRestriction(r)).Id()).Included(_Zt)
                  if (v_and8 == CFALSE) {g0956I_try09578 = EID{CFALSE.Id(),0}
                  } else /* arg:9 */{ 
                    v_and8 = ToType(Core.F_domain_I_restriction(ToRestriction(r)).Id()).Included(ToType(C_collection.Id()))
                    if (v_and8 == CFALSE) {g0956I_try09578 = EID{CFALSE.Id(),0}
                    } else /* arg:10 */{ 
                      v_and8 = ToMethod(r).Inline_ask
                      if (v_and8 == CFALSE) {g0956I_try09578 = EID{CFALSE.Id(),0}
                      } else /* arg:11 */{ 
                        var v_and8_try095812 EID 
                        v_and8_try095812 = Core.F_BELONG(v.Id(),ToRestriction(r).Domain.ValuesO()[2-1])
                        /* ERROR PROTECTION INSERTED (v_and8-g0956I_try09578) */
                        if ErrorIn(v_and8_try095812) {g0956I_try09578 = v_and8_try095812
                        } else {
                        v_and8 = ToBoolean(OBJ(v_and8_try095812))
                        if (v_and8 == CFALSE) {g0956I_try09578 = EID{CFALSE.Id(),0}
                        } else /* arg:12 */{ 
                          g0956I_try09578 = EID{CTRUE.Id(),0}/* arg-12 */} 
                        /* arg-11 */} 
                      /* arg-10 */} 
                    /* arg-9 */} 
                  }
                  /* and-8 */} 
                /* ERROR PROTECTION INSERTED (g0956I-void_try8) */
                if ErrorIn(g0956I_try09578) {void_try8 = g0956I_try09578
                } else {
                g0956I = ToBoolean(OBJ(g0956I_try09578))
                if (g0956I == CTRUE) /* If:8 */{ 
                  /* Let:9 */{ 
                    var vnew *ClaireVariable  
                    /* noccur = 4 */
                    /* Let:10 */{ 
                      var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
                      /* noccur = 7 */
                      _CL_obj.Pname = v.Pname
                      _CL_obj.Range = v.Range
                      _CL_obj.Index = v.Index
                      vnew = _CL_obj
                      /* Let-10 */} 
                    /* Let:10 */{ 
                      var narg *ClaireAny   = Language.F_substitution_any(self.Arg,v,vnew.Id())
                      /* noccur = 1 */
                      /* update:11 */{ 
                        var va_arg1 *Language.Let  
                        var va_arg2 *ClaireAny  
                        va_arg1 = n
                        var va_arg2_try095912 EID 
                        /* Let:12 */{ 
                          var g0960UU *Language.Call  
                          /* noccur = 1 */
                          /* Let:13 */{ 
                            var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                            /* noccur = 5 */
                            _CL_obj.Selector = ToProperty(C__Z.Id())
                            _CL_obj.Args = MakeConstantList(v2.Id(),ToRestriction(r).Domain.ValuesO()[1-1])
                            g0960UU = _CL_obj
                            /* Let-13 */} 
                          /* Let:13 */{ 
                            var g0961UU *ClaireAny  
                            /* noccur = 1 */
                            var g0961UU_try096214 EID 
                            if (F_Optimize_sort_abstract_ask_type(vnew.Range) == CTRUE) /* If:14 */{ 
                              vnew.Range = To_Union(v.Range.Id()).T2
                              /* If-14 */} 
                            g0961UU_try096214 = F_Optimize_c_inline_method1(ToMethod(r),MakeConstantList(v2.Id(),vnew.Id(),narg),s)
                            /* ERROR PROTECTION INSERTED (g0961UU-va_arg2_try095912) */
                            if ErrorIn(g0961UU_try096214) {va_arg2_try095912 = g0961UU_try096214
                            } else {
                            g0961UU = ANY(g0961UU_try096214)
                            va_arg2_try095912 = Language.C_If.Make(g0960UU.Id(),g0961UU,n.Arg).ToEID()
                            }
                            /* Let-13 */} 
                          /* Let-12 */} 
                        /* ERROR PROTECTION INSERTED (va_arg2-void_try8) */
                        if ErrorIn(va_arg2_try095912) {void_try8 = va_arg2_try095912
                        } else {
                        va_arg2 = ANY(va_arg2_try095912)
                        /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                        va_arg1.Arg = va_arg2
                        void_try8 = va_arg2.ToEID()
                        }
                        /* update-11 */} 
                      /* Let-10 */} 
                    /* Let-9 */} 
                  } else {
                  void_try8 = EID{CFALSE.Id(),0}
                  /* If-8 */} 
                }
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
            Result = Core.F_CALL(C_c_code,ARGS(EID{n.Id(),0},EID{s.Id(),0}))
            }
            }
            /* Let-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_code_multiple @ For (throw: true) 
func E_Optimize_c_code_multiple_For (self EID,_Zt EID,s EID) EID { 
    return /*(sm for c_code_multiple @ For= EID)*/ F_Optimize_c_code_multiple_For(Language.To_For(OBJ(self)),ToType(OBJ(_Zt)),ToClass(OBJ(s)) )} 
  
// ------------------------ Collect/ Image / Select / Lselect ------------------------------
// an Iteration builds a set
/* {1} OPT.The go function for: c_type(self:Iteration) [] */
func F_c_type_Iteration (self *Language.Iteration ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var _Zt *ClaireAny  
      /* noccur = 2 */
      var _Zt_try09633 EID 
      _Zt_try09633 = Core.F_CALL(C_c_type,ARGS(self.SetArg.ToEID()))
      /* ERROR PROTECTION INSERTED (_Zt-Result) */
      if ErrorIn(_Zt_try09633) {Result = _Zt_try09633
      } else {
      _Zt = ANY(_Zt_try09633)
      F_Optimize_range_infers_for_Variable(self.ClaireVar,F_Optimize_pmember_type(ToType(_Zt)),ToType(_Zt))
      if (Core.F_get_property(C_of,ToObject(self.Id())) != CNULL) /* If:3 */{ 
        /* Let:4 */{ 
          var g0964UU *ClaireClass  
          /* noccur = 1 */
          if ((self.Isa.IsIn(Language.C_Select) == CTRUE) || 
              (self.Isa.IsIn(Language.C_Image) == CTRUE)) /* If:5 */{ 
            g0964UU = C_set
            } else {
            g0964UU = C_list
            /* If-5 */} 
          Result = EID{Core.F_param_I_class(g0964UU,ToType(OBJ(Core.F_CALL(C_of,ARGS(EID{self.Id(),0}))))).Id(),0}
          /* Let-4 */} 
        } else {
        /* Let:4 */{ 
          var g0965UU *ClaireClass  
          /* noccur = 1 */
          if ((self.Isa.IsIn(Language.C_Select) == CTRUE) || 
              (self.Isa.IsIn(Language.C_Image) == CTRUE)) /* If:5 */{ 
            g0965UU = C_set
            } else {
            g0965UU = C_list
            /* If-5 */} 
          /* Let:5 */{ 
            var g0966UU *ClaireType  
            /* noccur = 1 */
            var g0966UU_try09676 EID 
            if ((self.Isa.IsIn(Language.C_Select) == CTRUE) || 
                (self.Isa.IsIn(Language.C_Lselect) == CTRUE)) /* If:6 */{ 
              /* Let:7 */{ 
                var g0968UU *ClaireType  
                /* noccur = 1 */
                var g0968UU_try09698 EID 
                g0968UU_try09698 = Core.F_CALL(C_c_type,ARGS(self.SetArg.ToEID()))
                /* ERROR PROTECTION INSERTED (g0968UU-g0966UU_try09676) */
                if ErrorIn(g0968UU_try09698) {g0966UU_try09676 = g0968UU_try09698
                } else {
                g0968UU = ToType(OBJ(g0968UU_try09698))
                g0966UU_try09676 = EID{F_Optimize_pmember_type(g0968UU).Id(),0}
                }
                /* Let-7 */} 
              } else {
              /* Let:7 */{ 
                var g0970UU *ClaireType  
                /* noccur = 1 */
                var g0970UU_try09718 EID 
                g0970UU_try09718 = Core.F_CALL(C_c_type,ARGS(self.Arg.ToEID()))
                /* ERROR PROTECTION INSERTED (g0970UU-g0966UU_try09676) */
                if ErrorIn(g0970UU_try09718) {g0966UU_try09676 = g0970UU_try09718
                } else {
                g0970UU = ToType(OBJ(g0970UU_try09718))
                g0966UU_try09676 = EID{F_Optimize_ptype_type(g0970UU).Id(),0}
                }
                /* Let-7 */} 
              /* If-6 */} 
            /* ERROR PROTECTION INSERTED (g0966UU-Result) */
            if ErrorIn(g0966UU_try09676) {Result = g0966UU_try09676
            } else {
            g0966UU = ToType(OBJ(g0966UU_try09676))
            Result = EID{Core.F_nth_class1(g0965UU,g0966UU).Id(),0}
            }
            /* Let-5 */} 
          /* Let-4 */} 
        /* If-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_type @ Iteration (throw: true) 
func E_c_type_Iteration (self EID) EID { 
    return /*(sm for c_type @ Iteration= EID)*/ F_c_type_Iteration(Language.To_Iteration(OBJ(self)) )} 
  
// They are all expended into a For except for Collect(bag : list or set)
/* {1} OPT.The go function for: c_code(self:Iteration) [] */
func F_c_code_Iteration (self *Language.Iteration ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var sx *ClaireAny   = self.SetArg
      /* noccur = 3 */
      /* Let:3 */{ 
        var _Zt *ClaireAny  
        /* noccur = 5 */
        var _Zt_try09744 EID 
        _Zt_try09744 = Core.F_CALL(C_c_type,ARGS(sx.ToEID()))
        /* ERROR PROTECTION INSERTED (_Zt-Result) */
        if ErrorIn(_Zt_try09744) {Result = _Zt_try09744
        } else {
        _Zt = ANY(_Zt_try09744)
        if (self.Isa.IsIn(Language.C_For) == CTRUE) /* If:4 */{ 
          Result = Core.F_CALL(C_c_code,ARGS(EID{self.Id(),0},EID{C_any.Id(),0}))
          } else {
          var g0975I *ClaireBoolean  
          if (self.Isa.IsIn(Language.C_Collect) == CTRUE) /* If:5 */{ 
            g0975I = MakeBoolean((ToType(_Zt).Included(ToType(C_list.Id())) == CTRUE) || (ToType(_Zt).Included(ToType(C_set.Id())) == CTRUE))
            } else {
            g0975I = CFALSE
            /* If-5 */} 
          if (g0975I == CTRUE) /* If:5 */{ 
            F_Optimize_range_infers_for_Variable(self.ClaireVar,F_Optimize_pmember_type(ToType(_Zt)),ToType(_Zt))
            /* Let:6 */{ 
              var ty *ClaireType  
              /* noccur = 2 */
              var ty_try09767 EID 
              /* Let:7 */{ 
                var g0977UU *ClaireType  
                /* noccur = 1 */
                var g0977UU_try09788 EID 
                g0977UU_try09788 = Core.F_CALL(C_c_type,ARGS(self.Arg.ToEID()))
                /* ERROR PROTECTION INSERTED (g0977UU-ty_try09767) */
                if ErrorIn(g0977UU_try09788) {ty_try09767 = g0977UU_try09788
                } else {
                g0977UU = ToType(OBJ(g0977UU_try09788))
                ty_try09767 = EID{F_Optimize_ptype_type(g0977UU).Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (ty-Result) */
              if ErrorIn(ty_try09767) {Result = ty_try09767
              } else {
              ty = ToType(OBJ(ty_try09767))
              /* Let:7 */{ 
                var x *Language.Collect  
                /* noccur = 5 */
                var x_try09798 EID 
                /* Let:8 */{ 
                  var _CL_obj *Language.Collect   = Language.To_Collect(new(Language.Collect).Is(Language.C_Collect))
                  /* noccur = 5 */
                  _CL_obj.ClaireVar = self.ClaireVar
                  /* update:9 */{ 
                    var va_arg1 *Language.Iteration  
                    var va_arg2 *ClaireAny  
                    va_arg1 = Language.To_Iteration(_CL_obj.Id())
                    var va_arg2_try098010 EID 
                    va_arg2_try098010 = F_Compile_c_strict_code_any(sx,ToTypeExpression(_Zt).Class_I())
                    /* ERROR PROTECTION INSERTED (va_arg2-x_try09798) */
                    if ErrorIn(va_arg2_try098010) {x_try09798 = va_arg2_try098010
                    } else {
                    va_arg2 = ANY(va_arg2_try098010)
                    /* ---------- now we compile update iClaire/set_arg(va_arg1) := va_arg2 ------- */
                    va_arg1.SetArg = va_arg2
                    x_try09798 = va_arg2.ToEID()
                    }
                    /* update-9 */} 
                  /* ERROR PROTECTION INSERTED (x_try09798-x_try09798) */
                  if !ErrorIn(x_try09798) {
                  /* update:9 */{ 
                    var va_arg1 *Language.Iteration  
                    var va_arg2 *ClaireAny  
                    va_arg1 = Language.To_Iteration(_CL_obj.Id())
                    var va_arg2_try098110 EID 
                    va_arg2_try098110 = Core.F_CALL(C_c_code,ARGS(self.Arg.ToEID(),EID{C_any.Id(),0}))
                    /* ERROR PROTECTION INSERTED (va_arg2-x_try09798) */
                    if ErrorIn(va_arg2_try098110) {x_try09798 = va_arg2_try098110
                    } else {
                    va_arg2 = ANY(va_arg2_try098110)
                    /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                    va_arg1.Arg = va_arg2
                    x_try09798 = va_arg2.ToEID()
                    }
                    /* update-9 */} 
                  /* ERROR PROTECTION INSERTED (x_try09798-x_try09798) */
                  if !ErrorIn(x_try09798) {
                  x_try09798 = EID{_CL_obj.Id(),0}
                  }}
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (x-Result) */
                if ErrorIn(x_try09798) {Result = x_try09798
                } else {
                x = Language.To_Collect(OBJ(x_try09798))
                if (ty.Id() == C_void.Id()) /* If:8 */{ 
                  Result = Core.F_CALL(C_Compile_Cerror,ARGS(EID{MakeString("[205] use of void expression ~S in ~S").Id(),0},self.Arg.ToEID(),EID{self.Id(),0}))
                  } else {
                  Result = EID{CFALSE.Id(),0}
                  /* If-8 */} 
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                if (Core.F_get_property(C_of,ToObject(self.Id())) != CNULL) /* If:8 */{ 
                  if ((C_compiler.Safety > 4) || 
                      (ty.Included(ToType(OBJ(Core.F_CALL(C_of,ARGS(EID{self.Id(),0}))))) == CTRUE)) /* If:9 */{ 
                    x.Of = ToType(OBJ(Core.F_CALL(C_of,ARGS(EID{self.Id(),0}))))
                    Result = EID{x.Id(),0}
                    } else {
                    F_Compile_warn_void()
                    /* Let:10 */{ 
                      var g0982UU *ClaireList  
                      /* noccur = 1 */
                      var g0982UU_try098311 EID 
                      /* Construct:11 */{ 
                        var v_bag_arg *ClaireAny  
                        g0982UU_try098311= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                        ToList(OBJ(g0982UU_try098311)).AddFast(self.Id())
                        var v_bag_arg_try098412 EID 
                        v_bag_arg_try098412 = Core.F_CALL(C_c_type,ARGS(self.Arg.ToEID()))
                        /* ERROR PROTECTION INSERTED (v_bag_arg-g0982UU_try098311) */
                        if ErrorIn(v_bag_arg_try098412) {g0982UU_try098311 = v_bag_arg_try098412
                        } else {
                        v_bag_arg = ANY(v_bag_arg_try098412)
                        ToList(OBJ(g0982UU_try098311)).AddFast(v_bag_arg)
                        ToList(OBJ(g0982UU_try098311)).AddFast(ANY(Core.F_CALL(C_of,ARGS(EID{self.Id(),0}))))}
                        /* Construct-11 */} 
                      /* ERROR PROTECTION INSERTED (g0982UU-Result) */
                      if ErrorIn(g0982UU_try098311) {Result = g0982UU_try098311
                      } else {
                      g0982UU = ToList(OBJ(g0982UU_try098311))
                      Result = Core.F_tformat_string(MakeString("unsafe typed collect (~S): ~S not in ~S [261]\n"),2,g0982UU)
                      }
                      /* Let-10 */} 
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    /* Let:10 */{ 
                      var g0985UU *Language.Call  
                      /* noccur = 1 */
                      /* Let:11 */{ 
                        var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                        /* noccur = 5 */
                        _CL_obj.Selector = Core.C_check_in
                        _CL_obj.Args = MakeConstantList(x.Id(),C_list.Id(),ANY(Core.F_CALL(C_of,ARGS(EID{self.Id(),0}))))
                        g0985UU = _CL_obj
                        /* Let-11 */} 
                      Result = Core.F_CALL(C_c_code,ARGS(EID{g0985UU.Id(),0},EID{C_list.Id(),0}))
                      /* Let-10 */} 
                    }
                    /* If-9 */} 
                  } else {
                  Result = EID{x.Id(),0}
                  /* If-8 */} 
                }
                }
                /* Let-7 */} 
              }
              /* Let-6 */} 
            } else {
            /* Let:6 */{ 
              var val *ClaireAny  
              /* noccur = 4 */
              if (self.Isa.IsIn(Language.C_Image) == CTRUE) /* If:7 */{ 
                val = ToType(CEMPTY.Id()).EmptySet().Id()
                } else {
                val = ToType(CEMPTY.Id()).EmptyList().Id()
                /* If-7 */} 
              /* Let:7 */{ 
                var v *ClaireVariable  
                /* noccur = 5 */
                /* Let:8 */{ 
                  var g0986UU int 
                  /* noccur = 1 */
                  C_OPT.MaxVars = (C_OPT.MaxVars+1)
                  g0986UU = 0
                  /* Let:9 */{ 
                    var g0987UU *ClaireClass  
                    /* noccur = 1 */
                    if (self.Isa.IsIn(Language.C_Image) == CTRUE) /* If:10 */{ 
                      g0987UU = C_set
                      } else {
                      g0987UU = C_list
                      /* If-10 */} 
                    v = F_Compile_Variable_I_symbol(F_append_symbol(self.ClaireVar.Pname,MakeString("_bag").Id()),g0986UU,g0987UU.Id())
                    /* Let-9 */} 
                  /* Let-8 */} 
                
                if (Core.F_get_property(C_of,ToObject(self.Id())) != CNULL) /* If:8 */{ 
                  /* Let:9 */{ 
                    var _ZtypeIn *ClaireType  
                    /* noccur = 2 */
                    var _ZtypeIn_try098810 EID 
                    _ZtypeIn_try098810 = Core.F_CALL(C_c_type,ARGS(self.Arg.ToEID()))
                    /* ERROR PROTECTION INSERTED (_ZtypeIn-Result) */
                    if ErrorIn(_ZtypeIn_try098810) {Result = _ZtypeIn_try098810
                    } else {
                    _ZtypeIn = ToType(OBJ(_ZtypeIn_try098810))
                    if ((F_Optimize_ptype_type(_ZtypeIn).Included(ToType(OBJ(Core.F_CALL(C_of,ARGS(EID{self.Id(),0}))))) != CTRUE) && 
                        (C_compiler.Safety <= 4)) /* If:10 */{ 
                      F_Compile_warn_void()
                      Core.F_tformat_string(MakeString("unsafe bag construction (~S) : a ~S is not a ~S [262]\n"),2,MakeConstantList(self.ClaireVar.Id(),_ZtypeIn.Id(),ANY(Core.F_CALL(C_of,ARGS(EID{self.Id(),0})))))
                      /* If-10 */} 
                    ToBag(val).Cast_I(ToType(OBJ(Core.F_CALL(C_of,ARGS(EID{self.Id(),0})))))
                    /* update:10 */{ 
                      var va_arg1 *ClaireVariable  
                      var va_arg2 *ClaireType  
                      va_arg1 = v
                      va_arg2 = Core.F_param_I_class(ToClass(v.Range.Id()),ToType(OBJ(Core.F_CALL(C_of,ARGS(EID{self.Id(),0})))))
                      /* ---------- now we compile update range(va_arg1) := va_arg2 ------- */
                      va_arg1.Range = va_arg2
                      Result = EID{va_arg2.Id(),0}
                      /* update-10 */} 
                    }
                    /* Let-9 */} 
                  } else {
                  if (C_set.Id() == val.Isa.Id()) /* If:9 */{ 
                    /* Let:10 */{ 
                      var _CL_obj *Language.Set   = Language.To_Set(new(Language.Set).Is(Language.C_Set))
                      /* noccur = 3 */
                      _CL_obj.Of = ToType(CEMPTY.Id())
                      val = _CL_obj.Id()
                      /* Let-10 */} 
                    } else {
                    /* Let:10 */{ 
                      var _CL_obj *Language.List   = Language.To_List(new(Language.List).Is(Language.C_List))
                      /* noccur = 3 */
                      _CL_obj.Of = ToType(CEMPTY.Id())
                      val = _CL_obj.Id()
                      /* Let-10 */} 
                    /* If-9 */} 
                  Result = val.ToEID()
                  /* If-8 */} 
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                /* Let:8 */{ 
                  var _CL_obj *Language.Let   = Language.To_Let(new(Language.Let).Is(Language.C_Let))
                  /* noccur = 22 */
                  _CL_obj.ClaireVar = v
                  _CL_obj.Value = val
                  /* update:9 */{ 
                    var va_arg1 *Language.Let  
                    var va_arg2 *ClaireAny  
                    va_arg1 = _CL_obj
                    var va_arg2_try098910 EID 
                    /* Let:10 */{ 
                      var _CL_obj *Language.Do   = Language.To_Do(new(Language.Do).Is(Language.C_Do))
                      /* noccur = 15 */
                      /* update:11 */{ 
                        var va_arg1 *Language.Do  
                        var va_arg2 *ClaireList  
                        va_arg1 = _CL_obj
                        var va_arg2_try099012 EID 
                        /* Construct:12 */{ 
                          var v_bag_arg *ClaireAny  
                          va_arg2_try099012= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                          var v_bag_arg_try099113 EID 
                          /* Let:13 */{ 
                            var g0992UU *Language.For  
                            /* noccur = 1 */
                            /* Let:14 */{ 
                              var _CL_obj *Language.For   = Language.To_For(new(Language.For).Is(Language.C_For))
                              /* noccur = 11 */
                              _CL_obj.ClaireVar = self.ClaireVar
                              _CL_obj.SetArg = sx
                              /* update:15 */{ 
                                var va_arg1 *Language.Iteration  
                                var va_arg2 *ClaireAny  
                                va_arg1 = Language.To_Iteration(_CL_obj.Id())
                                /* Let:16 */{ 
                                  var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                  /* noccur = 5 */
                                  _CL_obj.Selector = ToProperty(C_add_I.Id())
                                  _CL_obj.Args = MakeConstantList(v.Id(),self.Arg)
                                  va_arg2 = _CL_obj.Id()
                                  /* Let-16 */} 
                                /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                                va_arg1.Arg = va_arg2
                                /* update-15 */} 
                              g0992UU = _CL_obj
                              /* Let-14 */} 
                            v_bag_arg_try099113 = Core.F_CALL(C_c_code,ARGS(EID{g0992UU.Id(),0},EID{C_any.Id(),0}))
                            /* Let-13 */} 
                          /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try099012) */
                          if ErrorIn(v_bag_arg_try099113) {va_arg2_try099012 = v_bag_arg_try099113
                          } else {
                          v_bag_arg = ANY(v_bag_arg_try099113)
                          ToList(OBJ(va_arg2_try099012)).AddFast(v_bag_arg)
                          ToList(OBJ(va_arg2_try099012)).AddFast(v.Id())}
                          /* Construct-12 */} 
                        /* ERROR PROTECTION INSERTED (va_arg2-va_arg2_try098910) */
                        if ErrorIn(va_arg2_try099012) {va_arg2_try098910 = va_arg2_try099012
                        } else {
                        va_arg2 = ToList(OBJ(va_arg2_try099012))
                        /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                        va_arg1.Args = va_arg2
                        va_arg2_try098910 = EID{va_arg2.Id(),0}
                        }
                        /* update-11 */} 
                      /* ERROR PROTECTION INSERTED (va_arg2_try098910-va_arg2_try098910) */
                      if !ErrorIn(va_arg2_try098910) {
                      va_arg2_try098910 = EID{_CL_obj.Id(),0}
                      }
                      /* Let-10 */} 
                    /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                    if ErrorIn(va_arg2_try098910) {Result = va_arg2_try098910
                    } else {
                    va_arg2 = ANY(va_arg2_try098910)
                    /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                    va_arg1.Arg = va_arg2
                    Result = va_arg2.ToEID()
                    }
                    /* update-9 */} 
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  Result = EID{_CL_obj.Id(),0}
                  }
                  /* Let-8 */} 
                }
                /* Let-7 */} 
              /* Let-6 */} 
            /* If-5 */} 
          /* If-4 */} 
        }
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_code @ Iteration (throw: true) 
func E_c_code_Iteration (self EID) EID { 
    return /*(sm for c_code @ Iteration= EID)*/ F_c_code_Iteration(Language.To_Iteration(OBJ(self)) )} 
  
// new in v3.1.16
// selection has its own optimization method that takes care of the polymorphism
/* {1} OPT.The go function for: c_code(self:Select) [] */
func F_c_code_Select (self *Language.Select ) EID { 
    var Result EID 
    Result = F_Optimize_c_code_select_Iteration(Language.To_Iteration(self.Id()),C_set)
    return Result} 
  
// The EID go function for: c_code @ Select (throw: true) 
func E_c_code_Select (self EID) EID { 
    return /*(sm for c_code @ Select= EID)*/ F_c_code_Select(Language.To_Select(OBJ(self)) )} 
  
/* {1} OPT.The go function for: c_code(self:Lselect) [] */
func F_c_code_Lselect (self *Language.Lselect ) EID { 
    var Result EID 
    Result = F_Optimize_c_code_select_Iteration(Language.To_Iteration(self.Id()),C_list)
    return Result} 
  
// The EID go function for: c_code @ Lselect (throw: true) 
func E_c_code_Lselect (self EID) EID { 
    return /*(sm for c_code @ Lselect= EID)*/ F_c_code_Lselect(Language.To_Lselect(OBJ(self)) )} 
  
// changed in CLAIRE 4 (cf trans -> init.cl)
// x is set or list, tells what we want as output
/* {1} OPT.The go function for: c_code_select(self:Iteration,x:class) [] */
func F_Optimize_c_code_select_Iteration (self *Language.Iteration ,x *ClaireClass ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var sx *ClaireAny   = self.SetArg
      /* noccur = 4 */
      /* Let:3 */{ 
        var _Zt *ClaireAny  
        /* noccur = 4 */
        var _Zt_try09954 EID 
        _Zt_try09954 = Core.F_CALL(C_c_type,ARGS(sx.ToEID()))
        /* ERROR PROTECTION INSERTED (_Zt-Result) */
        if ErrorIn(_Zt_try09954) {Result = _Zt_try09954
        } else {
        _Zt = ANY(_Zt_try09954)
        /* Let:4 */{ 
          var st *ClaireAny  
          /* noccur = 1 */
          var st_try09965 EID 
          st_try09965 = F_Optimize_enumerate_code_any(sx,ToType(_Zt))
          /* ERROR PROTECTION INSERTED (st-Result) */
          if ErrorIn(st_try09965) {Result = st_try09965
          } else {
          st = ANY(st_try09965)
          /* Let:5 */{ 
            var val *ClaireBag  
            /* noccur = 2 */
            if (x.Id() == C_set.Id()) /* If:6 */{ 
              val = ToBag(ToType(CEMPTY.Id()).EmptySet().Id())
              } else {
              val = ToBag(ToType(CEMPTY.Id()).EmptyList().Id())
              /* If-6 */} 
            /* Let:6 */{ 
              var v1 *ClaireVariable  
              /* noccur = 3 */
              /* Let:7 */{ 
                var g0997UU int 
                /* noccur = 1 */
                C_OPT.MaxVars = (C_OPT.MaxVars+1)
                g0997UU = 0
                v1 = F_Compile_Variable_I_symbol(F_append_symbol(self.ClaireVar.Pname,MakeString("_in").Id()),g0997UU,_Zt)
                /* Let-7 */} 
              /* Let:7 */{ 
                var v2 *ClaireVariable  
                /* noccur = 4 */
                /* Let:8 */{ 
                  var g0998UU int 
                  /* noccur = 1 */
                  C_OPT.MaxVars = (C_OPT.MaxVars+1)
                  g0998UU = 0
                  v2 = F_Compile_Variable_I_symbol(F_append_symbol(self.ClaireVar.Pname,MakeString("_out").Id()),g0998UU,x.Id())
                  /* Let-8 */} 
                if (Core.F_get_property(C_of,ToObject(self.Id())) != CNULL) /* If:8 */{ 
                  /* Let:9 */{ 
                    var _ZtypeIn *ClaireType   = F_Optimize_pmember_type(ToType(_Zt))
                    /* noccur = 2 */
                    if ((F_Optimize_ptype_type(_ZtypeIn).Included(ToType(OBJ(Core.F_CALL(C_of,ARGS(EID{self.Id(),0}))))) != CTRUE) && 
                        (C_compiler.Safety <= 4)) /* If:10 */{ 
                      F_Compile_warn_void()
                      Core.F_tformat_string(MakeString("unsafe bag construction (~S) : a ~S is not a ~S [262]\n"),2,MakeConstantList(self.ClaireVar.Id(),_ZtypeIn.Id(),ANY(Core.F_CALL(C_of,ARGS(EID{self.Id(),0})))))
                      /* If-10 */} 
                    val.Cast_I(ToType(OBJ(Core.F_CALL(C_of,ARGS(EID{self.Id(),0})))))
                    v2.Range = Core.F_param_I_class(x,ToType(OBJ(Core.F_CALL(C_of,ARGS(EID{self.Id(),0})))))
                    Result = F_Optimize_inner_select_Iteration(self,v2.Id(),sx,val.Id())
                    /* Let-9 */} 
                  /* If!8 */}  else if (ToType(_Zt).Included(ToType(x.Id())) == CTRUE) /* If:8 */{ 
                  /* Let:9 */{ 
                    var _CL_obj *Language.Let   = Language.To_Let(new(Language.Let).Is(Language.C_Let))
                    /* noccur = 17 */
                    _CL_obj.ClaireVar = v1
                    _CL_obj.Value = st
                    /* update:10 */{ 
                      var va_arg1 *Language.Let  
                      var va_arg2 *ClaireAny  
                      va_arg1 = _CL_obj
                      var va_arg2_try099911 EID 
                      /* Let:11 */{ 
                        var g1000UU *CompileCCast  
                        /* noccur = 1 */
                        var g1000UU_try100112 EID 
                        /* Let:12 */{ 
                          var _CL_obj *CompileCCast   = To_CompileCCast(new(CompileCCast).Is(C_Compile_C_cast))
                          /* noccur = 10 */
                          /* update:13 */{ 
                            var va_arg1 *CompileCCast  
                            var va_arg2 *ClaireAny  
                            va_arg1 = _CL_obj
                            var va_arg2_try100214 EID 
                            /* Let:14 */{ 
                              var g1003UU *Language.Call  
                              /* noccur = 1 */
                              /* Let:15 */{ 
                                var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                /* noccur = 5 */
                                _CL_obj.Selector = C_empty
                                _CL_obj.Args = MakeConstantList(v1.Id())
                                g1003UU = _CL_obj
                                /* Let-15 */} 
                              va_arg2_try100214 = Core.F_CALL(C_c_code,ARGS(EID{g1003UU.Id(),0},EID{x.Id(),0}))
                              /* Let-14 */} 
                            /* ERROR PROTECTION INSERTED (va_arg2-g1000UU_try100112) */
                            if ErrorIn(va_arg2_try100214) {g1000UU_try100112 = va_arg2_try100214
                            } else {
                            va_arg2 = ANY(va_arg2_try100214)
                            /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                            va_arg1.Arg = va_arg2
                            g1000UU_try100112 = va_arg2.ToEID()
                            }
                            /* update-13 */} 
                          /* ERROR PROTECTION INSERTED (g1000UU_try100112-g1000UU_try100112) */
                          if !ErrorIn(g1000UU_try100112) {
                          _CL_obj.SetArg = x
                          g1000UU_try100112 = EID{_CL_obj.Id(),0}
                          }
                          /* Let-12 */} 
                        /* ERROR PROTECTION INSERTED (g1000UU-va_arg2_try099911) */
                        if ErrorIn(g1000UU_try100112) {va_arg2_try099911 = g1000UU_try100112
                        } else {
                        g1000UU = To_CompileCCast(OBJ(g1000UU_try100112))
                        va_arg2_try099911 = F_Optimize_inner_select_Iteration(self,v2.Id(),v1.Id(),g1000UU.Id())
                        }
                        /* Let-11 */} 
                      /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                      if ErrorIn(va_arg2_try099911) {Result = va_arg2_try099911
                      } else {
                      va_arg2 = ANY(va_arg2_try099911)
                      /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                      va_arg1.Arg = va_arg2
                      Result = va_arg2.ToEID()
                      }
                      /* update-10 */} 
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    Result = EID{_CL_obj.Id(),0}
                    }
                    /* Let-9 */} 
                  } else {
                  /* Let:9 */{ 
                    var g1004UU *Language.Construct  
                    /* noccur = 1 */
                    if (x.Id() == C_set.Id()) /* If:10 */{ 
                      /* Let:11 */{ 
                        var _CL_obj *Language.Set   = Language.To_Set(new(Language.Set).Is(Language.C_Set))
                        /* noccur = 3 */
                        _CL_obj.Of = ToType(CEMPTY.Id())
                        g1004UU = Language.To_Construct(_CL_obj.Id())
                        /* Let-11 */} 
                      } else {
                      /* Let:11 */{ 
                        var _CL_obj *Language.List   = Language.To_List(new(Language.List).Is(Language.C_List))
                        /* noccur = 3 */
                        _CL_obj.Of = ToType(CEMPTY.Id())
                        g1004UU = Language.To_Construct(_CL_obj.Id())
                        /* Let-11 */} 
                      /* If-10 */} 
                    Result = F_Optimize_inner_select_Iteration(self,v2.Id(),sx,g1004UU.Id())
                    /* Let-9 */} 
                  /* If-8 */} 
                /* Let-7 */} 
              /* Let-6 */} 
            /* Let-5 */} 
          }
          /* Let-4 */} 
        }
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_code_select @ Iteration (throw: true) 
func E_Optimize_c_code_select_Iteration (self EID,x EID) EID { 
    return /*(sm for c_code_select @ Iteration= EID)*/ F_Optimize_c_code_select_Iteration(Language.To_Iteration(OBJ(self)),ToClass(OBJ(x)) )} 
  
// v3.2.01
// sub-procedure : creates the iteration over sx, adds the selected value (var) into l2 (set or list)
/* {1} OPT.The go function for: inner_select(self:Iteration,v2:any,sx:any,val:any) [] */
func F_Optimize_inner_select_Iteration (self *Language.Iteration ,v2 *ClaireAny ,sx *ClaireAny ,val *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var _CL_obj *Language.Let   = Language.To_Let(new(Language.Let).Is(Language.C_Let))
      /* noccur = 28 */
      _CL_obj.ClaireVar = To_Variable(v2)
      _CL_obj.Value = val
      /* update:3 */{ 
        var va_arg1 *Language.Let  
        var va_arg2 *ClaireAny  
        va_arg1 = _CL_obj
        var va_arg2_try10054 EID 
        /* Let:4 */{ 
          var _CL_obj *Language.Do   = Language.To_Do(new(Language.Do).Is(Language.C_Do))
          /* noccur = 21 */
          /* update:5 */{ 
            var va_arg1 *Language.Do  
            var va_arg2 *ClaireList  
            va_arg1 = _CL_obj
            var va_arg2_try10066 EID 
            /* Construct:6 */{ 
              var v_bag_arg *ClaireAny  
              va_arg2_try10066= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
              var v_bag_arg_try10077 EID 
              /* Let:7 */{ 
                var g1008UU *Language.For  
                /* noccur = 1 */
                /* Let:8 */{ 
                  var _CL_obj *Language.For   = Language.To_For(new(Language.For).Is(Language.C_For))
                  /* noccur = 17 */
                  _CL_obj.ClaireVar = self.ClaireVar
                  _CL_obj.SetArg = sx
                  /* update:9 */{ 
                    var va_arg1 *Language.Iteration  
                    var va_arg2 *ClaireAny  
                    va_arg1 = Language.To_Iteration(_CL_obj.Id())
                    /* Let:10 */{ 
                      var _CL_obj *Language.If   = Language.To_If(new(Language.If).Is(Language.C_If))
                      /* noccur = 11 */
                      _CL_obj.Test = self.Arg
                      /* update:11 */{ 
                        var va_arg1 *Language.If  
                        var va_arg2 *ClaireAny  
                        va_arg1 = _CL_obj
                        /* Let:12 */{ 
                          var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                          /* noccur = 5 */
                          _CL_obj.Selector = ToProperty(C_add_I.Id())
                          _CL_obj.Args = MakeConstantList(v2,self.ClaireVar.Id())
                          va_arg2 = _CL_obj.Id()
                          /* Let-12 */} 
                        /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                        va_arg1.Arg = va_arg2
                        /* update-11 */} 
                      va_arg2 = _CL_obj.Id()
                      /* Let-10 */} 
                    /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                    va_arg1.Arg = va_arg2
                    /* update-9 */} 
                  g1008UU = _CL_obj
                  /* Let-8 */} 
                v_bag_arg_try10077 = Core.F_CALL(C_c_code,ARGS(EID{g1008UU.Id(),0},EID{C_any.Id(),0}))
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try10066) */
              if ErrorIn(v_bag_arg_try10077) {va_arg2_try10066 = v_bag_arg_try10077
              } else {
              v_bag_arg = ANY(v_bag_arg_try10077)
              ToList(OBJ(va_arg2_try10066)).AddFast(v_bag_arg)
              ToList(OBJ(va_arg2_try10066)).AddFast(v2)}
              /* Construct-6 */} 
            /* ERROR PROTECTION INSERTED (va_arg2-va_arg2_try10054) */
            if ErrorIn(va_arg2_try10066) {va_arg2_try10054 = va_arg2_try10066
            } else {
            va_arg2 = ToList(OBJ(va_arg2_try10066))
            /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
            va_arg1.Args = va_arg2
            va_arg2_try10054 = EID{va_arg2.Id(),0}
            }
            /* update-5 */} 
          /* ERROR PROTECTION INSERTED (va_arg2_try10054-va_arg2_try10054) */
          if !ErrorIn(va_arg2_try10054) {
          va_arg2_try10054 = EID{_CL_obj.Id(),0}
          }
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (va_arg2-Result) */
        if ErrorIn(va_arg2_try10054) {Result = va_arg2_try10054
        } else {
        va_arg2 = ANY(va_arg2_try10054)
        /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
        va_arg1.Arg = va_arg2
        Result = va_arg2.ToEID()
        }
        /* update-3 */} 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = EID{_CL_obj.Id(),0}
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: inner_select @ Iteration (throw: true) 
func E_Optimize_inner_select_Iteration (self EID,v2 EID,sx EID,val EID) EID { 
    return /*(sm for inner_select @ Iteration= EID)*/ F_Optimize_inner_select_Iteration(Language.To_Iteration(OBJ(self)),
      ANY(v2),
      ANY(sx),
      ANY(val) )} 
  
// if (other = unknown : some) the result is either the variable or unknown
/* {1} OPT.The go function for: c_type(self:Exists) [] */
func F_c_type_Exists (self *Language.Exists ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var _Zt *ClaireAny  
      /* noccur = 3 */
      var _Zt_try10093 EID 
      _Zt_try10093 = Core.F_CALL(C_c_type,ARGS(self.SetArg.ToEID()))
      /* ERROR PROTECTION INSERTED (_Zt-Result) */
      if ErrorIn(_Zt_try10093) {Result = _Zt_try10093
      } else {
      _Zt = ANY(_Zt_try10093)
      F_Optimize_range_infers_for_Variable(self.ClaireVar,F_Optimize_pmember_type(ToType(_Zt)),ToType(_Zt))
      if (self.Other == CNULL) /* If:3 */{ 
        Result = EID{F_Optimize_extends_type(F_Optimize_pmember_type(ToType(_Zt))).Id(),0}
        } else {
        Result = EID{C_boolean.Id(),0}
        /* If-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_type @ Exists (throw: true) 
func E_c_type_Exists (self EID) EID { 
    return /*(sm for c_type @ Exists= EID)*/ F_c_type_Exists(Language.To_Exists(OBJ(self)) )} 
  
// boolean, or any U boolean ?
/* {1} OPT.The go function for: c_code(self:Exists,s:class) [] */
func F_c_code_Exists (self *Language.Exists ,s *ClaireClass ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var _Zt *ClaireAny  
      /* noccur = 2 */
      var _Zt_try10113 EID 
      _Zt_try10113 = Core.F_CALL(C_c_type,ARGS(self.SetArg.ToEID()))
      /* ERROR PROTECTION INSERTED (_Zt-Result) */
      if ErrorIn(_Zt_try10113) {Result = _Zt_try10113
      } else {
      _Zt = ANY(_Zt_try10113)
      F_Optimize_range_infers_for_Variable(self.ClaireVar,F_Optimize_pmember_type(ToType(_Zt)),ToType(_Zt))
      if (self.Other == CTRUE.Id()) /* If:3 */{ 
        /* Let:4 */{ 
          var g1012UU *Language.Call  
          /* noccur = 1 */
          /* Let:5 */{ 
            var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
            /* noccur = 23 */
            _CL_obj.Selector = Core.C_not
            /* update:6 */{ 
              var va_arg1 *Language.Call  
              var va_arg2 *ClaireList  
              va_arg1 = _CL_obj
              /* Construct:7 */{ 
                var v_bag_arg *ClaireAny  
                va_arg2= ToType(CEMPTY.Id()).EmptyList()
                /* Let:8 */{ 
                  var _CL_obj *Language.For   = Language.To_For(new(Language.For).Is(Language.C_For))
                  /* noccur = 17 */
                  _CL_obj.ClaireVar = self.ClaireVar
                  _CL_obj.SetArg = self.SetArg
                  /* update:9 */{ 
                    var va_arg1 *Language.Iteration  
                    var va_arg2 *ClaireAny  
                    va_arg1 = Language.To_Iteration(_CL_obj.Id())
                    /* Let:10 */{ 
                      var _CL_obj *Language.If   = Language.To_If(new(Language.If).Is(Language.C_If))
                      /* noccur = 11 */
                      /* update:11 */{ 
                        var va_arg1 *Language.If  
                        var va_arg2 *ClaireAny  
                        va_arg1 = _CL_obj
                        /* Let:12 */{ 
                          var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                          /* noccur = 5 */
                          _CL_obj.Selector = Core.C_not
                          _CL_obj.Args = MakeConstantList(self.Arg)
                          va_arg2 = _CL_obj.Id()
                          /* Let-12 */} 
                        /* ---------- now we compile update iClaire/test(va_arg1) := va_arg2 ------- */
                        va_arg1.Test = va_arg2
                        /* update-11 */} 
                      _CL_obj.Arg = Language.C_Return.Make(CTRUE.Id())
                      va_arg2 = _CL_obj.Id()
                      /* Let-10 */} 
                    /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                    va_arg1.Arg = va_arg2
                    /* update-9 */} 
                  v_bag_arg = _CL_obj.Id()
                  /* Let-8 */} 
                va_arg2.AddFast(v_bag_arg)/* Construct-7 */} 
              /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
              va_arg1.Args = va_arg2
              /* update-6 */} 
            g1012UU = _CL_obj
            /* Let-5 */} 
          Result = Core.F_CALL(C_c_code,ARGS(EID{g1012UU.Id(),0},EID{s.Id(),0}))
          /* Let-4 */} 
        /* If!3 */}  else if (self.Other == CNULL) /* If:3 */{ 
        /* Let:4 */{ 
          var v *ClaireVariable  
          /* noccur = 3 */
          /* Let:5 */{ 
            var g1013UU int 
            /* noccur = 1 */
            C_OPT.MaxVars = (C_OPT.MaxVars+1)
            g1013UU = 0
            v = F_Compile_Variable_I_symbol(F_append_symbol(self.ClaireVar.Pname,MakeString("_some").Id()),g1013UU,F_Optimize_extends_type(self.ClaireVar.Range).Id())
            /* Let-5 */} 
          /* Let:5 */{ 
            var g1014UU *Language.Let  
            /* noccur = 1 */
            /* Let:6 */{ 
              var _CL_obj *Language.Let   = Language.To_Let(new(Language.Let).Is(Language.C_Let))
              /* noccur = 22 */
              _CL_obj.ClaireVar = v
              _CL_obj.Value = CNULL
              /* update:7 */{ 
                var va_arg1 *Language.Let  
                var va_arg2 *ClaireAny  
                va_arg1 = _CL_obj
                /* Let:8 */{ 
                  var _CL_obj *Language.Do   = Language.To_Do(new(Language.Do).Is(Language.C_Do))
                  /* noccur = 15 */
                  /* update:9 */{ 
                    var va_arg1 *Language.Do  
                    var va_arg2 *ClaireList  
                    va_arg1 = _CL_obj
                    /* Construct:10 */{ 
                      var v_bag_arg *ClaireAny  
                      va_arg2= ToType(CEMPTY.Id()).EmptyList()
                      /* Let:11 */{ 
                        var _CL_obj *Language.For   = Language.To_For(new(Language.For).Is(Language.C_For))
                        /* noccur = 11 */
                        _CL_obj.ClaireVar = self.ClaireVar
                        _CL_obj.SetArg = self.SetArg
                        /* update:12 */{ 
                          var va_arg1 *Language.Iteration  
                          var va_arg2 *ClaireAny  
                          va_arg1 = Language.To_Iteration(_CL_obj.Id())
                          /* Let:13 */{ 
                            var _CL_obj *Language.If   = Language.To_If(new(Language.If).Is(Language.C_If))
                            /* noccur = 5 */
                            _CL_obj.Test = self.Arg
                            _CL_obj.Arg = Language.C_Return.Make(Language.C_Assign.Make(v.Id(),self.ClaireVar.Id()))
                            va_arg2 = _CL_obj.Id()
                            /* Let-13 */} 
                          /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                          va_arg1.Arg = va_arg2
                          /* update-12 */} 
                        v_bag_arg = _CL_obj.Id()
                        /* Let-11 */} 
                      va_arg2.AddFast(v_bag_arg)
                      va_arg2.AddFast(v.Id())/* Construct-10 */} 
                    /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                    va_arg1.Args = va_arg2
                    /* update-9 */} 
                  va_arg2 = _CL_obj.Id()
                  /* Let-8 */} 
                /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                va_arg1.Arg = va_arg2
                /* update-7 */} 
              g1014UU = _CL_obj
              /* Let-6 */} 
            Result = Core.F_CALL(C_c_code,ARGS(EID{g1014UU.Id(),0},EID{s.Id(),0}))
            /* Let-5 */} 
          /* Let-4 */} 
        } else {
        /* Let:4 */{ 
          var g1015UU *Language.Call  
          /* noccur = 1 */
          /* Let:5 */{ 
            var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
            /* noccur = 17 */
            _CL_obj.Selector = C_boolean_I
            /* update:6 */{ 
              var va_arg1 *Language.Call  
              var va_arg2 *ClaireList  
              va_arg1 = _CL_obj
              /* Construct:7 */{ 
                var v_bag_arg *ClaireAny  
                va_arg2= ToType(CEMPTY.Id()).EmptyList()
                /* Let:8 */{ 
                  var _CL_obj *Language.For   = Language.To_For(new(Language.For).Is(Language.C_For))
                  /* noccur = 11 */
                  _CL_obj.ClaireVar = self.ClaireVar
                  _CL_obj.SetArg = self.SetArg
                  /* update:9 */{ 
                    var va_arg1 *Language.Iteration  
                    var va_arg2 *ClaireAny  
                    va_arg1 = Language.To_Iteration(_CL_obj.Id())
                    /* Let:10 */{ 
                      var _CL_obj *Language.If   = Language.To_If(new(Language.If).Is(Language.C_If))
                      /* noccur = 5 */
                      _CL_obj.Test = self.Arg
                      _CL_obj.Arg = Language.C_Return.Make(CTRUE.Id())
                      va_arg2 = _CL_obj.Id()
                      /* Let-10 */} 
                    /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                    va_arg1.Arg = va_arg2
                    /* update-9 */} 
                  v_bag_arg = _CL_obj.Id()
                  /* Let-8 */} 
                va_arg2.AddFast(v_bag_arg)/* Construct-7 */} 
              /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
              va_arg1.Args = va_arg2
              /* update-6 */} 
            g1015UU = _CL_obj
            /* Let-5 */} 
          Result = Core.F_CALL(C_c_code,ARGS(EID{g1015UU.Id(),0},EID{s.Id(),0}))
          /* Let-4 */} 
        /* If-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_code @ Exists (throw: true) 
func E_c_code_Exists (self EID,s EID) EID { 
    return /*(sm for c_code @ Exists= EID)*/ F_c_code_Exists(Language.To_Exists(OBJ(self)),ToClass(OBJ(s)) )} 
  
// exists
// am Image builds a set
/* {1} OPT.The go function for: c_type(self:Image) [] */
func F_c_type_Image (self *Language.Image ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var _Zt *ClaireType  
      /* noccur = 2 */
      var _Zt_try10163 EID 
      /* Let:3 */{ 
        var g1017UU *ClaireType  
        /* noccur = 1 */
        var g1017UU_try10184 EID 
        g1017UU_try10184 = Core.F_CALL(C_c_type,ARGS(self.SetArg.ToEID()))
        /* ERROR PROTECTION INSERTED (g1017UU-_Zt_try10163) */
        if ErrorIn(g1017UU_try10184) {_Zt_try10163 = g1017UU_try10184
        } else {
        g1017UU = ToType(OBJ(g1017UU_try10184))
        _Zt_try10163 = EID{F_Optimize_ptype_type(g1017UU).Id(),0}
        }
        /* Let-3 */} 
      /* ERROR PROTECTION INSERTED (_Zt-Result) */
      if ErrorIn(_Zt_try10163) {Result = _Zt_try10163
      } else {
      _Zt = ToType(OBJ(_Zt_try10163))
      F_Optimize_range_infers_for_Variable(self.ClaireVar,F_Optimize_pmember_type(_Zt),_Zt)
      if (self.Of.Id() != CNULL) /* If:3 */{ 
        Result = EID{Core.F_param_I_class(C_set,self.Of).Id(),0}
        } else {
        /* Let:4 */{ 
          var g1019UU *ClaireType  
          /* noccur = 1 */
          var g1019UU_try10205 EID 
          g1019UU_try10205 = Core.F_CALL(C_c_type,ARGS(self.Arg.ToEID()))
          /* ERROR PROTECTION INSERTED (g1019UU-Result) */
          if ErrorIn(g1019UU_try10205) {Result = g1019UU_try10205
          } else {
          g1019UU = ToType(OBJ(g1019UU_try10205))
          Result = EID{Core.F_nth_class1(C_set,g1019UU).Id(),0}
          }
          /* Let-4 */} 
        /* If-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_type @ Image (throw: true) 
func E_c_type_Image (self EID) EID { 
    return /*(sm for c_type @ Image= EID)*/ F_c_type_Image(Language.To_Image(OBJ(self)) )} 
  
/* {1} OPT.The go function for: c_type(self:Select) [] */
func F_c_type_Select (self *Language.Select ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var _Zt *ClaireAny  
      /* noccur = 2 */
      var _Zt_try10213 EID 
      _Zt_try10213 = Core.F_CALL(C_c_type,ARGS(self.SetArg.ToEID()))
      /* ERROR PROTECTION INSERTED (_Zt-Result) */
      if ErrorIn(_Zt_try10213) {Result = _Zt_try10213
      } else {
      _Zt = ANY(_Zt_try10213)
      F_Optimize_range_infers_for_Variable(self.ClaireVar,F_Optimize_pmember_type(ToType(_Zt)),ToType(_Zt))
      if (self.Of.Id() != CNULL) /* If:3 */{ 
        Result = EID{Core.F_param_I_class(C_set,self.Of).Id(),0}
        } else {
        /* Let:4 */{ 
          var g1022UU *ClaireType  
          /* noccur = 1 */
          var g1022UU_try10235 EID 
          /* Let:5 */{ 
            var g1024UU *ClaireType  
            /* noccur = 1 */
            var g1024UU_try10256 EID 
            g1024UU_try10256 = Core.F_CALL(C_c_type,ARGS(self.SetArg.ToEID()))
            /* ERROR PROTECTION INSERTED (g1024UU-g1022UU_try10235) */
            if ErrorIn(g1024UU_try10256) {g1022UU_try10235 = g1024UU_try10256
            } else {
            g1024UU = ToType(OBJ(g1024UU_try10256))
            g1022UU_try10235 = EID{F_Optimize_pmember_type(g1024UU).Id(),0}
            }
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (g1022UU-Result) */
          if ErrorIn(g1022UU_try10235) {Result = g1022UU_try10235
          } else {
          g1022UU = ToType(OBJ(g1022UU_try10235))
          Result = EID{Core.F_nth_class1(C_set,g1022UU).Id(),0}
          }
          /* Let-4 */} 
        /* If-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_type @ Select (throw: true) 
func E_c_type_Select (self EID) EID { 
    return /*(sm for c_type @ Select= EID)*/ F_c_type_Select(Language.To_Select(OBJ(self)) )} 
  
// new in v3.1.06 : proper type inference !
/* {1} OPT.The go function for: c_type(self:Lselect) [] */
func F_c_type_Lselect (self *Language.Lselect ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var _Zt *ClaireAny  
      /* noccur = 2 */
      var _Zt_try10263 EID 
      _Zt_try10263 = Core.F_CALL(C_c_type,ARGS(self.SetArg.ToEID()))
      /* ERROR PROTECTION INSERTED (_Zt-Result) */
      if ErrorIn(_Zt_try10263) {Result = _Zt_try10263
      } else {
      _Zt = ANY(_Zt_try10263)
      F_Optimize_range_infers_for_Variable(self.ClaireVar,F_Optimize_pmember_type(ToType(_Zt)),ToType(_Zt))
      if (self.Of.Id() != CNULL) /* If:3 */{ 
        Result = EID{Core.F_param_I_class(C_list,self.Of).Id(),0}
        } else {
        /* Let:4 */{ 
          var g1027UU *ClaireType  
          /* noccur = 1 */
          var g1027UU_try10285 EID 
          /* Let:5 */{ 
            var g1029UU *ClaireType  
            /* noccur = 1 */
            var g1029UU_try10306 EID 
            g1029UU_try10306 = Core.F_CALL(C_c_type,ARGS(self.SetArg.ToEID()))
            /* ERROR PROTECTION INSERTED (g1029UU-g1027UU_try10285) */
            if ErrorIn(g1029UU_try10306) {g1027UU_try10285 = g1029UU_try10306
            } else {
            g1029UU = ToType(OBJ(g1029UU_try10306))
            g1027UU_try10285 = EID{F_Optimize_pmember_type(g1029UU).Id(),0}
            }
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (g1027UU-Result) */
          if ErrorIn(g1027UU_try10285) {Result = g1027UU_try10285
          } else {
          g1027UU = ToType(OBJ(g1027UU_try10285))
          Result = EID{Core.F_nth_class1(C_list,g1027UU).Id(),0}
          }
          /* Let-4 */} 
        /* If-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_type @ Lselect (throw: true) 
func E_c_type_Lselect (self EID) EID { 
    return /*(sm for c_type @ Lselect= EID)*/ F_c_type_Lselect(Language.To_Lselect(OBJ(self)) )} 
  
//______________________  while/until  __________________________________
// similar to a For
/* {1} OPT.The go function for: c_type(self:While) [] */
func F_c_type_While (self *Language.While ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var g1031UU *ClaireType  
      /* noccur = 1 */
      var g1031UU_try10323 EID 
      g1031UU_try10323 = F_Compile_return_type_any(self.Arg)
      /* ERROR PROTECTION INSERTED (g1031UU-Result) */
      if ErrorIn(g1031UU_try10323) {Result = g1031UU_try10323
      } else {
      g1031UU = ToType(OBJ(g1031UU_try10323))
      Result = EID{F_Optimize_infers_from_type(g1031UU,self.Id()).Id(),0}
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_type @ While (throw: true) 
func E_c_type_While (self EID) EID { 
    return /*(sm for c_type @ While= EID)*/ F_c_type_While(Language.To_While(OBJ(self)) )} 
  
/* {1} OPT.The go function for: c_code(self:While,s:class) [] */
func F_c_code_While (self *Language.While ,s *ClaireClass ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var _CL_obj *Language.While   = Language.To_While(new(Language.While).Is(Language.C_While))
      /* noccur = 7 */
      /* update:3 */{ 
        var va_arg1 *Language.While  
        var va_arg2 *ClaireAny  
        va_arg1 = _CL_obj
        var va_arg2_try10334 EID 
        va_arg2_try10334 = F_Optimize_c_boolean_any(self.Test)
        /* ERROR PROTECTION INSERTED (va_arg2-Result) */
        if ErrorIn(va_arg2_try10334) {Result = va_arg2_try10334
        } else {
        va_arg2 = ANY(va_arg2_try10334)
        /* ---------- now we compile update iClaire/test(va_arg1) := va_arg2 ------- */
        va_arg1.Test = va_arg2
        Result = va_arg2.ToEID()
        }
        /* update-3 */} 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      /* update:3 */{ 
        var va_arg1 *Language.While  
        var va_arg2 *ClaireAny  
        va_arg1 = _CL_obj
        var va_arg2_try10344 EID 
        va_arg2_try10344 = Core.F_CALL(C_c_code,ARGS(self.Arg.ToEID(),EID{C_void.Id(),0}))
        /* ERROR PROTECTION INSERTED (va_arg2-Result) */
        if ErrorIn(va_arg2_try10344) {Result = va_arg2_try10344
        } else {
        va_arg2 = ANY(va_arg2_try10344)
        /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
        va_arg1.Arg = va_arg2
        Result = va_arg2.ToEID()
        }
        /* update-3 */} 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      _CL_obj.Other = self.Other
      Result = EID{_CL_obj.Id(),0}
      }}
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_code @ While (throw: true) 
func E_c_code_While (self EID,s EID) EID { 
    return /*(sm for c_code @ While= EID)*/ F_c_code_While(Language.To_While(OBJ(self)),ToClass(OBJ(s)) )} 
  
//      if (s != void & s != any)
//        (//[5] ... insert a to_C with s = ~S for ~S // s,self,
//         to_C(arg = r, set_arg = s)) // v3.3
//      else r)))
// *********************************************************************
// *     Part 6: Iterate                                               *
// *********************************************************************
// finds the right restriction of Iterate
// Iterate applies to the non-evaluated types (meta level)
/* {1} OPT.The go function for: Iterate!(self:Iteration) [] */
func F_Optimize_Iterate_I_Iteration (self *Language.Iteration ) *ClaireAny  { 
    // use function body compiling 
return  F_Optimize_restriction_I_property(Language.C_Iterate,MakeConstantList(MakeConstantSet(self.SetArg).Id(),MakeConstantSet(self.ClaireVar.Id()).Id(),C_any.Id()),CTRUE)
    } 
  
// The EID go function for: Iterate! @ Iteration (throw: false) 
func E_Optimize_Iterate_I_Iteration (self EID) EID { 
    return /*(sm for Iterate! @ Iteration= any)*/ F_Optimize_Iterate_I_Iteration(Language.To_Iteration(OBJ(self)) ).ToEID()} 
  
// iteration methods
// note the beauty of this: we only apply the code transformation if
// we actually get a constant Interval
/* {1} OPT.The go function for: iterate(x:Interval,v:Variable[range:(subtype[integer])],e:any) [] */
func F_iterate_Interval (x *ClaireInterval ,v *ClaireVariable ,e *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var v *ClaireAny  
      /* noccur = 3 */
      var v_try10353 EID 
      v_try10353 = F_eval_any2(MakeInteger(x.Arg1).Id(),C_Interval)
      /* ERROR PROTECTION INSERTED (v-Result) */
      if ErrorIn(v_try10353) {Result = v_try10353
      } else {
      v = ANY(v_try10353)
      /* Let:3 */{ 
        var _Zmax int 
        /* noccur = 1 */
        var _Zmax_try10364 EID 
        _Zmax_try10364 = F_eval_any2(MakeInteger(x.Arg2).Id(),C_Interval)
        /* ERROR PROTECTION INSERTED (_Zmax-Result) */
        if ErrorIn(_Zmax_try10364) {Result = _Zmax_try10364
        } else {
        _Zmax = INT(_Zmax_try10364)
        Result= EID{CFALSE.Id(),0}
        for (ToInteger(v).Value <= _Zmax) /* while:4 */{ 
          
          v = ANY(Core.F_CALL(ToProperty(Core.C__plus.Id()),ARGS(v.ToEID(),EID{C__INT,IVAL(1)})))
          /* while-4 */} 
        }
        /* Let-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: iterate @ Interval (throw: true) 
func E_iterate_Interval (x EID,v EID,e EID) EID { 
    return /*(sm for iterate @ Interval= EID)*/ F_iterate_Interval(To_Interval(OBJ(x)),To_Variable(OBJ(v)),ANY(e) )} 
  
/* {1} OPT.The go function for: iterate(x:array,v:Variable,e:any) [] */
func F_iterate_array (x *ClaireList ,v *ClaireVariable ,e *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var _Zi int  = 1
      /* noccur = 4 */
      /* Let:3 */{ 
        var _Za *ClaireList   = x
        /* noccur = 2 */
        /* Let:4 */{ 
          var _Zmax int  = _Za.Length()
          /* noccur = 1 */
          Result= EID{CFALSE.Id(),0}
          for (_Zi <= _Zmax) /* while:5 */{ 
            var void_try6 EID 
            _ = void_try6
            /* Let:6 */{ 
              var v *ClaireAny  
              /* noccur = 0 */
              _ = v
              var v_try10377 EID 
              v_try10377 = Core.F_nth_array(_Za,_Zi)
              /* ERROR PROTECTION INSERTED (v-void_try6) */
              if ErrorIn(v_try10377) {void_try6 = v_try10377
              } else {
              v = ANY(v_try10377)
              
              _Zi = (_Zi+1)
              void_try6 = EID{C__INT,IVAL(_Zi)}
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (void_try6-Result) */
            if ErrorIn(void_try6) {Result = void_try6
            Result = void_try6
            break
            } else {
            /* while-5 */} 
          }
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: iterate @ array (throw: true) 
func E_iterate_array (x EID,v EID,e EID) EID { 
    return /*(sm for iterate @ array= EID)*/ F_iterate_array(ToArray(OBJ(x)),To_Variable(OBJ(v)),ANY(e) )} 
  
/* {1} OPT.The go function for: Iterate(x:class,v:Variable,e:any) [] */
func F_Iterate_class (x *ClaireClass ,v *ClaireVariable ,e *ClaireAny ) *ClaireAny  { 
    // procedure body with s =  
var Result *ClaireAny  
    /* For:2 */{ 
      var _Zv_1 *ClaireAny  
      _ = _Zv_1
      Result= CFALSE.Id()
      for _,_Zv_1 = range(x.Descendents.Values)/* loop:3 */{ 
        /* Let:4 */{ 
          var _Zv_2 *ClaireBoolean  
          /* noccur = 2 */
          /* For:5 */{ 
            var v *ClaireAny  
            _ = v
            _Zv_2= CFALSE
            var v_support *ClaireList  
            v_support = ToClass(_Zv_1).Instances
            v_len := v_support.Length()
            for i_it := 0; i_it < v_len; i_it++ { 
              v = v_support.At(i_it)
              
              /* loop-6 */} 
            /* For-5 */} 
          if (_Zv_2 == CTRUE) /* If:5 */{ 
             /*v = Result, s =any*/
Result = _Zv_2.Id()
            break
            /* If-5 */} 
          /* Let-4 */} 
        /* loop-3 */} 
      /* For-2 */} 
    return Result} 
  
// The EID go function for: Iterate @ class (throw: false) 
func E_Iterate_class (x EID,v EID,e EID) EID { 
    return /*(sm for Iterate @ class= any)*/ F_Iterate_class(ToClass(OBJ(x)),To_Variable(OBJ(v)),ANY(e) ).ToEID()} 
  
/* {1} OPT.The go function for: Iterate(x:any,v:Variable,e:any) [] */
func F_Iterate_any1 (x *ClaireAny ,v *ClaireVariable ,e *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var v *ClaireAny  
      /* noccur = 3 */
      var v_try10383 EID 
      v_try10383 = EVAL(ToList(OBJ(Core.F_CALL(C_args,ARGS(x.ToEID())))).At(1-1))
      /* ERROR PROTECTION INSERTED (v-Result) */
      if ErrorIn(v_try10383) {Result = v_try10383
      } else {
      v = ANY(v_try10383)
      /* Let:3 */{ 
        var _Zmax *ClaireAny  
        /* noccur = 1 */
        var _Zmax_try10394 EID 
        _Zmax_try10394 = EVAL(ToList(OBJ(Core.F_CALL(C_args,ARGS(x.ToEID())))).At(2-1))
        /* ERROR PROTECTION INSERTED (_Zmax-Result) */
        if ErrorIn(_Zmax_try10394) {Result = _Zmax_try10394
        } else {
        _Zmax = ANY(_Zmax_try10394)
        Result= EID{CFALSE.Id(),0}
        for (ToBoolean(OBJ(Core.F_CALL(ToProperty(C__inf_equal.Id()),ARGS(v.ToEID(),_Zmax.ToEID())))) == CTRUE) /* while:4 */{ 
          
          v = ANY(Core.F_CALL(ToProperty(Core.C__plus.Id()),ARGS(v.ToEID(),EID{C__INT,IVAL(1)})))
          /* while-4 */} 
        }
        /* Let-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: Iterate @ list<type_expression>(..[tuple(integer,integer)], Variable, any) (throw: true) 
func E_Iterate_any1 (x EID,v EID,e EID) EID { 
    return /*(sm for Iterate @ list<type_expression>(..[tuple(integer,integer)], Variable, any)= EID)*/ F_Iterate_any1(ANY(x),To_Variable(OBJ(v)),ANY(e) )} 
  
/* {1} OPT.The go function for: Iterate(x:Lselect,v:Variable,e:any) [] */
func F_Iterate_Lselect (x *Language.Lselect ,v *ClaireVariable ,e *ClaireAny ) EID { 
    var Result EID 
    /* For:2 */{ 
      var v *ClaireAny  
      _ = v
      Result= EID{CFALSE.Id(),0}
      var v_support *ClaireList  
      var v_support_try10403 EID 
      /* Let:3 */{ 
        var g1041UU *ClaireAny  
        /* noccur = 1 */
        var g1041UU_try10424 EID 
        g1041UU_try10424 = EVAL(x.SetArg)
        /* ERROR PROTECTION INSERTED (g1041UU-v_support_try10403) */
        if ErrorIn(g1041UU_try10424) {v_support_try10403 = g1041UU_try10424
        } else {
        g1041UU = ANY(g1041UU_try10424)
        v_support_try10403 = Core.F_enumerate_any(g1041UU)
        }
        /* Let-3 */} 
      /* ERROR PROTECTION INSERTED (v_support-Result) */
      if ErrorIn(v_support_try10403) {Result = v_support_try10403
      } else {
      v_support = ToList(OBJ(v_support_try10403))
      v_len := v_support.Length()
      for i_it := 0; i_it < v_len; i_it++ { 
        v = v_support.At(i_it)
        var void_try4 EID 
        _ = void_try4
        var g1043I *ClaireBoolean  
        var g1043I_try10444 EID 
        /* Let:4 */{ 
          var g1045UU *ClaireAny  
          /* noccur = 1 */
          var g1045UU_try10465 EID 
          g1045UU_try10465 = EVAL(Language.F_substitution_any(x.Arg,x.ClaireVar,v))
          /* ERROR PROTECTION INSERTED (g1045UU-g1043I_try10444) */
          if ErrorIn(g1045UU_try10465) {g1043I_try10444 = g1045UU_try10465
          } else {
          g1045UU = ANY(g1045UU_try10465)
          g1043I_try10444 = EID{F_boolean_I_any(g1045UU).Id(),0}
          }
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (g1043I-void_try4) */
        if ErrorIn(g1043I_try10444) {void_try4 = g1043I_try10444
        } else {
        g1043I = ToBoolean(OBJ(g1043I_try10444))
        if (g1043I == CTRUE) /* If:4 */{ 
          void_try4 = e.ToEID()
          } else {
          void_try4 = EID{CFALSE.Id(),0}
          /* If-4 */} 
        }
        /* ERROR PROTECTION INSERTED (void_try4-Result) */
        if ErrorIn(void_try4) {Result = void_try4
        Result = void_try4
        break
        } else {
        }}
        /* loop-3 */} 
      /* For-2 */} 
    return Result} 
  
// The EID go function for: Iterate @ Lselect (throw: true) 
func E_Iterate_Lselect (x EID,v EID,e EID) EID { 
    return /*(sm for Iterate @ Lselect= EID)*/ F_Iterate_Lselect(Language.To_Lselect(OBJ(x)),To_Variable(OBJ(v)),ANY(e) )} 
  
/* {1} OPT.The go function for: Iterate(x:Select,v:Variable,e:any) [] */
func F_Iterate_Select (x *Language.Select ,v *ClaireVariable ,e *ClaireAny ) EID { 
    var Result EID 
    /* For:2 */{ 
      var v *ClaireAny  
      _ = v
      Result= EID{CFALSE.Id(),0}
      var v_support *ClaireList  
      var v_support_try10473 EID 
      /* Let:3 */{ 
        var g1048UU *ClaireAny  
        /* noccur = 1 */
        var g1048UU_try10494 EID 
        g1048UU_try10494 = EVAL(x.SetArg)
        /* ERROR PROTECTION INSERTED (g1048UU-v_support_try10473) */
        if ErrorIn(g1048UU_try10494) {v_support_try10473 = g1048UU_try10494
        } else {
        g1048UU = ANY(g1048UU_try10494)
        v_support_try10473 = Core.F_enumerate_any(g1048UU)
        }
        /* Let-3 */} 
      /* ERROR PROTECTION INSERTED (v_support-Result) */
      if ErrorIn(v_support_try10473) {Result = v_support_try10473
      } else {
      v_support = ToList(OBJ(v_support_try10473))
      v_len := v_support.Length()
      for i_it := 0; i_it < v_len; i_it++ { 
        v = v_support.At(i_it)
        var void_try4 EID 
        _ = void_try4
        var g1050I *ClaireBoolean  
        var g1050I_try10514 EID 
        /* Let:4 */{ 
          var g1052UU *ClaireAny  
          /* noccur = 1 */
          var g1052UU_try10535 EID 
          g1052UU_try10535 = EVAL(Language.F_substitution_any(x.Arg,x.ClaireVar,v))
          /* ERROR PROTECTION INSERTED (g1052UU-g1050I_try10514) */
          if ErrorIn(g1052UU_try10535) {g1050I_try10514 = g1052UU_try10535
          } else {
          g1052UU = ANY(g1052UU_try10535)
          g1050I_try10514 = EID{F_boolean_I_any(g1052UU).Id(),0}
          }
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (g1050I-void_try4) */
        if ErrorIn(g1050I_try10514) {void_try4 = g1050I_try10514
        } else {
        g1050I = ToBoolean(OBJ(g1050I_try10514))
        if (g1050I == CTRUE) /* If:4 */{ 
          void_try4 = e.ToEID()
          } else {
          void_try4 = EID{CFALSE.Id(),0}
          /* If-4 */} 
        }
        /* ERROR PROTECTION INSERTED (void_try4-Result) */
        if ErrorIn(void_try4) {Result = void_try4
        Result = void_try4
        break
        } else {
        }}
        /* loop-3 */} 
      /* For-2 */} 
    return Result} 
  
// The EID go function for: Iterate @ Select (throw: true) 
func E_Iterate_Select (x EID,v EID,e EID) EID { 
    return /*(sm for Iterate @ Select= EID)*/ F_Iterate_Select(Language.To_Select(OBJ(x)),To_Variable(OBJ(v)),ANY(e) )} 
  
/* {1} OPT.The go function for: Iterate(x:Collect,v:Variable,e:any) [] */
func F_Iterate_Collect (x *Language.Collect ,v *ClaireVariable ,e *ClaireAny ) EID { 
    var Result EID 
    /* For:2 */{ 
      var C_Zv *ClaireAny  
      _ = C_Zv
      Result= EID{CFALSE.Id(),0}
      var C_Zv_support *ClaireList  
      var C_Zv_support_try10543 EID 
      /* Let:3 */{ 
        var g1055UU *ClaireAny  
        /* noccur = 1 */
        var g1055UU_try10564 EID 
        g1055UU_try10564 = EVAL(x.SetArg)
        /* ERROR PROTECTION INSERTED (g1055UU-C_Zv_support_try10543) */
        if ErrorIn(g1055UU_try10564) {C_Zv_support_try10543 = g1055UU_try10564
        } else {
        g1055UU = ANY(g1055UU_try10564)
        C_Zv_support_try10543 = Core.F_enumerate_any(g1055UU)
        }
        /* Let-3 */} 
      /* ERROR PROTECTION INSERTED (C_Zv_support-Result) */
      if ErrorIn(C_Zv_support_try10543) {Result = C_Zv_support_try10543
      } else {
      C_Zv_support = ToList(OBJ(C_Zv_support_try10543))
      C_Zv_len := C_Zv_support.Length()
      for i_it := 0; i_it < C_Zv_len; i_it++ { 
        C_Zv = C_Zv_support.At(i_it)
        var void_try4 EID 
        _ = void_try4
        /* Let:4 */{ 
          var v *ClaireAny  
          /* noccur = 0 */
          _ = v
          var v_try10575 EID 
          v_try10575 = EVAL(Language.F_substitution_any(x.Arg,x.ClaireVar,C_Zv))
          /* ERROR PROTECTION INSERTED (v-void_try4) */
          if ErrorIn(v_try10575) {void_try4 = v_try10575
          } else {
          v = ANY(v_try10575)
          void_try4 = e.ToEID()
          }
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (void_try4-Result) */
        if ErrorIn(void_try4) {Result = void_try4
        Result = void_try4
        break
        } else {
        }}
        /* loop-3 */} 
      /* For-2 */} 
    return Result} 
  
// The EID go function for: Iterate @ Collect (throw: true) 
func E_Iterate_Collect (x EID,v EID,e EID) EID { 
    return /*(sm for Iterate @ Collect= EID)*/ F_Iterate_Collect(Language.To_Collect(OBJ(x)),To_Variable(OBJ(v)),ANY(e) )} 
  
/* {1} OPT.The go function for: Iterate(x:any,v:Variable,e:any) [] */
func F_Iterate_any2 (x *ClaireAny ,v *ClaireVariable ,e *ClaireAny ) EID { 
    var Result EID 
    /* For:2 */{ 
      var v *ClaireAny  
      _ = v
      Result= EID{CFALSE.Id(),0}
      var v_support *ClaireList  
      var v_support_try10583 EID 
      /* Let:3 */{ 
        var g1059UU *ClaireAny  
        /* noccur = 1 */
        var g1059UU_try10604 EID 
        g1059UU_try10604 = EVAL(ToList(OBJ(Core.F_CALL(C_args,ARGS(x.ToEID())))).At(1-1))
        /* ERROR PROTECTION INSERTED (g1059UU-v_support_try10583) */
        if ErrorIn(g1059UU_try10604) {v_support_try10583 = g1059UU_try10604
        } else {
        g1059UU = ANY(g1059UU_try10604)
        v_support_try10583 = Core.F_enumerate_any(g1059UU)
        }
        /* Let-3 */} 
      /* ERROR PROTECTION INSERTED (v_support-Result) */
      if ErrorIn(v_support_try10583) {Result = v_support_try10583
      } else {
      v_support = ToList(OBJ(v_support_try10583))
      v_len := v_support.Length()
      for i_it := 0; i_it < v_len; i_it++ { 
        v = v_support.At(i_it)
        var void_try4 EID 
        _ = void_try4
        var g1061I *ClaireBoolean  
        var g1061I_try10624 EID 
        /* Let:4 */{ 
          var g1063UU *ClaireAny  
          /* noccur = 1 */
          var g1063UU_try10645 EID 
          g1063UU_try10645 = EVAL(ToList(OBJ(Core.F_CALL(C_args,ARGS(x.ToEID())))).At(2-1))
          /* ERROR PROTECTION INSERTED (g1063UU-g1061I_try10624) */
          if ErrorIn(g1063UU_try10645) {g1061I_try10624 = g1063UU_try10645
          } else {
          g1063UU = ANY(g1063UU_try10645)
          g1061I_try10624 = EID{Core.F__I_equal_any(v,g1063UU).Id(),0}
          }
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (g1061I-void_try4) */
        if ErrorIn(g1061I_try10624) {void_try4 = g1061I_try10624
        } else {
        g1061I = ToBoolean(OBJ(g1061I_try10624))
        if (g1061I == CTRUE) /* If:4 */{ 
          void_try4 = e.ToEID()
          } else {
          void_try4 = EID{CFALSE.Id(),0}
          /* If-4 */} 
        }
        /* ERROR PROTECTION INSERTED (void_try4-Result) */
        if ErrorIn(void_try4) {Result = void_try4
        Result = void_try4
        break
        } else {
        }}
        /* loop-3 */} 
      /* For-2 */} 
    return Result} 
  
// The EID go function for: Iterate @ list<type_expression>(but[tuple(any,any)], Variable, any) (throw: true) 
func E_Iterate_any2 (x EID,v EID,e EID) EID { 
    return /*(sm for Iterate @ list<type_expression>(but[tuple(any,any)], Variable, any)= EID)*/ F_Iterate_any2(ANY(x),To_Variable(OBJ(v)),ANY(e) )} 
  
/* {1} OPT.The go function for: Iterate(x:any,v:Variable,e:any) [] */
func F_Iterate_any3 (x *ClaireAny ,v *ClaireVariable ,e *ClaireAny ) EID { 
    var Result EID 
    /* For:2 */{ 
      var v *ClaireAny  
      _ = v
      Result= EID{CFALSE.Id(),0}
      var v_support *ClaireList  
      var v_support_try10653 EID 
      /* Let:3 */{ 
        var g1066UU *ClaireAny  
        /* noccur = 1 */
        var g1066UU_try10674 EID 
        g1066UU_try10674 = EVAL(ToList(OBJ(Core.F_CALL(C_args,ARGS(x.ToEID())))).At(1-1))
        /* ERROR PROTECTION INSERTED (g1066UU-v_support_try10653) */
        if ErrorIn(g1066UU_try10674) {v_support_try10653 = g1066UU_try10674
        } else {
        g1066UU = ANY(g1066UU_try10674)
        v_support_try10653 = Core.F_enumerate_any(g1066UU)
        }
        /* Let-3 */} 
      /* ERROR PROTECTION INSERTED (v_support-Result) */
      if ErrorIn(v_support_try10653) {Result = v_support_try10653
      } else {
      v_support = ToList(OBJ(v_support_try10653))
      v_len := v_support.Length()
      for i_it := 0; i_it < v_len; i_it++ { 
        v = v_support.At(i_it)
        
        }
        /* loop-3 */} 
      /* For-2 */} 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    /* For:2 */{ 
      var v *ClaireAny  
      _ = v
      Result= EID{CFALSE.Id(),0}
      var v_support *ClaireList  
      var v_support_try10683 EID 
      /* Let:3 */{ 
        var g1069UU *ClaireAny  
        /* noccur = 1 */
        var g1069UU_try10704 EID 
        g1069UU_try10704 = EVAL(ToList(OBJ(Core.F_CALL(C_args,ARGS(x.ToEID())))).At(2-1))
        /* ERROR PROTECTION INSERTED (g1069UU-v_support_try10683) */
        if ErrorIn(g1069UU_try10704) {v_support_try10683 = g1069UU_try10704
        } else {
        g1069UU = ANY(g1069UU_try10704)
        v_support_try10683 = Core.F_enumerate_any(g1069UU)
        }
        /* Let-3 */} 
      /* ERROR PROTECTION INSERTED (v_support-Result) */
      if ErrorIn(v_support_try10683) {Result = v_support_try10683
      } else {
      v_support = ToList(OBJ(v_support_try10683))
      v_len := v_support.Length()
      for i_it := 0; i_it < v_len; i_it++ { 
        v = v_support.At(i_it)
        
        }
        /* loop-3 */} 
      /* For-2 */} 
    }
    return Result} 
  
// The EID go function for: Iterate @ list<type_expression>(/+[tuple(any,any)], Variable, any) (throw: true) 
func E_Iterate_any3 (x EID,v EID,e EID) EID { 
    return /*(sm for Iterate @ list<type_expression>(/+[tuple(any,any)], Variable, any)= EID)*/ F_Iterate_any3(ANY(x),To_Variable(OBJ(v)),ANY(e) )} 
  