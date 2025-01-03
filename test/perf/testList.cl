[lt(n:integer) : list<integer>
  -> let l := nil in
        (for i in (1 .. n) l := list<integer>(i, i + 1,i + 2,i + 3,i + 4,i + 5,i + 6,i + 7,i + 8,i + 9),
         l)]

[tlt(n:integer) -> time_set(), lt(n), time_show()]

// useful for debug : simple iteration
[el(n:integer) : integer
  -> let s := 0 in
       (for i in (1 .. n) s :+ n,
        s)]

// CLAIRE 1: tel(10M) -> 402ms
[tel(n:integer) -> time_set(), el(n), time_show()]


[lta(n:integer) : list<integer>
  -> let l := list<integer>(0) in
        (for i in (1 .. n)
           (l := list<integer>(i),
            for j in (1 .. i) add(l, j)),
         l)]

[tlta(n:integer) -> time_set(), lta(n), time_show()]


[ltr(n:integer) : integer
  -> let l := list<integer>(0), s := 0 in
        (for i in (1 .. n) l :add i,
         for k in (1 .. n)
            (s := 0,
             for i in (1 .. n)
                        for j in (1 .. i) s :+ l[i] * l[j]),
         s)]

[tltr(n:integer) -> time_set(), ltr(n), time_show()]

[ltw(n:integer)
  -> let l := list<integer>(0)  in
	(for i in (1 .. n) l :add i,
	 for k in (1 .. 100)
	    (for i in (1 .. n)
			 l[i] :+ 1)) ]

        
[tltw(n:integer) -> time_set(), ltw(n), time_show()]

[ltw1(n:integer)
  -> let l:list := list<integer>(0)  in
	(for i in (1 .. n) l :add i,
	 for k in (1 .. 100)
	    (for i in (1 .. n)
			 l[i] :+ 1)) ]

        
[tltw1(n:integer) -> time_set(), ltw1(n), time_show()]

[all()
  -> princ("tlt(1M)"),
     tlt(1000000),
     princ("tla(10K)"),
     tlta(10000),
     princ("tlr(200)"),
     tltr(200),
     princ("tlw(100K)"), 
     tltw(100000)]

[all2()
  -> princ("tlt(1M)"),
     tlt(1000000),
     princ("tla(10K)"),
     tlta(10000),
     princ("tlr(500)"),
     tltr(500),
     princ("tlw(100K)"), 
     tltw(100000)]


[main() -> 
  //[0] test lib ----- //,
  all(),
  princ("tlr(500)"),
  tltr(500),
  exit(0) ]
