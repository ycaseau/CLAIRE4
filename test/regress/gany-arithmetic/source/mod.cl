GINT:integer := 41
GANY:any := 41

[useInt() : integer -> GINT + 1]     // always worked: 42
[useAny() : integer -> GANY + 1]     // used to give 1

// the shape found in choco (CURRENT_PB:any := unknown)
[getAny() : integer -> GANY]
[indirect() : integer -> getAny() + 1]

[check() : any ->
   printf("useInt     = ~S (expected 42)\n", useInt()),
   printf("useAny     = ~S (expected 42)\n", useAny()),
   printf("indirect   = ~S (expected 42)\n", indirect())]
main() : any -> (check(), 0)
