//+-------------------------------------------------------------+
//| CLAIRE                                                      |
//| compiler examples (Go code generation)                      |
//+-------------------------------------------------------------+


this file contains go fragments that will be used as templates.
it replaces the written notebook (that will be lost eventually)

// *******************************************************************
// * Contents                                                        *
// *     Part 1: Modules                                             *
// *     Part 2: Classes                                             *
// *     Part 3: Methods                                             *
// *     Part 4: Primitive sorts                                     *
// *     Part 5: Objects                                             *
// *     Part 6: Expressions                                         *
// *     Part 7: Statements                                          *
// *     Part 8: Miscellaneous                                       *
// *******************************************************************
 

// *******************************************************************
// *     Part 1: Modules                                             *
// *******************************************************************

a claire module => a ClaireModule object + a package
  AXIOM : a claire module m is a Go package only if m.made_of is not nil (+ Kernel)


All named objects (things, including classes and global variables) are associated a go global variable
  object o in module m (module!(o.name)) =>   var m_o *ClaireC
  if m is Kernel, use short cut const_cast

Table 
  
  +-----------------+------------------------------+-------------------------------+
  |  4 kinds        |     defined in Kernel        |     defined in m              |
  +-----------------+------------------------------+-------------------------------+
  |  Class          |                              |                               |
  |    definition   |      ClaireC                 |     C (capitalized)           |
  |    use          |      ClaireC  (since *)      |     m.C                       |
  |    cast         |      ToC(...)                |     m.ToC(...)                |
  +-----------------+------------------------------+-------------------------------+
  |  Thing          |     C_thing                  |     C_thing                   |
  |    use          |     C_thing                  |     m.C_thing                 |
  +-----------------+------------------------------+-------------------------------+
  | Method foo      |      Foo()                   |     Foo()                     |
  |    use          |      ClaireC  (since *)      |     m.Foo()                   |
  +-----------------+------------------------------+-------------------------------+
  | function f_c    |      F_f_c                   |     F_f_c()                   |
  |    use          |      F_f_c  (since *)        |     m.F_f_c()                 |
  +-----------------+------------------------------+-------------------------------+

declarations
    package m  :  the rest of the file is in m
    import(   _ "m")    : uses m, _ avoids unused errors, note that we do not ? need the path, e.g. Compile/Generate
    import(   . "Kernel")  : avoids the Kernel.XXX

// modules are things (named objets) but they cannot be referenced recursively hence we use the previous IT trick
m.it ->  claire object with the module m 
However, Claire is not a go package, so we use C_claire (exception)

// what happens if x.name.defined=m2 is not x.name.module! = m1 (ex Claire) ?
Nothing ! go does not know about modules

// *******************************************************************
// *     Part 2: Classes                                             *
// *******************************************************************

// a class definition produces
C_class = MakeClass("name",superclass,module)     // module is where the class is visible Mod = tag(m) or C if m = claire
-- note : C_class is visible in the module/package, elsewhere it is mod.C_class
go_class(c) = ModClass,  class_ident = <mod.>C_class
examples : ClaireClass OptimizeCallSpecial => Kernel.C_class, Optimize.C_Call_special

because of forward, we need NewClass("name",superclass,module) which looks for a class with that name first

// adding a slot
func (c *ClaireClass) AddSlot(p *ClaireProperty, r *ClaireType, def *ClaireAny) EID {
    note that we do not handle the index 
     aha : do we want to manage the messy error case (multi-valued conflict) ?  => tricky

// class definition in Go
type Claire<X> struct {
	Claire<Y>                     // X is a subclass of Y
	Tag(s)      *Claire<R>        // s is a slot, Tag(s) is cident(s.selector.name)
	Num         int               // native form form for integer and floats, not for strings and chars
	Slots       *ClaireList       // list of slots (no prototype, use slots)
  buddy       *Claire<X>        // typical object slot
  friends     *ClaireSet        // multivalued slot

// instantiation
new!(c)    -> ToC(new(ClaireC).Is(M.C_c))


total instanciation is done with Make
func (c *ClaireClass) Make (x ClaireAny ...)
this is used with the anyObject! pattern (optimizer) => used for exceptions, should be used for instruction as well ?
Also : total instanciation can use go-genrerated-constructors
   Point(x := 1, y := 2)  -> MakePoint(1,2)

// *******************************************************************
// *     Part 3: Methods                                             *
// *******************************************************************

// properties
properties are created with MakeProperty( ....)   => defined in ClReflect.cl (missing at the moment)
     MakeProperty(name string, open int, module m, domain *ClaireClass)

In go (CLAIRE4) we got rid of "fast dispatch" with the dispatcher structure (a list of functions + a class)


// methods
methods are created with AddMethod(Signature, status, function) status tells about error
    func (p *ClaireProperty) AddMethod(f eMeth1, ld []*ClaireAny, err *ClaireBoolean) {
    example : C_copy.AddMethod1(Signature(Kernel.C_list.toType(), Kernel.C_list.toType()), 0, MakeFunction1(F_copy_list,"F_copy_list"))
              C.foo.AddMethod2(Signature(Kernel.C_integer.toType(),Kernel.C_integer,F_dot_dot(0,10)), 1, MakeFunction2(E_foo_integer,,"E_foo_integer")

the arity is 1 to 10 => 10 methods AddMethod<i> where f is eMeth<i>


// methods to function code generation 
two cases :
(a) module method : method defined in the same module as the domain class
  -> native is capitalized(m)  =>  x.M(...)
  -> interpreted is E_m_c

  func (a1 *ClaireA1) M (a1 type1, a2 type2 ....)

(b) external method : method added to a class
  -> native is F_m_c
  -> interpreted is E_m_c

  func m_c( a1 type1, a2 type2 ....)

// design decision : do we add the module/package name ?  (when the property foo is not claire/foo)
// yes F_M1_foo_integer  and F_M2_foo_integer may coexist

// key : the Functional (EID) go function is associated to the method in the "Functional" slot. The other is computed
// dynamically by the compiler as a string.

// type checking : if not statically type-checked, EID compiling
func foo() integer {
  var Result EID
  ...
  return RangeCheck(Result,C_integer) }
}

// void functions returns unknown under the interpreter 
EVOID = EID{C_void, CNULL}  where CNULL is a *ClaireAny
-> special pattern for the EID function
   E_m_c (....)  {  m(....)  return EVOID}


// *******************************************************************
// *     Part 4: Primitive sorts                                     *
// *******************************************************************

4 sorts (exception and object are similar - but flagging exception improved exception handling)

sort            go for object          Native         to_go           to_cl  
any             *ClaireAny              *ClaireAny      
exception       *ClaireAny              *ClaireAny      
int             *ClaireInt               int            x.Value       MakeInteger
float           *ClaireFloatr           float64         x.Value       MakeFloat
char            *ClaireChar             rune   
      
here the slots  are stored in a native form + the arguments are passed in a native form

// we also have primitive objects : Port, String, Functions, Booleans
slots are stored with ClaireX objects, they are passed the sames way

// THIS IS THE WAY FOR UPDATABLE strings !!
s[3] := 'F' is allowed and supported in claire

// boolean are implemented with object (ClaireBoolean) but the compiler also has a native mode
we use the following
   CBoolean(b:boolean)
   IfThenElse(b:boolean,a:ClaireAny,b:ClaireAny)

 EID in CLAIRE 4 are pairs {one *ClaireAny, One Value}
 5 kinds of EID
     Object x  -> EID{x,0}
     Error  x  -> EID{x,1}
     Integer x -> EID{c_INT, x}
     Float x -> EID{c_float, x}
     char r -> EID{C_char, r}  

we use 3 markers that are special and unique ClaireThings

// *******************************************************************
// *     Part 5: Objects & Lists                                     *
// *******************************************************************

Objects in Go are struct with embedded inheritance (sugar syntax)
the unknown object is represented by CNULL - assuming that the cast of CNULL to another struct does not 
break the GO garbage collector - WARNING: OTHERWISE WE NEED TO IMPLEMENT CNULL = ToAny(0=nil)
Slots are struct variables
- always 64 bits, 3 sorts : obj, int and float64
- direct access is made possible in ClKernel thanks to Dummy classes :)


Global variables => ClaireObjects with no native mode (except for global constants)
a global variable is a thing (with a name)
X:integer := 12
package.X = ClaireGlobalVariable(12)
use as : package.X.Value 


Bags: we separate lists (and tuple : of = nil + arrays : fixed size) and sets but we use commons
Bag is is a common for Array,List,Tuple,Set
  bag<t>(a1,a2, ... an) =>  Make<bag>(t,a1, ... an) for expressions
     =>  v = t.Empty<bag>() 
         v.AddFast(a1) ...
  bag(a1,a2,a3)  =>   Make<bag>(void,a1,... an)  ;  for list, tuple or set  ()
     =>  v = MakeNil<bag>(t)                     
         v.AddFast(a1) ...

KEY: (to enforce) list(1,2,3) and tuple(1,2,3) are the same  => suggest that list(1,2,3) returns tuple(1,2,3)


Access to lists
  l[i]    =>   l.Value[i - 1] in native form  or l.At(i - 1) if l is generic (we do not know the sort of the content)
               the test is g_member(l) != any

Arrays are simply lists (with the three kinds)  = this is a key decision ! by refusing to create ClaireArray
as a golang type, we avoid duplication in the go code but requires the CLAIRE compiler to be smarter.

// Note: DESIGN decision (made, to enforce)
(1) we have said that from CLAIRE, we only know ClaireList  (the subvariants are managed by Go)
(2) the compiler uses native forms for l<t> only when t is specific enough (t = int float or ob)
(3) WE DO NOT generate specific list types for slots 
(4) multivalued slots => use of sets, which is only generic :)

// *******************************************************************
// *     Part 6: Expressions                                         *
// *******************************************************************


MakeX(t,a1,...an)         // no type for tuple
MakeEmptyX(t)             // not for tuple


// Calls 

ATTENTION : calls cannot be expressions any more, if an error may be returned
(1) we need to compute this feature for all methods : canThrow? - it must be stored for each method ! (new slot)
    => REUSE STATUS !  ()  - status is bitvector, add a new bit
(2) a call_method with a throwing method => requires a statement
(3) for a call : better (code produces EID) but check the parameters IF 

compiled call: (Call_method)  => look at print_external_call in gexp.

dynamic calls:  relies on
     func (p *ClaireProperty) Call(args ...EID) EID {
  but also uses the fastcall pattern  (with dispatcher)
We can force the type of the first argument with Super :   p.Super(integer,a,b)


  func (p *ClaireProperty) Fastcall(args ...EID) EID {
     that is based on dispatch index (c.dispatch[m]).Call(args)


// Call slots

simple case : native or default value (WARNING : in CLAIRE 4 , native slots have default values)
x.c   =>   x.c
complex case (test = true)
x.c  =>   Known(c,x.c).ToY()   where Y is the range of the slot
warning: the complex case implies a statement ! 

// Call Tables   (cf. object.cl + define.cl : Defarray)
a claire 4 table has two slots
- param : an integer for 1-dim integer array, list<int> for 2-dim array, any for dictionary
- graph : a list or a map (type any) - map is handled via graph_get and grah_put (+ graph_init)



// *******************************************************************
// *     Part 7: Statements                                          *
// *******************************************************************

// indentation : 
//    we call statement(s) at the proper current indentation level => it produices n lines with the indentation
//    and stop after a break line, at the proper identation level

// blocks are opened and closed with new block (create a {, +Indent, nl}) and close block (}, -Indent, nl)


Let define local variables with three sorts
Let x := f in e   =>   var x <sort>     (use g_var_sort ?)
                       x = f            (statement)
                       e                (statement)
there are five sorts here : object, int, float, char, string
variant if g_throw(f)

Let+ => temporary assignment
Let* => variable tuple assignment

-------------------------------------
If statement : If create block
if (expression) {               => new_block
  statement(arg)
  }                              => close block
or
  } else {                       => pattern
  statement(other)
  }

Nested ifs (only with functional) :
 }  else if (...)

 When the test is not functional (or with a throw)
  

------------------------------------------

While (test) {stuff}  is now tricky in case there is an exception.
    g_statement(stuff,v)  produces an assignment to v if an error is possible (implies that expected is EID)

example with full error 
   var v_while *ClaireBoolean
   var v_try EID
   v_try = test
   if ErrorIN(v_try) {v = v_try
   } else { v_while = OBJ(v_try).ToBoolean()
        while (v_while == CTRUE) {
          v = stuff
          if ErrorIN(v) {
            break
          } else {var v_try EID
                  v_try = test
                  if ErrorIn(v_try) {
                      v = v_try
                      break
                  } else {
                    v_while = OBJ(v_try).ToBoolean()
                  }
          }
          }
        }
   }


KEY: While and For may have break(val) inside, which means that the type of the return value must be
passed in the loop parameter (a cls))

-----------------------------------------------------------

iteration : the optimizer does the hard work!
iteration for lists have two forms:
   - native (we know the type of x in go : int, float or object for lists + alsways for sets)
         for x in S e   => for _,x := range(S.value) e
   - generic (we do not know w)
         for x in S e =>   for i,_ := range(S.value) {x := S.}

Collect (f(x) | x in S) uses an expansion
   v_list = S
   v = CreateList(t,v_list.Length()).To<Fuck>()      // if v:Fuck
   for i := 0; i < v_list.Length(); i++ {
       x = v_list.At(i)
       y = f(x)
       v.ToList()?.PutAt(i,y)}
   }    // with the additional protection of two possible try/catch 



------------------------------------------------

Error Handling : smart chain pattern (notice that the return will happen later)
try e catch S Y    =>   v is of type EID
v = e
vcatch = S.contains(OBJ(v))
if vcatch { v = Y }
    
------------------------------------------------------

read a slot x.s => two cases
    not err :      v1 = x
                   v = v1.x
    err:           v_try = x
                   if ErrorIn(v_try) {v = vtry
                   } else {
                    v = OBJ(v_try).s

-----------------------------------------------------

updates are no longer expressions !
then r(x) = y    =>    r(x) = y only if the value is not needed
otherwise    var_val = y
             r(x) = var_val
             v = var_val
with all the complexity for error handling.

// *******************************************************************
// *     Part 8: Miscellaneous                                       *
// *******************************************************************

profiler : PRcount objects are defined in CLAIRE (inspect.cl)
start a method :
 var PR_x *ClairePRcount  PRget_property(~I).Start()
end a method
  PRend(PR_x)        (dans gogen.cl => return_result(...) anotates the return )




Debug: add to the generated method m(a1, ... an) entry and exit code
  entry:  DebugBind(mod,p,a1,...an)
  exit:  DebugUnbind(....)




Worlds:
o.a := x   =>   o.StoreObj/Int/Float(n,x,CTRUE)  (specialized version)
o.StoreAny(n,x,range,bool) is the generic version

lists:   store(l,n,val, bool)      - the control bool is for the pattern store(..., p.store?)
arrays are lists in CLAIRE 4 from an implementation perspective, but a separate class
store(l,n,val)  ->  l.Store(n int, y *ClaireAny, b *ClaireBoolean) *ClaireAny {
in the future we could see if we do specialized stores

Operations: world_push() : create a new world, world_pop(): return (defeasible updates) to the previous 
world, world_remove : commit the hypothetical changes and return to the previous world

 