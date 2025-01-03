// simple knapsack = complexity O(n2 x m)
LPOS:any :: nil
LSOL:any :: nil

// for debug
sum(s:list) -> (let n := 0 in (for x in s n :+ x, n))

// how to find the best combination of items to put in a knapsack of capacity val
[knap(l:list<integer>,val:integer) : void
  -> let n := length(l), m := max(< @ integer,l), mx := 0, best := 0,
      possible := list{ false | i in (1 .. (n * m))},
      sol := list{ nil | i in (1 .. (n * m))} in
      (// list of size ~S, max is ~S // n * m,m,
      // LPOS := possible,
      // LSOL := sol,
       for x in l 
        (let i := mx in
           (while (i > 0)
             (if possible[i] 
                 ( possible[i + x] := true,
                   sol[i + x] := sol[i] /+ list(x)),
              i :- 1)),
         possible[x] := true,
         sol[x] := list(x),
         mx := (mx + x)),
       for i in (1 .. min(mx,val)) 
            (if possible[i] best := i),
       list(best,sol[best]))]

// test set
S0 :: list<integer>(12,35,17,44)
S1 :: list<integer>(103, 57, 208, 68, 107, 47, 176, 60, 195, 107)
S2 :: list<integer>(103, 57, 208, 68, 107, 47, 176, 60, 195, 107, 28, 210, 44, 177, 81, 204, 99, 188, 57, 78)
S3 :: list<integer>(104, 56, 208, 68, 106, 46, 176, 60, 190, 108, 28, 210, 44, 172, 80, 204, 98, 188, 50, 76, 
                    120, 210, 30, 148, 50, 78, 104, 130, 74, 52, 146, 150, 150, 50)
S4 :: list<integer>(1103, 2057, 3208, 968, 1207, 3047, 2176, 860, 2195, 1307, 928, 2210, 
                    3044, 2177, 1081, 2204, 999, 2188, 1057, 2478)
S5 :: list<integer>(10104, 8956, 4208, 16368, 6106, 18946, 21076, 9460, 17190, 23108, 13028, 20210, 8844, 
                    5172, 980, 3204, 15098, 7188, 4850, 12376, 13220, 8210, 12330, 4148, 7750, 
                    10078, 20104, 16130, 12274, 7852, 19146, 4150, 9150, 5550)

// version 2 with integer sets
[knap2(l:list<integer>,val:integer) : void
  -> let n := length(l), m := max(< @ integer,l), mx := 0, best := 0,
      possible := list{ false | i in (1 .. (n * m))},
      sol := list{ 0 | i in (1 .. (n * m))} in
      (for i in (1 .. n)
        let x := l[i] in
        (let j := mx in
           (while (j > 0)
             (if possible[j] 
                 ( possible[j + x] := true,
                   sol[j + x] := sol[j] + ^2(i)),
              j :- 1)),
         possible[x] := true,
         sol[x] := ^2(i),
         mx := (mx + x)),
       for i in (1 .. min(mx,val)) 
            (if possible[i] best := i),
       list(best,list{l[j] | j in {j in (1 .. n) | sol[best][j]}}))]