/***** CLAIRE Compilation of file /Users/ycaseau/Dropbox/src/clairev4.12/src/compile/gomain.cl 
         [version 4.1.4 / safety 5] Friday 01-03-2025 14:52:13 *****/

package Generate
import (_ "fmt"
	. "Kernel"
	"Core"
	"Language"
	"Reader"
	"Optimize"
)

//-------- dumb function to prevent import errors --------
func import_g0209() { 
  _ = Core.It
  _ = Language.It
  _ = Reader.It
  _ = Optimize.It
  } 


//+-------------------------------------------------------------+
//| CLAIRE                                                      |
//| gomain.cl                                                   |
//| Copyright (C) 1994 - 2025 Yves Caseau. All Rights Reserved  |
//| cf. copyright info in file object.cl: about()               |
//+-------------------------------------------------------------+
// ----------------------------------------------------------------------
// this file contains the code that is necessary to create a claire-based 
// executable (a "system")
// CURRENT HYPOTHESIS : 
//    (1) we do not need makefiles 
//    (2) we always use modules  (hence we have a package)
//    (3) only two modes -cm : compile a module and produces an executable
//                       -cx : create a system file (called from a make)
//                       -cc : (slave of -cf) -> generate go files
// ----------------------------------------------------------------------
// -------------------------------------------------------------------
// Contents
//      Part 1: definition of the main function
//      Part 2: Generating a system file
//      Part 3: create an executable : generate a command line and calls it
//      Part 4: Profiler code  
// -------------------------------------------------------------------
// dumb utility
/* The go function for: external!(m:module) [status=0] */
func F_Generate_external_I_module (m *ClaireModule) *ClaireString { 
  if ((m.External).Id() != CNULL) { 
    return  m.External
    } else {
    return  m.Name.String_I()
    } 
  } 

// The EID go function for: external! @ module (throw: false) 
func E_Generate_external_I_module (m EID) EID { 
  return EID{F_Generate_external_I_module(ToModule(OBJ(m)) ).Id(),0}} 

/* The go function for: string2module(s:string) [status=1] */
func F_Generate_string2module_string (s *ClaireString) EID { 
  var Result EID
  { var m *ClaireAny = F_value_string(s)
    if (m.Isa.IsIn(C_module) == CTRUE) { 
      { var g0210 *ClaireModule = ToModule(m)
        Result = EID{g0210.Id(),0}
        } 
      } else {
      Result = ToException(Core.C_general_error.Make(MakeString("~A is not a module").Id(),MakeConstantList((s).Id()).Id())).Close()
      } 
    } 
  return Result} 

// The EID go function for: string2module @ string (throw: true) 
func E_Generate_string2module_string (s EID) EID { 
  return F_Generate_string2module_string(ToString(OBJ(s)) )} 

// *******************************************************************
// *       Part 1: definition of the main function                   *
// *******************************************************************
// help file
/* The go function for: printHelp(_CL_obj:void) [status=0] */
func F_Generate_printHelp_void ()  { 
  PRINC("------------- CLAIRE: The Art of Elegant Programming -----------\n\n")
  Core.F_about_void()
  PRINC("\noptions -s <int> : set memory allocation size  \n")
  PRINC("        -f <filename>  : load <filename>             \n")
  PRINC("        -e \"<call>\" : evaluate expression such as <p(listargs)>            \n")
  PRINC("        -n : do not load the init file               \n")
  PRINC("        -m <module> : load <module>                  \n")
  PRINC("        -mx <module> : load <module> and launch main()  \n")
  PRINC("        -v <int> : sets the verbosity level       \n")
  PRINC("        -S <flag> : sets the global variable <flag> to true  \n")
  PRINC("        -o <name> : sets the name of the executable  \n")
  PRINC("        -od <name> : sets the output directory  \n")
  PRINC("        -p : profiling mode                          \n")
  PRINC("        -D : debug mode                              \n")
  PRINC("        -safe : safe mode                            \n")
  PRINC("        -O : optimizing mode                         \n")
  PRINC("        -O2 : super optimizing mode  (no bound checks)      \n")
  PRINC("        -cm <module>: compiles a module -> executable        \n")
  PRINC("        -cc <module>: compiles a module -> target go files      \n")
  PRINC("        -cx <module> : compiles a module & launch main() \n")
  PRINC("         -sf <module> : generates system file associated to a module \n")
  PRINC("         -sx <module> : generates system file that includes main() \n")
  F_CL_exit(0)
  } 

// The EID go function for: printHelp @ void (throw: false) 
func E_Generate_printHelp_void (_CL_obj EID) EID { 
  F_Generate_printHelp_void( )
  return EVOID} 

// Claire's main - the rich version that starts the compiler
// a simpler version exists in inspect.cl (simple_main())
// -s is ignored because it is trapped earlier (see the file generator)
/* The go function for: complex_main(_CL_obj:void) [status=1] */
func F_Generate_complex_main_void () EID { 
  var Result EID
  { var _Zcm *ClaireString = MakeString("")
    { var _Zsf *ClaireString = MakeString("")
      { var _Zout *ClaireString = MakeString("")
        { var dblevel int = 1
          { var vlevel int = 1
            { var _Zinit_ask *ClaireBoolean = CTRUE
              { var _Zexe *ClaireBoolean = CFALSE
                { var _Zsafety *ClaireAny = CNULL
                  { var _Zmain *ClaireBoolean = CFALSE
                    { var l *ClaireList = ClEnv.Params.Copy()
                      { 
                        h_index := ClEnv.Index
                        h_base := ClEnv.Base
                        Result= EID{CFALSE.Id(),0}
                        for (l.Length() != 0) { 
                          var loop_1 EID
                          _ = loop_1
                          if ((ToString(l.ValuesO()[0]).Value == MakeString("?").Value) || 
                              ((ToString(l.ValuesO()[0]).Value == MakeString("-h").Value) || 
                                (ToString(l.ValuesO()[0]).Value == MakeString("-help").Value))) { 
                            F_Generate_printHelp_void()
                            loop_1 = EVOID
                            }  else if (ToString(l.ValuesO()[0]).Value == MakeString("-q").Value) { 
                            vlevel = 0
                            l = l.Skip(1)
                            loop_1 = EID{l.Id(),0}
                            }  else if (ToString(l.ValuesO()[0]).Value == MakeString("-v").Value) { 
                            vlevel = 2
                            l = l.Skip(1)
                            loop_1 = EID{l.Id(),0}
                            }  else if (ToString(l.ValuesO()[0]).Value == MakeString("-s").Value) { 
                            if (l.Length() >= 2) { 
                              l = l.Skip(2)
                              loop_1 = EID{l.Id(),0}
                              } else {
                              loop_1 = ToException(Core.C_general_error.Make(MakeString("option: -s <s1> <s2>").Id(),CNIL.Id())).Close()
                              } 
                            }  else if (ToString(l.ValuesO()[0]).Value == MakeString("-f").Value) { 
                            if (l.Length() >= 2) { 
                              loop_1 = Reader.F_load_string(ToString(l.ValuesO()[1]))
                              if ErrorIn(loop_1) {Result = loop_1
                              break
                              } else {
                              l = l.Skip(2)
                              loop_1 = EID{l.Id(),0}
                              }
                              } else {
                              loop_1 = ToException(Core.C_general_error.Make(MakeString("option: -f <filename>").Id(),CNIL.Id())).Close()
                              } 
                            }  else if (ToString(l.ValuesO()[0]).Value == MakeString("-e").Value) { 
                            if (l.Length() >= 2) { 
                              { var c *ClaireAny
                                var try_2 EID
                                try_2 = Reader.F_read_string(ToString(l.ValuesO()[1]))
                                if ErrorIn(try_2) {loop_1 = try_2
                                } else {
                                c = ANY(try_2)
                                l = l.Skip(2)
                                if (c.Isa.IsIn(Language.C_Call) == CTRUE) { 
                                  { var g0214 *Language.Call = Language.To_Call(c)
                                    loop_1 = EVAL(g0214.Id())
                                    } 
                                  } else {
                                  loop_1 = ToException(Core.C_general_error.Make(MakeString("-e arg ~S is not a call").Id(),MakeConstantList(c).Id())).Close()
                                  } 
                                if ErrorIn(loop_1) {Result = loop_1
                                break
                                } else {
                                }
                                }
                                } 
                              } else {
                              loop_1 = ToException(Core.C_general_error.Make(MakeString("option: -e <call>").Id(),CNIL.Id())).Close()
                              } 
                            }  else if (ToString(l.ValuesO()[0]).Value == MakeString("-m").Value) { 
                            if (l.Length() >= 2) { 
                              if (_Zinit_ask == CTRUE) { 
                                loop_1 = Reader.F_load_string(MakeString("init"))
                                if ErrorIn(loop_1) {Result = loop_1
                                break
                                } else {
                                _Zinit_ask = CFALSE
                                loop_1 = EID{_Zinit_ask.Id(),0}
                                }
                                } else {
                                loop_1 = EID{CFALSE.Id(),0}
                                } 
                              if ErrorIn(loop_1) {Result = loop_1
                              break
                              } else {
                              { var m *ClaireModule
                                var try_3 EID
                                try_3 = F_Generate_string2module_string(ToString(l.ValuesO()[1]))
                                if ErrorIn(try_3) {loop_1 = try_3
                                } else {
                                m = ToModule(OBJ(try_3))
                                loop_1 = Reader.F_load_module(m)
                                if ErrorIn(loop_1) {Result = loop_1
                                break
                                } else {
                                m.Begin()
                                l = l.Skip(2)
                                var v_gassign4 *ClaireAny
                                v_gassign4 = ToList(Optimize.C_claire_modules.Value).AddFast(m.Id()).Id()
                                Optimize.C_claire_modules.Value = v_gassign4
                                loop_1 = v_gassign4.ToEID()
                                }
                                }
                                } 
                              if ErrorIn(loop_1) {Result = loop_1
                              break
                              } else {
                              }}
                              } else {
                              loop_1 = ToException(Core.C_general_error.Make(MakeString("option: -m <module>").Id(),CNIL.Id())).Close()
                              } 
                            }  else if (ToString(l.ValuesO()[0]).Value == MakeString("-mx").Value) { 
                            if (l.Length() >= 2) { 
                              if (_Zinit_ask == CTRUE) { 
                                loop_1 = Reader.F_load_string(MakeString("init"))
                                if ErrorIn(loop_1) {Result = loop_1
                                break
                                } else {
                                _Zinit_ask = CFALSE
                                loop_1 = EID{_Zinit_ask.Id(),0}
                                }
                                } else {
                                loop_1 = EID{CFALSE.Id(),0}
                                } 
                              if ErrorIn(loop_1) {Result = loop_1
                              break
                              } else {
                              { var m *ClaireModule
                                var try_5 EID
                                try_5 = F_Generate_string2module_string(ToString(l.ValuesO()[1]))
                                if ErrorIn(try_5) {loop_1 = try_5
                                } else {
                                m = ToModule(OBJ(try_5))
                                loop_1 = Reader.F_load_module(m)
                                if ErrorIn(loop_1) {Result = loop_1
                                break
                                } else {
                                Core.F_CALL(Core.C_main,ARGS(EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}))
                                l = l.Skip(2)
                                var v_gassign6 *ClaireAny
                                v_gassign6 = ToList(Optimize.C_claire_modules.Value).AddFast(m.Id()).Id()
                                Optimize.C_claire_modules.Value = v_gassign6
                                loop_1 = v_gassign6.ToEID()
                                }
                                }
                                } 
                              if ErrorIn(loop_1) {Result = loop_1
                              break
                              } else {
                              }}
                              } else {
                              loop_1 = ToException(Core.C_general_error.Make(MakeString("option: -m <module>").Id(),CNIL.Id())).Close()
                              } 
                            }  else if (ToString(l.ValuesO()[0]).Value == MakeString("-v").Value) { 
                            if (l.Length() >= 2) { 
                              vlevel = (vlevel+F_integer_I_string(ToString(l.ValuesO()[1])))
                              l = l.Skip(2)
                              loop_1 = EID{l.Id(),0}
                              } else {
                              loop_1 = ToException(Core.C_general_error.Make(MakeString("option: -v <integer>").Id(),CNIL.Id())).Close()
                              } 
                            }  else if (ToString(l.ValuesO()[0]).Value == MakeString("-S").Value) { 
                            if (l.Length() >= 2) { 
                              { 
                                var va_arg1 *Core.GlobalVariable
                                var va_arg2 *ClaireAny
                                var try_7 EID
                                try_7 = Core.F_new_class2(Core.C_global_variable,Core.F_symbol_I_string2(ToString(l.ValuesO()[1])))
                                if ErrorIn(try_7) {loop_1 = try_7
                                } else {
                                va_arg1 = Core.ToGlobalVariable(OBJ(try_7))
                                va_arg2 = CTRUE.Id()
                                va_arg1.Value = va_arg2
                                loop_1 = va_arg2.ToEID()
                                }
                                } 
                              if ErrorIn(loop_1) {Result = loop_1
                              break
                              } else {
                              l = l.Skip(2)
                              loop_1 = EID{l.Id(),0}
                              }
                              } else {
                              loop_1 = ToException(Core.C_general_error.Make(MakeString("option: -S <FLAG>").Id(),CNIL.Id())).Close()
                              } 
                            }  else if (ToString(l.ValuesO()[0]).Value == MakeString("-od").Value) { 
                            if (l.Length() >= 2) { 
                              ToGenerateGoProducer(Optimize.C_PRODUCER.Value).Output = ToString(l.ValuesO()[1])
                              l = l.Skip(2)
                              loop_1 = EID{l.Id(),0}
                              } else {
                              loop_1 = ToException(Core.C_general_error.Make(MakeString("option: -od <directory>").Id(),CNIL.Id())).Close()
                              } 
                            }  else if (ToString(l.ValuesO()[0]).Value == MakeString("-o").Value) { 
                            if (l.Length() >= 2) { 
                              _Zout = ToString(l.ValuesO()[1])
                              l = l.Skip(2)
                              loop_1 = EID{l.Id(),0}
                              } else {
                              loop_1 = ToException(Core.C_general_error.Make(MakeString("option: -o <name>").Id(),CNIL.Id())).Close()
                              } 
                            }  else if (ToString(l.ValuesO()[0]).Value == MakeString("-p").Value) { 
                            Optimize.C_OPT.Profile_ask = CTRUE
                            if (dblevel <= 2) { 
                              dblevel = 2
                              } else {
                              dblevel = dblevel
                              } 
                            l = l.Skip(1)
                            loop_1 = EID{l.Id(),0}
                            }  else if (ToString(l.ValuesO()[0]).Value == MakeString("-D").Value) { 
                            dblevel = 0
                            l = l.Skip(1)
                            loop_1 = EID{l.Id(),0}
                            }  else if (ToString(l.ValuesO()[0]).Value == MakeString("-O").Value) { 
                            Optimize.C_compiler.Optimize_ask = CTRUE
                            _Zsafety = MakeInteger(2).Id()
                            dblevel = 2
                            l = l.Skip(1)
                            loop_1 = EID{l.Id(),0}
                            }  else if (ToString(l.ValuesO()[0]).Value == MakeString("-O2").Value) { 
                            Optimize.C_compiler.Optimize_ask = CTRUE
                            _Zsafety = MakeInteger(3).Id()
                            dblevel = 2
                            l = l.Skip(1)
                            loop_1 = EID{l.Id(),0}
                            }  else if (ToString(l.ValuesO()[0]).Value == MakeString("-cc").Value) { 
                            if (l.Length() >= 2) { 
                              _Zcm = ToString(l.ValuesO()[1])
                              l = l.Skip(2)
                              loop_1 = EID{l.Id(),0}
                              } else {
                              loop_1 = ToException(Core.C_general_error.Make(MakeString("option: -cc <module>").Id(),CNIL.Id())).Close()
                              } 
                            }  else if (ToString(l.ValuesO()[0]).Value == MakeString("-cm").Value) { 
                            if (l.Length() >= 2) { 
                              _Zexe = CTRUE
                              _Zcm = ToString(l.ValuesO()[1])
                              l = l.Skip(2)
                              loop_1 = EID{l.Id(),0}
                              } else {
                              loop_1 = ToException(Core.C_general_error.Make(MakeString("option: -cm <module>").Id(),CNIL.Id())).Close()
                              } 
                            }  else if (ToString(l.ValuesO()[0]).Value == MakeString("-cx").Value) { 
                            if (l.Length() >= 2) { 
                              _Zexe = CTRUE
                              _Zcm = ToString(l.ValuesO()[1])
                              _Zmain = CTRUE
                              l = l.Skip(2)
                              loop_1 = EID{l.Id(),0}
                              } else {
                              loop_1 = ToException(Core.C_general_error.Make(MakeString("option: -cm <module>").Id(),CNIL.Id())).Close()
                              } 
                            }  else if (ToString(l.ValuesO()[0]).Value == MakeString("-sf").Value) { 
                            if (l.Length() >= 2) { 
                              _Zsf = ToString(l.ValuesO()[1])
                              l = l.Skip(2)
                              loop_1 = EID{l.Id(),0}
                              } else {
                              loop_1 = ToException(Core.C_general_error.Make(MakeString("option: -sf <filename>").Id(),CNIL.Id())).Close()
                              } 
                            }  else if (ToString(l.ValuesO()[0]).Value == MakeString("-sx").Value) { 
                            if (l.Length() >= 2) { 
                              _Zsf = ToString(l.ValuesO()[1])
                              _Zmain = CTRUE
                              l = l.Skip(2)
                              loop_1 = EID{l.Id(),0}
                              } else {
                              loop_1 = ToException(Core.C_general_error.Make(MakeString("option: -sx <filename>").Id(),CNIL.Id())).Close()
                              } 
                            }  else if (ToString(l.ValuesO()[0]).Value == MakeString("-n").Value) { 
                            _Zinit_ask = CFALSE
                            l = l.Skip(1)
                            loop_1 = EID{l.Id(),0}
                            } else {
                            if (ToString(l.ValuesO()[0]).At(1) == '-') { 
                              loop_1 = Core.F_print_any(l.ValuesO()[0])
                              if ErrorIn(loop_1) {Result = loop_1
                              break
                              } else {
                              PRINC(" is an unvalid option\n")
                              loop_1 = EVOID
                              }
                              if ErrorIn(loop_1) {Result = loop_1
                              break
                              } else {
                              F_Generate_printHelp_void()
                              loop_1 = EVOID
                              }
                              } else {
                              loop_1 = EID{CFALSE.Id(),0}
                              } 
                            if ErrorIn(loop_1) {Result = loop_1
                            break
                            } else {
                            l = ToType(C_string.Id()).EmptyList()
                            loop_1 = EID{l.Id(),0}
                            }
                            } 
                          if ErrorIn(loop_1) {Result = loop_1
                          break
                          } else {
                          } 
                        }
                        if !ErrorIn(Result) {
                        if (_Zout.Value == MakeString("").Value) { 
                          if (_Zcm.Value != MakeString("").Value) { 
                            _Zout = _Zcm
                            }  else if (_Zsf.Value != MakeString("").Value) { 
                            _Zout = _Zsf
                            } 
                          } 
                        if (_Zinit_ask == CTRUE) { 
                          Result = Reader.F_load_string(MakeString("init"))
                          } else {
                          Result = EID{CFALSE.Id(),0}
                          } 
                        if !ErrorIn(Result) {
                        ClEnv.Verbose = vlevel
                        if (_Zsafety != CNULL) { 
                          Optimize.C_compiler.Safety = ToInteger(_Zsafety).Value
                          } 
                        if (_Zsf.Value != MakeString("").Value) { 
                          Result = Core.F_CALL(Reader.C_load,ARGS(F_value_string(MakeString("Compile")).ToEID()))
                          if !ErrorIn(Result) {
                          Optimize.C_compiler.Active_ask = CTRUE
                          { var arg_8 *ClaireModule
                            var try_9 EID
                            try_9 = F_Generate_string2module_string(_Zsf)
                            if ErrorIn(try_9) {Result = try_9
                            } else {
                            arg_8 = ToModule(OBJ(try_9))
                            Result = F_system_file_module(arg_8,_Zout,_Zmain)
                            }
                            } 
                          if !ErrorIn(Result) {
                          F_CL_exit(0)
                          Result = EVOID
                          }}
                          }  else if (_Zcm.Value != MakeString("").Value) { 
                          { var m *ClaireModule
                            var try_10 EID
                            try_10 = F_Generate_string2module_string(_Zcm)
                            if ErrorIn(try_10) {Result = try_10
                            } else {
                            m = ToModule(OBJ(try_10))
                            if (Equal(m.Uses.Id(),MakeConstantList(ToList(Optimize.C_claire_modules.Value).At(1)).Id()) == CTRUE) { 
                              Optimize.C_claire_modules.Value = ToList(Optimize.C_claire_modules.Value).Shrink(2).Id()
                              Core.F_tformat_string(MakeString("=== Light Module ~S:~S -> use ~S=== "),0,MakeConstantList(m.Id(),m.Uses.Id(),Optimize.C_claire_modules.Value))
                              } 
                            Optimize.C_claire_modules.Value = ToList(Optimize.C_claire_modules.Value).AddFast(m.Id()).Id()
                            Result = Core.F_CALL(Reader.C_load,ARGS(F_value_string(MakeString("Compile")).ToEID()))
                            if !ErrorIn(Result) {
                            ClEnv.Jito_ask = CFALSE
                            Optimize.C_compiler.Active_ask = CTRUE
                            if (_Zout.Value != MakeString("").Value) { 
                              m.External = _Zout
                              } 
                            Result = Reader.F_load_module(m)
                            if !ErrorIn(Result) {
                            if (dblevel < 1) { 
                              { 
                                var va_arg1 *Optimize.OptimizeMetaCompiler
                                var va_arg2 int
                                va_arg1 = Optimize.C_compiler
                                if (Optimize.C_compiler.Safety <= 1) { 
                                  va_arg2 = Optimize.C_compiler.Safety
                                  } else {
                                  va_arg2 = 1
                                  } 
                                va_arg1.Safety = va_arg2
                                } 
                              Optimize.C_compiler.Debug_ask = Optimize.C_compiler.Debug_ask.AddFast(m.Id())
                              } 
                            F_Generate_compile_dir_module(m)
                            Result = Core.F_CALL(C_compile,ARGS(EID{m.Id(),0}))
                            if !ErrorIn(Result) {
                            if (_Zexe == CTRUE) { 
                              Core.F_tformat_string(MakeString("==== create the system file for module ~S \n"),0,MakeConstantList((_Zout).Id()))
                              Result = F_system_file_module(m,_Zout,_Zmain)
                              if !ErrorIn(Result) {
                              F_Generate_compile_exe_string(_Zout)
                              Result = EVOID
                              }
                              } else {
                              Result = EID{CFALSE.Id(),0}
                              } 
                            if !ErrorIn(Result) {
                            F_CL_exit(0)
                            Result = EVOID
                            }}}}
                            }
                            } 
                          } else {
                          Result = Reader.C_reader.TopLevel()
                          } 
                        }}
                        if ErrorIn(Result){ 
                          ClEnv.Index = h_index
                          ClEnv.Base = h_base
                          Reader.C_reader.RestoreState()
                          PRINC("\nCLAIRE error during init [line ")
                          F_princ_integer(ClEnv.NLine)
                          PRINC("]:\n")
                          Result = Reader.F_debug_if_possible_void()
                          if !ErrorIn(Result) {
                          PRINC("\n")
                          Result = Reader.C_reader.TopLevel()
                          }
                          } 
                        } 
                      } 
                    } 
                  } 
                } 
              } 
            } 
          } 
        } 
      } 
    } 
  return Result} 

// The EID go function for: complex_main @ void (throw: true) 
func E_Generate_complex_main_void (_CL_obj EID) EID { 
  return F_Generate_complex_main_void( )} 

// *******************************************************************
// *       Part 2: System compiling methods                          *
// *******************************************************************
// generate a system file with
//   - the import
//   - the module definition
//   - calling the load() methods for the meta-descriptions
//   - the main function
/* The go function for: system_file(m:module,%out:string,%main:boolean) [status=1] */
func F_system_file_module (m *ClaireModule,_Zout *ClaireString,_Zmain *ClaireBoolean) EID { 
  var Result EID
  { var p *ClairePort
    var try_1 EID
    try_1 = F_fopen_string(F_append_string(Reader.F__7_string(Optimize.C_compiler.Source,_Zout),ToGenerateCodeProducer(Optimize.C_PRODUCER.Value).Extension),MakeString("w"))
    if ErrorIn(try_1) {Result = try_1
    } else {
    p = ToPort(OBJ(try_1))
    { var l_used *ClaireList = Reader.F_add_modules_list(MakeConstantList(m.Id()))
      { var l_necessary *ClaireList = F_Generate_parents_list(l_used)
        ToGenerateGoProducer(Optimize.C_PRODUCER.Value).Current = C_claire
        Optimize.C_OPT.Properties = ToType(C_property.Id()).EmptySet()
        Optimize.C_OPT.Objects = ToType(C_any.Id()).EmptyList()
        Optimize.C_OPT.Functions = ToType(C_any.Id()).EmptyList()
        Optimize.C_OPT.NeedToClose = ToType(C_any.Id()).EmptySet()
        Optimize.C_OPT.LegalModules = l_necessary.Set_I()
        p.UseAsOutput()
        PRINC("// --- System configuration file for ")
        Result = Core.F_print_any((_Zout).Id())
        if !ErrorIn(Result) {
        PRINC(" , [")
        Result = Core.F_print_any((F_date_I_integer(1)).Id())
        if !ErrorIn(Result) {
        PRINC("] ---\n\n")
        Result = EVOID
        }}
        if !ErrorIn(Result) {
        PRINC("package main\n")
        Result = F_Generate_system_imports_module(m)
        if !ErrorIn(Result) {
        Result = F_Generate_load_function_module(m,l_necessary)
        if !ErrorIn(Result) {
        F_Generate_main_function_module(m,l_used,_Zmain)
        p.Fclose()
        Result = EVOID
        }}}
        } 
      } 
    }
    } 
  return Result} 

// The EID go function for: system_file @ module (throw: true) 
func E_system_file_module (m EID,_Zout EID,_Zmain EID) EID { 
  return F_system_file_module(ToModule(OBJ(m)),ToString(OBJ(_Zout)),ToBoolean(OBJ(_Zmain)) )} 

// create the import declaration for this system file
/* The go function for: system_imports(m:module) [status=1] */
func F_Generate_system_imports_module (m *ClaireModule) EID { 
  var Result EID
  PRINC("import (\n")
  if (Optimize.C_OPT.Profile_ask == CTRUE) { 
    PRINC("\t\"os\"\n\t\"runtime/pprof\"\n")
    } 
  Result = F_Generate_import_declaration_module(m)
  if !ErrorIn(Result) {
  PRINC("\t\"")
  F_princ_string(m.Name.String_I())
  PRINC("\"\n")
  PRINC(")\n")
  Result = EVOID
  }
  return Result} 

// The EID go function for: system_imports @ module (throw: true) 
func E_Generate_system_imports_module (m EID) EID { 
  return F_Generate_system_imports_module(ToModule(OBJ(m)) )} 

// called by the run_system function. The goal of this method is
// to create the modules, and then loads all the meta-descriptions with m.MetaLoad()
// note that is status(m) = 5, we defer the load of the modules.
// we should have a declaration like deferred(m) :: m.status := 5
/* The go function for: load_function(m:module,l_necessary:list) [status=1] */
func F_Generate_load_function_module (m *ClaireModule,l_necessary *ClaireList) EID { 
  var Result EID
  { 
    var x *ClaireAny
    _ = x
    var x_support *ClaireList
    x_support = l_necessary
    x_len := x_support.Length()
    for i_it := 0; i_it < x_len; i_it++ { 
      x = x_support.At(i_it)
      if ANY(Core.F_CALL(C_mClaire_status,ARGS(x.ToEID()))).IsInt(5) { 
        PRINC("func load_")
        Core.F_CALL(Language.C_iClaire_ident,ARGS(Core.F_CALL(C_name,ARGS(x.ToEID()))))
        PRINC("() {")
        Core.F_CALL(Language.C_iClaire_ident,ARGS(Core.F_CALL(C_name,ARGS(x.ToEID()))))
        PRINC(".metaLoad();}\n")
        } 
      } 
    } 
  PRINC("\n//load function : create and load modules")
  F_Generate_breakline_void()
  PRINC("func Load() ")
  F_Generate_new_block_void()
  PRINC("It := C_claire")
  F_Generate_breakline_void()
  PRINC("//module definitions ")
  F_Generate_breakline_void()
  { 
    var x *ClaireAny
    _ = x
    Result= EID{CFALSE.Id(),0}
    var x_support *ClaireList
    x_support = l_necessary
    x_len := x_support.Length()
    for i_it := 0; i_it < x_len; i_it++ { 
      x = x_support.At(i_it)
      var loop_1 EID
      _ = loop_1
      if ((x != C_claire.Id()) && 
          ((x != C_mClaire.Id()) && 
            (x != C_Kernel.Id()))) { 
        loop_1 = Core.F_CALL(C_Generate_g_expression,ARGS(x.ToEID(),EID{C_module.Id(),0}))
        if ErrorIn(loop_1) {Result = loop_1
        break
        } else {
        PRINC(" = InitModule(")
        { var arg_2 *ClaireAny
          var try_3 EID
          try_3 = Core.F_CALL(C_string_I,ARGS(Core.F_CALL(C_name,ARGS(x.ToEID()))))
          if ErrorIn(try_3) {loop_1 = try_3
          } else {
          arg_2 = ANY(try_3)
          loop_1 = Core.F_print_any(arg_2)
          }
          } 
        if ErrorIn(loop_1) {Result = loop_1
        break
        } else {
        PRINC(",")
        loop_1 = F_Generate_g_expression_module(ToModule(x).PartOf,C_module)
        if ErrorIn(loop_1) {Result = loop_1
        break
        } else {
        PRINC(",")
        { var arg_4 *ClaireAny
          var try_5 EID
          try_5 = Core.F_CALL(Optimize.C_c_code,ARGS(EID{ToModule(x).Uses.Id(),0},EID{C_list.Id(),0}))
          if ErrorIn(try_5) {loop_1 = try_5
          } else {
          arg_4 = ANY(try_5)
          loop_1 = Core.F_CALL(C_Generate_g_expression,ARGS(arg_4.ToEID(),EID{C_list.Id(),0}))
          }
          } 
        if ErrorIn(loop_1) {Result = loop_1
        break
        } else {
        PRINC(",")
        F_Generate_breakline_void()
        PRINC("\t")
        loop_1 = Core.F_print_any(ANY(Core.F_CALL(C_source,ARGS(x.ToEID()))))
        if ErrorIn(loop_1) {Result = loop_1
        break
        } else {
        PRINC(",\n\t")
        { var arg_6 *ClaireAny
          var try_7 EID
          try_7 = Core.F_CALL(Optimize.C_c_code,ARGS(EID{ToModule(x).MadeOf.Id(),0},EID{C_list.Id(),0}))
          if ErrorIn(try_7) {loop_1 = try_7
          } else {
          arg_6 = ANY(try_7)
          loop_1 = Core.F_CALL(C_Generate_g_expression,ARGS(arg_6.ToEID(),EID{C_list.Id(),0}))
          }
          } 
        if ErrorIn(loop_1) {Result = loop_1
        break
        } else {
        PRINC(")")
        loop_1 = F_Generate_breakline_void().ToEID()
        }}}}}}
        } else {
        loop_1 = EID{CFALSE.Id(),0}
        } 
      if ErrorIn(loop_1) {Result = loop_1
      break
      } else {
      }
      } 
    } 
  if !ErrorIn(Result) {
  F_Generate_breakline_void()
  PRINC("// module load ")
  F_Generate_breakline_void()
  { 
    var x *ClaireAny
    _ = x
    var x_support *ClaireList
    x_support = l_necessary
    x_len := x_support.Length()
    for i_it := 0; i_it < x_len; i_it++ { 
      x = x_support.At(i_it)
      if ((ToModule(x).MadeOf.Length() != 0) && 
          !ANY(Core.F_CALL(C_mClaire_status,ARGS(x.ToEID()))).IsInt(5)) { 
        Core.F_CALL(Language.C_iClaire_ident,ARGS(Core.F_CALL(C_name,ARGS(x.ToEID()))))
        PRINC(".MetaLoad()")
        F_Generate_breakline_void()
        } 
      } 
    } 
  { 
    var x *ClaireAny
    _ = x
    Result= EID{CFALSE.Id(),0}
    var x_support *ClaireList
    x_support = l_necessary
    x_len := x_support.Length()
    for i_it := 0; i_it < x_len; i_it++ { 
      x = x_support.At(i_it)
      var loop_8 EID
      _ = loop_8
      if ANY(Core.F_CALL(C_mClaire_status,ARGS(x.ToEID()))).IsInt(5) { 
        Core.F_CALL(Language.C_iClaire_ident,ARGS(Core.F_CALL(C_name,ARGS(x.ToEID()))))
        PRINC(".it->evaluate = ")
        { var arg_9 *ClaireFunction
          var try_10 EID
          { var arg_11 *ClaireString
            var try_12 EID
            { var arg_13 *ClaireAny
              var try_14 EID
              try_14 = Core.F_CALL(C_string_I,ARGS(Core.F_CALL(C_name,ARGS(x.ToEID()))))
              if ErrorIn(try_14) {try_12 = try_14
              } else {
              arg_13 = ANY(try_14)
              try_12 = EID{F_append_string(MakeString("load_"),ToString(arg_13)).Id(),0}
              }
              } 
            if ErrorIn(try_12) {try_10 = try_12
            } else {
            arg_11 = ToString(OBJ(try_12))
            try_10 = F_make_function_string(arg_11).ToEID()
            }
            } 
          if ErrorIn(try_10) {loop_8 = try_10
          } else {
          arg_9 = ToFunction(OBJ(try_10))
          loop_8 = F_Generate_g_expression_function(arg_9,C_function)
          }
          } 
        if ErrorIn(loop_8) {Result = loop_8
        break
        } else {
        loop_8 = F_Generate_breakline_void().ToEID()
        }
        if ErrorIn(loop_8) {Result = loop_8
        break
        } else {
        Core.F_CALL(Language.C_iClaire_ident,ARGS(Core.F_CALL(C_name,ARGS(x.ToEID()))))
        PRINC(".it->status = 2;")
        loop_8 = F_Generate_breakline_void().ToEID()
        }
        } else {
        loop_8 = EID{CFALSE.Id(),0}
        } 
      if ErrorIn(loop_8) {Result = loop_8
      break
      } else {
      }
      } 
    } 
  if !ErrorIn(Result) {
  PRINC("ClEnv.Module_I = ")
  Result = F_Generate_g_expression_module(m,C_module)
  if !ErrorIn(Result) {
  PRINC("; ")
  Result = F_Generate_breakline_void().ToEID()
  }
  if !ErrorIn(Result) {
  F_Generate_close_block_void()
  Result = EVOID
  }}}
  return Result} 

// The EID go function for: load_function @ module (throw: true) 
func E_Generate_load_function_module (m EID,l_necessary EID) EID { 
  return F_Generate_load_function_module(ToModule(OBJ(m)),ToList(OBJ(l_necessary)) )} 

// create the main function
// %main = true means call main()
/* The go function for: main_function(m:module,l_used:list[module],%main:boolean) [status=0] */
func F_Generate_main_function_module (m *ClaireModule,l_used *ClaireList,_Zmain *ClaireBoolean)  { 
  PRINC("\n// reboot function is created by compiler \n")
  PRINC("func E_reboot(s EID) EID ")
  F_Generate_new_block_void()
  PRINC("Bootstrap()")
  F_Generate_breakline_void()
  PRINC("Load()")
  F_Generate_breakline_void()
  PRINC("Reader.C_reader.Fromp = ClEnv.Cin;")
  F_Generate_breakline_void()
  if (ToBoolean(l_used.Contain_ask(F_value_string(MakeString("Generate"))).Id()) == CTRUE) { 
    PRINC("ClEnv.Module_I = C_claire")
    F_Generate_breakline_void()
    } 
  PRINC("return EVOID")
  F_Generate_breakline_void()
  F_Generate_close_block_void()
  F_Generate_breakline_void()
  PRINC("\n// the main function \n")
  PRINC("func main() ")
  F_Generate_new_block_void()
  PRINC("MemoryFlags()")
  F_Generate_breakline_void()
  if (Optimize.C_OPT.Profile_ask == CTRUE) { 
    PRINC("// instruction for GO profiling - to be used with go tool pprof <m.prof>\n")
    F_Generate_new_block_void()
    PRINC("f,err := os.Create(\"")
    F_princ_string(m.Name.String_I())
    PRINC(".prof\")")
    F_Generate_breakline_void()
    PRINC("if err == nil ")
    F_Generate_new_block_void()
    PRINC("pprof.StartCPUProfile(f)")
    F_Generate_breakline_void()
    PRINC("defer pprof.StopCPUProfile()")
    F_Generate_close_block_void()
    F_Generate_close_block_void()
    } 
  PRINC("E_reboot(EVOID)")
  F_Generate_breakline_void()
  PRINC("ToMethod(C_reboot.Restrictions.ValuesO()[0]).Functional = MakeFunction1(E_reboot,\"reboot\")")
  F_Generate_breakline_void()
  PRINC("Reader.C_reader.Fromp = ClEnv.Cin")
  F_Generate_breakline_void()
  if (_Zmain == CTRUE) { 
    PRINC("Core.F_CALL(Core.C_main,ARGS(EID{ClEnv.Id(),0}))")
    }  else if (ToBoolean(l_used.Contain_ask(F_value_string(MakeString("Generate"))).Id()) == CTRUE) { 
    PRINC("Generate.F_Generate_complex_main_void()")
    } else {
    PRINC("Reader.F_Reader_simple_main_void()")
    } 
  F_Generate_breakline_void()
  F_Generate_close_block_void()
  } 

// The EID go function for: main_function @ module (throw: false) 
func E_Generate_main_function_module (m EID,l_used EID,_Zmain EID) EID { 
  F_Generate_main_function_module(ToModule(OBJ(m)),ToList(OBJ(l_used)),ToBoolean(OBJ(_Zmain)) )
  return EVOID} 

// *******************************************************************
// *       Part 3: module compiling : execute a command line         *
// *******************************************************************
// create a directory for the module (if it does not exist)
/* The go function for: compile_dir(m:module) [status=0] */
func F_Generate_compile_dir_module (m *ClaireModule) EID { 
  var Result EID
  { var s *ClaireString = Reader.F__7_string(Reader.F__7_string(Reader.F__7_string(F_append_string(MakeString("mkdir -p "),Optimize.F_home_void()),MakeString("go")),MakeString("src")),F_Generate_capitalize_string(m.Name.String_I()))
    if (Optimize.F_home_void().Value == MakeString("").Value) { 
      Result = ToException(Core.C_general_error.Make(MakeString("the environment variable CLAIRE_HOME is undefined\n").Id(),CNIL.Id())).Close()
      } else {
      Result = EID{CFALSE.Id(),0}
      } 
    if !ErrorIn(Result) {
    F_claire_shell(s)
    Result = EVOID
    }
    } 
  return Result} 

// The EID go function for: compile_dir @ module (throw: true) 
func E_Generate_compile_dir_module (m EID) EID { 
  return F_Generate_compile_dir_module(ToModule(OBJ(m)) )} 

// create the go
/* The go function for: compile_exe(%out:string) [status=0] */
func F_Generate_compile_exe_string (_Zout *ClaireString)  { 
  { var outdir *ClaireString
    if ((ToGenerateGoProducer(Optimize.C_PRODUCER.Value).Output).Id() != CNULL) { 
      outdir = F_append_string(F_append_string(MakeString(" -o "),ToGenerateGoProducer(Optimize.C_PRODUCER.Value).Output),MakeString(" "))
      } else {
      outdir = MakeString("")
      } 
    { var s *ClaireString = F_append_string(Reader.F__7_string(Reader.F__7_string(Reader.F__7_string(F_append_string(F_append_string(MakeString("go build "),outdir),Optimize.F_home_void()),MakeString("go")),MakeString("src")),_Zout),MakeString(".go"))
      F_claire_shell(s)
      } 
    } 
  } 

// The EID go function for: compile_exe @ string (throw: false) 
func E_Generate_compile_exe_string (_Zout EID) EID { 
  F_Generate_compile_exe_string(ToString(OBJ(_Zout)) )
  return EVOID} 

// *******************************************************************
// *       Part 4: Profiler code                                     *
// *******************************************************************
// ---------------------------------------------------------------