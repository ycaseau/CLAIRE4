(printf("Hello World, this is CLAIRE 4\n"))

mFib :: module(part_of = claire,
              source = "test",
              uses = list(Reader,mClaire),
              made_of = list("testFib"))

fib(n:integer) : integer -> (if (n < 2) 1 else (fib(n - 1) + fib(n - 2)))

g(n:integer) : any
  -> (time_set(), fib(n), time_show())

(printf("Done. \n"))

