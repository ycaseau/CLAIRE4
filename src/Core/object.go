/***** CLAIRE Compilation of file /Users/ycaseau/Dropbox/src/clairev4.0/src/meta/object.cl 
         [version 4.0.02 / safety 5] Saturday 10-30-2021 *****/

package Core
import (_ "fmt"
	. "Kernel"
)

//+-------------------------------------------------------------+
//| CLAIRE                                                      |
//| object.cl                                                   |
//| Copyright (C) 1994 - 2021 Yves Caseau. All Rights Reserved  |
//| cf. copyright info in about()                               |
//+-------------------------------------------------------------+
// ---------------------------------------------------------------------
// This file contains the definition of the objects that implement the
// core features of the microCLAIRE library: traceable & debug-able calls,
// tables, demons and exceptions
// ---------------------------------------------------------------------
// *********************************************************************
// *  Table of contents                                                *
// *   Part 1: Ask, debug & trace                                      *
// *   Part 2: Tables                                                  *
// *   Part 3: Demons & relations for the logic modules                *
// *   Part 4: Basics of Exceptions                                    *
// *********************************************************************
// release() should produce a version number
/* {1} OPT.The go function for: release(_CL_obj:void) [] */
func F_release_void () EID { 
    var Result EID 
    /* Let:2 */{ 
      var g0074UU *ClaireString  
      /* noccur = 1 */
      var g0074UU_try00753 EID 
      g0074UU_try00753 = F_string_I_float(ClEnv.Version)
      /* ERROR PROTECTION INSERTED (g0074UU-Result) */
      if ErrorIn(g0074UU_try00753) {Result = g0074UU_try00753
      } else {
      g0074UU = ToString(OBJ(g0074UU_try00753))
      Result = EID{F_append_string(MakeString("4."),g0074UU).Id(),0}
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: release @ void (throw: true) 
func E_release_void (_CL_obj EID) EID { 
    return /*(sm for release @ void= EID)*/ F_release_void( )} 
  
// the about method produces the legal warning, according to the GNU software
// recommendation
/* {1} OPT.The go function for: about(_CL_obj:void) [] */
func F_about_void () *ClaireAny  { 
    // use function body compiling 
PRINC("CLAIRE v4.")
    F_princ_float(ClEnv.Version)
    PRINC(" Copyright (C) 1994-2021 Yves Caseau. All Rights Reserved.\n")
    PRINC("use and redistribution in source code or binary forms are permitted\n")
    PRINC("resale is not permitted without the explicit agreement of Yves Caseau\n")
    PRINC("THIS SOFTWARE IS PROVIDED AS IS AND WITHOUT ANY WARRANTY, INCLUDING,\n")
    PRINC("WITHOUT LIMITATION, THE IMPLIED WARRANTIES OF MERCHANTABILTY AND FITNESS\n")
    PRINC("FOR A PARTICULAR PURPOSE\n")
    return  CTRUE.Id()
    } 
  
// The EID go function for: about @ void (throw: false) 
func E_about_void (_CL_obj EID) EID { 
    return /*(sm for about @ void= any)*/ F_about_void( ).ToEID()} 
  
// properties that are defined through compiling (eval would entail a loop)
// *********************************************************************
// *   Part 1: Ask, debug & trace                                      *
// *********************************************************************
// create the list of arguments if needed : allocate on the stack
/* {1} OPT.The go function for: mClaire/get_args(i:integer) [] */
func F_get_args_integer (i int) *ClaireList  { 
    // procedure body with s =  
var Result *ClaireList  
    /* Let:2 */{ 
      var liste *ClaireList   = ToType(C_any.Id()).EmptyList()
      /* noccur = 3 */
      for (i < ClEnv.Index) /* while:3 */{ 
        liste = liste.AddFast(ANY(ClEnv.EvalStack[i]))
        i = (i+1)
        /* while-3 */} 
      Result = liste
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: mClaire/get_args @ integer (throw: false) 
func E_get_args_integer (i EID) EID { 
    return EID{/*(sm for mClaire/get_args @ integer= list)*/ F_get_args_integer(INT(i) ).Id(),0}} 
  
// a simple method for a direct call with no argument
/* {1} OPT.The go function for: funcall(self:method,x:any) [] */
func F_funcall_method1 (self *ClaireMethod ,x *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var start int  = ClEnv.Index
      /* noccur = 1 */
      ClEnv.Push(x.ToEID())
      Result = F_execute_method(self,start,CFALSE)
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: funcall @ list<type_expression>(method, any) (throw: true) 
func E_funcall_method1 (self EID,x EID) EID { 
    return /*(sm for funcall @ list<type_expression>(method, any)= EID)*/ F_funcall_method1(ToMethod(OBJ(self)),ANY(x) )} 
  
// this is a simple method for calling directly a method with one argument
/* {1} OPT.The go function for: funcall(self:method,x:any,y:any) [] */
func F_funcall_method2 (self *ClaireMethod ,x *ClaireAny ,y *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var start int  = ClEnv.Index
      /* noccur = 1 */
      ClEnv.Push(x.ToEID())
      ClEnv.Push(y.ToEID())
      Result = F_execute_method(self,start,CFALSE)
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: funcall @ list<type_expression>(method, any, any) (throw: true) 
func E_funcall_method2 (self EID,x EID,y EID) EID { 
    return /*(sm for funcall @ list<type_expression>(method, any, any)= EID)*/ F_funcall_method2(ToMethod(OBJ(self)),ANY(x),ANY(y) )} 
  
// this is a simple method for calling directly a method with two arguments
/* {1} OPT.The go function for: funcall(self:method,x:any,y:any,z:any) [] */
func F_funcall_method3 (self *ClaireMethod ,x *ClaireAny ,y *ClaireAny ,z *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var start int  = ClEnv.Index
      /* noccur = 1 */
      ClEnv.Push(x.ToEID())
      ClEnv.Push(y.ToEID())
      ClEnv.Push(z.ToEID())
      Result = F_execute_method(self,start,CFALSE)
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: funcall @ list<type_expression>(method, any, any, any) (throw: true) 
func E_funcall_method3 (self EID,x EID,y EID,z EID) EID { 
    return /*(sm for funcall @ list<type_expression>(method, any, any, any)= EID)*/ F_funcall_method3(ToMethod(OBJ(self)),
      ANY(x),
      ANY(y),
      ANY(z) )} 
  
// how to apply a property to a list  (the function case is handled in Kernel - primitive go code)
/* {1} OPT.The go function for: call(p:property,l:listargs) [] */
func F_call_property (p *ClaireProperty ,l *ClaireList ) EID { 
    var Result EID 
    Result = F_apply_property(p,ToList(l.Id()))
    return Result} 
  
// The EID go function for: call @ property (throw: true) 
func E_call_property (p EID,l EID) EID { 
    return /*(sm for call @ property= EID)*/ F_call_property(ToProperty(OBJ(p)),ToList(OBJ(l)) )} 
  
/* {1} OPT.The go function for: apply(p:property,l:list) [] */
func F_apply_property (p *ClaireProperty ,l *ClaireList ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var start int  = ClEnv.Index
      /* noccur = 2 */
      /* For:3 */{ 
        var x *ClaireAny  
        _ = x
        var x_support *ClaireList  
        x_support = l
        x_len := x_support.Length()
        for i_it := 0; i_it < x_len; i_it++ { 
          x = x_support.At(i_it)
          ClEnv.Push(x.ToEID())
          /* loop-4 */} 
        /* For-3 */} 
      Result = F_eval_message_property(p,F_find_which_property(p,start,l.At(1-1).Isa),start,CTRUE)
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: apply @ property (throw: true) 
func E_apply_property (p EID,l EID) EID { 
    return /*(sm for apply @ property= EID)*/ F_apply_property(ToProperty(OBJ(p)),ToList(OBJ(l)) )} 
  
/* {1} OPT.The go function for: apply(m:method,l:list) [] */
func F_apply_method (m *ClaireMethod ,l *ClaireList ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var start int  = ClEnv.Index
      /* noccur = 1 */
      /* For:3 */{ 
        var x *ClaireAny  
        _ = x
        var x_support *ClaireList  
        x_support = l
        x_len := x_support.Length()
        for i_it := 0; i_it < x_len; i_it++ { 
          x = x_support.At(i_it)
          ClEnv.Push(x.ToEID())
          /* loop-4 */} 
        /* For-3 */} 
      Result = F_execute_method(m,start,CFALSE)
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: apply @ method (throw: true) 
func E_apply_method (m EID,l EID) EID { 
    return /*(sm for apply @ method= EID)*/ F_apply_method(ToMethod(OBJ(m)),ToList(OBJ(l)) )} 
  
// defined in Reader but tested in Core
// push and pop debug info on the stack
// this method also does the tracing and the steppping
/* {1} OPT.The go function for: push_debug(prop:property,arity:integer,start:integer) [] */
func F_push_debug_property (prop *ClaireProperty ,arity int,start int) EID { 
    var Result EID 
    /* Let:2 */{ 
      var i int  = ClEnv.Index
      /* noccur = 2 */
      /* Let:3 */{ 
        var n int  = ClEnv.Trace_I
        /* noccur = 3 */
        if ((n > 0) && 
            ((prop.Trace_I+ClEnv.Verbose) > 4)) /* If:4 */{ 
          /* Let:5 */{ 
            var p *ClairePort   = ClEnv.Ctrace.UseAsOutput()
            /* noccur = 1 */
            ClEnv.Trace_I = 0
            F_tr_indent_boolean(CFALSE,n)
            PRINC(" ")
            Result = F_print_any(prop.Id())
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC("(")
            Result = F_CALL(C_print,ARGS(ClEnv.EvalStack[start]))
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC("")
            Result = EVOID
            }}
            {
            /* Let:6 */{ 
              var j int  = (start+1)
              /* noccur = 4 */
              Result= EID{CFALSE.Id(),0}
              for (j < (start+arity)) /* while:7 */{ 
                var void_try8 EID 
                _ = void_try8
                { 
                PRINC(",")
                void_try8 = F_CALL(C_print,ARGS(ClEnv.EvalStack[j]))
                /* ERROR PROTECTION INSERTED (void_try8-void_try8) */
                if ErrorIn(void_try8) {Result = void_try8
                break
                } else {
                PRINC("")
                void_try8 = EVOID
                }
                {
                j = (j+1)
                }
                /* while-7 */} 
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            if (ClEnv.CountCall >= 0) /* If:6 */{ 
              ClEnv.CountCall = (ClEnv.CountCall+1)
              PRINC(" [")
              F_princ_integer(ClEnv.CountCall)
              PRINC("]")
              if (ClEnv.CountCall == ClEnv.CountLevel) /* If:7 */{ 
                if (ClEnv.CountTrigger == C_spy.Id()) /* If:8 */{ 
                  Result = F_update_property(C_spy_I,
                    ToObject(ClEnv.Id()),
                    16,
                    C_object,
                    F__at_property1(C_spy,C_void).Id())
                  } else {
                  /* update:9 */{ 
                    var va_arg1 *ClaireEnvironment  
                    var va_arg2 int 
                    va_arg1 = ClEnv
                    va_arg2 = ToInteger(ClEnv.CountTrigger).Value
                    /* ---------- now we compile update verbose(va_arg1) := va_arg2 ------- */
                    va_arg1.Verbose = va_arg2
                    Result = EID{C__INT,IVAL(va_arg2)}
                    /* update-9 */} 
                  /* If-8 */} 
                } else {
                Result = EID{CFALSE.Id(),0}
                /* If-7 */} 
              } else {
              Result = EID{CFALSE.Id(),0}
              /* If-6 */} 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC(")\n")
            ClEnv.Trace_I = (n+1)
            Result = p.UseAsOutput().ToEID()
            }}}
            /* Let-5 */} 
          } else {
          Result = EID{CFALSE.Id(),0}
          /* If-4 */} 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        if (F_get_table(C_Core_StopProperty,prop.Id()) != CNULL) /* If:4 */{ 
          var g0078I *ClaireBoolean  
          var g0078I_try00795 EID 
          /* or:5 */{ 
            var v_or5 *ClaireBoolean  
            
            var v_or5_try00806 EID 
            /* Let:6 */{ 
              var g0081UU *ClaireAny  
              /* noccur = 1 */
              var g0081UU_try00827 EID 
              g0081UU_try00827 = F_nth_table1(C_Core_StopProperty,prop.Id())
              /* ERROR PROTECTION INSERTED (g0081UU-v_or5_try00806) */
              if ErrorIn(g0081UU_try00827) {v_or5_try00806 = g0081UU_try00827
              } else {
              g0081UU = ANY(g0081UU_try00827)
              v_or5_try00806 = EID{Equal(g0081UU,CNIL.Id()).Id(),0}
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (v_or5-g0078I_try00795) */
            if ErrorIn(v_or5_try00806) {g0078I_try00795 = v_or5_try00806
            } else {
            v_or5 = ToBoolean(OBJ(v_or5_try00806))
            if (v_or5 == CTRUE) {g0078I_try00795 = EID{CTRUE.Id(),0}
            } else /* or:6 */{ 
              var v_or5_try00837 EID 
              /* Let:7 */{ 
                var g0084UU *ClaireAny  
                /* noccur = 1 */
                var g0084UU_try00858 EID 
                /* For:8 */{ 
                  var l2 *ClaireAny  
                  _ = l2
                  g0084UU_try00858= EID{CFALSE.Id(),0}
                  var l2_support *ClaireList  
                  var l2_support_try00869 EID 
                  l2_support_try00869 = F_nth_table1(C_Core_StopProperty,prop.Id())
                  /* ERROR PROTECTION INSERTED (l2_support-g0084UU_try00858) */
                  if ErrorIn(l2_support_try00869) {g0084UU_try00858 = l2_support_try00869
                  } else {
                  l2_support = ToList(OBJ(l2_support_try00869))
                  l2_len := l2_support.Length()
                  for i_it := 0; i_it < l2_len; i_it++ { 
                    l2 = l2_support.At(i_it)
                    var void_try10 EID 
                    _ = void_try10
                    var g0087I *ClaireBoolean  
                    var g0087I_try008810 EID 
                    /* Let:10 */{ 
                      var g0089UU *ClaireAny  
                      /* noccur = 1 */
                      var g0089UU_try009011 EID 
                      /* Let:11 */{ 
                        var j int  = 1
                        /* noccur = 6 */
                        /* Let:12 */{ 
                          var g0077 int  = INT(F_CALL(C_length,ARGS(l2.ToEID())))
                          /* noccur = 1 */
                          g0089UU_try009011= EID{CFALSE.Id(),0}
                          for (j <= g0077) /* while:13 */{ 
                            var void_try14 EID 
                            _ = void_try14
                            { 
                            var g0091I *ClaireBoolean  
                            var g0091I_try009214 EID 
                            /* Let:14 */{ 
                              var g0093UU *ClaireBoolean  
                              /* noccur = 1 */
                              var g0093UU_try009415 EID 
                              /* and:15 */{ 
                                var v_and15 *ClaireBoolean  
                                
                                v_and15 = F__inf_equal_integer((j+start),i)
                                if (v_and15 == CFALSE) {g0093UU_try009415 = EID{CFALSE.Id(),0}
                                } else /* arg:16 */{ 
                                  var v_and15_try009517 EID 
                                  /* Let:17 */{ 
                                    var g0096UU *ClaireAny  
                                    /* noccur = 1 */
                                    var g0096UU_try009718 EID 
                                    g0096UU_try009718 = F_CALL(C_nth,ARGS(l2.ToEID(),EID{C__INT,IVAL(j)}))
                                    /* ERROR PROTECTION INSERTED (g0096UU-v_and15_try009517) */
                                    if ErrorIn(g0096UU_try009718) {v_and15_try009517 = g0096UU_try009718
                                    } else {
                                    g0096UU = ANY(g0096UU_try009718)
                                    v_and15_try009517 = EID{Equal(g0096UU,ANY(ClEnv.EvalStack[((start+j)-1)])).Id(),0}
                                    }
                                    /* Let-17 */} 
                                  /* ERROR PROTECTION INSERTED (v_and15-g0093UU_try009415) */
                                  if ErrorIn(v_and15_try009517) {g0093UU_try009415 = v_and15_try009517
                                  } else {
                                  v_and15 = ToBoolean(OBJ(v_and15_try009517))
                                  if (v_and15 == CFALSE) {g0093UU_try009415 = EID{CFALSE.Id(),0}
                                  } else /* arg:17 */{ 
                                    g0093UU_try009415 = EID{CTRUE.Id(),0}/* arg-17 */} 
                                  /* arg-16 */} 
                                }
                                /* and-15 */} 
                              /* ERROR PROTECTION INSERTED (g0093UU-g0091I_try009214) */
                              if ErrorIn(g0093UU_try009415) {g0091I_try009214 = g0093UU_try009415
                              } else {
                              g0093UU = ToBoolean(OBJ(g0093UU_try009415))
                              g0091I_try009214 = EID{g0093UU.Not.Id(),0}
                              }
                              /* Let-14 */} 
                            /* ERROR PROTECTION INSERTED (g0091I-void_try14) */
                            if ErrorIn(g0091I_try009214) {void_try14 = g0091I_try009214
                            } else {
                            g0091I = ToBoolean(OBJ(g0091I_try009214))
                            if (g0091I == CTRUE) /* If:14 */{ 
                               /*v = g0089UU_try009011, s =EID*/
g0089UU_try009011 = EID{CTRUE.Id(),0}
                              break
                              } else {
                              void_try14 = EID{CFALSE.Id(),0}
                              /* If-14 */} 
                            }
                            /* ERROR PROTECTION INSERTED (void_try14-void_try14) */
                            if ErrorIn(void_try14) {g0089UU_try009011 = void_try14
                            break
                            } else {
                            j = (j+1)
                            }
                            /* while-13 */} 
                          }
                          /* Let-12 */} 
                        /* Let-11 */} 
                      /* ERROR PROTECTION INSERTED (g0089UU-g0087I_try008810) */
                      if ErrorIn(g0089UU_try009011) {g0087I_try008810 = g0089UU_try009011
                      } else {
                      g0089UU = ANY(g0089UU_try009011)
                      g0087I_try008810 = EID{F_not_any(g0089UU).Id(),0}
                      }
                      /* Let-10 */} 
                    /* ERROR PROTECTION INSERTED (g0087I-void_try10) */
                    if ErrorIn(g0087I_try008810) {void_try10 = g0087I_try008810
                    } else {
                    g0087I = ToBoolean(OBJ(g0087I_try008810))
                    if (g0087I == CTRUE) /* If:10 */{ 
                       /*v = g0084UU_try00858, s =EID*/
g0084UU_try00858 = EID{CTRUE.Id(),0}
                      break
                      } else {
                      void_try10 = EID{CFALSE.Id(),0}
                      /* If-10 */} 
                    }
                    /* ERROR PROTECTION INSERTED (void_try10-g0084UU_try00858) */
                    if ErrorIn(void_try10) {g0084UU_try00858 = void_try10
                    g0084UU_try00858 = void_try10
                    break
                    } else {
                    }}
                    /* loop-9 */} 
                  /* For-8 */} 
                /* ERROR PROTECTION INSERTED (g0084UU-v_or5_try00837) */
                if ErrorIn(g0084UU_try00858) {v_or5_try00837 = g0084UU_try00858
                } else {
                g0084UU = ANY(g0084UU_try00858)
                v_or5_try00837 = EID{F_boolean_I_any(g0084UU).Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (v_or5-g0078I_try00795) */
              if ErrorIn(v_or5_try00837) {g0078I_try00795 = v_or5_try00837
              } else {
              v_or5 = ToBoolean(OBJ(v_or5_try00837))
              if (v_or5 == CTRUE) {g0078I_try00795 = EID{CTRUE.Id(),0}
              } else /* or:7 */{ 
                g0078I_try00795 = EID{CFALSE.Id(),0}/* org-7 */} 
              /* org-6 */} 
            }}
            /* or-5 */} 
          /* ERROR PROTECTION INSERTED (g0078I-Result) */
          if ErrorIn(g0078I_try00795) {Result = g0078I_try00795
          } else {
          g0078I = ToBoolean(OBJ(g0078I_try00795))
          if (g0078I == CTRUE) /* If:5 */{ 
            Result = ToException(C_general_error.Make(MakeString("stop as required in ~S(~A)").Id(),MakeConstantList(prop.Id(),F_get_args_integer(start).Id()).Id())).Close()
            } else {
            Result = EID{CFALSE.Id(),0}
            /* If-5 */} 
          }
          } else {
          Result = EID{CFALSE.Id(),0}
          /* If-4 */} 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        ClEnv.Push(EID{C__INT,IVAL(ClEnv.Debug_I)})
        ClEnv.Push(EID{prop.Id(),0})
        ClEnv.Push(EID{C__INT,IVAL(arity)})
        ClEnv.Push(EID{C__INT,IVAL(start)})
        /* update:4 */{ 
          var va_arg1 *ClaireEnvironment  
          var va_arg2 int 
          va_arg1 = ClEnv
          va_arg2 = i
          /* ---------- now we compile update debug!(va_arg1) := va_arg2 ------- */
          va_arg1.Debug_I = va_arg2
          Result = EID{C__INT,IVAL(va_arg2)}
          /* update-4 */} 
        }}
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: push_debug @ property (throw: true) 
func E_push_debug_property (prop EID,arity EID,start EID) EID { 
    return /*(sm for push_debug @ property= EID)*/ F_push_debug_property(ToProperty(OBJ(prop)),INT(arity),INT(start) )} 
  
// value of the previous debug
// n is 0 for interpreted code and 1 for compiled code
/* {1} OPT.The go function for: pop_debug(self:property,n:integer,val:any) [] */
func F_pop_debug_property (self *ClaireProperty ,n int,val *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var v int  = ClEnv.Debug_I
      /* noccur = 2 */
      if (v > 0) /* If:3 */{ 
        if (n != 0) /* If:4 */{ 
          ClEnv.Index= INT(ClEnv.EvalStack[(v+3)])
          /* If-4 */} 
        ClEnv.Debug_I = INT(ClEnv.EvalStack[ClEnv.Debug_I])
        if (self.IfWrite == CNULL) /* If:4 */{ 
          /* Let:5 */{ 
            var m *ClaireObject   = ClEnv.Spy_I
            /* noccur = 3 */
            if (m.Id() != CNULL) /* If:6 */{ 
              ClEnv.Spy_I = ToObject(CNULL)
              Result = F_funcall_method1(ToMethod(m.Id()),ClEnv.Id())
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              /* update:7 */{ 
                var va_arg1 *ClaireEnvironment  
                var va_arg2 *ClaireObject  
                va_arg1 = ClEnv
                va_arg2 = m
                /* ---------- now we compile update spy!(va_arg1) := va_arg2 ------- */
                va_arg1.Spy_I = va_arg2
                Result = EID{va_arg2.Id(),0}
                /* update-7 */} 
              }
              } else {
              Result = EID{CFALSE.Id(),0}
              /* If-6 */} 
            /* Let-5 */} 
          } else {
          Result = EID{CFALSE.Id(),0}
          /* If-4 */} 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        if ((ClEnv.Trace_I > 1) && 
            ((self.Trace_I+ClEnv.Verbose) > 4)) /* If:4 */{ 
          /* Let:5 */{ 
            var i int  = ClEnv.Trace_I
            /* noccur = 2 */
            ClEnv.Trace_I = 0
            if ((self.Trace_I+ClEnv.Verbose) > 4) /* If:6 */{ 
              /* Let:7 */{ 
                var p *ClairePort   = ClEnv.Ctrace.UseAsOutput()
                /* noccur = 1 */
                F_tr_indent_boolean(CTRUE,(i-1))
                PRINC(" ")
                Result = F_CALL(C_print,ARGS(val.ToEID()))
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                PRINC("\n")
                Result = EVOID
                }
                {
                Result = p.UseAsOutput().ToEID()
                }
                /* Let-7 */} 
              } else {
              Result = EID{CFALSE.Id(),0}
              /* If-6 */} 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            /* update:6 */{ 
              var va_arg1 *ClaireEnvironment  
              var va_arg2 int 
              va_arg1 = ClEnv
              va_arg2 = (i-1)
              /* ---------- now we compile update trace!(va_arg1) := va_arg2 ------- */
              va_arg1.Trace_I = va_arg2
              Result = EID{C__INT,IVAL(va_arg2)}
              /* update-6 */} 
            }
            /* Let-5 */} 
          } else {
          Result = EID{CFALSE.Id(),0}
          /* If-4 */} 
        }
        } else {
        Result = EID{CFALSE.Id(),0}
        /* If-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: pop_debug @ property (throw: true) 
func E_pop_debug_property (self EID,n EID,val EID) EID { 
    return /*(sm for pop_debug @ property= EID)*/ F_pop_debug_property(ToProperty(OBJ(self)),INT(n),ANY(val) )} 
  
// print a nice indented mark
/* {1} OPT.The go function for: tr_indent(return?:boolean,n:integer) [] */
func F_tr_indent_boolean (return_ask *ClaireBoolean ,n int)  { 
    // procedure body with s =  
if (return_ask == CTRUE) /* If:2 */{ 
      PRINC("[")
      F_princ_integer(n)
      PRINC("]")
      } else {
      F_princ_integer(n)
      PRINC(":=")
      /* If-2 */} 
    for (n > 9) /* while:2 */{ 
      PRINC("=")
      n = (n-10)
      /* while-2 */} 
    for (n > 0) /* while:2 */{ 
      PRINC(">")
      n = (n-1)
      /* while-2 */} 
    } 
  
// The EID go function for: tr_indent @ boolean (throw: false) 
func E_tr_indent_boolean (return_ask EID,n EID) EID { 
    /*(sm for tr_indent @ boolean= void)*/ F_tr_indent_boolean(ToBoolean(OBJ(return_ask)),INT(n) )
    return EVOID} 
  
// *********************************************************************
// *   Part 2: Tables                                                  *
// *********************************************************************
// finds if objects are identified - unclear if there is any need for this
/* {1} OPT.The go function for: identified?(self:class) [] */
func F_identified_ask_class (self *ClaireClass ) *ClaireBoolean  { 
    // use function body compiling 
return  MakeBoolean((self.Id() == C_integer.Id()) || (self.IsIn(C_object) == CTRUE) || (self.Id() == C_symbol.Id()) || (self.Id() == C_boolean.Id()) || (self.Id() == C_char.Id()))
    } 
  
// The EID go function for: identified? @ class (throw: false) 
func E_identified_ask_class (self EID) EID { 
    return EID{/*(sm for identified? @ class= boolean)*/ F_identified_ask_class(ToClass(OBJ(self)) ).Id(),0}} 
  
// true pointer equality in go (used to be C++) => use externC form
/* {1} OPT.The go function for: identical?(x:any,y:any) [] */
func F_identical_ask_any (x *ClaireAny ,y *ClaireAny ) *ClaireBoolean  { 
    // use function body compiling 
return  ToBoolean(IfThenElse(x == y,CTRUE.Id(),CFALSE.Id()))
    } 
  
// The EID go function for: identical? @ any (throw: false) 
func E_identical_ask_any (x EID,y EID) EID { 
    return EID{/*(sm for identical? @ any= boolean)*/ F_identical_ask_any(ANY(x),ANY(y) ).Id(),0}} 
  
// writing a single value into a slot but does NOT trigger the rules !
// equivalent to is! of LAURE
// this definition should not be placed in the method.cl file
// (it requires some inheritance conflict processing)
/* {1} OPT.The go function for: put(self:property,x:object,y:any) [] */
func F_put_property2 (self *ClaireProperty ,x *ClaireObject ,y *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var s *ClaireObject   = F__at_property1(self,x.Id().Isa)
      /* noccur = 2 */
      if (C_slot.Id() == s.Isa.Id()) /* If:3 */{ 
        /* Let:4 */{ 
          var g0098 *ClaireSlot   = ToSlot(s.Id())
          /* noccur = 2 */
          Result = F_store_object(x,
            g0098.Index,
            g0098.Srange,
            y,
            self.Store_ask).ToEID()
          /* Let-4 */} 
        } else {
        Result = ToException(C_selector_error.Make(self.Id(),MakeConstantList(x.Id()).Id())).Close()
        /* If-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: put @ list<type_expression>(property, object, any) (throw: true) 
func E_put_property2 (self EID,x EID,y EID) EID { 
    return /*(sm for put @ list<type_expression>(property, object, any)= EID)*/ F_put_property2(ToProperty(OBJ(self)),ToObject(OBJ(x)),ANY(y) )} 
  
// v3.2 : same but multi valued
/* {1} OPT.The go function for: add_value(self:property,x:object,y:any) [] */
func F_add_value_property3 (self *ClaireProperty ,x *ClaireObject ,y *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var s *ClaireObject   = F__at_property1(self,x.Id().Isa)
      /* noccur = 2 */
      if (F_boolean_I_any(s.Id()).Id() != CTRUE.Id()) /* If:3 */{ 
        Result = ToException(C_selector_error.Make(self.Id(),MakeConstantList(x.Id()).Id())).Close()
        /* If!3 */}  else if (self.Multivalued_ask != CTRUE) /* If:3 */{ 
        Result = ToException(C_general_error.Make(MakeString("[134] Cannot apply add to ~S").Id(),MakeConstantList(self.Id()).Id())).Close()
        } else {
        /* Let:4 */{ 
          var n int  = ToSlot(s.Id()).Index
          /* noccur = 2 */
          /* Let:5 */{ 
            var l1 *ClaireSet   = ToSet(x.SlotGet(n,C_object))
            /* noccur = 1 */
            Result = EID{F_Core_add_value_I_property(self,
              x,
              n,
              l1,
              y).Id(),0}
            /* Let-5 */} 
          /* Let-4 */} 
        /* If-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: add_value @ property (throw: true) 
func E_add_value_property3 (self EID,x EID,y EID) EID { 
    return /*(sm for add_value @ property= EID)*/ F_add_value_property3(ToProperty(OBJ(self)),ToObject(OBJ(x)),ANY(y) )} 
  
// Claire4: add_value! is the internal form
// a table is implemented through its graph which is a list or a dictionary (a.params says which)
// graph_get(a:table,x:any) : any -> reads in a.graph
// graph_put(a:table,x:any,y:any) : void -> write in a.graph    
// access
// in Claire4 there is always a default hence the unknown check has disapeared
/* {1} OPT.The go function for: nth(a:table,x:any) [] */
func F_nth_table1 (a *ClaireTable ,x *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var p *ClaireAny   = a.Params
      /* noccur = 3 */
      if (a.Domain.Contains(x) != CTRUE) /* If:3 */{ 
        Result = ToException(C_general_error.Make(MakeString("[135] ~S does not belong to the domain of ~S").Id(),MakeConstantList(x,a.Id()).Id())).Close()
        } else {
        Result = EID{CFALSE.Id(),0}
        /* If-3 */} 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      if (C_integer.Id() == p.Isa.Id()) /* If:3 */{ 
        /* Let:4 */{ 
          var g0100 int  = ToInteger(p).Value
          /* noccur = 1 */
          Result = ToList(a.Graph).At((ToInteger(x).Value-g0100)-1).ToEID()
          /* Let-4 */} 
        /* If!3 */}  else if (p.Isa.IsIn(C_list) == CTRUE) /* If:3 */{ 
        Result = ToList(a.Graph).At(F_get_index_table2(a,ToInteger(ToList(x).At(1-1)).Value,ToInteger(ToList(x).At(2-1)).Value)-1).ToEID()
        } else {
        Result = a.GraphGet(x).ToEID()
        /* If-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: nth @ list<type_expression>(table, any) (throw: true) 
func E_nth_table1 (a EID,x EID) EID { 
    return /*(sm for nth @ list<type_expression>(table, any)= EID)*/ F_nth_table1(ToTable(OBJ(a)),ANY(x) )} 
  
/* {1} OPT.The go function for: nth_table1_type */
func F_nth_table1_type (a *ClaireType ,x *ClaireType ) EID { 
    /* eid body: (if unique?(a) the(a).range else any) */
    var Result EID 
    if (ToBoolean(OBJ(F_CALL(C_unique_ask,ARGS(EID{a.Id(),0})))) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0103UU *ClaireAny  
        /* noccur = 1 */
        var g0103UU_try01044 EID 
        g0103UU_try01044 = F_CALL(C_the,ARGS(EID{a.Id(),0}))
        /* ERROR PROTECTION INSERTED (g0103UU-Result) */
        if ErrorIn(g0103UU_try01044) {Result = g0103UU_try01044
        } else {
        g0103UU = ANY(g0103UU_try01044)
        Result = F_CALL(C_range,ARGS(g0103UU.ToEID()))
        }
        /* Let-3 */} 
      } else {
      Result = EID{C_any.Id(),0}
      /* If-2 */} 
    return Result} 
  
  
// The dual EID go function for: "nth_table1_type" 
func E_nth_table1_type (a EID,x EID) EID { 
    return F_nth_table1_type(ToType(OBJ(a)),ToType(OBJ(x)))} 
  
// get is the same, with no error            
/* {1} OPT.The go function for: get(a:table,x:any) [] */
func F_get_table (a *ClaireTable ,x *ClaireAny ) *ClaireAny  { 
    // procedure body with s =  
var Result *ClaireAny  
    /* Let:2 */{ 
      var p *ClaireAny   = a.Params
      /* noccur = 3 */
      if (C_integer.Id() == p.Isa.Id()) /* If:3 */{ 
        /* Let:4 */{ 
          var g0105 int  = ToInteger(p).Value
          /* noccur = 1 */
          Result = ToList(a.Graph).At((ToInteger(x).Value-g0105)-1)
          /* Let-4 */} 
        /* If!3 */}  else if (p.Isa.IsIn(C_list) == CTRUE) /* If:3 */{ 
        Result = ToList(a.Graph).At(F_get_index_table2(a,ToInteger(ToList(x).At(1-1)).Value,ToInteger(ToList(x).At(2-1)).Value)-1)
        } else {
        Result = a.GraphGet(x)
        /* If-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: get @ table (throw: false) 
func E_get_table (a EID,x EID) EID { 
    return /*(sm for get @ table= any)*/ F_get_table(ToTable(OBJ(a)),ANY(x) ).ToEID()} 
  
/* {1} OPT.The go function for: get_table_type */
func F_get_table_type (a *ClaireType ,x *ClaireType ) EID { 
    /* eid body: (if unique?(a) the(a).range else any) */
    var Result EID 
    if (ToBoolean(OBJ(F_CALL(C_unique_ask,ARGS(EID{a.Id(),0})))) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0108UU *ClaireAny  
        /* noccur = 1 */
        var g0108UU_try01094 EID 
        g0108UU_try01094 = F_CALL(C_the,ARGS(EID{a.Id(),0}))
        /* ERROR PROTECTION INSERTED (g0108UU-Result) */
        if ErrorIn(g0108UU_try01094) {Result = g0108UU_try01094
        } else {
        g0108UU = ANY(g0108UU_try01094)
        Result = F_CALL(C_range,ARGS(g0108UU.ToEID()))
        }
        /* Let-3 */} 
      } else {
      Result = EID{C_any.Id(),0}
      /* If-2 */} 
    return Result} 
  
  
// The dual EID go function for: "get_table_type" 
func E_get_table_type (a EID,x EID) EID { 
    return F_get_table_type(ToType(OBJ(a)),ToType(OBJ(x)))} 
  
// interface update method for a[x] := y
/* {1} OPT.The go function for: nth=(a:table,x:any,y:any) [] */
func F_nth_equal_table1 (a *ClaireTable ,x *ClaireAny ,y *ClaireAny ) EID { 
    var Result EID 
    if (a.Domain.Contains(x) != CTRUE) /* If:2 */{ 
      Result = ToException(C_general_error.Make(MakeString("[135] ~S does not belong to the domain of ~S").Id(),MakeConstantList(x,a.Id()).Id())).Close()
      } else {
      Result = EID{CFALSE.Id(),0}
      /* If-2 */} 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    if (a.Range.Contains(y) != CTRUE) /* If:2 */{ 
      Result = ToException(C_range_error.Make(a.Id(),y,a.Range.Id())).Close()
      } else {
      Result = EID{CFALSE.Id(),0}
      /* If-2 */} 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    Result = F_nth_put_table(a,x,y)
    }}
    return Result} 
  
// The EID go function for: nth= @ list<type_expression>(table, any, any) (throw: true) 
func E_nth_equal_table1 (a EID,x EID,y EID) EID { 
    return /*(sm for nth= @ list<type_expression>(table, any, any)= EID)*/ F_nth_equal_table1(ToTable(OBJ(a)),ANY(x),ANY(y) )} 
  
// internal form without checks
// equivalent of update = put + put_inverse
/* {1} OPT.The go function for: nth_put(a:table,x:any,y:any) [] */
func F_nth_put_table (a *ClaireTable ,x *ClaireAny ,y *ClaireAny ) EID { 
    var Result EID 
    if ((a.IfWrite != CNULL) && 
        (a.Multivalued_ask != CTRUE)) /* If:2 */{ 
      Result = F_fastcall_relation2(ToRelation(a.Id()),x,y)
      /* If!2 */}  else if (a.Multivalued_ask == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var r *ClaireAny   = F_get_property(C_inverse,ToObject(a.Id()))
        /* noccur = 2 */
        /* Let:4 */{ 
          var old *ClaireSet   = ToSet(F_get_table(a,x))
          /* noccur = 2 */
          if ((old.Id() != CNULL) && 
              (r != CNULL)) /* If:5 */{ 
            /* For:6 */{ 
              var z *ClaireAny  
              _ = z
              for _,z = range(old.Values)/* loop:7 */{ 
                F_update_dash_relation(ToRelation(r),z,x)
                /* loop-7 */} 
              /* For-6 */} 
            /* If-5 */} 
          F_put_table(a,x,y)
          /* For:5 */{ 
            var z *ClaireAny  
            _ = z
            Result= EID{CFALSE.Id(),0}
            var z_support *ClaireSet  
            z_support = ToSet(y)
            for _,z = range(z_support.Values)/* loop2:6 */{ 
              var void_try7 EID 
              _ = void_try7
              void_try7 = F_update_plus_relation(ToRelation(a.Id()),x,z)
              /* ERROR PROTECTION INSERTED (void_try7-Result) */
              if ErrorIn(void_try7) {Result = void_try7
              Result = void_try7
              break
              } else {
              }
              /* loop-6 */} 
            /* For-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      } else {
      /* Let:3 */{ 
        var r *ClaireAny   = F_get_property(C_inverse,ToObject(a.Id()))
        /* noccur = 3 */
        /* Let:4 */{ 
          var z *ClaireAny   = F_get_table(a,x)
          /* noccur = 1 */
          if (Equal(z,y) != CTRUE) /* If:5 */{ 
            if (r != CNULL) /* If:6 */{ 
              /* Let:7 */{ 
                var z *ClaireAny   = F_get_table(a,x)
                /* noccur = 3 */
                if ((z != CNULL) && 
                    ((r != a.Id()) || 
                        (Equal(x,z) != CTRUE))) /* If:8 */{ 
                  F_update_dash_relation(ToRelation(r),z,x)
                  /* If-8 */} 
                /* Let-7 */} 
              /* If-6 */} 
            F_put_table(a,x,y)
            Result = F_update_plus_relation(ToRelation(a.Id()),x,y)
            } else {
            Result = EID{CFALSE.Id(),0}
            /* If-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* If-2 */} 
    return Result} 
  
// The EID go function for: nth_put @ table (throw: true) 
func E_nth_put_table (a EID,x EID,y EID) EID { 
    return /*(sm for nth_put @ table= EID)*/ F_nth_put_table(ToTable(OBJ(a)),ANY(x),ANY(y) )} 
  
// update inverse 
// put does NOT update the inverse, but handles store ...
/* {1} OPT.The go function for: put(a:table,x:any,y:any) [] */
func F_put_table (a *ClaireTable ,x *ClaireAny ,y *ClaireAny )  { 
    // procedure body with s =  
/* Let:2 */{ 
      var p *ClaireAny   = a.Params
      /* noccur = 3 */
      /* Let:3 */{ 
        var z *ClaireAny   = F_get_table(a,x)
        /* noccur = 1 */
        if (Equal(z,y) != CTRUE) /* If:4 */{ 
          if (C_integer.Id() == p.Isa.Id()) /* If:5 */{ 
            /* Let:6 */{ 
              var g0110 int  = ToInteger(p).Value
              /* noccur = 1 */
              F_CALL(C_store,ARGS(a.Graph.ToEID(),
                F_CALL(ToProperty(C__dash.Id()),ARGS(x.ToEID(),EID{C__INT,IVAL(g0110)})),
                y.ToEID(),
                EID{a.Store_ask.Id(),0}))
              /* Let-6 */} 
            /* If!5 */}  else if (p.Isa.IsIn(C_list) == CTRUE) /* If:5 */{ 
            F_CALL(C_store,ARGS(a.Graph.ToEID(),
              EID{C__INT,IVAL(F_get_index_table2(a,ToInteger(ToList(x).At(1-1)).Value,ToInteger(ToList(x).At(2-1)).Value))},
              y.ToEID(),
              EID{a.Store_ask.Id(),0}))
            } else {
            a.GraphPut(x,y)
            /* If-5 */} 
          /* If-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    } 
  
// The EID go function for: put @ table (throw: false) 
func E_put_table (a EID,x EID,y EID) EID { 
    /*(sm for put @ table= void)*/ F_put_table(ToTable(OBJ(a)),ANY(x),ANY(y) )
    return EVOID} 
  
// takes care of the defeasible part :)
// adds a value to a multi-valued table: interface method
/* {1} OPT.The go function for: add(a:table,x:any,y:any) [] */
func F_add_table (a *ClaireTable ,x *ClaireAny ,y *ClaireAny ) EID { 
    var Result EID 
    if (a.Domain.Contains(x) != CTRUE) /* If:2 */{ 
      Result = ToException(C_general_error.Make(MakeString("[135] ~S does not belong to the domain of ~S").Id(),MakeConstantList(x,a.Id()).Id())).Close()
      } else {
      Result = EID{CFALSE.Id(),0}
      /* If-2 */} 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    if (F_member_type(a.Range).Contains(y) != CTRUE) /* If:2 */{ 
      Result = ToException(C_range_error.Make(a.Id(),y,a.Range.Id())).Close()
      } else {
      Result = EID{CFALSE.Id(),0}
      /* If-2 */} 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    Result = F_add_I_table(a,x,y)
    }}
    return Result} 
  
// The EID go function for: add @ table (throw: true) 
func E_add_table (a EID,x EID,y EID) EID { 
    return /*(sm for add @ table= EID)*/ F_add_table(ToTable(OBJ(a)),ANY(x),ANY(y) )} 
  
// adds a value to a multi-valued table: internal version without type checks
/* {1} OPT.The go function for: add!(a:table,x:any,y:any) [] */
func F_add_I_table (a *ClaireTable ,x *ClaireAny ,y *ClaireAny ) EID { 
    var Result EID 
    if (a.IfWrite != CNULL) /* If:2 */{ 
      Result = F_fastcall_relation2(ToRelation(a.Id()),x,y)
      } else {
      /* Let:3 */{ 
        var old *ClaireSet   = ToSet(F_get_table(a,x))
        /* noccur = 1 */
        if (F_Core_add_value_I_table(a,x,old,y) == CTRUE) /* If:4 */{ 
          Result = F_update_plus_relation(ToRelation(a.Id()),x,y)
          } else {
          Result = EID{CFALSE.Id(),0}
          /* If-4 */} 
        /* Let-3 */} 
      /* If-2 */} 
    return Result} 
  
// The EID go function for: add! @ table (throw: true) 
func E_add_I_table (a EID,x EID,y EID) EID { 
    return /*(sm for add! @ table= EID)*/ F_add_I_table(ToTable(OBJ(a)),ANY(x),ANY(y) )} 
  
// this methods adds a value to a multi-valued table (used by the compiler)
// s1 is the current value in the table
/* {1} OPT.The go function for: add_value!(self:table,x:any,s1:set,y:any) [] */
func F_Core_add_value_I_table (self *ClaireTable ,x *ClaireAny ,s1 *ClaireSet ,y *ClaireAny ) *ClaireBoolean  { 
    // procedure body with s =  
var Result *ClaireBoolean  
    if (s1.Contain_ask(y) != CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var s2 *ClaireSet  
        /* noccur = 1 */
        /* Let:4 */{ 
          var g0113UU *ClaireSet  
          /* noccur = 1 */
          if (self.Store_ask == CTRUE) /* If:5 */{ 
            g0113UU = s1.Copy()
            } else {
            g0113UU = s1
            /* If-5 */} 
          s2 = g0113UU.AddFast(y)
          /* Let-4 */} 
        F_put_table(self,x,s2.Id())
        Result = CTRUE
        /* Let-3 */} 
      } else {
      Result = CFALSE
      /* If-2 */} 
    return Result} 
  
// The EID go function for: add_value! @ table (throw: false) 
func E_Core_add_value_I_table (self EID,x EID,s1 EID,y EID) EID { 
    return EID{/*(sm for add_value! @ table= boolean)*/ F_Core_add_value_I_table(ToTable(OBJ(self)),
      ANY(x),
      ToSet(OBJ(s1)),
      ANY(y) ).Id(),0}} 
  
// a direct version (v3.2) that can be used in lieu of add!
/* {1} OPT.The go function for: add_value(self:table,x:any,y:any) [] */
func F_add_value_table3 (self *ClaireTable ,x *ClaireAny ,y *ClaireAny )  { 
    // procedure body with s =  
/* Let:2 */{ 
      var old *ClaireSet   = ToSet(F_get_table(self,x))
      /* noccur = 1 */
      F_Core_add_value_I_table(self,x,old,y)
      /* Let-2 */} 
    } 
  
// The EID go function for: add_value @ table (throw: false) 
func E_add_value_table3 (self EID,x EID,y EID) EID { 
    /*(sm for add_value @ table= void)*/ F_add_value_table3(ToTable(OBJ(self)),ANY(x),ANY(y) )
    return EVOID} 
  
// removes a value from an table (multivalued only)
/* {1} OPT.The go function for: delete(a:table,x:any,y:any) [] */
func F_delete_table (a *ClaireTable ,x *ClaireAny ,y *ClaireAny ) *ClaireAny  { 
    // procedure body with s =  
var Result *ClaireAny  
    /* Let:2 */{ 
      var old *ClaireSet   = ToSet(F_get_table(a,x))
      /* noccur = 4 */
      if (old.Contain_ask(y) == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var s *ClaireSet  
          /* noccur = 2 */
          /* Let:5 */{ 
            var g0114UU *ClaireSet  
            /* noccur = 1 */
            if (a.Store_ask == CTRUE) /* If:6 */{ 
              g0114UU = old.Copy()
              } else {
              g0114UU = old
              /* If-6 */} 
            s = g0114UU.Delete(y)
            /* Let-5 */} 
          F_put_table(a,x,s.Id())
          /* Let:5 */{ 
            var r *ClaireRelation   = a.Inverse
            /* noccur = 2 */
            if (r.Id() != CNULL) /* If:6 */{ 
              F_update_dash_relation(r,y,x)
              /* If-6 */} 
            /* Let-5 */} 
          Result = s.Id()
          /* Let-4 */} 
        } else {
        Result = old.Id()
        /* If-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: delete @ table (throw: false) 
func E_delete_table (a EID,x EID,y EID) EID { 
    return /*(sm for delete @ table= any)*/ F_delete_table(ToTable(OBJ(a)),ANY(x),ANY(y) ).ToEID()} 
  
// direct access to 2-dim tables
/* {1} OPT.The go function for: nth(a:table,x:any,y:any) [] */
func F_nth_table2 (a *ClaireTable ,x *ClaireAny ,y *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var p *ClaireAny   = a.Params
      /* noccur = 1 */
      if (p.Isa.IsIn(C_list) == CTRUE) /* If:3 */{ 
        if ((ToType(ToList(a.Domain.Id()).At(1-1)).Contains(x) != CTRUE) || 
            (ToType(ToList(a.Domain.Id()).At(2-1)).Contains(y) != CTRUE)) /* If:4 */{ 
          Result = ToException(C_general_error.Make(MakeString("[135] ~S does not belong to the domain of ~S").Id(),MakeConstantList(x,a.Id()).Id())).Close()
          } else {
          Result = EID{CFALSE.Id(),0}
          /* If-4 */} 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        Result = F_CALL(C_nth,ARGS(a.Graph.ToEID(),EID{C__INT,IVAL(F_get_index_table2(a,ToInteger(x).Value,ToInteger(y).Value))}))
        }
        } else {
        Result = F_nth_table1(a,MakeTuple(x,y).Id())
        /* If-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: nth @ list<type_expression>(table, any, any) (throw: true) 
func E_nth_table2 (a EID,x EID,y EID) EID { 
    return /*(sm for nth @ list<type_expression>(table, any, any)= EID)*/ F_nth_table2(ToTable(OBJ(a)),ANY(x),ANY(y) )} 
  
/* {1} OPT.The go function for: nth_table2_type */
func F_nth_table2_type (a *ClaireType ,x *ClaireType ,y *ClaireType ) EID { 
    /* eid body: (if unique?(a) the(a).range else any) */
    var Result EID 
    if (ToBoolean(OBJ(F_CALL(C_unique_ask,ARGS(EID{a.Id(),0})))) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0117UU *ClaireAny  
        /* noccur = 1 */
        var g0117UU_try01184 EID 
        g0117UU_try01184 = F_CALL(C_the,ARGS(EID{a.Id(),0}))
        /* ERROR PROTECTION INSERTED (g0117UU-Result) */
        if ErrorIn(g0117UU_try01184) {Result = g0117UU_try01184
        } else {
        g0117UU = ANY(g0117UU_try01184)
        Result = F_CALL(C_range,ARGS(g0117UU.ToEID()))
        }
        /* Let-3 */} 
      } else {
      Result = EID{C_any.Id(),0}
      /* If-2 */} 
    return Result} 
  
  
// The dual EID go function for: "nth_table2_type" 
func E_nth_table2_type (a EID,x EID,y EID) EID { 
    return F_nth_table2_type(ToType(OBJ(a)),ToType(OBJ(x)),ToType(OBJ(y)))} 
  
// sets a value in a 2-dim table
/* {1} OPT.The go function for: nth=(a:table,x:any,y:any,z:any) [] */
func F_nth_equal_table2 (a *ClaireTable ,x *ClaireAny ,y *ClaireAny ,z *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var p *ClaireAny   = a.Params
      /* noccur = 1 */
      if (p.Isa.IsIn(C_list) == CTRUE) /* If:3 */{ 
        if ((ToType(ToList(a.Domain.Id()).At(1-1)).Contains(x) != CTRUE) || 
            (ToType(ToList(a.Domain.Id()).At(2-1)).Contains(y) != CTRUE)) /* If:4 */{ 
          Result = ToException(C_general_error.Make(MakeString("[135] ~S does not belong to the domain of ~S").Id(),MakeConstantList(MakeConstantList(x,y).Id(),a.Id()).Id())).Close()
          } else {
          Result = EID{CFALSE.Id(),0}
          /* If-4 */} 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        if (a.Range.Contains(z) != CTRUE) /* If:4 */{ 
          Result = ToException(C_range_error.Make(a.Id(),z,a.Range.Id())).Close()
          } else {
          Result = EID{CFALSE.Id(),0}
          /* If-4 */} 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        if ((a.Inverse.Id() != CNULL) || 
            (a.IfWrite != CNULL)) /* If:4 */{ 
          Result = F_nth_put_table(a,MakeConstantList(x,y).Id(),z)
          } else {
          Result = F_store_list(ToList(a.Graph),F_get_index_table2(a,ToInteger(x).Value,ToInteger(y).Value),z,a.Store_ask).ToEID()
          /* If-4 */} 
        }}
        } else {
        Result = F_nth_equal_table1(a,MakeTuple(x,y).Id(),z)
        /* If-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: nth= @ list<type_expression>(table, any, any, any) (throw: true) 
func E_nth_equal_table2 (a EID,x EID,y EID,z EID) EID { 
    return /*(sm for nth= @ list<type_expression>(table, any, any, any)= EID)*/ F_nth_equal_table2(ToTable(OBJ(a)),
      ANY(x),
      ANY(y),
      ANY(z) )} 
  
// v3.2.16 tuple(a,b) is not list(a,b) !
/* {1} OPT.The go function for: get_index(a:table,x:any) [] */
func F_get_index_table1 (a *ClaireTable ,x *ClaireAny ) int { 
    // procedure body with s =  
var Result int 
    /* Let:2 */{ 
      var p *ClaireAny   = a.Params
      /* noccur = 3 */
      if (C_integer.Id() == p.Isa.Id()) /* If:3 */{ 
        /* Let:4 */{ 
          var g0121 int  = ToInteger(p).Value
          /* noccur = 1 */
          Result = (ToInteger(x).Value-g0121)
          /* Let-4 */} 
        /* If!3 */}  else if (p.Isa.IsIn(C_list) == CTRUE) /* If:3 */{ 
        Result = F_get_index_table2(a,ToInteger(ToList(x).At(1-1)).Value,ToInteger(ToList(x).At(2-1)).Value)
        } else {
        Result = 1
        /* If-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: get_index @ list<type_expression>(table, any) (throw: false) 
func E_get_index_table1 (a EID,x EID) EID { 
    return EID{C__INT,IVAL(/*(sm for get_index @ list<type_expression>(table, any)= integer)*/ F_get_index_table1(ToTable(OBJ(a)),ANY(x) ))}} 
  
/* {1} OPT.The go function for: get_index(a:table,x:integer,y:integer) [] */
func F_get_index_table2 (a *ClaireTable ,x int,y int) int { 
    // procedure body with s =  
var Result int 
    /* Let:2 */{ 
      var p *ClaireList   = ToList(a.Params)
      /* noccur = 2 */
      Result = (((p.ValuesI()[1-1]*x)+y)-p.ValuesI()[2-1])
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: get_index @ list<type_expression>(table, integer, integer) (throw: false) 
func E_get_index_table2 (a EID,x EID,y EID) EID { 
    return EID{C__INT,IVAL(/*(sm for get_index @ list<type_expression>(table, integer, integer)= integer)*/ F_get_index_table2(ToTable(OBJ(a)),INT(x),INT(y) ))}} 
  
// erase an table means to clean its graph so that it becomes empty.
/* {1} OPT.The go function for: erase(a:table) [] */
func F_erase_table (a *ClaireTable ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var p *ClaireAny   = a.Params
      /* noccur = 2 */
      if (C_integer.Id() == p.Isa.Id()) /* If:3 */{ 
        /* For:4 */{ 
          var i *ClaireAny  
          _ = i
          Result= EID{CFALSE.Id(),0}
          var i_support *ClaireList  
          var i_support_try01275 EID 
          i_support_try01275 = F_enumerate_any(a.Domain.Id())
          /* ERROR PROTECTION INSERTED (i_support-Result) */
          if ErrorIn(i_support_try01275) {Result = i_support_try01275
          } else {
          i_support = ToList(OBJ(i_support_try01275))
          i_len := i_support.Length()
          for i_it := 0; i_it < i_len; i_it++ { 
            i = i_support.At(i_it)
            var void_try6 EID 
            _ = void_try6
            void_try6 = F_CALL(C_nth_equal,ARGS(a.Graph.ToEID(),EID{C__INT,IVAL(F_get_index_table1(a,i))},a.Default.ToEID()))
            /* ERROR PROTECTION INSERTED (void_try6-Result) */
            if ErrorIn(void_try6) {Result = void_try6
            Result = void_try6
            break
            } else {
            }}
            /* loop-5 */} 
          /* For-4 */} 
        /* If!3 */}  else if (p.Isa.IsIn(C_list) == CTRUE) /* If:3 */{ 
        /* For:4 */{ 
          var l *ClaireAny  
          _ = l
          Result= EID{CFALSE.Id(),0}
          var l_support *ClaireList  
          var l_support_try01285 EID 
          l_support_try01285 = F_enumerate_any(a.Domain.Id())
          /* ERROR PROTECTION INSERTED (l_support-Result) */
          if ErrorIn(l_support_try01285) {Result = l_support_try01285
          } else {
          l_support = ToList(OBJ(l_support_try01285))
          l_len := l_support.Length()
          for i_it := 0; i_it < l_len; i_it++ { 
            l = l_support.At(i_it)
            var void_try6 EID 
            _ = void_try6
            void_try6 = F_CALL(C_nth_equal,ARGS(a.Graph.ToEID(),EID{C__INT,IVAL(F_get_index_table2(a,ToInteger(ToList(l).At(1-1)).Value,ToInteger(ToList(l).At(2-1)).Value))},a.Default.ToEID()))
            /* ERROR PROTECTION INSERTED (void_try6-Result) */
            if ErrorIn(void_try6) {Result = void_try6
            Result = void_try6
            break
            } else {
            }}
            /* loop-5 */} 
          /* For-4 */} 
        } else {
        Result = EID{CNIL.Id(),0}
        /* If-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: erase @ table (throw: true) 
func E_erase_table (a EID) EID { 
    return /*(sm for erase @ table= EID)*/ F_erase_table(ToTable(OBJ(a)) )} 
  
// the general case is waiting for the dictionary method like erase(a.graph as dictionary)
// new in v3.2.50 a constructor for building a table dynamically
/* {1} OPT.The go function for: make_table(%domain:type,%range:type,%default:any) [] */
func F_make_table_type (_Zdomain *ClaireType ,_Zrange *ClaireType ,_Zdefault *ClaireAny ) *ClaireTable  { 
    // procedure body with s =  
var Result *ClaireTable  
    /* Let:2 */{ 
      var t *ClaireTable   = ToTable(new(ClaireTable).Is(C_table))
      /* noccur = 11 */
      t.Range = _Zrange
      C_table.Instances = C_table.Instances.AddFast(t.Id())
      t.Domain = _Zdomain
      t.Default = _Zdefault
      t.Params = C_any.Id()
      t.GraphInit()
      Result = t
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: make_table @ type (throw: false) 
func E_make_table_type (_Zdomain EID,_Zrange EID,_Zdefault EID) EID { 
    return EID{/*(sm for make_table @ type= table)*/ F_make_table_type(ToType(OBJ(_Zdomain)),ToType(OBJ(_Zrange)),ANY(_Zdefault) ).Id(),0}} 
  
// Our first table: a debuging tool which stores a list of stopping values
// *********************************************************************
//   Part 3: Demons & relations for the logic modules                  *
// *********************************************************************
// applying a lambda to one argument
/* {1} OPT.The go function for: funcall(self:lambda,x:any) [] */
func F_funcall_lambda1 (self *ClaireLambda ,x *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var start int  = ClEnv.Index
      /* noccur = 2 */
      /* Let:3 */{ 
        var retour int  = ClEnv.Base
        /* noccur = 1 */
        ClEnv.Base= start
        ClEnv.Push(x.ToEID())
        F_stack_add(self.Dimension)
        /* LetE:4 */{ 
          var val EID 
          val = EVAL(self.Body)
          /* ERROR PROTECTION INSERTED (val-Result) */
          if ErrorIn(val) {Result = val
          } else {
          ClEnv.Base= retour
          ClEnv.Index= start
          Result = val}
          /* LetE-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: funcall @ list<type_expression>(lambda, any) (throw: true) 
func E_funcall_lambda1 (self EID,x EID) EID { 
    return /*(sm for funcall @ list<type_expression>(lambda, any)= EID)*/ F_funcall_lambda1(ToLambda(OBJ(self)),ANY(x) )} 
  
// applying a lambda to two argument
/* {1} OPT.The go function for: funcall(self:lambda,x:any,y:any) [] */
func F_funcall_lambda2 (self *ClaireLambda ,x *ClaireAny ,y *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var start int  = ClEnv.Index
      /* noccur = 2 */
      /* Let:3 */{ 
        var retour int  = ClEnv.Base
        /* noccur = 1 */
        ClEnv.Base= start
        ClEnv.Push(x.ToEID())
        ClEnv.Push(y.ToEID())
        F_stack_add(self.Dimension)
        /* LetE:4 */{ 
          var val EID 
          val = EVAL(self.Body)
          /* ERROR PROTECTION INSERTED (val-Result) */
          if ErrorIn(val) {Result = val
          } else {
          ClEnv.Base= retour
          ClEnv.Index= start
          Result = val}
          /* LetE-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: funcall @ list<type_expression>(lambda, any, any) (throw: true) 
func E_funcall_lambda2 (self EID,x EID,y EID) EID { 
    return /*(sm for funcall @ list<type_expression>(lambda, any, any)= EID)*/ F_funcall_lambda2(ToLambda(OBJ(self)),ANY(x),ANY(y) )} 
  
// applying a lambda to two argument
/* {1} OPT.The go function for: funcall(self:lambda,x:any,y:any,z:any) [] */
func F_funcall_lambda3 (self *ClaireLambda ,x *ClaireAny ,y *ClaireAny ,z *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var start int  = ClEnv.Index
      /* noccur = 2 */
      /* Let:3 */{ 
        var retour int  = ClEnv.Base
        /* noccur = 1 */
        ClEnv.Base= start
        ClEnv.Push(x.ToEID())
        ClEnv.Push(y.ToEID())
        ClEnv.Push(z.ToEID())
        F_stack_add(self.Dimension)
        /* LetE:4 */{ 
          var val EID 
          val = EVAL(self.Body)
          /* ERROR PROTECTION INSERTED (val-Result) */
          if ErrorIn(val) {Result = val
          } else {
          ClEnv.Base= retour
          ClEnv.Index= start
          Result = val}
          /* LetE-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: funcall @ list<type_expression>(lambda, any, any, any) (throw: true) 
func E_funcall_lambda3 (self EID,x EID,y EID,z EID) EID { 
    return /*(sm for funcall @ list<type_expression>(lambda, any, any, any)= EID)*/ F_funcall_lambda3(ToLambda(OBJ(self)),
      ANY(x),
      ANY(y),
      ANY(z) )} 
  
// dealing with inverse
/* {1} OPT.The go function for: check_inverse(%r1:any,%r2:any) [] */
func F_check_inverse_any (_Zr1 *ClaireAny ,_Zr2 *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var r1 *ClaireRelation   = ToRelation(_Zr1)
      /* noccur = 8 */
      /* Let:3 */{ 
        var r2 *ClaireRelation   = ToRelation(_Zr2)
        /* noccur = 8 */
        r1.Inverse = r2
        r2.Inverse = r1
        F_final_relation(r1)
        F_final_relation(r2)
        var g0130I *ClaireBoolean  
        /* or:4 */{ 
          var v_or4 *ClaireBoolean  
          
          /* Let:5 */{ 
            var g0131UU *ClaireBoolean  
            /* noccur = 1 */
            /* Let:6 */{ 
              var g0132UU *ClaireType  
              /* noccur = 1 */
              if (r2.Multivalued_ask == CTRUE) /* If:7 */{ 
                g0132UU = F_member_type(r2.Range)
                } else {
                g0132UU = r2.Range
                /* If-7 */} 
              g0131UU = r1.Domain.Included(g0132UU)
              /* Let-6 */} 
            v_or4 = g0131UU.Not
            /* Let-5 */} 
          if (v_or4 == CTRUE) {g0130I = CTRUE
          } else /* or:5 */{ 
            /* Let:6 */{ 
              var g0133UU *ClaireBoolean  
              /* noccur = 1 */
              /* Let:7 */{ 
                var g0134UU *ClaireType  
                /* noccur = 1 */
                if (r1.Multivalued_ask == CTRUE) /* If:8 */{ 
                  g0134UU = F_member_type(r1.Range)
                  } else {
                  g0134UU = r1.Range
                  /* If-8 */} 
                g0133UU = r2.Domain.Included(g0134UU)
                /* Let-7 */} 
              v_or4 = g0133UU.Not
              /* Let-6 */} 
            if (v_or4 == CTRUE) {g0130I = CTRUE
            } else /* or:6 */{ 
              g0130I = CFALSE/* org-6 */} 
            /* org-5 */} 
          /* or-4 */} 
        if (g0130I == CTRUE) /* If:4 */{ 
          Result = ToException(C_general_error.Make(MakeString("[137] ~S and ~S cannot be inverses for one another").Id(),MakeConstantList(r1.Id(),r2.Id()).Id())).Close()
          } else {
          Result = EID{CFALSE.Id(),0}
          /* If-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: check_inverse @ any (throw: true) 
func E_check_inverse_any (_Zr1 EID,_Zr2 EID) EID { 
    return /*(sm for check_inverse @ any= EID)*/ F_check_inverse_any(ANY(_Zr1),ANY(_Zr2) )} 
  
// very useful
/* {1} OPT.The go function for: invert(r:relation,x:any) [] */
func F_invert_relation (r *ClaireRelation ,x *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var r2 *ClaireAny   = F_get_property(C_inverse,ToObject(r.Id()))
      /* noccur = 4 */
      if (C_table.Id() == r2.Isa.Id()) /* If:3 */{ 
        /* Let:4 */{ 
          var g0135 *ClaireTable   = ToTable(r2)
          /* noccur = 2 */
          /* Let:5 */{ 
            var v *ClaireAny  
            /* noccur = 2 */
            var v_try01406 EID 
            v_try01406 = F_nth_table1(g0135,x)
            /* ERROR PROTECTION INSERTED (v-Result) */
            if ErrorIn(v_try01406) {Result = v_try01406
            } else {
            v = ANY(v_try01406)
            if (g0135.Multivalued_ask.Id() != CFALSE.Id()) /* If:6 */{ 
              Result = v.ToEID()
              } else {
              Result = EID{MakeConstantSet(v).Id(),0}
              /* If-6 */} 
            }
            /* Let-5 */} 
          /* Let-4 */} 
        /* If!3 */}  else if (r2.Isa.IsIn(C_property) == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var g0136 *ClaireProperty   = ToProperty(r2)
          /* noccur = 2 */
          /* Let:5 */{ 
            var v *ClaireAny   = F_get_property(g0136,ToObject(x))
            /* noccur = 2 */
            if (g0136.Multivalued_ask.Id() != CFALSE.Id()) /* If:6 */{ 
              Result = v.ToEID()
              } else {
              Result = EID{MakeConstantSet(v).Id(),0}
              /* If-6 */} 
            /* Let-5 */} 
          /* Let-4 */} 
        /* If!3 */}  else if (r.Isa.IsIn(C_property) == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var g0138 *ClaireProperty   = ToProperty(r.Id())
          /* noccur = 5 */
          if (g0138.Multivalued_ask.Id() != CFALSE.Id()) /* If:5 */{ 
            /* Let:6 */{ 
              var z_out *ClaireSet   = ToType(CEMPTY.Id()).EmptySet()
              /* noccur = 2 */
              /* For:7 */{ 
                var z *ClaireAny  
                _ = z
                Result= EID{CFALSE.Id(),0}
                var z_support *ClaireList  
                var z_support_try01418 EID 
                z_support_try01418 = F_enumerate_any(g0138.Domain.Id())
                /* ERROR PROTECTION INSERTED (z_support-Result) */
                if ErrorIn(z_support_try01418) {Result = z_support_try01418
                } else {
                z_support = ToList(OBJ(z_support_try01418))
                z_len := z_support.Length()
                for i_it := 0; i_it < z_len; i_it++ { 
                  z = z_support.At(i_it)
                  if (ToType(F_get_property(g0138,ToObject(z))).Contains(x) == CTRUE) /* If:9 */{ 
                    z_out.AddFast(z)
                    /* If-9 */} 
                  }
                  /* loop-8 */} 
                /* For-7 */} 
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              Result = EID{z_out.Id(),0}
              }
              /* Let-6 */} 
            } else {
            /* Let:6 */{ 
              var z_out *ClaireSet   = ToType(CEMPTY.Id()).EmptySet()
              /* noccur = 2 */
              /* For:7 */{ 
                var z *ClaireAny  
                _ = z
                Result= EID{CFALSE.Id(),0}
                var z_support *ClaireList  
                var z_support_try01428 EID 
                z_support_try01428 = F_enumerate_any(g0138.Domain.Id())
                /* ERROR PROTECTION INSERTED (z_support-Result) */
                if ErrorIn(z_support_try01428) {Result = z_support_try01428
                } else {
                z_support = ToList(OBJ(z_support_try01428))
                z_len := z_support.Length()
                for i_it := 0; i_it < z_len; i_it++ { 
                  z = z_support.At(i_it)
                  if (Equal(F_get_property(g0138,ToObject(z)),x) == CTRUE) /* If:9 */{ 
                    z_out.AddFast(z)
                    /* If-9 */} 
                  }
                  /* loop-8 */} 
                /* For-7 */} 
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              Result = EID{z_out.Id(),0}
              }
              /* Let-6 */} 
            /* If-5 */} 
          /* Let-4 */} 
        /* If!3 */}  else if (C_table.Id() == r.Isa.Id()) /* If:3 */{ 
        /* Let:4 */{ 
          var g0139 *ClaireTable   = ToTable(r.Id())
          /* noccur = 5 */
          if (g0139.Multivalued_ask.Id() != CFALSE.Id()) /* If:5 */{ 
            /* Let:6 */{ 
              var z_out *ClaireSet   = ToType(CEMPTY.Id()).EmptySet()
              /* noccur = 2 */
              /* For:7 */{ 
                var z *ClaireAny  
                _ = z
                Result= EID{CFALSE.Id(),0}
                var z_support *ClaireList  
                var z_support_try01438 EID 
                z_support_try01438 = F_enumerate_any(g0139.Domain.Id())
                /* ERROR PROTECTION INSERTED (z_support-Result) */
                if ErrorIn(z_support_try01438) {Result = z_support_try01438
                } else {
                z_support = ToList(OBJ(z_support_try01438))
                z_len := z_support.Length()
                for i_it := 0; i_it < z_len; i_it++ { 
                  z = z_support.At(i_it)
                  var void_try9 EID 
                  _ = void_try9
                  var g0144I *ClaireBoolean  
                  var g0144I_try01459 EID 
                  /* Let:9 */{ 
                    var g0146UU *ClaireAny  
                    /* noccur = 1 */
                    var g0146UU_try014710 EID 
                    g0146UU_try014710 = F_nth_table1(g0139,z)
                    /* ERROR PROTECTION INSERTED (g0146UU-g0144I_try01459) */
                    if ErrorIn(g0146UU_try014710) {g0144I_try01459 = g0146UU_try014710
                    } else {
                    g0146UU = ANY(g0146UU_try014710)
                    g0144I_try01459 = EID{ToType(g0146UU).Contains(x).Id(),0}
                    }
                    /* Let-9 */} 
                  /* ERROR PROTECTION INSERTED (g0144I-void_try9) */
                  if ErrorIn(g0144I_try01459) {void_try9 = g0144I_try01459
                  } else {
                  g0144I = ToBoolean(OBJ(g0144I_try01459))
                  if (g0144I == CTRUE) /* If:9 */{ 
                    void_try9 = EID{z_out.AddFast(z).Id(),0}
                    } else {
                    void_try9 = EID{CFALSE.Id(),0}
                    /* If-9 */} 
                  }
                  /* ERROR PROTECTION INSERTED (void_try9-Result) */
                  if ErrorIn(void_try9) {Result = void_try9
                  Result = void_try9
                  break
                  } else {
                  }}
                  /* loop-8 */} 
                /* For-7 */} 
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              Result = EID{z_out.Id(),0}
              }
              /* Let-6 */} 
            } else {
            /* Let:6 */{ 
              var z_out *ClaireSet   = ToType(CEMPTY.Id()).EmptySet()
              /* noccur = 2 */
              /* For:7 */{ 
                var z *ClaireAny  
                _ = z
                Result= EID{CFALSE.Id(),0}
                var z_support *ClaireList  
                var z_support_try01488 EID 
                z_support_try01488 = F_enumerate_any(g0139.Domain.Id())
                /* ERROR PROTECTION INSERTED (z_support-Result) */
                if ErrorIn(z_support_try01488) {Result = z_support_try01488
                } else {
                z_support = ToList(OBJ(z_support_try01488))
                z_len := z_support.Length()
                for i_it := 0; i_it < z_len; i_it++ { 
                  z = z_support.At(i_it)
                  var void_try9 EID 
                  _ = void_try9
                  var g0149I *ClaireBoolean  
                  var g0149I_try01509 EID 
                  /* Let:9 */{ 
                    var g0151UU *ClaireAny  
                    /* noccur = 1 */
                    var g0151UU_try015210 EID 
                    g0151UU_try015210 = F_nth_table1(g0139,z)
                    /* ERROR PROTECTION INSERTED (g0151UU-g0149I_try01509) */
                    if ErrorIn(g0151UU_try015210) {g0149I_try01509 = g0151UU_try015210
                    } else {
                    g0151UU = ANY(g0151UU_try015210)
                    g0149I_try01509 = EID{Equal(g0151UU,x).Id(),0}
                    }
                    /* Let-9 */} 
                  /* ERROR PROTECTION INSERTED (g0149I-void_try9) */
                  if ErrorIn(g0149I_try01509) {void_try9 = g0149I_try01509
                  } else {
                  g0149I = ToBoolean(OBJ(g0149I_try01509))
                  if (g0149I == CTRUE) /* If:9 */{ 
                    void_try9 = EID{z_out.AddFast(z).Id(),0}
                    } else {
                    void_try9 = EID{CFALSE.Id(),0}
                    /* If-9 */} 
                  }
                  /* ERROR PROTECTION INSERTED (void_try9-Result) */
                  if ErrorIn(void_try9) {Result = void_try9
                  Result = void_try9
                  break
                  } else {
                  }}
                  /* loop-8 */} 
                /* For-7 */} 
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              Result = EID{z_out.Id(),0}
              }
              /* Let-6 */} 
            /* If-5 */} 
          /* Let-4 */} 
        } else {
        Result = EID{CFALSE.Id(),0}
        /* If-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: invert @ relation (throw: true) 
func E_invert_relation (r EID,x EID) EID { 
    return /*(sm for invert @ relation= EID)*/ F_invert_relation(ToRelation(OBJ(r)),ANY(x) )} 
  
// same: two useful methods that are used often
/* {1} OPT.The go function for: domain!(x:restriction) [] */
func F_domain_I_restriction (x *ClaireRestriction ) *ClaireClass  { 
    // use function body compiling 
return  ToTypeExpression(x.Domain.ValuesO()[1-1]).Class_I()
    } 
  
// The EID go function for: domain! @ restriction (throw: false) 
func E_domain_I_restriction (x EID) EID { 
    return EID{/*(sm for domain! @ restriction= class)*/ F_domain_I_restriction(ToRestriction(OBJ(x)) ).Id(),0}} 
  
/* {1} OPT.The go function for: methods(d:class,r:class) [] */
func F_methods_class (d *ClaireClass ,r *ClaireClass ) *ClaireSet  { 
    // procedure body with s =  
var Result *ClaireSet  
    /* Let:2 */{ 
      var m_out *ClaireSet   = ToType(CEMPTY.Id()).EmptySet()
      /* noccur = 2 */
      /* For:3 */{ 
        var m *ClaireAny  
        _ = m
        var m_support *ClaireList  
        m_support = C_method.Instances
        m_len := m_support.Length()
        for i_it := 0; i_it < m_len; i_it++ { 
          m = m_support.At(i_it)
          if ((ToType(ToRestriction(m).Domain.ValuesO()[1-1]).Included(ToType(d.Id())) == CTRUE) && 
              (ToRestriction(m).Range.Included(ToType(r.Id())) == CTRUE)) /* If:5 */{ 
            m_out.AddFast(m)
            /* If-5 */} 
          /* loop-4 */} 
        /* For-3 */} 
      Result = m_out
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: methods @ class (throw: false) 
func E_methods_class (d EID,r EID) EID { 
    return EID{/*(sm for methods @ class= set)*/ F_methods_class(ToClass(OBJ(d)),ToClass(OBJ(r)) ).Id(),0}} 
  
// sets the reified flag
/* {1} OPT.The go function for: reify(l:listargs) [] */
func F_reify_listargs (l *ClaireList )  { 
    // procedure body with s =  
/* For:2 */{ 
      var p *ClaireAny  
      _ = p
      var p_support *ClaireList  
      p_support = ToList(l.Id())
      p_len := p_support.Length()
      for i_it := 0; i_it < p_len; i_it++ { 
        p = p_support.At(i_it)
        if (p.Isa.IsIn(C_property) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var g0153 *ClaireProperty   = ToProperty(p)
            /* noccur = 2 */
            g0153.Reified = CTRUE
            /* Let-5 */} 
          /* If-4 */} 
        /* loop-3 */} 
      /* For-2 */} 
    } 
  
// The EID go function for: reify @ listargs (throw: false) 
func E_reify_listargs (l EID) EID { 
    /*(sm for reify @ listargs= void)*/ F_reify_listargs(ToList(OBJ(l)) )
    return EVOID} 
  
// *********************************************************************
// *   Part 4: Basics of Exceptions                                    *
// *********************************************************************
// args :: property(open = 0)
// value :: property() - defined in kernel
// a generic error that is produced by the error(" ....") instruction
/* {1} OPT.The go function for: self_print(self:general_error) [] */
func (self *GeneralError ) SelfPrint () EID { 
    var Result EID 
    PRINC("**** An error has occurred.\n")
    Result = F_format_string(ToString(self.Cause),ToList(self.Arg))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("\n")
    Result = EVOID
    }
    return Result} 
  
// The EID go function for: self_print @ general_error (throw: true) 
func E_self_print_general_error_Core (self EID) EID { 
    return /*(sm for self_print @ general_error= EID)*/ ToGeneralError(OBJ(self)).SelfPrint( )} 
  
// a read_slot error is produced when an unknown value is found
/* {1} OPT.The go function for: self_print(self:read_slot_error) [] */
func (self *ReadSlotError ) SelfPrint () EID { 
    var Result EID 
    PRINC("****[138] The value of ")
    Result = F_CALL(C_print,ARGS(self.Wrong.ToEID()))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("(")
    Result = F_CALL(C_print,ARGS(self.Arg.ToEID()))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(") is unknown")
    Result = EVOID
    }}
    return Result} 
  
// The EID go function for: self_print @ read_slot_error (throw: true) 
func E_self_print_read_slot_error_Core (self EID) EID { 
    return /*(sm for self_print @ read_slot_error= EID)*/ ToReadSlotError(OBJ(self)).SelfPrint( )} 
  
// range errors
/* {1} OPT.The go function for: self_print(self:range_error) [] */
func (self *RangeError ) SelfPrint () EID { 
    var Result EID 
    PRINC("****[139] ")
    Result = F_CALL(C_print,ARGS(self.Cause.ToEID()))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(": range error, ")
    Result = F_CALL(C_print,ARGS(self.Arg.ToEID()))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(" does not belong? to ")
    Result = F_CALL(C_print,ARGS(self.Wrong.ToEID()))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(".\n")
    Result = EVOID
    }}}
    return Result} 
  
// The EID go function for: self_print @ range_error (throw: true) 
func E_self_print_range_error_Core (self EID) EID { 
    return /*(sm for self_print @ range_error= EID)*/ ToRangeError(OBJ(self)).SelfPrint( )} 
  
// selector errors
/* {1} OPT.The go function for: self_print(self:selector_error) [] */
func (self *SelectorError ) SelfPrint () EID { 
    var Result EID 
    /* Let:2 */{ 
      var p *ClaireAny   = self.Selector
      /* noccur = 3 */
      if (F_boolean_I_any(ToProperty(p).Restrictions.Id()).Id() != CTRUE.Id()) /* If:3 */{ 
        PRINC("[140] The property ")
        Result = F_CALL(C_print,ARGS(p.ToEID()))
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC(" is not defined (was applied to ")
        Result = F_CALL(C_print,ARGS(self.Arg.ToEID()))
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC(").\n")
        Result = EVOID
        }}
        } else {
        PRINC("****[141] ")
        Result = F_CALL(C_print,ARGS(self.Arg.ToEID()))
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC(" is a wrong arg list for ")
        Result = F_CALL(C_print,ARGS(p.ToEID()))
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC(".\n")
        Result = EVOID
        }}
        /* If-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: self_print @ selector_error (throw: true) 
func E_self_print_selector_error_Core (self EID) EID { 
    return /*(sm for self_print @ selector_error= EID)*/ ToSelectorError(OBJ(self)).SelfPrint( )} 
  
// produced by a return (usually trapped)
/* {1} OPT.The go function for: self_print(self:return_error) [] */
func (self *ReturnError ) SelfPrint ()  { 
    // procedure body with s =  
PRINC("****[142] return called outside of a loop (for or while).")
    } 
  
// The EID go function for: self_print @ return_error (throw: false) 
func E_self_print_return_error_Core (self EID) EID { 
    /*(sm for self_print @ return_error= void)*/ ToReturnError(OBJ(self)).SelfPrint( )
    return EVOID} 
  
// interpretation of all the error codes
/* {1} OPT.The go function for: self_print(self:system_error) [] */
func F_self_print_system_error_Core (self *ClaireSystemError ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var n int  = self.Index
      /* noccur = 42 */
      PRINC("**** An internal error [")
      F_princ_integer(n)
      PRINC("] has occured:\n")
      /* Let:3 */{ 
        var g0154UU *ClaireString  
        /* noccur = 1 */
        if (n == 1) /* If:4 */{ 
          g0154UU = MakeString("Slot Access : ~S(~S) is unknown")
          /* If!4 */}  else if (n == 2) /* If:4 */{ 
          g0154UU = MakeString("dynamic allocation, too large for available memory (~S)")
          /* If!4 */}  else if (n == 3) /* If:4 */{ 
          g0154UU = MakeString("object allocation, too large for available memory (~S)")
          /* If!4 */}  else if (n == 5) /* If:4 */{ 
          g0154UU = MakeString("nth[~S] outside of scope for ~S")
          /* If!4 */}  else if (n == 6) /* If:4 */{ 
          g0154UU = MakeString("invalid slot access: ~S does not have ~S")
          /* If!4 */}  else if (n == 7) /* If:4 */{ 
          g0154UU = MakeString("Skip applied on ~S with a negative argument ~S")
          /* If!4 */}  else if (n == 8) /* If:4 */{ 
          g0154UU = MakeString("List operation: cdr(()) is undefined")
          /* If!4 */}  else if (n == 9) /* If:4 */{ 
          g0154UU = MakeString("String buffer is full: ~S")
          /* If!4 */}  else if (n == 10) /* If:4 */{ 
          g0154UU = MakeString("Cannot create an imported entity from NULL reference")
          /* If!4 */}  else if (n == 11) /* If:4 */{ 
          g0154UU = MakeString("nth_string[~S]: string too short~S")
          /* If!4 */}  else if (n == 12) /* If:4 */{ 
          g0154UU = MakeString("Symbol Table table full")
          /* If!4 */}  else if (n == 13) /* If:4 */{ 
          g0154UU = MakeString("Cannot create a subclass for ~S [~A]")
          /* If!4 */}  else if (n == 16) /* If:4 */{ 
          g0154UU = MakeString("Temporary output string buffer too small")
          /* If!4 */}  else if (n == 17) /* If:4 */{ 
          g0154UU = MakeString("Bag Type Error: ~S does not belong to type ~S")
          /* If!4 */}  else if (n == 18) /* If:4 */{ 
          g0154UU = MakeString("definition of ~S is in conflict with an object from ~S")
          /* If!4 */}  else if (n == 19) /* If:4 */{ 
          g0154UU = MakeString("Integer overflow")
          /* If!4 */}  else if (n == 20) /* If:4 */{ 
          g0154UU = MakeString("Integer arithmetic: division/modulo of ~A by 0")
          /* If!4 */}  else if (n == 21) /* If:4 */{ 
          g0154UU = MakeString("Integer to character: ~S is a wrong value")
          /* If!4 */}  else if (n == 22) /* If:4 */{ 
          g0154UU = MakeString("Cannote create a string with negative length ~S")
          /* If!4 */}  else if (n == 23) /* If:4 */{ 
          g0154UU = MakeString("Not enough memory to instal claire")
          /* If!4 */}  else if (n == 24) /* If:4 */{ 
          g0154UU = MakeString("execution stack is full [~A]")
          /* If!4 */}  else if (n == 26) /* If:4 */{ 
          g0154UU = MakeString("Wrong usage of time counter [~A]")
          /* If!4 */}  else if (n == 27) /* If:4 */{ 
          g0154UU = MakeString("internal garbage protection stack overflow")
          /* If!4 */}  else if (n == 28) /* If:4 */{ 
          g0154UU = MakeString("the multivalued status of ~S is not compatible with ~S")
          /* If!4 */}  else if (n == 29) /* If:4 */{ 
          g0154UU = MakeString("There is no module ~S")
          /* If!4 */}  else if (n == 30) /* If:4 */{ 
          g0154UU = MakeString("Attempt to read a private symbol ~S")
          /* If!4 */}  else if (n == 31) /* If:4 */{ 
          g0154UU = MakeString("External function not compiled yet")
          /* If!4 */}  else if (n == 32) /* If:4 */{ 
          g0154UU = MakeString("Too many arguments (~S) for function ~S")
          /* If!4 */}  else if (n == 33) /* If:4 */{ 
          g0154UU = MakeString("Exception handling: stack overflow")
          /* If!4 */}  else if (n == 34) /* If:4 */{ 
          g0154UU = MakeString("User interrupt: EXECUTION ABORTED")
          /* If!4 */}  else if (n == 35) /* If:4 */{ 
          g0154UU = MakeString("reading char '~S': wrong char: ~S")
          /* If!4 */}  else if (n == 36) /* If:4 */{ 
          g0154UU = MakeString("cannot open file ~A")
          /* If!4 */}  else if (n == 37) /* If:4 */{ 
          g0154UU = MakeString("world stack is full")
          /* If!4 */}  else if (n == 38) /* If:4 */{ 
          g0154UU = MakeString("Undefined access to ~S")
          /* If!4 */}  else if (n == 39) /* If:4 */{ 
          g0154UU = MakeString("cannot convert ~S to an integer")
          /* If!4 */}  else if (n == 40) /* If:4 */{ 
          g0154UU = MakeString("integer multiplication overflow with ~S and ~S")
          /* If!4 */}  else if (n == 41) /* If:4 */{ 
          g0154UU = MakeString("wrong NTH access on ~S and ~S")
          /* If!4 */}  else if (n == 42) /* If:4 */{ 
          g0154UU = MakeString("Wrong array[~S] init value: ~S")
          /* If!4 */}  else if (n == 43) /* If:4 */{ 
          g0154UU = MakeString("Defeasible addition on list ~S requires pre-allocation (size ~S)")
          /* If!4 */}  else if (n == 50) /* If:4 */{ 
          g0154UU = MakeString("C++ imported error (~S) : ~S")
          } else {
          self.Value = MakeInteger(n).Id()
          g0154UU = MakeString("What the hell is this ! [code: ~S^]")
          /* If-4 */} 
        Result = F_format_string(g0154UU,MakeConstantList(self.Value,self.Arg))
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: self_print @ system_error (throw: true) 
func E_self_print_system_error_Core (self EID) EID { 
    return /*(sm for self_print @ system_error= EID)*/ F_self_print_system_error_Core(ToSystemError(OBJ(self)) )} 
  
// contradictions are nice exceptions
/* {1} OPT.The go function for: self_print(x:contradiction) [] */
func (x *Contradiction ) SelfPrint ()  { 
    // procedure body with s =  
PRINC("A contradiction has occured.")
    } 
  
// The EID go function for: self_print @ contradiction (throw: false) 
func E_self_print_contradiction_Core (x EID) EID { 
    /*(sm for self_print @ contradiction= void)*/ ToContradiction(OBJ(x)).SelfPrint( )
    return EVOID} 
  
// the format method is used to print error messages (similar to a printf)
/* {1} OPT.The go function for: format(self:string,larg:list) [] */
func F_format_string (self *ClaireString ,larg *ClaireList ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var s *ClaireString   = self
      /* noccur = 8 */
      /* Let:3 */{ 
        var n int  = F_get_string(s,'~')
        /* noccur = 6 */
        /* Let:4 */{ 
          var l *ClaireList   = larg.Copy()
          /* noccur = 4 */
          Result= EID{CFALSE.Id(),0}
          for (n != 0) /* while:5 */{ 
            var void_try6 EID 
            _ = void_try6
            /* Let:6 */{ 
              var m rune  = s.At((n+1))
              /* noccur = 4 */
              if (n > 1) /* If:7 */{ 
                F_princ_string(F_substring_string(s,1,(n-1)))
                /* If-7 */} 
              if ('A' == m) /* If:7 */{ 
                /* Let:8 */{ 
                  var g0155UU *ClaireAny  
                  /* noccur = 1 */
                  var g0155UU_try01569 EID 
                  g0155UU_try01569 = F_car_list(l)
                  /* ERROR PROTECTION INSERTED (g0155UU-void_try6) */
                  if ErrorIn(g0155UU_try01569) {void_try6 = g0155UU_try01569
                  } else {
                  g0155UU = ANY(g0155UU_try01569)
                  void_try6 = F_CALL(C_princ,ARGS(g0155UU.ToEID()))
                  }
                  /* Let-8 */} 
                /* If!7 */}  else if ('S' == m) /* If:7 */{ 
                /* Let:8 */{ 
                  var g0157UU *ClaireAny  
                  /* noccur = 1 */
                  var g0157UU_try01589 EID 
                  g0157UU_try01589 = F_car_list(l)
                  /* ERROR PROTECTION INSERTED (g0157UU-void_try6) */
                  if ErrorIn(g0157UU_try01589) {void_try6 = g0157UU_try01589
                  } else {
                  g0157UU = ANY(g0157UU_try01589)
                  void_try6 = F_CALL(C_print,ARGS(g0157UU.ToEID()))
                  }
                  /* Let-8 */} 
                /* If!7 */}  else if ('I' == m) /* If:7 */{ 
                void_try6 = ToException(C_general_error.Make(MakeString("[143] ~I not allowed in format").Id(),MakeConstantList(CNULL).Id())).Close()
                } else {
                void_try6 = EID{CFALSE.Id(),0}
                /* If-7 */} 
              /* ERROR PROTECTION INSERTED (void_try6-void_try6) */
              if ErrorIn(void_try6) {Result = void_try6
              break
              } else {
              if (m != '%') /* If:7 */{ 
                l = l.Skip(1)
                /* If-7 */} 
              s = F_substring_string(s,(n+2),1000)
              n = F_get_string(s,'~')
              void_try6 = EID{C__INT,IVAL(n)}
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (void_try6-Result) */
            if ErrorIn(void_try6) {Result = void_try6
            Result = void_try6
            break
            } else {
            /* while-5 */} 
          }
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          if (F_length_string(s) > 0) /* If:5 */{ 
            F_princ_string(s)
            Result = EVOID
            } else {
            Result = EID{CFALSE.Id(),0}
            /* If-5 */} 
          }
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: format @ string (throw: true) 
func E_format_string (self EID,larg EID) EID { 
    return /*(sm for format @ string= EID)*/ F_format_string(ToString(OBJ(self)),ToList(OBJ(larg)) )} 
  
// special version that prints in the trace port
/* {1} OPT.The go function for: tformat(self:string,i:integer,l:list) [] */
func F_tformat_string (self *ClaireString ,i int,l *ClaireList ) EID { 
    var Result EID 
    if (i <= ClEnv.Verbose) /* If:2 */{ 
      /* Let:3 */{ 
        var p *ClairePort   = ClEnv.Ctrace.UseAsOutput()
        /* noccur = 1 */
        Result = F_format_string(self,l)
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        Result = p.UseAsOutput().ToEID()
        }
        /* Let-3 */} 
      } else {
      Result = EID{CFALSE.Id(),0}
      /* If-2 */} 
    return Result} 
  
// The EID go function for: tformat @ string (throw: true) 
func E_tformat_string (self EID,i EID,l EID) EID { 
    return /*(sm for tformat @ string= EID)*/ F_tformat_string(ToString(OBJ(self)),INT(i),ToList(OBJ(l)) )} 
  
// printing a bag without ( ) separate between sets and lists in CLAIRE4
/* {1} OPT.The go function for: princ(s:list) [] */
func F_princ_list (s *ClaireList ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var f *ClaireBoolean   = CTRUE
      /* noccur = 2 */
      /* For:3 */{ 
        var x *ClaireAny  
        _ = x
        Result= EID{CFALSE.Id(),0}
        var x_support *ClaireList  
        x_support = s
        x_len := x_support.Length()
        for i_it := 0; i_it < x_len; i_it++ { 
          x = x_support.At(i_it)
          var void_try5 EID 
          _ = void_try5
          { 
          if (f == CTRUE) /* If:5 */{ 
            f = CFALSE
            } else {
            PRINC(",")
            /* If-5 */} 
          void_try5 = F_CALL(C_print,ARGS(x.ToEID()))
          /* ERROR PROTECTION INSERTED (void_try5-void_try5) */
          if ErrorIn(void_try5) {Result = void_try5
          break
          } else {
          }
          }
          /* loop-4 */} 
        /* For-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: princ @ list (throw: true) 
func E_princ_list (s EID) EID { 
    return /*(sm for princ @ list= EID)*/ F_princ_list(ToList(OBJ(s)) )} 
  
/* {1} OPT.The go function for: princ(s:set) [] */
func F_princ_set (s *ClaireSet ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var f *ClaireBoolean   = CTRUE
      /* noccur = 2 */
      /* For:3 */{ 
        var x *ClaireAny  
        _ = x
        Result= EID{CFALSE.Id(),0}
        for _,x = range(s.Values)/* loop:4 */{ 
          var void_try5 EID 
          _ = void_try5
          { 
          if (f == CTRUE) /* If:5 */{ 
            f = CFALSE
            } else {
            PRINC(",")
            /* If-5 */} 
          void_try5 = F_CALL(C_print,ARGS(x.ToEID()))
          /* ERROR PROTECTION INSERTED (void_try5-void_try5) */
          if ErrorIn(void_try5) {Result = void_try5
          break
          } else {
          }
          }
          /* loop-4 */} 
        /* For-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: princ @ set (throw: true) 
func E_princ_set (s EID) EID { 
    return /*(sm for princ @ set= EID)*/ F_princ_set(ToSet(OBJ(s)) )} 
  
// a global variable is a named object with a special evaluation
// NOTE: we need to refine the scheme for global constants !
// GV are defeasible
/* {1} OPT.The go function for: close(self:global_variable) [] */
func (self *GlobalVariable ) Close () EID { 
    var Result EID 
    if ((self.Value != CNULL) && 
        ((ToType(C_set.Id()).Contains(self.Range.Id()) != CTRUE) && 
            (self.Range.Contains(self.Value) != CTRUE))) /* If:2 */{ 
      Result = ToException(C_range_error.Make(self.Value,self.Id(),self.Range.Id())).Close()
      } else {
      Result = EID{CFALSE.Id(),0}
      /* If-2 */} 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    Result = EID{self.Id(),0}
    }
    return Result} 
  
// The EID go function for: close @ global_variable (throw: true) 
func E_close_global_variable (self EID) EID { 
    return /*(sm for close @ global_variable= EID)*/ ToGlobalVariable(OBJ(self)).Close( )} 
  
/* {1} OPT.The go function for: self_eval(self:global_variable) [] */
func (self *GlobalVariable ) SelfEval () EID { 
    var Result EID 
    Result = self.Value.ToEID()
    return Result} 
  
// The EID go function for: self_eval @ global_variable (throw: true) 
func E_self_eval_global_variable (self EID) EID { 
    return /*(sm for self_eval @ global_variable= EID)*/ ToGlobalVariable(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: global_variable 
func EVAL_global_variable (x *ClaireAny) EID { 
     return ToGlobalVariable(x).SelfEval()} 
  
// -> moved from pretty.cl
// we create a spcial contraidiction that we shall reuse
// how to use it
/* {1} OPT.The go function for: contradiction!(_CL_obj:void) [] */
func F_contradiction_I_void () EID { 
    var Result EID 
    Result = ToContradiction(C_contradiction_occurs.Value).Close()
    return Result} 
  
// The EID go function for: contradiction! @ void (throw: true) 
func E_contradiction_I_void (_CL_obj EID) EID { 
    return /*(sm for contradiction! @ void= EID)*/ F_contradiction_I_void( )} 
  
// v0.01
// end of file