// ========================== step 3: set creation, membership ==========================================

// iterative creation
[sta(n:integer) : set<integer>
  -> let l := set<integer>() in
	   (for i in (1 .. n)
	       (l := set<integer>(i + 1),
	        for j in (0 .. i) l :add j,
            for j in (0 .. i) l :add (i - j)),
         l)]


// insert : 2000 en 0.15s  -> 14 ms compiled
[tsta(n:integer) -> time_set(), sta(n), time_show()]

[sta0(n:integer) : set<integer>
  -> let l := set<integer>() in
	   (for i in (1 .. n)
	       (l := set<integer>(i + 1)),
         l)]

[tsta0(n:integer) -> time_set(), sta0(n), time_show()]         

// membership test
[stm(n:integer) : integer
   -> let s := set<integer>(),
          m := 0 in
       (for i in (1 .. n) s :add i,
        for i in (1 .. 100)
           (for j in (1 .. n)
               (if contain?(s,j) m := m + 1),
            for j in (1 .. n)
               (if contain?(s,j + n) m := m + 1)),
        m / 100)]

// membership : 100 000 en 6s / 1s (not good !)
[tstm(n:integer) -> time_set(), stm(n), time_show()]


all() ->
 (princ("tsta(2000) "), tsta(2000),
  princ("tstm(100000) "), tstm(100000))


[main() -> 
  //[0] test Set ----- //,
  all(),
  exit(0) ]