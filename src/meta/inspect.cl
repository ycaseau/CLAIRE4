//+-------------------------------------------------------------+
//| CLAIRE                                                      |
//| inspect.cl                                                  |
//| Copyright (C) 1994 - 2013 Yves Caseau. All Rights Reserved  |
//| cf. copyright info in file object.cl: about()               |
//+-------------------------------------------------------------+

// --------------------------------------------------------------------
// this file contains the CLAIRE run-time tools: inspect, trace & debug
// --------------------------------------------------------------------

// *********************************************************************
// * Contents                                                          *
// *      Part 1: Inspection                                           *
// *      Part 2: Trace                                                *
// *      Part 3: Debugger                                             *
// *      Part 4: Stepper & Measure                                    *
// *      Part 5: Profiler                                             *
// *********************************************************************

execute_do :: property()
execute_bk :: property()
inspect_loop :: property()
get_from_integer :: property()
top_debugger :: property()

// a useful global variable *last*
*last*:any := unknown
*index*:integer := unknown
*maxd*:integer := unknown
*curd*:integer := 0                     // v3.2.14 cleaner :-)
*showall*:boolean := true               // v0.01 stop the ... !

// this is the interface with the system
//
inspect_system(l:list) : any -> function!(InspectLoop)
debug_system() : any -> function!(DebugLoop)
step_system() : integer -> function!(StepLoop)
CommandLoop() : string -> function!(CommandLoopVoid)

// *********************************************************************
// *      Part 1: Inspection                                           *
// *********************************************************************
// this is the method that the user calls
//
inspect(self:any) : any
 -> let %read:any := {}, m0 := module!(), ix := 0 in
       (case self
         (bag for i in (1 .. length(self)) printf("~A: ~S\n", i, self[i]),
          object for rel in owner(self).slots
                  let m := module!(name(selector(rel))) in
                    (ix :+ 1,            // ix is the position in the slots list
                     if (m = m0 | m = claire | *showall*)
                       let val := get(rel,self) in
                        printf("~A: ~S = ~I\n", ix, rel.selector,
                           (case val (bag (if (length(val) < 10) pretty_print(val)
                                           else (pretty_print(
                                                   list{val[i] | i in (1 .. 9)}),
                                                 pretty_print("..."))),
                                      any pretty_print(val))))),
          any printf("~I\n", pretty_print(self))),
        inspect_system(list(self)),
        None)

// this is the inspect top_level
//
inspect_loop(%read:any,old:list) : any
 -> (let self:any := old[1] in
       (if (case %read (Call %read.selector = put))
           let n:integer := %read.args[1],
               s:symbol := extract_symbol(%read.args[2]) in
             (if not(n % integer) error("[128] ~S should be an integer", n),
              let val := get_from_integer(self, n) in
                (put(value, new(global_variable, s), val),
                 inspect(val),
                 old := val cons old))
        else if (%read = up)
           (if (length(old) > 1) (old := cdr(old), inspect(old[1])))
        else if (%read % integer)
           let val := get_from_integer(self, %read) in
             (old := val cons old, inspect(val))
        else if (%read % thing) (old := %read cons old, inspect(%read))
        else printf("=> given to inspector is wrong.\n")),
     inspect_system(old))

// get the information bound to the index
//
get_from_integer(self:any,n:integer) : any
 -> (if (self % bag)
        (if (n > 0 & n <= length(self)) self[n]
         else (printf("~A in not a good index for ~S.\n", n, self), self))
     else let l := owner(self).slots in
            (if (n > 0 & n <= length(l)) get@slot(l[n], self)
             else (printf("~A is not a good index for ~S.\n", n, self),
                   self)))

// *********************************************************************
// *      Part 2: Trace methods                                        *
// *********************************************************************
// instrument the code generated from the rules

trace_on(self:any) : any
 -> (case self
      (property (if (self = spy)
                   let m := (spy @ void) in (if m put(Kernel/spy!, system, m))      ;*/
                 else if (self = where) system.Kernel/call_count := 1
                 else put(trace!, self, 5 - system.verbose)),
       environment put(trace!, system, 1),
       module (if (self.mClaire/status > 2) self.mClaire/status := 4,
               for m in self.parts trace_on(m)),
       port put(ctrace, system, self),
       string put(ctrace, system, fopen(self, "w")),
       integer put(verbose,system,self),
       any error("[129] trace not implemented on ~S\n", self)),
      self)

untrace(self:any) : any
 -> (case self
      (property (if (self = spy) put(Kernel/spy!, system, unknown)
                 else if (self = where) system.Kernel/call_count := -1
                 else put(trace!, self, 0)),
       environment put(trace!, system, 0),
       module (if (self.mClaire/status = 4) self.mClaire/status := 3,
               for m in self.parts untrace(m)),
       port put(ctrace, system, stdout),
       any error("[130] untrace not implemented on ~S\n", self)),
      self)

// a filter to restrict the impact of spy
// we put the special value nil (emply list of demons => OK) to mark that spying
// should be waken up on properties from l
spy(l:listargs)  : void
  -> (let m := (spy @ void) in
       (if m (put(Kernel/spy!, system, m),                   ;*/
              for f in property (if (f % l) f.mClaire/if_write := nil))))


// note: trace behavior for output statements defined in CLAIRE1 (self_eval)
self_trace(self:Trace) : any
 ->  let a := self.args in
       (if (length(a) = 1)
           let a1 := eval(a[1]) in
             (if (system.trace! = 0) put(trace!, system, 1), trace_on(a1))
   ;     else if (length(a) = 2 & a[2] % integer & (a[1] % ({spy} U integer)))
   ;          (system.Kernel/call_level := a[2], system.Kernel/call_trigger := a[1])
        else self_eval(self))

(let x := get(functional, self_trace @ Trace) in
  (if known?(x) put(mClaire/evaluate, Trace, x)))

// used to trace the trigger of a rule
//
trace_rule(R:relation,s:string,x:any,y:any,u:any,v:any) : void
 -> (if ((mClaire/if_write.trace! + system.verbose) >= 5)
        let p := get(ctrace, system) in
          (if known?(p) p := use_as_output(p),
           printf("--- the rule ~A is triggered for (~S,~S) by an update ~S(~S) ~A ~S \n",
                  s, u, v, R, x, (if multi?(R) ":add" else ":="), y),
           if known?(p) use_as_output(p)))


// stores a set of stopping values
stop(p:property,l:listargs) : any
  ->  (if unknown?(Core/StopProperty,p) Core/StopProperty[p] := list(l)
       else if (l = nil) put(Core/StopProperty,p,unknown)
       else Core/StopProperty[p] :add list(l),
       true)


// ******************************************************************
// *    Part 3: The debugger interface                              *
// ******************************************************************

// toggle the debug mode
debug(system) : void
 -> (if (system.Kernel/debug! != -1)
        (write(Kernel/debug!, system, -1), printf("debugger removed\n"))
     else (write(Kernel/debug!, system, 0),
           put(ctrace, system, stdout),
           printf("debugger installed\n")))

// this method is called when an error has occured. The value of index
// is recalled with last_index, so that the actual content of the stack is
// preserved.
call_debug(system) : any
 -> (let top := system.Kernel/last_debug in
       (debug_system(),
        put(Kernel/spy!, system, unknown),       ;*/
        write(step!, system, 0),
        write(trace!, system, 0),
        mClaire/set_base(system.Kernel/last_index),
        mClaire/set_index(+@integer(system.Kernel/last_index, 1)),
        put(Kernel/debug!, system, top),
        print_exception(),
        write(fromp, reader, stdin),
        write(index, reader, 0),
        when c := Language/LastCall in
            try (if (DBline[c] > 0) printf(" \n---- Last call ~S in line ~A\n",c,DBline[c]))
            catch any nil,
        *index* := 1,
        *curd* := system.Kernel/debug!,
        *maxd* := system.Kernel/debug!))

// this method is called when an error has occured. The value of index
// is recalled with last_index, so that the actual content of the stack is
// preserved.
breakpoint() : void
 -> (let top := system.Kernel/debug!,
         s := system.step!,
         t := system.trace! in
       (write(step!, system, 0),
        write(trace!, system, 0),
        *index* := 0,
        *curd* := top,
        *maxd* := top,
        if (mClaire/get_stack(top) > 0 & system.Kernel/debug! > 0)
           let j := mClaire/get_stack(top),
               num_args := (mClaire/get_stack(j + 2) - 1),
               start := (mClaire/get_stack(j + 3) as integer),
               m := mClaire/get_stack(j + 1) in
            (printf("break in ~S(~S~I) [q] >", m,
                    mClaire/get_stack(start),
                    (for i in ((start + 1) .. (start + num_args))
                       printf(",~S", mClaire/get_stack(i)))),
             let n := 1,
                 m := 1,
                 c:any := read(CommandLoop()) in
              while (c != q) (eval(c),
                              princ("break>"),
                              c := read(CommandLoop()))),
        write(step!, system, s),
        write(trace!, system, t)))

// the four keyword
up :: property()
dn :: property()
where :: property()

[dn(x:integer) : void
 -> while (mClaire/get_stack(*curd*) > 0 & x > 0)
       (*curd* := mClaire/get_stack(*curd*), *index* :+ 1, x := x - 1) ]

[up(x:integer) : void
 ->  if (x > 0)
        let indices := nil,
            ind := *maxd* in
          (while (ind != *curd*)
             (indices := ind cons indices, ind := mClaire/get_stack(ind)),
           if (x > length(indices)) (*curd* := *maxd*, *index* := 1)
           else (*curd* := indices[x], *index* :- x)) ]

// top is the top position in this stack (the last entered message)
[where(x:integer) : void
 ->  let j := *curd*,
         stack_level := 0 in
       while (j > 0 & x > 0 & system.Kernel/debug! > 0)
         (print_debug_info(j, stack_level, *index*),
          stack_level :+ 1,
          x := x - 1,
          j := (mClaire/get_stack(j) as integer)) ]

// note for interpretted methods .. they should be pushing their restriction
// on the stack vs. properties
print_debug_info(index:integer,stack_level:integer,cur_index:integer) : void
 -> (let num_args := ((mClaire/get_stack(index + 2) as integer) - 1),
         start := (mClaire/get_stack(index + 3) as integer),
         m := mClaire/get_stack(index + 1) in
       printf("debug[~A]>~I ~S(~S~I)\n", cur_index + stack_level,
              (for x in (1 .. stack_level) printf(">")), m,
              mClaire/get_stack(start),
              (for i in ((start + 1) .. (start + num_args))
                 printf(",~S", mClaire/get_stack(i)))))

// debug version of the debugger :-)
Show(n:integer) : any
 -> let i := *curd* in          // i = top of stack
       (while (i > 0 & n > 0)
         let num_args := ((mClaire/get_stack(i + 2) as integer) - 1),
             start := (mClaire/get_stack(i + 3) as integer) in
           (printf("[~A - ~A]: p = ~S, narg = ~S \n", start, i,
                   mClaire/get_stack(i + 1), num_args),
            for j in (0 .. num_args)
              printf("  [~A]:~S \n", j + i, mClaire/get_stack(j + i)),
            n :- 1,
            i := (mClaire/get_stack(i) as integer)))         // go to next block

// top is the top position in this stack (the last entered message)
//
claire/block(x:integer) : void
 -> let j := *curd*,
         stack_level := 0 in
      (while (j > 0 & x > 0 & system.Kernel/debug! > 0)
         (let nargs := mClaire/get_stack(j + 2),
              start := (mClaire/get_stack(j + 3) as integer),
              z := mClaire/get_stack(j + 1) as property,
              m := Core/find_which(z.Kernel/definition,
                                    owner(mClaire/get_stack(start)), start,
                                    start + nargs) in
          (case m (method
              printf("debug[~A] > ~S(~I)\n", *index* + stack_level, m,
                   (if (known?(formula,m) & m.formula % lambda)
                       let n := 0 in
                         for v in closure_build(m.formula)
                           (printf("~S = ~S, ", v, mClaire/get_stack(start + n)),
                            n :+ 1)
                    else printf("<compiled:~S>", m.module!))),
                   any printf("debug[~A] > ~S -> ~S\n", *index* + stack_level,z,m))),
          stack_level :+ 1,
          x :- 1,
          j := (mClaire/get_stack(j) as integer)))

// computes the list of variables of a lambda, including everything
//
closure_build(self:lambda) : list
 -> (let lvar := make_list(self.dimension, {}) in
       (for x in self.vars lvar[x.index + 1] := x,
        closure_build(self.body, lvar),
        lvar))

// give to each lexical variable its right position in the stack
// answer with the number of lexical variable
//
closure_build(self:any,lvar:list) : void
 -> (case self
      (Variable lvar[self.index + 1] := self,
       Instruction for s in self.isa.slots
                    closure_build(get(s, self), lvar),
       bag for x in self closure_build(x, lvar),
       any nil))

// ******************************************************************
// *    Part 4: Stepper, Measure & mem                              *
// ******************************************************************
// the stepper interface is quite simple and could be improved
//
[Core/call_step(pr:property) : void
 ->  printf(") : [(i)n,(o)ut,e(x)it,(t)race,(b)reakpoint]\n"),
     let m := 1,
         c := char!(step_system()),
         n := system.step! in
       (if (c = 'i') write(step!, system, n + 1)
        else if (c = 'o') (if (n > 1) write(Kernel/step!, system, n - 1))
        else if (c = 'x') error("exit stepper")
        else if (c = 't') trace_on(pr)
        else if (c = 'b') breakpoint())]

// interface
// step => trace
//
claire/step(x:any) : void
 -> (if (system.trace! = 0) write(trace!, system, 1),
     case x
      (property write(trace!, x, x.trace! + 1000),
       integer (system.Kernel/count_trigger := step, system.Kernel/count_level := x),
       environment (if (system.step! = 0)  system.step! := system.trace!
                    else system.step! := 0),
       any 0))

// New in CLAIRE 3.4 - measure objects can be stored on a file and loaded later on
// a measure is a float value counter that stores the sum & sum of squares, to 
claire/measure <: object(
  m_index:integer = 1,          // each measure has an index
  sum_value:float = 0.0,        // keep sigma(values) to return average
  sum_square:float = 0.0,       // used for standard deviation
  num_value:float = 0.0)        // number of experiments

// simple methods add, mean, stdev
[close(x:measure) : measure 
  -> (x.m_index := length(measure.instances), x)]
[add(x:measure, f:float) : measure 
  -> x.num_value :+ 1.0, x.sum_value :+ f, x.sum_square :+ f * f, x ]
[claire/mean(x:measure) : float 
  -> if (x.num_value = 0.0) 0.0 else x.sum_value / x.num_value]
[claire/stdev(x:measure) : float
   -> let y := ((x.sum_square / x.num_value) - ((x.sum_value / x.num_value) ^ 2.0)) in
         (if (y > 0.0) sqrt(y) else 0.0) ]
[claire/stdev%(x:measure) : float -> stdev(x) / mean(x) ]
[claire/reset(x:measure) : void -> x.sum_square := 0.0, x.num_value := 0.0, x.sum_value := 0.0 ]
[self_print(m:measure) : void -> printf("~F2[~F0]",mean(m),m.num_value)]


// two simple methods to store and retreive measures
//   logMeasure(s:string)  : creates a file
//   load(s:string)        : loads the files, that containts addLog(i,s,ss,n) line
[claire/logMeasure(s:string) : void
  -> let p := fopen(s,"w"),n := size(measure) in
       (use_as_output(p),
        printf("// log file produced on ~A",date!(1)),
        for m in measure 
           printf("(addLog(~A,~A,~A,~A,~A))\n",m.m_index,m.sum_value,m.sum_square,m.num_value,n),
        fclose(p)) ]

// adds a set of measures to a measure object (represented by its index)
[claire/addLog(i:integer,x:float,y:float,n:float,s:integer) : void
  -> (if (size(measure) = s)
        let m := measure.instances[i] in   // i <= s by construction
           (m.sum_value :+ x, m.sum_square :+ y, m.num_value :+ n)
      else error("logMeasure not compatible with current set (~A vs ~A)",size(measure),s)) ]


// memory usage statistics for a class (useful for debug)
mem(c:class) : integer
  -> let n := 0 in
       (for x in c.instances
          (n :+ slot_get(x,0,integer),         // size of chunk
           for sl in c.slots
             let rs := sl.range in
               (if (rs = float) n :+ 5
                else if (rs = string)
                  (when st := get(sl,x) in 
                    n :+ min(5, length(st as string) / 2))    // upper bound !
                else if (rs <= bag)
                  (when l := get(sl,x) in
                    n :+ nth_get(l as bag,0)))),    
        n)



// *********************************************************************
// *      Part 5: Profiler                                             *
// *********************************************************************

// we use a counter object for the 5 interesting values  and
// we use the reified slot to store the counter (thus no profiling on reified)
claire/PRcount <: object(rtime:integer = 0,     // time that has elapsed in the property
                         rdepth:integer = 0,    // counter of recursion (only register 1st)
                         rnum:integer = 0,      // number of calls
                         rloop:integer = 0,     // counter of loops
                         rstart:integer = 0)    // start time (1st entry)

// get & create if needed a PRcounter
claire/PRget(p:property) : PRcount
 -> let x := p.reified in
     (if (owner(x) = PRcount) (x as PRcount)
      else if (x = true) error("[131] Cannot profile a reified property ~S",p)
      else (x := PRcount(), p.reified := x, x as PRcount))

// get & create if needed a PRcounter
claire/PRlook(p:property) : any -> show(PRget(p))

// show the profiler statistics on one property
claire/PRshow(p:property) : void
 -> let x := p.reified in
      (case x (PRcount printf("~S: ~A calls -> ~A clock tics\n",p, x.rnum, x.rtime)))

// elapsed time
claire/PRtime(p:property)  : integer
 -> let x := p.reified in
      (case x (PRcount x.rtime, any 0))

claire/PRcounter(p:property)  : integer
 -> let x := p.reified in
      (case x (PRcount x.rnum, any 0))


// show the profiler statistics on the 10 most important properties
claire/PRshow() : void
 -> let l := list<property>() in
      (for p in property
        (if exists(i in (1 .. min(10,length(l))) |
                    (if ((PRtime(p) > PRtime(l[i])) |
                         (PRtime(p) = PRtime(l[i]) & PRcounter(p) > PRcounter(l[i])))
                        (l := nth+(l,i,p), true))) nil
         else if (length(l) < 10) l :add p),
       shrink(l,10),
       for p in l
          (if (PRcounter(p) > 0)
              (printf("-----------------------------------\n"),
               PRshow(p),
               for p2 in PRdependent[p]
                  (if (PRtime(p2) > 0) printf("   * ~I",PRshow(p2))))))

// reuse from lexical_build in pretty.cl
// returns the list of properties that are used by a method
PRdependent[p:property] : set[property] := set<property>()
PRdependentOf[p:property] : set[property] := set<property>()
dependents(self:method) : set[property] //
 -> set<property>{p in dependents(self.formula.body) |
                  exists(r in p.restrictions | r % method)}

// this is really cute ....   v3.2.58: fix typing
dependents(self:any) : any
 -> (case self
       (Call add(dependents(self.args),self.selector),
        Instruction let s := set<property>() in
                      (for sl in self.isa.slots
                         s := s U dependents(get(sl, self)),
                       s),
         bag let s := set<property>() in
                (for x in self s := s U dependents(x),
                 s),
         property set<property>(self),
         any set<property>()))

// used to set up the dependence
(PRdependent.inverse := PRdependentOf)
claire/PRdepends(p:property,p2:property) : void
 -> (PRdependent[p] :add p2, unknown)

// end of file
