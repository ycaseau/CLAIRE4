// +------------------------------------------------------------+
// | bug15.cl                                                   |
// | last update: Dec 2001 - Y. Caseau                          |
// +------------------------------------------------------------+

// ---------------------------------------------------------------
// this file contains the compilation bugs & tests that may only
// be run under compiled mode
// ---------------------------------------------------------------

IntVar <: object()
EltTerm <: object()
EltVTerm <: object()

[nth(l:list[integer], i:IntVar) : EltTerm -> EltTerm()]
[nth(l:list[IntVar], i:IntVar) : EltVTerm -> EltVTerm()]

[f() : void 
-> let dL := list<IntVar>(IntVar(),IntVar(),IntVar()),
       d := IntVar(),
       f := IntVar(),
       et := nth(dL,f) in
    printf("hello world")]

// tuple allocation bug

// a classical example
seed:integer :: 0
[getPair() : tuple(integer,integer)
   -> seed :+ 1, tuple(12,seed)]

// a consumer
[eatPair() : integer
 -> let (p1,p2) := getPair() in      // direct assign to tuple component pattern
         let c := 12 in (p1 + p2)]

[testPair()
  ->   (for i in (1 .. 1000) eatPair(),
        let l := getPair(), y := copy(l) in
            (for i in (1 .. 1000) eatPair(),
             check("correct copy", l = y))) ]


(testPair())

// from FL
TOTO <: object
TOTO <: object(l1:list<list<integer>>, l2:list<TOTO>)

[gF(x:TOTO) : integer -> 3]
[fF(x:TOTO, i:integer) : integer -> i]

[titi(x:TOTO) : void
-> x.l1 := list<list<integer>>{
list<integer>{fF(y,i) | i in (1 .. gF(y))} | y in x.l2} ]

(titi(TOTO(l2 = list<TOTO>(TOTO()))))

// from ThB
[gTB() : void 
-> let a := tuple(1,2) in 
(case a 
(tuple(integer,integer) printf("YESn"), 
any error("bug")))] 

Ax <: object() 
Bx <: Ax() 
[foo(a:Ax,i:integer,x:float) : void 
-> if (x > 1000.0) error("Corrupted ~S",x) else printf("OK")] 
[foo(a:Bx,i:integer,x:float) : void -> 0.0] 
[gox() : void -> let a := Ax() in foo(a,2,2.0)] 

TA:Ax :: Bx()

(gTB(), gox())

(testOK())

