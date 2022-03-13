//+-------------------------------------------------------------+
//| CLAIRE                                                      |
//| types.cl                                                    |
//| Copyright (C) 1994 - 2021 Yves Caseau. All Rights Reserved  |
//| cf. copyright info in file object.cl: about()               |
//+-------------------------------------------------------------+

// --------------------------------------------------------------------
// This file contains the definition of the CLAIRE type system (a true lattice).
// that is used both at compile- and at run-time.
// --------------------------------------------------------------------

// ******************************************************************
// *  Table of contents                                             *
// *    Part 1: Common Set Methods                                  *
// *    Part 2: definition of the type operators                    *
// *    Part 3: Interface methods                                   *
// *    Part 4: Lattice methods                                     *
// *    Part 5: Type methods                                        *
// ******************************************************************

claire/-- :: operation(precedence = precedence(..))

// *********************************************************************
// *   Part 1: Common Set Methods                                      *
// *********************************************************************

// ----------------------- useful methods ------------------------------
[finite?(self:type) : boolean
 -> case self
      (set true,
       list not({ t in self | not(finite?(t))}),
       class let n := self.open in 
              ((n = open()) | (n = final()) | (n = close()) |
                    (n = abstract() &  forall(c in self.subclass | finite?(c)))),
       any false) ]


// making a set from an abstract_set  (CLAIRE 4 : bag is not longer a concrete type)
// this is a list since order matters in enumeration
[enumerate(self:any) : list
 -> case self
      (list self,
       set list!(self),
       array list!(self),
       class let l := list<object>() in
               (for c  in self.descendants  l := l /+ c.instances, l),
       Interval --!(self.arg1,self.arg2),
       integer list!(make_set(self)),
       collection list!(set!(self)),        // TODO : change to list!(self)
       any error("[178] cannot enumerate ~S",self)) ]

// =type? is an operation (equality on types)
=type? :: operation()
[=type?(self:type,ens:type) : boolean -> (self <=t ens & ens <=t self)]

// finds the sort associated to a type
[sort!(x:type) : class
 -> case x
      (class sort!(x),
       any sort!(class!(x))) ]

// the membership for classes
[%(self:any,ens:class) : boolean
  -> if inherit?(owner(self),ens) true else false]

/* an extension for %
[%type(x:any,y:any) : boolean
  -> if (x % type & y % set) exists(z in (y as set) | =type?(x,z))     // v3.2.28
     else (x % y) ] */

// v4.0 : belong is the unique method (static call for any) for membership
// replaces belong_to + member? in claire 3 => works on everything, collections and integer as well :)
// see belong_exp in gexp.cl to see how it is used + open-conding patterns
// note that belong may create an error => heavier => optimize with %t when possible 
[belong(x:any,y:any) : boolean
 -> case y
      (class x % y,       // open coded
       list x % y,       // open coded
       set x %t y,       // kernel definition (includes the =type? iteration)
       array x % y,
       tuple (x % tuple & length(x) = length(y) & forall(i in (1 .. length(x)) | x[i] % y[i])),     // tuple as a type !
       type_operator x %t y,      // kernel (closed) definition (contains)
       integer (case x (integer (x % y), any false)),
       any let start := index!() in         // this is the extensibility part for collections
             (push!(x),
              push!(y),
              let m := find_which(%, start, owner(x)) in
                (if (case m
                      (method
                         (length(m.domain) = 2 & m.domain[2] != any)))
                    (eval_message(%, m, start, true) as boolean)
                 else error("[179] (~S % ~S): not implemented!", x, y))))]

// x % y is a short cut 
// CLAIRE4 : cannot be a macro (too early)
[%(x:any,y:any) : boolean -> belong(x,y)]


// ****************************************************************
// *         Part 2: definition of the type operators             *
// ****************************************************************

// in CLAIRE4, types are defined in the Kernel go module
// type_operator <: type()

// union of two types ---------------------------------------------

// Disjonctive Union Axiom (DU): Each union (A U B) is stricly disjunctive:
//       (1) A ^B = 0
//       (2) x < A U B <=> x < A or x < B
// Producing disjunction union is a form of normalization (the previous notion
// of diustributivity was a lousy bug)
// DU Axiom is necessary to make <= and ^ easier to define
// This is achieved in the U method
self_print(self:Union) : void -> printf("(~S U ~S)", self.t1, self.t2)
finite?(self:Union) : boolean -> (finite?(self.t1) & finite?(self.t2))

// Intervals of integers ----------
self_print(self:Interval) : void
 -> (printf("(~S .. ~S)", self.arg1, self.arg2))

finite?(self:Interval) : boolean -> true

// true constructor
--(x:integer,y:integer) : Interval
  -> (if (x <= y) (x .. y) as Interval // Interval(arg1 = x, arg2 = y)
      else error("[182] the interval (~S -- ~S) is empty",x,y))

// Parameterized class. -------------------------------------------

[self_print(self:Param) : void
 -> if (length(self.params) = 1 & self.params[1] = of & self.args[1] % set)
       printf("~S<~S>", self.arg, the(self.args[1] as set))
    else printf("~S[~I]", self.arg,
           (for i in (1 .. length(self.args))
              (if (i != 1) princ(", "),
               printf("~S:(~S)", self.params[i], self.args[i])))) ]

finite?(self:Param) : boolean -> finite?(self.arg)

// subtype[X] ----------------------------------------------
// subtype[X] = {u in type | u <= t}
// for closure purposes, we add an arg Y -> Y inter st[X]
// Y can be any type class, but we forbid parametrisation on such classes !
// thus we can ensure that Y is a class
[self_print(self:subtype) : void
   -> if (self.arg = type) printf("subtype[~S]", self.t1)
      else printf("~S[~S]", self.arg, self.t1) ]            // v3.2

finite?(self:subtype) : boolean -> (self.arg = set & finite?(self.t1))

// creates a subtype, with some normalization
// v3.2 list[t] -> subtype 
// v4.0 => no error
[nth(self:class,x:type) : type
 -> (if (self = set | self = list)  subtype(arg = self, t1 = x)
            ;   Param(arg = self, params = list(of),
            ;         args = list(subtype( arg = type, t1 = x)))
      else if not(self inherit? type)  {} // error("[177] subtyping of ~S not allowed",self)
      else subtype(arg =  (if (self = subtype) type else self), t1 = x)) ]

// create a Param with a list of parameters (constant properties) l1 and a list
// of types l2
// v4.0 => no error
[nth(self:class,l1:list,l2:list) : type
 -> // patch in v3.2.36  -> rather ugly ... the whole processing of X[Y..] pattern should be revised
    if ((self = list | self = set) & l2[1] % subtype)
        nth(self, (l2[1] as subtype).t1)   // this is the patch
    else if ((self = list | self = set) & l1[1] != of) {}
        // error("[177] the subtyping expression ~S[~A] is not allowed",self,l1) // v3.2.36
    else Param(arg = self, params = l1, args = l2)]

// create a Param of the stack[X] kind
[param!(self:class,tx:type) : type
 ->  Param(arg = self, params = list(of), args = list(set(tx))) ]

// create the t[] param
[nth(self:type) : type
 ->  Param(arg = array, params = list(of), args = list(set(self))) ]

// tuple are types
finite?(self:tuple) : boolean -> forall(x in self | finite?(x as type))

// reference to a previous variable, not a type but a pattern -------
// index is the position of the stack of the referred type
// args is a list representing the path (a sequence of properties (parameters))
// a property is applied to the referred type
// if arg = true, the reference is the singleton containing the ref. value
// arg is set to true when we copy a reference in define.cl (unclear why)
self_print(self:Reference) : void -> printf("<ref:~S(ltype[~A])>",self.args,self.index)
get(self:Reference,y:any) : any
 -> (let l := self.args in
       (for i in (1 .. length(l)) y := (unsafe(funcall(l[i] as property, y))),
        y))

// we need a constructor
claire/Reference!(l:list,n:integer) : Reference
  -> Reference(args = l, index = n)       

// apply a reference to a type (l is args(self), passed for disambiguation)
@(self:Reference,l:list,y:any) : any
 -> (//[5] apply[@] ~S to ~S // self,y,
     for i in (1 .. length(l)) y := y @ (l[i] as property), y)


// type to set coercion  -------------------------------------------------

// new in v3.0.5 = use an interface method for type enumeration

// the default strategy is extensible: we look if there exists
// a proper definition that could be interpreted !
set!(x:collection) : set
  -> let m := (set! @ x.isa) in
       (if (domain!(m) != collection) (funcall(m,x) as set)
        else error("[178] cannot enumerate ~S",x))

size(x:collection) : integer
  -> let m := (size @ x.isa) in
       (if (domain!(m) != collection) (funcall(m,x) as integer)
        else size(set!(x)))                 // v3.2.34  -> makes the API simpler

// (interface(size))

// set is needed for recursive def
set!(x:set) : set -> x

// set is needed for recursive def
size(x:list) : integer -> size(set!(x))

// class  -> return a read-only list  (v3.2)
set!(x:class) : set
  -> let rep := list() in
       (for c in x.descendants
          (if (inherit?(c,primitive) & c != boolean)
              error("[178] cannot enumerate ~S",c)
           else rep := rep /+ c.instances),
        set!(rep))

size(self:class) : integer
  -> let n:integer := 0 in
        (for x in self.descendants n :+ length(x.instances), n)


// Union
set!(x:Union) : set -> (set!(x.t1) /+ set!(x.t2))

size(x:Union) : integer
  ->  (if (x.t1 % Interval | x.t1 % set)  (size(x.t1) + size(x.t2))
       else size(set!(x)))

// interval
set!(x:Interval) : set
  -> (x.arg1 claire/--? x.arg2)

size(self:Interval) : integer
  -> (self.arg2 + 1 - self.arg1)

// param
set!(x:Param) : set
  -> { y in set!(x.arg) | y % x}
size(x:Param) : integer -> size(set!(x))

// subtype
set!(x:subtype) : set
  ->  (if (x.arg = set) build_powerset(list!(set!(x.t1)))
       else error("[178] cannot enumerate ~S",x))
size(x:subtype) : integer
  -> (if (x.arg = set) ^2(size(x.t1))
      else error("[178] cannot enumerate ~S",x))

// tuple
set!(x:tuple) : set
  ->  let l := (x as list) in    // v3.0.54
        (if not(l) {{}}
         else let l1 := { list(y) | y in set!(l[1])} in
           (for n in (2 .. length(l))
             let l2 := set<any>() in
                (for z in set!(l[n])
                   for l3 in l1 l2 :add copy(l3) add z,
                 l1 := l2),
            l1))

size(l:tuple) : integer
  ->  (if not(l) 1
       else let m := size(l[1] as type) as integer in
          (for n in (2 .. length(l)) m :* size(l[n] as type),
           m))

// declarations (now useless in CLAIRE4)
(ephemeral(Union),
 ephemeral(Param),
 ephemeral(Interval),
 ephemeral(subtype))

// ********************************************************************
// *                Part 3: Interface Methods                         *
// ********************************************************************

// there is a special restriction for + to specify the way the inheritance
// conflict should be solved
//U(self:set,ens:type) : type -> (case ens (set self /+ ens, any ens U self))

// the union makes a partial reduction to the normal form. The complete
// reduction is done by enumeration if needed during the type subsumption
// union is left-associative: A U B U C is represented by (A U B) U C  => never(t2(x:Union) % union)
// a union of intervals is ALWAYS disjoint
U(x:type,y:type) : type
 -> (case x
      (set case y (set x /+ y, any y U x),
       any (if (y <=t x) x
           else if (x <=t y) y
           else if (y % Union) (x U y.t1) U y.t2
           else if (case x (Interval (y % Interval)))
              (if (y.arg1 - 1 <= x.arg2 & x.arg1 <= y.arg1)       // adjacent
                  x.arg1 .. y.arg2
               else if (x.arg1 - 1 <= y.arg2 & y.arg1 <= x.arg1)
                  y.arg1 .. x.arg2
               else Union(t1 = x, t2 = y))
           else if (case x (Union y % Interval))
              let z := (x.t2 U y) in
                case z
                 (Union Union(t1 = (x.t1 U y), t2 = x.t2),
                  any x.t1 U z)
           else if (case x (Interval (y % set & (x.arg1 - 1 % y | x.arg2 + 1 % y))))
               let a := x.arg1, b := x.arg2 in    // new in v0.30
                 (if (a - 1 % y) a :- 1,
                  if (b + 1 % y) b :+ 1,
                  (a .. b) U y)
           else (if (y % set) y := {z in (y as set) | not(z % x)},
                 Union(t1 = x, t2 = y)))))


// the Interval construction method has a smart second-order type  - fix on v3.1.06
[..(x:integer,y:integer) :  type[(if (unique?(x) & unique?(y) & the(x) <= the(y))
                                    set(the(x) .. the(y))
                                  else subtype[integer]) ]
 -> if (x <= y) Interval(arg1 = x, arg2 = y) else {} ]

// exception
but :: operation()
but(s:any,x:any) : type[(if (x <= list) nth(list,member(s)) else (if (x <= set) nth(set,member(s)) else any))]      // <yc> 16/3/98
 -> (case s (list list{y in s | y != x},         // v3.3.36 (thanks to fxj)
             set copy(s) delete x, 
             any Core/enumerate(s) delete x))

// a set difference (extended to types, with implicit enumeration)
\ :: operation(precedence = U.precedence)
\(x:type,y:type) : set -> {z in x | not(z % y)}


// ******************************************************************
// *    Part 4: Lattice methods                                     *
// ******************************************************************

// glb operation ---------------------------------------------------
// gbl is the extension of the lattice operator ^ for types to type_expressions

// new in v3.0.60: we reintroduce a glb method
claire/glb :: operation(precedence = precedence(^), domain = type_expression, range = type)

glb(x:set,y:type) : set -> { z in x | z % y}

glb(x:Union,y:type) : type ->  ((x.t1 glb y) U (x.t2 glb y))

glb(x:Interval,y:type) : type
 ->  (case y
       (class (if (integer <=t y) x else {}),
        set y glb x,
        Interval  (if (x.arg1 <= y.arg1)
                     (if (y.arg1 <= x.arg2)
                        (if (x.arg2 <= y.arg2) y.arg1 .. x.arg2
                         else y)
                      else {})
                   else y glb x),
        Union ((x glb y.t1) U (x glb y.t2)),    // v3.0.44
        any {}))

glb(x:class,y:type) : type
 -> (if (x.open = system.abstract & not(x.subclass)) { z in x.instances | z % y}
     else if (x.open = system.abstract & not(x.instances))
        Uall(list{ (z glb y) | z:type in x.subclass})
     else case y (class x join y, any y glb x))

glb(x:Param,y:type) : type
 -> (case y
      (Param let c := (x.arg join y.arg),
                 lp := list!(set!(x.params /+ y.params)),
                 l := list<any>() in
               (for p in lp
                  let t := ((x @ p) glb (y @ p)) in
                    (if (t != {}) l :add t else (c := {}, break(true))),
                if (c != {}) Param(arg = c, params = lp, args = l)
                else {}),
       class let c := (x.arg join y) in
               (if (c != {}) Param(arg = c, params = x.params, args = x.args)
                else {}),
       any y glb x))

// notice that a param whose class is a type must use of (only parameter allowed!)
// the result is a subtype
glb(x:subtype,y:type) : type
 -> (case y
      (class (if ((x.arg join y) != {}) nth(x.arg join y, x.t1)       // v3.00.07
             else {}),
       Param (if ((x.arg join y.arg) != {})
                 param!(x.arg join y.arg, member(x) glb member(y))
              else {}),
       subtype (if ((x.arg join y.arg) != {})
                  let t := (x.t1 glb y.t1) in
                    (if (t != {}) nth(x.arg join y.arg, t) else {})
               else {}),
       any y glb x))

// set, Interval, list
glb(x:tuple,y:type) : type
 -> (case y
      (class (if (tuple inherit? y) x else {}),    // v2.4 BUG
       Param {},
       tuple tuple!((x as list) ^ (y as list)),    // ^ on lists implements Cartesian product ^
       subtype (if (y.arg = tuple) tuple!(list{ (z glb y.t1) | z in x})
               else {}),
       any y glb x))


// a reference is seen as "any"
glb(x:Reference,y:type) : type -> y

// this will be greatly simplified in a few minutes !
^(x:type,y:type) : type -> (x glb y)

// the old lattice_glb
[join(x:class,y:class) : type
 -> let l1 := x.ancestors, n1 := length(l1),
        l2 := y.ancestors, n2 := length(l2) in
       (if (n1 < n2)
           (if (l2[n1] = x) y else {})
        else if (l1[n2] = y) x
        else {}) ]

// for lists
[^(x:list,y:list) : list
 ->  let n := length(x),
         r:list := list() in
       (if (n = length(y))
           for i in (1 .. n)
             let z := (x[i] glb y[i]) in
               (if (z != {}) r :add z else (r := nil, break(true))),
        r) ]


// a combined union
Uall(l:list) : type -> (let rep := {} in (for x in l rep :U x, rep))

// ------------------- The inclusion operation ------------------------
// the specialized versions %t and <=t are hard coded in Kernel, hence not extensible.
// if we create new types they will be used as patterns, not concrete types.

Core/<=t :: operation(precedence = precedence(<=), domain = type, range = boolean)   // hand-made

// v4 open coded (link to Included kernel method)
<=t(s:type,y:type) : boolean  -> (s <=t y)
 
// default order for types
<=(x:type_expression,y:type_expression) : boolean  
 -> (case x (set forall(z in x | (z % y)),
             type (case y (type (x <=t y), 
                   any let z:any := x in less?(z,y))),   // this is ugly (using z) and should be fixed later
             any less?(x,y)))

// membership for types
Core/%t :: operation(precedence = precedence(<=), domain = type, range = boolean)   // hand-made

// v4 open coded (link to Contains kernel method)
%t(x:any,y:type) : boolean  -> (x %t y)

// extensibility for type_expression is through less?, that always returns a value (hence no error returned)
[less?(x:type_expression, y:type_expression) : boolean 
   -> false]

// ******************************************************************
// *    Part 5: type methods                                        *
// ******************************************************************

// --------------------- extract tuple type information -------------

// extract a member type, that is a valid type for all members (z) of instances of
// the type x.This is much simpler in v3.0
[member(x:type) : type
 ->  case x
        (class (if (x = Interval) integer else any),
         Union member(x.t1) U member(x.t2),
         Interval {},
         Param member(x @ of),
         tuple Uall(x as list),
         subtype x.t1,
         set Uall(list{ (case y (list set!(y), type y, any {})) | y in x}),
         any {}) ]

// a simpler version (projection on bag subtypes)
// dumb code because it is used early in the bootstrap
[of_extract(x:type) : type
 -> let c := x.isa in
      (if (c = subtype) (x as subtype).t1
       else if  (c = Param)
         (if ((x as Param).params[1] = of)
             let y := ((x as Param).args[1] as type) in
               case y (set (list!(y)[1] as type),
                       subtype y.t1,
                       any any)
          else any)
       else any) ]

// useful type functions for the compiler
[unique?(x:type) : boolean
 -> case x
      (set size(x) = 1,
       class (x.open = 0 & size(x) = 1),
       any false) ]

// returns the unique element of the type
the(x:type) : any -> list!(set!(x))[1]

// bitvector made easy
// v0.01: should not use set[0 .. 29] => burden on caller is too heavy
[integer!(s:set[integer]) : integer
 -> let n := 0 in
       (for y in s (if (y % (0 .. 29))  n :+ ^2(y)), n) ]

claire/make_set(x:integer) : set -> {i in (0 .. 29) | x[i]}

// asbtract coercion of a set into an interval
[abstract_type(xt1:set) : type
 -> let m1 := 1,
        m2 := 0 in
       (for x in xt1
          case x
           (integer (if (m1 > m2) (m1 := x, m2 := x)
                    else if (x > m2) m2 := x
                    else if (x < m1) m1 := x),
            any (m1 := 1, m2 := 0, break(true))),
        m1 .. m2) ]

// abstract interpretation of integer arithmetique
[abstract_type(p:operation,xt1:type,xt2:type) : type
 -> case xt1
      (set (if (xt1 != {}) abstract_type(p, abstract_type(xt1), xt2) else xt1),
       Interval case xt2
                (Interval
                   (if (p = +)
                       (xt1.arg1 + xt2.arg1) .. (xt1.arg2 + xt2.arg2)
                    else if (p = -)
                       (xt1.arg1 - xt2.arg2) .. (xt1.arg2 - xt2.arg1)
                    else integer),
                 set (if (xt2 != {}) abstract_type(p, xt1, abstract_type(xt2))
                     else xt2),
                 Union abstract_type(p, xt1, xt2.t1) U abstract_type(p, xt1, xt2.t2),
                 any integer),
       Union abstract_type(p, xt1.t1, xt2) U abstract_type(p, xt1.t2, xt2),
       any integer) ]

// we create some types that we need
(set_range(subclass, class, set<class>),
 set_range(ancestors, class, list<class>),
 set_range(descendants, class, set<class>),
 set_range(definition, property, list<restriction>),
 set_range(restrictions, property, list<restriction>),
 set_range(domain,restriction,list<type_expression>),
 set_range(slots, class, list<slot>))

// a useful second ortder type
first_arg_type(x:type,y:type) : type -> x
first_arg_type(x:type,y:type,z:type) : type -> x
second_arg_type(x:type,y:type) : type -> y
meet_arg_types(x:type,y:type) : type -> (x U y)
first_member_type(x:type,y:type) : type -> member(x)  // v3.3.10
// nth@bag (list / set) is now in Kernel (CLAIRE4)
nth_arg_type(x:type,y:type) : type 
   -> (if (x % tuple & unique?(y)) x[the(y)] else member(x))

// we place here all methods that require second order types !!!!
nth_get(a:array,n:integer) : type[member(a)] -> nth_get(a,n)                    // managed by cross-compiler ?
/* v4.0.5 defined in Kernel
nth(self:array,x:integer) : type[member(self)]
 -> (if (x > 0 & x <= length(self)) nth_get(self,x)
     else error("[180] nth[~S] out of scope for ~S", x, self)) */
make_array(i:integer,t:type,v:any) : type[ (if unique?(t) (the(t))[] else array)]
 -> function!(make_array_integer)
 
make_list(n:integer,t:type,x:any) : type[ (if unique?(t) list[the(t)] else list)]
  -> (cast!(make_list(n,x),t) as list) 

make_set(self:array<X>) : type[(if (X = any) set else set<X>)]
  -> set!(list!(self))

// these four functions are defined in Core with Kernel functions because we want to
// add second order types
list!(a:array<X>) : type[(if (X = any) list else list<X>)]
 -> function!(list_I_array)
array!(a:list<X>) : type[(if (X = any) array else array<X>)]
 -> function!(array_I_list)      // v3.0.72

set!(l:list<X>) : type[(if (X = any) set else set<X>)]  // v3.1.06
  -> function!(set_I_list)
list!(l:set<X>) : type[(if (X = any) list else list<X>)]
  -> function!(list_I_set)

 // get the type from class if a constant
 thing_type_class(x:type,y:type) : type ->  (thing glb member(x))
 object_type_class(x:type) : type ->  (object glb member(x))

// new in v3.0.60 : second-order type for copy
(for r in copy.restrictions (r as method).Kernel/typing := Id,
 for r in empty.restrictions (r as method).Kernel/typing := Id,
 for r in sort.restrictions (r as method).Kernel/typing := second_arg_type,
 for r in /+.restrictions (r as method).Kernel/typing := meet_arg_types,
 // here we add some simple second orders
 put(Kernel/typing, (new! @ list(class,symbol)), thing_type_class),
 put(Kernel/typing, (new! @ list(class)), object_type_class),
 (nth_get @ array).Kernel/typing := first_member_type,
 (nth @ list).Kernel/typing := nth_arg_type, 
 (nth @ array).Kernel/typing := nth_arg_type,     // v4.0.5
 (nth @ set).Kernel/typing := nth_arg_type,
 for r in nth+.restrictions (r as method).Kernel/typing := first_arg_type,
 for r in add.restrictions
    (if (length(r.domain) = 2) (r as method).Kernel/typing := first_arg_type),
 for r in delete.restrictions
    (if (length(r.domain) = 2) (r as method).Kernel/typing := first_arg_type))

