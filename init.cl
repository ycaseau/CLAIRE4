(printf("Hello CLAIRE4, this is our init.cl file\n"))

// Mac version
*where* :: "/Users/ycaseau/claire/v4.0/go"                      // where the init file is
*output* :: "/Users/ycaseau/claire/v4.0/go/src"
*meta* :: "/Users/ycaseau/Dropbox/src/clairev4.12/src/meta"            // source files on dropbox (v2)
*compile* :: "/Users/ycaseau/Dropbox/src/clairev4.12/src/compile"      // source files on dropbox (v2)
*bsrc* :: "/Users/ycaseau/claire/v4.0/test/nonreg"
*tsrc* :: "/Users/ycaseau/claire/v4.0/test/perf"
*rsrc* :: "/Users/ycaseau/claire/v4.0/test/rules"

// these are the global variables expected by the compiler
RELEASE:float :: 1.2    // version of August 8th, 2024

// note : we should check that these optimizations are still valid


// additions  (comment out what you don't want)
begin(Language)

(#if (version() = 1.0) claire/IterateFast :: property())

(#if (version() = 1.0) claire/If? <: If())             // v4.12 conditional comment

end(Language)

begin(Optimize)
// Compile/Super_cast <: Compile/C_cast()                // used for Super
claire/DKEEP:any :: unknown

end(Optimize)

// stub for compiler new architecture
begin(Generate)

/*

// debug (to remove later)
[g_test(l:lambda) : void
  -> OPT.in_method := unknown,
     OPT.Optimize/use_string_update := false,   // v3.3.46
     OPT.Optimize/max_vars := 0,
     OPT.legal_modules := set!(module.instances),
     OPT.outfile := stdout,
     compiler.inline? := true, 
     PRODUCER.current := claire,
     trace(0,"\n---- code produced by the generator ------------------- \n"),
     make_c_function(l,"test", void),
     OPT.in_method := unknown ]

// this this the heart of the compiler : compiles module m into a set of files
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
       gen_files(p,m),                // (generic) files to files
       gen_mod_file(p,m),             // (generic) module file
       l1 := difference(set!(OPT.need_modules), OPT.legal_modules),
       if l1 (warn(),trace(1, "~S should be declared for ~S \n", l1, m))) ]

// the first part is to generate the files associated to the module claire files
// gen_files prepare the optimizer (OPT) and calls gen_file for each file
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
        // OPT.level := 1, // debug - to remove
        p.current_file := x,       // CLAIRE4 : keep the file name handy 
        gen_file(p, m.source / x, p.source / x)),
     end(m) ]

// main method: generate a target file associated to a CLAIRE file
[gen_file(p:code_producer, f1:string,f2:string)
 -> let p1 := fopen(f1 /+ ".cl", "r"), 
        b := reader.Reader/toplevel,
        p0 := reader.Reader/fromp,         // b, p0: reading context when we start
        out? := outfile?(p),
        prev_comment  := "" in
       (if out? OPT.outfile := fopen(f2 /+ p.Generate/extension, "w"),
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
        if out? fclose(OPT.outfile))] 

// generate a file associated with module m (called meta-m)
[gen_mod_file(p:code_producer, m:module) : void
 -> let prt := fopen(((p.source /+ *fs*) /+ string!(m.name) /+ "-meta") /+ p.Generate/extension, "w"),
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


// test the compiling of a method
// e.f. g_test(foo @ any)
[compile_method(p:code_producer,m:method) : void
  -> when l := get(formula,m) in
        (//[0] ---- Compiling ~S with following definition ---- // m,
        pretty_print(body(l)),
        OPT.in_method := m,
        OPT.Optimize/use_string_update := false,   // v3.3.46
        OPT.Optimize/max_vars := 0,
        OPT.legal_modules := set!(module.instances),
        OPT.outfile := stdout,
        compiler.inline? := true, 
        p.current := claire,
        trace(0,"\n---- code produced by the optimizer -------------------\n"),
        pretty_print(c_strict_code(formula(m).body,class!(m.range))),
        trace(0,"\n---- code produced by the generator ------------------- \n"),
        gen_function(p,formula(m),"test",m),
        OPT.in_method := unknown ) ]

// generate an expressions with boundaries (parenthesis if necessary)
[bounded_expression(p:code_producer,self:any,s:class) : void
  -> case self (Assign printf("(~I)",expression!(p,self,s)),
                integer (if (self < 0)  printf("(~I)",expression!(p,self,s))    // v3.2.44
                         else expression!(p,self,s)),                           // avoid (2--2)
                float   (if (self < 0.0)  printf("(~I)",expression!(p,self,s))
                         else expression!(p,self,s)),
                any    expression!(p,self,s)) ]

// CLAIRE uses 1 .. n as array range, target languages use 0 .. n-1
[at_index(p:code_producer,x:any) : void
  -> case x (integer princ(x - 1), any (expression!(p,x, integer), princ("-1"))) ]



[b_expression(p:code_producer,self:any,pos?:boolean) : void 
  -> if (self = true) princ("true")    // v4.0.
     else printf("(~I ~I ~I)", expression!(p,self, boolean), sign_equal(pos?), print_true(p)) ]

// strange : not clear why we should see a C_cast here
[b_expression(p:code_producer,self:C_cast,pos?:boolean) : void 
  -> b_expression(p,self.arg,pos?) ]


// If is supported with IfThenElse (means that all terms will be evaluated) [generic]
[b_expression(p:code_producer,self:If,pos?:boolean) : void
 -> if self.other
        printf("(~I ? ~I : ~I)", b_expression(p,self.test, true),
                                 b_expression(p,self.arg, pos?),
                                 b_expression(p,self.other, pos?))
     else printf("(~I ~I ~I)", b_expression(p,self.test, pos?),
                 Generate/sign_or(not(pos?)), b_expression(p,self.arg, pos?)) ]

// for a AND, we can used the && C operation [generic]
[b_expression(p:code_producer,self:And,pos?:boolean) : void
 -> let l := self.args, m := length(l), n := 0, %l := OPT.level in
       (OPT.level :+ 1,
        for x in l
          (n :+ 1,
           if (n = m) b_expression(p,x, pos?)
           else (printf("(~I ~I ", b_expression(p,x, pos?), Generate/sign_or(not(pos?))),
                 OPT.level :+ 1,
                 breakline())),
        for x in (2 .. m) princ(")"),
        OPT.level := %l)  ]

// idem for OR: we use ||  [generic]
[b_expression(p:code_producer,self:Or,pos?:boolean) : void
 -> let l := self.args, m := length(l), n := 0, %l := OPT.level in
       (OPT.level :+ 1,
        for x in l
          (n :+ 1,
           if (n = m) b_expression(p,x, pos?)
           else (printf("(~I ~I ", b_expression(p,x, pos?), Generate/sign_or(pos?)),
                 OPT.level :+ 1,
                 breakline())),
        for x in (2 .. m) princ(")"),
        OPT.level := %l) ]


// membership optization [generic, through belong_exp]
[b_expression(c:code_producer,self:Call,pos?:boolean) : void
 -> let p := self.selector in
       (if (p = %) printf("(~I ~I ~I)", belong_exp(p,self.args[1], self.args[2],boolean), sign_equal(pos?),
                          print_true(p)) 
        else printf("(~I ~I ~I)", expression!(p,self, boolean), sign_equal(pos?), print_true(p))) ]

// generic !
[b_expression(p:code_producer,self:Call_method1,pos?:boolean) : void
 -> let m := self.arg, a1 := self.args[1] in
       (if (m = *not*) b_expression(p,a1, not(pos?))     // v3.3.12 - was :  & a1 % to_CL
        else if (m = *known*) equal_exp(p,a1, not(pos?), unknown, true)
        else if (m = *unknown*) equal_exp(p,a1, pos?, unknown, true)
        else  printf("(~I ~I ~I)", expression!(p,self, boolean), sign_equal(pos?), print_true(p))) ]


// same thing for two arguments functions => generic
// equal_exp is in gogen.cl
[b_expression(c:code_producer,self:Call_method2,pos?:boolean) : void
 -> let m := self.arg, p := m.selector, lop := c.Generate/open_comparators,
        a1 := self.args[1], a2 := self.args[2] in
      (if (p = !=) equal_exp(c,a1, not(pos?), a2, false)
       else if (p = identical?) equal_exp(c,a1, pos?, a2, true)
       else if (p = =) equal_exp(c,a1, pos?, a2, false)
       else if (m = m_member) printf("(~I ~I ~I)",belong_exp(c,a1,a2,boolean),sign_equal(pos?),print_true(c)) 
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

[belong_exp(p:go_producer,a1:any,a2:any,s:class) : void
  ->  if (static_type(a2) <= type) 
         printf("~I~I.Contains(~I)~I",  cast_prefix(boolean,s),
                g_expression(a2,type), g_expression(a1,any),cast_post(boolean,s))
     else if (static_type(a2) <= integer & static_type(a1) <= integer)
        printf("~IBitVectorContains(~I,~I)~I",  cast_prefix(boolean,s),
                g_expression(a2,integer), g_expression(a1,integer),cast_post(boolean,s))
     else if (static_type(a2) <= list | static_type(a2) <= array)
        printf("~I~I.Contain_ask(~I)~I", cast_prefix(boolean,s),
                g_expression(a2,list), g_expression(a1,any),cast_post(boolean,s))
     else printf("~I~IF_BELONG(~I,~I)~I",  
                 cast_prefix(EID,s),
                 preCore?(),
                 g_expression(a1,any),
                 g_expression(a2,any),
                 cast_post(EID,s)) ]

[gen_functions(p:go_producer, m:module) : void
  -> nil ] 

[outfile?(p:go_producer) : boolean
  -> true]

// generic compiler method
[expression!(p:go_producer, self:any,s:class) : void
 -> g_expression(self,s) ]

// made generic with print_true
[print_true(p:go_producer) : void
   -> princ("CTRUE")]

// compiles one instrtuction from the CLAIRE file
// comments are printed (thanks to OPT.outfile) others are stacked in OPT.instructions  
[gen_instruction(p:go_producer,%instruction:any,prev_comment:string) : string  
  -> (if (%instruction % string)            // we have found a comment
               let pp := use_as_output(OPT.outfile) in
                 (printf("\n//~A", %instruction),
                  use_as_output(pp)) 
      else OPT.instructions :add c_code(%instruction, void),
      prev_comment)]

[gen_function(p:go_producer,self:lambda,%nom:string,m:method) : void
  -> let typeOK := check_range(m,self.body),
         s := class!(m.range),
         %body := c_strict_code(self.body,s),
         throw? := g_throw(%body) in
      (//[2] [~A] ~S: => simple=~S, throw=~S // n_line(), m, simple_body?(%body),throw?,
       compiler.n_methods :+ 1,
       p.varsym := 0,                                            // resets the ID for distinct vars
       if (m.status = -1) (m.status := (if throw? 1 else 0))     // nice case we have not seen m yet.
       else if (throw? != can_throw?(m))                         // avoids generating go code that will break
          (warn(),
           //[1] [CROSS] ~S body produces an error (g_throw = true) while status is 0 // m,
           if (m.status = 0) BadMethods :add m
           else throw? := true                         // avoid generating wrong code (rest of world assumes EID since m.status was 1)
          ),
       use_as_output(OPT.outfile),
       if ((typeOK | compiler.safety >= 2) & not(throw?) & (m.selector != self_eval))  // happy with the type inference => native function
          (//[5] --- Procedure generation (can throw = ~S) // throw?,
           if p.debug? printf("// DEBUG: g_throw=~S from body=~S ~I",throw?,%body, breakline()),
           generate_function_start(PRODUCER, self, s, m, %nom),        
           new_block(),
           if (need_debug?(m) |  OPT.profile? | not(simple_body?(%body)) | s = void) 
              procedure_body(p,m,self,%body,s)
           else (if p.debug? printf("// use function body compiling ~I",breakline()),
                 function_body(p,%body,s)))
       else (//[3] --- EID function generation (can throw = ~S) // throw?,
             throw? := true,  
             compiler.n_metheids :+ 1,                                     // this is the EID pathd
             generate_function_start(PRODUCER, self, EID, m, %nom),        
             new_block(),
             eid_body(m,%body,typeOK,s)),
       close_block(),
       generate_eid_function(self,m,throw?),
       if (m.selector = self_eval) generate_eval_function(self,m),
       use_as_output(stdout)) ]


// new open properties
(open(gen_instruction) := 3)
(open(outfile?) := 3)
(open(gen_classes) := 3)
(open(gen_objects) := 3)
(open(gen_functions) := 3)
(open(gen_meta_load) := 3)
(open(gen_function) := 3)  
(open(print_true) := 3)
(open(belong_exp) := 3)


[function_body(c:go_producer,self:any,s:class) : void
  -> let %ret := (if (s != void) "return " else "") in
      (if (s = boolean & (case self (Call_method (let p := self.arg.selector in 
                                      (p = = | p = < | p = > | p = >= | p = <=)  ))))       
                               // this is an old optimization - there is a debate if this is still needed with CLAIRE4
                               // reintroduced in v4.0.7 for mSend, but only for direct comparisons
          printf("if ~I {return CTRUE~I} else {return CFALSE}",b_expression(c,self,true),breakline())
       else if (c_type(self) = void & s != void)
         printf("~I~Ireturn ~I~I",
                 g_expression(self,void), breakline(),
                 g_expression(unknown,s), breakline())
      else printf("~A ~I~I", %ret, g_expression(self,s),breakline())) ]

// generate nice code for If function (inspired from g_statement@If)
[function_body(c:go_producer,self:If, s:class) : void
  -> printf("if ~I ~I",
            b_expression(c,self.test, true),
            new_block("body If")),
    function_body(c,self.arg,s),
    if (self.other = nil) close_block()
    else if (self.other % If) 
      printf("~I else ~I",finish_block(), function_body(c,self.other,s))  
    else if (s != void | not(designated?(self.other)))
        printf("} else {~I~I~I", breakline(),
                     function_body(c,self.other,s),
                     close_block("body If"))
    else close_block("body If") ]

// generate nice code for a Do
[function_body(c:go_producer,self:Do, s:class) : void
  ->  let l := self.args, %length := length(l), m := 0 in
        ( for x in l
            (m :+ 1,
             if (m = %length) function_body(c,x,s)
             else statement(x, void, "Unused", false)))
  ]

// default complex case : create a variable "Result"
[procedure_body(c:go_producer,m:method, %l:lambda, %body:any,s:class) : void
  ->  if need_debug?(m) debug_intro(PRODUCER,%l,m),
      if PRODUCER.debug? printf("// procedure body, with s = ~S~I",s,breakline()),
      if (s != void) 
         (var_declaration("Result",s,1),
          statement(%body,s,"Result",false))
      else statement(%body,void,"Unused",false),
      return_result(PRODUCER,s,m,"Result") ] 

  

(start_file.open := 3)
(compile.open := 3)
(b_expression.open := 3)
(inline_exp.open := 3)
*/

[gen_mod_file(p:code_producer, m:module) : void
 -> let prt := fopen(((p.source /+ *fs*) /+ modfile_name(p,m)) /+ p.Generate/extension, "w"),
        s := string!(m.name) in
      (//[2] ==== generate ~A file for module ~S ==== // p.comment,m,
       OPT.outfile := prt,
       start_file(p,s,m,true),            // true tells this is the module file
       use_as_output(OPT.outfile),
       gen_classes(p,m),                 // v4.0 : keep the class/struct definition in the module
       gen_objects(p,m),
       gen_functions(p,m),
       gen_meta_load(p,m),                 // reflective description (class, methods, vars)
       breakline(),
       close_block(),
       breakline(),                       // v3.0.3
       fclose(OPT.outfile)) ]

// name of the file generated for the module
[modfile_name(p:go_producer,m:module) : string
  -> string!(m.name) /+ "-meta"]
(modfile_name.open := 3)

end(Generate)

 // end of additions

/* debug for bugs  => foo(1,,3)bat()
[foo(x:integer,y:integer) 
   -> x + y]

// our sum macro  => sum(list{size(c) | c in class})
sum(l:list[integer]) : integer
 => (let x := 0 in (for y in l x :+ y, x))

 // a list arg method to see how it is compiled
 bar(l:listargs) : integer
  -> sum(list{(x * x) | x in l})

[claire/strange(m:module) : list
  -> let l := list{ m2 in (Reader/add_modules(list(m)) but m) |
                   (m2.made_of | m2 = Kernel) } in
       // clean_duplicates(l)
       l ]
       
*/


// ***************************************************************************
// *    Part 1: Modules & compiler environment                               *
// ***************************************************************************

// meta files are now the "official" github directory
(for m in {Core,Language,Reader} source(m) := *meta*,
 for m in {Optimize,Generate} source(m) := *compile*)

// where we want to generate the go code
(when c := get_value("compiler") in 
   (c.safety := 5,
    verbose() := 1,
    source(c) := *output*))


// ***************************************************************************
// *    Part 2: Performance test modules                                     *
// ***************************************************************************

// these are the performance test files of 2020
mFib :: module(part_of = claire,
              source = *tsrc*,
              uses = list(Reader,mClaire),
              made_of = list("testFib"))

mList :: module(part_of = claire,
              source = *tsrc*,
              uses = list(Reader,mClaire),
              made_of = list("testList"))

mSet :: module(part_of = claire,
              source = *tsrc*,
              uses = list(Reader,mClaire),
              made_of = list("testSet"))

mDict :: module(part_of = claire,
              source = *tsrc*,
              uses = list(Reader,mClaire),
              made_of = list("testDict"))

mObj :: module(part_of = claire,
              source = *tsrc*,
              uses = list(Reader,mClaire),
              made_of = list("testObj"))

mSend :: module(part_of = claire,
              source = *tsrc*,
              uses = list(Reader,mClaire),
              made_of = list("testSend"))

mCopy :: module(part_of = claire,
              source = *tsrc*,
              uses = list(Reader,mClaire),
              made_of = list("testCopy"))

// ***************************************************************************
// *    Part 3: Bugs for CLAIRE                                              *
// ***************************************************************************

// bu0 is a file of bugs that used to crash nastyly but now should trigger a panic error
bu0 :: module( uses = list(Reader), source = *bsrc*,
               made_of = list("bug0"))

// parsing bugs: things that cannot get read right
bu1 :: module( uses = list(Reader), source = *bsrc*,
               made_of = list("bstub","bug1"))

// array related bugs
bu2 :: module( uses = list(Reader), source = *bsrc*, 
               made_of = list("bstub", "bug2"))

// table related bugs
bu3 :: module( uses = list(Reader), source = *bsrc*,
               made_of = list("bstub", "bug3"))

// iteration of a union (interpreted) and other patterns
bu4 :: module( uses = list(Reader), source = *bsrc*,
               made_of = list("bstub", "bug4"))

// bugs with floats
bu5 :: module( uses = list(Reader), source = *bsrc*,
               made_of = list("bstub", "bug5"))

// bugs with class & method definitions
bu6 :: module( uses = list(Reader), source = *bsrc*,
               made_of = list("bstub", "bug6"))

// bugs with worlds
bu7 :: module( uses = list(Reader), source = *bsrc*,
               made_of = list("bstub", "bug7"))

// bugs with instantiation & primitive types
bu8 :: module( uses = list(Reader), source = *bsrc*,
               made_of = list("bstub", "bug8"))

// reversible cells from CLP, untyped version
bu9 :: module( uses = list(Reader), source = *bsrc*,
               made_of = list("bstub", "bug9"))

// reversible cells from CLP, typed version
bu10 :: module( uses = list(Reader), source = *bsrc*,
               made_of = list("bstub", "bug10"))

// famous examples (stack example, doc examples ...)
bu11 :: module( uses = list(Reader), source = *bsrc*,
                made_of = list("bstub", "bug11"))

// bug with tuples, lists and sets
bu12 :: module( uses = list(Reader), source = *bsrc*,
                made_of = list("bstub","bug12"))

// test file for handling unknown & inverses
bu13 :: module( uses = list(Reader), source = *bsrc*,
                made_of = list("bstub", "bug13"))

// test file for compilation bug (works with claire3)
bu14 :: module( uses = list(Reader), source = *bsrc*,
                made_of = list("bstub", "bug14"))

// another test file for testing compiler
bu15 :: module( uses = list(Reader), source = *bsrc*,
                made_of = list("bstub", "bug15"))

// sudoku example : shown in the tutorial - good example of rules & branch
bu16 :: module( uses = list(Reader), source = *bsrc*,
                made_of = list("bstub", "sudoku"))

// other examples from rge claire manual : need to put in the doc - good example of rules & branch
bu17 :: module( uses = list(Reader), source = *bsrc*,
                made_of = list("bstub", "manual"))



// ***************************************************************************
// *    Part 4: Simple rule examples                                              *
// ***************************************************************************

// dinner
mDinner :: module( uses = list(Reader), source = *rsrc*,
                    made_of = list("dinner"))


// filter
mFilter :: module( uses = list(Reader), source = *rsrc*,
                   made_of = list("filter"))

// monkey
mMonkey :: module( uses = list(Reader), source = *rsrc*,
                   made_of = list("monkey"))

// zebra
mZebra :: module( uses = list(Reader), source = *rsrc*,
                  made_of = list("zebra"))

// airline - WIP (old CLAIRE 2 example)
mAirline :: module( uses = list(Reader), source = *rsrc*,
                  made_of = list("Airline"))

// these are the old non-regression tests files (refreshed in July 2021)
(printf("Done. \n"))

