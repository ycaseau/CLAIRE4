// OPEN. Dividing the result of a function that may throw generates
//     cannot use try_n (Kernel.EID) as int value in assignment
// because "probe" is generated as a native function although it may raise.
DoDivision :: module(source = "source", uses = list(Core, Reader), made_of = list("mod"))
(load(DoDivision))
