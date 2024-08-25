//+-------------------------------------------------------------+
//| CLAIRE                                                      |
//| method.cl                                                   |
//| Copyright (C) 1994 - 2021 Yves Caseau. All Rights Reserved  |
//| cf. copyright info in file  object.cl: about()              |
//+-------------------------------------------------------------+

// ---------------------------------------------------------------------
// This file contains the reflective description of the most primitive
// CLAIRE kernel: the embryo of the class hierarchy and a set of methods
// to read/write objects and to evaluate messages.
// ---------------------------------------------------------------------

// CLAIRE4 revision : more objects are defined in Kernel

// *********************************************************************
// *  Table of contents                                                *
// *      Part 1: Lambda & Methods Evaluation                          *
// *      Part 2: Update methods                                       *
// *      Part 3: Management of definition(p)                          *
// *      Part 4: Matching Methods                                     *
// *********************************************************************

// catch what was missed in Kernel
(for s:slot in slot.instances close(s),
 for m:method in method.instances close(m))

// complete instanciation
[close(r:slot) : slot
  -> insert_definition(r.selector,r),
     r ]

[close(r:method) : method
  -> insert_definition(r.selector,r),
     r ]

// Claire 4: introduce the capacity to set the comment automatically at compile time
[attach(r:method,s:string) : method
  -> r.comment := "defined in file " /+ s, 
     close(r)]     

// *********************************************************************
// *      Part 1: Lambda & Methods Evaluation                          *
// *********************************************************************

// Lambda is now defined in Kernel

// explicit definition of the functions that are used in method [to avoid out_of-place implicit definitions]
execute :: property()
debug :: property()
eval_message :: property()
noeval_message :: property()
eval :: property()
call :: property()
self_eval :: property()
// read :: property() moved to Kernel
inlineok? :: property()
hold? :: property()
write :: property()
range_is_wrong :: property()
update+ :: property()
update- :: property()
add_value :: property()
known? :: property()
unknown?  :: property()
erase :: property()
set_range  :: property()
put_store :: property()
matching? :: property()
vmatch?  :: property()
tmatch? :: property()
find_which :: property()
claire/main :: property()

// the eval_message is a method that tells how to handle the message.
// it used to be distributed in CLAIRE (so that it was extensible) and each
// definition was called the behavior of a kind of restriction
// int? tells us if this is an interpreted message
eval_message(self:property,r:object,start:integer,int?:boolean) : any
 -> (if (r.isa = method)
       (if (system.debug! != -1) execute(r as method,start, int?)
        else let m:method := (r as method) in
          (if known?(formula,m)
            let retour := base!() in
              (set_base(start),
               stack_apply(m.formula.dimension),
               let val := eval(m.formula.body) in
                  (set_base(retour),                   // reset base to its entering value
                   set_index(start),                   // same for index (stack management)
                   val))
           else  stack_apply(m.functional, start, index!())))    // CLAIRE 4: use functional
    else if ( owner(r) = slot & index!() = (start + 1))
               let val := get(r as slot, get_stack(start) as object) in
                 (set_index(start),
                  let n := trace!() in
                    (if (n > 0 & trace!(self) + verbose() > 4 )
                      (put(trace!, system, 0),
                       printf("read: ~S(~S) = ~S\n",self,get_stack(start),val),
                       put(trace!, system, n))),
                  val)
    else noeval_message(self,start))


noeval_message(self:property, start:integer) : any
  -> let l := get_args(start) in
       (if (system.debug! != -1) push_debug(self, index!() - start, start),
        selector_error(selector = self, arg = l),
        nil)


// a generic method : same as previously but (1) can be called by other methods
// and (2) takes care of the debugging piece, which implies a slower run (GC)
execute(self:method,start:integer,int?:boolean) : any
 -> (let n := length(self.domain) in
       (if known?(formula, self)
           let retour := base!(),
               st? := (system.debug! != -1 & (int? | self.module!.status != 4)) in
             (set_base(start),
              stack_apply(self.formula.dimension),
              if st? push_debug(self.selector, n, start),
              let val := eval(self.formula.body) in         // error is possible 
                (if st? pop_debug(self.selector, 0, val),
                 set_base(retour),
                 set_index(start),
                 if (system.debug! != -1 & not(val % self.range))     // v3.2.01 
                     range_error(cause = self, arg = val, wrong = self.range), 
                 val))
        else let st? := (system.debug! != -1 &                     // system.debug > 0 <=> push for debug
                         (int? | self.module!.status != 3) &       
                         self.selector != debug),                  // do not record debug(...)
                 i := index!() in
               (if st? push_debug(self.selector, n, start),
                let val := stack_apply(self.functional, start, i) in
                  (if st? pop_debug(self.selector, 0, val),
                   if (system.debug! != -1 & not(val % self.range))    // v3.2.01
                      range_error(cause = self, arg = val, wrong = self.range), 
                   val))))

// the evaluator is open coded
eval(self:any) : any -> eval(self)

// this is the standard evaluation
// self_eval(self:object) : any -> self

// reads an inline definition for a method
// notice that it does not return an error
[inlineok?(self:method,s:string) : method
 -> try let p := read, l := call(p, s) in
        (self.inline? := true, self.formula := (l as lambda))
     catch any unsafe(trace(0,"---- WARNING: inline definition of ~S is wrong\n", self)),
     self ]

// reads a lambda - may return an error
[claire/read_lambda(s:string) : lambda
 -> try let p := read, l := call(p, s) in
          (case l (lambda l, any error("compiled lambda error with ~S (not a lambda!)",s)))
    catch any error("compiled lambda parse error with ~S",s)]     

// ****************************************************************
// *    Part 2: Update methods                                    *
// ****************************************************************

//get/put for a slot: should be inline
get(s:slot,x:object) : any -> slot_get(x, s.index, s.srange)
put(s:slot,x:object,y:any) : any
 -> store(x, s.index, s.srange, y, s.selector.store?)

// reading a value from a property (unknown is allowed)
// when unknown is not allowed, we use read which is defined in Kernel
get(self:property,x:object) : any
 -> (let s := (self @ owner(x)) in
       case s (slot slot_get(x, s.index, s.srange), any unknown))

// a more general value that is useful for types
funcall(self:property,x:any) : any
 -> (let s := (self @ owner(x)) in
       case s (slot slot_get(x as object, s.index, s.srange),
               method funcall(s,x),
               any unknown))

// verifying
hold?(self:property,x:object,y:any) : boolean
 -> (let s := (self @ owner(x)) in
       case s
        (slot let z := slot_get(x, s.index, s.srange) in
                case z (set contain?(z, y), any y = z),
         any false))

// writing a single value into a slot & trigger rules
//  write = check + put + put_inverse + propagate
//  if_write = put + put_inverse + propagate  (propagate => if_write)
//  update = put + put_inverse
// note in CLAIRE 4: with no inverse/store write_fast, defined in Kernel, works better
write(self:property,x:object,y:any) : void
 -> (let s := (self @ owner(x)) in
       case s
        (slot (if not(y % s.range) range_is_wrong(s, y)
              else if (self.open < 1 &
                       known?(slot_get(x, s.index, s.srange)))
                 error("[132] Cannot change ~S(~S)", self, x)
              else if (known?(if_write,self) & not(self.multivalued?))
                 fastcall(self,x,y)
              else update(self, x, s.index, s.srange, y)),
         any selector_error(selector = self, arg = list(x))),
     y)

// the value does not belong to the range: error!
range_is_wrong(self:slot,y:any) : void
 -> range_error(cause = self, arg = y, wrong = self.range)

// to remove
[put(p:property,x:object,n:integer,s:class,y:any) : void
 -> // are you still using this dead thing (put instead of update) ??? !! //,
    update(p,x,n,s,y) ]

// update (method called by the compiler)     // v3.0.20: renamed from put !
// update = put + put_inverse  (complex links) .. it does not trigger the rules (if_write)
// update uses two satellite methods: update+ and update-
// CLAIRE 4: inverse management only applies with set multivalued properties
update(p:property,x:object,n:integer,s:class,y:any) : void
 -> let old := slot_get(x, n, s) in
       (if (verbose() = 8) trace(0,"update ~S(~S) old = ~S\n", p,x,old),
        if (p.multivalued? = true)      // multivalued -> set based in CLAIRE 4       
           (if (length(old) > 0)
               let v := set()  in
                 (if (of(old) != void) cast!(v,of(old)),
                  store(x, n, s, v, p.store?)),
            when r := get(inverse, p) in for z in (old as set) update-(r, z, x),
            for z in (y as set) add!(p, x, n, z))
        else if (old != y)
           (when r := get(inverse, p) in
                (if (known?(old) & (r != p | x != old)) update-(r, old, x)),
            store(x, n, s, y, p.store?),
            update+(p, x, y)),
        y)

// this method checks the correctness of the inverse from a global view.
update+(self:relation,x:any,y:any) : void
 -> let r := self.inverse in
       (if (known?(r) & (r != self | x != y))
           (case r
             (property let s := (r @ owner(y)) in
                         case s
                          (slot
                             let old_y := get(s, (y as object)) in
                               (if (r.multivalued? != false)
                                   add_value!(r, (y as object), s.index, old_y as set, x)
                                else store(y, s.index, s.srange, x, r.store?)),
                           any error("[133] Inversion of ~S(~S,~S) impossible", self, x, y)),
              table let old_v := get(r,y) in
                      (if (r.multivalued? != false) add_value!(r as table, y, old_v as set, x)
                       else (if known?(old_v) update-(self, old_v, y),
                             put(r, y, x))))))

// this methods deletes a value in the inverse of a global_relation
update-(r:relation,x:any,y:any) : void
 -> (case r
      (property let s := (r @ owner(x)) in
                  case s
                   (slot let l := get(s, x) as any,            // Claire4
                             v := (case l
                                    (set (if r.store? copy(l as set) else (l as set)) delete y,
                                     any  unknown)) in
                           put(s, x, v)),
       table let l := get(r,x) as any,                          // previous r(x)
                 v := (case l (set (if r.store? copy(l as set) else (l as set)) delete y,   // remove v if l is a set
                               any unknown)) in
               put(r, x, v)))

// this methods adds a value to a multi-slot (used by the compiler)
// this is the multi-valued equivalent of update - we know self to be multivalued (hence a set in Claire 4)
[add!(self:property,x:object,n:integer,y:any) : void
 -> if known?(if_write,self) fastcall(self,x,y)
    else let s1 := (slot_get(x, n, object) as set) in
       (if add_value!(self, x, n, s1, y) update+(self, x, y)) ]

// this methods adds a value to a multi-slot (internal form)
// this is the multi-valued equivalent of put
// return true if the set is actually changed (y added to s)
[add_value!(self:property,x:object,n:integer,s1:set,y:any) : boolean
 ->   if not(y % s1)
        let s2 := (add!@set((if self.store? copy(s1) else s1), y)) in
                (store(x, n, object, s2, self.store?), true)
       else false]

// same method with error checking
[add(self:property,x:object,y:any) : void
 -> let s := (self @ owner(x)) in
       (// [5] add ~S(~S) <- ~S // self,x,y,
        if not(s) selector_error(selector = self, arg = list(x))
        else if not(multivalued?(self)) error("[134] Cannot apply add to ~S", self)
        else if (y % member((s as slot).range))
          (if known?(if_write,self) fastcall(self,x,y)
           else add!(self, x, (s as slot).index, y))
        else range_is_wrong((s as slot), y)),
     y ]

// known ?
known?(self:property,x:object) : boolean
 -> (let s := (self @ owner(x)) in
       case s (slot slot_get(x, s.index, s.srange) != unknown, any false))
unknown?(self:property,x:object) : boolean
 -> (let s := (self @ owner(x)) in
       case s (slot slot_get(x, s.index, s.srange) = unknown, any true))

// delete takes care of the inverse also
// assumes that self is multivalued -> should check !
delete(self:property,x:object,y:any) : any
 -> (let s := (self @ owner(x)) in
       (if not(s) selector_error(selector = self, arg = list(x))
        else if (self.multivalued? = true)
            let l1 := (slot_get(x, (s as slot).index, object) as set),
                 l := ((if self.store? copy(l1) else l1) delete y) in
               (store(x, (s as slot).index, object, l, self.store?),
                let r := self.inverse in (if known?(r) update-(r, y, x)),
                l)))

// erase is similar for mono-valued properties takes care of the inverse also
// v3.2.22: take care of multi-valued slot as well
erase(self:property,x:object) : any
 -> let s := (self @ owner(x)) in
       (if not(s) selector_error(selector = self, arg = list(x))
        else let y := slot_get(x, (s as slot).index, s.srange) in
               (if (self.multivalued? = true)       // multivalued is based on sets in CLAIRE 4
                   (when r := get(inverse,self) in for y1 in (y  as set) update-(r,y1,x),
                    let l := empty(y as set) in
                      (store(x, (s as slot).index, object, l, self.store?),
                       l))
                else (store(x, (s as slot).index, s.srange, default(s), self.store?),
                      let r := self.inverse in (if (known?(r) & known?(y)) update-(r, y, x)),
                      default(s))))



[set_range(p:property,c:class,r:type) : void
 -> let s := ((p @ c) as slot) in
       (s.range := r,
        s.srange := class!@type(r)) ]

// no longer needed because changing the range is not changing the prototype ?
// we should rather generate an error if the condition for dealing with
// defaults changes (TODO)
//        if (s.srange != any & s.srange != integer)
//           c.prototype[s.index] := 0))

// this method allows to bypass the storage mechanism - to be optimized ..
put_store(self:property,x:object,y:any,b:boolean) : void
 -> let s := (self @ owner(x)) in
      (case s
        (slot let z := slot_get(x, s.index, s.srange) in
                   (if (z != y) store(x, s.index, s.srange, y, b)),     // v3.2.04 same behavior compiled/interpreted !
         any selector_error(selector = self, arg = list(x))))


// tells if we have a multivalued relation nolonger used in CLAIRE 4
// multi? :: property()
// [multi?(x:any) : boolean
//   -> case x (relation (x.multivalued? != false), any false) ]

// new: (v3.0) we have a simpler management of demons thus fastcall can be
// written in CLAIRE. A demon is either a lambda or a function, which
// arguments have precise sorts that match the relation sort
// the demon does everything: put + inverse + propagation
// thus write(R,x,y) <=> fastcall(R,x,y) <=> f(x,y)
fastcall(r:relation,x:any,y:any) : void
 -> let f := r.if_write in
      (if (f % lambda) funcall(f as lambda, x , y)
       else funcall((f as function),x,y))

// *********************************************************************
// *   Part 3: Management of definition(p)                             *
// *********************************************************************
join :: operation()

// the dictionarty slot

// insertion in the definition tree
insert_definition(p:property,r:restriction) : void
 -> (put(definition, p,
         initialize(r, class!(r.domain[1]), p.definition)))

// insert a restriction with class-domain d into a property p
// claire4 : get rid of dispatcher
initialize(x:restriction,d:class,l:list) : list
 -> let p := x.selector in
       (if (length(p.restrictions) = 5 & uniform(p))
           // dictionary = true means uniform => restrictions are stored in c.dictionary
           (for r in p.restrictions hashinsert(r),
            p.dictionary := true),
        if p.dictionary
           (if uniform(x) (hashinsert(x))
            else (// printf("// ---- note: ~S is not uniform because of ~S \n", p, x),  // not necessarily a bug :-)
                  p.dictionary := false)),
        initialize(x, l))

// only uniform properties can use the dictionary representation
uniform(x:restriction) : boolean
 -> let l := x.domain,
         n := length(l) in
       forall(r in x.selector.restrictions |
         let l2 := r.domain in
           (l2[1] % class & length(l2) = n & l2[1] != listargs &
            (forall(i in (2 .. n) | 
                    (l[i] = l2[i] |                   // v3.3.34
                     (owner(l[i]) != class &          // introduce a protected call to =type !
                      owner(l[i]) = owner(l2[i]) & l[i] =type? l2[i]))))))    // v3.3.36      

// v3.0.54 check that a uniform property only uses methods !
uniform(p:property) : boolean
  -> (forall(x in p.restrictions | x % method) &  uniform(p.restrictions[1]))

// insert a restriction in a list with the good order
[initialize(x:restriction,l:list) : list
 -> let l1:list := nil in                       // no updates on nil
       (for i in (1 .. length(l))
          let l2 := (l[i] as restriction).domain in
            (if tmatch?(x.domain, l2)
                (if tmatch?(l2, x.domain)
                    (l[i] := x, l1 := l, break(true))
                 else (l1 := unsafe(nth+(l, i, x)) as list, break(true)))
             else if (not(tmatch?(l2, x.domain)) & x.domain join l2 & x.selector.open <= 1)
                unsafe(trace(2, ("Note: ~S and ~S are conflicting\n"), l[1], x))), // keep the trace
        if (length(l1) != 0) l1
        else add!(l, x)) ]

// definition of dictionary: standart hash-table
[hashinsert(m:restriction) : any
 -> // if (verbose() = 4) //[0] hashinsert(~S) // m,
    let c := (domain!(m) as class) in
       for c2 in c.descendants hashinsert(c2, (m as method)) ]

// insert into the hash table - since the order is not garanteed when we build the dictionary, we
// need to check that m is more suited than anything that could be there
[hashinsert(c:class,m:method) : any
 -> if (c.dictionary = unknown) c.dictionary := map!(property,method),   
    let m1 := dict_get(c.dictionary,m.selector) in
       (if (m1 = unknown |  domain!(m) inherit? domain!(m1 as method))
            dict_put(c.dictionary,m.selector,m)),
    c.dictionary ]

// read the value in the directory (a method or unknown)
[hashget(c:class,p:property) : any
 -> dict_get(c.dictionary,p) as object]        // UGLY CAST to remove

// look if two signature have a non-empty intersection
// note that the first case with classes is necessary for bootstraping
[join(x:list,y:list) : boolean
 -> let n := length(x) in
       (n = length(y) &
        forall(i in (1 .. n) | class!(x[i] as type) join class!(y[i] as type)) &
        forall(i in (1 .. n) | unsafe(x[i] glb y[i]))) ]

// *********************************************************************
// *      Part 4: Matching Methods                                     *
// *********************************************************************

// Key Axiom : this code is not using dynamic calls because we use the two closed forms %type and <=t 
// which are defined in Kernel as functions (Contains and Included)

// this is the method that matches the compilation pattern 
// n is the number of args that have been pushed in the stack
[stack_apply(p:property,n:integer) : any
  -> let i := index() - n in
       eval_message(p, find_which(p,i,owner(get_stack(i))),i,false) ]


// version where the class of first argument is forced (super)       
[super_apply(p:property,c:class,n:integer) : any
  -> let top := index(), i := top - n in
       eval_message(p, find_which(c, p.Core/definition,i,top),i,false) ]


// find the correct restrictions to be applied on a given set
// This is also optimized because it is very useful (it returns false if none is found)
[@(self:property,x:class) : object
 ->  if self.dictionary let rx := hashget(x, self) in
       (if known?(rx) rx else false)
     else let rx := some(r in self.definition |
                          (x inherit? class!(r.domain[1]))) in
       (if known?(rx) rx else false) ]


// finds a property through its full domain
[@(self:property,lt:list) : object
 ->  let rx := some(r in self.definition | tmatch?(lt, r.domain)) in
       (if known?(rx) rx else false) ]

// method's pattern matching : l is non nil, hence last(l) is safe  {called in find_which}
// we match a list of args in the stack [n ... m] to the list of type_expressions l
matching?(l:list,n:integer,m:integer) : boolean
 -> (let x := (m - n), z := length(l) in
       (if (z = x & nth_object(l,x) != listargs)
           not((for i in (1 .. x)
                  let y := ((n - 1) + i),
                      u := nth_object(l,i) in
                    (if (if (owner(u) = class)
                         not(owner(get_stack(y)) inherit? (u as class))
                      else not(vmatch?(u, get_stack(y), n))) break(true))))
        else if (unsafe(last(l)) = listargs & x >= z - 1)   // v3.2.24
           not((for i in (1 .. z)
                  let y := ((n - 1) + i) in
                    (if (l[i] = listargs)
                        (put_stack(y, get_args(y)),
                         set_index(y + 1),
                         break(false))
                     else if not(vmatch?(l[i], get_stack(y), n))
                        break(true))))
        else false))

// type's pattern matching - almost like % but accepts patterns such as Reference (extended in Optimizer)
// this is why we pass n (index in stack) as an argument
// t is the type expression and x is the value
[vmatch?(t:any,x:any,n:integer) : boolean
 -> case t
      (class owner(x) inherit? t,
       set contain?(t, x),
       subtype // (//[0] vmatch ~S with type ~S => ~S // x,t, (x as type) <=t  t.t1),
                ((if (t.arg = subtype) x % type else x % t.arg) &  (x as type) <=t  t.t1),
       Param (vmatch?(t.arg, x, n) &                   //  t = t.arg[p_i : t_i]
              forall(i in (1 .. length(t.params)) |
                      let %t := t.args[i], %v :=  unsafe(funcall((t.params[i] as property),x)) in
                         (if (%t % set & %v % type) // %type(%v,%t) :  pattern extension (not in <=t)
                            exists(z in (%t as set) | =type?(%v,z))   // different from 3.5, %type no longer exists
                          else vmatch?(%t, %v, n)))),     // %v = p(x) % t2
       Reference let v := get(t, get_stack(n + t.index)) in
                   (if t.arg x = v else x %t v),
       tuple case x
             (tuple (length(t) = length(x) &
                     forall( i in (1 .. length(x)) | vmatch?(t[i], x[i], n))),
              any false),
       any unsafe(x % t)) ]            // extensibility for type_expressions

// method's pattern matching based on type expressions (i.e. l2 is another list of type expressions).
// this is an extension of <=t to   all type expressions
tmatch?(l:list,l2:list) : boolean
 -> (let x := length(l2), z := length(l) in
       (if (z != x & (l2[x] != listargs | z < x - 1)) false           // v3.2.24
        else not((for i in (1 .. x)
                    (if (i = x & l2[i] = listargs) break(false)
                     else if not(tmatch?(l[i], l2[i], l)) break(true))))))

// type_expression pattern matching (t is the variable and t2 the pattern)
// this is an extension of <=t for the pattern Reference
[tmatch?(t:any,t2:any,l:list) : boolean
 -> case t2
      (Reference (if t2.arg false            // this is very unclear ! t2.arg -> t2 was copied in odefine.cl
                  else case t
                    (Reference ((t.index = t2.index) & (t.args = t2.args)),
                     type let tref := member(@(t2, t2.args, l[t2.index + 1])) in // member(X) because X is a container type (the value belongs to X)
                         (//[5] tmatch?: is ~S less than ~S ? reference gives ~S // t,t2,tref,
                          t <=t tref))),     // t is less than member(X), the "type contraint" extracted from the reference
       type (case t (type (t <=t t2),
                     any less?(t,t2))),      // extensibility with less?
       any less?(t,t2))]

// find the restriction (n is the position of the arglist start)
[find_which(p:property,n:integer,c:class) : object
 -> (if p.dictionary hashget(c, p) // v3.2.58  was ... (length(p.dictionary) != 0) hashget(c, p)
     else for r:restriction in p.definition
            (if matching?(r.domain, n, index!()) break(r))) as object]

// used by inspect.cl
[find_which(l:list,c:class,n:integer,m:integer) : object
  -> (for r:restriction in l (if matching?(r.domain, n, m) break(r))) as object]

// special version for super, where we give (n,m) -> position of arglist in the stack
[find_which(c:class,l:list,n:integer,m:integer) : object
 -> (for r:restriction in l
       (if (c <=t r.domain[1] & matching?(r.domain, n, m)) break(r))) as object]



