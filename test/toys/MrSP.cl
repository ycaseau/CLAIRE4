// Mr S+P Puzzle
// Mr S and Mr P are two persons.
// Mr S nows the sum of two numbers A + B, Mr P known the product A x B, A & B are integers greater that 1.
// When they meet they have the following dialog:
//   (1) Mr S says "the sum is less than 100"
//   (2) Mr P says "I cannot conclude"
//   (3) Mr S says "I knew you could not conclude"
//   (4) Mr P says "then I known (A and B)"
//   (5) Mr S says "then I know them two"
// what are A and B ?

// sets of A (A < B) from S (sum)
[AfromS(S:integer) 
   -> (2 .. (S / 2))]

// sets of A (A < B) from P (product)
[AfromP(S:integer) 
   -> {i in (2 .. integer!(sqrt(float!(S)))) | S mod i = 0}]

// first statement of Mr S
[statement1(S:integer) : boolean
   -> S <= 100]

// second statement of Mr P
[statement2(P:integer) : boolean
   -> size({A in AfromP(P) | A + (P / A) <= 100 }) > 1 ]

// third statement of Mr S : all decomposition of S satisfy statement2
[statement3(S:integer) : boolean
   -> forall(A in AfromS(S) | statement2(A * (S - A))) ]

// fourth statement of Mr P : only one decomposition of P satisfy statement3
[statement4(P:integer) : boolean
   -> size({A in AfromP(P) | statement1(A + P / A) & statement3(A + P / A)}) = 1]

// fifth statement of Mr S : only one decomposition of S satisfy statement4
[statement5(S:integer) : boolean
   -> size({A in AfromS(S) | statement3(S) & statement4(A * (S - A))}) = 1]

// solve the puzzle and show there is only one solution
[solve() : void
   -> for A in (2 .. 99)
       for B in (A .. (100 - A))       // iterate so that A < B
         let S := A + B, P := A * B in
           (if (statement2(P) & statement3(S) & statement4(P) & statement5(S))
            printf("found a solution : ~S and ~S\n", A, B))]

// debug
[check(A:integer,B:integer) 
  -> let S := A + B, P := A * B,
         lS := list{list(A1,S - A1) | A1 in AfromS(S)},
         lP := list{list(A1,P / A1) | A1 in AfromP(P)} in
        (printf("list of S decomposition:~S\n",lS),
         printf("list of P decomposition:~S\n",lP),
         printf("lS that satisfy S2: ~S\n",list{l1 in lS | statement2(l1[1] * l1[2])}),
         printf("lP that satisfy S3: ~S\n",list{l1 in lP | statement3(l1[1] + l1[2])}))]

// number of S decomposition that match P property

// variant of the puzzle
[statement0b(P:integer) : boolean
   -> P >= 30]

// first statement of Mr S
[statement1b(S:integer) : boolean
   -> S <= 100]

// second statement of Mr P
[statement2b(P:integer) : boolean
   -> size({A in AfromP(P) | (A + (P / A) <= 100) }) > 1 ]

// third statement of Mr S : all decomposition of S satisfy statement2
[statement3b(S:integer) : boolean
   -> forall(A in AfromS(S) | statement2b(A * (S - A))) ]

// fourth statement of Mr P : only one decomposition of P satisfy statement3
[statement4b(P:integer) : boolean
   -> size({A in AfromP(P) | statement1b(A + P / A) & statement3b(A + P / A)}) = 1  &  P >= 130 ]

// fifth statement of Mr S : only one decomposition of S satisfy statement4
[statement5b(S:integer) : boolean
   -> size({A in AfromS(S) |  statement3b(S) & statement4b(A * (S - A))}) = 1]

[count5b(S:integer) : integer
   -> size({A in AfromS(S) |  statement3b(S) & statement4b(A * (S - A))}) ]



[solve2() : void
   -> for A in (2 .. 99)
       for B in (A .. (100 - A))       // iterate so that A <= B
         let S := A + B, P := A * B in
           (if (statement2b(P) & statement3b(S) & statement4b(P) & statement5b(S))
            printf("found a solution : ~S and ~S -> ~A\n", A, B, count5b(S)))]


// ==================== third variant of the puzzle ====================

/* Mr S and Mr P are two persons. 
 Mr S nows the sum of two numbers A + B, Mr P known the product A x B, A & B are integers greater that 1. 
 When they meet they have the following dialog: 
 (1) Mr S says "the sum is less than 100" 
 (2) Mr P says "I cannot conclude" 
 (3) Mr S says "I knew you could not conclude" 
 (4) Mr P says "I still cannot conclude" 
(5) Mr S says "then I have found A and B" 
what are the possible values for A and B so that the previous dialog stands ? */

// first statement of Mr S
[statement1c(S:integer) : boolean
   -> S <= 100]

// second statement of Mr P
[statement2c(P:integer) : boolean
   -> size({A in AfromP(P) | A + (P / A) <= 100 }) > 1 ]

// third statement of Mr S : all decomposition of S satisfy statement2
[statement3c(S:integer) : boolean
   -> forall(A in AfromS(S) | statement2c(A * (S - A))) ]

// fourth statement of Mr P : more than onedecomposition of P satisfy statement3
[statement4c(P:integer) : boolean
   -> size({A in AfromP(P) | statement1c(A + P / A) & statement3c(A + P / A)}) > 1]

// fifth statement of Mr S : only one decomposition of S satisfy statement4
[statement5c(S:integer) : boolean
   -> size({A in AfromS(S) | statement3c(S) & statement4c(A * (S - A))}) = 1]


// solve the puzzle and show there is only one solution
[solve3() : void
   -> for A in (2 .. 99)
       for B in (A .. (100 - A))       // iterate so that A < B
         let S := A + B, P := A * B in
           (if (statement2c(P) & statement3c(S) & statement4c(P) & statement5c(S))
            printf("found a solution : ~S and ~S\n", A, B))]
