//+-------------------------------------------------------------+
//| CLAIRE                                                      |
//| gogen.cl                                                    |
//| Copyright (C) 2020 - 2023 Yves Caseau. All Rights Reserved  |
//| cf. copyright info in file object.cl: about()               |
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
JS_PRODUCER :: js_producer(
  Generate/open_comparators = list(<, >, >=, <=),     // do not change -> goexp uses the specific order !
  Generate/open_operators = list(+,-,*,>>),
  Generate/div_operators = list(/,mod),
  Generate/extension = ".js",
  comment = "Javascript",
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
    symbol!("Range"), symbol!("ClaireRegister"),symbol!("ClaireTemplate")),     // v3.3.46
 )

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

      


// *******************************************************************
// *       Part 2: utilities for file generation                     *
// *******************************************************************

// variables are untyped in JavaScript, which makes the code generation much simpler

// is a property diet ?
[diet?(p:property) : boolean
  -> forall(r in p.restrictions |
             (r.domain % class) & 
              not(exists(r2 in p.restrictions | r2 != r & r2.domain = r.domain)))]

      
// returns the name of the Javascript function
[jsFunction(p:js_producer,m:method) : string
  -> let s := m.selector in
       (if length(s.restrictions) = 1) string!(s.name)
        else string!(s.name) /+ "_" /+ string!(m.domain)) ]


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
                g_expression(a1,string), sign_equal(pos?), g_expression(a2,string))
     else if (char_exp?(c,a1) | char_exp?(c,a2))
        printf("(~I ~I ~I)", g_expression(a1,char), sign_equal(pos?),  g_expression(a2,char))
     else if ((id? | identifiable?(a1) | identifiable?(a2) | g_sort(a1) = float) & g_sort(a1) = g_sort(a2))
         (if (stupid_t(a1) glb stupid_t(a2) = {})
            (warn(), trace(1,"~S = ~S will fail ! [263]",a1,a2)),
         printf("(~I ~I ~I)", g_expression(a1, g_sort(a1)), sign_equal(pos?), 
                 g_expression(a2, g_sort(a1))))
    else if (stupid_t(a2) = integer)
       printf("~I~I.IsInt(~I)", (if not(pos?) princ("!")), g_expression(a1,any), g_expression(a2,integer))
    else printf("(Equal(~I,~I) ~I CTRUE)", g_expression(a1,any), g_expression(a2,any), sign_equal(pos?)) ]


             
// reads the member x from an expression self of expected type s
[c_member(c:js_producer,self:any,s:class,x:property) : void
 ->  printf("~I.~I", j_expression(self),  cap_short(x.name))]


// generic for bags
[bag_expression(c:go_producer,cl:class,l:list,t:type) : void
  -> if (length(l) = 0 & cl != tuple)
        printf("~I.Empty~I()",g_expression(c_code(t,object),type), cap_short(cl.name))
     else if (t = {} | t = void)              // constant set or tuple (constant list)
        printf("Make~A(~I)", 
                 (if (cl = set) "ConstantSet"  else if (cl = list) "ConstantList" else "Tuple"), args_list(l, any))
     else if (cl = list & t = integer)
        printf("MakeListInteger(~I)",args_list(l, integer))
     else  printf("Make~I(~I,~I)", cap_short(cl.name),g_expression(c_code(t,object),type),  args_list(l, any)) ]



// *******************************************************************
// *       Part 5: Utilities :                                       *
// *******************************************************************

// this is an attempt to get rid of useless parenthesis without creating ambiguous situations
// bounded_expression(x,loop) adds wrapping ( ) if needed     ==     bounded expression :)
// NOTE : should be a code producer method
// we could even make this more elegant if we implement
//     expression(c:js_producer,x,s) -> j_expression(x)
[bounded_expression_js(self:any,s:class) : void
  -> case self (Assign printf("(~I)",j_expression(self)),
                integer (if (self < 0)  printf("(~I)",_jexpression(self))    // v3.2.44
                         else j_expression(self)),                           // avoid (2--2)
                float   (if (self < 0.0)  printf("(~I)",j_expression(self))
                         else j_expression(self)),
                any    j_expression(self)) ]

// atIndex : print an integer "minus one"
// NOTE : should be a code producer method
[at_index_js(x:any) : void
  -> case x (integer princ(x - 1), any (g_expression(x, integer), princ("-1"))) ]
