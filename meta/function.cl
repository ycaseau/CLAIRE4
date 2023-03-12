//+-------------------------------------------------------------+
//| CLAIRE                                                      |
//| function.cl                                                 |
//| Copyright (C) 1994 - 2021 Yves Caseau. All Rights Reserved  |
//| cf. copyright info in file object.cl: about()               |
//+-------------------------------------------------------------+

// --------------------------------------------------------------------
// This file hold the methods that are defined by an external function
// and those that apply to a primitive type
// --------------------------------------------------------------------

// *********************************************************************
// *  Table of contents                                                *
// *   Part 1: Basics of pretty printing                               *
// *   Part 2: Methods for CLAIRE objects                              *
// *   Part 3: System Methods                                          *
// *   Part 4: Methods for Native entities                             *
// *********************************************************************


// we find here what is necessary for the minimal kernel of CLAIRE
// ==============================================================
!= :: operation(precedence = 60)
<< :: operation(precedence = 10)
>> :: operation(precedence = 10)
and :: operation()
or :: operation()
U :: operation(precedence = 50)
& :: operation()
// min :: operation(precedence = 20)
// max :: operation(precedence = 20)
meet :: operation()
inherit? :: operation()
claire/--? :: operation(precedence = 30)
--! :: operation(precedence = 30)

// *********************************************************************
// *   Part 1: Basics of pretty printing                               *
// *********************************************************************

cpstack :: property()

// we use a nice object
pretty_printer <: thing(cpretty:port,       	// a string port
                        cprevious:integer = 0,  // index of the current port in the stack 
                        index:integer = 0,      // indentation level
                        width:integer = 75,	    // size of window
                        pprint:boolean = false, // active
		                    pbreak:boolean = false,
                        cpstack:list)           // support reccursive print-in-string 

pretty :: pretty_printer(cpretty = port!(), cpstack = nil)

apply_self_print :: property()
short_enough :: property()
// print :: property() - defined in Kernel

// buffered print
// new in v3.3.26: unbounded recursion is supported :-)
[print_in_string() : void
 -> let n := pretty.cprevious + 1,
        p1 := (if (n < length(pretty.cpstack)) (pretty.cpstack[n + 1] as port)
               else port!()),
        p2 := use_as_output(p1) in
      (pretty.cprevious := n, 
       pretty.cpretty := p1,
       if (pretty.cpstack = nil) pretty.cpstack := list<port>(p2,p1)        // initialisation
       else (pretty.cpstack[n] := p2,
             if (n = length(pretty.cpstack)) pretty.cpstack :add p1)) ]


[end_of_string() : string
  -> if (pretty.cprevious = 0) error("[123] unbalanced use of print-in-string"),
     let n := pretty.cprevious,
         s := string!(pretty.cpretty), 
         p := (pretty.cpstack[n]) as port in
       (set_length(pretty.cpretty, 0),
        use_as_output(p),
        pretty.cpretty := p,
        pretty.cprevious :- 1,
        s) ]

[buffer_length() : integer
  -> length(pretty.cpretty) ]


[buffer_set_length(i:integer) : void
  -> set_length(pretty.cpretty,i) ]

// a method for calling the printer without issuing a message
// here we assume that self_print is always defined as a function
[apply_self_print(self:any) : void
 -> if (self = unknown) princ("unknown")
    else case self
      (thing princ(self.name),
       class princ(self.name),
       integer princ(self),
       string self_print(self),
       any let %prop := (self_print @ owner(self)) as method in
             (if (%prop & known?(functional, %prop)) funcall(%prop.functional, self)
              else try self_print(self)
                   catch any printf("<unprintable:~S>", owner(self)))) ]

// some basic definitions
[self_print(self:any) : void
 ->  if (self = unknown) princ("unknown")
     else let c := owner(self), n := length(c.params) in
        (if (n > 0)
            printf("~S(~I)",c,
                   for i in (1 .. n) (print(get(c.params[i],self)), if (i < n) princ(",")))
         else printf("<~S>", c)) ]

[self_print(self:boolean) : void
 -> if self princ("true") else princ("false") ]

[self_print(self:function) : void 
  -> printf("#'~A", string!(self)) ]




// prints the name of a restriction. If we have a close property and if a
// short-cut is possible, we use it.
self_print(self:restriction) : void
 -> (if (not(known?(selector, self)) | unknown?(domain, self))
        printf("<~S>", owner(self))
     else let p := self.selector,
              n := 0,
              c := domain!(self) in
            (for r in p.restrictions (if (domain!(r) = c) n :+ 1),
             printf("~A @ ~S", p.name, (if (n = 1) c else self.domain))))

// we are too far
much_too_far <: error()
print(x:any) : void
 -> let %l := pretty.index in 
      (if (pretty.pbreak & pretty.pprint)
        let b_index := buffer_length(), missed := false in
          (if not(short_enough(b_index + 10))                       // no room for pretty ...
              (pretty.pprint := false, 
               pretty.pbreak := false,
               print(x),                                          // safe print
               pretty.pbreak := true,
               pretty.pprint := true)
           else (try (pretty.pbreak := false,                       // try the regular version first
                      apply_self_print(x),
                      pretty.pbreak := true)
                 catch much_too_far (missed := true),
               if missed (pretty.pprint := true, pretty.pbreak := true,
                          buffer_set_length(b_index),
                          pretty.index := %l,
                          apply_self_print(x))))
      else apply_self_print(x),
      pretty.index := %l,             // added in claire4
      unknown)

// short_enough = we expect that what we want to print is short enough (more that 10 chars to the width)
short_enough(self:integer) : boolean -> (self < pretty.width)

// *********************************************************************
// *   Part 2: Methods for CLAIRE objects                              *
// *********************************************************************

// the instantiation body is a sequence of words from which the initialization
// of the object must be built.
//  copied_def = object (for object) + float (for float) + integer (for all)
//               + NULL for objects
// v3.2.12: use a condition that is coherent with ClReflect.cl : a slot defaut value must be placed
// unless it is a copied_def
// CLAIRE4 (fix OLD bug): lp is list of p such that o.p := unknown is requested !
new_defaults(self:object,lp:list) : object
 -> (for s:slot in self.isa.slots
       let p := s.selector,
           s2 := s.srange,
           d := get(default, s) in
         (if known?(d)
             let v := slot_get(self, s.index, s.srange) in
               (if (v = unknown & p % lp) nil    // explicit self.p := unknown
                else if (unknown?(v) & not(s2 = object | d % integer | s2 = float))   // v3.2.12 coherence
                   // was: (s2 = any & not(d % integer) & unknown?(v))
                   update(p, self, s.index, s.srange, d)
                else if (d = v)
                   (if p.multivalued? for y in d update+(p, self, y)
                    else update+(p, self, d)))),
     let m := (close @ owner(self)) in
       (case m (method funcall(m, self) as object,
                any self)))  // v3.0.41  obviously



//-------------------------- ENTITY   --------------------------------------

not(self:any) : boolean
  -> (if (self = true) false
      else if (self = false) true
      else if not(self) true         // catch special cases : empty list
      else false)
      
!=(self:any,x:any) : boolean -> (if (self = x) false else true)

// gives the type of any object. This is open_coded.
owner(self:any) : class -> owner(self)

// some useful methods
known?(self:any) : boolean -> (unknown != self)
unknown?(self:any) : boolean -> (unknown = self)

// needed by the compiled code for casts (inserted by ocontrol for dynamic type checks)
// Claire 4: TODO - add a second order type
check_in(self:any,y:type) : any
 -> (if (self % y) self
     else error("[124] the value ~S does not belong to the range ~S", self, y))

// used to cast dynamically a non-mutable bag to a typed mutable bag
// claire 4 : aplied to list and sets (bags)
check_in(self:bag,c:class,y:type) : list
 -> let OK := (case self (list forall(z in self | z % y),
                           any forall(z in (self as set) | z % y))) in  
    (if OK cast!(self,y)
     else error("[124] the value ~S does not belong to subtype[~S]",self, y))

// new in v3.00.48
<(self:any,x:any) : boolean -> (if (self = x) false else (self <= x))
>(self:any,x:any) : boolean -> (if (self = x) false else (x <= self))
// >= is defined as a macro in file.cl

// unsafe is a pragma : it tells the compiler not to catch a possible error 
// to add in the documentation
claire/unsafe(x:any) : any -> x

// ----------------------- CLASS ---------------------------------------------

// classes have a "open" status with three states
// forward : -2          - use open as a marker for forward definition
// close : -1           ????
// abstract :  open = 0  - cannot create an instance
// final : open = 1  - no more subclasses or instance
// default : open : 2 - default for ephemeral - no instance
// open :  open = 3  - default for things (named things can be queried)


// declares a class as ephemeral: the member set is not maintained
// v3.2.14 recusively applies to subclasses
ephemeral(self:class) : any
  -> (for c in self.descendants
       (if c.instances error("[187] cannot declare ~S as ephemeral because of ~S has instances",self,c)
        else put(open, c, system.default)))

// claire4 : opposite of ephemeral(c)
instanced(c:class) : any
  -> let n := c.open in
       (if (n = system.default | n = system.open) put(open,c,system.open)
        else error("[125] abstract classes cannot be instanced"))


// declares a class as an abtract class (without instances)
abstract(c:class) : any
 -> let n := c.open in
       (if (n = close()) error("[125] closed classes cannot be abstract")
        else if (length(c.instances) != 0)
          error("[125] instanced classes cannot be abstract")
        else (write(open, c, abstract())),
        c)

// declares a class with no subclasses (apply to things)
final(c:class) : any
 -> let n := c.open in
       (if (n = default()) error("[125] ephemetral classes cannot be final")
        else if (size(c.subclass) != 0)
          error("[125] a class with subclasses cannot be final")
        else (write(open, c, final())),
        c)

// instantiation with and without a name
// new! is a method-less property that is managed by the compiler
new(self:class) : type[object glb member(self)]
  -> let o := (if (self.open <= 0) error("[105] cannot instantiate ~S", self),  // v3.2.44
               new!(self)) in
        (// if (self.open = open()) add!(instances, self, o),                // v3.2.34
         new_defaults(o,nil))                                                          // v3.2.26
        
new(self:class,%nom:symbol) : type[thing glb member(self)]
   -> let o := (if (self.open <= 0) error("[105] cannot instantiate ~S", self),  // v3.2.44
                new!(self,%nom)) in
       (new_defaults(o,nil) as thing)      // v3.2.26

// internal version => moved to Kernel
// mClaire/new!(self:class) :  type[object glb member(self)]
//  -> function!(new_object_class)
// mClaire/new!(self:class,%nom:symbol) :  type[thing glb member(self)]
//   -> function!(new_thing_class)

// the smallest super_set of two sets
// there is always any, so it always returns a class
meet(self:class,ens:class) : class
 -> (let l1 := self.ancestors,
         l2 := ens.ancestors,
         m := (if (length(l1) < length(l2)) length(l1) else length(l2)) in
       (while (l1[m] != l2[m]) m :- 1, l1[m] as class))

// fast inclusion method for lattice_sets (lattice order). The argument is
// either a lattice_set or {}
inherit?(self:class,ens:class) : boolean
 -> (let l := self.ancestors,
         n := length(ens.ancestors) in
       (n <= length(l) & l[n] = ens))


//------------- PROPERTY ---------------------------------------------------

// the two methods to access open(r)
// an abstract property is extensible and can receive new restrictions
abstract(p:property) : any
 -> let n := p.open in
       (if (n < 2) error("[127] ~S can no longer become abstract", p)
        else write(open, p, 3),
        p)

// a final property is completely defined and cannot receive a new restriction
// v3.2.04: the new value 4 will be used to represent (compiled but open)
final(r:relation) : void
 -> (case r
      (property (if (r.open <= 2)
                   (write(open, r, 1),                  // v3.2.04
                    put(domain, r,
                        Uall(list{ x.domain[1] | x in r.restrictions})),
                    put(range, r,
                        Uall(list{ x.range | x in r.restrictions}))))))


//------------- MODULES   --------------------------------------------------

// book-keeping for a module : based on parts/part_of hierarchy
// propagates uses declaration + register a new associated namespace.
close(self:module) : module
 -> (if (self != claire)
        (if known?(part_of, self)
            let sup := self.part_of in
              (parts(sup) :add self,
               for x in sup.uses
                 (if (not(x % self.uses) & x % module) uses(self) :add x))
         else trace(3, "---- note: ~S is a root module !\n", self)),
     Kernel/namespace(self),                     // CLAIRE4: name spaces are defined in Kernel
     self)

// note: dynamic modules are no longer supported
claire/get_symbol(m:module,self:string) : any -> function!(get_symbol_module)
claire/get_symbol(self:string) : any -> get_symbol(claire,self)

// *********************************************************************
// *   Part 3: System Methods                                          *
// *********************************************************************

// all these methods will be open-coded by the compiler
//get_stack(self:integer) : any -> get_stack(self)
//put_stack(self:integer,x:any) : any -> put_stack(self, x)
//push!(self:meta_system,x:any) : void -> push!(self, x)

time_get() : integer -> function!(time_get_void)
time_read() : integer -> function!(time_read_void)
time_set() : void -> function!(time_set_void)
time_show() : void -> function!(time_show)
//stat() : void -> function!(claire_stat)
gensym(self:void) : symbol -> gensym("g")

// world management
store(l:list,n:integer,y:any) : any
 -> (if (n < 1 | n > length(l)) error("store @ list: ~S out of bounds for ~S",n,l)
     else store(l,n,y,true))

store(l:array,n:integer,y:any) : any 
 -> (if (n < 1 | n > length(l)) error("store @ array: ~S out of bounds for ~S",n,l)
     else store(l,n,y,true))

commit(n:integer) : void -> (while (n < world?()) commit())
backtrack(n:integer) : void -> (while (n < world?()) backtrack())
claire/world+ :: choice
claire/world- :: backtrack
claire/world-! :: commit

// allows to change the storage class
store(l:listargs) : any
  -> (for r in l
         case r (relation r.store? := true,
                 string let v := get_value(r) in
                          case v (global_variable v.store? := true)))


// *********************************************************************
// *   Part 4: Methods for Native entities                             *
// *********************************************************************

//------------------- STRING -----------------------------------------------
// length(self:string) : integer -> function!(length_string)  -> in Kernel

// make_function(self:string) : function -> function!(make_function_string)
symbol!(self:string) : symbol -> symbol!(self, claire)

// deprecated in claire 4 - do not use a string as a byte buffer
nth_get(s:string,n:integer,max:integer) : char
  -> (if (n <= max) s[n] else error("Buffer string access"))
nth_put(s:string,n:integer,c:char,max:integer) : void
  -> (if (n <= max) s[n] := c  else error("Buffer string access"))

// shell(self:string) : void -> function!(claire_shell)
claire/getenv(self:string) : string -> function!(getenv_string)
claire/get_value(self:string) : any -> function!(value_string)           // value @ string no longer supported in CLAIRE 4
claire/get_value(self:module,s:string) : any -> function!(value_module)  //  v3.2.14

// we keep the externC method name even if it now support go code
externC(s:string) : void -> error("cannot execute Go code: ~A", s)
externC(s:string,c:class) : type[member(c)] -> (error("cannot execute ~A",s), unknown)

// new in claire 4 - print with a constant number of char = truncate or extends
[princ(s:string,n:integer) : void
  -> let m := length(s) in
       (if (m > n) princ(substring(s,1,n))
        else (princ(s),
              for i in (m + 1 .. n) princ(' '))) ] 

//------------------- SYMBOL -----------------------------------------------
make_string(self:symbol) : string
 -> (print_in_string(), princ(self), end_of_string())
//princ(self:symbol) : any -> function!(princ_symbol)
self_print(self:symbol) : void -> printf("~A/~S", module!(self).name,string!(self))
//c_princ(self:symbol) : any -> function!(c_princ_symbol)
//gensym(self:string) : symbol -> function!(gensym_string, NEW_ALLOC)

//--------------------- INTEGER -----------------------------------------
+(self:integer,x:integer) : type[abstract_type(+, self, x)] -> (self + x)
-(self:integer,x:integer) : type[abstract_type(-, self, x)] -> (self - x)
//-(self:integer) : integer -> function!(ch_sign)

//float!(self:integer) : float -> function!(to_float)
//mod(self:integer,x:integer) : integer -> function!(mod_integer)
claire/--?(self:integer,x:integer) : set -> function!(sequence_integer)
--!(self:integer,x:integer) : list -> function!(list_integer)
// exit(self:integer) : void -> function!(CL_exit)
// less_code(n:integer,i:integer) : boolean -> function!(less_code_integer)
<<(x:integer,y:integer) : integer -> externC("(x << y)",integer)
>>(x:integer,y:integer) : integer -> externC("(x >> y)",integer)
and(x:integer,y:integer) : integer -> externC("(x & y)",integer)
or(x:integer,y:integer) : integer -> externC("(x | y)",integer)

// open-coded
<(self:integer,x:integer) : boolean -> (if (self < x) true else false)
<=(self:integer,x:integer) : boolean -> (if (self <= x) true else false)
>(self:integer,x:integer) : boolean -> (if (self > x) true else false)
nth(self:integer,y:integer) : boolean -> (if self[y] true else false)

claire/abs(x:integer) : integer -> (if (x >= 0) x else -(x))
claire/random(a:integer,b:integer) : integer -> (a + random(b + 1 - a))

// used by the logic
factor?(x:integer,y:integer) : boolean -> ((x mod y) = 0)
divide?(x:integer,y:integer) : boolean -> ((y mod x) = 0)
Id(x:any) : type[x] -> x

//------------------------ FLOAT ---------------------------------------------
self_print(self:float) : void -> function!(print_float)
+(self:float,x:float) : float ->  (self + x)  
-(self:float,x:float) : float -> (self - x) 
*(self:float,x:float) : float -> (self * x) 
/(self:float,x:float) : float -> unsafe(self / x)     // v4.0.7   -> cross compiled function without error 


-(self:float) : float -> (-1.0 * self)

string!(self:float) : string -> (print_in_string(), princ(self), end_of_string())
princ(self:float,i:integer) : void -> function!(print_format_float)        // v3.3.42

claire/abs(x:float) : float -> (if (x >= 0.0) x else -(x))

// the pF is my ugly duckling :) -------------------------------------------
// float print is now standard in v3.4.42 (princ(float_integer)  but this is still a cuter print ...
mClaire/printFDigit(x:float,i:integer) : void        // prinf i numbers
  -> (if (x < 0.0) (princ("-"), printFDigit(-(x),i))
      else let frac := x - float!(integer!(x + 1e-10)) + 1e-10 in
         printf("~A.~I", integer!(x + 1e-10),
                printFDigit(integer!(frac * (10.0 ^ float!(i))),i)))

// print the first i digits of an integer
mClaire/printFDigit(x:integer,i:integer) : void
  -> (if (i > 0) let f := 10 ^ (i - 1), d := x / f in
                   (princ(d), if (i > 1) printFDigit(x mod f, i - 1)))

//--------- BAG --------------------------------------------------------
// in CLAIRE 4, bag is no longer a concrete data type, hence bag methods must be divided between sets and lists

// length(self:bag ) : integer -> length(self)  => becomes part of go

// in CLAIRE 4, we can tell the compiler that the list is a list<object> (optimization purpose)
mClaire/nth_object(self:list,n:integer) : any -> nth_object(self,n)

/* nth only exists for lists
nth(self:list,x:integer) : type[(if (self % tuple & unique?(x)) self[the(x)]   // v3.3.22
                                else member(self))]
 -> (if (x > 0 & x <= length(self)) self[x]
     else error("[41] nth[~S] out of scope for ~S", x, self)) */
    
// nth_get(self:list,x:integer) : any -> nth_get(self, x)

// new in claire 4: tells the compiler that range check is required + EID optimized
nth_write(self:list,i:integer,v:any) : any 
  -> (if (v % of(self)) self[i] := v
      else system_error(index = 17,arg = v, value = of(self)))

// CLAIRE 4 duplication: define min/max for sets first
min(f:method,self:set) : type[member(self)]     // v3.1.08
 -> (if (size(self) != 0)
        let x:any := unknown in
                 (for y in self (if (x = unknown | funcall(f, y, x) as boolean) x := y), x)
     else error("[183] min of empty set is undefined"))

max(f:method,self:set) : type[member(self)]     // v3.1.08
 -> (if (size(self) != 0)
        let x:any := unknown in
                  (for y in self (if (x = unknown | not(funcall(f, y, x))) x := y), x)
     else error("[183] max of empty set is undefined"))

// CLAIRE 4 : optimize for lists
min(f:method,self:list) : type[member(self)]     // v3.1.08
 -> let n := length(self) in
      (if (n != 0)
        let x := self[1] in
                 (for i in (2 .. n) (if (funcall(f, nth_get(self,i), x) as boolean) x := nth_get(self,i)), x)
    else error("[183] min of empty list is undefined"))

max(f:method,self:list) : type[member(self)]     // v3.1.08
 -> let n := length(self) in
      (if (n != 0)
          let x := self[1] in
                  (for i in (2 .. n) (if not(funcall(f, nth_get(self,i), x)) x := nth_get(self,i)), x)
     else error("[183] max of empty list is undefined"))

// CLAIRE4 : /+ is native for list

// new for claire 3.4
random(self:list) : type[member(self)] -> let n := length(self) in self[1 + random(n)]

//--------- LIST --------------------------------------------------------
// last element of a list
last(self:list) : type[member(self)] 
  -> (if (length(self) > 0) self[length(self)] else error("[41] car(nil) is undefined"))
  
// remove the last element
rmlast(self:list) : list  -> (nth-(self,length(self)), self)

// the old LISP method
car(self:list) : any
 -> (if (length(self) > 0) self[1] else error("[41] car(nil) is undefined"))

// this method sorts a list according to an order
sort(f:method,self:list) : list
 -> (quicksort(self, f, 1, length(self)), self)

// v3.0.38: upgrade the quicksort algorithm with a better pivot selection cf.bag.cpp
// this is also proposed as a macro: cf. file.cl
quicksort(self:list,f:method,n:integer,m:integer) : void
 -> (if (m > n)
        let x := self[n] in
          (if (m = (n + 1))
              (if (funcall(f, self[m], x) as boolean)
                  (self[n] := self[m], self[m] := x))
           else let p := (m + n) >> 1, q := n in   // new: p is pivot's position
                  (x := self[p],
                   if (p != n) self[p] := self[n],
                   for p in ((n + 1) .. m)
                     (if (funcall(f, self[p], x) as boolean)
                         (self[n] := self[p],
                          n := n + 1,
                          if (p > n) self[p] := self[n])),
                   self[n] := x,
                   quicksort(self, f, q, n - 1),
                   quicksort(self, f, n + 1, m))))

 // destructive method that build the powerset
build_powerset(self:list) : set
 -> (if (length(self) != 0)
        let x := self[1],
            l1 := build_powerset(self << 1),
            l2 := l1 in
          (for y:set in l1 l2 :add (set(x) /+ y), l2)
     else {{}})

// skip 
<<(x:list,y:integer) : list -> function!(skip_list)

// new and useful (v3.1.06) - create a list with n replication of the default value - deprecated
claire/make_copy_list(n:integer,d:any) : list
  -> let l := make_list(n,d) in
       (case d (list  for i in (1 .. n) l[i] := copy(d)),
        l)

// new version : create a typed list for integer or floats
claire/typed_copy_list(t:type,n:integer,d:any) : list
  -> (if (t <= integer) make_list(n,integer,d as integer)
     else if (t <= float) make_list(n,float,d as float)
     else let l := make_list(n,d) in
       (case d (bag for i in (1 .. n) l[i] := copy(d)),
        l))

//----------------------  SET  ---------------------------------------------
difference(self:set,x:set) : set -> { y in self | not(contain?(x, y))}

//----------------------  TYPE ---------------------------------------------
@(self:type,p:property) : type -> function!(_at_type)

//--------- ARRAY --------------------------------------------------------

/* defined in Kernel in v4.0.5
nth=(self:array,x:integer,y:any) : void
 -> (if not(y % of(self)) error("type mismatch for array update ~S, ~S",y,self)
     else if (x > 0 & x <= length(self)) nth_put(self,x,y)
     else error("nth[~S] out of scope for ~S", x, self)) */

self_print(self:array) : void -> printf("array<~S>[~A]",of(self),length(self))
 
//---------------------- CHAR --------------------------------------------
self_print(self:char) : void -> printf("'~A'", self)
<=(c1:char,c2:char) : boolean -> (integer!(c1) <= integer!(c2))

// --------------------- BOOL -----------------------------------------------
claire/random(b:boolean) : boolean
  -> (if b (random(10000) >= 5000) else false)




