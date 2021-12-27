Thing <: object(name:string,x:integer,y:integer,prev:any)
(ephemeral(Thing))

[clo(k:integer)
  -> time_set(),
     let at := Thing(name = "obj", x = 0, y = 0) in
        (for i in (1 .. k)
            let bt := Thing(name = "obj", x = i, y = i + i) in
               (bt.prev := at, at := bt)),
     time_show() ]

[clo0(k:integer)
  -> time_set(),
     let at := Thing(name = "obj", x = 0, y = 0) in
        (for i in (1 .. k) Thing(name = "obj", x = i, y = i + i)),
     time_show() ]

/*[clo1(k:integer)
  -> time_set(),
     let at := Thing(name = "obj", x = 0, y = 0) in
        (for i in (1 .. k) Thing()),
     time_show() ]*/

// read the instance variables - 1M passes in  0.21s
[clr(n:integer)
  -> time_set(),
     let at := Thing(name = "obj", x = 1, y = 2), val := 0 in
        (for k in (1 .. n)
            val := (at.x + at.y * (at.x + at.y)) * (at.y - at.x)),
      time_show() ]

// write the instance variables
[clw(n:integer)
  -> time_set(),
     let at := Thing(name = "obj", x = 1, y = 2) in
        for k in (1 .. n)
             (at.x := n,
              at.y := n,
              at.x := at.y + n),
      time_show() ]

/* write the instance variables
[clw1(n:integer)
  -> time_set(),
     let at := Thing(name = "obj", x = 1, y = 2) in
        for k in (1 .. n)
          write_fast(x,at,n),
      time_show() ] */


[all() 
  -> princ("clo(1M)"),
     clo(1000000),
     princ("clr(1M)"),
     clr(1000000),
     princ("clw(1M)"),
     clw(1000000),
     nil ]      

[all2() 
  -> princ("clo(10M)"),
     clo(10000000),
     princ("clr(10M)"),
     clr(10000000),
     princ("clw(10M)"),
     clw(10000000),
     nil ]  

[main() -> 
  //[0] test Obj ----- //,
  all(),
  all2(),
  exit(0) ]    

