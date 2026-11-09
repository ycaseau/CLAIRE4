/***** CLAIRE Compilation of file /Users/ycaseau/Dropbox/src/clairev4.12/src/compile/gomain.cl 
         [version 4.1.6 / safety 5] Saturday 08-09-2025 06:42:35 *****/

package Generate
import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
	. "Kernel"
	"Core"
	"Language"
	"Reader"
	"Optimize"
)

//-------- dumb function to prevent import errors --------
func import_g0212() { 
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
      { var g0213 *ClaireModule = ToModule(m)
        Result = EID{g0213.Id(),0}
        } 
      } else {
      Result = ToException(Core.C_general_error.Make(MakeString("~A is not a module").Id(),MakeConstantList((s).Id()).Id())).Close()
      } 
    } 
  return Result} 

// dirname() : string -- name of the current directory (v4.1.6)
// Used when -cm / -cc / -cx is given without a module name.
func F_Generate_dirname_void () *ClaireString {
  d, err := os.Getwd()
  if err != nil { return MakeString("") }
  return MakeString(filepath.Base(d))
}

func E_Generate_dirname_void (_CL_obj EID) EID {
  return EID{F_Generate_dirname_void().Id(),0}}

// module_has_main(m:module) : boolean -- v4.1.6
//
// True when m defines a method main() with no argument (i.e. the body is
// annotated or takes only the implicit environment). In that case "claire4 -cm"
// behaves like "-cx": the generated binary calls main() and exits instead of
// opening the REPL.
func F_Generate_module_has_main_module (m *ClaireModule) *ClaireBoolean {
  if Core.C_main == nil || Core.C_main.Restrictions == nil { return CFALSE }
  rs := Core.C_main.Restrictions
  for i := 0; i < rs.Length(); i++ {
    r := ToRestriction(rs.At(i))
    if r.Module_I != m { continue }
    if r.Domain.Length() == 0 { return CTRUE }        // main()
    // main(x:void) is the same thing once the environment is dropped
    if r.Domain.Length() == 1 && (r.Domain.At(0) == C_void.Id() || r.Domain.At(0) == C_environment.Id()) { return CTRUE }
  }
  return CFALSE
}

func E_Generate_module_has_main_module (m EID) EID {
  return EID{F_Generate_module_has_main_module(ToModule(OBJ(m))).Id(),0}}

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
  PRINC("        -cm [<module>]: compiles a module -> executable (default: current directory) \n")
  PRINC("        -cc <module>: compiles a module -> target go files      \n")
  PRINC("        -cx <module> : compiles a module & launch main() \n")
  PRINC("         -sf <module> : generates system file associated to a module \n")
  PRINC("         -sx <module> : generates system file that includes main() \n")
  PRINC("        -publish [<dir>][@<version>] : publish the module for use_module() (version from VERSION) \n")
  PRINC("        -force : overwrite an existing published directory \n")
  PRINC("        -install : set up CLAIRE_HOME and put claire4 in your PATH \n")
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
                  { var _Zpublish *ClaireString = MakeString("")
                    { var _Zforce *ClaireBoolean = CFALSE
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
                                  { var g0217 *Language.Call = Language.To_Call(c)
                                    loop_1 = EVAL(g0217.Id())
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
                            // v4.1.6: the module name is optional. Without it we compile the
                            // module declared in ./init.cl that matches the directory name, so
                            // that "cd MyModule && claire4 -cm" works.
                            // the argument is optional, but must not swallow the next
                            // option: "claire4 -cm -publish dir" uses the directory name.
                            if ((l.Length() >= 2) && 
                                (ToString(l.ValuesO()[1]).At(1) != '-')) { 
                              _Zexe = CTRUE
                              _Zcm = ToString(l.ValuesO()[1])
                              l = l.Skip(2)
                              loop_1 = EID{l.Id(),0}
                              } else {
                              _Zexe = CTRUE
                              _Zcm = F_Generate_dirname_void()
                              l = l.Skip(1)
                              loop_1 = EID{l.Id(),0}
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
                            }  else if (ToString(l.ValuesO()[0]).Value == MakeString("-publish").Value) { 
                            // v4.1.6: the destination is optional. Without it, publish()
                            // uses $CLAIRE_LIBS (first entry), then $CLAIRE_HOME/lib.
                            // The destination may carry a version: -publish dir@v1.2.3
                            if ((l.Length() >= 2) && 
                                (ToString(l.ValuesO()[1]).At(1) != '-')) { 
                              _Zpublish = ToString(l.ValuesO()[1])
                              l = l.Skip(2)
                              loop_1 = EID{l.Id(),0}
                              } else {
                              _Zpublish = MakeString("@")
                              l = l.Skip(1)
                              loop_1 = EID{l.Id(),0}
                              } 
                            }  else if (ToString(l.ValuesO()[0]).Value == MakeString("-force").Value) { 
                            _Zforce = CTRUE
                            l = l.Skip(1)
                            loop_1 = EID{l.Id(),0}
                            }  else if (ToString(l.ValuesO()[0]).Value == MakeString("-install").Value) { 
                            // v4.1.6: create ~/.claire4 (or $CLAIRE_HOME) with lib/ and a default init.cl
                            F_Generate_install_void()
                            F_CL_exit(0)
                            loop_1 = EVOID
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
                            // v4.1.6: a module that declares main() is built as an
                            // executable that runs it and exits (no REPL), like -cx.
                            if (F_Generate_module_has_main_module(m) == CTRUE) { 
                              _Zmain = CTRUE
                              } 
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
                            // v4.1.6: -publish <dir> copies the module into a
                            // self-contained directory that use_module() can load.
                            if (!ErrorIn(Result) && !(_Zpublish.Value == MakeString("").Value)) { 
                              Result = F_Generate_publish_module(m,_Zpublish,_Zforce)
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
        Result = F_Generate_system_imports_module2(m,_Zmain)
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
  Result = F_Generate_import_declaration_module2(m,CFALSE)
  if !ErrorIn(Result) {
  // v4.1.6: main_function() always emits "Reader.C_reader" and calls either
  // Reader.F_Reader_simple_main_void() or Generate.F_Generate_complex_main_void().
  // The import must follow, otherwise Go fails with "undefined: Reader".
  // needed_modules() only includes dependencies of the module, so we add Reader
  // (and Generate when the compiler is loaded) explicitly, avoiding duplicates.
  if !(F_Generate_needed_modules_module(m).Contain_ask(Reader.It.Id()) == CTRUE) { 
    PRINC("\t\"Reader\"\n")
    } 
  if ((Optimize.C_compiler.Active_ask == CTRUE) && 
      !(F_Generate_needed_modules_module(m).Contain_ask(It.Id()) == CTRUE)) { 
    PRINC("\t\"Generate\"\n")
    } 
  PRINC("\t\"")
  F_princ_string(F_Generate_capitalize_string(m.Name.String_I()))
  PRINC("\"/**/\n")
  PRINC(")\n")
  Result = EVOID
  }
  return Result} 

// system_imports(m:module,%main:boolean) -- v4.1.6
// %main tells whether the generated binary calls main() (no REPL). In that case
// Reader and Generate are not needed: main_function() does not emit any
// reference to them.
func F_Generate_system_imports_module2 (m *ClaireModule,_Zmain *ClaireBoolean) EID {
  var Result EID
  PRINC("import (\n")
  if Optimize.C_OPT.Profile_ask == CTRUE {
    PRINC("\t\"os\"\n\t\"runtime/pprof\"\n")
  }
  // In %main mode the generated binary has no REPL: Reader and Generate are
  // not referenced, and importing them would be an "imported and not used"
  // error. We therefore build the list ourselves instead of calling
  // import_declaration(), which follows the module's "uses".
  // Reader stays: Load() calls Reader.It = InitModule(...) and Reader.MetaLoad()
  // for any module that declares it in "uses". Only Generate is dropped, since
  // nothing in the generated code refers to it when there is no REPL.
  List := F_Generate_needed_modules_module(m)
  for i := 0; i < List.Length(); i++ {
    x := List.At(i)
    if _Zmain == CTRUE && x == It.Id() { continue }
    PRINC("\t")
    if x == C_Kernel.Id() { PRINC(". ") }
    PRINC("\"")
    F_princ_string(F_Generate_capitalize_string(ToModule(x).Name.String_I()))
    PRINC("\"\n")
  }
  // Reader is needed by the REPL (main_function). Generate is not: the simple
  // REPL (Reader.F_Reader_simple_main_void) does not use it, and importing it
  // would be an unused import. A module that really needs the full compiler
  // declares Generate in its "uses" list.
  if _Zmain != CTRUE && !(List.Contain_ask(Reader.It.Id()) == CTRUE) {
    PRINC("\t\"Reader\"\n")
  }
  PRINC("\t\"")
  F_princ_string(F_Generate_capitalize_string(m.Name.String_I()))
  PRINC("\"/**/\n")
  PRINC(")\n")
  Result = EVOID
  return Result
}

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
        F_Generate_cap_short_symbol(ToSymbol(OBJ(Core.F_CALL(C_name,ARGS(x.ToEID())))))
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
  if (_Zmain != CTRUE) { 
    PRINC("Reader.C_reader.Fromp = ClEnv.Cin;")
    } 
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
  if (_Zmain != CTRUE) { 
    PRINC("Reader.C_reader.Fromp = ClEnv.Cin")
    } 
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
  // v4.1.6: use PRODUCER.source (= compiler.source / <Module>), which is the very
  // directory the code generator writes into. The previous version appended a
  // hard-coded "/go/src" to CLAIRE_HOME, which broke as soon as CLAIRE_HOME
  // already ended with "go/src" (the documented setting).
  { var dir *ClaireString = Optimize.C_compiler.Source
    if (dir.Value == MakeString("").Value) { 
      Result = ToException(Core.C_general_error.Make(MakeString("[213] the output directory is not defined: please set compiler.source\n").Id(),CNIL.Id())).Close()
      } else {
      Result = EID{CFALSE.Id(),0}
      } 
    if !ErrorIn(Result) {
    // create both the workspace (compiler.source) and the module package
    // directory (compiler.source / <Module>), which is where the generator writes
    { var pdir *ClaireString = F_append_string(F_append_string(dir,MakeString("/")),F_Generate_capitalize_string(m.Name.String_I()))
      F_claire_shell(F_append_string(F_append_string(F_append_string(F_append_string(F_append_string(MakeString("mkdir -p "),dir),MakeString(" ")),dir),MakeString("/")),F_Generate_capitalize_string(m.Name.String_I())))
      _ = pdir
      } 
    // v4.1.6: prepare the Go workspace so that the generated imports
    // (". \"Kernel\"", "\"Core\"", ...) resolve. We write one go.mod per base
    // module and a go.work that lists them all. Files are only created when
    // missing, so a user go.mod / go.work is never overwritten.
    F_Generate_prepare_workspace_module(m,dir)
    Result = EVOID
    }
    } 
  return Result} 

// The EID go function for: compile_dir @ module (throw: true) 
func E_Generate_compile_dir_module (m EID) EID { 
  return F_Generate_compile_dir_module(ToModule(OBJ(m)) )} 

// prepare_workspace(m:module,dir:string) : void  -- v4.1.6 (handwritten)
//
// The generated Go code imports the base packages by their bare name
// (". \"Kernel\"", "\"Core\"", "\"Language\"", ...). For `go build` to resolve
// them without GOPATH we create a small Go workspace in <dir>:
//
//   <dir>/go.work                 -- use (./Kernel ./Core ... ./<Module>)
//   <dir>/<Base>/go.mod           -- module <Base>       (one per base module)
//   <dir>/<Module>/go.mod         -- module <Module>     (the module being built)
//
// Existing files are left untouched (a user may provide its own go.mod).
// <dir>/<Base> is only declared when the directory exists: the base modules are
// expected to live next to the generated code (compiler.source = CLAIRE_HOME).
func F_Generate_prepare_workspace_module (m *ClaireModule,dir *ClaireString) {
  var bases []string = []string{"Kernel","Core","Language","Reader","Optimize","Generate"}
  var modname string = F_Generate_capitalize_string(m.Name.String_I()).Value
  var list []string = make([]string,0,8)
  // Where the base packages live. They are read-only inputs: we never copy
  // them, we point the workspace at them.
  //   1. next to the generated code (a full CLAIRE4 tree: src/<Base>)
  //   2. $CLAIRE_HOME/<Base>
  //   3. <home>/lib/claire4/<Base>  (installed by "claire4 -install")
  // v4.1.6: the base packages are looked for, in this order:
  //   - the output directory (a full CLAIRE4 tree)
  //   - $CLAIRE_HOME, then $CLAIRE_HOME/share/claire4
  //   - <dir of the claire4 binary>/../src  (when claire4 was built in src/)
  //   - <dir of the claire4 binary>/../share/claire4 (after "claire4 -install")
  //   - ~/.claire4/share/claire4
  // so that "claire4 -cm MyModule" works without setting CLAIRE_HOME.
  baseRoots := []string{dir.Value, os.Getenv("CLAIRE_HOME")}
  if h := os.Getenv("CLAIRE_HOME"); h != "" {
    baseRoots = append(baseRoots, filepath.Join(h, "share", "claire4"))
  }
  if exe, err := os.Executable(); err == nil {
    if real, err2 := filepath.EvalSymlinks(exe); err2 == nil { exe = real }
    d := filepath.Dir(exe)
    baseRoots = append(baseRoots,
      d,                                 // claire4 built inside src/ (go build -o claire4)
      filepath.Join(d, "..", "src"),     // claire4 installed in <root>/bin
      filepath.Join(d, "..", "share", "claire4"))
  }
  baseRoots = append(baseRoots, filepath.Join(claireHome(), "share", "claire4"))
  for _,b := range bases {
    found := ""
    for _,r := range baseRoots {
      if r == "" { continue }
      cand := filepath.Join(r, b)
      if st,err := os.Stat(cand) ; err == nil && st.IsDir() {
        if _,err := os.Stat(filepath.Join(cand,"go.mod")) ; err == nil { found = cand; break }
      }
    }
    if found != "" {
      // each base package must be a module for the workspace to accept it
      write_if_missing(filepath.Join(found,"go.mod"), "module "+b+"\n\ngo 1.25\n")
      list = append(list, found)
    }
  }
  p := filepath.Join(dir.Value, modname)
  write_if_ours(filepath.Join(p,"go.mod"), "module "+modname+"\n\ngo 1.25\n")
  list = append(list, p)
  var buf string = "go 1.25\n\n// generated by claire4 -cm " + m.Name.String_I().Value + "\n\nuse (\n"
  for _,x := range list { buf = buf + "\t" + x + "\n" }
  buf = buf + ")\n"
  write_if_ours(filepath.Join(dir.Value,"go.work"), buf)
}

// claireHome returns the CLAIRE home directory: $CLAIRE_HOME, else ~/.claire4.
func claireHome() string {
  if h := os.Getenv("CLAIRE_HOME"); h != "" { return h }
  if h, err := os.UserHomeDir(); err == nil { return filepath.Join(h, ".claire4") }
  return ".claire4"
}

// small helper: write a file only if it does not exist yet (v4.1.6)
func write_if_missing(path string, content string) {
  if _,err := os.Stat(path) ; err == nil { return }
  _ = os.WriteFile(path,[]byte(content),0644)
}

// write_if_ours replaces a file that we generated (go.work, go.mod) when its
// content differs, but leaves a file written by the user untouched.
// v4.1.6: a go.work that only lists the user module (because the base packages
// were not found at the time) must be repaired, otherwise "go build" keeps
// failing with "package Kernel is not in std".
//
// The heuristic is the marker "generated by claire4": our files carry it (or
// they are the trivially short "module X / go N" go.mod we write), anything
// else is considered user-written.
func write_if_ours(path string, content string) {
  old, err := os.ReadFile(path)
  if err == nil {
    if string(old) == content { return }                 // already up to date
    txt := string(old)
    ours := strings.Contains(txt, "generated by claire4") ||
            (strings.HasPrefix(txt, "module ") && strings.Count(txt, "\n") <= 3)
    if !ours { return }                                  // user file: leave it alone
  }
  _ = os.WriteFile(path,[]byte(content),0644)
}

// create the go
/* The go function for: compile_exe(%out:string) [status=0] */
func F_Generate_compile_exe_string (_Zout *ClaireString)  { 
  // v4.1.6: build the generated package ./<Module> from the directory that
  // holds the go.work of this module (compiler.source), with an output name
  // that does not clash with the module directory (./<Module>).
  { var outdir *ClaireString
    out := F_Generate_capitalize_string(_Zout).Value
    if ((ToGenerateGoProducer(Optimize.C_PRODUCER.Value).Output).Id() != CNULL) { 
      outdir = F_append_string(F_append_string(MakeString("-o "),ToGenerateGoProducer(Optimize.C_PRODUCER.Value).Output),MakeString(" "))
      } else {
      // default: ./<module> in the current directory (not inside the package dir)
      // <module>/bin/<module> -- a dedicated directory, so that the executable
      // never collides with the generated package directory ./<Module>.
      exe := filepath.Join("bin", strings.ToLower(out))
      if abs,err := filepath.Abs(exe) ; err == nil { exe = abs }
      _ = os.MkdirAll(filepath.Dir(exe), 0755)
      outdir = F_append_string(F_append_string(MakeString("-o "),MakeString(exe)),MakeString(" "))
      } 
    // The system file (<out>.go, package main) is what we build: the module
    // package ./<Module> only holds the generated API.
    { var pkg *ClaireString = F_append_string(F_append_string(MakeString("./"),_Zout),MakeString(".go"))
      { var s *ClaireString = F_append_string(F_append_string(F_append_string(F_append_string(MakeString("cd "),Optimize.C_compiler.Source),MakeString(" && go build ")),outdir),pkg)
        F_claire_shell(s)
        } 
      } 
    } 
  } 

// The EID go function for: compile_exe @ string (throw: false) 
func E_Generate_compile_exe_string (_Zout EID) EID { 
  F_Generate_compile_exe_string(ToString(OBJ(_Zout)) )
  return EVOID}  
// *******************************************************************
// *  Part 3b : publish a module  (action C of TODO-CLAIRE4)  v4.1.6  *
// *******************************************************************
// publish(m:module,%dir:string) : void
//
// Copies a compiled module into a self-contained, versioned directory that
// use_module() (Reader/file.cl) knows how to load:
//
//   <dir>/<name>/<version>/
//        init.cl        -- module declaration, copied from m.source/..
//        source/        -- the CLAIRE sources (for interpretation / recompilation)
//        build/         -- the generated Go package (<Module>/*.go) + go.mod
//        META.json      -- version, date, platform, base path, file list
//
// %dir defaults to the first entry of $CLAIRE_LIBS, then $CLAIRE_HOME/lib.
// An existing target directory is an error unless force, so that publishing
// twice by accident never corrupts a working install.
func F_Generate_publish_module (m *ClaireModule,_Zdir *ClaireString,_Zforce *ClaireBoolean) EID {
  name := m.Name.String_I().Value
  dest := _Zdir.Value
  // 1. version, by decreasing priority:
  //    a. -publish <dir>@<version>   (explicit, useful in CI)
  //    b. the file VERSION next to init.cl (see readVersionFile)
  //    c. "v0.0.0"
  // CLAIRE4 has no "version" slot on module (a free slot is rejected:
  // "[106] the object does not understand foo"), hence the file convention.
  ver := ""
  if k := strings.LastIndex(_Zdir.Value, "@"); k >= 0 {
    if v := _Zdir.Value[k+1:]; v != "" { ver = v }
    dest = _Zdir.Value[:k]
  }
  if ver == "" { ver = readVersionFile(m) }
  if ver == "" { ver = "v0.0.0" }
  // 2. destination
  if dest == "" {
    // note: dest was set from _Zdir, which is "@" when -publish had no argument
    if libs := os.Getenv("CLAIRE_LIBS"); libs != "" {
      dest = strings.Split(libs, string(os.PathListSeparator))[0]
    } else {
      // $CLAIRE_HOME/lib, or ~/.claire4/lib: the same place use_module() looks
      // for modules when no environment variable is set.
      hom := os.Getenv("CLAIRE_HOME")
      if hom == "" {
        if h, err := os.UserHomeDir(); err == nil { hom = h + "/.claire4" }
      }
      dest = hom + "/lib"
    }
  }
  target := dest + "/" + name + "/" + ver
  if _, err := os.Stat(target); err == nil && _Zforce != CTRUE {
    return ToException(Core.C_general_error.Make(
      F_append_string(MakeString("[214] the directory already exists (use -force to overwrite): "), MakeString(target)).Id(),
      CNIL.Id())).Close()
  }
  fmt.Printf("==== Publish %s %s -> %s ====\n", name, ver, target)
  // 3. copy init.cl
  srcDir := m.Source.Value
  initSrc := Reader.F__7_string(MakeString(srcDir), MakeString("../init.cl"))
  if F_isfile_ask_string(initSrc) != CTRUE {
    initSrc = Reader.F__7_string(MakeString(srcDir), MakeString("init.cl"))
  }
  if err := copyFile(initSrc.Value, target+"/init.cl"); err != nil {
    return ToException(Core.C_general_error.Make(
      F_append_string(MakeString("[214] cannot copy init.cl: "), MakeString(err.Error())).Id(),
      CNIL.Id())).Close()
  }
  // 4. copy the sources (m.source + the files listed in made_of)
  if err := copyTree(srcDir, target+"/"+filepath.Base(srcDir)); err != nil {
    return ToException(Core.C_general_error.Make(
      F_append_string(MakeString("[214] cannot copy sources: "), MakeString(err.Error())).Id(),
      CNIL.Id())).Close()
  }
  // 5. copy the generated package
  pkgDir := F_Generate_capitalize_string(m.Name.String_I()).Value
  if err := copyTree(Optimize.C_compiler.Source.Value+"/"+pkgDir, target+"/build/"+pkgDir); err != nil {
    return ToException(Core.C_general_error.Make(
      F_append_string(MakeString("[214] cannot copy the generated code: "), MakeString(err.Error())).Id(),
      CNIL.Id())).Close()
  }
  // 6. manifest
  meta := "{\n" +
    "  \"name\": \"" + name + "\",\n" +
    "  \"version\": \"" + ver + "\",\n" +
    "  \"date\": \"" + time.Now().Format(time.RFC3339) + "\",\n" +
    "  \"os\": \"" + runtime.GOOS + "\",\n" +
    "  \"arch\": \"" + runtime.GOARCH + "\",\n" +
    "  \"claire_version\": \"" + Optimize.C_compiler.Version.Prt() + "\",\n" +
    "  \"source\": \"" + filepath.Base(srcDir) + "\",\n" +
    "  \"package\": \"" + pkgDir + "\"\n" +
    "}\n"
  if err := os.WriteFile(target+"/META.json", []byte(meta), 0644); err != nil {
    return ToException(Core.C_general_error.Make(
      F_append_string(MakeString("[214] cannot write META.json: "), MakeString(err.Error())).Id(),
      CNIL.Id())).Close()
  }
  fmt.Printf("==== published: use_module(\"%s/%s\") ====\n", name, ver)
  return EVOID
}

func E_Generate_publish_module (m EID, dir EID, force EID) EID {
  return F_Generate_publish_module(ToModule(OBJ(m)), ToString(OBJ(dir)), ToBoolean(OBJ(force)))}

// install() : void  -- v4.1.6 (action D)
//
// One-time setup, the equivalent of "configure && make install" for the C++
// CLAIRE3. It creates the CLAIRE home ($CLAIRE_HOME if set, else ~/.claire4)
// with the layout expected by compiler.source and use_module():
//
//     <home>/            -- compiler.source: where "claire4 -cm" writes code
//     <home>/lib/        -- -publish / use_module
//     <home>/init.cl     -- the init file loaded at start-up
//
// The interpreter itself is a single Go binary: there is nothing else to
// install (Kernel, Core, Language, Reader, Optimize and Generate are compiled
// into it).
func F_Generate_install_void () {
  home := os.Getenv("CLAIRE_HOME")
  if home == "" {
    if h, err := os.UserHomeDir(); err == nil {
      home = filepath.Join(h, ".claire4")
    } else {
      home = ".claire4"
    }
  }
  fmt.Printf("==== CLAIRE4 install: %s ====\n", home)
  for _, d := range []string{home, filepath.Join(home, "lib")} {
    if err := os.MkdirAll(d, 0755); err != nil {
      fmt.Printf("---- cannot create %s: %v\n", d, err)
      return
    }
    fmt.Printf("   ok  %s\n", d)
  }

  // init.cl (never overwritten)
  initPath := filepath.Join(home, "init.cl")
  if _, err := os.Stat(initPath); err == nil {
    fmt.Printf("   keep existing %s\n", initPath)
  } else {
    content := "// CLAIRE4 init file -- created by claire4 -install\n" +
      "// " + time.Now().Format(time.RFC3339) + "\n" +
      "//\n" +
      "// Modules published with 'claire4 -cm <m> -publish <dir>' or listed in\n" +
      "// $CLAIRE_LIBS can be loaded with:  (use_module(\"mymod/v1.0.0\"))\n" +
      "\n(printf(\"Hello CLAIRE4\\n\"))\n"
    if err := os.WriteFile(initPath, []byte(content), 0644); err != nil {
      fmt.Printf("---- cannot write %s: %v\n", initPath, err)
      return
    }
    fmt.Printf("   new %s\n", initPath)
  }

  // copy the base packages (Kernel, Core, ...) so that "claire4 -cm" can
  // compile a module anywhere, without a CLAIRE4 source tree nearby.
  arch, _ := os.Executable()
  if real, err := filepath.EvalSymlinks(arch); err == nil { arch = real }
  bases := []string{"Kernel","Core","Language","Reader","Optimize","Generate"}
  srcRoots := []string{filepath.Dir(arch), filepath.Join(filepath.Dir(arch), "..", "src")}
  srcRoot := ""
  for _, r := range srcRoots {
    if st, err := os.Stat(filepath.Join(r, "Kernel", "ClKernel.go")); err == nil && !st.IsDir() {
      srcRoot = r
      break
    }
  }
  shareDir := filepath.Join(home, "share", "claire4")
  if srcRoot == "" {
    fmt.Printf("---- base packages not found next to %s: 'claire4 -cm' will\n", arch)
    fmt.Printf("     need CLAIRE_HOME to point to a CLAIRE4 source tree\n")
  } else {
    for _, b := range bases {
      if err := copyTree(filepath.Join(srcRoot, b), filepath.Join(shareDir, b)); err != nil {
        fmt.Printf("---- cannot copy %s: %v\n", b, err)
        break
      }
    }
    fmt.Printf("   ok  %s (base packages)\n", shareDir)
  }

  // install the binary in a directory reachable from the PATH
  exe, err := os.Executable()
  if err != nil {
    fmt.Printf("---- cannot locate the claire4 executable: %v\n", err)
    return
  }
  if exe, err = filepath.EvalSymlinks(exe); err != nil {
    exe = ""
  }
  dest, why := installBinDir(home)
  if dest == "" && exe != "" {
    // no directory from the PATH is usable: fall back on <home>/bin
    dest = filepath.Join(home, "bin")
    why = "not in the PATH -- add it"
  }
  if dest != "" && exe != "" {
    target := filepath.Join(dest, filepath.Base(exe))
    if samePath(exe, target) {
      fmt.Printf("   ok  binary already there: %s\n", target)
    } else if err := copyExec(exe, target); err != nil {
      fmt.Printf("---- cannot install the binary in %s: %v\n", dest, err)
      fmt.Printf("     (try: sudo cp %s %s)\n", exe, target)
    } else {
      fmt.Printf("   new %s%s\n", target, why)
    }
  }

  fmt.Printf("\nAdd this to your shell profile:\n")
  fmt.Printf("   export CLAIRE_HOME=%s\n", home)
  if dest != "" && !inPath(dest) {
    fmt.Printf("   export PATH=$PATH:%s\n", dest)
  }
  fmt.Printf("\nDone. Try:  claire4 -h\n")
}

// installBinDir picks the directory where the claire4 binary should be copied,
// following the usual conventions, in this order (v4.1.6):
//
//   1. $CLAIRE_BIN                                    -- explicit override
//   2. $GOBIN or $(go env GOPATH)/bin, if in PATH     -- Go convention
//   3. ~/.local/bin, if in PATH                       -- XDG, per-user
//   4. /usr/local/bin, only if CLAIRE_BIN_SYSTEM=1    -- FHS, opt-in
//   5. ""                                             -- caller uses <home>/bin
//
// /usr/local/bin is never selected automatically: writing outside the user's
// home during a simple "claire4 -install" would be surprising, and it may
// require administrator rights on a shared machine. Set CLAIRE_BIN_SYSTEM=1
// to opt in explicitly.
//
// The returned string explains the choice (shown next to the copied file).
func installBinDir(home string) (string, string) {
  if d := os.Getenv("CLAIRE_BIN"); d != "" {
    if canWriteDir(d) { return d, "  (from $CLAIRE_BIN)" }
  }
  if d := os.Getenv("GOBIN"); d != "" && inPath(d) && canWriteDir(d) {
    return d, "  (from $GOBIN)"
  }
  if gp := os.Getenv("GOPATH"); gp != "" {
    for _, d := range filepath.SplitList(gp) {
      b := filepath.Join(d, "bin")
      if inPath(b) && canWriteDir(b) { return b, "  (Go path, in $PATH)" }
    }
  }
  h, _ := os.UserHomeDir()
  if h != "" {
    d := filepath.Join(h, ".local", "bin")
    if inPath(d) && canWriteDir(d) { return d, "  (~/.local/bin, in $PATH)" }
  }
  if os.Getenv("CLAIRE_BIN_SYSTEM") == "1" && canWriteDir("/usr/local/bin") {
    return "/usr/local/bin", "  (/usr/local/bin, from $CLAIRE_BIN_SYSTEM)"
  }
  if h != "" {
    d := filepath.Join(h, ".local", "bin")
    if canWriteDir(d) { return d, "  (~/.local/bin, add it to your $PATH)" }
  }
  return "", ""
}

// inPath tells whether dir belongs to $PATH.
func inPath(dir string) bool {
  want := filepath.Clean(dir)
  for _, p := range filepath.SplitList(os.Getenv("PATH")) {
    if filepath.Clean(p) == want { return true }
  }
  return false
}

// canWriteDir tells whether we may create files in dir (creating it if needed).
func canWriteDir(dir string) bool {
  if err := os.MkdirAll(dir, 0755); err != nil { return false }
  f, err := os.CreateTemp(dir, ".claire4-probe-*")
  if err != nil { return false }
  name := f.Name()
  f.Close()
  os.Remove(name)
  return true
}

// samePath compares two paths, resolving symlinks when possible.
func samePath(a string, b string) bool {
  ra, erra := filepath.EvalSymlinks(a)
  rb, errb := filepath.EvalSymlinks(b)
  if erra == nil && errb == nil { return ra == rb }
  return filepath.Clean(a) == filepath.Clean(b)
}

// copyExec copies a binary, keeping the executable bit.
func copyExec(src string, dst string) error {
  in, err := os.Open(src)
  if err != nil { return err }
  defer in.Close()
  if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil { return err }
  tmp := dst + ".tmp"
  out, err := os.OpenFile(tmp, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0755)
  if err != nil { return err }
  if _, err := io.Copy(out, in); err != nil { out.Close(); os.Remove(tmp); return err }
  if err := out.Close(); err != nil { os.Remove(tmp); return err }
  return os.Rename(tmp, dst)
}

// readVersionFile looks for a VERSION file belonging to the module m and
// returns its first non-empty line, or "". It is looked for next to init.cl:
//   <module source>/../VERSION, then <module source>/VERSION, then ./VERSION
// Trailing and leading spaces are trimmed, and a leading "v" is kept as is.
func readVersionFile(m *ClaireModule) string {
  var cands []string
  if src := m.Source.Value; src != "" {
    cands = append(cands,
      filepath.Join(src, "..", "VERSION"),
      filepath.Join(src, "VERSION"))
  }
  cands = append(cands, "VERSION")
  for _, p := range cands {
    b, err := os.ReadFile(p)
    if err != nil { continue }
    for _, l := range strings.Split(string(b), "\n") {
      if t := strings.TrimSpace(l); t != "" && !strings.HasPrefix(t, "#") {
        return t
      }
    }
  }
  return ""
}

// copyFile copies a single file, creating the destination directory if needed.
func copyFile(src string, dst string) error {
  in, err := os.Open(src)
  if err != nil { return err }
  defer in.Close()
  if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil { return err }
  out, err := os.Create(dst)
  if err != nil { return err }
  defer out.Close()
  _, err = io.Copy(out, in)
  return err
}

// copyTree copies a directory recursively (used by publish).
func copyTree(src string, dst string) error {
  st, err := os.Stat(src)
  if err != nil { return err }
  if !st.IsDir() { return copyFile(src, dst) }
  if err := os.MkdirAll(dst, 0755); err != nil { return err }
  des, err := os.ReadDir(src)
  if err != nil { return err }
  for _, d := range des {
    if strings.HasPrefix(d.Name(), ".") { continue }   // .DS_Store, .git ...
    s := filepath.Join(src, d.Name())
    t := filepath.Join(dst, d.Name())
    if d.IsDir() {
      if err := copyTree(s, t); err != nil { return err }
    } else if err := copyFile(s, t); err != nil {
      return err
    }
  }
  return nil
}

 // *******************************************************************

// *******************************************************************
// *       Part 4: Profiler code                                     *
// *******************************************************************
// ---------------------------------------------------------------











