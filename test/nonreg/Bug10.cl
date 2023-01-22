// +------------------------------------------------------------+
// | bug10.cl                                                   |
// | last update: Jan 2000 - Y. Caseau                          |
// +------------------------------------------------------------+


// ---------------------------------------------------------------
// this file contains CLP's reversible cell in the typed version
// ---------------------------------------------------------------

// ---------------------------------------------------------------------------
// CLAIRE SCHEDULE                                              Claude Le Pape
// revlist.cl
//                                                          Copyright (C) 1996
// ---------------------------------------------------------------------------

// ***************************************************************************
// * Part 0: Forward declarations
// ***************************************************************************
// ephemeral_object <: object()

[private/reversible_cell <: ephemeral_object]
[private/reversible_list <: ephemeral_object]

// ***************************************************************************
// * Part 1: Classes
// ***************************************************************************

[private/reversible_cell <: ephemeral_object(
        private/_content:any,
        private/_next_cell:(reversible_cell U {nil}) = nil)
]
store(_next_cell)

// compiler pragma
// (known!(_content))

[reversible_list[of] <: ephemeral_object(
        private/of:type,
        private/_first_cell:(reversible_cell U {nil}) = nil)
]
store(_first_cell)

[activity <: ephemeral_object()]

[schedule <: ephemeral_object(
        private/_activities:reversible_list[of = activity])
]

// ***************************************************************************
// * Part 2: Methods
// ***************************************************************************

[get_activities(sch:schedule) : reversible_list[of = activity]
  => _activities(sch)
]

[car(rl:reversible_list[of = X]) : type[X]
  -> assert(_first_cell(rl)),
     _content(_first_cell(rl))
]

[test(sched:schedule) : activity
  -> car(get_activities(sched))
]

AA :: schedule()


// -------------- a problem from thomson ---------------------------

//
[problem <: ephemeral_object]
[order <: ephemeral_object]
[machine <: ephemeral_object]

// ***************************************************************************
// * Part 1: Classes
// ***************************************************************************

[problem <: ephemeral_object(_orders:list[order],
                             _machines:list[machine])
]

[order <: ephemeral_object(_order_number:integer = 0,
                           _bill:list[machine])
]

[machine <: ephemeral_object()
]

// ***************************************************************************
// * Part 2: Constructors
// ***************************************************************************

[make_problem() : problem
  -> let pb := problem()
     in (_machines(pb) :add machine(),
         _machines(pb) :add machine(),
         pb)
]

[make_order(pb:problem,
            order_number:integer,
            bill:list[machine]) : void
  -> _orders(pb) :add order(_order_number = order_number,
                            _bill = bill)
]

// ***************************************************************************
// * Part 4: Problem reader
// ***************************************************************************

[read_data() : problem
  -> let pb := make_problem()
     in (read_orders(pb),
         pb)
]

[read_orders(pb:problem) : void
  -> let p := fopen("ord_log.txt", "r"),
         order_number := read(p)
     in (while (order_number != 0)
           let bill := eval(read(p))
           in (make_order(pb,
                          order_number,
                          list{_machines(pb)[x] | x in bill}),
               order_number := read(p)),
         fclose(p))
]


(testOK())




