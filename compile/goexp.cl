//+-------------------------------------------------------------+
//| CLAIRE                                                      |
//| goexp.cl                                                    |
//| Copyright (C) 2020 - 2021 Yves Caseau. All Rights Reserved  |
//| cf. copyright info in file object.cl: about()               |
//+-------------------------------------------------------------+


// ---------------------------------------------------------------------
// Compiling is based upon three methods:
//  - g_func? tests if the CLAIRE form can be represented by a C/ expression.
//    In this case,
//  - g_expression transforms it into an equivalent go expression.
//    otherwise,
//  - gstatement takes also a variable as an argument, and transforms a CLAIRE
//    expression into a C statement that assigns the value of the expression
//    into the variable;
//
// A special case occurs when the expression represent a boolean value and is
// functional, we can use bool_exp that returns a C boolean
// ---------------------------------------------------------------------

// *********************************************************************
// *  Contents                                                         *
// *  Part 1: g_func & expression for objects                          *
// *  Part 2: expression for messages                                  *
// *  Part 3: the inline coding of function calls                      *
// *  Part 4: expression for structures                                *
// *  Part 5: boolean optimization                                     *
// *********************************************************************

// g_expression(x:any,s:class) produces a go expression based on expected go type
//     s = EID                            => produce an EID
//     s = any, object, c                 => produces a *ClaireAny  representation (default case)
//     s = integer, char, float, string   => produced a native representation

//**********************************************************************
//*          Part 1: g_func & expression for objects                   *
//**********************************************************************

// this methods tells if a CLAIRE instruction can be compiled as an expression,as opposed to a statement.
// CHANGE in CLAIRE 4 : everything that may throw an exception needs a statement (because of go limitation)
// HOWEVER : if a call produces the possible error, it should simply be compiled in EID mode
[g_func(self:any) : boolean
 -> case self
      (bag forall( x in self | g_func(x)),
       Construct (if (self % (Set U List U Tuple))
                    (length(self.args) < 15 &
                     forall( x in self.args | g_func(x)))
                  else false),
       If (g_func(self.test) & constant?(get(arg, self)) & constant?(get(other, self))),
       And g_func(self.args),
       Or g_func(self.args),
       Call (g_func(self.args) & self.selector != object! & forall(x in self.args | not(g_throw(x)))),    // v4: assignment are not expressions in Go
       Super g_func(self.args),
       Call_method (g_func(self.args) & (self.arg = m_unsafe | forall(x in self.args | not(g_throw(x))))),   // new in v4
       // Call_method (g_func(self.args) & (self.arg != *close_exception*) & not(g_throw(self))),   // new in v4
       Call_slot g_func(get(arg, self)),       // TODO  : refuse when selt.test !
       Call_table g_func(get(arg, self)),
       Call_array g_func(get(arg, self)),
       // note that Assign, Gassign and Update are statements in go
       Generate/Cast g_func(self.arg),
       Generate/C_cast g_func(self.arg),
       // Generate/to_C g_func(self.arg),
       // Generate/to_CL g_func(self.arg),
       any (self % thing | self % integer | self % string | self % char | self % lambda |
            self % float | self % Variable | self % global_variable |
            self % function | self % symbol | self = unknown | self % method |
            self % boolean | self % class | self % environment)) ]

// manages unknown + catch-all 
[g_expression(self:any,s:class) : void
  -> if (self != unknown) error("/!\\ design error: g_expression(~S: ~S) unknown",self,owner(self))
     else if (s = EID) to_eid(PRODUCER,self,object)
     else if (s = any) princ("CNULL")
     else printf("~ICNULL~I",object_prefix(any,s), object_post(any,s)) ]

// Things are represented by global variables in the associated go package
[g_expression(self:thing, s:class) : void 
  ->  if (s = EID) to_eid(PRODUCER,self,object) 
      else printf("~I~I~I", object_prefix(owner(self),s), 
                  thing_ident(self), 
                  object_post(owner(self),s)) ]

// note that there are two kinds of modules
//    - packages (when m.made_of != nil)  -> defined in their first members (iClaire in Language)
//    - node modules (abstractions) => need to be attached to packages
[g_expression(self:module, s:class) : void 
  ->  if (s = EID) to_eid(PRODUCER,self,object) 
      else printf("~I~I~I", object_prefix(owner(self),s), 
                  (if (self = PRODUCER.current) princ("It")
                   else if (self = Kernel) princ("C_Kernel")
                   else if (self.made_of = nil) 
                    let m := get_made(self) in
                       (if (m != Kernel & m != PRODUCER.current) (ident(m.name),princ(".")),
                        go_var(self.name))
                   else (ident(self.name), princ(".It"))),
                  object_post(owner(self),s)) ]

// A class is similar to a thing
[g_expression(self:class, s:class) : void 
  ->  if (s = EID) to_eid(PRODUCER,self,object) 
      // else if (self = bag) g_expression(list,s)           // ugly ! now that bag exist, let's use it
      else printf("~I~I~I", object_prefix(class,s),class_ident(self), object_post(class,s)) ]

// A named object is designed by a C identifier !
[g_expression(self:boolean, s:class) : void 
  ->  if (s = EID) to_eid(PRODUCER,self,object) 
      else printf("~I~A~I", object_prefix(boolean,s),(if self "CTRUE" else "CFALSE"),object_post(boolean,s)) ]

// Primitive types rely on the producer to generate code that uses their specific implementation
// this is done on purpose: supports the customization through another producer
[g_expression(self:integer,s:class) : void 
  ->  if (s = EID) to_eid(PRODUCER,self,integer)
      else if (s = integer)  princ(self)
      else (object_prefix(integer,s),
            to_cl(PRODUCER,self,integer),
            object_post(integer,s))]

g_expression(self:float,s:class) : void 
  -> (if (s = EID) to_eid(PRODUCER,self,float)
      else if (s = float) princ(self)
      else  (object_prefix(float,s),
             to_cl(PRODUCER,self,float),
             object_post(float,s)))
    
g_expression(self:char,s:class) : void 
  -> (if (s = EID) to_eid(PRODUCER,self,char)
      else if (s = char) print(self)
      else (object_prefix(char,s),
            to_cl(PRODUCER,self,char),
            object_post(char,s)))

// strings are primitive objects, same as function
g_expression(self:string,s:class) : void 
  -> (if (s = EID) to_eid(PRODUCER,self,string)
      else printf("~IMakeString(~S)~I", object_prefix(string,s),self,object_post(string,s)))

// symboles are primitive objects, same as function
g_expression(self:symbol,s:class) : void 
  -> (if (s = EID) to_eid(PRODUCER,self,object)
      else printf("~IMakeSymbol(~S,~I)~I", object_prefix(object,s),
                  string!(self),
                  g_expression(module!(self),module),
                  object_post(object,s)))

[g_expression(self:environment,s:class) : void 
   -> if (s = EID) to_eid(PRODUCER,self,object) 
      else printf("~IClEnv~I",object_prefix(environment,s),object_post(environment,s)) ]

[g_expression(self:function,s:class) : void 
   -> if (s = EID) to_eid(PRODUCER,self,object) 
      else printf("~IMakeFunction~A(E_~I,~I)~I",object_prefix(function,s),
                     Kernel/arity(self),                           // will be a call to arity(f)
                     c_princ(self), 
                     print(string!(self)),
                     object_post(function,s)) ]

// lexical variables are represented by C variables
// notice that we may need native to object conversion
[g_expression(self:Variable,s:class) : void 
  ->  //[5] g_expression(~S:Variable,~S) // self,s,
      let s2 := class!(range(self)) in
        (if (s = EID) 
            (if (s2 = EID) ident(PRODUCER,self)  // v:EID pragma was used :)
             else to_eid(PRODUCER,self,s2))
         else         // go_form : return native typ
            (cast_prefix(s2,s),
             ident(PRODUCER,self),
             cast_post(s2,s)))  ]

// global_variables are CLAIRE objects
// v4.0.4 : handle optimized variables (nativeVarG)
[g_expression(self:global_variable,s:class) : void 
  ->  if (s = EID) to_eid(PRODUCER,self,object) 
      else if (self.range = {} & (self.value % integer | self.value % float | self.value = nil))
          g_expression(self.value,s)            // global constant inlining
      else let s2 := (if nativeVar?(self) getRange(self) else any) in
           (cast_prefix(s2,s),
            globalVar(PRODUCER,self),
            cast_post(s2,s)) ]

// builds a set
g_expression(self:Set,s:class) : void 
  -> (if (s = EID) to_eid(PRODUCER,self,object) 
      else (cast_prefix(set,s),
            bag_expression(PRODUCER,set,self.args,(if known?(of,self) self.of else {})),
            cast_post(set,s)))

g_expression(self:set,s:class) : void
 -> (if (s = EID) to_eid(PRODUCER,self,object) 
     else if (size(self) = 0 & of(self) = {}) 
        printf("~ICEMPTY~I",cast_prefix(set,s), cast_post(set,s))
     else (cast_prefix(set,s),
           bag_expression(PRODUCER,set,list!(self),of(self)),
           cast_post(set,s)))

// builds a tuple
g_expression(self:Tuple,s:class) : void 
  -> (if (s = EID) to_eid(PRODUCER,self,object) 
      else (cast_prefix(tuple,s),
            bag_expression(PRODUCER,tuple,self.args,{}),
            cast_post(tuple,s)))

g_expression(self:tuple,s:class) : void 
  -> (if (s = EID) to_eid(PRODUCER,self,object) 
      else (cast_prefix(tuple,s),
            bag_expression(PRODUCER,tuple,list!(self),{}),
            cast_post(tuple,s)))

// builds a list
g_expression(self:List,s:class) : void 
  -> (if (s = EID) to_eid(PRODUCER,self,object) 
      else (cast_prefix(list,s),
            bag_expression(PRODUCER,list,self.args,(if known?(of,self) self.of else {})),
            cast_post(list,s)))

g_expression(self:list,s:class) : void 
  -> (if (s = EID) to_eid(PRODUCER,self,object) 
      else if (length(self) = 0 & of(self) = {}) 
        printf("~ICNIL~I",cast_prefix(list,s), cast_post(list,s))
      else (cast_prefix(list,s),
            bag_expression(PRODUCER,list,self,of(self)),
            cast_post(list,s)))

// new in CLAIRE 4 !! compilation of lambda is OK but requires the reader (similar to macros)
g_expression(self:lambda,s:class) : void
 -> (Optimize/legal?(Reader,self),
     printf("~ICore.F_read_lambda_string(MakeString(\"lambda[(~I),~S]\"))~I",
          eid_prefix(s),
          Language/ppvariable(self.vars), 
          self.body,
          eid_post(s)))

//**********************************************************************
//*          Part 2: expression for messages                         *
//**********************************************************************

// message compiling is tricky in go : Calls produce EID but for inline, Call_method produce native forms

// calls are expected to produce an EID
g_expression(self:Call,s:class) : void  -> inline_exp(PRODUCER,self,s)

// the other cases will be taken care in the optimization part
g_expression(self:Call_method1,s:class) : void  -> inline_exp(PRODUCER,self,s)
g_expression(self:Call_method2,s:class) : void -> inline_exp(PRODUCER,self,s)
g_expression(self:Call_method,s:class) : void -> inline_exp(PRODUCER,self,s)

// ---------------------------------------- dynamic call -------------------------------------------------------------------


// new in 3.0 : really low level method are virtual and only rely on inline compiling
// note the *_prefix(s) ... *_postfix(s) that add a conversion from * to exprected type s
// WARNING : we can use assignment (x = y) only when s = void (we do not care for the result)
[inline_exp(c:go_producer,self:Call,s:class) : void
 -> let  p := self.selector, a1 := car(self.args), n := length(self.args) in
       (if (p = mClaire/get_stack) 
           printf("~IClEnv.EvalStack[~I]~I", eid_prefix(s),g_expression(a1, integer),eid_post(s))
        else if (p = safe)                    // v3.2.12 : we had forgotten this !!!!
           let y := compiler.safety in
             (compiler.safety := 1,
              g_expression(self.args[1],s),
              compiler.safety := y)
        else if (p = mClaire/base!) 
          printf("~IClEnv.Base~I", integer_prefix(s), native_post(s))
        else if (p = Core/<=t)                                     // v4. <= is inline coded for types
           printf("~I~I.Included(~I)~I", object_prefix(boolean,s),g_expression(a1, type), 
                    g_expression(self.args[2], type),object_post(boolean,s))
        else if (p = mClaire/index! & n = 1) 
          printf("~IClEnv.Index~I", integer_prefix(s), native_post(s))
        else if (p = mClaire/push! & n = 1)
          printf("ClEnv.Push(~I)", g_expression(a1, EID))
        else if (p = mClaire/put_stack)      // we should produce an error if s /= void
          (if (s != void) (warn(), trace(1,"use ~S in non void context\n",self)),
           printf("ClEnv.EvalStack[~I]=~I", g_expression(a1, integer),
                                           g_expression(self.args[2], EID)))
        else if (p = mClaire/set_base & s = void)
          printf("ClEnv.Base= ~I", g_expression(a1, integer))
        else if (p = mClaire/set_index & s = void)
          printf("ClEnv.Index= ~I", g_expression(a1, integer))
        else if (p = anyObject!)                // special "total" instantiation for exception and others
            (if (a1 = Interval) printf("~I~I.MakeInts(~I)~I", object_prefix(any,s),
                        class_ident(a1 as class),        // v3.3.16
                        args_list(cdr(self.args), integer),
                        object_post(any,s))
             else  printf("~I~I.Make(~I)~I", object_prefix(any,s),
                        class_ident(a1 as class),        // v3.3.16
                        args_list(cdr(self.args), any),
                        object_post(any,s)))
        /* this is how we compile a store(.....,p,q, .....) declaration
        else if (length(self.args) > 20)            // v3.2.54
           (if (self.selector = store)
             let l := self.args, n := length(l), m := n / 10  in
               printf("(~I)",
                       (for i in (0 .. m)
                         (printf("(*~I)(~I)", expression(store, loop),
                                 args_list(list{l[j] |
                                                j in ((i * 10 + 1) ..
                                                      (if (i = m) n else (i * 10 + 10))) },
                                                loop, true)),
                          if (i != m) princ(","))))             // v3.1.06
            else error("[216] ~S has more than 10 parameters",self)) */
        // this is needed because add_slot is changing arity from CL 3 to CL 4
        else if (p  = add_slot)
           printf("~I~I.AddSlot(~I,~I,~I)~I", cast_prefix(slot,s),
                        g_expression(a1,class),
                        g_expression(self.args[2],property),                      // property
                        g_expression(self.args[3],type)   ,                // range 
                        g_expression(self.args[4],any),            // default value
                        cast_post(slot,s))
        else printf("~I~IF_CALL(~I,ARGS(~I))~I", eid_prefix(s), preCore?(),
                      g_expression(self.selector,property),
                      args_list(self.args, EID),
                      eid_post(s))) ]

// produces a list of C expressions, separated by commas
[args_list(self:list,s:class) : void
 -> let %first := true, bk? := (length(self) > 3) in
       (if bk? OPT.level :+ 1,
        for x in self
          (if %first (g_expression(x, s), %first := false)
           else printf(",~I~I", (if bk? breakline()),
                       g_expression(x, s))),
        if bk? OPT.level :- 1) ]    


 // CLAIRE4 : get rid of fast dispatch (fcall + dispatcher)

// Super is like a call
[g_expression(self:Super,s:class) : void
   -> printf("~I~IF_SUPER(~I, ~I, ARGS(~I))~I", eid_prefix(s),preCore?(),
              g_expression(self.selector,property),
              g_expression(self.cast_to,class),
              args_list(self.args, EID),
              eid_post(s)) ]


// *******************************************************************
// *       Part 3: the inline coding of function calls               *
// *******************************************************************

// CLAIRE4 Note : all inline optimization assume than can_throw?(m) = false


// these methods are important since they contain the open-coding optimisations. Some of the method calls are be replaced
// directly by  expressions. We always expect the native form (the sort s is passed as a parameter)
// functions with one argument
// note that we need the *_prefix / *_post 
[inline_exp(c:go_producer,self:Call_method1,s:class) : void
 -> let m := self.arg, p := m.selector, a1 := car(self.args), dm := domain!(m) in
       (if (p = - & ( dm = integer | dm = float))
           printf("~I(-~I)~I", cast_prefix(dm,s),bounded_expression(a1,s), cast_post(dm,s))
        else if (p = owner & eid_provide?(a1))
           printf("~IOWNER(~I)~I", object_prefix(class,s),g_expression(a1,EID),object_post(class,s))
        else if (p = owner & Compile/designated?(a1))
           printf("~I~I.Isa~I", object_prefix(class,s),g_expression(a1,any),object_post(class,s))
        else if (p = eval)   // in CLAIRE 4, EVAL is a function : use it
           printf("~IEVAL(~I)~I", eid_prefix(s),g_expression(a1, any),eid_post(s))
        else if (m.selector = externC) princ(a1)
        else if (m = *length_bag* & Compile/designated?(a1))
           printf("~I~I.Length()~I", integer_prefix(s), g_expression(a1, list), native_post(s))
        else if (p = integer! & domain!(m) = char & Compile/designated?(a1))
           printf("~Iint(~I)~I", integer_prefix(s),g_expression(a1, char),native_post(s))
        else if (m = *of_bag* | m = *of_array*)
           printf("~I~I.Of()~I", cast_prefix(type,s),g_expression(a1, list),cast_post(type,s))
        else if (m = m_unsafe)
           (if (s = EID) g_expression(a1,EID)
            else if not(g_throw(a1)) g_expression(a1,s)
            else printf("~IANY(~I)~I", cast_prefix(any,s),g_expression(a1,EID),cast_post(any,s)))
        else if (m = *princ_string* & a1 % string)
           printf("PRINC(~S)", a1)
        else if (m = *copy_list*)
           printf("~I~I.Copy()~I", cast_prefix(list,s),g_expression(a1, list),cast_post(list,s))
      // else if (m = *copy_set*)
      //     printf("~I~I.Copy()~I", cast_prefix(list,s),g_expression(a1, set),cast_post(list,s))
       else if (m = *length_array*)
           printf("~I~I.Length()~I", integer_prefix(s),g_expression(a1, array),native_post(s))
        else if (m = *not* & static_type(a1) <= boolean)     // v3.2.24 for Ali :-)
           printf("~I~I.Not~I", object_prefix(boolean,s),g_expression(a1,boolean),object_post(boolean,s))
        // v3.3.12 - change suggested by Sylvain
        else if (m = *new_class1* & a1 % class)
             printf("~Inew(~I).Is(~I)~I",object_prefix(any,s),
                    go_class(a1), g_expression(a1,class),
                    object_post(any,s))
        else print_external_call(c, self, s)) ]

// ===  functions with two arguments ===
[inline_exp(c:go_producer,self:Call_method2,s:class) : void
 -> let m := self.arg, p := m.selector, 
        a1 := self.args[1], a2 := self.args[2], s1 := class!(c_type(a1)) in
      ( if (p = class! & a1 % symbol)
           printf("~I = MakeClass(~S,~I,~I)", symbol_ident(a1),
                  string!(a1),                         // name  
                  g_expression(a2,class),              // superclass
                  g_expression(module!(a1),module))    // <yc>  7/98  safer (was current_module)
        else if ( (m.domain[1] = m.domain[2]) &  (s1 = integer | s1 = float) & 
                  (p % c.open_operators |
                   (p % c.div_operators & (case a2 (integer a2 != 0, 
                                                    float a2 != 0.0, 
                                                    any (compiler.safety >= 3))))))
           printf("~I(~I~A~I)~I", cast_prefix(s1,s),
                     bounded_expression(a1,s1), 
                     (if (p = mod) "%" else string!(p.name)), 
                     bounded_expression(a2, s1), cast_post(s1,s))
        else if (m = *contain_list* & Compile/identifiable?(a2))
           printf("~I~I.Memq(~I)~I", object_prefix(boolean,s),g_expression(a1, list), 
                    g_expression(a2, any),object_post(boolean,s))
        else if (m = *contain_set*)
           let sm := g_member(a1) in
           printf("~I~I.Contain~I(~I)~I", 
                    object_prefix(boolean,s),
                    g_expression(a1, set), 
                    (if (sm = integer) princ("SetInteger") else if (sm = float) princ("SetFloat") else princ("_ask")),
                    g_expression(a2, (if (sm = integer) integer else if (sm = float) float else any)), 
                    object_post(boolean,s))
        else if (m.selector = externC) princ(a1)
        else if (m = m_member) belong_exp(a1, a2,s)
        else if (m = *write_value* & eid_provide?(a2))   // returns EID hence should not be converted (error possible)
           printf("~I.WriteEID(~I)",g_expression(a1,Variable),g_expression(a2,EID))
        else if (m = *inherit*)
            printf("~I~I.IsIn(~I)~I", object_prefix(boolean,s),g_expression(a1, class), 
                    g_expression(a2, class),object_post(boolean,s))
        else if (m = *equal*)
           printf("~IEqual(~I,~I)~I", object_prefix(boolean,s),g_expression(a1, any), 
                    g_expression(a2, any),object_post(boolean,s))
       else if (m = *map*)   // FRAGILE : ADD prefix/ post SOON
           printf("~I.Map_I(~I)", g_expression(a1, type), g_expression(a2, type))
       else if (m = *%t*)            // v4. % is inline coded and only works for types
           printf("~I~I.Contains(~I)~I", object_prefix(boolean,s),g_expression(a2, type), 
                    g_expression(a1, any),object_post(boolean,s))
        else if (p = Core/<=t | m = *included*)                                     // v4. <= is inline coded for types
           printf("~I~I.Included(~I)~I", object_prefix(boolean,s),g_expression(a1, type), 
                    g_expression(a2, type),object_post(boolean,s))
        else if (((((m = *nth_list* | m = *nth_tuple*) & compiler.safety >= 3) |
                 m = *nth_1_list* | m = *nth_1_tuple* | m = *nth_1_array* ) & 
                 g_member(a1) != any) |                // will not apply if support is unknown
                 (m.selector = mClaire/nth_object))    // special case where we know the support (s1) of list a1
           let s1 := (if (m.selector = mClaire/nth_object) object 
                      else type_sort(g_member(a1))) in   // object or integer or float => type for values in list(a1); known by go compiler so no cast is necessary  
              printf("~I~I.~I[~I-1]~I", cast_prefix(s1,s),
                     g_expression(a1,list), 
                     valuesSlot(g_member(a1)),
                     g_expression(a2, integer), 
                     cast_post(s1,s))
        else if (((m = *nth_list* | m = *nth_tuple*) & compiler.safety >= 3) | 
                 (m = *nth_1_list* |  m = *nth_1_tuple* | m = *nth_1_array* ))     // use the .At method
             printf("~I~I.At(~I-1)~I", cast_prefix(any,s),
                     g_expression(a1,list), 
                     g_expression(a2, integer), 
                     cast_post(any,s))
        else if (p = add! & domain!(m) <= bag)
           let sbag := (if (domain!(m) = set) set else list), %type := g_member(a1) in
           (if (sbag = list & %type = integer & s = void)
               printf("~I.AddFastInteger(~I)",  g_expression(a1,list), g_expression(a2,integer))
            else if (sbag = set & %type = integer)
               printf("~I~I.AddSetInteger(~I)~I", cast_prefix(sbag,s), g_expression(a1,set), g_expression(a2,integer),cast_post(sbag,s))
            else printf("~I~I.AddFast(~I)~I", cast_prefix(sbag,s),
                        g_expression(a1,domain!(m)), g_expression(a2,any),
                        cast_post(sbag,s)))
        else if (m = *nth_1_string* | (m = *nth_string* & compiler.safety >= 3))
           printf("~I~I.At(~I)~I", char_prefix(s),g_expression(a1, string), 
                   g_expression(a2, integer),native_post(s))
        else if (a1 % table & p = nth & (c_type(a2) <= domain(a1 as table) | compiler.safety >= 2))
           printf("~I~IF_get_table(~I,~I)~I", cast_prefix(any,s), preCore?(),
                   g_expression(a1, table), 
                   g_expression(a2, any),cast_post(any,s))
        else if (m.selector = identical?)
           printf("~IMakeBoolean(~I)~I", cast_prefix(boolean,s), bool_exp(self, true), cast_post(boolean,s))
        else if (p = inlineok? & a2 % string)                        // define a macro method through an expression
           printf("~IF_inlineok_ask_method(~I~I,~IMakeString(~S))",preCore?(),
                   breakline(),
                   g_expression(a1,property),
                   breakline(), 
                   a2)            // inlineDef calls the reader with a sting argument
        // this assumes no name conflict at compile time => not safe 
        else if (m = *new_class2* & a1 % class & compiler.safety >= 2)
             printf("~Inew(~I).IsNamed(~I,~I)~I",object_prefix(any,s),
                    go_class(a1), g_expression(a1,class),g_expression(a2,symbol),
                    object_post(any,s))
        else print_external_call(c, self, s)) ]

// === functions with three arguments or more
[inline_exp(c:go_producer,self:Call_method,s:class) : void
 -> let m := self.arg, a1 := self.args[1], a2 := self.args[2], a3 := self.args[3] in
       (if (m = *nth=_list* & compiler.safety >= 3 & g_member(a1) != any & s = void)
           printf("~I.~I[~I-1]=~I", g_expression(a1,list),
                  valuesSlot(g_member(a1)),
                  g_expression(a2, integer), g_expression(a3, g_member(a1)),g_member(a1))
        else if (m = *nth_put_list* | m = *nth_put_array* | (compiler.safety >= 3 & m = *nth=_list*))
         printf("~I~I.NthPut(~I,~I)~I", cast_prefix(any,s), g_expression(a1,array),
                  g_expression(a2, integer), 
                  g_expression(a3, any), 
                  cast_post(any,s))
         else if (m = *make_list* & a3  = void)               // WATCH OUT to_CL should go away
             printf("~ICreateList(~I,~I)~I", cast_prefix(list,s), g_expression(a2,type),
                     g_expression(a1, integer), cast_post(list,s))
        /* else if (m.selector = store & (c_type(a1) <= list | c_type(a1) <= array) &
                 ((length(self.args) = 4 & self.args[4] = true) | length(self.args) = 3))
           printf("F_store_list(~I,~I,~I,CTRUE)",
                   g_expression(a1, list),
                   g_expression(a2, integer),
                   g_expression(a3, any)) */
        else if (m.selector = add_slot & a1 % class)
           printf("~IF_close_slot(~I.AddSlot(~I,~I,~I))", preCore?(),
                        g_expression(a1,class),
                        g_expression(a2,property),           // property
                        g_expression(a3,type),                // range 
                        g_expression(self.args[4],any))                 // default value
        else if (m.selector = add_method)           // form produced by compiler args = p,ls,range
          (if (a1 % property)
           let m := self.args[6] in         // nicely provided by the optimizer
           printf("~IF_attach_method(~I.~A(~I,~A,~I~I),MakeString(~S))",preCore?(),
                    g_expression(a1,owner(a1)),
                    (if (a1 = self_eval) "AddEvalMethod" else "AddMethod"),
                    signature!(c,full_signature(m)),
                    (if can_throw_status(m) 1 else 0),        // we encode with a status bit vector for future extensions
                    goEIDFunction(m),
                    (if (a1 = self_eval) goEvalFunction(m)),
                    FileOrigin[m])
          else printf("F_add_method_property(~I,~I,~I,~I,~I)", g_expression(a1,property),
                        g_expression(a2,list), g_expression(a3,type),
                        g_expression(self.args[4],integer), g_expression(self.args[5],function)))
        else print_external_call(c,self, s)) ]

// THIS IS ONE OF THE KEY PATTERNS: calls a method through its compiled function
// the arguments and the result are expected in native format
[print_external_call(c:go_producer,self:Call_method,s:class) : void
 -> let m := self.arg, l := self.args, n := 1, %sig := go_signature(m), sm := last(%sig) in
     (if (length(l) > 4) OPT.level :+ 1,
      if can_throw_status(m)   sm := EID,     // the function returns an EID !
      //[5] external_call ~S : s=~S and sm=~S // self,s,sm,
      cast_prefix(sm,s),
      if goMethod?(m)
        (external_casted_arg(l[1],%sig[1],0,length(l) > 4),
         printf(".~I(",goMethod(m)),
         for n in (2 .. length(l))
            external_casted_arg(l[n],%sig[n],n - 1,length(l) > 4))   // n -1 to get rid of , for arg #2 :)
      else 
        (printf("~I(", goFunction(m)),
         if (length(l) = 1 & domain!(m) = void)  l := nil,
         for n in (1 .. length(l))
            external_casted_arg(l[n],%sig[n],n,length(l) > 4)),
      princ(")"),
      if (length(l) > 4) OPT.level :- 1,
      cast_post(sm,s)) ]

 // prints the n-th arg with a possible cast if necessary since we expect the type t (hence the class class!(t))
 // n=0 is a special marker when the arg the receiver x in x.f(....)
 // in that case we can do with the static_type because of Go polymorphism
 [external_casted_arg(x:any,s:class,n:integer,nl?:boolean) : void
   ->  let st := static_type(x) in
        ( if (n > 1) (princ(","), if nl? breakline()),
          if (n = 0 & st <= s) g_expression(x,st)
          else g_expression(x,s)) ]
             

//**********************************************************************
//*          Part 4: expression for structures                       *
//**********************************************************************

// this is an attempt to get rid of useless parenthesis without creating ambuiguous situations
// bounded_expression(x,loop) adds wrapping ( ) if needed     ==     bounded expression :)
// here we assume that native is needed
[bounded_expression(self:any,s:class) : void
  -> case self (Assign printf("(~I)",g_expression(self,s)),
                // Generate/to_C  printf("(~I)",g_expression(self,s)),                     // v3.2.44
                integer (if (self < 0)  printf("(~I)",g_expression(self,s))    // v3.2.44
                         else g_expression(self,s)),                           // avoid (2--2)
                float   (if (self < 0.0)  printf("(~I)",g_expression(self,s))
                         else g_expression(self,s)),
                any    g_expression(self,s)) ]

// if can be represented by an expression if the two arguments are constants (evaluation does not cost)
[g_expression(self:If,s:class) : void
 -> (object_prefix(any,s),
     printf("IfThenElse(~I,", bool_exp(self.test, true)),
     OPT.level :+ 1,
     breakline(),
     printf("~I,", g_expression(get(arg, self),any)),
     breakline(),
     printf("~I)", g_expression(get(other, self),any)),
     object_post(any,s),
     OPT.level :- 1) ]
 
// a conjunction is also a C expression
// note that go requires && before the line break hence the more complex code
[g_expression(self:And,s:class) : void
 -> let b := (length(self.args) > 5), n := length(self.args) in
       (object_prefix(boolean,s),
        princ("MakeBoolean("),
        for i in (1 .. n)
          let x := self.args[i] in
            (bool_exp(x, true),
             if (i < n) printf(" && ~I", (if b breakline()))),
        princ(")"),
        object_post(boolean,s)) ]

// same thing for a disjunction
[g_expression(self:Or,s:class) : void
 -> let b := (length(self.args) > 5), n := length(self.args)  in
       (object_prefix(boolean,s),
        princ("MakeBoolean("),
        for i in (1 .. n)
          let x := self.args[i] in
            (bool_exp(x, true),
             if (i < n) printf(" || ~I", (if b breakline()))),
        princ(")"),
        object_post(boolean,s)) ]

// to_CL(x) produces a CLAIRE id from an external representation
// [g_expression(self:Generate/to_CL,s:class) : void
//  -> //[5] toCL -> ~S:~S // self.arg, owner(self.arg),
//    g_expression(self.arg, s)]

// to_C(x) produces an external representation from a CLAIRE id
// g_expression(self:Generate/to_C,s:class) : void
// -> g_expression(self.arg, s)

// C_cast(x) produces a cast for go  => unclear if it is still needed
g_expression(self:Generate/C_cast,s:class) : void
 -> g_expression(self.arg,s)    


               
// reads a slot : more complex that it looks
// when the test is on, we produce x.p.KNOWN(p) To transform CNULL into an error 
// because slots can be native, we need the generic pre/post to convert to the proper slot
[g_expression(self:Call_slot,s:class) : void
 -> let sc := class!(range(rootSlot(self.selector))),      // what we get from the GO (root) slot -> beware of covariant slots
        dc := static_type(self.arg),             // type of argument ... should be in domain
        s2 := (if (dc <= domain!(self.selector)) dc else class!(domain!(self.selector))),
        kt? := (known?(test,self) & self.test) in
     (if not(kt?) cast_prefix(sc,s),
      c_member(PRODUCER, self.arg, s2,self.selector.selector),
      if kt? printf(".KNOWN(~I)", g_expression(self.selector.selector, any))
      else cast_post(sc,s)) ]

// reads an (integer) table  = WARNING - this will change in the future when tables are implemented with dictionaries
// here we  assume that the table uses a list ....
[g_expression(self:Call_table,s:class) : void
  -> let a := self.selector,
         p := a.params,
         l := self.arg in
       (if (a.range <= integer) 
           (cast_prefix(integer,s),
            printf("ToList(~I.Graph).ValuesI()[~I-1]",g_expression(a, table),g_table_index(a,l)),
            cast_post(integer,s))
        else if (a.range = float)
           (cast_prefix(float,s),
            printf("ToList(~I.Graph).ValuesF()[~I-1]",g_expression(a, table),g_table_index(a,l)),
            cast_post(float,s))
        else 
         (object_prefix(any,s),
          printf("ToList(~I.Graph).At(~I-1)", g_expression(a, table),  g_table_index(a,l)),
          if self.test printf(".KNOWN(~I)", g_expression(a, any))               // assumes s is EID
          else object_post(any,s))) ]


// printf the code to access the index 
[g_table_index(a:table,l:any)
  ->  let p := a.params in 
        (case p
          (integer printf("~I - ~A", g_expression(l, integer), p),
           list (if not(l % List) error("shit with call_table ~S[~S]",a,l),
                printf("~I * ~A + ~I - ~A",                     // <yc> l is a List
                              g_expression(l.args[1], integer), p[1],
                              g_expression(l.args[2], integer), p[2])))) ]


// reads an array - remember that in CLAIRE 4, arrays are nothing but fixed size lists (with 3 sorts)
[g_expression(self:Call_array,s:class) : void
  -> let sa := type_sort(member(c_type(self.selector))),
         sm := g_member(self.selector) in
       (cast_prefix(sa,s),
        if (sm != any)
            printf("~I.~I[~I - 1]",g_expression(self.selector, list), valuesSlot(sm), g_expression(self.arg, integer))
        else printf("~I.At(~I - 1)",g_expression(self.selector, list), g_expression(self.arg, integer)),
        cast_post(sa,s)) ]


//**********************************************************************
//*          Part 5: the logical expression compilation                *
//**********************************************************************


// bool_exp(x,pos?) returns a native boolean go expression, assumes that g_func(x) !
// bool_expression(x) could be g_expression(x,boolean)
// however, boolean are not native in CLAIRE4 () to avoid conversions

// note : we drop bool_exp? and bool_exp!

// this is the boolean compiler. An automatic computation of negation is
// included. The flag pos? tells if the assertion is positive. When a
// negation occurs, we simply change the flag. At the end of compiling,
// the flag is used to generate == or != according to this method:

// generate the = or /=
sign_equal(self:boolean) : void -> (if self princ("==") else princ("!="))

// generate a conjunction/disjunction
sign_or(self:boolean) : void -> (if self princ("||") else princ("&&"))   

// default solution
[bool_exp(self:any,pos?:boolean) : void 
  -> printf("(~I ~I CTRUE)", g_expression(self, boolean), sign_equal(pos?)) ]

// strange : not clear why we should see a C_cast here
[bool_exp(self:C_cast,pos?:boolean) : void 
  -> bool_exp(self.arg,pos?) ]


// if we have a CL, we know that the self.arg is of type boolean
// [bool_exp(self:Generate/to_CL,pos?:boolean) : void
//  -> bool_exp(self.arg,pos?) ]

// If is supported with IfThenElse (means that all terms will be evaluated),
[bool_exp(self:If,pos?:boolean) : void
 -> if self.other
        printf("(~I ? ~I : ~I)", bool_exp(self.test, true),
                                 bool_exp(self.arg, pos?),
                                 bool_exp(self.other, pos?))
     else printf("(~I ~I ~I)", bool_exp(self.test, pos?),
                 Generate/sign_or(not(pos?)), bool_exp(self.arg, pos?)) ]

// for a AND, we can used the && C operation
[bool_exp(self:And,pos?:boolean) : void
 -> let l := self.args, m := length(l), n := 0, %l := OPT.level in
       (OPT.level :+ 1,
        for x in l
          (n :+ 1,
           if (n = m) bool_exp(x, pos?)
           else (printf("(~I ~I ", bool_exp(x, pos?), Generate/sign_or(not(pos?))),
                 OPT.level :+ 1,
                 breakline())),
        for x in (2 .. m) princ(")"),
        OPT.level := %l)  ]

// idem for OR: we use ||
[bool_exp(self:Or,pos?:boolean) : void
 -> let l := self.args, m := length(l), n := 0, %l := OPT.level in
       (OPT.level :+ 1,
        for x in l
          (n :+ 1,
           if (n = m) bool_exp(x, pos?)
           else (printf("(~I ~I ", bool_exp(x, pos?), Generate/sign_or(pos?)),
                 OPT.level :+ 1,
                 breakline())),
        for x in (2 .. m) princ(")"),
        OPT.level := %l) ]


// membership
[bool_exp(self:Call,pos?:boolean) : void
 -> let p := self.selector in
       (if (p = %) printf("(~I ~I CTRUE)", belong_exp(self.args[1], self.args[2],boolean), sign_equal(pos?)) 
        else bool_exp@any(self, pos?)) ]

// compile (a % ..), s is always a boolean but for EID mode
// the notOpt() test in gostat.cl ensures that the first three cases are seen as not-throw (not EID)
// however this fragment may be called to return an EID hence the global wrap with prefix/post
[belong_exp(a1:any,a2:any,s:class) : void
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


// some special functions are open coded when used in a logical test
[bool_exp(self:Call_method1,pos?:boolean) : void
 -> let m := self.arg, a1 := self.args[1] in
       (if (m = *not*) bool_exp(a1, not(pos?))     // v3.3.12 - was :  & a1 % to_CL
        else if (m = *known*) equal_exp(PRODUCER,a1, not(pos?), unknown, true)
        else if (m = *unknown*) equal_exp(PRODUCER,a1, pos?, unknown, true)
        else if (m.range <= boolean)
           printf("(~I ~I CTRUE)", g_expression(self, boolean), sign_equal(pos?))
        else bool_exp@any(self, pos?)) ]

// same thing for two arguments functions
// equal_exp is in gogen.cl
[bool_exp(self:Call_method2,pos?:boolean) : void
 -> let m := self.arg, p := m.selector, lop := PRODUCER.Generate/open_comparators,
        a1 := self.args[1], a2 := self.args[2] in
      (if (p = !=) equal_exp(PRODUCER,a1, not(pos?), a2, false)
       else if (p = identical?) equal_exp(PRODUCER,a1, pos?, a2, true)
       else if (p = =) equal_exp(PRODUCER,a1, pos?, a2, false)
       else if (m = m_member) printf("(~I ~I CTRUE)",belong_exp(a1,a2,boolean),sign_equal(pos?)) 
       else if (p % lop & domain!(m) % {float,integer})
           printf("(~I ~I ~I)", g_expression(a1, domain!(m)),
                  (if pos? print(p)
                   else print(lop[((get(lop, p) + 1) mod 4) + 1])),  // lop = (<, >, >=, <=)
                  g_expression(a2, domain!(m)))
        else if (m = *nth_integer*) // bit vectors  (a1 is a integer seen as a set, a2 is an integer)
           printf("(BitVectorContains(~I,~I) ~I CTRUE)", 
                      g_expression(a1,integer), g_expression(a2,integer), sign_equal(pos?))
        else if (p = inherit? & domain!(m) = class)
         printf("(~I.IsIn(~I) ~I CTRUE)", 
                  g_expression(a1,class), g_expression(a2,class),sign_equal(pos?))
       else if (m.range <= boolean)
          printf("(~I ~I CTRUE)", g_expression(self, boolean), sign_equal(pos?))
       else bool_exp@any(self, pos?)) ]

