// +------------------------------------------------------------+
// | bstub.cl                                                   |
// | last update: July 2021 - Y. Caseau                         |
// +------------------------------------------------------------+


// ---------------------------------------------------------------
// this file contains utilities
// ---------------------------------------------------------------

CompiledTest:boolean :: (#if (get_value("compile") != unknown) Id(compiler.active?) else false)
ClaireVersion:string :: Id(string!(integer!(version() * 100.0001)))

TestPort:port :: unknown
ModuleName:string :: ""

// check = assert with a symbolic tag + not sensible to compiler options
// note: we do not print when the check succeeds to keep the log small
[check(tag:string, x:any)
 -> (if not(x)
              (//[0] check ~A failed // tag,
               record("*** check failure *** " /+ tag, x)))]


// record -> write a value on the log so that difference analysis can be made
[record(tag:string,x:any)
 ->  let p := use_as_output(TestPort) in
           (printf("[~A] ~A : ~S\n",ModuleName, tag, x),
            use_as_output(p)) ]

// cute function that works with or without compiler
[exit_end() : boolean
  -> (#if (get_value("compile") != unknown) 
       Id(not(compiler.active?) | compiler.loading?)
      else true) ]

// close
[testOK(n:integer) : void
 -> printf("----------- end test --------------------\n"),
    record("OK:",true),
    fclose(TestPort),
    (if exit_end() exit(0) else nil),
    nil]

[testOK() => testOK(0) ]

// useful to have variants with the compiler -
//[compile_on() : boolean
//  -> (#if (value("compile") != unknown) Id(compiler.active?)
//      else false) ]

// open the test result port
// sets the compiler safety to the right level 
[startTest()
 -> ModuleName := string!(name(module!())),
    TestPort := fopen((if CompiledTest "clog" else "log") /+ ClaireVersion,"a"),
     printf("----------- start test for ~A with release 4.~A\n",  ModuleName, ClaireVersion) ]


(startTest())
