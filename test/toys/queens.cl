// a naive N-queen problem

N:integer := 64

column[n:(1 .. N)] : (0 .. N) := 0     // 0 for no values yet
possible[x:(1 .. N), y:(1 .. N)] : boolean := true
countPos[n:(1 .. N)] : (0 .. N) := N 
 	
// two defeasible attributes    
store(column, possible, countPos)

// remove a possible value
forbid(x:(1 .. N),y:(1 .. N)) : void 
  -> (if (column[x] = 0 & possible[x,y])
        (possible[x,y] := false, 
         countPos[x] :- 1,
         if (countPos[x] = 0) contradiction!()))

// three rules that implement the allDifferent constraints
r1() :: rule( 
 		column[x] := z => for y in ((1 .. N) but x) forbid(y,z))

r2() :: rule( 	
        column[x] := z => let d := x + z in
                            for y in (max(1,d - N) .. min(d - 1, N))
                                forbid(y,d - y))
r3() :: rule( 	
        column[x] := z => let d := (z - x) in 
                            for y in (max(1,1 - d) .. min(N,N - d))
                                forbid(y,y + d))
	
// finds the position with the fewest choice, 0 if none
tightest()
  -> let m := N + 1, best := 0 in
       (for j in (1 .. N)
          (if (column[j] = 0 & countPos[j] < m) 
              (best := j, m := countPos[j])),
        best)
            

queens(n:integer) : boolean
		-> (if (n = 0) true
            else let q := tightest() in
			    exists(p in (1 .. N) |
					(possible[q,p]	&
                      branch( (column[q] := p, queens(n - 1)) ))))
	
// print the solution
  solution() -> list{column[i] | i in (1 .. N)}


// check the solution before blogging :)
[checkSol()
 -> for i in (1 .. N)
      for j in ((1 .. N) but i)
        (if (column[i] = column[j] |
            column[i] + i = column[j] + j |
            column[i] - i = column[j] - j)
            printf("Error at ~S:~S and ~S:~S\n", i,column[i], j,column[j]))]    

  
