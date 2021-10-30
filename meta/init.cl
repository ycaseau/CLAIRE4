(printf("----------------- prototype of CLAIRE v2.9 compiler --------------------\n"),
 source(compiler) := "test",
 safety(compiler) := 5)

(debug())

RELEASE :: 9.0

gen :: module(source = "d:\\claire\\v2.9\\src\\compile",
              part_of = Compile,
              made_of = list("gsystem","gexp","gstat","cgen"))

;Core :: module(source = "d:\\claire\\v2.9\\src\\meta",
;               part_of = claire,
;               made_of = list("method","core","function","type"))
(source(System) := "d:\\claire\\v2.9\\src\\meta",
 source(Language) := "d:\\claire\\v2.9\\src\\meta",
 source(Reader) := "d:\\claire\\v2.9\\src\\meta",
 made_of(System) := list("method","object","function","types"))

[go() -> 
  //[0] ... generate(System)  //,
  generate(System) ]

[go1() -> 
  //[0] ... generate(Language + Reader)  //,
  generate(Language),
  generate(Reader) ]