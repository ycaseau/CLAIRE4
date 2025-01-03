
;;;======================================================
;;;   Who Drinks Water? And Who owns the Zebra?
;;;     
;;;   direct translation of a CLIPS Version 6.0 Example
;;;   this is a dangerous benchmark since performance
;;;   depends mostly on instantiation order
;;;   A much better approach is to use the ECLAIR library
;;;   However, the measure of rules/s is interesting
;;;
;;;======================================================

House :: (1 .. 5)
Problem <: thing(house:integer = 0)

// use naive backtracking to solve a problem
store(house)

// used to count triggers
RuleCount:integer :: 0

// check in CLAIRE
check(p:property,x:any,y:any) : void
 => (RuleCount :+ 1,
     if (get(p,x) = 0) write(p,x,y)
     else if (get(p,x) != y) contradiction!())

checknot(p:property,x:any,y:any) : void
 => (RuleCount :+ 1,
     if (get(p,x) = y) contradiction!())

// Data for the problem
Color <: Problem()
  red :: Color() 
  green :: Color() 
  ivory :: Color()
  yellow :: Color()
  blue :: Color()

Person <: Problem()
  englishman :: Person()
  spaniard :: Person()
  ukrainian :: Person() 
  norwegian :: Person()
  japanese :: Person()

Pet <: Problem()
  dog :: Pet()
  snails :: Pet()
  fox :: Pet()
  horse :: Pet()
  zebra :: Pet()

Drink <: Problem()
  water :: Drink()
  coffee :: Drink()
  milk :: Drink()
  orange-juice :: Drink()
  tea :: Drink()

Smoke <: Problem()
  old-golds :: Smoke()
  kools :: Smoke()
  chesterfields :: Smoke()
  lucky-strikes :: Smoke()
  parliaments :: Smoke()

final(Color)
final(Person)
final(Pet)
final(Drink)
final(Smoke)


; The Englishman lives in the red house.
rule1() :: rule(p.house := h & p = englishman => check(house,red,h))

; notice that the Alldifferent constraint is spelled in many rules in the original code
allDifferent() :: rule(p.house := h => use-house(p,h))

use-house(p:Problem,h:House)
   -> (case p (Person for p2 in (Person but p) checknot(house,p2,h),
               Color  for p2 in (Color but p) checknot(house,p2,h),
               Pet    for p2 in (Pet but p) checknot(house,p2,h),
               Drink  for p2 in (Drink but p) checknot(house,p2,h),
               Smoke  for p2 in (Smoke but p) checknot(house,p2,h)))

; The Spaniard owns the dog.
rule2() :: rule(p.house := h & p = spaniard  => check(house,dog,h)) 


; The ivory house is immediately to the left of the green house,
;  where the coffee drinker lives.
rule3() :: rule(p.house := h1 & p = ivory => 
                (if (h1 <= 4) check(house,green,h1 + 1) else contradiction!()))
rule4() :: rule(p.house := h & p = coffee => check(house,green,h))

; The milk drinker lives in the middle house.
(milk.house := 3)

; The man who smokes Old Golds also keeps snails.
rule5() :: rule(p.house := h & p = old-golds => check(house,snails,h)) 

; The Ukrainian drinks tea.
rule6() :: rule(p.house := h & p = ukrainian  => check(house,tea,h))  

; The Norwegian resides in the first house on the left.
(norwegian.house := 1)

; Chesterfields smoker lives next door to the fox owner.
rule7() :: rule(p.house := h & p = chesterfields  => next-to(fox,h))
rule7bis() :: rule(p.house := h & p = fox => next-to(chesterfields,h))

next-to(p:Problem,h:House) : void
  => (if (p.house > 0 & p.house != (h - 1)) check(house,p,h + 1))

; The Lucky Strike smoker drinks orange juice.
rule8() :: rule(p.house := h & p = lucky-strikes  => check(house,orange-juice,h))

; The Japanese smokes Parliaments
rule9() :: rule(p.house := h & p = japanese  => check(house,parliaments,h))

; The horse owner lives next to the Kools smoker, whose house is yellow.
rule10() :: rule(p.house := h & p = horse => next-to(kools,h)) 
rule10bis() :: rule(p.house := h & p = kools  => next-to(horse,h))

rule11() :: rule( p.house := h & p = kools  => check(house,yellow,h))

; The Norwegian lives next to the blue house.
rule12() :: rule(p.house := h & p = norwegian  => next-to(blue,h))
rule12bis() :: rule(p.house := h & p = blue  => next-to(norwegian,h))

// to solve we use the same instantiation order as the original CLIPS example
find-solution()
 -> solve(list<Problem>(
      englishman,
      red,
      spaniard,
      dog,
      ivory,
      green,
      coffee,
      milk,
      old-golds,
      snails,
      ukrainian,
      tea,
      norwegian,
      chesterfields,
      fox,
      lucky-strikes,
      orange-juice,
      japanese,
      parliaments,
      horse,
      kools,
      yellow,
      blue,
      water,
      zebra))

// generate and test
solve(l:list[Problem]) : boolean -> solve(l,1)
solve(l:list[Problem],n:integer)  : boolean
  -> (if (n > length(l))
        // (print-solution(), false)
         (if (verbose() >= 1) print-solution(),false)		// no more problems
      else let p := l[n] in              
         (if (p.house = 0)
             exists(h in (1 .. 5) |
                    branch( (//[5] (~A) try ~S -> ~S // n, p, h,
                             p.house := h,
                             solve(l, n + 1))))
          else solve(l, n + 1)))        // move to next


// prints the solution
print-solution()
; -> (for h in (1 .. 5)
;       printf("~A:~S\n",h, {p in Problem | p.house = h}))
 -> (for h in (1 .. 5)
       printf("~A:~S ~S ~S ~S ~S\n",h,
               some(x in Color | x.house = h),
               some(x in Person | x.house = h),
               some(x in Pet | x.house = h),
               some(x in Smoke | x.house = h),
               some(x in Drink | x.house = h)))

// start
do_test() : void
  -> (time_set(),
      RuleCount := 0,
      find-solution(),
      time_show(),
      printf("~A rules firings\n",RuleCount))

do_test1() : void
  -> (time_set(),
      verbose() := 0,
      RuleCount := 0,
      for i in (1 .. 100) find-solution(),
      time_show(),
      printf("~A rules firings\n",RuleCount))

main() : void -> 
  (//[0] look for a solution 100 time ..... //,
   do_test1(),
   exit(0))
