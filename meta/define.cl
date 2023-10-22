//+-------------------------------------------------------------+
//| CLAIRE                                                      |
//| define.cl                                                   |
//| Copyright (C) 1994 - 2021 Yves Caseau. All Rights Reserved  |
//| cf. copyright info in file object.cl: about()               |
//+-------------------------------------------------------------+

// --------------------------------------------------------------
// this file contains all definition & instanciation instructions
//---------------------------------------------------------------

// **************************************************************************
// * Contents:                                                              *
// *     Part 1: Definition instructions (Defobj, Defclass, Defmethod ...)  *
// *     Part 2: the instantiation macros                                   *
// *     Part 3: the useful stuff                                           *
// *     Part 4: the other macros                                           *
// *     Part 5: OFTO for methods                                           *
// **************************************************************************

extract_item :: property()
function! :: property()

iClaire/LastComment:any := unknown
claire/NeedComment:boolean := false

// *********************************************************************
// *     Part 1: Definition                                            *
// *********************************************************************
// this is the basic class instantiation
//
Defclaire <: Complex_instruction()
Definition <: Defclaire(arg:class,args:list)

self_print(self:Definition) : any
 -> printf("~S(~I)", self.arg, printbox(self.args))

// CLAIRE 4: fast definition when no close nor facy slots
DefFast <: Definition() 

// ------------- named object definition ------------------------------
//
Defobj <: Definition(ident:symbol)

self_print(self:Defobj) : void
 -> (if (arg(self) = global_variable)
       let r:any := any, v := unknown in
          (for x:Call in self.args
             (if (x.args[1] = value) v := x.args[2]
              else if (x.args[1] = range) r := x.args[2]),
           if r printf("~A:~S := ~I",self.ident,r,printexp(v,false))
           else printf("~A :: ~I",self.ident,printexp(v,false)))
     else printf("~A :: ~S(~I)", self.ident, (self.arg as any),
            printbox(self.args)))

// ------------- class definition ------------------------------------
//
Defclass <: Defobj(params:list, forward?:boolean)

self_print(self:Defclass) : void
 -> (if unknown?(ident,self) print("<Defclass>")
     else printf("~A~I <: ~S(~I)", self.ident,
           (if self.params printf("[~I]", princ(self.params))),
           self.arg,
           (let l := self.args,
                n := length(l) in
              for i in (1 .. n)
                (if (i = 1) set_level()
                 else lbreak(),
                 case l[i]
                  (Vardef
                     ppvariable(l[i]),
                   any printf("~I = ~S", ppvariable(l[i].args[1]),
                              l[i].args[2])),
                 if (i < n) princ(",")))))

// -------------- method definition ----------------------------------
//
Defmethod <: Defclaire(arg:Call,set_arg:any,body:any,inline?:boolean)

self_print(self:Defmethod) : void
 -> (printf("~S(~I) : ~I~I ~A ~I ", self.arg.selector,
            (if known?(args,self.arg) ppvariable(self.arg.args)), // v3.3.24
            printexp(self.set_arg,false),
            (lbreak(), pretty.index :+ 4),
            (if self.inline? "=>" else "->"),
            printexp(self.body, false)),
     pretty.index :- 4)

// -------------- array definition -----------------------------------
Defarray <: Defmethod()

self_print(self:Defarray) : void
 -> (printf("~S[~I] : ~S~I := ~I ", self.arg.args[1],
            ppvariable(cdr(self.arg.args)), self.set_arg,
            (lbreak(), pretty.index :+ 4), printexp(self.body, false)),
     pretty.index :- 4)

// -------------- rule definition ------------------------------------
Defrule <: Defclaire(ident:symbol,args:list,arg:any,body:any)

self_print(self:Defrule) : void
 -> (printf("~A(~I) :: rule(~I ~S ~I=> ~S)", self.ident, ppvariable(self.args),
            lbreak(4), self.arg, lbreak(4), self.body),
     pretty.index :- 4)

Defvar <: Defclaire(ident:Variable,arg:any)

self_print(self:Defvar) : void
 -> printf("~I := ~I", ppvariable(self.ident), printexp(self.arg, false))

// *********************************************************************
// *     Part 2: the general instantiation macro                       *
// *********************************************************************
// creation of a new object
//
self_eval(self:Definition) : any
 -> (let %c := self.arg,
         %o := (if (%c.open <= 0) error("[105] cannot instantiate ~S", %c),  // 1:final()
                mClaire/new!(%c)) in                   // v3.2.26
        Core/new_defaults(%o,new_writes(%o, self.args)))

// fast definition : no inverse management + no "close" method
fast_definition?(c:class) : boolean
  ->  (c.open > 1 & not( (close @ c) % method) & 
         forall(s: slot in c.slots | 
                  unknown?(inverse,s.selector) & not(s.selector.store?) & unknown?(if_write,s.selector)))

// then the evaluation is simpler ! write_fast does the range checking (may return an error) 
self_eval(self:DefFast) : any  
  -> let %o := mClaire/new!(self.arg) in
     (for x:Call in self.args
       let p := make_a_property(x.args[1]) in write_fast(p,%o,eval(x.args[2])),
      %o)    

// for a fast_definition, simpler eval
// the instantiation body is a sequence of words from which the initialization
// of the object must be built.
// CLAIRE4 : renamed complete(self:object,%l:list) to new_writes()
new_writes(self:object,%l:list) : list
 -> let lp := list() in
      (for x:Call in %l
       let p := make_a_property(x.args[1]),
           y := eval(x.args[2]),
           s := (p @ self.isa) in
         case s
          (slot (if (y = unknown) lp :add! p,
                 if not(y % s.range) range_is_wrong(s, y)
                 else mClaire/update(p, self, s.index, s.Core/srange, y)),
           any error("[106] the object ~S does not understand ~S", self, p)),
        lp)
         

// creation of a new named object
self_eval(self:Defobj) : any
 -> let %c := self.arg, %o:object := (unknown as object) in
       (if (%c.open <= 0) error("[105] cannot instantiate ~S", %c),
        if (%c inherit? thing)
            (%o := mClaire/new!(%c, self.ident),
             case %o (property (if (length(%o.restrictions) > 0)     // v3.2.58 : cause compiler problems !
                                  error("[188] the property ~S is already defined", %o))))
        else (%o := mClaire/new!(%c),
              if (%c.open = open()) add!(instances, %c, %o),
              put(self.ident, %o)),
        Core/new_defaults(%o, new_writes(%o, self.args)))

// creation of a new named object
// note that final() is the marker of a forward definition in CLAIRE4
self_eval(self:Defclass) : any
 -> (if (value(self.ident) % class &
                ( (value(self.ident) as class).open != final() |    // new in v2.5
                  self.arg != (value(self.ident) as class).superclass))
        error("[107] class re-definition is not valid: ~S",self)
     else let %o := class!(self.ident, self.arg) in
       (for x in self.args
          let v := unknown in
            (case x (Call (v := eval(x.args[2]), x := x.args[1])),    // Call(=)  x:t = v
             let rt := extract_type(x.range),
                 p := make_a_property(x.mClaire/pname) in
               (if (known?(v) & not(v % rt))
                   error("[108] default(~S) = ~S does not belong to ~S",x,v,rt),
                // new in CLAIRE 4 : check co-variant slot definition
                when s2 := some(sx in self.arg.slots | sx.selector = p) in
                   (if (p.open <= 0) error("[181] cannot overide a slot for a closed property ~S",p)
                    else if not(rt <= s2.range) error("[XXX] slot redefinition of ~S must be covariant, ~S is not a subtype",s2,rt)),
                // new in CLAIRE 4: float and integer slots must have a default value
                close(add_slot(%o, p,rt, getDefault(rt,v))))),             // index is computed by low level go method
        close(%o),
        if self.forward? write(open,%o,final()) 
        else if (%o.open = final()) write(open,%o,(self.arg as class).open),
        if (%o <= primitive)  %o.open := -1,                     // v3.2.40 avoid junk !
        put(params, %o, self.params),
        for p in self.params write(open, p as property, 0),
        attach_comment(%o),
        %o))

// we compute the proper default value (reused by compiler) - for int, float, sets and lists
[getDefault(rt:type,v:any) : any
  -> (if unknown?(v) 
        (if (rt <= integer) 0 
         else if (rt <= float)  0.0
         else if (rt <= set) empty_set(Core/of_extract(rt))
         else if (rt <= list) empty_list(Core/of_extract(rt))
         else unknown)
       else v) ]


// method definition
LDEF:any :: list() // v0.01
self_eval(self:Defmethod) : any
 -> (if not(self.arg % Call)
        error("[110] wrong signature definition ~S", self.arg),
     let p := make_a_property(self.arg.selector),
         l := self.arg.args,
         lv := (if (length(l) = 1 & l[1] = system)   list(Variable(mClaire/pname = symbol!("XfakeParameter"), range = void))
                else l),
         lp := extract_signature(lv),
         lrange := extract_range(self.set_arg, lv, LDEF),
         lbody := extract_status(self.body),
         m:method := add_method(p, lp, lrange[1], lbody[1], lbody[2]) in
       (if (p.open > 0 & p.open <= 1)       // interface methods are gone from CLAIRE 4
         when r := some(r in (p.restrictions but m) | (r.domain ^ m.domain)) in
           // error("[186] conflict between ~S and ~S is not allowed for ~S", m,r,p),
           trace(1,"--- WARNING ! [186] conflict between ~S and ~S is dangerous since ~S is closed\n", m,r,p), // v3.2.06
        LDEF := list<any>(),
        if (lbody[3] != body) (if jito?() trace(3,"---- jito for ~S\n",m),
                               put(formula, m, jito(lambda!(lv, lbody[3])))),
        if (length(lrange) > 1) put(Kernel/typing, m, lrange[2]),
        m.inline? := self.inline?,
        // write(inline?, m, self.inline?),
        attach_comment(m),
        close(m),                      // v4: must add close => insert_definition is not in Kernel
        if (p = close & not(m.range <= domain!(m)))                     // v3.2.01
          error("[184] the close method ~S has a wrong range",m),
        m))

// v3.2.24 : -1 : final
(%.open := -1,
 >=.open := -1,
 =.open := -1)

// attach a cute comment if needed ... to a defclass or a defmethod
attach_comment(x:any) : void
  -> (if (NeedComment & known?(LastComment))  x.comment := LastComment)


// returns the list of types AND modifies LDEF
[extract_signature(l:list) : list
 -> LDEF := list<any>(),
    let n := 0 in
       list{ (if not(v % Variable) error("[111] wrong typed argument ~S",v)  // v3.2.14
              else let p := extract_pattern(v.range, list(n)) in
                (n :+ 1,
                 if (p = unknown) error("[111] wrong typed argument ~S (~S)", v, v.range),
                 put(range, v, (type!@any(p))),
                 p)) |
         v:Variable in l} ]

// takes an <exp> that must belong to <type> and returns the CLAIRE type
// if LDEF is non-empty, it is used as a list of type variable and patterns
// may be returned. In addition, if the path list is non empty, new type
// variables may be defined. a syntax error will produce the unknown value
//
extract_pattern(x:any,path:list) : any
 -> (case x
      (class x,
       set let z := (if (size(x) = 1) extract_pattern(the(x), nil)) in
             case z
              (Reference let w:Reference := copy(z) in
                           (write(arg, w, true), w),
               any x),
       Tuple  let ltp := list{extract_pattern(z,path) | z in x.args} in
                    (if exists(y in ltp | unknown?(y)) unknown
                     else tuple!(ltp)),  // v3.0.56
       global_variable extract_pattern(x.value, path),
       Call let p := x.selector in
              (if (p = U)
                  let x1 := extract_pattern(x.args[1], nil),
                      x2 := extract_pattern(x.args[2], nil) in
                    (if (x1 = unknown | x2 = unknown) unknown else x1 U x2)  // v3.2.48
               else if (p = ^)
                  extract_pattern(x.args[1], nil) ^
                    extract_pattern(x.args[2], nil)
               else if (p = ..)
                  let v1 := extract_item(x.args[1], nil),
                      v2 := extract_item(x.args[2], nil) in
                    (if (v1 % integer & v2 % integer) (v1 .. v2) else unknown) //<yc> v3.0.02
               else if (p = nth) extract_pattern_nth(x.args, path)
               else if (p = *) when z := extract_pattern(x.args[1], path) in
                                 (z U {unknown})      // v3.1.14
                               else unknown
               else unknown),
       type x,
       unbound_symbol let s := extract_symbol(x),
                   v := some(z in LDEF | z.mClaire/pname = s) in
                 (if known?(v) v.range
                  else if (case path (list length(path) > 1))
                    let y := Reference!(cdr(path), path[1]),
                        v := Variable(mClaire/pname = s, range = (y as type)) in
                      (//[5] create a reference for ~S args=~S // s,y.args,
                       LDEF :add v, void)
                  else unknown),
       any unknown))

// takes an <exp> that must belong to <type> and returns the CLAIRE type
extract_type(x:any) : type_expression
 -> (LDEF := list<any>(),
     let r := extract_pattern(x, nil) in
       (if (r = unknown) error("[112] wrong type expression ~S", x)
        else r as type_expression))

// an item is an integer, a float, a symbol, a string or a type
extract_item(x:any,y:any) : any
 -> (if (x % ((((integer U float) U symbol) U string) U type)) x
     else if (x % global_variable) extract_item(x.value, y)
     else unknown)

// version for X[...] which is the most complex case - note the extensibility
// patch.
[extract_pattern_nth(l:list,path:list) : any
 -> let m := length(l), x := l[1] in
      (if (m = 1) let y := extract_pattern(l[1], nil) in
                   (if unknown?(y) unknown
                    else Param(arg = array, params = list(of),args = list(set(y))))
       else if (m = 2)
           (if (x % {list, set, subtype} | not(x % class))
               let y := extract_pattern(l[2], nil) in
                 try (if known?(y) l[1][y] else unknown) catch any unknown
            else unknown)
       else let l1 := l[2],               // list of properties pi
                l2 := l[3].args,          // list of expressions ei in C[pi:ei]
                l3 := list<any>() in
              (for n in (1 .. length(l1))
                  let y := l2[n] in
                    l3 :add
                      (case y
                        (Set
                           let v := extract_pattern(y.args[1],
                                                    copy(path) add l1[n]) in
                             (case v
                              ({void} any,
                               Reference  let z := copy(v) in  (put(arg, z, true), z),
                               any set((if known?(v) v else eval(y.args[1]))))),
                         any
                           extract_pattern(y, (if (length(path) != 0) path add l1[n])))),
                if (unknown % l3) unknown
                else Param(arg = x, params = l1, args = l3))) ]

// we perform some pre-processing on x[l] at reading time to make evaluation easier
[extract_class_call(self:class,l:list) : object
 ->  if (self % {list, set, subtype} & length(l) = 1 &   // v3.0.01
         (let y := l[1], z :=  extract_pattern(y,nil) in     // recognize the form list[t] and list<t> :: list[of = t]
            (case y (global_variable  y := l[1].value),
             (z % type | self = subtype |                // v3.0.48
             (case y (Call (y.selector != = | length(y.args) != 2),
                      Tuple true))))))
        Call(nth, self cons l) // ??  v3.2 ! list[t] -> subtype
     else if (self = lambda)
        (if (length(l) = 2 & (l[1] % Do | l[1] % Variable))
            let lv := (if (l[1] % Do) list{ v in l[1].args | v % Variable}
             else list(l[1])) in
              (extract_signature(lv), lambda!(lv, l[2]))
         else error("[113] Wrong lambda definition lambda[~S]", l))
     else let l1 := list<any>(),
              l2 := list<any>(),
              m := length(l) in
            (for n in (1 .. m)
               let y := l[n],
                   p := unknown,
                   v := unknown in
                 (case y
                   (Call (if not((y.selector = = & length(y.args) = 2))
                             error("[114] Wrong parametrization ~S", y),
                          p := make_a_property(y.args[1]),
                          v := Set(args = list(y.args[2]))),
                    Vardef
                      (p := make_a_property(y.mClaire/pname), v := y.range),
                    any (p := make_a_property(y), v := {})),
                  l1 :add p,
                  l2 :add v),
             Call( nth, (self cons list(l1, List(args = l2))))) ]

// extract the range (type and/or second-order function)
// lvar is the list of arguments that will serve as second-o. args
// ldef is the list of extra type variables that are defined in the sig.
[extract_range(x:any,lvar:list,ldef:list) : list
 -> if not((case x (Call (x.selector = nth & x.args[1] = type))))
        list(extract_type(x), {})
     else (//[5] extract the range from ~S with lval = ~S and ldedf = ~S // x,lvar,ldef,
           for v in ldef                           // transforms the reference in x into type expressions (using the paths)
             let r := v.range as Reference,
                 path := r.args,
                 n := length(path),
                 y := lvar[r.index + 1] in
               (for i in (1 .. n) y := Call(@, list(y, path[i])),
                x := substitution(x, v, Call(member, list(y)))),
           let lv2 := list<any>() in
             (for v in lvar
                let v2 := Variable(mClaire/pname = v.mClaire/pname, range = type) in
                  (lv2 :add v2, x := substitution(x, v, v2)),
              let lb := lambda!(lv2, x.args[2]),
                  ur := unknown in
                (//[5] extract range applies type lambda ~S to arg list ~S // lb, list{ v.range | v in lvar},
                 try ur := apply(lb, list{ v.range | v in lvar})
                 catch any (printf("The type expression ~S is not valid ... \n", x),
                            printf("context: lambda = ~S, lvars = ~S\n",lb,list{v.range | v in lvar}),
                            close(system.exception!)),
               if not(ur % type)
                  error("[115] the (resulting) range ~S is not a type", ur),
               list(ur, lb)))) ]

// create a bitvector from a list of flags
claire/bit_vector :: property()
claire/bit_vector(l:listargs) : integer
   -> let d := 0 in (for x in l d :+ ^2(x), d)

// parse the body and return (status, functional, body)
// the input is  body | (function!(f) | function!(f,s)) < | body> opt
// CLAIRE4: status is -1 : unknown, 0: no error, 1: an error may be thrown
//
extract_status(x:any) : list
 -> (let s := -1,
         f := (if (case x (Call x.selector = function!)) x else unknown) in
       (case x
         (And let y := x.args[1] in
                (if (case y (Call y.selector = function!))
                    (f := y, x := x.args[2])),
          Call (if (x.selector = function!) x := body),
          any nil),
        if known?(f)
           (x := body,
            if (length(f.args) > 1) s := 1
            else s := 0,
            f := imported_function(string!(extract_symbol(f.args[1])))),
        list(s, f, x)))

// new in CLAIRE4 : create a function with a syntactic marker ! for imported
claire/imported_function(s:string) : function 
  -> make_function("#" /+ s)        

// cleans a pattern into a type
type!(x:any) : type
 -> (case x
      (list list{ type!(y) | y in x},
       Param Param(arg = x.arg, params = x.params,
                   args = list{ type!(y) | y in x.args}),
       Reference any,
       type x,
       any any))        // for instance patterns

// creates a table
// to do in later versions: use an array if direct indexed access
// in the meanwhile, arrays of float should be used with care (indexed arrays)
//
self_eval(self:Defarray) : any
 -> (let a := (self.arg as Call).args,
         ar:table := mClaire/new!(table, extract_symbol(a[1])),
         v := (a[2] as Variable),
         s := extract_type(v.range),
         e := (let l := cdr(a),
                   b := lexical_index(self.body, l, 0,true) in
                 (if exists(va in l | occurrence(b, va) > 0) lambda!(l, b)
                  else self.body)),
         d := (case e (lambda unknown, any eval(self.body))) in
       (write(range, ar, extract_pattern(self.set_arg, nil)),
        if unknown?(range,ar) range_error(mClaire/cause = table, arg = self.set_arg, Core/wrong = type), // v3.3.18
        if (unknown?(d) & (ar.range <= integer | ar.range <= float)) 
            trace(0,"=== CLAIRE4 Warning: unknown not allowed as a default for table with range ~S\n ",ar.range),
        if known?(d) 
           (if not(d % ar.range)                  // v3.1.06
               range_error(mClaire/cause = ar,arg = d, Core/wrong = ar.range))
        else if (ar.range <= integer) d := 0
        else if (ar.range <= float) d := 0.0,                        // v4.0: unknown not allowed as a float or int
        put(range, v, s),
        attach_comment(ar),
        if (class!(ar.range) inherit? set) write(multivalued?, ar, true),
        if (length(a) = 2)
           (write(domain, ar, s),
            case s
             (Interval (write(params, ar, s.Core/arg1 - 1),      // v3.1.06 -> make_copy_list
                        write(Core/graph, ar, typed_copy_list(ar.range,size(s), d))),
              any (write(params, ar, any),
                   graph_init(ar))),
            case e
             (lambda for y in ar.domain ar[y] := funcall(e, y),
              any write(default, ar, d)))
        else let s2 := extract_type((a[3] as Variable).range) in
               (write(domain, ar, tuple!(list(s, s2))),
                put(range, (a[3] as Variable), s2),
                if (s % Interval & s2 % Interval)
                   (//[4] create a two dimensional array for ~S and ~S and d = ~S // s, s2,d,
                    write(params, ar, 
                          list<integer>(size(s2),
                               ((s.Core/arg1 * size(s2)) +
                                   s2.Core/arg1) -
                                 1)),
                    write(Core/graph, ar, typed_copy_list(ar.range, size(s) * size(s2), d)))
                else (//[4] create map dictionary for table ~S // ar, 
                      write(params, ar, any),
                      graph_init(ar)),
                //[4] --- start initialization for ~S --------------- // ar,
                case e
                 (lambda for y1 in s
                          for y2 in s2  ar[y1,y2] := funcall(e, y1, y2),
                  any write(default, ar, d))),
        ar))

// ------------------ NEW in v3.2 : definition of rules -----------------------
//

// a demon is a lambda with a name and a priority
demon <: lambda(Core/pname:symbol = symbol!("unamed"),
                priority:integer = 0,              // used by ClaireRules
                formula:lambda)

self_print(self:demon) : void -> princ(self.Core/pname)
funcall(self:demon,x:any,y:any) : any -> funcall(self.formula,x,y)
funcall(self:demon,x:any,y:any,z:any) : any -> funcall(self.formula,x,y,z)

// in the interpreted mode we store the list of demons using a table
claire/demons[r:relation] : list<demon> := list<demon>()  // list of relevant demons

claire/<- :: operation()
rule_object <: property()

// the last rule/axiom that was defined on each relation
// this is used to find when the relation may be compiled
relations[r:rule_object] : set := {}               // list of involved relations
last_rule[r:relation] : rule_object := unknown     // compile(ru) => may compile(r)
(write(inverse, relations, last_rule))

// evaluate a rule definition: create a new demon and, if needed, the if_write 
// function
eval_rule :: property(open = 3)
self_eval(self:Defrule) : any 
 -> (if (self.args[1] != system) eval_rule(self)   // hook for ClaireRules engine
     else let %condition := self.arg,
              ru := value(self.iClaire/ident) in        // name of the rule
       (put(isa, ru, rule_object), 
        add!(rule_object.instances,ru),
        let (R,lvar) := make_filter(%condition) in
         let d := make_demon(R,ru.name,
                            lvar,%condition,lexical_index(self.body,lvar,0,true)) in
         (if (R.if_write % function)
             error("cannot define a new rule on ~S which is closed", R),
          //[5] we have defined a demon ~S for ~S // d,R,
          demons[R] :add d,
          last_rule[R] := ru,
          if (length(demons[R]) = 1) eval_if_write(R),
          if (case R (property (length(R.restrictions) = 0)))
             eventMethod(R as property),
          ru)))

// an eventMethod is a property whose unique (?) restriction is a method
[eventMethod?(r:relation) : boolean
  -> case r (property forall(x in r.restrictions | not(x % slot))) ]
  
  
// check that condition is either a filter or the conjunction of a filter and a 
// condition
// a filter is R(x) := y | R(x) := (y <- z) | R(x) :add y | P(x,y)
// R(x) is x.r or A[x]
// the list of variable is of length 3 when R is mono-valued, whether we use a <- filter or a regular := 
[make_filter(cond:any) : tuple(relation,list[Variable])
  -> let c := (case cond (And cond.args[1], any cond)) in
       (//[5] make_filter : ~S (~S) // c, c.isa,
        if (case c (Call ((c.selector = write | c.selector = nth=) &
                          c.args[1] % relation)))
           let R := (c.args[1] as relation), 
               x := Variable(extract_symbol(c.args[2]),R.domain),
               y1 := c.args[3] in
             (if multivalued?(R as relation) 
                 error("[188] wrong event filter ~S for multi-valued relation",c,R),
              if (case y1 (Call (y1.selector = <-)))
                 tuple(R,list(x,Variable(extract_symbol(y1.args[1]),R.range),
                              Variable(extract_symbol(y1.args[2]),R.range)))
              else tuple(R,list(x,Variable(extract_symbol(y1),safeRange(R)),
                                  Variable(gensym(),safeRange(R)))))
        else if (case c (Call ((c.selector = add) & c.args[1] % relation)))
           let R := (c.args[1] as relation), 
               x := Variable(extract_symbol(c.args[2]),R.domain),
               y := Variable(extract_symbol(c.args[3]),R.range) in
             tuple(R,list(x,y))
        else if (case c (Call (length(c.args) = 2)))      // last case P(x,y) pattern
           let R := (c as Call).selector, 
               x := Variable(extract_symbol(c.args[1]),R.domain),
               y := Variable(extract_symbol(c.args[2]),R.range) in
             tuple(R,list(x,y)) 
        else error("[188] wrong event filter: ~S",c)) ]
       
  
// create a demon with lvar as list of variables
// notice that a demon may have 3 args if R is monovalued 
[make_demon(R:relation,n:symbol,lvar:list[Variable],cond:any,conc:any) : demon
   -> let x := lvar[1], y := lvar[2],
          %test:any := Call((if multivalued?(R) % else =), list(y, readCall(R,x))),
          %body:any := conc in
        (//[5] make a demon for ~S from ~S => ~S (name = ~S) // R, cond, conc,n,
          if (mClaire/trace!(if_write) > verbose())   // add a trace to the conclusion
           conc := Do(list(Call(format,list("--- trigger ~A(~S,~S)\n", 
                                             List(args = list(string!(n), x, y)))), 
                           conc)),
         %body := If(arg = conc),
         if eventMethod?(R)
            case cond (And %test := (if (length(cond.args) > 2)
                                        And(args = cdr(cond.args))
                                     else cond.args[2]),
                       any %body := conc)
         else case cond (And %test := And(args = list(%test) /+ cdr(cond.args))),        
         case %body (If %body.test := %test),
         demon(mClaire/pname = n,
               formula = lambda!(lvar,%body))) ]

// cute litle guy : create the read instruction both for a table and a property
[readCall(R:relation,x:any) : Call
  -> if (R % table) Call(get, list(R, x))                // v3.3.0
     else Call+(selector = R, args = list(x)) ]             

// a small brother
[putCall(R:relation,x:any,y:any) : Call
  -> if multivalued?(R) Call(add_value,list(R,x,y))
     else Call(put,list(R,x,y)) ]

// v3.3 : find the range when we read the current value     
[safeRange(x:relation) : type
  -> case x (property  (if forall(s in x.restrictions | (case s (slot s.default % s.range)))
                           x.range
                        else any),
             table  (if (x.default % x.range) x.range else any),
             any any) ]
          
// generate an if_write "daemon", only the first time, which uses
// the list in demons[R]
// the first step is to make the update (with inverse management)
eval_if_write(R:relation) : void
 -> let l := demons[R],
        lvar := l[1].formula.vars,  // list(x,y,?z) from 1st demon
        dv := Variable(gensym(),demon),
        l1 := list<any>(putCall(R,lvar[1],lvar[2])),
        l2 := list<any>(For(var = dv,
                            iClaire/set_arg = Call(nth,list(demons,R)),
                            arg = Call(funcall,list(dv) /+ lvar))) in
     (//[5] generate a if_write demon for ~S // R,
      for v in lvar put(range,v,class!(v.range)),
      if known?(inverse,R)
         (if not(multivalued?(R)) 
            l1 :add Call(Core/update-,list(R.inverse,lvar[3],lvar[1])),
          l1 :add putCall(R.inverse,lvar[2],lvar[1])),
      R.if_write := lambda!( list(lvar[1],lvar[2]),
         (if eventMethod?(R) Do(l2)
          else if multivalued?(R)
             If(test = Call(not,
                            list(Call(%,list(lvar[2],readCall(R,lvar[1]))))),
                arg = Do(l1 /+ l2))
          else Let(var = lvar[3],
                   value = readCall(R,lvar[1]),
                   arg = If(test = Call(!=,list(lvar[2],lvar[3])),
                            arg = Do(l1 /+ l2))))))
              
// create a restriction (method) that will trigger an event
eventMethod(p:property) : void
 -> let m:method := add_method(p, list(p.domain, p.range),void,0,unknown),
        %f := make_function(string!(p.name) /+ "_write") in
       (put(formula, m, p.if_write),              // how to execute a method ... 
        close(m),
        Kernel/set_arity(%f,2),
        put(functional, m, %f))                   // when we compile -> directly call the demon 


// **************************************************************************
// *     Part 5: JITO for methods                                           *
// **************************************************************************

// CLAIRE 4 reintroduced JITO : Just-In-Time Optimization
// we perform an on-the-fly optimization of lambdas through substitution (static calls)
// Jito(l:lambda) -> apply makeJito to the body (in place substitution)
[jito(self:any) : any
-> if not(jito?()) self                      // v4.0.6 : jito?() controls JITO
   else case self
      (list for x in self jito(x),
       Vardef (put(isa,self,Variable), self),
       lambda (jito(self.body), self),
       And jito(self.args),
       Or jito(self.args),
       Call (makeJito(self), true),
       Let letJito(self),
       Assign (if not(self.var % Variable) error("[101] ~S is not a variable but a ~S", self.var, owner(self.var)),    // moved this test from eval in v4.0 
               jito(self.arg)),
       Gassign (if (self.var.value % self.var.range) jito(self.arg)),    // watch out for unknown
       Do jito(self.args),
       If (jito(self.arg), jito(self.test), jito(self.other)),
       Iteration  let v := self.var, s := self.iClaire/set_arg,
                      o? := ((case s (Call s.selector = ..)) & unknown?(range,v)) in
                     (trace(3,"-- Iteration jito: ~S (~S)\n",self,static_type(self.iClaire/set_arg)),
                      if o? (put(range,v,integer), trace(3,"-- jito:put range ~S as integer\n",v)),
                      jito(s),
                      jito(self.arg),
                      if o? put(range,v,unknown)),
       While (jito(self.test), jito(self.arg)),
       Construct (trace(3,"-- Construct jito: ~S\n",self),
                  jito(self.args)),
       Exists (jito(self.iClaire/set_arg), jito(self.arg), jito(self.other)),
       Handle (if not(self.test % class) error("syntax: [try %S] must use a class",self.test), 
               jito(self.arg),
               jito(self.other)),
       Definition (if fast_definition?(self.arg) put(isa,self,DefFast)),        
       any false) ]

 // Let is special in CLAIRE4 : we implement the implicit typing found in the compiler = to infer
 // the type  from the value (when no range is given)
 // Note : this is doubtful ... 
 [letJito(self:Let) : any
   -> let v := self.var, x := self.value, 
          untyped:boolean := unknown?(range,v) in
         (trace(3,"Let Jito with var ~S => ~S\n",v,untyped),
          if untyped
            (if (x % List) let t := of(x) in 
                (if (t != {}) put(range,v,Core/param!(list,t)))
             else put(range,v,static_type(x)),
             trace(3,"--- let Jito ~S:~S (~S)\n",v,range(v),x)),
          jito(x), 
          jito(self.arg),
          if untyped put(range,v,unknown)) ]
            
// we optimize statically (Call(p) -> Call_method(m)) when
//   - only one restriction match 
//   - all domains are classes => class match
//   - the only one match is a compiled method
//   - the property is static (open = 1, vs extensible) and not too many restrictions
// note: the 12 hard limit is to avoid spending too much time with self_print or equivalent methods ... it is arbitrary
[makeJito(self:Call) : void
  -> jito(self.args),
     let p := self.selector, larg := self.args, n := length(larg), m := unknown in
       (if (p = write & (let p2 := self.args[1] in 
                  (case p2 (property (unknown?(inverse,p2) & not(p2.store?) & unknown?(if_write,p2)), 
                            any false))))
           (p := write_fast, self.selector := write_fast),
        if (p.open <= 1 & length(p.restrictions) <= 12 & 
            forall(x in p.restrictions | forall(t in x.domain | t % class)))
          let lt := list{static_type(x) | x in larg} in
            (trace(3,"-- call jito: ~S : ~S\n",self,lt),
             for x in p.mClaire/definition
                (if makeCallMatch(x,lt) (m := x, break()))),  
        if (case m (method known?(functional,m), any false))   // KEY: we only JITO compiled methods
            (put(isa,self, (if (n = 1) Call_method1 
                            else if (n = 2) Call_method2 
                            else if (n = 3) Call_method3 
                            else Call_method)),
             arg(self as Call_method) := (m as method))) ]


// tells if the restriction matches the type list lt : we know that the domain is made of classes
// only use for a compiled method, to help with debug
makeCallMatch(x:restriction,lt:list) : boolean
  ->  let n := length(lt), ld := x.domain in 
         (length(ld) = n & 
          forall(i in (1 .. n) | (lt[i] as class) <= (ld[i] as class)))   


// close some classes : final => no subclasses,  default() => ephemeral
// CLAIRE 4 : make sure that open statement for class are all here
(table.open := final(),
 class.open := final(),
 method.open := final(), 
 slot.open := final(), 
 boolean.open := -1,
 for x in Instruction.descendants (x.open := default()))  // instuctions are ephemeral

