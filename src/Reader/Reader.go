/***** CLAIRE Compilation of module Reader.cl 
         [version 4.0.02 / safety 5] Saturday 10-30-2021 *****/

package Reader
import (_ "fmt"
	"unsafe"
	. "Kernel"
	"Core"
	"Language"
)

//-------- dumb function to prevent import errors --------
func import_g0558() { 
    _ = Core.It
    _ = Language.It
    } 
  
  
// class file for delimiter in module Reader 
type Delimiter struct { 
   Core.GlobalVariable
   } 

// automatic cast function
func ToDelimiter(x *ClaireAny) *Delimiter {return (*Delimiter)(unsafe.Pointer(x))}

// automatic constructor function
func MakeDelimiter(name *ClaireSymbol ,value *ClaireAny ,Range *ClaireType ,store_ask *ClaireBoolean ) *Delimiter /* make:0 */{ 
  var o *Delimiter = new(Delimiter)
  o.Isa = C_delimiter
  o.Name = name
  o.Value = value
  o.Range = Range
  o.Store_ask = store_ask
  return o 
  /* make-0 */} 

// class file for reserved_keyword in module Reader 
type ReservedKeyword struct { 
   ClaireKeyword
   } 

// automatic cast function
func ToReservedKeyword(x *ClaireAny) *ReservedKeyword {return (*ReservedKeyword)(unsafe.Pointer(x))}

// automatic constructor function
func MakeReservedKeyword(name *ClaireSymbol ) *ReservedKeyword /* make:0 */{ 
  var o *ReservedKeyword = new(ReservedKeyword)
  o.Isa = C_reserved_keyword
  o.Name = name
  return o 
  /* make-0 */} 

// class file for meta_reader in module Reader 
type MetaReader struct { 
   ClaireThing
   Source *ClaireString 
  SIndex int
  Fromp *ClairePort 
  NbLine int
  External *ClaireString 
  Index int
  LastForm *ClaireAny 
  Maxstack int
  Toplevel *ClaireBoolean 
  Eof int
  Space int
  Tab int
  Bracket *ClaireAny 
  Paren *ClaireAny 
  Comma *ClaireAny 
  Curly *ClaireAny 
  LastArrow *ClaireBoolean 
  SProperties *ClaireSet 
  } 

// automatic cast function
func ToMetaReader(x *ClaireAny) *MetaReader {return (*MetaReader)(unsafe.Pointer(x))}

// class file for measure in module Reader 
type Measure struct { 
   ClaireObject
   MIndex int
  SumValue float64
  SumSquare float64
  NumValue float64
  } 

// automatic cast function
func ToMeasure(x *ClaireAny) *Measure {return (*Measure)(unsafe.Pointer(x))}

// automatic constructor function
func MakeMeasure(m_index int,sum_value float64,sum_square float64,num_value float64) *Measure /* make:0 */{ 
  var o *Measure = new(Measure)
  o.Isa = C_measure
  o.MIndex = m_index
  o.SumValue = sum_value
  o.SumSquare = sum_square
  o.NumValue = num_value
  return o 
  /* make-0 */} 

// class file for PRcount in module Reader 
type PRcount struct { 
   ClaireObject
   Rtime int
  Rdepth int
  Rnum int
  Rloop int
  Rstart int
  } 

// automatic cast function
func To_PRcount(x *ClaireAny) *PRcount {return (*PRcount)(unsafe.Pointer(x))}

var C_delimiter *ClaireClass  /*obj*/
var C_arrow *Core.GlobalVariable 
var C_triangle *Core.GlobalVariable 
var C_reserved_keyword *ClaireClass  /*obj*/
var C_else *ReservedKeyword  /*obj*/
var C_for *ReservedKeyword  /*obj*/
var C_case *ReservedKeyword  /*obj*/
var C_while *ReservedKeyword  /*obj*/
var C_until *ReservedKeyword  /*obj*/
var C_let *ReservedKeyword  /*obj*/
var C_when *ReservedKeyword  /*obj*/
var C_try *ReservedKeyword  /*obj*/
var C_if *ReservedKeyword  /*obj*/
var C_Zif *ReservedKeyword  /*obj*/
var C_branch *ReservedKeyword  /*obj*/
var C_forall *ClaireKeyword  /*obj*/
var C_none *ClaireKeyword  /*obj*/
var C_None *ClaireKeyword  /*obj*/
var C_L__equal *ClaireKeyword  /*obj*/
var C_L_ *ClaireKeyword  /*obj*/
var C_catch *ClaireKeyword  /*obj*/
var C_in *ClaireKeyword  /*obj*/
var C_as *ClaireKeyword  /*obj*/
var C_L_L_ *ClaireKeyword  /*obj*/
var C_printf *ClaireKeyword  /*obj*/
var C_assert *ClaireKeyword  /*obj*/
var C_return *ClaireKeyword  /*obj*/
var C_break *ClaireKeyword  /*obj*/
var C_trace *ClaireKeyword  /*obj*/
var C_exists *ClaireKeyword  /*obj*/
var C_some *ClaireKeyword  /*obj*/
var C__equal_sup *ClaireKeyword  /*obj*/
var C__ask *ClaireKeyword  /*obj*/
var C_rule *ClaireKeyword  /*obj*/
var C_quote *ClaireKeyword  /*obj*/
var C_inspect *ClaireProperty  /*obj*/
var C_known_I *ClaireProperty  /*obj*/
var C_meta_reader *ClaireClass  /*obj*/
var C_AND *Core.GlobalVariable 
var C_OR *Core.GlobalVariable 
var C_Reader_DBline *ClaireTable  /*obj*/
var C_reader *MetaReader  /*obj*/
var C_stdout *Core.GlobalVariable 
var C_stdin *Core.GlobalVariable 
var C__starfs_star *Core.GlobalVariable 
var C_q *ClaireKeyword  /*obj*/
var C_call_debug *ClaireProperty  /*obj*/
var C_EVAL *ClaireTable  /*obj*/
var C_pretty_show *ClaireProperty  /*obj*/
var C_execute_do *ClaireProperty  /*obj*/
var C_execute_bk *ClaireProperty  /*obj*/
var C_inspect_loop *ClaireProperty  /*obj*/
var C_get_from_integer *ClaireProperty  /*obj*/
var C_top_debugger *ClaireProperty  /*obj*/
var C__starlast_star *Core.GlobalVariable 
var C__starindex_star *Core.GlobalVariable 
var C__starmaxd_star *Core.GlobalVariable 
var C__starcurd_star *Core.GlobalVariable 
var C__starshowall_star *Core.GlobalVariable 
var C_Reader_InspectStack *Core.GlobalVariable 
var C_Reader_TopLevelMode *Core.GlobalVariable 
var C_Reader_TopCount *Core.GlobalVariable 
var C_Reader_TopIndex *Core.GlobalVariable 
var C_Reader_TopBase *Core.GlobalVariable 
var C_Reader_TopDebug *Core.GlobalVariable 
var C_up *ClaireProperty  /*obj*/
var C_dn *ClaireProperty  /*obj*/
var C_where *ClaireProperty  /*obj*/
var C_measure *ClaireClass  /*obj*/
var C_PRcount *ClaireClass  /*obj*/
var C_Reader_PRdependent *ClaireTable  /*obj*/
var C_Reader_PRdependentOf *ClaireTable  /*obj*/
var C_Reader_combine_I *ClaireProperty  // Reader/"combine!"
var C_Reader_bind_I *ClaireProperty  // Reader/"bind!"
var C_Reader_CommandLoop *ClaireProperty  // Reader/"CommandLoop"
var C_mean *ClaireProperty  // claire/"mean"
var C_Reader_rtime *ClaireProperty  // Reader/"rtime"
var C_Reader_rnum *ClaireProperty  // Reader/"rnum"
var C_Reader_space *ClaireProperty  // Reader/"space"
var C_Reader_nexti *ClaireProperty  // Reader/"nexti"
var C_kill *ClaireProperty  // claire/"kill"
var C_Reader_rdepth *ClaireProperty  // Reader/"rdepth"
var C_Reader_read_escape *ClaireProperty  // Reader/"read_escape"
var C_Reader_nextstruct *ClaireProperty  // Reader/"nextstruct"
var C_Reader_DBregister *ClaireProperty  // Reader/"DBregister"
var C_Reader_unbind_I *ClaireProperty  // Reader/"unbind!"
var C_addLog *ClaireProperty  // claire/"addLog"
var C_Reader_maxstack *ClaireProperty  // Reader/"maxstack"
var C_readblock *ClaireProperty  // claire/"readblock"
var C_Reader_nextdefinition *ClaireProperty  // Reader/"nextdefinition"
var C_Reader_extended_operator *ClaireProperty  // Reader/"extended_operator"
var C_Reader_readlet *ClaireProperty  // Reader/"readlet"
var C_Reader_readlet_star *ClaireProperty  // Reader/"readlet*"
var C_stdev *ClaireProperty  // claire/"stdev"
var C_Reader_inspect_system *ClaireProperty  // Reader/"inspect_system"
var C_Reader_fromp *ClaireProperty  // Reader/"fromp"
var C_Reader_nexts *ClaireProperty  // Reader/"nexts"
var C_Reader_nexte *ClaireProperty  // Reader/"nexte"
var C_Reader_extended_comment_I *ClaireProperty  // Reader/"extended_comment!"
var C_Reader_cnext *ClaireProperty  // Reader/"cnext"
var C_Reader_debugLoop *ClaireProperty  // Reader/"debugLoop"
var C_Reader_loopexp *ClaireProperty  // Reader/"loopexp"
var C_Reader_rstart *ClaireProperty  // Reader/"rstart"
var C_Reader_toplevel *ClaireProperty  // Reader/"toplevel"
var C_Reader_stop_ask *ClaireProperty  // Reader/"stop?"
var C_Reader_Call_I *ClaireProperty  // Reader/"Call!"
var C_Reader_operation_ask *ClaireProperty  // Reader/"operation?"
var C_PRshow *ClaireProperty  // claire/"PRshow"
var C_PRtime *ClaireProperty  // claire/"PRtime"
var C_Reader_nextunit *ClaireProperty  // Reader/"nextunit"
var C_Reader_nextexp *ClaireProperty  // Reader/"nextexp"
var C_Reader_nextseq *ClaireProperty  // Reader/"nextseq"
var C_Reader_nb_line *ClaireProperty  // Reader/"nb_line"
var C_Reader_nextvariable *ClaireProperty  // Reader/"nextvariable"
var C_Reader_skipc_I *ClaireProperty  // Reader/"skipc!"
var C_Reader_extract_variable *ClaireProperty  // Reader/"extract_variable"
var C_Reader_last_arrow *ClaireProperty  // Reader/"last_arrow"
var C_Reader_combine *ClaireProperty  // Reader/"combine"
var C_mClaire_evaluate *ClaireProperty  // mClaire/"evaluate"
var C_logMeasure *ClaireProperty  // claire/"logMeasure"
var C_stdev_Z *ClaireProperty  // claire/"stdev%"
var C_Reader_last_form *ClaireProperty  // Reader/"last_form"
var C_Reader_checkno *ClaireProperty  // Reader/"checkno"
var C_Reader_sum_square *ClaireProperty  // Reader/"sum_square"
var C_reset *ClaireProperty  // claire/"reset"
var C_Reader_dependents *ClaireProperty  // Reader/"dependents"
var C_Reader_tab *ClaireProperty  // Reader/"tab"
var C_Reader_load_file *ClaireProperty  // Reader/"load_file"
var C_load *ClaireProperty  // claire/"load"
var C_Reader_nexts_I *ClaireProperty  // Reader/"nexts!"
var C_Reader_trace_rule *ClaireProperty  // Reader/"trace_rule"
var C_Reader_dereference *ClaireProperty  // Reader/"dereference"
var C_Reader_nextinst *ClaireProperty  // Reader/"nextinst"
var C_Reader_verify *ClaireProperty  // Reader/"verify"
var C_Reader_print_debug_info *ClaireProperty  // Reader/"print_debug_info"
var C_Reader_num_value *ClaireProperty  // Reader/"num_value"
var C_Reader_keyword_ask *ClaireProperty  // Reader/"keyword?"
var C_sload *ClaireProperty  // claire/"sload"
var C_Reader_debug_if_possible *ClaireProperty  // Reader/"debug_if_possible"
var C_show *ClaireProperty  // claire/"show"
var C_Reader_Show *ClaireProperty  // Reader/"Show"
var C_sqr *ClaireProperty  // claire/"sqr"
var C_Reader_top_level *ClaireProperty  // Reader/"top_level"
var C_PRdepends *ClaireProperty  // claire/"PRdepends"
var C_Reader_readwhen *ClaireProperty  // Reader/"readwhen"
var C_Reader_readcase *ClaireProperty  // Reader/"readcase"
var C_Reader_paren *ClaireProperty  // Reader/"paren"
var C_Reader_Do_I *ClaireProperty  // Reader/"Do!"
var C_Reader_print_exception *ClaireProperty  // Reader/"print_exception"
var C_Reader_untrace *ClaireProperty  // Reader/"untrace"
var C_Reader_closure_build *ClaireProperty  // Reader/"closure_build"
var C_Reader_precedence_I *ClaireProperty  // Reader/"precedence!"
var C_Kernel_call_count *ClaireProperty  // Kernel/"call_count"
var C_block *ClaireProperty  // claire/"block"
var C_Reader_breakpoint *ClaireProperty  // Reader/"breakpoint"
var C_PRget *ClaireProperty  // claire/"PRget"
var C_Reader_s_properties *ClaireProperty  // Reader/"s_properties"
var C_Reader_next *ClaireProperty  // Reader/"next"
var C_Reader_nexte_I *ClaireProperty  // Reader/"nexte!"
var C_Reader_nextmethod *ClaireProperty  // Reader/"nextmethod"
var C_Reader_Serror *ClaireProperty  // Reader/"Serror"
var C_Reader_add_modules *ClaireProperty  // Reader/"add_modules"
var C_PRlook *ClaireProperty  // claire/"PRlook"
var C_Reader_curly *ClaireProperty  // Reader/"curly"
var C_Reader_readset *ClaireProperty  // Reader/"readset"
var C_Reader_extract_of_type *ClaireProperty  // Reader/"extract_of_type"
var C_Reader_nextDefclass *ClaireProperty  // Reader/"nextDefclass"
var C_Reader_bracket *ClaireProperty  // Reader/"bracket"
var C_Reader_firstc *ClaireProperty  // Reader/"firstc"
var C_Reader_eof *ClaireProperty  // Reader/"eof"
var C_Reader_extended_comment_ask *ClaireProperty  // Reader/"extended_comment?"
var C_Reader_operand_I *ClaireProperty  // Reader/"operand!"
var C_Reader_simple_main *ClaireProperty  // Reader/"simple_main"
var C_PRcounter *ClaireProperty  // claire/"PRcounter"
var C_Reader_s_index *ClaireProperty  // Reader/"s_index"
var C_eload *ClaireProperty  // claire/"eload"
var C_Reader_stop *ClaireProperty  // Reader/"stop"
var C_Reader_sum_value *ClaireProperty  // Reader/"sum_value"
var C_Reader_operation_I *ClaireProperty  // Reader/"operation!"
var C_Reader_readif *ClaireProperty  // Reader/"readif"
var C_Reader_findeol *ClaireProperty  // Reader/"findeol"
var C_Reader_comma *ClaireProperty  // Reader/"comma"
var C_Reader_readcall *ClaireProperty  // Reader/"readcall"
var C_Reader_useless_c *ClaireProperty  // Reader/"useless_c"
var C_Reader_skipc *ClaireProperty  // Reader/"skipc"
var C_Reader_m_index *ClaireProperty  // Reader/"m_index"
var C_Reader_rloop *ClaireProperty  // Reader/"rloop"
var It *ClaireModule

// definition of the meta-model for module Reader 
func MetaLoad() { 
  
  It = MakeModule("Reader",Language.C_iClaire)
  ClEnv.Module_I = It
  // definition of the properties 
  
  C_Reader_precedence_I = MakeProperty("precedence!",1,It)
  C_Kernel_call_count = MakeProperty("call_count",2,C_Kernel)
  C_block = MakeProperty("block",1,C_claire)
  C_Reader_add_modules = MakeProperty("add_modules",1,It)
  C_Reader_breakpoint = MakeProperty("breakpoint",1,It)
  C_PRget = MakeProperty("PRget",1,C_claire)
  C_Reader_s_properties = MakeProperty("s_properties",1,It)
  C_Reader_next = MakeProperty("next",1,It)
  C_Reader_nexte_I = MakeProperty("nexte!",1,It)
  C_Reader_nextmethod = MakeProperty("nextmethod",1,It)
  C_Reader_Serror = MakeProperty("Serror",1,It)
  C_PRlook = MakeProperty("PRlook",1,C_claire)
  C_Reader_curly = MakeProperty("curly",1,It)
  C_Reader_readset = MakeProperty("readset",1,It)
  C_Reader_extract_of_type = MakeProperty("extract_of_type",1,It)
  C_Reader_nextDefclass = MakeProperty("nextDefclass",1,It)
  C_Reader_bracket = MakeProperty("bracket",1,It)
  C_Reader_firstc = MakeProperty("firstc",1,It)
  C_Reader_eof = MakeProperty("eof",1,It)
  C_Reader_extended_comment_ask = MakeProperty("extended_comment?",1,It)
  C_Reader_operand_I = MakeProperty("operand!",1,It)
  C_Reader_simple_main = MakeProperty("simple_main",1,It)
  C_PRcounter = MakeProperty("PRcounter",1,C_claire)
  C_Reader_s_index = MakeProperty("s_index",1,It)
  C_eload = MakeProperty("eload",1,C_claire)
  C_Reader_stop = MakeProperty("stop",1,It)
  C_Reader_sum_value = MakeProperty("sum_value",1,It)
  C_Reader_operation_I = MakeProperty("operation!",1,It)
  C_Reader_readif = MakeProperty("readif",1,It)
  C_Reader_findeol = MakeProperty("findeol",1,It)
  C_Reader_comma = MakeProperty("comma",1,It)
  C_Reader_readcall = MakeProperty("readcall",1,It)
  C_Reader_useless_c = MakeProperty("useless_c",1,It)
  C_Reader_skipc = MakeProperty("skipc",1,It)
  C_Reader_m_index = MakeProperty("m_index",1,It)
  C_Reader_rloop = MakeProperty("rloop",2,It)
  C_Reader_rnum = MakeProperty("rnum",1,It)
  C_Reader_combine_I = MakeProperty("combine!",1,It)
  C_Reader_bind_I = MakeProperty("bind!",1,It)
  C_Reader_CommandLoop = MakeProperty("CommandLoop",1,It)
  C_mean = MakeProperty("mean",1,C_claire)
  C_Reader_rtime = MakeProperty("rtime",1,It)
  C_Reader_space = MakeProperty("space",1,It)
  C_Reader_nexti = MakeProperty("nexti",1,It)
  C_kill = MakeProperty("kill",1,C_claire)
  C_Reader_rdepth = MakeProperty("rdepth",2,It)
  C_Reader_read_escape = MakeProperty("read_escape",1,It)
  C_Reader_nextstruct = MakeProperty("nextstruct",1,It)
  C_Reader_DBregister = MakeProperty("DBregister",1,It)
  C_Reader_unbind_I = MakeProperty("unbind!",1,It)
  C_addLog = MakeProperty("addLog",1,C_claire)
  C_Reader_maxstack = MakeProperty("maxstack",1,It)
  C_readblock = MakeProperty("readblock",1,C_claire)
  C_Reader_nextdefinition = MakeProperty("nextdefinition",1,It)
  C_Reader_extended_operator = MakeProperty("extended_operator",1,It)
  C_Reader_readlet = MakeProperty("readlet",1,It)
  C_Reader_readlet_star = MakeProperty("readlet*",1,It)
  C_stdev = MakeProperty("stdev",1,C_claire)
  C_Reader_debugLoop = MakeProperty("debugLoop",1,It)
  C_Reader_inspect_system = MakeProperty("inspect_system",1,It)
  C_Reader_fromp = MakeProperty("fromp",1,It)
  C_Reader_nexts = MakeProperty("nexts",1,It)
  C_Reader_nexte = MakeProperty("nexte",1,It)
  C_Reader_extended_comment_I = MakeProperty("extended_comment!",1,It)
  C_Reader_cnext = MakeProperty("cnext",1,It)
  C_Reader_loopexp = MakeProperty("loopexp",1,It)
  C_Reader_rstart = MakeProperty("rstart",2,It)
  C_Reader_toplevel = MakeProperty("toplevel",1,It)
  C_Reader_stop_ask = MakeProperty("stop?",1,It)
  C_Reader_Call_I = MakeProperty("Call!",1,It)
  C_Reader_operation_ask = MakeProperty("operation?",1,It)
  C_PRshow = MakeProperty("PRshow",1,C_claire)
  C_PRtime = MakeProperty("PRtime",1,C_claire)
  C_Reader_nextunit = MakeProperty("nextunit",1,It)
  C_Reader_nextexp = MakeProperty("nextexp",1,It)
  C_Reader_nextseq = MakeProperty("nextseq",1,It)
  C_Reader_nb_line = MakeProperty("nb_line",1,It)
  C_Reader_nextvariable = MakeProperty("nextvariable",1,It)
  C_Reader_skipc_I = MakeProperty("skipc!",1,It)
  C_Reader_extract_variable = MakeProperty("extract_variable",1,It)
  C_Reader_last_arrow = MakeProperty("last_arrow",2,It)
  C_Reader_combine = MakeProperty("combine",1,It)
  C_mClaire_evaluate = MakeProperty("evaluate",2,C_mClaire)
  C_logMeasure = MakeProperty("logMeasure",1,C_claire)
  C_stdev_Z = MakeProperty("stdev%",1,C_claire)
  C_Reader_last_form = MakeProperty("last_form",1,It)
  C_Reader_checkno = MakeProperty("checkno",1,It)
  C_Reader_sum_square = MakeProperty("sum_square",1,It)
  C_reset = MakeProperty("reset",1,C_claire)
  C_Reader_dependents = MakeProperty("dependents",1,It)
  C_Reader_tab = MakeProperty("tab",1,It)
  C_Reader_load_file = MakeProperty("load_file",1,It)
  C_load = MakeProperty("load",1,C_claire)
  C_Reader_nexts_I = MakeProperty("nexts!",1,It)
  C_Reader_trace_rule = MakeProperty("trace_rule",1,It)
  C_Reader_dereference = MakeProperty("dereference",1,It)
  C_Reader_nextinst = MakeProperty("nextinst",1,It)
  C_Reader_verify = MakeProperty("verify",1,It)
  C_Reader_print_debug_info = MakeProperty("print_debug_info",1,It)
  C_Reader_num_value = MakeProperty("num_value",1,It)
  C_Reader_keyword_ask = MakeProperty("keyword?",1,It)
  C_sload = MakeProperty("sload",1,C_claire)
  C_Reader_debug_if_possible = MakeProperty("debug_if_possible",1,It)
  C_show = MakeProperty("show",1,C_claire)
  C_Reader_Show = MakeProperty("Show",1,It)
  C_sqr = MakeProperty("sqr",1,C_claire)
  C_Reader_top_level = MakeProperty("top_level",1,It)
  C_PRdepends = MakeProperty("PRdepends",1,C_claire)
  C_Reader_readwhen = MakeProperty("readwhen",1,It)
  C_Reader_readcase = MakeProperty("readcase",1,It)
  C_Reader_paren = MakeProperty("paren",1,It)
  C_Reader_Do_I = MakeProperty("Do!",1,It)
  C_Reader_print_exception = MakeProperty("print_exception",1,It)
  C_Reader_untrace = MakeProperty("untrace",1,It)
  C_Reader_closure_build = MakeProperty("closure_build",1,It)
  
  // instructions from module sources
  C_delimiter = MakeClass("delimiter",Core.C_global_variable,C_claire)
  
  { 
    /*PROTECT let _CL_obj:global_variable := (<Compile/object!(claire/"arrow", global_variable):global_variable>) in (range(_CL_obj) := any, value(_CL_obj) := let _CL_obj:keyword := (<mClaire/new! @ list<type_expression>(class)(keyword):keyword>) in (name(_CL_obj) := symbol! @ list<type_expression>(string)("->"), _CL_obj), close @ global_variable(_CL_obj)) */
    var expr EID 
    /* Let:2 */{ 
      var _CL_obj *Core.GlobalVariable  
      /* noccur = 9 */
      /* object!:claire/"arrow" ->Reader*/C_arrow = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("arrow",C_claire)))
      
      _CL_obj = C_arrow
      _CL_obj.Range = ToType(C_any.Id())
      /* update:3 */{ 
        var va_arg1 *Core.GlobalVariable  
        var va_arg2 *ClaireAny  
        va_arg1 = _CL_obj
        /* Let:4 */{ 
          var _CL_obj *ClaireKeyword   = ToKeyword(new(ClaireKeyword).Is(C_keyword))
          /* noccur = 3 */
          _CL_obj.Name = Core.F_symbol_I_string2(MakeString("->"))
          va_arg2 = _CL_obj.Id()
          /* Let-4 */} 
        /* ---------- now we compile update value(va_arg1) := va_arg2 ------- */
        va_arg1.Value = va_arg2
        /* update-3 */} 
      expr = _CL_obj.Close()
      /* Let-2 */} 
    ErrorCheck(expr)} 
  
  ToSymbol(OBJ(Core.F_CALL(C_name,ARGS(EID{C_arrow.Value,0})))).Put(C_arrow.Value)
  
  { 
    /*PROTECT let _CL_obj:global_variable := (<Compile/object!(claire/"triangle", global_variable):global_variable>) in (range(_CL_obj) := any, value(_CL_obj) := let _CL_obj:keyword := (<mClaire/new! @ list<type_expression>(class)(keyword):keyword>) in (name(_CL_obj) := symbol! @ list<type_expression>(string)("<:"), _CL_obj), close @ global_variable(_CL_obj)) */
    var expr EID 
    /* Let:2 */{ 
      var _CL_obj *Core.GlobalVariable  
      /* noccur = 9 */
      /* object!:claire/"triangle" ->Reader*/C_triangle = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("triangle",C_claire)))
      
      _CL_obj = C_triangle
      _CL_obj.Range = ToType(C_any.Id())
      /* update:3 */{ 
        var va_arg1 *Core.GlobalVariable  
        var va_arg2 *ClaireAny  
        va_arg1 = _CL_obj
        /* Let:4 */{ 
          var _CL_obj *ClaireKeyword   = ToKeyword(new(ClaireKeyword).Is(C_keyword))
          /* noccur = 3 */
          _CL_obj.Name = Core.F_symbol_I_string2(MakeString("<:"))
          va_arg2 = _CL_obj.Id()
          /* Let-4 */} 
        /* ---------- now we compile update value(va_arg1) := va_arg2 ------- */
        va_arg1.Value = va_arg2
        /* update-3 */} 
      expr = _CL_obj.Close()
      /* Let-2 */} 
    ErrorCheck(expr)} 
  
  C_reserved_keyword = MakeClass("reserved_keyword",C_keyword,C_claire)
  
  /* object!:claire/"else" ->Reader*/C_else = ToReservedKeyword(new(ReservedKeyword).IsNamed(C_reserved_keyword,MakeSymbol("else",C_claire)))
  
  
  /* object!:claire/"for" ->Reader*/C_for = ToReservedKeyword(new(ReservedKeyword).IsNamed(C_reserved_keyword,MakeSymbol("for",C_claire)))
  
  
  /* object!:claire/"case" ->Reader*/C_case = ToReservedKeyword(new(ReservedKeyword).IsNamed(C_reserved_keyword,MakeSymbol("case",C_claire)))
  
  
  /* object!:claire/"while" ->Reader*/C_while = ToReservedKeyword(new(ReservedKeyword).IsNamed(C_reserved_keyword,MakeSymbol("while",C_claire)))
  
  
  /* object!:claire/"until" ->Reader*/C_until = ToReservedKeyword(new(ReservedKeyword).IsNamed(C_reserved_keyword,MakeSymbol("until",C_claire)))
  
  
  /* object!:claire/"let" ->Reader*/C_let = ToReservedKeyword(new(ReservedKeyword).IsNamed(C_reserved_keyword,MakeSymbol("let",C_claire)))
  
  
  /* object!:claire/"when" ->Reader*/C_when = ToReservedKeyword(new(ReservedKeyword).IsNamed(C_reserved_keyword,MakeSymbol("when",C_claire)))
  
  
  /* object!:claire/"try" ->Reader*/C_try = ToReservedKeyword(new(ReservedKeyword).IsNamed(C_reserved_keyword,MakeSymbol("try",C_claire)))
  
  
  /* object!:claire/"if" ->Reader*/C_if = ToReservedKeyword(new(ReservedKeyword).IsNamed(C_reserved_keyword,MakeSymbol("if",C_claire)))
  
  
  /* object!:claire/"Zif" ->Reader*/C_Zif = ToReservedKeyword(new(ReservedKeyword).IsNamed(C_reserved_keyword,MakeSymbol("Zif",C_claire)))
  
  
  /* object!:claire/"branch" ->Reader*/C_branch = ToReservedKeyword(new(ReservedKeyword).IsNamed(C_reserved_keyword,MakeSymbol("branch",C_claire)))
  
  
  Core.F_attach_method(C_Reader_keyword_ask.AddMethod(Signature(C_any.Id(),C_boolean.Id()),0,MakeFunction1(E_keyword_ask_any,"keyword_ask_any")),MakeString("read.cl:49"))
  
  /* object!:claire/"forall" ->Reader*/C_forall = ToKeyword(new(ClaireKeyword).IsNamed(C_keyword,MakeSymbol("forall",C_claire)))
  
  
  /* object!:claire/"none" ->Reader*/C_none = ToKeyword(new(ClaireKeyword).IsNamed(C_keyword,MakeSymbol("none",C_claire)))
  
  
  /* object!:claire/"None" ->Reader*/C_None = ToKeyword(new(ClaireKeyword).IsNamed(C_keyword,MakeSymbol("None",C_claire)))
  
  
  /* object!:claire/":=" ->Reader*/C_L__equal = ToKeyword(new(ClaireKeyword).IsNamed(C_keyword,MakeSymbol(":=",C_claire)))
  
  
  /* object!:claire/":" ->Reader*/C_L_ = ToKeyword(new(ClaireKeyword).IsNamed(C_keyword,MakeSymbol(":",C_claire)))
  
  
  /* object!:claire/"catch" ->Reader*/C_catch = ToKeyword(new(ClaireKeyword).IsNamed(C_keyword,MakeSymbol("catch",C_claire)))
  
  
  /* object!:claire/"in" ->Reader*/C_in = ToKeyword(new(ClaireKeyword).IsNamed(C_keyword,MakeSymbol("in",C_claire)))
  
  
  /* object!:claire/"as" ->Reader*/C_as = ToKeyword(new(ClaireKeyword).IsNamed(C_keyword,MakeSymbol("as",C_claire)))
  
  
  /* object!:claire/"::" ->Reader*/C_L_L_ = ToKeyword(new(ClaireKeyword).IsNamed(C_keyword,MakeSymbol("::",C_claire)))
  
  
  /* object!:claire/"printf" ->Reader*/C_printf = ToKeyword(new(ClaireKeyword).IsNamed(C_keyword,MakeSymbol("printf",C_claire)))
  
  
  /* object!:claire/"assert" ->Reader*/C_assert = ToKeyword(new(ClaireKeyword).IsNamed(C_keyword,MakeSymbol("assert",C_claire)))
  
  
  /* object!:claire/"return" ->Reader*/C_return = ToKeyword(new(ClaireKeyword).IsNamed(C_keyword,MakeSymbol("return",C_claire)))
  
  
  /* object!:claire/"break" ->Reader*/C_break = ToKeyword(new(ClaireKeyword).IsNamed(C_keyword,MakeSymbol("break",C_claire)))
  
  
  /* object!:claire/"trace" ->Reader*/C_trace = ToKeyword(new(ClaireKeyword).IsNamed(C_keyword,MakeSymbol("trace",C_claire)))
  
  
  /* object!:claire/"exists" ->Reader*/C_exists = ToKeyword(new(ClaireKeyword).IsNamed(C_keyword,MakeSymbol("exists",C_claire)))
  
  
  /* object!:claire/"some" ->Reader*/C_some = ToKeyword(new(ClaireKeyword).IsNamed(C_keyword,MakeSymbol("some",C_claire)))
  
  
  /* object!:claire/"=>" ->Reader*/C__equal_sup = ToKeyword(new(ClaireKeyword).IsNamed(C_keyword,MakeSymbol("=>",C_claire)))
  
  
  /* object!:claire/"?" ->Reader*/C__ask = ToKeyword(new(ClaireKeyword).IsNamed(C_keyword,MakeSymbol("?",C_claire)))
  
  
  /* object!:claire/"rule" ->Reader*/C_rule = ToKeyword(new(ClaireKeyword).IsNamed(C_keyword,MakeSymbol("rule",C_claire)))
  
  
  /* object!:claire/"quote" ->Reader*/C_quote = ToKeyword(new(ClaireKeyword).IsNamed(C_keyword,MakeSymbol("quote",C_claire)))
  
  
  C_inspect = MakeProperty("inspect",1,C_claire)
  
  
  C_known_I = MakeProperty("known!",2,C_claire)
  
  
  C_meta_reader = MakeClass("meta_reader",C_thing,C_claire)
  Core.F_close_slot(C_meta_reader.AddSlot(C_source,ToType(C_string.Id()),CNULL))
  Core.F_close_slot(C_meta_reader.AddSlot(C_Reader_s_index,ToType(C_integer.Id()),MakeInteger(0).Id()))
  Core.F_close_slot(C_meta_reader.AddSlot(C_Reader_fromp,ToType(C_port.Id()),CNULL))
  Core.F_close_slot(C_meta_reader.AddSlot(C_Reader_nb_line,ToType(C_integer.Id()),MakeInteger(0).Id()))
  Core.F_close_slot(C_meta_reader.AddSlot(C_external,ToType(C_string.Id()),MakeString("toplevel").Id()))
  Core.F_close_slot(C_meta_reader.AddSlot(C_mClaire_index,ToType(C_integer.Id()),MakeInteger(0).Id()))
  Core.F_close_slot(C_meta_reader.AddSlot(C_Reader_last_form,ToType(C_any.Id()),CNULL))
  Core.F_close_slot(C_meta_reader.AddSlot(C_Reader_maxstack,ToType(C_integer.Id()),MakeInteger(0).Id()))
  Core.F_close_slot(C_meta_reader.AddSlot(C_Reader_toplevel,ToType(C_boolean.Id()),CNULL))
  Core.F_close_slot(C_meta_reader.AddSlot(C_Reader_eof,ToType(C_integer.Id()),MakeInteger(0).Id()))
  Core.F_close_slot(C_meta_reader.AddSlot(C_Reader_space,ToType(C_integer.Id()),MakeInteger(0).Id()))
  Core.F_close_slot(C_meta_reader.AddSlot(C_Reader_tab,ToType(C_integer.Id()),MakeInteger(0).Id()))
  Core.F_close_slot(C_meta_reader.AddSlot(C_Reader_bracket,ToType(C_any.Id()),CNULL))
  Core.F_close_slot(C_meta_reader.AddSlot(C_Reader_paren,ToType(C_any.Id()),CNULL))
  Core.F_close_slot(C_meta_reader.AddSlot(C_Reader_comma,ToType(C_any.Id()),CNULL))
  Core.F_close_slot(C_meta_reader.AddSlot(C_Reader_curly,ToType(C_any.Id()),CNULL))
  Core.F_close_slot(C_meta_reader.AddSlot(C_Reader_last_arrow,ToType(C_boolean.Id()),CFALSE.Id()))
  Core.F_close_slot(C_meta_reader.AddSlot(C_Reader_s_properties,Core.F_nth_class1(C_set,ToType(C_property.Id())),MakeSet(ToType(C_property.Id()),C_ephemeral.Id(),
    C_begin.Id(),
    C_end.Id(),
    C_store.Id(),
    Core.C_reify.Id(),
    C_known_I.Id(),
    C_abstract.Id(),
    C_final.Id()).Id()))
  
  Core.F_attach_method(C_Reader_next.AddMethod(Signature(C_meta_reader.Id(),C_void.Id()),0,MakeFunction1(E_next_meta_reader,"next_meta_reader")),MakeString("read.cl:102"))
  
  Core.F_attach_method(C_Reader_firstc.AddMethod(Signature(C_meta_reader.Id(),C_integer.Id()),0,MakeFunction1(E_firstc_meta_reader,"firstc_meta_reader")),MakeString("read.cl:103"))
  
  Core.F_attach_method(C_Reader_stop_ask.AddMethod(Signature(C_integer.Id(),C_any.Id()),0,MakeFunction1(E_stop_ask_integer,"stop_ask_integer")),MakeString("read.cl:106"))
  
  { 
    /*PROTECT let _CL_obj:global_variable := (<Compile/object!(claire/"AND", global_variable):global_variable>) in (range(_CL_obj) := any, value(_CL_obj) := &, close @ global_variable(_CL_obj)) */
    var expr EID 
    /* Let:2 */{ 
      var _CL_obj *Core.GlobalVariable  
      /* noccur = 5 */
      /* object!:claire/"AND" ->Reader*/C_AND = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("AND",C_claire)))
      
      _CL_obj = C_AND
      _CL_obj.Range = ToType(C_any.Id())
      _CL_obj.Value = Core.C__and.Id()
      expr = _CL_obj.Close()
      /* Let-2 */} 
    ErrorCheck(expr)} 
  
  { 
    /*PROTECT let _CL_obj:global_variable := (<Compile/object!(claire/"OR", global_variable):global_variable>) in (range(_CL_obj) := any, value(_CL_obj) := mClaire/new! @ list<type_expression>(class, symbol)(delimiter,symbol! @ list<type_expression>(string, module)("|",claire)), close @ global_variable(_CL_obj)) */
    var expr EID 
    /* Let:2 */{ 
      var _CL_obj *Core.GlobalVariable  
      /* noccur = 5 */
      /* object!:claire/"OR" ->Reader*/C_OR = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("OR",C_claire)))
      
      _CL_obj = C_OR
      _CL_obj.Range = ToType(C_any.Id())
      /* update:3 */{ 
        var va_arg1 *Core.GlobalVariable  
        var va_arg2 *ClaireAny  
        va_arg1 = _CL_obj
        var va_arg2_try05594 EID 
        va_arg2_try05594 = new(Delimiter).IsNamed(C_delimiter,F_symbol_I_string(MakeString("|"),C_claire)).ToEID()
        /* ERROR PROTECTION INSERTED (va_arg2-expr) */
        if ErrorIn(va_arg2_try05594) {expr = va_arg2_try05594
        } else {
        va_arg2 = ANY(va_arg2_try05594)
        /* ---------- now we compile update value(va_arg1) := va_arg2 ------- */
        va_arg1.Value = va_arg2
        expr = va_arg2.ToEID()
        }
        /* update-3 */} 
      /* ERROR PROTECTION INSERTED (expr-expr) */
      if !ErrorIn(expr) {
      expr = _CL_obj.Close()
      }
      /* Let-2 */} 
    ErrorCheck(expr)} 
  
  Core.F_attach_method(C_Reader_nextunit.AddMethod(Signature(C_meta_reader.Id(),C_any.Id()),1,MakeFunction1(E_nextunit_meta_reader,"nextunit_meta_reader")),MakeString("read.cl:138"))
  
  Core.F_attach_method(C_Reader_nexts.AddMethod(Signature(C_meta_reader.Id(),C_keyword.Id(),C_any.Id()),1,MakeFunction2(E_nexts_meta_reader,"nexts_meta_reader")),MakeString("read.cl:154"))
  
  Core.F_attach_method(C_Reader_loopexp.AddMethod(Signature(C_meta_reader.Id(),
    C_any.Id(),
    C_keyword.Id(),
    C_boolean.Id(),
    C_any.Id()),1,MakeFunction4(E_loopexp_meta_reader,"loopexp_meta_reader")),MakeString("read.cl:188"))
  
  Core.F_attach_method(C_Reader_extended_operator.AddMethod(Signature(C_property.Id(),
    C_any.Id(),
    C_any.Id(),
    C_any.Id()),1,MakeFunction3(E_extended_operator_property,"extended_operator_property")),MakeString("read.cl:199"))
  
  Core.F_attach_method(C_Reader_nexte.AddMethod(Signature(C_meta_reader.Id(),C_any.Id()),1,MakeFunction1(E_nexte_meta_reader,"nexte_meta_reader")),MakeString("read.cl:208"))
  
  Core.F_attach_method(C_Reader_nextexp.AddMethod(Signature(C_meta_reader.Id(),C_boolean.Id(),C_any.Id()),1,MakeFunction2(E_nextexp_meta_reader,"nextexp_meta_reader")),MakeString("read.cl:256"))
  
  Core.F_attach_method(C_Reader_nexti.AddMethod(Signature(C_meta_reader.Id(),C_any.Id(),C_any.Id()),1,MakeFunction2(E_nexti_meta_reader,"nexti_meta_reader")),MakeString("read.cl:305"))
  
  Core.F_attach_method(C_Reader_read_escape.AddMethod(Signature(C_meta_reader.Id(),C_any.Id()),1,MakeFunction1(E_read_escape_meta_reader,"read_escape_meta_reader")),MakeString("read.cl:316"))
  
  Core.F_attach_method(C_Reader_nextvariable.AddMethod(Signature(C_meta_reader.Id(),C_any.Id(),C_any.Id()),1,MakeFunction2(E_nextvariable_meta_reader,"nextvariable_meta_reader")),MakeString("read.cl:330"))
  
  Core.F_attach_method(C_Reader_nexts_I.AddMethod(Signature(C_meta_reader.Id(),C_keyword.Id(),C_any.Id()),1,MakeFunction2(E_nexts_I_meta_reader1,"nexts_I_meta_reader1")),MakeString("read.cl:337"))
  
  Core.F_attach_method(C_Reader_nexte_I.AddMethod(Signature(C_meta_reader.Id(),C_keyword.Id(),C_any.Id()),1,MakeFunction2(E_nexte_I_meta_reader,"nexte_I_meta_reader")),MakeString("read.cl:344"))
  
  Core.F_attach_method(C_Reader_nexts_I.AddMethod(Signature(C_meta_reader.Id(),C_integer.Id(),C_any.Id()),1,MakeFunction2(E_nexts_I_meta_reader2,"nexts_I_meta_reader2")),MakeString("read.cl:350"))
  
  Core.F_attach_method(C_Reader_nexts_I.AddMethod(Signature(C_meta_reader.Id(),
    C_keyword.Id(),
    C_integer.Id(),
    C_any.Id()),1,MakeFunction3(E_nexts_I_meta_reader3,"nexts_I_meta_reader3")),MakeString("read.cl:357"))
  
  Core.F_attach_method(C_Reader_extended_comment_ask.AddMethod(Signature(C_meta_reader.Id(),C_string.Id(),C_boolean.Id()),0,MakeFunction2(E_extended_comment_ask_meta_reader,"extended_comment_ask_meta_reader")),MakeString("read.cl:364"))
  
  Core.F_attach_method(C_Reader_extended_comment_I.AddMethod(Signature(C_meta_reader.Id(),C_string.Id(),C_any.Id()),1,MakeFunction2(E_extended_comment_I_meta_reader,"extended_comment_I_meta_reader")),MakeString("read.cl:387"))
  
  Core.F_attach_method(C_Reader_operation_ask.AddMethod(Signature(C_any.Id(),C_boolean.Id()),0,MakeFunction1(E_operation_ask_any,"operation_ask_any")),MakeString("syntax.cl:27"))
  
  Core.F_attach_method(C_Reader_combine.AddMethod(Signature(C_any.Id(),
    C_any.Id(),
    C_any.Id(),
    C_any.Id()),1,MakeFunction3(E_combine_any,"combine_any")),MakeString("syntax.cl:35"))
  
  Core.F_attach_method(C_Reader_combine_I.AddMethod(Signature(C_any.Id(),
    C_any.Id(),
    C_any.Id(),
    C_any.Id()),1,MakeFunction3(E_combine_I_any,"combine_I_any")),MakeString("syntax.cl:73"))
  
  Core.F_attach_method(C_Reader_operation_I.AddMethod(Signature(C_any.Id(),C_any.Id()),0,MakeFunction1(E_operation_I_any,"operation_I_any")),MakeString("syntax.cl:87"))
  
  Core.F_attach_method(C_Reader_operand_I.AddMethod(Signature(C_any.Id(),C_integer.Id(),C_any.Id()),1,MakeFunction2(E_operand_I_any,"operand_I_any")),MakeString("syntax.cl:105"))
  
  Core.F_attach_method(C_Reader_precedence_I.AddMethod(Signature(C_any.Id(),C_integer.Id()),0,MakeFunction1(E_precedence_I_any,"precedence_I_any")),MakeString("syntax.cl:114"))
  
  Core.F_attach_method(C_Reader_nextstruct.AddMethod(Signature(C_meta_reader.Id(),
    C_keyword.Id(),
    C_keyword.Id(),
    C_any.Id()),1,MakeFunction3(E_nextstruct_meta_reader,"nextstruct_meta_reader")),MakeString("syntax.cl:137"))
  
  Core.F_attach_method(C_Reader_readlet.AddMethod(Signature(C_meta_reader.Id(),C_keyword.Id(),C_any.Id()),1,MakeFunction2(E_readlet_meta_reader,"readlet_meta_reader")),MakeString("syntax.cl:166"))
  
  Core.F_attach_method(C_Reader_readlet_star.AddMethod(Signature(C_meta_reader.Id(),
    C_list.Id(),
    C_integer.Id(),
    C_keyword.Id(),
    C_any.Id()),1,MakeFunction4(E_readlet_star_meta_reader,"readlet_star_meta_reader")),MakeString("syntax.cl:177"))
  
  Core.F_attach_method(C_Reader_readwhen.AddMethod(Signature(C_meta_reader.Id(),C_keyword.Id(),C_any.Id()),1,MakeFunction2(E_readwhen_meta_reader,"readwhen_meta_reader")),MakeString("syntax.cl:192"))
  
  Core.F_attach_method(C_Reader_readif.AddMethod(Signature(C_meta_reader.Id(),C_integer.Id(),C_any.Id()),1,MakeFunction2(E_readif_meta_reader,"readif_meta_reader")),MakeString("syntax.cl:206"))
  
  Core.F_attach_method(C_Reader_readcase.AddMethod(Signature(C_meta_reader.Id(),C_keyword.Id(),C_any.Id()),1,MakeFunction2(E_readcase_meta_reader,"readcase_meta_reader")),MakeString("syntax.cl:224"))
  
  Core.F_attach_method(C_Reader_readset.AddMethod(Signature(C_meta_reader.Id(),C_any.Id(),C_any.Id()),1,MakeFunction2(E_readset_meta_reader,"readset_meta_reader")),MakeString("syntax.cl:254"))
  
  Core.F_attach_method(C_Reader_dereference.AddMethod(Signature(C_any.Id(),C_any.Id()),1,MakeFunction1(E_dereference_any,"dereference_any")),MakeString("syntax.cl:259"))
  
  Core.F_attach_method(C_Reader_nextseq.AddMethod(Signature(C_meta_reader.Id(),C_integer.Id(),C_any.Id()),1,MakeFunction2(E_nextseq_meta_reader,"nextseq_meta_reader")),MakeString("syntax.cl:271"))
  
  Core.F_attach_method(C_readblock.AddMethod(Signature(C_meta_reader.Id(),
    C_any.Id(),
    C_integer.Id(),
    C_any.Id()),1,MakeFunction3(E_readblock_meta_reader,"readblock_meta_reader")),MakeString("syntax.cl:289"))
  
  Core.F_attach_method(C_Reader_Do_I.AddMethod(Signature(C_any.Id(),C_any.Id(),C_any.Id()),1,MakeFunction2(E_Do_I_any,"Do_I_any")),MakeString("syntax.cl:295"))
  
  Core.F_attach_method(C_Reader_extract_of_type.AddMethod(Signature(Language.C_Call.Id(),C_type.Id()),1,MakeFunction1(E_extract_of_type_Call,"extract_of_type_Call")),MakeString("syntax.cl:306"))
  
  /* object!:Reader/"DBline" ->Reader*/C_Reader_DBline = ToTable(new(ClaireTable).IsNamed(C_table,MakeSymbol("DBline",It)))
  C_Reader_DBline.Range = ToType(C_integer.Id())
  C_Reader_DBline.Params = C_any.Id()
  C_Reader_DBline.Domain = ToType(Language.C_Call.Id())
  C_Reader_DBline.GraphInit()
  C_Reader_DBline.Default = MakeInteger(0).Id()
  
  Core.F_attach_method(C_Reader_DBregister.AddMethod(Signature(Language.C_Call.Id(),Language.C_Call.Id()),1,MakeFunction1(E_DBregister_Call,"DBregister_Call")),MakeString("syntax.cl:324"))
  
  Core.F_attach_method(C_Reader_Call_I.AddMethod(Signature(C_property.Id(),C_list.Id(),Language.C_Call.Id()),1,MakeFunction2(E_Call_I_property,"Call_I_property")),MakeString("syntax.cl:326"))
  
  Core.F_attach_method(C_Reader_readcall.AddMethod(Signature(C_meta_reader.Id(),
    C_any.Id(),
    C_any.Id(),
    C_any.Id()),1,MakeFunction3(E_readcall_meta_reader,"readcall_meta_reader")),MakeString("syntax.cl:368"))
  
  Core.F_attach_method(C_Reader_nextdefinition.AddMethod(Signature(C_meta_reader.Id(),
    C_any.Id(),
    C_any.Id(),
    C_boolean.Id(),
    C_any.Id()),1,MakeFunction4(E_nextdefinition_meta_reader,"nextdefinition_meta_reader")),MakeString("syntax.cl:401"))
  
  Core.F_attach_method(C_Reader_nextmethod.AddMethod(Signature(C_meta_reader.Id(),
    C_any.Id(),
    C_any.Id(),
    C_boolean.Id(),
    C_boolean.Id(),
    C_boolean.Id(),
    C_any.Id()),1,MakeFunction6(E_nextmethod_meta_reader,"nextmethod_meta_reader")),MakeString("syntax.cl:413"))
  
  Core.F_attach_method(C_Reader_nextinst.AddMethod(Signature(C_meta_reader.Id(),C_any.Id(),C_any.Id()),1,MakeFunction2(E_nextinst_meta_reader,"nextinst_meta_reader")),MakeString("syntax.cl:434"))
  
  Core.F_attach_method(C_Reader_nextDefclass.AddMethod(Signature(C_meta_reader.Id(),
    C_any.Id(),
    C_boolean.Id(),
    Language.C_Defclass.Id()),1,MakeFunction3(E_nextDefclass_meta_reader,"nextDefclass_meta_reader")),MakeString("syntax.cl:463"))
  
  Core.F_attach_method(Core.C_self_eval.AddEvalMethod(Signature(C_delimiter.Id(),C_void.Id()),1,MakeFunction1(E_self_eval_delimiter,"self_eval_delimiter"),EVAL_delimiter),MakeString("file.cl:30"))
  
  Core.F_attach_method(C_Reader_useless_c.AddMethod(Signature(C_integer.Id(),C_boolean.Id()),0,MakeFunction1(E_useless_c_integer,"useless_c_integer")),MakeString("file.cl:41"))
  
  Core.F_attach_method(C_Reader_skipc.AddMethod(Signature(C_meta_reader.Id(),C_integer.Id()),0,MakeFunction1(E_skipc_meta_reader,"skipc_meta_reader")),MakeString("file.cl:48"))
  
  Core.F_attach_method(C_Reader_skipc_I.AddMethod(Signature(C_meta_reader.Id(),C_integer.Id()),1,MakeFunction1(E_skipc_I_meta_reader,"skipc_I_meta_reader")),MakeString("file.cl:60"))
  
  Core.F_attach_method(C_Reader_cnext.AddMethod(Signature(C_meta_reader.Id(),C_meta_reader.Id()),0,MakeFunction1(E_cnext_meta_reader,"cnext_meta_reader")),MakeString("file.cl:62"))
  
  Core.F_attach_method(C_Reader_findeol.AddMethod(Signature(C_meta_reader.Id(),C_boolean.Id()),0,MakeFunction1(E_findeol_meta_reader,"findeol_meta_reader")),MakeString("file.cl:66"))
  
  Core.F_attach_method(C_Reader_checkno.AddMethod(Signature(C_meta_reader.Id(),
    C_integer.Id(),
    C_any.Id(),
    C_any.Id()),1,MakeFunction3(E_checkno_meta_reader,"checkno_meta_reader")),MakeString("file.cl:72"))
  
  Core.F_attach_method(C_Reader_verify.AddMethod(Signature(C_any.Id(),
    C_any.Id(),
    C_any.Id(),
    C_any.Id()),1,MakeFunction3(E_verify_any,"verify_any")),MakeString("file.cl:78"))
  
  Core.F_attach_method(C_Reader_Serror.AddMethod(Signature(C_string.Id(),C_list.Id(),CEMPTY.Id()),1,MakeFunction2(E_Serror_string,"Serror_string")),MakeString("file.cl:85"))
  
  { 
    /*PROTECT (Compile/object!(claire/"reader", meta_reader), Reader/space(reader) := 202, Reader/eof(reader) := -1, Reader/tab(reader) := 9, mClaire/index(reader) := 1, external(reader) := "toplevel", Reader/bracket(reader) := mClaire/new! @ list<type_expression>(class, symbol)(delimiter,symbol! @ list<type_expression>(string)("]")), Reader/paren(reader) := mClaire/new! @ list<type_expression>(class, symbol)(delimiter,symbol! @ list<type_expression>(string)(")")), Reader/comma(reader) := mClaire/new! @ list<type_expression>(class, symbol)(delimiter,symbol! @ list<type_expression>(string)(",")), Reader/curly(reader) := mClaire/new! @ list<type_expression>(class, symbol)(delimiter,symbol! @ list<type_expression>(string)("}")), reader) */
    var expr EID 
    /* object!:claire/"reader" ->Reader*/C_reader = ToMetaReader(new(MetaReader).IsNamed(C_meta_reader,MakeSymbol("reader",C_claire)))
    C_reader.Space = 202
    C_reader.Eof = -1
    C_reader.Tab = 9
    C_reader.Index = 1
    C_reader.External = MakeString("toplevel")
    /* update:2 */{ 
      var va_arg1 *MetaReader  
      var va_arg2 *ClaireAny  
      va_arg1 = C_reader
      var va_arg2_try05603 EID 
      va_arg2_try05603 = new(Delimiter).IsNamed(C_delimiter,Core.F_symbol_I_string2(MakeString("]"))).ToEID()
      /* ERROR PROTECTION INSERTED (va_arg2-expr) */
      if ErrorIn(va_arg2_try05603) {expr = va_arg2_try05603
      } else {
      va_arg2 = ANY(va_arg2_try05603)
      /* ---------- now we compile update Reader/bracket(va_arg1) := va_arg2 ------- */
      va_arg1.Bracket = va_arg2
      expr = va_arg2.ToEID()
      }
      /* update-2 */} 
    /* ERROR PROTECTION INSERTED (expr-expr) */
    if !ErrorIn(expr) {
    /* update:2 */{ 
      var va_arg1 *MetaReader  
      var va_arg2 *ClaireAny  
      va_arg1 = C_reader
      var va_arg2_try05613 EID 
      va_arg2_try05613 = new(Delimiter).IsNamed(C_delimiter,Core.F_symbol_I_string2(MakeString(")"))).ToEID()
      /* ERROR PROTECTION INSERTED (va_arg2-expr) */
      if ErrorIn(va_arg2_try05613) {expr = va_arg2_try05613
      } else {
      va_arg2 = ANY(va_arg2_try05613)
      /* ---------- now we compile update Reader/paren(va_arg1) := va_arg2 ------- */
      va_arg1.Paren = va_arg2
      expr = va_arg2.ToEID()
      }
      /* update-2 */} 
    /* ERROR PROTECTION INSERTED (expr-expr) */
    if !ErrorIn(expr) {
    /* update:2 */{ 
      var va_arg1 *MetaReader  
      var va_arg2 *ClaireAny  
      va_arg1 = C_reader
      var va_arg2_try05623 EID 
      va_arg2_try05623 = new(Delimiter).IsNamed(C_delimiter,Core.F_symbol_I_string2(MakeString(","))).ToEID()
      /* ERROR PROTECTION INSERTED (va_arg2-expr) */
      if ErrorIn(va_arg2_try05623) {expr = va_arg2_try05623
      } else {
      va_arg2 = ANY(va_arg2_try05623)
      /* ---------- now we compile update Reader/comma(va_arg1) := va_arg2 ------- */
      va_arg1.Comma = va_arg2
      expr = va_arg2.ToEID()
      }
      /* update-2 */} 
    /* ERROR PROTECTION INSERTED (expr-expr) */
    if !ErrorIn(expr) {
    /* update:2 */{ 
      var va_arg1 *MetaReader  
      var va_arg2 *ClaireAny  
      va_arg1 = C_reader
      var va_arg2_try05633 EID 
      va_arg2_try05633 = new(Delimiter).IsNamed(C_delimiter,Core.F_symbol_I_string2(MakeString("}"))).ToEID()
      /* ERROR PROTECTION INSERTED (va_arg2-expr) */
      if ErrorIn(va_arg2_try05633) {expr = va_arg2_try05633
      } else {
      va_arg2 = ANY(va_arg2_try05633)
      /* ---------- now we compile update Reader/curly(va_arg1) := va_arg2 ------- */
      va_arg1.Curly = va_arg2
      expr = va_arg2.ToEID()
      }
      /* update-2 */} 
    /* ERROR PROTECTION INSERTED (expr-expr) */
    if !ErrorIn(expr) {
    expr = EID{C_reader.Id(),0}
    }}}}
    ErrorCheck(expr)} 
  
  Core.F_attach_method(C_Reader_extract_variable.AddMethod(Signature(C_any.Id(),C_Variable.Id()),1,MakeFunction1(E_extract_variable_any,"extract_variable_any")),MakeString("file.cl:107"))
  
  Core.F_attach_method(C_Reader_bind_I.AddMethod(Signature(C_meta_reader.Id(),C_Variable.Id(),C_list.Id()),0,MakeFunction2(E_bind_I_meta_reader,"bind_I_meta_reader")),MakeString("file.cl:116"))
  
  Core.F_attach_method(C_Reader_unbind_I.AddMethod(Signature(C_meta_reader.Id(),C_list.Id(),C_any.Id()),0,MakeFunction2(E_unbind_I_meta_reader,"unbind_I_meta_reader")),MakeString("file.cl:123"))
  
  { 
    /*PROTECT let _CL_obj:global_variable := (<Compile/object!(claire/"stdout", global_variable):global_variable>) in (range(_CL_obj) := port, value(_CL_obj) := externC @ list<type_expression>(string)("ClEnv.Cout.Id()"), close @ global_variable(_CL_obj)) */
    var expr EID 
    /* Let:2 */{ 
      var _CL_obj *Core.GlobalVariable  
      /* noccur = 5 */
      /* object!:claire/"stdout" ->Reader*/C_stdout = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("stdout",C_claire)))
      
      _CL_obj = C_stdout
      _CL_obj.Range = ToType(C_port.Id())
      _CL_obj.Value = ClEnv.Cout.Id()
      expr = _CL_obj.Close()
      /* Let-2 */} 
    ErrorCheck(expr)} 
  
  ClEnv.Ctrace = ToPort(C_stdout.Value)
  
  { 
    /*PROTECT let _CL_obj:global_variable := (<Compile/object!(claire/"stdin", global_variable):global_variable>) in (range(_CL_obj) := port, value(_CL_obj) := externC @ list<type_expression>(string)("ClEnv.Cin.Id()"), close @ global_variable(_CL_obj)) */
    var expr EID 
    /* Let:2 */{ 
      var _CL_obj *Core.GlobalVariable  
      /* noccur = 5 */
      /* object!:claire/"stdin" ->Reader*/C_stdin = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("stdin",C_claire)))
      
      _CL_obj = C_stdin
      _CL_obj.Range = ToType(C_port.Id())
      _CL_obj.Value = ClEnv.Cin.Id()
      expr = _CL_obj.Close()
      /* Let-2 */} 
    ErrorCheck(expr)} 
  
  { 
    /*PROTECT let _CL_obj:global_variable := (<Compile/object!(claire/"*fs*", global_variable):global_variable>) in (range(_CL_obj) := string, value(_CL_obj) := "/", close @ global_variable(_CL_obj)) */
    var expr EID 
    /* Let:2 */{ 
      var _CL_obj *Core.GlobalVariable  
      /* noccur = 5 */
      /* object!:claire/"*fs*" ->Reader*/C__starfs_star = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("*fs*",C_claire)))
      
      _CL_obj = C__starfs_star
      _CL_obj.Range = ToType(C_string.Id())
      _CL_obj.Value = MakeString("/").Id()
      expr = _CL_obj.Close()
      /* Let-2 */} 
    ErrorCheck(expr)} 
  
  Core.F_attach_method(C__7.AddMethod(Signature(C_string.Id(),C_string.Id(),C_string.Id()),0,MakeFunction2(E__7_string,"_7_string")),MakeString("file.cl:135"))
  
  Core.F_attach_method(C_mClaire_restore_state.AddMethod(Signature(C_meta_reader.Id(),C_void.Id()),0,MakeFunction1(E_restore_state_meta_reader,"restore_state_meta_reader")),MakeString("file.cl:148"))
  
  Core.F_attach_method(C_Reader_load_file.AddMethod(Signature(C_string.Id(),C_boolean.Id(),C_any.Id()),1,MakeFunction2(E_load_file_string,"load_file_string")),MakeString("file.cl:199"))
  
  Core.F_attach_method(C_load.AddMethod(Signature(C_string.Id(),C_any.Id()),1,MakeFunction1(E_load_string,"load_string")),MakeString("file.cl:203"))
  
  Core.F_attach_method(C_sload.AddMethod(Signature(C_string.Id(),C_any.Id()),1,MakeFunction1(E_sload_string,"sload_string")),MakeString("file.cl:204"))
  
  Core.F_attach_method(C_Reader_load_file.AddMethod(Signature(C_module.Id(),C_boolean.Id(),C_void.Id()),1,MakeFunction2(E_load_file_module,"load_file_module")),MakeString("file.cl:218"))
  
  Core.F_attach_method(C_load.AddMethod(Signature(C_module.Id(),C_any.Id()),1,MakeFunction1(E_load_module,"load_module")),MakeString("file.cl:223"))
  
  Core.F_attach_method(C_sload.AddMethod(Signature(C_module.Id(),C_any.Id()),1,MakeFunction1(E_sload_module,"sload_module")),MakeString("file.cl:225"))
  
  Core.F_attach_method(C_Reader_add_modules.AddMethod(Signature(C_module.Id(),
    C_set.Id(),
    C_list.Id(),
    C_list.Id()),0,MakeFunction3(E_add_modules_module,"add_modules_module")),MakeString("file.cl:240"))
  
  Core.F_attach_method(C_Reader_add_modules.AddMethod(Signature(C_list.Id(),C_list.Id()),0,MakeFunction1(E_add_modules_list,"add_modules_list")),MakeString("file.cl:247"))
  
  Core.F_attach_method(C_eload.AddMethod(Signature(C_string.Id(),C_any.Id()),1,MakeFunction1(E_eload_string,"eload_string")),MakeString("file.cl:280"))
  
  Core.F_attach_method(C_readblock.AddMethod(Signature(C_port.Id(),C_any.Id()),1,MakeFunction1(E_readblock_port,"readblock_port")),MakeString("file.cl:300"))
  
  Core.F_attach_method(C_read.AddMethod(Signature(C_port.Id(),C_any.Id()),1,MakeFunction1(E_read_port,"read_port")),MakeString("file.cl:312"))
  
  Core.F_attach_method(C_read.AddMethod(Signature(C_string.Id(),C_any.Id()),1,MakeFunction1(E_read_string,"read_string")),MakeString("file.cl:327"))
  
  /* object!:claire/"q" ->Reader*/C_q = ToKeyword(new(ClaireKeyword).IsNamed(C_keyword,MakeSymbol("q",C_claire)))
  
  
  C_call_debug = MakeProperty("call_debug",2,C_claire)
  
  
  /* object!:claire/"EVAL" ->Reader*/C_EVAL = ToTable(new(ClaireTable).IsNamed(C_table,MakeSymbol("EVAL",C_claire)))
  C_EVAL.Range = ToType(C_any.Id())
  C_EVAL.Params = MakeInteger(-1).Id()
  C_EVAL.Domain = Core.F__dot_dot_integer(0,99)
  C_EVAL.Graph = Core.F_make_copy_list_integer(100,CNULL).Id()
  C_EVAL.Default = CNULL
  
  Core.F_attach_method(C_Reader_debug_if_possible.AddMethod(Signature(C_void.Id(),C_any.Id()),1,MakeFunction1(E_debug_if_possible_void,"debug_if_possible_void")),MakeString("file.cl:339"))
  
  Core.F_attach_method(C_Reader_print_exception.AddMethod(Signature(C_void.Id(),C_any.Id()),0,MakeFunction1(E_print_exception_void,"print_exception_void")),MakeString("file.cl:353"))
  
  C_pretty_show = MakeProperty("pretty_show",3,C_claire)
  C_pretty_show.Open = 3
  
  
  Core.F_attach_method(C_show.AddMethod(Signature(C_any.Id(),C_any.Id()),1,MakeFunction1(E_show_any,"show_any")),MakeString("file.cl:372"))
  
  Core.F_attach_method(C_kill.AddMethod(Signature(C_object.Id(),C_any.Id()),0,MakeFunction1(E_kill_object,"kill_object")),MakeString("file.cl:380"))
  
  Core.F_attach_method(C_kill.AddMethod(Signature(C_class.Id(),C_any.Id()),0,MakeFunction1(E_kill_class,"kill_class")),MakeString("file.cl:385"))
  
  Core.F_inlineok_ask_method(
  Core.F_attach_method(C_min.AddMethod(Signature(C_integer.Id(),C_integer.Id(),C_integer.Id()),0,MakeFunction2(E_min_integer,"min_integer")),MakeString("file.cl:388")),
  MakeString("lambda[(x:integer,y:integer),(if (x <= y) x else y)]"))
  
  Core.F_inlineok_ask_method(
  Core.F_attach_method(C_max.AddMethod(Signature(C_integer.Id(),C_integer.Id(),C_integer.Id()),0,MakeFunction2(E_max_integer,"max_integer")),MakeString("file.cl:389")),
  MakeString("lambda[(x:integer,y:integer),(if (x <= y) y else x)]"))
  
  Core.F_inlineok_ask_method(
  Core.F_attach_method(C_min.AddMethod(Signature(C_float.Id(),C_float.Id(),C_float.Id()),0,MakeFunction2(E_min_float,"min_float")),MakeString("file.cl:391")),
  MakeString("lambda[(x:float,y:float),(if (x <= y) x else y)]"))
  
  Core.F_inlineok_ask_method(
  Core.F_attach_method(C_max.AddMethod(Signature(C_float.Id(),C_float.Id(),C_float.Id()),0,MakeFunction2(E_max_float,"max_float")),MakeString("file.cl:392")),
  MakeString("lambda[(x:float,y:float),(if (x <= y) y else x)]"))
  
  Core.F_inlineok_ask_method(
  Core.F_attach_method(C_min.AddMethod(Signature(C_any.Id(),C_any.Id(),C_any.Id()),1,MakeFunction2(E_min_any,"min_any")),MakeString("file.cl:394")),
  MakeString("lambda[(x:any,y:any),(if (x <= y) x else y)]"))
  
  Core.F_inlineok_ask_method(
  Core.F_attach_method(C_max.AddMethod(Signature(C_any.Id(),C_any.Id(),C_any.Id()),1,MakeFunction2(E_max_any,"max_any")),MakeString("file.cl:395")),
  MakeString("lambda[(x:any,y:any),(if (x <= y) y else x)]"))
  
  Core.F_inlineok_ask_method(
  Core.F_attach_method(Core.C_known_ask.AddMethod(Signature(C_table.Id(),C_any.Id(),C_boolean.Id()),0,MakeFunction2(E_known_ask_table,"known_ask_table")),MakeString("file.cl:398")),
  MakeString("lambda[(a:table,x:any),get(a, x) != unknown]"))
  
  Core.F_inlineok_ask_method(
  Core.F_attach_method(Core.C_unknown_ask.AddMethod(Signature(C_table.Id(),C_any.Id(),C_boolean.Id()),0,MakeFunction2(E_unknown_ask_table,"unknown_ask_table")),MakeString("file.cl:399")),
  MakeString("lambda[(a:table,x:any),get(a, x) = unknown]"))
  
  Core.F_attach_method(C_float_I.AddMethod(Signature(C_string.Id(),C_float.Id()),1,MakeFunction1(E_float_I_string,"float_I_string")),MakeString("file.cl:405"))
  
  Core.F_inlineok_ask_method(
  Core.F_attach_method(C__sup_equal.AddMethod(Signature(C_any.Id(),C_any.Id(),C_boolean.Id()),1,MakeFunction2(E__sup_equal_any,"_sup_equal_any")),MakeString("file.cl:408")),
  MakeString("lambda[(self:any,x:any),x <= self]"))
  
  Core.F_inlineok_ask_method(
  Core.F_attach_method(Core.C__plus.AddMethod(Signature(C_integer.Id(),C_float.Id(),C_float.Id()),0,MakeFunction2(E__plus_integer2,"_plus_integer2")),MakeString("file.cl:411")),
  MakeString("lambda[(x:integer,y:float),float!(x) + y]"))
  
  Core.F_inlineok_ask_method(
  Core.F_attach_method(C__star.AddMethod(Signature(C_integer.Id(),C_float.Id(),C_float.Id()),0,MakeFunction2(E__star_integer2,"_star_integer2")),MakeString("file.cl:412")),
  MakeString("lambda[(x:integer,y:float),float!(x) * y]"))
  
  Core.F_inlineok_ask_method(
  Core.F_attach_method(C__7.AddMethod(Signature(C_integer.Id(),C_float.Id(),C_float.Id()),0,MakeFunction2(E__7_integer2,"_7_integer2")),MakeString("file.cl:413")),
  MakeString("lambda[(x:integer,y:float),float!(x) / y]"))
  
  Core.F_inlineok_ask_method(
  Core.F_attach_method(C__dash.AddMethod(Signature(C_integer.Id(),C_float.Id(),C_float.Id()),0,MakeFunction2(E__dash_integer3,"_dash_integer3")),MakeString("file.cl:414")),
  MakeString("lambda[(x:integer,y:float),float!(x) - y]"))
  
  Core.F_inlineok_ask_method(
  Core.F_attach_method(Core.C__plus.AddMethod(Signature(C_float.Id(),C_integer.Id(),C_float.Id()),0,MakeFunction2(E__plus_float2,"_plus_float2")),MakeString("file.cl:415")),
  MakeString("lambda[(x:float,y:integer),x + float!(y)]"))
  
  Core.F_inlineok_ask_method(
  Core.F_attach_method(C__star.AddMethod(Signature(C_float.Id(),C_integer.Id(),C_float.Id()),0,MakeFunction2(E__star_float2,"_star_float2")),MakeString("file.cl:416")),
  MakeString("lambda[(x:float,y:integer),x * float!(y)]"))
  
  Core.F_inlineok_ask_method(
  Core.F_attach_method(C__7.AddMethod(Signature(C_float.Id(),C_integer.Id(),C_float.Id()),0,MakeFunction2(E__7_float2,"_7_float2")),MakeString("file.cl:417")),
  MakeString("lambda[(x:float,y:integer),x / float!(y)]"))
  
  Core.F_inlineok_ask_method(
  Core.F_attach_method(C__dash.AddMethod(Signature(C_float.Id(),C_integer.Id(),C_float.Id()),0,MakeFunction2(E__dash_float3,"_dash_float3")),MakeString("file.cl:418")),
  MakeString("lambda[(x:float,y:integer),x - float!(y)]"))
  
  Core.F_inlineok_ask_method(
  Core.F_attach_method(C_sqr.AddMethod(Signature(C_integer.Id(),C_integer.Id()),0,MakeFunction1(E_sqr_integer,"sqr_integer")),MakeString("file.cl:421")),
  MakeString("lambda[(x:integer),x * x]"))
  
  Core.F_inlineok_ask_method(
  Core.F_attach_method(C_sqr.AddMethod(Signature(C_float.Id(),C_float.Id()),0,MakeFunction1(E_sqr_float,"sqr_float")),MakeString("file.cl:422")),
  MakeString("lambda[(x:float),x * x]"))
  
  C_execute_do = MakeProperty("execute_do",2,C_claire)
  
  
  C_execute_bk = MakeProperty("execute_bk",2,C_claire)
  
  
  C_inspect_loop = MakeProperty("inspect_loop",1,C_claire)
  
  
  C_get_from_integer = MakeProperty("get_from_integer",1,C_claire)
  
  
  C_top_debugger = MakeProperty("top_debugger",2,C_claire)
  
  
  { 
    /*PROTECT let _CL_obj:global_variable := (<Compile/object!(claire/"*last*", global_variable):global_variable>) in (range(_CL_obj) := any, value(_CL_obj) := unknown, close @ global_variable(_CL_obj)) */
    var expr EID 
    /* Let:2 */{ 
      var _CL_obj *Core.GlobalVariable  
      /* noccur = 5 */
      /* object!:claire/"*last*" ->Reader*/C__starlast_star = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("*last*",C_claire)))
      
      _CL_obj = C__starlast_star
      _CL_obj.Range = ToType(C_any.Id())
      _CL_obj.Value = CNULL
      expr = _CL_obj.Close()
      /* Let-2 */} 
    ErrorCheck(expr)} 
  
  { 
    /*PROTECT let _CL_obj:global_variable := (<Compile/object!(claire/"*index*", global_variable):global_variable>) in (range(_CL_obj) := integer, value(_CL_obj) := unknown, close @ global_variable(_CL_obj)) */
    var expr EID 
    /* Let:2 */{ 
      var _CL_obj *Core.GlobalVariable  
      /* noccur = 5 */
      /* object!:claire/"*index*" ->Reader*/C__starindex_star = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("*index*",C_claire)))
      
      _CL_obj = C__starindex_star
      _CL_obj.Range = ToType(C_integer.Id())
      _CL_obj.Value = CNULL
      expr = _CL_obj.Close()
      /* Let-2 */} 
    ErrorCheck(expr)} 
  
  { 
    /*PROTECT let _CL_obj:global_variable := (<Compile/object!(claire/"*maxd*", global_variable):global_variable>) in (range(_CL_obj) := integer, value(_CL_obj) := unknown, close @ global_variable(_CL_obj)) */
    var expr EID 
    /* Let:2 */{ 
      var _CL_obj *Core.GlobalVariable  
      /* noccur = 5 */
      /* object!:claire/"*maxd*" ->Reader*/C__starmaxd_star = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("*maxd*",C_claire)))
      
      _CL_obj = C__starmaxd_star
      _CL_obj.Range = ToType(C_integer.Id())
      _CL_obj.Value = CNULL
      expr = _CL_obj.Close()
      /* Let-2 */} 
    ErrorCheck(expr)} 
  
  { 
    /*PROTECT let _CL_obj:global_variable := (<Compile/object!(claire/"*curd*", global_variable):global_variable>) in (range(_CL_obj) := integer, value(_CL_obj) := 0, close @ global_variable(_CL_obj)) */
    var expr EID 
    /* Let:2 */{ 
      var _CL_obj *Core.GlobalVariable  
      /* noccur = 5 */
      /* object!:claire/"*curd*" ->Reader*/C__starcurd_star = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("*curd*",C_claire)))
      
      _CL_obj = C__starcurd_star
      _CL_obj.Range = ToType(C_integer.Id())
      _CL_obj.Value = MakeInteger(0).Id()
      expr = _CL_obj.Close()
      /* Let-2 */} 
    ErrorCheck(expr)} 
  
  { 
    /*PROTECT let _CL_obj:global_variable := (<Compile/object!(claire/"*showall*", global_variable):global_variable>) in (range(_CL_obj) := boolean, value(_CL_obj) := true, close @ global_variable(_CL_obj)) */
    var expr EID 
    /* Let:2 */{ 
      var _CL_obj *Core.GlobalVariable  
      /* noccur = 5 */
      /* object!:claire/"*showall*" ->Reader*/C__starshowall_star = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("*showall*",C_claire)))
      
      _CL_obj = C__starshowall_star
      _CL_obj.Range = ToType(C_boolean.Id())
      _CL_obj.Value = CTRUE.Id()
      expr = _CL_obj.Close()
      /* Let-2 */} 
    ErrorCheck(expr)} 
  
  Core.F_attach_method(C_Reader_CommandLoop.AddMethod(Signature(C_void.Id(),C_string.Id()),0,MakeFunction1(E_CommandLoopVoid,"#CommandLoopVoid")),MakeString("inspect.cl:35"))
  
  { 
    /*PROTECT let _CL_obj:global_variable := (<Compile/object!(Reader/"InspectStack", global_variable):global_variable>) in (range(_CL_obj) := list, value(_CL_obj) := nil, close @ global_variable(_CL_obj)) */
    var expr EID 
    /* Let:2 */{ 
      var _CL_obj *Core.GlobalVariable  
      /* noccur = 5 */
      /* object!:Reader/"InspectStack" ->Reader*/C_Reader_InspectStack = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("InspectStack",It)))
      
      _CL_obj = C_Reader_InspectStack
      _CL_obj.Range = ToType(C_list.Id())
      _CL_obj.Value = CNIL.Id()
      expr = _CL_obj.Close()
      /* Let-2 */} 
    ErrorCheck(expr)} 
  
  { 
    /*PROTECT let _CL_obj:global_variable := (<Compile/object!(Reader/"TopLevelMode", global_variable):global_variable>) in (range(_CL_obj) := integer, value(_CL_obj) := 1, close @ global_variable(_CL_obj)) */
    var expr EID 
    /* Let:2 */{ 
      var _CL_obj *Core.GlobalVariable  
      /* noccur = 5 */
      /* object!:Reader/"TopLevelMode" ->Reader*/C_Reader_TopLevelMode = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("TopLevelMode",It)))
      
      _CL_obj = C_Reader_TopLevelMode
      _CL_obj.Range = ToType(C_integer.Id())
      _CL_obj.Value = MakeInteger(1).Id()
      expr = _CL_obj.Close()
      /* Let-2 */} 
    ErrorCheck(expr)} 
  
  { 
    /*PROTECT let _CL_obj:global_variable := (<Compile/object!(Reader/"TopCount", global_variable):global_variable>) in (range(_CL_obj) := integer, value(_CL_obj) := 0, close @ global_variable(_CL_obj)) */
    var expr EID 
    /* Let:2 */{ 
      var _CL_obj *Core.GlobalVariable  
      /* noccur = 5 */
      /* object!:Reader/"TopCount" ->Reader*/C_Reader_TopCount = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("TopCount",It)))
      
      _CL_obj = C_Reader_TopCount
      _CL_obj.Range = ToType(C_integer.Id())
      _CL_obj.Value = MakeInteger(0).Id()
      expr = _CL_obj.Close()
      /* Let-2 */} 
    ErrorCheck(expr)} 
  
  { 
    /*PROTECT let _CL_obj:global_variable := (<Compile/object!(Reader/"TopIndex", global_variable):global_variable>) in (range(_CL_obj) := integer, value(_CL_obj) := 0, close @ global_variable(_CL_obj)) */
    var expr EID 
    /* Let:2 */{ 
      var _CL_obj *Core.GlobalVariable  
      /* noccur = 5 */
      /* object!:Reader/"TopIndex" ->Reader*/C_Reader_TopIndex = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("TopIndex",It)))
      
      _CL_obj = C_Reader_TopIndex
      _CL_obj.Range = ToType(C_integer.Id())
      _CL_obj.Value = MakeInteger(0).Id()
      expr = _CL_obj.Close()
      /* Let-2 */} 
    ErrorCheck(expr)} 
  
  { 
    /*PROTECT let _CL_obj:global_variable := (<Compile/object!(Reader/"TopBase", global_variable):global_variable>) in (range(_CL_obj) := integer, value(_CL_obj) := 0, close @ global_variable(_CL_obj)) */
    var expr EID 
    /* Let:2 */{ 
      var _CL_obj *Core.GlobalVariable  
      /* noccur = 5 */
      /* object!:Reader/"TopBase" ->Reader*/C_Reader_TopBase = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("TopBase",It)))
      
      _CL_obj = C_Reader_TopBase
      _CL_obj.Range = ToType(C_integer.Id())
      _CL_obj.Value = MakeInteger(0).Id()
      expr = _CL_obj.Close()
      /* Let-2 */} 
    ErrorCheck(expr)} 
  
  { 
    /*PROTECT let _CL_obj:global_variable := (<Compile/object!(Reader/"TopDebug", global_variable):global_variable>) in (range(_CL_obj) := integer, value(_CL_obj) := 0, close @ global_variable(_CL_obj)) */
    var expr EID 
    /* Let:2 */{ 
      var _CL_obj *Core.GlobalVariable  
      /* noccur = 5 */
      /* object!:Reader/"TopDebug" ->Reader*/C_Reader_TopDebug = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("TopDebug",It)))
      
      _CL_obj = C_Reader_TopDebug
      _CL_obj.Range = ToType(C_integer.Id())
      _CL_obj.Value = MakeInteger(0).Id()
      expr = _CL_obj.Close()
      /* Let-2 */} 
    ErrorCheck(expr)} 
  
  Core.F_attach_method(C_Reader_top_level.AddMethod(Signature(C_meta_reader.Id(),C_void.Id()),1,MakeFunction1(E_Reader_top_level_meta_reader,"Reader_top_level_meta_reader")),MakeString("inspect.cl:83"))
  
  Core.F_attach_method(C_Reader_debugLoop.AddMethod(Signature(C_meta_reader.Id(),C_void.Id()),0,MakeFunction1(E_Reader_debugLoop_meta_reader,"Reader_debugLoop_meta_reader")),MakeString("inspect.cl:93"))
  
  Core.F_attach_method(C_Reader_inspect_system.AddMethod(Signature(C_list.Id(),C_void.Id()),0,MakeFunction1(E_Reader_inspect_system_list2,"Reader_inspect_system_list2")),MakeString("inspect.cl:99"))
  
  Core.F_attach_method(C_Reader_simple_main.AddMethod(Signature(C_void.Id(),C_void.Id()),1,MakeFunction1(E_Reader_simple_main_void,"Reader_simple_main_void")),MakeString("inspect.cl:123"))
  
  Core.F_attach_method(C_inspect.AddMethod(Signature(C_any.Id(),C_any.Id()),1,MakeFunction1(E_inspect_any,"inspect_any")),MakeString("inspect.cl:148"))
  
  Core.F_attach_method(C_inspect_loop.AddMethod(Signature(C_any.Id(),C_list.Id(),C_void.Id()),1,MakeFunction2(E_inspect_loop_any,"inspect_loop_any")),MakeString("inspect.cl:169"))
  
  Core.F_attach_method(C_get_from_integer.AddMethod(Signature(C_any.Id(),C_integer.Id(),C_any.Id()),1,MakeFunction2(E_get_from_integer_any,"get_from_integer_any")),MakeString("inspect.cl:180"))
  
  Core.F_attach_method(Language.C_iClaire_trace_on.AddMethod(Signature(C_any.Id(),C_any.Id()),1,MakeFunction1(E_trace_on_any,"trace_on_any")),MakeString("inspect.cl:202"))
  
  Core.F_attach_method(C_Reader_untrace.AddMethod(Signature(C_any.Id(),C_any.Id()),1,MakeFunction1(E_untrace_any,"untrace_any")),MakeString("inspect.cl:214"))
  
  Core.F_attach_method(Core.C_spy.AddMethod(Signature(C_listargs.Id(),C_void.Id()),0,MakeFunction1(E_spy_listargs2_Reader,"spy_listargs2_Reader")),MakeString("inspect.cl:222"))
  
  Core.F_attach_method(C_Reader_trace_rule.AddMethod(Signature(C_relation.Id(),
    C_string.Id(),
    C_any.Id(),
    C_any.Id(),
    C_any.Id(),
    C_any.Id(),
    C_void.Id()),1,MakeFunction6(E_trace_rule_relation,"trace_rule_relation")),MakeString("inspect.cl:232"))
  
  Core.F_attach_method(C_Reader_stop.AddMethod(Signature(C_property.Id(),C_listargs.Id(),C_any.Id()),1,MakeFunction2(E_stop_property,"stop_property")),MakeString("inspect.cl:240"))
  
  Core.F_attach_method(Core.C_debug.AddMethod(Signature(C_void.Id(),C_void.Id()),0,MakeFunction1(E_debug_void,"debug_void")),MakeString("inspect.cl:253"))
  
  Core.F_attach_method(C_call_debug.AddMethod(Signature(C_void.Id(),C_any.Id()),0,MakeFunction1(E_call_debug_void,"call_debug_void")),MakeString("inspect.cl:274"))
  
  Core.F_attach_method(C_Reader_breakpoint.AddMethod(Signature(C_void.Id(),C_void.Id()),1,MakeFunction1(E_breakpoint_void,"breakpoint_void")),MakeString("inspect.cl:299"))
  
  C_up = MakeProperty("up",2,C_claire)
  
  
  C_dn = MakeProperty("dn",2,C_claire)
  
  
  C_where = MakeProperty("where",2,C_claire)
  
  
  Core.F_attach_method(C_dn.AddMethod(Signature(C_integer.Id(),C_void.Id()),0,MakeFunction1(E_dn_integer,"dn_integer")),MakeString("inspect.cl:309"))
  
  Core.F_attach_method(C_up.AddMethod(Signature(C_integer.Id(),C_void.Id()),0,MakeFunction1(E_up_integer,"up_integer")),MakeString("inspect.cl:318"))
  
  Core.F_attach_method(C_where.AddMethod(Signature(C_integer.Id(),C_void.Id()),1,MakeFunction1(E_where_integer,"where_integer")),MakeString("inspect.cl:328"))
  
  Core.F_attach_method(C_Reader_print_debug_info.AddMethod(Signature(C_integer.Id(),
    C_integer.Id(),
    C_integer.Id(),
    C_void.Id()),1,MakeFunction3(E_print_debug_info_integer,"print_debug_info_integer")),MakeString("inspect.cl:340"))
  
  Core.F_attach_method(C_Reader_Show.AddMethod(Signature(C_integer.Id(),C_any.Id()),1,MakeFunction1(E_Show_integer,"Show_integer")),MakeString("inspect.cl:353"))
  
  Core.F_attach_method(C_block.AddMethod(Signature(C_integer.Id(),C_void.Id()),1,MakeFunction1(E_block_integer,"block_integer")),MakeString("inspect.cl:378"))
  
  Core.F_attach_method(C_Reader_closure_build.AddMethod(Signature(C_lambda.Id(),C_list.Id()),0,MakeFunction1(E_closure_build_lambda,"closure_build_lambda")),MakeString("inspect.cl:386"))
  
  Core.F_attach_method(C_Reader_closure_build.AddMethod(Signature(C_any.Id(),C_list.Id(),C_void.Id()),0,MakeFunction2(E_closure_build_any,"closure_build_any")),MakeString("inspect.cl:397"))
  
  C_measure = MakeClass("measure",C_object,C_claire)
  Core.F_close_slot(C_measure.AddSlot(C_Reader_m_index,ToType(C_integer.Id()),MakeInteger(1).Id()))
  Core.F_close_slot(C_measure.AddSlot(C_Reader_sum_value,ToType(C_float.Id()),MakeFloat(0).Id()))
  Core.F_close_slot(C_measure.AddSlot(C_Reader_sum_square,ToType(C_float.Id()),MakeFloat(0).Id()))
  Core.F_close_slot(C_measure.AddSlot(C_Reader_num_value,ToType(C_float.Id()),MakeFloat(0).Id()))
  
  Core.F_attach_method(C_close.AddMethod(Signature(C_measure.Id(),C_measure.Id()),0,MakeFunction1(E_close_measure,"close_measure")),MakeString("inspect.cl:414"))
  
  Core.F_attach_method(C_add.AddMethod(Signature(C_measure.Id(),C_float.Id(),C_measure.Id()),0,MakeFunction2(E_add_measure,"add_measure")),MakeString("inspect.cl:416"))
  
  Core.F_attach_method(C_mean.AddMethod(Signature(C_measure.Id(),C_float.Id()),0,MakeFunction1(E_mean_measure,"mean_measure")),MakeString("inspect.cl:418"))
  
  Core.F_attach_method(C_stdev.AddMethod(Signature(C_measure.Id(),C_float.Id()),0,MakeFunction1(E_stdev_measure,"stdev_measure")),MakeString("inspect.cl:421"))
  
  Core.F_attach_method(C_stdev_Z.AddMethod(Signature(C_measure.Id(),C_float.Id()),0,MakeFunction1(E_stdev_Z_measure,"stdev_Z_measure")),MakeString("inspect.cl:422"))
  
  Core.F_attach_method(C_reset.AddMethod(Signature(C_measure.Id(),C_void.Id()),0,MakeFunction1(E_reset_measure,"reset_measure")),MakeString("inspect.cl:423"))
  
  Core.F_attach_method(C_self_print.AddMethod(Signature(C_measure.Id(),C_void.Id()),1,MakeFunction1(E_self_print_measure_Reader,"self_print_measure_Reader")),MakeString("inspect.cl:424"))
  
  Core.F_attach_method(C_logMeasure.AddMethod(Signature(C_string.Id(),C_void.Id()),1,MakeFunction1(E_logMeasure_string,"logMeasure_string")),MakeString("inspect.cl:436"))
  
  Core.F_attach_method(C_addLog.AddMethod(Signature(C_integer.Id(),
    C_float.Id(),
    C_float.Id(),
    C_float.Id(),
    C_integer.Id(),
    C_void.Id()),1,MakeFunction5(E_addLog_integer,"addLog_integer")),MakeString("inspect.cl:443"))
  
  C_PRcount = MakeClass("PRcount",C_object,C_claire)
  Core.F_close_slot(C_PRcount.AddSlot(C_Reader_rtime,ToType(C_integer.Id()),MakeInteger(0).Id()))
  Core.F_close_slot(C_PRcount.AddSlot(C_Reader_rdepth,ToType(C_integer.Id()),MakeInteger(0).Id()))
  Core.F_close_slot(C_PRcount.AddSlot(C_Reader_rnum,ToType(C_integer.Id()),MakeInteger(0).Id()))
  Core.F_close_slot(C_PRcount.AddSlot(C_Reader_rloop,ToType(C_integer.Id()),MakeInteger(0).Id()))
  Core.F_close_slot(C_PRcount.AddSlot(C_Reader_rstart,ToType(C_integer.Id()),MakeInteger(0).Id()))
  
  Core.F_attach_method(C_PRget.AddMethod(Signature(C_property.Id(),C_PRcount.Id()),1,MakeFunction1(E_PRget_property,"PRget_property")),MakeString("inspect.cl:475"))
  
  Core.F_attach_method(C_PRlook.AddMethod(Signature(C_property.Id(),C_any.Id()),1,MakeFunction1(E_PRlook_property2,"PRlook_property2")),MakeString("inspect.cl:478"))
  
  Core.F_attach_method(C_PRshow.AddMethod(Signature(C_property.Id(),C_void.Id()),1,MakeFunction1(E_PRshow_property,"PRshow_property")),MakeString("inspect.cl:483"))
  
  Core.F_attach_method(C_PRtime.AddMethod(Signature(C_property.Id(),C_integer.Id()),0,MakeFunction1(E_PRtime_property,"PRtime_property")),MakeString("inspect.cl:488"))
  
  Core.F_attach_method(C_PRcounter.AddMethod(Signature(C_property.Id(),C_integer.Id()),0,MakeFunction1(E_PRcounter_property,"PRcounter_property")),MakeString("inspect.cl:492"))
  
  Core.F_attach_method(C_PRshow.AddMethod(Signature(C_void.Id(),C_void.Id()),1,MakeFunction1(E_PRshow_void,"PRshow_void")),MakeString("inspect.cl:510"))
  
  /* object!:Reader/"PRdependent" ->Reader*/C_Reader_PRdependent = ToTable(new(ClaireTable).IsNamed(C_table,MakeSymbol("PRdependent",It)))
  C_Reader_PRdependent.Multivalued_ask = CTRUE
  C_Reader_PRdependent.Range = Core.F_nth_class1(C_set,ToType(C_property.Id()))
  C_Reader_PRdependent.Params = C_any.Id()
  C_Reader_PRdependent.Domain = ToType(C_property.Id())
  C_Reader_PRdependent.GraphInit()
  C_Reader_PRdependent.Default = ToType(C_property.Id()).EmptySet().Id()
  
  /* object!:Reader/"PRdependentOf" ->Reader*/C_Reader_PRdependentOf = ToTable(new(ClaireTable).IsNamed(C_table,MakeSymbol("PRdependentOf",It)))
  C_Reader_PRdependentOf.Multivalued_ask = CTRUE
  C_Reader_PRdependentOf.Range = Core.F_nth_class1(C_set,ToType(C_property.Id()))
  C_Reader_PRdependentOf.Params = C_any.Id()
  C_Reader_PRdependentOf.Domain = ToType(C_property.Id())
  C_Reader_PRdependentOf.GraphInit()
  C_Reader_PRdependentOf.Default = ToType(C_property.Id()).EmptySet().Id()
  
  Core.F_attach_method(C_Reader_dependents.AddMethod(Signature(C_method.Id(),Core.F_nth_class1(C_set,ToType(C_property.Id())).Id()),1,MakeFunction1(E_dependents_method,"dependents_method")),MakeString("inspect.cl:518"))
  
  Core.F_attach_method(C_Reader_dependents.AddMethod(Signature(C_any.Id(),C_any.Id()),1,MakeFunction1(E_dependents_any,"dependents_any")),MakeString("inspect.cl:532"))
  
  { 
    /*PROTECT mClaire/update @ property(inverse,Reader/PRdependent,8,relation,Reader/PRdependentOf) */
    var expr EID 
    expr = Core.F_update_property(C_inverse,
      ToObject(C_Reader_PRdependent.Id()),
      8,
      C_relation,
      C_Reader_PRdependentOf.Id())
    ErrorCheck(expr)} 
  
  Core.F_attach_method(C_PRdepends.AddMethod(Signature(C_property.Id(),C_property.Id(),C_void.Id()),1,MakeFunction2(E_PRdepends_property,"PRdepends_property")),MakeString("inspect.cl:537"))
  
  } 

