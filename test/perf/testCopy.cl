// ========================  Step7: file copy ===========================================================


F1 :: "/Users/ycaseau/proj/python/test1.py"
F2 :: "/Users/ycaseau/proj/python/test2.py"


// copy
[copy(f1:string,f2:string)
 -> let p1 := fopen(f1,"r"), m := 0,
        p2 := fopen(f2,"w") in
          (while true
             let c := getc(p1) in
               (m :+ 1,
                // printf("[~A] read char ~S(~A) vs ~S(~A)\n",m,c,integer!(c),EOF,-1),  
                if (c = EOF) break(),
                putc(c,p2)),
           fclose(p2)) ]

// test: 200 copies in 1.53s
[tcp()
  -> time_set(),
     for i in (1 .. 100)
         (copy(F1,F2),
          copy(F2,F1)),
     time_show() ]

all() -> tcp()

[tcp0() -> copy("testCopy.cl","junkFile")]


