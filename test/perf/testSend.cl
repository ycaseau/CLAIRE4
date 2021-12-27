// ================== Step 6: exploration with exception handling ==========================================
// SEND+MORE = MONEY in a naive manner
// this code uses try/catch to be ready for constraint propagagation that will add cuts
// this naive version has no cuts !
[Bucket <: thing(value:integer = -1, minValue:integer = 0)]
S :: Bucket(minValue = 1)
E :: Bucket()
N :: Bucket()
D :: Bucket()
M :: Bucket(minValue = 1)
O :: Bucket()
R :: Bucket()
Y :: Bucket()

LB :: list<Bucket>(S,E,N,D,M,O,R,Y)
lused :: list<boolean>{false | i in (0 .. 9)}
CInt <: thing(value:integer = 0)

nbk :: CInt()

// tells if the solution is OK
[solOK() : boolean
  -> (S.value * 1000 + E.value * 100 + N.value * 10 + D.value) +
     (M.value * 1000 + O.value * 100 + R.value * 10 + E.value) =
     (M.value * 10000 + O.value * 1000 + N.value * 100 + E.value * 10 + Y.value) ]

[solve() : boolean
   -> solve(LB, 8) ]

[solve(l:list<Bucket>,i:integer) : boolean
  -> if (i = 0)
        (if solOK() true else (nbk.value :+ 1, contradiction!(), false))
      else let b := l[i], found := false in
          (for j in (b.minValue .. 9)
             (if not(lused[j + 1])
                 try (b.value := j,
                      lused[j + 1] := true,
                      if solve(l, i - 1)
                         (found := true, break())
                      else lused[j + 1] := false)
                 catch contradiction lused[j + 1] := false),
          found) ]
         
[tsolve()
  -> time_set(),
     if solve()
        (time_show(),
         printf("found a solution\n ~A ~A ~A ~A + ~A ~A ~A ~A = ~A ~A ~A ~A ~A",
               S.value,E.value,N.value,D.value,
               M.value,O.value,R.value,E.value,
               M.value,O.value,N.value,E.value,Y.value))
     else time_show(),
     printf("nbk = ~A\n",nbk.value)]

[all() -> tsolve()]     

[main() -> 
  //[0] test Send+More=Money ----- //,
  all(),
  exit(0) ]