// simple knapsack = complexity O(n2 x m)
LPOS:any :: nil
LSOL:any :: nil

// for debug
sum(s:list) : integer -> (let n := 0 in (for x in s n :+ x, n))

// how to find the best combination of items to put in a knapsack of capacity val
[knap(l:list<integer>,val:integer) : list
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
S6 :: list<integer>(10104, 28956, 34208, 16368, 26106, 18946, 21075, 39460, 17190, 23108, 53028, 20210, 28844, 
                    105172, 22980, 123204, 15097, 47188, 44850, 12374, 13220, 68210, 12330, 34148, 77750, 
                    80078, 20104, 66130, 12274, 177852, 19146, 24150, 39150, 15550)

// 600000 en 3.5s
S6a :: list<integer>(10104, 28956, 34208, 16368, 26106, 18946, 21075, 39460, 17190, 23108, 53028, 20210, 28844, 
                    105172, 22980, 123204, 15097, 47188, 44850, 12374, 13220, 68210, 12330, 34148, 77750, 
                    80078, 20104, 66130, 12274, 177852, 19146, 24150, 39150, 15550)


// trying to make a difficult problem
S7 :: list<integer>(10104, 28956, 34208, 16368, 26106, 18946, 21075, 39460, 17190, 23108, 53028, 20210, 28844, 
                    105172, 22980, 123204, 15097, 47188, 44850, 112374, 13220, 68210, 12330, 34148, 77750, 
                    80078, 20104, 66130, 12274, 177852, 19146, 24150, 39150, 15550, 222222, 333333, 444444, 555555,
                    110104, 38956, 32208, 16999, 26106, 18946, 21075, 39460, 17190, 23108, 53028, 20210, 28844, 
                    100172, 2293, 123000, 15999, 47188, 44850, 12374, 13222, 68999, 201298, 34001, 177055,
                    268435, 198325, 422311, 90234,135603, 91919, 105675, 200200, 123456, 321099)

// version 2 with integer sets
[knap2(l:list<integer>,val:integer) : void
  -> let n := length(l), m := max(< @ integer,l), mx := 0, best := 0,
      possible := list<boolean>{ false | i in (1 .. (n * m))},
      sol := list<integer>{ 0 | i in (1 .. (n * m))} in
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
       //list(best,list<integer>{l[j] | j in {j in (1 .. n) | sol[best][j]}}))]
       best)]


  // test examples 
  [go5() 
     -> let l := nil in
          (time_set(),
           l := knap(S6,150000),
           time_show(),
           printf("check ~S vs ~S\n solution: ~A\n",l[1],sum(l[2]), l))] 

 [g25() 
     -> let l := nil in
          (time_set(),
           l := knap2(S5,150000),
           time_show(),
           printf("check ~S vs ~S\n solution: ~A\n",l[1],sum(l[2]), l))] 

// tryu with g26(1000000)
[g26(u:integer) 
     -> let l := nil in
          (time_set(),
           l := knap2(S6,u),
           time_show(),
           printf("check ~S vs ~S\n solution: ~A\n",l[1],sum(l[2]), l))] 

[g27(u:integer) 
     -> let l := nil in
          (time_set(),
           l := knap2(S7,u),
           time_show(),
           printf("check ~S vs ~S\n solution: ~A\n",l[1],sum(l[2]), l))] 

 // look for a hard problem with S7
 [sc7(a:integer,b:integer) : void
   -> for i in (a .. b)
        (time_set(),
         let l := knap(S7,i) in
          (time_show(),
           printf("~A:~A vs ~A\n",i,l[1],sum(l[2]))))]
          
