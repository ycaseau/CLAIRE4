// +---------------------------------------------------------+
// | MARIE                                       Yves Caseau |
// | test-monkey                                             |
// |                      Copyright (C) 1986 by Y. CASEAU    |
// |                          Du a l'amabilite de P. Dixneuf |
// +---------------------------------------------------------+

// this program is a LAURE adaptation of a ELOISE program given by
// Marc Porcheron (et P. Dixneuf).
// this is the ultimate :-) v3.2 version
// ------------------------------------------------------------------

TRACE:integer :: 4 // 1 for debug to see the rules

// the set of places
set_of_places <: thing()
Pa :: set_of_places()
Pb :: set_of_places()
Pc :: set_of_places()
Pd :: set_of_places()
Pe :: set_of_places()
nowhere :: set_of_places()

// the set of positions
positions <: thing()
ground :: positions()
high :: positions()
on_ladder :: positions()

// definition des classes
//
fruit <: thing(is_at:set_of_places = nowhere,
               is_on:positions = ground)
nothing :: fruit()

tool <: thing(is_at:set_of_places = nowhere)
ladder :: tool()

// the possible actions
//
actions <: thing()
catch_a :: actions()
move :: actions()
move_with_ladder :: actions()

// the goals
goal <: object(action:actions,
                         concerned_with:fruit,
                         to_be_found_in:set_of_places = nowhere)

[self_print(self:goal)  : void
  -> printf("~S[~S,~S]",action(self),concerned_with(self),
             get(to_be_found_in,self))]


// final(tool)

ape <: thing(current_place:set_of_places = nowhere,
              current_position:positions,
              hold:fruit,
              goals:set[goal] = {})

// we use an event which is the fact that something happened to the ape
consider :: property(domain = ape, range = goal)

// definition des regles ------------------------------------------

// taking a fruit
//
s-1() :: rule(
  (consider(x,y) & action(y) = catch_a &
   current_place(x) = is_at(concerned_with(y)) &
   hold(x) = nothing & current_position(x) = is_on(concerned_with(y)))
  => (  //[TRACE] ~S catches ~S // x,concerned_with(y),
        goals(x) :delete y,
        hold(x) := concerned_with(y)))

// taking a fruit with the ladder
//
s-2() :: rule(
  (consider(x,y) & action(y) = catch_a &
   current_place(x) = is_at(concerned_with(y)) &
   current_position(x) = on_ladder &
   is_on(concerned_with(y)) = high & hold(x) = nothing)
  => (  //[TRACE] ~S catches ~S // x,concerned_with(y),
        goals(x) :delete y,
        hold(x) := concerned_with(y)))


// Dropping a fruit
//
s-3() :: rule(
  (consider(x,y) & action(y) = catch_a & hold(x) != nothing)
  => ( //[TRACE] ~S drops ~S // x,hold(x),
       is_on(hold(x)) := ground,
       hold(x) := nothing) )


// moving to get a fruit on the ground
//
s-4() :: rule(
  (consider(x,y) & action(y) = catch_a &
   current_place(x) != is_at(concerned_with(y)) &
   is_on(concerned_with(y)) = ground)
  => ( goals(x) :add goal(concerned_with = concerned_with(y),
                          action = move,
                          to_be_found_in = is_at(concerned_with(y)))) )


// climbing the ladder
//
s-5() :: rule(
  (consider(x,y) & y.action = catch_a &
   x.current_place = y.concerned_with.is_at &
   x.current_position = ground & y.concerned_with.is_on = high &
   ladder.is_at = x.current_place)
  => ( //[TRACE] ~S climbs on the ladder // x,
       current_position(x) := on_ladder) )


// moving to get a fruit in high position
//
s-6() :: rule(
  (consider(x,y) & action(y) = catch_a &
   is_at(ladder) != is_at(concerned_with(y)) & is_on(concerned_with(y)) = high)
  => ( goals(x) :add  goal(action = move_with_ladder,
                           concerned_with = concerned_with(y),
                           to_be_found_in = is_at(concerned_with(y))) ))


// going down the ladder to move
//
s-7() :: rule(
  (consider(x,y) & action(y) = move &
   current_place(x) != is_at(concerned_with(y)) &
   current_position(x) = on_ladder )
  => (//[TRACE] ~S goes down the ladder // x,
      current_position(x) := ground) )


// going down the ladder to move with it
//
s-8() :: rule(
  (consider(x,y) & action(y) = move_with_ladder &
   current_place(x) != is_at(concerned_with(y)) &
   current_position(x) = on_ladder )
  => (//[TRACE] ~S goes down the ladder // x,
      current_position(x) := ground) )


// moving to get the ladder
//
s-9() :: rule(
  (consider(x,y) & action(y) = move_with_ladder &
   current_place(x) != is_at(ladder))
  => (  goals(x) :add goal(concerned_with = concerned_with(y),
                           action = move,
                           to_be_found_in = is_at(ladder))) )


// moving when the ladder is there
//
s-10() :: rule(
  (consider(x,y) & y % goals(x) & action(y) = move_with_ladder &
   current_place(x) = is_at(ladder) & current_position(x) = ground )
  => ( //[TRACE] ~S moves [~S] with the ladder to ~S // x,y,to_be_found_in(y),
       goals(x) :delete y,
       is_at(ladder) := to_be_found_in(y),
       current_place(x) := to_be_found_in(y)) )


// moving without the ladder
//
s-11() :: rule(
  (consider(x,y) & y % goals(x) & action(y) = move 
   & current_position(x) = ground )
   => ( //[TRACE] ~S moves [~S] (without ladder) to ~S // x,y,to_be_found_in(y),
       goals(x) :delete y,
       //[TRACE] now the list of goals is ~S // goals(x),
       current_place(x) := to_be_found_in(y)) )


// everything that is holded moves is the same way
//
s-12() :: rule(
  current_place(z) := y => when f := get(hold,z) in is_at(f) := y)

// re-evaluate rules
[consider(x:ape) : void 
  -> for y in x.goals consider(x,y) ]

// rules to re-evaluate
rs-1() :: rule(hold(x) := y => consider(x))
rs-2() :: rule(goals(x) :add y => consider(x,y))
rs-3() :: rule(current_position(x) := y => consider(x))
rs-4() :: rule(current_place(x) := y => consider(x))
  
// -----------------------------------------------------------------

banana :: fruit(is_at = Pa, is_on = high)
apple :: fruit(is_at = Pb)
orange :: fruit(is_at = Pd, is_on = high)
pear :: fruit(is_at = Pa, is_on = ground)
grape :: fruit(is_at = Pe, is_on = high)
ladder :: tool( is_at = Pc)
cheetah :: ape()

(put(current_place,cheetah,Pb),
 put(current_position,cheetah,ground),
 put(hold,cheetah,apple))

goal-OK:any  :: goal(action = catch_a,
                     concerned_with = banana,
                     to_be_found_in = Pa)

// test method
//
[do_test0() : void
  -> time_set(),
     goals(cheetah) :add goal-OK,
     time_show() ]
	 
	 
goal-1 :: goal(action = catch_a, concerned_with = banana)
goal-2 :: goal(action = catch_a, concerned_with = apple)
goal-3 :: goal(action = catch_a, concerned_with = orange)
goal-4 :: goal(action = catch_a, concerned_with = pear)
goal-5 :: goal(action = catch_a, concerned_with = grape)

// test method
// x 100
[do_test()  : void
  -> time_set(),
     for x in (1 .. 100000)     // 100 interpreted, 100000 compiled
       (goals(cheetah) :add goal-1,
        //[TRACE] do I have the banana?  //,
        goals(cheetah) :add goal-2,
        //[TRACE] do I have the apple ?  //,
        goals(cheetah) :add goal-3,
        //[TRACE] do I have the orange? //,
        goals(cheetah) :add goal-4,
        goals(cheetah) :add goal-5,
        is_on(grape) := high,
        is_on(pear) := high),
     time_show()]
     
[simple_test() : void
  -> goals(cheetah) :add goal-1 ]


main() : void ->
  (printf("start do_test 100 times\n"),
   do_test(),
   exit(0))


