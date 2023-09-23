// +------------------------------------------------------------+
// | bug4.cl                                                    |
// | last update: Jan 2000 - Y. Caseau                          |
// +------------------------------------------------------------+


// ---------------------------------------------------------------
// this file contains bugs related with iteration
// ---------------------------------------------------------------

// union
[count(s:any) : integer -> let n := 0 in (for x in s n :+ 1, n)]

(check("iterate union", 21 = count((1 .. 10) U (20 .. 30)) ))

// Tibor
*% :: operation(precedence = precedence(/))
[*%(x:integer,y:integer) : integer -> (x * y) / 100]


*profile*:list<integer> ::
    list<integer>(0,0, 0,0, 0,0, 0,0, 0,0, 0,0, 0,0, 0,0, 40,40, 50,60, 70,80, 90,80,
                  70,60, 50,40, 50,40, 60,50, 60,70, 80,60, 50,40, 30,20, 20,20,
                  0,0, 0,0, 0,0)

HL:list :: list(0,20,50,30)
HOURS :: (0 .. 47)

XX:list :: nil

XY:list :: list{ list{ (if (i <= 16 & i >= 42) 0
                        else ((13 *% (*profile*[i + 1] - random(20))) *% HL[k])) |
                      i in HOURS }  |
                k in (1 .. length(HL))} 


[m3()
  -> XX := list{ list{ (if (i <= 16 & i >= 42) 0
                        else ((13 *% (*profile*[i + 1] - random(20))) *% HL[k])) |
                      i in HOURS }  |
                k in (1 .. length(HL))} ]
(m3())

// from Tibor
[test3() 
  -> exists(e in {1,3,4} | e > 10) ]

(assert(test3() = false))

// an interesting bug because the update in fufu is optimized twice which may break if
// substitution are careless ...
Screen <: object()
Request <: object(to:Screen,
                  value:integer = 0)
cClass <: thing(score:integer, requests:set<Request>)
store(score)

[bar(x:integer,c:cClass) : integer 
 -> x]

valueMatch(s1:Screen,s2:Screen) : integer
 -> 12

valueMatch(r:Request,s:Screen) : integer 
  => valueMatch(r.to, s)

bestmatch(r:Request) : Screen -> r.to 

sum(s:any) : integer => (let d := 0 in (for x in s d :+ (x as integer), d))

[fufu(c:cClass)  : void
  ->  c.score := sum(list{ (r.value *% valueMatch(r,bestmatch(r))) | r in c.requests}) 
] 


// a test from ThB
[flatten(L:list[list]) : list
=> let resList := list<any>() in
(for l in L 
    for y in l 
      resList :add y, 
resList)]

// this is a poor list -> should not crash, should be caught
(try list<list>(1,2) 
 catch any (check("OK1",true)))


(try flatten(list<list>(1,2)) 
 catch any printf("OK"))


// test a pattern !  => requires compiler
/*
[Iterate(x:flatten[tuple(list[list])],itv:Variable,e:any)
=> let L:list<list> := (eval(nth(args(x),1)) as list<list>) in
(for l in L
  for i in l
     let itv := i in e)] */

[fooThB()
  -> let v := 0, l := list<list>(list(1,2,3),list(1,2,3)) in
         (for x in flatten(l) (v :+ x, printf("-> ~S\n",x)) ,
          check("fooThB",v = 12))  ]

(fooThB())

// new bug from SGSS
[createLists() 
  -> let l1 := list(1,2,3,4,5),
         l2 := list<float>{ 0.0 | y in l1} in
      check("sgss", length(l2) = 5) ]

(testOK())
