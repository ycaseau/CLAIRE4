// Hugarian Matching Algorithm
// credit : https://en.wikipedia.org/wiki/Hungarian_algorithm

// We use the primal-dual algorithm to find the optimal matching described in the Wikipedia article.
// we defined a potential pi on vertices that we grow incrementaly (pi(x) + pi(y) <= c(x,y))
// Gpi is the subgraph of tight edges (c(x,y) = pi(x) + pi(y))
// M is the subset of negative edges in Gpi (from Target to Source)
// Z is the set of vertices that are reachable from the unmatched vertices in Gpi
// note that this is the simpler O(n4) version with no incremental optimization
// this version solves the minimal weight perfect matching problem

TALK:integer :: 1

// Forward definition
Vertex <: object(index:integer)
(instanced(Vertex))
Source <: Vertex
Target <: Vertex

// bi-partite graph G(Source, Target)

Edge <: object(
  start:Source,
  end:Target,
  orientation:boolean = true,    // orientation in Gpi (true from Source to Target, false from Target to Source)
  weight:float)               // c(x,y) 

// constructor
edge(x:Source, y:Target, w:float) : Edge
  -> let e := Edge(start = x, end = y, weight = w) in
       (x.starts :add e,
        y.ends :add e,
        e)

self_print(e:Edge) : void
  -> printf("(~S,~S):~F2",e.start,e.end,e.weight)

Source <: Vertex(       
  pi:float = 0.0,                // potential
  rs?:boolean = true,            // not covered in M (Gpi + orientation)
  rPath:Edge,                    // keep the path that was found when computing Z
  inZ:boolean = false,           // in Z (reachable from unmatched vertices)
  starts:list<Edge>
)

Target <: Vertex(
  pi:float = 0.0,
  rt?:boolean = true,            // not covered in M
  rPath:Edge,                  // keep the path that was found when computing Z
  inZ:boolean = false,           // in Z (reachable from unmatched vertices)
  ends:list<Edge>
)

[self_print(x:Source)
   -> printf("S[~A]",x.index)]
[self_print(y:Target)
   -> printf("T[~A]",y.index)]


// a matching is a list of edges, at most one for each man
Matching :: list<Edge>

// Hungarian algorithm
[Hungarian(G:list<Edge>) : Matching 
  -> let n := computeZ() in
       (while (n > 0)
         (//[TALK] --- step n=~A, M is ~A // n,showM(G),
          when y := some(w in Target | w.rt? & w.inZ)  in
             (//[TALK] found an augmenting path ~A   // showPath(y),
              reversePath(y))
          else growthPotential(),
          n := computeZ()),
        showM(G))]


// tight?(e) 
[tight?(e:Edge) : boolean
  -> e.start.pi + e.end.pi = e.weight]

// match?(e) if e belongs to M
[match?(e:Edge) : boolean
  -> e.orientation = false & tight?(e) = true]

// compute Z and return number of Source in RS (not covered)
[computeZ()
  -> for m in Source m.rs? := forall(e in m.starts | not(match?(e))),
     for w in Target w.rt? := forall(e in w.ends | not(match?(e))),
     let n := 0 in
       (for m in Source m.inZ := false,
        for w in Target w.inZ := false,
        for m in list{m in Source | m.rs?}
             (n :+ 1, propagate(m,m.starts[1])),     // first edge does not matter
       n)
]

// Z is the set of vertices that are reachable from the unmatched vertices in Gpi
// propagation depth first search with memory
// note that we maintain the path in rPath (list of connected edges in Gpi)
[propagate(m:Source,from:Edge)
  -> if not(m.inZ) 
        (m.inZ := true,
         m.rPath := from,
         for e in list{e in m.starts | (tight?(e) & e.orientation)} propagate(e.end,e)) ]

[propagate(w:Target,from:Edge)
  -> if not(w.inZ)
        (w.inZ := true,
         w.rPath := from,
         for e in list{e in w.ends | (tight?(e) & not(e.orientation))} propagate(e.start,e))]

// reverse orientation on a path from Rs to y
[reversePath(y:Target)
  -> let e := y.rPath in
       (e.orientation := not(e.orientation),
        reversePath(e.start))]

[reversePath(x:Source)
  -> if x.rs? nil    // done
     else let e := x.rPath in
        (e.orientation := not(e.orientation),
         reversePath(e.end))]

// grow the potential
[growthPotential()
  -> let delta := 1e9 in
     (for m in {m in Source | m.inZ}
          (for e in m.starts
               (if not(e.end.inZ) (delta :min (e.weight - m.pi - e.end.pi)))),
      //[TALK] increase potential by ~F2 // delta,
      for m in {m in Source | m.inZ} m.pi :+ delta,
      for w in {w in Target | w.inZ} w.pi :- delta)
]

// show the matching
showM(G:list<Edge>) -> list<Edge>{e in G | match?(e)}

// show the weight
showWeight(M:list<Edge>) : float
  -> let w := 0.0 in
       (for e in M w :+ e.weight,
        printf("===> Solution is ~F1 for ~A",w,M))

// show the augmenting path
showPath(w:Target)
  -> (list(w.rPath) /+ showPath(w.rPath.start))

showPath(m:Source)
   -> (if m.rs? list() else list(m.rPath) /+ showPath(m.rPath.end))

// random graph
[rGraph(n:integer) : list<Edge>
  -> for i in (1 .. n) 
        (Source(index = i), Target(index = i)),
     for x in Source
        for y in Target edge(x,y,float!(random(1,100)))]

// Yves wedding problem
[yGraph(n:integer) : list<Edge>
  -> for i in (1 .. n)
        (Source(index = i), Target(index = i)),
     for x in Source
        for y in Target 
          let i := x.index, j := y.index in
           (if (i != j) edge(x,y, float!(ydist(i,j)))) ]


// modulo, always positive (should be fixed in CLAIRE)
pmod :: operation(precedence = 10)
[pmod(i:integer,n:integer) : integer
  -> let m := (i mod n) in
       (if (m >= 0) m else (m + n))]

// yves distance  (mod in Claire)
// the original problem is with 7,5 but 13,7 is also interesting
[ydist(i:integer, y:integer) : integer
  -> sqr((i - y) pmod 7) + sqr((i - y + 3) pmod 5)]



// test Hungarian
test1()
  -> let l := list<Edge>() in
       (for x in Source
         for e in x.starts
           (l :add e),
        showWeight(Hungarian(l)))

// validation through brute force (works only for small graphs)

[minMatch() : float
  -> for y in Target y.inZ := false,
     bestMatch(size(Source)) ]

[bestMatch(i:integer) : float
    -> if (i = 0) 0.0
       else let x := Source.instances[i], v := 1e10 in
              (for e in x.starts
                (if not(e.end.inZ) 
                   (e.end.inZ := true,
                    v :min (e.weight + bestMatch(i - 1)),
                    e.end.inZ := false)),
               v)]

// play with LLM solutions, expressed with list of indices
// GPT5 is correct  
LGPT5 :: list(8, 9, 16, 11, 12, 19, 20, 21,
             2, 3, 24, 5, 6, 7, 22, 23,
             10, 25, 26, 13, 14, 15, 1,
             17, 18, 4)

LMistral :: list{(i + 1) | i in 
               list(7,23,24,25,22,19,20,0,1,2,3,4,5,6,21,8,9,10,11,12,13,14,15,16,17,18)}

// check is simple
[check(l:list[integer])
 -> for y in Target y.inZ := false,
    let w := 0.0 in 
     (for i in (1 .. size(l))
        let x := Source.instances[i], y := Target.instances[l[i]] in
           (if y.inZ error("~S used twice",y),
            y.inZ := true,
            when em := some(e in x.starts | (e.end = y)) in (w :+ em.weight)
            else error("~S-~S not allowed",x,y)),
     printf("matching correct with weight ~F1\n",w))]


