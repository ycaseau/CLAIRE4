(printf("Hello CLAIRE4 on Linux, this is our init.cl file\n"))

// Ubuntu version
(osname() := "linux", *fs* := "/")
*root* :: "/home/ubuntu/claire"
*where* :: (*root* / "go")                                                // where the init file is
*output* :: (*root* / "go/src")
*meta* :: "~/claire/src/meta"                                          // source files from Github
*compile* :: "~/claire/src/compile"   
*bsrc* :: (*root* / "test/nonreg")
*tsrc* :: (*root* / "test/perf")
*rsrc* :: (*root* / "test/rules")

// these are the global variables expected by the compiler
RELEASE:float :: 1.4    // version of January 2025 

// ***************************************************************************
// *    Part 1: Modules & compiler environment                               *
// ***************************************************************************

// meta files are now the "official" github directory
(for m in {Core,Language,Reader} source(m) := *meta*,
 for m in {Optimize,Generate} source(m) := *compile*)

// where we want to generate the go code
(when c := get_value("compiler") in 
   (c.safety := 5,
    verbose() := 1,
    source(c) := *output*))


// ***************************************************************************
// *    Part 2: Performance test modules                                     *
// ***************************************************************************

// these are the performance test files of 2020
mFib :: module(part_of = claire,
              source = *tsrc*,
              uses = list(Reader,mClaire),
              made_of = list("testFib"))

mList :: module(part_of = claire,
              source = *tsrc*,
              uses = list(Reader,mClaire),
              made_of = list("testList"))

mSet :: module(part_of = claire,
              source = *tsrc*,
              uses = list(Reader,mClaire),
              made_of = list("testSet"))

mDict :: module(part_of = claire,
              source = *tsrc*,
              uses = list(Reader,mClaire),
              made_of = list("testDict"))

mObj :: module(part_of = claire,
              source = *tsrc*,
              uses = list(Reader,mClaire),
              made_of = list("testObj"))

mSend :: module(part_of = claire,
              source = *tsrc*,
              uses = list(Reader,mClaire),
              made_of = list("testSend"))

mCopy :: module(part_of = claire,
              source = *tsrc*,
              uses = list(Reader,mClaire),
              made_of = list("testCopy"))

// ***************************************************************************
// *    Part 3: Bugs for CLAIRE                                              *
// ***************************************************************************

// bu0 is a file of bugs that used to crash nastily but now should trigger a panic error
bu0 :: module( uses = list(Reader), source = *bsrc*,
               made_of = list("bug0"))

// parsing bugs: things that cannot get read right
bu1 :: module( uses = list(Reader), source = *bsrc*,
               made_of = list("bstub","bug1"))

// array related bugs
bu2 :: module( uses = list(Reader), source = *bsrc*, 
               made_of = list("bstub", "bug2"))

// table related bugs
bu3 :: module( uses = list(Reader), source = *bsrc*,
               made_of = list("bstub", "bug3"))

// iteration of a union (interpreted) and other patterns
bu4 :: module( uses = list(Reader), source = *bsrc*,
               made_of = list("bstub", "bug4"))

// bugs with floats
bu5 :: module( uses = list(Reader), source = *bsrc*,
               made_of = list("bstub", "bug5"))

// bugs with class & method definitions
bu6 :: module( uses = list(Reader), source = *bsrc*,
               made_of = list("bstub", "bug6"))

// bugs with worlds
bu7 :: module( uses = list(Reader), source = *bsrc*,
               made_of = list("bstub", "bug7"))

// bugs with instantiation & primitive types
bu8 :: module( uses = list(Reader), source = *bsrc*,
               made_of = list("bstub", "bug8"))

// reversible cells from CLP, untyped version
bu9 :: module( uses = list(Reader), source = *bsrc*,
               made_of = list("bstub", "bug9"))

// reversible cells from CLP, typed version
bu10 :: module( uses = list(Reader), source = *bsrc*,
               made_of = list("bstub", "bug10"))

// famous examples (stack example, doc examples ...)
bu11 :: module( uses = list(Reader), source = *bsrc*,
                made_of = list("bstub", "bug11"))

// bug with tuples, lists and sets
bu12 :: module( uses = list(Reader), source = *bsrc*,
                made_of = list("bstub","bug12"))

// test file for handling unknown & inverses
bu13 :: module( uses = list(Reader), source = *bsrc*,
                made_of = list("bstub", "bug13"))

// test file for compilation bug (works with claire3)
bu14 :: module( uses = list(Reader), source = *bsrc*,
                made_of = list("bstub", "bug14"))

// another test file for testing compiler
bu15 :: module( uses = list(Reader), source = *bsrc*,
                made_of = list("bstub", "bug15"))

// sudoku example : shown in the tutorial - good example of rules & branch
bu16 :: module( uses = list(Reader), source = *bsrc*,
                made_of = list("bstub", "sudoku"))

// other examples from rge claire manual : need to put in the doc - good example of rules & branch
bu17 :: module( uses = list(Reader), source = *bsrc*,
                made_of = list("bstub", "manual"))



// ***************************************************************************
// *    Part 4: Simple rule examples                                              *
// ***************************************************************************

// dinner
mDinner :: module( uses = list(Reader), source = *rsrc*,
                    made_of = list("dinner"))


// filter
mFilter :: module( uses = list(Reader), source = *rsrc*,
                   made_of = list("filter"))

// monkey
mMonkey :: module( uses = list(Reader), source = *rsrc*,
                   made_of = list("monkey"))

// zebra
mZebra :: module( uses = list(Reader), source = *rsrc*,
                  made_of = list("zebra"))

// airline - WIP (old CLAIRE 2 example)
mAirline :: module( uses = list(Reader), source = *rsrc*,
                  made_of = list("Airline"))

// these are the old non-regression tests files (refreshed in July 2021)
(printf("Done. \n"))
