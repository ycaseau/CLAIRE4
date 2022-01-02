//+-------------------------------------------------------------+
//| CLAIRE                                                      |
//| otool.cl                                                    |
//| Copyright (C) 1994 - 2021 Yves Caseau. All Rights Reserved  |
//| cf. copyright info in file object.cl: about()               |
//+-------------------------------------------------------------+

//-------------------------------------------------------------------
// this file contains the auxiliairy methods for the source optimizer
//-----------------------------------------------------------------

// ******************************************************************
// *  Table of contents                                             *
// *    Part 1: New Instructions & associated stuff                 *
// *    Part 2: Optimizer Warnings                                  *
// *    Part 3: Type Handling                                       *
// *    Part 4: Miscellaneous                                       *
// ******************************************************************

// ******************************************************************
// *    Part 1: New Instructions & associated stuff                 *
// ******************************************************************

/* three special instructions


// to_CL(x,s) produces a CLAIRE oid from an external thing of sort s
self_print(self:Compile/to_CL) : void -> printf("CL{~S}:~S", self.arg,self.set_arg)

c_type(self:Compile/to_CL) : type
 -> sort_abstract!(c_type(get(arg, self)))   // v2.4.06 (was any)

Compile/c_gc?(self:Compile/to_CL) : boolean
 -> (not(gcsafe?(self.set_arg)) &
      (self.set_arg = float |                  // v3.3.3 ! float must be protected when converted to OID
       self.set_arg Core/<=t import) |         // v3.00.10: similar rule = protect import if converted to OID
       Compile/c_gc?(self.arg))                // default: only protect if content if GC-fragile 
     

// to_C(x) produces an external thing from a CLAIRE oid of sort s.
self_print(self:Compile/to_C) : void -> printf("C{~S}:~S", self.arg,self.set_arg)

// new: (to retrofit in v2.5) an object pointer may be dangerous
Compile/c_gc?(self:Compile/to_C) : boolean
  -> (not(gcsafe?(self.set_arg)) & Compile/c_gc?(self.arg) &
      (self.set_arg Core/<=t object | self.set_arg = string) )      // v3.00.30 !! + v3.3.34

[c_type(self:Compile/to_C) : type
 -> self.set_arg glb ptype(c_type(get(arg, self))) ]                 // v3.2.28 (smart)
 */

// this is a same-sort (object) casting from one class to another because of the
// stupidity of the target type system
// its use is linked to stupid_t(x)
self_print(self:Compile/C_cast) : void
 -> printf("<~S:~S>", self.arg, self.set_arg)
Compile/c_gc?(self:Compile/C_cast) : boolean -> Compile/c_gc?(self.arg)
c_type(self:C_cast) : type -> self.set_arg    // v3.0 : better safe
c_code(self:Compile/C_cast,s:class) : any
 -> (if (s inherit? object)
        Compile/C_cast(arg = c_code(self.arg, s), set_arg = self.set_arg)
     else c_code(self.arg, s))

// we need a new type to express powerful Iterate rules
// Note: Patterns require the compiler !

self_print(self:Pattern) : void
 -> printf("~S[tuple(~A)]", self.selector, self.arg)

%(x:any,y:Pattern) : boolean
 -> (case x
      (Call (x.selector = y.selector &
             tmatch?(list{ c_type(z) | z in x.args}, y.arg)),
       any false))

// this is very lazy, we could do better
glb(x:Pattern,y:type_expression) : type_expression
  -> (case y (Pattern (if (x.selector = y.selector) Pattern(selector = x.selector,
                                                            arg = glb(x.arg,y.arg))
                      else {}),
             any {}))

// extension of <= for Patterns
[less?(x:Pattern,y:type_expression) : boolean
 -> case y
      (Pattern (x.selector = y.selector & length(x.arg) = length(y.arg) &        // v3.2.18
                forall(i in (1 .. length(x.arg)) | =type?(x.arg[i],y.arg[i]))),
       any Call Core/<=t y) ]

[less?(x:type_expression, y:Pattern) : boolean
 -> case x
      (set forall( z in x | z % y),
       Pattern (x.selector = y.selector & length(x.arg) = length(y.arg) &
                forall(i in (1 .. length(x.arg)) | =type?(x.arg[i],y.arg[i]))),
       any false) ]

// v0.03 must return a type
nth(p:property,x:tuple) : Pattern -> Pattern(selector = p, arg = list!(x))   

// ******************************************************************
// *    Part 2: Optimizer Warnings                                  *
// ******************************************************************

// unified warning
Compile/warn()  : void
 -> (compiler.n_warnings :+ 1,
     if known?(OPT.in_method) trace(1,"---- WARNING[in ~S, line ~A]: ",OPT.in_method,n_line())
     else trace(1,"---- WARNING[lien ~A]: ",n_line()))

Compile/Cerror(s:string,l:listargs) : {}
  -> (printf("---- Compiler Error[in ~S]:\n", OPT.in_method),
      printf("---- file read up to line ~A\n", n_line()),
      general_error(Kernel/cause = s, arg = l))

// a note
Compile/notice() : void
 -> (compiler.n_notes :+ 1,
     if known?(OPT.in_method) trace(2,"---- note[in ~S]: ",OPT.in_method)
     else trace(2,"---- note: "))

// Warning : compiling is impossible, wrong selector
[c_warn(self:Call,%type:any) : any
 -> let s := self.selector in
       (if (%type = void)  Cerror("[205] message ~S sent to void object", self)
        else if (not(s.restrictions) & not(s % OPT.ignore))
           (warn(),trace(1,"the property ~S is undefined [255]\n", s))
        else if (not(s % OPT.ignore) & (s.open <= 1 | s.open = 4) &
                 (case %type (list class!(%type[1]).open != 3)))
            (warn(), trace(1,"wrongly typed message ~S (~S) [256]\n", self, %type))
        else if compiler.optimize?
            (notice(), trace(3,"poorly typed message ~S [~S]\n", self, %type)),   // v3.3 poor opt. notice
        open_message(self.selector, self.args)) ]

[c_warn(self:Super,%type:any) : any
 -> let s := self.selector in
       (if (%type = void)  Cerror("[205] message ~S sent to void object", self)
        else if not(s.restrictions)
           (warn(),trace(1, "the property ~S is undefined [255]\n", s))
        else if (not(s % OPT.ignore) & s.open <= 1)
           trace(3,"---- note: wrongly typed message ~S [~S]\n", self, %type),
        let m := open_message(self.selector, self.args) in
          Super(selector = m.selector, cast_to = self.cast_to, args = m.args))  ]

// a message cannot be compiled into efficient code
// here the property does not allow the compilation and we want to see it
[c_warn(self:property,l:list,%type:list) : any
 -> if (self.open <= 1 & not(self % OPT.ignore) &  compiler.safety >= 2)
        trace(4, "---- note: poor type matching with ~S(~S) [~S]\n", self, l, %type),
    open_message(self, l) ]

// a variable should not be abused ! Either it is a true error or it is
// simply dangerous. The result is the value to be used (either x or
// ckeck_in(x,range(oself))
[c_warn(self:Variable,x:any,y:type) : any
 ->  if (self.index = -1) x   // self is a special variable created in a case expansion
     else if not(y ^ self.range)
        (if (compiler.safety >= 2)
            (warn(), trace(1,"~S of type ~S is put in the variable ~S:~S [257a]\n",
                            x, y, self, self.range))
         else Cerror("[212] the value ~S of type ~S cannot be placed in the variable ~S:~S",
                    x, y, self, self.range))
     else if (compiler.safety <= 1 | not(sort=(osort(self.range), osort(y))))
        (warn(),trace(1,"~S of type ~S is put in the variable ~S:~S (~A) [257b]\n",
                      x, y, self, self.range, self.index)),
     if (compiler.safety <= 1 & not(y <= self.range)) c_code(x, any) // forces a check in 
     else x ]


// ******************************************************************
// *    Part 3: Type Handling                                       *
// ******************************************************************

// we use  {any U type} to represent the change of sort  (to any)
//         {} U (c U t) to represent a change of psort   (to c)
// e.g.: (any U class) = class stored as an OID

 // tests if two sorts are similar
 // the compiler.overflow? test
Compile/sort=(c:class,c2:class) : any
 -> (if (c inherit? object) c2 inherit? object
     else (c = c2 |
          (not(compiler.overflow?)
           & (c = any & c2 = integer) | (c = integer & c2 = any))))

// give the "precise sort", i.e., a class under object is a sort
[Compile/psort(x:any) : class
 -> let c := class!(x) in
      (if (c inherit? object) c
       else sort!(c)) ]

// gives the "optimizer sort", which is one of
// any, object, float, X <= import,
[Compile/osort(x:any) : class
 -> sort!(class!(x)) ]

sort(x:Variable) : class
 -> (let r := x.range in
       (if (case r (Union r.Kernel/t1 = {}))
           Compile/psort(r.Kernel/t2.Kernel/t2)
        else Compile/psort(r)))

// this is a very stupid type inference that mimicks the go compiler - defined in pretty.cl with CLAIRE4
// it returns a class
[Compile/stupid_t(self:any) : class 
    -> static_type(self)]

// comparison
[Compile/stupid_t(self:any,x:any) : boolean
 ->  let c1 := Compile/stupid_t(self),
         c2 := Compile/stupid_t(x) in
       (c1 != any & c1 = c2) ]

// an extended type is of the kind (t U {unknown})
// CLAIRE4: got rid of optUnion
[extended?(self:type) : boolean
 -> case self
      (Union (self.Kernel/t2 % set & size(self.Kernel/t2) = 1 & the(self.Kernel/t2) = unknown),
       any false) ]

// creates an extended type (v0.02) that can be checked easily 
// CLAIRE4: we removed the syntactic marker optUnion for (X U {unknown})
// used in ocall and ocontrol
[extends(x:type) : type
   -> Union(Kernel/t1 = x, Kernel/t2 = {unknown})]

// a sort abstraction is the special union any U t, which is known to represent t by
// the type system (used for variables only) but tells the compiler that the sort is any
[sort_abstract!(x:type) : type
  -> if (sort!(x) != any & sort!(x) != integer & sort!(x) != float)
        Union(Kernel/t1 = any, Kernel/t2 = x)        // sort! on float is special..
     else x ]   
                                                                                            // v3.00.05
sort_abstract?(x:type) : boolean -> (case x (Union (x.Kernel/t1 = any), any false))

// since we introduce some fuzziness with types (any U t), we need a way to get
// the precise type t back
ptype(x:type) : type
 -> (case x (Union (if (x.Kernel/t1 = any) x.Kernel/t2 else x), any x))

// v3.1.06: member -> always apply to a ptype
pmember(x:type) : type -> member(ptype(x))

// transform an instruction representing a set into an instruction
// representing an enumeration
[enumerate_code(self:any,%t:type) : any
 -> if (ptype(%t) <= list | ptype(%t) <= set | ptype(%t) <= tuple)
       c_strict_code(self, class!(ptype(%t)))              // v3.2.01
    else (if compiler.optimize? 
             (notice(), trace(3,"explicit enmeration of ~S\n", self)),  // v3.3
          c_code_method(Core/enumerate @ any, list(self), list(%t))) ]


// range inference for a "for" structure: y is the new type and ts is the type of
// the collection structure. Note that except for the case of float arrays, the
// sort of the collection is assumed to be any or integer (thus we "correct" the
// type inference with sort_abstract)
[range_infers_for(self:Variable,y:type,ts:type) : any
 ->  if unknown?(range, self)
        (//[5] infer type ~S for ~S // y, self,
         if (y % Interval) y := integer,              // v3.1.06
         put(range, self, y))
     else if (not(y <= self.range) & compiler.safety <= 1)
        (if not((y & self.range))
            (warn(), trace(1, "range of variable in ~S is wrong [258]\n", self)))]
        // v3.1.06: remove complains because it traps the compiler's own inferences
        // to reintroduce, we need to distinguish between user and compiler
        // types for iteration variables !
     // if (sort(self) != any & (sort(self) != integer | compiler.overflow?) &
     //    not(ts <= array & y <= float))               // iteration of float array is a special case
     //  (//[5] protect original sort with ~S // sort_abstract!(self.range),
     //   put(range, self, sort_abstract!(self.range))) ]


// variable range inference, how to guess a type from the value ...
[range_infers(self:Variable,y:type) : any
 ->  if (unknown?(range, self) | extended?(self.range))
        (if (y % set) put(range, self, class!(y))
         else put(range, self, y)) ]


// temporary range inference for case, which may use a special form:
// {any U type} to represent the change of sort
// {} U (c U t) to represent a change of psort
[range_sets(self:any,y:type) : void
 -> case self (Variable put(range, self, y)) ]
    /*   (Variable (if sort=(osort(self.range), osort(y))
                   let c1 := psort(class!(self.range)) in
                     (if (c1 != psort(class!(y)))
                         put(range, self,
                             Union(Kernel/t1 = {},
                                   Kernel/t2 =  Union(Kernel/t1 = c1, Kernel/t2 = y)))
                      else put(range, self, y))
                 else if (osort(self.range) = any)
                   put(range, self, sort_abstract!(y)))) ] */

// the srange of a method = class!(range)
[c_srange(m:method) : class
 -> (last(m.srange) as class) ]

// v3.3 some of the global variables are compiled with a native var approach
// we require the range to be safe, no backtrack & local global var
[Compile/nativeVar?(x:global_variable) : boolean
  -> (compiler.optimize? & x.Kernel/store? = false &              // v3.3.04: only when optimized
      x.name.module! = defined(x.name)) ]

// v3.3 finds the possible return type of a block (within a loop)
// it returns a class for the time being ...
[Compile/return_type(self:any) : type
 -> case self
      (// to_C return_type(self.arg),
       Let return_type(self.arg),
       Do let x := {} in (for y in self.args x :^ return_type(y), x),
       If return_type(self.arg) ^ return_type(self.other),
       Return c_type(self.arg),
       Case let x := {} in (for y in self.args x :^ return_type(y), x),
       Handle return_type(self.arg),
       any {}) ]


// compiling a type expression --------------------------------------------
//
// creates the functional code that produce the code by evaluation
// note this is expensive -> we should encourage the use of global variables
c_code(self:(type_operator U Reference U Pattern),s:class) : any 
   -> c_code(self_code(self), s)

// to check - seems OK for 3.2 !
[self_code(self:subtype) : any
 -> Call(nth, list(self.arg, c_code(self.Kernel/t1,type)))]

// create a Param. Optimized in v3.2.28 for list<X>
[self_code(self:Param) : any
 -> if  (length(self.params) = 1 &  self.params[1] = of & self.args[1] % set)
        Call(Core/param!, list(self.arg, c_code(the(self.args[1]),type)))
    else Call(nth, list(self.arg, self.params,
                        list{c_code(y,type) | y in self.args})) ]

[self_code(self:Union) : any
 ->  Call(U,list( c_code(self.Kernel/t1,type), c_code(self.Kernel/t2,type))) ]

[self_code(self:Interval) : any
 -> Call(..,list(self.arg1,self.arg2)) ]

[self_code(self:Reference) : any
 -> Call(Reference!,list(self.args, self.Kernel/index)) ]

// compilation of a Pattern
self_code(self:Pattern) : any
 -> (if compiler.inline?
        Call(nth, list(self.selector,Tuple(args = self.arg)))
                    //      Call(tuple!, list(List(args = self.arg)))))
     else Call)


//-------------- membership compiling -------------------------------

// membership to a class : for final and closed classes => x.isa = c !
[member_code(self:class,x:any) : any
 -> let %xt := Call((if (c_type(x) <= object) isa else owner),list(x)) in
         (if ((self.open <= -1 | self.open = 1) & not(self.subclass))
              c_code(Call(=, list(self,%xt)))
          else c_code(Call(inherit?, list(%xt, self)))) ]

[member_code(self:type_operator,x:any) : any
 -> Call_method2(arg = (% @ list(any,any)),
                 args = list(c_code(x, any), c_code(self, any))) ]

[member_code(self:Union,x:any) : any
 -> Or(args = list(member_code(self.Kernel/t1, x),
                   member_code(self.Kernel/t2, x))) ]

[member_code(self:Interval,x:any) : any
 -> c_code(And(args = list(Call( >=, list(x, self.Kernel/arg1)),
                           Call( <=, list(x, self.Kernel/arg2)))),
           any) ]

[member_code(self:Param,x:any) : any
 -> c_code(And(args =
                 (list(Call(%, list(x, self.arg))) /+
                  list{ Call( %, list(Call(self.params[i],list(x)),self.args[i])) |
                        i in (1 .. length(self.params))})),
           any)]

// v3.3.14: specialized code for tuple
[member_code(self:tuple,x:any) : any
 -> if (x % Tuple)
       (if (length(x.args) != length(self)) false
        else c_code(And(args = list{Call(%, list(x.args[i], self[i]))  |
                                    i in (1 .. length(self)) }), any))
    else c_code_method(% @ list(any,any), list(x,self), list(any,any)) ]


[member_code(self:any,x:any) : any
 -> LDEF := nil,
    let %type := list(c_type(x), c_type(self)),
         r := Language/extract_pattern(self, nil) in
       (if (r = unknown | self = object | (case self (global_variable self.range))) // 2.4.06
           c_code_method(% @ %type, list(x, self), %type)
        else member_code(r, x)) ]

// membership optimization though inline definition of %
%(x:any,y:..[tuple(any, any)]) : boolean
  => (x <= eval(y.args[2]) & eval(y.args[1]) <= x)

%(x:any,y:but[tuple(any, any)]) : boolean
 => (x % eval(y.args[1]) & x != eval(y.args[2]))

// ******************************************************************
// *    Part 4: Miscellaneous                                       *
// ******************************************************************

// ------- variables ------------------------------------------------

[Compile/Variable!(s:symbol,n:integer,t:any) : Variable
 -> case t (type Variable(Kernel/pname = s, Kernel/index = n, range = t),
            any Variable(Kernel/pname = s, Kernel/index = n))]

get_indexed(c:class) : list -> c.slots

// simple C operations that can be duplicated at no cost {+, -, /, *}

// tells if an expression is a go simply designated object
[designated?(self:any) : boolean
 ->  self % thing | self % Variable | self % integer | self % boolean |
     self = nil | self = {} | self = unknown | self % float |
     (case self
       (Call let x := c_code(self) in
               ((not(x % Call) & designated?(x)) |
                 self.selector = mClaire/get_stack),
        Call_slot designated?(self.arg),
        Call_table designated?(self.arg),
        Call_array designated?(self.arg),
        // to_protect (not(need_protect(self.arg)) & designated?(self.arg)),
        Call_method ((self.arg.selector % OPT.simple_operations |
                     (self.arg = (unsafe @ any)) |   // v4 unsafe(..) is a just a marker
                     (self.arg = (nth @ list))) &    // v3.2.34: added nth
                    forall( y in self.args | designated?(y))),
        // to_CL designated?(self.arg),
        // to_C designated?(self.arg),
        any false)) ]


// OPT.non_identifiable_set: those sets who are identifiable (closure)
// set<class>{c in class | exists(c2 in c.descendants | c2.ident? = false)})

// equality is identity?
[Compile/identifiable?(self:any) : boolean
 -> (self = unknown |
       (let t := class!(c_type(self)) in not(t % OPT.non_identifiable_set))) ]

// inlinning ---------------------------------------------------------

// macro expansion of method self with argument list l
[c_inline(self:method,l:list,s:class) : any
 -> //[5] macroexpansion of ~S with method ~S // l,self,
    c_code(c_inline(self, l), s) ]
    
// apply the body of a macro definition
// notice that the name of the inner variables is changed except the second variable
// of iterate macros    
[c_inline(self:method,l:list) : any
 -> let f := self.formula,
        x := f.body,
        lbv := bound_variables(x),
        pv0 := (if (self.selector % {iterate, Iterate}) f.vars[2].pname
                else class.name) in
       (x := Language/instruction_copy(x),
        //[5] c_inline(~S) on ~S: ~S is bound : ~S // self,l,lbv,x,
        for v in lbv
          let v2 := Variable(pname = (if (v.pname = pv0) pv0
                                      else gensym()),     // v3.2.01, was (n :+ 1, v.pname /+ "_C_" /+ string!(n))),
                             Kernel/index = 1000) in      // force to be a local !
           (put(range,v2, get(range, v)),
            x := Language/substitution(x, v, v2)),
        OPT.max_vars :+ length(lbv),
        //[5] substitute f.vars = ~S with l = ~S // f.vars, l,
        c_substitution(x, f.vars, l, false)) ]


// returns the macro expanded code if a macro is involved and nil otherwise
[c_inline_arg?(self:any) : any
 -> case self
      (Call let l := self.args,
                m := restriction!(self.selector, list{ c_type(x) | x in l}, true) in
              case m
               (method (if (m.inline? & c_inline?(m, l)) c_inline(m, l)),
                any nil),
       any c_inline_arg?(Call(selector = set!, args = list(self)))) ]

// substitute any variable with same name as x with the value val. val is an expression
// when the special form eval() is found, it is "evaluated"
// NEW: in v3.0.5 -> eval(x,C) evals only if x is actually a C
[c_substitution(self:any,lx:list[Variable],val:list,eval?:boolean) : any
 ->  case self
      (Variable let i := some(j in (1 .. length(lx)) |
                              self.Kernel/pname = lx[j].Kernel/pname) in
                  (if known?(i) val[i] else self),
       list (for i in (1 .. length(self))
              self[i] := c_substitution(self[i], lx, val, eval?),
            self),
       Call (if (self.selector = eval)                     // two patterns eval(e) or eval(e,R)
               c_substitution(self.args[1], lx, val,       // new in v3.0.5 !
                              (length(self.args) = 1 |     // bool = true means
                                 (length(self.args) = 2 &  // we evaluate !
                                  val[1] % self.args[2]))) // 3.2.56: the range test holds on the first arg
             else if eval?
               try apply(self.selector,
                         list{ c_substitution(y, lx, val, true) |
                           y in self.args})
               catch any
                 (//[0] a strange problem happens ~A // verbose(),
                  warn(),trace(1,"failed substitution: ~S",system.exception!),
                  c_substitution(self.args, lx, val, false), self)
          else (c_substitution(self.args, lx, val, false), self)),
       Instruction (for s in owner(self).slots
                      let y := get(s, self) in
                        put(s, self, c_substitution(y, lx, val, eval?)),
                    self),
       any self) ]

// needed
[eval(x:any,y:class) : any -> eval(x) ]

// returns the list of bound variables in a piece of code
[bound_variables(self:any) : list
 -> let l := list<any>() in
       (case self (Instruction_with_var l := list<any>((self as Instruction_with_var).var)),
        case self
         (Variable nil,
          Instruction for s in self.isa.slots
                       l :add* bound_variables(get(s, self)),
          list for x in self l :add* bound_variables(x)),
        l) ]

// we must recognize true boolean ! coercion
[c_boolean(x:any) : any
  -> let tx:type := c_type(x), ptx:type := ptype(tx) in
       (if (ptx <= boolean)
           (case x (Call (if (x.selector = not & (ptype(c_type(x.args[1])) != boolean))
                             x := Call(!=,list(Call(boolean!,list(x.args[1])),true)))),
            (if (tx <= boolean) c_strict_code(x,boolean)
             else c_code(x,boolean)))      // v3.3
        else if (tx <= list) c_code(Call(!=,list(Call(length,list(x)),0)))
        else c_code(Call(boolean!,list(x)))) ]





