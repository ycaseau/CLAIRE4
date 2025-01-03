// +------------------------------------------------------------+
// | bug5.cl                                                    |
// | last update: Jan 2000 - Y. Caseau                          |
// +------------------------------------------------------------+


// ---------------------------------------------------------------
// this file contains bugs related with floats & imported primitives
// ---------------------------------------------------------------

// integer bugs --------------------------------------------------------------------
[day_(i:integer) : integer 
    -> (i mod (365 * 24)) / 24]


// bugs from XL -----------------------------------------------------------------------

C <: thing(i:integer, f:float , l:list)
ob :: C()

(check("default", ob.f = 0.0),
 check("erase f", erase(f,ob) = 0.0))

Z <: thing(ff:float, ss:set<integer>)  // si on ajoute un slot foat ->
zz :: Z()
(zz.ss := set<integer>(),
 zz.ss :add 3)

(check("zz.ss",(zz.ss = set<integer>(3))))


/// classical stuff ---------------------------------

[sum(s:any) : float
  -> let d := 0.0 in (for x in s d :+ x, d) ]
 
[sum2(s:any) : float
  -> let d := 0.0 in (for x in s d :+ x, d) ]

[sum3(s:any) : float
  -> let d := 0.0 in (for x:float in s d :+ x, d) ]

(check("sum", sum(list(1.0,2.2)) = 3.2))
(check("sum2", sum2(list(1.0,2.2)) = 3.2))
(check("sum3", sum3(list(1.0,2.2)) = 3.2))
  
this_is_a_global:any :: 12
  
[foo() : float -> 1.0 + 545.0 ]

[bar(sze:float) : float -> sze :/ 546.0]

(check("barfoo",bar(foo()) = 1.0))

// ---------- dumb float function --------------------
[sinis(x:float) : float
 -> let y := x / 6.28, f := y - float!(integer!(y)), z := f * 6.28 in
      (if (z > 3.14) -(sinis(z - 3.14))
       else (z / (3.14 / 2))) ]

// --------- parsing a float -------------------------

[private/Gparse_percentage(s:string) : (float U {false})
  -> let value1 := 0.0,
         value2 := 0.0,
         start := 1,
         end := start
     in (while (end <= length(s) & '0' <= s[end] & s[end] <= '9')
           end :+ 1,
         if (s[end] = '.')
           (if (start < end)
              value1 := float!(read(substring(s, start, end - 1))),
            start := end + 1,
            end := start,
            while (end <= length(s) & '0' <= s[end] & s[end] <= '9')
              end :+ 1,
            value2 := (float!(read(substring(s, start, end - 1)))
                       / (10.0 ^ float!(end - start))),
            value1 + value2))
]

[test() : (float U {false})
  -> Gparse_percentage("12.23")
]

[new_abs(x:float) : float -> if (x >= 0.0) x else -(x) ]

(check("parse", new_abs(test() - 12.23) < 1e-6))


// floating stuff (very cute bug)
[testPreviousBugs() : void
  -> let x := 1.0, y := 2.0 in (if (x != 0.0 | y != 0.0) true else false)
]

(check("previous",(testPreviousBugs(),true)))

// creates a bug if it is a macro ?
[sqr(x:integer) : float -> let y := float!(x) in (y * y) ]
[stdev(p:property,s:any) : any
  => let d := 0, %c := 0, d2 := 0.0 in
        (for x in s
           let y := p(x) in (d :+ y, %c :+ 1, d2 :+ sqr(y)),
         if (%c > 0) sqrt((d2 / float!(%c)) - sqr(d / %c))) ]


(record("stdev",stdev(random,list(100,100,100,100,100))))

// bugs from ThB and KA
[make2DArray() : float[][]
  -> array!(list<float[]>{make_array(10,float,0.0) | i in (1 .. 10)}) ]

[fooTH() : void
  -> let dist:float[][] := make2DArray() in nil ]


[buggy() : void
  -> let a:float := 1e100 in
       (for i in (1 .. 1)
          let b:float := 0.0 in
             a := b ) ]

// float alignment bugs - this should be run on all platforms !
A1 <: object(x:float = 1.01,y:integer,x2:float = 2.02)

B1 <: A1(u:integer,x3:float = 3.03,x5:float = 5.05)
B2 <: A1(x4:float = 4.04,vv:integer,x6:float = 6.06)

aB1 :: B1()
aB2 :: B2()
bB1 :: B1(x3 = 30.0)
bB2 :: B2(x4 = 40.0)

A2 <: object(x:float = 1.01,y:integer,x2:float = 2.02,y2:integer)

C1 <: A2(u:integer = 1234,x3:float = 3.03,x5:float = 5.05)
C2 <: A2(x4:float = 4.04,vv:integer,x6:float = 6.06)
C3 <: A2(ua:integer = 1234,ub:integer,x6:float = 6.06)

[foo(x:C1,a:float) : void -> x.x3 := a]
[bar(x:C1,a:integer) : void -> x.u := a]

aC1 :: C1()
aC2 :: C2()  
aC3 :: C3()

bC1 :: C1(x3 = 30.0, u = 2222, y2 = 20)
bC2 :: C2(x4 = 40.0, vv = 1277)
bC3 :: C3(ua = 321, ub = 322, x6 = 60.0)           
           
[testA() 
  -> check("aB1", (aB1.x3 = 3.03 & aB1.x5 = 5.05)),
     check("aB2", (aB2.x4 = 4.04 & aB2.x6 = 6.06)),
     let p := x5 in write(p,aB1,50.0),
     let p := x6 in write(p,aB2,60.0),
     check("dyn write B1/B2", aB1.x5 = 50.0 & aB2.x6 = 60.0),
     let p := x5 in write(p,bB1,50.0),
     let p := x6 in write(p,bB2,60.0),
     check("bB1", (bB1.x3 = 30.0 & bB1.x5 = 50.0)),
     check("bB2", (bB2.x4 = 40.0 & bB2.x6 = 60.0)),
     // second case
     check("aC1", (aC1.x3 = 3.03 & aC1.x5 = 5.05)),
     check("aC2", (aC2.x4 = 4.04 & aC2.x6 = 6.06)),
     let p := x5 in write(p,aC1,50.0),
     let p := x6 in write(p,aC2,60.0),
     check("dyn write C1/C2", aC1.x5 = 50.0 & aC2.x6 = 60.0),
     let p := x5 in write(p,bC1,50.0),
     let p := x6 in write(p,bC2,60.0),
     check("bC1", (bC1.x3 = 30.0 & bC1.x5 = 50.0)),
     check("bC2", (bC2.x4 = 40.0 & bC2.x6 = 60.0)),
     // C3 test
     check("aC3", (aC3.ua = 1234 & aC3.x6 = 6.06)),
     check("bC3", (bC3.ua = 321 & bC3.x6 = 60.0)),
     let p := ua in (write(p,aC3,9999), write(p,bC3,9999)),
     check("dyn write C3", aC3.ua = 9999 & bC3.ua = 9999) ]
     
(//[0] run test A ! //,
 testA())

// bug from ThB
A <: object(val:integer)
[f(a:A) : float -> max(2.0,float!(a.val))] 
aa :: A(val = 12)
[g(n:integer) : void 
-> let i := 1 in while (i < n) (f(aa),i :+ 1)]

     
Ax <: object(valx:float = 12.0)
(store(valx))
[f(a:Ax,doStore?:boolean) : void -> a.valx :- 1.2 * 3.0]
[testAx(n:integer) -> let a := Ax() in (for i in (1 .. n) f(a,false))]

Bx <: Ax()
[getMaxVal(a:Ax) : float -> 30.0]
[getMaxVal(b:Bx) : float -> 40.0]
[f(a:Ax) : float -> getMaxVal(a)]

AA:Ax :: Ax()

(//[0] test g and testAx //,
 g(100000),
 //[0] test Ax //,
 testAx(100000))


// ---------------- test the random number generator ------------------------

// distance
[dist(x:integer) : float -> if (x < 0) float!(-(x)) else float!(x) ]

AA:float := 0.0
BB:float := 0.0
// test random(n)
[testr(n:integer) : float
   -> let %max := -10, %min := n + 10, %sum := 0.0, %cut := 0, k := 100000, %q := 0.0 in 
         (for i in (1 .. k)
            let x := random(n) in 
               (%max :max x,
                %min :min x,
                %sum :+ float!(x),
                if ((2 * x) < n) %cut :+ 1),
          %q := (dist(%min) + dist(%max - (n - 1)) + 
                  abs((%sum / float!(k))  - float!(n - 1) / 2.0))  / float!(n),
          AA := (dist(%min) + dist(%max - (n - 1)) + 
                  abs((%sum / float!(k))  - float!(n - 1) / 2.0)),
          BB :=  float!(n),
          //[0] test random(~A) -> [~A - ~A]: ~A (~A%) -> ~A (quality) // n, %min, %max, %sum / float!(k), %cut / (k / 100), %q,
          //[0] q = (~A + ~A + ~A) / ~A // dist(%min), dist(%max - (n - 1)), abs((%sum / float!(k))  - float!(n - 1) / 2.0), float!(n),
          assert(%max < n & %min >= 0),
          %q) ]

// reproduce bug
// [foo()
//   -> let %q :=           


[testR()
  ->  check("R2", testr(2) < 0.002),
      check("R5", testr(5) < 0.002),
      check("R10", testr(10) < 0.002),
      check("R50", testr(50) < 0.002),
      check("R100", testr(100) < 0.002),  
      check("R200", testr(200) < 0.002),  
      check("R1000", testr(1000) < 0.002),
      check("R10000", testr(10000) < 0.002),
      check("R50000", testr(50000) < 0.002),
      check("R100000", testr(100000) < 0.002),
      check("R1000000", testr(1000000) < 0.002),
      check("R10000000", testr(10000000) < 0.002),
      check("R100000000", testr(100000000) < 0.002) ]


(testR())

// test lambdas with CLAIRE4 syntax -------------------------------------------------------
LL :: (x:integer){x + 1}

LL2 :: (x){princ("Hello World\n")}

(//[0] --- here is (x){x + 1} : 12 -> ~A ----- // funcall(LL,12),
 show(LL))

[testLambda()
  -> let l1 := (x:integer,y:integer){x + y},
         l2 := (x){x * x} in
       (funcall(LL2,void),
        check("apply",apply(l1,list(1,2)) = 3),
        check("funcall",funcall(l1,5,5) = 10),
        check("map",map(l2,{1,2}) = {1,4}),
        check("map list",length(map(l2,list{i | i in (1 .. 10)})) = 10)) ] 

(testLambda())



// test the use of ports -------------------------------------------
[openFile(s:string) : port
  -> let fn := s /+ ".log",
         p := fopen(fn ,"w") in
       (use_as_output(p),
        printf("------------------  Experiment version ~A on ~A",release(),date!(1)),
        p) ]

(fclose(openFile("sample")))

(testOK())

