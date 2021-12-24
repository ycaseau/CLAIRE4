(printf("Hello CLAIRE3, this is our init.cl file\n"))

// Mac version
*where* :: "/Users/ycaseau/claire/v4.0/go3"                      // where the init file is
*output* :: "/Users/ycaseau/claire/v4.0/go3/src"
*meta* :: "/Users/ycaseau/claire/v4.0/meta"                     // source files on github
*compile* :: "/Users/ycaseau/Dropbox/src/clairev4.0/src/compile"   // source files on dropbox
*bsrc* :: "/Users/ycaseau/claire/v4.0/go1/bsrc"
*tsrc* :: "/Users/ycaseau/claire/v4.0/go1/test"

// these are the global variables expected by the compiler
RELEASE:float :: 0.03    // August 4th, 2021 

// ***************************************************************************
// *    Part 1: Modules & compiler environment                               *
// ***************************************************************************

// meta files are now the "official" github directory
(for m in {Core,Language,Reader} source(m) := *meta*)

// where we want to generate the go code
(when c := get_value("compiler") in 
   (c.safety := 5,
    source(c) := *output*))

// debug 
[foo(n:integer) 
  -> if (n < 1) bar(n, 1, 0) else foo(n - 1) ] 

[bar(n:integer, l:listargs) : integer
  -> printf("l = ~S\n",l),
     12 / n ]

 // extensions
 begin(Language)
[self_print(x:pair) -> princ("pair")]

[self_eval(self:pair) : any
   -> pair(first = eval(self.first), second = eval(self.second))]

claire/Map <: Construct(domain:type,of:type)

// map is the most famous function on a lambda
[claire/map(self:lambda,%l:bag) : any
  -> case %l (set {funcall(self,x) | x in %l},
              any list{funcall(self,x) | x in (%l as list)})]

 // create a map from a list of pairs
self_eval(self:Map) : map_set
  -> let m := map!(self.domain,self.of) in
       (for x in self.args
          (case x (pair put(m,x.first,x.second),
                   any error("~S is not a pair, cannot be inserted in map ~S",x,m))),
        m)  


[self_print(self:map_set) : void
  -> printf("map<~S,~S>", domain(self), range(self)) ]               
 
 end(Language)  

 begin(Reader)
 
 // extended in CLAIRE4: reads the x[y] patterns
[readbracket(r:meta_reader,x:any) : any 
  -> let l := nextseq(cnext(r), #/]) in
      (if (x % class & x != type & l) extract_class_call(x,l)
       else if (l % pair | l % Vardef)        // slice CLAIRE4 syntax x[i:j]
         let i := (case l (Vardef l , pair l.first, any unknown)),
             j :=  (case l (Vardef range(l), pair l.second, any unknown)) in
           Call(slice,list(x,i,j)) 
       else Call!(nth, x cons l)) ]
                       

// new in CLAIRE4: reads map<t1,t2>(pairs*)
[readmap(r:meta_reader) : Map
 -> let l1 := nextseq(cnext(r), #/>) in 
     (//[0] readmap l=~S, char = ~<s // l1, firstc(r),
      let l2 := nextseq(cnext(r), #/)),
          m := Map(range = extract_type(l1[1]), of = l1[2]) in
        (for x in l2
           (case x (pair m.args :add x)),
          m))
 ]
        
 [revVar(x:Vardef) : any
   -> let s := mClaire/pname(x), v := value(s) in
        (if (v = unknown) unbound_symbol(name = s)
         else v) ]

// variant in CLAIRE4 when e = ), which can also read a lambda
readList(r:meta_reader, x:any) : any
  -> let y := readblock(r,x,#/)) in
       (if (firstc(r) = #/{) readlambda(r,y) else y)
       

readlambda(r:meta_reader,l:any) : any 
-> let e := nexte(r), lvar := list() in
        (case l
           (Vardef lvar :add l,
            list (for y in l
                    (case y (Vardef lvar :add y,
                             any Serror("[200] ~S is not a variable in lambda (...)->~S",l,e)))),
            any Serror("[200] ~S is not a variable in lambda (...)->~S",l,e)),
         lambda!(lvar,e))        

end(Reader)  


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

// sudoku example : need to put in the doc - good example of rules & branch
bu16 :: module( uses = list(Reader), source = *bsrc*,
                made_of = list("bstub", "sudoku"))


// ***************************************************************************
// *    Part 4: Simple rule examples                                              *
// ***************************************************************************

// these are the old non-regression tests files (refreshed in July 2021)
(printf("Done. \n"))

