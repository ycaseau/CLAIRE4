/***** CLAIRE Compilation of file /Users/ycaseau/Dropbox/src/clairev4.0/src/compile/gostat.cl 
         [version 4.0.01 / safety 5] Saturday 10-30-2021 *****/

package Generate
import (_ "fmt"
	. "Kernel"
	"Core"
	"Language"
	"Reader"
	"Optimize"
)

//-------- dumb function to prevent import errors --------
func import_g0535() { 
    _ = Core.It
    _ = Language.It
    _ = Reader.It
    _ = Optimize.It
    } 
  
  
//+-------------------------------------------------------------+
//| CLAIRE                                                      |
//| gostat.cl                                                   |
//| Copyright (C) 2020-2021   Yves Caseau. All Rights Reserved  |
//| cf. copyright info in file object.cl: about()               |
//+-------------------------------------------------------------+
// statement is implemented as a general method that calls a restriction
//        g_statement(self:any,e:class,v:string,err:boolean,loop:any)
// (1) e is the goType that the variable v must receive (HENCE goCast must be inserted)
//     a proper goType is a class, or EID, or void
// (2) The argument v is the named of the C variable in which the
//     result of the evaluation must be placed.
// (3) err tells if an error is possible, which forces to create a chain an not a block (see Do for example)
//     Note : if err = true, s is expected to be EID to (a) force a chain (b) place the error value in v
// (4) loop is either false (not within a loop) or a tuple(v,s) inside the compiling of While/For
//     This tuple describes the vreturn Variable in case a break(v) is encoutered
// there are two possible outputs: blocks (lines of code without {}, used to be call inner_statement)
// and chains  (we use chains to denote long nested ifs that manage error handling)
// indentation : 
//    we call statement(s) at the proper current indentation level => it produices n lines with the indentation
//    and stop after a break line, at the proper identation level
//**********************************************************************
//*  Table of contents:                                                *
//*          Part 1: Unfolding of complex expressions                  *
//*          Part 2: Error Management and EID Unfolding                *
//*          Part 3: Basic control structures                          *
//*          Part 4: iteration                                         *
//*          Part 5: CLAIRE-specific structures                        *
//**********************************************************************
//*************************************************************************
//*          Part 1: Unfolding of complex expressions                     *
//*************************************************************************
// when local CLAIRE expressions are not go expression, we need to unfold the global expression into a big Let
// HOWEVER, if only works for list of arguments whose evaluation order is not specified ! (because we move some of the evaluations earlier)
// this reintrant compiling (calling g_statement on a expanded Let) works because Let checks if g_expression can be used
// the same pattern is used for call_slot/call_table
// a clean expression is both a functional expression and one that does not throw an error
/* {1} OPT.The go function for: g_clean(x:any) [] */
func F_Generate_g_clean_any (x *ClaireAny ) EID { 
    var Result EID 
    /* and:2 */{ 
      var v_and2 *ClaireBoolean  
      
      var v_and2_try05363 EID 
      v_and2_try05363 = F_Generate_g_func_any(x)
      /* ERROR PROTECTION INSERTED (v_and2-Result) */
      if ErrorIn(v_and2_try05363) {Result = v_and2_try05363
      } else {
      v_and2 = ToBoolean(OBJ(v_and2_try05363))
      if (v_and2 == CFALSE) {Result = EID{CFALSE.Id(),0}
      } else /* arg:3 */{ 
        var v_and2_try05374 EID 
        /* Let:4 */{ 
          var g0538UU *ClaireBoolean  
          /* noccur = 1 */
          var g0538UU_try05395 EID 
          g0538UU_try05395 = Optimize.F_Compile_g_throw_any(x)
          /* ERROR PROTECTION INSERTED (g0538UU-v_and2_try05374) */
          if ErrorIn(g0538UU_try05395) {v_and2_try05374 = g0538UU_try05395
          } else {
          g0538UU = ToBoolean(OBJ(g0538UU_try05395))
          v_and2_try05374 = EID{g0538UU.Not.Id(),0}
          }
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (v_and2-Result) */
        if ErrorIn(v_and2_try05374) {Result = v_and2_try05374
        } else {
        v_and2 = ToBoolean(OBJ(v_and2_try05374))
        if (v_and2 == CFALSE) {Result = EID{CFALSE.Id(),0}
        } else /* arg:4 */{ 
          Result = EID{CTRUE.Id(),0}/* arg-4 */} 
        /* arg-3 */} 
      }}
      /* and-2 */} 
    return Result} 
  
// The EID go function for: g_clean @ any (throw: true) 
func E_Generate_g_clean_any (x EID) EID { 
    return /*(sm for g_clean @ any= EID)*/ F_Generate_g_clean_any(ANY(x) )} 
  
// this function is used to unfold complex expressions that should be compiled as
// expressions and not statements. It takes a list of arguments l and returns the
// embedded Lets that defines the necessary variable or nil (nothing is needed)
// this list is of the form  (a1 .. am) where
//     m is the number of statement args in l
//     ai is a Let that defines the i-th variable corresponding to the i-th bad guy
// CLAIRE 4: we unfold args that are not functional or args that can throw error
/* {1} OPT.The go function for: unfold_args(l:list) [] */
func F_Generate_unfold_args_list (l *ClaireList ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var lbad *ClaireList  
      /* noccur = 1 */
      var lbad_try05413 EID 
      /* Let:3 */{ 
        var i_out *ClaireList   = ToType(CEMPTY.Id()).EmptyList()
        /* noccur = 2 */
        /* Let:4 */{ 
          var i int  = 1
          /* noccur = 5 */
          /* Let:5 */{ 
            var g0540 int  = l.Length()
            /* noccur = 1 */
            lbad_try05413= EID{CFALSE.Id(),0}
            for (i <= g0540) /* while:6 */{ 
              var void_try7 EID 
              _ = void_try7
              { 
              var g0542I *ClaireBoolean  
              var g0542I_try05437 EID 
              /* Let:7 */{ 
                var g0544UU *ClaireBoolean  
                /* noccur = 1 */
                var g0544UU_try05458 EID 
                g0544UU_try05458 = F_Generate_g_clean_any(l.At(i-1))
                /* ERROR PROTECTION INSERTED (g0544UU-g0542I_try05437) */
                if ErrorIn(g0544UU_try05458) {g0542I_try05437 = g0544UU_try05458
                } else {
                g0544UU = ToBoolean(OBJ(g0544UU_try05458))
                g0542I_try05437 = EID{g0544UU.Not.Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (g0542I-void_try7) */
              if ErrorIn(g0542I_try05437) {void_try7 = g0542I_try05437
              } else {
              g0542I = ToBoolean(OBJ(g0542I_try05437))
              if (g0542I == CTRUE) /* If:7 */{ 
                void_try7 = EID{i_out.AddFast(MakeInteger(i).Id()).Id(),0}
                } else {
                void_try7 = EID{CFALSE.Id(),0}
                /* If-7 */} 
              }
              /* ERROR PROTECTION INSERTED (void_try7-void_try7) */
              if ErrorIn(void_try7) {lbad_try05413 = void_try7
              break
              } else {
              i = (i+1)
              }
              /* while-6 */} 
            }
            /* Let-5 */} 
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (lbad_try05413-lbad_try05413) */
        if !ErrorIn(lbad_try05413) {
        lbad_try05413 = EID{i_out.Id(),0}
        }
        /* Let-3 */} 
      /* ERROR PROTECTION INSERTED (lbad-Result) */
      if ErrorIn(lbad_try05413) {Result = lbad_try05413
      } else {
      lbad = ToList(OBJ(lbad_try05413))
      
      /* Iteration:3 */{ 
        var v_list3 *ClaireList  
        var i int 
        var v_local3 *ClaireAny  
        v_list3 = lbad
        Result = EID{CreateList(ToType(CEMPTY.Id()),v_list3.Length()).Id(),0}
        for CLcount := 0; CLcount < v_list3.Length(); CLcount++{ 
          i = ToInteger(v_list3.At(CLcount)).Value
          var v_local3_try05465 EID 
          /* Let:5 */{ 
            var _CL_obj *Language.Let   = Language.To_Let(new(Language.Let).Is(Language.C_Let))
            /* noccur = 4 */
            /* update:6 */{ 
              var va_arg1 *Language.InstructionWithVar  
              var va_arg2 *ClaireVariable  
              va_arg1 = Language.To_InstructionWithVar(_CL_obj.Id())
              var va_arg2_try05477 EID 
              /* Let:7 */{ 
                var g0548UU *ClaireClass  
                /* noccur = 1 */
                var g0548UU_try05498 EID 
                g0548UU_try05498 = Language.F_static_type_any(l.At(i-1))
                /* ERROR PROTECTION INSERTED (g0548UU-va_arg2_try05477) */
                if ErrorIn(g0548UU_try05498) {va_arg2_try05477 = g0548UU_try05498
                } else {
                g0548UU = ToClass(OBJ(g0548UU_try05498))
                va_arg2_try05477 = EID{Optimize.F_Compile_Variable_I_symbol(F_append_symbol(Core.F_gensym_void(),MakeString("UU").Id()),0,g0548UU.Id()).Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (va_arg2-v_local3_try05465) */
              if ErrorIn(va_arg2_try05477) {v_local3_try05465 = va_arg2_try05477
              } else {
              va_arg2 = To_Variable(OBJ(va_arg2_try05477))
              /* ---------- now we compile update var(va_arg1) := va_arg2 ------- */
              va_arg1.ClaireVar = va_arg2
              v_local3_try05465 = EID{va_arg2.Id(),0}
              }
              /* update-6 */} 
            /* ERROR PROTECTION INSERTED (v_local3_try05465-v_local3_try05465) */
            if ErrorIn(v_local3_try05465) {Result = v_local3_try05465
            break
            } else {
            _CL_obj.Value = l.At(i-1)
            v_local3_try05465 = EID{_CL_obj.Id(),0}
            }
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (v_local3-Result) */
          if ErrorIn(v_local3_try05465) {Result = v_local3_try05465
          Result = v_local3_try05465
          break
          } else {
          v_local3 = ANY(v_local3_try05465)
          ToList(OBJ(Result)).PutAt(CLcount,v_local3)
          } 
        }
        /* Iteration-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: unfold_args @ list (throw: true) 
func E_Generate_unfold_args_list (l EID) EID { 
    return /*(sm for unfold_args @ list= EID)*/ F_Generate_unfold_args_list(ToList(OBJ(l)) )} 
  
// uses the previous list to use the variable instead of the Fold.
// l is the list of arguments, ld is the previously build unfold_args(l)
/* {1} OPT.The go function for: unfold_arg(l:list,ld:list,x:any) [] */
func F_Generate_unfold_arg_list (l *ClaireList ,ld *ClaireList ,x *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var i int  = 1
      /* noccur = 7 */
      /* Let:3 */{ 
        var j int  = 0
        /* noccur = 3 */
        /* Let:4 */{ 
          var m int  = l.Length()
          /* noccur = 1 */
          var g0550I *ClaireBoolean  
          var g0550I_try05515 EID 
          g0550I_try05515= EID{CFALSE.Id(),0}
          for (i <= m) /* while:5 */{ 
            var void_try6 EID 
            _ = void_try6
            { 
            var g0552I *ClaireBoolean  
            var g0552I_try05536 EID 
            /* Let:6 */{ 
              var g0554UU *ClaireBoolean  
              /* noccur = 1 */
              var g0554UU_try05557 EID 
              g0554UU_try05557 = F_Generate_g_clean_any(l.At(i-1))
              /* ERROR PROTECTION INSERTED (g0554UU-g0552I_try05536) */
              if ErrorIn(g0554UU_try05557) {g0552I_try05536 = g0554UU_try05557
              } else {
              g0554UU = ToBoolean(OBJ(g0554UU_try05557))
              g0552I_try05536 = EID{g0554UU.Not.Id(),0}
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (g0552I-void_try6) */
            if ErrorIn(g0552I_try05536) {void_try6 = g0552I_try05536
            } else {
            g0552I = ToBoolean(OBJ(g0552I_try05536))
            if (g0552I == CTRUE) /* If:6 */{ 
              j = (j+1)
              if (Equal(l.At(i-1),x) == CTRUE) /* If:7 */{ 
                 /*v = g0550I_try05515, s =EID*/
g0550I_try05515 = EID{CTRUE.Id(),0}
                break
                } else {
                void_try6 = EID{CFALSE.Id(),0}
                /* If-7 */} 
              /* If!6 */}  else if (Equal(l.At(i-1),x) == CTRUE) /* If:6 */{ 
               /*v = g0550I_try05515, s =EID*/
g0550I_try05515 = EID{CFALSE.Id(),0}
              break
              } else {
              void_try6 = EID{CFALSE.Id(),0}
              /* If-6 */} 
            }
            /* ERROR PROTECTION INSERTED (void_try6-void_try6) */
            if ErrorIn(void_try6) {g0550I_try05515 = void_try6
            break
            } else {
            i = (i+1)
            }
            /* while-5 */} 
          }
          /* ERROR PROTECTION INSERTED (g0550I-Result) */
          if ErrorIn(g0550I_try05515) {Result = g0550I_try05515
          } else {
          g0550I = ToBoolean(OBJ(g0550I_try05515))
          if (g0550I == CTRUE) /* If:5 */{ 
            Result = Core.F_CALL(Language.C_var,ARGS(ld.At(j-1).ToEID()))
            } else {
            Result = l.At(i-1).ToEID()
            /* If-5 */} 
          }
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: unfold_arg @ list (throw: true) 
func E_Generate_unfold_arg_list (l EID,ld EID,x EID) EID { 
    return /*(sm for unfold_arg @ list= EID)*/ F_Generate_unfold_arg_list(ToList(OBJ(l)),ToList(OBJ(ld)),ANY(x) )} 
  
// creates the Let from the ldef definition and places the statement x in the body
// note that the error handling is done in the Let (with g_statement)
// x is the call form where the variable has been replaced if needed
/* {1} OPT.The go function for: unfold_use(ldef:list,x:any,s:class,v:string,err:boolean,loop:any) [] */
func F_Generate_unfold_use_list (ldef *ClaireList ,x *ClaireAny ,s *ClaireClass ,v *ClaireString ,err *ClaireBoolean ,loop *ClaireAny ) EID { 
    var Result EID 
    if (F_Generate_eid_require_ask_any(x) == CTRUE) /* If:2 */{ 
      Result = F_Generate_unfold_eid_list(ldef,
        x,
        s,
        (v).Id(),
        err,
        loop)
      } else {
      /* Let:3 */{ 
        var n int  = ldef.Length()
        /* noccur = 3 */
        /* Let:4 */{ 
          var vb int  = ClEnv.Verbose
          /* noccur = 1 */
          ClEnv.Verbose = 0
          if (F_boolean_I_any(ldef.Id()).Id() != CTRUE.Id()) /* If:5 */{ 
            Result = ToException(Core.C_general_error.Make(MakeString("[internal] design bug g_func(~S) should be true").Id(),MakeConstantList(x).Id())).Close()
            } else {
            Result = EID{CFALSE.Id(),0}
            /* If-5 */} 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          /* Let:5 */{ 
            var i int  = 1
            /* noccur = 5 */
            /* Let:6 */{ 
              var g0556 int  = (n-1)
              /* noccur = 1 */
              Result= EID{CFALSE.Id(),0}
              for (i <= g0556) /* while:7 */{ 
                var void_try8 EID 
                _ = void_try8
                { 
                void_try8 = Core.F_write_property(C_arg,ToObject(ldef.At(i-1)),ldef.At((i+1)-1))
                /* ERROR PROTECTION INSERTED (void_try8-void_try8) */
                if ErrorIn(void_try8) {Result = void_try8
                break
                } else {
                i = (i+1)
                }
                /* while-7 */} 
              }
              /* Let-6 */} 
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          Language.To_Let(ldef.At(n-1)).Arg = x
          ClEnv.Verbose = vb
          
          Result = F_Generate_g_statement_Let(Language.To_Let(ldef.At(1-1)),
            s,
            v,
            err,
            loop)
          }}
          /* Let-4 */} 
        /* Let-3 */} 
      /* If-2 */} 
    return Result} 
  
// The EID go function for: unfold_use @ list (throw: true) 
func E_Generate_unfold_use_list (ldef EID,x EID,s EID,v EID,err EID,loop EID) EID { 
    return /*(sm for unfold_use @ list= EID)*/ F_Generate_unfold_use_list(ToList(OBJ(ldef)),
      ANY(x),
      ToClass(OBJ(s)),
      ToString(OBJ(v)),
      ToBoolean(OBJ(err)),
      ANY(loop) )} 
  
//*************************************************************************
//*          Part 2: Error Management and EID Unfolding                   *
//*************************************************************************
// this is the error catching pattern: evaluate(self) and check if error then place it in vglobal,
// if no error we want the value in v with expected gotype e (a true gotype = class)
// if v2 is an EID variable, do not create an extra variable (we use it temporarily)
// in a loop we generate a break to exit to loop
// v is the variable that must receive self
// note : g_try produces a pattern   <e = code>, if Err(e) {res =e} else { ...
// that must be closed } with a close_try => and nothing after the close_try (nothing must if an error occured)
/* {1} OPT.The go function for: g_try(self:any,v:string,e:class,vglobal:string,loop:any) [] */
func F_Generate_g_try_any (self *ClaireAny ,v *ClaireString ,e *ClaireClass ,vglobal *ClaireString ,loop *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var v2 *ClaireString  
      /* noccur = 8 */
      if (e.Id() == Optimize.C_EID.Id()) /* If:3 */{ 
        v2 = v
        } else {
        v2 = F_Generate_check_var_string(F_append_string(v,F_gensym_string(MakeString("_try")).String_I()))
        /* If-3 */} 
      if (e.Id() != Optimize.C_EID.Id()) /* If:3 */{ 
        F_Generate_var_declaration_string(v2,Optimize.C_EID,1)
        /* If-3 */} 
      Result = Core.F_CALL(C_Generate_g_statement,ARGS(self.ToEID(),
        EID{Optimize.C_EID.Id(),0},
        EID{(v2).Id(),0},
        EID{CTRUE.Id(),0},
        loop.ToEID()))
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      if ((self.Isa.IsIn(Language.C_Do) == CTRUE) && 
          (F_boolean_I_any(loop) == CTRUE)) /* If:3 */{ 
        PRINC("{")
        F_Generate_breakline_void()
        PRINC("")
        Result = EVOID
        } else {
        PRINC("/* ERROR PROTECTION INSERTED (")
        F_princ_string(v)
        PRINC("-")
        F_princ_string(vglobal)
        PRINC(") */")
        F_Generate_breakline_void()
        PRINC("")
        if ((v.Value == vglobal.Value) && 
            ((e.Id() == Optimize.C_EID.Id()) && 
              (C_tuple.Id() != loop.Isa.Id()))) /* If:4 */{ 
          PRINC("if !ErrorIn(")
          F_c_princ_string(v2)
          PRINC(") {")
          F_Generate_breakline_void()
          PRINC("")
          Result = EVOID
          } else {
          PRINC("if ErrorIn(")
          F_c_princ_string(v2)
          PRINC(") {")
          if (v.Value != vglobal.Value) /* If:5 */{ 
            F_c_princ_string(vglobal)
            PRINC(" = ")
            F_c_princ_string(v2)
            F_Generate_breakline_void()
            PRINC("")
            /* If-5 */} 
          if (C_tuple.Id() == loop.Isa.Id()) /* If:5 */{ 
            /* Let:6 */{ 
              var g0562UU *ClaireAny  
              /* noccur = 1 */
              var g0562UU_try05637 EID 
              g0562UU_try05637 = Core.F_CALL(C_nth,ARGS(loop.ToEID(),EID{C__INT,IVAL(1)}))
              /* ERROR PROTECTION INSERTED (g0562UU-Result) */
              if ErrorIn(g0562UU_try05637) {Result = g0562UU_try05637
              } else {
              g0562UU = ANY(g0562UU_try05637)
              Result = Core.F_CALL(C_c_princ,ARGS(g0562UU.ToEID()))
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC(" = ")
            F_c_princ_string(v2)
            F_Generate_breakline_void()
            PRINC("")
            Result = EVOID
            }
            {
            PRINC("break")
            Result = F_Generate_breakline_void().ToEID()
            }
            } else {
            Result = EID{CFALSE.Id(),0}
            /* If-5 */} 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC("} else {")
          F_Generate_breakline_void()
          PRINC("")
          Result = EVOID
          }
          /* If-4 */} 
        /* If-3 */} 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      if ((e.Id() != C_void.Id()) && 
          (v.Value != v2.Value)) /* If:3 */{ 
        F_c_princ_string(v)
        PRINC(" = ")
        Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).FromEid(v2,e)
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        F_Generate_breakline_void()
        PRINC("")
        Result = EVOID
        }
        } else {
        Result = EID{CFALSE.Id(),0}
        /* If-3 */} 
      }}
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: g_try @ any (throw: true) 
func E_Generate_g_try_any (self EID,v EID,e EID,vglobal EID,loop EID) EID { 
    return /*(sm for g_try @ any= EID)*/ F_Generate_g_try_any(ANY(self),
      ToString(OBJ(v)),
      ToClass(OBJ(e)),
      ToString(OBJ(vglobal)),
      ANY(loop) )} 
  
// when the error is nested in the expression, the unfold pattern will make sure that we separate the sub_exp that may
// create the error, but assignment is not managed this way, hence this code to avoid double error check
/* {1} OPT.The go function for: g_try(self:Assign,v:string,e:class,vglobal:string,loop:any) [] */
func F_Generate_g_try_Assign (self *Language.Assign ,v *ClaireString ,e *ClaireClass ,vglobal *ClaireString ,loop *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var _Zvar *ClaireAny   = self.ClaireVar
      /* noccur = 2 */
      /* Let:3 */{ 
        var v1 *ClaireAny  
        /* noccur = 2 */
        var v1_try05644 EID 
        v1_try05644 = Core.F_CALL(C_Generate_c_string,ARGS(EID{Optimize.C_PRODUCER.Value,0},_Zvar.ToEID()))
        /* ERROR PROTECTION INSERTED (v1-Result) */
        if ErrorIn(v1_try05644) {Result = v1_try05644
        } else {
        v1 = ANY(v1_try05644)
        /* Let:4 */{ 
          var _Zrange *ClaireClass   = ToTypeExpression(OBJ(Core.F_CALL(C_range,ARGS(_Zvar.ToEID())))).Class_I()
          /* noccur = 2 */
          Result = Core.F_CALL(C_Generate_g_try,ARGS(self.Arg.ToEID(),
            v1.ToEID(),
            EID{_Zrange.Id(),0},
            EID{(vglobal).Id(),0},
            loop.ToEID()))
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          if (e.Id() != C_void.Id()) /* If:5 */{ 
            F_c_princ_string(v)
            PRINC(" = ")
            Result = F_Generate_use_variable_string(ToString(v1),e,_Zrange)
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            F_Generate_breakline_void()
            PRINC("")
            Result = EVOID
            }
            } else {
            Result = EID{CFALSE.Id(),0}
            /* If-5 */} 
          }
          /* Let-4 */} 
        }
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: g_try @ Assign (throw: true) 
func E_Generate_g_try_Assign (self EID,v EID,e EID,vglobal EID,loop EID) EID { 
    return /*(sm for g_try @ Assign= EID)*/ F_Generate_g_try_Assign(Language.To_Assign(OBJ(self)),
      ToString(OBJ(v)),
      ToClass(OBJ(e)),
      ToString(OBJ(vglobal)),
      ANY(loop) )} 
  
// each g_try produces a {, which we must balance before returning a new line 
// does NOT change OPT.level !
/* {1} OPT.The go function for: close_try(n:integer) [] */
func F_Generate_close_try_integer (n int)  { 
    // procedure body with s =  
/* Let:2 */{ 
      var i int  = 1
      /* noccur = 3 */
      /* Let:3 */{ 
        var g0565 int  = n
        /* noccur = 1 */
        for (i <= g0565) /* while:4 */{ 
          PRINC("}")
          i = (i+1)
          /* while-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    if (n > 0) /* If:2 */{ 
      F_Generate_breakline_void()
      /* If-2 */} 
    } 
  
// The EID go function for: close_try @ integer (throw: false) 
func E_Generate_close_try_integer (n EID) EID { 
    /*(sm for close_try @ integer= void)*/ F_Generate_close_try_integer(INT(n) )
    return EVOID} 
  
// special case when v is a g_func that can produce an error (s is assumed to be EID)
/* {1} OPT.The go function for: error_wrap(self:any,s:class,v:string) [] */
func F_Generate_error_wrap_any (self *ClaireAny ,s *ClaireClass ,v *ClaireString ) EID { 
    var Result EID 
    if (s.Id() != Optimize.C_EID.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var g0566UU *ClaireList  
        /* noccur = 1 */
        var g0566UU_try05674 EID 
        /* Construct:4 */{ 
          var v_bag_arg *ClaireAny  
          g0566UU_try05674= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
          ToList(OBJ(g0566UU_try05674)).AddFast(self)
          var v_bag_arg_try05685 EID 
          v_bag_arg_try05685 = Optimize.F_Compile_g_throw_any(self)
          /* ERROR PROTECTION INSERTED (v_bag_arg-g0566UU_try05674) */
          if ErrorIn(v_bag_arg_try05685) {g0566UU_try05674 = v_bag_arg_try05685
          } else {
          v_bag_arg = ANY(v_bag_arg_try05685)
          ToList(OBJ(g0566UU_try05674)).AddFast(v_bag_arg)}
          /* Construct-4 */} 
        /* ERROR PROTECTION INSERTED (g0566UU-Result) */
        if ErrorIn(g0566UU_try05674) {Result = g0566UU_try05674
        } else {
        g0566UU = ToList(OBJ(g0566UU_try05674))
        Result = Core.F_tformat_string(MakeString("---- g_throw(~S) = ~S\n"),0,g0566UU)
        }
        /* Let-3 */} 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = ToException(Core.C_general_error.Make(MakeString("design bug for error_wrap with ~S and s = ~S").Id(),MakeConstantList(self,s.Id()).Id())).Close()
      }
      } else {
      Result = EID{CFALSE.Id(),0}
      /* If-2 */} 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    F_princ_string(v)
    PRINC(" = ")
    Result = Core.F_CALL(C_Generate_g_expression,ARGS(self.ToEID(),EID{Optimize.C_EID.Id(),0}))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    F_Generate_breakline_void()
    PRINC("")
    Result = EVOID
    }
    }
    return Result} 
  
// The EID go function for: error_wrap @ any (throw: true) 
func E_Generate_error_wrap_any (self EID,s EID,v EID) EID { 
    return /*(sm for error_wrap @ any= EID)*/ F_Generate_error_wrap_any(ANY(self),ToClass(OBJ(s)),ToString(OBJ(v)) )} 
  
// this is a special case when the statement result is not needed (e = void) so we should
// not reuse v as a temporary variable (which we considered) 
// this is called inside a For/While, so loop is a tuple
/* {1} OPT.The go function for: g_try_void(self:any,vglobal:string,loop:any) [] */
func F_Generate_g_try_void_any (self *ClaireAny ,vglobal *ClaireString ,loop *ClaireAny ) EID { 
    var Result EID 
    var g0572I *ClaireBoolean  
    if (self.Isa.IsIn(Language.C_Assign) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0569 *Language.Assign   = Language.To_Assign(self)
        /* noccur = 1 */
        g0572I = Equal(ANY(Core.F_CALL(C_range,ARGS(g0569.ClaireVar.ToEID()))),Optimize.C_EID.Id())
        /* Let-3 */} 
      } else {
      g0572I = CFALSE
      /* If-2 */} 
    if (g0572I == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var _Zvar *ClaireAny   = ANY(Core.F_CALL(Language.C_var,ARGS(self.ToEID())))
        /* noccur = 1 */
        /* Let:4 */{ 
          var v1 *ClaireAny  
          /* noccur = 1 */
          var v1_try05735 EID 
          v1_try05735 = Core.F_CALL(C_Generate_c_string,ARGS(EID{Optimize.C_PRODUCER.Value,0},_Zvar.ToEID()))
          /* ERROR PROTECTION INSERTED (v1-Result) */
          if ErrorIn(v1_try05735) {Result = v1_try05735
          } else {
          v1 = ANY(v1_try05735)
          Result = Core.F_CALL(C_Generate_g_try,ARGS(Core.F_CALL(C_arg,ARGS(self.ToEID())),
            v1.ToEID(),
            EID{Optimize.C_EID.Id(),0},
            EID{(vglobal).Id(),0},
            loop.ToEID()))
          }
          /* Let-4 */} 
        /* Let-3 */} 
      } else {
      /* Let:3 */{ 
        var v2 *ClaireString   = F_Generate_check_var_string(MakeString("void_try"))
        /* noccur = 4 */
        F_Generate_var_declaration_string(v2,Optimize.C_EID,1)
        PRINC("_ = ")
        F_princ_string(v2)
        F_Generate_breakline_void()
        PRINC("")
        if (self.Isa.IsIn(Language.C_Do) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var g0570 *Language.Do   = Language.To_Do(self)
            /* noccur = 1 */
            PRINC("{ ")
            F_Generate_breakline_void()
            Result = F_Generate_do_statement_Do(g0570,
              Optimize.C_EID,
              v2,
              CTRUE,
              loop,
              CFALSE)
            /* Let-5 */} 
          } else {
          Result = Core.F_CALL(C_Generate_g_try,ARGS(self.ToEID(),
            EID{(v2).Id(),0},
            EID{Optimize.C_EID.Id(),0},
            EID{(vglobal).Id(),0},
            loop.ToEID()))
          /* If-4 */} 
        /* Let-3 */} 
      /* If-2 */} 
    return Result} 
  
// The EID go function for: g_try_void @ any (throw: true) 
func E_Generate_g_try_void_any (self EID,vglobal EID,loop EID) EID { 
    return /*(sm for g_try_void @ any= EID)*/ F_Generate_g_try_void_any(ANY(self),ToString(OBJ(vglobal)),ANY(loop) )} 
  
// eid_require means that the internal call should better be evaluated in EID mode
// this is really what we need for mClaire/push!(eval(x)) and funcall(f,...) ... but has been be extended to methods
// that do a better job (no allocation) in EID mode
/* {1} OPT.The go function for: eid_require?(x:any) [] */
func F_Generate_eid_require_ask_any (x *ClaireAny ) *ClaireBoolean  { 
    // procedure body with s =  
var Result *ClaireBoolean  
    if (x.Isa.IsIn(Language.C_Call) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0574 *Language.Call   = Language.To_Call(x)
        /* noccur = 1 */
        Result = Equal(g0574.Selector.Id(),Core.C_mClaire_push_I.Id())
        /* Let-3 */} 
      /* If!2 */}  else if (x.Isa.IsIn(Language.C_Call_method) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0575 *Language.CallMethod   = Language.To_CallMethod(x)
        /* noccur = 5 */
        Result = MakeBoolean((g0575.Arg.Selector.Id() == C_funcall.Id()) || (g0575.Arg.Id() == C_Generate__starwrite_value_star.Value) || (g0575.Arg.Id() == C_Generate__starread_property_star.Value) || (g0575.Arg.Selector.Id() == C_write_fast.Id()) || (g0575.Arg.Selector.Id() == Core.C_nth_write.Id()))
        /* Let-3 */} 
      } else {
      Result = CFALSE
      /* If-2 */} 
    return Result} 
  
// The EID go function for: eid_require? @ any (throw: false) 
func E_Generate_eid_require_ask_any (x EID) EID { 
    return EID{/*(sm for eid_require? @ any= boolean)*/ F_Generate_eid_require_ask_any(ANY(x) ).Id(),0}} 
  
// eid_provide? says that the call will produce first an EID
/* {1} OPT.The go function for: eid_provide?(x:any) [] */
func F_Generate_eid_provide_ask_any (x *ClaireAny ) *ClaireBoolean  { 
    // procedure body with s =  
var Result *ClaireBoolean  
    if (x.Isa.IsIn(Language.C_Call) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0577 *Language.Call   = Language.To_Call(x)
        /* noccur = 1 */
        Result = Equal(g0577.Selector.Id(),Core.C_mClaire_get_stack.Id())
        /* Let-3 */} 
      /* If!2 */}  else if (x.Isa.IsIn(Language.C_Call_method) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0578 *Language.CallMethod   = Language.To_CallMethod(x)
        /* noccur = 2 */
        Result = MakeBoolean((g0578.Arg.Selector.Id() == Core.C_eval.Id()) || (g0578.Arg.Id() == Optimize.C_Compile_m_unsafe.Value))
        /* Let-3 */} 
      /* If!2 */}  else if (x.Isa.IsIn(C_Variable) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0579 *ClaireVariable   = To_Variable(x)
        /* noccur = 1 */
        Result = g0579.Range.Included(ToType(C_integer.Id()))
        /* Let-3 */} 
      /* If!2 */}  else if (C_integer.Id() == x.Isa.Id()) /* If:2 */{ 
      Result = CTRUE
      } else {
      Result = CFALSE
      /* If-2 */} 
    return Result} 
  
// The EID go function for: eid_provide? @ any (throw: false) 
func E_Generate_eid_provide_ask_any (x EID) EID { 
    return EID{/*(sm for eid_provide? @ any= boolean)*/ F_Generate_eid_provide_ask_any(ANY(x) ).Id(),0}} 
  
// eid_unfold could use a more general "EID compling mode" (with a list of EID variables passed as context)
// this is a quickfix => we build the EID Let on our own (code borrowed from g_stat@Let)
/* {1} OPT.The go function for: unfold_eid(ldef:list,self:any,s:class,v:any,err:boolean,loop:any) [] */
func F_Generate_unfold_eid_list (ldef *ClaireList ,self *ClaireAny ,s *ClaireClass ,v *ClaireAny ,err *ClaireBoolean ,loop *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var n int  = ldef.Length()
      /* noccur = 1 */
      /* Let:3 */{ 
        var lvar *ClaireList   = ToType(C_Variable.Id()).EmptyList()
        /* noccur = 3 */
        /* Let:4 */{ 
          var count_try int  = 0
          /* noccur = 3 */
          F_Generate_new_block_string(MakeString("LetEID"))
          /* Let:5 */{ 
            var i int  = 1
            /* noccur = 4 */
            /* Let:6 */{ 
              var g0582 int  = n
              /* noccur = 1 */
              Result= EID{CFALSE.Id(),0}
              for (i <= g0582) /* while:7 */{ 
                var void_try8 EID 
                _ = void_try8
                { 
                /* Let:8 */{ 
                  var _Zl *Language.Let   = Language.To_Let(ldef.At(i-1))
                  /* noccur = 3 */
                  /* Let:9 */{ 
                    var v2 *ClaireString  
                    /* noccur = 3 */
                    var v2_try058310 EID 
                    v2_try058310 = F_Generate_c_string_go_producer1(ToGenerateGoProducer(Optimize.C_PRODUCER.Value),_Zl.ClaireVar)
                    /* ERROR PROTECTION INSERTED (v2-void_try8) */
                    if ErrorIn(v2_try058310) {void_try8 = v2_try058310
                    } else {
                    v2 = ToString(OBJ(v2_try058310))
                    /* Let:10 */{ 
                      var x *ClaireAny   = _Zl.Value
                      /* noccur = 3 */
                      /* Let:11 */{ 
                        var try_ask *ClaireBoolean  
                        /* noccur = 1 */
                        var try_ask_try058412 EID 
                        try_ask_try058412 = Optimize.F_Compile_g_throw_any(x)
                        /* ERROR PROTECTION INSERTED (try_ask-void_try8) */
                        if ErrorIn(try_ask_try058412) {void_try8 = try_ask_try058412
                        } else {
                        try_ask = ToBoolean(OBJ(try_ask_try058412))
                        F_Generate_var_declaration_string(v2,Optimize.C_EID,0)
                        lvar = lvar.AddFast(_Zl.ClaireVar.Id())
                        F_Generate_breakline_void()
                        if (try_ask == CTRUE) /* If:12 */{ 
                          count_try = (count_try+1)
                          void_try8 = Core.F_CALL(C_Generate_g_try,ARGS(x.ToEID(),
                            EID{(v2).Id(),0},
                            EID{Optimize.C_EID.Id(),0},
                            v.ToEID(),
                            EID{CFALSE.Id(),0}))
                          /* ERROR PROTECTION INSERTED (void_try8-void_try8) */
                          if ErrorIn(void_try8) {Result = void_try8
                          break
                          } else {
                          }
                          } else {
                          void_try8 = F_Generate_statement_any(x,Optimize.C_EID,v2,loop)
                          /* If-12 */} 
                        /* ERROR PROTECTION INSERTED (void_try8-void_try8) */
                        if ErrorIn(void_try8) {Result = void_try8
                        break
                        } else {
                        }
                        }
                        /* Let-11 */} 
                      /* Let-10 */} 
                    }
                    /* Let-9 */} 
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (void_try8-void_try8) */
                if ErrorIn(void_try8) {Result = void_try8
                break
                } else {
                i = (i+1)
                }
                /* while-7 */} 
              }
              /* Let-6 */} 
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          if (s.Id() != C_void.Id()) /* If:5 */{ 
            Result = Core.F_CALL(C_princ,ARGS(v.ToEID()))
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC(" = ")
            Result = EVOID
            }
            } else {
            Result = EID{CFALSE.Id(),0}
            /* If-5 */} 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          Result = F_Generate_eid_expression_any(self,Optimize.C_EID,lvar)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          F_Generate_close_try_integer(count_try)
          F_Generate_close_block_string(MakeString("LetEID"))
          Result = EVOID
          }}}
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: unfold_eid @ list (throw: true) 
func E_Generate_unfold_eid_list (ldef EID,self EID,s EID,v EID,err EID,loop EID) EID { 
    return /*(sm for unfold_eid @ list= EID)*/ F_Generate_unfold_eid_list(ToList(OBJ(ldef)),
      ANY(self),
      ToClass(OBJ(s)),
      ANY(v),
      ToBoolean(OBJ(err)),
      ANY(loop) )} 
  
// eid_expression compiles a call or call_method with one EID variable
// it performs all the compiler optimization (see the eid_fold? pattern in gostat.cl)
/* {1} OPT.The go function for: eid_expression(x:any,s:class,lvar:list<Variable>) [] */
func F_Generate_eid_expression_any (x *ClaireAny ,s *ClaireClass ,lvar *ClaireList ) EID { 
    var Result EID 
    if (x.Isa.IsIn(C_Variable) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0590 *ClaireVariable   = To_Variable(x)
        /* noccur = 3 */
        if ((lvar.Memq(g0590.Id()) == CTRUE) && 
            (s.Id() == Optimize.C_EID.Id())) /* If:4 */{ 
          /* Let:5 */{ 
            var g0594UU *ClaireString  
            /* noccur = 1 */
            var g0594UU_try05956 EID 
            g0594UU_try05956 = F_Generate_c_string_go_producer1(ToGenerateGoProducer(Optimize.C_PRODUCER.Value),g0590)
            /* ERROR PROTECTION INSERTED (g0594UU-Result) */
            if ErrorIn(g0594UU_try05956) {Result = g0594UU_try05956
            } else {
            g0594UU = ToString(OBJ(g0594UU_try05956))
            F_princ_string(g0594UU)
            Result = EVOID
            }
            /* Let-5 */} 
          } else {
          Result = F_Generate_g_expression_Variable(g0590,s)
          /* If-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (x.Isa.IsIn(Language.C_Call) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0591 *Language.Call   = Language.To_Call(x)
        /* noccur = 1 */
        PRINC("ClEnv.Push(")
        Result = F_Generate_eid_expression_any(g0591.Args.At(1-1),Optimize.C_EID,lvar)
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC(")")
        Result = EVOID
        }
        /* Let-3 */} 
      /* If!2 */}  else if (x.Isa.IsIn(Language.C_Call_method) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0592 *Language.CallMethod   = Language.To_CallMethod(x)
        /* noccur = 21 */
        if (g0592.Arg.Selector.Id() == C_funcall.Id()) /* If:4 */{ 
          PRINC("FASTCALL")
          F_princ_integer((g0592.Args.Length()-1))
          PRINC("(")
          Result = F_Generate_eid_expression_any(g0592.Args.At(1-1),C_method,lvar)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(",")
          Result = F_Generate_eid_expression_any(g0592.Args.At(2-1),Optimize.C_EID,lvar)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          if (g0592.Args.Length() >= 3) /* If:5 */{ 
            PRINC(",")
            Result = F_Generate_eid_expression_any(g0592.Args.At(3-1),Optimize.C_EID,lvar)
            } else {
            Result = EID{CFALSE.Id(),0}
            /* If-5 */} 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          if (g0592.Args.Length() == 4) /* If:5 */{ 
            PRINC(",")
            Result = F_Generate_eid_expression_any(g0592.Args.At(4-1),Optimize.C_EID,lvar)
            } else {
            Result = EID{CFALSE.Id(),0}
            /* If-5 */} 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(")")
          Result = EVOID
          }}}}
          /* If!4 */}  else if (g0592.Arg.Id() == C_Generate__starread_property_star.Value) /* If:4 */{ 
          Result = Core.F_CALL(C_Generate_g_expression,ARGS(g0592.Args.At(1-1).ToEID(),EID{C_property.Id(),0}))
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(".ReadEID(")
          Result = F_Generate_eid_expression_any(g0592.Args.At(2-1),Optimize.C_EID,lvar)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(")")
          Result = EVOID
          }}
          /* If!4 */}  else if (g0592.Arg.Selector.Id() == C_write_fast.Id()) /* If:4 */{ 
          Result = Core.F_CALL(C_Generate_g_expression,ARGS(g0592.Args.At(1-1).ToEID(),EID{C_property.Id(),0}))
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(".WriteEID(")
          Result = Core.F_CALL(C_Generate_g_expression,ARGS(g0592.Args.At(2-1).ToEID(),EID{C_object.Id(),0}))
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(",")
          Result = F_Generate_eid_expression_any(g0592.Args.At(3-1),Optimize.C_EID,lvar)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(")")
          Result = EVOID
          }}}
          /* If!4 */}  else if (g0592.Arg.Selector.Id() == Core.C_nth_write.Id()) /* If:4 */{ 
          Result = Core.F_CALL(C_Generate_g_expression,ARGS(g0592.Args.At(1-1).ToEID(),EID{C_list.Id(),0}))
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(".WriteEID(")
          Result = Core.F_CALL(C_Generate_g_expression,ARGS(g0592.Args.At(2-1).ToEID(),EID{C_integer.Id(),0}))
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(",")
          Result = F_Generate_eid_expression_any(g0592.Args.At(3-1),Optimize.C_EID,lvar)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(")")
          Result = EVOID
          }}}
          } else {
          Result = Core.F_CALL(C_Generate_g_expression,ARGS(g0592.Args.At(1-1).ToEID(),EID{C_Variable.Id(),0}))
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(".WriteEID(")
          Result = F_Generate_eid_expression_any(g0592.Args.At(2-1),Optimize.C_EID,lvar)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(")")
          Result = EVOID
          }}
          /* If-4 */} 
        /* Let-3 */} 
      } else {
      Result = Core.F_CALL(C_Generate_g_expression,ARGS(x.ToEID(),EID{s.Id(),0}))
      /* If-2 */} 
    return Result} 
  
// The EID go function for: eid_expression @ any (throw: true) 
func E_Generate_eid_expression_any (x EID,s EID,lvar EID) EID { 
    return /*(sm for eid_expression @ any= EID)*/ F_Generate_eid_expression_any(ANY(x),ToClass(OBJ(s)),ToList(OBJ(lvar)) )} 
  
//**********************************************************************
//*          Part 3: Basic control structures                          *
//**********************************************************************
// The re-entry definition (called within g_statement, not directly)
// if functional, the best compiling is into an expression
// s is the expected go type (as a class) + void + EID
// v is nil or a string (name of the variable)
// note that only 3 additional parameters are used since err is recomputed
/* {1} OPT.The go function for: statement(self:any,s:class,v:string,loop:any) [] */
func F_Generate_statement_any (self *ClaireAny ,s *ClaireClass ,v *ClaireString ,loop *ClaireAny ) EID { 
    var Result EID 
    var g0596I *ClaireBoolean  
    var g0596I_try05972 EID 
    g0596I_try05972 = F_Generate_g_clean_any(self)
    /* ERROR PROTECTION INSERTED (g0596I-Result) */
    if ErrorIn(g0596I_try05972) {Result = g0596I_try05972
    } else {
    g0596I = ToBoolean(OBJ(g0596I_try05972))
    if (g0596I == CTRUE) /* If:2 */{ 
      var g0598I *ClaireBoolean  
      var g0598I_try05993 EID 
      /* and:3 */{ 
        var v_and3 *ClaireBoolean  
        
        v_and3 = Core.F__I_equal_any(s.Id(),C_void.Id())
        if (v_and3 == CFALSE) {g0598I_try05993 = EID{CFALSE.Id(),0}
        } else /* arg:4 */{ 
          var v_and3_try06005 EID 
          /* Let:5 */{ 
            var g0601UU *ClaireClass  
            /* noccur = 1 */
            var g0601UU_try06026 EID 
            g0601UU_try06026 = Language.F_static_type_any(self)
            /* ERROR PROTECTION INSERTED (g0601UU-v_and3_try06005) */
            if ErrorIn(g0601UU_try06026) {v_and3_try06005 = g0601UU_try06026
            } else {
            g0601UU = ToClass(OBJ(g0601UU_try06026))
            v_and3_try06005 = EID{Core.F__I_equal_any(g0601UU.Id(),C_void.Id()).Id(),0}
            }
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (v_and3-g0598I_try05993) */
          if ErrorIn(v_and3_try06005) {g0598I_try05993 = v_and3_try06005
          } else {
          v_and3 = ToBoolean(OBJ(v_and3_try06005))
          if (v_and3 == CFALSE) {g0598I_try05993 = EID{CFALSE.Id(),0}
          } else /* arg:5 */{ 
            g0598I_try05993 = EID{CTRUE.Id(),0}/* arg-5 */} 
          /* arg-4 */} 
        }
        /* and-3 */} 
      /* ERROR PROTECTION INSERTED (g0598I-Result) */
      if ErrorIn(g0598I_try05993) {Result = g0598I_try05993
      } else {
      g0598I = ToBoolean(OBJ(g0598I_try05993))
      if (g0598I == CTRUE) /* If:3 */{ 
        F_c_princ_string(v)
        PRINC(" = ")
        Result = Core.F_CALL(C_Generate_g_expression,ARGS(self.ToEID(),EID{s.Id(),0}))
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        F_Generate_breakline_void()
        PRINC("")
        Result = EVOID
        }
        /* If!3 */}  else if (self.Isa.IsIn(Language.C_If) == CTRUE) /* If:3 */{ 
        Result = Core.F_CALL(C_Generate_g_statement,ARGS(self.ToEID(),
          EID{s.Id(),0},
          EID{(v).Id(),0},
          EID{CFALSE.Id(),0},
          loop.ToEID()))
        /* If!3 */}  else if (self.Isa.IsIn(Reader.C_delimiter) == CTRUE) /* If:3 */{ 
        Result = ToException(Core.C_general_error.Make(MakeString("[201] Loose delimiter in program: ~S").Id(),MakeConstantList(self).Id())).Close()
        } else {
        Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).StatExp(self,C_void)
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        if (s.Id() == Optimize.C_EID.Id()) /* If:4 */{ 
          F_c_princ_string(v)
          PRINC(" = EVOID")
          F_Generate_breakline_void()
          PRINC("")
          Result = EVOID
          } else {
          Result = EID{CFALSE.Id(),0}
          /* If-4 */} 
        }
        /* If-3 */} 
      }
      } else {
      /* Let:3 */{ 
        var g0603UU *ClaireBoolean  
        /* noccur = 1 */
        var g0603UU_try06044 EID 
        g0603UU_try06044 = Optimize.F_Compile_g_throw_any(self)
        /* ERROR PROTECTION INSERTED (g0603UU-Result) */
        if ErrorIn(g0603UU_try06044) {Result = g0603UU_try06044
        } else {
        g0603UU = ToBoolean(OBJ(g0603UU_try06044))
        Result = Core.F_CALL(C_Generate_g_statement,ARGS(self.ToEID(),
          EID{s.Id(),0},
          EID{(v).Id(),0},
          EID{g0603UU.Id(),0},
          loop.ToEID()))
        }
        /* Let-3 */} 
      /* If-2 */} 
    }
    return Result} 
  
// The EID go function for: statement @ any (throw: true) 
func E_Generate_statement_any (self EID,s EID,v EID,loop EID) EID { 
    return /*(sm for statement @ any= EID)*/ F_Generate_statement_any(ANY(self),
      ToClass(OBJ(s)),
      ToString(OBJ(v)),
      ANY(loop) )} 
  
// make a statement from an expression (in C++ we need a ; - with go a breakline is enough)
// we do not want to place the result in a variable (see upper)
// will get simpler once we have a stable compiler without to_C and to_CL
/* {1} OPT.The go function for: stat_exp(c:go_producer,self:any,s:class) [] */
func (c *GenerateGoProducer ) StatExp (self *ClaireAny ,s *ClaireClass ) EID { 
    var Result EID 
    var g0605I *ClaireBoolean  
    var g0605I_try06062 EID 
    g0605I_try06062 = Optimize.F_Compile_designated_ask_any(self)
    /* ERROR PROTECTION INSERTED (g0605I-Result) */
    if ErrorIn(g0605I_try06062) {Result = g0605I_try06062
    } else {
    g0605I = ToBoolean(OBJ(g0605I_try06062))
    if (g0605I == CTRUE) /* If:2 */{ 
      Result = F_Generate_breakline_void().ToEID()
      } else {
      Result = Core.F_CALL(C_Generate_g_expression,ARGS(self.ToEID(),EID{s.Id(),0}))
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      F_Generate_breakline_void()
      PRINC("")
      Result = EVOID
      }
      /* If-2 */} 
    }
    return Result} 
  
// The EID go function for: stat_exp @ go_producer (throw: true) 
func E_Generate_stat_exp_go_producer (c EID,self EID,s EID) EID { 
    return /*(sm for stat_exp @ go_producer= EID)*/ ToGenerateGoProducer(OBJ(c)).StatExp(ANY(self),ToClass(OBJ(s)) )} 
  
// a DO is a simple go block if there are not errors, a chain otherwise
// the chain means multiple nestings when an error occurs since the rest of the DO must not be
// this is why the close_try(count) are called at the end, to close the embedded ifs (ErrorIn(e))
// we use a specific method code_statement with an additional parameter %need which is true by default
/* {1} OPT.The go function for: g_statement(self:Do,e:class,v:string,err:boolean,loop:any) [] */
func F_Generate_g_statement_Do (self *Language.Do ,e *ClaireClass ,v *ClaireString ,err *ClaireBoolean ,loop *ClaireAny ) EID { 
    var Result EID 
    Result = F_Generate_do_statement_Do(self,
      e,
      v,
      err,
      loop,
      CTRUE)
    return Result} 
  
// The EID go function for: g_statement @ Do (throw: true) 
func E_Generate_g_statement_Do (self EID,e EID,v EID,err EID,loop EID) EID { 
    return /*(sm for g_statement @ Do= EID)*/ F_Generate_g_statement_Do(Language.To_Do(OBJ(self)),
      ToClass(OBJ(e)),
      ToString(OBJ(v)),
      ToBoolean(OBJ(err)),
      ANY(loop) )} 
  
// %need is true = the last argument is needed in v
/* {1} OPT.The go function for: do_statement(self:Do,e:class,v:string,err:boolean,loop:any,%need:boolean) [] */
func F_Generate_do_statement_Do (self *Language.Do ,e *ClaireClass ,v *ClaireString ,err *ClaireBoolean ,loop *ClaireAny ,_Zneed *ClaireBoolean ) EID { 
    var Result EID 
    if (self.Args.Length() == 1) /* If:2 */{ 
      Result = F_Generate_statement_any(self.Args.At(1-1),e,v,loop)
      /* If!2 */}  else if (err != CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var l *ClaireList   = self.Args
        /* noccur = 2 */
        /* Let:4 */{ 
          var m int  = l.Length()
          /* noccur = 2 */
          /* Let:5 */{ 
            var n int  = 1
            /* noccur = 5 */
            /* Let:6 */{ 
              var g0607 int  = m
              /* noccur = 1 */
              Result= EID{CFALSE.Id(),0}
              for (n <= g0607) /* while:7 */{ 
                var void_try8 EID 
                _ = void_try8
                { 
                /* Let:8 */{ 
                  var g0609UU *ClaireClass  
                  /* noccur = 1 */
                  if (n == m) /* If:9 */{ 
                    g0609UU = e
                    } else {
                    g0609UU = C_void
                    /* If-9 */} 
                  void_try8 = F_Generate_statement_any(l.At(n-1),g0609UU,v,loop)
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (void_try8-void_try8) */
                if ErrorIn(void_try8) {Result = void_try8
                break
                } else {
                n = (n+1)
                }
                /* while-7 */} 
              }
              /* Let-6 */} 
            /* Let-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      } else {
      /* Let:3 */{ 
        var l *ClaireList   = self.Args
        /* noccur = 2 */
        /* Let:4 */{ 
          var m int  = l.Length()
          /* noccur = 3 */
          /* Let:5 */{ 
            var count_if int  = 0
            /* noccur = 3 */
            /* Let:6 */{ 
              var n int  = 1
              /* noccur = 6 */
              /* Let:7 */{ 
                var g0608 int  = m
                /* noccur = 1 */
                Result= EID{CFALSE.Id(),0}
                for (n <= g0608) /* while:8 */{ 
                  var void_try9 EID 
                  _ = void_try9
                  { 
                  /* Let:9 */{ 
                    var x *ClaireAny   = l.At(n-1)
                    /* noccur = 4 */
                    var g0610I *ClaireBoolean  
                    var g0610I_try061110 EID 
                    g0610I_try061110 = Optimize.F_Compile_g_throw_any(x)
                    /* ERROR PROTECTION INSERTED (g0610I-void_try9) */
                    if ErrorIn(g0610I_try061110) {void_try9 = g0610I_try061110
                    } else {
                    g0610I = ToBoolean(OBJ(g0610I_try061110))
                    if (g0610I == CTRUE) /* If:10 */{ 
                      if ((n < m) || 
                          (C_tuple.Id() == loop.Isa.Id())) /* If:11 */{ 
                        count_if = (count_if+1)
                        void_try9 = Core.F_CALL(C_Generate_g_try,ARGS(x.ToEID(),
                          EID{(v).Id(),0},
                          EID{Optimize.C_EID.Id(),0},
                          EID{(v).Id(),0},
                          loop.ToEID()))
                        /* ERROR PROTECTION INSERTED (void_try9-void_try9) */
                        if ErrorIn(void_try9) {Result = void_try9
                        break
                        } else {
                        }
                        } else {
                        void_try9 = Core.F_CALL(C_Generate_g_statement,ARGS(x.ToEID(),
                          EID{Optimize.C_EID.Id(),0},
                          EID{(v).Id(),0},
                          EID{CTRUE.Id(),0},
                          loop.ToEID()))
                        /* If-11 */} 
                      } else {
                      /* Let:11 */{ 
                        var g0612UU *ClaireClass  
                        /* noccur = 1 */
                        if ((n == m) && 
                            (_Zneed == CTRUE)) /* If:12 */{ 
                          g0612UU = e
                          } else {
                          g0612UU = C_void
                          /* If-12 */} 
                        void_try9 = F_Generate_statement_any(x,g0612UU,v,loop)
                        /* Let-11 */} 
                      /* If-10 */} 
                    }
                    /* Let-9 */} 
                  /* ERROR PROTECTION INSERTED (void_try9-void_try9) */
                  if ErrorIn(void_try9) {Result = void_try9
                  break
                  } else {
                  n = (n+1)
                  }
                  /* while-8 */} 
                }
                /* Let-7 */} 
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            F_Generate_close_try_integer(count_if)
            Result = EVOID
            }
            /* Let-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* If-2 */} 
    return Result} 
  
// The EID go function for: do_statement @ Do (throw: true) 
func E_Generate_do_statement_Do (self EID,e EID,v EID,err EID,loop EID,_Zneed EID) EID { 
    return /*(sm for do_statement @ Do= EID)*/ F_Generate_do_statement_Do(Language.To_Do(OBJ(self)),
      ToClass(OBJ(e)),
      ToString(OBJ(v)),
      ToBoolean(OBJ(err)),
      ANY(loop),
      ToBoolean(OBJ(_Zneed)) )} 
  
// a Let is a local variable declaration 
// in CLAIRE 4, a block is anything that fits between {} hence inner/outer is not necessary
// AXIOM if err is true, we require that e = any
/* {1} OPT.The go function for: g_statement(self:Let,e:class,v:string,err:boolean,loop:any) [] */
func F_Generate_g_statement_Let (self *Language.Let ,e *ClaireClass ,v *ClaireString ,err *ClaireBoolean ,loop *ClaireAny ) EID { 
    var Result EID 
    var g0613I *ClaireBoolean  
    var g0613I_try06142 EID 
    g0613I_try06142 = F_Generate_let_eid_ask_Let(self)
    /* ERROR PROTECTION INSERTED (g0613I-Result) */
    if ErrorIn(g0613I_try06142) {Result = g0613I_try06142
    } else {
    g0613I = ToBoolean(OBJ(g0613I_try06142))
    if (g0613I == CTRUE) /* If:2 */{ 
      Result = F_Generate_g_eid_stat_Let(self,
        e,
        v,
        err,
        loop)
      } else {
      /* Let:3 */{ 
        var ns *ClaireString  
        /* noccur = 2 */
        var ns_try06164 EID 
        ns_try06164 = F_Generate_c_string_go_producer2(ToGenerateGoProducer(Optimize.C_PRODUCER.Value),self.ClaireVar.Pname)
        /* ERROR PROTECTION INSERTED (ns-Result) */
        if ErrorIn(ns_try06164) {Result = ns_try06164
        } else {
        ns = ToString(OBJ(ns_try06164))
        if ((ns.At(1) == 'C') && 
            (ns.At(2) == '%')) /* If:4 */{ 
          self.ClaireVar.Pname = Core.F_gensym_void()
          /* If-4 */} 
        /* Let:4 */{ 
          var v2 *ClaireString  
          /* noccur = 5 */
          var v2_try06175 EID 
          v2_try06175 = F_Generate_c_string_go_producer1(ToGenerateGoProducer(Optimize.C_PRODUCER.Value),self.ClaireVar)
          /* ERROR PROTECTION INSERTED (v2-Result) */
          if ErrorIn(v2_try06175) {Result = v2_try06175
          } else {
          v2 = ToString(OBJ(v2_try06175))
          /* Let:5 */{ 
            var x *ClaireAny   = self.Value
            /* noccur = 5 */
            /* Let:6 */{ 
              var f *ClaireBoolean  
              /* noccur = 2 */
              var f_try06187 EID 
              f_try06187 = F_Generate_g_clean_any(x)
              /* ERROR PROTECTION INSERTED (f-Result) */
              if ErrorIn(f_try06187) {Result = f_try06187
              } else {
              f = ToBoolean(OBJ(f_try06187))
              /* Let:7 */{ 
                var try_ask *ClaireBoolean  
                /* noccur = 2 */
                var try_ask_try06198 EID 
                try_ask_try06198 = Optimize.F_Compile_g_throw_any(x)
                /* ERROR PROTECTION INSERTED (try_ask-Result) */
                if ErrorIn(try_ask_try06198) {Result = try_ask_try06198
                } else {
                try_ask = ToBoolean(OBJ(try_ask_try06198))
                /* Let:8 */{ 
                  var ev *ClaireClass   = self.ClaireVar.Range.Class_I()
                  /* noccur = 4 */
                  F_Generate_new_block_string(MakeString("Let"))
                  F_Generate_var_declaration_string(v2,ev,0)
                  if (f == CTRUE) /* If:9 */{ 
                    PRINC(" = ")
                    Result = Core.F_CALL(C_Generate_g_expression,ARGS(x.ToEID(),EID{ev.Id(),0}))
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    PRINC("")
                    Result = EVOID
                    }
                    } else {
                    Result = EID{CFALSE.Id(),0}
                    /* If-9 */} 
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  F_Generate_breakline_void()
                  PRINC("/* noccur = ")
                  F_princ_integer(Language.F_Language_occurexact_any(self.Arg,self.ClaireVar))
                  PRINC(" */")
                  F_Generate_breakline_void()
                  if (Language.F_Language_occurexact_any(self.Arg,self.ClaireVar) == 0) /* If:9 */{ 
                    Core.F_tformat_string(MakeString(">>>>>>>>  variable ~S declared but unused  \n"),0,MakeConstantList((v2).Id()))
                    PRINC("_ = ")
                    F_princ_string(v2)
                    F_Generate_breakline_void()
                    PRINC("")
                    /* If-9 */} 
                  if (try_ask == CTRUE) /* If:9 */{ 
                    Result = Core.F_CALL(C_Generate_g_try,ARGS(x.ToEID(),
                      EID{(v2).Id(),0},
                      EID{ev.Id(),0},
                      EID{(v).Id(),0},
                      EID{CFALSE.Id(),0}))
                    /* If!9 */}  else if (f != CTRUE) /* If:9 */{ 
                    Result = F_Generate_statement_any(x,ev,v2,loop)
                    } else {
                    Result = EID{CFALSE.Id(),0}
                    /* If-9 */} 
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  Result = F_Generate_statement_any(self.Arg,e,v,loop)
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  if (try_ask == CTRUE) /* If:9 */{ 
                    F_Generate_close_try_integer(1)
                    /* If-9 */} 
                  F_Generate_close_block_string(MakeString("Let"))
                  Result = EVOID
                  }}}
                  /* Let-8 */} 
                }
                /* Let-7 */} 
              }
              /* Let-6 */} 
            /* Let-5 */} 
          }
          /* Let-4 */} 
        }
        /* Let-3 */} 
      /* If-2 */} 
    }
    return Result} 
  
// The EID go function for: g_statement @ Let (throw: true) 
func E_Generate_g_statement_Let (self EID,e EID,v EID,err EID,loop EID) EID { 
    return /*(sm for g_statement @ Let= EID)*/ F_Generate_g_statement_Let(Language.To_Let(OBJ(self)),
      ToClass(OBJ(e)),
      ToString(OBJ(v)),
      ToBoolean(OBJ(err)),
      ANY(loop) )} 
  
// then we must close the chain
// detect a pattern (that could be expanded) where the variable should be compiled as EID because
// the value may trigger an error and the body uses v once at the end (this could be : the body is EID-friendly)
// this current specific pattern is introduced to optimize eval_message
/* {1} OPT.The go function for: let_eid?(self:Let) [] */
func F_Generate_let_eid_ask_Let (self *Language.Let ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var v *ClaireVariable   = self.ClaireVar
      /* noccur = 2 */
      /* Let:3 */{ 
        var y *ClaireAny   = self.Arg
        /* noccur = 2 */
        /* or:4 */{ 
          var v_or4 *ClaireBoolean  
          
          var v_or4_try06215 EID 
          /* and:5 */{ 
            var v_and5 *ClaireBoolean  
            
            var v_and5_try06226 EID 
            /* or:6 */{ 
              var v_or6 *ClaireBoolean  
              
              var v_or6_try06237 EID 
              v_or6_try06237 = Optimize.F_Compile_g_throw_any(self.Value)
              /* ERROR PROTECTION INSERTED (v_or6-v_and5_try06226) */
              if ErrorIn(v_or6_try06237) {v_and5_try06226 = v_or6_try06237
              } else {
              v_or6 = ToBoolean(OBJ(v_or6_try06237))
              if (v_or6 == CTRUE) {v_and5_try06226 = EID{CTRUE.Id(),0}
              } else /* or:7 */{ 
                v_or6 = F_Generate_eid_provide_ask_any(self.Value)
                if (v_or6 == CTRUE) {v_and5_try06226 = EID{CTRUE.Id(),0}
                } else /* or:8 */{ 
                  v_and5_try06226 = EID{CFALSE.Id(),0}/* org-8 */} 
                /* org-7 */} 
              }
              /* or-6 */} 
            /* ERROR PROTECTION INSERTED (v_and5-v_or4_try06215) */
            if ErrorIn(v_and5_try06226) {v_or4_try06215 = v_and5_try06226
            } else {
            v_and5 = ToBoolean(OBJ(v_and5_try06226))
            if (v_and5 == CFALSE) {v_or4_try06215 = EID{CFALSE.Id(),0}
            } else /* arg:6 */{ 
              v_and5 = Equal(MakeInteger(Language.F_occurrence_any(self.Arg,self.ClaireVar)).Id(),MakeInteger(1).Id())
              if (v_and5 == CFALSE) {v_or4_try06215 = EID{CFALSE.Id(),0}
              } else /* arg:7 */{ 
                var v_and5_try06248 EID 
                if (y.Isa.IsIn(Language.C_Do) == CTRUE) /* If:8 */{ 
                  /* Let:9 */{ 
                    var g0620 *Language.Do   = Language.To_Do(y)
                    /* noccur = 2 */
                    /* and:10 */{ 
                      var v_and10 *ClaireBoolean  
                      
                      var v_and10_try062511 EID 
                      /* Let:11 */{ 
                        var g0626UU *ClaireBoolean  
                        /* noccur = 1 */
                        var g0626UU_try062712 EID 
                        g0626UU_try062712 = Optimize.F_Compile_g_throw_any(g0620.Id())
                        /* ERROR PROTECTION INSERTED (g0626UU-v_and10_try062511) */
                        if ErrorIn(g0626UU_try062712) {v_and10_try062511 = g0626UU_try062712
                        } else {
                        g0626UU = ToBoolean(OBJ(g0626UU_try062712))
                        v_and10_try062511 = EID{g0626UU.Not.Id(),0}
                        }
                        /* Let-11 */} 
                      /* ERROR PROTECTION INSERTED (v_and10-v_and5_try06248) */
                      if ErrorIn(v_and10_try062511) {v_and5_try06248 = v_and10_try062511
                      } else {
                      v_and10 = ToBoolean(OBJ(v_and10_try062511))
                      if (v_and10 == CFALSE) {v_and5_try06248 = EID{CFALSE.Id(),0}
                      } else /* arg:11 */{ 
                        var v_and10_try062812 EID 
                        /* Let:12 */{ 
                          var g0629UU *ClaireAny  
                          /* noccur = 1 */
                          var g0629UU_try063013 EID 
                          g0629UU_try063013 = Core.F_last_list(g0620.Args)
                          /* ERROR PROTECTION INSERTED (g0629UU-v_and10_try062812) */
                          if ErrorIn(g0629UU_try063013) {v_and10_try062812 = g0629UU_try063013
                          } else {
                          g0629UU = ANY(g0629UU_try063013)
                          v_and10_try062812 = EID{Equal(g0629UU,v.Id()).Id(),0}
                          }
                          /* Let-12 */} 
                        /* ERROR PROTECTION INSERTED (v_and10-v_and5_try06248) */
                        if ErrorIn(v_and10_try062812) {v_and5_try06248 = v_and10_try062812
                        } else {
                        v_and10 = ToBoolean(OBJ(v_and10_try062812))
                        if (v_and10 == CFALSE) {v_and5_try06248 = EID{CFALSE.Id(),0}
                        } else /* arg:12 */{ 
                          v_and5_try06248 = EID{CTRUE.Id(),0}/* arg-12 */} 
                        /* arg-11 */} 
                      }}
                      /* and-10 */} 
                    /* Let-9 */} 
                  } else {
                  v_and5_try06248 = EID{CFALSE.Id(),0}
                  /* If-8 */} 
                /* ERROR PROTECTION INSERTED (v_and5-v_or4_try06215) */
                if ErrorIn(v_and5_try06248) {v_or4_try06215 = v_and5_try06248
                } else {
                v_and5 = ToBoolean(OBJ(v_and5_try06248))
                if (v_and5 == CFALSE) {v_or4_try06215 = EID{CFALSE.Id(),0}
                } else /* arg:8 */{ 
                  v_or4_try06215 = EID{CTRUE.Id(),0}/* arg-8 */} 
                /* arg-7 */} 
              /* arg-6 */} 
            }}
            /* and-5 */} 
          /* ERROR PROTECTION INSERTED (v_or4-Result) */
          if ErrorIn(v_or4_try06215) {Result = v_or4_try06215
          } else {
          v_or4 = ToBoolean(OBJ(v_or4_try06215))
          if (v_or4 == CTRUE) {Result = EID{CTRUE.Id(),0}
          } else /* or:5 */{ 
            v_or4 = Equal(v.Range.Id(),Optimize.C_EID.Id())
            if (v_or4 == CTRUE) {Result = EID{CTRUE.Id(),0}
            } else /* or:6 */{ 
              Result = EID{CFALSE.Id(),0}/* org-6 */} 
            /* org-5 */} 
          }
          /* or-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: let_eid? @ Let (throw: true) 
func E_Generate_let_eid_ask_Let (self EID) EID { 
    return /*(sm for let_eid? @ Let= EID)*/ F_Generate_let_eid_ask_Let(Language.To_Let(OBJ(self)) )} 
  
// force EID compiling : back door :)
// the corresponding compiling (embeds the Do)
/* {1} OPT.The go function for: g_eid_stat(self:Let,e:class,v:string,err:boolean,loop:any) [] */
func F_Generate_g_eid_stat_Let (self *Language.Let ,e *ClaireClass ,v *ClaireString ,err *ClaireBoolean ,loop *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var v2 *ClaireString  
      /* noccur = 4 */
      var v2_try06323 EID 
      v2_try06323 = F_Generate_c_string_go_producer1(ToGenerateGoProducer(Optimize.C_PRODUCER.Value),self.ClaireVar)
      /* ERROR PROTECTION INSERTED (v2-Result) */
      if ErrorIn(v2_try06323) {Result = v2_try06323
      } else {
      v2 = ToString(OBJ(v2_try06323))
      /* Let:3 */{ 
        var x *ClaireAny   = self.Value
        /* noccur = 3 */
        /* Let:4 */{ 
          var try_ask *ClaireBoolean  
          /* noccur = 2 */
          var try_ask_try06335 EID 
          try_ask_try06335 = Optimize.F_Compile_g_throw_any(x)
          /* ERROR PROTECTION INSERTED (try_ask-Result) */
          if ErrorIn(try_ask_try06335) {Result = try_ask_try06335
          } else {
          try_ask = ToBoolean(OBJ(try_ask_try06335))
          F_Generate_new_block_string(MakeString("LetE"))
          F_Generate_var_declaration_string(v2,Optimize.C_EID,0)
          F_Generate_breakline_void()
          if (try_ask == CTRUE) /* If:5 */{ 
            Result = Core.F_CALL(C_Generate_g_try,ARGS(x.ToEID(),
              EID{(v2).Id(),0},
              EID{Optimize.C_EID.Id(),0},
              EID{(v).Id(),0},
              EID{CFALSE.Id(),0}))
            } else {
            F_princ_string(v2)
            PRINC(" = ")
            Result = Core.F_CALL(C_Generate_g_expression,ARGS(x.ToEID(),EID{Optimize.C_EID.Id(),0}))
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            F_Generate_breakline_void()
            PRINC("")
            Result = EVOID
            }
            /* If-5 */} 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          /* Let:5 */{ 
            var y *Language.Do   = Language.To_Do(self.Arg)
            /* noccur = 2 */
            /* Let:6 */{ 
              var n int  = y.Args.Length()
              /* noccur = 1 */
              /* Let:7 */{ 
                var i int  = 1
                /* noccur = 4 */
                /* Let:8 */{ 
                  var g0631 int  = (n-1)
                  /* noccur = 1 */
                  Result= EID{CFALSE.Id(),0}
                  for (i <= g0631) /* while:9 */{ 
                    var void_try10 EID 
                    _ = void_try10
                    { 
                    void_try10 = F_Generate_statement_any(y.Args.At(i-1),C_void,v,loop)
                    /* ERROR PROTECTION INSERTED (void_try10-void_try10) */
                    if ErrorIn(void_try10) {Result = void_try10
                    break
                    } else {
                    i = (i+1)
                    }
                    /* while-9 */} 
                  }
                  /* Let-8 */} 
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              F_princ_string(v)
              PRINC(" = ")
              F_princ_string(v2)
              PRINC("")
              Result = EVOID
              }
              /* Let-6 */} 
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          if (try_ask == CTRUE) /* If:5 */{ 
            F_Generate_close_try_integer(1)
            /* If-5 */} 
          F_Generate_close_block_string(MakeString("LetE"))
          Result = EVOID
          }}
          }
          /* Let-4 */} 
        /* Let-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: g_eid_stat @ Let (throw: true) 
func E_Generate_g_eid_stat_Let (self EID,e EID,v EID,err EID,loop EID) EID { 
    return /*(sm for g_eid_stat @ Let= EID)*/ F_Generate_g_eid_stat_Let(Language.To_Let(OBJ(self)),
      ToClass(OBJ(e)),
      ToString(OBJ(v)),
      ToBoolean(OBJ(err)),
      ANY(loop) )} 
  
// then we must close the chain
// makes a bag from a list of  statements.
// The value cannot be ignored: it is considered as an error (a do should have been used)
// there are two patterns depending if self.of is known : MakeEmptyX(t) or MakeEmptyX(any)
/* {1} OPT.The go function for: g_statement(self:Construct,s:class,v:string,err:boolean,loop:any) [] */
func F_Generate_g_statement_Construct (self *Language.Construct ,s *ClaireClass ,v *ClaireString ,err *ClaireBoolean ,loop *ClaireAny ) EID { 
    var Result EID 
    if (F_boolean_I_any((v).Id()).Id() != CTRUE.Id()) /* If:2 */{ 
      Result = ToException(Core.C_general_error.Make(MakeString("[202] A do should have been used for ~S").Id(),MakeConstantList(self.Id()).Id())).Close()
      } else {
      Result = EID{CFALSE.Id(),0}
      /* If-2 */} 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    /* Let:2 */{ 
      var v2 *ClaireString   = MakeString("v_bag_arg")
      /* noccur = 4 */
      /* Let:3 */{ 
        var kind *ClaireClass  
        /* noccur = 6 */
        var kind_try06384 EID 
        if (self.Isa.IsIn(Language.C_List) == CTRUE) /* If:4 */{ 
          kind_try06384 = EID{C_list.Id(),0}
          /* If!4 */}  else if (self.Isa.IsIn(Language.C_Set) == CTRUE) /* If:4 */{ 
          kind_try06384 = EID{C_set.Id(),0}
          /* If!4 */}  else if (self.Isa.IsIn(Language.C_Tuple) == CTRUE) /* If:4 */{ 
          kind_try06384 = EID{C_tuple.Id(),0}
          } else {
          kind_try06384 = ToException(Core.C_general_error.Make(MakeString("CONSTRUCT BUG: ~S").Id(),MakeConstantList(self.Id()).Id())).Close()
          /* If-4 */} 
        /* ERROR PROTECTION INSERTED (kind-Result) */
        if ErrorIn(kind_try06384) {Result = kind_try06384
        } else {
        kind = ToClass(OBJ(kind_try06384))
        /* Let:4 */{ 
          var count_try int  = 0
          /* noccur = 3 */
          /* Let:5 */{ 
            var t *ClaireAny  
            /* noccur = 1 */
            if (Core.F_get_property(C_of,ToObject(self.Id())) != CNULL) /* If:6 */{ 
              t = ANY(Core.F_CALL(C_of,ARGS(EID{self.Id(),0})))
              } else {
              t = C_void.Id()
              /* If-6 */} 
            F_Generate_new_block_string(MakeString("Construct"))
            var g0639I *ClaireBoolean  
            var g0639I_try06406 EID 
            /* Let:6 */{ 
              var g0641UU *ClaireAny  
              /* noccur = 1 */
              var g0641UU_try06427 EID 
              /* For:7 */{ 
                var x *ClaireAny  
                _ = x
                g0641UU_try06427= EID{CFALSE.Id(),0}
                var x_support *ClaireList  
                x_support = self.Args
                x_len := x_support.Length()
                for i_it := 0; i_it < x_len; i_it++ { 
                  x = x_support.At(i_it)
                  var void_try9 EID 
                  _ = void_try9
                  var g0643I *ClaireBoolean  
                  var g0643I_try06449 EID 
                  /* Let:9 */{ 
                    var g0645UU *ClaireBoolean  
                    /* noccur = 1 */
                    var g0645UU_try064610 EID 
                    g0645UU_try064610 = F_Generate_g_clean_any(x)
                    /* ERROR PROTECTION INSERTED (g0645UU-g0643I_try06449) */
                    if ErrorIn(g0645UU_try064610) {g0643I_try06449 = g0645UU_try064610
                    } else {
                    g0645UU = ToBoolean(OBJ(g0645UU_try064610))
                    g0643I_try06449 = EID{g0645UU.Not.Id(),0}
                    }
                    /* Let-9 */} 
                  /* ERROR PROTECTION INSERTED (g0643I-void_try9) */
                  if ErrorIn(g0643I_try06449) {void_try9 = g0643I_try06449
                  } else {
                  g0643I = ToBoolean(OBJ(g0643I_try06449))
                  if (g0643I == CTRUE) /* If:9 */{ 
                     /*v = g0641UU_try06427, s =EID*/
g0641UU_try06427 = EID{CTRUE.Id(),0}
                    break
                    } else {
                    void_try9 = EID{CFALSE.Id(),0}
                    /* If-9 */} 
                  }
                  /* ERROR PROTECTION INSERTED (void_try9-g0641UU_try06427) */
                  if ErrorIn(void_try9) {g0641UU_try06427 = void_try9
                  g0641UU_try06427 = void_try9
                  break
                  } else {
                  }
                  /* loop-8 */} 
                /* For-7 */} 
              /* ERROR PROTECTION INSERTED (g0641UU-g0639I_try06406) */
              if ErrorIn(g0641UU_try06427) {g0639I_try06406 = g0641UU_try06427
              } else {
              g0641UU = ANY(g0641UU_try06427)
              g0639I_try06406 = EID{F_boolean_I_any(g0641UU).Id(),0}
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (g0639I-Result) */
            if ErrorIn(g0639I_try06406) {Result = g0639I_try06406
            } else {
            g0639I = ToBoolean(OBJ(g0639I_try06406))
            if (g0639I == CTRUE) /* If:6 */{ 
              F_Generate_var_declaration_string(v2,C_any,1)
              Result = EVOID
              } else {
              Result = EID{CFALSE.Id(),0}
              /* If-6 */} 
            }
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            F_c_princ_string(v)
            PRINC("= ")
            Result = F_Generate_cast_prefix_class(kind,s)
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            if (kind.Id() == C_tuple.Id()) /* If:6 */{ 
              PRINC("Make")
              Result = EVOID
              } else {
              if (Core.F_get_property(C_of,ToObject(self.Id())) != CNULL) /* If:7 */{ 
                /* Let:8 */{ 
                  var g0647UU *ClaireAny  
                  /* noccur = 1 */
                  var g0647UU_try06489 EID 
                  g0647UU_try06489 = Core.F_CALL(Optimize.C_c_code,ARGS(t.ToEID(),EID{C_object.Id(),0}))
                  /* ERROR PROTECTION INSERTED (g0647UU-Result) */
                  if ErrorIn(g0647UU_try06489) {Result = g0647UU_try06489
                  } else {
                  g0647UU = ANY(g0647UU_try06489)
                  Result = Core.F_CALL(C_Generate_g_expression,ARGS(g0647UU.ToEID(),EID{C_type.Id(),0}))
                  }
                  /* Let-8 */} 
                } else {
                PRINC("ToType(CEMPTY.Id())")
                Result = EVOID
                /* If-7 */} 
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              PRINC(".")
              Result = EVOID
              }
              /* If-6 */} 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC("Empty")
            F_Generate_cap_short_symbol(kind.Name)
            PRINC("()")
            F_Generate_cast_post_class(kind,s)
            PRINC("")
            Result = EVOID
            }}
            {
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
                /* Let:8 */{ 
                  var f *ClaireBoolean  
                  /* noccur = 2 */
                  var f_try06499 EID 
                  f_try06499 = F_Generate_g_clean_any(x)
                  /* ERROR PROTECTION INSERTED (f-void_try8) */
                  if ErrorIn(f_try06499) {void_try8 = f_try06499
                  } else {
                  f = ToBoolean(OBJ(f_try06499))
                  /* Let:9 */{ 
                    var try_ask *ClaireBoolean  
                    /* noccur = 2 */
                    var try_ask_try065010 EID 
                    try_ask_try065010 = Optimize.F_Compile_g_throw_any(x)
                    /* ERROR PROTECTION INSERTED (try_ask-void_try8) */
                    if ErrorIn(try_ask_try065010) {void_try8 = try_ask_try065010
                    } else {
                    try_ask = ToBoolean(OBJ(try_ask_try065010))
                    F_Generate_breakline_void()
                    if (try_ask == CTRUE) /* If:10 */{ 
                      count_try = (count_try+1)
                      /* If-10 */} 
                    if (f != CTRUE) /* If:10 */{ 
                      if (try_ask == CTRUE) /* If:11 */{ 
                        void_try8 = Core.F_CALL(C_Generate_g_try,ARGS(x.ToEID(),
                          EID{(v2).Id(),0},
                          EID{C_any.Id(),0},
                          EID{(v).Id(),0},
                          EID{CFALSE.Id(),0}))
                        } else {
                        void_try8 = F_Generate_statement_any(x,C_any,v2,loop)
                        /* If-11 */} 
                      } else {
                      void_try8 = EID{CFALSE.Id(),0}
                      /* If-10 */} 
                    /* ERROR PROTECTION INSERTED (void_try8-void_try8) */
                    if ErrorIn(void_try8) {Result = void_try8
                    break
                    } else {
                    void_try8 = F_Generate_cast_prefix_class(s,kind)
                    /* ERROR PROTECTION INSERTED (void_try8-void_try8) */
                    if ErrorIn(void_try8) {Result = void_try8
                    break
                    } else {
                    F_c_princ_string(v)
                    F_Generate_cast_post_class(s,kind)
                    PRINC(".AddFast(")
                    if (f == CTRUE) /* If:10 */{ 
                      void_try8 = Core.F_CALL(C_Generate_g_expression,ARGS(x.ToEID(),EID{C_any.Id(),0}))
                      } else {
                      F_c_princ_string(v2)
                      void_try8 = EVOID
                      /* If-10 */} 
                    /* ERROR PROTECTION INSERTED (void_try8-void_try8) */
                    if ErrorIn(void_try8) {Result = void_try8
                    break
                    } else {
                    PRINC(")")
                    void_try8 = EVOID
                    }}
                    {
                    }}
                    }
                    /* Let-9 */} 
                  }
                  /* Let-8 */} 
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
            F_Generate_close_try_integer(count_try)
            F_Generate_close_block_string(MakeString("Construct"))
            Result = EVOID
            }}}
            /* Let-5 */} 
          /* Let-4 */} 
        }
        /* Let-3 */} 
      /* Let-2 */} 
    }
    return Result} 
  
// The EID go function for: g_statement @ Construct (throw: true) 
func E_Generate_g_statement_Construct (self EID,s EID,v EID,err EID,loop EID) EID { 
    return /*(sm for g_statement @ Construct= EID)*/ F_Generate_g_statement_Construct(Language.To_Construct(OBJ(self)),
      ToClass(OBJ(s)),
      ToString(OBJ(v)),
      ToBoolean(OBJ(err)),
      ANY(loop) )} 
  
// A if is easy to compile. We check if the logical compiler can be used
// we now assume that the test retuns a boolean !
// note that in GO the "} else " pattern is tricky
/* {1} OPT.The go function for: g_statement(self:If,s:class,v:string,err:boolean,loop:any) [] */
func F_Generate_g_statement_If (self *Language.If ,s *ClaireClass ,v *ClaireString ,err *ClaireBoolean ,loop *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var try_ask *ClaireBoolean  
      /* noccur = 2 */
      var try_ask_try06513 EID 
      try_ask_try06513 = Optimize.F_Compile_g_throw_any(self.Test)
      /* ERROR PROTECTION INSERTED (try_ask-Result) */
      if ErrorIn(try_ask_try06513) {Result = try_ask_try06513
      } else {
      try_ask = ToBoolean(OBJ(try_ask_try06513))
      var g0652I *ClaireBoolean  
      var g0652I_try06533 EID 
      g0652I_try06533 = F_Generate_g_clean_any(self.Test)
      /* ERROR PROTECTION INSERTED (g0652I-Result) */
      if ErrorIn(g0652I_try06533) {Result = g0652I_try06533
      } else {
      g0652I = ToBoolean(OBJ(g0652I_try06533))
      if (g0652I == CTRUE) /* If:3 */{ 
        PRINC("if ")
        Result = Core.F_CALL(Optimize.C_Compile_bool_exp,ARGS(self.Test.ToEID(),EID{CTRUE.Id(),0}))
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC(" ")
        Result = EVOID
        }
        {
        F_Generate_new_block_string(MakeString("If"))
        Result = EVOID
        }
        } else {
        /* Let:4 */{ 
          var v2 *ClaireString  
          /* noccur = 4 */
          var v2_try06545 EID 
          v2_try06545 = F_Generate_c_string_go_producer2(ToGenerateGoProducer(Optimize.C_PRODUCER.Value),F_append_symbol(Core.F_gensym_void(),MakeString("I").Id()))
          /* ERROR PROTECTION INSERTED (v2-Result) */
          if ErrorIn(v2_try06545) {Result = v2_try06545
          } else {
          v2 = ToString(OBJ(v2_try06545))
          F_Generate_var_declaration_string(v2,C_boolean,1)
          if (try_ask == CTRUE) /* If:5 */{ 
            Result = Core.F_CALL(C_Generate_g_try,ARGS(self.Test.ToEID(),
              EID{(v2).Id(),0},
              EID{C_boolean.Id(),0},
              EID{(v).Id(),0},
              EID{CFALSE.Id(),0}))
            } else {
            Result = F_Generate_statement_any(self.Test,C_boolean,v2,loop)
            /* If-5 */} 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC("if (")
          F_princ_string(v2)
          PRINC(" == CTRUE) ")
          F_Generate_new_block_string(MakeString("If"))
          PRINC("")
          Result = EVOID
          }
          }
          /* Let-4 */} 
        /* If-3 */} 
      }
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = F_Generate_statement_any(self.Arg,s,v,loop)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      if ((Equal(self.Other,CNIL.Id()) == CTRUE) || 
          ((self.Other == CFALSE.Id()) && 
              (s.Id() == C_void.Id()))) /* If:3 */{ 
        F_Generate_close_block_string(MakeString("If"))
        Result = EVOID
        } else {
        var g0655I *ClaireBoolean  
        var g0655I_try06564 EID 
        /* and:4 */{ 
          var v_and4 *ClaireBoolean  
          
          v_and4 = self.Other.Isa.IsIn(Language.C_If)
          if (v_and4 == CFALSE) {g0655I_try06564 = EID{CFALSE.Id(),0}
          } else /* arg:5 */{ 
            var v_and4_try06576 EID 
            v_and4_try06576 = F_Generate_g_func_any(ANY(Core.F_CALL(Language.C_iClaire_test,ARGS(self.Other.ToEID()))))
            /* ERROR PROTECTION INSERTED (v_and4-g0655I_try06564) */
            if ErrorIn(v_and4_try06576) {g0655I_try06564 = v_and4_try06576
            } else {
            v_and4 = ToBoolean(OBJ(v_and4_try06576))
            if (v_and4 == CFALSE) {g0655I_try06564 = EID{CFALSE.Id(),0}
            } else /* arg:6 */{ 
              var v_and4_try06587 EID 
              /* Let:7 */{ 
                var g0659UU *ClaireBoolean  
                /* noccur = 1 */
                var g0659UU_try06608 EID 
                g0659UU_try06608 = Optimize.F_Compile_g_throw_any(ANY(Core.F_CALL(Language.C_iClaire_test,ARGS(self.Other.ToEID()))))
                /* ERROR PROTECTION INSERTED (g0659UU-v_and4_try06587) */
                if ErrorIn(g0659UU_try06608) {v_and4_try06587 = g0659UU_try06608
                } else {
                g0659UU = ToBoolean(OBJ(g0659UU_try06608))
                v_and4_try06587 = EID{g0659UU.Not.Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (v_and4-g0655I_try06564) */
              if ErrorIn(v_and4_try06587) {g0655I_try06564 = v_and4_try06587
              } else {
              v_and4 = ToBoolean(OBJ(v_and4_try06587))
              if (v_and4 == CFALSE) {g0655I_try06564 = EID{CFALSE.Id(),0}
              } else /* arg:7 */{ 
                g0655I_try06564 = EID{CTRUE.Id(),0}/* arg-7 */} 
              /* arg-6 */} 
            /* arg-5 */} 
          }}
          /* and-4 */} 
        /* ERROR PROTECTION INSERTED (g0655I-Result) */
        if ErrorIn(g0655I_try06564) {Result = g0655I_try06564
        } else {
        g0655I = ToBoolean(OBJ(g0655I_try06564))
        if (g0655I == CTRUE) /* If:4 */{ 
          F_Generate_finish_block_string(MakeString("If"))
          PRINC(" else ")
          Result = Core.F_CALL(C_Generate_g_statement,ARGS(self.Other.ToEID(),
            EID{s.Id(),0},
            EID{(v).Id(),0},
            EID{CFALSE.Id(),0},
            loop.ToEID()))
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC("")
          Result = EVOID
          }
          } else {
          PRINC("} else {")
          F_Generate_breakline_void()
          Result = F_Generate_statement_any(self.Other,s,v,loop)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          F_Generate_close_block_string(MakeString("If"))
          PRINC("")
          Result = EVOID
          }
          /* If-4 */} 
        }
        /* If-3 */} 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      if (try_ask == CTRUE) /* If:3 */{ 
        F_Generate_close_try_integer(1)
        Result = EVOID
        } else {
        Result = EID{CFALSE.Id(),0}
        /* If-3 */} 
      }}}
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: g_statement @ If (throw: true) 
func E_Generate_g_statement_If (self EID,s EID,v EID,err EID,loop EID) EID { 
    return /*(sm for g_statement @ If= EID)*/ F_Generate_g_statement_If(Language.To_If(OBJ(self)),
      ToClass(OBJ(s)),
      ToString(OBJ(v)),
      ToBoolean(OBJ(err)),
      ANY(loop) )} 
  
// --------------- logical combinations and/or -------------------------------
// note: we cannot use unfolding because the order of evaluation is important !
// AND is compiled with IF: as soon as an argument is false, the result is false.
/* {1} OPT.The go function for: g_statement(self:And,s:class,v:string,err:boolean,loop:any) [] */
func F_Generate_g_statement_And (self *Language.And ,s *ClaireClass ,v *ClaireString ,err *ClaireBoolean ,loop *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var v2 *ClaireString   = F_Generate_check_var_string(MakeString("v_and"))
      /* noccur = 4 */
      /* Let:3 */{ 
        var count_try int  = 0
        /* noccur = 3 */
        F_Generate_new_block_string(MakeString("and"))
        F_Generate_var_declaration_string(v2,C_boolean,1)
        F_Generate_breakline_void()
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
            /* Let:6 */{ 
              var try_ask *ClaireBoolean  
              /* noccur = 1 */
              var try_ask_try06617 EID 
              try_ask_try06617 = Optimize.F_Compile_g_throw_any(x)
              /* ERROR PROTECTION INSERTED (try_ask-void_try6) */
              if ErrorIn(try_ask_try06617) {void_try6 = try_ask_try06617
              } else {
              try_ask = ToBoolean(OBJ(try_ask_try06617))
              if (try_ask == CTRUE) /* If:7 */{ 
                void_try6 = Core.F_CALL(C_Generate_g_try,ARGS(x.ToEID(),
                  EID{(v2).Id(),0},
                  EID{C_boolean.Id(),0},
                  EID{(v).Id(),0},
                  EID{CFALSE.Id(),0}))
                /* ERROR PROTECTION INSERTED (void_try6-void_try6) */
                if ErrorIn(void_try6) {Result = void_try6
                break
                } else {
                count_try = (count_try+1)
                void_try6 = EID{C__INT,IVAL(count_try)}
                }
                } else {
                void_try6 = F_Generate_statement_any(x,C_boolean,v2,loop)
                /* If-7 */} 
              /* ERROR PROTECTION INSERTED (void_try6-void_try6) */
              if ErrorIn(void_try6) {Result = void_try6
              break
              } else {
              PRINC("if (")
              F_c_princ_string(v2)
              PRINC(" == CFALSE) {")
              if (s.Id() != C_void.Id()) /* If:7 */{ 
                F_c_princ_string(v)
                PRINC(" = ")
                void_try6 = F_Generate_cast_prefix_class(C_boolean,s)
                /* ERROR PROTECTION INSERTED (void_try6-void_try6) */
                if ErrorIn(void_try6) {Result = void_try6
                break
                } else {
                PRINC("CFALSE")
                F_Generate_cast_post_class(C_boolean,s)
                PRINC("")
                void_try6 = EVOID
                }
                } else {
                void_try6 = EID{CFALSE.Id(),0}
                /* If-7 */} 
              /* ERROR PROTECTION INSERTED (void_try6-void_try6) */
              if ErrorIn(void_try6) {Result = void_try6
              break
              } else {
              F_Generate_breakline_void()
              PRINC("} else ")
              F_Generate_new_block_string(MakeString("arg"))
              PRINC("")
              void_try6 = EVOID
              }
              {
              }}
              }
              /* Let-6 */} 
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
        if (s.Id() != C_void.Id()) /* If:4 */{ 
          F_c_princ_string(v)
          PRINC(" = ")
          Result = F_Generate_cast_prefix_class(C_boolean,s)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC("CTRUE")
          F_Generate_cast_post_class(C_boolean,s)
          PRINC("")
          Result = EVOID
          }
          } else {
          Result = EID{CFALSE.Id(),0}
          /* If-4 */} 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        /* For:4 */{ 
          var x *ClaireAny  
          _ = x
          var x_support *ClaireList  
          x_support = self.Args
          x_len := x_support.Length()
          for i_it := 0; i_it < x_len; i_it++ { 
            x = x_support.At(i_it)
            F_Generate_close_block_string(MakeString("arg"))
            /* loop-5 */} 
          /* For-4 */} 
        F_Generate_close_try_integer(count_try)
        F_Generate_close_block_string(MakeString("and"))
        Result = EVOID
        }}
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: g_statement @ And (throw: true) 
func E_Generate_g_statement_And (self EID,s EID,v EID,err EID,loop EID) EID { 
    return /*(sm for g_statement @ And= EID)*/ F_Generate_g_statement_And(Language.To_And(OBJ(self)),
      ToClass(OBJ(s)),
      ToString(OBJ(v)),
      ToBoolean(OBJ(err)),
      ANY(loop) )} 
  
// same thing for OR
/* {1} OPT.The go function for: g_statement(self:Or,s:class,v:string,err:boolean,loop:any) [] */
func F_Generate_g_statement_Or (self *Language.Or ,s *ClaireClass ,v *ClaireString ,err *ClaireBoolean ,loop *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var v2 *ClaireString   = F_Generate_check_var_string(MakeString("v_or"))
      /* noccur = 4 */
      /* Let:3 */{ 
        var count_try int  = 0
        /* noccur = 3 */
        F_Generate_new_block_string(MakeString("or"))
        F_Generate_var_declaration_string(v2,C_boolean,1)
        F_Generate_breakline_void()
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
            /* Let:6 */{ 
              var try_ask *ClaireBoolean  
              /* noccur = 1 */
              var try_ask_try06627 EID 
              try_ask_try06627 = Optimize.F_Compile_g_throw_any(x)
              /* ERROR PROTECTION INSERTED (try_ask-void_try6) */
              if ErrorIn(try_ask_try06627) {void_try6 = try_ask_try06627
              } else {
              try_ask = ToBoolean(OBJ(try_ask_try06627))
              if (try_ask == CTRUE) /* If:7 */{ 
                void_try6 = Core.F_CALL(C_Generate_g_try,ARGS(x.ToEID(),
                  EID{(v2).Id(),0},
                  EID{C_boolean.Id(),0},
                  EID{(v).Id(),0},
                  EID{CFALSE.Id(),0}))
                /* ERROR PROTECTION INSERTED (void_try6-void_try6) */
                if ErrorIn(void_try6) {Result = void_try6
                break
                } else {
                count_try = (count_try+1)
                void_try6 = EID{C__INT,IVAL(count_try)}
                }
                } else {
                void_try6 = F_Generate_statement_any(x,C_boolean,v2,loop)
                /* If-7 */} 
              /* ERROR PROTECTION INSERTED (void_try6-void_try6) */
              if ErrorIn(void_try6) {Result = void_try6
              break
              } else {
              PRINC("if (")
              F_c_princ_string(v2)
              PRINC(" == CTRUE) {")
              if (s.Id() != C_void.Id()) /* If:7 */{ 
                F_c_princ_string(v)
                PRINC(" = ")
                void_try6 = F_Generate_cast_prefix_class(C_boolean,s)
                /* ERROR PROTECTION INSERTED (void_try6-void_try6) */
                if ErrorIn(void_try6) {Result = void_try6
                break
                } else {
                PRINC("CTRUE")
                F_Generate_cast_post_class(C_boolean,s)
                PRINC("")
                void_try6 = EVOID
                }
                } else {
                void_try6 = EID{CFALSE.Id(),0}
                /* If-7 */} 
              /* ERROR PROTECTION INSERTED (void_try6-void_try6) */
              if ErrorIn(void_try6) {Result = void_try6
              break
              } else {
              F_Generate_breakline_void()
              PRINC("} else ")
              F_Generate_new_block_string(MakeString("or"))
              PRINC("")
              void_try6 = EVOID
              }
              {
              }}
              }
              /* Let-6 */} 
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
        if (s.Id() != C_void.Id()) /* If:4 */{ 
          F_c_princ_string(v)
          PRINC(" = ")
          Result = F_Generate_cast_prefix_class(C_boolean,s)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC("CFALSE")
          F_Generate_cast_post_class(C_boolean,s)
          PRINC("")
          Result = EVOID
          }
          } else {
          Result = EID{CFALSE.Id(),0}
          /* If-4 */} 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        /* For:4 */{ 
          var x *ClaireAny  
          _ = x
          var x_support *ClaireList  
          x_support = self.Args
          x_len := x_support.Length()
          for i_it := 0; i_it < x_len; i_it++ { 
            x = x_support.At(i_it)
            F_Generate_close_block_string(MakeString("org"))
            /* loop-5 */} 
          /* For-4 */} 
        F_Generate_close_try_integer(count_try)
        F_Generate_close_block_string(MakeString("or"))
        Result = EVOID
        }}
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: g_statement @ Or (throw: true) 
func E_Generate_g_statement_Or (self EID,s EID,v EID,err EID,loop EID) EID { 
    return /*(sm for g_statement @ Or= EID)*/ F_Generate_g_statement_Or(Language.To_Or(OBJ(self)),
      ToClass(OBJ(s)),
      ToString(OBJ(v)),
      ToBoolean(OBJ(err)),
      ANY(loop) )} 
  
// Here this is the simple assignment, with a true variable
// note that the last line (assigning the value to result is only OK if no error)
/* {1} OPT.The go function for: g_statement(self:Assign,s:class,v:string,err:boolean,loop:any) [] */
func F_Generate_g_statement_Assign (self *Language.Assign ,s *ClaireClass ,v *ClaireString ,err *ClaireBoolean ,loop *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var _Zvar *ClaireAny   = self.ClaireVar
      /* noccur = 2 */
      /* Let:3 */{ 
        var x *ClaireAny   = self.Arg
        /* noccur = 3 */
        /* Let:4 */{ 
          var v2 *ClaireAny  
          /* noccur = 4 */
          var v2_try06635 EID 
          v2_try06635 = Core.F_CALL(C_Generate_c_string,ARGS(EID{Optimize.C_PRODUCER.Value,0},_Zvar.ToEID()))
          /* ERROR PROTECTION INSERTED (v2-Result) */
          if ErrorIn(v2_try06635) {Result = v2_try06635
          } else {
          v2 = ANY(v2_try06635)
          /* Let:5 */{ 
            var _Zrange *ClaireClass   = ToTypeExpression(OBJ(Core.F_CALL(C_range,ARGS(_Zvar.ToEID())))).Class_I()
            /* noccur = 4 */
            /* Let:6 */{ 
              var try_ask *ClaireBoolean  
              /* noccur = 1 */
              var try_ask_try06647 EID 
              try_ask_try06647 = Optimize.F_Compile_g_throw_any(x)
              /* ERROR PROTECTION INSERTED (try_ask-Result) */
              if ErrorIn(try_ask_try06647) {Result = try_ask_try06647
              } else {
              try_ask = ToBoolean(OBJ(try_ask_try06647))
              if (try_ask == CTRUE) /* If:7 */{ 
                Result = Core.F_CALL(C_Generate_g_try,ARGS(x.ToEID(),
                  v2.ToEID(),
                  EID{_Zrange.Id(),0},
                  EID{(v).Id(),0},
                  EID{CFALSE.Id(),0}))
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                if (s.Id() != C_void.Id()) /* If:8 */{ 
                  F_c_princ_string(v)
                  PRINC(" = ")
                  Result = F_Generate_use_variable_string(ToString(v2),s,_Zrange)
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  F_Generate_breakline_void()
                  PRINC("")
                  Result = EVOID
                  }
                  } else {
                  Result = EID{CFALSE.Id(),0}
                  /* If-8 */} 
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                F_Generate_close_try_integer(1)
                Result = EVOID
                }}
                } else {
                Result = F_Generate_statement_any(x,_Zrange,ToString(v2),loop)
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                if (s.Id() != C_void.Id()) /* If:8 */{ 
                  F_c_princ_string(v)
                  PRINC(" = ")
                  Result = F_Generate_use_variable_string(ToString(v2),s,_Zrange)
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  F_Generate_breakline_void()
                  PRINC("")
                  Result = EVOID
                  }
                  } else {
                  Result = EID{CFALSE.Id(),0}
                  /* If-8 */} 
                }
                /* If-7 */} 
              }
              /* Let-6 */} 
            /* Let-5 */} 
          }
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: g_statement @ Assign (throw: true) 
func E_Generate_g_statement_Assign (self EID,s EID,v EID,err EID,loop EID) EID { 
    return /*(sm for g_statement @ Assign= EID)*/ F_Generate_g_statement_Assign(Language.To_Assign(OBJ(self)),
      ToClass(OBJ(s)),
      ToString(OBJ(v)),
      ToBoolean(OBJ(err)),
      ANY(loop) )} 
  
// This is the global variable assignment - global variables exist in go so this is pretty simple
// note that the tricky part is the store management
/* {1} OPT.The go function for: g_statement(self:Gassign,s:class,v:string,err:boolean,loop:any) [] */
func F_Generate_g_statement_Gassign (self *Language.Gassign ,s *ClaireClass ,v *ClaireString ,err *ClaireBoolean ,loop *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var _Zvar *Core.GlobalVariable   = self.ClaireVar
      /* noccur = 0 */
      _ = _Zvar
      /* Let:3 */{ 
        var x *ClaireAny   = self.Arg
        /* noccur = 5 */
        var g0665I *ClaireBoolean  
        var g0665I_try06664 EID 
        /* and:4 */{ 
          var v_and4 *ClaireBoolean  
          
          var v_and4_try06675 EID 
          v_and4_try06675 = F_Generate_g_func_any(x)
          /* ERROR PROTECTION INSERTED (v_and4-g0665I_try06664) */
          if ErrorIn(v_and4_try06675) {g0665I_try06664 = v_and4_try06675
          } else {
          v_and4 = ToBoolean(OBJ(v_and4_try06675))
          if (v_and4 == CFALSE) {g0665I_try06664 = EID{CFALSE.Id(),0}
          } else /* arg:5 */{ 
            v_and4 = Equal(s.Id(),C_void.Id())
            if (v_and4 == CFALSE) {g0665I_try06664 = EID{CFALSE.Id(),0}
            } else /* arg:6 */{ 
              v_and4 = self.ClaireVar.Store_ask.Not
              if (v_and4 == CFALSE) {g0665I_try06664 = EID{CFALSE.Id(),0}
              } else /* arg:7 */{ 
                g0665I_try06664 = EID{CTRUE.Id(),0}/* arg-7 */} 
              /* arg-6 */} 
            /* arg-5 */} 
          }
          /* and-4 */} 
        /* ERROR PROTECTION INSERTED (g0665I-Result) */
        if ErrorIn(g0665I_try06664) {Result = g0665I_try06664
        } else {
        g0665I = ToBoolean(OBJ(g0665I_try06664))
        if (g0665I == CTRUE) /* If:4 */{ 
          ToGenerateGoProducer(Optimize.C_PRODUCER.Value).GlobalVar(self.ClaireVar)
          PRINC(" = ")
          Result = Core.F_CALL(C_Generate_g_expression,ARGS(x.ToEID(),EID{C_any.Id(),0}))
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          F_Generate_breakline_void()
          PRINC("")
          Result = EVOID
          }
          } else {
          /* Let:5 */{ 
            var v2 *ClaireString   = F_Generate_check_var_string(MakeString("v_gassign"))
            /* noccur = 8 */
            /* Let:6 */{ 
              var try_ask *ClaireBoolean  
              /* noccur = 3 */
              var try_ask_try06687 EID 
              try_ask_try06687 = Optimize.F_Compile_g_throw_any(x)
              /* ERROR PROTECTION INSERTED (try_ask-Result) */
              if ErrorIn(try_ask_try06687) {Result = try_ask_try06687
              } else {
              try_ask = ToBoolean(OBJ(try_ask_try06687))
              if ((try_ask != CTRUE) && 
                  ((s.Id() == C_void.Id()) || 
                      (s.Id() == C_any.Id()))) /* If:7 */{ 
                v2 = v
                } else {
                F_Generate_var_declaration_string(v2,C_any,1)
                /* If-7 */} 
              if (try_ask == CTRUE) /* If:7 */{ 
                Result = Core.F_CALL(C_Generate_g_try,ARGS(x.ToEID(),
                  EID{(v2).Id(),0},
                  EID{C_any.Id(),0},
                  EID{(v).Id(),0},
                  EID{CFALSE.Id(),0}))
                } else {
                Result = F_Generate_statement_any(x,C_any,v2,loop)
                /* If-7 */} 
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              if (self.ClaireVar.Store_ask == CTRUE) /* If:7 */{ 
                ToGenerateGoProducer(Optimize.C_PRODUCER.Value).GlobalVar(self.ClaireVar)
                PRINC(".StoreObj(2,")
                F_c_princ_string(v2)
                PRINC(");")
                } else {
                ToGenerateGoProducer(Optimize.C_PRODUCER.Value).GlobalVar(self.ClaireVar)
                PRINC(" = ")
                F_c_princ_string(v2)
                F_Generate_breakline_void()
                PRINC("")
                /* If-7 */} 
              if ((s.Id() != C_void.Id()) && 
                  (v.Value != v2.Value)) /* If:7 */{ 
                F_c_princ_string(v)
                PRINC(" = ")
                Result = F_Generate_use_variable_string(v2,s,C_any)
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                F_Generate_breakline_void()
                PRINC("")
                Result = EVOID
                }
                } else {
                Result = EID{CFALSE.Id(),0}
                /* If-7 */} 
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              if (try_ask == CTRUE) /* If:7 */{ 
                F_Generate_close_block_void()
                Result = EVOID
                } else {
                Result = EID{CFALSE.Id(),0}
                /* If-7 */} 
              }}
              }
              /* Let-6 */} 
            /* Let-5 */} 
          /* If-4 */} 
        }
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: g_statement @ Gassign (throw: true) 
func E_Generate_g_statement_Gassign (self EID,s EID,v EID,err EID,loop EID) EID { 
    return /*(sm for g_statement @ Gassign= EID)*/ F_Generate_g_statement_Gassign(Language.To_Gassign(OBJ(self)),
      ToClass(OBJ(s)),
      ToString(OBJ(v)),
      ToBoolean(OBJ(err)),
      ANY(loop) )} 
  
//**********************************************************************
//*          Part 3: iteration                                         *
//**********************************************************************
// we know to iterate sets or lists in Go
// the optimizer should give use something that is properly typed
/* {1} OPT.The go function for: bag_class(self:any) [] */
func F_Generate_bag_class_any (self *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var s *ClaireAny  
      /* noccur = 5 */
      var s_try06693 EID 
      s_try06693 = Core.F_CALL(Optimize.C_c_type,ARGS(self.ToEID()))
      /* ERROR PROTECTION INSERTED (s-Result) */
      if ErrorIn(s_try06693) {Result = s_try06693
      } else {
      s = ANY(s_try06693)
      if ((ToType(s).Included(ToType(C_list.Id())) == CTRUE) || 
          ((ToType(s).Included(ToType(C_tuple.Id())) == CTRUE) || 
            (ToType(s).Included(ToType(C_array.Id())) == CTRUE))) /* If:3 */{ 
        Result = EID{C_list.Id(),0}
        /* If!3 */}  else if (ToType(s).Included(ToType(C_set.Id())) == CTRUE) /* If:3 */{ 
        Result = EID{C_set.Id(),0}
        } else {
        Result = ToException(Core.C_general_error.Make(MakeString("bag_class(~S) returns ~S: cannot use in for").Id(),MakeConstantList(self,s).Id())).Close()
        /* If-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: bag_class @ any (throw: true) 
func E_Generate_bag_class_any (self EID) EID { 
    return /*(sm for bag_class @ any= EID)*/ F_Generate_bag_class_any(ANY(self) )} 
  
// generates the iteration code for a "for x in S ..." expression , once
// all optimization based on code substitution have been performed.
// very nice in go, except that we have to handle error
// if g_member(%set) is native (anything but any) we use the native go form
/* {1} OPT.The go function for: g_statement(self:For,s:class,v:string,err:boolean,loop:any) [] */
func F_Generate_g_statement_For (self *Language.For ,s *ClaireClass ,v *ClaireString ,err *ClaireBoolean ,loop *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var v2 *ClaireString  
      /* noccur = 6 */
      var v2_try06703 EID 
      v2_try06703 = F_Generate_c_string_go_producer1(ToGenerateGoProducer(Optimize.C_PRODUCER.Value),self.ClaireVar)
      /* ERROR PROTECTION INSERTED (v2-Result) */
      if ErrorIn(v2_try06703) {Result = v2_try06703
      } else {
      v2 = ToString(OBJ(v2_try06703))
      /* Let:3 */{ 
        var count_try int  = 0
        /* noccur = 5 */
        /* Let:4 */{ 
          var v2_range *ClaireClass   = self.ClaireVar.Range.Class_I()
          /* noccur = 5 */
          /* Let:5 */{ 
            var v4 *ClaireString   = F_append_string(v2,MakeString("_iter"))
            /* noccur = 6 */
            /* Let:6 */{ 
              var _Zset *ClaireAny   = self.SetArg
              /* noccur = 11 */
              /* Let:7 */{ 
                var sbag *ClaireClass  
                /* noccur = 8 */
                var sbag_try06718 EID 
                sbag_try06718 = F_Generate_bag_class_any(_Zset)
                /* ERROR PROTECTION INSERTED (sbag-Result) */
                if ErrorIn(sbag_try06718) {Result = sbag_try06718
                } else {
                sbag = ToClass(OBJ(sbag_try06718))
                F_Generate_new_block_string(MakeString("For"))
                F_Generate_var_declaration_string(v2,v2_range,2)
                if (v2_range.Id() != C_any.Id()) /* If:8 */{ 
                  F_Generate_var_declaration_string(v4,C_any,1)
                  } else {
                  v4 = v2
                  /* If-8 */} 
                if (s.Id() != C_void.Id()) /* If:8 */{ 
                  F_c_princ_string(v)
                  PRINC("= ")
                  Result = F_Generate_cast_prefix_class(C_boolean,s)
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  PRINC("CFALSE")
                  F_Generate_cast_post_class(C_boolean,s)
                  F_Generate_breakline_void()
                  PRINC("")
                  Result = EVOID
                  }
                  } else {
                  Result = EID{CFALSE.Id(),0}
                  /* If-8 */} 
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                var g0672I *ClaireBoolean  
                var g0672I_try06738 EID 
                /* and:8 */{ 
                  var v_and8 *ClaireBoolean  
                  
                  var v_and8_try06749 EID 
                  v_and8_try06749 = F_Generate_g_clean_any(_Zset)
                  /* ERROR PROTECTION INSERTED (v_and8-g0672I_try06738) */
                  if ErrorIn(v_and8_try06749) {g0672I_try06738 = v_and8_try06749
                  } else {
                  v_and8 = ToBoolean(OBJ(v_and8_try06749))
                  if (v_and8 == CFALSE) {g0672I_try06738 = EID{CFALSE.Id(),0}
                  } else /* arg:9 */{ 
                    var v_and8_try067510 EID 
                    v_and8_try067510 = Optimize.F_Compile_designated_ask_any(_Zset)
                    /* ERROR PROTECTION INSERTED (v_and8-g0672I_try06738) */
                    if ErrorIn(v_and8_try067510) {g0672I_try06738 = v_and8_try067510
                    } else {
                    v_and8 = ToBoolean(OBJ(v_and8_try067510))
                    if (v_and8 == CFALSE) {g0672I_try06738 = EID{CFALSE.Id(),0}
                    } else /* arg:10 */{ 
                      var v_and8_try067611 EID 
                      /* or:11 */{ 
                        var v_or11 *ClaireBoolean  
                        
                        var v_or11_try067712 EID 
                        /* Let:12 */{ 
                          var g0678UU *ClaireClass  
                          /* noccur = 1 */
                          var g0678UU_try067913 EID 
                          g0678UU_try067913 = F_Generate_g_member_any(_Zset)
                          /* ERROR PROTECTION INSERTED (g0678UU-v_or11_try067712) */
                          if ErrorIn(g0678UU_try067913) {v_or11_try067712 = g0678UU_try067913
                          } else {
                          g0678UU = ToClass(OBJ(g0678UU_try067913))
                          v_or11_try067712 = EID{Core.F__I_equal_any(g0678UU.Id(),C_any.Id()).Id(),0}
                          }
                          /* Let-12 */} 
                        /* ERROR PROTECTION INSERTED (v_or11-v_and8_try067611) */
                        if ErrorIn(v_or11_try067712) {v_and8_try067611 = v_or11_try067712
                        } else {
                        v_or11 = ToBoolean(OBJ(v_or11_try067712))
                        if (v_or11 == CTRUE) {v_and8_try067611 = EID{CTRUE.Id(),0}
                        } else /* or:12 */{ 
                          v_or11 = Equal(sbag.Id(),C_set.Id())
                          if (v_or11 == CTRUE) {v_and8_try067611 = EID{CTRUE.Id(),0}
                          } else /* or:13 */{ 
                            v_and8_try067611 = EID{CFALSE.Id(),0}/* org-13 */} 
                          /* org-12 */} 
                        }
                        /* or-11 */} 
                      /* ERROR PROTECTION INSERTED (v_and8-g0672I_try06738) */
                      if ErrorIn(v_and8_try067611) {g0672I_try06738 = v_and8_try067611
                      } else {
                      v_and8 = ToBoolean(OBJ(v_and8_try067611))
                      if (v_and8 == CFALSE) {g0672I_try06738 = EID{CFALSE.Id(),0}
                      } else /* arg:11 */{ 
                        g0672I_try06738 = EID{CTRUE.Id(),0}/* arg-11 */} 
                      /* arg-10 */} 
                    /* arg-9 */} 
                  }}}
                  /* and-8 */} 
                /* ERROR PROTECTION INSERTED (g0672I-Result) */
                if ErrorIn(g0672I_try06738) {Result = g0672I_try06738
                } else {
                g0672I = ToBoolean(OBJ(g0672I_try06738))
                if (g0672I == CTRUE) /* If:8 */{ 
                  PRINC("for _,")
                  F_c_princ_string(v4)
                  PRINC(" = range(")
                  Result = Core.F_CALL(C_Generate_g_expression,ARGS(_Zset.ToEID(),EID{sbag.Id(),0}))
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  /* Let:9 */{ 
                    var g0680UU *ClaireClass  
                    /* noccur = 1 */
                    var g0680UU_try068110 EID 
                    g0680UU_try068110 = F_Generate_g_member_any(_Zset)
                    /* ERROR PROTECTION INSERTED (g0680UU-Result) */
                    if ErrorIn(g0680UU_try068110) {Result = g0680UU_try068110
                    } else {
                    g0680UU = ToClass(OBJ(g0680UU_try068110))
                    F_Generate_list_cast_values_class(sbag,g0680UU)
                    Result = EVOID
                    }
                    /* Let-9 */} 
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  PRINC(")")
                  F_Generate_new_block_string(MakeString("loop"))
                  PRINC("")
                  Result = EVOID
                  }}
                  } else {
                  /* Let:9 */{ 
                    var v3 *ClaireString   = F_append_string(v2,MakeString("_support"))
                    /* noccur = 6 */
                    /* Let:10 */{ 
                      var try_ask *ClaireBoolean  
                      /* noccur = 1 */
                      var try_ask_try068211 EID 
                      try_ask_try068211 = Optimize.F_Compile_g_throw_any(_Zset)
                      /* ERROR PROTECTION INSERTED (try_ask-Result) */
                      if ErrorIn(try_ask_try068211) {Result = try_ask_try068211
                      } else {
                      try_ask = ToBoolean(OBJ(try_ask_try068211))
                      F_Generate_var_declaration_string(v3,sbag,1)
                      if (try_ask == CTRUE) /* If:11 */{ 
                        Result = Core.F_CALL(C_Generate_g_try,ARGS(_Zset.ToEID(),
                          EID{(v3).Id(),0},
                          EID{sbag.Id(),0},
                          EID{(v).Id(),0},
                          EID{CFALSE.Id(),0}))
                        /* ERROR PROTECTION INSERTED (Result-Result) */
                        if !ErrorIn(Result) {
                        count_try = (count_try+1)
                        Result = EID{C__INT,IVAL(count_try)}
                        }
                        } else {
                        Result = F_Generate_statement_any(_Zset,sbag,v3,CFALSE.Id())
                        /* If-11 */} 
                      /* ERROR PROTECTION INSERTED (Result-Result) */
                      if !ErrorIn(Result) {
                      var g0683I *ClaireBoolean  
                      var g0683I_try068411 EID 
                      /* or:11 */{ 
                        var v_or11 *ClaireBoolean  
                        
                        var v_or11_try068512 EID 
                        /* Let:12 */{ 
                          var g0686UU *ClaireClass  
                          /* noccur = 1 */
                          var g0686UU_try068713 EID 
                          g0686UU_try068713 = F_Generate_g_member_any(_Zset)
                          /* ERROR PROTECTION INSERTED (g0686UU-v_or11_try068512) */
                          if ErrorIn(g0686UU_try068713) {v_or11_try068512 = g0686UU_try068713
                          } else {
                          g0686UU = ToClass(OBJ(g0686UU_try068713))
                          v_or11_try068512 = EID{Core.F__I_equal_any(g0686UU.Id(),C_any.Id()).Id(),0}
                          }
                          /* Let-12 */} 
                        /* ERROR PROTECTION INSERTED (v_or11-g0683I_try068411) */
                        if ErrorIn(v_or11_try068512) {g0683I_try068411 = v_or11_try068512
                        } else {
                        v_or11 = ToBoolean(OBJ(v_or11_try068512))
                        if (v_or11 == CTRUE) {g0683I_try068411 = EID{CTRUE.Id(),0}
                        } else /* or:12 */{ 
                          v_or11 = Equal(sbag.Id(),C_set.Id())
                          if (v_or11 == CTRUE) {g0683I_try068411 = EID{CTRUE.Id(),0}
                          } else /* or:13 */{ 
                            g0683I_try068411 = EID{CFALSE.Id(),0}/* org-13 */} 
                          /* org-12 */} 
                        }
                        /* or-11 */} 
                      /* ERROR PROTECTION INSERTED (g0683I-Result) */
                      if ErrorIn(g0683I_try068411) {Result = g0683I_try068411
                      } else {
                      g0683I = ToBoolean(OBJ(g0683I_try068411))
                      if (g0683I == CTRUE) /* If:11 */{ 
                        PRINC("for _,")
                        F_c_princ_string(v4)
                        PRINC(" = range(")
                        F_c_princ_string(v3)
                        /* Let:12 */{ 
                          var g0688UU *ClaireClass  
                          /* noccur = 1 */
                          var g0688UU_try068913 EID 
                          g0688UU_try068913 = F_Generate_g_member_any(_Zset)
                          /* ERROR PROTECTION INSERTED (g0688UU-Result) */
                          if ErrorIn(g0688UU_try068913) {Result = g0688UU_try068913
                          } else {
                          g0688UU = ToClass(OBJ(g0688UU_try068913))
                          F_Generate_list_cast_values_class(sbag,g0688UU)
                          Result = EVOID
                          }
                          /* Let-12 */} 
                        /* ERROR PROTECTION INSERTED (Result-Result) */
                        if !ErrorIn(Result) {
                        PRINC(")")
                        F_Generate_new_block_string(MakeString("loop2"))
                        PRINC("")
                        Result = EVOID
                        }
                        } else {
                        /* Let:12 */{ 
                          var v5 *ClaireString   = F_append_string(v2,MakeString("_len"))
                          /* noccur = 2 */
                          F_c_princ_string(v5)
                          PRINC(" := ")
                          F_c_princ_string(v3)
                          PRINC(".Length()")
                          F_Generate_breakline_void()
                          PRINC("")
                          PRINC("for i_it := 0; i_it < ")
                          F_c_princ_string(v5)
                          PRINC("; i_it++ ")
                          F_Generate_new_block_void()
                          F_c_princ_string(v4)
                          PRINC(" = ")
                          F_c_princ_string(v3)
                          PRINC(".At(i_it)")
                          F_Generate_breakline_void()
                          PRINC("")
                          PRINC("")
                          Result = EVOID
                          /* Let-12 */} 
                        /* If-11 */} 
                      }
                      }
                      }
                      /* Let-10 */} 
                    /* Let-9 */} 
                  /* If-8 */} 
                }
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                if (v2_range.Id() != C_any.Id()) /* If:8 */{ 
                  F_princ_string(v2)
                  PRINC(" = ")
                  Result = F_Generate_cast_prefix_class(C_any,v2_range)
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  F_princ_string(v4)
                  F_Generate_cast_post_class(C_any,v2_range)
                  F_Generate_breakline_void()
                  PRINC("")
                  Result = EVOID
                  }
                  } else {
                  Result = EID{CFALSE.Id(),0}
                  /* If-8 */} 
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                var g0690I *ClaireBoolean  
                var g0690I_try06918 EID 
                g0690I_try06918 = Optimize.F_Compile_g_throw_any(self.Arg)
                /* ERROR PROTECTION INSERTED (g0690I-Result) */
                if ErrorIn(g0690I_try06918) {Result = g0690I_try06918
                } else {
                g0690I = ToBoolean(OBJ(g0690I_try06918))
                if (g0690I == CTRUE) /* If:8 */{ 
                  Result = F_Generate_g_try_void_any(self.Arg,v,MakeTuple((v).Id(),s.Id()).Id())
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  count_try = (count_try+1)
                  Result = EID{C__INT,IVAL(count_try)}
                  }
                  } else {
                  Result = F_Generate_statement_any(self.Arg,C_void,v,MakeTuple((v).Id(),s.Id()).Id())
                  /* If-8 */} 
                }
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                F_Generate_close_try_integer(count_try)
                F_Generate_close_block_string(MakeString("loop"))
                F_Generate_close_block_string(MakeString("For"))
                Result = EVOID
                }}}}
                }
                /* Let-7 */} 
              /* Let-6 */} 
            /* Let-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: g_statement @ For (throw: true) 
func E_Generate_g_statement_For (self EID,s EID,v EID,err EID,loop EID) EID { 
    return /*(sm for g_statement @ For= EID)*/ F_Generate_g_statement_For(Language.To_For(OBJ(self)),
      ToClass(OBJ(s)),
      ToString(OBJ(v)),
      ToBoolean(OBJ(err)),
      ANY(loop) )} 
  
// for statement
// here the value is expected to be important, otherwise an error is raised.
// THIS IS ONLY APPLIED TO COLLECT(f(x) | s in S) on lists or sets (bags) => Image is delt with
// we currently do not use the native form => use At and Put to work on generic lists
/* {1} OPT.The go function for: g_statement(self:Iteration,s:class,v:string,err:boolean,loop:any) [] */
func F_Generate_g_statement_Iteration (self *Language.Iteration ,s *ClaireClass ,v *ClaireString ,err *ClaireBoolean ,loop *ClaireAny ) EID { 
    var Result EID 
    if (s.Id() == C_void.Id()) /* If:2 */{ 
      Result = ToException(Core.C_general_error.Make(MakeString("[203] you should have used a FOR ere:~S").Id(),MakeConstantList(self.Id()).Id())).Close()
      } else {
      Result = EID{CFALSE.Id(),0}
      /* If-2 */} 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    /* Let:2 */{ 
      var v2 *ClaireString  
      /* noccur = 3 */
      var v2_try06953 EID 
      v2_try06953 = F_Generate_c_string_go_producer1(ToGenerateGoProducer(Optimize.C_PRODUCER.Value),self.ClaireVar)
      /* ERROR PROTECTION INSERTED (v2-Result) */
      if ErrorIn(v2_try06953) {Result = v2_try06953
      } else {
      v2 = ToString(OBJ(v2_try06953))
      /* Let:3 */{ 
        var v2_range *ClaireClass   = self.ClaireVar.Range.Class_I()
        /* noccur = 3 */
        /* Let:4 */{ 
          var vlist *ClaireString   = F_Generate_check_var_string(MakeString("v_list"))
          /* noccur = 7 */
          /* Let:5 */{ 
            var vlocal *ClaireString   = F_Generate_check_var_string(MakeString("v_local"))
            /* noccur = 4 */
            /* Let:6 */{ 
              var bag_type *ClaireClass  
              /* noccur = 5 */
              var bag_type_try06967 EID 
              var g0697I *ClaireBoolean  
              var g0697I_try06987 EID 
              /* Let:7 */{ 
                var g0699UU *ClaireType  
                /* noccur = 1 */
                var g0699UU_try07008 EID 
                g0699UU_try07008 = Core.F_CALL(Optimize.C_c_type,ARGS(self.SetArg.ToEID()))
                /* ERROR PROTECTION INSERTED (g0699UU-g0697I_try06987) */
                if ErrorIn(g0699UU_try07008) {g0697I_try06987 = g0699UU_try07008
                } else {
                g0699UU = ToType(OBJ(g0699UU_try07008))
                g0697I_try06987 = EID{g0699UU.Included(ToType(C_set.Id())).Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (g0697I-bag_type_try06967) */
              if ErrorIn(g0697I_try06987) {bag_type_try06967 = g0697I_try06987
              } else {
              g0697I = ToBoolean(OBJ(g0697I_try06987))
              if (g0697I == CTRUE) /* If:7 */{ 
                bag_type_try06967 = EID{C_set.Id(),0}
                } else {
                bag_type_try06967 = EID{C_list.Id(),0}
                /* If-7 */} 
              }
              /* ERROR PROTECTION INSERTED (bag_type-Result) */
              if ErrorIn(bag_type_try06967) {Result = bag_type_try06967
              } else {
              bag_type = ToClass(OBJ(bag_type_try06967))
              /* Let:7 */{ 
                var try_count int  = 0
                /* noccur = 5 */
                
                F_Generate_new_block_string(MakeString("Iteration"))
                F_Generate_var_declaration_string(vlist,bag_type,1)
                if (bag_type.Id() == C_list.Id()) /* If:8 */{ 
                  F_Generate_var_declaration_string(v2,v2_range,1)
                  /* If-8 */} 
                F_Generate_var_declaration_string(vlocal,C_any,1)
                var g0701I *ClaireBoolean  
                var g0701I_try07028 EID 
                g0701I_try07028 = Optimize.F_Compile_g_throw_any(self.SetArg)
                /* ERROR PROTECTION INSERTED (g0701I-Result) */
                if ErrorIn(g0701I_try07028) {Result = g0701I_try07028
                } else {
                g0701I = ToBoolean(OBJ(g0701I_try07028))
                if (g0701I == CTRUE) /* If:8 */{ 
                  try_count = (try_count+1)
                  Result = Core.F_CALL(C_Generate_g_try,ARGS(self.SetArg.ToEID(),
                    EID{(vlist).Id(),0},
                    EID{bag_type.Id(),0},
                    EID{(v).Id(),0},
                    loop.ToEID()))
                  } else {
                  Result = F_Generate_statement_any(self.SetArg,bag_type,vlist,loop)
                  /* If-8 */} 
                }
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                F_princ_string(v)
                PRINC(" = ")
                Result = F_Generate_cast_prefix_class(C_list,s)
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                PRINC("CreateList(")
                if (Core.F_get_property(C_of,ToObject(self.Id())) != CNULL) /* If:8 */{ 
                  /* Let:9 */{ 
                    var g0703UU *ClaireAny  
                    /* noccur = 1 */
                    var g0703UU_try070410 EID 
                    g0703UU_try070410 = Core.F_CALL(Optimize.C_c_code,ARGS(Core.F_CALL(C_of,ARGS(EID{self.Id(),0})),EID{C_type.Id(),0}))
                    /* ERROR PROTECTION INSERTED (g0703UU-Result) */
                    if ErrorIn(g0703UU_try070410) {Result = g0703UU_try070410
                    } else {
                    g0703UU = ANY(g0703UU_try070410)
                    Result = Core.F_CALL(C_Generate_g_expression,ARGS(g0703UU.ToEID(),EID{C_type.Id(),0}))
                    }
                    /* Let-9 */} 
                  } else {
                  PRINC("ToType(CEMPTY.Id())")
                  Result = EVOID
                  /* If-8 */} 
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                PRINC(",")
                F_princ_string(vlist)
                PRINC(".Length())")
                F_Generate_cast_post_class(C_list,s)
                F_Generate_breakline_void()
                PRINC("")
                Result = EVOID
                }}
                {
                if (bag_type.Id() == C_set.Id()) /* If:8 */{ 
                  PRINC("var CLcount = -1")
                  F_Generate_breakline_void()
                  PRINC("")
                  PRINC("for _,")
                  F_princ_string(v2)
                  PRINC(" := range(")
                  F_princ_string(vlist)
                  PRINC(".Values) ")
                  F_Generate_new_block_void()
                  PRINC("")
                  PRINC("CLcount++")
                  F_Generate_breakline_void()
                  PRINC("")
                  Result = EVOID
                  } else {
                  PRINC("for CLcount := 0; CLcount < ")
                  F_princ_string(vlist)
                  PRINC(".Length(); CLcount++")
                  F_Generate_new_block_void()
                  PRINC("")
                  F_princ_string(v2)
                  PRINC(" = ")
                  Result = F_Generate_cast_prefix_class(C_any,v2_range)
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  F_princ_string(vlist)
                  PRINC(".At(CLcount)")
                  F_Generate_cast_post_class(C_any,v2_range)
                  F_Generate_breakline_void()
                  PRINC("")
                  Result = EVOID
                  }
                  /* If-8 */} 
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                var g0705I *ClaireBoolean  
                var g0705I_try07068 EID 
                g0705I_try07068 = Optimize.F_Compile_g_throw_any(self.Arg)
                /* ERROR PROTECTION INSERTED (g0705I-Result) */
                if ErrorIn(g0705I_try07068) {Result = g0705I_try07068
                } else {
                g0705I = ToBoolean(OBJ(g0705I_try07068))
                if (g0705I == CTRUE) /* If:8 */{ 
                  try_count = (try_count+1)
                  Result = Core.F_CALL(C_Generate_g_try,ARGS(self.Arg.ToEID(),
                    EID{(vlocal).Id(),0},
                    EID{C_any.Id(),0},
                    EID{(v).Id(),0},
                    EID{MakeTuple((v).Id(),s.Id()).Id(),0}))
                  } else {
                  Result = F_Generate_statement_any(self.Arg,C_any,vlocal,MakeTuple((v).Id(),s.Id()).Id())
                  /* If-8 */} 
                }
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                Result = F_Generate_cast_prefix_class(s,C_list)
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                F_princ_string(v)
                F_Generate_cast_post_class(s,C_list)
                PRINC(".PutAt(CLcount,")
                F_princ_string(vlocal)
                PRINC(")")
                F_Generate_breakline_void()
                PRINC("")
                Result = EVOID
                }
                {
                F_Generate_close_block_void()
                F_Generate_close_try_integer(try_count)
                F_Generate_close_block_string(MakeString("Iteration"))
                Result = EVOID
                }}}}}
                /* Let-7 */} 
              }
              /* Let-6 */} 
            /* Let-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      }
      /* Let-2 */} 
    }
    return Result} 
  
// The EID go function for: g_statement @ Iteration (throw: true) 
func E_Generate_g_statement_Iteration (self EID,s EID,v EID,err EID,loop EID) EID { 
    return /*(sm for g_statement @ Iteration= EID)*/ F_Generate_g_statement_Iteration(Language.To_Iteration(OBJ(self)),
      ToClass(OBJ(s)),
      ToString(OBJ(v)),
      ToBoolean(OBJ(err)),
      ANY(loop) )} 
  
// --------------- WHILE   ------------------------------------------
// if it is possible the logical compiler is used to produce a better code
// self.other = true => until(....) was used 
// error is more tricky => we produce a chain with 3 more blocks
/* {1} OPT.The go function for: g_statement(self:While,s:class,v:string,err:boolean,loop:any) [] */
func F_Generate_g_statement_While (self *Language.While ,s *ClaireClass ,v *ClaireString ,err *ClaireBoolean ,loop *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var f_ask *ClaireBoolean  
      /* noccur = 3 */
      var f_ask_try07073 EID 
      /* and:3 */{ 
        var v_and3 *ClaireBoolean  
        
        var v_and3_try07084 EID 
        v_and3_try07084 = F_Generate_g_clean_any(self.Test)
        /* ERROR PROTECTION INSERTED (v_and3-f_ask_try07073) */
        if ErrorIn(v_and3_try07084) {f_ask_try07073 = v_and3_try07084
        } else {
        v_and3 = ToBoolean(OBJ(v_and3_try07084))
        if (v_and3 == CFALSE) {f_ask_try07073 = EID{CFALSE.Id(),0}
        } else /* arg:4 */{ 
          v_and3 = self.Other.Not
          if (v_and3 == CFALSE) {f_ask_try07073 = EID{CFALSE.Id(),0}
          } else /* arg:5 */{ 
            f_ask_try07073 = EID{CTRUE.Id(),0}/* arg-5 */} 
          /* arg-4 */} 
        }
        /* and-3 */} 
      /* ERROR PROTECTION INSERTED (f_ask-Result) */
      if ErrorIn(f_ask_try07073) {Result = f_ask_try07073
      } else {
      f_ask = ToBoolean(OBJ(f_ask_try07073))
      /* Let:3 */{ 
        var try_ask *ClaireBoolean  
        /* noccur = 3 */
        var try_ask_try07094 EID 
        /* and:4 */{ 
          var v_and4 *ClaireBoolean  
          
          v_and4 = self.Other.Not
          if (v_and4 == CFALSE) {try_ask_try07094 = EID{CFALSE.Id(),0}
          } else /* arg:5 */{ 
            var v_and4_try07106 EID 
            v_and4_try07106 = Optimize.F_Compile_g_throw_any(self.Test)
            /* ERROR PROTECTION INSERTED (v_and4-try_ask_try07094) */
            if ErrorIn(v_and4_try07106) {try_ask_try07094 = v_and4_try07106
            } else {
            v_and4 = ToBoolean(OBJ(v_and4_try07106))
            if (v_and4 == CFALSE) {try_ask_try07094 = EID{CFALSE.Id(),0}
            } else /* arg:6 */{ 
              try_ask_try07094 = EID{CTRUE.Id(),0}/* arg-6 */} 
            /* arg-5 */} 
          }
          /* and-4 */} 
        /* ERROR PROTECTION INSERTED (try_ask-Result) */
        if ErrorIn(try_ask_try07094) {Result = try_ask_try07094
        } else {
        try_ask = ToBoolean(OBJ(try_ask_try07094))
        /* Let:4 */{ 
          var try2_ask *ClaireBoolean  
          /* noccur = 2 */
          var try2_ask_try07115 EID 
          try2_ask_try07115 = Optimize.F_Compile_g_throw_any(self.Arg)
          /* ERROR PROTECTION INSERTED (try2_ask-Result) */
          if ErrorIn(try2_ask_try07115) {Result = try2_ask_try07115
          } else {
          try2_ask = ToBoolean(OBJ(try2_ask_try07115))
          /* Let:5 */{ 
            var v2 *ClaireString   = F_Generate_check_var_string(MakeString("v_while"))
            /* noccur = 6 */
            if (f_ask != CTRUE) /* If:6 */{ 
              F_Generate_var_declaration_string(v2,C_boolean,1)
              /* If-6 */} 
            if (s.Id() != C_void.Id()) /* If:6 */{ 
              F_c_princ_string(v)
              PRINC("= ")
              Result = F_Generate_cast_prefix_class(C_boolean,s)
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              PRINC("CFALSE")
              F_Generate_cast_post_class(C_boolean,s)
              F_Generate_breakline_void()
              PRINC("")
              Result = EVOID
              }
              } else {
              Result = EID{CFALSE.Id(),0}
              /* If-6 */} 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            if (f_ask == CTRUE) /* If:6 */{ 
              PRINC("for ")
              Result = Core.F_CALL(Optimize.C_Compile_bool_exp,ARGS(self.Test.ToEID(),EID{self.Other.Not.Id(),0}))
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              PRINC(" ")
              Result = EVOID
              }
              } else {
              if (try_ask == CTRUE) /* If:7 */{ 
                Result = Core.F_CALL(C_Generate_g_try,ARGS(self.Test.ToEID(),
                  EID{(v2).Id(),0},
                  EID{C_boolean.Id(),0},
                  EID{(v).Id(),0},
                  EID{CFALSE.Id(),0}))
                } else {
                /* Let:8 */{ 
                  var g0712UU *ClaireAny  
                  /* noccur = 1 */
                  if (self.Other == CTRUE) /* If:9 */{ 
                    g0712UU = CFALSE.Id()
                    } else {
                    g0712UU = self.Test
                    /* If-9 */} 
                  Result = F_Generate_statement_any(g0712UU,C_boolean,v2,CFALSE.Id())
                  /* Let-8 */} 
                /* If-7 */} 
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              F_Generate_breakline_void()
              PRINC("for ")
              F_princ_string(v2)
              PRINC(" ")
              if (self.Other == CTRUE) /* If:7 */{ 
                PRINC("!=")
                } else {
                PRINC("==")
                /* If-7 */} 
              PRINC(" CTRUE ")
              Result = EVOID
              }
              /* If-6 */} 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            F_Generate_new_block_string(MakeString("while"))
            if (try2_ask == CTRUE) /* If:6 */{ 
              Result = F_Generate_g_try_void_any(self.Arg,v,MakeTuple((v).Id(),s.Id()).Id())
              } else {
              Result = F_Generate_statement_any(self.Arg,C_void,v,MakeTuple((v).Id(),s.Id()).Id())
              /* If-6 */} 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            if (try_ask == CTRUE) /* If:6 */{ 
              Result = Core.F_CALL(C_Generate_g_try,ARGS(self.Test.ToEID(),
                EID{(v2).Id(),0},
                EID{C_boolean.Id(),0},
                EID{(v).Id(),0},
                EID{MakeTuple((v).Id(),s.Id()).Id(),0}))
              /* If!6 */}  else if (f_ask != CTRUE) /* If:6 */{ 
              Result = F_Generate_statement_any(self.Test,C_boolean,v2,CFALSE.Id())
              } else {
              Result = EID{CFALSE.Id(),0}
              /* If-6 */} 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            F_Generate_close_block_string(MakeString("while"))
            if (try_ask == CTRUE) /* If:6 */{ 
              F_Generate_close_try_integer(2)
              /* If-6 */} 
            if (try2_ask == CTRUE) /* If:6 */{ 
              F_Generate_close_try_integer(1)
              Result = EVOID
              } else {
              Result = EID{CFALSE.Id(),0}
              /* If-6 */} 
            }}}}
            /* Let-5 */} 
          }
          /* Let-4 */} 
        }
        /* Let-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: g_statement @ While (throw: true) 
func E_Generate_g_statement_While (self EID,s EID,v EID,err EID,loop EID) EID { 
    return /*(sm for g_statement @ While= EID)*/ F_Generate_g_statement_While(Language.To_While(OBJ(self)),
      ToClass(OBJ(s)),
      ToString(OBJ(v)),
      ToBoolean(OBJ(err)),
      ANY(loop) )} 
  
//------------- compiling a return -------------------------------------
// a return inside a loop is compiled with a break, the go variable is provided
// in the loop argument
/* {1} OPT.The go function for: g_statement(self:Return,s:class,v:string,err:boolean,loop:any) [] */
func F_Generate_g_statement_Return (self *Language.Return ,s *ClaireClass ,v *ClaireString ,err *ClaireBoolean ,loop *ClaireAny ) EID { 
    var Result EID 
    if (C_tuple.Id() == loop.Isa.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var g0713 *ClaireTuple   = ToTuple(loop)
        /* noccur = 2 */
        /* Let:4 */{ 
          var vreturn *ClaireString   = ToString(ToList(g0713.Id()).At(1-1))
          /* noccur = 2 */
          /* Let:5 */{ 
            var sreturn *ClaireClass   = ToClass(ToList(g0713.Id()).At(2-1))
            /* noccur = 2 */
            PRINC(" /*v = ")
            F_princ_string(vreturn)
            PRINC(", s =")
            Result = Core.F_print_any(sreturn.Id())
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC("*/\n")
            Result = EVOID
            }
            {
            Result = F_Generate_statement_any(self.Arg,sreturn,vreturn,CFALSE.Id())
            }
            /* Let-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      } else {
      Result = EID{CFALSE.Id(),0}
      /* If-2 */} 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("break")
    Result = F_Generate_breakline_void().ToEID()
    }
    return Result} 
  
// The EID go function for: g_statement @ Return (throw: true) 
func E_Generate_g_statement_Return (self EID,s EID,v EID,err EID,loop EID) EID { 
    return /*(sm for g_statement @ Return= EID)*/ F_Generate_g_statement_Return(Language.To_Return(OBJ(self)),
      ToClass(OBJ(s)),
      ToString(OBJ(v)),
      ToBoolean(OBJ(err)),
      ANY(loop) )} 
  
//**********************************************************************
//*          Part 4: CLAIRE-specific structures                        *
//**********************************************************************
// ------------- Messages and optimized instructions ------------------------
// this is one example on how to unfold: a Call
// note that if the error is returned it should be passed away
// we also add inline_stat in v4 for special cases
/* {1} OPT.The go function for: g_statement(self:Call,s:class,v:string,err:boolean,loop:any) [] */
func F_Generate_g_statement_Call (self *Language.Call ,s *ClaireClass ,v *ClaireString ,err *ClaireBoolean ,loop *ClaireAny ) EID { 
    var Result EID 
    var g0714I *ClaireBoolean  
    var g0714I_try07152 EID 
    g0714I_try07152 = F_Generate_g_clean_any(self.Args.Id())
    /* ERROR PROTECTION INSERTED (g0714I-Result) */
    if ErrorIn(g0714I_try07152) {Result = g0714I_try07152
    } else {
    g0714I = ToBoolean(OBJ(g0714I_try07152))
    if (g0714I == CTRUE) /* If:2 */{ 
      Result = F_Generate_inline_stat_Call(self,s,v)
      } else {
      /* Let:3 */{ 
        var l *ClaireList   = self.Args
        /* noccur = 3 */
        /* Let:4 */{ 
          var ld *ClaireList  
          /* noccur = 3 */
          var ld_try07175 EID 
          ld_try07175 = F_Generate_unfold_args_list(l)
          /* ERROR PROTECTION INSERTED (ld-Result) */
          if ErrorIn(ld_try07175) {Result = ld_try07175
          } else {
          ld = ToList(OBJ(ld_try07175))
          if (Equal(ld.Id(),CNIL.Id()) == CTRUE) /* If:5 */{ 
            Result = F_Generate_error_wrap_any(self.Id(),s,v)
            } else {
            /* Let:6 */{ 
              var g0718UU *Language.Call  
              /* noccur = 1 */
              var g0718UU_try07197 EID 
              /* Let:7 */{ 
                var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                /* noccur = 5 */
                _CL_obj.Selector = self.Selector
                /* update:8 */{ 
                  var va_arg1 *Language.Call  
                  var va_arg2 *ClaireList  
                  va_arg1 = _CL_obj
                  var va_arg2_try07209 EID 
                  /* Iteration:9 */{ 
                    var v_list9 *ClaireList  
                    var z *ClaireAny  
                    var v_local9 *ClaireAny  
                    v_list9 = l
                    va_arg2_try07209 = EID{CreateList(ToType(CEMPTY.Id()),v_list9.Length()).Id(),0}
                    for CLcount := 0; CLcount < v_list9.Length(); CLcount++{ 
                      z = v_list9.At(CLcount)
                      var v_local9_try072111 EID 
                      v_local9_try072111 = F_Generate_unfold_arg_list(l,ld,z)
                      /* ERROR PROTECTION INSERTED (v_local9-va_arg2_try07209) */
                      if ErrorIn(v_local9_try072111) {va_arg2_try07209 = v_local9_try072111
                      va_arg2_try07209 = v_local9_try072111
                      break
                      } else {
                      v_local9 = ANY(v_local9_try072111)
                      ToList(OBJ(va_arg2_try07209)).PutAt(CLcount,v_local9)
                      } 
                    }
                    /* Iteration-9 */} 
                  /* ERROR PROTECTION INSERTED (va_arg2-g0718UU_try07197) */
                  if ErrorIn(va_arg2_try07209) {g0718UU_try07197 = va_arg2_try07209
                  } else {
                  va_arg2 = ToList(OBJ(va_arg2_try07209))
                  /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                  va_arg1.Args = va_arg2
                  g0718UU_try07197 = EID{va_arg2.Id(),0}
                  }
                  /* update-8 */} 
                /* ERROR PROTECTION INSERTED (g0718UU_try07197-g0718UU_try07197) */
                if !ErrorIn(g0718UU_try07197) {
                g0718UU_try07197 = EID{_CL_obj.Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (g0718UU-Result) */
              if ErrorIn(g0718UU_try07197) {Result = g0718UU_try07197
              } else {
              g0718UU = Language.To_Call(OBJ(g0718UU_try07197))
              Result = F_Generate_unfold_use_list(ld,
                g0718UU.Id(),
                s,
                v,
                err,
                loop)
              }
              /* Let-6 */} 
            /* If-5 */} 
          }
          /* Let-4 */} 
        /* Let-3 */} 
      /* If-2 */} 
    }
    return Result} 
  
// The EID go function for: g_statement @ Call (throw: true) 
func E_Generate_g_statement_Call (self EID,s EID,v EID,err EID,loop EID) EID { 
    return /*(sm for g_statement @ Call= EID)*/ F_Generate_g_statement_Call(Language.To_Call(OBJ(self)),
      ToClass(OBJ(s)),
      ToString(OBJ(v)),
      ToBoolean(OBJ(err)),
      ANY(loop) )} 
  
// this is our special inling that requires an assignment (not allowed as an expression in go)
/* {1} OPT.The go function for: inline_stat(self:Call,s:class,v:string) [] */
func F_Generate_inline_stat_Call (self *Language.Call ,s *ClaireClass ,v *ClaireString ) EID { 
    var Result EID 
    if (self.Selector.Id() == Optimize.C_Compile_object_I.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var a1 *ClaireAny   = self.Args.At(1-1)
        /* noccur = 9 */
        /* Let:4 */{ 
          var a2 *ClaireAny   = self.Args.At(2-1)
          /* noccur = 5 */
          if ((a2 == C_property.Id()) && 
              (Core.F_owner_any(ANY(Core.F_CALL(C_value,ARGS(a1.ToEID())))).IsIn(C_property) == CTRUE)) /* If:5 */{ 
            F_Generate_symbol_ident_symbol(ToSymbol(a1))
            PRINC(" = ")
            Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).Declare(ToProperty(OBJ(Core.F_CALL(C_value,ARGS(a1.ToEID())))))
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            F_Generate_breakline_void()
            PRINC("")
            Result = EVOID
            }
            } else {
            PRINC("/* object!:")
            Result = Core.F_CALL(C_print,ARGS(a1.ToEID()))
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC(" ->")
            Result = Core.F_print_any(ToSymbol(a1).Defined().Id())
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC("*/")
            Result = EVOID
            }}
            {
            F_Generate_symbol_ident_symbol(ToSymbol(a1))
            PRINC(" = ")
            F_Generate_object_prefix_class(C_any,ToClass(a2))
            PRINC("new(")
            F_Generate_go_class_class(ToClass(a2))
            PRINC(").IsNamed(")
            F_Generate_class_ident_class(ToClass(a2))
            PRINC(",MakeSymbol(")
            /* Let:6 */{ 
              var g0723UU *ClaireAny  
              /* noccur = 1 */
              var g0723UU_try07247 EID 
              g0723UU_try07247 = Core.F_CALL(C_string_I,ARGS(a1.ToEID()))
              /* ERROR PROTECTION INSERTED (g0723UU-Result) */
              if ErrorIn(g0723UU_try07247) {Result = g0723UU_try07247
              } else {
              g0723UU = ANY(g0723UU_try07247)
              Result = Core.F_print_any(g0723UU)
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC(",")
            Result = F_Generate_g_expression_module(ToModule(OBJ(Core.F_CALL(C_module_I,ARGS(a1.ToEID())))),C_module)
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC("))")
            F_Generate_object_post_class(C_any,ToClass(a2))
            F_Generate_breakline_void()
            PRINC("")
            Result = EVOID
            }}
            }
            /* If-5 */} 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          if (s.Id() != C_void.Id()) /* If:5 */{ 
            F_Generate_breakline_void()
            F_c_princ_string(v)
            PRINC(" = ")
            F_Generate_symbol_ident_symbol(ToSymbol(a1))
            F_Generate_breakline_void()
            PRINC("")
            Result = EVOID
            } else {
            Result = EID{CFALSE.Id(),0}
            /* If-5 */} 
          }
          /* Let-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (s.Id() == Optimize.C_EID.Id()) /* If:2 */{ 
      F_c_princ_string(v)
      PRINC(" = ")
      Result = F_Generate_g_expression_Call(self,s)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      F_Generate_breakline_void()
      PRINC("")
      Result = EVOID
      }
      } else {
      Result = ToException(Core.C_general_error.Make(MakeString("desing error : inline_stat for ~S").Id(),MakeConstantList(self.Id()).Id())).Close()
      /* If-2 */} 
    return Result} 
  
// The EID go function for: inline_stat @ Call (throw: true) 
func E_Generate_inline_stat_Call (self EID,s EID,v EID) EID { 
    return /*(sm for inline_stat @ Call= EID)*/ F_Generate_inline_stat_Call(Language.To_Call(OBJ(self)),ToClass(OBJ(s)),ToString(OBJ(v)) )} 
  
// A call method is now simpler with unfolding ! very similar structucture
/* {1} OPT.The go function for: g_statement(self:Call_method,s:class,v:string,err:boolean,loop:any) [] */
func F_Generate_g_statement_Call_method (self *Language.CallMethod ,s *ClaireClass ,v *ClaireString ,err *ClaireBoolean ,loop *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var l *ClaireList   = self.Args
      /* noccur = 3 */
      /* Let:3 */{ 
        var ld *ClaireList  
        /* noccur = 3 */
        var ld_try07254 EID 
        ld_try07254 = F_Generate_unfold_args_list(l)
        /* ERROR PROTECTION INSERTED (ld-Result) */
        if ErrorIn(ld_try07254) {Result = ld_try07254
        } else {
        ld = ToList(OBJ(ld_try07254))
        if (Equal(ld.Id(),CNIL.Id()) == CTRUE) /* If:4 */{ 
          Result = F_Generate_error_wrap_any(self.Id(),s,v)
          } else {
          /* Let:5 */{ 
            var g0726UU *Language.CallMethod  
            /* noccur = 1 */
            var g0726UU_try07276 EID 
            /* Let:6 */{ 
              var _CL_obj *Language.CallMethod   = Language.To_CallMethod(new(Language.CallMethod).Is(Language.C_Call_method))
              /* noccur = 5 */
              _CL_obj.Arg = self.Arg
              /* update:7 */{ 
                var va_arg1 *Language.CallMethod  
                var va_arg2 *ClaireList  
                va_arg1 = _CL_obj
                var va_arg2_try07288 EID 
                /* Iteration:8 */{ 
                  var v_list8 *ClaireList  
                  var z *ClaireAny  
                  var v_local8 *ClaireAny  
                  v_list8 = l
                  va_arg2_try07288 = EID{CreateList(ToType(CEMPTY.Id()),v_list8.Length()).Id(),0}
                  for CLcount := 0; CLcount < v_list8.Length(); CLcount++{ 
                    z = v_list8.At(CLcount)
                    var v_local8_try072910 EID 
                    v_local8_try072910 = F_Generate_unfold_arg_list(l,ld,z)
                    /* ERROR PROTECTION INSERTED (v_local8-va_arg2_try07288) */
                    if ErrorIn(v_local8_try072910) {va_arg2_try07288 = v_local8_try072910
                    va_arg2_try07288 = v_local8_try072910
                    break
                    } else {
                    v_local8 = ANY(v_local8_try072910)
                    ToList(OBJ(va_arg2_try07288)).PutAt(CLcount,v_local8)
                    } 
                  }
                  /* Iteration-8 */} 
                /* ERROR PROTECTION INSERTED (va_arg2-g0726UU_try07276) */
                if ErrorIn(va_arg2_try07288) {g0726UU_try07276 = va_arg2_try07288
                } else {
                va_arg2 = ToList(OBJ(va_arg2_try07288))
                /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                va_arg1.Args = va_arg2
                g0726UU_try07276 = EID{va_arg2.Id(),0}
                }
                /* update-7 */} 
              /* ERROR PROTECTION INSERTED (g0726UU_try07276-g0726UU_try07276) */
              if !ErrorIn(g0726UU_try07276) {
              g0726UU_try07276 = EID{_CL_obj.Id(),0}
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (g0726UU-Result) */
            if ErrorIn(g0726UU_try07276) {Result = g0726UU_try07276
            } else {
            g0726UU = Language.To_CallMethod(OBJ(g0726UU_try07276))
            Result = F_Generate_unfold_use_list(ld,
              g0726UU.Id(),
              s,
              v,
              err,
              loop)
            }
            /* Let-5 */} 
          /* If-4 */} 
        }
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: g_statement @ Call_method (throw: true) 
func E_Generate_g_statement_Call_method (self EID,s EID,v EID,err EID,loop EID) EID { 
    return /*(sm for g_statement @ Call_method= EID)*/ F_Generate_g_statement_Call_method(Language.To_CallMethod(OBJ(self)),
      ToClass(OBJ(s)),
      ToString(OBJ(v)),
      ToBoolean(OBJ(err)),
      ANY(loop) )} 
  
/* {1} OPT.The go function for: g_statement(self:Call_method1,s:class,v:string,err:boolean,loop:any) [] */
func F_Generate_g_statement_Call_method1 (self *Language.CallMethod1 ,s *ClaireClass ,v *ClaireString ,err *ClaireBoolean ,loop *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var l *ClaireList   = self.Args
      /* noccur = 3 */
      /* Let:3 */{ 
        var ld *ClaireList  
        /* noccur = 3 */
        var ld_try07304 EID 
        ld_try07304 = F_Generate_unfold_args_list(l)
        /* ERROR PROTECTION INSERTED (ld-Result) */
        if ErrorIn(ld_try07304) {Result = ld_try07304
        } else {
        ld = ToList(OBJ(ld_try07304))
        
        if (Equal(ld.Id(),CNIL.Id()) == CTRUE) /* If:4 */{ 
          Result = F_Generate_error_wrap_any(self.Id(),s,v)
          } else {
          /* Let:5 */{ 
            var g0731UU *Language.CallMethod1  
            /* noccur = 1 */
            var g0731UU_try07326 EID 
            /* Let:6 */{ 
              var _CL_obj *Language.CallMethod1   = Language.To_CallMethod1(new(Language.CallMethod1).Is(Language.C_Call_method1))
              /* noccur = 5 */
              _CL_obj.Arg = self.Arg
              /* update:7 */{ 
                var va_arg1 *Language.CallMethod  
                var va_arg2 *ClaireList  
                va_arg1 = Language.To_CallMethod(_CL_obj.Id())
                var va_arg2_try07338 EID 
                /* Iteration:8 */{ 
                  var v_list8 *ClaireList  
                  var z *ClaireAny  
                  var v_local8 *ClaireAny  
                  v_list8 = l
                  va_arg2_try07338 = EID{CreateList(ToType(CEMPTY.Id()),v_list8.Length()).Id(),0}
                  for CLcount := 0; CLcount < v_list8.Length(); CLcount++{ 
                    z = v_list8.At(CLcount)
                    var v_local8_try073410 EID 
                    v_local8_try073410 = F_Generate_unfold_arg_list(l,ld,z)
                    /* ERROR PROTECTION INSERTED (v_local8-va_arg2_try07338) */
                    if ErrorIn(v_local8_try073410) {va_arg2_try07338 = v_local8_try073410
                    va_arg2_try07338 = v_local8_try073410
                    break
                    } else {
                    v_local8 = ANY(v_local8_try073410)
                    ToList(OBJ(va_arg2_try07338)).PutAt(CLcount,v_local8)
                    } 
                  }
                  /* Iteration-8 */} 
                /* ERROR PROTECTION INSERTED (va_arg2-g0731UU_try07326) */
                if ErrorIn(va_arg2_try07338) {g0731UU_try07326 = va_arg2_try07338
                } else {
                va_arg2 = ToList(OBJ(va_arg2_try07338))
                /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                va_arg1.Args = va_arg2
                g0731UU_try07326 = EID{va_arg2.Id(),0}
                }
                /* update-7 */} 
              /* ERROR PROTECTION INSERTED (g0731UU_try07326-g0731UU_try07326) */
              if !ErrorIn(g0731UU_try07326) {
              g0731UU_try07326 = EID{_CL_obj.Id(),0}
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (g0731UU-Result) */
            if ErrorIn(g0731UU_try07326) {Result = g0731UU_try07326
            } else {
            g0731UU = Language.To_CallMethod1(OBJ(g0731UU_try07326))
            Result = F_Generate_unfold_use_list(ld,
              g0731UU.Id(),
              s,
              v,
              err,
              loop)
            }
            /* Let-5 */} 
          /* If-4 */} 
        }
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: g_statement @ Call_method1 (throw: true) 
func E_Generate_g_statement_Call_method1 (self EID,s EID,v EID,err EID,loop EID) EID { 
    return /*(sm for g_statement @ Call_method1= EID)*/ F_Generate_g_statement_Call_method1(Language.To_CallMethod1(OBJ(self)),
      ToClass(OBJ(s)),
      ToString(OBJ(v)),
      ToBoolean(OBJ(err)),
      ANY(loop) )} 
  
/* {1} OPT.The go function for: g_statement(self:Call_method2,s:class,v:string,err:boolean,loop:any) [] */
func F_Generate_g_statement_Call_method2 (self *Language.CallMethod2 ,s *ClaireClass ,v *ClaireString ,err *ClaireBoolean ,loop *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var l *ClaireList   = self.Args
      /* noccur = 3 */
      /* Let:3 */{ 
        var ld *ClaireList  
        /* noccur = 3 */
        var ld_try07354 EID 
        ld_try07354 = F_Generate_unfold_args_list(l)
        /* ERROR PROTECTION INSERTED (ld-Result) */
        if ErrorIn(ld_try07354) {Result = ld_try07354
        } else {
        ld = ToList(OBJ(ld_try07354))
        if (Equal(ld.Id(),CNIL.Id()) == CTRUE) /* If:4 */{ 
          Result = F_Generate_error_wrap_any(self.Id(),s,v)
          } else {
          /* Let:5 */{ 
            var g0736UU *Language.CallMethod2  
            /* noccur = 1 */
            var g0736UU_try07376 EID 
            /* Let:6 */{ 
              var _CL_obj *Language.CallMethod2   = Language.To_CallMethod2(new(Language.CallMethod2).Is(Language.C_Call_method2))
              /* noccur = 5 */
              _CL_obj.Arg = self.Arg
              /* update:7 */{ 
                var va_arg1 *Language.CallMethod  
                var va_arg2 *ClaireList  
                va_arg1 = Language.To_CallMethod(_CL_obj.Id())
                var va_arg2_try07388 EID 
                /* Iteration:8 */{ 
                  var v_list8 *ClaireList  
                  var z *ClaireAny  
                  var v_local8 *ClaireAny  
                  v_list8 = l
                  va_arg2_try07388 = EID{CreateList(ToType(CEMPTY.Id()),v_list8.Length()).Id(),0}
                  for CLcount := 0; CLcount < v_list8.Length(); CLcount++{ 
                    z = v_list8.At(CLcount)
                    var v_local8_try073910 EID 
                    v_local8_try073910 = F_Generate_unfold_arg_list(l,ld,z)
                    /* ERROR PROTECTION INSERTED (v_local8-va_arg2_try07388) */
                    if ErrorIn(v_local8_try073910) {va_arg2_try07388 = v_local8_try073910
                    va_arg2_try07388 = v_local8_try073910
                    break
                    } else {
                    v_local8 = ANY(v_local8_try073910)
                    ToList(OBJ(va_arg2_try07388)).PutAt(CLcount,v_local8)
                    } 
                  }
                  /* Iteration-8 */} 
                /* ERROR PROTECTION INSERTED (va_arg2-g0736UU_try07376) */
                if ErrorIn(va_arg2_try07388) {g0736UU_try07376 = va_arg2_try07388
                } else {
                va_arg2 = ToList(OBJ(va_arg2_try07388))
                /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                va_arg1.Args = va_arg2
                g0736UU_try07376 = EID{va_arg2.Id(),0}
                }
                /* update-7 */} 
              /* ERROR PROTECTION INSERTED (g0736UU_try07376-g0736UU_try07376) */
              if !ErrorIn(g0736UU_try07376) {
              g0736UU_try07376 = EID{_CL_obj.Id(),0}
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (g0736UU-Result) */
            if ErrorIn(g0736UU_try07376) {Result = g0736UU_try07376
            } else {
            g0736UU = Language.To_CallMethod2(OBJ(g0736UU_try07376))
            Result = F_Generate_unfold_use_list(ld,
              g0736UU.Id(),
              s,
              v,
              err,
              loop)
            }
            /* Let-5 */} 
          /* If-4 */} 
        }
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: g_statement @ Call_method2 (throw: true) 
func E_Generate_g_statement_Call_method2 (self EID,s EID,v EID,err EID,loop EID) EID { 
    return /*(sm for g_statement @ Call_method2= EID)*/ F_Generate_g_statement_Call_method2(Language.To_CallMethod2(OBJ(self)),
      ToClass(OBJ(s)),
      ToString(OBJ(v)),
      ToBoolean(OBJ(err)),
      ANY(loop) )} 
  
// in Claire 4, Super is handled as a Call
/* {1} OPT.The go function for: g_statement(self:Super,s:class,v:string,err:boolean,loop:any) [] */
func F_Generate_g_statement_Super (self *Language.Super ,s *ClaireClass ,v *ClaireString ,err *ClaireBoolean ,loop *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var l *ClaireList   = self.Args
      /* noccur = 3 */
      /* Let:3 */{ 
        var ld *ClaireList  
        /* noccur = 3 */
        var ld_try07404 EID 
        ld_try07404 = F_Generate_unfold_args_list(l)
        /* ERROR PROTECTION INSERTED (ld-Result) */
        if ErrorIn(ld_try07404) {Result = ld_try07404
        } else {
        ld = ToList(OBJ(ld_try07404))
        if (Equal(ld.Id(),CNIL.Id()) == CTRUE) /* If:4 */{ 
          Result = F_Generate_error_wrap_any(self.Id(),s,v)
          } else {
          /* Let:5 */{ 
            var g0741UU *Language.Super  
            /* noccur = 1 */
            var g0741UU_try07426 EID 
            /* Let:6 */{ 
              var _CL_obj *Language.Super   = Language.To_Super(new(Language.Super).Is(Language.C_Super))
              /* noccur = 7 */
              _CL_obj.Selector = self.Selector
              _CL_obj.CastTo = self.CastTo
              /* update:7 */{ 
                var va_arg1 *Language.Super  
                var va_arg2 *ClaireList  
                va_arg1 = _CL_obj
                var va_arg2_try07438 EID 
                /* Iteration:8 */{ 
                  var v_list8 *ClaireList  
                  var z *ClaireAny  
                  var v_local8 *ClaireAny  
                  v_list8 = l
                  va_arg2_try07438 = EID{CreateList(ToType(CEMPTY.Id()),v_list8.Length()).Id(),0}
                  for CLcount := 0; CLcount < v_list8.Length(); CLcount++{ 
                    z = v_list8.At(CLcount)
                    var v_local8_try074410 EID 
                    v_local8_try074410 = F_Generate_unfold_arg_list(l,ld,z)
                    /* ERROR PROTECTION INSERTED (v_local8-va_arg2_try07438) */
                    if ErrorIn(v_local8_try074410) {va_arg2_try07438 = v_local8_try074410
                    va_arg2_try07438 = v_local8_try074410
                    break
                    } else {
                    v_local8 = ANY(v_local8_try074410)
                    ToList(OBJ(va_arg2_try07438)).PutAt(CLcount,v_local8)
                    } 
                  }
                  /* Iteration-8 */} 
                /* ERROR PROTECTION INSERTED (va_arg2-g0741UU_try07426) */
                if ErrorIn(va_arg2_try07438) {g0741UU_try07426 = va_arg2_try07438
                } else {
                va_arg2 = ToList(OBJ(va_arg2_try07438))
                /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                va_arg1.Args = va_arg2
                g0741UU_try07426 = EID{va_arg2.Id(),0}
                }
                /* update-7 */} 
              /* ERROR PROTECTION INSERTED (g0741UU_try07426-g0741UU_try07426) */
              if !ErrorIn(g0741UU_try07426) {
              g0741UU_try07426 = EID{_CL_obj.Id(),0}
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (g0741UU-Result) */
            if ErrorIn(g0741UU_try07426) {Result = g0741UU_try07426
            } else {
            g0741UU = Language.To_Super(OBJ(g0741UU_try07426))
            Result = F_Generate_unfold_use_list(ld,
              g0741UU.Id(),
              s,
              v,
              err,
              loop)
            }
            /* Let-5 */} 
          /* If-4 */} 
        }
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: g_statement @ Super (throw: true) 
func E_Generate_g_statement_Super (self EID,s EID,v EID,err EID,loop EID) EID { 
    return /*(sm for g_statement @ Super= EID)*/ F_Generate_g_statement_Super(Language.To_Super(OBJ(self)),
      ToClass(OBJ(s)),
      ToString(OBJ(v)),
      ToBoolean(OBJ(err)),
      ANY(loop) )} 
  
// trivial 
/* {1} OPT.The go function for: g_statement(self:Cast,s:class,v:string,err:boolean,loop:any) [] */
func F_Generate_g_statement_Cast (self *Language.Cast ,s *ClaireClass ,v *ClaireString ,err *ClaireBoolean ,loop *ClaireAny ) EID { 
    var Result EID 
    Result = Core.F_CALL(C_Generate_statement,ARGS(self.Arg.ToEID(),
      EID{s.Id(),0},
      EID{s.Id(),0},
      EID{(v).Id(),0},
      loop.ToEID()))
    return Result} 
  
// The EID go function for: g_statement @ Cast (throw: true) 
func E_Generate_g_statement_Cast (self EID,s EID,v EID,err EID,loop EID) EID { 
    return /*(sm for g_statement @ Cast= EID)*/ F_Generate_g_statement_Cast(Language.To_Cast(OBJ(self)),
      ToClass(OBJ(s)),
      ToString(OBJ(v)),
      ToBoolean(OBJ(err)),
      ANY(loop) )} 
  
//-------------- compiling a handle -------------------------------------
// In most cases, s = EID (err = true) and v is an EID variable => reuse v
// in some cases (s != EID => test = any) .. we need a special variable (v2)
// we see if the catch applied (bool : e % S) 
// in CLAIRE4, we know that self.test is a class
/* {1} OPT.The go function for: g_statement(self:Handle,s:class,v:string,err:boolean,loop:any) [] */
func F_Generate_g_statement_Handle (self *Language.ClaireHandle ,s *ClaireClass ,v *ClaireString ,err *ClaireBoolean ,loop *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var v2 *ClaireString  
      /* noccur = 6 */
      if (s.Id() == Optimize.C_EID.Id()) /* If:3 */{ 
        v2 = v
        } else {
        v2 = F_Generate_check_var_string(F_append_string(v,MakeString("_try")))
        /* If-3 */} 
      if (s.Id() != Optimize.C_EID.Id()) /* If:3 */{ 
        F_Generate_var_declaration_string(v2,Optimize.C_EID,1)
        /* If-3 */} 
      PRINC("h_index := ClEnv.Index /* Handle */")
      F_Generate_breakline_void()
      PRINC("")
      PRINC("h_base := ClEnv.Base")
      F_Generate_breakline_void()
      PRINC("")
      Result = F_Generate_statement_any(self.Arg,Optimize.C_EID,v2,CFALSE.Id())
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      if (self.Test == C_any.Id()) /* If:3 */{ 
        PRINC("if ErrorIn(")
        F_princ_string(v2)
        PRINC(")")
        Result = EVOID
        } else {
        PRINC("if ErrorIn(")
        F_princ_string(v2)
        PRINC(") && ")
        Result = Core.F_CALL(C_Generate_g_expression,ARGS(self.Test.ToEID(),EID{C_type.Id(),0}))
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC(".Contains(ANY(")
        F_princ_string(v2)
        PRINC(")) == CTRUE ")
        Result = EVOID
        }
        /* If-3 */} 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      F_Generate_new_block_void()
      PRINC("/* s=")
      Result = Core.F_print_any(s.Id())
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(" */")
      Result = EVOID
      }
      {
      PRINC("ClEnv.Index = h_index")
      F_Generate_breakline_void()
      PRINC("")
      PRINC("ClEnv.Base = h_base")
      F_Generate_breakline_void()
      PRINC("")
      Result = F_Generate_statement_any(self.Other,s,v,loop)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      if ((s.Id() == Optimize.C_EID.Id()) || 
          (s.Id() == C_void.Id())) /* If:3 */{ 
        F_Generate_close_block_void()
        Result = EVOID
        } else {
        PRINC("} else {")
        F_Generate_breakline_void()
        F_c_princ_string(v)
        PRINC(" = ")
        Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).FromEid(v2,s)
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        F_Generate_breakline_void()
        F_Generate_close_block_void()
        PRINC("")
        Result = EVOID
        }
        /* If-3 */} 
      }}}}
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: g_statement @ Handle (throw: true) 
func E_Generate_g_statement_Handle (self EID,s EID,v EID,err EID,loop EID) EID { 
    return /*(sm for g_statement @ Handle= EID)*/ F_Generate_g_statement_Handle(Language.To_ClaireHandle(OBJ(self)),
      ToClass(OBJ(s)),
      ToString(OBJ(v)),
      ToBoolean(OBJ(err)),
      ANY(loop) )} 
  
// to_CL and to_C are presently ignored in CLAIRE 4
// [g_statement(self:Generate/to_CL,s:class,v:string,err:boolean,loop:any) : void
// -> g_statement(self.arg, s, v, err, loop) ]
// [g_statement(self:Generate/to_C,s:class,v:string,err:boolean,loop:any) : void
// -> g_statement(self.arg, s, v, err, loop) ]
// same for a cast
// v3.2.06: the case where self.arg is of type any is painful => it is forbiden in osystem.cl
/* {1} OPT.The go function for: g_statement(self:Compile/C_cast,s:class,v:string,err:boolean,loop:any) [] */
func F_Generate_g_statement_C_cast (self *Optimize.CompileCCast ,s *ClaireClass ,v *ClaireString ,err *ClaireBoolean ,loop *ClaireAny ) EID { 
    var Result EID 
    Result = Core.F_CALL(C_Generate_g_statement,ARGS(self.Arg.ToEID(),
      EID{s.Id(),0},
      EID{(v).Id(),0},
      EID{err.Id(),0},
      loop.ToEID()))
    return Result} 
  
// The EID go function for: g_statement @ Compile/C_cast (throw: true) 
func E_Generate_g_statement_C_cast (self EID,s EID,v EID,err EID,loop EID) EID { 
    return /*(sm for g_statement @ Compile/C_cast= EID)*/ F_Generate_g_statement_C_cast(Optimize.To_CompileCCast(OBJ(self)),
      ToClass(OBJ(s)),
      ToString(OBJ(v)),
      ToBoolean(OBJ(err)),
      ANY(loop) )} 
  
//------------- compiling slot read/write -------------------------------
// new in CLAIRE 4 : there are two kinds => err (EID required) or not (self.arg is just too complex)
// we will follow a pattern similar to unfold => create the let then call g_statement on it
// reads a slot.
/* {1} OPT.The go function for: g_statement(self:Call_slot,s:class,v:string,err:boolean,loop:any) [] */
func F_Generate_g_statement_Call_slot (self *Language.CallSlot ,s *ClaireClass ,v *ClaireString ,err *ClaireBoolean ,loop *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var varg *ClaireVariable  
      /* noccur = 2 */
      var varg_try07453 EID 
      /* Let:3 */{ 
        var g0746UU *ClaireType  
        /* noccur = 1 */
        var g0746UU_try07474 EID 
        g0746UU_try07474 = Core.F_CALL(Optimize.C_c_type,ARGS(self.Arg.ToEID()))
        /* ERROR PROTECTION INSERTED (g0746UU-varg_try07453) */
        if ErrorIn(g0746UU_try07474) {varg_try07453 = g0746UU_try07474
        } else {
        g0746UU = ToType(OBJ(g0746UU_try07474))
        varg_try07453 = EID{F_Generate_build_Variable_string(MakeString("v_slot"),g0746UU.Id()).Id(),0}
        }
        /* Let-3 */} 
      /* ERROR PROTECTION INSERTED (varg-Result) */
      if ErrorIn(varg_try07453) {Result = varg_try07453
      } else {
      varg = To_Variable(OBJ(varg_try07453))
      /* Let:3 */{ 
        var unfold *Language.Let  
        /* noccur = 1 */
        /* Let:4 */{ 
          var _CL_obj *Language.Let   = Language.To_Let(new(Language.Let).Is(Language.C_Let))
          /* noccur = 12 */
          _CL_obj.ClaireVar = varg
          _CL_obj.Value = self.Arg
          /* update:5 */{ 
            var va_arg1 *Language.Let  
            var va_arg2 *ClaireAny  
            va_arg1 = _CL_obj
            /* Let:6 */{ 
              var _CL_obj *Language.CallSlot   = Language.To_CallSlot(new(Language.CallSlot).Is(Language.C_Call_slot))
              /* noccur = 5 */
              _CL_obj.Selector = self.Selector
              _CL_obj.Arg = varg.Id()
              va_arg2 = _CL_obj.Id()
              /* Let-6 */} 
            /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
            va_arg1.Arg = va_arg2
            /* update-5 */} 
          unfold = _CL_obj
          /* Let-4 */} 
        Result = F_Generate_g_statement_Let(unfold,
          s,
          v,
          err,
          loop)
        /* Let-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: g_statement @ Call_slot (throw: true) 
func E_Generate_g_statement_Call_slot (self EID,s EID,v EID,err EID,loop EID) EID { 
    return /*(sm for g_statement @ Call_slot= EID)*/ F_Generate_g_statement_Call_slot(Language.To_CallSlot(OBJ(self)),
      ToClass(OBJ(s)),
      ToString(OBJ(v)),
      ToBoolean(OBJ(err)),
      ANY(loop) )} 
  
// reads an table.
/* {1} OPT.The go function for: g_statement(self:Call_table,s:class,v:string,err:boolean,loop:any) [] */
func F_Generate_g_statement_Call_table (self *Language.CallTable ,s *ClaireClass ,v *ClaireString ,err *ClaireBoolean ,loop *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var varg *ClaireVariable  
      /* noccur = 2 */
      var varg_try07483 EID 
      /* Let:3 */{ 
        var g0749UU *ClaireType  
        /* noccur = 1 */
        var g0749UU_try07504 EID 
        g0749UU_try07504 = Core.F_CALL(Optimize.C_c_type,ARGS(self.Arg.ToEID()))
        /* ERROR PROTECTION INSERTED (g0749UU-varg_try07483) */
        if ErrorIn(g0749UU_try07504) {varg_try07483 = g0749UU_try07504
        } else {
        g0749UU = ToType(OBJ(g0749UU_try07504))
        varg_try07483 = EID{F_Generate_build_Variable_string(MakeString("v_table"),g0749UU.Id()).Id(),0}
        }
        /* Let-3 */} 
      /* ERROR PROTECTION INSERTED (varg-Result) */
      if ErrorIn(varg_try07483) {Result = varg_try07483
      } else {
      varg = To_Variable(OBJ(varg_try07483))
      /* Let:3 */{ 
        var unfold *Language.Let  
        /* noccur = 1 */
        /* Let:4 */{ 
          var _CL_obj *Language.Let   = Language.To_Let(new(Language.Let).Is(Language.C_Let))
          /* noccur = 12 */
          _CL_obj.ClaireVar = varg
          _CL_obj.Value = self.Arg
          /* update:5 */{ 
            var va_arg1 *Language.Let  
            var va_arg2 *ClaireAny  
            va_arg1 = _CL_obj
            /* Let:6 */{ 
              var _CL_obj *Language.CallTable   = Language.To_CallTable(new(Language.CallTable).Is(Language.C_Call_table))
              /* noccur = 5 */
              _CL_obj.Selector = self.Selector
              _CL_obj.Arg = varg.Id()
              va_arg2 = _CL_obj.Id()
              /* Let-6 */} 
            /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
            va_arg1.Arg = va_arg2
            /* update-5 */} 
          unfold = _CL_obj
          /* Let-4 */} 
        Result = F_Generate_g_statement_Let(unfold,
          s,
          v,
          err,
          loop)
        /* Let-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: g_statement @ Call_table (throw: true) 
func E_Generate_g_statement_Call_table (self EID,s EID,v EID,err EID,loop EID) EID { 
    return /*(sm for g_statement @ Call_table= EID)*/ F_Generate_g_statement_Call_table(Language.To_CallTable(OBJ(self)),
      ToClass(OBJ(s)),
      ToString(OBJ(v)),
      ToBoolean(OBJ(err)),
      ANY(loop) )} 
  
// reads an array.
/* {1} OPT.The go function for: g_statement(self:Call_array,s:class,v:string,err:boolean,loop:any) [] */
func F_Generate_g_statement_Call_array (self *Language.CallArray ,s *ClaireClass ,v *ClaireString ,err *ClaireBoolean ,loop *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var varg1 *ClaireVariable   = F_Generate_build_Variable_string(MakeString("va_arg1"),C_array.Id())
      /* noccur = 2 */
      /* Let:3 */{ 
        var varg2 *ClaireVariable   = F_Generate_build_Variable_string(MakeString("va_arg2"),C_integer.Id())
        /* noccur = 2 */
        /* Let:4 */{ 
          var unfold *Language.Let  
          /* noccur = 1 */
          /* Let:5 */{ 
            var _CL_obj *Language.Let   = Language.To_Let(new(Language.Let).Is(Language.C_Let))
            /* noccur = 13 */
            _CL_obj.ClaireVar = varg1
            _CL_obj.Value = self.Selector
            /* update:6 */{ 
              var va_arg1 *Language.Let  
              var va_arg2 *ClaireAny  
              va_arg1 = _CL_obj
              /* Let:7 */{ 
                var _CL_obj *Language.Let   = Language.To_Let(new(Language.Let).Is(Language.C_Let))
                /* noccur = 6 */
                _CL_obj.ClaireVar = varg2
                _CL_obj.Value = self.Arg
                _CL_obj.Arg = Language.C_Call_array.Make(varg1.Id(),varg2.Id(),self.Test)
                va_arg2 = _CL_obj.Id()
                /* Let-7 */} 
              /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
              va_arg1.Arg = va_arg2
              /* update-6 */} 
            unfold = _CL_obj
            /* Let-5 */} 
          Result = F_Generate_g_statement_Let(unfold,
            s,
            v,
            err,
            loop)
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: g_statement @ Call_array (throw: true) 
func E_Generate_g_statement_Call_array (self EID,s EID,v EID,err EID,loop EID) EID { 
    return /*(sm for g_statement @ Call_array= EID)*/ F_Generate_g_statement_Call_array(Language.To_CallArray(OBJ(self)),
      ToClass(OBJ(s)),
      ToString(OBJ(v)),
      ToBoolean(OBJ(err)),
      ANY(loop) )} 
  
// places a value in a slot with similar conventions ------------------------------------------------------------------------
// Update = [R(x) := y] where R(x) is a Call_slot, a call_array or a call_table 
// THIS USE OF self.arg IS MEGA UGLY AND SHOULD BE SIMPLIFIED IN THE OPTIMIZER LATER ON ... THERE SHOULD AT LEAST EXIST SOME COMMENTS !
// self.arg is a meta parameter /  it is a property (add or put ...) unless a demon if_write is used 
// self.value is Y and self.var is R(x)  => look in goexp.cl 
/* {1} OPT.The go function for: g_statement(self:Update,s:class,v:string,err:boolean,loop:any) [] */
func F_Generate_g_statement_Update (self *Language.Update ,s *ClaireClass ,v *ClaireString ,err *ClaireBoolean ,loop *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var X *ClaireAny   = self.ClaireVar
      /* noccur = 11 */
      /* Let:3 */{ 
        var p *ClaireAny   = self.Selector
        /* noccur = 3 */
        /* Let:4 */{ 
          var sr *ClaireType  
          /* noccur = 3 */
          if (X.Isa.IsIn(Language.C_Call_slot) == CTRUE) /* If:5 */{ 
            /* Let:6 */{ 
              var g0751 *Language.CallSlot   = Language.To_CallSlot(X)
              /* noccur = 2 */
              if (self.Arg == C_add.Id()) /* If:7 */{ 
                sr = Core.F_member_type(g0751.Selector.Range)
                } else {
                sr = g0751.Selector.Range
                /* If-7 */} 
              /* Let-6 */} 
            /* If!5 */}  else if (X.Isa.IsIn(Language.C_Call_array) == CTRUE) /* If:5 */{ 
            /* Let:6 */{ 
              var g0752 *Language.CallArray   = Language.To_CallArray(X)
              /* noccur = 1 */
              if (ToType(g0752.Test).Included(ToType(C_float.Id())) == CTRUE) /* If:7 */{ 
                sr = ToType(C_float.Id())
                } else {
                sr = ToType(C_any.Id())
                /* If-7 */} 
              /* Let-6 */} 
            } else {
            /* Let:6 */{ 
              var g0757UU *ClaireAny  
              /* noccur = 1 */
              if (self.Arg == C_add.Id()) /* If:7 */{ 
                g0757UU = Core.F_member_type(ToType(OBJ(Core.F_CALL(C_range,ARGS(p.ToEID()))))).Id()
                } else {
                g0757UU = ANY(Core.F_CALL(C_range,ARGS(p.ToEID())))
                /* If-7 */} 
              sr = Core.F_U_type(ToType(C_any.Id()),ToType(g0757UU))
              /* Let-6 */} 
            /* If-5 */} 
          
          var g0758I *ClaireBoolean  
          var g0758I_try07595 EID 
          /* and:5 */{ 
            var v_and5 *ClaireBoolean  
            
            v_and5 = err.Not
            if (v_and5 == CFALSE) {g0758I_try07595 = EID{CFALSE.Id(),0}
            } else /* arg:6 */{ 
              var v_and5_try07607 EID 
              v_and5_try07607 = F_Generate_g_func_any(ANY(Core.F_CALL(C_arg,ARGS(self.ClaireVar.ToEID()))))
              /* ERROR PROTECTION INSERTED (v_and5-g0758I_try07595) */
              if ErrorIn(v_and5_try07607) {g0758I_try07595 = v_and5_try07607
              } else {
              v_and5 = ToBoolean(OBJ(v_and5_try07607))
              if (v_and5 == CFALSE) {g0758I_try07595 = EID{CFALSE.Id(),0}
              } else /* arg:7 */{ 
                var v_and5_try07618 EID 
                v_and5_try07618 = F_Generate_g_func_any(self.Value)
                /* ERROR PROTECTION INSERTED (v_and5-g0758I_try07595) */
                if ErrorIn(v_and5_try07618) {g0758I_try07595 = v_and5_try07618
                } else {
                v_and5 = ToBoolean(OBJ(v_and5_try07618))
                if (v_and5 == CFALSE) {g0758I_try07595 = EID{CFALSE.Id(),0}
                } else /* arg:8 */{ 
                  v_and5 = Equal(s.Id(),C_void.Id())
                  if (v_and5 == CFALSE) {g0758I_try07595 = EID{CFALSE.Id(),0}
                  } else /* arg:9 */{ 
                    g0758I_try07595 = EID{CTRUE.Id(),0}/* arg-9 */} 
                  /* arg-8 */} 
                /* arg-7 */} 
              /* arg-6 */} 
            }}
            /* and-5 */} 
          /* ERROR PROTECTION INSERTED (g0758I-Result) */
          if ErrorIn(g0758I_try07595) {Result = g0758I_try07595
          } else {
          g0758I = ToBoolean(OBJ(g0758I_try07595))
          if (g0758I == CTRUE) /* If:5 */{ 
            Result = F_Generate_update_statement_Update(self,sr.Class_I())
            } else {
            /* Let:6 */{ 
              var try_count int  = 0
              /* noccur = 5 */
              /* Let:7 */{ 
                var varg1 *ClaireVariable  
                /* noccur = 5 */
                /* Let:8 */{ 
                  var g0764UU *ClaireType  
                  /* noccur = 1 */
                  if (X.Isa.IsIn(Language.C_Call_slot) == CTRUE) /* If:9 */{ 
                    /* Let:10 */{ 
                      var g0754 *Language.CallSlot   = Language.To_CallSlot(X)
                      /* noccur = 1 */
                      g0764UU = ToType(Core.F_domain_I_restriction(ToRestriction(g0754.Selector.Id())).Id())
                      /* Let-10 */} 
                    /* If!9 */}  else if (X.Isa.IsIn(Language.C_Call_array) == CTRUE) /* If:9 */{ 
                    g0764UU = ToType(C_integer.Id())
                    } else {
                    g0764UU = Core.F_U_type(ToType(C_any.Id()),ToType(OBJ(Core.F_CALL(C_domain,ARGS(p.ToEID())))))
                    /* If-9 */} 
                  varg1 = F_Generate_build_Variable_string(MakeString("va_arg1"),g0764UU.Id())
                  /* Let-8 */} 
                /* Let:8 */{ 
                  var varg2 *ClaireVariable   = F_Generate_build_Variable_string(MakeString("va_arg2"),sr.Id())
                  /* noccur = 7 */
                  /* Let:9 */{ 
                    var _Zcall *ClaireAny  
                    /* noccur = 1 */
                    var _Zcall_try076510 EID 
                    /* Let:10 */{ 
                      var xx *ClaireAny   = ANY(Core.F_CALL(C_copy,ARGS(X.ToEID())))
                      /* noccur = 2 */
                      _Zcall_try076510 = Core.F_put_property2(C_arg,ToObject(xx),varg1.Id())
                      /* ERROR PROTECTION INSERTED (_Zcall_try076510-_Zcall_try076510) */
                      if !ErrorIn(_Zcall_try076510) {
                      _Zcall_try076510 = xx.ToEID()
                      }
                      /* Let-10 */} 
                    /* ERROR PROTECTION INSERTED (_Zcall-Result) */
                    if ErrorIn(_Zcall_try076510) {Result = _Zcall_try076510
                    } else {
                    _Zcall = ANY(_Zcall_try076510)
                    /* Let:10 */{ 
                      var _Zunfold *Language.Update  
                      /* noccur = 2 */
                      /* Let:11 */{ 
                        var _CL_obj *Language.Update   = Language.To_Update(new(Language.Update).Is(Language.C_Update))
                        /* noccur = 8 */
                        _CL_obj.Selector = self.Selector
                        _CL_obj.Value = varg2.Id()
                        _CL_obj.Arg = self.Arg
                        _CL_obj.ClaireVar = _Zcall
                        _Zunfold = _CL_obj
                        /* Let-11 */} 
                      F_Generate_new_block_string(MakeString("update"))
                      /* Let:11 */{ 
                        var g0766UU *ClaireString  
                        /* noccur = 1 */
                        var g0766UU_try076712 EID 
                        g0766UU_try076712 = F_Generate_c_string_go_producer1(ToGenerateGoProducer(Optimize.C_PRODUCER.Value),varg1)
                        /* ERROR PROTECTION INSERTED (g0766UU-Result) */
                        if ErrorIn(g0766UU_try076712) {Result = g0766UU_try076712
                        } else {
                        g0766UU = ToString(OBJ(g0766UU_try076712))
                        F_Generate_var_declaration_string(g0766UU,F_Generate_go_range_Variable(varg1),1)
                        Result = EVOID
                        }
                        /* Let-11 */} 
                      /* ERROR PROTECTION INSERTED (Result-Result) */
                      if !ErrorIn(Result) {
                      /* Let:11 */{ 
                        var g0768UU *ClaireString  
                        /* noccur = 1 */
                        var g0768UU_try076912 EID 
                        g0768UU_try076912 = F_Generate_c_string_go_producer1(ToGenerateGoProducer(Optimize.C_PRODUCER.Value),varg2)
                        /* ERROR PROTECTION INSERTED (g0768UU-Result) */
                        if ErrorIn(g0768UU_try076912) {Result = g0768UU_try076912
                        } else {
                        g0768UU = ToString(OBJ(g0768UU_try076912))
                        F_Generate_var_declaration_string(g0768UU,F_Generate_go_range_Variable(varg2),1)
                        Result = EVOID
                        }
                        /* Let-11 */} 
                      /* ERROR PROTECTION INSERTED (Result-Result) */
                      if !ErrorIn(Result) {
                      var g0770I *ClaireBoolean  
                      var g0770I_try077111 EID 
                      g0770I_try077111 = Optimize.F_Compile_g_throw_any(ANY(Core.F_CALL(C_arg,ARGS(X.ToEID()))))
                      /* ERROR PROTECTION INSERTED (g0770I-Result) */
                      if ErrorIn(g0770I_try077111) {Result = g0770I_try077111
                      } else {
                      g0770I = ToBoolean(OBJ(g0770I_try077111))
                      if (g0770I == CTRUE) /* If:11 */{ 
                        try_count = (try_count+1)
                        Result = Core.F_CALL(C_Generate_g_try,ARGS(Core.F_CALL(C_arg,ARGS(X.ToEID())),
                          EID{MakeString("va_arg1").Id(),0},
                          EID{F_Generate_go_range_Variable(varg1).Id(),0},
                          EID{(v).Id(),0},
                          EID{CFALSE.Id(),0}))
                        } else {
                        Result = F_Generate_statement_any(ANY(Core.F_CALL(C_arg,ARGS(X.ToEID()))),F_Generate_go_range_Variable(varg1),MakeString("va_arg1"),loop)
                        /* If-11 */} 
                      }
                      /* ERROR PROTECTION INSERTED (Result-Result) */
                      if !ErrorIn(Result) {
                      var g0772I *ClaireBoolean  
                      var g0772I_try077311 EID 
                      g0772I_try077311 = Optimize.F_Compile_g_throw_any(self.Value)
                      /* ERROR PROTECTION INSERTED (g0772I-Result) */
                      if ErrorIn(g0772I_try077311) {Result = g0772I_try077311
                      } else {
                      g0772I = ToBoolean(OBJ(g0772I_try077311))
                      if (g0772I == CTRUE) /* If:11 */{ 
                        try_count = (try_count+1)
                        Result = Core.F_CALL(C_Generate_g_try,ARGS(self.Value.ToEID(),
                          EID{MakeString("va_arg2").Id(),0},
                          EID{F_Generate_go_range_Variable(varg2).Id(),0},
                          EID{(v).Id(),0},
                          EID{CFALSE.Id(),0}))
                        } else {
                        Result = F_Generate_statement_any(self.Value,F_Generate_go_range_Variable(varg2),MakeString("va_arg2"),loop)
                        /* If-11 */} 
                      }
                      /* ERROR PROTECTION INSERTED (Result-Result) */
                      if !ErrorIn(Result) {
                      PRINC("/* ---------- now we compile update ")
                      Result = Core.F_print_any(_Zunfold.Id())
                      /* ERROR PROTECTION INSERTED (Result-Result) */
                      if !ErrorIn(Result) {
                      PRINC(" ------- */")
                      F_Generate_breakline_void()
                      PRINC("")
                      Result = EVOID
                      }
                      {
                      Result = F_Generate_update_statement_Update(_Zunfold,sr.Class_I())
                      /* ERROR PROTECTION INSERTED (Result-Result) */
                      if !ErrorIn(Result) {
                      if (s.Id() != C_void.Id()) /* If:11 */{ 
                        F_princ_string(v)
                        PRINC(" = ")
                        Result = F_Generate_cast_prefix_class(F_Generate_go_range_Variable(varg2),s)
                        /* ERROR PROTECTION INSERTED (Result-Result) */
                        if !ErrorIn(Result) {
                        PRINC("va_arg2")
                        F_Generate_cast_post_class(F_Generate_go_range_Variable(varg2),s)
                        F_Generate_breakline_void()
                        PRINC("")
                        Result = EVOID
                        }
                        } else {
                        Result = EID{CFALSE.Id(),0}
                        /* If-11 */} 
                      /* ERROR PROTECTION INSERTED (Result-Result) */
                      if !ErrorIn(Result) {
                      F_Generate_close_try_integer(try_count)
                      F_Generate_close_block_string(MakeString("update"))
                      Result = EVOID
                      }}}}}}}
                      /* Let-10 */} 
                    }
                    /* Let-9 */} 
                  /* Let-8 */} 
                /* Let-7 */} 
              /* Let-6 */} 
            /* If-5 */} 
          }
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: g_statement @ Update (throw: true) 
func E_Generate_g_statement_Update (self EID,s EID,v EID,err EID,loop EID) EID { 
    return /*(sm for g_statement @ Update= EID)*/ F_Generate_g_statement_Update(Language.To_Update(OBJ(self)),
      ToClass(OBJ(s)),
      ToString(OBJ(v)),
      ToBoolean(OBJ(err)),
      ANY(loop) )} 
  
// this produce the code for an update assuming that self is error-free and functional
// this methiod handles
//    if_write demons (that perform the update)  p_write(x:any,y:any)
//    defeasible updates   o.StoreX(n,v,CTRUE)
// if we cannot find n (type too generic) => revert to a generic Update method
/* {1} OPT.The go function for: update_statement(self:Update,s:class) [] */
func F_Generate_update_statement_Update (self *Language.Update ,s *ClaireClass ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var p *ClaireAny   = self.Selector
      /* noccur = 6 */
      /* Let:3 */{ 
        var a *ClaireAny   = self.Arg
        /* noccur = 3 */
        /* Let:4 */{ 
          var v *ClaireAny   = self.Value
          /* noccur = 3 */
          /* Let:5 */{ 
            var x *ClaireAny   = self.ClaireVar
            /* noccur = 3 */
            var g0776I *ClaireBoolean  
            if (p.Isa.IsIn(C_relation) == CTRUE) /* If:6 */{ 
              /* Let:7 */{ 
                var g0774 *ClaireRelation   = ToRelation(p)
                /* noccur = 1 */
                g0776I = MakeBoolean((g0774.IfWrite != CNULL) && (a != C_put.Id()) && (a != Core.C_put_store.Id()))
                /* Let-7 */} 
              } else {
              g0776I = CFALSE
              /* If-6 */} 
            if (g0776I == CTRUE) /* If:6 */{ 
              F_c_princ_string(ToSymbol(OBJ(Core.F_CALL(C_name,ARGS(p.ToEID())))).String_I())
              PRINC("_write(")
              Result = Core.F_CALL(C_Generate_g_expression,ARGS(Core.F_CALL(C_arg,ARGS(x.ToEID())),EID{C_any.Id(),0}))
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              PRINC(",")
              Result = Core.F_CALL(C_Generate_g_expression,ARGS(v.ToEID(),EID{C_any.Id(),0}))
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              PRINC(")")
              Result = EVOID
              }}
              } else {
              var g0777I *ClaireBoolean  
              if (p.Isa.IsIn(C_relation) == CTRUE) /* If:7 */{ 
                /* Let:8 */{ 
                  var g0775 *ClaireRelation   = ToRelation(p)
                  /* noccur = 1 */
                  g0777I = MakeBoolean((g0775.Store_ask == CTRUE) || (a == Core.C_put_store.Id()))
                  /* Let-8 */} 
                } else {
                g0777I = CFALSE
                /* If-7 */} 
              if (g0777I == CTRUE) /* If:7 */{ 
                /* Let:8 */{ 
                  var s2 *ClaireClass   = ToTypeExpression(OBJ(Core.F_CALL(C_arg,ARGS(x.ToEID())))).Class_I()
                  /* noccur = 1 */
                  /* Let:9 */{ 
                    var n int  = INT(Core.F_CALL(C_mClaire_index,ARGS(EID{Core.F__at_property1(ToProperty(p),s2).Id(),0})))
                    /* noccur = 1 */
                    Result = Core.F_CALL(C_Generate_g_expression,ARGS(EID{ClEnv.Id(),0}))
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    PRINC(".Store")
                    /* Let:10 */{ 
                      var g0778UU *ClaireString  
                      /* noccur = 1 */
                      if (s.Id() == C_integer.Id()) /* If:11 */{ 
                        g0778UU = MakeString("Integer")
                        /* If!11 */}  else if (s.Id() == C_float.Id()) /* If:11 */{ 
                        g0778UU = MakeString("Float")
                        } else {
                        g0778UU = MakeString("Object")
                        /* If-11 */} 
                      F_princ_string(g0778UU)
                      /* Let-10 */} 
                    PRINC("(")
                    F_princ_integer(n)
                    PRINC(",")
                    Result = Core.F_CALL(C_Generate_g_expression,ARGS(v.ToEID(),EID{s.Id(),0}))
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    PRINC(",CTRUE)")
                    Result = EVOID
                    }}
                    /* Let-9 */} 
                  /* Let-8 */} 
                } else {
                Result = Core.F_CALL(C_Generate_g_expression,ARGS(x.ToEID(),EID{s.Id(),0}))
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                PRINC(" = ")
                Result = Core.F_CALL(C_Generate_g_expression,ARGS(v.ToEID(),EID{s.Id(),0}))
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                F_Generate_breakline_void()
                PRINC("")
                Result = EVOID
                }}
                /* If-7 */} 
              /* If-6 */} 
            /* Let-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: update_statement @ Update (throw: true) 
func E_Generate_update_statement_Update (self EID,s EID) EID { 
    return /*(sm for update_statement @ Update= EID)*/ F_Generate_update_statement_Update(Language.To_Update(OBJ(self)),ToClass(OBJ(s)) )} 
  
// in the expansion of Defarray, we generate x.graph := make_list(29,unknonw) that we need to trap
/* {1} OPT.The go function for: need_shortcut(v:any) [] */
func F_Generate_need_shortcut_any (v *ClaireAny ) *ClaireBoolean  { 
    // procedure body with s =  
var Result *ClaireBoolean  
    if (v.Isa.IsIn(Language.C_Call_method) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0779 *Language.CallMethod   = Language.To_CallMethod(v)
        /* noccur = 1 */
        Result = Equal(g0779.Arg.Selector.Id(),C_make_list.Id())
        /* Let-3 */} 
      } else {
      Result = CFALSE
      /* If-2 */} 
    return Result} 
  
// The EID go function for: need_shortcut @ any (throw: false) 
func E_Generate_need_shortcut_any (v EID) EID { 
    return EID{/*(sm for need_shortcut @ any= boolean)*/ F_Generate_need_shortcut_any(ANY(v) ).Id(),0}} 
  