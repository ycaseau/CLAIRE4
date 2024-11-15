/***** CLAIRE Compilation of file /Users/ycaseau/Dropbox/src/clairev4.12/src/compile/gosystem.cl 
         [version 4.1.2 / safety 5] Sunday 08-11-2024 08:17:31 *****/

package Generate
import (_ "fmt"
	. "Kernel"
	"Core"
	"Language"
	"Reader"
	"Optimize"
)

//-------- dumb function to prevent import errors --------
func import_g0000() { 
  _ = Core.It
  _ = Language.It
  _ = Reader.It
  _ = Optimize.It
  } 


//+-------------------------------------------------------------+
//| CLAIRE                                                      |
//| gosystem.cl                                                 |
//| Copyright (C) 2020-2023 Yves Caseau. All Rights Reserved    |
//| cf. copyright info in file object.cl: about()               |
//+-------------------------------------------------------------+
//**********************************************************************
//* Contents                                                           *
//*          Part 1: Global_variables & producer interface             *
//*          Part 2: Generic Compiler Methods                          *
//*          Part 3: Module Compiler Interface                         *
//*          Part 4: File Compiler                                     *
//*          Part 5: Function Compiler                                 *
//**********************************************************************
// content map (represent the tree with a indented hierarchy :))
// compile [Part 2]
//      - gen_files
//            - gen_file
//      - gen_module_file
//            - start_file
//            - gen_objects
//            - gen_classes
//  compile_lambda -> ... -> gen_function [Part 4]
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
// v3.2.58  */
// force goMethod()
// new: the target code production (the part that depends on the target language) is
// encapsulated with a producer object
// CLAIRE 4 is focused on go, but we try to keep the previous structure of CLAIRE3 to be ready
// for Java or Swift compiling. However, the GC management stuff is lost forever :)
// used to translate imported to C/.. entities
// add the go_producer here  (replaces the C++ producer)
// note that the double list bad/good names is ugly and should be replaced by a dictionary later 
// where to place the executable
// this is a special case : the function may return an error but the optimized form does not
// most standard method: call the producer to print the ident from a symbol
/* The go function for: iClaire/ident(self:symbol) [status=0] */
func F_iClaire_ident_symbol (self *ClaireSymbol)  { 
  F_iClaire_ident_go_producer2(ToGenerateGoProducer(Optimize.C_PRODUCER.Value),self)
  } 

// The EID go function for: iClaire/ident @ symbol (throw: false) 
func E_iClaire_ident_symbol (self EID) EID { 
  F_iClaire_ident_symbol(ToSymbol(OBJ(self)) )
  return EVOID} 

/* The go function for: iClaire/ident(self:thing) [status=0] */
func F_iClaire_ident_thing (self *ClaireThing)  { 
  F_iClaire_ident_go_producer2(ToGenerateGoProducer(Optimize.C_PRODUCER.Value),self.Name)
  } 

// The EID go function for: iClaire/ident @ thing (throw: false) 
func E_iClaire_ident_thing (self EID) EID { 
  F_iClaire_ident_thing(ToThing(OBJ(self)) )
  return EVOID} 

/* The go function for: iClaire/ident(self:class) [status=0] */
func F_iClaire_ident_class (self *ClaireClass)  { 
  F_iClaire_ident_go_producer2(ToGenerateGoProducer(Optimize.C_PRODUCER.Value),self.Name)
  } 

// The EID go function for: iClaire/ident @ class (throw: false) 
func E_iClaire_ident_class (self EID) EID { 
  F_iClaire_ident_class(ToClass(OBJ(self)) )
  return EVOID} 

// we simply use some smart indentation. True pretty_printing will be left to bc
/* The go function for: indent_c(_CL_obj:void) [status=0] */
func F_Generate_indent_c_void () *ClaireAny { 
  var Result *ClaireAny
  { var x int = Optimize.C_OPT.Level
    Result= CFALSE.Id()
    for (x > 0) { 
      PRINC("  ")
      x = (x-1)
      } 
    } 
  return Result} 

// The EID go function for: indent_c @ void (throw: false) 
func E_Generate_indent_c_void (_CL_obj EID) EID { 
  return F_Generate_indent_c_void( ).ToEID()} 

/* The go function for: breakline(_CL_obj:void) [status=0] */
func F_Generate_breakline_void () *ClaireAny { 
  PRINC("\n")
  return  F_Generate_indent_c_void()
  } 

// The EID go function for: breakline @ void (throw: false) 
func E_Generate_breakline_void (_CL_obj EID) EID { 
  return F_Generate_breakline_void( ).ToEID()} 

// adds a new C block with the condensed option
/* The go function for: new_block(_CL_obj:void) [status=0] */
func F_Generate_new_block_void ()  { 
  Optimize.C_OPT.Level = (Optimize.C_OPT.Level+1)
  PRINC("{ ")
  F_Generate_breakline_void()
  } 

// The EID go function for: new_block @ void (throw: false) 
func E_Generate_new_block_void (_CL_obj EID) EID { 
  F_Generate_new_block_void( )
  return EVOID} 

// adds a new block without the breaklines for Let
/* The go function for: let_block(_CL_obj:void) [status=0] */
func F_Generate_let_block_void ()  { 
  Optimize.C_OPT.Level = (Optimize.C_OPT.Level+1)
  PRINC("{ ")
  } 

// The EID go function for: let_block @ void (throw: false) 
func E_Generate_let_block_void (_CL_obj EID) EID { 
  F_Generate_let_block_void( )
  return EVOID} 

// closes the current C block
/* The go function for: close_block(_CL_obj:void) [status=0] */
func F_Generate_close_block_void ()  { 
  Optimize.C_OPT.Level = (Optimize.C_OPT.Level-1)
  PRINC("} ")
  F_Generate_breakline_void()
  } 

// The EID go function for: close_block @ void (throw: false) 
func E_Generate_close_block_void (_CL_obj EID) EID { 
  F_Generate_close_block_void( )
  return EVOID} 

// prints the } without a new line - used for nested If
/* The go function for: finish_block(_CL_obj:void) [status=0] */
func F_Generate_finish_block_void ()  { 
  Optimize.C_OPT.Level = (Optimize.C_OPT.Level-1)
  PRINC("} ")
  } 

// The EID go function for: finish_block @ void (throw: false) 
func E_Generate_finish_block_void (_CL_obj EID) EID { 
  F_Generate_finish_block_void( )
  return EVOID} 

//*********************************************************************
//*          Part 2: Generic Compiler Methods                         *
//*********************************************************************
// these are the code_producer methods that works with go, C++ or Javascript
// compiling.
// this this the heart of the compiler : compiles module m into a set of files
/* The go function for: compile(p:code_producer,m:module) [status=1] */
func (p *GenerateCodeProducer) Compile (m *ClaireModule) EID { 
  var Result EID
  Optimize.C_OPT.NeedModules = CEMPTY
  Optimize.C_compiler.Inline_ask = CTRUE
  Optimize.C_compiler.NLoc = 0
  Optimize.C_compiler.NWarnings = 0
  Optimize.C_compiler.NNotes = 0
  { var l1 *ClaireBag = ToBag(F_Generate_parents_list(Reader.F_add_modules_list(MakeConstantList(m.Id()))).Id())
    Core.F_tformat_string(MakeString("==========  START GO COMPILING (~S) with ~S ================ \n"),3,MakeConstantList(m.Id(),l1.Id()))
    { 
      var va_arg1 *Optimize.OptimizeMetaOPT
      var va_arg2 *ClaireSet
      va_arg1 = Optimize.C_OPT
      var try_1 EID
      try_1 = Core.F_CALL(C_set_I,ARGS(EID{l1.Id(),0}))
      if ErrorIn(try_1) {Result = try_1
      } else {
      va_arg2 = ToSet(OBJ(try_1))
      va_arg1.LegalModules = va_arg2
      Result = EID{va_arg2.Id(),0}
      }
      } 
    if !ErrorIn(Result) {
    ToGenerateGoProducer(p.Id()).Current = m
    Result = ToGenerateGoProducer(p.Id()).GenFiles(m)
    if !ErrorIn(Result) {
    Result = Core.F_CALL(C_Generate_gen_mod_file,ARGS(EID{p.Id(),0},EID{m.Id(),0}))
    if !ErrorIn(Result) {
    l1 = ToBag(Core.F_difference_set(Core.F_set_I_set(Optimize.C_OPT.NeedModules),Optimize.C_OPT.LegalModules).Id())
    if (F_boolean_I_any(l1.Id()) == CTRUE) { 
      Optimize.F_Compile_warn_void()
      Result = Core.F_tformat_string(MakeString("~S should be declared for ~S \n"),1,MakeConstantList(l1.Id(),m.Id()))
      } else {
      Result = EID{CFALSE.Id(),0}
      } 
    }}}
    } 
  return Result} 

// The EID go function for: compile @ code_producer (throw: true) 
func E_compile_code_producer (p EID,m EID) EID { 
  return ToGenerateCodeProducer(OBJ(p)).Compile(ToModule(OBJ(m)) )} 

// the first part is to generate the files associated to the module claire files
// gen_files prepare the optimizer (OPT) and calls gen_file for each file
/* The go function for: gen_files(p:go_producer,m:module) [status=1] */
func (p *GenerateGoProducer) GenFiles (m *ClaireModule) EID { 
  var Result EID
  Core.F_tformat_string(MakeString("==== Generate ~A files for module ~S [verbose = ~A, Opt? = ~S] \n"),0,MakeConstantList((p.Comment).Id(),
    m.Id(),
    MakeInteger(ClEnv.Verbose).Id(),
    Optimize.C_compiler.Optimize_ask.Id()))
  Optimize.C_OPT.Instructions = ToType(C_any.Id()).EmptyList()
  Optimize.C_OPT.Properties = ToType(C_property.Id()).EmptySet()
  Optimize.C_OPT.Objects = ToType(C_object.Id()).EmptyList()
  Optimize.C_OPT.Functions = ToType(C_any.Id()).EmptyList()
  Optimize.C_OPT.NeedToClose = ToType(C_any.Id()).EmptySet()
  m.Begin()
  { 
    var x *ClaireAny
    _ = x
    Result= EID{CFALSE.Id(),0}
    var x_support *ClaireList
    x_support = m.MadeOf
    x_len := x_support.Length()
    for i_it := 0; i_it < x_len; i_it++ { 
      x = x_support.At(i_it)
      var loop_1 EID
      _ = loop_1
      { 
      Core.F_tformat_string(MakeString("++++ Compiling the file ~A.cl [v. 4.~A - safety:~A] \n"),1,MakeConstantList(x,Optimize.C_compiler.Version,MakeInteger(Optimize.C_compiler.Safety).Id()))
      if (Equal(x,(m.Name.String_I()).Id()) == CTRUE) { 
        loop_1 = Optimize.F_Compile_Cerror_string(MakeString("[211]  ~S cannot be used both as a file and module name"),ToList(MakeConstantList(x).Id()))
        } else {
        loop_1 = EID{CFALSE.Id(),0}
        } 
      if ErrorIn(loop_1) {Result = loop_1
      break
      } else {
      p.CurrentFile = ToString(x)
      loop_1 = p.GenFile(Reader.F__7_string(m.Source,ToString(x)),Reader.F__7_string(p.Source,ToString(x)))
      if ErrorIn(loop_1) {Result = loop_1
      break
      } else {
      }}
      }
      } 
    } 
  if !ErrorIn(Result) {
  m.End()
  Result = EVOID
  }
  return Result} 

// The EID go function for: gen_files @ go_producer (throw: true) 
func E_Generate_gen_files_go_producer2 (p EID,m EID) EID { 
  return ToGenerateGoProducer(OBJ(p)).GenFiles(ToModule(OBJ(m)) )} 

// main method: generate a target file associated to a CLAIRE file
/* The go function for: gen_file(p:code_producer,f1:string,f2:string) [status=1] */
func (p *GenerateCodeProducer) GenFile (f1 *ClaireString,f2 *ClaireString) EID { 
  var Result EID
  { var p1 *ClairePort
    var try_1 EID
    try_1 = F_fopen_string(F_append_string(f1,MakeString(".cl")),MakeString("r"))
    if ErrorIn(try_1) {Result = try_1
    } else {
    p1 = ToPort(OBJ(try_1))
    { var b *ClaireBoolean = Reader.C_reader.Toplevel
      { var p0 *ClairePort = Reader.C_reader.Fromp
        { var out_ask *ClaireBoolean
          var try_2 EID
          try_2 = Core.F_CALL(C_Generate_outfile_ask,ARGS(EID{p.Id(),0}))
          if ErrorIn(try_2) {Result = try_2
          } else {
          out_ask = ToBoolean(OBJ(try_2))
          { var prev_comment *ClaireString = MakeString("")
            if (out_ask == CTRUE) { 
              { 
                var va_arg1 *Optimize.OptimizeMetaOPT
                var va_arg2 *ClairePort
                va_arg1 = Optimize.C_OPT
                var try_3 EID
                try_3 = F_fopen_string(F_append_string(f2,p.Extension),MakeString("w"))
                if ErrorIn(try_3) {Result = try_3
                } else {
                va_arg2 = ToPort(OBJ(try_3))
                va_arg1.Outfile = va_arg2
                Result = va_arg2.ToEID()
                }
                } 
              } else {
              Result = EID{CFALSE.Id(),0}
              } 
            if !ErrorIn(Result) {
            Reader.C_reader.Toplevel = CFALSE
            Optimize.C_compiler.Loading_ask = CTRUE
            ClEnv.NLine = 1
            Reader.C_reader.External = f1
            Reader.C_reader.Fromp = p1
            if (out_ask == CTRUE) { 
              Result = Core.F_CALL(C_Generate_start_file,ARGS(EID{p.Id(),0},
                EID{(f1).Id(),0},
                EID{ClEnv.Module_I.Id(),0},
                EID{CFALSE.Id(),0}))
              } else {
              Result = EID{CFALSE.Id(),0}
              } 
            if !ErrorIn(Result) {
            { var _Zinstruction *ClaireAny
              var try_4 EID
              try_4 = Reader.F_readblock_port(p1)
              if ErrorIn(try_4) {Result = try_4
              } else {
              _Zinstruction = ANY(try_4)
              Result= EID{CFALSE.Id(),0}
              for (_Zinstruction != Reader.C_Reader_eof.Id()) { 
                var loop_5 EID
                _ = loop_5
                { 
                var try_6 EID
                try_6 = Core.F_CALL(C_Generate_gen_instruction,ARGS(EID{p.Id(),0},_Zinstruction.ToEID(),EID{(prev_comment).Id(),0}))
                if ErrorIn(try_6) {loop_5 = try_6
                Result = try_6
                break
                } else {
                prev_comment = ToString(OBJ(try_6))
                loop_5 = EID{prev_comment.Id(),0}
                var try_7 EID
                try_7 = Reader.F_readblock_port(p1)
                if ErrorIn(try_7) {loop_5 = try_7
                Result = try_7
                break
                } else {
                _Zinstruction = ANY(try_7)
                loop_5 = _Zinstruction.ToEID()
                }}
                } 
              }
              }
              } 
            if !ErrorIn(Result) {
            Optimize.C_compiler.NLoc = (Optimize.C_compiler.NLoc+ClEnv.NLine)
            p1.Fclose()
            Optimize.C_compiler.Loading_ask = CFALSE
            Reader.C_reader.Toplevel = b
            Reader.C_reader.External = MakeString("toplevel")
            Reader.C_reader.Fromp = p0
            if (out_ask == CTRUE) { 
              Optimize.C_OPT.Outfile.Fclose()
              Result = EVOID
              } else {
              Result = EID{CFALSE.Id(),0}
              } 
            }}}
            } 
          }
          } 
        } 
      } 
    }
    } 
  return Result} 

// The EID go function for: gen_file @ code_producer (throw: true) 
func E_Generate_gen_file_code_producer (p EID,f1 EID,f2 EID) EID { 
  return ToGenerateCodeProducer(OBJ(p)).GenFile(ToString(OBJ(f1)),ToString(OBJ(f2)) )} 

// generate a file associated with module m (called meta-m)
/* The go function for: gen_mod_file(p:code_producer,m:module) [status=1] */
func (p *GenerateCodeProducer) GenModFile (m *ClaireModule) EID { 
  var Result EID
  { var prt *ClairePort
    var try_1 EID
    { var arg_2 *ClaireString
      var try_3 EID
      { var arg_4 *ClaireString
        var try_5 EID
        { var arg_6 *ClaireAny
          var try_7 EID
          try_7 = Core.F_CALL(C_Generate_modfile_name,ARGS(EID{p.Id(),0},EID{m.Id(),0}))
          if ErrorIn(try_7) {try_5 = try_7
          } else {
          arg_6 = ANY(try_7)
          try_5 = EID{F_append_string(F_append_string(ToGenerateGoProducer(p.Id()).Source,ToString(Reader.C__starfs_star.Value)),ToString(arg_6)).Id(),0}
          }
          } 
        if ErrorIn(try_5) {try_3 = try_5
        } else {
        arg_4 = ToString(OBJ(try_5))
        try_3 = EID{F_append_string(arg_4,p.Extension).Id(),0}
        }
        } 
      if ErrorIn(try_3) {try_1 = try_3
      } else {
      arg_2 = ToString(OBJ(try_3))
      try_1 = F_fopen_string(arg_2,MakeString("w"))
      }
      } 
    if ErrorIn(try_1) {Result = try_1
    } else {
    prt = ToPort(OBJ(try_1))
    { var s *ClaireString = m.Name.String_I()
      { var arg_8 *ClaireList
        var try_9 EID
        { 
          var v_bag_arg *ClaireAny
          try_9= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
          ToList(OBJ(try_9)).AddFast((p.Comment).Id())
          ToList(OBJ(try_9)).AddFast(m.Id())
          var try_10 EID
          try_10 = Core.F_CALL(C_Generate_modfile_name,ARGS(EID{p.Id(),0},EID{m.Id(),0}))
          if ErrorIn(try_10) {try_9 = try_10
          } else {
          v_bag_arg = ANY(try_10)
          ToList(OBJ(try_9)).AddFast(v_bag_arg)}
          } 
        if ErrorIn(try_9) {Result = try_9
        } else {
        arg_8 = ToList(OBJ(try_9))
        Result = Core.F_tformat_string(MakeString("==== generate ~A file for module ~S [~A] ==== \n"),0,arg_8)
        }
        } 
      if !ErrorIn(Result) {
      Optimize.C_OPT.Outfile = prt
      Result = Core.F_CALL(C_Generate_start_file,ARGS(EID{p.Id(),0},
        EID{(s).Id(),0},
        EID{m.Id(),0},
        EID{CTRUE.Id(),0}))
      if !ErrorIn(Result) {
      Optimize.C_OPT.Outfile.UseAsOutput()
      Result = Core.F_CALL(C_Generate_gen_classes,ARGS(EID{p.Id(),0},EID{m.Id(),0}))
      if !ErrorIn(Result) {
      Result = Core.F_CALL(C_Generate_gen_objects,ARGS(EID{p.Id(),0},EID{m.Id(),0}))
      if !ErrorIn(Result) {
      Result = Core.F_CALL(C_Generate_gen_functions,ARGS(EID{p.Id(),0},EID{m.Id(),0}))
      if !ErrorIn(Result) {
      Result = Core.F_CALL(C_Generate_gen_meta_load,ARGS(EID{p.Id(),0},EID{m.Id(),0}))
      if !ErrorIn(Result) {
      F_Generate_breakline_void()
      Optimize.C_OPT.Outfile.Fclose()
      Result = EVOID
      }}}}}}
      } 
    }
    } 
  return Result} 

// The EID go function for: gen_mod_file @ list<type_expression>(code_producer, module) (throw: true) 
func E_Generate_gen_mod_file_code_producer2 (p EID,m EID) EID { 
  return ToGenerateCodeProducer(OBJ(p)).GenModFile(ToModule(OBJ(m)) )} 

// test the compiling of a method
// e.f. g_test(foo @ any)
/* The go function for: compile_method(p:code_producer,m:method) [status=1] */
func (p *GenerateCodeProducer) CompileMethod (m *ClaireMethod) EID { 
  var Result EID
  { var l *ClaireLambda = m.Formula
    if (l.Id() == CNULL) { 
      Result = EID{CNULL,0}
      } else {
      Core.F_tformat_string(MakeString("---- Compiling ~S with following definition ---- \n"),0,MakeConstantList(m.Id()))
      Result = Language.F_pretty_print_any(l.Body)
      if !ErrorIn(Result) {
      Optimize.C_OPT.InMethod = m.Id()
      Optimize.C_OPT.UseStringUpdate = CFALSE
      Optimize.C_OPT.MaxVars = 0
      Optimize.C_OPT.LegalModules = C_module.Instances.Set_I()
      Optimize.C_OPT.Outfile = ToPort(Reader.C_stdout.Value)
      Optimize.C_compiler.Inline_ask = CTRUE
      ToGenerateGoProducer(p.Id()).Current = C_claire
      Core.F_tformat_string(MakeString("\n---- code produced by the optimizer -------------------\n"),0,ToType(CEMPTY.Id()).EmptyList())
      { var arg_1 *ClaireAny
        var try_2 EID
        try_2 = Optimize.F_Compile_c_strict_code_any(m.Formula.Body,m.Range.Class_I())
        if ErrorIn(try_2) {Result = try_2
        } else {
        arg_1 = ANY(try_2)
        Result = Language.F_pretty_print_any(arg_1)
        }
        } 
      if !ErrorIn(Result) {
      Core.F_tformat_string(MakeString("\n---- code produced by the generator ------------------- \n"),0,ToType(CEMPTY.Id()).EmptyList())
      Result = Core.F_CALL(C_Generate_gen_function,ARGS(EID{p.Id(),0},
        EID{m.Formula.Id(),0},
        EID{MakeString("test").Id(),0},
        EID{m.Id(),0}))
      if !ErrorIn(Result) {
      { 
        var va_arg1 *Optimize.OptimizeMetaOPT
        var va_arg2 *ClaireAny
        va_arg1 = Optimize.C_OPT
        va_arg2 = CNULL
        va_arg1.InMethod = va_arg2
        Result = va_arg2.ToEID()
        } 
      }}}
      } 
    } 
  return Result} 

// The EID go function for: compile_method @ code_producer (throw: true) 
func E_Generate_compile_method_code_producer (p EID,m EID) EID { 
  return ToGenerateCodeProducer(OBJ(p)).CompileMethod(ToMethod(OBJ(m)) )} 

// these are the open properties that CLAIRE compiler expects (specific to the target language)
// open for future producers (such as Javascript)
//*********************************************************************
//*          Part 3: Module Compiler Interface                        *
//*********************************************************************
// a small g_test function for the compiler : very convenient to test the compiler
/* The go function for: g_test(x:any) [status=1] */
func F_g_test_any (x *ClaireAny) EID { 
  var Result EID
  Result = F_g_test_module(C_claire,x)
  return Result} 

// The EID go function for: g_test @ any (throw: true) 
func E_g_test_any (x EID) EID { 
  return F_g_test_any(ANY(x) )} 

/* The go function for: g_test(m:module,x:any) [status=1] */
func F_g_test_module (m *ClaireModule,x *ClaireAny) EID { 
  var Result EID
  { var t *ClaireType
    var try_1 EID
    try_1 = Core.F_CALL(Optimize.C_c_type,ARGS(x.ToEID()))
    if ErrorIn(try_1) {Result = try_1
    } else {
    t = ToType(OBJ(try_1))
    { var s *ClaireClass = Optimize.F_Compile_osort_any(t.Id())
      { var u *ClaireAny
        var try_2 EID
        try_2 = Core.F_CALL(Optimize.C_c_code,ARGS(x.ToEID(),EID{s.Id(),0}))
        if ErrorIn(try_2) {Result = try_2
        } else {
        u = ANY(try_2)
        { var f *ClaireBoolean
          var try_3 EID
          try_3 = F_Generate_g_func_any(u)
          if ErrorIn(try_3) {Result = try_3
          } else {
          f = ToBoolean(OBJ(try_3))
          { var gt *ClaireBoolean
            var try_4 EID
            try_4 = Optimize.F_Compile_g_throw_any(u)
            if ErrorIn(try_4) {Result = try_4
            } else {
            gt = ToBoolean(OBJ(try_4))
            ToGenerateGoProducer(Optimize.C_PRODUCER.Value).Current = m
            PRINC("type -> ")
            Result = Core.F_print_any(t.Id())
            if !ErrorIn(Result) {
            PRINC(" [sort ")
            Result = Core.F_print_any(s.Id())
            if !ErrorIn(Result) {
            PRINC("]\n")
            Result = EVOID
            }}
            if !ErrorIn(Result) {
            PRINC("opt[")
            Result = Core.F_print_any(u.Isa.Id())
            if !ErrorIn(Result) {
            PRINC("] -> ")
            Result = Core.F_CALL(C_print,ARGS(u.ToEID()))
            if !ErrorIn(Result) {
            PRINC(" \n")
            Result = EVOID
            }}
            if !ErrorIn(Result) {
            if (gt == CTRUE) { 
              PRINC("----------------------- Error is possible => EID (func:")
              Result = Core.F_print_any(f.Id())
              if !ErrorIn(Result) {
              PRINC(")  ----------------\n")
              Result = EVOID
              }
              } else {
              Result = EID{CFALSE.Id(),0}
              } 
            if !ErrorIn(Result) {
            if (f == CTRUE) { 
              PRINC("exp  -> ")
              Result = Core.F_CALL(C_Generate_g_expression,ARGS(u.ToEID(),EID{t.Class_I().Id(),0}))
              if !ErrorIn(Result) {
              PRINC("\n")
              Result = EVOID
              }
              } else {
              PRINC("stat -> ")
              { var arg_5 *ClaireClass
                if (gt == CTRUE) { 
                  arg_5 = Optimize.C_EID
                  } else {
                  arg_5 = t.Class_I()
                  } 
                Result = F_Generate_statement_any(u,arg_5,MakeString("result"),CFALSE.Id())
                } 
              if !ErrorIn(Result) {
              PRINC("\n")
              Result = EVOID
              }
              } 
            }}}
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

// The EID go function for: g_test @ module (throw: true) 
func E_g_test_module (m EID,x EID) EID { 
  return F_g_test_module(ToModule(OBJ(m)),ANY(x) )} 

// even more fun 
/* The go function for: gtop(_CL_obj:void) [status=1] */
func F_Generate_gtop_void () EID { 
  var Result EID
  PRINC("in> ")
  { var x *ClaireAny
    var try_1 EID
    try_1 = Reader.F_read_port(ToPort(Reader.C_stdin.Value))
    if ErrorIn(try_1) {Result = try_1
    } else {
    x = ANY(try_1)
    if (x == Reader.C_q.Id()) { 
      PRINC("bye.\n")
      Result = EVOID
      } else {
      Result = Core.F_CALL(C_g_test,ARGS(x.ToEID()))
      if !ErrorIn(Result) {
      Result = F_Generate_gtop_void()
      }
      } 
    }
    } 
  return Result} 

// The EID go function for: gtop @ void (throw: true) 
func E_Generate_gtop_void (_CL_obj EID) EID { 
  return F_Generate_gtop_void( )} 

// test the compiling of a method
// e.f. g_test(foo @ any)
/* The go function for: g_test(m:method) [status=1] */
func F_g_test_method (m *ClaireMethod) EID { 
  var Result EID
  Result = ToGenerateCodeProducer(Optimize.C_PRODUCER.Value).CompileMethod(m)
  return Result} 

// The EID go function for: g_test @ method (throw: true) 
func E_g_test_method (m EID) EID { 
  return F_g_test_method(ToMethod(OBJ(m)) )} 

// debug (to remove later)
/* The go function for: g_test(l:lambda) [status=1] */
func F_g_test_lambda (l *ClaireLambda) EID { 
  var Result EID
  Optimize.C_OPT.InMethod = CNULL
  Optimize.C_OPT.UseStringUpdate = CFALSE
  Optimize.C_OPT.MaxVars = 0
  Optimize.C_OPT.LegalModules = C_module.Instances.Set_I()
  Optimize.C_OPT.Outfile = ToPort(Reader.C_stdout.Value)
  Optimize.C_compiler.Inline_ask = CTRUE
  ToGenerateGoProducer(Optimize.C_PRODUCER.Value).Current = C_claire
  Core.F_tformat_string(MakeString("\n---- code produced by the generator ------------------- \n"),0,ToType(CEMPTY.Id()).EmptyList())
  Result = Core.F_CALL(Optimize.C_Compile_make_c_function,ARGS(EID{l.Id(),0},EID{MakeString("test").Id(),0},EID{C_void.Id(),0}))
  if !ErrorIn(Result) {
  { 
    var va_arg1 *Optimize.OptimizeMetaOPT
    var va_arg2 *ClaireAny
    va_arg1 = Optimize.C_OPT
    va_arg2 = CNULL
    va_arg1.InMethod = va_arg2
    Result = va_arg2.ToEID()
    } 
  }
  return Result} 

// The EID go function for: g_test @ lambda (throw: true) 
func E_g_test_lambda (l EID) EID { 
  return F_g_test_lambda(ToLambda(OBJ(l)) )} 

// debug (to remove later)
// compile the modules and check that no necessary modules is not
// declared
/* The go function for: compile(m:module) [status=1] */
func F_compile_module (m *ClaireModule) EID { 
  var Result EID
  Result = Core.F_CALL(C_compile,ARGS(EID{Optimize.C_PRODUCER.Value,0},EID{m.Id(),0}))
  return Result} 

// The EID go function for: compile @ module (throw: true) 
func E_compile_module (m EID) EID { 
  return F_compile_module(ToModule(OBJ(m)) )} 

//  shortcut that already exists
// adds a little book_keeping + more detailed summary to the generic method in Part 2
/* The go function for: compile(p:go_producer,m:module) [status=1] */
func (p *GenerateGoProducer) Compile (m *ClaireModule) EID { 
  var Result EID
  C_BadMethods.Value = ToType(C_method.Id()).EmptyList().Id()
  Optimize.C_compiler.NDynamic = 0
  Optimize.C_compiler.NMetheids = 0
  p.Source = Reader.F__7_string(Optimize.C_compiler.Source,m.Name.String_I())
  Result = Core.F_SUPER(C_compile, C_Generate_code_producer, ARGS(EID{p.Id(),0},EID{m.Id(),0}))
  if !ErrorIn(Result) {
  { var arg_1 *ClaireList
    { 
      var v_bag_arg *ClaireAny
      arg_1= ToType(CEMPTY.Id()).EmptyList()
      arg_1.AddFast(m.Id())
      arg_1.AddFast(MakeInteger(Optimize.C_compiler.NLoc).Id())
      arg_1.AddFast(MakeInteger(Optimize.C_compiler.NWarnings).Id())
      arg_1.AddFast(MakeInteger(Optimize.C_compiler.NNotes).Id())
      arg_1.AddFast(MakeInteger(Optimize.C_compiler.NDynamic).Id())
      if (Optimize.C_compiler.NMethods == 0) { 
        v_bag_arg = MakeInteger(0).Id()
        } else {
        v_bag_arg = MakeInteger(((100*Optimize.C_compiler.NMetheids)/Optimize.C_compiler.NMethods)).Id()
        } 
      arg_1.AddFast(v_bag_arg)} 
    Result = Core.F_tformat_string(MakeString("~S: ~A lines of code compiled. ~A warnings, ~A notes. ~A dynamic calls, ~A% exception-ready methods\n"),1,arg_1)
    } 
  }
  return Result} 

// The EID go function for: compile @ go_producer (throw: true) 
func E_compile_go_producer (p EID,m EID) EID { 
  return ToGenerateGoProducer(OBJ(p)).Compile(ToModule(OBJ(m)) )} 

// Creates the "meta" file for the module m.
/* The go function for: gen_mod_file(p:go_producer,m:module) [status=1] */
func (p *GenerateGoProducer) GenModFile (m *ClaireModule) EID { 
  var Result EID
  Result = ToGenerateCodeProducer(p.Id()).GenModFile(m)
  if !ErrorIn(Result) {
  if (Optimize.C_compiler.Safety > 4) { 
    Result = Core.F_tformat_string(MakeString("===== [CROSS]  ~A BAD METHODS : ~S  \n"),1,MakeConstantList(MakeInteger(ToList(C_BadMethods.Value).Length()).Id(),C_BadMethods.Value))
    } else {
    Result = EID{CFALSE.Id(),0}
    } 
  }
  return Result} 

// The EID go function for: gen_mod_file @ go_producer (throw: true) 
func E_Generate_gen_mod_file_go_producer (p EID,m EID) EID { 
  return ToGenerateGoProducer(OBJ(p)).GenModFile(ToModule(OBJ(m)) )} 

// functions are produced in continuous mode for go compiling (gen_functions does nothing)
/* The go function for: gen_functions(p:go_producer,m:module) [status=0] */
func (p *GenerateGoProducer) GenFunctions (m *ClaireModule)  { 
  
  } 

// The EID go function for: gen_functions @ go_producer (throw: false) 
func E_Generate_gen_functions_go_producer (p EID,m EID) EID { 
  ToGenerateGoProducer(OBJ(p)).GenFunctions(ToModule(OBJ(m)) )
  return EVOID} 

// start the produced go file
// Puts the reference to the packages, and some useful comments
// we limit the use of "unsafe" Go package to the module file
/* The go function for: start_file(p:go_producer,s:string,m:module,module?:boolean) [status=1] */
func (p *GenerateGoProducer) StartFile (s *ClaireString,m *ClaireModule,module_ask *ClaireBoolean) EID { 
  var Result EID
  Optimize.C_OPT.Outfile.UseAsOutput()
  PRINC("/***** CLAIRE Compilation of ")
  F_princ_string(ToString(IfThenElse((module_ask == CTRUE),
    MakeString("module").Id(),
    MakeString("file").Id())))
  PRINC(" ")
  F_princ_string(s)
  PRINC(".cl \n         [version ")
  { var arg_1 *ClaireAny
    var try_2 EID
    try_2 = Core.F_release_void()
    if ErrorIn(try_2) {Result = try_2
    } else {
    arg_1 = ANY(try_2)
    Result = Core.F_CALL(C_princ,ARGS(arg_1.ToEID()))
    }
    } 
  if !ErrorIn(Result) {
  PRINC(" / safety ")
  Result = Core.F_print_any(MakeInteger(Optimize.C_compiler.Safety).Id())
  if !ErrorIn(Result) {
  PRINC("] ")
  F_princ_string(F_date_I_integer(0))
  PRINC(" *****/\n\n")
  Result = EVOID
  }}
  if !ErrorIn(Result) {
  p.Namespace_I(m)
  PRINC("import (_ \"fmt\"\n")
  if (module_ask == CTRUE) { 
    var g0006I *ClaireBoolean
    { var arg_3 *ClaireAny
      { 
        var c *ClaireAny
        _ = c
        arg_3= CFALSE.Id()
        var c_support *ClaireList
        c_support = Optimize.C_OPT.Objects
        c_len := c_support.Length()
        for i_it := 0; i_it < c_len; i_it++ { 
          c = c_support.At(i_it)
          if (C_class.Id() == c.Isa.Id()) { 
            arg_3 = CTRUE.Id()
            break
            } 
          } 
        } 
      g0006I = F_boolean_I_any(arg_3)
      } 
    if (g0006I == CTRUE) { 
      PRINC("\t\"unsafe\"\n")
      } 
    } 
  Result = F_Generate_import_declaration_module(m)
  if !ErrorIn(Result) {
  PRINC(")\n")
  F_Generate_dumb_import_module(m)
  Result = ToPort(Reader.C_stdout.Value).UseAsOutput().ToEID()
  }}
  return Result} 

// The EID go function for: start_file @ go_producer (throw: true) 
func E_Generate_start_file_go_producer (p EID,s EID,m EID,module_ask EID) EID { 
  return ToGenerateGoProducer(OBJ(p)).StartFile(ToString(OBJ(s)),
    ToModule(OBJ(m)),
    ToBoolean(OBJ(module_ask)) )} 

// import declarations
/* The go function for: import_declaration(m:module) [status=1] */
func F_Generate_import_declaration_module (m *ClaireModule) EID { 
  var Result EID
  { 
    var x *ClaireAny
    _ = x
    Result= EID{CFALSE.Id(),0}
    var x_support *ClaireList
    x_support = F_Generate_needed_modules_module(m)
    x_len := x_support.Length()
    for i_it := 0; i_it < x_len; i_it++ { 
      x = x_support.At(i_it)
      var loop_1 EID
      _ = loop_1
      { 
      PRINC("\t")
      if (x == C_Kernel.Id()) { 
        PRINC(". ")
        } 
      PRINC("\"")
      { var arg_2 *ClaireAny
        var try_3 EID
        try_3 = Core.F_CALL(C_string_I,ARGS(Core.F_CALL(C_name,ARGS(x.ToEID()))))
        if ErrorIn(try_3) {loop_1 = try_3
        } else {
        arg_2 = ANY(try_3)
        F_princ_string(ToString(arg_2))
        loop_1 = EVOID
        }
        } 
      if ErrorIn(loop_1) {Result = loop_1
      break
      } else {
      PRINC("\"\n")
      }
      }
      } 
    } 
  return Result} 

// The EID go function for: import_declaration @ module (throw: true) 
func E_Generate_import_declaration_module (m EID) EID { 
  return F_Generate_import_declaration_module(ToModule(OBJ(m)) )} 

// go requires an import list without redundancy + we only import
/* The go function for: needed_modules(m:module) [status=0] */
func F_Generate_needed_modules_module (m *ClaireModule) *ClaireList { 
  var Result *ClaireList
  { var l *ClaireList
    { var m2_out *ClaireList = ToType(CEMPTY.Id()).EmptyList()
      { 
        var m2 *ClaireAny
        _ = m2
        var m2_support *ClaireList
        m2_support = Reader.F_add_modules_list(MakeConstantList(m.Id()))
        m2_len := m2_support.Length()
        for i_it := 0; i_it < m2_len; i_it++ { 
          m2 = m2_support.At(i_it)
          if (m2 != m.Id()) { 
            if ((ToModule(m2).MadeOf.Length() != 0) || 
                (m2 == C_Kernel.Id())) { 
              m2_out.AddFast(m2)
              } 
            } 
          } 
        } 
      l = m2_out
      } 
    Result = l
    } 
  return Result} 

// The EID go function for: needed_modules @ module (throw: false) 
func E_Generate_needed_modules_module (m EID) EID { 
  return EID{F_Generate_needed_modules_module(ToModule(OBJ(m)) ).Id(),0}} 

// create a dumb function that prevents the go compiler to complain
/* The go function for: dumb_import(m:module) [status=0] */
func F_Generate_dumb_import_module (m *ClaireModule)  { 
  { var l *ClaireList = F_Generate_needed_modules_module(m)
    if (l.Length() > 1) { 
      PRINC("\n//-------- dumb function to prevent import errors --------\n")
      PRINC("func import_")
      Core.F_gensym_void().Princ()
      PRINC("() ")
      F_Generate_new_block_void()
      { 
        var x *ClaireAny
        _ = x
        var x_support *ClaireList
        x_support = l
        x_len := x_support.Length()
        for i_it := 0; i_it < x_len; i_it++ { 
          x = x_support.At(i_it)
          if (x != C_Kernel.Id()) { 
            PRINC("_ = ")
            F_Generate_cap_short_symbol(ToSymbol(OBJ(Core.F_CALL(C_name,ARGS(x.ToEID())))))
            PRINC(".It")
            F_Generate_breakline_void()
            } 
          } 
        } 
      F_Generate_close_block_void()
      F_Generate_breakline_void()
      } 
    } 
  } 

// The EID go function for: dumb_import @ module (throw: false) 
func E_Generate_dumb_import_module (m EID) EID { 
  F_Generate_dumb_import_module(ToModule(OBJ(m)) )
  return EVOID} 

// pick a thing in module m
/* The go function for: representative(m:module) [status=1] */
func F_Generate_representative_module (m *ClaireModule) EID { 
  var Result EID
  { var x_some *ClaireAny = CNULL
    { 
      var x *ClaireAny
      _ = x
      Result= EID{CFALSE.Id(),0}
      var x_support *ClaireList
      var try_1 EID
      try_1 = Core.F_enumerate_any(Core.F_U_type(ToType(C_class.Id()),ToType(C_property.Id())).Id())
      if ErrorIn(try_1) {Result = try_1
      } else {
      x_support = ToList(OBJ(try_1))
      x_len := x_support.Length()
      for i_it := 0; i_it < x_len; i_it++ { 
        x = x_support.At(i_it)
        var g0009I *ClaireBoolean
        { 
          var v_and4 *ClaireBoolean
          
          v_and4 = Equal(ToSymbol(OBJ(Core.F_CALL(C_name,ARGS(x.ToEID())))).Defined().Id(),m.Id())
          if (v_and4 == CFALSE) {g0009I = CFALSE
          } else { 
            if (x.Isa.IsIn(C_property) == CTRUE) { 
              { var g0007 *ClaireProperty = ToProperty(x)
                { var arg_2 *ClaireAny
                  { 
                    var y *ClaireRestriction
                    _ = y
                    var y_iter *ClaireAny
                    arg_2= CFALSE.Id()
                    for _,y_iter = range(g0007.Restrictions.ValuesO()){ 
                      y = ToRestriction(y_iter)
                      if (y.Module_I.Id() == m.Id()) { 
                        arg_2 = CTRUE.Id()
                        break
                        } 
                      } 
                    } 
                  v_and4 = F_boolean_I_any(arg_2)
                  } 
                } 
              } else {
              v_and4 = CTRUE
              } 
            if (v_and4 == CFALSE) {g0009I = CFALSE
            } else { 
              g0009I = CTRUE} 
            } 
          } 
        if (g0009I == CTRUE) { 
          x_some = x
          Result = x_some.ToEID()
          break
          } 
        }
        } 
      } 
    if !ErrorIn(Result) {
    Result = x_some.ToEID()
    }
    } 
  return Result} 

// The EID go function for: representative @ module (throw: true) 
func E_Generate_representative_module (m EID) EID { 
  return F_Generate_representative_module(ToModule(OBJ(m)) )} 

// remove dual imports (hopefully, works if the import path is simple enough)
/* The go function for: clean_duplicates(l:list) [status=0] */
func F_Generate_clean_duplicates_list (l *ClaireList) *ClaireList { 
  var Result *ClaireList
  { var l2 *ClaireList = l.Copy()
    { var n int = l.Length()
      { var i int = (n-1)
        for (i > 1) { 
          var g0011I *ClaireBoolean
          { var arg_1 *ClaireAny
            { var j int = (i+1)
              { var g0010 int = n
                arg_1= CFALSE.Id()
                for (j <= g0010) { 
                  if (ToBoolean(Reader.F_add_modules_list(MakeConstantList(l.At(j-1))).Contain_ask(l.At(i-1)).Id()) == CTRUE) { 
                    arg_1 = CTRUE.Id()
                    break
                    } 
                  j = (j+1)
                  } 
                } 
              } 
            g0011I = F_boolean_I_any(arg_1)
            } 
          if (g0011I == CTRUE) { 
            l2 = l2.Delete(l.At(i-1))
            } 
          i = (i-1)
          } 
        Result = l2
        } 
      } 
    } 
  return Result} 

// The EID go function for: clean_duplicates @ list (throw: false) 
func E_Generate_clean_duplicates_list (l EID) EID { 
  return EID{F_Generate_clean_duplicates_list(ToList(OBJ(l)) ).Id(),0}} 

// For each class we produce two things in the module-generated-file
//   - the struct (with embedded inheritance)
//   - the cast method
//   - we also gerenate a constructor  makeC(a1, ... , an) when there are no inverses 
/* The go function for: gen_classes(p:go_producer,m:module) [status=1] */
func (p *GenerateGoProducer) GenClasses (m *ClaireModule) EID { 
  var Result EID
  Core.F_tformat_string(MakeString("===== generate classes for ~S ==== \n"),3,MakeConstantList(m.Id()))
  { 
    var c *ClaireAny
    _ = c
    Result= EID{CFALSE.Id(),0}
    var c_support *ClaireList
    c_support = Optimize.C_OPT.Objects
    c_len := c_support.Length()
    for i_it := 0; i_it < c_len; i_it++ { 
      c = c_support.At(i_it)
      var loop_1 EID
      _ = loop_1
      if (C_class.Id() == c.Isa.Id()) { 
        Optimize.C_OPT.Level = 0
        PRINC("\n// class file for ")
        loop_1 = Core.F_CALL(C_print,ARGS(c.ToEID()))
        if ErrorIn(loop_1) {Result = loop_1
        break
        } else {
        PRINC(" in module ")
        loop_1 = Core.F_print_any(m.Id())
        if ErrorIn(loop_1) {Result = loop_1
        break
        } else {
        PRINC(" ")
        loop_1 = EVOID
        }}
        if ErrorIn(loop_1) {Result = loop_1
        break
        } else {
        F_Generate_breakline_void()
        loop_1 = p.GenClassDef(ToClass(c))
        if ErrorIn(loop_1) {Result = loop_1
        break
        } else {
        p.GenCastFunction(ToClass(c))
        if (F_Generate_construct_class_ask_class(ToClass(c)) == CTRUE) { 
          loop_1 = p.GenConstruct(ToClass(c))
          } else {
          loop_1 = EID{CFALSE.Id(),0}
          } 
        if ErrorIn(loop_1) {Result = loop_1
        break
        } else {
        }}}
        } else {
        loop_1 = EID{CFALSE.Id(),0}
        } 
      if ErrorIn(loop_1) {Result = loop_1
      break
      } else {
      }
      } 
    } 
  return Result} 

// The EID go function for: gen_classes @ go_producer (throw: true) 
func E_Generate_gen_classes_go_producer (p EID,m EID) EID { 
  return ToGenerateGoProducer(OBJ(p)).GenClasses(ToModule(OBJ(m)) )} 

// beware of covariant slot redefinition : go only knows the rootSlot
/* The go function for: rootSlot(s:slot) [status=0] */
func F_Generate_rootSlot_slot (s *ClaireSlot) *ClaireSlot { 
  var Result *ClaireSlot
  { var p *ClaireProperty = s.Selector
    { var c *ClaireClass = Core.F_domain_I_restriction(ToRestriction(s.Id()))
      { var s2 *ClaireSlot = s
        { var i int = s.Index
          for (c.Id() != C_any.Id()) { 
            c = c.Superclass
            if (c.Slots.Length() < i) { 
              
              break
              } else {
              { var s3 *ClaireSlot = ToSlot(c.Slots.ValuesO()[i-1])
                if (s3.Selector.Id() == p.Id()) { 
                  s2 = s3
                  } 
                } 
              } 
            } 
          Result = s2
          } 
        } 
      } 
    } 
  return Result} 

// The EID go function for: rootSlot @ slot (throw: false) 
func E_Generate_rootSlot_slot (s EID) EID { 
  return EID{F_Generate_rootSlot_slot(ToSlot(OBJ(s)) ).Id(),0}} 

// how to generate a struct associated to a class
// notice that we only add the slots that are defined for c, not those inherited from a super class (even with covariant redefinition)
/* The go function for: gen_class_def(p:go_producer,c:class) [status=1] */
func (p *GenerateGoProducer) GenClassDef (c *ClaireClass) EID { 
  var Result EID
  PRINC("type ")
  F_Generate_go_class_class(c)
  PRINC(" struct ")
  F_Generate_new_block_void()
  PRINC(" ")
  F_Generate_go_class_class(c.Superclass)
  F_Generate_breakline_void()
  PRINC(" ")
  { 
    var y *ClaireSlot
    _ = y
    var y_iter *ClaireAny
    Result= EID{CFALSE.Id(),0}
    var y_support *ClaireList
    var try_1 EID
    try_1 = Core.F_CALL(Optimize.C_Compile_get_indexed,ARGS(EID{c.Id(),0}))
    if ErrorIn(try_1) {Result = try_1
    } else {
    y_support = ToList(OBJ(try_1))
    y_len := y_support.Length()
    for i_it := 0; i_it < y_len; i_it++ { 
      y_iter = y_support.At(i_it)
      y = ToSlot(y_iter)
      if (Core.F_domain_I_restriction(ToRestriction(F_Generate_rootSlot_slot(y).Id())).Id() == c.Id()) { 
        F_Generate_cap_short_symbol(y.Selector.Name)
        PRINC(" ")
        F_Generate_interface_I_class(y.Range.Class_I())
        F_Generate_breakline_void()
        } 
      }
      } 
    } 
  if !ErrorIn(Result) {
  F_Generate_close_block_void()
  Result = EVOID
  }
  return Result} 

// The EID go function for: gen_class_def @ go_producer (throw: true) 
func E_Generate_gen_class_def_go_producer (p EID,c EID) EID { 
  return ToGenerateGoProducer(OBJ(p)).GenClassDef(ToClass(OBJ(c)) )} 

// how to produce the ToC() cast function that applies to any pointer (using unsafe)
/* The go function for: gen_cast_function(p:go_producer,c:class) [status=0] */
func (p *GenerateGoProducer) GenCastFunction (c *ClaireClass)  { 
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
  ToGenerateGoProducer(OBJ(p)).GenCastFunction(ToClass(OBJ(c)) )
  return EVOID} 

// when we want a constructor ? when slots are simple (no inverse, no store ...)
// TODO : to complete with the proper test
/* The go function for: construct_class?(c:class) [status=0] */
func F_Generate_construct_class_ask_class (c *ClaireClass) *ClaireBoolean { 
  return  MakeBoolean((ToType(c.Id()).Included(ToType(C_object.Id())) == CTRUE) && (c.Slots.Length() <= 5))
  } 

// The EID go function for: construct_class? @ class (throw: false) 
func E_Generate_construct_class_ask_class (c EID) EID { 
  return EID{F_Generate_construct_class_ask_class(ToClass(OBJ(c)) ).Id(),0}} 

// generate a constructor
/* The go function for: gen_construct(p:go_producer,c:class) [status=1] */
func (p *GenerateGoProducer) GenConstruct (c *ClaireClass) EID { 
  var Result EID
  { var first *ClaireBoolean = CTRUE
    PRINC("\n// automatic constructor function\n")
    PRINC("func Make")
    F_Generate_add_underscore_symbol(c.Name)
    F_Generate_go_class_class(c)
    PRINC("(")
    { 
      var y *ClaireSlot
      _ = y
      var y_iter *ClaireAny
      Result= EID{CFALSE.Id(),0}
      var y_support *ClaireList
      var try_1 EID
      { var arg_2 *ClaireList
        var try_3 EID
        try_3 = Core.F_CALL(Optimize.C_Compile_get_indexed,ARGS(EID{c.Id(),0}))
        if ErrorIn(try_3) {try_1 = try_3
        } else {
        arg_2 = ToList(OBJ(try_3))
        try_1 = arg_2.Cdr()
        }
        } 
      if ErrorIn(try_1) {Result = try_1
      } else {
      y_support = ToList(OBJ(try_1))
      y_len := y_support.Length()
      for i_it := 0; i_it < y_len; i_it++ { 
        y_iter = y_support.At(i_it)
        y = ToSlot(y_iter)
        if (first == CTRUE) { 
          first = CFALSE
          } else {
          PRINC(",")
          } 
        F_iClaire_ident_symbol(y.Selector.Name)
        PRINC(" ")
        F_Generate_interface_I_class(y.Range.Class_I())
        }
        } 
      } 
    if !ErrorIn(Result) {
    PRINC(") *")
    F_Generate_go_class_class(c)
    PRINC(" ")
    F_Generate_new_block_string(MakeString("make"))
    Result = EVOID
    }
    if !ErrorIn(Result) {
    PRINC("var o *")
    F_Generate_go_class_class(c)
    PRINC(" = new(")
    F_Generate_go_class_class(c)
    PRINC(")")
    F_Generate_breakline_void()
    PRINC("o.Isa = ")
    F_Generate_class_ident_class(c)
    F_Generate_breakline_void()
    { 
      var y *ClaireSlot
      _ = y
      var y_iter *ClaireAny
      Result= EID{CFALSE.Id(),0}
      var y_support *ClaireList
      var try_4 EID
      { var arg_5 *ClaireList
        var try_6 EID
        try_6 = Core.F_CALL(Optimize.C_Compile_get_indexed,ARGS(EID{c.Id(),0}))
        if ErrorIn(try_6) {try_4 = try_6
        } else {
        arg_5 = ToList(OBJ(try_6))
        try_4 = arg_5.Cdr()
        }
        } 
      if ErrorIn(try_4) {Result = try_4
      } else {
      y_support = ToList(OBJ(try_4))
      y_len := y_support.Length()
      for i_it := 0; i_it < y_len; i_it++ { 
        y_iter = y_support.At(i_it)
        y = ToSlot(y_iter)
        var loop_7 EID
        _ = loop_7
        { 
        PRINC("o.")
        F_Generate_cap_short_symbol(y.Selector.Name)
        PRINC(" = ")
        loop_7 = F_Generate_cast_prefix_class(y.Range.Class_I(),F_Generate_rootSlot_slot(y).Range.Class_I())
        if ErrorIn(loop_7) {Result = loop_7
        break
        } else {
        F_iClaire_ident_symbol(y.Selector.Name)
        F_Generate_cast_post_class(y.Range.Class_I(),F_Generate_rootSlot_slot(y).Range.Class_I())
        F_Generate_breakline_void()
        }
        }}
        } 
      } 
    if !ErrorIn(Result) {
    PRINC("return o ")
    F_Generate_breakline_void()
    F_Generate_close_block_string(MakeString("make"))
    Result = EVOID
    }}
    } 
  return Result} 

// The EID go function for: gen_construct @ go_producer (throw: true) 
func E_Generate_gen_construct_go_producer (p EID,c EID) EID { 
  return ToGenerateGoProducer(OBJ(p)).GenConstruct(ToClass(OBJ(c)) )} 

// generate the definition of the named objects from the module (used in both modes)
// must move to the producer
/* The go function for: gen_objects(p:go_producer,m:module) [status=1] */
func (p *GenerateGoProducer) GenObjects (m *ClaireModule) EID { 
  var Result EID
  Core.F_tformat_string(MakeString("===== generate objects for ~S [graph : ~S] ==== \n"),3,MakeConstantList(m.Id(),Optimize.C_OPT.Properties.Contain_ask(C_mClaire_graph.Id()).Id()))
  { 
    var x *ClaireAny
    _ = x
    var x_support *ClaireList
    x_support = Optimize.C_OPT.Objects
    x_len := x_support.Length()
    for i_it := 0; i_it < x_len; i_it++ { 
      x = x_support.At(i_it)
      F_Generate_breakline_void()
      if (x.Isa.IsIn(Core.C_global_variable) == CTRUE) { 
        { var g0012 *Core.GlobalVariable = Core.ToGlobalVariable(x)
          PRINC("var ")
          F_Generate_go_var_symbol(g0012.Name)
          PRINC(" ")
          { var arg_1 *ClaireClass
            if (Optimize.F_Compile_nativeVar_ask_global_variable(g0012) == CTRUE) { 
              arg_1 = F_Generate_getRange_global_variable(g0012)
              } else {
              arg_1 = Core.C_global_variable
              } 
            F_Generate_interface_I_class(arg_1)
            } 
          } 
        } else {
        PRINC("var ")
        F_Generate_go_var_symbol(ToSymbol(OBJ(Core.F_CALL(C_name,ARGS(x.ToEID())))))
        PRINC(" ")
        F_Generate_interface_I_class(Optimize.F_Compile_psort_any(x.Isa.Id()))
        PRINC(" /*obj*/")
        } 
      } 
    } 
  { 
    var x *ClaireProperty
    _ = x
    var x_iter *ClaireAny
    Result= EID{CFALSE.Id(),0}
    var x_support *ClaireSet
    x_support = Optimize.C_OPT.Properties
    for i_it := 0; i_it < x_support.Count; i_it++ { 
      x_iter = x_support.At(i_it)
      x = ToProperty(x_iter)
      var loop_2 EID
      _ = loop_2
      if (Optimize.C_OPT.Objects.Memq(x.Id()) != CTRUE) { 
        { var p2test *ClaireAny
          { var p2_some *ClaireAny = CNULL
            { 
              var p2 *ClaireProperty
              _ = p2
              var p2_iter *ClaireAny
              var p2_support *ClaireSet
              p2_support = Optimize.C_OPT.Properties
              for i_it := 0; i_it < p2_support.Count; i_it++ { 
                p2_iter = p2_support.At(i_it)
                p2 = ToProperty(p2_iter)
                if (p2.Id() != x.Id()) { 
                  if (p2.Name.String_I().Value == x.Name.String_I().Value) { 
                    p2_some = p2.Id()
                    break
                    } 
                  } 
                } 
              } 
            p2test = p2_some
            } 
          if (p2test != CNULL) { 
            { var p2 *ClaireProperty = ToProperty(p2test)
              loop_2 = ToException(Core.C_general_error.Make(MakeString("[217] ~S and ~S cannot be defined in the same module").Id(),MakeConstantList(p2.Id(),x.Id()).Id())).Close()
              } 
            } else {
            loop_2 = EID{CNULL,0}
            } 
          } 
        if ErrorIn(loop_2) {Result = loop_2
        break
        } else {
        F_Generate_breakline_void()
        PRINC("var ")
        F_Generate_thing_ident_thing(ToThing(x.Id()))
        PRINC(" ")
        F_Generate_interface_I_class(Optimize.F_Compile_psort_any(x.Id().Isa.Id()))
        PRINC(" // ")
        loop_2 = Core.F_print_any(x.Name.Id())
        if ErrorIn(loop_2) {Result = loop_2
        break
        } else {
        }
        if ErrorIn(loop_2) {Result = loop_2
        break
        } else {
        }}
        } else {
        loop_2 = EID{CFALSE.Id(),0}
        } 
      if ErrorIn(loop_2) {Result = loop_2
      break
      } else {
      }
      } 
    } 
  if !ErrorIn(Result) {
  F_Generate_breakline_void()
  PRINC("var It *ClaireModule")
  F_Generate_breakline_void()
  { var m1 *ClaireModule = m
    { var m2 *ClaireModule = m.PartOf
      for ((m2.Id() != C_claire.Id()) && 
          (m2.Parts.At(0) == m1.Id())) { 
        if (Equal(m2.MadeOf.Id(),CNIL.Id()) == CTRUE) { 
          PRINC("var ")
          F_Generate_go_var_symbol(m2.Name)
          PRINC(" *ClaireModule ")
          } 
        m1 = m2
        m2 = m2.PartOf
        } 
      } 
    } 
  Result = F_Generate_breakline_void().ToEID()
  }
  return Result} 

// The EID go function for: gen_objects @ go_producer (throw: true) 
func E_Generate_gen_objects_go_producer (p EID,m EID) EID { 
  return ToGenerateGoProducer(OBJ(p)).GenObjects(ToModule(OBJ(m)) )} 

// extract the range for a global_variable
/* The go function for: getRange(x:global_variable) [status=0] */
func F_Generate_getRange_global_variable (x *Core.GlobalVariable) *ClaireClass { 
  if (Equal(x.Range.Id(),CEMPTY.Id()) == CTRUE) { 
    return  x.Value.Isa
    } else {
    return  x.Range.Class_I()
    } 
  } 

// The EID go function for: getRange @ global_variable (throw: false) 
func E_Generate_getRange_global_variable (x EID) EID { 
  return EID{F_Generate_getRange_global_variable(Core.ToGlobalVariable(OBJ(x)) ).Id(),0}} 

// generate the meta_load function
// in go the load function for M is M_load()
/* The go function for: gen_meta_load(p:go_producer,m:module) [status=1] */
func (p *GenerateGoProducer) GenMetaLoad (m *ClaireModule) EID { 
  var Result EID
  Core.F_tformat_string(MakeString("===== generate meta_load function for ~S ==== \n"),3,MakeConstantList(m.Id()))
  PRINC("// definition of the meta-model for module ")
  Result = Core.F_print_any(m.Id())
  if !ErrorIn(Result) {
  PRINC(" ")
  Result = F_Generate_breakline_void().ToEID()
  }
  if !ErrorIn(Result) {
  PRINC("func MetaLoad() ")
  F_Generate_new_block_void()
  F_Generate_breakline_void()
  Result = p.GenModule(m,m)
  if !ErrorIn(Result) {
  F_Generate_breakline_void()
  PRINC("// definition of the properties")
  { 
    var x *ClaireProperty
    _ = x
    var x_iter *ClaireAny
    Result= EID{CFALSE.Id(),0}
    var x_support *ClaireSet
    x_support = Optimize.C_OPT.Properties
    for i_it := 0; i_it < x_support.Count; i_it++ { 
      x_iter = x_support.At(i_it)
      x = ToProperty(x_iter)
      var loop_1 EID
      _ = loop_1
      if ((Optimize.C_OPT.Objects.Memq(x.Id()) != CTRUE) && 
          ((x.Id() != C_value.Id()) && 
            (x.Id() != C_vars.Id()))) { 
        F_Generate_breakline_void()
        F_Generate_thing_ident_thing(ToThing(x.Id()))
        PRINC(" = ")
        loop_1 = p.Declare(x)
        if ErrorIn(loop_1) {Result = loop_1
        break
        } else {
        }
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
  F_Generate_breakline_void()
  PRINC("// instructions from module sources")
  { var j *ClaireAny = CNULL
    { 
      var i *ClaireAny
      _ = i
      Result= EID{CFALSE.Id(),0}
      var i_support *ClaireList
      i_support = Optimize.C_OPT.Instructions
      i_len := i_support.Length()
      for i_it := 0; i_it < i_len; i_it++ { 
        i = i_support.At(i_it)
        var loop_2 EID
        _ = loop_2
        { 
        F_Generate_breakline_void()
        if (C_string.Id() == i.Isa.Id()) { 
          if (C_string.Id() != j.Isa.Id()) { 
            F_Generate_breakline_void()
            } 
          PRINC("// ")
          loop_2 = Core.F_CALL(C_princ,ARGS(i.ToEID()))
          if ErrorIn(loop_2) {Result = loop_2
          break
          } else {
          }
          } else {
          var g0014I *ClaireBoolean
          var try_3 EID
          try_3 = Optimize.F_Compile_g_throw_any(i)
          if ErrorIn(try_3) {loop_2 = try_3
          } else {
          g0014I = ToBoolean(OBJ(try_3))
          if (g0014I == CTRUE) { 
            F_Generate_new_block_void()
            if (p.Debug_ask == CTRUE) { 
              PRINC("/*PROTECT ")
              loop_2 = Core.F_CALL(C_print,ARGS(i.ToEID()))
              if ErrorIn(loop_2) {Result = loop_2
              break
              } else {
              PRINC(" */")
              loop_2 = F_Generate_breakline_void().ToEID()
              }
              } else {
              loop_2 = EID{CFALSE.Id(),0}
              } 
            if ErrorIn(loop_2) {Result = loop_2
            break
            } else {
            F_Generate_var_declaration_string(MakeString("expr"),Optimize.C_EID,1)
            loop_2 = Core.F_CALL(C_Generate_g_statement,ARGS(i.ToEID(),
              EID{Optimize.C_EID.Id(),0},
              EID{MakeString("expr").Id(),0},
              EID{CTRUE.Id(),0},
              EID{CFALSE.Id(),0}))
            if ErrorIn(loop_2) {Result = loop_2
            break
            } else {
            PRINC("ErrorCheck(expr)")
            F_Generate_close_block_void()
            loop_2 = EVOID
            }}
            } else {
            var g0015I *ClaireBoolean
            var try_4 EID
            try_4 = F_Generate_simple_func_ask_any(i)
            if ErrorIn(try_4) {loop_2 = try_4
            } else {
            g0015I = ToBoolean(OBJ(try_4))
            if (g0015I == CTRUE) { 
              PRINC("_ = ")
              { var arg_5 *ClaireClass
                var try_6 EID
                { var arg_7 *ClaireType
                  var try_8 EID
                  try_8 = Core.F_CALL(Optimize.C_c_type,ARGS(i.ToEID()))
                  if ErrorIn(try_8) {try_6 = try_8
                  } else {
                  arg_7 = ToType(OBJ(try_8))
                  try_6 = EID{arg_7.Class_I().Id(),0}
                  }
                  } 
                if ErrorIn(try_6) {loop_2 = try_6
                } else {
                arg_5 = ToClass(OBJ(try_6))
                loop_2 = Core.F_CALL(C_Generate_g_expression,ARGS(i.ToEID(),EID{arg_5.Id(),0}))
                }
                } 
              if ErrorIn(loop_2) {Result = loop_2
              break
              } else {
              loop_2 = F_Generate_breakline_void().ToEID()
              }
              } else {
              loop_2 = F_Generate_statement_any(i,C_void,MakeString("Niet"),CFALSE.Id())
              } 
            }
            } 
          }
          } 
        if ErrorIn(loop_2) {Result = loop_2
        break
        } else {
        j = i
        }
        }
        } 
      } 
    } 
  if !ErrorIn(Result) {
  F_Generate_breakline_void()
  F_Generate_close_block_void()
  Result = EVOID
  }}}}
  return Result} 

// The EID go function for: gen_meta_load @ go_producer (throw: true) 
func E_Generate_gen_meta_load_go_producer (p EID,m EID) EID { 
  return ToGenerateGoProducer(OBJ(p)).GenMetaLoad(ToModule(OBJ(m)) )} 

// generate the module definition - only the module structure (the decoration is found in the system file)
// cool recursive method that ensures that all non-package modules are visible
// load_m() has an implicit begin(m) so that new methods are assigned to m
/* The go function for: gen_module(p:go_producer,m:module,%package:module) [status=1] */
func (p *GenerateGoProducer) GenModule (m *ClaireModule,_Zpackage *ClaireModule) EID { 
  var Result EID
  if (m.Id() == _Zpackage.Id()) { 
    PRINC("It")
    } else {
    F_Generate_go_var_symbol(m.Name)
    } 
  PRINC(" = MakeModule(")
  Result = Core.F_print_any((m.Name.String_I()).Id())
  if !ErrorIn(Result) {
  PRINC(",")
  Result = F_Generate_g_expression_module(m.PartOf,C_module)
  if !ErrorIn(Result) {
  PRINC(")")
  Result = F_Generate_breakline_void().ToEID()
  }}
  if !ErrorIn(Result) {
  { var s *ClaireString
    var try_1 EID
    if (((m.Comment).Id() != CNULL) && 
        ((F_length_string(m.Comment) <= 8) || 
            (F_substring_string(m.Comment,1,8).Value != MakeString("Compiled").Value))) { 
      try_1 = EID{m.Comment.Id(),0}
      } else {
      { var arg_2 *ClaireString
        var try_3 EID
        { var arg_4 *ClaireString
          var try_5 EID
          { var arg_6 *ClaireString
            var try_7 EID
            { var arg_8 *ClaireString
              var try_9 EID
              { var arg_10 *ClaireString
                var try_11 EID
                { var arg_12 *ClaireString
                  var try_13 EID
                  { var arg_14 *ClaireString
                    var try_15 EID
                    try_15 = Core.F_string_I_float(ClEnv.Version)
                    if ErrorIn(try_15) {try_13 = try_15
                    } else {
                    arg_14 = ToString(OBJ(try_15))
                    try_13 = EID{F_append_string(F_append_string(F_append_string(MakeString("Compiled on "),F_date_I_integer(0)),MakeString("(v4.")),arg_14).Id(),0}
                    }
                    } 
                  if ErrorIn(try_13) {try_11 = try_13
                  } else {
                  arg_12 = ToString(OBJ(try_13))
                  try_11 = EID{F_append_string(arg_12,MakeString("), lines:")).Id(),0}
                  }
                  } 
                if ErrorIn(try_11) {try_9 = try_11
                } else {
                arg_10 = ToString(OBJ(try_11))
                try_9 = EID{F_append_string(arg_10,F_string_I_integer(Optimize.C_compiler.NLoc)).Id(),0}
                }
                } 
              if ErrorIn(try_9) {try_7 = try_9
              } else {
              arg_8 = ToString(OBJ(try_9))
              try_7 = EID{F_append_string(arg_8,MakeString(", warnings:")).Id(),0}
              }
              } 
            if ErrorIn(try_7) {try_5 = try_7
            } else {
            arg_6 = ToString(OBJ(try_7))
            try_5 = EID{F_append_string(arg_6,F_string_I_integer(Optimize.C_compiler.NWarnings)).Id(),0}
            }
            } 
          if ErrorIn(try_5) {try_3 = try_5
          } else {
          arg_4 = ToString(OBJ(try_5))
          try_3 = EID{F_append_string(arg_4,MakeString(",safety:")).Id(),0}
          }
          } 
        if ErrorIn(try_3) {try_1 = try_3
        } else {
        arg_2 = ToString(OBJ(try_3))
        try_1 = EID{F_append_string(arg_2,F_string_I_integer(Optimize.C_compiler.Safety)).Id(),0}
        }
        } 
      } 
    if ErrorIn(try_1) {Result = try_1
    } else {
    s = ToString(OBJ(try_1))
    PRINC("It.Comment = MakeString(")
    Result = Core.F_print_any((s).Id())
    if !ErrorIn(Result) {
    PRINC(")")
    Result = F_Generate_breakline_void().ToEID()
    }
    }
    } 
  if !ErrorIn(Result) {
  if (Optimize.C_compiler.Debug_ask.Memq(m.Id()) == CTRUE) { 
    PRINC("It.Status = 4")
    F_Generate_breakline_void()
    } 
  PRINC("ClEnv.Module_I = It")
  Result = F_Generate_breakline_void().ToEID()
  }}
  return Result} 

// The EID go function for: gen_module @ go_producer (throw: true) 
func E_Generate_gen_module_go_producer (p EID,m EID,_Zpackage EID) EID { 
  return ToGenerateGoProducer(OBJ(p)).GenModule(ToModule(OBJ(m)),ToModule(OBJ(_Zpackage)) )} 

// implicit begin(m)
// reciprocate : finds the concrete module where a package module must be defined.
/* The go function for: get_made(self:module) [status=0] */
func F_Generate_get_made_module (self *ClaireModule) *ClaireModule { 
  var Result *ClaireModule
  { var m *ClaireAny = self.Parts.At(0)
    if ((m == C_Kernel.Id()) || 
        (ToModule(m).MadeOf.Length() != 0)) { 
      Result = ToModule(m)
      } else {
      Result = F_Generate_get_made_module(ToModule(m))
      } 
    } 
  return Result} 

// The EID go function for: get_made @ module (throw: false) 
func E_Generate_get_made_module (self EID) EID { 
  return EID{F_Generate_get_made_module(ToModule(OBJ(self)) ).Id(),0}} 

// called by gosystem.cl : declare a property or an operation (handles the dispatch case)
/* The go function for: declare(c:go_producer,p:property) [status=1] */
func (c *GenerateGoProducer) Declare (p *ClaireProperty) EID { 
  var Result EID
  PRINC("Make")
  F_princ_string(ToString(IfThenElse((p.Isa.IsIn(C_operation) == CTRUE),
    MakeString("Operation").Id(),
    MakeString("Property").Id())))
  PRINC("(")
  Result = Core.F_print_any((p.Name.String_I()).Id())
  if !ErrorIn(Result) {
  PRINC(",")
  F_princ_integer(p.Open)
  PRINC(",")
  Result = F_Generate_g_expression_module(p.Name.Module_I(),C_module)
  if !ErrorIn(Result) {
  if (p.Isa.IsIn(C_operation) == CTRUE) { 
    PRINC(",")
    F_princ_integer(ToOperation(p.Id()).Precedence)
    } 
  PRINC(")")
  Result = EVOID
  }}
  return Result} 

// The EID go function for: declare @ go_producer (throw: true) 
func E_Generate_declare_go_producer (c EID,p EID) EID { 
  return ToGenerateGoProducer(OBJ(c)).Declare(ToProperty(OBJ(p)) )} 

// This is a similar method which places all the necessary modules
// in the right order so that self can be defined
/* The go function for: parents(self:module,l:list) [status=0] */
func F_Generate_parents_module (self *ClaireModule,l *ClaireList) *ClaireList { 
  if (l.Memq(self.Id()) == CTRUE) { 
    return  l
    } else {
    if (self.PartOf.Id() != CNULL) { 
      l = F_Generate_parents_module(self.PartOf,l)
      } 
    l = l.AddFast(self.Id())
    return  l
    } 
  } 

// The EID go function for: parents @ module (throw: false) 
func E_Generate_parents_module (self EID,l EID) EID { 
  return EID{F_Generate_parents_module(ToModule(OBJ(self)),ToList(OBJ(l)) ).Id(),0}} 

// this methods takes a list of modules that must be loaded and returns
// a list of modules that are necessary for the definition
/* The go function for: parents(self:list) [status=0] */
func F_Generate_parents_list (self *ClaireList) *ClaireList { 
  var Result *ClaireList
  { var l *ClaireList = ToType(C_module.Id()).EmptyList()
    { 
      var x *ClaireAny
      _ = x
      var x_support *ClaireList
      x_support = self
      x_len := x_support.Length()
      for i_it := 0; i_it < x_len; i_it++ { 
        x = x_support.At(i_it)
        l = F_Generate_parents_module(ToModule(x),l)
        } 
      } 
    Result = l
    } 
  return Result} 

// The EID go function for: parents @ list (throw: false) 
func E_Generate_parents_list (self EID) EID { 
  return EID{F_Generate_parents_list(ToList(OBJ(self)) ).Id(),0}} 

// useful (v3.0.06)
/* The go function for: get(m:module) [status=1] */
func F_get_module2 (m *ClaireModule) EID { 
  var Result EID
  Result = Reader.F_load_module(m)
  if !ErrorIn(Result) {
  m.Begin()
  Result = EVOID
  }
  return Result} 

// The EID go function for: get @ list<type_expression>(module) (throw: true) 
func E_get_module2 (m EID) EID { 
  return F_get_module2(ToModule(OBJ(m)) )} 

// *********************************************************************
// *     Part 4: File compilation                                      *
// *********************************************************************
// Go compiler need the output port to be open for continuous generation
/* The go function for: outfile?(p:go_producer) [status=0] */
func (p *GenerateGoProducer) Outfile_ask () *ClaireBoolean { 
  return  CTRUE
  } 

// The EID go function for: outfile? @ go_producer (throw: false) 
func E_Generate_outfile_ask_go_producer (p EID) EID { 
  return EID{ToGenerateGoProducer(OBJ(p)).Outfile_ask( ).Id(),0}} 

// name of the file generated for the module
/* The go function for: modfile_name(p:go_producer,m:module) [status=0] */
func (p *GenerateGoProducer) ModfileName (m *ClaireModule) *ClaireString { 
  return  F_append_string(m.Name.String_I(),MakeString("-meta"))
  } 

// The EID go function for: modfile_name @ list<type_expression>(go_producer, module) (throw: false) 
func E_Generate_modfile_name_go_producer2 (p EID,m EID) EID { 
  return EID{ToGenerateGoProducer(OBJ(p)).ModfileName(ToModule(OBJ(m)) ).Id(),0}} 

// compiles one instruction from the CLAIRE file
// comments are printed (thanks to OPT.outfile) others are stacked in OPT.instructions  
/* The go function for: gen_instruction(p:go_producer,%instruction:any,prev_comment:string) [status=1] */
func (p *GenerateGoProducer) GenInstruction (_Zinstruction *ClaireAny,prev_comment *ClaireString) EID { 
  var Result EID
  if (C_string.Id() == _Zinstruction.Isa.Id()) { 
    { var pp *ClairePort = Optimize.C_OPT.Outfile.UseAsOutput()
      PRINC("\n//")
      Result = Core.F_CALL(C_princ,ARGS(_Zinstruction.ToEID()))
      if !ErrorIn(Result) {
      Result = pp.UseAsOutput().ToEID()
      }
      } 
    } else {
    { 
      var va_arg1 *Optimize.OptimizeMetaOPT
      var va_arg2 *ClaireList
      va_arg1 = Optimize.C_OPT
      var try_1 EID
      { var arg_2 *ClaireAny
        var try_3 EID
        try_3 = Core.F_CALL(Optimize.C_c_code,ARGS(_Zinstruction.ToEID(),EID{C_void.Id(),0}))
        if ErrorIn(try_3) {try_1 = try_3
        } else {
        arg_2 = ANY(try_3)
        try_1 = EID{Optimize.C_OPT.Instructions.AddFast(arg_2).Id(),0}
        }
        } 
      if ErrorIn(try_1) {Result = try_1
      } else {
      va_arg2 = ToList(OBJ(try_1))
      va_arg1.Instructions = va_arg2
      Result = EID{va_arg2.Id(),0}
      }
      } 
    } 
  if !ErrorIn(Result) {
  Result = EID{(prev_comment).Id(),0}
  }
  return Result} 

// The EID go function for: gen_instruction @ go_producer (throw: true) 
func E_Generate_gen_instruction_go_producer (p EID,_Zinstruction EID,prev_comment EID) EID { 
  return ToGenerateGoProducer(OBJ(p)).GenInstruction(ANY(_Zinstruction),ToString(OBJ(prev_comment)) )} 

// sugar
/* The go function for: fileName(s:string) [status=0] */
func F_Generate_fileName_string (s *ClaireString) *ClaireString { 
  var Result *ClaireString
  { var n int = F_length_string(s)
    { var i int = F_get_string(s,ToString(Reader.C__starfs_star.Value).At(1))
      if (i > 0) { 
        Result = F_Generate_fileName_string(F_substring_string(s,(i+1),n))
        } else {
        Result = s
        } 
      } 
    } 
  return Result} 

// The EID go function for: fileName @ string (throw: false) 
func E_Generate_fileName_string (s EID) EID { 
  return EID{F_Generate_fileName_string(ToString(OBJ(s)) ).Id(),0}} 

//**********************************************************************
//*     Part 5: the lambda-to-function compiler                        *
//**********************************************************************
// This is simplified in CLAIRE4 since the class2file mode is no longer supported
// we could re-introduce it from CLAIRE 3.5 if we want to support Java compiling
/* The go function for: Compile/make_c_function(self:lambda,%nom:string,m:any) [status=1] */
func F_Compile_make_c_function_lambda (self *ClaireLambda,_Znom *ClaireString,m *ClaireAny) EID { 
  var Result EID
  if (C_method.Id() == m.Isa.Id()) { 
    Result = Core.F_CALL(C_Generate_gen_function,ARGS(EID{Optimize.C_PRODUCER.Value,0},
      EID{self.Id(),0},
      EID{(_Znom).Id(),0},
      m.ToEID()))
    } else {
    Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).MakeLambdaFunction(self,_Znom)
    } 
  return Result} 

// The EID go function for: Compile/make_c_function @ lambda (throw: true) 
func E_Compile_make_c_function_lambda (self EID,_Znom EID,m EID) EID { 
  return F_Compile_make_c_function_lambda(ToLambda(OBJ(self)),ToString(OBJ(_Znom)),ANY(m) )} 

// In CLAIRE 4 we separate methods from free lambdas (used for demons, but which could be used to compile lambda blocks)
// this is used for demons as well as second-order-types
// create an EID lambda  
/* The go function for: make_lambda_function(p:go_producer,self:lambda,%nom:string) [status=1] */
func (p *GenerateGoProducer) MakeLambdaFunction (self *ClaireLambda,_Znom *ClaireString) EID { 
  var Result EID
  { var _Zbody *ClaireAny
    var try_1 EID
    try_1 = Core.F_CALL(Optimize.C_c_code,ARGS(self.Body.ToEID(),EID{C_any.Id(),0}))
    if ErrorIn(try_1) {Result = try_1
    } else {
    _Zbody = ANY(try_1)
    Core.F_tformat_string(MakeString("===== generate an EID function from a lambda for ~A \n"),0,MakeConstantList((_Znom).Id()))
    PRINC("Optimization:\nlambda = ")
    Result = Core.F_print_any(self.Id())
    if !ErrorIn(Result) {
    PRINC("\n optimized = ")
    Result = Core.F_CALL(C_print,ARGS(_Zbody.ToEID()))
    if !ErrorIn(Result) {
    PRINC("\n")
    Result = EVOID
    }}
    if !ErrorIn(Result) {
    Optimize.C_OPT.Outfile.UseAsOutput()
    Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).GenerateFunctionStart(self,
      Optimize.C_EID,
      CNIL.Id(),
      _Znom)
    if !ErrorIn(Result) {
    F_Generate_new_block_void()
    if (p.Debug_ask == CTRUE) { 
      PRINC("/* eid body: ")
      Result = Core.F_CALL(C_print,ARGS(self.Body.ToEID()))
      if !ErrorIn(Result) {
      PRINC(" */")
      Result = F_Generate_breakline_void().ToEID()
      }
      } else {
      Result = EID{CFALSE.Id(),0}
      } 
    if !ErrorIn(Result) {
    Result = F_Generate_eid_body_any(_Zbody,CTRUE,Optimize.C_EID)
    if !ErrorIn(Result) {
    F_Generate_close_block_void()
    F_Generate_breakline_void()
    Result = F_Generate_generate_eid_dual_lambda(self,_Znom)
    if !ErrorIn(Result) {
    Result = ToPort(Reader.C_stdout.Value).UseAsOutput().ToEID()
    }}}}}
    }
    } 
  return Result} 

// The EID go function for: make_lambda_function @ go_producer (throw: true) 
func E_Generate_make_lambda_function_go_producer (p EID,self EID,_Znom EID) EID { 
  return ToGenerateGoProducer(OBJ(p)).MakeLambdaFunction(ToLambda(OBJ(self)),ToString(OBJ(_Znom)) )} 

// how to declare a function in the interface file and its header in the
// output file
/* The go function for: generate_function_start(p:go_producer,self:lambda,s:class,m:any,%nom:string) [status=1] */
func (p *GenerateGoProducer) GenerateFunctionStart (self *ClaireLambda,s *ClaireClass,m *ClaireAny,_Znom *ClaireString) EID { 
  var Result EID
  { var _Zdom *ClaireType
    if (self.Vars.Length() != 0) { 
      _Zdom = ToType(OBJ(Core.F_CALL(C_range,ARGS(self.Vars.At(0).ToEID()))))
      } else {
      _Zdom = ToType(C_any.Id())
      } 
    { var _Zf *ClaireFunction = F_make_function_string(_Znom)
      { var lv *ClaireList
        if ((self.Vars.Length() == 1) && 
            ((_Zdom.Id() == C_void.Id()) || 
                (_Zdom.Id() == C_environment.Id()))) { 
          lv = CNIL
          } else {
          lv = self.Vars
          } 
        Optimize.C_OPT.Functions = Optimize.C_OPT.Functions.AddFast(MakeConstantList(_Zf.Id(),lv.Id(),s.Id()).Id())
        PRINC("\n/* The go function for: ")
        if (C_method.Id() == m.Isa.Id()) { 
          { var g0018 *ClaireMethod = ToMethod(m)
            Result = Core.F_print_any(g0018.Selector.Id())
            if !ErrorIn(Result) {
            PRINC("(")
            Result = Language.F_ppvariable_list(self.Vars)
            if !ErrorIn(Result) {
            PRINC(") [status=")
            F_princ_integer(g0018.Status)
            PRINC("]")
            Result = EVOID
            }}
            } 
          } else {
          F_princ_string(F_string_I_function(_Zf))
          Result = EVOID
          } 
        if !ErrorIn(Result) {
        PRINC(" */\n")
        Result = EVOID
        }
        if !ErrorIn(Result) {
        if (F_Generate_goMethod_ask_any(m) == CTRUE) { 
          PRINC("func (")
          p.GoVariable(To_Variable(self.Vars.At(0)))
          PRINC(") ")
          F_Generate_goMethod_method(ToMethod(m))
          PRINC(" (")
          { var arg_1 *ClaireList
            var try_2 EID
            try_2 = self.Vars.Cdr()
            if ErrorIn(try_2) {Result = try_2
            } else {
            arg_1 = ToList(OBJ(try_2))
            Result = p.GoVariables(arg_1).ToEID()
            }
            } 
          if !ErrorIn(Result) {
          PRINC(") ")
          if (s.Id() != C_void.Id()) { 
            F_Generate_interface_I_class(s)
            } 
          PRINC(" ")
          Result = EVOID
          }
          }  else if (Equal(m,CNIL.Id()) == CTRUE) { 
          PRINC("func F_")
          F_c_princ_string(_Znom)
          PRINC(" (")
          p.GoVariables(lv)
          PRINC(") EID ")
          Result = EVOID
          } else {
          PRINC("func ")
          Result = F_Generate_goFunction_method(ToMethod(m))
          if !ErrorIn(Result) {
          PRINC(" (")
          p.GoVariables(lv)
          PRINC(") ")
          if (s.Id() != C_void.Id()) { 
            F_Generate_interface_I_class(s)
            } 
          PRINC(" ")
          Result = EVOID
          }
          } 
        }
        } 
      } 
    } 
  return Result} 

// The EID go function for: generate_function_start @ go_producer (throw: true) 
func E_Generate_generate_function_start_go_producer (p EID,self EID,s EID,m EID,_Znom EID) EID { 
  return ToGenerateGoProducer(OBJ(p)).GenerateFunctionStart(ToLambda(OBJ(self)),
    ToClass(OBJ(s)),
    ANY(m),
    ToString(OBJ(_Znom)) )} 

// This method creates a go function from a claire lambda for a method m.
// %name is the name that was proposed for the lambda (or derived from method m)
// we either use function_body to try a simple approach or (procedure_body | eid_body) that add all the trimmings
/* The go function for: gen_function(p:go_producer,self:lambda,%nom:string,m:method) [status=1] */
func (p *GenerateGoProducer) GenFunction (self *ClaireLambda,_Znom *ClaireString,m *ClaireMethod) EID { 
  var Result EID
  { var typeOK *ClaireAny
    var try_1 EID
    try_1 = F_Generate_check_range_method(m,self.Body)
    if ErrorIn(try_1) {Result = try_1
    } else {
    typeOK = ANY(try_1)
    { var s *ClaireClass = m.Range.Class_I()
      { var _Zbody *ClaireAny
        var try_2 EID
        try_2 = Optimize.F_Compile_c_strict_code_any(self.Body,s)
        if ErrorIn(try_2) {Result = try_2
        } else {
        _Zbody = ANY(try_2)
        { var throw_ask *ClaireBoolean
          var try_3 EID
          try_3 = Optimize.F_Compile_g_throw_any(_Zbody)
          if ErrorIn(try_3) {Result = try_3
          } else {
          throw_ask = ToBoolean(OBJ(try_3))
          { var arg_4 *ClaireList
            var try_5 EID
            { 
              var v_bag_arg *ClaireAny
              try_5= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
              ToList(OBJ(try_5)).AddFast(MakeInteger(ClEnv.NLine).Id())
              ToList(OBJ(try_5)).AddFast(m.Id())
              var try_6 EID
              try_6 = F_Generate_simple_body_ask_any(_Zbody)
              if ErrorIn(try_6) {try_5 = try_6
              } else {
              v_bag_arg = ANY(try_6)
              ToList(OBJ(try_5)).AddFast(v_bag_arg)
              ToList(OBJ(try_5)).AddFast(throw_ask.Id())}
              } 
            if ErrorIn(try_5) {Result = try_5
            } else {
            arg_4 = ToList(OBJ(try_5))
            Result = Core.F_tformat_string(MakeString("[~A] ~S: => simple=~S, throw=~S \n"),2,arg_4)
            }
            } 
          if !ErrorIn(Result) {
          Optimize.C_compiler.NMethods = (Optimize.C_compiler.NMethods+1)
          p.Varsym = 0
          if (m.Status == -1) { 
            { 
              var va_arg1 *ClaireMethod
              var va_arg2 int
              va_arg1 = m
              if (throw_ask == CTRUE) { 
                va_arg2 = 1
                } else {
                va_arg2 = 0
                } 
              va_arg1.Status = va_arg2
              Result = EID{C__INT,IVAL(va_arg2)}
              } 
            } else {
            var g0022I *ClaireBoolean
            var try_7 EID
            { var arg_8 *ClaireBoolean
              var try_9 EID
              try_9 = Optimize.F_Compile_can_throw_ask_method(m)
              if ErrorIn(try_9) {try_7 = try_9
              } else {
              arg_8 = ToBoolean(OBJ(try_9))
              try_7 = EID{Core.F__I_equal_any(throw_ask.Id(),arg_8.Id()).Id(),0}
              }
              } 
            if ErrorIn(try_7) {Result = try_7
            } else {
            g0022I = ToBoolean(OBJ(try_7))
            if (g0022I == CTRUE) { 
              Optimize.F_Compile_warn_void()
              Core.F_tformat_string(MakeString("[CROSS] ~S body produces an error (g_throw = true) while status is 0 \n"),1,MakeConstantList(m.Id()))
              if (m.Status == 0) { 
                var v_gassign10 *ClaireAny
                v_gassign10 = ToList(C_BadMethods.Value).AddFast(m.Id()).Id()
                C_BadMethods.Value = v_gassign10
                Result = v_gassign10.ToEID()
                } else {
                throw_ask = CTRUE
                Result = EID{throw_ask.Id(),0}
                } 
              } else {
              Result = EID{CFALSE.Id(),0}
              } 
            }
            } 
          if !ErrorIn(Result) {
          Optimize.C_OPT.Outfile.UseAsOutput()
          if (((F_boolean_I_any(typeOK) == CTRUE) || 
                (Optimize.C_compiler.Safety >= 2)) && 
              ((throw_ask != CTRUE) && 
                (m.Selector.Id() != Core.C_self_eval.Id()))) { 
            
            if (p.Debug_ask == CTRUE) { 
              PRINC("// DEBUG: g_throw=")
              Result = Core.F_print_any(throw_ask.Id())
              if !ErrorIn(Result) {
              PRINC(" from body=")
              Result = Core.F_CALL(C_print,ARGS(_Zbody.ToEID()))
              if !ErrorIn(Result) {
              PRINC(" ")
              Result = F_Generate_breakline_void().ToEID()
              }}
              } else {
              Result = EID{CFALSE.Id(),0}
              } 
            if !ErrorIn(Result) {
            Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).GenerateFunctionStart(self,
              s,
              m.Id(),
              _Znom)
            if !ErrorIn(Result) {
            F_Generate_new_block_void()
            var g0023I *ClaireBoolean
            var try_11 EID
            { 
              var v_or6 *ClaireBoolean
              
              v_or6 = F_Generate_need_debug_ask_any(m.Id())
              if (v_or6 == CTRUE) {try_11 = EID{CTRUE.Id(),0}
              } else { 
                v_or6 = Optimize.C_OPT.Profile_ask
                if (v_or6 == CTRUE) {try_11 = EID{CTRUE.Id(),0}
                } else { 
                  var try_12 EID
                  { var arg_13 *ClaireBoolean
                    var try_14 EID
                    try_14 = F_Generate_simple_body_ask_any(_Zbody)
                    if ErrorIn(try_14) {try_12 = try_14
                    } else {
                    arg_13 = ToBoolean(OBJ(try_14))
                    try_12 = EID{arg_13.Not.Id(),0}
                    }
                    } 
                  if ErrorIn(try_12) {try_11 = try_12
                  } else {
                  v_or6 = ToBoolean(OBJ(try_12))
                  if (v_or6 == CTRUE) {try_11 = EID{CTRUE.Id(),0}
                  } else { 
                    v_or6 = Equal(s.Id(),C_void.Id())
                    if (v_or6 == CTRUE) {try_11 = EID{CTRUE.Id(),0}
                    } else { 
                      try_11 = EID{CFALSE.Id(),0}} 
                    } 
                  } 
                } 
              }
              } 
            if ErrorIn(try_11) {Result = try_11
            } else {
            g0023I = ToBoolean(OBJ(try_11))
            if (g0023I == CTRUE) { 
              Result = p.ProcedureBody(m,
                self,
                _Zbody,
                s)
              } else {
              if (p.Debug_ask == CTRUE) { 
                PRINC("// use function body compiling ")
                F_Generate_breakline_void()
                } 
              Result = Core.F_CALL(C_Generate_function_body,ARGS(EID{p.Id(),0},_Zbody.ToEID(),EID{s.Id(),0}))
              } 
            }
            }}
            } else {
            Core.F_tformat_string(MakeString("--- EID function generation (can throw = ~S) \n"),3,MakeConstantList(throw_ask.Id()))
            throw_ask = CTRUE
            Optimize.C_compiler.NMetheids = (Optimize.C_compiler.NMetheids+1)
            Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).GenerateFunctionStart(self,
              Optimize.C_EID,
              m.Id(),
              _Znom)
            if !ErrorIn(Result) {
            F_Generate_new_block_void()
            Result = F_Generate_eid_body_method(m,_Zbody,ToBoolean(typeOK),s)
            }
            } 
          if !ErrorIn(Result) {
          F_Generate_close_block_void()
          Result = F_Generate_generate_eid_function_lambda(self,m,throw_ask)
          if !ErrorIn(Result) {
          if (m.Selector.Id() == Core.C_self_eval.Id()) { 
            Result = F_Generate_generate_eval_function_lambda(self,m)
            } else {
            Result = EID{CFALSE.Id(),0}
            } 
          if !ErrorIn(Result) {
          Result = ToPort(Reader.C_stdout.Value).UseAsOutput().ToEID()
          }}}}}
          }
          } 
        }
        } 
      } 
    }
    } 
  return Result} 

// The EID go function for: gen_function @ go_producer (throw: true) 
func E_Generate_gen_function_go_producer (p EID,self EID,_Znom EID,m EID) EID { 
  return ToGenerateGoProducer(OBJ(p)).GenFunction(ToLambda(OBJ(self)),
    ToString(OBJ(_Znom)),
    ToMethod(OBJ(m)) )} 

// check that we may call function_body  (replaces the print_body method of CLAIRE 3 compiler)  
// simple : we can generate ... return X directly without the need for a "Result" variable 
/* The go function for: simple_body?(self:any) [status=1] */
func F_Generate_simple_body_ask_any (self *ClaireAny) EID { 
  var Result EID
  if (self.Isa.IsIn(Language.C_If) == CTRUE) { 
    { var g0024 *Language.If = Language.To_If(self)
      { 
        var v_and3 *ClaireBoolean
        
        var try_1 EID
        try_1 = F_Generate_g_func_any(g0024.Test)
        if ErrorIn(try_1) {Result = try_1
        } else {
        v_and3 = ToBoolean(OBJ(try_1))
        if (v_and3 == CFALSE) {Result = EID{CFALSE.Id(),0}
        } else { 
          var try_2 EID
          try_2 = F_Generate_simple_body_ask_any(g0024.Arg)
          if ErrorIn(try_2) {Result = try_2
          } else {
          v_and3 = ToBoolean(OBJ(try_2))
          if (v_and3 == CFALSE) {Result = EID{CFALSE.Id(),0}
          } else { 
            var try_3 EID
            try_3 = F_Generate_simple_body_ask_any(g0024.Other)
            if ErrorIn(try_3) {Result = try_3
            } else {
            v_and3 = ToBoolean(OBJ(try_3))
            if (v_and3 == CFALSE) {Result = EID{CFALSE.Id(),0}
            } else { 
              Result = EID{CTRUE.Id(),0}} 
            } 
          } 
        }}}
        } 
      } 
    }  else if (self.Isa.IsIn(Language.C_Do) == CTRUE) { 
    { var g0025 *Language.Do = Language.To_Do(self)
      { var arg_4 *ClaireAny
        var try_5 EID
        try_5 = Core.F_last_list(g0025.Args)
        if ErrorIn(try_5) {Result = try_5
        } else {
        arg_4 = ANY(try_5)
        Result = F_Generate_simple_body_ask_any(arg_4)
        }
        } 
      } 
    } else {
    Result = F_Generate_g_func_any(self)
    } 
  return Result} 

// The EID go function for: simple_body? @ any (throw: true) 
func E_Generate_simple_body_ask_any (self EID) EID { 
  return F_Generate_simple_body_ask_any(ANY(self) )} 

// simpler case that we apply for Do, Ifs and functional expressions
// however is c_type(exp) is void we need to return CNULL
/* The go function for: function_body(c:go_producer,self:any,s:class) [status=1] */
func F_Generate_function_body_go_producer1 (c *GenerateGoProducer,self *ClaireAny,s *ClaireClass) EID { 
  var Result EID
  { var _Zret *ClaireString = ToString(IfThenElse((s.Id() != C_void.Id()),
      MakeString("return ").Id(),
      MakeString("").Id()))
    var g0028I *ClaireBoolean
    { 
      var v_and2 *ClaireBoolean
      
      v_and2 = Equal(s.Id(),C_boolean.Id())
      if (v_and2 == CFALSE) {g0028I = CFALSE
      } else { 
        if (self.Isa.IsIn(Language.C_Call_method) == CTRUE) { 
          { var g0027 *Language.CallMethod = Language.To_CallMethod(self)
            { var p *ClaireProperty = g0027.Arg.Selector
              v_and2 = MakeBoolean((p.Id() == C__equal.Id()) || (p.Id() == C__inf.Id()) || (p.Id() == C__sup.Id()) || (p.Id() == C__sup_equal.Id()) || (p.Id() == C__inf_equal.Id()))
              } 
            } 
          } else {
          v_and2 = CFALSE
          } 
        if (v_and2 == CFALSE) {g0028I = CFALSE
        } else { 
          g0028I = CTRUE} 
        } 
      } 
    if (g0028I == CTRUE) { 
      PRINC("if ")
      Result = Core.F_CALL(C_Generate_b_expression,ARGS(EID{c.Id(),0},self.ToEID(),EID{CTRUE.Id(),0}))
      if !ErrorIn(Result) {
      PRINC(" {return CTRUE")
      F_Generate_breakline_void()
      PRINC("} else {return CFALSE}")
      Result = EVOID
      }
      } else {
      var g0029I *ClaireBoolean
      var try_1 EID
      { 
        var v_and3 *ClaireBoolean
        
        var try_2 EID
        { var arg_3 *ClaireType
          var try_4 EID
          try_4 = Core.F_CALL(Optimize.C_c_type,ARGS(self.ToEID()))
          if ErrorIn(try_4) {try_2 = try_4
          } else {
          arg_3 = ToType(OBJ(try_4))
          try_2 = EID{Equal(arg_3.Id(),C_void.Id()).Id(),0}
          }
          } 
        if ErrorIn(try_2) {try_1 = try_2
        } else {
        v_and3 = ToBoolean(OBJ(try_2))
        if (v_and3 == CFALSE) {try_1 = EID{CFALSE.Id(),0}
        } else { 
          v_and3 = Core.F__I_equal_any(s.Id(),C_void.Id())
          if (v_and3 == CFALSE) {try_1 = EID{CFALSE.Id(),0}
          } else { 
            try_1 = EID{CTRUE.Id(),0}} 
          } 
        }
        } 
      if ErrorIn(try_1) {Result = try_1
      } else {
      g0029I = ToBoolean(OBJ(try_1))
      if (g0029I == CTRUE) { 
        Result = Core.F_CALL(C_Generate_g_expression,ARGS(self.ToEID(),EID{C_void.Id(),0}))
        if !ErrorIn(Result) {
        F_Generate_breakline_void()
        PRINC("return ")
        Result = F_Generate_g_expression_any(CNULL,s)
        if !ErrorIn(Result) {
        Result = F_Generate_breakline_void().ToEID()
        }}
        } else {
        F_princ_string(_Zret)
        PRINC(" ")
        Result = Core.F_CALL(C_Generate_g_expression,ARGS(self.ToEID(),EID{s.Id(),0}))
        if !ErrorIn(Result) {
        Result = F_Generate_breakline_void().ToEID()
        }
        } 
      }
      } 
    } 
  return Result} 

// The EID go function for: function_body @ list<type_expression>(go_producer, any, class) (throw: true) 
func E_Generate_function_body_go_producer1 (c EID,self EID,s EID) EID { 
  return F_Generate_function_body_go_producer1(ToGenerateGoProducer(OBJ(c)),ANY(self),ToClass(OBJ(s)) )} 

// generate nice code for If function (inspired from g_statement@If)
/* The go function for: function_body(c:go_producer,self:If,s:class) [status=1] */
func F_Generate_function_body_go_producer2 (c *GenerateGoProducer,self *Language.If,s *ClaireClass) EID { 
  var Result EID
  PRINC("if ")
  Result = Core.F_CALL(C_Generate_b_expression,ARGS(EID{c.Id(),0},self.Test.ToEID(),EID{CTRUE.Id(),0}))
  if !ErrorIn(Result) {
  PRINC(" ")
  F_Generate_new_block_string(MakeString("body If"))
  Result = EVOID
  }
  if !ErrorIn(Result) {
  Result = Core.F_CALL(C_Generate_function_body,ARGS(EID{c.Id(),0},self.Arg.ToEID(),EID{s.Id(),0}))
  if !ErrorIn(Result) {
  if (Equal(self.Other,CNIL.Id()) == CTRUE) { 
    F_Generate_close_block_void()
    Result = EVOID
    }  else if (self.Other.Isa.IsIn(Language.C_If) == CTRUE) { 
    F_Generate_finish_block_void()
    PRINC(" else ")
    Result = Core.F_CALL(C_Generate_function_body,ARGS(EID{c.Id(),0},self.Other.ToEID(),EID{s.Id(),0}))
    } else {
    var g0030I *ClaireBoolean
    var try_1 EID
    { 
      var v_or2 *ClaireBoolean
      
      v_or2 = Core.F__I_equal_any(s.Id(),C_void.Id())
      if (v_or2 == CTRUE) {try_1 = EID{CTRUE.Id(),0}
      } else { 
        var try_2 EID
        { var arg_3 *ClaireBoolean
          var try_4 EID
          try_4 = Optimize.F_Compile_designated_ask_any(self.Other)
          if ErrorIn(try_4) {try_2 = try_4
          } else {
          arg_3 = ToBoolean(OBJ(try_4))
          try_2 = EID{arg_3.Not.Id(),0}
          }
          } 
        if ErrorIn(try_2) {try_1 = try_2
        } else {
        v_or2 = ToBoolean(OBJ(try_2))
        if (v_or2 == CTRUE) {try_1 = EID{CTRUE.Id(),0}
        } else { 
          try_1 = EID{CFALSE.Id(),0}} 
        } 
      }
      } 
    if ErrorIn(try_1) {Result = try_1
    } else {
    g0030I = ToBoolean(OBJ(try_1))
    if (g0030I == CTRUE) { 
      PRINC("} else {")
      F_Generate_breakline_void()
      Result = Core.F_CALL(C_Generate_function_body,ARGS(EID{c.Id(),0},self.Other.ToEID(),EID{s.Id(),0}))
      if !ErrorIn(Result) {
      F_Generate_close_block_string(MakeString("body If"))
      Result = EVOID
      }
      } else {
      F_Generate_close_block_string(MakeString("body If"))
      Result = EVOID
      } 
    }
    } 
  }}
  return Result} 

// The EID go function for: function_body @ list<type_expression>(go_producer, If, class) (throw: true) 
func E_Generate_function_body_go_producer2 (c EID,self EID,s EID) EID { 
  return F_Generate_function_body_go_producer2(ToGenerateGoProducer(OBJ(c)),Language.To_If(OBJ(self)),ToClass(OBJ(s)) )} 

// generate nice code for a Do
/* The go function for: function_body(c:go_producer,self:Do,s:class) [status=1] */
func F_Generate_function_body_go_producer3 (c *GenerateGoProducer,self *Language.Do,s *ClaireClass) EID { 
  var Result EID
  { var l *ClaireList = self.Args
    { var _Zlength int = l.Length()
      { var m int = 0
        { 
          var x *ClaireAny
          _ = x
          Result= EID{CFALSE.Id(),0}
          var x_support *ClaireList
          x_support = l
          x_len := x_support.Length()
          for i_it := 0; i_it < x_len; i_it++ { 
            x = x_support.At(i_it)
            var loop_1 EID
            _ = loop_1
            { 
            m = (m+1)
            if (m == _Zlength) { 
              loop_1 = Core.F_CALL(C_Generate_function_body,ARGS(EID{c.Id(),0},x.ToEID(),EID{s.Id(),0}))
              } else {
              loop_1 = F_Generate_statement_any(x,C_void,MakeString("Unused"),CFALSE.Id())
              } 
            if ErrorIn(loop_1) {Result = loop_1
            break
            } else {
            }
            }
            } 
          } 
        } 
      } 
    } 
  return Result} 

// The EID go function for: function_body @ list<type_expression>(go_producer, Do, class) (throw: true) 
func E_Generate_function_body_go_producer3 (c EID,self EID,s EID) EID { 
  return F_Generate_function_body_go_producer3(ToGenerateGoProducer(OBJ(c)),Language.To_Do(OBJ(self)),ToClass(OBJ(s)) )} 

// default complex case : create a variable "Result"
/* The go function for: procedure_body(c:go_producer,m:method,%l:lambda,%body:any,s:class) [status=1] */
func (c *GenerateGoProducer) ProcedureBody (m *ClaireMethod,_Zl *ClaireLambda,_Zbody *ClaireAny,s *ClaireClass) EID { 
  var Result EID
  if (F_Generate_need_debug_ask_any(m.Id()) == CTRUE) { 
    Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).DebugIntro(_Zl,m)
    } else {
    Result = EID{CFALSE.Id(),0}
    } 
  if !ErrorIn(Result) {
  if (ToGenerateGoProducer(Optimize.C_PRODUCER.Value).Debug_ask == CTRUE) { 
    PRINC("// procedure body, with s = ")
    Result = Core.F_print_any(s.Id())
    if !ErrorIn(Result) {
    Result = F_Generate_breakline_void().ToEID()
    }
    } else {
    Result = EID{CFALSE.Id(),0}
    } 
  if !ErrorIn(Result) {
  if (s.Id() != C_void.Id()) { 
    F_Generate_var_declaration_string(MakeString("Result"),s,1)
    Result = F_Generate_statement_any(_Zbody,s,MakeString("Result"),CFALSE.Id())
    } else {
    Result = F_Generate_statement_any(_Zbody,C_void,MakeString("Unused"),CFALSE.Id())
    } 
  if !ErrorIn(Result) {
  Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).ReturnResult(s,m,MakeString("Result"))
  }}}
  return Result} 

// The EID go function for: procedure_body @ go_producer (throw: true) 
func E_Generate_procedure_body_go_producer (c EID,m EID,_Zl EID,_Zbody EID,s EID) EID { 
  return ToGenerateGoProducer(OBJ(c)).ProcedureBody(ToMethod(OBJ(m)),
    ToLambda(OBJ(_Zl)),
    ANY(_Zbody),
    ToClass(OBJ(s)) )} 

// generate an EID function (lambda)    
/* The go function for: eid_body(%body:any,typeOK:boolean,s:class) [status=1] */
func F_Generate_eid_body_any (_Zbody *ClaireAny,typeOK *ClaireBoolean,s *ClaireClass) EID { 
  var Result EID
  F_Generate_var_declaration_string(MakeString("Result"),Optimize.C_EID,1)
  { var arg_1 *ClaireBoolean
    var try_2 EID
    try_2 = Optimize.F_Compile_g_throw_any(_Zbody)
    if ErrorIn(try_2) {Result = try_2
    } else {
    arg_1 = ToBoolean(OBJ(try_2))
    Result = F_Generate_statement_any(_Zbody,Optimize.C_EID,MakeString("Result"),arg_1.Id())
    }
    } 
  if !ErrorIn(Result) {
  PRINC("return Result")
  Result = EVOID
  }
  return Result} 

// The EID go function for: eid_body @ any (throw: true) 
func E_Generate_eid_body_any (_Zbody EID,typeOK EID,s EID) EID { 
  return F_Generate_eid_body_any(ANY(_Zbody),ToBoolean(OBJ(typeOK)),ToClass(OBJ(s)) )} 

// generate an EID body for a method 
// call for the debug/profile is needed     
/* The go function for: eid_body(m:method,%body:any,typeOK:boolean,s:class) [status=1] */
func F_Generate_eid_body_method (m *ClaireMethod,_Zbody *ClaireAny,typeOK *ClaireBoolean,s *ClaireClass) EID { 
  var Result EID
  if (F_Generate_need_debug_ask_any(m.Id()) == CTRUE) { 
    Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).DebugIntro(m.Formula,m)
    } else {
    Result = EID{CFALSE.Id(),0}
    } 
  if !ErrorIn(Result) {
  if (ToGenerateGoProducer(Optimize.C_PRODUCER.Value).Debug_ask == CTRUE) { 
    PRINC("// eid body s = ")
    Result = Core.F_print_any(s.Id())
    if !ErrorIn(Result) {
    Result = F_Generate_breakline_void().ToEID()
    }
    } else {
    Result = EID{CFALSE.Id(),0}
    } 
  if !ErrorIn(Result) {
  F_Generate_var_declaration_string(MakeString("Result"),Optimize.C_EID,1)
  { var arg_1 *ClaireBoolean
    var try_2 EID
    try_2 = Optimize.F_Compile_g_throw_any(_Zbody)
    if ErrorIn(try_2) {Result = try_2
    } else {
    arg_1 = ToBoolean(OBJ(try_2))
    Result = F_Generate_statement_any(_Zbody,Optimize.C_EID,MakeString("Result"),arg_1.Id())
    }
    } 
  if !ErrorIn(Result) {
  if (F_Generate_need_debug_ask_any(m.Id()) == CTRUE) { 
    Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).ReturnResult(Optimize.C_EID,m,MakeString("Result"))
    }  else if ((typeOK == CTRUE) || 
      (Optimize.C_compiler.Safety >= 2)) { 
    PRINC("return Result")
    Result = EVOID
    } else {
    PRINC("return RangeCheck(")
    Result = F_Generate_g_expression_class(s,C_type)
    if !ErrorIn(Result) {
    PRINC(",Result)")
    Result = EVOID
    }
    } 
  }}}
  return Result} 

// The EID go function for: eid_body @ method (throw: true) 
func E_Generate_eid_body_method (m EID,_Zbody EID,typeOK EID,s EID) EID { 
  return F_Generate_eid_body_method(ToMethod(OBJ(m)),
    ANY(_Zbody),
    ToBoolean(OBJ(typeOK)),
    ToClass(OBJ(s)) )} 

// generate the EID function associated to each method (used by the interpreter - EID mode)
/* The go function for: generate_eid_function(self:lambda,m:method,throw?:boolean) [status=1] */
func F_Generate_generate_eid_function_lambda (self *ClaireLambda,m *ClaireMethod,throw_ask *ClaireBoolean) EID { 
  var Result EID
  { var _Zsig *ClaireList = F_Generate_go_signature_method(m)
    { var lv *ClaireList = self.Vars
      PRINC("\n// The EID go function for: ")
      Result = Core.F_print_any(m.Id())
      if !ErrorIn(Result) {
      PRINC(" (throw: ")
      Result = Core.F_print_any(throw_ask.Id())
      if !ErrorIn(Result) {
      PRINC(") \n")
      Result = EVOID
      }}
      if !ErrorIn(Result) {
      PRINC("func ")
      Result = F_Generate_goEIDFunctionName_method(m)
      if !ErrorIn(Result) {
      PRINC(" (")
      ToGenerateGoProducer(Optimize.C_PRODUCER.Value).GoEIDVariables(lv)
      PRINC(") EID ")
      Result = EVOID
      }
      if !ErrorIn(Result) {
      F_Generate_new_block_void()
      if ((m.Range.Id() == C_void.Id()) && 
          (throw_ask != CTRUE)) { 
        Result = F_Generate_print_EID_call_method(m,lv,_Zsig,throw_ask)
        if !ErrorIn(Result) {
        F_Generate_breakline_void()
        PRINC("return EVOID")
        Result = EVOID
        }
        } else {
        PRINC("return ")
        Result = F_Generate_print_EID_call_method(m,lv,_Zsig,throw_ask)
        } 
      if !ErrorIn(Result) {
      F_Generate_close_block_void()
      Result = EVOID
      }}}
      } 
    } 
  return Result} 

// The EID go function for: generate_eid_function @ lambda (throw: true) 
func E_Generate_generate_eid_function_lambda (self EID,m EID,throw_ask EID) EID { 
  return F_Generate_generate_eid_function_lambda(ToLambda(OBJ(self)),ToMethod(OBJ(m)),ToBoolean(OBJ(throw_ask)) )} 

// similar but simpler for a lambda associated to a name (e.g. 2nd order types) => E_C(nom)      
/* The go function for: generate_eid_dual(self:lambda,%nom:string) [status=1] */
func F_Generate_generate_eid_dual_lambda (self *ClaireLambda,_Znom *ClaireString) EID { 
  var Result EID
  { var lv *ClaireList = self.Vars
    { var nl_ask *ClaireBoolean = Core.F__sup_integer(lv.Length(),3)
      PRINC("\n// The dual EID go function for: ")
      Result = Core.F_print_any((_Znom).Id())
      if !ErrorIn(Result) {
      PRINC(" \n")
      Result = EVOID
      }
      if !ErrorIn(Result) {
      PRINC("func E_")
      F_c_princ_string(_Znom)
      PRINC(" (")
      ToGenerateGoProducer(Optimize.C_PRODUCER.Value).GoEIDVariables(lv)
      PRINC(") EID ")
      F_Generate_new_block_void()
      PRINC("return F_")
      F_c_princ_string(_Znom)
      PRINC("(")
      { var n int = 1
        { var g0031 int = lv.Length()
          Result= EID{CFALSE.Id(),0}
          for (n <= g0031) { 
            var loop_1 EID
            _ = loop_1
            { 
            loop_1 = F_Generate_external_EID_arg_Variable(To_Variable(lv.At(n-1)),ToTypeExpression(OBJ(Core.F_CALL(C_range,ARGS(lv.At(n-1).ToEID())))).Class_I(),n,nl_ask)
            if ErrorIn(loop_1) {Result = loop_1
            break
            } else {
            n = (n+1)
            }
            } 
          }
          } 
        } 
      if !ErrorIn(Result) {
      PRINC(")")
      F_Generate_close_block_void()
      Result = EVOID
      }}
      } 
    } 
  return Result} 

// The EID go function for: generate_eid_dual @ lambda (throw: true) 
func E_Generate_generate_eid_dual_lambda (self EID,_Znom EID) EID { 
  return F_Generate_generate_eid_dual_lambda(ToLambda(OBJ(self)),ToString(OBJ(_Znom)) )} 

// EID function calls the compiled native function - uses a code that looks like print_external_call
// watch out: a method that can throw returns an EID directly ! (same as goexp.cl : print_ext_call)
/* The go function for: print_EID_call(m:method,l:list,%sig:list<class>,throw?:boolean) [status=1] */
func F_Generate_print_EID_call_method (m *ClaireMethod,l *ClaireList,_Zsig *ClaireList,throw_ask *ClaireBoolean) EID { 
  var Result EID
  { var n int = 1
    _ = n
    { var sm *ClaireClass
      var try_1 EID
      try_1 = Core.F_last_list(_Zsig)
      if ErrorIn(try_1) {Result = try_1
      } else {
      sm = ToClass(OBJ(try_1))
      { var nl_ask *ClaireBoolean = Core.F__sup_integer(l.Length(),3)
        if (nl_ask == CTRUE) { 
          Optimize.C_OPT.Level = (Optimize.C_OPT.Level+1)
          } 
        if ((throw_ask == CTRUE) || 
            ((ToSet(C_Generate_EIDSET.Value).Contain_ask(m.Id()) == CTRUE) || 
              (m.Selector.Id() == Core.C_self_eval.Id()))) { 
          sm = Optimize.C_EID
          } 
        Result = F_Generate_cast_prefix_class(sm,Optimize.C_EID)
        if !ErrorIn(Result) {
        if (ToGenerateGoProducer(Optimize.C_PRODUCER.Value).Debug_ask == CTRUE) { 
          PRINC("/*(sm for ")
          Result = Core.F_print_any(m.Id())
          if !ErrorIn(Result) {
          PRINC("= ")
          Result = Core.F_print_any(sm.Id())
          if !ErrorIn(Result) {
          PRINC(")*/ ")
          Result = EVOID
          }}
          } else {
          Result = EID{CFALSE.Id(),0}
          } 
        if !ErrorIn(Result) {
        if (F_Generate_goMethod_ask_any(m.Id()) == CTRUE) { 
          Result = F_Generate_external_EID_arg_Variable(To_Variable(l.At(0)),ToClass(_Zsig.ValuesO()[0]),1,nl_ask)
          if !ErrorIn(Result) {
          PRINC(".")
          F_Generate_goMethod_method(m)
          PRINC("(")
          { var n int = 2
            { var g0034 int = l.Length()
              Result= EID{CFALSE.Id(),0}
              for (n <= g0034) { 
                var loop_2 EID
                _ = loop_2
                { 
                loop_2 = F_Generate_external_EID_arg_Variable(To_Variable(l.At(n-1)),ToClass(_Zsig.ValuesO()[n-1]),(n-1),nl_ask)
                if ErrorIn(loop_2) {Result = loop_2
                break
                } else {
                n = (n+1)
                }
                } 
              }
              } 
            } 
          }
          } else {
          Result = F_Generate_goFunction_method(m)
          if !ErrorIn(Result) {
          PRINC("(")
          Result = EVOID
          }
          if !ErrorIn(Result) {
          if ((l.Length() == 1) && 
              (Core.F_domain_I_restriction(ToRestriction(m.Id())).Id() == C_void.Id())) { 
            l = CNIL
            } 
          { var n int = 1
            { var g0035 int = l.Length()
              Result= EID{CFALSE.Id(),0}
              for (n <= g0035) { 
                var loop_3 EID
                _ = loop_3
                { 
                loop_3 = F_Generate_external_EID_arg_Variable(To_Variable(l.At(n-1)),ToClass(_Zsig.ValuesO()[n-1]),n,nl_ask)
                if ErrorIn(loop_3) {Result = loop_3
                break
                } else {
                n = (n+1)
                }
                } 
              }
              } 
            } 
          }
          } 
        if !ErrorIn(Result) {
        PRINC(" )")
        if (nl_ask == CTRUE) { 
          Optimize.C_OPT.Level = (Optimize.C_OPT.Level-1)
          } 
        F_Generate_cast_post_class(sm,Optimize.C_EID)
        Result = EVOID
        }}}
        } 
      }
      } 
    } 
  return Result} 

// The EID go function for: print_EID_call @ method (throw: true) 
func E_Generate_print_EID_call_method (m EID,l EID,_Zsig EID,throw_ask EID) EID { 
  return F_Generate_print_EID_call_method(ToMethod(OBJ(m)),
    ToList(OBJ(l)),
    ToList(OBJ(_Zsig)),
    ToBoolean(OBJ(throw_ask)) )} 

// here v is a EID-range variable and we need to extract the native s representation
// n=0 is a special marker when the arg the receiver x in x.f(....)
/* The go function for: external_EID_arg(v:Variable,s:class,n:integer,nl?:boolean) [status=1] */
func F_Generate_external_EID_arg_Variable (v *ClaireVariable,s *ClaireClass,n int,nl_ask *ClaireBoolean) EID { 
  var Result EID
  if (n > 1) { 
    PRINC(",")
    if (nl_ask == CTRUE) { 
      F_Generate_breakline_void()
      } 
    } 
  Result = F_Generate_eid_prefix_class(s)
  if !ErrorIn(Result) {
  F_iClaire_ident_go_producer1(ToGenerateGoProducer(Optimize.C_PRODUCER.Value),v)
  F_Generate_eid_post_class(s)
  Result = EVOID
  }
  return Result} 

// The EID go function for: external_EID_arg @ Variable (throw: true) 
func E_Generate_external_EID_arg_Variable (v EID,s EID,n EID,nl_ask EID) EID { 
  return F_Generate_external_EID_arg_Variable(To_Variable(OBJ(v)),
    ToClass(OBJ(s)),
    INT(n),
    ToBoolean(OBJ(nl_ask)) )} 

// prints a list of arguments with types / replaces typed_args_list
/* The go function for: goEIDVariables(p:go_producer,self:list) [status=0] */
func (p *GenerateGoProducer) GoEIDVariables (self *ClaireList) *ClaireAny { 
  var Result *ClaireAny
  { var prems *ClaireBoolean = CTRUE
    { 
      var v *ClaireVariable
      _ = v
      var v_iter *ClaireAny
      Result= CFALSE.Id()
      var v_support *ClaireList
      v_support = self
      v_len := v_support.Length()
      for i_it := 0; i_it < v_len; i_it++ { 
        v_iter = v_support.At(i_it)
        v = To_Variable(v_iter)
        if (prems == CTRUE) { 
          prems = CFALSE
          } else {
          PRINC(",")
          } 
        F_iClaire_ident_go_producer1(p,v)
        PRINC(" EID")
        } 
      } 
    } 
  return Result} 

// The EID go function for: goEIDVariables @ go_producer (throw: false) 
func E_Generate_goEIDVariables_go_producer (p EID,self EID) EID { 
  return ToGenerateGoProducer(OBJ(p)).GoEIDVariables(ToList(OBJ(self)) ).ToEID()} 

// check the range & sort of the method through type inference. 
// returns true if OK and false otherwise (can produce an error at run-time)
// notice that %body is the lambda body before compilation => use c_type
/* The go function for: check_range(self:method,%body:any) [status=1] */
func F_Generate_check_range_method (self *ClaireMethod,_Zbody *ClaireAny) EID { 
  var Result EID
  { var s1 *ClaireClass = self.Range.Class_I()
    { var s2 *ClaireClass
      var try_1 EID
      { var arg_2 *ClaireType
        var try_3 EID
        try_3 = Core.F_CALL(Optimize.C_c_type,ARGS(_Zbody.ToEID()))
        if ErrorIn(try_3) {try_1 = try_3
        } else {
        arg_2 = ToType(OBJ(try_3))
        try_1 = EID{arg_2.Class_I().Id(),0}
        }
        } 
      if ErrorIn(try_1) {Result = try_1
      } else {
      s2 = ToClass(OBJ(try_1))
      if ((s1.Id() == C_void.Id()) || 
          (ToType(s2.Id()).Included(ToType(s1.Id())) == CTRUE)) { 
        Result = EID{CTRUE.Id(),0}
        } else {
        Optimize.F_Compile_notice_void()
        C_ABODY.Value = _Zbody
        Core.F_tformat_string(MakeString("~S's range was found to be ~S (vs. ~S) \n"),2,MakeConstantList(self.Id(),s2.Id(),s1.Id()))
        { var arg_4 *ClaireList
          var try_5 EID
          { 
            var v_bag_arg *ClaireAny
            try_5= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
            ToList(OBJ(try_5)).AddFast(_Zbody)
            var try_6 EID
            try_6 = Core.F_CALL(Optimize.C_c_type,ARGS(_Zbody.ToEID()))
            if ErrorIn(try_6) {try_5 = try_6
            } else {
            v_bag_arg = ANY(try_6)
            ToList(OBJ(try_5)).AddFast(v_bag_arg)}
            } 
          if ErrorIn(try_5) {Result = try_5
          } else {
          arg_4 = ToList(OBJ(try_5))
          Result = Core.F_tformat_string(MakeString("BODY is ~S -> type=~S \n"),2,arg_4)
          }
          } 
        if !ErrorIn(Result) {
        if (((s1.Id() != C_void.Id()) && 
              ((s2.Id() == C_void.Id()) && 
                (s1.Id() != C_error.Id()))) || 
            (Equal(Core.F__exp_type(ToType(s1.Id()),ToType(s2.Id())).Id(),CEMPTY.Id()) == CTRUE)) { 
          Result = Optimize.F_Compile_Cerror_string(MakeString("[218] Sort error: Cannot compile ~S (~S cannot be ~S)."),ToList(MakeConstantList(self.Id(),s1.Id(),s2.Id()).Id()))
          } else {
          Result = EID{CFALSE.Id(),0}
          } 
        }
        } 
      }
      } 
    } 
  return Result} 

// The EID go function for: check_range @ method (throw: true) 
func E_Generate_check_range_method (self EID,_Zbody EID) EID { 
  return F_Generate_check_range_method(ToMethod(OBJ(self)),ANY(_Zbody) )} 

// generate the eval function associated to each self_eval method (type *any -> EID)
// EVAL_C(x *ClaireAny) EID {return ToC(x).SelfEval()}
/* The go function for: generate_eval_function(self:lambda,m:method) [status=1] */
func F_Generate_generate_eval_function_lambda (self *ClaireLambda,m *ClaireMethod) EID { 
  var Result EID
  { var c *ClaireClass = Core.F_domain_I_restriction(ToRestriction(m.Id()))
    if (c.Id() != C_Variable.Id()) { 
      PRINC("\n// The EVAL go function for: ")
      Result = Core.F_print_any(c.Id())
      if !ErrorIn(Result) {
      PRINC(" \n")
      Result = EVOID
      }
      if !ErrorIn(Result) {
      PRINC("func EVAL_")
      c.Name.CPrinc()
      PRINC(" (x *ClaireAny) EID ")
      F_Generate_new_block_void()
      if (F_Generate_goMethod_ask_any(m.Id()) == CTRUE) { 
        PRINC(" return ")
        F_Generate_cast_class_class(c)
        PRINC("(x).SelfEval()")
        } else {
        PRINC(" return F_self_eval_")
        c.Name.CPrinc()
        PRINC("(")
        F_Generate_cast_class_class(c)
        PRINC("(x))")
        } 
      F_Generate_close_block_void()
      Result = EVOID
      }
      } else {
      Result = EID{CFALSE.Id(),0}
      } 
    } 
  return Result} 

// The EID go function for: generate_eval_function @ lambda (throw: true) 
func E_Generate_generate_eval_function_lambda (self EID,m EID) EID { 
  return F_Generate_generate_eval_function_lambda(ToLambda(OBJ(self)),ToMethod(OBJ(m)) )} 

// tells if a method needs debug instrumentation
/* The go function for: need_debug?(m:any) [status=0] */
func F_Generate_need_debug_ask_any (m *ClaireAny) *ClaireBoolean { 
  var Result *ClaireBoolean
  if (C_method.Id() == m.Isa.Id()) { 
    { var g0036 *ClaireMethod = ToMethod(m)
      { var p *ClaireProperty = g0036.Selector
        Result = MakeBoolean((Optimize.C_compiler.Debug_ask.Memq(g0036.Module_I.Id()) == CTRUE) && 
        (g0036.Module_I.Id() != C_claire.Id()) && 
        (p.Id() != Core.C_self_eval.Id()) && 
        (p.Id() != Core.C_execute.Id()) && 
        (p.Id() != Core.C_eval_message.Id()) && 
        (p.Id() != Core.C_Core_push_debug.Id()) && 
        (p.Id() != Core.C_Core_pop_debug.Id()) && 
        (p.Id() != Core.C_Core_tr_indent.Id()) && 
        (p.Id() != Core.C_Core_find_which.Id()) && 
        (p.Id() != Core.C_Core_matching_ask.Id()) && 
        (p.Id() != Core.C_Core_vmatch_ask.Id()))
        } 
      } 
    } else {
    Result = CFALSE
    } 
  return Result} 

// The EID go function for: need_debug? @ any (throw: false) 
func E_Generate_need_debug_ask_any (m EID) EID { 
  return EID{F_Generate_need_debug_ask_any(ANY(m) ).Id(),0}} 

// produce the debugging code introduction
// db_bind is defined in  method.cl
/* The go function for: debug_intro(c:go_producer,self:lambda,x:method) [status=1] */
func (c *GenerateGoProducer) DebugIntro (self *ClaireLambda,x *ClaireMethod) EID { 
  var Result EID
  { var m *ClaireModule
    if (C_method.Id() == x.Isa.Id()) { 
      m = x.Module_I
      } else {
      m = ToModule(CFALSE.Id())
      } 
    { var n int = 1
      PRINC("Core.F_Core_db_bind_module(")
      Result = F_Generate_g_expression_module(m,C_module)
      if !ErrorIn(Result) {
      PRINC(",")
      Result = F_Generate_g_expression_thing(ToThing(x.Selector.Id()),C_property)
      if !ErrorIn(Result) {
      PRINC(",ARGS(")
      if ((self.Vars.Length() == 1) && 
          (ANY(Core.F_CALL(C_range,ARGS(self.Vars.At(0).ToEID()))) == C_void.Id())) { 
        PRINC("EID{ClEnv.Id(),0}")
        Result = EVOID
        } else {
        { 
          var v *ClaireAny
          _ = v
          Result= EID{CFALSE.Id(),0}
          var v_support *ClaireList
          v_support = self.Vars
          v_len := v_support.Length()
          for i_it := 0; i_it < v_len; i_it++ { 
            v = v_support.At(i_it)
            var loop_1 EID
            _ = loop_1
            { 
            if (n > 1) { 
              PRINC(",")
              } 
            loop_1 = c.ToEid(v,ToClass(F_Generate_go_signature_method(x).ValuesO()[n-1]))
            if ErrorIn(loop_1) {Result = loop_1
            break
            } else {
            n = (n+1)
            }
            }
            } 
          } 
        } 
      if !ErrorIn(Result) {
      PRINC("));")
      Result = F_Generate_breakline_void().ToEID()
      }}}
      } 
    } 
  return Result} 

// The EID go function for: debug_intro @ go_producer (throw: true) 
func E_Generate_debug_intro_go_producer (c EID,self EID,x EID) EID { 
  return ToGenerateGoProducer(OBJ(c)).DebugIntro(ToLambda(OBJ(self)),ToMethod(OBJ(x)) )} 

// auxiliary to produce the end statement for the function. s tells if the result is needed.
// generates a "... return" if the result is needed or just an empy string
// we also add the debugging unbind if needed.  (used to be called protect_result)
/* The go function for: return_result(p:go_producer,s:class,x:method,%res:string) [status=1] */
func (p *GenerateGoProducer) ReturnResult (s *ClaireClass,x *ClaireMethod,_Zres *ClaireString) EID { 
  var Result EID
  if (F_Generate_need_debug_ask_any(x.Id()) == CTRUE) { 
    PRINC("Core.F_Core_db_unbind_module(")
    Result = F_Generate_g_expression_module(x.Module_I,C_module)
    if !ErrorIn(Result) {
    PRINC(",")
    Result = F_Generate_g_expression_thing(ToThing(x.Selector.Id()),C_property)
    if !ErrorIn(Result) {
    PRINC(",")
    { var arg_1 *ClaireAny
      if (s.Id() == C_void.Id()) { 
        arg_1 = CNULL
        } else {
        arg_1 = F_Generate_build_Variable_string(MakeString("Result"),s.Id()).Id()
        } 
      Result = Core.F_CALL(C_Generate_g_expression,ARGS(arg_1.ToEID(),EID{C_any.Id(),0}))
      } 
    if !ErrorIn(Result) {
    PRINC(")")
    Result = F_Generate_breakline_void().ToEID()
    }}}
    } else {
    Result = EID{CFALSE.Id(),0}
    } 
  if !ErrorIn(Result) {
  if (s.Id() != C_void.Id()) { 
    PRINC("return ")
    F_princ_string(_Zres)
    Result = EVOID
    } else {
    Result = EID{CFALSE.Id(),0}
    } 
  }
  return Result} 

// The EID go function for: return_result @ go_producer (throw: true) 
func E_Generate_return_result_go_producer (p EID,s EID,x EID,_Zres EID) EID { 
  return ToGenerateGoProducer(OBJ(p)).ReturnResult(ToClass(OBJ(s)),
    ToMethod(OBJ(x)),
    ToString(OBJ(_Zres)) )} 

// prints a function name without the # syntactic marker for imported
/* The go function for: c_princ(self:function) [status=0] */
func F_c_princ_function (self *ClaireFunction)  { 
  F_Generate_import_princ_string(F_string_I_function(self))
  } 

// The EID go function for: c_princ @ function (throw: false) 
func E_c_princ_function (self EID) EID { 
  F_c_princ_function(ToFunction(OBJ(self)) )
  return EVOID} 

/* The go function for: import_princ(s:string) [status=0] */
func F_Generate_import_princ_string (s *ClaireString)  { 
  { var i int = 1
    { var g0039 int = F_length_string(s)
      for (i <= g0039) { 
        if ((i > 1) || 
            (s.At(i) != '#')) { 
          F_c_princ_char(s.At(i))
          } 
        i = (i+1)
        } 
      } 
    } 
  } 

// The EID go function for: import_princ @ string (throw: false) 
func E_Generate_import_princ_string (s EID) EID { 
  F_Generate_import_princ_string(ToString(OBJ(s)) )
  return EVOID} 

// end of file