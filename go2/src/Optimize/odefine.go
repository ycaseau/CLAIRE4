/***** CLAIRE Compilation of file /Users/ycaseau/Dropbox/src/clairev4.0/src/compile/odefine.cl 
         [version 4.0.01 / safety 5] Saturday 10-30-2021 *****/

package Optimize
import (_ "fmt"
	. "Kernel"
	"Core"
	"Language"
	"Reader"
)

//-------- dumb function to prevent import errors --------
func import_g1210() { 
    _ = Core.It
    _ = Language.It
    _ = Reader.It
    } 
  
  
//+-------------------------------------------------------------+
//| CLAIRE                                                      |
//| odefine.cl                                                  |
//| Copyright (C) 1994 - 2013 Yves Caseau. All Rights Reserved  |
//| cf. copyright info in file object.cl: about()               |
//+-------------------------------------------------------------+
// *********************************************************************
// *  Table of contents                                                *
// *     Part 1: Set, List and Tuple creation                          *
// *     Part 2: Object definition                                     *
// *     Part 3: Method instantiation                                  *
// *     Part 4: Inverse Management                                    *
// *********************************************************************
// *********************************************************************
// *     Part 1: Set, List and Tuple creation                          *
// *********************************************************************
// type inference has changed in v3.2:
/* {1} OPT.The go function for: c_type(self:List) [] */
func F_c_type_List (self *Language.List ) EID { 
    var Result EID 
    
    if (self.Of.Id() != CNULL) /* If:2 */{ 
      Result = EID{Core.F_param_I_class(C_list,self.Of).Id(),0}
      } else {
      /* Let:3 */{ 
        var _Zres *ClaireAny   = CEMPTY.Id()
        /* noccur = 5 */
        /* For:4 */{ 
          var _Zx *ClaireAny  
          _ = _Zx
          Result= EID{CFALSE.Id(),0}
          var _Zx_support *ClaireList  
          _Zx_support = self.Args
          _Zx_len := _Zx_support.Length()
          for i_it := 0; i_it < _Zx_len; i_it++ { 
            _Zx = _Zx_support.At(i_it)
            var void_try6 EID 
            _ = void_try6
            if (F_boolean_I_any(_Zres) == CTRUE) /* If:6 */{ 
              var _Zres_try12117 EID 
              /* Let:7 */{ 
                var g1212UU *ClaireClass  
                /* noccur = 1 */
                var g1212UU_try12138 EID 
                /* Let:8 */{ 
                  var g1214UU *ClaireType  
                  /* noccur = 1 */
                  var g1214UU_try12159 EID 
                  /* Let:9 */{ 
                    var g1216UU *ClaireType  
                    /* noccur = 1 */
                    var g1216UU_try121710 EID 
                    g1216UU_try121710 = Core.F_CALL(C_c_type,ARGS(_Zx.ToEID()))
                    /* ERROR PROTECTION INSERTED (g1216UU-g1214UU_try12159) */
                    if ErrorIn(g1216UU_try121710) {g1214UU_try12159 = g1216UU_try121710
                    } else {
                    g1216UU = ToType(OBJ(g1216UU_try121710))
                    g1214UU_try12159 = EID{F_Optimize_ptype_type(g1216UU).Id(),0}
                    }
                    /* Let-9 */} 
                  /* ERROR PROTECTION INSERTED (g1214UU-g1212UU_try12138) */
                  if ErrorIn(g1214UU_try12159) {g1212UU_try12138 = g1214UU_try12159
                  } else {
                  g1214UU = ToType(OBJ(g1214UU_try12159))
                  g1212UU_try12138 = EID{g1214UU.Class_I().Id(),0}
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (g1212UU-_Zres_try12117) */
                if ErrorIn(g1212UU_try12138) {_Zres_try12117 = g1212UU_try12138
                } else {
                g1212UU = ToClass(OBJ(g1212UU_try12138))
                _Zres_try12117 = EID{Core.F_meet_class(ToClass(_Zres),g1212UU).Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (_Zres-void_try6) */
              if ErrorIn(_Zres_try12117) {void_try6 = _Zres_try12117
              } else {
              _Zres = ANY(_Zres_try12117)
              void_try6 = _Zres.ToEID()
              }
              } else {
              var _Zres_try12187 EID 
              /* Let:7 */{ 
                var g1219UU *ClaireType  
                /* noccur = 1 */
                var g1219UU_try12208 EID 
                /* Let:8 */{ 
                  var g1221UU *ClaireType  
                  /* noccur = 1 */
                  var g1221UU_try12229 EID 
                  g1221UU_try12229 = Core.F_CALL(C_c_type,ARGS(_Zx.ToEID()))
                  /* ERROR PROTECTION INSERTED (g1221UU-g1219UU_try12208) */
                  if ErrorIn(g1221UU_try12229) {g1219UU_try12208 = g1221UU_try12229
                  } else {
                  g1221UU = ToType(OBJ(g1221UU_try12229))
                  g1219UU_try12208 = EID{F_Optimize_ptype_type(g1221UU).Id(),0}
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (g1219UU-_Zres_try12187) */
                if ErrorIn(g1219UU_try12208) {_Zres_try12187 = g1219UU_try12208
                } else {
                g1219UU = ToType(OBJ(g1219UU_try12208))
                _Zres_try12187 = EID{g1219UU.Class_I().Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (_Zres-void_try6) */
              if ErrorIn(_Zres_try12187) {void_try6 = _Zres_try12187
              } else {
              _Zres = ANY(_Zres_try12187)
              void_try6 = _Zres.ToEID()
              }
              /* If-6 */} 
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
        Result = EID{Core.F_nth_class1(C_list,ToType(_Zres)).Id(),0}
        }
        /* Let-3 */} 
      /* If-2 */} 
    return Result} 
  
// The EID go function for: c_type @ List (throw: true) 
func E_c_type_List (self EID) EID { 
    return /*(sm for c_type @ List= EID)*/ F_c_type_List(Language.To_List(OBJ(self)) )} 
  
// compile a List: take the of parameter into account !
/* {1} OPT.The go function for: c_code(self:List) [] */
func F_c_code_List (self *Language.List ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var x *Language.List  
      /* noccur = 5 */
      var x_try12233 EID 
      /* Let:3 */{ 
        var _CL_obj *Language.List   = Language.To_List(new(Language.List).Is(Language.C_List))
        /* noccur = 3 */
        /* update:4 */{ 
          var va_arg1 *Language.Construct  
          var va_arg2 *ClaireList  
          va_arg1 = Language.To_Construct(_CL_obj.Id())
          var va_arg2_try12245 EID 
          /* Iteration:5 */{ 
            var v_list5 *ClaireList  
            var _Zx *ClaireAny  
            var v_local5 *ClaireAny  
            v_list5 = self.Args
            va_arg2_try12245 = EID{CreateList(ToType(CEMPTY.Id()),v_list5.Length()).Id(),0}
            for CLcount := 0; CLcount < v_list5.Length(); CLcount++{ 
              _Zx = v_list5.At(CLcount)
              var v_local5_try12257 EID 
              v_local5_try12257 = Core.F_CALL(C_c_code,ARGS(_Zx.ToEID(),EID{C_any.Id(),0}))
              /* ERROR PROTECTION INSERTED (v_local5-va_arg2_try12245) */
              if ErrorIn(v_local5_try12257) {va_arg2_try12245 = v_local5_try12257
              va_arg2_try12245 = v_local5_try12257
              break
              } else {
              v_local5 = ANY(v_local5_try12257)
              ToList(OBJ(va_arg2_try12245)).PutAt(CLcount,v_local5)
              } 
            }
            /* Iteration-5 */} 
          /* ERROR PROTECTION INSERTED (va_arg2-x_try12233) */
          if ErrorIn(va_arg2_try12245) {x_try12233 = va_arg2_try12245
          } else {
          va_arg2 = ToList(OBJ(va_arg2_try12245))
          /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
          va_arg1.Args = va_arg2
          x_try12233 = EID{va_arg2.Id(),0}
          }
          /* update-4 */} 
        /* ERROR PROTECTION INSERTED (x_try12233-x_try12233) */
        if !ErrorIn(x_try12233) {
        x_try12233 = EID{_CL_obj.Id(),0}
        }
        /* Let-3 */} 
      /* ERROR PROTECTION INSERTED (x-Result) */
      if ErrorIn(x_try12233) {Result = x_try12233
      } else {
      x = Language.To_List(OBJ(x_try12233))
      if (self.Of.Id() != CNULL) /* If:3 */{ 
        var g1226I *ClaireBoolean  
        var g1226I_try12274 EID 
        /* or:4 */{ 
          var v_or4 *ClaireBoolean  
          
          v_or4 = Core.F__sup_integer(C_compiler.Safety,4)
          if (v_or4 == CTRUE) {g1226I_try12274 = EID{CTRUE.Id(),0}
          } else /* or:5 */{ 
            v_or4 = Equal(self.Of.Id(),CEMPTY.Id())
            if (v_or4 == CTRUE) {g1226I_try12274 = EID{CTRUE.Id(),0}
            } else /* or:6 */{ 
              var v_or4_try12287 EID 
              /* Let:7 */{ 
                var g1229UU *ClaireAny  
                /* noccur = 1 */
                var g1229UU_try12308 EID 
                /* For:8 */{ 
                  var _Zx *ClaireAny  
                  _ = _Zx
                  g1229UU_try12308= EID{CFALSE.Id(),0}
                  var _Zx_support *ClaireList  
                  _Zx_support = self.Args
                  _Zx_len := _Zx_support.Length()
                  for i_it := 0; i_it < _Zx_len; i_it++ { 
                    _Zx = _Zx_support.At(i_it)
                    var void_try10 EID 
                    _ = void_try10
                    var g1231I *ClaireBoolean  
                    var g1231I_try123210 EID 
                    /* Let:10 */{ 
                      var g1233UU *ClaireBoolean  
                      /* noccur = 1 */
                      var g1233UU_try123411 EID 
                      /* Let:11 */{ 
                        var g1235UU *ClaireType  
                        /* noccur = 1 */
                        var g1235UU_try123612 EID 
                        g1235UU_try123612 = Core.F_CALL(C_c_type,ARGS(_Zx.ToEID()))
                        /* ERROR PROTECTION INSERTED (g1235UU-g1233UU_try123411) */
                        if ErrorIn(g1235UU_try123612) {g1233UU_try123411 = g1235UU_try123612
                        } else {
                        g1235UU = ToType(OBJ(g1235UU_try123612))
                        g1233UU_try123411 = EID{g1235UU.Included(self.Of).Id(),0}
                        }
                        /* Let-11 */} 
                      /* ERROR PROTECTION INSERTED (g1233UU-g1231I_try123210) */
                      if ErrorIn(g1233UU_try123411) {g1231I_try123210 = g1233UU_try123411
                      } else {
                      g1233UU = ToBoolean(OBJ(g1233UU_try123411))
                      g1231I_try123210 = EID{g1233UU.Not.Id(),0}
                      }
                      /* Let-10 */} 
                    /* ERROR PROTECTION INSERTED (g1231I-void_try10) */
                    if ErrorIn(g1231I_try123210) {void_try10 = g1231I_try123210
                    } else {
                    g1231I = ToBoolean(OBJ(g1231I_try123210))
                    if (g1231I == CTRUE) /* If:10 */{ 
                       /*v = g1229UU_try12308, s =EID*/
g1229UU_try12308 = EID{CTRUE.Id(),0}
                      break
                      } else {
                      void_try10 = EID{CFALSE.Id(),0}
                      /* If-10 */} 
                    }
                    /* ERROR PROTECTION INSERTED (void_try10-g1229UU_try12308) */
                    if ErrorIn(void_try10) {g1229UU_try12308 = void_try10
                    g1229UU_try12308 = void_try10
                    break
                    } else {
                    }
                    /* loop-9 */} 
                  /* For-8 */} 
                /* ERROR PROTECTION INSERTED (g1229UU-v_or4_try12287) */
                if ErrorIn(g1229UU_try12308) {v_or4_try12287 = g1229UU_try12308
                } else {
                g1229UU = ANY(g1229UU_try12308)
                v_or4_try12287 = EID{Core.F_not_any(g1229UU).Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (v_or4-g1226I_try12274) */
              if ErrorIn(v_or4_try12287) {g1226I_try12274 = v_or4_try12287
              } else {
              v_or4 = ToBoolean(OBJ(v_or4_try12287))
              if (v_or4 == CTRUE) {g1226I_try12274 = EID{CTRUE.Id(),0}
              } else /* or:7 */{ 
                g1226I_try12274 = EID{CFALSE.Id(),0}/* org-7 */} 
              /* org-6 */} 
            /* org-5 */} 
          }
          /* or-4 */} 
        /* ERROR PROTECTION INSERTED (g1226I-Result) */
        if ErrorIn(g1226I_try12274) {Result = g1226I_try12274
        } else {
        g1226I = ToBoolean(OBJ(g1226I_try12274))
        if (g1226I == CTRUE) /* If:4 */{ 
          x.Of = self.Of
          Result = EID{x.Id(),0}
          } else {
          F_Compile_warn_void()
          /* Let:5 */{ 
            var g1237UU *ClaireList  
            /* noccur = 1 */
            var g1237UU_try12386 EID 
            /* Construct:6 */{ 
              var v_bag_arg *ClaireAny  
              g1237UU_try12386= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
              var v_bag_arg_try12397 EID 
              /* Iteration:7 */{ 
                var v_list7 *ClaireList  
                var _Zx *ClaireAny  
                var v_local7 *ClaireAny  
                v_list7 = self.Args
                v_bag_arg_try12397 = EID{CreateList(ToType(CEMPTY.Id()),v_list7.Length()).Id(),0}
                for CLcount := 0; CLcount < v_list7.Length(); CLcount++{ 
                  _Zx = v_list7.At(CLcount)
                  var v_local7_try12409 EID 
                  v_local7_try12409 = Core.F_CALL(C_c_type,ARGS(_Zx.ToEID()))
                  /* ERROR PROTECTION INSERTED (v_local7-v_bag_arg_try12397) */
                  if ErrorIn(v_local7_try12409) {v_bag_arg_try12397 = v_local7_try12409
                  v_bag_arg_try12397 = v_local7_try12409
                  break
                  } else {
                  v_local7 = ANY(v_local7_try12409)
                  ToList(OBJ(v_bag_arg_try12397)).PutAt(CLcount,v_local7)
                  } 
                }
                /* Iteration-7 */} 
              /* ERROR PROTECTION INSERTED (v_bag_arg-g1237UU_try12386) */
              if ErrorIn(v_bag_arg_try12397) {g1237UU_try12386 = v_bag_arg_try12397
              } else {
              v_bag_arg = ANY(v_bag_arg_try12397)
              ToList(OBJ(g1237UU_try12386)).AddFast(v_bag_arg)
              ToList(OBJ(g1237UU_try12386)).AddFast(self.Of.Id())}
              /* Construct-6 */} 
            /* ERROR PROTECTION INSERTED (g1237UU-Result) */
            if ErrorIn(g1237UU_try12386) {Result = g1237UU_try12386
            } else {
            g1237UU = ToList(OBJ(g1237UU_try12386))
            Result = Core.F_tformat_string(MakeString("unsafe typed list: ~S not in ~S [262]\n"),2,g1237UU)
            }
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          /* Let:5 */{ 
            var g1241UU *Language.Call  
            /* noccur = 1 */
            /* Let:6 */{ 
              var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
              /* noccur = 5 */
              _CL_obj.Selector = Core.C_check_in
              _CL_obj.Args = MakeConstantList(x.Id(),C_list.Id(),self.Of.Id())
              g1241UU = _CL_obj
              /* Let-6 */} 
            Result = Core.F_CALL(C_c_code,ARGS(EID{g1241UU.Id(),0},EID{C_list.Id(),0}))
            /* Let-5 */} 
          }
          /* If-4 */} 
        }
        } else {
        Result = EID{x.Id(),0}
        /* If-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_code @ List (throw: true) 
func E_c_code_List (self EID) EID { 
    return /*(sm for c_code @ List= EID)*/ F_c_code_List(Language.To_List(OBJ(self)) )} 
  
// new in v3.2: static list have type inference !         
/* {1} OPT.The go function for: c_type(self:Set) [] */
func F_c_type_Set (self *Language.Set ) EID { 
    var Result EID 
    
    if (self.Of.Id() != CNULL) /* If:2 */{ 
      Result = EID{Core.F_param_I_class(C_set,self.Of).Id(),0}
      } else {
      /* Let:3 */{ 
        var _Zres *ClaireAny   = CEMPTY.Id()
        /* noccur = 5 */
        /* For:4 */{ 
          var _Zx *ClaireAny  
          _ = _Zx
          Result= EID{CFALSE.Id(),0}
          var _Zx_support *ClaireList  
          _Zx_support = self.Args
          _Zx_len := _Zx_support.Length()
          for i_it := 0; i_it < _Zx_len; i_it++ { 
            _Zx = _Zx_support.At(i_it)
            var void_try6 EID 
            _ = void_try6
            if (F_boolean_I_any(_Zres) == CTRUE) /* If:6 */{ 
              var _Zres_try12427 EID 
              /* Let:7 */{ 
                var g1243UU *ClaireClass  
                /* noccur = 1 */
                var g1243UU_try12448 EID 
                /* Let:8 */{ 
                  var g1245UU *ClaireType  
                  /* noccur = 1 */
                  var g1245UU_try12469 EID 
                  g1245UU_try12469 = Core.F_CALL(C_c_type,ARGS(_Zx.ToEID()))
                  /* ERROR PROTECTION INSERTED (g1245UU-g1243UU_try12448) */
                  if ErrorIn(g1245UU_try12469) {g1243UU_try12448 = g1245UU_try12469
                  } else {
                  g1245UU = ToType(OBJ(g1245UU_try12469))
                  g1243UU_try12448 = EID{g1245UU.Class_I().Id(),0}
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (g1243UU-_Zres_try12427) */
                if ErrorIn(g1243UU_try12448) {_Zres_try12427 = g1243UU_try12448
                } else {
                g1243UU = ToClass(OBJ(g1243UU_try12448))
                _Zres_try12427 = EID{Core.F_meet_class(ToClass(_Zres),g1243UU).Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (_Zres-void_try6) */
              if ErrorIn(_Zres_try12427) {void_try6 = _Zres_try12427
              } else {
              _Zres = ANY(_Zres_try12427)
              void_try6 = _Zres.ToEID()
              }
              } else {
              var _Zres_try12477 EID 
              /* Let:7 */{ 
                var g1248UU *ClaireType  
                /* noccur = 1 */
                var g1248UU_try12498 EID 
                g1248UU_try12498 = Core.F_CALL(C_c_type,ARGS(_Zx.ToEID()))
                /* ERROR PROTECTION INSERTED (g1248UU-_Zres_try12477) */
                if ErrorIn(g1248UU_try12498) {_Zres_try12477 = g1248UU_try12498
                } else {
                g1248UU = ToType(OBJ(g1248UU_try12498))
                _Zres_try12477 = EID{g1248UU.Class_I().Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (_Zres-void_try6) */
              if ErrorIn(_Zres_try12477) {void_try6 = _Zres_try12477
              } else {
              _Zres = ANY(_Zres_try12477)
              void_try6 = _Zres.ToEID()
              }
              /* If-6 */} 
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
        Result = EID{Core.F_nth_class1(C_set,ToType(_Zres)).Id(),0}
        }
        /* Let-3 */} 
      /* If-2 */} 
    return Result} 
  
// The EID go function for: c_type @ Set (throw: true) 
func E_c_type_Set (self EID) EID { 
    return /*(sm for c_type @ Set= EID)*/ F_c_type_Set(Language.To_Set(OBJ(self)) )} 
  
/* {1} OPT.The go function for: c_code(self:Set) [] */
func F_c_code_Set (self *Language.Set ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var x *Language.Set  
      /* noccur = 5 */
      var x_try12503 EID 
      /* Let:3 */{ 
        var _CL_obj *Language.Set   = Language.To_Set(new(Language.Set).Is(Language.C_Set))
        /* noccur = 3 */
        /* update:4 */{ 
          var va_arg1 *Language.Construct  
          var va_arg2 *ClaireList  
          va_arg1 = Language.To_Construct(_CL_obj.Id())
          var va_arg2_try12515 EID 
          /* Iteration:5 */{ 
            var v_list5 *ClaireList  
            var _Zx *ClaireAny  
            var v_local5 *ClaireAny  
            v_list5 = self.Args
            va_arg2_try12515 = EID{CreateList(ToType(CEMPTY.Id()),v_list5.Length()).Id(),0}
            for CLcount := 0; CLcount < v_list5.Length(); CLcount++{ 
              _Zx = v_list5.At(CLcount)
              var v_local5_try12527 EID 
              v_local5_try12527 = Core.F_CALL(C_c_code,ARGS(_Zx.ToEID(),EID{C_any.Id(),0}))
              /* ERROR PROTECTION INSERTED (v_local5-va_arg2_try12515) */
              if ErrorIn(v_local5_try12527) {va_arg2_try12515 = v_local5_try12527
              va_arg2_try12515 = v_local5_try12527
              break
              } else {
              v_local5 = ANY(v_local5_try12527)
              ToList(OBJ(va_arg2_try12515)).PutAt(CLcount,v_local5)
              } 
            }
            /* Iteration-5 */} 
          /* ERROR PROTECTION INSERTED (va_arg2-x_try12503) */
          if ErrorIn(va_arg2_try12515) {x_try12503 = va_arg2_try12515
          } else {
          va_arg2 = ToList(OBJ(va_arg2_try12515))
          /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
          va_arg1.Args = va_arg2
          x_try12503 = EID{va_arg2.Id(),0}
          }
          /* update-4 */} 
        /* ERROR PROTECTION INSERTED (x_try12503-x_try12503) */
        if !ErrorIn(x_try12503) {
        x_try12503 = EID{_CL_obj.Id(),0}
        }
        /* Let-3 */} 
      /* ERROR PROTECTION INSERTED (x-Result) */
      if ErrorIn(x_try12503) {Result = x_try12503
      } else {
      x = Language.To_Set(OBJ(x_try12503))
      if (self.Of.Id() != CNULL) /* If:3 */{ 
        var g1253I *ClaireBoolean  
        var g1253I_try12544 EID 
        /* or:4 */{ 
          var v_or4 *ClaireBoolean  
          
          v_or4 = Core.F__sup_integer(C_compiler.Safety,4)
          if (v_or4 == CTRUE) {g1253I_try12544 = EID{CTRUE.Id(),0}
          } else /* or:5 */{ 
            v_or4 = Equal(self.Of.Id(),CEMPTY.Id())
            if (v_or4 == CTRUE) {g1253I_try12544 = EID{CTRUE.Id(),0}
            } else /* or:6 */{ 
              var v_or4_try12557 EID 
              /* Let:7 */{ 
                var g1256UU *ClaireAny  
                /* noccur = 1 */
                var g1256UU_try12578 EID 
                /* For:8 */{ 
                  var _Zx *ClaireAny  
                  _ = _Zx
                  g1256UU_try12578= EID{CFALSE.Id(),0}
                  var _Zx_support *ClaireList  
                  _Zx_support = self.Args
                  _Zx_len := _Zx_support.Length()
                  for i_it := 0; i_it < _Zx_len; i_it++ { 
                    _Zx = _Zx_support.At(i_it)
                    var void_try10 EID 
                    _ = void_try10
                    var g1258I *ClaireBoolean  
                    var g1258I_try125910 EID 
                    /* Let:10 */{ 
                      var g1260UU *ClaireBoolean  
                      /* noccur = 1 */
                      var g1260UU_try126111 EID 
                      /* Let:11 */{ 
                        var g1262UU *ClaireType  
                        /* noccur = 1 */
                        var g1262UU_try126312 EID 
                        g1262UU_try126312 = Core.F_CALL(C_c_type,ARGS(_Zx.ToEID()))
                        /* ERROR PROTECTION INSERTED (g1262UU-g1260UU_try126111) */
                        if ErrorIn(g1262UU_try126312) {g1260UU_try126111 = g1262UU_try126312
                        } else {
                        g1262UU = ToType(OBJ(g1262UU_try126312))
                        g1260UU_try126111 = EID{g1262UU.Included(self.Of).Id(),0}
                        }
                        /* Let-11 */} 
                      /* ERROR PROTECTION INSERTED (g1260UU-g1258I_try125910) */
                      if ErrorIn(g1260UU_try126111) {g1258I_try125910 = g1260UU_try126111
                      } else {
                      g1260UU = ToBoolean(OBJ(g1260UU_try126111))
                      g1258I_try125910 = EID{g1260UU.Not.Id(),0}
                      }
                      /* Let-10 */} 
                    /* ERROR PROTECTION INSERTED (g1258I-void_try10) */
                    if ErrorIn(g1258I_try125910) {void_try10 = g1258I_try125910
                    } else {
                    g1258I = ToBoolean(OBJ(g1258I_try125910))
                    if (g1258I == CTRUE) /* If:10 */{ 
                       /*v = g1256UU_try12578, s =EID*/
g1256UU_try12578 = EID{CTRUE.Id(),0}
                      break
                      } else {
                      void_try10 = EID{CFALSE.Id(),0}
                      /* If-10 */} 
                    }
                    /* ERROR PROTECTION INSERTED (void_try10-g1256UU_try12578) */
                    if ErrorIn(void_try10) {g1256UU_try12578 = void_try10
                    g1256UU_try12578 = void_try10
                    break
                    } else {
                    }
                    /* loop-9 */} 
                  /* For-8 */} 
                /* ERROR PROTECTION INSERTED (g1256UU-v_or4_try12557) */
                if ErrorIn(g1256UU_try12578) {v_or4_try12557 = g1256UU_try12578
                } else {
                g1256UU = ANY(g1256UU_try12578)
                v_or4_try12557 = EID{Core.F_not_any(g1256UU).Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (v_or4-g1253I_try12544) */
              if ErrorIn(v_or4_try12557) {g1253I_try12544 = v_or4_try12557
              } else {
              v_or4 = ToBoolean(OBJ(v_or4_try12557))
              if (v_or4 == CTRUE) {g1253I_try12544 = EID{CTRUE.Id(),0}
              } else /* or:7 */{ 
                g1253I_try12544 = EID{CFALSE.Id(),0}/* org-7 */} 
              /* org-6 */} 
            /* org-5 */} 
          }
          /* or-4 */} 
        /* ERROR PROTECTION INSERTED (g1253I-Result) */
        if ErrorIn(g1253I_try12544) {Result = g1253I_try12544
        } else {
        g1253I = ToBoolean(OBJ(g1253I_try12544))
        if (g1253I == CTRUE) /* If:4 */{ 
          x.Of = self.Of
          Result = EID{x.Id(),0}
          } else {
          F_Compile_warn_void()
          /* Let:5 */{ 
            var g1264UU *ClaireList  
            /* noccur = 1 */
            var g1264UU_try12656 EID 
            /* Construct:6 */{ 
              var v_bag_arg *ClaireAny  
              g1264UU_try12656= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
              var v_bag_arg_try12667 EID 
              /* Iteration:7 */{ 
                var v_list7 *ClaireList  
                var _Zx *ClaireAny  
                var v_local7 *ClaireAny  
                v_list7 = self.Args
                v_bag_arg_try12667 = EID{CreateList(ToType(CEMPTY.Id()),v_list7.Length()).Id(),0}
                for CLcount := 0; CLcount < v_list7.Length(); CLcount++{ 
                  _Zx = v_list7.At(CLcount)
                  var v_local7_try12679 EID 
                  v_local7_try12679 = Core.F_CALL(C_c_type,ARGS(_Zx.ToEID()))
                  /* ERROR PROTECTION INSERTED (v_local7-v_bag_arg_try12667) */
                  if ErrorIn(v_local7_try12679) {v_bag_arg_try12667 = v_local7_try12679
                  v_bag_arg_try12667 = v_local7_try12679
                  break
                  } else {
                  v_local7 = ANY(v_local7_try12679)
                  ToList(OBJ(v_bag_arg_try12667)).PutAt(CLcount,v_local7)
                  } 
                }
                /* Iteration-7 */} 
              /* ERROR PROTECTION INSERTED (v_bag_arg-g1264UU_try12656) */
              if ErrorIn(v_bag_arg_try12667) {g1264UU_try12656 = v_bag_arg_try12667
              } else {
              v_bag_arg = ANY(v_bag_arg_try12667)
              ToList(OBJ(g1264UU_try12656)).AddFast(v_bag_arg)
              ToList(OBJ(g1264UU_try12656)).AddFast(self.Of.Id())}
              /* Construct-6 */} 
            /* ERROR PROTECTION INSERTED (g1264UU-Result) */
            if ErrorIn(g1264UU_try12656) {Result = g1264UU_try12656
            } else {
            g1264UU = ToList(OBJ(g1264UU_try12656))
            Result = Core.F_tformat_string(MakeString("unsafe typed set: ~S not in ~S [262]\n"),2,g1264UU)
            }
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          /* Let:5 */{ 
            var g1268UU *Language.Call  
            /* noccur = 1 */
            /* Let:6 */{ 
              var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
              /* noccur = 5 */
              _CL_obj.Selector = Core.C_check_in
              _CL_obj.Args = MakeConstantList(x.Id(),C_set.Id(),self.Of.Id())
              g1268UU = _CL_obj
              /* Let-6 */} 
            Result = Core.F_CALL(C_c_code,ARGS(EID{g1268UU.Id(),0},EID{C_set.Id(),0}))
            /* Let-5 */} 
          }
          /* If-4 */} 
        }
        } else {
        Result = EID{x.Id(),0}
        /* If-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_code @ Set (throw: true) 
func E_c_code_Set (self EID) EID { 
    return /*(sm for c_code @ Set= EID)*/ F_c_code_Set(Language.To_Set(OBJ(self)) )} 
  
/* {1} OPT.The go function for: c_type(self:Tuple) [] */
func F_c_type_Tuple (self *Language.Tuple ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var g1269UU *ClaireList  
      /* noccur = 1 */
      var g1269UU_try12703 EID 
      /* Iteration:3 */{ 
        var v_list3 *ClaireList  
        var x *ClaireAny  
        var v_local3 *ClaireAny  
        v_list3 = self.Args
        g1269UU_try12703 = EID{CreateList(ToType(CEMPTY.Id()),v_list3.Length()).Id(),0}
        for CLcount := 0; CLcount < v_list3.Length(); CLcount++{ 
          x = v_list3.At(CLcount)
          var v_local3_try12715 EID 
          v_local3_try12715 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
          /* ERROR PROTECTION INSERTED (v_local3-g1269UU_try12703) */
          if ErrorIn(v_local3_try12715) {g1269UU_try12703 = v_local3_try12715
          g1269UU_try12703 = v_local3_try12715
          break
          } else {
          v_local3 = ANY(v_local3_try12715)
          ToList(OBJ(g1269UU_try12703)).PutAt(CLcount,v_local3)
          } 
        }
        /* Iteration-3 */} 
      /* ERROR PROTECTION INSERTED (g1269UU-Result) */
      if ErrorIn(g1269UU_try12703) {Result = g1269UU_try12703
      } else {
      g1269UU = ToList(OBJ(g1269UU_try12703))
      Result = EID{g1269UU.Tuple_I().Id(),0}
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_type @ Tuple (throw: true) 
func E_c_type_Tuple (self EID) EID { 
    return /*(sm for c_type @ Tuple= EID)*/ F_c_type_Tuple(Language.To_Tuple(OBJ(self)) )} 
  
/* {1} OPT.The go function for: c_code(self:Tuple) [] */
func F_c_code_Tuple (self *Language.Tuple ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var _CL_obj *Language.Tuple   = Language.To_Tuple(new(Language.Tuple).Is(Language.C_Tuple))
      /* noccur = 3 */
      /* update:3 */{ 
        var va_arg1 *Language.Construct  
        var va_arg2 *ClaireList  
        va_arg1 = Language.To_Construct(_CL_obj.Id())
        var va_arg2_try12724 EID 
        /* Iteration:4 */{ 
          var v_list4 *ClaireList  
          var _Zx *ClaireAny  
          var v_local4 *ClaireAny  
          v_list4 = self.Args
          va_arg2_try12724 = EID{CreateList(ToType(CEMPTY.Id()),v_list4.Length()).Id(),0}
          for CLcount := 0; CLcount < v_list4.Length(); CLcount++{ 
            _Zx = v_list4.At(CLcount)
            var v_local4_try12736 EID 
            v_local4_try12736 = Core.F_CALL(C_c_code,ARGS(_Zx.ToEID(),EID{C_any.Id(),0}))
            /* ERROR PROTECTION INSERTED (v_local4-va_arg2_try12724) */
            if ErrorIn(v_local4_try12736) {va_arg2_try12724 = v_local4_try12736
            va_arg2_try12724 = v_local4_try12736
            break
            } else {
            v_local4 = ANY(v_local4_try12736)
            ToList(OBJ(va_arg2_try12724)).PutAt(CLcount,v_local4)
            } 
          }
          /* Iteration-4 */} 
        /* ERROR PROTECTION INSERTED (va_arg2-Result) */
        if ErrorIn(va_arg2_try12724) {Result = va_arg2_try12724
        } else {
        va_arg2 = ToList(OBJ(va_arg2_try12724))
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
  
// The EID go function for: c_code @ Tuple (throw: true) 
func E_c_code_Tuple (self EID) EID { 
    return /*(sm for c_code @ Tuple= EID)*/ F_c_code_Tuple(Language.To_Tuple(OBJ(self)) )} 
  
// ******************************************************************
// *      Part 2: Compiling Definitions                             *
// ******************************************************************
/* {1} OPT.The go function for: c_type(self:Definition) [] */
func F_c_type_Definition (self *Language.Definition ) *ClaireType  { 
    // use function body compiling 
if (ToType(self.Arg.Id()).Included(ToType(C_exception.Id())) == CTRUE) /* body If:2 */{ 
      return  ToType(CEMPTY.Id())
      } else {
      return  ToType(self.Arg.Id())
      /* body If-2 */} 
    } 
  
// The EID go function for: c_type @ Definition (throw: false) 
func E_c_type_Definition (self EID) EID { 
    return EID{/*(sm for c_type @ Definition= type)*/ F_c_type_Definition(Language.To_Definition(OBJ(self)) ).Id(),0}} 
  
// */
// creation of a new object
/* {1} OPT.The go function for: c_code(self:Definition,s:class) [] */
func F_c_code_Definition (self *Language.Definition ,s *ClaireClass ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var _Zc *ClaireClass   = self.Arg
      /* noccur = 7 */
      /* Let:3 */{ 
        var _Zv *ClaireVariable  
        /* noccur = 2 */
        /* Let:4 */{ 
          var g1276UU int 
          /* noccur = 1 */
          C_OPT.MaxVars = (C_OPT.MaxVars+1)
          g1276UU = 0
          _Zv = F_Compile_Variable_I_symbol(ToSymbol(C_Compile__starname_star.Value),g1276UU,_Zc.Id())
          /* Let-4 */} 
        /* Let:4 */{ 
          var _Zx *ClaireAny  
          /* noccur = 2 */
          var _Zx_try12775 EID 
          _Zx_try12775 = F_Optimize_total_ask_class(_Zc,self.Args)
          /* ERROR PROTECTION INSERTED (_Zx-Result) */
          if ErrorIn(_Zx_try12775) {Result = _Zx_try12775
          } else {
          _Zx = ANY(_Zx_try12775)
          if (_Zc.Open <= 0) /* If:5 */{ 
            Result = ToException(Core.C_general_error.Make(MakeString("[105] cannot instantiate ~S").Id(),MakeConstantList(_Zc.Id()).Id())).Close()
            } else {
            Result = EID{CFALSE.Id(),0}
            /* If-5 */} 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          if (F_boolean_I_any(_Zx) == CTRUE) /* If:5 */{ 
            Result = Core.F_CALL(C_c_code,ARGS(_Zx.ToEID(),EID{s.Id(),0}))
            } else {
            /* Let:6 */{ 
              var g1278UU *Language.Let  
              /* noccur = 1 */
              var g1278UU_try12797 EID 
              /* Let:7 */{ 
                var _CL_obj *Language.Let   = Language.To_Let(new(Language.Let).Is(Language.C_Let))
                /* noccur = 21 */
                _CL_obj.ClaireVar = _Zv
                /* update:8 */{ 
                  var va_arg1 *Language.Let  
                  var va_arg2 *ClaireAny  
                  va_arg1 = _CL_obj
                  var va_arg2_try12809 EID 
                  /* Let:9 */{ 
                    var _CL_obj *Language.Cast   = Language.To_Cast(new(Language.Cast).Is(Language.C_Cast))
                    /* noccur = 10 */
                    /* update:10 */{ 
                      var va_arg1 *Language.Cast  
                      var va_arg2 *ClaireAny  
                      va_arg1 = _CL_obj
                      var va_arg2_try128111 EID 
                      /* Let:11 */{ 
                        var g1282UU *Language.Call  
                        /* noccur = 1 */
                        /* Let:12 */{ 
                          var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                          /* noccur = 5 */
                          _CL_obj.Selector = C_mClaire_new_I
                          _CL_obj.Args = MakeConstantList(_Zc.Id())
                          g1282UU = _CL_obj
                          /* Let-12 */} 
                        va_arg2_try128111 = Core.F_CALL(C_c_code,ARGS(EID{g1282UU.Id(),0},EID{C_object.Id(),0}))
                        /* Let-11 */} 
                      /* ERROR PROTECTION INSERTED (va_arg2-va_arg2_try12809) */
                      if ErrorIn(va_arg2_try128111) {va_arg2_try12809 = va_arg2_try128111
                      } else {
                      va_arg2 = ANY(va_arg2_try128111)
                      /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                      va_arg1.Arg = va_arg2
                      va_arg2_try12809 = va_arg2.ToEID()
                      }
                      /* update-10 */} 
                    /* ERROR PROTECTION INSERTED (va_arg2_try12809-va_arg2_try12809) */
                    if !ErrorIn(va_arg2_try12809) {
                    _CL_obj.SetArg = ToType(_Zc.Id())
                    va_arg2_try12809 = EID{_CL_obj.Id(),0}
                    }
                    /* Let-9 */} 
                  /* ERROR PROTECTION INSERTED (va_arg2-g1278UU_try12797) */
                  if ErrorIn(va_arg2_try12809) {g1278UU_try12797 = va_arg2_try12809
                  } else {
                  va_arg2 = ANY(va_arg2_try12809)
                  /* ---------- now we compile update value(va_arg1) := va_arg2 ------- */
                  va_arg1.Value = va_arg2
                  g1278UU_try12797 = va_arg2.ToEID()
                  }
                  /* update-8 */} 
                /* ERROR PROTECTION INSERTED (g1278UU_try12797-g1278UU_try12797) */
                if !ErrorIn(g1278UU_try12797) {
                /* update:8 */{ 
                  var va_arg1 *Language.Let  
                  var va_arg2 *ClaireAny  
                  va_arg1 = _CL_obj
                  var va_arg2_try12839 EID 
                  /* Let:9 */{ 
                    var _CL_obj *Language.Do   = Language.To_Do(new(Language.Do).Is(Language.C_Do))
                    /* noccur = 3 */
                    /* update:10 */{ 
                      var va_arg1 *Language.Do  
                      var va_arg2 *ClaireList  
                      va_arg1 = _CL_obj
                      var va_arg2_try128411 EID 
                      va_arg2_try128411 = F_Optimize_analyze_I_class(_Zc,_Zv.Id(),self.Args,ToType(CEMPTY.Id()).EmptyList())
                      /* ERROR PROTECTION INSERTED (va_arg2-va_arg2_try12839) */
                      if ErrorIn(va_arg2_try128411) {va_arg2_try12839 = va_arg2_try128411
                      } else {
                      va_arg2 = ToList(OBJ(va_arg2_try128411))
                      /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                      va_arg1.Args = va_arg2
                      va_arg2_try12839 = EID{va_arg2.Id(),0}
                      }
                      /* update-10 */} 
                    /* ERROR PROTECTION INSERTED (va_arg2_try12839-va_arg2_try12839) */
                    if !ErrorIn(va_arg2_try12839) {
                    va_arg2_try12839 = EID{_CL_obj.Id(),0}
                    }
                    /* Let-9 */} 
                  /* ERROR PROTECTION INSERTED (va_arg2-g1278UU_try12797) */
                  if ErrorIn(va_arg2_try12839) {g1278UU_try12797 = va_arg2_try12839
                  } else {
                  va_arg2 = ANY(va_arg2_try12839)
                  /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                  va_arg1.Arg = va_arg2
                  g1278UU_try12797 = va_arg2.ToEID()
                  }
                  /* update-8 */} 
                /* ERROR PROTECTION INSERTED (g1278UU_try12797-g1278UU_try12797) */
                if !ErrorIn(g1278UU_try12797) {
                g1278UU_try12797 = EID{_CL_obj.Id(),0}
                }}
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (g1278UU-Result) */
              if ErrorIn(g1278UU_try12797) {Result = g1278UU_try12797
              } else {
              g1278UU = Language.To_Let(OBJ(g1278UU_try12797))
              Result = Core.F_CALL(C_c_code,ARGS(EID{g1278UU.Id(),0},EID{s.Id(),0}))
              }
              /* Let-6 */} 
            /* If-5 */} 
          }
          }
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_code @ Definition (throw: true) 
func E_c_code_Definition (self EID,s EID) EID { 
    return /*(sm for c_code @ Definition= EID)*/ F_c_code_Definition(Language.To_Definition(OBJ(self)),ToClass(OBJ(s)) )} 
  
// tells if a "total instantiation" is appropriate (for exceptions)
// we actually check that the srange is OID or integer for all slots
/* {1} OPT.The go function for: total?(self:class,l:list) [] */
func F_Optimize_total_ask_class (self *ClaireClass ,l *ClaireList ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var lp *ClaireList  
      /* noccur = 3 */
      var lp_try12863 EID 
      lp_try12863 = Core.F_CALL(C_Compile_get_indexed,ARGS(EID{self.Id(),0}))
      /* ERROR PROTECTION INSERTED (lp-Result) */
      if ErrorIn(lp_try12863) {Result = lp_try12863
      } else {
      lp = ToList(OBJ(lp_try12863))
      /* Let:3 */{ 
        var n int  = lp.Length()
        /* noccur = 3 */
        var g1287I *ClaireBoolean  
        /* and:4 */{ 
          var v_and4 *ClaireBoolean  
          
          v_and4 = C_compiler.Diet_ask.Not
          if (v_and4 == CFALSE) {g1287I = CFALSE
          } else /* arg:5 */{ 
            v_and4 = Equal(MakeInteger(l.Length()).Id(),MakeInteger((n-1)).Id())
            if (v_and4 == CFALSE) {g1287I = CFALSE
            } else /* arg:6 */{ 
              v_and4 = MakeBoolean((self.Open == ClEnv.Default) || (ToType(self.Id()).Included(ToType(C_exception.Id())) == CTRUE))
              if (v_and4 == CFALSE) {g1287I = CFALSE
              } else /* arg:7 */{ 
                v_and4 = Core.F__inf_equal_integer(n,4)
                if (v_and4 == CFALSE) {g1287I = CFALSE
                } else /* arg:8 */{ 
                  /* Let:9 */{ 
                    var g1288UU *ClaireAny  
                    /* noccur = 1 */
                    /* Let:10 */{ 
                      var i int  = 2
                      /* noccur = 5 */
                      /* Let:11 */{ 
                        var g1285 int  = n
                        /* noccur = 1 */
                        g1288UU= CFALSE.Id()
                        for (i <= g1285) /* while:12 */{ 
                          if ((ANY(Core.F_CALL(C_mClaire_srange,ARGS(lp.At(i-1).ToEID()))) != C_any.Id()) && 
                              (ANY(Core.F_CALL(C_mClaire_srange,ARGS(lp.At(i-1).ToEID()))) != C_integer.Id())) /* If:13 */{ 
                             /*v = g1288UU, s =any*/
g1288UU = CTRUE.Id()
                            break
                            /* If-13 */} 
                          i = (i+1)
                          /* while-12 */} 
                        /* Let-11 */} 
                      /* Let-10 */} 
                    v_and4 = Core.F_not_any(g1288UU)
                    /* Let-9 */} 
                  if (v_and4 == CFALSE) {g1287I = CFALSE
                  } else /* arg:9 */{ 
                    g1287I = CTRUE/* arg-9 */} 
                  /* arg-8 */} 
                /* arg-7 */} 
              /* arg-6 */} 
            /* arg-5 */} 
          /* and-4 */} 
        if (g1287I == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var _Zc *ClaireAny  
            /* noccur = 4 */
            var _Zc_try12896 EID 
            /* Let:6 */{ 
              var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
              /* noccur = 5 */
              _CL_obj.Selector = ToProperty(IfThenElse((l.Length() == 0),
                C_mClaire_new_I.Id(),
                C_Compile_anyObject_I.Id()))
              /* update:7 */{ 
                var va_arg1 *Language.Call  
                var va_arg2 *ClaireList  
                va_arg1 = _CL_obj
                var va_arg2_try12908 EID 
                /* Let:8 */{ 
                  var g1291UU *ClaireList  
                  /* noccur = 1 */
                  var g1291UU_try12929 EID 
                  /* Iteration:9 */{ 
                    var v_list9 *ClaireList  
                    var x *ClaireAny  
                    var v_local9 *ClaireAny  
                    v_list9 = l
                    g1291UU_try12929 = EID{CreateList(ToType(CEMPTY.Id()),v_list9.Length()).Id(),0}
                    for CLcount := 0; CLcount < v_list9.Length(); CLcount++{ 
                      x = v_list9.At(CLcount)
                      var v_local9_try129311 EID 
                      v_local9_try129311 = Core.F_CALL(C_c_code,ARGS(ToList(OBJ(Core.F_CALL(C_args,ARGS(x.ToEID())))).At(2-1).ToEID(),EID{C_any.Id(),0}))
                      /* ERROR PROTECTION INSERTED (v_local9-g1291UU_try12929) */
                      if ErrorIn(v_local9_try129311) {g1291UU_try12929 = v_local9_try129311
                      g1291UU_try12929 = v_local9_try129311
                      break
                      } else {
                      v_local9 = ANY(v_local9_try129311)
                      ToList(OBJ(g1291UU_try12929)).PutAt(CLcount,v_local9)
                      } 
                    }
                    /* Iteration-9 */} 
                  /* ERROR PROTECTION INSERTED (g1291UU-va_arg2_try12908) */
                  if ErrorIn(g1291UU_try12929) {va_arg2_try12908 = g1291UU_try12929
                  } else {
                  g1291UU = ToList(OBJ(g1291UU_try12929))
                  va_arg2_try12908 = EID{F_cons_any(self.Id(),g1291UU).Id(),0}
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (va_arg2-_Zc_try12896) */
                if ErrorIn(va_arg2_try12908) {_Zc_try12896 = va_arg2_try12908
                } else {
                va_arg2 = ToList(OBJ(va_arg2_try12908))
                /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                va_arg1.Args = va_arg2
                _Zc_try12896 = EID{va_arg2.Id(),0}
                }
                /* update-7 */} 
              /* ERROR PROTECTION INSERTED (_Zc_try12896-_Zc_try12896) */
              if !ErrorIn(_Zc_try12896) {
              _Zc_try12896 = EID{_CL_obj.Id(),0}
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (_Zc-Result) */
            if ErrorIn(_Zc_try12896) {Result = _Zc_try12896
            } else {
            _Zc = ANY(_Zc_try12896)
            /* Let:6 */{ 
              var m *ClaireAny   = Core.F__at_property1(C_close,self).Id()
              /* noccur = 2 */
              if (l.Length() == 0) /* If:7 */{ 
                var _Zc_try12948 EID 
                _Zc_try12948 = Core.F_CALL(C_c_code,ARGS(_Zc.ToEID()))
                /* ERROR PROTECTION INSERTED (_Zc-Result) */
                if ErrorIn(_Zc_try12948) {Result = _Zc_try12948
                } else {
                _Zc = ANY(_Zc_try12948)
                Result = _Zc.ToEID()
                }
                } else {
                Result = EID{CFALSE.Id(),0}
                /* If-7 */} 
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              if (F_boolean_I_any(m) == CTRUE) /* If:7 */{ 
                /* Let:8 */{ 
                  var _CL_obj *Language.CallMethod1   = Language.To_CallMethod1(new(Language.CallMethod1).Is(Language.C_Call_method1))
                  /* noccur = 5 */
                  _CL_obj.Arg = ToMethod(m)
                  _CL_obj.Args = MakeConstantList(_Zc)
                  Result = EID{_CL_obj.Id(),0}
                  /* Let-8 */} 
                } else {
                Result = _Zc.ToEID()
                /* If-7 */} 
              }
              /* Let-6 */} 
            }
            /* Let-5 */} 
          } else {
          Result = EID{CFALSE.Id(),0}
          /* If-4 */} 
        /* Let-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: total? @ class (throw: true) 
func E_Optimize_total_ask_class (self EID,l EID) EID { 
    return /*(sm for total? @ class= EID)*/ F_Optimize_total_ask_class(ToClass(OBJ(self)),ToList(OBJ(l)) )} 
  
// the instantiation body is a sequence of words from which the initialization
// of the object must be built. This method produces a list of CLAIRE instructions
// self is the object (if named) or a variable if unamed
// lp will become the list of properties with explicit value setup 
// in CLAIRE 4, we assume that instantiation will put all the default values
// so we need to add write(p,self,def) for all default of p not in lp with complex (inverse or rules) management 
/* {1} OPT.The go function for: analyze!(c:class,self:any,%l:list,lp:list) [] */
func F_Optimize_analyze_I_class (c *ClaireClass ,self *ClaireAny ,_Zl *ClaireList ,lp *ClaireList ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var ins_ask *ClaireBoolean   = MakeBoolean(((c.Open == 3) || 
          (c.Open == 1)) && (F_boolean_I_any(lp.Id()).Id() != CTRUE.Id()))
      /* noccur = 0 */
      _ = ins_ask
      /* Let:3 */{ 
        var r *ClaireList  
        /* noccur = 5 */
        var r_try12954 EID 
        /* Iteration:4 */{ 
          var v_list4 *ClaireList  
          var x *ClaireAny  
          var v_local4 *ClaireAny  
          v_list4 = _Zl
          r_try12954 = EID{CreateList(ToType(C_any.Id()),v_list4.Length()).Id(),0}
          for CLcount := 0; CLcount < v_list4.Length(); CLcount++{ 
            x = v_list4.At(CLcount)
            var v_local4_try12966 EID 
            /* Let:6 */{ 
              var p *ClaireAny   = Language.To_Call(x).Args.At(1-1)
              /* noccur = 4 */
              /* Let:7 */{ 
                var y *ClaireAny   = Language.To_Call(x).Args.At(2-1)
                /* noccur = 3 */
                /* Let:8 */{ 
                  var s *ClaireObject   = Core.F__at_property1(ToProperty(p),c)
                  /* noccur = 2 */
                  /* Let:9 */{ 
                    var special_ask *ClaireBoolean   = MakeBoolean((Equal(ANY(Core.F_CALL(C_open,ARGS(p.ToEID()))),MakeInteger(0).Id()) == CTRUE) && (C_slot.Id() == s.Isa.Id()))
                    /* noccur = 2 */
                    lp = lp.AddFast(p)
                    /* Let:10 */{ 
                      var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                      /* noccur = 5 */
                      _CL_obj.Selector = ToProperty(IfThenElse((special_ask == CTRUE),
                        C_put.Id(),
                        Core.C_write.Id()))
                      /* update:11 */{ 
                        var va_arg1 *Language.Call  
                        var va_arg2 *ClaireList  
                        va_arg1 = _CL_obj
                        var va_arg2_try129712 EID 
                        /* Construct:12 */{ 
                          var v_bag_arg *ClaireAny  
                          va_arg2_try129712= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                          ToList(OBJ(va_arg2_try129712)).AddFast(p)
                          ToList(OBJ(va_arg2_try129712)).AddFast(self)
                          var v_bag_arg_try129813 EID 
                          var g1299I *ClaireBoolean  
                          var g1299I_try130013 EID 
                          /* or:13 */{ 
                            var v_or13 *ClaireBoolean  
                            
                            v_or13 = special_ask.Not
                            if (v_or13 == CTRUE) {g1299I_try130013 = EID{CTRUE.Id(),0}
                            } else /* or:14 */{ 
                              var v_or13_try130115 EID 
                              /* Let:15 */{ 
                                var g1302UU *ClaireType  
                                /* noccur = 1 */
                                var g1302UU_try130316 EID 
                                g1302UU_try130316 = Core.F_CALL(C_c_type,ARGS(y.ToEID()))
                                /* ERROR PROTECTION INSERTED (g1302UU-v_or13_try130115) */
                                if ErrorIn(g1302UU_try130316) {v_or13_try130115 = g1302UU_try130316
                                } else {
                                g1302UU = ToType(OBJ(g1302UU_try130316))
                                v_or13_try130115 = EID{g1302UU.Included(ToType(OBJ(Core.F_CALL(C_range,ARGS(EID{s.Id(),0}))))).Id(),0}
                                }
                                /* Let-15 */} 
                              /* ERROR PROTECTION INSERTED (v_or13-g1299I_try130013) */
                              if ErrorIn(v_or13_try130115) {g1299I_try130013 = v_or13_try130115
                              } else {
                              v_or13 = ToBoolean(OBJ(v_or13_try130115))
                              if (v_or13 == CTRUE) {g1299I_try130013 = EID{CTRUE.Id(),0}
                              } else /* or:15 */{ 
                                g1299I_try130013 = EID{CFALSE.Id(),0}/* org-15 */} 
                              /* org-14 */} 
                            }
                            /* or-13 */} 
                          /* ERROR PROTECTION INSERTED (g1299I-v_bag_arg_try129813) */
                          if ErrorIn(g1299I_try130013) {v_bag_arg_try129813 = g1299I_try130013
                          } else {
                          g1299I = ToBoolean(OBJ(g1299I_try130013))
                          if (g1299I == CTRUE) /* If:13 */{ 
                            v_bag_arg_try129813 = y.ToEID()
                            } else {
                            v_bag_arg_try129813 = Core.F_CALL(C_c_code,ARGS(y.ToEID(),EID{C_any.Id(),0}))
                            /* If-13 */} 
                          }
                          /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try129712) */
                          if ErrorIn(v_bag_arg_try129813) {va_arg2_try129712 = v_bag_arg_try129813
                          } else {
                          v_bag_arg = ANY(v_bag_arg_try129813)
                          ToList(OBJ(va_arg2_try129712)).AddFast(v_bag_arg)}
                          /* Construct-12 */} 
                        /* ERROR PROTECTION INSERTED (va_arg2-v_local4_try12966) */
                        if ErrorIn(va_arg2_try129712) {v_local4_try12966 = va_arg2_try129712
                        } else {
                        va_arg2 = ToList(OBJ(va_arg2_try129712))
                        /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                        va_arg1.Args = va_arg2
                        v_local4_try12966 = EID{va_arg2.Id(),0}
                        }
                        /* update-11 */} 
                      /* ERROR PROTECTION INSERTED (v_local4_try12966-v_local4_try12966) */
                      if ErrorIn(v_local4_try12966) {r_try12954 = v_local4_try12966
                      break
                      } else {
                      v_local4_try12966 = EID{_CL_obj.Id(),0}
                      }
                      /* Let-10 */} 
                    /* ERROR PROTECTION INSERTED (v_local4_try12966-v_local4_try12966) */
                    if ErrorIn(v_local4_try12966) {r_try12954 = v_local4_try12966
                    break
                    } else {
                    }
                    /* Let-9 */} 
                  /* Let-8 */} 
                /* Let-7 */} 
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (v_local4-r_try12954) */
            if ErrorIn(v_local4_try12966) {r_try12954 = v_local4_try12966
            r_try12954 = v_local4_try12966
            break
            } else {
            v_local4 = ANY(v_local4_try12966)
            ToList(OBJ(r_try12954)).PutAt(CLcount,v_local4)
            } 
          }
          /* Iteration-4 */} 
        /* ERROR PROTECTION INSERTED (r-Result) */
        if ErrorIn(r_try12954) {Result = r_try12954
        } else {
        r = ToList(OBJ(r_try12954))
        /* For:4 */{ 
          var s *ClaireAny  
          _ = s
          Result= EID{CFALSE.Id(),0}
          var s_support *ClaireList  
          var s_support_try13045 EID 
          s_support_try13045 = Core.F_CALL(C_Compile_get_indexed,ARGS(EID{c.Id(),0}))
          /* ERROR PROTECTION INSERTED (s_support-Result) */
          if ErrorIn(s_support_try13045) {Result = s_support_try13045
          } else {
          s_support = ToList(OBJ(s_support_try13045))
          s_len := s_support.Length()
          for i_it := 0; i_it < s_len; i_it++ { 
            s = s_support.At(i_it)
            var void_try6 EID 
            _ = void_try6
            /* Let:6 */{ 
              var p *ClaireProperty   = ToRestriction(s).Selector
              /* noccur = 5 */
              /* Let:7 */{ 
                var v *ClaireAny   = ToSlot(s).Default
                /* noccur = 3 */
                if ((v != CNULL) && 
                    ((lp.Memq(p.Id()) != CTRUE) && 
                      ((p.Inverse.Id() != CNULL) || 
                          (p.IfWrite != CNULL)))) /* If:8 */{ 
                  /* Let:9 */{ 
                    var defExp *ClaireAny  
                    /* noccur = 1 */
                    var defExp_try130510 EID 
                    var g1306I *ClaireBoolean  
                    var g1306I_try130710 EID 
                    g1306I_try130710 = F_Compile_designated_ask_any(v)
                    /* ERROR PROTECTION INSERTED (g1306I-defExp_try130510) */
                    if ErrorIn(g1306I_try130710) {defExp_try130510 = g1306I_try130710
                    } else {
                    g1306I = ToBoolean(OBJ(g1306I_try130710))
                    if (g1306I == CTRUE) /* If:10 */{ 
                      defExp_try130510 = v.ToEID()
                      } else {
                      /* Let:11 */{ 
                        var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                        /* noccur = 16 */
                        _CL_obj.Selector = C_default
                        /* update:12 */{ 
                          var va_arg1 *Language.Call  
                          var va_arg2 *ClaireList  
                          va_arg1 = _CL_obj
                          /* Construct:13 */{ 
                            var v_bag_arg *ClaireAny  
                            va_arg2= ToType(CEMPTY.Id()).EmptyList()
                            /* Let:14 */{ 
                              var _CL_obj *Language.Cast   = Language.To_Cast(new(Language.Cast).Is(Language.C_Cast))
                              /* noccur = 10 */
                              /* update:15 */{ 
                                var va_arg1 *Language.Cast  
                                var va_arg2 *ClaireAny  
                                va_arg1 = _CL_obj
                                /* Let:16 */{ 
                                  var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                  /* noccur = 5 */
                                  _CL_obj.Selector = ToProperty(Core.C__at.Id())
                                  _CL_obj.Args = MakeConstantList(p.Id(),c.Id())
                                  va_arg2 = _CL_obj.Id()
                                  /* Let-16 */} 
                                /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                                va_arg1.Arg = va_arg2
                                /* update-15 */} 
                              _CL_obj.SetArg = ToType(C_slot.Id())
                              v_bag_arg = _CL_obj.Id()
                              /* Let-14 */} 
                            va_arg2.AddFast(v_bag_arg)/* Construct-13 */} 
                          /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                          va_arg1.Args = va_arg2
                          /* update-12 */} 
                        defExp_try130510 = EID{_CL_obj.Id(),0}
                        /* Let-11 */} 
                      /* If-10 */} 
                    }
                    /* ERROR PROTECTION INSERTED (defExp-void_try6) */
                    if ErrorIn(defExp_try130510) {void_try6 = defExp_try130510
                    } else {
                    defExp = ANY(defExp_try130510)
                    /* Let:10 */{ 
                      var g1308UU *Language.Call  
                      /* noccur = 1 */
                      /* Let:11 */{ 
                        var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                        /* noccur = 5 */
                        _CL_obj.Selector = Core.C_write
                        _CL_obj.Args = MakeConstantList(p.Id(),self,defExp)
                        g1308UU = _CL_obj
                        /* Let-11 */} 
                      r = r.AddFast(g1308UU.Id())
                      /* Let-10 */} 
                    void_try6 = EID{r.Id(),0}
                    }
                    /* Let-9 */} 
                  } else {
                  void_try6 = EID{CFALSE.Id(),0}
                  /* If-8 */} 
                /* Let-7 */} 
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (void_try6-Result) */
            if ErrorIn(void_try6) {Result = void_try6
            Result = void_try6
            break
            } else {
            }}
            /* loop-5 */} 
          /* For-4 */} 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        /* Let:4 */{ 
          var m *ClaireAny   = Core.F__at_property1(C_close,c).Id()
          /* noccur = 2 */
          /* Let:5 */{ 
            var g1309UU *ClaireAny  
            /* noccur = 1 */
            if (F_boolean_I_any(m) == CTRUE) /* If:6 */{ 
              /* Let:7 */{ 
                var _CL_obj *Language.CallMethod1   = Language.To_CallMethod1(new(Language.CallMethod1).Is(Language.C_Call_method1))
                /* noccur = 5 */
                _CL_obj.Arg = ToMethod(m)
                _CL_obj.Args = MakeConstantList(self)
                g1309UU = _CL_obj.Id()
                /* Let-7 */} 
              } else {
              g1309UU = self
              /* If-6 */} 
            r = r.AddFast(g1309UU)
            /* Let-5 */} 
          /* Let-4 */} 
        Result = EID{r.Id(),0}
        }
        }
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: analyze! @ class (throw: true) 
func E_Optimize_analyze_I_class (c EID,self EID,_Zl EID,lp EID) EID { 
    return /*(sm for analyze! @ class= EID)*/ F_Optimize_analyze_I_class(ToClass(OBJ(c)),
      ANY(self),
      ToList(OBJ(_Zl)),
      ToList(OBJ(lp)) )} 
  
// creation of a new named object
/* {1} OPT.The go function for: c_code(self:Defobj,s:class) [] */
func F_c_code_Defobj (self *Language.Defobj ,s *ClaireClass ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var _Zc *ClaireClass   = self.Arg
      /* noccur = 5 */
      /* Let:3 */{ 
        var o *ClaireAny   = self.Ident.Get()
        /* noccur = 7 */
        /* Let:4 */{ 
          var _Zv *ClaireAny  
          /* noccur = 3 */
          if ((o != CNULL) && 
              (o.Isa.IsIn(Core.C_global_variable) != CTRUE)) /* If:5 */{ 
            _Zv = o
            } else {
            /* Let:6 */{ 
              var g1312UU int 
              /* noccur = 1 */
              C_OPT.MaxVars = (C_OPT.MaxVars+1)
              g1312UU = 0
              _Zv = F_Compile_Variable_I_symbol(ToSymbol(C_Compile__starname_star.Value),g1312UU,_Zc.Id()).Id()
              /* Let-6 */} 
            /* If-5 */} 
          /* Let:5 */{ 
            var _Zy1 *Language.Call  
            /* noccur = 2 */
            /* Let:6 */{ 
              var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
              /* noccur = 5 */
              _CL_obj.Selector = C_Compile_object_I
              _CL_obj.Args = MakeConstantList(self.Ident.Id(),_Zc.Id())
              _Zy1 = _CL_obj
              /* Let-6 */} 
            /* Let:6 */{ 
              var _Zy2 *ClaireAny  
              /* noccur = 2 */
              var _Zy2_try13137 EID 
              _Zy2_try13137 = F_Optimize_analyze_I_class(_Zc,_Zv,self.Args,MakeConstantList(C_name.Id()))
              /* ERROR PROTECTION INSERTED (_Zy2-Result) */
              if ErrorIn(_Zy2_try13137) {Result = _Zy2_try13137
              } else {
              _Zy2 = ANY(_Zy2_try13137)
              /* Let:7 */{ 
                var _Zx *ClaireAny  
                /* noccur = 3 */
                if (_Zv.Isa.IsIn(C_Variable) != CTRUE) /* If:8 */{ 
                  /* Let:9 */{ 
                    var _CL_obj *Language.Do   = Language.To_Do(new(Language.Do).Is(Language.C_Do))
                    /* noccur = 3 */
                    _CL_obj.Args = F_cons_any(_Zy1.Id(),ToList(_Zy2))
                    _Zx = _CL_obj.Id()
                    /* Let-9 */} 
                  } else {
                  /* Let:9 */{ 
                    var _CL_obj *Language.Let   = Language.To_Let(new(Language.Let).Is(Language.C_Let))
                    /* noccur = 10 */
                    _CL_obj.ClaireVar = To_Variable(_Zv)
                    _CL_obj.Value = _Zy1.Id()
                    /* update:10 */{ 
                      var va_arg1 *Language.Let  
                      var va_arg2 *ClaireAny  
                      va_arg1 = _CL_obj
                      /* Let:11 */{ 
                        var _CL_obj *Language.Do   = Language.To_Do(new(Language.Do).Is(Language.C_Do))
                        /* noccur = 3 */
                        _CL_obj.Args = ToList(_Zy2)
                        va_arg2 = _CL_obj.Id()
                        /* Let-11 */} 
                      /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                      va_arg1.Arg = va_arg2
                      /* update-10 */} 
                    _Zx = _CL_obj.Id()
                    /* Let-9 */} 
                  /* If-8 */} 
                
                if (_Zc.Open <= 0) /* If:8 */{ 
                  Result = ToException(Core.C_general_error.Make(MakeString("[105] cannot instantiate ~S").Id(),MakeConstantList(_Zc.Id()).Id())).Close()
                  } else {
                  Result = EID{CFALSE.Id(),0}
                  /* If-8 */} 
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                if (o != CNULL) /* If:8 */{ 
                  if (ToBoolean(C_OPT.Objects.Contain_ask(o).Id()) != CTRUE) /* If:9 */{ 
                    C_OPT.Objects = C_OPT.Objects.AddFast(o)
                    Core.F_CALL(C_Optimize_c_register,ARGS(o.ToEID()))
                    /* If-9 */} 
                  } else {
                  F_Compile_warn_void()
                  Core.F_tformat_string(MakeString("~S is unknown [265]\n"),2,MakeConstantList(self.Ident.Id()))
                  /* If-8 */} 
                var _Zx_try13148 EID 
                _Zx_try13148 = Core.F_CALL(C_c_code,ARGS(_Zx.ToEID(),EID{s.Id(),0}))
                /* ERROR PROTECTION INSERTED (_Zx-Result) */
                if ErrorIn(_Zx_try13148) {Result = _Zx_try13148
                } else {
                _Zx = ANY(_Zx_try13148)
                Result = _Zx.ToEID()
                Result = _Zx.ToEID()
                }}
                /* Let-7 */} 
              }
              /* Let-6 */} 
            /* Let-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_code @ Defobj (throw: true) 
func E_c_code_Defobj (self EID,s EID) EID { 
    return /*(sm for c_code @ Defobj= EID)*/ F_c_code_Defobj(Language.To_Defobj(OBJ(self)),ToClass(OBJ(s)) )} 
  
// creation of a new named object
/* {1} OPT.The go function for: c_code(self:Defclass,s:class) [] */
func F_c_code_Defclass (self *Language.Defclass ,s *ClaireClass ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var _Zname *ClaireSymbol   = self.Ident
      /* noccur = 3 */
      /* Let:3 */{ 
        var o *ClaireAny   = _Zname.Get()
        /* noccur = 6 */
        /* Let:4 */{ 
          var _Zcreate *Language.Call  
          /* noccur = 1 */
          var _Zcreate_try13185 EID 
          if (o != CNULL) /* If:5 */{ 
            /* Let:6 */{ 
              var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
              /* noccur = 5 */
              _CL_obj.Selector = C_class_I
              _CL_obj.Args = MakeConstantList(_Zname.Id(),self.Arg.Id())
              _Zcreate_try13185 = EID{_CL_obj.Id(),0}
              /* Let-6 */} 
            } else {
            _Zcreate_try13185 = ToException(Core.C_general_error.Make(MakeString("[internal] cannot compile unknown class ~S").Id(),MakeConstantList(_Zname.Id()).Id())).Close()
            /* If-5 */} 
          /* ERROR PROTECTION INSERTED (_Zcreate-Result) */
          if ErrorIn(_Zcreate_try13185) {Result = _Zcreate_try13185
          } else {
          _Zcreate = Language.To_Call(OBJ(_Zcreate_try13185))
          /* Let:5 */{ 
            var _Zx *Language.Do  
            /* noccur = 1 */
            var _Zx_try13196 EID 
            /* Let:6 */{ 
              var _CL_obj *Language.Do   = Language.To_Do(new(Language.Do).Is(Language.C_Do))
              /* noccur = 15 */
              /* update:7 */{ 
                var va_arg1 *Language.Do  
                var va_arg2 *ClaireList  
                va_arg1 = _CL_obj
                var va_arg2_try13208 EID 
                /* Let:8 */{ 
                  var g1321UU *ClaireList  
                  /* noccur = 1 */
                  var g1321UU_try13229 EID 
                  /* Let:9 */{ 
                    var g1323UU *ClaireList  
                    /* noccur = 1 */
                    var g1323UU_try132510 EID 
                    /* Iteration:10 */{ 
                      var v_list10 *ClaireList  
                      var x *ClaireAny  
                      var v_local10 *ClaireAny  
                      v_list10 = self.Args
                      g1323UU_try132510 = EID{CreateList(ToType(CEMPTY.Id()),v_list10.Length()).Id(),0}
                      for CLcount := 0; CLcount < v_list10.Length(); CLcount++{ 
                        x = v_list10.At(CLcount)
                        var v_local10_try132612 EID 
                        /* Let:12 */{ 
                          var v *ClaireAny   = CNULL
                          /* noccur = 4 */
                          if (x.Isa.IsIn(Language.C_Call) == CTRUE) /* If:13 */{ 
                            /* Let:14 */{ 
                              var g1316 *Language.Call   = Language.To_Call(x)
                              /* noccur = 4 */
                              v = g1316.Args.At(2-1)
                              g1316 = Language.To_Call(g1316.Args.At(1-1))
                              x = g1316.Id()
                              v_local10_try132612 = x.ToEID()
                              /* Let-14 */} 
                            /* If!13 */}  else if (x.Isa.IsIn(C_Variable) == CTRUE) /* If:13 */{ 
                            /* Let:14 */{ 
                              var g1317 *ClaireVariable   = To_Variable(x)
                              /* noccur = 1 */
                              var v_try132715 EID 
                              /* Let:15 */{ 
                                var g1328UU *ClaireAny  
                                /* noccur = 1 */
                                var g1328UU_try132916 EID 
                                g1328UU_try132916 = Language.F_extract_type_any(g1317.Range.Id())
                                /* ERROR PROTECTION INSERTED (g1328UU-v_try132715) */
                                if ErrorIn(g1328UU_try132916) {v_try132715 = g1328UU_try132916
                                } else {
                                g1328UU = ANY(g1328UU_try132916)
                                v_try132715 = Language.F_Language_getDefault_type(ToType(g1328UU),v).ToEID()
                                }
                                /* Let-15 */} 
                              /* ERROR PROTECTION INSERTED (v-v_local10_try132612) */
                              if ErrorIn(v_try132715) {v_local10_try132612 = v_try132715
                              } else {
                              v = ANY(v_try132715)
                              v_local10_try132612 = v.ToEID()
                              }
                              /* Let-14 */} 
                            } else {
                            v_local10_try132612 = EID{CFALSE.Id(),0}
                            /* If-13 */} 
                          /* ERROR PROTECTION INSERTED (v_local10_try132612-v_local10_try132612) */
                          if ErrorIn(v_local10_try132612) {g1323UU_try132510 = v_local10_try132612
                          break
                          } else {
                          /* Let:13 */{ 
                            var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                            /* noccur = 5 */
                            _CL_obj.Selector = C_add_slot
                            /* update:14 */{ 
                              var va_arg1 *Language.Call  
                              var va_arg2 *ClaireList  
                              va_arg1 = _CL_obj
                              var va_arg2_try133015 EID 
                              /* Construct:15 */{ 
                                var v_bag_arg *ClaireAny  
                                va_arg2_try133015= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                                ToList(OBJ(va_arg2_try133015)).AddFast(o)
                                var v_bag_arg_try133116 EID 
                                v_bag_arg_try133116 = Language.F_make_a_property_any(ANY(Core.F_CALL(C_mClaire_pname,ARGS(x.ToEID()))))
                                /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try133015) */
                                if ErrorIn(v_bag_arg_try133116) {va_arg2_try133015 = v_bag_arg_try133116
                                } else {
                                v_bag_arg = ANY(v_bag_arg_try133116)
                                ToList(OBJ(va_arg2_try133015)).AddFast(v_bag_arg)
                                ToList(OBJ(va_arg2_try133015)).AddFast(ANY(Core.F_CALL(C_range,ARGS(x.ToEID()))))
                                ToList(OBJ(va_arg2_try133015)).AddFast(v)}
                                /* Construct-15 */} 
                              /* ERROR PROTECTION INSERTED (va_arg2-v_local10_try132612) */
                              if ErrorIn(va_arg2_try133015) {v_local10_try132612 = va_arg2_try133015
                              } else {
                              va_arg2 = ToList(OBJ(va_arg2_try133015))
                              /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                              va_arg1.Args = va_arg2
                              v_local10_try132612 = EID{va_arg2.Id(),0}
                              }
                              /* update-14 */} 
                            /* ERROR PROTECTION INSERTED (v_local10_try132612-v_local10_try132612) */
                            if ErrorIn(v_local10_try132612) {g1323UU_try132510 = v_local10_try132612
                            break
                            } else {
                            v_local10_try132612 = EID{_CL_obj.Id(),0}
                            }
                            /* Let-13 */} 
                          /* ERROR PROTECTION INSERTED (v_local10_try132612-v_local10_try132612) */
                          if ErrorIn(v_local10_try132612) {g1323UU_try132510 = v_local10_try132612
                          break
                          } else {
                          }}
                          /* Let-12 */} 
                        /* ERROR PROTECTION INSERTED (v_local10-g1323UU_try132510) */
                        if ErrorIn(v_local10_try132612) {g1323UU_try132510 = v_local10_try132612
                        g1323UU_try132510 = v_local10_try132612
                        break
                        } else {
                        v_local10 = ANY(v_local10_try132612)
                        ToList(OBJ(g1323UU_try132510)).PutAt(CLcount,v_local10)
                        } 
                      }
                      /* Iteration-10 */} 
                    /* ERROR PROTECTION INSERTED (g1323UU-g1321UU_try13229) */
                    if ErrorIn(g1323UU_try132510) {g1321UU_try13229 = g1323UU_try132510
                    } else {
                    g1323UU = ToList(OBJ(g1323UU_try132510))
                    /* Let:10 */{ 
                      var g1324UU *ClaireList  
                      /* noccur = 1 */
                      if (self.Params.Length() != 0) /* If:11 */{ 
                        /* Construct:12 */{ 
                          var v_bag_arg *ClaireAny  
                          g1324UU= ToType(CEMPTY.Id()).EmptyList()
                          /* Let:13 */{ 
                            var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                            /* noccur = 5 */
                            _CL_obj.Selector = C_put
                            _CL_obj.Args = MakeConstantList(C_params.Id(),o,self.Params.Id())
                            v_bag_arg = _CL_obj.Id()
                            /* Let-13 */} 
                          g1324UU.AddFast(v_bag_arg)/* Construct-12 */} 
                        } else {
                        g1324UU = ToType(CEMPTY.Id()).EmptyList()
                        /* If-11 */} 
                      g1321UU_try13229 = EID{g1323UU.Append(g1324UU).Id(),0}
                      /* Let-10 */} 
                    }
                    /* Let-9 */} 
                  /* ERROR PROTECTION INSERTED (g1321UU-va_arg2_try13208) */
                  if ErrorIn(g1321UU_try13229) {va_arg2_try13208 = g1321UU_try13229
                  } else {
                  g1321UU = ToList(OBJ(g1321UU_try13229))
                  va_arg2_try13208 = EID{F_cons_any(_Zcreate.Id(),g1321UU).Id(),0}
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (va_arg2-_Zx_try13196) */
                if ErrorIn(va_arg2_try13208) {_Zx_try13196 = va_arg2_try13208
                } else {
                va_arg2 = ToList(OBJ(va_arg2_try13208))
                /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                va_arg1.Args = va_arg2
                _Zx_try13196 = EID{va_arg2.Id(),0}
                }
                /* update-7 */} 
              /* ERROR PROTECTION INSERTED (_Zx_try13196-_Zx_try13196) */
              if !ErrorIn(_Zx_try13196) {
              _Zx_try13196 = EID{_CL_obj.Id(),0}
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (_Zx-Result) */
            if ErrorIn(_Zx_try13196) {Result = _Zx_try13196
            } else {
            _Zx = Language.To_Do(OBJ(_Zx_try13196))
            if (ToBoolean(C_OPT.Objects.Contain_ask(o).Id()) != CTRUE) /* If:6 */{ 
              C_OPT.Objects = C_OPT.Objects.AddFast(o)
              Core.F_CALL(C_Optimize_c_register,ARGS(o.ToEID()))
              /* If-6 */} 
            Result = Core.F_CALL(C_c_code,ARGS(EID{_Zx.Id(),0},EID{s.Id(),0}))
            }
            /* Let-5 */} 
          }
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_code @ Defclass (throw: true) 
func E_c_code_Defclass (self EID,s EID) EID { 
    return /*(sm for c_code @ Defclass= EID)*/ F_c_code_Defclass(Language.To_Defclass(OBJ(self)),ToClass(OBJ(s)) )} 
  
// method definition
// note (3.4): using the un-compiled code for c_status is weak, it would be much better to
//
/* {1} OPT.The go function for: c_type(self:Defmethod) [] */
func F_c_type_Defmethod (self *Language.Defmethod ) *ClaireType  { 
    // use function body compiling 
return  ToType(C_any.Id())
    } 
  
// The EID go function for: c_type @ Defmethod (throw: false) 
func E_c_type_Defmethod (self EID) EID { 
    return EID{/*(sm for c_type @ Defmethod= type)*/ F_c_type_Defmethod(Language.To_Defmethod(OBJ(self)) ).Id(),0}} 
  
/* {1} OPT.The go function for: c_code(self:Defmethod) [] */
func F_c_code_Defmethod (self *Language.Defmethod ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var px *ClaireProperty   = self.Arg.Selector
      /* noccur = 7 */
      /* Let:3 */{ 
        var l *ClaireList   = self.Arg.Args
        /* noccur = 3 */
        /* Let:4 */{ 
          var lv *ClaireList  
          /* noccur = 6 */
          if ((l.Length() == 1) && 
              (l.At(1-1) == ClEnv.Id())) /* If:5 */{ 
            lv = MakeConstantList(F_Compile_Variable_I_symbol(ToSymbol(C_Compile__starname_star.Value),0,C_void.Id()).Id())
            } else {
            lv = l
            /* If-5 */} 
          /* Let:5 */{ 
            var ls *ClaireList  
            /* noccur = 5 */
            var ls_try13346 EID 
            ls_try13346 = F_Optimize_extract_signature_I_list(lv)
            /* ERROR PROTECTION INSERTED (ls-Result) */
            if ErrorIn(ls_try13346) {Result = ls_try13346
            } else {
            ls = ToList(OBJ(ls_try13346))
            /* Let:6 */{ 
              var lrange *ClaireList  
              /* noccur = 8 */
              var lrange_try13357 EID 
              /* Iteration:7 */{ 
                var v_list7 *ClaireList  
                var x *ClaireAny  
                var v_local7 *ClaireAny  
                var v_list7_try13368 EID 
                v_list7_try13368 = Language.F_extract_range_any(self.SetArg,lv,ToList(Language.C_LDEF.Value))
                /* ERROR PROTECTION INSERTED (v_list7-lrange_try13357) */
                if ErrorIn(v_list7_try13368) {lrange_try13357 = v_list7_try13368
                } else {
                v_list7 = ToList(OBJ(v_list7_try13368))
                lrange_try13357 = EID{CreateList(ToType(C_any.Id()),v_list7.Length()).Id(),0}
                for CLcount := 0; CLcount < v_list7.Length(); CLcount++{ 
                  x = v_list7.At(CLcount)
                  v_local7 = x
                  ToList(OBJ(lrange_try13357)).PutAt(CLcount,v_local7)
                  } 
                }
                /* Iteration-7 */} 
              /* ERROR PROTECTION INSERTED (lrange-Result) */
              if ErrorIn(lrange_try13357) {Result = lrange_try13357
              } else {
              lrange = ToList(OBJ(lrange_try13357))
              /* Let:7 */{ 
                var sdef *ClaireAny  
                /* noccur = 1 */
                var sdef_try13378 EID 
                if ((self.Inline_ask == CTRUE) && 
                    (C_compiler.Inline_ask == CTRUE)) /* If:8 */{ 
                  Core.F_print_in_string_void()
                  PRINC("lambda[(")
                  sdef_try13378 = Language.F_ppvariable_list(lv)
                  /* ERROR PROTECTION INSERTED (sdef_try13378-sdef_try13378) */
                  if !ErrorIn(sdef_try13378) {
                  PRINC("),")
                  sdef_try13378 = Core.F_CALL(C_print,ARGS(self.Body.ToEID()))
                  /* ERROR PROTECTION INSERTED (sdef_try13378-sdef_try13378) */
                  if !ErrorIn(sdef_try13378) {
                  PRINC("]")
                  sdef_try13378 = EVOID
                  }}
                  /* ERROR PROTECTION INSERTED (sdef_try13378-sdef_try13378) */
                  if !ErrorIn(sdef_try13378) {
                  sdef_try13378 = Core.F_end_of_string_void()
                  }
                  } else {
                  sdef_try13378 = EID{CFALSE.Id(),0}
                  /* If-8 */} 
                /* ERROR PROTECTION INSERTED (sdef-Result) */
                if ErrorIn(sdef_try13378) {Result = sdef_try13378
                } else {
                sdef = ANY(sdef_try13378)
                /* Let:8 */{ 
                  var lbody *ClaireList  
                  /* noccur = 11 */
                  var lbody_try13389 EID 
                  lbody_try13389 = Language.F_extract_status_any(self.Body)
                  /* ERROR PROTECTION INSERTED (lbody-Result) */
                  if ErrorIn(lbody_try13389) {Result = lbody_try13389
                  } else {
                  lbody = ToList(OBJ(lbody_try13389))
                  /* Let:9 */{ 
                    var getm *ClaireAny   = ANY(Core.F_CALL(ToProperty(Core.C__at.Id()),ARGS(EID{px.Id(),0},ls.At(2-1).ToEID())))
                    /* noccur = 2 */
                    /* Let:10 */{ 
                      var m *ClaireMethod  
                      /* noccur = 8 */
                      var m_try133911 EID 
                      if (C_method.Id() == getm.Isa.Id()) /* If:11 */{ 
                        /* Let:12 */{ 
                          var g1332 *ClaireMethod   = ToMethod(getm)
                          /* noccur = 1 */
                          m_try133911 = EID{g1332.Id(),0}
                          /* Let-12 */} 
                        } else {
                        m_try133911 = ToException(Core.C_general_error.Make(MakeString("[internal] the method ~S @ ~S is not known").Id(),MakeConstantList(px.Id(),ls.At(2-1)).Id())).Close()
                        /* If-11 */} 
                      /* ERROR PROTECTION INSERTED (m-Result) */
                      if ErrorIn(m_try133911) {Result = m_try133911
                      } else {
                      m = ToMethod(OBJ(m_try133911))
                      ToArray(lbody.Id()).NthPut(2,Core.F_get_property(C_functional,ToObject(m.Id())))
                      Core.F_put_table(C_Compile_FileOrigin,m.Id(),(F_append_string(F_append_string(ToCompileProducer(C_PRODUCER.Value).CurrentFile,MakeString(".cl:")),F_string_I_integer(ClEnv.NLine))).Id())
                      if ((C_compiler.Inline_ask != CTRUE) && 
                          ((px.Id() == Language.C_Iterate.Id()) || 
                              (px.Id() == Language.C_iterate.Id()))) /* If:11 */{ 
                        Result = EID{CNIL.Id(),0}
                        } else {
                        var g1343I *ClaireBoolean  
                        var g1343I_try134412 EID 
                        /* and:12 */{ 
                          var v_and12 *ClaireBoolean  
                          
                          v_and12 = Equal(lrange.At(1-1),C_void.Id())
                          if (v_and12 == CFALSE) {g1343I_try134412 = EID{CFALSE.Id(),0}
                          } else /* arg:13 */{ 
                            var v_and12_try134514 EID 
                            v_and12_try134514 = F_Optimize_sort_pattern_ask_list(lv,self.Body)
                            /* ERROR PROTECTION INSERTED (v_and12-g1343I_try134412) */
                            if ErrorIn(v_and12_try134514) {g1343I_try134412 = v_and12_try134514
                            } else {
                            v_and12 = ToBoolean(OBJ(v_and12_try134514))
                            if (v_and12 == CFALSE) {g1343I_try134412 = EID{CFALSE.Id(),0}
                            } else /* arg:14 */{ 
                              g1343I_try134412 = EID{CTRUE.Id(),0}/* arg-14 */} 
                            /* arg-13 */} 
                          }
                          /* and-12 */} 
                        /* ERROR PROTECTION INSERTED (g1343I-Result) */
                        if ErrorIn(g1343I_try134412) {Result = g1343I_try134412
                        } else {
                        g1343I = ToBoolean(OBJ(g1343I_try134412))
                        if (g1343I == CTRUE) /* If:12 */{ 
                          Result = F_Optimize_sort_code_Defmethod(self,lv)
                          } else {
                          if (lbody.At(3-1) != C_body.Id()) /* If:13 */{ 
                            /* Let:14 */{ 
                              var na *ClaireString  
                              /* noccur = 2 */
                              var na_try134615 EID 
                              na_try134615 = Core.F_CALL(C_Compile_function_name,ARGS(EID{px.Id(),0},ls.At(2-1).ToEID(),lbody.At(2-1).ToEID()))
                              /* ERROR PROTECTION INSERTED (na-Result) */
                              if ErrorIn(na_try134615) {Result = na_try134615
                              } else {
                              na = ToString(OBJ(na_try134615))
                              /* Let:15 */{ 
                                var la *ClaireLambda  
                                /* noccur = 1 */
                                var la_try134716 EID 
                                la_try134716 = Language.F_lambda_I_list(lv,lbody.At(3-1))
                                /* ERROR PROTECTION INSERTED (la-Result) */
                                if ErrorIn(la_try134716) {Result = la_try134716
                                } else {
                                la = ToLambda(OBJ(la_try134716))
                                /* Let:16 */{ 
                                  var news int 
                                  /* noccur = 1 */
                                  var news_try134917 EID 
                                  var g1350I *ClaireBoolean  
                                  var g1350I_try135117 EID 
                                  if (C_OPT.Recompute == CTRUE) /* If:17 */{ 
                                    g1350I_try135117 = F_Compile_g_throw_any(lbody.At(2-1))
                                    } else {
                                    g1350I_try135117 = F_Compile_can_throw_ask_method(m)
                                    /* If-17 */} 
                                  /* ERROR PROTECTION INSERTED (g1350I-news_try134917) */
                                  if ErrorIn(g1350I_try135117) {news_try134917 = g1350I_try135117
                                  } else {
                                  g1350I = ToBoolean(OBJ(g1350I_try135117))
                                  if (g1350I == CTRUE) /* If:17 */{ 
                                    news_try134917 = EID{C__INT,IVAL(1)}
                                    } else {
                                    news_try134917 = EID{C__INT,IVAL(0)}
                                    /* If-17 */} 
                                  }
                                  /* ERROR PROTECTION INSERTED (news-Result) */
                                  if ErrorIn(news_try134917) {Result = news_try134917
                                  } else {
                                  news = INT(news_try134917)
                                  Result = F_Compile_compile_lambda_string(na,la,m.Id())
                                  /* ERROR PROTECTION INSERTED (Result-Result) */
                                  if !ErrorIn(Result) {
                                  if ((lbody.At(1-1) == CNULL) || 
                                      (C_OPT.Recompute == CTRUE)) /* If:17 */{ 
                                    ToArray(lbody.Id()).NthPut(1,MakeInteger(news).Id())
                                    /* If-17 */} 
                                  Result = ToArray(lbody.Id()).NthPut(2,F_make_function_string(na).Id()).ToEID()
                                  }
                                  }
                                  /* Let-16 */} 
                                }
                                /* Let-15 */} 
                              }
                              /* Let-14 */} 
                            } else {
                            Result = EID{CFALSE.Id(),0}
                            /* If-13 */} 
                          /* ERROR PROTECTION INSERTED (Result-Result) */
                          if !ErrorIn(Result) {
                          if (self.SetArg.Isa.IsIn(Core.C_global_variable) == CTRUE) /* If:13 */{ 
                            ToArray(lrange.Id()).NthPut(1,self.SetArg)
                            /* If!13 */}  else if ((C_class.Id() == m.Range.Isa.Id()) && 
                              (C_class.Id() != lrange.At(1-1).Isa.Id())) /* If:13 */{ 
                            ToArray(lrange.Id()).NthPut(1,m.Range.Id())
                            /* If-13 */} 
                          /* Let:13 */{ 
                            var _Zm *ClaireAny  
                            /* noccur = 4 */
                            var _Zm_try135214 EID 
                            _Zm_try135214 = F_Optimize_add_method_I_method(m,
                              ToList(ls.At(1-1)),
                              lrange.At(1-1),
                              lbody.At(1-1),
                              ToFunction(lbody.At(2-1)))
                            /* ERROR PROTECTION INSERTED (_Zm-Result) */
                            if ErrorIn(_Zm_try135214) {Result = _Zm_try135214
                            } else {
                            _Zm = ANY(_Zm_try135214)
                            Core.F_tformat_string(MakeString("c_code@defmethod produces ~S with lrange = ~S \n"),0,MakeConstantList(_Zm,lrange.Id()))
                            if (px.Id() == C_nth.Id()) /* If:14 */{ 
                              Core.F_tformat_string(MakeString("*********************** LOOK ********************* \n"),0,ToType(CEMPTY.Id()).EmptyList())
                              /* If-14 */} 
                            /* Let:14 */{ 
                              var g1353UU *ClaireAny  
                              /* noccur = 1 */
                              var g1353UU_try135415 EID 
                              if ((self.Inline_ask == CTRUE) && 
                                  ((C_compiler.Inline_ask == CTRUE) && 
                                    (C_compiler.Diet_ask != CTRUE))) /* If:15 */{ 
                                /* Let:16 */{ 
                                  var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                  /* noccur = 5 */
                                  _CL_obj.Selector = Core.C_inlineok_ask
                                  _CL_obj.Args = MakeConstantList(_Zm,sdef)
                                  g1353UU_try135415 = EID{_CL_obj.Id(),0}
                                  /* Let-16 */} 
                                /* If!15 */}  else if ((F_boolean_I_any(lrange.At(2-1)) == CTRUE) && 
                                  (C_compiler.Diet_ask != CTRUE)) /* If:15 */{ 
                                /* Let:16 */{ 
                                  var na *ClaireString  
                                  /* noccur = 2 */
                                  var na_try135517 EID 
                                  /* Let:17 */{ 
                                    var g1356UU *ClaireAny  
                                    /* noccur = 1 */
                                    var g1356UU_try135718 EID 
                                    g1356UU_try135718 = Core.F_CALL(C_Compile_function_name,ARGS(EID{px.Id(),0},ls.At(2-1).ToEID(),lbody.At(2-1).ToEID()))
                                    /* ERROR PROTECTION INSERTED (g1356UU-na_try135517) */
                                    if ErrorIn(g1356UU_try135718) {na_try135517 = g1356UU_try135718
                                    } else {
                                    g1356UU = ANY(g1356UU_try135718)
                                    na_try135517 = EID{F_Optimize_type_extension_string(ToString(g1356UU)).Id(),0}
                                    }
                                    /* Let-17 */} 
                                  /* ERROR PROTECTION INSERTED (na-g1353UU_try135415) */
                                  if ErrorIn(na_try135517) {g1353UU_try135415 = na_try135517
                                  } else {
                                  na = ToString(OBJ(na_try135517))
                                  /* Let:17 */{ 
                                    var _Zf *ClaireFunction   = F_make_function_string(na)
                                    /* noccur = 2 */
                                    g1353UU_try135415 = F_Compile_compile_lambda_string(na,ToLambda(lrange.At(2-1)),C_type.Id())
                                    /* ERROR PROTECTION INSERTED (g1353UU_try135415-g1353UU_try135415) */
                                    if !ErrorIn(g1353UU_try135415) {
                                    F_set_arity_function(_Zf,m.Domain.Length())
                                    /* Let:18 */{ 
                                      var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                      /* noccur = 5 */
                                      _CL_obj.Selector = Core.C_write
                                      _CL_obj.Args = MakeConstantList(Language.C_iClaire_typing.Value,_Zm,_Zf.Id())
                                      g1353UU_try135415 = EID{_CL_obj.Id(),0}
                                      /* Let-18 */} 
                                    }
                                    /* Let-17 */} 
                                  }
                                  /* Let-16 */} 
                                } else {
                                g1353UU_try135415 = _Zm.ToEID()
                                /* If-15 */} 
                              /* ERROR PROTECTION INSERTED (g1353UU-Result) */
                              if ErrorIn(g1353UU_try135415) {Result = g1353UU_try135415
                              } else {
                              g1353UU = ANY(g1353UU_try135415)
                              Result = Core.F_CALL(C_c_code,ARGS(g1353UU.ToEID()))
                              }
                              /* Let-14 */} 
                            }
                            /* Let-13 */} 
                          }
                          /* If-12 */} 
                        }
                        /* If-11 */} 
                      }
                      /* Let-10 */} 
                    /* Let-9 */} 
                  }
                  /* Let-8 */} 
                }
                /* Let-7 */} 
              }
              /* Let-6 */} 
            }
            /* Let-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_code @ Defmethod (throw: true) 
func E_c_code_Defmethod (self EID) EID { 
    return /*(sm for c_code @ Defmethod= EID)*/ F_c_code_Defmethod(Language.To_Defmethod(OBJ(self)) )} 
  
// create a type function name by adding a "_type" - will not be imported
/* {1} OPT.The go function for: type_extension(s:string) [] */
func F_Optimize_type_extension_string (s *ClaireString ) *ClaireString  { 
    // procedure body with s =  
var Result *ClaireString  
    /* Let:2 */{ 
      var n int  = F_length_string(s)
      /* noccur = 1 */
      /* Let:3 */{ 
        var f *ClaireString  
        /* noccur = 1 */
        if (s.At(1) == '#') /* If:4 */{ 
          f = F_substring_string(s,2,n)
          } else {
          f = s
          /* If-4 */} 
        Result = F_append_string(f,MakeString("_type"))
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: type_extension @ string (throw: false) 
func E_Optimize_type_extension_string (s EID) EID { 
    return EID{/*(sm for type_extension @ string= string)*/ F_Optimize_type_extension_string(ToString(OBJ(s)) ).Id(),0}} 
  
// v3.3 : we optimize a single sort definition -----------------------------------------------
// [foo(x:list) : list -> sort(m,x) ]
/* {1} OPT.The go function for: sort_pattern?(lv:list,%body:any) [] */
func F_Optimize_sort_pattern_ask_list (lv *ClaireList ,_Zbody *ClaireAny ) EID { 
    var Result EID 
    /* and:2 */{ 
      var v_and2 *ClaireBoolean  
      
      v_and2 = Equal(MakeInteger(lv.Length()).Id(),MakeInteger(1).Id())
      if (v_and2 == CFALSE) {Result = EID{CFALSE.Id(),0}
      } else /* arg:3 */{ 
        var v_and2_try13614 EID 
        if (_Zbody.Isa.IsIn(Language.C_Call) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var g1358 *Language.Call   = Language.To_Call(_Zbody)
            /* noccur = 3 */
            /* and:6 */{ 
              var v_and6 *ClaireBoolean  
              
              v_and6 = Equal(g1358.Selector.Id(),Core.C_sort.Id())
              if (v_and6 == CFALSE) {v_and2_try13614 = EID{CFALSE.Id(),0}
              } else /* arg:7 */{ 
                /* Let:8 */{ 
                  var a1 *ClaireAny   = g1358.Args.At(1-1)
                  /* noccur = 2 */
                  if (a1.Isa.IsIn(Language.C_Call) == CTRUE) /* If:9 */{ 
                    /* Let:10 */{ 
                      var g1359 *Language.Call   = Language.To_Call(a1)
                      /* noccur = 2 */
                      v_and6 = MakeBoolean((g1359.Selector.Id() == Core.C__at.Id()) && (g1359.Args.At(1-1).Isa.IsIn(C_property) == CTRUE))
                      /* Let-10 */} 
                    } else {
                    v_and6 = CFALSE
                    /* If-9 */} 
                  /* Let-8 */} 
                if (v_and6 == CFALSE) {v_and2_try13614 = EID{CFALSE.Id(),0}
                } else /* arg:8 */{ 
                  var v_and6_try13629 EID 
                  /* Let:9 */{ 
                    var g1363UU *ClaireAny  
                    /* noccur = 1 */
                    var g1363UU_try136410 EID 
                    g1363UU_try136410 = Language.F_lexical_build_any(g1358.Args.At(2-1),lv,0)
                    /* ERROR PROTECTION INSERTED (g1363UU-v_and6_try13629) */
                    if ErrorIn(g1363UU_try136410) {v_and6_try13629 = g1363UU_try136410
                    } else {
                    g1363UU = ANY(g1363UU_try136410)
                    v_and6_try13629 = EID{Equal(g1363UU,lv.At(1-1)).Id(),0}
                    }
                    /* Let-9 */} 
                  /* ERROR PROTECTION INSERTED (v_and6-v_and2_try13614) */
                  if ErrorIn(v_and6_try13629) {v_and2_try13614 = v_and6_try13629
                  } else {
                  v_and6 = ToBoolean(OBJ(v_and6_try13629))
                  if (v_and6 == CFALSE) {v_and2_try13614 = EID{CFALSE.Id(),0}
                  } else /* arg:9 */{ 
                    v_and2_try13614 = EID{CTRUE.Id(),0}/* arg-9 */} 
                  /* arg-8 */} 
                /* arg-7 */} 
              }
              /* and-6 */} 
            /* Let-5 */} 
          } else {
          v_and2_try13614 = EID{CFALSE.Id(),0}
          /* If-4 */} 
        /* ERROR PROTECTION INSERTED (v_and2-Result) */
        if ErrorIn(v_and2_try13614) {Result = v_and2_try13614
        } else {
        v_and2 = ToBoolean(OBJ(v_and2_try13614))
        if (v_and2 == CFALSE) {Result = EID{CFALSE.Id(),0}
        } else /* arg:4 */{ 
          Result = EID{CTRUE.Id(),0}/* arg-4 */} 
        /* arg-3 */} 
      }
      /* and-2 */} 
    return Result} 
  
// The EID go function for: sort_pattern? @ list (throw: true) 
func E_Optimize_sort_pattern_ask_list (lv EID,_Zbody EID) EID { 
    return /*(sm for sort_pattern? @ list= EID)*/ F_Optimize_sort_pattern_ask_list(ToList(OBJ(lv)),ANY(_Zbody) )} 
  
// this is the macroexpansion of the quick_sort which is difficult because of the dual recursion
// Thus, we generate two methods for one definition, and produce the explicit code for the specialized
// quicksort (v3.3)
/* {1} OPT.The go function for: sort_code(self:Defmethod,lv:list) [] */
func F_Optimize_sort_code_Defmethod (self *Language.Defmethod ,lv *ClaireList ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var l *ClaireAny   = lv.At(1-1)
      /* noccur = 19 */
      /* Let:3 */{ 
        var f *ClaireAny   = ToList(OBJ(Core.F_CALL(C_args,ARGS(ToList(OBJ(Core.F_CALL(C_args,ARGS(self.Body.ToEID())))).At(1-1).ToEID())))).At(1-1)
        /* noccur = 2 */
        /* Let:4 */{ 
          var m *ClaireVariable   = F_Compile_Variable_I_symbol(Core.F_symbol_I_string2(MakeString("m")),0,C_integer.Id())
          /* noccur = 9 */
          /* Let:5 */{ 
            var n *ClaireVariable   = F_Compile_Variable_I_symbol(Core.F_symbol_I_string2(MakeString("n")),0,C_integer.Id())
            /* noccur = 18 */
            /* Let:6 */{ 
              var x *ClaireVariable   = F_Compile_Variable_I_symbol(Core.F_symbol_I_string2(MakeString("x")),0,Core.F_member_type(ToType(OBJ(Core.F_CALL(C_range,ARGS(l.ToEID()))))).Id())
              /* noccur = 6 */
              /* Let:7 */{ 
                var p *ClaireVariable   = F_Compile_Variable_I_symbol(Core.F_symbol_I_string2(MakeString("p")),0,C_integer.Id())
                /* noccur = 9 */
                /* Let:8 */{ 
                  var q *ClaireVariable   = F_Compile_Variable_I_symbol(Core.F_symbol_I_string2(MakeString("q")),0,C_integer.Id())
                  /* noccur = 2 */
                  /* Let:9 */{ 
                    var def1 *Language.Defmethod  
                    /* noccur = 1 */
                    /* Let:10 */{ 
                      var _CL_obj *Language.Defmethod   = Language.To_Defmethod(new(Language.Defmethod).Is(Language.C_Defmethod))
                      /* noccur = 20 */
                      _CL_obj.Arg = self.Arg
                      _CL_obj.Inline_ask = CFALSE
                      _CL_obj.SetArg = self.SetArg
                      /* update:11 */{ 
                        var va_arg1 *Language.Defmethod  
                        var va_arg2 *ClaireAny  
                        va_arg1 = _CL_obj
                        /* Let:12 */{ 
                          var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                          /* noccur = 11 */
                          _CL_obj.Selector = self.Arg.Selector
                          /* update:13 */{ 
                            var va_arg1 *Language.Call  
                            var va_arg2 *ClaireList  
                            va_arg1 = _CL_obj
                            /* Construct:14 */{ 
                              var v_bag_arg *ClaireAny  
                              va_arg2= ToType(CEMPTY.Id()).EmptyList()
                              va_arg2.AddFast(MakeInteger(1).Id())
                              /* Let:15 */{ 
                                var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                /* noccur = 5 */
                                _CL_obj.Selector = C_length
                                _CL_obj.Args = MakeConstantList(lv.At(1-1))
                                v_bag_arg = _CL_obj.Id()
                                /* Let-15 */} 
                              va_arg2.AddFast(v_bag_arg)
                              va_arg2.AddFast(l)/* Construct-14 */} 
                            /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                            va_arg1.Args = va_arg2
                            /* update-13 */} 
                          va_arg2 = _CL_obj.Id()
                          /* Let-12 */} 
                        /* ---------- now we compile update body(va_arg1) := va_arg2 ------- */
                        va_arg1.Body = va_arg2
                        /* update-11 */} 
                      def1 = _CL_obj
                      /* Let-10 */} 
                    /* Let:10 */{ 
                      var _Zbd *Language.If  
                      /* noccur = 1 */
                      /* Let:11 */{ 
                        var _CL_obj *Language.If   = Language.To_If(new(Language.If).Is(Language.C_If))
                        /* noccur = 248 */
                        /* update:12 */{ 
                          var va_arg1 *Language.If  
                          var va_arg2 *ClaireAny  
                          va_arg1 = _CL_obj
                          /* Let:13 */{ 
                            var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                            /* noccur = 5 */
                            _CL_obj.Selector = ToProperty(C__sup.Id())
                            _CL_obj.Args = MakeConstantList(m.Id(),n.Id())
                            va_arg2 = _CL_obj.Id()
                            /* Let-13 */} 
                          /* ---------- now we compile update iClaire/test(va_arg1) := va_arg2 ------- */
                          va_arg1.Test = va_arg2
                          /* update-12 */} 
                        /* update:12 */{ 
                          var va_arg1 *Language.If  
                          var va_arg2 *ClaireAny  
                          va_arg1 = _CL_obj
                          /* Let:13 */{ 
                            var _CL_obj *Language.Let   = Language.To_Let(new(Language.Let).Is(Language.C_Let))
                            /* noccur = 236 */
                            _CL_obj.ClaireVar = x
                            /* update:14 */{ 
                              var va_arg1 *Language.Let  
                              var va_arg2 *ClaireAny  
                              va_arg1 = _CL_obj
                              /* Let:15 */{ 
                                var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                /* noccur = 5 */
                                _CL_obj.Selector = C_nth
                                _CL_obj.Args = MakeConstantList(l,n.Id())
                                va_arg2 = _CL_obj.Id()
                                /* Let-15 */} 
                              /* ---------- now we compile update value(va_arg1) := va_arg2 ------- */
                              va_arg1.Value = va_arg2
                              /* update-14 */} 
                            /* update:14 */{ 
                              var va_arg1 *Language.Let  
                              var va_arg2 *ClaireAny  
                              va_arg1 = _CL_obj
                              /* Let:15 */{ 
                                var g1365UU *Language.Call  
                                /* noccur = 1 */
                                /* Let:16 */{ 
                                  var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                  /* noccur = 11 */
                                  _CL_obj.Selector = ToProperty(C__equal.Id())
                                  /* update:17 */{ 
                                    var va_arg1 *Language.Call  
                                    var va_arg2 *ClaireList  
                                    va_arg1 = _CL_obj
                                    /* Construct:18 */{ 
                                      var v_bag_arg *ClaireAny  
                                      va_arg2= ToType(CEMPTY.Id()).EmptyList()
                                      va_arg2.AddFast(m.Id())
                                      /* Let:19 */{ 
                                        var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                        /* noccur = 5 */
                                        _CL_obj.Selector = ToProperty(Core.C__plus.Id())
                                        _CL_obj.Args = MakeConstantList(n.Id(),MakeInteger(1).Id())
                                        v_bag_arg = _CL_obj.Id()
                                        /* Let-19 */} 
                                      va_arg2.AddFast(v_bag_arg)/* Construct-18 */} 
                                    /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                                    va_arg1.Args = va_arg2
                                    /* update-17 */} 
                                  g1365UU = _CL_obj
                                  /* Let-16 */} 
                                /* Let:16 */{ 
                                  var g1366UU *Language.If  
                                  /* noccur = 1 */
                                  /* Let:17 */{ 
                                    var _CL_obj *Language.If   = Language.To_If(new(Language.If).Is(Language.C_If))
                                    /* noccur = 39 */
                                    /* update:18 */{ 
                                      var va_arg1 *Language.If  
                                      var va_arg2 *ClaireAny  
                                      va_arg1 = _CL_obj
                                      /* Let:19 */{ 
                                        var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                        /* noccur = 11 */
                                        _CL_obj.Selector = ToProperty(f)
                                        /* update:20 */{ 
                                          var va_arg1 *Language.Call  
                                          var va_arg2 *ClaireList  
                                          va_arg1 = _CL_obj
                                          /* Construct:21 */{ 
                                            var v_bag_arg *ClaireAny  
                                            va_arg2= ToType(CEMPTY.Id()).EmptyList()
                                            /* Let:22 */{ 
                                              var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                              /* noccur = 5 */
                                              _CL_obj.Selector = C_nth
                                              _CL_obj.Args = MakeConstantList(l,m.Id())
                                              v_bag_arg = _CL_obj.Id()
                                              /* Let-22 */} 
                                            va_arg2.AddFast(v_bag_arg)
                                            va_arg2.AddFast(x.Id())/* Construct-21 */} 
                                          /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                                          va_arg1.Args = va_arg2
                                          /* update-20 */} 
                                        va_arg2 = _CL_obj.Id()
                                        /* Let-19 */} 
                                      /* ---------- now we compile update iClaire/test(va_arg1) := va_arg2 ------- */
                                      va_arg1.Test = va_arg2
                                      /* update-18 */} 
                                    /* update:18 */{ 
                                      var va_arg1 *Language.If  
                                      var va_arg2 *ClaireAny  
                                      va_arg1 = _CL_obj
                                      /* Let:19 */{ 
                                        var _CL_obj *Language.Do   = Language.To_Do(new(Language.Do).Is(Language.C_Do))
                                        /* noccur = 21 */
                                        /* update:20 */{ 
                                          var va_arg1 *Language.Do  
                                          var va_arg2 *ClaireList  
                                          va_arg1 = _CL_obj
                                          /* Construct:21 */{ 
                                            var v_bag_arg *ClaireAny  
                                            va_arg2= ToType(CEMPTY.Id()).EmptyList()
                                            /* Let:22 */{ 
                                              var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                              /* noccur = 11 */
                                              _CL_obj.Selector = C_nth_equal
                                              /* update:23 */{ 
                                                var va_arg1 *Language.Call  
                                                var va_arg2 *ClaireList  
                                                va_arg1 = _CL_obj
                                                /* Construct:24 */{ 
                                                  var v_bag_arg *ClaireAny  
                                                  va_arg2= ToType(CEMPTY.Id()).EmptyList()
                                                  va_arg2.AddFast(l)
                                                  va_arg2.AddFast(n.Id())
                                                  /* Let:25 */{ 
                                                    var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                                    /* noccur = 5 */
                                                    _CL_obj.Selector = C_nth
                                                    _CL_obj.Args = MakeConstantList(l,m.Id())
                                                    v_bag_arg = _CL_obj.Id()
                                                    /* Let-25 */} 
                                                  va_arg2.AddFast(v_bag_arg)/* Construct-24 */} 
                                                /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                                                va_arg1.Args = va_arg2
                                                /* update-23 */} 
                                              v_bag_arg = _CL_obj.Id()
                                              /* Let-22 */} 
                                            va_arg2.AddFast(v_bag_arg)
                                            /* Let:22 */{ 
                                              var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                              /* noccur = 5 */
                                              _CL_obj.Selector = C_nth_equal
                                              _CL_obj.Args = MakeConstantList(l,m.Id(),x.Id())
                                              v_bag_arg = _CL_obj.Id()
                                              /* Let-22 */} 
                                            va_arg2.AddFast(v_bag_arg)/* Construct-21 */} 
                                          /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                                          va_arg1.Args = va_arg2
                                          /* update-20 */} 
                                        va_arg2 = _CL_obj.Id()
                                        /* Let-19 */} 
                                      /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                                      va_arg1.Arg = va_arg2
                                      /* update-18 */} 
                                    g1366UU = _CL_obj
                                    /* Let-17 */} 
                                  /* Let:17 */{ 
                                    var g1367UU *Language.Let  
                                    /* noccur = 1 */
                                    /* Let:18 */{ 
                                      var _CL_obj *Language.Let   = Language.To_Let(new(Language.Let).Is(Language.C_Let))
                                      /* noccur = 171 */
                                      _CL_obj.ClaireVar = p
                                      /* update:19 */{ 
                                        var va_arg1 *Language.Let  
                                        var va_arg2 *ClaireAny  
                                        va_arg1 = _CL_obj
                                        /* Let:20 */{ 
                                          var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                          /* noccur = 11 */
                                          _CL_obj.Selector = ToProperty(Core.C__sup_sup.Id())
                                          /* update:21 */{ 
                                            var va_arg1 *Language.Call  
                                            var va_arg2 *ClaireList  
                                            va_arg1 = _CL_obj
                                            /* Construct:22 */{ 
                                              var v_bag_arg *ClaireAny  
                                              va_arg2= ToType(CEMPTY.Id()).EmptyList()
                                              /* Let:23 */{ 
                                                var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                                /* noccur = 5 */
                                                _CL_obj.Selector = ToProperty(Core.C__plus.Id())
                                                _CL_obj.Args = MakeConstantList(n.Id(),m.Id())
                                                v_bag_arg = _CL_obj.Id()
                                                /* Let-23 */} 
                                              va_arg2.AddFast(v_bag_arg)
                                              va_arg2.AddFast(MakeInteger(1).Id())/* Construct-22 */} 
                                            /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                                            va_arg1.Args = va_arg2
                                            /* update-21 */} 
                                          va_arg2 = _CL_obj.Id()
                                          /* Let-20 */} 
                                        /* ---------- now we compile update value(va_arg1) := va_arg2 ------- */
                                        va_arg1.Value = va_arg2
                                        /* update-19 */} 
                                      /* update:19 */{ 
                                        var va_arg1 *Language.Let  
                                        var va_arg2 *ClaireAny  
                                        va_arg1 = _CL_obj
                                        /* Let:20 */{ 
                                          var _CL_obj *Language.Let   = Language.To_Let(new(Language.Let).Is(Language.C_Let))
                                          /* noccur = 152 */
                                          _CL_obj.ClaireVar = q
                                          _CL_obj.Value = n.Id()
                                          /* update:21 */{ 
                                            var va_arg1 *Language.Let  
                                            var va_arg2 *ClaireAny  
                                            va_arg1 = _CL_obj
                                            /* Let:22 */{ 
                                              var _CL_obj *Language.Do   = Language.To_Do(new(Language.Do).Is(Language.C_Do))
                                              /* noccur = 145 */
                                              /* update:23 */{ 
                                                var va_arg1 *Language.Do  
                                                var va_arg2 *ClaireList  
                                                va_arg1 = _CL_obj
                                                /* Construct:24 */{ 
                                                  var v_bag_arg *ClaireAny  
                                                  va_arg2= ToType(CEMPTY.Id()).EmptyList()
                                                  /* Let:25 */{ 
                                                    var g1368UU *Language.Call  
                                                    /* noccur = 1 */
                                                    /* Let:26 */{ 
                                                      var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                                      /* noccur = 5 */
                                                      _CL_obj.Selector = C_nth
                                                      _CL_obj.Args = MakeConstantList(l,p.Id())
                                                      g1368UU = _CL_obj
                                                      /* Let-26 */} 
                                                    v_bag_arg = Language.C_Assign.Make(x.Id(),g1368UU.Id())
                                                    /* Let-25 */} 
                                                  va_arg2.AddFast(v_bag_arg)
                                                  /* Let:25 */{ 
                                                    var _CL_obj *Language.If   = Language.To_If(new(Language.If).Is(Language.C_If))
                                                    /* noccur = 23 */
                                                    /* update:26 */{ 
                                                      var va_arg1 *Language.If  
                                                      var va_arg2 *ClaireAny  
                                                      va_arg1 = _CL_obj
                                                      /* Let:27 */{ 
                                                        var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                                        /* noccur = 5 */
                                                        _CL_obj.Selector = ToProperty(Core.C__I_equal.Id())
                                                        _CL_obj.Args = MakeConstantList(p.Id(),n.Id())
                                                        va_arg2 = _CL_obj.Id()
                                                        /* Let-27 */} 
                                                      /* ---------- now we compile update iClaire/test(va_arg1) := va_arg2 ------- */
                                                      va_arg1.Test = va_arg2
                                                      /* update-26 */} 
                                                    /* update:26 */{ 
                                                      var va_arg1 *Language.If  
                                                      var va_arg2 *ClaireAny  
                                                      va_arg1 = _CL_obj
                                                      /* Let:27 */{ 
                                                        var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                                        /* noccur = 11 */
                                                        _CL_obj.Selector = C_nth_equal
                                                        /* update:28 */{ 
                                                          var va_arg1 *Language.Call  
                                                          var va_arg2 *ClaireList  
                                                          va_arg1 = _CL_obj
                                                          /* Construct:29 */{ 
                                                            var v_bag_arg *ClaireAny  
                                                            va_arg2= ToType(CEMPTY.Id()).EmptyList()
                                                            va_arg2.AddFast(l)
                                                            va_arg2.AddFast(p.Id())
                                                            /* Let:30 */{ 
                                                              var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                                              /* noccur = 5 */
                                                              _CL_obj.Selector = C_nth
                                                              _CL_obj.Args = MakeConstantList(l,n.Id())
                                                              v_bag_arg = _CL_obj.Id()
                                                              /* Let-30 */} 
                                                            va_arg2.AddFast(v_bag_arg)/* Construct-29 */} 
                                                          /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                                                          va_arg1.Args = va_arg2
                                                          /* update-28 */} 
                                                        va_arg2 = _CL_obj.Id()
                                                        /* Let-27 */} 
                                                      /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                                                      va_arg1.Arg = va_arg2
                                                      /* update-26 */} 
                                                    v_bag_arg = _CL_obj.Id()
                                                    /* Let-25 */} 
                                                  va_arg2.AddFast(v_bag_arg)
                                                  /* Let:25 */{ 
                                                    var _CL_obj *Language.For   = Language.To_For(new(Language.For).Is(Language.C_For))
                                                    /* noccur = 81 */
                                                    _CL_obj.ClaireVar = p
                                                    /* update:26 */{ 
                                                      var va_arg1 *Language.Iteration  
                                                      var va_arg2 *ClaireAny  
                                                      va_arg1 = Language.To_Iteration(_CL_obj.Id())
                                                      /* Let:27 */{ 
                                                        var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                                        /* noccur = 11 */
                                                        _CL_obj.Selector = ToProperty(C__dot_dot.Id())
                                                        /* update:28 */{ 
                                                          var va_arg1 *Language.Call  
                                                          var va_arg2 *ClaireList  
                                                          va_arg1 = _CL_obj
                                                          /* Construct:29 */{ 
                                                            var v_bag_arg *ClaireAny  
                                                            va_arg2= ToType(CEMPTY.Id()).EmptyList()
                                                            /* Let:30 */{ 
                                                              var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                                              /* noccur = 5 */
                                                              _CL_obj.Selector = ToProperty(Core.C__plus.Id())
                                                              _CL_obj.Args = MakeConstantList(n.Id(),MakeInteger(1).Id())
                                                              v_bag_arg = _CL_obj.Id()
                                                              /* Let-30 */} 
                                                            va_arg2.AddFast(v_bag_arg)
                                                            va_arg2.AddFast(m.Id())/* Construct-29 */} 
                                                          /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                                                          va_arg1.Args = va_arg2
                                                          /* update-28 */} 
                                                        va_arg2 = _CL_obj.Id()
                                                        /* Let-27 */} 
                                                      /* ---------- now we compile update iClaire/set_arg(va_arg1) := va_arg2 ------- */
                                                      va_arg1.SetArg = va_arg2
                                                      /* update-26 */} 
                                                    /* update:26 */{ 
                                                      var va_arg1 *Language.Iteration  
                                                      var va_arg2 *ClaireAny  
                                                      va_arg1 = Language.To_Iteration(_CL_obj.Id())
                                                      /* Let:27 */{ 
                                                        var _CL_obj *Language.If   = Language.To_If(new(Language.If).Is(Language.C_If))
                                                        /* noccur = 63 */
                                                        /* update:28 */{ 
                                                          var va_arg1 *Language.If  
                                                          var va_arg2 *ClaireAny  
                                                          va_arg1 = _CL_obj
                                                          /* Let:29 */{ 
                                                            var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                                            /* noccur = 11 */
                                                            _CL_obj.Selector = ToProperty(f)
                                                            /* update:30 */{ 
                                                              var va_arg1 *Language.Call  
                                                              var va_arg2 *ClaireList  
                                                              va_arg1 = _CL_obj
                                                              /* Construct:31 */{ 
                                                                var v_bag_arg *ClaireAny  
                                                                va_arg2= ToType(CEMPTY.Id()).EmptyList()
                                                                /* Let:32 */{ 
                                                                  var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                                                  /* noccur = 5 */
                                                                  _CL_obj.Selector = C_nth
                                                                  _CL_obj.Args = MakeConstantList(l,p.Id())
                                                                  v_bag_arg = _CL_obj.Id()
                                                                  /* Let-32 */} 
                                                                va_arg2.AddFast(v_bag_arg)
                                                                va_arg2.AddFast(x.Id())/* Construct-31 */} 
                                                              /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                                                              va_arg1.Args = va_arg2
                                                              /* update-30 */} 
                                                            va_arg2 = _CL_obj.Id()
                                                            /* Let-29 */} 
                                                          /* ---------- now we compile update iClaire/test(va_arg1) := va_arg2 ------- */
                                                          va_arg1.Test = va_arg2
                                                          /* update-28 */} 
                                                        /* update:28 */{ 
                                                          var va_arg1 *Language.If  
                                                          var va_arg2 *ClaireAny  
                                                          va_arg1 = _CL_obj
                                                          /* Let:29 */{ 
                                                            var _CL_obj *Language.Do   = Language.To_Do(new(Language.Do).Is(Language.C_Do))
                                                            /* noccur = 45 */
                                                            /* update:30 */{ 
                                                              var va_arg1 *Language.Do  
                                                              var va_arg2 *ClaireList  
                                                              va_arg1 = _CL_obj
                                                              /* Construct:31 */{ 
                                                                var v_bag_arg *ClaireAny  
                                                                va_arg2= ToType(CEMPTY.Id()).EmptyList()
                                                                /* Let:32 */{ 
                                                                  var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                                                  /* noccur = 11 */
                                                                  _CL_obj.Selector = C_nth_equal
                                                                  /* update:33 */{ 
                                                                    var va_arg1 *Language.Call  
                                                                    var va_arg2 *ClaireList  
                                                                    va_arg1 = _CL_obj
                                                                    /* Construct:34 */{ 
                                                                      var v_bag_arg *ClaireAny  
                                                                      va_arg2= ToType(CEMPTY.Id()).EmptyList()
                                                                      va_arg2.AddFast(l)
                                                                      va_arg2.AddFast(n.Id())
                                                                      /* Let:35 */{ 
                                                                        var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                                                        /* noccur = 5 */
                                                                        _CL_obj.Selector = C_nth
                                                                        _CL_obj.Args = MakeConstantList(l,p.Id())
                                                                        v_bag_arg = _CL_obj.Id()
                                                                        /* Let-35 */} 
                                                                      va_arg2.AddFast(v_bag_arg)/* Construct-34 */} 
                                                                    /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                                                                    va_arg1.Args = va_arg2
                                                                    /* update-33 */} 
                                                                  v_bag_arg = _CL_obj.Id()
                                                                  /* Let-32 */} 
                                                                va_arg2.AddFast(v_bag_arg)
                                                                /* Let:32 */{ 
                                                                  var g1369UU *Language.Call  
                                                                  /* noccur = 1 */
                                                                  /* Let:33 */{ 
                                                                    var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                                                    /* noccur = 5 */
                                                                    _CL_obj.Selector = ToProperty(Core.C__plus.Id())
                                                                    _CL_obj.Args = MakeConstantList(n.Id(),MakeInteger(1).Id())
                                                                    g1369UU = _CL_obj
                                                                    /* Let-33 */} 
                                                                  v_bag_arg = Language.C_Assign.Make(n.Id(),g1369UU.Id())
                                                                  /* Let-32 */} 
                                                                va_arg2.AddFast(v_bag_arg)
                                                                /* Let:32 */{ 
                                                                  var _CL_obj *Language.If   = Language.To_If(new(Language.If).Is(Language.C_If))
                                                                  /* noccur = 23 */
                                                                  /* update:33 */{ 
                                                                    var va_arg1 *Language.If  
                                                                    var va_arg2 *ClaireAny  
                                                                    va_arg1 = _CL_obj
                                                                    /* Let:34 */{ 
                                                                      var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                                                      /* noccur = 5 */
                                                                      _CL_obj.Selector = ToProperty(C__sup.Id())
                                                                      _CL_obj.Args = MakeConstantList(p.Id(),n.Id())
                                                                      va_arg2 = _CL_obj.Id()
                                                                      /* Let-34 */} 
                                                                    /* ---------- now we compile update iClaire/test(va_arg1) := va_arg2 ------- */
                                                                    va_arg1.Test = va_arg2
                                                                    /* update-33 */} 
                                                                  /* update:33 */{ 
                                                                    var va_arg1 *Language.If  
                                                                    var va_arg2 *ClaireAny  
                                                                    va_arg1 = _CL_obj
                                                                    /* Let:34 */{ 
                                                                      var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                                                      /* noccur = 11 */
                                                                      _CL_obj.Selector = C_nth_equal
                                                                      /* update:35 */{ 
                                                                        var va_arg1 *Language.Call  
                                                                        var va_arg2 *ClaireList  
                                                                        va_arg1 = _CL_obj
                                                                        /* Construct:36 */{ 
                                                                          var v_bag_arg *ClaireAny  
                                                                          va_arg2= ToType(CEMPTY.Id()).EmptyList()
                                                                          va_arg2.AddFast(l)
                                                                          va_arg2.AddFast(p.Id())
                                                                          /* Let:37 */{ 
                                                                            var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                                                            /* noccur = 5 */
                                                                            _CL_obj.Selector = C_nth
                                                                            _CL_obj.Args = MakeConstantList(l,n.Id())
                                                                            v_bag_arg = _CL_obj.Id()
                                                                            /* Let-37 */} 
                                                                          va_arg2.AddFast(v_bag_arg)/* Construct-36 */} 
                                                                        /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                                                                        va_arg1.Args = va_arg2
                                                                        /* update-35 */} 
                                                                      va_arg2 = _CL_obj.Id()
                                                                      /* Let-34 */} 
                                                                    /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                                                                    va_arg1.Arg = va_arg2
                                                                    /* update-33 */} 
                                                                  v_bag_arg = _CL_obj.Id()
                                                                  /* Let-32 */} 
                                                                va_arg2.AddFast(v_bag_arg)/* Construct-31 */} 
                                                              /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                                                              va_arg1.Args = va_arg2
                                                              /* update-30 */} 
                                                            va_arg2 = _CL_obj.Id()
                                                            /* Let-29 */} 
                                                          /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                                                          va_arg1.Arg = va_arg2
                                                          /* update-28 */} 
                                                        va_arg2 = _CL_obj.Id()
                                                        /* Let-27 */} 
                                                      /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                                                      va_arg1.Arg = va_arg2
                                                      /* update-26 */} 
                                                    v_bag_arg = _CL_obj.Id()
                                                    /* Let-25 */} 
                                                  va_arg2.AddFast(v_bag_arg)
                                                  /* Let:25 */{ 
                                                    var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                                    /* noccur = 5 */
                                                    _CL_obj.Selector = C_nth_equal
                                                    _CL_obj.Args = MakeConstantList(l,n.Id(),x.Id())
                                                    v_bag_arg = _CL_obj.Id()
                                                    /* Let-25 */} 
                                                  va_arg2.AddFast(v_bag_arg)
                                                  /* Let:25 */{ 
                                                    var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                                    /* noccur = 11 */
                                                    _CL_obj.Selector = self.Arg.Selector
                                                    /* update:26 */{ 
                                                      var va_arg1 *Language.Call  
                                                      var va_arg2 *ClaireList  
                                                      va_arg1 = _CL_obj
                                                      /* Construct:27 */{ 
                                                        var v_bag_arg *ClaireAny  
                                                        va_arg2= ToType(CEMPTY.Id()).EmptyList()
                                                        va_arg2.AddFast(q.Id())
                                                        /* Let:28 */{ 
                                                          var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                                          /* noccur = 5 */
                                                          _CL_obj.Selector = ToProperty(C__dash.Id())
                                                          _CL_obj.Args = MakeConstantList(n.Id(),MakeInteger(1).Id())
                                                          v_bag_arg = _CL_obj.Id()
                                                          /* Let-28 */} 
                                                        va_arg2.AddFast(v_bag_arg)
                                                        va_arg2.AddFast(l)/* Construct-27 */} 
                                                      /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                                                      va_arg1.Args = va_arg2
                                                      /* update-26 */} 
                                                    v_bag_arg = _CL_obj.Id()
                                                    /* Let-25 */} 
                                                  va_arg2.AddFast(v_bag_arg)
                                                  /* Let:25 */{ 
                                                    var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                                    /* noccur = 11 */
                                                    _CL_obj.Selector = self.Arg.Selector
                                                    /* update:26 */{ 
                                                      var va_arg1 *Language.Call  
                                                      var va_arg2 *ClaireList  
                                                      va_arg1 = _CL_obj
                                                      /* Construct:27 */{ 
                                                        var v_bag_arg *ClaireAny  
                                                        va_arg2= ToType(CEMPTY.Id()).EmptyList()
                                                        /* Let:28 */{ 
                                                          var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                                          /* noccur = 5 */
                                                          _CL_obj.Selector = ToProperty(Core.C__plus.Id())
                                                          _CL_obj.Args = MakeConstantList(n.Id(),MakeInteger(1).Id())
                                                          v_bag_arg = _CL_obj.Id()
                                                          /* Let-28 */} 
                                                        va_arg2.AddFast(v_bag_arg)
                                                        va_arg2.AddFast(m.Id())
                                                        va_arg2.AddFast(l)/* Construct-27 */} 
                                                      /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                                                      va_arg1.Args = va_arg2
                                                      /* update-26 */} 
                                                    v_bag_arg = _CL_obj.Id()
                                                    /* Let-25 */} 
                                                  va_arg2.AddFast(v_bag_arg)/* Construct-24 */} 
                                                /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                                                va_arg1.Args = va_arg2
                                                /* update-23 */} 
                                              va_arg2 = _CL_obj.Id()
                                              /* Let-22 */} 
                                            /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                                            va_arg1.Arg = va_arg2
                                            /* update-21 */} 
                                          va_arg2 = _CL_obj.Id()
                                          /* Let-20 */} 
                                        /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                                        va_arg1.Arg = va_arg2
                                        /* update-19 */} 
                                      g1367UU = _CL_obj
                                      /* Let-18 */} 
                                    va_arg2 = Language.C_If.Make(g1365UU.Id(),g1366UU.Id(),g1367UU.Id())
                                    /* Let-17 */} 
                                  /* Let-16 */} 
                                /* Let-15 */} 
                              /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                              va_arg1.Arg = va_arg2
                              /* update-14 */} 
                            va_arg2 = _CL_obj.Id()
                            /* Let-13 */} 
                          /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                          va_arg1.Arg = va_arg2
                          /* update-12 */} 
                        _Zbd = _CL_obj
                        /* Let-11 */} 
                      /* Let:11 */{ 
                        var def2 *Language.Defmethod  
                        /* noccur = 2 */
                        /* Let:12 */{ 
                          var _CL_obj *Language.Defmethod   = Language.To_Defmethod(new(Language.Defmethod).Is(Language.C_Defmethod))
                          /* noccur = 14 */
                          /* update:13 */{ 
                            var va_arg1 *Language.Defmethod  
                            var va_arg2 *Language.Call  
                            va_arg1 = _CL_obj
                            /* Let:14 */{ 
                              var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                              /* noccur = 5 */
                              _CL_obj.Selector = self.Arg.Selector
                              _CL_obj.Args = MakeConstantList(n.Id(),m.Id(),l)
                              va_arg2 = _CL_obj
                              /* Let-14 */} 
                            /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                            va_arg1.Arg = va_arg2
                            /* update-13 */} 
                          _CL_obj.Inline_ask = CFALSE
                          _CL_obj.SetArg = self.SetArg
                          _CL_obj.Body = _Zbd.Id()
                          def2 = _CL_obj
                          /* Let-12 */} 
                        Core.F_tformat_string(MakeString("---- note: quick sort optimisation for ~S ---- \n"),2,MakeConstantList(self.Arg.Selector.Id()))
                        Result = EVAL(def2.Id())
                        /* ERROR PROTECTION INSERTED (Result-Result) */
                        if !ErrorIn(Result) {
                        /* Let:12 */{ 
                          var _CL_obj *Language.Do   = Language.To_Do(new(Language.Do).Is(Language.C_Do))
                          /* noccur = 3 */
                          /* update:13 */{ 
                            var va_arg1 *Language.Do  
                            var va_arg2 *ClaireList  
                            va_arg1 = _CL_obj
                            var va_arg2_try137014 EID 
                            /* Construct:14 */{ 
                              var v_bag_arg *ClaireAny  
                              va_arg2_try137014= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                              var v_bag_arg_try137115 EID 
                              v_bag_arg_try137115 = Core.F_CALL(C_c_code,ARGS(EID{def1.Id(),0}))
                              /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try137014) */
                              if ErrorIn(v_bag_arg_try137115) {va_arg2_try137014 = v_bag_arg_try137115
                              } else {
                              v_bag_arg = ANY(v_bag_arg_try137115)
                              ToList(OBJ(va_arg2_try137014)).AddFast(v_bag_arg)
                              var v_bag_arg_try137215 EID 
                              v_bag_arg_try137215 = Core.F_CALL(C_c_code,ARGS(EID{def2.Id(),0}))
                              /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try137014) */
                              if ErrorIn(v_bag_arg_try137215) {va_arg2_try137014 = v_bag_arg_try137215
                              } else {
                              v_bag_arg = ANY(v_bag_arg_try137215)
                              ToList(OBJ(va_arg2_try137014)).AddFast(v_bag_arg)}}
                              /* Construct-14 */} 
                            /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                            if ErrorIn(va_arg2_try137014) {Result = va_arg2_try137014
                            } else {
                            va_arg2 = ToList(OBJ(va_arg2_try137014))
                            /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                            va_arg1.Args = va_arg2
                            Result = EID{va_arg2.Id(),0}
                            }
                            /* update-13 */} 
                          /* ERROR PROTECTION INSERTED (Result-Result) */
                          if !ErrorIn(Result) {
                          Result = EID{_CL_obj.Id(),0}
                          }
                          /* Let-12 */} 
                        }
                        /* Let-11 */} 
                      /* Let-10 */} 
                    /* Let-9 */} 
                  /* Let-8 */} 
                /* Let-7 */} 
              /* Let-6 */} 
            /* Let-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: sort_code @ Defmethod (throw: true) 
func E_Optimize_sort_code_Defmethod (self EID,lv EID) EID { 
    return /*(sm for sort_code @ Defmethod= EID)*/ F_Optimize_sort_code_Defmethod(Language.To_Defmethod(OBJ(self)),ToList(OBJ(lv)) )} 
  
// new: we deal with floats --------------------------------------
// create a restriction so that OPT is happy
/* {1} OPT.The go function for: add_method(p:property,ls:list,rg:type,st:integer,f1:function,f2:function) [] */
func F_add_method_property2 (p *ClaireProperty ,ls *ClaireList ,rg *ClaireType ,st int,f1 *ClaireFunction ,f2 *ClaireFunction ) *ClaireMethod  { 
    // use function body compiling 
return  F_add_method_property(p,ls,rg,st,f1)
    } 
  
// The EID go function for: add_method @ list<type_expression>(property, list, type, integer, function, function) (throw: false) 
func E_add_method_property2 (p EID,ls EID,rg EID,st EID,f1 EID,f2 EID) EID { 
    return EID{/*(sm for add_method @ list<type_expression>(property, list, type, integer, function, function)= method)*/ F_add_method_property2(ToProperty(OBJ(p)),
      ToList(OBJ(ls)),
      ToType(OBJ(rg)),
      INT(st),
      ToFunction(OBJ(f1)),
      ToFunction(OBJ(f2)) ).Id(),0}} 
  
/* {1} OPT.The go function for: add_method!(m:method,ls:list,rg:any,stat:any,fu:function) [] */
func F_Optimize_add_method_I_method (m *ClaireMethod ,ls *ClaireList ,rg *ClaireAny ,stat *ClaireAny ,fu *ClaireFunction ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var _Zc *Language.CallMethod  
      /* noccur = 4 */
      var _Zc_try13743 EID 
      /* Let:3 */{ 
        var _CL_obj *Language.CallMethod   = Language.To_CallMethod(new(Language.CallMethod).Is(Language.C_Call_method))
        /* noccur = 5 */
        _CL_obj.Arg = ToMethod(Core.F__at_property1(C_add_method,C_property).Id())
        /* update:4 */{ 
          var va_arg1 *Language.CallMethod  
          var va_arg2 *ClaireList  
          va_arg1 = _CL_obj
          var va_arg2_try13755 EID 
          /* Construct:5 */{ 
            var v_bag_arg *ClaireAny  
            va_arg2_try13755= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
            var v_bag_arg_try13766 EID 
            v_bag_arg_try13766 = Core.F_CALL(C_c_code,ARGS(EID{m.Selector.Id(),0},EID{C_property.Id(),0}))
            /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try13755) */
            if ErrorIn(v_bag_arg_try13766) {va_arg2_try13755 = v_bag_arg_try13766
            } else {
            v_bag_arg = ANY(v_bag_arg_try13766)
            ToList(OBJ(va_arg2_try13755)).AddFast(v_bag_arg)
            var v_bag_arg_try13776 EID 
            v_bag_arg_try13776 = Core.F_CALL(C_c_code,ARGS(EID{ls.Id(),0},EID{C_list.Id(),0}))
            /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try13755) */
            if ErrorIn(v_bag_arg_try13776) {va_arg2_try13755 = v_bag_arg_try13776
            } else {
            v_bag_arg = ANY(v_bag_arg_try13776)
            ToList(OBJ(va_arg2_try13755)).AddFast(v_bag_arg)
            var v_bag_arg_try13786 EID 
            v_bag_arg_try13786 = Core.F_CALL(C_c_code,ARGS(rg.ToEID(),EID{C_type.Id(),0}))
            /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try13755) */
            if ErrorIn(v_bag_arg_try13786) {va_arg2_try13755 = v_bag_arg_try13786
            } else {
            v_bag_arg = ANY(v_bag_arg_try13786)
            ToList(OBJ(va_arg2_try13755)).AddFast(v_bag_arg)
            ToList(OBJ(va_arg2_try13755)).AddFast(stat)
            ToList(OBJ(va_arg2_try13755)).AddFast(fu.Id())}}}
            /* Construct-5 */} 
          /* ERROR PROTECTION INSERTED (va_arg2-_Zc_try13743) */
          if ErrorIn(va_arg2_try13755) {_Zc_try13743 = va_arg2_try13755
          } else {
          va_arg2 = ToList(OBJ(va_arg2_try13755))
          /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
          va_arg1.Args = va_arg2
          _Zc_try13743 = EID{va_arg2.Id(),0}
          }
          /* update-4 */} 
        /* ERROR PROTECTION INSERTED (_Zc_try13743-_Zc_try13743) */
        if !ErrorIn(_Zc_try13743) {
        _Zc_try13743 = EID{_CL_obj.Id(),0}
        }
        /* Let-3 */} 
      /* ERROR PROTECTION INSERTED (_Zc-Result) */
      if ErrorIn(_Zc_try13743) {Result = _Zc_try13743
      } else {
      _Zc = Language.To_CallMethod(OBJ(_Zc_try13743))
      if ((m.Range.Id() == C_float.Id()) || 
          ((m.Domain.Memq(C_float.Id()) == CTRUE) || 
            (C_tuple.Id() == m.Range.Isa.Id()))) /* If:3 */{ 
        _Zc.Args = _Zc.Args.AddFast(F_make_function_string(F_append_string(F_string_I_function(fu),MakeString("_"))).Id())
        /* If-3 */} 
      Result = EID{_Zc.Id(),0}
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: add_method! @ method (throw: true) 
func E_Optimize_add_method_I_method (m EID,ls EID,rg EID,stat EID,fu EID) EID { 
    return /*(sm for add_method! @ method= EID)*/ F_Optimize_add_method_I_method(ToMethod(OBJ(m)),
      ToList(OBJ(ls)),
      ANY(rg),
      ANY(stat),
      ToFunction(OBJ(fu)) )} 
  
// this signature extraction is more subtle since it also builds an external
// list. (l1 is the domain (may use global variables), l2 is the "pure"
// list of patterns)
/* {1} OPT.The go function for: extract_signature!(l:list) [] */
func F_Optimize_extract_signature_I_list (l *ClaireList ) EID { 
    var Result EID 
    Language.C_LDEF.Value = ToType(C_any.Id()).EmptyList().Id()
    /* Let:2 */{ 
      var n int  = 0
      /* noccur = 3 */
      /* Let:3 */{ 
        var l1 *ClaireList   = ToType(C_type_expression.Id()).EmptyList()
        /* noccur = 3 */
        /* Let:4 */{ 
          var l2 *ClaireList  
          /* noccur = 1 */
          var l2_try13795 EID 
          /* Iteration:5 */{ 
            var v_list5 *ClaireList  
            var v *ClaireAny  
            var v_local5 *ClaireAny  
            v_list5 = l
            l2_try13795 = EID{CreateList(ToType(C_any.Id()),v_list5.Length()).Id(),0}
            for CLcount := 0; CLcount < v_list5.Length(); CLcount++{ 
              v = v_list5.At(CLcount)
              var v_local5_try13807 EID 
              /* Let:7 */{ 
                var p *ClaireAny  
                /* noccur = 3 */
                var p_try13818 EID 
                p_try13818 = Language.F_extract_pattern_any(To_Variable(v).Range.Id(),MakeConstantList(MakeInteger(n).Id()))
                /* ERROR PROTECTION INSERTED (p-v_local5_try13807) */
                if ErrorIn(p_try13818) {v_local5_try13807 = p_try13818
                } else {
                p = ANY(p_try13818)
                n = (n+1)
                /* Let:8 */{ 
                  var g1382UU *ClaireAny  
                  /* noccur = 1 */
                  if (To_Variable(v).Range.Isa.IsIn(Core.C_global_variable) == CTRUE) /* If:9 */{ 
                    g1382UU = To_Variable(v).Range.Id()
                    } else {
                    g1382UU = p
                    /* If-9 */} 
                  l1 = l1.AddFast(g1382UU)
                  /* Let-8 */} 
                To_Variable(v).Range = Language.F_type_I_any(p)
                v_local5_try13807 = p.ToEID()
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (v_local5-l2_try13795) */
              if ErrorIn(v_local5_try13807) {l2_try13795 = v_local5_try13807
              l2_try13795 = v_local5_try13807
              break
              } else {
              v_local5 = ANY(v_local5_try13807)
              ToList(OBJ(l2_try13795)).PutAt(CLcount,v_local5)
              } 
            }
            /* Iteration-5 */} 
          /* ERROR PROTECTION INSERTED (l2-Result) */
          if ErrorIn(l2_try13795) {Result = l2_try13795
          } else {
          l2 = ToList(OBJ(l2_try13795))
          Result = EID{MakeConstantList(l1.Id(),l2.Id()).Id(),0}
          }
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: extract_signature! @ list (throw: true) 
func E_Optimize_extract_signature_I_list (l EID) EID { 
    return /*(sm for extract_signature! @ list= EID)*/ F_Optimize_extract_signature_I_list(ToList(OBJ(l)) )} 
  
// check signature equality
/* {1} OPT.The go function for: =sig?(x:list,y:list) [] */
func F_Optimize__equalsig_ask_list (x *ClaireList ,y *ClaireList ) *ClaireBoolean  { 
    // use function body compiling 
return  MakeBoolean((Core.F_tmatch_ask_list(x,y) == CTRUE) && (Core.F_tmatch_ask_list(y,x) == CTRUE))
    } 
  
// The EID go function for: =sig? @ list (throw: false) 
func E_Optimize__equalsig_ask_list (x EID,y EID) EID { 
    return EID{/*(sm for =sig? @ list= boolean)*/ F_Optimize__equalsig_ask_list(ToList(OBJ(x)),ToList(OBJ(y)) ).Id(),0}} 
  
// creates a name for a restriction from the full domain
// Note that we suppose that a new restriction is not allowed to be inserted
// in a list of restrictions when the property is closed.
//
/* {1} OPT.The go function for: Compile/function_name(p:property,l:list,x:any) [] */
func F_Compile_function_name_property1 (p *ClaireProperty ,l *ClaireList ,x *ClaireAny ) *ClaireString  { 
    // procedure body with s =  
var Result *ClaireString  
    if (C_function.Id() == x.Isa.Id()) /* If:2 */{ 
      Result = F_string_I_function(ToFunction(x))
      } else {
      /* Let:3 */{ 
        var n int  = 0
        /* noccur = 4 */
        /* Let:4 */{ 
          var m int  = 0
          /* noccur = 2 */
          /* Let:5 */{ 
            var md *ClaireModule   = p.Name.Module_I()
            /* noccur = 1 */
            /* Let:6 */{ 
              var c *ClaireClass   = ToTypeExpression(l.At(1-1)).Class_I()
              /* noccur = 2 */
              /* Let:7 */{ 
                var r *ClaireString   = F_append_string(F_append_string(p.Name.String_I(),MakeString("_")),c.Name.String_I())
                /* noccur = 7 */
                if ((C_compiler.Naming == 0) && 
                    (p.Id() != Core.C_main.Id())) /* If:8 */{ 
                  r = F_append_string(F_append_string(md.Name.String_I(),MakeString("_")),r)
                  /* If-8 */} 
                /* For:8 */{ 
                  var r *ClaireAny  
                  _ = r
                  for _,r = range(p.Restrictions.ValuesO())/* loop:9 */{ 
                    if (c.Id() == Core.F_domain_I_restriction(ToRestriction(r)).Id()) /* If:10 */{ 
                      n = (n+1)
                      /* If-10 */} 
                    if (F_Optimize__equalsig_ask_list(l,ToRestriction(r).Domain) == CTRUE) /* If:10 */{ 
                      m = n
                      /* If-10 */} 
                    /* loop-9 */} 
                  /* For-8 */} 
                if (n <= 1) /* If:8 */{ 
                  r = r
                  } else {
                  r = F_append_string(r,F_string_I_integer(m))
                  /* If-8 */} 
                if ((F_Optimize_stable_ask_relation(ToRelation(p.Id())) == CTRUE) || 
                    (p.Id() == Core.C_main.Id())) /* If:8 */{ 
                  Result = r
                  } else {
                  Result = F_append_string(F_append_string(r,MakeString("_")),ClEnv.Module_I.Name.String_I())
                  /* If-8 */} 
                /* Let-7 */} 
              /* Let-6 */} 
            /* Let-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* If-2 */} 
    return Result} 
  
// The EID go function for: Compile/function_name @ list<type_expression>(property, list, any) (throw: false) 
func E_Compile_function_name_property1 (p EID,l EID,x EID) EID { 
    return EID{/*(sm for Compile/function_name @ list<type_expression>(property, list, any)= string)*/ F_Compile_function_name_property1(ToProperty(OBJ(p)),ToList(OBJ(l)),ANY(x) ).Id(),0}} 
  
// this compiles a lambda into a C method with name oself.
// the use_new flag will be raised if a new object is created inside the
// function.
// m is either the associated method,or the expected range
//
/* {1} OPT.The go function for: Compile/compile_lambda(self:string,l:lambda,m:any) [] */
func F_Compile_compile_lambda_string (self *ClaireString ,l *ClaireLambda ,m *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var x int  = C_compiler.Safety
      /* noccur = 1 */
      /* Let:3 */{ 
        var y *ClaireLambda   = l
        /* noccur = 0 */
        _ = y
        Core.F_tformat_string(MakeString("---- Compiling ~A,\n"),3,MakeConstantList((self).Id()))
        if (C_method.Id() == m.Isa.Id()) /* If:4 */{ 
          /* Let:5 */{ 
            var g1383 *ClaireMethod   = ToMethod(m)
            /* noccur = 1 */
            C_OPT.InMethod = g1383.Id()
            /* Let-5 */} 
          /* If-4 */} 
        if (C_OPT.LoopIndex > 0) /* If:4 */{ 
          C_OPT.LoopIndex = 0
          /* If-4 */} 
        C_OPT.MaxVars = 0
        if (ToBoolean(C_OPT.Unsure.Contain_ask(m).Id()) == CTRUE) /* If:4 */{ 
          C_compiler.Safety = 1
          /* If-4 */} 
        Result = Core.F_CALL(C_Compile_make_c_function,ARGS(EID{l.Id(),0},EID{(self).Id(),0},m.ToEID()))
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        C_OPT.InMethod = CNULL
        C_compiler.Safety = x
        Result = EID{CTRUE.Id(),0}
        }
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: Compile/compile_lambda @ string (throw: true) 
func E_Compile_compile_lambda_string (self EID,l EID,m EID) EID { 
    return /*(sm for Compile/compile_lambda @ string= EID)*/ F_Compile_compile_lambda_string(ToString(OBJ(self)),ToLambda(OBJ(l)),ANY(m) )} 
  
// how to compile an table definition
/* {1} OPT.The go function for: c_code(self:Defarray) [] */
func F_c_code_Defarray (self *Language.Defarray ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var a *ClaireList   = self.Arg.Args
      /* noccur = 14 */
      /* Let:3 */{ 
        var _Za *ClaireAny  
        /* noccur = 13 */
        var _Za_try13934 EID 
        /* Let:4 */{ 
          var g1394UU *ClaireSymbol  
          /* noccur = 1 */
          var g1394UU_try13955 EID 
          g1394UU_try13955 = Language.F_extract_symbol_any(a.At(1-1))
          /* ERROR PROTECTION INSERTED (g1394UU-_Za_try13934) */
          if ErrorIn(g1394UU_try13955) {_Za_try13934 = g1394UU_try13955
          } else {
          g1394UU = ToSymbol(OBJ(g1394UU_try13955))
          _Za_try13934 = g1394UU.Get().ToEID()
          }
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (_Za-Result) */
        if ErrorIn(_Za_try13934) {Result = _Za_try13934
        } else {
        _Za = ANY(_Za_try13934)
        /* Let:4 */{ 
          var _Zv *ClaireTable  
          /* noccur = 11 */
          var _Zv_try13965 EID 
          if (C_table.Id() == _Za.Isa.Id()) /* If:5 */{ 
            /* Let:6 */{ 
              var g1385 *ClaireTable   = ToTable(_Za)
              /* noccur = 1 */
              _Zv_try13965 = EID{g1385.Id(),0}
              /* Let-6 */} 
            } else {
            _Zv_try13965 = ToException(Core.C_general_error.Make(MakeString("[internal] the table ~S is unknown").Id(),MakeConstantList(a.At(1-1)).Id())).Close()
            /* If-5 */} 
          /* ERROR PROTECTION INSERTED (_Zv-Result) */
          if ErrorIn(_Zv_try13965) {Result = _Zv_try13965
          } else {
          _Zv = ToTable(OBJ(_Zv_try13965))
          /* Let:5 */{ 
            var s *ClaireAny   = ANY(Core.F_CALL(C_domain,ARGS(_Za.ToEID())))
            /* noccur = 5 */
            /* Let:6 */{ 
              var e *ClaireAny  
              /* noccur = 5 */
              var e_try13977 EID 
              /* Let:7 */{ 
                var l *ClaireList  
                /* noccur = 3 */
                var l_try13988 EID 
                l_try13988 = a.Cdr()
                /* ERROR PROTECTION INSERTED (l-e_try13977) */
                if ErrorIn(l_try13988) {e_try13977 = l_try13988
                } else {
                l = ToList(OBJ(l_try13988))
                /* Let:8 */{ 
                  var b *ClaireAny  
                  /* noccur = 2 */
                  var b_try13999 EID 
                  b_try13999 = Language.F_lexical_build_any(self.Body,l,0)
                  /* ERROR PROTECTION INSERTED (b-e_try13977) */
                  if ErrorIn(b_try13999) {e_try13977 = b_try13999
                  } else {
                  b = ANY(b_try13999)
                  var g1400I *ClaireBoolean  
                  /* Let:9 */{ 
                    var g1401UU *ClaireAny  
                    /* noccur = 1 */
                    /* For:10 */{ 
                      var va *ClaireAny  
                      _ = va
                      g1401UU= CFALSE.Id()
                      var va_support *ClaireList  
                      va_support = l
                      va_len := va_support.Length()
                      for i_it := 0; i_it < va_len; i_it++ { 
                        va = va_support.At(i_it)
                        if (Language.F_occurrence_any(b,To_Variable(va)) > 0) /* If:12 */{ 
                           /*v = g1401UU, s =any*/
g1401UU = CTRUE.Id()
                          break
                          /* If-12 */} 
                        /* loop-11 */} 
                      /* For-10 */} 
                    g1400I = F_boolean_I_any(g1401UU)
                    /* Let-9 */} 
                  if (g1400I == CTRUE) /* If:9 */{ 
                    e_try13977 = Language.F_lambda_I_list(l,b)
                    } else {
                    e_try13977 = self.Body.ToEID()
                    /* If-9 */} 
                  }
                  /* Let-8 */} 
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (e-Result) */
              if ErrorIn(e_try13977) {Result = e_try13977
              } else {
              e = ANY(e_try13977)
              /* Let:7 */{ 
                var d *ClaireAny  
                /* noccur = 3 */
                if (e.Isa.IsIn(C_lambda) == CTRUE) /* If:8 */{ 
                  d = CNULL
                  } else {
                  d = self.Body
                  /* If-8 */} 
                /* Let:8 */{ 
                  var _Zl1 *ClaireList  
                  /* noccur = 1 */
                  if (ToRelation(_Za).Multivalued_ask == CTRUE) /* If:9 */{ 
                    /* Construct:10 */{ 
                      var v_bag_arg *ClaireAny  
                      _Zl1= ToType(C_any.Id()).EmptyList()
                      /* Let:11 */{ 
                        var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                        /* noccur = 5 */
                        _CL_obj.Selector = C_put
                        _CL_obj.Args = MakeConstantList(C_multivalued_ask.Id(),_Zv.Id(),ToRelation(_Za).Multivalued_ask.Id())
                        v_bag_arg = _CL_obj.Id()
                        /* Let-11 */} 
                      _Zl1.AddFast(v_bag_arg)/* Construct-10 */} 
                    } else {
                    _Zl1 = ToType(C_any.Id()).EmptyList()
                    /* If-9 */} 
                  /* Let:9 */{ 
                    var _Zl2 *ClaireList  
                    /* noccur = 9 */
                    /* Construct:10 */{ 
                      var v_bag_arg *ClaireAny  
                      _Zl2= ToType(C_any.Id()).EmptyList()
                      /* Let:11 */{ 
                        var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                        /* noccur = 5 */
                        _CL_obj.Selector = C_put
                        _CL_obj.Args = MakeConstantList(C_range.Id(),_Zv.Id(),ANY(Core.F_CALL(C_range,ARGS(_Za.ToEID()))))
                        v_bag_arg = _CL_obj.Id()
                        /* Let-11 */} 
                      _Zl2.AddFast(v_bag_arg)
                      /* Let:11 */{ 
                        var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                        /* noccur = 5 */
                        _CL_obj.Selector = C_put
                        _CL_obj.Args = MakeConstantList(C_params.Id(),_Zv.Id(),ANY(Core.F_CALL(C_params,ARGS(_Za.ToEID()))))
                        v_bag_arg = _CL_obj.Id()
                        /* Let-11 */} 
                      _Zl2.AddFast(v_bag_arg)
                      /* Let:11 */{ 
                        var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                        /* noccur = 5 */
                        _CL_obj.Selector = C_put
                        _CL_obj.Args = MakeConstantList(C_domain.Id(),_Zv.Id(),s)
                        v_bag_arg = _CL_obj.Id()
                        /* Let-11 */} 
                      _Zl2.AddFast(v_bag_arg)/* Construct-10 */} 
                    /* update:10 */{ 
                      var va_arg1 *ClaireVariable  
                      var va_arg2 *ClaireType  
                      va_arg1 = To_Variable(a.At(2-1))
                      var va_arg2_try140211 EID 
                      va_arg2_try140211 = Language.F_extract_type_any(To_Variable(a.At(2-1)).Range.Id())
                      /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                      if ErrorIn(va_arg2_try140211) {Result = va_arg2_try140211
                      } else {
                      va_arg2 = ToType(OBJ(va_arg2_try140211))
                      /* ---------- now we compile update range(va_arg1) := va_arg2 ------- */
                      va_arg1.Range = va_arg2
                      Result = EID{va_arg2.Id(),0}
                      }
                      /* update-10 */} 
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    if (a.Length() == 2) /* If:10 */{ 
                      var _Zl2_try140311 EID 
                      /* Let:11 */{ 
                        var g1404UU *Language.Call  
                        /* noccur = 1 */
                        var g1404UU_try140512 EID 
                        if (s.Isa.IsIn(C_Interval) == CTRUE) /* If:12 */{ 
                          /* Let:13 */{ 
                            var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                            /* noccur = 11 */
                            _CL_obj.Selector = C_put
                            /* update:14 */{ 
                              var va_arg1 *Language.Call  
                              var va_arg2 *ClaireList  
                              va_arg1 = _CL_obj
                              var va_arg2_try140615 EID 
                              /* Construct:15 */{ 
                                var v_bag_arg *ClaireAny  
                                va_arg2_try140615= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                                ToList(OBJ(va_arg2_try140615)).AddFast(C_mClaire_graph.Id())
                                ToList(OBJ(va_arg2_try140615)).AddFast(_Zv.Id())
                                var v_bag_arg_try140716 EID 
                                /* Let:16 */{ 
                                  var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                  /* noccur = 5 */
                                  _CL_obj.Selector = Core.C_make_copy_list
                                  /* update:17 */{ 
                                    var va_arg1 *Language.Call  
                                    var va_arg2 *ClaireList  
                                    va_arg1 = _CL_obj
                                    var va_arg2_try140818 EID 
                                    /* Construct:18 */{ 
                                      var v_bag_arg *ClaireAny  
                                      va_arg2_try140818= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                                      var v_bag_arg_try140919 EID 
                                      v_bag_arg_try140919 = Core.F_CALL(C_size,ARGS(s.ToEID()))
                                      /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try140818) */
                                      if ErrorIn(v_bag_arg_try140919) {va_arg2_try140818 = v_bag_arg_try140919
                                      } else {
                                      v_bag_arg = ANY(v_bag_arg_try140919)
                                      ToList(OBJ(va_arg2_try140818)).AddFast(v_bag_arg)
                                      ToList(OBJ(va_arg2_try140818)).AddFast(d)}
                                      /* Construct-18 */} 
                                    /* ERROR PROTECTION INSERTED (va_arg2-v_bag_arg_try140716) */
                                    if ErrorIn(va_arg2_try140818) {v_bag_arg_try140716 = va_arg2_try140818
                                    } else {
                                    va_arg2 = ToList(OBJ(va_arg2_try140818))
                                    /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                                    va_arg1.Args = va_arg2
                                    v_bag_arg_try140716 = EID{va_arg2.Id(),0}
                                    }
                                    /* update-17 */} 
                                  /* ERROR PROTECTION INSERTED (v_bag_arg_try140716-v_bag_arg_try140716) */
                                  if !ErrorIn(v_bag_arg_try140716) {
                                  v_bag_arg_try140716 = EID{_CL_obj.Id(),0}
                                  }
                                  /* Let-16 */} 
                                /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try140615) */
                                if ErrorIn(v_bag_arg_try140716) {va_arg2_try140615 = v_bag_arg_try140716
                                } else {
                                v_bag_arg = ANY(v_bag_arg_try140716)
                                ToList(OBJ(va_arg2_try140615)).AddFast(v_bag_arg)}
                                /* Construct-15 */} 
                              /* ERROR PROTECTION INSERTED (va_arg2-g1404UU_try140512) */
                              if ErrorIn(va_arg2_try140615) {g1404UU_try140512 = va_arg2_try140615
                              } else {
                              va_arg2 = ToList(OBJ(va_arg2_try140615))
                              /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                              va_arg1.Args = va_arg2
                              g1404UU_try140512 = EID{va_arg2.Id(),0}
                              }
                              /* update-14 */} 
                            /* ERROR PROTECTION INSERTED (g1404UU_try140512-g1404UU_try140512) */
                            if !ErrorIn(g1404UU_try140512) {
                            g1404UU_try140512 = EID{_CL_obj.Id(),0}
                            }
                            /* Let-13 */} 
                          } else {
                          /* Let:13 */{ 
                            var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                            /* noccur = 5 */
                            _CL_obj.Selector = C_graph_init
                            _CL_obj.Args = MakeConstantList(_Zv.Id())
                            g1404UU_try140512 = EID{_CL_obj.Id(),0}
                            /* Let-13 */} 
                          /* If-12 */} 
                        /* ERROR PROTECTION INSERTED (g1404UU-_Zl2_try140311) */
                        if ErrorIn(g1404UU_try140512) {_Zl2_try140311 = g1404UU_try140512
                        } else {
                        g1404UU = Language.To_Call(OBJ(g1404UU_try140512))
                        _Zl2_try140311 = EID{_Zl2.AddFast(g1404UU.Id()).Id(),0}
                        }
                        /* Let-11 */} 
                      /* ERROR PROTECTION INSERTED (_Zl2-Result) */
                      if ErrorIn(_Zl2_try140311) {Result = _Zl2_try140311
                      } else {
                      _Zl2 = ToList(OBJ(_Zl2_try140311))
                      Result = EID{_Zl2.Id(),0}
                      /* Let:11 */{ 
                        var g1410UU *Language.ComplexInstruction  
                        /* noccur = 1 */
                        if (e.Isa.IsIn(C_lambda) == CTRUE) /* If:12 */{ 
                          /* Let:13 */{ 
                            var g1389 *ClaireLambda   = ToLambda(e)
                            /* noccur = 1 */
                            /* Let:14 */{ 
                              var _CL_obj *Language.For   = Language.To_For(new(Language.For).Is(Language.C_For))
                              /* noccur = 11 */
                              _CL_obj.ClaireVar = To_Variable(a.At(2-1))
                              _CL_obj.SetArg = s
                              /* update:15 */{ 
                                var va_arg1 *Language.Iteration  
                                var va_arg2 *ClaireAny  
                                va_arg1 = Language.To_Iteration(_CL_obj.Id())
                                /* Let:16 */{ 
                                  var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                  /* noccur = 5 */
                                  _CL_obj.Selector = C_nth_equal
                                  _CL_obj.Args = MakeConstantList(_Zv.Id(),a.At(2-1),g1389.Body)
                                  va_arg2 = _CL_obj.Id()
                                  /* Let-16 */} 
                                /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                                va_arg1.Arg = va_arg2
                                /* update-15 */} 
                              g1410UU = Language.To_ComplexInstruction(_CL_obj.Id())
                              /* Let-14 */} 
                            /* Let-13 */} 
                          } else {
                          /* Let:13 */{ 
                            var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                            /* noccur = 5 */
                            _CL_obj.Selector = C_put
                            _CL_obj.Args = MakeConstantList(C_default.Id(),_Zv.Id(),d)
                            g1410UU = Language.To_ComplexInstruction(_CL_obj.Id())
                            /* Let-13 */} 
                          /* If-12 */} 
                        _Zl2 = _Zl2.AddFast(g1410UU.Id())
                        /* Let-11 */} 
                      Result = EID{_Zl2.Id(),0}
                      }
                      } else {
                      /* Let:11 */{ 
                        var s2 *ClaireTypeExpression  
                        /* noccur = 2 */
                        var s2_try141112 EID 
                        s2_try141112 = Language.F_extract_type_any(To_Variable(a.At(3-1)).Range.Id())
                        /* ERROR PROTECTION INSERTED (s2-Result) */
                        if ErrorIn(s2_try141112) {Result = s2_try141112
                        } else {
                        s2 = ToTypeExpression(OBJ(s2_try141112))
                        To_Variable(a.At(3-1)).Range = ToType(s2.Id())
                        /* Let:12 */{ 
                          var g1412UU *Language.Call  
                          /* noccur = 1 */
                          /* Let:13 */{ 
                            var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                            /* noccur = 11 */
                            _CL_obj.Selector = C_put
                            /* update:14 */{ 
                              var va_arg1 *Language.Call  
                              var va_arg2 *ClaireList  
                              va_arg1 = _CL_obj
                              /* Construct:15 */{ 
                                var v_bag_arg *ClaireAny  
                                va_arg2= ToType(CEMPTY.Id()).EmptyList()
                                va_arg2.AddFast(C_mClaire_graph.Id())
                                va_arg2.AddFast(_Zv.Id())
                                /* Let:16 */{ 
                                  var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                  /* noccur = 5 */
                                  _CL_obj.Selector = Core.C_make_copy_list
                                  /* update:17 */{ 
                                    var va_arg1 *Language.Call  
                                    var va_arg2 *ClaireList  
                                    va_arg1 = _CL_obj
                                    /* Construct:18 */{ 
                                      var v_bag_arg *ClaireAny  
                                      va_arg2= ToType(CEMPTY.Id()).EmptyList()
                                      va_arg2.AddFast(ANY(Core.F_CALL(C_length,ARGS(ToTable(_Za).Graph.ToEID()))))
                                      if (ANY(Core.F_CALL(C_params,ARGS(_Za.ToEID()))) == C_any.Id()) /* If:19 */{ 
                                        v_bag_arg = CNULL
                                        } else {
                                        v_bag_arg = ANY(Core.F_CALL(C_default,ARGS(_Za.ToEID())))
                                        /* If-19 */} 
                                      va_arg2.AddFast(v_bag_arg)/* Construct-18 */} 
                                    /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                                    va_arg1.Args = va_arg2
                                    /* update-17 */} 
                                  v_bag_arg = _CL_obj.Id()
                                  /* Let-16 */} 
                                va_arg2.AddFast(v_bag_arg)/* Construct-15 */} 
                              /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                              va_arg1.Args = va_arg2
                              /* update-14 */} 
                            g1412UU = _CL_obj
                            /* Let-13 */} 
                          _Zl2 = _Zl2.AddFast(g1412UU.Id())
                          /* Let-12 */} 
                        var _Zl2_try141312 EID 
                        /* Let:12 */{ 
                          var g1414UU *Language.ComplexInstruction  
                          /* noccur = 1 */
                          var g1414UU_try141513 EID 
                          if (e.Isa.IsIn(C_lambda) == CTRUE) /* If:13 */{ 
                            /* Let:14 */{ 
                              var g1391 *ClaireLambda   = ToLambda(e)
                              /* noccur = 1 */
                              /* Let:15 */{ 
                                var _CL_obj *Language.For   = Language.To_For(new(Language.For).Is(Language.C_For))
                                /* noccur = 17 */
                                _CL_obj.ClaireVar = To_Variable(a.At(2-1))
                                /* update:16 */{ 
                                  var va_arg1 *Language.Iteration  
                                  var va_arg2 *ClaireAny  
                                  va_arg1 = Language.To_Iteration(_CL_obj.Id())
                                  var va_arg2_try141617 EID 
                                  va_arg2_try141617 = Core.F_CALL(C_nth,ARGS(s.ToEID(),EID{C__INT,IVAL(1)}))
                                  /* ERROR PROTECTION INSERTED (va_arg2-g1414UU_try141513) */
                                  if ErrorIn(va_arg2_try141617) {g1414UU_try141513 = va_arg2_try141617
                                  } else {
                                  va_arg2 = ANY(va_arg2_try141617)
                                  /* ---------- now we compile update iClaire/set_arg(va_arg1) := va_arg2 ------- */
                                  va_arg1.SetArg = va_arg2
                                  g1414UU_try141513 = va_arg2.ToEID()
                                  }
                                  /* update-16 */} 
                                /* ERROR PROTECTION INSERTED (g1414UU_try141513-g1414UU_try141513) */
                                if !ErrorIn(g1414UU_try141513) {
                                /* update:16 */{ 
                                  var va_arg1 *Language.Iteration  
                                  var va_arg2 *ClaireAny  
                                  va_arg1 = Language.To_Iteration(_CL_obj.Id())
                                  /* Let:17 */{ 
                                    var _CL_obj *Language.For   = Language.To_For(new(Language.For).Is(Language.C_For))
                                    /* noccur = 11 */
                                    _CL_obj.ClaireVar = To_Variable(a.At(3-1))
                                    _CL_obj.SetArg = s2.Id()
                                    /* update:18 */{ 
                                      var va_arg1 *Language.Iteration  
                                      var va_arg2 *ClaireAny  
                                      va_arg1 = Language.To_Iteration(_CL_obj.Id())
                                      /* Let:19 */{ 
                                        var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                        /* noccur = 5 */
                                        _CL_obj.Selector = C_nth_equal
                                        _CL_obj.Args = MakeConstantList(_Zv.Id(),
                                          a.At(2-1),
                                          a.At(3-1),
                                          g1391.Body)
                                        va_arg2 = _CL_obj.Id()
                                        /* Let-19 */} 
                                      /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                                      va_arg1.Arg = va_arg2
                                      /* update-18 */} 
                                    va_arg2 = _CL_obj.Id()
                                    /* Let-17 */} 
                                  /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                                  va_arg1.Arg = va_arg2
                                  /* update-16 */} 
                                g1414UU_try141513 = EID{_CL_obj.Id(),0}
                                }
                                /* Let-15 */} 
                              /* Let-14 */} 
                            } else {
                            /* Let:14 */{ 
                              var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                              /* noccur = 5 */
                              _CL_obj.Selector = C_put
                              _CL_obj.Args = MakeConstantList(C_default.Id(),_Zv.Id(),d)
                              g1414UU_try141513 = EID{_CL_obj.Id(),0}
                              /* Let-14 */} 
                            /* If-13 */} 
                          /* ERROR PROTECTION INSERTED (g1414UU-_Zl2_try141312) */
                          if ErrorIn(g1414UU_try141513) {_Zl2_try141312 = g1414UU_try141513
                          } else {
                          g1414UU = Language.To_ComplexInstruction(OBJ(g1414UU_try141513))
                          _Zl2_try141312 = EID{_Zl2.AddFast(g1414UU.Id()).Id(),0}
                          }
                          /* Let-12 */} 
                        /* ERROR PROTECTION INSERTED (_Zl2-Result) */
                        if ErrorIn(_Zl2_try141312) {Result = _Zl2_try141312
                        } else {
                        _Zl2 = ToList(OBJ(_Zl2_try141312))
                        Result = EID{_Zl2.Id(),0}
                        }
                        }
                        /* Let-11 */} 
                      /* If-10 */} 
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    C_OPT.Objects = C_OPT.Objects.AddFast(_Za)
                    Core.F_CALL(C_Optimize_c_register,ARGS(_Za.ToEID()))
                    /* Let:10 */{ 
                      var g1417UU *Language.Do  
                      /* noccur = 1 */
                      var g1417UU_try141811 EID 
                      /* Let:11 */{ 
                        var _CL_obj *Language.Do   = Language.To_Do(new(Language.Do).Is(Language.C_Do))
                        /* noccur = 9 */
                        /* update:12 */{ 
                          var va_arg1 *Language.Do  
                          var va_arg2 *ClaireList  
                          va_arg1 = _CL_obj
                          var va_arg2_try141913 EID 
                          /* Let:13 */{ 
                            var g1420UU *Language.Call  
                            /* noccur = 1 */
                            /* Let:14 */{ 
                              var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                              /* noccur = 5 */
                              _CL_obj.Selector = C_Compile_object_I
                              _CL_obj.Args = MakeConstantList(ANY(Core.F_CALL(C_name,ARGS(_Za.ToEID()))),C_table.Id())
                              g1420UU = _CL_obj
                              /* Let-14 */} 
                            /* Let:14 */{ 
                              var g1421UU *ClaireList  
                              /* noccur = 1 */
                              var g1421UU_try142215 EID 
                              g1421UU_try142215 = _Zl1.Add_star(_Zl2)
                              /* ERROR PROTECTION INSERTED (g1421UU-va_arg2_try141913) */
                              if ErrorIn(g1421UU_try142215) {va_arg2_try141913 = g1421UU_try142215
                              } else {
                              g1421UU = ToList(OBJ(g1421UU_try142215))
                              va_arg2_try141913 = EID{F_cons_any(g1420UU.Id(),g1421UU).Id(),0}
                              }
                              /* Let-14 */} 
                            /* Let-13 */} 
                          /* ERROR PROTECTION INSERTED (va_arg2-g1417UU_try141811) */
                          if ErrorIn(va_arg2_try141913) {g1417UU_try141811 = va_arg2_try141913
                          } else {
                          va_arg2 = ToList(OBJ(va_arg2_try141913))
                          /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                          va_arg1.Args = va_arg2
                          g1417UU_try141811 = EID{va_arg2.Id(),0}
                          }
                          /* update-12 */} 
                        /* ERROR PROTECTION INSERTED (g1417UU_try141811-g1417UU_try141811) */
                        if !ErrorIn(g1417UU_try141811) {
                        g1417UU_try141811 = EID{_CL_obj.Id(),0}
                        }
                        /* Let-11 */} 
                      /* ERROR PROTECTION INSERTED (g1417UU-Result) */
                      if ErrorIn(g1417UU_try141811) {Result = g1417UU_try141811
                      } else {
                      g1417UU = Language.To_Do(OBJ(g1417UU_try141811))
                      Result = Core.F_CALL(C_c_code,ARGS(EID{g1417UU.Id(),0},EID{C_any.Id(),0}))
                      }
                      /* Let-10 */} 
                    }}
                    /* Let-9 */} 
                  /* Let-8 */} 
                /* Let-7 */} 
              }
              /* Let-6 */} 
            /* Let-5 */} 
          }
          /* Let-4 */} 
        }
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_code @ Defarray (throw: true) 
func E_c_code_Defarray (self EID) EID { 
    return /*(sm for c_code @ Defarray= EID)*/ F_c_code_Defarray(Language.To_Defarray(OBJ(self)) )} 
  
// *********************************************************************
// *     Part 4: Inverse Management (new in v3.0.50)                   *
// *********************************************************************
// this method creates an if_write demon that takes care of the inverse
/* {1} OPT.The go function for: Compile/compute_if_write_inverse(R:relation) [] */
func F_Compile_compute_if_write_inverse_relation (R *ClaireRelation ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var x *ClaireVariable  
      /* noccur = 9 */
      /* Let:3 */{ 
        var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
        /* noccur = 5 */
        _CL_obj.Pname = Core.F_symbol_I_string2(MakeString("XX"))
        _CL_obj.Range = R.Domain
        x = _CL_obj
        /* Let-3 */} 
      /* Let:3 */{ 
        var y *ClaireVariable  
        /* noccur = 8 */
        /* Let:4 */{ 
          var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
          /* noccur = 5 */
          _CL_obj.Pname = Core.F_symbol_I_string2(MakeString("YY"))
          /* update:5 */{ 
            var va_arg1 *ClaireVariable  
            var va_arg2 *ClaireType  
            va_arg1 = _CL_obj
            if (F_boolean_I_any(ANY(Core.F_CALL(C_Optimize_multi_ask,ARGS(EID{R.Id(),0})))) == CTRUE) /* If:6 */{ 
              va_arg2 = Core.F_member_type(R.Range)
              } else {
              va_arg2 = R.Range
              /* If-6 */} 
            /* ---------- now we compile update range(va_arg1) := va_arg2 ------- */
            va_arg1.Range = va_arg2
            /* update-5 */} 
          y = _CL_obj
          /* Let-4 */} 
        /* Let:4 */{ 
          var z *ClaireVariable  
          /* noccur = 4 */
          /* Let:5 */{ 
            var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
            /* noccur = 5 */
            _CL_obj.Pname = Core.F_symbol_I_string2(MakeString("ZZ"))
            _CL_obj.Range = R.Range
            z = _CL_obj
            /* Let-5 */} 
          /* Let:5 */{ 
            var l1 *ClaireList   = ToType(C_any.Id()).EmptyList()
            /* noccur = 10 */
            if (F_boolean_I_any(ANY(Core.F_CALL(C_Optimize_multi_ask,ARGS(EID{R.Id(),0})))) == CTRUE) /* If:6 */{ 
              var l1_try14237 EID 
              /* Construct:7 */{ 
                var v_bag_arg *ClaireAny  
                l1_try14237= EID{ToType(C_any.Id()).EmptyList().Id(),0}
                var v_bag_arg_try14248 EID 
                v_bag_arg_try14248 = Core.F_CALL(C_Optimize_Produce_put,ARGS(EID{R.Id(),0},EID{x.Id(),0},EID{y.Id(),0}))
                /* ERROR PROTECTION INSERTED (v_bag_arg-l1_try14237) */
                if ErrorIn(v_bag_arg_try14248) {l1_try14237 = v_bag_arg_try14248
                } else {
                v_bag_arg = ANY(v_bag_arg_try14248)
                ToList(OBJ(l1_try14237)).AddFast(v_bag_arg)}
                /* Construct-7 */} 
              /* ERROR PROTECTION INSERTED (l1-Result) */
              if ErrorIn(l1_try14237) {Result = l1_try14237
              } else {
              l1 = ToList(OBJ(l1_try14237))
              Result = EID{l1.Id(),0}
              if (R.Inverse.Id() != CNULL) /* If:7 */{ 
                var l1_try14258 EID 
                /* Let:8 */{ 
                  var g1426UU *ClaireAny  
                  /* noccur = 1 */
                  var g1426UU_try14279 EID 
                  g1426UU_try14279 = Core.F_CALL(C_Optimize_Produce_put,ARGS(EID{R.Inverse.Id(),0},EID{y.Id(),0},EID{x.Id(),0}))
                  /* ERROR PROTECTION INSERTED (g1426UU-l1_try14258) */
                  if ErrorIn(g1426UU_try14279) {l1_try14258 = g1426UU_try14279
                  } else {
                  g1426UU = ANY(g1426UU_try14279)
                  l1_try14258 = EID{l1.AddFast(g1426UU).Id(),0}
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (l1-Result) */
                if ErrorIn(l1_try14258) {Result = l1_try14258
                } else {
                l1 = ToList(OBJ(l1_try14258))
                Result = EID{l1.Id(),0}
                }
                } else {
                Result = EID{CFALSE.Id(),0}
                /* If-7 */} 
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              /* update:7 */{ 
                var va_arg1 *ClaireRelation  
                var va_arg2 *ClaireAny  
                va_arg1 = R
                var va_arg2_try14308 EID 
                /* Let:8 */{ 
                  var g1431UU *Language.If  
                  /* noccur = 1 */
                  var g1431UU_try14329 EID 
                  /* Let:9 */{ 
                    var _CL_obj *Language.If   = Language.To_If(new(Language.If).Is(Language.C_If))
                    /* noccur = 21 */
                    /* update:10 */{ 
                      var va_arg1 *Language.If  
                      var va_arg2 *ClaireAny  
                      va_arg1 = _CL_obj
                      var va_arg2_try143311 EID 
                      /* Let:11 */{ 
                        var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                        /* noccur = 11 */
                        _CL_obj.Selector = Core.C_not
                        /* update:12 */{ 
                          var va_arg1 *Language.Call  
                          var va_arg2 *ClaireList  
                          va_arg1 = _CL_obj
                          var va_arg2_try143413 EID 
                          /* Construct:13 */{ 
                            var v_bag_arg *ClaireAny  
                            va_arg2_try143413= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                            var v_bag_arg_try143514 EID 
                            /* Let:14 */{ 
                              var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                              /* noccur = 5 */
                              _CL_obj.Selector = ToProperty(C__Z.Id())
                              /* update:15 */{ 
                                var va_arg1 *Language.Call  
                                var va_arg2 *ClaireList  
                                va_arg1 = _CL_obj
                                var va_arg2_try143616 EID 
                                /* Construct:16 */{ 
                                  var v_bag_arg *ClaireAny  
                                  va_arg2_try143616= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                                  ToList(OBJ(va_arg2_try143616)).AddFast(y.Id())
                                  var v_bag_arg_try143717 EID 
                                  v_bag_arg_try143717 = F_Optimize_Produce_get_relation(R,x)
                                  /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try143616) */
                                  if ErrorIn(v_bag_arg_try143717) {va_arg2_try143616 = v_bag_arg_try143717
                                  } else {
                                  v_bag_arg = ANY(v_bag_arg_try143717)
                                  ToList(OBJ(va_arg2_try143616)).AddFast(v_bag_arg)}
                                  /* Construct-16 */} 
                                /* ERROR PROTECTION INSERTED (va_arg2-v_bag_arg_try143514) */
                                if ErrorIn(va_arg2_try143616) {v_bag_arg_try143514 = va_arg2_try143616
                                } else {
                                va_arg2 = ToList(OBJ(va_arg2_try143616))
                                /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                                va_arg1.Args = va_arg2
                                v_bag_arg_try143514 = EID{va_arg2.Id(),0}
                                }
                                /* update-15 */} 
                              /* ERROR PROTECTION INSERTED (v_bag_arg_try143514-v_bag_arg_try143514) */
                              if !ErrorIn(v_bag_arg_try143514) {
                              v_bag_arg_try143514 = EID{_CL_obj.Id(),0}
                              }
                              /* Let-14 */} 
                            /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try143413) */
                            if ErrorIn(v_bag_arg_try143514) {va_arg2_try143413 = v_bag_arg_try143514
                            } else {
                            v_bag_arg = ANY(v_bag_arg_try143514)
                            ToList(OBJ(va_arg2_try143413)).AddFast(v_bag_arg)}
                            /* Construct-13 */} 
                          /* ERROR PROTECTION INSERTED (va_arg2-va_arg2_try143311) */
                          if ErrorIn(va_arg2_try143413) {va_arg2_try143311 = va_arg2_try143413
                          } else {
                          va_arg2 = ToList(OBJ(va_arg2_try143413))
                          /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                          va_arg1.Args = va_arg2
                          va_arg2_try143311 = EID{va_arg2.Id(),0}
                          }
                          /* update-12 */} 
                        /* ERROR PROTECTION INSERTED (va_arg2_try143311-va_arg2_try143311) */
                        if !ErrorIn(va_arg2_try143311) {
                        va_arg2_try143311 = EID{_CL_obj.Id(),0}
                        }
                        /* Let-11 */} 
                      /* ERROR PROTECTION INSERTED (va_arg2-g1431UU_try14329) */
                      if ErrorIn(va_arg2_try143311) {g1431UU_try14329 = va_arg2_try143311
                      } else {
                      va_arg2 = ANY(va_arg2_try143311)
                      /* ---------- now we compile update iClaire/test(va_arg1) := va_arg2 ------- */
                      va_arg1.Test = va_arg2
                      g1431UU_try14329 = va_arg2.ToEID()
                      }
                      /* update-10 */} 
                    /* ERROR PROTECTION INSERTED (g1431UU_try14329-g1431UU_try14329) */
                    if !ErrorIn(g1431UU_try14329) {
                    /* update:10 */{ 
                      var va_arg1 *Language.If  
                      var va_arg2 *ClaireAny  
                      va_arg1 = _CL_obj
                      /* Let:11 */{ 
                        var _CL_obj *Language.Do   = Language.To_Do(new(Language.Do).Is(Language.C_Do))
                        /* noccur = 3 */
                        _CL_obj.Args = l1
                        va_arg2 = _CL_obj.Id()
                        /* Let-11 */} 
                      /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                      va_arg1.Arg = va_arg2
                      /* update-10 */} 
                    g1431UU_try14329 = EID{_CL_obj.Id(),0}
                    }
                    /* Let-9 */} 
                  /* ERROR PROTECTION INSERTED (g1431UU-va_arg2_try14308) */
                  if ErrorIn(g1431UU_try14329) {va_arg2_try14308 = g1431UU_try14329
                  } else {
                  g1431UU = Language.To_If(OBJ(g1431UU_try14329))
                  va_arg2_try14308 = Language.F_lambda_I_list(MakeConstantList(x.Id(),y.Id()),g1431UU.Id())
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                if ErrorIn(va_arg2_try14308) {Result = va_arg2_try14308
                } else {
                va_arg2 = ANY(va_arg2_try14308)
                /* ---------- now we compile update if_write(va_arg1) := va_arg2 ------- */
                va_arg1.IfWrite = va_arg2
                Result = va_arg2.ToEID()
                }
                /* update-7 */} 
              }}
              } else {
              var l1_try14387 EID 
              /* Construct:7 */{ 
                var v_bag_arg *ClaireAny  
                l1_try14387= EID{ToType(C_any.Id()).EmptyList().Id(),0}
                var v_bag_arg_try14398 EID 
                v_bag_arg_try14398 = Core.F_CALL(C_Optimize_Produce_put,ARGS(EID{R.Id(),0},EID{x.Id(),0},EID{y.Id(),0}))
                /* ERROR PROTECTION INSERTED (v_bag_arg-l1_try14387) */
                if ErrorIn(v_bag_arg_try14398) {l1_try14387 = v_bag_arg_try14398
                } else {
                v_bag_arg = ANY(v_bag_arg_try14398)
                ToList(OBJ(l1_try14387)).AddFast(v_bag_arg)}
                /* Construct-7 */} 
              /* ERROR PROTECTION INSERTED (l1-Result) */
              if ErrorIn(l1_try14387) {Result = l1_try14387
              } else {
              l1 = ToList(OBJ(l1_try14387))
              Result = EID{l1.Id(),0}
              if (R.Inverse.Id() != CNULL) /* If:7 */{ 
                var l1_try14408 EID 
                /* Let:8 */{ 
                  var g1441UU *Language.If  
                  /* noccur = 1 */
                  var g1441UU_try14429 EID 
                  /* Let:9 */{ 
                    var _CL_obj *Language.If   = Language.To_If(new(Language.If).Is(Language.C_If))
                    /* noccur = 11 */
                    /* update:10 */{ 
                      var va_arg1 *Language.If  
                      var va_arg2 *ClaireAny  
                      va_arg1 = _CL_obj
                      /* Let:11 */{ 
                        var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                        /* noccur = 5 */
                        _CL_obj.Selector = Core.C_known_ask
                        _CL_obj.Args = MakeConstantList(z.Id())
                        va_arg2 = _CL_obj.Id()
                        /* Let-11 */} 
                      /* ---------- now we compile update iClaire/test(va_arg1) := va_arg2 ------- */
                      va_arg1.Test = va_arg2
                      /* update-10 */} 
                    /* update:10 */{ 
                      var va_arg1 *Language.If  
                      var va_arg2 *ClaireAny  
                      va_arg1 = _CL_obj
                      var va_arg2_try144311 EID 
                      va_arg2_try144311 = Core.F_CALL(C_Optimize_Produce_remove,ARGS(EID{R.Inverse.Id(),0},EID{z.Id(),0},EID{x.Id(),0}))
                      /* ERROR PROTECTION INSERTED (va_arg2-g1441UU_try14429) */
                      if ErrorIn(va_arg2_try144311) {g1441UU_try14429 = va_arg2_try144311
                      } else {
                      va_arg2 = ANY(va_arg2_try144311)
                      /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                      va_arg1.Arg = va_arg2
                      g1441UU_try14429 = va_arg2.ToEID()
                      }
                      /* update-10 */} 
                    /* ERROR PROTECTION INSERTED (g1441UU_try14429-g1441UU_try14429) */
                    if !ErrorIn(g1441UU_try14429) {
                    g1441UU_try14429 = EID{_CL_obj.Id(),0}
                    }
                    /* Let-9 */} 
                  /* ERROR PROTECTION INSERTED (g1441UU-l1_try14408) */
                  if ErrorIn(g1441UU_try14429) {l1_try14408 = g1441UU_try14429
                  } else {
                  g1441UU = Language.To_If(OBJ(g1441UU_try14429))
                  l1_try14408 = EID{l1.AddFast(g1441UU.Id()).Id(),0}
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (l1-Result) */
                if ErrorIn(l1_try14408) {Result = l1_try14408
                } else {
                l1 = ToList(OBJ(l1_try14408))
                Result = EID{l1.Id(),0}
                var l1_try14448 EID 
                /* Let:8 */{ 
                  var g1445UU *ClaireAny  
                  /* noccur = 1 */
                  var g1445UU_try14469 EID 
                  g1445UU_try14469 = Core.F_CALL(C_Optimize_Produce_put,ARGS(EID{R.Inverse.Id(),0},EID{y.Id(),0},EID{x.Id(),0}))
                  /* ERROR PROTECTION INSERTED (g1445UU-l1_try14448) */
                  if ErrorIn(g1445UU_try14469) {l1_try14448 = g1445UU_try14469
                  } else {
                  g1445UU = ANY(g1445UU_try14469)
                  l1_try14448 = EID{l1.AddFast(g1445UU).Id(),0}
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (l1-Result) */
                if ErrorIn(l1_try14448) {Result = l1_try14448
                } else {
                l1 = ToList(OBJ(l1_try14448))
                Result = EID{l1.Id(),0}
                }
                }
                } else {
                Result = EID{CFALSE.Id(),0}
                /* If-7 */} 
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              /* update:7 */{ 
                var va_arg1 *ClaireRelation  
                var va_arg2 *ClaireAny  
                va_arg1 = R
                var va_arg2_try14478 EID 
                /* Let:8 */{ 
                  var g1448UU *Language.Let  
                  /* noccur = 1 */
                  var g1448UU_try14499 EID 
                  /* Let:9 */{ 
                    var _CL_obj *Language.Let   = Language.To_Let(new(Language.Let).Is(Language.C_Let))
                    /* noccur = 22 */
                    _CL_obj.ClaireVar = z
                    /* update:10 */{ 
                      var va_arg1 *Language.Let  
                      var va_arg2 *ClaireAny  
                      va_arg1 = _CL_obj
                      var va_arg2_try145011 EID 
                      va_arg2_try145011 = F_Optimize_Produce_get_relation(R,x)
                      /* ERROR PROTECTION INSERTED (va_arg2-g1448UU_try14499) */
                      if ErrorIn(va_arg2_try145011) {g1448UU_try14499 = va_arg2_try145011
                      } else {
                      va_arg2 = ANY(va_arg2_try145011)
                      /* ---------- now we compile update value(va_arg1) := va_arg2 ------- */
                      va_arg1.Value = va_arg2
                      g1448UU_try14499 = va_arg2.ToEID()
                      }
                      /* update-10 */} 
                    /* ERROR PROTECTION INSERTED (g1448UU_try14499-g1448UU_try14499) */
                    if !ErrorIn(g1448UU_try14499) {
                    /* update:10 */{ 
                      var va_arg1 *Language.Let  
                      var va_arg2 *ClaireAny  
                      va_arg1 = _CL_obj
                      /* Let:11 */{ 
                        var _CL_obj *Language.If   = Language.To_If(new(Language.If).Is(Language.C_If))
                        /* noccur = 15 */
                        /* update:12 */{ 
                          var va_arg1 *Language.If  
                          var va_arg2 *ClaireAny  
                          va_arg1 = _CL_obj
                          /* Let:13 */{ 
                            var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                            /* noccur = 5 */
                            _CL_obj.Selector = ToProperty(Core.C__I_equal.Id())
                            _CL_obj.Args = MakeConstantList(y.Id(),z.Id())
                            va_arg2 = _CL_obj.Id()
                            /* Let-13 */} 
                          /* ---------- now we compile update iClaire/test(va_arg1) := va_arg2 ------- */
                          va_arg1.Test = va_arg2
                          /* update-12 */} 
                        /* update:12 */{ 
                          var va_arg1 *Language.If  
                          var va_arg2 *ClaireAny  
                          va_arg1 = _CL_obj
                          /* Let:13 */{ 
                            var _CL_obj *Language.Do   = Language.To_Do(new(Language.Do).Is(Language.C_Do))
                            /* noccur = 3 */
                            _CL_obj.Args = l1
                            va_arg2 = _CL_obj.Id()
                            /* Let-13 */} 
                          /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                          va_arg1.Arg = va_arg2
                          /* update-12 */} 
                        va_arg2 = _CL_obj.Id()
                        /* Let-11 */} 
                      /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                      va_arg1.Arg = va_arg2
                      /* update-10 */} 
                    g1448UU_try14499 = EID{_CL_obj.Id(),0}
                    }
                    /* Let-9 */} 
                  /* ERROR PROTECTION INSERTED (g1448UU-va_arg2_try14478) */
                  if ErrorIn(g1448UU_try14499) {va_arg2_try14478 = g1448UU_try14499
                  } else {
                  g1448UU = Language.To_Let(OBJ(g1448UU_try14499))
                  va_arg2_try14478 = Language.F_lambda_I_list(MakeConstantList(x.Id(),y.Id()),g1448UU.Id())
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                if ErrorIn(va_arg2_try14478) {Result = va_arg2_try14478
                } else {
                va_arg2 = ANY(va_arg2_try14478)
                /* ---------- now we compile update if_write(va_arg1) := va_arg2 ------- */
                va_arg1.IfWrite = va_arg2
                Result = va_arg2.ToEID()
                }
                /* update-7 */} 
              }}
              /* If-6 */} 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            /* Let:6 */{ 
              var dn *ClaireString   = F_append_string(R.Name.String_I(),MakeString("_write"))
              /* noccur = 1 */
              Result = F_Compile_compile_lambda_string(dn,ToLambda(R.IfWrite),C_void.Id())
              /* Let-6 */} 
            }
            /* Let-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: Compile/compute_if_write_inverse @ relation (throw: true) 
func E_Compile_compute_if_write_inverse_relation (R EID) EID { 
    return /*(sm for Compile/compute_if_write_inverse @ relation= EID)*/ F_Compile_compute_if_write_inverse_relation(ToRelation(OBJ(R)) )} 
  
// generate a demon to perform x.R := s (s is a set)
/* {1} OPT.The go function for: Compile/compute_set_write(R:relation) [] */
func F_Compile_compute_set_write_relation (R *ClaireRelation ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var x *ClaireVariable  
      /* noccur = 5 */
      /* Let:3 */{ 
        var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
        /* noccur = 5 */
        _CL_obj.Pname = Core.F_symbol_I_string2(MakeString("XX"))
        _CL_obj.Range = R.Domain
        x = _CL_obj
        /* Let-3 */} 
      /* Let:3 */{ 
        var y *ClaireVariable  
        /* noccur = 2 */
        /* Let:4 */{ 
          var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
          /* noccur = 5 */
          _CL_obj.Pname = Core.F_symbol_I_string2(MakeString("YY"))
          _CL_obj.Range = ToType(C_bag.Id())
          y = _CL_obj
          /* Let-4 */} 
        /* Let:4 */{ 
          var z *ClaireVariable  
          /* noccur = 4 */
          /* Let:5 */{ 
            var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
            /* noccur = 5 */
            _CL_obj.Pname = Core.F_symbol_I_string2(MakeString("ZZ"))
            _CL_obj.Range = Core.F_member_type(R.Range)
            z = _CL_obj
            /* Let-5 */} 
          /* Let:5 */{ 
            var l1 *ClaireList   = ToType(C_any.Id()).EmptyList()
            /* noccur = 7 */
            Core.F_tformat_string(MakeString("compute set_write for ~S \n"),0,MakeConstantList(R.Id()))
            if (R.Inverse.Id() != CNULL) /* If:6 */{ 
              var l1_try14517 EID 
              /* Let:7 */{ 
                var g1452UU *Language.For  
                /* noccur = 1 */
                var g1452UU_try14538 EID 
                /* Let:8 */{ 
                  var _CL_obj *Language.For   = Language.To_For(new(Language.For).Is(Language.C_For))
                  /* noccur = 5 */
                  _CL_obj.ClaireVar = z
                  /* update:9 */{ 
                    var va_arg1 *Language.Iteration  
                    var va_arg2 *ClaireAny  
                    va_arg1 = Language.To_Iteration(_CL_obj.Id())
                    var va_arg2_try145410 EID 
                    va_arg2_try145410 = F_Optimize_Produce_get_relation(R,x)
                    /* ERROR PROTECTION INSERTED (va_arg2-g1452UU_try14538) */
                    if ErrorIn(va_arg2_try145410) {g1452UU_try14538 = va_arg2_try145410
                    } else {
                    va_arg2 = ANY(va_arg2_try145410)
                    /* ---------- now we compile update iClaire/set_arg(va_arg1) := va_arg2 ------- */
                    va_arg1.SetArg = va_arg2
                    g1452UU_try14538 = va_arg2.ToEID()
                    }
                    /* update-9 */} 
                  /* ERROR PROTECTION INSERTED (g1452UU_try14538-g1452UU_try14538) */
                  if !ErrorIn(g1452UU_try14538) {
                  /* update:9 */{ 
                    var va_arg1 *Language.Iteration  
                    var va_arg2 *ClaireAny  
                    va_arg1 = Language.To_Iteration(_CL_obj.Id())
                    var va_arg2_try145510 EID 
                    va_arg2_try145510 = Core.F_CALL(C_Optimize_Produce_remove,ARGS(EID{R.Inverse.Id(),0},EID{z.Id(),0},EID{x.Id(),0}))
                    /* ERROR PROTECTION INSERTED (va_arg2-g1452UU_try14538) */
                    if ErrorIn(va_arg2_try145510) {g1452UU_try14538 = va_arg2_try145510
                    } else {
                    va_arg2 = ANY(va_arg2_try145510)
                    /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                    va_arg1.Arg = va_arg2
                    g1452UU_try14538 = va_arg2.ToEID()
                    }
                    /* update-9 */} 
                  /* ERROR PROTECTION INSERTED (g1452UU_try14538-g1452UU_try14538) */
                  if !ErrorIn(g1452UU_try14538) {
                  g1452UU_try14538 = EID{_CL_obj.Id(),0}
                  }}
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (g1452UU-l1_try14517) */
                if ErrorIn(g1452UU_try14538) {l1_try14517 = g1452UU_try14538
                } else {
                g1452UU = Language.To_For(OBJ(g1452UU_try14538))
                l1_try14517 = EID{l1.AddFast(g1452UU.Id()).Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (l1-Result) */
              if ErrorIn(l1_try14517) {Result = l1_try14517
              } else {
              l1 = ToList(OBJ(l1_try14517))
              Result = EID{l1.Id(),0}
              }
              } else {
              Result = EID{CFALSE.Id(),0}
              /* If-6 */} 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            var l1_try14566 EID 
            /* Let:6 */{ 
              var g1457UU *ClaireAny  
              /* noccur = 1 */
              var g1457UU_try14587 EID 
              g1457UU_try14587 = F_Optimize_Produce_erase_property(ToProperty(R.Id()),x)
              /* ERROR PROTECTION INSERTED (g1457UU-l1_try14566) */
              if ErrorIn(g1457UU_try14587) {l1_try14566 = g1457UU_try14587
              } else {
              g1457UU = ANY(g1457UU_try14587)
              l1_try14566 = EID{l1.AddFast(g1457UU).Id(),0}
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (l1-Result) */
            if ErrorIn(l1_try14566) {Result = l1_try14566
            } else {
            l1 = ToList(OBJ(l1_try14566))
            Result = EID{l1.Id(),0}
            var l1_try14596 EID 
            /* Let:6 */{ 
              var g1460UU *Language.For  
              /* noccur = 1 */
              var g1460UU_try14617 EID 
              /* Let:7 */{ 
                var _CL_obj *Language.For   = Language.To_For(new(Language.For).Is(Language.C_For))
                /* noccur = 5 */
                _CL_obj.ClaireVar = z
                _CL_obj.SetArg = y.Id()
                /* update:8 */{ 
                  var va_arg1 *Language.Iteration  
                  var va_arg2 *ClaireAny  
                  va_arg1 = Language.To_Iteration(_CL_obj.Id())
                  var va_arg2_try14629 EID 
                  va_arg2_try14629 = Core.F_CALL(C_Optimize_Produce_put,ARGS(EID{R.Id(),0},EID{x.Id(),0},EID{z.Id(),0}))
                  /* ERROR PROTECTION INSERTED (va_arg2-g1460UU_try14617) */
                  if ErrorIn(va_arg2_try14629) {g1460UU_try14617 = va_arg2_try14629
                  } else {
                  va_arg2 = ANY(va_arg2_try14629)
                  /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                  va_arg1.Arg = va_arg2
                  g1460UU_try14617 = va_arg2.ToEID()
                  }
                  /* update-8 */} 
                /* ERROR PROTECTION INSERTED (g1460UU_try14617-g1460UU_try14617) */
                if !ErrorIn(g1460UU_try14617) {
                g1460UU_try14617 = EID{_CL_obj.Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (g1460UU-l1_try14596) */
              if ErrorIn(g1460UU_try14617) {l1_try14596 = g1460UU_try14617
              } else {
              g1460UU = Language.To_For(OBJ(g1460UU_try14617))
              l1_try14596 = EID{l1.AddFast(g1460UU.Id()).Id(),0}
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (l1-Result) */
            if ErrorIn(l1_try14596) {Result = l1_try14596
            } else {
            l1 = ToList(OBJ(l1_try14596))
            Result = EID{l1.Id(),0}
            /* Let:6 */{ 
              var dn *ClaireString   = F_append_string(R.Name.String_I(),MakeString("_set_write"))
              /* noccur = 1 */
              /* Let:7 */{ 
                var g1463UU *ClaireLambda  
                /* noccur = 1 */
                var g1463UU_try14648 EID 
                /* Let:8 */{ 
                  var g1465UU *Language.Do  
                  /* noccur = 1 */
                  /* Let:9 */{ 
                    var _CL_obj *Language.Do   = Language.To_Do(new(Language.Do).Is(Language.C_Do))
                    /* noccur = 3 */
                    _CL_obj.Args = l1
                    g1465UU = _CL_obj
                    /* Let-9 */} 
                  g1463UU_try14648 = Language.F_lambda_I_list(MakeConstantList(x.Id(),y.Id()),g1465UU.Id())
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (g1463UU-Result) */
                if ErrorIn(g1463UU_try14648) {Result = g1463UU_try14648
                } else {
                g1463UU = ToLambda(OBJ(g1463UU_try14648))
                Result = F_Compile_compile_lambda_string(dn,g1463UU,C_void.Id())
                }
                /* Let-7 */} 
              /* Let-6 */} 
            }}}
            /* Let-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: Compile/compute_set_write @ relation (throw: true) 
func E_Compile_compute_set_write_relation (R EID) EID { 
    return /*(sm for Compile/compute_set_write @ relation= EID)*/ F_Compile_compute_set_write_relation(ToRelation(OBJ(R)) )} 
  
// generate a simple put for a property => generate a case to make sure
// that we get the fastest possible code
/* {1} OPT.The go function for: Produce_put(r:property,x:Variable,y:any) [] */
func F_Optimize_Produce_put_property (r *ClaireProperty ,x *ClaireVariable ,y *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var l *ClaireList   = ToType(C_any.Id()).EmptyList()
      /* noccur = 5 */
      /* For:3 */{ 
        var xs *ClaireAny  
        _ = xs
        Result= EID{CFALSE.Id(),0}
        for _,xs = range(r.Restrictions.ValuesO())/* loop:4 */{ 
          var void_try5 EID 
          _ = void_try5
          if ((C_slot.Id() == xs.Isa.Id()) && 
              (F_boolean_I_any(Core.F__exp_type(F_Optimize_ptype_type(x.Range),ToType(Core.F_domain_I_restriction(ToRestriction(xs)).Id())).Id()) == CTRUE)) /* If:5 */{ 
            var l_try14666 EID 
            /* Let:6 */{ 
              var g1467UU *ClaireList  
              /* noccur = 1 */
              /* Construct:7 */{ 
                var v_bag_arg *ClaireAny  
                g1467UU= ToType(CEMPTY.Id()).EmptyList()
                g1467UU.AddFast(Core.F_domain_I_restriction(ToRestriction(xs)).Id())
                if (r.Multivalued_ask == CTRUE) /* If:8 */{ 
                  /* Let:9 */{ 
                    var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                    /* noccur = 16 */
                    _CL_obj.Selector = ToProperty(C_add_I.Id())
                    /* update:10 */{ 
                      var va_arg1 *Language.Call  
                      var va_arg2 *ClaireList  
                      va_arg1 = _CL_obj
                      /* Construct:11 */{ 
                        var v_bag_arg *ClaireAny  
                        va_arg2= ToType(CEMPTY.Id()).EmptyList()
                        /* Let:12 */{ 
                          var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                          /* noccur = 10 */
                          _CL_obj.Selector = r
                          /* update:13 */{ 
                            var va_arg1 *Language.Call  
                            var va_arg2 *ClaireList  
                            va_arg1 = _CL_obj
                            /* Construct:14 */{ 
                              var v_bag_arg *ClaireAny  
                              va_arg2= ToType(CEMPTY.Id()).EmptyList()
                              /* Let:15 */{ 
                                var _CL_obj *Language.Cast   = Language.To_Cast(new(Language.Cast).Is(Language.C_Cast))
                                /* noccur = 4 */
                                _CL_obj.Arg = x.Id()
                                _CL_obj.SetArg = ToType(Core.F_domain_I_restriction(ToRestriction(xs)).Id())
                                v_bag_arg = _CL_obj.Id()
                                /* Let-15 */} 
                              va_arg2.AddFast(v_bag_arg)/* Construct-14 */} 
                            /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                            va_arg1.Args = va_arg2
                            /* update-13 */} 
                          v_bag_arg = _CL_obj.Id()
                          /* Let-12 */} 
                        va_arg2.AddFast(v_bag_arg)
                        va_arg2.AddFast(y)/* Construct-11 */} 
                      /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                      va_arg1.Args = va_arg2
                      /* update-10 */} 
                    v_bag_arg = _CL_obj.Id()
                    /* Let-9 */} 
                  } else {
                  /* Let:9 */{ 
                    var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                    /* noccur = 10 */
                    _CL_obj.Selector = C_put
                    /* update:10 */{ 
                      var va_arg1 *Language.Call  
                      var va_arg2 *ClaireList  
                      va_arg1 = _CL_obj
                      /* Construct:11 */{ 
                        var v_bag_arg *ClaireAny  
                        va_arg2= ToType(CEMPTY.Id()).EmptyList()
                        va_arg2.AddFast(r.Id())
                        /* Let:12 */{ 
                          var _CL_obj *Language.Cast   = Language.To_Cast(new(Language.Cast).Is(Language.C_Cast))
                          /* noccur = 4 */
                          _CL_obj.Arg = x.Id()
                          _CL_obj.SetArg = ToType(Core.F_domain_I_restriction(ToRestriction(xs)).Id())
                          v_bag_arg = _CL_obj.Id()
                          /* Let-12 */} 
                        va_arg2.AddFast(v_bag_arg)
                        va_arg2.AddFast(y)/* Construct-11 */} 
                      /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                      va_arg1.Args = va_arg2
                      /* update-10 */} 
                    v_bag_arg = _CL_obj.Id()
                    /* Let-9 */} 
                  /* If-8 */} 
                g1467UU.AddFast(v_bag_arg)/* Construct-7 */} 
              l_try14666 = l.Add_star(g1467UU)
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (l-void_try5) */
            if ErrorIn(l_try14666) {void_try5 = l_try14666
            } else {
            l = ToList(OBJ(l_try14666))
            void_try5 = EID{l.Id(),0}
            }
            } else {
            void_try5 = EID{CFALSE.Id(),0}
            /* If-5 */} 
          /* ERROR PROTECTION INSERTED (void_try5-Result) */
          if ErrorIn(void_try5) {Result = void_try5
          Result = void_try5
          break
          } else {
          }
          /* loop-4 */} 
        /* For-3 */} 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      if (l.Length() == 2) /* If:3 */{ 
        Result = l.At(2-1).ToEID()
        } else {
        /* Let:4 */{ 
          var _CL_obj *Language.Case   = Language.To_Case(new(Language.Case).Is(Language.C_Case))
          /* noccur = 4 */
          _CL_obj.ClaireVar = x.Id()
          _CL_obj.Args = l
          Result = EID{_CL_obj.Id(),0}
          /* Let-4 */} 
        /* If-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: Produce_put @ property (throw: true) 
func E_Optimize_Produce_put_property (r EID,x EID,y EID) EID { 
    return /*(sm for Produce_put @ property= EID)*/ F_Optimize_Produce_put_property(ToProperty(OBJ(r)),To_Variable(OBJ(x)),ANY(y) )} 
  
// generate a simple erase (the inverse management has been done)
// v3.2.50: use ptype(x.range) for variable whose type is t U any :-)
/* {1} OPT.The go function for: Produce_erase(r:property,x:Variable) [] */
func F_Optimize_Produce_erase_property (r *ClaireProperty ,x *ClaireVariable ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var l *ClaireList   = ToType(C_any.Id()).EmptyList()
      /* noccur = 5 */
      /* Let:3 */{ 
        var val *ClaireBag  
        /* noccur = 2 */
        if (r.Multivalued_ask.Id() == C_list.Id()) /* If:4 */{ 
          val = ToBag(ToType(C_any.Id()).EmptyList().Id())
          } else {
          val = ToBag(ToType(C_any.Id()).EmptySet().Id())
          /* If-4 */} 
        val.Cast_I(Core.F_member_type(r.Range))
        /* For:4 */{ 
          var xs *ClaireAny  
          _ = xs
          Result= EID{CFALSE.Id(),0}
          for _,xs = range(r.Restrictions.ValuesO())/* loop:5 */{ 
            var void_try6 EID 
            _ = void_try6
            if ((C_slot.Id() == xs.Isa.Id()) && 
                (F_boolean_I_any(Core.F__exp_type(F_Optimize_ptype_type(x.Range),ToType(Core.F_domain_I_restriction(ToRestriction(xs)).Id())).Id()) == CTRUE)) /* If:6 */{ 
              var l_try14687 EID 
              /* Let:7 */{ 
                var g1469UU *ClaireList  
                /* noccur = 1 */
                /* Construct:8 */{ 
                  var v_bag_arg *ClaireAny  
                  g1469UU= ToType(CEMPTY.Id()).EmptyList()
                  g1469UU.AddFast(Core.F_domain_I_restriction(ToRestriction(xs)).Id())
                  /* Let:9 */{ 
                    var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                    /* noccur = 10 */
                    _CL_obj.Selector = C_put
                    /* update:10 */{ 
                      var va_arg1 *Language.Call  
                      var va_arg2 *ClaireList  
                      va_arg1 = _CL_obj
                      /* Construct:11 */{ 
                        var v_bag_arg *ClaireAny  
                        va_arg2= ToType(CEMPTY.Id()).EmptyList()
                        va_arg2.AddFast(r.Id())
                        /* Let:12 */{ 
                          var _CL_obj *Language.Cast   = Language.To_Cast(new(Language.Cast).Is(Language.C_Cast))
                          /* noccur = 4 */
                          _CL_obj.Arg = x.Id()
                          _CL_obj.SetArg = ToType(Core.F_domain_I_restriction(ToRestriction(xs)).Id())
                          v_bag_arg = _CL_obj.Id()
                          /* Let-12 */} 
                        va_arg2.AddFast(v_bag_arg)
                        if (r.Multivalued_ask == CTRUE) /* If:12 */{ 
                          v_bag_arg = val.Id()
                          } else {
                          v_bag_arg = ToSlot(xs).Default
                          /* If-12 */} 
                        va_arg2.AddFast(v_bag_arg)/* Construct-11 */} 
                      /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                      va_arg1.Args = va_arg2
                      /* update-10 */} 
                    v_bag_arg = _CL_obj.Id()
                    /* Let-9 */} 
                  g1469UU.AddFast(v_bag_arg)/* Construct-8 */} 
                l_try14687 = l.Add_star(g1469UU)
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (l-void_try6) */
              if ErrorIn(l_try14687) {void_try6 = l_try14687
              } else {
              l = ToList(OBJ(l_try14687))
              void_try6 = EID{l.Id(),0}
              }
              } else {
              void_try6 = EID{CFALSE.Id(),0}
              /* If-6 */} 
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
        if (l.Length() == 2) /* If:4 */{ 
          Result = l.At(2-1).ToEID()
          } else {
          /* Let:5 */{ 
            var _CL_obj *Language.Case   = Language.To_Case(new(Language.Case).Is(Language.C_Case))
            /* noccur = 4 */
            _CL_obj.ClaireVar = x.Id()
            _CL_obj.Args = l
            Result = EID{_CL_obj.Id(),0}
            /* Let-5 */} 
          /* If-4 */} 
        }
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: Produce_erase @ property (throw: true) 
func E_Optimize_Produce_erase_property (r EID,x EID) EID { 
    return /*(sm for Produce_erase @ property= EID)*/ F_Optimize_Produce_erase_property(ToProperty(OBJ(r)),To_Variable(OBJ(x)) )} 
  
// note:  (a) Simpler because of v3.0 !! (siude-effects on lists or sets)
//        (b) if |l|= 1 domain!(r) = domain!(x) because of tighten
// same for a table
/* {1} OPT.The go function for: Produce_put(r:table,x:Variable,y:any) [] */
func F_Optimize_Produce_put_table (r *ClaireTable ,x *ClaireVariable ,y *ClaireAny ) *ClaireAny  { 
    // procedure body with s =  
var Result *ClaireAny  
    /* Let:2 */{ 
      var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
      /* noccur = 11 */
      _CL_obj.Selector = C_put
      /* update:3 */{ 
        var va_arg1 *Language.Call  
        var va_arg2 *ClaireList  
        va_arg1 = _CL_obj
        /* Construct:4 */{ 
          var v_bag_arg *ClaireAny  
          va_arg2= ToType(CEMPTY.Id()).EmptyList()
          va_arg2.AddFast(r.Id())
          va_arg2.AddFast(x.Id())
          if (r.Multivalued_ask == CTRUE) /* If:5 */{ 
            /* Let:6 */{ 
              var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
              /* noccur = 5 */
              _CL_obj.Selector = ToProperty(C_add.Id())
              _CL_obj.Args = MakeConstantList(MakeConstantList(C_nth.Id(),MakeConstantList(r.Id(),x.Id()).Id()).Id(),y)
              v_bag_arg = _CL_obj.Id()
              /* Let-6 */} 
            } else {
            v_bag_arg = y
            /* If-5 */} 
          va_arg2.AddFast(v_bag_arg)/* Construct-4 */} 
        /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
        va_arg1.Args = va_arg2
        /* update-3 */} 
      Result = _CL_obj.Id()
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: Produce_put @ table (throw: false) 
func E_Optimize_Produce_put_table (r EID,x EID,y EID) EID { 
    return /*(sm for Produce_put @ table= any)*/ F_Optimize_Produce_put_table(ToTable(OBJ(r)),To_Variable(OBJ(x)),ANY(y) ).ToEID()} 
  
/* {1} OPT.The go function for: Produce_get(r:relation,x:Variable) [] */
func F_Optimize_Produce_get_relation (r *ClaireRelation ,x *ClaireVariable ) EID { 
    var Result EID 
    if (C_table.Id() == r.Isa.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var g1470 *ClaireTable   = ToTable(r.Id())
        /* noccur = 1 */
        /* Let:4 */{ 
          var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
          /* noccur = 5 */
          _CL_obj.Selector = C_nth
          _CL_obj.Args = MakeConstantList(g1470.Id(),x.Id())
          Result = EID{_CL_obj.Id(),0}
          /* Let-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (r.Isa.IsIn(C_property) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g1471 *ClaireProperty   = ToProperty(r.Id())
        /* noccur = 2 */
        /* Let:4 */{ 
          var l *ClaireList   = ToType(C_any.Id()).EmptyList()
          /* noccur = 5 */
          /* For:5 */{ 
            var xs *ClaireAny  
            _ = xs
            Result= EID{CFALSE.Id(),0}
            for _,xs = range(g1471.Restrictions.ValuesO())/* loop:6 */{ 
              var void_try7 EID 
              _ = void_try7
              if ((C_slot.Id() == xs.Isa.Id()) && 
                  (F_boolean_I_any(Core.F__exp_type(F_Optimize_ptype_type(x.Range),ToType(Core.F_domain_I_restriction(ToRestriction(xs)).Id())).Id()) == CTRUE)) /* If:7 */{ 
                var l_try14728 EID 
                /* Let:8 */{ 
                  var g1473UU *ClaireList  
                  /* noccur = 1 */
                  /* Construct:9 */{ 
                    var v_bag_arg *ClaireAny  
                    g1473UU= ToType(CEMPTY.Id()).EmptyList()
                    g1473UU.AddFast(Core.F_domain_I_restriction(ToRestriction(xs)).Id())
                    /* Let:10 */{ 
                      var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                      /* noccur = 10 */
                      _CL_obj.Selector = g1471
                      /* update:11 */{ 
                        var va_arg1 *Language.Call  
                        var va_arg2 *ClaireList  
                        va_arg1 = _CL_obj
                        /* Construct:12 */{ 
                          var v_bag_arg *ClaireAny  
                          va_arg2= ToType(CEMPTY.Id()).EmptyList()
                          /* Let:13 */{ 
                            var _CL_obj *Language.Cast   = Language.To_Cast(new(Language.Cast).Is(Language.C_Cast))
                            /* noccur = 4 */
                            _CL_obj.Arg = x.Id()
                            _CL_obj.SetArg = ToType(Core.F_domain_I_restriction(ToRestriction(xs)).Id())
                            v_bag_arg = _CL_obj.Id()
                            /* Let-13 */} 
                          va_arg2.AddFast(v_bag_arg)/* Construct-12 */} 
                        /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                        va_arg1.Args = va_arg2
                        /* update-11 */} 
                      v_bag_arg = _CL_obj.Id()
                      /* Let-10 */} 
                    g1473UU.AddFast(v_bag_arg)/* Construct-9 */} 
                  l_try14728 = l.Add_star(g1473UU)
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (l-void_try7) */
                if ErrorIn(l_try14728) {void_try7 = l_try14728
                } else {
                l = ToList(OBJ(l_try14728))
                void_try7 = EID{l.Id(),0}
                }
                } else {
                void_try7 = EID{CFALSE.Id(),0}
                /* If-7 */} 
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
          if (l.Length() == 2) /* If:5 */{ 
            Result = l.At(2-1).ToEID()
            } else {
            /* Let:6 */{ 
              var _CL_obj *Language.Case   = Language.To_Case(new(Language.Case).Is(Language.C_Case))
              /* noccur = 4 */
              _CL_obj.ClaireVar = x.Id()
              _CL_obj.Args = l
              Result = EID{_CL_obj.Id(),0}
              /* Let-6 */} 
            /* If-5 */} 
          }
          /* Let-4 */} 
        /* Let-3 */} 
      } else {
      Result = EID{CFALSE.Id(),0}
      /* If-2 */} 
    return Result} 
  
// The EID go function for: Produce_get @ relation (throw: true) 
func E_Optimize_Produce_get_relation (r EID,x EID) EID { 
    return /*(sm for Produce_get @ relation= EID)*/ F_Optimize_Produce_get_relation(ToRelation(OBJ(r)),To_Variable(OBJ(x)) )} 
  
// generate a remove
/* {1} OPT.The go function for: Produce_remove(r:property,x:Variable,y:any) [] */
func F_Optimize_Produce_remove_property (r *ClaireProperty ,x *ClaireVariable ,y *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var l *ClaireList   = ToType(C_any.Id()).EmptyList()
      /* noccur = 5 */
      /* For:3 */{ 
        var xs *ClaireAny  
        _ = xs
        Result= EID{CFALSE.Id(),0}
        for _,xs = range(r.Restrictions.ValuesO())/* loop:4 */{ 
          var void_try5 EID 
          _ = void_try5
          if (C_slot.Id() == xs.Isa.Id()) /* If:5 */{ 
            var l_try14746 EID 
            /* Let:6 */{ 
              var g1475UU *ClaireList  
              /* noccur = 1 */
              /* Construct:7 */{ 
                var v_bag_arg *ClaireAny  
                g1475UU= ToType(CEMPTY.Id()).EmptyList()
                g1475UU.AddFast(Core.F_domain_I_restriction(ToRestriction(xs)).Id())
                if (r.Multivalued_ask == CTRUE) /* If:8 */{ 
                  /* Let:9 */{ 
                    var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                    /* noccur = 11 */
                    _CL_obj.Selector = ToProperty(C_delete.Id())
                    /* update:10 */{ 
                      var va_arg1 *Language.Call  
                      var va_arg2 *ClaireList  
                      va_arg1 = _CL_obj
                      /* Construct:11 */{ 
                        var v_bag_arg *ClaireAny  
                        va_arg2= ToType(CEMPTY.Id()).EmptyList()
                        /* Let:12 */{ 
                          var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                          /* noccur = 5 */
                          _CL_obj.Selector = r
                          _CL_obj.Args = MakeConstantList(x.Id())
                          v_bag_arg = _CL_obj.Id()
                          /* Let-12 */} 
                        va_arg2.AddFast(v_bag_arg)
                        va_arg2.AddFast(y)/* Construct-11 */} 
                      /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                      va_arg1.Args = va_arg2
                      /* update-10 */} 
                    v_bag_arg = _CL_obj.Id()
                    /* Let-9 */} 
                  } else {
                  /* Let:9 */{ 
                    var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                    /* noccur = 5 */
                    _CL_obj.Selector = C_put
                    _CL_obj.Args = MakeConstantList(r.Id(),x.Id(),CNULL)
                    v_bag_arg = _CL_obj.Id()
                    /* Let-9 */} 
                  /* If-8 */} 
                g1475UU.AddFast(v_bag_arg)/* Construct-7 */} 
              l_try14746 = l.Add_star(g1475UU)
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (l-void_try5) */
            if ErrorIn(l_try14746) {void_try5 = l_try14746
            } else {
            l = ToList(OBJ(l_try14746))
            void_try5 = EID{l.Id(),0}
            }
            } else {
            void_try5 = EID{CFALSE.Id(),0}
            /* If-5 */} 
          /* ERROR PROTECTION INSERTED (void_try5-Result) */
          if ErrorIn(void_try5) {Result = void_try5
          Result = void_try5
          break
          } else {
          }
          /* loop-4 */} 
        /* For-3 */} 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      if (l.Length() == 2) /* If:3 */{ 
        Result = l.At(2-1).ToEID()
        } else {
        /* Let:4 */{ 
          var _CL_obj *Language.Case   = Language.To_Case(new(Language.Case).Is(Language.C_Case))
          /* noccur = 4 */
          _CL_obj.ClaireVar = x.Id()
          _CL_obj.Args = l
          Result = EID{_CL_obj.Id(),0}
          /* Let-4 */} 
        /* If-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: Produce_remove @ property (throw: true) 
func E_Optimize_Produce_remove_property (r EID,x EID,y EID) EID { 
    return /*(sm for Produce_remove @ property= EID)*/ F_Optimize_Produce_remove_property(ToProperty(OBJ(r)),To_Variable(OBJ(x)),ANY(y) )} 
  
// same for a table
/* {1} OPT.The go function for: Produce_remove(r:table,x:Variable,y:any) [] */
func F_Optimize_Produce_remove_table (r *ClaireTable ,x *ClaireVariable ,y *ClaireAny ) *ClaireAny  { 
    // procedure body with s =  
var Result *ClaireAny  
    /* Let:2 */{ 
      var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
      /* noccur = 11 */
      _CL_obj.Selector = C_put
      /* update:3 */{ 
        var va_arg1 *Language.Call  
        var va_arg2 *ClaireList  
        va_arg1 = _CL_obj
        /* Construct:4 */{ 
          var v_bag_arg *ClaireAny  
          va_arg2= ToType(CEMPTY.Id()).EmptyList()
          va_arg2.AddFast(r.Id())
          va_arg2.AddFast(x.Id())
          if (r.Multivalued_ask == CTRUE) /* If:5 */{ 
            /* Let:6 */{ 
              var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
              /* noccur = 5 */
              _CL_obj.Selector = ToProperty(C_delete.Id())
              _CL_obj.Args = MakeConstantList(MakeConstantList(C_nth.Id(),MakeConstantList(r.Id(),x.Id()).Id()).Id(),y)
              v_bag_arg = _CL_obj.Id()
              /* Let-6 */} 
            } else {
            v_bag_arg = CNULL
            /* If-5 */} 
          va_arg2.AddFast(v_bag_arg)/* Construct-4 */} 
        /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
        va_arg1.Args = va_arg2
        /* update-3 */} 
      Result = _CL_obj.Id()
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: Produce_remove @ table (throw: false) 
func E_Optimize_Produce_remove_table (r EID,x EID,y EID) EID { 
    return /*(sm for Produce_remove @ table= any)*/ F_Optimize_Produce_remove_table(ToTable(OBJ(r)),To_Variable(OBJ(x)),ANY(y) ).ToEID()} 
  
/* {1} OPT.The go function for: Tighten(r:relation) [] */
func F_Optimize_Tighten_relation (r *ClaireRelation )  { 
    // procedure body with s =  
if (r.Isa.IsIn(C_property) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g1476 *ClaireProperty   = ToProperty(r.Id())
        /* noccur = 8 */
        /* Let:4 */{ 
          var ad *ClaireType   = ToType(ToType(CEMPTY.Id()).EmptySet().Id())
          /* noccur = 3 */
          /* Let:5 */{ 
            var ar *ClaireType   = ToType(ToType(CEMPTY.Id()).EmptySet().Id())
            /* noccur = 5 */
            /* For:6 */{ 
              var s *ClaireAny  
              _ = s
              for _,s = range(g1476.Restrictions.ValuesO())/* loop:7 */{ 
                if (C_slot.Id() == s.Isa.Id()) /* If:8 */{ 
                  ad = Core.F_U_type(ad,ToType(Core.F_domain_I_restriction(ToRestriction(s)).Id()))
                  /* Let:9 */{ 
                    var g1477UU *ClaireType  
                    /* noccur = 1 */
                    if (F_boolean_I_any(ANY(Core.F_CALL(C_Optimize_multi_ask,ARGS(EID{g1476.Id(),0})))) == CTRUE) /* If:10 */{ 
                      g1477UU = Core.F_member_type(ToRestriction(s).Range)
                      } else {
                      g1477UU = ToRestriction(s).Range
                      /* If-10 */} 
                    ar = Core.F_U_type(ar,g1477UU)
                    /* Let-9 */} 
                  /* If-8 */} 
                /* loop-7 */} 
              /* For-6 */} 
            g1476.Open = 1
            g1476.Domain = ToType(ad.Class_I().Id())
            /* update:6 */{ 
              var va_arg1 *ClaireRelation  
              var va_arg2 *ClaireType  
              va_arg1 = ToRelation(g1476.Id())
              if (g1476.Multivalued_ask.Id() == C_list.Id()) /* If:7 */{ 
                va_arg2 = Core.F_param_I_class(C_list,ToType(ar.Class_I().Id()))
                /* If!7 */}  else if (g1476.Multivalued_ask.Id() == C_set.Id()) /* If:7 */{ 
                va_arg2 = Core.F_param_I_class(C_set,ToType(ar.Class_I().Id()))
                } else {
                va_arg2 = ar
                /* If-7 */} 
              /* ---------- now we compile update range(va_arg1) := va_arg2 ------- */
              va_arg1.Range = va_arg2
              /* update-6 */} 
            
            /* Let-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* If-2 */} 
    } 
  
// The EID go function for: Tighten @ relation (throw: false) 
func E_Optimize_Tighten_relation (r EID) EID { 
    /*(sm for Tighten @ relation= void)*/ F_Optimize_Tighten_relation(ToRelation(OBJ(r)) )
    return EVOID} 
  
// new: re-compute the numbering but without the side-effects of the interpreter version (v3.067)
/* {1} OPT.The go function for: Compile/lexical_num(self:any,n:integer) [] */
func F_Compile_lexical_num_any (self *ClaireAny ,n int) EID { 
    var Result EID 
    if (self.Isa.IsIn(Language.C_Call) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g1478 *Language.Call   = Language.To_Call(self)
        /* noccur = 1 */
        Result = F_Compile_lexical_num_any(g1478.Args.Id(),n)
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(C_Instruction) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g1479 *ClaireInstruction   = To_Instruction(self)
        /* noccur = 3 */
        /* Let:4 */{ 
          var _Ztype *ClaireClass   = g1479.Isa
          /* noccur = 2 */
          if (Language.C_Instruction_with_var.Descendents.Contain_ask(_Ztype.Id()) == CTRUE) /* If:5 */{ 
            Result = Core.F_put_property2(C_mClaire_index,ToObject(OBJ(Core.F_CALL(Language.C_var,ARGS(EID{g1479.Id(),0})))),MakeInteger(n).Id())
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            n = (n+1)
            if (n > ToInteger(Language.C__starvariable_index_star.Value).Value) /* If:6 */{ 
              var v_gassign7 *ClaireAny  
              v_gassign7 = MakeInteger(n).Id()
              Language.C__starvariable_index_star.Value = v_gassign7
              Result = v_gassign7.ToEID()
              } else {
              Result = EID{CFALSE.Id(),0}
              /* If-6 */} 
            }
            } else {
            Result = EID{CFALSE.Id(),0}
            /* If-5 */} 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          /* For:5 */{ 
            var s *ClaireAny  
            _ = s
            Result= EID{CFALSE.Id(),0}
            for _,s = range(_Ztype.Slots.ValuesO())/* loop:6 */{ 
              var void_try7 EID 
              _ = void_try7
              void_try7 = F_Compile_lexical_num_any(Core.F_get_slot(ToSlot(s),ToObject(g1479.Id())),n)
              /* ERROR PROTECTION INSERTED (void_try7-Result) */
              if ErrorIn(void_try7) {Result = void_try7
              Result = void_try7
              break
              } else {
              }
              /* loop-6 */} 
            /* For-5 */} 
          }
          /* Let-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(C_bag) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g1480 *ClaireBag   = ToBag(self)
        /* noccur = 1 */
        /* For:4 */{ 
          var x *ClaireAny  
          _ = x
          Result= EID{CFALSE.Id(),0}
          var x_support *ClaireList  
          var x_support_try14865 EID 
          x_support_try14865 = Core.F_enumerate_any(g1480.Id())
          /* ERROR PROTECTION INSERTED (x_support-Result) */
          if ErrorIn(x_support_try14865) {Result = x_support_try14865
          } else {
          x_support = ToList(OBJ(x_support_try14865))
          x_len := x_support.Length()
          for i_it := 0; i_it < x_len; i_it++ { 
            x = x_support.At(i_it)
            var void_try6 EID 
            _ = void_try6
            void_try6 = F_Compile_lexical_num_any(x,n)
            /* ERROR PROTECTION INSERTED (void_try6-Result) */
            if ErrorIn(void_try6) {Result = void_try6
            Result = void_try6
            break
            } else {
            }}
            /* loop-5 */} 
          /* For-4 */} 
        /* Let-3 */} 
      } else {
      Result = EID{CNIL.Id(),0}
      /* If-2 */} 
    return Result} 
  
// The EID go function for: Compile/lexical_num @ any (throw: true) 
func E_Compile_lexical_num_any (self EID,n EID) EID { 
    return /*(sm for Compile/lexical_num @ any= EID)*/ F_Compile_lexical_num_any(ANY(self),INT(n) )} 
  
// v3.2 -----------------------------------------------------------------
/* {1} OPT.The go function for: c_type(self:Defrule) [] */
func F_c_type_Defrule (self *Language.Defrule ) *ClaireType  { 
    // use function body compiling 
return  ToType(C_any.Id())
    } 
  
// The EID go function for: c_type @ Defrule (throw: false) 
func E_c_type_Defrule (self EID) EID { 
    return EID{/*(sm for c_type @ Defrule= type)*/ F_c_type_Defrule(Language.To_Defrule(OBJ(self)) ).Id(),0}} 
  
// compile a rule definition
/* {1} OPT.The go function for: c_code(self:Defrule,s:class) [] */
func F_c_code_Defrule (self *Language.Defrule ,s *ClaireClass ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var ru *ClaireAny   = self.Ident.Get()
      /* noccur = 4 */
      /* Let:3 */{ 
        var l *ClaireList   = ToType(C_any.Id()).EmptyList()
        /* noccur = 6 */
        Core.F_tformat_string(MakeString("compile a rule ~S \n"),0,MakeConstantList(ru))
        /* For:4 */{ 
          var r *ClaireAny  
          _ = r
          Result= EID{CFALSE.Id(),0}
          var r_support *ClaireSet  
          var r_support_try14875 EID 
          r_support_try14875 = Core.F_nth_table1(Language.C_Language_relations,ru)
          /* ERROR PROTECTION INSERTED (r_support-Result) */
          if ErrorIn(r_support_try14875) {Result = r_support_try14875
          } else {
          r_support = ToSet(OBJ(r_support_try14875))
          for _,r = range(r_support.Values)/* loop2:5 */{ 
            if (Language.F_eventMethod_ask_relation2(ToRelation(r)) != CTRUE) /* If:6 */{ 
              F_Optimize_Tighten_relation(ToRelation(r))
              /* If-6 */} 
            }
            /* loop-5 */} 
          /* For-4 */} 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        /* For:4 */{ 
          var r *ClaireAny  
          _ = r
          Result= EID{CFALSE.Id(),0}
          var r_support *ClaireSet  
          var r_support_try14895 EID 
          r_support_try14895 = Core.F_nth_table1(Language.C_Language_relations,ru)
          /* ERROR PROTECTION INSERTED (r_support-Result) */
          if ErrorIn(r_support_try14895) {Result = r_support_try14895
          } else {
          r_support = ToSet(OBJ(r_support_try14895))
          for _,r = range(r_support.Values)/* loop2:5 */{ 
            var void_try6 EID 
            _ = void_try6
            { 
            if (INT(Core.F_CALL(C_open,ARGS(r.ToEID()))) < 2) /* If:6 */{ 
              /* Let:7 */{ 
                var g1490UU *Language.Call  
                /* noccur = 1 */
                /* Let:8 */{ 
                  var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                  /* noccur = 5 */
                  _CL_obj.Selector = C_final
                  _CL_obj.Args = MakeConstantList(r)
                  g1490UU = _CL_obj
                  /* Let-8 */} 
                l = l.AddFast(g1490UU.Id())
                /* Let-7 */} 
              /* If-6 */} 
            void_try6 = F_Optimize_compile_if_write_relation(ToRelation(r))
            /* ERROR PROTECTION INSERTED (void_try6-void_try6) */
            if ErrorIn(void_try6) {Result = void_try6
            break
            } else {
            /* Let:6 */{ 
              var dn *ClaireAny   = ANY(Core.F_CALL(ToProperty(C__7_plus.Id()),ARGS(Core.F_CALL(C_name,ARGS(r.ToEID())),EID{MakeString("_write").Id(),0})))
              /* noccur = 1 */
              /* Let:7 */{ 
                var s *ClaireString   = ToSymbol(dn).String_I()
                /* noccur = 2 */
                /* Let:8 */{ 
                  var lb *ClaireAny   = ANY(Core.F_CALL(C_if_write,ARGS(r.ToEID())))
                  /* noccur = 1 */
                  void_try6 = F_Compile_compile_lambda_string(s,ToLambda(lb),C_void.Id())
                  /* ERROR PROTECTION INSERTED (void_try6-void_try6) */
                  if ErrorIn(void_try6) {Result = void_try6
                  break
                  } else {
                  /* Let:9 */{ 
                    var g1491UU *Language.Call  
                    /* noccur = 1 */
                    /* Let:10 */{ 
                      var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                      /* noccur = 5 */
                      _CL_obj.Selector = C_put
                      _CL_obj.Args = MakeConstantList(C_if_write.Id(),r,F_make_function_string(s).Id())
                      g1491UU = _CL_obj
                      /* Let-10 */} 
                    void_try6 = EID{l.AddFast(g1491UU.Id()).Id(),0}
                    /* Let-9 */} 
                  }
                  /* Let-8 */} 
                /* Let-7 */} 
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (void_try6-void_try6) */
            if ErrorIn(void_try6) {Result = void_try6
            break
            } else {
            }}
            }}
            /* loop-5 */} 
          /* For-4 */} 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        /* For:4 */{ 
          var r *ClaireAny  
          _ = r
          Result= EID{CFALSE.Id(),0}
          var r_support *ClaireSet  
          var r_support_try14925 EID 
          r_support_try14925 = Core.F_nth_table1(Language.C_Language_relations,ru)
          /* ERROR PROTECTION INSERTED (r_support-Result) */
          if ErrorIn(r_support_try14925) {Result = r_support_try14925
          } else {
          r_support = ToSet(OBJ(r_support_try14925))
          for _,r = range(r_support.Values)/* loop2:5 */{ 
            var void_try6 EID 
            _ = void_try6
            if (Language.F_eventMethod_ask_relation2(ToRelation(r)) == CTRUE) /* If:6 */{ 
              var l_try14937 EID 
              /* Let:7 */{ 
                var g1494UU *ClaireAny  
                /* noccur = 1 */
                var g1494UU_try14958 EID 
                g1494UU_try14958 = F_Optimize_compileEventMethod_property(ToProperty(r))
                /* ERROR PROTECTION INSERTED (g1494UU-l_try14937) */
                if ErrorIn(g1494UU_try14958) {l_try14937 = g1494UU_try14958
                } else {
                g1494UU = ANY(g1494UU_try14958)
                l_try14937 = EID{l.AddFast(g1494UU).Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (l-void_try6) */
              if ErrorIn(l_try14937) {void_try6 = l_try14937
              } else {
              l = ToList(OBJ(l_try14937))
              void_try6 = EID{l.Id(),0}
              }
              } else {
              void_try6 = EID{CFALSE.Id(),0}
              /* If-6 */} 
            /* ERROR PROTECTION INSERTED (void_try6-Result) */
            if ErrorIn(void_try6) {Result = void_try6
            Result = void_try6
            break
            } else {
            }}
            /* loop-5 */} 
          /* For-4 */} 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        /* Let:4 */{ 
          var g1496UU *Language.Do  
          /* noccur = 1 */
          /* Let:5 */{ 
            var _CL_obj *Language.Do   = Language.To_Do(new(Language.Do).Is(Language.C_Do))
            /* noccur = 3 */
            _CL_obj.Args = l
            g1496UU = _CL_obj
            /* Let-5 */} 
          Result = Core.F_CALL(C_c_code,ARGS(EID{g1496UU.Id(),0},EID{s.Id(),0}))
          /* Let-4 */} 
        }}}
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_code @ Defrule (throw: true) 
func E_c_code_Defrule (self EID,s EID) EID { 
    return /*(sm for c_code @ Defrule= EID)*/ F_c_code_Defrule(Language.To_Defrule(OBJ(self)),ToClass(OBJ(s)) )} 
  
// produce a beautiful if_write demon
/* {1} OPT.The go function for: compile_if_write(R:relation) [] */
func F_Optimize_compile_if_write_relation (R *ClaireRelation ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var l *ClaireAny  
      /* noccur = 2 */
      var l_try14973 EID 
      l_try14973 = Core.F_nth_table1(Language.C_demons,R.Id())
      /* ERROR PROTECTION INSERTED (l-Result) */
      if ErrorIn(l_try14973) {Result = l_try14973
      } else {
      l = ANY(l_try14973)
      /* Let:3 */{ 
        var lvar *ClaireList   = Language.ToLanguageDemon(ToList(l).At(1-1)).Formula.Vars
        /* noccur = 20 */
        /* Let:4 */{ 
          var l1 *ClaireList  
          /* noccur = 6 */
          var l1_try14985 EID 
          /* Construct:5 */{ 
            var v_bag_arg *ClaireAny  
            l1_try14985= EID{ToType(C_any.Id()).EmptyList().Id(),0}
            var v_bag_arg_try14996 EID 
            v_bag_arg_try14996 = Core.F_CALL(C_Optimize_Produce_put,ARGS(EID{R.Id(),0},lvar.At(1-1).ToEID(),lvar.At(2-1).ToEID()))
            /* ERROR PROTECTION INSERTED (v_bag_arg-l1_try14985) */
            if ErrorIn(v_bag_arg_try14996) {l1_try14985 = v_bag_arg_try14996
            } else {
            v_bag_arg = ANY(v_bag_arg_try14996)
            ToList(OBJ(l1_try14985)).AddFast(v_bag_arg)}
            /* Construct-5 */} 
          /* ERROR PROTECTION INSERTED (l1-Result) */
          if ErrorIn(l1_try14985) {Result = l1_try14985
          } else {
          l1 = ToList(OBJ(l1_try14985))
          /* Let:5 */{ 
            var l2 *ClaireList  
            /* noccur = 10 */
            /* Iteration:6 */{ 
              var v_list6 *ClaireList  
              var x *ClaireAny  
              var v_local6 *ClaireAny  
              v_list6 = ToList(l)
              l2 = CreateList(ToType(C_any.Id()),v_list6.Length())
              for CLcount := 0; CLcount < v_list6.Length(); CLcount++{ 
                x = v_list6.At(CLcount)
                v_local6 = Language.F_substitution_any(Language.F_substitution_any(Language.F_substitution_any(Language.ToLanguageDemon(x).Formula.Body,To_Variable(Language.ToLanguageDemon(x).Formula.Vars.At(3-1)),lvar.At(3-1)),To_Variable(Language.ToLanguageDemon(x).Formula.Vars.At(1-1)),lvar.At(1-1)),To_Variable(Language.ToLanguageDemon(x).Formula.Vars.At(2-1)),lvar.At(2-1))
                l2.PutAt(CLcount,v_local6)
                } 
              /* Iteration-6 */} 
            Result = Core.F_put_property2(C_range,ToObject(lvar.At(1-1)),R.Domain.Id())
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            Result = Core.F_put_property2(C_range,ToObject(lvar.At(2-1)),R.Range.Id())
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            /* For:6 */{ 
              var v *ClaireAny  
              _ = v
              Result= EID{CFALSE.Id(),0}
              var v_support *ClaireList  
              v_support = lvar
              v_len := v_support.Length()
              for i_it := 0; i_it < v_len; i_it++ { 
                v = v_support.At(i_it)
                var void_try8 EID 
                _ = void_try8
                void_try8 = Core.F_put_property2(C_range,ToObject(v),ToTypeExpression(OBJ(Core.F_CALL(C_range,ARGS(v.ToEID())))).Class_I().Id())
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
            if ((l2.At(1-1).Isa.IsIn(Language.C_If) == CTRUE) && 
                (Language.F_eventMethod_ask_relation2(R) != CTRUE)) /* If:6 */{ 
              if (Core.F_owner_any(Language.To_If(l2.At(1-1)).Test).IsIn(Language.C_And) == CTRUE) /* If:7 */{ 
                /* update:8 */{ 
                  var va_arg1 *Language.If  
                  var va_arg2 *ClaireAny  
                  va_arg1 = Language.To_If(l2.At(1-1))
                  var va_arg2_try15009 EID 
                  /* Let:9 */{ 
                    var _CL_obj *Language.And   = Language.To_And(new(Language.And).Is(Language.C_And))
                    /* noccur = 3 */
                    /* update:10 */{ 
                      var va_arg1 *Language.And  
                      var va_arg2 *ClaireList  
                      va_arg1 = _CL_obj
                      var va_arg2_try150111 EID 
                      va_arg2_try150111 = ToList(OBJ(Core.F_CALL(C_args,ARGS(Language.To_If(l2.At(1-1)).Test.ToEID())))).Cdr()
                      /* ERROR PROTECTION INSERTED (va_arg2-va_arg2_try15009) */
                      if ErrorIn(va_arg2_try150111) {va_arg2_try15009 = va_arg2_try150111
                      } else {
                      va_arg2 = ToList(OBJ(va_arg2_try150111))
                      /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                      va_arg1.Args = va_arg2
                      va_arg2_try15009 = EID{va_arg2.Id(),0}
                      }
                      /* update-10 */} 
                    /* ERROR PROTECTION INSERTED (va_arg2_try15009-va_arg2_try15009) */
                    if !ErrorIn(va_arg2_try15009) {
                    va_arg2_try15009 = EID{_CL_obj.Id(),0}
                    }
                    /* Let-9 */} 
                  /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                  if ErrorIn(va_arg2_try15009) {Result = va_arg2_try15009
                  } else {
                  va_arg2 = ANY(va_arg2_try15009)
                  /* ---------- now we compile update iClaire/test(va_arg1) := va_arg2 ------- */
                  va_arg1.Test = va_arg2
                  Result = va_arg2.ToEID()
                  }
                  /* update-8 */} 
                } else {
                Result = ToArray(l2.Id()).NthPut(1,Language.To_If(l2.At(1-1)).Arg).ToEID()
                /* If-7 */} 
              } else {
              Result = EID{CFALSE.Id(),0}
              /* If-6 */} 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            if (R.Inverse.Id() != CNULL) /* If:6 */{ 
              if (F_boolean_I_any(ANY(Core.F_CALL(C_Optimize_multi_ask,ARGS(EID{R.Id(),0})))).Id() != CTRUE.Id()) /* If:7 */{ 
                var l1_try15028 EID 
                /* Let:8 */{ 
                  var g1503UU *ClaireAny  
                  /* noccur = 1 */
                  var g1503UU_try15049 EID 
                  g1503UU_try15049 = Core.F_CALL(C_Optimize_Produce_remove,ARGS(EID{R.Inverse.Id(),0},lvar.At(3-1).ToEID(),lvar.At(1-1).ToEID()))
                  /* ERROR PROTECTION INSERTED (g1503UU-l1_try15028) */
                  if ErrorIn(g1503UU_try15049) {l1_try15028 = g1503UU_try15049
                  } else {
                  g1503UU = ANY(g1503UU_try15049)
                  l1_try15028 = EID{l1.AddFast(g1503UU).Id(),0}
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (l1-Result) */
                if ErrorIn(l1_try15028) {Result = l1_try15028
                } else {
                l1 = ToList(OBJ(l1_try15028))
                Result = EID{l1.Id(),0}
                }
                } else {
                Result = EID{CFALSE.Id(),0}
                /* If-7 */} 
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              var l1_try15057 EID 
              /* Let:7 */{ 
                var g1506UU *ClaireAny  
                /* noccur = 1 */
                var g1506UU_try15078 EID 
                g1506UU_try15078 = Core.F_CALL(C_Optimize_Produce_put,ARGS(EID{R.Inverse.Id(),0},lvar.At(2-1).ToEID(),lvar.At(1-1).ToEID()))
                /* ERROR PROTECTION INSERTED (g1506UU-l1_try15057) */
                if ErrorIn(g1506UU_try15078) {l1_try15057 = g1506UU_try15078
                } else {
                g1506UU = ANY(g1506UU_try15078)
                l1_try15057 = EID{l1.AddFast(g1506UU).Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (l1-Result) */
              if ErrorIn(l1_try15057) {Result = l1_try15057
              } else {
              l1 = ToList(OBJ(l1_try15057))
              Result = EID{l1.Id(),0}
              }
              }
              } else {
              Result = EID{CFALSE.Id(),0}
              /* If-6 */} 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            /* update:6 */{ 
              var va_arg1 *ClaireRelation  
              var va_arg2 *ClaireAny  
              va_arg1 = R
              var va_arg2_try15087 EID 
              /* Let:7 */{ 
                var g1509UU *Language.ComplexInstruction  
                /* noccur = 1 */
                if (Language.F_eventMethod_ask_relation2(R) == CTRUE) /* If:8 */{ 
                  /* Let:9 */{ 
                    var _CL_obj *Language.Do   = Language.To_Do(new(Language.Do).Is(Language.C_Do))
                    /* noccur = 3 */
                    _CL_obj.Args = l2
                    g1509UU = Language.To_ComplexInstruction(_CL_obj.Id())
                    /* Let-9 */} 
                  /* If!8 */}  else if (F_boolean_I_any(ANY(Core.F_CALL(C_Optimize_multi_ask,ARGS(EID{R.Id(),0})))) == CTRUE) /* If:8 */{ 
                  /* Let:9 */{ 
                    var _CL_obj *Language.If   = Language.To_If(new(Language.If).Is(Language.C_If))
                    /* noccur = 21 */
                    /* update:10 */{ 
                      var va_arg1 *Language.If  
                      var va_arg2 *ClaireAny  
                      va_arg1 = _CL_obj
                      /* Let:11 */{ 
                        var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                        /* noccur = 11 */
                        _CL_obj.Selector = Core.C_not
                        /* update:12 */{ 
                          var va_arg1 *Language.Call  
                          var va_arg2 *ClaireList  
                          va_arg1 = _CL_obj
                          /* Construct:13 */{ 
                            var v_bag_arg *ClaireAny  
                            va_arg2= ToType(CEMPTY.Id()).EmptyList()
                            /* Let:14 */{ 
                              var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                              /* noccur = 5 */
                              _CL_obj.Selector = ToProperty(C__Z.Id())
                              _CL_obj.Args = MakeConstantList(lvar.At(2-1),Language.F_readCall_relation(R,lvar.At(1-1)).Id())
                              v_bag_arg = _CL_obj.Id()
                              /* Let-14 */} 
                            va_arg2.AddFast(v_bag_arg)/* Construct-13 */} 
                          /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                          va_arg1.Args = va_arg2
                          /* update-12 */} 
                        va_arg2 = _CL_obj.Id()
                        /* Let-11 */} 
                      /* ---------- now we compile update iClaire/test(va_arg1) := va_arg2 ------- */
                      va_arg1.Test = va_arg2
                      /* update-10 */} 
                    /* update:10 */{ 
                      var va_arg1 *Language.If  
                      var va_arg2 *ClaireAny  
                      va_arg1 = _CL_obj
                      /* Let:11 */{ 
                        var _CL_obj *Language.Do   = Language.To_Do(new(Language.Do).Is(Language.C_Do))
                        /* noccur = 3 */
                        _CL_obj.Args = l1.Append(l2)
                        va_arg2 = _CL_obj.Id()
                        /* Let-11 */} 
                      /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                      va_arg1.Arg = va_arg2
                      /* update-10 */} 
                    g1509UU = Language.To_ComplexInstruction(_CL_obj.Id())
                    /* Let-9 */} 
                  } else {
                  /* Let:9 */{ 
                    var _CL_obj *Language.Let   = Language.To_Let(new(Language.Let).Is(Language.C_Let))
                    /* noccur = 22 */
                    _CL_obj.ClaireVar = To_Variable(lvar.At(3-1))
                    _CL_obj.Value = Language.F_readCall_relation(R,lvar.At(1-1)).Id()
                    /* update:10 */{ 
                      var va_arg1 *Language.Let  
                      var va_arg2 *ClaireAny  
                      va_arg1 = _CL_obj
                      /* Let:11 */{ 
                        var _CL_obj *Language.If   = Language.To_If(new(Language.If).Is(Language.C_If))
                        /* noccur = 15 */
                        /* update:12 */{ 
                          var va_arg1 *Language.If  
                          var va_arg2 *ClaireAny  
                          va_arg1 = _CL_obj
                          /* Let:13 */{ 
                            var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                            /* noccur = 5 */
                            _CL_obj.Selector = ToProperty(Core.C__I_equal.Id())
                            _CL_obj.Args = MakeConstantList(lvar.At(2-1),lvar.At(3-1))
                            va_arg2 = _CL_obj.Id()
                            /* Let-13 */} 
                          /* ---------- now we compile update iClaire/test(va_arg1) := va_arg2 ------- */
                          va_arg1.Test = va_arg2
                          /* update-12 */} 
                        /* update:12 */{ 
                          var va_arg1 *Language.If  
                          var va_arg2 *ClaireAny  
                          va_arg1 = _CL_obj
                          /* Let:13 */{ 
                            var _CL_obj *Language.Do   = Language.To_Do(new(Language.Do).Is(Language.C_Do))
                            /* noccur = 3 */
                            _CL_obj.Args = l1.Append(l2)
                            va_arg2 = _CL_obj.Id()
                            /* Let-13 */} 
                          /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                          va_arg1.Arg = va_arg2
                          /* update-12 */} 
                        va_arg2 = _CL_obj.Id()
                        /* Let-11 */} 
                      /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                      va_arg1.Arg = va_arg2
                      /* update-10 */} 
                    g1509UU = Language.To_ComplexInstruction(_CL_obj.Id())
                    /* Let-9 */} 
                  /* If-8 */} 
                va_arg2_try15087 = Language.F_lambda_I_list(MakeConstantList(lvar.At(1-1),lvar.At(2-1)),g1509UU.Id())
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (va_arg2-Result) */
              if ErrorIn(va_arg2_try15087) {Result = va_arg2_try15087
              } else {
              va_arg2 = ANY(va_arg2_try15087)
              /* ---------- now we compile update if_write(va_arg1) := va_arg2 ------- */
              va_arg1.IfWrite = va_arg2
              Result = va_arg2.ToEID()
              }
              /* update-6 */} 
            }}}}}
            /* Let-5 */} 
          }
          /* Let-4 */} 
        /* Let-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: compile_if_write @ relation (throw: true) 
func E_Optimize_compile_if_write_relation (R EID) EID { 
    return /*(sm for compile_if_write @ relation= EID)*/ F_Optimize_compile_if_write_relation(ToRelation(OBJ(R)) )} 
  
// create a simple method that will trigger the event
/* {1} OPT.The go function for: compileEventMethod(p:property) [] */
func F_Optimize_compileEventMethod_property (p *ClaireProperty ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var m *ClaireMethod   = ToMethod(p.Restrictions.ValuesO()[1-1])
      /* noccur = 1 */
      /* Let:3 */{ 
        var na *ClaireString   = F_append_string(p.Name.String_I(),MakeString("_write"))
        /* noccur = 1 */
        Result = F_Optimize_add_method_I_method(m,
          MakeConstantList(p.Domain.Id(),p.Range.Id()),
          C_void.Id(),
          MakeInteger(0).Id(),
          F_make_function_string(na))
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: compileEventMethod @ property (throw: true) 
func E_Optimize_compileEventMethod_property (p EID) EID { 
    return /*(sm for compileEventMethod @ property= EID)*/ F_Optimize_compileEventMethod_property(ToProperty(OBJ(p)) )} 
  