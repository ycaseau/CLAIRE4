//+-------------------------------------------------------------+
//| CLAIRE                                                      |
//| jsexp.cl                                                    |
//| Copyright (C) 2023 - 2024 Yves Caseau. All Rights Reserved  |
//+-------------------------------------------------------------+


// ---------------------------------------------------------------------
// Compiling is based upon three methods:
//  - j_func tests if the CLAIRE form can be represented by a C/ expression.
//  - j_expression transforms it into an equivalent go expression.
//    otherwise,
//  - j_statement takes also a variable as an argument, and transforms a CLAIRE
//    expression into a C statement that assigns the value of the expression
//    into the variable;
//
// A special case occurs when the expression represent a boolean value and is
// functional, we can use b_expression that returns a C boolean
// ---------------------------------------------------------------------

// *********************************************************************
// *  Contents                                                         *
// *  Part 1: g_func & expression for objects                          *
// *  Part 2: expression for messages                                  *
// *  Part 3: the inline coding of function calls                      *
// *  Part 4: expression for structures                                *
// *  Part 5: boolean optimization                                     *
// *********************************************************************


//**********************************************************************
//*          Part 1: j_func & expression for objects                   *
//**********************************************************************

// this methods tells if a CLAIRE instruction can be compiled as an expression,as opposed to a statement.
// simpler for Javascript because of exception handling
[j_func(self:any) : boolean
 -> case self
      (bag forall( x in self | j_func(x)),
       Construct (if (self % (Set U List U Tuple))
                    (length(self.args) < 15 &
                     forall( x in self.args | j_func(x)))
                  else false),
       If (j_func(self.test) & j_func(get(arg, self)) & j_func(get(other, self))),
       And j_func(self.args),
       Or j_func(self.args),
       Call j_func(self.args),
       Super j_func(self.args),
       Call_method (self.arg.selector != mClaire/update & j_func(self.args)),   
       Call_slot j_func(get(arg, self)),       // TODO  : refuse when self.test !
       Call_table j_func(get(arg, self)),
       Call_array j_func(get(arg, self)),
       // note that Assign, Gassign and Update are statements in go
       Generate/Cast j_func(self.arg),
       Generate/C_cast j_func(self.arg),
       any (self % thing | self % integer | self % string | self % char | self % lambda |
            self % float | self % Variable | self % global_variable |
            self % function | self % symbol | self = unknown | self % method |
            self % boolean | self % class | self % environment)) ]

// manages unknown + catch-all 
[j_expression(self:any) : void
  -> if (self != unknown) 
        error("/!\\ design error: j_expression(~S: ~S) unknown, not in Diet Claire",self,owner(self))
     else princ("null")]

// Things are represented by global variables 
[j_expression(self:thing) : void 
  ->  js_ident(self.name)]

// note that there are two kinds of modules
//    - packages (when m.made_of != nil)  -> defined in their first members (iClaire in Language)
//    - node modules (abstractions) => need to be attached to packages
[j_expression(self:module) : void 
  ->  error("explicit mention of module ~S is not Diet CLAIRE",self)]

// A class is similar to a thing
[j_expression(self:class) : void 
  ->  js_ident(self.name) ]

// A named object is designed by a C identifier !
[j_expression(self:boolean) : void 
  ->  if self princ("true") else princ("false") ]

// Primitive types rely on the producer to generate code that uses their specific implementation
// this is done on purpose: supports the customization through another producer
[j_expression(self:integer) : void 
  ->  princ(self)]

j_expression(self:float) : void 
  -> princ(self)
    
j_expression(self:char) : void 
  -> print(self)

// strings are primitive objects, same as function
j_expression(self:string) : void 
  -> print(self)

// symbols are forced to be strings (modules are ignored in Diet Claire)
j_expression(self:symbol) : void 
  -> print(string!(self))

// we need an environment object to get verbose()
[j_expression(self:environment) : void 
   -> princ("kernel.ClEnv")]

// lexical variables are represented by JS variables
[j_expression(self:Variable) : void 
  ->  ident(PRODUCER,self)]

// global_variables are translated to JavaScript variables
[j_expression(self:global_variable) : void 
  ->  js_ident(self.name) ]

// builds a set
j_expression(self:Set) : void 
  -> bag_expression(PRODUCER,set,self.args)

j_expression(self:set) : void
 -> bag_expression(PRODUCER,set,list!(self))

// tuples are list
j_expression(self:Tuple) : void 
  -> bag_expression(PRODUCER,list,self.args)

j_expression(self:tuple) : void 
  -> bag_expression(PRODUCER,list,list!(self))

// builds a list
j_expression(self:List) : void 
  -> bag_expression(PRODUCER,list,self.args)

j_expression(self:list) : void 
  -> bag_expression(PRODUCER,list,self)


//**********************************************************************
//*          Part 2: expression for messages                         *
//**********************************************************************

// message compiling is tricky in go : Calls produce EID but for inline, Call_method produce native forms

// calls are expected to be proper calls -> use Javascript method
// however, HERE WE FIX A BUG = listarg should be handled by the optimiser !
[j_expression(self:Call) : void  
 -> let p := self.selector, a1 := self.args[1] in 
     (if (p = Compile/anyObject!)
         (if (a1 = general_error) printf("kernel.MakeError(~I,~I)",
                   j_expression(self.args[2]), j_expression(self.args[3]))
          else error("fast instantiation ~S is not diet",self))
      else if (p = Compile/object!)
          printf("(new ~I(~S)).Is(~S)",js_class(self.args[2]),string!(a1),j_expression(a1))
      else if exists(m in p.restrictions | domain!(m) = listargs)
          printf("~I([~I])", c_princ(string!(p.name)), js_args_list(self.args))
      else if (p = float!) j_expression(a1)
      else if (p % PRODUCER.Generate/open_operators |
               p % PRODUCER.Generate/div_operators) 
           printf("(~I~A~I)", 
                  bounded_expression_js(a1), 
                  (if (p = mod) "%" else string!(p.name)), 
                  bounded_expression_js(self.args[2]))
      else if (p = nth) 
          printf("~I[~I-1]",j_expression(a1),j_expression(self.args[2]))
      else if not(diet?(p)) error("Non diet call ~S",self)
      else let l := self.args, n := 1, nl? := length(l) > 4 in
        (if nl? OPT.level :+ 1,
         j_expression(l[1]),
         printf(".~I(",c_princ(capitalize(p.name))),
         for n in (2 .. length(l))
            (if (n > 2) (princ(","),if nl? breakline()),
            j_expression(l[n])),
         princ(")"),
         if nl? OPT.level :- 1)) ]

// the other cases will be taken care in the optimization part
j_expression(self:Call_method1) : void  -> inline_exp(PRODUCER,self)
j_expression(self:Call_method2) : void -> inline_exp(PRODUCER,self)
j_expression(self:Call_method) : void -> inline_exp(PRODUCER,self)


// produces a list of C expressions, separated by commas
[js_args_list(self:list) : void
 -> let %first := true, bk? := (length(self) > 3) in
       (if bk? OPT.level :+ 1,
        for x in self
          (if %first (j_expression(x), %first := false)
           else printf(",~I~I", (if bk? breakline()),
                       j_expression(x))),
        if bk? OPT.level :- 1) ]    


 // CLAIRE4 : get rid of fast dispatch (fcall + dispatcher)

// Super is like a call
[j_expression(self:Super) : void
   -> error("super calls such as ~S are not supported in Diet CLAIRE",self)]


// *******************************************************************
// *       Part 3: the inline coding of function calls               *
// *******************************************************************

// CLAIRE4 Note : all inline optimization assume than can_throw?(m) = false


// these methods are important since they contain the open-coding optimisations. Some of the method calls are be replaced
// directly by  expressions. We always expect the native form (the sort s is passed as a parameter)
// functions with one argument
// note that we need the *_prefix / *_post 
[inline_exp(c:js_producer,self:Call_method1) : void
 -> let m := self.arg, p := m.selector, a1 := car(self.args), dm := domain!(m) in
       (if (p = - & ( dm = integer | dm = float))
           printf("(-~I)", bounded_expression_js(a1))
        else if (p = float!) bounded_expression_js(a1)
        else if (p = owner & Compile/designated?(a1))
           (if (c_type(a1) <= object) printf("~I.isa", j_expression(a1))
            else printf("kernel.owner_any(~I)", j_expression(a1)))
        else if (m = (length @ list) |  m = Generate/*length_array*) // */
           printf("~I.length", j_expression(a1))
        else if (p = integer! & domain!(m) = char & Compile/designated?(a1)) j_expression(a1)
        else if (m = Generate/*princ_string*)   // */
           printf("kernel.PRINC(~S)", a1)
        else if (m = Generate/*copy_list*)    // */
           printf("Array.from(~I)", j_expression(a1))
        else if (m = *not* & static_type(a1) <= boolean)     // v3.2.24 for Ali :-)
           printf("!~I", j_expression(a1))
        else if (m = *new_class1* & a1 % class)
           printf("(new ~I()).Is(~I)", js_class(a1),j_expression(a1))
        else print_external_call(c, self)) ]

// ===  functions with two arguments ===
[inline_exp(c:js_producer,self:Call_method2) : void
 -> let m := self.arg, p := m.selector, a1 := self.args[1], a2 := self.args[2] in
      ( if (p = class! & a1 % symbol)
           printf("~I = new kernel.ClaireClass(~S,~I,~I)", js_ident(a1),
                  string!(a1),                         // name  
                  j_expression(a2),                    // super class
                  j_expression(value(a1).open = 3))    // keep instances
        else if ( (m.domain[1] = m.domain[2]) &  
                  (p % c.Generate/open_operators |
                   (p % c.Generate/div_operators & (case a2 (integer false,   // x / y works for float not for int in Javascript
                                                             float a2 != 0.0, 
                                                             any (compiler.safety >= 3))))))
           printf("(~I~A~I)", 
                  bounded_expression_js(a1), 
                  (if (p = mod) "%" else string!(p.name)), 
                  bounded_expression_js(a2))
        else if (m = *contain_list* & Compile/identifiable?(a2))
           printf("Memq(~I,~I)", j_expression(a1), j_expression(a2))
        else if (m = *contain_set*)
           printf("~I.has(~I)", j_expression(a1), j_expression(a2))
        else if (m = m_member) belong_exp(c,a1, a2,s)
        else if (m = *inherit*)
            printf("IsIn(~I,~I)", j_expression(a1), j_expression(a2))
        else if (m = *equal*)
           printf("Equal(~I,~I)", j_expression(a1), j_expression(a2))
        else if ( ((m = *nth_list* | m = *nth_tuple*) & compiler.safety >= 3) |
                  (m = *nth_1_list* | m = *nth_1_tuple* | m = *nth_1_array* ) |
                  (m.selector = mClaire/nth_object))    // special case where we know the support (s1) of list a1
           printf("~I[~I]",  j_expression(a1),  at_index_js(a2))
        else if (((m = *nth_list* | m = *nth_tuple*) & compiler.safety >= 3) | 
                 (m = *nth_1_list* |  m = *nth_1_tuple* | m = *nth_1_array* ))     // use the .At method
             printf("at_bag(~I)", j_expression(a1), at_index_js(a2))                       // new in v4.0.6 
        else if (p = add! & domain!(m) <= list)
           printf("kernel.add_list(~I,~I)", j_expression(a1),j_expression(a2))
        else if (p = add! & domain!(m) <= set)
           printf("kernel.add_set(~I,~I)", j_expression(a1),j_expression(a2))
        else if (m = *nth_1_string* | (m = *nth_string* & compiler.safety >= 3))
           printf("nth_string(~I,~I)", j_expression(a1), j_expression(a2))
        else if (m.selector = identical?)
           printf("~IMakeBoolean(~I)~I", cast_prefix(boolean,s), b_expression(c,self, true), cast_post(boolean,s))
        else if (p = inlineok?)                        // define a macro method through an expression
              nil    // ignored 
        // this assumes no name conflict at compile time => not safe 
        else if (m = *new_class2* & a1 % class & compiler.safety >= 2)
             printf("(new ~I(~I)).Is(~I)", js_class(a1), j_expression(a2),j_expression(a1))
        else print_external_call(c, self)) ]
       

// === functions with three arguments or more
[inline_exp(c:js_producer,self:Call_method) : void
 -> let m := self.arg, a1 := self.args[1], a2 := self.args[2], a3 := self.args[3] in
       (if (m = *nth=_list* & compiler.safety >= 3)
           printf("~I[~I]=~I", j_expression(a1), at_index_js(a2), j_expression(a3))
        else if (m = *nth_put_list* | m = *nth_put_array* | (compiler.safety >= 3 & m = *nth=_list*))
         printf("nth_put_bag(~I,~I,~I)~I", j_expression(a1), j_expression(a2), j_expression(a3))
         else if (m = *make_list* & a3  = void)               // WATCH OUT to_CL should go away
             printf("make_list_integer(~I)", j_expression(a1))
         else if (m.selector = add_method | m.selector = add_slot)           // form produced by compiler args = p,ls,range
              nil // no reflection of methods or slot in JavaScript 
         else print_external_call(c,self)) ]

// THIS IS ONE OF THE KEY PATTERNS: calls a method through its compiled function
// the arguments and the result are expected in native format
// we reuse from Go the distinction between CLAIRE methods that are compiled as methods vs those compiled as functions
[print_external_call(c:js_producer,self:Call_method) : void
 -> let m := self.arg, l := self.args, n := 1, nl? := length(l) > 4 in
     (if nl? OPT.level :+ 1,
      if jsMethod?(m)                           
        (j_expression(l[1]),
         printf(".~I(",goMethod(m)),
         for n in (2 .. length(l))
            (if (n > 2) (princ(","),if nl? breakline()),
            j_expression(l[n])))
      else 
        (printf("~I(", c_princ(jsFunction(c,m))),
         if (length(l) = 1 & domain!(m) = void)  l := nil,
         for n in (1 .. length(l))
            (if (n > 1) (princ(","), if nl? breakline()),
             j_expression(l[n]))),
      princ(")"),
      if nl? OPT.level :- 1)]


//**********************************************************************
//*          Part 4: expression for structures                       *
//**********************************************************************

// if can be represented by an expression if the two arguments are constants (evaluation does not cost)
[j_expression(self:If) : void
 -> printf("(~I ? ", b_expression(PRODUCER,self.test, true)),
    OPT.level :+ 1,
    breakline(),
    printf("~I :", j_expression(get(arg, self))),
    breakline(),
    printf("~I)", j_expression(get(other, self))),
    OPT.level :- 1 ]
 
// a conjunction is also a C expression
// note that go requires && before the line break hence the more complex code
[j_expression(self:And) : void
 -> let b := (length(self.args) > 5), n := length(self.args) in
       (princ("()"),
        for i in (1 .. n)
          let x := self.args[i] in
            (b_expression(PRODUCER,x, true),
             if (i < n) printf(" && ~I", (if b breakline()))),
        princ(")")) ]

// same thing for a disjunction
[j_expression(self:Or) : void
 -> let b := (length(self.args) > 5), n := length(self.args)  in
       (princ("("),
        for i in (1 .. n)
          let x := self.args[i] in
            (b_expression(PRODUCER,x, true),
             if (i < n) printf(" || ~I", (if b breakline()))),
        princ(")")) ]


// C_cast(x) produces a cast for go  => unclear if it is still needed
j_expression(self:Generate/C_cast) : void
 -> j_expression(self.arg)    

               
// reads a slot : more complex that it looks
// when the test is on, we produce x.p.KNOWN(p) To transform CNULL into an error 
// because slots can be native, we need the generic pre/post to convert to the proper slot
[j_expression(self:Call_slot) : void
 -> c_member(PRODUCER, self.arg, any,self.selector.selector) ]
      
// for the time being we do not support tables in DietClaire
// however a generic table with a global variable that points to a dictionary (each Object() is 
// indeed a dictionary) would work
//     var Dict = Object()     // how to declare it
//     Dict["a"] = 1           // how to use it
[j_expression(self:Call_table) : void
  -> error("tables such as ~S are not supported in Diet CLAIRE",self)]

// reads an array - remember that in CLAIRE 4, arrays are nothing but fixed size lists (with 3 sorts)
[j_expression(self:Call_array) : void
  -> printf("~I.At(~I)",j_expression(self.selector), at_index(self.arg)) ]


//**********************************************************************
//*          Part 5: the logical expression compilation                *
//**********************************************************************


// it is sad to have to duplicate the code for javascript since it is so close to go 

// default solution
[b_expression(c:js_producer,self:any,pos?:boolean) : void 
  -> if (self = true) princ("true")    // v4.0.
     else printf("(~I ~I true)", j_expression(self), sign_equal(pos?)) ]

// strange : not clear why we should see a C_cast here
[b_expression(c:js_producer,self:Compile/C_cast,pos?:boolean) : void 
  -> b_expression(c,self.arg,pos?) ]


// If is supported with IfThenElse (means that all terms will be evaluated),
[b_expression(c:js_producer,self:If,pos?:boolean) : void
 -> if self.other
        printf("(~I ? ~I : ~I)", b_expression(c,self.test, true),
                                 b_expression(c,self.arg, pos?),
                                 b_expression(c,self.other, pos?))
     else printf("(~I ~I ~I)", b_expression(c,self.test, pos?),
                 Generate/sign_or(not(pos?)), b_expression(c,self.arg, pos?)) ]


// membership
[b_expression(c:js_producer,self:Call,pos?:boolean) : void
 -> let p := self.selector in
       (if (p = %) printf("(~I ~I true)", belong_exp(c,self.args[1], self.args[2],boolean),
                          sign_equal(pos?)) 
        else printf("(~I ~I true)", j_expression(self), sign_equal(pos?)))]




// some special functions are open coded when used in a logical test
[b_expression(c:js_producer,self:Call_method1,pos?:boolean) : void
 -> let m := self.arg, a1 := self.args[1] in
       (if (m = *not*) b_expression(c,a1, not(pos?))     // v3.3.12 - was :  & a1 % to_CL
        else if (m = *known*) equal_exp(PRODUCER,a1, not(pos?), unknown, true)
        else if (m = *unknown*) equal_exp(PRODUCER,a1, pos?, unknown, true)
        else printf("(~I ~I true)", j_expression(self), sign_equal(pos?))) ]

// same thing for two arguments functions
// equal_exp is in gogen.cl
[b_expression(c:js_producer,self:Call_method2,pos?:boolean) : void
 -> let m := self.arg, p := m.selector, lop := PRODUCER.Generate/open_comparators,
        a1 := self.args[1], a2 := self.args[2] in
      (if (p = !=) equal_exp(PRODUCER,a1, not(pos?), a2, false)
       else if (p = identical?) equal_exp(PRODUCER,a1, pos?, a2, true)
       else if (p = =) equal_exp(PRODUCER,a1, pos?, a2, false)
       else if (m = m_member) printf("(~I ~I true)",belong_exp(c,a1,a2,boolean),sign_equal(pos?)) 
       else if (p % lop & domain!(m) % {float,integer})
           printf("(~I ~I ~I)", j_expression(a1),
                  (if pos? print(p)
                   else print(lop[((get(lop, p) + 1) mod 4) + 1])),  // lop = (<, >, >=, <=)
                  j_expression(a2))
        else if (m = *nth_integer*) // bit vectors  (a1 is a integer seen as a set, a2 is an integer)
           printf("(BitVectorContains(~I,~I) ~I true)", 
                      j_expression(a1), j_expression(a2), sign_equal(pos?))
        else if (p = inherit? & domain!(m) = class)
         printf("(~I.IsIn(~I) ~I true)",j_expression(a1), j_expression(a2),sign_equal(pos?))
       else if (m.range <= boolean)
          printf("(~I ~I true)", j_expression(self), sign_equal(pos?))
       else printf("(~I ~I true)", j_expression(self), sign_equal(pos?))) ]

// compile (a % ..), s is always a boolean but for EID mode
// the notOpt() test in gostat.cl ensures that the first three cases are seen as not-throw (not EID)
// however this fragment may be called to return an EID hence the global wrap with prefix/post
[belong_exp(p:js_producer,a1:any,a2:any,s:class) : void
 ->  printf("F_BELONG(~I,~I)", j_expression(a1),j_expression(a2)) ]