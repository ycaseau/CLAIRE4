// *******************************************************************************************
// *     log file : simplified on July 29th 2024 (see previous versions for full history)    *
// *******************************************************************************************

// THIS LOG FILE sits in the claire/dropbox/compile directory : how we create CLAIRE 2 (our first interpreter)
// from Kernel go files and compiled CLAIRE files (thanks to 3.5 + trans)
// the first step is trans !

shared context:
- developped on Mac / will need to synchronize on cloud to save
- use Visual Studio Editor



// move to CLAIRE v4.0.7 on June 20th 2022 ------------------------------------------------------------------

// resume for Xmas vacacations 2022 ! 
-  add m.resources: list of string that represent useful files (great to copy/upload module)
-  add two gitUpload and gitDownload methods  ? unclear (more like fragments)

// first step : recompile Claire4 using the new file directories

// test 2 key modules (load, compile, compare output with CLAIRE 3.5)
(a) MMS
- load (m3 vs m2 in Claire3)
- compile : error ! forward declarations do not work
BIG CHANGE : introduce NewClass(name,c,m) which checks if the class is already there (support forward def)
-> compiled(m3) works

(b) EMLA
- load (m5)
added ephemeral_object for upward_compatibility in object.cl
bug in gauge.cl => line 329 : a piece of code yields strange "enumerate primitive"
min.open was 2 in claire3.5 -> reset in 4.0
[foo(x:list<int>) -> 12] does not complain (list<X> may be parametric)

load(m5) fails .. nth error (inside code)
=> big error in nth+ (ClBag.go)  -> add the empty case nth+((),1,X)
Aha ! characters in CLAIRE4 requires their equality tests (only a small part is cashed)
=> now load(m5) works

// Dec 28th
compile(m5) fails with two problems ....
(1) soft: fractional @ integer (use a / 10) produces status = 1 (while code is safe)
    need to understand

(2) hard: 
g_test(quote(for x in LIST_LEVEL print(x)))  fails with a member_type
car pmember({list<float>(0.2, 0.3, 0.5, 0.7)}) fails !!
=> resolved (big bug in ClBag : set!(list<float>) ... used s := ... vs =

(3) Attention: slots cannot have the same name as methods in Go.
    thus the "simple" goMethod?(m) must check that no other restrictions is a slot (with intersecting domain)

(4) true compiler error with use of indirect read(p) where p = list[i]

// Jan 1st : MMS & EMLA OK + tests and perf

// when we restart : copy sclaire in a safe place and recompile everything 

// ============= CLAIRE 4.08 =====================================================

- osystem.cl : error checking for division only if compiler.safety < 5
- makefile: added "make cross" to be used before "make"

Decision
(1) always run cross-compilation and fixed point
(2) odd version numbers are private ... even versions are public

// ============= CLAIRE 4.10 =====================================================

create v4.09 then move to v4.10 when tested

- extend the size of possible classes to 50 (simple + compatible with SGSS)
   -> ClReflect: add a test and an error (fatal error ? )
   -> ClKernel: create, read and set objet slots

- create reboot()
  (a) in claire1.go : add a method E_reboot() that calls Bootstrap() and Load()
  (b) add ToMethod(C_reboot.Restrictions.ValuesO()[0]).Functional = MakeFunction0(E_reboot,"reboot")
  (c) tune compiler to add the previous line (gomain.cl)
  (d) add reboot() in ClReflect.go
  (e) make top-level independant from local variable (to avoid r:reader becoming obsolete)

- load SGSS as a test

   (1) added random! as defined in the documentations
   (2) renamed stat() into statistics()  ( short cuts and abbreviation are not supported) 

- load FBID as a test

- compile SGSS & FBID : fixed a number of compiler bugs :)

September 2nd
- add m.resources: list of string that represent useful files (great to copy/upload module)
  it works AS LONG AS THE SLOT is put in the right position
  it was also the opportunity to put import where it should be 
- now supports creating the executable on another directory
   claire4 -od <dir> -cm <module>
  this allowed to clean the test generation to avoid the mess in go directory
- fixed stupid bug with value@symbol that called .value vs .Value()

TODO on PC

============

(1) move to 4.10
(2) publier CLAIRE v4.10 sur GitHub
(3) test on PC !!!
(4) modifier https://sites.google.com/view/claire4/home
(5) update the CLAIRE programming language page !!
(6) update Facebook page
(7) publier SGSS et FBID sur github

// October 22nd, 2023: This is the day when CLAIRE 4.10 is published on GitHub  !!!!!!!!!!!!!!!!

- fixed the exception!() bug by adding a symmetrical method for glb@Pattern
- fixed the SHOW := true bug by prefixing wrong in pretty.cl with Core/
- updated to gitHub

// February 1st : open CLAIRE 4.1.11 -> publish 4.1.12 in April

// introduced Language/occurbreak(self:any)
allows to compile "for x in c ..." in a better way (without checking for return)

// TODO
(a) bugs from Cl2js (look at the file)
(b) redefinition of a method should complain (but not an error, useful when debug)

// ============= CLAIRE 4.12 =====================================================

create v4.1.1 then move to v4.1.2 when tested

// [A]  Bug fixes 
    (3a): done length_bag does not exist -> should be length_list
        -> verifier toutes les constantes de gosystem.cl - DONE

    (3b) affine(l:listarg)   -> les calls sont mal compilés 
     => correction dans 4.12 [ocall.cl]
      - restriction!() renvoie la methode m =p@(.. listargs) 
      - grace a dmatch? (qui était non utilisée !!!)
      - et c_code_call traite le cas d'une listargs avec lisrargsFormat(l,n)

    (3c) CRAZY BUG in CLAIRE4 to fix  foo(1,,2)
       it looks like the eval of call(foo,list(1,,,3)) creates a bug with ','
        => fixed in 4.12: there was a next(c) in self_eval(delimiter) that was not needed [file.cl]
    
    (3d) bug with sum( ...) : produce a for with a break with no reasons
    see ActualEnergy in gw0.js => seems to work with go
    
    (3d) poor typing for (class but class) because second-order typing of (X but Y) is not supported
    => look at but-compile :)
         Iterate(but[..] )  works well (ocontrol.cl)
         but second-order type was false (should be type[self])

    BUG fix in [ocontrol.c] : c_code_select
       -> removed a path that was not needed (and was wrong) : %t vs x (infered type) ... better typing of but 
       made the code less precise.
    
    (3e) doublon entre get@symbol et value@symbol -> get rid of get@symbol / DONE


// [B] what to do in this release
(a)   ??? introduce //(test)  code   [read.cl]
      - it must be one line, hence code is a short printf or a function call 
      - we use () as syntactic markers but we count for balancing parenthesis
      -> cf conditional_comment <!|?> in read.cl
      we use "assert" as a marker (else part) to recognize conditional comments, we should create If?

(b) create option -e "exp" that will allow to work faster [ccmain.cl]
          claire -m ... -e "jcompile(gw0)" 
          ajouter le code dans gomain ... attention exp doit etre un func call sans whitespace

(c) use € versus % for belong
    change the name of the property and create an alias % -> €
    (a) create the property with name € [ClReflect] and create alias for "%"
    (b) extend c_princ_char in [ClString.go] so that % = €

(d) dual definition of global constant should complain ? 
    toto :: 12 -> creates Defobj, with  arg = global_variable args = list(range = {}, value = 12)
    and iClaire/ident = claire/"toto"
    todo : add a warning if self_eval@Defobj when c = global variable (no redefinition of global var expected)

// [C] compiler architecture  (cf analysis made in cl2js-log.cl)

(1) [gosystem.cl] create generic methods for code_producer
(2) for go compiler, leverage generic methods
(3) create init.cl stub  -> done
(4) compile (call this version 4.12)
    Aha ! a "]" left in the code creates an infinite loop .... (FIX LATER)
    NOTE: How to do a diff = compare file with 
    (a) select the command palette in view
    (b) type "Compare Active File with ..."
    (c) click on the desired file (visible in the list if open)
(5)  check that fixed point has been reached  => done on 8/10/2024 at 7:55AM

//[D] update the javascript compiler using the new architecture
//    (a) create a new version of cl2js
//    (b) test the new version with gw5.js
//    done on 8/8 :)

// [E] before releasing 4.1.2 (August)
- run perf tests -> OK
- update the documentation
     -> check that #if is documented
     -> add //(..) 
     -> check that //[?] expression
- upload on GitHub : cf GitNotes
    (a) make git (on root/go)
    (b) gitAdd (on root)
    (c) git commit -m $(COMMIT)
	  (d) git push -f origin master
- update website
   

(b) talk about new 16 bits char in the claire documentation

// =============== CLAIRE 4.1.3 (will be 4.1.4 in December) ========================

Octobre 31st (Halloween)
- define.cl => check that methods have less than 10 arguments
- goexp.cl -> when compiling a lambda, use print_in_string to get its body as as string
- bug5.cl  -> add a test with (x){print("hello")}  

January 1st, 2025

KEY= moved directory to 4.1 to simplify (the go1, go2, ... subdirectories are no longer necessary)
THIS REQUIRES TO CHANGE GO PATH TO ~/claire/v4.1/go
we find the subdirecties that will be pushed to GitHub:
- src
- test
- meta
- compile
- interpreter

As before, the work is done in the go directory

January 3rd: publish 4.1.4 on GitHub x PC x Ubuntu


A year of reading about energy/economy/climate scenarios, as well as doing my own simulations with the CCEM world modem, has led me the following four insights:

(1) Borrowing from Hannah Richie’s great book, what most 21st century simulations show is “not the end of the world”.
 The state of the world today is more what could be called “slow transformation as usual” (versus "Business as usual" or versus Paris Agreement targets): 
 the energy transition has already begun, focusing on clean energy adoption, efficiency improvements, and air quality advancements, but progress is too slow.
 
(2) Decarbonation of energy will happen, but there seems to be a critical time gap between the peak of fossil fuel reliance and the widespread availability 
of clean energy at sufficient levels. Such a gap will lead to economic strain caused by energy shortages while the quantity of fossil fuels 
connsumed will cause global temperatures to rise by 2.5°C to 3.0°C by 2100. 

(3) This leads to the idea that it may not be the end of the world, but the end of our 20th century world is definitely coming 
  (as is beautifully explained by Langdon Morris book "Hello Future!"). Sustained global warming with severe impacts necessitates adaptation,
   while the era of cheap energy has ended, leading to energy shortages, market distortions, and regional disparities intensified by conflicts. 

(4)  For Europe, and this aligns perfectly with Mario Draghi’s report, we need to act quickly on an aggressive decarbonation strategy, 
     not to save the planet, but to save our economic competitiveness. 
    
You can read abpout those on my last 2024 blog post: https://informationsystemsbiology.blogspot.com/2024/12/energy-matters-modeling-manifesto.html









// ============================ OLD BACKLOG =====================================

// New : add a reboot function in CLAIRE (useful to play with server)
// reboot() : recreate all objects (rebuild a clean state)

- ajouter la ligne de code dans le code généré  pour claire1.go
    (a) dans Load (à la fin)
    C_reboot.AddMethod(Signature(C_void.Id(), C_void.Id()), 0, MakeFunction1(E_reboot_void, "reboot_void"))
    (b) deux méthodes associées
    func reboot_void() {
      Bootstrap()
      Load()
      Reader.C_reader.Fromp = ClEnv.Cin
    }
    func E_reboot_void(s EID) EID {
      reboot_void()
      return EVOID
    }
	
// enrich the Go unit tests ! 
// create a nice command to launch tests from the makefile
list: 
   -  nth+(list<integer>(),1,1) = list<integer>(1)
   - set!(list<float>(1.2)) = {1.2}
test Equal('a','a'), Equal('a','b'), 


// extensions

- play with spy
- imports is not used !!!

- add something for methods with no errors but poor range ...
     simplest is when we test can_throw / check that c_type <= range.  
- Introduce CheckRange(type,Result,"cause")
- add the "close world assumption" for a Call in g_throw => restriction! is not empty
-  the c_substitution does not work with jito on !!!!! find out why 
      try g_test(for i in (1 .. (2 + 3)) print(i))
- fix the compilation of self_print @ string (self_print is open) => in gexp
  optimize print(integer) ?
- why do we set compiler.inline? to false, it should be true !
- when we detect a discrepancy between actual throw and can_throw => 
    a. we should get rid of "good/bad lists" (unless we keep it as a debug back door)
    b. we should generate the code that is expected by the status
    c. a proper warning should be issued
- add a second-order type to check_in

=> Réfléchir à une méthode pour automatiser des tests Claire via Go
    Claire -ct mTests  -> run tests via un go test ...
    Suppose de bien comprendre le monde Go 





