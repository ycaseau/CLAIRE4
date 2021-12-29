// +------------------------------------------------------------+
// | bug7.cl                                                    |
// | last update: Jan 2000 - Y. Caseau                          |
// +------------------------------------------------------------+



// ---------------------------------------------------------------
// this file contains bugs related with worlds
// ---------------------------------------------------------------

// a really dumb test file that needs to be extended ...

homme <: thing(age:integer = 0, money:float = 0.0, knows:class)
store(age,money,knows)

Albert :: homme()

[testH() : void
  -> let titi := Albert in
        (check("default",age(titi) = 0),
         titi.knows := integer,
         choice(),
         age(titi) := 1,
         titi.money := 12.4,
         commit(),
         check("age",age(titi) = 1),
         check("money",money(titi) = 12.4),
         choice(),
         age(titi) := 2,
         money(titi) := 12.1,
         knows(titi) := float,
         backtrack(),
         check("age2",titi.age = 1),
         check("money2",titi.money = 12.4),
         check("knows",titi.knows = integer)) ]

(testH())

femme <: thing(friends:set[homme],
               xs:set[integer],
               ys:list[float])

store(friends,xs,ys)

Arthur :: homme()
Mat :: homme()
Mary :: femme()


[testF() : void
  -> let titi:femme := Mary in
        (check("default",xs(titi) = {}),
         titi.friends :add Arthur,
         choice(),
         titi.xs :add 1,
         titi.ys := list<float>(12.34),
         backtrack(),
         check("xs1",titi.xs = {}),
         check("ys1",titi.ys = list<float>()),
         titi.xs :add 1,
         titi.ys := add(copy(titi.ys), 1.34),       // add is no longer defeasible in CLAIRE 4
         choice(),
         titi.xs :add 2,
         titi.ys := add(copy(titi.ys), 2.34),
         titi.friends :add Mat,
         backtrack(),
         check("xs2",titi.xs = set<integer>(1)),
         check("ys2",titi.ys = list<float>(1.34)),
         check("friends",titi.friends = {Arthur})) ]

(testF())
(printf("----------- start testing tables ------------------------"))

// testing array-based tables
Table1[x:(0 .. 10)] : integer := 1
Table2[x:(0 .. 10)] : float := 0.0
Table3[x:(0 .. 10)] : class := integer
Var1:integer := 0
Var2:list[integer] := list<integer>{0 | i in (1 .. 10)}
store(Table1, Table2, Table3)
store(Var1)

[testK()
 ->  Table1[2] := 2,
     Table2[2] := 2.01,
     Table3[2] := float,
     Var1 := 123,
     store(Var2,3,123),
     choice(),
     Table1[2] := 3,
     Table2[2] := 3.01,
     Table3[2] := list,
     Var1 := 2,
     store(Var2,3,14423),
     backtrack(),
     check("K1",Table1[2] = 2),
     check("K2",Table2[2] = 2.01),
     check("K3",Table3[2] = float),
     check("K4",Var1 = 123),
     check("K5",Var2[3] = 123) ]

(testK())

// a test that creates a lot of worlds and expands
Table4[x:integer] : integer := 0
store(Table4)

[foo+(i:integer)  : void
  -> if (i > 1) foo+(i - 1),
     choice(),
     for j in (1 .. i * i) Table4[j] := i + j ]

[foo-(i:integer) : void
 -> let i2 := i - 1 in
      (if (i2 = 0) backtrack()
       else (backtrack(),
             for j in (1 .. i * i)
                (if (j <= (i2 * i2))
                   ( check("loop",Table4[j] = i2 + j),
                     if (Table4[j] != i2 + j)
                        error("Table[~A] = ~A != ~A",j,Table4[j], i2 + j) )
                 else check("loop0",Table4[j] = 0)),
             foo-(i2))) ]

[testA(n:integer) -> foo+(n), foo-(n), nil]

(testA(10))

// same test with a different implementation
Table5[x:(0 .. 200)] : integer := 0
store(Table5)

[foo2+(i:integer)  : void
  -> if (i > 1) foo2+(i - 1),
     choice(),
     for j in (1 .. i * i) Table5[j] := i + j ]

[foo2-(i:integer) : void
 -> let i2 := i - 1 in
      (if (i2 = 0) backtrack()
       else (backtrack(),
             for j in (1 .. i * i)
                (if (j <= (i2 * i2))
                    check("loop2",Table5[j] = i2 + j)
                 else check("loop02",Table5[j] = 0)),
             foo2-(i2))) ]

[testB(n:integer) -> foo2+(n), foo2-(n), nil]

(testB(10))

// test Store[X] --------------------------------------------------------

// =========================== useful and reusable ================================
Store[of] <: object(of:type, value:any, world:integer = -1)
ephemeral(Store)

[self_print(x:Store) -> printf("store(~S)",get(value,x)) ]

[write(x:Store[of = X],y:X) : void
  -> (if (world_id() > x.world & x.value != y)
         (//[0] ===> call put_store value, ~S, ~S (old value is ~S)// x,y,x.value,
          put_store(value,x,y,true), 
          put_store(world,x,world_id(),true))
      else (//[5] ===> simple write ~S.value = ~S ... // x,y,
            x.value := y)) ]

// debug
[fob(x:Store[of = X],y:X) : void -> x.value := y ]


[read(x:Store[of = X]) : type[X] => x.value]

// ================================= end of source code library ======================

A <: thing(x:Store<integer>, y:Store<string>)

// will support the syntax a.x := value => translated into write(a.x,value)
reify(x,y)

// create a test object with two store slots
a :: A(x = Store(integer), y = Store(string))

[test(n:integer)
  -> (if (n = 0) nil
      else let x1 := a.x, y1 := a.y in
         (//[0] ==== create a world with a.x = ~S and a.y = ~S // x1, y1,
          choice(),
          for j in (1 .. 10) (a.x := j, a.y := string!(j)),
          test(n - 1),
          backtrack(),
          //[0] === after backtrack a.x = ~S and a.y = ~S // a.x, a.y,
          check("Store", x1 = a.x & y1 = a.y))) ]


b :: get(x,a)      // debug hold to bucket

(let y:any := 3 in (a.x := y),      // check that poorly typed works !
a.x := 0,                           // should be faster :)
test(10))              // need to initialize

// bug from ThB & EGaudin
[test1()
  -> a.x := 0,
     world+(),
     a.x := 0,
     a.x := 1,
     world-(),
     check("test1",a.x = 0) ]

[test2()
  -> a.x := 123,
     world+(),
     for i in (1 .. 10)
        (world+(),
         a.x := i,
         commit(),
         check("commit", a.x = i)),
     world-(),
     check("test2",a.x = 123) ]


(test1(), test2(), test2(),
 check("World number", world?() = 0))

// mise a jour 
[toto <: object()]

//Claire4 => we will need a declaration : instanced(toto)
(instanced(toto))

Abruti :: toto()

(check("instances", length(instances(toto)) = 1))
(printf("point3 ------------------ \n"))

// A cute bug from Pierre Novat
*world_res*:boolean :: false

*Y*:float :: 0.0

store(*Y*)

[claire/barNone() : any ->
   backtrack(0),
	*Y* := 0.0,
	for x in (1 .. 1000)
		(choice(), 
       *Y* :+ 1.0,
       if (x mod 100 = 0) printf("world = ~A, Y = ~A\n",world?(),*Y*)
       ),
   printf("end of 1st loop : world = ~A, Y = ~A\n",world?(),*Y*),
	*world_res* := true,
  	for y in (1 .. 100)
		(backtrack(1000 - (10 * y)),
       if (y mod 10 = 0) printf("world = ~A, Y = ~A\n",world?(),*Y*),
       if (*Y* != float!(1000 - (10 * y)))
	    (*world_res* := false)),
	*world_res*]

(check("Novat's test",barNone() = true))

// small test with cast
[fooCast()
 -> let x:any := 1, y:any := 2, z:integer := 3 in
        (z := (x as integer) + (y as integer),
         check("cast+", (x as integer) + (y as integer) = 3)) ]

(fooCast()) 

(testOK())
