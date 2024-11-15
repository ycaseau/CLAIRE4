//+-------------------------------------------------------------+
//| CLAIRE                                                      |
//| jsgen.cl                                                    |
//| Copyright (C) 2023 - 2024 Yves Caseau. All Rights Reserved  |
//+-------------------------------------------------------------+

// *******************************************************************
// * Contents                                                        *
// *     Part 1: definition of the code producer                     *
// *     Part 2: utilities for file generation                       *
// *     Part 3: interface declarations                              *
// *     Part 4: use of language dependent patterns (macros)         *
// *     Part 5: Utilities                                           *
// *******************************************************************

// *******************************************************************
// *       Part 1: definition of the code producer               *
// *******************************************************************

// definition of the instance
// to do : update the reserved names progressively - note that classes do not need protection since
// ClaireX is added to X
claire/JS_PRODUCER  :: js_producer(
  Generate/open_comparators = list(<, >, >=, <=),     // do not change -> goexp uses the specific order !
  Generate/open_operators = list(+,-,*,>>),
  Generate/div_operators = list(/,mod),
  Generate/extension = ".js",
  comment = "Javascript",
  source = "node",                                 // default source generation directory
  bad_names = list(
    symbol!("do"), symbol!("if"), symbol!("and"), symbol!("or"),symbol!("not"),
    symbol!("printf"), symbol!("void"), Optimize/Pattern.name,
    symbol!("return"), symbol!("new"), symbol!("default"),
    symbol!("private"), symbol!("operator"), symbol!("EID"),
    symbol!("Handle"), symbol!("import"),symbol!("var"),
    symbol!("catch"), symbol!("stdout"), symbol!("stdin"),
    symbol!("break"), symbol!("char"), symbol!("interface"), symbol!("EOF"),
    symbol!("System"), symbol!("delete"), symbol!("package"),
    symbol!("abstract"), symbol!("final"), symbol!("system_object"),
    symbol!("range"), symbol!("register"),symbol!("template")),                 // v3.3.36
 good_names = list(
    symbol!("DO"),symbol!("IF"),symbol!("ClaireAnd"),symbol!("ClaireOr"),symbol!("NOT"),
    symbol!("PRINTF"), symbol!("ClaireVoid"), symbol!("ClairePattern"),
    symbol!("RETURN"), symbol!("NEW"), symbol!("Default"),
    symbol!("PRIVATE"), symbol!("ClaireOperator"), symbol!("ClaireEID"),
    symbol!("ClaireHandle"), symbol!("ClaireImport"),symbol!("ClaireVar"),
    symbol!("CATCH"), symbol!("STDOUT"), symbol!("STDIN"),
    symbol!("BREAK"), symbol!("ClaireChar"), symbol!("ClaireInterface"), symbol!("_eof"),
    symbol!("Core"), symbol!("Delete"), symbol!("ClairePackage"),
    symbol!("ABSTRACT"), symbol!("Final"),symbol!("SystemObject"),
    symbol!("Range"), symbol!("ClaireRegister"),symbol!("ClaireTemplate"))     // v3.3.46
 )

// think about a way to make this generic
[claire/jtrace(l:listargs) : void 
  -> for p in l JS_PRODUCER.traces :add p]

// this should be a pragma / comment in or out
// (jtrace(gw0/go,gw0/run,gw0/solve,gw0/getNeed,gw0/getProd))


// NOTE: this code should be defined at the code_producer level (refactoring)
// 

// makes an ident (string) from a variable's name - in CLAIRE4 we got rid of .naming option
c_string(c:js_producer, self:Variable) : string
 -> (//[5]  WARNING : unsafe call to c_string(Var) for ~S, should be ident // self,
     Core/print_in_string(), ident(c,self.mClaire/pname), Core/end_of_string())

c_string(c:js_producer, self:symbol) : string
 -> (Core/print_in_string(), ident(c,self), Core/end_of_string())

// print a symbol for a variable
// two issues : replace with a dictionary some day (CLAIRE4) + why does c_string exist ?
// notice that ident should only exist for <strings> that will exist directly in Go code 
[ident(c:js_producer,v:Variable) : void
 ->  let s := v.mClaire/pname, n := get(c.bad_names, s) in
       (if (n = 0) c_princ(string!(s))
        else c_princ@symbol(c.good_names[n])) ]


// print a symbol for the structure definition  => use c_princ to get rid of special chars
// [should be refactored / define at producer level]
[ident(c:js_producer,s:symbol) : void
 -> let n := get(c.bad_names, s) in
       (if (n = 0) c_princ(string!(s))
        else c_princ@symbol(c.good_names[n])) ]


// we only support native vars
[globalVar(c:js_producer,x:global_variable) : void
  ->  if (module!(x.name) != defined(x.name)) 
            error("global variable ~S is not diet (local)",x)
      else  thing_ident(x) ]
     
// simpler in JavaScript : forget the module 
[js_class(self:class) : void
  -> let m := defined(self.name) in                // where the name is defined
       (if (m = Kernel) princ("kernel.Claire"),
        c_princ(capitalize(string!(self.name)))) ]


// how a named object is designated in Javascript 
// non local symbols are imported from "kernel"
[js_ident(s:symbol) 
  -> let m := defined(s) in
       (if (m != PRODUCER.current) princ("kernel."),
        js_var(s)) ]

 // this produced the C_s identifier which are go global variables, 
 // all compiler code should use this (get rid of C_ in code)
 [js_var(s:symbol) : void 
     -> princ("C_"),
        c_princ(string!(s)) ] 

// *******************************************************************
// *       Part 2: utilities for file generation                     *
// *******************************************************************

// variables are untyped in JavaScript, which makes the code generation much simpler

// is a property diet ? specific for a call with n args
[diet?(p:property) : boolean
  -> forall(r in p.restrictions |
          (domain!(r) % class) & 
            (module!(r) != PRODUCER.current |
              forall(r2 in p.restrictions | not(jsConflict?(r,r2)))))]

[jsConflict?(r1:restriction,r2:restriction) : boolean
    -> r2 != r1 & module!(r2) = module!(r1) & domain!(r2) = domain!(r1) ]

[nodiet(p:property) : any
  -> some(r in p.restrictions |
            ( module!(r) = PRODUCER.current  &
              domain!(r) % class) & 
              exists(r2 in p.restrictions | 
                   module!(r2) = module!(r) & r2 != r & domain!(r2) = domain!(r)))  ]

// this is simpler than go -> we want a method uniquely defined by its domain! and the arity
[jsMethod?(m:method) : boolean
 ->  let firstc := string!(m.selector.name)[1],
         n := length(domain(m)), c := domain!(m) in 
     ((firstc >= 'A' & firstc <= 'z' & firstc != '^') &
      c <= object  &  m.module! = defined(c.name) &        // 
      forall(m2 in m.selector.restrictions | not(jsConflict?(m,m2)))) ]
 

// debug
[jsMethod!(m:method) : any
 ->  some(m2 in m.selector.restrictions | 
            not((case m2 (method 
                          (domain!(m2) % class &
                            (m2 = m | length(m2.domain) != length(m.domain) | 
                            domain!(m2) != domain!(m))),
                        any false)))) ]       
    
// returns the name of the Javascript function
[jsFunction(p:js_producer,m:method) : string
  -> let s := m.selector in
       (if (m.module! != PRODUCER.current) printf("kernel."),
        if (length(s.restrictions) = 1) string!(s.name)
        else let fn := string!(s.name) /+ "_" /+ string!(class!(m.domain[1]).name), 
                 i := 0, im := 0, n := 0 in
           (for m2 in m.selector.restrictions 
              (i :+ 1,
               if (m = m2) im := i,
               if (class!(m2.domain[1]) = class!(m.domain[1])) n :+ 1),
            if (n = 1) fn
            else fn /+ string!(im))) ]

// this should be a code_producer method (refactoring)
[var_declaration(p:js_producer,v:string,mode:integer) : void
  -> printf("var ~I ~I", c_princ(v), (if (mode > 0) breakline())) ]

// prints a list of arguments with types / replaces typed_args_list
[jsVariables(p:js_producer,self:list) : any
 -> let prems := true in
       for x:Variable in self
         (if prems prems := false else printf(","),
          ident(p,x)) ]


// *******************************************************************
// *       Part 4: use of language dependent patterns (macros)       *
// *******************************************************************


// when we print an equality, we do not need to_CL !
// id is used to force the identifiability (use = vs equal)
[equal_exp(c:js_producer, a1:any,pos?:boolean,a2:any,id?:any) : void
 ->  if (static_type(a1) = string & static_type(a2) = string) 
        printf("(~I.Value ~I ~I.Value)",   
                j_expression(a1), sign_equal(pos?), j_expression(a2))
     else if ((id? | identifiable?(a1) | identifiable?(a2) | g_sort(a1) = float) & g_sort(a1) = g_sort(a2))
         (if (stupid_t(a1) glb stupid_t(a2) = {})
            (warn(), trace(1,"~S = ~S will fail ! [263]",a1,a2)),
         printf("(~I ~I ~I)", j_expression(a1), sign_equal(pos?), j_expression(a2)))
    else if (stupid_t(a2) = integer)
       printf("~I~I.IsInt(~I)", (if not(pos?) princ("!")), j_expression(a1), j_expression(a2))
    else printf("(Equal(~I,~I) ~I CTRUE)", j_expression(a1), j_expression(a2), sign_equal(pos?)) ]


             
// reads the member x from an expression self of expected type s
[c_member(c:js_producer,self:any,s:class,x:property) : void
 ->  printf("~I.~I", j_expression(self),  c_princ(string!(x.name)))]


// generic for bags
[bag_expression(c:js_producer,cl:class,l:list) : void
  -> if (cl = set) printf("new Set(~I)", js_args_list(l))
     else if (length(l) = 0 & cl != tuple) printf("[]")
     else  printf("[~I]", js_args_list(l)) ]



// *******************************************************************
// *       Part 5: Utilities :                                       *
// *******************************************************************

// this is an attempt to get rid of useless parenthesis without creating ambiguous situations
// bounded_expression(x,loop) adds wrapping ( ) if needed     ==     bounded expression :)
// NOTE : should be a code producer method
// we could even make this more elegant if we implement
//     expression(c:js_producer,x,s) -> j_expression(x)
[bounded_expression_js(self:any) : void
  -> case self (Assign printf("(~I)",j_expression(self)),
                integer (if (self < 0)  printf("(~I)",_jexpression(self))    // v3.2.44
                         else j_expression(self)),                           // avoid (2--2)
                float   (if (self < 0.0)  printf("(~I)",j_expression(self))
                         else j_expression(self)),
                any    j_expression(self)) ]

// atIndex : print an integer "minus one"
// NOTE : should be a code producer method
[at_index_js(x:any) : void
  -> case x (integer princ(x - 1), any (j_expression(x), princ("-1"))) ]
