//+-------------------------------------------------------------+
//| CLAIRE                                                      |
//| control.cl                                                  |
//| Copyright (C) 1994 - 2021 Yves Caseau. All Rights Reserved  |
//| cf. copyright info in file object.cl: about()               |
//+-------------------------------------------------------------+

// *********************************************************************
// *     Part 1: If, Do, Let                                           *
// *     Part 2: set control structures                                *
// *     Part 3: other control structures                              *
// *     Part 4: the constructs                                        *
// *********************************************************************

// *********************************************************************
// *     Part 1: If, Do, Let                                           *
// *********************************************************************

//--------------- the IF --------------------------------------------
If <: Control_structure(test:any,arg:any,other:any = false)
claire/If? <: If()              // v4.12 conditional comment

self_print(self:If) : void
 -> printf("(~I~I~I)", (pretty.index :+ 1), printstat(self),
           (pretty.index :- 1))

printstat(self:If) : void
 -> printf("if ~I ~I~I~I", (printexp(self.test, false), checkfar()),
           printif(self.arg), (pretty.index :- 3), printelse(self))

printif(self:any) : void
 -> (pretty.index :+ 3,
     if pretty.mClaire/pbreak
        let b_index := mClaire/buffer_length(),
            %l := pretty.index in
          (pretty.mClaire/pbreak := false,
           try print(self) catch much_too_far {},
         pretty.mClaire/pbreak := true,
         if short_enough(mClaire/buffer_length()) {}
         else (mClaire/buffer_set_length(b_index),
               pretty.index := %l,
               lbreak(),
               print(self)))
     else print(self))

printelse(self:If) : void
 -> (let e := get(other, self) in
       case e
        (If printf(" ~Ielse if ~I ~I~I~I", lbreak(),
                   printexp(e.test, false), printif(e.arg),
                   (pretty.index :- 3), printelse(e)),
         any (if (e != nil)
                let %l := pretty.index in
                  (printf(" ~Ielse ~I~S", lbreak(), set_level(1), e),
                   pretty.index := %l))))

// notice that the eval(test) is not a boolean thus the compiler will add
// something
// TODO: check that is is not too slow (may use a constant for _oid_(true))
self_eval(self:If) : any
 -> let x := eval(self.test) in
      (if (x = true) eval(self.arg)
       else if (x = false) eval(self.other)
       else if x eval(self.arg)
       else eval(self.other))

//--------------------- block structure------------------------------
Do[args] <: Control_structure(args:list)

self_print(self:Do) : void
 -> (let %l := pretty.index in
       (printf("("),
        set_level(1),
        printdo(self.args, true),
        pretty.index := %l))

printdo(l:list,clo:boolean) : void
 -> (let n := length(l) in
       for x in l
         (case x (If printstat(x), any print(x)),
          n :- 1,
          if (n = 0) (if clo princ(")"))
          else (printf(", "), lbreak())))

printblock(x:any) : void
 -> (case x (Do printdo(x.args, false), If printstat(x), any print(x)))

// use res:EID pragma when compiled with CLAIRE4, res:any for CLAIRE3
self_eval(self:Do) : any
 -> (let res:any := {} in (for %x in self.args res := eval(%x), res))

// ----------------- lexical variable definition -----------------------
Let <: Instruction_with_var(value:any,arg:any)

self_print(self:Let) : void
 -> (let %l := pretty.index in
       (set_level(1),
        printf("let ~I := ~I~I", ppvariable(self.var),
               printexp(get(value, self), false), printbody(self)),
        pretty.index := %l))

printbody(self:Let) : void
 -> (let a := self.arg in
       case a
        (Let printf(",~I~I := ~I~I~I", lbreak(4), ppvariable(a.var),
                    printexp(get(value, a), false), (pretty.index :- 4),
                    printbody(a)),
         any printf(" in ~I~S", lbreak(2), a)))

self_eval(self:Let) : any
 -> let val := eval(self.value) in
             (write_value(self.var as Variable, val),
              eval(self.arg))

// a when is a special Let that filters out the unknown value !
//
When <: Let(other:any = unknown)
self_print(self:When) : void
 -> (let %l := pretty.index in
       (set_level(1),
        printf("when ~I := ~I in ~I~S", ppvariable(self.var),
               printexp(get(value, self), false), lbreak(2), self.arg),
        if known?(other, self)
           printf(" ~Ielse ~I~S", lbreak(), set_level(1), self.other),
        pretty.index := %l))

self_eval(self:When) : any
 -> (let val := eval(self.value),
         n := system.Core/trace! in
       (if (val != unknown)
           (write_value@Variable(self.var, val), eval(self.arg))
        else eval(self.other)))

// two special forms of Let:
// Let+(v,r(x),(r(x) := y),Let(v2,e,(r(x) := v,v2)))    <=>  let r(x) = y in e
// Let*(v,f(),Let(v1,v[1],...(Let(vn,v[n],e))   <=> let (v1,v2,...vn) := f() in e
//
Let+ <: Let()
Let* <: Let()

//note: the Let* is also used for multi-assignments
// Let*(v,f(),(v1 := v[1], v2 := v[2], ...))   <=>  (v1,v2,...vn) := f()
//
self_print(self:Let+) : void
 -> (let %l := pretty.index,
         l := (self.arg as Do).args in
       (set_level(1),
        printf("let ~I := ~I in ~I~S", printexp(self.value, false),
               printexp(l[1].args[3], false), lbreak(2), (l[2] as Let).value),
        pretty.index := %l))

self_print(self:Let*) : void
 -> (let %l := pretty.index,
         l := self.arg in
       (set_level(1),
        if (l % Let)
           printf("let (~I) := ~I~I",
                  (while true
                     (ppvariable(l.var),
                      let lnext := l.arg in
                        (if (case lnext
                              (Let
                                 (lnext.value % Call &
                                  lnext.value.args[1] = self.var)))
                            (princ(","), l := lnext)
                         else break(true)))),
                  printexp(get(value, self), false), printbody(l))
        else printf("(~I) := ~I",
                    (let %f := true in
                       for %a in l.args
                         (if %f %f := false
                          else princ(","),
                          ppvariable(%a.var))),
                    printexp(get(value, self), false)),
        pretty.index := %l))

// *********************************************************************
// *     Part 2: set control structures                                *
// *********************************************************************
Iteration <: Instruction_with_var(set_arg:any,arg:any)
iterate :: property()
Iterate :: property()
IterateFast :: property()

// for is the simplest evaluation loop
//
For <: Iteration()
self_print(self:For) : void
 -> printf("for ~I in ~I ~I", ppvariable(self.var),
           (let %l := pretty.index in
              (set_level(),
               printexp(self.set_arg, false),
               pretty.index := %l)),
           (pretty.index :+ 2,
            lbreak(),
            print(self.arg),
            pretty.index :- 2))

self_eval(self:For) : any
 -> (let x := eval(self.set_arg) in
       (try case x
         (class for y in x.descendants
                 for z in y.instances
                   (write_value(self.var, z), eval(self.arg)),
          list for z in x
                    (write_value(self.var, z), eval(self.arg)),
         array let n := length(x) in
                  for z in list{nth_get(x,i) | i in (1 .. n)}
                    (write_value(self.var, z), eval(self.arg)),
          Interval for y in (x.arg1 .. x.arg2)
                      (write_value(self.var, y), eval(self.arg)),
          collection for y in x (write_value(self.var, y), eval(self.arg)),
          any error("[136] ~S is not a collection !", x))
        catch return_error system.exception!.arg))

// [collect VAR in SET_EXPR, ...] is the same as a "for", but returns the list of values
//
Collect <: Iteration(of:type)

self_print(self:Collect) : void
 -> printf("list{ ~I | ~I~I in ~I}",
           (pretty.index :+ 2, printexp(self.arg, false)),
           lbreak(),
           ppvariable(self.var),
           (let %l := pretty.index in
              (set_level(),
               printexp(self.set_arg, false),
               pretty.index := %l - 2)))

// list image : preserve the order for lists and intervals (v4)
self_eval(self:Collect) : any
 -> (let x := eval(self.set_arg),
         res:list := empty_list((if known?(of,self) self.of else {})) in
       (case x
         (class for y in x.descendants
                 for z in y.instances
                   (write_value(self.var, z), res :add eval(self.arg)),
          list for y in x
               (write_value(self.var, y), res :add eval(self.arg)),
          Interval for y in x
               (write_value(self.var, y), res :add eval(self.arg)),
          any for y in x
               (write_value(self.var, y), res :add eval(self.arg))),
         res))

// this is a set image version, that produces a set
//
Image <: Iteration(of:type)

self_print(self:Image) : void
 -> printf("{ ~I | ~I~I in ~I}",
           (pretty.index :+ 2, printexp(self.arg, false)), lbreak(),
           ppvariable(self.var),
           (let %l := pretty.index in
              (set_level(),
               printexp(self.set_arg, false),
               pretty.index := %l - 2)))

self_eval(self:Image) : any
 -> (let x := eval(self.set_arg),
         res:set :=  empty_set((if known?(of,self) self.of else {})) in
       (for y in x (write_value(self.var, y), res :add eval(self.arg)),
        res))

// [select VAR in SET_EXPR, ...] is the same as a "for" but returns the subset of
//  members that produce a true value
//
Select <: Iteration(of:type)
self_print(self:Select) : void
 -> printf("{ ~I in ~I | ~I}", ppvariable(self.var),
           (let %l := pretty.index in
              (set_level(),
               printexp(self.set_arg, false),
               pretty.index := %l)),
           (lbreak(2), print(self.arg), pretty.index :- 2))

self_eval(self:Select) : any
 -> (let x := eval(self.set_arg), 
         res:set :=  empty_set((if known?(of,self) self.of else {})) in
       (case x
         (class for y in x.descendants
                 for z in y.instances
                   (write_value(self.var, z),
                    if (eval(self.arg) != false) res :add z),
          Interval for y in x
               (write_value(self.var, y), 
                if (eval(self.arg) != false) res :add y),
          any for y in x
               (write_value(self.var, y), 
                if (eval(self.arg) != false) res :add y)),
        res))

// [select VAR in SET_EXPR, ...] is the same as a "for" but returns the subset of
//  members that produce a true value
//
Lselect <: Iteration(of:type)

self_print(self:Lselect) : void
 -> printf("list{ ~I in ~I | ~I}", ppvariable(self.var),
           (let %l := pretty.index in
              (set_level(),
               printexp(self.set_arg, false),
               pretty.index := %l)),
           (lbreak(2), print(self.arg), pretty.index :- 2))

self_eval(self:Lselect) : any
 -> (let x := eval(self.set_arg),
         res:list := (case x (list empty(x), any list())) in
       (case x
         (class for y in x.descendants
                 for z in y.instances
                   (write_value(self.var, z),
                    if (eval(self.arg) != false) res :add z),
          any for y in x
               (write_value(self.var, y), if (eval(self.arg) != false) res :add y)),
        if known?(of,self)
          (when x := some(x in res | not(x % self.of)) in   // v3.1.06
               range_error(mClaire/cause = self,arg = x, Core/wrong = self.of),
           Kernel/cast!(res,self.of)),
        res))

// Exists is an iteration that checks a condition
// other = true => forall,  other = false => exists, other = unknown => some
Exists <: Iteration(other:any = false)

self_print(self:Exists) : void
 -> (if (self.other = true) princ("forall")
     else if (self.other = false) princ("exists")
     else princ("some"),
     if (self.set_arg = any)
        printf("(~I,~I)", ppvariable(self.var), print(self.arg))
     else printf("(~I in ~I | ~I)", ppvariable(self.var),
                 (let %l := pretty.index in
                    (set_level(),
                     printexp(self.set_arg, false),
                     pretty.index := %l)),
                 (lbreak(2), print(self.arg), pretty.index :- 2)))

self_eval(self:Exists) : any
 -> (let x := eval(self.set_arg),
         b := self.other,
         res:any := b in
       (case x
         (class for y in x.descendants
                 for z in y.instances
                   (write_value(self.var, z),
                    if (eval(self.arg) != false)
                       (if (b != true) break(res := (if b z else true)))
                    else if (b = true) break(res := false)),
          any for y in x
               (write_value(self.var, y),
                if (eval(self.arg) != false)
                   (if (b != true) break(res := (if b y else true)))
                else if (b = true) break(res := false))),
        res))

// *********************************************************************
// *     Part 3: other control structures                              *
// *********************************************************************

// ----------------- case  --------------------------------------
Case <: Control_structure(var:any,args:list)

self_print(self:Case) : void
 -> (printf("case ~S ~I(", self.var, lbreak(1)),
     let n := 1,
         m := length(self.args) in
       (pretty.index :+ 1,
        while (n <= m)
          let %l := pretty.index in
            (printf("~I ~I~I", printexp(self.args[n], false),
                    (if (mClaire/buffer_length() > (pretty.mClaire/width - 50))
                        lbreak(2)
                     else set_level(),
                     print(self.args[n + 1])),
                    (pretty.index := %l,
                     if ((n + 1) != m) printf(", ~I", lbreak()))),
             n :+ 2),
        printf(")"),
        pretty.index :- 2))

[self_eval(self:Case) : any
 -> let truc := eval(self.var),
         flip:boolean := true,
         previous:any := false in
     (if (for x in self.args
          (if flip (flip := false, previous := eval(x))
           else if (truc % previous)
                  (previous := eval(x), break(true))
           else flip := true))
        previous
      else false) ]

// ------------------ WHILE  and UNTIL  -----------------------------
// the "other" while is until, where the first test is skipped
While <: Control_structure(test:any,arg:any,other:boolean = false)

self_print(self:While) : void
 -> (printf("~A ~I ~I~S", (if self.other "until" else "while"),
            printexp(self.test, false), lbreak(2), self.arg),
     pretty.index :- 2)

// other = true => self means  repeat self.arg until self.test = true
self_eval(self:While) : any
 -> (let a := self.other,
         b := a in
       try while (b | not(eval(self.test)) = a)
         (b := false, eval(self.arg))
       catch return_error system.exception!.arg)

//-------------- handling errors -----------------------------------
// This is the control structure associated with these errors. Its real
// semantics is defined in the C compiler file
//
Handle <: Control_structure(test:any,arg:any,other:any)

self_print(self:Handle) : void
 -> (printf("try ~S ~Icatch ~S ~S", self.arg, lbreak(0), self.test,
            self.other),
     pretty.index :- 2)

// original code
// self_eval(self:Handle) : any
//  -> (let x := (self.test as class) in
//       try eval(self.arg)
//       catch x (if (exception!() % return_error) close(exception!())
//                else eval(self.other)))     // <yc> 6/98
// CLAIRE 4 VERSION, because catch x => x is a constant class
// notice that return_error should be called return_exception since they travel through interpreted
// not a problem at compile time since return_exceptions are handled with break(x)
self_eval(self:Handle) : any
 ->  (try eval(self.arg)
      catch any let e := exception!(), x := (self.test as class) in
         (if (e % return_error |  not(e % x)) close(e)    // move to next handle
          else eval(self.other)))     // <yc> 6/98

// *********************************************************************
// *     Part 4: the constructs                                         *
// *********************************************************************

Construct <: Complex_instruction(args:list)

List <: Construct(of:type)
Tuple <: Construct()
Set <: Construct(of:type)
Array <: Construct(of:type)        // v3.2.16   constructor for arrays
Printf <: Construct()
Error <: Construct()
Branch <: Construct()
claire/Map <: Construct(domain:type,of:type)

self_print(self:Construct) : void
 -> (let %l := pretty.index in
       (printf("~A~I(~I~I)",       // v3.2.56-58 add a <type> when needed
               (case self
                 (List "list",
                  Set "set",
                  Tuple "tuple",
                  Printf "printf",
                  Error "error",
                  Trace "trace",
                  Assert "assert",
                  Branch "branch",
                  Map "map",
                  any string!(self.isa.name))),
               (case self ((List U Set)
                            when %t := get(of,self) in (if (%t != {}) printf("<~S>", %t) ),
                           Map printf("<~S,~S>",self.domain,self.of))),
               set_level(), printbox(self.args)),
        pretty.index := %l))

// constructors: how to create a list, a set, a tuple or an array
// note that the constructor is typed
// CLAIRE4: must build the list with the proper type from the beginning, so that Srange is correct
self_eval(self:List) : list
  -> let type? := known?(of,self), n := length(self.args) in
        (if type? 
           let l := make_list(n, self.of, void)  in                          // compiler optimization
               (for i in (1 .. n)  nth_write(l,i, eval(self.args[i])),       // write with a test
               l)
         else let l := make_list(n, {}, void)  in              // compiler optimization
               (for i in (1 .. n)  l[i] := eval(self.args[i]),              // write without a test
               l))

// here we use the CLAIRE 3 style of post-typing with a cast! 
self_eval(self:Set) : set
 -> let type? := known?(of,self), n := length(self.args) in
        (if type? 
           let l := empty_set(self.of)  in                          // compiler optimization
               (for i in (1 .. n)  add(l, eval(self.args[i])),       // adds with a test
               l)
         else let l := empty_set({})  in              // compiler optimization
               (for i in (1 .. n)  add!(l,eval(self.args[i])),              // adds without a test
               l))
/* -> let s := { eval(x) | x in self.args} in
      (if known?(of,self)
          (when x := some(x in s | not(x % self.of)) in   // v3.0.72
               range_error(cause = self,arg = x,wrong = self.of),
           Kernel/cast!(s,self.of))   // v0.01
       else Kernel/cast!(s,{})) */
       
self_eval(self:Tuple) : any
 -> tuple!(list{ eval(x) | x in self.args})

// same as creating a list (same constraints since same underlying structure)
self_eval(self:Array) : any
 -> let type? := known?(of,self), n := length(self.args) in
        (if type? 
           let l := make_list(n, self.of, void)  in                          // compiler optimization
               (for i in (1 .. n)  nth_write(l,i, eval(self.args[i])),       // write with a test
                array!(l))
         else let l := make_list(n, {}, void)  in              // compiler optimization
               (for i in (1 .. n)  l[i] := eval(self.args[i]),              // write without a test
                array!(l)))
       
 // create a map from a list of pairs
self_eval(self:Map) : map_set
  -> let m := map!(self.domain,self.of) in
       (for x in self.args
          (case x (pair put(m,eval(x.first),eval(x.second)),
                   any error("~S is not a pair, cannot be inserted in map ~S",x,m))),
        m)        

// Macros are a nice but undocumented feature of CLAIRE. This is deliberate :)
// it is an advanced feature for those who want to expand the language. This
// makes CLAIRE a nice framework for DSL
//
Macro <: Construct()
macroexpand :: property(open = 3)
self_eval(self:Macro) : any -> eval(call(macroexpand,self))


// error produces an exception of type general_error
self_eval(self:Error) : error
 -> (if (not(self.args) | not(self.args[1] % string))
        error("Syntax error: ~S", self),
     let x:general_error := mClaire/new!(general_error) in      // v3.2.26
        (x.mClaire/cause := car(self.args),
         x.arg := list{ eval(x) | x in cdr(self.args)},
         close@exception(x)))

// this is the basic tool for printing in CLAIRE. A complex statement
// is macroexpanded into basic printing instructions
//
[self_eval(self:Printf) : any
 -> let l := self.args,
        s := l[1] in
       (if not(s % string)
           error("[102] the first argument in ~S must be a string", self)
        else let i := 2,
                 n := get(s as string, '~') in
               (while not(n = 0)
                  let m := s[n + 1] in
                    (if (i > length(l))
                        error("[103] not enough arguments in ~S", self),
                     if (n > 1) princ(substring(s, 1, n - 1)),
                     if ('A' = m) princ(eval(l[i]))
                     else if ('S' = m) print(eval(l[i]))
                     else if ('F' = m)  // v3.4
                        let fv := eval(l[i]),                              // float value
                            p% := false,                                   // print a %
                            j := integer!(nth_get(s,n + 2,n + 2)) - 48 in
                          (if ('%' = s[n + 2]) (p% := true, j := 1, fv :* 100.0)
                           else if (j < 0 | j > 9) error("[189] F requires a single digit integer in ~S",self),
                           if (not(p%) & length(s) > n + 2 & '%' = s[n + 3]) 
                              (p% := true, fv :* 100.0, n :+ 1),
                           mClaire/printFDigit(fv,j),
                           if p% princ("%"),
                           n :+ 1)
                     else if ('I' = m) eval(l[i]),
                     i :+ 1,
                     s := substring(s, n + 2, 1000),
                     n := get(s, '~')),
                if s princ(s)),
                unknown) ]

// trace is refined in inspect.cl
// If trace_output() is known, use it, else use current output.
iClaire/trace_on :: property()      // defined in inspect.cl
Trace <: Construct()

// CLAIRE4: self_eval is defined once for all, hence extended
self_eval(self:Trace) : any
 -> (if (length(self.args) = 0) (if (system.trace! = 0) "inactive" else "active")
     else let a := self.args,
         l := list{ eval(x) | x in a},
         i := l[1],
         a2 := (if (length(a) > 1) a[2]) in
       (if (length(a) = 1)
           let a1 := eval(a[1]), p := trace_on in
             (if p.restrictions
                (if (system.trace! = 0) put(trace!, system, 1), 
                 call(p,a1)))
        else if (a2 % string & (case i (integer i <= system.verbose)))
           let p := get(ctrace, system) in
             (if known?(p) p := use_as_output(p),
              format(a2, l << 2),
              if known?(p) use_as_output(p),
              {})))

// assert is refined in trace.la
//
Assert <: Construct(Core/index:integer,external:string)

self_eval(self:Assert) : any
 -> (let a := self.args in
       (if (length(a) > 0 & known?(get(ctrace,system)) & not(eval(a[1])))
           let p := use_as_output(system.ctrace) in
             (printf("~S,line=~A: (ASSERT) ~S\n", self.external,
                     self.index, a[1]),
              use_as_output(p),
              if (system.Core/debug! >= 0) error("Assertion Violation"),
              {})))

self_eval(self:Branch) : any
 -> (if (length(self.args) != 1)
         error("[104] Syntax error with ~S (one arg. expected)",self),
     try (choice(), if (eval(self.args[1]) != false) true else (backtrack(), false))
     catch contradiction (backtrack(), false))

// end of file
