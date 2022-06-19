(printf("Hello CLAIRE4, this is our init.cl file\n"))

// Mac version
*where* :: "/Users/ycaseau/claire/v4.0/go"                      // where the init file is
*output* :: "/Users/ycaseau/claire/v4.0/go/src"
*meta* :: "/Users/ycaseau/Dropbox/src/clairev4.05/src/meta"            // source files on dropbox (v2)
*compile* :: "/Users/ycaseau/Dropbox/src/clairev4.05/src/compile"      // source files on dropbox (v2)
*bsrc* :: "/Users/ycaseau/claire/v4.0/test/nonreg"
*tsrc* :: "/Users/ycaseau/claire/v4.0/test/perf"
*rsrc* :: "/Users/ycaseau/claire/v4.0/test/rules"

// these are the global variables expected by the compiler
RELEASE:float :: 0.06    // version of March 11th, 2022

// new in v4.0.6
// atIndex : print an integer "minus one"
[at_index(x:any) : void
  -> case x (integer princ(x - 1), any (g_expression(x, integer), princ(" - 1"))) ]

// additions
/*
begin(Core)
[princ(s:string,n:integer) : void
  -> let m := length(s) in
       (if (m > n) princ(substring(s,1,n))
        else (princ(s),
              for i in (m + 1 .. n) princ(' '))) ] 
end(Core)

begin(Language)
iClaire/lexical_index(self:any,lvar:list,n:integer,final?:boolean) : any
 -> (if (self % thing | self % unbound_symbol) lexical_change(self, lvar)
     else (case self
            (Variable (if unknown?(index,self)                          // v3.1.12
                          error("[145] the symbol ~A is unbound",  self.mClaire/pname),
                       self),
             Call let s := lexical_change(self.selector, lvar) in
                    (lexical_index(self.args, lvar, n,final?),
                     if (self.selector != s)
                        (put(selector, self, call),
                         put(args, self, s cons self.args))),
             Instruction let %type:class := self.isa in
                           (if (%type % Instruction_with_var.descendants)
                               (put(index, self.var, n),
                                n := n + 1,
                                if (n > *variable_index*)
                                   *variable_index* := n),
                            for s in %type.slots
                              let x := get(s, self) in
                                (if ((x % thing | x % unbound_symbol) &
                                     s.range = any)
                                    put(s, self, lexical_change(x, lvar))
                                 else lexical_index(x, lvar, n, final?)),
                            if (%type = Assign & (self as Assign).var % unbound_symbol & final?)                // CLAIRE4
                                error("[101] ~S is not a variable but a ~S", (self as Assign).var, owner((self as Assign).var))),             // moved from self_eval @ Assign
             list let %n := length(self) in
                   while (%n > 0)
                     (let x := (nth@list(self, %n)) in
                        (if (x % thing | x % unbound_symbol)
                            nth=@list(self, %n, lexical_change(x, lvar))
                         else lexical_index(x, lvar, n, final?)),
                      %n :- 1),
             any nil),
           self))

end(Language)   

begin(Generate)
[g_expected(s:class) : class 
  -> if (s = float | s = integer) s else any ]  

// debug
[totul?(self:class,l:list) : any
 ->  let lp := get_indexed(self),
         n := length(lp) in
       (if (length(l) = n - 1 & 
            forall(i in (2 .. n) | selector(lp[i]) = (l[i - 1] as Call).args[1]) &          // args are passed in the proper order !
            (self.open = default() | self Core/<=t exception) &
            n <= 4 & forall(i in (2 .. n) | srange(lp[i]) % {any,integer}))
         let %c:any := Call((if (length(l) = 0) mClaire/new! else anyObject!),
                        self cons list{ c_code(x.args[2],any) | x in l}),  // v3.00.10
             m := (close @ self) in
           (if (length(l) = 0) %c := c_code(%c),
            if m Call_method1(arg = m, args = list(%c)) else %c)
        else false) ]

end(Generate)
*/

// ***************************************************************************
// *    Part 1: Modules & compiler environment                               *
// ***************************************************************************

// meta files are now the "official" github directory
(for m in {Core,Language,Reader} source(m) := *meta*,
 for m in {Optimize,Generate} source(m) := *compile*)

// where we want to generate the go code
(when c := get_value("compiler") in 
   (c.safety := 5,
    verbose() := 1,
    source(c) := *output*))


// ***************************************************************************
// *    Part 2: Performance test modules                                     *
// ***************************************************************************

// these are the performance test files of 2020
mFib :: module(part_of = claire,
              source = *tsrc*,
              uses = list(Reader,mClaire),
              made_of = list("testFib"))

mList :: module(part_of = claire,
              source = *tsrc*,
              uses = list(Reader,mClaire),
              made_of = list("testList"))

mSet :: module(part_of = claire,
              source = *tsrc*,
              uses = list(Reader,mClaire),
              made_of = list("testSet"))

mDict :: module(part_of = claire,
              source = *tsrc*,
              uses = list(Reader,mClaire),
              made_of = list("testDict"))

mObj :: module(part_of = claire,
              source = *tsrc*,
              uses = list(Reader,mClaire),
              made_of = list("testObj"))

mSend :: module(part_of = claire,
              source = *tsrc*,
              uses = list(Reader,mClaire),
              made_of = list("testSend"))

mCopy :: module(part_of = claire,
              source = *tsrc*,
              uses = list(Reader,mClaire),
              made_of = list("testCopy"))

// ***************************************************************************
// *    Part 3: Bugs for CLAIRE                                              *
// ***************************************************************************


// parsing bugs: things that cannot get read right
bu1 :: module( uses = list(Reader), source = *bsrc*,
               made_of = list("bstub","bug1"))

// array related bugs
bu2 :: module( uses = list(Reader), source = *bsrc*, 
               made_of = list("bstub", "bug2"))

// table related bugs
bu3 :: module( uses = list(Reader), source = *bsrc*,
               made_of = list("bstub", "bug3"))

// iteration of a union (interpreted) and other patterns
bu4 :: module( uses = list(Reader), source = *bsrc*,
               made_of = list("bstub", "bug4"))

// bugs with floats
bu5 :: module( uses = list(Reader), source = *bsrc*,
               made_of = list("bstub", "bug5"))

// bugs with class & method definitions
bu6 :: module( uses = list(Reader), source = *bsrc*,
               made_of = list("bstub", "bug6"))

// bugs with worlds
bu7 :: module( uses = list(Reader), source = *bsrc*,
               made_of = list("bstub", "bug7"))

// bugs with instantiation & primitive types
bu8 :: module( uses = list(Reader), source = *bsrc*,
               made_of = list("bstub", "bug8"))

// reversible cells from CLP, untyped version
bu9 :: module( uses = list(Reader), source = *bsrc*,
               made_of = list("bstub", "bug9"))

// reversible cells from CLP, typed version
bu10 :: module( uses = list(Reader), source = *bsrc*,
               made_of = list("bstub", "bug10"))

// famous examples (stack example, doc examples ...)
bu11 :: module( uses = list(Reader), source = *bsrc*,
                made_of = list("bstub", "bug11"))

// bug with tuples, lists and sets
bu12 :: module( uses = list(Reader), source = *bsrc*,
                made_of = list("bstub","bug12"))

// test file for handling unknown & inverses
bu13 :: module( uses = list(Reader), source = *bsrc*,
                made_of = list("bstub", "bug13"))

// test file for compilation bug (works with claire3)
bu14 :: module( uses = list(Reader), source = *bsrc*,
                made_of = list("bstub", "bug14"))

// another test file for testing compiler
bu15 :: module( uses = list(Reader), source = *bsrc*,
                made_of = list("bstub", "bug15"))

// sudoku example : shown in the tutorial - good example of rules & branch
bu16 :: module( uses = list(Reader), source = *bsrc*,
                made_of = list("bstub", "sudoku"))

// other examples from rge claire manual : need to put in the doc - good example of rules & branch
bu17 :: module( uses = list(Reader), source = *bsrc*,
                made_of = list("bstub", "manual"))



// ***************************************************************************
// *    Part 4: Simple rule examples                                              *
// ***************************************************************************

// dinner
mDinner :: module( uses = list(Reader), source = *rsrc*,
                    made_of = list("dinner"))


// filter
mFilter :: module( uses = list(Reader), source = *rsrc*,
                   made_of = list("filter"))

// monkey
mMonkey :: module( uses = list(Reader), source = *rsrc*,
                   made_of = list("monkey"))

// zebra
mZebra :: module( uses = list(Reader), source = *rsrc*,
                  made_of = list("zebra"))

// airline - WIP (old CLAIRE 2 example)
mAirline :: module( uses = list(Reader), source = *rsrc*,
                  made_of = list("Airline"))

// these are the old non-regression tests files (refreshed in July 2021)
(printf("Done. \n"))

