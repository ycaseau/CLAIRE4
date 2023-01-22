//+-------------------------------------------------------------+
//| CLAIRE                                                      |
//| pretty.cl                                                   |
//| Copyright (C) 1994 - 2021 Yves Caseau. All Rights Reserved  |
//| cf. copyright info in file object.cl: about()               |
//+-------------------------------------------------------------+

// *********************************************************************
// *  Contents:                                                        *
// *  Part 1: unbound_symbol and variables                             *
// *  Part 2: lambdas                                                  *
// *  Part 3: close methods for lattice_set instantiation              *
// *  Part 4: Pretty printing                                          *
// *********************************************************************

// Instruction <: system_object()  : in Kernel (CLAIRE 4)
Basic_instruction <: Instruction()

no_eval(self:Instruction) : void
 -> error("[144] evaluate(~S) is not defined", owner(self))

// import => cannot work in CLAIRE4
iClaire/typing :: Kernel/typing
iClaire/index :: mClaire/index

// *********************************************************************
// *   Part 1: unbound_symbol and variables                            *
// *********************************************************************

// An unbound_symbol is created by the reader when a symbol is not bound
//
//unbound_symbol <: Basic_instruction(identifier:symbol)
self_print(self:unbound_symbol) : void
   -> printf("~A", self.name)
self_eval(self:unbound_symbol) : any
   -> (if (value(self.name) % thing)  eval(value(self.name))
       else error("[145] the symbol ~A is unbound",  self.name))

// A lexical variable is defined by a "Let" or inside a method's definition
// Lexical variables --------------------------------------------------
//
/*Variable[mClaire/pname,range] <: Basic_instruction(
     mClaire/pname:symbol,              // name of the variable
     range:type,                        //
     index:integer)                     // position in the stack */

self_print(self:Variable) : void ->
  (when s := get(mClaire/pname,self) in princ(s) else princ("V?"))

ppvariable(self:Variable) : void
 -> (if known?(range, self)
        printf("~A:~I", self.mClaire/pname, printexp(self.range, false))
     else princ(self.mClaire/pname))

ppvariable(self:list) : void
 -> (let f := true in
       for v in self
         (if f f := false
          else princ(","),
          case v (Variable ppvariable(v), any print(v))))

self_eval(self:Variable) : any -> mClaire/get_stack(mClaire/base!() + self.index)

write_value(self:Variable,val:any) : any
 -> (if (unknown?(range, self) | val % self.range)
        (mClaire/put_stack(mClaire/base!() + self.index, val), val)
     else range_error(arg = self, mClaire/cause = val, wrong = self.range))

// this is the definition of a typed variable / Vardef is a syntactic marker
// in CLAIRE 4, Vardef are transformed in Var at run time
Vardef <: Variable()

// this is strange and should be fixed  or understood
self_eval(self:Vardef) : any
  ->  let i := get(index,self) in 
         (if (i >= 0) mClaire/get_stack(mClaire/base!() + i)
          else error("[146] The variable ~S is not defined",self))

//   [self_print(self:Vardef) : any -> ppvariable(self) ]
Complex_instruction <: Instruction()
Instruction_with_var <: Complex_instruction(var:Variable)
Control_structure <: Complex_instruction()

// global_variables are defined in exception ? ---------------------------
// a global variable is a named object with a special evaluation
//
// self_eval(self:global_variable) : any -> self.value  -> moved to object.cl
write_value(self:global_variable,val:any) : any
 -> (if (val % self.range)
        (put_store(value,self,val,self.store?), val)
     else range_error(mClaire/cause = self, arg = val, wrong = self.range)) // v0.01


// same as C  (used externC("((int) EOF",integer))
EOF :: global_variable(range = char, value = char!(-1)) // v3.2.52
EOS :: global_variable(range = char, value = char!(0))

// v3.4
claire/MAX_INTEGER :: 1073741822

// *********************************************************************
// *   Part 2: CLAIRE Lambdas                                           *
// *********************************************************************

// CLAIRE lambdas are the basic functional objects, defined by a filter
// and a piece of code. Lambda is defined in the "method" file.
// applying a lambda to a list of arguments
//
apply(self:lambda,%l:list) : any
 -> (let start := mClaire/index!(),
         retour := mClaire/base!() in
       (mClaire/set_base(start),
        for %x in %l mClaire/push!(%x),
        mClaire/stack_apply(self.dimension),
        let val := eval(self.body) in
          (mClaire/set_base(retour), mClaire/set_index(start), val)))
call(self:lambda,l:listargs) : any -> apply(self, l)

// printing a lambda
self_print(self:lambda) : void
 -> printf("lambda[(~I),~I~S~I]", ppvariable(self.vars), lbreak(1),
           self.body, (pretty.index :- 1))


// map is the most famous function on a lambda
[claire/map(self:lambda,%l:bag) : any
  -> case %l (set {funcall(self,x) | x in %l},
              any list{funcall(self,x) | x in (%l as list)}) ]

// lambda! and lexical_index communicate via a global_variable, which
// however is only used in this file (and also by odefine.cl :-) ):
//
*variable_index* :: global_variable(range = integer, value = 0)

// creating a lambda from an instruction and a list of variables
lambda!(lvar:list,self:any) : lambda
 -> (*variable_index* := 0,
     for v:Variable in lvar
       (put(index, v, *variable_index*),
        put(isa, v, Variable),
        *variable_index* :+ 1),
     let corps := lexical_index(self, lvar, *variable_index*,true),
         resultat:lambda := mClaire/new!(lambda) in
       (put(vars, resultat, lvar),
        put(body, resultat, corps),
        put(dimension, resultat, *variable_index*),
        resultat))

// Give to each lexical variable its right position in the stack.
// We look for a named object or an unbound symbol to replace by a lexical
// variable.
// The number of variables is kept in the global_variable *variable_index*.
// On entry, n need not be equal to size(lvar) (see [case ...instruction]).
/*
lexical_build(self:any,lvar:list,n:integer) : any
 -> (if (self % thing | self % unbound_symbol) lexical_change(self, lvar)
     else (case self
            (Variable (if unknown?(index,self)                          // v3.1.12
                          error("[145] the symbol ~A is unbound",  self.mClaire/pname),
                       self),
             Call let s := lexical_change(self.selector, lvar) in
                    (lexical_build(self.args, lvar, n),
                     if (self.selector != s)
                        (put(selector, self, call),
                         put(args, self, s cons self.args))),
             Instruction let %type:class := self.isa in
                           (if (%type % Instruction_with_var.descendants)
                               (put(index, self.var, n),
                                n := n + 1,
                                if (n > *variable_index*)
                                   *variable_index* := n),
                            for s in %type.slots
                              let x := get(s, self) in
                                (if ((x % thing | x % unbound_symbol) &
                                     s.range = any)
                                    put(s, self, lexical_change(x, lvar))
                                 else lexical_build(x, lvar, n)),
                            if (%type = Assign & (self as Assign).var % unbound_symbol)                // CLAIRE4
                               (printf("--- in lexical_build(~S,~S,~S)\n",self,lvar,n),
                                exit(-1),
                                error("[101] ~S is not a variable but a ~S", (self as Assign).var, owner((self as Assign).var)))),             // moved from self_eval @ Assign
             list let %n := length(self) in
                   while (%n > 0)
                     (let x := (nth@list(self, %n)) in
                        (if (x % thing | x % unbound_symbol)
                            nth=@list(self, %n, lexical_change(x, lvar))
                         else lexical_build(x, lvar, n)),
                      %n :- 1),
             any nil),
           self)) */

lexical_change(self:any,lvar:list) : any
 -> (let rep:any := self,
         %name:symbol := (case self  (Variable self.mClaire/pname,
                                      any extract_symbol(self))) in
       (for x:Variable in lvar (if (x.mClaire/pname = %name) rep := x), rep))


// Give to each lexical variable its right position in the stack.
// We look for a named object or an unbound symbol to replace by a lexical variable.
// The number of variables is kept in the global_variable *variable_index*.
// On entry, n need not be equal to size(lvar) (see [case ...instruction]).
// in claire4, lexical_index replaces lexical_build with an additional variable : 
// final? = true means all Assign must contain a variable <v4.0.6>
lexical_index(self:any,lvar:list,n:integer,final?:boolean) : any
 -> (//[5] call lexical index on ~S (~S) // self,owner(self),
     if (self % thing | self % unbound_symbol) lexical_change(self, lvar)
     else (case self
            (Variable (if unknown?(index,self)                          // v3.1.12
                          error("[145] the symbol ~A is unbound",  self.mClaire/pname),
                       self),
             Call let s := lexical_change(self.selector, lvar) in
                    (lexical_index(self.args, lvar, n,final?),
                     if (self.selector != s)
                        (put(selector, self, call),
                         put(args, self, s cons self.args))),
             Instruction let %type:class := self.isa in
                           (if (%type % Instruction_with_var.descendants)
                               (put(index, self.var, n),
                                n := n + 1,
                                if (n > *variable_index*)
                                   *variable_index* := n),
                            for s in %type.slots
                              let x := get(s, self) in
                                (if ((x % thing | x % unbound_symbol) &
                                     s.range = any)
                                    put(s, self, lexical_change(x, lvar))
                                 else lexical_index(x, lvar, n, final?)),
                            if (%type = Assign & not((self as Assign).var % Variable) & final?)                // CLAIRE4
                                error("[101] ~S is not a variable but a ~S", (self as Assign).var, owner((self as Assign).var))),             // moved from self_eval @ Assign
             list let %n := length(self) in
                   while (%n > 0)
                     (let x := (nth@list(self, %n)) in
                        (if (x % thing | x % unbound_symbol)
                            nth=@list(self, %n, lexical_change(x, lvar))
                         else lexical_index(x, lvar, n, final?)),
                      %n :- 1),
             any nil),
           self))

// *******************************************************************
// *       Part 3: functions for lattice_set instantiation           *
// *******************************************************************
// close is the basic method called by an instantiation.
// Once the indexed list is built, we never call it again.
//
close(self:class) : class -> self

// Extract the symbol associated with self.
// This is useful e.g. when using read() (read@port, read@string).
//
extract_symbol(self:any) : symbol
 -> (case self
      (unbound_symbol self.name,
       thing self.name,
       class self.name,
       symbol self,
       Variable self.mClaire/pname,
       boolean (if self symbol!("true") else symbol!("nil")),
       any error("[147] a name cannot be made from ~S", self)))

// we must be sure that the selector (in a has statement or in a message)
// is a property.
//
make_a_property(self:any) : property
 -> (case self
      (global_variable make_a_property(value(self)),
       property self,
       symbol let x := value(self) in
               (case x (property make_a_property(x),
                        global_variable  make_a_property(value(x)),
                        any  let p := (mClaire/new!(property, self) as property) in
                                 (p.comment := string!(self),
                                  put(domain, p, any),
                                  put(range, p, any),
                                  p))),
       unbound_symbol make_a_property(self.name),
       any error("[148] Wrong selector: ~S, cannot make a property\n", self)))

printl :: property()

// *********************************************************************
// *  Part 4: Pretty printing                                          *
// *********************************************************************

PPC:integer := 0   // debug

// create a line break
// if the pretty mode is here ... 
//    (1) pbreak = true means that we create a new line (whatever the length)
//    (2) break = false => we generate a much too far exception
lbreak() : any
 -> (if pretty.mClaire/pprint
        (if (pretty.mClaire/pbreak)
            (princ("\n"),
             put_buffer(),    // time to output current buffer
             indent(pretty.index))
         else if (mClaire/buffer_length() > pretty.mClaire/width)  much_too_far()))

put_buffer() : any
 -> (let buffer := end_of_string() in
       (princ(buffer), print_in_string(), {}))

checkfar() : any
 -> (if (pretty.mClaire/pprint & not(pretty.mClaire/pbreak) &
         mClaire/buffer_length() > pretty.mClaire/width) much_too_far())

lbreak(n:integer) : any -> (pretty.index :+ n, lbreak())

// indentation
//
indent(limit:integer) : any
 -> (let x := mClaire/buffer_length() in while (x < limit) (princ(" "), x :+ 1))

// sets the current_level
set_level() : void
 -> (pretty.index := mClaire/buffer_length() - 1)
set_level(n:integer) : void -> (set_level(), pretty.index :+ n)

// prints a list as a box in character zone [start, finish], s is the separator (",")
// pbreak = true means that we will print step by step; false => try to add to current place
// if impossible of if pbreak = false, we will switch to printl
// the tricky part is that this method can generate a too far error
printbox(self:list,start:integer,finish:integer,s:string) : any
 -> (let i := 1,
         startline := true,
         n := length(self),
         %l := pretty.index in
       (pretty.index := start,
        if (not(pretty.mClaire/pprint) | (not(short_enough(start + 10))
             & pretty.mClaire/pbreak))
           printl(self, s)
        else if not(pretty.mClaire/pbreak) printl(self, s)        // call reccursively in no-break mode
        else while (i <= n)                                        // do our step by step
               (while (Core/buffer_length() < start) printf(" "),  // move to a new line
                let idx := Core/buffer_length() in
                  (try (pretty.mClaire/pbreak := false,           // will force to create a much too far
                        printexp(self[i], true),
                        pretty.mClaire/pbreak := true)
                   catch much_too_far (pretty.mClaire/pbreak := true,
                                       pretty.index := start),
                 if (i != n) princ(s),
                 if (Core/buffer_length() < finish)       // happy with result
                    (i :+ 1, startline := false)          // startline = false => keep adding
                 else (Core/buffer_set_length(idx),
                       if not(startline) (lbreak(), startline := true)   // i does not change, we will come back
                       else (set_level(),
                             pretty.index :+ 1,
                             printexp(self[i], true),
                             pretty.index := %l,
                             if (i != n) (princ(s), lbreak()),
                             i :+ 1)))),
        pretty.index := %l,
        unknown))

// default value of arguments
//
printbox(self:list) : any
 -> printbox(self, mClaire/buffer_length(), pretty.mClaire/width, ", ")
printbox(self:list,s:string) : any
 -> printbox(self, mClaire/buffer_length(), pretty.mClaire/width, s)

// this is a tricky method : first try to print without pretty (box) 
printl(self:list,s:string) : void
 -> (let f := true, 
         b := pretty.mClaire/pprint in
       (pretty.mClaire/pprint := false,
        try for x in self
          (if f f := false
           else princ(s),
           printexp(x, true),
           if (b & not(pretty.mClaire/pbreak) &
                   mClaire/buffer_length() > pretty.mClaire/width)
              (pretty.mClaire/pprint := b, much_too_far()))          // only if pbreak = false 
        catch system_error let x := (system.exception! as exception) in
                             (if (b & x.index = 16)                   // buffer too small
                                 (pretty.mClaire/pprint := b, much_too_far())
                              else close(x)),
       pretty.mClaire/pprint := b))

// print bounded prints a bounded expression using ( and )
[printexp(self:any,comp:boolean) : void
 ->  if ((case self
           (Call not((self.selector % operation & not(comp) &
                      length(self.args) = 2)))) |
         self % Collect | self % Select | self % Definition |
         self % Construct | self % Do | self = unknown | self % And |
         self % primitive | self % Or | self % If | self % restriction |
         self % unbound_symbol | self % Variable | not(self % Instruction)) print(self)
     else let %l := pretty.index in
            (printf("(~I~S)", set_level(1), self), 
             pretty.index := %l) ]

// pretty print is using the buffered print (into a string)
pretty_print(self:any) : void
 -> (print_in_string(),
     pretty.mClaire/pprint := true,
     pretty.mClaire/pbreak := true,
     pretty.index := 0,
     print(self),                        // will apply print twice (with break = false and break = true)
     pretty.mClaire/pprint := false,
     princ(end_of_string()))

// self_print uses the default boxing
[self_print(self:list) : void
 -> if (of(self) != {}) printf("list<~S>",of(self)) else princ("list"),
    printf("(~I)", printbox(self)) ]

[self_print(self:set) : void
  -> if (of(self) = {}) printf("{~I}", printbox(list!(self)))
     else  (printf("set<~S>",of(self)),
            printf("(~I)", printbox(list!(self)))) ]  

// to remove !
[self_print(self:tuple) : void
 -> printf("tuple(~I)", printbox(self as list)) ]

// a map_set 
[self_print(self:map_set) : void
  -> printf("map<~S,~S>", domain(self), range(self)) ]  

// a pair
[self_print(x:pair) : void
    -> printexp(x.first,false),
       princ(":"),
       printexp(x.second,false) ]

// *********************************************************************
// *  Part 5: simple type inference  (class based)                     *
// *********************************************************************

// this is a simple, self-contained, type inference method that mimicks what GO is bound to know
// it is used to check the type safety of the gerenated code in the Generate module and it is used
// in call.cl to produce OFTO (on-the-fly optimization) => see readcall
// s_type =  static type, or stupid_type  (we should remove stupid_t)
[static_type(self:any) : class
 -> case self
      (Variable class!(self.range), 
       global_variable let r := self.range in                          // was missing ! v3.0.62
                         (if r class!(r) else owner(self.value)),
       And boolean,
       Or boolean,
       environment environment,
       Call+ let p := self.selector, s := (p @ static_type(self.args[1])) in
                (case s (slot s.range, any p.range)),
       Call_slot let s := self.selector, p := s.selector in
                   (for s2 in p.mClaire/definition
                      (case s2 (slot (if (domain!(s) <= domain!(s2)) s := s2))),      // v3.2.30 C++ is really stupid :-)
                    class!(s.range)),
       Call_method let p := self.arg.selector in            // this is extreme :) we catch l<X>[i]
              (if (p = nth) static_type_nth(self.args[1]) else class!(self.arg.range)),
       Call let p := self.selector in            // this is extreme :) we catch l<X>[i]
              (if (p = nth) static_type_nth(self.args[1]) else class!(p.range)),
       Assign static_type(self.arg),
       Let static_type(self.arg),
       Do static_type(last(self.args)),
       If static_type(self.arg) meet static_type(self.other),
       Collect list,
       Image set,
       Select set,
       Lselect list,
       List list,
       Set set,
       Tuple tuple,
       Exists (if (self.other = unknown) any else boolean),
       Definition arg(self),
       Instruction any,
       any owner(self)) ]

// second order pattern for a very common case l[i] where l:list<X>
[static_type_nth(x:any) : class
  -> case x (Variable let s := x.range in
                      (case s (Param (if (s.params[1] = of) class!(the(s.args[1])) 
                                      else any), 
                               any any)),
             any any)]

// end of file
