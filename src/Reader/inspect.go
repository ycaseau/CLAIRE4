/***** CLAIRE Compilation of file /Users/ycaseau/Dropbox/src/clairev4.12/src/meta/inspect.cl 
         [version 4.1.4 / safety 5] Friday 01-03-2025 16:21:02 *****/

package Reader
import (_ "fmt"
	. "Kernel"
	"Core"
	"Language"
)

//-------- dumb function to prevent import errors --------
func import_g0150() { 
  _ = Core.It
  _ = Language.It
  } 


//+-------------------------------------------------------------+
//| CLAIRE                                                      |
//| inspect.cl                                                  |
//| Copyright (C) 1994 - 2025 Yves Caseau. All Rights Reserved  |
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
// v4.10 use directly "reader" object versus r variable, to support reboot
/* The go function for: top_level(r:meta_reader) [status=1] */
func (r *MetaReader) TopLevel () EID { 
  var Result EID
  { var res *ClaireAny = MakeInteger(0).Id()
    Result= EID{CFALSE.Id(),0}
    for (res != C_q.Id()) { 
      var loop_1 EID
      _ = loop_1
      { 
      { var arg_2 *ClaireString
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
      { 
        h_index := ClEnv.Index
        h_base := ClEnv.Base
        C_reader.Toplevel = CTRUE
        if (ClEnv.CountCall > 0) { 
          ClEnv.CountCall = 1
          } 
        var try_3 EID
        { var arg_4 *ClaireAny
          var try_5 EID
          try_5 = C_reader.Nextunit()
          if ErrorIn(try_5) {try_3 = try_5
          } else {
          arg_4 = ANY(try_5)
          try_3 = Language.F_iClaire_lexical_index_any2(arg_4,CNIL,0,CTRUE)
          }
          } 
        if ErrorIn(try_3) {loop_1 = try_3
        } else {
        res = ANY(try_3)
        loop_1 = res.ToEID()
        if (ToInteger(C_Reader_TopLevelMode.Value).Value == 1) { 
          ClEnv.Index = 20
          } 
        if ((ToInteger(C_Reader_TopLevelMode.Value).Value == 3) && 
            (res != C_q.Id())) { 
          loop_1 = F_inspect_loop_any(res,ToList(C_Reader_InspectStack.Value))
          } else {
          if (ToInteger(C_Reader_TopLevelMode.Value).Value == 1) { 
            PRINC("eval[")
            { var arg_6 *ClaireAny
              arg_6 = MakeInteger((ToInteger(C_Reader_TopCount.Value).Value+1)).Id()
              C_Reader_TopCount.Value = arg_6
              loop_1 = Core.F_print_any(arg_6)
              } 
            if !ErrorIn(loop_1) {
            PRINC("]> ")
            loop_1 = EVOID
            }
            } else {
            PRINC("> ")
            loop_1 = EVOID
            } 
          if !ErrorIn(loop_1) {
          var try_8 EID
          try_8 = EVAL(res)
          if ErrorIn(try_8) {loop_1 = try_8
          } else {
          res = ANY(try_8)
          loop_1 = res.ToEID()
          if (res != C_q.Id()) { 
            loop_1 = Core.F_CALL(C_print,ARGS(res.ToEID()))
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
          var g0151I *ClaireBoolean
          { var e *ClaireException = ClEnv.Exception_I
            g0151I = MakeBoolean((e.Isa.Id() == C_system_error.Id()) && (ToSystemError(e.Id()).Index == -1))
            } 
          if (g0151I == CTRUE) { 
            C_Reader_TopLevelMode.Value = MakeInteger(1).Id()
            res = C_q.Id()
            loop_1 = res.ToEID()
            } else {
            C_reader.RestoreState()
            if (C_reader.External.Value != MakeString("toplevel").Value) { 
              PRINC("---- file: ")
              F_princ_string(C_reader.External)
              PRINC(", line: ")
              F_princ_integer(ClEnv.NLine)
              PRINC("\n")
              } 
            loop_1 = F_debug_if_possible_void()
            if ErrorIn(loop_1) {Result = loop_1
            break
            } else {
            PRINC("\n")
            loop_1 = EVOID
            }
            } 
          } 
        } 
      if ErrorIn(loop_1) {Result = loop_1
      break
      } else {
      if ((ToInteger(C_Reader_TopLevelMode.Value).Value != 1) && 
          (res == C_q.Id())) { 
        if (ToInteger(C_Reader_TopLevelMode.Value).Value == 2) { 
          ClEnv.Index = ToInteger(C_Reader_TopIndex.Value).Value
          ClEnv.Base = ToInteger(C_Reader_TopBase.Value).Value
          ClEnv.Trace_I = 1
          ClEnv.Debug_I = ToInteger(C_Reader_TopDebug.Value).Value
          } 
        res = CNULL
        C_Reader_TopLevelMode.Value = MakeInteger(1).Id()
        } 
      }
      } 
    }
    if !ErrorIn(Result) {
    PRINC("\n")
    Result = EVOID
    }
    } 
  return Result} 

// The EID go function for: top_level @ meta_reader (throw: true) 
func E_Reader_top_level_meta_reader (r EID) EID { 
  return ToMetaReader(OBJ(r)).TopLevel( )} 

//        exit(1)) ]
// start a debug loop - aha 
/* The go function for: debugLoop(r:meta_reader) [status=0] */
func (r *MetaReader) DebugLoop ()  { 
  C_Reader_TopDebug.Value = MakeInteger(0).Id()
  C_Reader_TopBase.Value = MakeInteger(ClEnv.Base).Id()
  C_Reader_TopIndex.Value = MakeInteger(ClEnv.Index).Id()
  r.Toplevel = CTRUE
  C_Reader_TopLevelMode.Value = MakeInteger(2).Id()
  PRINC("--------------- Debug -------------------\n")
  } 

// The EID go function for: debugLoop @ meta_reader (throw: false) 
func E_Reader_debugLoop_meta_reader (r EID) EID { 
  ToMetaReader(OBJ(r)).DebugLoop( )
  return EVOID} 

// starts an inspector  on a list
/* The go function for: inspect_system(l:list) [status=0] */
func F_Reader_inspect_system_list2 (l *ClaireList)  { 
  C_Reader_InspectStack.Value = l.Id()
  if (ToInteger(C_Reader_TopLevelMode.Value).Value == 2) { 
    ClEnv.Trace_I = 1
    } 
  C_Reader_TopLevelMode.Value = MakeInteger(3).Id()
  } 

// The EID go function for: inspect_system @ list (throw: false) 
func E_Reader_inspect_system_list2 (l EID) EID { 
  F_Reader_inspect_system_list2(ToList(OBJ(l)) )
  return EVOID} 

// INSPECT   
// simple main (to be enriched later)
/* The go function for: simple_main(_CL_obj:void) [status=1] */
func F_Reader_simple_main_void () EID { 
  var Result EID
  Result = C_reader.TopLevel()
  if !ErrorIn(Result) {
  PRINC("[regular exit] Bye.\n")
  Result = EVOID
  }
  return Result} 

// The EID go function for: simple_main @ void (throw: true) 
func E_Reader_simple_main_void (_CL_obj EID) EID { 
  return F_Reader_simple_main_void( )} 

// *********************************************************************
// *      Part 2: Inspection                                           *
// *********************************************************************
// this is the method that the user calls
//
/* The go function for: inspect(self:any) [status=1] */
func F_inspect_any (self *ClaireAny) EID { 
  var Result EID
  { var m0 *ClaireModule = ClEnv.Module_I
    { var ix int = 0
      if (self.Isa.IsIn(C_list) == CTRUE) { 
        { var g0152 *ClaireList = ToList(self)
          { var i int = 1
            { var g0153 int = g0152.Length()
              Result= EID{CFALSE.Id(),0}
              for (i <= g0153) { 
                var loop_1 EID
                _ = loop_1
                { 
                F_princ_integer(i)
                PRINC(": ")
                loop_1 = Core.F_CALL(C_print,ARGS(g0152.At(i-1).ToEID()))
                if ErrorIn(loop_1) {Result = loop_1
                break
                } else {
                PRINC("\n")
                loop_1 = EVOID
                }
                if ErrorIn(loop_1) {Result = loop_1
                break
                } else {
                i = (i+1)
                }
                } 
              }
              } 
            } 
          } 
        }  else if (C_set.Id() == self.Isa.Id()) { 
        { var g0154 *ClaireSet = ToSet(self)
          Result = F_inspect_any(g0154.List_I().Id())
          } 
        }  else if (Core.F__Z_any1(self,C_object) == CTRUE) { 
        { var g0155 *ClaireObject = ToObject(self)
          { 
            var rel *ClaireSlot
            _ = rel
            var rel_iter *ClaireAny
            Result= EID{CFALSE.Id(),0}
            var rel_support *ClaireList
            rel_support = g0155.Id().Isa.Slots
            for _,rel_iter = range(rel_support.ValuesO()){ 
              rel = ToSlot(rel_iter)
              var loop_2 EID
              _ = loop_2
              { var m *ClaireModule = rel.Selector.Name.Module_I()
                ix = (ix+1)
                if ((m.Id() == m0.Id()) || 
                    ((m.Id() == C_claire.Id()) || 
                      (ToBoolean(C__starshowall_star.Value) == CTRUE))) { 
                  { var val *ClaireAny = Core.F_get_slot(rel,g0155)
                    F_princ_integer(ix)
                    PRINC(": ")
                    loop_2 = Core.F_print_any(rel.Selector.Id())
                    if ErrorIn(loop_2) {Result = loop_2
                    break
                    } else {
                    PRINC(" = ")
                    if (val.Isa.IsIn(C_list) == CTRUE) { 
                      { var g0156 *ClaireList = ToList(val)
                        if (g0156.Length() < 10) { 
                          loop_2 = Language.F_pretty_print_any(g0156.Id())
                          } else {
                          { var arg_3 *ClaireList
                            { var i_bag *ClaireList = ToType(CEMPTY.Id()).EmptyList()
                              { var i int = 1
                                { var g0157 int = 9
                                  for (i <= g0157) { 
                                    i_bag.AddFast(g0156.At(i-1))
                                    i = (i+1)
                                    } 
                                  } 
                                } 
                              arg_3 = i_bag
                              } 
                            loop_2 = Language.F_pretty_print_any(arg_3.Id())
                            } 
                          if ErrorIn(loop_2) {Result = loop_2
                          break
                          } else {
                          PRINC("...")
                          loop_2 = EVOID
                          }
                          } 
                        } 
                      } else {
                      loop_2 = Language.F_pretty_print_any(val)
                      } 
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
                if ErrorIn(loop_2) {Result = loop_2
                break
                } else {
                }
                } 
              if ErrorIn(loop_2) {Result = loop_2
              break
              } else {
              }
              } 
            } 
          } 
        } else {
        Result = Language.F_pretty_print_any(self)
        if !ErrorIn(Result) {
        PRINC("\n")
        Result = EVOID
        }
        } 
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
/* The go function for: inspect_loop(%read:any,old:list) [status=1] */
func F_inspect_loop_any (_Zread *ClaireAny,old *ClaireList) EID { 
  var Result EID
  { var self *ClaireAny = old.At(0)
    var g0161I *ClaireBoolean
    if (_Zread.Isa.IsIn(Language.C_Call) == CTRUE) { 
      { var g0160 *Language.Call = Language.To_Call(_Zread)
        g0161I = Equal(g0160.Selector.Id(),C_put.Id())
        } 
      } else {
      g0161I = CFALSE
      } 
    if (g0161I == CTRUE) { 
      { var n int = ToInteger(ToList(OBJ(Core.F_CALL(C_args,ARGS(_Zread.ToEID())))).At(0)).Value
        { var s *ClaireSymbol
          var try_1 EID
          try_1 = Language.F_extract_symbol_any(ToList(OBJ(Core.F_CALL(C_args,ARGS(_Zread.ToEID())))).At(1))
          if ErrorIn(try_1) {Result = try_1
          } else {
          s = ToSymbol(OBJ(try_1))
          if (C_integer.Id() != OWNER(EID{C__INT,IVAL(n)}).Id()) { 
            Result = ToException(Core.C_general_error.Make(MakeString("[128] ~S should be an integer").Id(),MakeConstantList(MakeInteger(n).Id()).Id())).Close()
            } else {
            Result = EID{CFALSE.Id(),0}
            } 
          if !ErrorIn(Result) {
          { var val *ClaireAny
            var try_2 EID
            try_2 = F_get_from_integer_any(self,n)
            if ErrorIn(try_2) {Result = try_2
            } else {
            val = ANY(try_2)
            { 
              var va_arg1 *Core.GlobalVariable
              var va_arg2 *ClaireAny
              var try_3 EID
              try_3 = Core.F_new_class2(Core.C_global_variable,s)
              if ErrorIn(try_3) {Result = try_3
              } else {
              va_arg1 = Core.ToGlobalVariable(OBJ(try_3))
              va_arg2 = val
              va_arg1.Value = va_arg2
              Result = va_arg2.ToEID()
              }
              } 
            if !ErrorIn(Result) {
            Result = F_inspect_any(val)
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
        try_4 = old.Cdr()
        if ErrorIn(try_4) {Result = try_4
        } else {
        old = ToList(OBJ(try_4))
        Result = EID{old.Id(),0}
        Result = F_inspect_any(old.At(0))
        }
        } else {
        Result = EID{CFALSE.Id(),0}
        } 
      }  else if (C_integer.Id() == _Zread.Isa.Id()) { 
      { var val *ClaireAny
        var try_5 EID
        try_5 = F_get_from_integer_any(self,ToInteger(_Zread).Value)
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
/* The go function for: get_from_integer(self:any,n:integer) [status=1] */
func F_get_from_integer_any (self *ClaireAny,n int) EID { 
  var Result EID
  if (self.Isa.IsIn(C_list) == CTRUE) { 
    if ((n > 0) && 
        (n <= ToList(self).Length())) { 
      Result = Core.F_CALL(C_nth,ARGS(self.ToEID(),EID{C__INT,IVAL(n)}))
      } else {
      F_princ_integer(n)
      PRINC(" in not a good index for ")
      Result = Core.F_CALL(C_print,ARGS(self.ToEID()))
      if !ErrorIn(Result) {
      PRINC(".\n")
      Result = EVOID
      }
      if !ErrorIn(Result) {
      Result = self.ToEID()
      }
      } 
    } else {
    { var l *ClaireList = self.Isa.Slots
      if ((n > 0) && 
          (n <= l.Length())) { 
        Result = Core.F_SUPER(C_get, C_slot, ARGS(l.ValuesO()[n-1].ToEID(),self.ToEID()))
        } else {
        F_princ_integer(n)
        PRINC(" is not a good index for ")
        Result = Core.F_CALL(C_print,ARGS(self.ToEID()))
        if !ErrorIn(Result) {
        PRINC(".\n")
        Result = EVOID
        }
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
/* The go function for: iClaire/trace_on(self:any) [status=1] */
func F_trace_on_any (self *ClaireAny) EID { 
  var Result EID
  if (self.Isa.IsIn(C_property) == CTRUE) { 
    { var g0162 *ClaireProperty = ToProperty(self)
      if (g0162.Id() == Core.C_spy.Id()) { 
        { var m *ClaireAny = Core.F__at_property1(Core.C_spy,C_void).Id()
          if (F_boolean_I_any(m) == CTRUE) { 
            Result = F_store_object(ToObject(ClEnv.Id()),
              17,
              C_object,
              m,
              CFALSE).ToEID()
            } else {
            Result = EID{CFALSE.Id(),0}
            } 
          } 
        }  else if (g0162.Id() == C_where.Id()) { 
        Result = Core.F_write_property(C_Kernel_call_count,ToObject(ClEnv.Id()),MakeInteger(1).Id())
        } else {
        { 
          var va_arg1 *ClaireProperty
          var va_arg2 int
          va_arg1 = g0162
          va_arg2 = (5-ClEnv.Verbose)
          va_arg1.Trace_I = va_arg2
          Result = EID{C__INT,IVAL(va_arg2)}
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
      Result = EID{C__INT,IVAL(va_arg2)}
      } 
    }  else if (self.Isa.IsIn(C_module) == CTRUE) { 
    { var g0164 *ClaireModule = ToModule(self)
      if (g0164.Status > 2) { 
        g0164.Status = 4
        } 
      { 
        var m *ClaireAny
        _ = m
        Result= EID{CFALSE.Id(),0}
        var m_support *ClaireList
        m_support = g0164.Parts
        m_len := m_support.Length()
        for i_it := 0; i_it < m_len; i_it++ { 
          m = m_support.At(i_it)
          var loop_1 EID
          _ = loop_1
          loop_1 = F_trace_on_any(m)
          if ErrorIn(loop_1) {Result = loop_1
          break
          } else {
          }
          } 
        } 
      } 
    }  else if (C_port.Id() == self.Isa.Id()) { 
    { var g0165 *ClairePort = ToPort(self)
      { 
        var va_arg1 *ClaireEnvironment
        var va_arg2 *ClairePort
        va_arg1 = ClEnv
        va_arg2 = g0165
        va_arg1.Ctrace = va_arg2
        Result = va_arg2.ToEID()
        } 
      } 
    }  else if (C_string.Id() == self.Isa.Id()) { 
    { var g0166 *ClaireString = ToString(self)
      { 
        var va_arg1 *ClaireEnvironment
        var va_arg2 *ClairePort
        va_arg1 = ClEnv
        var try_2 EID
        try_2 = F_fopen_string(g0166,MakeString("w"))
        if ErrorIn(try_2) {Result = try_2
        } else {
        va_arg2 = ToPort(OBJ(try_2))
        va_arg1.Ctrace = va_arg2
        Result = va_arg2.ToEID()
        }
        } 
      } 
    }  else if (C_integer.Id() == self.Isa.Id()) { 
    { var g0167 int = ToInteger(self).Value
      { 
        var va_arg1 *ClaireEnvironment
        var va_arg2 int
        va_arg1 = ClEnv
        va_arg2 = g0167
        va_arg1.Verbose = va_arg2
        Result = EID{C__INT,IVAL(va_arg2)}
        } 
      } 
    } else {
    Result = ToException(Core.C_general_error.Make(MakeString("[129] trace not implemented on ~S\n").Id(),MakeConstantList(self).Id())).Close()
    } 
  if !ErrorIn(Result) {
  Result = self.ToEID()
  }
  return Result} 

// The EID go function for: iClaire/trace_on @ any (throw: true) 
func E_trace_on_any (self EID) EID { 
  return F_trace_on_any(ANY(self) )} 

/* The go function for: untrace(self:any) [status=1] */
func F_untrace_any (self *ClaireAny) EID { 
  var Result EID
  if (self.Isa.IsIn(C_property) == CTRUE) { 
    { var g0169 *ClaireProperty = ToProperty(self)
      if (g0169.Id() == Core.C_spy.Id()) { 
        { 
          var va_arg1 *ClaireEnvironment
          var va_arg2 *ClaireObject
          va_arg1 = ClEnv
          va_arg2 = ToObject(CNULL)
          va_arg1.Spy_I = va_arg2
          Result = EID{va_arg2.Id(),0}
          } 
        }  else if (g0169.Id() == C_where.Id()) { 
        Result = Core.F_write_property(C_Kernel_call_count,ToObject(ClEnv.Id()),MakeInteger(-1).Id())
        } else {
        { 
          var va_arg1 *ClaireProperty
          var va_arg2 int
          va_arg1 = g0169
          va_arg2 = 0
          va_arg1.Trace_I = va_arg2
          Result = EID{C__INT,IVAL(va_arg2)}
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
      Result = EID{C__INT,IVAL(va_arg2)}
      } 
    }  else if (self.Isa.IsIn(C_module) == CTRUE) { 
    { var g0171 *ClaireModule = ToModule(self)
      if (g0171.Status == 4) { 
        g0171.Status = 3
        } 
      { 
        var m *ClaireAny
        _ = m
        Result= EID{CFALSE.Id(),0}
        var m_support *ClaireList
        m_support = g0171.Parts
        m_len := m_support.Length()
        for i_it := 0; i_it < m_len; i_it++ { 
          m = m_support.At(i_it)
          var loop_1 EID
          _ = loop_1
          loop_1 = F_untrace_any(m)
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
      Result = va_arg2.ToEID()
      } 
    } else {
    Result = ToException(Core.C_general_error.Make(MakeString("[130] untrace not implemented on ~S\n").Id(),MakeConstantList(self).Id())).Close()
    } 
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
/* The go function for: spy(l:listargs) [status=0] */
func F_spy_listargs2_Reader (l *ClaireList)  { 
  { var m *ClaireAny = Core.F__at_property1(Core.C_spy,C_void).Id()
    if (F_boolean_I_any(m) == CTRUE) { 
      F_store_object(ToObject(ClEnv.Id()),
        17,
        C_object,
        m,
        CFALSE)
      { 
        var g0174 *ClaireClass
        _ = g0174
        var g0174_iter *ClaireAny
        var g0174_support *ClaireSet
        g0174_support = C_property.Descendants
        for i_it := 0; i_it < g0174_support.Count; i_it++ { 
          g0174_iter = g0174_support.At(i_it)
          g0174 = ToClass(g0174_iter)
          { 
            var f *ClaireProperty
            _ = f
            var f_iter *ClaireAny
            var f_support *ClaireList
            f_support = g0174.Instances
            f_len := f_support.Length()
            for i_it := 0; i_it < f_len; i_it++ { 
              f_iter = f_support.At(i_it)
              f = ToProperty(f_iter)
              if (ToList(l.Id()).Memq(f.Id()) == CTRUE) { 
                f.IfWrite = CNIL.Id()
                } 
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
/* The go function for: trace_rule(R:relation,s:string,x:any,y:any,u:any,v:any) [status=1] */
func F_trace_rule_relation (R *ClaireRelation,s *ClaireString,x *ClaireAny,y *ClaireAny,u *ClaireAny,v *ClaireAny) EID { 
  var Result EID
  if ((C_if_write.Trace_I+ClEnv.Verbose) >= 5) { 
    { var p *ClaireAny = Core.F_get_property(C_ctrace,ToObject(ClEnv.Id()))
      if (p != CNULL) { 
        p = ToPort(p).UseAsOutput().Id()
        } 
      PRINC("--- the rule ")
      F_princ_string(s)
      PRINC(" is triggered for (")
      Result = Core.F_CALL(C_print,ARGS(u.ToEID()))
      if !ErrorIn(Result) {
      PRINC(",")
      Result = Core.F_CALL(C_print,ARGS(v.ToEID()))
      if !ErrorIn(Result) {
      PRINC(") by an update ")
      Result = Core.F_print_any(R.Id())
      if !ErrorIn(Result) {
      PRINC("(")
      Result = Core.F_CALL(C_print,ARGS(x.ToEID()))
      if !ErrorIn(Result) {
      PRINC(") ")
      F_princ_string(ToString(IfThenElse((R.Multivalued_ask == CTRUE),
        MakeString(":add").Id(),
        MakeString(":=").Id())))
      PRINC(" ")
      Result = Core.F_CALL(C_print,ARGS(y.ToEID()))
      if !ErrorIn(Result) {
      PRINC(" \n")
      Result = EVOID
      }}}}}
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
/* The go function for: stop(p:property,l:listargs) [status=1] */
func F_stop_property (p *ClaireProperty,l *ClaireList) EID { 
  var Result EID
  if (Core.F_get_table(Core.C_Core_StopProperty,p.Id()) == CNULL) { 
    Core.F_put_table(Core.C_Core_StopProperty,p.Id(),MakeConstantList(l.Id()).Id())
    }  else if (l.Id() == CNIL.Id()) { 
    Core.F_put_table(Core.C_Core_StopProperty,p.Id(),CNULL)
    } else {
    Core.F_put_table(Core.C_Core_StopProperty,p.Id(),ToList(Core.F_get_table(Core.C_Core_StopProperty,p.Id())).AddFast(MakeConstantList(l.Id()).Id()).Id())
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
/* The go function for: debug(_CL_obj:void) [status=0] */
func F_debug_void ()  { 
  if (ClEnv.Debug_I != -1) { 
    ClEnv.Debug_I = -1
    PRINC("debugger removed\n")
    } else {
    ClEnv.Debug_I = 0
    ClEnv.Ctrace = ToPort(C_stdout.Value)
    ClEnv.Jito_ask = CFALSE
    PRINC("debugger installed\n")
    } 
  } 

// The EID go function for: debug @ void (throw: false) 
func E_debug_void (_CL_obj EID) EID { 
  F_debug_void( )
  return EVOID} 

// this method is called when an error has occurred. The value of index
// is recalled with last_index, so that the actual content of the stack is
// preserved.
/* The go function for: call_debug(_CL_obj:void) [status=0] */
func F_call_debug_void () *ClaireAny { 
  var Result *ClaireAny
  { var top int = ClEnv.LastDebug
    C_reader.DebugLoop()
    ClEnv.Spy_I = ToObject(CNULL)
    ClEnv.Trace_I = 0
    ClEnv.Base= ClEnv.LastIndex
    ClEnv.Index= (ClEnv.LastIndex+1)
    ClEnv.Debug_I = top
    if (ClEnv.Verbose > -1) { 
      F_print_exception_void()
      } 
    C_reader.Fromp = ToPort(C_stdin.Value)
    C_reader.Index = 0
    { var c *ClaireAny = Language.C_iClaire_LastCall.Value
      if (c != CNULL) { 
        { 
          var Result_H EID
          h_index := ClEnv.Index
          h_base := ClEnv.Base
          if (ToInteger(Core.F_get_table(C_Reader_DBline,c)).Value > 0) { 
            PRINC(" \n---- Last call ")
            Result_H = Core.F_CALL(C_print,ARGS(c.ToEID()))
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

// this method is called when an error has occurred. The value of index
// is recalled with last_index, so that the actual content of the stack is
// preserved.
/* The go function for: breakpoint(_CL_obj:void) [status=1] */
func F_breakpoint_void () EID { 
  var Result EID
  { var top int = ClEnv.Debug_I
    { var t int = ClEnv.Trace_I
      ClEnv.Trace_I = 0
      C__starindex_star.Value = MakeInteger(0).Id()
      C__starcurd_star.Value = MakeInteger(top).Id()
      C__starmaxd_star.Value = MakeInteger(top).Id()
      var g0177I *ClaireBoolean
      var try_1 EID
      { 
        var v_and3 *ClaireBoolean
        
        var try_2 EID
        try_2 = Core.F_CALL(ToProperty(C__sup.Id()),ARGS(ClEnv.EvalStack[top],EID{C__INT,IVAL(0)}))
        if ErrorIn(try_2) {try_1 = try_2
        } else {
        v_and3 = ToBoolean(OBJ(try_2))
        if (v_and3 == CFALSE) {try_1 = EID{CFALSE.Id(),0}
        } else { 
          v_and3 = Core.F__sup_integer(ClEnv.Debug_I,0)
          if (v_and3 == CFALSE) {try_1 = EID{CFALSE.Id(),0}
          } else { 
            try_1 = EID{CTRUE.Id(),0}} 
          } 
        }
        } 
      if ErrorIn(try_1) {Result = try_1
      } else {
      g0177I = ToBoolean(OBJ(try_1))
      if (g0177I == CTRUE) { 
        { var j int = INT(ClEnv.EvalStack[top])
          { var num_args int = (INT(ClEnv.EvalStack[(j+2)])-1)
            { var start int = INT(ClEnv.EvalStack[(j+3)])
              { var m *ClaireAny = ANY(ClEnv.EvalStack[(j+1)])
                PRINC("break in ")
                Result = Core.F_CALL(C_print,ARGS(m.ToEID()))
                if !ErrorIn(Result) {
                PRINC("(")
                Result = Core.F_CALL(C_print,ARGS(ClEnv.EvalStack[start]))
                if !ErrorIn(Result) {
                { var i int = (start+1)
                  { var g0176 int = (start+num_args)
                    Result= EID{CFALSE.Id(),0}
                    for (i <= g0176) { 
                      var loop_3 EID
                      _ = loop_3
                      { 
                      PRINC(",")
                      loop_3 = Core.F_CALL(C_print,ARGS(ClEnv.EvalStack[i]))
                      if ErrorIn(loop_3) {Result = loop_3
                      break
                      } else {
                      }
                      if ErrorIn(loop_3) {Result = loop_3
                      break
                      } else {
                      i = (i+1)
                      }
                      } 
                    }
                    } 
                  } 
                if !ErrorIn(Result) {
                PRINC(") [q] >")
                Result = EVOID
                }}}
                if !ErrorIn(Result) {
                { var c *ClaireAny
                  var try_4 EID
                  try_4 = F_read_string(F_CommandLoopVoid())
                  if ErrorIn(try_4) {Result = try_4
                  } else {
                  c = ANY(try_4)
                  Result= EID{CFALSE.Id(),0}
                  for (c != C_q.Id()) { 
                    var loop_5 EID
                    _ = loop_5
                    { 
                    loop_5 = EVAL(c)
                    if ErrorIn(loop_5) {Result = loop_5
                    break
                    } else {
                    PRINC("break>")
                    var try_6 EID
                    try_6 = F_read_string(F_CommandLoopVoid())
                    if ErrorIn(try_6) {loop_5 = try_6
                    Result = try_6
                    break
                    } else {
                    c = ANY(try_6)
                    loop_5 = c.ToEID()
                    }}
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
      if !ErrorIn(Result) {
      { 
        var va_arg1 *ClaireEnvironment
        var va_arg2 int
        va_arg1 = ClEnv
        va_arg2 = t
        va_arg1.Trace_I = va_arg2
        Result = EID{C__INT,IVAL(va_arg2)}
        } 
      }
      } 
    } 
  return Result} 

// The EID go function for: breakpoint @ void (throw: true) 
func E_breakpoint_void (_CL_obj EID) EID { 
  return F_breakpoint_void( )} 

// the four keyword
/* The go function for: dn(x:integer) [status=0] */
func F_dn_integer (x int)  { 
  for ((INT(ClEnv.EvalStack[ToInteger(C__starcurd_star.Value).Value]) > 0) && 
      (x > 0)) { 
    C__starcurd_star.Value = ANY(ClEnv.EvalStack[ToInteger(C__starcurd_star.Value).Value])
    C__starindex_star.Value = MakeInteger((ToInteger(C__starindex_star.Value).Value+1)).Id()
    x = (x-1)
    } 
  } 

// The EID go function for: dn @ integer (throw: false) 
func E_dn_integer (x EID) EID { 
  F_dn_integer(INT(x) )
  return EVOID} 

/* The go function for: up(x:integer) [status=0] */
func F_up_integer (x int)  { 
  if (x > 0) { 
    { var indices *ClaireList = CNIL
      { var ind int = ToInteger(C__starmaxd_star.Value).Value
        for (ind != ToInteger(C__starcurd_star.Value).Value) { 
          indices = F_cons_any(MakeInteger(ind).Id(),indices)
          ind = INT(ClEnv.EvalStack[ind])
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
/* The go function for: where(x:integer) [status=1] */
func F_where_integer (x int) EID { 
  var Result EID
  { var j int = ToInteger(C__starcurd_star.Value).Value
    { var stack_level int = 0
      Result= EID{CFALSE.Id(),0}
      for ((j > 0) && 
          ((x > 0) && 
            (ClEnv.Debug_I > 0))) { 
        var loop_1 EID
        _ = loop_1
        { 
        loop_1 = F_print_debug_info_integer(j,stack_level,ToInteger(C__starindex_star.Value).Value)
        if ErrorIn(loop_1) {Result = loop_1
        break
        } else {
        stack_level = (stack_level+1)
        x = (x-1)
        j = INT(ClEnv.EvalStack[j])
        }
        } 
      }
      } 
    } 
  return Result} 

// The EID go function for: where @ integer (throw: true) 
func E_where_integer (x EID) EID { 
  return F_where_integer(INT(x) )} 

// note for interpreted methods .. they should be pushing their restriction
// on the stack vs. properties
/* The go function for: print_debug_info(iClaire/index:integer,stack_level:integer,cur_index:integer) [status=1] */
func F_print_debug_info_integer (index int,stack_level int,cur_index int) EID { 
  var Result EID
  { var num_args int = (INT(ClEnv.EvalStack[(index+2)])-1)
    { var start int = INT(ClEnv.EvalStack[(index+3)])
      { var m *ClaireAny = ANY(ClEnv.EvalStack[(index+1)])
        PRINC("debug[")
        F_princ_integer((cur_index+stack_level))
        PRINC("]>")
        { var x int = 1
          { var g0178 int = stack_level
            for (x <= g0178) { 
              PRINC(">")
              x = (x+1)
              } 
            } 
          } 
        PRINC(" ")
        Result = Core.F_CALL(C_print,ARGS(m.ToEID()))
        if !ErrorIn(Result) {
        PRINC("(")
        Result = Core.F_CALL(C_print,ARGS(ClEnv.EvalStack[start]))
        if !ErrorIn(Result) {
        { var i int = (start+1)
          { var g0179 int = (start+num_args)
            Result= EID{CFALSE.Id(),0}
            for (i <= g0179) { 
              var loop_1 EID
              _ = loop_1
              { 
              PRINC(",")
              loop_1 = Core.F_CALL(C_print,ARGS(ClEnv.EvalStack[i]))
              if ErrorIn(loop_1) {Result = loop_1
              break
              } else {
              }
              if ErrorIn(loop_1) {Result = loop_1
              break
              } else {
              i = (i+1)
              }
              } 
            }
            } 
          } 
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
/* The go function for: Show(n:integer) [status=1] */
func F_Show_integer (n int) EID { 
  var Result EID
  { var i int = ToInteger(C__starcurd_star.Value).Value
    Result= EID{CFALSE.Id(),0}
    for ((i > 0) && 
        (n > 0)) { 
      var loop_1 EID
      _ = loop_1
      { var num_args int = (INT(ClEnv.EvalStack[(i+2)])-1)
        { var start int = INT(ClEnv.EvalStack[(i+3)])
          PRINC("[")
          F_princ_integer(start)
          PRINC(" - ")
          F_princ_integer(i)
          PRINC("]: p = ")
          loop_1 = Core.F_CALL(C_print,ARGS(ClEnv.EvalStack[(i+1)]))
          if ErrorIn(loop_1) {Result = loop_1
          break
          } else {
          PRINC(", narg = ")
          loop_1 = Core.F_print_any(MakeInteger(num_args).Id())
          if ErrorIn(loop_1) {Result = loop_1
          break
          } else {
          PRINC(" \n")
          loop_1 = EVOID
          }}
          if ErrorIn(loop_1) {Result = loop_1
          break
          } else {
          { var j int = 0
            { var g0180 int = num_args
              loop_1= EID{CFALSE.Id(),0}
              for (j <= g0180) { 
                var loop_2 EID
                _ = loop_2
                { 
                PRINC("  [")
                F_princ_integer((j+i))
                PRINC("]:")
                loop_2 = Core.F_CALL(C_print,ARGS(ClEnv.EvalStack[(j+i)]))
                if ErrorIn(loop_2) {loop_1 = loop_2
                break
                } else {
                PRINC(" \n")
                loop_2 = EVOID
                }
                if ErrorIn(loop_2) {loop_1 = loop_2
                break
                } else {
                j = (j+1)
                }
                } 
              }
              } 
            } 
          if ErrorIn(loop_1) {Result = loop_1
          break
          } else {
          n = (n-1)
          i = INT(ClEnv.EvalStack[i])
          loop_1 = EID{C__INT,IVAL(i)}
          }}
          } 
        } 
      if ErrorIn(loop_1) {Result = loop_1
      break
      } else {
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
/* The go function for: block(x:integer) [status=1] */
func F_block_integer (x int) EID { 
  var Result EID
  { var j int = ToInteger(C__starcurd_star.Value).Value
    { var stack_level int = 0
      Result= EID{CFALSE.Id(),0}
      for ((j > 0) && 
          ((x > 0) && 
            (ClEnv.Debug_I > 0))) { 
        var loop_1 EID
        _ = loop_1
        { 
        { var nargs *ClaireAny = ANY(ClEnv.EvalStack[(j+2)])
          { var start int = INT(ClEnv.EvalStack[(j+3)])
            { var z *ClaireProperty = ToProperty(OBJ(ClEnv.EvalStack[(j+1)]))
              { var m *ClaireObject = Core.F_find_which_list(z.Definition,OWNER(ClEnv.EvalStack[start]),start,INT(Core.F_CALL(ToProperty(Core.C__plus.Id()),ARGS(EID{C__INT,IVAL(start)},nargs.ToEID()))))
                if (C_method.Id() == m.Isa.Id()) { 
                  { var g0181 *ClaireMethod = ToMethod(m.Id())
                    PRINC("debug[")
                    F_princ_integer((ToInteger(C__starindex_star.Value).Value+stack_level))
                    PRINC("] > ")
                    loop_1 = Core.F_print_any(g0181.Id())
                    if ErrorIn(loop_1) {Result = loop_1
                    break
                    } else {
                    PRINC("(")
                    if ((g0181.Formula.Id() != CNULL) && 
                        (g0181.Formula.Isa.IsIn(C_lambda) == CTRUE)) { 
                      { var n int = 0
                        { 
                          var v *ClaireAny
                          _ = v
                          loop_1= EID{CFALSE.Id(),0}
                          var v_support *ClaireList
                          v_support = F_closure_build_lambda(g0181.Formula)
                          v_len := v_support.Length()
                          for i_it := 0; i_it < v_len; i_it++ { 
                            v = v_support.At(i_it)
                            var loop_2 EID
                            _ = loop_2
                            { 
                            loop_2 = Core.F_CALL(C_print,ARGS(v.ToEID()))
                            if ErrorIn(loop_2) {loop_1 = loop_2
                            break
                            } else {
                            PRINC(" = ")
                            loop_2 = Core.F_CALL(C_print,ARGS(ClEnv.EvalStack[(start+n)]))
                            if ErrorIn(loop_2) {loop_1 = loop_2
                            break
                            } else {
                            PRINC(", ")
                            loop_2 = EVOID
                            }}
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
                      loop_1 = Core.F_print_any(g0181.Module_I.Id())
                      if ErrorIn(loop_1) {Result = loop_1
                      break
                      } else {
                      PRINC(">")
                      loop_1 = EVOID
                      }
                      } 
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
                  loop_1 = Core.F_print_any(z.Id())
                  if ErrorIn(loop_1) {Result = loop_1
                  break
                  } else {
                  PRINC(" -> ")
                  loop_1 = Core.F_print_any(m.Id())
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
        if ErrorIn(loop_1) {Result = loop_1
        break
        } else {
        stack_level = (stack_level+1)
        x = (x-1)
        j = INT(ClEnv.EvalStack[j])
        }
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
/* The go function for: closure_build(self:lambda) [status=0] */
func F_closure_build_lambda (self *ClaireLambda) *ClaireList { 
  var Result *ClaireList
  { var lvar *ClaireList = F_make_list_integer(self.Dimension,CEMPTY.Id())
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
/* The go function for: closure_build(self:any,lvar:list) [status=0] */
func F_closure_build_any (self *ClaireAny,lvar *ClaireList)  { 
  if (self.Isa.IsIn(C_Variable) == CTRUE) { 
    { var g0183 *ClaireVariable = To_Variable(self)
      ToArray(lvar.Id()).NthPut((g0183.Index+1),g0183.Id())
      } 
    }  else if (self.Isa.IsIn(C_Instruction) == CTRUE) { 
    { var g0184 *ClaireInstruction = To_Instruction(self)
      { 
        var s *ClaireSlot
        _ = s
        var s_iter *ClaireAny
        for _,s_iter = range(g0184.Isa.Slots.ValuesO()){ 
          s = ToSlot(s_iter)
          F_closure_build_any(Core.F_get_slot(s,ToObject(g0184.Id())),lvar)
          } 
        } 
      } 
    }  else if (self.Isa.IsIn(C_list) == CTRUE) { 
    { var g0185 *ClaireList = ToList(self)
      { 
        var x *ClaireAny
        _ = x
        var x_support *ClaireList
        x_support = g0185
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
/* The go function for: close(x:measure) [status=0] */
func F_close_measure (x *Measure) *Measure { 
  x.MIndex = C_measure.Instances.Length()
  return  x
  } 

// The EID go function for: close @ measure (throw: false) 
func E_close_measure (x EID) EID { 
  return EID{F_close_measure(ToMeasure(OBJ(x)) ).Id(),0}} 

/* The go function for: add(x:measure,f:float) [status=0] */
func (x *Measure) Add (f float64) *Measure { 
  x.NumValue = (x.NumValue+1)
  x.SumValue = (x.SumValue+f)
  x.SumSquare = (x.SumSquare+(f*f))
  return  x
  } 

// The EID go function for: add @ measure (throw: false) 
func E_add_measure (x EID,f EID) EID { 
  return EID{ToMeasure(OBJ(x)).Add(FLOAT(f) ).Id(),0}} 

/* The go function for: mean(x:measure) [status=0] */
func (x *Measure) Mean () float64 { 
  if (x.NumValue == 0) { 
    return  0
    } else {
    return  (x.SumValue/x.NumValue)
    } 
  } 

// The EID go function for: mean @ measure (throw: false) 
func E_mean_measure (x EID) EID { 
  return EID{C__FLOAT,FVAL(ToMeasure(OBJ(x)).Mean( ))}} 

/* The go function for: stdev(x:measure) [status=0] */
func (x *Measure) Stdev () float64 { 
  var Result float64
  { var y float64 = ((x.SumSquare/x.NumValue)-F__exp_float((x.SumValue/x.NumValue),2))
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

/* The go function for: stdev%(x:measure) [status=0] */
func (x *Measure) Stdev_Z () float64 { 
  return  (x.Stdev()/x.Mean())
  } 

// The EID go function for: stdev% @ measure (throw: false) 
func E_stdev_Z_measure (x EID) EID { 
  return EID{C__FLOAT,FVAL(ToMeasure(OBJ(x)).Stdev_Z( ))}} 

/* The go function for: reset(x:measure) [status=0] */
func (x *Measure) Reset ()  { 
  x.SumSquare = 0
  x.NumValue = 0
  x.SumValue = 0
  } 

// The EID go function for: reset @ measure (throw: false) 
func E_reset_measure (x EID) EID { 
  ToMeasure(OBJ(x)).Reset( )
  return EVOID} 

/* The go function for: self_print(m:measure) [status=1] */
func (m *Measure) SelfPrint () EID { 
  var Result EID
  Result = Core.F_printFDigit_float(m.Mean(),2)
  if !ErrorIn(Result) {
  PRINC("[")
  Result = Core.F_printFDigit_float(m.NumValue,0)
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
/* The go function for: logMeasure(s:string) [status=1] */
func F_logMeasure_string (s *ClaireString) EID { 
  var Result EID
  { var p *ClairePort
    var try_1 EID
    try_1 = F_fopen_string(s,MakeString("w"))
    if ErrorIn(try_1) {Result = try_1
    } else {
    p = ToPort(OBJ(try_1))
    { var n int = Core.F_size_class(C_measure)
      p.UseAsOutput()
      PRINC("// log file produced on ")
      F_princ_string(F_date_I_integer(1))
      { 
        var g0190 *ClaireClass
        _ = g0190
        var g0190_iter *ClaireAny
        var g0190_support *ClaireSet
        g0190_support = C_measure.Descendants
        for i_it := 0; i_it < g0190_support.Count; i_it++ { 
          g0190_iter = g0190_support.At(i_it)
          g0190 = ToClass(g0190_iter)
          { 
            var m *Measure
            _ = m
            var m_iter *ClaireAny
            var m_support *ClaireList
            m_support = g0190.Instances
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
/* The go function for: addLog(i:integer,x:float,y:float,n:float,s:integer) [status=1] */
func F_addLog_integer (i int,x float64,y float64,n float64,s int) EID { 
  var Result EID
  if (Core.F_size_class(C_measure) == s) { 
    { var m *Measure = ToMeasure(C_measure.Instances.At(i-1))
      m.SumValue = (m.SumValue+x)
      m.SumSquare = (m.SumSquare+y)
      { 
        var va_arg1 *Measure
        var va_arg2 float64
        va_arg1 = m
        va_arg2 = (m.NumValue+n)
        va_arg1.NumValue = va_arg2
        Result = EID{C__FLOAT,FVAL(va_arg2)}
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

// end of file