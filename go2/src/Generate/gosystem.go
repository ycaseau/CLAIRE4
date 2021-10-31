/***** CLAIRE Compilation of file /Users/ycaseau/Dropbox/src/clairev4.0/src/compile/gosystem.cl 
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
func import_g0035() { 
    _ = Core.It
    _ = Language.It
    _ = Reader.It
    _ = Optimize.It
    } 
  
  
//+-------------------------------------------------------------+
//| CLAIRE                                                      |
//| gosystem.cl                                                 |
//| Copyright (C) 2020-2021 Yves Caseau. All Rights Reserved    |
//| cf. copyright info in file object.cl: about()               |
//+-------------------------------------------------------------+
//**********************************************************************
//* Contents                                                           *
//*          Part 1: Global_variables & producer interface             *
//*          Part 2: Module Compiler Interface                         *
//*          Part 3: File Compiler                                     *
//*          Part 4: Function Compiler                                 *
//**********************************************************************
// content map (represent the tree with a indented hierarchy :))
// compile [Part 2]
//      - gen_files
//            - gen_file
//      - gen_module_file
//            - start_file
//            - gen_objects
//            - gen_classes
//  compile_lambda -> ... -> make_go_function [Part 4]
//       - gen_func_start   [gogen]
//       - function_body, procedure_body, eid_body
//       - generate_eid_function
//       - check_sort
// the form is the expected go type : a ClaireX, a native (4) or EID
//**********************************************************************
//*          Part 1: Global_variables                                  *
//**********************************************************************
// ----------------------- inline coding --------------------------------
// here we have a list of methods that we want to handle in a special way
// CLAIRE4.0 duplicate (list -> tuple)
// special method in types => dispatch to any
// Compile/*min_integer* :: (min @ integer)
// Compile/*max_integer* :: (max @ integer)
// v3.2.58  */
// bag methods could be ommited in the future - these methods are used to force compiling (could be removed later)
// force goMethod()
// new: the target code production (the part that depends on the target language) is
// encapsulated with a producer object
// CLAIRE 4 is focused on go, but we try to keep the previous structure of CLAIRE3 to be ready
// for Java or Swift compiling. However, the GC management stuff is lost forever :)
// v3.3.32: stats about GC protection  */
// add the go_producer here  (replaces the C++ producer)
// note that the double list bad/good names is ugly and should be replaced by a dictionary later 
// where to place the go code
// TODO: define a status = 3 for the PRODUCER class that tells that is it extensible
// (Genearate/producer.open := 3)
// this is a special case : the function may return an error but the optimized form does not
// most standard method: call the producer to print the ident from a symbol
/* {1} OPT.The go function for: iClaire/ident(self:symbol) [] */
func F_iClaire_ident_symbol (self *ClaireSymbol )  { 
    // procedure body with s =  
F_iClaire_ident_go_producer2(ToGenerateGoProducer(Optimize.C_PRODUCER.Value),self)
    } 
  
// The EID go function for: iClaire/ident @ symbol (throw: false) 
func E_iClaire_ident_symbol (self EID) EID { 
    /*(sm for iClaire/ident @ symbol= void)*/ F_iClaire_ident_symbol(ToSymbol(OBJ(self)) )
    return EVOID} 
  
/* {1} OPT.The go function for: iClaire/ident(self:thing) [] */
func F_iClaire_ident_thing (self *ClaireThing )  { 
    // procedure body with s =  
F_iClaire_ident_go_producer2(ToGenerateGoProducer(Optimize.C_PRODUCER.Value),self.Name)
    } 
  
// The EID go function for: iClaire/ident @ thing (throw: false) 
func E_iClaire_ident_thing (self EID) EID { 
    /*(sm for iClaire/ident @ thing= void)*/ F_iClaire_ident_thing(ToThing(OBJ(self)) )
    return EVOID} 
  
/* {1} OPT.The go function for: iClaire/ident(self:class) [] */
func F_iClaire_ident_class (self *ClaireClass )  { 
    // procedure body with s =  
F_iClaire_ident_go_producer2(ToGenerateGoProducer(Optimize.C_PRODUCER.Value),self.Name)
    } 
  
// The EID go function for: iClaire/ident @ class (throw: false) 
func E_iClaire_ident_class (self EID) EID { 
    /*(sm for iClaire/ident @ class= void)*/ F_iClaire_ident_class(ToClass(OBJ(self)) )
    return EVOID} 
  
// we simply use some smart identation. True pretty_printing will be left to bc
/* {1} OPT.The go function for: indent_c(_CL_obj:void) [] */
func F_Generate_indent_c_void () *ClaireAny  { 
    // procedure body with s =  
var Result *ClaireAny  
    /* Let:2 */{ 
      var x int  = Optimize.C_OPT.Level
      /* noccur = 3 */
      Result= CFALSE.Id()
      for (x > 0) /* while:3 */{ 
        PRINC("  ")
        x = (x-1)
        /* while-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: indent_c @ void (throw: false) 
func E_Generate_indent_c_void (_CL_obj EID) EID { 
    return /*(sm for indent_c @ void= any)*/ F_Generate_indent_c_void( ).ToEID()} 
  
/* {1} OPT.The go function for: breakline(_CL_obj:void) [] */
func F_Generate_breakline_void () *ClaireAny  { 
    // use function body compiling 
PRINC("\n")
    return  F_Generate_indent_c_void()
    } 
  
// The EID go function for: breakline @ void (throw: false) 
func E_Generate_breakline_void (_CL_obj EID) EID { 
    return /*(sm for breakline @ void= any)*/ F_Generate_breakline_void( ).ToEID()} 
  
// adds a new C block with the condensed option
/* {1} OPT.The go function for: new_block(_CL_obj:void) [] */
func F_Generate_new_block_void ()  { 
    // procedure body with s =  
Optimize.C_OPT.Level = (Optimize.C_OPT.Level+1)
    PRINC("{ ")
    F_Generate_breakline_void()
    } 
  
// The EID go function for: new_block @ void (throw: false) 
func E_Generate_new_block_void (_CL_obj EID) EID { 
    /*(sm for new_block @ void= void)*/ F_Generate_new_block_void( )
    return EVOID} 
  
// closes the current C block
/* {1} OPT.The go function for: close_block(_CL_obj:void) [] */
func F_Generate_close_block_void ()  { 
    // procedure body with s =  
Optimize.C_OPT.Level = (Optimize.C_OPT.Level-1)
    PRINC("} ")
    F_Generate_breakline_void()
    } 
  
// The EID go function for: close_block @ void (throw: false) 
func E_Generate_close_block_void (_CL_obj EID) EID { 
    /*(sm for close_block @ void= void)*/ F_Generate_close_block_void( )
    return EVOID} 
  
// prints the } without a new line - used for nested If
/* {1} OPT.The go function for: finish_block(_CL_obj:void) [] */
func F_Generate_finish_block_void ()  { 
    // procedure body with s =  
Optimize.C_OPT.Level = (Optimize.C_OPT.Level-1)
    PRINC("} ")
    } 
  
// The EID go function for: finish_block @ void (throw: false) 
func E_Generate_finish_block_void (_CL_obj EID) EID { 
    /*(sm for finish_block @ void= void)*/ F_Generate_finish_block_void( )
    return EVOID} 
  
//*********************************************************************
//*          Part 2: Module Compiler Interface                        *
//*********************************************************************
// a small test function for the compiler
/* {1} OPT.The go function for: g_test(x:any) [] */
func F_g_test_any (x *ClaireAny ) EID { 
    var Result EID 
    Result = F_g_test_module(C_claire,x)
    return Result} 
  
// The EID go function for: g_test @ any (throw: true) 
func E_g_test_any (x EID) EID { 
    return /*(sm for g_test @ any= EID)*/ F_g_test_any(ANY(x) )} 
  
/* {1} OPT.The go function for: g_test(m:module,x:any) [] */
func F_g_test_module (m *ClaireModule ,x *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var t *ClaireAny  
      /* noccur = 4 */
      var t_try00393 EID 
      t_try00393 = Core.F_CALL(Optimize.C_c_type,ARGS(x.ToEID()))
      /* ERROR PROTECTION INSERTED (t-Result) */
      if ErrorIn(t_try00393) {Result = t_try00393
      } else {
      t = ANY(t_try00393)
      /* Let:3 */{ 
        var s *ClaireClass   = Optimize.F_Compile_osort_any(t)
        /* noccur = 2 */
        /* Let:4 */{ 
          var u *ClaireAny  
          /* noccur = 6 */
          var u_try00405 EID 
          u_try00405 = Core.F_CALL(Optimize.C_c_code,ARGS(x.ToEID(),EID{s.Id(),0}))
          /* ERROR PROTECTION INSERTED (u-Result) */
          if ErrorIn(u_try00405) {Result = u_try00405
          } else {
          u = ANY(u_try00405)
          /* Let:5 */{ 
            var f *ClaireBoolean  
            /* noccur = 2 */
            var f_try00556 EID 
            f_try00556 = F_Generate_g_func_any(u)
            /* ERROR PROTECTION INSERTED (f-Result) */
            if ErrorIn(f_try00556) {Result = f_try00556
            } else {
            f = ToBoolean(OBJ(f_try00556))
            /* Let:6 */{ 
              var gt *ClaireBoolean  
              /* noccur = 2 */
              var gt_try00807 EID 
              gt_try00807 = Optimize.F_Compile_g_throw_any(u)
              /* ERROR PROTECTION INSERTED (gt-Result) */
              if ErrorIn(gt_try00807) {Result = gt_try00807
              } else {
              gt = ToBoolean(OBJ(gt_try00807))
              ToGenerateGoProducer(Optimize.C_PRODUCER.Value).Current = m
              PRINC("type -> ")
              Result = Core.F_print_any(t)
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              PRINC(" [sort ")
              Result = Core.F_print_any(s.Id())
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              PRINC("]\n")
              Result = EVOID
              }}
              {
              PRINC("opt[")
              Result = Core.F_print_any(u.Isa.Id())
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              PRINC("] -> ")
              Result = Core.F_CALL(C_print,ARGS(u.ToEID()))
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              PRINC(" \n")
              Result = EVOID
              }}
              {
              if (gt == CTRUE) /* If:7 */{ 
                PRINC("----------------------- Error is possible => EID (func:")
                Result = Core.F_print_any(f.Id())
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                PRINC(")  ----------------\n")
                Result = EVOID
                }
                } else {
                Result = EID{CFALSE.Id(),0}
                /* If-7 */} 
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              if (f == CTRUE) /* If:7 */{ 
                PRINC("exp  -> ")
                Result = Core.F_CALL(C_Generate_g_expression,ARGS(u.ToEID(),EID{ToTypeExpression(t).Class_I().Id(),0}))
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                PRINC("\n")
                Result = EVOID
                }
                } else {
                PRINC("stat -> ")
                /* Let:8 */{ 
                  var g0081UU *ClaireClass  
                  /* noccur = 1 */
                  if (gt == CTRUE) /* If:9 */{ 
                    g0081UU = Optimize.C_EID
                    } else {
                    g0081UU = ToTypeExpression(t).Class_I()
                    /* If-9 */} 
                  Result = F_Generate_statement_any(u,g0081UU,MakeString("result"),C_void.Id())
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                PRINC("\n")
                Result = EVOID
                }
                /* If-7 */} 
              }}}
              }
              /* Let-6 */} 
            }
            /* Let-5 */} 
          }
          /* Let-4 */} 
        /* Let-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: g_test @ module (throw: true) 
func E_g_test_module (m EID,x EID) EID { 
    return /*(sm for g_test @ module= EID)*/ F_g_test_module(ToModule(OBJ(m)),ANY(x) )} 
  
// even more fun 
/* {1} OPT.The go function for: gtop(_CL_obj:void) [] */
func F_Generate_gtop_void () EID { 
    var Result EID 
    PRINC("in> ")
    /* Let:2 */{ 
      var x *ClaireAny  
      /* noccur = 2 */
      var x_try00823 EID 
      x_try00823 = Reader.F_read_port(ToPort(Reader.C_stdin.Value))
      /* ERROR PROTECTION INSERTED (x-Result) */
      if ErrorIn(x_try00823) {Result = x_try00823
      } else {
      x = ANY(x_try00823)
      if (x == Reader.C_q.Id()) /* If:3 */{ 
        PRINC("bye.\n")
        Result = EVOID
        } else {
        Result = Core.F_CALL(C_g_test,ARGS(x.ToEID()))
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        Result = F_Generate_gtop_void()
        }
        /* If-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: gtop @ void (throw: true) 
func E_Generate_gtop_void (_CL_obj EID) EID { 
    return /*(sm for gtop @ void= EID)*/ F_Generate_gtop_void( )} 
  
// test the compiling of a method
// e.f. g_test(foo @ any)
/* {1} OPT.The go function for: g_test(m:method) [] */
func F_g_test_method (m *ClaireMethod ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var l *ClaireLambda   = m.Formula
      /* noccur = 2 */
      if (l.Id() == CNULL) /* If:3 */{ 
        Result = EID{CNULL,0}
        } else {
        Core.F_tformat_string(MakeString("---- Compiling ~S with following definition ---- \n"),0,MakeConstantList(m.Id()))
        Result = Language.F_pretty_print_any(l.Body)
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        Optimize.C_OPT.InMethod = m.Id()
        Optimize.C_OPT.UseStringUpdate = CFALSE
        Optimize.C_OPT.MaxVars = 0
        Optimize.C_OPT.LegalModules = C_module.Instances.Set_I()
        Optimize.C_OPT.Outfile = ToPort(Reader.C_stdout.Value)
        Optimize.C_compiler.Inline_ask = CTRUE
        ToGenerateGoProducer(Optimize.C_PRODUCER.Value).Current = C_claire
        Core.F_tformat_string(MakeString("---- code produced by the optimizer ------------------- \n"),0,ToType(CEMPTY.Id()).EmptyList())
        Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).MakeGoFunction(m.Formula,MakeString("test"),m)
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        /* update:4 */{ 
          var va_arg1 *Optimize.OptimizeMetaOPT  
          var va_arg2 *ClaireAny  
          va_arg1 = Optimize.C_OPT
          va_arg2 = CNULL
          /* ---------- now we compile update Compile/in_method(va_arg1) := va_arg2 ------- */
          va_arg1.InMethod = va_arg2
          Result = va_arg2.ToEID()
          /* update-4 */} 
        }}
        /* If-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: g_test @ method (throw: true) 
func E_g_test_method (m EID) EID { 
    return /*(sm for g_test @ method= EID)*/ F_g_test_method(ToMethod(OBJ(m)) )} 
  
// debug (to remove later)
// compile the modules and check that no necessary modules is not
// declared
/* {1} OPT.The go function for: compile(m:module) [] */
func F_compile_module (m *ClaireModule ) EID { 
    var Result EID 
    Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).Compile(m)
    return Result} 
  
// The EID go function for: compile @ module (throw: true) 
func E_compile_module (m EID) EID { 
    return /*(sm for compile @ module= EID)*/ F_compile_module(ToModule(OBJ(m)) )} 
  
//  shortcut that already exists
/* {1} OPT.The go function for: compile(p:go_producer,m:module) [] */
func (p *GenerateGoProducer ) Compile (m *ClaireModule ) EID { 
    var Result EID 
    Core.F_tformat_string(MakeString("--------- start compile ------------------- \n"),0,ToType(CEMPTY.Id()).EmptyList())
    Optimize.C_OPT.NeedModules = CEMPTY
    C_BadMethods.Value = ToType(C_method.Id()).EmptyList().Id()
    Optimize.C_compiler.Inline_ask = CTRUE
    /* Let:2 */{ 
      var l1 *ClaireBag   = ToBag(F_Generate_parents_list(Reader.F_add_modules_list(MakeConstantList(m.Id()))).Id())
      /* noccur = 5 */
      Core.F_tformat_string(MakeString("==========  START GO COMPILING (~S) with ~S ================ \n"),0,MakeConstantList(m.Id(),l1.Id()))
      /* update:3 */{ 
        var va_arg1 *Optimize.OptimizeMetaOPT  
        var va_arg2 *ClaireSet  
        va_arg1 = Optimize.C_OPT
        var va_arg2_try00834 EID 
        va_arg2_try00834 = Core.F_CALL(C_set_I,ARGS(EID{l1.Id(),0}))
        /* ERROR PROTECTION INSERTED (va_arg2-Result) */
        if ErrorIn(va_arg2_try00834) {Result = va_arg2_try00834
        } else {
        va_arg2 = ToSet(OBJ(va_arg2_try00834))
        /* ---------- now we compile update Compile/legal_modules(va_arg1) := va_arg2 ------- */
        va_arg1.LegalModules = va_arg2
        Result = EID{va_arg2.Id(),0}
        }
        /* update-3 */} 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      p.Current = m
      p.Source = Reader.F__7_string(Optimize.C_compiler.Source,m.Name.String_I())
      Result = p.GenFiles(m)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = p.GenModFile(m)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      l1 = ToBag(Core.F_difference_set(Core.F_set_I_set(Optimize.C_OPT.NeedModules),Optimize.C_OPT.LegalModules).Id())
      if (F_boolean_I_any(l1.Id()) == CTRUE) /* If:3 */{ 
        Result = Core.F_tformat_string(MakeString("---- WARNING: ~S should be declared for ~S \n"),2,MakeConstantList(l1.Id(),m.Id()))
        } else {
        Result = EID{CFALSE.Id(),0}
        /* If-3 */} 
      }}}
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: compile @ go_producer (throw: true) 
func E_compile_go_producer (p EID,m EID) EID { 
    return /*(sm for compile @ go_producer= EID)*/ ToGenerateGoProducer(OBJ(p)).Compile(ToModule(OBJ(m)) )} 
  
// the first part is to generate the go files in the FileToFile mode
/* {1} OPT.The go function for: gen_files(p:go_producer,m:module) [] */
func (p *GenerateGoProducer ) GenFiles (m *ClaireModule ) EID { 
    var Result EID 
    Core.F_tformat_string(MakeString("==== Generate ~A files for module ~S [verbose = ~A, Opt? = ~S] \n"),0,MakeConstantList((ToGenerateCodeProducer(Optimize.C_PRODUCER.Value).Comment).Id(),
      m.Id(),
      MakeInteger(ClEnv.Verbose).Id(),
      Optimize.C_compiler.Optimize_ask.Id()))
    Optimize.C_OPT.Instructions = ToType(C_any.Id()).EmptyList()
    Optimize.C_OPT.Properties = ToType(C_property.Id()).EmptySet()
    Optimize.C_OPT.Objects = ToType(C_object.Id()).EmptyList()
    Optimize.C_OPT.Functions = ToType(C_any.Id()).EmptyList()
    Optimize.C_OPT.NeedToClose = ToType(C_any.Id()).EmptySet()
    m.Begin()
    /* For:2 */{ 
      var x *ClaireAny  
      _ = x
      Result= EID{CFALSE.Id(),0}
      var x_support *ClaireList  
      x_support = m.MadeOf
      x_len := x_support.Length()
      for i_it := 0; i_it < x_len; i_it++ { 
        x = x_support.At(i_it)
        var void_try4 EID 
        _ = void_try4
        { 
        Core.F_tformat_string(MakeString("++++ Compiling the file ~A.cl [v. 3.~A - safety:~A] \n"),2,MakeConstantList(x,Optimize.C_compiler.Version,MakeInteger(Optimize.C_compiler.Safety).Id()))
        if (Equal(x,(m.Name.String_I()).Id()) == CTRUE) /* If:4 */{ 
          void_try4 = Core.F_CALL(Optimize.C_Compile_Cerror,ARGS(EID{MakeString("[211]  ~S cannot be used both as a file and module name").Id(),0},x.ToEID()))
          } else {
          void_try4 = EID{CFALSE.Id(),0}
          /* If-4 */} 
        /* ERROR PROTECTION INSERTED (void_try4-void_try4) */
        if ErrorIn(void_try4) {Result = void_try4
        break
        } else {
        Optimize.C_OPT.Level = 1
        p.CurrentFile = ToString(x)
        void_try4 = p.GenFile(Reader.F__7_string(m.Source,ToString(x)),Reader.F__7_string(p.Source,ToString(x)))
        /* ERROR PROTECTION INSERTED (void_try4-void_try4) */
        if ErrorIn(void_try4) {Result = void_try4
        break
        } else {
        }}
        }
        /* loop-3 */} 
      /* For-2 */} 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    m.End()
    Result = EVOID
    }
    return Result} 
  
// The EID go function for: gen_files @ go_producer (throw: true) 
func E_Generate_gen_files_go_producer (p EID,m EID) EID { 
    return /*(sm for gen_files @ go_producer= EID)*/ ToGenerateGoProducer(OBJ(p)).GenFiles(ToModule(OBJ(m)) )} 
  
// This makes the initial loading function by compilinf all the claire
// expression placed in the list oself. *new_objects* holds all the new
// objects defined in this file.
// The name of the function is built from the file name (s argument)
//
/* {1} OPT.The go function for: gen_mod_file(p:go_producer,m:module) [] */
func (p *GenerateGoProducer ) GenModFile (m *ClaireModule ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var prt *ClairePort  
      /* noccur = 1 */
      var prt_try00853 EID 
      prt_try00853 = F_fopen_string(F_append_string(F_append_string(F_append_string(ToGenerateGoProducer(Optimize.C_PRODUCER.Value).Source,ToString(Reader.C__starfs_star.Value)),m.Name.String_I()),ToGenerateCodeProducer(Optimize.C_PRODUCER.Value).Extension),MakeString("w"))
      /* ERROR PROTECTION INSERTED (prt-Result) */
      if ErrorIn(prt_try00853) {Result = prt_try00853
      } else {
      prt = ToPort(OBJ(prt_try00853))
      /* Let:3 */{ 
        var s *ClaireString   = m.Name.String_I()
        /* noccur = 1 */
        Core.F_tformat_string(MakeString("==== generate file for module ~S ==== \n"),0,MakeConstantList(m.Id()))
        Optimize.C_OPT.Outfile = prt
        ClEnv.Verbose = 3
        Result = p.StartFile(s,m,CTRUE)
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        Optimize.C_OPT.Outfile.UseAsOutput()
        Result = p.GenClasses(m)
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        Result = p.GenObjects(m)
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        Result = p.GenMetaLoad(m)
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        F_Generate_breakline_void()
        F_Generate_close_block_void()
        F_Generate_breakline_void()
        Core.F_tformat_string(MakeString("===== ~A BAD METHODS : ~S \n"),0,MakeConstantList(MakeInteger(ToList(C_BadMethods.Value).Length()).Id(),C_BadMethods.Value))
        Optimize.C_OPT.Outfile.Fclose()
        Result = EVOID
        }}}}
        /* Let-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: gen_mod_file @ go_producer (throw: true) 
func E_Generate_gen_mod_file_go_producer (p EID,m EID) EID { 
    return /*(sm for gen_mod_file @ go_producer= EID)*/ ToGenerateGoProducer(OBJ(p)).GenModFile(ToModule(OBJ(m)) )} 
  
// start the produced go file
// Puts the reference to the packages, and some useful comments
// we limit the use of "unsafe" Go package to the module file
/* {1} OPT.The go function for: start_file(p:go_producer,s:string,m:module,module?:boolean) [] */
func (p *GenerateGoProducer ) StartFile (s *ClaireString ,m *ClaireModule ,module_ask *ClaireBoolean ) EID { 
    var Result EID 
    Optimize.C_OPT.Outfile.UseAsOutput()
    PRINC("/***** CLAIRE Compilation of ")
    F_princ_string(ToString(IfThenElse((module_ask == CTRUE),
      MakeString("module").Id(),
      MakeString("file").Id())))
    PRINC(" ")
    F_princ_string(s)
    PRINC(".cl \n         [version ")
    /* Let:2 */{ 
      var g0088UU *ClaireAny  
      /* noccur = 1 */
      var g0088UU_try00893 EID 
      g0088UU_try00893 = Core.F_release_void()
      /* ERROR PROTECTION INSERTED (g0088UU-Result) */
      if ErrorIn(g0088UU_try00893) {Result = g0088UU_try00893
      } else {
      g0088UU = ANY(g0088UU_try00893)
      Result = Core.F_CALL(C_princ,ARGS(g0088UU.ToEID()))
      }
      /* Let-2 */} 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(" / safety ")
    Result = Core.F_print_any(MakeInteger(Optimize.C_compiler.Safety).Id())
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("] ")
    F_princ_string(F_substring_string(F_date_I_integer(1),1,24))
    PRINC(" *****/\n\n")
    Result = EVOID
    }}
    {
    p.Namespace_I(m)
    PRINC("import (_ \"fmt\"\n")
    if (module_ask == CTRUE) /* If:2 */{ 
      PRINC("\t\"unsafe\"\n")
      /* If-2 */} 
    F_Generate_import_declaration_module(m)
    PRINC(")\n")
    F_Generate_dumb_import_module(m)
    Result = ToPort(Reader.C_stdout.Value).UseAsOutput().ToEID()
    }
    return Result} 
  
// The EID go function for: start_file @ go_producer (throw: true) 
func E_Generate_start_file_go_producer (p EID,s EID,m EID,module_ask EID) EID { 
    return /*(sm for start_file @ go_producer= EID)*/ ToGenerateGoProducer(OBJ(p)).StartFile(ToString(OBJ(s)),
      ToModule(OBJ(m)),
      ToBoolean(OBJ(module_ask)) )} 
  
// import declarations
/* {1} OPT.The go function for: import_declaration(m:module) [] */
func F_Generate_import_declaration_module (m *ClaireModule )  { 
    // procedure body with s =  
/* For:2 */{ 
      var x *ClaireAny  
      _ = x
      var x_support *ClaireList  
      x_support = F_Generate_needed_modules_module(m)
      x_len := x_support.Length()
      for i_it := 0; i_it < x_len; i_it++ { 
        x = x_support.At(i_it)
        PRINC("\t")
        if (x == C_Kernel.Id()) /* If:4 */{ 
          PRINC(". ")
          /* If-4 */} 
        PRINC("\"")
        F_princ_string(ToSymbol(OBJ(Core.F_CALL(C_name,ARGS(x.ToEID())))).String_I())
        PRINC("\"\n")
        /* loop-3 */} 
      /* For-2 */} 
    } 
  
// The EID go function for: import_declaration @ module (throw: false) 
func E_Generate_import_declaration_module (m EID) EID { 
    /*(sm for import_declaration @ module= void)*/ F_Generate_import_declaration_module(ToModule(OBJ(m)) )
    return EVOID} 
  
// go requires an import list without redundancy + we only import
/* {1} OPT.The go function for: needed_modules(m:module) [] */
func F_Generate_needed_modules_module (m *ClaireModule ) *ClaireList  { 
    // procedure body with s =  
var Result *ClaireList  
    /* Let:2 */{ 
      var l *ClaireList  
      /* noccur = 1 */
      /* Let:3 */{ 
        var m2_out *ClaireList   = ToType(CEMPTY.Id()).EmptyList()
        /* noccur = 2 */
        /* For:4 */{ 
          var m2 *ClaireAny  
          _ = m2
          var m2_support *ClaireList  
          m2_support = Reader.F_add_modules_list(MakeConstantList(m.Id()))
          m2_len := m2_support.Length()
          for i_it := 0; i_it < m2_len; i_it++ { 
            m2 = m2_support.At(i_it)
            if (m2 != m.Id()) /* If:6 */{ 
              if ((ToModule(m2).MadeOf.Length() != 0) || 
                  (m2 == C_Kernel.Id())) /* If:7 */{ 
                m2_out.AddFast(m2)
                /* If-7 */} 
              /* If-6 */} 
            /* loop-5 */} 
          /* For-4 */} 
        l = m2_out
        /* Let-3 */} 
      Result = l
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: needed_modules @ module (throw: false) 
func E_Generate_needed_modules_module (m EID) EID { 
    return EID{/*(sm for needed_modules @ module= list)*/ F_Generate_needed_modules_module(ToModule(OBJ(m)) ).Id(),0}} 
  
// create a dumb function that prevents the go compiler to complain
/* {1} OPT.The go function for: dumb_import(m:module) [] */
func F_Generate_dumb_import_module (m *ClaireModule )  { 
    // procedure body with s =  
/* Let:2 */{ 
      var l *ClaireList   = F_Generate_needed_modules_module(m)
      /* noccur = 2 */
      if (l.Length() > 1) /* If:3 */{ 
        PRINC("\n//-------- dumb function to prevent import errors --------\n")
        PRINC("func import_")
        Core.F_gensym_void().Princ()
        PRINC("() ")
        F_Generate_new_block_void()
        PRINC("")
        /* For:4 */{ 
          var x *ClaireAny  
          _ = x
          var x_support *ClaireList  
          x_support = l
          x_len := x_support.Length()
          for i_it := 0; i_it < x_len; i_it++ { 
            x = x_support.At(i_it)
            if (x != C_Kernel.Id()) /* If:6 */{ 
              PRINC("_ = ")
              F_Generate_cap_short_symbol(ToSymbol(OBJ(Core.F_CALL(C_name,ARGS(x.ToEID())))))
              PRINC(".It")
              F_Generate_breakline_void()
              PRINC("")
              /* If-6 */} 
            /* loop-5 */} 
          /* For-4 */} 
        F_Generate_close_block_void()
        F_Generate_breakline_void()
        /* If-3 */} 
      /* Let-2 */} 
    } 
  
// The EID go function for: dumb_import @ module (throw: false) 
func E_Generate_dumb_import_module (m EID) EID { 
    /*(sm for dumb_import @ module= void)*/ F_Generate_dumb_import_module(ToModule(OBJ(m)) )
    return EVOID} 
  
// pick a thing in module m
/* {1} OPT.The go function for: representative(m:module) [] */
func F_Generate_representative_module (m *ClaireModule ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var x_some *ClaireAny   = CNULL
      /* noccur = 2 */
      /* For:3 */{ 
        var x *ClaireAny  
        _ = x
        Result= EID{CFALSE.Id(),0}
        var x_support *ClaireList  
        var x_support_try00924 EID 
        x_support_try00924 = Core.F_enumerate_any(Core.F_U_type(ToType(C_class.Id()),ToType(C_property.Id())).Id())
        /* ERROR PROTECTION INSERTED (x_support-Result) */
        if ErrorIn(x_support_try00924) {Result = x_support_try00924
        } else {
        x_support = ToList(OBJ(x_support_try00924))
        x_len := x_support.Length()
        for i_it := 0; i_it < x_len; i_it++ { 
          x = x_support.At(i_it)
          var g0093I *ClaireBoolean  
          /* and:5 */{ 
            var v_and5 *ClaireBoolean  
            
            v_and5 = Equal(ToSymbol(OBJ(Core.F_CALL(C_name,ARGS(x.ToEID())))).Defined().Id(),m.Id())
            if (v_and5 == CFALSE) {g0093I = CFALSE
            } else /* arg:6 */{ 
              if (x.Isa.IsIn(C_property) == CTRUE) /* If:7 */{ 
                /* Let:8 */{ 
                  var g0090 *ClaireProperty   = ToProperty(x)
                  /* noccur = 1 */
                  /* Let:9 */{ 
                    var g0094UU *ClaireAny  
                    /* noccur = 1 */
                    /* For:10 */{ 
                      var y *ClaireAny  
                      _ = y
                      g0094UU= CFALSE.Id()
                      for _,y = range(g0090.Restrictions.ValuesO())/* loop:11 */{ 
                        if (ToRestriction(y).Module_I.Id() == m.Id()) /* If:12 */{ 
                           /*v = g0094UU, s =any*/
g0094UU = CTRUE.Id()
                          break
                          /* If-12 */} 
                        /* loop-11 */} 
                      /* For-10 */} 
                    v_and5 = F_boolean_I_any(g0094UU)
                    /* Let-9 */} 
                  /* Let-8 */} 
                } else {
                v_and5 = CTRUE
                /* If-7 */} 
              if (v_and5 == CFALSE) {g0093I = CFALSE
              } else /* arg:7 */{ 
                g0093I = CTRUE/* arg-7 */} 
              /* arg-6 */} 
            /* and-5 */} 
          if (g0093I == CTRUE) /* If:5 */{ 
             /*v = Result, s =EID*/
x_some = x
            Result = x_some.ToEID()
            break
            /* If-5 */} 
          }
          /* loop-4 */} 
        /* For-3 */} 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = x_some.ToEID()
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: representative @ module (throw: true) 
func E_Generate_representative_module (m EID) EID { 
    return /*(sm for representative @ module= EID)*/ F_Generate_representative_module(ToModule(OBJ(m)) )} 
  
// remove dual imports (hopefully, works if the import path is simple enough)
/* {1} OPT.The go function for: clean_duplicates(l:list) [] */
func F_Generate_clean_duplicates_list (l *ClaireList ) *ClaireList  { 
    // procedure body with s =  
var Result *ClaireList  
    /* Let:2 */{ 
      var l2 *ClaireList   = l.Copy()
      /* noccur = 3 */
      /* Let:3 */{ 
        var n int  = l.Length()
        /* noccur = 2 */
        /* Let:4 */{ 
          var i int  = (n-1)
          /* noccur = 6 */
          for (i > 1) /* while:5 */{ 
            var g0096I *ClaireBoolean  
            /* Let:6 */{ 
              var g0097UU *ClaireAny  
              /* noccur = 1 */
              /* Let:7 */{ 
                var j int  = (i+1)
                /* noccur = 4 */
                /* Let:8 */{ 
                  var g0095 int  = n
                  /* noccur = 1 */
                  g0097UU= CFALSE.Id()
                  for (j <= g0095) /* while:9 */{ 
                    if (ToBoolean(Reader.F_add_modules_list(MakeConstantList(l.At(j-1))).Contain_ask(l.At(i-1)).Id()) == CTRUE) /* If:10 */{ 
                       /*v = g0097UU, s =any*/
g0097UU = CTRUE.Id()
                      break
                      /* If-10 */} 
                    j = (j+1)
                    /* while-9 */} 
                  /* Let-8 */} 
                /* Let-7 */} 
              g0096I = F_boolean_I_any(g0097UU)
              /* Let-6 */} 
            if (g0096I == CTRUE) /* If:6 */{ 
              l2 = l2.Delete(l.At(i-1))
              /* If-6 */} 
            i = (i-1)
            /* while-5 */} 
          Result = l2
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: clean_duplicates @ list (throw: false) 
func E_Generate_clean_duplicates_list (l EID) EID { 
    return EID{/*(sm for clean_duplicates @ list= list)*/ F_Generate_clean_duplicates_list(ToList(OBJ(l)) ).Id(),0}} 
  
// For each class we produce two things in the module-generated-file
//   - the struct (with embedded inheritance)
//   - the cast method
//   - we also gerenate a constructor  makeC(a1, ... , an) when there are no inverses 
/* {1} OPT.The go function for: gen_classes(p:go_producer,m:module) [] */
func (p *GenerateGoProducer ) GenClasses (m *ClaireModule ) EID { 
    var Result EID 
    Core.F_tformat_string(MakeString("===== generate classes for ~S ==== \n"),3,MakeConstantList(m.Id()))
    /* For:2 */{ 
      var c *ClaireAny  
      _ = c
      Result= EID{CFALSE.Id(),0}
      var c_support *ClaireList  
      c_support = Optimize.C_OPT.Objects
      c_len := c_support.Length()
      for i_it := 0; i_it < c_len; i_it++ { 
        c = c_support.At(i_it)
        var void_try4 EID 
        _ = void_try4
        if (C_class.Id() == c.Isa.Id()) /* If:4 */{ 
          Optimize.C_OPT.Level = 0
          PRINC("\n// class file for ")
          void_try4 = Core.F_CALL(C_print,ARGS(c.ToEID()))
          /* ERROR PROTECTION INSERTED (void_try4-void_try4) */
          if ErrorIn(void_try4) {Result = void_try4
          break
          } else {
          PRINC(" in module ")
          void_try4 = Core.F_print_any(m.Id())
          /* ERROR PROTECTION INSERTED (void_try4-void_try4) */
          if ErrorIn(void_try4) {Result = void_try4
          break
          } else {
          PRINC(" ")
          void_try4 = EVOID
          }}
          {
          F_Generate_breakline_void()
          void_try4 = p.GenClassDef(ToClass(c))
          /* ERROR PROTECTION INSERTED (void_try4-void_try4) */
          if ErrorIn(void_try4) {Result = void_try4
          break
          } else {
          p.GenCastFunction(ToClass(c))
          if (F_Generate_construct_class_ask_class(ToClass(c)) == CTRUE) /* If:5 */{ 
            void_try4 = p.GenConstruct(ToClass(c))
            } else {
            void_try4 = EID{CFALSE.Id(),0}
            /* If-5 */} 
          /* ERROR PROTECTION INSERTED (void_try4-void_try4) */
          if ErrorIn(void_try4) {Result = void_try4
          break
          } else {
          }}}
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
    return Result} 
  
// The EID go function for: gen_classes @ go_producer (throw: true) 
func E_Generate_gen_classes_go_producer (p EID,m EID) EID { 
    return /*(sm for gen_classes @ go_producer= EID)*/ ToGenerateGoProducer(OBJ(p)).GenClasses(ToModule(OBJ(m)) )} 
  
// how to generate a struct associated to a class
/* {1} OPT.The go function for: gen_class_def(p:go_producer,c:class) [] */
func (p *GenerateGoProducer ) GenClassDef (c *ClaireClass ) EID { 
    var Result EID 
    PRINC("type ")
    F_Generate_go_class_class(c)
    PRINC(" struct ")
    F_Generate_new_block_void()
    PRINC(" ")
    F_Generate_go_class_class(c.Superclass)
    F_Generate_breakline_void()
    PRINC(" ")
    /* For:2 */{ 
      var y *ClaireAny  
      _ = y
      Result= EID{CFALSE.Id(),0}
      var y_support *ClaireList  
      var y_support_try00983 EID 
      y_support_try00983 = Core.F_CALL(Optimize.C_Compile_get_indexed,ARGS(EID{c.Id(),0}))
      /* ERROR PROTECTION INSERTED (y_support-Result) */
      if ErrorIn(y_support_try00983) {Result = y_support_try00983
      } else {
      y_support = ToList(OBJ(y_support_try00983))
      y_len := y_support.Length()
      for i_it := 0; i_it < y_len; i_it++ { 
        y = y_support.At(i_it)
        if (Core.F_domain_I_restriction(ToRestriction(y)).Id() == c.Id()) /* If:4 */{ 
          F_Generate_cap_short_symbol(ToRestriction(y).Selector.Name)
          PRINC(" ")
          F_Generate_interface_I_class(ToRestriction(y).Range.Class_I())
          PRINC("")
          F_Generate_breakline_void()
          /* If-4 */} 
        }
        /* loop-3 */} 
      /* For-2 */} 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    F_Generate_close_block_void()
    PRINC("")
    Result = EVOID
    }
    return Result} 
  
// The EID go function for: gen_class_def @ go_producer (throw: true) 
func E_Generate_gen_class_def_go_producer (p EID,c EID) EID { 
    return /*(sm for gen_class_def @ go_producer= EID)*/ ToGenerateGoProducer(OBJ(p)).GenClassDef(ToClass(OBJ(c)) )} 
  
// how to produce the ToC() cast function that applies to any pointer (using unsafe)
/* {1} OPT.The go function for: gen_cast_function(p:go_producer,c:class) [] */
func (p *GenerateGoProducer ) GenCastFunction (c *ClaireClass )  { 
    // procedure body with s =  
PRINC("\n// automatic cast function\n")
    PRINC("func ")
    F_Generate_cast_class_class(c)
    PRINC("(x *ClaireAny) *")
    F_Generate_go_class_class(c)
    PRINC(" {return (*")
    F_Generate_go_class_class(c)
    PRINC(")(unsafe.Pointer(x))}")
    F_Generate_breakline_void()
    } 
  
// The EID go function for: gen_cast_function @ go_producer (throw: false) 
func E_Generate_gen_cast_function_go_producer (p EID,c EID) EID { 
    /*(sm for gen_cast_function @ go_producer= void)*/ ToGenerateGoProducer(OBJ(p)).GenCastFunction(ToClass(OBJ(c)) )
    return EVOID} 
  
// when we want a constructor ? when slots are simple (no inverse, no store ...)
// TODO : to complete with the proper test
/* {1} OPT.The go function for: construct_class?(c:class) [] */
func F_Generate_construct_class_ask_class (c *ClaireClass ) *ClaireBoolean  { 
    // use function body compiling 
return  MakeBoolean((ToType(c.Id()).Included(ToType(C_object.Id())) == CTRUE) && (c.Slots.Length() <= 5))
    } 
  
// The EID go function for: construct_class? @ class (throw: false) 
func E_Generate_construct_class_ask_class (c EID) EID { 
    return EID{/*(sm for construct_class? @ class= boolean)*/ F_Generate_construct_class_ask_class(ToClass(OBJ(c)) ).Id(),0}} 
  
// generate a constructor
/* {1} OPT.The go function for: gen_construct(p:go_producer,c:class) [] */
func (p *GenerateGoProducer ) GenConstruct (c *ClaireClass ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var first *ClaireBoolean   = CTRUE
      /* noccur = 2 */
      PRINC("\n// automatic constructor function\n")
      PRINC("func Make")
      F_Generate_addUnderscore_class(c)
      F_Generate_go_class_class(c)
      PRINC("(")
      /* For:3 */{ 
        var y *ClaireAny  
        _ = y
        Result= EID{CFALSE.Id(),0}
        var y_support *ClaireList  
        var y_support_try00994 EID 
        /* Let:4 */{ 
          var g0100UU *ClaireList  
          /* noccur = 1 */
          var g0100UU_try01015 EID 
          g0100UU_try01015 = Core.F_CALL(Optimize.C_Compile_get_indexed,ARGS(EID{c.Id(),0}))
          /* ERROR PROTECTION INSERTED (g0100UU-y_support_try00994) */
          if ErrorIn(g0100UU_try01015) {y_support_try00994 = g0100UU_try01015
          } else {
          g0100UU = ToList(OBJ(g0100UU_try01015))
          y_support_try00994 = g0100UU.Cdr()
          }
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (y_support-Result) */
        if ErrorIn(y_support_try00994) {Result = y_support_try00994
        } else {
        y_support = ToList(OBJ(y_support_try00994))
        y_len := y_support.Length()
        for i_it := 0; i_it < y_len; i_it++ { 
          y = y_support.At(i_it)
          if (first == CTRUE) /* If:5 */{ 
            first = CFALSE
            } else {
            PRINC(",")
            /* If-5 */} 
          F_iClaire_ident_symbol(ToRestriction(y).Selector.Name)
          PRINC(" ")
          F_Generate_interface_I_class(ToRestriction(y).Range.Class_I())
          PRINC("")
          }
          /* loop-4 */} 
        /* For-3 */} 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(") *")
      F_Generate_go_class_class(c)
      PRINC(" ")
      F_Generate_new_block_string(MakeString("make"))
      PRINC("")
      Result = EVOID
      }
      {
      PRINC("var o *")
      F_Generate_go_class_class(c)
      PRINC(" = new(")
      F_Generate_go_class_class(c)
      PRINC(")")
      F_Generate_breakline_void()
      PRINC("")
      PRINC("o.Isa = ")
      F_Generate_class_ident_class(c)
      F_Generate_breakline_void()
      PRINC("")
      /* For:3 */{ 
        var y *ClaireAny  
        _ = y
        Result= EID{CFALSE.Id(),0}
        var y_support *ClaireList  
        var y_support_try01024 EID 
        /* Let:4 */{ 
          var g0103UU *ClaireList  
          /* noccur = 1 */
          var g0103UU_try01045 EID 
          g0103UU_try01045 = Core.F_CALL(Optimize.C_Compile_get_indexed,ARGS(EID{c.Id(),0}))
          /* ERROR PROTECTION INSERTED (g0103UU-y_support_try01024) */
          if ErrorIn(g0103UU_try01045) {y_support_try01024 = g0103UU_try01045
          } else {
          g0103UU = ToList(OBJ(g0103UU_try01045))
          y_support_try01024 = g0103UU.Cdr()
          }
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (y_support-Result) */
        if ErrorIn(y_support_try01024) {Result = y_support_try01024
        } else {
        y_support = ToList(OBJ(y_support_try01024))
        y_len := y_support.Length()
        for i_it := 0; i_it < y_len; i_it++ { 
          y = y_support.At(i_it)
          PRINC("o.")
          F_Generate_cap_short_symbol(ToRestriction(y).Selector.Name)
          PRINC(" = ")
          F_iClaire_ident_symbol(ToRestriction(y).Selector.Name)
          F_Generate_breakline_void()
          PRINC("")
          }
          /* loop-4 */} 
        /* For-3 */} 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC("return o ")
      F_Generate_breakline_void()
      F_Generate_close_block_string(MakeString("make"))
      PRINC("")
      Result = EVOID
      }}
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: gen_construct @ go_producer (throw: true) 
func E_Generate_gen_construct_go_producer (p EID,c EID) EID { 
    return /*(sm for gen_construct @ go_producer= EID)*/ ToGenerateGoProducer(OBJ(p)).GenConstruct(ToClass(OBJ(c)) )} 
  
// generate the definition of the named objects from the module (used in both modes)
// must move to the producer
/* {1} OPT.The go function for: gen_objects(p:go_producer,m:module) [] */
func (p *GenerateGoProducer ) GenObjects (m *ClaireModule ) EID { 
    var Result EID 
    Core.F_tformat_string(MakeString("===== generate objects for ~S [graph : ~S] ==== \n"),3,MakeConstantList(m.Id(),Optimize.C_OPT.Properties.Contain_ask(C_mClaire_graph.Id()).Id()))
    /* For:2 */{ 
      var x *ClaireAny  
      _ = x
      var x_support *ClaireList  
      x_support = Optimize.C_OPT.Objects
      x_len := x_support.Length()
      for i_it := 0; i_it < x_len; i_it++ { 
        x = x_support.At(i_it)
        F_Generate_breakline_void()
        if (x.Isa.IsIn(Core.C_global_variable) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var g0105 *Core.GlobalVariable   = Core.ToGlobalVariable(x)
            /* noccur = 3 */
            PRINC("var ")
            F_Generate_go_var_symbol(g0105.Name)
            PRINC(" ")
            /* Let:6 */{ 
              var g0107UU *ClaireAny  
              /* noccur = 1 */
              if (Optimize.F_Compile_nativeVar_ask_global_variable(g0105) == CTRUE) /* If:7 */{ 
                g0107UU = ANY(Core.F_CALL(C_Generate_getRange,ARGS(EID{g0105.Id(),0})))
                } else {
                g0107UU = Core.C_global_variable.Id()
                /* If-7 */} 
              F_Generate_interface_I_class(ToClass(g0107UU))
              /* Let-6 */} 
            PRINC("")
            /* Let-5 */} 
          } else {
          PRINC("var ")
          F_Generate_go_var_symbol(ToSymbol(OBJ(Core.F_CALL(C_name,ARGS(x.ToEID())))))
          PRINC(" ")
          F_Generate_interface_I_class(Optimize.F_Compile_psort_any(x.Isa.Id()))
          PRINC(" /*obj*/")
          /* If-4 */} 
        /* loop-3 */} 
      /* For-2 */} 
    /* For:2 */{ 
      var x *ClaireAny  
      _ = x
      Result= EID{CFALSE.Id(),0}
      for _,x = range(Optimize.C_OPT.Properties.Values)/* loop:3 */{ 
        var void_try4 EID 
        _ = void_try4
        if (Optimize.C_OPT.Objects.Memq(x) != CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var p2 *ClaireAny  
            /* noccur = 2 */
            /* Let:6 */{ 
              var p2_some *ClaireAny   = CNULL
              /* noccur = 2 */
              /* For:7 */{ 
                var p2 *ClaireAny  
                _ = p2
                for _,p2 = range(Optimize.C_OPT.Properties.Values)/* loop:8 */{ 
                  if (p2 != x) /* If:9 */{ 
                    if (ToSymbol(OBJ(Core.F_CALL(C_name,ARGS(p2.ToEID())))).String_I().Value == ToThing(x).Name.String_I().Value) /* If:10 */{ 
                       /*v = p2, s =void*/
p2_some = p2
                      break
                      /* If-10 */} 
                    /* If-9 */} 
                  /* loop-8 */} 
                /* For-7 */} 
              p2 = p2_some
              /* Let-6 */} 
            if (p2 != CNULL) /* If:6 */{ 
              void_try4 = ToException(Core.C_general_error.Make(MakeString("[217] ~S and ~S cannot be defined in the same module").Id(),MakeConstantList(p2,x).Id())).Close()
              } else {
              void_try4 = EID{CNULL,0}
              /* If-6 */} 
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (void_try4-void_try4) */
          if ErrorIn(void_try4) {Result = void_try4
          break
          } else {
          F_Generate_breakline_void()
          PRINC("var ")
          F_Generate_thing_ident_thing(ToThing(x))
          PRINC(" ")
          F_Generate_interface_I_class(Optimize.F_Compile_psort_any(x.Isa.Id()))
          PRINC(" // ")
          void_try4 = Core.F_print_any(ToThing(x).Name.Id())
          /* ERROR PROTECTION INSERTED (void_try4-void_try4) */
          if ErrorIn(void_try4) {Result = void_try4
          break
          } else {
          PRINC("")
          void_try4 = EVOID
          }
          {
          }}
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
    PRINC("var It *ClaireModule")
    F_Generate_breakline_void()
    PRINC("")
    /* Let:2 */{ 
      var m1 *ClaireModule   = m
      /* noccur = 2 */
      /* Let:3 */{ 
        var m2 *ClaireModule   = m.PartOf
        /* noccur = 7 */
        for ((m2.Id() != C_claire.Id()) && 
            (m2.Parts.At(1-1) == m1.Id())) /* while:4 */{ 
          if (Equal(m2.MadeOf.Id(),CNIL.Id()) == CTRUE) /* If:5 */{ 
            PRINC("var ")
            F_Generate_go_var_symbol(m2.Name)
            PRINC(" *ClaireModule ")
            /* If-5 */} 
          m1 = m2
          m2 = m2.PartOf
          /* while-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    Result = F_Generate_breakline_void().ToEID()
    }
    return Result} 
  
// The EID go function for: gen_objects @ go_producer (throw: true) 
func E_Generate_gen_objects_go_producer (p EID,m EID) EID { 
    return /*(sm for gen_objects @ go_producer= EID)*/ ToGenerateGoProducer(OBJ(p)).GenObjects(ToModule(OBJ(m)) )} 
  
// generate the meta_load function
// in go the load function for M is M_load()
/* {1} OPT.The go function for: gen_meta_load(p:go_producer,m:module) [] */
func (p *GenerateGoProducer ) GenMetaLoad (m *ClaireModule ) EID { 
    var Result EID 
    Core.F_tformat_string(MakeString("===== generate meta_load function for ~S ==== \n"),3,MakeConstantList(m.Id()))
    PRINC("// definition of the meta-model for module ")
    Result = Core.F_print_any(m.Id())
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(" ")
    F_Generate_breakline_void()
    PRINC("")
    Result = EVOID
    }
    {
    PRINC("func MetaLoad() ")
    F_Generate_new_block_void()
    F_Generate_breakline_void()
    PRINC("")
    Result = p.GenModule(m,m)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("// definition of the properties ")
    F_Generate_breakline_void()
    PRINC("")
    /* For:2 */{ 
      var x *ClaireAny  
      _ = x
      Result= EID{CFALSE.Id(),0}
      for _,x = range(Optimize.C_OPT.Properties.Values)/* loop:3 */{ 
        var void_try4 EID 
        _ = void_try4
        if ((Optimize.C_OPT.Objects.Memq(x) != CTRUE) && 
            ((x != C_value.Id()) && 
              (x != C_vars.Id()))) /* If:4 */{ 
          F_Generate_breakline_void()
          F_Generate_thing_ident_thing(ToThing(x))
          PRINC(" = ")
          void_try4 = p.Declare(ToProperty(x))
          /* ERROR PROTECTION INSERTED (void_try4-void_try4) */
          if ErrorIn(void_try4) {Result = void_try4
          break
          } else {
          PRINC("")
          void_try4 = EVOID
          }
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
    F_Generate_breakline_void()
    PRINC("// instructions from module sources")
    /* Let:2 */{ 
      var j *ClaireAny   = CNULL
      /* noccur = 2 */
      /* For:3 */{ 
        var i *ClaireAny  
        _ = i
        Result= EID{CFALSE.Id(),0}
        var i_support *ClaireList  
        i_support = Optimize.C_OPT.Instructions
        i_len := i_support.Length()
        for i_it := 0; i_it < i_len; i_it++ { 
          i = i_support.At(i_it)
          var void_try5 EID 
          _ = void_try5
          { 
          F_Generate_breakline_void()
          if (C_string.Id() == i.Isa.Id()) /* If:5 */{ 
            if (C_string.Id() != j.Isa.Id()) /* If:6 */{ 
              F_Generate_breakline_void()
              /* If-6 */} 
            PRINC("// ")
            void_try5 = Core.F_CALL(C_princ,ARGS(i.ToEID()))
            /* ERROR PROTECTION INSERTED (void_try5-void_try5) */
            if ErrorIn(void_try5) {Result = void_try5
            break
            } else {
            PRINC("")
            void_try5 = EVOID
            }
            } else {
            var g0108I *ClaireBoolean  
            var g0108I_try01096 EID 
            g0108I_try01096 = Optimize.F_Compile_g_throw_any(i)
            /* ERROR PROTECTION INSERTED (g0108I-void_try5) */
            if ErrorIn(g0108I_try01096) {void_try5 = g0108I_try01096
            } else {
            g0108I = ToBoolean(OBJ(g0108I_try01096))
            if (g0108I == CTRUE) /* If:6 */{ 
              F_Generate_new_block_void()
              PRINC("/*PROTECT ")
              void_try5 = Core.F_CALL(C_print,ARGS(i.ToEID()))
              /* ERROR PROTECTION INSERTED (void_try5-void_try5) */
              if ErrorIn(void_try5) {Result = void_try5
              break
              } else {
              PRINC(" */")
              F_Generate_breakline_void()
              PRINC("")
              void_try5 = EVOID
              }
              {
              F_Generate_var_declaration_string(MakeString("expr"),Optimize.C_EID,1)
              void_try5 = Core.F_CALL(C_Generate_g_statement,ARGS(i.ToEID(),
                EID{Optimize.C_EID.Id(),0},
                EID{MakeString("expr").Id(),0},
                EID{CTRUE.Id(),0},
                EID{CFALSE.Id(),0}))
              /* ERROR PROTECTION INSERTED (void_try5-void_try5) */
              if ErrorIn(void_try5) {Result = void_try5
              break
              } else {
              PRINC("ErrorCheck(expr)")
              F_Generate_close_block_void()
              void_try5 = EVOID
              }}
              } else {
              void_try5 = F_Generate_statement_any(i,C_void,MakeString("Niet"),CFALSE.Id())
              /* If-6 */} 
            }
            /* If-5 */} 
          /* ERROR PROTECTION INSERTED (void_try5-void_try5) */
          if ErrorIn(void_try5) {Result = void_try5
          break
          } else {
          j = i
          }
          }
          /* loop-4 */} 
        /* For-3 */} 
      /* Let-2 */} 
    }}}
    return Result} 
  
// The EID go function for: gen_meta_load @ go_producer (throw: true) 
func E_Generate_gen_meta_load_go_producer (p EID,m EID) EID { 
    return /*(sm for gen_meta_load @ go_producer= EID)*/ ToGenerateGoProducer(OBJ(p)).GenMetaLoad(ToModule(OBJ(m)) )} 
  
// generate the module definition - only the module structure (the decoration is found in the system file)
// cool recursive method that ensures that all non-package modules are visible
// load_m() has an implicit begin(m) so that new methods are assigned to m
/* {1} OPT.The go function for: gen_module(p:go_producer,m:module,%package:module) [] */
func (p *GenerateGoProducer ) GenModule (m *ClaireModule ,_Zpackage *ClaireModule ) EID { 
    var Result EID 
    if (m.Id() == _Zpackage.Id()) /* If:2 */{ 
      PRINC("It")
      } else {
      F_Generate_go_var_symbol(m.Name)
      PRINC("")
      /* If-2 */} 
    PRINC(" = MakeModule(")
    Result = Core.F_print_any((m.Name.String_I()).Id())
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(",")
    Result = F_Generate_g_expression_module(m.PartOf,C_module)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(")")
    F_Generate_breakline_void()
    PRINC("")
    Result = EVOID
    }}
    {
    PRINC("ClEnv.Module_I = It")
    F_Generate_breakline_void()
    PRINC("")
    Result = EVOID
    }
    return Result} 
  
// The EID go function for: gen_module @ go_producer (throw: true) 
func E_Generate_gen_module_go_producer (p EID,m EID,_Zpackage EID) EID { 
    return /*(sm for gen_module @ go_producer= EID)*/ ToGenerateGoProducer(OBJ(p)).GenModule(ToModule(OBJ(m)),ToModule(OBJ(_Zpackage)) )} 
  
// implicit begin(m)
// reciprocate : finds the concrete module where a package module must be defined.
/* {1} OPT.The go function for: get_made(self:module) [] */
func F_Generate_get_made_module (self *ClaireModule ) *ClaireModule  { 
    // procedure body with s =  
var Result *ClaireModule  
    /* Let:2 */{ 
      var m *ClaireAny   = self.Parts.At(1-1)
      /* noccur = 4 */
      if ((m == C_Kernel.Id()) || 
          (ToModule(m).MadeOf.Length() != 0)) /* If:3 */{ 
        Result = ToModule(m)
        } else {
        Result = F_Generate_get_made_module(ToModule(m))
        /* If-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: get_made @ module (throw: false) 
func E_Generate_get_made_module (self EID) EID { 
    return EID{/*(sm for get_made @ module= module)*/ F_Generate_get_made_module(ToModule(OBJ(self)) ).Id(),0}} 
  
// called by gosystem.cl : declare a property or an operation (handles the dispatch case)
/* {1} OPT.The go function for: declare(c:go_producer,p:property) [] */
func (c *GenerateGoProducer ) Declare (p *ClaireProperty ) EID { 
    var Result EID 
    PRINC("Make")
    F_princ_string(ToString(IfThenElse((p.Isa.IsIn(C_operation) == CTRUE),
      MakeString("Operation").Id(),
      MakeString("Property").Id())))
    PRINC("(")
    Result = Core.F_print_any((p.Name.String_I()).Id())
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(",")
    F_princ_integer(p.Open)
    PRINC(",")
    Result = F_Generate_g_expression_module(p.Name.Module_I(),C_module)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    if (p.Isa.IsIn(C_operation) == CTRUE) /* If:2 */{ 
      PRINC(",")
      F_princ_integer(ToOperation(p.Id()).Precedence)
      PRINC("")
      /* If-2 */} 
    PRINC(")")
    Result = EVOID
    }}
    return Result} 
  
// The EID go function for: declare @ go_producer (throw: true) 
func E_Generate_declare_go_producer (c EID,p EID) EID { 
    return /*(sm for declare @ go_producer= EID)*/ ToGenerateGoProducer(OBJ(c)).Declare(ToProperty(OBJ(p)) )} 
  
// This is a similar method which places all the necessary modules
// in the right order so that self can be defined
/* {1} OPT.The go function for: parents(self:module,l:list) [] */
func F_Generate_parents_module (self *ClaireModule ,l *ClaireList ) *ClaireList  { 
    // use function body compiling 
if (l.Memq(self.Id()) == CTRUE) /* body If:2 */{ 
      return  l
      } else {
      if (self.PartOf.Id() != CNULL) /* If:3 */{ 
        l = F_Generate_parents_module(self.PartOf,l)
        /* If-3 */} 
      l = l.AddFast(self.Id())
      return  l
      /* body If-2 */} 
    } 
  
// The EID go function for: parents @ module (throw: false) 
func E_Generate_parents_module (self EID,l EID) EID { 
    return EID{/*(sm for parents @ module= list)*/ F_Generate_parents_module(ToModule(OBJ(self)),ToList(OBJ(l)) ).Id(),0}} 
  
// this methods takes a list of modules that must be loaded and returns
// a list of modules that are necessary for the definition
//
/* {1} OPT.The go function for: parents(self:list) [] */
func F_Generate_parents_list (self *ClaireList ) *ClaireList  { 
    // procedure body with s =  
var Result *ClaireList  
    /* Let:2 */{ 
      var l *ClaireList   = ToType(C_module.Id()).EmptyList()
      /* noccur = 3 */
      /* For:3 */{ 
        var x *ClaireAny  
        _ = x
        var x_support *ClaireList  
        x_support = self
        x_len := x_support.Length()
        for i_it := 0; i_it < x_len; i_it++ { 
          x = x_support.At(i_it)
          l = F_Generate_parents_module(ToModule(x),l)
          /* loop-4 */} 
        /* For-3 */} 
      Result = l
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: parents @ list (throw: false) 
func E_Generate_parents_list (self EID) EID { 
    return EID{/*(sm for parents @ list= list)*/ F_Generate_parents_list(ToList(OBJ(self)) ).Id(),0}} 
  
// useful (v3.0.06)
/* {1} OPT.The go function for: get(m:module) [] */
func F_get_module2 (m *ClaireModule ) EID { 
    var Result EID 
    Result = Reader.F_load_module(m)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    m.Begin()
    Result = EVOID
    }
    return Result} 
  
// The EID go function for: get @ list<type_expression>(module) (throw: true) 
func E_get_module2 (m EID) EID { 
    return /*(sm for get @ list<type_expression>(module)= EID)*/ F_get_module2(ToModule(OBJ(m)) )} 
  
// *********************************************************************
// *     Part 3: File compilation                                      *
// *********************************************************************
// this is the basic file cross_compiler, which translates from claire to go
// this file compiler runs only in the good environment (the file to be compiled must be already loaded).
// it generates methods definitions in f2 and stores the instructions into OPT.instructions
/* {1} OPT.The go function for: gen_file(p:go_producer,f1:string,f2:string) [] */
func (p *GenerateGoProducer ) GenFile (f1 *ClaireString ,f2 *ClaireString ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var p1 *ClairePort  
      /* noccur = 4 */
      var p1_try01143 EID 
      p1_try01143 = F_fopen_string(F_append_string(f1,MakeString(".cl")),MakeString("r"))
      /* ERROR PROTECTION INSERTED (p1-Result) */
      if ErrorIn(p1_try01143) {Result = p1_try01143
      } else {
      p1 = ToPort(OBJ(p1_try01143))
      /* Let:3 */{ 
        var b *ClaireBoolean   = Reader.C_reader.Toplevel
        /* noccur = 1 */
        /* Let:4 */{ 
          var p0 *ClairePort   = Reader.C_reader.Fromp
          /* noccur = 1 */
          /* update:5 */{ 
            var va_arg1 *Optimize.OptimizeMetaOPT  
            var va_arg2 *ClairePort  
            va_arg1 = Optimize.C_OPT
            var va_arg2_try01156 EID 
            va_arg2_try01156 = F_fopen_string(F_append_string(f2,p.Extension),MakeString("w"))
            /* ERROR PROTECTION INSERTED (va_arg2-Result) */
            if ErrorIn(va_arg2_try01156) {Result = va_arg2_try01156
            } else {
            va_arg2 = ToPort(OBJ(va_arg2_try01156))
            /* ---------- now we compile update Compile/outfile(va_arg1) := va_arg2 ------- */
            va_arg1.Outfile = va_arg2
            Result = va_arg2.ToEID()
            }
            /* update-5 */} 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          Reader.C_reader.Toplevel = CFALSE
          Optimize.C_compiler.Loading_ask = CTRUE
          ClEnv.NLine = 1
          Reader.C_reader.External = f1
          Reader.C_reader.Fromp = p1
          Result = p.StartFile(f1,ClEnv.Module_I,CFALSE)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          /* Let:5 */{ 
            var _Zinstruction *ClaireAny  
            /* noccur = 11 */
            var _Zinstruction_try01166 EID 
            _Zinstruction_try01166 = Reader.F_readblock_port(p1)
            /* ERROR PROTECTION INSERTED (_Zinstruction-Result) */
            if ErrorIn(_Zinstruction_try01166) {Result = _Zinstruction_try01166
            } else {
            _Zinstruction = ANY(_Zinstruction_try01166)
            Result= EID{CFALSE.Id(),0}
            for (_Zinstruction != Reader.C_Reader_eof.Id()) /* while:6 */{ 
              var void_try7 EID 
              _ = void_try7
              { 
              if ((ClEnv.Verbose > -1) && 
                  (C_string.Id() != _Zinstruction.Isa.Id())) /* If:7 */{ 
                PRINC("[")
                void_try7 = Core.F_print_any(ClEnv.Module_I.Id())
                /* ERROR PROTECTION INSERTED (void_try7-void_try7) */
                if ErrorIn(void_try7) {Result = void_try7
                break
                } else {
                PRINC("/")
                F_princ_string(F_Generate_fileName_string(f1))
                PRINC(":")
                F_princ_integer(ClEnv.NLine)
                PRINC("] ")
                void_try7 = Core.F_print_any(_Zinstruction.Isa.Id())
                /* ERROR PROTECTION INSERTED (void_try7-void_try7) */
                if ErrorIn(void_try7) {Result = void_try7
                break
                } else {
                PRINC(" (")
                void_try7 = Core.F_print_any(Optimize.C_OPT.NeedModules.Id())
                /* ERROR PROTECTION INSERTED (void_try7-void_try7) */
                if ErrorIn(void_try7) {Result = void_try7
                break
                } else {
                PRINC(")\n")
                void_try7 = EVOID
                }}}
                } else {
                void_try7 = EID{CFALSE.Id(),0}
                /* If-7 */} 
              /* ERROR PROTECTION INSERTED (void_try7-void_try7) */
              if ErrorIn(void_try7) {Result = void_try7
              break
              } else {
              if (_Zinstruction.Isa.IsIn(Language.C_Defobj) == CTRUE) /* If:7 */{ 
                /* Let:8 */{ 
                  var g0113 *Language.Defobj   = Language.To_Defobj(_Zinstruction)
                  /* noccur = 1 */
                  /* Let:9 */{ 
                    var _Zs *ClaireSymbol   = g0113.Ident
                    /* noccur = 2 */
                    PRINC("[defobj ident is ")
                    void_try7 = Core.F_print_any(_Zs.Id())
                    /* ERROR PROTECTION INSERTED (void_try7-void_try7) */
                    if ErrorIn(void_try7) {Result = void_try7
                    break
                    } else {
                    PRINC("->")
                    void_try7 = Core.F_print_any(_Zs.Defined().Id())
                    /* ERROR PROTECTION INSERTED (void_try7-void_try7) */
                    if ErrorIn(void_try7) {Result = void_try7
                    break
                    } else {
                    PRINC("]\n")
                    void_try7 = EVOID
                    }}
                    /* Let-9 */} 
                  /* Let-8 */} 
                } else {
                void_try7 = EID{CFALSE.Id(),0}
                /* If-7 */} 
              /* ERROR PROTECTION INSERTED (void_try7-void_try7) */
              if ErrorIn(void_try7) {Result = void_try7
              break
              } else {
              if (C_string.Id() == _Zinstruction.Isa.Id()) /* If:7 */{ 
                PRINC("// ")
                F_princ_string(F_substring_string(ToString(_Zinstruction),1,(INT(Core.F_CALL(C_length,ARGS(_Zinstruction.ToEID())))-1)))
                PRINC("\n")
                if (Optimize.C_compiler.Naming < 2) /* If:8 */{ 
                  /* Let:9 */{ 
                    var pp *ClairePort   = Optimize.C_OPT.Outfile.UseAsOutput()
                    /* noccur = 1 */
                    PRINC("\n//")
                    void_try7 = Core.F_CALL(C_princ,ARGS(_Zinstruction.ToEID()))
                    /* ERROR PROTECTION INSERTED (void_try7-void_try7) */
                    if ErrorIn(void_try7) {Result = void_try7
                    break
                    } else {
                    PRINC("")
                    void_try7 = EVOID
                    }
                    {
                    void_try7 = pp.UseAsOutput().ToEID()
                    }
                    /* Let-9 */} 
                  } else {
                  void_try7 = EID{CFALSE.Id(),0}
                  /* If-8 */} 
                /* ERROR PROTECTION INSERTED (void_try7-void_try7) */
                if ErrorIn(void_try7) {Result = void_try7
                break
                } else {
                }
                } else {
                /* update:8 */{ 
                  var va_arg1 *Optimize.OptimizeMetaOPT  
                  var va_arg2 *ClaireList  
                  va_arg1 = Optimize.C_OPT
                  var va_arg2_try01179 EID 
                  /* Let:9 */{ 
                    var g0118UU *ClaireAny  
                    /* noccur = 1 */
                    var g0118UU_try011910 EID 
                    g0118UU_try011910 = Core.F_CALL(Optimize.C_c_code,ARGS(_Zinstruction.ToEID(),EID{C_void.Id(),0}))
                    /* ERROR PROTECTION INSERTED (g0118UU-va_arg2_try01179) */
                    if ErrorIn(g0118UU_try011910) {va_arg2_try01179 = g0118UU_try011910
                    } else {
                    g0118UU = ANY(g0118UU_try011910)
                    va_arg2_try01179 = EID{Optimize.C_OPT.Instructions.AddFast(g0118UU).Id(),0}
                    }
                    /* Let-9 */} 
                  /* ERROR PROTECTION INSERTED (va_arg2-void_try7) */
                  if ErrorIn(va_arg2_try01179) {void_try7 = va_arg2_try01179
                  } else {
                  va_arg2 = ToList(OBJ(va_arg2_try01179))
                  /* ---------- now we compile update Compile/instructions(va_arg1) := va_arg2 ------- */
                  va_arg1.Instructions = va_arg2
                  void_try7 = EID{va_arg2.Id(),0}
                  }
                  /* update-8 */} 
                /* If-7 */} 
              /* ERROR PROTECTION INSERTED (void_try7-void_try7) */
              if ErrorIn(void_try7) {Result = void_try7
              break
              } else {
              var _Zinstruction_try01207 EID 
              _Zinstruction_try01207 = Reader.F_readblock_port(p1)
              /* ERROR PROTECTION INSERTED (_Zinstruction-void_try7) */
              if ErrorIn(_Zinstruction_try01207) {void_try7 = _Zinstruction_try01207
              Result = _Zinstruction_try01207
              break
              } else {
              _Zinstruction = ANY(_Zinstruction_try01207)
              void_try7 = _Zinstruction.ToEID()
              }}}}
              /* while-6 */} 
            }
            }
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          p1.Fclose()
          Optimize.C_compiler.Loading_ask = CFALSE
          Reader.C_reader.Toplevel = b
          Reader.C_reader.External = MakeString("toplevel")
          Reader.C_reader.Fromp = p0
          Optimize.C_OPT.Outfile.Fclose()
          Result = EVOID
          }}}
          /* Let-4 */} 
        /* Let-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: gen_file @ go_producer (throw: true) 
func E_Generate_gen_file_go_producer (p EID,f1 EID,f2 EID) EID { 
    return /*(sm for gen_file @ go_producer= EID)*/ ToGenerateGoProducer(OBJ(p)).GenFile(ToString(OBJ(f1)),ToString(OBJ(f2)) )} 
  
// sugar
/* {1} OPT.The go function for: fileName(s:string) [] */
func F_Generate_fileName_string (s *ClaireString ) *ClaireString  { 
    // procedure body with s =  
var Result *ClaireString  
    /* Let:2 */{ 
      var n int  = F_length_string(s)
      /* noccur = 1 */
      /* Let:3 */{ 
        var i int  = F_get_string(s,ToString(Reader.C__starfs_star.Value).At(1))
        /* noccur = 2 */
        if (i > 0) /* If:4 */{ 
          Result = F_Generate_fileName_string(F_substring_string(s,(i+1),n))
          } else {
          Result = s
          /* If-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: fileName @ string (throw: false) 
func E_Generate_fileName_string (s EID) EID { 
    return EID{/*(sm for fileName @ string= string)*/ F_Generate_fileName_string(ToString(OBJ(s)) ).Id(),0}} 
  
//**********************************************************************
//*     Part 4: the lambda-to-function compiler                        *
//**********************************************************************
// This is simplified in CLAIRE4 since the class2file mode is no longer supported
// we could re-introduce it from CLAIRE 3.5 if we want to support Java compiling
/* {1} OPT.The go function for: Compile/make_c_function(self:lambda,%nom:string,m:any) [] */
func F_Compile_make_c_function_lambda (self *ClaireLambda ,_Znom *ClaireString ,m *ClaireAny ) EID { 
    var Result EID 
    if (C_method.Id() == m.Isa.Id()) /* If:2 */{ 
      Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).MakeGoFunction(self,_Znom,ToMethod(m))
      } else {
      Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).MakeLambdaFunction(self,_Znom)
      /* If-2 */} 
    return Result} 
  
// The EID go function for: Compile/make_c_function @ lambda (throw: true) 
func E_Compile_make_c_function_lambda (self EID,_Znom EID,m EID) EID { 
    return /*(sm for Compile/make_c_function @ lambda= EID)*/ F_Compile_make_c_function_lambda(ToLambda(OBJ(self)),ToString(OBJ(_Znom)),ANY(m) )} 
  
// In CLAIRE 4 we separate methods from free lambdas (used for demons, but which could be used to compile lambda blocks)
// this is used for demons as well as second-order-types
// create an EID lambda  
/* {1} OPT.The go function for: make_lambda_function(p:go_producer,self:lambda,%nom:string) [] */
func (p *GenerateGoProducer ) MakeLambdaFunction (self *ClaireLambda ,_Znom *ClaireString ) EID { 
    var Result EID 
    Core.F_tformat_string(MakeString("===== generate an EID function from a lambda for ~A \n"),0,MakeConstantList((_Znom).Id()))
    Optimize.C_OPT.Outfile.UseAsOutput()
    Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).GenerateFunctionStart(self,
      Optimize.C_EID,
      CNIL.Id(),
      _Znom)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    F_Generate_new_block_void()
    PRINC("/* eid body: ")
    Result = Core.F_CALL(C_print,ARGS(self.Body.ToEID()))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(" */")
    F_Generate_breakline_void()
    PRINC("")
    Result = EVOID
    }
    {
    Result = F_Generate_eid_body_any(self.Body,CTRUE,Optimize.C_EID)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    F_Generate_close_block_void()
    F_Generate_breakline_void()
    Result = F_Generate_generate_eid_dual_lambda(self,_Znom)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    Result = ToPort(Reader.C_stdout.Value).UseAsOutput().ToEID()
    }}}}
    return Result} 
  
// The EID go function for: make_lambda_function @ go_producer (throw: true) 
func E_Generate_make_lambda_function_go_producer (p EID,self EID,_Znom EID) EID { 
    return /*(sm for make_lambda_function @ go_producer= EID)*/ ToGenerateGoProducer(OBJ(p)).MakeLambdaFunction(ToLambda(OBJ(self)),ToString(OBJ(_Znom)) )} 
  
// how to declare a function in the interface file and its header in the
// output file
/* {1} OPT.The go function for: generate_function_start(p:go_producer,self:lambda,s:class,m:any,%nom:string) [] */
func (p *GenerateGoProducer ) GenerateFunctionStart (self *ClaireLambda ,s *ClaireClass ,m *ClaireAny ,_Znom *ClaireString ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var _Zdom *ClaireAny  
      /* noccur = 2 */
      if (self.Vars.Length() != 0) /* If:3 */{ 
        _Zdom = ANY(Core.F_CALL(C_range,ARGS(self.Vars.At(1-1).ToEID())))
        } else {
        _Zdom = C_any.Id()
        /* If-3 */} 
      /* Let:3 */{ 
        var _Zf *ClaireFunction   = F_make_function_string(_Znom)
        /* noccur = 2 */
        /* Let:4 */{ 
          var lv *ClaireList  
          /* noccur = 3 */
          if ((self.Vars.Length() == 1) && 
              ((_Zdom == C_void.Id()) || 
                  (_Zdom == C_environment.Id()))) /* If:5 */{ 
            lv = CNIL
            } else {
            lv = self.Vars
            /* If-5 */} 
          Optimize.C_OPT.Functions = Optimize.C_OPT.Functions.AddFast(MakeConstantList(_Zf.Id(),lv.Id(),s.Id()).Id())
          if (Optimize.C_compiler.Naming != 2) /* If:5 */{ 
            PRINC("\n/* {")
            F_princ_integer(Optimize.C_OPT.Level)
            PRINC("} OPT.The go function for: ")
            if (C_method.Id() == m.Isa.Id()) /* If:6 */{ 
              /* Let:7 */{ 
                var g0125 *ClaireMethod   = ToMethod(m)
                /* noccur = 2 */
                Result = Core.F_print_any(g0125.Selector.Id())
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                PRINC("(")
                Result = Language.F_ppvariable_list(self.Vars)
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                PRINC(") [")
                if (Optimize.C_compiler.Naming == 1) /* If:8 */{ 
                  Core.F_CALL(C_Generate_bitvectorSum,ARGS(Core.F_CALL(C_Generate_status_I,ARGS(EID{g0125.Id(),0}))))
                  /* If-8 */} 
                PRINC("]")
                Result = EVOID
                }}
                /* Let-7 */} 
              } else {
              F_princ_string(F_string_I_function(_Zf))
              Result = EVOID
              /* If-6 */} 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC(" */\n")
            Result = EVOID
            }
            } else {
            Result = EID{CFALSE.Id(),0}
            /* If-5 */} 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          if (F_Generate_goMethod_ask_any(m) == CTRUE) /* If:5 */{ 
            PRINC("func (")
            Result = p.GoVariable(To_Variable(self.Vars.At(1-1)))
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC(") ")
            F_Generate_goMethod_method(ToMethod(m))
            PRINC(" (")
            /* Let:6 */{ 
              var g0135UU *ClaireList  
              /* noccur = 1 */
              var g0135UU_try01367 EID 
              g0135UU_try01367 = self.Vars.Cdr()
              /* ERROR PROTECTION INSERTED (g0135UU-Result) */
              if ErrorIn(g0135UU_try01367) {Result = g0135UU_try01367
              } else {
              g0135UU = ToList(OBJ(g0135UU_try01367))
              Result = p.GoVariables(g0135UU)
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC(") ")
            if (s.Id() != C_void.Id()) /* If:6 */{ 
              F_Generate_interface_I_class(s)
              /* If-6 */} 
            PRINC(" ")
            Result = EVOID
            }}
            /* If!5 */}  else if (Equal(m,CNIL.Id()) == CTRUE) /* If:5 */{ 
            PRINC("func F_")
            F_c_princ_string(_Znom)
            PRINC(" (")
            Result = p.GoVariables(lv)
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC(") EID ")
            Result = EVOID
            }
            } else {
            PRINC("func ")
            Result = F_Generate_goFunction_method(ToMethod(m))
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC(" (")
            Result = p.GoVariables(lv)
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC(") ")
            if (s.Id() != C_void.Id()) /* If:6 */{ 
              F_Generate_interface_I_class(s)
              /* If-6 */} 
            PRINC(" ")
            Result = EVOID
            }}
            /* If-5 */} 
          }
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: generate_function_start @ go_producer (throw: true) 
func E_Generate_generate_function_start_go_producer (p EID,self EID,s EID,m EID,_Znom EID) EID { 
    return /*(sm for generate_function_start @ go_producer= EID)*/ ToGenerateGoProducer(OBJ(p)).GenerateFunctionStart(ToLambda(OBJ(self)),
      ToClass(OBJ(s)),
      ANY(m),
      ToString(OBJ(_Znom)) )} 
  
// This method creates a go function from a claire lambda for a method m.
// %name is the name that was proposed for the lambda (or derived from method m)
// we either use function_body to try a simple approach or (procedure_body | eid_body) that add all the trimmings
/* {1} OPT.The go function for: make_go_function(p:go_producer,self:lambda,%nom:string,m:method) [] */
func (p *GenerateGoProducer ) MakeGoFunction (self *ClaireLambda ,_Znom *ClaireString ,m *ClaireMethod ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var typeOK *ClaireAny  
      /* noccur = 2 */
      var typeOK_try01373 EID 
      typeOK_try01373 = F_Generate_check_range_method(m,self.Body)
      /* ERROR PROTECTION INSERTED (typeOK-Result) */
      if ErrorIn(typeOK_try01373) {Result = typeOK_try01373
      } else {
      typeOK = ANY(typeOK_try01373)
      /* Let:3 */{ 
        var s *ClaireClass   = m.Range.Class_I()
        /* noccur = 6 */
        /* Let:4 */{ 
          var _Zbody *ClaireAny  
          /* noccur = 7 */
          var _Zbody_try01385 EID 
          _Zbody_try01385 = Optimize.F_Compile_c_strict_code_any(self.Body,s)
          /* ERROR PROTECTION INSERTED (_Zbody-Result) */
          if ErrorIn(_Zbody_try01385) {Result = _Zbody_try01385
          } else {
          _Zbody = ANY(_Zbody_try01385)
          /* Let:5 */{ 
            var throw_ask *ClaireBoolean  
            /* noccur = 5 */
            var throw_ask_try01396 EID 
            throw_ask_try01396 = Optimize.F_Compile_g_throw_any(_Zbody)
            /* ERROR PROTECTION INSERTED (throw_ask-Result) */
            if ErrorIn(throw_ask_try01396) {Result = throw_ask_try01396
            } else {
            throw_ask = ToBoolean(OBJ(throw_ask_try01396))
            /* Let:6 */{ 
              var g0143UU *ClaireList  
              /* noccur = 1 */
              var g0143UU_try01447 EID 
              /* Construct:7 */{ 
                var v_bag_arg *ClaireAny  
                g0143UU_try01447= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                ToList(OBJ(g0143UU_try01447)).AddFast(m.Id())
                ToList(OBJ(g0143UU_try01447)).AddFast(_Zbody)
                var v_bag_arg_try01458 EID 
                v_bag_arg_try01458 = F_Generate_simple_body_ask_any(_Zbody)
                /* ERROR PROTECTION INSERTED (v_bag_arg-g0143UU_try01447) */
                if ErrorIn(v_bag_arg_try01458) {g0143UU_try01447 = v_bag_arg_try01458
                } else {
                v_bag_arg = ANY(v_bag_arg_try01458)
                ToList(OBJ(g0143UU_try01447)).AddFast(v_bag_arg)}
                /* Construct-7 */} 
              /* ERROR PROTECTION INSERTED (g0143UU-Result) */
              if ErrorIn(g0143UU_try01447) {Result = g0143UU_try01447
              } else {
              g0143UU = ToList(OBJ(g0143UU_try01447))
              Result = Core.F_tformat_string(MakeString("---- ~S: make_go(~S) => simple=~S \n"),0,g0143UU)
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            var g0146I *ClaireBoolean  
            var g0146I_try01476 EID 
            /* and:6 */{ 
              var v_and6 *ClaireBoolean  
              
              v_and6 = Core.F__I_equal_any(MakeInteger(m.Status).Id(),MakeInteger(-1).Id())
              if (v_and6 == CFALSE) {g0146I_try01476 = EID{CFALSE.Id(),0}
              } else /* arg:7 */{ 
                var v_and6_try01488 EID 
                /* Let:8 */{ 
                  var g0149UU *ClaireBoolean  
                  /* noccur = 1 */
                  var g0149UU_try01509 EID 
                  g0149UU_try01509 = Optimize.F_Compile_can_throw_ask_method(m)
                  /* ERROR PROTECTION INSERTED (g0149UU-v_and6_try01488) */
                  if ErrorIn(g0149UU_try01509) {v_and6_try01488 = g0149UU_try01509
                  } else {
                  g0149UU = ToBoolean(OBJ(g0149UU_try01509))
                  v_and6_try01488 = EID{Core.F__I_equal_any(throw_ask.Id(),g0149UU.Id()).Id(),0}
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (v_and6-g0146I_try01476) */
                if ErrorIn(v_and6_try01488) {g0146I_try01476 = v_and6_try01488
                } else {
                v_and6 = ToBoolean(OBJ(v_and6_try01488))
                if (v_and6 == CFALSE) {g0146I_try01476 = EID{CFALSE.Id(),0}
                } else /* arg:8 */{ 
                  g0146I_try01476 = EID{CTRUE.Id(),0}/* arg-8 */} 
                /* arg-7 */} 
              }
              /* and-6 */} 
            /* ERROR PROTECTION INSERTED (g0146I-Result) */
            if ErrorIn(g0146I_try01476) {Result = g0146I_try01476
            } else {
            g0146I = ToBoolean(OBJ(g0146I_try01476))
            if (g0146I == CTRUE) /* If:6 */{ 
              Core.F_tformat_string(MakeString("======================== WARNING ======================================== \n"),0,ToType(CEMPTY.Id()).EmptyList())
              Core.F_tformat_string(MakeString(">>>>> ~S body produces an error (g_throw = true) while status is 0 <<<<<<< \n"),0,MakeConstantList(m.Id()))
              if (m.Status == 0) /* If:7 */{ 
                var v_gassign8 *ClaireAny  
                v_gassign8 = ToList(C_BadMethods.Value).AddFast(m.Id()).Id()
                C_BadMethods.Value = v_gassign8
                Result = v_gassign8.ToEID()
                } else {
                Result = EID{CFALSE.Id(),0}
                /* If-7 */} 
              } else {
              Result = EID{CFALSE.Id(),0}
              /* If-6 */} 
            }
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            Optimize.C_OPT.Outfile.UseAsOutput()
            if (((F_boolean_I_any(typeOK) == CTRUE) || 
                  (Optimize.C_compiler.Safety > 3)) && 
                ((throw_ask != CTRUE) && 
                  (m.Selector.Id() != Core.C_self_eval.Id()))) /* If:6 */{ 
              
              Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).GenerateFunctionStart(self,
                s,
                m.Id(),
                _Znom)
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              F_Generate_new_block_void()
              var g0153I *ClaireBoolean  
              var g0153I_try01547 EID 
              /* or:7 */{ 
                var v_or7 *ClaireBoolean  
                
                v_or7 = F_Generate_need_debug_ask_any(m.Id())
                if (v_or7 == CTRUE) {g0153I_try01547 = EID{CTRUE.Id(),0}
                } else /* or:8 */{ 
                  v_or7 = Optimize.C_OPT.Profile_ask
                  if (v_or7 == CTRUE) {g0153I_try01547 = EID{CTRUE.Id(),0}
                  } else /* or:9 */{ 
                    var v_or7_try015510 EID 
                    /* Let:10 */{ 
                      var g0156UU *ClaireBoolean  
                      /* noccur = 1 */
                      var g0156UU_try015711 EID 
                      g0156UU_try015711 = F_Generate_simple_body_ask_any(_Zbody)
                      /* ERROR PROTECTION INSERTED (g0156UU-v_or7_try015510) */
                      if ErrorIn(g0156UU_try015711) {v_or7_try015510 = g0156UU_try015711
                      } else {
                      g0156UU = ToBoolean(OBJ(g0156UU_try015711))
                      v_or7_try015510 = EID{g0156UU.Not.Id(),0}
                      }
                      /* Let-10 */} 
                    /* ERROR PROTECTION INSERTED (v_or7-g0153I_try01547) */
                    if ErrorIn(v_or7_try015510) {g0153I_try01547 = v_or7_try015510
                    } else {
                    v_or7 = ToBoolean(OBJ(v_or7_try015510))
                    if (v_or7 == CTRUE) {g0153I_try01547 = EID{CTRUE.Id(),0}
                    } else /* or:10 */{ 
                      v_or7 = Equal(s.Id(),C_void.Id())
                      if (v_or7 == CTRUE) {g0153I_try01547 = EID{CTRUE.Id(),0}
                      } else /* or:11 */{ 
                        g0153I_try01547 = EID{CFALSE.Id(),0}/* org-11 */} 
                      /* org-10 */} 
                    /* org-9 */} 
                  /* org-8 */} 
                }
                /* or-7 */} 
              /* ERROR PROTECTION INSERTED (g0153I-Result) */
              if ErrorIn(g0153I_try01547) {Result = g0153I_try01547
              } else {
              g0153I = ToBoolean(OBJ(g0153I_try01547))
              if (g0153I == CTRUE) /* If:7 */{ 
                Result = F_Generate_procedure_body_method(m,self,_Zbody,s)
                } else {
                PRINC("// use function body compiling \n")
                Result = Core.F_CALL(C_Generate_function_body,ARGS(_Zbody.ToEID(),EID{s.Id(),0}))
                /* If-7 */} 
              }
              }
              } else {
              Core.F_tformat_string(MakeString("--- EID function generation (can throw = ~S) \n"),0,MakeConstantList(throw_ask.Id()))
              throw_ask = CTRUE
              Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).GenerateFunctionStart(self,
                Optimize.C_EID,
                m.Id(),
                _Znom)
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              F_Generate_new_block_void()
              Result = F_Generate_eid_body_any(_Zbody,ToBoolean(typeOK),s)
              }
              /* If-6 */} 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            F_Generate_close_block_void()
            Result = F_Generate_generate_eid_function_lambda(self,m,throw_ask)
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            if (m.Selector.Id() == Core.C_self_eval.Id()) /* If:6 */{ 
              Result = F_Generate_generate_eval_function_lambda(self,m)
              } else {
              Result = EID{CFALSE.Id(),0}
              /* If-6 */} 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            Result = ToPort(Reader.C_stdout.Value).UseAsOutput().ToEID()
            }}}}}
            }
            /* Let-5 */} 
          }
          /* Let-4 */} 
        /* Let-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: make_go_function @ go_producer (throw: true) 
func E_Generate_make_go_function_go_producer (p EID,self EID,_Znom EID,m EID) EID { 
    return /*(sm for make_go_function @ go_producer= EID)*/ ToGenerateGoProducer(OBJ(p)).MakeGoFunction(ToLambda(OBJ(self)),
      ToString(OBJ(_Znom)),
      ToMethod(OBJ(m)) )} 
  
// check that we may call function_body  (replaces the print_body method of CLAIRE 3 compiler)  
// simple : we can generate ... return X directly without the need for a "Result" variable 
/* {1} OPT.The go function for: simple_body?(self:any) [] */
func F_Generate_simple_body_ask_any (self *ClaireAny ) EID { 
    var Result EID 
    if (self.Isa.IsIn(Language.C_If) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0160 *Language.If   = Language.To_If(self)
        /* noccur = 3 */
        /* and:4 */{ 
          var v_and4 *ClaireBoolean  
          
          var v_and4_try01635 EID 
          v_and4_try01635 = F_Generate_g_func_any(g0160.Test)
          /* ERROR PROTECTION INSERTED (v_and4-Result) */
          if ErrorIn(v_and4_try01635) {Result = v_and4_try01635
          } else {
          v_and4 = ToBoolean(OBJ(v_and4_try01635))
          if (v_and4 == CFALSE) {Result = EID{CFALSE.Id(),0}
          } else /* arg:5 */{ 
            var v_and4_try01646 EID 
            v_and4_try01646 = F_Generate_simple_body_ask_any(g0160.Arg)
            /* ERROR PROTECTION INSERTED (v_and4-Result) */
            if ErrorIn(v_and4_try01646) {Result = v_and4_try01646
            } else {
            v_and4 = ToBoolean(OBJ(v_and4_try01646))
            if (v_and4 == CFALSE) {Result = EID{CFALSE.Id(),0}
            } else /* arg:6 */{ 
              var v_and4_try01657 EID 
              v_and4_try01657 = F_Generate_simple_body_ask_any(g0160.Other)
              /* ERROR PROTECTION INSERTED (v_and4-Result) */
              if ErrorIn(v_and4_try01657) {Result = v_and4_try01657
              } else {
              v_and4 = ToBoolean(OBJ(v_and4_try01657))
              if (v_and4 == CFALSE) {Result = EID{CFALSE.Id(),0}
              } else /* arg:7 */{ 
                Result = EID{CTRUE.Id(),0}/* arg-7 */} 
              /* arg-6 */} 
            /* arg-5 */} 
          }}}
          /* and-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(Language.C_Do) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0161 *Language.Do   = Language.To_Do(self)
        /* noccur = 1 */
        /* Let:4 */{ 
          var g0166UU *ClaireAny  
          /* noccur = 1 */
          var g0166UU_try01675 EID 
          g0166UU_try01675 = Core.F_last_list(g0161.Args)
          /* ERROR PROTECTION INSERTED (g0166UU-Result) */
          if ErrorIn(g0166UU_try01675) {Result = g0166UU_try01675
          } else {
          g0166UU = ANY(g0166UU_try01675)
          Result = F_Generate_simple_body_ask_any(g0166UU)
          }
          /* Let-4 */} 
        /* Let-3 */} 
      } else {
      Result = F_Generate_g_func_any(self)
      /* If-2 */} 
    return Result} 
  
// The EID go function for: simple_body? @ any (throw: true) 
func E_Generate_simple_body_ask_any (self EID) EID { 
    return /*(sm for simple_body? @ any= EID)*/ F_Generate_simple_body_ask_any(ANY(self) )} 
  
// generic case (g_func is true)
// simpler case that we apply for Do, Ifs and functional expressions
/* {1} OPT.The go function for: function_body(self:any,s:class) [] */
func F_Generate_function_body_any (self *ClaireAny ,s *ClaireClass ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var _Zret *ClaireString   = ToString(IfThenElse((s.Id() != C_void.Id()),
        MakeString("return ").Id(),
        MakeString("").Id()))
      /* noccur = 1 */
      F_princ_string(_Zret)
      PRINC(" ")
      Result = Core.F_CALL(C_Generate_g_expression,ARGS(self.ToEID(),EID{s.Id(),0}))
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      F_Generate_breakline_void()
      PRINC("")
      Result = EVOID
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: function_body @ any (throw: true) 
func E_Generate_function_body_any (self EID,s EID) EID { 
    return /*(sm for function_body @ any= EID)*/ F_Generate_function_body_any(ANY(self),ToClass(OBJ(s)) )} 
  
// generate nice code for If function (inspired from g_statement@If)
/* {1} OPT.The go function for: function_body(self:If,s:class) [] */
func F_Generate_function_body_If (self *Language.If ,s *ClaireClass ) EID { 
    var Result EID 
    PRINC("if ")
    Result = Core.F_CALL(Optimize.C_Compile_bool_exp,ARGS(self.Test.ToEID(),EID{CTRUE.Id(),0}))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(" ")
    F_Generate_new_block_string(MakeString("body If"))
    PRINC("")
    Result = EVOID
    }
    {
    Result = Core.F_CALL(C_Generate_function_body,ARGS(self.Arg.ToEID(),EID{s.Id(),0}))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    if (Equal(self.Other,CNIL.Id()) == CTRUE) /* If:2 */{ 
      F_Generate_close_block_void()
      Result = EVOID
      /* If!2 */}  else if (self.Other.Isa.IsIn(Language.C_If) == CTRUE) /* If:2 */{ 
      F_Generate_finish_block_void()
      PRINC(" else ")
      Result = Core.F_CALL(C_Generate_function_body,ARGS(self.Other.ToEID(),EID{s.Id(),0}))
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC("")
      Result = EVOID
      }
      } else {
      var g0173I *ClaireBoolean  
      var g0173I_try01743 EID 
      /* or:3 */{ 
        var v_or3 *ClaireBoolean  
        
        v_or3 = Core.F__I_equal_any(s.Id(),C_void.Id())
        if (v_or3 == CTRUE) {g0173I_try01743 = EID{CTRUE.Id(),0}
        } else /* or:4 */{ 
          var v_or3_try01755 EID 
          /* Let:5 */{ 
            var g0176UU *ClaireBoolean  
            /* noccur = 1 */
            var g0176UU_try01776 EID 
            g0176UU_try01776 = Optimize.F_Compile_designated_ask_any(self.Other)
            /* ERROR PROTECTION INSERTED (g0176UU-v_or3_try01755) */
            if ErrorIn(g0176UU_try01776) {v_or3_try01755 = g0176UU_try01776
            } else {
            g0176UU = ToBoolean(OBJ(g0176UU_try01776))
            v_or3_try01755 = EID{g0176UU.Not.Id(),0}
            }
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (v_or3-g0173I_try01743) */
          if ErrorIn(v_or3_try01755) {g0173I_try01743 = v_or3_try01755
          } else {
          v_or3 = ToBoolean(OBJ(v_or3_try01755))
          if (v_or3 == CTRUE) {g0173I_try01743 = EID{CTRUE.Id(),0}
          } else /* or:5 */{ 
            g0173I_try01743 = EID{CFALSE.Id(),0}/* org-5 */} 
          /* org-4 */} 
        }
        /* or-3 */} 
      /* ERROR PROTECTION INSERTED (g0173I-Result) */
      if ErrorIn(g0173I_try01743) {Result = g0173I_try01743
      } else {
      g0173I = ToBoolean(OBJ(g0173I_try01743))
      if (g0173I == CTRUE) /* If:3 */{ 
        PRINC("} else {")
        F_Generate_breakline_void()
        Result = Core.F_CALL(C_Generate_function_body,ARGS(self.Other.ToEID(),EID{s.Id(),0}))
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        F_Generate_close_block_string(MakeString("body If"))
        PRINC("")
        Result = EVOID
        }
        } else {
        F_Generate_close_block_string(MakeString("body If"))
        Result = EVOID
        /* If-3 */} 
      }
      /* If-2 */} 
    }}
    return Result} 
  
// The EID go function for: function_body @ If (throw: true) 
func E_Generate_function_body_If (self EID,s EID) EID { 
    return /*(sm for function_body @ If= EID)*/ F_Generate_function_body_If(Language.To_If(OBJ(self)),ToClass(OBJ(s)) )} 
  
// generate nice code for a Do
/* {1} OPT.The go function for: function_body(self:Do,s:class) [] */
func F_Generate_function_body_Do (self *Language.Do ,s *ClaireClass ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var l *ClaireList   = self.Args
      /* noccur = 2 */
      /* Let:3 */{ 
        var _Zlength int  = l.Length()
        /* noccur = 1 */
        /* Let:4 */{ 
          var m int  = 0
          /* noccur = 3 */
          /* For:5 */{ 
            var x *ClaireAny  
            _ = x
            Result= EID{CFALSE.Id(),0}
            var x_support *ClaireList  
            x_support = l
            x_len := x_support.Length()
            for i_it := 0; i_it < x_len; i_it++ { 
              x = x_support.At(i_it)
              var void_try7 EID 
              _ = void_try7
              { 
              m = (m+1)
              if (m == _Zlength) /* If:7 */{ 
                void_try7 = Core.F_CALL(C_Generate_function_body,ARGS(x.ToEID(),EID{s.Id(),0}))
                } else {
                void_try7 = F_Generate_statement_any(x,C_void,MakeString("Unused"),C_void.Id())
                /* If-7 */} 
              /* ERROR PROTECTION INSERTED (void_try7-void_try7) */
              if ErrorIn(void_try7) {Result = void_try7
              break
              } else {
              }
              }
              /* loop-6 */} 
            /* For-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: function_body @ Do (throw: true) 
func E_Generate_function_body_Do (self EID,s EID) EID { 
    return /*(sm for function_body @ Do= EID)*/ F_Generate_function_body_Do(Language.To_Do(OBJ(self)),ToClass(OBJ(s)) )} 
  
// default complex case : create a variable "Result"
/* {1} OPT.The go function for: procedure_body(m:method,%l:lambda,%body:any,s:class) [] */
func F_Generate_procedure_body_method (m *ClaireMethod ,_Zl *ClaireLambda ,_Zbody *ClaireAny ,s *ClaireClass ) EID { 
    var Result EID 
    if (Optimize.C_OPT.Profile_ask == CTRUE) /* If:2 */{ 
      Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).GenerateProfile(m.Id())
      } else {
      Result = EID{CFALSE.Id(),0}
      /* If-2 */} 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    if (F_Generate_need_debug_ask_any(m.Id()) == CTRUE) /* If:2 */{ 
      Result = Core.F_CALL(C_Generate_debug_intro,ARGS(EID{Optimize.C_PRODUCER.Value,0},EID{m.Id(),0}))
      } else {
      Result = EID{CFALSE.Id(),0}
      /* If-2 */} 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("// procedure body with s = ")
    
    PRINC(" \n")
    if (s.Id() != C_void.Id()) /* If:2 */{ 
      F_Generate_var_declaration_string(MakeString("Result"),s,1)
      Result = F_Generate_statement_any(_Zbody,s,MakeString("Result"),CFALSE.Id())
      } else {
      Result = F_Generate_statement_any(_Zbody,C_void,MakeString("Unused"),CFALSE.Id())
      /* If-2 */} 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).ReturnResult(s,m,MakeString("Result"))
    }}}
    return Result} 
  
// The EID go function for: procedure_body @ method (throw: true) 
func E_Generate_procedure_body_method (m EID,_Zl EID,_Zbody EID,s EID) EID { 
    return /*(sm for procedure_body @ method= EID)*/ F_Generate_procedure_body_method(ToMethod(OBJ(m)),
      ToLambda(OBJ(_Zl)),
      ANY(_Zbody),
      ToClass(OBJ(s)) )} 
  
// generate an EID function 
// call for the debug/profile is needed     
/* {1} OPT.The go function for: eid_body(%body:any,typeOK:boolean,s:class) [] */
func F_Generate_eid_body_any (_Zbody *ClaireAny ,typeOK *ClaireBoolean ,s *ClaireClass ) EID { 
    var Result EID 
    F_Generate_var_declaration_string(MakeString("Result"),Optimize.C_EID,1)
    /* Let:2 */{ 
      var g0179UU *ClaireBoolean  
      /* noccur = 1 */
      var g0179UU_try01803 EID 
      g0179UU_try01803 = Optimize.F_Compile_g_throw_any(_Zbody)
      /* ERROR PROTECTION INSERTED (g0179UU-Result) */
      if ErrorIn(g0179UU_try01803) {Result = g0179UU_try01803
      } else {
      g0179UU = ToBoolean(OBJ(g0179UU_try01803))
      Result = F_Generate_statement_any(_Zbody,Optimize.C_EID,MakeString("Result"),g0179UU.Id())
      }
      /* Let-2 */} 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    if (typeOK == CTRUE) /* If:2 */{ 
      PRINC("return Result")
      Result = EVOID
      } else {
      PRINC("return RangeCheck(")
      Result = F_Generate_g_expression_class(s,C_type)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(",Result)")
      Result = EVOID
      }
      /* If-2 */} 
    }
    return Result} 
  
// The EID go function for: eid_body @ any (throw: true) 
func E_Generate_eid_body_any (_Zbody EID,typeOK EID,s EID) EID { 
    return /*(sm for eid_body @ any= EID)*/ F_Generate_eid_body_any(ANY(_Zbody),ToBoolean(OBJ(typeOK)),ToClass(OBJ(s)) )} 
  
// generate the EID function associated to each method (used by the interpreter - EID mode)
/* {1} OPT.The go function for: generate_eid_function(self:lambda,m:method,throw?:boolean) [] */
func F_Generate_generate_eid_function_lambda (self *ClaireLambda ,m *ClaireMethod ,throw_ask *ClaireBoolean ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var _Zsig *ClaireList   = F_Generate_go_signature_method(m)
      /* noccur = 2 */
      /* Let:3 */{ 
        var lv *ClaireList   = self.Vars
        /* noccur = 3 */
        PRINC("\n// The EID go function for: ")
        Result = Core.F_print_any(m.Id())
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC(" (throw: ")
        Result = Core.F_print_any(throw_ask.Id())
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC(") \n")
        Result = EVOID
        }}
        {
        PRINC("func ")
        Result = F_Generate_goEIDFunctionName_method(m)
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC(" (")
        Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).GoEIDVariables(lv)
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC(") EID ")
        Result = EVOID
        }}
        {
        F_Generate_new_block_void()
        if ((m.Range.Id() == C_void.Id()) && 
            (throw_ask != CTRUE)) /* If:4 */{ 
          Result = F_Generate_print_EID_call_method(m,lv,_Zsig,throw_ask)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          F_Generate_breakline_void()
          PRINC("return EVOID")
          Result = EVOID
          }
          } else {
          PRINC("return ")
          Result = F_Generate_print_EID_call_method(m,lv,_Zsig,throw_ask)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC("")
          Result = EVOID
          }
          /* If-4 */} 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        F_Generate_close_block_void()
        Result = EVOID
        }}}
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: generate_eid_function @ lambda (throw: true) 
func E_Generate_generate_eid_function_lambda (self EID,m EID,throw_ask EID) EID { 
    return /*(sm for generate_eid_function @ lambda= EID)*/ F_Generate_generate_eid_function_lambda(ToLambda(OBJ(self)),ToMethod(OBJ(m)),ToBoolean(OBJ(throw_ask)) )} 
  
// similar but simpler for a lambda associated to a name (e.g. 2nd order types) => E_C(nom)      
/* {1} OPT.The go function for: generate_eid_dual(self:lambda,%nom:string) [] */
func F_Generate_generate_eid_dual_lambda (self *ClaireLambda ,_Znom *ClaireString ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var lv *ClaireList   = self.Vars
      /* noccur = 5 */
      /* Let:3 */{ 
        var nl_ask *ClaireBoolean   = Core.F__sup_integer(lv.Length(),3)
        /* noccur = 1 */
        PRINC("\n// The dual EID go function for: ")
        Result = Core.F_print_any((_Znom).Id())
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC(" \n")
        Result = EVOID
        }
        {
        PRINC("func E_")
        F_c_princ_string(_Znom)
        PRINC(" (")
        Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).GoEIDVariables(lv)
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC(") EID ")
        Result = EVOID
        }
        {
        F_Generate_new_block_void()
        PRINC("return F_")
        F_c_princ_string(_Znom)
        PRINC("(")
        /* Let:4 */{ 
          var n int  = 1
          /* noccur = 6 */
          /* Let:5 */{ 
            var g0183 int  = lv.Length()
            /* noccur = 1 */
            Result= EID{CFALSE.Id(),0}
            for (n <= g0183) /* while:6 */{ 
              var void_try7 EID 
              _ = void_try7
              { 
              void_try7 = F_Generate_external_EID_arg_Variable(To_Variable(lv.At(n-1)),ToTypeExpression(OBJ(Core.F_CALL(C_range,ARGS(lv.At(n-1).ToEID())))).Class_I(),n,nl_ask)
              /* ERROR PROTECTION INSERTED (void_try7-void_try7) */
              if ErrorIn(void_try7) {Result = void_try7
              break
              } else {
              n = (n+1)
              }
              /* while-6 */} 
            }
            /* Let-5 */} 
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC(")")
        F_Generate_close_block_void()
        Result = EVOID
        }}}
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: generate_eid_dual @ lambda (throw: true) 
func E_Generate_generate_eid_dual_lambda (self EID,_Znom EID) EID { 
    return /*(sm for generate_eid_dual @ lambda= EID)*/ F_Generate_generate_eid_dual_lambda(ToLambda(OBJ(self)),ToString(OBJ(_Znom)) )} 
  
// EID function calls the compiled native function - uses a code that looks like print_external_call
// watch out: a method that can throw returns an EID directly ! (same as goexp.cl : print_ext_call)
/* {1} OPT.The go function for: print_EID_call(m:method,l:list,%sig:list<class>,throw?:boolean) [] */
func F_Generate_print_EID_call_method (m *ClaireMethod ,l *ClaireList ,_Zsig *ClaireList ,throw_ask *ClaireBoolean ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var n int  = 1
      /* noccur = 0 */
      _ = n
      /* Let:3 */{ 
        var sm *ClaireAny  
        /* noccur = 4 */
        var sm_try01884 EID 
        sm_try01884 = Core.F_last_list(_Zsig)
        /* ERROR PROTECTION INSERTED (sm-Result) */
        if ErrorIn(sm_try01884) {Result = sm_try01884
        } else {
        sm = ANY(sm_try01884)
        /* Let:4 */{ 
          var nl_ask *ClaireBoolean   = Core.F__sup_integer(l.Length(),3)
          /* noccur = 5 */
          if (nl_ask == CTRUE) /* If:5 */{ 
            Optimize.C_OPT.Level = (Optimize.C_OPT.Level+1)
            /* If-5 */} 
          if ((throw_ask == CTRUE) || 
              ((ToSet(C_Generate_EIDSET.Value).Contain_ask(m.Id()) == CTRUE) || 
                (m.Selector.Id() == Core.C_self_eval.Id()))) /* If:5 */{ 
            sm = Optimize.C_EID.Id()
            /* If-5 */} 
          Result = F_Generate_cast_prefix_class(ToClass(sm),Optimize.C_EID)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC("/*(sm for ")
          Result = Core.F_print_any(m.Id())
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC("= ")
          Result = Core.F_print_any(sm)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(")*/ ")
          Result = EVOID
          }}
          {
          if (F_Generate_goMethod_ask_any(m.Id()) == CTRUE) /* If:5 */{ 
            Result = F_Generate_external_EID_arg_Variable(To_Variable(l.At(1-1)),ToClass(_Zsig.ValuesO()[1-1]),1,nl_ask)
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC(".")
            F_Generate_goMethod_method(m)
            PRINC("(")
            /* Let:6 */{ 
              var n int  = 2
              /* noccur = 6 */
              /* Let:7 */{ 
                var g0186 int  = l.Length()
                /* noccur = 1 */
                Result= EID{CFALSE.Id(),0}
                for (n <= g0186) /* while:8 */{ 
                  var void_try9 EID 
                  _ = void_try9
                  { 
                  void_try9 = F_Generate_external_EID_arg_Variable(To_Variable(l.At(n-1)),ToClass(_Zsig.ValuesO()[n-1]),(n-1),nl_ask)
                  /* ERROR PROTECTION INSERTED (void_try9-void_try9) */
                  if ErrorIn(void_try9) {Result = void_try9
                  break
                  } else {
                  n = (n+1)
                  }
                  /* while-8 */} 
                }
                /* Let-7 */} 
              /* Let-6 */} 
            }
            } else {
            Result = F_Generate_goFunction_method(m)
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC("(")
            Result = EVOID
            }
            {
            if ((l.Length() == 1) && 
                (Core.F_domain_I_restriction(ToRestriction(m.Id())).Id() == C_void.Id())) /* If:6 */{ 
              l = CNIL
              /* If-6 */} 
            /* Let:6 */{ 
              var n int  = 1
              /* noccur = 6 */
              /* Let:7 */{ 
                var g0187 int  = l.Length()
                /* noccur = 1 */
                Result= EID{CFALSE.Id(),0}
                for (n <= g0187) /* while:8 */{ 
                  var void_try9 EID 
                  _ = void_try9
                  { 
                  void_try9 = F_Generate_external_EID_arg_Variable(To_Variable(l.At(n-1)),ToClass(_Zsig.ValuesO()[n-1]),n,nl_ask)
                  /* ERROR PROTECTION INSERTED (void_try9-void_try9) */
                  if ErrorIn(void_try9) {Result = void_try9
                  break
                  } else {
                  n = (n+1)
                  }
                  /* while-8 */} 
                }
                /* Let-7 */} 
              /* Let-6 */} 
            }
            /* If-5 */} 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(" )")
          if (nl_ask == CTRUE) /* If:5 */{ 
            Optimize.C_OPT.Level = (Optimize.C_OPT.Level-1)
            /* If-5 */} 
          F_Generate_cast_post_class(ToClass(sm),Optimize.C_EID)
          Result = EVOID
          }}}
          /* Let-4 */} 
        }
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: print_EID_call @ method (throw: true) 
func E_Generate_print_EID_call_method (m EID,l EID,_Zsig EID,throw_ask EID) EID { 
    return /*(sm for print_EID_call @ method= EID)*/ F_Generate_print_EID_call_method(ToMethod(OBJ(m)),
      ToList(OBJ(l)),
      ToList(OBJ(_Zsig)),
      ToBoolean(OBJ(throw_ask)) )} 
  
// here v is a EID-range variable and we need to extract the native s representation
// n=0 is a special marker when the arg the receiver x in x.f(....)
/* {1} OPT.The go function for: external_EID_arg(v:Variable,s:class,n:integer,nl?:boolean) [] */
func F_Generate_external_EID_arg_Variable (v *ClaireVariable ,s *ClaireClass ,n int,nl_ask *ClaireBoolean ) EID { 
    var Result EID 
    if (n > 1) /* If:2 */{ 
      PRINC(",")
      if (nl_ask == CTRUE) /* If:3 */{ 
        F_Generate_breakline_void()
        /* If-3 */} 
      /* If-2 */} 
    Result = F_Generate_eid_prefix_class(s)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    Result = F_iClaire_ident_go_producer1(ToGenerateGoProducer(Optimize.C_PRODUCER.Value),v)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    F_Generate_eid_post_class(s)
    Result = EVOID
    }}
    return Result} 
  
// The EID go function for: external_EID_arg @ Variable (throw: true) 
func E_Generate_external_EID_arg_Variable (v EID,s EID,n EID,nl_ask EID) EID { 
    return /*(sm for external_EID_arg @ Variable= EID)*/ F_Generate_external_EID_arg_Variable(To_Variable(OBJ(v)),
      ToClass(OBJ(s)),
      INT(n),
      ToBoolean(OBJ(nl_ask)) )} 
  
// prints a list of arguments with types / replaces typed_args_list
/* {1} OPT.The go function for: goEIDVariables(p:go_producer,self:list) [] */
func (p *GenerateGoProducer ) GoEIDVariables (self *ClaireList ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var prems *ClaireBoolean   = CTRUE
      /* noccur = 2 */
      /* For:3 */{ 
        var v *ClaireAny  
        _ = v
        Result= EID{CFALSE.Id(),0}
        var v_support *ClaireList  
        v_support = self
        v_len := v_support.Length()
        for i_it := 0; i_it < v_len; i_it++ { 
          v = v_support.At(i_it)
          var void_try5 EID 
          _ = void_try5
          { 
          if (prems == CTRUE) /* If:5 */{ 
            prems = CFALSE
            } else {
            PRINC(",")
            /* If-5 */} 
          void_try5 = F_iClaire_ident_go_producer1(p,To_Variable(v))
          /* ERROR PROTECTION INSERTED (void_try5-void_try5) */
          if ErrorIn(void_try5) {Result = void_try5
          break
          } else {
          PRINC(" EID")
          void_try5 = EVOID
          }
          {
          }
          }
          /* loop-4 */} 
        /* For-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: goEIDVariables @ go_producer (throw: true) 
func E_Generate_goEIDVariables_go_producer (p EID,self EID) EID { 
    return /*(sm for goEIDVariables @ go_producer= EID)*/ ToGenerateGoProducer(OBJ(p)).GoEIDVariables(ToList(OBJ(self)) )} 
  
// check the range & sort of the method through type inference. 
// returns true if OK and false otherwise (can produce an error at run-time)
// notice that %body is the lambda body before compilation => use c_type
/* {1} OPT.The go function for: check_range(self:method,%body:any) [] */
func F_Generate_check_range_method (self *ClaireMethod ,_Zbody *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var s1 *ClaireClass   = self.Range.Class_I()
      /* noccur = 8 */
      /* Let:3 */{ 
        var s2 *ClaireClass  
        /* noccur = 6 */
        var s2_try01894 EID 
        /* Let:4 */{ 
          var g0190UU *ClaireType  
          /* noccur = 1 */
          var g0190UU_try01915 EID 
          g0190UU_try01915 = Core.F_CALL(Optimize.C_c_type,ARGS(_Zbody.ToEID()))
          /* ERROR PROTECTION INSERTED (g0190UU-s2_try01894) */
          if ErrorIn(g0190UU_try01915) {s2_try01894 = g0190UU_try01915
          } else {
          g0190UU = ToType(OBJ(g0190UU_try01915))
          s2_try01894 = EID{g0190UU.Class_I().Id(),0}
          }
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (s2-Result) */
        if ErrorIn(s2_try01894) {Result = s2_try01894
        } else {
        s2 = ToClass(OBJ(s2_try01894))
        Core.F_tformat_string(MakeString("---- info: ~S's range was found to be ~S (vs. ~S) \n"),0,MakeConstantList(self.Id(),s2.Id(),s1.Id()))
        if ((s1.Id() == C_void.Id()) || 
            (ToType(s2.Id()).Included(ToType(s1.Id())) == CTRUE)) /* If:4 */{ 
          Result = EID{CTRUE.Id(),0}
          } else {
          Core.F_tformat_string(MakeString("---- note: ~S's range was found to be ~S (vs. ~S) \n"),0,MakeConstantList(self.Id(),s2.Id(),s1.Id()))
          if (((s1.Id() != C_void.Id()) && 
                ((s2.Id() == C_void.Id()) && 
                  (s1.Id() != C_error.Id()))) || 
              (Equal(Core.F__exp_type(ToType(s1.Id()),ToType(s2.Id())).Id(),CEMPTY.Id()) == CTRUE)) /* If:5 */{ 
            Result = Core.F_CALL(Optimize.C_Compile_Cerror,ARGS(EID{MakeString("[218] Sort error: Cannot compile ~S (~S cannot be ~S).").Id(),0},
              EID{self.Id(),0},
              EID{s1.Id(),0},
              EID{s2.Id(),0}))
            } else {
            Result = EID{CFALSE.Id(),0}
            /* If-5 */} 
          /* If-4 */} 
        }
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: check_range @ method (throw: true) 
func E_Generate_check_range_method (self EID,_Zbody EID) EID { 
    return /*(sm for check_range @ method= EID)*/ F_Generate_check_range_method(ToMethod(OBJ(self)),ANY(_Zbody) )} 
  
// generate the eval function associated to each self_eval method (type *any -> EID)
// EVAL_C(x *ClaireAny) EID {return ToC(x).SelfEval()}
/* {1} OPT.The go function for: generate_eval_function(self:lambda,m:method) [] */
func F_Generate_generate_eval_function_lambda (self *ClaireLambda ,m *ClaireMethod ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var c *ClaireClass   = Core.F_domain_I_restriction(ToRestriction(m.Id()))
      /* noccur = 6 */
      if (c.Id() != C_Variable.Id()) /* If:3 */{ 
        PRINC("\n// The EVAL go function for: ")
        Result = Core.F_print_any(c.Id())
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC(" \n")
        Result = EVOID
        }
        {
        PRINC("func EVAL_")
        c.Name.CPrinc()
        PRINC(" (x *ClaireAny) EID ")
        F_Generate_new_block_void()
        if (F_Generate_goMethod_ask_any(m.Id()) == CTRUE) /* If:4 */{ 
          PRINC(" return ")
          F_Generate_cast_class_class(c)
          PRINC("(x).SelfEval()")
          } else {
          PRINC(" return F_self_eval_")
          c.Name.CPrinc()
          PRINC("(")
          F_Generate_cast_class_class(c)
          PRINC("(x))")
          /* If-4 */} 
        F_Generate_close_block_void()
        Result = EVOID
        }
        } else {
        Result = EID{CFALSE.Id(),0}
        /* If-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: generate_eval_function @ lambda (throw: true) 
func E_Generate_generate_eval_function_lambda (self EID,m EID) EID { 
    return /*(sm for generate_eval_function @ lambda= EID)*/ F_Generate_generate_eval_function_lambda(ToLambda(OBJ(self)),ToMethod(OBJ(m)) )} 
  
// tells if a method needs debug instrumentation
/* {1} OPT.The go function for: need_debug?(m:any) [] */
func F_Generate_need_debug_ask_any (m *ClaireAny ) *ClaireBoolean  { 
    // procedure body with s =  
var Result *ClaireBoolean  
    if (C_method.Id() == m.Isa.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var g0192 *ClaireMethod   = ToMethod(m)
        /* noccur = 4 */
        /* Let:4 */{ 
          var p *ClaireProperty   = g0192.Selector
          /* noccur = 9 */
          Result = MakeBoolean((Optimize.C_compiler.Debug_ask.Memq(g0192.Module_I.Id()) == CTRUE) && 
          (Core.F_domain_I_restriction(ToRestriction(g0192.Id())).Id() != C_environment.Id()) && 
          (g0192.Module_I.Id() != C_claire.Id()) && 
          (p.Id() != Core.C_self_eval.Id()) && 
          (p.Id() != Core.C_execute.Id()) && 
          (p.Id() != Core.C_eval_message.Id()) && 
          (p.Id() != Core.C_Core_push_debug.Id()) && 
          (p.Id() != Core.C_Core_pop_debug.Id()) && 
          (p.Id() != Core.C_Core_tr_indent.Id()) && 
          (p.Id() != Core.C_Core_find_which.Id()) && 
          (p.Id() != Core.C_Core_matching_ask.Id()) && 
          (p.Id() != Core.C_Core_vmatch_ask.Id()))
          /* Let-4 */} 
        /* Let-3 */} 
      } else {
      Result = CFALSE
      /* If-2 */} 
    return Result} 
  
// The EID go function for: need_debug? @ any (throw: false) 
func E_Generate_need_debug_ask_any (m EID) EID { 
    return EID{/*(sm for need_debug? @ any= boolean)*/ F_Generate_need_debug_ask_any(ANY(m) ).Id(),0}} 
  
// profiler code 
/* {1} OPT.The go function for: generate_profile(c:go_producer,m:any) [] */
func (c *GenerateGoProducer ) GenerateProfile (m *ClaireAny ) EID { 
    var Result EID 
    if (C_method.Id() == m.Isa.Id()) /* If:2 */{ 
      Result = F_Generate_get_dependents_method(ToMethod(m))
      } else {
      Result = EID{CFALSE.Id(),0}
      /* If-2 */} 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("   var PR_x *ClairePRcount PRget_property(")
    /* Let:2 */{ 
      var g0196UU *ClaireProperty  
      /* noccur = 1 */
      if (C_method.Id() == m.Isa.Id()) /* If:3 */{ 
        /* Let:4 */{ 
          var g0194 *ClaireMethod   = ToMethod(m)
          /* noccur = 1 */
          g0196UU = g0194.Selector
          /* Let-4 */} 
        } else {
        g0196UU = C_fastcall
        /* If-3 */} 
      Core.F_CALL(C_Generate_expression,ARGS(EID{g0196UU.Id(),0},EID{CNIL.Id(),0}))
      /* Let-2 */} 
    PRINC(").Start();")
    F_Generate_breakline_void()
    PRINC("")
    Result = EVOID
    }
    return Result} 
  
// The EID go function for: generate_profile @ go_producer (throw: true) 
func E_Generate_generate_profile_go_producer (c EID,m EID) EID { 
    return /*(sm for generate_profile @ go_producer= EID)*/ ToGenerateGoProducer(OBJ(c)).GenerateProfile(ANY(m) )} 
  
// produce the debugging code introduction
// in go we have no macros but functions with variable number of args
// assumes DebugBind(module, method, ClaireAny args* ) 
/* {1} OPT.The go function for: debug_intro(c:go_producer,self:lambda,x:method) [] */
func (c *GenerateGoProducer ) DebugIntro (self *ClaireLambda ,x *ClaireMethod ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var m *ClaireModule  
      /* noccur = 1 */
      if (C_method.Id() == x.Isa.Id()) /* If:3 */{ 
        m = x.Module_I
        } else {
        m = ToModule(CFALSE.Id())
        /* If-3 */} 
      /* Let:3 */{ 
        var n int  = 1
        /* noccur = 3 */
        PRINC("DebugBind(")
        F_iClaire_ident_thing(ToThing(m.Id()))
        PRINC(",")
        Core.F_CALL(C_Generate_expression,ARGS(EID{x.Selector.Id(),0},EID{CEMPTY.Id(),0}))
        if ((self.Vars.Length() == 1) && 
            (ANY(Core.F_CALL(C_range,ARGS(self.Vars.At(1-1).ToEID()))) == C_void.Id())) /* If:4 */{ 
          PRINC(",EID{C_object,ClEnv.Uip()}));")
          Result = EVOID
          } else {
          /* For:5 */{ 
            var v *ClaireAny  
            _ = v
            Result= EID{CFALSE.Id(),0}
            var v_support *ClaireList  
            v_support = self.Vars
            v_len := v_support.Length()
            for i_it := 0; i_it < v_len; i_it++ { 
              v = v_support.At(i_it)
              var void_try7 EID 
              _ = void_try7
              { 
              PRINC(", ")
              void_try7 = c.ToEid(v,ToClass(F_Generate_go_signature_method(x).ValuesO()[n-1]))
              /* ERROR PROTECTION INSERTED (void_try7-void_try7) */
              if ErrorIn(void_try7) {Result = void_try7
              break
              } else {
              PRINC("")
              void_try7 = EVOID
              }
              {
              n = (n+1)
              }
              }
              /* loop-6 */} 
            /* For-5 */} 
          /* If-4 */} 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC(");")
        F_Generate_breakline_void()
        PRINC("")
        Result = EVOID
        }
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: debug_intro @ go_producer (throw: true) 
func E_Generate_debug_intro_go_producer (c EID,self EID,x EID) EID { 
    return /*(sm for debug_intro @ go_producer= EID)*/ ToGenerateGoProducer(OBJ(c)).DebugIntro(ToLambda(OBJ(self)),ToMethod(OBJ(x)) )} 
  
// auxiliary to produce the end statement for the function. s tells if the result is needed.
// generates a "... return" if the result is needed or just an empy string
// we also add the debugging unbind if needed.  (used to be called protect_result)
/* {1} OPT.The go function for: return_result(p:go_producer,s:class,x:method,%res:string) [] */
func (p *GenerateGoProducer ) ReturnResult (s *ClaireClass ,x *ClaireMethod ,_Zres *ClaireString ) EID { 
    var Result EID 
    if (F_Generate_need_debug_ask_any(x.Id()) == CTRUE) /* If:2 */{ 
      PRINC("DebugUnbind(")
      Result = Core.F_CALL(Language.C_iClaire_ident,ARGS(EID{x.Id(),0}))
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(",")
      Result = F_Generate_g_expression_thing(ToThing(x.Selector.Id()),C_property)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(",")
      /* Let:3 */{ 
        var g0200UU *ClaireAny  
        /* noccur = 1 */
        if (s.Id() == C_void.Id()) /* If:4 */{ 
          g0200UU = CNULL
          } else {
          g0200UU = F_Generate_build_Variable_string(MakeString("Result"),s.Id()).Id()
          /* If-4 */} 
        Result = p.ToEid(g0200UU,C_object)
        /* Let-3 */} 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(")")
      F_Generate_breakline_void()
      PRINC("")
      Result = EVOID
      }}}
      /* If!2 */}  else if (Optimize.C_OPT.Profile_ask == CTRUE) /* If:2 */{ 
      PRINC("PRend(PR_x)")
      F_Generate_breakline_void()
      PRINC("")
      Result = EVOID
      } else {
      Result = EID{CFALSE.Id(),0}
      /* If-2 */} 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    if (s.Id() != C_void.Id()) /* If:2 */{ 
      PRINC("return ")
      F_princ_string(_Zres)
      PRINC("")
      Result = EVOID
      } else {
      Result = EID{CFALSE.Id(),0}
      /* If-2 */} 
    }
    return Result} 
  
// The EID go function for: return_result @ go_producer (throw: true) 
func E_Generate_return_result_go_producer (p EID,s EID,x EID,_Zres EID) EID { 
    return /*(sm for return_result @ go_producer= EID)*/ ToGenerateGoProducer(OBJ(p)).ReturnResult(ToClass(OBJ(s)),
      ToMethod(OBJ(x)),
      ToString(OBJ(_Zres)) )} 
  
// computes the inter-module dependence
/* {1} OPT.The go function for: get_dependents(m:method) [] */
func F_Generate_get_dependents_method (m *ClaireMethod ) EID { 
    var Result EID 
    /* For:2 */{ 
      var p *ClaireAny  
      _ = p
      Result= EID{CFALSE.Id(),0}
      var p_support *ClaireSet  
      var p_support_try02023 EID 
      p_support_try02023 = Reader.F_dependents_method(m)
      /* ERROR PROTECTION INSERTED (p_support-Result) */
      if ErrorIn(p_support_try02023) {Result = p_support_try02023
      } else {
      p_support = ToSet(OBJ(p_support_try02023))
      for _,p = range(p_support.Values)/* loop2:3 */{ 
        var void_try4 EID 
        _ = void_try4
        /* Let:4 */{ 
          var g0201 *ClaireProperty   = m.Selector
          /* noccur = 1 */
          void_try4 = Core.F_add_table(Reader.C_Reader_PRdependent,g0201.Id(),p)
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (void_try4-Result) */
        if ErrorIn(void_try4) {Result = void_try4
        Result = void_try4
        break
        } else {
        }}
        /* loop-3 */} 
      /* For-2 */} 
    return Result} 
  
// The EID go function for: get_dependents @ method (throw: true) 
func E_Generate_get_dependents_method (m EID) EID { 
    return /*(sm for get_dependents @ method= EID)*/ F_Generate_get_dependents_method(ToMethod(OBJ(m)) )} 
  
// prints a function name without the # syntactic marker for imported
/* {1} OPT.The go function for: c_princ(self:function) [] */
func F_c_princ_function (self *ClaireFunction )  { 
    // procedure body with s =  
F_Generate_import_princ_string(F_string_I_function(self))
    } 
  
// The EID go function for: c_princ @ function (throw: false) 
func E_c_princ_function (self EID) EID { 
    /*(sm for c_princ @ function= void)*/ F_c_princ_function(ToFunction(OBJ(self)) )
    return EVOID} 
  
/* {1} OPT.The go function for: import_princ(s:string) [] */
func F_Generate_import_princ_string (s *ClaireString )  { 
    // procedure body with s =  
/* Let:2 */{ 
      var i int  = 1
      /* noccur = 6 */
      /* Let:3 */{ 
        var g0203 int  = F_length_string(s)
        /* noccur = 1 */
        for (i <= g0203) /* while:4 */{ 
          if ((i > 1) || 
              (s.At(i) != '#')) /* If:5 */{ 
            F_c_princ_char(s.At(i))
            /* If-5 */} 
          i = (i+1)
          /* while-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    } 
  
// The EID go function for: import_princ @ string (throw: false) 
func E_Generate_import_princ_string (s EID) EID { 
    /*(sm for import_princ @ string= void)*/ F_Generate_import_princ_string(ToString(OBJ(s)) )
    return EVOID} 
  
// v3.2.06 - some properties may be extended
//(put(open,Generate/set_outfile,4),
// put(open,Generate/inline_exp,4))
// end of file