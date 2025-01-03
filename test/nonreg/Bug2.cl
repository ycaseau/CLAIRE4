// +------------------------------------------------------------+
// | bug2.cl                                                    |
// | last update: Jan 2000 - Y. Caseau                          |
// +------------------------------------------------------------+


// ---------------------------------------------------------------
// this file contains bugs related to arrays and maps (CLAIRE4)
// ---------------------------------------------------------------

// heavy tests on float arrays

FObj <: object(x:integer,y:float[])

Global:float[] := make_array(100,float,0.0)

Ftest:float := 0.0

[sumF(x:any) : float => let d := 0.0 in (for y in x d :+ y, d) ]


[Finit()
  -> for i in (1 .. 100) Global[i] := float!(i * i),
     Ftest := sumF(Global),
     record("Ftest",Ftest)]

[Foo()
  -> Ftest := sumF(Global)]

[Floop()
 -> let A := make_array(100,float,0.0),
        u := FObj(y = make_array(100,float,0.0)) in
       (u.x := length(instances(FObj)),
        //[0] Floop (~A) ------- // u.x,
        for i in (1 .. 100) A[i] := Global[i],
        for i in (1 .. 100) u.y[i] := A[i],
        for i in (1 .. 100) Global[i] := u.y[i],
        if (sumF(Global) != Ftest |
             exists(u2 in FObj | sumF(u2.y) != Ftest)) error("bug")) ]

[Fgo(i:integer)
  -> Finit(),
     for j in (1 .. i) Floop() ]

(Fgo(20))
    
// a bug from ALI using matrix (arrays of floats)

cToto <: thing(x:float[])
myTheta:float[] := make_array(10,float,1.0)

[access(i:integer) : float => myTheta[i]]

tata :: cToto(x = make_array(10,float,1.0))

[buggy() : void -> tata.x[2] := access(2) ]

// (verbose() := 3, error("sop1"))

Matrix <: object(_n:integer,       // number of lines
                 _m:integer,       // number of columns
                 _v:float[])       // content (size n*m, line by line)  // note:array => bug

ephemeral(Matrix)
          
self_print(M:Matrix) -> printf("Matrix[~Ax~A]",M._n,M._m)

see(M:Matrix)
  -> (for i in (1 .. M._n)
       printf("~I~I\n",
              (if (i = 1) princ("\n(") else princ(" ")),
              for j in (1 .. M._m) (print(M[i,j]), princ(" "))))


// Creators -------------------------------------------------------------
[makeNullMatrix(n:integer,m:integer) : Matrix
-> let M := Matrix(_n = n,_m = m,_v = make_array(n * m,float,0.0))
    in M]

[makeIdMatrix(n:integer) : Matrix
-> let M := Matrix(_n = n,_m = n,_v = make_array(n * n,float,0.0))
    in (for i in (1 .. n)
    	(M._v[(i - 1) * n + i] := 1.0),
       M)]

// Access  -------------------------------------------------------------
[nth(M:Matrix,i:integer,j:integer) : float
=>	M._v[(i - 1) * M._m + j]]

[nth=(M:Matrix,i:integer,j:integer,x:float) : float
=> let v := x in (M._v[(i - 1) * M._m + j] := v, v) ]

// check simple functions
foo(M:Matrix) : float -> M[1,1]

foobar(M:Matrix) : float -> (if (M[1,1]  > 0.0) M[1,1] else 0.0)

bar(M:Matrix) : float -> (M[1,2] + 3.0)

Theta:Matrix :: makeIdMatrix(10)
AA:list[float] :: list<float>()

[macro1(x:list[float],y:float) : float
  => (macro2(x,y) / y) ]

[macro2(x:list[float],y:float) : float
  => let v := y in
       (for k in x v :+ k, v) ] 

makeMatrix(i:integer,j:integer) : Matrix
  -> Matrix(_m = i, _n = j, _v = make_array(i * j,float,0.0))

M2 :: makeMatrix(10,10)

(access(2),
 buggy(),
 see(M2),
 for i in (1 .. 10)
  for j in (1 .. 10)
    M2[i,j] := float!(i) * float!(j))


// test
[YAbug()
   -> tata.x[2] := Theta[1,2],
      Theta[3,4] := Theta[1,2],
      let eqx := 0.0 in (AA :add (eqx := macro1(AA,eqx))) ]

(YAbug())

// check that it would work with an untyped version -----------------------------------------
Motrix <: object(_n:integer,       // number of lines
                 _m:integer,       // number of columns
                 _v:any[])       // content (size n*m, line by line)  // note:array => bug

ephemeral(Motrix)
self_print(M:Motrix) -> printf("Motrix[~Ax~A]",M._n,M._m)

makeMotrix(i:integer,j:integer) : Motrix
  -> Motrix(_m = i, _n = j, _v = make_array(i * j,any,0.0))

// Access  -------------------------------------------------------------
[nth(M:Motrix,i:integer,j:integer) : any
=>	M._v[(i - 1) * M._m + j]]

[nth=(M:Motrix,i:integer,j:integer,x:float) : any
=> let v := x in (M._v[(i - 1) * M._m + j] := v, v) ]


Mo2 :: makeMotrix(10,10)

(for i in (1 .. 10)
  for j in (1 .. 10)
    Mo2[i,j] := float!(i) * float!(j),
 check("Motrix",Mo2[2,2] = 4.0))


// bug Arnaud (actually a compilation bug that involves arrays) --------------------------------
toto <: object()

//c_test(`(let a := 1,b := 2,c := make_array(3,toto,toto()) in (for i in (a .. b) (c[i] := toto()))))
[foo()
 -> let a := 1,b := 2,c := make_array(3,toto,toto()) in
        (for i in (a .. b) c[i] := toto()) ]


(foo())

// that should be allowed
tota:float[][] :=  make_array(10,float[],make_array(5,float,0.0))

// this one works
toti:any :=  make_array(10,float[],make_array(5,float,0.0))

// a bug from ThB
arr :: make_array(100,boolean,false)

[foo_array()
  -> //[0] start store .... //,
     store(arr,21,true,true),
     //[0] first set store .... //,
     for i in (1 .. 100) store(arr,i, (i mod 2) = 0, true),
     printf("world ~A: ~S\n",world?(),list{arr[i] | i in (1 .. 20)}),
     world+(),
     //[0] second store .... //,
     for i in (1 .. 100) store(arr,i, (i mod 5) = 0, true),
     printf("world ~A: ~S\n",world?(),list{arr[i] | i in (1 .. 20)}),
     world+(),
     //[0] third store .... //,
     for i in (1 .. 100) store(arr,i, (i mod 10) = 0, true),
     printf("world ~A: ~S\n",world?(),list{arr[i] | i in (1 .. 20)}),
     world-(),
     for i in (1 .. 100) check("store on array (2)",arr[i] = ((i mod 5) = 0)),
     printf("world ~A: ~S\n",world?(),list{arr[i] | i in (1 .. 20)}),
     world-(),
     for i in (1 .. 100) check("store on array (1)",arr[i] = ((i mod 2) = 0)),
     printf("world ~A: ~S\n",world?(),list{arr[i] | i in (1 .. 20)}) ]
    
(foo_array())

// --------------------------------- part 2 : simple map tests -------------------------------

M1 :: map<any,any>("foo":(1 + 3),class:"class",(3 * 4):12)
M2 :: map<integer,integer>()

[testMap()  
 -> check("M1", get(M1,12) = 12),
    for i in (1 .. 100) put(M2,i,i + 1),
    check("M2", get(M2,50) = 51) ]

(testMap())

(testOK())

  
