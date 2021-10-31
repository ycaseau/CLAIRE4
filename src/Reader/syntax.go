/***** CLAIRE Compilation of file /Users/ycaseau/Dropbox/src/clairev4.0/src/meta/syntax.cl 
         [version 4.0.02 / safety 5] Saturday 10-30-2021 *****/

package Reader
import (_ "fmt"
	. "Kernel"
	"Core"
	"Language"
)

//-------- dumb function to prevent import errors --------
func import_g0141() { 
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
/* {1} OPT.The go function for: operation?(y:any) [] */
func F_operation_ask_any (y *ClaireAny ) *ClaireBoolean  { 
    // use function body compiling 
return  MakeBoolean((y == C_as.Id()) || 
    (y == C_L__equal.Id()) || 
    (y.Isa.IsIn(C_operation) == CTRUE) || 
    (Equal(y,C_OR.Value) == CTRUE) || 
    (y == C__Z.Id()) || 
    (y == C_add.Id()))
    } 
  
// The EID go function for: operation? @ any (throw: false) 
func E_operation_ask_any (y EID) EID { 
    return EID{/*(sm for operation? @ any= boolean)*/ F_operation_ask_any(ANY(y) ).Id(),0}} 
  
// produce an expression from an operation
// apply precedence rules ((x1 op x2) y  z) -> x1 op (x2 y z)
/* {1} OPT.The go function for: combine(x:any,y:any,z:any) [] */
func F_combine_any (x *ClaireAny ,y *ClaireAny ,z *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var p *ClaireAny   = F_operation_I_any(x)
      /* noccur = 3 */
      if ((F_boolean_I_any(p) == CTRUE) && 
          (F_precedence_I_any(y) < F_precedence_I_any(p))) /* If:3 */{ 
        /* Let:4 */{ 
          var g0142UU *ClaireAny  
          /* noccur = 1 */
          var g0142UU_try01445 EID 
          g0142UU_try01445 = F_operand_I_any(x,1)
          /* ERROR PROTECTION INSERTED (g0142UU-Result) */
          if ErrorIn(g0142UU_try01445) {Result = g0142UU_try01445
          } else {
          g0142UU = ANY(g0142UU_try01445)
          /* Let:5 */{ 
            var g0143UU *ClaireAny  
            /* noccur = 1 */
            var g0143UU_try01456 EID 
            /* Let:6 */{ 
              var g0146UU *ClaireAny  
              /* noccur = 1 */
              var g0146UU_try01477 EID 
              g0146UU_try01477 = F_operand_I_any(x,2)
              /* ERROR PROTECTION INSERTED (g0146UU-g0143UU_try01456) */
              if ErrorIn(g0146UU_try01477) {g0143UU_try01456 = g0146UU_try01477
              } else {
              g0146UU = ANY(g0146UU_try01477)
              g0143UU_try01456 = F_combine_any(g0146UU,y,z)
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (g0143UU-Result) */
            if ErrorIn(g0143UU_try01456) {Result = g0143UU_try01456
            } else {
            g0143UU = ANY(g0143UU_try01456)
            Result = F_combine_I_any(g0142UU,p,g0143UU)
            }
            /* Let-5 */} 
          }
          /* Let-4 */} 
        } else {
        Result = F_combine_I_any(x,y,z)
        /* If-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: combine @ any (throw: true) 
func E_combine_any (x EID,y EID,z EID) EID { 
    return /*(sm for combine @ any= EID)*/ F_combine_any(ANY(x),ANY(y),ANY(z) )} 
  
// produces x op=y z
// replace r(x) :add y with add(r,x,y) for multivalued or defeasible .. also with delete
/* {1} OPT.The go function for: combine!(x:any,y:any,z:any) [] */
func F_combine_I_any (x *ClaireAny ,y *ClaireAny ,z *ClaireAny ) EID { 
    var Result EID 
    if (y == C_as.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var _CL_obj *Language.Cast   = Language.To_Cast(new(Language.Cast).Is(Language.C_Cast))
        /* noccur = 4 */
        _CL_obj.Arg = x
        /* update:4 */{ 
          var va_arg1 *Language.Cast  
          var va_arg2 *ClaireType  
          va_arg1 = _CL_obj
          var va_arg2_try01615 EID 
          va_arg2_try01615 = Language.F_extract_type_any(z)
          /* ERROR PROTECTION INSERTED (va_arg2-Result) */
          if ErrorIn(va_arg2_try01615) {Result = va_arg2_try01615
          } else {
          va_arg2 = ToType(OBJ(va_arg2_try01615))
          /* ---------- now we compile update iClaire/set_arg(va_arg1) := va_arg2 ------- */
          va_arg1.SetArg = va_arg2
          Result = EID{va_arg2.Id(),0}
          }
          /* update-4 */} 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        Result = EID{_CL_obj.Id(),0}
        }
        /* Let-3 */} 
      /* If!2 */}  else if (y == C_L__equal.Id()) /* If:2 */{ 
      if (x.Isa.IsIn(Core.C_global_variable) == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var g0148 *Core.GlobalVariable   = Core.ToGlobalVariable(x)
          /* noccur = 1 */
          /* Let:5 */{ 
            var _CL_obj *Language.Gassign   = Language.To_Gassign(new(Language.Gassign).Is(Language.C_Gassign))
            /* noccur = 4 */
            _CL_obj.ClaireVar = g0148
            _CL_obj.Arg = z
            Result = EID{_CL_obj.Id(),0}
            /* Let-5 */} 
          /* Let-4 */} 
        /* If!3 */}  else if (x.Isa.IsIn(Language.C_Call) == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var g0149 *Language.Call   = Language.To_Call(x)
          /* noccur = 18 */
          var g0162I *ClaireBoolean  
          if (z.Isa.IsIn(Language.C_Call) == CTRUE) /* If:5 */{ 
            /* Let:6 */{ 
              var g0150 *Language.Call   = Language.To_Call(z)
              /* noccur = 3 */
              /* and:7 */{ 
                var v_and7 *ClaireBoolean  
                
                v_and7 = MakeBoolean((g0150.Selector.Id() == C_add.Id()) || (g0150.Selector.Id() == C_delete.Id()))
                if (v_and7 == CFALSE) {g0162I = CFALSE
                } else /* arg:8 */{ 
                  v_and7 = Equal(g0150.Args.At(1-1),g0149.Id())
                  if (v_and7 == CFALSE) {g0162I = CFALSE
                  } else /* arg:9 */{ 
                    if (g0149.Args.Length() == 1) /* If:10 */{ 
                      v_and7 = g0149.Selector.Multivalued_ask
                      /* If!10 */}  else if ((g0149.Selector.Id() == C_nth.Id()) && 
                        (g0149.Args.Length() == 2)) /* If:10 */{ 
                      /* Let:11 */{ 
                        var p *ClaireAny   = g0149.Args.At(1-1)
                        /* noccur = 2 */
                        if (p.Isa.IsIn(C_relation) == CTRUE) /* If:12 */{ 
                          /* Let:13 */{ 
                            var g0151 *ClaireRelation   = ToRelation(p)
                            /* noccur = 1 */
                            v_and7 = g0151.Multivalued_ask
                            /* Let-13 */} 
                          } else {
                          v_and7 = CFALSE
                          /* If-12 */} 
                        /* Let-11 */} 
                      } else {
                      v_and7 = CFALSE
                      /* If-10 */} 
                    if (v_and7 == CFALSE) {g0162I = CFALSE
                    } else /* arg:10 */{ 
                      g0162I = CTRUE/* arg-10 */} 
                    /* arg-9 */} 
                  /* arg-8 */} 
                /* and-7 */} 
              /* Let-6 */} 
            } else {
            g0162I = CFALSE
            /* If-5 */} 
          if (g0162I == CTRUE) /* If:5 */{ 
            /* Let:6 */{ 
              var g0163UU *ClaireList  
              /* noccur = 1 */
              /* Construct:7 */{ 
                var v_bag_arg *ClaireAny  
                g0163UU= ToType(CEMPTY.Id()).EmptyList()
                if (g0149.Args.Length() == 1) /* If:8 */{ 
                  v_bag_arg = g0149.Selector.Id()
                  } else {
                  v_bag_arg = g0149.Args.At(1-1)
                  /* If-8 */} 
                g0163UU.AddFast(v_bag_arg)
                if (g0149.Args.Length() == 1) /* If:8 */{ 
                  v_bag_arg = g0149.Args.At(1-1)
                  } else {
                  v_bag_arg = g0149.Args.At(2-1)
                  /* If-8 */} 
                g0163UU.AddFast(v_bag_arg)
                g0163UU.AddFast(ToList(OBJ(Core.F_CALL(C_args,ARGS(z.ToEID())))).At(2-1))/* Construct-7 */} 
              Result = F_Call_I_property(ToProperty(OBJ(Core.F_CALL(C_selector,ARGS(z.ToEID())))),g0163UU)
              /* Let-6 */} 
            /* If!5 */}  else if (g0149.Selector.Id() == C_nth.Id()) /* If:5 */{ 
            Result = F_Call_I_property(C_nth_equal,g0149.Args.Copy().AddFast(z))
            /* If!5 */}  else if (g0149.Args.Length() == 1) /* If:5 */{ 
            /* Let:6 */{ 
              var p *ClaireProperty  
              /* noccur = 2 */
              var p_try01647 EID 
              p_try01647 = Language.F_make_a_property_any(g0149.Selector.Id())
              /* ERROR PROTECTION INSERTED (p-Result) */
              if ErrorIn(p_try01647) {Result = p_try01647
              } else {
              p = ToProperty(OBJ(p_try01647))
              /* Let:7 */{ 
                var y *ClaireAny   = g0149.Args.At(1-1)
                /* noccur = 4 */
                var g0165I *ClaireBoolean  
                /* and:8 */{ 
                  var v_and8 *ClaireBoolean  
                  
                  v_and8 = Equal(p.Id(),C_read.Id())
                  if (v_and8 == CFALSE) {g0165I = CFALSE
                  } else /* arg:9 */{ 
                    if (y.Isa.IsIn(Language.C_Call_plus) == CTRUE) /* If:10 */{ 
                      /* Let:11 */{ 
                        var g0153 *Language.Call_plus   = Language.To_Call_plus(y)
                        /* noccur = 1 */
                        v_and8 = Equal(g0153.Selector.Reified.Id(),CTRUE.Id())
                        /* Let-11 */} 
                      } else {
                      v_and8 = CFALSE
                      /* If-10 */} 
                    if (v_and8 == CFALSE) {g0165I = CFALSE
                    } else /* arg:10 */{ 
                      g0165I = CTRUE/* arg-10 */} 
                    /* arg-9 */} 
                  /* and-8 */} 
                if (g0165I == CTRUE) /* If:8 */{ 
                  Result = F_Call_I_property(Core.C_write,MakeConstantList(y,z))
                  } else {
                  Result = F_Call_I_property(Core.C_write,MakeConstantList(p.Id(),y,z))
                  /* If-8 */} 
                /* Let-7 */} 
              }
              /* Let-6 */} 
            } else {
            Result = F_Serror_string(MakeString("[164] ~S cannot be assigned with :="),MakeConstantList(g0149.Id()))
            /* If-5 */} 
          /* Let-4 */} 
        /* If!3 */}  else if (x.Isa.IsIn(Language.C_Do) == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var g0154 *Language.Do   = Language.To_Do(x)
          /* noccur = 1 */
          /* Let:5 */{ 
            var l *ClaireList   = g0154.Args
            /* noccur = 2 */
            /* Let:6 */{ 
              var m int  = l.Length()
              /* noccur = 1 */
              /* Let:7 */{ 
                var v *ClaireVariable  
                /* noccur = 2 */
                /* Let:8 */{ 
                  var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
                  /* noccur = 3 */
                  _CL_obj.Pname = Core.F_gensym_void()
                  v = _CL_obj
                  /* Let-8 */} 
                /* Let:8 */{ 
                  var _CL_obj *Language.Let_star   = Language.To_Let_star(new(Language.Let_star).Is(Language.C_Let_star))
                  /* noccur = 10 */
                  _CL_obj.ClaireVar = v
                  _CL_obj.Value = z
                  /* update:9 */{ 
                    var va_arg1 *Language.Let  
                    var va_arg2 *ClaireAny  
                    va_arg1 = Language.To_Let(_CL_obj.Id())
                    var va_arg2_try016610 EID 
                    /* Let:10 */{ 
                      var _CL_obj *Language.Do   = Language.To_Do(new(Language.Do).Is(Language.C_Do))
                      /* noccur = 3 */
                      /* update:11 */{ 
                        var va_arg1 *Language.Do  
                        var va_arg2 *ClaireList  
                        va_arg1 = _CL_obj
                        var va_arg2_try016712 EID 
                        /* Let:12 */{ 
                          var i_bag *ClaireList   = ToType(C_any.Id()).EmptyList()
                          /* noccur = 2 */
                          /* Let:13 */{ 
                            var i int  = 1
                            /* noccur = 5 */
                            /* Let:14 */{ 
                              var g0155 int  = m
                              /* noccur = 1 */
                              va_arg2_try016712= EID{CFALSE.Id(),0}
                              for (i <= g0155) /* while:15 */{ 
                                var void_try16 EID 
                                _ = void_try16
                                { 
                                /* Let:16 */{ 
                                  var g0168UU *ClaireObject  
                                  /* noccur = 1 */
                                  var g0168UU_try016917 EID 
                                  /* Let:17 */{ 
                                    var g0170UU *Language.Call  
                                    /* noccur = 1 */
                                    var g0170UU_try017118 EID 
                                    g0170UU_try017118 = F_Call_I_property(C_nth,MakeConstantList(v.Id(),MakeInteger(i).Id()))
                                    /* ERROR PROTECTION INSERTED (g0170UU-g0168UU_try016917) */
                                    if ErrorIn(g0170UU_try017118) {g0168UU_try016917 = g0170UU_try017118
                                    } else {
                                    g0170UU = Language.To_Call(OBJ(g0170UU_try017118))
                                    g0168UU_try016917 = Language.C_Assign.Make(l.At(i-1),g0170UU.Id()).ToEID()
                                    }
                                    /* Let-17 */} 
                                  /* ERROR PROTECTION INSERTED (g0168UU-void_try16) */
                                  if ErrorIn(g0168UU_try016917) {void_try16 = g0168UU_try016917
                                  } else {
                                  g0168UU = ToObject(OBJ(g0168UU_try016917))
                                  void_try16 = EID{i_bag.AddFast(g0168UU.Id()).Id(),0}
                                  }
                                  /* Let-16 */} 
                                /* ERROR PROTECTION INSERTED (void_try16-void_try16) */
                                if ErrorIn(void_try16) {va_arg2_try016712 = void_try16
                                break
                                } else {
                                i = (i+1)
                                }
                                /* while-15 */} 
                              }
                              /* Let-14 */} 
                            /* Let-13 */} 
                          /* ERROR PROTECTION INSERTED (va_arg2_try016712-va_arg2_try016712) */
                          if !ErrorIn(va_arg2_try016712) {
                          va_arg2_try016712 = EID{i_bag.Id(),0}
                          }
                          /* Let-12 */} 
                        /* ERROR PROTECTION INSERTED (va_arg2-va_arg2_try016610) */
                        if ErrorIn(va_arg2_try016712) {va_arg2_try016610 = va_arg2_try016712
                        } else {
                        va_arg2 = ToList(OBJ(va_arg2_try016712))
                        /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                        va_arg1.Args = va_arg2
                        va_arg2_try016610 = EID{va_arg2.Id(),0}
                        }
                        /* update-11 */} 
                      /* ERROR PROTECTION INSERTED (va_arg2_try016610-va_arg2_try016610) */
                      if !ErrorIn(va_arg2_try016610) {
                      va_arg2_try016610 = EID{_CL_obj.Id(),0}
                      }
                      /* Let-10 */} 
                    /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                    if ErrorIn(va_arg2_try016610) {Result = va_arg2_try016610
                    } else {
                    va_arg2 = ANY(va_arg2_try016610)
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
                /* Let-7 */} 
              /* Let-6 */} 
            /* Let-5 */} 
          /* Let-4 */} 
        } else {
        Result = Language.C_Assign.Make(x,z).ToEID()
        /* If-3 */} 
      /* If!2 */}  else if (Equal(y,C_OR.Value) == CTRUE) /* If:2 */{ 
      if (x.Isa.IsIn(Language.C_Or) == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var g0157 *Language.Or   = Language.To_Or(x)
          /* noccur = 2 */
          g0157.Args.AddFast(z)
          Result = EID{g0157.Id(),0}
          /* Let-4 */} 
        } else {
        /* Let:4 */{ 
          var _CL_obj *Language.Or   = Language.To_Or(new(Language.Or).Is(Language.C_Or))
          /* noccur = 3 */
          _CL_obj.Args = MakeConstantList(x,z)
          Result = EID{_CL_obj.Id(),0}
          /* Let-4 */} 
        /* If-3 */} 
      /* If!2 */}  else if (Equal(y,C_AND.Value) == CTRUE) /* If:2 */{ 
      if (x.Isa.IsIn(Language.C_And) == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var g0159 *Language.And   = Language.To_And(x)
          /* noccur = 2 */
          g0159.Args.AddFast(z)
          Result = EID{g0159.Id(),0}
          /* Let-4 */} 
        } else {
        /* Let:4 */{ 
          var _CL_obj *Language.And   = Language.To_And(new(Language.And).Is(Language.C_And))
          /* noccur = 3 */
          _CL_obj.Args = MakeConstantList(x,z)
          Result = EID{_CL_obj.Id(),0}
          /* Let-4 */} 
        /* If-3 */} 
      /* If!2 */}  else if (y == C__Z.Id()) /* If:2 */{ 
      Result = F_Call_I_property(ToProperty(C__Z.Id()),MakeConstantList(x,z))
      } else {
      /* Let:3 */{ 
        var g0172UU *Language.Call_star  
        /* noccur = 1 */
        /* Let:4 */{ 
          var _CL_obj *Language.Call_star   = Language.To_Call_star(new(Language.Call_star).Is(Language.C_Call_star))
          /* noccur = 5 */
          _CL_obj.Selector = ToProperty(y)
          _CL_obj.Args = MakeConstantList(x,z)
          g0172UU = _CL_obj
          /* Let-4 */} 
        Result = F_DBregister_Call(Language.To_Call(g0172UU.Id()))
        /* Let-3 */} 
      /* If-2 */} 
    return Result} 
  
// The EID go function for: combine! @ any (throw: true) 
func E_combine_I_any (x EID,y EID,z EID) EID { 
    return /*(sm for combine! @ any= EID)*/ F_combine_I_any(ANY(x),ANY(y),ANY(z) )} 
  
// Call* says that combining is OK
// allows to treats Calls, Assigns, Gassign in an homogeneous way
// return false if the pattern is not (x OP y) and OP otherwise
/* {1} OPT.The go function for: operation!(x:any) [] */
func F_operation_I_any (x *ClaireAny ) *ClaireAny  { 
    // procedure body with s =  
var Result *ClaireAny  
    if (x.Isa.IsIn(Language.C_Or) == CTRUE) /* If:2 */{ 
      Result = C_OR.Value
      /* If!2 */}  else if (x.Isa.IsIn(Language.C_And) == CTRUE) /* If:2 */{ 
      Result = C_AND.Value
      /* If!2 */}  else if (x.Isa.IsIn(Language.C_Assign) == CTRUE) /* If:2 */{ 
      Result = C_L__equal.Id()
      /* If!2 */}  else if (x.Isa.IsIn(Language.C_Gassign) == CTRUE) /* If:2 */{ 
      Result = C_L__equal.Id()
      /* If!2 */}  else if (x.Isa.IsIn(Language.C_Call) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0177 *Language.Call   = Language.To_Call(x)
        /* noccur = 2 */
        /* Let:4 */{ 
          var p *ClaireProperty   = g0177.Selector
          /* noccur = 4 */
          if ((g0177.Isa.IsIn(Language.C_Call_star) == CTRUE) && 
              (F_operation_ask_any(p.Id()) == CTRUE)) /* If:5 */{ 
            Result = p.Id()
            /* If!5 */}  else if (p.Id() == C_nth_equal.Id()) /* If:5 */{ 
            Result = C_L__equal.Id()
            /* If!5 */}  else if (p.Id() == Core.C_write.Id()) /* If:5 */{ 
            Result = C_L__equal.Id()
            } else {
            Result = CFALSE.Id()
            /* If-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      } else {
      Result = CFALSE.Id()
      /* If-2 */} 
    return Result} 
  
// The EID go function for: operation! @ any (throw: false) 
func E_operation_I_any (x EID) EID { 
    return /*(sm for operation! @ any= any)*/ F_operation_I_any(ANY(x) ).ToEID()} 
  
// extract the two operands from an expression x such that operation!(x) != false
/* {1} OPT.The go function for: operand!(x:any,n:integer) [] */
func F_operand_I_any (x *ClaireAny ,n int) EID { 
    var Result EID 
    if (x.Isa.IsIn(Language.C_Or) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0179 *Language.Or   = Language.To_Or(x)
        /* noccur = 2 */
        if (n == 1) /* If:4 */{ 
          /* Let:5 */{ 
            var _CL_obj *Language.Or   = Language.To_Or(new(Language.Or).Is(Language.C_Or))
            /* noccur = 3 */
            /* update:6 */{ 
              var va_arg1 *Language.Or  
              var va_arg2 *ClaireList  
              va_arg1 = _CL_obj
              var va_arg2_try01857 EID 
              va_arg2_try01857 = Core.F_rmlast_list(g0179.Args.Copy())
              /* ERROR PROTECTION INSERTED (va_arg2-Result) */
              if ErrorIn(va_arg2_try01857) {Result = va_arg2_try01857
              } else {
              va_arg2 = ToList(OBJ(va_arg2_try01857))
              /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
              va_arg1.Args = va_arg2
              Result = EID{va_arg2.Id(),0}
              }
              /* update-6 */} 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            Result = EID{_CL_obj.Id(),0}
            }
            /* Let-5 */} 
          } else {
          Result = Core.F_last_list(g0179.Args)
          /* If-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (x.Isa.IsIn(Language.C_And) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0180 *Language.And   = Language.To_And(x)
        /* noccur = 2 */
        if (n == 1) /* If:4 */{ 
          /* Let:5 */{ 
            var _CL_obj *Language.And   = Language.To_And(new(Language.And).Is(Language.C_And))
            /* noccur = 3 */
            /* update:6 */{ 
              var va_arg1 *Language.And  
              var va_arg2 *ClaireList  
              va_arg1 = _CL_obj
              var va_arg2_try01867 EID 
              va_arg2_try01867 = Core.F_rmlast_list(g0180.Args.Copy())
              /* ERROR PROTECTION INSERTED (va_arg2-Result) */
              if ErrorIn(va_arg2_try01867) {Result = va_arg2_try01867
              } else {
              va_arg2 = ToList(OBJ(va_arg2_try01867))
              /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
              va_arg1.Args = va_arg2
              Result = EID{va_arg2.Id(),0}
              }
              /* update-6 */} 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            Result = EID{_CL_obj.Id(),0}
            }
            /* Let-5 */} 
          } else {
          Result = Core.F_last_list(g0180.Args)
          /* If-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (x.Isa.IsIn(Language.C_Assign) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0181 *Language.Assign   = Language.To_Assign(x)
        /* noccur = 2 */
        if (n == 1) /* If:4 */{ 
          Result = g0181.ClaireVar.ToEID()
          } else {
          Result = g0181.Arg.ToEID()
          /* If-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (x.Isa.IsIn(Language.C_Gassign) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0182 *Language.Gassign   = Language.To_Gassign(x)
        /* noccur = 2 */
        if (n == 1) /* If:4 */{ 
          Result = EID{g0182.ClaireVar.Id(),0}
          } else {
          Result = g0182.Arg.ToEID()
          /* If-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (x.Isa.IsIn(Language.C_Call) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0183 *Language.Call   = Language.To_Call(x)
        /* noccur = 8 */
        if (g0183.Selector.Id() == Core.C_write.Id()) /* If:4 */{ 
          if (n == 2) /* If:5 */{ 
            Result = g0183.Args.At(3-1).ToEID()
            } else {
            Result = F_Call_I_property(ToProperty(g0183.Args.At(1-1)),MakeConstantList(g0183.Args.At(2-1)))
            /* If-5 */} 
          /* If!4 */}  else if (g0183.Selector.Id() == C_nth_equal.Id()) /* If:4 */{ 
          if (n == 2) /* If:5 */{ 
            Result = Core.F_last_list(g0183.Args)
            } else {
            /* Let:6 */{ 
              var g0187UU *ClaireList  
              /* noccur = 1 */
              var g0187UU_try01887 EID 
              g0187UU_try01887 = Core.F_rmlast_list(g0183.Args.Copy())
              /* ERROR PROTECTION INSERTED (g0187UU-Result) */
              if ErrorIn(g0187UU_try01887) {Result = g0187UU_try01887
              } else {
              g0187UU = ToList(OBJ(g0187UU_try01887))
              Result = F_Call_I_property(C_nth,g0187UU)
              }
              /* Let-6 */} 
            /* If-5 */} 
          } else {
          Result = g0183.Args.At(n-1).ToEID()
          /* If-4 */} 
        /* Let-3 */} 
      } else {
      Result = EID{CFALSE.Id(),0}
      /* If-2 */} 
    return Result} 
  
// The EID go function for: operand! @ any (throw: true) 
func E_operand_I_any (x EID,n EID) EID { 
    return /*(sm for operand! @ any= EID)*/ F_operand_I_any(ANY(x),INT(n) )} 
  
// precedence
//
/* {1} OPT.The go function for: precedence!(y:any) [] */
func F_precedence_I_any (y *ClaireAny ) int { 
    // use function body compiling 
if (y == C_as.Id()) /* body If:2 */{ 
      return  0
      }  else if (y == C_L__equal.Id()) /* body If:2 */{ 
      return  100
      }  else if (Equal(y,C_AND.Value) == CTRUE) /* body If:2 */{ 
      return  1000
      }  else if (Equal(y,C_OR.Value) == CTRUE) /* body If:2 */{ 
      return  1010
      } else {
      return  ToOperation(y).Precedence
      /* body If-2 */} 
    } 
  
// The EID go function for: precedence! @ any (throw: false) 
func E_precedence_I_any (y EID) EID { 
    return EID{C__INT,IVAL(/*(sm for precedence! @ any= integer)*/ F_precedence_I_any(ANY(y) ))}} 
  
// **********************************************************************
// *   Part 2: read control structures                                  *
// **********************************************************************
/* {1} OPT.The go function for: nextstruct(r:meta_reader,%first:keyword,e:keyword) [] */
func (r *MetaReader ) Nextstruct (_Zfirst *ClaireKeyword ,e *ClaireKeyword ) EID { 
    var Result EID 
    if (_Zfirst.Id() == C_let.Id()) /* If:2 */{ 
      Result = r.Readlet(e)
      /* If!2 */}  else if (_Zfirst.Id() == C_when.Id()) /* If:2 */{ 
      Result = r.Readwhen(e)
      /* If!2 */}  else if (_Zfirst.Id() == C_case.Id()) /* If:2 */{ 
      Result = r.Readcase(e)
      /* If!2 */}  else if (_Zfirst.Id() == C_for.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var _Zvar *ClaireVariable  
        /* noccur = 2 */
        var _Zvar_try01894 EID 
        /* Let:4 */{ 
          var g0190UU *ClaireAny  
          /* noccur = 1 */
          var g0190UU_try01915 EID 
          g0190UU_try01915 = F_nexts_I_meta_reader1(r,C_in)
          /* ERROR PROTECTION INSERTED (g0190UU-_Zvar_try01894) */
          if ErrorIn(g0190UU_try01915) {_Zvar_try01894 = g0190UU_try01915
          } else {
          g0190UU = ANY(g0190UU_try01915)
          _Zvar_try01894 = F_extract_variable_any(g0190UU)
          }
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (_Zvar-Result) */
        if ErrorIn(_Zvar_try01894) {Result = _Zvar_try01894
        } else {
        _Zvar = To_Variable(OBJ(_Zvar_try01894))
        /* Let:4 */{ 
          var _Zset *ClaireAny  
          /* noccur = 1 */
          var _Zset_try01925 EID 
          _Zset_try01925 = r.Nexte()
          /* ERROR PROTECTION INSERTED (_Zset-Result) */
          if ErrorIn(_Zset_try01925) {Result = _Zset_try01925
          } else {
          _Zset = ANY(_Zset_try01925)
          /* Let:5 */{ 
            var _Zbind *ClaireList   = r.Bind_I(_Zvar)
            /* noccur = 1 */
            /* LetE:6 */{ 
              var x EID 
              if (r.Firstc() == 44) /* If:7 */{ 
                r.Next()
                /* If-7 */} 
              /* Let:7 */{ 
                var _CL_obj *Language.For   = Language.To_For(new(Language.For).Is(Language.C_For))
                /* noccur = 5 */
                _CL_obj.ClaireVar = _Zvar
                _CL_obj.SetArg = _Zset
                /* update:8 */{ 
                  var va_arg1 *Language.Iteration  
                  var va_arg2 *ClaireAny  
                  va_arg1 = Language.To_Iteration(_CL_obj.Id())
                  var va_arg2_try01939 EID 
                  va_arg2_try01939 = r.Nexts(e)
                  /* ERROR PROTECTION INSERTED (va_arg2-x) */
                  if ErrorIn(va_arg2_try01939) {x = va_arg2_try01939
                  } else {
                  va_arg2 = ANY(va_arg2_try01939)
                  /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                  va_arg1.Arg = va_arg2
                  x = va_arg2.ToEID()
                  }
                  /* update-8 */} 
                /* ERROR PROTECTION INSERTED (x-x) */
                if !ErrorIn(x) {
                x = EID{_CL_obj.Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (x-Result) */
              if ErrorIn(x) {Result = x
              } else {
              r.Unbind_I(_Zbind)
              Result = x}
              /* LetE-6 */} 
            /* Let-5 */} 
          }
          /* Let-4 */} 
        }
        /* Let-3 */} 
      /* If!2 */}  else if (_Zfirst.Id() == C_while.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var _CL_obj *Language.While   = Language.To_While(new(Language.While).Is(Language.C_While))
        /* noccur = 7 */
        /* update:4 */{ 
          var va_arg1 *Language.While  
          var va_arg2 *ClaireAny  
          va_arg1 = _CL_obj
          var va_arg2_try01945 EID 
          va_arg2_try01945 = r.Nexte()
          /* ERROR PROTECTION INSERTED (va_arg2-Result) */
          if ErrorIn(va_arg2_try01945) {Result = va_arg2_try01945
          } else {
          va_arg2 = ANY(va_arg2_try01945)
          /* ---------- now we compile update iClaire/test(va_arg1) := va_arg2 ------- */
          va_arg1.Test = va_arg2
          Result = va_arg2.ToEID()
          }
          /* update-4 */} 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        /* update:4 */{ 
          var va_arg1 *Language.While  
          var va_arg2 *ClaireAny  
          va_arg1 = _CL_obj
          var va_arg2_try01955 EID 
          va_arg2_try01955 = r.Nexts(e)
          /* ERROR PROTECTION INSERTED (va_arg2-Result) */
          if ErrorIn(va_arg2_try01955) {Result = va_arg2_try01955
          } else {
          va_arg2 = ANY(va_arg2_try01955)
          /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
          va_arg1.Arg = va_arg2
          Result = va_arg2.ToEID()
          }
          /* update-4 */} 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        _CL_obj.Other = CFALSE
        Result = EID{_CL_obj.Id(),0}
        }}
        /* Let-3 */} 
      /* If!2 */}  else if (_Zfirst.Id() == C_until.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var _CL_obj *Language.While   = Language.To_While(new(Language.While).Is(Language.C_While))
        /* noccur = 7 */
        /* update:4 */{ 
          var va_arg1 *Language.While  
          var va_arg2 *ClaireAny  
          va_arg1 = _CL_obj
          var va_arg2_try01965 EID 
          va_arg2_try01965 = r.Nexte()
          /* ERROR PROTECTION INSERTED (va_arg2-Result) */
          if ErrorIn(va_arg2_try01965) {Result = va_arg2_try01965
          } else {
          va_arg2 = ANY(va_arg2_try01965)
          /* ---------- now we compile update iClaire/test(va_arg1) := va_arg2 ------- */
          va_arg1.Test = va_arg2
          Result = va_arg2.ToEID()
          }
          /* update-4 */} 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        /* update:4 */{ 
          var va_arg1 *Language.While  
          var va_arg2 *ClaireAny  
          va_arg1 = _CL_obj
          var va_arg2_try01975 EID 
          va_arg2_try01975 = r.Nexts(e)
          /* ERROR PROTECTION INSERTED (va_arg2-Result) */
          if ErrorIn(va_arg2_try01975) {Result = va_arg2_try01975
          } else {
          va_arg2 = ANY(va_arg2_try01975)
          /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
          va_arg1.Arg = va_arg2
          Result = va_arg2.ToEID()
          }
          /* update-4 */} 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        _CL_obj.Other = CTRUE
        Result = EID{_CL_obj.Id(),0}
        }}
        /* Let-3 */} 
      /* If!2 */}  else if (_Zfirst.Id() == C_try.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var _Za *ClaireAny  
        /* noccur = 1 */
        var _Za_try01984 EID 
        _Za_try01984 = F_nexts_I_meta_reader1(r,C_catch)
        /* ERROR PROTECTION INSERTED (_Za-Result) */
        if ErrorIn(_Za_try01984) {Result = _Za_try01984
        } else {
        _Za = ANY(_Za_try01984)
        /* Let:4 */{ 
          var _Zt *ClaireAny  
          /* noccur = 3 */
          var _Zt_try01995 EID 
          _Zt_try01995 = r.Nexte()
          /* ERROR PROTECTION INSERTED (_Zt-Result) */
          if ErrorIn(_Zt_try01995) {Result = _Zt_try01995
          } else {
          _Zt = ANY(_Zt_try01995)
          if (C_class.Id() == _Zt.Isa.Id()) /* If:5 */{ 
            /* Let:6 */{ 
              var g0200UU *ClaireAny  
              /* noccur = 1 */
              var g0200UU_try02017 EID 
              g0200UU_try02017 = r.Nexts(e)
              /* ERROR PROTECTION INSERTED (g0200UU-Result) */
              if ErrorIn(g0200UU_try02017) {Result = g0200UU_try02017
              } else {
              g0200UU = ANY(g0200UU_try02017)
              Result = Language.C_Handle.Make(_Zt,_Za,g0200UU).ToEID()
              }
              /* Let-6 */} 
            } else {
            Result = F_Serror_string(MakeString("[00] in try/catch, ~S is not a class"),MakeConstantList(_Zt))
            /* If-5 */} 
          }
          /* Let-4 */} 
        }
        /* Let-3 */} 
      } else {
      Result = EID{_Zfirst.Id(),0}
      /* If-2 */} 
    return Result} 
  
// The EID go function for: nextstruct @ meta_reader (throw: true) 
func E_nextstruct_meta_reader (r EID,_Zfirst EID,e EID) EID { 
    return /*(sm for nextstruct @ meta_reader= EID)*/ ToMetaReader(OBJ(r)).Nextstruct(ToKeyword(OBJ(_Zfirst)),ToKeyword(OBJ(e)) )} 
  
// reads a let expression
//
/* {1} OPT.The go function for: readlet(r:meta_reader,e:keyword) [] */
func (r *MetaReader ) Readlet (e *ClaireKeyword ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var _Zdef *ClaireAny  
      /* noccur = 8 */
      var _Zdef_try02063 EID 
      _Zdef_try02063 = F_nexts_I_meta_reader3(r,C_in,44)
      /* ERROR PROTECTION INSERTED (_Zdef-Result) */
      if ErrorIn(_Zdef_try02063) {Result = _Zdef_try02063
      } else {
      _Zdef = ANY(_Zdef_try02063)
      if (_Zdef.Isa.IsIn(Language.C_Assign) == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var g0202 *Language.Assign   = Language.To_Assign(_Zdef)
          /* noccur = 2 */
          /* Let:5 */{ 
            var v *ClaireVariable  
            /* noccur = 2 */
            var v_try02076 EID 
            v_try02076 = F_extract_variable_any(g0202.ClaireVar)
            /* ERROR PROTECTION INSERTED (v-Result) */
            if ErrorIn(v_try02076) {Result = v_try02076
            } else {
            v = To_Variable(OBJ(v_try02076))
            /* Let:6 */{ 
              var _Zbind *ClaireList   = r.Bind_I(v)
              /* noccur = 1 */
              /* LetE:7 */{ 
                var x EID 
                /* Let:8 */{ 
                  var _CL_obj *Language.Let   = Language.To_Let(new(Language.Let).Is(Language.C_Let))
                  /* noccur = 6 */
                  _CL_obj.ClaireVar = v
                  _CL_obj.Value = g0202.Arg
                  /* update:9 */{ 
                    var va_arg1 *Language.Let  
                    var va_arg2 *ClaireAny  
                    va_arg1 = _CL_obj
                    var va_arg2_try020810 EID 
                    if (r.Firstc() == 44) /* If:10 */{ 
                      va_arg2_try020810 = r.Cnext().Readlet(e)
                      } else {
                      va_arg2_try020810 = r.Nexts(e)
                      /* If-10 */} 
                    /* ERROR PROTECTION INSERTED (va_arg2-x) */
                    if ErrorIn(va_arg2_try020810) {x = va_arg2_try020810
                    } else {
                    va_arg2 = ANY(va_arg2_try020810)
                    /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                    va_arg1.Arg = va_arg2
                    x = va_arg2.ToEID()
                    }
                    /* update-9 */} 
                  /* ERROR PROTECTION INSERTED (x-x) */
                  if !ErrorIn(x) {
                  x = EID{_CL_obj.Id(),0}
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (x-Result) */
                if ErrorIn(x) {Result = x
                } else {
                r.Unbind_I(_Zbind)
                Result = x}
                /* LetE-7 */} 
              /* Let-6 */} 
            }
            /* Let-5 */} 
          /* Let-4 */} 
        /* If!3 */}  else if (_Zdef.Isa.IsIn(Language.C_Let_star) == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var g0203 *Language.Let_star   = Language.To_Let_star(_Zdef)
          /* noccur = 4 */
          /* update:5 */{ 
            var va_arg1 *Language.Let  
            var va_arg2 *ClaireAny  
            va_arg1 = Language.To_Let(g0203.Id())
            var va_arg2_try02096 EID 
            va_arg2_try02096 = r.Readlet_star(Language.To_Do(g0203.Arg).Args,1,e)
            /* ERROR PROTECTION INSERTED (va_arg2-Result) */
            if ErrorIn(va_arg2_try02096) {Result = va_arg2_try02096
            } else {
            va_arg2 = ANY(va_arg2_try02096)
            /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
            va_arg1.Arg = va_arg2
            Result = va_arg2.ToEID()
            }
            /* update-5 */} 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          Result = EID{g0203.Id(),0}
          }
          /* Let-4 */} 
        /* If!3 */}  else if ((_Zdef.Isa.IsIn(Language.C_Call) == CTRUE) && 
          (ANY(Core.F_CALL(C_selector,ARGS(_Zdef.ToEID()))) == Core.C_write.Id())) /* If:3 */{ 
        /* Let:4 */{ 
          var g0204 *Language.Call   = Language.To_Call(_Zdef)
          /* noccur = 2 */
          /* Let:5 */{ 
            var v1 *ClaireVariable  
            /* noccur = 2 */
            /* Let:6 */{ 
              var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
              /* noccur = 5 */
              _CL_obj.Pname = Core.F_gensym_void()
              _CL_obj.Range = ToType(C_any.Id())
              v1 = _CL_obj
              /* Let-6 */} 
            /* Let:6 */{ 
              var v2 *ClaireVariable  
              /* noccur = 2 */
              /* Let:7 */{ 
                var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
                /* noccur = 5 */
                _CL_obj.Pname = Core.F_gensym_void()
                _CL_obj.Range = ToType(C_any.Id())
                v2 = _CL_obj
                /* Let-7 */} 
              /* Let:7 */{ 
                var _Za *ClaireList   = g0204.Args
                /* noccur = 4 */
                /* Let:8 */{ 
                  var _Ze *ClaireAny  
                  /* noccur = 1 */
                  var _Ze_try02109 EID 
                  _Ze_try02109 = r.Nexts(e)
                  /* ERROR PROTECTION INSERTED (_Ze-Result) */
                  if ErrorIn(_Ze_try02109) {Result = _Ze_try02109
                  } else {
                  _Ze = ANY(_Ze_try02109)
                  /* Let:9 */{ 
                    var _CL_obj *Language.Let_plus   = Language.To_Let_plus(new(Language.Let_plus).Is(Language.C_Let_plus))
                    /* noccur = 27 */
                    _CL_obj.ClaireVar = v1
                    /* update:10 */{ 
                      var va_arg1 *Language.Let  
                      var va_arg2 *ClaireAny  
                      va_arg1 = Language.To_Let(_CL_obj.Id())
                      /* Let:11 */{ 
                        var _CL_obj *Language.Call_plus   = Language.To_Call_plus(new(Language.Call_plus).Is(Language.C_Call_plus))
                        /* noccur = 5 */
                        _CL_obj.Selector = ToProperty(_Za.At(1-1))
                        _CL_obj.Args = MakeConstantList(_Za.At(2-1))
                        va_arg2 = _CL_obj.Id()
                        /* Let-11 */} 
                      /* ---------- now we compile update value(va_arg1) := va_arg2 ------- */
                      va_arg1.Value = va_arg2
                      /* update-10 */} 
                    /* update:10 */{ 
                      var va_arg1 *Language.Let  
                      var va_arg2 *ClaireAny  
                      va_arg1 = Language.To_Let(_CL_obj.Id())
                      var va_arg2_try021111 EID 
                      /* Let:11 */{ 
                        var _CL_obj *Language.Do   = Language.To_Do(new(Language.Do).Is(Language.C_Do))
                        /* noccur = 14 */
                        /* update:12 */{ 
                          var va_arg1 *Language.Do  
                          var va_arg2 *ClaireList  
                          va_arg1 = _CL_obj
                          var va_arg2_try021213 EID 
                          /* Construct:13 */{ 
                            var v_bag_arg *ClaireAny  
                            va_arg2_try021213= EID{ToType(C_any.Id()).EmptyList().Id(),0}
                            ToList(OBJ(va_arg2_try021213)).AddFast(g0204.Id())
                            var v_bag_arg_try021314 EID 
                            /* Let:14 */{ 
                              var _CL_obj *Language.Let   = Language.To_Let(new(Language.Let).Is(Language.C_Let))
                              /* noccur = 10 */
                              _CL_obj.ClaireVar = v2
                              _CL_obj.Value = _Ze
                              /* update:15 */{ 
                                var va_arg1 *Language.Let  
                                var va_arg2 *ClaireAny  
                                va_arg1 = _CL_obj
                                var va_arg2_try021416 EID 
                                /* Let:16 */{ 
                                  var _CL_obj *Language.Do   = Language.To_Do(new(Language.Do).Is(Language.C_Do))
                                  /* noccur = 3 */
                                  /* update:17 */{ 
                                    var va_arg1 *Language.Do  
                                    var va_arg2 *ClaireList  
                                    va_arg1 = _CL_obj
                                    var va_arg2_try021518 EID 
                                    /* Construct:18 */{ 
                                      var v_bag_arg *ClaireAny  
                                      va_arg2_try021518= EID{ToType(C_any.Id()).EmptyList().Id(),0}
                                      var v_bag_arg_try021619 EID 
                                      v_bag_arg_try021619 = F_Call_I_property(Core.C_write,MakeConstantList(_Za.At(1-1),_Za.At(2-1),v1.Id()))
                                      /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try021518) */
                                      if ErrorIn(v_bag_arg_try021619) {va_arg2_try021518 = v_bag_arg_try021619
                                      } else {
                                      v_bag_arg = ANY(v_bag_arg_try021619)
                                      ToList(OBJ(va_arg2_try021518)).AddFast(v_bag_arg)
                                      ToList(OBJ(va_arg2_try021518)).AddFast(v2.Id())}
                                      /* Construct-18 */} 
                                    /* ERROR PROTECTION INSERTED (va_arg2-va_arg2_try021416) */
                                    if ErrorIn(va_arg2_try021518) {va_arg2_try021416 = va_arg2_try021518
                                    } else {
                                    va_arg2 = ToList(OBJ(va_arg2_try021518))
                                    /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                                    va_arg1.Args = va_arg2
                                    va_arg2_try021416 = EID{va_arg2.Id(),0}
                                    }
                                    /* update-17 */} 
                                  /* ERROR PROTECTION INSERTED (va_arg2_try021416-va_arg2_try021416) */
                                  if !ErrorIn(va_arg2_try021416) {
                                  va_arg2_try021416 = EID{_CL_obj.Id(),0}
                                  }
                                  /* Let-16 */} 
                                /* ERROR PROTECTION INSERTED (va_arg2-v_bag_arg_try021314) */
                                if ErrorIn(va_arg2_try021416) {v_bag_arg_try021314 = va_arg2_try021416
                                } else {
                                va_arg2 = ANY(va_arg2_try021416)
                                /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                                va_arg1.Arg = va_arg2
                                v_bag_arg_try021314 = va_arg2.ToEID()
                                }
                                /* update-15 */} 
                              /* ERROR PROTECTION INSERTED (v_bag_arg_try021314-v_bag_arg_try021314) */
                              if !ErrorIn(v_bag_arg_try021314) {
                              v_bag_arg_try021314 = EID{_CL_obj.Id(),0}
                              }
                              /* Let-14 */} 
                            /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try021213) */
                            if ErrorIn(v_bag_arg_try021314) {va_arg2_try021213 = v_bag_arg_try021314
                            } else {
                            v_bag_arg = ANY(v_bag_arg_try021314)
                            ToList(OBJ(va_arg2_try021213)).AddFast(v_bag_arg)}
                            /* Construct-13 */} 
                          /* ERROR PROTECTION INSERTED (va_arg2-va_arg2_try021111) */
                          if ErrorIn(va_arg2_try021213) {va_arg2_try021111 = va_arg2_try021213
                          } else {
                          va_arg2 = ToList(OBJ(va_arg2_try021213))
                          /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                          va_arg1.Args = va_arg2
                          va_arg2_try021111 = EID{va_arg2.Id(),0}
                          }
                          /* update-12 */} 
                        /* ERROR PROTECTION INSERTED (va_arg2_try021111-va_arg2_try021111) */
                        if !ErrorIn(va_arg2_try021111) {
                        va_arg2_try021111 = EID{_CL_obj.Id(),0}
                        }
                        /* Let-11 */} 
                      /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                      if ErrorIn(va_arg2_try021111) {Result = va_arg2_try021111
                      } else {
                      va_arg2 = ANY(va_arg2_try021111)
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
                  }
                  /* Let-8 */} 
                /* Let-7 */} 
              /* Let-6 */} 
            /* Let-5 */} 
          /* Let-4 */} 
        } else {
        Result = F_Serror_string(MakeString("[165] ~S is illegal after a let"),MakeConstantList(_Zdef))
        /* If-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: readlet @ meta_reader (throw: true) 
func E_readlet_meta_reader (r EID,e EID) EID { 
    return /*(sm for readlet @ meta_reader= EID)*/ ToMetaReader(OBJ(r)).Readlet(ToKeyword(OBJ(e)) )} 
  
// recursive construction of the tail of a Let*
//
/* {1} OPT.The go function for: readlet*(r:meta_reader,l:list,n:integer,e:keyword) [] */
func (r *MetaReader ) Readlet_star (l *ClaireList ,n int,e *ClaireKeyword ) EID { 
    var Result EID 
    if (n > l.Length()) /* If:2 */{ 
      Result = r.Nexts(e)
      } else {
      /* Let:3 */{ 
        var v *ClaireVariable  
        /* noccur = 2 */
        var v_try02174 EID 
        v_try02174 = F_extract_variable_any(ANY(Core.F_CALL(Language.C_var,ARGS(l.At(n-1).ToEID()))))
        /* ERROR PROTECTION INSERTED (v-Result) */
        if ErrorIn(v_try02174) {Result = v_try02174
        } else {
        v = To_Variable(OBJ(v_try02174))
        /* Let:4 */{ 
          var _Zbind *ClaireList   = r.Bind_I(v)
          /* noccur = 1 */
          /* LetE:5 */{ 
            var x EID 
            /* Let:6 */{ 
              var _CL_obj *Language.Let   = Language.To_Let(new(Language.Let).Is(Language.C_Let))
              /* noccur = 6 */
              _CL_obj.ClaireVar = v
              _CL_obj.Value = ANY(Core.F_CALL(C_arg,ARGS(l.At(n-1).ToEID())))
              /* update:7 */{ 
                var va_arg1 *Language.Let  
                var va_arg2 *ClaireAny  
                va_arg1 = _CL_obj
                var va_arg2_try02188 EID 
                va_arg2_try02188 = r.Readlet_star(l,(n+1),e)
                /* ERROR PROTECTION INSERTED (va_arg2-x) */
                if ErrorIn(va_arg2_try02188) {x = va_arg2_try02188
                } else {
                va_arg2 = ANY(va_arg2_try02188)
                /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                va_arg1.Arg = va_arg2
                x = va_arg2.ToEID()
                }
                /* update-7 */} 
              /* ERROR PROTECTION INSERTED (x-x) */
              if !ErrorIn(x) {
              x = EID{_CL_obj.Id(),0}
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (x-Result) */
            if ErrorIn(x) {Result = x
            } else {
            r.Unbind_I(_Zbind)
            Result = x}
            /* LetE-5 */} 
          /* Let-4 */} 
        }
        /* Let-3 */} 
      /* If-2 */} 
    return Result} 
  
// The EID go function for: readlet* @ meta_reader (throw: true) 
func E_readlet_star_meta_reader (r EID,l EID,n EID,e EID) EID { 
    return /*(sm for readlet* @ meta_reader= EID)*/ ToMetaReader(OBJ(r)).Readlet_star(ToList(OBJ(l)),
      INT(n),
      ToKeyword(OBJ(e)) )} 
  
// reads a when expression
//
/* {1} OPT.The go function for: readwhen(r:meta_reader,e:keyword) [] */
func (r *MetaReader ) Readwhen (e *ClaireKeyword ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var _Zdef *ClaireAny  
      /* noccur = 3 */
      var _Zdef_try02213 EID 
      _Zdef_try02213 = F_nexts_I_meta_reader3(r,C_in,44)
      /* ERROR PROTECTION INSERTED (_Zdef-Result) */
      if ErrorIn(_Zdef_try02213) {Result = _Zdef_try02213
      } else {
      _Zdef = ANY(_Zdef_try02213)
      if (_Zdef.Isa.IsIn(Language.C_Assign) == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var g0219 *Language.Assign   = Language.To_Assign(_Zdef)
          /* noccur = 2 */
          /* Let:5 */{ 
            var v *ClaireVariable  
            /* noccur = 2 */
            var v_try02226 EID 
            v_try02226 = F_extract_variable_any(g0219.ClaireVar)
            /* ERROR PROTECTION INSERTED (v-Result) */
            if ErrorIn(v_try02226) {Result = v_try02226
            } else {
            v = To_Variable(OBJ(v_try02226))
            /* Let:6 */{ 
              var _Zbind *ClaireList   = r.Bind_I(v)
              /* noccur = 1 */
              /* Let:7 */{ 
                var _Za *ClaireAny  
                /* noccur = 1 */
                var _Za_try02238 EID 
                _Za_try02238 = r.Nexts(ToKeyword(C_else.Id()))
                /* ERROR PROTECTION INSERTED (_Za-Result) */
                if ErrorIn(_Za_try02238) {Result = _Za_try02238
                } else {
                _Za = ANY(_Za_try02238)
                /* LetE:8 */{ 
                  var x EID 
                  /* Let:9 */{ 
                    var _CL_obj *Language.When   = Language.To_When(new(Language.When).Is(Language.C_When))
                    /* noccur = 8 */
                    _CL_obj.ClaireVar = v
                    _CL_obj.Value = g0219.Arg
                    _CL_obj.Arg = _Za
                    /* update:10 */{ 
                      var va_arg1 *Language.When  
                      var va_arg2 *ClaireAny  
                      va_arg1 = _CL_obj
                      var va_arg2_try022411 EID 
                      if (F_boolean_I_any(F_stop_ask_integer(r.Firstc())) == CTRUE) /* If:11 */{ 
                        va_arg2_try022411 = EID{CNULL,0}
                        } else {
                        va_arg2_try022411 = r.Nexts(e)
                        /* If-11 */} 
                      /* ERROR PROTECTION INSERTED (va_arg2-x) */
                      if ErrorIn(va_arg2_try022411) {x = va_arg2_try022411
                      } else {
                      va_arg2 = ANY(va_arg2_try022411)
                      /* ---------- now we compile update iClaire/other(va_arg1) := va_arg2 ------- */
                      va_arg1.Other = va_arg2
                      x = va_arg2.ToEID()
                      }
                      /* update-10 */} 
                    /* ERROR PROTECTION INSERTED (x-x) */
                    if !ErrorIn(x) {
                    x = EID{_CL_obj.Id(),0}
                    }
                    /* Let-9 */} 
                  /* ERROR PROTECTION INSERTED (x-Result) */
                  if ErrorIn(x) {Result = x
                  } else {
                  r.Unbind_I(_Zbind)
                  Result = x}
                  /* LetE-8 */} 
                }
                /* Let-7 */} 
              /* Let-6 */} 
            }
            /* Let-5 */} 
          /* Let-4 */} 
        } else {
        Result = F_Serror_string(MakeString("[165] ~S is illegal after a when"),MakeConstantList(_Zdef))
        /* If-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: readwhen @ meta_reader (throw: true) 
func E_readwhen_meta_reader (r EID,e EID) EID { 
    return /*(sm for readwhen @ meta_reader= EID)*/ ToMetaReader(OBJ(r)).Readwhen(ToKeyword(OBJ(e)) )} 
  
// read an if
//
/* {1} OPT.The go function for: readif(r:meta_reader,e:integer) [] */
func (r *MetaReader ) Readif (e int) EID { 
    var Result EID 
    /* Let:2 */{ 
      var _Za1 *ClaireAny  
      /* noccur = 1 */
      var _Za1_try02253 EID 
      _Za1_try02253 = r.Nexte()
      /* ERROR PROTECTION INSERTED (_Za1-Result) */
      if ErrorIn(_Za1_try02253) {Result = _Za1_try02253
      } else {
      _Za1 = ANY(_Za1_try02253)
      /* Let:3 */{ 
        var _Za2 *ClaireAny  
        /* noccur = 1 */
        var _Za2_try02264 EID 
        _Za2_try02264 = r.Nexts(ToKeyword(C_else.Id()))
        /* ERROR PROTECTION INSERTED (_Za2-Result) */
        if ErrorIn(_Za2_try02264) {Result = _Za2_try02264
        } else {
        _Za2 = ANY(_Za2_try02264)
        /* Let:4 */{ 
          var g0227UU *ClaireAny  
          /* noccur = 1 */
          var g0227UU_try02285 EID 
          if ((r.Firstc() == 44) || 
              (r.Firstc() == e)) /* If:5 */{ 
            g0227UU_try02285 = EID{CFALSE.Id(),0}
            } else {
            /* Let:6 */{ 
              var x *ClaireAny  
              /* noccur = 4 */
              var x_try02297 EID 
              x_try02297 = r.Nexte()
              /* ERROR PROTECTION INSERTED (x-g0227UU_try02285) */
              if ErrorIn(x_try02297) {g0227UU_try02285 = x_try02297
              } else {
              x = ANY(x_try02297)
              if (x == C_if.Id()) /* If:7 */{ 
                g0227UU_try02285 = r.Readif(e)
                /* If!7 */}  else if (F_keyword_ask_any(x) == CTRUE) /* If:7 */{ 
                g0227UU_try02285 = r.Nextstruct(ToKeyword(x),C_none)
                } else {
                g0227UU_try02285 = r.Loopexp(x,C_none,CFALSE)
                /* If-7 */} 
              }
              /* Let-6 */} 
            /* If-5 */} 
          /* ERROR PROTECTION INSERTED (g0227UU-Result) */
          if ErrorIn(g0227UU_try02285) {Result = g0227UU_try02285
          } else {
          g0227UU = ANY(g0227UU_try02285)
          Result = Language.C_If.Make(_Za1,_Za2,g0227UU).ToEID()
          }
          /* Let-4 */} 
        }
        /* Let-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: readif @ meta_reader (throw: true) 
func E_readif_meta_reader (r EID,e EID) EID { 
    return /*(sm for readif @ meta_reader= EID)*/ ToMetaReader(OBJ(r)).Readif(INT(e) )} 
  
// reads a member_of
//
/* {1} OPT.The go function for: readcase(r:meta_reader,e:keyword) [] */
func (r *MetaReader ) Readcase (e *ClaireKeyword ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var _Zv *ClaireAny  
      /* noccur = 2 */
      var _Zv_try02303 EID 
      _Zv_try02303 = r.Nexte()
      /* ERROR PROTECTION INSERTED (_Zv-Result) */
      if ErrorIn(_Zv_try02303) {Result = _Zv_try02303
      } else {
      _Zv = ANY(_Zv_try02303)
      var g0231I *ClaireBoolean  
      var g0231I_try02323 EID 
      /* Let:3 */{ 
        var g0233UU int 
        /* noccur = 1 */
        var g0233UU_try02344 EID 
        g0233UU_try02344 = r.Skipc_I()
        /* ERROR PROTECTION INSERTED (g0233UU-g0231I_try02323) */
        if ErrorIn(g0233UU_try02344) {g0231I_try02323 = g0233UU_try02344
        } else {
        g0233UU = INT(g0233UU_try02344)
        g0231I_try02323 = EID{Core.F__I_equal_any(MakeInteger(g0233UU).Id(),MakeInteger(40).Id()).Id(),0}
        }
        /* Let-3 */} 
      /* ERROR PROTECTION INSERTED (g0231I-Result) */
      if ErrorIn(g0231I_try02323) {Result = g0231I_try02323
      } else {
      g0231I = ToBoolean(OBJ(g0231I_try02323))
      if (g0231I == CTRUE) /* If:3 */{ 
        Result = F_Serror_string(MakeString("[166] Missing ( after case ~S"),MakeConstantList(_Zv))
        } else {
        Result = EID{CFALSE.Id(),0}
        /* If-3 */} 
      }
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      /* Let:3 */{ 
        var _Zx *Language.Case  
        /* noccur = 5 */
        /* Let:4 */{ 
          var _CL_obj *Language.Case   = Language.To_Case(new(Language.Case).Is(Language.C_Case))
          /* noccur = 4 */
          _CL_obj.ClaireVar = _Zv
          _CL_obj.Args = ToType(CEMPTY.Id()).EmptyList()
          _Zx = _CL_obj
          /* Let-4 */} 
        /* Let:4 */{ 
          var _Zt *ClaireAny   = C_any.Id()
          /* noccur = 2 */
          Result= EID{CFALSE.Id(),0}
          for (r.Firstc() != 41) /* while:5 */{ 
            var void_try6 EID 
            _ = void_try6
            { 
            r.Next()
            var _Zt_try02356 EID 
            /* Let:6 */{ 
              var g0236UU *ClaireAny  
              /* noccur = 1 */
              var g0236UU_try02377 EID 
              g0236UU_try02377 = r.Nexte()
              /* ERROR PROTECTION INSERTED (g0236UU-_Zt_try02356) */
              if ErrorIn(g0236UU_try02377) {_Zt_try02356 = g0236UU_try02377
              } else {
              g0236UU = ANY(g0236UU_try02377)
              _Zt_try02356 = Language.F_extract_type_any(g0236UU)
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (_Zt-void_try6) */
            if ErrorIn(_Zt_try02356) {void_try6 = _Zt_try02356
            Result = _Zt_try02356
            break
            } else {
            _Zt = ANY(_Zt_try02356)
            void_try6 = _Zt.ToEID()
            /* update:6 */{ 
              var va_arg1 *Language.Case  
              var va_arg2 *ClaireList  
              va_arg1 = _Zx
              var va_arg2_try02387 EID 
              /* Let:7 */{ 
                var g0239UU *ClaireAny  
                /* noccur = 1 */
                var g0239UU_try02408 EID 
                g0239UU_try02408 = r.Nexts(C_none)
                /* ERROR PROTECTION INSERTED (g0239UU-va_arg2_try02387) */
                if ErrorIn(g0239UU_try02408) {va_arg2_try02387 = g0239UU_try02408
                } else {
                g0239UU = ANY(g0239UU_try02408)
                va_arg2_try02387 = EID{_Zx.Args.AddFast(_Zt).AddFast(g0239UU).Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (va_arg2-void_try6) */
              if ErrorIn(va_arg2_try02387) {void_try6 = va_arg2_try02387
              } else {
              va_arg2 = ToList(OBJ(va_arg2_try02387))
              /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
              va_arg1.Args = va_arg2
              void_try6 = EID{va_arg2.Id(),0}
              }
              /* update-6 */} 
            /* ERROR PROTECTION INSERTED (void_try6-void_try6) */
            if ErrorIn(void_try6) {Result = void_try6
            break
            } else {
            if ((F_boolean_I_any(F_stop_ask_integer(r.Firstc())).Id() != CTRUE.Id()) && 
                (F_boolean_I_any(F_stop_ask_integer(r.Skipc())).Id() != CTRUE.Id())) /* If:6 */{ 
              void_try6 = F_Serror_string(MakeString("[167] missing ) or , after ~S"),MakeConstantList(_Zx.Id()))
              } else {
              void_try6 = EID{CFALSE.Id(),0}
              /* If-6 */} 
            /* ERROR PROTECTION INSERTED (void_try6-void_try6) */
            if ErrorIn(void_try6) {Result = void_try6
            break
            } else {
            }}}
            /* while-5 */} 
          }
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          r.Next()
          var g0241I *ClaireBoolean  
          var g0241I_try02425 EID 
          /* and:5 */{ 
            var v_and5 *ClaireBoolean  
            
            v_and5 = Core.F__I_equal_any(e.Id(),C_none.Id())
            if (v_and5 == CFALSE) {g0241I_try02425 = EID{CFALSE.Id(),0}
            } else /* arg:6 */{ 
              v_and5 = Core.F__I_equal_any(F_boolean_I_any(F_stop_ask_integer(r.Skipc())).Id(),CTRUE.Id())
              if (v_and5 == CFALSE) {g0241I_try02425 = EID{CFALSE.Id(),0}
              } else /* arg:7 */{ 
                var v_and5_try02438 EID 
                /* Let:8 */{ 
                  var g0244UU *ClaireAny  
                  /* noccur = 1 */
                  var g0244UU_try02459 EID 
                  g0244UU_try02459 = r.Nexte()
                  /* ERROR PROTECTION INSERTED (g0244UU-v_and5_try02438) */
                  if ErrorIn(g0244UU_try02459) {v_and5_try02438 = g0244UU_try02459
                  } else {
                  g0244UU = ANY(g0244UU_try02459)
                  v_and5_try02438 = EID{Core.F__I_equal_any(g0244UU,e.Id()).Id(),0}
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (v_and5-g0241I_try02425) */
                if ErrorIn(v_and5_try02438) {g0241I_try02425 = v_and5_try02438
                } else {
                v_and5 = ToBoolean(OBJ(v_and5_try02438))
                if (v_and5 == CFALSE) {g0241I_try02425 = EID{CFALSE.Id(),0}
                } else /* arg:8 */{ 
                  g0241I_try02425 = EID{CTRUE.Id(),0}/* arg-8 */} 
                /* arg-7 */} 
              /* arg-6 */} 
            }
            /* and-5 */} 
          /* ERROR PROTECTION INSERTED (g0241I-Result) */
          if ErrorIn(g0241I_try02425) {Result = g0241I_try02425
          } else {
          g0241I = ToBoolean(OBJ(g0241I_try02425))
          if (g0241I == CTRUE) /* If:5 */{ 
            Result = F_Serror_string(MakeString("[161] missing ~S after ~S"),MakeConstantList(e.Id(),_Zx.Id()))
            } else {
            Result = EID{CFALSE.Id(),0}
            /* If-5 */} 
          }
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          Result = EID{_Zx.Id(),0}
          }}
          /* Let-4 */} 
        /* Let-3 */} 
      }
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: readcase @ meta_reader (throw: true) 
func E_readcase_meta_reader (r EID,e EID) EID { 
    return /*(sm for readcase @ meta_reader= EID)*/ ToMetaReader(OBJ(r)).Readcase(ToKeyword(OBJ(e)) )} 
  
// if the expression begins with "{"
//
/* {1} OPT.The go function for: readset(r:meta_reader,%a1:any) [] */
func (r *MetaReader ) Readset (_Za1 *ClaireAny ) EID { 
    var Result EID 
    if (Equal(_Za1,r.Curly) == CTRUE) /* If:2 */{ 
      r.Next()
      Result = EID{CEMPTY.Id(),0}
      } else {
      if (F_keyword_ask_any(_Za1) == CTRUE) /* If:3 */{ 
        var _Za1_try02464 EID 
        _Za1_try02464 = r.Nextstruct(ToKeyword(_Za1),C_none)
        /* ERROR PROTECTION INSERTED (_Za1-Result) */
        if ErrorIn(_Za1_try02464) {Result = _Za1_try02464
        } else {
        _Za1 = ANY(_Za1_try02464)
        Result = _Za1.ToEID()
        }
        } else {
        Result = EID{CFALSE.Id(),0}
        /* If-3 */} 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      /* Let:3 */{ 
        var _Za2 *ClaireAny  
        /* noccur = 7 */
        var _Za2_try02474 EID 
        _Za2_try02474 = r.Nexte()
        /* ERROR PROTECTION INSERTED (_Za2-Result) */
        if ErrorIn(_Za2_try02474) {Result = _Za2_try02474
        } else {
        _Za2 = ANY(_Za2_try02474)
        if (Equal(_Za2,r.Comma) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var g0248UU *ClaireSet  
            /* noccur = 1 */
            var g0248UU_try02496 EID 
            /* Let:6 */{ 
              var u_bag *ClaireSet   = ToType(CEMPTY.Id()).EmptySet()
              /* noccur = 2 */
              /* For:7 */{ 
                var u *ClaireAny  
                _ = u
                g0248UU_try02496= EID{CFALSE.Id(),0}
                var u_support *ClaireList  
                var u_support_try02508 EID 
                /* Let:8 */{ 
                  var g0251UU *ClaireAny  
                  /* noccur = 1 */
                  var g0251UU_try02529 EID 
                  g0251UU_try02529 = r.Cnext().Nextseq(125)
                  /* ERROR PROTECTION INSERTED (g0251UU-u_support_try02508) */
                  if ErrorIn(g0251UU_try02529) {u_support_try02508 = g0251UU_try02529
                  } else {
                  g0251UU = ANY(g0251UU_try02529)
                  u_support_try02508 = EID{F_cons_any(_Za1,ToList(g0251UU)).Id(),0}
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (u_support-g0248UU_try02496) */
                if ErrorIn(u_support_try02508) {g0248UU_try02496 = u_support_try02508
                } else {
                u_support = ToList(OBJ(u_support_try02508))
                u_len := u_support.Length()
                for i_it := 0; i_it < u_len; i_it++ { 
                  u = u_support.At(i_it)
                  var void_try9 EID 
                  _ = void_try9
                  /* Let:9 */{ 
                    var g0253UU *ClaireAny  
                    /* noccur = 1 */
                    var g0253UU_try025410 EID 
                    g0253UU_try025410 = F_dereference_any(u)
                    /* ERROR PROTECTION INSERTED (g0253UU-void_try9) */
                    if ErrorIn(g0253UU_try025410) {void_try9 = g0253UU_try025410
                    } else {
                    g0253UU = ANY(g0253UU_try025410)
                    void_try9 = EID{u_bag.AddFast(g0253UU).Id(),0}
                    }
                    /* Let-9 */} 
                  /* ERROR PROTECTION INSERTED (void_try9-g0248UU_try02496) */
                  if ErrorIn(void_try9) {g0248UU_try02496 = void_try9
                  g0248UU_try02496 = void_try9
                  break
                  } else {
                  }}
                  /* loop-8 */} 
                /* For-7 */} 
              /* ERROR PROTECTION INSERTED (g0248UU_try02496-g0248UU_try02496) */
              if !ErrorIn(g0248UU_try02496) {
              g0248UU_try02496 = EID{u_bag.Id(),0}
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (g0248UU-Result) */
            if ErrorIn(g0248UU_try02496) {Result = g0248UU_try02496
            } else {
            g0248UU = ToSet(OBJ(g0248UU_try02496))
            Result = EID{g0248UU.Cast_I(ToType(CEMPTY.Id())).Id(),0}
            }
            /* Let-5 */} 
          /* If!4 */}  else if (Equal(_Za2,r.Curly) == CTRUE) /* If:4 */{ 
          r.Next()
          /* Let:5 */{ 
            var g0255UU *ClaireSet  
            /* noccur = 1 */
            var g0255UU_try02566 EID 
            /* Construct:6 */{ 
              var v_bag_arg *ClaireAny  
              g0255UU_try02566= EID{ToType(CEMPTY.Id()).EmptySet().Id(),0}
              var v_bag_arg_try02577 EID 
              v_bag_arg_try02577 = F_dereference_any(_Za1)
              /* ERROR PROTECTION INSERTED (v_bag_arg-g0255UU_try02566) */
              if ErrorIn(v_bag_arg_try02577) {g0255UU_try02566 = v_bag_arg_try02577
              } else {
              v_bag_arg = ANY(v_bag_arg_try02577)
              ToSet(OBJ(g0255UU_try02566)).AddFast(v_bag_arg)}
              /* Construct-6 */} 
            /* ERROR PROTECTION INSERTED (g0255UU-Result) */
            if ErrorIn(g0255UU_try02566) {Result = g0255UU_try02566
            } else {
            g0255UU = ToSet(OBJ(g0255UU_try02566))
            Result = EID{g0255UU.Cast_I(ToType(CEMPTY.Id())).Id(),0}
            }
            /* Let-5 */} 
          /* If!4 */}  else if (_Za2 == C_in.Id()) /* If:4 */{ 
          /* Let:5 */{ 
            var v *ClaireVariable  
            /* noccur = 2 */
            var v_try02586 EID 
            v_try02586 = F_extract_variable_any(_Za1)
            /* ERROR PROTECTION INSERTED (v-Result) */
            if ErrorIn(v_try02586) {Result = v_try02586
            } else {
            v = To_Variable(OBJ(v_try02586))
            /* Let:6 */{ 
              var _CL_obj *Language.Select   = Language.To_Select(new(Language.Select).Is(Language.C_Select))
              /* noccur = 5 */
              _CL_obj.ClaireVar = v
              /* update:7 */{ 
                var va_arg1 *Language.Iteration  
                var va_arg2 *ClaireAny  
                va_arg1 = Language.To_Iteration(_CL_obj.Id())
                var va_arg2_try02598 EID 
                va_arg2_try02598 = r.Nexte()
                /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                if ErrorIn(va_arg2_try02598) {Result = va_arg2_try02598
                } else {
                va_arg2 = ANY(va_arg2_try02598)
                /* ---------- now we compile update iClaire/set_arg(va_arg1) := va_arg2 ------- */
                va_arg1.SetArg = va_arg2
                Result = va_arg2.ToEID()
                }
                /* update-7 */} 
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              /* update:7 */{ 
                var va_arg1 *Language.Iteration  
                var va_arg2 *ClaireAny  
                va_arg1 = Language.To_Iteration(_CL_obj.Id())
                var va_arg2_try02608 EID 
                /* Let:8 */{ 
                  var _Zbind *ClaireList   = r.Bind_I(v)
                  /* noccur = 1 */
                  /* LetE:9 */{ 
                    var x EID 
                    var g0261I *ClaireBoolean  
                    var g0261I_try026210 EID 
                    /* Let:10 */{ 
                      var g0263UU *ClaireAny  
                      /* noccur = 1 */
                      var g0263UU_try026411 EID 
                      g0263UU_try026411 = r.Nexte()
                      /* ERROR PROTECTION INSERTED (g0263UU-g0261I_try026210) */
                      if ErrorIn(g0263UU_try026411) {g0261I_try026210 = g0263UU_try026411
                      } else {
                      g0263UU = ANY(g0263UU_try026411)
                      g0261I_try026210 = EID{Core.F__I_equal_any(g0263UU,C_OR.Value).Id(),0}
                      }
                      /* Let-10 */} 
                    /* ERROR PROTECTION INSERTED (g0261I-x) */
                    if ErrorIn(g0261I_try026210) {x = g0261I_try026210
                    } else {
                    g0261I = ToBoolean(OBJ(g0261I_try026210))
                    if (g0261I == CTRUE) /* If:10 */{ 
                      x = F_Serror_string(MakeString("[168] missing | in selection"),CNIL)
                      } else {
                      x = F_nexts_I_meta_reader2(r,125)
                      /* If-10 */} 
                    }
                    /* ERROR PROTECTION INSERTED (x-va_arg2_try02608) */
                    if ErrorIn(x) {va_arg2_try02608 = x
                    } else {
                    r.Unbind_I(_Zbind)
                    va_arg2_try02608 = x}
                    /* LetE-9 */} 
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                if ErrorIn(va_arg2_try02608) {Result = va_arg2_try02608
                } else {
                va_arg2 = ANY(va_arg2_try02608)
                /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                va_arg1.Arg = va_arg2
                Result = va_arg2.ToEID()
                }
                /* update-7 */} 
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              Result = EID{_CL_obj.Id(),0}
              }}
              /* Let-6 */} 
            }
            /* Let-5 */} 
          /* If!4 */}  else if (Equal(_Za2,C_OR.Value) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var v *ClaireVariable  
            /* noccur = 4 */
            var v_try02656 EID 
            /* Let:6 */{ 
              var g0266UU *ClaireAny  
              /* noccur = 1 */
              var g0266UU_try02677 EID 
              g0266UU_try02677 = F_nexts_I_meta_reader1(r,C_in)
              /* ERROR PROTECTION INSERTED (g0266UU-v_try02656) */
              if ErrorIn(g0266UU_try02677) {v_try02656 = g0266UU_try02677
              } else {
              g0266UU = ANY(g0266UU_try02677)
              v_try02656 = F_extract_variable_any(g0266UU)
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (v-Result) */
            if ErrorIn(v_try02656) {Result = v_try02656
            } else {
            v = To_Variable(OBJ(v_try02656))
            /* Let:6 */{ 
              var g0268UU *Language.Image  
              /* noccur = 1 */
              var g0268UU_try02697 EID 
              /* Let:7 */{ 
                var _CL_obj *Language.Image   = Language.To_Image(new(Language.Image).Is(Language.C_Image))
                /* noccur = 5 */
                _CL_obj.ClaireVar = v
                /* update:8 */{ 
                  var va_arg1 *Language.Iteration  
                  var va_arg2 *ClaireAny  
                  va_arg1 = Language.To_Iteration(_CL_obj.Id())
                  var va_arg2_try02709 EID 
                  va_arg2_try02709 = F_nexts_I_meta_reader2(r,125)
                  /* ERROR PROTECTION INSERTED (va_arg2-g0268UU_try02697) */
                  if ErrorIn(va_arg2_try02709) {g0268UU_try02697 = va_arg2_try02709
                  } else {
                  va_arg2 = ANY(va_arg2_try02709)
                  /* ---------- now we compile update iClaire/set_arg(va_arg1) := va_arg2 ------- */
                  va_arg1.SetArg = va_arg2
                  g0268UU_try02697 = va_arg2.ToEID()
                  }
                  /* update-8 */} 
                /* ERROR PROTECTION INSERTED (g0268UU_try02697-g0268UU_try02697) */
                if !ErrorIn(g0268UU_try02697) {
                _CL_obj.Arg = Language.F_substitution_any(_Za1,v,v.Id())
                g0268UU_try02697 = EID{_CL_obj.Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (g0268UU-Result) */
              if ErrorIn(g0268UU_try02697) {Result = g0268UU_try02697
              } else {
              g0268UU = Language.To_Image(OBJ(g0268UU_try02697))
              Result = Language.F_lexical_build_any(g0268UU.Id(),MakeConstantList(v.Id()),0)
              }
              /* Let-6 */} 
            }
            /* Let-5 */} 
          /* If!4 */}  else if (F_operation_ask_any(_Za2) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var g0271UU *ClaireAny  
            /* noccur = 1 */
            var g0271UU_try02726 EID 
            /* Let:6 */{ 
              var g0273UU *ClaireAny  
              /* noccur = 1 */
              var g0273UU_try02747 EID 
              /* Let:7 */{ 
                var g0275UU *ClaireAny  
                /* noccur = 1 */
                var g0275UU_try02768 EID 
                g0275UU_try02768 = r.Nexte()
                /* ERROR PROTECTION INSERTED (g0275UU-g0273UU_try02747) */
                if ErrorIn(g0275UU_try02768) {g0273UU_try02747 = g0275UU_try02768
                } else {
                g0275UU = ANY(g0275UU_try02768)
                g0273UU_try02747 = F_combine_any(_Za1,_Za2,g0275UU)
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (g0273UU-g0271UU_try02726) */
              if ErrorIn(g0273UU_try02747) {g0271UU_try02726 = g0273UU_try02747
              } else {
              g0273UU = ANY(g0273UU_try02747)
              g0271UU_try02726 = r.Loopexp(g0273UU,C_none,CFALSE)
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (g0271UU-Result) */
            if ErrorIn(g0271UU_try02726) {Result = g0271UU_try02726
            } else {
            g0271UU = ANY(g0271UU_try02726)
            Result = r.Readset(g0271UU)
            }
            /* Let-5 */} 
          } else {
          Result = F_Serror_string(MakeString("[169] missing separation between ~S and ~S"),MakeConstantList(_Za1,_Za2))
          /* If-4 */} 
        }
        /* Let-3 */} 
      }
      /* If-2 */} 
    return Result} 
  
// The EID go function for: readset @ meta_reader (throw: true) 
func E_readset_meta_reader (r EID,_Za1 EID) EID { 
    return /*(sm for readset @ meta_reader= EID)*/ ToMetaReader(OBJ(r)).Readset(ANY(_Za1) )} 
  
/* {1} OPT.The go function for: dereference(x:any) [] */
func F_dereference_any (x *ClaireAny ) EID { 
    var Result EID 
    if (x.Isa.IsIn(C_unbound_symbol) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0277 *ClaireUnboundSymbol   = ToUnboundSymbol(x)
        /* noccur = 1 */
        Result = ToException(Core.C_general_error.Make(MakeString("[170] cannot use ~S in a set constant").Id(),MakeConstantList(g0277.Id()).Id())).Close()
        /* Let-3 */} 
      /* If!2 */}  else if (x.Isa.IsIn(C_Variable) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0278 *ClaireVariable   = To_Variable(x)
        /* noccur = 1 */
        Result = ToException(Core.C_general_error.Make(MakeString("[170] cannot use a variable (~S) in a set constant").Id(),MakeConstantList(g0278.Id()).Id())).Close()
        /* Let-3 */} 
      } else {
      Result = EVAL(x)
      /* If-2 */} 
    return Result} 
  
// The EID go function for: dereference @ any (throw: true) 
func E_dereference_any (x EID) EID { 
    return /*(sm for dereference @ any= EID)*/ F_dereference_any(ANY(x) )} 
  
// reads a sequence of exp. Must end with a e = ) | ] | }
//
/* {1} OPT.The go function for: nextseq(r:meta_reader,e:integer) [] */
func (r *MetaReader ) Nextseq (e int) EID { 
    var Result EID 
    if (r.Firstc() == e) /* If:2 */{ 
      r.Next()
      Result = EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
      } else {
      /* Let:3 */{ 
        var x *ClaireAny  
        /* noccur = 2 */
        var x_try02804 EID 
        x_try02804 = r.Nexts(C_none)
        /* ERROR PROTECTION INSERTED (x-Result) */
        if ErrorIn(x_try02804) {Result = x_try02804
        } else {
        x = ANY(x_try02804)
        if ((r.Firstc() == 10) && 
            (r.Toplevel == CTRUE)) /* If:4 */{ 
          r.Skipc()
          /* If-4 */} 
        if (r.Firstc() == e) /* If:4 */{ 
          r.Next()
          Result = EID{MakeConstantList(x).Id(),0}
          /* If!4 */}  else if (r.Firstc() == 44) /* If:4 */{ 
          /* Let:5 */{ 
            var g0281UU *ClaireAny  
            /* noccur = 1 */
            var g0281UU_try02826 EID 
            g0281UU_try02826 = r.Cnext().Nextseq(e)
            /* ERROR PROTECTION INSERTED (g0281UU-Result) */
            if ErrorIn(g0281UU_try02826) {Result = g0281UU_try02826
            } else {
            g0281UU = ANY(g0281UU_try02826)
            Result = EID{F_cons_any(x,ToList(g0281UU)).Id(),0}
            }
            /* Let-5 */} 
          } else {
          Result = F_Serror_string(MakeString("[171] Read the character ~S inside a sequence"),MakeConstantList(MakeChar(F_char_I_integer(r.Firstc())).Id()))
          /* If-4 */} 
        }
        /* Let-3 */} 
      /* If-2 */} 
    return Result} 
  
// The EID go function for: nextseq @ meta_reader (throw: true) 
func E_nextseq_meta_reader (r EID,e EID) EID { 
    return /*(sm for nextseq @ meta_reader= EID)*/ ToMetaReader(OBJ(r)).Nextseq(INT(e) )} 
  
// read the next block: a sequence of exp. Must end with a e = ) | ] | }
//
/* {1} OPT.The go function for: readblock(r:meta_reader,x:any,e:integer) [] */
func (r *MetaReader ) Readblock (x *ClaireAny ,e int) EID { 
    var Result EID 
    r.Skipc()
    if (Equal(x,r.Paren) == CTRUE) /* If:2 */{ 
      Result = EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
      /* If!2 */}  else if (r.Firstc() == 44) /* If:2 */{ 
      /* Let:3 */{ 
        var g0284UU *ClaireAny  
        /* noccur = 1 */
        var g0284UU_try02854 EID 
        /* Let:4 */{ 
          var g0286UU *ClaireAny  
          /* noccur = 1 */
          var g0286UU_try02875 EID 
          g0286UU_try02875 = r.Cnext().Nexte()
          /* ERROR PROTECTION INSERTED (g0286UU-g0284UU_try02854) */
          if ErrorIn(g0286UU_try02875) {g0284UU_try02854 = g0286UU_try02875
          } else {
          g0286UU = ANY(g0286UU_try02875)
          g0284UU_try02854 = r.Readblock(g0286UU,e)
          }
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (g0284UU-Result) */
        if ErrorIn(g0284UU_try02854) {Result = g0284UU_try02854
        } else {
        g0284UU = ANY(g0284UU_try02854)
        Result = F_Do_I_any(x,g0284UU)
        }
        /* Let-3 */} 
      /* If!2 */}  else if (r.Firstc() == e) /* If:2 */{ 
      r.Cnext()
      Result = x.ToEID()
      /* If!2 */}  else if (F_boolean_I_any(F_stop_ask_integer(r.Firstc())) == CTRUE) /* If:2 */{ 
      Result = F_Serror_string(MakeString("[172] the sequence ...~S must end with ~A"),MakeConstantList(x,MakeChar(F_char_I_integer(e)).Id()))
      /* If!2 */}  else if (x == C_if.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var g0288UU *ClaireAny  
        /* noccur = 1 */
        var g0288UU_try02894 EID 
        g0288UU_try02894 = r.Readif(e)
        /* ERROR PROTECTION INSERTED (g0288UU-Result) */
        if ErrorIn(g0288UU_try02894) {Result = g0288UU_try02894
        } else {
        g0288UU = ANY(g0288UU_try02894)
        Result = r.Readblock(g0288UU,e)
        }
        /* Let-3 */} 
      /* If!2 */}  else if (x == C_Zif.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var _Zi *Language.If  
        /* noccur = 3 */
        var _Zi_try02904 EID 
        _Zi_try02904 = r.Readif(e)
        /* ERROR PROTECTION INSERTED (_Zi-Result) */
        if ErrorIn(_Zi_try02904) {Result = _Zi_try02904
        } else {
        _Zi = Language.To_If(OBJ(_Zi_try02904))
        /* Let:4 */{ 
          var g0291UU *ClaireAny  
          /* noccur = 1 */
          var g0291UU_try02925 EID 
          var g0293I *ClaireBoolean  
          var g0293I_try02945 EID 
          /* Let:5 */{ 
            var g0295UU *ClaireAny  
            /* noccur = 1 */
            var g0295UU_try02966 EID 
            g0295UU_try02966 = EVAL(_Zi.Test)
            /* ERROR PROTECTION INSERTED (g0295UU-g0293I_try02945) */
            if ErrorIn(g0295UU_try02966) {g0293I_try02945 = g0295UU_try02966
            } else {
            g0295UU = ANY(g0295UU_try02966)
            g0293I_try02945 = EID{F_boolean_I_any(g0295UU).Id(),0}
            }
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (g0293I-g0291UU_try02925) */
          if ErrorIn(g0293I_try02945) {g0291UU_try02925 = g0293I_try02945
          } else {
          g0293I = ToBoolean(OBJ(g0293I_try02945))
          if (g0293I == CTRUE) /* If:5 */{ 
            g0291UU_try02925 = _Zi.Arg.ToEID()
            } else {
            g0291UU_try02925 = _Zi.Other.ToEID()
            /* If-5 */} 
          }
          /* ERROR PROTECTION INSERTED (g0291UU-Result) */
          if ErrorIn(g0291UU_try02925) {Result = g0291UU_try02925
          } else {
          g0291UU = ANY(g0291UU_try02925)
          Result = r.Readblock(g0291UU,e)
          }
          /* Let-4 */} 
        }
        /* Let-3 */} 
      /* If!2 */}  else if (x == C_else.Id()) /* If:2 */{ 
      Result = F_Serror_string(MakeString("[173] Expression starting with else"),CNIL)
      /* If!2 */}  else if (F_keyword_ask_any(x) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0297UU *ClaireAny  
        /* noccur = 1 */
        var g0297UU_try02984 EID 
        g0297UU_try02984 = r.Nextstruct(ToKeyword(x),C_none)
        /* ERROR PROTECTION INSERTED (g0297UU-Result) */
        if ErrorIn(g0297UU_try02984) {Result = g0297UU_try02984
        } else {
        g0297UU = ANY(g0297UU_try02984)
        Result = r.Readblock(g0297UU,e)
        }
        /* Let-3 */} 
      } else {
      /* Let:3 */{ 
        var y *ClaireAny  
        /* noccur = 3 */
        var y_try02994 EID 
        y_try02994 = r.Loopexp(x,C_none,CFALSE)
        /* ERROR PROTECTION INSERTED (y-Result) */
        if ErrorIn(y_try02994) {Result = y_try02994
        } else {
        y = ANY(y_try02994)
        if (y.Isa.IsIn(Language.C_Call_star) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var g0283 *Language.Call_star   = Language.To_Call_star(y)
            /* noccur = 1 */
            g0283.Isa = Language.C_Call
            /* Let-5 */} 
          /* If-4 */} 
        Result = r.Readblock(y,e)
        }
        /* Let-3 */} 
      /* If-2 */} 
    return Result} 
  
// The EID go function for: readblock @ meta_reader (throw: true) 
func E_readblock_meta_reader (r EID,x EID,e EID) EID { 
    return /*(sm for readblock @ meta_reader= EID)*/ ToMetaReader(OBJ(r)).Readblock(ANY(x),INT(e) )} 
  
/* {1} OPT.The go function for: Do!(x:any,y:any) [] */
func F_Do_I_any (x *ClaireAny ,y *ClaireAny ) EID { 
    var Result EID 
    if (y.Isa.IsIn(Language.C_Do) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0300 *Language.Do   = Language.To_Do(y)
        /* noccur = 3 */
        /* update:4 */{ 
          var va_arg1 *Language.Do  
          var va_arg2 *ClaireList  
          va_arg1 = g0300
          var va_arg2_try03025 EID 
          va_arg2_try03025 = g0300.Args.Nth_plus(1,x)
          /* ERROR PROTECTION INSERTED (va_arg2-Result) */
          if ErrorIn(va_arg2_try03025) {Result = va_arg2_try03025
          } else {
          va_arg2 = ToList(OBJ(va_arg2_try03025))
          /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
          va_arg1.Args = va_arg2
          Result = EID{va_arg2.Id(),0}
          }
          /* update-4 */} 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        Result = EID{g0300.Id(),0}
        }
        /* Let-3 */} 
      } else {
      /* Let:3 */{ 
        var _CL_obj *Language.Do   = Language.To_Do(new(Language.Do).Is(Language.C_Do))
        /* noccur = 3 */
        _CL_obj.Args = MakeList(ToType(C_any.Id()),x,y)
        Result = EID{_CL_obj.Id(),0}
        /* Let-3 */} 
      /* If-2 */} 
    return Result} 
  
// The EID go function for: Do! @ any (throw: true) 
func E_Do_I_any (x EID,y EID) EID { 
    return /*(sm for Do! @ any= EID)*/ F_Do_I_any(ANY(x),ANY(y) )} 
  
// extract the type from a list<X> expression
/* {1} OPT.The go function for: extract_of_type(x:Call) [] */
func F_extract_of_type_Call (x *Language.Call ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var l *ClaireList   = x.Args
      /* noccur = 2 */
      if (l.Length() > 2) /* If:3 */{ 
        /* Let:4 */{ 
          var y *ClaireAny   = l.At(3-1)
          /* noccur = 2 */
          if (y.Isa.IsIn(Language.C_List) == CTRUE) /* If:5 */{ 
            /* Let:6 */{ 
              var g0303 *Language.List   = Language.To_List(y)
              /* noccur = 1 */
              /* Let:7 */{ 
                var z *ClaireAny   = g0303.Args.At(1-1)
                /* noccur = 2 */
                if (z.Isa.IsIn(Language.C_Set) == CTRUE) /* If:8 */{ 
                  /* Let:9 */{ 
                    var g0304 *Language.Set   = Language.To_Set(z)
                    /* noccur = 1 */
                    Result = Language.F_extract_type_any(g0304.Args.At(1-1))
                    /* Let-9 */} 
                  } else {
                  Result = EID{C_any.Id(),0}
                  /* If-8 */} 
                /* Let-7 */} 
              /* Let-6 */} 
            } else {
            Result = EID{C_any.Id(),0}
            /* If-5 */} 
          /* Let-4 */} 
        } else {
        Result = EID{C_any.Id(),0}
        /* If-3 */} 
      /* Let-2 */} 
    return RangeCheck(ToType(C_type.Id()),Result)} 
  
// The EID go function for: extract_of_type @ Call (throw: true) 
func E_extract_of_type_Call (x EID) EID { 
    return /*(sm for extract_of_type @ Call= EID)*/ F_extract_of_type_Call(Language.To_Call(OBJ(x)) )} 
  
// **********************************************************************
// *   Part 3: read functional calls                                    *
// **********************************************************************
// store the line number in debug mode
// in v4.0 we will not do this for JITO calls :)
// this is a cool trick when operating in debug mode: we store the last evaluated
// call so we can tell very simply which last call triggered the error
//
/* {1} OPT.The go function for: DBregister(c:Call) [] */
func F_DBregister_Call (c *Language.Call ) EID { 
    var Result EID 
    if (ClEnv.Debug_I >= 0) /* If:2 */{ 
      Language.C_iClaire_LastCall.Value = c.Id()
      Core.F_put_table(C_Reader_DBline,c.Id(),MakeInteger(ClEnv.NLine).Id())
      /* If-2 */} 
    if ((c.Selector.Id() == C_store.Id()) && 
        (c.Args.Length() == 1)) /* If:2 */{ 
      /* Let:3 */{ 
        var l *ClaireList   = c.Args
        /* noccur = 3 */
        if (l.At(1-1).Isa.IsIn(Core.C_global_variable) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var g0307UU *ClaireString  
            /* noccur = 1 */
            var g0307UU_try03086 EID 
            g0307UU_try03086 = Core.F_make_string_symbol(ToSymbol(OBJ(Core.F_CALL(C_name,ARGS(l.At(1-1).ToEID())))))
            /* ERROR PROTECTION INSERTED (g0307UU-Result) */
            if ErrorIn(g0307UU_try03086) {Result = g0307UU_try03086
            } else {
            g0307UU = ToString(OBJ(g0307UU_try03086))
            Result = ToArray(l.Id()).NthPut(1,(g0307UU).Id()).ToEID()
            }
            /* Let-5 */} 
          } else {
          Result = EID{CFALSE.Id(),0}
          /* If-4 */} 
        /* Let-3 */} 
      } else {
      Result = EID{CFALSE.Id(),0}
      /* If-2 */} 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    Result = EID{c.Id(),0}
    }
    return Result} 
  
// The EID go function for: DBregister @ Call (throw: true) 
func E_DBregister_Call (c EID) EID { 
    return /*(sm for DBregister @ Call= EID)*/ F_DBregister_Call(Language.To_Call(OBJ(c)) )} 
  
/* {1} OPT.The go function for: Call!(p:property,l:list) [] */
func F_Call_I_property (p *ClaireProperty ,l *ClaireList ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var g0309UU *Language.Call  
      /* noccur = 1 */
      /* Let:3 */{ 
        var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
        /* noccur = 5 */
        _CL_obj.Selector = p
        _CL_obj.Args = l
        g0309UU = _CL_obj
        /* Let-3 */} 
      Result = F_DBregister_Call(g0309UU)
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: Call! @ property (throw: true) 
func E_Call_I_property (p EID,l EID) EID { 
    return /*(sm for Call! @ property= EID)*/ F_Call_I_property(ToProperty(OBJ(p)),ToList(OBJ(l)) )} 
  
// if the expression is a call -------------------------------------------
// x is the first token that we have read x(...)
// there are many special case (x is not a propery) then the regular case
// t is a type when x was read as (p@t)
/* {1} OPT.The go function for: readcall(r:meta_reader,x:any,t:any) [] */
func (r *MetaReader ) Readcall (x *ClaireAny ,t *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var l *ClaireAny  
      /* noccur = 26 */
      var l_try03183 EID 
      l_try03183 = r.Cnext().Nextseq(41)
      /* ERROR PROTECTION INSERTED (l-Result) */
      if ErrorIn(l_try03183) {Result = l_try03183
      } else {
      l = ANY(l_try03183)
      if (x == C_printf.Id()) /* If:3 */{ 
        /* Let:4 */{ 
          var _CL_obj *Language.Printf   = Language.To_Printf(new(Language.Printf).Is(Language.C_Printf))
          /* noccur = 3 */
          _CL_obj.Args = ToList(l)
          Result = EID{_CL_obj.Id(),0}
          /* Let-4 */} 
        /* If!3 */}  else if (x == C_error.Id()) /* If:3 */{ 
        /* Let:4 */{ 
          var _CL_obj *Language.Error   = Language.To_Error(new(Language.Error).Is(Language.C_Error))
          /* noccur = 3 */
          _CL_obj.Args = ToList(l)
          Result = EID{_CL_obj.Id(),0}
          /* Let-4 */} 
        /* If!3 */}  else if (x == C_assert.Id()) /* If:3 */{ 
        /* Let:4 */{ 
          var _CL_obj *Language.Assert   = Language.To_Assert(new(Language.Assert).Is(Language.C_Assert))
          /* noccur = 7 */
          _CL_obj.Args = ToList(l)
          _CL_obj.Index = ClEnv.NLine
          _CL_obj.External = r.External
          Result = EID{_CL_obj.Id(),0}
          /* Let-4 */} 
        /* If!3 */}  else if (x == C_trace.Id()) /* If:3 */{ 
        /* Let:4 */{ 
          var _CL_obj *Language.Trace   = Language.To_Trace(new(Language.Trace).Is(Language.C_Trace))
          /* noccur = 3 */
          _CL_obj.Args = ToList(l)
          Result = EID{_CL_obj.Id(),0}
          /* Let-4 */} 
        /* If!3 */}  else if (x == C_branch.Id()) /* If:3 */{ 
        /* Let:4 */{ 
          var _CL_obj *Language.Branch   = Language.To_Branch(new(Language.Branch).Is(Language.C_Branch))
          /* noccur = 3 */
          _CL_obj.Args = ToList(l)
          Result = EID{_CL_obj.Id(),0}
          /* Let-4 */} 
        /* If!3 */}  else if (x == C_quote.Id()) /* If:3 */{ 
        /* Let:4 */{ 
          var g0319UU *ClaireAny  
          /* noccur = 1 */
          var g0319UU_try03205 EID 
          if (F_boolean_I_any(l) == CTRUE) /* If:5 */{ 
            g0319UU_try03205 = Core.F_CALL(C_nth,ARGS(l.ToEID(),EID{C__INT,IVAL(1)}))
            } else {
            g0319UU_try03205 = EID{CFALSE.Id(),0}
            /* If-5 */} 
          /* ERROR PROTECTION INSERTED (g0319UU-Result) */
          if ErrorIn(g0319UU_try03205) {Result = g0319UU_try03205
          } else {
          g0319UU = ANY(g0319UU_try03205)
          Result = Language.C_Quote.Make(g0319UU).ToEID()
          }
          /* Let-4 */} 
        /* If!3 */}  else if (x == C_tuple.Id()) /* If:3 */{ 
        /* Let:4 */{ 
          var _CL_obj *Language.Tuple   = Language.To_Tuple(new(Language.Tuple).Is(Language.C_Tuple))
          /* noccur = 3 */
          _CL_obj.Args = ToList(l)
          Result = EID{_CL_obj.Id(),0}
          /* Let-4 */} 
        /* If!3 */}  else if (x == C_list.Id()) /* If:3 */{ 
        /* Let:4 */{ 
          var _CL_obj *Language.List   = Language.To_List(new(Language.List).Is(Language.C_List))
          /* noccur = 3 */
          _CL_obj.Args = ToList(l)
          Result = EID{_CL_obj.Id(),0}
          /* Let-4 */} 
        } else {
        var g0321I *ClaireBoolean  
        if (x.Isa.IsIn(Language.C_Call) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var g0310 *Language.Call   = Language.To_Call(x)
            /* noccur = 1 */
            g0321I = Equal(g0310.Args.At(1-1),C_list.Id())
            /* Let-5 */} 
          } else {
          g0321I = CFALSE
          /* If-4 */} 
        if (g0321I == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var _CL_obj *Language.List   = Language.To_List(new(Language.List).Is(Language.C_List))
            /* noccur = 5 */
            /* update:6 */{ 
              var va_arg1 *Language.List  
              var va_arg2 *ClaireType  
              va_arg1 = _CL_obj
              var va_arg2_try03227 EID 
              va_arg2_try03227 = F_extract_of_type_Call(Language.To_Call(x))
              /* ERROR PROTECTION INSERTED (va_arg2-Result) */
              if ErrorIn(va_arg2_try03227) {Result = va_arg2_try03227
              } else {
              va_arg2 = ToType(OBJ(va_arg2_try03227))
              /* ---------- now we compile update of(va_arg1) := va_arg2 ------- */
              va_arg1.Of = va_arg2
              Result = EID{va_arg2.Id(),0}
              }
              /* update-6 */} 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            _CL_obj.Args = ToList(l)
            Result = EID{_CL_obj.Id(),0}
            }
            /* Let-5 */} 
          } else {
          var g0323I *ClaireBoolean  
          if (x.Isa.IsIn(Language.C_Call) == CTRUE) /* If:5 */{ 
            /* Let:6 */{ 
              var g0311 *Language.Call   = Language.To_Call(x)
              /* noccur = 1 */
              g0323I = Equal(g0311.Args.At(1-1),C_array.Id())
              /* Let-6 */} 
            } else {
            g0323I = CFALSE
            /* If-5 */} 
          if (g0323I == CTRUE) /* If:5 */{ 
            /* Let:6 */{ 
              var _CL_obj *Language.Array   = Language.To_Array(new(Language.Array).Is(Language.C_Array))
              /* noccur = 5 */
              /* update:7 */{ 
                var va_arg1 *Language.Array  
                var va_arg2 *ClaireType  
                va_arg1 = _CL_obj
                var va_arg2_try03248 EID 
                va_arg2_try03248 = F_extract_of_type_Call(Language.To_Call(x))
                /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                if ErrorIn(va_arg2_try03248) {Result = va_arg2_try03248
                } else {
                va_arg2 = ToType(OBJ(va_arg2_try03248))
                /* ---------- now we compile update of(va_arg1) := va_arg2 ------- */
                va_arg1.Of = va_arg2
                Result = EID{va_arg2.Id(),0}
                }
                /* update-7 */} 
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              _CL_obj.Args = ToList(l)
              Result = EID{_CL_obj.Id(),0}
              }
              /* Let-6 */} 
            } else {
            var g0325I *ClaireBoolean  
            if (x.Isa.IsIn(Language.C_Call) == CTRUE) /* If:6 */{ 
              /* Let:7 */{ 
                var g0312 *Language.Call   = Language.To_Call(x)
                /* noccur = 1 */
                g0325I = Equal(g0312.Args.At(1-1),C_set.Id())
                /* Let-7 */} 
              } else {
              g0325I = CFALSE
              /* If-6 */} 
            if (g0325I == CTRUE) /* If:6 */{ 
              /* Let:7 */{ 
                var _CL_obj *Language.Set   = Language.To_Set(new(Language.Set).Is(Language.C_Set))
                /* noccur = 5 */
                /* update:8 */{ 
                  var va_arg1 *Language.Set  
                  var va_arg2 *ClaireType  
                  va_arg1 = _CL_obj
                  var va_arg2_try03269 EID 
                  va_arg2_try03269 = F_extract_of_type_Call(Language.To_Call(x))
                  /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                  if ErrorIn(va_arg2_try03269) {Result = va_arg2_try03269
                  } else {
                  va_arg2 = ToType(OBJ(va_arg2_try03269))
                  /* ---------- now we compile update of(va_arg1) := va_arg2 ------- */
                  va_arg1.Of = va_arg2
                  Result = EID{va_arg2.Id(),0}
                  }
                  /* update-8 */} 
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                _CL_obj.Args = ToList(l)
                Result = EID{_CL_obj.Id(),0}
                }
                /* Let-7 */} 
              /* If!6 */}  else if ((C_class.Id() == x.Isa.Id()) && 
                (ToClass(x).IsIn(Language.C_Macro) == CTRUE)) /* If:6 */{ 
              /* Let:7 */{ 
                var o *ClaireObject   = F_new_object_class(ToClass(x))
                /* noccur = 2 */
                Result = Core.F_put_property2(C_args,o,l)
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                Result = EID{o.Id(),0}
                }
                /* Let-7 */} 
              /* If!6 */}  else if (x == C_set.Id()) /* If:6 */{ 
              /* Let:7 */{ 
                var _CL_obj *Language.Set   = Language.To_Set(new(Language.Set).Is(Language.C_Set))
                /* noccur = 3 */
                _CL_obj.Args = ToList(l)
                Result = EID{_CL_obj.Id(),0}
                /* Let-7 */} 
              /* If!6 */}  else if ((x == C_return.Id()) || 
                (x == C_break.Id())) /* If:6 */{ 
              /* Let:7 */{ 
                var g0327UU *ClaireAny  
                /* noccur = 1 */
                var g0327UU_try03288 EID 
                if (F_boolean_I_any(l) == CTRUE) /* If:8 */{ 
                  g0327UU_try03288 = Core.F_CALL(C_nth,ARGS(l.ToEID(),EID{C__INT,IVAL(1)}))
                  } else {
                  g0327UU_try03288 = EID{CTRUE.Id(),0}
                  /* If-8 */} 
                /* ERROR PROTECTION INSERTED (g0327UU-Result) */
                if ErrorIn(g0327UU_try03288) {Result = g0327UU_try03288
                } else {
                g0327UU = ANY(g0327UU_try03288)
                Result = Language.C_Return.Make(g0327UU).ToEID()
                }
                /* Let-7 */} 
              /* If!6 */}  else if (C_class.Id() == x.Isa.Id()) /* If:6 */{ 
              var g0329I *ClaireBoolean  
              var g0329I_try03307 EID 
              /* Let:7 */{ 
                var g0331UU *ClaireBoolean  
                /* noccur = 1 */
                var g0331UU_try03328 EID 
                /* Let:8 */{ 
                  var g0333UU *ClaireAny  
                  /* noccur = 1 */
                  var g0333UU_try03349 EID 
                  /* For:9 */{ 
                    var y *ClaireAny  
                    _ = y
                    g0333UU_try03349= EID{CFALSE.Id(),0}
                    var y_support *ClaireList  
                    var y_support_try033510 EID 
                    y_support_try033510 = Core.F_enumerate_any(l)
                    /* ERROR PROTECTION INSERTED (y_support-g0333UU_try03349) */
                    if ErrorIn(y_support_try033510) {g0333UU_try03349 = y_support_try033510
                    } else {
                    y_support = ToList(OBJ(y_support_try033510))
                    y_len := y_support.Length()
                    for i_it := 0; i_it < y_len; i_it++ { 
                      y = y_support.At(i_it)
                      var void_try11 EID 
                      _ = void_try11
                      var g0336I *ClaireBoolean  
                      var g0336I_try033711 EID 
                      /* Let:11 */{ 
                        var g0338UU *ClaireBoolean  
                        /* noccur = 1 */
                        var g0338UU_try033912 EID 
                        if (y.Isa.IsIn(Language.C_Call) == CTRUE) /* If:12 */{ 
                          /* Let:13 */{ 
                            var g0313 *Language.Call   = Language.To_Call(y)
                            /* noccur = 3 */
                            if (g0313.Selector.Id() == C__equal.Id()) /* If:14 */{ 
                              /* Let:15 */{ 
                                var g0340UU *ClaireProperty  
                                /* noccur = 1 */
                                var g0340UU_try034116 EID 
                                g0340UU_try034116 = Language.F_make_a_property_any(g0313.Args.At(1-1))
                                /* ERROR PROTECTION INSERTED (g0340UU-g0338UU_try033912) */
                                if ErrorIn(g0340UU_try034116) {g0338UU_try033912 = g0340UU_try034116
                                } else {
                                g0340UU = ToProperty(OBJ(g0340UU_try034116))
                                g0338UU_try033912 = ToArray(g0313.Args.Id()).NthPut(1,g0340UU.Id()).ToEID()
                                }
                                /* Let-15 */} 
                              /* ERROR PROTECTION INSERTED (g0338UU_try033912-g0338UU_try033912) */
                              if !ErrorIn(g0338UU_try033912) {
                              g0338UU_try033912 = EID{CTRUE.Id(),0}
                              }
                              } else {
                              g0338UU_try033912 = EID{CFALSE.Id(),0}
                              /* If-14 */} 
                            /* Let-13 */} 
                          } else {
                          g0338UU_try033912 = EID{CFALSE.Id(),0}
                          /* If-12 */} 
                        /* ERROR PROTECTION INSERTED (g0338UU-g0336I_try033711) */
                        if ErrorIn(g0338UU_try033912) {g0336I_try033711 = g0338UU_try033912
                        } else {
                        g0338UU = ToBoolean(OBJ(g0338UU_try033912))
                        g0336I_try033711 = EID{g0338UU.Not.Id(),0}
                        }
                        /* Let-11 */} 
                      /* ERROR PROTECTION INSERTED (g0336I-void_try11) */
                      if ErrorIn(g0336I_try033711) {void_try11 = g0336I_try033711
                      } else {
                      g0336I = ToBoolean(OBJ(g0336I_try033711))
                      if (g0336I == CTRUE) /* If:11 */{ 
                         /*v = g0333UU_try03349, s =EID*/
g0333UU_try03349 = EID{CTRUE.Id(),0}
                        break
                        } else {
                        void_try11 = EID{CFALSE.Id(),0}
                        /* If-11 */} 
                      }
                      /* ERROR PROTECTION INSERTED (void_try11-g0333UU_try03349) */
                      if ErrorIn(void_try11) {g0333UU_try03349 = void_try11
                      g0333UU_try03349 = void_try11
                      break
                      } else {
                      }}
                      /* loop-10 */} 
                    /* For-9 */} 
                  /* ERROR PROTECTION INSERTED (g0333UU-g0331UU_try03328) */
                  if ErrorIn(g0333UU_try03349) {g0331UU_try03328 = g0333UU_try03349
                  } else {
                  g0333UU = ANY(g0333UU_try03349)
                  g0331UU_try03328 = EID{Core.F_not_any(g0333UU).Id(),0}
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (g0331UU-g0329I_try03307) */
                if ErrorIn(g0331UU_try03328) {g0329I_try03307 = g0331UU_try03328
                } else {
                g0331UU = ToBoolean(OBJ(g0331UU_try03328))
                g0329I_try03307 = EID{g0331UU.Not.Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (g0329I-Result) */
              if ErrorIn(g0329I_try03307) {Result = g0329I_try03307
              } else {
              g0329I = ToBoolean(OBJ(g0329I_try03307))
              if (g0329I == CTRUE) /* If:7 */{ 
                /* Let:8 */{ 
                  var l2 *ClaireList   = ToClass(x).Params
                  /* noccur = 3 */
                  /* Let:9 */{ 
                    var n int 
                    /* noccur = 2 */
                    if (l2.Isa.IsIn(C_list) == CTRUE) /* If:10 */{ 
                      n = l2.Length()
                      } else {
                      n = 0
                      /* If-10 */} 
                    if (Equal(ANY(Core.F_CALL(C_length,ARGS(l.ToEID()))),MakeInteger(n).Id()) == CTRUE) /* If:10 */{ 
                      var l_try034211 EID 
                      /* Let:11 */{ 
                        var i_bag *ClaireList   = ToType(CEMPTY.Id()).EmptyList()
                        /* noccur = 2 */
                        /* Let:12 */{ 
                          var i int  = 1
                          /* noccur = 5 */
                          /* Let:13 */{ 
                            var g0316 int  = n
                            /* noccur = 1 */
                            l_try034211= EID{CFALSE.Id(),0}
                            for (i <= g0316) /* while:14 */{ 
                              var void_try15 EID 
                              _ = void_try15
                              { 
                              /* Let:15 */{ 
                                var g0343UU *Language.Call  
                                /* noccur = 1 */
                                var g0343UU_try034416 EID 
                                /* Let:16 */{ 
                                  var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                  /* noccur = 5 */
                                  _CL_obj.Selector = ToProperty(C__equal.Id())
                                  /* update:17 */{ 
                                    var va_arg1 *Language.Call  
                                    var va_arg2 *ClaireList  
                                    va_arg1 = _CL_obj
                                    var va_arg2_try034518 EID 
                                    /* Construct:18 */{ 
                                      var v_bag_arg *ClaireAny  
                                      va_arg2_try034518= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                                      ToList(OBJ(va_arg2_try034518)).AddFast(l2.At(i-1))
                                      var v_bag_arg_try034619 EID 
                                      v_bag_arg_try034619 = Core.F_CALL(C_nth,ARGS(l.ToEID(),EID{C__INT,IVAL(i)}))
                                      /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try034518) */
                                      if ErrorIn(v_bag_arg_try034619) {va_arg2_try034518 = v_bag_arg_try034619
                                      } else {
                                      v_bag_arg = ANY(v_bag_arg_try034619)
                                      ToList(OBJ(va_arg2_try034518)).AddFast(v_bag_arg)}
                                      /* Construct-18 */} 
                                    /* ERROR PROTECTION INSERTED (va_arg2-g0343UU_try034416) */
                                    if ErrorIn(va_arg2_try034518) {g0343UU_try034416 = va_arg2_try034518
                                    } else {
                                    va_arg2 = ToList(OBJ(va_arg2_try034518))
                                    /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                                    va_arg1.Args = va_arg2
                                    g0343UU_try034416 = EID{va_arg2.Id(),0}
                                    }
                                    /* update-17 */} 
                                  /* ERROR PROTECTION INSERTED (g0343UU_try034416-g0343UU_try034416) */
                                  if !ErrorIn(g0343UU_try034416) {
                                  g0343UU_try034416 = EID{_CL_obj.Id(),0}
                                  }
                                  /* Let-16 */} 
                                /* ERROR PROTECTION INSERTED (g0343UU-void_try15) */
                                if ErrorIn(g0343UU_try034416) {void_try15 = g0343UU_try034416
                                } else {
                                g0343UU = Language.To_Call(OBJ(g0343UU_try034416))
                                void_try15 = EID{i_bag.AddFast(g0343UU.Id()).Id(),0}
                                }
                                /* Let-15 */} 
                              /* ERROR PROTECTION INSERTED (void_try15-void_try15) */
                              if ErrorIn(void_try15) {l_try034211 = void_try15
                              break
                              } else {
                              i = (i+1)
                              }
                              /* while-14 */} 
                            }
                            /* Let-13 */} 
                          /* Let-12 */} 
                        /* ERROR PROTECTION INSERTED (l_try034211-l_try034211) */
                        if !ErrorIn(l_try034211) {
                        l_try034211 = EID{i_bag.Id(),0}
                        }
                        /* Let-11 */} 
                      /* ERROR PROTECTION INSERTED (l-Result) */
                      if ErrorIn(l_try034211) {Result = l_try034211
                      } else {
                      l = ANY(l_try034211)
                      Result = l.ToEID()
                      }
                      } else {
                      Result = F_Serror_string(MakeString("[174] Wrong instantiation list ~S(~S..."),MakeConstantList(x,MakeConstantList(l).Id()))
                      /* If-10 */} 
                    /* Let-9 */} 
                  /* Let-8 */} 
                } else {
                Result = EID{CFALSE.Id(),0}
                /* If-7 */} 
              }
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              /* Let:7 */{ 
                var _CL_obj *Language.Definition   = Language.To_Definition(new(Language.Definition).Is(Language.C_Definition))
                /* noccur = 5 */
                _CL_obj.Arg = ToClass(x)
                _CL_obj.Args = ToList(l)
                Result = EID{_CL_obj.Id(),0}
                /* Let-7 */} 
              }
              } else {
              var g0347I *ClaireBoolean  
              /* or:7 */{ 
                var v_or7 *ClaireBoolean  
                
                v_or7 = x.Isa.IsIn(C_Variable)
                if (v_or7 == CTRUE) {g0347I = CTRUE
                } else /* or:8 */{ 
                  /* Let:9 */{ 
                    var g0348UU *ClaireObject  
                    /* noccur = 1 */
                    if (x.Isa.IsIn(Core.C_global_variable) == CTRUE) /* If:10 */{ 
                      /* Let:11 */{ 
                        var g0317 *Core.GlobalVariable   = Core.ToGlobalVariable(x)
                        /* noccur = 1 */
                        g0348UU = ToObject(g0317.Range.Id())
                        /* Let-11 */} 
                      } else {
                      g0348UU = ToObject(CFALSE.Id())
                      /* If-10 */} 
                    v_or7 = F_boolean_I_any(g0348UU.Id())
                    /* Let-9 */} 
                  if (v_or7 == CTRUE) {g0347I = CTRUE
                  } else /* or:9 */{ 
                    g0347I = CFALSE/* org-9 */} 
                  /* org-8 */} 
                /* or-7 */} 
              if (g0347I == CTRUE) /* If:7 */{ 
                /* Let:8 */{ 
                  var g0349UU *ClaireList  
                  /* noccur = 1 */
                  /* Let:9 */{ 
                    var g0350UU *ClaireAny  
                    /* noccur = 1 */
                    if (F_boolean_I_any(l) == CTRUE) /* If:10 */{ 
                      g0350UU = l
                      } else {
                      g0350UU = MakeConstantList(ClEnv.Id()).Id()
                      /* If-10 */} 
                    g0349UU = F_cons_any(x,ToList(g0350UU))
                    /* Let-9 */} 
                  Result = F_Call_I_property(Core.C_call,g0349UU)
                  /* Let-8 */} 
                } else {
                /* Let:8 */{ 
                  var p *ClaireProperty  
                  /* noccur = 2 */
                  var p_try03519 EID 
                  p_try03519 = Language.F_make_a_property_any(x)
                  /* ERROR PROTECTION INSERTED (p-Result) */
                  if ErrorIn(p_try03519) {Result = p_try03519
                  } else {
                  p = ToProperty(OBJ(p_try03519))
                  /* Let:9 */{ 
                    var l2 *ClaireAny  
                    /* noccur = 2 */
                    if (F_boolean_I_any(l) == CTRUE) /* If:10 */{ 
                      l2 = l
                      } else {
                      l2 = MakeConstantList(ClEnv.Id()).Id()
                      /* If-10 */} 
                    if (t != CNULL) /* If:10 */{ 
                      /* Let:11 */{ 
                        var _CL_obj *Language.Super   = Language.To_Super(new(Language.Super).Is(Language.C_Super))
                        /* noccur = 7 */
                        _CL_obj.Selector = p
                        _CL_obj.CastTo = ToType(t)
                        _CL_obj.Args = ToList(l2)
                        Result = EID{_CL_obj.Id(),0}
                        /* Let-11 */} 
                      } else {
                      Result = F_Call_I_property(p,ToList(l2))
                      /* If-10 */} 
                    /* Let-9 */} 
                  }
                  /* Let-8 */} 
                /* If-7 */} 
              /* If-6 */} 
            /* If-5 */} 
          /* If-4 */} 
        /* If-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: readcall @ meta_reader (throw: true) 
func E_readcall_meta_reader (r EID,x EID,t EID) EID { 
    return /*(sm for readcall @ meta_reader= EID)*/ ToMetaReader(OBJ(r)).Readcall(ANY(x),ANY(t) )} 
  
// **********************************************************************
// *   Part 4: read definitions                                         *
// **********************************************************************
// reads a definition (CLAIRE2 syntax)   - x and y are two expressions that have been read
//
/* {1} OPT.The go function for: nextdefinition(r:meta_reader,x:any,y:any,old?:boolean) [] */
func (r *MetaReader ) Nextdefinition (x *ClaireAny ,y *ClaireAny ,old_ask *ClaireBoolean ) EID { 
    var Result EID 
    r.LastArrow = CFALSE
    if (Equal(y,C_triangle.Value) == CTRUE) /* If:2 */{ 
      Result = r.Cnext().NextDefclass(x,old_ask)
      /* If!2 */}  else if (y == C_L_.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var table_ask *ClaireBoolean  
        /* noccur = 2 */
        if (x.Isa.IsIn(Language.C_Call) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var g0352 *Language.Call   = Language.To_Call(x)
            /* noccur = 3 */
            table_ask = MakeBoolean((g0352.Selector.Id() == C_nth.Id()) && ((g0352.Args.At(1-1).Isa.IsIn(C_unbound_symbol) == CTRUE) || 
                (C_table.Id() == g0352.Args.At(1-1).Isa.Id())))
            /* Let-5 */} 
          } else {
          table_ask = CFALSE
          /* If-4 */} 
        /* Let:4 */{ 
          var z *ClaireAny  
          /* noccur = 2 */
          var z_try03555 EID 
          z_try03555 = r.Nexte()
          /* ERROR PROTECTION INSERTED (z-Result) */
          if ErrorIn(z_try03555) {Result = z_try03555
          } else {
          z = ANY(z_try03555)
          /* Let:5 */{ 
            var w *ClaireAny  
            /* noccur = 5 */
            var w_try03566 EID 
            w_try03566 = r.Nexte()
            /* ERROR PROTECTION INSERTED (w-Result) */
            if ErrorIn(w_try03566) {Result = w_try03566
            } else {
            w = ANY(w_try03566)
            var g0357I *ClaireBoolean  
            if (table_ask == CTRUE) /* If:6 */{ 
              g0357I = Equal(w,C_L__equal.Id())
              } else {
              g0357I = MakeBoolean((Equal(w,C_arrow.Value) == CTRUE) || (w == C__equal_sup.Id()))
              /* If-6 */} 
            if (g0357I == CTRUE) /* If:6 */{ 
              Result = EID{CNIL.Id(),0}
              } else {
              Result = F_Serror_string(MakeString("[149] wrong keyword (~S) after ~S"),MakeConstantList(w,z))
              /* If-6 */} 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            Result = r.Nextmethod(x,
              z,
              table_ask,
              old_ask,
              Equal(w,C__equal_sup.Id()))
            }
            }
            /* Let-5 */} 
          }
          /* Let-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (y == C_L_L_.Id()) /* If:2 */{ 
      if (x.Isa.IsIn(Language.C_Call) == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var g0353 *Language.Call   = Language.To_Call(x)
          /* noccur = 2 */
          /* Let:5 */{ 
            var ru *ClaireAny  
            /* noccur = 0 */
            _ = ru
            var ru_try03586 EID 
            ru_try03586 = r.Nexte()
            /* ERROR PROTECTION INSERTED (ru-Result) */
            if ErrorIn(ru_try03586) {Result = ru_try03586
            } else {
            ru = ANY(ru_try03586)
            /* Let:6 */{ 
              var z *ClaireAny  
              /* noccur = 1 */
              var z_try03597 EID 
              z_try03597 = r.Nexts(C__equal_sup)
              /* ERROR PROTECTION INSERTED (z-Result) */
              if ErrorIn(z_try03597) {Result = z_try03597
              } else {
              z = ANY(z_try03597)
              /* Let:7 */{ 
                var _CL_obj *Language.Defrule   = Language.To_Defrule(new(Language.Defrule).Is(Language.C_Defrule))
                /* noccur = 9 */
                _CL_obj.Ident = g0353.Selector.Name
                _CL_obj.Args = g0353.Args
                _CL_obj.Arg = z
                /* update:8 */{ 
                  var va_arg1 *Language.Defrule  
                  var va_arg2 *ClaireAny  
                  va_arg1 = _CL_obj
                  var va_arg2_try03609 EID 
                  if (r.Firstc() == 41) /* If:9 */{ 
                    r.Next()
                    va_arg2_try03609 = EID{CNIL.Id(),0}
                    } else {
                    /* Let:10 */{ 
                      var g0361UU *ClaireAny  
                      /* noccur = 1 */
                      var g0361UU_try036211 EID 
                      g0361UU_try036211 = r.Nexte()
                      /* ERROR PROTECTION INSERTED (g0361UU-va_arg2_try03609) */
                      if ErrorIn(g0361UU_try036211) {va_arg2_try03609 = g0361UU_try036211
                      } else {
                      g0361UU = ANY(g0361UU_try036211)
                      va_arg2_try03609 = r.Readblock(g0361UU,41)
                      }
                      /* Let-10 */} 
                    /* If-9 */} 
                  /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                  if ErrorIn(va_arg2_try03609) {Result = va_arg2_try03609
                  } else {
                  va_arg2 = ANY(va_arg2_try03609)
                  /* ---------- now we compile update body(va_arg1) := va_arg2 ------- */
                  va_arg1.Body = va_arg2
                  Result = va_arg2.ToEID()
                  }
                  /* update-8 */} 
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                Result = EID{_CL_obj.Id(),0}
                }
                /* Let-7 */} 
              }
              /* Let-6 */} 
            }
            /* Let-5 */} 
          /* Let-4 */} 
        } else {
        Result = r.Nextinst(x)
        /* If-3 */} 
      /* If!2 */}  else if ((Equal(y,C_arrow.Value) == CTRUE) || 
        (y == C__equal_sup.Id())) /* If:2 */{ 
      r.LastArrow = Equal(y,C__equal_sup.Id())
      
      Result = r.Nextmethod(x,
        C_void.Id(),
        CFALSE,
        old_ask,
        Equal(y,C__equal_sup.Id()))
      /* If!2 */}  else if ((y == C_L__equal.Id()) && 
        (x.Isa.IsIn(Language.C_Vardef) == CTRUE)) /* If:2 */{ 
      /* Let:3 */{ 
        var _CL_obj *Language.Defobj   = Language.To_Defobj(new(Language.Defobj).Is(Language.C_Defobj))
        /* noccur = 7 */
        _CL_obj.Ident = ToSymbol(OBJ(Core.F_CALL(C_mClaire_pname,ARGS(x.ToEID()))))
        _CL_obj.Arg = Core.C_global_variable
        /* update:4 */{ 
          var va_arg1 *Language.Definition  
          var va_arg2 *ClaireList  
          va_arg1 = Language.To_Definition(_CL_obj.Id())
          var va_arg2_try03635 EID 
          /* Construct:5 */{ 
            var v_bag_arg *ClaireAny  
            va_arg2_try03635= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
            var v_bag_arg_try03646 EID 
            /* Let:6 */{ 
              var g0365UU *ClaireList  
              /* noccur = 1 */
              var g0365UU_try03667 EID 
              /* Construct:7 */{ 
                var v_bag_arg *ClaireAny  
                g0365UU_try03667= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                ToList(OBJ(g0365UU_try03667)).AddFast(C_range.Id())
                var v_bag_arg_try03678 EID 
                v_bag_arg_try03678 = Language.F_extract_type_any(ANY(Core.F_CALL(C_range,ARGS(x.ToEID()))))
                /* ERROR PROTECTION INSERTED (v_bag_arg-g0365UU_try03667) */
                if ErrorIn(v_bag_arg_try03678) {g0365UU_try03667 = v_bag_arg_try03678
                } else {
                v_bag_arg = ANY(v_bag_arg_try03678)
                ToList(OBJ(g0365UU_try03667)).AddFast(v_bag_arg)}
                /* Construct-7 */} 
              /* ERROR PROTECTION INSERTED (g0365UU-v_bag_arg_try03646) */
              if ErrorIn(g0365UU_try03667) {v_bag_arg_try03646 = g0365UU_try03667
              } else {
              g0365UU = ToList(OBJ(g0365UU_try03667))
              v_bag_arg_try03646 = F_Call_I_property(ToProperty(C__equal.Id()),g0365UU)
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try03635) */
            if ErrorIn(v_bag_arg_try03646) {va_arg2_try03635 = v_bag_arg_try03646
            } else {
            v_bag_arg = ANY(v_bag_arg_try03646)
            ToList(OBJ(va_arg2_try03635)).AddFast(v_bag_arg)
            var v_bag_arg_try03686 EID 
            /* Let:6 */{ 
              var g0369UU *ClaireList  
              /* noccur = 1 */
              var g0369UU_try03707 EID 
              /* Construct:7 */{ 
                var v_bag_arg *ClaireAny  
                g0369UU_try03707= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                ToList(OBJ(g0369UU_try03707)).AddFast(C_value.Id())
                var v_bag_arg_try03718 EID 
                v_bag_arg_try03718 = r.Nexte()
                /* ERROR PROTECTION INSERTED (v_bag_arg-g0369UU_try03707) */
                if ErrorIn(v_bag_arg_try03718) {g0369UU_try03707 = v_bag_arg_try03718
                } else {
                v_bag_arg = ANY(v_bag_arg_try03718)
                ToList(OBJ(g0369UU_try03707)).AddFast(v_bag_arg)}
                /* Construct-7 */} 
              /* ERROR PROTECTION INSERTED (g0369UU-v_bag_arg_try03686) */
              if ErrorIn(g0369UU_try03707) {v_bag_arg_try03686 = g0369UU_try03707
              } else {
              g0369UU = ToList(OBJ(g0369UU_try03707))
              v_bag_arg_try03686 = F_Call_I_property(ToProperty(C__equal.Id()),g0369UU)
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try03635) */
            if ErrorIn(v_bag_arg_try03686) {va_arg2_try03635 = v_bag_arg_try03686
            } else {
            v_bag_arg = ANY(v_bag_arg_try03686)
            ToList(OBJ(va_arg2_try03635)).AddFast(v_bag_arg)}}
            /* Construct-5 */} 
          /* ERROR PROTECTION INSERTED (va_arg2-Result) */
          if ErrorIn(va_arg2_try03635) {Result = va_arg2_try03635
          } else {
          va_arg2 = ToList(OBJ(va_arg2_try03635))
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
        var _CL_obj *Language.Do   = Language.To_Do(new(Language.Do).Is(Language.C_Do))
        /* noccur = 3 */
        _CL_obj.Args = MakeList(ToType(C_any.Id()),x,y)
        Result = EID{_CL_obj.Id(),0}
        /* Let-3 */} 
      /* If-2 */} 
    return Result} 
  
// The EID go function for: nextdefinition @ meta_reader (throw: true) 
func E_nextdefinition_meta_reader (r EID,x EID,y EID,old_ask EID) EID { 
    return /*(sm for nextdefinition @ meta_reader= EID)*/ ToMetaReader(OBJ(r)).Nextdefinition(ANY(x),
      ANY(y),
      ToBoolean(OBJ(old_ask)) )} 
  
/* {1} OPT.The go function for: nextmethod(r:meta_reader,x:any,y:any,table?:boolean,old?:boolean,inl?:boolean) [] */
func (r *MetaReader ) Nextmethod (x *ClaireAny ,y *ClaireAny ,table_ask *ClaireBoolean ,old_ask *ClaireBoolean ,inl_ask *ClaireBoolean ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var n int  = r.Skipc()
      /* noccur = 1 */
      /* Let:3 */{ 
        var z *ClaireAny  
        /* noccur = 2 */
        var z_try03724 EID 
        if (old_ask == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var g0373UU *ClaireAny  
            /* noccur = 1 */
            var g0373UU_try03746 EID 
            g0373UU_try03746 = r.Nexte()
            /* ERROR PROTECTION INSERTED (g0373UU-z_try03724) */
            if ErrorIn(g0373UU_try03746) {z_try03724 = g0373UU_try03746
            } else {
            g0373UU = ANY(g0373UU_try03746)
            z_try03724 = r.Readblock(g0373UU,93)
            }
            /* Let-5 */} 
          /* If!4 */}  else if (n == 40) /* If:4 */{ 
          if (r.Toplevel == CTRUE) /* If:5 */{ 
            z_try03724 = r.Nexts(C_none)
            } else {
            /* Let:6 */{ 
              var g0375UU *ClaireAny  
              /* noccur = 1 */
              var g0375UU_try03767 EID 
              g0375UU_try03767 = r.Cnext().Nexte()
              /* ERROR PROTECTION INSERTED (g0375UU-z_try03724) */
              if ErrorIn(g0375UU_try03767) {z_try03724 = g0375UU_try03767
              } else {
              g0375UU = ANY(g0375UU_try03767)
              z_try03724 = r.Readblock(g0375UU,41)
              }
              /* Let-6 */} 
            /* If-5 */} 
          } else {
          z_try03724 = r.Nexte()
          /* If-4 */} 
        /* ERROR PROTECTION INSERTED (z-Result) */
        if ErrorIn(z_try03724) {Result = z_try03724
        } else {
        z = ANY(z_try03724)
        /* Let:4 */{ 
          var rs *Language.Defmethod  
          /* noccur = 2 */
          var rs_try03775 EID 
          /* Let:5 */{ 
            var _CL_obj *Language.Defmethod   = Language.To_Defmethod(new(Language.Defmethod).Is(Language.C_Defmethod))
            /* noccur = 8 */
            _CL_obj.Arg = Language.To_Call(x)
            _CL_obj.SetArg = y
            /* update:6 */{ 
              var va_arg1 *Language.Defmethod  
              var va_arg2 *ClaireAny  
              va_arg1 = _CL_obj
              var va_arg2_try03787 EID 
              if (z == C_let.Id()) /* If:7 */{ 
                va_arg2_try03787 = r.Readlet(C_None)
                } else {
                va_arg2_try03787 = z.ToEID()
                /* If-7 */} 
              /* ERROR PROTECTION INSERTED (va_arg2-rs_try03775) */
              if ErrorIn(va_arg2_try03787) {rs_try03775 = va_arg2_try03787
              } else {
              va_arg2 = ANY(va_arg2_try03787)
              /* ---------- now we compile update body(va_arg1) := va_arg2 ------- */
              va_arg1.Body = va_arg2
              rs_try03775 = va_arg2.ToEID()
              }
              /* update-6 */} 
            /* ERROR PROTECTION INSERTED (rs_try03775-rs_try03775) */
            if !ErrorIn(rs_try03775) {
            _CL_obj.Inline_ask = inl_ask
            rs_try03775 = EID{_CL_obj.Id(),0}
            }
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (rs-Result) */
          if ErrorIn(rs_try03775) {Result = rs_try03775
          } else {
          rs = Language.To_Defmethod(OBJ(rs_try03775))
          if (table_ask == CTRUE) /* If:5 */{ 
            rs.Isa = Language.C_Defarray
            /* If-5 */} 
          Result = EID{rs.Id(),0}
          }
          /* Let-4 */} 
        }
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: nextmethod @ meta_reader (throw: true) 
func E_nextmethod_meta_reader (r EID,x EID,y EID,table_ask EID,old_ask EID,inl_ask EID) EID { 
    return /*(sm for nextmethod @ meta_reader= EID)*/ ToMetaReader(OBJ(r)).Nextmethod(ANY(x),
      ANY(y),
      ToBoolean(OBJ(table_ask)),
      ToBoolean(OBJ(old_ask)),
      ToBoolean(OBJ(inl_ask)) )} 
  
// reads an instantiation
//
/* {1} OPT.The go function for: nextinst(r:meta_reader,x:any) [] */
func (r *MetaReader ) Nextinst (x *ClaireAny ) EID { 
    var Result EID 
    if (x.Isa.IsIn(C_Variable) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0379 *ClaireVariable   = To_Variable(x)
        /* noccur = 2 */
        /* Let:4 */{ 
          var _CL_obj *Language.Defobj   = Language.To_Defobj(new(Language.Defobj).Is(Language.C_Defobj))
          /* noccur = 19 */
          _CL_obj.Ident = g0379.Pname
          _CL_obj.Arg = Core.C_global_variable
          /* update:5 */{ 
            var va_arg1 *Language.Definition  
            var va_arg2 *ClaireList  
            va_arg1 = Language.To_Definition(_CL_obj.Id())
            var va_arg2_try03846 EID 
            /* Construct:6 */{ 
              var v_bag_arg *ClaireAny  
              va_arg2_try03846= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
              var v_bag_arg_try03857 EID 
              /* Let:7 */{ 
                var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                /* noccur = 5 */
                _CL_obj.Selector = ToProperty(C__equal.Id())
                /* update:8 */{ 
                  var va_arg1 *Language.Call  
                  var va_arg2 *ClaireList  
                  va_arg1 = _CL_obj
                  var va_arg2_try03869 EID 
                  /* Construct:9 */{ 
                    var v_bag_arg *ClaireAny  
                    va_arg2_try03869= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                    ToList(OBJ(va_arg2_try03869)).AddFast(C_range.Id())
                    var v_bag_arg_try038710 EID 
                    v_bag_arg_try038710 = Language.F_extract_type_any(g0379.Range.Id())
                    /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try03869) */
                    if ErrorIn(v_bag_arg_try038710) {va_arg2_try03869 = v_bag_arg_try038710
                    } else {
                    v_bag_arg = ANY(v_bag_arg_try038710)
                    ToList(OBJ(va_arg2_try03869)).AddFast(v_bag_arg)}
                    /* Construct-9 */} 
                  /* ERROR PROTECTION INSERTED (va_arg2-v_bag_arg_try03857) */
                  if ErrorIn(va_arg2_try03869) {v_bag_arg_try03857 = va_arg2_try03869
                  } else {
                  va_arg2 = ToList(OBJ(va_arg2_try03869))
                  /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                  va_arg1.Args = va_arg2
                  v_bag_arg_try03857 = EID{va_arg2.Id(),0}
                  }
                  /* update-8 */} 
                /* ERROR PROTECTION INSERTED (v_bag_arg_try03857-v_bag_arg_try03857) */
                if !ErrorIn(v_bag_arg_try03857) {
                v_bag_arg_try03857 = EID{_CL_obj.Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try03846) */
              if ErrorIn(v_bag_arg_try03857) {va_arg2_try03846 = v_bag_arg_try03857
              } else {
              v_bag_arg = ANY(v_bag_arg_try03857)
              ToList(OBJ(va_arg2_try03846)).AddFast(v_bag_arg)
              var v_bag_arg_try03887 EID 
              /* Let:7 */{ 
                var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                /* noccur = 5 */
                _CL_obj.Selector = ToProperty(C__equal.Id())
                /* update:8 */{ 
                  var va_arg1 *Language.Call  
                  var va_arg2 *ClaireList  
                  va_arg1 = _CL_obj
                  var va_arg2_try03899 EID 
                  /* Construct:9 */{ 
                    var v_bag_arg *ClaireAny  
                    va_arg2_try03899= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                    ToList(OBJ(va_arg2_try03899)).AddFast(C_value.Id())
                    var v_bag_arg_try039010 EID 
                    v_bag_arg_try039010 = r.Nexte()
                    /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try03899) */
                    if ErrorIn(v_bag_arg_try039010) {va_arg2_try03899 = v_bag_arg_try039010
                    } else {
                    v_bag_arg = ANY(v_bag_arg_try039010)
                    ToList(OBJ(va_arg2_try03899)).AddFast(v_bag_arg)}
                    /* Construct-9 */} 
                  /* ERROR PROTECTION INSERTED (va_arg2-v_bag_arg_try03887) */
                  if ErrorIn(va_arg2_try03899) {v_bag_arg_try03887 = va_arg2_try03899
                  } else {
                  va_arg2 = ToList(OBJ(va_arg2_try03899))
                  /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                  va_arg1.Args = va_arg2
                  v_bag_arg_try03887 = EID{va_arg2.Id(),0}
                  }
                  /* update-8 */} 
                /* ERROR PROTECTION INSERTED (v_bag_arg_try03887-v_bag_arg_try03887) */
                if !ErrorIn(v_bag_arg_try03887) {
                v_bag_arg_try03887 = EID{_CL_obj.Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try03846) */
              if ErrorIn(v_bag_arg_try03887) {va_arg2_try03846 = v_bag_arg_try03887
              } else {
              v_bag_arg = ANY(v_bag_arg_try03887)
              ToList(OBJ(va_arg2_try03846)).AddFast(v_bag_arg)}}
              /* Construct-6 */} 
            /* ERROR PROTECTION INSERTED (va_arg2-Result) */
            if ErrorIn(va_arg2_try03846) {Result = va_arg2_try03846
            } else {
            va_arg2 = ToList(OBJ(va_arg2_try03846))
            /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
            va_arg1.Args = va_arg2
            Result = EID{va_arg2.Id(),0}
            }
            /* update-5 */} 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          Result = EID{_CL_obj.Id(),0}
          }
          /* Let-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (x.Isa.IsIn(Language.C_Call) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0380 *Language.Call   = Language.To_Call(x)
        /* noccur = 2 */
        /* Let:4 */{ 
          var ru *ClaireAny  
          /* noccur = 0 */
          _ = ru
          var ru_try03915 EID 
          ru_try03915 = r.Nexte()
          /* ERROR PROTECTION INSERTED (ru-Result) */
          if ErrorIn(ru_try03915) {Result = ru_try03915
          } else {
          ru = ANY(ru_try03915)
          /* Let:5 */{ 
            var z *ClaireAny  
            /* noccur = 1 */
            var z_try03926 EID 
            z_try03926 = r.Nexts(C__equal_sup)
            /* ERROR PROTECTION INSERTED (z-Result) */
            if ErrorIn(z_try03926) {Result = z_try03926
            } else {
            z = ANY(z_try03926)
            /* Let:6 */{ 
              var _CL_obj *Language.Defrule   = Language.To_Defrule(new(Language.Defrule).Is(Language.C_Defrule))
              /* noccur = 9 */
              _CL_obj.Ident = g0380.Selector.Name
              _CL_obj.Args = g0380.Args
              _CL_obj.Arg = z
              /* update:7 */{ 
                var va_arg1 *Language.Defrule  
                var va_arg2 *ClaireAny  
                va_arg1 = _CL_obj
                var va_arg2_try03938 EID 
                if (r.Firstc() == 41) /* If:8 */{ 
                  r.Next()
                  va_arg2_try03938 = EID{CNIL.Id(),0}
                  } else {
                  /* Let:9 */{ 
                    var g0394UU *ClaireAny  
                    /* noccur = 1 */
                    var g0394UU_try039510 EID 
                    g0394UU_try039510 = r.Nexte()
                    /* ERROR PROTECTION INSERTED (g0394UU-va_arg2_try03938) */
                    if ErrorIn(g0394UU_try039510) {va_arg2_try03938 = g0394UU_try039510
                    } else {
                    g0394UU = ANY(g0394UU_try039510)
                    va_arg2_try03938 = r.Readblock(g0394UU,41)
                    }
                    /* Let-9 */} 
                  /* If-8 */} 
                /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                if ErrorIn(va_arg2_try03938) {Result = va_arg2_try03938
                } else {
                va_arg2 = ANY(va_arg2_try03938)
                /* ---------- now we compile update body(va_arg1) := va_arg2 ------- */
                va_arg1.Body = va_arg2
                Result = va_arg2.ToEID()
                }
                /* update-7 */} 
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              Result = EID{_CL_obj.Id(),0}
              }
              /* Let-6 */} 
            }
            /* Let-5 */} 
          }
          /* Let-4 */} 
        /* Let-3 */} 
      } else {
      /* Let:3 */{ 
        var y *ClaireAny  
        /* noccur = 7 */
        var y_try03964 EID 
        y_try03964 = r.Nexte()
        /* ERROR PROTECTION INSERTED (y-Result) */
        if ErrorIn(y_try03964) {Result = y_try03964
        } else {
        y = ANY(y_try03964)
        var g0397I *ClaireBoolean  
        if (x.Isa.IsIn(Core.C_global_variable) == CTRUE) /* If:4 */{ 
          g0397I = Core.F_unknown_ask_any(y)
          } else {
          g0397I = CFALSE
          /* If-4 */} 
        if (g0397I == CTRUE) /* If:4 */{ 
          Result = y.ToEID()
          } else {
          var g0398I *ClaireBoolean  
          if (y.Isa.IsIn(Language.C_Definition) == CTRUE) /* If:5 */{ 
            /* Let:6 */{ 
              var g0383 *Language.Definition   = Language.To_Definition(y)
              /* noccur = 1 */
              g0398I = g0383.Arg.IsIn(C_thing)
              /* Let-6 */} 
            } else {
            g0398I = CFALSE
            /* If-5 */} 
          if (g0398I == CTRUE) /* If:5 */{ 
            /* Let:6 */{ 
              var _CL_obj *Language.Defobj   = Language.To_Defobj(new(Language.Defobj).Is(Language.C_Defobj))
              /* noccur = 7 */
              /* update:7 */{ 
                var va_arg1 *Language.Defobj  
                var va_arg2 *ClaireSymbol  
                va_arg1 = _CL_obj
                var va_arg2_try03998 EID 
                va_arg2_try03998 = Language.F_extract_symbol_any(x)
                /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                if ErrorIn(va_arg2_try03998) {Result = va_arg2_try03998
                } else {
                va_arg2 = ToSymbol(OBJ(va_arg2_try03998))
                /* ---------- now we compile update iClaire/ident(va_arg1) := va_arg2 ------- */
                va_arg1.Ident = va_arg2
                Result = EID{va_arg2.Id(),0}
                }
                /* update-7 */} 
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              _CL_obj.Arg = ToClass(OBJ(Core.F_CALL(C_arg,ARGS(y.ToEID()))))
              _CL_obj.Args = ToList(OBJ(Core.F_CALL(C_args,ARGS(y.ToEID()))))
              Result = EID{_CL_obj.Id(),0}
              }
              /* Let-6 */} 
            } else {
            /* Let:6 */{ 
              var _CL_obj *Language.Defobj   = Language.To_Defobj(new(Language.Defobj).Is(Language.C_Defobj))
              /* noccur = 19 */
              /* update:7 */{ 
                var va_arg1 *Language.Defobj  
                var va_arg2 *ClaireSymbol  
                va_arg1 = _CL_obj
                var va_arg2_try04008 EID 
                va_arg2_try04008 = Language.F_extract_symbol_any(x)
                /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                if ErrorIn(va_arg2_try04008) {Result = va_arg2_try04008
                } else {
                va_arg2 = ToSymbol(OBJ(va_arg2_try04008))
                /* ---------- now we compile update iClaire/ident(va_arg1) := va_arg2 ------- */
                va_arg1.Ident = va_arg2
                Result = EID{va_arg2.Id(),0}
                }
                /* update-7 */} 
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              _CL_obj.Arg = Core.C_global_variable
              /* update:7 */{ 
                var va_arg1 *Language.Definition  
                var va_arg2 *ClaireList  
                va_arg1 = Language.To_Definition(_CL_obj.Id())
                /* Construct:8 */{ 
                  var v_bag_arg *ClaireAny  
                  va_arg2= ToType(CEMPTY.Id()).EmptyList()
                  /* Let:9 */{ 
                    var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                    /* noccur = 5 */
                    _CL_obj.Selector = ToProperty(C__equal.Id())
                    _CL_obj.Args = MakeConstantList(C_range.Id(),CEMPTY.Id())
                    v_bag_arg = _CL_obj.Id()
                    /* Let-9 */} 
                  va_arg2.AddFast(v_bag_arg)
                  /* Let:9 */{ 
                    var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                    /* noccur = 5 */
                    _CL_obj.Selector = ToProperty(C__equal.Id())
                    _CL_obj.Args = MakeConstantList(C_value.Id(),y)
                    v_bag_arg = _CL_obj.Id()
                    /* Let-9 */} 
                  va_arg2.AddFast(v_bag_arg)/* Construct-8 */} 
                /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                va_arg1.Args = va_arg2
                /* update-7 */} 
              Result = EID{_CL_obj.Id(),0}
              }
              /* Let-6 */} 
            /* If-5 */} 
          /* If-4 */} 
        }
        /* Let-3 */} 
      /* If-2 */} 
    return Result} 
  
// The EID go function for: nextinst @ meta_reader (throw: true) 
func E_nextinst_meta_reader (r EID,x EID) EID { 
    return /*(sm for nextinst @ meta_reader= EID)*/ ToMetaReader(OBJ(r)).Nextinst(ANY(x) )} 
  
// reads a class Definition of the form C(p:t | p:t = v *)
// new in v2.5
/* {1} OPT.The go function for: nextDefclass(r:meta_reader,x:any,old?:boolean) [] */
func (r *MetaReader ) NextDefclass (x *ClaireAny ,old_ask *ClaireBoolean ) EID { 
    var Result EID 
    r.Skipc()
    /* Let:2 */{ 
      var c *ClaireAny  
      /* noccur = 3 */
      var c_try04053 EID 
      /* Let:3 */{ 
        var g0406UU *ClaireAny  
        /* noccur = 1 */
        var g0406UU_try04074 EID 
        g0406UU_try04074 = r.Fromp.ReadIdent()
        /* ERROR PROTECTION INSERTED (g0406UU-c_try04053) */
        if ErrorIn(g0406UU_try04074) {c_try04053 = g0406UU_try04074
        } else {
        g0406UU = ANY(g0406UU_try04074)
        c_try04053 = F_verify_any(C_class.Id(),g0406UU,Language.C_Defclass.Id())
        }
        /* Let-3 */} 
      /* ERROR PROTECTION INSERTED (c-Result) */
      if ErrorIn(c_try04053) {Result = c_try04053
      } else {
      c = ANY(c_try04053)
      /* Let:3 */{ 
        var y *Language.Defclass  
        /* noccur = 6 */
        var y_try04084 EID 
        if (r.Firstc() != 40) /* If:4 */{ 
          /* Let:5 */{ 
            var _CL_obj *Language.Defclass   = Language.To_Defclass(new(Language.Defclass).Is(Language.C_Defclass))
            /* noccur = 7 */
            _CL_obj.Arg = ToClass(c)
            _CL_obj.Args = CNIL
            _CL_obj.Forward_ask = CTRUE
            y_try04084 = EID{_CL_obj.Id(),0}
            /* Let-5 */} 
          } else {
          /* Let:5 */{ 
            var l *ClaireAny  
            /* noccur = 3 */
            var l_try04096 EID 
            l_try04096 = r.Cnext().Nextseq(41)
            /* ERROR PROTECTION INSERTED (l-y_try04084) */
            if ErrorIn(l_try04096) {y_try04084 = l_try04096
            } else {
            l = ANY(l_try04096)
            /* For:6 */{ 
              var y1 *ClaireAny  
              _ = y1
              y_try04084= EID{CFALSE.Id(),0}
              var y1_support *ClaireList  
              var y1_support_try04107 EID 
              y1_support_try04107 = Core.F_enumerate_any(l)
              /* ERROR PROTECTION INSERTED (y1_support-y_try04084) */
              if ErrorIn(y1_support_try04107) {y_try04084 = y1_support_try04107
              } else {
              y1_support = ToList(OBJ(y1_support_try04107))
              y1_len := y1_support.Length()
              for i_it := 0; i_it < y1_len; i_it++ { 
                y1 = y1_support.At(i_it)
                var void_try8 EID 
                _ = void_try8
                var g0411I *ClaireBoolean  
                /* Let:8 */{ 
                  var g0412UU *ClaireBoolean  
                  /* noccur = 1 */
                  if (y1.Isa.IsIn(Language.C_Call) == CTRUE) /* If:9 */{ 
                    /* Let:10 */{ 
                      var g0401 *Language.Call   = Language.To_Call(y1)
                      /* noccur = 2 */
                      g0412UU = MakeBoolean((g0401.Selector.Id() == C__equal.Id()) && (g0401.Args.At(1-1).Isa.IsIn(Language.C_Vardef) == CTRUE))
                      /* Let-10 */} 
                    /* If!9 */}  else if (y1.Isa.IsIn(Language.C_Vardef) == CTRUE) /* If:9 */{ 
                    g0412UU = CTRUE
                    } else {
                    g0412UU = CFALSE
                    /* If-9 */} 
                  g0411I = g0412UU.Not
                  /* Let-8 */} 
                if (g0411I == CTRUE) /* If:8 */{ 
                  void_try8 = F_Serror_string(MakeString("[175] Wrong form ~S in ~S(~S)"),MakeConstantList(y1,c,l))
                  } else {
                  void_try8 = EID{CFALSE.Id(),0}
                  /* If-8 */} 
                /* ERROR PROTECTION INSERTED (void_try8-y_try04084) */
                if ErrorIn(void_try8) {y_try04084 = void_try8
                y_try04084 = void_try8
                break
                } else {
                }}
                /* loop-7 */} 
              /* For-6 */} 
            /* ERROR PROTECTION INSERTED (y_try04084-y_try04084) */
            if !ErrorIn(y_try04084) {
            /* Let:6 */{ 
              var _CL_obj *Language.Defclass   = Language.To_Defclass(new(Language.Defclass).Is(Language.C_Defclass))
              /* noccur = 7 */
              _CL_obj.Arg = ToClass(c)
              _CL_obj.Args = ToList(l)
              _CL_obj.Forward_ask = CFALSE
              y_try04084 = EID{_CL_obj.Id(),0}
              /* Let-6 */} 
            }
            }
            /* Let-5 */} 
          /* If-4 */} 
        /* ERROR PROTECTION INSERTED (y-Result) */
        if ErrorIn(y_try04084) {Result = y_try04084
        } else {
        y = Language.To_Defclass(OBJ(y_try04084))
        /* Let:4 */{ 
          var lp *ClaireList   = CNIL
          /* noccur = 3 */
          /* Let:5 */{ 
            var idt *ClaireSymbol  
            /* noccur = 1 */
            var idt_try04136 EID 
            var g0414I *ClaireBoolean  
            if (x.Isa.IsIn(Language.C_Call) == CTRUE) /* If:6 */{ 
              /* Let:7 */{ 
                var g0404 *Language.Call   = Language.To_Call(x)
                /* noccur = 1 */
                g0414I = Equal(g0404.Selector.Id(),C_nth.Id())
                /* Let-7 */} 
              } else {
              g0414I = CFALSE
              /* If-6 */} 
            if (g0414I == CTRUE) /* If:6 */{ 
              /* Let:7 */{ 
                var l *ClaireList   = Language.To_Call(x).Args
                /* noccur = 4 */
                if (C_class.Id() == l.At(1-1).Isa.Id()) /* If:8 */{ 
                  lp = ToList(l.At(2-1))
                  idt_try04136 = EID{lp.Id(),0}
                  } else {
                  var lp_try04159 EID 
                  /* Iteration:9 */{ 
                    var v_list9 *ClaireList  
                    var y2 *ClaireAny  
                    var v_local9 *ClaireAny  
                    var v_list9_try041610 EID 
                    v_list9_try041610 = l.Cdr()
                    /* ERROR PROTECTION INSERTED (v_list9-lp_try04159) */
                    if ErrorIn(v_list9_try041610) {lp_try04159 = v_list9_try041610
                    } else {
                    v_list9 = ToList(OBJ(v_list9_try041610))
                    lp_try04159 = EID{CreateList(ToType(CEMPTY.Id()),v_list9.Length()).Id(),0}
                    for CLcount := 0; CLcount < v_list9.Length(); CLcount++{ 
                      y2 = v_list9.At(CLcount)
                      var v_local9_try041711 EID 
                      v_local9_try041711 = Language.F_make_a_property_any(y2)
                      /* ERROR PROTECTION INSERTED (v_local9-lp_try04159) */
                      if ErrorIn(v_local9_try041711) {lp_try04159 = v_local9_try041711
                      lp_try04159 = v_local9_try041711
                      break
                      } else {
                      v_local9 = ANY(v_local9_try041711)
                      ToList(OBJ(lp_try04159)).PutAt(CLcount,v_local9)
                      } 
                    }}
                    /* Iteration-9 */} 
                  /* ERROR PROTECTION INSERTED (lp-idt_try04136) */
                  if ErrorIn(lp_try04159) {idt_try04136 = lp_try04159
                  } else {
                  lp = ToList(OBJ(lp_try04159))
                  idt_try04136 = EID{lp.Id(),0}
                  }
                  /* If-8 */} 
                /* ERROR PROTECTION INSERTED (idt_try04136-idt_try04136) */
                if !ErrorIn(idt_try04136) {
                idt_try04136 = Language.F_extract_symbol_any(l.At(1-1))
                }
                /* Let-7 */} 
              } else {
              idt_try04136 = Language.F_extract_symbol_any(x)
              /* If-6 */} 
            /* ERROR PROTECTION INSERTED (idt-Result) */
            if ErrorIn(idt_try04136) {Result = idt_try04136
            } else {
            idt = ToSymbol(OBJ(idt_try04136))
            if ((old_ask == CTRUE) && 
                (r.Skipc() != 93)) /* If:6 */{ 
              Result = F_Serror_string(MakeString("[176] Missing ] after ~S "),MakeConstantList(y.Id()))
              /* If!6 */}  else if (old_ask == CTRUE) /* If:6 */{ 
              r.Next()
              Result = EVOID
              } else {
              Result = EID{CFALSE.Id(),0}
              /* If-6 */} 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            y.Ident = idt
            y.Params = lp
            Result = EID{y.Id(),0}
            }
            }
            /* Let-5 */} 
          /* Let-4 */} 
        }
        /* Let-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: nextDefclass @ meta_reader (throw: true) 
func E_nextDefclass_meta_reader (r EID,x EID,old_ask EID) EID { 
    return /*(sm for nextDefclass @ meta_reader= EID)*/ ToMetaReader(OBJ(r)).NextDefclass(ANY(x),ToBoolean(OBJ(old_ask)) )} 
  
// end of file