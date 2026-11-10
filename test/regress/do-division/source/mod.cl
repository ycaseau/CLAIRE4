Stored <: object(latest:integer = 0)
AFIELD:any := unknown

[getF() : integer -> (AFIELD as Stored).latest]
[probe() : integer -> getF() / 2]

main() : any -> (printf("probe = ~S\n", probe()), 0)
