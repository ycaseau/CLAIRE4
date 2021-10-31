/***** CLAIRE Compilation of file /Users/ycaseau/Dropbox/src/clairev4.0/src/meta/inspect.cl 
         [version 4.0.02 / safety 5] Saturday 10-30-2021 *****/

package Reader
import (_ "fmt"
	. "Kernel"
	"Core"
	"Language"
)

//-------- dumb function to prevent import errors --------
func import_g0450() { 
    _ = Core.It
    _ = Language.It
    } 
  
  
//+-------------------------------------------------------------+
//| CLAIRE                                                      |
//| inspect.cl                                                  |
//| Copyright (C) 1994 - 2021 Yves Caseau. All Rights Reserved  |
//| cf. copyright info in file object.cl: about()               |
//+-------------------------------------------------------------+
// --------------------------------------------------------------------
// this file contains the CLAIRE run-time tools: toplevel, inspect, trace & debug
// --------------------------------------------------------------------
// *********************************************************************
// * Contents                                                          *
// *      Part 1: Top Level                                            *
// *      Part 2: Inspection                                           *
// *      Part 3: Trace                                                *
// *      Part 4: Debugger                                             *
// *      Part 5: Measures & Profiler                                  *
// *********************************************************************
// a useful global variable *last*
// v3.2.14 cleaner :-) : current debug stack top
// v0.01 stop the ... !
// this is the interface with the system - used to manage breakpoints
// *********************************************************************
// *      Part 1: Toplevel                                             *
// *********************************************************************
// we use six global variables (used to be C++)
// 1 : regular, 2: debug; 3: inspect
// for debug loop, store the stack context
// this is the classical print(eval(read)) LISP top level :)
// error are caught
/* {1} OPT.The go function for: top_level(r:meta_reader) [] */
func (r *MetaReader ) TopLevel () EID { 
    var Result EID 
    /* Let:2 */{ 
      var res *ClaireAny   = MakeInteger(0).Id()
      /* noccur = 11 */
      Result= EID{CFALSE.Id(),0}
      for (res != C_q.Id()) /* while:3 */{ 
        var void_try4 EID 
        _ = void_try4
        { 
        /* Let:4 */{ 
          var g0451UU *ClaireString  
          /* noccur = 1 */
          if (ToInteger(C_Reader_TopLevelMode.Value).Value == 1) /* If:5 */{ 
            g0451UU = ClEnv.Module_I.Name.String_I()
            /* If!5 */}  else if (ToInteger(C_Reader_TopLevelMode.Value).Value == 2) /* If:5 */{ 
            g0451UU = MakeString("debug")
            } else {
            g0451UU = MakeString("inspect")
            /* If-5 */} 
          F_princ_string(g0451UU)
          /* Let-4 */} 
        PRINC("> ")
        h_index := ClEnv.Index /* Handle */
        h_base := ClEnv.Base
        r.Toplevel = CTRUE
        if (ClEnv.CountCall > 0) /* If:4 */{ 
          ClEnv.CountCall = 1
          /* If-4 */} 
        var res_try04524 EID 
        res_try04524 = r.Nextunit()
        /* ERROR PROTECTION INSERTED (res-void_try4) */
        if ErrorIn(res_try04524) {void_try4 = res_try04524
        } else {
        res = ANY(res_try04524)
        void_try4 = res.ToEID()
        if (ToInteger(C_Reader_TopLevelMode.Value).Value == 1) /* If:4 */{ 
          ClEnv.Index = 20
          /* If-4 */} 
        if ((ToInteger(C_Reader_TopLevelMode.Value).Value == 3) && 
            (res != C_q.Id())) /* If:4 */{ 
          void_try4 = F_inspect_loop_any(res,ToList(C_Reader_InspectStack.Value))
          } else {
          if (ToInteger(C_Reader_TopLevelMode.Value).Value == 1) /* If:5 */{ 
            PRINC("eval[")
            /* Let:6 */{ 
              var g0453UU *ClaireAny  
              /* noccur = 1 */
              g0453UU = MakeInteger((ToInteger(C_Reader_TopCount.Value).Value+1)).Id()
              C_Reader_TopCount.Value = g0453UU
              void_try4 = Core.F_print_any(g0453UU)
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (void_try4-void_try4) */
            if !ErrorIn(void_try4) {
            PRINC("]> ")
            void_try4 = EVOID
            }
            } else {
            PRINC("> ")
            void_try4 = EVOID
            /* If-5 */} 
          /* ERROR PROTECTION INSERTED (void_try4-void_try4) */
          if !ErrorIn(void_try4) {
          var res_try04545 EID 
          res_try04545 = EVAL(res)
          /* ERROR PROTECTION INSERTED (res-void_try4) */
          if ErrorIn(res_try04545) {void_try4 = res_try04545
          } else {
          res = ANY(res_try04545)
          void_try4 = res.ToEID()
          if (res != C_q.Id()) /* If:5 */{ 
            void_try4 = Core.F_CALL(C_print,ARGS(res.ToEID()))
            /* ERROR PROTECTION INSERTED (void_try4-void_try4) */
            if !ErrorIn(void_try4) {
            PRINC("\n")
            void_try4 = EVOID
            }
            } else {
            void_try4 = EID{CFALSE.Id(),0}
            /* If-5 */} 
          }}
          /* If-4 */} 
        }
        if ErrorIn(void_try4){ 
          /* s=EID */ClEnv.Index = h_index
          ClEnv.Base = h_base
          var g0455I *ClaireBoolean  
          /* Let:5 */{ 
            var e *ClaireException   = ClEnv.Exception_I
            /* noccur = 2 */
            g0455I = MakeBoolean((e.Isa.Id() == C_system_error.Id()) && (ToSystemError(e.Id()).Index == -1))
            /* Let-5 */} 
          if (g0455I == CTRUE) /* If:5 */{ 
            C_Reader_TopLevelMode.Value = MakeInteger(1).Id()
            res = C_q.Id()
            void_try4 = res.ToEID()
            } else {
            r.RestoreState()
            if (r.External.Value != MakeString("toplevel").Value) /* If:6 */{ 
              PRINC("---- file: ")
              F_princ_string(r.External)
              PRINC(", line: ")
              F_princ_integer(ClEnv.NLine)
              PRINC("\n")
              /* If-6 */} 
            void_try4 = F_debug_if_possible_void()
            /* ERROR PROTECTION INSERTED (void_try4-void_try4) */
            if ErrorIn(void_try4) {Result = void_try4
            break
            } else {
            PRINC("\n")
            void_try4 = EVOID
            }
            /* If-5 */} 
          } 
        /* ERROR PROTECTION INSERTED (void_try4-void_try4) */
        if ErrorIn(void_try4) {Result = void_try4
        break
        } else {
        if ((ToInteger(C_Reader_TopLevelMode.Value).Value != 1) && 
            (res == C_q.Id())) /* If:4 */{ 
          if (ToInteger(C_Reader_TopLevelMode.Value).Value == 2) /* If:5 */{ 
            ClEnv.Index = ToInteger(C_Reader_TopIndex.Value).Value
            ClEnv.Base = ToInteger(C_Reader_TopBase.Value).Value
            ClEnv.Trace_I = 1
            PRINC("--------- quit debug base:")
            F_princ_integer(ToInteger(C_Reader_TopBase.Value).Value)
            PRINC(" index:")
            F_princ_integer(ToInteger(C_Reader_TopIndex.Value).Value)
            PRINC(" debug:")
            F_princ_integer(ToInteger(C_Reader_TopDebug.Value).Value)
            PRINC("\n")
            ClEnv.Debug_I = ToInteger(C_Reader_TopDebug.Value).Value
            /* If-5 */} 
          res = CNULL
          C_Reader_TopLevelMode.Value = MakeInteger(1).Id()
          /* If-4 */} 
        }
        /* while-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: top_level @ meta_reader (throw: true) 
func E_Reader_top_level_meta_reader (r EID) EID { 
    return /*(sm for top_level @ meta_reader= EID)*/ ToMetaReader(OBJ(r)).TopLevel( )} 
  
//        exit(1)) ]
// start a debug loop - aha 
/* {1} OPT.The go function for: debugLoop(r:meta_reader) [] */
func (r *MetaReader ) DebugLoop ()  { 
    // procedure body with s =  
C_Reader_TopDebug.Value = MakeInteger(0).Id()
    C_Reader_TopBase.Value = MakeInteger(ClEnv.Base).Id()
    C_Reader_TopIndex.Value = MakeInteger(ClEnv.Index).Id()
    r.Toplevel = CTRUE
    C_Reader_TopLevelMode.Value = MakeInteger(2).Id()
    PRINC("--------------- Debug -------------------\n")
    } 
  
// The EID go function for: debugLoop @ meta_reader (throw: false) 
func E_Reader_debugLoop_meta_reader (r EID) EID { 
    /*(sm for debugLoop @ meta_reader= void)*/ ToMetaReader(OBJ(r)).DebugLoop( )
    return EVOID} 
  
// starts an inspector  on a list
/* {1} OPT.The go function for: inspect_system(l:list) [] */
func F_Reader_inspect_system_list2 (l *ClaireList )  { 
    // procedure body with s =  
C_Reader_InspectStack.Value = l.Id()
    if (ToInteger(C_Reader_TopLevelMode.Value).Value == 2) /* If:2 */{ 
      ClEnv.Trace_I = 1
      /* If-2 */} 
    C_Reader_TopLevelMode.Value = MakeInteger(3).Id()
    } 
  
// The EID go function for: inspect_system @ list (throw: false) 
func E_Reader_inspect_system_list2 (l EID) EID { 
    /*(sm for inspect_system @ list= void)*/ F_Reader_inspect_system_list2(ToList(OBJ(l)) )
    return EVOID} 
  
// INSPECT   
// simple main (to be enriched later)
/* {1} OPT.The go function for: simple_main(_CL_obj:void) [] */
func F_Reader_simple_main_void () EID { 
    var Result EID 
    /* Let:2 */{ 
      var _Zinit_ask *ClaireBoolean   = CTRUE
      /* noccur = 4 */
      /* Let:3 */{ 
        var l *ClaireList   = ClEnv.Params.Copy()
        /* noccur = 22 */
        h_index := ClEnv.Index /* Handle */
        h_base := ClEnv.Base
        Result= EID{CFALSE.Id(),0}
        for (l.Length() != 0) /* while:4 */{ 
          var void_try5 EID 
          _ = void_try5
          if (ToString(l.ValuesO()[1-1]).Value == MakeString("-s").Value) /* If:5 */{ 
            if (l.Length() >= 2) /* If:6 */{ 
              l = l.Skip(2)
              void_try5 = EID{l.Id(),0}
              } else {
              void_try5 = ToException(Core.C_general_error.Make(MakeString("option: -s <s1> <s2>").Id(),CNIL.Id())).Close()
              /* If-6 */} 
            /* If!5 */}  else if (ToString(l.ValuesO()[1-1]).Value == MakeString("-f").Value) /* If:5 */{ 
            if (l.Length() >= 2) /* If:6 */{ 
              void_try5 = F_load_string(ToString(l.ValuesO()[2-1]))
              /* ERROR PROTECTION INSERTED (void_try5-void_try5) */
              if ErrorIn(void_try5) {Result = void_try5
              break
              } else {
              l = l.Skip(2)
              void_try5 = EID{l.Id(),0}
              }
              } else {
              void_try5 = ToException(Core.C_general_error.Make(MakeString("option: -f <filename>").Id(),CNIL.Id())).Close()
              /* If-6 */} 
            /* If!5 */}  else if (ToString(l.ValuesO()[1-1]).Value == MakeString("-m").Value) /* If:5 */{ 
            if (l.Length() >= 2) /* If:6 */{ 
              if (_Zinit_ask == CTRUE) /* If:7 */{ 
                void_try5 = F_load_string(MakeString("init"))
                /* ERROR PROTECTION INSERTED (void_try5-void_try5) */
                if ErrorIn(void_try5) {Result = void_try5
                break
                } else {
                _Zinit_ask = CFALSE
                void_try5 = EID{_Zinit_ask.Id(),0}
                }
                } else {
                void_try5 = EID{CFALSE.Id(),0}
                /* If-7 */} 
              /* ERROR PROTECTION INSERTED (void_try5-void_try5) */
              if ErrorIn(void_try5) {Result = void_try5
              break
              } else {
              /* Let:7 */{ 
                var m *ClaireAny   = F_value_string(ToString(l.ValuesO()[2-1]))
                /* noccur = 3 */
                if (m.Isa.IsIn(C_module) != CTRUE) /* If:8 */{ 
                  void_try5 = ToException(Core.C_general_error.Make(MakeString("~S is not a module").Id(),MakeConstantList(l.ValuesO()[2-1]).Id())).Close()
                  } else {
                  void_try5 = EID{CFALSE.Id(),0}
                  /* If-8 */} 
                /* ERROR PROTECTION INSERTED (void_try5-void_try5) */
                if ErrorIn(void_try5) {Result = void_try5
                break
                } else {
                void_try5 = Core.F_CALL(C_load,ARGS(m.ToEID()))
                /* ERROR PROTECTION INSERTED (void_try5-void_try5) */
                if ErrorIn(void_try5) {Result = void_try5
                break
                } else {
                ToModule(m).Begin()
                l = l.Skip(2)
                void_try5 = EID{l.Id(),0}
                }}
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (void_try5-void_try5) */
              if ErrorIn(void_try5) {Result = void_try5
              break
              } else {
              }}
              } else {
              void_try5 = ToException(Core.C_general_error.Make(MakeString("option: -m <module>").Id(),CNIL.Id())).Close()
              /* If-6 */} 
            /* If!5 */}  else if (ToString(l.ValuesO()[1-1]).Value == MakeString("-n").Value) /* If:5 */{ 
            _Zinit_ask = CFALSE
            l = l.Skip(1)
            void_try5 = EID{l.Id(),0}
            } else {
            if (ToString(l.ValuesO()[1-1]).At(1) == '-') /* If:6 */{ 
              void_try5 = Core.F_print_any(l.ValuesO()[1-1])
              /* ERROR PROTECTION INSERTED (void_try5-void_try5) */
              if ErrorIn(void_try5) {Result = void_try5
              break
              } else {
              PRINC(" is an unvalid option\n")
              void_try5 = EVOID
              }
              } else {
              void_try5 = EID{CFALSE.Id(),0}
              /* If-6 */} 
            /* ERROR PROTECTION INSERTED (void_try5-void_try5) */
            if ErrorIn(void_try5) {Result = void_try5
            break
            } else {
            l = ToType(C_string.Id()).EmptyList()
            void_try5 = EID{l.Id(),0}
            }
            /* If-5 */} 
          /* ERROR PROTECTION INSERTED (void_try5-Result) */
          if ErrorIn(void_try5) {Result = void_try5
          Result = void_try5
          break
          } else {
          /* while-4 */} 
        }
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        if (_Zinit_ask == CTRUE) /* If:4 */{ 
          Result = F_load_string(MakeString("init"))
          } else {
          Result = EID{CFALSE.Id(),0}
          /* If-4 */} 
        }
        if ErrorIn(Result){ 
          /* s=EID */ClEnv.Index = h_index
          ClEnv.Base = h_base
          C_reader.RestoreState()
          Result = F_debug_if_possible_void()
          } 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        Result = C_reader.TopLevel()
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC("[regular exit] Bye.\n")
        Result = EVOID
        }}
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: simple_main @ void (throw: true) 
func E_Reader_simple_main_void (_CL_obj EID) EID { 
    return /*(sm for simple_main @ void= EID)*/ F_Reader_simple_main_void( )} 
  
// *********************************************************************
// *      Part 2: Inspection                                           *
// *********************************************************************
// this is the method that the user calls
//
/* {1} OPT.The go function for: inspect(self:any) [] */
func F_inspect_any (self *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var m0 *ClaireModule   = ClEnv.Module_I
      /* noccur = 1 */
      /* Let:3 */{ 
        var ix int  = 0
        /* noccur = 3 */
        if (self.Isa.IsIn(C_list) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var g0456 *ClaireList   = ToList(self)
            /* noccur = 2 */
            /* Let:6 */{ 
              var i int  = 1
              /* noccur = 5 */
              /* Let:7 */{ 
                var g0457 int  = g0456.Length()
                /* noccur = 1 */
                Result= EID{CFALSE.Id(),0}
                for (i <= g0457) /* while:8 */{ 
                  var void_try9 EID 
                  _ = void_try9
                  { 
                  F_princ_integer(i)
                  PRINC(": ")
                  void_try9 = Core.F_CALL(C_print,ARGS(g0456.At(i-1).ToEID()))
                  /* ERROR PROTECTION INSERTED (void_try9-void_try9) */
                  if ErrorIn(void_try9) {Result = void_try9
                  break
                  } else {
                  PRINC("\n")
                  void_try9 = EVOID
                  }
                  {
                  i = (i+1)
                  }
                  /* while-8 */} 
                }
                /* Let-7 */} 
              /* Let-6 */} 
            /* Let-5 */} 
          /* If!4 */}  else if (Core.F__Z_any1(self,C_object) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var g0458 *ClaireObject   = ToObject(self)
            /* noccur = 2 */
            /* For:6 */{ 
              var rel *ClaireAny  
              _ = rel
              Result= EID{CFALSE.Id(),0}
              var rel_support *ClaireList  
              rel_support = g0458.Id().Isa.Slots
              for _,rel = range(rel_support.ValuesO())/* loop2:7 */{ 
                var void_try8 EID 
                _ = void_try8
                /* Let:8 */{ 
                  var m *ClaireModule   = ToRestriction(rel).Selector.Name.Module_I()
                  /* noccur = 2 */
                  ix = (ix+1)
                  if ((m.Id() == m0.Id()) || 
                      ((m.Id() == C_claire.Id()) || 
                        (ToBoolean(C__starshowall_star.Value) == CTRUE))) /* If:9 */{ 
                    /* Let:10 */{ 
                      var val *ClaireAny   = Core.F_get_slot(ToSlot(rel),g0458)
                      /* noccur = 3 */
                      F_princ_integer(ix)
                      PRINC(": ")
                      void_try8 = Core.F_print_any(ToRestriction(rel).Selector.Id())
                      /* ERROR PROTECTION INSERTED (void_try8-void_try8) */
                      if ErrorIn(void_try8) {Result = void_try8
                      break
                      } else {
                      PRINC(" = ")
                      if (val.Isa.IsIn(C_list) == CTRUE) /* If:11 */{ 
                        /* Let:12 */{ 
                          var g0459 *ClaireList   = ToList(val)
                          /* noccur = 3 */
                          if (g0459.Length() < 10) /* If:13 */{ 
                            void_try8 = Language.F_pretty_print_any(g0459.Id())
                            } else {
                            /* Let:14 */{ 
                              var g0463UU *ClaireList  
                              /* noccur = 1 */
                              /* Let:15 */{ 
                                var i_bag *ClaireList   = ToType(CEMPTY.Id()).EmptyList()
                                /* noccur = 2 */
                                /* Let:16 */{ 
                                  var i int  = 1
                                  /* noccur = 4 */
                                  /* Let:17 */{ 
                                    var g0460 int  = 9
                                    /* noccur = 1 */
                                    for (i <= g0460) /* while:18 */{ 
                                      i_bag.AddFast(g0459.At(i-1))
                                      i = (i+1)
                                      /* while-18 */} 
                                    /* Let-17 */} 
                                  /* Let-16 */} 
                                g0463UU = i_bag
                                /* Let-15 */} 
                              void_try8 = Language.F_pretty_print_any(g0463UU.Id())
                              /* Let-14 */} 
                            /* ERROR PROTECTION INSERTED (void_try8-void_try8) */
                            if ErrorIn(void_try8) {Result = void_try8
                            break
                            } else {
                            void_try8 = Language.F_pretty_print_any(MakeString("...").Id())
                            /* ERROR PROTECTION INSERTED (void_try8-void_try8) */
                            if ErrorIn(void_try8) {Result = void_try8
                            break
                            } else {
                            }}
                            /* If-13 */} 
                          /* Let-12 */} 
                        } else {
                        void_try8 = Language.F_pretty_print_any(val)
                        /* If-11 */} 
                      /* ERROR PROTECTION INSERTED (void_try8-void_try8) */
                      if ErrorIn(void_try8) {Result = void_try8
                      break
                      } else {
                      PRINC("\n")
                      void_try8 = EVOID
                      }}
                      /* Let-10 */} 
                    } else {
                    void_try8 = EID{CFALSE.Id(),0}
                    /* If-9 */} 
                  /* ERROR PROTECTION INSERTED (void_try8-void_try8) */
                  if ErrorIn(void_try8) {Result = void_try8
                  break
                  } else {
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
            /* Let-5 */} 
          } else {
          Result = Language.F_pretty_print_any(self)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC("\n")
          Result = EVOID
          }
          /* If-4 */} 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        F_Reader_inspect_system_list2(MakeConstantList(self))
        Result = EID{C_None.Id(),0}
        }
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: inspect @ any (throw: true) 
func E_inspect_any (self EID) EID { 
    return /*(sm for inspect @ any= EID)*/ F_inspect_any(ANY(self) )} 
  
// this is the inspect top_level
//
/* {1} OPT.The go function for: inspect_loop(%read:any,old:list) [] */
func F_inspect_loop_any (_Zread *ClaireAny ,old *ClaireList ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var self *ClaireAny   = old.At(1-1)
      /* noccur = 2 */
      var g0465I *ClaireBoolean  
      if (_Zread.Isa.IsIn(Language.C_Call) == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var g0464 *Language.Call   = Language.To_Call(_Zread)
          /* noccur = 1 */
          g0465I = Equal(g0464.Selector.Id(),C_put.Id())
          /* Let-4 */} 
        } else {
        g0465I = CFALSE
        /* If-3 */} 
      if (g0465I == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var n int  = ToInteger(ToList(OBJ(Core.F_CALL(C_args,ARGS(_Zread.ToEID())))).At(1-1)).Value
          /* noccur = 3 */
          /* Let:5 */{ 
            var s *ClaireSymbol  
            /* noccur = 1 */
            var s_try04666 EID 
            s_try04666 = Language.F_extract_symbol_any(ToList(OBJ(Core.F_CALL(C_args,ARGS(_Zread.ToEID())))).At(2-1))
            /* ERROR PROTECTION INSERTED (s-Result) */
            if ErrorIn(s_try04666) {Result = s_try04666
            } else {
            s = ToSymbol(OBJ(s_try04666))
            if (C_integer.Id() != OWNER(EID{C__INT,IVAL(n)}).Id()) /* If:6 */{ 
              Result = ToException(Core.C_general_error.Make(MakeString("[128] ~S should be an integer").Id(),MakeConstantList(MakeInteger(n).Id()).Id())).Close()
              } else {
              Result = EID{CFALSE.Id(),0}
              /* If-6 */} 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            /* Let:6 */{ 
              var val *ClaireAny  
              /* noccur = 3 */
              var val_try04677 EID 
              val_try04677 = F_get_from_integer_any(self,n)
              /* ERROR PROTECTION INSERTED (val-Result) */
              if ErrorIn(val_try04677) {Result = val_try04677
              } else {
              val = ANY(val_try04677)
              Core.ToGlobalVariable(OBJ(Core.F_new_class2(Core.C_global_variable,s))).Value = val
              Result = F_inspect_any(val)
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              old = F_cons_any(val,old)
              Result = EID{old.Id(),0}
              }
              }
              /* Let-6 */} 
            }
            }
            /* Let-5 */} 
          /* Let-4 */} 
        /* If!3 */}  else if (_Zread == C_up.Id()) /* If:3 */{ 
        if (old.Length() > 1) /* If:4 */{ 
          var old_try04685 EID 
          old_try04685 = old.Cdr()
          /* ERROR PROTECTION INSERTED (old-Result) */
          if ErrorIn(old_try04685) {Result = old_try04685
          } else {
          old = ToList(OBJ(old_try04685))
          Result = EID{old.Id(),0}
          Result = F_inspect_any(old.At(1-1))
          }
          } else {
          Result = EID{CFALSE.Id(),0}
          /* If-4 */} 
        /* If!3 */}  else if (C_integer.Id() == _Zread.Isa.Id()) /* If:3 */{ 
        /* Let:4 */{ 
          var val *ClaireAny  
          /* noccur = 2 */
          var val_try04695 EID 
          val_try04695 = F_get_from_integer_any(self,ToInteger(_Zread).Value)
          /* ERROR PROTECTION INSERTED (val-Result) */
          if ErrorIn(val_try04695) {Result = val_try04695
          } else {
          val = ANY(val_try04695)
          old = F_cons_any(val,old)
          Result = F_inspect_any(val)
          }
          /* Let-4 */} 
        /* If!3 */}  else if (_Zread.Isa.IsIn(C_thing) == CTRUE) /* If:3 */{ 
        old = F_cons_any(_Zread,old)
        Result = F_inspect_any(_Zread)
        } else {
        PRINC("=> given to inspector is wrong.\n")
        Result = EVOID
        /* If-3 */} 
      /* Let-2 */} 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    F_Reader_inspect_system_list2(old)
    Result = EVOID
    }
    return Result} 
  
// The EID go function for: inspect_loop @ any (throw: true) 
func E_inspect_loop_any (_Zread EID,old EID) EID { 
    return /*(sm for inspect_loop @ any= EID)*/ F_inspect_loop_any(ANY(_Zread),ToList(OBJ(old)) )} 
  
// get the information bound to the index
//
/* {1} OPT.The go function for: get_from_integer(self:any,n:integer) [] */
func F_get_from_integer_any (self *ClaireAny ,n int) EID { 
    var Result EID 
    if (self.Isa.IsIn(C_list) == CTRUE) /* If:2 */{ 
      if ((n > 0) && 
          (n <= ToList(self).Length())) /* If:3 */{ 
        Result = Core.F_CALL(C_nth,ARGS(self.ToEID(),EID{C__INT,IVAL(n)}))
        } else {
        F_princ_integer(n)
        PRINC(" in not a good index for ")
        Result = Core.F_CALL(C_print,ARGS(self.ToEID()))
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC(".\n")
        Result = EVOID
        }
        {
        Result = self.ToEID()
        }
        /* If-3 */} 
      } else {
      /* Let:3 */{ 
        var l *ClaireList   = self.Isa.Slots
        /* noccur = 2 */
        if ((n > 0) && 
            (n <= l.Length())) /* If:4 */{ 
          Result = Core.F_SUPER(C_get, C_slot, ARGS(l.ValuesO()[n-1].ToEID(),self.ToEID()))
          } else {
          F_princ_integer(n)
          PRINC(" is not a good index for ")
          Result = Core.F_CALL(C_print,ARGS(self.ToEID()))
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(".\n")
          Result = EVOID
          }
          {
          Result = self.ToEID()
          }
          /* If-4 */} 
        /* Let-3 */} 
      /* If-2 */} 
    return Result} 
  
// The EID go function for: get_from_integer @ any (throw: true) 
func E_get_from_integer_any (self EID,n EID) EID { 
    return /*(sm for get_from_integer @ any= EID)*/ F_get_from_integer_any(ANY(self),INT(n) )} 
  
// *********************************************************************
// *      Part 2: Trace methods                                        *
// *********************************************************************
// instrument the code generated from the rules
// this is the control method to CLAIRE tracer
// notice that trace(where) activates the call_count
/* {1} OPT.The go function for: iClaire/trace_on(self:any) [] */
func F_trace_on_any (self *ClaireAny ) EID { 
    var Result EID 
    if (self.Isa.IsIn(C_property) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0470 *ClaireProperty   = ToProperty(self)
        /* noccur = 3 */
        if (g0470.Id() == Core.C_spy.Id()) /* If:4 */{ 
          /* Let:5 */{ 
            var m *ClaireAny   = Core.F__at_property1(Core.C_spy,C_void).Id()
            /* noccur = 2 */
            if (F_boolean_I_any(m) == CTRUE) /* If:6 */{ 
              Result = F_store_object(ToObject(ClEnv.Id()),
                16,
                C_object,
                m,
                CFALSE).ToEID()
              } else {
              Result = EID{CFALSE.Id(),0}
              /* If-6 */} 
            /* Let-5 */} 
          /* If!4 */}  else if (g0470.Id() == C_where.Id()) /* If:4 */{ 
          Result = Core.F_write_property(C_Kernel_call_count,ToObject(ClEnv.Id()),MakeInteger(1).Id())
          } else {
          /* update:5 */{ 
            var va_arg1 *ClaireProperty  
            var va_arg2 int 
            va_arg1 = g0470
            va_arg2 = (5-ClEnv.Verbose)
            /* ---------- now we compile update trace!(va_arg1) := va_arg2 ------- */
            va_arg1.Trace_I = va_arg2
            Result = EID{C__INT,IVAL(va_arg2)}
            /* update-5 */} 
          /* If-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(C_environment) == CTRUE) /* If:2 */{ 
      /* update:3 */{ 
        var va_arg1 *ClaireEnvironment  
        var va_arg2 int 
        va_arg1 = ClEnv
        va_arg2 = 1
        /* ---------- now we compile update trace!(va_arg1) := va_arg2 ------- */
        va_arg1.Trace_I = va_arg2
        Result = EID{C__INT,IVAL(va_arg2)}
        /* update-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(C_module) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0472 *ClaireModule   = ToModule(self)
        /* noccur = 4 */
        if (g0472.Status > 2) /* If:4 */{ 
          g0472.Status = 4
          /* If-4 */} 
        /* For:4 */{ 
          var m *ClaireAny  
          _ = m
          Result= EID{CFALSE.Id(),0}
          var m_support *ClaireList  
          m_support = g0472.Parts
          m_len := m_support.Length()
          for i_it := 0; i_it < m_len; i_it++ { 
            m = m_support.At(i_it)
            var void_try6 EID 
            _ = void_try6
            void_try6 = F_trace_on_any(m)
            /* ERROR PROTECTION INSERTED (void_try6-Result) */
            if ErrorIn(void_try6) {Result = void_try6
            Result = void_try6
            break
            } else {
            }
            /* loop-5 */} 
          /* For-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (C_port.Id() == self.Isa.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var g0473 *ClairePort   = ToPort(self)
        /* noccur = 1 */
        /* update:4 */{ 
          var va_arg1 *ClaireEnvironment  
          var va_arg2 *ClairePort  
          va_arg1 = ClEnv
          va_arg2 = g0473
          /* ---------- now we compile update ctrace(va_arg1) := va_arg2 ------- */
          va_arg1.Ctrace = va_arg2
          Result = va_arg2.ToEID()
          /* update-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (C_string.Id() == self.Isa.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var g0474 *ClaireString   = ToString(self)
        /* noccur = 1 */
        /* update:4 */{ 
          var va_arg1 *ClaireEnvironment  
          var va_arg2 *ClairePort  
          va_arg1 = ClEnv
          var va_arg2_try04775 EID 
          va_arg2_try04775 = F_fopen_string(g0474,MakeString("w"))
          /* ERROR PROTECTION INSERTED (va_arg2-Result) */
          if ErrorIn(va_arg2_try04775) {Result = va_arg2_try04775
          } else {
          va_arg2 = ToPort(OBJ(va_arg2_try04775))
          /* ---------- now we compile update ctrace(va_arg1) := va_arg2 ------- */
          va_arg1.Ctrace = va_arg2
          Result = va_arg2.ToEID()
          }
          /* update-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (C_integer.Id() == self.Isa.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var g0475 int  = ToInteger(self).Value
        /* noccur = 1 */
        /* update:4 */{ 
          var va_arg1 *ClaireEnvironment  
          var va_arg2 int 
          va_arg1 = ClEnv
          va_arg2 = g0475
          /* ---------- now we compile update verbose(va_arg1) := va_arg2 ------- */
          va_arg1.Verbose = va_arg2
          Result = EID{C__INT,IVAL(va_arg2)}
          /* update-4 */} 
        /* Let-3 */} 
      } else {
      Result = ToException(Core.C_general_error.Make(MakeString("[129] trace not implemented on ~S\n").Id(),MakeConstantList(self).Id())).Close()
      /* If-2 */} 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    Result = self.ToEID()
    }
    return Result} 
  
// The EID go function for: iClaire/trace_on @ any (throw: true) 
func E_trace_on_any (self EID) EID { 
    return /*(sm for iClaire/trace_on @ any= EID)*/ F_trace_on_any(ANY(self) )} 
  
/* {1} OPT.The go function for: untrace(self:any) [] */
func F_untrace_any (self *ClaireAny ) EID { 
    var Result EID 
    if (self.Isa.IsIn(C_property) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0478 *ClaireProperty   = ToProperty(self)
        /* noccur = 3 */
        if (g0478.Id() == Core.C_spy.Id()) /* If:4 */{ 
          /* update:5 */{ 
            var va_arg1 *ClaireEnvironment  
            var va_arg2 *ClaireObject  
            va_arg1 = ClEnv
            va_arg2 = ToObject(CNULL)
            /* ---------- now we compile update spy!(va_arg1) := va_arg2 ------- */
            va_arg1.Spy_I = va_arg2
            Result = EID{va_arg2.Id(),0}
            /* update-5 */} 
          /* If!4 */}  else if (g0478.Id() == C_where.Id()) /* If:4 */{ 
          Result = Core.F_write_property(C_Kernel_call_count,ToObject(ClEnv.Id()),MakeInteger(-1).Id())
          } else {
          /* update:5 */{ 
            var va_arg1 *ClaireProperty  
            var va_arg2 int 
            va_arg1 = g0478
            va_arg2 = 0
            /* ---------- now we compile update trace!(va_arg1) := va_arg2 ------- */
            va_arg1.Trace_I = va_arg2
            Result = EID{C__INT,IVAL(va_arg2)}
            /* update-5 */} 
          /* If-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(C_environment) == CTRUE) /* If:2 */{ 
      /* update:3 */{ 
        var va_arg1 *ClaireEnvironment  
        var va_arg2 int 
        va_arg1 = ClEnv
        va_arg2 = 0
        /* ---------- now we compile update trace!(va_arg1) := va_arg2 ------- */
        va_arg1.Trace_I = va_arg2
        Result = EID{C__INT,IVAL(va_arg2)}
        /* update-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(C_module) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0480 *ClaireModule   = ToModule(self)
        /* noccur = 4 */
        if (g0480.Status == 4) /* If:4 */{ 
          g0480.Status = 3
          /* If-4 */} 
        /* For:4 */{ 
          var m *ClaireAny  
          _ = m
          Result= EID{CFALSE.Id(),0}
          var m_support *ClaireList  
          m_support = g0480.Parts
          m_len := m_support.Length()
          for i_it := 0; i_it < m_len; i_it++ { 
            m = m_support.At(i_it)
            var void_try6 EID 
            _ = void_try6
            void_try6 = F_untrace_any(m)
            /* ERROR PROTECTION INSERTED (void_try6-Result) */
            if ErrorIn(void_try6) {Result = void_try6
            Result = void_try6
            break
            } else {
            }
            /* loop-5 */} 
          /* For-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (C_port.Id() == self.Isa.Id()) /* If:2 */{ 
      /* update:3 */{ 
        var va_arg1 *ClaireEnvironment  
        var va_arg2 *ClairePort  
        va_arg1 = ClEnv
        va_arg2 = ToPort(C_stdout.Value)
        /* ---------- now we compile update ctrace(va_arg1) := va_arg2 ------- */
        va_arg1.Ctrace = va_arg2
        Result = va_arg2.ToEID()
        /* update-3 */} 
      } else {
      Result = ToException(Core.C_general_error.Make(MakeString("[130] untrace not implemented on ~S\n").Id(),MakeConstantList(self).Id())).Close()
      /* If-2 */} 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    Result = self.ToEID()
    }
    return Result} 
  
// The EID go function for: untrace @ any (throw: true) 
func E_untrace_any (self EID) EID { 
    return /*(sm for untrace @ any= EID)*/ F_untrace_any(ANY(self) )} 
  
// a filter to restrict the impact of spy
// we put the special value nil (emply list of demons => OK) to mark that spying
// should be waken up on properties from l
/* {1} OPT.The go function for: spy(l:listargs) [] */
func F_spy_listargs2_Reader (l *ClaireList )  { 
    // procedure body with s =  
/* Let:2 */{ 
      var m *ClaireAny   = Core.F__at_property1(Core.C_spy,C_void).Id()
      /* noccur = 2 */
      if (F_boolean_I_any(m) == CTRUE) /* If:3 */{ 
        F_store_object(ToObject(ClEnv.Id()),
          16,
          C_object,
          m,
          CFALSE)
        /* For:4 */{ 
          var g0483 *ClaireAny  
          _ = g0483
          var g0483_support *ClaireSet  
          g0483_support = C_property.Descendents
          for _,g0483 = range(g0483_support.Values)/* loop2:5 */{ 
            /* Let:6 */{ 
              var g0484 *ClaireBoolean  
              /* noccur = 2 */
              /* For:7 */{ 
                var f *ClaireAny  
                _ = f
                g0484= CFALSE
                var f_support *ClaireList  
                f_support = ToClass(g0483).Instances
                f_len := f_support.Length()
                for i_it := 0; i_it < f_len; i_it++ { 
                  f = f_support.At(i_it)
                  if (ToList(l.Id()).Memq(f) == CTRUE) /* If:9 */{ 
                    ToRelation(f).IfWrite = CNIL.Id()
                    /* If-9 */} 
                  /* loop-8 */} 
                /* For-7 */} 
              if (g0484 == CTRUE) /* If:7 */{ 
                 /*v = Unused, s =void*/

                break
                /* If-7 */} 
              /* Let-6 */} 
            /* loop-5 */} 
          /* For-4 */} 
        /* If-3 */} 
      /* Let-2 */} 
    } 
  
// The EID go function for: spy @ listargs (throw: false) 
func E_spy_listargs2_Reader (l EID) EID { 
    /*(sm for spy @ listargs= void)*/ F_spy_listargs2_Reader(ToList(OBJ(l)) )
    return EVOID} 
  
// used to trace the trigger of a rule
/* {1} OPT.The go function for: trace_rule(R:relation,s:string,x:any,y:any,u:any,v:any) [] */
func F_trace_rule_relation (R *ClaireRelation ,s *ClaireString ,x *ClaireAny ,y *ClaireAny ,u *ClaireAny ,v *ClaireAny ) EID { 
    var Result EID 
    if ((C_if_write.Trace_I+ClEnv.Verbose) >= 5) /* If:2 */{ 
      /* Let:3 */{ 
        var p *ClaireAny   = Core.F_get_property(C_ctrace,ToObject(ClEnv.Id()))
        /* noccur = 5 */
        if (p != CNULL) /* If:4 */{ 
          p = ToPort(p).UseAsOutput().Id()
          /* If-4 */} 
        PRINC("--- the rule ")
        F_princ_string(s)
        PRINC(" is triggered for (")
        Result = Core.F_CALL(C_print,ARGS(u.ToEID()))
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC(",")
        Result = Core.F_CALL(C_print,ARGS(v.ToEID()))
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC(") by an update ")
        Result = Core.F_print_any(R.Id())
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC("(")
        Result = Core.F_CALL(C_print,ARGS(x.ToEID()))
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC(") ")
        F_princ_string(ToString(IfThenElse((R.Multivalued_ask == CTRUE),
          MakeString(":add").Id(),
          MakeString(":=").Id())))
        PRINC(" ")
        Result = Core.F_CALL(C_print,ARGS(y.ToEID()))
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC(" \n")
        Result = EVOID
        }}}}}
        {
        if (p != CNULL) /* If:4 */{ 
          Result = ToPort(p).UseAsOutput().ToEID()
          } else {
          Result = EID{CFALSE.Id(),0}
          /* If-4 */} 
        }
        /* Let-3 */} 
      } else {
      Result = EID{CFALSE.Id(),0}
      /* If-2 */} 
    return Result} 
  
// The EID go function for: trace_rule @ relation (throw: true) 
func E_trace_rule_relation (R EID,s EID,x EID,y EID,u EID,v EID) EID { 
    return /*(sm for trace_rule @ relation= EID)*/ F_trace_rule_relation(ToRelation(OBJ(R)),
      ToString(OBJ(s)),
      ANY(x),
      ANY(y),
      ANY(u),
      ANY(v) )} 
  
// stores a set of stopping values
// this is a cool feature : stop(p, list(a1,a2)) => p(x,y) will stop if x = a1 and y = a2
/* {1} OPT.The go function for: stop(p:property,l:listargs) [] */
func F_stop_property (p *ClaireProperty ,l *ClaireList ) EID { 
    var Result EID 
    if (Core.F_get_table(Core.C_Core_StopProperty,p.Id()) == CNULL) /* If:2 */{ 
      Core.F_put_table(Core.C_Core_StopProperty,p.Id(),MakeConstantList(l.Id()).Id())
      Result = EVOID
      /* If!2 */}  else if (l.Id() == CNIL.Id()) /* If:2 */{ 
      Core.F_put_table(Core.C_Core_StopProperty,p.Id(),CNULL)
      Result = EVOID
      } else {
      /* Let:3 */{ 
        var g0486UU *ClaireList  
        /* noccur = 1 */
        var g0486UU_try04874 EID 
        /* Let:4 */{ 
          var g0488UU *ClaireAny  
          /* noccur = 1 */
          var g0488UU_try04895 EID 
          g0488UU_try04895 = Core.F_nth_table1(Core.C_Core_StopProperty,p.Id())
          /* ERROR PROTECTION INSERTED (g0488UU-g0486UU_try04874) */
          if ErrorIn(g0488UU_try04895) {g0486UU_try04874 = g0488UU_try04895
          } else {
          g0488UU = ANY(g0488UU_try04895)
          g0486UU_try04874 = ToList(g0488UU).Add(MakeConstantList(l.Id()).Id())
          }
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (g0486UU-Result) */
        if ErrorIn(g0486UU_try04874) {Result = g0486UU_try04874
        } else {
        g0486UU = ToList(OBJ(g0486UU_try04874))
        Core.F_put_table(Core.C_Core_StopProperty,p.Id(),g0486UU.Id())
        Result = EVOID
        }
        /* Let-3 */} 
      /* If-2 */} 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    Result = EID{CTRUE.Id(),0}
    }
    return Result} 
  
// The EID go function for: stop @ property (throw: true) 
func E_stop_property (p EID,l EID) EID { 
    return /*(sm for stop @ property= EID)*/ F_stop_property(ToProperty(OBJ(p)),ToList(OBJ(l)) )} 
  
// ******************************************************************
// *    Part 3: The debugger interface                              *
// ******************************************************************
// toggle the debug mode
/* {1} OPT.The go function for: debug(_CL_obj:void) [] */
func F_debug_void ()  { 
    // procedure body with s =  
if (ClEnv.Debug_I != -1) /* If:2 */{ 
      ClEnv.Debug_I = -1
      PRINC("debugger removed\n")
      } else {
      ClEnv.Debug_I = 0
      ClEnv.Ctrace = ToPort(C_stdout.Value)
      PRINC("debugger installed\n")
      /* If-2 */} 
    } 
  
// The EID go function for: debug @ void (throw: false) 
func E_debug_void (_CL_obj EID) EID { 
    /*(sm for debug @ void= void)*/ F_debug_void( )
    return EVOID} 
  
// this method is called when an error has occured. The value of index
// is recalled with last_index, so that the actual content of the stack is
// preserved.
/* {1} OPT.The go function for: call_debug(_CL_obj:void) [] */
func F_call_debug_void () *ClaireAny  { 
    // procedure body with s =  
var Result *ClaireAny  
    /* Let:2 */{ 
      var top int  = ClEnv.LastDebug
      /* noccur = 1 */
      C_reader.DebugLoop()
      ClEnv.Spy_I = ToObject(CNULL)
      ClEnv.Trace_I = 0
      ClEnv.Base= ClEnv.LastIndex
      ClEnv.Index= (ClEnv.LastIndex+1)
      ClEnv.Debug_I = top
      if (ClEnv.Verbose > -1) /* If:3 */{ 
        F_print_exception_void()
        /* If-3 */} 
      C_reader.Fromp = ToPort(C_stdin.Value)
      C_reader.Index = 0
      /* Let:3 */{ 
        var c *ClaireAny   = Language.C_iClaire_LastCall.Value
        /* noccur = 4 */
        if (c != CNULL) /* If:4 */{ 
          var Result_try5 EID 
          h_index := ClEnv.Index /* Handle */
          h_base := ClEnv.Base
          var g0490I *ClaireBoolean  
          var g0490I_try04915 EID 
          /* Let:5 */{ 
            var g0492UU *ClaireAny  
            /* noccur = 1 */
            var g0492UU_try04936 EID 
            g0492UU_try04936 = Core.F_nth_table1(C_Reader_DBline,c)
            /* ERROR PROTECTION INSERTED (g0492UU-g0490I_try04915) */
            if ErrorIn(g0492UU_try04936) {g0490I_try04915 = g0492UU_try04936
            } else {
            g0492UU = ANY(g0492UU_try04936)
            g0490I_try04915 = EID{Core.F__sup_integer(ToInteger(g0492UU).Value,0).Id(),0}
            }
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (g0490I-Result_try5) */
          if ErrorIn(g0490I_try04915) {Result_try5 = g0490I_try04915
          } else {
          g0490I = ToBoolean(OBJ(g0490I_try04915))
          if (g0490I == CTRUE) /* If:5 */{ 
            PRINC(" \n---- Last call ")
            Result_try5 = Core.F_CALL(C_print,ARGS(c.ToEID()))
            /* ERROR PROTECTION INSERTED (Result_try5-Result_try5) */
            if !ErrorIn(Result_try5) {
            PRINC(" in line ")
            /* Let:6 */{ 
              var g0494UU *ClaireAny  
              /* noccur = 1 */
              var g0494UU_try04957 EID 
              g0494UU_try04957 = Core.F_nth_table1(C_Reader_DBline,c)
              /* ERROR PROTECTION INSERTED (g0494UU-Result_try5) */
              if ErrorIn(g0494UU_try04957) {Result_try5 = g0494UU_try04957
              } else {
              g0494UU = ANY(g0494UU_try04957)
              F_princ_integer(ToInteger(g0494UU).Value)
              Result_try5 = EVOID
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (Result_try5-Result_try5) */
            if !ErrorIn(Result_try5) {
            PRINC("\n")
            Result_try5 = EVOID
            }}
            } else {
            Result_try5 = EID{CFALSE.Id(),0}
            /* If-5 */} 
          }
          if ErrorIn(Result_try5){ 
            /* s=void */ClEnv.Index = h_index
            ClEnv.Base = h_base
            
            } 
          } else {
          
          /* If-4 */} 
        /* Let-3 */} 
      C__starindex_star.Value = MakeInteger(1).Id()
      C__starcurd_star.Value = MakeInteger(ClEnv.Debug_I).Id()
      Result = MakeInteger(ClEnv.Debug_I).Id()
      C__starmaxd_star.Value = Result
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: call_debug @ void (throw: false) 
func E_call_debug_void (_CL_obj EID) EID { 
    return /*(sm for call_debug @ void= any)*/ F_call_debug_void( ).ToEID()} 
  
// this method is called when an error has occured. The value of index
// is recalled with last_index, so that the actual content of the stack is
// preserved.
/* {1} OPT.The go function for: breakpoint(_CL_obj:void) [] */
func F_breakpoint_void () EID { 
    var Result EID 
    /* Let:2 */{ 
      var top int  = ClEnv.Debug_I
      /* noccur = 4 */
      /* Let:3 */{ 
        var t int  = ClEnv.Trace_I
        /* noccur = 1 */
        ClEnv.Trace_I = 0
        C__starindex_star.Value = MakeInteger(0).Id()
        C__starcurd_star.Value = MakeInteger(top).Id()
        C__starmaxd_star.Value = MakeInteger(top).Id()
        var g0497I *ClaireBoolean  
        var g0497I_try04984 EID 
        /* and:4 */{ 
          var v_and4 *ClaireBoolean  
          
          var v_and4_try04995 EID 
          v_and4_try04995 = Core.F_CALL(ToProperty(C__sup.Id()),ARGS(ClEnv.EvalStack[top],EID{C__INT,IVAL(0)}))
          /* ERROR PROTECTION INSERTED (v_and4-g0497I_try04984) */
          if ErrorIn(v_and4_try04995) {g0497I_try04984 = v_and4_try04995
          } else {
          v_and4 = ToBoolean(OBJ(v_and4_try04995))
          if (v_and4 == CFALSE) {g0497I_try04984 = EID{CFALSE.Id(),0}
          } else /* arg:5 */{ 
            v_and4 = Core.F__sup_integer(ClEnv.Debug_I,0)
            if (v_and4 == CFALSE) {g0497I_try04984 = EID{CFALSE.Id(),0}
            } else /* arg:6 */{ 
              g0497I_try04984 = EID{CTRUE.Id(),0}/* arg-6 */} 
            /* arg-5 */} 
          }
          /* and-4 */} 
        /* ERROR PROTECTION INSERTED (g0497I-Result) */
        if ErrorIn(g0497I_try04984) {Result = g0497I_try04984
        } else {
        g0497I = ToBoolean(OBJ(g0497I_try04984))
        if (g0497I == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var j int  = INT(ClEnv.EvalStack[top])
            /* noccur = 3 */
            /* Let:6 */{ 
              var num_args int  = (INT(ClEnv.EvalStack[(j+2)])-1)
              /* noccur = 1 */
              /* Let:7 */{ 
                var start int  = INT(ClEnv.EvalStack[(j+3)])
                /* noccur = 3 */
                /* Let:8 */{ 
                  var m *ClaireAny   = ANY(ClEnv.EvalStack[(j+1)])
                  /* noccur = 1 */
                  PRINC("break in ")
                  Result = Core.F_CALL(C_print,ARGS(m.ToEID()))
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  PRINC("(")
                  Result = Core.F_CALL(C_print,ARGS(ClEnv.EvalStack[start]))
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  /* Let:9 */{ 
                    var i int  = (start+1)
                    /* noccur = 4 */
                    /* Let:10 */{ 
                      var g0496 int  = (start+num_args)
                      /* noccur = 1 */
                      Result= EID{CFALSE.Id(),0}
                      for (i <= g0496) /* while:11 */{ 
                        var void_try12 EID 
                        _ = void_try12
                        { 
                        PRINC(",")
                        void_try12 = Core.F_CALL(C_print,ARGS(ClEnv.EvalStack[i]))
                        /* ERROR PROTECTION INSERTED (void_try12-void_try12) */
                        if ErrorIn(void_try12) {Result = void_try12
                        break
                        } else {
                        PRINC("")
                        void_try12 = EVOID
                        }
                        {
                        i = (i+1)
                        }
                        /* while-11 */} 
                      }
                      /* Let-10 */} 
                    /* Let-9 */} 
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  PRINC(") [q] >")
                  Result = EVOID
                  }}}
                  {
                  /* Let:9 */{ 
                    var c *ClaireAny  
                    /* noccur = 3 */
                    var c_try050010 EID 
                    c_try050010 = F_read_string(F_CommandLoopVoid())
                    /* ERROR PROTECTION INSERTED (c-Result) */
                    if ErrorIn(c_try050010) {Result = c_try050010
                    } else {
                    c = ANY(c_try050010)
                    Result= EID{CFALSE.Id(),0}
                    for (c != C_q.Id()) /* while:10 */{ 
                      var void_try11 EID 
                      _ = void_try11
                      { 
                      void_try11 = EVAL(c)
                      /* ERROR PROTECTION INSERTED (void_try11-void_try11) */
                      if ErrorIn(void_try11) {Result = void_try11
                      break
                      } else {
                      PRINC("break>")
                      var c_try050111 EID 
                      c_try050111 = F_read_string(F_CommandLoopVoid())
                      /* ERROR PROTECTION INSERTED (c-void_try11) */
                      if ErrorIn(c_try050111) {void_try11 = c_try050111
                      Result = c_try050111
                      break
                      } else {
                      c = ANY(c_try050111)
                      void_try11 = c.ToEID()
                      }}
                      /* while-10 */} 
                    }
                    }
                    /* Let-9 */} 
                  }
                  /* Let-8 */} 
                /* Let-7 */} 
              /* Let-6 */} 
            /* Let-5 */} 
          } else {
          Result = EID{CFALSE.Id(),0}
          /* If-4 */} 
        }
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        /* update:4 */{ 
          var va_arg1 *ClaireEnvironment  
          var va_arg2 int 
          va_arg1 = ClEnv
          va_arg2 = t
          /* ---------- now we compile update trace!(va_arg1) := va_arg2 ------- */
          va_arg1.Trace_I = va_arg2
          Result = EID{C__INT,IVAL(va_arg2)}
          /* update-4 */} 
        }
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: breakpoint @ void (throw: true) 
func E_breakpoint_void (_CL_obj EID) EID { 
    return /*(sm for breakpoint @ void= EID)*/ F_breakpoint_void( )} 
  
// the four keyword
/* {1} OPT.The go function for: dn(x:integer) [] */
func F_dn_integer (x int)  { 
    // procedure body with s =  
for ((INT(ClEnv.EvalStack[ToInteger(C__starcurd_star.Value).Value]) > 0) && 
        (x > 0)) /* while:2 */{ 
      C__starcurd_star.Value = ANY(ClEnv.EvalStack[ToInteger(C__starcurd_star.Value).Value])
      C__starindex_star.Value = MakeInteger((ToInteger(C__starindex_star.Value).Value+1)).Id()
      x = (x-1)
      /* while-2 */} 
    } 
  
// The EID go function for: dn @ integer (throw: false) 
func E_dn_integer (x EID) EID { 
    /*(sm for dn @ integer= void)*/ F_dn_integer(INT(x) )
    return EVOID} 
  
/* {1} OPT.The go function for: up(x:integer) [] */
func F_up_integer (x int)  { 
    // procedure body with s =  
if (x > 0) /* If:2 */{ 
      /* Let:3 */{ 
        var indices *ClaireList   = CNIL
        /* noccur = 4 */
        /* Let:4 */{ 
          var ind int  = ToInteger(C__starmaxd_star.Value).Value
          /* noccur = 4 */
          for (ind != ToInteger(C__starcurd_star.Value).Value) /* while:5 */{ 
            indices = F_cons_any(MakeInteger(ind).Id(),indices)
            ind = INT(ClEnv.EvalStack[ind])
            /* while-5 */} 
          if (x > indices.Length()) /* If:5 */{ 
            C__starcurd_star.Value = C__starmaxd_star.Value
            C__starindex_star.Value = MakeInteger(1).Id()
            } else {
            C__starcurd_star.Value = indices.At(x-1)
            C__starindex_star.Value = MakeInteger((ToInteger(C__starindex_star.Value).Value-x)).Id()
            /* If-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* If-2 */} 
    } 
  
// The EID go function for: up @ integer (throw: false) 
func E_up_integer (x EID) EID { 
    /*(sm for up @ integer= void)*/ F_up_integer(INT(x) )
    return EVOID} 
  
// top is the top position in this stack (the last entered message)
/* {1} OPT.The go function for: where(x:integer) [] */
func F_where_integer (x int) EID { 
    var Result EID 
    /* Let:2 */{ 
      var j int  = ToInteger(C__starcurd_star.Value).Value
      /* noccur = 4 */
      /* Let:3 */{ 
        var stack_level int  = 0
        /* noccur = 3 */
        Result= EID{CFALSE.Id(),0}
        for ((j > 0) && 
            ((x > 0) && 
              (ClEnv.Debug_I > 0))) /* while:4 */{ 
          var void_try5 EID 
          _ = void_try5
          { 
          void_try5 = F_print_debug_info_integer(j,stack_level,ToInteger(C__starindex_star.Value).Value)
          /* ERROR PROTECTION INSERTED (void_try5-void_try5) */
          if ErrorIn(void_try5) {Result = void_try5
          break
          } else {
          stack_level = (stack_level+1)
          x = (x-1)
          j = INT(ClEnv.EvalStack[j])
          }
          /* while-4 */} 
        }
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: where @ integer (throw: true) 
func E_where_integer (x EID) EID { 
    return /*(sm for where @ integer= EID)*/ F_where_integer(INT(x) )} 
  
// note for interpretted methods .. they should be pushing their restriction
// on the stack vs. properties
/* {1} OPT.The go function for: print_debug_info(iClaire/index:integer,stack_level:integer,cur_index:integer) [] */
func F_print_debug_info_integer (index int,stack_level int,cur_index int) EID { 
    var Result EID 
    /* Let:2 */{ 
      var num_args int  = (INT(ClEnv.EvalStack[(index+2)])-1)
      /* noccur = 1 */
      /* Let:3 */{ 
        var start int  = INT(ClEnv.EvalStack[(index+3)])
        /* noccur = 3 */
        /* Let:4 */{ 
          var m *ClaireAny   = ANY(ClEnv.EvalStack[(index+1)])
          /* noccur = 1 */
          PRINC("debug[")
          F_princ_integer((cur_index+stack_level))
          PRINC("]>")
          /* Let:5 */{ 
            var x int  = 1
            /* noccur = 3 */
            /* Let:6 */{ 
              var g0502 int  = stack_level
              /* noccur = 1 */
              for (x <= g0502) /* while:7 */{ 
                PRINC(">")
                x = (x+1)
                /* while-7 */} 
              /* Let-6 */} 
            /* Let-5 */} 
          PRINC(" ")
          Result = Core.F_CALL(C_print,ARGS(m.ToEID()))
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC("(")
          Result = Core.F_CALL(C_print,ARGS(ClEnv.EvalStack[start]))
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          /* Let:5 */{ 
            var i int  = (start+1)
            /* noccur = 4 */
            /* Let:6 */{ 
              var g0503 int  = (start+num_args)
              /* noccur = 1 */
              Result= EID{CFALSE.Id(),0}
              for (i <= g0503) /* while:7 */{ 
                var void_try8 EID 
                _ = void_try8
                { 
                PRINC(",")
                void_try8 = Core.F_CALL(C_print,ARGS(ClEnv.EvalStack[i]))
                /* ERROR PROTECTION INSERTED (void_try8-void_try8) */
                if ErrorIn(void_try8) {Result = void_try8
                break
                } else {
                PRINC("")
                void_try8 = EVOID
                }
                {
                i = (i+1)
                }
                /* while-7 */} 
              }
              /* Let-6 */} 
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(")\n")
          Result = EVOID
          }}}
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: print_debug_info @ integer (throw: true) 
func E_print_debug_info_integer (index EID,stack_level EID,cur_index EID) EID { 
    return /*(sm for print_debug_info @ integer= EID)*/ F_print_debug_info_integer(INT(index),INT(stack_level),INT(cur_index) )} 
  
// debug version of the debugger :-)  => use as Reader/Show(n)
/* {1} OPT.The go function for: Show(n:integer) [] */
func F_Show_integer (n int) EID { 
    var Result EID 
    /* Let:2 */{ 
      var i int  = ToInteger(C__starcurd_star.Value).Value
      /* noccur = 9 */
      Result= EID{CFALSE.Id(),0}
      for ((i > 0) && 
          (n > 0)) /* while:3 */{ 
        var void_try4 EID 
        _ = void_try4
        /* Let:4 */{ 
          var num_args int  = (INT(ClEnv.EvalStack[(i+2)])-1)
          /* noccur = 2 */
          /* Let:5 */{ 
            var start int  = INT(ClEnv.EvalStack[(i+3)])
            /* noccur = 1 */
            PRINC("[")
            F_princ_integer(start)
            PRINC(" - ")
            F_princ_integer(i)
            PRINC("]: p = ")
            void_try4 = Core.F_CALL(C_print,ARGS(ClEnv.EvalStack[(i+1)]))
            /* ERROR PROTECTION INSERTED (void_try4-void_try4) */
            if ErrorIn(void_try4) {Result = void_try4
            break
            } else {
            PRINC(", narg = ")
            void_try4 = Core.F_print_any(MakeInteger(num_args).Id())
            /* ERROR PROTECTION INSERTED (void_try4-void_try4) */
            if ErrorIn(void_try4) {Result = void_try4
            break
            } else {
            PRINC(" \n")
            void_try4 = EVOID
            }}
            {
            /* Let:6 */{ 
              var j int  = 0
              /* noccur = 5 */
              /* Let:7 */{ 
                var g0504 int  = num_args
                /* noccur = 1 */
                void_try4= EID{CFALSE.Id(),0}
                for (j <= g0504) /* while:8 */{ 
                  var void_try9 EID 
                  _ = void_try9
                  { 
                  PRINC("  [")
                  F_princ_integer((j+i))
                  PRINC("]:")
                  void_try9 = Core.F_CALL(C_print,ARGS(ClEnv.EvalStack[(j+i)]))
                  /* ERROR PROTECTION INSERTED (void_try9-void_try9) */
                  if ErrorIn(void_try9) {void_try4 = void_try9
                  break
                  } else {
                  PRINC(" \n")
                  void_try9 = EVOID
                  }
                  {
                  j = (j+1)
                  }
                  /* while-8 */} 
                }
                /* Let-7 */} 
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (void_try4-void_try4) */
            if ErrorIn(void_try4) {Result = void_try4
            break
            } else {
            n = (n-1)
            i = INT(ClEnv.EvalStack[i])
            void_try4 = EID{C__INT,IVAL(i)}
            }}
            /* Let-5 */} 
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (void_try4-Result) */
        if ErrorIn(void_try4) {Result = void_try4
        Result = void_try4
        break
        } else {
        /* while-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: Show @ integer (throw: true) 
func E_Show_integer (n EID) EID { 
    return /*(sm for Show @ integer= EID)*/ F_Show_integer(INT(n) )} 
  
// go to next block
// top is the top position in this stack (the last entered message)
//
/* {1} OPT.The go function for: block(x:integer) [] */
func F_block_integer (x int) EID { 
    var Result EID 
    /* Let:2 */{ 
      var j int  = ToInteger(C__starcurd_star.Value).Value
      /* noccur = 6 */
      /* Let:3 */{ 
        var stack_level int  = 0
        /* noccur = 4 */
        Result= EID{CFALSE.Id(),0}
        for ((j > 0) && 
            ((x > 0) && 
              (ClEnv.Debug_I > 0))) /* while:4 */{ 
          var void_try5 EID 
          _ = void_try5
          { 
          /* Let:5 */{ 
            var nargs *ClaireAny   = ANY(ClEnv.EvalStack[(j+2)])
            /* noccur = 1 */
            /* Let:6 */{ 
              var start int  = INT(ClEnv.EvalStack[(j+3)])
              /* noccur = 4 */
              /* Let:7 */{ 
                var z *ClaireProperty   = ToProperty(OBJ(ClEnv.EvalStack[(j+1)]))
                /* noccur = 2 */
                /* Let:8 */{ 
                  var m *ClaireObject   = Core.F_find_which_list(z.Definition,OWNER(ClEnv.EvalStack[start]),start,INT(Core.F_CALL(ToProperty(Core.C__plus.Id()),ARGS(EID{C__INT,IVAL(start)},nargs.ToEID()))))
                  /* noccur = 3 */
                  if (C_method.Id() == m.Isa.Id()) /* If:9 */{ 
                    /* Let:10 */{ 
                      var g0505 *ClaireMethod   = ToMethod(m.Id())
                      /* noccur = 5 */
                      PRINC("debug[")
                      F_princ_integer((ToInteger(C__starindex_star.Value).Value+stack_level))
                      PRINC("] > ")
                      void_try5 = Core.F_print_any(g0505.Id())
                      /* ERROR PROTECTION INSERTED (void_try5-void_try5) */
                      if ErrorIn(void_try5) {Result = void_try5
                      break
                      } else {
                      PRINC("(")
                      if ((g0505.Formula.Id() != CNULL) && 
                          (g0505.Formula.Isa.IsIn(C_lambda) == CTRUE)) /* If:11 */{ 
                        /* Let:12 */{ 
                          var n int  = 0
                          /* noccur = 3 */
                          /* For:13 */{ 
                            var v *ClaireAny  
                            _ = v
                            void_try5= EID{CFALSE.Id(),0}
                            var v_support *ClaireList  
                            v_support = F_closure_build_lambda(g0505.Formula)
                            v_len := v_support.Length()
                            for i_it := 0; i_it < v_len; i_it++ { 
                              v = v_support.At(i_it)
                              var void_try15 EID 
                              _ = void_try15
                              { 
                              void_try15 = Core.F_CALL(C_print,ARGS(v.ToEID()))
                              /* ERROR PROTECTION INSERTED (void_try15-void_try15) */
                              if ErrorIn(void_try15) {void_try5 = void_try15
                              break
                              } else {
                              PRINC(" = ")
                              void_try15 = Core.F_CALL(C_print,ARGS(ClEnv.EvalStack[(start+n)]))
                              /* ERROR PROTECTION INSERTED (void_try15-void_try15) */
                              if ErrorIn(void_try15) {void_try5 = void_try15
                              break
                              } else {
                              PRINC(", ")
                              void_try15 = EVOID
                              }}
                              {
                              n = (n+1)
                              }
                              }
                              /* loop-14 */} 
                            /* For-13 */} 
                          /* Let-12 */} 
                        } else {
                        PRINC("<compiled:")
                        void_try5 = Core.F_print_any(g0505.Module_I.Id())
                        /* ERROR PROTECTION INSERTED (void_try5-void_try5) */
                        if ErrorIn(void_try5) {Result = void_try5
                        break
                        } else {
                        PRINC(">")
                        void_try5 = EVOID
                        }
                        /* If-11 */} 
                      /* ERROR PROTECTION INSERTED (void_try5-void_try5) */
                      if ErrorIn(void_try5) {Result = void_try5
                      break
                      } else {
                      PRINC(")\n")
                      void_try5 = EVOID
                      }}
                      /* Let-10 */} 
                    } else {
                    PRINC("debug[")
                    F_princ_integer((ToInteger(C__starindex_star.Value).Value+stack_level))
                    PRINC("] > ")
                    void_try5 = Core.F_print_any(z.Id())
                    /* ERROR PROTECTION INSERTED (void_try5-void_try5) */
                    if ErrorIn(void_try5) {Result = void_try5
                    break
                    } else {
                    PRINC(" -> ")
                    void_try5 = Core.F_print_any(m.Id())
                    /* ERROR PROTECTION INSERTED (void_try5-void_try5) */
                    if ErrorIn(void_try5) {Result = void_try5
                    break
                    } else {
                    PRINC("\n")
                    void_try5 = EVOID
                    }}
                    /* If-9 */} 
                  /* Let-8 */} 
                /* Let-7 */} 
              /* Let-6 */} 
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (void_try5-void_try5) */
          if ErrorIn(void_try5) {Result = void_try5
          break
          } else {
          stack_level = (stack_level+1)
          x = (x-1)
          j = INT(ClEnv.EvalStack[j])
          }
          /* while-4 */} 
        }
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: block @ integer (throw: true) 
func E_block_integer (x EID) EID { 
    return /*(sm for block @ integer= EID)*/ F_block_integer(INT(x) )} 
  
// computes the list of variables of a lambda, including everything
//
/* {1} OPT.The go function for: closure_build(self:lambda) [] */
func F_closure_build_lambda (self *ClaireLambda ) *ClaireList  { 
    // procedure body with s =  
var Result *ClaireList  
    /* Let:2 */{ 
      var lvar *ClaireList   = F_make_list_integer(self.Dimension,CEMPTY.Id())
      /* noccur = 3 */
      /* For:3 */{ 
        var x *ClaireAny  
        _ = x
        var x_support *ClaireList  
        x_support = self.Vars
        x_len := x_support.Length()
        for i_it := 0; i_it < x_len; i_it++ { 
          x = x_support.At(i_it)
          ToArray(lvar.Id()).NthPut((INT(Core.F_CALL(C_mClaire_index,ARGS(x.ToEID())))+1),x)
          /* loop-4 */} 
        /* For-3 */} 
      F_closure_build_any(self.Body,lvar)
      Result = lvar
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: closure_build @ lambda (throw: false) 
func E_closure_build_lambda (self EID) EID { 
    return EID{/*(sm for closure_build @ lambda= list)*/ F_closure_build_lambda(ToLambda(OBJ(self)) ).Id(),0}} 
  
// give to each lexical variable its right position in the stack
// answer with the number of lexical variable
//
/* {1} OPT.The go function for: closure_build(self:any,lvar:list) [] */
func F_closure_build_any (self *ClaireAny ,lvar *ClaireList )  { 
    // procedure body with s =  
if (self.Isa.IsIn(C_Variable) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0507 *ClaireVariable   = To_Variable(self)
        /* noccur = 2 */
        ToArray(lvar.Id()).NthPut((g0507.Index+1),g0507.Id())
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(C_Instruction) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0508 *ClaireInstruction   = To_Instruction(self)
        /* noccur = 2 */
        /* For:4 */{ 
          var s *ClaireAny  
          _ = s
          for _,s = range(g0508.Isa.Slots.ValuesO())/* loop:5 */{ 
            F_closure_build_any(Core.F_get_slot(ToSlot(s),ToObject(g0508.Id())),lvar)
            /* loop-5 */} 
          /* For-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(C_list) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0509 *ClaireList   = ToList(self)
        /* noccur = 1 */
        /* For:4 */{ 
          var x *ClaireAny  
          _ = x
          var x_support *ClaireList  
          x_support = g0509
          x_len := x_support.Length()
          for i_it := 0; i_it < x_len; i_it++ { 
            x = x_support.At(i_it)
            F_closure_build_any(x,lvar)
            /* loop-5 */} 
          /* For-4 */} 
        /* Let-3 */} 
      } else {
      
      /* If-2 */} 
    } 
  
// The EID go function for: closure_build @ any (throw: false) 
func E_closure_build_any (self EID,lvar EID) EID { 
    /*(sm for closure_build @ any= void)*/ F_closure_build_any(ANY(self),ToList(OBJ(lvar)) )
    return EVOID} 
  
// ******************************************************************
// *    Part 5:  Measure &  Profile                                 *
// ******************************************************************
// New in CLAIRE 3.4 - measure objects can be stored on a file and loaded later on
// a measure is a float value counter that stores the sum & sum of squares, to 
// number of experiments
// simple methods add, mean, stdev
/* {1} OPT.The go function for: close(x:measure) [] */
func (x *Measure ) Close () *Measure  { 
    // use function body compiling 
x.MIndex = C_measure.Instances.Length()
    return  x
    } 
  
// The EID go function for: close @ measure (throw: false) 
func E_close_measure (x EID) EID { 
    return EID{/*(sm for close @ measure= measure)*/ ToMeasure(OBJ(x)).Close( ).Id(),0}} 
  
/* {1} OPT.The go function for: add(x:measure,f:float) [] */
func (x *Measure ) Add (f float64) *Measure  { 
    // use function body compiling 
x.NumValue = (x.NumValue+1)
    x.SumValue = (x.SumValue+f)
    x.SumSquare = (x.SumSquare+(f*f))
    return  x
    } 
  
// The EID go function for: add @ measure (throw: false) 
func E_add_measure (x EID,f EID) EID { 
    return EID{/*(sm for add @ measure= measure)*/ ToMeasure(OBJ(x)).Add(FLOAT(f) ).Id(),0}} 
  
/* {1} OPT.The go function for: mean(x:measure) [] */
func (x *Measure ) Mean () float64 { 
    // use function body compiling 
if (x.NumValue == 0) /* body If:2 */{ 
      return  0
      } else {
      return  (x.SumValue/x.NumValue)
      /* body If-2 */} 
    } 
  
// The EID go function for: mean @ measure (throw: false) 
func E_mean_measure (x EID) EID { 
    return EID{C__FLOAT,FVAL(/*(sm for mean @ measure= float)*/ ToMeasure(OBJ(x)).Mean( ))}} 
  
/* {1} OPT.The go function for: stdev(x:measure) [] */
func (x *Measure ) Stdev () float64 { 
    // procedure body with s =  
var Result float64 
    /* Let:2 */{ 
      var y float64  = ((x.SumSquare/x.NumValue)-F__exp_float((x.SumValue/x.NumValue),2))
      /* noccur = 2 */
      if (y > 0) /* If:3 */{ 
        Result = F_sqrt_float(y)
        } else {
        Result = 0
        /* If-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: stdev @ measure (throw: false) 
func E_stdev_measure (x EID) EID { 
    return EID{C__FLOAT,FVAL(/*(sm for stdev @ measure= float)*/ ToMeasure(OBJ(x)).Stdev( ))}} 
  
/* {1} OPT.The go function for: stdev%(x:measure) [] */
func (x *Measure ) Stdev_Z () float64 { 
    // use function body compiling 
return  (x.Stdev()/x.Mean())
    } 
  
// The EID go function for: stdev% @ measure (throw: false) 
func E_stdev_Z_measure (x EID) EID { 
    return EID{C__FLOAT,FVAL(/*(sm for stdev% @ measure= float)*/ ToMeasure(OBJ(x)).Stdev_Z( ))}} 
  
/* {1} OPT.The go function for: reset(x:measure) [] */
func (x *Measure ) Reset ()  { 
    // procedure body with s =  
x.SumSquare = 0
    x.NumValue = 0
    x.SumValue = 0
    } 
  
// The EID go function for: reset @ measure (throw: false) 
func E_reset_measure (x EID) EID { 
    /*(sm for reset @ measure= void)*/ ToMeasure(OBJ(x)).Reset( )
    return EVOID} 
  
/* {1} OPT.The go function for: self_print(m:measure) [] */
func (m *Measure ) SelfPrint () EID { 
    var Result EID 
    Result = Core.F_printFDigit_float(m.Mean(),2)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("[")
    Result = Core.F_printFDigit_float(m.NumValue,0)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("]")
    Result = EVOID
    }}
    return Result} 
  
// The EID go function for: self_print @ measure (throw: true) 
func E_self_print_measure_Reader (m EID) EID { 
    return /*(sm for self_print @ measure= EID)*/ ToMeasure(OBJ(m)).SelfPrint( )} 
  
// two simple methods to store and retreive measures
//   logMeasure(s:string)  : creates a file
//   load(s:string)        : loads the files, that containts addLog(i,s,ss,n) line
/* {1} OPT.The go function for: logMeasure(s:string) [] */
func F_logMeasure_string (s *ClaireString ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var p *ClairePort  
      /* noccur = 2 */
      var p_try05163 EID 
      p_try05163 = F_fopen_string(s,MakeString("w"))
      /* ERROR PROTECTION INSERTED (p-Result) */
      if ErrorIn(p_try05163) {Result = p_try05163
      } else {
      p = ToPort(OBJ(p_try05163))
      /* Let:3 */{ 
        var n int  = Core.F_size_class(C_measure)
        /* noccur = 1 */
        p.UseAsOutput()
        PRINC("// log file produced on ")
        F_princ_string(F_date_I_integer(1))
        PRINC("")
        /* For:4 */{ 
          var g0514 *ClaireAny  
          _ = g0514
          var g0514_support *ClaireSet  
          g0514_support = C_measure.Descendents
          for _,g0514 = range(g0514_support.Values)/* loop2:5 */{ 
            /* Let:6 */{ 
              var g0515 *ClaireBoolean  
              /* noccur = 2 */
              /* For:7 */{ 
                var m *ClaireAny  
                _ = m
                g0515= CFALSE
                var m_support *ClaireList  
                m_support = ToClass(g0514).Instances
                m_len := m_support.Length()
                for i_it := 0; i_it < m_len; i_it++ { 
                  m = m_support.At(i_it)
                  PRINC("(addLog(")
                  F_princ_integer(ToMeasure(m).MIndex)
                  PRINC(",")
                  F_princ_float(ToMeasure(m).SumValue)
                  PRINC(",")
                  F_princ_float(ToMeasure(m).SumSquare)
                  PRINC(",")
                  F_princ_float(ToMeasure(m).NumValue)
                  PRINC(",")
                  F_princ_integer(n)
                  PRINC("))\n")
                  /* loop-8 */} 
                /* For-7 */} 
              if (g0515 == CTRUE) /* If:7 */{ 
                 /*v = Result, s =void*/

                break
                /* If-7 */} 
              /* Let-6 */} 
            /* loop-5 */} 
          /* For-4 */} 
        p.Fclose()
        Result = EVOID
        /* Let-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: logMeasure @ string (throw: true) 
func E_logMeasure_string (s EID) EID { 
    return /*(sm for logMeasure @ string= EID)*/ F_logMeasure_string(ToString(OBJ(s)) )} 
  
// adds a set of measures to a measure object (represented by its index)
/* {1} OPT.The go function for: addLog(i:integer,x:float,y:float,n:float,s:integer) [] */
func F_addLog_integer (i int,x float64,y float64,n float64,s int) EID { 
    var Result EID 
    if (Core.F_size_class(C_measure) == s) /* If:2 */{ 
      /* Let:3 */{ 
        var m *ClaireAny   = C_measure.Instances.At(i-1)
        /* noccur = 9 */
        ToMeasure(m).SumValue = (ToMeasure(m).SumValue+x)
        ToMeasure(m).SumSquare = (ToMeasure(m).SumSquare+y)
        /* update:4 */{ 
          var va_arg1 *Measure  
          var va_arg2 float64 
          va_arg1 = ToMeasure(m)
          va_arg2 = (ToMeasure(m).NumValue+n)
          /* ---------- now we compile update num_value(va_arg1) := va_arg2 ------- */
          va_arg1.NumValue = va_arg2
          Result = EID{C__FLOAT,FVAL(va_arg2)}
          /* update-4 */} 
        /* Let-3 */} 
      } else {
      Result = ToException(Core.C_general_error.Make(MakeString("logMeasure not compatible with current set (~A vs ~A)").Id(),MakeConstantList(MakeInteger(Core.F_size_class(C_measure)).Id(),MakeInteger(s).Id()).Id())).Close()
      /* If-2 */} 
    return Result} 
  
// The EID go function for: addLog @ integer (throw: true) 
func E_addLog_integer (i EID,x EID,y EID,n EID,s EID) EID { 
    return /*(sm for addLog @ integer= EID)*/ F_addLog_integer(INT(i),
      FLOAT(x),
      FLOAT(y),
      FLOAT(n),
      INT(s) )} 
  
//
// we use a counter object for the 5 interesting values  and
// we use the reified slot to store the counter (thus no profiling on reified)
// start time (1st entry)
// get & create if needed a PRcounter
/* {1} OPT.The go function for: PRget(p:property) [] */
func F_PRget_property (p *ClaireProperty ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var x *ClaireBoolean   = p.Reified
      /* noccur = 6 */
      if (x.Id().Isa.Id() == C_PRcount.Id()) /* If:3 */{ 
        Result = EID{x.Id(),0}
        /* If!3 */}  else if (x.Id() == CTRUE.Id()) /* If:3 */{ 
        Result = ToException(Core.C_general_error.Make(MakeString("[131] Cannot profile a reified property ~S").Id(),MakeConstantList(p.Id()).Id())).Close()
        } else {
        /* Let:4 */{ 
          var _CL_obj *PRcount   = To_PRcount(new(PRcount).Is(C_PRcount))
          /* noccur = 1 */
          x = ToBoolean(_CL_obj.Id())
          /* Let-4 */} 
        p.Reified = x
        Result = EID{x.Id(),0}
        /* If-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: PRget @ property (throw: true) 
func E_PRget_property (p EID) EID { 
    return /*(sm for PRget @ property= EID)*/ F_PRget_property(ToProperty(OBJ(p)) )} 
  
// get & create if needed a PRcounter
/* {1} OPT.The go function for: PRlook(p:property) [] */
func F_PRlook_property2 (p *ClaireProperty ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var g0520UU *PRcount  
      /* noccur = 1 */
      var g0520UU_try05213 EID 
      g0520UU_try05213 = F_PRget_property(p)
      /* ERROR PROTECTION INSERTED (g0520UU-Result) */
      if ErrorIn(g0520UU_try05213) {Result = g0520UU_try05213
      } else {
      g0520UU = To_PRcount(OBJ(g0520UU_try05213))
      Result = F_show_any(g0520UU.Id())
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: PRlook @ property (throw: true) 
func E_PRlook_property2 (p EID) EID { 
    return /*(sm for PRlook @ property= EID)*/ F_PRlook_property2(ToProperty(OBJ(p)) )} 
  
// show the profiler statistics on one property
/* {1} OPT.The go function for: PRshow(p:property) [] */
func F_PRshow_property (p *ClaireProperty ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var x *ClaireBoolean   = p.Reified
      /* noccur = 2 */
      if (x.Isa.IsIn(C_PRcount) == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var g0522 *PRcount   = To_PRcount(x.Id())
          /* noccur = 2 */
          Result = Core.F_print_any(p.Id())
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(": ")
          F_princ_integer(g0522.Rnum)
          PRINC(" calls -> ")
          F_princ_integer(g0522.Rtime)
          PRINC(" clock tics\n")
          Result = EVOID
          }
          /* Let-4 */} 
        } else {
        Result = EID{CFALSE.Id(),0}
        /* If-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: PRshow @ property (throw: true) 
func E_PRshow_property (p EID) EID { 
    return /*(sm for PRshow @ property= EID)*/ F_PRshow_property(ToProperty(OBJ(p)) )} 
  
// elapsed time
/* {1} OPT.The go function for: PRtime(p:property) [] */
func F_PRtime_property (p *ClaireProperty ) int { 
    // procedure body with s =  
var Result int 
    /* Let:2 */{ 
      var x *ClaireBoolean   = p.Reified
      /* noccur = 2 */
      if (x.Isa.IsIn(C_PRcount) == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var g0523 *PRcount   = To_PRcount(x.Id())
          /* noccur = 1 */
          Result = g0523.Rtime
          /* Let-4 */} 
        } else {
        Result = 0
        /* If-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: PRtime @ property (throw: false) 
func E_PRtime_property (p EID) EID { 
    return EID{C__INT,IVAL(/*(sm for PRtime @ property= integer)*/ F_PRtime_property(ToProperty(OBJ(p)) ))}} 
  
/* {1} OPT.The go function for: PRcounter(p:property) [] */
func F_PRcounter_property (p *ClaireProperty ) int { 
    // procedure body with s =  
var Result int 
    /* Let:2 */{ 
      var x *ClaireBoolean   = p.Reified
      /* noccur = 2 */
      if (x.Isa.IsIn(C_PRcount) == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var g0525 *PRcount   = To_PRcount(x.Id())
          /* noccur = 1 */
          Result = g0525.Rnum
          /* Let-4 */} 
        } else {
        Result = 0
        /* If-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: PRcounter @ property (throw: false) 
func E_PRcounter_property (p EID) EID { 
    return EID{C__INT,IVAL(/*(sm for PRcounter @ property= integer)*/ F_PRcounter_property(ToProperty(OBJ(p)) ))}} 
  
// show the profiler statistics on the 10 most important properties
/* {1} OPT.The go function for: PRshow(_CL_obj:void) [] */
func F_PRshow_void () EID { 
    var Result EID 
    /* Let:2 */{ 
      var l *ClaireList   = ToType(C_property.Id()).EmptyList()
      /* noccur = 11 */
      /* For:3 */{ 
        var g0527 *ClaireAny  
        _ = g0527
        Result= EID{CFALSE.Id(),0}
        var g0527_support *ClaireSet  
        g0527_support = C_property.Descendents
        for _,g0527 = range(g0527_support.Values)/* loop2:4 */{ 
          var void_try5 EID 
          _ = void_try5
          /* Let:5 */{ 
            var g0528 *ClaireBoolean  
            /* noccur = 2 */
            var g0528_try05306 EID 
            /* For:6 */{ 
              var p *ClaireAny  
              _ = p
              g0528_try05306= EID{CFALSE.Id(),0}
              var p_support *ClaireList  
              p_support = ToClass(g0527).Instances
              p_len := p_support.Length()
              for i_it := 0; i_it < p_len; i_it++ { 
                p = p_support.At(i_it)
                var void_try8 EID 
                _ = void_try8
                var g0531I *ClaireBoolean  
                var g0531I_try05328 EID 
                /* Let:8 */{ 
                  var g0533UU *ClaireAny  
                  /* noccur = 1 */
                  var g0533UU_try05349 EID 
                  /* Let:9 */{ 
                    var i int  = 1
                    /* noccur = 7 */
                    /* Let:10 */{ 
                      var g0529 int  = F_min_integer(10,l.Length())
                      /* noccur = 1 */
                      g0533UU_try05349= EID{CFALSE.Id(),0}
                      for (i <= g0529) /* while:11 */{ 
                        var void_try12 EID 
                        _ = void_try12
                        { 
                        var g0535I *ClaireBoolean  
                        var g0535I_try053612 EID 
                        if ((F_PRtime_property(ToProperty(p)) > F_PRtime_property(ToProperty(l.ValuesO()[i-1]))) || 
                            ((F_PRtime_property(ToProperty(p)) == F_PRtime_property(ToProperty(l.ValuesO()[i-1]))) && 
                                (F_PRcounter_property(ToProperty(p)) > F_PRcounter_property(ToProperty(l.ValuesO()[i-1]))))) /* If:12 */{ 
                          var l_try053713 EID 
                          l_try053713 = l.Nth_plus(i,p)
                          /* ERROR PROTECTION INSERTED (l-g0535I_try053612) */
                          if ErrorIn(l_try053713) {g0535I_try053612 = l_try053713
                          } else {
                          l = ToList(OBJ(l_try053713))
                          g0535I_try053612 = EID{l.Id(),0}
                          g0535I_try053612 = EID{CTRUE.Id(),0}
                          }
                          } else {
                          g0535I_try053612 = EID{CFALSE.Id(),0}
                          /* If-12 */} 
                        /* ERROR PROTECTION INSERTED (g0535I-void_try12) */
                        if ErrorIn(g0535I_try053612) {void_try12 = g0535I_try053612
                        } else {
                        g0535I = ToBoolean(OBJ(g0535I_try053612))
                        if (g0535I == CTRUE) /* If:12 */{ 
                           /*v = g0533UU_try05349, s =EID*/
g0533UU_try05349 = EID{CTRUE.Id(),0}
                          break
                          } else {
                          void_try12 = EID{CFALSE.Id(),0}
                          /* If-12 */} 
                        }
                        /* ERROR PROTECTION INSERTED (void_try12-void_try12) */
                        if ErrorIn(void_try12) {g0533UU_try05349 = void_try12
                        break
                        } else {
                        i = (i+1)
                        }
                        /* while-11 */} 
                      }
                      /* Let-10 */} 
                    /* Let-9 */} 
                  /* ERROR PROTECTION INSERTED (g0533UU-g0531I_try05328) */
                  if ErrorIn(g0533UU_try05349) {g0531I_try05328 = g0533UU_try05349
                  } else {
                  g0533UU = ANY(g0533UU_try05349)
                  g0531I_try05328 = EID{F_boolean_I_any(g0533UU).Id(),0}
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (g0531I-void_try8) */
                if ErrorIn(g0531I_try05328) {void_try8 = g0531I_try05328
                } else {
                g0531I = ToBoolean(OBJ(g0531I_try05328))
                if (g0531I == CTRUE) /* If:8 */{ 
                  void_try8 = EID{CNIL.Id(),0}
                  /* If!8 */}  else if (l.Length() < 10) /* If:8 */{ 
                  l = l.AddFast(p)
                  void_try8 = EID{l.Id(),0}
                  } else {
                  void_try8 = EID{CFALSE.Id(),0}
                  /* If-8 */} 
                }
                /* ERROR PROTECTION INSERTED (void_try8-g0528_try05306) */
                if ErrorIn(void_try8) {g0528_try05306 = void_try8
                g0528_try05306 = void_try8
                break
                } else {
                }
                /* loop-7 */} 
              /* For-6 */} 
            /* ERROR PROTECTION INSERTED (g0528-void_try5) */
            if ErrorIn(g0528_try05306) {void_try5 = g0528_try05306
            } else {
            g0528 = ToBoolean(OBJ(g0528_try05306))
            if (g0528 == CTRUE) /* If:6 */{ 
               /*v = Result, s =EID*/
Result = EID{g0528.Id(),0}
              break
              } else {
              void_try5 = EID{CFALSE.Id(),0}
              /* If-6 */} 
            }
            /* Let-5 */} 
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
      l.Shrink(10)
      /* For:3 */{ 
        var p *ClaireAny  
        _ = p
        Result= EID{CFALSE.Id(),0}
        for _,p = range(l.ValuesO())/* loop:4 */{ 
          var void_try5 EID 
          _ = void_try5
          if (F_PRcounter_property(ToProperty(p)) > 0) /* If:5 */{ 
            PRINC("-----------------------------------\n")
            void_try5 = F_PRshow_property(ToProperty(p))
            /* ERROR PROTECTION INSERTED (void_try5-void_try5) */
            if ErrorIn(void_try5) {Result = void_try5
            break
            } else {
            /* For:6 */{ 
              var p2 *ClaireAny  
              _ = p2
              void_try5= EID{CFALSE.Id(),0}
              var p2_support *ClaireSet  
              var p2_support_try05387 EID 
              p2_support_try05387 = Core.F_nth_table1(C_Reader_PRdependent,p)
              /* ERROR PROTECTION INSERTED (p2_support-void_try5) */
              if ErrorIn(p2_support_try05387) {void_try5 = p2_support_try05387
              } else {
              p2_support = ToSet(OBJ(p2_support_try05387))
              for _,p2 = range(p2_support.Values)/* loop2:7 */{ 
                var void_try8 EID 
                _ = void_try8
                if (F_PRtime_property(ToProperty(p2)) > 0) /* If:8 */{ 
                  PRINC("   * ")
                  void_try8 = F_PRshow_property(ToProperty(p2))
                  /* ERROR PROTECTION INSERTED (void_try8-void_try8) */
                  if ErrorIn(void_try8) {void_try5 = void_try8
                  break
                  } else {
                  PRINC("")
                  void_try8 = EVOID
                  }
                  } else {
                  void_try8 = EID{CFALSE.Id(),0}
                  /* If-8 */} 
                /* ERROR PROTECTION INSERTED (void_try8-void_try5) */
                if ErrorIn(void_try8) {void_try5 = void_try8
                void_try5 = void_try8
                break
                } else {
                }}
                /* loop-7 */} 
              /* For-6 */} 
            /* ERROR PROTECTION INSERTED (void_try5-void_try5) */
            if ErrorIn(void_try5) {Result = void_try5
            break
            } else {
            }}
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
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: PRshow @ void (throw: true) 
func E_PRshow_void (_CL_obj EID) EID { 
    return /*(sm for PRshow @ void= EID)*/ F_PRshow_void( )} 
  
// reuse from lexical_build in pretty.cl
// returns the list of properties that are used by a method
/* {1} OPT.The go function for: dependents(self:method) [] */
func F_dependents_method (self *ClaireMethod ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var p_out *ClaireSet   = ToType(C_property.Id()).EmptySet()
      /* noccur = 2 */
      /* For:3 */{ 
        var p *ClaireAny  
        _ = p
        Result= EID{CFALSE.Id(),0}
        var p_support *ClaireList  
        var p_support_try05394 EID 
        /* Let:4 */{ 
          var g0540UU *ClaireAny  
          /* noccur = 1 */
          var g0540UU_try05415 EID 
          g0540UU_try05415 = Core.F_CALL(C_Reader_dependents,ARGS(self.Formula.Body.ToEID()))
          /* ERROR PROTECTION INSERTED (g0540UU-p_support_try05394) */
          if ErrorIn(g0540UU_try05415) {p_support_try05394 = g0540UU_try05415
          } else {
          g0540UU = ANY(g0540UU_try05415)
          p_support_try05394 = Core.F_enumerate_any(g0540UU)
          }
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (p_support-Result) */
        if ErrorIn(p_support_try05394) {Result = p_support_try05394
        } else {
        p_support = ToList(OBJ(p_support_try05394))
        p_len := p_support.Length()
        for i_it := 0; i_it < p_len; i_it++ { 
          p = p_support.At(i_it)
          var g0542I *ClaireBoolean  
          /* Let:5 */{ 
            var g0543UU *ClaireAny  
            /* noccur = 1 */
            /* For:6 */{ 
              var r *ClaireAny  
              _ = r
              g0543UU= CFALSE.Id()
              for _,r = range(ToProperty(p).Restrictions.ValuesO())/* loop:7 */{ 
                if (C_method.Id() == r.Isa.Id()) /* If:8 */{ 
                   /*v = g0543UU, s =any*/
g0543UU = CTRUE.Id()
                  break
                  /* If-8 */} 
                /* loop-7 */} 
              /* For-6 */} 
            g0542I = F_boolean_I_any(g0543UU)
            /* Let-5 */} 
          if (g0542I == CTRUE) /* If:5 */{ 
            p_out.AddFast(p)
            /* If-5 */} 
          }
          /* loop-4 */} 
        /* For-3 */} 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = EID{p_out.Id(),0}
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: dependents @ method (throw: true) 
func E_dependents_method (self EID) EID { 
    return /*(sm for dependents @ method= EID)*/ F_dependents_method(ToMethod(OBJ(self)) )} 
  
// this is really cute ....   v3.2.58: fix typing
/* {1} OPT.The go function for: dependents(self:any) [] */
func F_dependents_any (self *ClaireAny ) EID { 
    var Result EID 
    if (self.Isa.IsIn(Language.C_Call) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0544 *Language.Call   = Language.To_Call(self)
        /* noccur = 2 */
        /* Let:4 */{ 
          var g0549UU *ClaireAny  
          /* noccur = 1 */
          var g0549UU_try05505 EID 
          g0549UU_try05505 = F_dependents_any(g0544.Args.Id())
          /* ERROR PROTECTION INSERTED (g0549UU-Result) */
          if ErrorIn(g0549UU_try05505) {Result = g0549UU_try05505
          } else {
          g0549UU = ANY(g0549UU_try05505)
          Result = Core.F_CALL(ToProperty(C_add.Id()),ARGS(g0549UU.ToEID(),EID{g0544.Selector.Id(),0}))
          }
          /* Let-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(C_Instruction) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0545 *ClaireInstruction   = To_Instruction(self)
        /* noccur = 2 */
        /* Let:4 */{ 
          var s *ClaireSet   = ToType(C_property.Id()).EmptySet()
          /* noccur = 3 */
          /* For:5 */{ 
            var sl *ClaireAny  
            _ = sl
            Result= EID{CFALSE.Id(),0}
            for _,sl = range(g0545.Isa.Slots.ValuesO())/* loop:6 */{ 
              var void_try7 EID 
              _ = void_try7
              var s_try05517 EID 
              /* Let:7 */{ 
                var g0552UU *ClaireAny  
                /* noccur = 1 */
                var g0552UU_try05538 EID 
                g0552UU_try05538 = Core.F_CALL(C_Reader_dependents,ARGS(Core.F_get_slot(ToSlot(sl),ToObject(g0545.Id())).ToEID()))
                /* ERROR PROTECTION INSERTED (g0552UU-s_try05517) */
                if ErrorIn(g0552UU_try05538) {s_try05517 = g0552UU_try05538
                } else {
                g0552UU = ANY(g0552UU_try05538)
                s_try05517 = EID{Core.F_U_type(ToType(s.Id()),ToType(g0552UU)).Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (s-Result) */
              if ErrorIn(s_try05517) {Result = s_try05517
              Result = s_try05517
              break
              } else {
              s = ToSet(OBJ(s_try05517))
              void_try7 = EID{s.Id(),0}
              }
              /* loop-6 */} 
            /* For-5 */} 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          Result = EID{s.Id(),0}
          }
          /* Let-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(C_list) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0546 *ClaireList   = ToList(self)
        /* noccur = 1 */
        /* Let:4 */{ 
          var s *ClaireSet   = ToType(C_property.Id()).EmptySet()
          /* noccur = 3 */
          /* For:5 */{ 
            var x *ClaireAny  
            _ = x
            Result= EID{CFALSE.Id(),0}
            var x_support *ClaireList  
            x_support = g0546
            x_len := x_support.Length()
            for i_it := 0; i_it < x_len; i_it++ { 
              x = x_support.At(i_it)
              var void_try7 EID 
              _ = void_try7
              var s_try05547 EID 
              /* Let:7 */{ 
                var g0555UU *ClaireAny  
                /* noccur = 1 */
                var g0555UU_try05568 EID 
                g0555UU_try05568 = Core.F_CALL(C_Reader_dependents,ARGS(x.ToEID()))
                /* ERROR PROTECTION INSERTED (g0555UU-s_try05547) */
                if ErrorIn(g0555UU_try05568) {s_try05547 = g0555UU_try05568
                } else {
                g0555UU = ANY(g0555UU_try05568)
                s_try05547 = EID{Core.F_U_type(ToType(s.Id()),ToType(g0555UU)).Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (s-Result) */
              if ErrorIn(s_try05547) {Result = s_try05547
              Result = s_try05547
              break
              } else {
              s = ToSet(OBJ(s_try05547))
              void_try7 = EID{s.Id(),0}
              }
              /* loop-6 */} 
            /* For-5 */} 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          Result = EID{s.Id(),0}
          }
          /* Let-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(C_property) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0547 *ClaireProperty   = ToProperty(self)
        /* noccur = 1 */
        Result = EID{MakeSet(ToType(C_property.Id()),g0547.Id()).Id(),0}
        /* Let-3 */} 
      } else {
      Result = EID{ToType(C_property.Id()).EmptySet().Id(),0}
      /* If-2 */} 
    return Result} 
  
// The EID go function for: dependents @ any (throw: true) 
func E_dependents_any (self EID) EID { 
    return /*(sm for dependents @ any= EID)*/ F_dependents_any(ANY(self) )} 
  
// used to set up the dependence
/* {1} OPT.The go function for: PRdepends(p:property,p2:property) [] */
func F_PRdepends_property (p *ClaireProperty ,p2 *ClaireProperty ) EID { 
    var Result EID 
    Result = Core.F_add_table(C_Reader_PRdependent,p.Id(),p2.Id())
    return Result} 
  
// The EID go function for: PRdepends @ property (throw: true) 
func E_PRdepends_property (p EID,p2 EID) EID { 
    return /*(sm for PRdepends @ property= EID)*/ F_PRdepends_property(ToProperty(OBJ(p)),ToProperty(OBJ(p2)) )} 
  
// end of file