// +------------------------------------------------------------+
// | bug14.cl                                                   |
// | last update: Jan 2000 - Y. Caseau                          |
// +------------------------------------------------------------+

// ---------------------------------------------------------------
// this file contains the compilation bugs
// ---------------------------------------------------------------


// bug de thomson  -----------------------------------------------------

MISSILE <: thing(idx:integer)
PISTE <: thing(missile1:(MISSILE U {unknown}))

M0 :: MISSILE(idx = 0)
P0 :: PISTE(missile1 = M0)

[f30(p:PISTE)
 -> when m := get(missile1, p) in printf("~S",m.idx)]
[g30() -> f30(P0)]

// dans le code compile on genere
;        print_any(m->idx);         // FAUX !!!!!!!!!
(g30())


// bug du 16/2/99

// la variable lastl n'est pas prot�g�e contre le gc.
// La cellule est recup�r�e par l'appel gc()
// et on a n'importe quoi quand on cherche a imprimer en fin de procedure.
[bug3(n:integer)
 -> let lastl:list<tuple(integer,integer)> := list<tuple(integer,integer)>(),
        bestval := -1000 in
      (for i in (1 .. n)
          let t := list<tuple(integer,integer)>(tuple(i,i)) in
              lastl := t,
       // gc(),
       check("no harm done", lastl = list(tuple(n,n)))) ]

(bug3(10))

// bug du 20/12/99
Flow <: object
Flow <: object(coef:float = 1.0,
               superflow:Flow,
               idx:integer)

Parallel <: Flow()
Series <: Flow()
Edge <: Flow()

[buggy(self:Flow)
 -> let f := self.superflow, c := self.coef in
        let n := Series() in
            (case f (Parallel true,
                     Edge n.idx := 1)) ]

// a compilation bug  -------------------------------------------------

Interface <: object()
NULL_INTERFACE:Interface := Interface()
COMMethod <: object()
COMObject <: object(automat:boolean,
			  idispatch:Interface = NULL_INTERFACE) 
                                                                                                                                                                                                                             
[getDispinterface(obj:COMObject) : Interface 
=> if (obj.automat) 
	obj.idispatch
   else 
	(error("~S",obj),NULL_INTERFACE)]

[getMethodList(idisp:Interface) : list[COMMethod] 
 -> nil]

[getMethodList(obj:COMObject) : list[COMMethod] 
 => getMethodList(getDispinterface(obj))]

coco :: COMObject(automat = true, idispatch = NULL_INTERFACE)



// last but not least, a strange compiler bug

private/claire_libacadD:string :: "e:\\app\\acad13\\dos\\ads\\wcads100"
private/ahaha:string :: "check private"

[test4() ->
      printf("  xxxxxxxxx~A xxxxxx ~I xxxxxxx ~A xxxx\n",
             "aa",
             nil,
             "xxxxxxxxxxxx" /+ claire_libacadD) ]


(test4())

// a new bug by Francois
Problem <: object(name:string = "")
[makeProblem(s:string, n:integer) : Problem
   -> let pb := Problem(name = copy(s)) in pb ]


// compilation bug  by ThB
A <: object(k:integer = 10)
ListA:list[A] := list(A(),A(),A())

[shortList() : list[A] => list{ListA[i] | i in (1 .. 2)}] 
[first(l:list) : any  //Same as "some" but respect the order of the list
  => let res := unknown in 
    (for firstElt in l 
        (res := firstElt, 
         break()),
     res)]   

[f() -> let p := first(list{x.k | x in shortList()}) in nil]

(f())

// a bug from M. LeMaitre

MyIntVar <: object(
value:(integer U {unknown}) = unknown)

// Note that if myGetVar is changed into a function (not a macro)
// then everything is OK
[myGetVar(v:MyIntVar, i:integer) : MyIntVar =>
assert(i <= 999),
v]

[toto(v:MyIntVar) : boolean ->
if known?(value, myGetVar(v, 10)) true 
else false ]

[theBug() : boolean ->
let v := MyIntVar() in
   (check("macroexpansion check", toto(v) = false),
    toto(v))]

VV:MyIntVar :: MyIntVar()

(theBug())

// a new compilation bug from Naren
[h() : void -> nil]
[k() : list 
-> let x := 0
in (
while (x >= 0) h()
) as list ]             // this is dumb !!! but is will compile


// check the safe(x) pattern
** :: operation(precedence = precedence(*))
**(x:integer,y:integer) : integer -> (try safe( x * y ) catch any 100000)

(check("a large product", 1000000 ** 1000000 + 100000))


// a funny bug from Francois
AAA <: object(x:integer)
BBB <: AAA()

[getX(a:AAA) : integer -> a.x]
[getX(a:BBB) : integer -> a.x + 1]

[titi() : list<AAA>
  -> let la := toto(AAA),
         lax := list<integer>{getX(a) | a in la} in
     (printf("~S:~S",la,lax),
      la)]

[toto(c:class) : list<AAA>
-> list<AAA>(AAA(x = 1), AAA(x = 2), BBB(x = 3))]



(testOK())
