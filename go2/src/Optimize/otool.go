/***** CLAIRE Compilation of file /Users/ycaseau/Dropbox/src/clairev4.0/src/compile/otool.cl 
         [version 4.0.01 / safety 5] Saturday 10-30-2021 *****/

package Optimize
import (_ "fmt"
	. "Kernel"
	"Core"
	"Language"
	"Reader"
)

//-------- dumb function to prevent import errors --------
func import_g0246() { 
    _ = Core.It
    _ = Language.It
    _ = Reader.It
    } 
  
  
//+-------------------------------------------------------------+
//| CLAIRE                                                      |
//| otool.cl                                                    |
//| Copyright (C) 1994 - 2013 Yves Caseau. All Rights Reserved  |
//| cf. copyright info in file object.cl: about()               |
//+-------------------------------------------------------------+
//-------------------------------------------------------------------
// this file contains the auxiliairy methods for the source optimizer
//-----------------------------------------------------------------
// ******************************************************************
// *  Table of contents                                             *
// *    Part 1: New Instructions & associated stuff                 *
// *    Part 2: Optimizer Warnings                                  *
// *    Part 3: Type Handling                                       *
// *    Part 4: Miscellaneous                                       *
// ******************************************************************
// ******************************************************************
// *    Part 1: New Instructions & associated stuff                 *
// ******************************************************************
//
// this is a same-sort (object) casting from one class to another because of the
// stupidity of the target type system
// its use is linked to stupid_t(x)
/* {1} OPT.The go function for: self_print(self:Compile/C_cast) [] */
func (self *CompileCCast ) SelfPrint () EID { 
    var Result EID 
    PRINC("<")
    Result = Core.F_CALL(C_print,ARGS(self.Arg.ToEID()))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(":")
    Result = Core.F_print_any(self.SetArg.Id())
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(">")
    Result = EVOID
    }}
    return Result} 
  
// The EID go function for: self_print @ Compile/C_cast (throw: true) 
func E_self_print_C_cast (self EID) EID { 
    return /*(sm for self_print @ Compile/C_cast= EID)*/ To_CompileCCast(OBJ(self)).SelfPrint( )} 
  
/* {1} OPT.The go function for: Compile/c_gc?(self:Compile/C_cast) [] */
func (self *CompileCCast ) CGc_ask () *ClaireBoolean  { 
    // use function body compiling 
return  To_CompileCCast(self.Arg).CGc_ask()
    } 
  
// The EID go function for: Compile/c_gc? @ Compile/C_cast (throw: false) 
func E_Compile_c_gc_ask_C_cast (self EID) EID { 
    return EID{/*(sm for Compile/c_gc? @ Compile/C_cast= boolean)*/ To_CompileCCast(OBJ(self)).CGc_ask( ).Id(),0}} 
  
/* {1} OPT.The go function for: c_type(self:Compile/C_cast) [] */
func (self *CompileCCast ) CType () *ClaireType  { 
    // use function body compiling 
return  ToType(self.SetArg.Id())
    } 
  
// The EID go function for: c_type @ Compile/C_cast (throw: false) 
func E_c_type_C_cast (self EID) EID { 
    return EID{/*(sm for c_type @ Compile/C_cast= type)*/ To_CompileCCast(OBJ(self)).CType( ).Id(),0}} 
  
// v3.0 : better safe
/* {1} OPT.The go function for: c_code(self:Compile/C_cast,s:class) [] */
func F_c_code_C_cast (self *CompileCCast ,s *ClaireClass ) EID { 
    var Result EID 
    if (s.IsIn(C_object) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var _CL_obj *CompileCCast   = To_CompileCCast(new(CompileCCast).Is(C_Compile_C_cast))
        /* noccur = 4 */
        /* update:4 */{ 
          var va_arg1 *CompileCCast  
          var va_arg2 *ClaireAny  
          va_arg1 = _CL_obj
          var va_arg2_try02475 EID 
          va_arg2_try02475 = Core.F_CALL(C_c_code,ARGS(self.Arg.ToEID(),EID{s.Id(),0}))
          /* ERROR PROTECTION INSERTED (va_arg2-Result) */
          if ErrorIn(va_arg2_try02475) {Result = va_arg2_try02475
          } else {
          va_arg2 = ANY(va_arg2_try02475)
          /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
          va_arg1.Arg = va_arg2
          Result = va_arg2.ToEID()
          }
          /* update-4 */} 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        _CL_obj.SetArg = self.SetArg
        Result = EID{_CL_obj.Id(),0}
        }
        /* Let-3 */} 
      } else {
      Result = Core.F_CALL(C_c_code,ARGS(self.Arg.ToEID(),EID{s.Id(),0}))
      /* If-2 */} 
    return Result} 
  
// The EID go function for: c_code @ Compile/C_cast (throw: true) 
func E_c_code_C_cast (self EID,s EID) EID { 
    return /*(sm for c_code @ Compile/C_cast= EID)*/ F_c_code_C_cast(To_CompileCCast(OBJ(self)),ToClass(OBJ(s)) )} 
  
// we need a new type to express powerful Iterate rules
// Note: Patterns require the compiler !
/* {1} OPT.The go function for: self_print(self:Pattern) [] */
func (self *ClairePattern ) SelfPrint () EID { 
    var Result EID 
    Result = Core.F_print_any(self.Selector.Id())
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("[tuple(")
    Result = Core.F_princ_list(self.Arg)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(")]")
    Result = EVOID
    }}
    return Result} 
  
// The EID go function for: self_print @ Pattern (throw: true) 
func E_self_print_Pattern (self EID) EID { 
    return /*(sm for self_print @ Pattern= EID)*/ To_ClairePattern(OBJ(self)).SelfPrint( )} 
  
/* {1} OPT.The go function for: %(x:any,y:Pattern) [] */
func F__Z_any3 (x *ClaireAny ,y *ClairePattern ) EID { 
    var Result EID 
    if (x.Isa.IsIn(Language.C_Call) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0248 *Language.Call   = Language.To_Call(x)
        /* noccur = 2 */
        /* and:4 */{ 
          var v_and4 *ClaireBoolean  
          
          v_and4 = Equal(g0248.Selector.Id(),y.Selector.Id())
          if (v_and4 == CFALSE) {Result = EID{CFALSE.Id(),0}
          } else /* arg:5 */{ 
            var v_and4_try02506 EID 
            /* Let:6 */{ 
              var g0251UU *ClaireList  
              /* noccur = 1 */
              var g0251UU_try02527 EID 
              /* Iteration:7 */{ 
                var v_list7 *ClaireList  
                var z *ClaireAny  
                var v_local7 *ClaireAny  
                v_list7 = g0248.Args
                g0251UU_try02527 = EID{CreateList(ToType(CEMPTY.Id()),v_list7.Length()).Id(),0}
                for CLcount := 0; CLcount < v_list7.Length(); CLcount++{ 
                  z = v_list7.At(CLcount)
                  var v_local7_try02539 EID 
                  v_local7_try02539 = Core.F_CALL(C_c_type,ARGS(z.ToEID()))
                  /* ERROR PROTECTION INSERTED (v_local7-g0251UU_try02527) */
                  if ErrorIn(v_local7_try02539) {g0251UU_try02527 = v_local7_try02539
                  g0251UU_try02527 = v_local7_try02539
                  break
                  } else {
                  v_local7 = ANY(v_local7_try02539)
                  ToList(OBJ(g0251UU_try02527)).PutAt(CLcount,v_local7)
                  } 
                }
                /* Iteration-7 */} 
              /* ERROR PROTECTION INSERTED (g0251UU-v_and4_try02506) */
              if ErrorIn(g0251UU_try02527) {v_and4_try02506 = g0251UU_try02527
              } else {
              g0251UU = ToList(OBJ(g0251UU_try02527))
              v_and4_try02506 = EID{Core.F_tmatch_ask_list(g0251UU,y.Arg).Id(),0}
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (v_and4-Result) */
            if ErrorIn(v_and4_try02506) {Result = v_and4_try02506
            } else {
            v_and4 = ToBoolean(OBJ(v_and4_try02506))
            if (v_and4 == CFALSE) {Result = EID{CFALSE.Id(),0}
            } else /* arg:6 */{ 
              Result = EID{CTRUE.Id(),0}/* arg-6 */} 
            /* arg-5 */} 
          }
          /* and-4 */} 
        /* Let-3 */} 
      } else {
      Result = EID{CFALSE.Id(),0}
      /* If-2 */} 
    return Result} 
  
// The EID go function for: % @ list<type_expression>(any, Pattern) (throw: true) 
func E__Z_any3 (x EID,y EID) EID { 
    return /*(sm for % @ list<type_expression>(any, Pattern)= EID)*/ F__Z_any3(ANY(x),To_ClairePattern(OBJ(y)) )} 
  
// this is very lazy, we could do better
/* {1} OPT.The go function for: glb(x:Pattern,y:type_expression) [] */
func (x *ClairePattern ) Glb (y *ClaireTypeExpression ) EID { 
    var Result EID 
    if (y.Isa.IsIn(C_Optimize_Pattern) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0254 *ClairePattern   = To_ClairePattern(y.Id())
        /* noccur = 2 */
        if (x.Selector.Id() == g0254.Selector.Id()) /* If:4 */{ 
          /* Let:5 */{ 
            var _CL_obj *ClairePattern   = To_ClairePattern(new(ClairePattern).Is(C_Optimize_Pattern))
            /* noccur = 4 */
            _CL_obj.Selector = x.Selector
            Result = Core.F_write_property(C_args,ToObject(_CL_obj.Id()),ANY(Core.F_CALL(ToProperty(Core.C_glb.Id()),ARGS(Core.F_CALL(C_args,ARGS(EID{x.Id(),0})),Core.F_CALL(C_args,ARGS(EID{g0254.Id(),0}))))))
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            Result = EID{_CL_obj.Id(),0}
            }
            /* Let-5 */} 
          } else {
          Result = EID{CEMPTY.Id(),0}
          /* If-4 */} 
        /* Let-3 */} 
      } else {
      Result = EID{CEMPTY.Id(),0}
      /* If-2 */} 
    return Result} 
  
// The EID go function for: glb @ Pattern (throw: true) 
func E_glb_Pattern (x EID,y EID) EID { 
    return /*(sm for glb @ Pattern= EID)*/ To_ClairePattern(OBJ(x)).Glb(ToTypeExpression(OBJ(y)) )} 
  
// extension of <= for Patterns
/* {1} OPT.The go function for: less?(x:Pattern,y:type_expression) [] */
func F_less_ask_Pattern (x *ClairePattern ,y *ClaireTypeExpression ) *ClaireBoolean  { 
    // procedure body with s =  
var Result *ClaireBoolean  
    if (y.Isa.IsIn(C_Optimize_Pattern) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0256 *ClairePattern   = To_ClairePattern(y.Id())
        /* noccur = 3 */
        /* and:4 */{ 
          var v_and4 *ClaireBoolean  
          
          v_and4 = Equal(x.Selector.Id(),g0256.Selector.Id())
          if (v_and4 == CFALSE) {Result = CFALSE
          } else /* arg:5 */{ 
            v_and4 = Equal(MakeInteger(x.Arg.Length()).Id(),MakeInteger(g0256.Arg.Length()).Id())
            if (v_and4 == CFALSE) {Result = CFALSE
            } else /* arg:6 */{ 
              /* Let:7 */{ 
                var g0259UU *ClaireAny  
                /* noccur = 1 */
                /* Let:8 */{ 
                  var i int  = 1
                  /* noccur = 5 */
                  /* Let:9 */{ 
                    var g0257 int  = x.Arg.Length()
                    /* noccur = 1 */
                    g0259UU= CFALSE.Id()
                    for (i <= g0257) /* while:10 */{ 
                      if (Core.F__equaltype_ask_any(ToType(x.Arg.At(i-1)),ToType(g0256.Arg.At(i-1))) != CTRUE) /* If:11 */{ 
                         /*v = g0259UU, s =any*/
g0259UU = CTRUE.Id()
                        break
                        /* If-11 */} 
                      i = (i+1)
                      /* while-10 */} 
                    /* Let-9 */} 
                  /* Let-8 */} 
                v_and4 = Core.F_not_any(g0259UU)
                /* Let-7 */} 
              if (v_and4 == CFALSE) {Result = CFALSE
              } else /* arg:7 */{ 
                Result = CTRUE/* arg-7 */} 
              /* arg-6 */} 
            /* arg-5 */} 
          /* and-4 */} 
        /* Let-3 */} 
      } else {
      Result = ToType(Language.C_Call.Id()).Included(ToType(y.Id()))
      /* If-2 */} 
    return Result} 
  
// The EID go function for: less? @ Pattern (throw: false) 
func E_less_ask_Pattern (x EID,y EID) EID { 
    return EID{/*(sm for less? @ Pattern= boolean)*/ F_less_ask_Pattern(To_ClairePattern(OBJ(x)),ToTypeExpression(OBJ(y)) ).Id(),0}} 
  
/* {1} OPT.The go function for: less?(x:type_expression,y:Pattern) [] */
func F_less_ask_type_expression2 (x *ClaireTypeExpression ,y *ClairePattern ) EID { 
    var Result EID 
    if (C_set.Id() == x.Isa.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var g0260 *ClaireSet   = ToSet(x.Id())
        /* noccur = 1 */
        /* Let:4 */{ 
          var g0266UU *ClaireAny  
          /* noccur = 1 */
          var g0266UU_try02675 EID 
          /* For:5 */{ 
            var z *ClaireAny  
            _ = z
            g0266UU_try02675= EID{CFALSE.Id(),0}
            for _,z = range(g0260.Values)/* loop:6 */{ 
              var void_try7 EID 
              _ = void_try7
              var g0268I *ClaireBoolean  
              var g0268I_try02697 EID 
              /* Let:7 */{ 
                var g0270UU *ClaireBoolean  
                /* noccur = 1 */
                var g0270UU_try02718 EID 
                g0270UU_try02718 = F__Z_any3(z,y)
                /* ERROR PROTECTION INSERTED (g0270UU-g0268I_try02697) */
                if ErrorIn(g0270UU_try02718) {g0268I_try02697 = g0270UU_try02718
                } else {
                g0270UU = ToBoolean(OBJ(g0270UU_try02718))
                g0268I_try02697 = EID{g0270UU.Not.Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (g0268I-void_try7) */
              if ErrorIn(g0268I_try02697) {void_try7 = g0268I_try02697
              } else {
              g0268I = ToBoolean(OBJ(g0268I_try02697))
              if (g0268I == CTRUE) /* If:7 */{ 
                 /*v = g0266UU_try02675, s =EID*/
g0266UU_try02675 = EID{CTRUE.Id(),0}
                break
                } else {
                void_try7 = EID{CFALSE.Id(),0}
                /* If-7 */} 
              }
              /* ERROR PROTECTION INSERTED (void_try7-g0266UU_try02675) */
              if ErrorIn(void_try7) {g0266UU_try02675 = void_try7
              g0266UU_try02675 = void_try7
              break
              } else {
              }
              /* loop-6 */} 
            /* For-5 */} 
          /* ERROR PROTECTION INSERTED (g0266UU-Result) */
          if ErrorIn(g0266UU_try02675) {Result = g0266UU_try02675
          } else {
          g0266UU = ANY(g0266UU_try02675)
          Result = EID{Core.F_not_any(g0266UU).Id(),0}
          }
          /* Let-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (x.Isa.IsIn(C_Optimize_Pattern) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0261 *ClairePattern   = To_ClairePattern(x.Id())
        /* noccur = 4 */
        /* and:4 */{ 
          var v_and4 *ClaireBoolean  
          
          v_and4 = Equal(g0261.Selector.Id(),y.Selector.Id())
          if (v_and4 == CFALSE) {Result = EID{CFALSE.Id(),0}
          } else /* arg:5 */{ 
            v_and4 = Equal(MakeInteger(g0261.Arg.Length()).Id(),MakeInteger(y.Arg.Length()).Id())
            if (v_and4 == CFALSE) {Result = EID{CFALSE.Id(),0}
            } else /* arg:6 */{ 
              /* Let:7 */{ 
                var g0272UU *ClaireAny  
                /* noccur = 1 */
                /* Let:8 */{ 
                  var i int  = 1
                  /* noccur = 5 */
                  /* Let:9 */{ 
                    var g0262 int  = g0261.Arg.Length()
                    /* noccur = 1 */
                    g0272UU= CFALSE.Id()
                    for (i <= g0262) /* while:10 */{ 
                      if (Core.F__equaltype_ask_any(ToType(g0261.Arg.At(i-1)),ToType(y.Arg.At(i-1))) != CTRUE) /* If:11 */{ 
                         /*v = g0272UU, s =any*/
g0272UU = CTRUE.Id()
                        break
                        /* If-11 */} 
                      i = (i+1)
                      /* while-10 */} 
                    /* Let-9 */} 
                  /* Let-8 */} 
                v_and4 = Core.F_not_any(g0272UU)
                /* Let-7 */} 
              if (v_and4 == CFALSE) {Result = EID{CFALSE.Id(),0}
              } else /* arg:7 */{ 
                Result = EID{CTRUE.Id(),0}/* arg-7 */} 
              /* arg-6 */} 
            /* arg-5 */} 
          /* and-4 */} 
        /* Let-3 */} 
      } else {
      Result = EID{CFALSE.Id(),0}
      /* If-2 */} 
    return Result} 
  
// The EID go function for: less? @ list<type_expression>(type_expression, Pattern) (throw: true) 
func E_less_ask_type_expression2 (x EID,y EID) EID { 
    return /*(sm for less? @ list<type_expression>(type_expression, Pattern)= EID)*/ F_less_ask_type_expression2(ToTypeExpression(OBJ(x)),To_ClairePattern(OBJ(y)) )} 
  
// v0.03 must return a type
/* {1} OPT.The go function for: nth(p:property,x:tuple) [] */
func F_nth_property (p *ClaireProperty ,x *ClaireTuple ) *ClairePattern  { 
    // procedure body with s =  
var Result *ClairePattern  
    /* Let:2 */{ 
      var _CL_obj *ClairePattern   = To_ClairePattern(new(ClairePattern).Is(C_Optimize_Pattern))
      /* noccur = 5 */
      _CL_obj.Selector = p
      _CL_obj.Arg = x.List_I()
      Result = _CL_obj
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: nth @ property (throw: false) 
func E_nth_property (p EID,x EID) EID { 
    return EID{/*(sm for nth @ property= Pattern)*/ F_nth_property(ToProperty(OBJ(p)),ToTuple(OBJ(x)) ).Id(),0}} 
  
// ******************************************************************
// *    Part 2: Optimizer Warnings                                  *
// ******************************************************************
// unified warning
/* {1} OPT.The go function for: Compile/warn(_CL_obj:void) [] */
func F_Compile_warn_void ()  { 
    // procedure body with s =  
if (C_OPT.InMethod != CNULL) /* If:2 */{ 
      Core.F_tformat_string(MakeString("---- WARNING[in ~S]: "),2,MakeConstantList(C_OPT.InMethod))
      } else {
      Core.F_tformat_string(MakeString("---- WARNING: "),2,ToType(CEMPTY.Id()).EmptyList())
      /* If-2 */} 
    } 
  
// The EID go function for: Compile/warn @ void (throw: false) 
func E_Compile_warn_void (_CL_obj EID) EID { 
    /*(sm for Compile/warn @ void= void)*/ F_Compile_warn_void( )
    return EVOID} 
  
/* {1} OPT.The go function for: Compile/Cerror(s:string,l:listargs) [] */
func F_Compile_Cerror_string (s *ClaireString ,l *ClaireList ) EID { 
    var Result EID 
    PRINC("---- Compiler Error[in ")
    Result = Core.F_CALL(C_print,ARGS(C_OPT.InMethod.ToEID()))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("]:\n")
    Result = EVOID
    }
    {
    PRINC("---- file read up to line ")
    F_princ_integer(ClEnv.NLine)
    PRINC("\n")
    Result = ToException(Core.C_general_error.Make((s).Id(),l.Id())).Close()
    }
    return Result} 
  
// The EID go function for: Compile/Cerror @ string (throw: true) 
func E_Compile_Cerror_string (s EID,l EID) EID { 
    return /*(sm for Compile/Cerror @ string= EID)*/ F_Compile_Cerror_string(ToString(OBJ(s)),ToList(OBJ(l)) )} 
  
// a note
/* {1} OPT.The go function for: Compile/notice(_CL_obj:void) [] */
func F_Compile_notice_void ()  { 
    // procedure body with s =  
if (C_OPT.InMethod != CNULL) /* If:2 */{ 
      Core.F_tformat_string(MakeString("---- note[in ~S]: "),3,MakeConstantList(C_OPT.InMethod))
      } else {
      Core.F_tformat_string(MakeString("---- note: "),3,ToType(CEMPTY.Id()).EmptyList())
      /* If-2 */} 
    } 
  
// The EID go function for: Compile/notice @ void (throw: false) 
func E_Compile_notice_void (_CL_obj EID) EID { 
    /*(sm for Compile/notice @ void= void)*/ F_Compile_notice_void( )
    return EVOID} 
  
// Warning : compiling is impossible, wrong selector
/* {1} OPT.The go function for: c_warn(self:Call,%type:any) [] */
func F_Optimize_c_warn_Call (self *Language.Call ,_Ztype *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var s *ClaireProperty   = self.Selector
      /* noccur = 6 */
      if (_Ztype == C_void.Id()) /* If:3 */{ 
        Result = Core.F_CALL(C_Compile_Cerror,ARGS(EID{MakeString("[205] message ~S sent to void object").Id(),0},EID{self.Id(),0}))
        /* If!3 */}  else if ((F_boolean_I_any(s.Restrictions.Id()).Id() != CTRUE.Id()) && 
          (C_OPT.Ignore.Contain_ask(s.Id()) != CTRUE)) /* If:3 */{ 
        F_Compile_warn_void()
        Result = Core.F_tformat_string(MakeString("the property ~S is undefined [255]\n"),2,MakeConstantList(s.Id()))
        } else {
        var g0274I *ClaireBoolean  
        /* and:4 */{ 
          var v_and4 *ClaireBoolean  
          
          v_and4 = C_OPT.Ignore.Contain_ask(s.Id()).Not
          if (v_and4 == CFALSE) {g0274I = CFALSE
          } else /* arg:5 */{ 
            v_and4 = MakeBoolean((s.Open <= 1) || (s.Open == 4))
            if (v_and4 == CFALSE) {g0274I = CFALSE
            } else /* arg:6 */{ 
              if (_Ztype.Isa.IsIn(C_list) == CTRUE) /* If:7 */{ 
                /* Let:8 */{ 
                  var g0273 *ClaireList   = ToList(_Ztype)
                  /* noccur = 1 */
                  v_and4 = Core.F__I_equal_any(MakeInteger(ToTypeExpression(g0273.At(1-1)).Class_I().Open).Id(),MakeInteger(3).Id())
                  /* Let-8 */} 
                } else {
                v_and4 = CFALSE
                /* If-7 */} 
              if (v_and4 == CFALSE) {g0274I = CFALSE
              } else /* arg:7 */{ 
                g0274I = CTRUE/* arg-7 */} 
              /* arg-6 */} 
            /* arg-5 */} 
          /* and-4 */} 
        if (g0274I == CTRUE) /* If:4 */{ 
          F_Compile_warn_void()
          Result = Core.F_tformat_string(MakeString("wrongly typed message ~S (~S) [256]\n"),2,MakeConstantList(self.Id(),_Ztype))
          /* If!4 */}  else if (C_compiler.Optimize_ask == CTRUE) /* If:4 */{ 
          F_Compile_notice_void()
          Result = Core.F_tformat_string(MakeString("poorly typed message ~S [~S]\n"),3,MakeConstantList(self.Id(),_Ztype))
          } else {
          Result = EID{CFALSE.Id(),0}
          /* If-4 */} 
        /* If-3 */} 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = F_Optimize_open_message_property(self.Selector,self.Args)
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_warn @ Call (throw: true) 
func E_Optimize_c_warn_Call (self EID,_Ztype EID) EID { 
    return /*(sm for c_warn @ Call= EID)*/ F_Optimize_c_warn_Call(Language.To_Call(OBJ(self)),ANY(_Ztype) )} 
  
/* {1} OPT.The go function for: c_warn(self:Super,%type:any) [] */
func F_Optimize_c_warn_Super (self *Language.Super ,_Ztype *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var s *ClaireProperty   = self.Selector
      /* noccur = 4 */
      if (_Ztype == C_void.Id()) /* If:3 */{ 
        Result = Core.F_CALL(C_Compile_Cerror,ARGS(EID{MakeString("[205] message ~S sent to void object").Id(),0},EID{self.Id(),0}))
        /* If!3 */}  else if (F_boolean_I_any(s.Restrictions.Id()).Id() != CTRUE.Id()) /* If:3 */{ 
        F_Compile_warn_void()
        Result = Core.F_tformat_string(MakeString("the property ~S is undefined [255]\n"),2,MakeConstantList(s.Id()))
        /* If!3 */}  else if ((C_OPT.Ignore.Contain_ask(s.Id()) != CTRUE) && 
          (s.Open <= 1)) /* If:3 */{ 
        Result = Core.F_tformat_string(MakeString("---- note: wrongly typed message ~S [~S]\n"),3,MakeConstantList(self.Id(),_Ztype))
        } else {
        Result = EID{CFALSE.Id(),0}
        /* If-3 */} 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      /* Let:3 */{ 
        var m *Language.Call  
        /* noccur = 2 */
        var m_try02754 EID 
        m_try02754 = F_Optimize_open_message_property(self.Selector,self.Args)
        /* ERROR PROTECTION INSERTED (m-Result) */
        if ErrorIn(m_try02754) {Result = m_try02754
        } else {
        m = Language.To_Call(OBJ(m_try02754))
        /* Let:4 */{ 
          var _CL_obj *Language.Super   = Language.To_Super(new(Language.Super).Is(Language.C_Super))
          /* noccur = 7 */
          _CL_obj.Selector = m.Selector
          _CL_obj.CastTo = self.CastTo
          _CL_obj.Args = m.Args
          Result = EID{_CL_obj.Id(),0}
          /* Let-4 */} 
        }
        /* Let-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_warn @ Super (throw: true) 
func E_Optimize_c_warn_Super (self EID,_Ztype EID) EID { 
    return /*(sm for c_warn @ Super= EID)*/ F_Optimize_c_warn_Super(Language.To_Super(OBJ(self)),ANY(_Ztype) )} 
  
// a message cannot be compiled into efficient code
// here the property does not allow the compilation and we want to see it
/* {1} OPT.The go function for: c_warn(self:property,l:list,%type:list) [] */
func F_Optimize_c_warn_property (self *ClaireProperty ,l *ClaireList ,_Ztype *ClaireList ) EID { 
    var Result EID 
    if ((self.Open <= 1) && 
        ((C_OPT.Ignore.Contain_ask(self.Id()) != CTRUE) && 
          (C_compiler.Safety > 1))) /* If:2 */{ 
      
      /* If-2 */} 
    Result = F_Optimize_open_message_property(self,l)
    return Result} 
  
// The EID go function for: c_warn @ property (throw: true) 
func E_Optimize_c_warn_property (self EID,l EID,_Ztype EID) EID { 
    return /*(sm for c_warn @ property= EID)*/ F_Optimize_c_warn_property(ToProperty(OBJ(self)),ToList(OBJ(l)),ToList(OBJ(_Ztype)) )} 
  
// a variable should not be abused ! Either it is a true error or it is
// simply dangerous. The result is the value to be used (either x or
// ckeck_in(x,range(oself))
/* {1} OPT.The go function for: c_warn(self:Variable,x:any,y:type) [] */
func F_Optimize_c_warn_Variable (self *ClaireVariable ,x *ClaireAny ,y *ClaireType ) EID { 
    var Result EID 
    var g0276I *ClaireBoolean  
    var g0276I_try02772 EID 
    /* Let:2 */{ 
      var g0278UU *ClaireBoolean  
      /* noccur = 1 */
      var g0278UU_try02793 EID 
      /* Let:3 */{ 
        var g0280UU *ClaireAny  
        /* noccur = 1 */
        var g0280UU_try02814 EID 
        g0280UU_try02814 = Core.F_CALL(ToProperty(C__exp.Id()),ARGS(EID{y.Id(),0},EID{self.Range.Id(),0}))
        /* ERROR PROTECTION INSERTED (g0280UU-g0278UU_try02793) */
        if ErrorIn(g0280UU_try02814) {g0278UU_try02793 = g0280UU_try02814
        } else {
        g0280UU = ANY(g0280UU_try02814)
        g0278UU_try02793 = EID{F_boolean_I_any(g0280UU).Id(),0}
        }
        /* Let-3 */} 
      /* ERROR PROTECTION INSERTED (g0278UU-g0276I_try02772) */
      if ErrorIn(g0278UU_try02793) {g0276I_try02772 = g0278UU_try02793
      } else {
      g0278UU = ToBoolean(OBJ(g0278UU_try02793))
      g0276I_try02772 = EID{Core.F__I_equal_any(g0278UU.Id(),CTRUE.Id()).Id(),0}
      }
      /* Let-2 */} 
    /* ERROR PROTECTION INSERTED (g0276I-Result) */
    if ErrorIn(g0276I_try02772) {Result = g0276I_try02772
    } else {
    g0276I = ToBoolean(OBJ(g0276I_try02772))
    if (g0276I == CTRUE) /* If:2 */{ 
      if (C_compiler.Safety > 4) /* If:3 */{ 
        F_Compile_warn_void()
        Result = Core.F_tformat_string(MakeString("~S of type ~S is put in the variable ~S:~S [257]\n"),2,MakeConstantList(x,
          y.Id(),
          self.Id(),
          self.Range.Id()))
        } else {
        Result = Core.F_CALL(C_Compile_Cerror,ARGS(EID{MakeString("[212] the value ~S of type ~S cannot be placed in the variable ~S:~S").Id(),0},
          x.ToEID(),
          EID{y.Id(),0},
          EID{self.Id(),0},
          EID{self.Range.Id(),0}))
        /* If-3 */} 
      /* If!2 */}  else if ((C_compiler.Safety <= 1) || 
        (F_boolean_I_any(F_Compile_sort_equal_class(F_Compile_osort_any(self.Range.Id()),F_Compile_osort_any(y.Id()))).Id() != CTRUE.Id())) /* If:2 */{ 
      F_Compile_warn_void()
      Result = Core.F_tformat_string(MakeString("~S of type ~S is put in the variable ~S:~S [257]\n"),2,MakeConstantList(x,
        y.Id(),
        self.Id(),
        self.Range.Id()))
      } else {
      Result = EID{CFALSE.Id(),0}
      /* If-2 */} 
    }
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    if ((C_compiler.Safety <= 1) && 
        (y.Included(self.Range) != CTRUE)) /* If:2 */{ 
      Result = Core.F_CALL(C_c_code,ARGS(x.ToEID(),EID{C_any.Id(),0}))
      } else {
      Result = x.ToEID()
      /* If-2 */} 
    }
    return Result} 
  
// The EID go function for: c_warn @ Variable (throw: true) 
func E_Optimize_c_warn_Variable (self EID,x EID,y EID) EID { 
    return /*(sm for c_warn @ Variable= EID)*/ F_Optimize_c_warn_Variable(To_Variable(OBJ(self)),ANY(x),ToType(OBJ(y)) )} 
  
// ******************************************************************
// *    Part 3: Type Handling                                       *
// ******************************************************************
// we use  {any U type} to represent the change of sort  (to any)
//         {} U (c U t) to represent a change of psort   (to c)
// e.g.: (any U class) = class stored as an OID
// tests if two sorts are similar
// the compiler.overflow? test
/* {1} OPT.The go function for: Compile/sort=(c:class,c2:class) [] */
func F_Compile_sort_equal_class (c *ClaireClass ,c2 *ClaireClass ) *ClaireAny  { 
    // use function body compiling 
if (c.IsIn(C_object) == CTRUE) /* body If:2 */{ 
      return  c2.IsIn(C_object).Id()
      } else {
      return  MakeBoolean((c.Id() == c2.Id()) || (((C_compiler.Overflow_ask != CTRUE) && 
            ((c.Id() == C_any.Id()) && 
                (c2.Id() == C_integer.Id()))) || 
          ((c.Id() == C_integer.Id()) && 
              (c2.Id() == C_any.Id())))).Id()
      /* body If-2 */} 
    } 
  
// The EID go function for: Compile/sort= @ class (throw: false) 
func E_Compile_sort_equal_class (c EID,c2 EID) EID { 
    return /*(sm for Compile/sort= @ class= any)*/ F_Compile_sort_equal_class(ToClass(OBJ(c)),ToClass(OBJ(c2)) ).ToEID()} 
  
// give the "precise sort", i.e., a class under object is a sort
/* {1} OPT.The go function for: Compile/psort(x:any) [] */
func F_Compile_psort_any (x *ClaireAny ) *ClaireClass  { 
    // procedure body with s =  
var Result *ClaireClass  
    /* Let:2 */{ 
      var c *ClaireClass   = ToTypeExpression(x).Class_I()
      /* noccur = 3 */
      if (c.IsIn(C_object) == CTRUE) /* If:3 */{ 
        Result = c
        } else {
        Result = c.Sort_I()
        /* If-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: Compile/psort @ any (throw: false) 
func E_Compile_psort_any (x EID) EID { 
    return EID{/*(sm for Compile/psort @ any= class)*/ F_Compile_psort_any(ANY(x) ).Id(),0}} 
  
// gives the "optimizer sort", which is one of
// any, object, float, X <= import,
/* {1} OPT.The go function for: Compile/osort(x:any) [] */
func F_Compile_osort_any (x *ClaireAny ) *ClaireClass  { 
    // use function body compiling 
return  ToTypeExpression(x).Class_I().Sort_I()
    } 
  
// The EID go function for: Compile/osort @ any (throw: false) 
func E_Compile_osort_any (x EID) EID { 
    return EID{/*(sm for Compile/osort @ any= class)*/ F_Compile_osort_any(ANY(x) ).Id(),0}} 
  
/* {1} OPT.The go function for: sort(x:Variable) [] */
func F_sort_Variable (x *ClaireVariable ) *ClaireClass  { 
    // procedure body with s =  
var Result *ClaireClass  
    /* Let:2 */{ 
      var r *ClaireType   = x.Range
      /* noccur = 4 */
      var g0283I *ClaireBoolean  
      if (r.Isa.IsIn(C_Union) == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var g0282 *ClaireUnion   = To_Union(r.Id())
          /* noccur = 1 */
          g0283I = Equal(g0282.T1.Id(),CEMPTY.Id())
          /* Let-4 */} 
        } else {
        g0283I = CFALSE
        /* If-3 */} 
      if (g0283I == CTRUE) /* If:3 */{ 
        Result = F_Compile_psort_any(To_Union(To_Union(r.Id()).T2.Id()).T2.Id())
        } else {
        Result = F_Compile_psort_any(r.Id())
        /* If-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: sort @ Variable (throw: false) 
func E_sort_Variable (x EID) EID { 
    return EID{/*(sm for sort @ Variable= class)*/ F_sort_Variable(To_Variable(OBJ(x)) ).Id(),0}} 
  
// this is a very stupid type inference that mimicks the go compiler - defined in pretty.cl with CLAIRE4
// it returns a class
/* {1} OPT.The go function for: Compile/stupid_t(self:any) [] */
func F_Compile_stupid_t_any1 (self *ClaireAny ) EID { 
    var Result EID 
    Result = Language.F_static_type_any(self)
    return Result} 
  
// The EID go function for: Compile/stupid_t @ list<type_expression>(any) (throw: true) 
func E_Compile_stupid_t_any1 (self EID) EID { 
    return /*(sm for Compile/stupid_t @ list<type_expression>(any)= EID)*/ F_Compile_stupid_t_any1(ANY(self) )} 
  
// comparison
/* {1} OPT.The go function for: Compile/stupid_t(self:any,x:any) [] */
func F_Compile_stupid_t_any2 (self *ClaireAny ,x *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var c1 *ClaireClass  
      /* noccur = 2 */
      var c1_try02843 EID 
      c1_try02843 = F_Compile_stupid_t_any1(self)
      /* ERROR PROTECTION INSERTED (c1-Result) */
      if ErrorIn(c1_try02843) {Result = c1_try02843
      } else {
      c1 = ToClass(OBJ(c1_try02843))
      /* Let:3 */{ 
        var c2 *ClaireClass  
        /* noccur = 1 */
        var c2_try02854 EID 
        c2_try02854 = F_Compile_stupid_t_any1(x)
        /* ERROR PROTECTION INSERTED (c2-Result) */
        if ErrorIn(c2_try02854) {Result = c2_try02854
        } else {
        c2 = ToClass(OBJ(c2_try02854))
        Result = EID{MakeBoolean((c1.Id() != C_any.Id()) && (c1.Id() == c2.Id())).Id(),0}
        }
        /* Let-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: Compile/stupid_t @ list<type_expression>(any, any) (throw: true) 
func E_Compile_stupid_t_any2 (self EID,x EID) EID { 
    return /*(sm for Compile/stupid_t @ list<type_expression>(any, any)= EID)*/ F_Compile_stupid_t_any2(ANY(self),ANY(x) )} 
  
// an extended type is of the kind (t U {unknown})
// CLAIRE4: got rid of optUnion
/* {1} OPT.The go function for: extended?(self:type) [] */
func F_Optimize_extended_ask_type (self *ClaireType ) EID { 
    var Result EID 
    if (self.Isa.IsIn(C_Union) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0286 *ClaireUnion   = To_Union(self.Id())
        /* noccur = 3 */
        /* and:4 */{ 
          var v_and4 *ClaireBoolean  
          
          v_and4 = Equal(C_set.Id(),g0286.T2.Isa.Id())
          if (v_and4 == CFALSE) {Result = EID{CFALSE.Id(),0}
          } else /* arg:5 */{ 
            var v_and4_try02886 EID 
            /* Let:6 */{ 
              var g0289UU *ClaireAny  
              /* noccur = 1 */
              var g0289UU_try02907 EID 
              g0289UU_try02907 = Core.F_CALL(C_size,ARGS(EID{g0286.T2.Id(),0}))
              /* ERROR PROTECTION INSERTED (g0289UU-v_and4_try02886) */
              if ErrorIn(g0289UU_try02907) {v_and4_try02886 = g0289UU_try02907
              } else {
              g0289UU = ANY(g0289UU_try02907)
              v_and4_try02886 = EID{Equal(g0289UU,MakeInteger(1).Id()).Id(),0}
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (v_and4-Result) */
            if ErrorIn(v_and4_try02886) {Result = v_and4_try02886
            } else {
            v_and4 = ToBoolean(OBJ(v_and4_try02886))
            if (v_and4 == CFALSE) {Result = EID{CFALSE.Id(),0}
            } else /* arg:6 */{ 
              var v_and4_try02917 EID 
              /* Let:7 */{ 
                var g0292UU *ClaireAny  
                /* noccur = 1 */
                var g0292UU_try02938 EID 
                g0292UU_try02938 = Core.F_the_type(g0286.T2)
                /* ERROR PROTECTION INSERTED (g0292UU-v_and4_try02917) */
                if ErrorIn(g0292UU_try02938) {v_and4_try02917 = g0292UU_try02938
                } else {
                g0292UU = ANY(g0292UU_try02938)
                v_and4_try02917 = EID{Equal(g0292UU,CNULL).Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (v_and4-Result) */
              if ErrorIn(v_and4_try02917) {Result = v_and4_try02917
              } else {
              v_and4 = ToBoolean(OBJ(v_and4_try02917))
              if (v_and4 == CFALSE) {Result = EID{CFALSE.Id(),0}
              } else /* arg:7 */{ 
                Result = EID{CTRUE.Id(),0}/* arg-7 */} 
              /* arg-6 */} 
            /* arg-5 */} 
          }}
          /* and-4 */} 
        /* Let-3 */} 
      } else {
      Result = EID{CFALSE.Id(),0}
      /* If-2 */} 
    return Result} 
  
// The EID go function for: extended? @ type (throw: true) 
func E_Optimize_extended_ask_type (self EID) EID { 
    return /*(sm for extended? @ type= EID)*/ F_Optimize_extended_ask_type(ToType(OBJ(self)) )} 
  
// creates an extended type (v0.02) that can be checked easily 
// CLAIRE4: we removed the syntactic marker optUnion for (X U {unknown})
// used in ocall and ocontrol
/* {1} OPT.The go function for: extends(x:type) [] */
func F_Optimize_extends_type (x *ClaireType ) *ClaireType  { 
    // procedure body with s =  
var Result *ClaireType  
    /* Let:2 */{ 
      var _CL_obj *ClaireUnion   = To_Union(new(ClaireUnion).Is(C_Union))
      /* noccur = 5 */
      _CL_obj.T1 = x
      _CL_obj.T2 = ToType(MakeConstantSet(CNULL).Id())
      Result = ToType(_CL_obj.Id())
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: extends @ type (throw: false) 
func E_Optimize_extends_type (x EID) EID { 
    return EID{/*(sm for extends @ type= type)*/ F_Optimize_extends_type(ToType(OBJ(x)) ).Id(),0}} 
  
// a sort abstraction is the special union any U t, which is known to represent t by
// the type system (used for variables only) but tells the compiler that the sort is any
/* {1} OPT.The go function for: sort_abstract!(x:type) [] */
func F_Optimize_sort_abstract_I_type (x *ClaireType ) *ClaireType  { 
    // procedure body with s =  
var Result *ClaireType  
    if (((ANY(Core.F_CALL(C_sort_I,ARGS(EID{x.Id(),0}))) != C_any.Id()) && 
          ((ANY(Core.F_CALL(C_sort_I,ARGS(EID{x.Id(),0}))) != C_integer.Id()) || 
              (C_compiler.Overflow_ask == CTRUE))) || 
        (x.Id() == C_float.Id())) /* If:2 */{ 
      /* Let:3 */{ 
        var _CL_obj *ClaireUnion   = To_Union(new(ClaireUnion).Is(C_Union))
        /* noccur = 5 */
        _CL_obj.T1 = ToType(C_any.Id())
        _CL_obj.T2 = x
        Result = ToType(_CL_obj.Id())
        /* Let-3 */} 
      } else {
      Result = x
      /* If-2 */} 
    return Result} 
  
// The EID go function for: sort_abstract! @ type (throw: false) 
func E_Optimize_sort_abstract_I_type (x EID) EID { 
    return EID{/*(sm for sort_abstract! @ type= type)*/ F_Optimize_sort_abstract_I_type(ToType(OBJ(x)) ).Id(),0}} 
  
// v3.00.05
/* {1} OPT.The go function for: sort_abstract?(x:type) [] */
func F_Optimize_sort_abstract_ask_type (x *ClaireType ) *ClaireBoolean  { 
    // procedure body with s =  
var Result *ClaireBoolean  
    if (x.Isa.IsIn(C_Union) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0294 *ClaireUnion   = To_Union(x.Id())
        /* noccur = 1 */
        Result = Equal(g0294.T1.Id(),C_any.Id())
        /* Let-3 */} 
      } else {
      Result = CFALSE
      /* If-2 */} 
    return Result} 
  
// The EID go function for: sort_abstract? @ type (throw: false) 
func E_Optimize_sort_abstract_ask_type (x EID) EID { 
    return EID{/*(sm for sort_abstract? @ type= boolean)*/ F_Optimize_sort_abstract_ask_type(ToType(OBJ(x)) ).Id(),0}} 
  
// since we introduce some fuzziness with types (any U t), we need a way to get
// the precise type t back
/* {1} OPT.The go function for: ptype(x:type) [] */
func F_Optimize_ptype_type (x *ClaireType ) *ClaireType  { 
    // procedure body with s =  
var Result *ClaireType  
    if (x.Isa.IsIn(C_Union) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0296 *ClaireUnion   = To_Union(x.Id())
        /* noccur = 3 */
        if (g0296.T1.Id() == C_any.Id()) /* If:4 */{ 
          Result = g0296.T2
          } else {
          Result = ToType(g0296.Id())
          /* If-4 */} 
        /* Let-3 */} 
      } else {
      Result = x
      /* If-2 */} 
    return Result} 
  
// The EID go function for: ptype @ type (throw: false) 
func E_Optimize_ptype_type (x EID) EID { 
    return EID{/*(sm for ptype @ type= type)*/ F_Optimize_ptype_type(ToType(OBJ(x)) ).Id(),0}} 
  
// v3.1.06: member -> always apply to a ptype
/* {1} OPT.The go function for: pmember(x:type) [] */
func F_Optimize_pmember_type (x *ClaireType ) *ClaireType  { 
    // use function body compiling 
return  Core.F_member_type(F_Optimize_ptype_type(x))
    } 
  
// The EID go function for: pmember @ type (throw: false) 
func E_Optimize_pmember_type (x EID) EID { 
    return EID{/*(sm for pmember @ type= type)*/ F_Optimize_pmember_type(ToType(OBJ(x)) ).Id(),0}} 
  
// transform an instruction representing a set into an instruction
// representing an enumeration
/* {1} OPT.The go function for: enumerate_code(self:any,%t:type) [] */
func F_Optimize_enumerate_code_any (self *ClaireAny ,_Zt *ClaireType ) EID { 
    var Result EID 
    if ((F_Optimize_ptype_type(_Zt).Included(ToType(C_list.Id())) == CTRUE) || 
        ((F_Optimize_ptype_type(_Zt).Included(ToType(C_set.Id())) == CTRUE) || 
          (F_Optimize_ptype_type(_Zt).Included(ToType(C_tuple.Id())) == CTRUE))) /* If:2 */{ 
      Result = F_Compile_c_strict_code_any(self,F_Optimize_ptype_type(_Zt).Class_I())
      } else {
      if (C_compiler.Optimize_ask == CTRUE) /* If:3 */{ 
        F_Compile_notice_void()
        Core.F_tformat_string(MakeString("explicit enmeration of ~S\n"),3,MakeConstantList(self))
        /* If-3 */} 
      Result = F_Optimize_c_code_method_method1(ToMethod(Core.F__at_property1(Core.C_Core_enumerate,C_any).Id()),MakeConstantList(self),MakeConstantList(_Zt.Id()))
      /* If-2 */} 
    return Result} 
  
// The EID go function for: enumerate_code @ any (throw: true) 
func E_Optimize_enumerate_code_any (self EID,_Zt EID) EID { 
    return /*(sm for enumerate_code @ any= EID)*/ F_Optimize_enumerate_code_any(ANY(self),ToType(OBJ(_Zt)) )} 
  
// range inference for a "for" structure: y is the new type and ts is the type of
// the collection structure. Note that except for the case of float arrays, the
// sort of the collection is assumed to be any or integer (thus we "correct" the
// type inference with sort_abstract)
/* {1} OPT.The go function for: range_infers_for(self:Variable,y:type,ts:type) [] */
func F_Optimize_range_infers_for_Variable (self *ClaireVariable ,y *ClaireType ,ts *ClaireType ) *ClaireAny  { 
    // procedure body with s =  
var Result *ClaireAny  
    if (self.Range.Id() == CNULL) /* If:2 */{ 
      
      if (y.Isa.IsIn(C_Interval) == CTRUE) /* If:3 */{ 
        y = ToType(C_integer.Id())
        /* If-3 */} 
      self.Range = y
      /* If!2 */}  else if ((y.Included(self.Range) != CTRUE) && 
        (C_compiler.Safety <= 1)) /* If:2 */{ 
      if ((F_boolean_I_any(y.Id()) != CTRUE) || 
          (F_boolean_I_any(self.Range.Id()) != CTRUE)) /* If:3 */{ 
        F_Compile_warn_void()
        Core.F_tformat_string(MakeString("range of variable in ~S is wrong [258]\n"),2,MakeConstantList(self.Id()))
        /* If-3 */} 
      /* If-2 */} 
    if ((F_sort_Variable(self).Id() != C_any.Id()) && 
        (((F_sort_Variable(self).Id() != C_integer.Id()) || 
            (C_compiler.Overflow_ask == CTRUE)) && 
          ((ts.Included(ToType(C_array.Id())) != CTRUE) || 
              (y.Included(ToType(C_float.Id())) != CTRUE)))) /* If:2 */{ 
      
      /* update:3 */{ 
        var va_arg1 *ClaireVariable  
        var va_arg2 *ClaireType  
        va_arg1 = self
        va_arg2 = F_Optimize_sort_abstract_I_type(self.Range)
        /* ---------- now we compile update range(va_arg1) := va_arg2 ------- */
        va_arg1.Range = va_arg2
        Result = va_arg2.Id()
        /* update-3 */} 
      } else {
      Result = CFALSE.Id()
      /* If-2 */} 
    return Result} 
  
// The EID go function for: range_infers_for @ Variable (throw: false) 
func E_Optimize_range_infers_for_Variable (self EID,y EID,ts EID) EID { 
    return /*(sm for range_infers_for @ Variable= any)*/ F_Optimize_range_infers_for_Variable(To_Variable(OBJ(self)),ToType(OBJ(y)),ToType(OBJ(ts)) ).ToEID()} 
  
// variable range inference, how to guess a type from the value ...
/* {1} OPT.The go function for: range_infers(self:Variable,y:type) [] */
func F_Optimize_range_infers_Variable (self *ClaireVariable ,y *ClaireType ) EID { 
    var Result EID 
    var g0300I *ClaireBoolean  
    var g0300I_try03012 EID 
    /* or:2 */{ 
      var v_or2 *ClaireBoolean  
      
      v_or2 = MakeBoolean((self.Range.Id() == CNULL))
      if (v_or2 == CTRUE) {g0300I_try03012 = EID{CTRUE.Id(),0}
      } else /* or:3 */{ 
        var v_or2_try03024 EID 
        v_or2_try03024 = F_Optimize_extended_ask_type(self.Range)
        /* ERROR PROTECTION INSERTED (v_or2-g0300I_try03012) */
        if ErrorIn(v_or2_try03024) {g0300I_try03012 = v_or2_try03024
        } else {
        v_or2 = ToBoolean(OBJ(v_or2_try03024))
        if (v_or2 == CTRUE) {g0300I_try03012 = EID{CTRUE.Id(),0}
        } else /* or:4 */{ 
          g0300I_try03012 = EID{CFALSE.Id(),0}/* org-4 */} 
        /* org-3 */} 
      }
      /* or-2 */} 
    /* ERROR PROTECTION INSERTED (g0300I-Result) */
    if ErrorIn(g0300I_try03012) {Result = g0300I_try03012
    } else {
    g0300I = ToBoolean(OBJ(g0300I_try03012))
    if (g0300I == CTRUE) /* If:2 */{ 
      if (C_set.Id() == y.Isa.Id()) /* If:3 */{ 
        /* update:4 */{ 
          var va_arg1 *ClaireVariable  
          var va_arg2 *ClaireType  
          va_arg1 = self
          va_arg2 = ToType(y.Class_I().Id())
          /* ---------- now we compile update range(va_arg1) := va_arg2 ------- */
          va_arg1.Range = va_arg2
          Result = EID{va_arg2.Id(),0}
          /* update-4 */} 
        } else {
        /* update:4 */{ 
          var va_arg1 *ClaireVariable  
          var va_arg2 *ClaireType  
          va_arg1 = self
          va_arg2 = y
          /* ---------- now we compile update range(va_arg1) := va_arg2 ------- */
          va_arg1.Range = va_arg2
          Result = EID{va_arg2.Id(),0}
          /* update-4 */} 
        /* If-3 */} 
      } else {
      Result = EID{CFALSE.Id(),0}
      /* If-2 */} 
    }
    return Result} 
  
// The EID go function for: range_infers @ Variable (throw: true) 
func E_Optimize_range_infers_Variable (self EID,y EID) EID { 
    return /*(sm for range_infers @ Variable= EID)*/ F_Optimize_range_infers_Variable(To_Variable(OBJ(self)),ToType(OBJ(y)) )} 
  
// temporary range inference for case, which may use a special form:
// {any U type} to represent the change of sort
// {} U (c U t) to represent a change of psort
/* {1} OPT.The go function for: range_infer_case(self:any,y:type) [] */
func F_Optimize_range_infer_case_any (self *ClaireAny ,y *ClaireType )  { 
    // procedure body with s =  
if (self.Isa.IsIn(C_Variable) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0303 *ClaireVariable   = To_Variable(self)
        /* noccur = 6 */
        if (F_boolean_I_any(F_Compile_sort_equal_class(F_Compile_osort_any(g0303.Range.Id()),F_Compile_osort_any(y.Id()))) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var c1 *ClaireClass   = F_Compile_psort_any(g0303.Range.Class_I().Id())
            /* noccur = 2 */
            if (c1.Id() != F_Compile_psort_any(y.Class_I().Id()).Id()) /* If:6 */{ 
              /* update:7 */{ 
                var va_arg1 *ClaireVariable  
                var va_arg2 *ClaireType  
                va_arg1 = g0303
                /* Let:8 */{ 
                  var _CL_obj *ClaireUnion   = To_Union(new(ClaireUnion).Is(C_Union))
                  /* noccur = 11 */
                  _CL_obj.T1 = ToType(CEMPTY.Id())
                  /* update:9 */{ 
                    var va_arg1 *ClaireUnion  
                    var va_arg2 *ClaireType  
                    va_arg1 = _CL_obj
                    /* Let:10 */{ 
                      var _CL_obj *ClaireUnion   = To_Union(new(ClaireUnion).Is(C_Union))
                      /* noccur = 5 */
                      _CL_obj.T1 = ToType(c1.Id())
                      _CL_obj.T2 = y
                      va_arg2 = ToType(_CL_obj.Id())
                      /* Let-10 */} 
                    /* ---------- now we compile update mClaire/t2(va_arg1) := va_arg2 ------- */
                    va_arg1.T2 = va_arg2
                    /* update-9 */} 
                  va_arg2 = ToType(_CL_obj.Id())
                  /* Let-8 */} 
                /* ---------- now we compile update range(va_arg1) := va_arg2 ------- */
                va_arg1.Range = va_arg2
                /* update-7 */} 
              } else {
              g0303.Range = y
              /* If-6 */} 
            /* Let-5 */} 
          /* If!4 */}  else if (F_Compile_osort_any(g0303.Range.Id()).Id() == C_any.Id()) /* If:4 */{ 
          g0303.Range = F_Optimize_sort_abstract_I_type(y)
          /* If-4 */} 
        /* Let-3 */} 
      /* If-2 */} 
    } 
  
// The EID go function for: range_infer_case @ any (throw: false) 
func E_Optimize_range_infer_case_any (self EID,y EID) EID { 
    /*(sm for range_infer_case @ any= void)*/ F_Optimize_range_infer_case_any(ANY(self),ToType(OBJ(y)) )
    return EVOID} 
  
//
// temporary range inference for case, which may use a special form:
// {any U type} to represent the change of sort
// {} U (c U t) to represent a change of psort
/* {1} OPT.The go function for: range_sets(self:any,y:type) [] */
func F_Optimize_range_sets_any (self *ClaireAny ,y *ClaireType )  { 
    // procedure body with s =  
if (self.Isa.IsIn(C_Variable) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0304 *ClaireVariable   = To_Variable(self)
        /* noccur = 6 */
        if (F_boolean_I_any(F_Compile_sort_equal_class(F_Compile_osort_any(g0304.Range.Id()),F_Compile_osort_any(y.Id()))) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var c1 *ClaireClass   = F_Compile_psort_any(g0304.Range.Class_I().Id())
            /* noccur = 2 */
            if (c1.Id() != F_Compile_psort_any(y.Class_I().Id()).Id()) /* If:6 */{ 
              /* update:7 */{ 
                var va_arg1 *ClaireVariable  
                var va_arg2 *ClaireType  
                va_arg1 = g0304
                /* Let:8 */{ 
                  var _CL_obj *ClaireUnion   = To_Union(new(ClaireUnion).Is(C_Union))
                  /* noccur = 11 */
                  _CL_obj.T1 = ToType(CEMPTY.Id())
                  /* update:9 */{ 
                    var va_arg1 *ClaireUnion  
                    var va_arg2 *ClaireType  
                    va_arg1 = _CL_obj
                    /* Let:10 */{ 
                      var _CL_obj *ClaireUnion   = To_Union(new(ClaireUnion).Is(C_Union))
                      /* noccur = 5 */
                      _CL_obj.T1 = ToType(c1.Id())
                      _CL_obj.T2 = y
                      va_arg2 = ToType(_CL_obj.Id())
                      /* Let-10 */} 
                    /* ---------- now we compile update mClaire/t2(va_arg1) := va_arg2 ------- */
                    va_arg1.T2 = va_arg2
                    /* update-9 */} 
                  va_arg2 = ToType(_CL_obj.Id())
                  /* Let-8 */} 
                /* ---------- now we compile update range(va_arg1) := va_arg2 ------- */
                va_arg1.Range = va_arg2
                /* update-7 */} 
              } else {
              g0304.Range = y
              /* If-6 */} 
            /* Let-5 */} 
          /* If!4 */}  else if (F_Compile_osort_any(g0304.Range.Id()).Id() == C_any.Id()) /* If:4 */{ 
          g0304.Range = F_Optimize_sort_abstract_I_type(y)
          /* If-4 */} 
        /* Let-3 */} 
      /* If-2 */} 
    } 
  
// The EID go function for: range_sets @ any (throw: false) 
func E_Optimize_range_sets_any (self EID,y EID) EID { 
    /*(sm for range_sets @ any= void)*/ F_Optimize_range_sets_any(ANY(self),ToType(OBJ(y)) )
    return EVOID} 
  
// the srange of a method = class!(range)
/* {1} OPT.The go function for: c_srange(m:method) [] */
func F_Optimize_c_srange_method (m *ClaireMethod ) EID { 
    var Result EID 
    Result = Core.F_last_list(m.Srange)
    return Result} 
  
// The EID go function for: c_srange @ method (throw: true) 
func E_Optimize_c_srange_method (m EID) EID { 
    return /*(sm for c_srange @ method= EID)*/ F_Optimize_c_srange_method(ToMethod(OBJ(m)) )} 
  
// v3.3 some of the global variables are compiled with a native var approach
// we require the range to be safe, no backtrack & local global var
/* {1} OPT.The go function for: Compile/nativeVar?(x:global_variable) [] */
func F_Compile_nativeVar_ask_global_variable (x *Core.GlobalVariable ) *ClaireBoolean  { 
    // use function body compiling 
return  MakeBoolean((C_compiler.Optimize_ask == CTRUE) && (x.Store_ask.Id() == CFALSE.Id()) && (x.Name.Module_I().Id() == ANY(Core.F_CALL(C_mClaire_definition,ARGS(EID{x.Name.Id(),0})))) && (F_boolean_I_any(ANY(Core.F_CALL(C_Optimize_gcsafe_ask,ARGS(EID{x.Range.Id(),0})))) == CTRUE))
    } 
  
// The EID go function for: Compile/nativeVar? @ global_variable (throw: false) 
func E_Compile_nativeVar_ask_global_variable (x EID) EID { 
    return EID{/*(sm for Compile/nativeVar? @ global_variable= boolean)*/ F_Compile_nativeVar_ask_global_variable(Core.ToGlobalVariable(OBJ(x)) ).Id(),0}} 
  
// v3.3 finds the possible return type of a block (within a loop)
// it returns a class for the time being ...
/* {1} OPT.The go function for: Compile/return_type(self:any) [] */
func F_Compile_return_type_any (self *ClaireAny ) EID { 
    var Result EID 
    if (self.Isa.IsIn(Language.C_Let) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0305 *Language.Let   = Language.To_Let(self)
        /* noccur = 1 */
        Result = F_Compile_return_type_any(g0305.Arg)
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(Language.C_Do) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0306 *Language.Do   = Language.To_Do(self)
        /* noccur = 1 */
        /* Let:4 */{ 
          var x *ClaireSet   = CEMPTY
          /* noccur = 3 */
          /* For:5 */{ 
            var y *ClaireAny  
            _ = y
            Result= EID{CFALSE.Id(),0}
            var y_support *ClaireList  
            y_support = g0306.Args
            y_len := y_support.Length()
            for i_it := 0; i_it < y_len; i_it++ { 
              y = y_support.At(i_it)
              var void_try7 EID 
              _ = void_try7
              var x_try03197 EID 
              /* Let:7 */{ 
                var g0320UU *ClaireType  
                /* noccur = 1 */
                var g0320UU_try03218 EID 
                g0320UU_try03218 = F_Compile_return_type_any(y)
                /* ERROR PROTECTION INSERTED (g0320UU-x_try03197) */
                if ErrorIn(g0320UU_try03218) {x_try03197 = g0320UU_try03218
                } else {
                g0320UU = ToType(OBJ(g0320UU_try03218))
                x_try03197 = Core.F_CALL(ToProperty(C__exp.Id()),ARGS(EID{x.Id(),0},EID{g0320UU.Id(),0}))
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (x-Result) */
              if ErrorIn(x_try03197) {Result = x_try03197
              Result = x_try03197
              break
              } else {
              x = ToSet(OBJ(x_try03197))
              void_try7 = EID{x.Id(),0}
              }
              /* loop-6 */} 
            /* For-5 */} 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          Result = EID{x.Id(),0}
          }
          /* Let-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(Language.C_If) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0307 *Language.If   = Language.To_If(self)
        /* noccur = 2 */
        /* Let:4 */{ 
          var g0322UU *ClaireType  
          /* noccur = 1 */
          var g0322UU_try03245 EID 
          g0322UU_try03245 = F_Compile_return_type_any(g0307.Arg)
          /* ERROR PROTECTION INSERTED (g0322UU-Result) */
          if ErrorIn(g0322UU_try03245) {Result = g0322UU_try03245
          } else {
          g0322UU = ToType(OBJ(g0322UU_try03245))
          /* Let:5 */{ 
            var g0323UU *ClaireType  
            /* noccur = 1 */
            var g0323UU_try03256 EID 
            g0323UU_try03256 = F_Compile_return_type_any(g0307.Other)
            /* ERROR PROTECTION INSERTED (g0323UU-Result) */
            if ErrorIn(g0323UU_try03256) {Result = g0323UU_try03256
            } else {
            g0323UU = ToType(OBJ(g0323UU_try03256))
            Result = Core.F_CALL(ToProperty(C__exp.Id()),ARGS(EID{g0322UU.Id(),0},EID{g0323UU.Id(),0}))
            }
            /* Let-5 */} 
          }
          /* Let-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(Language.C_Return) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0308 *Language.Return   = Language.To_Return(self)
        /* noccur = 1 */
        Result = Core.F_CALL(C_c_type,ARGS(g0308.Arg.ToEID()))
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(Language.C_Case) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0309 *Language.Case   = Language.To_Case(self)
        /* noccur = 1 */
        /* Let:4 */{ 
          var x *ClaireSet   = CEMPTY
          /* noccur = 3 */
          /* For:5 */{ 
            var y *ClaireAny  
            _ = y
            Result= EID{CFALSE.Id(),0}
            var y_support *ClaireList  
            y_support = g0309.Args
            y_len := y_support.Length()
            for i_it := 0; i_it < y_len; i_it++ { 
              y = y_support.At(i_it)
              var void_try7 EID 
              _ = void_try7
              var x_try03267 EID 
              /* Let:7 */{ 
                var g0327UU *ClaireType  
                /* noccur = 1 */
                var g0327UU_try03288 EID 
                g0327UU_try03288 = F_Compile_return_type_any(y)
                /* ERROR PROTECTION INSERTED (g0327UU-x_try03267) */
                if ErrorIn(g0327UU_try03288) {x_try03267 = g0327UU_try03288
                } else {
                g0327UU = ToType(OBJ(g0327UU_try03288))
                x_try03267 = Core.F_CALL(ToProperty(C__exp.Id()),ARGS(EID{x.Id(),0},EID{g0327UU.Id(),0}))
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (x-Result) */
              if ErrorIn(x_try03267) {Result = x_try03267
              Result = x_try03267
              break
              } else {
              x = ToSet(OBJ(x_try03267))
              void_try7 = EID{x.Id(),0}
              }
              /* loop-6 */} 
            /* For-5 */} 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          Result = EID{x.Id(),0}
          }
          /* Let-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(Language.C_Handle) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0310 *Language.ClaireHandle   = Language.To_ClaireHandle(self)
        /* noccur = 1 */
        Result = F_Compile_return_type_any(g0310.Arg)
        /* Let-3 */} 
      } else {
      Result = EID{CEMPTY.Id(),0}
      /* If-2 */} 
    return RangeCheck(ToType(C_type.Id()),Result)} 
  
// The EID go function for: Compile/return_type @ any (throw: true) 
func E_Compile_return_type_any (self EID) EID { 
    return /*(sm for Compile/return_type @ any= EID)*/ F_Compile_return_type_any(ANY(self) )} 
  
// compiling a type expression --------------------------------------------
//
// creates the functional code that produce the code by evaluation
// note this is expensive -> we should encourage the use of global variables
/* {1} OPT.The go function for: c_code(self:((type_operator U Reference) U Pattern),s:class) [] */
func F_c_code_type_expression (self *ClaireTypeExpression ,s *ClaireClass ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var g0329UU *ClaireAny  
      /* noccur = 1 */
      var g0329UU_try03303 EID 
      g0329UU_try03303 = Core.F_CALL(C_Compile_self_code,ARGS(EID{self.Id(),0}))
      /* ERROR PROTECTION INSERTED (g0329UU-Result) */
      if ErrorIn(g0329UU_try03303) {Result = g0329UU_try03303
      } else {
      g0329UU = ANY(g0329UU_try03303)
      Result = Core.F_CALL(C_c_code,ARGS(g0329UU.ToEID(),EID{s.Id(),0}))
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_code @ type_expression (throw: true) 
func E_c_code_type_expression (self EID,s EID) EID { 
    return /*(sm for c_code @ type_expression= EID)*/ F_c_code_type_expression(ToTypeExpression(OBJ(self)),ToClass(OBJ(s)) )} 
  
// to check - seems OK for 3.2 !
/* {1} OPT.The go function for: Compile/self_code(self:subtype) [] */
func F_Compile_self_code_subtype (self *ClaireSubtype ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
      /* noccur = 5 */
      _CL_obj.Selector = C_nth
      /* update:3 */{ 
        var va_arg1 *Language.Call  
        var va_arg2 *ClaireList  
        va_arg1 = _CL_obj
        var va_arg2_try03314 EID 
        /* Construct:4 */{ 
          var v_bag_arg *ClaireAny  
          va_arg2_try03314= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
          ToList(OBJ(va_arg2_try03314)).AddFast(self.Arg.Id())
          var v_bag_arg_try03325 EID 
          v_bag_arg_try03325 = Core.F_CALL(C_c_code,ARGS(EID{self.T1.Id(),0},EID{C_type.Id(),0}))
          /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try03314) */
          if ErrorIn(v_bag_arg_try03325) {va_arg2_try03314 = v_bag_arg_try03325
          } else {
          v_bag_arg = ANY(v_bag_arg_try03325)
          ToList(OBJ(va_arg2_try03314)).AddFast(v_bag_arg)}
          /* Construct-4 */} 
        /* ERROR PROTECTION INSERTED (va_arg2-Result) */
        if ErrorIn(va_arg2_try03314) {Result = va_arg2_try03314
        } else {
        va_arg2 = ToList(OBJ(va_arg2_try03314))
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
  
// The EID go function for: Compile/self_code @ subtype (throw: true) 
func E_Compile_self_code_subtype (self EID) EID { 
    return /*(sm for Compile/self_code @ subtype= EID)*/ F_Compile_self_code_subtype(ToSubtype(OBJ(self)) )} 
  
// create a Param. Optimized in v3.2.28 for list<X>
/* {1} OPT.The go function for: Compile/self_code(self:Param) [] */
func F_Compile_self_code_Param (self *ClaireParam ) EID { 
    var Result EID 
    if ((self.Params.Length() == 1) && 
        ((self.Params.At(1-1) == C_of.Id()) && 
          (C_set.Id() == self.Args.At(1-1).Isa.Id()))) /* If:2 */{ 
      /* Let:3 */{ 
        var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
        /* noccur = 5 */
        _CL_obj.Selector = Core.C_Core_param_I
        /* update:4 */{ 
          var va_arg1 *Language.Call  
          var va_arg2 *ClaireList  
          va_arg1 = _CL_obj
          var va_arg2_try03335 EID 
          /* Construct:5 */{ 
            var v_bag_arg *ClaireAny  
            va_arg2_try03335= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
            ToList(OBJ(va_arg2_try03335)).AddFast(self.Arg.Id())
            var v_bag_arg_try03346 EID 
            /* Let:6 */{ 
              var g0335UU *ClaireAny  
              /* noccur = 1 */
              var g0335UU_try03367 EID 
              g0335UU_try03367 = Core.F_the_type(ToType(self.Args.At(1-1)))
              /* ERROR PROTECTION INSERTED (g0335UU-v_bag_arg_try03346) */
              if ErrorIn(g0335UU_try03367) {v_bag_arg_try03346 = g0335UU_try03367
              } else {
              g0335UU = ANY(g0335UU_try03367)
              v_bag_arg_try03346 = Core.F_CALL(C_c_code,ARGS(g0335UU.ToEID(),EID{C_type.Id(),0}))
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try03335) */
            if ErrorIn(v_bag_arg_try03346) {va_arg2_try03335 = v_bag_arg_try03346
            } else {
            v_bag_arg = ANY(v_bag_arg_try03346)
            ToList(OBJ(va_arg2_try03335)).AddFast(v_bag_arg)}
            /* Construct-5 */} 
          /* ERROR PROTECTION INSERTED (va_arg2-Result) */
          if ErrorIn(va_arg2_try03335) {Result = va_arg2_try03335
          } else {
          va_arg2 = ToList(OBJ(va_arg2_try03335))
          /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
          va_arg1.Args = va_arg2
          Result = EID{va_arg2.Id(),0}
          }
          /* update-4 */} 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        Result = EID{_CL_obj.Id(),0}
        }
        /* Let-3 */} 
      } else {
      /* Let:3 */{ 
        var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
        /* noccur = 5 */
        _CL_obj.Selector = C_nth
        /* update:4 */{ 
          var va_arg1 *Language.Call  
          var va_arg2 *ClaireList  
          va_arg1 = _CL_obj
          var va_arg2_try03375 EID 
          /* Construct:5 */{ 
            var v_bag_arg *ClaireAny  
            va_arg2_try03375= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
            ToList(OBJ(va_arg2_try03375)).AddFast(self.Arg.Id())
            ToList(OBJ(va_arg2_try03375)).AddFast(self.Params.Id())
            var v_bag_arg_try03386 EID 
            /* Iteration:6 */{ 
              var v_list6 *ClaireList  
              var y *ClaireAny  
              var v_local6 *ClaireAny  
              v_list6 = self.Args
              v_bag_arg_try03386 = EID{CreateList(ToType(CEMPTY.Id()),v_list6.Length()).Id(),0}
              for CLcount := 0; CLcount < v_list6.Length(); CLcount++{ 
                y = v_list6.At(CLcount)
                var v_local6_try03398 EID 
                v_local6_try03398 = Core.F_CALL(C_c_code,ARGS(y.ToEID(),EID{C_type.Id(),0}))
                /* ERROR PROTECTION INSERTED (v_local6-v_bag_arg_try03386) */
                if ErrorIn(v_local6_try03398) {v_bag_arg_try03386 = v_local6_try03398
                v_bag_arg_try03386 = v_local6_try03398
                break
                } else {
                v_local6 = ANY(v_local6_try03398)
                ToList(OBJ(v_bag_arg_try03386)).PutAt(CLcount,v_local6)
                } 
              }
              /* Iteration-6 */} 
            /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try03375) */
            if ErrorIn(v_bag_arg_try03386) {va_arg2_try03375 = v_bag_arg_try03386
            } else {
            v_bag_arg = ANY(v_bag_arg_try03386)
            ToList(OBJ(va_arg2_try03375)).AddFast(v_bag_arg)}
            /* Construct-5 */} 
          /* ERROR PROTECTION INSERTED (va_arg2-Result) */
          if ErrorIn(va_arg2_try03375) {Result = va_arg2_try03375
          } else {
          va_arg2 = ToList(OBJ(va_arg2_try03375))
          /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
          va_arg1.Args = va_arg2
          Result = EID{va_arg2.Id(),0}
          }
          /* update-4 */} 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        Result = EID{_CL_obj.Id(),0}
        }
        /* Let-3 */} 
      /* If-2 */} 
    return Result} 
  
// The EID go function for: Compile/self_code @ Param (throw: true) 
func E_Compile_self_code_Param (self EID) EID { 
    return /*(sm for Compile/self_code @ Param= EID)*/ F_Compile_self_code_Param(To_Param(OBJ(self)) )} 
  
/* {1} OPT.The go function for: Compile/self_code(self:Union) [] */
func F_Compile_self_code_Union (self *ClaireUnion ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
      /* noccur = 5 */
      _CL_obj.Selector = ToProperty(Core.C_U.Id())
      /* update:3 */{ 
        var va_arg1 *Language.Call  
        var va_arg2 *ClaireList  
        va_arg1 = _CL_obj
        var va_arg2_try03404 EID 
        /* Construct:4 */{ 
          var v_bag_arg *ClaireAny  
          va_arg2_try03404= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
          var v_bag_arg_try03415 EID 
          v_bag_arg_try03415 = Core.F_CALL(C_c_code,ARGS(EID{self.T1.Id(),0},EID{C_type.Id(),0}))
          /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try03404) */
          if ErrorIn(v_bag_arg_try03415) {va_arg2_try03404 = v_bag_arg_try03415
          } else {
          v_bag_arg = ANY(v_bag_arg_try03415)
          ToList(OBJ(va_arg2_try03404)).AddFast(v_bag_arg)
          var v_bag_arg_try03425 EID 
          v_bag_arg_try03425 = Core.F_CALL(C_c_code,ARGS(EID{self.T2.Id(),0},EID{C_type.Id(),0}))
          /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try03404) */
          if ErrorIn(v_bag_arg_try03425) {va_arg2_try03404 = v_bag_arg_try03425
          } else {
          v_bag_arg = ANY(v_bag_arg_try03425)
          ToList(OBJ(va_arg2_try03404)).AddFast(v_bag_arg)}}
          /* Construct-4 */} 
        /* ERROR PROTECTION INSERTED (va_arg2-Result) */
        if ErrorIn(va_arg2_try03404) {Result = va_arg2_try03404
        } else {
        va_arg2 = ToList(OBJ(va_arg2_try03404))
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
  
// The EID go function for: Compile/self_code @ Union (throw: true) 
func E_Compile_self_code_Union (self EID) EID { 
    return /*(sm for Compile/self_code @ Union= EID)*/ F_Compile_self_code_Union(To_Union(OBJ(self)) )} 
  
/* {1} OPT.The go function for: Compile/self_code(self:Interval) [] */
func F_Compile_self_code_Interval (self *ClaireInterval ) *ClaireAny  { 
    // procedure body with s =  
var Result *ClaireAny  
    /* Let:2 */{ 
      var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
      /* noccur = 5 */
      _CL_obj.Selector = ToProperty(C__dot_dot.Id())
      _CL_obj.Args = MakeConstantList(MakeInteger(self.Arg1).Id(),MakeInteger(self.Arg2).Id())
      Result = _CL_obj.Id()
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: Compile/self_code @ Interval (throw: false) 
func E_Compile_self_code_Interval (self EID) EID { 
    return /*(sm for Compile/self_code @ Interval= any)*/ F_Compile_self_code_Interval(To_Interval(OBJ(self)) ).ToEID()} 
  
/* {1} OPT.The go function for: Compile/self_code(self:Reference) [] */
func F_Compile_self_code_Reference (self *ClaireReference ) *ClaireAny  { 
    // procedure body with s =  
var Result *ClaireAny  
    /* Let:2 */{ 
      var _CL_obj *Language.Definition   = Language.To_Definition(new(Language.Definition).Is(Language.C_Definition))
      /* noccur = 23 */
      _CL_obj.Arg = C_Reference
      /* update:3 */{ 
        var va_arg1 *Language.Definition  
        var va_arg2 *ClaireList  
        va_arg1 = _CL_obj
        /* Construct:4 */{ 
          var v_bag_arg *ClaireAny  
          va_arg2= ToType(CEMPTY.Id()).EmptyList()
          /* Let:5 */{ 
            var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
            /* noccur = 5 */
            _CL_obj.Selector = ToProperty(C__equal.Id())
            _CL_obj.Args = MakeConstantList(C_args.Id(),self.Args.Id())
            v_bag_arg = _CL_obj.Id()
            /* Let-5 */} 
          va_arg2.AddFast(v_bag_arg)
          /* Let:5 */{ 
            var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
            /* noccur = 5 */
            _CL_obj.Selector = ToProperty(C__equal.Id())
            _CL_obj.Args = MakeConstantList(C_mClaire_index.Id(),MakeInteger(self.Index).Id())
            v_bag_arg = _CL_obj.Id()
            /* Let-5 */} 
          va_arg2.AddFast(v_bag_arg)
          /* Let:5 */{ 
            var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
            /* noccur = 5 */
            _CL_obj.Selector = ToProperty(C__equal.Id())
            _CL_obj.Args = MakeConstantList(C_arg.Id(),self.Arg.Id())
            v_bag_arg = _CL_obj.Id()
            /* Let-5 */} 
          va_arg2.AddFast(v_bag_arg)/* Construct-4 */} 
        /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
        va_arg1.Args = va_arg2
        /* update-3 */} 
      Result = _CL_obj.Id()
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: Compile/self_code @ Reference (throw: false) 
func E_Compile_self_code_Reference (self EID) EID { 
    return /*(sm for Compile/self_code @ Reference= any)*/ F_Compile_self_code_Reference(To_Reference(OBJ(self)) ).ToEID()} 
  
// compilation of a Pattern
/* {1} OPT.The go function for: Compile/self_code(self:Pattern) [] */
func (self *ClairePattern ) SelfCode () *ClaireAny  { 
    // procedure body with s =  
var Result *ClaireAny  
    if (C_compiler.Inline_ask == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
        /* noccur = 9 */
        _CL_obj.Selector = C_nth
        /* update:4 */{ 
          var va_arg1 *Language.Call  
          var va_arg2 *ClaireList  
          va_arg1 = _CL_obj
          /* Construct:5 */{ 
            var v_bag_arg *ClaireAny  
            va_arg2= ToType(CEMPTY.Id()).EmptyList()
            va_arg2.AddFast(self.Selector.Id())
            /* Let:6 */{ 
              var _CL_obj *Language.Tuple   = Language.To_Tuple(new(Language.Tuple).Is(Language.C_Tuple))
              /* noccur = 3 */
              _CL_obj.Args = self.Arg
              v_bag_arg = _CL_obj.Id()
              /* Let-6 */} 
            va_arg2.AddFast(v_bag_arg)/* Construct-5 */} 
          /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
          va_arg1.Args = va_arg2
          /* update-4 */} 
        Result = _CL_obj.Id()
        /* Let-3 */} 
      } else {
      Result = Language.C_Call.Id()
      /* If-2 */} 
    return Result} 
  
// The EID go function for: Compile/self_code @ Pattern (throw: false) 
func E_Compile_self_code_Pattern (self EID) EID { 
    return /*(sm for Compile/self_code @ Pattern= any)*/ To_ClairePattern(OBJ(self)).SelfCode( ).ToEID()} 
  
//-------------- membership compiling -------------------------------
// membership to a class : for final and closed classes => x.isa = c !
/* {1} OPT.The go function for: member_code(self:class,x:any) [] */
func F_Optimize_member_code_class (self *ClaireClass ,x *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var _Zxt *Language.Call  
      /* noccur = 2 */
      var _Zxt_try03433 EID 
      /* Let:3 */{ 
        var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
        /* noccur = 5 */
        /* update:4 */{ 
          var va_arg1 *Language.Call  
          var va_arg2 *ClaireProperty  
          va_arg1 = _CL_obj
          var va_arg2_try03445 EID 
          var g0345I *ClaireBoolean  
          var g0345I_try03465 EID 
          /* Let:5 */{ 
            var g0347UU *ClaireType  
            /* noccur = 1 */
            var g0347UU_try03486 EID 
            g0347UU_try03486 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
            /* ERROR PROTECTION INSERTED (g0347UU-g0345I_try03465) */
            if ErrorIn(g0347UU_try03486) {g0345I_try03465 = g0347UU_try03486
            } else {
            g0347UU = ToType(OBJ(g0347UU_try03486))
            g0345I_try03465 = EID{g0347UU.Included(ToType(C_object.Id())).Id(),0}
            }
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (g0345I-va_arg2_try03445) */
          if ErrorIn(g0345I_try03465) {va_arg2_try03445 = g0345I_try03465
          } else {
          g0345I = ToBoolean(OBJ(g0345I_try03465))
          if (g0345I == CTRUE) /* If:5 */{ 
            va_arg2_try03445 = EID{C_isa.Id(),0}
            } else {
            va_arg2_try03445 = EID{Core.C_owner.Id(),0}
            /* If-5 */} 
          }
          /* ERROR PROTECTION INSERTED (va_arg2-_Zxt_try03433) */
          if ErrorIn(va_arg2_try03445) {_Zxt_try03433 = va_arg2_try03445
          } else {
          va_arg2 = ToProperty(OBJ(va_arg2_try03445))
          /* ---------- now we compile update selector(va_arg1) := va_arg2 ------- */
          va_arg1.Selector = va_arg2
          _Zxt_try03433 = EID{va_arg2.Id(),0}
          }
          /* update-4 */} 
        /* ERROR PROTECTION INSERTED (_Zxt_try03433-_Zxt_try03433) */
        if !ErrorIn(_Zxt_try03433) {
        _CL_obj.Args = MakeConstantList(x)
        _Zxt_try03433 = EID{_CL_obj.Id(),0}
        }
        /* Let-3 */} 
      /* ERROR PROTECTION INSERTED (_Zxt-Result) */
      if ErrorIn(_Zxt_try03433) {Result = _Zxt_try03433
      } else {
      _Zxt = Language.To_Call(OBJ(_Zxt_try03433))
      if (((self.Open <= -1) || 
            (self.Open == 1)) && 
          (F_boolean_I_any(self.Subclass.Id()).Id() != CTRUE.Id())) /* If:3 */{ 
        /* Let:4 */{ 
          var g0349UU *Language.Call  
          /* noccur = 1 */
          /* Let:5 */{ 
            var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
            /* noccur = 5 */
            _CL_obj.Selector = ToProperty(C__equal.Id())
            _CL_obj.Args = MakeConstantList(self.Id(),_Zxt.Id())
            g0349UU = _CL_obj
            /* Let-5 */} 
          Result = Core.F_CALL(C_c_code,ARGS(EID{g0349UU.Id(),0}))
          /* Let-4 */} 
        } else {
        /* Let:4 */{ 
          var g0350UU *Language.Call  
          /* noccur = 1 */
          /* Let:5 */{ 
            var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
            /* noccur = 5 */
            _CL_obj.Selector = ToProperty(Core.C_inherit_ask.Id())
            _CL_obj.Args = MakeConstantList(_Zxt.Id(),self.Id())
            g0350UU = _CL_obj
            /* Let-5 */} 
          Result = Core.F_CALL(C_c_code,ARGS(EID{g0350UU.Id(),0}))
          /* Let-4 */} 
        /* If-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: member_code @ class (throw: true) 
func E_Optimize_member_code_class (self EID,x EID) EID { 
    return /*(sm for member_code @ class= EID)*/ F_Optimize_member_code_class(ToClass(OBJ(self)),ANY(x) )} 
  
/* {1} OPT.The go function for: member_code(self:type_operator,x:any) [] */
func F_Optimize_member_code_type_operator (self *ClaireTypeOperator ,x *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var _CL_obj *Language.CallMethod2   = Language.To_CallMethod2(new(Language.CallMethod2).Is(Language.C_Call_method2))
      /* noccur = 5 */
      _CL_obj.Arg = ToMethod(Core.F__at_property2(ToProperty(C__Z.Id()),MakeConstantList(C_any.Id(),C_any.Id())).Id())
      /* update:3 */{ 
        var va_arg1 *Language.CallMethod  
        var va_arg2 *ClaireList  
        va_arg1 = Language.To_CallMethod(_CL_obj.Id())
        var va_arg2_try03514 EID 
        /* Construct:4 */{ 
          var v_bag_arg *ClaireAny  
          va_arg2_try03514= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
          var v_bag_arg_try03525 EID 
          v_bag_arg_try03525 = Core.F_CALL(C_c_code,ARGS(x.ToEID(),EID{C_any.Id(),0}))
          /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try03514) */
          if ErrorIn(v_bag_arg_try03525) {va_arg2_try03514 = v_bag_arg_try03525
          } else {
          v_bag_arg = ANY(v_bag_arg_try03525)
          ToList(OBJ(va_arg2_try03514)).AddFast(v_bag_arg)
          var v_bag_arg_try03535 EID 
          v_bag_arg_try03535 = Core.F_CALL(C_c_code,ARGS(EID{self.Id(),0},EID{C_any.Id(),0}))
          /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try03514) */
          if ErrorIn(v_bag_arg_try03535) {va_arg2_try03514 = v_bag_arg_try03535
          } else {
          v_bag_arg = ANY(v_bag_arg_try03535)
          ToList(OBJ(va_arg2_try03514)).AddFast(v_bag_arg)}}
          /* Construct-4 */} 
        /* ERROR PROTECTION INSERTED (va_arg2-Result) */
        if ErrorIn(va_arg2_try03514) {Result = va_arg2_try03514
        } else {
        va_arg2 = ToList(OBJ(va_arg2_try03514))
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
  
// The EID go function for: member_code @ type_operator (throw: true) 
func E_Optimize_member_code_type_operator (self EID,x EID) EID { 
    return /*(sm for member_code @ type_operator= EID)*/ F_Optimize_member_code_type_operator(ToTypeOperator(OBJ(self)),ANY(x) )} 
  
/* {1} OPT.The go function for: member_code(self:Union,x:any) [] */
func F_Optimize_member_code_Union (self *ClaireUnion ,x *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var _CL_obj *Language.Or   = Language.To_Or(new(Language.Or).Is(Language.C_Or))
      /* noccur = 3 */
      /* update:3 */{ 
        var va_arg1 *Language.Or  
        var va_arg2 *ClaireList  
        va_arg1 = _CL_obj
        var va_arg2_try03544 EID 
        /* Construct:4 */{ 
          var v_bag_arg *ClaireAny  
          va_arg2_try03544= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
          var v_bag_arg_try03555 EID 
          v_bag_arg_try03555 = Core.F_CALL(C_Optimize_member_code,ARGS(EID{self.T1.Id(),0},x.ToEID()))
          /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try03544) */
          if ErrorIn(v_bag_arg_try03555) {va_arg2_try03544 = v_bag_arg_try03555
          } else {
          v_bag_arg = ANY(v_bag_arg_try03555)
          ToList(OBJ(va_arg2_try03544)).AddFast(v_bag_arg)
          var v_bag_arg_try03565 EID 
          v_bag_arg_try03565 = Core.F_CALL(C_Optimize_member_code,ARGS(EID{self.T2.Id(),0},x.ToEID()))
          /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try03544) */
          if ErrorIn(v_bag_arg_try03565) {va_arg2_try03544 = v_bag_arg_try03565
          } else {
          v_bag_arg = ANY(v_bag_arg_try03565)
          ToList(OBJ(va_arg2_try03544)).AddFast(v_bag_arg)}}
          /* Construct-4 */} 
        /* ERROR PROTECTION INSERTED (va_arg2-Result) */
        if ErrorIn(va_arg2_try03544) {Result = va_arg2_try03544
        } else {
        va_arg2 = ToList(OBJ(va_arg2_try03544))
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
  
// The EID go function for: member_code @ Union (throw: true) 
func E_Optimize_member_code_Union (self EID,x EID) EID { 
    return /*(sm for member_code @ Union= EID)*/ F_Optimize_member_code_Union(To_Union(OBJ(self)),ANY(x) )} 
  
/* {1} OPT.The go function for: member_code(self:Interval,x:any) [] */
func F_Optimize_member_code_Interval (self *ClaireInterval ,x *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var g0357UU *Language.And  
      /* noccur = 1 */
      /* Let:3 */{ 
        var _CL_obj *Language.And   = Language.To_And(new(Language.And).Is(Language.C_And))
        /* noccur = 15 */
        /* update:4 */{ 
          var va_arg1 *Language.And  
          var va_arg2 *ClaireList  
          va_arg1 = _CL_obj
          /* Construct:5 */{ 
            var v_bag_arg *ClaireAny  
            va_arg2= ToType(CEMPTY.Id()).EmptyList()
            /* Let:6 */{ 
              var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
              /* noccur = 5 */
              _CL_obj.Selector = ToProperty(C__sup_equal.Id())
              _CL_obj.Args = MakeConstantList(x,MakeInteger(self.Arg1).Id())
              v_bag_arg = _CL_obj.Id()
              /* Let-6 */} 
            va_arg2.AddFast(v_bag_arg)
            /* Let:6 */{ 
              var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
              /* noccur = 5 */
              _CL_obj.Selector = ToProperty(C__inf_equal.Id())
              _CL_obj.Args = MakeConstantList(x,MakeInteger(self.Arg2).Id())
              v_bag_arg = _CL_obj.Id()
              /* Let-6 */} 
            va_arg2.AddFast(v_bag_arg)/* Construct-5 */} 
          /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
          va_arg1.Args = va_arg2
          /* update-4 */} 
        g0357UU = _CL_obj
        /* Let-3 */} 
      Result = Core.F_CALL(C_c_code,ARGS(EID{g0357UU.Id(),0},EID{C_any.Id(),0}))
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: member_code @ Interval (throw: true) 
func E_Optimize_member_code_Interval (self EID,x EID) EID { 
    return /*(sm for member_code @ Interval= EID)*/ F_Optimize_member_code_Interval(To_Interval(OBJ(self)),ANY(x) )} 
  
/* {1} OPT.The go function for: member_code(self:Param,x:any) [] */
func F_Optimize_member_code_Param (self *ClaireParam ,x *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var g0359UU *Language.And  
      /* noccur = 1 */
      /* Let:3 */{ 
        var _CL_obj *Language.And   = Language.To_And(new(Language.And).Is(Language.C_And))
        /* noccur = 21 */
        /* update:4 */{ 
          var va_arg1 *Language.And  
          var va_arg2 *ClaireList  
          va_arg1 = _CL_obj
          /* Let:5 */{ 
            var g0360UU *ClaireList  
            /* noccur = 1 */
            /* Construct:6 */{ 
              var v_bag_arg *ClaireAny  
              g0360UU= ToType(CEMPTY.Id()).EmptyList()
              /* Let:7 */{ 
                var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                /* noccur = 5 */
                _CL_obj.Selector = ToProperty(C__Z.Id())
                _CL_obj.Args = MakeConstantList(x,self.Arg.Id())
                v_bag_arg = _CL_obj.Id()
                /* Let-7 */} 
              g0360UU.AddFast(v_bag_arg)/* Construct-6 */} 
            /* Let:6 */{ 
              var g0361UU *ClaireList  
              /* noccur = 1 */
              /* Let:7 */{ 
                var i_bag *ClaireList   = ToType(CEMPTY.Id()).EmptyList()
                /* noccur = 2 */
                /* Let:8 */{ 
                  var i int  = 1
                  /* noccur = 5 */
                  /* Let:9 */{ 
                    var g0358 int  = self.Params.Length()
                    /* noccur = 1 */
                    for (i <= g0358) /* while:10 */{ 
                      /* Let:11 */{ 
                        var g0362UU *Language.Call  
                        /* noccur = 1 */
                        /* Let:12 */{ 
                          var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                          /* noccur = 11 */
                          _CL_obj.Selector = ToProperty(C__Z.Id())
                          /* update:13 */{ 
                            var va_arg1 *Language.Call  
                            var va_arg2 *ClaireList  
                            va_arg1 = _CL_obj
                            /* Construct:14 */{ 
                              var v_bag_arg *ClaireAny  
                              va_arg2= ToType(CEMPTY.Id()).EmptyList()
                              /* Let:15 */{ 
                                var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                /* noccur = 5 */
                                _CL_obj.Selector = ToProperty(self.Params.At(i-1))
                                _CL_obj.Args = MakeConstantList(x)
                                v_bag_arg = _CL_obj.Id()
                                /* Let-15 */} 
                              va_arg2.AddFast(v_bag_arg)
                              va_arg2.AddFast(self.Args.At(i-1))/* Construct-14 */} 
                            /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                            va_arg1.Args = va_arg2
                            /* update-13 */} 
                          g0362UU = _CL_obj
                          /* Let-12 */} 
                        i_bag.AddFast(g0362UU.Id())
                        /* Let-11 */} 
                      i = (i+1)
                      /* while-10 */} 
                    /* Let-9 */} 
                  /* Let-8 */} 
                g0361UU = i_bag
                /* Let-7 */} 
              va_arg2 = g0360UU.Append(g0361UU)
              /* Let-6 */} 
            /* Let-5 */} 
          /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
          va_arg1.Args = va_arg2
          /* update-4 */} 
        g0359UU = _CL_obj
        /* Let-3 */} 
      Result = Core.F_CALL(C_c_code,ARGS(EID{g0359UU.Id(),0},EID{C_any.Id(),0}))
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: member_code @ Param (throw: true) 
func E_Optimize_member_code_Param (self EID,x EID) EID { 
    return /*(sm for member_code @ Param= EID)*/ F_Optimize_member_code_Param(To_Param(OBJ(self)),ANY(x) )} 
  
// v3.3.14: specialized code for tuple
/* {1} OPT.The go function for: member_code(self:tuple,x:any) [] */
func F_Optimize_member_code_tuple (self *ClaireTuple ,x *ClaireAny ) EID { 
    var Result EID 
    if (x.Isa.IsIn(Language.C_Tuple) == CTRUE) /* If:2 */{ 
      if (ToList(OBJ(Core.F_CALL(C_args,ARGS(x.ToEID())))).Length() != self.Length()) /* If:3 */{ 
        Result = EID{CFALSE.Id(),0}
        } else {
        /* Let:4 */{ 
          var g0364UU *Language.And  
          /* noccur = 1 */
          /* Let:5 */{ 
            var _CL_obj *Language.And   = Language.To_And(new(Language.And).Is(Language.C_And))
            /* noccur = 9 */
            /* update:6 */{ 
              var va_arg1 *Language.And  
              var va_arg2 *ClaireList  
              va_arg1 = _CL_obj
              /* Let:7 */{ 
                var i_bag *ClaireList   = ToType(CEMPTY.Id()).EmptyList()
                /* noccur = 2 */
                /* Let:8 */{ 
                  var i int  = 1
                  /* noccur = 5 */
                  /* Let:9 */{ 
                    var g0363 int  = self.Length()
                    /* noccur = 1 */
                    for (i <= g0363) /* while:10 */{ 
                      /* Let:11 */{ 
                        var g0365UU *Language.Call  
                        /* noccur = 1 */
                        /* Let:12 */{ 
                          var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                          /* noccur = 5 */
                          _CL_obj.Selector = ToProperty(C__Z.Id())
                          _CL_obj.Args = MakeConstantList(ToList(OBJ(Core.F_CALL(C_args,ARGS(x.ToEID())))).At(i-1),ToList(self.Id()).At(i-1))
                          g0365UU = _CL_obj
                          /* Let-12 */} 
                        i_bag.AddFast(g0365UU.Id())
                        /* Let-11 */} 
                      i = (i+1)
                      /* while-10 */} 
                    /* Let-9 */} 
                  /* Let-8 */} 
                va_arg2 = i_bag
                /* Let-7 */} 
              /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
              va_arg1.Args = va_arg2
              /* update-6 */} 
            g0364UU = _CL_obj
            /* Let-5 */} 
          Result = Core.F_CALL(C_c_code,ARGS(EID{g0364UU.Id(),0},EID{C_any.Id(),0}))
          /* Let-4 */} 
        /* If-3 */} 
      } else {
      Result = F_Optimize_c_code_method_method1(ToMethod(Core.F__at_property2(ToProperty(C__Z.Id()),MakeConstantList(C_any.Id(),C_any.Id())).Id()),MakeConstantList(x,self.Id()),MakeConstantList(C_any.Id(),C_any.Id()))
      /* If-2 */} 
    return Result} 
  
// The EID go function for: member_code @ tuple (throw: true) 
func E_Optimize_member_code_tuple (self EID,x EID) EID { 
    return /*(sm for member_code @ tuple= EID)*/ F_Optimize_member_code_tuple(ToTuple(OBJ(self)),ANY(x) )} 
  
/* {1} OPT.The go function for: member_code(self:any,x:any) [] */
func F_Optimize_member_code_any (self *ClaireAny ,x *ClaireAny ) EID { 
    var Result EID 
    Language.C_LDEF.Value = CNIL.Id()
    /* Let:2 */{ 
      var _Ztype *ClaireList  
      /* noccur = 2 */
      var _Ztype_try03673 EID 
      /* Construct:3 */{ 
        var v_bag_arg *ClaireAny  
        _Ztype_try03673= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
        var v_bag_arg_try03684 EID 
        v_bag_arg_try03684 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
        /* ERROR PROTECTION INSERTED (v_bag_arg-_Ztype_try03673) */
        if ErrorIn(v_bag_arg_try03684) {_Ztype_try03673 = v_bag_arg_try03684
        } else {
        v_bag_arg = ANY(v_bag_arg_try03684)
        ToList(OBJ(_Ztype_try03673)).AddFast(v_bag_arg)
        var v_bag_arg_try03694 EID 
        v_bag_arg_try03694 = Core.F_CALL(C_c_type,ARGS(self.ToEID()))
        /* ERROR PROTECTION INSERTED (v_bag_arg-_Ztype_try03673) */
        if ErrorIn(v_bag_arg_try03694) {_Ztype_try03673 = v_bag_arg_try03694
        } else {
        v_bag_arg = ANY(v_bag_arg_try03694)
        ToList(OBJ(_Ztype_try03673)).AddFast(v_bag_arg)}}
        /* Construct-3 */} 
      /* ERROR PROTECTION INSERTED (_Ztype-Result) */
      if ErrorIn(_Ztype_try03673) {Result = _Ztype_try03673
      } else {
      _Ztype = ToList(OBJ(_Ztype_try03673))
      /* Let:3 */{ 
        var r *ClaireAny  
        /* noccur = 2 */
        var r_try03704 EID 
        r_try03704 = Language.F_extract_pattern_any(self,CNIL)
        /* ERROR PROTECTION INSERTED (r-Result) */
        if ErrorIn(r_try03704) {Result = r_try03704
        } else {
        r = ANY(r_try03704)
        var g0371I *ClaireBoolean  
        /* or:4 */{ 
          var v_or4 *ClaireBoolean  
          
          v_or4 = Equal(r,CNULL)
          if (v_or4 == CTRUE) {g0371I = CTRUE
          } else /* or:5 */{ 
            v_or4 = Equal(self,C_object.Id())
            if (v_or4 == CTRUE) {g0371I = CTRUE
            } else /* or:6 */{ 
              /* Let:7 */{ 
                var g0372UU *ClaireObject  
                /* noccur = 1 */
                if (self.Isa.IsIn(Core.C_global_variable) == CTRUE) /* If:8 */{ 
                  /* Let:9 */{ 
                    var g0366 *Core.GlobalVariable   = Core.ToGlobalVariable(self)
                    /* noccur = 1 */
                    g0372UU = ToObject(g0366.Range.Id())
                    /* Let-9 */} 
                  } else {
                  g0372UU = ToObject(CFALSE.Id())
                  /* If-8 */} 
                v_or4 = F_boolean_I_any(g0372UU.Id())
                /* Let-7 */} 
              if (v_or4 == CTRUE) {g0371I = CTRUE
              } else /* or:7 */{ 
                g0371I = CFALSE/* org-7 */} 
              /* org-6 */} 
            /* org-5 */} 
          /* or-4 */} 
        if (g0371I == CTRUE) /* If:4 */{ 
          Result = F_Optimize_c_code_method_method1(ToMethod(Core.F__at_property2(ToProperty(C__Z.Id()),_Ztype).Id()),MakeConstantList(x,self),_Ztype)
          } else {
          Result = Core.F_CALL(C_Optimize_member_code,ARGS(r.ToEID(),x.ToEID()))
          /* If-4 */} 
        }
        /* Let-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: member_code @ any (throw: true) 
func E_Optimize_member_code_any (self EID,x EID) EID { 
    return /*(sm for member_code @ any= EID)*/ F_Optimize_member_code_any(ANY(self),ANY(x) )} 
  
// membership optimization though inline definition of %
/* {1} OPT.The go function for: %(x:any,y:any) [] */
func F__Z_any4 (x *ClaireAny ,y *ClaireAny ) EID { 
    var Result EID 
    /* and:2 */{ 
      var v_and2 *ClaireBoolean  
      
      var v_and2_try03733 EID 
      /* Let:3 */{ 
        var g0374UU *ClaireAny  
        /* noccur = 1 */
        var g0374UU_try03754 EID 
        g0374UU_try03754 = EVAL(ToList(OBJ(Core.F_CALL(C_args,ARGS(y.ToEID())))).At(2-1))
        /* ERROR PROTECTION INSERTED (g0374UU-v_and2_try03733) */
        if ErrorIn(g0374UU_try03754) {v_and2_try03733 = g0374UU_try03754
        } else {
        g0374UU = ANY(g0374UU_try03754)
        v_and2_try03733 = Core.F_CALL(ToProperty(C__inf_equal.Id()),ARGS(x.ToEID(),g0374UU.ToEID()))
        }
        /* Let-3 */} 
      /* ERROR PROTECTION INSERTED (v_and2-Result) */
      if ErrorIn(v_and2_try03733) {Result = v_and2_try03733
      } else {
      v_and2 = ToBoolean(OBJ(v_and2_try03733))
      if (v_and2 == CFALSE) {Result = EID{CFALSE.Id(),0}
      } else /* arg:3 */{ 
        var v_and2_try03764 EID 
        /* Let:4 */{ 
          var g0377UU *ClaireAny  
          /* noccur = 1 */
          var g0377UU_try03785 EID 
          g0377UU_try03785 = EVAL(ToList(OBJ(Core.F_CALL(C_args,ARGS(y.ToEID())))).At(1-1))
          /* ERROR PROTECTION INSERTED (g0377UU-v_and2_try03764) */
          if ErrorIn(g0377UU_try03785) {v_and2_try03764 = g0377UU_try03785
          } else {
          g0377UU = ANY(g0377UU_try03785)
          v_and2_try03764 = Core.F_CALL(ToProperty(C__inf_equal.Id()),ARGS(g0377UU.ToEID(),x.ToEID()))
          }
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (v_and2-Result) */
        if ErrorIn(v_and2_try03764) {Result = v_and2_try03764
        } else {
        v_and2 = ToBoolean(OBJ(v_and2_try03764))
        if (v_and2 == CFALSE) {Result = EID{CFALSE.Id(),0}
        } else /* arg:4 */{ 
          Result = EID{CTRUE.Id(),0}/* arg-4 */} 
        /* arg-3 */} 
      }}
      /* and-2 */} 
    return Result} 
  
// The EID go function for: % @ list<type_expression>(any, ..[tuple(any,any)]) (throw: true) 
func E__Z_any4 (x EID,y EID) EID { 
    return /*(sm for % @ list<type_expression>(any, ..[tuple(any,any)])= EID)*/ F__Z_any4(ANY(x),ANY(y) )} 
  
/* {1} OPT.The go function for: %(x:any,y:any) [] */
func F__Z_any5 (x *ClaireAny ,y *ClaireAny ) EID { 
    var Result EID 
    /* and:2 */{ 
      var v_and2 *ClaireBoolean  
      
      var v_and2_try03793 EID 
      /* Let:3 */{ 
        var g0380UU *ClaireAny  
        /* noccur = 1 */
        var g0380UU_try03814 EID 
        g0380UU_try03814 = EVAL(ToList(OBJ(Core.F_CALL(C_args,ARGS(y.ToEID())))).At(1-1))
        /* ERROR PROTECTION INSERTED (g0380UU-v_and2_try03793) */
        if ErrorIn(g0380UU_try03814) {v_and2_try03793 = g0380UU_try03814
        } else {
        g0380UU = ANY(g0380UU_try03814)
        v_and2_try03793 = Core.F_BELONG(x,g0380UU)
        }
        /* Let-3 */} 
      /* ERROR PROTECTION INSERTED (v_and2-Result) */
      if ErrorIn(v_and2_try03793) {Result = v_and2_try03793
      } else {
      v_and2 = ToBoolean(OBJ(v_and2_try03793))
      if (v_and2 == CFALSE) {Result = EID{CFALSE.Id(),0}
      } else /* arg:3 */{ 
        var v_and2_try03824 EID 
        /* Let:4 */{ 
          var g0383UU *ClaireAny  
          /* noccur = 1 */
          var g0383UU_try03845 EID 
          g0383UU_try03845 = EVAL(ToList(OBJ(Core.F_CALL(C_args,ARGS(y.ToEID())))).At(2-1))
          /* ERROR PROTECTION INSERTED (g0383UU-v_and2_try03824) */
          if ErrorIn(g0383UU_try03845) {v_and2_try03824 = g0383UU_try03845
          } else {
          g0383UU = ANY(g0383UU_try03845)
          v_and2_try03824 = EID{Core.F__I_equal_any(x,g0383UU).Id(),0}
          }
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (v_and2-Result) */
        if ErrorIn(v_and2_try03824) {Result = v_and2_try03824
        } else {
        v_and2 = ToBoolean(OBJ(v_and2_try03824))
        if (v_and2 == CFALSE) {Result = EID{CFALSE.Id(),0}
        } else /* arg:4 */{ 
          Result = EID{CTRUE.Id(),0}/* arg-4 */} 
        /* arg-3 */} 
      }}
      /* and-2 */} 
    return Result} 
  
// The EID go function for: % @ list<type_expression>(any, but[tuple(any,any)]) (throw: true) 
func E__Z_any5 (x EID,y EID) EID { 
    return /*(sm for % @ list<type_expression>(any, but[tuple(any,any)])= EID)*/ F__Z_any5(ANY(x),ANY(y) )} 
  
// ******************************************************************
// *    Part 4: Miscellaneous                                       *
// ******************************************************************
// ------- variables ------------------------------------------------
/* {1} OPT.The go function for: Compile/Variable!(s:symbol,n:integer,t:any) [] */
func F_Compile_Variable_I_symbol (s *ClaireSymbol ,n int,t *ClaireAny ) *ClaireVariable  { 
    // procedure body with s =  
var Result *ClaireVariable  
    if (t.Isa.IsIn(C_type) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0385 *ClaireType   = ToType(t)
        /* noccur = 1 */
        /* Let:4 */{ 
          var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
          /* noccur = 7 */
          _CL_obj.Pname = s
          _CL_obj.Index = n
          _CL_obj.Range = g0385
          Result = _CL_obj
          /* Let-4 */} 
        /* Let-3 */} 
      } else {
      /* Let:3 */{ 
        var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
        /* noccur = 5 */
        _CL_obj.Pname = s
        _CL_obj.Index = n
        Result = _CL_obj
        /* Let-3 */} 
      /* If-2 */} 
    return Result} 
  
// The EID go function for: Compile/Variable! @ symbol (throw: false) 
func E_Compile_Variable_I_symbol (s EID,n EID,t EID) EID { 
    return EID{/*(sm for Compile/Variable! @ symbol= Variable)*/ F_Compile_Variable_I_symbol(ToSymbol(OBJ(s)),INT(n),ANY(t) ).Id(),0}} 
  
/* {1} OPT.The go function for: Compile/get_indexed(c:class) [] */
func F_Compile_get_indexed_class (c *ClaireClass ) *ClaireList  { 
    // use function body compiling 
return  c.Slots
    } 
  
// The EID go function for: Compile/get_indexed @ class (throw: false) 
func E_Compile_get_indexed_class (c EID) EID { 
    return EID{/*(sm for Compile/get_indexed @ class= list)*/ F_Compile_get_indexed_class(ToClass(OBJ(c)) ).Id(),0}} 
  
// simple C operations that can be duplicated at no cost {+, -, /, *}
// tells if an expression is a C simply designated object
/* {1} OPT.The go function for: Compile/designated?(self:any) [] */
func F_Compile_designated_ask_any (self *ClaireAny ) EID { 
    var Result EID 
    /* or:2 */{ 
      var v_or2 *ClaireBoolean  
      
      v_or2 = self.Isa.IsIn(C_thing)
      if (v_or2 == CTRUE) {Result = EID{CTRUE.Id(),0}
      } else /* or:3 */{ 
        v_or2 = self.Isa.IsIn(C_Variable)
        if (v_or2 == CTRUE) {Result = EID{CTRUE.Id(),0}
        } else /* or:4 */{ 
          v_or2 = Equal(C_integer.Id(),self.Isa.Id())
          if (v_or2 == CTRUE) {Result = EID{CTRUE.Id(),0}
          } else /* or:5 */{ 
            v_or2 = Equal(C_boolean.Id(),self.Isa.Id())
            if (v_or2 == CTRUE) {Result = EID{CTRUE.Id(),0}
            } else /* or:6 */{ 
              v_or2 = Equal(self,CNIL.Id())
              if (v_or2 == CTRUE) {Result = EID{CTRUE.Id(),0}
              } else /* or:7 */{ 
                v_or2 = Equal(self,CEMPTY.Id())
                if (v_or2 == CTRUE) {Result = EID{CTRUE.Id(),0}
                } else /* or:8 */{ 
                  v_or2 = Equal(self,CNULL)
                  if (v_or2 == CTRUE) {Result = EID{CTRUE.Id(),0}
                  } else /* or:9 */{ 
                    v_or2 = Equal(C_float.Id(),self.Isa.Id())
                    if (v_or2 == CTRUE) {Result = EID{CTRUE.Id(),0}
                    } else /* or:10 */{ 
                      var v_or2_try039211 EID 
                      if (self.Isa.IsIn(Language.C_Call) == CTRUE) /* If:11 */{ 
                        /* Let:12 */{ 
                          var g0387 *Language.Call   = Language.To_Call(self)
                          /* noccur = 2 */
                          /* Let:13 */{ 
                            var x *ClaireAny  
                            /* noccur = 2 */
                            var x_try039314 EID 
                            x_try039314 = Core.F_CALL(C_c_code,ARGS(EID{g0387.Id(),0}))
                            /* ERROR PROTECTION INSERTED (x-v_or2_try039211) */
                            if ErrorIn(x_try039314) {v_or2_try039211 = x_try039314
                            } else {
                            x = ANY(x_try039314)
                            /* or:14 */{ 
                              var v_or14 *ClaireBoolean  
                              
                              var v_or14_try039415 EID 
                              /* and:15 */{ 
                                var v_and15 *ClaireBoolean  
                                
                                v_and15 = x.Isa.IsIn(Language.C_Call).Not
                                if (v_and15 == CFALSE) {v_or14_try039415 = EID{CFALSE.Id(),0}
                                } else /* arg:16 */{ 
                                  var v_and15_try039517 EID 
                                  v_and15_try039517 = F_Compile_designated_ask_any(x)
                                  /* ERROR PROTECTION INSERTED (v_and15-v_or14_try039415) */
                                  if ErrorIn(v_and15_try039517) {v_or14_try039415 = v_and15_try039517
                                  } else {
                                  v_and15 = ToBoolean(OBJ(v_and15_try039517))
                                  if (v_and15 == CFALSE) {v_or14_try039415 = EID{CFALSE.Id(),0}
                                  } else /* arg:17 */{ 
                                    v_or14_try039415 = EID{CTRUE.Id(),0}/* arg-17 */} 
                                  /* arg-16 */} 
                                }
                                /* and-15 */} 
                              /* ERROR PROTECTION INSERTED (v_or14-v_or2_try039211) */
                              if ErrorIn(v_or14_try039415) {v_or2_try039211 = v_or14_try039415
                              } else {
                              v_or14 = ToBoolean(OBJ(v_or14_try039415))
                              if (v_or14 == CTRUE) {v_or2_try039211 = EID{CTRUE.Id(),0}
                              } else /* or:15 */{ 
                                v_or14 = Equal(g0387.Selector.Id(),Core.C_mClaire_get_stack.Id())
                                if (v_or14 == CTRUE) {v_or2_try039211 = EID{CTRUE.Id(),0}
                                } else /* or:16 */{ 
                                  v_or2_try039211 = EID{CFALSE.Id(),0}/* org-16 */} 
                                /* org-15 */} 
                              }
                              /* or-14 */} 
                            }
                            /* Let-13 */} 
                          /* Let-12 */} 
                        /* If!11 */}  else if (self.Isa.IsIn(Language.C_Call_slot) == CTRUE) /* If:11 */{ 
                        /* Let:12 */{ 
                          var g0388 *Language.CallSlot   = Language.To_CallSlot(self)
                          /* noccur = 1 */
                          v_or2_try039211 = F_Compile_designated_ask_any(g0388.Arg)
                          /* Let-12 */} 
                        /* If!11 */}  else if (self.Isa.IsIn(Language.C_Call_table) == CTRUE) /* If:11 */{ 
                        /* Let:12 */{ 
                          var g0389 *Language.CallTable   = Language.To_CallTable(self)
                          /* noccur = 1 */
                          v_or2_try039211 = F_Compile_designated_ask_any(g0389.Arg)
                          /* Let-12 */} 
                        /* If!11 */}  else if (self.Isa.IsIn(Language.C_Call_method) == CTRUE) /* If:11 */{ 
                        /* Let:12 */{ 
                          var g0390 *Language.CallMethod   = Language.To_CallMethod(self)
                          /* noccur = 3 */
                          /* and:13 */{ 
                            var v_and13 *ClaireBoolean  
                            
                            v_and13 = MakeBoolean((C_OPT.SimpleOperations.Contain_ask(g0390.Arg.Selector.Id()) == CTRUE) || (g0390.Arg.Id() == Core.F__at_property1(C_nth,C_list).Id()))
                            if (v_and13 == CFALSE) {v_or2_try039211 = EID{CFALSE.Id(),0}
                            } else /* arg:14 */{ 
                              var v_and13_try039615 EID 
                              /* Let:15 */{ 
                                var g0397UU *ClaireAny  
                                /* noccur = 1 */
                                var g0397UU_try039816 EID 
                                /* For:16 */{ 
                                  var y *ClaireAny  
                                  _ = y
                                  g0397UU_try039816= EID{CFALSE.Id(),0}
                                  var y_support *ClaireList  
                                  y_support = g0390.Args
                                  y_len := y_support.Length()
                                  for i_it := 0; i_it < y_len; i_it++ { 
                                    y = y_support.At(i_it)
                                    var void_try18 EID 
                                    _ = void_try18
                                    var g0399I *ClaireBoolean  
                                    var g0399I_try040018 EID 
                                    /* Let:18 */{ 
                                      var g0401UU *ClaireBoolean  
                                      /* noccur = 1 */
                                      var g0401UU_try040219 EID 
                                      g0401UU_try040219 = F_Compile_designated_ask_any(y)
                                      /* ERROR PROTECTION INSERTED (g0401UU-g0399I_try040018) */
                                      if ErrorIn(g0401UU_try040219) {g0399I_try040018 = g0401UU_try040219
                                      } else {
                                      g0401UU = ToBoolean(OBJ(g0401UU_try040219))
                                      g0399I_try040018 = EID{g0401UU.Not.Id(),0}
                                      }
                                      /* Let-18 */} 
                                    /* ERROR PROTECTION INSERTED (g0399I-void_try18) */
                                    if ErrorIn(g0399I_try040018) {void_try18 = g0399I_try040018
                                    } else {
                                    g0399I = ToBoolean(OBJ(g0399I_try040018))
                                    if (g0399I == CTRUE) /* If:18 */{ 
                                       /*v = g0397UU_try039816, s =EID*/
g0397UU_try039816 = EID{CTRUE.Id(),0}
                                      break
                                      } else {
                                      void_try18 = EID{CFALSE.Id(),0}
                                      /* If-18 */} 
                                    }
                                    /* ERROR PROTECTION INSERTED (void_try18-g0397UU_try039816) */
                                    if ErrorIn(void_try18) {g0397UU_try039816 = void_try18
                                    g0397UU_try039816 = void_try18
                                    break
                                    } else {
                                    }
                                    /* loop-17 */} 
                                  /* For-16 */} 
                                /* ERROR PROTECTION INSERTED (g0397UU-v_and13_try039615) */
                                if ErrorIn(g0397UU_try039816) {v_and13_try039615 = g0397UU_try039816
                                } else {
                                g0397UU = ANY(g0397UU_try039816)
                                v_and13_try039615 = EID{Core.F_not_any(g0397UU).Id(),0}
                                }
                                /* Let-15 */} 
                              /* ERROR PROTECTION INSERTED (v_and13-v_or2_try039211) */
                              if ErrorIn(v_and13_try039615) {v_or2_try039211 = v_and13_try039615
                              } else {
                              v_and13 = ToBoolean(OBJ(v_and13_try039615))
                              if (v_and13 == CFALSE) {v_or2_try039211 = EID{CFALSE.Id(),0}
                              } else /* arg:15 */{ 
                                v_or2_try039211 = EID{CTRUE.Id(),0}/* arg-15 */} 
                              /* arg-14 */} 
                            }
                            /* and-13 */} 
                          /* Let-12 */} 
                        } else {
                        v_or2_try039211 = EID{CFALSE.Id(),0}
                        /* If-11 */} 
                      /* ERROR PROTECTION INSERTED (v_or2-Result) */
                      if ErrorIn(v_or2_try039211) {Result = v_or2_try039211
                      } else {
                      v_or2 = ToBoolean(OBJ(v_or2_try039211))
                      if (v_or2 == CTRUE) {Result = EID{CTRUE.Id(),0}
                      } else /* or:11 */{ 
                        Result = EID{CFALSE.Id(),0}/* org-11 */} 
                      /* org-10 */} 
                    /* org-9 */} 
                  /* org-8 */} 
                /* org-7 */} 
              /* org-6 */} 
            /* org-5 */} 
          /* org-4 */} 
        /* org-3 */} 
      }
      /* or-2 */} 
    return Result} 
  
// The EID go function for: Compile/designated? @ any (throw: true) 
func E_Compile_designated_ask_any (self EID) EID { 
    return /*(sm for Compile/designated? @ any= EID)*/ F_Compile_designated_ask_any(ANY(self) )} 
  
// OPT.non_identifiable_set: those sets who are identifiable (closure)
// set<class>{c in class | exists(c2 in c.descendents | c2.ident? = false)})
// equality is identity?
/* {1} OPT.The go function for: Compile/identifiable?(self:any) [] */
func F_Compile_identifiable_ask_any (self *ClaireAny ) EID { 
    var Result EID 
    /* or:2 */{ 
      var v_or2 *ClaireBoolean  
      
      v_or2 = Equal(self,CNULL)
      if (v_or2 == CTRUE) {Result = EID{CTRUE.Id(),0}
      } else /* or:3 */{ 
        var v_or2_try04034 EID 
        /* Let:4 */{ 
          var t *ClaireClass  
          /* noccur = 1 */
          var t_try04045 EID 
          /* Let:5 */{ 
            var g0405UU *ClaireType  
            /* noccur = 1 */
            var g0405UU_try04066 EID 
            g0405UU_try04066 = Core.F_CALL(C_c_type,ARGS(self.ToEID()))
            /* ERROR PROTECTION INSERTED (g0405UU-t_try04045) */
            if ErrorIn(g0405UU_try04066) {t_try04045 = g0405UU_try04066
            } else {
            g0405UU = ToType(OBJ(g0405UU_try04066))
            t_try04045 = EID{g0405UU.Class_I().Id(),0}
            }
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (t-v_or2_try04034) */
          if ErrorIn(t_try04045) {v_or2_try04034 = t_try04045
          } else {
          t = ToClass(OBJ(t_try04045))
          v_or2_try04034 = EID{C_OPT.NonIdentifiableSet.Contain_ask(t.Id()).Not.Id(),0}
          }
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (v_or2-Result) */
        if ErrorIn(v_or2_try04034) {Result = v_or2_try04034
        } else {
        v_or2 = ToBoolean(OBJ(v_or2_try04034))
        if (v_or2 == CTRUE) {Result = EID{CTRUE.Id(),0}
        } else /* or:4 */{ 
          Result = EID{CFALSE.Id(),0}/* org-4 */} 
        /* org-3 */} 
      }
      /* or-2 */} 
    return Result} 
  
// The EID go function for: Compile/identifiable? @ any (throw: true) 
func E_Compile_identifiable_ask_any (self EID) EID { 
    return /*(sm for Compile/identifiable? @ any= EID)*/ F_Compile_identifiable_ask_any(ANY(self) )} 
  
// inlinning ---------------------------------------------------------
// macro expansion of method self with argument list l
/* {1} OPT.The go function for: c_inline(self:method,l:list,s:class) [] */
func F_Optimize_c_inline_method1 (self *ClaireMethod ,l *ClaireList ,s *ClaireClass ) EID { 
    var Result EID 
    Core.F_tformat_string(MakeString("macroexpansion of ~S with method ~S \n"),0,MakeConstantList(l.Id(),self.Id()))
    /* Let:2 */{ 
      var g0411UU *ClaireAny  
      /* noccur = 1 */
      var g0411UU_try04123 EID 
      g0411UU_try04123 = F_Optimize_c_inline_method2(self,l)
      /* ERROR PROTECTION INSERTED (g0411UU-Result) */
      if ErrorIn(g0411UU_try04123) {Result = g0411UU_try04123
      } else {
      g0411UU = ANY(g0411UU_try04123)
      Result = Core.F_CALL(C_c_code,ARGS(g0411UU.ToEID(),EID{s.Id(),0}))
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_inline @ list<type_expression>(method, list, class) (throw: true) 
func E_Optimize_c_inline_method1 (self EID,l EID,s EID) EID { 
    return /*(sm for c_inline @ list<type_expression>(method, list, class)= EID)*/ F_Optimize_c_inline_method1(ToMethod(OBJ(self)),ToList(OBJ(l)),ToClass(OBJ(s)) )} 
  
// apply the body of a macro definition
// notice that the name of the inner variables is changed except the second variable
// of iterate macros    
/* {1} OPT.The go function for: c_inline(self:method,l:list) [] */
func F_Optimize_c_inline_method2 (self *ClaireMethod ,l *ClaireList ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var f *ClaireLambda   = self.Formula
      /* noccur = 4 */
      /* Let:3 */{ 
        var x *ClaireAny   = f.Body
        /* noccur = 7 */
        /* Let:4 */{ 
          var lbv *ClaireList  
          /* noccur = 3 */
          var lbv_try04145 EID 
          lbv_try04145 = F_Optimize_bound_variables_any(x)
          /* ERROR PROTECTION INSERTED (lbv-Result) */
          if ErrorIn(lbv_try04145) {Result = lbv_try04145
          } else {
          lbv = ToList(OBJ(lbv_try04145))
          /* Let:5 */{ 
            var pv0 *ClaireAny  
            /* noccur = 2 */
            if ((self.Selector.Id() == Language.C_iterate.Id()) || 
                (self.Selector.Id() == Language.C_Iterate.Id())) /* If:6 */{ 
              pv0 = ANY(Core.F_CALL(C_mClaire_pname,ARGS(f.Vars.At(2-1).ToEID())))
              } else {
              pv0 = C_class.Name.Id()
              /* If-6 */} 
            x = Language.F_instruction_copy_any(x)
            Core.F_tformat_string(MakeString("c_inline(~S) on ~S: ~S is bound : ~S \n"),0,MakeConstantList(self.Id(),
              l.Id(),
              lbv.Id(),
              x))
            /* For:6 */{ 
              var v *ClaireAny  
              _ = v
              var v_support *ClaireList  
              v_support = lbv
              v_len := v_support.Length()
              for i_it := 0; i_it < v_len; i_it++ { 
                v = v_support.At(i_it)
                /* Let:8 */{ 
                  var v2 *ClaireVariable  
                  /* noccur = 2 */
                  /* Let:9 */{ 
                    var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
                    /* noccur = 5 */
                    /* update:10 */{ 
                      var va_arg1 *ClaireVariable  
                      var va_arg2 *ClaireSymbol  
                      va_arg1 = _CL_obj
                      if (ANY(Core.F_CALL(C_mClaire_pname,ARGS(v.ToEID()))) == pv0) /* If:11 */{ 
                        va_arg2 = ToSymbol(pv0)
                        } else {
                        va_arg2 = Core.F_gensym_void()
                        /* If-11 */} 
                      /* ---------- now we compile update mClaire/pname(va_arg1) := va_arg2 ------- */
                      va_arg1.Pname = va_arg2
                      /* update-10 */} 
                    _CL_obj.Index = 1000
                    v2 = _CL_obj
                    /* Let-9 */} 
                  v2.Range = ToType(Core.F_get_property(C_range,ToObject(v)))
                  x = Language.F_substitution_any(x,To_Variable(v),v2.Id())
                  /* Let-8 */} 
                /* loop-7 */} 
              /* For-6 */} 
            C_OPT.MaxVars = (C_OPT.MaxVars+lbv.Length())
            Core.F_tformat_string(MakeString("substitute f.vars = ~S with l = ~S \n"),0,MakeConstantList(f.Vars.Id(),l.Id()))
            Result = F_Optimize_c_substitution_any(x,f.Vars,l,CFALSE)
            /* Let-5 */} 
          }
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_inline @ list<type_expression>(method, list) (throw: true) 
func E_Optimize_c_inline_method2 (self EID,l EID) EID { 
    return /*(sm for c_inline @ list<type_expression>(method, list)= EID)*/ F_Optimize_c_inline_method2(ToMethod(OBJ(self)),ToList(OBJ(l)) )} 
  
// returns the macro expanded code if a macro is involved and nil otherwise
/* {1} OPT.The go function for: c_inline_arg?(self:any) [] */
func F_Optimize_c_inline_arg_ask_any (self *ClaireAny ) EID { 
    var Result EID 
    if (self.Isa.IsIn(Language.C_Call) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0422 *Language.Call   = Language.To_Call(self)
        /* noccur = 2 */
        /* Let:4 */{ 
          var l *ClaireList   = g0422.Args
          /* noccur = 3 */
          /* Let:5 */{ 
            var m *ClaireAny  
            /* noccur = 2 */
            var m_try04266 EID 
            /* Let:6 */{ 
              var g0427UU *ClaireList  
              /* noccur = 1 */
              var g0427UU_try04287 EID 
              /* Iteration:7 */{ 
                var v_list7 *ClaireList  
                var x *ClaireAny  
                var v_local7 *ClaireAny  
                v_list7 = l
                g0427UU_try04287 = EID{CreateList(ToType(CEMPTY.Id()),v_list7.Length()).Id(),0}
                for CLcount := 0; CLcount < v_list7.Length(); CLcount++{ 
                  x = v_list7.At(CLcount)
                  var v_local7_try04299 EID 
                  v_local7_try04299 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
                  /* ERROR PROTECTION INSERTED (v_local7-g0427UU_try04287) */
                  if ErrorIn(v_local7_try04299) {g0427UU_try04287 = v_local7_try04299
                  g0427UU_try04287 = v_local7_try04299
                  break
                  } else {
                  v_local7 = ANY(v_local7_try04299)
                  ToList(OBJ(g0427UU_try04287)).PutAt(CLcount,v_local7)
                  } 
                }
                /* Iteration-7 */} 
              /* ERROR PROTECTION INSERTED (g0427UU-m_try04266) */
              if ErrorIn(g0427UU_try04287) {m_try04266 = g0427UU_try04287
              } else {
              g0427UU = ToList(OBJ(g0427UU_try04287))
              m_try04266 = F_Optimize_restriction_I_property(g0422.Selector,g0427UU,CTRUE).ToEID()
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (m-Result) */
            if ErrorIn(m_try04266) {Result = m_try04266
            } else {
            m = ANY(m_try04266)
            if (C_method.Id() == m.Isa.Id()) /* If:6 */{ 
              /* Let:7 */{ 
                var g0423 *ClaireMethod   = ToMethod(m)
                /* noccur = 3 */
                var g0431I *ClaireBoolean  
                var g0431I_try04328 EID 
                /* and:8 */{ 
                  var v_and8 *ClaireBoolean  
                  
                  v_and8 = g0423.Inline_ask
                  if (v_and8 == CFALSE) {g0431I_try04328 = EID{CFALSE.Id(),0}
                  } else /* arg:9 */{ 
                    var v_and8_try043310 EID 
                    v_and8_try043310 = F_Optimize_c_inline_ask_method(g0423,l)
                    /* ERROR PROTECTION INSERTED (v_and8-g0431I_try04328) */
                    if ErrorIn(v_and8_try043310) {g0431I_try04328 = v_and8_try043310
                    } else {
                    v_and8 = ToBoolean(OBJ(v_and8_try043310))
                    if (v_and8 == CFALSE) {g0431I_try04328 = EID{CFALSE.Id(),0}
                    } else /* arg:10 */{ 
                      g0431I_try04328 = EID{CTRUE.Id(),0}/* arg-10 */} 
                    /* arg-9 */} 
                  }
                  /* and-8 */} 
                /* ERROR PROTECTION INSERTED (g0431I-Result) */
                if ErrorIn(g0431I_try04328) {Result = g0431I_try04328
                } else {
                g0431I = ToBoolean(OBJ(g0431I_try04328))
                if (g0431I == CTRUE) /* If:8 */{ 
                  Result = F_Optimize_c_inline_method2(g0423,l)
                  } else {
                  Result = EID{CFALSE.Id(),0}
                  /* If-8 */} 
                }
                /* Let-7 */} 
              } else {
              Result = EID{CNIL.Id(),0}
              /* If-6 */} 
            }
            /* Let-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      } else {
      /* Let:3 */{ 
        var g0438UU *Language.Call  
        /* noccur = 1 */
        /* Let:4 */{ 
          var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
          /* noccur = 5 */
          _CL_obj.Selector = C_set_I
          _CL_obj.Args = MakeConstantList(self)
          g0438UU = _CL_obj
          /* Let-4 */} 
        Result = F_Optimize_c_inline_arg_ask_any(g0438UU.Id())
        /* Let-3 */} 
      /* If-2 */} 
    return Result} 
  
// The EID go function for: c_inline_arg? @ any (throw: true) 
func E_Optimize_c_inline_arg_ask_any (self EID) EID { 
    return /*(sm for c_inline_arg? @ any= EID)*/ F_Optimize_c_inline_arg_ask_any(ANY(self) )} 
  
// substitute any variable with same name as x with the value val. val is an expression
// when the special form eval() is found, it is "evaluated"
// NEW: in v3.0.5 -> eval(x,C) evals only if x is actually a C
/* {1} OPT.The go function for: c_substitution(self:any,lx:list[Variable],val:list,eval?:boolean) [] */
func F_Optimize_c_substitution_any (self *ClaireAny ,lx *ClaireList ,val *ClaireList ,eval_ask *ClaireBoolean ) EID { 
    var Result EID 
    if (self.Isa.IsIn(C_Variable) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0439 *ClaireVariable   = To_Variable(self)
        /* noccur = 2 */
        /* Let:4 */{ 
          var i *ClaireAny  
          /* noccur = 2 */
          /* Let:5 */{ 
            var j_some *ClaireAny   = CNULL
            /* noccur = 2 */
            /* Let:6 */{ 
              var j int  = 1
              /* noccur = 5 */
              /* Let:7 */{ 
                var g0440 int  = lx.Length()
                /* noccur = 1 */
                for (j <= g0440) /* while:8 */{ 
                  if (g0439.Pname.Id() == To_Variable(lx.At(j-1)).Pname.Id()) /* If:9 */{ 
                     /*v = i, s =void*/
j_some = MakeInteger(j).Id()
                    break
                    /* If-9 */} 
                  j = (j+1)
                  /* while-8 */} 
                /* Let-7 */} 
              /* Let-6 */} 
            i = j_some
            /* Let-5 */} 
          if (i != CNULL) /* If:5 */{ 
            Result = val.At(ToInteger(i).Value-1).ToEID()
            } else {
            Result = EID{g0439.Id(),0}
            /* If-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(C_list) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0441 *ClaireList   = ToList(self)
        /* noccur = 4 */
        /* Let:4 */{ 
          var i int  = 1
          /* noccur = 5 */
          /* Let:5 */{ 
            var g0442 int  = g0441.Length()
            /* noccur = 1 */
            Result= EID{CFALSE.Id(),0}
            for (i <= g0442) /* while:6 */{ 
              var void_try7 EID 
              _ = void_try7
              { 
              /* Let:7 */{ 
                var g0446UU *ClaireAny  
                /* noccur = 1 */
                var g0446UU_try04478 EID 
                g0446UU_try04478 = F_Optimize_c_substitution_any(g0441.At(i-1),lx,val,eval_ask)
                /* ERROR PROTECTION INSERTED (g0446UU-void_try7) */
                if ErrorIn(g0446UU_try04478) {void_try7 = g0446UU_try04478
                } else {
                g0446UU = ANY(g0446UU_try04478)
                void_try7 = ToArray(g0441.Id()).NthPut(i,g0446UU).ToEID()
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (void_try7-void_try7) */
              if ErrorIn(void_try7) {Result = void_try7
              break
              } else {
              i = (i+1)
              }
              /* while-6 */} 
            }
            /* Let-5 */} 
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        Result = EID{g0441.Id(),0}
        }
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(Language.C_Call) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0443 *Language.Call   = Language.To_Call(self)
        /* noccur = 11 */
        if (g0443.Selector.Id() == Core.C_eval.Id()) /* If:4 */{ 
          /* Let:5 */{ 
            var g0448UU *ClaireBoolean  
            /* noccur = 1 */
            var g0448UU_try04496 EID 
            /* or:6 */{ 
              var v_or6 *ClaireBoolean  
              
              v_or6 = Equal(MakeInteger(g0443.Args.Length()).Id(),MakeInteger(1).Id())
              if (v_or6 == CTRUE) {g0448UU_try04496 = EID{CTRUE.Id(),0}
              } else /* or:7 */{ 
                var v_or6_try04508 EID 
                /* and:8 */{ 
                  var v_and8 *ClaireBoolean  
                  
                  v_and8 = Equal(MakeInteger(g0443.Args.Length()).Id(),MakeInteger(2).Id())
                  if (v_and8 == CFALSE) {v_or6_try04508 = EID{CFALSE.Id(),0}
                  } else /* arg:9 */{ 
                    var v_and8_try045110 EID 
                    v_and8_try045110 = Core.F_BELONG(val.At(1-1),g0443.Args.At(2-1))
                    /* ERROR PROTECTION INSERTED (v_and8-v_or6_try04508) */
                    if ErrorIn(v_and8_try045110) {v_or6_try04508 = v_and8_try045110
                    } else {
                    v_and8 = ToBoolean(OBJ(v_and8_try045110))
                    if (v_and8 == CFALSE) {v_or6_try04508 = EID{CFALSE.Id(),0}
                    } else /* arg:10 */{ 
                      v_or6_try04508 = EID{CTRUE.Id(),0}/* arg-10 */} 
                    /* arg-9 */} 
                  }
                  /* and-8 */} 
                /* ERROR PROTECTION INSERTED (v_or6-g0448UU_try04496) */
                if ErrorIn(v_or6_try04508) {g0448UU_try04496 = v_or6_try04508
                } else {
                v_or6 = ToBoolean(OBJ(v_or6_try04508))
                if (v_or6 == CTRUE) {g0448UU_try04496 = EID{CTRUE.Id(),0}
                } else /* or:8 */{ 
                  g0448UU_try04496 = EID{CFALSE.Id(),0}/* org-8 */} 
                /* org-7 */} 
              }
              /* or-6 */} 
            /* ERROR PROTECTION INSERTED (g0448UU-Result) */
            if ErrorIn(g0448UU_try04496) {Result = g0448UU_try04496
            } else {
            g0448UU = ToBoolean(OBJ(g0448UU_try04496))
            Result = F_Optimize_c_substitution_any(g0443.Args.At(1-1),lx,val,g0448UU)
            }
            /* Let-5 */} 
          /* If!4 */}  else if (eval_ask == CTRUE) /* If:4 */{ 
          h_index := ClEnv.Index /* Handle */
          h_base := ClEnv.Base
          /* Let:5 */{ 
            var g0452UU *ClaireList  
            /* noccur = 1 */
            var g0452UU_try04536 EID 
            /* Iteration:6 */{ 
              var v_list6 *ClaireList  
              var y *ClaireAny  
              var v_local6 *ClaireAny  
              v_list6 = g0443.Args
              g0452UU_try04536 = EID{CreateList(ToType(CEMPTY.Id()),v_list6.Length()).Id(),0}
              for CLcount := 0; CLcount < v_list6.Length(); CLcount++{ 
                y = v_list6.At(CLcount)
                var v_local6_try04548 EID 
                v_local6_try04548 = F_Optimize_c_substitution_any(y,lx,val,CTRUE)
                /* ERROR PROTECTION INSERTED (v_local6-g0452UU_try04536) */
                if ErrorIn(v_local6_try04548) {g0452UU_try04536 = v_local6_try04548
                g0452UU_try04536 = v_local6_try04548
                break
                } else {
                v_local6 = ANY(v_local6_try04548)
                ToList(OBJ(g0452UU_try04536)).PutAt(CLcount,v_local6)
                } 
              }
              /* Iteration-6 */} 
            /* ERROR PROTECTION INSERTED (g0452UU-Result) */
            if ErrorIn(g0452UU_try04536) {Result = g0452UU_try04536
            } else {
            g0452UU = ToList(OBJ(g0452UU_try04536))
            Result = Core.F_apply_property(g0443.Selector,g0452UU)
            }
            /* Let-5 */} 
          if ErrorIn(Result){ 
            /* s=EID */ClEnv.Index = h_index
            ClEnv.Base = h_base
            Core.F_tformat_string(MakeString("a strange problem happens ~A \n"),0,MakeConstantList(MakeInteger(ClEnv.Verbose).Id()))
            F_Compile_warn_void()
            Core.F_tformat_string(MakeString("failed substitution: ~S"),2,MakeConstantList(ClEnv.Exception_I.Id()))
            Result = F_Optimize_c_substitution_any(g0443.Args.Id(),lx,val,CFALSE)
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            Result = EID{g0443.Id(),0}
            }
            } 
          } else {
          Result = F_Optimize_c_substitution_any(g0443.Args.Id(),lx,val,CFALSE)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          Result = EID{g0443.Id(),0}
          }
          /* If-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(C_Instruction) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0444 *ClaireInstruction   = To_Instruction(self)
        /* noccur = 4 */
        /* For:4 */{ 
          var s *ClaireAny  
          _ = s
          Result= EID{CFALSE.Id(),0}
          var s_support *ClaireList  
          s_support = g0444.Id().Isa.Slots
          for _,s = range(s_support.ValuesO())/* loop2:5 */{ 
            var void_try6 EID 
            _ = void_try6
            /* Let:6 */{ 
              var y *ClaireAny   = Core.F_get_slot(ToSlot(s),ToObject(g0444.Id()))
              /* noccur = 1 */
              /* Let:7 */{ 
                var g0455UU *ClaireAny  
                /* noccur = 1 */
                var g0455UU_try04568 EID 
                g0455UU_try04568 = F_Optimize_c_substitution_any(y,lx,val,eval_ask)
                /* ERROR PROTECTION INSERTED (g0455UU-void_try6) */
                if ErrorIn(g0455UU_try04568) {void_try6 = g0455UU_try04568
                } else {
                g0455UU = ANY(g0455UU_try04568)
                void_try6 = Core.F_put_slot(ToSlot(s),ToObject(g0444.Id()),g0455UU).ToEID()
                }
                /* Let-7 */} 
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
        Result = EID{g0444.Id(),0}
        }
        /* Let-3 */} 
      } else {
      Result = self.ToEID()
      /* If-2 */} 
    return Result} 
  
// The EID go function for: c_substitution @ any (throw: true) 
func E_Optimize_c_substitution_any (self EID,lx EID,val EID,eval_ask EID) EID { 
    return /*(sm for c_substitution @ any= EID)*/ F_Optimize_c_substitution_any(ANY(self),
      ToList(OBJ(lx)),
      ToList(OBJ(val)),
      ToBoolean(OBJ(eval_ask)) )} 
  
// needed
/* {1} OPT.The go function for: eval(x:any,y:class) [] */
func F_eval_any2 (x *ClaireAny ,y *ClaireClass ) EID { 
    var Result EID 
    Result = EVAL(x)
    return Result} 
  
// The EID go function for: eval @ list<type_expression>(any, class) (throw: true) 
func E_eval_any2 (x EID,y EID) EID { 
    return /*(sm for eval @ list<type_expression>(any, class)= EID)*/ F_eval_any2(ANY(x),ToClass(OBJ(y)) )} 
  
// returns the list of bound variables in a piece of code
/* {1} OPT.The go function for: bound_variables(self:any) [] */
func F_Optimize_bound_variables_any (self *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var l *ClaireList   = ToType(C_any.Id()).EmptyList()
      /* noccur = 6 */
      if (self.Isa.IsIn(Language.C_Instruction_with_var) == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var g0457 *Language.InstructionWithVar   = Language.To_InstructionWithVar(self)
          /* noccur = 1 */
          l = MakeList(ToType(C_any.Id()),g0457.ClaireVar.Id())
          /* Let-4 */} 
        /* If-3 */} 
      if (self.Isa.IsIn(C_Variable) == CTRUE) /* If:3 */{ 
        Result = EID{CNIL.Id(),0}
        /* If!3 */}  else if (self.Isa.IsIn(C_Instruction) == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var g0459 *ClaireInstruction   = To_Instruction(self)
          /* noccur = 2 */
          /* For:5 */{ 
            var s *ClaireAny  
            _ = s
            Result= EID{CFALSE.Id(),0}
            for _,s = range(g0459.Isa.Slots.ValuesO())/* loop:6 */{ 
              var void_try7 EID 
              _ = void_try7
              var l_try04617 EID 
              /* Let:7 */{ 
                var g0462UU *ClaireList  
                /* noccur = 1 */
                var g0462UU_try04638 EID 
                g0462UU_try04638 = F_Optimize_bound_variables_any(Core.F_get_slot(ToSlot(s),ToObject(g0459.Id())))
                /* ERROR PROTECTION INSERTED (g0462UU-l_try04617) */
                if ErrorIn(g0462UU_try04638) {l_try04617 = g0462UU_try04638
                } else {
                g0462UU = ToList(OBJ(g0462UU_try04638))
                l_try04617 = l.Add_star(g0462UU)
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (l-Result) */
              if ErrorIn(l_try04617) {Result = l_try04617
              Result = l_try04617
              break
              } else {
              l = ToList(OBJ(l_try04617))
              void_try7 = EID{l.Id(),0}
              }
              /* loop-6 */} 
            /* For-5 */} 
          /* Let-4 */} 
        /* If!3 */}  else if (self.Isa.IsIn(C_list) == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var g0460 *ClaireList   = ToList(self)
          /* noccur = 1 */
          /* For:5 */{ 
            var x *ClaireAny  
            _ = x
            Result= EID{CFALSE.Id(),0}
            var x_support *ClaireList  
            x_support = g0460
            x_len := x_support.Length()
            for i_it := 0; i_it < x_len; i_it++ { 
              x = x_support.At(i_it)
              var void_try7 EID 
              _ = void_try7
              var l_try04647 EID 
              /* Let:7 */{ 
                var g0465UU *ClaireList  
                /* noccur = 1 */
                var g0465UU_try04668 EID 
                g0465UU_try04668 = F_Optimize_bound_variables_any(x)
                /* ERROR PROTECTION INSERTED (g0465UU-l_try04647) */
                if ErrorIn(g0465UU_try04668) {l_try04647 = g0465UU_try04668
                } else {
                g0465UU = ToList(OBJ(g0465UU_try04668))
                l_try04647 = l.Add_star(g0465UU)
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (l-Result) */
              if ErrorIn(l_try04647) {Result = l_try04647
              Result = l_try04647
              break
              } else {
              l = ToList(OBJ(l_try04647))
              void_try7 = EID{l.Id(),0}
              }
              /* loop-6 */} 
            /* For-5 */} 
          /* Let-4 */} 
        } else {
        Result = EID{CFALSE.Id(),0}
        /* If-3 */} 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = EID{l.Id(),0}
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: bound_variables @ any (throw: true) 
func E_Optimize_bound_variables_any (self EID) EID { 
    return /*(sm for bound_variables @ any= EID)*/ F_Optimize_bound_variables_any(ANY(self) )} 
  
// we must recognize true boolean ! coercion
/* {1} OPT.The go function for: c_boolean(x:any) [] */
func F_Optimize_c_boolean_any (x *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var tx *ClaireType  
      /* noccur = 3 */
      var tx_try04683 EID 
      tx_try04683 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
      /* ERROR PROTECTION INSERTED (tx-Result) */
      if ErrorIn(tx_try04683) {Result = tx_try04683
      } else {
      tx = ToType(OBJ(tx_try04683))
      /* Let:3 */{ 
        var ptx *ClaireType   = F_Optimize_ptype_type(tx)
        /* noccur = 1 */
        if (ptx.Included(ToType(C_boolean.Id())) == CTRUE) /* If:4 */{ 
          if (x.Isa.IsIn(Language.C_Call) == CTRUE) /* If:5 */{ 
            /* Let:6 */{ 
              var g0467 *Language.Call   = Language.To_Call(x)
              /* noccur = 5 */
              var g0469I *ClaireBoolean  
              var g0469I_try04707 EID 
              /* and:7 */{ 
                var v_and7 *ClaireBoolean  
                
                v_and7 = Equal(g0467.Selector.Id(),Core.C_not.Id())
                if (v_and7 == CFALSE) {g0469I_try04707 = EID{CFALSE.Id(),0}
                } else /* arg:8 */{ 
                  var v_and7_try04719 EID 
                  /* Let:9 */{ 
                    var g0472UU *ClaireType  
                    /* noccur = 1 */
                    var g0472UU_try047310 EID 
                    /* Let:10 */{ 
                      var g0474UU *ClaireType  
                      /* noccur = 1 */
                      var g0474UU_try047511 EID 
                      g0474UU_try047511 = Core.F_CALL(C_c_type,ARGS(g0467.Args.At(1-1).ToEID()))
                      /* ERROR PROTECTION INSERTED (g0474UU-g0472UU_try047310) */
                      if ErrorIn(g0474UU_try047511) {g0472UU_try047310 = g0474UU_try047511
                      } else {
                      g0474UU = ToType(OBJ(g0474UU_try047511))
                      g0472UU_try047310 = EID{F_Optimize_ptype_type(g0474UU).Id(),0}
                      }
                      /* Let-10 */} 
                    /* ERROR PROTECTION INSERTED (g0472UU-v_and7_try04719) */
                    if ErrorIn(g0472UU_try047310) {v_and7_try04719 = g0472UU_try047310
                    } else {
                    g0472UU = ToType(OBJ(g0472UU_try047310))
                    v_and7_try04719 = EID{Core.F__I_equal_any(g0472UU.Id(),C_boolean.Id()).Id(),0}
                    }
                    /* Let-9 */} 
                  /* ERROR PROTECTION INSERTED (v_and7-g0469I_try04707) */
                  if ErrorIn(v_and7_try04719) {g0469I_try04707 = v_and7_try04719
                  } else {
                  v_and7 = ToBoolean(OBJ(v_and7_try04719))
                  if (v_and7 == CFALSE) {g0469I_try04707 = EID{CFALSE.Id(),0}
                  } else /* arg:9 */{ 
                    g0469I_try04707 = EID{CTRUE.Id(),0}/* arg-9 */} 
                  /* arg-8 */} 
                }
                /* and-7 */} 
              /* ERROR PROTECTION INSERTED (g0469I-Result) */
              if ErrorIn(g0469I_try04707) {Result = g0469I_try04707
              } else {
              g0469I = ToBoolean(OBJ(g0469I_try04707))
              if (g0469I == CTRUE) /* If:7 */{ 
                /* Let:8 */{ 
                  var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                  /* noccur = 11 */
                  _CL_obj.Selector = ToProperty(Core.C__I_equal.Id())
                  /* update:9 */{ 
                    var va_arg1 *Language.Call  
                    var va_arg2 *ClaireList  
                    va_arg1 = _CL_obj
                    /* Construct:10 */{ 
                      var v_bag_arg *ClaireAny  
                      va_arg2= ToType(CEMPTY.Id()).EmptyList()
                      /* Let:11 */{ 
                        var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                        /* noccur = 5 */
                        _CL_obj.Selector = C_boolean_I
                        _CL_obj.Args = MakeConstantList(g0467.Args.At(1-1))
                        v_bag_arg = _CL_obj.Id()
                        /* Let-11 */} 
                      va_arg2.AddFast(v_bag_arg)
                      va_arg2.AddFast(CTRUE.Id())/* Construct-10 */} 
                    /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                    va_arg1.Args = va_arg2
                    /* update-9 */} 
                  g0467 = _CL_obj
                  /* Let-8 */} 
                Result = EID{g0467.Id(),0}
                } else {
                Result = EID{CFALSE.Id(),0}
                /* If-7 */} 
              }
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              x = g0467.Id()
              Result = x.ToEID()
              }
              /* Let-6 */} 
            } else {
            Result = EID{CFALSE.Id(),0}
            /* If-5 */} 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          if (tx.Included(ToType(C_boolean.Id())) == CTRUE) /* If:5 */{ 
            Result = F_Compile_c_strict_code_any(x,C_boolean)
            } else {
            Result = Core.F_CALL(C_c_code,ARGS(x.ToEID(),EID{C_boolean.Id(),0}))
            /* If-5 */} 
          }
          /* If!4 */}  else if (tx.Included(ToType(C_list.Id())) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var g0476UU *Language.Call  
            /* noccur = 1 */
            /* Let:6 */{ 
              var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
              /* noccur = 11 */
              _CL_obj.Selector = ToProperty(Core.C__I_equal.Id())
              /* update:7 */{ 
                var va_arg1 *Language.Call  
                var va_arg2 *ClaireList  
                va_arg1 = _CL_obj
                /* Construct:8 */{ 
                  var v_bag_arg *ClaireAny  
                  va_arg2= ToType(CEMPTY.Id()).EmptyList()
                  /* Let:9 */{ 
                    var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                    /* noccur = 5 */
                    _CL_obj.Selector = C_length
                    _CL_obj.Args = MakeConstantList(x)
                    v_bag_arg = _CL_obj.Id()
                    /* Let-9 */} 
                  va_arg2.AddFast(v_bag_arg)
                  va_arg2.AddFast(MakeInteger(0).Id())/* Construct-8 */} 
                /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                va_arg1.Args = va_arg2
                /* update-7 */} 
              g0476UU = _CL_obj
              /* Let-6 */} 
            Result = Core.F_CALL(C_c_code,ARGS(EID{g0476UU.Id(),0}))
            /* Let-5 */} 
          } else {
          /* Let:5 */{ 
            var g0477UU *Language.Call  
            /* noccur = 1 */
            /* Let:6 */{ 
              var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
              /* noccur = 5 */
              _CL_obj.Selector = C_boolean_I
              _CL_obj.Args = MakeConstantList(x)
              g0477UU = _CL_obj
              /* Let-6 */} 
            Result = Core.F_CALL(C_c_code,ARGS(EID{g0477UU.Id(),0}))
            /* Let-5 */} 
          /* If-4 */} 
        /* Let-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_boolean @ any (throw: true) 
func E_Optimize_c_boolean_any (x EID) EID { 
    return /*(sm for c_boolean @ any= EID)*/ F_Optimize_c_boolean_any(ANY(x) )} 
  