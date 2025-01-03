//+-------------------------------------------------------------+
//| CLAIRE                                                      |
//| jssystem.cl                                                 |
//| Copyright (C) 2023-2024 Yves Caseau. All Rights Reserved    |
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
    varsym:integer = 0,            // disambiguate variables by adding a number
    methods:list,                  // stack of tuples(method, lambda, name)
    traces:list)                   // list of properties (for debugging)

claire/jtrace :: property(open = 3)

// made generic with print_true
[print_true(p:go_producer) : void
   -> princ("true")]

//*********************************************************************
//*          Part 2: Module Compiler Interface                        *
//*********************************************************************

// a small test function for the compiler
[claire/j_test(x:any) : void 
 -> j_test(claire,x)]

// debug handle
claire/LASTU:any :: unknown

[claire/j_test(m:module,x:any) : void
  ->  let t := c_type(x),         // type inference (reused)
         s := Compile/osort(t),           
         u := Compile/c_code(x, s),       // claire optimizer (reused)
         f := j_func(u) in  
       (PRODUCER := claire/JS_PRODUCER,
        PRODUCER.current := m,
        LASTU := u,
        printf("type -> ~S [sort ~S]\n", t, s),
        printf("opt[~S] -> ~S \n", owner(u), u),
        if f printf("exp  -> ~I\n", j_expression(u))
        else printf("stat -> ~I\n", j_statement(u,"result",nil))) ]

// test the compiling of a method
// e.f. g_test(foo @ any)
[j_test(m:method) : void
  -> when l := get(formula,m) in
        (//[0] ---- JS compiling ~S with following definition ---- // m,
        PRODUCER := claire/JS_PRODUCER,
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
        make_js_function(PRODUCER,formula(m),m),
        OPT.in_method := unknown ) ]


// compile the modules and check that no necessary modules is not
// declared
[claire/jcompile(m:module) 
   -> PRODUCER := claire/JS_PRODUCER,
      PRODUCER.debug? := false,
      compile(PRODUCER,m)]              //  shortcut that already exists

// reuses the v4.12 generic compiler
[compile(p:js_producer, m:module) : void
 ->  compile@code_producer(p,m),
     trace(1, "~S: ~A lines of code compiled. ~A warnings, ~A notes.\n",
              m, compiler.n_loc, compiler.n_warnings, compiler.n_notes) ]


// start the produced go file
// Puts the reference to the packages, and some useful comments
// we limit the use of "unsafe" Go package to the module file
// in javascript we produce only one file, so the last parameter is not used
[start_file(p:js_producer,s:string,m:module,unused?:boolean) : any
 ->  //[0] ===== start file for ~A ==== // s,
     use_as_output(OPT.outfile),
     printf("/***** CLAIRE Compilation of module ~A into Javascript \n",s),
     printf("       [version ~A / safety ~S] ~A *****/\n\n",release(),compiler.safety, date!(0)),
     printf("const kernel = require('./ClaireKernel')\n\n"),
     use_as_output(stdout) ]


// For each class we produce two things in the module-generated-file
//   - the struct (with embedded inheritance)
//   - we also gerenate a constructor  C() or C(name)
[gen_classes(p:js_producer,m:module) : void
 ->  //[0] ===== generate classes for ~S ==== // m,
     for c in {c in OPT.objects | c % class}
       (OPT.level := 0,
        printf("\n// class file for ~S in module ~S ",c,m), // v0.02
        gen_class_def(p,c))] 


 // how to generate a struct associated to a class
[gen_class_def(p:js_producer,c:class) : void
 ->  let ls := cdr(c.slots),n := length(ls) in
        (//[0] ------------- generate class ~S (~A) -------------------- // c,length(c.comment),
         if (length(c.comment) > 20) printf("// ~A",c.comment)
         else breakline(),
         printf("class ~I extends ~I~I ~I",js_class(c),js_class(c.superclass),new_block("class def"),breakline()),
         printf("constructor(~I) ~I", 
                  (if (c <= thing) princ("name")),
                   new_block()),
         if (c <= thing)
             (ls := cdr(ls),
              printf("super(name)~I",breakline()))
         else printf("super()~I",breakline()),
         for s in list{s2 in ls | known?(default,s2) } 
             printf("this.~I = ~I~I",ident(PRODUCER,s.selector.name),
                        j_expression(s.default),breakline()),
          close_block("constructor"),
          breakline(),
          gen_class_methods(p,c),
          close_block("class def"),
          breakline())]

   //       printf("var Meta~I = new MetaClass(~A,Meta~I)~I",js_class(c),
   //                string!(c.name),js_class(c.superclass),breakline()))]

// print and indent a string that starts with \n
[princIndent(s:string) : void
  -> let n := length(s), m := get(s,char!(10)) in 
       (if (m > 0) 
          (princ(slice(s,1,m - 1)),
           breakline(),
           princIndent(slice(s,m + 1,n)))
        else (princ(s),indent_c()))]


// generate all the methods associated to c (look in the stack)
// we use the comment slot to attach the file's comment lines to the method
[gen_class_methods(p:js_producer,c:class) : void
  ->  let prev_comment := "" in 
        for x in p.methods
          (case x (string prev_comment := x,
                   list let m := x[1] as method in
                       (if (jsMethod?(m) & m.domain[1] = c) 
                           (printf("// ----- class method ~S ------------- ",m),
                            if (prev_comment != "") princIndent(prev_comment)
                            else breakline(),
                            make_js_function(p,x[2],m))
                        else prev_comment := ""))) ]


// generate all the other methods as functions
[gen_functions(p:js_producer,m:module) : void
  ->  let prev_comment := "" in 
        for x in p.methods
          (case x (string prev_comment := x,
                   list let m := x[1] as method in
                       (if not(jsMethod?(m)) 
                           (printf("// ----- function from method ~S ------------- ",m),
                            if (prev_comment != "") princIndent(prev_comment) 
                            else breakline(),
                            make_js_function(p,x[2],m))
                        else prev_comment := ""))) ]
  

// generate the definition of the named objects from the module (used in both modes)
// must move to the producer
[gen_objects(p:js_producer,m:module) : void
 -> //[3] ===== generate objects for ~S [graph : ~S] ==== // m, Core/graph % OPT.properties,
     for x in OPT.objects     
        printf("var ~I~I", go_var(x.name),breakline()),
     breakline()]
  
// generate the meta_load function 
//    + contains all the floating statements in the code)
//    - does not contain any meta (module, properties, etc.)
[gen_meta_load(p:js_producer,m:module) : void   // v0.02
  -> //[3] ===== generate meta_load function for ~S ==== // m,
     printf("\n//--------------- meta description + top-level instructions ----\n"),
     printf("function MetaLoad() ~I~I",new_block(), breakline()),
     printf("// instructions from module sources~I",breakline()),
     for i in OPT.instructions 
         (// printf("//(debug) -> ~S~I",owner(i),breakline()),
          if ignore_exp?(i) 
             let l := extract_global_variables(i) in
               (if (length(l) = 2)
                  printf("~I = ~I ~I",js_ident(l[1]),j_expression(l[2]),breakline())
                else trace(0,"[debug] ignored ~S -> ~S\n",i,l))
                // if not j_func(l[2]) this code will break
          else if (i % string)     // comment
              printf("// ~A", i)
          else if j_func(i)
              printf("~I~I",j_expression(i),breakline())
          else  call_j_statement(i, "niet", nil)),
       printf("console.log(\"------------- end of ~S meta_load --------------\")~I",m,breakline()),
       close_block("MetaLoad"),
       printf("MetaLoad()"),
       breakline(),                       // v3.0.3
       let p := get_value("jsmain") in 
         (if (p % property & ((p @ void) % method))
            printf("jsmain()~I",breakline()) ) ]



// these are the expressions that we want to ignore
[ignore_exp?(self:any) : boolean
  -> case self
      (Call_method  let p := self.arg.selector in 
              (p = add_slot | p = add_method),
       Call  let p := self.selector in
                ((p = object! & self.args[2] = global_variable) | p = claire/jtrace),
       C_cast ignore_exp?(self.arg),
       Let ignore_exp?(self.value),
       any false)]

// extract the global variables from an expression
// returns nil or a list (name:symbol, value:any)
[extract_global_variables(self:any) : list
  -> case self
      (Let let x := (self.value as C_cast).arg, y := self.arg in
             (//[5] extract looks at x:~S and y:~S // x,y,
              case x (Call  let p := x.selector in
                              (if (p = object! & x.args[2] = global_variable)
                                   list(x.args[1], (((y as Do).args[2]) as Update).value)
                               else nil),
                      any nil)),
       any nil) ]


// *********************************************************************
// *     Part 3: File compilation                                      *
// *********************************************************************

// gen_file is now generic in v4.12, we just need 
// (1) to declare outfile?
// (2) to implement gen_insruction
[outfile?(p:js_producer) : boolean
  -> false]
[modfile_name(p:js_producer,m:module) : string
  -> string!(m.name) ]

// the only specific part is the instruction (because of comment management)
// NOTE: this is actually cool and could be reused in Go compiling
[gen_instruction(p:js_producer,%instruction:any,prev_comment:string) : string  
   -> if (%instruction % string)            // we have found a comment => put on methods stack
         (trace(5,"READ COMMENT [~A]:~A\n", %instruction, length(%instruction)),
          prev_comment :/+ ("\n// " /+ %instruction))
      else 
         (if (%instruction % Defclass) trace(0,"READ DEFCLASS ~S [comment :~A]\n",%instruction,length(prev_comment)),
          if (%instruction % Defmethod) p.methods : add prev_comment    // store the comment associated to m
          else if (%instruction % Defclass) 
             value(iClaire/ident(%instruction as Defclass)).comment := prev_comment,
          prev_comment := "",
          OPT.instructions :add c_code(%instruction, void)),   // instructions stack => deferred module generation
        prev_comment ]

//**********************************************************************
//*     Part 4: the lambda-to-function compiler                        *
//**********************************************************************

// This is the function that is called by the optimizer
// For javascript compiling, we store the method+lambda in a stack since everything is placed in the m.js file
[make_c_function(self:lambda,%nom:string,m:any) : void
 ->  if (m % method) PRODUCER.methods :add list(m,self)
     else error("lambdas such as ~S are not supported with Diet Claire",self) ]

 // JS function declaration
[generate_function_start(p:js_producer,self:lambda,m:method) : void
 ->  let lv := (if (length(self.vars) = 1 & m.domain[1] % {void, environment})  nil       // v3.0.05 was : & not(s = float)
                else self.vars) in
      (if jsMethod?(m)
          printf("~I (~I) ",goMethod(m as method), jsVariables(p,cdr(self.vars)))
       else printf("function ~I (~I) ", 
                     c_princ(jsFunction(p,m)), jsVariables(p,lv)))   ]

// This method creates a js function or method from a claire lambda for a method m.
// 
[make_js_function(p:js_producer,self:lambda,m:method) : void
  -> let typeOK := check_range(m,self.body),
         s := class!(m.range),
         %body := c_strict_code(self.body,s),
         throw? := g_throw(%body) in
      (//[5] >>>>>>>>>>>> JS_FUNCTION applied to lambda ~S >>>>>> (~S) // self,m,
       compiler.n_methods :+ 1,
       p.varsym := 0,                                            // resets the ID for distinct vars
       if jsMethod?(m) self.vars[1].mClaire/pname := symbol!("this"),
       generate_function_start(PRODUCER, self, m),        
       new_block(),
       if (m.selector % p.traces) 
          printf("console.log(\"[trace] start ~S\")~I",m,breakline()),
       if not(j_simple_body?(%body))  procedure_js_body(m,self,%body,s)
       else function_js_body(%body,s),
       close_block(),
       breakline()) ]

// check that we may call function_body  (replaces the print_body method of CLAIRE 3 compiler)  
// simple : we can generate ... return X directly without the need for a "Result" variable 
[j_simple_body?(self:any) : boolean
   -> case self
       (If j_func(self.test) & j_simple_body?(self.arg) & j_simple_body?(self.other),
        Do j_simple_body?(last(self.args)),
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
            b_expression(PRODUCER,self.test, true),
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
             else call_j_statement(x, "niet",nil)))
]

// default complex case : create a variable "Result"
[procedure_js_body(m:method, %l:lambda, %body:any,s:class) : void
  ->  if (s = void) j_statement(%body,"niet",nil)
      else (printf("var Result ~I",breakline()),
            call_j_statement(%body,"Result",nil),
            printf("return Result~I",breakline())) ]


// end of file
