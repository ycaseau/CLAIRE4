;; This programs compute the cities that one can visit from another city
;;
FLIGHT <: object
CITY <: thing

CITY <: thing(flight:set[CITY],
               directflight:set[CITY],
               flight_of:set[CITY],
               directflight_of:set[CITY],
               departs:set[FLIGHT],
               arrives:set[FLIGHT],
			   query:boolean = false)
			   
FLIGHT <: object(leaves:CITY,
                  go-to:CITY,
                  departure:integer,
                  arrival:integer,
                  id:symbol)


// set inverses
(inverse(flight) := flight_of,
 inverse(directflight) := directflight_of,
 inverse(leaves) := departs,
 inverse(go-to) := arrives)
 
event(query,flight,directflight)

// rules
Rr1(x:CITY, y:CITY) :: rule(
  query(x) = true & exists(z:FLIGHT, leaves(z) = x & go-to(z) = y) 
  =>  (directflight(x) :add y, query(y) := true))

Rr2(x:CITY, y:CITY) :: rule(
  y % directflight(x) | exists(z:CITY, z % flight(x) & y % directflight(z))
  => flight(x) :add y )

[do_test()
  -> let m := value("muenchen") in
       (time_set(),
        query(m) := true,
        time_show(),
        printf("there are ~A cities where you can go from Muenchen.~%",
                size(flight(m)))) ]
		
[do_load() -> load(home() / "lib\\other\\rule\\airdata") ]
