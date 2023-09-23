//+-------------------------------------------------------------+
//| CLAIRE                                                      |
//| jssystem.cl                                                 |
//| Copyright (C) 2023-2023 Yves Caseau. All Rights Reserved    |
//| cf. copyright info in file object.cl: about()               |
//+-------------------------------------------------------------+

//**********************************************************************
//* Contents                                                           *
//*          Part 1: Global_variables & producer interface             *
//*          Part 2: Module Compiler Interface                         *
//*          Part 3: File Compiler                                     *
//*          Part 4: Function Compiler                                 *
//**********************************************************************


//**********************************************************************
//*          Part 1: Global_variables & JavaScript Producer            *
//**********************************************************************

// ----------------------- inline coding --------------------------------
// we reuse the global variables that are present in the go compiler

// add the go_producer here  (replaces the C++ producer)
// note that the double list bad/good names is ugly and should be replaced by a dictionary later 
js_producer <: code_producer(
    current:module,                        // module that is being compiled
    bad_names:list[symbol],        // avoid generating !
    good_names:list[symbol],       // replacements (same order)
    kernel_methods:list,               // dictionary for go "sugar" (nice methods in Kernel versus functions)
    source:string,                 // where to place the go code
    debug?:boolean = false,        // if debug, add /* explanation */ into code
    varsym:integer = 0)            // desambiguate variables by adding a number


// we can reuse the indent / *_block methods from Go !

//*********************************************************************
//*          Part 2: Module Compiler Interface                        *
//*********************************************************************

// a small test function for the compiler
[claire/j_test(x:any) : void 
 -> j_test(claire,x)]

[claire/j_test(m:module,x:any) : void
  ->  let t := c_type(x),         // type inference (reused)
         s := osort(t),           
         u := c_code(x, s),       // claire optimizer (reused)
         f := j_func(u) in  
       (PRODUCER := JS_PRODUCER
        PRODUCER.current := m,
        printf("type -> ~S [sort ~S]\n", t, s),
        printf("opt[~S] -> ~S \n", owner(u), u),
        if f printf("exp  -> ~I\n", j_expression(u)
        else printf("stat -> ~I\n", j_statement(u,"result")) ]

// test the compiling of a method
// e.f. g_test(foo @ any)
[j_test(m:method) : void
  -> when l := get(formula,m) in
        (//[0] ---- JS compiling ~S with following definition ---- // m,
        PRODUCER := JS_PRODUCER
        pretty_print(body(l)),
        OPT.in_method := m,
        OPT.Optimize/use_string_update := false,   // v3.3.46
        OPT.Optimize/max_vars := 0,
        OPT.legal_modules := set!(module.instances),
        OPT.outfile := stdout,
        compiler.inline? := true, 
        PRODUCER.current := claire,
        trace(0,"\n---- code produced by the optimizer -------------------\n"),
        pretty_print(c_strict_code(formula(m).body,class!(m.range))),
        trace(0,"\n---- code produced by the generator ------------------- \n"),
        make_go_function(PRODUCER,formula(m),"test",m),
        OPT.in_method := unknown ) ]


// compile the modules and check that no necessary modules is not
// declared
[claire/jcompile(m:module) 
   -> PRODUCER := JS_PRODUCER
      compile(PRODUCER,m)]              //  shortcut that already exists

// note: this could be the same code as the go compiler ...  [note: should be generic]
// only difference is the simpler stats (Diet => no dynamic calls)
[compile(p:js_producer, m:module) : void
 ->   OPT.need_modules := {},
      compiler.inline? := true, 
      compiler.n_loc := 0,                // number of lines of code
      compiler.n_warnings := 0,           // number of warnings
      compiler.n_notes := 0,              // number of notes
      let l1:bag := parents(Reader/add_modules(list(m))) in
      (//[3] ==========  START JS COMPILING (~S) with ~S ================ // m, l1,
       OPT.legal_modules := set!(l1),
       p.current := m,                                              // v4: we need to know in which module we are
       p.source := compiler.source / string!(m.name),               // produce code in the <src>/<module> directory
       js_files(p,m),                // files to files
       js_mod_file(p,m),
       l1 := difference(set!(OPT.need_modules), OPT.legal_modules),
       if l1 (warn(),trace(1, "~S should be declared for ~S \n", l1, m)),
       trace(1, "~S: ~A lines of code compiled. ~A warnings, ~A notes.\n",
              m, compiler.n_loc, compiler.n_warnings, compiler.n_notes)) ]

// the first part is to generate the js files in the FileToFile mode [note: should be generic]
[js_files(p:js_producer,m:module) : void
 ->  //[0] ==== Generate ~A files for module ~S [verbose = ~A, Opt? = ~S] // PRODUCER.comment, m, verbose(),compiler.optimize?,
     OPT.instructions := list<any>(),
     OPT.properties := set<property>(),
     OPT.objects := list<object>(),
     OPT.functions := list<any>(),
     OPT.need_to_close := set<any>(),
     begin(m),
     for x in m.made_of
       (//[1] ++++ Compiling the file ~A.cl [v. 4.~A - safety:~A] // x, compiler.version, compiler.safety,
        if (x = string!(m.name))
           Cerror("[211]  ~S cannot be used both as a file and module name",x),
        OPT.level := 1, // debug - to remove
        p.current_file := x,       // CLAIRE4 : keep the file name handy 
        js_file(p, m.source / x, p.source / x)),
     end(m) ]


// Creates the "meta" file for the module m. [note: should be generic]
[js_mod_file(p:go_producer, m:module) : void
 -> let prt := fopen(((PRODUCER.source /+ *fs*) /+ string!(m.name) /+ "-meta") /+ PRODUCER.Generate/extension, "w"),
        s := string!(m.name) in
      (//[2] ==== generate file for module ~S ==== // m,
       OPT.outfile := prt,
       start_file(p,s,m,true),            // true tells this is the module file
       use_as_output(OPT.outfile),
       gen_classes(p,m),                 // v4.0 : keep the class/struct definition in the module
       gen_objects(p,m),
       gen_meta_load(p,m),                 // reflective description (class, methods, vars)
       breakline(),
       close_block(),
       breakline(),                       // v3.0.3
       if (compiler.safety > 4) //[1] ===== [CROSS]  ~A BAD METHODS : ~S  // length(BadMethods), BadMethods,
       fclose(OPT.outfile)) ]

// start the produced go file
// Puts the reference to the packages, and some useful comments
// we limit the use of "unsafe" Go package to the module file
[start_file(p:js_producer,s:string,m:module, module?:boolean) : any
 ->  use_as_output(OPT.outfile),
     printf("/***** CLAIRE Compilation of ~A ~A.cl \n         [version ~A / safety ~S] ~A *****/\n\n",
            (if module? "module" else "file"), s, 
            release(),compiler.safety,
            date!(0)),
     use_as_output(stdout) ]


// For each class we produce two things in the module-generated-file
//   - the struct (with embedded inheritance)
//   - we also gerenate a constructor  C() or C(name)
[gen_classes(p:js_producer,m:module) : void
 ->  //[3] ===== generate classes for ~S ==== // m,
     for c in {c in OPT.objects | c % class}
       (OPT.level := 0,
        printf("\n// class file for ~S in module ~S ",c,m), // v0.02
        breakline(), 
        gen_class_def(p,c)] 


 // how to generate a struct associated to a class
 [gen_class_def(p:go_producer,c:class) : void
 ->  let ls := cdr(c.slots),n := length(ls) in
        (printf("Class ~I ~I~I",go_class(c),new_block(),breakline())
         printf("constructor(~I) ~I", 
                  (if (c <= thing) princ("name")),
                   new_block()),
          printf("this.isa = Meta~I~I",go_class(c),breakline()),
          if (c <= thing) printf("this.name = name~I",breakline())
          for s in list{s2 in ls | known?(default,s2)} 
             printf("this.~I = ~I~I",ident(PRODUCER,s.selector.name),
                        j_expression(s.default),breakline()),
          if (c.open = 3)
             printf("Meta~I.instances.push(this)~I",go_class(c),breakline()),
          close_block("constructor"),
          close_block(),
          breakline(),
          printf("var Meta~I = new MetaClass(~A,Meta~I)~I",go_class(c),go_class(c),breakline()))]


// generate the definition of the named objects from the module (used in both modes)
// must move to the producer
[gen_objects(p:jss_producer,m:module) : void
 -> //[3] ===== generate objects for ~S [graph : ~S] ==== // m, Core/graph % OPT.properties,
     for x in OPT.objects     
        printf("var ~I~I", go_var(x.name),breakline()),
     breakline()]
  
// generate the meta_load function 
//    + contains all the floating statements in the code)
//    - does not contain any meta (module, properties, etc.)
[gen_meta_load(p:go_producer,m:module) : void   // v0.02
  -> //[3] ===== generate meta_load function for ~S ==== // m,
     printf("func MetaLoad() ~I~I",new_block(), breakline()),
     printf("// instructions from module sources",breakline()),
     for i in OPT.instructions 
         (breakline(),
          if (i % string)     // comment
              printf("~I// ~A", (if not(j % string) breakline()), i)
          else if j_func?(i)
              printf("~I~I",j_expression(i),breakline())
          else  statement(i, "Unused")) ]          // Unused : "not needed" marker

// *********************************************************************
// *     Part 3: File compilation                                      *
// *********************************************************************

// this is the basic file cross_compiler, which translates from claire to javascript
// [note: should be generic] : this method should be attacted to code_producer
[js_file(p:js_producer, f1:string,f2:string) : void
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
            (if (%instruction % string)            // we have found a comment
               let pp := use_as_output(OPT.outfile) in
                 (printf("\n//~A", %instruction),
                  use_as_output(pp))
             else OPT.instructions :add c_code(%instruction, void),
             %instruction := Reader/readblock(p1)),
       compiler.n_loc :+ n_line(),
       fclose(p1),
       compiler.loading? := false,
       // restore reading context
       reader.Reader/toplevel := b,
       reader.external := "toplevel",
       reader.Reader/fromp := p0,
       fclose(OPT.outfile)) ]


//**********************************************************************
//*     Part 4: the lambda-to-function compiler                        *
//**********************************************************************

// This is simplified in CLAIRE4 since the class2file mode is no longer supported
// we could re-introduce it from CLAIRE 3.5 if we want to support Java compiling
[make_c_function(self:lambda,%nom:string,m:any) : void
 ->  if (m % method) make_js_function(PRODUCER,self,%nom,m)
     else error("lambdas such as ~S are not supported with Diet Claire",self) ]

 // JS function declaration
[generate_function_start(p:go_producer,self:lambda,s:class,m:any, %nom:string) : void
 ->  let %dom := (if self.vars self.vars[1].range else any),
         %f := make_function(%nom),
         lv := (if (length(self.vars) = 1 & %dom % {void, environment})  nil       // v3.0.05 was : & not(s = float)
                else self.vars) in
      (OPT.functions :add list(%f, lv, s),        // register the function in the API list
       printf("\n/* The js function for: ~I */\n",     // for debug: add {~A}, OPT.Compile/level,
               (case m
                 (method printf("~S(~I)", m.selector, Language/ppvariable(self.vars)),
                  any princ(string!(%f))))),
       printf("function ~I (~I) ", jsFunction(p,m), jsVariables(p,lv)))   ]

  
// This method creates a js function from a claire lambda for a method m.
// 
[make_js_function(p:js_producer,self:lambda,%nom:string,m:method) : void
  -> let typeOK := check_range(m,self.body),
         s := class!(m.range),
         %body := c_strict_code(self.body,s),
         throw? := g_throw(%body) in
      (compiler.n_methods :+ 1,
       p.varsym := 0,                                            // resets the ID for distinct vars
       use_as_output(OPT.outfile),
       generate_function_start(PRODUCER, self, s, m, %nom),        
       new_block(),
      if not(simple_body?(%body)  procedure_js_body(m,self,%body,s)
      else function_js_body(%body,s)))
        close_block(),
       use_as_output(stdout)) ]

// check that we may call function_body  (replaces the print_body method of CLAIRE 3 compiler)  
// simple : we can generate ... return X directly without the need for a "Result" variable 
[simple_body?(self:any) : boolean
   -> case self
       (If j_func(self.test) & simple_body?(self.arg) & simple_body?(self.other),
        Do simple_body?(last(self.args)),
        any j_func(self))]


// simpler case that we apply for Do, Ifs and functional expressions
// however is c_type(exp) is void we need to return CNULL
[function_js_body(self:any,s:class) : void
  -> let %ret := (if (s != void) "return " else "") in
      (if (s = boolean & (case self (Call_method (let p := self.arg.selector in 
                                      (p = = | p = < | p = > | p = >= | p = <=)  ))))       
                               // this is an old optimization - there is a debate if this is still needed with CLAIRE4
                               // reintroduced in v4.0.7 for mSend, but only for direct comparisons
          printf("if ~I {return true~I} else {return false}",bool_exp(self,true),breakline())
       else if (c_type(self) = void & s != void)
         printf("~I~Ireturn ~I~I",
                 j_expression(self), breakline(),
                 j_expression(unknown), breakline())
      else printf("~A ~I~I", %ret, j_expression(self),breakline())) ]

// generate nice code for If function (inspired from g_statement@If)
[function_js_body(self:If,s:class) : void
  -> printf("if ~I ~I",
            js_bool_exp(self.test, true),
            new_block("body If")),
    function_js_body(self.arg,s),
    if (self.other = nil) close_block()
    else if (self.other % If) 
      printf("~I else ~I",finish_block(), function_js_body(self.other,s))  
    else if (s != void | not(designated?(self.other)))
        printf("} else {~I~I~I", breakline(),
                     function_js_body(self.other,s),
                     close_block("body If"))
    else close_block("body If") ]

// generate nice code for a Do
[function_js_body(self:Do, s:class) : void
  ->  let l := self.args, %length := length(l), m := 0 in
        ( for x in l
            (m :+ 1,
             if (m = %length) function_js_body(x,s)
             else j_statement(x, "Unused")))
]

// default complex case : create a variable "Result"
[procedure_body(m:method, %l:lambda, %body:any,s:class) : void
  ->  if (s = void) j_statement(%body,"Unused")
      else (printf("var Result ~I",breakline()),
            j_statement(%body,"Result",)
            printf("return Result~I",breakline())) ]


// end of file
