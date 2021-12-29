// +------------------------------------------------------------+
// | bug6.cl                                                    |
// | last update: Jan 2000 - Y. Caseau                          |
// +------------------------------------------------------------+


// ---------------------------------------------------------------
// this file contains bugs related with class and method definition
// ---------------------------------------------------------------

// "bug" found by F. Laburthe -----------------------------------

// test the definition of a class
myclass <: thing(uu:integer)

[private/g(x:integer) : integer -> 1]

(check("1",g(0) = 1))

[claire/g2(x:(0 .. 3)) : integer -> 2]

[g2(x:integer) -> 1]

(check("2", g2(0) = 2))

AnIn :: myclass(uu = 2)

AnOut :: myclass(uu = 1)

// typing ... (compilation bug) --------------------------------------

[C1 <: thing()]

[C2 <: C1(li:list[C1])]

[m1(x:C1) : C1
  -> case x 
      (C2 x.li := (list{m1(y) | y in x.li} as list<C1>),
       any nil),
     x ]

[m2()
 -> let a := new(C2,symbol!("aa")) in
     (1,
      a.li := list<C1>(new(C1,symbol!("bb"))),
      a.li := list{m1(y) | y in a.li},
              // (list{m1(y) | y in a.li} as list<C1>),
      2) ]

//(m2())

// testing mutivalued slots
multiC <: thing                         // forward
multiC <: thing(ages:set<integer>,
                friends:set<multiC>,
                incomes:set<float>)

Peter :: multiC(ages = set<integer>(30), incomes = set<float>(15.0, 20.0))
Paul :: multiC()
Mary :: multiC(ages = set<integer>(30), friends = set<multiC>(Paul)) 

[multiTest()
  -> Peter.ages :add 40,
     Peter.incomes :delete 15.0,
     check("Peter1", Peter.ages = {30,40} & Peter.incomes = {20.0}),
     Peter.friends :add Paul,
     Peter.friends :add Mary,
     check("Peter2", Peter.friends = {Paul,Mary}),
     Mary.incomes :add 70.0,
     Mary.incomes :add 70.0,
     Paul.incomes := (Peter.incomes U Mary.incomes) as set<float>, 
     record("Paul incomes",Paul.incomes),
     check("Mary", Mary.incomes = {70.0}) ]

(multiTest())

// a new bug by Tibor -----------------------------------------------

obj <: thing(l0:integer,l1:integer,l2:integer,l3:integer)

ats2(t:obj) : tuple(integer,integer,integer,integer) 
 -> tuple(l0(t),l1(t),l2(t),l3(t))

ats(t:obj) : list[integer] 
  -> list(l0(t),l1(t),l2(t),l3(t))

oo :: obj(l0 = 123, l1 = 321, l2 = 789, l3 = 987)

(check("3", ats(oo) != ats2(oo)))

// small bug from G. Salaun -----------------------------------------

[TOTO <: thing(idex:integer, nom:string)]
BB :: TOTO(nom = "truc", idex = 2)

[get_data(i:integer) : string
   -> nom(some(t in TOTO | idex(t) = i)) ]

(check("4", get_data(2) = "truc"))

// deuxieme bug (compilation) ---------------------------------------

A <: thing()
B <: A()
C <: A()

[f10(x:A) -> true]
[f10(x:(B U C)) -> f10@A(x)]


ASylvain <: object(a:integer)
[bugSylvain() : void -> ASylvain(a = 12)] 
[nobugSylvain() : void -> ASylvain()] 


// introduce a new char for ALI
;[�(x:integer) -> x * x]
;(check("sqr �",�(2) = 4))

// a forward bug() -------------------------------------------------

A1 <: A
A2 <: A1
A3 <: A2

A1 <: A(x:integer)
A2 <: A1()
A3 <: A2(y:integer)

(let o := A3() in
  (o.x := 234,
   o.y := 123,
   check("conflict",o.x = 234),
   check("forward",A3 % A1.descendents)))

// tiny problem :-)
[tinyfoo() -> let x:integer := 1 in (if ((x := 2) != 3) printf("jhgjh") ) ]

// -----------------------------------------------------------

// v3.1 : test the interface


// a class
NiceClass <: object(x:integer, y:float)
SubClass <: NiceClass(x:integer = 12)

[method1(x:NiceClass) : integer -> 12]

[method2(self:NiceClass,a:float) : void -> self.y := a ]

[method3(a:NiceClass, y:integer) : boolean -> (a.x = y)]

(abstract(method3))

// this is a compiler pragma
// interface(NiceClass,method1,method2,method3)

niceTest :: SubClass()

[testInterface() -> method1(niceTest), method2(niceTest,12.0),
                    method3(niceTest,222) ]
                    
(testInterface())                    

// a method
[foo(x:property) : string -> "property"]
[foo(x:relation) : string -> "relation"]
[foo(x:table) : string -> "table"]

// interface(foo)

[testfoo(x:relation) : string -> foo(x) ]

(check("testfoo", testfoo(foo) = "property"))

// a bug about >= and <= 's redefinition
Ephemeral <: object()
AbstractVar <: Ephemeral()
IntVar <: AbstractVar()
AbstractConstraint <: Ephemeral()
IntConstraint <: AbstractConstraint(cste:integer = 0)
UnIntConstraint <: IntConstraint(v1:IntVar, idx1:integer)
GreaterOrEqualxc <: UnIntConstraint()
LessOrEqualxc <: UnIntConstraint()
Term <: Ephemeral(cste:integer = 0)
UnTerm <: Term(v1:IntVar, sign1:boolean = true)

[-(t:UnTerm) : UnTerm -> t.sign1 := not(t.sign1), t.cste := -(t.cste), t]
[>=(a:integer, t:UnTerm) : AbstractConstraint -> -(t) >= -(a)]
[>=(t:UnTerm, c:integer) : UnIntConstraint
-> if t.sign1
GreaterOrEqualxc(v1 = t.v1, cste = c - t.cste)
else LessOrEqualxc(v1 = t.v1, cste = t.cste - c)]

// test listargs

[sum(x:list) : integer
  => let d := 0 in (for y in x d :+ y, d) ]

[foo(x:listargs) : integer -> print(x), sum(x) ]
[bar(x:any,y:listargs) : integer -> sum(y) ]

[testFooBar()
 -> check("bar(1)",bar(1) = 0),
    check("bar(1,2,3)",bar(1,2,3) = 5),
    check("foo(1,2,3)",foo(1,2,3) = 6) ]

(testFooBar())

// test with global variables of all kinds

vA1 :: 10
vB1 :: 10.2
vC1 :: class

vA2:integer :: 10
vB2:float :: 10.2
vC2:class :: class

claire/wA3 :: 10
claire/wB3 :: 10.2
claire/wC3 :: class

claire/wA4:integer :: 10
claire/wB4:float :: 10.2
claire/wC4:class :: class

BTEST:boolean := true


[testV()
  ->  if not(BTEST) printf("---- false\n") else printf("--- true\n"),
      check("vars A",vA1 = vA2 & wA3 = wA4),
      check("vars B",vB1 = vB2 & wB3 = wB4),
      check("vars C",vC1 = vC2 & wC3 = wC4),
      vA2 := wA3, wA4 := vA1,
      vB2 := wB3, wB4 := vB1,
      vC2 := wC3, wC4 := vC1,
      check("2 vars A",vA1 = vA2 & wA3 = wA4),
      check("2 vars B",vB1 = vB2 & wB3 = wB4),
      check("2 vars C",vC1 = vC2 & wC3 = wC4),
      vA2 := 10 + wA3 - 2, wA4 := 10 + vA1 - 4,
      vB2 := 2.0 * wB3 * 2.0, wB4 := 2.2 * vB1 * 2.1,
      vC2 := wC3.superclass, wC4 := vC1.superclass ]
     
// something that must work
MONTHV:Interval :: (1 .. 12)
MONTHS :: (1 .. 12)

[fooru(x:MONTHS) : MONTHS
   -> MONTHV := MONTHS,
      printf("~S",MONTHS),
      for y in MONTHS print(y),
      MONTHV := (2 .. 3),
      x ]


(testV(), fooru(1))

// a bug from FXJ ? with v3.3.2
// INTERPRETED :: not(Id(compiler.active?))


(testOK())





