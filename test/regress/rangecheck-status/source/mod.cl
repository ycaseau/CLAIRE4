GLOBAL_PB:any := 41

[getPb() : integer -> GLOBAL_PB]
[useIt() : integer -> getPb() + 1]

[check() : any ->
   printf("getPb  = ~S (expected 41)\n", getPb()),
   printf("useIt  = ~S (expected 42)\n", useIt())]
main() : any -> (check(), 0)
