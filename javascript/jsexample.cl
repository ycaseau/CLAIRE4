//+-------------------------------------------------------------+
//| Diet CLAIRE                                                 |
//| compiler examples (Javascript code generation)              |
//+-------------------------------------------------------------+


this file contains javascript fragments that will be used as templates.
it replaces the written notebook (that will be lost eventually)

This is "Diet Claire Compiler" (DCC) that generates javascript code from CLAIRE code.
Diet means that only simple constructs are allowed (no calls, no rules, no complex types)

// *******************************************************************
// * Contents                                                        *
// *     Part 1: Modules                                             *
// *     Part 2: Classes                                             *
// *     Part 3: Methods                                             *
// *     Part 4: Primitive sorts                                     *
// *     Part 5: Objects                                             *
// *     Part 6: Expressions                                         *
// *     Part 7: Statements                                          *
// *     Part 8: Error Handling                                      *
// *******************************************************************
 

// *******************************************************************
// *     Part 1: Modules                                             *
// *******************************************************************

a claire module => a set of JavaScript files (no namespace is supported)

m becomes m.js
with require('./ClaireKernel.js')
  
All named objects (things, including classes and global variables) are associated a go global variable
  object o in module m (module!(o.name)) =>   var m_o *ClaireC
  

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



// *******************************************************************
// *     Part 2: Classes                                             *
// *******************************************************************

// there is no reflection in Diet Claire,but we want to have meta classes: 
//     supports membership
//     supports class iteration
// se we follow the same code generation as Go

// a class definition produces
C_class = MakeClass("name",superclass)    

// slots and properties are not reified -> AddSlot and AddMethod are not needed

// class definition in Javascript
// simple since no slots need to be defined
// a cool idea could be to allow dynamic calls when methods are dispatched
// (same signature, but only the class changes)
class c1 extends c2 {
        constructor(a) {
                this.isa = C_c1
                this.slot = default_value
                c1_class.instances.push(this)
                }}

// instantiation
new!(c)    -> new C()

// *******************************************************************
// *     Part 3: Methods                                             *
// *******************************************************************

// we borrow Go distinction between methods and functions
jfmethod(m:method)
   -> all respections of m.selector have 
       - a class from the current module as the domain
       - the same codomain

// methods to function code generation 
two cases :
(a) module method : method defined in the same module as the domain class
  -> native is capitalized(m)  =>  x.M(...)

  func (a1 *ClaireA1) M (a1 type1, a2 type2 ....)

(b) external method : method added to a class
  -> native is F_m_c
  -> interpreted is E_m_c

  func m_c( a1 type1, a2 type2 ....)

Reference on how to add a method:
https://stackoverflow.com/questions/68020041/how-do-you-declare-a-javascript-class-method-outside-of-the-class
// *******************************************************************
// *     Part 4: Primitive sorts                                     *
// *******************************************************************

Class.prototype.selector := function() {  } 

// *******************************************************************
// *     Part 4: Primitive sorts                                     *
// *******************************************************************

// no typing in Javascript, so we do not need sorts

Primitive types translation
integer -> int
float -> double
char ->  char
string -> string
Port -> not supported yet (no i/o in Diet)
function -> not DIET
boolean -> true and false     
list -> arrays
set -> Set  (new Set())
map_set -> Object
  (this is the fun part : all JS objects are implemented as a map_set !)

// *******************************************************************
// *     Part 5: Objects & Lists                                     *
// *******************************************************************

Objects in Javascript are organized into classes


Global variables => ClaireObjects with no native mode (except for global constants)
a global variable is a thing (with a name)
X:integer := 12
package.X = ClaireGlobalVariable(12)
use as : package.X.Value 

// bags are not diet : we need explicit distinction between
- list, arrayn, tupple -> Javascript arrays
- set -> Javascript Set()


Access to lists
  l[i]    =>   l[i - 1] in native form 
list iteration 
- L.foreach(function(x) {e(x)})
- for (let i = 0; i < L.length; i++) { x = L[i]; e(x)}

// sets are native in JavaScript
s = new Set()
s.add(i)
for (const x) in s ...

// *******************************************************************
// *     Part 6: Expressions                                         *
// *******************************************************************


MakeX(t,a1,...an)         // no type for tuple
MakeEmptyX(t)             // not for tuple


// Calls  

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
r(x) := y    =>    r(x) = y only if the value is not needed
otherwise    var_val = y
             r(x) = var_val
             v = var_val
Also, in diet Claire, we manage inverses explicitly
r(x) := y  ->   x.r = y
                y.r2 = x

// *******************************************************************
// *     Part 8: Error Handling                                      *
// *******************************************************************

error("pattern", arg*)
is compiled into
- MakeError(pattern,list)
- Close() a method of ClaireClass

ClaireKernel.js implements ClaireError
- with a method Close()
- with a method Format(string,list)
