// =======================  step4: dictionnaries ===========================================
// this test file uses the "relation dict" that is a standard CLAIRE 4 feature
// later we shall add the new Dict objects
stlist :: list<string>("ab","fg","sh","ai","po","gu","ms","fx","wz","ks")


// create a list of n string
[lstring(n:integer) : list<string>
  -> let l := list<string>() in
         (for i in (1 .. n)
             l :add (stlist[i mod 10 + 1] /+ stlist[(i / 10) mod 10 + 1] /+ string!(i)),
          l) ]

// test the speed of string creation, 1000 -> 2s (0.4s compiled)
[tls(n:integer)
  -> time_set(),
     for i in (1 .. n) lstring(n),
     time_show()]

Dict[s:string] : integer := 0

// creation
[dt(n:integer) : void
   -> let l := lstring(n) in
        (for k in (1 .. 100)
            for i in (1 .. n)
                Dict[l[i]] := i + k) ]

// test speed with 100 update of dict of size 10000 : failure
// there is a bug in the hashing dispertion => conflict rate is obscene => quadratic behavior
// TO BE FIXED BEFORE PUBLISHING
// Claire4 : tdt(10K) : 900ms :)
[tdt(n:integer)
   -> time_set(), dt(n), time_show()]


// test the values
[dr(n:integer) : void
   -> let l := lstring(n) in
         (for k in (1 .. 100)
            for i in (1 .. n)
             (if (Dict[l[i]] != i + 100) 
                printf("error at ~A: ~A vs ~A",i,Dict[l[i]],i + 100))) ]

// 
[tdr(n:integer)
   -> time_set(), dr(n), time_show()]

all() ->
 (princ("tls(1000)"), tls(1000),
  princ("tdt(10000)"), tdt(10000),
  princ("tdr(10000)"), tdr(10000))

[main() -> 
  //[0] test Dict ----- //,
  all(),
  exit(0) ]