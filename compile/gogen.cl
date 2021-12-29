//+-------------------------------------------------------------+
//| CLAIRE                                                      |
//| gogen.cl                                                    |
//| Copyright (C) 2020 - 2021 Yves Caseau. All Rights Reserved  |
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

// renaming philosophy:
// keyword => become allcaps and if not good enough, add _CL_
// class => add Claire

// debug
[new_block(tag:string) 
  -> if PRODUCER.debug? printf("/* ~A:~A */",tag,OPT.level),
     new_block()]

[close_block(tag:string)
  -> if PRODUCER.debug? printf("/* ~A-~A */",tag,OPT.level - 1 ),
     close_block()]

[finish_block(tag:string)
  -> if PRODUCER.debug? printf("/* ~A!~A */",tag,OPT.level - 1),
     finish_block()]

// adds a distinct ID to a variable name that may be reused
[genvar(v:string) : string
  -> PRODUCER.varsym :+ 1,
    v /+ string!(PRODUCER.varsym)]

// *******************************************************************
// *       Part 1: definition of the code producer               *
// *******************************************************************

// definition of the instance
// to do : update the reserved names progressively - note that classes do not need protection since
// ClaireX is added to X
GO_PRODUCER :: go_producer(
  Generate/open_comparators = list(<, >, >=, <=),     // do not change -> goexp uses the specific order !
  Generate/open_operators = list(+,-,*,>>),
  Generate/div_operators = list(/,mod),
  Generate/extension = ".go",
  comment = "Go",
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
 // a list of interface
 Generate/interfaces = list(integer, "int", char, "rune", string, "string", float, "float64 "),
 kernel_methods = list<any>(// empty @ set, class! @ type, copy @ set, length @ bag,
                             @ @ type,"At", array! @ list, // size @ set, empty @ bag, 
                             list! @ set, set! @ list, tuple! @ list,    // defined in Core (2nd order type)
                             list! @ tuple, /+ @ list,"Append", << @ list, "Skip")
)

// use this producer
(PRODUCER := GO_PRODUCER)

// makes an ident (string) from a variable's name - in CLAIRE4 we got rid of .naming option
c_string(c:go_producer, self:Variable) : string
 -> (//[5]  WARNING : unsafe call to c_string(Var) for ~S, should be ident // self,
     Core/print_in_string(), ident(c,self.mClaire/pname), Core/end_of_string())

c_string(c:go_producer, self:symbol) : string
 -> (Core/print_in_string(), ident(c,self), Core/end_of_string())

// print a symbol for a variable
// two issues : replace with a dictionary some day (CLAIRE4) + why does c_string exist ?
// notice that ident should only exist for <strings> that will exist directly in Go code 
[ident(c:go_producer,v:Variable) : void
 ->  let s := v.mClaire/pname, n := get(c.bad_names, s) in
       (if (n = 0) c_princ(string!(s))
        else c_princ@symbol(c.good_names[n])) ]


// print a symbol for the structure definition  => use c_princ to get rid of special chars
[ident(c:go_producer,s:symbol) : void
 -> let n := get(c.bad_names, s) in
       (if (n = 0) c_princ(string!(s))
        else c_princ@symbol(c.good_names[n])) ]

// new in claire4: printd the go identifier asociated with symbol s
// cap_ident(c,x) uses capitalization : used for Class and Method, required by Go for identifiers to be visible
// notice that we print explicitly s.module! (namespace) if not claire, to avoid c name conflicts
[cap_ident(s:symbol) : void 
   -> capitalized_ident(s,module!(s))]

// this is the capitalized ident for s in namespace m
[capitalized_ident(s:symbol, m:module) : void
  -> let n := get(PRODUCER.bad_names, s) in
       (if (n = 0) 
          (if (m != claire) 
              (c_princ(capitalize(string!(m.name))),      // for objects that are not exposed to claire we add the module name
               add_underscore(s)),                        // to disambiguate between aClass and AClass in m
           c_princ(capitalize(string!(s))))
        else c_princ@symbol(PRODUCER.good_names[n])) ]     

 // short version (we do not care about the namespace) 
[cap_short(s:symbol) : void 
    -> capitalized_ident(s,claire)]   

// CLAIRE 4 NEW ! a class name is printed with the module identifier
// go_class is the the go name ModuleClass 
// class_ident => thing_ident is the name of the global variable that contains the CLAIRE object 
[go_class(self:class) : void
  -> let m := defined(self.name) in                // where the name is defined
       (if (m = Kernel) princ("Claire")
        else if (m != PRODUCER.current) 
            (cap_ident(m.name),princ(".")),
        if (self = array | self = listargs) c_princ("List")
        else cap_ident(self.name)) ]

// small subtlety : in cast names (ToX) we ommit the "Claire" for simplicity
[cast_class(self:class) : void 
  -> let m := defined(self.name) in 
        (if (m != Kernel & m != PRODUCER.current)
            (cap_ident(m.name),princ(".")),
         princ("To"),
         add_underscore(self.name),
         if (self = listargs) princ("List")
         else cap_ident(self.name)) ]

// class_ident(c) = C_c
[class_ident(self:class) : void 
  -> symbol_ident(self.name) ]

// same : remember that a class is  not a thing
[thing_ident(self:thing) : void 
  -> symbol_ident(self.name) ]

// how a named object is designated in go (through a global variable from the package = module). 
// CLAIRE v4: No prefix needed for current or Kernel
[symbol_ident(s:symbol) 
  -> let m := defined(s) in
       (if (m != PRODUCER.current & m != Kernel & m != claire)  (cap_short(m.name), princ(".")),
        go_var(s)) ]

 // this produced the C_s identifier which are go global variables, 
 // all compiler code should use this (get rid of C_ in code)
 [go_var(s:symbol) : void -> 
   let m := module!(s) in
     (princ("C_"),
      if (m != claire) (c_princ(string!(m.name)), c_princ("_")),
      c_princ(string!(s))) ]         


// when we capitalize the name of class, we may create a conflict (list vs List)
[add_underscore(name:symbol) : void 
  -> let s := string!(name) in
       (if (integer!(s[1]) % (65 .. 90)) princ("_")) ]    

// the Go code producer uses Capitalization as a strategy for name generation
//  capitalize(s)  => capitalize the first letter + search for _, remove and capitalize next letter
//  capitalize("foo_bar") = "FooBar"
[capitalize(s:string) : string
   -> let n := length(s), i := get(s,'_') in
        (if (i = 0)
            let s2 := copy(s) in
               (s2[1] := capitalize(s[1]), 
                s2)
         else capitalize(substring(s,1,i - 1)) /+ capitalize(substring(s,i + 1,n))) ]

[capitalize(c:char) : char
   -> let i := integer!(c) in 
        (if (i % (97 .. 122)) char!(i - 32) else c)]

[capitalize(s:symbol) : string -> capitalize(string!(s))]


// v3.3 : new ! a global variable contains the native value
// range = {} for global constant
[globalVar(c:go_producer,x:global_variable) : void
  ->  thing_ident(x),
      princ(".Value") ]

// the go expression that represents a global variable, as a string (reused for Gassign)

// Five sorts in go : categories to distinguish between native, object, EID
//    x:object       x,    x,       EID(x.Id(),0)
//    x:int,float,char         x,    MakeX(x),      EID{C__C,xVAL(x)}
//    x:exception     x,  x,      EID{x,1}
//  notice that Boolean is a an object but it could be handled with a native form in the future
[type_sort(x:type) : class 
  -> let c := class!(x) in
       (if (c = float | c = integer | c = char | c = EID) c 
        else any)  ]

// sorts in go are much simpler : int, float, any or EID
[g_sort(x:any) : class
  -> type_sort(static_type(x)) ]

// access the proper values slot for a list whose member type s is determined (not any)
[valuesSlot(s:class)
  -> printf("Values~A()", (if (s = integer) "I" else if (s = float) "F" else "O")) ]


// *******************************************************************
// *       Part 2: utilities for file generation                     *
// *******************************************************************

// note that all code to produce interfaces is gone :)
// a module is simply the combination of (1) a Go package (2) a ClaireModule (3) a load function

// generate a namespace definition (Go package)
[namespace!(c:go_producer,m:module) : void
 -> printf("package ~I\n",ident(m.name)) ]

// note : we have removed module!(c:go_producer,m:module) => nothing to add to the previous line

// define a new typed variable named v (called in go_stat.cl )
// short cut : var declaration without an initialisation + a breakline
// CRAZY: go compiler gets confused with some variables not being used .. the dump forces to issue a dumb
// statement to get rid of this
// mode : 0 : normal no newline, 1 : newline, 2: special
[var_declaration(v:string,s:class,mode:integer) : void
  -> printf("var ~I ~I ~I", c_princ(v), interface!(s), (if (mode > 0) breakline())),
     if (mode = 2) printf("_ = ~I~I",c_princ(v),breakline()) ]

// ! is a semantic marker for imported
[claire/imported_function?(f:any) : boolean
  -> case f (function string!(f)[1] = '#', any false) ] 


// tells if we can compile the CLAIRE method as a go method or if we shoud use a function
// we use the go method if the class is defined in the same 
// remember that Go does not support polymorphism on parameters : we can use a method only if there is one match 
// based on first argument - howver this restriction is package based (to be checked)
// we first check that the first char of the name is a proper letter
// also methods defined with #'#foo are forced to use foo :)
[goMethod?(m:any) : boolean
 -> (if (m % PRODUCER.kernel_methods) true // force to the x.M(...) pattern
     else case m (method  
             let firstc := string!(m.selector.name)[1],
                %sig := go_signature(m), c := %sig[1] in 
                 ((firstc >= 'A' & firstc <= 'z' & firstc != '^') &
                  (c <= object | c = port | c = environment) &
                  m.module! = defined(c.name) &        // was definition in claire 3.5
                  unknown?(if_write,m.selector) &      // in claire4 : EventMethod require f-compiling
                  //  not(c <= type_expression) &    // in claire4 all type expressions are defined in Kernel
                  (unknown?(functional,m) | not(imported_function?(m.functional))) &
                  forall(m2 in m.selector.restrictions | 
                          (case m2 (method (if not(m2.domain[1] % class) false
                                            else if (m2.module! = m.module! & c ^ m2.domain[1] != {})
                                               arg_match(go_signature(m2), %sig)
                                            else true), 
                                    any true)))),
            any false)) ]

    

// useful for debug - notice that a CLAIRE method defined on a class which is NOT in the same module
// is always compiled as a function
[claire/dMethod?(m:any) 
 -> let firstc := string!(m.selector.name)[1],
       %sig := go_signature(m), c := %sig[1] in 
       (printf("char -> ~S\n", (firstc >= 'A' & firstc <= 'z' & firstc != '^')),
        printf("hierarchy -> ~S\n", (c <= object | c = port | c = environment)),
        printf("module [~S] -> ~S\n",  m.module!, m.module! = defined(c.name)),
        printf("all m -> ~S\n", forall(m2 in m.selector.restrictions | 
                          (printf("---- try m2 = ~S in ~S\n",m2,m2.module!),
                           case m2 (method (if (m2.module! = m.module! & c ^ m2.domain[1] != {})
                                               arg_match(go_signature(m2), %sig)
                                            else true), 
                                    any true)))))]

            
// same argument types for all restrictions, excluding the range (that is included in go_signature)
[arg_match(l1:list<class>,l2:list<class>) : boolean
  -> let n := length(l1) in
      (length(l2) = n & forall(i in (2 .. (n - 1)) | l1[i] = l2[i])) ]

// create the function (a name) for a method with selector p and signature l
// the name of the module where p was defined is included (until claire => public)
[function_name(p:property,l:list) : string
 -> let n := 0,        // iterate all restrictions
        m := 0,        // index in polymorphic case
        md := module!(name(p)),
        c := class!(l[1]),
        r:string := ((string!(p.name) /+ "_") /+ string!(c.name)) in  // default name p_c
      (if (p != main & md != claire)     // v3.1.04
          r := (string!(md.name) /+ "_") /+ r,
       for r in p.restrictions
          (if (c = domain!(r)) n :+ 1, if (l Optimize/=sig? r.domain) m := n),
       if (n <= 1) r else r /+ string!(m))  ]


[at(p:go_producer) : void 
  -> princ(".")]

// prints a list of arguments with types / replaces typed_args_list
[goVariables(p:go_producer,self:list) : any
 -> let prems := true in
       for x:Variable in self
         (if prems prems := false else printf(","),
          goVariable(p,x)) ]

// prints a variable declaration (inside an arg list
[goVariable(p:go_producer,v:Variable)
  -> printf("~I ~I", ident(p,v), interface!(class!(v.range))) ]


// prints the name of a method as a go method 
// Here we use the list of exceptions (kernel_methods) to force a "go method syntax" (with possibly a forced name)
// this is convenient when cross-compiling (when method move from one module/package to another)
[goMethod(m:method) 
  ->  let lm := PRODUCER.kernel_methods, i := get(lm,m) in
        (if (i > 0 & length(lm) > i & lm[i + 1] % string) c_princ(lm[i + 1])
         else c_princ(capitalize(m.selector.name))) ]

// prints the name of a function as a go function F_f
// NOTE : the link method <=> go function is not stored (the function is not known by CLAIRE)
// imported functions do not refer to the module/package
[goFunction(m:method) : void
   ->  let md := m.module! in
         (if (md != Kernel & md != PRODUCER.current & md != claire &
              (unknown?(functional,m) | not(imported_function?(m.functional))))
             (cap_short(md.name), princ(".")),  
          princ("F_"), 
          import_princ(getFunctionName(m))) ]


// specialized version for Core method
[preCore?() : void
  -> if (PRODUCER.current != Core) princ("Core.") ]         

// prints the name of the EID method that is linked by the compiler to the method
[goEIDFunction(m:method) : void
  -> let s := getFunctionName(m) in
      printf("MakeFunction~A(E_~I,~I)", length(m.domain),
               import_princ(s), print(s)) ]

// prints the function MakeFunction(...) expression
[goEIDFunctionName(m:method) : void
  -> let f := getFunctionName(m) in
      printf("E_~I", c_princ(f)) ]


// special function for self_eval of type  => added as an extra paramer of type eFunc
// cf. goexp: AddMethod -> AddEvalMethod     
[goEvalFunction(m:method) : void
  -> let c := domain!(m) in
      printf(",EVAL_~I",c_princ(c.name)) ]

// exceptions
*length_string* :: (length @ string)
*set!_list* :: (set! @ list)

// get function name
[getFunctionName(m:method) : string
  ->  if (m = *length_string*) "length_string"
      else if (m = *nth_list*) "nth_list"
      else if (m = *set!_list*) "set_I_list"
      else if (m = *stack_apply*) "CALL"
      else if (m = *super_apply*) "SUPER"
      else if (m = *belong*) "BELONG"
      else if known?(functional,m) string!(m.functional)
      else function_name(m.selector,m.domain)  ]

 // ugly : reverse engineer a compiled definition into a method
 // we need to do something better
[retreive_method(p:any,lf:any) : method
   -> case p (property 
                let m := p @ retreive_list(lf) in
                  (case m (method m,
                           any error("there is no method ~S @ ~S",p,lf))),
              any error("we have a problem to retreive ~S (not a property) at ~S",p,lf)) ]

// constrained eval in disguise : returns a type or a list of types from CLAIRE expressions
[retreive_list(x:any) : any 
  -> case x
       (type x,
        integer x,
        property x,
        // Generate/to_CL retreive_list(x.arg),
        // Generate/to_protect retreive_list(x.arg),
        global_variable (if (x.range = {}) retreive_list(x.value)
                         else error("we cannot retreive a type from a variable ~S",x)),
        List list{ retreive_list(y) | y in x.args},
        Tuple tuple!(list{ retreive_list(y) | y in x.args}),
        Call_method (if (x.arg.selector = nth &  length(x.args) = 2)
                        nth(retreive_list(x.args[1]), retreive_list(x.args[2]))
                     else if (x.arg.selector = nth &  length(x.args) = 3)
                        nth(retreive_list(x.args[1]), retreive_list(x.args[2]), retreive_list(x.args[3]))
                     else if (x.arg.selector = Core/param! &  length(x.args) = 2)
                        Core/param!(retreive_list(x.args[1]), retreive_list(x.args[2]))
                     else if (x.arg.selector = U &  length(x.args) = 2)
                        retreive_list(x.args[1]) U retreive_list(x.args[2])
                     else if (x.arg.selector = .. &  length(x.args) = 2)
                        retreive_list(x.args[1]) .. retreive_list(x.args[2])
                     else  error("we need to extend retreive_list to handle a type call: ~S",x)),
        any error("we need to extend retreive_list to handle ~S",x)) ]
    
// *******************************************************************
// *       Part 3: interface declarations                            *
// *******************************************************************

// How to declare a sort in Go. The boolean tells if we are in an external
// mode , in which case we produce the C sort. Otherwise, we use OIDs.
// THERE are 5 sorts in go : int, float, char,  any (object) and EID
// there are 7 sorts in CLAIRE : int, float, char, object, string, function, any
[interface!(self:class) : void
 ->   if (self = void) princ("void ")
      else if (self = integer) princ("int")
       else if (self = float) princ("float64")
      else if (self = char)  princ("rune")
      else if (self = EID) princ("EID")
      else  printf("*~I ", go_class(self))  ]

// general translation method: x is an expression that must be translated
// to a CLAIRE object (*ClaireX). x is known to be functional ! s is the sort for x.
[to_cl(c:go_producer,x:any,s:class) : void
 -> if (s = void) printf("Void(~I)", g_expression(x,any))
    else if (s = integer) printf("MakeInteger(~I)", g_expression(x,integer))
    else if (s = float) printf("MakeFloat(~I)", g_expression(x,float))
    else if (s = char) printf("MakeChar(~I)", g_expression(x,char))
    // else if (s = string) printf("MakeString(~I)", g_expression(x,string))  // ??? why the comment out ?
    else if (s inherit? object | s = any | s = primitive) g_expression(x,s)
    else error("[internal] to_cl for a ~S is not implemented", s) ]

/* reverse function : produce a native forme from a claire object
// quite simple with go since for object, OID is the object
[to_c(c:go_producer,x:any,s:class) : void
  -> if (x = unknown) printf("CNULL")
     // else if (x % global_variable & x.range = {} & x.value = nil)   printf("Kernel.nil")
     else if (s = integer | s = float | s = string | s = char | s = function) 
       printf("~I.Value", g_expression(x, s))
     else g_expression(x,s) ] */

// new for go: compile to an EID form (128 bit generic representation)
// s is the expected sort
[to_eid(c:go_producer,x:any,s:class) : void
 -> if (s = void) princ("EVOID")
    else if (s = integer) printf("EID{C__INT,IVAL(~I)}", g_expression(x,integer))
    else if (s = float) printf("EID{C__FLOAT,FVAL(~I)}", g_expression(x,float))
    else if (s = char) printf("EID{C__CHAR,CVAL(~I)}", g_expression(x,char))
    else if (s = string | s = function | s inherit? object) printf("EID{~I,0}", g_expression(x,any))
    else if (s = any | s = primitive) printf("~I.ToEID()",to_cl(c,x,s))
    else error("[internal] to_eid for a ~S is not implemented", s) ]

// reciprocate with an expected class e / used for variables
[from_eid(c:go_producer,x:string,e:class) : void
  -> let s := class!(e) in
      (eid_prefix(s),
       c_princ(x),
       eid_post(s))]

// reciprocate : move from EID to a sort s (if s = any, do nothing )
[eid_prefix(s:class) : void
  -> if (s = EID |  s = void) nil
     else if (s = integer) printf("INT(")
     else if (s = float) printf("FLOAT(")
     else if (s = char) printf("CHAR(")
     else if (s = any | s = primitive) printf("ANY(") 
     else if (s <= object | s = array | s = string | s = port | s = function)
        printf("~I(OBJ(",cast_class(s))               
     else if (s != any) error("what the fuck: eid prefix for ~S",s) ]

[eid_post(s:class) : void
  ->  if (s = EID | s = void) nil
      else if (s = char | s = any | s = primitive) princ(")")
      else if (s <= object |  s = array | s = string | s = port | s = function) printf("))")
      else if (s != any) princ(")") ]

 // move from an integer to a EID or Object
[integer_prefix(s:class) : void
   -> if (s = EID) princ("EID{C__INT,IVAL(")
      else if (s = any) princ("MakeInteger(") ]

// move from an integer to a EID or Object
[float_prefix(s:class) : void
   -> if (s = EID) princ("EID{C__FLOAT,FVAL(")
      else if (s = any) princ("MakeFloat(") ]

// move from an integer to a EID or Object
[char_prefix(s:class) : void
   -> if (s = EID) princ("EID{C__CHAR,CVAL(")
      else if (s = any) princ("MakeChar(") ]

// move from an integer to a EID or Object
[string_prefix(s:class) : void
   -> if (s = EID) princ("EID{")
      else if (s = any) princ("(") ]

[string_post(s:class) : void 
   -> if (s = EID) princ(".Id(),0}") 
      else if (s = any) princ(").Id()") ]

// works for integer, float, char
[native_post(s:class) : void 
   -> if (s = EID) princ(")}") 
      else if (s = any) princ(").Id()") ]

// move from ClaireId (inferred) to s (expected)
[object_prefix(inferred:class, expected:class) : void
   -> if (expected = EID) 
         (if (inferred <= object) princ("EID{"))     // objects use EID{x,0} undetermined must use toEID()
      else if (expected = inferred) nil
      else if (expected = char) princ("ToChar(")        // TODO : fix this mess : why is char special => bad name exception
      else if (expected <= primitive) (cast_class(expected),princ("("))
      else if (expected <= object) (cast_class(expected),princ("(")) ]


// s: expected
[object_post(inferred:class, s:class) : void 
   -> if (s = EID) 
         (if not(inferred <= object) princ(".ToEID()")  
          else princ(".Id(),0}"))
      else if (s = inferred) nil
      else if (s = integer | s = float | s = char) 
          (if (inferred = any) princ(").Value") else princ(".Id()).Value"))
      else if (s <= object | s <= primitive) 
         (if (inferred = any) princ(")") else princ(".Id())"))
      else if (s = any) princ(".Id()") ]

// generic version that applies to everything (s1:infered) => *_prefix(s2:expected)
[cast_prefix(s1:class, s2: class) : void 
  -> if (s1 = EID) eid_prefix(s2) 
     else if (s1 = void) nil
     else if (s1 = integer) integer_prefix(s2)
     else if (s1 = float) float_prefix(s2) 
     else if (s1 = char) char_prefix(s2) 
     else if (s1 = string) string_prefix(s2) 
     else object_prefix(s1,s2)]

// generic version that applies to everything (s1) => *_post(s2)
// s1 is the goType of the expression, s2 is the expected
[cast_post(s1:class, s2: class) : void 
  -> if (s1 = EID) eid_post(s2)
     else if (s1 = void) nil 
     else if (s1 = integer |  s1 = float | s1 = char) native_post(s2)
     else if (s1 = string) string_post(s2)
     else  object_post(s1,s2)]

    
// *******************************************************************
// *       Part 4: use of language dependent patterns (macros)       *
// *******************************************************************


// when we print an equality, we do not need to_CL !
// id is used to force the identifiability (use = vs equal)
[equal_exp(c:go_producer, a1:any,pos?:boolean,a2:any,id?:any) : void
 -> // if false (a1 % to_CL & a2 % to_CL &
   //           osort(a1.set_arg) = osort(a2.set_arg) &
        //  (identifiable?(a1.arg) |              // NEW:make sure id? := true or string
         // identifiable?(a2.arg) |
        //  a1.set_arg = string | a1.set_arg = float))   // or float => will generate nice form
       // equal_exp(c,a1.arg, pos?, a2.arg, true)
     if (static_type(a1) = string & static_type(a2) = string) 
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

// new: special code for char
// CLAIRE 4 : removed char_exp => g_expression(x,char) should work
[char_exp?(c:go_producer,x:any) : boolean
  -> case x (char true,
             Call_method let m := x.arg in
                           (m = *nth_1_string* |
                            (m = *nth_string* & compiler.safety >= 2)),
             any false) ]
             
// reads the member x from an expression self of expected type s
[c_member(c:go_producer,self:any,s:class,x:property) : void
 ->  printf("~I.~I", g_expression(self,s),  cap_short(x.name))]


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

// a constant can be evaluated with no cost in an IfThenElse(test,a,b)
[constant?(self: any) : boolean
  -> self % thing | self % boolean | self % Variable | self % string | self = unknown |
      self = nil | self = {} | self % global_variable ]

// short cut for variable
[go_range(v:Variable) : class
  -> class!(v.range)] 

// in claire 4, srange(m:method) is gone, replaced by signature => this is temporary method
[go_signature(m:method) : list<class>
  -> list<class>{class!(t) | t in m.domain} add class!(m.range) ]

// probably should exist elsewhere
[full_signature(m:method) : list<type>
  -> list<type>{t | t in m.domain} add m.range ]

// print a signature in a AddMethod (goexp.cl)
[signature!(c:go_producer, l:list<type>) : void
 -> printf("Signature(~I)", args_list(list{c_code(x,type) | x in l}, any)) ]


// this is a specialized form for list expressions => see if Go should know if a ListObject, ListInt, ListFloat will be used versus generic List
[g_member(x:any) : class
  -> if (x % Call_method | x % Construct | x % Variable | x % Call_slot | x % Cast | x % global_variable) 
        // g_sort(member(c_type(x))) //  is too  strong
        let t1 := (c_type(x) @ of) in
          (if unique?(t1) the(t1) else any)
    else any ]

// this is a way to access the low-level native slices (for list and sets)
[cast_Values(sbag:class,gmem:class) : void 
   ->  let short := (if (gmem = integer) "I" else if (gmem = float) "F" else "O") in
          printf(".Values~A()",short) ]

// this method does nothing. It used to check if a name could create a naming conflict.
// we keep it until we have tested that it is safe to remove it
// we could use a stack of names that have been used (reset for each method)
[check_var(self:string) : string  
 -> self /+ string!(OPT.level)]

build_Variable(s:string,t:any) : Variable
 -> Variable!(symbol!(s), 0, t)

// use a variable v with inferred type when expected : add the casts
[use_variable(v:string,expected:class,inferred:class) : void
  -> cast_prefix(inferred,expected),
     c_princ(v),
     cast_post(inferred,expected)]

// a clean expression is both a functional expression and one that does not throw an error
[g_clean(x:any) : boolean 
 -> g_func(x) & not(g_throw(x)) ]
        
// a simple func expression that should not be left in go code
[simple_func?(x:any) : boolean 
  -> if (g_clean(x) & c_type(x) != void) true
     else false ]
