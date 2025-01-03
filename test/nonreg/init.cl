(printf("Hello World, this is our init.cl file\n"))

// play with debugger
foo(n:integer) 
  -> (if (n = 0)  2 / n
      else foo(n - 1))

// ***************************************************************************
// *    Part 1: Performance tests                                            *
// ***************************************************************************

// these are the performance test files of 2020
mFib :: module(part_of = claire,
              source = "test",
              uses = list(Reader,mClaire),
              made_of = list("testFib"))

mList :: module(part_of = claire,
              source = "test",
              uses = list(Reader,mClaire),
              made_of = list("testList"))

mSet :: module(part_of = claire,
              source = "test",
              uses = list(Reader,mClaire),
              made_of = list("testSet"))

mDict :: module(part_of = claire,
              source = "test",
              uses = list(Reader,mClaire),
              made_of = list("testDict"))

mObj :: module(part_of = claire,
              source = "test",
              uses = list(Reader,mClaire),
              made_of = list("testObj"))

mSend :: module(part_of = claire,
              source = "test",
              uses = list(Reader,mClaire),
              made_of = list("testSend"))

mCopy :: module(part_of = claire,
              source = "test",
              uses = list(Reader,mClaire),
              made_of = list("testCopy"))

// ***************************************************************************
// *    Part 2: Bugs for CLAIRE                                              *
// ***************************************************************************

*where* :: "/Users/ycaseau/claire/v4.0/go"

// parsing bugs: things that cannot get read right
bu1 :: module( uses = list(Reader), source = *where* / "bsrc",
               made_of = list("bstub","bug1"))

// array related bugs
bu2 :: module( uses = list(Reader), source = *where* / "bsrc", 
               made_of = list("bstub", "bug2"))

// table related bugs
bu3 :: module( uses = list(Reader), source = *where* / "bsrc",
               made_of = list("bstub", "bug3"))

// iteration of a union (interpreted) and other patterns
bu4 :: module( uses = list(Reader), source = *where* / "bsrc",
               made_of = list("bstub", "bug4"))

// bugs with floats
bu5 :: module( uses = list(Reader), source = *where* / "bsrc",
               made_of = list("bstub", "bug5"))

// bugs with class & method definitions
bu6 :: module( uses = list(Reader), source = *where* / "bsrc",
               made_of = list("bstub", "bug6"))

// bugs with worlds
bu7 :: module( uses = list(Reader), source = *where* / "bsrc",
               made_of = list("bstub", "bug7"))

// bugs with instantiation & primitive types
bu8 :: module( uses = list(Reader), source = *where* / "bsrc",
               made_of = list("bstub", "bug8"))

// reversible cells from CLP, untyped version
bu9 :: module( uses = list(Reader), source = *where* / "bsrc",
               made_of = list("bstub", "bug9"))

// reversible cells from CLP, typed version
bu10 :: module( uses = list(Reader), source = *where* / "bsrc",
               made_of = list("bstub", "bug10"))

// famous examples (stack example, doc examples ...)
bu11 :: module( uses = list(Reader), source = *where* / "bsrc",
                made_of = list("bstub", "bug11"))

// bug with tuples, lists and sets
bu12 :: module( uses = list(Reader), source = *where* / "bsrc",
                made_of = list("bstub","bug12"))

// test file for handling unknown & inverses
bu13 :: module( uses = list(Reader), source = *where* / "bsrc",
                made_of = list("bstub", "bug13"))

// sudoku example
bu14 :: module( uses = list(Reader), source = *where* / "bsrc",
                made_of = list("bstub", "sudoku"))

// compilation bugs are not tested with claire1

// ***************************************************************************
// *    Part 3: Simple rule examples                                              *
// ***************************************************************************

// these are the old non-regression tests files (refreshed in July 2021)
(printf("Done. \n"))

