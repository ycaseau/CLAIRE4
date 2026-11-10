// A method whose body is typed but not exact ends with RangeCheck, which can
// raise [17]. It must therefore be marked exception-ready (EID), otherwise its
// callers generate a native function that consumes the EID as a value:
//     F_getPb_void() + 1   (mismatched types Kernel.EID and untyped int)
RangecheckStatus :: module(source = "source", uses = list(Core, Reader), made_of = list("mod"))
(load(RangecheckStatus))
