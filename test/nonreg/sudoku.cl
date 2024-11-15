// sudoku example (to include in CLAIRE doc)

// data structure
Cell <: object
CellSet <: object

// cell from a 9 x 9 Sudoku grid
Cell[x,y] <: object(x:integer,
                    y:integer,
                    possible:list<boolean>,  // list of possible digits for the cell
                    count:integer = 9,       // number of possible digit
                    digit:integer = 0,       // assigned value to the cell (0 = none)
                    line:CellSet,            // each cell belongs to 3 CellSets: line,
                    column:CellSet,          // its column
                    square:CellSet)          // its 3x3 square

// a set of cells: line, column, square that holds the AllDiff constraint
CellSet[contents] <: object(contents:list<Cell>,          // cells that belong to the set
                            counts:list<integer>)         // a possible digit counter

// two defeasible slots for hypothetical reasoning, but possible uses direct store
store(digit,count)

// creates a cell
makeCell(a:integer,b:integer) : Cell
  -> Cell(x = a, y = b, possible = list<boolean>{true | i in (1 .. 9)}, digit = 0)

// A sudoku grid
Grid <: object(cells:list<Cell>,
               lines:list<CellSet>,
               columns:list<CellSet>,
               squares:list<CellSet>)


// useful for debug
nth(g:Grid,i:(1 .. 9),j:(1 .. 9)) : Cell
  -> some(c in g.cells | c.x = i & c.y = j)

// creates a grid
makeGrid() : Grid
  -> let g := Grid() in
        (for i in (1 .. 9)
           for j in (1 .. 9) g.cells :add makeCell(i,j),
         for i in (1 .. 9)
            let li := list<Cell>{c in g.cells | c.x = i},
                cs := CellSet(contents = li, counts = list<integer>{9 | i in (1 .. 9)}) in
               (g.lines :add cs,
                for c in li c.line := cs),
         for j in (1 .. 9)
            let co := list<Cell>{c in g.cells | c.y = j},
                cs := CellSet(contents = co, counts = list<integer>{9 | i in (1 .. 9)}) in
               (g.columns :add cs,
                for c in co c.column := cs),
         for k1 in (1 .. 3)
           for k2 in (1 .. 3)
              let sq := list<Cell>{c in g.cells | abs(3 * k1 - c.x - 1) <= 1 & abs(3 * k2 - c.y - 1) <= 1},
                  cs := CellSet(contents = sq, counts = list<integer>{9 | i in (1 .. 9)}) in
               (g.squares :add cs,
                for c in sq c.square := cs),
         g)


// This program uses two propagation rules
// the first one propagates a choice by forbidding the digit in the line, column, square
r1() :: rule(
     c.digit := v => (//[1] propagation of ~S <- ~A // c,v,
                      store(c.line.counts,v,0),         // disable counts[v] since v was found !
                      store(c.column.counts,v,0),
                      store(c.square.counts,v,0),
                      for v2 in (1 .. 9)
                        (if (v != v2 & c.possible[v2])
                            (store(c.possible,v2,false),       // avoid double count
                             oneLess(c.line,v2),
                             oneLess(c.column,v2),
                             oneLess(c.square,v2))),
                      for c2 in (c.line.contents but c) forbid(c2,v),
                      for c2 in (c.column.contents but c) forbid(c2,v),
                      for c2 in (c.square.contents but c) forbid(c2,v)))

// forbid a digit
// Attention: if we forbid a digit that is assigned, we must raise a contradiction
forbid(c:Cell,v:(1 .. 9))
  -> (//[3] forbid ~S(~A) -> ~A // c,c.digit,v,
      if (c.x = 5 & c.y = 9) trace(1,">>>> forbid ~S(~A) -> ~A [~A]\n",c,c.digit,v,c.count),
      if (c.digit = v) (//[5] contradiction while propagation //,
                        contradiction!())
      else if (c.digit = 0 & c.possible[v])
         (store(c.possible,v,false),
          c.count :- 1,
          oneLess(c.line,v),
          oneLess(c.column,v),
          oneLess(c.square,v)))

// debug
CS:any :: 1

// remove a digit v in a CellSet cs
oneLess(cs:CellSet,v:(1 .. 9)) : void
  -> let cpos := cs.counts[v] in
      (if (cpos > 0)
         (store(cs.counts,v,cpos - 1),
          if (cpos <= 2)
            when c := some(c in cs.contents | c.digit = 0 & c.possible[v]) in
              (//[1] CellSet inference ~S (~A) -> ~S = ~A // cs, list{c.digit | c in cs.contents},c,v,
               CS := cs,
               c.digit := v)
            else (//[1] contradiction because ~S has no ~A // cs,v,
                  contradiction!())))

// second rule says that once only one digit is possible, we should deduce it !
r2() :: rule(
     c.count := y & y = 1 =>  (trace(1,"--- r2 finds a digit for ~S:~A\n ",c,some(y in (1 .. 9) | c.possible[y])),
                              c.digit := some(y in (1 .. 9) | c.possible[y])))


// create a grid from a problem
[grid(l1:list[list[integer]]) : Grid
   -> let g := makeGrid() in
        (assert(length(l1) = 9),
         for c in g.cells
           let i := c.x, j := c.y, val := l1[i][j] in
             (if (val != 0)  c.digit := val),
          g) ]

// classical heuristic for branching : finds a cell with a min count
[findPivot(g:Grid) : any
  -> let minv := 10, cmin := unknown in
        (for c in g.cells
           (if (c.digit = 0 & c.count < minv)
               (minv := c.count, cmin := c)),
         cmin) ]

NBK:integer :: 0

// solve a sudoku : branch on possible digits
[solve(n:integer,g:Grid) : boolean
  -> when c := findPivot(g) in
        exists(v in (1 .. 9) |
               (if c.possible[v] branch((trace(1,"[~A] create a branch ~S <- ~A\n",n,c,v),
                                         c.digit := v,
                                         let b := solve(n + 1, g) in
                                           (if not(b) trace(1,"[~A] ~S := ~A failed\n",n,c,v),
                                            if not(b) NBK :+ 1,
                                            b)))
                else false))
     else true]

// show the solution
[see(g:Grid)
  -> printf("\n\t------------------\n"),
     for i in (1 .. 9) printf("\t~I\n",(for j in (1 .. 9) printf("~A ",g[i,j].digit))) ]

// count the number of solutions
[count(g:Grid) : integer
   -> when c := findPivot(g) in
        let ct := 0 in
           (for v in (1 .. 9)
               (if c.possible[v] branch((trace(1,"create a branch ~S <- ~A\n",c,v),
                                         c.digit := v, ct :+ count(g), false))),
            ct)
     else (see(g), check(g), 1) ]

// example from Yvette
S1:any :: 1

foo()
 -> (S1 := grid(list(list(0,3,0,0,9,0,0,1,0),
                list(0,0,7,0,0,0,0,0,6),
                list(0,0,0,0,3,4,0,0,7),
                list(0,0,0,0,0,0,0,0,3),
                list(8,2,1,0,5,0,4,7,9),
                list(9,0,0,0,0,0,0,0,0),
                list(4,0,0,5,2,0,0,0,0),
                list(3,0,0,0,0,0,2,0,0),
                list(0,6,0,0,4,0,0,5,0))))

// another example                
foo2()
 -> (S1 := grid(list(list(0,6,0,0,0,2,0,8,4),
                list(2,3,0,4,0,0,0,6,0),
                list(0,8,4,0,0,0,3,0,0),
                list(0,0,0,6,4,0,0,0,5),
                list(8,0,1,7,0,0,6,0,0),
                list(0,0,0,0,3,0,0,2,0),
                list(0,0,0,2,6,0,0,0,0),
                list(0,0,0,0,0,0,0,1,3),
                list(0,0,2,8,0,0,0,0,0))))

go() -> (verbose() := 0, solve(1,S1))

foo3()
 -> (S1 := grid(list(list(0,0,3,4,9,2,6,8,0),
                list(0,1,0,0,6,0,0,0,2),
                list(2,4,0,0,8,5,7,3,9),
                list(1,5,4,0,3,7,0,0,0),
                list(0,2,7,0,4,9,0,1,0),
                list(3,8,9,0,1,6,2,0,0),
                list(4,0,5,3,7,1,8,0,0),
                list(0,6,2,9,5,0,1,0,0),
                list(7,3,0,6,0,8,5,0,0))))

// tests
claire/testS1() : integer -> (verbose() := 0, foo(), solve(1,S1), see(S1), NBK)
claire/testS2() : integer -> (verbose() := 1, foo2(), solve(1,S1), see(S1), NBK)
claire/testS3() : integer -> (verbose() := 0, foo2(), solve(1,S1), see(S1), NBK)
claire/testS4() : integer -> (verbose() := 0, time_set(),foo4(), solve(1,S1), time_show(),see(S1), NBK)
claire/countS1() : integer -> (verbose() := 0, foo(), count(S1))
claire/countS2() : integer -> (verbose() := 0, foo2(), count(S1))
claire/countS3() : integer -> (verbose() := 0, foo3(), count(S1))


// debug
claire/bar1() -> (verbose() := 1, foo2())
claire/bar2() -> (verbose() := 1, foo2(), S1[1,3].digit := 5)


// check the validity of a solution
[check(g:Grid) 
  -> forall(cs in g.lines | (check("line",cs), true)),
     forall(cs in g.columns | (check("column",cs), true)),
     forall(cs in g.squares | (check("square",cs), true)) ]

[check(s:string, cs:CellSet) 
  -> for i in (1 .. 9)
      (if not(exists(c in cs.contents | c.digit = i))
           error("~A :cell set ~S has no ~A",s,cs,i)) ]

// the proper test is to count .. since it implies a check (solution is good) and exploration
// (check("count sols", (countS1() = 1)))
// (record("test with #backtracks = ",NBK))

(testOK())

// S5: hard sudoku from sudoku.com
foo4()
  -> (S1 :: grid(list(list(0,2,0,0,0,4,0,0,1),
                      list(8,0,0,0,0,0,0,0,7),
                      list(0,0,7,0,5,3,0,6,0),
                      list(0,0,5,4,0,0,0,0,0),
                      list(0,0,0,8,0,0,0,2,0),
                      list(9,0,0,0,2,5,3,0,0),
                      list(0,1,0,0,0,0,9,0,0),
                      list(0,0,0,2,0,0,0,0,0),
                      list(0,0,8,0,6,7,0,3,0))))
                

// example from Yvette (diabolique + hard to solve)
foo5()
  -> (S1 :: grid(list(list(6,0,0,0,5,4,3,0,0),
                      list(9,0,0,0,8,0,0,0,0),
                      list(2,0,0,0,0,0,6,9,0),
                      list(5,2,0,0,0,0,0,0,0),
                      list(0,3,0,0,7,0,0,0,0),
                      list(0,0,0,0,0,0,1,0,8),
                      list(0,0,4,0,0,9,0,7,5),
                      list(0,0,5,1,0,0,0,0,0),
                      list(0,0,0,6,0,0,0,3,0))))


/*
// median example from iPhone
S4 :: grid(list(list(0,0,0,3,0,0,0,0,0),
                list(5,0,7,0,0,0,9,8,6),
                list(0,8,0,0,0,0,0,0,0),
                list(0,0,0,0,0,0,0,0,0),
                list(4,0,0,0,7,2,0,5,0),
                list(1,0,3,0,6,5,0,0,0),
                list(6,4,0,2,3,9,0,0,1),
                list(0,0,1,7,5,4,0,0,9),
                list(0,0,5,0,0,8,2,4,0)))
*/


