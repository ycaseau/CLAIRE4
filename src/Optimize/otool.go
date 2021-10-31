/***** CLAIRE Compilation of file /Users/ycaseau/Dropbox/src/clairev4.0/src/compile/otool.cl 
         [version 4.0.02 / safety 5] Saturday 10-30-2021 *****/

package Optimize
import (_ "fmt"
	. "Kernel"
	"Core"
	"Language"
	"Reader"
)

//-------- dumb function to prevent import errors --------
func import_g0172() { 
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
          var va_arg2_try01735 EID 
          va_arg2_try01735 = Core.F_CALL(C_c_code,ARGS(self.Arg.ToEID(),EID{s.Id(),0}))
          /* ERROR PROTECTION INSERTED (va_arg2-Result) */
          if ErrorIn(va_arg2_try01735) {Result = va_arg2_try01735
          } else {
          va_arg2 = ANY(va_arg2_try01735)
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
        var g0174 *Language.Call   = Language.To_Call(x)
        /* noccur = 2 */
        /* and:4 */{ 
          var v_and4 *ClaireBoolean  
          
          v_and4 = Equal(g0174.Selector.Id(),y.Selector.Id())
          if (v_and4 == CFALSE) {Result = EID{CFALSE.Id(),0}
          } else /* arg:5 */{ 
            var v_and4_try01766 EID 
            /* Let:6 */{ 
              var g0177UU *ClaireList  
              /* noccur = 1 */
              var g0177UU_try01787 EID 
              /* Iteration:7 */{ 
                var v_list7 *ClaireList  
                var z *ClaireAny  
                var v_local7 *ClaireAny  
                v_list7 = g0174.Args
                g0177UU_try01787 = EID{CreateList(ToType(CEMPTY.Id()),v_list7.Length()).Id(),0}
                for CLcount := 0; CLcount < v_list7.Length(); CLcount++{ 
                  z = v_list7.At(CLcount)
                  var v_local7_try01799 EID 
                  v_local7_try01799 = Core.F_CALL(C_c_type,ARGS(z.ToEID()))
                  /* ERROR PROTECTION INSERTED (v_local7-g0177UU_try01787) */
                  if ErrorIn(v_local7_try01799) {g0177UU_try01787 = v_local7_try01799
                  g0177UU_try01787 = v_local7_try01799
                  break
                  } else {
                  v_local7 = ANY(v_local7_try01799)
                  ToList(OBJ(g0177UU_try01787)).PutAt(CLcount,v_local7)
                  } 
                }
                /* Iteration-7 */} 
              /* ERROR PROTECTION INSERTED (g0177UU-v_and4_try01766) */
              if ErrorIn(g0177UU_try01787) {v_and4_try01766 = g0177UU_try01787
              } else {
              g0177UU = ToList(OBJ(g0177UU_try01787))
              v_and4_try01766 = EID{Core.F_tmatch_ask_list(g0177UU,y.Arg).Id(),0}
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (v_and4-Result) */
            if ErrorIn(v_and4_try01766) {Result = v_and4_try01766
            } else {
            v_and4 = ToBoolean(OBJ(v_and4_try01766))
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
        var g0180 *ClairePattern   = To_ClairePattern(y.Id())
        /* noccur = 2 */
        if (x.Selector.Id() == g0180.Selector.Id()) /* If:4 */{ 
          /* Let:5 */{ 
            var _CL_obj *ClairePattern   = To_ClairePattern(new(ClairePattern).Is(C_Optimize_Pattern))
            /* noccur = 4 */
            _CL_obj.Selector = x.Selector
            Result = Core.F_write_property(C_args,ToObject(_CL_obj.Id()),ANY(Core.F_CALL(ToProperty(Core.C_glb.Id()),ARGS(Core.F_CALL(C_args,ARGS(EID{x.Id(),0})),Core.F_CALL(C_args,ARGS(EID{g0180.Id(),0}))))))
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
        var g0182 *ClairePattern   = To_ClairePattern(y.Id())
        /* noccur = 3 */
        /* and:4 */{ 
          var v_and4 *ClaireBoolean  
          
          v_and4 = Equal(x.Selector.Id(),g0182.Selector.Id())
          if (v_and4 == CFALSE) {Result = CFALSE
          } else /* arg:5 */{ 
            v_and4 = Equal(MakeInteger(x.Arg.Length()).Id(),MakeInteger(g0182.Arg.Length()).Id())
            if (v_and4 == CFALSE) {Result = CFALSE
            } else /* arg:6 */{ 
              /* Let:7 */{ 
                var g0185UU *ClaireAny  
                /* noccur = 1 */
                /* Let:8 */{ 
                  var i int  = 1
                  /* noccur = 5 */
                  /* Let:9 */{ 
                    var g0183 int  = x.Arg.Length()
                    /* noccur = 1 */
                    g0185UU= CFALSE.Id()
                    for (i <= g0183) /* while:10 */{ 
                      if (Core.F__equaltype_ask_any(ToType(x.Arg.At(i-1)),ToType(g0182.Arg.At(i-1))) != CTRUE) /* If:11 */{ 
                         /*v = g0185UU, s =any*/
g0185UU = CTRUE.Id()
                        break
                        /* If-11 */} 
                      i = (i+1)
                      /* while-10 */} 
                    /* Let-9 */} 
                  /* Let-8 */} 
                v_and4 = Core.F_not_any(g0185UU)
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
        var g0186 *ClaireSet   = ToSet(x.Id())
        /* noccur = 1 */
        /* Let:4 */{ 
          var g0190UU *ClaireAny  
          /* noccur = 1 */
          var g0190UU_try01915 EID 
          /* For:5 */{ 
            var z *ClaireAny  
            _ = z
            g0190UU_try01915= EID{CFALSE.Id(),0}
            for _,z = range(g0186.Values)/* loop:6 */{ 
              var void_try7 EID 
              _ = void_try7
              var g0192I *ClaireBoolean  
              var g0192I_try01937 EID 
              /* Let:7 */{ 
                var g0194UU *ClaireBoolean  
                /* noccur = 1 */
                var g0194UU_try01958 EID 
                g0194UU_try01958 = F__Z_any3(z,y)
                /* ERROR PROTECTION INSERTED (g0194UU-g0192I_try01937) */
                if ErrorIn(g0194UU_try01958) {g0192I_try01937 = g0194UU_try01958
                } else {
                g0194UU = ToBoolean(OBJ(g0194UU_try01958))
                g0192I_try01937 = EID{g0194UU.Not.Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (g0192I-void_try7) */
              if ErrorIn(g0192I_try01937) {void_try7 = g0192I_try01937
              } else {
              g0192I = ToBoolean(OBJ(g0192I_try01937))
              if (g0192I == CTRUE) /* If:7 */{ 
                 /*v = g0190UU_try01915, s =EID*/
g0190UU_try01915 = EID{CTRUE.Id(),0}
                break
                } else {
                void_try7 = EID{CFALSE.Id(),0}
                /* If-7 */} 
              }
              /* ERROR PROTECTION INSERTED (void_try7-g0190UU_try01915) */
              if ErrorIn(void_try7) {g0190UU_try01915 = void_try7
              g0190UU_try01915 = void_try7
              break
              } else {
              }
              /* loop-6 */} 
            /* For-5 */} 
          /* ERROR PROTECTION INSERTED (g0190UU-Result) */
          if ErrorIn(g0190UU_try01915) {Result = g0190UU_try01915
          } else {
          g0190UU = ANY(g0190UU_try01915)
          Result = EID{Core.F_not_any(g0190UU).Id(),0}
          }
          /* Let-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (x.Isa.IsIn(C_Optimize_Pattern) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0187 *ClairePattern   = To_ClairePattern(x.Id())
        /* noccur = 4 */
        /* and:4 */{ 
          var v_and4 *ClaireBoolean  
          
          v_and4 = Equal(g0187.Selector.Id(),y.Selector.Id())
          if (v_and4 == CFALSE) {Result = EID{CFALSE.Id(),0}
          } else /* arg:5 */{ 
            v_and4 = Equal(MakeInteger(g0187.Arg.Length()).Id(),MakeInteger(y.Arg.Length()).Id())
            if (v_and4 == CFALSE) {Result = EID{CFALSE.Id(),0}
            } else /* arg:6 */{ 
              /* Let:7 */{ 
                var g0196UU *ClaireAny  
                /* noccur = 1 */
                /* Let:8 */{ 
                  var i int  = 1
                  /* noccur = 5 */
                  /* Let:9 */{ 
                    var g0188 int  = g0187.Arg.Length()
                    /* noccur = 1 */
                    g0196UU= CFALSE.Id()
                    for (i <= g0188) /* while:10 */{ 
                      if (Core.F__equaltype_ask_any(ToType(g0187.Arg.At(i-1)),ToType(y.Arg.At(i-1))) != CTRUE) /* If:11 */{ 
                         /*v = g0196UU, s =any*/
g0196UU = CTRUE.Id()
                        break
                        /* If-11 */} 
                      i = (i+1)
                      /* while-10 */} 
                    /* Let-9 */} 
                  /* Let-8 */} 
                v_and4 = Core.F_not_any(g0196UU)
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
        var g0198I *ClaireBoolean  
        /* and:4 */{ 
          var v_and4 *ClaireBoolean  
          
          v_and4 = C_OPT.Ignore.Contain_ask(s.Id()).Not
          if (v_and4 == CFALSE) {g0198I = CFALSE
          } else /* arg:5 */{ 
            v_and4 = MakeBoolean((s.Open <= 1) || (s.Open == 4))
            if (v_and4 == CFALSE) {g0198I = CFALSE
            } else /* arg:6 */{ 
              if (_Ztype.Isa.IsIn(C_list) == CTRUE) /* If:7 */{ 
                /* Let:8 */{ 
                  var g0197 *ClaireList   = ToList(_Ztype)
                  /* noccur = 1 */
                  v_and4 = Core.F__I_equal_any(MakeInteger(ToTypeExpression(g0197.At(1-1)).Class_I().Open).Id(),MakeInteger(3).Id())
                  /* Let-8 */} 
                } else {
                v_and4 = CFALSE
                /* If-7 */} 
              if (v_and4 == CFALSE) {g0198I = CFALSE
              } else /* arg:7 */{ 
                g0198I = CTRUE/* arg-7 */} 
              /* arg-6 */} 
            /* arg-5 */} 
          /* and-4 */} 
        if (g0198I == CTRUE) /* If:4 */{ 
          F_Compile_warn_void()
          Result = Core.F_tformat_string(MakeString("wrongly typed message ~S (~S) [256]\n"),2,MakeConstantList(self.Id(),_Ztype))
          /* If!4 */}  else if (C_compiler.Optimize_ask == CTRUE) /* If:4 */{ 
          F_Compile_notice_void()
          Result = EID{CFALSE.Id(),0}
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
        Result = EID{CFALSE.Id(),0}
        } else {
        Result = EID{CFALSE.Id(),0}
        /* If-3 */} 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      /* Let:3 */{ 
        var m *Language.Call  
        /* noccur = 2 */
        var m_try01994 EID 
        m_try01994 = F_Optimize_open_message_property(self.Selector,self.Args)
        /* ERROR PROTECTION INSERTED (m-Result) */
        if ErrorIn(m_try01994) {Result = m_try01994
        } else {
        m = Language.To_Call(OBJ(m_try01994))
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
    var g0200I *ClaireBoolean  
    var g0200I_try02012 EID 
    /* Let:2 */{ 
      var g0202UU *ClaireBoolean  
      /* noccur = 1 */
      var g0202UU_try02033 EID 
      /* Let:3 */{ 
        var g0204UU *ClaireAny  
        /* noccur = 1 */
        var g0204UU_try02054 EID 
        g0204UU_try02054 = Core.F_CALL(ToProperty(C__exp.Id()),ARGS(EID{y.Id(),0},EID{self.Range.Id(),0}))
        /* ERROR PROTECTION INSERTED (g0204UU-g0202UU_try02033) */
        if ErrorIn(g0204UU_try02054) {g0202UU_try02033 = g0204UU_try02054
        } else {
        g0204UU = ANY(g0204UU_try02054)
        g0202UU_try02033 = EID{F_boolean_I_any(g0204UU).Id(),0}
        }
        /* Let-3 */} 
      /* ERROR PROTECTION INSERTED (g0202UU-g0200I_try02012) */
      if ErrorIn(g0202UU_try02033) {g0200I_try02012 = g0202UU_try02033
      } else {
      g0202UU = ToBoolean(OBJ(g0202UU_try02033))
      g0200I_try02012 = EID{Core.F__I_equal_any(g0202UU.Id(),CTRUE.Id()).Id(),0}
      }
      /* Let-2 */} 
    /* ERROR PROTECTION INSERTED (g0200I-Result) */
    if ErrorIn(g0200I_try02012) {Result = g0200I_try02012
    } else {
    g0200I = ToBoolean(OBJ(g0200I_try02012))
    if (g0200I == CTRUE) /* If:2 */{ 
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
      var g0207I *ClaireBoolean  
      if (r.Isa.IsIn(C_Union) == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var g0206 *ClaireUnion   = To_Union(r.Id())
          /* noccur = 1 */
          g0207I = Equal(g0206.T1.Id(),CEMPTY.Id())
          /* Let-4 */} 
        } else {
        g0207I = CFALSE
        /* If-3 */} 
      if (g0207I == CTRUE) /* If:3 */{ 
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
      var c1_try02083 EID 
      c1_try02083 = F_Compile_stupid_t_any1(self)
      /* ERROR PROTECTION INSERTED (c1-Result) */
      if ErrorIn(c1_try02083) {Result = c1_try02083
      } else {
      c1 = ToClass(OBJ(c1_try02083))
      /* Let:3 */{ 
        var c2 *ClaireClass  
        /* noccur = 1 */
        var c2_try02094 EID 
        c2_try02094 = F_Compile_stupid_t_any1(x)
        /* ERROR PROTECTION INSERTED (c2-Result) */
        if ErrorIn(c2_try02094) {Result = c2_try02094
        } else {
        c2 = ToClass(OBJ(c2_try02094))
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
        var g0210 *ClaireUnion   = To_Union(self.Id())
        /* noccur = 3 */
        /* and:4 */{ 
          var v_and4 *ClaireBoolean  
          
          v_and4 = Equal(C_set.Id(),g0210.T2.Isa.Id())
          if (v_and4 == CFALSE) {Result = EID{CFALSE.Id(),0}
          } else /* arg:5 */{ 
            var v_and4_try02126 EID 
            /* Let:6 */{ 
              var g0213UU *ClaireAny  
              /* noccur = 1 */
              var g0213UU_try02147 EID 
              g0213UU_try02147 = Core.F_CALL(C_size,ARGS(EID{g0210.T2.Id(),0}))
              /* ERROR PROTECTION INSERTED (g0213UU-v_and4_try02126) */
              if ErrorIn(g0213UU_try02147) {v_and4_try02126 = g0213UU_try02147
              } else {
              g0213UU = ANY(g0213UU_try02147)
              v_and4_try02126 = EID{Equal(g0213UU,MakeInteger(1).Id()).Id(),0}
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (v_and4-Result) */
            if ErrorIn(v_and4_try02126) {Result = v_and4_try02126
            } else {
            v_and4 = ToBoolean(OBJ(v_and4_try02126))
            if (v_and4 == CFALSE) {Result = EID{CFALSE.Id(),0}
            } else /* arg:6 */{ 
              var v_and4_try02157 EID 
              /* Let:7 */{ 
                var g0216UU *ClaireAny  
                /* noccur = 1 */
                var g0216UU_try02178 EID 
                g0216UU_try02178 = Core.F_the_type(g0210.T2)
                /* ERROR PROTECTION INSERTED (g0216UU-v_and4_try02157) */
                if ErrorIn(g0216UU_try02178) {v_and4_try02157 = g0216UU_try02178
                } else {
                g0216UU = ANY(g0216UU_try02178)
                v_and4_try02157 = EID{Equal(g0216UU,CNULL).Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (v_and4-Result) */
              if ErrorIn(v_and4_try02157) {Result = v_and4_try02157
              } else {
              v_and4 = ToBoolean(OBJ(v_and4_try02157))
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
        var g0218 *ClaireUnion   = To_Union(x.Id())
        /* noccur = 1 */
        Result = Equal(g0218.T1.Id(),C_any.Id())
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
        var g0220 *ClaireUnion   = To_Union(x.Id())
        /* noccur = 3 */
        if (g0220.T1.Id() == C_any.Id()) /* If:4 */{ 
          Result = g0220.T2
          } else {
          Result = ToType(g0220.Id())
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
    var g0222I *ClaireBoolean  
    var g0222I_try02232 EID 
    /* or:2 */{ 
      var v_or2 *ClaireBoolean  
      
      v_or2 = MakeBoolean((self.Range.Id() == CNULL))
      if (v_or2 == CTRUE) {g0222I_try02232 = EID{CTRUE.Id(),0}
      } else /* or:3 */{ 
        var v_or2_try02244 EID 
        v_or2_try02244 = F_Optimize_extended_ask_type(self.Range)
        /* ERROR PROTECTION INSERTED (v_or2-g0222I_try02232) */
        if ErrorIn(v_or2_try02244) {g0222I_try02232 = v_or2_try02244
        } else {
        v_or2 = ToBoolean(OBJ(v_or2_try02244))
        if (v_or2 == CTRUE) {g0222I_try02232 = EID{CTRUE.Id(),0}
        } else /* or:4 */{ 
          g0222I_try02232 = EID{CFALSE.Id(),0}/* org-4 */} 
        /* org-3 */} 
      }
      /* or-2 */} 
    /* ERROR PROTECTION INSERTED (g0222I-Result) */
    if ErrorIn(g0222I_try02232) {Result = g0222I_try02232
    } else {
    g0222I = ToBoolean(OBJ(g0222I_try02232))
    if (g0222I == CTRUE) /* If:2 */{ 
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
        var g0225 *ClaireVariable   = To_Variable(self)
        /* noccur = 6 */
        if (F_boolean_I_any(F_Compile_sort_equal_class(F_Compile_osort_any(g0225.Range.Id()),F_Compile_osort_any(y.Id()))) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var c1 *ClaireClass   = F_Compile_psort_any(g0225.Range.Class_I().Id())
            /* noccur = 2 */
            if (c1.Id() != F_Compile_psort_any(y.Class_I().Id()).Id()) /* If:6 */{ 
              /* update:7 */{ 
                var va_arg1 *ClaireVariable  
                var va_arg2 *ClaireType  
                va_arg1 = g0225
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
              g0225.Range = y
              /* If-6 */} 
            /* Let-5 */} 
          /* If!4 */}  else if (F_Compile_osort_any(g0225.Range.Id()).Id() == C_any.Id()) /* If:4 */{ 
          g0225.Range = F_Optimize_sort_abstract_I_type(y)
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
        var g0226 *ClaireVariable   = To_Variable(self)
        /* noccur = 6 */
        if (F_boolean_I_any(F_Compile_sort_equal_class(F_Compile_osort_any(g0226.Range.Id()),F_Compile_osort_any(y.Id()))) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var c1 *ClaireClass   = F_Compile_psort_any(g0226.Range.Class_I().Id())
            /* noccur = 2 */
            if (c1.Id() != F_Compile_psort_any(y.Class_I().Id()).Id()) /* If:6 */{ 
              /* update:7 */{ 
                var va_arg1 *ClaireVariable  
                var va_arg2 *ClaireType  
                va_arg1 = g0226
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
              g0226.Range = y
              /* If-6 */} 
            /* Let-5 */} 
          /* If!4 */}  else if (F_Compile_osort_any(g0226.Range.Id()).Id() == C_any.Id()) /* If:4 */{ 
          g0226.Range = F_Optimize_sort_abstract_I_type(y)
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
        var g0227 *Language.Let   = Language.To_Let(self)
        /* noccur = 1 */
        Result = F_Compile_return_type_any(g0227.Arg)
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(Language.C_Do) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0228 *Language.Do   = Language.To_Do(self)
        /* noccur = 1 */
        /* Let:4 */{ 
          var x *ClaireSet   = CEMPTY
          /* noccur = 3 */
          /* For:5 */{ 
            var y *ClaireAny  
            _ = y
            Result= EID{CFALSE.Id(),0}
            var y_support *ClaireList  
            y_support = g0228.Args
            y_len := y_support.Length()
            for i_it := 0; i_it < y_len; i_it++ { 
              y = y_support.At(i_it)
              var void_try7 EID 
              _ = void_try7
              var x_try02347 EID 
              /* Let:7 */{ 
                var g0235UU *ClaireType  
                /* noccur = 1 */
                var g0235UU_try02368 EID 
                g0235UU_try02368 = F_Compile_return_type_any(y)
                /* ERROR PROTECTION INSERTED (g0235UU-x_try02347) */
                if ErrorIn(g0235UU_try02368) {x_try02347 = g0235UU_try02368
                } else {
                g0235UU = ToType(OBJ(g0235UU_try02368))
                x_try02347 = Core.F_CALL(ToProperty(C__exp.Id()),ARGS(EID{x.Id(),0},EID{g0235UU.Id(),0}))
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (x-Result) */
              if ErrorIn(x_try02347) {Result = x_try02347
              Result = x_try02347
              break
              } else {
              x = ToSet(OBJ(x_try02347))
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
        var g0229 *Language.If   = Language.To_If(self)
        /* noccur = 2 */
        /* Let:4 */{ 
          var g0237UU *ClaireType  
          /* noccur = 1 */
          var g0237UU_try02395 EID 
          g0237UU_try02395 = F_Compile_return_type_any(g0229.Arg)
          /* ERROR PROTECTION INSERTED (g0237UU-Result) */
          if ErrorIn(g0237UU_try02395) {Result = g0237UU_try02395
          } else {
          g0237UU = ToType(OBJ(g0237UU_try02395))
          /* Let:5 */{ 
            var g0238UU *ClaireType  
            /* noccur = 1 */
            var g0238UU_try02406 EID 
            g0238UU_try02406 = F_Compile_return_type_any(g0229.Other)
            /* ERROR PROTECTION INSERTED (g0238UU-Result) */
            if ErrorIn(g0238UU_try02406) {Result = g0238UU_try02406
            } else {
            g0238UU = ToType(OBJ(g0238UU_try02406))
            Result = Core.F_CALL(ToProperty(C__exp.Id()),ARGS(EID{g0237UU.Id(),0},EID{g0238UU.Id(),0}))
            }
            /* Let-5 */} 
          }
          /* Let-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(Language.C_Return) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0230 *Language.Return   = Language.To_Return(self)
        /* noccur = 1 */
        Result = Core.F_CALL(C_c_type,ARGS(g0230.Arg.ToEID()))
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(Language.C_Case) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0231 *Language.Case   = Language.To_Case(self)
        /* noccur = 1 */
        /* Let:4 */{ 
          var x *ClaireSet   = CEMPTY
          /* noccur = 3 */
          /* For:5 */{ 
            var y *ClaireAny  
            _ = y
            Result= EID{CFALSE.Id(),0}
            var y_support *ClaireList  
            y_support = g0231.Args
            y_len := y_support.Length()
            for i_it := 0; i_it < y_len; i_it++ { 
              y = y_support.At(i_it)
              var void_try7 EID 
              _ = void_try7
              var x_try02417 EID 
              /* Let:7 */{ 
                var g0242UU *ClaireType  
                /* noccur = 1 */
                var g0242UU_try02438 EID 
                g0242UU_try02438 = F_Compile_return_type_any(y)
                /* ERROR PROTECTION INSERTED (g0242UU-x_try02417) */
                if ErrorIn(g0242UU_try02438) {x_try02417 = g0242UU_try02438
                } else {
                g0242UU = ToType(OBJ(g0242UU_try02438))
                x_try02417 = Core.F_CALL(ToProperty(C__exp.Id()),ARGS(EID{x.Id(),0},EID{g0242UU.Id(),0}))
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (x-Result) */
              if ErrorIn(x_try02417) {Result = x_try02417
              Result = x_try02417
              break
              } else {
              x = ToSet(OBJ(x_try02417))
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
        var g0232 *Language.ClaireHandle   = Language.To_ClaireHandle(self)
        /* noccur = 1 */
        Result = F_Compile_return_type_any(g0232.Arg)
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
      var g0244UU *ClaireAny  
      /* noccur = 1 */
      var g0244UU_try02453 EID 
      g0244UU_try02453 = Core.F_CALL(C_Compile_self_code,ARGS(EID{self.Id(),0}))
      /* ERROR PROTECTION INSERTED (g0244UU-Result) */
      if ErrorIn(g0244UU_try02453) {Result = g0244UU_try02453
      } else {
      g0244UU = ANY(g0244UU_try02453)
      Result = Core.F_CALL(C_c_code,ARGS(g0244UU.ToEID(),EID{s.Id(),0}))
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
        var va_arg2_try02464 EID 
        /* Construct:4 */{ 
          var v_bag_arg *ClaireAny  
          va_arg2_try02464= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
          ToList(OBJ(va_arg2_try02464)).AddFast(self.Arg.Id())
          var v_bag_arg_try02475 EID 
          v_bag_arg_try02475 = Core.F_CALL(C_c_code,ARGS(EID{self.T1.Id(),0},EID{C_type.Id(),0}))
          /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try02464) */
          if ErrorIn(v_bag_arg_try02475) {va_arg2_try02464 = v_bag_arg_try02475
          } else {
          v_bag_arg = ANY(v_bag_arg_try02475)
          ToList(OBJ(va_arg2_try02464)).AddFast(v_bag_arg)}
          /* Construct-4 */} 
        /* ERROR PROTECTION INSERTED (va_arg2-Result) */
        if ErrorIn(va_arg2_try02464) {Result = va_arg2_try02464
        } else {
        va_arg2 = ToList(OBJ(va_arg2_try02464))
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
          var va_arg2_try02485 EID 
          /* Construct:5 */{ 
            var v_bag_arg *ClaireAny  
            va_arg2_try02485= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
            ToList(OBJ(va_arg2_try02485)).AddFast(self.Arg.Id())
            var v_bag_arg_try02496 EID 
            /* Let:6 */{ 
              var g0250UU *ClaireAny  
              /* noccur = 1 */
              var g0250UU_try02517 EID 
              g0250UU_try02517 = Core.F_the_type(ToType(self.Args.At(1-1)))
              /* ERROR PROTECTION INSERTED (g0250UU-v_bag_arg_try02496) */
              if ErrorIn(g0250UU_try02517) {v_bag_arg_try02496 = g0250UU_try02517
              } else {
              g0250UU = ANY(g0250UU_try02517)
              v_bag_arg_try02496 = Core.F_CALL(C_c_code,ARGS(g0250UU.ToEID(),EID{C_type.Id(),0}))
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try02485) */
            if ErrorIn(v_bag_arg_try02496) {va_arg2_try02485 = v_bag_arg_try02496
            } else {
            v_bag_arg = ANY(v_bag_arg_try02496)
            ToList(OBJ(va_arg2_try02485)).AddFast(v_bag_arg)}
            /* Construct-5 */} 
          /* ERROR PROTECTION INSERTED (va_arg2-Result) */
          if ErrorIn(va_arg2_try02485) {Result = va_arg2_try02485
          } else {
          va_arg2 = ToList(OBJ(va_arg2_try02485))
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
          var va_arg2_try02525 EID 
          /* Construct:5 */{ 
            var v_bag_arg *ClaireAny  
            va_arg2_try02525= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
            ToList(OBJ(va_arg2_try02525)).AddFast(self.Arg.Id())
            ToList(OBJ(va_arg2_try02525)).AddFast(self.Params.Id())
            var v_bag_arg_try02536 EID 
            /* Iteration:6 */{ 
              var v_list6 *ClaireList  
              var y *ClaireAny  
              var v_local6 *ClaireAny  
              v_list6 = self.Args
              v_bag_arg_try02536 = EID{CreateList(ToType(CEMPTY.Id()),v_list6.Length()).Id(),0}
              for CLcount := 0; CLcount < v_list6.Length(); CLcount++{ 
                y = v_list6.At(CLcount)
                var v_local6_try02548 EID 
                v_local6_try02548 = Core.F_CALL(C_c_code,ARGS(y.ToEID(),EID{C_type.Id(),0}))
                /* ERROR PROTECTION INSERTED (v_local6-v_bag_arg_try02536) */
                if ErrorIn(v_local6_try02548) {v_bag_arg_try02536 = v_local6_try02548
                v_bag_arg_try02536 = v_local6_try02548
                break
                } else {
                v_local6 = ANY(v_local6_try02548)
                ToList(OBJ(v_bag_arg_try02536)).PutAt(CLcount,v_local6)
                } 
              }
              /* Iteration-6 */} 
            /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try02525) */
            if ErrorIn(v_bag_arg_try02536) {va_arg2_try02525 = v_bag_arg_try02536
            } else {
            v_bag_arg = ANY(v_bag_arg_try02536)
            ToList(OBJ(va_arg2_try02525)).AddFast(v_bag_arg)}
            /* Construct-5 */} 
          /* ERROR PROTECTION INSERTED (va_arg2-Result) */
          if ErrorIn(va_arg2_try02525) {Result = va_arg2_try02525
          } else {
          va_arg2 = ToList(OBJ(va_arg2_try02525))
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
        var va_arg2_try02554 EID 
        /* Construct:4 */{ 
          var v_bag_arg *ClaireAny  
          va_arg2_try02554= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
          var v_bag_arg_try02565 EID 
          v_bag_arg_try02565 = Core.F_CALL(C_c_code,ARGS(EID{self.T1.Id(),0},EID{C_type.Id(),0}))
          /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try02554) */
          if ErrorIn(v_bag_arg_try02565) {va_arg2_try02554 = v_bag_arg_try02565
          } else {
          v_bag_arg = ANY(v_bag_arg_try02565)
          ToList(OBJ(va_arg2_try02554)).AddFast(v_bag_arg)
          var v_bag_arg_try02575 EID 
          v_bag_arg_try02575 = Core.F_CALL(C_c_code,ARGS(EID{self.T2.Id(),0},EID{C_type.Id(),0}))
          /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try02554) */
          if ErrorIn(v_bag_arg_try02575) {va_arg2_try02554 = v_bag_arg_try02575
          } else {
          v_bag_arg = ANY(v_bag_arg_try02575)
          ToList(OBJ(va_arg2_try02554)).AddFast(v_bag_arg)}}
          /* Construct-4 */} 
        /* ERROR PROTECTION INSERTED (va_arg2-Result) */
        if ErrorIn(va_arg2_try02554) {Result = va_arg2_try02554
        } else {
        va_arg2 = ToList(OBJ(va_arg2_try02554))
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
      var _Zxt_try02583 EID 
      /* Let:3 */{ 
        var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
        /* noccur = 5 */
        /* update:4 */{ 
          var va_arg1 *Language.Call  
          var va_arg2 *ClaireProperty  
          va_arg1 = _CL_obj
          var va_arg2_try02595 EID 
          var g0260I *ClaireBoolean  
          var g0260I_try02615 EID 
          /* Let:5 */{ 
            var g0262UU *ClaireType  
            /* noccur = 1 */
            var g0262UU_try02636 EID 
            g0262UU_try02636 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
            /* ERROR PROTECTION INSERTED (g0262UU-g0260I_try02615) */
            if ErrorIn(g0262UU_try02636) {g0260I_try02615 = g0262UU_try02636
            } else {
            g0262UU = ToType(OBJ(g0262UU_try02636))
            g0260I_try02615 = EID{g0262UU.Included(ToType(C_object.Id())).Id(),0}
            }
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (g0260I-va_arg2_try02595) */
          if ErrorIn(g0260I_try02615) {va_arg2_try02595 = g0260I_try02615
          } else {
          g0260I = ToBoolean(OBJ(g0260I_try02615))
          if (g0260I == CTRUE) /* If:5 */{ 
            va_arg2_try02595 = EID{C_isa.Id(),0}
            } else {
            va_arg2_try02595 = EID{Core.C_owner.Id(),0}
            /* If-5 */} 
          }
          /* ERROR PROTECTION INSERTED (va_arg2-_Zxt_try02583) */
          if ErrorIn(va_arg2_try02595) {_Zxt_try02583 = va_arg2_try02595
          } else {
          va_arg2 = ToProperty(OBJ(va_arg2_try02595))
          /* ---------- now we compile update selector(va_arg1) := va_arg2 ------- */
          va_arg1.Selector = va_arg2
          _Zxt_try02583 = EID{va_arg2.Id(),0}
          }
          /* update-4 */} 
        /* ERROR PROTECTION INSERTED (_Zxt_try02583-_Zxt_try02583) */
        if !ErrorIn(_Zxt_try02583) {
        _CL_obj.Args = MakeConstantList(x)
        _Zxt_try02583 = EID{_CL_obj.Id(),0}
        }
        /* Let-3 */} 
      /* ERROR PROTECTION INSERTED (_Zxt-Result) */
      if ErrorIn(_Zxt_try02583) {Result = _Zxt_try02583
      } else {
      _Zxt = Language.To_Call(OBJ(_Zxt_try02583))
      if (((self.Open <= -1) || 
            (self.Open == 1)) && 
          (F_boolean_I_any(self.Subclass.Id()).Id() != CTRUE.Id())) /* If:3 */{ 
        /* Let:4 */{ 
          var g0264UU *Language.Call  
          /* noccur = 1 */
          /* Let:5 */{ 
            var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
            /* noccur = 5 */
            _CL_obj.Selector = ToProperty(C__equal.Id())
            _CL_obj.Args = MakeConstantList(self.Id(),_Zxt.Id())
            g0264UU = _CL_obj
            /* Let-5 */} 
          Result = Core.F_CALL(C_c_code,ARGS(EID{g0264UU.Id(),0}))
          /* Let-4 */} 
        } else {
        /* Let:4 */{ 
          var g0265UU *Language.Call  
          /* noccur = 1 */
          /* Let:5 */{ 
            var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
            /* noccur = 5 */
            _CL_obj.Selector = ToProperty(Core.C_inherit_ask.Id())
            _CL_obj.Args = MakeConstantList(_Zxt.Id(),self.Id())
            g0265UU = _CL_obj
            /* Let-5 */} 
          Result = Core.F_CALL(C_c_code,ARGS(EID{g0265UU.Id(),0}))
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
        var va_arg2_try02664 EID 
        /* Construct:4 */{ 
          var v_bag_arg *ClaireAny  
          va_arg2_try02664= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
          var v_bag_arg_try02675 EID 
          v_bag_arg_try02675 = Core.F_CALL(C_c_code,ARGS(x.ToEID(),EID{C_any.Id(),0}))
          /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try02664) */
          if ErrorIn(v_bag_arg_try02675) {va_arg2_try02664 = v_bag_arg_try02675
          } else {
          v_bag_arg = ANY(v_bag_arg_try02675)
          ToList(OBJ(va_arg2_try02664)).AddFast(v_bag_arg)
          var v_bag_arg_try02685 EID 
          v_bag_arg_try02685 = Core.F_CALL(C_c_code,ARGS(EID{self.Id(),0},EID{C_any.Id(),0}))
          /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try02664) */
          if ErrorIn(v_bag_arg_try02685) {va_arg2_try02664 = v_bag_arg_try02685
          } else {
          v_bag_arg = ANY(v_bag_arg_try02685)
          ToList(OBJ(va_arg2_try02664)).AddFast(v_bag_arg)}}
          /* Construct-4 */} 
        /* ERROR PROTECTION INSERTED (va_arg2-Result) */
        if ErrorIn(va_arg2_try02664) {Result = va_arg2_try02664
        } else {
        va_arg2 = ToList(OBJ(va_arg2_try02664))
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
        var va_arg2_try02694 EID 
        /* Construct:4 */{ 
          var v_bag_arg *ClaireAny  
          va_arg2_try02694= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
          var v_bag_arg_try02705 EID 
          v_bag_arg_try02705 = Core.F_CALL(C_Optimize_member_code,ARGS(EID{self.T1.Id(),0},x.ToEID()))
          /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try02694) */
          if ErrorIn(v_bag_arg_try02705) {va_arg2_try02694 = v_bag_arg_try02705
          } else {
          v_bag_arg = ANY(v_bag_arg_try02705)
          ToList(OBJ(va_arg2_try02694)).AddFast(v_bag_arg)
          var v_bag_arg_try02715 EID 
          v_bag_arg_try02715 = Core.F_CALL(C_Optimize_member_code,ARGS(EID{self.T2.Id(),0},x.ToEID()))
          /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try02694) */
          if ErrorIn(v_bag_arg_try02715) {va_arg2_try02694 = v_bag_arg_try02715
          } else {
          v_bag_arg = ANY(v_bag_arg_try02715)
          ToList(OBJ(va_arg2_try02694)).AddFast(v_bag_arg)}}
          /* Construct-4 */} 
        /* ERROR PROTECTION INSERTED (va_arg2-Result) */
        if ErrorIn(va_arg2_try02694) {Result = va_arg2_try02694
        } else {
        va_arg2 = ToList(OBJ(va_arg2_try02694))
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
      var g0272UU *Language.And  
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
        g0272UU = _CL_obj
        /* Let-3 */} 
      Result = Core.F_CALL(C_c_code,ARGS(EID{g0272UU.Id(),0},EID{C_any.Id(),0}))
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: member_code @ Interval (throw: true) 
func E_Optimize_member_code_Interval (self EID,x EID) EID { 
    return /*(sm for member_code @ Interval= EID)*/ F_Optimize_member_code_Interval(To_Interval(OBJ(self)),ANY(x) )} 
  
/* {1} OPT.The go function for: member_code(self:Param,x:any) [] */
func F_Optimize_member_code_Param (self *ClaireParam ,x *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var g0274UU *Language.And  
      /* noccur = 1 */
      /* Let:3 */{ 
        var _CL_obj *Language.And   = Language.To_And(new(Language.And).Is(Language.C_And))
        /* noccur = 21 */
        /* update:4 */{ 
          var va_arg1 *Language.And  
          var va_arg2 *ClaireList  
          va_arg1 = _CL_obj
          /* Let:5 */{ 
            var g0275UU *ClaireList  
            /* noccur = 1 */
            /* Construct:6 */{ 
              var v_bag_arg *ClaireAny  
              g0275UU= ToType(CEMPTY.Id()).EmptyList()
              /* Let:7 */{ 
                var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                /* noccur = 5 */
                _CL_obj.Selector = ToProperty(C__Z.Id())
                _CL_obj.Args = MakeConstantList(x,self.Arg.Id())
                v_bag_arg = _CL_obj.Id()
                /* Let-7 */} 
              g0275UU.AddFast(v_bag_arg)/* Construct-6 */} 
            /* Let:6 */{ 
              var g0276UU *ClaireList  
              /* noccur = 1 */
              /* Let:7 */{ 
                var i_bag *ClaireList   = ToType(CEMPTY.Id()).EmptyList()
                /* noccur = 2 */
                /* Let:8 */{ 
                  var i int  = 1
                  /* noccur = 5 */
                  /* Let:9 */{ 
                    var g0273 int  = self.Params.Length()
                    /* noccur = 1 */
                    for (i <= g0273) /* while:10 */{ 
                      /* Let:11 */{ 
                        var g0277UU *Language.Call  
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
                          g0277UU = _CL_obj
                          /* Let-12 */} 
                        i_bag.AddFast(g0277UU.Id())
                        /* Let-11 */} 
                      i = (i+1)
                      /* while-10 */} 
                    /* Let-9 */} 
                  /* Let-8 */} 
                g0276UU = i_bag
                /* Let-7 */} 
              va_arg2 = g0275UU.Append(g0276UU)
              /* Let-6 */} 
            /* Let-5 */} 
          /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
          va_arg1.Args = va_arg2
          /* update-4 */} 
        g0274UU = _CL_obj
        /* Let-3 */} 
      Result = Core.F_CALL(C_c_code,ARGS(EID{g0274UU.Id(),0},EID{C_any.Id(),0}))
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
          var g0279UU *Language.And  
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
                    var g0278 int  = self.Length()
                    /* noccur = 1 */
                    for (i <= g0278) /* while:10 */{ 
                      /* Let:11 */{ 
                        var g0280UU *Language.Call  
                        /* noccur = 1 */
                        /* Let:12 */{ 
                          var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                          /* noccur = 5 */
                          _CL_obj.Selector = ToProperty(C__Z.Id())
                          _CL_obj.Args = MakeConstantList(ToList(OBJ(Core.F_CALL(C_args,ARGS(x.ToEID())))).At(i-1),ToList(self.Id()).At(i-1))
                          g0280UU = _CL_obj
                          /* Let-12 */} 
                        i_bag.AddFast(g0280UU.Id())
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
            g0279UU = _CL_obj
            /* Let-5 */} 
          Result = Core.F_CALL(C_c_code,ARGS(EID{g0279UU.Id(),0},EID{C_any.Id(),0}))
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
      var _Ztype_try02823 EID 
      /* Construct:3 */{ 
        var v_bag_arg *ClaireAny  
        _Ztype_try02823= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
        var v_bag_arg_try02834 EID 
        v_bag_arg_try02834 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
        /* ERROR PROTECTION INSERTED (v_bag_arg-_Ztype_try02823) */
        if ErrorIn(v_bag_arg_try02834) {_Ztype_try02823 = v_bag_arg_try02834
        } else {
        v_bag_arg = ANY(v_bag_arg_try02834)
        ToList(OBJ(_Ztype_try02823)).AddFast(v_bag_arg)
        var v_bag_arg_try02844 EID 
        v_bag_arg_try02844 = Core.F_CALL(C_c_type,ARGS(self.ToEID()))
        /* ERROR PROTECTION INSERTED (v_bag_arg-_Ztype_try02823) */
        if ErrorIn(v_bag_arg_try02844) {_Ztype_try02823 = v_bag_arg_try02844
        } else {
        v_bag_arg = ANY(v_bag_arg_try02844)
        ToList(OBJ(_Ztype_try02823)).AddFast(v_bag_arg)}}
        /* Construct-3 */} 
      /* ERROR PROTECTION INSERTED (_Ztype-Result) */
      if ErrorIn(_Ztype_try02823) {Result = _Ztype_try02823
      } else {
      _Ztype = ToList(OBJ(_Ztype_try02823))
      /* Let:3 */{ 
        var r *ClaireAny  
        /* noccur = 2 */
        var r_try02854 EID 
        r_try02854 = Language.F_extract_pattern_any(self,CNIL)
        /* ERROR PROTECTION INSERTED (r-Result) */
        if ErrorIn(r_try02854) {Result = r_try02854
        } else {
        r = ANY(r_try02854)
        var g0286I *ClaireBoolean  
        /* or:4 */{ 
          var v_or4 *ClaireBoolean  
          
          v_or4 = Equal(r,CNULL)
          if (v_or4 == CTRUE) {g0286I = CTRUE
          } else /* or:5 */{ 
            v_or4 = Equal(self,C_object.Id())
            if (v_or4 == CTRUE) {g0286I = CTRUE
            } else /* or:6 */{ 
              /* Let:7 */{ 
                var g0287UU *ClaireObject  
                /* noccur = 1 */
                if (self.Isa.IsIn(Core.C_global_variable) == CTRUE) /* If:8 */{ 
                  /* Let:9 */{ 
                    var g0281 *Core.GlobalVariable   = Core.ToGlobalVariable(self)
                    /* noccur = 1 */
                    g0287UU = ToObject(g0281.Range.Id())
                    /* Let-9 */} 
                  } else {
                  g0287UU = ToObject(CFALSE.Id())
                  /* If-8 */} 
                v_or4 = F_boolean_I_any(g0287UU.Id())
                /* Let-7 */} 
              if (v_or4 == CTRUE) {g0286I = CTRUE
              } else /* or:7 */{ 
                g0286I = CFALSE/* org-7 */} 
              /* org-6 */} 
            /* org-5 */} 
          /* or-4 */} 
        if (g0286I == CTRUE) /* If:4 */{ 
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
      
      var v_and2_try02883 EID 
      /* Let:3 */{ 
        var g0289UU *ClaireAny  
        /* noccur = 1 */
        var g0289UU_try02904 EID 
        g0289UU_try02904 = EVAL(ToList(OBJ(Core.F_CALL(C_args,ARGS(y.ToEID())))).At(2-1))
        /* ERROR PROTECTION INSERTED (g0289UU-v_and2_try02883) */
        if ErrorIn(g0289UU_try02904) {v_and2_try02883 = g0289UU_try02904
        } else {
        g0289UU = ANY(g0289UU_try02904)
        v_and2_try02883 = Core.F_CALL(ToProperty(C__inf_equal.Id()),ARGS(x.ToEID(),g0289UU.ToEID()))
        }
        /* Let-3 */} 
      /* ERROR PROTECTION INSERTED (v_and2-Result) */
      if ErrorIn(v_and2_try02883) {Result = v_and2_try02883
      } else {
      v_and2 = ToBoolean(OBJ(v_and2_try02883))
      if (v_and2 == CFALSE) {Result = EID{CFALSE.Id(),0}
      } else /* arg:3 */{ 
        var v_and2_try02914 EID 
        /* Let:4 */{ 
          var g0292UU *ClaireAny  
          /* noccur = 1 */
          var g0292UU_try02935 EID 
          g0292UU_try02935 = EVAL(ToList(OBJ(Core.F_CALL(C_args,ARGS(y.ToEID())))).At(1-1))
          /* ERROR PROTECTION INSERTED (g0292UU-v_and2_try02914) */
          if ErrorIn(g0292UU_try02935) {v_and2_try02914 = g0292UU_try02935
          } else {
          g0292UU = ANY(g0292UU_try02935)
          v_and2_try02914 = Core.F_CALL(ToProperty(C__inf_equal.Id()),ARGS(g0292UU.ToEID(),x.ToEID()))
          }
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (v_and2-Result) */
        if ErrorIn(v_and2_try02914) {Result = v_and2_try02914
        } else {
        v_and2 = ToBoolean(OBJ(v_and2_try02914))
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
      
      var v_and2_try02943 EID 
      /* Let:3 */{ 
        var g0295UU *ClaireAny  
        /* noccur = 1 */
        var g0295UU_try02964 EID 
        g0295UU_try02964 = EVAL(ToList(OBJ(Core.F_CALL(C_args,ARGS(y.ToEID())))).At(1-1))
        /* ERROR PROTECTION INSERTED (g0295UU-v_and2_try02943) */
        if ErrorIn(g0295UU_try02964) {v_and2_try02943 = g0295UU_try02964
        } else {
        g0295UU = ANY(g0295UU_try02964)
        v_and2_try02943 = Core.F_BELONG(x,g0295UU)
        }
        /* Let-3 */} 
      /* ERROR PROTECTION INSERTED (v_and2-Result) */
      if ErrorIn(v_and2_try02943) {Result = v_and2_try02943
      } else {
      v_and2 = ToBoolean(OBJ(v_and2_try02943))
      if (v_and2 == CFALSE) {Result = EID{CFALSE.Id(),0}
      } else /* arg:3 */{ 
        var v_and2_try02974 EID 
        /* Let:4 */{ 
          var g0298UU *ClaireAny  
          /* noccur = 1 */
          var g0298UU_try02995 EID 
          g0298UU_try02995 = EVAL(ToList(OBJ(Core.F_CALL(C_args,ARGS(y.ToEID())))).At(2-1))
          /* ERROR PROTECTION INSERTED (g0298UU-v_and2_try02974) */
          if ErrorIn(g0298UU_try02995) {v_and2_try02974 = g0298UU_try02995
          } else {
          g0298UU = ANY(g0298UU_try02995)
          v_and2_try02974 = EID{Core.F__I_equal_any(x,g0298UU).Id(),0}
          }
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (v_and2-Result) */
        if ErrorIn(v_and2_try02974) {Result = v_and2_try02974
        } else {
        v_and2 = ToBoolean(OBJ(v_and2_try02974))
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
        var g0300 *ClaireType   = ToType(t)
        /* noccur = 1 */
        /* Let:4 */{ 
          var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
          /* noccur = 7 */
          _CL_obj.Pname = s
          _CL_obj.Index = n
          _CL_obj.Range = g0300
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
                      var v_or2_try030711 EID 
                      if (self.Isa.IsIn(Language.C_Call) == CTRUE) /* If:11 */{ 
                        /* Let:12 */{ 
                          var g0302 *Language.Call   = Language.To_Call(self)
                          /* noccur = 2 */
                          /* Let:13 */{ 
                            var x *ClaireAny  
                            /* noccur = 2 */
                            var x_try030814 EID 
                            x_try030814 = Core.F_CALL(C_c_code,ARGS(EID{g0302.Id(),0}))
                            /* ERROR PROTECTION INSERTED (x-v_or2_try030711) */
                            if ErrorIn(x_try030814) {v_or2_try030711 = x_try030814
                            } else {
                            x = ANY(x_try030814)
                            /* or:14 */{ 
                              var v_or14 *ClaireBoolean  
                              
                              var v_or14_try030915 EID 
                              /* and:15 */{ 
                                var v_and15 *ClaireBoolean  
                                
                                v_and15 = x.Isa.IsIn(Language.C_Call).Not
                                if (v_and15 == CFALSE) {v_or14_try030915 = EID{CFALSE.Id(),0}
                                } else /* arg:16 */{ 
                                  var v_and15_try031017 EID 
                                  v_and15_try031017 = F_Compile_designated_ask_any(x)
                                  /* ERROR PROTECTION INSERTED (v_and15-v_or14_try030915) */
                                  if ErrorIn(v_and15_try031017) {v_or14_try030915 = v_and15_try031017
                                  } else {
                                  v_and15 = ToBoolean(OBJ(v_and15_try031017))
                                  if (v_and15 == CFALSE) {v_or14_try030915 = EID{CFALSE.Id(),0}
                                  } else /* arg:17 */{ 
                                    v_or14_try030915 = EID{CTRUE.Id(),0}/* arg-17 */} 
                                  /* arg-16 */} 
                                }
                                /* and-15 */} 
                              /* ERROR PROTECTION INSERTED (v_or14-v_or2_try030711) */
                              if ErrorIn(v_or14_try030915) {v_or2_try030711 = v_or14_try030915
                              } else {
                              v_or14 = ToBoolean(OBJ(v_or14_try030915))
                              if (v_or14 == CTRUE) {v_or2_try030711 = EID{CTRUE.Id(),0}
                              } else /* or:15 */{ 
                                v_or14 = Equal(g0302.Selector.Id(),Core.C_mClaire_get_stack.Id())
                                if (v_or14 == CTRUE) {v_or2_try030711 = EID{CTRUE.Id(),0}
                                } else /* or:16 */{ 
                                  v_or2_try030711 = EID{CFALSE.Id(),0}/* org-16 */} 
                                /* org-15 */} 
                              }
                              /* or-14 */} 
                            }
                            /* Let-13 */} 
                          /* Let-12 */} 
                        /* If!11 */}  else if (self.Isa.IsIn(Language.C_Call_slot) == CTRUE) /* If:11 */{ 
                        /* Let:12 */{ 
                          var g0303 *Language.CallSlot   = Language.To_CallSlot(self)
                          /* noccur = 1 */
                          v_or2_try030711 = F_Compile_designated_ask_any(g0303.Arg)
                          /* Let-12 */} 
                        /* If!11 */}  else if (self.Isa.IsIn(Language.C_Call_table) == CTRUE) /* If:11 */{ 
                        /* Let:12 */{ 
                          var g0304 *Language.CallTable   = Language.To_CallTable(self)
                          /* noccur = 1 */
                          v_or2_try030711 = F_Compile_designated_ask_any(g0304.Arg)
                          /* Let-12 */} 
                        /* If!11 */}  else if (self.Isa.IsIn(Language.C_Call_method) == CTRUE) /* If:11 */{ 
                        /* Let:12 */{ 
                          var g0305 *Language.CallMethod   = Language.To_CallMethod(self)
                          /* noccur = 3 */
                          /* and:13 */{ 
                            var v_and13 *ClaireBoolean  
                            
                            v_and13 = MakeBoolean((C_OPT.SimpleOperations.Contain_ask(g0305.Arg.Selector.Id()) == CTRUE) || (g0305.Arg.Id() == Core.F__at_property1(C_nth,C_list).Id()))
                            if (v_and13 == CFALSE) {v_or2_try030711 = EID{CFALSE.Id(),0}
                            } else /* arg:14 */{ 
                              var v_and13_try031115 EID 
                              /* Let:15 */{ 
                                var g0312UU *ClaireAny  
                                /* noccur = 1 */
                                var g0312UU_try031316 EID 
                                /* For:16 */{ 
                                  var y *ClaireAny  
                                  _ = y
                                  g0312UU_try031316= EID{CFALSE.Id(),0}
                                  var y_support *ClaireList  
                                  y_support = g0305.Args
                                  y_len := y_support.Length()
                                  for i_it := 0; i_it < y_len; i_it++ { 
                                    y = y_support.At(i_it)
                                    var void_try18 EID 
                                    _ = void_try18
                                    var g0314I *ClaireBoolean  
                                    var g0314I_try031518 EID 
                                    /* Let:18 */{ 
                                      var g0316UU *ClaireBoolean  
                                      /* noccur = 1 */
                                      var g0316UU_try031719 EID 
                                      g0316UU_try031719 = F_Compile_designated_ask_any(y)
                                      /* ERROR PROTECTION INSERTED (g0316UU-g0314I_try031518) */
                                      if ErrorIn(g0316UU_try031719) {g0314I_try031518 = g0316UU_try031719
                                      } else {
                                      g0316UU = ToBoolean(OBJ(g0316UU_try031719))
                                      g0314I_try031518 = EID{g0316UU.Not.Id(),0}
                                      }
                                      /* Let-18 */} 
                                    /* ERROR PROTECTION INSERTED (g0314I-void_try18) */
                                    if ErrorIn(g0314I_try031518) {void_try18 = g0314I_try031518
                                    } else {
                                    g0314I = ToBoolean(OBJ(g0314I_try031518))
                                    if (g0314I == CTRUE) /* If:18 */{ 
                                       /*v = g0312UU_try031316, s =EID*/
g0312UU_try031316 = EID{CTRUE.Id(),0}
                                      break
                                      } else {
                                      void_try18 = EID{CFALSE.Id(),0}
                                      /* If-18 */} 
                                    }
                                    /* ERROR PROTECTION INSERTED (void_try18-g0312UU_try031316) */
                                    if ErrorIn(void_try18) {g0312UU_try031316 = void_try18
                                    g0312UU_try031316 = void_try18
                                    break
                                    } else {
                                    }
                                    /* loop-17 */} 
                                  /* For-16 */} 
                                /* ERROR PROTECTION INSERTED (g0312UU-v_and13_try031115) */
                                if ErrorIn(g0312UU_try031316) {v_and13_try031115 = g0312UU_try031316
                                } else {
                                g0312UU = ANY(g0312UU_try031316)
                                v_and13_try031115 = EID{Core.F_not_any(g0312UU).Id(),0}
                                }
                                /* Let-15 */} 
                              /* ERROR PROTECTION INSERTED (v_and13-v_or2_try030711) */
                              if ErrorIn(v_and13_try031115) {v_or2_try030711 = v_and13_try031115
                              } else {
                              v_and13 = ToBoolean(OBJ(v_and13_try031115))
                              if (v_and13 == CFALSE) {v_or2_try030711 = EID{CFALSE.Id(),0}
                              } else /* arg:15 */{ 
                                v_or2_try030711 = EID{CTRUE.Id(),0}/* arg-15 */} 
                              /* arg-14 */} 
                            }
                            /* and-13 */} 
                          /* Let-12 */} 
                        } else {
                        v_or2_try030711 = EID{CFALSE.Id(),0}
                        /* If-11 */} 
                      /* ERROR PROTECTION INSERTED (v_or2-Result) */
                      if ErrorIn(v_or2_try030711) {Result = v_or2_try030711
                      } else {
                      v_or2 = ToBoolean(OBJ(v_or2_try030711))
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
        var v_or2_try03184 EID 
        /* Let:4 */{ 
          var t *ClaireClass  
          /* noccur = 1 */
          var t_try03195 EID 
          /* Let:5 */{ 
            var g0320UU *ClaireType  
            /* noccur = 1 */
            var g0320UU_try03216 EID 
            g0320UU_try03216 = Core.F_CALL(C_c_type,ARGS(self.ToEID()))
            /* ERROR PROTECTION INSERTED (g0320UU-t_try03195) */
            if ErrorIn(g0320UU_try03216) {t_try03195 = g0320UU_try03216
            } else {
            g0320UU = ToType(OBJ(g0320UU_try03216))
            t_try03195 = EID{g0320UU.Class_I().Id(),0}
            }
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (t-v_or2_try03184) */
          if ErrorIn(t_try03195) {v_or2_try03184 = t_try03195
          } else {
          t = ToClass(OBJ(t_try03195))
          v_or2_try03184 = EID{C_OPT.NonIdentifiableSet.Contain_ask(t.Id()).Not.Id(),0}
          }
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (v_or2-Result) */
        if ErrorIn(v_or2_try03184) {Result = v_or2_try03184
        } else {
        v_or2 = ToBoolean(OBJ(v_or2_try03184))
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
      var g0322UU *ClaireAny  
      /* noccur = 1 */
      var g0322UU_try03233 EID 
      g0322UU_try03233 = F_Optimize_c_inline_method2(self,l)
      /* ERROR PROTECTION INSERTED (g0322UU-Result) */
      if ErrorIn(g0322UU_try03233) {Result = g0322UU_try03233
      } else {
      g0322UU = ANY(g0322UU_try03233)
      Result = Core.F_CALL(C_c_code,ARGS(g0322UU.ToEID(),EID{s.Id(),0}))
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
          var lbv_try03255 EID 
          lbv_try03255 = F_Optimize_bound_variables_any(x)
          /* ERROR PROTECTION INSERTED (lbv-Result) */
          if ErrorIn(lbv_try03255) {Result = lbv_try03255
          } else {
          lbv = ToList(OBJ(lbv_try03255))
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
        var g0326 *Language.Call   = Language.To_Call(self)
        /* noccur = 2 */
        /* Let:4 */{ 
          var l *ClaireList   = g0326.Args
          /* noccur = 3 */
          /* Let:5 */{ 
            var m *ClaireAny  
            /* noccur = 2 */
            var m_try03306 EID 
            /* Let:6 */{ 
              var g0331UU *ClaireList  
              /* noccur = 1 */
              var g0331UU_try03327 EID 
              /* Iteration:7 */{ 
                var v_list7 *ClaireList  
                var x *ClaireAny  
                var v_local7 *ClaireAny  
                v_list7 = l
                g0331UU_try03327 = EID{CreateList(ToType(CEMPTY.Id()),v_list7.Length()).Id(),0}
                for CLcount := 0; CLcount < v_list7.Length(); CLcount++{ 
                  x = v_list7.At(CLcount)
                  var v_local7_try03339 EID 
                  v_local7_try03339 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
                  /* ERROR PROTECTION INSERTED (v_local7-g0331UU_try03327) */
                  if ErrorIn(v_local7_try03339) {g0331UU_try03327 = v_local7_try03339
                  g0331UU_try03327 = v_local7_try03339
                  break
                  } else {
                  v_local7 = ANY(v_local7_try03339)
                  ToList(OBJ(g0331UU_try03327)).PutAt(CLcount,v_local7)
                  } 
                }
                /* Iteration-7 */} 
              /* ERROR PROTECTION INSERTED (g0331UU-m_try03306) */
              if ErrorIn(g0331UU_try03327) {m_try03306 = g0331UU_try03327
              } else {
              g0331UU = ToList(OBJ(g0331UU_try03327))
              m_try03306 = F_Optimize_restriction_I_property(g0326.Selector,g0331UU,CTRUE).ToEID()
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (m-Result) */
            if ErrorIn(m_try03306) {Result = m_try03306
            } else {
            m = ANY(m_try03306)
            if (C_method.Id() == m.Isa.Id()) /* If:6 */{ 
              /* Let:7 */{ 
                var g0327 *ClaireMethod   = ToMethod(m)
                /* noccur = 3 */
                var g0334I *ClaireBoolean  
                var g0334I_try03358 EID 
                /* and:8 */{ 
                  var v_and8 *ClaireBoolean  
                  
                  v_and8 = g0327.Inline_ask
                  if (v_and8 == CFALSE) {g0334I_try03358 = EID{CFALSE.Id(),0}
                  } else /* arg:9 */{ 
                    var v_and8_try033610 EID 
                    v_and8_try033610 = F_Optimize_c_inline_ask_method(g0327,l)
                    /* ERROR PROTECTION INSERTED (v_and8-g0334I_try03358) */
                    if ErrorIn(v_and8_try033610) {g0334I_try03358 = v_and8_try033610
                    } else {
                    v_and8 = ToBoolean(OBJ(v_and8_try033610))
                    if (v_and8 == CFALSE) {g0334I_try03358 = EID{CFALSE.Id(),0}
                    } else /* arg:10 */{ 
                      g0334I_try03358 = EID{CTRUE.Id(),0}/* arg-10 */} 
                    /* arg-9 */} 
                  }
                  /* and-8 */} 
                /* ERROR PROTECTION INSERTED (g0334I-Result) */
                if ErrorIn(g0334I_try03358) {Result = g0334I_try03358
                } else {
                g0334I = ToBoolean(OBJ(g0334I_try03358))
                if (g0334I == CTRUE) /* If:8 */{ 
                  Result = F_Optimize_c_inline_method2(g0327,l)
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
        var g0337UU *Language.Call  
        /* noccur = 1 */
        /* Let:4 */{ 
          var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
          /* noccur = 5 */
          _CL_obj.Selector = C_set_I
          _CL_obj.Args = MakeConstantList(self)
          g0337UU = _CL_obj
          /* Let-4 */} 
        Result = F_Optimize_c_inline_arg_ask_any(g0337UU.Id())
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
        var g0338 *ClaireVariable   = To_Variable(self)
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
                var g0339 int  = lx.Length()
                /* noccur = 1 */
                for (j <= g0339) /* while:8 */{ 
                  if (g0338.Pname.Id() == To_Variable(lx.At(j-1)).Pname.Id()) /* If:9 */{ 
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
            Result = EID{g0338.Id(),0}
            /* If-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(C_list) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0340 *ClaireList   = ToList(self)
        /* noccur = 4 */
        /* Let:4 */{ 
          var i int  = 1
          /* noccur = 5 */
          /* Let:5 */{ 
            var g0341 int  = g0340.Length()
            /* noccur = 1 */
            Result= EID{CFALSE.Id(),0}
            for (i <= g0341) /* while:6 */{ 
              var void_try7 EID 
              _ = void_try7
              { 
              /* Let:7 */{ 
                var g0345UU *ClaireAny  
                /* noccur = 1 */
                var g0345UU_try03468 EID 
                g0345UU_try03468 = F_Optimize_c_substitution_any(g0340.At(i-1),lx,val,eval_ask)
                /* ERROR PROTECTION INSERTED (g0345UU-void_try7) */
                if ErrorIn(g0345UU_try03468) {void_try7 = g0345UU_try03468
                } else {
                g0345UU = ANY(g0345UU_try03468)
                void_try7 = ToArray(g0340.Id()).NthPut(i,g0345UU).ToEID()
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
        Result = EID{g0340.Id(),0}
        }
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(Language.C_Call) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0342 *Language.Call   = Language.To_Call(self)
        /* noccur = 11 */
        if (g0342.Selector.Id() == Core.C_eval.Id()) /* If:4 */{ 
          /* Let:5 */{ 
            var g0347UU *ClaireBoolean  
            /* noccur = 1 */
            var g0347UU_try03486 EID 
            /* or:6 */{ 
              var v_or6 *ClaireBoolean  
              
              v_or6 = Equal(MakeInteger(g0342.Args.Length()).Id(),MakeInteger(1).Id())
              if (v_or6 == CTRUE) {g0347UU_try03486 = EID{CTRUE.Id(),0}
              } else /* or:7 */{ 
                var v_or6_try03498 EID 
                /* and:8 */{ 
                  var v_and8 *ClaireBoolean  
                  
                  v_and8 = Equal(MakeInteger(g0342.Args.Length()).Id(),MakeInteger(2).Id())
                  if (v_and8 == CFALSE) {v_or6_try03498 = EID{CFALSE.Id(),0}
                  } else /* arg:9 */{ 
                    var v_and8_try035010 EID 
                    v_and8_try035010 = Core.F_BELONG(val.At(1-1),g0342.Args.At(2-1))
                    /* ERROR PROTECTION INSERTED (v_and8-v_or6_try03498) */
                    if ErrorIn(v_and8_try035010) {v_or6_try03498 = v_and8_try035010
                    } else {
                    v_and8 = ToBoolean(OBJ(v_and8_try035010))
                    if (v_and8 == CFALSE) {v_or6_try03498 = EID{CFALSE.Id(),0}
                    } else /* arg:10 */{ 
                      v_or6_try03498 = EID{CTRUE.Id(),0}/* arg-10 */} 
                    /* arg-9 */} 
                  }
                  /* and-8 */} 
                /* ERROR PROTECTION INSERTED (v_or6-g0347UU_try03486) */
                if ErrorIn(v_or6_try03498) {g0347UU_try03486 = v_or6_try03498
                } else {
                v_or6 = ToBoolean(OBJ(v_or6_try03498))
                if (v_or6 == CTRUE) {g0347UU_try03486 = EID{CTRUE.Id(),0}
                } else /* or:8 */{ 
                  g0347UU_try03486 = EID{CFALSE.Id(),0}/* org-8 */} 
                /* org-7 */} 
              }
              /* or-6 */} 
            /* ERROR PROTECTION INSERTED (g0347UU-Result) */
            if ErrorIn(g0347UU_try03486) {Result = g0347UU_try03486
            } else {
            g0347UU = ToBoolean(OBJ(g0347UU_try03486))
            Result = F_Optimize_c_substitution_any(g0342.Args.At(1-1),lx,val,g0347UU)
            }
            /* Let-5 */} 
          /* If!4 */}  else if (eval_ask == CTRUE) /* If:4 */{ 
          h_index := ClEnv.Index /* Handle */
          h_base := ClEnv.Base
          /* Let:5 */{ 
            var g0351UU *ClaireList  
            /* noccur = 1 */
            var g0351UU_try03526 EID 
            /* Iteration:6 */{ 
              var v_list6 *ClaireList  
              var y *ClaireAny  
              var v_local6 *ClaireAny  
              v_list6 = g0342.Args
              g0351UU_try03526 = EID{CreateList(ToType(CEMPTY.Id()),v_list6.Length()).Id(),0}
              for CLcount := 0; CLcount < v_list6.Length(); CLcount++{ 
                y = v_list6.At(CLcount)
                var v_local6_try03538 EID 
                v_local6_try03538 = F_Optimize_c_substitution_any(y,lx,val,CTRUE)
                /* ERROR PROTECTION INSERTED (v_local6-g0351UU_try03526) */
                if ErrorIn(v_local6_try03538) {g0351UU_try03526 = v_local6_try03538
                g0351UU_try03526 = v_local6_try03538
                break
                } else {
                v_local6 = ANY(v_local6_try03538)
                ToList(OBJ(g0351UU_try03526)).PutAt(CLcount,v_local6)
                } 
              }
              /* Iteration-6 */} 
            /* ERROR PROTECTION INSERTED (g0351UU-Result) */
            if ErrorIn(g0351UU_try03526) {Result = g0351UU_try03526
            } else {
            g0351UU = ToList(OBJ(g0351UU_try03526))
            Result = Core.F_apply_property(g0342.Selector,g0351UU)
            }
            /* Let-5 */} 
          if ErrorIn(Result){ 
            /* s=EID */ClEnv.Index = h_index
            ClEnv.Base = h_base
            Core.F_tformat_string(MakeString("a strange problem happens ~A \n"),0,MakeConstantList(MakeInteger(ClEnv.Verbose).Id()))
            F_Compile_warn_void()
            Core.F_tformat_string(MakeString("failed substitution: ~S"),2,MakeConstantList(ClEnv.Exception_I.Id()))
            Result = F_Optimize_c_substitution_any(g0342.Args.Id(),lx,val,CFALSE)
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            Result = EID{g0342.Id(),0}
            }
            } 
          } else {
          Result = F_Optimize_c_substitution_any(g0342.Args.Id(),lx,val,CFALSE)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          Result = EID{g0342.Id(),0}
          }
          /* If-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(C_Instruction) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0343 *ClaireInstruction   = To_Instruction(self)
        /* noccur = 4 */
        /* For:4 */{ 
          var s *ClaireAny  
          _ = s
          Result= EID{CFALSE.Id(),0}
          var s_support *ClaireList  
          s_support = g0343.Id().Isa.Slots
          for _,s = range(s_support.ValuesO())/* loop2:5 */{ 
            var void_try6 EID 
            _ = void_try6
            /* Let:6 */{ 
              var y *ClaireAny   = Core.F_get_slot(ToSlot(s),ToObject(g0343.Id()))
              /* noccur = 1 */
              /* Let:7 */{ 
                var g0354UU *ClaireAny  
                /* noccur = 1 */
                var g0354UU_try03558 EID 
                g0354UU_try03558 = F_Optimize_c_substitution_any(y,lx,val,eval_ask)
                /* ERROR PROTECTION INSERTED (g0354UU-void_try6) */
                if ErrorIn(g0354UU_try03558) {void_try6 = g0354UU_try03558
                } else {
                g0354UU = ANY(g0354UU_try03558)
                void_try6 = Core.F_put_slot(ToSlot(s),ToObject(g0343.Id()),g0354UU).ToEID()
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
        Result = EID{g0343.Id(),0}
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
          var g0356 *Language.InstructionWithVar   = Language.To_InstructionWithVar(self)
          /* noccur = 1 */
          l = MakeList(ToType(C_any.Id()),g0356.ClaireVar.Id())
          /* Let-4 */} 
        /* If-3 */} 
      if (self.Isa.IsIn(C_Variable) == CTRUE) /* If:3 */{ 
        Result = EID{CNIL.Id(),0}
        /* If!3 */}  else if (self.Isa.IsIn(C_Instruction) == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var g0358 *ClaireInstruction   = To_Instruction(self)
          /* noccur = 2 */
          /* For:5 */{ 
            var s *ClaireAny  
            _ = s
            Result= EID{CFALSE.Id(),0}
            for _,s = range(g0358.Isa.Slots.ValuesO())/* loop:6 */{ 
              var void_try7 EID 
              _ = void_try7
              var l_try03607 EID 
              /* Let:7 */{ 
                var g0361UU *ClaireList  
                /* noccur = 1 */
                var g0361UU_try03628 EID 
                g0361UU_try03628 = F_Optimize_bound_variables_any(Core.F_get_slot(ToSlot(s),ToObject(g0358.Id())))
                /* ERROR PROTECTION INSERTED (g0361UU-l_try03607) */
                if ErrorIn(g0361UU_try03628) {l_try03607 = g0361UU_try03628
                } else {
                g0361UU = ToList(OBJ(g0361UU_try03628))
                l_try03607 = l.Add_star(g0361UU)
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (l-Result) */
              if ErrorIn(l_try03607) {Result = l_try03607
              Result = l_try03607
              break
              } else {
              l = ToList(OBJ(l_try03607))
              void_try7 = EID{l.Id(),0}
              }
              /* loop-6 */} 
            /* For-5 */} 
          /* Let-4 */} 
        /* If!3 */}  else if (self.Isa.IsIn(C_list) == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var g0359 *ClaireList   = ToList(self)
          /* noccur = 1 */
          /* For:5 */{ 
            var x *ClaireAny  
            _ = x
            Result= EID{CFALSE.Id(),0}
            var x_support *ClaireList  
            x_support = g0359
            x_len := x_support.Length()
            for i_it := 0; i_it < x_len; i_it++ { 
              x = x_support.At(i_it)
              var void_try7 EID 
              _ = void_try7
              var l_try03637 EID 
              /* Let:7 */{ 
                var g0364UU *ClaireList  
                /* noccur = 1 */
                var g0364UU_try03658 EID 
                g0364UU_try03658 = F_Optimize_bound_variables_any(x)
                /* ERROR PROTECTION INSERTED (g0364UU-l_try03637) */
                if ErrorIn(g0364UU_try03658) {l_try03637 = g0364UU_try03658
                } else {
                g0364UU = ToList(OBJ(g0364UU_try03658))
                l_try03637 = l.Add_star(g0364UU)
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (l-Result) */
              if ErrorIn(l_try03637) {Result = l_try03637
              Result = l_try03637
              break
              } else {
              l = ToList(OBJ(l_try03637))
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
      var tx_try03673 EID 
      tx_try03673 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
      /* ERROR PROTECTION INSERTED (tx-Result) */
      if ErrorIn(tx_try03673) {Result = tx_try03673
      } else {
      tx = ToType(OBJ(tx_try03673))
      /* Let:3 */{ 
        var ptx *ClaireType   = F_Optimize_ptype_type(tx)
        /* noccur = 1 */
        if (ptx.Included(ToType(C_boolean.Id())) == CTRUE) /* If:4 */{ 
          if (x.Isa.IsIn(Language.C_Call) == CTRUE) /* If:5 */{ 
            /* Let:6 */{ 
              var g0366 *Language.Call   = Language.To_Call(x)
              /* noccur = 5 */
              var g0368I *ClaireBoolean  
              var g0368I_try03697 EID 
              /* and:7 */{ 
                var v_and7 *ClaireBoolean  
                
                v_and7 = Equal(g0366.Selector.Id(),Core.C_not.Id())
                if (v_and7 == CFALSE) {g0368I_try03697 = EID{CFALSE.Id(),0}
                } else /* arg:8 */{ 
                  var v_and7_try03709 EID 
                  /* Let:9 */{ 
                    var g0371UU *ClaireType  
                    /* noccur = 1 */
                    var g0371UU_try037210 EID 
                    /* Let:10 */{ 
                      var g0373UU *ClaireType  
                      /* noccur = 1 */
                      var g0373UU_try037411 EID 
                      g0373UU_try037411 = Core.F_CALL(C_c_type,ARGS(g0366.Args.At(1-1).ToEID()))
                      /* ERROR PROTECTION INSERTED (g0373UU-g0371UU_try037210) */
                      if ErrorIn(g0373UU_try037411) {g0371UU_try037210 = g0373UU_try037411
                      } else {
                      g0373UU = ToType(OBJ(g0373UU_try037411))
                      g0371UU_try037210 = EID{F_Optimize_ptype_type(g0373UU).Id(),0}
                      }
                      /* Let-10 */} 
                    /* ERROR PROTECTION INSERTED (g0371UU-v_and7_try03709) */
                    if ErrorIn(g0371UU_try037210) {v_and7_try03709 = g0371UU_try037210
                    } else {
                    g0371UU = ToType(OBJ(g0371UU_try037210))
                    v_and7_try03709 = EID{Core.F__I_equal_any(g0371UU.Id(),C_boolean.Id()).Id(),0}
                    }
                    /* Let-9 */} 
                  /* ERROR PROTECTION INSERTED (v_and7-g0368I_try03697) */
                  if ErrorIn(v_and7_try03709) {g0368I_try03697 = v_and7_try03709
                  } else {
                  v_and7 = ToBoolean(OBJ(v_and7_try03709))
                  if (v_and7 == CFALSE) {g0368I_try03697 = EID{CFALSE.Id(),0}
                  } else /* arg:9 */{ 
                    g0368I_try03697 = EID{CTRUE.Id(),0}/* arg-9 */} 
                  /* arg-8 */} 
                }
                /* and-7 */} 
              /* ERROR PROTECTION INSERTED (g0368I-Result) */
              if ErrorIn(g0368I_try03697) {Result = g0368I_try03697
              } else {
              g0368I = ToBoolean(OBJ(g0368I_try03697))
              if (g0368I == CTRUE) /* If:7 */{ 
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
                        _CL_obj.Args = MakeConstantList(g0366.Args.At(1-1))
                        v_bag_arg = _CL_obj.Id()
                        /* Let-11 */} 
                      va_arg2.AddFast(v_bag_arg)
                      va_arg2.AddFast(CTRUE.Id())/* Construct-10 */} 
                    /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                    va_arg1.Args = va_arg2
                    /* update-9 */} 
                  g0366 = _CL_obj
                  /* Let-8 */} 
                Result = EID{g0366.Id(),0}
                } else {
                Result = EID{CFALSE.Id(),0}
                /* If-7 */} 
              }
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              x = g0366.Id()
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
            var g0375UU *Language.Call  
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
              g0375UU = _CL_obj
              /* Let-6 */} 
            Result = Core.F_CALL(C_c_code,ARGS(EID{g0375UU.Id(),0}))
            /* Let-5 */} 
          } else {
          /* Let:5 */{ 
            var g0376UU *Language.Call  
            /* noccur = 1 */
            /* Let:6 */{ 
              var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
              /* noccur = 5 */
              _CL_obj.Selector = C_boolean_I
              _CL_obj.Args = MakeConstantList(x)
              g0376UU = _CL_obj
              /* Let-6 */} 
            Result = Core.F_CALL(C_c_code,ARGS(EID{g0376UU.Id(),0}))
            /* Let-5 */} 
          /* If-4 */} 
        /* Let-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_boolean @ any (throw: true) 
func E_Optimize_c_boolean_any (x EID) EID { 
    return /*(sm for c_boolean @ any= EID)*/ F_Optimize_c_boolean_any(ANY(x) )} 
  