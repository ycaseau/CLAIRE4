// +------------------------------------------------------------+
// | bug11.cl                                                   |
// | last update: August 2021 - Y. Caseau                       |
// +------------------------------------------------------------+


// ---------------------------------------------------------------
// this file contains bugs related with classical examples
// for instance, examples from the documentation !
// ---------------------------------------------------------------


// classical stack example


[stack[of] <: thing(of:type, content:list)]

[top(s:stack[of = X]) : type[X]
  -> last(content(s)) ]

[pop(s:stack[of = X]) : type[X]
  -> let l := content(s), n := length(l), x := l[n] in
      (content(s) := shrink(l, n - 1), x) ]

[push(s:stack[of = X],y:X) : type[s] 
  -> content(s) :add y, s ]

S :: stack(integer)

check("stack1", top(push(S,1)) = 1)
check("stack2", pop(S) = 1)
check("stack3", (push(push(push(S,1),2),3), pop(S) = 3))
check("stack4", top(S) = 2)

assert(try (push(S,1.1), false) catch error true) 

[member(s:{stack},t:type) -> t @ of ]
[iterate(s:stack,v:Variable,e:any)
  => for v in content(s) e ]

// bug from CLAUDE
activity <: thing()
A1 :: activity()
ST :: stack(of = activity)

[tolist(s:stack[of = activity]) : list[activity]
  -> list(top(s)) ]

[ID(s:stack[of = activity]) : stack[of = activity]
  -> s ]

[car(s:stack[of = X]) : type[X]
  -> car(s.content)
]

/* another arnaud bug from the doc !!!  - 
Clong <: import()
(c_interface(Clong,"long "))

[long!(xa:integer) : Clong -> let x := xa in externC("(long)(x)",Clong)]
[+(x:Clong,y:Clong) : Clong -> externC("(x + y)",Clong) ]

[checkLong()
  -> printf("test long ---------------------------\n"),
     show(long!(12) + long!(12))   ]

(#if (compiler.active? & compiler.loading?) checkLong())

*/

// phone application
// definition of the module

// value is a table that stores the phone #
private/value_table[s:string] : string := "unknown"

// lower returns the lower case version of a string
// (i.e. lower("aBcD") = "abcd")
lower(s:string) : string
  -> let s2 := copy(s) in
	(for i in (1 .. length(s))
	    (if (integer!(s2[i]) % (integer!('A') .. integer!('Z')))
	        s2[i] := char!(integer!(s2[i]) + 32)),
         s2)


claire/store(name:string,phone:string)
	  -> (value_table[lower(name)] := phone)
claire/dial(name:string) : string   // returns the phone #
	  -> value_table[lower(name)]

(store("Yves","0102"),
 store("Mr X","1234"),
 check("phone",dial("yves") = "0102"))


// HIFI example --------------------------------------------------------------------

component <: thing(price:integer, brand:string)

amplifier <: component(	power:integer, input:integer,
			ohms:set[{4,8}])

speaker <: component(maxpower:integer, ohm:{4,8})
headphone <: component(maxpower:integer, ohm:{4,8})
musical_source <: component(sensitivity:integer)
CDplayer <: musical_source(laser_beams:(1 .. 3))
turntable <: musical_source()
tuner <: musical_source()

B :: thing() 	C :: thing() 	 nodolby :: thing()

tape <: musical_source(dolby:{nodolby,B,C})
stereo <: object(	sources:set[musical_source],
			amp:amplifier,
			out:set[speaker U headphone],
			warranty:boolean = false)



// instantiation rule example -----------------------------------------

instantiation :: property(domain = amplifier, range = string)
[close(x:amplifier) : amplifier -> instantiation(x,date!(1)), x ]

controlRule() :: rule( instantiation(x,s) 
                        => printf("--- create ~S at ~A \n",x,s))


// very simple rule to see if we track 
simpleRule() :: rule( price(x) := y   
                      =>    printf("--- change ~S cost to ~A \n",x,y))                   


//instances
amp1 :: amplifier(power = 120, input = 4, ohms = {4,8},
 	          price = 400, brand = "Okyonino")
amp2 :: amplifier(power = 50, input = 2, ohms = {4},
                  price = 130, brand = "Cheapy")
tuner1 :: tuner(	sensitivity = 10, price = 200, brand = "Okyonino")
tuner2 :: tuner(	sensitivity = 30, price = 80, brand = "Cheapy")
CD1 :: CDplayer(	sensitivity = 3, price = 300,
					laser_beams = 3, brand = "Okyonino")
CD2 :: CDplayer(	sensitivity = 7, price = 180,
					laser_beams = 2, brand = "Okyonino")
CD3 :: CDplayer(	sensitivity = 15, price = 110,
					laser_beams = 1, brand = "Cheapy")
t1 :: tape(	sensitivity = 40, price = 70,
				dolby = nodolby, brand = "Cheapy")

s1 :: speaker(	ohm = 8, maxpower = 150, price = 1000, brand = "Magisound")
s2 :: speaker(	ohm = 8, maxpower = 80,	price = 400, brand = "Magisound")
s3 :: speaker(	ohm = 4, maxpower = 40, price = 150, brand = "Cheapy")
ph :: speaker( ohm = 4, maxpower = 50, price = 50, brand = "Okyonino")
hd :: headphone( ohm = 4, maxpower = 200, price = 50, brand = "Okyonino")


// methods
sum(s:list[integer]) : integer
	-> let n := 0 in (for y in s n :+ y, n)
total_price(s:stereo) : integer
	-> sum(list{x.price | x in s.sources U set(s.amp) U s.out})
InventoryTotal:integer :: 0

discount[s:set[component]] : integer := 0
(discount[{amp1,s1}] := 1200,
 discount[{amp1,CD1}] := 600)

// check that dictionary works on sets 
(check("set dictionary", discount[{amp1,s1}] = 1200))

[best_price(s:set[component]) : integer
   -> let p := 100000 in
	(if (size(s) = 0) p := 0
   else if (size(s) = 1) p := price(the(s))
   else for s2 in nth(set,s)				// decompose s = s2 U ...
	    let x := size(s2),
		    p2 := (if (x > 1) discount[s2]
			         else if (x = 1) price(the(s2))
			         else 0)  in
		(if (p2 > 0) p :min (p2 + best_price(difference(s,s2)))),
	 p) ]

(check("find best discount",best_price({amp1,s1,CD1,t1}) = 1200 + 300 + 70),
 check("discount2", best_price({amp1,s2,CD1,t1}) = 600 + 400 + 70))

// Rules
technical_problem <: exception(s:string)

compatibility1() :: rule(
		(st.out :add sp) & not(sp.ohm % st.amp.ohms)
		=> technical_problem(s = "conflict speakers-amp"))

compatibility2() :: rule(
		(st.sources :add x) & size(st.sources) > st.amp.input
		=> technical_problem(s = "too many sources"))

compatibility3() :: rule(
		(st.out :add x) & x.maxpower < st.amp.power
		=> technical_problem(s = "amp to strong for the speakers"))

my_system :: stereo(amp = amp2)

(exists(sp in speaker |
    (//[0] will try to add ~S // sp,
     try (my_system.out :add sp, true)
     catch technical_problem
        (//[0] rejects ~S because ~A // sp, exception!().s,
         my_system.out :delete sp,
         false))),

check("my system",my_system.out = set(ph)))

store(out,sources)

DBC:integer :: 0

all_possible_stereos() : list[stereo]
  -> let solutions := list<stereo>() , syst:stereo := stereo() in
	(for a in amplifier
	    (syst.amp := a,
       for sp in speaker try
		       (choice(),
            //[0] try ~S speaker ~S ..... // syst, sp,
             syst.out := set(sp),
	           for h in headphone try
			         (choice(),
                DBC :+ 1,
                //[0] [~A] try ~S headphones ~S ..... // DBC, syst, h,
                syst.out :add h,
                //[0] tried headphones ~S ..... // h,
		    	      for s1 in musical_source try
			              (choice(),
                      //[0] try main source ~S ..... // s1,
                      syst.sources := set(s1),
                      for s2 in {s in musical_source | owner(s) != owner(s1) & s.price < s1.price}
			               	try (choice(),
				                   syst.sources :add s2,
                           //[0] add ~S // syst,
				                   solutions :add copy(syst),
                           backtrack())
				              catch technical_problem backtrack(),
                      backtrack())
			         catch technical_problem backtrack(),
               backtrack())
	        catch technical_problem backtrack(),
          backtrack())
	    catch technical_problem  backtrack()),
      solutions)   

price_order(s1:stereo, s2:stereo) : boolean -> (total_price(s1) <= total_price(s2))

cheapest() : list[stereo] ->
let l := all_possible_stereos() in sort(price_order @ stereo, l)

self_print(s:stereo) -> printf("~A:~S->~A@ ~A$",s.sources,s.amp,s.out,total_price(s))

(check("best price",total_price(cheapest()[1]) = 380))


// queens example -----------------------------------------------------------

column[n:(1 .. 8)] : (1 .. 8) := unknown
possible[x:(1 .. 8), y:(1 .. 8)] : boolean := true

store(column, possible)

r1() :: rule(
 		column[x] := z => for y in ((1 .. 8) but x) possible[y,z] := false)

r2() :: rule(
        column[x] := z => let d := x + z in
                            for y in (max(1,d - 8) .. min(d - 1, 8))
                                (//[5] r2: forbid  ~S ~S // y, d - y,
                                 possible[y,d - y] := false))
r3() :: rule(
        column[x] := z => let d := z - x in 
                            for y in (max(1,1 - d) .. min(8,(8 - d)))
                                (//[5] r3: forbid ~S ~S (d = ~A) // y, y + d,d,
                                 possible[y,y + d] := false))

queens(n:(0 .. 8)) : boolean
 -> (	if (n = 0) true
	 else exists(p in (1 .. 8) |
					(possible[n,p]	&
                 branch( (//[5] try column ~A = ~A // n,p,
                          column[n] := p,
                          queens(n - 1)) ))))

(check("queens",queens(8)),
 printf("sol : ~A \n",list{column[i] | i in (1 .. 8)}))


// small rule example from Sylvain
Anobject <: thing(status:integer = 0)
Atest :: Anobject()
Acontrol:integer :: 0 

check_status() :: rule(x.status := (y <- z) 
                       => (printf("status(~S) changed from ~S to ~S\n",x,z,y),
                           Acontrol :+ (y - z)))
    
goRule()
  -> (Atest.status := 1,
      check("rule",Acontrol = 1))
    
(goRule())

(testOK())

