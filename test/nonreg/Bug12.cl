// +------------------------------------------------------------+
// | bug12.cl                                                   |
// | last update: Jan 2000 - Y. Caseau                          |
// +------------------------------------------------------------+


// ---------------------------------------------------------------
// this file contains bugs related with tuples, lists and sets
// ---------------------------------------------------------------

;----------------------------------------
Person <: thing(age:integer)

p1:Person := unknown

[go()
 -> p1 := new(Person),
    p1.age := 4,
    record("tuple",foo(p1)),
    let (a,b) := foo(p1) in
       (check("tuple1",a = p1.age),
        record("tuple2",b = p1.age + 10)) ]

[foo(p:Person) : tuple(integer,integer)
 -> tuple(p.age, (p.age + 10))]

(go())

;----------------------------------------


// a bug from tibor: does not work when compiled
[test() : void
 -> let count := 0, ordlist:list[list] := list<list>() in
    (
      for o in (1 .. 2)
      ( ordlist :add list(o,0,1),
        ordlist :add list(o,1,-1) ),
      printf("ordlist = ~S\n",ordlist),
      for u in ordlist
      ( count :+ u[2],
        printf("u= ~S \n",u))
      )]

[aha() 
 -> let count := 0, ordlist:list[list] := list<list>() in
    (
      for o in (1 .. 2)
      ( ordlist :add list(o,0,1),
        ordlist :add list(o,1,-1) ),
      printf("ordlist = ~S\n",ordlist),
      ordlist) ]


(test(), aha())

// a bug from Yves: insert a enumeration when interpreted
//
L:list :: list(1,2,3,2,1)
TV:integer :: 0

[test2() 
 -> for x in (L but 3) TV :+ 1 ]

(test2(), assert(TV = 4))


// this code from Francois causes a strange warning
WEEKDAY :: (1 .. 7)
MAXWEEK :: 24
WEEK :: (1 .. MAXWEEK)

Nweek:integer := 6
allWEEKS:Interval := (1 .. Nweek)

DATE :: tuple(WEEK,WEEKDAY)
FirstMonday :: tuple(1,1)
[nextWD(d:DATE) : DATE
-> let (w,wd) := d in
(if (wd != 7) tuple(w,wd + 1)
else if (w != Nweek) tuple(w + 1, 1)
else FirstMonday)]


// check that boundaries are OK
[checkBound()
  -> let l := list<integer>(1,2,3) in
       (nth+(l,2,12),
        printf("after nth, l=~S\n",l),
        l[4] := 123,
        check("nth+",l[2] = 12 & l[3] = 2 & l[4] = 123)) ]

(checkBound())


// a bug from Francois
TOTO <: object
TOTO <: object(l1:list<list<integer>>, l2:list<TOTO>)

[g(x:TOTO) : integer -> 3]
[f(x:TOTO, i:integer) : integer -> i]

[titi(x:TOTO) : void
-> x.l1 := list<list<integer>>{
      list<integer>{f(y,i) | i in (1 .. g(y))} | y in x.l2} ]

(titi(TOTO(l2 = list<TOTO>(TOTO()))))

// a new bug !

[essai() : void
  ->
  let l1 := list<integer>(1,2,3,4), l := list<integer>(1,2,3,4) in (
      store(l),  // useless but harmless ?
      choice(),
      store(l, 1, 100, true),
      try store(l, 1)
      catch any printf("error caught"),
      printf("before l = ~S \n",l),
      backtrack(),
      printf("after l = ~S \n",l),
      check("list bk", l = l1)) ]

[essai2() : void
  ->
  let l1 := list<integer>(1,2,3,4), l := make_list(100,integer,0) in (
      shrink(l,0),
      l :add 1, l :add 2, l :add 3, l :add 4,
      choice(),
      store(l, 1, 100, true),
      store(l, 5),
      printf("2 before l = ~S \n",l),
      backtrack(),
      printf("2 after l = ~S \n",l),
      check("(2) list bk", l = l1)) ]

(essai(), essai2())

// some set examples
[testSet() 
 -> check("intersection", {1,2,4,7} ^ {8,4,0,2} = {4,2}),
    check("union", {"a","b"} U {"c","d"} = {"a","b","c","d"} ),
    check("size", size({class,property,class,method}) = 3 ),
    check("but", {1,2,3} but 1 = {2,3})
 ]

 (testSet())

 // tests with hybrid sets
 [testSet2()
   -> let a := TOTO(), b := list(1,2,3), c := Call(selector = +, args = list(1,2)),
          s := set(a,b,c,"d",123,12.3), nota := TOTO() in
       (check("object membership", (a % s)),
        check("call membership", (c % s)),
        check("string membership", ("d" % s)),
        check("list membership", (list(1,2,3) % s)),
        check("int membership", (123 % s)),
        check("float membership", (12.3 % s)),
        check("object negative membership", not(nota % s)),
        check("string negative membership", not("e" % s)),
        check("list negative membership", not(list(1,2,4) % s)),
        check("int negative membership", not(124 % s)),
        check("float negative membership", not(12.4 % s))) ]

 (testSet2())       


 (testOK())

