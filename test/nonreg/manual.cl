// queens example (included in CLAIRE doc)
column[n:(1 .. 8)] : (0 .. 8) := 0                      // 0 means unknown
possible[x:(1 .. 8), y:(1 .. 8)] : boolean := true

store(column, possible)
	
r1() :: rule( 
 		column[x] := z => for y in ((1 .. 8) but x) possible[y,z] := false)

r2() :: rule( 	
        column[x] := z => let d := x + z in
                            for y in (max(1,d - 8) .. min(d - 1, 8))
                                possible[y,d - y] := false )

r3() :: rule( 	
        column[x] := z => let d := (z - x) in 
                            for y in (max(1,1 - d) .. min(8,8 - d))
                                possible[y,y + d] := false)


[queens(n:(0 .. 8)) : boolean
		-> (if (n = 0) true
			  else exists(p in (1 .. 8) |
					(possible[n,p]	&
                 branch( (column[n] := p, queens(n - 1)) ))))]

[qsolve() 
  -> check("queens", (queens(8) = true)),
     printf("solution is ~A\n",list{column[i] | i in (1 .. 8)}) ]

(qsolve())


// Hset data structure
htable <: object(	count:integer = 0,
					index:integer = 4,
					arg:list<any> = list<any>())

set!(x:htable) : set -> {y in x.arg | known?(y)}
self_print(h:htable) -> printf("#~S",set!(h))

htable!() : htable -> 
  let h := htable() in 
    (h.arg := make_list(^2(h.index - 3), unknown),
     h)

insert(x:htable,y:any) : void
		-> let l := x.arg in
				(if (x.count >= length(l) / 2)
			 		(x.arg := make_list(^2(x.index - 3), unknown),
					 x.index :+ 1, x.count := 0,
					 for z in {y in l | known?(y)} insert(x,z),
					 insert(x,y))
				 else let i := hash(l,y) in
					(until (l[i] = unknown | l[i] = y)
						   (if (i = length(l)) i := 1 else i :+ 1),
					 if (l[i] = unknown)
						(x.count :+ 1, 
             l[i] := y)))


[testh() 
  -> let h := htable!(), testS := {"class", 12, 18, class, 34.56, list(1,2)} in
       (for i in testS insert(h,i),
        check("htable", set!(h) = testS))]

(testh())

(testOK())
