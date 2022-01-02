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
[external!(m:module) : string
 -> if known?(external,m) m.external else string!(m.name) ]

[string2module(s:string) : module
  -> let m := get_value(s) in
       case m (module m,
               any error("~A is not a module",s)) ]

// *******************************************************************
// *       Part 1: definition of the main function                   *
// *******************************************************************

// help file
[printHelp() : void
 -> printf("------------- CLAIRE: The Art of Elegant Programming -----------\n\n"),
    about(),
    printf("\noptions -s <int> : set memory allocation size  \n"),
    printf("        -f <filename>  : load <filename>             \n"),
    printf("        -n : do not load the init file               \n"),
    printf("        -m <module> : load <module>                  \n"),
    printf("        -mx <module> : load <module> and launch main()  \n"),
    printf("        -v <int> : sets the verbosity level       \n"),
    printf("        -S <flag> : sets the global variable <flag> to true  \n"),
    printf("        -o <name> : sets the name of the executable  \n"),
    printf("        -od <name> : sets the output directory  \n"),
    printf("        -p : profiling mode                          \n"),
    printf("        -D : debug mode                              \n"),
    printf("        -safe : safe mode                            \n"),
    printf("        -O : optimizing mode                         \n"),
    printf("        -O2 : super optimizing mode  (no bound checks)      \n"),
    printf("        -cm <module>: compiles a module -> executable        \n"),
    printf("        -cc <module>: compiles a module -> target go files      \n"),
    printf("        -cx <module> : compiles a module & launch main() \n"),
    printf("         -sf <module> : generates system file associated to a module \n"),
    printf("         -sx <module> : generates system file that includes main() \n"),
    exit(0) ]


// Claire's main - the rich version that starts the compiler
// a simpler version exists in inspect.cl (simple_main())
// -s is ignored because it is trapped earlier (see the file generator)
[complex_main() : void
  -> let %cm := "", %sf := "", %out := "",        // names of files/modules
         dblevel := 1, vlevel := 1,               // defaults that can be overiden
         %init? := true, %exe := false,           // do we want to load init.cl ? an executable ?
         %safety := unknown,                      // safety level for compiler
         %main := false,                           // do we cant to call main() ?
         l := (copy(params()) as list<string>) in  // args list
 (try
  (while (l)
   (case l[1]
    ({"?", "-help"} printHelp(),
     {"-q"} (vlevel := 0, l :<< 1),
     {"-v"} (vlevel := 2, l :<< 1),
     {"-s"}  (if (length(l) >= 2)  l :<< 2 else error("option: -s <s1> <s2>")),
     {"-f"}  (if (length(l) >= 2)  (load(l[2]), l :<< 2)
              else error("option: -f <filename>")),
     {"-m"}  (if (length(l) >= 2)
               (if %init? (load("init"), %init? := false),
                let m := string2module(l[2]) in
                  (load(m), begin(m), l :<< 2 , claire_modules :add m))
              else error("option: -m <module>")),
     {"-mx"}  (if (length(l) >= 2)
               (if %init? (load("init"), %init? := false),
                let m := string2module(l[2]) in
                  (load(m), call(main,list()), l :<< 2 , claire_modules :add m))
              else error("option: -m <module>")),
     {"-v"} (if (length(l) >= 2) (vlevel :+ integer!(l[2]), l :<< 2)
             else error("option: -v <integer>")),
     {"-S"} (if (length(l) >= 2) 
                (value(new(global_variable,symbol!(l[2]))) := true,
                 l :<< 2)
             else error("option: -S <FLAG>")),
     {"-od"} (if (length(l) >= 2) (compiler.source := l[2], l :<< 2)
              else error("option: -od <directory>")),
     {"-o"} (if (length(l) >= 2) (%out := l[2], l :<< 2) 
             else error("option: -o <name>")),
     {"-p"} (OPT.Compile/profile? := true, dblevel :max 2, l :<< 1),
     {"-D"} (dblevel := 0, l :<< 1),
     {"-O"} (compiler.optimize? := true, %safety := 2, dblevel := 2, l :<< 1),
     {"-O2"} (compiler.optimize? := true, %safety := 3, dblevel := 2, l :<< 1),
     {"-cc"} (if (length(l) >= 2)  (%cm := l[2], l :<< 2)
               else error("option: -cc <module>")),
     {"-cm"}  (if (length(l) >= 2)   (%exe := true, %cm := l[2], l :<< 2)
               else error("option: -cm <module>")),
     {"-cx"}  (if (length(l) >= 2)   (%exe := true, %cm := l[2], %main := true, l :<< 2)
               else error("option: -cm <module>")),
     {"-sf"} (if (length(l) >= 2)   (%sf := l[2], l :<< 2)
               else error("option: -sf <filename>")),
     {"-sx"} (if (length(l) >= 2)   (%sf := l[2], %main := true, l :<< 2)
               else error("option: -sx <filename>")),
     {"-n"} (%init? := false, l :<< 1),
     any (if (l[1][1] = '-') (printf("~S is an unvalid option\n",l[1]),
                              printHelp()),
           l := list<string>() ) )),
   if (%out = "") (if (%cm != "") %out := %cm else if (%sf != "") %out := %sf),
   if %init? load("init"), 
   system.verbose := vlevel,
   if known?(%safety) compiler.safety := %safety,     // if changed through -O*
   if (%sf != "")                                // we want a system file
      (load(get_value("Compile")),
       compiler.active? := true,
       claire/system_file(string2module(%sf),%out,%main),
       exit(0))
   else if (%cm != "")                        // we have asked to compile a module
      let m := string2module(%cm) in
        (if (m.uses = list(claire_modules[2]))
            (claire_modules := shrink(claire_modules,2),
             trace(0,"=== Light Module ~S:~S -> use ~S=== ", m, m.uses,claire_modules)),
         claire_modules :add m,
         load(get_value("Compile")),     // load the compiler
         compiler.active? := true,
	       if (%out != "") external(m) := %out,
	       load(m),                        // load the module
         if (dblevel < 1) (compiler.safety :min 1, compiler.debug? :add m),
         compile_dir(m),
         compile(m),
         if (%exe) 
            (//[0] ==== create the systel file for module ~S // %out,
             claire/system_file(m,%out,%main),   // v3.2.12: level = 0 => do nothing ....
             compile_exe(%out)),
         exit(0))
    else Reader/top_level(reader)
    )                                                                
   catch any (mClaire/restore_state(reader),
              printf("\nCLAIRE error during init [line ~I]:\n",princ(n_line())),
              Reader/debug_if_possible(), // print_exception(),
              princ("\n"),
              Reader/top_level(reader))) ]


// *******************************************************************
// *       Part 2: System compiling methods                          *
// *******************************************************************

// generate a system file with
//   - the import
//   - the module definition
//   - calling the load() methods for the meta-descriptions
//   - the main function
[claire/system_file(m:module,%out:string,%main:boolean) : void
 -> let p := fopen((compiler.source / %out) /+ PRODUCER.extension, "w"),
        l_used:list := Reader/add_modules(list(m)),
        l_necessary:list := parents(l_used) in
       (// setup the Optimizer
        PRODUCER.current := claire,
        OPT.properties := set<property>(),
        OPT.objects := list<any>(),
        OPT.functions := list<any>(),
        OPT.need_to_close := set<any>(),
        OPT.legal_modules := set!(l_necessary),
        // start the file
        use_as_output(p),
        printf("// --- System configuration file for ~S , [~S] ---\n\n", %out, date!(1)),
        printf("package main\n"),
        system_imports(m),
        load_function(m,l_necessary),
        main_function(m,l_used,%main),
        fclose(p)) ]

// create the import declaration for this system file
[system_imports(m:module) : void
  ->  printf("import (\n"),
      // printf("\t\"fmt\"\n"),                // not necessary if no printf statement is in the code
      if OPT.profile? printf("\t\"os\"\n\t\"runtime/pprof\"\n"),
      import_declaration(m),
      printf("\t\"~A\"\n",string!(m.name)),
      printf(")\n")]
    

// called by the run_system function. The goal of this method is
// to create the modules, and then loads all the meta-descriptions with m.MetaLoad()
// note that is status(m) = 5, we defer the load of the modules.
// we should have a declaration like deferred(m) :: m.status := 5
[load_function(m:module, l_necessary:list) : void
 -> for x in {m in l_necessary | m.status = 5}
       printf("func load_~I() {~I.metaLoad();}\n",ident(x.name), ident(x.name)),
    printf("\n//load function : create and load modules~I", breakline()),
    printf("func Load() ~I",new_block()),
    printf("It := C_claire~I",breakline()),
    printf("//module definitions ~I", breakline()),
    for x in (l_necessary)
      (if not(x % {claire,mClaire,Kernel})
       printf("~I = InitModule(~S,~I,~I,~I\t~S,\n\t~I)~I", 
              g_expression(x,module),
              string!(x.name),
              g_expression(x.part_of,module),
              g_expression(c_code(x.uses,list),list),
              breakline(),
              source(x),
              g_expression(c_code(x.made_of,list),list),
              breakline())),
    printf("~I// module load ~I", breakline(), breakline()),
    for x in list{m in l_necessary | m.made_of & m.status != 5}
       printf("~I.MetaLoad()~I",ident(x.name),breakline()),
    for x in list{m in l_necessary | m.status = 5}
      ( printf("~I.it->evaluate = ~I~I",ident(x.name),
               g_expression(make_function("load_" /+ string!(x.name)), function),
               breakline()),
        printf("~I.it->status = 2;~I",ident(x.name), breakline())),
    printf("ClEnv.Module_I = ~I; ~I", g_expression( m, module), breakline()),
    close_block() ]


// create the main function
// %main = true means call main()
[main_function(m:module,l_used:list[module],%main:boolean) : void
 -> // stuff that is useful to parse
    printf("\n// the main function \n"),
    printf("func main() ~I",new_block()),
    printf("MemoryFlags()~I",breakline()),
    if OPT.profile?
       (printf("// instruction for GO profiling - to be used with go tool pprof <m.prof>\n"),
        new_block(),
        printf("f,err := os.Create(\"~A.prof\")~I",string!(m.name),breakline()),
        printf("if err == nil "),
        new_block(),
        printf("pprof.StartCPUProfile(f)~Idefer pprof.StopCPUProfile()",breakline()),
        close_block(),
        close_block()),
    // printf("fmt.Printf(\"=== CLAIRE4 interpreter version 1.0    ===\\n\")~I",breakline()),
    printf("Bootstrap()~I",breakline()),
    printf("Load()~I",breakline()),
    if (get_value("Generate") % l_used)
        printf("ClEnv.Module_I = C_claire~I",breakline()),
	  printf("Reader.C_reader.Fromp = ClEnv.Cin~I",breakline()),
    if %main printf("Core.F_CALL(Core.C_main,ARGS(EID{ClEnv.Id(),0}))")
    else if (get_value("Generate") % l_used) printf("Generate.F_Generate_complex_main_void()")
    else printf("Reader.F_Reader_simple_main_void()"),
    breakline(),
    close_block() ]

// *******************************************************************
// *       Part 3: module compiling : execute a command line         *
// *******************************************************************


// create a directory for the module (if it does not exist)
[compile_dir(m:module): void
 -> let s := "mkdir -p src" / capitalize(string!(m.name)) in 
     (//[0] ask shell : ~S // s,
      shell(s))]

// create the go
[compile_exe(%out:string): void
 -> let s := "go build src" / %out /+ ".go" in 
     (//[0] ask shell : ~S // s,
      shell(s))]



// *******************************************************************
// *       Part 4: Profiler code                                     *
// *******************************************************************

// ---------------------------------------------------------------




