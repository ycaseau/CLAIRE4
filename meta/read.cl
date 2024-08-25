//+-------------------------------------------------------------+
//| CLAIRE                                                      |
//| reader.cl                                                   |
//| Copyright (C) 1994 - 2021 Yves Caseau. All Rights Reserved  |
//| cf. copyright info in file object.cl: about()               |
//+-------------------------------------------------------------+

// ---------------------------------------------------------------------
// this file contains the reader object and the top-level read functions
// the second part of the reader is in syntax.cl
// ---------------------------------------------------------------------

// **********************************************************************
// *  Content:                                                          *
// *   Part 1: The reader object                                        *
// *   Part 2: reading blocks                                           *
// *   Part 3: reading expressions                                      *
// *   Part 4: miscellaneous                                            *
// **********************************************************************

// **********************************************************************
// *   Part 1: The reader object                                        *
// **********************************************************************

// global definitions
delimiter <: global_variable()

arrow:any :: keyword(name = symbol!("->"))
(put(name(arrow), arrow))

triangle:any :: (keyword(name = symbol!("<:")))
// *arrow*:boolean :: false

// here we define the basic keywords
reserved_keyword <: keyword()

else :: reserved_keyword()
for :: reserved_keyword()
case :: reserved_keyword()
while :: reserved_keyword()
until :: reserved_keyword()
let :: reserved_keyword()
when :: reserved_keyword()
try :: reserved_keyword()
if :: reserved_keyword()
Zif :: reserved_keyword()
branch :: reserved_keyword()

keyword?(x:any) : boolean -> (x % reserved_keyword)

forall :: keyword()
none :: keyword()
None :: keyword()
:= :: keyword()
: :: keyword()
catch :: keyword()
in :: keyword()
as :: keyword()
:: :: keyword()
printf :: keyword()
assert :: keyword()
return :: keyword()
break :: keyword()
trace :: keyword()
exists :: keyword()
some :: keyword()
=> :: keyword()
? :: keyword()
rule :: keyword()
quote :: keyword()
claire/inspect :: property()
claire/known! :: property()


// the meta_class of the reader --------------------------------------
// The key values are placed in indexed so that they can be changed (eof ...).
// The slot *internal* is used to give addresses to lexical variables.
// The function next reads a character and places it in the slot first.
//
meta_reader <: thing(source:string,
                     s_index:integer = 0,
                     fromp:port,
                     nb_line:integer,          // to remove ! replaced by n_line()
                     external:string = "toplevel",
                     index:integer,
                     last_form:any,              // v3.3: replaces LastExpRead
                     maxstack:integer,
                     toplevel:boolean,
                     eof:integer, space:integer, tab:integer, bracket:any,
                     paren:any, comma:any, curly:any,
                     last_arrow:boolean = false,     // v3.3
                     s_properties:set[property] = Id(reader.s_properties))


// **********************************************************************
// *   Part 2: reading blocks                                           *
// **********************************************************************

// these are the two low level functions found in the Kernel - with a direct import pattern :)
// for old historical reasons, the reader code uses integers and the macro pattern #/a
// we could re-write everything using char now that char are natives (rune)
next(r:meta_reader) : void -> externC("r.Fromp.GetNext()",void)
firstc(r:meta_reader) : integer -> externC("r.Fromp.CharInt()",integer)

// when to stop
stop?(n:integer) : any -> (n = #/, | n = #/) | n = #/] | n = #/}) 

AND:any :: &
OR:any :: mClaire/new!(delimiter, symbol!("|", claire))

// read the next unit (definition, block or expression)
//
nextunit(r:meta_reader) : any
 -> let n := skipc(r) in
      (if (n = eof(r)) (next(r), eof)
       else if (n = #/[) let z := nexte(cnext(r)) in nextdefinition(r,z,nexte(r),true)
       else if (n = #/() (if toplevel(r) nexts(r, none)
                          else readList(r, nexte(cnext(r))))
       else if (n = #/`) Quote(arg = nextunit(cnext(r)))
       else if (n = #/;)
          (while (firstc(r) != eof(r) & firstc(r) != 10) next(r),
           if (firstc(r) = eof(r)) eof
           else (n_line() :+ 1, next(r), nextunit(r)))
       else let x := (if toplevel(r) nexts(r, none) else nextexp(r,true)) in
           (if (toplevel(r) & (case x (Assign (var(x) % Vardef))))
               Defobj(ident = mClaire/pname(var(x)), arg = global_variable,
                      args = list(Call(=, list(range, 
                                               extract_type(range(var(x))))),
                                  Call(=, list(value, arg(x)))))
            else if (x % string) x
            else if (case x (Call (x.selector % r.s_properties &               // v3.3
                                   forall(y in x.args | not(y % Vardef)))))
              let z := (x as Call), a := z.args[1] in
                 (if (z.selector = begin & a % unbound_symbol) z.args[1] := string!(extract_symbol(a)),
                  if (z.selector = end & a % module)  z.args[1] := claire,
                  x)
            else if not(toplevel(r) | x % Assert) nextdefinition(r,x,nexte(r),false)
            else x))


// read the next statement & stops at the keyword e or at a delimiter
// the keyword has been read but not the delimiter, so we know which case
// by testing stop?(first(r))
// Note: it actually reads a fragment
//
nexts(r:meta_reader, e:keyword) : any
 -> let n := skipc(r) in
      (if (n = eof(r)) (next(r), eof)
       else if (n = #/[) let z := nexte(cnext(r)) in nextdefinition(r,z,nexte(r),true)
       else if (e = None) nexte(r)
       else let x := nexte(r) in
              (if keyword?(x) nextstruct(r, x, e)
               else if (e = none & (case x (Call (x.selector % r.s_properties)))) x
               else loopexp(r, x, e, false)))

// loops until the right expression is built (ends with e ',', '}' or ')')
// x is the first expression that was read
loopexp(r:meta_reader, x:any, e:keyword, loop:boolean) : any
 -> let c := firstc(r) in                                        // last char read
   (if (toplevel(r) & e = none & findeol(r)) x
    else if (x = ?) Call(inspect, list(nexte(r)))
    else if (skipc(r) = #/:)     // reads :* forms
       let y := nexte(cnext(r)) in
         (if (y = =) loopexp(r, combine(x, :=, nexte(r)), e, true)
     //     else if (toplevel(r) & y = :) nextinst(r, x)
          else if (y = :) nextinst(r, x)    // AHA (v3.0.05)
          else if operation?(y) extended_operator(y,x,loopexp(r, nexte(r), e, false)) // v3.3.32
          else if (x % Call & c = 32)                            // expects a space (CLAIRE4)
              (//[5] Call+: c = ~S // c,
               let w := nexte(r) in
               (if (w = =>) r.last_arrow := true                  // v3.3.00
                else if not(w = arrow | w = :=) 
                  Serror("[149] wrong keyword (~S) after ~S",list(w,y)),
                nextmethod(r,x,y,(w = :=),false,(w = =>))))
          else (//[5] create a pair ~S ~S// x,y,
                pair(first = x, second = y)))
    else let y := nexte(r) in
           (if (y = e | (y = => & e = arrow))
               (if (y != e) r.last_arrow := true,                   // v3.3
                if stop?(firstc(r))
                   Serror("[151] ~S not allowed after ~S\n",
                          list(char!(firstc(r)), e))
                else x)
            else if (y = triangle | y = arrow | y = : | y = :: | y = =>)
                nextdefinition(r,x,y,false)
            else if (y % delimiter & stop?(firstc(r))) x
            else if operation?(y)       // this is where
               (if loop loopexp(r, combine(x, y, nexte(r)), e, true)      // in a loop, apply precedence
                else loopexp(r, combine!(x, y, nexte(r)), e, true))
            else Serror("[152] Separation missing between ~S \nand ~S [~S?]",
                        list(x, y, e))))

// this is the special form for x :op y - new in v3.3.32
extended_operator(p:property,x:any,y:any) : any
  -> (case x (Call let r := (if (x.selector = nth) x.args[2] else x.args[1]),
                       v := Variable(mClaire/pname = gensym()),
                       x2 := (if (x.selector = nth) Call(nth,list(x.args[1],v))
                              else Call(x.selector,list(v))) in
                     (if (r % Call)
                        Let(var = v, value = r, arg = combine(x2, :=, combine(x2,p,y)))
                      else combine(x, :=, combine(x, p, y))),
              any combine(x, :=, combine(x, p, y))))

// **********************************************************************
// *   Part 3: reading expressions                                      *
// **********************************************************************

// reading the next compact expression - comments are ignored but they can
// be attached to the last read expression
nexte(r:meta_reader) : any
 -> let x := nextexp(r,false) in (if (x % Instruction) r.last_form := x, x)     // v3.3

// reading the next compact expression/ same
// str = true <=> a comment is not ignored and returned as a string
// nextexp handles the specific cases based on first char, nexti handles what starts with an identifier
nextexp(r:meta_reader,str:boolean) : any
 -> (let n:integer := skipc(r) in
      (if (n = #/)) paren(r)
       else if (n = #/}) curly(r)
       else if (n = #/]) bracket(r)
 //      else if (n = #/>) angular(r)
       else if (n = #/|) (next(r), OR)
       else if (n = #/,) comma(r)
       else if (n = eof(r)) Serror("[153] eof inside an expression", nil)
       else if (n = #/;)
          (while (firstc(r) != eof(r) & firstc(r) != 10) next(r),
           if (firstc(r) = eof(r)) eof
           else (n_line() :+ 1, next(r), nexte(r)))
       else if (n = #/#) read_escape(r)
       else if (n = #/`) Quote(arg = nexte(cnext(r)))
       else let y:any := unknown,
                x := (if (n = #/") Kernel/read_string(cnext(r).fromp)  ;"
                      else if (n = #/() readList(r,nexte(cnext(r)))
                      else if (n >= #/0 & n <= #/9) Kernel/read_number(r.fromp)         // read an int or a float
                      else if (n = #/{) readset(r, nexte(cnext(r)))                     // read a set
                      else (y := Kernel/read_ident(r.fromp),                            // we read an identifier ...
                            if (y % string) y 
                            else nexti(r, y))) in                                       // Note that nexti(r) is the default case
              (if (y % string)
                  (if extended_comment?(r,y as string)
				              extended_comment!(r,y as string)
                   else if conditional_comment?(r,y as string)
                      conditional_comment!(r,y as string)
                   else if str y              // comment is returned
                   else nexte(r))             // read a comment => ignored
               else (while (firstc(r) = #/[ | firstc(r) = #/. | firstc(r) = #/<)
	              (if (firstc(r) = #/<)
                    (if (x = map) (x := readmap(r))
                     else let y := nexte(cnext(r)) in
                          (if (x % class & firstc(r) = #/>)
                            (cnext(r),
                             x := extract_class_call(x,list(Call(=,list(of,y)))),
                             x := nexti(r,x))
                           else Serror("[154] ~S<~S not allowed",list(x,y))))
                 else if (firstc(r) = #/[) (x := readbracket(r,x))
                 else let y := Kernel/read_ident(cnext(r).fromp),
                                p := make_a_property(y)  in
                         (x := Call+(selector = p, args = list(x)),
                          if (p.reified = true) x := Call( read, list(x)))),
                 x))))

// extended in CLAIRE4: reads the x[y] patterns
[readbracket(r:meta_reader,x:any) : any 
  -> let l := nextseq(cnext(r), #/]) in
      (if (x % class & x != type & l) extract_class_call(x,l)
       else Call!(nth, x cons l)) ]
                       
// new in CLAIRE4: reads map<t1,t2>(pairs*)
[readmap(r:meta_reader) : Map
 -> //[5] enter readmap with char=~S // firstc(r),
    let l1 := nextseq(cnext(r), #/>) in 
     (//[5] readmap l=~S, char = ~<s // l1, firstc(r),
      if (length(l1) != 2) Serror("[XXX] map<~A requires two types",list(l1)),
      let l2 := nextseq(cnext(r), #/)),
          m := Map(domain = extract_type(l1[1]), of = l1[2]) in
        (for x in l2
           (case x 
             (pair m.args :add x,
              Vardef m.args :add pair(first = revVar(x), second = range(x)),
              any Serror("~S in map<~A>(... is not a pair ",list(x,l1)))),
          m))
 ]

// returns to the original form from which the Vardef was created (name -> symbol)
[revVar(x:Vardef) : any
   -> let s := mClaire/pname(x), v := value(s) in
        (if (v = unknown) unbound_symbol(name = s)
         else v) ]

// reads a compact expression that starts with an ident
nexti(r:meta_reader, val:any) : any
 -> (if (firstc(r) = #/()
      (if (val % {exists, forall, some})
          let v :=  extract_variable(nexte(cnext(r))), %a2 := nexte(r),
		      %a3:any := any in
            (if (%a2 = in) (%a3 := nexte(r),
			    if (nexte(r) != OR)
                              Serror("[155] missing | in exists / forall",nil))
             else if (%a2 = comma(r)) cnext(r)
             else Serror("[156] wrong use of exists(~S ~S ...",list(v, %a2)),
             Exists( var = v, set_arg  = %a3,
                     arg = (let %bind := bind!(r,v), x := nexts!(r,#/)) in
                              (unbind!(r,%bind), x)),
                     other = (if (val = forall) true else if (val = exists) false else unknown)))
        else if (val = rule) (cnext(r), val)
        else readcall(r, val, unknown))
    else if (val = list & firstc(r) = #/{)
       let s := readset(r, nexte(cnext(r))) in
         case s
          (Image (put(isa,s,Collect), s),
           Select (put(isa,s,Lselect), s),
           any Serror("[157] ~S cannot follow list{", list(s)))
    else if ((case val (Call (val.selector = nth & val.args[1] = list),
                        any false)) & firstc(r) = #/{)
       let s := readset(r, nexte(cnext(r))),
           x := extract_of_type(val as Call) in
         case s
          (Image (put(isa,s,Collect), put(of,s,x), s),
           Select (put(isa,s,Lselect), put(of,s,x), s),
           any Serror("[157] ~S cannot follow list{", list(s)))
    else if ((case val (Call (val.selector = nth & val.args[1] = set),
                        any false)) & firstc(r) = #/{)
       let s := readset(r, nexte(cnext(r))),
           x := extract_of_type(val as Call) in
         case s
          (Image (put(of,s,x), s),
           Select (put(of,s,x), s),
           any Serror("[157] ~S cannot follow list{", list(s)))
    else if (firstc(r) = #/:) nextvariable(r, val)
    else if (firstc(r) = #/@)
       let %a1 := Kernel/read_ident(cnext(r).fromp) in
         (if not(%a1 % class)
             Serror("[158] wrong type in call ~S@~S", list(val, %a1)),
          if (firstc(r) = #/() readcall(r, val, %a1)
          else Serror("[159] missing ( after ~S@~S", list(val, %a1)))
    else val) 

// we have read the escape character #
//
read_escape(r:meta_reader) : any
 -> (if (firstc(cnext(r)) = #//)
       let val := firstc(cnext(r)) in (next(r), val)
    else if (firstc(r) = #/')   // close'
       make_function(string!(extract_symbol(Kernel/read_ident(cnext(r).fromp))))
    else if ((firstc(r) = #/i) & (firstc(cnext(r)) = #/f))
       (next(r), Zif)
    else Serror("[160] wrong use of special char #",nil))
        

// **********************************************************************
// *   Part 4: miscellaneous                                            *
// **********************************************************************


// in CLAIRE4: we record the use of classes from other modules, to trigger the do_import pragma
// note : this is not enough, hence the pragma may need to be setup manually
// we could add explicit casts ...
nextvariable(r:meta_reader, val:any) : any
 -> (if (val = <) (skipc(r), triangle)
     else let x := nexte(cnext(r)) in    // type x in val:x
              Vardef(mClaire/pname = extract_symbol(val), range = x))

// reads an expression, then the exact keyword e
//
nexts!(r:meta_reader, e:keyword) : any
 -> (let x := nexts(r, e) in
      (if not(stop?(firstc(r))) x
       else Serror("[161] Missing keyword ~S after ~S", list(e, x))))

// reads an expression, then the exact keyword e
//
nexte!(r:meta_reader, e:keyword) : any
 -> (let x := nexte(r) in
      (if (nexte(r) = e) x
       else Serror("[161] Missing keyword ~S after ~S", list(e, x))))

// ... exact separator
nexts!(r:meta_reader, e:integer) : any
 -> (let x := nexts(r, none) in
      (if (firstc(r) = e) (cnext(r), x)
       else Serror("[162] Missing separator ~S after ~S", list(char!(e), x))))

// ... keyword e or separator n. DOES NOT SKIP the last character
//
nexts!(r:meta_reader, e:keyword, n:integer) : any
 -> (let x := nexts(r, e) in
      (if (firstc(r) = n | not(stop?(firstc(r)))) x
       else Serror("[163] wrong separator ~S after ~S", list(char!(firstc(r)), x))))

// checks if s is an extended comment
//
extended_comment?(r:meta_reader,s:string) : boolean
  -> (let n := get(s,']') in
        (if (s[1] = EOS | s[1] != '[' | n = 0) false // v0.01
         else forall(i in (2 .. n) | s[n] != '[')))

// produce the equivalent extended comment
//
extended_comment!(r:meta_reader,s:string) : any
  -> (let i := get(s,']'),
          k := substring(s,"//",true),
          m := length(s),
          cx := firstc(r) in         // int code for the last char
       (// printf(" we have an extended comment [~A,~A]\n",s,cx),
        print_in_string(),
	      while useless_c(integer!(s[m])) m :- 1,
	    	if (s[m] = ',') (cx := #/, , m :- 1),
	    	if (k = 0) k := m + 1,
        if (i = 3 & s[2] = '?')
           printf("assert(~I)", for j in ((i + 2) .. m) princ(s[j]))
        else printf("trace(~I,\"~I\\n\"~I)",
               for j in (2 .. (i - 1)) princ(s[j]),
               for j in ((i + 2) .. (k - 1)) princ(s[j]),
               (if (k + 3 <= m) (princ(","),
	             for j in ((k + 3) .. m) princ(s[j])))),
		let s2 := read(end_of_string()) in
         (n_line() :+ 1,            // the waiting '\n' is lost
          flush(reader.fromp,cx),          // push back the int
          s2)))

// v4.12 checks if s is a conditional comment //(s) expression
// note : checks the balanced parenthesis
conditional_comment?(r:meta_reader,s:string) : boolean
  -> (if (s[1] = EOS | s[1] != '(') false
      else let n := length(s), ct := 1 in
           (for i in (2 .. n) 
              (if (s[i] = '(') ct := ct + 1
               else if (s[i] = ')') ct := ct - 1,
               if (ct = 0) break(true))))

// v 4.12 produce the equivalent conditional comment
// produces (if (s) expression else assert)  // assert as a marker :)
conditional_comment!(r:meta_reader,s:string) : any
  -> (let n := length(s), ct := 1, k := 0, 
          cx := firstc(r) in
       (// printf(" we have a conditional comment [~A,~A]\n",s,cx),
        for i in (2 .. n) 
          (if (s[i] = '(') ct := ct + 1
           else if (s[i] = ')') ct := ct - 1,
           if (ct = 0) (k := i, break(true))),
        print_in_string(),
        while useless_c(integer!(s[n])) n :- 1,
        if (s[n] = ',') (cx := #/, , n :- 1),
        printf("(if (~I) ~I else assert)",
               for j in (2 .. (k - 1)) princ(s[j]),
               for j in ((k + 1) .. n) princ(s[j])),
        let s2 := read(end_of_string()) in
           (n_line() :+ 1,            // the waiting '\n' is lost
            flush(reader.fromp,cx),   // push back the int
            s2)))