// + ------------------------------------------------------- +
// | Micro Benchmark Suite for Inference Engines             |
// |                                collected by Y.Caseau    |
// | airline.cl                                              |
// + --------------------------------------------------------+

// a travel example: a transitive closure with a parameter (time)
// we only use two attributes to propagate to make the compilation easier
// (i.e. shorter). A more proper program would use 4.

// forward definition
flight <: thing
travel_plan <: object
request <: thing

city <: thing(time_zone:integer = 0,
               previous_request:set<request>,
               starting:set<flight>,
               arriving:set<flight>,
               starting_plan:set<travel_plan>,
               arriving_plan:set<travel_plan>)

flight <: thing(fromc:city,to:city,
                depart:integer = 0,   ; departure time
                time:integer)        ; flight duration


travel_plan <: object( depart:integer = 0,     ; departure
                       time:integer = 0,       ; actual flight time
                       end_at:integer,         ; arrival
                       start_plan:city,
                       arrive_plan:city,
                       transit:list<city>,
                       content:list<flight>)

self_print(x:travel_plan)
 -> printf("~A:~S:~A",depart(x),content(x),end_at(x))

request <: thing(depart:integer = 7,
                  end_at:integer = 20,
                  solutions:list<travel_plan>,
                  start:city,
                  arrive:city)

// setup inverses (could be done using rules)
(inverse(starting) := fromc,
 inverse(arriving) := to,
 inverse(arrive_plan) := arriving_plan,
 inverse(start_plan) := starting_plan)

// rule compiler declarations
(known!(to,fromc,start_plan,arrive_plan))
// mode(exists)   WHAT THE FUCK IS THIS ? 
// event(arrive)

DEBUG:integer := 5
Cnt:integer :: 0

// create a request (if needed)
request!(c1:city,c2:city,dep:integer,arr:integer)
 -> let r := unknown in
      (for r2 in previous_request(c1)
         (if (arrive(r2) = c2 & depart(r2) = dep & end_at(r2) >= arr)
             (r := r2, break(true))),
       if known?(r) r
       else request(name = (name(c1) /+ name(c2)),
                    depart = dep, end_at = arr))

// register a solution if it is new
registerSol(r:request,t:travel_plan)
 -> (if not(for t2 in solutions(r)
                (if (content(t) = content(t2)) break()))
        solutions(r) :add t)

// a simple travel plan !
simple_plan(r:request,f:flight) :: rule(
    (fromc(f) = start(r)  & to(f) = arrive(r) &
     depart(f) >= depart(r) & (depart(f) + time(f)) <= end_at(r))
    => let t := travel_plan(depart = depart(f),  end_at = (depart(f) + time(f)),
                            content = list<flight>(f), time = time(f)) in
        (//[DEBUG] create simple travel plan ~S // t,
         Cnt :+ 1,
         start_plan(t) := fromc(f),
         arrive_plan(t) := to(f),
         registerSol(r,t)))

// create an auxiliary request
complex_plan(r:request,f:flight) :: rule(
  (fromc(f) = start(r) & to(f) != arrive(r) &
   depart(f) >= depart(r) & (depart(f) + time(f)) <= end_at(r) )
   => let r2 := request!(to(f),arrive(r),(depart(f) + time(f)),end_at(r)) in
        (Cnt :+ 1,
         if not(r2 % previous_request(to(f)))
           (//[DEBUG] create req. ~S: ~A -> ~S // r2, depart(r),end_at(r),
            previous_request(to(f)) :add r2,
            start(r2) := to(f),
            arrive(r2) := arrive(r))))


// how to merge 2 plans: transitive closure
merge_plan(r:request,t:travel_plan,f:flight) :: rule(
   (arrive(r) = arrive_plan(t) & end_at(t) <= end_at(r) &
    fromc(f) = start(r) & to(f) = start_plan(t) &
    depart(f) >= depart(r) & (time(f) + depart(f)) <= depart(t))
    => (Cnt :+ 1,
        if (to(f) % transit(t)) nil    ;; prevent loops !
        else let t2 := travel_plan(depart = depart(f), end_at = end_at(t),
                                   content = list<flight>(f) /+ content(t),
                                   transit = list<city>(to(f)) /+ transit(t),
                                   time = (time(f) + time(t))) in
          (//[DEBUG] create complex travel plan ~S for ~S // t2,r,
           start_plan(t2) := fromc(f),
           arrive_plan(t2) := arrive_plan(t),
           registerSol(r,t2))))


// interface to create a query
howto(c1:city,c2:city,dep:integer,arr:integer) : void
  -> let r := request!(c1,c2,dep,arr) in
       (if not(solutions(r))
           (start(r) := c1, arrive(r) := c2),
        printf("solutions = ~S\n",solutions(r)),
        printf("best solutions = ~S\n",solutions!(r)))


// list of locally optimal solutions
[better(t1:travel_plan,t2:travel_plan) : boolean
  -> (time(t1) <= time(t2) & end_at(t1) <= end_at(t2) &
      length(transit(t1)) <= length(transit(t2))) ]

[solutions!(r:request) : list
  -> let l1 := list(), l2 := solutions(r) in
       (for t2 in l2
         (if exists(t1 in l1 | better(t1,t2)) nil
          else if (for t1 in l1
               (if better(t2,t1)
                   (l1 :delete t1, l1 :add t2, break()))) nil
          else l1 :add t2),
        l1) ]

// schedule
PRINT:boolean :: true
[howto(c1:city,c2:city,l:integer) : void
 -> time_set(),
    for i in (7 .. 19)
     (let r := request!(c1,c2,i,i + l) in
        (start(r) := c1, arrive(r) := c2,
         if PRINT printf("~A == ~S\n",i,solutions!(r)),
         cleanTP())),
    time_show() ]

// remove the previous travel plans
[cleanTP()
  -> for c in city
       (c.starting_plan := set<travel_plan>(),
        c.arriving_plan := set<travel_plan>()) ]


// ----------------- flight database ------------------------------

Paris :: city()
London :: city()
Frankfurt :: city()
Rome :: city()
Berlin :: city()
Athens :: city()
Madrid :: city()


AF001 :: flight(fromc = Paris, to = London, time = 1, depart = 7)
AF002 :: flight(fromc = London, to = Paris, time = 1, depart = 9)
AF011 :: flight(fromc = Paris, to = London, time = 1, depart = 12)
AF012 :: flight(fromc = London, to = Paris, time = 1, depart = 13)
AF021 :: flight(fromc = Paris, to = London, time = 1, depart = 18)
AF022 :: flight(fromc = London, to = Paris, time = 1, depart = 19)
AF201 :: flight(fromc = Paris, to = Rome, time = 2, depart = 10)
AF202 :: flight(fromc = Rome, to = Paris, time = 2, depart = 13)
AF301 :: flight(fromc = Paris, to = Frankfurt, time = 1, depart = 8)
AF302 :: flight(fromc = Frankfurt, to = Paris, time = 1, depart = 9)
AF311 :: flight(fromc = Paris, to = Frankfurt, time = 1, depart = 10)
AF312 :: flight(fromc = Frankfurt, to = Paris, time = 1, depart = 12)
AF321 :: flight(fromc = Paris, to = Frankfurt, time = 1, depart = 15)
AF322 :: flight(fromc = Frankfurt, to = Paris, time = 1, depart = 16)
AF331 :: flight(fromc = Paris, to = Frankfurt, time = 1, depart = 19)
AF332 :: flight(fromc = Frankfurt, to = Paris, time = 1, depart = 20)
AF401 :: flight(fromc = Paris, to = Athens, time = 4, depart = 9)
AF402 :: flight(fromc = Athens, to = Paris, time = 4, depart = 16)
AF501 :: flight(fromc = Paris, to = Berlin, time = 2, depart = 8)
AF502 :: flight(fromc = Berlin, to = Paris, time = 2, depart = 18)
AF601 :: flight(fromc = Paris, to = Madrid, time = 2, depart = 11)
AF602 :: flight(fromc = Madrid, to = Paris, time = 2, depart = 17)

BA001 :: flight(fromc = London, to = Paris, time = 1, depart = 8)
BA002 :: flight(fromc = Paris, to = London, time = 1, depart = 10)
BA011 :: flight(fromc = London, to = Paris, time = 1, depart = 16)
BA012 :: flight(fromc = Paris, to = London, time = 1, depart = 18)
BA021 :: flight(fromc = London, to = Berlin, time = 2, depart = 8)
BA022 :: flight(fromc = Berlin, to = London, time = 2, depart = 14)
BA201 :: flight(fromc = London, to = Rome, time = 2, depart = 11)
BA202 :: flight(fromc = Rome, to = London, time = 2, depart = 14)
BA301 :: flight(fromc = London, to = Frankfurt, time = 1, depart = 8)
BA302 :: flight(fromc = Frankfurt, to = London, time = 1, depart = 9)
BA311 :: flight(fromc = London, to = Frankfurt, time = 1, depart = 11)
BA312 :: flight(fromc = Frankfurt, to = London, time = 1, depart = 12)
BA321 :: flight(fromc = London, to = Frankfurt, time = 1, depart = 16)
BA322 :: flight(fromc = Frankfurt, to = London, time = 1, depart = 17)
BA331 :: flight(fromc = London, to = Frankfurt, time = 1, depart = 19)
BA332 :: flight(fromc = Frankfurt, to = London, time = 1, depart = 20)
BA401 :: flight(fromc = London, to = Athens, time = 4, depart = 15)
BA402 :: flight(fromc = Athens, to = London, time = 4, depart = 9)
BA501 :: flight(fromc = London, to = Berlin, time = 2, depart = 8)
BA502 :: flight(fromc = Berlin, to = London, time = 2, depart = 18)
BA601 :: flight(fromc = London, to = Madrid, time = 2, depart = 8)
BA602 :: flight(fromc = Madrid, to = Frankfurt, time = 2, depart = 15)

LU001 :: flight(fromc = Frankfurt, to = Paris, time = 1, depart = 8)
LU002 :: flight(fromc = Paris, to = Frankfurt, time = 1, depart = 10)
LU011 :: flight(fromc = Frankfurt, to = Paris, time = 1, depart = 16)
LU012 :: flight(fromc = Paris, to = Frankfurt, time = 1, depart = 18)
LU201 :: flight(fromc = Frankfurt, to = Rome, time = 2, depart = 11)
LU202 :: flight(fromc = Rome, to = Frankfurt, time = 2, depart = 14)
LU211 :: flight(fromc = Frankfurt, to = Rome, time = 2, depart = 17)
LU212 :: flight(fromc = Rome, to = Frankfurt, time = 2, depart = 19)
LU213 :: flight(fromc = Rome, to = Frankfurt, time = 2, depart = 7)
LU214 :: flight(fromc = Frankfurt, to = Rome, time = 2, depart = 19)
LU301 :: flight(fromc = Berlin, to = London, time = 1, depart = 8)
LU302 :: flight(fromc = London, to = Berlin, time = 1, depart = 9)
LU311 :: flight(fromc = Frankfurt, to = London, time = 1, depart = 11)
LU312 :: flight(fromc = London, to = Frankfurt, time = 1, depart = 12)
LU321 :: flight(fromc = Frankfurt, to = London, time = 1, depart = 16)
LU322 :: flight(fromc = London, to = Frankfurt, time = 1, depart = 17)
LU401 :: flight(fromc = Frankfurt, to = Athens, time = 3, depart = 9)
LU402 :: flight(fromc = Athens, to = Frankfurt, time = 3, depart = 13)
LU411 :: flight(fromc = Berlin, to = Athens, time = 2, depart = 14)
LU412 :: flight(fromc = Athens, to = Berlin, time = 2, depart = 18)
LU501 :: flight(fromc = Frankfurt, to = Berlin, time = 1, depart = 7)
LU502 :: flight(fromc = Berlin, to = Frankfurt, time = 1, depart = 9)
LU511 :: flight(fromc = Frankfurt, to = Berlin, time = 1, depart = 12)
LU512 :: flight(fromc = Berlin, to = Frankfurt, time = 1, depart = 13)
LU521 :: flight(fromc = Frankfurt, to = Berlin, time = 1, depart = 18)
LU522 :: flight(fromc = Berlin, to = Frankfurt, time = 1, depart = 19)
LU601 :: flight(fromc = Berlin, to = Madrid, time = 3, depart = 19)
LU602 :: flight(fromc = Madrid, to = Paris, time = 2, depart = 8)

// test
do_test() : void -> howto(Paris,Athens,10)

[do_test1() : void
 -> time_set(),
    for i in (1 .. 200) request!(Paris,Athens,i,10 + i),
    time_show() ]

