//+-------------------------------------------------------------+
//| CLAIRE                                                      |
//| jsstat.cl                                                   |
//| Copyright (C) 2023-2024   Yves Caseau. All Rights Reserved  |
//+-------------------------------------------------------------+

// j_statement is implemented as a general method that calls a restriction
//       j_statement(self:any,v:string,loop:any) 
//       produced the Javascript code that puts the results of the expression self in v
//       if v = "niet" the result is ignored (notice that the compiler does not carry s, the type)
//       loop is the context for the expression (used for iteration)

//**********************************************************************
//*  Table of contents:                                                *
//*          Part 1: Unfolding of complex expressions                     *
//*          Part 2: Basic control structures                          *
//*          Part 3: iteration                                         *
//*          Part 4: CLAIRE-specific structures                        *
//**********************************************************************

//*************************************************************************
//*          Part 1: Unfolding of complex expressions                     *
//*************************************************************************

// when local CLAIRE expressions are not JS expression, we need to unfold the global expression 
// into a big Let (unfolding is a common pattern for C++, Java, Go ...)
// HOWEVER, if only works for list of arguments whose evaluation order is not specified ! (because we move some of the evaluations earlier)
// this reintrant compiling (calling g_statement on a expanded Let) works because Let checks if g_expression can be used

// this function is used to unfold complex expressions that should be compiled as
// expressions and not statements. this is almost the same as Go but we call j_func, and we know that
// the result is a non empty list
[j_unfold_args(l:list) : list
  ->  let lbad := list{i in (1 .. length(l)) | not(j_func(l[i])) } in // list of indices
        (//[5] unfold -> ~S : ~S // list{l[x] | x in lbad}, list{stupid_t(l[x]) | x in lbad},
         list{ Let(var = build_Variable(genvar("arg_"), static_type(l[i])),
                   value = l[i]) | i in lbad}   ) ]


// uses the previous list to use the variable instead of the Fold.
// l is the list of arguments, ld is the previously build unfold_args(l)
j_unfold_arg(l:list,ld:list,x:any) : any
 -> let i := 1, j := 0, m := length(l) in
       (if (while (i <= m)
             (if not(j_func(l[i]))
                 (j :+ 1, (if (l[i] = x) break(true)))
              else if (l[i] = x) break(false),
              i :+ 1))
           var(ld[j])
        else l[i])

// creates the Let from the ldef definition and places the statement x in the body
// note that the error handling is done in the Let (with g_statement)
// x is the call form where the variable has been replaced if needed
j_unfold_use(ldef:list,x:any,v:string,loop:any) : void
 ->  let  n := length(ldef), vb := verbose() in
      (verbose() := 0,                  // v3.1.06
       if not(ldef) error("[internal] design bug g_func(~S) should be true",x),
       for i in (1 .. (n - 1)) ldef[i].arg := ldef[i + 1],
       arg(ldef[n] as Let) := x,
       verbose() := vb,                  // not very elegant !
       //[5] unfold use: ~S (tail = ~S)// ldef[1],x,
       j_statement(ldef[1] as Let,v,loop))



//**********************************************************************
//*          Part 2: Basic control structures                          *
//**********************************************************************

// The re-entry definition (called within g_statement, not directly)
// if functional, the best compiling is into an expression
// s is the expected go type (as a class) + void + EID
// v is nil or a string (name of the variable)
// note that only 3 additional parameters are used since err is recomputed
[call_j_statement(self:any,v:string,loop:any) : void
 -> if j_func(self)
       (if (v != "niet")
           printf("~I = ~I~I", c_princ(v),  j_expression(self), breakline())
       else if (self % If) j_statement(self, v, loop)
       else if (self % delimiter)
             error("[201] Loose delimiter in program: ~S", self)
       else (// we do not need the result (marked by niet)
             stat_exp(PRODUCER,self,void)))   // new: v3.0.60: include a few tricks
    else j_statement(self, v, loop) ]


// make a statement from an expression (in C++ we need a ; - with go a breakline is enough)
// we do not want to place the result in a variable (see upper)
// will get simpler once we have a stable compiler without to_C and to_CL
[stat_exp(c:js_producer,self:any,s:class) : void
  -> if (case self (Call (self.selector = Compile/object!)))
        printf("~I = (new ~I(~S)).Is(~I)~I", js_ident(self.args[1]), 
                 js_class(self.args[2]), 
                 string!(self.args[1]), 
                 j_expression(self.args[2]),
                 breakline())
     else if Compile/designated?(self)  breakline()                        // designated should be ignored if unused
     else if ignore_exp?(self) nil // ignore completely, no breakline necessary
     else printf("~I~I", j_expression(self), breakline()) ]

// produce a inverse dual statement
[update_inverse(c:js_producer,r:relation,x:any,y:any) : void
  ->  (if (r = inverse) nil
       else if known?(inverse,r)
          (update_simple(c,r,x,y),
           breakline(),
           update_simple(c,r.inverse,y,x),
           breakline())
       else error("udate ~S is not Diet - use of properties",self)) ]
       
[update_simple(c:js_producer,r:relation,x:any,y:any) : void
   -> if r.multivalued? 
         printf("~I.push(~I)",c_member(c, x, any,r),j_expression(y))
      else printf("~I = ~I",c_member(c, x, any,r),j_expression(y)) ]


// a DO is a simple go block if there are not errors, a chain otherwise
// the chain means multiple nestings when an error occurs since the rest of the DO must not be
// this is why the close_try(count) are called at the end, to close the embedded ifs (ErrorIn(e))
// we use a specific method code_statement with an additional parameter %need which is true by default
[j_statement(self:Do,v:string,loop:any) : void
 -> do_statement(self,v,loop,true)]

// %need is true = the last argument is needed in v
[do_statement(self:Do,v:string,loop:any,%need:boolean) : void
 -> if (length(self.args) = 1) call_j_statement(self.args[1], v, loop)
    else let l := self.args, m := length(l) in 
           (for n in (1 .. m)
               call_j_statement(l[n],(if (n = m) v else "niet"),loop))]
   
// a Let is a local variable declaration 
// Java script has no equivalent so what should exist is 
//   (a) looking for embedded lets with same variable name (not diet today)
//   (b) issue an error, or change the name of the variable
[j_statement(self:Let,v:string,loop:any) : void
  -> if (self.arg = self.var & v = "niet")           // stupid case where             
         call_j_statement(self.value,v,loop)              // self.var is not needed !                          
     else let ns := c_string(PRODUCER,self.var.mClaire/pname) in
      (if (ns[1] = 'C' & ns[2] = '%') self.var.mClaire/pname := gensym(),     // used in Iterate (C% variables are expanded): ocontrol.cl
       let v2 := c_string(PRODUCER,self.var),  x := self.value,
           f := j_func(x) in
              (var_declaration(PRODUCER,v2,0),                    // no trailing " " in v4.0.6
               if f printf(" = ~I", j_expression(x)),          // 
               breakline(),
               if not(f) call_j_statement(x, v2,loop),
               call_j_statement(self.arg, v, loop))) ]                           // then we must close the chain


// makes a bag from a list of  statements.
// The value cannot be ignored: it is considered as an error (a do should have been used)
// there are two patterns depending if self.of is known : MakeEmptyX(t) or MakeEmptyX(any)
[j_statement(self:Construct, v:string, loop:any) : void
  -> if not(v) error("[202] A do should have been used for ~S", self),
     let v2 := "v_bag_arg",
         kind := (case self (List list, Set set, Tuple tuple, any error("CONSTRUCT BUG: ~S", self))) in
       (new_block("Construct"),                                           // the need for a new block is debatable, but accepted 
        if exists(x in self.args | not(j_func(x))) var_declaration(PRODUCER,v2,1),
        printf("~I= ~A~I", c_princ(v), j_construct(kind), breakline()),
        for x in self.args
          let f := j_func(x) in
             (breakline(),
              if not(f) call_j_statement(x, v2,loop),
              printf("~I.~A(~I)", c_princ(v),(if (kind = set) "add" else "push"),
                      (if f j_expression(x) else c_princ(v2)))),
        close_block("Construct")) ]

// this shoud also be a code_producer method (refactoring)
[j_construct(kind:class) : string
   -> (if (kind = set) "new Set()" else "[]") ]

// A if is easy to compile. We check if the logical compiler can be used
// we now assume that the test retuns a boolean !
 [j_statement(self:If,v:string,loop:any) : void
   ->  if j_func(self.test)                                       // easy case 
          (printf("if ~I ", b_expression(PRODUCER,self.test, true)), new_block("If"))
        else let v2 := c_string(PRODUCER,gensym() /+ "I") in       // use intermediary variable
          (var_declaration(v2, boolean, 1),
           call_j_statement(self.test, v2, loop),
           printf("if (~A == CTRUE) ~I",v2,new_block("If"))),
        call_j_statement(self.arg, v, loop),
        if (self.other = nil | (self.other = false &  v = "niet")) close_block("If")
        else if (self.other % If & g_func(self.other.test))
              printf("~I else ~I",finish_block("If"),j_statement(self.other,v,loop))
        else  printf("} else {~I~I~I",
                     breakline(),
                     call_j_statement(self.other,v,loop),
                     close_block("If")) ]


// --------------- logical combinations and/or -------------------------------

// note: we cannot use unfolding because the order of evaluation is important !
// AND is compiled with IF: as soon as an argument is false, the result is false.
[j_statement(self:And,v:string,loop:any) : void
  -> let v2 := check_var("v_and") in
       (new_block("and"),
        var_declaration(PRODUCER,v2,1),
        breakline(),
        for x in self.args
            (call_j_statement(x,v2,loop),
             printf("if (~I == false) {~I~I} else ~I", c_princ(v2),
                         (if (v != "niet") printf("~I = false", c_princ(v)),
                         breakline(),
                         new_block("arg")))),
        if (v != "niet") printf("~I = true", c_princ(v)),                 
        for x in self.args close_block("arg"),
        close_block("and")) ]

// same thing for OR
[j_statement(self:Or,v:string,loop:any) : void
 -> let v2 := check_var("v_or") in
       (new_block("or"),
        var_declaration(PRODUCER,v2,1),
        breakline(),
        for x in self.args
            (call_j_statement(x,v2,loop),
             printf("if (~I == true) {~I~I} else ~I", c_princ(v2),
                         (if (v != "niet") printf("~I = true", c_princ(v)), 
                         breakline(),
                         new_block("or")))),
        if (v != "niet") printf("~I = false", c_princ(v)),
        for x in self.args close_block("org"),
        close_block("or")) ]


// Here this is the simple assignment, with a true variable
// note that the last line (assigning the value to result is only OK if no error)
[j_statement(self:Assign,v:string,loop:any) : void
 -> let %var := self.var, x := self.arg, v2 := c_string(PRODUCER,%var) in
     (call_j_statement(x,v2,loop),
      if (v != "niet") printf("~I = ~I~I",c_princ(v),c_princ(v2),breakline())) ]
   
// This is the global variable assignment - global variables exist in go so this is pretty simple
// note that the tricky part is the store management
// v4.0.4 : if nativeVar, we need to produce the go object, not an any  (any is now replaced by %srange)
[j_statement(self:Gassign,v:string,loop:any) : void
   ->  let %var := self.var, x := self.arg  in
           (if (j_func(x) & v = "niet")    // simple case
               printf("~I = ~I~I", globalVar(PRODUCER,%var), j_expression(x), breakline())   
            else let v2 := genvar("v_gassign") in
              (var_declaration(PRODUCER,v2,1),
               call_j_statement(x,v2,loop),
               printf("~I = ~I~I", globalVar(PRODUCER,%var), c_princ(v2),breakline()),
               if (v != "niet" & v != v2) 
                   printf("~I = ~I~I",c_princ(v),c_princ(v2),breakline()))) ]

//**********************************************************************
//*          Part 3: Iteration                                         *
//**********************************************************************

// generates the iteration code for a "for x in S ..." expression , once
// all optimization based on code substitution have been performed.
// very nice in go, except that we have to handle error
// if g_member(%set) is native (anything but any) we use the native go form
[j_statement(self:For,v:string,loop:any) : void
-> let v2 := c_string(PRODUCER,self.var), 
       v3 := v2 /+ "_support",                                              // v3 : variable that holds the bag
       v4 := v2 /+ "_iter",                                                 // *Any variable for iteration
       %set := self.set_arg,                                                // set expression
       sbag := bag_class(%set) in                                             // list or set 
      (if (v != "niet") printf("~I= false~I", c_princ(v), breakline()),
       iteration_statement(self,%set,sbag,v,v3,v2),
       call_j_statement(self.arg, "niet",loopPush(v)),     // ... set the loop parameter to vreturn 
       close_block("loop")) ]    // foreach statement

// push v into the loop context
[loopPush(v:string) : any
  ->  tuple(v,any)]
  // -> if (v != "niet") tuple(v,any) else nil]

// iteration_statement produces the bulk of the iteration code
[iteration_statement(self:For,%set:any,sbag:class,v:string,v3:string,v4:string) : void
 ->  let direct := j_func(%set) in
       (if not(direct) 
           (var_declaration(PRODUCER,v3,1),                                            // v3 will contain the bag
            call_j_statement(%set, v3, false)),
        printf("for (const ~I of ~I)~I", 
              c_princ(v4), 
              (if direct j_expression(%set) else c_princ(v3)), 
              new_block("loop2"))) ]



[iteration_statement_old(self:For,%set:any,sbag:class,v:string,v3:string,v4:string) : void
 ->  var_declaration(PRODUCER,v3,1),                                            // v3 will contain the bag
     call_j_statement(%set, v3, false),
     if (sbag = list)                                                           // set iteration 
        printf("for (let i_it = 0; i_it < ~I.length; i_it++) ~I~I",  c_princ(v3), 
                     new_block(),
                     printf("~I = ~I[i_it]~I", c_princ(v4), c_princ(v3),breakline()))
     else printf("for (const ~I of ~I)~I", c_princ(v4), c_princ(v3), new_block("loop2")) ]
          
// here the value is expected to be important, otherwise an error is raised.
// THIS IS ONLY APPLIED TO COLLECT(f(x) | s in S) on lists => Image is delt with
// we currently do not use the native form => use At and Put to work on generic lists
[j_statement(self:Iteration,v:string,loop:any) : void
-> if (v = "niet") error("[203] you should have used a FOR ere:~S", self),         // we shall use v to place the new list !
   let v2 := c_string(PRODUCER,self.var),         // iteration variable x with proper type
       vlist := check_var("v_list"),       // where we place S
       vlocal := check_var("v_local") in     // where we place f(x)
      (//[5] COOL : an Iteration for g_statement:~S // self,
       var_declaration(PRODUCER,vlist,1),          // the list that is being built
       var_declaration(PRODUCER,v2,1),              // variable associated to f(x)
       var_declaration(PRODUCER,vlocal,1),         // need a *ClaireAny
       call_j_statement(self.set_arg,vlist,loop),
       printf("~A = new Array(~A.length)~I", v, vlist, breakline()),
       printf("for (let CLcount = 0; CLcount < ~A.length; CLcount++)~I", vlist,  new_block()),
       printf("~A = ~A[CLcount]~I", v2,vlist, breakline()),
       call_j_statement(self.arg, vlocal,loopPush(v)),
       printf("~A[CLcount] = ~A~I", v,vlocal,breakline()),
       close_block()) ]       


// --------------- WHILE   ------------------------------------------

// if it is possible the logical compiler is used to produce a better code
// self.other = true => until(....) was used 
// error is more tricky => we produce a chain with 3 more blocks
[j_statement(self:While,v:string,loop:any) : void
 -> let f? := (j_func(self.test) & not(self.other)),            // simple while pattern 
        v2 := check_var("v_while") in
       (if not(f?) var_declaration(v2,boolean,1),
        if (v != "niet") printf("~I= false~I", c_princ(v), breakline()),
        if f?  printf("while ~I ",b_expression(PRODUCER,self.iClaire/test, true))      // while self.test = true ...
        else 
             (call_j_statement((if self.iClaire/other false else self.iClaire/test), v2,false),
              breakline(),
              printf("while ~A ~I true ", v2, (if self.other princ("!=") else princ("==")))),      // v3.00.05
        new_block("while"),
        call_j_statement(self.arg,"niet",loopPush(v)),               // notice the loop parameter is set
        if not(f?) call_j_statement(self.iClaire/test,v2,false),
        close_block("while")) ]
   

//------------- compiling a return -------------------------------------
// a return inside a loop is compiled with a break, the go variable is provided
// in the loop argument
[j_statement(self:Return,v:string,loop:any) : void
 -> case loop 
      (tuple let vreturn := loop[1] as string in
          ( call_j_statement(self.arg, vreturn, false))),      
    printf("break // loop = ~S",loop),
    breakline() ]
         

//**********************************************************************
//*          Part 4: CLAIRE-specific structures                        *
//**********************************************************************

// ------------- Messages and optimized instructions ------------------------

// dynamic calls are not supported in CLAIRE-TO-JAVACSRIPT
// it would make sense to support the special case of homogeneous signature + class based dispatch
[j_statement(self:Call,v:string,loop:any) : void
 -> let l := args(self), ld := j_unfold_args(l) in
      j_unfold_use(ld, Call(selector = self.selector,
                            args = list{j_unfold_arg(l,ld,z) | z in l}), v, loop) ]

// A call method is now simpler with unfolding ! very similar structucture
[j_statement(self:Call_method,v:string,loop:any) : void
 -> if (self.arg.selector = mClaire/update)
        update_inverse(PRODUCER,self.args[1],self.args[2],self.args[5])
    else let l := args(self), ld := j_unfold_args(l) in
       j_unfold_use(ld, Call_method(arg = self.arg,
                                  args = list{j_unfold_arg(l,ld,z) | z in l}), v, loop) ]

[j_statement(self:Call_method1,v:string,loop:any) : void
 -> let l := args(self), ld := j_unfold_args(l) in
       j_unfold_use(ld, Call_method1(arg = self.arg,
                                  args = list{j_unfold_arg(l,ld,z) | z in l}), v, loop)  ]

[j_statement(self:Call_method2,v:string,loop:any) : void
 -> let l := args(self), ld := j_unfold_args(l) in
       j_unfold_use(ld, Call_method2(arg = self.arg,
                                  args = list{j_unfold_arg(l,ld,z) | z in l}), v, loop) ]


// in Claire 4, Super is handled as a Call
[j_statement(self:Super,v:string,loop:any) : void
 -> let l := args(self),  ld := j_unfold_args(l) in  // produces the list of (v,expr)
      j_unfold_use(ld, Super(selector = selector(self), cast_to = cast_to(self), 
                           args = list{j_unfold_arg(l,ld,z) | z in l}), 
                 v, loop) ]


// trivial 
[j_statement(self:Cast,v:string,loop:any) : void
 -> call_j_statement(self.arg, v, loop)]


//-------------- compiling a handle -------------------------------------
// This is easy with Javascript that supports exception handling
[j_statement(self:Handle,s:class,v:string,err:boolean,loop:any) : void
 -> printf("try ~I~I~I catch (e) ~I",
            new_block("handle"),
            call_j_statement(self.arg,v,false),    // no exit from loop inside  the try part
            close_block("handle"),
            printf("{ if e.belongs(~I) {~I~I}~I else throw e }", 
                    c_princ(self.test), 
                    breakline(), 
                    call_j_statement(self.other,v,false),
                    breakline())) ]

// same for a cast
// v3.2.06: the case where self.arg is of type any is painful => it is forbiden in osystem.cl
[j_statement(self:Generate/C_cast,v:string,loop:any) : void
 -> call_j_statement(self.arg, v, loop) ]


//------------- compiling slot read/write -------------------------------

// reads a slot.
// there are two reasons for requiring a statement : complex arg or possible error when reading ! hence we check before using unfold ...
[j_statement(self:Call_slot,v:string,loop:any) : void
  -> if j_func(self.arg) printf("~I = ~I~I", c_princ(v),  j_expression(self), breakline())
     else let varg := build_Variable("v_slot", c_type(self.arg)),
              unfold := Let( var = varg, value = self.arg, arg = Call_slot(selector = self.selector,  arg = varg)) in
       j_statement(unfold,v,loop)]


// reads an table.
// there are two reasons for requiring a statement : complex arg or possible error when reading !
[j_statement(self:Call_table,s:class,v:string,err:boolean,loop:any) : void
 -> if j_func(self.arg) printf("~I = ~I~I", c_princ(v),  j_expression(self), breakline())
     else let varg := build_Variable("v_table", c_type(self.arg)),
              unfold := Let( var = varg, value = self.arg, arg = Call_slot(selector = self.selector,  arg = varg)) in
       j_statement(unfold,v,loop)]
       
// reads an array.
[j_statement(self:Call_array,v:string,loop:any) : void
 -> let varg1 :=  build_Variable("va_arg1", array),
        varg2 :=  build_Variable("va_arg2", integer),
        unfold := Let( var = varg1, value = self.selector, 
                       arg = Let(var = varg2, value = self.arg,
                       arg = Call_array(selector = varg1, arg = varg2, test = self.test))) in
       j_statement(unfold,v,loop)]


// places a value in a slot with similar conventions ------------------------------------------------------------------------
// Update = [R(x) := y] where R(x) is a Call_slot, a call_array or a call_table 
// Diet version excludes rules or defeasible updates (store)
[j_statement(self:Update,v:string,loop:any) : void
 -> let X := self.var, p:any := self.selector in
    (//[5] g_stat@Update X.selector = ~S  // X.selector,
     if (j_func(self.var.arg) & j_func(self.value) & v = "niet")  
         j_update_statement(self)
     else if (j_func(self.var.arg) & j_func(self.value))
         printf("~A = ~I",v,j_update_statement(self))
     else let varg1 :=  (if j_func(X.arg) unknown 
                         else build_Variable("va_arg1", any)),   
              varg2 :=  build_Variable("va_arg2", any),
              %call := (let xx := copy(X) in (put(arg, xx, varg1), xx)),          // simpler version of X
              %unfold := Update(selector = self.selector,  value = varg2, arg = self.arg, 
                                var = (if known?(varg1) %call else X)) in
    (new_block("update"),
     if known?(varg1) var_declaration(PRODUCER,c_string(PRODUCER,varg1),1),
     var_declaration(PRODUCER,c_string(PRODUCER,varg2),1),
     if known?(varg1) call_j_statement(X.arg,"va_arg1",loop),              // produces assignment of v1
     call_j_statement(self.value,"va_arg2",loop),           // produces assignment of v1
     if (v != "niet") printf("~A = ",v),        
     j_update_statement(%unfold),
     close_block("update"))) ]
 

// this produce the code for an update assuming that self is error-free and functional
// this method handles many things:
//   Update(p:property, arg: ss | c_code(x,any),  var:call_slot, value:y) 
//   Update(p:property, arg: add,  var:call_slot, value:y) multivalued
//   Update(p:table, arg: put | c_code(x,any), var: call_table, value:y)   // only with a list-based table!
//   Update(p:exp<array>, arg:put, var:call_array(p,x), value:y)
[j_update_statement(self:Update) : void
 ->  let p:any := self.selector, a := self.arg,
         v := self.value, x := self.var, s2 := class!(c_type(self.var)) in
       (if (case x (Call_array true, any false))
            (printf("~I[~I] = ~I~I", 
               j_expression(c_code(p,array)), 
               j_expression(arg(x)),
               j_expression(v),
               breakline()))
        else if (case x (Call_table not(range(p as table) % {integer,float}), any false))
           (printf("WHAT-THE-FUCK-TABLE(~I,~I,~I)~I", 
               j_expression(p,table), 
               j_expression(arg(x)),
               j_expression(v),
               breakline()))
        // this assumes that x is an adressable expression: Call_slot, sorted Call_array, sorted Call_table
        else  printf("~I = ~I~I", j_expression(x), j_expression(v),breakline()))]
          
    

