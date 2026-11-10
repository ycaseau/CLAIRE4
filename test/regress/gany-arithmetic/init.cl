// A global declared ":any" must be converted to a well formed EID when read.
// With the bug, "GANY + 1" returned 1 instead of 42 (CLAIRE3 returns 42).
GanyArithmetic :: module(source = "source", uses = list(Core, Reader), made_of = list("mod"))
(load(GanyArithmetic))
