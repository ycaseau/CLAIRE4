//+-------------------------------------------------------------+
//| CLAIRE                                                      |
//| osystem.cl                                                  |
//| Copyright (C) 1994 - 2021 Yves Caseau. All Rights Reserved  |
//| cf. copyright info in file object.cl: about()               |
//+-------------------------------------------------------------+

//-------------------------------------------------------------------
// This file contains the gloabal parameter objects and the key methods
// This are the key methods
//
//  c_type(x)  is the CLAIRE type of x
//  c_code(x)  is an optimized instruction
//  & c_code(x,s) is an optimized expression of sort s
//  c_sort(x)  the sort of the expression x (precise sort)
//  g_throw(x) boolean that says if x may throw an exception
//
// the sorts are integer, float, char, object, EID
//-----------------------------------------------------------------

// ******************************************************************
// *   Table of contents                                            *
// *    Part 1: General Global Variables and Properties             *
// *    Part 2: The defaults for c_type, c_code, c_gc and c_sort    *
// *    Part 3: g_throw and status(m:method)                        *
// *    Part 4: Names & identifiers management                      *
// ******************************************************************

/* membership optimization though inline definition of %
%(x:any,y:..[tuple(any, any)]) : boolean
 => (x <= eval(y.args[2]) & eval(y.args[1]) <= x) */

// import
// Compile/index :: Kernel/index   (1) should not be needed (inherited by iClaire)
Compile/srange :: Kernel/srange
// Compile/typing :: Kernel/typing
Compile/status :: Kernel/status
Compile/tmatch? :: Core/tmatch?
Compile/pname :: Kernel/pname

// where to find the CLAIRE libraries
claire/home() : string -> getenv("CLAIRE_HOME")     // TO CHANGE -> simply read claire_home

// CLAIRE4 uses EID where CLAIRE used C++ OID (integer representation)
claire/EID <: object()                               // used as a marker for form EID

// ******************************************************************
// *    Part 1: General Global Variables and Properties             *
// ******************************************************************

// we use an optimizer object with all the necessary resources
// they are all private.
meta_OPT <: thing(
	Compile/instructions:list,        // list of compiled instructions
        Compile/objects:list,          // new named objects that are defined - v3.3.3: the order is important
        Compile/properties:set<property>, // new properties (implicit)
        Compile/functions:list,           // new functions
        Compile/need_to_close:set,        // properties that need to be closed
        Compile/need_modules:set,         // other modules that are needed
        Compile/legal_modules:set,        // modules that are allowed
        Compile/ignore:set,               // do not print
        Compile/to_remove:set,            // do not compile
        Compile/outfile:port,             // port for the .Cpp output
        Compile/max_vars:integer = 0,          // max number of local variables
        Compile/loop_index:integer = 0,        // max index of var in loop
        Compile/level:integer = 0,            // indentation level
        Compile/in_method:any = unknown,      // current method (used to print context info)
        Compile/profile?:boolean = false,     // do we use the profiler
        Compile/cfile:any = false,            // used for file compilation mode
        // use_update:boolean = false,    // used to record update for status
        // use_nth=:boolean = false,      // update on bags
        // online?:boolean = false,       // online optimization
        recompute:boolean = true,      // force to recompute the status
        unsure:list = nil,             // methods that need to be compiled at safety 1
        knowns:set<relation>,          // properties that are safe (no unknown)
        Compile/simple_operations:set<property>,   //  v3.3
        Compile/non_identifiable_set:set<class>,   // v 3.3
        Compile/use_string_update:boolean = false)     // update on strings   v3.3.46
       

// The meta_compiler contains the definition of the compiler flags and slots
// that are important for the user. Other stuff is hidden in OPT
meta_compiler <: thing(external:string,             // name of the output
                       source:string,               // where to put the produced code
                       claire/debug?:list<module>,  // generate instrumented code for modules
                       version:any,
                       claire/active?:boolean = false,     // active = (loading | compiling)
                       claire/safety:integer = 1,           // cf .. later
                       claire/env:string,                   // OS info
                       claire/libraries:list<string>,      // libs to be included
                       claire/inline?:boolean = false,     // do we want to use inlining (should be TRUE)
                       claire/loading?:boolean = false,    // mode (read the file vs compile)
                       claire/overflow?:boolean = false,   // safe arithmetic
                       claire/optimize?:boolean = false,   // v3.2.56: record -O option
                       claire/n_loc:integer,                // number of lines of code
                       claire/n_warnings:integer,           // number of warnings
                       claire/n_notes:integer,              // number of notes
                       claire/n_dynamic:integer,            // number of dynamic calls
                       claire/n_methods:integer,           // number of methods compiled 
                       claire/n_metheids:integer)           // number of methods compiled with EID (error handling)

// code producer are defined in Generate
// but the stub is define in Optimize to have access to current_file
Compile/producer <: thing(
    Compile/current_file:string = "")           // name of the file being compiled

// we use a global variable to hide the indirection through the producer
// this is kept in CLAIRE 4.0 so that the C++ compiler could be re-introduced
claire/PRODUCER:producer :: unknown

// new in CLAIRE4: create an automated comment
Compile/FileOrigin[m:method] : string := ""

// the three variables that are used in the main files
claire/claire_modules:list :: list{get_value(x) | x in list("Kernel","Core","Language","Reader")}

// safety:
//       0  -> super-safe (keep assertion)
//       1  -> safe (regular)
//       2  -> we trus typing
//       3  -> no overflow checking (integer & arrays)
claire/compiler :: meta_compiler(
   external = "go", // Id(compiler.external),
   env = "MacOS",   // Id(compiler.env),
   version = Id(version()),
   source = "",
   libraries = list<string>("Kernel"))

// re-definable items for bootstrap modifications
claire/c_type :: property(open = 3, range = type)
claire/c_code :: property(open = 3)
Compile/get_index :: property(range = integer, open = 3)
Compile/get_indexed :: property(range = list, open = 3)      
Compile/make_c_function :: property(open = 3)
// Compile/make_float_function :: property(Core/open = 3)
// Compile/c_expression :: property(Core/open = 3)
Compile/bool_exp :: property(open = 3)
Compile/c_statement :: property()
claire/c_interface :: property(open = 3)
Compile/c_sort :: property(open = 3)

// other useful properties shared between Optimize & Generate
Compile/g_throw :: property()
Compile/can_throw? :: property()
Compile/can_throw! :: property()
Compile/designated? :: property()
Compile/sort= :: property()
Compile/psort :: property()
Compile/osort :: property()                        // Optimizer version of sorts
Compile/compile_lambda :: property()
Compile/need_protect :: property()
member_code :: property()
Compile/c_strict_code :: property()                // code with strict (stupid) type
Compile/c_strict_check :: property(open = 3)  // new: allow future overload !!
Compile/object! :: property()                      // compiler instantiation
Compile/anyObject! :: property(range = object)     // fast instantiation if all any slots are known
Compile/Cerror :: property()
Compile/self_code :: property()                    // how to compile a type expression
Compile/get_module :: property()
Compile/function_name :: property(open = 3)
Compile/m_unsafe :: (unsafe @ any)
Compile/m_member :: (% @ list(any,any))
Compile/warn :: property()
Compile/nativeVar? :: property()

// these are the classes defined especially for this module
// Compile/to_CL <: Optimized_instruction(arg:any,set_arg:class)
// Compile/to_C <: Optimized_instruction(arg:any,set_arg:class)
Compile/C_cast <: Optimized_instruction(arg:any,set_arg:class)   // was to_C()

// Patterns are calls p(X) that are seen as a type expression
Pattern <: type_expression(
         selector:property,       //
         arg:list)                // the tuple is made into a list

// OPT contains all the parameters for the optimizer
claire/OPT :: meta_OPT(
    outfile = stdin, 
    ignore = set(mClaire/index!, mClaire/set_index, object!, mClaire/base!, mClaire/set_base,
                 mClaire/push!, anyObject!, mClaire/get_stack, mClaire/put_stack),
    to_remove = {},                                         // set(interface),
    knowns = set<relation>(arg1,arg2),                     // v3.1.12
    unsure = list(+ @ integer, * @ integer, - @ integer),
    simple_operations = set<property>(+, -, /, *),                      // v3.3
    non_identifiable_set =
                    Id(set<class>{c in ((class but integer) but float) | 
                                   exists(c2 in c.descendants | c2.ident? = false)}))


// pragma for the compiler  => MOVED TO LANGUAGE in CLAIRE 4
// this pragma tells to compile with full safety (include arithmetic checks)
claire/safe(x:any) : type[x] -> x

// ******************************************************************
// *    Part 2: The defaults for c_type, c_code and c_sort          *
// ******************************************************************


// basic type inference
[c_type(self:any) : type
 -> case self
      (Variable let r := get(range, self) in
                  (if (r = unknown | r = EID) any         // v:EID is a pragma for compiler in CLAIRE4
                   else if (case r (Union r.Kernel/t1 = {})) r.Kernel/t2.Kernel/t2
                   else ptype(r)),
       global_variable let r := self.range in
                         (if r r else set(self.value)),
       unbound_symbol Cerror("[215] the symbol ~A is unbound_symbol", self.name),
       error {},
       Update c_type(self.value),
       Construct (if not(self % (List U Set)) any
                 else let %res:type := {} in
                        (for %x in self.args
                           (if %res %res :meet class!(c_type(%x))
                            else %res := class!(c_type(%x))),
                         nth((case self (Set set, any list)),
                             %res))),
       Instruction error("c_type of ~S is not defined",owner(self)), //<yc:v0.01>
       any set(self)) ]

// compile into a sort and checks strict type matching (naive/stupid)
[c_strict_code(x:any,s:class) : any
 -> c_strict_check(c_code(x, s), s) ]

// CLAIRE 4: introduce C_cast so that psort(x) is what is expected (s)
[c_strict_check(x:any,s:class) : any
 -> if (s inherit? object & not(static_type(x) inherit? s))
       (// [5] c_strict_check is unhappy with ~S: expecting ~S and found ~S [claire:~S] // x,s,static_type(x),c_type(x),
        // if (c_type(x) = any)     // v3.2.06 - avoid C++ compiler error !
        // Cerror("Need explict cast: ~S is not a ~S",x,s),
        C_cast(arg = x, set_arg = s))
     else x ]

// using conversions. s is a sort or void (we do not need the value).
// note: we need s to be the precise sort for C++
// the is the default version that uses c_code(x)/ c_sort(x)
// in CLAIRE 4, we do not generate conversion at optim time
[c_code(x:any,s:class) : any
 -> let y := (case x (Call c_code_call(x,s), 
                      any c_code(x))),   // v2.4.9 safe sort for inline !
        z := c_sort(y) in
      (if (s = void | z = s) //  | OPT.online?)
           (if (s = void & (case x (Call x.selector = =)))
               (warn(),
                trace(2,"-- Equality meant as an assignment: ~S [264]\n",x)),    // v3.3
            y)       // v3.0.44 BIG CHANGE
      // else if (s = any)
      //   (if (z = integer & y % Call_slot &                       // need a proper slot
       //       not(compiler.overflow? & compiler.class2file?)) y   // UGLY: v3.0.42
          //else if (y % to_C) arg(y as to_C)
          // else to_CL(arg = y, set_arg = psort(c_type(y))))
       // else if (z = any) to_C(arg = y, set_arg = s)
       else y) ]

// basic code generation
// c_code without a sort parameter means that we do not care about the resulting sort,
// which will be checked later on using c_sort
[c_code(self:any) : any
 -> case self
      (unbound_symbol Cerror("[215] the symbol ~A is unbound_symbol", self.name),
       Variable self,
       global_variable (c_register(self), self),
       Optimized_instruction self,
       Instruction Cerror("[internal] c_code(~S) should have a parameter", self),
       set (if self
              let x :=  Set(args = list!(self)) in
                 (if (of(self) != void) x.of := of(self),
                  c_code(x))
            else self),
       list (if self
               let x := List(args = self) in
                    (if (of(self) != void) x.of := of(self),
                  c_code(x))
             else self),
       tuple c_code(Tuple(args = list!(self))),
       any (if (self % thing) c_register(self),
            self)) ]

// suggestion for claire4 : get rid of c_sort
[get_sort(self:any) : class 
   -> static_type(self)]

// gives the sort of a compiled expression (does not apply to instructions that
// have a direct c_code(x,s)
// v2.4.9: special type => special sorts !!!
[c_sort(self:any) : class
 -> case self
      (global_variable (if nativeVar?(self) (if (self.range = {}) osort(owner(self.value))
                                             else osort(self.range))
                        else any),    // v3.3 ! was any,
       Instruction case self
                   (Variable sort(self),
                    Assign sort(self.var),
                    Call osort(selector_psort(self)),
                    Call_method (if (self.arg.selector = externC & length(self.args) = 2)
                                    psort(self.args[2] as class)
                                 else c_srange(self.arg)),
                    Call_slot srange(self.selector),
                    Call_table any,
                    Call_array (if (self.test = float) float else any),
                    Definition object,
                    C_cast self.set_arg,     // was to_C (includes C_cast)
                    // to_CL any,
                    Update c_sort(self.value),
                    If psort(c_sort(self.arg) meet c_sort(self.other)),
                    Handle  psort(c_sort(self.arg) meet c_sort(self.other)),
                    Let c_sort(self.arg),
                    Return  any,
                    List object,
                    Set object,
                    Tuple object,
                    Construct any,
                    Gassign  any,
                    Super any,
                    For any,
                    Exists (if (self.other = unknown) any else object),
                    Iteration object,
                    And boolean,
                    Or boolean,
                    While any,
                    Do c_sort(last(self.args)),
                    any error("[internal] c_sort is not implemented for ~S",
                              owner(self))),
       float float,
       any psort(c_type(self))) ]

// for the special compiler properties, we need to tell the sort of the optimized
// form
[selector_psort(self:Call) : class
  -> let p := self.selector in
       (if (p = mClaire/base! | p = mClaire/index!) integer
        else if (p = anyObject!) (self.args[1] as class)
        else if (p = object!) (self.args[2] as class)
        else any) ]


// ******************************************************************
// *    Part 3: g_throw and status(m:method)                        *
// ******************************************************************

// NEW in claire4 : optimization when compiler.safety is high may prevent throwing exceptions
Compile/NoErrorOptimize :: list<any>(nth @ list, nth= @ list, nth @ tuple, nth @ string,
     less?, glb,
     // traces (tformat) are assumed to be correct in cross_compiling
     Core/tformat @ string,
     // <= @ type_expression is safe when closed
     <= @ type_expression, >=,
     // optimize -> do not check for arithmetic overflow
     * @ list(integer,integer))

// these two variabler are used for cross-compiling, when the status changes from the existing(compiled) version to the
// new one being compiled
Compile/ForceThrow :: list<method>()
Compile/ForceNotThrow :: list<method>()  


// NEW in claire 4, because error handling is mananaged by the compiler
// tells if an expression can throw an exception, based on can_throw?(p or m)

// debug loop
claire/DSHOW:boolean := false

[g_throw(self:any) : boolean
  -> let v := g_throw1(self) in
        (if DSHOW trace(0,"-> g_throw(~S)=~S\n",self,v),
     v)]
       
[g_throw1(self:any) : boolean
 -> case self
      (bag exists(x in self | g_throw(x)),
       Construct exists(x in self.args | g_throw(x)),
       Assign g_throw(get(arg, self)),
       Gassign g_throw(get(arg, self)),
       And g_throw(self.args),
       Or g_throw(self.args),
       Call  (self.selector != unsafe & (g_throw(self.args) | can_throw?(self.selector))),
       Call_method (self.arg != m_unsafe & notOpt(self) & self.arg.selector != externC &
                     (g_throw(self.args) | can_throw?(self.arg))),
       Call_slot (g_throw(get(arg, self)) | (known?(test,self) & self.test)),
       Call_table (g_throw(get(arg, self)) | (known?(test,self) & self.test)),
       Call_array (g_throw(get(selector,self)) | g_throw(get(arg, self))),
       Super (g_throw(self.args) | can_throw?(self.selector)),                                  // selector can throw
       Update ((g_throw(get(value, self)) | g_throw(get(var, self))) | Compile/update_write?(self)),    // if_write may throw
       Cast g_throw(self.arg),
       C_cast g_throw(self.arg),
       // Generate/to_C g_throw(self.arg),
       // Generate/to_CL g_throw(self.arg),
       // add all control instructions
       Let (g_throw(self.value) | g_throw(self.arg)),
       Do g_throw(self.args),
       While (g_throw(self.test) | g_throw(self.arg)),
       Construct g_throw(self.args),
       If (g_throw(self.test) | g_throw(self.arg) | g_throw(self.other)),
       For (g_throw(self.set_arg) | g_throw(self.arg)),
       Iteration (g_throw(self.set_arg) | g_throw(self.arg)),
       Handle (self.iClaire/test != any | g_throw(self.other)),       // test = any => safe: catch all
       // anything else is simple
       any false) ]

// return true in regular case, false if the optimization means that no error will occur.
//  this is ugly, so use sparingly for truly critical code optimization:  
//    - (x % y) can raise an error in the generic case (using F_belong) but not in the  optimized case
//    - class!(...) can raise an error in interpreted mode nut not at compile time
//    - division by non-zero integer constant is OK
//    - etc (extensible)   ... hopefully go will support exceptions one day so I can get rid of this junk :)
[notOpt(self:Call_method) : boolean 
    -> if (self.arg = m_member)  
             (let t2 := static_type(self.args[2]) in
                   not( t2 <= type | t2 <= list | t2 <= integer | t2 <= array))    // these are the 4 exceptions (optimized)
       else if (self.arg.selector = class!) not(self.args[1] % symbol)    // constant symbol => conflicts have been checked
       else if (self.arg.selector = / | self.arg.selector = mod) (not(self.args[2] % integer) | self.args[2] = 0)
       else if (self.args[1] % table & self.arg.selector = nth)  not(static_type(self.args[2]) <= domain(self.args[1])) & (compiler.safety < 2)
       else true ]        //regular case !


// can_throw is based on restrictions analysis ... unless it is open => could always return an error
[can_throw?(p:property) : boolean
  -> (p.open = 3 |  (not( p % Compile/NoErrorOptimize) & 
          exists(m in p.restrictions | (case m (method can_throw?(m), any false))))) ]

// access to status ... -1 means that it was never computed 
// Force*Throw is used to adjust for cross-compiling with a status change
// the reference to safety is here to ensure cross-compilation mode
[can_throw?(m:method) : boolean
  -> if ((safety(compiler) >= 2 & (m % NoErrorOptimize | (m.selector % NoErrorOptimize))) | 
         m % ForceNotThrow) false
     else if (m % ForceThrow) true
     else if (m.status != -1 | unknown?(formula,m)) (m.status != 0)
     else can_throw!(m) ]

// debug handle
claire/DTHROW:any :: unknown

// here we recursively call g_throw on the body => forced re-compute of status(m) (status!(m) in CLAIRE3)
[can_throw!(m:method) : boolean
  -> m.status := 0,                                             // optimistic .. to avoid recursive problems
     DTHROW := c_code(m.formula.body, class!(m.range)),
     if g_throw(c_code(m.formula.body, class!(m.range)))       // QUESTION: compile with c_code ?
        (m.status := 1, true)
     else false ]

// read can_throw from the status, not influenced by exceptions (for code generation)
// however, for a new method, compute the status
[Compile/can_throw_status(m:method) : boolean
  -> if (m.status = -1) can_throw!(m),
     (m.status != 0) ]   


// useful #2: provoke a recomputation of status
[claire/s_throw(m:method) : void
  -> let  la := m.formula,
          news := g_throw(la.body) in
       (//[0] status(~S) := ~S // m, news,
        Core/status(m) := (if news 1 else 0)) ]

// ******************************************************************
// *    Part 4: Names & identifiers management                      *
// ******************************************************************

// check that the module is allowed and otherwise complain because of x;
// this should raise an error, it simply returns false if there is a problem
[legal?(self:module,x:any) : boolean
 -> if (x = object! | x = anyObject!) true
    else if (OPT.legal_modules)
       (if (not(self % OPT.legal_modules)  &
            (case x (method (x.selector != add_method & (x.inline? = false | not(compiler.inline?))))))
        (trace(0, "legal_modules = ~S\n", OPT.legal_modules),
         trace(0, "---- ERROR: ~S implies using ~S !\n\n", x, self),
         false)
        else true)
     else (need_modules(OPT) :add self, true) ]

legal?(self:environment,x:any) : any -> true

// A named object is used, thus it must be declared if it belongs to the
// current module - returns true if OK
[c_register(self:(thing U class)) : any
 -> let x := get_module(self) in
       (if (x != system) legal?(x, self) else true) ]

// looks if a property may be implicit and then add it in the right list
[c_register(self:property) : any
 -> let m := module!(), m2 := get_module(self) in
     (// while (m.loaded = 0) m := m.part_of,     // <yc> 7/98: we need the compilation unit
      if ((m2 = claire | m2 = m) & //   v3.2.4  (self % property) &
          not(self % OPT.objects)) //   v3.2.4  & not(self % OPT.properties))
        OPT.properties :add self,
      c_register@thing(self)) ]

// declare the property as used and check if a property may allocate
[selector_register(self:property) : any
 ->  c_register(self),
     self ]


// this method looks if the open slot is less than 1 or can be set to 1
// v3.3.48 note - weaken the open semantic to get a better c_status
[stable?(self:relation) : boolean
 -> let m := get_module(self) in
       (if (self.open = 2) // v3.3.48  - was :  & (m % OPT.legal_modules | m = system))
           write(open, self, 1)),
     self.open <= 1 | self.open = 4 ]                       // v3.2.04

// returns the module (i.e. the compilation unit, not the namespace) in which self is
// defined
[get_module(self:(thing U class)) : any
  -> defined(self.name) ]
//      (while (m.loaded = 0) m := m.part_of, m) ]


// allows to optimize the access
[known!(l:listargs) : any
 -> (to_remove(OPT) :add known!,
     for r in l (case r (property OPT.knowns :add r))) ]




