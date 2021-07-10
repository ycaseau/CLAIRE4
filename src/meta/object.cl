//+-------------------------------------------------------------+
//| CLAIRE                                                      |
//| object.cl                                                   |
//| Copyright (C) 1994 - 2013 Yves Caseau. All Rights Reserved  |
//| cf. copyright info in about()                               |
//+-------------------------------------------------------------+

// ---------------------------------------------------------------------
// This file contains the definition of the objects that implement the
// core features of the microCLAIRE library: traceable & debug-able calls,
// tables, demons and exceptions
// ---------------------------------------------------------------------

// *********************************************************************
// *  Table of contents                                                *
// *   Part 1: Ask, debug & trace                                      *
// *   Part 2: Tables                                                  *
// *   Part 3: Demons & relations for the logic modules                *
// *   Part 4: Basics of Exceptions                                    *
// *********************************************************************


// release() should produce a version number
(system.version := Id(RELEASE),
 printf("-- CLAIRE run-time library v 3.~A [os: ~A, C++:~A ] --\n",
        Id(RELEASE),Id(compiler.env), Id(compiler.external)))

[release() : any -> ("3." /+ string!(system.version)) ]

// the about method produces the legal warning, according to the GNU software
// recommendation
about() : any
 -> (printf("CLAIRE v3.~A Copyright (C) 1994-2013 Yves Caseau. All Rights Reserved.\n",
            version()),
     printf("use and redistribution in source code or binary forms are permitted\n"),
     printf("resale is not permitted without the explicit agreement of Yves Caseau\n"),
     printf("THIS SOFTWARE IS PROVIDED AS IS AND WITHOUT ANY WARRANTY, INCLUDING,\n"),
     printf("WITHOUT LIMITATION, THE IMPLIED WARRANTIES OF MERCHANTABILTY AND FITNESS\n"),
     printf("FOR A PARTICULAR PURPOSE\n"),
     true)

// *********************************************************************
// *   Part 1: Ask, debug & trace                                      *
// *********************************************************************

// create the list of arguments if needed : allocate on the stack
[get_args(i:integer) : list
 -> let liste := list<any>() in
       (while (i < index!()) (liste :add get_stack(i), i := i + 1), liste) ]

// evaluation of a message without the message structure, with a list
// of arguments. This method must be garbage-protected, because it is
// used as an entry point.
// to remove !!!!
;ask(self:any,p:property,arg:list) : any
; -> (let start := index!() in
;       (push!(self),
;        for x in arg push!(x),
;        eval_message(p, find_which(p, start, owner(self)), start, true)))

// a simple method for a direct call with no argument
funcall(self:method,x:any) : any
 -> (let start := index!() in
       (push!(x), execute(self, start, false)))

// this is a simple method for calling directly a method with one argument
funcall(self:method,x:any,y:any) : any
 -> (let start := index!() in
       (push!(x), push!(y), execute(self, start, false)))

// this is a simple method for calling directly a method with two arguments
funcall(self:method,x:any,y:any,z:any) : any
 -> (let start := index!() in
       (push!(x),
        push!(y),
        push!(z),
        execute(self, start, false)))

// how to apply a function to a list
apply(self:function,ls:list,l:list) : any
 -> (let start := index!() in
       (for x in l push!(x),
        stack_apply(self, ls, start, index!())))
call(p:property,l:listargs) : any -> apply(p, l)

apply(p:property,l:list) : any
 -> let start := index!() in
            (for x in l push!(x),
             eval_message(p, find_which(p, start, owner(l[1])), start,
                          true))

apply(m:method,l:list) : any
 -> (let start := index!() in
       (for x in l push!(x), execute(m, start, false)))


// push and pop debug info on the stack
// this method also does the tracing and the steppping
// NOTE: self should be either a property or a restriction
push_debug(prop:property,arity:integer,start:integer) : void
 -> (let i := index!(),
         n := system.trace! in
       (if (n > 0 &
            ((prop.trace! + system.verbose) > 4 | n = system.step!))
           let p := use_as_output(system.ctrace) in
             (put(trace!, system, 0),
              printf("~I ~S(~S", tr_indent(false, n), prop, get_stack(start)),
              let j := (start + 1) in
                while (j < (start + arity)) (printf(",~S", get_stack(j)), j :+ 1),
              if (prop.trace! >= 1000) system.step! := n,
              if (system.count_call >= 0)
                 (system.count_call :+ 1, printf(" [~A]",system.count_call),
                  if (system.count_call = system.count_level)
                     (if (system.count_trigger = call_step) system.step! := n
                      else if (system.count_trigger = spy) system.spy! := spy @ void
                      else system.verbose := system.count_trigger) ),
              if (n = system.step!) call_step(prop)
              else printf(")\n"),
              system.trace! := n + 1,
              use_as_output(p))
        else if (n > 0 & system.step! > 0)
           (if (system.count_call >= 0) system.count_call :+ 1, system.trace! := n + 1),
        if known?(StopProperty,prop)
           (if (StopProperty[prop] = nil |
                exists( l2 in StopProperty[prop] |
                        forall(j in (1 .. length(l2)) |(j + start <= i &
                                                 l2[j] = get_stack(start + j - 1)))))
                error("stop as required in ~S(~A)",prop,get_args(start))),
        push!(system.debug!),
        push!(prop),
        push!(arity),
        push!(start),
        put(debug!, system, i)))

// value of the previous debug
// n is 0 for interpreted code and 1 for compiled code
pop_debug(self:property,n:integer,val:any) : void
 -> (let v:integer := system.debug! in
       (if (v > 0)
           (if (n != 0) set_index(get_stack(v + 3) as integer),
            put(debug!, system, get_stack(system.debug!)),
            if (self.if_write = unknown) 
              let m := system.spy! in
                (if (m != unknown)                  // v3.3.14
                   (put(spy!, system, unknown),     // v3.0.3
                    funcall(m as method, system),
                    write(spy!, system, m))),
            if (system.trace! > 1 &
                ((self.trace! + system.verbose) > 4 | system.step! != 0))
               let i:integer := system.trace! in
                 (write(trace!, system, 0),
                  if ((self.trace! + system.verbose) > 4)
                     let p := use_as_output(system.ctrace) in
                       (printf("~I ~S\n", tr_indent(true, i - 1), val),
                        use_as_output(p)),
                  if (i <= system.step!) write(step!, system, i - 1),
                  write(trace!, system, i - 1),
                  if (system.trace! = 1) write(step!, system, 0)))))

// print a nice indented mark
tr_indent(return?:boolean,n:integer) : void
 -> (if return? printf("[~A]", n)
     else printf("~A:=", n),
     while (n > 9) (princ("="), n := n - 10),
     while (n > 0) (princ(">"), n := n - 1))

// *********************************************************************
// *   Part 2: Tables                                                  *
// *********************************************************************

// finds if objects are identified
identified?(self:class) : boolean
 -> (self = integer | self inherit? object | self = symbol | self = boolean |
     self = char)

identical?(x:any,y:any) : boolean
  ->  externC("((x == y) ? CTRUE : CFALSE)",boolean)

//  let x1: (if ((x as boolean) = (y as boolean)) true else false)

// writing a single value into a slot but does NOT trigger the rules !
// equivalent to is! of LAURE
// this definition should not be placed in the method.cl file
// (it requires some inheritance conflict processing)
put(self:property,x:object,y:any) : any
 -> (let s := (self @ owner(x)) in
       case s
        (slot store(x, s.index, s.srange, y, self.store?),
         any selector_error(selector = self, arg = list(x))))

         
// v3.2 : same but multi valued
[add_value(self:property,x:object,y:any) : void
  -> let s := (self @ owner(x)) in
    (if not(s) selector_error(selector = self, arg = list(x))
     else if not(multi?(self)) error("[134] Cannot apply add to ~S", self)
     else let n := (s as slot).index,
              l1 := (slot_get(x, n, object) as bag) in
            add_value(self, x, n, l1, y)) ]

         
// access
nth(a:table,x:any) : type[(if unique?(a) the(a).range else any)]
 -> let p := a.params in
       (if not(x % a.domain)
           error("[135] ~S does not belong to the domain of ~S", x, a),
        let v := (case p
                  (integer a.graph[(x as integer) - p],
                   list a.graph[get_index(a, (x as list)[1], (x as list)[2])],
                   any let i := index(a, x) in a.graph[i])) in
          (if (known?(v) | v % a.range) v
           else error("[138] the value ~S(~S) is unknown !",a,x)))

get(a:table,x:any) : type[(if unique?(a) (the(a).range U {unknown}) else any)]
 -> (let p := a.params in
       (if not(x % a.domain) unknown
        else let i := get_index(a, x) in a.graph[i]))

// interface update method for a[x] := y
nth=(a:table,x:any,y:any) : void
 -> (if not(x % a.domain)
        error("[135] ~S does not belong to the domain of ~S", x, a),
     if not(y % a.range) range_error(cause = a, arg = y, wrong = a.range),
     nth_put(a, x, y)) 

// internal form without checks
// equivalent of update = put + put_inverse
nth_put(a:table,x:any,y:any) : void
 -> (if  (known?(if_write,a) & not(multi?(a))) fastcall(a,x,y)
     else if multi?(a)
        let r := get(inverse, a),
            old := get(a,x) in  // v3.3.38 : redo (thanks to <sb>)
          (a.graph[get_index(a, x)] :=
             (if (length((y as set)) = 0) y    // we install a new value -> direct write 
              else if (a.multivalued? = list) 
                    make_list(0,of_extract(a.range),0)  // watch out: a.default does not always exist
              else cast!(set(),of_extract(a.range))),   // of_extract is a fast member(..)
           if (old != unknown & known?(r)) for z in old update-(r, z, x),
           for z in (y as set) add!(a, x, z))
     else let  r := get(inverse, a), z := get(a,x) in
             (if (z != y)
                 (if known?(r)
                   let z := get(a,x) in
                      (if (known?(z) & (r != a | x != z)) update-(r, z, x)),
                       put(a, x, y),
                       update+(a, x, y))))

// put does NOT update the inverse
put(a:table,x:any,y:any) : void
 -> let p := a.params, z := get(a,x)  in
       (if (z != y)
           (case p
             (integer store(a.graph, x - p, y, a.store?),
              list store(a.graph,
                         get_index(a, (x as list)[1], (x as list)[2]), y,
                         a.store?),
              any let i := index(a, x) in store(a.graph, i, y, a.store?)),
            true))

// adds a value to a multi-valued table: interface method
add(a:table,x:any,y:any) : void
 -> (if not(x % a.domain)
        error("[135] ~S does not belong to the domain of ~S", x, a),
     if not(y % member(a.range)) range_error(cause = a, arg = y, wrong = a.range),
     add!(a, x, y))

// adds a value to a multi-valued table: internal version without type checks
add!(a:table,x:any,y:any) : void
 -> (if known?(if_write,a) fastcall(a,x,y)
     else let p := a.params,
         i := get_index(a, x),
         l := (a.graph[i] as bag) in
       (if add_value(a, i, l, y) update+(a, x, y)))

// this methods adds a value to a multi-slot (used by the compiler)
add_value(self:table,n:integer,l:bag,y:any) : boolean
 -> (if (self.multivalued? = true)
        (if not(y % l)
            let l1 := (add!@set((if self.store? copy(l) else l), y)) in
              (store(self.graph, n, l1, self.store?), true)
         else false)
     else let l1 := (if self.store? store(l, y) else add!@list(l, y)) in
            (store(self.graph, n, l1, self.store?), true))

// a direct version (v3.2) that can be used in lieu of add!
add_value(self:table,x:any,y:any) : void
 -> (let p := self.params,
         i := get_index(self, x),
         l := (self.graph[i] as bag) in
       add_value(self, i, l, y))
            
// removes a value from an table
delete(a:table,x:any,y:any) : any
 -> (let p := a.params,
         i := get_index(a, x),
         l1 := (a.graph[i] as set),
         l := ((if a.store? copy(l1) else l1) delete y) in
       (store(a.graph, i, l, a.store?),
        let r := a.inverse in (if known?(r) update-(r, y, x)),
        l))

// direct access to 2-dim tables
[nth(a:table,x:any,y:any) : type[(if unique?(a) the(a).range else any)]
 -> let p := a.params,
        v := (case p
                (list (if not((x % (a.domain as tuple)[1] &
                               y % (a.domain as tuple)[2]))
                         error("[135] ~S does not belong to the domain of ~S", x, a),
                       a.graph[get_index(a, x, y)]),
                 any index(a, x, y)))  in
         (if (known?(v) | v % a.range) v else error("~S(~S) is unknown !",a,x))]

// sets a value in a 2-dim table
nth=(a:table,x:any,y:any,z:any) : void
 -> let p := a.params in
      (case p
        (list (if not((x % (a.domain as tuple)[1] &
                       y % (a.domain as tuple)[2]))
                  error("[135] ~S does not belong to the domain of ~S", list(x,y), a),
               if not(z % a.range)
                  range_error(cause = a, arg = z, wrong = a.range),
               if (known?(inverse, a) | known?(if_write, a))
                  nth_put(a, list(x, y), z)
               else store(a.graph, get_index(a, x, y), z, a.store?)),
         any nth=(a, tuple(x, y), z)))          // v3.2.16 tuple(a,b) is not list(a,b) !

get_index(a:table,x:any) : integer
 -> (let p := a.params in
       case p
        (integer (x as integer) - p,
         list get_index(a, (x as list)[1], (x as list)[2]),
         any index(a, x)))

get_index(a:table,x:integer,y:integer) : integer
 -> let p := (a.params as list<integer>) in (((p[1] * x) + y) - p[2])


// erase an table means to clean its graph so that it becomes empty.
erase(a:table)  : void
  -> let p := a.params in
       (case p
         (integer (for i in domain(a) a.graph[get_index(a,i)] := a.default),
          list (for l:list in domain(a) a.graph[get_index(a,l[1],l[2])] := a.default),
          any (for i in (1 .. length(a.graph))
                  (a.graph as list)[i] := unknown)))

// new in v3.2.50 a constructor for building a table dynamically
claire/make_table(%domain:type, %range:type, %default:any) : table
  -> let t := (mClaire/new!(table) as table) in   
       (t.range := %range, 
        table.instances :add t,                          // v3.3.3
        t.domain := %domain,
        t.default := %default,
        t.params := any,
        t.mClaire/graph := make_list(29,unknown),
        t)


// Our first table: a debuging tool which stores a list of stopping values
StopProperty[p:property] : list := unknown

// *********************************************************************
//   Part 3: Demons & relations for the logic modules                  *
// *********************************************************************

// applying a lambda to one argument
funcall(self:lambda,x:any) : any
 -> (let start := mClaire/index!(),
         retour := mClaire/base!() in
       (mClaire/set_base(start),
        mClaire/push!(x),
        mClaire/stack_apply(self.dimension),
        let val := eval(self.body) in
          (mClaire/set_base(retour), mClaire/set_index(start), val)))

// applying a lambda to two argument
[funcall(self:lambda,x:any,y:any) : any
 ->  let start := mClaire/index!(),
         retour := mClaire/base!() in
       (mClaire/set_base(start),
        mClaire/push!(x),
        mClaire/push!(y),
        mClaire/stack_apply(self.dimension),
        let val := eval(self.body) in
          (mClaire/set_base(retour),
           mClaire/set_index(start),
           val)) ]

// applying a lambda to two argument
[funcall(self:lambda,x:any,y:any,z:any) : any
 ->  let start := mClaire/index!(),
         retour := mClaire/base!() in
       (mClaire/set_base(start),
        mClaire/push!(x),
        mClaire/push!(y),
        mClaire/push!(z),
        mClaire/stack_apply(self.dimension),
        let val := eval(self.body) in
          (mClaire/set_base(retour),
           mClaire/set_index(start),
           val)) ]


// for historical reasons
mClaire/pname :: property()

// dealing with inverse
check_inverse(%r1:any,%r2:any) : void
 -> (let r1 := (%r1 as relation),
         r2 := (%r2 as relation) in
       (put(inverse,r1,r2),
        put(inverse,r2,r1),
        final(r1),
        final(r2),
        if (not(r1.domain <=  (if multi?(r2) member(r2.range) else r2.range)) |
            not(r2.domain <=  (if multi?(r1) member(r1.range) else r1.range)))
           error("[137] ~S and ~S cannot be inverses for one another", r1, r2)))

(relation.open := 0,
 write(inverse, inverse, inverse),
 write(if_write, inverse, (check_inverse @ any).functional))


// very useful
invert(r:relation,x:any) : bag
 -> (let r2 := get(inverse, r) in
       case r2
        (table let v := r2[x] in (if (r2.multivalued? != false) (v as bag) else set(v)),
         property let v := get(r2, x) in
                    (if (r2.multivalued? != false) (v as bag) else set(v)),
         any case r
             (property (if (r.multivalued? != false) { z in r.domain | x % get(r, z)}
                       else { z in r.domain | get(r, z) = x}),
              table (if (r.multivalued? != false) { z in r.domain | x % r[z]}
                     else { z in r.domain | r[z] = x}))))

// same: two useful methods that are used often
domain!(x:restriction) : class -> class!(x.domain[1])
methods(d:class,r:class) : set
 -> { m in method | (m.domain[1] <= d & m.range <= r)}

// sets the reified flag
claire/reify  :: property()
reify(l:listargs) : void
  -> (for p in l
       (case p (property p.reified := true)))

// *********************************************************************
// *   Part 4: Basics of Exceptions                                    *
// *********************************************************************

**arg :: property(open = 0)
args :: property(open = 0)
value :: property()

// a generic error that is produced by the error(" ....") instruction
general_error <: error(cause:any,arg:any)
self_print(self:general_error) : void
 -> printf("**** An error has occurred.\n~I\n", format(self.cause, self.arg))

// a read_slot error is produced when an unknown value is found
read_slot_error <: error(arg:any,wrong:any)
self_print(self:read_slot_error) : void
    -> printf("****[138] The value of ~S(~S) is unknown", self.wrong, self.arg)

// range errors
range_error <: error(cause:any,arg:any,wrong:any)
self_print(self:range_error) : void
    -> printf("****[139] ~S: range error, ~S does not belong? to ~S.\n",
              self.cause, self.arg, self.wrong)

// selector errors
selector_error <: error(selector:any,arg:any)
self_print(self:selector_error) : void
  -> (let p := self.selector in
        (if not(p.restrictions)
          printf("[140] The property ~S is not defined (was applied to ~S).\n", p, self.arg)
         else printf("****[141] ~S is a wrong arg list for ~S.\n",self.arg, p)))

// produced by a return (usually trapped)
return_error <: error(arg:any)
self_print(self:return_error) : void
  -> printf("****[142] return called outside of a loop (for or while).")

// interpretation of all the error codes
self_print(self:system_error) : void
 -> (let n := self.index in
       (printf("**** An internal error [~A] has occured:\n", n),
        format((if (n = 1) "dynamic allocation, item is too big (~S)"
                else if (n = 2) "dynamic allocation, too large for available memory (~S)"
                else if (n = 3) "object allocation, too large for available memory (~S)"
                else if (n = 5) "nth[~S] outside of scope for ~S"
                else if (n = 7) "Skip applied on ~S with a negative argument ~S"
                else if (n = 8) "List operation: cdr(()) is undefined"
                else if (n = 9) "String buffer is full: ~S"
                else if (n = 10) "Cannot create an imported entity from NULL reference"
                else if (n = 11) "nth_string[~S]: string too short~S"
                else if (n = 12) "Symbol Table table full"
                else if (n = 13) "Cannot create a subclass for ~S [~A]"
                else if (n = 16) "Temporary output string buffer too small"
                else if (n = 17) "Bag Type Error: ~S cannot be added to ~S"
                else if (n = 18) "definition of ~S is in conflict with an object from ~S"
                else if (n = 19) "Integer overflow"
                else if (n = 20) "Integer arithmetic: division/modulo of ~A by 0"
                else if (n = 21) "Integer to character: ~S is a wrong value"
                else if (n = 22) "Cannote create a string with negative length ~S"
                else if (n = 23) "Not enough memory to instal claire"
                else if (n = 24) "execution stack is full [~A]"
                else if (n = 26) "Wrong usage of time counter [~A]"
                else if (n = 27) "internal garbage protection stack overflow"
                else if (n = 28) "the multivalued status of ~S is not compatible with ~S"
                else if (n = 29) "There is no module ~S"
                else if (n = 30) "Attempt to read a private symbol ~S"
                else if (n = 31) "External function not compiled yet"
                else if (n = 32) "Too many arguments (~S) for function ~S"
                else if (n = 33) "Exception handling: stack overflow"
                else if (n = 34) "User interrupt: EXECUTION ABORTED"
                else if (n = 35) "reading char '~S': wrong char: ~S"
                else if (n = 36) "cannot open file ~A"
                else if (n = 37) "world stack is full"
                else if (n = 38) "Undefined access to ~S"
                else if (n = 39) "cannot convert ~S to an integer"
                else if (n = 40) "integer multiplication overflow with ~S and ~S"
                else if (n = 41) "wrong NTH access on ~S and ~S"
                else if (n = 42) "Wrong array[~S] init value: ~S"
                else if (n = 43) "Defeasible addition on list ~S requires pre-allocation (size ~S)" // v3.3.06
                else if (n = 50) "C++ imported error (~S) : ~S"   // NEW IN v3.1.04 (backdoor)
                else (self.value := n,
                      "What the hell is this ! [code: ~S^]")),
               list(self.value, self.arg))))

// contradictions are nice exceptions
contradiction <: exception()
self_print(x:contradiction) : void
 -> printf("A contradiction has occured.")

// the format method is used to print error messages (similar to a printf)
[format(self:string,larg:list) : void
 ->  let s := self,
         n := get(s, '~'),
         l := copy(larg) in
       (while not(n = 0)
          let m := s[n + 1] in
            (if (n > 1) princ(substring(s, 1, n - 1)),
             if ('A' = m) princ(car(l))
             else if ('S' = m) print(car(l))
             else if ('I' = m) error("[143] ~I not allowed in format", unknown),
             if (m != '%') l := l << 1,
             s := substring(s, n + 2, 1000),
             n := get(s, '~')),
        if (length(s) > 0) princ(s)) ]

// special version that prints in the trace port
tformat(self:string,i:integer,l:list) : any
 -> (if (i <= system.verbose)
        let p := use_as_output(system.ctrace) in
          (format(self, l), use_as_output(p)))

// printing a bag without ( )
princ(s:bag) : void
 -> (let f := true in
       for x in s (if f f := false else princ(","), print(x)))

// a global variable is a named object with a special evaluation
// NOTE: we need to refine the scheme for global constants !
global_variable <: system_thing(
    value:any,                    // the value
    range:type = any,             // the range is a type, {} means a global constant
    store?:boolean = false)       // GV are defeasible

close(self:global_variable) : global_variable
 -> (if not((unknown?(value, self) |
             (self.range % set | self.value % self.range)))
        range_error(arg = self.value, cause = self, wrong = self.range),
     self)

<=2 :: operation()

// we create a spcial contraidiction that we shall reuse
contradiction_occurs :: global_variable(range = contradiction,
                                        value = new!(contradiction))

// how to use it
contradiction!(system) : void -> close(contradiction_occurs)
nil :: global_variable(range = {}, value = Id(nil))       // v0.01
claire_date:string :: Id(date!(1))


// end of file
