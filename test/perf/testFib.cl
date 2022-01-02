// fib test - the oldest claire file :)
fib(n:integer) : integer -> (if (n < 2) n else (fib(n - 1) + fib(n - 2)))

g(n:integer) : void
  -> (time_set(), fib(n), time_show())

// version with floats
[fab(n:float) : float -> if (n < 2.0) 1.0001 else fab(n - 1.0) + fab(n - 2.0)]

[tfab(n:float) 
  -> time_set(), fab(n), time_show()]

// bug for debugging
foo(n:integer) : integer
  -> (if (n < 1) (1 / n) else 1 + foo(n - 1)) 

// quicksort was added from Julia benchmark ----------

// float random generator
[random(f:float) : float
  -> let m := 10000000, n := random(m) in
       (f * float!(n) / float!(m))]


// call quicksort on a random array of 5000 
// too fast when compiled => added a x10
testSort(N:integer) 
 -> let lst := make_array(N,float,0.0) in 
   	(time_set(),
   	 for k in (1 .. 10)
        (for i in (1 .. N) lst[i] := random(1.0),
         qsort(lst,1,N)),
     time_show())

// test and see :)
tests(N:integer)
  -> let lst := make_array(N, float, 0.0) in 
   	  (for i in (1 .. N) lst[i] := random(1.0),
       qsort(lst,1,N),
       printf("sorted = ~S\n",list{lst[i] | i in (1 .. N)})) 

// quick sort from the Julia benchmark
qsort(a:float[], lo:integer, hi:integer) : float[]
-> let i := lo, j := hi in 
 (while (i < hi)
   let pivot := a[(lo + hi) / 2] in
       (while (i <= j)
           (while (a[i] < pivot) i :+ 1,
            while (a[j] > pivot) j :- 1, 
            if (i <= j)
              let z := a[i] in 
                 (a[i] := a[j], 
                  a[j] := z,
                  i :+ 1,
                  j :- 1)),
      if (lo < j) qsort(a, lo, j),
      lo := i,
      j := hi),
 a)


// execute all tests

all() ->
  (princ("fib(30)"),
   g(30),
   princ("fib(35)"),
   g(35),
   princ("fib(35.0)"),
   tfab(35.0),
   princ("10 x quicksort(5000)"),
   testSort(5000))

[main() -> 
  //[0] test fib ----- //,
  all(),
  exit(0) ]

