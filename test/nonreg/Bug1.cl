// +------------------------------------------------------------+
// | bug1.cl                                                    |
// | last update: Jan 2000 - Y. Caseau                          |
// +------------------------------------------------------------+


// ---------------------------------------------------------------
// this file contains simple parsing bugs + list/set image/select use cases
// ---------------------------------------------------------------

// this caused a parsing bug
// (in v2.5 [ and ] around class defs are no longer supported)
GUI2 <: thing(x:integer, // 1
              y:integer,  // 2
			  w:integer)  
                         // 4


GUI <: thing(x:integer,  // 1
              y:integer,  // 2
			  w:integer)  // 3
                         // 4

[screen <: object(day:integer, time:integer) ]    // this is a comment

// do not reindent !
[s(d:integer,t:integer) : any
  -> 
    for x in screen
       (if (day(x) = d & time(x) = t) return(x)) ]


// the reader is lazy
C <: thing(st:list)
c :: C(st = list(0,0,0,0))
(c.st[3])

// should complain cannot be compiled !
// this is a bug about Vardef
(#if not(CompiledTest)
(try (hello:world = class,
      check("aha",false))
 catch any record("error caught",true))
 else nil)

 // float sum
 sum(l:list<float>) : float 
  -> let d := 0.0 in (for x in l d :+ x, d)

 // this are 8 case
 case1() -> check("case1", {(x + 1) | x in list(1,2,3)} = {2,3,4})
 case2() -> check("case2", set<integer>{(x * x) | x in (1 .. 3)} = set<integer>(9,4,1))
 case3() -> check("case3", list{(x + 1) | x in (3 .. 7)} = list(4,5,6,7,8))
 case4() -> check("case4", sum(list<float>{(x + 1.0) | x in {3.4,3.6}}) = 9.0)      // order may vary 
 case5() -> check("case5", {x in (1 .. 10) | x > 6} = {7,8,9,10})
 case6() -> check("case6", set<integer>{x in list(1,2,3,4) | x > 0} = set<integer>(1,2,3,4))
 case7() -> check("case7", {x in class | x.name = class.name} = {class})
 case8() -> check("case8", set<string>{x in list("Peter","Paul","Mary") | x[1] = 'P'} = set<string>("Peter","Paul"))
 case9() -> check("case0", list{x in (1 .. 3) | true} = list(1,2,3))

foo4() : float -> sum(list<float>{(x + 1.0) | x in {3.4,3.2}})
 
(//[0] run 9 cases ... //,
  case1(), case2(), case3(), case4(), case5(), case6(), case7(), case8(), case9())

(testOK())


