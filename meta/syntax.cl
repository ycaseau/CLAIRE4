//+-------------------------------------------------------------+
//| CLAIRE                                                      |
//| syntax.cl                                                   |
//| Copyright (C) 1994 - 2021 Yves Caseau. All Rights Reserved  |
//| cf. copyright info in file object.cl: about()               |
//+-------------------------------------------------------------+

// --------------------------------------------------------------
// this file contains specialized reading methods
// --------------------------------------------------------------

// **********************************************************************
// *  Content:                                                          *
// *   Part 1: read operation expressions (<exp> <op> <exp>)            *
// *   Part 2: read control structures                                  *
// *   Part 3: read functional calls                                    *
// *   Part 4: read definitions                                         *
// **********************************************************************

// **********************************************************************
// *   Part 1: read operation expressions (<exp> <op> <exp>)            *
// **********************************************************************

// who is an operation?
//
operation?(y:any) : boolean
 -> (y = as | y = := | y % operation | y = OR | y = % | y = add) 

// produce an expression from an operation
// apply precedence rules ((x1 op x2) y  z) -> x1 op (x2 y z)
combine(x:any, y:any, z:any) : any
 -> let p := operation!(x) in
      (if (p & precedence!(y) < precedence!(p))
          combine!(operand!(x, 1), p, combine(operand!(x, 2), y, z))
       else combine!(x, y, z)) 

// produces x op=y z
// replace r(x) :add y with add(r,x,y) for multivalued or defeasible .. also with delete
combine!(x:any, y:any, z:any) : any
 -> (if (y = as) Cast(arg = x, set_arg = extract_type(z))
    else if (y = :=)
       case x
        (global_variable Gassign(var = x, arg = z),
         Call (if (case z (Call (((selector(z) = add) | (selector(z) = delete)) &
                                 (args(z)[1] = x) &
                                 (if (length(args(x)) = 1) x.selector.multivalued?
                                  else if (x.selector = nth & length(x.args) = 2)
                                       (let p := x.args[1] in (case p (relation p.multivalued?, any false)))))))
                   (Call!(selector(z),
                         list(
                           (if (length(args(x)) = 1) selector(x) else args(x)[1]),
                           (if (length(args(x)) = 1) args(x)[1] else args(x)[2]),
                           args(z)[2])))
               else if (selector(x) = nth)
                  Call!(nth=, add(copy(args(x)), z))
               else if (length(args(x)) = 1)
                 let p := make_a_property(selector(x)), y := args(x)[1] in
                   (if (p = read & (case y (Call+ (y.selector.reified = true)))) Call!(write,list(y,z))
                    else Call!(write, list(p,y,z)))
               else Serror("[164] ~S cannot be assigned with :=", list(x))),
         Do let l := args(x), m := length(l),
                v := Variable(mClaire/pname = gensym()) in
              Let*(var = v, value = z,
                   arg = Do(list<any>{ 
                               Assign(var = l[i], arg = Call!(nth,list(v,i))) |
                               i in (1 .. m)})),
         any Assign(var = x, arg = z))
    else if (y = OR)
       case x (Or (add!(args, x, z), x), any Or(args = list(x, z)))
    else if (y = AND)
       case x (And (add!(args, x, z), x), any And(args = list(x, z)))
    else if (y = %) Call!(%,list(x, z))
    else DBregister(Call*(selector = y, args = list(x, z))))     // Call* says that combining is OK

// allows to treats Calls, Assigns, Gassign in an homogeneous way
// return false if the pattern is not (x OP y) and OP otherwise
operation!(x:any) : any
 -> (case x
     (Or OR,
      And AND,
      Assign :=,
      Gassign :=,
      Call let p := selector(x) in
             (if (x % Call* & operation?(p)) p    // Call* is a marker (produced in the loopexp)
              else if (p = nth=) :=
              else if (p = write) :=),
      any false))

// extract the two operands from an expression x such that operation!(x) != false
operand!(x:any, n:integer) : any
 -> (case x
     (Or (if (n = 1) Or(args = rmlast@list(copy(args(x))))
                        else last@list(args(x))),
      And (if (n = 1) And(args = rmlast@list(copy(args(x))))
                          else last@list(args(x))),
      Assign (if (n = 1) var(x) else arg(x)),
      Gassign (if (n = 1) var(x) else arg(x)),
      Call (if (selector(x) = write)
               (if (n = 2) args(x)[3]
                else Call!(args(x)[1],list(args(x)[2])))
            else if (selector(x) = nth=)
               (if (n = 2) last(args(x))
                else Call!(nth,rmlast(copy(args(x)))))
            else args(x)[n]),
      any false))

// precedence
//
precedence!(y:any) : integer
 -> (if (y = as) 0
    else if (y = :=) 100
    else if (y = AND) 1000
    else if (y = OR) 1010
    else precedence(y as operation))

// **********************************************************************
// *   Part 2: read control structures                                  *
// **********************************************************************

nextstruct(r:meta_reader, %first:keyword, e:keyword) : any
 -> (if (%first = let) readlet(r, e)
    else if (%first = when) readwhen(r, e)
    else if (%first = case) readcase(r, e)
    else if (%first = for)
       let %var := extract_variable(nexts!(r, in)),
           %set := nexte(r),
           %bind := bind!(r, %var),
           x := (if (firstc(r) = #/,) next(r), ; to remove later
                 For(var = %var, set_arg = %set, arg = nexts(r, e))) in
         (unbind!(r, %bind), x)
    else if (%first = while) While(test = nexte(r), arg = nexts(r, e), other = false)
    else if (%first = until) While(test = nexte(r), arg = nexts(r, e), other = true)
    else if (%first = try)
       let %a := nexts!(r, catch), %t := nexte(r) in
         (if (%t % class) Handle(test = %t, arg = %a, other = nexts(r, e))
          else Serror("[00] in try/catch, ~S is not a class",list(%t)))
    else %first)

// reads a let expression
//
readlet(r:meta_reader, e:keyword) : any
 -> (let %def := nexts!(r, in, #/,) in
       case %def
        (Assign let v := extract_variable(var(%def)),
                    %bind := bind!(r, v),
                    x := Let(var = v, value = get(arg, %def),
                            arg = (if (firstc(r) = #/,) readlet(cnext(r), e)
                                   else nexts(r, e))) in
                   (unbind!(r, %bind), x),
         Let*  (arg(%def) := readlet*(r,args(arg(%def) as Do),1,e), %def),
         Call[selector = write]
           let v1 := Variable(gensym(), any),
               v2 := Variable(gensym(), any),
               %a := args(%def), %e := nexts(r, e) in
              Let+(var = v1, value = Call+(selector = %a[1], args = list(%a[2])),
                   arg = Do(list<any>(%def,
                                 Let(var = v2, value = %e,
                                     arg = Do(list<any>(Call!(write,list(%a[1],%a[2],v1)), v2)))))),
         any Serror("[165] ~S is illegal after a let", list(%def))))
         


// recursive construction of the tail of a Let*
readlet*(r:meta_reader, l:list, n:integer, e:keyword) : any
 -> (if (n > length(l)) nexts(r, e)
    else let v := extract_variable(var(l[n])),
             %bind := bind!(r, v),
             x := Let(var = v, value = arg(l[n]), arg = readlet*(r,l,n + 1, e)) in
           (unbind!(r, %bind), x))


// reads a when expression
readwhen(r:meta_reader, e:keyword) : any
 -> (let %def := nexts!(r, in, #/,) in
      case %def
        (Assign let v := extract_variable(var(%def)),
                    %bind := bind!(r, v),
                    %a := nexts(r,else),
                    x := When(var = v, value = get(arg, %def), arg = %a,
                              other = (if stop?(firstc(r)) unknown
                                       else nexts(r, e))) in
                   (unbind!(r, %bind), x),
        any Serror("[165] ~S is illegal after a when", list(%def))))
         

// read an if
readif(r:meta_reader, e:integer) : any
 -> (let %a1 := nexte(r),
        %a2 := nexts(r, else) in
      If(test = %a1, arg = %a2,
         other =
           ((if (firstc(r) = #/, | firstc(r) = e) false
             else let x := nexte(r) in
                    (if (x = if) readif(r, e)
                     else if keyword?(x) nextstruct(r, x, none)
                     else loopexp(r, x, none, false))))))

// reads a member_of
//
readcase(r:meta_reader, e:keyword) : any
 -> (let %v := nexte(r) in
      (if (skipc!(r) != #/() Serror("[166] Missing ( after case ~S", list(%v)),
       let %x:Case := (Case(var = %v, args = list())), // v0.01
           %t:any := any in
         (while (firstc(r) != #/))
            (next(r),
             %t := extract_type(nexte(r)),
             put(args, %x, add(add(args(%x), %t), nexts(r, none))),
             if (not(stop?(firstc(r))) & not(stop?(skipc(r))))          ; because of toplevel ....
                Serror("[167] missing ) or , after ~S",list(%x))),
          next(r),
          if (e != none & not(stop?(skipc(r))) & nexte(r) != e)
             Serror("[161] missing ~S after ~S", list(e, %x)),
          %x)))

// if the expression begins with "{"
//
readset(r:meta_reader, %a1:any) : any
 -> (if (%a1 = curly(r)) (next(r), {})
    else (if keyword?(%a1) %a1 := nextstruct(r, %a1, none),
          let %a2 := nexte(r) in
            (if (%a2 = comma(r))
	       Kernel/cast!({dereference(u) | u in (%a1 cons nextseq(cnext(r), #/}))},{})
             else if (%a2 = curly(r))
               (next(r), Kernel/cast!(set(dereference(%a1)),{}))
             else if (%a2 = in)
                let v := extract_variable(%a1) in
                  Select(var = v, set_arg = nexte(r),
                         arg =
                           (let %bind := bind!(r, v),
                                x := ((if (nexte(r) != OR)
                                          Serror("[168] missing | in selection",nil)
                                       else nexts!(r, #/}))) in
                              (unbind!(r, %bind), x)))
             else if (%a2 = OR)
                let v := extract_variable(nexts!(r, in)) in
                  lexical_index(Image(var = v, set_arg = nexts!(r, #/}),
                                      arg = substitution(%a1,v,v)), 
                                list(v), 0, false)                     // v4.0.6 = partial lexical indexing
             else if operation?(%a2)
                readset(r,
                        loopexp(r, combine(%a1, %a2, nexte(r)), none,
                                false))
             else Serror("[169] missing separation between ~S and ~S",
                         list(%a1, %a2)))))

dereference(x:any) : any
  -> (case x (unbound_symbol error("[170] cannot use ~S in a set constant",x),
             Variable error("[170] cannot use a variable (~S) in a set constant",x),
             any eval(x)))
			
			
// reads a sequence of exp. Must end with a e = ) | ] | }
//  <actually returns a list>
nextseq(r:meta_reader, e:integer) : any
 -> (//[5] enter nextseq(~S) first=~S // e, firstc(r),
    if (firstc(r) = e) (next(r), list())  // this allows reading empty lists ()
    else let x := nexts(r, (if (e = #/>) None else none)) in
           (if (firstc(r) = 10 & r.toplevel) skipc(r),         // v3.2.22
            if (firstc(r) = e) (next(r), list(x))
            else if (firstc(r) = #/,) 
              let y := nextseq(cnext(r), e) in
                 (if (case y (list (length(y) = 0)))           // v4.0.6 : once a , is read, an value must follow before )
                     Serror("[171] Read the character ) inside a sequence",list())
                  else (x cons y))    // builds the sequence recursively
            // (x cons nextseq(cnext(r), e))
            else Serror("[171] Read the character ~S inside a sequence",
                        list(char!(firstc(r))))))

// read the next block: a sequence of exp. Must end with a e = ) | ] | }
//
readblock(r:meta_reader, x:any, e:integer) : any
 -> (skipc(r),
    if (x = paren(r)) list()
    else if (firstc(r) = #/,)  
          let y := nexte(cnext(r)) in
            (case y (delimiter Serror("[172] delimiter ~S found too soon after ~S + comma(,)",list(y,x)),
                     any Do!(x,readblock(r,y,e))))               // v4.0.6 check that ,] is trapped
         // Do!(x, readblock(r, nexte(cnext(r)), e))
    else if (firstc(r) = e) (cnext(r), x)
    else if stop?(firstc(r))
       Serror("[172] the sequence ...~S must end with ~A", list(x, char!(e)))
    else if (x = if) readblock(r, readif(r, e), e)
    else if (x = Zif) let %i := (readif(r,e) as If) in
                         readblock(r,(if eval(test(%i)) arg(%i) else other(%i)),e)
    else if (x = else) Serror("[173] Expression starting with else", nil)
    else if keyword?(x) readblock(r, nextstruct(r, x, none), e)
    else let y := loopexp(r, x, none, false) in
          (case y (Call* put(isa,y,Call)),    // this is how () are implemented -> forbids combining
           readblock(r, y, e)))
           
// variant in CLAIRE4 when e = ), which can also read a lambda of for (...){...}
readList(r:meta_reader, x:any) : any
  -> let y := readblock(r,x,#/)) in
       (//[5] DEBUG: readList gets block ~S char=~A // y, firstc(r),
        if (firstc(r) = #/{) readlambda(r,y) else y)    
 
// create the lambda
readlambda(r:meta_reader,l:any) : any 
-> let lbody := nextseq(cnext(r),#/}), lvar := list() in
        (//[5] read a lambda ~S with body ~S // l, lbody,
         case l
           (Vardef lvar :add! l,
            Do (for y in l.args
                  (case y (Vardef lvar :add! y,
                           any lvar :add! Variable(mClaire/pname = extract_symbol(y), range = any)))),
            list (for y in l
                    (case y (Vardef lvar :add! y,
                             any lvar :add! Variable(mClaire/pname = extract_symbol(y), range = any)))),
            any lvar :add! Variable(mClaire/pname = extract_symbol(l), range = any)),
         lambda!(lvar,(if (length(lbody) = 1) lbody[1] else Do(lbody))))                

Do!(x:any, y:any) : any
 -> (case y
     (Do (put(args, y, nth+(args(y), 1, x)), y),
      any Do(args = list<any>(x, y))))

// extract the type from a list<X> expression
extract_of_type(x:Call) : type
  -> let l := x.args in
       (if (length(l) > 2)
           let y := l[3] in
             (case y (List let z := y.args[1] in
                             (case z (Set extract_type(z.args[1] as type),
                                      any any)),
                      any any))
        else any)

// **********************************************************************
// *   Part 3: read functional calls                                    *
// **********************************************************************

// store the line number in debug mode
// in v4.0 we will not do this for JITO calls :)
DBline[c:Call] : integer := 0

// this is a cool trick when operating in debug mode: we store the last evaluated
// call so we can tell very simply which last call triggered the error
//
DBregister(c:Call) : Call
 -> (if (system.Core/debug! >= 0) (Language/LastCall := c,DBline[c] := n_line()),
     if (c.selector = store & length(c.args) = 1)        // otherwise it is ambiguous !
        let l := c.args in
          (if (l[1] % global_variable) l[1] := make_string(name(l[1]))),
     c)

Call!(p:property,l:list) : Call -> DBregister(Call(p,l))

// if the expression is a call -------------------------------------------
// x is the first token that we have read x(...)
// there are many special case (x is not a propery) then the regular case
// t is a type when x was read as (p@t)
readcall(r:meta_reader, x:any, t:any) : any
 -> (let l := nextseq(cnext(r), #/)) in              // read the arg list
      (if (x = printf) Printf(args = l)
       else if (x = error) Error(args = l)
       else if (x = assert)
           Assert(args = l, index = n_line() , external = external(r))
       else if (x = trace) Trace(args = l)
       else if (x = branch) Branch(args = l)
       else if (x = quote) Quote(arg = (if l l[1]))
       else if (x = tuple) Tuple(args = l)
       else if (x = list) List(args = l)     // unmutable list => of = unknown
       else if (case x (Call (x.args[1] = list)))     // ------ the 3 parameterized constructor
           List(of = extract_of_type(x as Call), args = l)
       else if (case x (Call (x.args[1] = array)))              // new in v3.2.16
           Array(of = extract_of_type(x as Call), args = l)
       else if (case x (Call (x.args[1] = set)))
           Set(of = extract_of_type(x as Call), args = l)
       else if (x % class & x inherit? Macro) let o := mClaire/new!(x) in (put(args,o,l), o)
       else if (x = set) Set(args = l)
       else if (x = return | x = break) Return(arg = (if l l[1] else true))
       else if (x % class)
          (if not(forall( y in l |
                          case y (Call (if (selector(y) = =)
                                       (args(y)[1] := make_a_property(args(y)[1]), true)))))
            let l2 := params(x as class),
                n :=  (case l2 (list length(l2), any 0)) in
             (if (length(l) = n)
                l := list{Call(selector = =, args = list(l2[i],l[i])) |
                          i in (1 .. n)}
              else Serror("[174] Wrong instantiation list ~S(~S...",list(x,list(l)))),
           Definition(arg = x, args = l))
       else if (x % Variable | (case x (global_variable x.range)))
           Call!(call,cons(x,(if l l else list(system))))
       else let p := make_a_property(x),
                l2 := ((if l l else list(system))) in        // introduce the fake system arg
              (if known?(t) Super(selector = p, cast_to = t, args = l2)
               else Call!(p,l2))))

// **********************************************************************
// *   Part 4: read definitions                                         *
// **********************************************************************

// reads a definition (CLAIRE2 syntax)   - x and y are two expressions that have been read
//
nextdefinition(r:meta_reader,x:any,y:any,old?:boolean) : any
 -> (r.last_arrow := false,
     if (y = triangle) nextDefclass(cnext(r),x,old?)
     else if (y = :)
     let table? := (case x (Call (selector(x) = nth &
                                  args(x)[1] % (unbound_symbol U table)))),
         z := nexte(r), w := nexte(r) in
          (if (if table? (w = :=) else (w = arrow | w = =>)) nil         //v3.3
           else Serror("[149] wrong keyword (~S) after ~S",list(w,z)),
           nextmethod(r,x,z,table?,old?,( w = =>)))
    else if (y = ::)
     case x (Call let ru := nexte(r), z := nexts(r, =>) in
                   Defrule(ident = name(selector(x)), args = args(x),
                           arg = z,
                           body = (if (firstc(r) = #/)) (next(r),nil)
                                   else readblock(r, nexte(r), #/))) ),
             any nextinst(r,x))
    else if (y = arrow | y = =>)
     (r.last_arrow := (y = =>),
      //[3] ---- note: ~S - method's range is assumed to be void // x,
      nextmethod(r,x,void,false,old?,(y = =>)))       // v3.3 must be void
    else if (y = := & x % Vardef)
          Defobj(ident = mClaire/pname(x), arg = global_variable,
                 args = list(Call!(=, list(range, extract_type(range(x)))),     // v3.1.14
                             Call!(=, list(value, nexte(r)))))
     else Do(args = list<any>(x,y)))
    
// read a method - old? = true means that the method definition is between brackets    
nextmethod(r:meta_reader,x:any,y:any,table?:boolean,old?:boolean,inl?:boolean)  : any
  -> (let n := skipc(r),
          z := ( if old? readblock(r, nexte(r), #/])
                 else if (n = #/() (if toplevel(r) nexts(r, none)
                               else readblock(r, nexte(cnext(r)), #/)))
                 else nexte(r)),
          rs := Defmethod(arg = x, set_arg =  y, 
                          body = (if (z = let) readlet(r, None) else z), 
                          inline? = inl?) in
       (if table? put(isa,rs,Defarray), rs))

// reads an instantiation
//
nextinst(r:meta_reader, x:any) : any
 -> (case x
     (Variable (Defobj(ident = mClaire/pname(x), arg = global_variable,
                      args = list(Call(=, list(range, extract_type(range(x)))),  // v3.1.14
                                  Call(=, list(value, nexte(r)))))),
      Call let ru := nexte(r), z := nexts(r, =>) in
                   Defrule(ident = name(selector(x)), args = args(x),
                           arg = z,
                           body = (if (firstc(r) = #/)) (next(r),nil)
                                   else readblock(r, nexte(r), #/))) ),
      any let y := nexte(r) in
            (if (case x (global_variable unknown?(y))) y
			 else if (case y (Definition (arg(y) inherit? thing)))
                 Defobj(ident = extract_symbol(x), arg = arg(y),
                                args = args(y))
             else Defobj(ident = extract_symbol(x), arg = global_variable,
                         args =  list(Call(=,list(range, {})),
                                      Call(=,list(value, y)))))))


// reads a class Definition of the form C(p:t | p:t = v *)
// new in v2.5
nextDefclass(r:meta_reader,x:any,old?:boolean) : Defclass
 -> (skipc(r),
    let c := verify(class, Kernel/read_ident(r.fromp), Defclass),     // superclass
        y:Defclass :=                                    // create the Defclass
          (if (firstc(r) != #/() Defclass(arg = c, args = nil, forward? = true)
           else let l := nextseq(cnext(r), #/)) in
              (for y1 in l,
                 (if not(case y1
                          (Call (selector(y1) = = & args(y1)[1] % Vardef),
                           Vardef true,
                           any false))
                     Serror("[175] Wrong form ~S in ~S(~S)", list(y1, c, l))),
               Defclass(arg = c, args = l, forward? = false))),
        lp := nil,
        idt := (if (case x (Call selector(x) = nth))    // extract ident + list
                 let l := (x as Call).args in           // of parameters
                     (if (l[1] % class) lp := (l[2] as list)
                      else lp := list{make_a_property(y2) | y2 in cdr(l)},
                      extract_symbol(l[1]))
                 else extract_symbol(x)) in
        (if (old? & (skipc(r) != #/])) Serror("[176] Missing ] after ~S ", list(y))
         else if old? next(r),
         y.ident := idt,
         y.params := lp,
         y))


// end of file
