//+-------------------------------------------------------------+
//| CLAIRE                                                      |
//| ccmain.cl                                                   |
//| Copyright (C) 1994 - 2003 Yves Caseau. All Rights Reserved  |
//| cf. copyright info in file object.cl: about()               |
//+-------------------------------------------------------------+

// ----------------------------------------------------------------------
// this file contains the claire definition of the main for the cclaire
// executable. The four envt are supported ntw, ntv, unix, win32v
// ----------------------------------------------------------------------


// -------------------------------------------------------------------
// Contents
//      Part 1: definition of system variables
//      Part 2: definition of the main function
//      Part 3: Generating makefiles
// -------------------------------------------------------------------


// *******************************************************************
// *       Part 1: definition of the system variables                *
// *******************************************************************


// dumb utility
[external!(m:module) : string
 -> if known?(external,m) m.external else string!(m.name) ]

[string2module(s:string) : module
  -> let m := get_value(s) in
       case m (module m,
               any error("~A is not a module",s)) ]

// *******************************************************************
// *       Part 2: definition of the main function
// *******************************************************************

// help file
[printHelp() : void
 -> printf("------------- CLAIRE: The Art of Elegant Programming -----------\n\n"),
    about(),
    printf("\noptions -s <int> <int> : set memory allocation size  \n"),
    printf("        -f <filename>  : load <filename>             \n"),
    printf("        -env <sys> : compile for a different OS target \n"),
    printf("        -n : do not load the init file               \n"),
    printf("        -m <module> : load <module>                  \n"),
    printf("        -v <int> : upgrade the verbosity level       \n"),
    printf("        -S <flag> : sets the global variable <flag> to true  \n"),
    printf("        -o <name> : sets the name of the executable  \n"),
    printf("        -ld <name> : sets the library directory  \n"),
    printf("        -od <name> : sets the output directory  \n"),
    printf("        -p : profiling mode                          \n"),
    printf("        -D : debug mode                              \n"),
    printf("        -safe : safe mode                            \n"),
    printf("        -O : optimizing mode                         \n"),
    printf("        -os <int> : sets the optimizer safety level          \n"),
    printf("        -l <lib> : adds <lib> to the list of needed libs     \n"),
    printf("        -cm <module>: compiles a module -> executable        \n"),
    printf("        -cc <module>: compiles a module -> target files      \n"),
    printf("        -cl <module>: compiles a module -> library           \n"),
    printf("        -cx <main file> : generates an executable from a file\n"),
    exit(0) ]


//Claire's main
[main(lp:list[string]) : void
  -> let rCode := true, %cm := "", %cf := "", dblevel := 1, %out := "", %cj := "", slevel := 0,
         clevel := 1, %init? := true, vlevel := 2, l := (copy(lp) as list<string>) in
 (try
  (*fs* := Id(*fs*),
   PRODUCER.Generate/extension := Id(PRODUCER.Generate/extension),
   compiler.libraries_dir := Id(compiler.libraries_dir),
   compiler.headers_dir := Id(compiler.headers_dir),
   compiler.options := Id(compiler.options),
   compiler.env := Id(compiler.env),
   claire_lib := "",                                                  // erase
   while (l)
   (case l[1]
    ({"?", "-help"} printHelp(),
     {"-s"}  (if (length(l) >= 3)  l :<< 3 else error("option: -s <s1> <s2>")),
     {"-f"}  (if (length(l) >= 2)  (load(l[2]), l :<< 2)
              else error("option: -f <filename>")),
     {"-env"}  (if (length(l) >= 2)  (compiler.env := l[2], l :<< 2)
                else error("option: -env <OS name>")),
     {"-m"}  (if (length(l) >= 2)
               (if %init? (load("init"), %init? := false),
                let m := string2module(l[2]) in
                  (load(m), begin(m), l :<< 2 , claire_modules :add m))
              else error("option: -m <module>")),
     {"-v"} (if (length(l) >= 2) (vlevel :+ integer!(l[2]), l :<< 2)
             else error("option: -v <integer>")),
     {"-ld"} (if (length(l) >= 2) (claire_lib := l[2], l :<< 2)
             else error("option: -od <directory>")),
     {"-od"} (if (length(l) >= 2) (compiler.source := l[2], l :<< 2)
              else error("option: -od <directory>")),
     {"-os"} (if (length(l) >= 2) (slevel := integer!(l[2]), l :<< 2)
              else error("option: -ol <int>")),
     {"-S"} (if (length(l) >= 2) 
                (value(new(global_variable,symbol!(l[2]))) := true,
                 l :<< 2)
             else error("option: -S <FLAG>")),
     {"-o"} (if (length(l) >= 2) (%out := l[2], l :<< 2) 
             else error("option: -o <name>")),
     {"-p"} (OPT.Compile/profile? := true, dblevel :max 1, l :<< 1),
     {"-D"} (dblevel := 0, l :<< 1),
     {"-safe"} (safety(compiler) := (if (dblevel = 0) 0 else 1),
                claire_lib := compiler.libraries_dir[2],
                claire_options := compiler.options[2],
                l :<< 1),
     {"-O"} (compiler.optimize? := true, dblevel := 2, l :<< 1),
     {"-l"} (if (length(l) >= 2)  (compiler.libraries :add l[2], l :<< 2)
             else error("option: -l <library>")),
     {"-cl"} (if (length(l) >= 2) (%cm := l[2], l :<< 2)
              else error("option: -cl <module>")),
     {"-cc"} (if (length(l) >= 2)  (clevel := 0, %cm := l[2], l :<< 2)
               else error("option: -cc <module>")),
     {"-cm"}  (if (length(l) >= 2)   (clevel := 2, %cm := l[2], l :<< 2)
               else error("option: -cm <module>")),
     {"-cj"}  (if (length(l) >= 2) (%cj := l[2], l :<< 2)),
     {"-cjx"}  (if (length(l) >= 2) (%cj := l[2], clevel := 0, l :<< 2)),
     {"-cx"} (if (length(l) >= 2)   (%cf := l[2], l :<< 2, clevel := 2)
               else error("option: -cx <filename>")),
     {"-n"} (%init? := false, l :<< 1),
     any (if (l[1][1] = '-') (printf("~S is an unvalid option\n",l[1]),
                              printHelp()),
          rCode := false, l := list<string>() ) )),
   if (%out = "") (if (%cm != "") %out := %cm else if (%cf != "") %out := %cf),
   if %init? load("init"), 
   claire_options :=  compiler.options[(if (dblevel = 0) 2 else if (dblevel = 2) 1 else 3)],
   if (claire_lib = "")
       claire_lib := compiler.libraries_dir[(if (dblevel = 0) 2 else if (dblevel = 2) 1 else 3)],
   system.verbose := vlevel,
   if (slevel > 0) compiler.safety := slevel,        // v3.3.26
   if (%cm != "")
      let m := string2module(%cm) in
        (compiler.active? := true,
         if (m.uses = list(claire_modules[2]))
            (claire_modules := shrink(claire_modules,2),
             trace(0,"=== Light Module ~S:~S -> use ~S=== ", m, m.uses,claire_modules)),
         claire_modules :add m,
         load(get_value("Compile")),
	 if (%out != "") external(m) := %out,
	 load(m),
         if (dblevel < 1) (compiler.safety :min 4, compiler.debug? :add m),
         compile(m),
         if (clevel = 1)
            (if (%out != "") m.external := %out, cmakefile(m,""))
         else if (clevel = 2) cmakefile(m,%out),                       // v3.2.12: level = 0 => do nothing ....
         exit(0))
    else if (%cj != "")
       call(get_value("jcmakefile"), string2module(%cj), %out, (clevel = 0))
    else if (%cf != "")
       (compiler.active? := true,
        load(get_value("Compile")),
        load(%cf),
        function_compile(%cf,%cf),
	cmakefile(%cf,%out),
        exit(0))
   )                                                                ; while (l)
   catch any (mClaire/restore_state(reader),
              Reader/debug_if_possible())) ]



 
// *******************************************************************
// *       Part 3: single file compiling                             *
// *******************************************************************

// compile a single command file: can only generate functions
// since there is no associated module
// we assume that the file only contains function definitions
//
[function_compile(self:string, fullname:string) : void
 -> Compile/need_modules(OPT) := set(),
    let %interface:string := compiler.headers_dir / self /+ ".h" in
      (OPT.Compile/legal_modules := set!(module),
       OPT.Compile/properties := set<property>(),
       OPT.Compile/objects := set<object>(),
       OPT.Compile/functions := list(),
       OPT.Compile/cinterface := fopen(%interface,"w"),
       OPT.Compile/cfile := self,
       Generate/generate_file(fullname, source(compiler) / self),
       use_as_output(OPT.Compile/cinterface),
       Generate/breakline(),
       for x in Compile/need_modules(OPT),
          (if (made_of(x) & not(x % claire_modules))
              printf("#include <~I.h>\n",Generate/ident(name(x)))),
       // generate the function definitions
       for l in OPT.Compile/functions     // prints the API methods
        let %f := l[1], %vars := l[2], s := l[3] in
          printf("\nextern ~I ~I(~I);",
                Generate/interface!(PRODUCER,s), c_princ(%f),
                Generate/typed_args_list(%vars)),
       fclose(OPT.Compile/cinterface)) ]


// *******************************************************************
// *       Part 4: Generating makefiles                              *
// *******************************************************************


// prints the list of lib files that are needed for m
[lib!(m:any,l:list) : void             // + Tibor
  -> let s_end := ".lib", %env := env(compiler), first := true,
         s_sep := (case %env ({"unix","win32v","ntv"} " $Z" /+ *fs*,  any ",")) in
      (case %env
	 ({"ntw"} printf("LIBP $Z L "),
          {"unix"} printf("$Z/"),
	  {"win32v","ntv"} printf("$Z\\")),
       for m2 in {x in l  | made_of(x)}
	      printf("~I~A~A~I",
                     (if first (first := false) else princ(s_sep)),
                     external!(m2),s_end,
	             for s in {s1 in uses(m2) | s1 % string}
		    	      printf("~A~A~A",s_sep,s,s_end)),
       for m2 in compiler.libraries printf("~A~A~A",s_sep,m2,s_end)) ]

// prints the necessary files for the two compilation modes
// if link? is true we need a comma-separated list
// if m is a module, we compile fi*.cpp + m.cpp -> m.lib
// if m=f is a file, we compile f-s.cpp (system), f.cpp (functions) -> exe
[files!(m:any,link?:boolean,%out:string) : void
 ->  let %bef := (if link? "" else "$T" /+ *fs*),
         %end := (if (env(compiler) = "unix") "o" else "obj"),
         %sep := (if link? "," else " ") in
        printf("~A~A.~A~I",%bef,                              // added by Tibor
	        (if (link? | m % string) (%out /+ "-s") else name(m)),
                %end,
	        (case m
	          (module  (if not(link?)
                              for ff:string in made_of(m)
	                             printf("~A~A~A.~A ",%sep,%bef, ff, %end)),
	           string printf("~A~A~A.~A",%sep, %bef, m, %end))))]
		 

// module linker - dispatch according to hardware and OS
[cmakefile(m:any, out:string) : void
  -> let l :=  Reader/add_modules(claire_modules),
         f := (case m (module string!(name(m)), any out)), // name of the .mk file
         %os := env(compiler) in
       (if (out != "") Generate/generate_s_file(out,l,m), // v3.2.54 !!!!
        if (%os = "win32v") compiler.libraries :add "gui"
        else if (length(claire_modules) = 3) compiler.libraries :add "noConsole"
        else compiler.libraries :add "Console",
        if (%os % {"ntw", "ntv","win32v"}) create_makefile_nt(m,out,l)    // Visual
        else if (%os % {"unix","osx"}) create_makefile_unix(m, out, l)
        else error("Unknown environment, should be one of :'ntv','ntw','win32v','unix','osx'\n"),
        print_in_string(),
        if (%os = "unix") printf("make -f ~A.mk",f)
        else printf("nmake /c /f ~A.mk",f),
        shell(end_of_string())) ]

    
    
;----------------------- NT makefile (MS Visual C++ or Watcom) ----------------------

// creates the nt makefile for a module or a string
// m is a module or nothing
// out is a string or nothing
// l is a list of library modules
// ... with the help of Arnaud Linz
[create_makefile_nt(m:any,out:string,l:list)  : void
 -> let f := (case m (module string!(name(m)), any out)), // name of the .mk file
        p := fopen(f /+ ".mk", "w"), %env := compiler.env,
        %I := (if (%env = "ntw") "/i=" else "/I"), %O := (if (%env = "ntw") "/fo=" else "/Fo"),
        sis := (if (%env = "ntw") "nt" else %env) in
     (use_as_output(p),
      printf("option = ~A~A ~A /DCLPC /DCLWIN\n",%I, compiler.headers_dir,claire_options),
      printf(".SUFFIXES : .exe .obj .cpp\n\n"),
      printf("Z = ~A\nT = ~A\n",claire_lib,source(compiler)),
      printf("CC = ~A\n", (if (compiler.env = "ntw") "wpp386" else "cl")),
      printf("FILES = ~I\n",files!(m,false,out)),
      printf("{$T}.cpp{$T}.obj:\n"),
      printf("	$(CC) $(option) ~A$T\\$(<B).obj ~I$<\n", %O,
               (if (%env != "ntw") princ("/Tp ") )),
      printf("all: ~I\n",
              (if (out != "") printf("~A.exe",out) else printf("$Z\\~S.lib",m))),
      case m
       (module
          (printf("$Z\\~A.lib: $(FILES)\n",external!(m)),
           if (%env = "ntw")
              printf("\t!wlib /q /c / b $Z\\~A.lib +-$?\n",external!(m))
           else printf("\tlib /NOLOGO /OUT:$@ $(FILES)\n"))),
        if (out != "")
        (if (%env != "ntw") printf("JUNK = /NOLOGO /DEBUG /MAP /STACK:1600000 user32.lib gdi32.lib shell32.lib comdlg32.lib\n"),
         printf("~A.exe: ~I\n",out,
                 (case m (module printf("$Z\\~A.lib $T\\~A-s.obj",external!(m),out),
                          any princ("$(FILES)")))),
         if (%env = "ntw")
               printf("\twlink sys ~A N ~A d all option ~A P $T F ~I ~I\n\n",
                      sis, out,"q,d,ST=600K,c,mang",
                      files!(m,true,out),                   // includes main,
                      lib!(m,l))                            // includes all that is needed
         else printf("\tlink /subsystem:~A $(JUNK) ~I /OUT:~A.exe ~I\n\n",
                     (if (%env = "win32v") "windows" else "console"),
                      lib!(m,l), out,
                      (if (m % module) printf("$T\\~A-s.obj",out)
                       else princ("$(FILES)")))),
       fclose(p)) ]



; ------------------------  Unix (Linux, Solaris, ...) -----------------------
  
// creates the unix makefile for a module or a string (cf. nt makefile)
// this was copied from v2.4.28
// with the help of Francois Laburthe !
[create_makefile_unix(m:any,out:string,l:list)  : void
 -> let f := (case m (module string!(name(m)), any out)), // name of the .mk file
        p := fopen(f /+ ".mk", "w") in
     (//[5] =========== CREATE UNIX MAKEFILE ~A.mk (in progress) ======= // f,
      use_as_output(p),
      printf("# --- unix makefile --- version ~A ----\n",release()),
      printf("Z = ~A\n",claire_lib),              // Z: where to find the libs
      printf("T = ~A\n\n",source(compiler)),      // T: where to find the .cc
      printf("FILES = ~I\n",files!(m,false,out)), // list of c++ files
      printf("CXX = g++\n"),                      //  C++ compiler
      printf("LINK = ld -r\n"),                      // linker
      printf("CXXFLAGS = -I~A -DCLUNIX ~A\n\n",   // C++ compiler options
             compiler.headers_dir, claire_options),
      // dependency line for producing .o from .cc s
      princ("$T%.o:	$T%.cc\n	$(CXX) $(CXXFLAGS) -c $(@:.o=.cc) -o $@\n"),
      // global target
      printf("all: ~I\n",
              (if (out != "") printf("~A",out) else printf("$Z~A~S.lib",*fs*,m))),
      // compile the module m into m.lib
      case m
       (module
          (printf("$Z~A~A.lib: $(FILES)\n",*fs*,external!(m)),
           printf("\t$(LINK) -o $Z~A~A.lib $(FILES)\n",*fs*,external!(m)))),
      // generate an executable from m (a module) or a file (string)
      if (out != "")
        (printf("~A: ~I\n",out,
           (case m (module printf("$Z~A~A.lib $T~A~A-s.o",
                                  *fs*,external!(m),*fs*,out),
                          any princ("$(FILES)")))),
         printf("\t$(CXX) -o ~A ~I ~I\n\n",
                out,
                lib!(m,l),
                (if (m % module) printf("$T~A~A-s.o",*fs*,out)
                 else princ("$(FILES)")))),
      fclose(p)) ]

// ---------------------------------------------------------------




