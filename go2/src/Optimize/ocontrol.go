/***** CLAIRE Compilation of file /Users/ycaseau/Dropbox/src/clairev4.0/src/compile/ocontrol.cl 
         [version 4.0.01 / safety 5] Saturday 10-30-2021 *****/

package Optimize
import (_ "fmt"
	. "Kernel"
	"Core"
	"Language"
	"Reader"
)

//-------- dumb function to prevent import errors --------
func import_g0835() { 
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
          var _Ztype_try08365 EID 
          /* Let:5 */{ 
            var g0837UU *ClaireType  
            /* noccur = 1 */
            var g0837UU_try08386 EID 
            g0837UU_try08386 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
            /* ERROR PROTECTION INSERTED (g0837UU-_Ztype_try08365) */
            if ErrorIn(g0837UU_try08386) {_Ztype_try08365 = g0837UU_try08386
            } else {
            g0837UU = ToType(OBJ(g0837UU_try08386))
            _Ztype_try08365 = EID{F_Optimize_ptype_type(g0837UU).Id(),0}
            }
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (_Ztype-Result) */
          if ErrorIn(_Ztype_try08365) {Result = _Ztype_try08365
          } else {
          _Ztype = ToType(OBJ(_Ztype_try08365))
          if (v.Isa.IsIn(C_Variable) != CTRUE) /* If:5 */{ 
            Result = Core.F_CALL(C_Compile_Cerror,ARGS(EID{MakeString("[213] ~S is not a variable").Id(),0},v.ToEID()))
            } else {
            Result = EID{CFALSE.Id(),0}
            /* If-5 */} 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          if (_Ztype.Included(ToType(OBJ(Core.F_CALL(C_range,ARGS(v.ToEID()))))) != CTRUE) /* If:5 */{ 
            var x_try08396 EID 
            x_try08396 = Core.F_CALL(C_Optimize_c_warn,ARGS(self.ClaireVar.ToEID(),x.ToEID(),EID{_Ztype.Id(),0}))
            /* ERROR PROTECTION INSERTED (x-Result) */
            if ErrorIn(x_try08396) {Result = x_try08396
            } else {
            x = ANY(x_try08396)
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
            var _Zarg_try08406 EID 
            _Zarg_try08406 = F_Compile_c_strict_code_any(x,F_Compile_psort_any(ANY(Core.F_CALL(C_range,ARGS(v.ToEID())))))
            /* ERROR PROTECTION INSERTED (_Zarg-Result) */
            if ErrorIn(_Zarg_try08406) {Result = _Zarg_try08406
            } else {
            _Zarg = ANY(_Zarg_try08406)
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
        var _Ztype_try08414 EID 
        /* Let:4 */{ 
          var g0842UU *ClaireType  
          /* noccur = 1 */
          var g0842UU_try08435 EID 
          g0842UU_try08435 = Core.F_CALL(C_c_type,ARGS(_Zv.ToEID()))
          /* ERROR PROTECTION INSERTED (g0842UU-_Ztype_try08414) */
          if ErrorIn(g0842UU_try08435) {_Ztype_try08414 = g0842UU_try08435
          } else {
          g0842UU = ToType(OBJ(g0842UU_try08435))
          _Ztype_try08414 = EID{F_Optimize_ptype_type(g0842UU).Id(),0}
          }
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (_Ztype-Result) */
        if ErrorIn(_Ztype_try08414) {Result = _Ztype_try08414
        } else {
        _Ztype = ToType(OBJ(_Ztype_try08414))
        if (F_boolean_I_any(self.ClaireVar.Range.Id()).Id() != CTRUE.Id()) /* If:4 */{ 
          Result = Core.F_CALL(C_Compile_Cerror,ARGS(EID{MakeString("[214] cannot assign ~S").Id(),0},EID{self.Id(),0}))
          } else {
          Result = EID{CFALSE.Id(),0}
          /* If-4 */} 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        if (_Ztype.Included(self.ClaireVar.Range) != CTRUE) /* If:4 */{ 
          var _Zv_try08445 EID 
          _Zv_try08445 = Core.F_CALL(C_c_code,ARGS(_Zv.ToEID(),EID{C_any.Id(),0}))
          /* ERROR PROTECTION INSERTED (_Zv-Result) */
          if ErrorIn(_Zv_try08445) {Result = _Zv_try08445
          } else {
          _Zv = ANY(_Zv_try08445)
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
            var va_arg2_try08456 EID 
            va_arg2_try08456 = Core.F_CALL(C_c_code,ARGS(EID{self.ClaireVar.Id(),0}))
            /* ERROR PROTECTION INSERTED (va_arg2-Result) */
            if ErrorIn(va_arg2_try08456) {Result = va_arg2_try08456
            } else {
            va_arg2 = Core.ToGlobalVariable(OBJ(va_arg2_try08456))
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
            var va_arg2_try08466 EID 
            if (F_Compile_nativeVar_ask_global_variable(self.ClaireVar) == CTRUE) /* If:6 */{ 
              va_arg2_try08466 = F_Compile_c_strict_code_any(_Zv,F_Compile_psort_any(self.ClaireVar.Range.Id()))
              } else {
              va_arg2_try08466 = Core.F_CALL(C_c_code,ARGS(_Zv.ToEID(),EID{C_any.Id(),0}))
              /* If-6 */} 
            /* ERROR PROTECTION INSERTED (va_arg2-Result) */
            if ErrorIn(va_arg2_try08466) {Result = va_arg2_try08466
            } else {
            va_arg2 = ANY(va_arg2_try08466)
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
        var va_arg2_try08474 EID 
        /* Iteration:4 */{ 
          var v_list4 *ClaireList  
          var x *ClaireAny  
          var v_local4 *ClaireAny  
          v_list4 = self.Args
          va_arg2_try08474 = EID{CreateList(ToType(CEMPTY.Id()),v_list4.Length()).Id(),0}
          for CLcount := 0; CLcount < v_list4.Length(); CLcount++{ 
            x = v_list4.At(CLcount)
            var v_local4_try08486 EID 
            var g0849I *ClaireBoolean  
            var g0849I_try08506 EID 
            /* Let:6 */{ 
              var g0851UU *ClaireType  
              /* noccur = 1 */
              var g0851UU_try08527 EID 
              g0851UU_try08527 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
              /* ERROR PROTECTION INSERTED (g0851UU-g0849I_try08506) */
              if ErrorIn(g0851UU_try08527) {g0849I_try08506 = g0851UU_try08527
              } else {
              g0851UU = ToType(OBJ(g0851UU_try08527))
              g0849I_try08506 = EID{Equal(g0851UU.Id(),C_void.Id()).Id(),0}
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (g0849I-v_local4_try08486) */
            if ErrorIn(g0849I_try08506) {v_local4_try08486 = g0849I_try08506
            } else {
            g0849I = ToBoolean(OBJ(g0849I_try08506))
            if (g0849I == CTRUE) /* If:6 */{ 
              v_local4_try08486 = Core.F_CALL(C_Compile_Cerror,ARGS(EID{MakeString("[206] void ~S in ~S").Id(),0},x.ToEID(),EID{self.Id(),0}))
              } else {
              v_local4_try08486 = EID{CFALSE.Id(),0}
              /* If-6 */} 
            }
            /* ERROR PROTECTION INSERTED (v_local4_try08486-v_local4_try08486) */
            if ErrorIn(v_local4_try08486) {va_arg2_try08474 = v_local4_try08486
            break
            } else {
            v_local4_try08486 = F_Optimize_c_boolean_any(x)
            /* ERROR PROTECTION INSERTED (v_local4_try08486-v_local4_try08486) */
            if ErrorIn(v_local4_try08486) {va_arg2_try08474 = v_local4_try08486
            break
            } else {
            }}
            {
            v_local4 = ANY(v_local4_try08486)
            ToList(OBJ(va_arg2_try08474)).PutAt(CLcount,v_local4)
            } 
          }
          /* Iteration-4 */} 
        /* ERROR PROTECTION INSERTED (va_arg2-Result) */
        if ErrorIn(va_arg2_try08474) {Result = va_arg2_try08474
        } else {
        va_arg2 = ToList(OBJ(va_arg2_try08474))
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
        var va_arg2_try08544 EID 
        /* Iteration:4 */{ 
          var v_list4 *ClaireList  
          var x *ClaireAny  
          var v_local4 *ClaireAny  
          v_list4 = self.Args
          va_arg2_try08544 = EID{CreateList(ToType(CEMPTY.Id()),v_list4.Length()).Id(),0}
          for CLcount := 0; CLcount < v_list4.Length(); CLcount++{ 
            x = v_list4.At(CLcount)
            var v_local4_try08556 EID 
            var g0856I *ClaireBoolean  
            var g0856I_try08576 EID 
            /* Let:6 */{ 
              var g0858UU *ClaireType  
              /* noccur = 1 */
              var g0858UU_try08597 EID 
              g0858UU_try08597 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
              /* ERROR PROTECTION INSERTED (g0858UU-g0856I_try08576) */
              if ErrorIn(g0858UU_try08597) {g0856I_try08576 = g0858UU_try08597
              } else {
              g0858UU = ToType(OBJ(g0858UU_try08597))
              g0856I_try08576 = EID{Equal(g0858UU.Id(),C_void.Id()).Id(),0}
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (g0856I-v_local4_try08556) */
            if ErrorIn(g0856I_try08576) {v_local4_try08556 = g0856I_try08576
            } else {
            g0856I = ToBoolean(OBJ(g0856I_try08576))
            if (g0856I == CTRUE) /* If:6 */{ 
              v_local4_try08556 = Core.F_CALL(C_Compile_Cerror,ARGS(EID{MakeString("[206] void ~S in ~S").Id(),0},x.ToEID(),EID{self.Id(),0}))
              } else {
              v_local4_try08556 = EID{CFALSE.Id(),0}
              /* If-6 */} 
            }
            /* ERROR PROTECTION INSERTED (v_local4_try08556-v_local4_try08556) */
            if ErrorIn(v_local4_try08556) {va_arg2_try08544 = v_local4_try08556
            break
            } else {
            v_local4_try08556 = F_Optimize_c_boolean_any(x)
            /* ERROR PROTECTION INSERTED (v_local4_try08556-v_local4_try08556) */
            if ErrorIn(v_local4_try08556) {va_arg2_try08544 = v_local4_try08556
            break
            } else {
            }}
            {
            v_local4 = ANY(v_local4_try08556)
            ToList(OBJ(va_arg2_try08544)).PutAt(CLcount,v_local4)
            } 
          }
          /* Iteration-4 */} 
        /* ERROR PROTECTION INSERTED (va_arg2-Result) */
        if ErrorIn(va_arg2_try08544) {Result = va_arg2_try08544
        } else {
        va_arg2 = ToList(OBJ(va_arg2_try08544))
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
      var g0860UU *ClaireAny  
      /* noccur = 1 */
      var g0860UU_try08613 EID 
      g0860UU_try08613 = Core.F_CALL(C_c_code,ARGS(self.Arg.ToEID(),EID{C_any.Id(),0}))
      /* ERROR PROTECTION INSERTED (g0860UU-Result) */
      if ErrorIn(g0860UU_try08613) {Result = g0860UU_try08613
      } else {
      g0860UU = ANY(g0860UU_try08613)
      Result = Language.C_Return.Make(g0860UU).ToEID()
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
      var g0862UU *ClaireType  
      /* noccur = 1 */
      var g0862UU_try08643 EID 
      g0862UU_try08643 = Core.F_CALL(C_c_type,ARGS(self.Arg.ToEID()))
      /* ERROR PROTECTION INSERTED (g0862UU-Result) */
      if ErrorIn(g0862UU_try08643) {Result = g0862UU_try08643
      } else {
      g0862UU = ToType(OBJ(g0862UU_try08643))
      /* Let:3 */{ 
        var g0863UU *ClaireType  
        /* noccur = 1 */
        var g0863UU_try08654 EID 
        g0863UU_try08654 = Core.F_CALL(C_c_type,ARGS(self.Other.ToEID()))
        /* ERROR PROTECTION INSERTED (g0863UU-Result) */
        if ErrorIn(g0863UU_try08654) {Result = g0863UU_try08654
        } else {
        g0863UU = ToType(OBJ(g0863UU_try08654))
        Result = EID{Core.F_U_type(g0862UU,g0863UU).Id(),0}
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
      var x_try08663 EID 
      /* Let:3 */{ 
        var g0867UU *ClaireAny  
        /* noccur = 1 */
        var g0867UU_try08694 EID 
        g0867UU_try08694 = Core.F_CALL(C_c_code,ARGS(self.Arg.ToEID(),EID{s.Id(),0}))
        /* ERROR PROTECTION INSERTED (g0867UU-x_try08663) */
        if ErrorIn(g0867UU_try08694) {x_try08663 = g0867UU_try08694
        } else {
        g0867UU = ANY(g0867UU_try08694)
        /* Let:4 */{ 
          var g0868UU *ClaireAny  
          /* noccur = 1 */
          var g0868UU_try08705 EID 
          g0868UU_try08705 = Core.F_CALL(C_c_code,ARGS(self.Other.ToEID(),EID{s.Id(),0}))
          /* ERROR PROTECTION INSERTED (g0868UU-x_try08663) */
          if ErrorIn(g0868UU_try08705) {x_try08663 = g0868UU_try08705
          } else {
          g0868UU = ANY(g0868UU_try08705)
          x_try08663 = Language.C_Handle.Make(C_any.Id(),g0867UU,g0868UU).ToEID()
          }
          /* Let-4 */} 
        }
        /* Let-3 */} 
      /* ERROR PROTECTION INSERTED (x-Result) */
      if ErrorIn(x_try08663) {Result = x_try08663
      } else {
      x = Language.To_ClaireHandle(OBJ(x_try08663))
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
        var g0872I *ClaireBoolean  
        if (y.Isa.IsIn(C_Param) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var g0871 *ClaireParam   = To_Param(y.Id())
            /* noccur = 3 */
            g0872I = MakeBoolean(((g0871.Arg.Id() == C_list.Id()) || 
                (g0871.Arg.Id() == C_set.Id())) && (C_set.Id() == g0871.Args.At(1-1).Isa.Id()))
            /* Let-5 */} 
          } else {
          g0872I = CFALSE
          /* If-4 */} 
        if (g0872I == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var utype *ClaireAny  
            /* noccur = 2 */
            var utype_try08736 EID 
            utype_try08736 = Core.F_the_type(ToType(To_Param(y.Id()).Args.At(1-1)))
            /* ERROR PROTECTION INSERTED (utype-Result) */
            if ErrorIn(utype_try08736) {Result = utype_try08736
            } else {
            utype = ANY(utype_try08736)
            var g0874I *ClaireBoolean  
            var g0874I_try08756 EID 
            /* or:6 */{ 
              var v_or6 *ClaireBoolean  
              
              var v_or6_try08767 EID 
              /* Let:7 */{ 
                var g0877UU *ClaireType  
                /* noccur = 1 */
                var g0877UU_try08788 EID 
                /* Let:8 */{ 
                  var g0879UU *ClaireType  
                  /* noccur = 1 */
                  var g0879UU_try08809 EID 
                  g0879UU_try08809 = Core.F_CALL(C_c_type,ARGS(self.Arg.ToEID()))
                  /* ERROR PROTECTION INSERTED (g0879UU-g0877UU_try08788) */
                  if ErrorIn(g0879UU_try08809) {g0877UU_try08788 = g0879UU_try08809
                  } else {
                  g0879UU = ToType(OBJ(g0879UU_try08809))
                  g0877UU_try08788 = EID{g0879UU.At(C_of).Id(),0}
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (g0877UU-v_or6_try08767) */
                if ErrorIn(g0877UU_try08788) {v_or6_try08767 = g0877UU_try08788
                } else {
                g0877UU = ToType(OBJ(g0877UU_try08788))
                v_or6_try08767 = EID{Equal(g0877UU.Id(),utype).Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (v_or6-g0874I_try08756) */
              if ErrorIn(v_or6_try08767) {g0874I_try08756 = v_or6_try08767
              } else {
              v_or6 = ToBoolean(OBJ(v_or6_try08767))
              if (v_or6 == CTRUE) {g0874I_try08756 = EID{CTRUE.Id(),0}
              } else /* or:7 */{ 
                v_or6 = Core.F__sup_integer(C_compiler.Safety,4)
                if (v_or6 == CTRUE) {g0874I_try08756 = EID{CTRUE.Id(),0}
                } else /* or:8 */{ 
                  g0874I_try08756 = EID{CFALSE.Id(),0}/* org-8 */} 
                /* org-7 */} 
              }
              /* or-6 */} 
            /* ERROR PROTECTION INSERTED (g0874I-Result) */
            if ErrorIn(g0874I_try08756) {Result = g0874I_try08756
            } else {
            g0874I = ToBoolean(OBJ(g0874I_try08756))
            if (g0874I == CTRUE) /* If:6 */{ 
              Result = Core.F_CALL(C_c_code,ARGS(self.Arg.ToEID(),EID{ftype.Id(),0}))
              } else {
              /* Let:7 */{ 
                var g0881UU *Language.Call  
                /* noccur = 1 */
                /* Let:8 */{ 
                  var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                  /* noccur = 5 */
                  _CL_obj.Selector = Core.C_check_in
                  _CL_obj.Args = MakeConstantList(self.Arg,To_Param(y.Id()).Arg.Id(),utype)
                  g0881UU = _CL_obj
                  /* Let-8 */} 
                Result = Core.F_CALL(C_c_code,ARGS(EID{g0881UU.Id(),0},EID{ftype.Id(),0}))
                /* Let-7 */} 
              /* If-6 */} 
            }
            }
            /* Let-5 */} 
          } else {
          var g0882I *ClaireBoolean  
          var g0882I_try08835 EID 
          /* Let:5 */{ 
            var g0884UU *ClaireType  
            /* noccur = 1 */
            var g0884UU_try08856 EID 
            g0884UU_try08856 = Core.F_CALL(C_c_type,ARGS(self.Arg.ToEID()))
            /* ERROR PROTECTION INSERTED (g0884UU-g0882I_try08835) */
            if ErrorIn(g0884UU_try08856) {g0882I_try08835 = g0884UU_try08856
            } else {
            g0884UU = ToType(OBJ(g0884UU_try08856))
            g0882I_try08835 = EID{g0884UU.Included(y).Id(),0}
            }
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (g0882I-Result) */
          if ErrorIn(g0882I_try08835) {Result = g0882I_try08835
          } else {
          g0882I = ToBoolean(OBJ(g0882I_try08835))
          if (g0882I == CTRUE) /* If:5 */{ 
            Result = Core.F_CALL(C_c_code,ARGS(self.Arg.ToEID(),EID{ftype.Id(),0}))
            /* If!5 */}  else if (C_compiler.Safety > 1) /* If:5 */{ 
            /* Let:6 */{ 
              var _CL_obj *CompileCCast   = To_CompileCCast(new(CompileCCast).Is(C_Compile_C_cast))
              /* noccur = 4 */
              /* update:7 */{ 
                var va_arg1 *CompileCCast  
                var va_arg2 *ClaireAny  
                va_arg1 = _CL_obj
                var va_arg2_try08868 EID 
                va_arg2_try08868 = Core.F_CALL(C_c_code,ARGS(self.Arg.ToEID(),EID{ftype.Id(),0}))
                /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                if ErrorIn(va_arg2_try08868) {Result = va_arg2_try08868
                } else {
                va_arg2 = ANY(va_arg2_try08868)
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
              var g0887UU *Language.Call  
              /* noccur = 1 */
              /* Let:7 */{ 
                var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                /* noccur = 5 */
                _CL_obj.Selector = Core.C_check_in
                _CL_obj.Args = MakeConstantList(self.Arg,y.Id())
                g0887UU = _CL_obj
                /* Let-7 */} 
              Result = Core.F_CALL(C_c_code,ARGS(EID{g0887UU.Id(),0},EID{ftype.Id(),0}))
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
      var _Ztype_try08913 EID 
      /* Iteration:3 */{ 
        var v_list3 *ClaireList  
        var x *ClaireAny  
        var v_local3 *ClaireAny  
        v_list3 = self.Args
        _Ztype_try08913 = EID{CreateList(ToType(CEMPTY.Id()),v_list3.Length()).Id(),0}
        for CLcount := 0; CLcount < v_list3.Length(); CLcount++{ 
          x = v_list3.At(CLcount)
          var v_local3_try08925 EID 
          v_local3_try08925 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
          /* ERROR PROTECTION INSERTED (v_local3-_Ztype_try08913) */
          if ErrorIn(v_local3_try08925) {_Ztype_try08913 = v_local3_try08925
          _Ztype_try08913 = v_local3_try08925
          break
          } else {
          v_local3 = ANY(v_local3_try08925)
          ToList(OBJ(_Ztype_try08913)).PutAt(CLcount,v_local3)
          } 
        }
        /* Iteration-3 */} 
      /* ERROR PROTECTION INSERTED (_Ztype-Result) */
      if ErrorIn(_Ztype_try08913) {Result = _Ztype_try08913
      } else {
      _Ztype = ToList(OBJ(_Ztype_try08913))
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
              var g0888 *ClaireSlot   = ToSlot(prop)
              /* noccur = 1 */
              Result = EID{g0888.Range.Id(),0}
              /* Let-6 */} 
            /* If!5 */}  else if (C_method.Id() == prop.Isa.Id()) /* If:5 */{ 
            /* Let:6 */{ 
              var g0889 *ClaireMethod   = ToMethod(prop)
              /* noccur = 1 */
              Result = F_Optimize_use_range_method(g0889,_Ztype)
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
          var _Ztype_try08965 EID 
          /* Iteration:5 */{ 
            var v_list5 *ClaireList  
            var x *ClaireAny  
            var v_local5 *ClaireAny  
            v_list5 = self.Args
            _Ztype_try08965 = EID{CreateList(ToType(CEMPTY.Id()),v_list5.Length()).Id(),0}
            for CLcount := 0; CLcount < v_list5.Length(); CLcount++{ 
              x = v_list5.At(CLcount)
              var v_local5_try08977 EID 
              v_local5_try08977 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
              /* ERROR PROTECTION INSERTED (v_local5-_Ztype_try08965) */
              if ErrorIn(v_local5_try08977) {_Ztype_try08965 = v_local5_try08977
              _Ztype_try08965 = v_local5_try08977
              break
              } else {
              v_local5 = ANY(v_local5_try08977)
              ToList(OBJ(_Ztype_try08965)).PutAt(CLcount,v_local5)
              } 
            }
            /* Iteration-5 */} 
          /* ERROR PROTECTION INSERTED (_Ztype-Result) */
          if ErrorIn(_Ztype_try08965) {Result = _Ztype_try08965
          } else {
          _Ztype = ToList(OBJ(_Ztype_try08965))
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
                var g0893 *ClaireSlot   = ToSlot(prop)
                /* noccur = 4 */
                /* Let:8 */{ 
                  var _CL_obj *Language.CallSlot   = Language.To_CallSlot(new(Language.CallSlot).Is(Language.C_Call_slot))
                  /* noccur = 7 */
                  _CL_obj.Selector = g0893
                  /* update:9 */{ 
                    var va_arg1 *Language.CallSlot  
                    var va_arg2 *ClaireAny  
                    va_arg1 = _CL_obj
                    var va_arg2_try089810 EID 
                    va_arg2_try089810 = Core.F_CALL(C_c_code,ARGS(l.At(1-1).ToEID(),EID{F_Compile_psort_any(Core.F_domain_I_restriction(ToRestriction(g0893.Id())).Id()).Id(),0}))
                    /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                    if ErrorIn(va_arg2_try089810) {Result = va_arg2_try089810
                    } else {
                    va_arg2 = ANY(va_arg2_try089810)
                    /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                    va_arg1.Arg = va_arg2
                    Result = va_arg2.ToEID()
                    }
                    /* update-9 */} 
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  _CL_obj.Test = MakeBoolean((g0893.Range.Contains(g0893.Default) != CTRUE) && (C_compiler.Safety < 5))
                  Result = EID{_CL_obj.Id(),0}
                  }
                  /* Let-8 */} 
                /* Let-7 */} 
              /* If!6 */}  else if (C_method.Id() == prop.Isa.Id()) /* If:6 */{ 
              /* Let:7 */{ 
                var g0894 *ClaireMethod   = ToMethod(prop)
                /* noccur = 1 */
                Result = F_Optimize_c_code_method_method1(g0894,l,_Ztype)
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
        var va_arg2_try08994 EID 
        /* Iteration:4 */{ 
          var v_list4 *ClaireList  
          var x *ClaireAny  
          var v_local4 *ClaireAny  
          v_list4 = self.Args
          va_arg2_try08994 = EID{CreateList(ToType(CEMPTY.Id()),v_list4.Length()).Id(),0}
          for CLcount := 0; CLcount < v_list4.Length(); CLcount++{ 
            x = v_list4.At(CLcount)
            var v_local4_try09006 EID 
            v_local4_try09006 = Core.F_CALL(C_c_code,ARGS(x.ToEID(),EID{C_any.Id(),0}))
            /* ERROR PROTECTION INSERTED (v_local4-va_arg2_try08994) */
            if ErrorIn(v_local4_try09006) {va_arg2_try08994 = v_local4_try09006
            va_arg2_try08994 = v_local4_try09006
            break
            } else {
            v_local4 = ANY(v_local4_try09006)
            ToList(OBJ(va_arg2_try08994)).PutAt(CLcount,v_local4)
            } 
          }
          /* Iteration-4 */} 
        /* ERROR PROTECTION INSERTED (va_arg2-Result) */
        if ErrorIn(va_arg2_try08994) {Result = va_arg2_try08994
        } else {
        va_arg2 = ToList(OBJ(va_arg2_try08994))
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
        var g0901UU *ClaireObject  
        /* noccur = 1 */
        /* Let:4 */{ 
          var g0902UU *Language.Call  
          /* noccur = 1 */
          /* Let:5 */{ 
            var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
            /* noccur = 5 */
            _CL_obj.Selector = Core.C_not
            _CL_obj.Args = MakeConstantList(self.Args.At(1-1))
            g0902UU = _CL_obj
            /* Let-5 */} 
          /* Let:5 */{ 
            var g0903UU *Language.Call  
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
              g0903UU = _CL_obj
              /* Let-6 */} 
            g0901UU = ToObject(Language.C_If.Make(g0902UU.Id(),g0903UU.Id(),CFALSE.Id()))
            /* Let-5 */} 
          /* Let-4 */} 
        Result = Core.F_CALL(C_c_code,ARGS(EID{g0901UU.Id(),0},EID{C_any.Id(),0}))
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
      var g0904I *ClaireBoolean  
      var g0904I_try09053 EID 
      /* and:3 */{ 
        var v_and3 *ClaireBoolean  
        
        v_and3 = Equal(MakeInteger(a.Length()).Id(),MakeInteger(1).Id())
        if (v_and3 == CFALSE) {g0904I_try09053 = EID{CFALSE.Id(),0}
        } else /* arg:4 */{ 
          var v_and3_try09065 EID 
          /* Let:5 */{ 
            var g0907UU *ClaireType  
            /* noccur = 1 */
            var g0907UU_try09086 EID 
            g0907UU_try09086 = Core.F_CALL(C_c_type,ARGS(a.At(1-1).ToEID()))
            /* ERROR PROTECTION INSERTED (g0907UU-v_and3_try09065) */
            if ErrorIn(g0907UU_try09086) {v_and3_try09065 = g0907UU_try09086
            } else {
            g0907UU = ToType(OBJ(g0907UU_try09086))
            v_and3_try09065 = EID{g0907UU.Included(ToType(C_integer.Id())).Id(),0}
            }
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (v_and3-g0904I_try09053) */
          if ErrorIn(v_and3_try09065) {g0904I_try09053 = v_and3_try09065
          } else {
          v_and3 = ToBoolean(OBJ(v_and3_try09065))
          if (v_and3 == CFALSE) {g0904I_try09053 = EID{CFALSE.Id(),0}
          } else /* arg:5 */{ 
            g0904I_try09053 = EID{CTRUE.Id(),0}/* arg-5 */} 
          /* arg-4 */} 
        }
        /* and-3 */} 
      /* ERROR PROTECTION INSERTED (g0904I-Result) */
      if ErrorIn(g0904I_try09053) {Result = g0904I_try09053
      } else {
      g0904I = ToBoolean(OBJ(g0904I_try09053))
      if (g0904I == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var g0909UU *Language.Call  
          /* noccur = 1 */
          /* Let:5 */{ 
            var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
            /* noccur = 5 */
            _CL_obj.Selector = Core.C_write
            _CL_obj.Args = MakeConstantList(C_verbose.Id(),ClEnv.Id(),a.At(1-1))
            g0909UU = _CL_obj
            /* Let-5 */} 
          Result = Core.F_CALL(C_c_code,ARGS(EID{g0909UU.Id(),0}))
          /* Let-4 */} 
        } else {
        var g0910I *ClaireBoolean  
        /* and:4 */{ 
          var v_and4 *ClaireBoolean  
          
          v_and4 = Core.F__sup_integer(a.Length(),1)
          if (v_and4 == CFALSE) {g0910I = CFALSE
          } else /* arg:5 */{ 
            v_and4 = Equal(C_string.Id(),a.At(2-1).Isa.Id())
            if (v_and4 == CFALSE) {g0910I = CFALSE
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
                    var g0911UU *ClaireAny  
                    /* noccur = 1 */
                    var g0911UU_try091210 EID 
                    g0911UU_try091210 = EVAL(a.At(1-1))
                    /* ERROR PROTECTION INSERTED (g0911UU-v_or7_try9) */
                    if ErrorIn(g0911UU_try091210) {v_or7_try9 = g0911UU_try091210
                    } else {
                    g0911UU = ANY(g0911UU_try091210)
                    v_or7_try9 = EID{Core.F__inf_equal_integer(ToInteger(g0911UU).Value,ClEnv.Verbose).Id(),0}
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
              if (v_and4 == CFALSE) {g0910I = CFALSE
              } else /* arg:7 */{ 
                g0910I = CTRUE/* arg-7 */} 
              /* arg-6 */} 
            /* arg-5 */} 
          /* and-4 */} 
        if (g0910I == CTRUE) /* If:4 */{ 
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
              var g0913UU *ClaireObject  
              /* noccur = 1 */
              if (C_integer.Id() != a.At(1-1).Isa.Id()) /* If:7 */{ 
                /* Let:8 */{ 
                  var g0914UU *Language.Call  
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
                    g0914UU = _CL_obj
                    /* Let-9 */} 
                  g0913UU = ToObject(Language.C_If.Make(g0914UU.Id(),_Zc.Id(),CFALSE.Id()))
                  /* Let-8 */} 
                } else {
                g0913UU = ToObject(_Zc.Id())
                /* If-7 */} 
              Result = Core.F_CALL(C_c_code,ARGS(EID{g0913UU.Id(),0},EID{C_any.Id(),0}))
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
      var g0915UU *ClaireObject  
      /* noccur = 1 */
      /* Let:3 */{ 
        var g0916UU *Language.Do  
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
                var g0918UU *Language.Do  
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
                  g0918UU = _CL_obj
                  /* Let-8 */} 
                v_bag_arg = Language.C_If.Make(self.Args.At(1-1),CTRUE.Id(),g0918UU.Id())
                /* Let-7 */} 
              va_arg2.AddFast(v_bag_arg)/* Construct-6 */} 
            /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
            va_arg1.Args = va_arg2
            /* update-5 */} 
          g0916UU = _CL_obj
          /* Let-4 */} 
        /* Let:4 */{ 
          var g0917UU *Language.Do  
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
            g0917UU = _CL_obj
            /* Let-5 */} 
          g0915UU = ToObject(Language.C_Handle.Make(g0916UU.Id(),Core.C_contradiction.Id(),g0917UU.Id()))
          /* Let-4 */} 
        /* Let-3 */} 
      Result = Core.F_CALL(C_c_code,ARGS(EID{g0915UU.Id(),0},EID{C_any.Id(),0}))
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_code @ Branch (throw: true) 
func E_c_code_Branch (self EID) EID { 
    return /*(sm for c_code @ Branch= EID)*/ F_c_code_Branch(Language.To_Branch(OBJ(self)) )} 
  
/* {1} OPT.The go function for: c_code(self:Macro,s:class) [] */
func F_c_code_Macro (self *Language.Macro ,s *ClaireClass ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var g0919UU *ClaireAny  
      /* noccur = 1 */
      var g0919UU_try09203 EID 
      g0919UU_try09203 = Core.F_CALL(Language.C_macroexpand,ARGS(EID{self.Id(),0}))
      /* ERROR PROTECTION INSERTED (g0919UU-Result) */
      if ErrorIn(g0919UU_try09203) {Result = g0919UU_try09203
      } else {
      g0919UU = ANY(g0919UU_try09203)
      Result = Core.F_CALL(C_c_code,ARGS(g0919UU.ToEID(),EID{s.Id(),0}))
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
      var g0921UU *ClaireAny  
      /* noccur = 1 */
      var g0921UU_try09223 EID 
      g0921UU_try09223 = Core.F_CALL(Language.C_macroexpand,ARGS(EID{self.Id(),0}))
      /* ERROR PROTECTION INSERTED (g0921UU-Result) */
      if ErrorIn(g0921UU_try09223) {Result = g0921UU_try09223
      } else {
      g0921UU = ANY(g0921UU_try09223)
      Result = Core.F_CALL(C_c_type,ARGS(g0921UU.ToEID()))
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
                        var g0923UU *Language.Call  
                        /* noccur = 1 */
                        /* Let:12 */{ 
                          var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                          /* noccur = 5 */
                          _CL_obj.Selector = C_princ
                          _CL_obj.Args = MakeConstantList((F_substring_string(s,1,(n-1))).Id())
                          g0923UU = _CL_obj
                          /* Let-12 */} 
                        r = r.AddFast(g0923UU.Id())
                        /* Let-11 */} 
                      /* If-10 */} 
                    var r_try092410 EID 
                    /* Let:10 */{ 
                      var g0925UU *ClaireAny  
                      /* noccur = 1 */
                      var g0925UU_try092611 EID 
                      if ('A' == m) /* If:11 */{ 
                        /* Let:12 */{ 
                          var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                          /* noccur = 5 */
                          _CL_obj.Selector = C_princ
                          _CL_obj.Args = MakeConstantList(l.At(i-1))
                          g0925UU_try092611 = EID{_CL_obj.Id(),0}
                          /* Let-12 */} 
                        /* If!11 */}  else if ('S' == m) /* If:11 */{ 
                        /* Let:12 */{ 
                          var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                          /* noccur = 5 */
                          _CL_obj.Selector = C_print
                          _CL_obj.Args = MakeConstantList(l.At(i-1))
                          g0925UU_try092611 = EID{_CL_obj.Id(),0}
                          /* Let-12 */} 
                        /* If!11 */}  else if ('F' == m) /* If:11 */{ 
                        /* Let:12 */{ 
                          var p_Z *ClaireBoolean   = CFALSE
                          /* noccur = 4 */
                          /* Let:13 */{ 
                            var j int 
                            /* noccur = 4 */
                            var j_try092714 EID 
                            /* Let:14 */{ 
                              var g0928UU int 
                              /* noccur = 1 */
                              var g0928UU_try092915 EID 
                              /* Let:15 */{ 
                                var g0930UU rune 
                                /* noccur = 1 */
                                var g0930UU_try093116 EID 
                                g0930UU_try093116 = Core.F_nth_get_string(s,(n+2),(n+2))
                                /* ERROR PROTECTION INSERTED (g0930UU-g0928UU_try092915) */
                                if ErrorIn(g0930UU_try093116) {g0928UU_try092915 = g0930UU_try093116
                                } else {
                                g0930UU = CHAR(g0930UU_try093116)
                                g0928UU_try092915 = EID{C__INT,IVAL(int(g0930UU))}
                                }
                                /* Let-15 */} 
                              /* ERROR PROTECTION INSERTED (g0928UU-j_try092714) */
                              if ErrorIn(g0928UU_try092915) {j_try092714 = g0928UU_try092915
                              } else {
                              g0928UU = INT(g0928UU_try092915)
                              j_try092714 = EID{C__INT,IVAL((g0928UU-48))}
                              }
                              /* Let-14 */} 
                            /* ERROR PROTECTION INSERTED (j-g0925UU_try092611) */
                            if ErrorIn(j_try092714) {g0925UU_try092611 = j_try092714
                            } else {
                            j = INT(j_try092714)
                            if ('%' == s.At((n+2))) /* If:14 */{ 
                              p_Z = CTRUE
                              j = 1
                              g0925UU_try092611 = EID{C__INT,IVAL(j)}
                              /* If!14 */}  else if ((j < 0) || 
                                (j > 9)) /* If:14 */{ 
                              g0925UU_try092611 = ToException(Core.C_general_error.Make(MakeString("[189] F requires a single digit integer in ~S").Id(),MakeConstantList(self.Id()).Id())).Close()
                              } else {
                              g0925UU_try092611 = EID{CFALSE.Id(),0}
                              /* If-14 */} 
                            /* ERROR PROTECTION INSERTED (g0925UU_try092611-g0925UU_try092611) */
                            if !ErrorIn(g0925UU_try092611) {
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
                              g0925UU_try092611 = EID{_CL_obj.Id(),0}
                              /* Let-14 */} 
                            }
                            }
                            /* Let-13 */} 
                          /* Let-12 */} 
                        /* If!11 */}  else if ('I' == m) /* If:11 */{ 
                        g0925UU_try092611 = l.At(i-1).ToEID()
                        } else {
                        g0925UU_try092611 = EID{CFALSE.Id(),0}
                        /* If-11 */} 
                      /* ERROR PROTECTION INSERTED (g0925UU-r_try092410) */
                      if ErrorIn(g0925UU_try092611) {r_try092410 = g0925UU_try092611
                      } else {
                      g0925UU = ANY(g0925UU_try092611)
                      r_try092410 = EID{r.AddFast(g0925UU).Id(),0}
                      }
                      /* Let-10 */} 
                    /* ERROR PROTECTION INSERTED (r-void_try9) */
                    if ErrorIn(r_try092410) {void_try9 = r_try092410
                    Result = r_try092410
                    break
                    } else {
                    r = ToList(OBJ(r_try092410))
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
                    var g0932UU *Language.Call  
                    /* noccur = 1 */
                    /* Let:10 */{ 
                      var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                      /* noccur = 5 */
                      _CL_obj.Selector = C_princ
                      _CL_obj.Args = MakeConstantList((s).Id())
                      g0932UU = _CL_obj
                      /* Let-10 */} 
                    r = r.AddFast(g0932UU.Id())
                    /* Let-9 */} 
                  /* If-8 */} 
                /* Let:8 */{ 
                  var g0933UU *Language.Do  
                  /* noccur = 1 */
                  /* Let:9 */{ 
                    var _CL_obj *Language.Do   = Language.To_Do(new(Language.Do).Is(Language.C_Do))
                    /* noccur = 3 */
                    _CL_obj.Args = r
                    g0933UU = _CL_obj
                    /* Let-9 */} 
                  Result = Core.F_CALL(C_c_code,ARGS(EID{g0933UU.Id(),0},EID{C_any.Id(),0}))
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
      var g0934UU *Language.Call  
      /* noccur = 1 */
      var g0934UU_try09353 EID 
      /* Let:3 */{ 
        var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
        /* noccur = 20 */
        _CL_obj.Selector = C_close
        /* update:4 */{ 
          var va_arg1 *Language.Call  
          var va_arg2 *ClaireList  
          va_arg1 = _CL_obj
          var va_arg2_try09365 EID 
          /* Construct:5 */{ 
            var v_bag_arg *ClaireAny  
            va_arg2_try09365= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
            var v_bag_arg_try09376 EID 
            /* Let:6 */{ 
              var _CL_obj *Language.Cast   = Language.To_Cast(new(Language.Cast).Is(Language.C_Cast))
              /* noccur = 14 */
              /* update:7 */{ 
                var va_arg1 *Language.Cast  
                var va_arg2 *ClaireAny  
                va_arg1 = _CL_obj
                var va_arg2_try09388 EID 
                /* Let:8 */{ 
                  var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                  /* noccur = 9 */
                  _CL_obj.Selector = C_Compile_anyObject_I
                  /* update:9 */{ 
                    var va_arg1 *Language.Call  
                    var va_arg2 *ClaireList  
                    va_arg1 = _CL_obj
                    var va_arg2_try093910 EID 
                    /* Construct:10 */{ 
                      var v_bag_arg *ClaireAny  
                      va_arg2_try093910= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                      ToList(OBJ(va_arg2_try093910)).AddFast(Core.C_general_error.Id())
                      var v_bag_arg_try094011 EID 
                      /* Let:11 */{ 
                        var g0941UU *ClaireAny  
                        /* noccur = 1 */
                        var g0941UU_try094212 EID 
                        g0941UU_try094212 = Core.F_car_list(self.Args)
                        /* ERROR PROTECTION INSERTED (g0941UU-v_bag_arg_try094011) */
                        if ErrorIn(g0941UU_try094212) {v_bag_arg_try094011 = g0941UU_try094212
                        } else {
                        g0941UU = ANY(g0941UU_try094212)
                        v_bag_arg_try094011 = Core.F_CALL(C_c_code,ARGS(g0941UU.ToEID(),EID{C_any.Id(),0}))
                        }
                        /* Let-11 */} 
                      /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try093910) */
                      if ErrorIn(v_bag_arg_try094011) {va_arg2_try093910 = v_bag_arg_try094011
                      } else {
                      v_bag_arg = ANY(v_bag_arg_try094011)
                      ToList(OBJ(va_arg2_try093910)).AddFast(v_bag_arg)
                      var v_bag_arg_try094311 EID 
                      /* Let:11 */{ 
                        var g0944UU *ClaireObject  
                        /* noccur = 1 */
                        var g0944UU_try094512 EID 
                        if (self.Args.Length() != 1) /* If:12 */{ 
                          /* Let:13 */{ 
                            var _CL_obj *Language.List   = Language.To_List(new(Language.List).Is(Language.C_List))
                            /* noccur = 3 */
                            /* update:14 */{ 
                              var va_arg1 *Language.Construct  
                              var va_arg2 *ClaireList  
                              va_arg1 = Language.To_Construct(_CL_obj.Id())
                              var va_arg2_try094615 EID 
                              va_arg2_try094615 = self.Args.Cdr()
                              /* ERROR PROTECTION INSERTED (va_arg2-g0944UU_try094512) */
                              if ErrorIn(va_arg2_try094615) {g0944UU_try094512 = va_arg2_try094615
                              } else {
                              va_arg2 = ToList(OBJ(va_arg2_try094615))
                              /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                              va_arg1.Args = va_arg2
                              g0944UU_try094512 = EID{va_arg2.Id(),0}
                              }
                              /* update-14 */} 
                            /* ERROR PROTECTION INSERTED (g0944UU_try094512-g0944UU_try094512) */
                            if !ErrorIn(g0944UU_try094512) {
                            g0944UU_try094512 = EID{_CL_obj.Id(),0}
                            }
                            /* Let-13 */} 
                          } else {
                          g0944UU_try094512 = EID{CNIL.Id(),0}
                          /* If-12 */} 
                        /* ERROR PROTECTION INSERTED (g0944UU-v_bag_arg_try094311) */
                        if ErrorIn(g0944UU_try094512) {v_bag_arg_try094311 = g0944UU_try094512
                        } else {
                        g0944UU = ToObject(OBJ(g0944UU_try094512))
                        v_bag_arg_try094311 = Core.F_CALL(C_c_code,ARGS(EID{g0944UU.Id(),0},EID{C_any.Id(),0}))
                        }
                        /* Let-11 */} 
                      /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try093910) */
                      if ErrorIn(v_bag_arg_try094311) {va_arg2_try093910 = v_bag_arg_try094311
                      } else {
                      v_bag_arg = ANY(v_bag_arg_try094311)
                      ToList(OBJ(va_arg2_try093910)).AddFast(v_bag_arg)}}
                      /* Construct-10 */} 
                    /* ERROR PROTECTION INSERTED (va_arg2-va_arg2_try09388) */
                    if ErrorIn(va_arg2_try093910) {va_arg2_try09388 = va_arg2_try093910
                    } else {
                    va_arg2 = ToList(OBJ(va_arg2_try093910))
                    /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                    va_arg1.Args = va_arg2
                    va_arg2_try09388 = EID{va_arg2.Id(),0}
                    }
                    /* update-9 */} 
                  /* ERROR PROTECTION INSERTED (va_arg2_try09388-va_arg2_try09388) */
                  if !ErrorIn(va_arg2_try09388) {
                  va_arg2_try09388 = EID{_CL_obj.Id(),0}
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (va_arg2-v_bag_arg_try09376) */
                if ErrorIn(va_arg2_try09388) {v_bag_arg_try09376 = va_arg2_try09388
                } else {
                va_arg2 = ANY(va_arg2_try09388)
                /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                va_arg1.Arg = va_arg2
                v_bag_arg_try09376 = va_arg2.ToEID()
                }
                /* update-7 */} 
              /* ERROR PROTECTION INSERTED (v_bag_arg_try09376-v_bag_arg_try09376) */
              if !ErrorIn(v_bag_arg_try09376) {
              _CL_obj.SetArg = ToType(C_exception.Id())
              v_bag_arg_try09376 = EID{_CL_obj.Id(),0}
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try09365) */
            if ErrorIn(v_bag_arg_try09376) {va_arg2_try09365 = v_bag_arg_try09376
            } else {
            v_bag_arg = ANY(v_bag_arg_try09376)
            ToList(OBJ(va_arg2_try09365)).AddFast(v_bag_arg)}
            /* Construct-5 */} 
          /* ERROR PROTECTION INSERTED (va_arg2-g0934UU_try09353) */
          if ErrorIn(va_arg2_try09365) {g0934UU_try09353 = va_arg2_try09365
          } else {
          va_arg2 = ToList(OBJ(va_arg2_try09365))
          /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
          va_arg1.Args = va_arg2
          g0934UU_try09353 = EID{va_arg2.Id(),0}
          }
          /* update-4 */} 
        /* ERROR PROTECTION INSERTED (g0934UU_try09353-g0934UU_try09353) */
        if !ErrorIn(g0934UU_try09353) {
        g0934UU_try09353 = EID{_CL_obj.Id(),0}
        }
        /* Let-3 */} 
      /* ERROR PROTECTION INSERTED (g0934UU-Result) */
      if ErrorIn(g0934UU_try09353) {Result = g0934UU_try09353
      } else {
      g0934UU = Language.To_Call(OBJ(g0934UU_try09353))
      Result = Core.F_CALL(C_c_code,ARGS(EID{g0934UU.Id(),0},EID{C_void.Id(),0}))
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
          var g0947 *Language.Call   = Language.To_Call(_Zt)
          /* noccur = 3 */
          if ((g0947.Args.At(1-1).Isa.IsIn(C_Variable) == CTRUE) && 
              (g0947.Selector.Id() == Core.C_known_ask.Id())) /* If:5 */{ 
            Result = ToType(OBJ(Core.F_CALL(C_range,ARGS(g0947.Args.At(1-1).ToEID()))))
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
      var g0951I *ClaireBoolean  
      var g0951I_try09523 EID 
      g0951I_try09523 = F_Optimize_extended_ask_type(_Zr)
      /* ERROR PROTECTION INSERTED (g0951I-Result) */
      if ErrorIn(g0951I_try09523) {Result = g0951I_try09523
      } else {
      g0951I = ToBoolean(OBJ(g0951I_try09523))
      if (g0951I == CTRUE) /* If:3 */{ 
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
        var result_try09544 EID 
        /* Let:4 */{ 
          var g0955UU *ClaireType  
          /* noccur = 1 */
          var g0955UU_try09575 EID 
          g0955UU_try09575 = Core.F_CALL(C_c_type,ARGS(self.Arg.ToEID()))
          /* ERROR PROTECTION INSERTED (g0955UU-result_try09544) */
          if ErrorIn(g0955UU_try09575) {result_try09544 = g0955UU_try09575
          } else {
          g0955UU = ToType(OBJ(g0955UU_try09575))
          /* Let:5 */{ 
            var g0956UU *ClaireType  
            /* noccur = 1 */
            var g0956UU_try09586 EID 
            g0956UU_try09586 = Core.F_CALL(C_c_type,ARGS(self.Other.ToEID()))
            /* ERROR PROTECTION INSERTED (g0956UU-result_try09544) */
            if ErrorIn(g0956UU_try09586) {result_try09544 = g0956UU_try09586
            } else {
            g0956UU = ToType(OBJ(g0956UU_try09586))
            result_try09544 = EID{Core.F_U_type(g0955UU,g0956UU).Id(),0}
            }
            /* Let-5 */} 
          }
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (result-Result) */
        if ErrorIn(result_try09544) {Result = result_try09544
        } else {
        result = ToType(OBJ(result_try09544))
        var g0959I *ClaireBoolean  
        var g0959I_try09604 EID 
        g0959I_try09604 = F_Optimize_extended_ask_type(_Zr)
        /* ERROR PROTECTION INSERTED (g0959I-Result) */
        if ErrorIn(g0959I_try09604) {Result = g0959I_try09604
        } else {
        g0959I = ToBoolean(OBJ(g0959I_try09604))
        if (g0959I == CTRUE) /* If:4 */{ 
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
      var g0961I *ClaireBoolean  
      var g0961I_try09623 EID 
      g0961I_try09623 = F_Optimize_extended_ask_type(_Zr)
      /* ERROR PROTECTION INSERTED (g0961I-Result) */
      if ErrorIn(g0961I_try09623) {Result = g0961I_try09623
      } else {
      g0961I = ToBoolean(OBJ(g0961I_try09623))
      if (g0961I == CTRUE) /* If:3 */{ 
        F_Optimize_range_sets_any(ToList(OBJ(Core.F_CALL(C_args,ARGS(self.Test.ToEID())))).At(1-1),F_Optimize_sort_abstract_I_type(ToType(OBJ(Core.F_CALL(C_mClaire_t1,ARGS(EID{_Zr.Id(),0}))))))
        Result = EVOID
        } else {
        Result = EID{CFALSE.Id(),0}
        /* If-3 */} 
      }
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      var g0963I *ClaireBoolean  
      var g0963I_try09643 EID 
      /* and:3 */{ 
        var v_and3 *ClaireBoolean  
        
        var v_and3_try09654 EID 
        /* Let:4 */{ 
          var g0966UU *ClaireBoolean  
          /* noccur = 1 */
          var g0966UU_try09675 EID 
          /* Let:5 */{ 
            var g0968UU *ClaireType  
            /* noccur = 1 */
            var g0968UU_try09696 EID 
            /* Let:6 */{ 
              var g0970UU *ClaireType  
              /* noccur = 1 */
              var g0970UU_try09717 EID 
              g0970UU_try09717 = Core.F_CALL(C_c_type,ARGS(self.Test.ToEID()))
              /* ERROR PROTECTION INSERTED (g0970UU-g0968UU_try09696) */
              if ErrorIn(g0970UU_try09717) {g0968UU_try09696 = g0970UU_try09717
              } else {
              g0970UU = ToType(OBJ(g0970UU_try09717))
              g0968UU_try09696 = EID{F_Optimize_ptype_type(g0970UU).Id(),0}
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (g0968UU-g0966UU_try09675) */
            if ErrorIn(g0968UU_try09696) {g0966UU_try09675 = g0968UU_try09696
            } else {
            g0968UU = ToType(OBJ(g0968UU_try09696))
            g0966UU_try09675 = EID{g0968UU.Included(ToType(C_boolean.Id())).Id(),0}
            }
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (g0966UU-v_and3_try09654) */
          if ErrorIn(g0966UU_try09675) {v_and3_try09654 = g0966UU_try09675
          } else {
          g0966UU = ToBoolean(OBJ(g0966UU_try09675))
          v_and3_try09654 = EID{g0966UU.Not.Id(),0}
          }
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (v_and3-g0963I_try09643) */
        if ErrorIn(v_and3_try09654) {g0963I_try09643 = v_and3_try09654
        } else {
        v_and3 = ToBoolean(OBJ(v_and3_try09654))
        if (v_and3 == CFALSE) {g0963I_try09643 = EID{CFALSE.Id(),0}
        } else /* arg:4 */{ 
          v_and3 = Equal(C_PENIBLE.Value,CTRUE.Id())
          if (v_and3 == CFALSE) {g0963I_try09643 = EID{CFALSE.Id(),0}
          } else /* arg:5 */{ 
            g0963I_try09643 = EID{CTRUE.Id(),0}/* arg-5 */} 
          /* arg-4 */} 
        }
        /* and-3 */} 
      /* ERROR PROTECTION INSERTED (g0963I-Result) */
      if ErrorIn(g0963I_try09643) {Result = g0963I_try09643
      } else {
      g0963I = ToBoolean(OBJ(g0963I_try09643))
      if (g0963I == CTRUE) /* If:3 */{ 
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
        var result_try09724 EID 
        /* Let:4 */{ 
          var g0973UU *ClaireAny  
          /* noccur = 1 */
          var g0973UU_try09765 EID 
          g0973UU_try09765 = F_Optimize_c_boolean_any(self.Test)
          /* ERROR PROTECTION INSERTED (g0973UU-result_try09724) */
          if ErrorIn(g0973UU_try09765) {result_try09724 = g0973UU_try09765
          } else {
          g0973UU = ANY(g0973UU_try09765)
          /* Let:5 */{ 
            var g0974UU *ClaireAny  
            /* noccur = 1 */
            var g0974UU_try09776 EID 
            g0974UU_try09776 = Core.F_CALL(C_c_code,ARGS(self.Arg.ToEID(),EID{s.Id(),0}))
            /* ERROR PROTECTION INSERTED (g0974UU-result_try09724) */
            if ErrorIn(g0974UU_try09776) {result_try09724 = g0974UU_try09776
            } else {
            g0974UU = ANY(g0974UU_try09776)
            /* Let:6 */{ 
              var g0975UU *ClaireAny  
              /* noccur = 1 */
              var g0975UU_try09787 EID 
              g0975UU_try09787 = Core.F_CALL(C_c_code,ARGS(self.Other.ToEID(),EID{s.Id(),0}))
              /* ERROR PROTECTION INSERTED (g0975UU-result_try09724) */
              if ErrorIn(g0975UU_try09787) {result_try09724 = g0975UU_try09787
              } else {
              g0975UU = ANY(g0975UU_try09787)
              result_try09724 = Language.C_If.Make(g0973UU,g0974UU,g0975UU).ToEID()
              }
              /* Let-6 */} 
            }
            /* Let-5 */} 
          }
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (result-Result) */
        if ErrorIn(result_try09724) {Result = result_try09724
        } else {
        result = Language.To_If(OBJ(result_try09724))
        var g0979I *ClaireBoolean  
        var g0979I_try09804 EID 
        g0979I_try09804 = F_Optimize_extended_ask_type(_Zr)
        /* ERROR PROTECTION INSERTED (g0979I-Result) */
        if ErrorIn(g0979I_try09804) {Result = g0979I_try09804
        } else {
        g0979I = ToBoolean(OBJ(g0979I_try09804))
        if (g0979I == CTRUE) /* If:4 */{ 
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
            var g0981 *ClaireVariable   = To_Variable(_Zvar)
            /* noccur = 1 */
            _Ztype = Core.F_get_property(C_range,ToObject(g0981.Id()))
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
                    var g0984UU *ClaireAny  
                    /* noccur = 1 */
                    var g0984UU_try098510 EID 
                    g0984UU_try098510 = Core.F_car_list(l)
                    /* ERROR PROTECTION INSERTED (g0984UU-void_try8) */
                    if ErrorIn(g0984UU_try098510) {void_try8 = g0984UU_try098510
                    } else {
                    g0984UU = ANY(g0984UU_try098510)
                    void_try8 = Core.F_CALL(C_Compile_Cerror,ARGS(EID{MakeString("[208] wrong type declaration for case: ~S in ~S").Id(),0},g0984UU.ToEID(),EID{self.Id(),0}))
                    }
                    /* Let-9 */} 
                  /* If-8 */} 
                /* ERROR PROTECTION INSERTED (void_try8-void_try8) */
                if ErrorIn(void_try8) {Result = void_try8
                break
                } else {
                var rtype_try09868 EID 
                /* Let:8 */{ 
                  var g0987UU *ClaireType  
                  /* noccur = 1 */
                  var g0987UU_try09889 EID 
                  g0987UU_try09889 = Core.F_CALL(C_c_type,ARGS(l.At(2-1).ToEID()))
                  /* ERROR PROTECTION INSERTED (g0987UU-rtype_try09868) */
                  if ErrorIn(g0987UU_try09889) {rtype_try09868 = g0987UU_try09889
                  } else {
                  g0987UU = ToType(OBJ(g0987UU_try09889))
                  rtype_try09868 = EID{Core.F_U_type(rtype,g0987UU).Id(),0}
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (rtype-void_try8) */
                if ErrorIn(rtype_try09868) {void_try8 = rtype_try09868
                Result = rtype_try09868
                break
                } else {
                rtype = ToType(OBJ(rtype_try09868))
                void_try8 = EID{rtype.Id(),0}
                
                if (_Zvar.Isa.IsIn(C_Variable) == CTRUE) /* If:8 */{ 
                  /* Let:9 */{ 
                    var g0983 *ClaireVariable   = To_Variable(_Zvar)
                    /* noccur = 1 */
                    g0983.Range = ToType(_Ztype)
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
        var g0989 *ClaireVariable   = To_Variable(_Zvar)
        /* noccur = 4 */
        /* Let:4 */{ 
          var vsub *ClaireVariable   = F_Compile_Variable_I_symbol(Core.F_gensym_void(),0,_Ztype.Id())
          /* noccur = 2 */
          if ((Equal(_Ztype.Id(),g0989.Range.Id()) != CTRUE) && 
              ((_Ztype.Id() != C_any.Id()) && 
                (Language.F_occurrence_any(x,g0989) > 0))) /* If:5 */{ 
            /* Let:6 */{ 
              var _CL_obj *Language.Let   = Language.To_Let(new(Language.Let).Is(Language.C_Let))
              /* noccur = 6 */
              _CL_obj.ClaireVar = vsub
              _CL_obj.Value = g0989.Id()
              _CL_obj.Arg = F_Optimize_case_substitution_any(x,g0989,vsub)
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
            var g0993 *ClaireVariable   = To_Variable(_Zvar)
            /* noccur = 1 */
            _Ztype = Core.F_get_property(C_range,ToObject(g0993.Id()))
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
              var ctest1_try09967 EID 
              /* Let:7 */{ 
                var g0997UU *Language.Call  
                /* noccur = 1 */
                /* Let:8 */{ 
                  var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                  /* noccur = 5 */
                  _CL_obj.Selector = ToProperty(C__Z.Id())
                  _CL_obj.Args = MakeConstantList(_Zvar,l.At(1-1))
                  g0997UU = _CL_obj
                  /* Let-8 */} 
                ctest1_try09967 = F_Optimize_c_boolean_any(g0997UU.Id())
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (ctest1-Result) */
              if ErrorIn(ctest1_try09967) {Result = ctest1_try09967
              } else {
              ctest1 = ANY(ctest1_try09967)
              /* Let:7 */{ 
                var rep *Language.If  
                /* noccur = 2 */
                var rep_try10008 EID 
                /* Let:8 */{ 
                  var g1001UU *ClaireAny  
                  /* noccur = 1 */
                  var g1001UU_try10039 EID 
                  g1001UU_try10039 = Core.F_CALL(C_c_code,ARGS(F_Optimize_case_branch_any(l.At(2-1),_Zvar,ToType(l.At(1-1))).ToEID(),EID{s.Id(),0}))
                  /* ERROR PROTECTION INSERTED (g1001UU-rep_try10008) */
                  if ErrorIn(g1001UU_try10039) {rep_try10008 = g1001UU_try10039
                  } else {
                  g1001UU = ANY(g1001UU_try10039)
                  /* Let:9 */{ 
                    var g1002UU *ClaireAny  
                    /* noccur = 1 */
                    var g1002UU_try100410 EID 
                    g1002UU_try100410 = Core.F_CALL(C_c_code,ARGS(EID{CFALSE.Id(),0},EID{s.Id(),0}))
                    /* ERROR PROTECTION INSERTED (g1002UU-rep_try10008) */
                    if ErrorIn(g1002UU_try100410) {rep_try10008 = g1002UU_try100410
                    } else {
                    g1002UU = ANY(g1002UU_try100410)
                    rep_try10008 = Language.C_If.Make(ctest1,g1001UU,g1002UU).ToEID()
                    }
                    /* Let-9 */} 
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (rep-Result) */
                if ErrorIn(rep_try10008) {Result = rep_try10008
                } else {
                rep = Language.To_If(OBJ(rep_try10008))
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
                        var va_arg2_try100512 EID 
                        va_arg2_try100512 = Core.F_CALL(C_c_code,ARGS(F_Optimize_case_branch_any(l.At(2-1),_Zvar,ToType(l.At(1-1))).ToEID(),EID{s.Id(),0}))
                        /* ERROR PROTECTION INSERTED (va_arg2-void_try10) */
                        if ErrorIn(va_arg2_try100512) {void_try10 = va_arg2_try100512
                        } else {
                        va_arg2 = ANY(va_arg2_try100512)
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
                        var ctest_try100612 EID 
                        /* Let:12 */{ 
                          var g1007UU *Language.Call  
                          /* noccur = 1 */
                          /* Let:13 */{ 
                            var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                            /* noccur = 5 */
                            _CL_obj.Selector = ToProperty(C__Z.Id())
                            _CL_obj.Args = MakeConstantList(_Zvar,l.At(1-1))
                            g1007UU = _CL_obj
                            /* Let-13 */} 
                          ctest_try100612 = F_Optimize_c_boolean_any(g1007UU.Id())
                          /* Let-12 */} 
                        /* ERROR PROTECTION INSERTED (ctest-void_try10) */
                        if ErrorIn(ctest_try100612) {void_try10 = ctest_try100612
                        } else {
                        ctest = ANY(ctest_try100612)
                        /* update:12 */{ 
                          var va_arg1 *Language.If  
                          var va_arg2 *ClaireAny  
                          va_arg1 = pointer
                          var va_arg2_try100813 EID 
                          /* Let:13 */{ 
                            var g1009UU *ClaireAny  
                            /* noccur = 1 */
                            var g1009UU_try101114 EID 
                            g1009UU_try101114 = Core.F_CALL(C_c_code,ARGS(F_Optimize_case_branch_any(l.At(2-1),_Zvar,ToType(l.At(1-1))).ToEID(),EID{s.Id(),0}))
                            /* ERROR PROTECTION INSERTED (g1009UU-va_arg2_try100813) */
                            if ErrorIn(g1009UU_try101114) {va_arg2_try100813 = g1009UU_try101114
                            } else {
                            g1009UU = ANY(g1009UU_try101114)
                            /* Let:14 */{ 
                              var g1010UU *ClaireAny  
                              /* noccur = 1 */
                              var g1010UU_try101215 EID 
                              g1010UU_try101215 = Core.F_CALL(C_c_code,ARGS(EID{CFALSE.Id(),0},EID{s.Id(),0}))
                              /* ERROR PROTECTION INSERTED (g1010UU-va_arg2_try100813) */
                              if ErrorIn(g1010UU_try101215) {va_arg2_try100813 = g1010UU_try101215
                              } else {
                              g1010UU = ANY(g1010UU_try101215)
                              va_arg2_try100813 = Language.C_If.Make(ctest,g1009UU,g1010UU).ToEID()
                              }
                              /* Let-14 */} 
                            }
                            /* Let-13 */} 
                          /* ERROR PROTECTION INSERTED (va_arg2-void_try10) */
                          if ErrorIn(va_arg2_try100813) {void_try10 = va_arg2_try100813
                          } else {
                          va_arg2 = ANY(va_arg2_try100813)
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
                  var g1013I *ClaireBoolean  
                  if (_Zvar.Isa.IsIn(Language.C_Definition) == CTRUE) /* If:9 */{ 
                    /* Let:10 */{ 
                      var g0995 *Language.Definition   = Language.To_Definition(_Zvar)
                      /* noccur = 1 */
                      g1013I = g0995.Arg.Isa.IsIn(C_exception)
                      /* Let-10 */} 
                    } else {
                    g1013I = CFALSE
                    /* If-9 */} 
                  if (g1013I == CTRUE) /* If:9 */{ 
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
      var g1014UU *ClaireAny  
      /* noccur = 1 */
      var g1014UU_try10153 EID 
      g1014UU_try10153 = Core.F_last_list(self.Args)
      /* ERROR PROTECTION INSERTED (g1014UU-Result) */
      if ErrorIn(g1014UU_try10153) {Result = g1014UU_try10153
      } else {
      g1014UU = ANY(g1014UU_try10153)
      Result = Core.F_CALL(C_c_type,ARGS(g1014UU.ToEID()))
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
        var va_arg2_try10164 EID 
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
              va_arg2_try10164 = EID{CreateList(ToType(CEMPTY.Id()),v_list6.Length()).Id(),0}
              for CLcount := 0; CLcount < v_list6.Length(); CLcount++{ 
                x = v_list6.At(CLcount)
                var v_local6_try10178 EID 
                n = (n+1)
                /* Let:8 */{ 
                  var g1018UU *ClaireClass  
                  /* noccur = 1 */
                  if (n == m) /* If:9 */{ 
                    g1018UU = s
                    } else {
                    g1018UU = C_void
                    /* If-9 */} 
                  v_local6_try10178 = Core.F_CALL(C_c_code,ARGS(x.ToEID(),EID{g1018UU.Id(),0}))
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (v_local6_try10178-v_local6_try10178) */
                if ErrorIn(v_local6_try10178) {va_arg2_try10164 = v_local6_try10178
                break
                } else {
                }
                {
                v_local6 = ANY(v_local6_try10178)
                ToList(OBJ(va_arg2_try10164)).PutAt(CLcount,v_local6)
                } 
              }
              /* Iteration-6 */} 
            /* Let-5 */} 
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (va_arg2-Result) */
        if ErrorIn(va_arg2_try10164) {Result = va_arg2_try10164
        } else {
        va_arg2 = ToList(OBJ(va_arg2_try10164))
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
      var g1019UU *ClaireType  
      /* noccur = 1 */
      var g1019UU_try10203 EID 
      g1019UU_try10203 = Core.F_CALL(C_c_type,ARGS(self.Value.ToEID()))
      /* ERROR PROTECTION INSERTED (g1019UU-Result) */
      if ErrorIn(g1019UU_try10203) {Result = g1019UU_try10203
      } else {
      g1019UU = ToType(OBJ(g1019UU_try10203))
      Result = F_Optimize_range_infers_Variable(self.ClaireVar,g1019UU)
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
        var _Ztype_try10214 EID 
        /* Let:4 */{ 
          var g1022UU *ClaireType  
          /* noccur = 1 */
          var g1022UU_try10235 EID 
          g1022UU_try10235 = Core.F_CALL(C_c_type,ARGS(_Zv.ToEID()))
          /* ERROR PROTECTION INSERTED (g1022UU-_Ztype_try10214) */
          if ErrorIn(g1022UU_try10235) {_Ztype_try10214 = g1022UU_try10235
          } else {
          g1022UU = ToType(OBJ(g1022UU_try10235))
          _Ztype_try10214 = EID{F_Optimize_ptype_type(g1022UU).Id(),0}
          }
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (_Ztype-Result) */
        if ErrorIn(_Ztype_try10214) {Result = _Ztype_try10214
        } else {
        _Ztype = ToType(OBJ(_Ztype_try10214))
        Result = F_Optimize_range_infers_Variable(self.ClaireVar,_Ztype)
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        if (_Ztype.Included(self.ClaireVar.Range) != CTRUE) /* If:4 */{ 
          var _Zv_try10245 EID 
          _Zv_try10245 = F_Optimize_c_warn_Variable(self.ClaireVar,_Zv,_Ztype)
          /* ERROR PROTECTION INSERTED (_Zv-Result) */
          if ErrorIn(_Zv_try10245) {Result = _Zv_try10245
          } else {
          _Zv = ANY(_Zv_try10245)
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
          var x_try10255 EID 
          /* Let:5 */{ 
            var _CL_obj *Language.Let   = Language.To_Let(new(Language.Let).Is(Language.C_Let))
            /* noccur = 6 */
            _CL_obj.ClaireVar = self.ClaireVar
            /* update:6 */{ 
              var va_arg1 *Language.Let  
              var va_arg2 *ClaireAny  
              va_arg1 = _CL_obj
              var va_arg2_try10267 EID 
              va_arg2_try10267 = F_Compile_c_strict_code_any(_Zv,F_Compile_psort_any(self.ClaireVar.Range.Id()))
              /* ERROR PROTECTION INSERTED (va_arg2-x_try10255) */
              if ErrorIn(va_arg2_try10267) {x_try10255 = va_arg2_try10267
              } else {
              va_arg2 = ANY(va_arg2_try10267)
              /* ---------- now we compile update value(va_arg1) := va_arg2 ------- */
              va_arg1.Value = va_arg2
              x_try10255 = va_arg2.ToEID()
              }
              /* update-6 */} 
            /* ERROR PROTECTION INSERTED (x_try10255-x_try10255) */
            if !ErrorIn(x_try10255) {
            /* update:6 */{ 
              var va_arg1 *Language.Let  
              var va_arg2 *ClaireAny  
              va_arg1 = _CL_obj
              var va_arg2_try10277 EID 
              va_arg2_try10277 = Core.F_CALL(C_c_code,ARGS(self.Arg.ToEID(),EID{s.Id(),0}))
              /* ERROR PROTECTION INSERTED (va_arg2-x_try10255) */
              if ErrorIn(va_arg2_try10277) {x_try10255 = va_arg2_try10277
              } else {
              va_arg2 = ANY(va_arg2_try10277)
              /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
              va_arg1.Arg = va_arg2
              x_try10255 = va_arg2.ToEID()
              }
              /* update-6 */} 
            /* ERROR PROTECTION INSERTED (x_try10255-x_try10255) */
            if !ErrorIn(x_try10255) {
            x_try10255 = EID{_CL_obj.Id(),0}
            }}
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (x-Result) */
          if ErrorIn(x_try10255) {Result = x_try10255
          } else {
          x = Language.To_Let(OBJ(x_try10255))
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
          var d_try10285 EID 
          d_try10285 = F_Optimize_daccess_any(_Zv,CTRUE)
          /* ERROR PROTECTION INSERTED (d-Result) */
          if ErrorIn(d_try10285) {Result = d_try10285
          } else {
          d = ANY(d_try10285)
          /* Let:5 */{ 
            var _Ztype *ClaireAny  
            /* noccur = 4 */
            var _Ztype_try10296 EID 
            if (d != CNULL) /* If:6 */{ 
              _Ztype_try10296 = Core.F_CALL(C_c_type,ARGS(d.ToEID()))
              } else {
              _Ztype_try10296 = Core.F_CALL(C_c_type,ARGS(_Zv.ToEID()))
              /* If-6 */} 
            /* ERROR PROTECTION INSERTED (_Ztype-Result) */
            if ErrorIn(_Ztype_try10296) {Result = _Ztype_try10296
            } else {
            _Ztype = ANY(_Ztype_try10296)
            var g1030I *ClaireBoolean  
            var g1030I_try10316 EID 
            g1030I_try10316 = F_Optimize_extended_ask_type(ToType(_Ztype))
            /* ERROR PROTECTION INSERTED (g1030I-Result) */
            if ErrorIn(g1030I_try10316) {Result = g1030I_try10316
            } else {
            g1030I = ToBoolean(OBJ(g1030I_try10316))
            if (g1030I == CTRUE) /* If:6 */{ 
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
              var g1032UU *ClaireType  
              /* noccur = 1 */
              var g1032UU_try10347 EID 
              g1032UU_try10347 = Core.F_CALL(C_c_type,ARGS(self.Arg.ToEID()))
              /* ERROR PROTECTION INSERTED (g1032UU-Result) */
              if ErrorIn(g1032UU_try10347) {Result = g1032UU_try10347
              } else {
              g1032UU = ToType(OBJ(g1032UU_try10347))
              /* Let:7 */{ 
                var g1033UU *ClaireType  
                /* noccur = 1 */
                var g1033UU_try10358 EID 
                g1033UU_try10358 = Core.F_CALL(C_c_type,ARGS(self.Other.ToEID()))
                /* ERROR PROTECTION INSERTED (g1033UU-Result) */
                if ErrorIn(g1033UU_try10358) {Result = g1033UU_try10358
                } else {
                g1033UU = ToType(OBJ(g1033UU_try10358))
                Result = EID{Core.F_U_type(g1032UU,g1033UU).Id(),0}
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
          var d_try10365 EID 
          d_try10365 = F_Optimize_daccess_any(_Zv,CTRUE)
          /* ERROR PROTECTION INSERTED (d-Result) */
          if ErrorIn(d_try10365) {Result = d_try10365
          } else {
          d = ANY(d_try10365)
          /* Let:5 */{ 
            var v2 *ClaireVariable   = F_Compile_Variable_I_symbol(F_append_symbol(self.ClaireVar.Pname,MakeString("test").Id()),self.ClaireVar.Index,C_any.Id())
            /* noccur = 3 */
            /* Let:6 */{ 
              var _Ztype *ClaireAny  
              /* noccur = 6 */
              var _Ztype_try10377 EID 
              if (d != CNULL) /* If:7 */{ 
                _Ztype_try10377 = Core.F_CALL(C_c_type,ARGS(d.ToEID()))
                } else {
                _Ztype_try10377 = Core.F_CALL(C_c_type,ARGS(_Zv.ToEID()))
                /* If-7 */} 
              /* ERROR PROTECTION INSERTED (_Ztype-Result) */
              if ErrorIn(_Ztype_try10377) {Result = _Ztype_try10377
              } else {
              _Ztype = ANY(_Ztype_try10377)
              var g1038I *ClaireBoolean  
              var g1038I_try10397 EID 
              g1038I_try10397 = F_Optimize_extended_ask_type(ToType(_Ztype))
              /* ERROR PROTECTION INSERTED (g1038I-Result) */
              if ErrorIn(g1038I_try10397) {Result = g1038I_try10397
              } else {
              g1038I = ToBoolean(OBJ(g1038I_try10397))
              if (g1038I == CTRUE) /* If:7 */{ 
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
              var g1040I *ClaireBoolean  
              var g1040I_try10417 EID 
              /* and:7 */{ 
                var v_and7 *ClaireBoolean  
                
                v_and7 = Core.F_known_ask_any(d)
                if (v_and7 == CFALSE) {g1040I_try10417 = EID{CFALSE.Id(),0}
                } else /* arg:8 */{ 
                  var v_and7_try10429 EID 
                  /* Let:9 */{ 
                    var g1043UU *ClaireBoolean  
                    /* noccur = 1 */
                    var g1043UU_try104410 EID 
                    g1043UU_try104410 = F_Optimize_extended_ask_type(ToType(OBJ(Core.F_CALL(C_range,ARGS(Core.F_CALL(C_selector,ARGS(d.ToEID())))))))
                    /* ERROR PROTECTION INSERTED (g1043UU-v_and7_try10429) */
                    if ErrorIn(g1043UU_try104410) {v_and7_try10429 = g1043UU_try104410
                    } else {
                    g1043UU = ToBoolean(OBJ(g1043UU_try104410))
                    v_and7_try10429 = EID{g1043UU.Not.Id(),0}
                    }
                    /* Let-9 */} 
                  /* ERROR PROTECTION INSERTED (v_and7-g1040I_try10417) */
                  if ErrorIn(v_and7_try10429) {g1040I_try10417 = v_and7_try10429
                  } else {
                  v_and7 = ToBoolean(OBJ(v_and7_try10429))
                  if (v_and7 == CFALSE) {g1040I_try10417 = EID{CFALSE.Id(),0}
                  } else /* arg:9 */{ 
                    g1040I_try10417 = EID{CTRUE.Id(),0}/* arg-9 */} 
                  /* arg-8 */} 
                }
                /* and-7 */} 
              /* ERROR PROTECTION INSERTED (g1040I-Result) */
              if ErrorIn(g1040I_try10417) {Result = g1040I_try10417
              } else {
              g1040I = ToBoolean(OBJ(g1040I_try10417))
              if (g1040I == CTRUE) /* If:7 */{ 
                /* Let:8 */{ 
                  var _CL_obj *Language.Let   = Language.To_Let(new(Language.Let).Is(Language.C_Let))
                  /* noccur = 12 */
                  _CL_obj.ClaireVar = v
                  _CL_obj.Value = d
                  /* update:9 */{ 
                    var va_arg1 *Language.Let  
                    var va_arg2 *ClaireAny  
                    va_arg1 = _CL_obj
                    var va_arg2_try104510 EID 
                    /* Let:10 */{ 
                      var g1046UU *Language.CallMethod2  
                      /* noccur = 1 */
                      var g1046UU_try104911 EID 
                      /* Let:11 */{ 
                        var _CL_obj *Language.CallMethod2   = Language.To_CallMethod2(new(Language.CallMethod2).Is(Language.C_Call_method2))
                        /* noccur = 5 */
                        _CL_obj.Arg = ToMethod(Core.F__at_property1(Core.C_identical_ask,C_any).Id())
                        /* update:12 */{ 
                          var va_arg1 *Language.CallMethod  
                          var va_arg2 *ClaireList  
                          va_arg1 = Language.To_CallMethod(_CL_obj.Id())
                          var va_arg2_try105013 EID 
                          /* Construct:13 */{ 
                            var v_bag_arg *ClaireAny  
                            va_arg2_try105013= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                            ToList(OBJ(va_arg2_try105013)).AddFast(v.Id())
                            var v_bag_arg_try105114 EID 
                            /* Let:14 */{ 
                              var g1052UU *ClaireAny  
                              /* noccur = 1 */
                              var g1052UU_try105315 EID 
                              g1052UU_try105315 = Core.F_CALL(C_Compile_c_sort,ARGS(EID{v.Id(),0}))
                              /* ERROR PROTECTION INSERTED (g1052UU-v_bag_arg_try105114) */
                              if ErrorIn(g1052UU_try105315) {v_bag_arg_try105114 = g1052UU_try105315
                              } else {
                              g1052UU = ANY(g1052UU_try105315)
                              v_bag_arg_try105114 = Core.F_CALL(C_c_code,ARGS(EID{CNULL,0},g1052UU.ToEID()))
                              }
                              /* Let-14 */} 
                            /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try105013) */
                            if ErrorIn(v_bag_arg_try105114) {va_arg2_try105013 = v_bag_arg_try105114
                            } else {
                            v_bag_arg = ANY(v_bag_arg_try105114)
                            ToList(OBJ(va_arg2_try105013)).AddFast(v_bag_arg)}
                            /* Construct-13 */} 
                          /* ERROR PROTECTION INSERTED (va_arg2-g1046UU_try104911) */
                          if ErrorIn(va_arg2_try105013) {g1046UU_try104911 = va_arg2_try105013
                          } else {
                          va_arg2 = ToList(OBJ(va_arg2_try105013))
                          /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                          va_arg1.Args = va_arg2
                          g1046UU_try104911 = EID{va_arg2.Id(),0}
                          }
                          /* update-12 */} 
                        /* ERROR PROTECTION INSERTED (g1046UU_try104911-g1046UU_try104911) */
                        if !ErrorIn(g1046UU_try104911) {
                        g1046UU_try104911 = EID{_CL_obj.Id(),0}
                        }
                        /* Let-11 */} 
                      /* ERROR PROTECTION INSERTED (g1046UU-va_arg2_try104510) */
                      if ErrorIn(g1046UU_try104911) {va_arg2_try104510 = g1046UU_try104911
                      } else {
                      g1046UU = Language.To_CallMethod2(OBJ(g1046UU_try104911))
                      /* Let:11 */{ 
                        var g1047UU *ClaireAny  
                        /* noccur = 1 */
                        var g1047UU_try105412 EID 
                        g1047UU_try105412 = Core.F_CALL(C_c_code,ARGS(self.Other.ToEID(),EID{s.Id(),0}))
                        /* ERROR PROTECTION INSERTED (g1047UU-va_arg2_try104510) */
                        if ErrorIn(g1047UU_try105412) {va_arg2_try104510 = g1047UU_try105412
                        } else {
                        g1047UU = ANY(g1047UU_try105412)
                        /* Let:12 */{ 
                          var g1048UU *ClaireAny  
                          /* noccur = 1 */
                          var g1048UU_try105513 EID 
                          g1048UU_try105513 = Core.F_CALL(C_c_code,ARGS(self.Arg.ToEID(),EID{s.Id(),0}))
                          /* ERROR PROTECTION INSERTED (g1048UU-va_arg2_try104510) */
                          if ErrorIn(g1048UU_try105513) {va_arg2_try104510 = g1048UU_try105513
                          } else {
                          g1048UU = ANY(g1048UU_try105513)
                          va_arg2_try104510 = Language.C_If.Make(g1046UU.Id(),g1047UU,g1048UU).ToEID()
                          }
                          /* Let-12 */} 
                        }
                        /* Let-11 */} 
                      }
                      /* Let-10 */} 
                    /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                    if ErrorIn(va_arg2_try104510) {Result = va_arg2_try104510
                    } else {
                    va_arg2 = ANY(va_arg2_try104510)
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
                var g1056I *ClaireBoolean  
                var g1056I_try10578 EID 
                /* and:8 */{ 
                  var v_and8 *ClaireBoolean  
                  
                  var v_and8_try10589 EID 
                  /* Let:9 */{ 
                    var g1059UU *ClaireAny  
                    /* noccur = 1 */
                    var g1059UU_try106010 EID 
                    g1059UU_try106010 = Core.F_CALL(C_Compile_c_sort,ARGS(EID{v.Id(),0}))
                    /* ERROR PROTECTION INSERTED (g1059UU-v_and8_try10589) */
                    if ErrorIn(g1059UU_try106010) {v_and8_try10589 = g1059UU_try106010
                    } else {
                    g1059UU = ANY(g1059UU_try106010)
                    v_and8_try10589 = EID{Equal(g1059UU,C_any.Id()).Id(),0}
                    }
                    /* Let-9 */} 
                  /* ERROR PROTECTION INSERTED (v_and8-g1056I_try10578) */
                  if ErrorIn(v_and8_try10589) {g1056I_try10578 = v_and8_try10589
                  } else {
                  v_and8 = ToBoolean(OBJ(v_and8_try10589))
                  if (v_and8 == CFALSE) {g1056I_try10578 = EID{CFALSE.Id(),0}
                  } else /* arg:9 */{ 
                    v_and8 = ToType(_Ztype).Included(v.Range)
                    if (v_and8 == CFALSE) {g1056I_try10578 = EID{CFALSE.Id(),0}
                    } else /* arg:10 */{ 
                      v_and8 = F__sup_equal_integer(C_compiler.Safety,3)
                      if (v_and8 == CFALSE) {g1056I_try10578 = EID{CFALSE.Id(),0}
                      } else /* arg:11 */{ 
                        g1056I_try10578 = EID{CTRUE.Id(),0}/* arg-11 */} 
                      /* arg-10 */} 
                    /* arg-9 */} 
                  }
                  /* and-8 */} 
                /* ERROR PROTECTION INSERTED (g1056I-Result) */
                if ErrorIn(g1056I_try10578) {Result = g1056I_try10578
                } else {
                g1056I = ToBoolean(OBJ(g1056I_try10578))
                if (g1056I == CTRUE) /* If:8 */{ 
                  /* Let:9 */{ 
                    var g1061UU *Language.Let  
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
                          var g1062UU *Language.Call  
                          /* noccur = 1 */
                          /* Let:13 */{ 
                            var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                            /* noccur = 5 */
                            _CL_obj.Selector = ToProperty(Core.C__I_equal.Id())
                            _CL_obj.Args = MakeConstantList(v.Id(),CNULL)
                            g1062UU = _CL_obj
                            /* Let-13 */} 
                          va_arg2 = Language.C_If.Make(g1062UU.Id(),self.Arg,self.Other)
                          /* Let-12 */} 
                        /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                        va_arg1.Arg = va_arg2
                        /* update-11 */} 
                      g1061UU = _CL_obj
                      /* Let-10 */} 
                    Result = Core.F_CALL(C_c_code,ARGS(EID{g1061UU.Id(),0},EID{s.Id(),0}))
                    /* Let-9 */} 
                  } else {
                  /* Let:9 */{ 
                    var g1063UU *Language.Let  
                    /* noccur = 1 */
                    var g1063UU_try106410 EID 
                    /* Let:10 */{ 
                      var _CL_obj *Language.Let   = Language.To_Let(new(Language.Let).Is(Language.C_Let))
                      /* noccur = 24 */
                      _CL_obj.ClaireVar = v2
                      _CL_obj.Value = _Zv
                      /* update:11 */{ 
                        var va_arg1 *Language.Let  
                        var va_arg2 *ClaireAny  
                        va_arg1 = _CL_obj
                        var va_arg2_try106512 EID 
                        /* Let:12 */{ 
                          var g1066UU *Language.Call  
                          /* noccur = 1 */
                          /* Let:13 */{ 
                            var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                            /* noccur = 5 */
                            _CL_obj.Selector = ToProperty(Core.C__I_equal.Id())
                            _CL_obj.Args = MakeConstantList(v2.Id(),CNULL)
                            g1066UU = _CL_obj
                            /* Let-13 */} 
                          /* Let:13 */{ 
                            var g1067UU *Language.Let  
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
                              g1067UU = _CL_obj
                              /* Let-14 */} 
                            /* Let:14 */{ 
                              var g1068UU *ClaireAny  
                              /* noccur = 1 */
                              var g1068UU_try106915 EID 
                              g1068UU_try106915 = Core.F_CALL(C_c_code,ARGS(self.Other.ToEID(),EID{s.Id(),0}))
                              /* ERROR PROTECTION INSERTED (g1068UU-va_arg2_try106512) */
                              if ErrorIn(g1068UU_try106915) {va_arg2_try106512 = g1068UU_try106915
                              } else {
                              g1068UU = ANY(g1068UU_try106915)
                              va_arg2_try106512 = Language.C_If.Make(g1066UU.Id(),g1067UU.Id(),g1068UU).ToEID()
                              }
                              /* Let-14 */} 
                            /* Let-13 */} 
                          /* Let-12 */} 
                        /* ERROR PROTECTION INSERTED (va_arg2-g1063UU_try106410) */
                        if ErrorIn(va_arg2_try106512) {g1063UU_try106410 = va_arg2_try106512
                        } else {
                        va_arg2 = ANY(va_arg2_try106512)
                        /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                        va_arg1.Arg = va_arg2
                        g1063UU_try106410 = va_arg2.ToEID()
                        }
                        /* update-11 */} 
                      /* ERROR PROTECTION INSERTED (g1063UU_try106410-g1063UU_try106410) */
                      if !ErrorIn(g1063UU_try106410) {
                      g1063UU_try106410 = EID{_CL_obj.Id(),0}
                      }
                      /* Let-10 */} 
                    /* ERROR PROTECTION INSERTED (g1063UU-Result) */
                    if ErrorIn(g1063UU_try106410) {Result = g1063UU_try106410
                    } else {
                    g1063UU = Language.To_Let(OBJ(g1063UU_try106410))
                    Result = Core.F_CALL(C_c_code,ARGS(EID{g1063UU.Id(),0},EID{s.Id(),0}))
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
      var g1070UU *ClaireType  
      /* noccur = 1 */
      var g1070UU_try10713 EID 
      g1070UU_try10713 = F_Compile_return_type_any(self.Arg)
      /* ERROR PROTECTION INSERTED (g1070UU-Result) */
      if ErrorIn(g1070UU_try10713) {Result = g1070UU_try10713
      } else {
      g1070UU = ToType(OBJ(g1070UU_try10713))
      Result = EID{F_Optimize_infers_from_type(g1070UU,self.Id()).Id(),0}
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
                var scs_try10788 EID 
                scs_try10788 = F_Optimize_c_inline_arg_ask_any(sx)
                /* ERROR PROTECTION INSERTED (scs-Result) */
                if ErrorIn(scs_try10788) {Result = scs_try10788
                } else {
                scs = ANY(scs_try10788)
                if (sx.Isa.IsIn(Core.C_global_variable) == CTRUE) /* If:8 */{ 
                  /* Let:9 */{ 
                    var g1072 *Core.GlobalVariable   = Core.ToGlobalVariable(sx)
                    /* noccur = 5 */
                    if (F_boolean_I_any(g1072.Range.Id()).Id() != CTRUE.Id()) /* If:10 */{ 
                      self.SetArg = g1072.Value
                      g1072 = Core.ToGlobalVariable(g1072.Value)
                      /* If-10 */} 
                    sx = g1072.Id()
                    Result = sx.ToEID()
                    /* Let-9 */} 
                  /* If!8 */}  else if (sx.Isa.IsIn(Language.C_Select) == CTRUE) /* If:8 */{ 
                  /* Let:9 */{ 
                    var g1073 *Language.Select   = Language.To_Select(sx)
                    /* noccur = 1 */
                    /* Let:10 */{ 
                      var _Zt *ClaireType  
                      /* noccur = 2 */
                      var _Zt_try107911 EID 
                      _Zt_try107911 = Core.F_CALL(C_c_type,ARGS(EID{g1073.Id(),0}))
                      /* ERROR PROTECTION INSERTED (_Zt-Result) */
                      if ErrorIn(_Zt_try107911) {Result = _Zt_try107911
                      } else {
                      _Zt = ToType(OBJ(_Zt_try107911))
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
                    var g1074 *ClaireClass   = ToClass(sx)
                    /* noccur = 3 */
                    if ((g1074.Open <= 1) && 
                        (F_boolean_I_any(g1074.Subclass.Id()).Id() != CTRUE.Id())) /* If:10 */{ 
                      /* update:11 */{ 
                        var va_arg1 *Language.Iteration  
                        var va_arg2 *ClaireAny  
                        va_arg1 = Language.To_Iteration(self.Id())
                        /* Let:12 */{ 
                          var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                          /* noccur = 5 */
                          _CL_obj.Selector = C_instances
                          _CL_obj.Args = MakeConstantList(g1074.Id())
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
                  var _Zt_try10809 EID 
                  _Zt_try10809 = Core.F_CALL(C_c_type,ARGS(self.SetArg.ToEID()))
                  /* ERROR PROTECTION INSERTED (_Zt-Result) */
                  if ErrorIn(_Zt_try10809) {Result = _Zt_try10809
                  } else {
                  _Zt = ANY(_Zt_try10809)
                  /* Let:9 */{ 
                    var _Zt2 *ClaireType   = F_Optimize_pmember_type(ToType(_Zt))
                    /* noccur = 6 */
                    /* Let:10 */{ 
                      var _Zt3 *ClaireType  
                      /* noccur = 2 */
                      var _Zt3_try108111 EID 
                      /* Let:11 */{ 
                        var g1082UU *ClaireType  
                        /* noccur = 1 */
                        var g1082UU_try108312 EID 
                        /* Let:12 */{ 
                          var g1084UU *Language.Call  
                          /* noccur = 1 */
                          /* Let:13 */{ 
                            var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                            /* noccur = 5 */
                            _CL_obj.Selector = C_set_I
                            _CL_obj.Args = MakeConstantList(sx)
                            g1084UU = _CL_obj
                            /* Let-13 */} 
                          g1082UU_try108312 = Core.F_CALL(C_c_type,ARGS(EID{g1084UU.Id(),0}))
                          /* Let-12 */} 
                        /* ERROR PROTECTION INSERTED (g1082UU-_Zt3_try108111) */
                        if ErrorIn(g1082UU_try108312) {_Zt3_try108111 = g1082UU_try108312
                        } else {
                        g1082UU = ToType(OBJ(g1082UU_try108312))
                        _Zt3_try108111 = EID{F_Optimize_pmember_type(g1082UU).Id(),0}
                        }
                        /* Let-11 */} 
                      /* ERROR PROTECTION INSERTED (_Zt3-Result) */
                      if ErrorIn(_Zt3_try108111) {Result = _Zt3_try108111
                      } else {
                      _Zt3 = ToType(OBJ(_Zt3_try108111))
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
                                var g1075 *ClaireMethod   = ToMethod(m2)
                                /* noccur = 1 */
                                m = g1075.Id()
                                /* Let-15 */} 
                              /* If-14 */} 
                            /* Let-13 */} 
                          /* If-12 */} 
                        C_compiler.Safety = ns
                        v.Range = _Zt2
                        var g1087I *ClaireBoolean  
                        if (C_method.Id() == m.Isa.Id()) /* If:12 */{ 
                          /* Let:13 */{ 
                            var g1076 *ClaireMethod   = ToMethod(m)
                            /* noccur = 1 */
                            g1087I = g1076.Inline_ask
                            /* Let-13 */} 
                          } else {
                          g1087I = CFALSE
                          /* If-12 */} 
                        if (g1087I == CTRUE) /* If:12 */{ 
                          
                          if (F_Optimize_sort_abstract_ask_type(v.Range) == CTRUE) /* If:13 */{ 
                            v.Range = To_Union(v.Range.Id()).T2
                            /* If-13 */} 
                          Result = F_Optimize_c_inline_method1(ToMethod(m),MakeConstantList(Language.F_instruction_copy_any(self.SetArg),v.Id(),narg),s)
                          /* If!12 */}  else if (F_boolean_I_any(scs) == CTRUE) /* If:12 */{ 
                          /* Let:13 */{ 
                            var g1088UU *Language.For  
                            /* noccur = 1 */
                            /* Let:14 */{ 
                              var _CL_obj *Language.For   = Language.To_For(new(Language.For).Is(Language.C_For))
                              /* noccur = 5 */
                              _CL_obj.ClaireVar = v
                              _CL_obj.SetArg = scs
                              _CL_obj.Arg = narg
                              g1088UU = _CL_obj
                              /* Let-14 */} 
                            Result = Core.F_CALL(C_c_code,ARGS(EID{g1088UU.Id(),0},EID{s.Id(),0}))
                            /* Let-13 */} 
                          } else {
                          var g1089I *ClaireBoolean  
                          if (sx.Isa.IsIn(Language.C_Call) == CTRUE) /* If:13 */{ 
                            /* Let:14 */{ 
                              var g1077 *Language.Call   = Language.To_Call(sx)
                              /* noccur = 1 */
                              g1089I = Equal(g1077.Selector.Id(),Core.C_Id.Id())
                              /* Let-14 */} 
                            } else {
                            g1089I = CFALSE
                            /* If-13 */} 
                          if (g1089I == CTRUE) /* If:13 */{ 
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
                                var va_arg2_try109016 EID 
                                va_arg2_try109016 = F_Optimize_enumerate_code_any(self.SetArg,ToType(_Zt))
                                /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                                if ErrorIn(va_arg2_try109016) {Result = va_arg2_try109016
                                } else {
                                va_arg2 = ANY(va_arg2_try109016)
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
                                var va_arg2_try109116 EID 
                                va_arg2_try109116 = Core.F_CALL(C_c_code,ARGS(narg.ToEID(),EID{C_void.Id(),0}))
                                /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                                if ErrorIn(va_arg2_try109116) {Result = va_arg2_try109116
                                } else {
                                va_arg2 = ANY(va_arg2_try109116)
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
            var n_try10926 EID 
            /* Let:6 */{ 
              var _CL_obj *Language.Let   = Language.To_Let(new(Language.Let).Is(Language.C_Let))
              /* noccur = 12 */
              _CL_obj.ClaireVar = v2
              _CL_obj.Value = sx
              /* update:7 */{ 
                var va_arg1 *Language.Let  
                var va_arg2 *ClaireAny  
                va_arg1 = _CL_obj
                var va_arg2_try10938 EID 
                /* Let:8 */{ 
                  var _CL_obj *Language.For   = Language.To_For(new(Language.For).Is(Language.C_For))
                  /* noccur = 5 */
                  _CL_obj.ClaireVar = self.ClaireVar
                  /* update:9 */{ 
                    var va_arg1 *Language.Iteration  
                    var va_arg2 *ClaireAny  
                    va_arg1 = Language.To_Iteration(_CL_obj.Id())
                    var va_arg2_try109410 EID 
                    va_arg2_try109410 = F_Optimize_enumerate_code_any(sx,_Zt)
                    /* ERROR PROTECTION INSERTED (va_arg2-va_arg2_try10938) */
                    if ErrorIn(va_arg2_try109410) {va_arg2_try10938 = va_arg2_try109410
                    } else {
                    va_arg2 = ANY(va_arg2_try109410)
                    /* ---------- now we compile update iClaire/set_arg(va_arg1) := va_arg2 ------- */
                    va_arg1.SetArg = va_arg2
                    va_arg2_try10938 = va_arg2.ToEID()
                    }
                    /* update-9 */} 
                  /* ERROR PROTECTION INSERTED (va_arg2_try10938-va_arg2_try10938) */
                  if !ErrorIn(va_arg2_try10938) {
                  _CL_obj.Arg = self.Arg
                  va_arg2_try10938 = EID{_CL_obj.Id(),0}
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (va_arg2-n_try10926) */
                if ErrorIn(va_arg2_try10938) {n_try10926 = va_arg2_try10938
                } else {
                va_arg2 = ANY(va_arg2_try10938)
                /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                va_arg1.Arg = va_arg2
                n_try10926 = va_arg2.ToEID()
                }
                /* update-7 */} 
              /* ERROR PROTECTION INSERTED (n_try10926-n_try10926) */
              if !ErrorIn(n_try10926) {
              n_try10926 = EID{_CL_obj.Id(),0}
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (n-Result) */
            if ErrorIn(n_try10926) {Result = n_try10926
            } else {
            n = Language.To_Let(OBJ(n_try10926))
            Core.F_tformat_string(MakeString("---- note: use an expended iteration for {~S} \n"),0,MakeConstantList(self.Id()))
            /* For:6 */{ 
              var r *ClaireAny  
              _ = r
              Result= EID{CFALSE.Id(),0}
              for _,r = range(Language.C_iterate.Restrictions.ValuesO())/* loop:7 */{ 
                var void_try8 EID 
                _ = void_try8
                var g1095I *ClaireBoolean  
                var g1095I_try10968 EID 
                /* and:8 */{ 
                  var v_and8 *ClaireBoolean  
                  
                  v_and8 = ToType(Core.F_domain_I_restriction(ToRestriction(r)).Id()).Included(_Zt)
                  if (v_and8 == CFALSE) {g1095I_try10968 = EID{CFALSE.Id(),0}
                  } else /* arg:9 */{ 
                    v_and8 = ToType(Core.F_domain_I_restriction(ToRestriction(r)).Id()).Included(ToType(C_collection.Id()))
                    if (v_and8 == CFALSE) {g1095I_try10968 = EID{CFALSE.Id(),0}
                    } else /* arg:10 */{ 
                      v_and8 = ToMethod(r).Inline_ask
                      if (v_and8 == CFALSE) {g1095I_try10968 = EID{CFALSE.Id(),0}
                      } else /* arg:11 */{ 
                        var v_and8_try109712 EID 
                        v_and8_try109712 = Core.F_BELONG(v.Id(),ToRestriction(r).Domain.ValuesO()[2-1])
                        /* ERROR PROTECTION INSERTED (v_and8-g1095I_try10968) */
                        if ErrorIn(v_and8_try109712) {g1095I_try10968 = v_and8_try109712
                        } else {
                        v_and8 = ToBoolean(OBJ(v_and8_try109712))
                        if (v_and8 == CFALSE) {g1095I_try10968 = EID{CFALSE.Id(),0}
                        } else /* arg:12 */{ 
                          g1095I_try10968 = EID{CTRUE.Id(),0}/* arg-12 */} 
                        /* arg-11 */} 
                      /* arg-10 */} 
                    /* arg-9 */} 
                  }
                  /* and-8 */} 
                /* ERROR PROTECTION INSERTED (g1095I-void_try8) */
                if ErrorIn(g1095I_try10968) {void_try8 = g1095I_try10968
                } else {
                g1095I = ToBoolean(OBJ(g1095I_try10968))
                if (g1095I == CTRUE) /* If:8 */{ 
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
                        var va_arg2_try109812 EID 
                        /* Let:12 */{ 
                          var g1099UU *Language.Call  
                          /* noccur = 1 */
                          /* Let:13 */{ 
                            var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                            /* noccur = 5 */
                            _CL_obj.Selector = ToProperty(C__Z.Id())
                            _CL_obj.Args = MakeConstantList(v2.Id(),ToRestriction(r).Domain.ValuesO()[1-1])
                            g1099UU = _CL_obj
                            /* Let-13 */} 
                          /* Let:13 */{ 
                            var g1100UU *ClaireAny  
                            /* noccur = 1 */
                            var g1100UU_try110114 EID 
                            if (F_Optimize_sort_abstract_ask_type(vnew.Range) == CTRUE) /* If:14 */{ 
                              vnew.Range = To_Union(v.Range.Id()).T2
                              /* If-14 */} 
                            g1100UU_try110114 = F_Optimize_c_inline_method1(ToMethod(r),MakeConstantList(v2.Id(),vnew.Id(),narg),s)
                            /* ERROR PROTECTION INSERTED (g1100UU-va_arg2_try109812) */
                            if ErrorIn(g1100UU_try110114) {va_arg2_try109812 = g1100UU_try110114
                            } else {
                            g1100UU = ANY(g1100UU_try110114)
                            va_arg2_try109812 = Language.C_If.Make(g1099UU.Id(),g1100UU,n.Arg).ToEID()
                            }
                            /* Let-13 */} 
                          /* Let-12 */} 
                        /* ERROR PROTECTION INSERTED (va_arg2-void_try8) */
                        if ErrorIn(va_arg2_try109812) {void_try8 = va_arg2_try109812
                        } else {
                        va_arg2 = ANY(va_arg2_try109812)
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
      var _Zt_try11023 EID 
      _Zt_try11023 = Core.F_CALL(C_c_type,ARGS(self.SetArg.ToEID()))
      /* ERROR PROTECTION INSERTED (_Zt-Result) */
      if ErrorIn(_Zt_try11023) {Result = _Zt_try11023
      } else {
      _Zt = ANY(_Zt_try11023)
      F_Optimize_range_infers_for_Variable(self.ClaireVar,F_Optimize_pmember_type(ToType(_Zt)),ToType(_Zt))
      if (Core.F_get_property(C_of,ToObject(self.Id())) != CNULL) /* If:3 */{ 
        /* Let:4 */{ 
          var g1103UU *ClaireClass  
          /* noccur = 1 */
          if ((self.Isa.IsIn(Language.C_Select) == CTRUE) || 
              (self.Isa.IsIn(Language.C_Image) == CTRUE)) /* If:5 */{ 
            g1103UU = C_set
            } else {
            g1103UU = C_list
            /* If-5 */} 
          Result = EID{Core.F_param_I_class(g1103UU,ToType(OBJ(Core.F_CALL(C_of,ARGS(EID{self.Id(),0}))))).Id(),0}
          /* Let-4 */} 
        } else {
        /* Let:4 */{ 
          var g1104UU *ClaireClass  
          /* noccur = 1 */
          if ((self.Isa.IsIn(Language.C_Select) == CTRUE) || 
              (self.Isa.IsIn(Language.C_Image) == CTRUE)) /* If:5 */{ 
            g1104UU = C_set
            } else {
            g1104UU = C_list
            /* If-5 */} 
          /* Let:5 */{ 
            var g1105UU *ClaireType  
            /* noccur = 1 */
            var g1105UU_try11066 EID 
            if ((self.Isa.IsIn(Language.C_Select) == CTRUE) || 
                (self.Isa.IsIn(Language.C_Lselect) == CTRUE)) /* If:6 */{ 
              /* Let:7 */{ 
                var g1107UU *ClaireType  
                /* noccur = 1 */
                var g1107UU_try11088 EID 
                g1107UU_try11088 = Core.F_CALL(C_c_type,ARGS(self.SetArg.ToEID()))
                /* ERROR PROTECTION INSERTED (g1107UU-g1105UU_try11066) */
                if ErrorIn(g1107UU_try11088) {g1105UU_try11066 = g1107UU_try11088
                } else {
                g1107UU = ToType(OBJ(g1107UU_try11088))
                g1105UU_try11066 = EID{F_Optimize_pmember_type(g1107UU).Id(),0}
                }
                /* Let-7 */} 
              } else {
              /* Let:7 */{ 
                var g1109UU *ClaireType  
                /* noccur = 1 */
                var g1109UU_try11108 EID 
                g1109UU_try11108 = Core.F_CALL(C_c_type,ARGS(self.Arg.ToEID()))
                /* ERROR PROTECTION INSERTED (g1109UU-g1105UU_try11066) */
                if ErrorIn(g1109UU_try11108) {g1105UU_try11066 = g1109UU_try11108
                } else {
                g1109UU = ToType(OBJ(g1109UU_try11108))
                g1105UU_try11066 = EID{F_Optimize_ptype_type(g1109UU).Id(),0}
                }
                /* Let-7 */} 
              /* If-6 */} 
            /* ERROR PROTECTION INSERTED (g1105UU-Result) */
            if ErrorIn(g1105UU_try11066) {Result = g1105UU_try11066
            } else {
            g1105UU = ToType(OBJ(g1105UU_try11066))
            Result = EID{Core.F_nth_class1(g1104UU,g1105UU).Id(),0}
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
        var _Zt_try11134 EID 
        _Zt_try11134 = Core.F_CALL(C_c_type,ARGS(sx.ToEID()))
        /* ERROR PROTECTION INSERTED (_Zt-Result) */
        if ErrorIn(_Zt_try11134) {Result = _Zt_try11134
        } else {
        _Zt = ANY(_Zt_try11134)
        if (self.Isa.IsIn(Language.C_For) == CTRUE) /* If:4 */{ 
          Result = Core.F_CALL(C_c_code,ARGS(EID{self.Id(),0},EID{C_any.Id(),0}))
          } else {
          var g1114I *ClaireBoolean  
          if (self.Isa.IsIn(Language.C_Collect) == CTRUE) /* If:5 */{ 
            g1114I = MakeBoolean((ToType(_Zt).Included(ToType(C_list.Id())) == CTRUE) || (ToType(_Zt).Included(ToType(C_set.Id())) == CTRUE))
            } else {
            g1114I = CFALSE
            /* If-5 */} 
          if (g1114I == CTRUE) /* If:5 */{ 
            F_Optimize_range_infers_for_Variable(self.ClaireVar,F_Optimize_pmember_type(ToType(_Zt)),ToType(_Zt))
            /* Let:6 */{ 
              var ty *ClaireType  
              /* noccur = 2 */
              var ty_try11157 EID 
              /* Let:7 */{ 
                var g1116UU *ClaireType  
                /* noccur = 1 */
                var g1116UU_try11178 EID 
                g1116UU_try11178 = Core.F_CALL(C_c_type,ARGS(self.Arg.ToEID()))
                /* ERROR PROTECTION INSERTED (g1116UU-ty_try11157) */
                if ErrorIn(g1116UU_try11178) {ty_try11157 = g1116UU_try11178
                } else {
                g1116UU = ToType(OBJ(g1116UU_try11178))
                ty_try11157 = EID{F_Optimize_ptype_type(g1116UU).Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (ty-Result) */
              if ErrorIn(ty_try11157) {Result = ty_try11157
              } else {
              ty = ToType(OBJ(ty_try11157))
              /* Let:7 */{ 
                var x *Language.Collect  
                /* noccur = 5 */
                var x_try11188 EID 
                /* Let:8 */{ 
                  var _CL_obj *Language.Collect   = Language.To_Collect(new(Language.Collect).Is(Language.C_Collect))
                  /* noccur = 5 */
                  _CL_obj.ClaireVar = self.ClaireVar
                  /* update:9 */{ 
                    var va_arg1 *Language.Iteration  
                    var va_arg2 *ClaireAny  
                    va_arg1 = Language.To_Iteration(_CL_obj.Id())
                    var va_arg2_try111910 EID 
                    va_arg2_try111910 = F_Compile_c_strict_code_any(sx,ToTypeExpression(_Zt).Class_I())
                    /* ERROR PROTECTION INSERTED (va_arg2-x_try11188) */
                    if ErrorIn(va_arg2_try111910) {x_try11188 = va_arg2_try111910
                    } else {
                    va_arg2 = ANY(va_arg2_try111910)
                    /* ---------- now we compile update iClaire/set_arg(va_arg1) := va_arg2 ------- */
                    va_arg1.SetArg = va_arg2
                    x_try11188 = va_arg2.ToEID()
                    }
                    /* update-9 */} 
                  /* ERROR PROTECTION INSERTED (x_try11188-x_try11188) */
                  if !ErrorIn(x_try11188) {
                  /* update:9 */{ 
                    var va_arg1 *Language.Iteration  
                    var va_arg2 *ClaireAny  
                    va_arg1 = Language.To_Iteration(_CL_obj.Id())
                    var va_arg2_try112010 EID 
                    va_arg2_try112010 = Core.F_CALL(C_c_code,ARGS(self.Arg.ToEID(),EID{C_any.Id(),0}))
                    /* ERROR PROTECTION INSERTED (va_arg2-x_try11188) */
                    if ErrorIn(va_arg2_try112010) {x_try11188 = va_arg2_try112010
                    } else {
                    va_arg2 = ANY(va_arg2_try112010)
                    /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                    va_arg1.Arg = va_arg2
                    x_try11188 = va_arg2.ToEID()
                    }
                    /* update-9 */} 
                  /* ERROR PROTECTION INSERTED (x_try11188-x_try11188) */
                  if !ErrorIn(x_try11188) {
                  x_try11188 = EID{_CL_obj.Id(),0}
                  }}
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (x-Result) */
                if ErrorIn(x_try11188) {Result = x_try11188
                } else {
                x = Language.To_Collect(OBJ(x_try11188))
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
                      var g1121UU *ClaireList  
                      /* noccur = 1 */
                      var g1121UU_try112211 EID 
                      /* Construct:11 */{ 
                        var v_bag_arg *ClaireAny  
                        g1121UU_try112211= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                        ToList(OBJ(g1121UU_try112211)).AddFast(self.Id())
                        var v_bag_arg_try112312 EID 
                        v_bag_arg_try112312 = Core.F_CALL(C_c_type,ARGS(self.Arg.ToEID()))
                        /* ERROR PROTECTION INSERTED (v_bag_arg-g1121UU_try112211) */
                        if ErrorIn(v_bag_arg_try112312) {g1121UU_try112211 = v_bag_arg_try112312
                        } else {
                        v_bag_arg = ANY(v_bag_arg_try112312)
                        ToList(OBJ(g1121UU_try112211)).AddFast(v_bag_arg)
                        ToList(OBJ(g1121UU_try112211)).AddFast(ANY(Core.F_CALL(C_of,ARGS(EID{self.Id(),0}))))}
                        /* Construct-11 */} 
                      /* ERROR PROTECTION INSERTED (g1121UU-Result) */
                      if ErrorIn(g1121UU_try112211) {Result = g1121UU_try112211
                      } else {
                      g1121UU = ToList(OBJ(g1121UU_try112211))
                      Result = Core.F_tformat_string(MakeString("unsafe typed collect (~S): ~S not in ~S [261]\n"),2,g1121UU)
                      }
                      /* Let-10 */} 
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    /* Let:10 */{ 
                      var g1124UU *Language.Call  
                      /* noccur = 1 */
                      /* Let:11 */{ 
                        var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                        /* noccur = 5 */
                        _CL_obj.Selector = Core.C_check_in
                        _CL_obj.Args = MakeConstantList(x.Id(),C_list.Id(),ANY(Core.F_CALL(C_of,ARGS(EID{self.Id(),0}))))
                        g1124UU = _CL_obj
                        /* Let-11 */} 
                      Result = Core.F_CALL(C_c_code,ARGS(EID{g1124UU.Id(),0},EID{C_list.Id(),0}))
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
                  var g1125UU int 
                  /* noccur = 1 */
                  C_OPT.MaxVars = (C_OPT.MaxVars+1)
                  g1125UU = 0
                  /* Let:9 */{ 
                    var g1126UU *ClaireClass  
                    /* noccur = 1 */
                    if (self.Isa.IsIn(Language.C_Image) == CTRUE) /* If:10 */{ 
                      g1126UU = C_set
                      } else {
                      g1126UU = C_list
                      /* If-10 */} 
                    v = F_Compile_Variable_I_symbol(F_append_symbol(self.ClaireVar.Pname,MakeString("_bag").Id()),g1125UU,g1126UU.Id())
                    /* Let-9 */} 
                  /* Let-8 */} 
                
                if (Core.F_get_property(C_of,ToObject(self.Id())) != CNULL) /* If:8 */{ 
                  /* Let:9 */{ 
                    var _ZtypeIn *ClaireType  
                    /* noccur = 2 */
                    var _ZtypeIn_try112710 EID 
                    _ZtypeIn_try112710 = Core.F_CALL(C_c_type,ARGS(self.Arg.ToEID()))
                    /* ERROR PROTECTION INSERTED (_ZtypeIn-Result) */
                    if ErrorIn(_ZtypeIn_try112710) {Result = _ZtypeIn_try112710
                    } else {
                    _ZtypeIn = ToType(OBJ(_ZtypeIn_try112710))
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
                    var va_arg2_try112810 EID 
                    /* Let:10 */{ 
                      var _CL_obj *Language.Do   = Language.To_Do(new(Language.Do).Is(Language.C_Do))
                      /* noccur = 15 */
                      /* update:11 */{ 
                        var va_arg1 *Language.Do  
                        var va_arg2 *ClaireList  
                        va_arg1 = _CL_obj
                        var va_arg2_try112912 EID 
                        /* Construct:12 */{ 
                          var v_bag_arg *ClaireAny  
                          va_arg2_try112912= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                          var v_bag_arg_try113013 EID 
                          /* Let:13 */{ 
                            var g1131UU *Language.For  
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
                              g1131UU = _CL_obj
                              /* Let-14 */} 
                            v_bag_arg_try113013 = Core.F_CALL(C_c_code,ARGS(EID{g1131UU.Id(),0},EID{C_any.Id(),0}))
                            /* Let-13 */} 
                          /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try112912) */
                          if ErrorIn(v_bag_arg_try113013) {va_arg2_try112912 = v_bag_arg_try113013
                          } else {
                          v_bag_arg = ANY(v_bag_arg_try113013)
                          ToList(OBJ(va_arg2_try112912)).AddFast(v_bag_arg)
                          ToList(OBJ(va_arg2_try112912)).AddFast(v.Id())}
                          /* Construct-12 */} 
                        /* ERROR PROTECTION INSERTED (va_arg2-va_arg2_try112810) */
                        if ErrorIn(va_arg2_try112912) {va_arg2_try112810 = va_arg2_try112912
                        } else {
                        va_arg2 = ToList(OBJ(va_arg2_try112912))
                        /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                        va_arg1.Args = va_arg2
                        va_arg2_try112810 = EID{va_arg2.Id(),0}
                        }
                        /* update-11 */} 
                      /* ERROR PROTECTION INSERTED (va_arg2_try112810-va_arg2_try112810) */
                      if !ErrorIn(va_arg2_try112810) {
                      va_arg2_try112810 = EID{_CL_obj.Id(),0}
                      }
                      /* Let-10 */} 
                    /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                    if ErrorIn(va_arg2_try112810) {Result = va_arg2_try112810
                    } else {
                    va_arg2 = ANY(va_arg2_try112810)
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
        var _Zt_try11344 EID 
        _Zt_try11344 = Core.F_CALL(C_c_type,ARGS(sx.ToEID()))
        /* ERROR PROTECTION INSERTED (_Zt-Result) */
        if ErrorIn(_Zt_try11344) {Result = _Zt_try11344
        } else {
        _Zt = ANY(_Zt_try11344)
        /* Let:4 */{ 
          var st *ClaireAny  
          /* noccur = 1 */
          var st_try11355 EID 
          st_try11355 = F_Optimize_enumerate_code_any(sx,ToType(_Zt))
          /* ERROR PROTECTION INSERTED (st-Result) */
          if ErrorIn(st_try11355) {Result = st_try11355
          } else {
          st = ANY(st_try11355)
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
                var g1136UU int 
                /* noccur = 1 */
                C_OPT.MaxVars = (C_OPT.MaxVars+1)
                g1136UU = 0
                v1 = F_Compile_Variable_I_symbol(F_append_symbol(self.ClaireVar.Pname,MakeString("_in").Id()),g1136UU,_Zt)
                /* Let-7 */} 
              /* Let:7 */{ 
                var v2 *ClaireVariable  
                /* noccur = 4 */
                /* Let:8 */{ 
                  var g1137UU int 
                  /* noccur = 1 */
                  C_OPT.MaxVars = (C_OPT.MaxVars+1)
                  g1137UU = 0
                  v2 = F_Compile_Variable_I_symbol(F_append_symbol(self.ClaireVar.Pname,MakeString("_out").Id()),g1137UU,x.Id())
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
                      var va_arg2_try113811 EID 
                      /* Let:11 */{ 
                        var g1139UU *CompileCCast  
                        /* noccur = 1 */
                        var g1139UU_try114012 EID 
                        /* Let:12 */{ 
                          var _CL_obj *CompileCCast   = To_CompileCCast(new(CompileCCast).Is(C_Compile_C_cast))
                          /* noccur = 10 */
                          /* update:13 */{ 
                            var va_arg1 *CompileCCast  
                            var va_arg2 *ClaireAny  
                            va_arg1 = _CL_obj
                            var va_arg2_try114114 EID 
                            /* Let:14 */{ 
                              var g1142UU *Language.Call  
                              /* noccur = 1 */
                              /* Let:15 */{ 
                                var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                /* noccur = 5 */
                                _CL_obj.Selector = C_empty
                                _CL_obj.Args = MakeConstantList(v1.Id())
                                g1142UU = _CL_obj
                                /* Let-15 */} 
                              va_arg2_try114114 = Core.F_CALL(C_c_code,ARGS(EID{g1142UU.Id(),0},EID{x.Id(),0}))
                              /* Let-14 */} 
                            /* ERROR PROTECTION INSERTED (va_arg2-g1139UU_try114012) */
                            if ErrorIn(va_arg2_try114114) {g1139UU_try114012 = va_arg2_try114114
                            } else {
                            va_arg2 = ANY(va_arg2_try114114)
                            /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                            va_arg1.Arg = va_arg2
                            g1139UU_try114012 = va_arg2.ToEID()
                            }
                            /* update-13 */} 
                          /* ERROR PROTECTION INSERTED (g1139UU_try114012-g1139UU_try114012) */
                          if !ErrorIn(g1139UU_try114012) {
                          _CL_obj.SetArg = x
                          g1139UU_try114012 = EID{_CL_obj.Id(),0}
                          }
                          /* Let-12 */} 
                        /* ERROR PROTECTION INSERTED (g1139UU-va_arg2_try113811) */
                        if ErrorIn(g1139UU_try114012) {va_arg2_try113811 = g1139UU_try114012
                        } else {
                        g1139UU = To_CompileCCast(OBJ(g1139UU_try114012))
                        va_arg2_try113811 = F_Optimize_inner_select_Iteration(self,v2.Id(),v1.Id(),g1139UU.Id())
                        }
                        /* Let-11 */} 
                      /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                      if ErrorIn(va_arg2_try113811) {Result = va_arg2_try113811
                      } else {
                      va_arg2 = ANY(va_arg2_try113811)
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
                    var g1143UU *Language.Construct  
                    /* noccur = 1 */
                    if (x.Id() == C_set.Id()) /* If:10 */{ 
                      /* Let:11 */{ 
                        var _CL_obj *Language.Set   = Language.To_Set(new(Language.Set).Is(Language.C_Set))
                        /* noccur = 3 */
                        _CL_obj.Of = ToType(CEMPTY.Id())
                        g1143UU = Language.To_Construct(_CL_obj.Id())
                        /* Let-11 */} 
                      } else {
                      /* Let:11 */{ 
                        var _CL_obj *Language.List   = Language.To_List(new(Language.List).Is(Language.C_List))
                        /* noccur = 3 */
                        _CL_obj.Of = ToType(CEMPTY.Id())
                        g1143UU = Language.To_Construct(_CL_obj.Id())
                        /* Let-11 */} 
                      /* If-10 */} 
                    Result = F_Optimize_inner_select_Iteration(self,v2.Id(),sx,g1143UU.Id())
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
        var va_arg2_try11444 EID 
        /* Let:4 */{ 
          var _CL_obj *Language.Do   = Language.To_Do(new(Language.Do).Is(Language.C_Do))
          /* noccur = 21 */
          /* update:5 */{ 
            var va_arg1 *Language.Do  
            var va_arg2 *ClaireList  
            va_arg1 = _CL_obj
            var va_arg2_try11456 EID 
            /* Construct:6 */{ 
              var v_bag_arg *ClaireAny  
              va_arg2_try11456= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
              var v_bag_arg_try11467 EID 
              /* Let:7 */{ 
                var g1147UU *Language.For  
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
                  g1147UU = _CL_obj
                  /* Let-8 */} 
                v_bag_arg_try11467 = Core.F_CALL(C_c_code,ARGS(EID{g1147UU.Id(),0},EID{C_any.Id(),0}))
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try11456) */
              if ErrorIn(v_bag_arg_try11467) {va_arg2_try11456 = v_bag_arg_try11467
              } else {
              v_bag_arg = ANY(v_bag_arg_try11467)
              ToList(OBJ(va_arg2_try11456)).AddFast(v_bag_arg)
              ToList(OBJ(va_arg2_try11456)).AddFast(v2)}
              /* Construct-6 */} 
            /* ERROR PROTECTION INSERTED (va_arg2-va_arg2_try11444) */
            if ErrorIn(va_arg2_try11456) {va_arg2_try11444 = va_arg2_try11456
            } else {
            va_arg2 = ToList(OBJ(va_arg2_try11456))
            /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
            va_arg1.Args = va_arg2
            va_arg2_try11444 = EID{va_arg2.Id(),0}
            }
            /* update-5 */} 
          /* ERROR PROTECTION INSERTED (va_arg2_try11444-va_arg2_try11444) */
          if !ErrorIn(va_arg2_try11444) {
          va_arg2_try11444 = EID{_CL_obj.Id(),0}
          }
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (va_arg2-Result) */
        if ErrorIn(va_arg2_try11444) {Result = va_arg2_try11444
        } else {
        va_arg2 = ANY(va_arg2_try11444)
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
      var _Zt_try11483 EID 
      _Zt_try11483 = Core.F_CALL(C_c_type,ARGS(self.SetArg.ToEID()))
      /* ERROR PROTECTION INSERTED (_Zt-Result) */
      if ErrorIn(_Zt_try11483) {Result = _Zt_try11483
      } else {
      _Zt = ANY(_Zt_try11483)
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
      var _Zt_try11503 EID 
      _Zt_try11503 = Core.F_CALL(C_c_type,ARGS(self.SetArg.ToEID()))
      /* ERROR PROTECTION INSERTED (_Zt-Result) */
      if ErrorIn(_Zt_try11503) {Result = _Zt_try11503
      } else {
      _Zt = ANY(_Zt_try11503)
      F_Optimize_range_infers_for_Variable(self.ClaireVar,F_Optimize_pmember_type(ToType(_Zt)),ToType(_Zt))
      if (self.Other == CTRUE.Id()) /* If:3 */{ 
        /* Let:4 */{ 
          var g1151UU *Language.Call  
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
            g1151UU = _CL_obj
            /* Let-5 */} 
          Result = Core.F_CALL(C_c_code,ARGS(EID{g1151UU.Id(),0},EID{s.Id(),0}))
          /* Let-4 */} 
        /* If!3 */}  else if (self.Other == CNULL) /* If:3 */{ 
        /* Let:4 */{ 
          var v *ClaireVariable  
          /* noccur = 3 */
          /* Let:5 */{ 
            var g1152UU int 
            /* noccur = 1 */
            C_OPT.MaxVars = (C_OPT.MaxVars+1)
            g1152UU = 0
            v = F_Compile_Variable_I_symbol(F_append_symbol(self.ClaireVar.Pname,MakeString("_some").Id()),g1152UU,F_Optimize_extends_type(self.ClaireVar.Range).Id())
            /* Let-5 */} 
          /* Let:5 */{ 
            var g1153UU *Language.Let  
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
              g1153UU = _CL_obj
              /* Let-6 */} 
            Result = Core.F_CALL(C_c_code,ARGS(EID{g1153UU.Id(),0},EID{s.Id(),0}))
            /* Let-5 */} 
          /* Let-4 */} 
        } else {
        /* Let:4 */{ 
          var g1154UU *Language.Call  
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
            g1154UU = _CL_obj
            /* Let-5 */} 
          Result = Core.F_CALL(C_c_code,ARGS(EID{g1154UU.Id(),0},EID{s.Id(),0}))
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
      var _Zt_try11553 EID 
      /* Let:3 */{ 
        var g1156UU *ClaireType  
        /* noccur = 1 */
        var g1156UU_try11574 EID 
        g1156UU_try11574 = Core.F_CALL(C_c_type,ARGS(self.SetArg.ToEID()))
        /* ERROR PROTECTION INSERTED (g1156UU-_Zt_try11553) */
        if ErrorIn(g1156UU_try11574) {_Zt_try11553 = g1156UU_try11574
        } else {
        g1156UU = ToType(OBJ(g1156UU_try11574))
        _Zt_try11553 = EID{F_Optimize_ptype_type(g1156UU).Id(),0}
        }
        /* Let-3 */} 
      /* ERROR PROTECTION INSERTED (_Zt-Result) */
      if ErrorIn(_Zt_try11553) {Result = _Zt_try11553
      } else {
      _Zt = ToType(OBJ(_Zt_try11553))
      F_Optimize_range_infers_for_Variable(self.ClaireVar,F_Optimize_pmember_type(_Zt),_Zt)
      if (self.Of.Id() != CNULL) /* If:3 */{ 
        Result = EID{Core.F_param_I_class(C_set,self.Of).Id(),0}
        } else {
        /* Let:4 */{ 
          var g1158UU *ClaireType  
          /* noccur = 1 */
          var g1158UU_try11595 EID 
          g1158UU_try11595 = Core.F_CALL(C_c_type,ARGS(self.Arg.ToEID()))
          /* ERROR PROTECTION INSERTED (g1158UU-Result) */
          if ErrorIn(g1158UU_try11595) {Result = g1158UU_try11595
          } else {
          g1158UU = ToType(OBJ(g1158UU_try11595))
          Result = EID{Core.F_nth_class1(C_set,g1158UU).Id(),0}
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
      var _Zt_try11603 EID 
      _Zt_try11603 = Core.F_CALL(C_c_type,ARGS(self.SetArg.ToEID()))
      /* ERROR PROTECTION INSERTED (_Zt-Result) */
      if ErrorIn(_Zt_try11603) {Result = _Zt_try11603
      } else {
      _Zt = ANY(_Zt_try11603)
      F_Optimize_range_infers_for_Variable(self.ClaireVar,F_Optimize_pmember_type(ToType(_Zt)),ToType(_Zt))
      if (self.Of.Id() != CNULL) /* If:3 */{ 
        Result = EID{Core.F_param_I_class(C_set,self.Of).Id(),0}
        } else {
        /* Let:4 */{ 
          var g1161UU *ClaireType  
          /* noccur = 1 */
          var g1161UU_try11625 EID 
          /* Let:5 */{ 
            var g1163UU *ClaireType  
            /* noccur = 1 */
            var g1163UU_try11646 EID 
            g1163UU_try11646 = Core.F_CALL(C_c_type,ARGS(self.SetArg.ToEID()))
            /* ERROR PROTECTION INSERTED (g1163UU-g1161UU_try11625) */
            if ErrorIn(g1163UU_try11646) {g1161UU_try11625 = g1163UU_try11646
            } else {
            g1163UU = ToType(OBJ(g1163UU_try11646))
            g1161UU_try11625 = EID{F_Optimize_pmember_type(g1163UU).Id(),0}
            }
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (g1161UU-Result) */
          if ErrorIn(g1161UU_try11625) {Result = g1161UU_try11625
          } else {
          g1161UU = ToType(OBJ(g1161UU_try11625))
          Result = EID{Core.F_nth_class1(C_set,g1161UU).Id(),0}
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
      var _Zt_try11653 EID 
      _Zt_try11653 = Core.F_CALL(C_c_type,ARGS(self.SetArg.ToEID()))
      /* ERROR PROTECTION INSERTED (_Zt-Result) */
      if ErrorIn(_Zt_try11653) {Result = _Zt_try11653
      } else {
      _Zt = ANY(_Zt_try11653)
      F_Optimize_range_infers_for_Variable(self.ClaireVar,F_Optimize_pmember_type(ToType(_Zt)),ToType(_Zt))
      if (self.Of.Id() != CNULL) /* If:3 */{ 
        Result = EID{Core.F_param_I_class(C_list,self.Of).Id(),0}
        } else {
        /* Let:4 */{ 
          var g1166UU *ClaireType  
          /* noccur = 1 */
          var g1166UU_try11675 EID 
          /* Let:5 */{ 
            var g1168UU *ClaireType  
            /* noccur = 1 */
            var g1168UU_try11696 EID 
            g1168UU_try11696 = Core.F_CALL(C_c_type,ARGS(self.SetArg.ToEID()))
            /* ERROR PROTECTION INSERTED (g1168UU-g1166UU_try11675) */
            if ErrorIn(g1168UU_try11696) {g1166UU_try11675 = g1168UU_try11696
            } else {
            g1168UU = ToType(OBJ(g1168UU_try11696))
            g1166UU_try11675 = EID{F_Optimize_pmember_type(g1168UU).Id(),0}
            }
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (g1166UU-Result) */
          if ErrorIn(g1166UU_try11675) {Result = g1166UU_try11675
          } else {
          g1166UU = ToType(OBJ(g1166UU_try11675))
          Result = EID{Core.F_nth_class1(C_list,g1166UU).Id(),0}
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
      var g1170UU *ClaireType  
      /* noccur = 1 */
      var g1170UU_try11713 EID 
      g1170UU_try11713 = F_Compile_return_type_any(self.Arg)
      /* ERROR PROTECTION INSERTED (g1170UU-Result) */
      if ErrorIn(g1170UU_try11713) {Result = g1170UU_try11713
      } else {
      g1170UU = ToType(OBJ(g1170UU_try11713))
      Result = EID{F_Optimize_infers_from_type(g1170UU,self.Id()).Id(),0}
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
        var va_arg2_try11724 EID 
        va_arg2_try11724 = F_Optimize_c_boolean_any(self.Test)
        /* ERROR PROTECTION INSERTED (va_arg2-Result) */
        if ErrorIn(va_arg2_try11724) {Result = va_arg2_try11724
        } else {
        va_arg2 = ANY(va_arg2_try11724)
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
        var va_arg2_try11734 EID 
        va_arg2_try11734 = Core.F_CALL(C_c_code,ARGS(self.Arg.ToEID(),EID{C_void.Id(),0}))
        /* ERROR PROTECTION INSERTED (va_arg2-Result) */
        if ErrorIn(va_arg2_try11734) {Result = va_arg2_try11734
        } else {
        va_arg2 = ANY(va_arg2_try11734)
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
      var v_try11743 EID 
      v_try11743 = F_eval_any2(MakeInteger(x.Arg1).Id(),C_Interval)
      /* ERROR PROTECTION INSERTED (v-Result) */
      if ErrorIn(v_try11743) {Result = v_try11743
      } else {
      v = ANY(v_try11743)
      /* Let:3 */{ 
        var _Zmax int 
        /* noccur = 1 */
        var _Zmax_try11754 EID 
        _Zmax_try11754 = F_eval_any2(MakeInteger(x.Arg2).Id(),C_Interval)
        /* ERROR PROTECTION INSERTED (_Zmax-Result) */
        if ErrorIn(_Zmax_try11754) {Result = _Zmax_try11754
        } else {
        _Zmax = INT(_Zmax_try11754)
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
              var v_try11767 EID 
              v_try11767 = Core.F_nth_array(_Za,_Zi)
              /* ERROR PROTECTION INSERTED (v-void_try6) */
              if ErrorIn(v_try11767) {void_try6 = v_try11767
              } else {
              v = ANY(v_try11767)
              
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
      var v_try11773 EID 
      v_try11773 = EVAL(ToList(OBJ(Core.F_CALL(C_args,ARGS(x.ToEID())))).At(1-1))
      /* ERROR PROTECTION INSERTED (v-Result) */
      if ErrorIn(v_try11773) {Result = v_try11773
      } else {
      v = ANY(v_try11773)
      /* Let:3 */{ 
        var _Zmax *ClaireAny  
        /* noccur = 1 */
        var _Zmax_try11784 EID 
        _Zmax_try11784 = EVAL(ToList(OBJ(Core.F_CALL(C_args,ARGS(x.ToEID())))).At(2-1))
        /* ERROR PROTECTION INSERTED (_Zmax-Result) */
        if ErrorIn(_Zmax_try11784) {Result = _Zmax_try11784
        } else {
        _Zmax = ANY(_Zmax_try11784)
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
      var v_support_try11793 EID 
      /* Let:3 */{ 
        var g1180UU *ClaireAny  
        /* noccur = 1 */
        var g1180UU_try11814 EID 
        g1180UU_try11814 = EVAL(x.SetArg)
        /* ERROR PROTECTION INSERTED (g1180UU-v_support_try11793) */
        if ErrorIn(g1180UU_try11814) {v_support_try11793 = g1180UU_try11814
        } else {
        g1180UU = ANY(g1180UU_try11814)
        v_support_try11793 = Core.F_enumerate_any(g1180UU)
        }
        /* Let-3 */} 
      /* ERROR PROTECTION INSERTED (v_support-Result) */
      if ErrorIn(v_support_try11793) {Result = v_support_try11793
      } else {
      v_support = ToList(OBJ(v_support_try11793))
      v_len := v_support.Length()
      for i_it := 0; i_it < v_len; i_it++ { 
        v = v_support.At(i_it)
        var void_try4 EID 
        _ = void_try4
        var g1182I *ClaireBoolean  
        var g1182I_try11834 EID 
        /* Let:4 */{ 
          var g1184UU *ClaireAny  
          /* noccur = 1 */
          var g1184UU_try11855 EID 
          g1184UU_try11855 = EVAL(Language.F_substitution_any(x.Arg,x.ClaireVar,v))
          /* ERROR PROTECTION INSERTED (g1184UU-g1182I_try11834) */
          if ErrorIn(g1184UU_try11855) {g1182I_try11834 = g1184UU_try11855
          } else {
          g1184UU = ANY(g1184UU_try11855)
          g1182I_try11834 = EID{F_boolean_I_any(g1184UU).Id(),0}
          }
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (g1182I-void_try4) */
        if ErrorIn(g1182I_try11834) {void_try4 = g1182I_try11834
        } else {
        g1182I = ToBoolean(OBJ(g1182I_try11834))
        if (g1182I == CTRUE) /* If:4 */{ 
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
      var v_support_try11863 EID 
      /* Let:3 */{ 
        var g1187UU *ClaireAny  
        /* noccur = 1 */
        var g1187UU_try11884 EID 
        g1187UU_try11884 = EVAL(x.SetArg)
        /* ERROR PROTECTION INSERTED (g1187UU-v_support_try11863) */
        if ErrorIn(g1187UU_try11884) {v_support_try11863 = g1187UU_try11884
        } else {
        g1187UU = ANY(g1187UU_try11884)
        v_support_try11863 = Core.F_enumerate_any(g1187UU)
        }
        /* Let-3 */} 
      /* ERROR PROTECTION INSERTED (v_support-Result) */
      if ErrorIn(v_support_try11863) {Result = v_support_try11863
      } else {
      v_support = ToList(OBJ(v_support_try11863))
      v_len := v_support.Length()
      for i_it := 0; i_it < v_len; i_it++ { 
        v = v_support.At(i_it)
        var void_try4 EID 
        _ = void_try4
        var g1189I *ClaireBoolean  
        var g1189I_try11904 EID 
        /* Let:4 */{ 
          var g1191UU *ClaireAny  
          /* noccur = 1 */
          var g1191UU_try11925 EID 
          g1191UU_try11925 = EVAL(Language.F_substitution_any(x.Arg,x.ClaireVar,v))
          /* ERROR PROTECTION INSERTED (g1191UU-g1189I_try11904) */
          if ErrorIn(g1191UU_try11925) {g1189I_try11904 = g1191UU_try11925
          } else {
          g1191UU = ANY(g1191UU_try11925)
          g1189I_try11904 = EID{F_boolean_I_any(g1191UU).Id(),0}
          }
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (g1189I-void_try4) */
        if ErrorIn(g1189I_try11904) {void_try4 = g1189I_try11904
        } else {
        g1189I = ToBoolean(OBJ(g1189I_try11904))
        if (g1189I == CTRUE) /* If:4 */{ 
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
      var C_Zv_support_try11933 EID 
      /* Let:3 */{ 
        var g1194UU *ClaireAny  
        /* noccur = 1 */
        var g1194UU_try11954 EID 
        g1194UU_try11954 = EVAL(x.SetArg)
        /* ERROR PROTECTION INSERTED (g1194UU-C_Zv_support_try11933) */
        if ErrorIn(g1194UU_try11954) {C_Zv_support_try11933 = g1194UU_try11954
        } else {
        g1194UU = ANY(g1194UU_try11954)
        C_Zv_support_try11933 = Core.F_enumerate_any(g1194UU)
        }
        /* Let-3 */} 
      /* ERROR PROTECTION INSERTED (C_Zv_support-Result) */
      if ErrorIn(C_Zv_support_try11933) {Result = C_Zv_support_try11933
      } else {
      C_Zv_support = ToList(OBJ(C_Zv_support_try11933))
      C_Zv_len := C_Zv_support.Length()
      for i_it := 0; i_it < C_Zv_len; i_it++ { 
        C_Zv = C_Zv_support.At(i_it)
        var void_try4 EID 
        _ = void_try4
        /* Let:4 */{ 
          var v *ClaireAny  
          /* noccur = 0 */
          _ = v
          var v_try11965 EID 
          v_try11965 = EVAL(Language.F_substitution_any(x.Arg,x.ClaireVar,C_Zv))
          /* ERROR PROTECTION INSERTED (v-void_try4) */
          if ErrorIn(v_try11965) {void_try4 = v_try11965
          } else {
          v = ANY(v_try11965)
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
      var v_support_try11973 EID 
      /* Let:3 */{ 
        var g1198UU *ClaireAny  
        /* noccur = 1 */
        var g1198UU_try11994 EID 
        g1198UU_try11994 = EVAL(ToList(OBJ(Core.F_CALL(C_args,ARGS(x.ToEID())))).At(1-1))
        /* ERROR PROTECTION INSERTED (g1198UU-v_support_try11973) */
        if ErrorIn(g1198UU_try11994) {v_support_try11973 = g1198UU_try11994
        } else {
        g1198UU = ANY(g1198UU_try11994)
        v_support_try11973 = Core.F_enumerate_any(g1198UU)
        }
        /* Let-3 */} 
      /* ERROR PROTECTION INSERTED (v_support-Result) */
      if ErrorIn(v_support_try11973) {Result = v_support_try11973
      } else {
      v_support = ToList(OBJ(v_support_try11973))
      v_len := v_support.Length()
      for i_it := 0; i_it < v_len; i_it++ { 
        v = v_support.At(i_it)
        var void_try4 EID 
        _ = void_try4
        var g1200I *ClaireBoolean  
        var g1200I_try12014 EID 
        /* Let:4 */{ 
          var g1202UU *ClaireAny  
          /* noccur = 1 */
          var g1202UU_try12035 EID 
          g1202UU_try12035 = EVAL(ToList(OBJ(Core.F_CALL(C_args,ARGS(x.ToEID())))).At(2-1))
          /* ERROR PROTECTION INSERTED (g1202UU-g1200I_try12014) */
          if ErrorIn(g1202UU_try12035) {g1200I_try12014 = g1202UU_try12035
          } else {
          g1202UU = ANY(g1202UU_try12035)
          g1200I_try12014 = EID{Core.F__I_equal_any(v,g1202UU).Id(),0}
          }
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (g1200I-void_try4) */
        if ErrorIn(g1200I_try12014) {void_try4 = g1200I_try12014
        } else {
        g1200I = ToBoolean(OBJ(g1200I_try12014))
        if (g1200I == CTRUE) /* If:4 */{ 
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
      var v_support_try12043 EID 
      /* Let:3 */{ 
        var g1205UU *ClaireAny  
        /* noccur = 1 */
        var g1205UU_try12064 EID 
        g1205UU_try12064 = EVAL(ToList(OBJ(Core.F_CALL(C_args,ARGS(x.ToEID())))).At(1-1))
        /* ERROR PROTECTION INSERTED (g1205UU-v_support_try12043) */
        if ErrorIn(g1205UU_try12064) {v_support_try12043 = g1205UU_try12064
        } else {
        g1205UU = ANY(g1205UU_try12064)
        v_support_try12043 = Core.F_enumerate_any(g1205UU)
        }
        /* Let-3 */} 
      /* ERROR PROTECTION INSERTED (v_support-Result) */
      if ErrorIn(v_support_try12043) {Result = v_support_try12043
      } else {
      v_support = ToList(OBJ(v_support_try12043))
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
      var v_support_try12073 EID 
      /* Let:3 */{ 
        var g1208UU *ClaireAny  
        /* noccur = 1 */
        var g1208UU_try12094 EID 
        g1208UU_try12094 = EVAL(ToList(OBJ(Core.F_CALL(C_args,ARGS(x.ToEID())))).At(2-1))
        /* ERROR PROTECTION INSERTED (g1208UU-v_support_try12073) */
        if ErrorIn(g1208UU_try12094) {v_support_try12073 = g1208UU_try12094
        } else {
        g1208UU = ANY(g1208UU_try12094)
        v_support_try12073 = Core.F_enumerate_any(g1208UU)
        }
        /* Let-3 */} 
      /* ERROR PROTECTION INSERTED (v_support-Result) */
      if ErrorIn(v_support_try12073) {Result = v_support_try12073
      } else {
      v_support = ToList(OBJ(v_support_try12073))
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
  