// ==================================================================
// golang experiments Phase 2 - Copyright (C) Yves Caseau, 2020-22
// started on June 21st, 2020
// clEnv.go
// ==================================================================

// this is the ClEnv.cpp file (system, exception, worlds) + Dict + Int / Float

package Kernel

import (
	"fmt"
//	"bytes"
	"math"
	"os"
	"runtime"
	"os/exec"
	"strconv"
	"time"
	"unicode/utf8"
)

// +---------------------------------------------------------------------------+
// |  Table of contents                                                        |
// |  Part 1: System object (clEnv )                                           |
// |  Part 2: Exceptions                                                       |
// |  Part 3: Worlds  (clRes)                                                  |
// |  Part 4: Dictionary & Tables                                              |
// |  Part 5: Reader functions                                                 |
// +---------------------------------------------------------------------------+

// +---------------------------------------------------------------------------+
// |  Part 1: System object (clEnv)                                            |
// +---------------------------------------------------------------------------+

/* (isa @ object, verbose @ environment, exception! @ environment, module! @ environment,
name @ environment, version @ environment, ctrace @ environment, cout @ environment,
mClaire/index @ environment, mClaire/base @ environment, mClaire/debug! @ environment,
trace! @ environment, step! @ environment, Kernel/last_debug @ environment,
Kernel/last_index @ environment, spy! @ environment, mClaire/count_call @ environment,
mClaire/count_level @ environment, mClaire/count_trigger @ environment,
params @ environment, close @ environment, abstract @ environment,
final @ environment, default @ environment, open @ environment, ephemeral @ environment) */

// size factor (set with -s option)
var CLAIRESIZE = 1

// as for the C++ code, we put all global variables into one struct
// note that this struct may be seen as a system object
type ClaireEnvironment struct {
	// --- this part is visible from CLAIRE through a meta-description ---------------
	ClaireSystemObject // this allows : system = &ClEnv
	Verbose            int
	Exception_I        *ClaireException // last
	Module_I           *ClaireModule    // current module
	Name               *ClaireString    // name of the environment ()
	Osname             *ClaireString    // the name of the env "macos","win","linux",...
	Version            float64     // version number ! (4.0 implicit)
	Ctrace             *ClairePort //  current trace port
	Cout               *ClairePort //  current out port
	Cin                *ClairePort //  current in port
	// -------- slots for stack management
	Index int // top of stack
	Base  int // current marker
	// -------- debug & perf slots ----------------
	Debug_I      int           // debug!
	Trace_I      int           // for tracing
	LastDebug    int           // last value of the debug index
	LastIndex    int           // last value of the top of eval stack
	Spy_I        *ClaireObject // store the spy method if any
	CountCall    int           // count the numbers of call
	CountLevel   int           // level at which something happens ...
	CountTrigger *ClaireAny    // what should happen
	// system slots
	Params 		*ClaireList // parameters passed from the shell
	// global constants: open status for classes,   properties and modules
	Close     int // c: do not touch,   p:read only (-1)
	ABSTRACT  int // c: no more instances           (0)
	Final int // c: no more subclasses          (1)
	Default   int //                                (2)
	Open      int // p: open property (extensible)  (3)
	Ephemeral int // c: no more subclasses          (1)
	Jito_ask  *ClaireBoolean     // allows Just In Time Optimization
	NLine   int                // numbers of lines read
	
	// --- stack management  -------------------------------------------------------
	maxStack   int   // size of stack
	EvalStack  []EID // value stack
	indexStack []int // index markers
	step       int   // top of the indexStack

	// --- this part is only accessible through the Kernel -------------------------
	abortOnError int               // 0: OK, 1: stop if error (useful for debug)
	moduleStack  *ClaireListObject // stack for begin/end (NEW)
	tIndex       int               // time counter index /* time counters                      */
	tStack       [10]int           // stack of time counters
	maxBuffer    int               // max size of char buffer
	buffer       []byte            // a local buffer for string creation
	bLength      int               // its length
	gensym       int               // a seed for symbol generation

}

// one object
var ClEnv *ClaireEnvironment = nil

// open is a key slot - here we define the values
// -2   forward       trick used in define.cl to note the forward definition
// -1   closed        c: no more instance, no subclass (but with instances)
// 0   abstract       c: no instances
// 1   final          c: thing with no subclasses        p: compiled
// 2   default        c: default = ephemeral             p: default
// 3    open          c: keep instances                  p: extensible    

// initialisation : ClEnv & ClRes
// TODO : pass as parameters the MaxStack, MaxBuffer, MaxHist values (not constants !)
func InitClEnv() {
	ClEnv = new(ClaireEnvironment)
	ClEnv.Close = -1
	ClEnv.ABSTRACT = 0  // c: no more instances           (0)
	ClEnv.Final = 1     // c: no more subclasses          (1)
	ClEnv.Default = 2   //                                (2)
	ClEnv.Open = 3      	// 				p: open property (extensible)  
	ClEnv.Ephemeral = 4     // c: no more subclasses          (1)
	ClEnv.Jito_ask = CTRUE
	
	ClEnv.maxBuffer = 10000 * CLAIRESIZE
	ClEnv.Module_I = C_claire
	ClEnv.buffer = make([]byte, ClEnv.maxBuffer) // a local buffer for string creation
	ClEnv.bLength = 0
	ClEnv.gensym = 0
	// stack initialisation
	ClEnv.maxStack = 100000 * CLAIRESIZE
	ClEnv.EvalStack = make([]EID, ClEnv.maxStack)
	ClEnv.indexStack = make([]int, ClEnv.maxStack)
	ClEnv.Base = 0
	ClEnv.Index = 0
	ClEnv.tIndex = 0
	ClEnv.Trace_I = 0
	ClEnv.Debug_I = -1         // -1 means no debug
	ClEnv.CountCall = 0        // count the numbers of call
	
	// two slots with CNULL values are set up later (in Bootcore)
	// ClRes is our world set of stacks - see part 3
	ClRes = new(ClaireResource)
	ClRes.maxHist = 1000000 * CLAIRESIZE
	ClRes.init()
	
}

// there is a two step process because the firt step is called very early
func finishClEnv() {
	ClEnv.Isa = C_environment
	C_claire.register(makeSymbol("system"), ClEnv.Id()) // link symbol to object
	ClEnv.Name = MakeString("Kernel")
	ClEnv.Osname = MakeString("macos")
	ClEnv.Exception_I = ToException(CNULL)
	ClEnv.Spy_I = ToObject(CNULL)
	ClEnv.moduleStack = makeEmptyObjectList(0)
	ClEnv.Jito_ask = CTRUE
	ClEnv.CountTrigger = CNULL 				// moved from ClEnvInit() because now CNULL exists
	ClEnv.Params = ToType(C_string.Id()).EmptyList()
	// get the args from the command line
	for _,s := range(os.Args[1:]) {ClEnv.Params.AddFast(MakeString(s).Id())}	
}

// parse the args to see if it starts with -s n
func MemoryFlags() {
 if len(os.Args) > 2 && os.Args[1] == "-s" {
		 size,err := strconv.Atoi(os.Args[2])
         if err == nil {
			 fmt.Printf("==> change CAIRE size factor to %d (string buffer & stacks) ===\n",size)
             CLAIRESIZE = size }}
}
    



// stack management methods --------------------------------------------------

//  CLAIRE 4 uses a dual stack:
//     evalStack (EID) is the stack onto which we store EID values (args for calls)
//     indexStack (int) is the call structure stack (push and pop a call)

/* step : create a new stack block
func (c *ClaireEnvironment) StepStack(base int) {
	c.indexStack[c.step] = c.Base
	c.Base = base
	// fmt.Printf("============== Add a STEP STACK [%d] ==================\n",sBase)
	c.step = c.step + 1
} */

// push: add to the stack
func (c *ClaireEnvironment) Push(x EID) EID {
	/* debug warning + debug check ---------------------
	if c.Index > c.maxStack - 10  {
		fmt.Printf("Push %s in stack at position %d\n",PEID(x),c.Index)
		for i := -20; i < 0; i++ {
			j := c.Index + i
			fmt.Printf("stack[%d] = %s\n",j,PEID(c.EvalStack[j]))}
		panic("stack overflow, maxStack too small")
	}
	BadI(x,"ClEnv.Push")       // do no push an integer OID
	// -------- end of debug ----------------------------- */
	c.EvalStack[c.Index] = x
	// fmt.Printf("=== pushStack %s at [%d]\n", PEID(x), c.Index)
	c.Index = c.Index + 1
	if c.Index > c.maxStack - 1 {
		panic("stack overflow, maxStack too small or infinite loop")
	}
	return x
}

// moves the stack pointer up by x units */
func F_stack_add(x int) {
	ClEnv.Index = ClEnv.Base + x
	if ClEnv.Index >= ClEnv.maxStack {
		panic("stack overflow, maxStack too small")
	}
}

func E_stack_add(x EID) EID { 
	F_stack_add(INT(x))
	return EVOID
}

// pop: return to previous block in the stack (current + top)
func (c *ClaireEnvironment) PopStack() {
	c.step = c.step - 1
	if c.step < 0 {
		panic("stack management error : Fatal Failure")
	}
	c.Index = c.Base
	c.Base = c.indexStack[c.step] // previous value of curIndex
	// fmt.Printf("============== Remove a STEP STACK [%d] ==================\n", sBase)
}

//  p.stackMethod() => find the method that matches all the args
// this version uses a signature match (class only) -> to extend later (TODO)
func (p *ClaireProperty) StackMethod(narg int) *ClaireAny {
	n := len(p.Restrictions.ValuesO())
	for i := 0; i < n; i++ {
		m := ToMethod(p.Restrictions.ValuesO()[i])
		// fmt.Printf("method m of type %s\n",m.isa.name)
		if m.Isa.IsIn(C_method) == CTRUE && len(m.Srange.ValuesO()) == narg+1 {
			// check that all EID in stack matches
			var matchOK bool = true
			for j := 0; j < narg; j++ {
				matchOK = true
				// fmt.Printf("match arg[%d] with type %s\n",j,ClEnv.EvalStack[ClEnv.Base + j].Isa.Name.Key)
				if OWNER(ClEnv.EvalStack[ClEnv.Base+j]).IsIn(ToClass(m.Srange.ValuesO()[j])) != CTRUE {
					matchOK = false
					break
				}
			}
			if matchOK {
				return m.ToAny()
			}
		}
	}
	return CNULL
}

// ----------- variable optimization methods --------------------------------------------------
// these are used directly by the compiler

// The EVAL go function for: Kernel/Variable
func EVAL_Variable(x *ClaireAny) EID {
	return ClEnv.EvalStack[(ClEnv.Base + To_Variable(x).Index)]
}

// Write a value (EID)
func (self *ClaireVariable) WriteEID(val EID) EID {
	if (self.Range.Id() == CNULL) || (self.Range.CONTAINS(val) == CTRUE) {
		ClEnv.EvalStack[(ClEnv.Base + self.Index)] = val
		return val
	} else {
		return Cerror(300, self.Id(), self.Range.Id())
	}
}

// member methods --------------------------------------------------------------

// output a char to Cout - prints a char on the current port
func (c *ClaireEnvironment) put(r rune) {
	ClEnv.Cout.Putc(r)
}

// func (c *ClaireEnvironment) PutInteger(n int) { }
// func (c *ClaireEnvironment) PutString(s string) { }

// call to system exit
func (c *ClaireEnvironment) Abort() {
	os.Exit(1)
}

func E_abort_system(s EID) EID {
	os.Exit(1)
	return EVOID
}

// returns the appropriate file separator
func (c *ClaireEnvironment) FileSeparator() *ClaireString {
	if c.Osname.Value == "win" {return MakeString("\\")
    } else {return MakeString("/")} 
}

func E_file_separator (c EID) EID {return EID{ClEnv.FileSeparator().Id(),0}}

// restore a good working state.
// used to be called clean_state -> called from the meta-reader more complex form
func F_restore_state_void() {
	ClEnv.Index = 3
	ClEnv.Base = 1
	ClEnv.Cout = makeFile(os.Stdout)
}

func E_restore_state_void(s EID) EID {
	F_restore_state_void()
	return EVOID
}

// buffer methods : note that these methods are local to the Kernelgrep
// Start a buffer
func (c *ClaireEnvironment) bufferStart() { c.bLength = 0 }

// push a char in the better
func (c *ClaireEnvironment) pushChar(r rune) {
	n := c.bLength
	if n >= c.maxBuffer {
		panic("ClEnv.buffer has reached its maximum size")
	} else if r < 128 { // simple case
		c.buffer[n] = (byte)(r)
		c.bLength = n + 1
		// fmt.Printf("PushChar(%c) at position %d -> length=%d\n", r, n, n+1)
	} else {
		size := utf8.EncodeRune(c.buffer[n:n+3], r)
		c.bLength = n + size
	}
}

// prints an integer in the string buffer
func (c *ClaireEnvironment) pushInteger(n int) { c.pushString(strconv.Itoa(n)) }

// prints a string in the string buffer
func (c *ClaireEnvironment) pushString(s string) {
	for _, r := range s {
		c.pushChar(r)
	}
}

// end buffer, returns a safe copy
func (c *ClaireEnvironment) bufferCopy() string {
	n := c.bLength
	s := make([]byte, n)
	copy(s, c.buffer[0:n])
	return string(s)
}

// new in CLAIRE 4 (since playing with '0' is no longer OK)
// returns a pointer to the buffer seen as a string => read Only
// NOTICE THE CRAZY SEMANTIC OF SLICE (0:n picks n chars ... up to n-1)
func (c *ClaireEnvironment) bufferPeek() string {
	return string(c.buffer[0:c.bLength])
}

// -----------------------  time management -------------------------------------------------------

//  we measure the time in ms as a float
func makeTimestamp() int {
	return int(time.Now().UnixNano())
}

// set the time counter : notice that we use panic and not an error
func F_time_set_void() {
	ClEnv.tIndex++
	if ClEnv.tIndex > 9 {
		fmt.Printf("tIndex = %d\n", ClEnv.tIndex)
		panic("time stack overflow, too many embedded timings")
	}
	ClEnv.tStack[ClEnv.tIndex] = makeTimestamp()
}

func E_time_set_void(void EID) EID {
	F_time_set_void()
	return EVOID
}

// shows the elapsed time and pop the counters
func F_time_get_void() int {
	now := makeTimestamp()
	if ClEnv.tIndex == 0 {
		panic("time stack error, unbalanced embedded timings")
	} else {
		ClEnv.tIndex--
	}
	return (now - ClEnv.tStack[ClEnv.tIndex+1])
}

func E_time_get_void(void EID) EID { return EID{C__INT, IVAL(F_time_get_void())} }

// reads the elapsed time in microseconds
func F_time_read_void() int {
	now := makeTimestamp()
	return (now - ClEnv.tStack[ClEnv.tIndex]) / 1000
}

func E_time_read_void(void EID) EID { return EID{C__INT, IVAL(F_time_read_void())} }

// shows the elapsed time in milliseconds
func F_time_show_void() {
	now := makeTimestamp()
	if ClEnv.Index == 0 {
		panic("time stack error, unbalanced embedded timings")
	} else {
		ClEnv.tIndex--
	}
	F_princ_string(MakeString("Counter["))
	F_princ_integer(ClEnv.tIndex)
	F_princ_string(MakeString("] Elapsed time: "))
	F_princ_integer((now - ClEnv.tStack[ClEnv.tIndex+1]) / 1000000)
	F_princ_string(MakeString("ms. \n"))
}

func E_time_show_void(void EID) EID {
	F_time_show_void()
	return EVOID
}

// shell
func F_claire_shell(s *ClaireString) {
	fmt.Printf("execute: %s\n",s.Value)
	var cmd *exec.Cmd
	// var stdout bytes.Buffer        -> left if one day we want the result
	if ClEnv.Osname.Value == "win" {cmd = exec.Command("cmd","/C",s.Value)
	} else { cmd = exec.Command("bash","-c",s.Value)}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	// cmd.Stdout = &stdout
    cmd.Run()
	// fmt.Printf("result : %s\n",stdout.String())
}

func E_claire_shell(s EID) EID {
	F_claire_shell(ToString(OBJ(s)))
	return EVOID
}

// exit : extended with -1 option that triggers a failure (and the use of go debugging)
func F_CL_exit(i int) {
	if i == -1 {panic("stop and see why")}
	os.Exit(i)
}

func E_CL_exit(s EID) EID {
	F_CL_exit(INT(s))        // was os.Exit(INT(s))
	return EVOID
}

// memory management is left to go in CLAIRE4 ! 
// this function prints some information
func F_claire_stat() {
   fmt.Printf("--------------- Claire memory statistics ---------------------------\n")
   fmt.Printf("eval stack size:%d \n",ClEnv.maxStack)
   fmt.Printf("string buffer size:%d \n",ClEnv.maxBuffer)
   m := new(runtime.MemStats)
   runtime.ReadMemStats(m)
   fmt.Printf("total allocated objects:%d \n",m.TotalAlloc)
   fmt.Printf("currently allocated objects:%d \n",m.HeapAlloc)
   fmt.Printf("number of Go GC:%d \n",m.NumGC)
  
}

func E_claire_stat(x EID) EID {
	F_claire_stat()
	return EVOID
}


// +---------------------------------------------------------------------------+
// |  Part 2: Exceptions                                                       |
// +---------------------------------------------------------------------------+

// exceptions are handled through EID: every function that may return an error
// is compiled (and interpreted obviously) as an EID-returning function
// the key fragment will be
// x = f() ; if ErrorIn(x) {return x}

// we hope that golang sees this as a macro :)
func ErrorIn(x EID) bool {
	return x.VAL == 1 && x.PTR != C__INT && x.PTR != C__FLOAT && x.PTR != C__CHAR
}

// create a system error
func Cerror(n int, a *ClaireAny, b *ClaireAny) EID {
	o := new(ClaireSystemError)
	o.Isa = C_system_error
	o.Index = n
	o.Value = a
	o.Arg = b
	if ClEnv.Verbose > 5 || n == 12 {panic("stop and see why")}
   	o.Close()    // very unclear that we need this 
	return EID{o.Id(), 1}
}

// close simply transforms an exception into an EID :)
func (e *ClaireException) Close() EID {
	ClEnv.Exception_I = e
	ClEnv.LastDebug = ClEnv.Debug_I
	ClEnv.LastIndex = ClEnv.Index
	if ClEnv.Verbose > 4 {fmt.Printf("======== >  Close_exception %s \n", e.Prt())}
	if ClEnv.Verbose > 5 { panic("stop and see why")}  // debug line: convenient to see where the error occurred
	return EID{e.Id(), 1}
}

// this is a callable method from the interpreter
func E_close_exception (e EID) EID {return ToException(OBJ(e)).Close() }


// check if an error occurred
func ErrorCheck(x EID) {
	if ErrorIn(x) {
		fmt.Printf("----- an error has occurred in compiled code ------\n")
		e := OBJ(x)
		for _, m := range C_self_print.Restrictions.ValuesO() {
			if e.Isa.IsIn(ToClass(ToMethod(m).Srange.At(0))) == CTRUE {
				toFunction1(ToMethod(m).Functional).call(x)
				break
			}
		}
		panic("execution must stop")
	}
}

// check that a value is not unknown
func (val *ClaireAny) KNOWN (cause *ClaireAny) EID {
	if (val == CNULL) {return Cerror(24,cause,val)
	} else {return val.ToEID()}
}



// +---------------------------------------------------------------------------+
// |  Part 3: Worlds                                                           |
// +---------------------------------------------------------------------------+

// ClRes is the struct (mirror to ClEnv) that contains the world structure
//
type ClaireResource struct {
	ascii          []*ClaireChar        // cashing ascii chars
	maxHist        int                  // size of stacks
	cWorld         int                  // world number
	cWorldId       int                  // unique ID (useful for debug)
	oBase          int                  // start of current world in stack
	oIndex         int                  // top of stack
	slotObjRec     []*ClaireObject      // receiver of the defeasible update
	slotObjIndex   []int                // index
	slotObjVal     []*ClaireAny         // value
	iBase          int                  // start of current world in stack
	iIndex         int                  // top of stack
	slotIntRec     []*ClaireObject      // receiver of the defeasible update
	slotIntIndex   []int                // index
	slotIntVal     []int                // value
	fBase          int                  // start of current world in stack
	fIndex         int                  // top of stack
	slotFloatRec   []*ClaireObject      // receiver of the defeasible update
	slotFloatIndex []int                // index
	slotFloatVal   []float64            // value
	olBase         int                  // start of current world in stack
	olIndex        int                  // top of stack
	listObjRec     []*ClaireListObject  // receiver of the defeasible update
	listObjIndex   []int                // index
	listObjVal     []*ClaireAny         // value
	ilBase         int                  // start of current world in stack
	ilIndex        int                  // top of list_integer stack
	listIntRec     []*ClaireListInteger // receiver of the defeasible update
	listIntIndex   []int                // index
	listIntVal     []int                // value
	flBase         int                  // start of current world in stack
	flIndex        int                  // top of stack
	listFloatRec   []*ClaireListFloat   // receiver of the defeasible update
	listFloatIndex []int                // index
	listFloatVal   []float64            // value
	odBase         int                  // start of current world in stack
	odIndex        int                  // top of stack
	dictObjRec     []*ClaireMapSet         // receiver of the defeasible update for dict
	dictObjIndex   []*ClaireAny         // index is x in d[x]=y
	dictObjVal     []*ClaireAny         // value

}

var ClRes *ClaireResource

// init
func (c *ClaireResource) init() {
	c.cWorld = 0
	c.cWorldId = 0
	c.oBase = 0
	c.oIndex = 0
	c.slotObjRec = make([]*ClaireObject, c.maxHist)
	c.slotObjIndex = make([]int, c.maxHist)
	c.slotObjVal = make([]*ClaireAny, c.maxHist)
	c.iBase = 0
	c.iIndex = 0
	c.slotIntRec = make([]*ClaireObject, c.maxHist)
	c.slotIntIndex = make([]int, c.maxHist)
	c.slotIntVal = make([]int, c.maxHist)
	c.fBase = 0
	c.fIndex = 0
	c.slotFloatRec = make([]*ClaireObject, c.maxHist)
	c.slotFloatIndex = make([]int, c.maxHist)
	c.slotFloatVal = make([]float64, c.maxHist)
	c.olBase = 0
	c.olIndex = 0
	c.listObjRec = make([]*ClaireListObject, c.maxHist)
	c.listObjIndex = make([]int, c.maxHist)
	c.listObjVal = make([]*ClaireAny, c.maxHist)
	c.ilBase = 0
	c.ilIndex = 0
	c.listIntRec = make([]*ClaireListInteger, c.maxHist)
	c.listIntIndex = make([]int, c.maxHist)
	c.listIntVal = make([]int, c.maxHist)
	c.flBase = 0
	c.flIndex = 0
	c.listFloatRec = make([]*ClaireListFloat, c.maxHist)
	c.listFloatIndex = make([]int, c.maxHist)
	c.listFloatVal = make([]float64, c.maxHist)
	c.odBase = 0
	c.odIndex = 0
	c.dictObjRec = make([]*ClaireMapSet, c.maxHist)
	c.dictObjIndex = make([]*ClaireAny, c.maxHist)
	c.dictObjVal = make([]*ClaireAny, c.maxHist)

}

// defeasible update on a list  (not a method because there are many variations)
// WARNING: this version does not look for size ...
func F_store_list(l *ClaireList,n int, y *ClaireAny, b *ClaireBoolean) *ClaireAny {
	// fmt.Printf("store on list %s at n=%d with y=%s\n",l.Prt(),n,y.Prt())
	if l.Srange == C_integer {
		l.toInteger().store(n, ToInteger(y).Value, b)
	} else if l.Srange == C_float {
		l.toFloat().store(n, ToFloat(y).Value, b)
	} else {
		l.toObject().store(n, y, b)
	}
	return y
}

// this is optimized to avoid ANY(y) allocation
func E_store_list(leid EID, n EID, y EID, b EID) EID {
	l := ToList(OBJ(leid))
    // fmt.Printf("EID store on list %s at n=%d with y=%s\n",l.Prt(),INT(n),PEID(y))
	if l.Srange == C_integer {
		l.toInteger().store(INT(n), INT(y), ToBoolean(OBJ(b)))
	} else if l.Srange == C_float {
		l.toFloat().store(INT(n), FLOAT(y), ToBoolean(OBJ(b)))
	} else {
		l.toObject().store(INT(n), OBJ(y), ToBoolean(OBJ(b)))
	}
	return y
}

// three versions that are sort dependent
func (l *ClaireListObject) store(n int, y *ClaireAny, b *ClaireBoolean)  {
	if b == CTRUE {
		ClRes.olIndex++
		if ClRes.olIndex > ClRes.maxHist {
			panic("History stack overflow")
		}
		ClRes.listObjRec[ClRes.olIndex] = l
		ClRes.listObjIndex[ClRes.olIndex] = n-1
		ClRes.listObjVal[ClRes.olIndex] = l.Values[n-1]
	}
	l.Values[n-1] = y
}

func (l *ClaireListInteger) store(n int, y int, b *ClaireBoolean) {
	if b == CTRUE {
		ClRes.ilIndex++
		if ClRes.ilIndex > ClRes.maxHist {
			panic("History stack overflow")
		}
		// fmt.Printf("store list Int [%d] %p at %d = %p\n",ClRes.ilIndex,l,n-1,l.Values[n-1])
		ClRes.listIntRec[ClRes.ilIndex] = l
		ClRes.listIntIndex[ClRes.ilIndex] = n-1
		ClRes.listIntVal[ClRes.ilIndex] = l.Values[n-1]
	}
	l.Values[n-1] = y
}

func (l *ClaireListFloat) store(n int, y float64, b *ClaireBoolean) {
	if b == CTRUE {
		ClRes.flIndex = ClRes.flIndex + 1
		if ClRes.flIndex > ClRes.maxHist {
			panic("History stack overflow")
		}
		ClRes.listFloatRec[ClRes.flIndex] = l
		ClRes.listFloatIndex[ClRes.flIndex] = n-1
		ClRes.listFloatVal[ClRes.flIndex] = l.Values[n-1]
	}
	l.Values[n-1] = y
}

// there are three variants, depending on the sort
// notice that we use the proper update methods (to prevent GC problems) versus pointer handling
// this is the generic version (equivalent of store_object in C++)
func F_store_object (x *ClaireObject, n int, s *ClaireClass, y *ClaireAny, b *ClaireBoolean) *ClaireAny {
	if (ClEnv.Verbose > 10) {fmt.Printf("compiled store on object %s at n=%d with y=%s\n",x.Prt(),n,y.Prt())}
	if s == C_integer {
		x.StoreInt(n, ToInteger(y).Value, b)
	} else if s == C_float {
		x.StoreFloat(n, ToFloat(y).Value, b)
	} else {
		x.StoreObj(n, y, b)
	}
	return y
}

// this is optimized to avoid ANY(y)
func E_store_object(xeid EID, n EID, seid EID, y EID, b EID) EID {
	s := ToClass(OBJ(seid))
	x := ToObject(OBJ(xeid))
	// fmt.Printf("EID store on object %s at n=%d with y=%s\n",x.Prt(),INT(n),PEID(y))
	if s == C_integer {
		x.StoreInt(INT(n), INT(y), ToBoolean(OBJ(b)))
	} else if s == C_float {
		x.StoreFloat(INT(n), FLOAT(y), ToBoolean(OBJ(b)))
	} else {
		x.StoreObj(INT(n), OBJ(y), ToBoolean(OBJ(b)))
	}
	return y
}

// same pattern: three specialized versions
func (x *ClaireObject) StoreObj(n int, y *ClaireAny, b *ClaireBoolean) {
	if b == CTRUE {
		ClRes.oIndex = ClRes.oIndex + 1
		if ClRes.oIndex >= ClRes.maxHist {
			panic("History stack overflow")
		}
		// fmt.Printf("[%d] store slot%d(%s) = %s\n",ClRes.oIndex,n,x.Prt(),x.GetObj(n).Prt())
		ClRes.slotObjRec[ClRes.oIndex] = x
		ClRes.slotObjIndex[ClRes.oIndex] = n
		ClRes.slotObjVal[ClRes.oIndex] = x.GetObj(n)
	}
	x.SetObj(n, y)
}

func (x *ClaireObject) StoreInt(n int, y int, b *ClaireBoolean) {
	if b == CTRUE {
		ClRes.iIndex = ClRes.iIndex + 1
		if ClRes.iIndex >= ClRes.maxHist {
			panic("History stack overflow")
		}
		ClRes.slotIntRec[ClRes.iIndex] = x
		ClRes.slotIntIndex[ClRes.iIndex] = n
		ClRes.slotIntVal[ClRes.iIndex] = x.GetInt(n)
	}
	x.SetInt(n, y)
}

func (x *ClaireObject) StoreFloat(n int, y float64, b *ClaireBoolean) {
	if b == CTRUE {
		ClRes.fIndex = ClRes.fIndex + 1
		if ClRes.fIndex >= ClRes.maxHist {
			panic("History stack overflow")
		}
		ClRes.slotFloatRec[ClRes.fIndex] = x
		ClRes.slotFloatIndex[ClRes.fIndex] = n
		ClRes.slotFloatVal[ClRes.fIndex] = x.GetFloat(n)
	}
	x.SetFloat(n, y)
}

/* performs an addition to a list and store the relevant changes
// THIS IS NOT IMPLEMENTED IN CLAIRE4/Go -> would require another kind of stack
func F_store_add(l *ClaireList,
{if (l->length + 1 == (*l)[0]) Cerror(43, _oid_(l), l->length);
 if (ClRes->cWorld) STOREOID(&(l->length),l->length)
 return l->addFast(y); } */

// add one new world : the top of each stacks become the new base,
func F_world_push() {
	// fmt.Printf("------- world push @ %d -------------------------\n",ClRes.cWorld)
	c := ClRes
	if c.iIndex >= c.maxHist || c.oIndex >= c.maxHist || c.fIndex >= c.maxHist ||
		c.olIndex >= c.maxHist || c.ilIndex >= c.maxHist || c.flIndex >= c.maxHist {
		panic("History stack overflow")
	}
	c.cWorld++
	c.cWorldId++
	// add a new layer onto each stack, using Index stack for step back (push previous base)
	c.oIndex++
	c.iIndex++
	c.fIndex++
	c.slotObjIndex[c.oIndex] = c.oBase // we use the index stack to chain bases
	c.slotIntIndex[c.iIndex] = c.iBase //
	c.slotFloatIndex[c.fIndex] = c.fBase
	if ClEnv.Verbose > 10 {fmt.Printf("obj slot stack base : %d -> %d = new index\n",c.oBase,c.oIndex)}
	c.oBase = c.oIndex
	c.iBase = c.iIndex
	c.fBase = c.fIndex
	// dictionary indexes
	c.odIndex++                                            // index +1
	c.dictObjIndex[c.odIndex] = MakeInteger(c.odBase).Id() // store the base value in the stack
	c.odBase = c.odIndex
	// same thing for the list stacks
	c.olIndex++
	c.ilIndex++
	c.flIndex++
	c.listObjIndex[c.olIndex] = c.olBase
	c.listIntIndex[c.ilIndex] = c.ilBase
	c.listFloatIndex[c.flIndex] = c.flBase
	c.olBase = c.olIndex
	c.ilBase = c.ilIndex
	c.flBase = c.flIndex
}

func E_world_push(c EID) EID {
	F_world_push()
	return EVOID
}

// remove a world and perform all modifications stored in the stack
func F_world_pop() {
	// fmt.Printf("------- world pop @ %d -------------------------\n",ClRes.cWorld)
	c := ClRes
	c.cWorldId++ // v3.2.04
	c.cWorld--
	if c.cWorld < 0 {
		c.cWorld++ // cannot go further
	} else {
		// objects slots : base N becomes previous base(w-1) & index become N-1
		j := c.oBase
		for i := c.oIndex; i > j; i-- {
			c.slotObjRec[i].SetObj(c.slotObjIndex[i], c.slotObjVal[i])
		}
		if ClEnv.Verbose > 10 {fmt.Printf("obj slot: index de %d a %d, new base is %d\n",c.oIndex,j + 1,c.slotObjIndex[j])}
		c.oIndex = j - 1
		c.oBase = c.slotObjIndex[j]           // use base chaining
		// integer slots
		j = c.iBase
		for i := c.iIndex; i > j; i-- {
			c.slotIntRec[i].SetInt(c.slotIntIndex[i], c.slotIntVal[i])
		}
		// fmt.Printf("int slot: index de %d a %d, new base is %d\n",c.iIndex,j + 1,c.slotIntIndex[j])
		c.iIndex = j - 1
		c.iBase = c.slotIntIndex[j]
		j = c.fBase
		for i := c.fIndex; i > j; i-- {
			c.slotFloatRec[i].SetFloat(c.slotFloatIndex[i], c.slotFloatVal[i])
		}
		c.fIndex = j - 1
		c.fBase = c.slotFloatIndex[j]
		// dict updates
		j = c.odBase
		for i := c.odIndex; i > j; i-- {
			c.dictObjRec[i].Value[c.dictObjIndex[i].Key()] = c.dictObjVal[i]
		}
		c.odIndex = j - 1
		c.odBase = ToInteger(c.dictObjIndex[j]).Value
		// list updates
		j = c.olBase
		for i := c.olIndex; i > j; i-- {
			c.listObjRec[i].Values[c.listObjIndex[i]] = c.listObjVal[i]
		}
		c.olIndex = j - 1
		c.olBase = c.listObjIndex[j]
		j = c.ilBase
		for i := c.ilIndex; i > j; i-- {
			c.listIntRec[i].Values[c.listIntIndex[i]] = c.listIntVal[i]
		}
		c.ilIndex = j - 1
		c.ilBase = c.listIntIndex[j]
		j = c.flBase
		for i := c.flIndex; i > j; i-- {
			c.listFloatRec[i].Values[c.listFloatIndex[i]] = c.listFloatVal[i]
		}
		c.flIndex = j - 1
		c.flBase = c.listFloatIndex[j]
	}
}

func E_world_pop(c EID) EID {
	F_world_pop()
	return EVOID
}

// commit: all updates are accepted and the backtrack values are moves the previous world
// this is a base manipulation : the previous base is the current top, and what was on top of previous base is now below
// Notice that if the world is not empty, the "base line" that is used for based chaining is overwritten with a copy of upper line
// this maintains the structure: stack of value lines on top of the "base lines" (chaining)
func F_world_remove() {
	// fmt.Printf("------- world remove @ %d -------------------------\n",ClRes.cWorld)
	c := ClRes
	c.cWorldId++ // v3.2.04
	c.cWorld--
	if c.cWorld < 0 {
		c.cWorld++ // cannot go further
	} else if c.cWorld == 0 { // complete reset of base & index for 7 blocks
		c.iBase = 0
		c.iIndex = 0
		c.oBase = 0
		c.oIndex = 0
		c.fBase = 0
		c.fIndex = 0
		c.ilBase = 0
		c.ilIndex = 0
		c.odIndex = 0
		c.odBase = 0
		c.olBase = 0
		c.olIndex = 0
		c.flBase = 0
		c.flIndex = 0
	} else {
		// integer slot commit - detailed explanation
		j := c.iBase                                    // index does not change if the world is full
		c.iBase = c.slotIntIndex[j]                     // base is the previous base
		if (c.iIndex > j)   { // the world contains updates
			c.slotIntRec[j] = c.slotIntRec[j + 1]        // duplication of world line to erase the base chaining
			c.slotIntIndex[j] = c.slotIntIndex[j + 1]        // ensures that world-() will not cause an issue
			c.slotIntVal[j] = c.slotIntVal[j + 1]        // although it is a waste of time
		} else {c.iIndex = j - 1}                                               // if 
		// object slot commit
		j = c.oBase
		// k := c.oIndex // debug
		c.oBase = c.slotObjIndex[j] //
		if (c.oIndex > j)   { // the world contains updates
			c.slotObjRec[j] = c.slotObjRec[j + 1]        // duplication of world line to erase the base chaining
			c.slotObjIndex[j] = c.slotObjIndex[j + 1]        // ensures that world-() will not cause an issue
			c.slotObjVal[j] = c.slotObjVal[j + 1]        // although it is a waste of time
		} else {fmt.Printf("empty slice for obj, set index to %d\n",j - 1)
			    c.oIndex = j - 1} 
		// fmt.Printf("remove : iBase %d to %d, iIndex %d to %d\n",j,c.oBase,k,c.oIndex)                                              // if 
		// float slot commit
		j = c.fBase
		c.fBase = c.slotFloatIndex[j] //
		if (c.fIndex > j)   { // the world contains updates
			c.slotFloatRec[j] = c.slotFloatRec[j + 1]        // duplication of world line to erase the base chaining
			c.slotFloatIndex[j] = c.slotFloatIndex[j + 1]        // ensures that world-() will not cause an issue
			c.slotFloatVal[j] = c.slotFloatVal[j + 1]        // although it is a waste of time
		} else {c.fIndex = j - 1}                                               // if 
		// dictionary part
		j = c.odBase
		c.odBase = ToInteger(c.dictObjIndex[j]).Value 
		if (c.odIndex > j)   { // the world contains updates
			c.dictObjRec[j] = c.dictObjRec[j + 1]        // duplication of world line to erase the base chaining
			c.dictObjIndex[j] = c.dictObjIndex[j + 1]        // ensures that world-() will not cause an issue
			c.dictObjVal[j] = c.dictObjVal[j + 1]        // although it is a waste of time
		} else {c.odIndex = j - 1}                                               // if 
		// lists
		j = c.ilBase
		c.ilBase = c.listIntIndex[j] // base is the previous base
		if (c.ilIndex > j)   { // the world contains updates
			c.listIntRec[j] = c.listIntRec[j + 1]        // duplication of world line to erase the base chaining
			c.listIntIndex[j] = c.listIntIndex[j + 1]        // ensures that world-() will not cause an issue
			c.listIntVal[j] = c.listIntVal[j + 1]        // although it is a waste of time
		} else {c.ilIndex = j - 1}                                               // if 
		// object lists
		j = c.olBase
		c.olBase = c.listObjIndex[j] //
		if (c.olIndex > j)   { // the world contains updates
			c.listObjRec[j] = c.listObjRec[j + 1]        // duplication of world line to erase the base chaining
			c.listObjIndex[j] = c.listObjIndex[j + 1]        // ensures that world-() will not cause an issue
			c.listObjVal[j] = c.listObjVal[j + 1]        // although it is a waste of time
		} else {c.olIndex = j - 1}                                               // if 
		// float lists
		j = c.flBase
		c.flBase = c.listFloatIndex[j] 
		if (c.flIndex > j)   { // the world contains updates
			c.listFloatRec[j] = c.listFloatRec[j + 1]        // duplication of world line to erase the base chaining
			c.listFloatIndex[j] = c.listFloatIndex[j + 1]        // ensures that world-() will not cause an issue
			c.listFloatVal[j] = c.listFloatVal[j + 1]        // although it is a waste of time
		} else {c.flIndex = j - 1}                                               // if
	}
}

// how to call world-() from the interpreter
func E_world_remove(c EID) EID {
	F_world_remove()
	return EVOID
}

// Note : world slaughter does not exist in CLAIRE 4 any more

// give the current world
func F_world_number() int      { return ClRes.cWorld }
func E_world_number(c EID) EID { return EID{C__INT, IVAL(ClRes.cWorld)} }

// give the current world
func F_world_get_id() int      { return ClRes.cWorldId }
func E_world_get_id(c EID) EID { return EID{C__INT, IVAL(ClRes.cWorldId)} }

// +---------------------------------------------------------------------------+
// |  Part 4: Dictionary                                                       |
// +---------------------------------------------------------------------------+

// this is our hash function
// 
func (x *ClaireAny) Key() string {
	if x.Isa == C_integer {
		return  strconv.Itoa(ToInteger(x).Value)          // golang primitive function => tried to rewrite with no benefits
	} else if x.Isa == C_float {
		return fmt.Sprintf("#F%f", ToFloat(x).Value)
	} else if x.Isa == C_string {
		return "#S" + ToString(x).Value
	} else if x.Isa == C_tuple {
		return ToList(x).tupleKey()
	} else if x.Isa == C_set {
		return ToSet(x).setKey()
	} else {
		return fmt.Sprintf("#O%p", x)
	}
}

// multiple hash for Tuple (works, but expensive for big tuples)
func (x *ClaireList) tupleKey() string {
	rep := "#T"
	for _,v := range(x.ValuesO()) { rep = rep + "," + v.Key()}
	return rep
}

// multiple hash for Sets (works, but really expensive for big tuples)
// sort the keys with naive insert
func (x *ClaireSet) setKey() string {
	rep := "#S"
	var l []string = make([]string, 0)
	var n int = 0
	for k := 0; k < x.Count; k++ {
	    s := x.At(k).Key()
		var j int = 0
		for j < n && l[j] < s { j++ }   // j is the first position such that s < l[j]
		l = append(l, "0")   // Step 1 : increase capacity
		copy(l[j+1:], l[j:]) // Step 2 : shift
		l[j] = s        	 // Step 3 : put s at position j
		n++
	}
	for i := 0; i < n; i++ { rep = rep + "," + l[i]}
	return rep
}

// this is a special version for EID : avoids allocation
func KEY(x EID) string {
	if x.PTR == C__INT {
		return strconv.Itoa(INT(x))
	} else if x.PTR == C__FLOAT {
		return fmt.Sprintf("#F%f",FLOAT(x))
	} else if x.PTR == C__CHAR {
		return fmt.Sprintf("#C%c",CHAR(x))
	} else if x.PTR.Isa == C_string {
		return "#S" + ToString(x.PTR).Value
	} else if x.PTR.Isa == C_tuple {
		return ToList(x.PTR).tupleKey()
	} else if x.PTR.Isa == C_set {
		return ToSet(x.PTR).setKey()
	} else {		
		return fmt.Sprintf("#O%p",x.PTR)
	}
}


// generic hash function ... to be done later


// API functions ------------------------------------------------------------

// create a dict  (CLAIRE 4 : [type:type])  -> map!(t,t)
func (in *ClaireType) Map_I(out *ClaireType) *ClaireMapSet {
	d := new(ClaireMapSet)
	d.Isa = C_map_set
	d.domain = in
	d.mrange = out
	d.Value = make(map[string]*ClaireAny)
	// fmt.Printf("create a map %x\n", d.Uip())
	return d
}

func E_map_I_type(in EID, out EID) EID {
	return EID{ToType(OBJ(in)).Map_I(ToType(OBJ(out))).Id(), 0}
}

// read from a dictionary - reuse the Key function defined for Sets
func (d *ClaireMapSet) Get(x *ClaireAny) *ClaireAny {
	y := d.Value[x.Key()]
	if y == nil {
		return CNULL
	} else {
		return y
	}
}

func E_get_map(d EID, x EID) EID {
	return EID{ToMapSet(OBJ(d)).Get(ANY(x)), 0}
}

// Writes into a dict - TODO : add a specific Cerror
func (d *ClaireMapSet) Put(x *ClaireAny, y *ClaireAny) EID {
	if d.domain.Contains(x) == CFALSE {
		return Cerror(17, x, d.Id())
	}
	if d.mrange.Contains(y) == CFALSE {
		return Cerror(17, y, d.Id())
	}
	d.Value[x.Key()] = y
	return EVOID
}

func E_put_map(d EID, x EID, y EID) EID { return ToMapSet(OBJ(d)).Put(ANY(x), ANY(y)) }

// Fast version for compiler
func (d *ClaireMapSet) AddFast(x *ClaireAny, y *ClaireAny) {
	// fmt.Printf("add to a map %x\n", d.Uip())
	d.Value[x.Key()] = y
}

// these are temporary functions : get & put on anything => the last one should
func F_dict_get_any(d *ClaireAny, x *ClaireAny) *ClaireAny {
	// fmt.Printf(">>>>>  Dict_get with x=%s \n", x.Prt())
	return ToMapSet(d).Get(x)
}

func E_dict_get_any(d EID, x EID) EID {
	return EID{F_dict_get_any(ANY(d), ANY(x)), 0}
}

// write in a dict
func F_dict_put_any(d *ClaireAny, x *ClaireAny, y *ClaireAny) {
	// fmt.Printf(">>>>>  Dict_put x=%s and y=%s\n", x.Prt(), y.Prt())
	ToMapSet(d).AddFast(x, y)
}

func E_dict_put_any(d EID, x EID, y EID) EID {
	F_dict_put_any(ANY(d), ANY(x), ANY(y))
	return EVOID
}

// copy a map (thanks to stack Overflow)
func (d *ClaireMapSet) Copy() *ClaireMapSet {
	d2 := new(ClaireMapSet)
	d2.Isa = C_map_set
	d2.domain = d.domain
	d2.mrange = d.mrange
	d2.Value = make(map[string]*ClaireAny)
	// Copy from the original map to the target map
	for key, value := range d.Value {
		d2.Value[key] = value
	}
	return d2
}

func E_copy_map(d EID) EID {
	return EID{ToMapSet(OBJ(d)).Copy().Id(), 0}
}

// access to range and domain
func (d *ClaireMapSet) Domain() *ClaireType { return d.domain}
func (d *ClaireMapSet) Range() *ClaireType { return d.mrange}

func E_domain_map_set (d EID) EID { return EID{ToMapSet(OBJ(d)).Domain().Id(),0}}
func E_range_map_set (d EID) EID { return EID{ToMapSet(OBJ(d)).Range().Id(),0}}


// set extension of a map = list of values	
func (d *ClaireMapSet) Set_I() *ClaireSet {
	s := d.mrange.EmptySet()
	for _, value := range d.Value {
		s.AddFast(value)
	}
	return s
}	

func E_set_I_map_set(d EID) EID {return EID{ToMapSet(OBJ(d)).Set_I().Id(),0}}

// note: this implementation does not support projection of a map = set of keys


// table function ----------------------------------------------------------------

// implement default copy (cf index_table in ClReflect.cpp)
func (a *ClaireTable) GraphGet(x *ClaireAny) *ClaireAny {
	d := ToMapSet(a.Graph)
	y := d.Value[x.Key()]
	if y == nil {
		d := a.Default
		if d.Isa == C_list { return ToList(d).Copy().Id()
		} else if d.Isa == C_set { return ToSet(d).Copy().Id()
		} else {return d}
	} else {
		return y
	}
}

func E_graph_get_table(a EID, x EID) EID {
	return EID{ToTable(OBJ(a)).GraphGet(ANY(x)), 0}
}

// write into a table using a defeasible update if needed
func (a *ClaireTable) GraphPut(x *ClaireAny, y *ClaireAny) {
	d := ToMapSet(a.Graph)
	if a.Store_ask == CTRUE {
		ClRes.odIndex++
		if ClRes.odIndex > ClRes.maxHist {
			panic("History stack overflow")
		}
		y2 := d.Value[x.Key()]
		fmt.Printf("[%d] graph put - %s(x = %s) was %p\n",ClRes.odIndex, d.Prt(),x.Prt(),y2)
		ClRes.dictObjRec[ClRes.odIndex] = d
		ClRes.dictObjIndex[ClRes.odIndex] = x
		ClRes.dictObjVal[ClRes.odIndex] = d.Value[x.Key()]
	}
	d.AddFast(x, y)
}

func E_graph_put_table(a EID, x EID, y EID) EID {
	ToTable(OBJ(a)).GraphPut(ANY(x), ANY(y))
	return EVOID
}

// initialization of the graph according to param
func (a *ClaireTable) GraphInit() {
	a.Graph = ToType(C_any.Id()).Map_I(ToType(C_any.Id())).Id()
}

func E_graph_init_table(a EID) EID {
	ToTable(OBJ(a)).GraphInit()
	return EVOID
}

// generic hash : useful feature borrowed from CLAIRE 3.5
// private method
func hash_modulo(x *ClaireAny, n int) int {
	// fmt.Printf("hash %s (ident=%s)\n",x.Prt(),x.Isa.Ident_ask.Prt())
	if (x.Isa == C_integer) { return (ToInteger(x).Value % n)
	} else if (x.Isa == C_float) { return (int(FVAL(ToFloat(x).Value)) % n)
	} else if (x.Isa == C_char) { return (int(ToChar(x).Value) % n)
	} else if (x.Isa == C_string) {
		m := F_length_string(ToString(x))
		var v int = 0
		if (m > 10) {m = 10}
		i := 1
	    for _, r := range ToString(x).Value {
			if i > m {break}
			v += ((int)(r) % n)}
		return (v % n)
    }  else if (x.Isa == C_list) {
		m := ToList(x).Length()
		var v int = 0
		if (m > 10) {m = 10}
		for i := 0; i < m; i++ {v += hash_modulo(ToList(x).At(i),n)}
		return (v % n)    
    // to be continued ! (sets)
	} else if x.Isa.Ident_ask == CTRUE {
		return ((int(x.ui64()) >> 8) % n)                 // shift by 8 to avoid alignment constraints
    }  else {return ((int(x.Isa.ui64()) >> 8) % n)} 
	}

func (l *ClaireList) Hash(x *ClaireAny) int {return hash_modulo(x,l.Length())}
func E_hash_list(l EID, x EID) EID {return EID{C__INT,IVAL(ToList(OBJ(l)).Hash(ANY(x)))}}


// +---------------------------------------------------------------------------+
// |  Part 5: Reader functions                                                 |
// +---------------------------------------------------------------------------+

// these are 4 medium-level reading functions that are defined on ports, on top of which the
// reader is built
// they are based on p.getNext, p.Firstc

// reading a string in a port - assumes that " was read and that the string will end with "
func (p *ClairePort) ReadString() *ClaireString {
	cur := p.firstc
	ClEnv.bufferStart()
	for cur != '"' {
		if cur == CEOF {
			break
		}
		if cur == '\\' {
			p.GetNext()
			cur = p.firstc
			if cur == 't' {
				cur = '\t'
			} else if cur == 'n' {
				cur = '\n'
			}
		}
		ClEnv.pushChar(cur)
		p.GetNext()
		cur = p.firstc
	}
	p.GetNext()
	return MakeString(ClEnv.bufferCopy())
}

func E_read_string_port(p EID) EID { return EID{ToPort(OBJ(p)).ReadString().Id(), 0} }

// test
var ETEST *ClaireAny

// reading an ident, which is either a symbol, a number or a special case
// may return an error
func (p *ClairePort) ReadIdent() EID {
	cur := p.firstc
	// fmt.Printf("======================== READ IDENT (%c) ===================\n", cur)
	p.GetNext()
	if cur == '-' && ('0' <= p.firstc && '9' >= p.firstc) {
		value := p.ReadNumber()
		if value.Isa == C_integer {
			return EID{C__INT, IVAL(-(ToInteger(value).Value))}
		} else {
			return EID{C__FLOAT, FVAL(-(ToFloat(value).Value))}
		}
	} else if cur == '\'' {
		cur = p.firstc
		p.GetNext()
		if '\'' != p.firstc {
			fmt.Printf("Cannot read a char\n")
			return Cerror(35, MakeChar(cur).ToAny(), MakeChar(p.firstc).ToAny())
		} else {
			p.GetNext()
			// fmt.Printf("we have read a char\n")
			// fmt.Printf("we return %s\n", PEID(EID{C__CHAR, CVAL(cur)}))
			return EID{C__CHAR, CVAL(cur)}
		}
	} else {
		// fmt.Printf("we read a thing\n")  // debug
		return p.ReadThing(ClEnv.Module_I, cur, ClEnv.Module_I)
	}
}

func E_read_ident_port(p EID) EID { return ToPort(OBJ(p)).ReadIdent() }

// min and max
const (
	MaxInt64 = float64(1<<63 - 1)
	MinInt64 = float64(-1 << 63)
)

// read a number, either a float or an integer
// changed in v3.0.70 to read long floats
func (p *ClairePort) ReadNumber() *ClaireAny {
	var res float64 = (float64)(p.firstc - '0')
	p.GetNext()
	for p.firstc >= '0' && p.firstc <= '9' {
		res = (res * 10.0) + (float64)(p.firstc-'0')
		p.GetNext()
	}
	if p.firstc == '%' {              // new in v4.0.6 => 12% = 0.12
		p.GetNext()
		return MakeFloat(res / 100.0).ToAny()
	} else if p.firstc != '.' && p.firstc != 'e' {   // regular number read
		if res >= MinInt64 && res <= MaxInt64 {
			return MakeInteger((int)(res)).ToAny() // read an int
		} else {
			return MakeFloat(res).ToAny()          // convert to a float
		} // overflow -> float (v3.0.70)
	} else {  
		possible := res      // read a float (saw a e or a .)
		if p.firstc == '.' { // read the decimal part
			res = 10.0
			p.GetNext()
			for p.firstc >= '0' && p.firstc <= '9' {
				possible = possible + (((float64)(p.firstc - '0')) / res)
				res = res * 10.0
				p.GetNext()
			}
		}
		if p.firstc == 'e' { // read the exponent part
			var signe rune = '+'
			res = 0.0
			p.GetNext()
			if p.firstc == '-' {
				signe = '-'
				p.GetNext()
			}
			if p.firstc == '+' {
				p.GetNext()
			}
			for p.firstc >= '0' && p.firstc <= '9' {
				res = (res * 10.0) + (float64)(p.firstc-'0')
				p.GetNext()
			}
			if signe == '-' {
				possible = possible / math.Pow(10.0, res)
			} else {
				possible = possible * math.Pow(10.0, res)
			}
		}
		// new in v4.0.6 (check on previous versions) : accept % at the end of a number
		if p.firstc == '%' { // read the exponent part
			possible = possible / 100.0
			p.GetNext()
		}
		return MakeFloat(possible).ToAny()
	}
}

func E_read_number_port(p EID) EID { return EID{ToPort(OBJ(p)).ReadNumber(), 0} }

// reading a true identifier (symbol or object), if we see a comment we return it as a string
// app is the module in which the stream is read, cur is the current
// character
// def tells where the module is read, def = NULL means that we read a private name
func (p *ClairePort) ReadThing(app *ClaireModule, cur rune, def *ClaireModule) EID {
	if cur == '"' {
		return EID{app.createSymbol(p.ReadString().Value).Id(), 0}
	} // strange : TODO = explain
	if cur == '/' && p.firstc == '*' { // C-style comments
		for cur != '*' || p.firstc != '/' {
			cur = p.firstc
			if cur == '\n' {ClEnv.NLine = ClEnv.NLine + 1}
			p.GetNext()
		}
		p.GetNext()
		return EID{MakeString("").Id(), 0}        // old-style comments are ignored !  (hence nb_line is wrong)
	}
	ClEnv.bufferStart()
	if cur == '/' && p.firstc == '/' { // C++ comment
		p.GetNext()
		for cur = p.firstc; cur != '\n' && cur != CEOF; {
			ClEnv.pushChar(cur)
			p.GetNext()
			cur = p.firstc
		}
		return EID{MakeString(ClEnv.bufferCopy()).Id(), 0}   // returned to CLAIRE as a string (used in generated code)
	}
	ClEnv.pushChar(cur)
	if cur == ':' && p.firstc == ':' {
		ClEnv.pushChar(cur)
		p.GetNext() // :: trap
	} else if cur == '.' && p.firstc == '.' {
		ClEnv.pushChar(cur)
		p.GetNext() // .. trap
	} else if cur == '<' && p.firstc == '<' {
		ClEnv.pushChar(cur)
		p.GetNext() // << trap
	} else if (cur == '-' || cur == '=' || cur == '>') && p.firstc == '>' {
		ClEnv.pushChar('>')
		p.GetNext()
	} // -> trap for *>
	cur = p.firstc
	for alpha_char(cur) == CFALSE {
		ClEnv.pushChar(cur)
		p.GetNext()
		cur = p.firstc
	}
	if cur == '/' { // read a qualified ident
		mname := MakeString(ClEnv.bufferPeek())
		s := F_get_symbol_module(app,mname) // symbol or CNULL (*CANY)
		p.GetNext()
		cx := p.firstc
		p.GetNext()
		// CLAIRE 4 : check that s is a bound symbol !
		if s.Id() == PRIVATE.Id() {
			return p.ReadThing(app,cx, ToModule(CNULL))
		} else if s == CNULL {return Cerror(29, mname.Id(), CNULL)
		} else if m := ToSymbol(s).Value(); m.Isa != C_module {
					return Cerror(29, s, CNULL)
		} else {
			return p.ReadThing(ToModule(m), cx, def)
		}
	} else {
		s := app.createSymbol(ClEnv.bufferPeek()) // create the symbol
		// if (ClEnv.Verbose == 102) { fmt.Printf(">>>>>>>>>>>>>> read symbol: %s def:%x\n", s.Prt(), s.definition)}
		if app == ClEnv.Module_I || s.definition != nil { // allowed to read
			return EID{s.makeValue(), 0} // return ident value or unbound symbol
		} else {
			fmt.Printf("read private symbol %s ... ERROR \n", s.Prt())
			return Cerror(30, s.Id(), CNULL)
		}
	}
}

func E_read_thing_port(p EID, app EID, cur EID, def EID) EID {
	return ToPort(OBJ(p)).ReadThing(ToModule(OBJ(app)), CHAR(cur), ToModule(OBJ(def)))
}
