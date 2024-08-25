// *******************************************************************************************
// *     log file : started on December 31st, 2023                                           *
// *******************************************************************************************

// THIS LOG FILE sits in the claire/dropbox/compile directory : it tells the stoty of the 
// CLAIRE-to-Javascript compiler

// *******************************************************************
// * Contents                                                        *
// *     Part 1: Structure / README                                  *
// *     Part 2: Memento about CLAIRE to JavaScript                  *
// *     Part 3: Log                                                 *
// *******************************************************************


// *******************************************************************
// *     Part 1: Structure / README                                  *
// *******************************************************************

// (1) README (to extend)
files
- jssystem.cl
- jsexp.cl
- jsstat.cl
- jsgen.cl

Note: this javascript code generator is much simpler for two reasons
- no types in javascript + Javascript supports exception handling
- Diet compiling : assumes that the code is statically typed from CLAIRE
    -> no Calls
    -> no mention of types
    -> no complex/dynamic type expressions
    -> no rules

Key methods
- j_test( expression or module or method)
- PRODUCER = JS_OBJECT => contains the key slots for code generation
    PRODUCER.source = "node" by default

// (2)  -------------------------------- FILE STRUCTURE ----------------------------------------

//+-------------------------------------------------------------+
//| jssystem.cl                                                 |
//+-------------------------------------------------------------+
//*          Part 1: Global_variables & producer interface             *
//*          Part 2: Module Compiler Interface                         *
//*          Part 3: File Compiler                                     *
//*          Part 4: Function Compiler                                 *

//+-------------------------------------------------------------+
//| gogen.cl                                                    |
//+-------------------------------------------------------------+
// *     Part 1: definition of the code producer                     *
// *     Part 2: utilities for file generation                       *
// *     Part 3: interface declarations                              *
// *     Part 4: use of language dependent patterns (macros)         *
// *     Part 5: Utilities                                           *

//+-------------------------------------------------------------+
//| jsexp.cl                                                    |
//+-------------------------------------------------------------+
// *  Part 1: g_func & expression for objects                          *
// *  Part 2: expression for messages                                  *
// *  Part 3: the inline coding of function calls                      *
// *  Part 4: expression for structures                                *
// *  Part 5: boolean optimization                                     *

//+-------------------------------------------------------------+
//| jsstat.cl                                                   |
//+-------------------------------------------------------------+
//*          Part 1: Unfolding of complex expressions                  *
//*          Part 2: Basic control structures                          *
//*          Part 3: iteration                                         *
//*          Part 4: CLAIRE-specific structures                        *



// *******************************************************************
// *     Part 2: Memento about CLAIRE to JavaScript                  *
// *******************************************************************

// this is inspired from the main.js file

c1 <: c2(a:integer)  ->  
class c1 extends c2 {
        constructor(a) {
                this.a = a
                }}

fib(x:integer) : integer -> fib(x-1) + fib(x-2)    ->
function fib(x) {
        if (x <= 1) return 1;
        else return fib(x-1) + fib(x-2);
        }

for x in l e(x)  ->
   for (let i = 0; i < l.length; i++) { e(l[i]); }
for x in s e(x)  ->
   for (const s of x) { e(x)}

try e1 catch (set e2)  ->
   try { e1 } catch (e) { if e.belongs(s) {e2}  else {throw e;}} 



// *******************************************************************
// *     Part 3: Log                                                 *
// *******************************************************************

// this code started in August 2023
- file structure borrowed from Java compiler from 2000 :)
- code is 90% completed but was neither loaded nor tested

// reopen on December 31st, 2023 !
step1: finish & load the module

// Jannuary 1-3
- code is finished and loaded :)
- I need to find out how method works in Javascript -> create a small file test.js
  that file will become ClaireKernel.js

// Jan 14th
- code works at the test_j level :) move to jcompile(gw4)
- create test.js to see if my code generation option works !
- support dynamic calls if the only restrictions that could apply are iso-co-domain and belong to the current module
  then we can generate x.m(y_i)

// Jan 19th (Sylvie's bday)
- resume jcompile(gw0) => works !
- need to fix the comment extraction/ restitution

// Jan 20th : work on producing code that node.js can load
- anyObjecty! only allowed for general_error (from error(...))
- got the whole file gw0.js to load !

Key: we call ClaireKernel = kernel module
Hence everything that is not in the current module should have a kernel.

// Jan 27-28th : finish ClaireKernel.js
(a) Class : instanciation and subclass/superclass closure 
    - size_class
(b) belong -> use javascript primitive type recognition

// Jan 29th
- fixed get@Affine which was not diet (two restrictions with same class)
INCROYABLE: Javscript does not support arity !!!!!!
diet means : no other method from the same module has the same domain!
DECISION : implement jsMethod? which is specific to Javascript

// COOL : we have a code that runs ... with some iteration problem


// May 12th : compiled gw1 with success !!!
  However, for c in (C but x) requires explicit typing, which is a bug (cf. 3d in next list)


// ------ summer 2024 : resume the work on the CLAIRE-to-Javascript compiler -------

// August 7th
 (a)  list the methods that should be generic (look for notes) and js_producer s methods  (js in name)
Design goal : 
    (a)  attach JS methods to js_producer
    (b) reuse as much as possible => generic methods
    (c) enrich code_producer 
    code_producer <: producer(
         open_comparators:list[operation],      // list of comparison ops that are inlined (the order matters!!)
         open_operators:list[operation],        // list of arithmetic operators that are inlined
         div_operators:list[operation],         // list of division operators that are inlined
         body:any = 0,                          // used to store the body of the current method
         extension:string,                      // extension for generated files
         comment:string,                        // a string that designates the target language
         interfaces:list)     


generic list:  (see generic analysis below )
============================================
(91) [compile(p:js_producer, m:module) : void
(110) [js_files(p:js_producer,m:module) : void
(129) [js_mod_file(p:js_producer, m:module) : void
(291) [js_file(p:js_producer, f1:string) : void


list of other js* methods:
===========================
(jssystem.cl)  => very specific methods / having 
[make_js_function(p:js_producer,self:lambda,m:method) : void
is different from make_go_function(p:go_producer,self:lambda,%nom:string,m:method) : void
because we used %nom

DECISION: use a generic make_function(p:go_producer,self:lambda,%nom:string,m:method) : void

[function_js_body(self:any,s:class) : void
 => [function_body(p:go_producer, self:any,s:class) : void

[procedure_js_body(m:method, %l:lambda, %body:any,s:class) : void
 => [procedure_body(p:go_producer, self:any,s:class) : void

(jsgen.cl)
ignore the following : small and speficic ....
js_class, js_ident, js_var ...
(158) [jsFunction(p:js_producer,m:method) : string
(176) [jsVariables(p:js_producer,self:list) : any

(226) [bounded_expression_js(self:any) : void
=> made generic with expression!(p:code_producer,x:any,s:class)

expression!(p:go_producer,x:any,s:class) : void
  -> g_expression(x,s)
expression!(p:go_producer,x:any,s:class) : void
  -> j_expression(x)
[bounded_expression(p:code_producer,self:any,s:class) : void
  -> case self (Assign printf("(~I)",expression!(p,self,s)),
                integer (if (self < 0)  printf("(~I)",expression!(p,self,s))    // v3.2.44
                         else expression!(p,self,s)),                           // avoid (2--2)
                float   (if (self < 0.0)  printf("(~I)",expression!(p,self,s))
                         else expression!(p,self,s)),
                any    expression!(p,self,s)) ]

(236) [at_index_js(x:any) : void
=> [at_index(p:js_producer,x:any) : void
  -> case x (integer princ(x - 1), any (expression!(p,x, integer), princ("-1"))) ]

(jsexp.cl)
[j_bool_exp(self:any,pos?:boolean) : void 
=> becomes generic (see below)
[j_belong_exp(a1:any,a2:any) : void
=> belong_exp(p:go_producer/js_producer, a1:any, a2:any)
   Note assumes that if is always expected to return a bool

(jsstat.cl)
(176) [call_j_statement(self:any,v:string,loop:any) : void
reintrant but very specific to JS

TODO from previous notes
------------------------
    (3e) add traces as a PRODUCER slot witg a ctrace method to update it that is not compiled.
         This is clearly useful to debug the c2j compiler but it could be useful with go as well.
         we could even have a Kernel traceStart(p,args) method in the future
         => stupide on a l option pour cela ... vérifier si elle est là ?
         need_debug?(m)  =>  debug_intro(c:go_producer,self:lambda,x:method)      
              => printf("Core.F_Core_db_bind_module(~I,~I,ARGS(~I));~I", 

     (4) there are many useless variables in the generated code, would be nice to get rid of them


// August 10th 2024 : now that we have a CLAIRE4 v4.12 we start

(a) jssystem.cl




generic analysis:  (kill later)
===============================

(91) [compile(p:js_producer, m:module) : void
---------------------------------------------

Go code

// generic version
[compile(p:code_producer, m:module) : void
 ->   OPT.need_modules := {},
      compiler.inline? := true, 
      compiler.n_loc := 0,                // number of lines of code
      compiler.n_warnings := 0,           // number of warnings
      compiler.n_notes := 0,              // number of notes
      let l1:bag := parents(Reader/add_modules(list(m))) in
      (//[3] ==========  START GO COMPILING (~S) with ~S ================ // m, l1,
       OPT.legal_modules := set!(l1),
       p.current := m,                                              // v4: we need to know in which module we are
       gen_files(p,m),                // files to files
       gen_mod_file(p,m),
       l1 := difference(set!(OPT.need_modules), OPT.legal_modules),
       if l1 (warn(),trace(1, "~S should be declared for ~S \n", l1, m))) ]

[compile(p:go_producer, m:module) : void
 ->   BadMethods := list<method>(),       // specific to cross-compiling
      compiler.n_dynamic := 0,            // number of dyn calls
      compiler.n_metheids := 0,           // methods that may return an error (useful for Go)
      p.source := compiler.source / string!(m.name),               // produce code in the <src>/<module> directory
      compile@code_producer(p,m),
      trace(1, "~S: ~A lines of code compiled. ~A warnings, ~A notes. ~A dynamic calls, ~A% exception-ready methods\n",
              m, compiler.n_loc, compiler.n_warnings, compiler.n_notes, compiler.n_dynamic, 
            (if (compiler.n_methods = 0) 0 else (100 * compiler.n_metheids) / compiler.n_methods))) ]


JS code
[compile(p:js_producer, m:module) : void
 ->  compile@code_producer(p,m),
     trace(1, "~S: ~A lines of code compiled. ~A warnings, ~A notes.\n",
              m, compiler.n_loc, compiler.n_warnings, compiler.n_notes) ]

VERDICT: this is indeed generic (nothing much to do )

(110) [js_files(p:js_producer,m:module) : void
----------------------------------------------
Go code:
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
       (//[1] ++++ Compiling the file ~A.cl [v. 4.~A - safety:~A] // x, compiler.version, compiler.safety,
        if (x = string!(m.name))
           Cerror("[211]  ~S cannot be used both as a file and module name",x),
        OPT.level := 1, // debug - to remove
        p.current_file := x,       // CLAIRE4 : keep the file name handy 
        gen_file(p, m.source / x, p.source / x)),
     end(m) ]

JS code:
[js_files(p:js_producer,m:module) : void
 ->  //[0] ==== Generate ~A files for module ~S [verbose = ~A, Opt? = ~S] // PRODUCER.comment, m, verbose(),compiler.optimize?,
     OPT.instructions := list<any>(),
     OPT.properties := set<property>(),
     OPT.objects := list<object>(),
     OPT.functions := list<any>(),
     OPT.need_to_close := set<any>(),
     begin(m),
     for x in m.made_of
       (//[0] ++++ Compiling the file ~A.cl [v. 4.~A - safety:~A] // x, compiler.version, compiler.safety,
        if (x = string!(m.name))
           Cerror("[211]  ~S cannot be used both as a file and module name",x),
        OPT.level := 1, // debug - to remove
        p.current_file := x,       // CLAIRE4 : keep the file name handy 
        js_file(p, m.source / x)),
     end(m) ]

VERDICT: should be generic(code_producer)



(129) [js_mod_file(p:js_producer, m:module) : void
--------------------------------------------------

Generic code
[gen_mod_file(p:code_producer, m:module) : void
 -> let prt := fopen(((PRODUCER.source /+ *fs*) /+ string!(m.name) /+ "-meta") /+ PRODUCER.Generate/extension, "w"),
        s := string!(m.name) in
      (//[2] ==== generate file for module ~S ==== // m,
       OPT.outfile := prt,
       start_file(p,s,m,true),            // true tells this is the module file
       use_as_output(OPT.outfile),
       gen_classes(p,m),                 // v4.0 : keep the class/struct definition in the module
       gen_objects(p,m),
       gen_functions(p),
       gen_meta_load(p,m),                 // reflective description (class, methods, vars)
       breakline(),
       close_block(),
       breakline(),                       // v3.0.3
       fclose(OPT.outfile)) ]

Go code:
[gen_functions(p:go_producer, m:module) : void
  -> nil ]   //   not used for go 


[gen_mod_file(p:go_producer, m:module) : void
 -> gen_mod_file@code_producer(p,m,false)
    if (compiler.safety > 4) //[1] ===== [CROSS]  ~A BAD METHODS : ~S  // length(BadMethods), BadMethods
    ]

JS code: (no code : same as generic !)
[js_mod_file(p:js_producer, m:module) : void
 -> let prt := fopen((p.source / string!(m.name)) /+ PRODUCER.Generate/extension, "w"),
        s := string!(m.name) in
      (//[0] ==== generate file for module ~S ==== // m,
       OPT.outfile := prt,
       start_file(p,s,m),                
       use_as_output(OPT.outfile),
       gen_classes(p,m),                 // v4.0 : keep the class/struct definition in the module
       gen_objects(p,m),
       gen_functions(p),
       gen_meta_load(p,m),                 // reflective description (class, methods, vars)
       breakline(),
       fclose(OPT.outfile)) ]

VERDICT=> possible with a dummy gen_functions(p) and two parameters

(291) [js_file(p:js_producer, f1:string) : void
------------------------------------------------

unclear since a heavy method but maybe a common frame with additional 
  - bool to open/close OPT.outfile = outfile?(p)
  - inner method gen_instruction(p:go_producer,%instruction:any)

Go code:
[gen_file(p:code_producer, f1:string,f2:string, outfile?:boolean)
 -> let p1 := fopen(f1 /+ ".cl", "r"), 
        b := reader.Reader/toplevel,
        p0 := reader.Reader/fromp,         // b, p0: reading context when we start
        prev_comment  := "" in
       (if outfile? OPT.outfile := fopen(f2 /+ p.Generate/extension, "w"),
        reader.Reader/toplevel := false,
        compiler.loading? := true,
        n_line() := 1,
        reader.external := f1,
        reader.Reader/fromp := p1,                     // <yc> ensures automatic fclose !
        start_file(p,f1,module!(),false),               // CLAIRE 4 : always add a header
        let %instruction := Reader/readblock(p1) in
          while not(%instruction = Reader/eof)
             (prev_comment := gen_instruction(p,%instruction,prev_comment),
              %instruction := Reader/readblock(p1)),
        compiler.n_loc :+ n_line(),
        fclose(p1),
        compiler.loading? := false,
        // restore reading context
        reader.Reader/toplevel := b,
        reader.external := "toplevel",
        reader.Reader/fromp := p0,
        if outfile? fclose(OPT.outfile))] 

// the only specific part is the instruction (because of comment management)
[gen_instruction(p:go_producer,%instruction:any,prev_comment:string) : string  
  -> (if (%instruction % string)            // we have found a comment
               let pp := use_as_output(OPT.outfile) in
                 (printf("\n//~A", %instruction),
                  use_as_output(pp)) 
      else OPT.instructions :add c_code(%instruction, void),
      prev_comment)]

Previous JS code (2.10)
// this is the basic file cross_compiler, which translates from claire to javascript
// [note: should be generic] : this method should be attacted to code_producer in the go version as well
// for Javascript, like Java, this pass does not produce code but fills the stacks
[js_file(p:js_producer, f1:string) : void
 -> let p1 := fopen(f1 /+ ".cl", "r"), 
        b := reader.Reader/toplevel,         // we are about to divert the top-level flow
        p0 := reader.Reader/fromp,         // b, p0: reading context when we start
        prev_comment  := "" in
       (reader.Reader/toplevel := false,
        compiler.loading? := true,
        n_line() := 1,
        reader.external := f1,
        reader.Reader/fromp := p1,                     // <yc> ensures automatic fclose !
        let %instruction := Reader/readblock(p1) in
          while not(%instruction = Reader/eof)
            (if (%instruction % string)            // we have found a comment => put on methods stack
               (trace(5,"READ COMMENT [~A]:~A\n", %instruction, length(%instruction)),
                prev_comment :/+ ("\n// " /+ %instruction))
             else 
               (if (%instruction % Defclass) trace(0,"READ DEFCLASS ~S [comment :~A]\n",%instruction,length(prev_comment)),
                if (%instruction % Defmethod) p.methods : add prev_comment    // store the comment associated to m
                else if (%instruction % Defclass) 
                   value(iClaire/ident(%instruction as Defclass)).comment := prev_comment,
                prev_comment := "",
                OPT.instructions :add c_code(%instruction, void)),   // instructions stack => deferred module generation
             %instruction := Reader/readblock(p1)),
       compiler.n_loc :+ n_line(),
       compiler.loading? := false,
       // restore reading context
       reader.Reader/toplevel := b,
       reader.external := "toplevel",
       reader.Reader/fromp := p0) ]
JS code:
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
            (iClaire/ident(%instruction as Defclass)).comment := prev_comment,
             prev_comment := "",
             OPT.instructions :add c_code(%instruction, void)),   // instructions stack => deferred module generation
        prev_commment ]

// ----------- generic bool_exp candidates ----------------------------------------------------

// made generic with print_true
[print_true(p:go_producer) : void
   -> princ("CTRUE)")]

[bool_exp(p:code_producer,self:any,pos?:boolean) : void 
  -> if (self = true) print_true(p)    // v4.0.
     else printf("(~I ~I ~I)", g_expression(self, boolean), sign_equal(pos?), print_true(p))) ]

// strange : not clear why we should see a C_cast here
[bool_exp(p:code_producer,self:C_cast,pos?:boolean) : void 
  -> bool_exp(self.arg,pos?) ]


// If is supported with IfThenElse (means that all terms will be evaluated) [generic]
[bool_exp(p:code_producer,self:If,pos?:boolean) : void
 -> if self.other
        printf("(~I ? ~I : ~I)", bool_exp(p,self.test, true),
                                 bool_exp(p,self.arg, pos?),
                                 bool_exp(p,self.other, pos?))
     else printf("(~I ~I ~I)", bool_exp(p,self.test, pos?),
                 Generate/sign_or(not(pos?)), bool_exp(p,self.arg, pos?)) ]

// for a AND, we can used the && C operation [generic]
[bool_exp(p:code_producer,self:And,pos?:boolean) : void
 -> let l := self.args, m := length(l), n := 0, %l := OPT.level in
       (OPT.level :+ 1,
        for x in l
          (n :+ 1,
           if (n = m) bool_exp(p,x, pos?)
           else (printf("(~I ~I ", bool_exp(p,x, pos?), Generate/sign_or(not(pos?))),
                 OPT.level :+ 1,
                 breakline())),
        for x in (2 .. m) princ(")"),
        OPT.level := %l)  ]

// idem for OR: we use ||  [generic]
[bool_exp(p:code_producer,self:Or,pos?:boolean) : void
 -> let l := self.args, m := length(l), n := 0, %l := OPT.level in
       (OPT.level :+ 1,
        for x in l
          (n :+ 1,
           if (n = m) bool_exp(p,x, pos?)
           else (printf("(~I ~I ", bool_exp(p,x, pos?), Generate/sign_or(pos?)),
                 OPT.level :+ 1,
                 breakline())),
        for x in (2 .. m) princ(")"),
        OPT.level := %l) ]


// membership optization [generic, through belong_exp]
[bool_exp(p:go_producer,self:Call,pos?:boolean) : void
 -> let p := self.selector in
       (if (p = %) printf("(~I ~I ~I)", belong_exp(p,self.args[1], self.args[2]), sign_equal(pos?),
                          print_true(p)) 
        else printf("(~I ~I ~I)", expression!(p,self, boolean), sign_equal(pos?), print_true(p))) ]

[bool_exp(p:js_producer,self:Call,pos?:boolean) : void
 -> let p := self.selector in
       (if (p = %) printf("(~I ~I true)", belong_exp(p,self.args[1], self.args[2]), sign_equal(pos?)) 
        else printf("(~I ~I true)", j_expression(self), sign_equal(pos?)))]


// generic !
[bool_exp(p:code_producer,self:Call_method1,pos?:boolean) : void
 -> let m := self.arg, a1 := self.args[1] in
       (if (m = *not*) bool_exp(ap,1, not(pos?))     // v3.3.12 - was :  & a1 % to_CL
        else if (m = *known*) equal_exp(p,a1, not(pos?), unknown, true)
        else if (m = *unknown*) equal_exp(p,a1, pos?, unknown, true)
        else if (m.range <= boolean)
           printf("(~I ~I ~I)", expression!(p,self, boolean), sign_equal(pos?),print_true(p))
        else bool_exp@any(self, pos?)) ]

// same thing for two arguments functions => generic
// equal_exp is in gogen.cl
[bool_exp(c:code_producer,self:Call_method2,pos?:boolean) : void
 -> let m := self.arg, p := m.selector, lop := c.Generate/open_comparators,
        a1 := self.args[1], a2 := self.args[2] in
      (if (p = !=) equal_exp(c,a1, not(pos?), a2, false)
       else if (p = identical?) equal_exp(c,a1, pos?, a2, true)
       else if (p = =) equal_exp(c,a1, pos?, a2, false)
       else if (m = m_member) printf("(~I ~I ~I)",belong_exp(c,a1,a2),sign_equal(pos?),print_true(c)) 
       else if (p % lop & domain!(m) % {float,integer})
           printf("(~I ~I ~I)", expression!(c,a1, domain!(m)),
                  (if pos? print(p)
                   else print(lop[((get(lop, p) + 1) mod 4) + 1])),  // lop = (<, >, >=, <=)
                  expression!(c,a2, domain!(m)))
       else if (m = *nth_integer*) // bit vectors  (a1 is a integer seen as a set, a2 is an integer)
           printf("(BitVectorContains(~I,~I) ~I ~I)", 
                      expression!(c,a1,integer), expression!(c,a2,integer), sign_equal(pos?),
                      print_true(c))
        else if (p = inherit? & domain!(m) = class)
         printf("(~I.IsIn(~I) ~I ~I)", 
                  expression!(c,a1,class), expression!(c,a2,class),sign_equal(pos?),print_true(c))
       else 
          printf("(~I ~I ~I)", expression!(c,self, boolean), sign_equal(pos?),print_true(c))) ]



// belong_exp is not generic but is a producer method
// assumption : we require a boolean (s = boolean input constraint)
[belong_exp(p:go_producer,a1:any,a2:any) : void
 ->  if (static_type(a2) <= type) 
         printf("~I.Contains(~I)",  g_expression(a2,type), g_expression(a1,any))
     else if (static_type(a2) <= integer & static_type(a1) <= integer)
        printf("BitVectorContains(~I,~I)", g_expression(a2,integer), g_expression(a1,integer))
     else if (static_type(a2) <= list | static_type(a2) <= array)
        printf("~I.Contain_ask(~I)",  g_expression(a2,list), g_expression(a1,any))
     else printf("~I~IF_BELONG(~I,~I)~I",  
                 cast_prefix(EID,boolean),
                 preCore?(),
                 g_expression(a1,any),
                 g_expression(a2,any),
                 cast_post(EID,boolean)) ]


