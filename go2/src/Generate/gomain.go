/***** CLAIRE Compilation of file /Users/ycaseau/Dropbox/src/clairev4.0/src/compile/gomain.cl 
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
func import_g0781() { 
    _ = Core.It
    _ = Language.It
    _ = Reader.It
    _ = Optimize.It
    } 
  
  
//+-------------------------------------------------------------+
//| CLAIRE                                                      |
//| gomain.cl                                                   |
//| Copyright (C) 1994 - 2021 Yves Caseau. All Rights Reserved  |
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
//                       -cc : (slave of -cf) -> generage go files
// ----------------------------------------------------------------------
// -------------------------------------------------------------------
// Contents
//      Part 1: definition of the main function
//      Part 2: Generating a system file
//      Part 3: create an executable : generate a command line and calls it
//      Part 4: Profiler code  
// -------------------------------------------------------------------
// dumb utility
/* {1} OPT.The go function for: external!(m:module) [] */
func F_Generate_external_I_module (m *ClaireModule ) *ClaireString  { 
    // use function body compiling 
if ((m.External).Id() != CNULL) /* body If:2 */{ 
      return  m.External
      } else {
      return  m.Name.String_I()
      /* body If-2 */} 
    } 
  
// The EID go function for: external! @ module (throw: false) 
func E_Generate_external_I_module (m EID) EID { 
    return EID{/*(sm for external! @ module= string)*/ F_Generate_external_I_module(ToModule(OBJ(m)) ).Id(),0}} 
  
/* {1} OPT.The go function for: string2module(s:string) [] */
func F_Generate_string2module_string (s *ClaireString ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var m *ClaireAny   = F_value_string(s)
      /* noccur = 2 */
      if (m.Isa.IsIn(C_module) == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var g0782 *ClaireModule   = ToModule(m)
          /* noccur = 1 */
          Result = EID{g0782.Id(),0}
          /* Let-4 */} 
        } else {
        Result = ToException(Core.C_general_error.Make(MakeString("~A is not a module").Id(),MakeConstantList((s).Id()).Id())).Close()
        /* If-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: string2module @ string (throw: true) 
func E_Generate_string2module_string (s EID) EID { 
    return /*(sm for string2module @ string= EID)*/ F_Generate_string2module_string(ToString(OBJ(s)) )} 
  
// *******************************************************************
// *       Part 1: definition of the main function                   *
// *******************************************************************
// help file
/* {1} OPT.The go function for: printHelp(_CL_obj:void) [] */
func F_Generate_printHelp_void ()  { 
    // procedure body with s =  
PRINC("------------- CLAIRE: The Art of Elegant Programming -----------\n\n")
    Core.F_about_void()
    PRINC("\noptions -s <int> : set memory allocation size  \n")
    PRINC("        -f <filename>  : load <filename>             \n")
    PRINC("        -n : do not load the init file               \n")
    PRINC("        -m <module> : load <module>                  \n")
    PRINC("        -v <int> : sets the verbosity level       \n")
    PRINC("        -S <flag> : sets the global variable <flag> to true  \n")
    PRINC("        -o <name> : sets the name of the executable  \n")
    PRINC("        -od <name> : sets the output directory  \n")
    PRINC("        -p : profiling mode                          \n")
    PRINC("        -D : debug mode                              \n")
    PRINC("        -safe : safe mode                            \n")
    PRINC("        -O : optimizing mode                         \n")
    PRINC("        -cm <module>: compiles a module -> executable        \n")
    PRINC("        -cc <module>: compiles a module -> target go files      \n")
    PRINC("        -cx <module> : generates system file associated to a module \n")
    F_CL_exit(0)
    } 
  
// The EID go function for: printHelp @ void (throw: false) 
func E_Generate_printHelp_void (_CL_obj EID) EID { 
    /*(sm for printHelp @ void= void)*/ F_Generate_printHelp_void( )
    return EVOID} 
  
// Claire's main - the rich version that starts the compiler
// a simpler version exists in inspect.cl (simple_main())
// -s is ignored because it is trapped earlier (see the file generator)
/* {1} OPT.The go function for: complex_main(_CL_obj:void) [] */
func F_Generate_complex_main_void () EID { 
    var Result EID 
    /* Let:2 */{ 
      var _Zcm *ClaireString   = MakeString("")
      /* noccur = 6 */
      /* Let:3 */{ 
        var _Zcx *ClaireString   = MakeString("")
        /* noccur = 5 */
        /* Let:4 */{ 
          var _Zout *ClaireString   = MakeString("")
          /* noccur = 8 */
          /* Let:5 */{ 
            var dblevel int  = 1
            /* noccur = 7 */
            /* Let:6 */{ 
              var vlevel int  = 2
              /* noccur = 3 */
              /* Let:7 */{ 
                var _Zinit_ask *ClaireBoolean   = CTRUE
                /* noccur = 4 */
                /* Let:8 */{ 
                  var _Zexe *ClaireBoolean   = CFALSE
                  /* noccur = 2 */
                  /* Let:9 */{ 
                    var l *ClaireList   = ClEnv.Params.Copy()
                    /* noccur = 71 */
                    h_index := ClEnv.Index /* Handle */
                    h_base := ClEnv.Base
                    Reader.C__starfs_star.Value = MakeString("/").Id()
                    Optimize.C_compiler.Env = MakeString("MacOS")
                    Result= EID{CFALSE.Id(),0}
                    for (l.Length() != 0) /* while:10 */{ 
                      var void_try11 EID 
                      _ = void_try11
                      if ((ToString(l.ValuesO()[1-1]).Value == MakeString("-help").Value) || 
                          (ToString(l.ValuesO()[1-1]).Value == MakeString("?").Value)) /* If:11 */{ 
                        F_Generate_printHelp_void()
                        void_try11 = EVOID
                        /* If!11 */}  else if (ToString(l.ValuesO()[1-1]).Value == MakeString("-s").Value) /* If:11 */{ 
                        if (l.Length() >= 2) /* If:12 */{ 
                          l = l.Skip(2)
                          void_try11 = EID{l.Id(),0}
                          } else {
                          void_try11 = ToException(Core.C_general_error.Make(MakeString("option: -s <s1> <s2>").Id(),CNIL.Id())).Close()
                          /* If-12 */} 
                        /* If!11 */}  else if (ToString(l.ValuesO()[1-1]).Value == MakeString("-f").Value) /* If:11 */{ 
                        if (l.Length() >= 2) /* If:12 */{ 
                          void_try11 = Reader.F_load_string(ToString(l.ValuesO()[2-1]))
                          /* ERROR PROTECTION INSERTED (void_try11-void_try11) */
                          if ErrorIn(void_try11) {Result = void_try11
                          break
                          } else {
                          l = l.Skip(2)
                          void_try11 = EID{l.Id(),0}
                          }
                          } else {
                          void_try11 = ToException(Core.C_general_error.Make(MakeString("option: -f <filename>").Id(),CNIL.Id())).Close()
                          /* If-12 */} 
                        /* If!11 */}  else if (ToString(l.ValuesO()[1-1]).Value == MakeString("-m").Value) /* If:11 */{ 
                        if (l.Length() >= 2) /* If:12 */{ 
                          if (_Zinit_ask == CTRUE) /* If:13 */{ 
                            void_try11 = Reader.F_load_string(MakeString("init"))
                            /* ERROR PROTECTION INSERTED (void_try11-void_try11) */
                            if ErrorIn(void_try11) {Result = void_try11
                            break
                            } else {
                            _Zinit_ask = CFALSE
                            void_try11 = EID{_Zinit_ask.Id(),0}
                            }
                            } else {
                            void_try11 = EID{CFALSE.Id(),0}
                            /* If-13 */} 
                          /* ERROR PROTECTION INSERTED (void_try11-void_try11) */
                          if ErrorIn(void_try11) {Result = void_try11
                          break
                          } else {
                          /* Let:13 */{ 
                            var m *ClaireModule  
                            /* noccur = 3 */
                            var m_try078814 EID 
                            m_try078814 = F_Generate_string2module_string(ToString(l.ValuesO()[2-1]))
                            /* ERROR PROTECTION INSERTED (m-void_try11) */
                            if ErrorIn(m_try078814) {void_try11 = m_try078814
                            } else {
                            m = ToModule(OBJ(m_try078814))
                            void_try11 = Reader.F_load_module(m)
                            /* ERROR PROTECTION INSERTED (void_try11-void_try11) */
                            if ErrorIn(void_try11) {Result = void_try11
                            break
                            } else {
                            m.Begin()
                            l = l.Skip(2)
                            var v_gassign14 *ClaireAny  
                            v_gassign14 = ToList(Optimize.C_claire_modules.Value).AddFast(m.Id()).Id()
                            Optimize.C_claire_modules.Value = v_gassign14
                            void_try11 = v_gassign14.ToEID()
                            }
                            }
                            /* Let-13 */} 
                          /* ERROR PROTECTION INSERTED (void_try11-void_try11) */
                          if ErrorIn(void_try11) {Result = void_try11
                          break
                          } else {
                          }}
                          } else {
                          void_try11 = ToException(Core.C_general_error.Make(MakeString("option: -m <module>").Id(),CNIL.Id())).Close()
                          /* If-12 */} 
                        /* If!11 */}  else if (ToString(l.ValuesO()[1-1]).Value == MakeString("-v").Value) /* If:11 */{ 
                        if (l.Length() >= 2) /* If:12 */{ 
                          vlevel = (vlevel+F_integer_I_string(ToString(l.ValuesO()[2-1])))
                          l = l.Skip(2)
                          void_try11 = EID{l.Id(),0}
                          } else {
                          void_try11 = ToException(Core.C_general_error.Make(MakeString("option: -v <integer>").Id(),CNIL.Id())).Close()
                          /* If-12 */} 
                        /* If!11 */}  else if (ToString(l.ValuesO()[1-1]).Value == MakeString("-S").Value) /* If:11 */{ 
                        if (l.Length() >= 2) /* If:12 */{ 
                          /* update:13 */{ 
                            var va_arg1 *Core.GlobalVariable  
                            var va_arg2 *ClaireAny  
                            var va_arg1_try078914 EID 
                            va_arg1_try078914 = Core.F_new_class2(Core.C_global_variable,Core.F_symbol_I_string2(ToString(l.ValuesO()[2-1])))
                            /* ERROR PROTECTION INSERTED (va_arg1-void_try11) */
                            if ErrorIn(va_arg1_try078914) {void_try11 = va_arg1_try078914
                            } else {
                            va_arg1 = Core.ToGlobalVariable(OBJ(va_arg1_try078914))
                            va_arg2 = CTRUE.Id()
                            /* ---------- now we compile update value(va_arg1) := va_arg2 ------- */
                            va_arg1.Value = va_arg2
                            void_try11 = va_arg2.ToEID()
                            }
                            /* update-13 */} 
                          /* ERROR PROTECTION INSERTED (void_try11-void_try11) */
                          if ErrorIn(void_try11) {Result = void_try11
                          break
                          } else {
                          l = l.Skip(2)
                          void_try11 = EID{l.Id(),0}
                          }
                          } else {
                          void_try11 = ToException(Core.C_general_error.Make(MakeString("option: -S <FLAG>").Id(),CNIL.Id())).Close()
                          /* If-12 */} 
                        /* If!11 */}  else if (ToString(l.ValuesO()[1-1]).Value == MakeString("-od").Value) /* If:11 */{ 
                        if (l.Length() >= 2) /* If:12 */{ 
                          Optimize.C_compiler.Source = ToString(l.ValuesO()[2-1])
                          l = l.Skip(2)
                          void_try11 = EID{l.Id(),0}
                          } else {
                          void_try11 = ToException(Core.C_general_error.Make(MakeString("option: -od <directory>").Id(),CNIL.Id())).Close()
                          /* If-12 */} 
                        /* If!11 */}  else if (ToString(l.ValuesO()[1-1]).Value == MakeString("-o").Value) /* If:11 */{ 
                        if (l.Length() >= 2) /* If:12 */{ 
                          _Zout = ToString(l.ValuesO()[2-1])
                          l = l.Skip(2)
                          void_try11 = EID{l.Id(),0}
                          } else {
                          void_try11 = ToException(Core.C_general_error.Make(MakeString("option: -o <name>").Id(),CNIL.Id())).Close()
                          /* If-12 */} 
                        /* If!11 */}  else if (ToString(l.ValuesO()[1-1]).Value == MakeString("-p").Value) /* If:11 */{ 
                        Optimize.C_OPT.Profile_ask = CTRUE
                        if (dblevel <= 1) /* If:12 */{ 
                          dblevel = 1
                          } else {
                          dblevel = dblevel
                          /* If-12 */} 
                        l = l.Skip(1)
                        void_try11 = EID{l.Id(),0}
                        /* If!11 */}  else if (ToString(l.ValuesO()[1-1]).Value == MakeString("-D").Value) /* If:11 */{ 
                        dblevel = 0
                        l = l.Skip(1)
                        void_try11 = EID{l.Id(),0}
                        /* If!11 */}  else if (ToString(l.ValuesO()[1-1]).Value == MakeString("-safe").Value) /* If:11 */{ 
                        /* update:12 */{ 
                          var va_arg1 *Optimize.OptimizeMetaCompiler  
                          var va_arg2 int 
                          va_arg1 = Optimize.C_compiler
                          if (dblevel == 0) /* If:13 */{ 
                            va_arg2 = 0
                            } else {
                            va_arg2 = 1
                            /* If-13 */} 
                          /* ---------- now we compile update safety(va_arg1) := va_arg2 ------- */
                          va_arg1.Safety = va_arg2
                          /* update-12 */} 
                        Optimize.C_claire_lib.Value = Optimize.C_compiler.LibrariesDir.At(2-1)
                        Optimize.C_claire_options.Value = Optimize.C_compiler.Options.At(2-1)
                        l = l.Skip(1)
                        void_try11 = EID{l.Id(),0}
                        /* If!11 */}  else if (ToString(l.ValuesO()[1-1]).Value == MakeString("-O").Value) /* If:11 */{ 
                        Optimize.C_compiler.Optimize_ask = CTRUE
                        dblevel = 2
                        l = l.Skip(1)
                        void_try11 = EID{l.Id(),0}
                        /* If!11 */}  else if (ToString(l.ValuesO()[1-1]).Value == MakeString("-cc").Value) /* If:11 */{ 
                        if (l.Length() >= 2) /* If:12 */{ 
                          _Zcm = ToString(l.ValuesO()[2-1])
                          l = l.Skip(2)
                          void_try11 = EID{l.Id(),0}
                          } else {
                          void_try11 = ToException(Core.C_general_error.Make(MakeString("option: -cc <module>").Id(),CNIL.Id())).Close()
                          /* If-12 */} 
                        /* If!11 */}  else if (ToString(l.ValuesO()[1-1]).Value == MakeString("-cm").Value) /* If:11 */{ 
                        if (l.Length() >= 2) /* If:12 */{ 
                          _Zexe = CTRUE
                          _Zcm = ToString(l.ValuesO()[2-1])
                          l = l.Skip(2)
                          void_try11 = EID{l.Id(),0}
                          } else {
                          void_try11 = ToException(Core.C_general_error.Make(MakeString("option: -cm <module>").Id(),CNIL.Id())).Close()
                          /* If-12 */} 
                        /* If!11 */}  else if (ToString(l.ValuesO()[1-1]).Value == MakeString("-cx").Value) /* If:11 */{ 
                        if (l.Length() >= 2) /* If:12 */{ 
                          _Zcx = ToString(l.ValuesO()[2-1])
                          l = l.Skip(2)
                          void_try11 = EID{l.Id(),0}
                          } else {
                          void_try11 = ToException(Core.C_general_error.Make(MakeString("option: -cx <filename>").Id(),CNIL.Id())).Close()
                          /* If-12 */} 
                        /* If!11 */}  else if (ToString(l.ValuesO()[1-1]).Value == MakeString("-n").Value) /* If:11 */{ 
                        _Zinit_ask = CFALSE
                        l = l.Skip(1)
                        void_try11 = EID{l.Id(),0}
                        } else {
                        if (ToString(l.ValuesO()[1-1]).At(1) == '-') /* If:12 */{ 
                          void_try11 = Core.F_print_any(l.ValuesO()[1-1])
                          /* ERROR PROTECTION INSERTED (void_try11-void_try11) */
                          if ErrorIn(void_try11) {Result = void_try11
                          break
                          } else {
                          PRINC(" is an unvalid option\n")
                          void_try11 = EVOID
                          }
                          {
                          F_Generate_printHelp_void()
                          void_try11 = EVOID
                          }
                          } else {
                          void_try11 = EID{CFALSE.Id(),0}
                          /* If-12 */} 
                        /* ERROR PROTECTION INSERTED (void_try11-void_try11) */
                        if ErrorIn(void_try11) {Result = void_try11
                        break
                        } else {
                        l = ToType(C_string.Id()).EmptyList()
                        void_try11 = EID{l.Id(),0}
                        }
                        /* If-11 */} 
                      /* ERROR PROTECTION INSERTED (void_try11-Result) */
                      if ErrorIn(void_try11) {Result = void_try11
                      Result = void_try11
                      break
                      } else {
                      /* while-10 */} 
                    }
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    if (_Zout.Value == MakeString("").Value) /* If:10 */{ 
                      if (_Zcm.Value != MakeString("").Value) /* If:11 */{ 
                        _Zout = _Zcm
                        /* If!11 */}  else if (_Zcx.Value != MakeString("").Value) /* If:11 */{ 
                        _Zout = _Zcx
                        /* If-11 */} 
                      /* If-10 */} 
                    if (_Zinit_ask == CTRUE) /* If:10 */{ 
                      Result = Reader.F_load_string(MakeString("init"))
                      } else {
                      Result = EID{CFALSE.Id(),0}
                      /* If-10 */} 
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    ClEnv.Verbose = vlevel
                    if (_Zcx.Value != MakeString("").Value) /* If:10 */{ 
                      Optimize.C_compiler.Active_ask = CTRUE
                      Result = Core.F_CALL(Reader.C_load,ARGS(F_value_string(MakeString("Compile")).ToEID()))
                      /* ERROR PROTECTION INSERTED (Result-Result) */
                      if !ErrorIn(Result) {
                      /* Let:11 */{ 
                        var g0790UU *ClaireModule  
                        /* noccur = 1 */
                        var g0790UU_try079112 EID 
                        g0790UU_try079112 = F_Generate_string2module_string(_Zcx)
                        /* ERROR PROTECTION INSERTED (g0790UU-Result) */
                        if ErrorIn(g0790UU_try079112) {Result = g0790UU_try079112
                        } else {
                        g0790UU = ToModule(OBJ(g0790UU_try079112))
                        Result = F_system_file_module(g0790UU,_Zout)
                        }
                        /* Let-11 */} 
                      /* ERROR PROTECTION INSERTED (Result-Result) */
                      if !ErrorIn(Result) {
                      F_CL_exit(0)
                      Result = EVOID
                      }}
                      /* If!10 */}  else if (_Zcm.Value != MakeString("").Value) /* If:10 */{ 
                      /* Let:11 */{ 
                        var m *ClaireModule  
                        /* noccur = 10 */
                        var m_try079212 EID 
                        m_try079212 = F_Generate_string2module_string(_Zcm)
                        /* ERROR PROTECTION INSERTED (m-Result) */
                        if ErrorIn(m_try079212) {Result = m_try079212
                        } else {
                        m = ToModule(OBJ(m_try079212))
                        Optimize.C_compiler.Active_ask = CTRUE
                        if (Equal(m.Uses.Id(),MakeConstantList(ToList(Optimize.C_claire_modules.Value).At(2-1)).Id()) == CTRUE) /* If:12 */{ 
                          Optimize.C_claire_modules.Value = ToList(Optimize.C_claire_modules.Value).Shrink(2).Id()
                          Core.F_tformat_string(MakeString("=== Light Module ~S:~S -> use ~S=== "),0,MakeConstantList(m.Id(),m.Uses.Id(),Optimize.C_claire_modules.Value))
                          /* If-12 */} 
                        Optimize.C_claire_modules.Value = ToList(Optimize.C_claire_modules.Value).AddFast(m.Id()).Id()
                        Result = Core.F_CALL(Reader.C_load,ARGS(F_value_string(MakeString("Compile")).ToEID()))
                        /* ERROR PROTECTION INSERTED (Result-Result) */
                        if !ErrorIn(Result) {
                        if (_Zout.Value != MakeString("").Value) /* If:12 */{ 
                          m.External = _Zout
                          /* If-12 */} 
                        Result = Reader.F_load_module(m)
                        /* ERROR PROTECTION INSERTED (Result-Result) */
                        if !ErrorIn(Result) {
                        if (dblevel < 1) /* If:12 */{ 
                          /* update:13 */{ 
                            var va_arg1 *Optimize.OptimizeMetaCompiler  
                            var va_arg2 int 
                            va_arg1 = Optimize.C_compiler
                            if (Optimize.C_compiler.Safety <= 4) /* If:14 */{ 
                              va_arg2 = Optimize.C_compiler.Safety
                              } else {
                              va_arg2 = 4
                              /* If-14 */} 
                            /* ---------- now we compile update safety(va_arg1) := va_arg2 ------- */
                            va_arg1.Safety = va_arg2
                            /* update-13 */} 
                          Optimize.C_compiler.Debug_ask = Optimize.C_compiler.Debug_ask.AddFast(m.Id())
                          /* If-12 */} 
                        Result = F_compile_module(m)
                        /* ERROR PROTECTION INSERTED (Result-Result) */
                        if !ErrorIn(Result) {
                        if (_Zexe == CTRUE) /* If:12 */{ 
                          Result = F_system_file_module(m,_Zout)
                          } else {
                          Result = EID{CFALSE.Id(),0}
                          /* If-12 */} 
                        /* ERROR PROTECTION INSERTED (Result-Result) */
                        if !ErrorIn(Result) {
                        F_CL_exit(0)
                        Result = EVOID
                        }}}}
                        }
                        /* Let-11 */} 
                      } else {
                      Result = Reader.C_reader.TopLevel()
                      /* If-10 */} 
                    }}
                    if ErrorIn(Result){ 
                      /* s=EID */ClEnv.Index = h_index
                      ClEnv.Base = h_base
                      Reader.C_reader.RestoreState()
                      Result = Reader.F_debug_if_possible_void()
                      /* ERROR PROTECTION INSERTED (Result-Result) */
                      if !ErrorIn(Result) {
                      Result = Reader.C_reader.TopLevel()
                      }
                      } 
                    /* Let-9 */} 
                  /* Let-8 */} 
                /* Let-7 */} 
              /* Let-6 */} 
            /* Let-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: complex_main @ void (throw: true) 
func E_Generate_complex_main_void (_CL_obj EID) EID { 
    return /*(sm for complex_main @ void= EID)*/ F_Generate_complex_main_void( )} 
  
// *******************************************************************
// *       Part 2: System compiling methods                          *
// *******************************************************************
// generate a system file with
//   - the import
//   - the module definition
//   - calling the load() methods for the meta-descriptions
//   - the main function
/* {1} OPT.The go function for: system_file(m:module,%out:string) [] */
func F_system_file_module (m *ClaireModule ,_Zout *ClaireString ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var p *ClairePort  
      /* noccur = 2 */
      var p_try07933 EID 
      p_try07933 = F_fopen_string(F_append_string(Reader.F__7_string(Optimize.C_compiler.Source,_Zout),ToGenerateCodeProducer(Optimize.C_PRODUCER.Value).Extension),MakeString("w"))
      /* ERROR PROTECTION INSERTED (p-Result) */
      if ErrorIn(p_try07933) {Result = p_try07933
      } else {
      p = ToPort(OBJ(p_try07933))
      /* Let:3 */{ 
        var l_used *ClaireList   = Reader.F_add_modules_list(MakeConstantList(m.Id()))
        /* noccur = 2 */
        /* Let:4 */{ 
          var l_necessary *ClaireList   = F_Generate_parents_list(l_used)
          /* noccur = 2 */
          ToGenerateGoProducer(Optimize.C_PRODUCER.Value).Current = C_claire
          Optimize.C_OPT.Properties = ToType(C_property.Id()).EmptySet()
          Optimize.C_OPT.Objects = ToType(C_any.Id()).EmptyList()
          Optimize.C_OPT.Functions = ToType(C_any.Id()).EmptyList()
          Optimize.C_OPT.NeedToClose = ToType(C_any.Id()).EmptySet()
          Optimize.C_OPT.LegalModules = l_necessary.Set_I()
          p.UseAsOutput()
          PRINC("// --- System configuration file for ")
          Result = Core.F_print_any((_Zout).Id())
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(" , [")
          Result = Core.F_print_any((F_date_I_integer(1)).Id())
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC("] ---\n\n")
          Result = EVOID
          }}
          {
          PRINC("package main\n")
          F_Generate_system_imports_module(m)
          Result = F_Generate_load_function_module(m,l_necessary)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          F_Generate_main_function_list(l_used)
          p.Fclose()
          Result = EVOID
          }}
          /* Let-4 */} 
        /* Let-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: system_file @ module (throw: true) 
func E_system_file_module (m EID,_Zout EID) EID { 
    return /*(sm for system_file @ module= EID)*/ F_system_file_module(ToModule(OBJ(m)),ToString(OBJ(_Zout)) )} 
  
// create the import declaration for this system file
/* {1} OPT.The go function for: system_imports(m:module) [] */
func F_Generate_system_imports_module (m *ClaireModule )  { 
    // procedure body with s =  
PRINC("import (\n")
    PRINC("\t\"fmt\"\n")
    F_Generate_import_declaration_module(m)
    PRINC("\t\"")
    F_princ_string(m.Name.String_I())
    PRINC("\"\n")
    PRINC(")\n")
    } 
  
// The EID go function for: system_imports @ module (throw: false) 
func E_Generate_system_imports_module (m EID) EID { 
    /*(sm for system_imports @ module= void)*/ F_Generate_system_imports_module(ToModule(OBJ(m)) )
    return EVOID} 
  
// called by the run_system function. The goal of this method is
// to create the modules, and then loads all the meta-descriptions with m.MetaLoad()
// note that is status(m) = 5, we defer the load of the modules.
// we should have a declaration like deferred(m) :: m.status := 5
/* {1} OPT.The go function for: load_function(m:module,l_necessary:list) [] */
func F_Generate_load_function_module (m *ClaireModule ,l_necessary *ClaireList ) EID { 
    var Result EID 
    /* For:2 */{ 
      var x *ClaireAny  
      _ = x
      var x_support *ClaireList  
      x_support = l_necessary
      x_len := x_support.Length()
      for i_it := 0; i_it < x_len; i_it++ { 
        x = x_support.At(i_it)
        if (Equal(ANY(Core.F_CALL(C_mClaire_status,ARGS(x.ToEID()))),MakeInteger(5).Id()) == CTRUE) /* If:4 */{ 
          PRINC("func load_")
          F_iClaire_ident_symbol(ToSymbol(OBJ(Core.F_CALL(C_name,ARGS(x.ToEID())))))
          PRINC("() {")
          F_iClaire_ident_symbol(ToSymbol(OBJ(Core.F_CALL(C_name,ARGS(x.ToEID())))))
          PRINC(".metaLoad();}\n")
          /* If-4 */} 
        /* loop-3 */} 
      /* For-2 */} 
    PRINC("\n//load function : create and load modules")
    F_Generate_breakline_void()
    PRINC("")
    PRINC("func Load() ")
    F_Generate_new_block_void()
    PRINC("")
    PRINC("It := C_claire")
    F_Generate_breakline_void()
    PRINC("")
    PRINC("//module definitions ")
    F_Generate_breakline_void()
    PRINC("")
    /* For:2 */{ 
      var x *ClaireAny  
      _ = x
      Result= EID{CFALSE.Id(),0}
      var x_support *ClaireList  
      x_support = l_necessary
      x_len := x_support.Length()
      for i_it := 0; i_it < x_len; i_it++ { 
        x = x_support.At(i_it)
        var void_try4 EID 
        _ = void_try4
        if ((x != C_Kernel.Id()) && 
            ((x != C_claire.Id()) && 
              (x != C_mClaire.Id()))) /* If:4 */{ 
          void_try4 = Core.F_CALL(C_Generate_g_expression,ARGS(x.ToEID(),EID{C_module.Id(),0}))
          /* ERROR PROTECTION INSERTED (void_try4-void_try4) */
          if ErrorIn(void_try4) {Result = void_try4
          break
          } else {
          PRINC(" = InitModule(")
          void_try4 = Core.F_print_any((ToSymbol(OBJ(Core.F_CALL(C_name,ARGS(x.ToEID())))).String_I()).Id())
          /* ERROR PROTECTION INSERTED (void_try4-void_try4) */
          if ErrorIn(void_try4) {Result = void_try4
          break
          } else {
          PRINC(",")
          void_try4 = F_Generate_g_expression_module(ToModule(x).PartOf,C_module)
          /* ERROR PROTECTION INSERTED (void_try4-void_try4) */
          if ErrorIn(void_try4) {Result = void_try4
          break
          } else {
          PRINC(",")
          /* Let:5 */{ 
            var g0794UU *ClaireAny  
            /* noccur = 1 */
            var g0794UU_try07956 EID 
            g0794UU_try07956 = Core.F_CALL(Optimize.C_c_code,ARGS(EID{ToModule(x).Uses.Id(),0},EID{C_list.Id(),0}))
            /* ERROR PROTECTION INSERTED (g0794UU-void_try4) */
            if ErrorIn(g0794UU_try07956) {void_try4 = g0794UU_try07956
            } else {
            g0794UU = ANY(g0794UU_try07956)
            void_try4 = Core.F_CALL(C_Generate_g_expression,ARGS(g0794UU.ToEID(),EID{C_list.Id(),0}))
            }
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (void_try4-void_try4) */
          if ErrorIn(void_try4) {Result = void_try4
          break
          } else {
          PRINC(",")
          F_Generate_breakline_void()
          PRINC("\t")
          void_try4 = Core.F_print_any(ANY(Core.F_CALL(C_source,ARGS(x.ToEID()))))
          /* ERROR PROTECTION INSERTED (void_try4-void_try4) */
          if ErrorIn(void_try4) {Result = void_try4
          break
          } else {
          PRINC(",\n\t")
          /* Let:5 */{ 
            var g0796UU *ClaireAny  
            /* noccur = 1 */
            var g0796UU_try07976 EID 
            g0796UU_try07976 = Core.F_CALL(Optimize.C_c_code,ARGS(EID{ToModule(x).MadeOf.Id(),0},EID{C_list.Id(),0}))
            /* ERROR PROTECTION INSERTED (g0796UU-void_try4) */
            if ErrorIn(g0796UU_try07976) {void_try4 = g0796UU_try07976
            } else {
            g0796UU = ANY(g0796UU_try07976)
            void_try4 = Core.F_CALL(C_Generate_g_expression,ARGS(g0796UU.ToEID(),EID{C_list.Id(),0}))
            }
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (void_try4-void_try4) */
          if ErrorIn(void_try4) {Result = void_try4
          break
          } else {
          PRINC(");")
          F_Generate_breakline_void()
          PRINC("")
          void_try4 = EVOID
          }}}}}}
          } else {
          void_try4 = EID{CFALSE.Id(),0}
          /* If-4 */} 
        /* ERROR PROTECTION INSERTED (void_try4-Result) */
        if ErrorIn(void_try4) {Result = void_try4
        Result = void_try4
        break
        } else {
        }
        /* loop-3 */} 
      /* For-2 */} 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    F_Generate_breakline_void()
    PRINC("// module load ")
    F_Generate_breakline_void()
    PRINC("")
    /* For:2 */{ 
      var x *ClaireAny  
      _ = x
      var x_support *ClaireList  
      x_support = l_necessary
      x_len := x_support.Length()
      for i_it := 0; i_it < x_len; i_it++ { 
        x = x_support.At(i_it)
        if ((ToModule(x).MadeOf.Length() != 0) && 
            (Equal(ANY(Core.F_CALL(C_mClaire_status,ARGS(x.ToEID()))),MakeInteger(5).Id()) != CTRUE)) /* If:4 */{ 
          F_iClaire_ident_symbol(ToSymbol(OBJ(Core.F_CALL(C_name,ARGS(x.ToEID())))))
          PRINC(".MetaLoad();")
          F_Generate_breakline_void()
          PRINC("")
          /* If-4 */} 
        /* loop-3 */} 
      /* For-2 */} 
    /* For:2 */{ 
      var x *ClaireAny  
      _ = x
      var x_support *ClaireList  
      x_support = l_necessary
      x_len := x_support.Length()
      for i_it := 0; i_it < x_len; i_it++ { 
        x = x_support.At(i_it)
        if (Equal(ANY(Core.F_CALL(C_mClaire_status,ARGS(x.ToEID()))),MakeInteger(5).Id()) == CTRUE) /* If:4 */{ 
          F_iClaire_ident_symbol(ToSymbol(OBJ(Core.F_CALL(C_name,ARGS(x.ToEID())))))
          PRINC(".it->evaluate = ")
          Core.F_CALL(C_Generate_expression,ARGS(F_make_function_string(F_append_string(MakeString("load_"),ToSymbol(OBJ(Core.F_CALL(C_name,ARGS(x.ToEID())))).String_I())).ToEID(),EID{CFALSE.Id(),0}))
          F_Generate_breakline_void()
          PRINC("")
          F_iClaire_ident_symbol(ToSymbol(OBJ(Core.F_CALL(C_name,ARGS(x.ToEID())))))
          PRINC(".it->status = 2;")
          F_Generate_breakline_void()
          PRINC("")
          /* If-4 */} 
        /* loop-3 */} 
      /* For-2 */} 
    PRINC("ClEnv.Module_I = ")
    Result = F_Generate_g_expression_module(m,C_module)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("; ")
    F_Generate_breakline_void()
    PRINC("")
    Result = EVOID
    }
    {
    F_Generate_close_block_void()
    Result = EVOID
    }}
    return Result} 
  
// The EID go function for: load_function @ module (throw: true) 
func E_Generate_load_function_module (m EID,l_necessary EID) EID { 
    return /*(sm for load_function @ module= EID)*/ F_Generate_load_function_module(ToModule(OBJ(m)),ToList(OBJ(l_necessary)) )} 
  
// create the main function
/* {1} OPT.The go function for: main_function(l_used:list[module]) [] */
func F_Generate_main_function_list (l_used *ClaireList )  { 
    // procedure body with s =  
PRINC("\n// the main function \n")
    PRINC("func main() ")
    F_Generate_new_block_void()
    PRINC("")
    PRINC("MemoryFlags()")
    F_Generate_breakline_void()
    PRINC("")
    PRINC("fmt.Printf(\"=== CLAIRE4 interpreter version 1.0    ===\\n\")")
    F_Generate_breakline_void()
    PRINC("")
    PRINC("Bootstrap()")
    F_Generate_breakline_void()
    PRINC("")
    PRINC("Load()")
    F_Generate_breakline_void()
    PRINC("")
    PRINC("ClEnv.Module_I = C_claire")
    F_Generate_breakline_void()
    PRINC("")
    PRINC("Reader.C_reader.Fromp = ClEnv.Cin")
    F_Generate_breakline_void()
    PRINC("")
    if (ToBoolean(l_used.Contain_ask(F_value_string(MakeString("Generate"))).Id()) == CTRUE) /* If:2 */{ 
      PRINC("Generate.F_Generate_complex_main_void()")
      } else {
      PRINC("Reader.F_Reader_simple_main_void()")
      /* If-2 */} 
    F_Generate_breakline_void()
    F_Generate_close_block_void()
    } 
  
// The EID go function for: main_function @ list (throw: false) 
func E_Generate_main_function_list (l_used EID) EID { 
    /*(sm for main_function @ list= void)*/ F_Generate_main_function_list(ToList(OBJ(l_used)) )
    return EVOID} 
  
// *******************************************************************
// *       Part 3: module compiling : execute a command line         *
// *******************************************************************
// *******************************************************************
// *       Part 4: Profiler code                                     *
// *******************************************************************
// ---------------------------------------------------------------