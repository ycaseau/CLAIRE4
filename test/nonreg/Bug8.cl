// +------------------------------------------------------------+
// | bug8.cl                                                    |
// | last update: Jan 2000 - Y. Caseau                          |
// +------------------------------------------------------------+


// ---------------------------------------------------------------
// this file contains bugs related with instantiation
// ---------------------------------------------------------------

// old bugs (Yves) with instantiation
[toto <: thing(x:integer, y:integer = 3)]
[comment(o:toto) -> true]

[tata <: object(x:integer, z:integer = 33, ss:string = "default")]

[f() ->
 let a := 1,
     b := 2,
     c := toto(x = 3, y = 5) in
        (check("def value 1",tata(x = 2).z = 33),
         check("def value 2",tata(z = 12).ss = "default"))  ]

(f())

// this has changed since unknown is not allowed for int
Obj <: thing(col:(integer U {unknown}))

[go()
-> let t := Obj() in
   (if (known?(get(col,t))) (check("error!",false),
                             printf("Oui\n")) else
       (record("2nd branch",true),
        printf("Non => OK !\n"))) ]

(go())


// check that nth_put & get are OK - changed in CLAIRE4
// 
[testNTH()
  -> let s := "abcdef" in
       (nth_put(s,3,'a',10),
        check("put",s = "abadef"),
        nth_put(s,6,'b',10),
        check("nth",nth_get(s,6,10) = 'b')) ]

(testNTH())


// a new test - checks co-variant redefinition of slot
AC1 <: thing(x:any = 1,y:integer = 2)

AC2 <: AC1(x:integer = 101,y:integer = 202)

ATestC1 :: AC1()
ATestC2 :: AC2()
BTestC2 :: AC2(x = 3, y = 4)

(check("default", ATestC1.x = 1 & ATestC1.y = 2 & ATestC2.x = 101 & ATestC2.y = 202),
 check("covariant", BTestC2.x = 3))

// a bug from ThB

A <: object(val:any = 12)
B <: A(val:integer = 15)
C <: A(val:float = 15.3)
D <: A(val:(1 .. 20) = 12)
E <: A(val:string = "1")

// removed safe in interpreted version
[f(la:list<A>) : void
  -> for a in la (a.val :+ 1) ]

[checkf(tag:string, la:list, expected:boolean)
  -> try (f(la), check(tag /+ " correct arg: ", expected = true))
     catch any  check(tag /+ " detect bug: ", expected = false) ]


(checkf("A",list<A>(A()),true),
 checkf("B",list<A>(B()),true),
 checkf("C",list<A>(C()),true),
 checkf("D",list<A>(D()),true),
 checkf("E",list<A>(E()),false))


// another one ...
Ax <: object()
Bx <: Ax(val:float)

Cx <: object(a:Ax)
Dx <: Cx(a:Bx)                     //narrows type of slot 'a'

[f(d:Dx) : float -> d.a.val]

XX:Dx :: Dx()

(testOK())

