/***** CLAIRE Compilation of module Language.cl 
         [version 4.1.4 / safety 5] Friday 01-03-2025 14:52:12 *****/

package Language
import (_ "fmt"
	"unsafe"
	. "Kernel"
	"Core"
)

//-------- dumb function to prevent import errors --------
func import_g0332() { 
_ = Core.It
} 


// class file for Basic_instruction in module Language 
type BasicInstruction struct { 
   ClaireInstruction
   } 

// automatic cast function
func To_BasicInstruction(x *ClaireAny) *BasicInstruction {return (*BasicInstruction)(unsafe.Pointer(x))}

// automatic constructor function
func Make_BasicInstruction() *BasicInstruction { 
  var o *BasicInstruction = new(BasicInstruction)
  o.Isa = C_Basic_instruction
  return o 
  } 

// class file for Vardef in module Language 
type Vardef struct { 
   ClaireVariable
   } 

// automatic cast function
func To_Vardef(x *ClaireAny) *Vardef {return (*Vardef)(unsafe.Pointer(x))}

// automatic constructor function
func Make_Vardef(pname *ClaireSymbol,Range *ClaireType,index int) *Vardef { 
  var o *Vardef = new(Vardef)
  o.Isa = C_Vardef
  o.Pname = pname
  o.Range = Range
  o.Index = index
  return o 
  } 

// class file for Complex_instruction in module Language 
type ComplexInstruction struct { 
   ClaireInstruction
   } 

// automatic cast function
func To_ComplexInstruction(x *ClaireAny) *ComplexInstruction {return (*ComplexInstruction)(unsafe.Pointer(x))}

// automatic constructor function
func Make_ComplexInstruction() *ComplexInstruction { 
  var o *ComplexInstruction = new(ComplexInstruction)
  o.Isa = C_Complex_instruction
  return o 
  } 

// class file for Instruction_with_var in module Language 
type InstructionWithVar struct { 
   ComplexInstruction
   ClaireVar *ClaireVariable
  } 

// automatic cast function
func To_InstructionWithVar(x *ClaireAny) *InstructionWithVar {return (*InstructionWithVar)(unsafe.Pointer(x))}

// automatic constructor function
func Make_InstructionWithVar(ClaireVar *ClaireVariable) *InstructionWithVar { 
  var o *InstructionWithVar = new(InstructionWithVar)
  o.Isa = C_Instruction_with_var
  o.ClaireVar = ClaireVar
  return o 
  } 

// class file for Control_structure in module Language 
type ControlStructure struct { 
   ComplexInstruction
   } 

// automatic cast function
func To_ControlStructure(x *ClaireAny) *ControlStructure {return (*ControlStructure)(unsafe.Pointer(x))}

// automatic constructor function
func Make_ControlStructure() *ControlStructure { 
  var o *ControlStructure = new(ControlStructure)
  o.Isa = C_Control_structure
  return o 
  } 

// class file for Call in module Language 
type Call struct { 
   ControlStructure
   Selector *ClaireProperty
  Args *ClaireList
  } 

// automatic cast function
func To_Call(x *ClaireAny) *Call {return (*Call)(unsafe.Pointer(x))}

// automatic constructor function
func Make_Call(selector *ClaireProperty,args *ClaireList) *Call { 
  var o *Call = new(Call)
  o.Isa = C_Call
  o.Selector = selector
  o.Args = args
  return o 
  } 

// class file for Call* in module Language 
type Call_star struct { 
   Call
   } 

// automatic cast function
func To_Call_star(x *ClaireAny) *Call_star {return (*Call_star)(unsafe.Pointer(x))}

// automatic constructor function
func Make_Call_star(selector *ClaireProperty,args *ClaireList) *Call_star { 
  var o *Call_star = new(Call_star)
  o.Isa = C_Call_star
  o.Selector = selector
  o.Args = args
  return o 
  } 

// class file for Call+ in module Language 
type Call_plus struct { 
   Call
   } 

// automatic cast function
func To_Call_plus(x *ClaireAny) *Call_plus {return (*Call_plus)(unsafe.Pointer(x))}

// automatic constructor function
func Make_Call_plus(selector *ClaireProperty,args *ClaireList) *Call_plus { 
  var o *Call_plus = new(Call_plus)
  o.Isa = C_Call_plus
  o.Selector = selector
  o.Args = args
  return o 
  } 

// class file for Assign in module Language 
type Assign struct { 
   BasicInstruction
   ClaireVar *ClaireAny
  Arg *ClaireAny
  } 

// automatic cast function
func To_Assign(x *ClaireAny) *Assign {return (*Assign)(unsafe.Pointer(x))}

// automatic constructor function
func Make_Assign(ClaireVar *ClaireAny,arg *ClaireAny) *Assign { 
  var o *Assign = new(Assign)
  o.Isa = C_Assign
  o.ClaireVar = ClaireVar
  o.Arg = arg
  return o 
  } 

// class file for Gassign in module Language 
type Gassign struct { 
   BasicInstruction
   ClaireVar *Core.GlobalVariable
  Arg *ClaireAny
  } 

// automatic cast function
func To_Gassign(x *ClaireAny) *Gassign {return (*Gassign)(unsafe.Pointer(x))}

// automatic constructor function
func Make_Gassign(ClaireVar *Core.GlobalVariable,arg *ClaireAny) *Gassign { 
  var o *Gassign = new(Gassign)
  o.Isa = C_Gassign
  o.ClaireVar = ClaireVar
  o.Arg = arg
  return o 
  } 

// class file for And in module Language 
type And struct { 
   ControlStructure
   Args *ClaireList
  } 

// automatic cast function
func To_And(x *ClaireAny) *And {return (*And)(unsafe.Pointer(x))}

// automatic constructor function
func Make_And(args *ClaireList) *And { 
  var o *And = new(And)
  o.Isa = C_And
  o.Args = args
  return o 
  } 

// class file for Or in module Language 
type Or struct { 
   ControlStructure
   Args *ClaireList
  } 

// automatic cast function
func To_Or(x *ClaireAny) *Or {return (*Or)(unsafe.Pointer(x))}

// automatic constructor function
func Make_Or(args *ClaireList) *Or { 
  var o *Or = new(Or)
  o.Isa = C_Or
  o.Args = args
  return o 
  } 

// class file for Quote in module Language 
type Quote struct { 
   BasicInstruction
   Arg *ClaireAny
  } 

// automatic cast function
func To_Quote(x *ClaireAny) *Quote {return (*Quote)(unsafe.Pointer(x))}

// automatic constructor function
func Make_Quote(arg *ClaireAny) *Quote { 
  var o *Quote = new(Quote)
  o.Isa = C_Quote
  o.Arg = arg
  return o 
  } 

// class file for Optimized_instruction in module Language 
type OptimizedInstruction struct { 
   ComplexInstruction
   } 

// automatic cast function
func To_OptimizedInstruction(x *ClaireAny) *OptimizedInstruction {return (*OptimizedInstruction)(unsafe.Pointer(x))}

// automatic constructor function
func Make_OptimizedInstruction() *OptimizedInstruction { 
  var o *OptimizedInstruction = new(OptimizedInstruction)
  o.Isa = C_Optimized_instruction
  return o 
  } 

// class file for Call_method in module Language 
type CallMethod struct { 
   OptimizedInstruction
   Arg *ClaireMethod
  Args *ClaireList
  } 

// automatic cast function
func To_CallMethod(x *ClaireAny) *CallMethod {return (*CallMethod)(unsafe.Pointer(x))}

// automatic constructor function
func Make_CallMethod(arg *ClaireMethod,args *ClaireList) *CallMethod { 
  var o *CallMethod = new(CallMethod)
  o.Isa = C_Call_method
  o.Arg = arg
  o.Args = args
  return o 
  } 

// class file for Call_method1 in module Language 
type CallMethod1 struct { 
   CallMethod
   } 

// automatic cast function
func To_CallMethod1(x *ClaireAny) *CallMethod1 {return (*CallMethod1)(unsafe.Pointer(x))}

// automatic constructor function
func Make_CallMethod1(arg *ClaireMethod,args *ClaireList) *CallMethod1 { 
  var o *CallMethod1 = new(CallMethod1)
  o.Isa = C_Call_method1
  o.Arg = arg
  o.Args = args
  return o 
  } 

// class file for Call_method2 in module Language 
type CallMethod2 struct { 
   CallMethod
   } 

// automatic cast function
func To_CallMethod2(x *ClaireAny) *CallMethod2 {return (*CallMethod2)(unsafe.Pointer(x))}

// automatic constructor function
func Make_CallMethod2(arg *ClaireMethod,args *ClaireList) *CallMethod2 { 
  var o *CallMethod2 = new(CallMethod2)
  o.Isa = C_Call_method2
  o.Arg = arg
  o.Args = args
  return o 
  } 

// class file for Language/Call_method3 in module Language 
type Language_CallMethod3 struct { 
   CallMethod
   } 

// automatic cast function
func To_Language_CallMethod3(x *ClaireAny) *Language_CallMethod3 {return (*Language_CallMethod3)(unsafe.Pointer(x))}

// automatic constructor function
func Make_Language_CallMethod3(arg *ClaireMethod,args *ClaireList) *Language_CallMethod3 { 
  var o *Language_CallMethod3 = new(Language_CallMethod3)
  o.Isa = C_Language_Call_method3
  o.Arg = arg
  o.Args = args
  return o 
  } 

// class file for Call_slot in module Language 
type CallSlot struct { 
   OptimizedInstruction
   Selector *ClaireSlot
  Arg *ClaireAny
  Test *ClaireBoolean
  } 

// automatic cast function
func To_CallSlot(x *ClaireAny) *CallSlot {return (*CallSlot)(unsafe.Pointer(x))}

// automatic constructor function
func Make_CallSlot(selector *ClaireSlot,arg *ClaireAny,test *ClaireBoolean) *CallSlot { 
  var o *CallSlot = new(CallSlot)
  o.Isa = C_Call_slot
  o.Selector = selector
  o.Arg = arg
  o.Test = test
  return o 
  } 

// class file for Call_array in module Language 
type CallArray struct { 
   OptimizedInstruction
   Selector *ClaireAny
  Arg *ClaireAny
  Test *ClaireAny
  } 

// automatic cast function
func To_CallArray(x *ClaireAny) *CallArray {return (*CallArray)(unsafe.Pointer(x))}

// automatic constructor function
func Make_CallArray(selector *ClaireAny,arg *ClaireAny,test *ClaireAny) *CallArray { 
  var o *CallArray = new(CallArray)
  o.Isa = C_Call_array
  o.Selector = selector
  o.Arg = arg
  o.Test = test
  return o 
  } 

// class file for Call_table in module Language 
type CallTable struct { 
   OptimizedInstruction
   Selector *ClaireTable
  Arg *ClaireAny
  Test *ClaireBoolean
  } 

// automatic cast function
func To_CallTable(x *ClaireAny) *CallTable {return (*CallTable)(unsafe.Pointer(x))}

// automatic constructor function
func Make_CallTable(selector *ClaireTable,arg *ClaireAny,test *ClaireBoolean) *CallTable { 
  var o *CallTable = new(CallTable)
  o.Isa = C_Call_table
  o.Selector = selector
  o.Arg = arg
  o.Test = test
  return o 
  } 

// class file for Update in module Language 
type Update struct { 
   OptimizedInstruction
   Selector *ClaireAny
  Arg *ClaireAny
  Value *ClaireAny
  ClaireVar *ClaireAny
  } 

// automatic cast function
func To_Update(x *ClaireAny) *Update {return (*Update)(unsafe.Pointer(x))}

// automatic constructor function
func Make_Update(selector *ClaireAny,arg *ClaireAny,value *ClaireAny,ClaireVar *ClaireAny) *Update { 
  var o *Update = new(Update)
  o.Isa = C_Update
  o.Selector = selector
  o.Arg = arg
  o.Value = value
  o.ClaireVar = ClaireVar
  return o 
  } 

// class file for Super in module Language 
type Super struct { 
   ControlStructure
   Selector *ClaireProperty
  CastTo *ClaireType
  Args *ClaireList
  } 

// automatic cast function
func To_Super(x *ClaireAny) *Super {return (*Super)(unsafe.Pointer(x))}

// automatic constructor function
func Make_Super(selector *ClaireProperty,cast_to *ClaireType,args *ClaireList) *Super { 
  var o *Super = new(Super)
  o.Isa = C_Super
  o.Selector = selector
  o.CastTo = cast_to
  o.Args = args
  return o 
  } 

// class file for Cast in module Language 
type Cast struct { 
   BasicInstruction
   Arg *ClaireAny
  SetArg *ClaireType
  } 

// automatic cast function
func To_Cast(x *ClaireAny) *Cast {return (*Cast)(unsafe.Pointer(x))}

// automatic constructor function
func Make_Cast(arg *ClaireAny,set_arg *ClaireType) *Cast { 
  var o *Cast = new(Cast)
  o.Isa = C_Cast
  o.Arg = arg
  o.SetArg = set_arg
  return o 
  } 

// class file for Return in module Language 
type Return struct { 
   BasicInstruction
   Arg *ClaireAny
  } 

// automatic cast function
func To_Return(x *ClaireAny) *Return {return (*Return)(unsafe.Pointer(x))}

// automatic constructor function
func Make_Return(arg *ClaireAny) *Return { 
  var o *Return = new(Return)
  o.Isa = C_Return
  o.Arg = arg
  return o 
  } 

// class file for If in module Language 
type If struct { 
   ControlStructure
   Test *ClaireAny
  Arg *ClaireAny
  Other *ClaireAny
  } 

// automatic cast function
func To_If(x *ClaireAny) *If {return (*If)(unsafe.Pointer(x))}

// automatic constructor function
func Make_If(test *ClaireAny,arg *ClaireAny,other *ClaireAny) *If { 
  var o *If = new(If)
  o.Isa = C_If
  o.Test = test
  o.Arg = arg
  o.Other = other
  return o 
  } 

// class file for If? in module Language 
type If_ask struct { 
   If
   } 

// automatic cast function
func To_If_ask(x *ClaireAny) *If_ask {return (*If_ask)(unsafe.Pointer(x))}

// automatic constructor function
func Make_If_ask(test *ClaireAny,arg *ClaireAny,other *ClaireAny) *If_ask { 
  var o *If_ask = new(If_ask)
  o.Isa = C_If_ask
  o.Test = test
  o.Arg = arg
  o.Other = other
  return o 
  } 

// class file for Do in module Language 
type Do struct { 
   ControlStructure
   Args *ClaireList
  } 

// automatic cast function
func To_Do(x *ClaireAny) *Do {return (*Do)(unsafe.Pointer(x))}

// automatic constructor function
func Make_Do(args *ClaireList) *Do { 
  var o *Do = new(Do)
  o.Isa = C_Do
  o.Args = args
  return o 
  } 

// class file for Let in module Language 
type Let struct { 
   InstructionWithVar
   Value *ClaireAny
  Arg *ClaireAny
  } 

// automatic cast function
func To_Let(x *ClaireAny) *Let {return (*Let)(unsafe.Pointer(x))}

// automatic constructor function
func Make_Let(ClaireVar *ClaireVariable,value *ClaireAny,arg *ClaireAny) *Let { 
  var o *Let = new(Let)
  o.Isa = C_Let
  o.ClaireVar = ClaireVar
  o.Value = value
  o.Arg = arg
  return o 
  } 

// class file for When in module Language 
type When struct { 
   Let
   Other *ClaireAny
  } 

// automatic cast function
func To_When(x *ClaireAny) *When {return (*When)(unsafe.Pointer(x))}

// automatic constructor function
func Make_When(ClaireVar *ClaireVariable,value *ClaireAny,arg *ClaireAny,other *ClaireAny) *When { 
  var o *When = new(When)
  o.Isa = C_When
  o.ClaireVar = ClaireVar
  o.Value = value
  o.Arg = arg
  o.Other = other
  return o 
  } 

// class file for Let+ in module Language 
type Let_plus struct { 
   Let
   } 

// automatic cast function
func To_Let_plus(x *ClaireAny) *Let_plus {return (*Let_plus)(unsafe.Pointer(x))}

// automatic constructor function
func Make_Let_plus(ClaireVar *ClaireVariable,value *ClaireAny,arg *ClaireAny) *Let_plus { 
  var o *Let_plus = new(Let_plus)
  o.Isa = C_Let_plus
  o.ClaireVar = ClaireVar
  o.Value = value
  o.Arg = arg
  return o 
  } 

// class file for Let* in module Language 
type Let_star struct { 
   Let
   } 

// automatic cast function
func To_Let_star(x *ClaireAny) *Let_star {return (*Let_star)(unsafe.Pointer(x))}

// automatic constructor function
func Make_Let_star(ClaireVar *ClaireVariable,value *ClaireAny,arg *ClaireAny) *Let_star { 
  var o *Let_star = new(Let_star)
  o.Isa = C_Let_star
  o.ClaireVar = ClaireVar
  o.Value = value
  o.Arg = arg
  return o 
  } 

// class file for Iteration in module Language 
type Iteration struct { 
   InstructionWithVar
   SetArg *ClaireAny
  Arg *ClaireAny
  } 

// automatic cast function
func To_Iteration(x *ClaireAny) *Iteration {return (*Iteration)(unsafe.Pointer(x))}

// automatic constructor function
func Make_Iteration(ClaireVar *ClaireVariable,set_arg *ClaireAny,arg *ClaireAny) *Iteration { 
  var o *Iteration = new(Iteration)
  o.Isa = C_Iteration
  o.ClaireVar = ClaireVar
  o.SetArg = set_arg
  o.Arg = arg
  return o 
  } 

// class file for For in module Language 
type For struct { 
   Iteration
   } 

// automatic cast function
func To_For(x *ClaireAny) *For {return (*For)(unsafe.Pointer(x))}

// automatic constructor function
func Make_For(ClaireVar *ClaireVariable,set_arg *ClaireAny,arg *ClaireAny) *For { 
  var o *For = new(For)
  o.Isa = C_For
  o.ClaireVar = ClaireVar
  o.SetArg = set_arg
  o.Arg = arg
  return o 
  } 

// class file for Collect in module Language 
type Collect struct { 
   Iteration
   Of *ClaireType
  } 

// automatic cast function
func To_Collect(x *ClaireAny) *Collect {return (*Collect)(unsafe.Pointer(x))}

// automatic constructor function
func Make_Collect(ClaireVar *ClaireVariable,set_arg *ClaireAny,arg *ClaireAny,of *ClaireType) *Collect { 
  var o *Collect = new(Collect)
  o.Isa = C_Collect
  o.ClaireVar = ClaireVar
  o.SetArg = set_arg
  o.Arg = arg
  o.Of = of
  return o 
  } 

// class file for Image in module Language 
type Image struct { 
   Iteration
   Of *ClaireType
  } 

// automatic cast function
func To_Image(x *ClaireAny) *Image {return (*Image)(unsafe.Pointer(x))}

// automatic constructor function
func Make_Image(ClaireVar *ClaireVariable,set_arg *ClaireAny,arg *ClaireAny,of *ClaireType) *Image { 
  var o *Image = new(Image)
  o.Isa = C_Image
  o.ClaireVar = ClaireVar
  o.SetArg = set_arg
  o.Arg = arg
  o.Of = of
  return o 
  } 

// class file for Select in module Language 
type Select struct { 
   Iteration
   Of *ClaireType
  } 

// automatic cast function
func To_Select(x *ClaireAny) *Select {return (*Select)(unsafe.Pointer(x))}

// automatic constructor function
func Make_Select(ClaireVar *ClaireVariable,set_arg *ClaireAny,arg *ClaireAny,of *ClaireType) *Select { 
  var o *Select = new(Select)
  o.Isa = C_Select
  o.ClaireVar = ClaireVar
  o.SetArg = set_arg
  o.Arg = arg
  o.Of = of
  return o 
  } 

// class file for Lselect in module Language 
type Lselect struct { 
   Iteration
   Of *ClaireType
  } 

// automatic cast function
func To_Lselect(x *ClaireAny) *Lselect {return (*Lselect)(unsafe.Pointer(x))}

// automatic constructor function
func Make_Lselect(ClaireVar *ClaireVariable,set_arg *ClaireAny,arg *ClaireAny,of *ClaireType) *Lselect { 
  var o *Lselect = new(Lselect)
  o.Isa = C_Lselect
  o.ClaireVar = ClaireVar
  o.SetArg = set_arg
  o.Arg = arg
  o.Of = of
  return o 
  } 

// class file for Exists in module Language 
type Exists struct { 
   Iteration
   Other *ClaireAny
  } 

// automatic cast function
func To_Exists(x *ClaireAny) *Exists {return (*Exists)(unsafe.Pointer(x))}

// automatic constructor function
func Make_Exists(ClaireVar *ClaireVariable,set_arg *ClaireAny,arg *ClaireAny,other *ClaireAny) *Exists { 
  var o *Exists = new(Exists)
  o.Isa = C_Exists
  o.ClaireVar = ClaireVar
  o.SetArg = set_arg
  o.Arg = arg
  o.Other = other
  return o 
  } 

// class file for Case in module Language 
type Case struct { 
   ControlStructure
   ClaireVar *ClaireAny
  Args *ClaireList
  } 

// automatic cast function
func To_Case(x *ClaireAny) *Case {return (*Case)(unsafe.Pointer(x))}

// automatic constructor function
func Make_Case(ClaireVar *ClaireAny,args *ClaireList) *Case { 
  var o *Case = new(Case)
  o.Isa = C_Case
  o.ClaireVar = ClaireVar
  o.Args = args
  return o 
  } 

// class file for While in module Language 
type While struct { 
   ControlStructure
   Test *ClaireAny
  Arg *ClaireAny
  Other *ClaireBoolean
  } 

// automatic cast function
func To_While(x *ClaireAny) *While {return (*While)(unsafe.Pointer(x))}

// automatic constructor function
func Make_While(test *ClaireAny,arg *ClaireAny,other *ClaireBoolean) *While { 
  var o *While = new(While)
  o.Isa = C_While
  o.Test = test
  o.Arg = arg
  o.Other = other
  return o 
  } 

// class file for Handle in module Language 
type ClaireHandle struct { 
   ControlStructure
   Test *ClaireAny
  Arg *ClaireAny
  Other *ClaireAny
  } 

// automatic cast function
func To_ClaireHandle(x *ClaireAny) *ClaireHandle {return (*ClaireHandle)(unsafe.Pointer(x))}

// automatic constructor function
func Make_ClaireHandle(test *ClaireAny,arg *ClaireAny,other *ClaireAny) *ClaireHandle { 
  var o *ClaireHandle = new(ClaireHandle)
  o.Isa = C_Handle
  o.Test = test
  o.Arg = arg
  o.Other = other
  return o 
  } 

// class file for Construct in module Language 
type Construct struct { 
   ComplexInstruction
   Args *ClaireList
  } 

// automatic cast function
func To_Construct(x *ClaireAny) *Construct {return (*Construct)(unsafe.Pointer(x))}

// automatic constructor function
func Make_Construct(args *ClaireList) *Construct { 
  var o *Construct = new(Construct)
  o.Isa = C_Construct
  o.Args = args
  return o 
  } 

// class file for List in module Language 
type List struct { 
   Construct
   Of *ClaireType
  } 

// automatic cast function
func To_List(x *ClaireAny) *List {return (*List)(unsafe.Pointer(x))}

// automatic constructor function
func Make_List(args *ClaireList,of *ClaireType) *List { 
  var o *List = new(List)
  o.Isa = C_List
  o.Args = args
  o.Of = of
  return o 
  } 

// class file for Tuple in module Language 
type Tuple struct { 
   Construct
   } 

// automatic cast function
func To_Tuple(x *ClaireAny) *Tuple {return (*Tuple)(unsafe.Pointer(x))}

// automatic constructor function
func Make_Tuple(args *ClaireList) *Tuple { 
  var o *Tuple = new(Tuple)
  o.Isa = C_Tuple
  o.Args = args
  return o 
  } 

// class file for Set in module Language 
type Set struct { 
   Construct
   Of *ClaireType
  } 

// automatic cast function
func To_Set(x *ClaireAny) *Set {return (*Set)(unsafe.Pointer(x))}

// automatic constructor function
func Make_Set(args *ClaireList,of *ClaireType) *Set { 
  var o *Set = new(Set)
  o.Isa = C_Set
  o.Args = args
  o.Of = of
  return o 
  } 

// class file for Array in module Language 
type Array struct { 
   Construct
   Of *ClaireType
  } 

// automatic cast function
func To_Array(x *ClaireAny) *Array {return (*Array)(unsafe.Pointer(x))}

// automatic constructor function
func Make_Array(args *ClaireList,of *ClaireType) *Array { 
  var o *Array = new(Array)
  o.Isa = C_Array
  o.Args = args
  o.Of = of
  return o 
  } 

// class file for Printf in module Language 
type Printf struct { 
   Construct
   } 

// automatic cast function
func To_Printf(x *ClaireAny) *Printf {return (*Printf)(unsafe.Pointer(x))}

// automatic constructor function
func Make_Printf(args *ClaireList) *Printf { 
  var o *Printf = new(Printf)
  o.Isa = C_Printf
  o.Args = args
  return o 
  } 

// class file for Error in module Language 
type Error struct { 
   Construct
   } 

// automatic cast function
func To_Error(x *ClaireAny) *Error {return (*Error)(unsafe.Pointer(x))}

// automatic constructor function
func Make_Error(args *ClaireList) *Error { 
  var o *Error = new(Error)
  o.Isa = C_Error
  o.Args = args
  return o 
  } 

// class file for Branch in module Language 
type Branch struct { 
   Construct
   } 

// automatic cast function
func To_Branch(x *ClaireAny) *Branch {return (*Branch)(unsafe.Pointer(x))}

// automatic constructor function
func Make_Branch(args *ClaireList) *Branch { 
  var o *Branch = new(Branch)
  o.Isa = C_Branch
  o.Args = args
  return o 
  } 

// class file for Map in module Language 
type Map struct { 
   Construct
   Domain *ClaireType
  Of *ClaireType
  } 

// automatic cast function
func To_Map(x *ClaireAny) *Map {return (*Map)(unsafe.Pointer(x))}

// automatic constructor function
func Make_Map(args *ClaireList,domain *ClaireType,of *ClaireType) *Map { 
  var o *Map = new(Map)
  o.Isa = C_Map
  o.Args = args
  o.Domain = domain
  o.Of = of
  return o 
  } 

// class file for Macro in module Language 
type Macro struct { 
   Construct
   } 

// automatic cast function
func To_Macro(x *ClaireAny) *Macro {return (*Macro)(unsafe.Pointer(x))}

// automatic constructor function
func Make_Macro(args *ClaireList) *Macro { 
  var o *Macro = new(Macro)
  o.Isa = C_Macro
  o.Args = args
  return o 
  } 

// class file for Trace in module Language 
type Trace struct { 
   Construct
   } 

// automatic cast function
func To_Trace(x *ClaireAny) *Trace {return (*Trace)(unsafe.Pointer(x))}

// automatic constructor function
func Make_Trace(args *ClaireList) *Trace { 
  var o *Trace = new(Trace)
  o.Isa = C_Trace
  o.Args = args
  return o 
  } 

// class file for Assert in module Language 
type Assert struct { 
   Construct
   Index int
  External *ClaireString
  } 

// automatic cast function
func To_Assert(x *ClaireAny) *Assert {return (*Assert)(unsafe.Pointer(x))}

// automatic constructor function
func Make_Assert(args *ClaireList,index int,external *ClaireString) *Assert { 
  var o *Assert = new(Assert)
  o.Isa = C_Assert
  o.Args = args
  o.Index = index
  o.External = external
  return o 
  } 

// class file for Defclaire in module Language 
type Defclaire struct { 
   ComplexInstruction
   } 

// automatic cast function
func To_Defclaire(x *ClaireAny) *Defclaire {return (*Defclaire)(unsafe.Pointer(x))}

// automatic constructor function
func Make_Defclaire() *Defclaire { 
  var o *Defclaire = new(Defclaire)
  o.Isa = C_Defclaire
  return o 
  } 

// class file for Definition in module Language 
type Definition struct { 
   Defclaire
   Arg *ClaireClass
  Args *ClaireList
  } 

// automatic cast function
func To_Definition(x *ClaireAny) *Definition {return (*Definition)(unsafe.Pointer(x))}

// automatic constructor function
func Make_Definition(arg *ClaireClass,args *ClaireList) *Definition { 
  var o *Definition = new(Definition)
  o.Isa = C_Definition
  o.Arg = arg
  o.Args = args
  return o 
  } 

// class file for Language/DefFast in module Language 
type Language_DefFast struct { 
   Definition
   } 

// automatic cast function
func To_Language_DefFast(x *ClaireAny) *Language_DefFast {return (*Language_DefFast)(unsafe.Pointer(x))}

// automatic constructor function
func Make_Language_DefFast(arg *ClaireClass,args *ClaireList) *Language_DefFast { 
  var o *Language_DefFast = new(Language_DefFast)
  o.Isa = C_Language_DefFast
  o.Arg = arg
  o.Args = args
  return o 
  } 

// class file for Defobj in module Language 
type Defobj struct { 
   Definition
   Ident *ClaireSymbol
  } 

// automatic cast function
func To_Defobj(x *ClaireAny) *Defobj {return (*Defobj)(unsafe.Pointer(x))}

// automatic constructor function
func Make_Defobj(arg *ClaireClass,args *ClaireList,ident *ClaireSymbol) *Defobj { 
  var o *Defobj = new(Defobj)
  o.Isa = C_Defobj
  o.Arg = arg
  o.Args = args
  o.Ident = ident
  return o 
  } 

// class file for Defclass in module Language 
type Defclass struct { 
   Defobj
   Params *ClaireList
  Forward_ask *ClaireBoolean
  } 

// automatic cast function
func To_Defclass(x *ClaireAny) *Defclass {return (*Defclass)(unsafe.Pointer(x))}

// class file for Defmethod in module Language 
type Defmethod struct { 
   Defclaire
   Arg *Call
  SetArg *ClaireAny
  Body *ClaireAny
  Inline_ask *ClaireBoolean
  } 

// automatic cast function
func To_Defmethod(x *ClaireAny) *Defmethod {return (*Defmethod)(unsafe.Pointer(x))}

// automatic constructor function
func Make_Defmethod(arg *Call,set_arg *ClaireAny,body *ClaireAny,inline_ask *ClaireBoolean) *Defmethod { 
  var o *Defmethod = new(Defmethod)
  o.Isa = C_Defmethod
  o.Arg = arg
  o.SetArg = set_arg
  o.Body = body
  o.Inline_ask = inline_ask
  return o 
  } 

// class file for Defarray in module Language 
type Defarray struct { 
   Defmethod
   } 

// automatic cast function
func To_Defarray(x *ClaireAny) *Defarray {return (*Defarray)(unsafe.Pointer(x))}

// automatic constructor function
func Make_Defarray(arg *Call,set_arg *ClaireAny,body *ClaireAny,inline_ask *ClaireBoolean) *Defarray { 
  var o *Defarray = new(Defarray)
  o.Isa = C_Defarray
  o.Arg = arg
  o.SetArg = set_arg
  o.Body = body
  o.Inline_ask = inline_ask
  return o 
  } 

// class file for Defrule in module Language 
type Defrule struct { 
   Defclaire
   Ident *ClaireSymbol
  Args *ClaireList
  Arg *ClaireAny
  Body *ClaireAny
  } 

// automatic cast function
func To_Defrule(x *ClaireAny) *Defrule {return (*Defrule)(unsafe.Pointer(x))}

// automatic constructor function
func Make_Defrule(ident *ClaireSymbol,args *ClaireList,arg *ClaireAny,body *ClaireAny) *Defrule { 
  var o *Defrule = new(Defrule)
  o.Isa = C_Defrule
  o.Ident = ident
  o.Args = args
  o.Arg = arg
  o.Body = body
  return o 
  } 

// class file for Defvar in module Language 
type Defvar struct { 
   Defclaire
   Ident *ClaireVariable
  Arg *ClaireAny
  } 

// automatic cast function
func To_Defvar(x *ClaireAny) *Defvar {return (*Defvar)(unsafe.Pointer(x))}

// automatic constructor function
func Make_Defvar(ident *ClaireVariable,arg *ClaireAny) *Defvar { 
  var o *Defvar = new(Defvar)
  o.Isa = C_Defvar
  o.Ident = ident
  o.Arg = arg
  return o 
  } 

// class file for Language/demon in module Language 
type LanguageDemon struct { 
   ClaireLambda
   Pname *ClaireSymbol
  Priority int
  Formula *ClaireLambda
  } 

// automatic cast function
func ToLanguageDemon(x *ClaireAny) *LanguageDemon {return (*LanguageDemon)(unsafe.Pointer(x))}

// class file for Language/rule_object in module Language 
type LanguageRuleObject struct { 
   ClaireProperty
   } 

// automatic cast function
func ToLanguageRuleObject(x *ClaireAny) *LanguageRuleObject {return (*LanguageRuleObject)(unsafe.Pointer(x))}

var C_Basic_instruction *ClaireClass /*obj*/
var C_iClaire_typing *Core.GlobalVariable
var C_iClaire_index *Core.GlobalVariable
var C_Vardef *ClaireClass /*obj*/
var C_Complex_instruction *ClaireClass /*obj*/
var C_Instruction_with_var *ClaireClass /*obj*/
var C_Control_structure *ClaireClass /*obj*/
var C_EOF *Core.GlobalVariable
var C_EOS *Core.GlobalVariable
var C_MAX_INTEGER *Core.GlobalVariable
var C__starvariable_index_star *Core.GlobalVariable
var C_printl *ClaireProperty /*obj*/
var C_Language_PPC *Core.GlobalVariable
var C_iClaire_LastCall *Core.GlobalVariable
var C_Call *ClaireClass /*obj*/
var C_Call_star *ClaireClass /*obj*/
var C_Call_plus *ClaireClass /*obj*/
var C_Assign *ClaireClass /*obj*/
var C_Gassign *ClaireClass /*obj*/
var C_And *ClaireClass /*obj*/
var C_Or *ClaireClass /*obj*/
var C_Quote *ClaireClass /*obj*/
var C_Optimized_instruction *ClaireClass /*obj*/
var C_Call_method *ClaireClass /*obj*/
var C_Call_method1 *ClaireClass /*obj*/
var C_Call_method2 *ClaireClass /*obj*/
var C_Language_Call_method3 *ClaireClass /*obj*/
var C_Call_slot *ClaireClass /*obj*/
var C_Call_array *ClaireClass /*obj*/
var C_Call_table *ClaireClass /*obj*/
var C_Update *ClaireClass /*obj*/
var C_Super *ClaireClass /*obj*/
var C_Cast *ClaireClass /*obj*/
var C_Return *ClaireClass /*obj*/
var C_If *ClaireClass /*obj*/
var C_If_ask *ClaireClass /*obj*/
var C_Do *ClaireClass /*obj*/
var C_Let *ClaireClass /*obj*/
var C_When *ClaireClass /*obj*/
var C_Let_plus *ClaireClass /*obj*/
var C_Let_star *ClaireClass /*obj*/
var C_Iteration *ClaireClass /*obj*/
var C_iterate *ClaireProperty /*obj*/
var C_Iterate *ClaireProperty /*obj*/
var C_IterateFast *ClaireProperty /*obj*/
var C_For *ClaireClass /*obj*/
var C_Collect *ClaireClass /*obj*/
var C_Image *ClaireClass /*obj*/
var C_Select *ClaireClass /*obj*/
var C_Lselect *ClaireClass /*obj*/
var C_Exists *ClaireClass /*obj*/
var C_Case *ClaireClass /*obj*/
var C_While *ClaireClass /*obj*/
var C_Handle *ClaireClass /*obj*/
var C_Construct *ClaireClass /*obj*/
var C_List *ClaireClass /*obj*/
var C_Tuple *ClaireClass /*obj*/
var C_Set *ClaireClass /*obj*/
var C_Array *ClaireClass /*obj*/
var C_Printf *ClaireClass /*obj*/
var C_Error *ClaireClass /*obj*/
var C_Branch *ClaireClass /*obj*/
var C_Map *ClaireClass /*obj*/
var C_Macro *ClaireClass /*obj*/
var C_macroexpand *ClaireProperty /*obj*/
var C_iClaire_trace_on *ClaireProperty /*obj*/
var C_Trace *ClaireClass /*obj*/
var C_Assert *ClaireClass /*obj*/
var C_extract_item *ClaireProperty /*obj*/
var C_function_I *ClaireProperty /*obj*/
var C_iClaire_LastComment *Core.GlobalVariable
var C_NeedComment *Core.GlobalVariable
var C_Defclaire *ClaireClass /*obj*/
var C_Definition *ClaireClass /*obj*/
var C_Language_DefFast *ClaireClass /*obj*/
var C_Defobj *ClaireClass /*obj*/
var C_Defclass *ClaireClass /*obj*/
var C_Defmethod *ClaireClass /*obj*/
var C_Defarray *ClaireClass /*obj*/
var C_Defrule *ClaireClass /*obj*/
var C_Defvar *ClaireClass /*obj*/
var C_LDEF *Core.GlobalVariable
var C_bit_vector *ClaireProperty /*obj*/
var C_Language_demon *ClaireClass /*obj*/
var C_demons *ClaireTable /*obj*/
var C__inf_dash *ClaireOperation /*obj*/
var C_Language_rule_object *ClaireClass /*obj*/
var C_Language_relations *ClaireTable /*obj*/
var C_Language_last_rule *ClaireTable /*obj*/
var C_eval_rule *ClaireProperty /*obj*/
var C_Language_putCall *ClaireProperty // Language/"putCall"
var C_Language_safeRange *ClaireProperty // Language/"safeRange"
var C_imported_function *ClaireProperty // claire/"imported_function"
var C_occurchange *ClaireProperty // claire/"occurchange"
var C_Language_getDefault *ClaireProperty // Language/"getDefault"
var C_Language_fast_definition_ask *ClaireProperty // Language/"fast_definition?"
var C_static_type *ClaireProperty // claire/"static_type"
var C_Language_static_type_nth *ClaireProperty // Language/"static_type_nth"
var C_Language_jito *ClaireProperty // Language/"jito"
var C_Language_makeJito *ClaireProperty // Language/"makeJito"
var C_Language_letJito *ClaireProperty // Language/"letJito"
var C_Language_makeCallMatch *ClaireProperty // Language/"makeCallMatch"
var C_Language_new_writes *ClaireProperty // Language/"new_writes"
var C_Language_occurexact *ClaireProperty // Language/"occurexact"
var C_map *ClaireProperty // claire/"map"
var C_Language_no_eval *ClaireProperty // Language/"no_eval"
var C_Language_ppvariable *ClaireProperty // Language/"ppvariable"
var C_Language_write_value *ClaireProperty // Language/"write_value"
var C_var *ClaireProperty // claire/"var"
var C_iClaire_lambda_I *ClaireProperty // iClaire/"lambda!"
var C_iClaire_lexical_change *ClaireProperty // iClaire/"lexical_change"
var C_iClaire_extract_symbol *ClaireProperty // iClaire/"extract_symbol"
var C_iClaire_make_a_property *ClaireProperty // iClaire/"make_a_property"
var C_Language_lbreak *ClaireProperty // Language/"lbreak"
var C_Language_put_buffer *ClaireProperty // Language/"put_buffer"
var C_Language_checkfar *ClaireProperty // Language/"checkfar"
var C_Language_indent *ClaireProperty // Language/"indent"
var C_Language_set_level *ClaireProperty // Language/"set_level"
var C_Language_printbox *ClaireProperty // Language/"printbox"
var C_Language_printexp *ClaireProperty // Language/"printexp"
var C_pretty_print *ClaireProperty // claire/"pretty_print"
var C_Language_assign *ClaireProperty // Language/"assign"
var C_Language_printe *ClaireProperty // Language/"printe"
var C_Language_sugar_ask *ClaireProperty // Language/"sugar?"
var C_iClaire_cast_to *ClaireProperty // iClaire/"cast_to"
var C_iClaire_set_arg *ClaireProperty // iClaire/"set_arg"
var C_substitution *ClaireProperty // claire/"substitution"
var C_Language_occurrence *ClaireProperty // Language/"occurrence"
var C_Language_instruction_copy *ClaireProperty // Language/"instruction_copy"
var C_iClaire_other *ClaireProperty // iClaire/"other"
var C_iClaire_test *ClaireProperty // iClaire/"test"
var C_Language_printstat *ClaireProperty // Language/"printstat"
var C_Language_printif *ClaireProperty // Language/"printif"
var C_Language_printelse *ClaireProperty // Language/"printelse"
var C_Language_printdo *ClaireProperty // Language/"printdo"
var C_Language_printblock *ClaireProperty // Language/"printblock"
var C_Language_printbody *ClaireProperty // Language/"printbody"
var C_iClaire_ident *ClaireProperty // iClaire/"ident"
var C_Language_attach_comment *ClaireProperty // Language/"attach_comment"
var C_iClaire_extract_signature *ClaireProperty // iClaire/"extract_signature"
var C_iClaire_extract_pattern *ClaireProperty // iClaire/"extract_pattern"
var C_iClaire_extract_type *ClaireProperty // iClaire/"extract_type"
var C_Language_extract_pattern_nth *ClaireProperty // Language/"extract_pattern_nth"
var C_iClaire_extract_class_call *ClaireProperty // iClaire/"extract_class_call"
var C_iClaire_extract_range *ClaireProperty // iClaire/"extract_range"
var C_iClaire_extract_status *ClaireProperty // iClaire/"extract_status"
var C_iClaire_type_I *ClaireProperty // iClaire/"type!"
var C_iClaire_forward_ask *ClaireProperty // iClaire/"forward?"
var C_Language_priority *ClaireProperty // Language/"priority"
var C_Language_make_filter *ClaireProperty // Language/"make_filter"
var C_Language_make_demon *ClaireProperty // Language/"make_demon"
var C_Language_eval_if_write *ClaireProperty // Language/"eval_if_write"
var C_Language_readCall *ClaireProperty // Language/"readCall"
var C_Language_eventMethod *ClaireProperty // Language/"eventMethod"
var C_Language_eventMethod_ask *ClaireProperty // Language/"eventMethod?"
var C_iClaire_lexical_index *ClaireProperty // iClaire/"lexical_index"
var C_Language_occurbreak *ClaireProperty // Language/"occurbreak"
var It *ClaireModule
var C_iClaire *ClaireModule 
// definition of the meta-model for module Language 
func MetaLoad() { 
  
  It = MakeModule("Language",C_iClaire)
  It.Comment = MakeString("Compiled on Friday 01-03-2025 14:52:12(v4.1.4), lines:2275, warnings:1,safety:5")
  ClEnv.Module_I = It
  
  // definition of the properties
  C_Language_putCall = MakeProperty("putCall",1,It)
  C_Language_safeRange = MakeProperty("safeRange",1,It)
  C_imported_function = MakeProperty("imported_function",1,C_claire)
  C_occurchange = MakeProperty("occurchange",1,C_claire)
  C_Language_getDefault = MakeProperty("getDefault",1,It)
  C_Language_fast_definition_ask = MakeProperty("fast_definition?",1,It)
  C_static_type = MakeProperty("static_type",1,C_claire)
  C_Language_static_type_nth = MakeProperty("static_type_nth",1,It)
  C_Language_jito = MakeProperty("jito",1,It)
  C_Language_makeJito = MakeProperty("makeJito",1,It)
  C_Language_letJito = MakeProperty("letJito",1,It)
  C_Language_makeCallMatch = MakeProperty("makeCallMatch",1,It)
  C_Language_new_writes = MakeProperty("new_writes",1,It)
  C_Language_occurexact = MakeProperty("occurexact",1,It)
  C_map = MakeProperty("map",1,C_claire)
  C_Language_no_eval = MakeProperty("no_eval",1,It)
  C_Language_ppvariable = MakeProperty("ppvariable",1,It)
  C_Language_write_value = MakeProperty("write_value",1,It)
  C_var = MakeProperty("var",0,C_claire)
  C_iClaire_lambda_I = MakeProperty("lambda!",1,C_iClaire)
  C_iClaire_lexical_change = MakeProperty("lexical_change",1,C_iClaire)
  C_iClaire_extract_symbol = MakeProperty("extract_symbol",1,C_iClaire)
  C_iClaire_make_a_property = MakeProperty("make_a_property",1,C_iClaire)
  C_Language_lbreak = MakeProperty("lbreak",1,It)
  C_Language_put_buffer = MakeProperty("put_buffer",1,It)
  C_Language_checkfar = MakeProperty("checkfar",1,It)
  C_Language_indent = MakeProperty("indent",1,It)
  C_Language_set_level = MakeProperty("set_level",1,It)
  C_Language_printbox = MakeProperty("printbox",1,It)
  C_Language_printexp = MakeProperty("printexp",1,It)
  C_pretty_print = MakeProperty("pretty_print",1,C_claire)
  C_Language_assign = MakeProperty("assign",1,It)
  C_Language_printe = MakeProperty("printe",1,It)
  C_Language_sugar_ask = MakeProperty("sugar?",1,It)
  C_iClaire_cast_to = MakeProperty("cast_to",1,C_iClaire)
  C_iClaire_set_arg = MakeProperty("set_arg",0,C_iClaire)
  C_substitution = MakeProperty("substitution",1,C_claire)
  C_Language_occurrence = MakeProperty("occurrence",1,It)
  C_Language_instruction_copy = MakeProperty("instruction_copy",1,It)
  C_iClaire_other = MakeProperty("other",1,C_iClaire)
  C_iClaire_test = MakeProperty("test",1,C_iClaire)
  C_Language_printstat = MakeProperty("printstat",1,It)
  C_Language_printif = MakeProperty("printif",1,It)
  C_Language_printelse = MakeProperty("printelse",1,It)
  C_Language_printdo = MakeProperty("printdo",1,It)
  C_Language_printblock = MakeProperty("printblock",1,It)
  C_Language_printbody = MakeProperty("printbody",1,It)
  C_iClaire_ident = MakeProperty("ident",1,C_iClaire)
  C_Language_attach_comment = MakeProperty("attach_comment",1,It)
  C_iClaire_extract_signature = MakeProperty("extract_signature",1,C_iClaire)
  C_iClaire_extract_pattern = MakeProperty("extract_pattern",1,C_iClaire)
  C_iClaire_extract_type = MakeProperty("extract_type",1,C_iClaire)
  C_Language_extract_pattern_nth = MakeProperty("extract_pattern_nth",1,It)
  C_iClaire_extract_class_call = MakeProperty("extract_class_call",1,C_iClaire)
  C_iClaire_extract_range = MakeProperty("extract_range",1,C_iClaire)
  C_iClaire_extract_status = MakeProperty("extract_status",1,C_iClaire)
  C_iClaire_type_I = MakeProperty("type!",1,C_iClaire)
  C_iClaire_forward_ask = MakeProperty("forward?",1,C_iClaire)
  C_Language_priority = MakeProperty("priority",2,It)
  C_Language_make_filter = MakeProperty("make_filter",1,It)
  C_Language_make_demon = MakeProperty("make_demon",1,It)
  C_Language_eval_if_write = MakeProperty("eval_if_write",1,It)
  C_Language_readCall = MakeProperty("readCall",1,It)
  C_Language_eventMethod = MakeProperty("eventMethod",1,It)
  C_Language_eventMethod_ask = MakeProperty("eventMethod?",1,It)
  C_iClaire_lexical_index = MakeProperty("lexical_index",1,C_iClaire)
  C_Language_occurbreak = MakeProperty("occurbreak",1,It)
  
  // instructions from module sources
  C_Basic_instruction = NewClass("Basic_instruction",C_Instruction,C_claire)
  
  _ = Core.F_attach_method(C_Language_no_eval.AddMethod(Signature(C_Instruction.Id(),C_void.Id()),1,MakeFunction1(E_no_eval_Instruction,"no_eval_Instruction")),MakeString("pretty.cl:20"))
  
  { 
    var expr EID
    { var _CL_obj *Core.GlobalVariable
      C_iClaire_typing = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("typing",C_iClaire)))
      
      _CL_obj = C_iClaire_typing
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = C_Kernel_typing.Id()
      expr = Core.F_close_global_variable(_CL_obj)
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID
    { var _CL_obj *Core.GlobalVariable
      C_iClaire_index = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("index",C_iClaire)))
      
      _CL_obj = C_iClaire_index
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = C_mClaire_index.Id()
      expr = Core.F_close_global_variable(_CL_obj)
      } 
    ErrorCheck(expr)} 
  
  _ = Core.F_attach_method(C_self_print.AddMethod(Signature(C_unbound_symbol.Id(),C_void.Id()),0,MakeFunction1(E_self_print_unbound_symbol_Language,"self_print_unbound_symbol_Language")),MakeString("pretty.cl:34"))
  
  _ = Core.F_attach_method(Core.C_self_eval.AddEvalMethod(Signature(C_unbound_symbol.Id(),C_any.Id()),1,MakeFunction1(E_self_eval_unbound_symbol,"self_eval_unbound_symbol"),EVAL_unbound_symbol),MakeString("pretty.cl:37"))
  
  _ = Core.F_attach_method(C_self_print.AddMethod(Signature(C_Variable.Id(),C_void.Id()),0,MakeFunction1(E_self_print_Variable_Language,"self_print_Variable_Language")),MakeString("pretty.cl:48"))
  
  _ = Core.F_attach_method(C_Language_ppvariable.AddMethod(Signature(C_Variable.Id(),C_void.Id()),1,MakeFunction1(E_ppvariable_Variable,"ppvariable_Variable")),MakeString("pretty.cl:53"))
  
  _ = Core.F_attach_method(C_Language_ppvariable.AddMethod(Signature(C_list.Id(),C_void.Id()),1,MakeFunction1(E_ppvariable_list,"ppvariable_list")),MakeString("pretty.cl:60"))
  
  _ = Core.F_attach_method(Core.C_self_eval.AddEvalMethod(Signature(C_Variable.Id(),C_any.Id()),0,MakeFunction1(E_self_eval_Variable,"self_eval_Variable"),EVAL_Variable),MakeString("pretty.cl:62"))
  
  _ = Core.F_attach_method(C_Language_write_value.AddMethod(Signature(C_Variable.Id(),C_any.Id(),C_any.Id()),1,MakeFunction2(E_write_value_Variable,"write_value_Variable")),MakeString("pretty.cl:67"))
  
  C_Vardef = NewClass("Vardef",C_Variable,C_claire)
  
  _ = Core.F_attach_method(Core.C_self_eval.AddEvalMethod(Signature(C_Vardef.Id(),C_any.Id()),1,MakeFunction1(E_self_eval_Vardef,"self_eval_Vardef"),EVAL_Vardef),MakeString("pretty.cl:77"))
  
  C_Complex_instruction = NewClass("Complex_instruction",C_Instruction,C_claire)
  
  C_Instruction_with_var = NewClass("Instruction_with_var",C_Complex_instruction,C_claire)
  Core.F_close_slot(C_Instruction_with_var.AddSlot(C_var,ToType(C_Variable.Id()),CNULL))
  
  C_Control_structure = NewClass("Control_structure",C_Complex_instruction,C_claire)
  
  _ = Core.F_attach_method(C_Language_write_value.AddMethod(Signature(Core.C_global_variable.Id(),C_any.Id(),C_any.Id()),1,MakeFunction2(E_write_value_global_variable,"write_value_global_variable")),MakeString("pretty.cl:91"))
  
  { 
    var expr EID
    { var _CL_obj *Core.GlobalVariable
      C_EOF = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("EOF",C_claire)))
      
      _CL_obj = C_EOF
      _CL_obj.Range = ToType(C_char.Id())
      _CL_obj.Value = MakeChar(F_char_I_integer(-1)).Id()
      expr = Core.F_close_global_variable(_CL_obj)
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID
    { var _CL_obj *Core.GlobalVariable
      C_EOS = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("EOS",C_claire)))
      
      _CL_obj = C_EOS
      _CL_obj.Range = ToType(C_char.Id())
      _CL_obj.Value = MakeChar(F_char_I_integer(0)).Id()
      expr = Core.F_close_global_variable(_CL_obj)
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID
    { var _CL_obj *Core.GlobalVariable
      C_MAX_INTEGER = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("MAX_INTEGER",C_claire)))
      
      _CL_obj = C_MAX_INTEGER
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = MakeInteger(1073741822).Id()
      expr = Core.F_close_global_variable(_CL_obj)
      } 
    ErrorCheck(expr)} 
  
  _ = Core.F_attach_method(C_apply.AddMethod(Signature(C_lambda.Id(),C_list.Id(),C_any.Id()),1,MakeFunction2(E_apply_lambda,"apply_lambda")),MakeString("pretty.cl:116"))
  
  _ = Core.F_attach_method(Core.C_call.AddMethod(Signature(C_lambda.Id(),C_listargs.Id(),C_any.Id()),1,MakeFunction2(E_call_lambda2,"call_lambda2")),MakeString("pretty.cl:117"))
  
  _ = Core.F_attach_method(C_self_print.AddMethod(Signature(C_lambda.Id(),C_void.Id()),1,MakeFunction1(E_self_print_lambda_Language,"self_print_lambda_Language")),MakeString("pretty.cl:122"))
  
  _ = Core.F_attach_method(C_map.AddMethod(Signature(C_lambda.Id(),C_bag.Id(),C_any.Id()),1,MakeFunction2(E_map_lambda,"map_lambda")),MakeString("pretty.cl:128"))
  
  { 
    var expr EID
    { var _CL_obj *Core.GlobalVariable
      C__starvariable_index_star = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("*variable_index*",C_claire)))
      
      _CL_obj = C__starvariable_index_star
      _CL_obj.Range = ToType(C_integer.Id())
      _CL_obj.Value = MakeInteger(0).Id()
      expr = Core.F_close_global_variable(_CL_obj)
      } 
    ErrorCheck(expr)} 
  
  _ = Core.F_attach_method(C_iClaire_lambda_I.AddMethod(Signature(C_list.Id(),C_any.Id(),C_lambda.Id()),1,MakeFunction2(E_lambda_I_list,"lambda_I_list")),MakeString("pretty.cl:147"))
  
  _ = Core.F_attach_method(C_iClaire_lexical_change.AddMethod(Signature(C_any.Id(),C_list.Id(),C_any.Id()),1,MakeFunction2(E_lexical_change_any,"lexical_change_any")),MakeString("pretty.cl:196"))
  
  _ = Core.F_attach_method(C_iClaire_lexical_index.AddMethod(Signature(C_any.Id(),
    C_list.Id(),
    C_integer.Id(),
    C_boolean.Id(),
    C_any.Id()),1,MakeFunction4(E_iClaire_lexical_index_any2,"iClaire_lexical_index_any2")),MakeString("pretty.cl:239"))
  
  _ = Core.F_attach_method(C_close.AddMethod(Signature(C_class.Id(),C_class.Id()),0,MakeFunction1(E_close_class,"close_class")),MakeString("pretty.cl:247"))
  
  _ = Core.F_attach_method(C_iClaire_extract_symbol.AddMethod(Signature(C_any.Id(),C_symbol.Id()),1,MakeFunction1(E_extract_symbol_any,"extract_symbol_any")),MakeString("pretty.cl:260"))
  
  _ = Core.F_attach_method(C_iClaire_make_a_property.AddMethod(Signature(C_any.Id(),C_property.Id()),1,MakeFunction1(E_make_a_property_any,"make_a_property_any")),MakeString("pretty.cl:278"))
  
  C_printl = MakeProperty("printl",1,C_claire)
  
  
  { 
    var expr EID
    { var _CL_obj *Core.GlobalVariable
      C_Language_PPC = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("PPC",It)))
      
      _CL_obj = C_Language_PPC
      _CL_obj.Range = ToType(C_integer.Id())
      _CL_obj.Value = MakeInteger(0).Id()
      expr = Core.F_close_global_variable(_CL_obj)
      } 
    ErrorCheck(expr)} 
  
  _ = Core.F_attach_method(C_Language_lbreak.AddMethod(Signature(C_void.Id(),C_any.Id()),1,MakeFunction1(E_lbreak_void,"lbreak_void")),MakeString("pretty.cl:298"))
  
  _ = Core.F_attach_method(C_Language_put_buffer.AddMethod(Signature(C_void.Id(),C_any.Id()),1,MakeFunction1(E_put_buffer_void,"put_buffer_void")),MakeString("pretty.cl:302"))
  
  _ = Core.F_attach_method(C_Language_checkfar.AddMethod(Signature(C_void.Id(),C_any.Id()),1,MakeFunction1(E_checkfar_void,"checkfar_void")),MakeString("pretty.cl:306"))
  
  _ = Core.F_attach_method(C_Language_lbreak.AddMethod(Signature(C_integer.Id(),C_any.Id()),1,MakeFunction1(E_lbreak_integer,"lbreak_integer")),MakeString("pretty.cl:308"))
  
  _ = Core.F_attach_method(C_Language_indent.AddMethod(Signature(C_integer.Id(),C_any.Id()),0,MakeFunction1(E_indent_integer,"indent_integer")),MakeString("pretty.cl:313"))
  
  _ = Core.F_attach_method(C_Language_set_level.AddMethod(Signature(C_void.Id(),C_void.Id()),0,MakeFunction1(E_set_level_void,"set_level_void")),MakeString("pretty.cl:317"))
  
  _ = Core.F_attach_method(C_Language_set_level.AddMethod(Signature(C_integer.Id(),C_void.Id()),0,MakeFunction1(E_set_level_integer,"set_level_integer")),MakeString("pretty.cl:318"))
  
  _ = Core.F_attach_method(C_Language_printbox.AddMethod(Signature(C_list.Id(),
    C_integer.Id(),
    C_integer.Id(),
    C_string.Id(),
    C_any.Id()),1,MakeFunction4(E_Language_printbox_list1,"Language_printbox_list1")),MakeString("pretty.cl:354"))
  
  _ = Core.F_attach_method(C_Language_printbox.AddMethod(Signature(C_list.Id(),C_any.Id()),1,MakeFunction1(E_Language_printbox_list2,"Language_printbox_list2")),MakeString("pretty.cl:359"))
  
  _ = Core.F_attach_method(C_Language_printbox.AddMethod(Signature(C_list.Id(),C_string.Id(),C_any.Id()),1,MakeFunction2(E_Language_printbox_list3,"Language_printbox_list3")),MakeString("pretty.cl:361"))
  
  _ = Core.F_attach_method(C_printl.AddMethod(Signature(C_list.Id(),C_string.Id(),C_void.Id()),1,MakeFunction2(E_printl_list,"printl_list")),MakeString("pretty.cl:379"))
  
  _ = Core.F_attach_method(C_Language_printexp.AddMethod(Signature(C_any.Id(),C_boolean.Id(),C_void.Id()),1,MakeFunction2(E_printexp_any,"printexp_any")),MakeString("pretty.cl:392"))
  
  _ = Core.F_attach_method(C_pretty_print.AddMethod(Signature(C_any.Id(),C_void.Id()),1,MakeFunction1(E_pretty_print_any,"pretty_print_any")),MakeString("pretty.cl:402"))
  
  _ = Core.F_attach_method(C_self_print.AddMethod(Signature(C_list.Id(),C_void.Id()),1,MakeFunction1(E_self_print_list_Language,"self_print_list_Language")),MakeString("pretty.cl:407"))
  
  _ = Core.F_attach_method(C_self_print.AddMethod(Signature(C_set.Id(),C_void.Id()),1,MakeFunction1(E_self_print_set_Language,"self_print_set_Language")),MakeString("pretty.cl:412"))
  
  _ = Core.F_attach_method(C_self_print.AddMethod(Signature(C_tuple.Id(),C_void.Id()),1,MakeFunction1(E_self_print_tuple_Language,"self_print_tuple_Language")),MakeString("pretty.cl:416"))
  
  _ = Core.F_attach_method(C_self_print.AddMethod(Signature(C_map_set.Id(),C_void.Id()),1,MakeFunction1(E_self_print_map_set,"self_print_map_set")),MakeString("pretty.cl:420"))
  
  _ = Core.F_attach_method(C_self_print.AddMethod(Signature(C_pair.Id(),C_void.Id()),1,MakeFunction1(E_self_print_pair,"self_print_pair")),MakeString("pretty.cl:426"))
  
  _ = Core.F_attach_method(C_static_type.AddMethod(Signature(C_any.Id(),C_class.Id()),1,MakeFunction1(E_static_type_any,"static_type_any")),MakeString("pretty.cl:468"))
  
  _ = Core.F_attach_method(C_Language_static_type_nth.AddMethod(Signature(C_any.Id(),C_class.Id()),1,MakeFunction1(E_Language_static_type_nth_any,"Language_static_type_nth_any")),MakeString("pretty.cl:476"))
  
  { 
    var expr EID
    { var _CL_obj *Core.GlobalVariable
      C_iClaire_LastCall = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("LastCall",C_iClaire)))
      
      _CL_obj = C_iClaire_LastCall
      _CL_obj.Range = ToType(C_any.Id())
      _CL_obj.Value = CNULL
      expr = Core.F_close_global_variable(_CL_obj)
      } 
    ErrorCheck(expr)} 
  
  C_Call = NewClass("Call",C_Control_structure,C_claire)
  Core.F_close_slot(C_Call.AddSlot(C_selector,ToType(C_property.Id()),CNULL))
  Core.F_close_slot(C_Call.AddSlot(C_args,ToType(C_list.Id()),ToType(C_any.Id()).EmptyList().Id()))
  C_Call.Params = MakeList(ToType(C_any.Id()),C_selector.Id(),C_args.Id())
  
  C_Call_star = NewClass("Call*",C_Call,C_claire)
  C_Call_star.Params = MakeList(ToType(C_any.Id()),C_selector.Id(),C_args.Id())
  
  C_Call_plus = NewClass("Call+",C_Call,C_claire)
  C_Call_plus.Params = MakeList(ToType(C_any.Id()),C_selector.Id(),C_args.Id())
  
  _ = Core.F_attach_method(C_self_print.AddMethod(Signature(C_Call.Id(),C_void.Id()),1,MakeFunction1(E_self_print_Call_Language,"self_print_Call_Language")),MakeString("call.cl:68"))
  
  _ = Core.F_attach_method(C_self_print.AddMethod(Signature(C_Call_plus.Id(),C_void.Id()),1,MakeFunction1(E_self_print_Call_plus_Language,"self_print_Call_plus_Language")),MakeString("call.cl:71"))
  
  _ = Core.F_attach_method(Core.C_self_eval.AddEvalMethod(Signature(C_Call.Id(),C_any.Id()),1,MakeFunction1(E_self_eval_Call,"self_eval_Call"),EVAL_Call),MakeString("call.cl:84"))
  
  _ = Core.F_attach_method(Core.C_self_eval.AddEvalMethod(Signature(C_Call_plus.Id(),C_any.Id()),1,MakeFunction1(E_self_eval_Call_plus,"self_eval_Call_plus"),EVAL_Call_plus),MakeString("call.cl:88"))
  
  _ = Core.F_attach_method(Core.C_self_eval.AddEvalMethod(Signature(C_Call_star.Id(),C_any.Id()),1,MakeFunction1(E_self_eval_Call_star,"self_eval_Call_star"),EVAL_Call_star),MakeString("call.cl:102"))
  
  _ = Core.F_attach_method(C_Language_printe.AddMethod(Signature(C_any.Id(),C_property.Id(),C_void.Id()),1,MakeFunction2(E_printe_any,"printe_any")),MakeString("call.cl:110"))
  
  _ = Core.F_attach_method(C_Language_sugar_ask.AddMethod(Signature(C_any.Id(),
    C_any.Id(),
    C_any.Id(),
    C_any.Id(),
    C_boolean.Id()),0,MakeFunction4(E_sugar_ask_any,"sugar_ask_any")),MakeString("call.cl:120"))
  
  C_Assign = NewClass("Assign",C_Basic_instruction,C_claire)
  Core.F_close_slot(C_Assign.AddSlot(C_var,ToType(C_any.Id()),CNULL))
  Core.F_close_slot(C_Assign.AddSlot(C_arg,ToType(C_any.Id()),CNULL))
  
  _ = Core.F_attach_method(C_self_print.AddMethod(Signature(C_Assign.Id(),C_void.Id()),1,MakeFunction1(E_self_print_Assign_Language,"self_print_Assign_Language")),MakeString("call.cl:141"))
  
  _ = Core.F_attach_method(Core.C_self_eval.AddEvalMethod(Signature(C_Assign.Id(),C_any.Id()),1,MakeFunction1(E_self_eval_Assign,"self_eval_Assign"),EVAL_Assign),MakeString("call.cl:145"))
  
  C_Gassign = NewClass("Gassign",C_Basic_instruction,C_claire)
  Core.F_close_slot(C_Gassign.AddSlot(C_var,ToType(Core.C_global_variable.Id()),CNULL))
  Core.F_close_slot(C_Gassign.AddSlot(C_arg,ToType(C_any.Id()),CNULL))
  
  _ = Core.F_attach_method(C_self_print.AddMethod(Signature(C_Gassign.Id(),C_void.Id()),1,MakeFunction1(E_self_print_Gassign_Language,"self_print_Gassign_Language")),MakeString("call.cl:156"))
  
  _ = Core.F_attach_method(Core.C_self_eval.AddEvalMethod(Signature(C_Gassign.Id(),C_any.Id()),1,MakeFunction1(E_self_eval_Gassign,"self_eval_Gassign"),EVAL_Gassign),MakeString("call.cl:159"))
  
  C_And = NewClass("And",C_Control_structure,C_claire)
  Core.F_close_slot(C_And.AddSlot(C_args,ToType(C_list.Id()),ToType(C_any.Id()).EmptyList().Id()))
  
  _ = Core.F_attach_method(C_self_print.AddMethod(Signature(C_And.Id(),C_void.Id()),1,MakeFunction1(E_self_print_And_Language,"self_print_And_Language")),MakeString("call.cl:165"))
  
  _ = Core.F_attach_method(Core.C_self_eval.AddEvalMethod(Signature(C_And.Id(),C_any.Id()),1,MakeFunction1(E_self_eval_And,"self_eval_And"),EVAL_And),MakeString("call.cl:167"))
  
  C_Or = NewClass("Or",C_Control_structure,C_claire)
  Core.F_close_slot(C_Or.AddSlot(C_args,ToType(C_list.Id()),ToType(C_any.Id()).EmptyList().Id()))
  
  _ = Core.F_attach_method(C_self_print.AddMethod(Signature(C_Or.Id(),C_void.Id()),1,MakeFunction1(E_self_print_Or_Language,"self_print_Or_Language")),MakeString("call.cl:172"))
  
  _ = Core.F_attach_method(Core.C_self_eval.AddEvalMethod(Signature(C_Or.Id(),C_any.Id()),1,MakeFunction1(E_self_eval_Or,"self_eval_Or"),EVAL_Or),MakeString("call.cl:174"))
  
  C_Quote = NewClass("Quote",C_Basic_instruction,C_claire)
  Core.F_close_slot(C_Quote.AddSlot(C_arg,ToType(C_any.Id()),CNULL))
  
  _ = Core.F_attach_method(C_self_print.AddMethod(Signature(C_Quote.Id(),C_void.Id()),1,MakeFunction1(E_self_print_Quote_Language,"self_print_Quote_Language")),MakeString("call.cl:179"))
  
  _ = Core.F_attach_method(Core.C_self_eval.AddEvalMethod(Signature(C_Quote.Id(),C_any.Id()),0,MakeFunction1(E_self_eval_Quote,"self_eval_Quote"),EVAL_Quote),MakeString("call.cl:180"))
  
  C_Optimized_instruction = NewClass("Optimized_instruction",C_Complex_instruction,C_claire)
  
  C_Call_method = NewClass("Call_method",C_Optimized_instruction,C_claire)
  Core.F_close_slot(C_Call_method.AddSlot(C_arg,ToType(C_method.Id()),CNULL))
  Core.F_close_slot(C_Call_method.AddSlot(C_args,ToType(C_list.Id()),ToType(C_any.Id()).EmptyList().Id()))
  
  _ = Core.F_attach_method(C_self_print.AddMethod(Signature(C_Call_method.Id(),C_void.Id()),1,MakeFunction1(E_self_print_Call_method_Language,"self_print_Call_method_Language")),MakeString("call.cl:197"))
  
  _ = Core.F_attach_method(Core.C_self_eval.AddEvalMethod(Signature(C_Call_method.Id(),C_any.Id()),1,MakeFunction1(E_self_eval_Call_method,"self_eval_Call_method"),EVAL_Call_method),MakeString("call.cl:202"))
  
  C_Call_method1 = NewClass("Call_method1",C_Call_method,C_claire)
  
  _ = Core.F_attach_method(Core.C_self_eval.AddEvalMethod(Signature(C_Call_method1.Id(),C_any.Id()),1,MakeFunction1(E_self_eval_Call_method1,"self_eval_Call_method1"),EVAL_Call_method1),MakeString("call.cl:207"))
  
  C_Call_method2 = NewClass("Call_method2",C_Call_method,C_claire)
  
  _ = Core.F_attach_method(Core.C_self_eval.AddEvalMethod(Signature(C_Call_method2.Id(),C_any.Id()),1,MakeFunction1(E_self_eval_Call_method2,"self_eval_Call_method2"),EVAL_Call_method2),MakeString("call.cl:214"))
  
  C_Language_Call_method3 = NewClass("Call_method3",C_Call_method,It)
  
  _ = Core.F_attach_method(Core.C_self_eval.AddEvalMethod(Signature(C_Language_Call_method3.Id(),C_any.Id()),1,MakeFunction1(E_self_eval_Call_method3,"self_eval_Call_method3"),EVAL_Language_Call_method3),MakeString("call.cl:221"))
  
  C_Call_slot = NewClass("Call_slot",C_Optimized_instruction,C_claire)
  Core.F_close_slot(C_Call_slot.AddSlot(C_selector,ToType(C_slot.Id()),CNULL))
  Core.F_close_slot(C_Call_slot.AddSlot(C_arg,ToType(C_any.Id()),CNULL))
  Core.F_close_slot(C_Call_slot.AddSlot(C_iClaire_test,ToType(C_boolean.Id()),CNULL))
  
  _ = Core.F_attach_method(C_self_print.AddMethod(Signature(C_Call_slot.Id(),C_void.Id()),1,MakeFunction1(E_self_print_Call_slot_Language,"self_print_Call_slot_Language")),MakeString("call.cl:227"))
  
  _ = Core.F_attach_method(Core.C_self_eval.AddEvalMethod(Signature(C_Call_slot.Id(),C_any.Id()),1,MakeFunction1(E_self_eval_Call_slot,"self_eval_Call_slot"),EVAL_Call_slot),MakeString("call.cl:228"))
  
  C_Call_array = NewClass("Call_array",C_Optimized_instruction,C_claire)
  Core.F_close_slot(C_Call_array.AddSlot(C_selector,ToType(C_any.Id()),CNULL))
  Core.F_close_slot(C_Call_array.AddSlot(C_arg,ToType(C_any.Id()),CNULL))
  Core.F_close_slot(C_Call_array.AddSlot(C_iClaire_test,ToType(C_any.Id()),CNULL))
  
  _ = Core.F_attach_method(C_self_print.AddMethod(Signature(C_Call_array.Id(),C_void.Id()),1,MakeFunction1(E_self_print_Call_array_Language,"self_print_Call_array_Language")),MakeString("call.cl:236"))
  
  _ = Core.F_attach_method(Core.C_self_eval.AddEvalMethod(Signature(C_Call_array.Id(),C_any.Id()),1,MakeFunction1(E_self_eval_Call_array,"self_eval_Call_array"),EVAL_Call_array),MakeString("call.cl:238"))
  
  C_Call_table = NewClass("Call_table",C_Optimized_instruction,C_claire)
  Core.F_close_slot(C_Call_table.AddSlot(C_selector,ToType(C_table.Id()),CNULL))
  Core.F_close_slot(C_Call_table.AddSlot(C_arg,ToType(C_any.Id()),CNULL))
  Core.F_close_slot(C_Call_table.AddSlot(C_iClaire_test,ToType(C_boolean.Id()),CNULL))
  
  _ = Core.F_attach_method(C_self_print.AddMethod(Signature(C_Call_table.Id(),C_void.Id()),1,MakeFunction1(E_self_print_Call_table_Language,"self_print_Call_table_Language")),MakeString("call.cl:244"))
  
  _ = Core.F_attach_method(Core.C_self_eval.AddEvalMethod(Signature(C_Call_table.Id(),C_any.Id()),1,MakeFunction1(E_self_eval_Call_table,"self_eval_Call_table"),EVAL_Call_table),MakeString("call.cl:247"))
  
  C_Update = NewClass("Update",C_Optimized_instruction,C_claire)
  Core.F_close_slot(C_Update.AddSlot(C_selector,ToType(C_any.Id()),CNULL))
  Core.F_close_slot(C_Update.AddSlot(C_arg,ToType(C_any.Id()),CNULL))
  Core.F_close_slot(C_Update.AddSlot(C_value,ToType(C_any.Id()),CNULL))
  Core.F_close_slot(C_Update.AddSlot(C_var,ToType(C_any.Id()),CNULL))
  
  _ = Core.F_attach_method(C_self_print.AddMethod(Signature(C_Update.Id(),C_void.Id()),1,MakeFunction1(E_self_print_Update_Language,"self_print_Update_Language")),MakeString("call.cl:258"))
  
  _ = Core.F_attach_method(Core.C_self_eval.AddEvalMethod(Signature(C_Update.Id(),C_any.Id()),1,MakeFunction1(E_self_eval_Update,"self_eval_Update"),EVAL_Update),MakeString("call.cl:264"))
  
  C_Super = NewClass("Super",C_Control_structure,C_claire)
  Core.F_close_slot(C_Super.AddSlot(C_selector,ToType(C_property.Id()),CNULL))
  Core.F_close_slot(C_Super.AddSlot(C_iClaire_cast_to,ToType(C_type.Id()),CNULL))
  Core.F_close_slot(C_Super.AddSlot(C_args,ToType(C_list.Id()),ToType(C_any.Id()).EmptyList().Id()))
  
  _ = Core.F_attach_method(C_self_print.AddMethod(Signature(C_Super.Id(),C_void.Id()),1,MakeFunction1(E_self_print_Super_Language,"self_print_Super_Language")),MakeString("call.cl:280"))
  
  _ = Core.F_attach_method(Core.C_self_eval.AddEvalMethod(Signature(C_Super.Id(),C_any.Id()),1,MakeFunction1(E_self_eval_Super,"self_eval_Super"),EVAL_Super),MakeString("call.cl:289"))
  
  C_Cast = NewClass("Cast",C_Basic_instruction,C_claire)
  Core.F_close_slot(C_Cast.AddSlot(C_arg,ToType(C_any.Id()),CNULL))
  Core.F_close_slot(C_Cast.AddSlot(C_iClaire_set_arg,ToType(C_type.Id()),CNULL))
  
  _ = Core.F_attach_method(C_self_print.AddMethod(Signature(C_Cast.Id(),C_void.Id()),1,MakeFunction1(E_self_print_Cast_Language,"self_print_Cast_Language")),MakeString("call.cl:299"))
  
  _ = Core.F_attach_method(Core.C_self_eval.AddEvalMethod(Signature(C_Cast.Id(),C_any.Id()),1,MakeFunction1(E_self_eval_Cast,"self_eval_Cast"),EVAL_Cast),MakeString("call.cl:305"))
  
  C_Return = NewClass("Return",C_Basic_instruction,C_claire)
  Core.F_close_slot(C_Return.AddSlot(C_arg,ToType(C_any.Id()),CNULL))
  
  _ = Core.F_attach_method(C_self_print.AddMethod(Signature(C_Return.Id(),C_void.Id()),1,MakeFunction1(E_self_print_Return_Language,"self_print_Return_Language")),MakeString("call.cl:316"))
  
  _ = Core.F_attach_method(Core.C_self_eval.AddEvalMethod(Signature(C_Return.Id(),C_error.Id()),1,MakeFunction1(E_self_eval_Return,"self_eval_Return"),EVAL_Return),MakeString("call.cl:317"))
  
  _ = Core.F_attach_method(C_substitution.AddMethod(Signature(C_any.Id(),
    C_Variable.Id(),
    C_any.Id(),
    C_any.Id()),0,MakeFunction3(E_substitution_any,"substitution_any")),MakeString("call.cl:338"))
  
  _ = Core.F_attach_method(C_Language_occurrence.AddMethod(Signature(C_any.Id(),C_Variable.Id(),C_integer.Id()),0,MakeFunction2(E_occurrence_any,"occurrence_any")),MakeString("call.cl:351"))
  
  _ = Core.F_attach_method(C_occurchange.AddMethod(Signature(C_any.Id(),C_Variable.Id(),C_boolean.Id()),0,MakeFunction2(E_occurchange_any,"occurchange_any")),MakeString("call.cl:359"))
  
  _ = Core.F_attach_method(C_Language_occurexact.AddMethod(Signature(C_any.Id(),C_Variable.Id(),C_integer.Id()),0,MakeFunction2(E_Language_occurexact_any,"Language_occurexact_any")),MakeString("call.cl:375"))
  
  _ = Core.F_attach_method(C_Language_occurbreak.AddMethod(Signature(C_any.Id(),C_boolean.Id()),0,MakeFunction1(E_Language_occurbreak_any,"Language_occurbreak_any")),MakeString("call.cl:384"))
  
  _ = Core.F_attach_method(C_Language_instruction_copy.AddMethod(Signature(C_any.Id(),C_any.Id()),0,MakeFunction1(E_instruction_copy_any,"instruction_copy_any")),MakeString("call.cl:398"))
  
  C_If = NewClass("If",C_Control_structure,C_claire)
  Core.F_close_slot(C_If.AddSlot(C_iClaire_test,ToType(C_any.Id()),CNULL))
  Core.F_close_slot(C_If.AddSlot(C_arg,ToType(C_any.Id()),CNULL))
  Core.F_close_slot(C_If.AddSlot(C_iClaire_other,ToType(C_any.Id()),CFALSE.Id()))
  
  C_If_ask = NewClass("If?",C_If,C_claire)
  
  _ = Core.F_attach_method(C_self_print.AddMethod(Signature(C_If.Id(),C_void.Id()),1,MakeFunction1(E_self_print_If_Language,"self_print_If_Language")),MakeString("control.cl:25"))
  
  _ = Core.F_attach_method(C_Language_printstat.AddMethod(Signature(C_If.Id(),C_void.Id()),1,MakeFunction1(E_printstat_If,"printstat_If")),MakeString("control.cl:29"))
  
  _ = Core.F_attach_method(C_Language_printif.AddMethod(Signature(C_any.Id(),C_void.Id()),1,MakeFunction1(E_printif_any,"printif_any")),MakeString("control.cl:44"))
  
  _ = Core.F_attach_method(C_Language_printelse.AddMethod(Signature(C_If.Id(),C_void.Id()),1,MakeFunction1(E_printelse_If,"printelse_If")),MakeString("control.cl:55"))
  
  _ = Core.F_attach_method(Core.C_self_eval.AddEvalMethod(Signature(C_If.Id(),C_any.Id()),1,MakeFunction1(E_self_eval_If,"self_eval_If"),EVAL_If),MakeString("control.cl:65"))
  
  C_Do = NewClass("Do",C_Control_structure,C_claire)
  Core.F_close_slot(C_Do.AddSlot(C_args,ToType(C_list.Id()),ToType(C_any.Id()).EmptyList().Id()))
  C_Do.Params = MakeList(ToType(C_any.Id()),C_args.Id())
  
  _ = Core.F_attach_method(C_self_print.AddMethod(Signature(C_Do.Id(),C_void.Id()),1,MakeFunction1(E_self_print_Do_Language,"self_print_Do_Language")),MakeString("control.cl:75"))
  
  _ = Core.F_attach_method(C_Language_printdo.AddMethod(Signature(C_list.Id(),C_boolean.Id(),C_void.Id()),1,MakeFunction2(E_printdo_list,"printdo_list")),MakeString("control.cl:83"))
  
  _ = Core.F_attach_method(C_Language_printblock.AddMethod(Signature(C_any.Id(),C_void.Id()),1,MakeFunction1(E_printblock_any,"printblock_any")),MakeString("control.cl:86"))
  
  _ = Core.F_attach_method(Core.C_self_eval.AddEvalMethod(Signature(C_Do.Id(),C_any.Id()),1,MakeFunction1(E_self_eval_Do,"self_eval_Do"),EVAL_Do),MakeString("control.cl:90"))
  
  C_Let = NewClass("Let",C_Instruction_with_var,C_claire)
  Core.F_close_slot(C_Let.AddSlot(C_value,ToType(C_any.Id()),CNULL))
  Core.F_close_slot(C_Let.AddSlot(C_arg,ToType(C_any.Id()),CNULL))
  
  _ = Core.F_attach_method(C_self_print.AddMethod(Signature(C_Let.Id(),C_void.Id()),1,MakeFunction1(E_self_print_Let_Language,"self_print_Let_Language")),MakeString("control.cl:100"))
  
  _ = Core.F_attach_method(C_Language_printbody.AddMethod(Signature(C_Let.Id(),C_void.Id()),1,MakeFunction1(E_printbody_Let,"printbody_Let")),MakeString("control.cl:108"))
  
  _ = Core.F_attach_method(Core.C_self_eval.AddEvalMethod(Signature(C_Let.Id(),C_any.Id()),1,MakeFunction1(E_self_eval_Let,"self_eval_Let"),EVAL_Let),MakeString("control.cl:113"))
  
  C_When = NewClass("When",C_Let,C_claire)
  Core.F_close_slot(C_When.AddSlot(C_iClaire_other,ToType(C_any.Id()),CNULL))
  
  _ = Core.F_attach_method(C_self_print.AddMethod(Signature(C_When.Id(),C_void.Id()),1,MakeFunction1(E_self_print_When_Language,"self_print_When_Language")),MakeString("control.cl:125"))
  
  _ = Core.F_attach_method(Core.C_self_eval.AddEvalMethod(Signature(C_When.Id(),C_any.Id()),1,MakeFunction1(E_self_eval_When,"self_eval_When"),EVAL_When),MakeString("control.cl:132"))
  
  C_Let_plus = NewClass("Let+",C_Let,C_claire)
  
  C_Let_star = NewClass("Let*",C_Let,C_claire)
  
  _ = Core.F_attach_method(C_self_print.AddMethod(Signature(C_Let_plus.Id(),C_void.Id()),1,MakeFunction1(E_self_print_Let_plus_Language,"self_print_Let_plus_Language")),MakeString("control.cl:150"))
  
  _ = Core.F_attach_method(C_self_print.AddMethod(Signature(C_Let_star.Id(),C_void.Id()),1,MakeFunction1(E_self_print_Let_star_Language,"self_print_Let_star_Language")),MakeString("control.cl:175"))
  
  C_Iteration = NewClass("Iteration",C_Instruction_with_var,C_claire)
  Core.F_close_slot(C_Iteration.AddSlot(C_iClaire_set_arg,ToType(C_any.Id()),CNULL))
  Core.F_close_slot(C_Iteration.AddSlot(C_arg,ToType(C_any.Id()),CNULL))
  
  C_iterate = MakeProperty("iterate",2,C_claire)
  
  
  C_Iterate = MakeProperty("Iterate",2,C_claire)
  
  
  C_IterateFast = MakeProperty("IterateFast",2,C_claire)
  
  
  C_For = NewClass("For",C_Iteration,C_claire)
  
  _ = Core.F_attach_method(C_self_print.AddMethod(Signature(C_For.Id(),C_void.Id()),1,MakeFunction1(E_self_print_For_Language,"self_print_For_Language")),MakeString("control.cl:197"))
  
  _ = Core.F_attach_method(Core.C_self_eval.AddEvalMethod(Signature(C_For.Id(),C_any.Id()),1,MakeFunction1(E_self_eval_For,"self_eval_For"),EVAL_For),MakeString("control.cl:214"))
  
  C_Collect = NewClass("Collect",C_Iteration,C_claire)
  Core.F_close_slot(C_Collect.AddSlot(C_of,ToType(C_type.Id()),CNULL))
  
  _ = Core.F_attach_method(C_self_print.AddMethod(Signature(C_Collect.Id(),C_void.Id()),1,MakeFunction1(E_self_print_Collect_Language,"self_print_Collect_Language")),MakeString("control.cl:228"))
  
  _ = Core.F_attach_method(Core.C_self_eval.AddEvalMethod(Signature(C_Collect.Id(),C_any.Id()),1,MakeFunction1(E_self_eval_Collect,"self_eval_Collect"),EVAL_Collect),MakeString("control.cl:244"))
  
  C_Image = NewClass("Image",C_Iteration,C_claire)
  Core.F_close_slot(C_Image.AddSlot(C_of,ToType(C_type.Id()),CNULL))
  
  _ = Core.F_attach_method(C_self_print.AddMethod(Signature(C_Image.Id(),C_void.Id()),1,MakeFunction1(E_self_print_Image_Language,"self_print_Image_Language")),MakeString("control.cl:257"))
  
  _ = Core.F_attach_method(Core.C_self_eval.AddEvalMethod(Signature(C_Image.Id(),C_any.Id()),1,MakeFunction1(E_self_eval_Image,"self_eval_Image"),EVAL_Image),MakeString("control.cl:263"))
  
  C_Select = NewClass("Select",C_Iteration,C_claire)
  Core.F_close_slot(C_Select.AddSlot(C_of,ToType(C_type.Id()),CNULL))
  
  _ = Core.F_attach_method(C_self_print.AddMethod(Signature(C_Select.Id(),C_void.Id()),1,MakeFunction1(E_self_print_Select_Language,"self_print_Select_Language")),MakeString("control.cl:275"))
  
  _ = Core.F_attach_method(Core.C_self_eval.AddEvalMethod(Signature(C_Select.Id(),C_any.Id()),1,MakeFunction1(E_self_eval_Select,"self_eval_Select"),EVAL_Select),MakeString("control.cl:291"))
  
  C_Lselect = NewClass("Lselect",C_Iteration,C_claire)
  Core.F_close_slot(C_Lselect.AddSlot(C_of,ToType(C_type.Id()),CNULL))
  
  _ = Core.F_attach_method(C_self_print.AddMethod(Signature(C_Lselect.Id(),C_void.Id()),1,MakeFunction1(E_self_print_Lselect_Language,"self_print_Lselect_Language")),MakeString("control.cl:304"))
  
  _ = Core.F_attach_method(Core.C_self_eval.AddEvalMethod(Signature(C_Lselect.Id(),C_any.Id()),1,MakeFunction1(E_self_eval_Lselect,"self_eval_Lselect"),EVAL_Lselect),MakeString("control.cl:320"))
  
  C_Exists = NewClass("Exists",C_Iteration,C_claire)
  Core.F_close_slot(C_Exists.AddSlot(C_iClaire_other,ToType(C_any.Id()),CFALSE.Id()))
  
  _ = Core.F_attach_method(C_self_print.AddMethod(Signature(C_Exists.Id(),C_void.Id()),1,MakeFunction1(E_self_print_Exists_Language,"self_print_Exists_Language")),MakeString("control.cl:337"))
  
  _ = Core.F_attach_method(Core.C_self_eval.AddEvalMethod(Signature(C_Exists.Id(),C_any.Id()),1,MakeFunction1(E_self_eval_Exists,"self_eval_Exists"),EVAL_Exists),MakeString("control.cl:355"))
  
  C_Case = NewClass("Case",C_Control_structure,C_claire)
  Core.F_close_slot(C_Case.AddSlot(C_var,ToType(C_any.Id()),CNULL))
  Core.F_close_slot(C_Case.AddSlot(C_args,ToType(C_list.Id()),ToType(C_any.Id()).EmptyList().Id()))
  
  _ = Core.F_attach_method(C_self_print.AddMethod(Signature(C_Case.Id(),C_void.Id()),1,MakeFunction1(E_self_print_Case_Language,"self_print_Case_Language")),MakeString("control.cl:380"))
  
  _ = Core.F_attach_method(Core.C_self_eval.AddEvalMethod(Signature(C_Case.Id(),C_any.Id()),1,MakeFunction1(E_self_eval_Case,"self_eval_Case"),EVAL_Case),MakeString("control.cl:392"))
  
  C_While = NewClass("While",C_Control_structure,C_claire)
  Core.F_close_slot(C_While.AddSlot(C_iClaire_test,ToType(C_any.Id()),CNULL))
  Core.F_close_slot(C_While.AddSlot(C_arg,ToType(C_any.Id()),CNULL))
  Core.F_close_slot(C_While.AddSlot(C_iClaire_other,ToType(C_boolean.Id()),CFALSE.Id()))
  
  _ = Core.F_attach_method(C_self_print.AddMethod(Signature(C_While.Id(),C_void.Id()),1,MakeFunction1(E_self_print_While_Language,"self_print_While_Language")),MakeString("control.cl:401"))
  
  _ = Core.F_attach_method(Core.C_self_eval.AddEvalMethod(Signature(C_While.Id(),C_any.Id()),1,MakeFunction1(E_self_eval_While,"self_eval_While"),EVAL_While),MakeString("control.cl:409"))
  
  C_Handle = NewClass("Handle",C_Control_structure,C_claire)
  Core.F_close_slot(C_Handle.AddSlot(C_iClaire_test,ToType(C_any.Id()),CNULL))
  Core.F_close_slot(C_Handle.AddSlot(C_arg,ToType(C_any.Id()),CNULL))
  Core.F_close_slot(C_Handle.AddSlot(C_iClaire_other,ToType(C_any.Id()),CNULL))
  
  _ = Core.F_attach_method(C_self_print.AddMethod(Signature(C_Handle.Id(),C_void.Id()),1,MakeFunction1(E_self_print_Handle_Language,"self_print_Handle_Language")),MakeString("control.cl:420"))
  
  _ = Core.F_attach_method(Core.C_self_eval.AddEvalMethod(Signature(C_Handle.Id(),C_any.Id()),1,MakeFunction1(E_self_eval_Handle,"self_eval_Handle"),EVAL_Handle),MakeString("control.cl:435"))
  
  C_Construct = NewClass("Construct",C_Complex_instruction,C_claire)
  Core.F_close_slot(C_Construct.AddSlot(C_args,ToType(C_list.Id()),ToType(C_any.Id()).EmptyList().Id()))
  
  C_List = NewClass("List",C_Construct,C_claire)
  Core.F_close_slot(C_List.AddSlot(C_of,ToType(C_type.Id()),CNULL))
  
  C_Tuple = NewClass("Tuple",C_Construct,C_claire)
  
  C_Set = NewClass("Set",C_Construct,C_claire)
  Core.F_close_slot(C_Set.AddSlot(C_of,ToType(C_type.Id()),CNULL))
  
  C_Array = NewClass("Array",C_Construct,C_claire)
  Core.F_close_slot(C_Array.AddSlot(C_of,ToType(C_type.Id()),CNULL))
  
  C_Printf = NewClass("Printf",C_Construct,C_claire)
  
  C_Error = NewClass("Error",C_Construct,C_claire)
  
  C_Branch = NewClass("Branch",C_Construct,C_claire)
  
  C_Map = NewClass("Map",C_Construct,C_claire)
  Core.F_close_slot(C_Map.AddSlot(C_domain,ToType(C_type.Id()),CNULL))
  Core.F_close_slot(C_Map.AddSlot(C_of,ToType(C_type.Id()),CNULL))
  
  _ = Core.F_attach_method(C_self_print.AddMethod(Signature(C_Construct.Id(),C_void.Id()),1,MakeFunction1(E_self_print_Construct_Language,"self_print_Construct_Language")),MakeString("control.cl:470"))
  
  _ = Core.F_attach_method(Core.C_self_eval.AddEvalMethod(Signature(C_List.Id(),C_any.Id()),1,MakeFunction1(E_self_eval_List,"self_eval_List"),EVAL_List),MakeString("control.cl:483"))
  
  _ = Core.F_attach_method(Core.C_self_eval.AddEvalMethod(Signature(C_Set.Id(),C_any.Id()),1,MakeFunction1(E_self_eval_Set,"self_eval_Set"),EVAL_Set),MakeString("control.cl:494"))
  
  _ = Core.F_attach_method(Core.C_self_eval.AddEvalMethod(Signature(C_Tuple.Id(),C_any.Id()),1,MakeFunction1(E_self_eval_Tuple,"self_eval_Tuple"),EVAL_Tuple),MakeString("control.cl:503"))
  
  _ = Core.F_attach_method(Core.C_self_eval.AddEvalMethod(Signature(C_Array.Id(),C_any.Id()),1,MakeFunction1(E_self_eval_Array2,"self_eval_Array2"),EVAL_Array),MakeString("control.cl:514"))
  
  _ = Core.F_attach_method(Core.C_self_eval.AddEvalMethod(Signature(C_Map.Id(),C_map_set.Id()),1,MakeFunction1(E_self_eval_Map,"self_eval_Map"),EVAL_Map),MakeString("control.cl:522"))
  
  C_Macro = NewClass("Macro",C_Construct,C_claire)
  
  C_macroexpand = MakeProperty("macroexpand",3,C_claire)
  C_macroexpand.Open = 3
  
  
  _ = Core.F_attach_method(Core.C_self_eval.AddEvalMethod(Signature(C_Macro.Id(),C_any.Id()),1,MakeFunction1(E_self_eval_Macro2,"self_eval_Macro2"),EVAL_Macro),MakeString("control.cl:530"))
  
  _ = Core.F_attach_method(Core.C_self_eval.AddEvalMethod(Signature(C_Error.Id(),C_error.Id()),1,MakeFunction1(E_self_eval_Error,"self_eval_Error"),EVAL_Error),MakeString("control.cl:540"))
  
  _ = Core.F_attach_method(Core.C_self_eval.AddEvalMethod(Signature(C_Printf.Id(),C_any.Id()),1,MakeFunction1(E_self_eval_Printf,"self_eval_Printf"),EVAL_Printf),MakeString("control.cl:574"))
  
  C_iClaire_trace_on = MakeProperty("trace_on",1,C_iClaire)
  
  
  C_Trace = NewClass("Trace",C_Construct,C_claire)
  
  _ = Core.F_attach_method(Core.C_self_eval.AddEvalMethod(Signature(C_Trace.Id(),C_any.Id()),1,MakeFunction1(E_self_eval_Trace,"self_eval_Trace"),EVAL_Trace),MakeString("control.cl:598"))
  
  C_Assert = NewClass("Assert",C_Construct,C_claire)
  Core.F_close_slot(C_Assert.AddSlot(C_mClaire_index,ToType(C_integer.Id()),MakeInteger(0).Id()))
  Core.F_close_slot(C_Assert.AddSlot(C_external,ToType(C_string.Id()),CNULL))
  
  _ = Core.F_attach_method(Core.C_self_eval.AddEvalMethod(Signature(C_Assert.Id(),C_any.Id()),1,MakeFunction1(E_self_eval_Assert,"self_eval_Assert"),EVAL_Assert),MakeString("control.cl:612"))
  
  _ = Core.F_attach_method(Core.C_self_eval.AddEvalMethod(Signature(C_Branch.Id(),C_any.Id()),1,MakeFunction1(E_self_eval_Branch,"self_eval_Branch"),EVAL_Branch),MakeString("control.cl:618"))
  
  C_extract_item = MakeProperty("extract_item",1,C_claire)
  
  
  C_function_I = MakeProperty("function!",2,C_claire)
  
  
  { 
    var expr EID
    { var _CL_obj *Core.GlobalVariable
      C_iClaire_LastComment = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("LastComment",C_iClaire)))
      
      _CL_obj = C_iClaire_LastComment
      _CL_obj.Range = ToType(C_any.Id())
      _CL_obj.Value = CNULL
      expr = Core.F_close_global_variable(_CL_obj)
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID
    { var _CL_obj *Core.GlobalVariable
      C_NeedComment = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("NeedComment",C_claire)))
      
      _CL_obj = C_NeedComment
      _CL_obj.Range = ToType(C_boolean.Id())
      _CL_obj.Value = CFALSE.Id()
      expr = Core.F_close_global_variable(_CL_obj)
      } 
    ErrorCheck(expr)} 
  
  C_Defclaire = NewClass("Defclaire",C_Complex_instruction,C_claire)
  
  C_Definition = NewClass("Definition",C_Defclaire,C_claire)
  Core.F_close_slot(C_Definition.AddSlot(C_arg,ToType(C_class.Id()),CNULL))
  Core.F_close_slot(C_Definition.AddSlot(C_args,ToType(C_list.Id()),ToType(C_any.Id()).EmptyList().Id()))
  
  _ = Core.F_attach_method(C_self_print.AddMethod(Signature(C_Definition.Id(),C_any.Id()),1,MakeFunction1(E_self_print_Definition_Language,"self_print_Definition_Language")),MakeString("define.cl:36"))
  
  C_Language_DefFast = NewClass("DefFast",C_Definition,It)
  
  C_Defobj = NewClass("Defobj",C_Definition,C_claire)
  Core.F_close_slot(C_Defobj.AddSlot(C_iClaire_ident,ToType(C_symbol.Id()),CNULL))
  
  _ = Core.F_attach_method(C_self_print.AddMethod(Signature(C_Defobj.Id(),C_void.Id()),1,MakeFunction1(E_self_print_Defobj_Language,"self_print_Defobj_Language")),MakeString("define.cl:54"))
  
  C_Defclass = NewClass("Defclass",C_Defobj,C_claire)
  Core.F_close_slot(C_Defclass.AddSlot(C_params,ToType(C_list.Id()),ToType(C_any.Id()).EmptyList().Id()))
  Core.F_close_slot(C_Defclass.AddSlot(C_iClaire_forward_ask,ToType(C_boolean.Id()),CNULL))
  
  _ = Core.F_attach_method(C_self_print.AddMethod(Signature(C_Defclass.Id(),C_void.Id()),1,MakeFunction1(E_self_print_Defclass_Language,"self_print_Defclass_Language")),MakeString("define.cl:75"))
  
  C_Defmethod = NewClass("Defmethod",C_Defclaire,C_claire)
  Core.F_close_slot(C_Defmethod.AddSlot(C_arg,ToType(C_Call.Id()),CNULL))
  Core.F_close_slot(C_Defmethod.AddSlot(C_iClaire_set_arg,ToType(C_any.Id()),CNULL))
  Core.F_close_slot(C_Defmethod.AddSlot(C_body,ToType(C_any.Id()),CNULL))
  Core.F_close_slot(C_Defmethod.AddSlot(C_inline_ask,ToType(C_boolean.Id()),CNULL))
  
  _ = Core.F_attach_method(C_self_print.AddMethod(Signature(C_Defmethod.Id(),C_void.Id()),1,MakeFunction1(E_self_print_Defmethod_Language,"self_print_Defmethod_Language")),MakeString("define.cl:88"))
  
  C_Defarray = NewClass("Defarray",C_Defmethod,C_claire)
  
  _ = Core.F_attach_method(C_self_print.AddMethod(Signature(C_Defarray.Id(),C_void.Id()),1,MakeFunction1(E_self_print_Defarray_Language,"self_print_Defarray_Language")),MakeString("define.cl:97"))
  
  C_Defrule = NewClass("Defrule",C_Defclaire,C_claire)
  Core.F_close_slot(C_Defrule.AddSlot(C_iClaire_ident,ToType(C_symbol.Id()),CNULL))
  Core.F_close_slot(C_Defrule.AddSlot(C_args,ToType(C_list.Id()),ToType(C_any.Id()).EmptyList().Id()))
  Core.F_close_slot(C_Defrule.AddSlot(C_arg,ToType(C_any.Id()),CNULL))
  Core.F_close_slot(C_Defrule.AddSlot(C_body,ToType(C_any.Id()),CNULL))
  
  _ = Core.F_attach_method(C_self_print.AddMethod(Signature(C_Defrule.Id(),C_void.Id()),1,MakeFunction1(E_self_print_Defrule_Language,"self_print_Defrule_Language")),MakeString("define.cl:105"))
  
  C_Defvar = NewClass("Defvar",C_Defclaire,C_claire)
  Core.F_close_slot(C_Defvar.AddSlot(C_iClaire_ident,ToType(C_Variable.Id()),CNULL))
  Core.F_close_slot(C_Defvar.AddSlot(C_arg,ToType(C_any.Id()),CNULL))
  
  _ = Core.F_attach_method(C_self_print.AddMethod(Signature(C_Defvar.Id(),C_void.Id()),1,MakeFunction1(E_self_print_Defvar_Language,"self_print_Defvar_Language")),MakeString("define.cl:110"))
  
  _ = Core.F_attach_method(Core.C_self_eval.AddEvalMethod(Signature(C_Definition.Id(),C_any.Id()),1,MakeFunction1(E_self_eval_Definition,"self_eval_Definition"),EVAL_Definition),MakeString("define.cl:121"))
  
  _ = Core.F_attach_method(C_Language_fast_definition_ask.AddMethod(Signature(C_class.Id(),C_boolean.Id()),0,MakeFunction1(E_Language_fast_definition_ask_class,"Language_fast_definition_ask_class")),MakeString("define.cl:127"))
  
  _ = Core.F_attach_method(Core.C_self_eval.AddEvalMethod(Signature(C_Language_DefFast.Id(),C_any.Id()),1,MakeFunction1(E_self_eval_DefFast,"self_eval_DefFast"),EVAL_Language_DefFast),MakeString("define.cl:134"))
  
  _ = Core.F_attach_method(C_Language_new_writes.AddMethod(Signature(C_object.Id(),C_list.Id(),C_list.Id()),1,MakeFunction2(E_Language_new_writes_object,"Language_new_writes_object")),MakeString("define.cl:151"))
  
  _ = Core.F_attach_method(Core.C_self_eval.AddEvalMethod(Signature(C_Defobj.Id(),C_any.Id()),1,MakeFunction1(E_self_eval_Defobj,"self_eval_Defobj"),EVAL_Defobj),MakeString("define.cl:167"))
  
  _ = Core.F_attach_method(Core.C_self_eval.AddEvalMethod(Signature(C_Defclass.Id(),C_any.Id()),1,MakeFunction1(E_self_eval_Defclass,"self_eval_Defclass"),EVAL_Defclass),MakeString("define.cl:197"))
  
  _ = Core.F_attach_method(C_Language_getDefault.AddMethod(Signature(C_type.Id(),C_any.Id(),C_any.Id()),0,MakeFunction2(E_Language_getDefault_type,"Language_getDefault_type")),MakeString("define.cl:207"))
  
  { 
    var expr EID
    { var _CL_obj *Core.GlobalVariable
      C_LDEF = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("LDEF",C_claire)))
      
      _CL_obj = C_LDEF
      _CL_obj.Range = ToType(C_any.Id())
      _CL_obj.Value = ToType(CEMPTY.Id()).EmptyList().Id()
      expr = Core.F_close_global_variable(_CL_obj)
      } 
    ErrorCheck(expr)} 
  
  _ = Core.F_attach_method(Core.C_self_eval.AddEvalMethod(Signature(C_Defmethod.Id(),C_any.Id()),1,MakeFunction1(E_self_eval_Defmethod,"self_eval_Defmethod"),EVAL_Defmethod),MakeString("define.cl:238"))
  
  C__Z.Open = -1
  C__sup_equal.Open = -1
  C__equal.Open = -1
  
  _ = Core.F_attach_method(C_Language_attach_comment.AddMethod(Signature(C_any.Id(),C_void.Id()),1,MakeFunction1(E_attach_comment_any,"attach_comment_any")),MakeString("define.cl:247"))
  
  _ = Core.F_attach_method(C_iClaire_extract_signature.AddMethod(Signature(C_list.Id(),C_list.Id()),1,MakeFunction1(E_extract_signature_list,"extract_signature_list")),MakeString("define.cl:260"))
  
  _ = Core.F_attach_method(C_iClaire_extract_pattern.AddMethod(Signature(C_any.Id(),C_list.Id(),C_any.Id()),1,MakeFunction2(E_extract_pattern_any,"extract_pattern_any")),MakeString("define.cl:306"))
  
  _ = Core.F_attach_method(C_iClaire_extract_type.AddMethod(Signature(C_any.Id(),C_type_expression.Id()),1,MakeFunction1(E_extract_type_any,"extract_type_any")),MakeString("define.cl:313"))
  
  _ = Core.F_attach_method(C_extract_item.AddMethod(Signature(C_any.Id(),C_any.Id(),C_any.Id()),0,MakeFunction2(E_extract_item_any,"extract_item_any")),MakeString("define.cl:319"))
  
  _ = Core.F_attach_method(C_Language_extract_pattern_nth.AddMethod(Signature(C_list.Id(),C_list.Id(),C_any.Id()),1,MakeFunction2(E_extract_pattern_nth_list,"extract_pattern_nth_list")),MakeString("define.cl:350"))
  
  _ = Core.F_attach_method(C_iClaire_extract_class_call.AddMethod(Signature(C_class.Id(),C_list.Id(),C_object.Id()),1,MakeFunction2(E_extract_class_call_class,"extract_class_call_class")),MakeString("define.cl:384"))
  
  _ = Core.F_attach_method(C_iClaire_extract_range.AddMethod(Signature(C_any.Id(),
    C_list.Id(),
    C_list.Id(),
    C_list.Id()),1,MakeFunction3(E_extract_range_any,"extract_range_any")),MakeString("define.cl:413"))
  
  C_bit_vector = MakeProperty("bit_vector",2,C_claire)
  
  
  _ = Core.F_attach_method(C_bit_vector.AddMethod(Signature(C_listargs.Id(),C_integer.Id()),1,MakeFunction1(E_bit_vector_listargs2,"bit_vector_listargs2")),MakeString("define.cl:418"))
  
  _ = Core.F_attach_method(C_iClaire_extract_status.AddMethod(Signature(C_any.Id(),C_list.Id()),1,MakeFunction1(E_extract_status_any,"extract_status_any")),MakeString("define.cl:438"))
  
  _ = Core.F_attach_method(C_imported_function.AddMethod(Signature(C_string.Id(),C_function.Id()),0,MakeFunction1(E_imported_function_string,"imported_function_string")),MakeString("define.cl:442"))
  
  _ = Core.F_attach_method(C_iClaire_type_I.AddMethod(Signature(C_any.Id(),C_type.Id()),0,MakeFunction1(E_type_I_any,"type_I_any")),MakeString("define.cl:452"))
  
  _ = Core.F_attach_method(Core.C_self_eval.AddEvalMethod(Signature(C_Defarray.Id(),C_any.Id()),1,MakeFunction1(E_self_eval_Defarray,"self_eval_Defarray"),EVAL_Defarray),MakeString("define.cl:509"))
  
  C_Language_demon = NewClass("demon",C_lambda,It)
  Core.F_close_slot(C_Language_demon.AddSlot(C_mClaire_pname,ToType(C_symbol.Id()),Core.F_symbol_I_string2(MakeString("unamed")).Id()))
  Core.F_close_slot(C_Language_demon.AddSlot(C_Language_priority,ToType(C_integer.Id()),MakeInteger(0).Id()))
  Core.F_close_slot(C_Language_demon.AddSlot(C_formula,ToType(C_lambda.Id()),CNULL))
  
  _ = Core.F_attach_method(C_self_print.AddMethod(Signature(C_Language_demon.Id(),C_void.Id()),1,MakeFunction1(E_self_print_demon,"self_print_demon")),MakeString("define.cl:519"))
  
  _ = Core.F_attach_method(C_funcall.AddMethod(Signature(C_Language_demon.Id(),
    C_any.Id(),
    C_any.Id(),
    C_any.Id()),1,MakeFunction3(E_funcall_demon1,"funcall_demon1")),MakeString("define.cl:520"))
  
  _ = Core.F_attach_method(C_funcall.AddMethod(Signature(C_Language_demon.Id(),
    C_any.Id(),
    C_any.Id(),
    C_any.Id(),
    C_any.Id()),1,MakeFunction4(E_funcall_demon2,"funcall_demon2")),MakeString("define.cl:521"))
  
  C_demons = ToTable(new(ClaireTable).IsNamed(C_table,MakeSymbol("demons",C_claire)))
  C_demons.Range = Core.F_param_I_class(C_list,ToType(C_Language_demon.Id()))
  C_demons.Params = C_any.Id()
  C_demons.Domain = ToType(C_relation.Id())
  C_demons.GraphInit()
  
  C_demons.Default = ToType(C_Language_demon.Id()).EmptyList().Id()
  
  C__inf_dash = ToOperation(new(ClaireOperation).IsNamed(C_operation,MakeSymbol("<-",C_claire)))
  
  
  C_Language_rule_object = NewClass("rule_object",C_property,It)
  
  C_Language_relations = ToTable(new(ClaireTable).IsNamed(C_table,MakeSymbol("relations",It)))
  C_Language_relations.Multivalued_ask = CTRUE
  C_Language_relations.Range = ToType(C_set.Id())
  C_Language_relations.Params = C_any.Id()
  C_Language_relations.Domain = ToType(C_Language_rule_object.Id())
  C_Language_relations.GraphInit()
  
  C_Language_relations.Default = CEMPTY.Id()
  
  C_Language_last_rule = ToTable(new(ClaireTable).IsNamed(C_table,MakeSymbol("last_rule",It)))
  C_Language_last_rule.Range = ToType(C_Language_rule_object.Id())
  C_Language_last_rule.Params = C_any.Id()
  C_Language_last_rule.Domain = ToType(C_relation.Id())
  C_Language_last_rule.GraphInit()
  
  C_Language_last_rule.Default = CNULL
  
  { 
    var expr EID
    expr = Core.F_update_property(C_inverse,
      ToObject(C_Language_relations.Id()),
      8,
      C_relation,
      C_Language_last_rule.Id())
    ErrorCheck(expr)} 
  
  C_eval_rule = MakeProperty("eval_rule",3,C_claire)
  C_eval_rule.Open = 3
  
  
  _ = Core.F_attach_method(Core.C_self_eval.AddEvalMethod(Signature(C_Defrule.Id(),C_any.Id()),1,MakeFunction1(E_self_eval_Defrule,"self_eval_Defrule"),EVAL_Defrule),MakeString("define.cl:555"))
  
  _ = Core.F_attach_method(C_Language_eventMethod_ask.AddMethod(Signature(C_relation.Id(),C_boolean.Id()),0,MakeFunction1(E_eventMethod_ask_relation2,"eventMethod_ask_relation2")),MakeString("define.cl:559"))
  
  _ = Core.F_attach_method(C_Language_make_filter.AddMethod(Signature(C_any.Id(),MakeTuple(C_relation.Id(),Core.F_nth_class1(C_list,ToType(C_Variable.Id())).Id()).Id()),1,MakeFunction1(E_make_filter_any,"make_filter_any")),MakeString("define.cl:592"))
  
  _ = Core.F_attach_method(C_Language_make_demon.AddMethod(Signature(C_relation.Id(),
    C_symbol.Id(),
    Core.F_nth_class1(C_list,ToType(C_Variable.Id())).Id(),
    C_any.Id(),
    C_any.Id(),
    C_Language_demon.Id()),1,MakeFunction5(E_make_demon_relation,"make_demon_relation")),MakeString("define.cl:615"))
  
  _ = Core.F_attach_method(C_Language_readCall.AddMethod(Signature(C_relation.Id(),C_any.Id(),C_Call.Id()),0,MakeFunction2(E_readCall_relation,"readCall_relation")),MakeString("define.cl:620"))
  
  _ = Core.F_attach_method(C_Language_putCall.AddMethod(Signature(C_relation.Id(),
    C_any.Id(),
    C_any.Id(),
    C_Call.Id()),0,MakeFunction3(E_putCall_relation2,"putCall_relation2")),MakeString("define.cl:625"))
  
  _ = Core.F_attach_method(C_Language_safeRange.AddMethod(Signature(C_relation.Id(),C_type.Id()),0,MakeFunction1(E_safeRange_relation,"safeRange_relation")),MakeString("define.cl:633"))
  
  _ = Core.F_attach_method(C_Language_eval_if_write.AddMethod(Signature(C_relation.Id(),C_void.Id()),1,MakeFunction1(E_eval_if_write_relation,"eval_if_write_relation")),MakeString("define.cl:661"))
  
  _ = Core.F_attach_method(C_Language_eventMethod.AddMethod(Signature(C_property.Id(),C_void.Id()),0,MakeFunction1(E_eventMethod_property,"eventMethod_property")),MakeString("define.cl:670"))
  
  _ = Core.F_attach_method(C_Language_jito.AddMethod(Signature(C_any.Id(),C_any.Id()),1,MakeFunction1(E_Language_jito_any,"Language_jito_any")),MakeString("define.cl:710"))
  
  _ = Core.F_attach_method(C_Language_letJito.AddMethod(Signature(C_Let.Id(),C_any.Id()),1,MakeFunction1(E_Language_letJito_Let,"Language_letJito_Let")),MakeString("define.cl:726"))
  
  _ = Core.F_attach_method(C_Language_makeJito.AddMethod(Signature(C_Call.Id(),C_void.Id()),1,MakeFunction1(E_Language_makeJito_Call,"Language_makeJito_Call")),MakeString("define.cl:752"))
  
  _ = Core.F_attach_method(C_Language_makeCallMatch.AddMethod(Signature(C_restriction.Id(),C_list.Id(),C_boolean.Id()),0,MakeFunction2(E_Language_makeCallMatch_restriction,"Language_makeCallMatch_restriction")),MakeString("define.cl:760"))
  
  C_table.Open = ClEnv.Final
  C_class.Open = ClEnv.Final
  C_method.Open = ClEnv.Final
  C_slot.Open = ClEnv.Final
  C_boolean.Open = -1
  { 
    var x *ClaireClass
    _ = x
    var x_iter *ClaireAny
    var x_support *ClaireSet
    x_support = C_Instruction.Descendants
    for i_it := 0; i_it < x_support.Count; i_it++ { 
      x_iter = x_support.At(i_it)
      x = ToClass(x_iter)
      x.Open = ClEnv.Default
      } 
    } 
  
  } 

