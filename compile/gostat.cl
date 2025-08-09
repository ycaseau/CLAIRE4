//+-------------------------------------------------------------+
//| CLAIRE                                                      |
//| gostat.cl                                                   |
//| Copyright (C) 2020-2023   Yves Caseau. All Rights Reserved  |
//| cf. copyright info in file object.cl: about()               |
//+-------------------------------------------------------------+

// statement is implemented as a general method that calls a restriction
//        g_statement(self:any,e:class,v:string,err:boolean,loop:any)
// (1) e is the goType that the variable v must receive (HENCE goCast must be inserted)
//     a proper goType is a class, or EID, or void
// (2) The argument v is the named of the C variable in which the
//     result of the evaluation must be placed.
// (3) err tells if an error is possible, which forces to create a chain an not a block (see Do for example)
//     Note : if err = true, s is expected to be EID to (a) force a chain (b) place the error value in v
// (4) loop is either false (not within a loop) or a tuple(v,s) inside the compiling of While/For
//     This tuple describes the vreturn Variable in case a break(v) is encountered

// there are two possible outputs: blocks (lines of code without {}, used to be call inner_statement)
// and chains  (we use chains to denote long nested ifs that manage error handling)

// indentation : 
//    we call statement(s) at the proper current indentation level => it produices n lines with the indentation
//    and stop after a break line, at the proper indentation level


//**********************************************************************
//*  Table of contents:                                                *
//*          Part 1: Unfolding of complex expressions                  *
//*          Part 2: Error Management and EID Unfolding                *
//*          Part 3: Basic control structures                          *
//*          Part 4: iteration                                         *
//*          Part 5: CLAIRE-specific structures                        *
//**********************************************************************

//*************************************************************************
//*          Part 1: Unfolding of complex expressions                     *
//*************************************************************************

// when local CLAIRE expressions are not go expression, we need to unfold the global expression into a big Let
// HOWEVER, if only works for list of arguments whose evaluation order is not specified ! (because we move some of the evaluations earlier)
// this reentrant compiling (calling g_statement on a expanded Let) works because Let checks if g_expression can be used
// the same pattern is used for call_slot/call_table

// this function is used to unfold complex expressions that should be compiled as
// expressions and not statements. It takes a list of arguments l and returns the
// embedded Lets that defines the necessary variable or nil (nothing is needed)
// this list is of the form  (a1 .. am) where
//     m is the number of statement args in l
//     ai is a Let that defines the i-th variable corresponding to the i-th bad guy
// CLAIRE 4: we unfold args that are not functional or args that can throw error
unfold_args(l:list) : list
  ->  let lbad := list{i in (1 .. length(l)) | not(g_clean(l[i])) } in // list of indices
        (//[5] unfold -> ~S : ~S // list{l[x] | x in lbad}, list{stupid_t(l[x]) | x in lbad},
         list{ Let(var = build_Variable(genvar("arg_"), static_type(l[i])),
                   value = l[i]) | i in lbad}   )

// uses the previous list to use the variable instead of the Fold.
// l is the list of arguments, ld is the previously build unfold_args(l)
unfold_arg(l:list,ld:list,x:any) : any
 -> let i := 1, j := 0, m := length(l) in
       (if (while (i <= m)
             (if not(g_clean(l[i]))
                 (j :+ 1, (if (l[i] = x) break(true)))
              else if (l[i] = x) break(false),
              i :+ 1))
           var(ld[j])
        else l[i])

// creates the Let from the ldef definition and places the statement x in the body
// note that the error handling is done in the Let (with g_statement)
// x is the call form where the variable has been replaced if needed
unfold_use(ldef:list,x:any,s:class,v:string,err:boolean,loop:any) : void
 -> (if eid_require?(x) unfold_eid(ldef,x,s,v,err,loop)
     else let  n := length(ldef), vb := verbose() in
      (verbose() := 0,                  // v3.1.06
       if not(ldef) error("[internal] design bug g_func(~S) should be true",x),
       for i in (1 .. (n - 1)) ldef[i].arg := ldef[i + 1],
       arg(ldef[n] as Let) := x,
       verbose() := vb,                  // not very elegant !
       //[5] unfold use: ~S (tail = ~S)// ldef[1],x,
       g_statement(ldef[1] as Let,s,v,err,loop)))


//*************************************************************************
//*          Part 2: Error Management and EID Unfolding                   *
//*************************************************************************

// this is the error catching pattern: evaluate(self) and check if error then place it in vglobal,
// if no error we want the value in v with expected gotype e (a true gotype = class)
// if v is an EID variable, do not create an extra variable (we use it temporarily)
// in a loop we generate a break to exit to loop
// v is the variable that must receive self
// note : g_try produces a pattern   <e = code>, if Err(e) {res =e} else { ...
// that must be closed } with a close_try => and nothing after the close_try (nothing must if an error occurred)
[g_try(self:any,v:string,e:class,vglobal:string,loop:any) : void
  -> let v2 := (if (e = EID) v else genvar("try_")) in 
        (if (e != EID) var_declaration(v2,EID,1),
         if PRODUCER.debug? printf("/*g_try(v2:~S,loop:~S,e:~S) */~I",v2,loop,e,breakline()),
         g_statement(self,EID,v2,true,loop),
         // AUDACIEUX: if self is a Do, and we have a loop, break statements cover the error case
         if (self % Do & loop % Tuple) printf("{~I",breakline())
         else
           (if PRODUCER.debug? printf("/* ERROR PROTECTION INSERTED (~A-~A) */~I",v,vglobal,breakline()),   // for debug
            if (v = vglobal & e = EID & not(loop % tuple))   // simpler since the value is already in vglobal !
              printf("if !ErrorIn(~I) {~I", c_princ(v2), breakline())
           else (printf("if ErrorIn(~I) {", c_princ(v2)),
               if (v != vglobal) printf("~I = ~I~I",c_princ(vglobal), c_princ(v2), breakline()),
               if (loop % tuple) 
                  (if (loop[1] !=  vglobal) printf("~I = ~I~I",c_princ(loop[1]),c_princ(v2),breakline()),
                   princ("break"), 
                   breakline()), 
               printf("} else {~I",breakline()))),
         if (e != void & v != v2)                           // place the right go form in v
              printf("~I = ~I~I", c_princ(v),from_eid(PRODUCER,v2,e),breakline())) ]

// when the error is nested in the expression, the unfold pattern will make sure that we separate the sub_exp that may
// create the error, but assignment is not managed this way, hence this code to avoid double error check
[g_try(self:Assign,v:string,e:class,vglobal:string,loop:any) : void
  -> let %var := self.var,  v1 := c_string(PRODUCER,%var), %range := class!(range(%var)) in
        (g_try(self.arg,v1,%range,vglobal,loop),
         if (e != void) printf("~I = ~I~I",c_princ(v),use_variable(v1,e,%range),breakline()))]
   
// each g_try produces a {, which we must balance before returning a new line 
// does NOT change OPT.level !
[close_try(n:integer) : void 
  ->  for i in (1 .. n) princ("}"),
      if (n > 0) breakline()]        

// special case when v is a g_func that can produce an error (s is assumed to be EID)
[error_wrap(self:any,s:class,v:string) : void 
  -> if (s != EID)  
        (trace(0,"---- g_throw(~S) = ~S\n",self, g_throw(self)),
         error("design bug for error_wrap with ~S and s = ~S",self,s)),
     printf("~A = ~I~I",v,g_expression(self,EID),breakline()) ]

// this is a special case when the statement result is not needed (e = void) so we should
// not reuse v as a temporary variable (which we considered) 
// this is called inside a For/While, so loop is a tuple
[g_try_void(self:any,vglobal:string,loop:any) : void   
   -> if (case self (Assign (range(self.var) = EID)))    // pragma v:EID was used
        let %var := self.var, v1 := c_string(PRODUCER,%var) in
             g_try(self.arg,v1,EID,vglobal,loop)
      else let v2 := genvar("loop_") in  
        (// printf("/* try_void(~S) : throw = ~S */~I",self,g_throw(self), breakline()),
         var_declaration(v2,EID,1),       // this is the EID variable for error check (assign to vglobal only if error)
         printf("_ = ~A~I", v2,breakline()),     // sometimes this variable is not used
         case self
           (Do  (princ("{ "), 
                 breakline(),                      // g_try opens one { + no change on OPT
                 do_statement(self,EID,v2,true,loop,false)),   // no need for the last Do value
            any g_try(self,v2,EID,vglobal,loop))) ]

// eid_require means that the internal call should better be evaluated in EID mode
// this is really what we need for mClaire/push!(eval(x)) and funcall(f,...) ... but has been be extended to methods
// that do a better job (no allocation) in EID mode
eid_require?(x:any) : boolean
  -> (case x (Call x.selector = mClaire/push!, 
              Call_method (x.arg.selector = funcall | x.arg = *write_value* | x.arg = *read_property* | 
                           x.arg.selector = write_fast | x.arg.selector = nth_write),
              any false))

// eid_provide? says that the call will produce first an EID
eid_provide?(x:any) : boolean
  -> (case x (Call x.selector = mClaire/get_stack,        // gets and EID directly
              Call_method (x.arg.selector = eval | x.arg = m_unsafe),        // eval returns an EID :)
              Variable (x.range <= integer),              // avoid MakeInteger() conversion+alloc
              integer true,
              // to_CL eid_provide?(x.arg),                  // to get rid of ! 
              any false))


// eid_unfold could use a more general "EID compiling mode" (with a list of EID variables passed as context)
// this is a quickfix => we build the EID Let on our own (code borrowed from g_stat@Let)
unfold_eid(ldef:list,self:any,s:class, v:any,err:boolean,loop:any) : void
  -> let n := length(ldef), lvar := list<Variable>(), count_try := 0 in
        (new_block("LetEID"),
         for i in (1 .. n)
           let %l := (ldef[i] as Let), v2 :=  c_string(PRODUCER,%l.var), 
               x := %l.value, try? := g_throw(x) in 
            (var_declaration(v2,EID,0), 
             lvar :add %l.var,
             breakline(),
             if  try? 
                (count_try :+ 1,
                 g_try(x,v2,EID,v,false))          // if the value may be an error => start if chain
             else  statement(x, EID, v2,loop)),
         if (s != void) printf("~A = ", v),
         eid_expression(self,EID,lvar),                 // compiles the expression self in EID mode
         close_try(count_try),
         close_block("LetEID"))

 // eid_expression compiles a call or call_method with one EID variable
 // it performs all the compiler optimization (see the eid_fold? pattern in gostat.cl)
 [eid_expression(x:any,s:class,lvar:list<Variable>) : void
    -> case x
        (Variable (if (x % lvar & s = EID) princ(c_string(PRODUCER,x))
                   else g_expression(x,s)),
         Call printf("ClEnv.Push(~I)",eid_expression(x.args[1],EID,lvar)),
         Call_method (if (x.arg.selector = funcall)
                        printf("FASTCALL~A(~I,~I~I~I)", (length(x.args) - 1),   // 1,2 or 3
                         eid_expression(x.args[1],method,lvar), 
                         eid_expression(x.args[2],EID,lvar),
                         (if (length(x.args) >= 3) (princ(","), eid_expression(x.args[3],EID,lvar))),
                         (if (length(x.args) = 4) (princ(","), eid_expression(x.args[4],EID,lvar))))
                      else if (x.arg = *read_property*)
                          printf("~I.ReadEID(~I)",g_eid_expression(x.args[1],property,lvar),
                                    eid_expression(x.args[2],EID,lvar))
                     else if  (x.arg.selector = write_fast)
                          printf("~I.WriteEID(~I,~I)",g_expression(x.args[1],property),
                                    g_eid_expression(x.args[2],object,lvar),
                                    eid_expression(x.args[3],EID,lvar))
                     else if  (x.arg.selector = nth_write)
                          printf("~I.WriteEID(~I,~I)",g_eid_expression(x.args[1],list,lvar),
                                    g_eid_expression(x.args[2],integer,lvar),
                                    eid_expression(x.args[3],EID,lvar))
                      else  printf("~I.WriteEID(~I)",g_eid_expression(x.args[1],Variable,lvar),
                                   eid_expression(x.args[2],EID,lvar))),
         any g_expression(x,s)) ]

// reverse from eid (the args of the call have been EIDed : represented by an EID var)
// hence when we need a regular object, we must check that the arg x is not in the var list
[g_eid_expression(x:any,s:class,lvar:list<Variable>) : void
    -> case x
        (Variable (if (x % lvar) (eid_prefix(s),
                                  princ(c_string(PRODUCER,x)),
                                  eid_post(s))
                   else g_expression(x,s)),
         any g_expression(x,s)) ]

//**********************************************************************
//*          Part 3: Basic control structures                          *
//**********************************************************************

// The re-entry definition (called within g_statement, not directly)
// if functional, the best compiling is into an expression
// s is the expected go type (as a class) + void + EID
// v is nil or a string (name of the variable)
// note that only 3 additional parameters are used since err is recomputed
[statement(self:any,s:class,v:string,loop:any) : void
 -> if g_clean(self)
       (if (s != void & static_type(self) != void)
           printf("~I = ~I~I", c_princ(v),  g_expression(self, s), breakline())
       else if (self % If) g_statement(self, s, v, false,loop)
       else if (self % delimiter)
             error("[201] Loose delimiter in program: ~S", self)
       else (// s = void or static_type = void
             stat_exp(PRODUCER,self,void),    // new: v3.0.60: include a few tricks
             if (s = EID) printf("~I = EVOID~I", c_princ(v),breakline())))
    else g_statement(self, s, v, g_throw(self),loop) ]


// make a statement from an expression (in C++ we need a ; - with go a breakline is enough)
// we do not want to place the result in a variable (see upper)
// will get simpler once we have a stable compiler without to_C and to_CL
[stat_exp(c:go_producer,self:any,s:class) : void
  -> // if (self % Generate/to_C) stat_exp(c, (self as Generate/to_C).arg, s)
     // else if (self % Generate/to_CL) stat_exp(c, (self as Generate/to_CL).arg, s)
     if Compile/designated?(self)  breakline()                        // designated should be ignored if unused
     else printf("~I~I", g_expression(self, s), breakline()) ]


// a DO is a simple go block if there are not errors, a chain otherwise
// the chain means multiple nestings when an error occurs since the rest of the DO must not be
// this is why the close_try(count) are called at the end, to close the embedded ifs (ErrorIn(e))
// we use a specific method code_statement with an additional parameter %need which is true by default
[g_statement(self:Do,e:class,v:string,err:boolean,loop:any)
 -> do_statement(self,e,v,err,loop,true)]

// %need is true = the last argument is needed in v
[do_statement(self:Do,e:class,v:string,err:boolean,loop:any,%need:boolean)
 -> if (length(self.args) = 1) statement(self.args[1], e,v, loop)
    else if not(err) 
       let l := self.args, m := length(l) in 
           (for n in (1 .. m)
               (statement(l[n],(if (n = m) e else void),v,loop)))
     else let l := self.args, m := length(l), count_if := 0 in 
           (for n in (1 .. m)
              let x := l[n] in
                 (if g_throw(x)                  // an error may occur => trap it in a loop
                    (if (n < m | loop % tuple)
                       (count_if :+ 1,                      // balance parenthesis later
                        g_try(x,v,EID,v,loop))              // ok to use v + trap 
                     else g_statement(x, EID, v, true, loop))     // v needed but no trap
                  else statement(x,(if (n = m & %need) e else void),v,loop)),
            close_try(count_if)) ]  // balance 

   
// a Let is a local variable declaration 
// in CLAIRE 4, a block is anything that fits between {} hence inner/outer is not necessary
// AXIOM if err is true, we require that e = any
[g_statement(self:Let,e:class,v:string,err:boolean,loop:any) : void
  -> if let_eid?(self) g_eid_stat(self,e,v,err,loop)
     else if (self.arg = self.var & e = void)        // stupid case where             
         statement(self.value,e,v,loop)              // self.var is not needed !                          
     else let ns := c_string(PRODUCER,self.var.mClaire/pname) in
      (if (ns[1] = 'C' & ns[2] = '%') self.var.mClaire/pname := gensym(),     // used in Iterate (C% variables are expanded): ocontrol.cl
       let v2 := c_string(PRODUCER,self.var),  x := self.value,
           f := g_clean(x), try? := g_throw(x), ev := class!(self.var.range) in
              (let_block(),
               var_declaration(v2,ev,0),                    // no trailing " " in v4.0.6
               if f printf(" = ~I", g_expression(x,ev)),    // 
               breakline(),
               // printf("/* noccur = ~A */~I",Language/occurexact(self.arg, self.var),breakline()),   // occurexact should discard setup !
               if (Language/occurexact(self.arg, self.var) < 1)          // avoid unused variable error (1 safe, 0 optimized)
                  (// THIS SHOULD BE A PROPER WARNING ==============
                   //[5] >>>>>>>>  variable ~S declared but unused  // v2,          
                   printf("_ = ~A~I", v2,breakline())),
               if  try? g_try(x,v2,ev,v,false)                           // if the value may be an error => start if chain
               else if not(f) statement(x, ev, v2,loop),
               statement(self.arg, e, v, loop),                          // calling statement is critical for reentrant pattern :)
               if try? close_try(1),
               close_block("Let"))) ]                           // then we must close the chain

// detect a pattern (that could be expanded) where the variable should be compiled as EID because
// the value may trigger an error and the body uses v once at the end (this could be : the body is EID-friendly)
// this current specific pattern is introduced to optimize eval_message
[let_eid?(self:Let) : boolean 
  -> let v := self.var, y := self.arg in 
        (((g_throw(self.value) | eid_provide?(self.value)) &     // special pattern
         Language/occurrence(self.arg, self.var) = 1 &
         (case y (Do not(g_throw(y)) & last(y.args) = v)))   |
         range(v) = EID)    ]                                     // force EID compiling : back door :)

// the corresponding compiling (embeds the Do)
[g_eid_stat(self:Let,e:class,v:string,err:boolean,loop:any) : void
  -> let v2 := c_string(PRODUCER,self.var),  x := self.value, try? := g_throw(x) in
     (new_block("LetE"),
      var_declaration(v2,EID,0), 
      breakline(),
      if try? g_try(x,v2,EID,v,false)
      else printf("~A = ~I~I",v2,g_expression(x,EID),breakline()),
      let y := (self.arg as Do), n := length(y.args) in
         (for i in (1 .. (n - 1))  statement(y.args[i], void,v,loop), 
          printf("~A = ~A",v,v2)),
      if try? close_try(1),
      close_block("LetE")) ]                           // then we must close the chain


// makes a bag from a list of  statements.
// The value cannot be ignored: it is considered as an error (a do should have been used)
// there are two patterns depending if self.of is known : MakeEmptyX(t) or MakeEmptyX(any)
[g_statement(self:Construct,s:class,v:string,err:boolean, loop:any) : void
  -> if not(v) error("[202] A do should have been used for ~S", self),
     let v2 := "v_bag_arg",
         kind := (case self (List list, Set set, Tuple tuple, any error("CONSTRUCT BUG: ~S", self))),
         count_try := 0,                         // count the number of args that may throw an error
         t := (if known?(of,self) self.of else void) in
       (//  [0] stat @ Construct self = ~S, v = ~S // self,v,
        new_block("Construct"),                                           // the need for a new block is debatable, but accepted 
        if exists(x in self.args | not(g_clean(x))) var_declaration(v2,any,1),
        printf("~I= ~I~IEmpty~I()~I", c_princ(v),
               cast_prefix(kind,s),
               (if (kind = tuple) princ("Make")        // MakeEmptyTuple()
                else ( if known?(of,self)  g_expression(c_code(t,object),type) else  princ("ToType(CEMPTY.Id())"),
                       princ("."))),                  // t.EmptyList()
                cap_short(kind.name),
                cast_post(kind,s)),
        for x in self.args
          let f := g_clean(x), try? := g_throw(x) in
             (breakline(),
              if try? count_try :+ 1,
              if not(f) (if try? g_try(x,v2,any,v,false) 
                         else statement(x, any, v2,loop)),
              printf("~I~I~I.AddFast(~I)",cast_prefix(s,kind), c_princ(v),cast_post(s,kind),
                   (if f g_expression(x, any) else c_princ(v2)))),
        close_try(count_try),
        close_block("Construct")) ]


// A if is easy to compile. We check if the logical compiler can be used
// we now assume that the test returns a boolean !
// note that in GO the "} else " pattern is tricky
 [g_statement(self:If,s:class,v:string,err:boolean,loop:any) : void
 ->  let try? := g_throw(self.test) in
       (if g_clean(self.test)                                       // easy case 
          (printf("if ~I ", b_expression(PRODUCER,self.test, true)),
                  new_block("If"))
        else let v2 := c_string(PRODUCER,gensym() /+ "I") in       // use intermediary variable
          (var_declaration(v2, boolean, 1),
           if try? g_try(self.test,v2,boolean,v,false)            //  [A]
           else  statement(self.test, boolean, v2, loop),
           printf("if (~A == CTRUE) ~I",v2,new_block("If"))),
        statement(self.arg,s,v, loop),
        if (self.other = nil | (self.other = false &  s = void)) close_block("If")
        else if (self.other % If & g_func(self.other.test) & not(g_throw(self.other.test)))
              printf("~I else ~I",finish_block("If"),g_statement(self.other,s,v,false,loop))
        else  printf("} else {~I~I~I",
                     breakline(),
                     statement(self.other,s,v,loop),
                     close_block("If")),
        if try? close_try(1)) ]
     



// --------------- logical combinations and/or -------------------------------

// note: we cannot use unfolding because the order of evaluation is important !
// AND is compiled with IF: as soon as an argument is false, the result is false.
[g_statement(self:And,s:class,v:string,err:boolean,loop:any) : void
  -> let v2 := check_var("v_and"), count_try := 0 in
       (new_block("and"),
        var_declaration(v2,boolean,1),
        breakline(),
        for x in self.args
              let try? := g_throw(x) in
                (if try? 
                     (g_try(x,v2,boolean,v,false), count_try :+ 1)
                 else statement(x,boolean,v2,loop),
                 printf("if (~I == CFALSE) {~I~I} else ~I", c_princ(v2),
                         (if (s != void) printf("~I = ~ICFALSE~I", c_princ(v),
                                                 cast_prefix(boolean,s),
                                                 cast_post(boolean,s))),
                         breakline(),
                         new_block("arg"))),
        if (s != void) printf("~I = ~ICTRUE~I", c_princ(v),
                                cast_prefix(boolean,s),
                                cast_post(boolean,s)),                 
        for x in self.args close_block("arg"),
        close_try(count_try),
        close_block("and")) ]

// same thing for OR
[g_statement(self:Or,s:class,v:string,err:boolean,loop:any) : void
 -> let v2 := check_var("v_or"), count_try := 0 in
       (new_block("or"),
        var_declaration(v2,boolean,1),
        breakline(),
        for x in self.args
              let try? := g_throw(x) in
                (if try? 
                     (g_try(x,v2,boolean,v,loop), count_try :+ 1)           // if loop ... it must be used !
                 else statement(x,boolean,v2,loop),
                 printf("if (~I == CTRUE) {~I~I} else ~I", c_princ(v2),
                         (if (s != void) printf("~I = ~ICTRUE~I", c_princ(v), 
                                                 cast_prefix(boolean,s),
                                                 cast_post(boolean,s))),
                         breakline(),
                         new_block("or"))),
        if (s != void) printf("~I = ~ICFALSE~I", c_princ(v),
                                cast_prefix(boolean,s),
                                cast_post(boolean,s)),              
        for x in self.args close_block("org"),
        close_try(count_try),
        close_block("or")) ]


// Here this is the simple assignment, with a true variable
// note that the last line (assigning the value to result is only OK if no error)
[g_statement(self:Assign,s:class,v:string,err:boolean,loop:any) : void
 -> let %var := self.var, x := self.arg, 
        v2 := c_string(PRODUCER,%var), %range := class!(range(%var)),
        try? := g_throw(x) in
       (if try? 
         (g_try(x,v2,%range,v,false),
          if (s != void) printf("~I = ~I~I",c_princ(v),use_variable(v2,s,%range),breakline()),
          close_try(1))
        else 
          (statement(x,%range,v2,loop),
           if (s != void) printf("~I = ~I~I",c_princ(v),use_variable(v2,s,%range),breakline()))) ]
   
// This is the global variable assignment - global variables exist in go so this is pretty simple
// note that the tricky part is the store management
// v4.0.4 : if nativeVar, we need to produce the go object, not an any  (any is now replaced by %srange)
[g_statement(self:Gassign,s:class,v:string,err:boolean,loop:any) : void
   ->  let %var := self.var, x := self.arg, 
           %range := (if nativeVar?(%var) getRange(%var) else any)  in
           (if (g_func(x) & s = void & not(%var.Kernel/store?))    // simple case
               printf("~I = ~I~I", globalVar(PRODUCER,%var), g_expression(x,%range), breakline())   
            else let v2 := genvar("v_gassign"), try? := g_throw(x) in
              (if (not(try?) & (s = any & %range = any)) v2 := v        // save intermediate the variable
               else var_declaration(v2,%range,1),
               if try? g_try(x,v2,%range,v,false)
               else statement(x,%range,v2,loop),
               if self.var.Kernel/store?  
                   printf("~I.StoreObj(3,~I,CTRUE)~I", thing_ident(%var), c_princ(v2),breakline())
               else printf("~I = ~I~I", globalVar(PRODUCER,%var), c_princ(v2),breakline()),
               if (s != void & v != v2) 
                   printf("~I = ~I~I",c_princ(v),use_variable(v2,s,%range),breakline()),
               if try? close_block()))]


//**********************************************************************
//*          Part 3: iteration                                         *
//**********************************************************************

// we know to iterate sets or lists in Go
// the optimizer should give use something that is properly typed
[bag_class(self:any) : class
  -> let s := c_type(self) in
       (if (s <= list | s <= tuple | s <= array) list
        else if (s <= set) set
        else error("bag_class(~S) returns ~S: cannot use in for",self,s))
]

// generates the iteration code for a "for x in S ..." expression , once
// all optimization based on code substitution have been performed.
// very nice in go, except that we have to handle error
// if g_member(%set) is native (anything but any) we use the native go form
[g_statement(self:For,s:class,v:string,err:boolean,loop:any) : void
-> let v2 := c_string(PRODUCER,self.var), count_try := 0,
       v2_range := class!(self.var.range),                                 // what claire expects for v2
       v3 := v2 /+ "_support",                                              // v3 : variable that holds the bag
       v4 := v2 /+ "_iter",                                                 // *Any variable for iteration
       %set := self.set_arg,                                                // set expression
       sbag := bag_class(%set), smember := g_member(%set),                  // list or set + member_type
       %direct := (v2_range = any | smember = integer | smember = float) in     // direct means v4 not necessary
      (new_block("For"),
       var_declaration(v2,v2_range,2),                                        // 2:Special : _ = v needed
       if not(%direct) var_declaration(v4,any,1) else v4 := v2,                // do not use v4 (v2 is filled directly)
       if (s != void) printf("~I= ~ICFALSE~I~I", c_princ(v), 
                              cast_prefix(boolean,s),
                              cast_post(boolean,s),
                              breakline()),  // v3.3.42 - Sylvain's optim 
       count_try :+ iteration_statement(self,%set,sbag,smember,v,v3,v4),
       if not(%direct) printf("~A = ~I~A~I~I",v2,cast_prefix(any,v2_range),v4,
                                   cast_post(any,v2_range),breakline()),
       if g_throw(self.arg)
          (g_try_void(self.arg, v, tuple(v,s)),     // loop = context for inner compiling
           count_try :+ 1)
       else statement(self.arg, void, v,tuple(v,s)),     // ... set the loop parameter to vreturn 
       close_try(count_try),
       close_block("loop"),      // for loop content
       close_block("For")) ]    // for statement

// iteration_statement produces the bulk of the iteration code
// returns 1 if we use a try/pattern for error protection
[iteration_statement(self:For,%set:any,sbag:class,smember:class,v:string,v3:string,v4:string) : integer
 -> if (g_clean(%set) & designated?(%set) & (smember != any & sbag = list))         // simple forms for list (%set is used once)
          (printf("for _,~I = range(~I.~I)~I", c_princ(v4),                           // typed list iteration pattern
                   g_expression(%set, sbag),                                        // notice that v4 occurs once
                   valuesSlot(smember),                                      // access through Values*()
                   new_block("loop")),
           0)
    else let try? := g_throw(%set) in
            (var_declaration(v3,sbag,1),                                            // v3 will contain the bag
             if try? g_try(%set,v3,sbag,v,false)                                    // use g_try, will return 1 to balance the }
             else statement(%set, sbag, v3, false),
             if (sbag = set)                                                  // set iteration 
                printf("for i_it := 0; i_it < ~I.Count; i_it++ ~I~I",  c_princ(v3), 
                     new_block(),
                     (if (smember = integer | smember = float)
                        printf("~I = ~I.~I[i_it]~I", c_princ(v4), c_princ(v3),valuesSlot(smember),breakline())
                      else printf("~I = ~I.At(i_it)~I", c_princ(v4), c_princ(v3),breakline())))
             else if (g_member(%set) != any)                                 // use native pattern for list
                printf("for _,~I = range(~I.~I)~I", c_princ(v4), c_princ(v3), 
                         valuesSlot(smember),
                         new_block("loop2"))
            else let v5 := c_string(PRODUCER,self.var) /+ "_len" in       // length of bag, used for regular pattern for complex list expr
              (printf("~I := ~I.Length()~I", c_princ(v5),c_princ(v3),breakline()),
               printf("for i_it := 0; i_it < ~I; i_it++ ~I~I",  c_princ(v5), 
                     new_block(),
                     printf("~I = ~I.At(i_it)~I", c_princ(v4), c_princ(v3),breakline()))),
            if try? 1 else 0) ]
       


// here the value is expected to be important, otherwise an error is raised.
// THIS IS ONLY APPLIED TO COLLECT(f(x) | s in S) on lists => Image is delt with
// we currently do not use the native form => use At and Put to work on generic lists
[g_statement(self:Iteration,s:class,v:string,err:boolean,loop:any) : void
-> if (s = void) error("[203] you should have used a FOR ere:~S", self),         // we shall use v to place the new list !
    let v2 := c_string(PRODUCER,self.var),         // iteration variable x with proper type
        v2_range := class!(self.var.range),
        vlist := check_var("v_list"),       // where we place S
        vlocal := check_var("v_local"),     // where we place f(x)
        bag_type := (if (c_type(self.set_arg) <= set) set else list),  // assumes here that set_arg is a set or a list 
        try_count := 0 in   // 
      (//[5] COOL : an Iteration for g_statement:~S // self,
       new_block("Iteration"),
       var_declaration(vlist,bag_type,1),         // the list that is being built
       var_declaration(v2,v2_range,                // variable associated to f(x)
          (if (Language/occurexact(self.arg, self.var) < 1) 2 else 1)),   // v4.10 : check if x unused    
       var_declaration(vlocal,any,1),         // need a *ClaireAny
       if g_throw(self.set_arg)
          (try_count :+ 1,
           g_try(self.set_arg,vlist,bag_type,v,loop))
       else  statement(self.set_arg,bag_type,vlist,loop),
       printf("~A = ~ICreateList(~I,~A.Length())~I~I", v,    // v is the list that is built
              cast_prefix(list,s),
              (if known?(of,self) g_expression(c_code(of(self),type),type) 
               else printf("ToType(CEMPTY.Id())")),
              vlist,
              cast_post(list,s),             // we know we have a list, may cast to put in v
              breakline()),
       printf("for CLcount := 0; CLcount < ~A.~A; CLcount++~I", vlist,       // works for sets and lists :)
              (if (bag_type = set) "Count" else "Length()"), new_block()),
       printf("~A = ~I~A.At(CLcount)~I~I", v2, cast_prefix(any,v2_range) ,vlist, 
                        cast_post(any,v2_range),
                        breakline()),
       if g_throw(self.arg)
          (try_count :+ 1,
           g_try(self.arg,vlocal,any,v,tuple(v,s)))
       else statement(self.arg, any, vlocal,tuple(v,s)),
       printf("~I~A~I.PutAt(CLcount,~A)~I", cast_prefix(s,list), v, cast_post(s,list),vlocal,breakline()),
       close_block(),
       close_try(try_count),
       close_block("Iteration")) ]       


// --------------- WHILE   ------------------------------------------

// if it is possible the logical compiler is used to produce a better code
// self.other = true => until(....) was used 
// error is more tricky => we produce a chain with 3 more blocks
[g_statement(self:While,s:class,v:string,err:boolean,loop:any) : void
 -> let f? := (g_clean(self.test) & not(self.other)),            // simple while pattern 
        try? := g_throw(self.test),                              // error catch pattern for test
        try2? := g_throw(self.arg),                              // error catch pattern for arg
        v2 := check_var("v_while") in
       (if not(f?) var_declaration(v2,boolean,1),
        if (s != void) printf("~I= ~ICFALSE~I~I", c_princ(v), 
                                cast_prefix(boolean,s),
                                cast_post(boolean,s), 
                                breakline()),
        if f?  printf("for ~I ",b_expression(PRODUCER,self.iClaire/test, true))      // while self.test = true ...
        else 
             (if try? g_try(self.iClaire/test,v2,boolean,v,false)
              else statement((if self.iClaire/other false else self.iClaire/test), boolean, v2,false),
              breakline(),
              printf("for ~A ~I CTRUE ", v2,
                      (if self.other princ("!=") else princ("==")))),      // v3.00.05
        new_block("while"),
        if try2? g_try_void(self.arg,v,tuple(v,s))           //    Aspecial form of g_try (no need)
        else statement(self.arg,void,v,tuple(v,s)),               // notice the loop parameter is set
        if try? g_try(self.iClaire/test,v2,boolean,v,tuple(v,s))  // will generate a break if an error
        else if not(f?) statement(self.iClaire/test,boolean,v2,false),
        close_block("while"),
        if try? (close_try(2)),
        if try2? (close_try(1))) ]
   

//------------- compiling a return -------------------------------------
// a return inside a loop is compiled with a break, the go variable is provided
// in the loop argument
[g_statement(self:Return,s:class,v:string,err:boolean,loop:any) : void
 -> case loop 
      (tuple let vreturn := loop[1] as string, sreturn := loop[2] as class in
          ( if PRODUCER.debug? printf(" /*v = ~A, s =~S*/~I",vreturn, sreturn,breakline()),
            statement(self.arg, sreturn, vreturn, false))),      
    princ("break"),
    breakline() ]
         

//**********************************************************************
//*          Part 4: CLAIRE-specific structures                        *
//**********************************************************************

// ------------- Messages and optimized instructions ------------------------

// this is one example on how to unfold: a Call
// note that if the error is returned it should be passed away
// we also add inline_stat in v4 for special cases
[g_statement(self:Call,s:class,v:string,err:boolean,loop:any) : void
 -> if g_clean(self.args) inline_stat(self,s,v)        // ASSUMES err = false or s = EID
    else let l := args(self),         // list of arguments in the call
        ld := unfold_args(l) in  // produces the list of (v,expr)
      (if (ld = nil)                // we are here if the call is simple but an error is possible
        error_wrap(self,s,v)
       else unfold_use(ld, Call(selector(self), list{unfold_arg(l,ld,z) | z in l}), s, v, err, loop)) ]

// this is our special inling that requires an assignment (not allowed as an expression in go)
[inline_stat(self:Call,s:class,v:string)
  -> if (self.selector = object! & s != EID)           // object!(...) is our instantiation macro
          let a1 := self.args[1], a2 := self.args[2] in     // a class
             (if (a2 = property & (value(a1) % property))
                 printf("~I = ~I~I",symbol_ident(a1),
                           declare(PRODUCER,value(a1)), 
                           breakline())
              else  (printf("~I = ~Inew(~I).IsNamed(~I,MakeSymbol(~S,~I))~I~I",symbol_ident(a1),
                                  object_prefix(any,a2),
                                  go_class(a2),                  // v3.3.14
                                  class_ident(a2),
                                  string!(a1),
                                  g_expression(module!(a1),module),
                                  object_post(any,a2),
                                  breakline())),
              if (s != void)
                  printf("~I~I = ~I~I", breakline(),c_princ(v),symbol_ident(a1),breakline()))
      else if (s = EID) printf("~I = ~I~I", c_princ(v),g_expression(self,s),breakline())
      else error("design error : inline_stat for ~S", self) ]
    

// A call method is now simpler with unfolding ! very similar structure
[g_statement(self:Call_method,s:class,v:string,err:boolean,loop:any) : void
 -> let l := args(self), ld := unfold_args(l) in
      (if (ld = nil)                // we are here if the call is simple but an error is possible
        error_wrap(self,s,v)
       else  unfold_use(ld, Call_method(arg = self.arg,
                                  args = list{unfold_arg(l,ld,z) | z in l}),  s, v, err, loop)) ]

[g_statement(self:Call_method1,s:class,v:string,err:boolean,loop:any) : void
 -> let l := args(self), ld := unfold_args(l) in
       (//[5] g_stat(~S): ld = ~S // self,ld,
        if (ld = nil)                // we are here if the call is simple but an error is possible
            error_wrap(self,s,v)
        else unfold_use(ld, Call_method1(arg = self.arg,
                                  args = list{unfold_arg(l,ld,z) | z in l}),  s, v, err, loop))  ]

[g_statement(self:Call_method2,s:class,v:string,err:boolean,loop:any) : void
 -> let l := args(self), ld := unfold_args(l) in
       (if (ld = nil)                // we are here if the call is simple but an error is possible
         error_wrap(self,s,v)
       else unfold_use(ld, Call_method2(arg = self.arg,
                                  args = list{unfold_arg(l,ld,z) | z in l}),  s, v, err, loop)) ]


// in Claire 4, Super is handled as a Call
[g_statement(self:Super,s:class,v:string,err:boolean,loop:any) : void
 -> let l := args(self),         // list of arguments in the call
        ld := unfold_args(l) in  // produces the list of (v,expr)
      (if (ld = nil)                // we are here if the call is simple but an error is possible
         error_wrap(self,s,v)
       else unfold_use(ld, Super(selector = selector(self), cast_to = cast_to(self), 
                           args = list{unfold_arg(l,ld,z) | z in l}), 
                 s, v, err, loop)) ]


// trivial 
[g_statement(self:Cast,s:class,v:string,err:boolean,loop:any) : void
 -> statement(self.arg, s, v, loop)]


//-------------- compiling a handle -------------------------------------
// In most cases, s = EID (err = true) and v is an EID variable => reuse v
// in some cases (s != EID => test = any) .. we need a special variable (v2)
// we see if the catch applied (bool : e % S) 
// in CLAIRE4, we know that self.test is a class
[g_statement(self:Handle,s:class,v:string,err:boolean,loop:any) : void
 -> let v2 := (if (s = EID) v else v /+ "_H") in 
        (new_block("handle"),
         if (s != EID) var_declaration(v2,EID,1), 
         printf("h_index := ClEnv.Index~I",breakline()),
         printf("h_base := ClEnv.Base~I",breakline()),
         statement(self.arg,EID,v2,false),    // no exit from loop inside  the try part
         if (self.test = any)   // catch any errors
            printf("if ErrorIn(~A)",v2)
          else printf("if ErrorIn(~A) && ~I.Contains(ANY(~A)) == CTRUE ",v2, 
                         g_expression(self.test,type),v2),
         new_block(),
         if PRODUCER.debug? printf("/* s=~S */",s),
         printf("ClEnv.Index = h_index~I",breakline()),
         printf("ClEnv.Base = h_base~I",breakline()),
         statement(self.other,s,v,loop),    // here we may exit from the loop
         if (s = EID | s = void) close_block()   // we do not need v
         else printf("} else {~I~I = ~I~I~I", breakline(),
                              c_princ(v),from_eid(PRODUCER,v2,s),
                              breakline(),close_block()),
         close_block("handle"))]
     
// to_CL and to_C are presently ignored in CLAIRE 4
// [g_statement(self:Generate/to_CL,s:class,v:string,err:boolean,loop:any) : void
// -> g_statement(self.arg, s, v, err, loop) ]

// [g_statement(self:Generate/to_C,s:class,v:string,err:boolean,loop:any) : void
// -> g_statement(self.arg, s, v, err, loop) ]

// same for a cast
// v3.2.06: the case where self.arg is of type any is painful => it is forbidden in osystem.cl
[g_statement(self:Generate/C_cast,s:class,v:string,err:boolean,loop:any) : void
 -> g_statement(self.arg, s, v, err, loop) ]


//------------- compiling slot read/write -------------------------------

// new in CLAIRE 4 : there are two kinds => err (EID required) or not (self.arg is just too complex)
// we will follow a pattern similar to unfold => create the let then call g_statement on it

// reads a slot.
// there are two reasons for requiring a statement : complex arg or possible error when reading ! hence we check before using unfold ...
[g_statement(self:Call_slot,s:class,v:string,err:boolean,loop:any) : void
  -> if g_clean(self.arg) printf("~I = ~I~I", c_princ(v),  g_expression(self, s), breakline())
     else let varg := build_Variable("v_slot", c_type(self.arg)),
         unfold := Let( var = varg, value = self.arg, arg = Call_slot(selector = self.selector,  arg = varg)) in
       g_statement(unfold,s,v,err,loop)]


// reads an table.
// there are two reasons for requiring a statement : complex arg or possible error when reading !
[g_statement(self:Call_table,s:class,v:string,err:boolean,loop:any) : void
 -> if g_clean(self.arg) printf("~I = ~I~I", c_princ(v),  g_expression(self, s), breakline())
    else let varg := build_Variable("v_table", c_type(self.arg)),
        unfold := Let( var = varg, value = self.arg, arg = Call_table(selector = self.selector,  arg = varg)) in
       g_statement(unfold,s,v,err,loop)]
       
// reads an array.
[g_statement(self:Call_array,s:class,v:string,err:boolean,loop:any) : void
 -> let varg1 :=  build_Variable("va_arg1", array),
        varg2 :=  build_Variable("va_arg2", integer),
        unfold := Let( var = varg1, value = self.selector, 
                       arg = Let(var = varg2, value = self.arg,
                                 arg = Call_array(selector = varg1, arg = varg2, test = self.test))) in
       g_statement(unfold,s,v,err,loop)]


// places a value in a slot with similar conventions ------------------------------------------------------------------------
// Update = [R(x) := y] where R(x) is a Call_slot, a call_array or a call_table 
// THIS USE OF self.arg IS MEGA UGLY AND SHOULD BE SIMPLIFIED IN THE OPTIMIZER LATER ON ... THERE SHOULD AT LEAST EXIST SOME COMMENTS !
// self.arg is a meta parameter /  it is a property (add or put ...) unless a demon if_write is used 
// self.value is Y and self.var is R(x)  => look in goexp.cl 
[g_statement(self:Update,s:class,v:string,err:boolean,loop:any) : void
 -> let X := self.var, p:any := self.selector,
        sr := (case X (Call_slot (if (self.arg = add) member(X.selector.range)
                                  else X.selector.range),
                       Call_array (if (X.test Core/<=t float) float else any),
                       any any U (if (self.arg = add) member(p.range) else p.range))) in
    (//[5] g_stat@Update X.selector = ~S => range = ~S // X.selector, sr,
     if (not(err) & g_func(self.var.arg) & g_func(self.value) & s = void)  
         update_statement(self,class!(sr))
     else if (g_clean(self.var.arg) & g_clean(self.value) & update_write?(self))
        (printf("~A = ",v),
         update_statement(self,class!(sr)))
     else let try_count := 0,
          varg1 :=  build_Variable("va_arg1", (case X
                     (Call_slot domain!(X.selector),
                      Call_array integer,
                      any any U p.domain))),              // Call_table,
           varg2 :=  build_Variable("va_arg2", sr),
           %call := (let xx := copy(X) in (put(arg, xx, varg1), xx)),          // simpler version of X
           %unfold := Update(selector = self.selector,  value = varg2, arg = self.arg, var = %call) in
    (new_block("update"),
     var_declaration(c_string(PRODUCER,varg1),go_range(varg1),1),
     var_declaration(c_string(PRODUCER,varg2),go_range(varg2),1),
     if g_throw(X.arg)
          (try_count :+ 1,
           g_try(X.arg,"va_arg1",go_range(varg1),v,false))              // checks error, then ...
     else  statement(X.arg,go_range(varg1),"va_arg1",loop),           // produces assignment of v1
     if g_throw(self.value)
          (try_count :+ 1,
           g_try(self.value,"va_arg2",go_range(varg2),v,false))              // checks error, then ...
     else  statement(self.value,go_range(varg2),"va_arg2",loop),           // produces assignment of v1
     if PRODUCER.debug? printf("/* ---------- now we compile update ~S ------- */~I",%unfold,breakline()),
     if update_write?(self) printf("~A = ",v),        // then s = EID
     update_statement(%unfold,class!(sr)),
     if (s != void & not(update_write?(self)))        // when we want the result
        printf("~A = ~I~A~I~I",v,
          cast_prefix(go_range(varg2),s),
          "va_arg2",
          cast_post(go_range(varg2),s),
          breakline()),
     close_try(try_count),
     close_block("update"))) ]
 

// this produce the code for an update assuming that self is error-free and functional
// this method handles
//    if_write demons (that perform the update)  p_write(x:any,y:any)
//    defeasible updates   o.StoreX(n,v,CTRUE)
// if we cannot find n (type too generic) => revert to a generic Update method
// NOTE: Update is used for many things:
//   Update(p:property, arg: ss | c_code(x,any),  var:call_slot, value:y) 
//   Update(p:property, arg: add,  var:call_slot, value:y) multivalued
//   Update(p:table, arg: put | c_code(x,any), var: call_table, value:y)   // only with a list-based table!
//   Update(p:exp<array>, arg:put, var:call_array(p,x), value:y)
[update_statement(self:Update,s:class) : void
 ->  let p:any := self.selector, a := self.arg,
         v := self.value, x := self.var, s2 := class!(c_type(self.var)) in
       (if update_write?(self)     // hopefully s = EID
           (Compile/Tighten!(p),                                      // use the best value for r.domain & r.range
            printf("F_~I_write(~I,~I)~I",
                  c_princ(string!(p.name)),
                  g_expression(x.arg, class!(domain(p))),     // here we use arg !
                  g_expression(v,class!(range(p))),
                  breakline()))
        else if (case p (relation (p.Kernel/store? | a = put_store)))         // defeasible update : Call_table or Call_slot
           (if (x % Call_table)
              printf("F_store_list(ToList(~I.Graph),~I,~I,CTRUE)~I",g_expression(p, table),
                      g_table_index(p as table,x.arg),                  
                      g_expression(v, any),
                      breakline())
            else let s2 := class!(c_type((x as Call_slot).arg)), n := (p @ s2).mClaire/index in
               printf("~I.Store~A(~A,~I,CTRUE)~I",
                     g_expression(x.arg,object),
                     (if (s = integer) "Int" else if (s = float) "Float" else "Obj"),
                     n,
                     g_expression(v, (if (s = integer | s = float) s else any)),
                     breakline()))
        else if (case x (Call_array (sort!(member(c_type(p))) = any), any false))
            (//[5] UNSORTED array p is ~S, index is ~S, type : ~S// p, arg(x), c_type(p),
             printf("~I.PutAt(~I,~I)~I", 
               g_expression(c_code(p,array), list), 
               g_expression(arg(x),integer),
               g_expression(v, any),
               breakline()))
        else if (case x (Call_table not(range(p as table) % {integer,float}), any false))
           (printf("Core.F_put_table(~I,~I,~I)~I", 
               g_expression(p,table), 
               g_expression(arg(x),any),
               g_expression(v, any),
               breakline()))
        // this assumes that x is an addressable expression: Call_slot, sorted Call_array, sorted Call_table
        else (if (x % Call_array) 
                 (s2 := sort!(member(c_type(p))),
                  if (s2 = object) s2 := any)
              else if (x % Call_slot) s2 := class!(range(rootSlot(x.selector))),    // what go expects
              printf("~I = ~I~I", g_expression(x, s2), g_expression(v, s2),breakline()))) ]
          
    
// in the expansion of Defarray, we generate x.graph := make_list(29,unknown) that we need to trap
[need_shortcut(v:any) : boolean
  -> case v
         (// to_CL need_shortcut(v.arg),
          Call_method (selector(v.arg) = make_list),
          any false)]
