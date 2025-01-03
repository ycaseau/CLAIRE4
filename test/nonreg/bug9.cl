// +------------------------------------------------------------+
// | bug9.cl                                                    |
// | last update: Jan 2000 - Y. Caseau                          |
// +------------------------------------------------------------+


// ---------------------------------------------------------------
// this file contains CLP's reversible cell in the untyped version
// ---------------------------------------------------------------

// ****************************************************************************
// * Part 0: Forward declarations
// ****************************************************************************

[reversible_cell <: object]
[reversible_list <: collection]


// ****************************************************************************
// * Part 1: Classes
// ****************************************************************************

[reversible_cell
  <: object(_content:any,
		        _next_cell:reversible_cell)
]
store(_next_cell)

[reversible_list
  <: collection(_first_cell:reversible_cell)
]
store(_first_cell)
ephemeral(reversible_list)

// ****************************************************************************
// * Part 2: Methods
// ****************************************************************************

[add(rl:reversible_list, content:any, queue:boolean) : void
  -> (if (not(known?(_first_cell, rl)))
        _first_cell(rl) := reversible_cell(_content = content)
      else if (queue)
        let rc := _first_cell(rl)
        in (while (known?(_next_cell, rc))
              rc := _next_cell(rc),
            _next_cell(rc) := reversible_cell(_content = content))
      else
        _first_cell(rl) := reversible_cell(_content = content,
					   _next_cell = _first_cell(rl)))
]

[set!(rl:reversible_list) : set
  -> let rc := get(_first_cell, rl),
	 s := set<any>()
     in (while (known?(rc))
           (s :add _content(rc),
	    rc := get(_next_cell, rc)),
	 s)
]

[iterate(rl:reversible_list, v:Variable, e:any)
  => let v:any := 0, rc := get(_first_cell, rl)
     in while (known?(rc))
       (v := _content(rc),
	e,
	rc := get(_next_cell, rc))
]

[test()
  -> let l1:reversible_list := reversible_list(),
	 l2:reversible_list := reversible_list()
     in (add(l1, 1, true), add(l1, 2, true), add(l1, 3, true),
	 add(l2, 4, true), add(l2, 5, true), add(l2, 6, true),
	 printf("\n"),
         for i in set!(l1) print(i),
	 printf("\n"),
	 for i in l1 print(i),
	 printf("\n"),
         for i in set!(l1) (for j in set!(l2) (print(i), print(j))),
	 printf("\n"),
	 for i in l1 (for j in l2 (print(i), print(j))),
	 printf("\n"))
]

(test())

(testOK())
