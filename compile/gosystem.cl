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
*?_interval* :: (set! @ Interval) 
*--_integer* :: (.. @ integer)
*+_integer* :: (+ @ integer)
*nth_integer* :: (nth @ integer)
*nth_list* :: (nth @ list)
*nth_tuple* :: (nth @ tuple)      // CLAIRE4.0 duplicate (list -> tuple)
*nth_1_list* :: (nth_get @ list)
*nth_1_tuple* :: (nth_get @ tuple)
*nth_1_array* :: (nth_get @ array)
*nth_string* :: (nth @ string)
*nth_1_string* :: (nth_get @ string)
*nth=_list* :: (nth= @ list)
*nth_put_list* :: (nth_put @ list)
*make_list* :: (make_list @ list(integer,type,any))
*not* :: (not @ any)
*known* :: (known? @ any)
*unknown* :: (unknown? @ any)
*not_equal* :: (!= @ any)
*equal* :: (= @ any)
*belong* :: (Core/belong @ any)         // special method in types => dispatch to any
*contain* :: (contain? @ list)
// Compile/*min_integer* :: (min @ integer)
// Compile/*max_integer* :: (max @ integer)
*length_array* :: (length @ array)
*length_bag* :: (length @ bag)
*close_exception* :: (close @ exception)            // v3.2.58  */
*new_class1* :: (mClaire/new! @ list(class))
*new_class2* :: (mClaire/new! @ list(class,symbol))
*slot_get* :: (slot_get @ object)
*map* :: (map! @ type)
// bag methods could be ommited in the future - these methods are used to force compiling (could be removed later)
*of_bag* :: (of @ bag)
*of_array* :: (of @ array)
*copy_list* :: (copy @ list)
*copy_set* :: (copy @ set)
*empty_set* :: (empty @ set)          // force goMethod()
*nth_put_array* :: (nth_put @ array)
*%t* ::  (Core/%t @ any)
*lesst* :: (Core/<=t @ type)
*included* :: (<= @ type)
*stack_apply* :: (stack_apply @ property)
*super_apply* :: (Core/super_apply @ property)
*princ_string* :: (princ @ string)
*inherit* :: (inherit? @ class)
*write_value* :: (Language/write_value @ Variable)
*read_property* :: (read @ property)


// new: the target code production (the part that depends on the target language) is
// encapsulated with a producer object
// CLAIRE 4 is focused on go, but we try to keep the previous structure of CLAIRE3 to be ready
// for Java or Swift compiling. However, the GC management stuff is lost forever :)
code_producer <: producer(
   open_comparators:list[operation],      // list of comparison ops that are inlined (the order matters!!)
   open_operators:list[operation],        // list of arithmetic operators that are inlined
   body:any = 0,                          // used to store the body of the current method
   extension:string,                      // extension for generated files
   comment:string,                        // a string that designates the target language
   interfaces:list,                       // used to translate imported to C/.. entities
   stat:integer = 0)                      // v3.3.32: stats about GC protection  */

// add the go_producer here  (replaces the C++ producer)
// note that the double list bad/good names is ugly and should be replaced by a dictionary later 
go_producer <: code_producer(
    current:module,                        // module that is being compiled
    bad_names:list[symbol],        // avoid generating !
    good_names:list[symbol],       // replacements (same order)
    kernel_methods:list,               // dictionary for go "sugar" (nice methods in Kernel versus functions)
    source:string)                 // where to place the go code


// TODO: define a status = 3 for the PRODUCER class that tells that is it extensible
// (Genearate/producer.open := 3)
// this is a special case : the function may return an error but the optimized form does not
EIDSET:set<any> := set<any>(nth @ list)

// most standard method: call the producer to print the ident from a symbol
ident(self:symbol) : void -> ident(PRODUCER,self)
ident(self:thing) : void -> ident(PRODUCER,self.name)
ident(self:class) : void -> ident(PRODUCER,self.name)

// we simply use some smart identation. True pretty_printing will be left to bc
[indent_c() : any
 -> let x := OPT.level in while (x > 0) (princ("  "), x :- 1) ]

breakline() : any -> (princ("\n"), indent_c())

// adds a new C block with the condensed option
[new_block() : void
 -> OPT.level :+ 1, princ("{ "), breakline()]

// closes the current C block
[close_block() : void 
 -> OPT.level :- 1, princ("} "), breakline() ]

 // prints the } without a new line - used for nested If
[finish_block() : void 
 -> OPT.level :- 1, princ("} ")]


//*********************************************************************
//*          Part 2: Module Compiler Interface                        *
//*********************************************************************

// a small test function for the compiler
[claire/g_test(x:any) : void 
 -> g_test(claire,x)]

[claire/g_test(m:module,x:any) : void
  ->  let t := c_type(x),
         s := osort(t),
         u := c_code(x, s),
         f := g_func(u),
         gt := g_throw(u) in  
       (PRODUCER.current := m,
        printf("type -> ~S [sort ~S]\n", t, s),
        printf("opt[~S] -> ~S \n", owner(u), u),
        if gt printf( "----------------------- Error is possible => EID (func:~S)  ----------------\n",f),
        if f printf("exp  -> ~I\n", g_expression(u, class!(t)))
        else printf("stat -> ~I\n", statement(u, (if gt EID else class!(t)), "result", void))) ]

// even more fun 
[gtop() : void 
   -> princ("in> "),
      let x := read(stdin) in
        (if (x = q) princ("bye.\n")
         else (g_test(x), gtop())) ]


// test the compiling of a method
// e.f. g_test(foo @ any)
[g_test(m:method) : void
  -> when l := get(formula,m) in
        (//[0] ---- Compiling ~S with following definition ---- // m,
        pretty_print(body(l)),
        OPT.in_method := m,
        OPT.Optimize/use_string_update := false,   // v3.3.46
        OPT.Optimize/max_vars := 0,
        OPT.legal_modules := set!(module.instances),
        OPT.outfile := stdout,
        compiler.inline? := true, 
        PRODUCER.current := claire,
        //[0] ---- code produced by the optimizer ------------------- //,
        make_go_function(PRODUCER,formula(m),"test",m),
        OPT.in_method := unknown ) ]

// debug (to remove later)
claire/BadMethods:list<method> :: unknown

// compile the modules and check that no necessary modules is not
// declared
[claire/compile(m:module) -> compile(PRODUCER,m)]   //  shortcut that already exists
[compile(p:go_producer, m:module) : void
 ->   //[0] --------- start compile ------------------- //,
      OPT.need_modules := {},
      BadMethods := list<method>(),
      compiler.inline? := true, 
      let l1:bag := parents(Reader/add_modules(list(m))) in
      (//[0] ==========  START GO COMPILING (~S) with ~S ================ // m, l1,
       OPT.legal_modules := set!(l1),
       p.current := m,                                 // v4: we need to know in which module we are
       p.source := compiler.source / string!(m.name),
       gen_files(p,m),                // files to files
       gen_mod_file(p,m),
       l1 := difference(set!(OPT.need_modules), OPT.legal_modules),
       if l1 trace(2, "---- WARNING: ~S should be declared for ~S \n", l1, m)) ]

// the first part is to generate the go files in the FileToFile mode
[gen_files(p:go_producer,m:module) : void
 ->  //[0] ==== Generate ~A files for module ~S [verbose = ~A, Opt? = ~S] // PRODUCER.comment, m, verbose(),compiler.optimize?,
     OPT.instructions := list<any>(),
     OPT.properties := set<property>(),
     OPT.objects := list<object>(),
     OPT.functions := list<any>(),
     OPT.need_to_close := set<any>(),
     begin(m),
     for x in m.made_of
       (//[2] ++++ Compiling the file ~A.cl [v. 3.~A - safety:~A] // x, compiler.version, compiler.safety,
        if (x = string!(m.name))
           Cerror("[211]  ~S cannot be used both as a file and module name",x),
        OPT.level := 1, // debug - to remove
        p.current_file := x,       // CLAIRE4 : keep the file name handy 
        gen_file(p, m.source / x, p.source / x)),
     end(m) ]


// This makes the initial loading function by compilinf all the claire
// expression placed in the list oself. *new_objects* holds all the new
// objects defined in this file.
// The name of the function is built from the file name (s argument)
//
[gen_mod_file(p:go_producer, m:module) : void
 -> let prt := fopen(((PRODUCER.source /+ *fs*) /+ string!(m.name)) /+ PRODUCER.Generate/extension, "w"),
        s := string!(m.name) in
      (//[0] ==== generate file for module ~S ==== // m,
       OPT.outfile := prt,
       verbose() := 3,
       start_file(p,s,m,true),            // true tells this is the module file
       use_as_output(OPT.outfile),
       gen_classes(p,m),                 // v4.0 : keep the class/struct definition in the module
       gen_objects(p,m),
       gen_meta_load(p,m),                 // reflective description (class, methods, vars)
       breakline(),
       close_block(),
       breakline(),                       // v3.0.3
       //[0] ===== ~A BAD METHODS : ~S // length(BadMethods), BadMethods,
       fclose(OPT.outfile)) ]

// start the produced go file
// Puts the reference to the packages, and some useful comments
// we limit the use of "unsafe" Go package to the module file
[start_file(p:go_producer,s:string,m:module, module?:boolean) : any
 ->  use_as_output(OPT.outfile),
     printf("/***** CLAIRE Compilation of ~A ~A.cl \n         [version ~A / safety ~S] ~A *****/\n\n",
            (if module? "module" else "file"), s, 
            release(),compiler.safety,
            substring(date!(1),1,24)),
     namespace!(p,m),
     printf("import (_ \"fmt\"\n"),
     if module? printf("\t\"unsafe\"\n"),            // import go.unsafe (support pointer cast)
     import_declaration(m),
     printf(")\n"),                                  // end of import 
     dumb_import(m),
     use_as_output(stdout) ]

// import declarations
[import_declaration(m:module) : void
 ->  for x in needed_modules(m)
       printf("\t~I\"~A\"\n", 
          (if (x = Kernel) princ(". ")),  // Kernel is always included in namespace
          string!(x.name)) ]
     

// go requires an import list without redundancy + we only import
[needed_modules(m:module) : list
  -> let l := list{ m2 in (Reader/add_modules(list(m)) but m) |
                   (m2.made_of | m2 = Kernel) } in
       // clean_duplicates(l)
       l ]

// create a dumb function that prevents the go compiler to complain
[dumb_import(m:module)
  -> let l := needed_modules(m) in
      (if (length(l) > 1)
        (printf("\n//-------- dumb function to prevent import errors --------\n"),
         printf("func import_~A() ~I",gensym(),new_block()),
         for x in l
            (if (x != Kernel) printf("_ = ~I.It~I",cap_short(x.name),breakline())),
               // let y := representative(x) in
                 // printf("_ = ~I~I",g_expression(y,owner(y)),breakline())),
         close_block(),
         breakline())) ]

// pick a thing in module m
[representative(m:module) : any
   -> some(x in (class U property) | 
        defined(x.name) = m & 
         (case x (property exists(y in x.restrictions | module!(y) = m), any true))) ]

// remove dual imports (hopefully, works if the import path is simple enough)
[clean_duplicates(l:list) : list
  -> let l2 := copy(l), n := length(l), i := n - 1 in
       (while (i > 1)       // 0 : remove Kernel 1: Keep it
          (if exists(j in (i + 1 .. n) | l[i] % Reader/add_modules(list(l[j]))) 
              l2 :delete l[i],
            i :- 1),
        l2) ]

// For each class we produce two things in the module-generated-file
//   - the struct (with embedded inheritance)
//   - the cast method
//   - we also gerenate a constructor  makeC(a1, ... , an) when there are no inverses 
[gen_classes(p:go_producer,m:module) : void
 ->  //[3] ===== generate classes for ~S ==== // m,
     for c in {c in OPT.objects | c % class}
       (OPT.level := 0,
        printf("\n// class file for ~S in module ~S ",c,m), // v0.02
        breakline(), 
        gen_class_def(p,c),
        gen_cast_function(p,c),
        if construct_class?(c) gen_construct(p,c))] 

 // how to generate a struct associated to a class
[gen_class_def(p:go_producer,c:class) : void
 -> printf("type ~I struct ~I ~I~I ~I~I",
               go_class(c),
               new_block(),
               go_class(c.superclass),
               breakline(),
               for y:slot in list{s in get_indexed(c) | domain!(s) = c}
                  (printf("~I ~I", cap_short(y.selector.name),interface!(class!(y.range))),
                   breakline()),
               close_block())]

 // how to produce the ToC() cast function that applies to any pointer (using unsafe)
 [gen_cast_function(p:go_producer,c:class) : void
  -> printf("\n// automatic cast function\n"),
     printf("func ~I(x *ClaireAny) *~I {return (*~I)(unsafe.Pointer(x))}",
               cast_class(c),
               go_class(c),
               go_class(c)),
     breakline()]

// when we want a constructor ? when slots are simple (no inverse, no store ...)
// TODO : to complete with the proper test
[construct_class?(c:class) : boolean 
  -> c <= object & length(c.slots) <= 5 ]


// generate a constructor
[gen_construct(p:go_producer,c:class) : void
  -> let first := true in 
      (printf("\n// automatic constructor function\n"),
       printf("func Make~I~I(~I) *~I ~I",
            addUnderscore(c),
            go_class(c),
            for y:slot in cdr(get_indexed(c))
               (if first (first := false) else princ(","),
                printf("~I ~I", ident(y.selector.name), interface!(class!(y.range)))),
            go_class(c),
            new_block("make")),
      printf("var o *~I = new(~I)~I",go_class(c),go_class(c),breakline()),
      printf("o.Isa = ~I~I",class_ident(c),breakline()),
      for y:slot in cdr(get_indexed(c))
         printf("o.~I = ~I~I",cap_short(y.selector.name),ident(y.selector.name),breakline()),
      printf("return o ~I~I",breakline(),close_block("make"))) ]


// generate the definition of the named objects from the module (used in both modes)
// must move to the producer
[gen_objects(p:go_producer,m:module) : void
 -> //[3] ===== generate objects for ~S [graph : ~S] ==== // m, Core/graph % OPT.properties,
     for x in OPT.objects     
      (breakline(),
       case x
        (global_variable printf("var ~I ~I", go_var(x.name),
                                 interface!((if nativeVar?(x) getRange(x) else global_variable))),
         any (printf("var ~I ~I /*obj*/", go_var(x.name),
                    interface!(Compile/psort(owner(x))))))),
    for x in {p in OPT.properties | not(p % OPT.objects) }  
       (when p2 := some(p2 in (OPT.properties but x) |            // v0.01
                        string!(p2.name) = string!(x.name)) in
         error("[217] ~S and ~S cannot be defined in the same module",p2,x),
        printf("~Ivar ~I ~I // ~S",breakline(),
                thing_ident(x),
                interface!(Compile/psort(owner(x))),
                x.name)),
    breakline(),
    printf("var It *ClaireModule~I",breakline()),
    // we will add the definition of all ancestors which are not package in a first-born line
    let m1 := m, m2 := m.part_of in
       (while (m2 != claire & m2.parts[1] = m1)      // similar to the gen_module loop
          (if (m2.made_of = nil) printf("var ~I *ClaireModule ",go_var(m2.name)),
           m1 := m2,
           m2 := m2.part_of)),
    breakline()]

// generate the meta_load function
// in go the load function for M is M_load()
[gen_meta_load(p:go_producer,m:module) : void   // v0.02
  -> //[3] ===== generate meta_load function for ~S ==== // m,
     printf("// definition of the meta-model for module ~S ~I",m, breakline()),
     printf("func MetaLoad() ~I~I",new_block(), breakline()),
     gen_module(p,m,m),
     printf("// definition of the properties ~I",breakline()),
     for x in {p in OPT.properties | not(p % OPT.objects) & p != value & p != vars}      // vars & value moved to kernel
        printf("~I~I = ~I", breakline(), thing_ident(x), declare(p,x)),
     breakline(), 
     breakline(),
     printf("// instructions from module sources",breakline()),
     let j:any := unknown in 
       for i in OPT.instructions 
         (breakline(),
          if (i % string)  printf("~I// ~A", (if not(j % string) breakline()), i)
          else if g_throw(i)                   // need to protect from an error in compiled code
             (new_block(),
              printf("/*PROTECT ~S */~I",i,breakline()),
              var_declaration("expr",EID,1),
              g_statement(i,EID,"expr",true,false),
              printf("ErrorCheck(expr)"),
              close_block())
          else  statement(i, void, "Niet", false),          // Niet is a junk EID variable
          j := i) ]

// generate the module definition - only the module structure (the decoration is found in the system file)
// cool recursive method that ensures that all non-package modules are visible
// load_m() has an implicit begin(m) so that new methods are assigned to m
[gen_module(p:go_producer,m:module,%package:module) : void   // v0.02
  -> // if (m.part_of != claire & m.part_of.parts[1] = m)    // we are part of a composite ...
     //   gen_module(p,m.part_of,%package),                 // add the definition of the composite
     printf("~I = MakeModule(~S,~I)~I",
            (if (m = %package) princ("It") else printf("~I",go_var(m.name))),
            string!(m.name),
            g_expression(m.part_of,module),              // now we can use m.part_of (gen_module reccursion)
            breakline()),
     printf("ClEnv.Module_I = It~I",breakline())]        // implicit begin(m)

// reciprocate : finds the concrete module where a package module must be defined.
[get_made(self:module) : module
  -> let m := self.parts[1] in
      (if (m = Kernel | m.made_of) m else get_made(m)) ]

// called by gosystem.cl : declare a property or an operation (handles the dispatch case)
[declare(c:go_producer,p:property) : void
 -> printf("Make~A(~S,~A,~I~I)",
           (if (p % operation) "Operation" 
            else "Property"),
           string!(p.name),
           p.open, 
           g_expression(module!(p.name),module),
           (if (p % operation) printf(",~A",p.precedence))) ]

// This is a similar method which places all the necessary modules
// in the right order so that self can be defined
parents(self:module,l:list) : list
 -> (if (self % l) l
     else (if known?(part_of, self) l := parents(self.part_of, l),
           l :add self, l))

// this methods takes a list of modules that must be loaded and returns
// a list of modules that are necessary for the definition
//
parents(self:list) : list
 -> (let l := list<module>() in (for x in self l := parents(x, l), l))

// useful (v3.0.06)
[claire/get(m:module) : void 
  ->  load(m), begin(m) ]

// *********************************************************************
// *     Part 3: File compilation                                      *
// *********************************************************************

// this is the basic file cross_compiler, which translates from claire to go
// this file compiler runs only in the good environment (the file to be compiled must be already loaded).
// it generates methods definitions in f2 and stores the instructions into OPT.instructions
[gen_file(p:go_producer, f1:string,f2:string) : void
 -> let p1 := fopen(f1 /+ ".cl", "r"), b := reader.Reader/toplevel,
        p0 := reader.Reader/fromp in         // b, p0: reading context when we start
       (OPT.outfile := fopen(f2 /+ p.Generate/extension, "w"),
        reader.Reader/toplevel := false,
        compiler.loading? := true,
        n_line() := 1,
        reader.external := f1,
        reader.Reader/fromp := p1,                     // <yc> ensures automatic fclose !
        start_file(p,f1,module!(),false),               // CLAIRE 4 : always add a header
        let %instruction := Reader/readblock(p1) in
          while not(%instruction = Reader/eof)
            (if (system.verbose > -1 & not(%instruction % string))
             printf("[~S/~A:~A] ~S (~S)\n", module!(), fileName(f1), n_line(), owner(%instruction), OPT.need_modules),  
             case %instruction (Defobj 
                let %s := %instruction.iClaire/ident in
                   printf("[defobj ident is ~S->~S]\n",%s,defined(%s))),
             if (%instruction % string)            // we have found a comment
              (printf("// ~A\n", substring(%instruction,1,length(%instruction) - 1)),
               if (compiler.Optimize/naming < 2)   // which is put back in the files
                let pp := use_as_output(OPT.outfile) in
                 (printf("\n//~A", %instruction),
                  use_as_output(pp)))
             else OPT.instructions :add c_code(%instruction, void),
             %instruction := Reader/readblock(p1)),
       fclose(p1),
       compiler.loading? := false,
       // restore reading context
       reader.Reader/toplevel := b,
       reader.external := "toplevel",
       reader.Reader/fromp := p0,
       fclose(OPT.outfile)) ]

 // sugar
 [fileName(s:string) : string
   -> let n := length(s), i := get(s,*fs*[1]) in
        (if (i > 0) fileName(substring(s,i + 1, n)) else s) ]   

//**********************************************************************
//*     Part 4: the lambda-to-function compiler                        *
//**********************************************************************


// This is simplified in CLAIRE4 since the class2file mode is no longer supported
// we could re-introduce it from CLAIRE 3.5 if we want to support Java compiling
[make_c_function(self:lambda,%nom:string,m:any) : void
 ->  if (m % method) make_go_function(PRODUCER,self,%nom,m)
     else make_lambda_function(PRODUCER,self,%nom) ]


// In CLAIRE 4 we separate methods from free lambdas (used for demons, but which could be used to compile lambda blocks)
// this is used for demons as well as second-order-types

// create an EID lambda  
[make_lambda_function(p:go_producer,self:lambda,%nom:string) : void
  -> //[0] ===== generate an EID function from a lambda for ~A // %nom,
     use_as_output(OPT.outfile),
     generate_function_start(PRODUCER, self, EID, nil, %nom),        // defined in gogen
     new_block(),
     printf("/* eid body: ~S */~I",self.body,breakline()),
     eid_body(self.body,true,EID),
     close_block(),
     breakline(),
     generate_eid_dual(self,%nom),                          // generate E_C(nom)
     use_as_output(stdout) ]

 // how to declare a function in the interface file and its header in the
// output file
[generate_function_start(p:go_producer,self:lambda,s:class,m:any, %nom:string) : void
 ->  let %dom := (if self.vars self.vars[1].range else any),
         %f := make_function(%nom),
         lv := (if (length(self.vars) = 1 & %dom % {void, environment})  nil       // v3.0.05 was : & not(s = float)
                else self.vars) in
      (OPT.functions :add list(%f, lv, s),        // register the function in the API list
       if (compiler.naming != 2)           // TODO: shoud we keep the naming option - 2 (generate code that is hard to reverse engineer)
       printf("\n/* {~A} OPT.The go function for: ~I */\n", OPT.Compile/level,
               (case m
                 (method printf("~S(~I) [~I]", m.selector,
                                Language/ppvariable(self.vars),
                               (if (compiler.naming = 1) bitvectorSum(status!(m)))),
                  any princ(string!(%f))))),
       if goMethod?(m)
          printf("func (~I) ~I (~I) ~I ",goVariable(p,self.vars[1]), goMethod(m as method), 
                  goVariables(p,cdr(self.vars)),
                  (if (s != void) interface!(s)))
       else if (m = nil) printf("func F_~I (~I) EID ", c_princ(%nom), goVariables(p,lv))
       else printf("func ~I (~I) ~I ", goFunction(m), goVariables(p,lv), 
                  (if (s != void) interface!(s))))   ]

  

// This method creates a go function from a claire lambda for a method m.
// %name is the name that was proposed for the lambda (or derived from method m)
// we either use function_body to try a simple approach or (procedure_body | eid_body) that add all the trimmings
[make_go_function(p:go_producer,self:lambda,%nom:string,m:method) : void
  -> let typeOK := check_range(m,self.body),
         s := class!(m.range),
         %body := c_strict_code(self.body,s),
         throw? := g_throw(%body) in
      (//[0] ---- ~S: make_go(~S) => simple=~S // m, %body, simple_body?(%body),
       if (m.status != -1 & throw? != can_throw?(m))        // avoids generating go code that will break
          (//[0] ======================== WARNING ======================================== //,
           //[0] >>>>> ~S body produces an error (g_throw = true) while status is 0 <<<<<<< // m,
           if (m.status = 0) BadMethods :add m
           // error("==== cross-compiling error with ~S: the new error status means that it must be in ~S",
           //      m, (if throw? "ForceThrow" else "ForceNotThrow"))
          ),
       use_as_output(OPT.outfile),
       if ((typeOK | compiler.safety > 3) & not(throw?) & (m.selector != self_eval))  // happy with the type inference => native function
          (//[5] --- Procedure generation (can throw = ~S) // throw?,
           generate_function_start(PRODUCER, self, s, m, %nom),        
           new_block(),
           if (need_debug?(m) |  OPT.profile? | not(simple_body?(%body)) | s = void) 
              procedure_body(m,self,%body,s)
           else (princ("// use function body compiling \n"),
                 function_body(%body,s)))
       else (//[0] --- EID function generation (can throw = ~S) // throw?,
             throw? := true,                                       // this is the EID pathd
             generate_function_start(PRODUCER, self, EID, m, %nom),        
             new_block(),
             // printf("/*G_throw = ~S for ~S, s:~S*/~I",throw?,%body,s,breakline()),
             eid_body(%body,typeOK,s)),
       close_block(),
       generate_eid_function(self,m,throw?),
       if (m.selector = self_eval) generate_eval_function(self,m),
       use_as_output(stdout)) ]

// check that we may call function_body  (replaces the print_body method of CLAIRE 3 compiler)  
// simple : we can generate ... return X directly without the need for a "Result" variable 
[simple_body?(self:any) : boolean
   -> case self
       (If g_func(self.test) & simple_body?(self.arg) & simple_body?(self.other),
        Do simple_body?(last(self.args)),
        any g_func(self))]

// generic case (g_func is true)

// simpler case that we apply for Do, Ifs and functional expressions
[function_body(self:any,s:class) : void
  -> let %ret := (if (s != void) "return " else "") in
      printf("~A ~I~I", %ret, g_expression(self,s),breakline()) ]

// generate nice code for If function (inspired from g_statement@If)
[function_body(self:If, s:class) : void
  -> printf("if ~I ~I",
            bool_exp(self.test, true),
            new_block("body If")),
    function_body(self.arg,s),
    if (self.other = nil) close_block()
    else if (self.other % If) 
      printf("~I else ~I",finish_block(), function_body(self.other,s))  
    else if (s != void | not(designated?(self.other)))
        printf("} else {~I~I~I", breakline(),
                     function_body(self.other,s),
                     close_block("body If"))
    else close_block("body If") ]

// generate nice code for a Do
[function_body(self:Do, s:class) : void
  ->  let l := self.args, %length := length(l), m := 0 in
        ( for x in l
            (m :+ 1,
             if (m = %length) function_body(x,s)
             else statement(x, void, "Unused", void)))
  ]

// default complex case : create a variable "Result"
[procedure_body(m:method, %l:lambda, %body:any,s:class) : void
  ->  if OPT.profile? generate_profile(PRODUCER,m),
      if need_debug?(m) debug_intro(PRODUCER,m),
      printf("// procedure body with s = ~s \n",s),
      if (s != void) 
         (var_declaration("Result",s,1),
          statement(%body,s,"Result",false))
      else statement(%body,void,"Unused",false),
      return_result(PRODUCER,s,m,"Result") ]

// generate an EID function 
// call for the debug/profile is needed     
[eid_body(%body:any,typeOK:boolean, s:class) : void
 -> // printf("/* eid_body s = ~S */\n",s),
    var_declaration("Result",EID,1),
    statement(%body,EID,"Result",g_throw(%body)),
    if typeOK printf("return Result")
    else printf("return RangeCheck(~I,Result)",g_expression(s,type))]

// generate the EID function associated to each method (used by the interpreter - EID mode)
[generate_eid_function(self:lambda,m:method,throw?:boolean) : void
  -> let %sig := go_signature(m), lv := self.vars in
      (printf("\n// The EID go function for: ~S (throw: ~S) \n", m, throw?),
       printf("func ~I (~I) EID ", goEIDFunctionName(m), goEIDVariables(PRODUCER,lv)),
       new_block(),
       if (m.range = void & not(throw?))
          printf("~I~Ireturn EVOID", print_EID_call(m,lv,%sig,throw?), breakline())
       else printf("return ~I", print_EID_call(m,lv,%sig,throw?)),  
       close_block()) ]


 // similar but simpler for a lambda associated to a name (e.g. 2nd order types) => E_C(nom)      
[generate_eid_dual(self:lambda,%nom:string) : void
  -> let lv := self.vars, nl? := length(lv) > 3  in
      (printf("\n// The dual EID go function for: ~S \n", %nom),
       printf("func E_~I (~I) EID ", c_princ(%nom), goEIDVariables(PRODUCER,lv)),
       new_block(),
       printf("return F_~I(", c_princ(%nom)),  
       for n in (1 .. length(lv))
            external_EID_arg(lv[n],class!(range(lv[n])),n,nl?),
       princ(")"),
       close_block()) ]

// EID function calls the compiled native function - uses a code that looks like print_external_call
// watch out: a method that can throw returns an EID directly ! (same as goexp.cl : print_ext_call)
[print_EID_call(m:method,l:list,%sig:list<class>,throw?:boolean) : void
 -> let  n := 1,  sm := last(%sig), nl? := length(l) > 3 in
     (if nl? OPT.level :+ 1,
      if (throw? | m % EIDSET | m.selector = self_eval) 
          sm := EID,     // the function returns an EID !
      cast_prefix(sm,EID),
      printf("/*(sm for ~S= ~S)*/ ",m,sm),
      if goMethod?(m)
        (external_EID_arg(l[1],%sig[1],1,nl?),
         printf(".~I(",goMethod(m)),
         for n in (2 .. length(l))
            external_EID_arg(l[n],%sig[n],n - 1,nl?))
      else 
        (printf("~I(", goFunction(m)),
         if (length(l) = 1 & domain!(m) = void)  l := nil,
         for n in (1 .. length(l))
            external_EID_arg(l[n],%sig[n],n,nl?)),
      princ(" )"),
      if nl? OPT.level :- 1,
      cast_post(sm,EID)) ]

 // here v is a EID-range variable and we need to extract the native s representation
 // n=0 is a special marker when the arg the receiver x in x.f(....)
 [external_EID_arg(v:Variable,s:class,n:integer,nl?:boolean) : void
   -> if (n > 1) (princ(","), if nl? breakline()),
      // printf("/* ext_arg s=~S/*",s),
      eid_prefix(s),
      ident(PRODUCER,v),
      eid_post(s) ]
             
// prints a list of arguments with types / replaces typed_args_list
[goEIDVariables(p:go_producer,self:list) : any
 -> let prems := true in
       for v:Variable in self
         (if prems prems := false else printf(","),
          printf("~I EID", ident(p,v))) ]

// check the range & sort of the method through type inference. 
// returns true if OK and false otherwise (can produce an error at run-time)
// notice that %body is the lambda body before compilation => use c_type
[check_range(self:method,%body:any) : any
 -> let s1 := class!(self.range),                // declared
        s2 := class!(c_type(%body)) in           // inferred
      (//[0] ---- info: ~S's range was found to be ~S (vs. ~S) // self, s2, s1,
       if (s1 = void | s2 <= s1) true       // method is statically type-safe  
       else 
         (//[0] ---- note: ~S's range was found to be ~S (vs. ~S) // self, s2, s1,
          if ((s1 != void & s2 = void & s1 != error) | (s1 ^ s2) = {})
                  Cerror("[218] Sort error: Cannot compile ~S (~S cannot be ~S).",self,s1,s2)
          else false)) ]

// generate the eval function associated to each self_eval method (type *any -> EID)
// EVAL_C(x *ClaireAny) EID {return ToC(x).SelfEval()}
[generate_eval_function(self:lambda,m:method) : void
  -> let c := domain!(m) in
      (if (c != Variable)                                       // Variable is managed in Kernel (clEnv.go)
      (printf("\n// The EVAL go function for: ~S \n", c),
       printf("func EVAL_~I (x *ClaireAny) EID ", c_princ(c.name)),
       new_block(),
       if goMethod?(m) printf(" return ~I(x).SelfEval()", cast_class(c))
       else printf(" return F_self_eval_~I(~I(x))", c_princ(c.name),cast_class(c)),
       close_block()))  ]

// tells if a method needs debug instrumentation
[need_debug?(m:any) : boolean
 -> case m
      (method let p := m.selector in
                (m.module! % compiler.debug? &
                 domain!(m) != environment &
                 m.module! != claire & p != self_eval & p != execute &
                 p != eval_message & p != Core/push_debug & p != Core/pop_debug &
                 p != Core/tr_indent & p != Core/find_which & p != Core/matching? &
                 p != Core/vmatch?),
       any false) ]

// profiler code 
[generate_profile(c:go_producer,m:any) : void
 -> if (m % method) get_dependents(m),
    printf("   var PR_x *ClairePRcount PRget_property(~I).Start();~I",
             expression( (case m (method m.selector, any fastcall)), nil),
             breakline())  ]


// produce the debugging code introduction
// in go we have no macros but functions with variable number of args
// assumes DebugBind(module, method, ClaireAny args* ) 
[debug_intro(c:go_producer,self:lambda,x:method) : void
 -> let m := (case x (method x.module!)),
        n := 1 in
       printf("DebugBind(~I,~I~I);~I", ident(m),
              expression(x.selector, {}), 
              (if (length(self.vars) = 1 & (self.vars[1]).range = void)       // foo() means foo(system) ?
                  printf(",EID{C_object,ClEnv.Uip()}));")
              else (for v in self.vars
                      (printf(", ~I", to_eid(c, v, go_signature(x)[n])),
                       n :+ 1))),
               breakline()) ]
                  

// auxiliary to produce the end statement for the function. s tells if the result is needed.
// generates a "... return" if the result is needed or just an empy string
// we also add the debugging unbind if needed.  (used to be called protect_result)
[return_result(p:go_producer,s:class,x:method,%res:string) : void
 ->  if need_debug?(x)
           printf("DebugUnbind(~I,~I,~I)~I", ident(x), g_expression(x.selector, property),
                     to_eid(p, (if (s = void) unknown else build_Variable("Result", s)), object),
                     breakline())
      else if OPT.profile? printf("PRend(PR_x)~I",breakline()),
      if (s != void) printf("return ~A", %res) ]
          

// computes the inter-module dependence
get_dependents(m:method) : void
  -> (for p in Reader/dependents(m) Reader/PRdependent[m.selector] :add p)

// prints a function name without the # syntactic marker for imported
[c_princ(self:function) : void  -> import_princ(string!(self)) ]
[import_princ(s:string) : void 
   -> for i in (1 .. length(s))
        (if (i > 1 | s[i] != '#') c_princ(s[i])) ]


// v3.2.06 - some properties may be extended
//(put(open,Generate/set_outfile,4),
// put(open,Generate/inline_exp,4))

// end of file
