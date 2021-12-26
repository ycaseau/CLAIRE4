/***** CLAIRE Compilation of file /Users/ycaseau/Dropbox/src/clairev4.03/src/meta/inspect.cl 
         [version 4.0.04 / safety 5] Sunday 12-26-2021 17:16:10 *****/

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
/* {1} The go function for: top_level(r:meta_reader) [status=1] */
func (r *MetaReader ) TopLevel () EID { 
    var Result EID 
    { var res *ClaireAny   = MakeInteger(0).Id()
      Result= EID{CFALSE.Id(),0}
      for (res != C_q.Id()) { 
        /* While stat, v:"Result" loop:true */
        var loop_1 EID 
        _ = loop_1
        { 
        { var arg_2 *ClaireString  
          _ = arg_2
          if (ToInteger(C_Reader_TopLevelMode.Value).Value == 1) { 
            arg_2 = ClEnv.Module_I.Name.String_I()
            }  else if (ToInteger(C_Reader_TopLevelMode.Value).Value == 2) { 
            arg_2 = MakeString("debug")
            } else {
            arg_2 = MakeString("inspect")
            } 
          F_princ_string(arg_2)
          } 
        PRINC("> ")
        /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
        { 
          h_index := ClEnv.Index
          h_base := ClEnv.Base
          r.Toplevel = CTRUE
          /*boolean->boolean*/if (ClEnv.CountCall > 0) { 
            ClEnv.CountCall = 1
            /*integer->integer*/} 
          var try_3 EID 
          /*g_try(v2:"try_3",loop:false) */
          try_3 = r.Nextunit()
          /* ERROR PROTECTION INSERTED (res-loop_1) */
          if ErrorIn(try_3) {loop_1 = try_3
          } else {
          res = ANY(try_3)
          loop_1 = res.ToEID()
          if (ToInteger(C_Reader_TopLevelMode.Value).Value == 1) { 
            ClEnv.Index = 20
            /*integer->integer*/} 
          if ((ToInteger(C_Reader_TopLevelMode.Value).Value == 3) && 
              (res != C_q.Id())) { 
            loop_1 = F_inspect_loop_any(res,ToList(C_Reader_InspectStack.Value))
            } else {
            /*g_try(v2:"loop_1",loop:false) */
            if (ToInteger(C_Reader_TopLevelMode.Value).Value == 1) { 
              PRINC("eval[")
              /*g_try(v2:"loop_1",loop:false) */
              { var arg_4 *ClaireAny  
                _ = arg_4
                arg_4 = MakeInteger((ToInteger(C_Reader_TopCount.Value).Value+1)).Id()
                C_Reader_TopCount.Value = arg_4
                loop_1 = Core.F_print_any(arg_4)
                } 
              /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
              if !ErrorIn(loop_1) {
              PRINC("]> ")
              loop_1 = EVOID
              }
              } else {
              PRINC("> ")
              loop_1 = EVOID
              } 
            /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
            if !ErrorIn(loop_1) {
            var try_6 EID 
            /*g_try(v2:"try_6",loop:false) */
            try_6 = EVAL(res)
            /* ERROR PROTECTION INSERTED (res-loop_1) */
            if ErrorIn(try_6) {loop_1 = try_6
            } else {
            res = ANY(try_6)
            loop_1 = res.ToEID()
            if (res != C_q.Id()) { 
              /*g_try(v2:"loop_1",loop:false) */
              loop_1 = Core.F_CALL(C_print,ARGS(res.ToEID()))
              /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
              if !ErrorIn(loop_1) {
              PRINC("\n")
              loop_1 = EVOID
              }
              } else {
              loop_1 = EID{CFALSE.Id(),0}
              } 
            }}
            } 
          }
          if ErrorIn(loop_1){ 
            ClEnv.Index = h_index
            ClEnv.Base = h_base
            var g0142I *ClaireBoolean  
            { var e *ClaireException   = ClEnv.Exception_I
              g0142I = MakeBoolean((e.Isa.Id() == C_system_error.Id()) && (ToSystemError(e.Id()).Index == -1))
              } 
            if (g0142I == CTRUE) { 
              C_Reader_TopLevelMode.Value = MakeInteger(1).Id()
              res = C_q.Id()
              loop_1 = res.ToEID()
              } else {
              r.RestoreState()
              if (r.External.Value != MakeString("toplevel").Value) { 
                PRINC("---- file: ")
                F_princ_string(r.External)
                PRINC(", line: ")
                F_princ_integer(ClEnv.NLine)
                PRINC("\n")
                } 
              /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
              loop_1 = F_debug_if_possible_void()
              /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
              if ErrorIn(loop_1) {Result = loop_1
              break
              } else {
              PRINC("\n")
              loop_1 = EVOID
              }
              } 
            } 
          } 
        /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
        if ErrorIn(loop_1) {Result = loop_1
        break
        } else {
        if ((ToInteger(C_Reader_TopLevelMode.Value).Value != 1) && 
            (res == C_q.Id())) { 
          if (ToInteger(C_Reader_TopLevelMode.Value).Value == 2) { 
            ClEnv.Index = ToInteger(C_Reader_TopIndex.Value).Value
            /*integer->integer*/ClEnv.Base = ToInteger(C_Reader_TopBase.Value).Value
            /*integer->integer*/ClEnv.Trace_I = 1
            /*integer->integer*/ClEnv.Debug_I = ToInteger(C_Reader_TopDebug.Value).Value
            /*integer->integer*/} 
          res = CNULL
          C_Reader_TopLevelMode.Value = MakeInteger(1).Id()
          } 
        }
        /* try?:false, v2:"v_while3" loop will be:tuple("Result", EID) */
        } 
      }
      } 
    return Result} 
  
// The EID go function for: top_level @ meta_reader (throw: true) 
func E_Reader_top_level_meta_reader (r EID) EID { 
    return ToMetaReader(OBJ(r)).TopLevel( )} 
  
//        exit(1)) ]
// start a debug loop - aha 
/* {1} The go function for: debugLoop(r:meta_reader) [status=0] */
func (r *MetaReader ) DebugLoop ()  { 
    // procedure body with s = void 
C_Reader_TopDebug.Value = MakeInteger(0).Id()
    C_Reader_TopBase.Value = MakeInteger(ClEnv.Base).Id()
    C_Reader_TopIndex.Value = MakeInteger(ClEnv.Index).Id()
    r.Toplevel = CTRUE
    /*boolean->boolean*/C_Reader_TopLevelMode.Value = MakeInteger(2).Id()
    PRINC("--------------- Debug -------------------\n")
    } 
  
// The EID go function for: debugLoop @ meta_reader (throw: false) 
func E_Reader_debugLoop_meta_reader (r EID) EID { 
    ToMetaReader(OBJ(r)).DebugLoop( )
    return EVOID} 
  
// starts an inspector  on a list
/* {1} The go function for: inspect_system(l:list) [status=0] */
func F_Reader_inspect_system_list2 (l *ClaireList )  { 
    // procedure body with s = void 
C_Reader_InspectStack.Value = l.Id()
    if (ToInteger(C_Reader_TopLevelMode.Value).Value == 2) { 
      ClEnv.Trace_I = 1
      /*integer->integer*/} 
    C_Reader_TopLevelMode.Value = MakeInteger(3).Id()
    } 
  
// The EID go function for: inspect_system @ list (throw: false) 
func E_Reader_inspect_system_list2 (l EID) EID { 
    F_Reader_inspect_system_list2(ToList(OBJ(l)) )
    return EVOID} 
  
// INSPECT   
// simple main (to be enriched later)
/* {1} The go function for: simple_main(_CL_obj:void) [status=1] */
func F_Reader_simple_main_void () EID { 
    var Result EID 
    /*g_try(v2:"Result",loop:true) */
    Result = C_reader.TopLevel()
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("[regular exit] Bye.\n")
    Result = EVOID
    }
    return Result} 
  
// The EID go function for: simple_main @ void (throw: true) 
func E_Reader_simple_main_void (_CL_obj EID) EID { 
    return F_Reader_simple_main_void( )} 
  
// unclear that we need this (simple_main is used with -cm)
//
// *********************************************************************
// *      Part 2: Inspection                                           *
// *********************************************************************
// this is the method that the user calls
//
/* {1} The go function for: inspect(self:any) [status=1] */
func F_inspect_any (self *ClaireAny ) EID { 
    var Result EID 
    { var m0 *ClaireModule   = ClEnv.Module_I
      _ = m0
      { var ix int  = 0
        _ = ix
        /*g_try(v2:"Result",loop:true) */
        if (self.Isa.IsIn(C_list) == CTRUE) { 
          { var g0143 *ClaireList   = ToList(self)
            { var i int  = 1
              { var g0144 int  = g0143.Length()
                _ = g0144
                Result= EID{CFALSE.Id(),0}
                for (i <= g0144) { 
                  /* While stat, v:"Result" loop:true */
                  var loop_1 EID 
                  _ = loop_1
                  { 
                  /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
                  F_princ_integer(i)
                  PRINC(": ")
                  /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
                  loop_1 = Core.F_CALL(C_print,ARGS(g0143.At(i-1).ToEID()))
                  /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
                  if ErrorIn(loop_1) {Result = loop_1
                  break
                  } else {
                  PRINC("\n")
                  loop_1 = EVOID
                  }
                  /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
                  if ErrorIn(loop_1) {Result = loop_1
                  break
                  } else {
                  i = (i+1)
                  }
                  /* try?:false, v2:"v_while8" loop will be:tuple("Result", EID) */
                  } 
                }
                } 
              } 
            } 
          }  else if (Core.F__Z_any1(self,C_object) == CTRUE) { 
          { var g0145 *ClaireObject   = ToObject(self)
            { 
              var rel *ClaireSlot  
              _ = rel
              var rel_iter *ClaireAny  
              Result= EID{CFALSE.Id(),0}
              var rel_support *ClaireList  
              rel_support = g0145.Id().Isa.Slots
              for _,rel_iter = range(rel_support.ValuesO()){ 
                rel = ToSlot(rel_iter)
                var loop_2 EID 
                _ = loop_2
                /*g_try(v2:"loop_2",loop:tuple("Result", EID)) */
                { var m *ClaireModule   = rel.Selector.Name.Module_I()
                  ix = (ix+1)
                  /*g_try(v2:"loop_2",loop:tuple("Result", EID)) */
                  if ((m.Id() == m0.Id()) || 
                      ((m.Id() == C_claire.Id()) || 
                        (ToBoolean(C__starshowall_star.Value) == CTRUE))) { 
                    { var val *ClaireAny   = Core.F_get_slot(rel,g0145)
                      F_princ_integer(ix)
                      PRINC(": ")
                      /*g_try(v2:"loop_2",loop:tuple("Result", EID)) */
                      loop_2 = Core.F_print_any(rel.Selector.Id())
                      /* ERROR PROTECTION INSERTED (loop_2-loop_2) */
                      if ErrorIn(loop_2) {Result = loop_2
                      break
                      } else {
                      PRINC(" = ")
                      /*g_try(v2:"loop_2",loop:tuple("Result", EID)) */
                      if (val.Isa.IsIn(C_list) == CTRUE) { 
                        { var g0146 *ClaireList   = ToList(val)
                          if (g0146.Length() < 10) { 
                            loop_2 = Language.F_pretty_print_any(g0146.Id())
                            } else {
                            /*g_try(v2:"loop_2",loop:tuple("Result", EID)) */
                            { var arg_3 *ClaireList  
                              _ = arg_3
                              { var i_bag *ClaireList   = ToType(CEMPTY.Id()).EmptyList()
                                { var i int  = 1
                                  { var g0147 int  = 9
                                    _ = g0147
                                    for (i <= g0147) { 
                                      /* While stat, v:"arg_3" loop:tuple("Result", EID) */
                                      i_bag.AddFast(g0146.At(i-1))/*t=any,s=void*/
                                      i = (i+1)
                                      /* try?:false, v2:"v_while18" loop will be:tuple("arg_3", void) */
                                      } 
                                    } 
                                  } 
                                arg_3 = i_bag
                                } 
                              loop_2 = Language.F_pretty_print_any(arg_3.Id())
                              } 
                            /* ERROR PROTECTION INSERTED (loop_2-loop_2) */
                            if ErrorIn(loop_2) {Result = loop_2
                            break
                            } else {
                            /*g_try(v2:"loop_2",loop:tuple("Result", EID)) */
                            loop_2 = Language.F_pretty_print_any(MakeString("...").Id())
                            /* ERROR PROTECTION INSERTED (loop_2-loop_2) */
                            if ErrorIn(loop_2) {Result = loop_2
                            break
                            } else {
                            }}
                            } 
                          } 
                        } else {
                        loop_2 = Language.F_pretty_print_any(val)
                        } 
                      /* ERROR PROTECTION INSERTED (loop_2-loop_2) */
                      if ErrorIn(loop_2) {Result = loop_2
                      break
                      } else {
                      PRINC("\n")
                      loop_2 = EVOID
                      }}
                      } 
                    } else {
                    loop_2 = EID{CFALSE.Id(),0}
                    } 
                  /* ERROR PROTECTION INSERTED (loop_2-loop_2) */
                  if ErrorIn(loop_2) {Result = loop_2
                  break
                  } else {
                  }
                  } 
                /* ERROR PROTECTION INSERTED (loop_2-Result) */
                if ErrorIn(loop_2) {Result = loop_2
                break
                } else {
                }
                } 
              } 
            } 
          } else {
          /*g_try(v2:"Result",loop:true) */
          Result = Language.F_pretty_print_any(self)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC("\n")
          Result = EVOID
          }
          } 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        F_Reader_inspect_system_list2(MakeConstantList(self))
        Result = EID{C_None.Id(),0}
        }
        } 
      } 
    return Result} 
  
// The EID go function for: inspect @ any (throw: true) 
func E_inspect_any (self EID) EID { 
    return F_inspect_any(ANY(self) )} 
  
// this is the inspect top_level
//
/* {1} The go function for: inspect_loop(%read:any,old:list) [status=1] */
func F_inspect_loop_any (_Zread *ClaireAny ,old *ClaireList ) EID { 
    var Result EID 
    /*g_try(v2:"Result",loop:true) */
    { var self *ClaireAny   = old.At(1-1)
      var g0151I *ClaireBoolean  
      if (_Zread.Isa.IsIn(Language.C_Call) == CTRUE) { 
        { var g0150 *Language.Call   = Language.To_Call(_Zread)
          _ = g0150
          g0151I = Equal(g0150.Selector.Id(),C_put.Id())
          } 
        } else {
        g0151I = CFALSE
        } 
      if (g0151I == CTRUE) { 
        { var n int  = ToInteger(ToList(OBJ(Core.F_CALL(C_args,ARGS(_Zread.ToEID())))).At(1-1)).Value
          { var s *ClaireSymbol  
            _ = s
            var try_1 EID 
            /*g_try(v2:"try_1",loop:false) */
            try_1 = Language.F_extract_symbol_any(ToList(OBJ(Core.F_CALL(C_args,ARGS(_Zread.ToEID())))).At(2-1))
            /* ERROR PROTECTION INSERTED (s-Result) */
            if ErrorIn(try_1) {Result = try_1
            } else {
            s = ToSymbol(OBJ(try_1))
            /*g_try(v2:"Result",loop:true) */
            if (C_integer.Id() != OWNER(EID{C__INT,IVAL(n)}).Id()) { 
              Result = ToException(Core.C_general_error.Make(MakeString("[128] ~S should be an integer").Id(),MakeConstantList(MakeInteger(n).Id()).Id())).Close()
              } else {
              Result = EID{CFALSE.Id(),0}
              } 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            { var val *ClaireAny  
              var try_2 EID 
              /*g_try(v2:"try_2",loop:false) */
              try_2 = F_get_from_integer_any(self,n)
              /* ERROR PROTECTION INSERTED (val-Result) */
              if ErrorIn(try_2) {Result = try_2
              } else {
              val = ANY(try_2)
              /*g_try(v2:"Result",loop:true) */
              { 
                var va_arg1 *Core.GlobalVariable  
                var va_arg2 *ClaireAny  
                var try_3 EID 
                /*g_try(v2:"try_3",loop:false) */
                try_3 = Core.F_new_class2(Core.C_global_variable,s)
                /* ERROR PROTECTION INSERTED (va_arg1-Result) */
                if ErrorIn(try_3) {Result = try_3
                } else {
                va_arg1 = Core.ToGlobalVariable(OBJ(try_3))
                va_arg2 = val
                va_arg1.Value = va_arg2
                /*any->any*/Result = va_arg2.ToEID()
                }
                } 
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              /*g_try(v2:"Result",loop:true) */
              Result = F_inspect_any(val)
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              old = F_cons_any(val,old)
              Result = EID{old.Id(),0}
              }}
              }
              } 
            }
            }
            } 
          } 
        }  else if (_Zread == C_up.Id()) { 
        if (old.Length() > 1) { 
          var try_4 EID 
          /*g_try(v2:"try_4",loop:true) */
          try_4 = old.Cdr()
          /* ERROR PROTECTION INSERTED (old-Result) */
          if ErrorIn(try_4) {Result = try_4
          } else {
          old = ToList(OBJ(try_4))
          Result = EID{old.Id(),0}
          Result = F_inspect_any(old.At(1-1))
          }
          } else {
          Result = EID{CFALSE.Id(),0}
          } 
        }  else if (C_integer.Id() == _Zread.Isa.Id()) { 
        { var val *ClaireAny  
          _ = val
          var try_5 EID 
          /*g_try(v2:"try_5",loop:false) */
          try_5 = F_get_from_integer_any(self,ToInteger(_Zread).Value)
          /* ERROR PROTECTION INSERTED (val-Result) */
          if ErrorIn(try_5) {Result = try_5
          } else {
          val = ANY(try_5)
          old = F_cons_any(val,old)
          Result = F_inspect_any(val)
          }
          } 
        }  else if (_Zread.Isa.IsIn(C_thing) == CTRUE) { 
        old = F_cons_any(_Zread,old)
        Result = F_inspect_any(_Zread)
        } else {
        PRINC("=> given to inspector is wrong.\n")
        Result = EVOID
        } 
      } 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    F_Reader_inspect_system_list2(old)
    Result = EVOID
    }
    return Result} 
  
// The EID go function for: inspect_loop @ any (throw: true) 
func E_inspect_loop_any (_Zread EID,old EID) EID { 
    return F_inspect_loop_any(ANY(_Zread),ToList(OBJ(old)) )} 
  
// get the information bound to the index
//
/* {1} The go function for: get_from_integer(self:any,n:integer) [status=1] */
func F_get_from_integer_any (self *ClaireAny ,n int) EID { 
    var Result EID 
    if (self.Isa.IsIn(C_list) == CTRUE) { 
      if ((n > 0) && 
          (n <= ToList(self).Length())) { 
        Result = Core.F_CALL(C_nth,ARGS(self.ToEID(),EID{C__INT,IVAL(n)}))
        } else {
        /*g_try(v2:"Result",loop:true) */
        F_princ_integer(n)
        PRINC(" in not a good index for ")
        /*g_try(v2:"Result",loop:true) */
        Result = Core.F_CALL(C_print,ARGS(self.ToEID()))
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC(".\n")
        Result = EVOID
        }
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        Result = self.ToEID()
        }
        } 
      } else {
      { var l *ClaireList   = self.Isa.Slots
        if ((n > 0) && 
            (n <= l.Length())) { 
          Result = Core.F_SUPER(C_get, C_slot, ARGS(l.ValuesO()[n-1].ToEID(),self.ToEID()))
          } else {
          /*g_try(v2:"Result",loop:true) */
          F_princ_integer(n)
          PRINC(" is not a good index for ")
          /*g_try(v2:"Result",loop:true) */
          Result = Core.F_CALL(C_print,ARGS(self.ToEID()))
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(".\n")
          Result = EVOID
          }
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          Result = self.ToEID()
          }
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: get_from_integer @ any (throw: true) 
func E_get_from_integer_any (self EID,n EID) EID { 
    return F_get_from_integer_any(ANY(self),INT(n) )} 
  
// *********************************************************************
// *      Part 2: Trace methods                                        *
// *********************************************************************
// instrument the code generated from the rules
// this is the control method to CLAIRE tracer
// notice that trace(where) activates the call_count
/* {1} The go function for: iClaire/trace_on(self:any) [status=1] */
func F_trace_on_any (self *ClaireAny ) EID { 
    var Result EID 
    /*g_try(v2:"Result",loop:true) */
    if (self.Isa.IsIn(C_property) == CTRUE) { 
      { var g0152 *ClaireProperty   = ToProperty(self)
        if (g0152.Id() == Core.C_spy.Id()) { 
          { var m *ClaireAny   = Core.F__at_property1(Core.C_spy,C_void).Id()
            if (F_boolean_I_any(m) == CTRUE) { 
              Result = F_store_object(ToObject(ClEnv.Id()),
                16,
                C_object,
                m,
                CFALSE).ToEID()
              } else {
              Result = EID{CFALSE.Id(),0}
              } 
            } 
          }  else if (g0152.Id() == C_where.Id()) { 
          Result = Core.F_write_property(C_Kernel_call_count,ToObject(ClEnv.Id()),MakeInteger(1).Id())
          } else {
          { 
            var va_arg1 *ClaireProperty  
            var va_arg2 int 
            va_arg1 = g0152
            va_arg2 = (5-ClEnv.Verbose)
            va_arg1.Trace_I = va_arg2
            /*integer->integer*/Result = EID{C__INT,IVAL(va_arg2)}
            } 
          } 
        } 
      }  else if (self.Isa.IsIn(C_environment) == CTRUE) { 
      { 
        var va_arg1 *ClaireEnvironment  
        var va_arg2 int 
        va_arg1 = ClEnv
        va_arg2 = 1
        va_arg1.Trace_I = va_arg2
        /*integer->integer*/Result = EID{C__INT,IVAL(va_arg2)}
        } 
      }  else if (self.Isa.IsIn(C_module) == CTRUE) { 
      { var g0154 *ClaireModule   = ToModule(self)
        if (g0154.Status > 2) { 
          g0154.Status = 4
          /*integer->integer*/} 
        { 
          var m *ClaireAny  
          _ = m
          Result= EID{CFALSE.Id(),0}
          var m_support *ClaireList  
          m_support = g0154.Parts
          m_len := m_support.Length()
          for i_it := 0; i_it < m_len; i_it++ { 
            m = m_support.At(i_it)
            var loop_1 EID 
            _ = loop_1
            /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
            loop_1 = F_trace_on_any(m)
            /* ERROR PROTECTION INSERTED (loop_1-Result) */
            if ErrorIn(loop_1) {Result = loop_1
            break
            } else {
            }
            } 
          } 
        } 
      }  else if (C_port.Id() == self.Isa.Id()) { 
      { var g0155 *ClairePort   = ToPort(self)
        _ = g0155
        { 
          var va_arg1 *ClaireEnvironment  
          var va_arg2 *ClairePort  
          va_arg1 = ClEnv
          va_arg2 = g0155
          va_arg1.Ctrace = va_arg2
          /*port->port*/Result = va_arg2.ToEID()
          } 
        } 
      }  else if (C_string.Id() == self.Isa.Id()) { 
      { var g0156 *ClaireString   = ToString(self)
        _ = g0156
        { 
          var va_arg1 *ClaireEnvironment  
          var va_arg2 *ClairePort  
          va_arg1 = ClEnv
          var try_2 EID 
          /*g_try(v2:"try_2",loop:false) */
          try_2 = F_fopen_string(g0156,MakeString("w"))
          /* ERROR PROTECTION INSERTED (va_arg2-Result) */
          if ErrorIn(try_2) {Result = try_2
          } else {
          va_arg2 = ToPort(OBJ(try_2))
          va_arg1.Ctrace = va_arg2
          /*port->port*/Result = va_arg2.ToEID()
          }
          } 
        } 
      }  else if (C_integer.Id() == self.Isa.Id()) { 
      { var g0157 int  = ToInteger(self).Value
        _ = g0157
        { 
          var va_arg1 *ClaireEnvironment  
          var va_arg2 int 
          va_arg1 = ClEnv
          va_arg2 = g0157
          va_arg1.Verbose = va_arg2
          /*integer->integer*/Result = EID{C__INT,IVAL(va_arg2)}
          } 
        } 
      } else {
      Result = ToException(Core.C_general_error.Make(MakeString("[129] trace not implemented on ~S\n").Id(),MakeConstantList(self).Id())).Close()
      } 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    Result = self.ToEID()
    }
    return Result} 
  
// The EID go function for: iClaire/trace_on @ any (throw: true) 
func E_trace_on_any (self EID) EID { 
    return F_trace_on_any(ANY(self) )} 
  
/* {1} The go function for: untrace(self:any) [status=1] */
func F_untrace_any (self *ClaireAny ) EID { 
    var Result EID 
    /*g_try(v2:"Result",loop:true) */
    if (self.Isa.IsIn(C_property) == CTRUE) { 
      { var g0159 *ClaireProperty   = ToProperty(self)
        if (g0159.Id() == Core.C_spy.Id()) { 
          { 
            var va_arg1 *ClaireEnvironment  
            var va_arg2 *ClaireObject  
            va_arg1 = ClEnv
            va_arg2 = ToObject(CNULL)
            va_arg1.Spy_I = va_arg2
            /*object->object*/Result = EID{va_arg2.Id(),0}
            } 
          }  else if (g0159.Id() == C_where.Id()) { 
          Result = Core.F_write_property(C_Kernel_call_count,ToObject(ClEnv.Id()),MakeInteger(-1).Id())
          } else {
          { 
            var va_arg1 *ClaireProperty  
            var va_arg2 int 
            va_arg1 = g0159
            va_arg2 = 0
            va_arg1.Trace_I = va_arg2
            /*integer->integer*/Result = EID{C__INT,IVAL(va_arg2)}
            } 
          } 
        } 
      }  else if (self.Isa.IsIn(C_environment) == CTRUE) { 
      { 
        var va_arg1 *ClaireEnvironment  
        var va_arg2 int 
        va_arg1 = ClEnv
        va_arg2 = 0
        va_arg1.Trace_I = va_arg2
        /*integer->integer*/Result = EID{C__INT,IVAL(va_arg2)}
        } 
      }  else if (self.Isa.IsIn(C_module) == CTRUE) { 
      { var g0161 *ClaireModule   = ToModule(self)
        if (g0161.Status == 4) { 
          g0161.Status = 3
          /*integer->integer*/} 
        { 
          var m *ClaireAny  
          _ = m
          Result= EID{CFALSE.Id(),0}
          var m_support *ClaireList  
          m_support = g0161.Parts
          m_len := m_support.Length()
          for i_it := 0; i_it < m_len; i_it++ { 
            m = m_support.At(i_it)
            var loop_1 EID 
            _ = loop_1
            /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
            loop_1 = F_untrace_any(m)
            /* ERROR PROTECTION INSERTED (loop_1-Result) */
            if ErrorIn(loop_1) {Result = loop_1
            break
            } else {
            }
            } 
          } 
        } 
      }  else if (C_port.Id() == self.Isa.Id()) { 
      { 
        var va_arg1 *ClaireEnvironment  
        var va_arg2 *ClairePort  
        va_arg1 = ClEnv
        va_arg2 = ToPort(C_stdout.Value)
        va_arg1.Ctrace = va_arg2
        /*port->port*/Result = va_arg2.ToEID()
        } 
      } else {
      Result = ToException(Core.C_general_error.Make(MakeString("[130] untrace not implemented on ~S\n").Id(),MakeConstantList(self).Id())).Close()
      } 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    Result = self.ToEID()
    }
    return Result} 
  
// The EID go function for: untrace @ any (throw: true) 
func E_untrace_any (self EID) EID { 
    return F_untrace_any(ANY(self) )} 
  
// a filter to restrict the impact of spy
// we put the special value nil (emply list of demons => OK) to mark that spying
// should be waken up on properties from l
/* {1} The go function for: spy(l:listargs) [status=0] */
func F_spy_listargs2_Reader (l *ClaireList )  { 
    // procedure body with s = void 
{ var m *ClaireAny   = Core.F__at_property1(Core.C_spy,C_void).Id()
      if (F_boolean_I_any(m) == CTRUE) { 
        F_store_object(ToObject(ClEnv.Id()),
          16,
          C_object,
          m,
          CFALSE)
        { 
          var g0164 *ClaireClass  
          _ = g0164
          var g0164_iter *ClaireAny  
          var g0164_support *ClaireSet  
          g0164_support = C_property.Descendents
          for i_it := 0; i_it < g0164_support.Count; i_it++ { 
            g0164_iter = g0164_support.At(i_it)
            g0164 = ToClass(g0164_iter)
            { var g0165 *ClaireBoolean  
              { 
                var f *ClaireProperty  
                _ = f
                var f_iter *ClaireAny  
                g0165= CFALSE
                var f_support *ClaireList  
                f_support = g0164.Instances
                f_len := f_support.Length()
                for i_it := 0; i_it < f_len; i_it++ { 
                  f_iter = f_support.At(i_it)
                  f = ToProperty(f_iter)
                  if (ToList(l.Id()).Memq(f.Id()) == CTRUE) { 
                    f.IfWrite = CNIL.Id()
                    /*any->any*/} 
                  } 
                } 
              if (g0165 == CTRUE) { 
                
                break
                } 
              } 
            } 
          } 
        } 
      } 
    } 
  
// The EID go function for: spy @ listargs (throw: false) 
func E_spy_listargs2_Reader (l EID) EID { 
    F_spy_listargs2_Reader(ToList(OBJ(l)) )
    return EVOID} 
  
// used to trace the trigger of a rule
/* {1} The go function for: trace_rule(R:relation,s:string,x:any,y:any,u:any,v:any) [status=1] */
func F_trace_rule_relation (R *ClaireRelation ,s *ClaireString ,x *ClaireAny ,y *ClaireAny ,u *ClaireAny ,v *ClaireAny ) EID { 
    var Result EID 
    if ((C_if_write.Trace_I+ClEnv.Verbose) >= 5) { 
      { var p *ClaireAny   = Core.F_get_property(C_ctrace,ToObject(ClEnv.Id()))
        if (p != CNULL) { 
          p = ToPort(p).UseAsOutput().Id()
          } 
        /*g_try(v2:"Result",loop:true) */
        PRINC("--- the rule ")
        F_princ_string(s)
        PRINC(" is triggered for (")
        /*g_try(v2:"Result",loop:true) */
        Result = Core.F_CALL(C_print,ARGS(u.ToEID()))
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC(",")
        /*g_try(v2:"Result",loop:true) */
        Result = Core.F_CALL(C_print,ARGS(v.ToEID()))
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC(") by an update ")
        /*g_try(v2:"Result",loop:true) */
        Result = Core.F_print_any(R.Id())
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC("(")
        /*g_try(v2:"Result",loop:true) */
        Result = Core.F_CALL(C_print,ARGS(x.ToEID()))
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC(") ")
        F_princ_string(ToString(IfThenElse((R.Multivalued_ask == CTRUE),
          MakeString(":add").Id(),
          MakeString(":=").Id())))
        PRINC(" ")
        /*g_try(v2:"Result",loop:true) */
        Result = Core.F_CALL(C_print,ARGS(y.ToEID()))
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC(" \n")
        Result = EVOID
        }}}}}
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        if (p != CNULL) { 
          Result = ToPort(p).UseAsOutput().ToEID()
          } else {
          Result = EID{CFALSE.Id(),0}
          } 
        }
        } 
      } else {
      Result = EID{CFALSE.Id(),0}
      } 
    return Result} 
  
// The EID go function for: trace_rule @ relation (throw: true) 
func E_trace_rule_relation (R EID,s EID,x EID,y EID,u EID,v EID) EID { 
    return F_trace_rule_relation(ToRelation(OBJ(R)),
      ToString(OBJ(s)),
      ANY(x),
      ANY(y),
      ANY(u),
      ANY(v) )} 
  
// stores a set of stopping values
// this is a cool feature : stop(p, list(a1,a2)) => p(x,y) will stop if x = a1 and y = a2
/* {1} The go function for: stop(p:property,l:listargs) [status=1] */
func F_stop_property (p *ClaireProperty ,l *ClaireList ) EID { 
    var Result EID 
    if (Core.F_get_table(Core.C_Core_StopProperty,p.Id()) == CNULL) { 
      Core.F_put_table(Core.C_Core_StopProperty,p.Id(),MakeConstantList(l.Id()).Id())
      }  else if (l.Id() == CNIL.Id()) { 
      Core.F_put_table(Core.C_Core_StopProperty,p.Id(),CNULL)
      } else {
      Core.F_put_table(Core.C_Core_StopProperty,p.Id(),ToList(Core.F_get_table(Core.C_Core_StopProperty,p.Id())).AddFast(MakeConstantList(l.Id()).Id()).Id()/*t=any,s=any*/)
      } 
    Result = EID{CTRUE.Id(),0}
    return Result} 
  
// The EID go function for: stop @ property (throw: true) 
func E_stop_property (p EID,l EID) EID { 
    return F_stop_property(ToProperty(OBJ(p)),ToList(OBJ(l)) )} 
  
// ******************************************************************
// *    Part 3: The debugger interface                              *
// ******************************************************************
// toggle the debug mode
/* {1} The go function for: debug(_CL_obj:void) [status=0] */
func F_debug_void ()  { 
    // procedure body with s = void 
if (ClEnv.Debug_I != -1) { 
      ClEnv.Debug_I = -1
      /*integer->integer*/PRINC("debugger removed\n")
      } else {
      ClEnv.Debug_I = 0
      /*integer->integer*/ClEnv.Ctrace = ToPort(C_stdout.Value)
      /*port->port*/PRINC("debugger installed\n")
      } 
    } 
  
// The EID go function for: debug @ void (throw: false) 
func E_debug_void (_CL_obj EID) EID { 
    F_debug_void( )
    return EVOID} 
  
// this method is called when an error has occured. The value of index
// is recalled with last_index, so that the actual content of the stack is
// preserved.
/* {1} The go function for: call_debug(_CL_obj:void) [status=0] */
func F_call_debug_void () *ClaireAny  { 
    // procedure body with s = any 
var Result *ClaireAny  
    { var top int  = ClEnv.LastDebug
      _ = top
      C_reader.DebugLoop()
      ClEnv.Spy_I = ToObject(CNULL)
      /*object->object*/ClEnv.Trace_I = 0
      /*integer->integer*/ClEnv.Base= ClEnv.LastIndex
      ClEnv.Index= (ClEnv.LastIndex+1)
      ClEnv.Debug_I = top
      /*integer->integer*/if (ClEnv.Verbose > -1) { 
        F_print_exception_void()
        } 
      C_reader.Fromp = ToPort(C_stdin.Value)
      /*port->port*/C_reader.Index = 0
      /*integer->integer*/{ var c *ClaireAny   = Language.C_iClaire_LastCall.Value
        if (c != CNULL) { 
          { 
            var Result_H EID 
            h_index := ClEnv.Index
            h_base := ClEnv.Base
            if (ToInteger(Core.F_get_table(C_Reader_DBline,c)).Value > 0) { 
              PRINC(" \n---- Last call ")
              /*g_try(v2:"Result_H",loop:false) */
              Result_H = Core.F_CALL(C_print,ARGS(c.ToEID()))
              /* ERROR PROTECTION INSERTED (Result_H-Result_H) */
              if !ErrorIn(Result_H) {
              PRINC(" in line ")
              F_princ_integer(ToInteger(Core.F_get_table(C_Reader_DBline,c)).Value)
              PRINC("\n")
              Result_H = EVOID
              }
              } else {
              Result_H = EID{CFALSE.Id(),0}
              } 
            if ErrorIn(Result_H){ 
              ClEnv.Index = h_index
              ClEnv.Base = h_base
              
              } 
            } 
          } else {
          
          } 
        } 
      C__starindex_star.Value = MakeInteger(1).Id()
      C__starcurd_star.Value = MakeInteger(ClEnv.Debug_I).Id()
      Result = MakeInteger(ClEnv.Debug_I).Id()
      C__starmaxd_star.Value = Result
      } 
    return Result} 
  
// The EID go function for: call_debug @ void (throw: false) 
func E_call_debug_void (_CL_obj EID) EID { 
    return F_call_debug_void( ).ToEID()} 
  
// this method is called when an error has occured. The value of index
// is recalled with last_index, so that the actual content of the stack is
// preserved.
/* {1} The go function for: breakpoint(_CL_obj:void) [status=1] */
func F_breakpoint_void () EID { 
    var Result EID 
    { var top int  = ClEnv.Debug_I
      { var t int  = ClEnv.Trace_I
        _ = t
        ClEnv.Trace_I = 0
        /*integer->integer*/C__starindex_star.Value = MakeInteger(0).Id()
        C__starcurd_star.Value = MakeInteger(top).Id()
        C__starmaxd_star.Value = MakeInteger(top).Id()
        /*g_try(v2:"Result",loop:true) */
        var g0168I *ClaireBoolean  
        var try_1 EID 
        /*g_try(v2:"try_1",loop:false) */
        { 
          var v_and4 *ClaireBoolean  
          
          var try_2 EID 
          /*g_try(v2:"try_2",loop:false) */
          try_2 = Core.F_CALL(ToProperty(C__sup.Id()),ARGS(ClEnv.EvalStack[top],EID{C__INT,IVAL(0)}))
          /* ERROR PROTECTION INSERTED (v_and4-try_1) */
          if ErrorIn(try_2) {try_1 = try_2
          } else {
          v_and4 = ToBoolean(OBJ(try_2))
          if (v_and4 == CFALSE) {try_1 = EID{CFALSE.Id(),0}
          } else { 
            v_and4 = Core.F__sup_integer(ClEnv.Debug_I,0)
            if (v_and4 == CFALSE) {try_1 = EID{CFALSE.Id(),0}
            } else { 
              try_1 = EID{CTRUE.Id(),0}} 
            } 
          }
          } 
        /* ERROR PROTECTION INSERTED (g0168I-Result) */
        if ErrorIn(try_1) {Result = try_1
        } else {
        g0168I = ToBoolean(OBJ(try_1))
        if (g0168I == CTRUE) { 
          { var j int  = INT(ClEnv.EvalStack[top])
            { var num_args int  = (INT(ClEnv.EvalStack[(j+2)])-1)
              _ = num_args
              { var start int  = INT(ClEnv.EvalStack[(j+3)])
                { var m *ClaireAny   = ANY(ClEnv.EvalStack[(j+1)])
                  _ = m
                  /*g_try(v2:"Result",loop:true) */
                  PRINC("break in ")
                  /*g_try(v2:"Result",loop:true) */
                  Result = Core.F_CALL(C_print,ARGS(m.ToEID()))
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  PRINC("(")
                  /*g_try(v2:"Result",loop:true) */
                  Result = Core.F_CALL(C_print,ARGS(ClEnv.EvalStack[start]))
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  /*g_try(v2:"Result",loop:true) */
                  { var i int  = (start+1)
                    { var g0167 int  = (start+num_args)
                      _ = g0167
                      Result= EID{CFALSE.Id(),0}
                      for (i <= g0167) { 
                        /* While stat, v:"Result" loop:true */
                        var loop_3 EID 
                        _ = loop_3
                        { 
                        /*g_try(v2:"loop_3",loop:tuple("Result", EID)) */
                        PRINC(",")
                        /*g_try(v2:"loop_3",loop:tuple("Result", EID)) */
                        loop_3 = Core.F_CALL(C_print,ARGS(ClEnv.EvalStack[i]))
                        /* ERROR PROTECTION INSERTED (loop_3-loop_3) */
                        if ErrorIn(loop_3) {Result = loop_3
                        break
                        } else {
                        PRINC("")
                        loop_3 = EVOID
                        }
                        /* ERROR PROTECTION INSERTED (loop_3-loop_3) */
                        if ErrorIn(loop_3) {Result = loop_3
                        break
                        } else {
                        i = (i+1)
                        }
                        /* try?:false, v2:"v_while11" loop will be:tuple("Result", EID) */
                        } 
                      }
                      } 
                    } 
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  PRINC(") [q] >")
                  Result = EVOID
                  }}}
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  { var c *ClaireAny  
                    var try_4 EID 
                    /*g_try(v2:"try_4",loop:false) */
                    try_4 = F_read_string(F_CommandLoopVoid())
                    /* ERROR PROTECTION INSERTED (c-Result) */
                    if ErrorIn(try_4) {Result = try_4
                    } else {
                    c = ANY(try_4)
                    Result= EID{CFALSE.Id(),0}
                    for (c != C_q.Id()) { 
                      /* While stat, v:"Result" loop:true */
                      var loop_5 EID 
                      _ = loop_5
                      { 
                      /*g_try(v2:"loop_5",loop:tuple("Result", EID)) */
                      loop_5 = EVAL(c)
                      /* ERROR PROTECTION INSERTED (loop_5-loop_5) */
                      if ErrorIn(loop_5) {Result = loop_5
                      break
                      } else {
                      PRINC("break>")
                      var try_6 EID 
                      /*g_try(v2:"try_6",loop:tuple("Result", EID)) */
                      try_6 = F_read_string(F_CommandLoopVoid())
                      /* ERROR PROTECTION INSERTED (c-loop_5) */
                      if ErrorIn(try_6) {loop_5 = try_6
                      Result = try_6
                      break
                      } else {
                      c = ANY(try_6)
                      loop_5 = c.ToEID()
                      }}
                      /* try?:false, v2:"v_while10" loop will be:tuple("Result", EID) */
                      } 
                    }
                    }
                    } 
                  }
                  } 
                } 
              } 
            } 
          } else {
          Result = EID{CFALSE.Id(),0}
          } 
        }
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        { 
          var va_arg1 *ClaireEnvironment  
          var va_arg2 int 
          va_arg1 = ClEnv
          va_arg2 = t
          va_arg1.Trace_I = va_arg2
          /*integer->integer*/Result = EID{C__INT,IVAL(va_arg2)}
          } 
        }
        } 
      } 
    return Result} 
  
// The EID go function for: breakpoint @ void (throw: true) 
func E_breakpoint_void (_CL_obj EID) EID { 
    return F_breakpoint_void( )} 
  
// the four keyword
/* {1} The go function for: dn(x:integer) [status=0] */
func F_dn_integer (x int)  { 
    // procedure body with s = void 
for ((INT(ClEnv.EvalStack[ToInteger(C__starcurd_star.Value).Value]) > 0) && 
        (x > 0)) { 
      /* While stat, v:"Unused" loop:false */
      C__starcurd_star.Value = ANY(ClEnv.EvalStack[ToInteger(C__starcurd_star.Value).Value])
      C__starindex_star.Value = MakeInteger((ToInteger(C__starindex_star.Value).Value+1)).Id()
      x = (x-1)
      /* try?:false, v2:"v_while2" loop will be:tuple("Unused", void) */
      } 
    } 
  
// The EID go function for: dn @ integer (throw: false) 
func E_dn_integer (x EID) EID { 
    F_dn_integer(INT(x) )
    return EVOID} 
  
/* {1} The go function for: up(x:integer) [status=0] */
func F_up_integer (x int)  { 
    // procedure body with s = void 
if (x > 0) { 
      { var indices *ClaireList   = CNIL
        { var ind int  = ToInteger(C__starmaxd_star.Value).Value
          _ = ind
          for (ind != ToInteger(C__starcurd_star.Value).Value) { 
            /* While stat, v:"Unused" loop:false */
            indices = F_cons_any(MakeInteger(ind).Id(),indices)
            ind = INT(ClEnv.EvalStack[ind])
            /* try?:false, v2:"v_while5" loop will be:tuple("Unused", void) */
            } 
          if (x > indices.Length()) { 
            C__starcurd_star.Value = C__starmaxd_star.Value
            C__starindex_star.Value = MakeInteger(1).Id()
            } else {
            C__starcurd_star.Value = indices.At(x-1)
            C__starindex_star.Value = MakeInteger((ToInteger(C__starindex_star.Value).Value-x)).Id()
            } 
          } 
        } 
      } 
    } 
  
// The EID go function for: up @ integer (throw: false) 
func E_up_integer (x EID) EID { 
    F_up_integer(INT(x) )
    return EVOID} 
  
// top is the top position in this stack (the last entered message)
/* {1} The go function for: where(x:integer) [status=1] */
func F_where_integer (x int) EID { 
    var Result EID 
    { var j int  = ToInteger(C__starcurd_star.Value).Value
      { var stack_level int  = 0
        _ = stack_level
        Result= EID{CFALSE.Id(),0}
        for ((j > 0) && 
            ((x > 0) && 
              (ClEnv.Debug_I > 0))) { 
          /* While stat, v:"Result" loop:true */
          var loop_1 EID 
          _ = loop_1
          { 
          /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
          loop_1 = F_print_debug_info_integer(j,stack_level,ToInteger(C__starindex_star.Value).Value)
          /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
          if ErrorIn(loop_1) {Result = loop_1
          break
          } else {
          stack_level = (stack_level+1)
          x = (x-1)
          j = INT(ClEnv.EvalStack[j])
          }
          /* try?:false, v2:"v_while4" loop will be:tuple("Result", EID) */
          } 
        }
        } 
      } 
    return Result} 
  
// The EID go function for: where @ integer (throw: true) 
func E_where_integer (x EID) EID { 
    return F_where_integer(INT(x) )} 
  
// note for interpretted methods .. they should be pushing their restriction
// on the stack vs. properties
/* {1} The go function for: print_debug_info(iClaire/index:integer,stack_level:integer,cur_index:integer) [status=1] */
func F_print_debug_info_integer (index int,stack_level int,cur_index int) EID { 
    var Result EID 
    { var num_args int  = (INT(ClEnv.EvalStack[(index+2)])-1)
      _ = num_args
      { var start int  = INT(ClEnv.EvalStack[(index+3)])
        { var m *ClaireAny   = ANY(ClEnv.EvalStack[(index+1)])
          _ = m
          PRINC("debug[")
          F_princ_integer((cur_index+stack_level))
          PRINC("]>")
          { var x int  = 1
            _ = x
            { var g0169 int  = stack_level
              _ = g0169
              for (x <= g0169) { 
                /* While stat, v:"Result" loop:true */
                PRINC(">")
                x = (x+1)
                /* try?:false, v2:"v_while7" loop will be:tuple("Result", void) */
                } 
              } 
            } 
          PRINC(" ")
          /*g_try(v2:"Result",loop:true) */
          Result = Core.F_CALL(C_print,ARGS(m.ToEID()))
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC("(")
          /*g_try(v2:"Result",loop:true) */
          Result = Core.F_CALL(C_print,ARGS(ClEnv.EvalStack[start]))
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          /*g_try(v2:"Result",loop:true) */
          { var i int  = (start+1)
            { var g0170 int  = (start+num_args)
              _ = g0170
              Result= EID{CFALSE.Id(),0}
              for (i <= g0170) { 
                /* While stat, v:"Result" loop:true */
                var loop_1 EID 
                _ = loop_1
                { 
                /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
                PRINC(",")
                /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
                loop_1 = Core.F_CALL(C_print,ARGS(ClEnv.EvalStack[i]))
                /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
                if ErrorIn(loop_1) {Result = loop_1
                break
                } else {
                PRINC("")
                loop_1 = EVOID
                }
                /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
                if ErrorIn(loop_1) {Result = loop_1
                break
                } else {
                i = (i+1)
                }
                /* try?:false, v2:"v_while7" loop will be:tuple("Result", EID) */
                } 
              }
              } 
            } 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(")\n")
          Result = EVOID
          }}}
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: print_debug_info @ integer (throw: true) 
func E_print_debug_info_integer (index EID,stack_level EID,cur_index EID) EID { 
    return F_print_debug_info_integer(INT(index),INT(stack_level),INT(cur_index) )} 
  
// debug version of the debugger :-)  => use as Reader/Show(n)
/* {1} The go function for: Show(n:integer) [status=1] */
func F_Show_integer (n int) EID { 
    var Result EID 
    { var i int  = ToInteger(C__starcurd_star.Value).Value
      Result= EID{CFALSE.Id(),0}
      for ((i > 0) && 
          (n > 0)) { 
        /* While stat, v:"Result" loop:true */
        var loop_1 EID 
        _ = loop_1
        /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
        { var num_args int  = (INT(ClEnv.EvalStack[(i+2)])-1)
          { var start int  = INT(ClEnv.EvalStack[(i+3)])
            _ = start
            /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
            PRINC("[")
            F_princ_integer(start)
            PRINC(" - ")
            F_princ_integer(i)
            PRINC("]: p = ")
            /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
            loop_1 = Core.F_CALL(C_print,ARGS(ClEnv.EvalStack[(i+1)]))
            /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
            if ErrorIn(loop_1) {Result = loop_1
            break
            } else {
            PRINC(", narg = ")
            /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
            loop_1 = Core.F_print_any(MakeInteger(num_args).Id())
            /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
            if ErrorIn(loop_1) {Result = loop_1
            break
            } else {
            PRINC(" \n")
            loop_1 = EVOID
            }}
            /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
            if ErrorIn(loop_1) {Result = loop_1
            break
            } else {
            /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
            { var j int  = 0
              { var g0171 int  = num_args
                _ = g0171
                loop_1= EID{CFALSE.Id(),0}
                for (j <= g0171) { 
                  /* While stat, v:"loop_1" loop:tuple("Result", EID) */
                  var loop_2 EID 
                  _ = loop_2
                  { 
                  /*g_try(v2:"loop_2",loop:tuple("loop_1", EID)) */
                  PRINC("  [")
                  F_princ_integer((j+i))
                  PRINC("]:")
                  /*g_try(v2:"loop_2",loop:tuple("loop_1", EID)) */
                  loop_2 = Core.F_CALL(C_print,ARGS(ClEnv.EvalStack[(j+i)]))
                  /* ERROR PROTECTION INSERTED (loop_2-loop_2) */
                  if ErrorIn(loop_2) {loop_1 = loop_2
                  break
                  } else {
                  PRINC(" \n")
                  loop_2 = EVOID
                  }
                  /* ERROR PROTECTION INSERTED (loop_2-loop_2) */
                  if ErrorIn(loop_2) {loop_1 = loop_2
                  break
                  } else {
                  j = (j+1)
                  }
                  /* try?:false, v2:"v_while8" loop will be:tuple("loop_1", EID) */
                  } 
                }
                } 
              } 
            /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
            if ErrorIn(loop_1) {Result = loop_1
            break
            } else {
            n = (n-1)
            i = INT(ClEnv.EvalStack[i])
            loop_1 = EID{C__INT,IVAL(i)}
            }}
            } 
          } 
        /* ERROR PROTECTION INSERTED (loop_1-Result) */
        if ErrorIn(loop_1) {Result = loop_1
        break
        } else {
        /* try?:false, v2:"v_while3" loop will be:tuple("Result", EID) */
        } 
      }
      } 
    return Result} 
  
// The EID go function for: Show @ integer (throw: true) 
func E_Show_integer (n EID) EID { 
    return F_Show_integer(INT(n) )} 
  
// go to next block
// top is the top position in this stack (the last entered message)
//
/* {1} The go function for: block(x:integer) [status=1] */
func F_block_integer (x int) EID { 
    var Result EID 
    { var j int  = ToInteger(C__starcurd_star.Value).Value
      { var stack_level int  = 0
        Result= EID{CFALSE.Id(),0}
        for ((j > 0) && 
            ((x > 0) && 
              (ClEnv.Debug_I > 0))) { 
          /* While stat, v:"Result" loop:true */
          var loop_1 EID 
          _ = loop_1
          { 
          /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
          { var nargs *ClaireAny   = ANY(ClEnv.EvalStack[(j+2)])
            _ = nargs
            { var start int  = INT(ClEnv.EvalStack[(j+3)])
              { var z *ClaireProperty   = ToProperty(OBJ(ClEnv.EvalStack[(j+1)]))
                { var m *ClaireObject   = Core.F_find_which_list(z.Definition,OWNER(ClEnv.EvalStack[start]),start,INT(Core.F_CALL(ToProperty(Core.C__plus.Id()),ARGS(EID{C__INT,IVAL(start)},nargs.ToEID()))))
                  if (C_method.Id() == m.Isa.Id()) { 
                    { var g0172 *ClaireMethod   = ToMethod(m.Id())
                      PRINC("debug[")
                      F_princ_integer((ToInteger(C__starindex_star.Value).Value+stack_level))
                      PRINC("] > ")
                      /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
                      loop_1 = Core.F_print_any(g0172.Id())
                      /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
                      if ErrorIn(loop_1) {Result = loop_1
                      break
                      } else {
                      PRINC("(")
                      /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
                      if ((g0172.Formula.Id() != CNULL) && 
                          (g0172.Formula.Isa.IsIn(C_lambda) == CTRUE)) { 
                        { var n int  = 0
                          _ = n
                          { 
                            var v *ClaireAny  
                            _ = v
                            loop_1= EID{CFALSE.Id(),0}
                            var v_support *ClaireList  
                            v_support = F_closure_build_lambda(g0172.Formula)
                            v_len := v_support.Length()
                            for i_it := 0; i_it < v_len; i_it++ { 
                              v = v_support.At(i_it)
                              var loop_2 EID 
                              _ = loop_2
                              { 
                              /*g_try(v2:"loop_2",loop:tuple("loop_1", EID)) */
                              /*g_try(v2:"loop_2",loop:tuple("loop_1", EID)) */
                              loop_2 = Core.F_CALL(C_print,ARGS(v.ToEID()))
                              /* ERROR PROTECTION INSERTED (loop_2-loop_2) */
                              if ErrorIn(loop_2) {loop_1 = loop_2
                              break
                              } else {
                              PRINC(" = ")
                              /*g_try(v2:"loop_2",loop:tuple("loop_1", EID)) */
                              loop_2 = Core.F_CALL(C_print,ARGS(ClEnv.EvalStack[(start+n)]))
                              /* ERROR PROTECTION INSERTED (loop_2-loop_2) */
                              if ErrorIn(loop_2) {loop_1 = loop_2
                              break
                              } else {
                              PRINC(", ")
                              loop_2 = EVOID
                              }}
                              /* ERROR PROTECTION INSERTED (loop_2-loop_2) */
                              if ErrorIn(loop_2) {loop_1 = loop_2
                              break
                              } else {
                              n = (n+1)
                              }
                              }
                              } 
                            } 
                          } 
                        } else {
                        PRINC("<compiled:")
                        /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
                        loop_1 = Core.F_print_any(g0172.Module_I.Id())
                        /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
                        if ErrorIn(loop_1) {Result = loop_1
                        break
                        } else {
                        PRINC(">")
                        loop_1 = EVOID
                        }
                        } 
                      /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
                      if ErrorIn(loop_1) {Result = loop_1
                      break
                      } else {
                      PRINC(")\n")
                      loop_1 = EVOID
                      }}
                      } 
                    } else {
                    PRINC("debug[")
                    F_princ_integer((ToInteger(C__starindex_star.Value).Value+stack_level))
                    PRINC("] > ")
                    /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
                    loop_1 = Core.F_print_any(z.Id())
                    /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
                    if ErrorIn(loop_1) {Result = loop_1
                    break
                    } else {
                    PRINC(" -> ")
                    /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
                    loop_1 = Core.F_print_any(m.Id())
                    /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
                    if ErrorIn(loop_1) {Result = loop_1
                    break
                    } else {
                    PRINC("\n")
                    loop_1 = EVOID
                    }}
                    } 
                  } 
                } 
              } 
            } 
          /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
          if ErrorIn(loop_1) {Result = loop_1
          break
          } else {
          stack_level = (stack_level+1)
          x = (x-1)
          j = INT(ClEnv.EvalStack[j])
          }
          /* try?:false, v2:"v_while4" loop will be:tuple("Result", EID) */
          } 
        }
        } 
      } 
    return Result} 
  
// The EID go function for: block @ integer (throw: true) 
func E_block_integer (x EID) EID { 
    return F_block_integer(INT(x) )} 
  
// computes the list of variables of a lambda, including everything
//
/* {1} The go function for: closure_build(self:lambda) [status=0] */
func F_closure_build_lambda (self *ClaireLambda ) *ClaireList  { 
    // procedure body with s = list 
var Result *ClaireList  
    { var lvar *ClaireList   = F_make_list_integer(self.Dimension,CEMPTY.Id())
      { 
        var x *ClaireAny  
        _ = x
        var x_support *ClaireList  
        x_support = self.Vars
        x_len := x_support.Length()
        for i_it := 0; i_it < x_len; i_it++ { 
          x = x_support.At(i_it)
          ToArray(lvar.Id()).NthPut((INT(Core.F_CALL(C_mClaire_index,ARGS(x.ToEID())))+1),x)
          } 
        } 
      F_closure_build_any(self.Body,lvar)
      Result = lvar
      } 
    return Result} 
  
// The EID go function for: closure_build @ lambda (throw: false) 
func E_closure_build_lambda (self EID) EID { 
    return EID{F_closure_build_lambda(ToLambda(OBJ(self)) ).Id(),0}} 
  
// give to each lexical variable its right position in the stack
// answer with the number of lexical variable
//
/* {1} The go function for: closure_build(self:any,lvar:list) [status=0] */
func F_closure_build_any (self *ClaireAny ,lvar *ClaireList )  { 
    // procedure body with s = void 
if (self.Isa.IsIn(C_Variable) == CTRUE) { 
      { var g0174 *ClaireVariable   = To_Variable(self)
        ToArray(lvar.Id()).NthPut((g0174.Index+1),g0174.Id())
        } 
      }  else if (self.Isa.IsIn(C_Instruction) == CTRUE) { 
      { var g0175 *ClaireInstruction   = To_Instruction(self)
        { 
          var s *ClaireSlot  
          _ = s
          var s_iter *ClaireAny  
          for _,s_iter = range(g0175.Isa.Slots.ValuesO()){ 
            s = ToSlot(s_iter)
            F_closure_build_any(Core.F_get_slot(s,ToObject(g0175.Id())),lvar)
            } 
          } 
        } 
      }  else if (self.Isa.IsIn(C_list) == CTRUE) { 
      { var g0176 *ClaireList   = ToList(self)
        _ = g0176
        { 
          var x *ClaireAny  
          _ = x
          var x_support *ClaireList  
          x_support = g0176
          x_len := x_support.Length()
          for i_it := 0; i_it < x_len; i_it++ { 
            x = x_support.At(i_it)
            F_closure_build_any(x,lvar)
            } 
          } 
        } 
      } else {
      
      } 
    } 
  
// The EID go function for: closure_build @ any (throw: false) 
func E_closure_build_any (self EID,lvar EID) EID { 
    F_closure_build_any(ANY(self),ToList(OBJ(lvar)) )
    return EVOID} 
  
// ******************************************************************
// *    Part 5:  Measure &  Profile                                 *
// ******************************************************************
// New in CLAIRE 3.4 - measure objects can be stored on a file and loaded later on
// a measure is a float value counter that stores the sum & sum of squares, to 
// number of experiments
// simple methods add, mean, stdev
/* {1} The go function for: close(x:measure) [status=0] */
func (x *Measure ) Close () *Measure  { 
    x.MIndex = C_measure.Instances.Length()
    /*integer->integer*/return  x
    } 
  
// The EID go function for: close @ measure (throw: false) 
func E_close_measure (x EID) EID { 
    return EID{ToMeasure(OBJ(x)).Close( ).Id(),0}} 
  
/* {1} The go function for: add(x:measure,f:float) [status=0] */
func (x *Measure ) Add (f float64) *Measure  { 
    x.NumValue = (x.NumValue+1)
    /*float->float*/x.SumValue = (x.SumValue+f)
    /*float->float*/x.SumSquare = (x.SumSquare+(f*f))
    /*float->float*/return  x
    } 
  
// The EID go function for: add @ measure (throw: false) 
func E_add_measure (x EID,f EID) EID { 
    return EID{ToMeasure(OBJ(x)).Add(FLOAT(f) ).Id(),0}} 
  
/* {1} The go function for: mean(x:measure) [status=0] */
func (x *Measure ) Mean () float64 { 
    if (x.NumValue == 0) { 
      return  0
      } else {
      return  (x.SumValue/x.NumValue)
      } 
    } 
  
// The EID go function for: mean @ measure (throw: false) 
func E_mean_measure (x EID) EID { 
    return EID{C__FLOAT,FVAL(ToMeasure(OBJ(x)).Mean( ))}} 
  
/* {1} The go function for: stdev(x:measure) [status=0] */
func (x *Measure ) Stdev () float64 { 
    // procedure body with s = float 
var Result float64 
    { var y float64  = ((x.SumSquare/x.NumValue)-F__exp_float((x.SumValue/x.NumValue),2))
      if (y > 0) { 
        Result = F_sqrt_float(y)
        } else {
        Result = 0
        } 
      } 
    return Result} 
  
// The EID go function for: stdev @ measure (throw: false) 
func E_stdev_measure (x EID) EID { 
    return EID{C__FLOAT,FVAL(ToMeasure(OBJ(x)).Stdev( ))}} 
  
/* {1} The go function for: stdev%(x:measure) [status=0] */
func (x *Measure ) Stdev_Z () float64 { 
    return  (x.Stdev()/x.Mean())
    } 
  
// The EID go function for: stdev% @ measure (throw: false) 
func E_stdev_Z_measure (x EID) EID { 
    return EID{C__FLOAT,FVAL(ToMeasure(OBJ(x)).Stdev_Z( ))}} 
  
/* {1} The go function for: reset(x:measure) [status=0] */
func (x *Measure ) Reset ()  { 
    // procedure body with s = void 
x.SumSquare = 0
    /*float->float*/x.NumValue = 0
    /*float->float*/x.SumValue = 0
    /*float->float*/} 
  
// The EID go function for: reset @ measure (throw: false) 
func E_reset_measure (x EID) EID { 
    ToMeasure(OBJ(x)).Reset( )
    return EVOID} 
  
/* {1} The go function for: self_print(m:measure) [status=1] */
func (m *Measure ) SelfPrint () EID { 
    var Result EID 
    /*g_try(v2:"Result",loop:true) */
    Result = Core.F_printFDigit_float(m.Mean(),2)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("[")
    /*g_try(v2:"Result",loop:true) */
    Result = Core.F_printFDigit_float(m.NumValue,0)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("]")
    Result = EVOID
    }}
    return Result} 
  
// The EID go function for: self_print @ measure (throw: true) 
func E_self_print_measure_Reader (m EID) EID { 
    return ToMeasure(OBJ(m)).SelfPrint( )} 
  
// two simple methods to store and retreive measures
//   logMeasure(s:string)  : creates a file
//   load(s:string)        : loads the files, that containts addLog(i,s,ss,n) line
/* {1} The go function for: logMeasure(s:string) [status=1] */
func F_logMeasure_string (s *ClaireString ) EID { 
    var Result EID 
    { var p *ClairePort  
      var try_1 EID 
      /*g_try(v2:"try_1",loop:false) */
      try_1 = F_fopen_string(s,MakeString("w"))
      /* ERROR PROTECTION INSERTED (p-Result) */
      if ErrorIn(try_1) {Result = try_1
      } else {
      p = ToPort(OBJ(try_1))
      { var n int  = Core.F_size_class(C_measure)
        _ = n
        p.UseAsOutput()
        PRINC("// log file produced on ")
        F_princ_string(F_date_I_integer(1))
        PRINC("")
        { 
          var g0181 *ClaireClass  
          _ = g0181
          var g0181_iter *ClaireAny  
          var g0181_support *ClaireSet  
          g0181_support = C_measure.Descendents
          for i_it := 0; i_it < g0181_support.Count; i_it++ { 
            g0181_iter = g0181_support.At(i_it)
            g0181 = ToClass(g0181_iter)
            { var g0182 *ClaireBoolean  
              { 
                var m *Measure  
                _ = m
                var m_iter *ClaireAny  
                g0182= CFALSE
                var m_support *ClaireList  
                m_support = g0181.Instances
                m_len := m_support.Length()
                for i_it := 0; i_it < m_len; i_it++ { 
                  m_iter = m_support.At(i_it)
                  m = ToMeasure(m_iter)
                  PRINC("(addLog(")
                  F_princ_integer(m.MIndex)
                  PRINC(",")
                  F_princ_float(m.SumValue)
                  PRINC(",")
                  F_princ_float(m.SumSquare)
                  PRINC(",")
                  F_princ_float(m.NumValue)
                  PRINC(",")
                  F_princ_integer(n)
                  PRINC("))\n")
                  } 
                } 
              if (g0182 == CTRUE) { 
                
                break
                } 
              } 
            } 
          } 
        p.Fclose()
        Result = EVOID
        } 
      }
      } 
    return Result} 
  
// The EID go function for: logMeasure @ string (throw: true) 
func E_logMeasure_string (s EID) EID { 
    return F_logMeasure_string(ToString(OBJ(s)) )} 
  
// adds a set of measures to a measure object (represented by its index)
/* {1} The go function for: addLog(i:integer,x:float,y:float,n:float,s:integer) [status=1] */
func F_addLog_integer (i int,x float64,y float64,n float64,s int) EID { 
    var Result EID 
    if (Core.F_size_class(C_measure) == s) { 
      { var m *Measure   = ToMeasure(C_measure.Instances.At(i-1))
        m.SumValue = (m.SumValue+x)
        /*float->float*/m.SumSquare = (m.SumSquare+y)
        /*float->float*/{ 
          var va_arg1 *Measure  
          var va_arg2 float64 
          va_arg1 = m
          va_arg2 = (m.NumValue+n)
          va_arg1.NumValue = va_arg2
          /*float->float*/Result = EID{C__FLOAT,FVAL(va_arg2)}
          } 
        } 
      } else {
      Result = ToException(Core.C_general_error.Make(MakeString("logMeasure not compatible with current set (~A vs ~A)").Id(),MakeConstantList(MakeInteger(Core.F_size_class(C_measure)).Id(),MakeInteger(s).Id()).Id())).Close()
      } 
    return Result} 
  
// The EID go function for: addLog @ integer (throw: true) 
func E_addLog_integer (i EID,x EID,y EID,n EID,s EID) EID { 
    return F_addLog_integer(INT(i),
      FLOAT(x),
      FLOAT(y),
      FLOAT(n),
      INT(s) )} 
  
//
// we use a counter object for the 5 interesting values  and
// we use the reified slot to store the counter (thus no profiling on reified)
// start time (1st entry)
// get & create if needed a PRcounter
/* {1} The go function for: PRget(p:property) [status=1] */
func F_PRget_property (p *ClaireProperty ) EID { 
    var Result EID 
    { var x *ClaireBoolean   = p.Reified
      if (x.Id().Isa.Id() == C_PRcount.Id()) { 
        Result = EID{x.Id(),0}
        }  else if (x.Id() == CTRUE.Id()) { 
        Result = ToException(Core.C_general_error.Make(MakeString("[131] Cannot profile a reified property ~S").Id(),MakeConstantList(p.Id()).Id())).Close()
        } else {
        { var _CL_obj *PRcount   = To_PRcount(new(PRcount).Is(C_PRcount))
          _ = _CL_obj
          x = ToBoolean(_CL_obj.Id())
          } 
        p.Reified = x
        /*boolean->boolean*/Result = EID{x.Id(),0}
        } 
      } 
    return Result} 
  
// The EID go function for: PRget @ property (throw: true) 
func E_PRget_property (p EID) EID { 
    return F_PRget_property(ToProperty(OBJ(p)) )} 
  
// get & create if needed a PRcounter
/* {1} The go function for: PRlook(p:property) [status=1] */
func F_PRlook_property2 (p *ClaireProperty ) EID { 
    var Result EID 
    { var arg_1 *PRcount  
      _ = arg_1
      var try_2 EID 
      /*g_try(v2:"try_2",loop:false) */
      try_2 = F_PRget_property(p)
      /* ERROR PROTECTION INSERTED (arg_1-Result) */
      if ErrorIn(try_2) {Result = try_2
      } else {
      arg_1 = To_PRcount(OBJ(try_2))
      Result = F_show_any(arg_1.Id())
      }
      } 
    return Result} 
  
// The EID go function for: PRlook @ property (throw: true) 
func E_PRlook_property2 (p EID) EID { 
    return F_PRlook_property2(ToProperty(OBJ(p)) )} 
  
// show the profiler statistics on one property
/* {1} The go function for: PRshow(p:property) [status=1] */
func F_PRshow_property (p *ClaireProperty ) EID { 
    var Result EID 
    { var x *ClaireBoolean   = p.Reified
      if (x.Isa.IsIn(C_PRcount) == CTRUE) { 
        { var g0186 *PRcount   = To_PRcount(x.Id())
          /*g_try(v2:"Result",loop:true) */
          Result = Core.F_print_any(p.Id())
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(": ")
          F_princ_integer(g0186.Rnum)
          PRINC(" calls -> ")
          F_princ_integer(g0186.Rtime)
          PRINC(" clock tics\n")
          Result = EVOID
          }
          } 
        } else {
        Result = EID{CFALSE.Id(),0}
        } 
      } 
    return Result} 
  
// The EID go function for: PRshow @ property (throw: true) 
func E_PRshow_property (p EID) EID { 
    return F_PRshow_property(ToProperty(OBJ(p)) )} 
  
// elapsed time
/* {1} The go function for: PRtime(p:property) [status=0] */
func F_PRtime_property (p *ClaireProperty ) int { 
    // procedure body with s = integer 
var Result int 
    { var x *ClaireBoolean   = p.Reified
      if (x.Isa.IsIn(C_PRcount) == CTRUE) { 
        { var g0187 *PRcount   = To_PRcount(x.Id())
          _ = g0187
          Result = g0187.Rtime
          } 
        } else {
        Result = 0
        } 
      } 
    return Result} 
  
// The EID go function for: PRtime @ property (throw: false) 
func E_PRtime_property (p EID) EID { 
    return EID{C__INT,IVAL(F_PRtime_property(ToProperty(OBJ(p)) ))}} 
  
/* {1} The go function for: PRcounter(p:property) [status=0] */
func F_PRcounter_property (p *ClaireProperty ) int { 
    // procedure body with s = integer 
var Result int 
    { var x *ClaireBoolean   = p.Reified
      if (x.Isa.IsIn(C_PRcount) == CTRUE) { 
        { var g0189 *PRcount   = To_PRcount(x.Id())
          _ = g0189
          Result = g0189.Rnum
          } 
        } else {
        Result = 0
        } 
      } 
    return Result} 
  
// The EID go function for: PRcounter @ property (throw: false) 
func E_PRcounter_property (p EID) EID { 
    return EID{C__INT,IVAL(F_PRcounter_property(ToProperty(OBJ(p)) ))}} 
  
// show the profiler statistics on the 10 most important properties
/* {1} The go function for: PRshow(_CL_obj:void) [status=1] */
func F_PRshow_void () EID { 
    var Result EID 
    { var l *ClaireList   = ToType(C_property.Id()).EmptyList()
      /*g_try(v2:"Result",loop:true) */
      { 
        var g0191 *ClaireClass  
        _ = g0191
        var g0191_iter *ClaireAny  
        Result= EID{CFALSE.Id(),0}
        var g0191_support *ClaireSet  
        g0191_support = C_property.Descendents
        for i_it := 0; i_it < g0191_support.Count; i_it++ { 
          g0191_iter = g0191_support.At(i_it)
          g0191 = ToClass(g0191_iter)
          var loop_1 EID 
          _ = loop_1
          /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
          { var g0192 *ClaireBoolean  
            var try_2 EID 
            /*g_try(v2:"try_2",loop:false) */
            { 
              var p *ClaireProperty  
              _ = p
              var p_iter *ClaireAny  
              try_2= EID{CFALSE.Id(),0}
              var p_support *ClaireList  
              p_support = g0191.Instances
              p_len := p_support.Length()
              for i_it := 0; i_it < p_len; i_it++ { 
                p_iter = p_support.At(i_it)
                p = ToProperty(p_iter)
                var loop_3 EID 
                _ = loop_3
                /*g_try(v2:"loop_3",loop:tuple("try_2", EID)) */
                var g0194I *ClaireBoolean  
                var try_4 EID 
                /*g_try(v2:"try_4",loop:false) */
                { var arg_5 *ClaireAny  
                  _ = arg_5
                  var try_6 EID 
                  /*g_try(v2:"try_6",loop:false) */
                  { var i int  = 1
                    { var g0193 int  = F_min_integer(10,l.Length())
                      _ = g0193
                      try_6= EID{CFALSE.Id(),0}
                      for (i <= g0193) { 
                        /* While stat, v:"try_6" loop:false */
                        var loop_7 EID 
                        _ = loop_7
                        { 
                        /*g_try(v2:"loop_7",loop:tuple("try_6", EID)) */
                        var g0195I *ClaireBoolean  
                        var try_8 EID 
                        /*g_try(v2:"try_8",loop:false) */
                        if ((F_PRtime_property(p) > F_PRtime_property(ToProperty(l.ValuesO()[i-1]))) || 
                            ((F_PRtime_property(p) == F_PRtime_property(ToProperty(l.ValuesO()[i-1]))) && 
                                (F_PRcounter_property(p) > F_PRcounter_property(ToProperty(l.ValuesO()[i-1]))))) { 
                          var try_9 EID 
                          /*g_try(v2:"try_9",loop:false) */
                          try_9 = l.Nth_plus(i,p.Id())
                          /* ERROR PROTECTION INSERTED (l-try_8) */
                          if ErrorIn(try_9) {try_8 = try_9
                          } else {
                          l = ToList(OBJ(try_9))
                          try_8 = EID{l.Id(),0}
                          try_8 = EID{CTRUE.Id(),0}
                          }
                          } else {
                          try_8 = EID{CFALSE.Id(),0}
                          } 
                        /* ERROR PROTECTION INSERTED (g0195I-loop_7) */
                        if ErrorIn(try_8) {loop_7 = try_8
                        } else {
                        g0195I = ToBoolean(OBJ(try_8))
                        if (g0195I == CTRUE) { 
                          try_6 = EID{CTRUE.Id(),0}
                          break
                          } else {
                          loop_7 = EID{CFALSE.Id(),0}
                          } 
                        }
                        /* ERROR PROTECTION INSERTED (loop_7-loop_7) */
                        if ErrorIn(loop_7) {try_6 = loop_7
                        break
                        } else {
                        i = (i+1)
                        }
                        /* try?:false, v2:"v_while11" loop will be:tuple("try_6", EID) */
                        } 
                      }
                      } 
                    } 
                  /* ERROR PROTECTION INSERTED (arg_5-try_4) */
                  if ErrorIn(try_6) {try_4 = try_6
                  } else {
                  arg_5 = ANY(try_6)
                  try_4 = EID{F_boolean_I_any(arg_5).Id(),0}
                  }
                  } 
                /* ERROR PROTECTION INSERTED (g0194I-loop_3) */
                if ErrorIn(try_4) {loop_3 = try_4
                } else {
                g0194I = ToBoolean(OBJ(try_4))
                if (g0194I == CTRUE) { 
                  loop_3 = EID{CNIL.Id(),0}
                  }  else if (l.Length() < 10) { 
                  l = l.AddFast(p.Id())/*t=property,s=list*/
                  loop_3 = EID{l.Id(),0}
                  } else {
                  loop_3 = EID{CFALSE.Id(),0}
                  } 
                }
                /* ERROR PROTECTION INSERTED (loop_3-try_2) */
                if ErrorIn(loop_3) {try_2 = loop_3
                break
                } else {
                }
                } 
              } 
            /* ERROR PROTECTION INSERTED (g0192-loop_1) */
            if ErrorIn(try_2) {loop_1 = try_2
            } else {
            g0192 = ToBoolean(OBJ(try_2))
            if (g0192 == CTRUE) { 
              Result = EID{g0192.Id(),0}
              break
              } else {
              loop_1 = EID{CFALSE.Id(),0}
              } 
            }
            } 
          /* ERROR PROTECTION INSERTED (loop_1-Result) */
          if ErrorIn(loop_1) {Result = loop_1
          break
          } else {
          }
          } 
        } 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      l.Shrink(10)
      { 
        var p *ClaireProperty  
        _ = p
        var p_iter *ClaireAny  
        Result= EID{CFALSE.Id(),0}
        for _,p_iter = range(l.ValuesO()){ 
          p = ToProperty(p_iter)
          var loop_10 EID 
          _ = loop_10
          /*g_try(v2:"loop_10",loop:tuple("Result", EID)) */
          if (F_PRcounter_property(p) > 0) { 
            PRINC("-----------------------------------\n")
            /*g_try(v2:"loop_10",loop:tuple("Result", EID)) */
            loop_10 = F_PRshow_property(p)
            /* ERROR PROTECTION INSERTED (loop_10-loop_10) */
            if ErrorIn(loop_10) {Result = loop_10
            break
            } else {
            /*g_try(v2:"loop_10",loop:tuple("Result", EID)) */
            { 
              var p2 *ClaireProperty  
              _ = p2
              var p2_iter *ClaireAny  
              loop_10= EID{CFALSE.Id(),0}
              var p2_support *ClaireSet  
              p2_support = ToSet(Core.F_get_table(C_Reader_PRdependent,p.Id()))
              for i_it := 0; i_it < p2_support.Count; i_it++ { 
                p2_iter = p2_support.At(i_it)
                p2 = ToProperty(p2_iter)
                var loop_11 EID 
                _ = loop_11
                /*g_try(v2:"loop_11",loop:tuple("loop_10", EID)) */
                if (F_PRtime_property(p2) > 0) { 
                  PRINC("   * ")
                  /*g_try(v2:"loop_11",loop:tuple("loop_10", EID)) */
                  loop_11 = F_PRshow_property(p2)
                  /* ERROR PROTECTION INSERTED (loop_11-loop_11) */
                  if ErrorIn(loop_11) {loop_10 = loop_11
                  break
                  } else {
                  PRINC("")
                  loop_11 = EVOID
                  }
                  } else {
                  loop_11 = EID{CFALSE.Id(),0}
                  } 
                /* ERROR PROTECTION INSERTED (loop_11-loop_10) */
                if ErrorIn(loop_11) {loop_10 = loop_11
                break
                } else {
                }
                } 
              } 
            /* ERROR PROTECTION INSERTED (loop_10-loop_10) */
            if ErrorIn(loop_10) {Result = loop_10
            break
            } else {
            }}
            } else {
            loop_10 = EID{CFALSE.Id(),0}
            } 
          /* ERROR PROTECTION INSERTED (loop_10-Result) */
          if ErrorIn(loop_10) {Result = loop_10
          break
          } else {
          }
          } 
        } 
      }
      } 
    return Result} 
  
// The EID go function for: PRshow @ void (throw: true) 
func E_PRshow_void (_CL_obj EID) EID { 
    return F_PRshow_void( )} 
  
// reuse from lexical_build in pretty.cl
// returns the list of properties that are used by a method
/* {1} The go function for: dependents(self:method) [status=1] */
func F_dependents_method (self *ClaireMethod ) EID { 
    var Result EID 
    { var p_out *ClaireSet   = ToType(C_property.Id()).EmptySet()
      /*g_try(v2:"Result",loop:true) */
      { 
        var p *ClaireAny  
        _ = p
        Result= EID{CFALSE.Id(),0}
        var p_support *ClaireList  
        var try_1 EID 
        /*g_try(v2:"try_1",loop:false) */
        { var arg_2 *ClaireAny  
          _ = arg_2
          var try_3 EID 
          /*g_try(v2:"try_3",loop:false) */
          try_3 = Core.F_CALL(C_Reader_dependents,ARGS(self.Formula.Body.ToEID()))
          /* ERROR PROTECTION INSERTED (arg_2-try_1) */
          if ErrorIn(try_3) {try_1 = try_3
          } else {
          arg_2 = ANY(try_3)
          try_1 = Core.F_enumerate_any(arg_2)
          }
          } 
        /* ERROR PROTECTION INSERTED (p_support-Result) */
        if ErrorIn(try_1) {Result = try_1
        } else {
        p_support = ToList(OBJ(try_1))
        p_len := p_support.Length()
        for i_it := 0; i_it < p_len; i_it++ { 
          p = p_support.At(i_it)
          var g0196I *ClaireBoolean  
          { var arg_4 *ClaireAny  
            _ = arg_4
            { 
              var r *ClaireRestriction  
              _ = r
              var r_iter *ClaireAny  
              arg_4= CFALSE.Id()
              for _,r_iter = range(ToProperty(p).Restrictions.ValuesO()){ 
                r = ToRestriction(r_iter)
                if (C_method.Id() == r.Isa.Id()) { 
                  arg_4 = CTRUE.Id()
                  break
                  } 
                } 
              } 
            g0196I = F_boolean_I_any(arg_4)
            } 
          if (g0196I == CTRUE) { 
            p_out.AddFast(p)/*t=property,s=void*/
            } 
          }
          } 
        } 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = EID{p_out.Id(),0}
      }
      } 
    return Result} 
  
// The EID go function for: dependents @ method (throw: true) 
func E_dependents_method (self EID) EID { 
    return F_dependents_method(ToMethod(OBJ(self)) )} 
  
// this is really cute ....   v3.2.58: fix typing
/* {1} The go function for: dependents(self:any) [status=1] */
func F_dependents_any (self *ClaireAny ) EID { 
    var Result EID 
    if (self.Isa.IsIn(Language.C_Call) == CTRUE) { 
      { var g0197 *Language.Call   = Language.To_Call(self)
        { var arg_1 *ClaireAny  
          _ = arg_1
          var try_2 EID 
          /*g_try(v2:"try_2",loop:false) */
          try_2 = F_dependents_any(g0197.Args.Id())
          /* ERROR PROTECTION INSERTED (arg_1-Result) */
          if ErrorIn(try_2) {Result = try_2
          } else {
          arg_1 = ANY(try_2)
          Result = Core.F_CALL(ToProperty(C_add.Id()),ARGS(arg_1.ToEID(),EID{g0197.Selector.Id(),0}))
          }
          } 
        } 
      }  else if (self.Isa.IsIn(C_Instruction) == CTRUE) { 
      { var g0198 *ClaireInstruction   = To_Instruction(self)
        _ = g0198
        { var s *ClaireSet   = ToType(C_property.Id()).EmptySet()
          _ = s
          /*g_try(v2:"Result",loop:true) */
          { 
            var sl *ClaireSlot  
            _ = sl
            var sl_iter *ClaireAny  
            Result= EID{CFALSE.Id(),0}
            for _,sl_iter = range(g0198.Isa.Slots.ValuesO()){ 
              sl = ToSlot(sl_iter)
              var loop_3 EID 
              _ = loop_3
              var try_4 EID 
              /*g_try(v2:"try_4",loop:tuple("Result", EID)) */
              { var arg_5 *ClaireAny  
                _ = arg_5
                var try_6 EID 
                /*g_try(v2:"try_6",loop:false) */
                try_6 = Core.F_CALL(C_Reader_dependents,ARGS(Core.F_get_slot(sl,ToObject(g0198.Id())).ToEID()))
                /* ERROR PROTECTION INSERTED (arg_5-try_4) */
                if ErrorIn(try_6) {try_4 = try_6
                } else {
                arg_5 = ANY(try_6)
                try_4 = EID{Core.F_U_type(ToType(s.Id()),ToType(arg_5)).Id(),0}
                }
                } 
              /* ERROR PROTECTION INSERTED (s-Result) */
              if ErrorIn(try_4) {Result = try_4
              break
              } else {
              s = ToSet(OBJ(try_4))
              loop_3 = EID{s.Id(),0}
              }
              } 
            } 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          Result = EID{s.Id(),0}
          }
          } 
        } 
      }  else if (self.Isa.IsIn(C_list) == CTRUE) { 
      { var g0199 *ClaireList   = ToList(self)
        _ = g0199
        { var s *ClaireSet   = ToType(C_property.Id()).EmptySet()
          _ = s
          /*g_try(v2:"Result",loop:true) */
          { 
            var x *ClaireAny  
            _ = x
            Result= EID{CFALSE.Id(),0}
            var x_support *ClaireList  
            x_support = g0199
            x_len := x_support.Length()
            for i_it := 0; i_it < x_len; i_it++ { 
              x = x_support.At(i_it)
              var loop_7 EID 
              _ = loop_7
              var try_8 EID 
              /*g_try(v2:"try_8",loop:tuple("Result", EID)) */
              { var arg_9 *ClaireAny  
                _ = arg_9
                var try_10 EID 
                /*g_try(v2:"try_10",loop:false) */
                try_10 = Core.F_CALL(C_Reader_dependents,ARGS(x.ToEID()))
                /* ERROR PROTECTION INSERTED (arg_9-try_8) */
                if ErrorIn(try_10) {try_8 = try_10
                } else {
                arg_9 = ANY(try_10)
                try_8 = EID{Core.F_U_type(ToType(s.Id()),ToType(arg_9)).Id(),0}
                }
                } 
              /* ERROR PROTECTION INSERTED (s-Result) */
              if ErrorIn(try_8) {Result = try_8
              break
              } else {
              s = ToSet(OBJ(try_8))
              loop_7 = EID{s.Id(),0}
              }
              } 
            } 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          Result = EID{s.Id(),0}
          }
          } 
        } 
      }  else if (self.Isa.IsIn(C_property) == CTRUE) { 
      { var g0200 *ClaireProperty   = ToProperty(self)
        _ = g0200
        Result = EID{MakeSet(ToType(C_property.Id()),g0200.Id()).Id(),0}
        } 
      } else {
      Result = EID{ToType(C_property.Id()).EmptySet().Id(),0}
      } 
    return Result} 
  
// The EID go function for: dependents @ any (throw: true) 
func E_dependents_any (self EID) EID { 
    return F_dependents_any(ANY(self) )} 
  
// used to set up the dependence
/* {1} The go function for: PRdepends(p:property,p2:property) [status=1] */
func F_PRdepends_property (p *ClaireProperty ,p2 *ClaireProperty ) EID { 
    var Result EID 
    Result = Core.F_add_table(C_Reader_PRdependent,p.Id(),p2.Id())
    return Result} 
  
// The EID go function for: PRdepends @ property (throw: true) 
func E_PRdepends_property (p EID,p2 EID) EID { 
    return F_PRdepends_property(ToProperty(OBJ(p)),ToProperty(OBJ(p2)) )} 
  
// end of file