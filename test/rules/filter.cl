// A simple rule test that was proposed by Talarian as benchmark for IPS (inference per second)

// number of frames
N :: (#if compiler.active? 1000000 else 1000)

TEST <: object(
       nc1:integer = 3,
       nc2:integer = 3,
       nc3:integer = 3,
       nc4:integer = 3,
       nc5:integer = 3,
       nc6:integer = 3,
       nc7:integer = 3,
       nc8:integer = 3,
       nc9:integer = 3,
       nc0:integer = 3,
       N1:integer = 0,
       N2:integer = 0,
       N3:integer = 0,
       N4:integer = 0,
       N5:integer = 0,
       N6:integer = 0,
       N7:integer = 0,
       N8:integer = 0,
       N9:integer = 0,
       N0:integer = 0)


a1() :: rule(
  N1(x) := y & y > 0 => (nc1(x) := (nc1(x) + 1)))

a2() :: rule(
  N2(x) := y & y > 0 => (nc2(x) := (nc2(x) + 1)))

a3() :: rule(
  N3(x) := y & y > 0 => (nc3(x) := (nc3(x) + 1)))

a4() :: rule(
  N4(x) := y & y > 0 => (nc4(x) := (nc4(x) + 1)))

a5() :: rule(
  N5(x) := y & y > 0 => (nc5(x) := (nc5(x) + 1)))

a6() :: rule(
  N6(x) := y & y > 0 => (nc6(x) := (nc6(x) + 1)))

a7() :: rule(
  N7(x) := y & y > 0 => (nc7(x) := (nc7(x) + 1)))

a8() :: rule(
  N8(x) := y & y > 0 => (nc8(x) := (nc8(x) + 1)))

a9() :: rule(
  N9(x) := y & y > 0 => (nc9(x) := (nc9(x) + 1)))

a0() :: rule(
  N0(x) := y & y > 0 => (nc0(x) := (nc0(x) + 1)))

// Talarian's test
do_test() : void
  -> (time_set(),
      for x in (1 .. N)
       let y := TEST() in
        (N1(y) := x,
         N2(y) := x,
         N3(y) := x,
         N4(y) := x,
         N5(y) := x,
         N6(y) := x,
         N7(y) := x,
         N8(y) := x,
         N9(y) := x,
         N0(y) := x),
      time_show())

// instantiation
do_test0() : void
 -> (time_set(),
     for x in (1 .. N)
       let y := TEST() in y,
	 time_show())
	 

// performance test
// N = 1000000 -> 500ms => 20 MIPS (million inference per second)
// CLAIRE4: (delta = 40ms) => 250 MIPS
main() : void ->
  (printf("Talarian filter test with ~S frames\n",N),
   printf("without rules\n"),
   do_test0(),
   printf("with rules:\n"),
   do_test(),
   exit(0))

