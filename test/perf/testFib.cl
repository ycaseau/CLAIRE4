// fib test - the oldest claire file :)

fib(n:integer) : integer -> (if (n < 2) n else (fib(n - 1) + fib(n - 2)))

g(n:integer) : void
  -> (time_set(), fib(n), time_show())

(printf("Done. \n"))

[fab(n:float) : float -> if (n < 2.0) 1.0001 else fab(n - 1.0) + fab(n - 2.0)]

[tfab(n:float) -> time_set(), fab(n), time_show()]


all() ->
  (princ("fib(30)"),
   g(30),
   princ("fib(35)"),
   g(35),
   princ("fib(35.0)"),
   tfab(35.0))

[main() -> 
  //[0] test fib ----- //,
  all(),
  exit(0) ]

