// +------------------------------------------------------------+
// | bug3.cl                                                    |
// | last update: Jan 2000 - Y. Caseau                          |
// +------------------------------------------------------------+


// ---------------------------------------------------------------
// this file contains table related bugs
// ---------------------------------------------------------------

toto[i:any,j:any] : integer := 0
tata[t:tuple(any,any)] : integer := 1
tutu[i:(1 .. 19), j:(1 .. 19)] : integer := 2

[test()
  -> if (toto[class,class] = 0 & tata[class,class] = 1 & tutu[2,2] = 2) nil
     else error("bug")]

// bug from Xavier (22/2/99)
a[x:integer] : integer := 0
(a.comment := "22")

// a bug by Yves
Truc[x:(1 .. 10),y:(1 .. 10)] : integer := (x + y)

(for x in (1 .. 10)
  for y in (1 .. 10)
    assert(Truc[x,y] = x + y))

// was impossible
(for x in (1 .. 10) Truc[x,1] :+ 1)

// a check on tables
// (Id(compiler.safety := 2))

XX:any :: 1
AB[x:integer] : integer := unknown

(XX := AB[12], AB[11] := 11, check("writeAB", XX = 0 & AB[11] = 11))
(try (AB["21"] := 12, record("failureMEGA BUG I",1)) 
 catch any record("OK1",1))

BA[x:(1 .. 20)] : integer := unknown

(XX := BA[12], BA[11] := 11, check("writeBA", XX = 0 & BA[11] = 11))

(try (BA[21] := 12, record("failure: MEGA BUG II",2)) 
 catch any record("OK2",2))

goli() -> BA[12]

ABC[i:(0 .. 10),j:(0 .. 10)] : integer  := 13

(XX := ABC[1,2], ABC[1,1] := 11, check("writeABC", XX = 13 & ABC[1,1] = 11))
(try (ABC[1,20] := 13, record("failure, MEGA BUG III",3)) 
 catch any record("OK3",3))


// a bug by Francois --------------------------------------------------------
INT :: (1 .. 3)
Aircraft <: thing(idx:INT = 2)
Tab[x:INT,y:INT] : integer := 0

[Test()
 -> let N := 2, l := list{Aircraft() | i in (1 .. N)} in
     (//[0] l = ~S // l,
      for i in (1 .. N)
        for j in (1 .. N)
           printf("~S, ~S",Tab[2,2],l[i].idx),
      for i in (1 .. N)
        for j in (1 .. N)
           printf("~S ~S -> ~S  ",l[i].idx,l[j].idx,
               Tab[l[i].idx, l[j].idx]),
      princ("\n")) ]

(Test())


;; a bug by FXJ --------------------------------------------------------------
// demander a FXJ -> ou est le bug ?

; Ce fichier permet de mettre en evidence un probleme dans la
; gestion de la table de hash pour les array[integer].
; si tu traces le system.graph de ancetre[], on decouvre que entre le
; noeud 34  et 17, il y a un probleme dans le graphe de relation...
; ainsi ancetre n'est pu du tout alimenté comme il faut et il est
; impossible de dépiler la solution : depile-solution()

; Le probleme consiste a savoir si il est possible
; de reduire tous les nombres a la valeur 1 en leur
; appliquant 3 regles particulieres

; rule 1 : add 0 at the end of the value
; rule 2 : add 4 at the end of the value
; rule 3 : if the even value, divide by 2

max_val:integer := 10000
val:integer := 4 ;  // 4 est une valeur par defaut, modifiable par main(n)
store(val)  // et c'est tout !!

rule1()	-> (val := val * 10)

rule2()	-> (val := val * 10 + 4)

rule3()	-> (if (( val mod 2 ) = 0) val := val / 2
            else contradiction() )

; exploration meilleure d'abord (meilleure = lowest explored value)
;
; Afin d'eviter les bouclages, on ne reitere pas
; sur les noeuds deja explores...
; valeur[i] = 0 unexplored
; valeur[i] = 1 exploring
; valeur[i] = 2 already explored
valeur[n:integer] : integer := 0

; indice du noeud le plus faible non exploré
lower:integer := 1    

// DEBUG si pas hash table ca marche.
;ancetre[i:(1 .. max_val)] : integer := 0  ; pour depiler la solution
ancetre[n:integer] : integer := 0  ; pour depiler la solution

main(value:integer)
  -> (val := value,
      lower := 1,
      ancetre[value] := -1,
      valeur[value] := 1,
      best_branch(),
      if (val != 1) printf(" \n no solution ")
;      else (printf(" \n valeur = ~S  : \n",value),depile_solution())
     )

best_branch()
 -> (
	 until ((val = 1) | (not(new_node())))
			 ( indente(world?()), printf("\n val:~S  ",val),
       	 let root := val in
	       for i in (1 .. 3)
  	      (try (
        	world+(),
    	        regle(i),
      	        add_node(root),
                world-(),
        	false)
	       catch contradiction (world-(),false))
              )
     )

regle(x:(1 .. 3))
	-> ( case x (
	     {1} rule1(),
	     {2} rule2(),
	     {3} rule3()
                    )
           )

new_node() : boolean
	-> (until ((valeur[lower] = 1) | (lower > max_val))
           	lower :+ 1,
       valeur[lower] := 2,
       val := lower,
       (lower <= max_val)
     )

add_node(root:integer)
 -> (if (val < max_val)
       (if (valeur[val] = 0)
  	   ( if (val < lower) lower := val - 1,
    	     valeur[val] := 1,
      	     ancetre[val] := root,
// DEBUG si printf, on voit le probleme apparaitre entre les noeuds 17 et 34 ..
       	     printf(" genere ~S pere=~S[~S]  [~S]\n",
                 val,root,ancetre[val],ancetre.Core/graph))))

depile_solution()
  -> (let level := 0, value := 1 in
        until (ancetre[value] = -1)
          (level :+ 1,
           indente(level),
            printf(" v = ~S -> ~S\n",value,ancetre[value]),
            value := ancetre[value])
            )

indente(indentation:integer)
	-> (if (indentation != 0)
		(if (indentation > 70) indentation :- 50,
                 printf(" "),
                 indente(indentation - 1)
                )
           )

// record("main(6)",main(6))


// a new bug with 2-dim tables
Dist[x:integer,y:integer] : integer := 0

[naivebug()
  ->  Dist[1,2] := 123,
      check("write 1", Dist[1,2] = 123),
      Dist[tuple(1,2)] := 345,
      check("write 2", Dist[1,2] = 345) ]
 
(naivebug())

// a new bug from SGSS
City <: object(name:string)
(instanced(City))

CY[i:(1 .. 10)] : City := unknown

[testSGSS()
  -> for i in (1 .. 10)
       CY[i] := City(name = "City" /+ string!(i)),
    check("instanced works", size(City) = 10)]

(testSGSS()) 

(testOK())
