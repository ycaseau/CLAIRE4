//+-------------------------------------------------------------+
//| CLAIRE                                                      |
//| inspect.cl                                                  |
//| Copyright (C) 1994 - 2021 Yves Caseau. All Rights Reserved  |
//| cf. copyright info in file object.cl: about()               |
//+-------------------------------------------------------------+

// --------------------------------------------------------------------
// this file contains the CLAIRE run-time tools: toplevel, inspect, trace & debug
// --------------------------------------------------------------------

// *********************************************************************
// * Contents                                                          *
// *      Part 1: Top Level                                            *
// *      Part 2: Inspection                                           *
// *      Part 3: Trace                                                *
// *      Part 4: Debugger                                             *
// *      Part 5: Measures & Profiler                                  *
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
*curd*:integer := 0                     // v3.2.14 cleaner :-) : current debug stack top
*showall*:boolean := true               // v0.01 stop the ... !

// this is the interface with the system - used to manage breakpoints
CommandLoop() : string -> function!(CommandLoopVoid)

// *********************************************************************
// *      Part 1: Toplevel                                             *
// *********************************************************************

// we use six global variables (used to be C++)
InspectStack:list := nil
TopLevelMode:integer := 1        // 1 : regular, 2: debug; 3: inspect
TopCount:integer := 0
TopIndex:integer := 0            // for debug loop, store the stack context
TopBase:integer := 0
TopDebug:integer := 0

// this is the classical print(eval(read)) LISP top level :)
// error are caught
[top_level(r:meta_reader) : void
  ->  let res:any := 0 in
        (while (res != q)
          (princ((if (TopLevelMode = 1) string!(name(module!()))
                  else if (TopLevelMode = 2) "debug" else "inspect")),
           princ("> "),
           try (r.toplevel := true,
                if (system.mClaire/count_call > 0) system.mClaire/count_call := 1,
                res := lexical_index(nextunit(r),nil,0,true),
                if (TopLevelMode = 1) system.index := 20,
                if (TopLevelMode = 3 & res != q) inspect_loop(res,InspectStack)
                else (if (TopLevelMode = 1) printf("eval[~S]> ",TopCount :+ 1)
                      else princ("> "),
                      res := eval(res),
                      if (res != q)
                         (print(res),princ("\n"))))
           catch any
             (if (let e := exception!() in (e.isa = system_error & e.index = -1)) // abort
                 (TopLevelMode := 1, res := q)
              else (mClaire/restore_state(r),        // unclear what it does 
                    if (r.external != "toplevel")
                       printf("---- file: ~I, line: ~I\n", princ(r.external), princ(n_line())), 
                    debug_if_possible(),
                    princ("\n"))),
           if (TopLevelMode != 1 & res = q)            // quit debug or inspect loop 
              (if (TopLevelMode = 2)                   // quit Debug mode
                 (system.mClaire/index := TopIndex,
                  system.mClaire/base := TopBase,
                  system.trace! := 1,
                  // printf("--------- quit debug base:~A index:~A debug:~A\n",TopBase,TopIndex,TopDebug),
                  system.mClaire/debug! := TopDebug),
              res := unknown,
              TopLevelMode := 1))) ]
  //        exit(1)) ]

// start a debug loop - aha 
[debugLoop(r:meta_reader) : void
 ->  TopDebug := 0,
     TopBase := system.mClaire/base,
     TopIndex := system.mClaire/index,
     r.toplevel := true,
     TopLevelMode := 2,
     princ("--------------- Debug -------------------\n") ]

// starts an inspector  on a list
[inspect_system(l:list) : void
  ->  InspectStack := l,
      if (TopLevelMode = 2) system.trace! := 1,
      TopLevelMode := 3 ] // INSPECT   
  
// simple main (to be enriched later)
[simple_main() : void
  -> top_level(reader),
	   printf("[regular exit] Bye.\n") ]

// unclear that we need this (simple_main is used with -cm)
/*
  ->  let %init? := true, l := (copy(params()) as list<string>) in  // args list
    (try
      (while (l)
       (case l[1]
        ({"-s"}  (if (length(l) >= 2)  l :<< 2 else error("option: -s <s1> <s2>")),
         {"-f"}  (if (length(l) >= 2)  (load(l[2]), l :<< 2)
                  else error("option: -f <filename>")),
         {"-m"}  (if (length(l) >= 2)
                 (if %init? (load("init"), %init? := false),
                  let m := value(l[2]) in
                     (if not(m % module) error("~S is not a module",l[2]),
                      load(m), begin(m), l :<< 2))
                  else error("option: -m <module>")),
         {"-n"} (%init? := false, l :<< 1),
         any (if (l[1][1] = '-') (printf("~S is an unvalid option\n",l[1])),
              l := list<string>()))),
       if %init? load("init"))
       catch any (mClaire/restore_state(reader),
                  Reader/debug_if_possible()),
     top_level(reader),
	   printf("[regular exit] Bye.\n")) ]  */


// *********************************************************************
// *      Part 2: Inspection                                           *
// *********************************************************************
// this is the method that the user calls
//
inspect(self:any) : any
 -> let m0 := module!(), ix := 0 in
       (case self
         (list for i in (1 .. length(self)) printf("~A: ~S\n", i, self[i]),
          object for rel in owner(self).slots
                  let m := module!(name(selector(rel))) in
                    (ix :+ 1,            // ix is the position in the slots list
                     if (m = m0 | m = claire | *showall*)
                       let val := get(rel,self) in
                        printf("~A: ~S = ~I\n", ix, rel.selector,
                           (case val (list (if (length(val) < 10) pretty_print(val)
                                           else (pretty_print(
                                                   list{val[i] | i in (1 .. 9)}),
                                                 princ("..."))),              // v4.0.6 
                                      any pretty_print(val))))),
          any printf("~I\n", pretty_print(self))),
        inspect_system(list(self)),
        None)

// this is the inspect top_level
//
inspect_loop(%read:any,old:list) : void
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
 -> (if (self % list)
        (if (n > 0 & n <= length(self as list)) self[n]
         else (printf("~A in not a good index for ~S.\n", n, self), self))
     else let l := owner(self).slots in
            (if (n > 0 & n <= length(l)) get@slot(l[n], self)
             else (printf("~A is not a good index for ~S.\n", n, self),
                   self)))

// *********************************************************************
// *      Part 2: Trace methods                                        *
// *********************************************************************
// instrument the code generated from the rules

// this is the control method to CLAIRE tracer
// notice that trace(where) activates the call_count
trace_on(self:any) : any
 -> (case self
      (property (if (self = spy)
                   let m := (spy @ void) in (if m put(Kernel/spy!, system, m))      ;*/
                 else if (self = where) system.Kernel/call_count := 1
                 else put(trace!, self, 5 - system.verbose)),          // makes trace(p) active
       environment put(trace!, system, 1),                             // activates tracing
       module (if (self.mClaire/status > 2) self.mClaire/status := 4,
               for m in self.parts trace_on(m)),
       port put(ctrace, system, self),                                 // sets the trace output port
       string put(ctrace, system, fopen(self, "w")),
       integer put(verbose,system,self),                               // alias for verbosity
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


// used to trace the trigger of a rule
trace_rule(R:relation,s:string,x:any,y:any,u:any,v:any) : void
 -> (if ((mClaire/if_write.trace! + system.verbose) >= 5)
        let p := get(ctrace, system) in
          (if known?(p) p := use_as_output(p),
           printf("--- the rule ~A is triggered for (~S,~S) by an update ~S(~S) ~A ~S \n",
                  s, u, v, R, x, (if R.multivalued? ":add" else ":="), y),
           if known?(p) use_as_output(p)))

// stores a set of stopping values
// this is a cool feature : stop(p, list(a1,a2)) => p(x,y) will stop if x = a1 and y = a2
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
           put(jito?,system,false),               // v4.0.6 : debug prevents jito (tracability)
           printf("debugger installed\n")))

// this method is called when an error has occured. The value of index
// is recalled with last_index, so that the actual content of the stack is
// preserved.
call_debug(system) : any
 -> (let top := system.Kernel/last_debug in
       (debugLoop(reader),
        put(Kernel/spy!, system, unknown),       ;*/
        write(trace!, system, 0),
        mClaire/set_base(system.Kernel/last_index),
        mClaire/set_index(+@integer(system.Kernel/last_index, 1)),
        put(Kernel/debug!, system, top),
        if (system.verbose > -1) print_exception(),  // meta-debug hook
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
         t := system.trace! in
       (write(trace!, system, 0),
        *index* := 0,
        *curd* := top,
        *maxd* := top,
        if (mClaire/get_stack(top) > 0 & system.Kernel/debug! > 0)
           let j := mClaire/get_stack(top) as integer,
               num_args := ((mClaire/get_stack(j + 2) as integer) - 1),
               start := (mClaire/get_stack(j + 3) as integer),
               m := mClaire/get_stack(j + 1) in
            (printf("break in ~S(~S~I) [q] >", m,
                    mClaire/get_stack(start),
                    (for i in ((start + 1) .. (start + num_args))
                       printf(",~S", mClaire/get_stack(i)))),
             let c:any := read(CommandLoop()) in
              while (c != q) (eval(c),
                              princ("break>"),
                              c := read(CommandLoop()))),
        write(trace!, system, t)))

// the four keyword
up :: property()
dn :: property()
where :: property()

[dn(x:integer) : void
 -> while (mClaire/get_stack(*curd*) as integer > 0 & x > 0)
       (*curd* := mClaire/get_stack(*curd*) as integer,
        *index* :+ 1, x := x - 1) ]

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

// debug version of the debugger :-)  => use as Reader/Show(n)
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
       list for x in self closure_build(x, lvar),
       any nil))

// ******************************************************************
// *    Part 5:  Measure &  Profile                                 *
// ******************************************************************


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



// end of file
