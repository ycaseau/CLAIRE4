// A Bench from Ilog Rules .....
// =============================

// this is a v3.2 version
// to do: get a clean 3.0 version !! (hopefully on Desktop)
// this looks like a piece of ... :-(

DEBUG :: 5
N :: 32
NUP :: (N + 1)

Guest <: object(male?:boolean = true,
                hobbies:integer = 0,       // actually, a bit vector
                seated?:boolean = false,
                canBeNext?:boolean = true,
                idx:integer = 0)

// track the instances
(instanced(Guest))

[self_print(g:Guest)  : void -> printf("G~A",g.idx) ]

table1[i:(1 .. NUP)] : Guest := unknown

MasterOfCeremonies <: thing(currentPlace:(1 .. NUP) = 1)

FXJ :: MasterOfCeremonies()

// this program uses a search
store(table1, canBeNext?, seated?, currentPlace)

// various optimization for compile-time
known!(table1,hobbies)
// final(Guest)    

Nfiring:integer := 0                     // couter

// this is a rule to place the special gest G2
placeSecond() :: rule(
   F.currentPlace := p & p = 2 
   => for G2 in {x in Guest | not(x.seated?)}
        (if (G2.hobbies and table1[1].hobbies = 0)
            (G2.canBeNext? := false, Nfiring :+ 1 )))

placeNext()  :: rule(
   F.currentPlace := p & p > 2
   =>  for G2 in {x in Guest | not(x.seated?)}
        (if (G2.hobbies and table1[p - 1].hobbies and table1[p - 2].hobbies = 0)
            (G2.canBeNext? := false,
             Nfiring :+ 1,
             trace(DEBUG,"impossible to seat ~S at place ~S\n",G2,world?()) ) ))

alternateMenAndWomen()  :: rule(
  F.currentPlace := p &  p > 1 
  => for G2 in {x in Guest | not(x.seated?)}
        (if (G2.male? = table1[p - 1].male?)
            (G2.canBeNext? := false, Nfiring :+ 1 )))

[create(n:integer)
  -> (
     Guest(idx = 1, male? = true, hobbies = 1),
     Guest(idx = 2, male? = false, hobbies = 2),
     Guest(idx = 3, male? = true, hobbies = 6),
     Guest(idx = 4, male? = false, hobbies = 3),
if (n > 4) (
     Guest(idx = 5, male? = true, hobbies = 7),
     Guest(idx = 6, male? = true, hobbies = 7),
     Guest(idx = 7, male? = false, hobbies = 7),
     Guest(idx = 8, male? = true, hobbies = 5),
     Guest(idx = 9, male? = true, hobbies = 7),
     Guest(idx = 10, male? = true, hobbies = 7),
     Guest(idx = 11, male? = false, hobbies = 7),
     Guest(idx = 12, male? = false, hobbies = 7),
     Guest(idx = 13, male? = false, hobbies = 6),
     Guest(idx = 14, male? = false, hobbies = 3),
     Guest(idx = 15, male? = false, hobbies = 7),
     Guest(idx = 16, male? = false, hobbies = 6) ),
  if (n > 16) (
     Guest(idx = 17, male? = false, hobbies = 6),
     Guest(idx = 18, male? = true, hobbies = 7),
     Guest(idx = 19, male? = false, hobbies = 5),
     Guest(idx = 20, male? = false, hobbies = 7),
     Guest(idx = 21, male? = true, hobbies = 6),
     Guest(idx = 22, male? = true, hobbies = 6),
     Guest(idx = 23, male? = false, hobbies = 3),
     Guest(idx = 24, male? = false, hobbies = 7),
     Guest(idx = 25, male? = false, hobbies = 7),
     Guest(idx = 26, male? = true, hobbies = 7),
     Guest(idx = 27, male? = false, hobbies = 7),
     Guest(idx = 28, male? = true, hobbies = 3),
     Guest(idx = 29, male? = false, hobbies = 7),
     Guest(idx = 30, male? = false, hobbies = 7),
     Guest(idx = 31, male? = true, hobbies = 7),
     Guest(idx = 32, male? = true, hobbies = 3) ),
  if (n > 32) (
     Guest(idx = 33, male? = true, hobbies = 7),
     Guest(idx = 34, male? = false, hobbies = 7),
     Guest(idx = 35, male? = false, hobbies = 6),
     Guest(idx = 36, male? = true, hobbies = 3),
     Guest(idx = 37, male? = true, hobbies = 3),
     Guest(idx = 38, male? = false, hobbies = 7),
     Guest(idx = 39, male? = true, hobbies = 7),
     Guest(idx = 40, male? = false, hobbies = 7),
     Guest(idx = 41, male? = true, hobbies = 7),
     Guest(idx = 42, male? = true, hobbies = 5),
     Guest(idx = 43, male? = true, hobbies = 7),
     Guest(idx = 44, male? = true, hobbies = 7),
     Guest(idx = 45, male? = true, hobbies = 3),
     Guest(idx = 46, male? = false, hobbies = 7),
     Guest(idx = 47, male? = true, hobbies = 3),
     Guest(idx = 48, male? = false, hobbies = 6),
     Guest(idx = 49, male? = true, hobbies = 6),
     Guest(idx = 50, male? = true, hobbies = 6),
     Guest(idx = 51, male? = false, hobbies = 7),
     Guest(idx = 52, male? = true, hobbies = 7),
     Guest(idx = 53, male? = true, hobbies = 3),
     Guest(idx = 54, male? = false, hobbies = 7),
     Guest(idx = 55, male? = false, hobbies = 7),
     Guest(idx = 56, male? = true, hobbies = 7),
     Guest(idx = 57, male? = false, hobbies = 7),
     Guest(idx = 58, male? = true, hobbies = 7),
     Guest(idx = 59, male? = true, hobbies = 7),
     Guest(idx = 60, male? = false, hobbies = 5),
     Guest(idx = 61, male? = false, hobbies = 6),
     Guest(idx = 62, male? = false, hobbies = 7),
     Guest(idx = 63, male? = false, hobbies = 7),
     Guest(idx = 64, male? = true, hobbies = 6) ),
  if (n > 64) (
     Guest(idx = 65, male? = true, hobbies = 6),
     Guest(idx = 66, male? = false, hobbies = 3),
     Guest(idx = 67, male? = true, hobbies = 5),
     Guest(idx = 68, male? = true, hobbies = 3),
     Guest(idx = 69, male? = false, hobbies = 5),
     Guest(idx = 70, male? = true, hobbies = 7),
     Guest(idx = 71, male? = true, hobbies = 3),
     Guest(idx = 72, male? = true, hobbies = 6),
     Guest(idx = 73, male? = true, hobbies = 6),
     Guest(idx = 74, male? = true, hobbies = 7),
     Guest(idx = 75, male? = false, hobbies = 7),
     Guest(idx = 76, male? = false, hobbies = 5),
     Guest(idx = 77, male? = false, hobbies = 7),
     Guest(idx = 78, male? = false, hobbies = 5),
     Guest(idx = 79, male? = false, hobbies = 7),
     Guest(idx = 80, male? = true, hobbies = 5),
     Guest(idx = 81, male? = false, hobbies = 6),
     Guest(idx = 82, male? = false, hobbies = 5),
     Guest(idx = 83, male? = true, hobbies = 6),
     Guest(idx = 84, male? = true, hobbies = 3),
     Guest(idx = 85, male? = true, hobbies = 6),
     Guest(idx = 86, male? = true, hobbies = 7),
     Guest(idx = 87, male? = true, hobbies = 5),
     Guest(idx = 88, male? = true, hobbies = 7),
     Guest(idx = 89, male? = true, hobbies = 7),
     Guest(idx = 90, male? = true, hobbies = 7),
     Guest(idx = 91, male? = false, hobbies = 5),
     Guest(idx = 92, male? = false, hobbies = 3),
     Guest(idx = 93, male? = false, hobbies = 3),
     Guest(idx = 94, male? = true, hobbies = 5),
     Guest(idx = 95, male? = false, hobbies = 6),
     Guest(idx = 96, male? = false, hobbies = 6),
     Guest(idx = 97, male? = false, hobbies = 6),
     Guest(idx = 98, male? = false, hobbies = 5),
     Guest(idx = 99, male? = false, hobbies = 7),
     Guest(idx = 100, male? = true, hobbies = 3),
     Guest(idx = 101, male? = false, hobbies = 7),
     Guest(idx = 102, male? = false, hobbies = 3),
     Guest(idx = 103, male? = true, hobbies = 3),
     Guest(idx = 104, male? = false, hobbies = 6),
     Guest(idx = 105, male? = false, hobbies = 3),
     Guest(idx = 106, male? = false, hobbies = 6),
     Guest(idx = 107, male? = false, hobbies = 3),
     Guest(idx = 108, male? = false, hobbies = 7),
     Guest(idx = 109, male? = true, hobbies = 5),
     Guest(idx = 110, male? = true, hobbies = 7),
     Guest(idx = 111, male? = true, hobbies = 3),
     Guest(idx = 112, male? = true, hobbies = 7),
     Guest(idx = 113, male? = true, hobbies = 5),
     Guest(idx = 114, male? = true, hobbies = 5),
     Guest(idx = 115, male? = false, hobbies = 3),
     Guest(idx = 116, male? = false, hobbies = 6),
     Guest(idx = 117, male? = false, hobbies = 6),
     Guest(idx = 118, male? = true, hobbies = 7),
     Guest(idx = 119, male? = false, hobbies = 7),
     Guest(idx = 120, male? = false, hobbies = 3),
     Guest(idx = 121, male? = false, hobbies = 7),
     Guest(idx = 122, male? = false, hobbies = 6),
     Guest(idx = 123, male? = false, hobbies = 3),
     Guest(idx = 124, male? = false, hobbies = 7),
     Guest(idx = 125, male? = false, hobbies = 3),
     Guest(idx = 126, male? = false, hobbies = 7),
     Guest(idx = 127, male? = false, hobbies = 6),
     Guest(idx = 128, male? = false, hobbies = 7) ) ) ]


[createBis()
 ->  for i in (1 .. random(N)) random(6),
     for i in (1 .. N)
       Guest(hobbies = randomDrawHobbies(),
             idx = i,
             male? = ((i mod 2) = 0)  ) ]

[randomDrawHobbies() : integer
   -> let i := random(4) in
      (case i ({0} 7,
               {1} 6,
               {2} 5,
               any 3) )]

// search loop
[solve(n:integer) : boolean
 -> ((n > N) |
     exists(G in Guest |
             G.canBeNext? &
             branch((//[DEBUG] try table[~A] = ~S // n, G,
                     table1[n] := G, G.seated? := true, G.canBeNext? := false,
                     // propagation
                     for G1 in {G1 in Guest | not(G1.seated?)}
                         G1.canBeNext? := true,
                     FXJ.currentPlace :+ 1,    // produce the event
                     solve(n + 1))) ))]

[showSol()
 -> printf("\n"),
    for j in (1 .. N)
      printf("~S", table1[j]),
    printf("\n"),
    for i in (0 .. 2)
      (for j in (1 .. N)
         printf("~A ",(if table1[j].hobbies[i] "X" else ".")),
       printf("\n") )]

[clean() : void
   -> for g in Guest (
         put(seated?,g,false),
         put(canBeNext?,g,true) ),
      erase(table1),
      FXJ.currentPlace := 1 ]


[go(n:integer) : void
    ->  create(N),
        let t:integer := 0 in (
            for i in (1 .. n) (
                clean(),
                time_set(),
                solve(1),
                t :+ time_get() ) ,
	   // showSol(),
            printf("~S rules fired in ~S ms\n",Nfiring,t) ) ]

[do_test() : void
  -> createBis(),
     time_set(),
     solve(1),
     time_show(),
     printf("~S rules fired \n",Nfiring) ]
     
[do_test1() : void
  -> go(5) ]


// quelques resultats sur Pentium Pro 200
// (ca permet de comparer avec les chiffres ILOG, eux aussi sur PPro200
// N = 32, #firings = 40, CPU = 15ms.
// N = 64, #firings = 181, CPU = 15ms.
// N = 128, #firings = 653, CPU = 15ms.
// N = 250, #firings = 2704, CPU = 47ms.
// N = 500, #firings = 9476, CPU = 156ms.
// N = 1000, #firings = 41926, CPU = 578ms.
// N = 2000, #firings = 164273, CPU = 1828ms.



// sur P II bi-pro 450 MHz, 128 Mo RAM  -  claire v2.3.13
// compile :
//   pour N = 16, 32, 64 ou 128 : CPU = 0  =>  non mesurable
// interprete :
//   N = 16  : #firings = 57,   CPU = 16ms,  ips = 3562 (n.s.)
//   N = 32  : #firings = 250,  CPU = 31ms,  ips = 8065
//   N = 64  : #firings = 1021, CPU = 140ms, ips = 7293
//   N = 128 : #firings = 4254, CPU = 531ms, ips = 8011

main() : void 
  -> (//[0] runs the dinner problem with N = // N,
      do_test(),
      exit(0))

