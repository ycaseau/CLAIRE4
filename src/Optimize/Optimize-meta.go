/***** CLAIRE Compilation of module Optimize.cl 
         [version 4.1 / safety 5] Saturday 09-23-2023 07:22:33 *****/

package Optimize
import (_ "fmt"
	"unsafe"
	. "Kernel"
	"Core"
	"Language"
	"Reader"
)

//-------- dumb function to prevent import errors --------
func import_g0331() { 
    _ = Core.It
    _ = Language.It
    _ = Reader.It
    } 
  
  
// class file for EID in module Optimize 
type ClaireEID struct { 
   ClaireObject
   } 

// automatic cast function
func To_ClaireEID(x *ClaireAny) *ClaireEID {return (*ClaireEID)(unsafe.Pointer(x))}

// automatic constructor function
func Make_ClaireEID() *ClaireEID { 
  var o *ClaireEID = new(ClaireEID)
  o.Isa = C_EID
  return o 
  } 

// class file for Optimize/meta_OPT in module Optimize 
type OptimizeMetaOPT struct { 
   ClaireThing
   Instructions *ClaireList
  Objects *ClaireList
  Properties *ClaireSet
  Functions *ClaireList
  NeedToClose *ClaireSet
  NeedModules *ClaireSet
  LegalModules *ClaireSet
  Ignore *ClaireSet
  ToRemove *ClaireSet
  Outfile *ClairePort
  MaxVars int
  LoopIndex int
  Level int
  InMethod *ClaireAny
  Profile_ask *ClaireBoolean
  Cfile *ClaireAny
  Recompute *ClaireBoolean
  Unsure *ClaireList
  Knowns *ClaireSet
  SimpleOperations *ClaireSet
  NonIdentifiableSet *ClaireSet
  UseStringUpdate *ClaireBoolean
  } 

// automatic cast function
func ToOptimizeMetaOPT(x *ClaireAny) *OptimizeMetaOPT {return (*OptimizeMetaOPT)(unsafe.Pointer(x))}

// class file for Optimize/meta_compiler in module Optimize 
type OptimizeMetaCompiler struct { 
   ClaireThing
   External *ClaireString
  Source *ClaireString
  Debug_ask *ClaireList
  Version *ClaireAny
  Active_ask *ClaireBoolean
  Safety int
  Env *ClaireString
  Libraries *ClaireList
  Inline_ask *ClaireBoolean
  Loading_ask *ClaireBoolean
  Overflow_ask *ClaireBoolean
  Optimize_ask *ClaireBoolean
  NLoc int
  NWarnings int
  NNotes int
  NDynamic int
  NMethods int
  NMetheids int
  } 

// automatic cast function
func ToOptimizeMetaCompiler(x *ClaireAny) *OptimizeMetaCompiler {return (*OptimizeMetaCompiler)(unsafe.Pointer(x))}

// class file for Compile/producer in module Optimize 
type CompileProducer struct { 
   ClaireThing
   CurrentFile *ClaireString
  } 

// automatic cast function
func ToCompileProducer(x *ClaireAny) *CompileProducer {return (*CompileProducer)(unsafe.Pointer(x))}

// automatic constructor function
func MakeCompileProducer(name *ClaireSymbol,current_file *ClaireString) *CompileProducer { 
  var o *CompileProducer = new(CompileProducer)
  o.Isa = C_Compile_producer
  o.Name = name
  o.CurrentFile = current_file
  return o 
  } 

// class file for Compile/C_cast in module Optimize 
type Compile_CCast struct { 
   Language.OptimizedInstruction
   Arg *ClaireAny
  SetArg *ClaireClass
  } 

// automatic cast function
func To_Compile_CCast(x *ClaireAny) *Compile_CCast {return (*Compile_CCast)(unsafe.Pointer(x))}

// automatic constructor function
func Make_Compile_CCast(arg *ClaireAny,set_arg *ClaireClass) *Compile_CCast { 
  var o *Compile_CCast = new(Compile_CCast)
  o.Isa = C_Compile_C_cast
  o.Arg = arg
  o.SetArg = set_arg
  return o 
  } 

// class file for Optimize/Pattern in module Optimize 
type ClairePattern struct { 
   ClaireTypeExpression
   Selector *ClaireProperty
  Arg *ClaireList
  } 

// automatic cast function
func To_ClairePattern(x *ClaireAny) *ClairePattern {return (*ClairePattern)(unsafe.Pointer(x))}

// automatic constructor function
func Make_ClairePattern(selector *ClaireProperty,arg *ClaireList) *ClairePattern { 
  var o *ClairePattern = new(ClairePattern)
  o.Isa = C_Optimize_Pattern
  o.Selector = selector
  o.Arg = arg
  return o 
  } 

// class file for Optimize/Call_function2 in module Optimize 
type Optimize_CallFunction2 struct { 
   Language.OptimizedInstruction
   Arg *ClaireFunction
  Args *ClaireList
  } 

// automatic cast function
func To_Optimize_CallFunction2(x *ClaireAny) *Optimize_CallFunction2 {return (*Optimize_CallFunction2)(unsafe.Pointer(x))}

// automatic constructor function
func Make_Optimize_CallFunction2(arg *ClaireFunction,args *ClaireList) *Optimize_CallFunction2 { 
  var o *Optimize_CallFunction2 = new(Optimize_CallFunction2)
  o.Isa = C_Optimize_Call_function2
  o.Arg = arg
  o.Args = args
  return o 
  } 

var C_Compile_srange *Core.GlobalVariable
var C_Compile_status *Core.GlobalVariable
var C_Compile_tmatch_ask *Core.GlobalVariable
var C_Compile_pname *Core.GlobalVariable
var C_EID *ClaireClass /*obj*/
var C_Optimize_meta_OPT *ClaireClass /*obj*/
var C_Optimize_meta_compiler *ClaireClass /*obj*/
var C_Compile_producer *ClaireClass /*obj*/
var C_PRODUCER *Core.GlobalVariable
var C_Compile_FileOrigin *ClaireTable /*obj*/
var C_claire_modules *Core.GlobalVariable
var C_compiler *OptimizeMetaCompiler /*obj*/
var C_c_type *ClaireProperty /*obj*/
var C_c_code *ClaireProperty /*obj*/
var C_Compile_get_index *ClaireProperty /*obj*/
var C_Compile_get_indexed *ClaireProperty /*obj*/
var C_Compile_make_c_function *ClaireProperty /*obj*/
var C_Compile_bool_exp *ClaireProperty /*obj*/
var C_Compile_c_statement *ClaireProperty /*obj*/
var C_c_interface *ClaireProperty /*obj*/
var C_Compile_c_sort *ClaireProperty /*obj*/
var C_Compile_g_throw *ClaireProperty /*obj*/
var C_Compile_can_throw_ask *ClaireProperty /*obj*/
var C_Compile_can_throw_I *ClaireProperty /*obj*/
var C_Compile_designated_ask *ClaireProperty /*obj*/
var C_Compile_sort_equal *ClaireProperty /*obj*/
var C_Compile_psort *ClaireProperty /*obj*/
var C_Compile_osort *ClaireProperty /*obj*/
var C_Compile_compile_lambda *ClaireProperty /*obj*/
var C_Compile_need_protect *ClaireProperty /*obj*/
var C_Optimize_member_code *ClaireProperty /*obj*/
var C_Compile_c_strict_code *ClaireProperty /*obj*/
var C_Compile_c_strict_check *ClaireProperty /*obj*/
var C_Compile_object_I *ClaireProperty /*obj*/
var C_Compile_anyObject_I *ClaireProperty /*obj*/
var C_Compile_Cerror *ClaireProperty /*obj*/
var C_Compile_self_code *ClaireProperty /*obj*/
var C_Compile_get_module *ClaireProperty /*obj*/
var C_Compile_function_name *ClaireProperty /*obj*/
var C_Compile_m_unsafe *Core.GlobalVariable
var C_Compile_m_member *Core.GlobalVariable
var C_Compile_warn *ClaireProperty /*obj*/
var C_Compile_nativeVar_ask *ClaireProperty /*obj*/
var C_Compile_C_cast *ClaireClass /*obj*/
var C_Optimize_Pattern *ClaireClass /*obj*/
var C_OPT *OptimizeMetaOPT /*obj*/
var C_Compile_NoErrorOptimize *Core.GlobalVariable
var C_Compile_ForceThrow *Core.GlobalVariable
var C_Compile_ForceNotThrow *Core.GlobalVariable
var C_Optimize__startimes_integer_star *Core.GlobalVariable
var C_Optimize__stardiv_integer_star *Core.GlobalVariable
var C_Optimize__stardiv_float_star *Core.GlobalVariable
var C_Optimize__starmod_integer_star *Core.GlobalVariable
var C_DSHOW *Core.GlobalVariable
var C_DTHROW *Core.GlobalVariable
var C_Optimize_ambiguous *ClaireKeyword /*obj*/
var C_Compile_functional_I *ClaireProperty /*obj*/
var C_Optimize_Call_function2 *ClaireClass /*obj*/
var C_PENIBLE *Core.GlobalVariable
var C_Compile__starname_star *Core.GlobalVariable
var C_Optimize__staradd_method2_star *Core.GlobalVariable
var C_Optimize__equalsig_ask *ClaireOperation /*obj*/
var C_Optimize_selector_register *ClaireProperty // Optimize/"selector_register"
var C_Optimize_stable_ask *ClaireProperty // Optimize/"stable?"
var C_Compile_c_gc_ask *ClaireProperty // Compile/"c_gc?"
var C_Compile_notice *ClaireProperty // Compile/"notice"
var C_Optimize_c_warn *ClaireProperty // Optimize/"c_warn"
var C_Optimize_open_message *ClaireProperty // Optimize/"open_message"
var C_Compile_stupid_t *ClaireProperty // Compile/"stupid_t"
var C_Optimize_extended_ask *ClaireProperty // Optimize/"extended?"
var C_Optimize_extends *ClaireProperty // Optimize/"extends"
var C_Optimize_sort_abstract_I *ClaireProperty // Optimize/"sort_abstract!"
var C_Optimize_notOpt *ClaireProperty // Optimize/"notOpt"
var C_Compile_update_write_ask *ClaireProperty // Compile/"update_write?"
var C_Optimize_sort_abstract_ask *ClaireProperty // Optimize/"sort_abstract?"
var C_Optimize_pmember *ClaireProperty // Optimize/"pmember"
var C_Optimize_enumerate_code *ClaireProperty // Optimize/"enumerate_code"
var C_Optimize_c_code_method *ClaireProperty // Optimize/"c_code_method"
var C_Optimize_range_infers_for *ClaireProperty // Optimize/"range_infers_for"
var C_Optimize_range_infers *ClaireProperty // Optimize/"range_infers"
var C_Optimize_range_sets *ClaireProperty // Optimize/"range_sets"
var C_Compile_return_type *ClaireProperty // Compile/"return_type"
var C_Compile_Variable_I *ClaireProperty // Compile/"Variable!"
var C_Compile_identifiable_ask *ClaireProperty // Compile/"identifiable?"
var C_Optimize_c_inline *ClaireProperty // Optimize/"c_inline"
var C_Optimize_bound_variables *ClaireProperty // Optimize/"bound_variables"
var C_Optimize_c_substitution *ClaireProperty // Optimize/"c_substitution"
var C_Optimize_c_inline_arg_ask *ClaireProperty // Optimize/"c_inline_arg?"
var C_Optimize_restriction_I *ClaireProperty // Optimize/"restriction!"
var C_Optimize_c_inline_ask *ClaireProperty // Optimize/"c_inline?"
var C_Optimize_c_boolean *ClaireProperty // Optimize/"c_boolean"
var C_findr *ClaireProperty // claire/"findr"
var C_dmatch_ask *ClaireProperty // claire/"dmatch?"
var C_Optimize_use_range *ClaireProperty // Optimize/"use_range"
var C_Optimize_inline_optimize_ask *ClaireProperty // Optimize/"inline_optimize?"
var C_Optimize_daccess *ClaireProperty // Optimize/"daccess"
var C_Optimize_c_code_write *ClaireProperty // Optimize/"c_code_write"
var C_Optimize_c_code_hold *ClaireProperty // Optimize/"c_code_hold"
var C_Optimize_Produce_erase *ClaireProperty // Optimize/"Produce_erase"
var C_Optimize_c_code_add *ClaireProperty // Optimize/"c_code_add"
var C_Optimize_c_code_add_bag *ClaireProperty // Optimize/"c_code_add_bag"
var C_Optimize_c_code_delete *ClaireProperty // Optimize/"c_code_delete"
var C_Optimize_c_code_table *ClaireProperty // Optimize/"c_code_table"
var C_Optimize_c_code_array *ClaireProperty // Optimize/"c_code_array"
var C_Optimize_c_code_nth *ClaireProperty // Optimize/"c_code_nth"
var C_Optimize_c_code_belong *ClaireProperty // Optimize/"c_code_belong"
var C_Optimize_c_code_not *ClaireProperty // Optimize/"c_code_not"
var C_Optimize_Update_ask *ClaireProperty // Optimize/"Update?"
var C_Optimize_Call_method_I *ClaireProperty // Optimize/"Call_method!"
var C_Optimize_nth_type_check *ClaireProperty // Optimize/"nth_type_check"
var C_Optimize_extendedTest_ask *ClaireProperty // Optimize/"extendedTest?"
var C_safety *ClaireProperty // claire/"safety"
var C_home *ClaireProperty // claire/"home"
var C_Compile_instructions *ClaireProperty // Compile/"instructions"
var C_Compile_objects *ClaireProperty // Compile/"objects"
var C_Compile_properties *ClaireProperty // Compile/"properties"
var C_Compile_functions *ClaireProperty // Compile/"functions"
var C_Compile_need_to_close *ClaireProperty // Compile/"need_to_close"
var C_Compile_need_modules *ClaireProperty // Compile/"need_modules"
var C_Compile_legal_modules *ClaireProperty // Compile/"legal_modules"
var C_Compile_ignore *ClaireProperty // Compile/"ignore"
var C_Compile_to_remove *ClaireProperty // Compile/"to_remove"
var C_Compile_outfile *ClaireProperty // Compile/"outfile"
var C_Compile_max_vars *ClaireProperty // Compile/"max_vars"
var C_Compile_loop_index *ClaireProperty // Compile/"loop_index"
var C_Compile_level *ClaireProperty // Compile/"level"
var C_Compile_in_method *ClaireProperty // Compile/"in_method"
var C_Compile_profile_ask *ClaireProperty // Compile/"profile?"
var C_Compile_cfile *ClaireProperty // Compile/"cfile"
var C_Optimize_recompute *ClaireProperty // Optimize/"recompute"
var C_Optimize_unsure *ClaireProperty // Optimize/"unsure"
var C_Optimize_knowns *ClaireProperty // Optimize/"knowns"
var C_Compile_simple_operations *ClaireProperty // Compile/"simple_operations"
var C_Compile_non_identifiable_set *ClaireProperty // Compile/"non_identifiable_set"
var C_Compile_use_string_update *ClaireProperty // Compile/"use_string_update"
var C_debug_ask *ClaireProperty // claire/"debug?"
var C_active_ask *ClaireProperty // claire/"active?"
var C_libraries *ClaireProperty // claire/"libraries"
var C_loading_ask *ClaireProperty // claire/"loading?"
var C_overflow_ask *ClaireProperty // claire/"overflow?"
var C_optimize_ask *ClaireProperty // claire/"optimize?"
var C_n_loc *ClaireProperty // claire/"n_loc"
var C_n_warnings *ClaireProperty // claire/"n_warnings"
var C_n_notes *ClaireProperty // claire/"n_notes"
var C_n_dynamic *ClaireProperty // claire/"n_dynamic"
var C_n_methods *ClaireProperty // claire/"n_methods"
var C_n_metheids *ClaireProperty // claire/"n_metheids"
var C_Compile_current_file *ClaireProperty // Compile/"current_file"
var C_safe *ClaireProperty // claire/"safe"
var C_Optimize_ptype *ClaireProperty // Optimize/"ptype"
var C_Optimize_c_code_call *ClaireProperty // Optimize/"c_code_call"
var C_Optimize_c_register *ClaireProperty // Optimize/"c_register"
var C_Optimize_get_sort *ClaireProperty // Optimize/"get_sort"
var C_Optimize_selector_psort *ClaireProperty // Optimize/"selector_psort"
var C_Optimize_c_srange *ClaireProperty // Optimize/"c_srange"
var C_Optimize_g_throw1 *ClaireProperty // Optimize/"g_throw1"
var C_Optimize_case_branch *ClaireProperty // Optimize/"case_branch"
var C_Optimize_case_substitution *ClaireProperty // Optimize/"case_substitution"
var C_Optimize_infers_from *ClaireProperty // Optimize/"infers_from"
var C_Optimize_Iterate_I *ClaireProperty // Optimize/"Iterate!"
var C_Optimize_c_code_multiple *ClaireProperty // Optimize/"c_code_multiple"
var C_Optimize_c_code_select *ClaireProperty // Optimize/"c_code_select"
var C_Optimize_inner_select *ClaireProperty // Optimize/"inner_select"
var C_Optimize_total_ask *ClaireProperty // Optimize/"total?"
var C_Optimize_analyze_I *ClaireProperty // Optimize/"analyze!"
var C_Optimize_extract_signature_I *ClaireProperty // Optimize/"extract_signature!"
var C_Optimize_sort_pattern_ask *ClaireProperty // Optimize/"sort_pattern?"
var C_Optimize_sort_code *ClaireProperty // Optimize/"sort_code"
var C_Optimize_add_method_I *ClaireProperty // Optimize/"add_method!"
var C_Optimize_type_extension *ClaireProperty // Optimize/"type_extension"
var C_Compile_compute_if_write_inverse *ClaireProperty // Compile/"compute_if_write_inverse"
var C_Optimize_Produce_put *ClaireProperty // Optimize/"Produce_put"
var C_Optimize_Produce_get *ClaireProperty // Optimize/"Produce_get"
var C_Optimize_Produce_remove *ClaireProperty // Optimize/"Produce_remove"
var C_Compile_compute_set_write *ClaireProperty // Compile/"compute_set_write"
var C_Optimize_Tighten *ClaireProperty // Optimize/"Tighten"
var C_Compile_Tighten_I *ClaireProperty // Compile/"Tighten!"
var C_Compile_lexical_num *ClaireProperty // Compile/"lexical_num"
var C_Optimize_compile_if_write *ClaireProperty // Optimize/"compile_if_write"
var C_Optimize_demon_function *ClaireProperty // Optimize/"demon_function"
var C_Optimize_compileEventMethod *ClaireProperty // Optimize/"compileEventMethod"
var C_Optimize_demon_substitution *ClaireProperty // Optimize/"demon_substitution"
var C_Compile_Do_I *ClaireProperty // Compile/"Do!"
var C_env *ClaireProperty // claire/"env"
var C_Compile_can_throw_status *ClaireProperty // Compile/"can_throw_status"
var C_s_throw *ClaireProperty // claire/"s_throw"
var C_Optimize_legal_ask *ClaireProperty // Optimize/"legal?"
var It *ClaireModule
var C_Compile *ClaireModule 
// definition of the meta-model for module Optimize 
func MetaLoad() { 
  
  It = MakeModule("Optimize",C_Compile)
  It.Comment = MakeString("Compiled on Saturday 09-23-2023 07:22:33(v4.1), lines:3034, warnings:1,safety:5")
  ClEnv.Module_I = It
  
  // definition of the properties
  C_Optimize_selector_register = MakeProperty("selector_register",1,It)
  C_Optimize_stable_ask = MakeProperty("stable?",1,It)
  C_Compile_c_gc_ask = MakeProperty("c_gc?",1,C_Compile)
  C_Compile_notice = MakeProperty("notice",1,C_Compile)
  C_Optimize_c_warn = MakeProperty("c_warn",1,It)
  C_Optimize_open_message = MakeProperty("open_message",1,It)
  C_Compile_stupid_t = MakeProperty("stupid_t",1,C_Compile)
  C_Optimize_extended_ask = MakeProperty("extended?",1,It)
  C_Optimize_extends = MakeProperty("extends",1,It)
  C_Optimize_sort_abstract_I = MakeProperty("sort_abstract!",1,It)
  C_Optimize_notOpt = MakeProperty("notOpt",1,It)
  C_Compile_update_write_ask = MakeProperty("update_write?",1,C_Compile)
  C_Optimize_sort_abstract_ask = MakeProperty("sort_abstract?",1,It)
  C_Optimize_pmember = MakeProperty("pmember",1,It)
  C_Optimize_enumerate_code = MakeProperty("enumerate_code",1,It)
  C_Optimize_c_code_method = MakeProperty("c_code_method",1,It)
  C_Optimize_range_infers_for = MakeProperty("range_infers_for",1,It)
  C_Optimize_range_infers = MakeProperty("range_infers",1,It)
  C_Optimize_range_sets = MakeProperty("range_sets",1,It)
  C_Compile_return_type = MakeProperty("return_type",1,C_Compile)
  C_Compile_Variable_I = MakeProperty("Variable!",1,C_Compile)
  C_Compile_identifiable_ask = MakeProperty("identifiable?",1,C_Compile)
  C_Optimize_c_inline = MakeProperty("c_inline",1,It)
  C_Optimize_bound_variables = MakeProperty("bound_variables",1,It)
  C_Optimize_c_substitution = MakeProperty("c_substitution",1,It)
  C_Optimize_c_inline_arg_ask = MakeProperty("c_inline_arg?",1,It)
  C_Optimize_restriction_I = MakeProperty("restriction!",1,It)
  C_Optimize_c_inline_ask = MakeProperty("c_inline?",1,It)
  C_Optimize_c_boolean = MakeProperty("c_boolean",1,It)
  C_findr = MakeProperty("findr",1,C_claire)
  C_dmatch_ask = MakeProperty("dmatch?",1,C_claire)
  C_Optimize_use_range = MakeProperty("use_range",1,It)
  C_Optimize_inline_optimize_ask = MakeProperty("inline_optimize?",1,It)
  C_Optimize_daccess = MakeProperty("daccess",1,It)
  C_Optimize_c_code_write = MakeProperty("c_code_write",1,It)
  C_Optimize_c_code_hold = MakeProperty("c_code_hold",1,It)
  C_Optimize_Produce_erase = MakeProperty("Produce_erase",1,It)
  C_Optimize_c_code_add = MakeProperty("c_code_add",1,It)
  C_Optimize_c_code_add_bag = MakeProperty("c_code_add_bag",1,It)
  C_Optimize_c_code_delete = MakeProperty("c_code_delete",1,It)
  C_Optimize_c_code_table = MakeProperty("c_code_table",1,It)
  C_Optimize_c_code_array = MakeProperty("c_code_array",1,It)
  C_Optimize_c_code_nth = MakeProperty("c_code_nth",1,It)
  C_Optimize_c_code_belong = MakeProperty("c_code_belong",1,It)
  C_Optimize_c_code_not = MakeProperty("c_code_not",1,It)
  C_Optimize_Update_ask = MakeProperty("Update?",1,It)
  C_Optimize_Call_method_I = MakeProperty("Call_method!",1,It)
  C_Optimize_nth_type_check = MakeProperty("nth_type_check",1,It)
  C_Optimize_extendedTest_ask = MakeProperty("extendedTest?",1,It)
  C_safety = MakeProperty("safety",2,C_claire)
  C_home = MakeProperty("home",1,C_claire)
  C_Compile_instructions = MakeProperty("instructions",2,C_Compile)
  C_Compile_objects = MakeProperty("objects",2,C_Compile)
  C_Compile_properties = MakeProperty("properties",2,C_Compile)
  C_Compile_functions = MakeProperty("functions",2,C_Compile)
  C_Compile_need_to_close = MakeProperty("need_to_close",2,C_Compile)
  C_Compile_need_modules = MakeProperty("need_modules",2,C_Compile)
  C_Compile_legal_modules = MakeProperty("legal_modules",2,C_Compile)
  C_Compile_ignore = MakeProperty("ignore",2,C_Compile)
  C_Compile_to_remove = MakeProperty("to_remove",2,C_Compile)
  C_Compile_outfile = MakeProperty("outfile",2,C_Compile)
  C_Compile_max_vars = MakeProperty("max_vars",2,C_Compile)
  C_Compile_loop_index = MakeProperty("loop_index",2,C_Compile)
  C_Compile_level = MakeProperty("level",2,C_Compile)
  C_Compile_in_method = MakeProperty("in_method",2,C_Compile)
  C_Compile_profile_ask = MakeProperty("profile?",2,C_Compile)
  C_Compile_cfile = MakeProperty("cfile",2,C_Compile)
  C_Optimize_recompute = MakeProperty("recompute",2,It)
  C_Optimize_unsure = MakeProperty("unsure",2,It)
  C_Optimize_knowns = MakeProperty("knowns",2,It)
  C_Compile_simple_operations = MakeProperty("simple_operations",2,C_Compile)
  C_Compile_non_identifiable_set = MakeProperty("non_identifiable_set",2,C_Compile)
  C_Compile_use_string_update = MakeProperty("use_string_update",2,C_Compile)
  C_debug_ask = MakeProperty("debug?",2,C_claire)
  C_active_ask = MakeProperty("active?",2,C_claire)
  C_libraries = MakeProperty("libraries",2,C_claire)
  C_loading_ask = MakeProperty("loading?",2,C_claire)
  C_overflow_ask = MakeProperty("overflow?",2,C_claire)
  C_optimize_ask = MakeProperty("optimize?",2,C_claire)
  C_n_loc = MakeProperty("n_loc",2,C_claire)
  C_n_warnings = MakeProperty("n_warnings",2,C_claire)
  C_n_notes = MakeProperty("n_notes",2,C_claire)
  C_n_dynamic = MakeProperty("n_dynamic",2,C_claire)
  C_n_methods = MakeProperty("n_methods",2,C_claire)
  C_n_metheids = MakeProperty("n_metheids",2,C_claire)
  C_Compile_current_file = MakeProperty("current_file",2,C_Compile)
  C_safe = MakeProperty("safe",1,C_claire)
  C_Optimize_ptype = MakeProperty("ptype",1,It)
  C_Optimize_c_code_call = MakeProperty("c_code_call",1,It)
  C_Optimize_c_register = MakeProperty("c_register",1,It)
  C_Optimize_get_sort = MakeProperty("get_sort",1,It)
  C_Optimize_selector_psort = MakeProperty("selector_psort",1,It)
  C_Optimize_c_srange = MakeProperty("c_srange",1,It)
  C_Optimize_g_throw1 = MakeProperty("g_throw1",1,It)
  C_Optimize_case_branch = MakeProperty("case_branch",1,It)
  C_Optimize_case_substitution = MakeProperty("case_substitution",1,It)
  C_Optimize_infers_from = MakeProperty("infers_from",1,It)
  C_Optimize_Iterate_I = MakeProperty("Iterate!",1,It)
  C_Optimize_c_code_multiple = MakeProperty("c_code_multiple",1,It)
  C_Optimize_c_code_select = MakeProperty("c_code_select",1,It)
  C_Optimize_inner_select = MakeProperty("inner_select",1,It)
  C_Optimize_total_ask = MakeProperty("total?",1,It)
  C_Optimize_analyze_I = MakeProperty("analyze!",1,It)
  C_Optimize_extract_signature_I = MakeProperty("extract_signature!",1,It)
  C_Optimize_sort_pattern_ask = MakeProperty("sort_pattern?",1,It)
  C_Optimize_sort_code = MakeProperty("sort_code",1,It)
  C_Optimize_add_method_I = MakeProperty("add_method!",1,It)
  C_Optimize_type_extension = MakeProperty("type_extension",1,It)
  C_Compile_compute_if_write_inverse = MakeProperty("compute_if_write_inverse",1,C_Compile)
  C_Optimize_Produce_put = MakeProperty("Produce_put",1,It)
  C_Optimize_Produce_get = MakeProperty("Produce_get",1,It)
  C_Optimize_Produce_remove = MakeProperty("Produce_remove",1,It)
  C_Compile_compute_set_write = MakeProperty("compute_set_write",1,C_Compile)
  C_Optimize_Tighten = MakeProperty("Tighten",1,It)
  C_Compile_Tighten_I = MakeProperty("Tighten!",1,C_Compile)
  C_Compile_lexical_num = MakeProperty("lexical_num",1,C_Compile)
  C_Optimize_compile_if_write = MakeProperty("compile_if_write",1,It)
  C_Optimize_demon_function = MakeProperty("demon_function",1,It)
  C_Optimize_compileEventMethod = MakeProperty("compileEventMethod",1,It)
  C_Optimize_demon_substitution = MakeProperty("demon_substitution",1,It)
  C_Compile_Do_I = MakeProperty("Do!",1,C_Compile)
  C_env = MakeProperty("env",2,C_claire)
  C_Compile_can_throw_status = MakeProperty("can_throw_status",1,C_Compile)
  C_s_throw = MakeProperty("s_throw",1,C_claire)
  C_Optimize_legal_ask = MakeProperty("legal?",1,It)
  
  // instructions from module sources
  { 
    var expr EID
    { var _CL_obj *Core.GlobalVariable
      C_Compile_srange = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("srange",C_Compile)))
      
      _CL_obj = C_Compile_srange
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = C_mClaire_srange.Id()
      expr = Core.F_close_global_variable(_CL_obj)
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID
    { var _CL_obj *Core.GlobalVariable
      C_Compile_status = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("status",C_Compile)))
      
      _CL_obj = C_Compile_status
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = C_mClaire_status.Id()
      expr = Core.F_close_global_variable(_CL_obj)
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID
    { var _CL_obj *Core.GlobalVariable
      C_Compile_tmatch_ask = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("tmatch?",C_Compile)))
      
      _CL_obj = C_Compile_tmatch_ask
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = Core.C_Core_tmatch_ask.Id()
      expr = Core.F_close_global_variable(_CL_obj)
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID
    { var _CL_obj *Core.GlobalVariable
      C_Compile_pname = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("pname",C_Compile)))
      
      _CL_obj = C_Compile_pname
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = C_mClaire_pname.Id()
      expr = Core.F_close_global_variable(_CL_obj)
      } 
    ErrorCheck(expr)} 
  
  _ = Core.F_attach_method(C_home.AddMethod(Signature(C_void.Id(),C_string.Id()),0,MakeFunction1(E_home_void,"home_void")),MakeString("osystem.cl:42"))
  
  C_EID = NewClass("EID",C_object,C_claire)
  
  C_Optimize_meta_OPT = NewClass("meta_OPT",C_thing,It)
  Core.F_close_slot(C_Optimize_meta_OPT.AddSlot(C_Compile_instructions,ToType(C_list.Id()),ToType(C_any.Id()).EmptyList().Id()))
  Core.F_close_slot(C_Optimize_meta_OPT.AddSlot(C_Compile_objects,ToType(C_list.Id()),ToType(C_any.Id()).EmptyList().Id()))
  Core.F_close_slot(C_Optimize_meta_OPT.AddSlot(C_Compile_properties,Core.F_nth_class2(C_set,MakeList(ToType(C_any.Id()),C_of.Id()),MakeConstantList(MakeConstantSet(C_property.Id()).Id())),ToType(C_property.Id()).EmptySet().Id()))
  Core.F_close_slot(C_Optimize_meta_OPT.AddSlot(C_Compile_functions,ToType(C_list.Id()),ToType(C_any.Id()).EmptyList().Id()))
  Core.F_close_slot(C_Optimize_meta_OPT.AddSlot(C_Compile_need_to_close,ToType(C_set.Id()),ToType(C_any.Id()).EmptySet().Id()))
  Core.F_close_slot(C_Optimize_meta_OPT.AddSlot(C_Compile_need_modules,ToType(C_set.Id()),ToType(C_any.Id()).EmptySet().Id()))
  Core.F_close_slot(C_Optimize_meta_OPT.AddSlot(C_Compile_legal_modules,ToType(C_set.Id()),ToType(C_any.Id()).EmptySet().Id()))
  Core.F_close_slot(C_Optimize_meta_OPT.AddSlot(C_Compile_ignore,ToType(C_set.Id()),ToType(C_any.Id()).EmptySet().Id()))
  Core.F_close_slot(C_Optimize_meta_OPT.AddSlot(C_Compile_to_remove,ToType(C_set.Id()),ToType(C_any.Id()).EmptySet().Id()))
  Core.F_close_slot(C_Optimize_meta_OPT.AddSlot(C_Compile_outfile,ToType(C_port.Id()),CNULL))
  Core.F_close_slot(C_Optimize_meta_OPT.AddSlot(C_Compile_max_vars,ToType(C_integer.Id()),MakeInteger(0).Id()))
  Core.F_close_slot(C_Optimize_meta_OPT.AddSlot(C_Compile_loop_index,ToType(C_integer.Id()),MakeInteger(0).Id()))
  Core.F_close_slot(C_Optimize_meta_OPT.AddSlot(C_Compile_level,ToType(C_integer.Id()),MakeInteger(0).Id()))
  Core.F_close_slot(C_Optimize_meta_OPT.AddSlot(C_Compile_in_method,ToType(C_any.Id()),CNULL))
  Core.F_close_slot(C_Optimize_meta_OPT.AddSlot(C_Compile_profile_ask,ToType(C_boolean.Id()),CFALSE.Id()))
  Core.F_close_slot(C_Optimize_meta_OPT.AddSlot(C_Compile_cfile,ToType(C_any.Id()),CFALSE.Id()))
  Core.F_close_slot(C_Optimize_meta_OPT.AddSlot(C_Optimize_recompute,ToType(C_boolean.Id()),CTRUE.Id()))
  Core.F_close_slot(C_Optimize_meta_OPT.AddSlot(C_Optimize_unsure,ToType(C_list.Id()),CNIL.Id()))
  Core.F_close_slot(C_Optimize_meta_OPT.AddSlot(C_Optimize_knowns,Core.F_nth_class2(C_set,MakeList(ToType(C_any.Id()),C_of.Id()),MakeConstantList(MakeConstantSet(C_relation.Id()).Id())),ToType(C_relation.Id()).EmptySet().Id()))
  Core.F_close_slot(C_Optimize_meta_OPT.AddSlot(C_Compile_simple_operations,Core.F_nth_class2(C_set,MakeList(ToType(C_any.Id()),C_of.Id()),MakeConstantList(MakeConstantSet(C_property.Id()).Id())),ToType(C_property.Id()).EmptySet().Id()))
  Core.F_close_slot(C_Optimize_meta_OPT.AddSlot(C_Compile_non_identifiable_set,Core.F_nth_class2(C_set,MakeList(ToType(C_any.Id()),C_of.Id()),MakeConstantList(MakeConstantSet(C_class.Id()).Id())),ToType(C_class.Id()).EmptySet().Id()))
  Core.F_close_slot(C_Optimize_meta_OPT.AddSlot(C_Compile_use_string_update,ToType(C_boolean.Id()),CFALSE.Id()))
  
  C_Optimize_meta_compiler = NewClass("meta_compiler",C_thing,It)
  Core.F_close_slot(C_Optimize_meta_compiler.AddSlot(C_external,ToType(C_string.Id()),CNULL))
  Core.F_close_slot(C_Optimize_meta_compiler.AddSlot(C_source,ToType(C_string.Id()),CNULL))
  Core.F_close_slot(C_Optimize_meta_compiler.AddSlot(C_debug_ask,Core.F_nth_class2(C_list,MakeList(ToType(C_any.Id()),C_of.Id()),MakeConstantList(MakeConstantSet(C_module.Id()).Id())),ToType(C_module.Id()).EmptyList().Id()))
  Core.F_close_slot(C_Optimize_meta_compiler.AddSlot(C_version,ToType(C_any.Id()),CNULL))
  Core.F_close_slot(C_Optimize_meta_compiler.AddSlot(C_active_ask,ToType(C_boolean.Id()),CFALSE.Id()))
  Core.F_close_slot(C_Optimize_meta_compiler.AddSlot(C_safety,ToType(C_integer.Id()),MakeInteger(1).Id()))
  Core.F_close_slot(C_Optimize_meta_compiler.AddSlot(C_env,ToType(C_string.Id()),CNULL))
  Core.F_close_slot(C_Optimize_meta_compiler.AddSlot(C_libraries,Core.F_nth_class2(C_list,MakeList(ToType(C_any.Id()),C_of.Id()),MakeConstantList(MakeConstantSet(C_string.Id()).Id())),ToType(C_string.Id()).EmptyList().Id()))
  Core.F_close_slot(C_Optimize_meta_compiler.AddSlot(C_inline_ask,ToType(C_boolean.Id()),CFALSE.Id()))
  Core.F_close_slot(C_Optimize_meta_compiler.AddSlot(C_loading_ask,ToType(C_boolean.Id()),CFALSE.Id()))
  Core.F_close_slot(C_Optimize_meta_compiler.AddSlot(C_overflow_ask,ToType(C_boolean.Id()),CFALSE.Id()))
  Core.F_close_slot(C_Optimize_meta_compiler.AddSlot(C_optimize_ask,ToType(C_boolean.Id()),CFALSE.Id()))
  Core.F_close_slot(C_Optimize_meta_compiler.AddSlot(C_n_loc,ToType(C_integer.Id()),MakeInteger(0).Id()))
  Core.F_close_slot(C_Optimize_meta_compiler.AddSlot(C_n_warnings,ToType(C_integer.Id()),MakeInteger(0).Id()))
  Core.F_close_slot(C_Optimize_meta_compiler.AddSlot(C_n_notes,ToType(C_integer.Id()),MakeInteger(0).Id()))
  Core.F_close_slot(C_Optimize_meta_compiler.AddSlot(C_n_dynamic,ToType(C_integer.Id()),MakeInteger(0).Id()))
  Core.F_close_slot(C_Optimize_meta_compiler.AddSlot(C_n_methods,ToType(C_integer.Id()),MakeInteger(0).Id()))
  Core.F_close_slot(C_Optimize_meta_compiler.AddSlot(C_n_metheids,ToType(C_integer.Id()),MakeInteger(0).Id()))
  
  C_Compile_producer = NewClass("producer",C_thing,C_Compile)
  Core.F_close_slot(C_Compile_producer.AddSlot(C_Compile_current_file,ToType(C_string.Id()),MakeString("").Id()))
  
  { 
    var expr EID
    { var _CL_obj *Core.GlobalVariable
      C_PRODUCER = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("PRODUCER",C_claire)))
      
      _CL_obj = C_PRODUCER
      _CL_obj.Range = ToType(C_Compile_producer.Id())
      _CL_obj.Value = CNULL
      expr = Core.F_close_global_variable(_CL_obj)
      } 
    ErrorCheck(expr)} 
  
  C_Compile_FileOrigin = ToTable(new(ClaireTable).IsNamed(C_table,MakeSymbol("FileOrigin",C_Compile)))
  C_Compile_FileOrigin.Range = ToType(C_string.Id())
  C_Compile_FileOrigin.Params = C_any.Id()
  C_Compile_FileOrigin.Domain = ToType(C_method.Id())
  C_Compile_FileOrigin.GraphInit()
  
  C_Compile_FileOrigin.Default = MakeString("").Id()
  
  { 
    var expr EID
    { var _CL_obj *Core.GlobalVariable
      C_claire_modules = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("claire_modules",C_claire)))
      
      _CL_obj = C_claire_modules
      _CL_obj.Range = ToType(C_list.Id())
      { 
        var va_arg1 *Core.GlobalVariable
        var va_arg2 *ClaireAny
        va_arg1 = _CL_obj
        { 
          var v_list4 *ClaireList
          var x *ClaireString
          var v_local4 *ClaireAny
          v_list4 = MakeConstantList(MakeString("Kernel").Id(),
            MakeString("Core").Id(),
            MakeString("Language").Id(),
            MakeString("Reader").Id())
          va_arg2 = CreateList(ToType(CEMPTY.Id()),v_list4.Length()).Id()
          for CLcount := 0; CLcount < v_list4.Length(); CLcount++{ 
            x = ToString(v_list4.At(CLcount))
            v_local4 = F_value_string(x)
            ToList(va_arg2).PutAt(CLcount,v_local4)
            } 
          } 
        va_arg1.Value = va_arg2
        } 
      expr = Core.F_close_global_variable(_CL_obj)
      } 
    ErrorCheck(expr)} 
  
  C_compiler = ToOptimizeMetaCompiler(new(OptimizeMetaCompiler).IsNamed(C_Optimize_meta_compiler,MakeSymbol("compiler",C_claire)))
  C_compiler.External = MakeString("go")
  C_compiler.Env = MakeString("MacOS")
  C_compiler.Version = MakeFloat(1).Id()
  C_compiler.Source = MakeString("/Users/ycaseau/claire/v4.0/go/src")
  C_compiler.Libraries = MakeList(ToType(C_string.Id()),MakeString("Kernel").Id())
  
  
  C_c_type = MakeProperty("c_type",3,C_claire)
  C_c_type.Open = 3
  C_c_type.Range = ToType(C_type.Id())
  
  
  C_c_code = MakeProperty("c_code",3,C_claire)
  C_c_code.Open = 3
  
  
  C_Compile_get_index = MakeProperty("get_index",3,C_Compile)
  C_Compile_get_index.Range = ToType(C_integer.Id())
  C_Compile_get_index.Open = 3
  
  
  C_Compile_get_indexed = MakeProperty("get_indexed",3,C_Compile)
  C_Compile_get_indexed.Range = ToType(C_list.Id())
  C_Compile_get_indexed.Open = 3
  
  
  C_Compile_make_c_function = MakeProperty("make_c_function",3,C_Compile)
  C_Compile_make_c_function.Open = 3
  
  
  C_Compile_bool_exp = MakeProperty("bool_exp",3,C_Compile)
  C_Compile_bool_exp.Open = 3
  
  
  C_Compile_c_statement = MakeProperty("c_statement",2,C_Compile)
  
  
  C_c_interface = MakeProperty("c_interface",3,C_claire)
  C_c_interface.Open = 3
  
  
  C_Compile_c_sort = MakeProperty("c_sort",3,C_Compile)
  C_Compile_c_sort.Open = 3
  
  
  C_Compile_g_throw = MakeProperty("g_throw",1,C_Compile)
  
  
  C_Compile_can_throw_ask = MakeProperty("can_throw?",1,C_Compile)
  
  
  C_Compile_can_throw_I = MakeProperty("can_throw!",1,C_Compile)
  
  
  C_Compile_designated_ask = MakeProperty("designated?",1,C_Compile)
  
  
  C_Compile_sort_equal = MakeProperty("sort=",1,C_Compile)
  
  
  C_Compile_psort = MakeProperty("psort",1,C_Compile)
  
  
  C_Compile_osort = MakeProperty("osort",1,C_Compile)
  
  
  C_Compile_compile_lambda = MakeProperty("compile_lambda",1,C_Compile)
  
  
  C_Compile_need_protect = MakeProperty("need_protect",2,C_Compile)
  
  
  C_Optimize_member_code = MakeProperty("member_code",1,It)
  
  
  C_Compile_c_strict_code = MakeProperty("c_strict_code",1,C_Compile)
  
  
  C_Compile_c_strict_check = MakeProperty("c_strict_check",3,C_Compile)
  C_Compile_c_strict_check.Open = 3
  
  
  C_Compile_object_I = MakeProperty("object!",2,C_Compile)
  
  
  C_Compile_anyObject_I = MakeProperty("anyObject!",2,C_Compile)
  C_Compile_anyObject_I.Range = ToType(C_object.Id())
  
  
  C_Compile_Cerror = MakeProperty("Cerror",1,C_Compile)
  
  
  C_Compile_self_code = MakeProperty("self_code",1,C_Compile)
  
  
  C_Compile_get_module = MakeProperty("get_module",1,C_Compile)
  
  
  C_Compile_function_name = MakeProperty("function_name",3,C_Compile)
  C_Compile_function_name.Open = 3
  
  
  { 
    var expr EID
    { var _CL_obj *Core.GlobalVariable
      C_Compile_m_unsafe = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("m_unsafe",C_Compile)))
      
      _CL_obj = C_Compile_m_unsafe
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = Core.F__at_property1(Core.C_unsafe,C_any).Id()
      expr = Core.F_close_global_variable(_CL_obj)
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID
    { var _CL_obj *Core.GlobalVariable
      C_Compile_m_member = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("m_member",C_Compile)))
      
      _CL_obj = C_Compile_m_member
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = Core.F__at_property2(ToProperty(C__Z.Id()),MakeConstantList(C_any.Id(),C_any.Id())).Id()
      expr = Core.F_close_global_variable(_CL_obj)
      } 
    ErrorCheck(expr)} 
  
  C_Compile_warn = MakeProperty("warn",1,C_Compile)
  
  
  C_Compile_nativeVar_ask = MakeProperty("nativeVar?",1,C_Compile)
  
  
  C_Compile_C_cast = NewClass("C_cast",Language.C_Optimized_instruction,C_Compile)
  Core.F_close_slot(C_Compile_C_cast.AddSlot(C_arg,ToType(C_any.Id()),CNULL))
  Core.F_close_slot(C_Compile_C_cast.AddSlot(Language.C_iClaire_set_arg,ToType(C_class.Id()),CNULL))
  
  C_Optimize_Pattern = NewClass("Pattern",C_type_expression,It)
  Core.F_close_slot(C_Optimize_Pattern.AddSlot(C_selector,ToType(C_property.Id()),CNULL))
  Core.F_close_slot(C_Optimize_Pattern.AddSlot(C_arg,ToType(C_list.Id()),ToType(C_any.Id()).EmptyList().Id()))
  
  C_OPT = ToOptimizeMetaOPT(new(OptimizeMetaOPT).IsNamed(C_Optimize_meta_OPT,MakeSymbol("OPT",C_claire)))
  C_OPT.Outfile = ToPort(Reader.C_stdin.Value)
  C_OPT.Ignore = MakeConstantSet(Core.C_mClaire_index_I.Id(),
    Core.C_mClaire_set_index.Id(),
    C_Compile_object_I.Id(),
    Core.C_mClaire_base_I.Id(),
    Core.C_mClaire_set_base.Id(),
    Core.C_mClaire_push_I.Id(),
    C_Compile_anyObject_I.Id(),
    Core.C_mClaire_get_stack.Id(),
    Core.C_mClaire_put_stack.Id(),
    Core.C_main.Id())
  C_OPT.ToRemove = CEMPTY
  C_OPT.Knowns = MakeSet(ToType(C_relation.Id()),C_arg1.Id(),C_arg2.Id())
  C_OPT.Unsure = MakeConstantList(Core.F__at_property1(ToProperty(Core.C__plus.Id()),C_integer).Id(),Core.F__at_property1(ToProperty(C__star.Id()),C_integer).Id(),Core.F__at_property1(ToProperty(C__dash.Id()),C_integer).Id())
  C_OPT.SimpleOperations = MakeSet(ToType(C_property.Id()),Core.C__plus.Id(),
    C__dash.Id(),
    C__7.Id(),
    C__star.Id())
  C_OPT.NonIdentifiableSet = MakeSet(ToType(C_class.Id()),C_object.Id(),
    C_list.Id(),
    C_set.Id(),
    C_void.Id(),
    C_any.Id(),
    C_collection.Id(),
    C_type_expression.Id(),
    C_type.Id(),
    C_primitive.Id(),
    C_string.Id(),
    C_char.Id(),
    C_bag.Id(),
    C_tuple.Id(),
    C_port.Id())
  
  
  Core.F_attach_method(C_safe.AddMethod(Signature(C_any.Id(),C_any.Id()),0,MakeFunction1(E_safe_any,"safe_any")),MakeString("osystem.cl:192")).Typing = MakeFunction1(E_safe_any_type,"safe_any_type").Id()
  
  _ = Core.F_attach_method(C_c_type.AddMethod(Signature(C_any.Id(),C_type.Id()),1,MakeFunction1(E_c_type_any,"c_type_any")),MakeString("osystem.cl:219"))
  
  _ = Core.F_attach_method(C_Compile_c_strict_code.AddMethod(Signature(C_any.Id(),C_class.Id(),C_any.Id()),1,MakeFunction2(E_Compile_c_strict_code_any,"Compile_c_strict_code_any")),MakeString("osystem.cl:223"))
  
  _ = Core.F_attach_method(C_Compile_c_strict_check.AddMethod(Signature(C_any.Id(),C_class.Id(),C_any.Id()),1,MakeFunction2(E_Compile_c_strict_check_any,"Compile_c_strict_check_any")),MakeString("osystem.cl:232"))
  
  _ = Core.F_attach_method(C_c_code.AddMethod(Signature(C_any.Id(),C_class.Id(),C_any.Id()),1,MakeFunction2(E_c_code_any1,"c_code_any1")),MakeString("osystem.cl:253"))
  
  _ = Core.F_attach_method(C_c_code.AddMethod(Signature(C_any.Id(),C_any.Id()),1,MakeFunction1(E_c_code_any2,"c_code_any2")),MakeString("osystem.cl:277"))
  
  _ = Core.F_attach_method(C_Optimize_get_sort.AddMethod(Signature(C_any.Id(),C_class.Id()),1,MakeFunction1(E_Optimize_get_sort_any,"Optimize_get_sort_any")),MakeString("osystem.cl:281"))
  
  _ = Core.F_attach_method(C_Compile_c_sort.AddMethod(Signature(C_any.Id(),C_class.Id()),1,MakeFunction1(E_Compile_c_sort_any,"Compile_c_sort_any")),MakeString("osystem.cl:325"))
  
  _ = Core.F_attach_method(C_Optimize_selector_psort.AddMethod(Signature(Language.C_Call.Id(),C_class.Id()),0,MakeFunction1(E_Optimize_selector_psort_Call,"Optimize_selector_psort_Call")),MakeString("osystem.cl:334"))
  
  { 
    var expr EID
    { var _CL_obj *Core.GlobalVariable
      C_Compile_NoErrorOptimize = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("NoErrorOptimize",C_Compile)))
      
      _CL_obj = C_Compile_NoErrorOptimize
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = MakeList(ToType(C_any.Id()),Core.F__at_property1(C_nth,C_list).Id(),
        Core.F__at_property1(C_nth_equal,C_list).Id(),
        Core.F__at_property1(C_nth,C_tuple).Id(),
        Core.F__at_property1(C_nth,C_string).Id(),
        Core.C_less_ask.Id(),
        Core.C_glb.Id(),
        Core.F__at_property1(Core.C_Core_tformat,C_string).Id(),
        Core.F__at_property1(ToProperty(C__inf_equal.Id()),C_type_expression).Id(),
        C__sup_equal.Id(),
        Core.F__at_property2(ToProperty(C__star.Id()),MakeConstantList(C_integer.Id(),C_integer.Id())).Id()).Id()
      expr = Core.F_close_global_variable(_CL_obj)
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID
    { var _CL_obj *Core.GlobalVariable
      C_Compile_ForceThrow = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("ForceThrow",C_Compile)))
      
      _CL_obj = C_Compile_ForceThrow
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = ToType(C_method.Id()).EmptyList().Id()
      expr = Core.F_close_global_variable(_CL_obj)
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID
    { var _CL_obj *Core.GlobalVariable
      C_Compile_ForceNotThrow = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("ForceNotThrow",C_Compile)))
      
      _CL_obj = C_Compile_ForceNotThrow
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = ToType(C_method.Id()).EmptyList().Id()
      expr = Core.F_close_global_variable(_CL_obj)
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID
    { var _CL_obj *Core.GlobalVariable
      C_Optimize__startimes_integer_star = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("*times_integer*",It)))
      
      _CL_obj = C_Optimize__startimes_integer_star
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = Core.F__at_property1(ToProperty(C__star.Id()),C_integer).Id()
      expr = Core.F_close_global_variable(_CL_obj)
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID
    { var _CL_obj *Core.GlobalVariable
      C_Optimize__stardiv_integer_star = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("*div_integer*",It)))
      
      _CL_obj = C_Optimize__stardiv_integer_star
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = Core.F__at_property1(ToProperty(C__7.Id()),C_integer).Id()
      expr = Core.F_close_global_variable(_CL_obj)
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID
    { var _CL_obj *Core.GlobalVariable
      C_Optimize__stardiv_float_star = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("*div_float*",It)))
      
      _CL_obj = C_Optimize__stardiv_float_star
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = Core.F__at_property1(ToProperty(C__7.Id()),C_float).Id()
      expr = Core.F_close_global_variable(_CL_obj)
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID
    { var _CL_obj *Core.GlobalVariable
      C_Optimize__starmod_integer_star = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("*mod_integer*",It)))
      
      _CL_obj = C_Optimize__starmod_integer_star
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = Core.F__at_property1(ToProperty(C_mod.Id()),C_integer).Id()
      expr = Core.F_close_global_variable(_CL_obj)
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID
    { var _CL_obj *Core.GlobalVariable
      C_DSHOW = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("DSHOW",C_claire)))
      
      _CL_obj = C_DSHOW
      _CL_obj.Range = ToType(C_boolean.Id())
      _CL_obj.Value = CFALSE.Id()
      expr = Core.F_close_global_variable(_CL_obj)
      } 
    ErrorCheck(expr)} 
  
  _ = Core.F_attach_method(C_Compile_g_throw.AddMethod(Signature(C_any.Id(),C_boolean.Id()),1,MakeFunction1(E_Compile_g_throw_any,"Compile_g_throw_any")),MakeString("osystem.cl:372"))
  
  _ = Core.F_attach_method(C_Optimize_g_throw1.AddMethod(Signature(C_any.Id(),C_boolean.Id()),1,MakeFunction1(E_Optimize_g_throw1_any,"Optimize_g_throw1_any")),MakeString("osystem.cl:413"))
  
  _ = Core.F_attach_method(C_Optimize_notOpt.AddMethod(Signature(Language.C_Call_method.Id(),C_boolean.Id()),1,MakeFunction1(E_Optimize_notOpt_Call_method,"Optimize_notOpt_Call_method")),MakeString("osystem.cl:428"))
  
  _ = Core.F_attach_method(C_Compile_can_throw_ask.AddMethod(Signature(C_property.Id(),C_boolean.Id()),1,MakeFunction1(E_Compile_can_throw_ask_property,"Compile_can_throw?_property")),MakeString("osystem.cl:434"))
  
  _ = Core.F_attach_method(C_Compile_can_throw_ask.AddMethod(Signature(C_method.Id(),C_boolean.Id()),1,MakeFunction1(E_Compile_can_throw_ask_method,"Compile_can_throw?_method")),MakeString("osystem.cl:444"))
  
  { 
    var expr EID
    { var _CL_obj *Core.GlobalVariable
      C_DTHROW = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("DTHROW",C_claire)))
      
      _CL_obj = C_DTHROW
      _CL_obj.Range = ToType(C_any.Id())
      _CL_obj.Value = CNULL
      expr = Core.F_close_global_variable(_CL_obj)
      } 
    ErrorCheck(expr)} 
  
  _ = Core.F_attach_method(C_Compile_can_throw_I.AddMethod(Signature(C_method.Id(),C_boolean.Id()),1,MakeFunction1(E_Compile_can_throw_I_method,"Compile_can_throw!_method")),MakeString("osystem.cl:455"))
  
  _ = Core.F_attach_method(C_Compile_can_throw_status.AddMethod(Signature(C_method.Id(),C_boolean.Id()),1,MakeFunction1(E_Compile_can_throw_status_method,"Compile_can_throw_status_method")),MakeString("osystem.cl:461"))
  
  _ = Core.F_attach_method(C_s_throw.AddMethod(Signature(C_method.Id(),C_void.Id()),1,MakeFunction1(E_s_throw_method,"s_throw_method")),MakeString("osystem.cl:469"))
  
  _ = Core.F_attach_method(C_Optimize_legal_ask.AddMethod(Signature(C_module.Id(),C_any.Id(),C_boolean.Id()),0,MakeFunction2(E_Optimize_legal_ask_module,"Optimize_legal?_module")),MakeString("osystem.cl:486"))
  
  _ = Core.F_attach_method(C_Optimize_legal_ask.AddMethod(Signature(C_environment.Id(),C_any.Id(),C_any.Id()),0,MakeFunction2(E_Optimize_legal_ask_environment,"Optimize_legal?_environment")),MakeString("osystem.cl:488"))
  
  _ = Core.F_attach_method(C_Optimize_c_register.AddMethod(Signature(Core.F_U_type(ToType(C_thing.Id()),ToType(C_class.Id())).Id(),C_any.Id()),0,MakeFunction1(E_Optimize_c_register_object,"Optimize_c_register_object")),MakeString("osystem.cl:494"))
  
  _ = Core.F_attach_method(C_Optimize_c_register.AddMethod(Signature(C_property.Id(),C_any.Id()),0,MakeFunction1(E_Optimize_c_register_property,"Optimize_c_register_property")),MakeString("osystem.cl:503"))
  
  _ = Core.F_attach_method(C_Optimize_selector_register.AddMethod(Signature(C_property.Id(),C_any.Id()),0,MakeFunction1(E_Optimize_selector_register_property,"Optimize_selector_register_property")),MakeString("osystem.cl:508"))
  
  _ = Core.F_attach_method(C_Optimize_stable_ask.AddMethod(Signature(C_relation.Id(),C_boolean.Id()),0,MakeFunction1(E_Optimize_stable_ask_relation,"Optimize_stable?_relation")),MakeString("osystem.cl:517"))
  
  _ = Core.F_attach_method(C_Compile_get_module.AddMethod(Signature(Core.F_U_type(ToType(C_thing.Id()),ToType(C_class.Id())).Id(),C_any.Id()),0,MakeFunction1(E_Compile_get_module_object,"Compile_get_module_object")),MakeString("osystem.cl:522"))
  
  _ = Core.F_attach_method(Reader.C_known_I.AddMethod(Signature(C_listargs.Id(),C_any.Id()),0,MakeFunction1(E_known_I_listargs,"known!_listargs")),MakeString("osystem.cl:529"))
  
  _ = Core.F_attach_method(C_self_print.AddMethod(Signature(C_Compile_C_cast.Id(),C_void.Id()),1,MakeFunction1(E_self_print_C_cast,"self_print_C_cast")),MakeString("otool.cl:56"))
  
  _ = Core.F_attach_method(C_Compile_c_gc_ask.AddMethod(Signature(C_Compile_C_cast.Id(),C_boolean.Id()),0,MakeFunction1(E_Compile_c_gc_ask_C_cast,"Compile_c_gc?_C_cast")),MakeString("otool.cl:57"))
  
  _ = Core.F_attach_method(C_c_type.AddMethod(Signature(C_Compile_C_cast.Id(),C_type.Id()),0,MakeFunction1(E_c_type_C_cast,"c_type_C_cast")),MakeString("otool.cl:58"))
  
  _ = Core.F_attach_method(C_c_code.AddMethod(Signature(C_Compile_C_cast.Id(),C_class.Id(),C_any.Id()),1,MakeFunction2(E_c_code_C_cast,"c_code_C_cast")),MakeString("otool.cl:62"))
  
  _ = Core.F_attach_method(C_self_print.AddMethod(Signature(C_Optimize_Pattern.Id(),C_void.Id()),1,MakeFunction1(E_self_print_Pattern,"self_print_Pattern")),MakeString("otool.cl:68"))
  
  _ = Core.F_attach_method(C__Z.AddMethod(Signature(C_any.Id(),C_Optimize_Pattern.Id(),C_boolean.Id()),1,MakeFunction2(E__Z_any3,"%_any3")),MakeString("otool.cl:74"))
  
  _ = Core.F_attach_method(Core.C_glb.AddMethod(Signature(C_Optimize_Pattern.Id(),C_type_expression.Id(),C_type_expression.Id()),1,MakeFunction2(E_glb_Pattern,"glb_Pattern")),MakeString("otool.cl:81"))
  
  _ = Core.F_attach_method(Core.C_less_ask.AddMethod(Signature(C_Optimize_Pattern.Id(),C_type_expression.Id(),C_boolean.Id()),0,MakeFunction2(E_less_ask_Pattern,"less?_Pattern")),MakeString("otool.cl:88"))
  
  _ = Core.F_attach_method(Core.C_less_ask.AddMethod(Signature(C_type_expression.Id(),C_Optimize_Pattern.Id(),C_boolean.Id()),1,MakeFunction2(E_less_ask_type_expression2,"less?_type_expression2")),MakeString("otool.cl:95"))
  
  _ = Core.F_attach_method(C_nth.AddMethod(Signature(C_property.Id(),C_tuple.Id(),C_Optimize_Pattern.Id()),0,MakeFunction2(E_nth_property,"nth_property")),MakeString("otool.cl:98"))
  
  _ = Core.F_attach_method(C_Compile_warn.AddMethod(Signature(C_void.Id(),C_void.Id()),0,MakeFunction1(E_Compile_warn_void,"Compile_warn_void")),MakeString("otool.cl:108"))
  
  _ = Core.F_attach_method(C_Compile_Cerror.AddMethod(Signature(C_string.Id(),C_listargs.Id(),CEMPTY.Id()),1,MakeFunction2(E_Compile_Cerror_string,"Compile_Cerror_string")),MakeString("otool.cl:113"))
  
  _ = Core.F_attach_method(C_Compile_notice.AddMethod(Signature(C_void.Id(),C_void.Id()),0,MakeFunction1(E_Compile_notice_void,"Compile_notice_void")),MakeString("otool.cl:119"))
  
  _ = Core.F_attach_method(C_Optimize_c_warn.AddMethod(Signature(Language.C_Call.Id(),C_any.Id(),C_any.Id()),1,MakeFunction2(E_Optimize_c_warn_Call,"Optimize_c_warn_Call")),MakeString("otool.cl:133"))
  
  _ = Core.F_attach_method(C_Optimize_c_warn.AddMethod(Signature(Language.C_Super.Id(),C_any.Id(),C_any.Id()),1,MakeFunction2(E_Optimize_c_warn_Super,"Optimize_c_warn_Super")),MakeString("otool.cl:143"))
  
  _ = Core.F_attach_method(C_Optimize_c_warn.AddMethod(Signature(C_property.Id(),
    C_list.Id(),
    C_list.Id(),
    C_any.Id()),1,MakeFunction3(E_Optimize_c_warn_property,"Optimize_c_warn_property")),MakeString("otool.cl:150"))
  
  _ = Core.F_attach_method(C_Optimize_c_warn.AddMethod(Signature(C_Variable.Id(),
    C_any.Id(),
    C_type.Id(),
    C_any.Id()),1,MakeFunction3(E_Optimize_c_warn_Variable,"Optimize_c_warn_Variable")),MakeString("otool.cl:167"))
  
  _ = Core.F_attach_method(C_Compile_sort_equal.AddMethod(Signature(C_class.Id(),C_class.Id(),C_any.Id()),0,MakeFunction2(E_Compile_sort_equal_class,"Compile_sort=_class")),MakeString("otool.cl:184"))
  
  _ = Core.F_attach_method(C_Compile_psort.AddMethod(Signature(C_any.Id(),C_class.Id()),0,MakeFunction1(E_Compile_psort_any,"Compile_psort_any")),MakeString("otool.cl:190"))
  
  _ = Core.F_attach_method(C_Compile_osort.AddMethod(Signature(C_any.Id(),C_class.Id()),0,MakeFunction1(E_Compile_osort_any,"Compile_osort_any")),MakeString("otool.cl:195"))
  
  _ = Core.F_attach_method(Core.C_sort.AddMethod(Signature(C_Variable.Id(),C_class.Id()),0,MakeFunction1(E_sort_Variable,"sort_Variable")),MakeString("otool.cl:201"))
  
  _ = Core.F_attach_method(C_Compile_stupid_t.AddMethod(Signature(C_any.Id(),C_class.Id()),1,MakeFunction1(E_Compile_stupid_t_any1,"Compile_stupid_t_any1")),MakeString("otool.cl:206"))
  
  _ = Core.F_attach_method(C_Compile_stupid_t.AddMethod(Signature(C_any.Id(),C_any.Id(),C_boolean.Id()),1,MakeFunction2(E_Compile_stupid_t_any2,"Compile_stupid_t_any2")),MakeString("otool.cl:212"))
  
  _ = Core.F_attach_method(C_Optimize_extended_ask.AddMethod(Signature(C_type.Id(),C_boolean.Id()),1,MakeFunction1(E_Optimize_extended_ask_type,"Optimize_extended?_type")),MakeString("otool.cl:219"))
  
  _ = Core.F_attach_method(C_Optimize_extends.AddMethod(Signature(C_type.Id(),C_type.Id()),0,MakeFunction1(E_Optimize_extends_type,"Optimize_extends_type")),MakeString("otool.cl:225"))
  
  _ = Core.F_attach_method(C_Optimize_sort_abstract_I.AddMethod(Signature(C_type.Id(),C_type.Id()),0,MakeFunction1(E_Optimize_sort_abstract_I_type,"Optimize_sort_abstract!_type")),MakeString("otool.cl:232"))
  
  _ = Core.F_attach_method(C_Optimize_sort_abstract_ask.AddMethod(Signature(C_type.Id(),C_boolean.Id()),0,MakeFunction1(E_Optimize_sort_abstract_ask_type,"Optimize_sort_abstract?_type")),MakeString("otool.cl:234"))
  
  _ = Core.F_attach_method(C_Optimize_ptype.AddMethod(Signature(C_type.Id(),C_type.Id()),0,MakeFunction1(E_Optimize_ptype_type,"Optimize_ptype_type")),MakeString("otool.cl:239"))
  
  _ = Core.F_attach_method(C_Optimize_pmember.AddMethod(Signature(C_type.Id(),C_type.Id()),0,MakeFunction1(E_Optimize_pmember_type,"Optimize_pmember_type")),MakeString("otool.cl:243"))
  
  _ = Core.F_attach_method(C_Optimize_enumerate_code.AddMethod(Signature(C_any.Id(),C_type.Id(),C_any.Id()),1,MakeFunction2(E_Optimize_enumerate_code_any,"Optimize_enumerate_code_any")),MakeString("otool.cl:252"))
  
  _ = Core.F_attach_method(C_Optimize_range_infers_for.AddMethod(Signature(C_Variable.Id(),
    C_type.Id(),
    C_type.Id(),
    C_any.Id()),0,MakeFunction3(E_Optimize_range_infers_for_Variable,"Optimize_range_infers_for_Variable")),MakeString("otool.cl:266"))
  
  _ = Core.F_attach_method(C_Optimize_range_infers.AddMethod(Signature(C_Variable.Id(),C_type.Id(),C_any.Id()),1,MakeFunction2(E_Optimize_range_infers_Variable,"Optimize_range_infers_Variable")),MakeString("otool.cl:280"))
  
  _ = Core.F_attach_method(C_Optimize_range_sets.AddMethod(Signature(C_any.Id(),C_type.Id(),C_void.Id()),0,MakeFunction2(E_Optimize_range_sets_any,"Optimize_range_sets_any")),MakeString("otool.cl:287"))
  
  _ = Core.F_attach_method(C_Optimize_c_srange.AddMethod(Signature(C_method.Id(),C_class.Id()),1,MakeFunction1(E_Optimize_c_srange_method,"Optimize_c_srange_method")),MakeString("otool.cl:300"))
  
  _ = Core.F_attach_method(C_Compile_nativeVar_ask.AddMethod(Signature(Core.C_global_variable.Id(),C_boolean.Id()),0,MakeFunction1(E_Compile_nativeVar_ask_global_variable,"Compile_nativeVar?_global_variable")),MakeString("otool.cl:306"))
  
  _ = Core.F_attach_method(C_Compile_return_type.AddMethod(Signature(C_any.Id(),C_type.Id()),1,MakeFunction1(E_Compile_return_type_any,"Compile_return_type_any")),MakeString("otool.cl:319"))
  
  _ = Core.F_attach_method(C_c_code.AddMethod(Signature(Core.F_U_type(Core.F_U_type(ToType(C_type_operator.Id()),ToType(C_Reference.Id())),ToType(C_Optimize_Pattern.Id())).Id(),C_class.Id(),C_any.Id()),1,MakeFunction2(E_c_code_type_expression,"c_code_type_expression")),MakeString("otool.cl:327"))
  
  _ = Core.F_attach_method(C_Compile_self_code.AddMethod(Signature(C_subtype.Id(),C_any.Id()),1,MakeFunction1(E_Compile_self_code_subtype,"Compile_self_code_subtype")),MakeString("otool.cl:331"))
  
  _ = Core.F_attach_method(C_Compile_self_code.AddMethod(Signature(C_Param.Id(),C_any.Id()),1,MakeFunction1(E_Compile_self_code_Param,"Compile_self_code_Param")),MakeString("otool.cl:338"))
  
  _ = Core.F_attach_method(C_Compile_self_code.AddMethod(Signature(C_Union.Id(),C_any.Id()),1,MakeFunction1(E_Compile_self_code_Union,"Compile_self_code_Union")),MakeString("otool.cl:341"))
  
  _ = Core.F_attach_method(C_Compile_self_code.AddMethod(Signature(C_Interval.Id(),C_any.Id()),0,MakeFunction1(E_Compile_self_code_Interval,"Compile_self_code_Interval")),MakeString("otool.cl:344"))
  
  _ = Core.F_attach_method(C_Compile_self_code.AddMethod(Signature(C_Reference.Id(),C_any.Id()),0,MakeFunction1(E_Compile_self_code_Reference,"Compile_self_code_Reference")),MakeString("otool.cl:347"))
  
  _ = Core.F_attach_method(C_Compile_self_code.AddMethod(Signature(C_Optimize_Pattern.Id(),C_any.Id()),0,MakeFunction1(E_Compile_self_code_Pattern,"Compile_self_code_Pattern")),MakeString("otool.cl:354"))
  
  _ = Core.F_attach_method(C_Optimize_member_code.AddMethod(Signature(C_class.Id(),C_any.Id(),C_any.Id()),1,MakeFunction2(E_Optimize_member_code_class,"Optimize_member_code_class")),MakeString("otool.cl:364"))
  
  _ = Core.F_attach_method(C_Optimize_member_code.AddMethod(Signature(C_type_operator.Id(),C_any.Id(),C_any.Id()),1,MakeFunction2(E_Optimize_member_code_type_operator,"Optimize_member_code_type_operator")),MakeString("otool.cl:368"))
  
  _ = Core.F_attach_method(C_Optimize_member_code.AddMethod(Signature(C_Union.Id(),C_any.Id(),C_any.Id()),1,MakeFunction2(E_Optimize_member_code_Union,"Optimize_member_code_Union")),MakeString("otool.cl:372"))
  
  _ = Core.F_attach_method(C_Optimize_member_code.AddMethod(Signature(C_Interval.Id(),C_any.Id(),C_any.Id()),1,MakeFunction2(E_Optimize_member_code_Interval,"Optimize_member_code_Interval")),MakeString("otool.cl:377"))
  
  _ = Core.F_attach_method(C_Optimize_member_code.AddMethod(Signature(C_Param.Id(),C_any.Id(),C_any.Id()),1,MakeFunction2(E_Optimize_member_code_Param,"Optimize_member_code_Param")),MakeString("otool.cl:384"))
  
  _ = Core.F_attach_method(C_Optimize_member_code.AddMethod(Signature(C_tuple.Id(),C_any.Id(),C_any.Id()),1,MakeFunction2(E_Optimize_member_code_tuple,"Optimize_member_code_tuple")),MakeString("otool.cl:392"))
  
  _ = Core.F_attach_method(C_Optimize_member_code.AddMethod(Signature(C_any.Id(),C_any.Id(),C_any.Id()),1,MakeFunction2(E_Optimize_member_code_any,"Optimize_member_code_any")),MakeString("otool.cl:401"))
  
  _ = Core.F_inlineok_ask_method(
  Core.F_attach_method(C__Z.AddMethod(Signature(C_any.Id(),F_nth_property(ToProperty(C__dot_dot.Id()),MakeTuple(C_any.Id(),C_any.Id())).Id(),C_boolean.Id()),1,MakeFunction2(E__Z_any4,"%_any4")),MakeString("otool.cl:405")),
  MakeString("lambda[(x:any,y:any),(x <= eval(y.args[2]) & eval(y.args[1]) <= x)]"))
  
  _ = Core.F_inlineok_ask_method(
  Core.F_attach_method(C__Z.AddMethod(Signature(C_any.Id(),F_nth_property(ToProperty(Core.C_but.Id()),MakeTuple(C_any.Id(),C_any.Id())).Id(),C_boolean.Id()),1,MakeFunction2(E__Z_any5,"%_any5")),MakeString("otool.cl:408")),
  MakeString("lambda[(x:any,y:any),(x % eval(y.args[1]) & x != eval(y.args[2]))]"))
  
  _ = Core.F_attach_method(C_Compile_Variable_I.AddMethod(Signature(C_symbol.Id(),
    C_integer.Id(),
    C_any.Id(),
    C_Variable.Id()),0,MakeFunction3(E_Compile_Variable_I_symbol,"Compile_Variable!_symbol")),MakeString("otool.cl:418"))
  
  _ = Core.F_attach_method(C_Compile_get_indexed.AddMethod(Signature(C_class.Id(),C_list.Id()),0,MakeFunction1(E_Compile_get_indexed_class,"Compile_get_indexed_class")),MakeString("otool.cl:420"))
  
  _ = Core.F_attach_method(C_Compile_designated_ask.AddMethod(Signature(C_any.Id(),C_boolean.Id()),1,MakeFunction1(E_Compile_designated_ask_any,"Compile_designated?_any")),MakeString("otool.cl:442"))
  
  _ = Core.F_attach_method(C_Compile_identifiable_ask.AddMethod(Signature(C_any.Id(),C_boolean.Id()),1,MakeFunction1(E_Compile_identifiable_ask_any,"Compile_identifiable?_any")),MakeString("otool.cl:451"))
  
  _ = Core.F_attach_method(C_Optimize_c_inline.AddMethod(Signature(C_method.Id(),
    C_list.Id(),
    C_class.Id(),
    C_any.Id()),1,MakeFunction3(E_Optimize_c_inline_method1,"Optimize_c_inline_method1")),MakeString("otool.cl:458"))
  
  _ = Core.F_attach_method(C_Optimize_c_inline.AddMethod(Signature(C_method.Id(),C_list.Id(),C_any.Id()),1,MakeFunction2(E_Optimize_c_inline_method2,"Optimize_c_inline_method2")),MakeString("otool.cl:479"))
  
  _ = Core.F_attach_method(C_Optimize_c_inline_arg_ask.AddMethod(Signature(C_any.Id(),C_any.Id()),1,MakeFunction1(E_Optimize_c_inline_arg_ask_any,"Optimize_c_inline_arg?_any")),MakeString("otool.cl:490"))
  
  _ = Core.F_attach_method(C_Optimize_c_substitution.AddMethod(Signature(C_any.Id(),
    Core.F_nth_class1(C_list,ToType(C_Variable.Id())).Id(),
    C_list.Id(),
    C_boolean.Id(),
    C_any.Id()),1,MakeFunction4(E_Optimize_c_substitution_any,"Optimize_c_substitution_any")),MakeString("otool.cl:521"))
  
  _ = Core.F_attach_method(Core.C_eval.AddMethod(Signature(C_any.Id(),C_class.Id(),C_any.Id()),1,MakeFunction2(E_eval_any2,"eval_any2")),MakeString("otool.cl:524"))
  
  _ = Core.F_attach_method(C_Optimize_bound_variables.AddMethod(Signature(C_any.Id(),C_list.Id()),1,MakeFunction1(E_Optimize_bound_variables_any,"Optimize_bound_variables_any")),MakeString("otool.cl:535"))
  
  _ = Core.F_attach_method(C_Optimize_c_boolean.AddMethod(Signature(C_any.Id(),C_any.Id()),1,MakeFunction1(E_Optimize_c_boolean_any,"Optimize_c_boolean_any")),MakeString("otool.cl:546"))
  
  _ = Core.F_attach_method(C_Compile_Do_I.AddMethod(Signature(C_list.Id(),C_any.Id()),0,MakeFunction1(E_Compile_Do_I_list,"Compile_Do!_list")),MakeString("otool.cl:550"))
  
  C_Optimize_ambiguous = ToKeyword(new(ClaireKeyword).IsNamed(C_keyword,MakeSymbol("ambiguous",It)))
  
  
  _ = Core.F_attach_method(C_Optimize_restriction_I.AddMethod(Signature(C_property.Id(),
    C_list.Id(),
    C_boolean.Id(),
    C_any.Id()),0,MakeFunction3(E_Optimize_restriction_I_property,"Optimize_restriction!_property")),MakeString("ocall.cl:32"))
  
  _ = Core.F_attach_method(C_Optimize_restriction_I.AddMethod(Signature(C_list.Id(),
    C_list.Id(),
    C_boolean.Id(),
    C_any.Id()),0,MakeFunction3(E_Optimize_restriction_I_list,"Optimize_restriction!_list")),MakeString("ocall.cl:53"))
  
  _ = Core.F_attach_method(C_findr.AddMethod(Signature(C_property.Id(),C_list.Id(),C_void.Id()),1,MakeFunction2(E_findr_property,"findr_property")),MakeString("ocall.cl:60"))
  
  _ = Core.F_attach_method(C_dmatch_ask.AddMethod(Signature(C_list.Id(),C_list.Id(),C_boolean.Id()),0,MakeFunction2(E_dmatch_ask_list,"dmatch?_list")),MakeString("ocall.cl:70"))
  
  _ = Core.F_attach_method(C_Optimize_restriction_I.AddMethod(Signature(C_class.Id(),
    C_list.Id(),
    C_list.Id(),
    C_any.Id()),0,MakeFunction3(E_Optimize_restriction_I_class,"Optimize_restriction!_class")),MakeString("ocall.cl:79"))
  
  _ = Core.F_attach_method(C_Optimize_use_range.AddMethod(Signature(C_method.Id(),C_list.Id(),C_type.Id()),1,MakeFunction2(E_Optimize_use_range_method,"Optimize_use_range_method")),MakeString("ocall.cl:108"))
  
  _ = Core.F_attach_method(C_c_type.AddMethod(Signature(Language.C_Call.Id(),C_type.Id()),1,MakeFunction1(E_c_type_Call,"c_type_Call")),MakeString("ocall.cl:161"))
  
  _ = Core.F_attach_method(C_c_code.AddMethod(Signature(Language.C_Call.Id(),C_any.Id()),1,MakeFunction1(E_c_code_Call,"c_code_Call")),MakeString("ocall.cl:165"))
  
  _ = Core.F_attach_method(C_Optimize_c_code_call.AddMethod(Signature(Language.C_Call.Id(),C_class.Id(),C_any.Id()),1,MakeFunction2(E_Optimize_c_code_call_Call,"Optimize_c_code_call_Call")),MakeString("ocall.cl:234"))
  
  _ = Core.F_attach_method(C_Optimize_open_message.AddMethod(Signature(C_property.Id(),C_list.Id(),Language.C_Call.Id()),1,MakeFunction2(E_Optimize_open_message_property,"Optimize_open_message_property")),MakeString("ocall.cl:244"))
  
  _ = Core.F_attach_method(C_Optimize_daccess.AddMethod(Signature(C_any.Id(),C_boolean.Id(),C_any.Id()),1,MakeFunction2(E_Optimize_daccess_any,"Optimize_daccess_any")),MakeString("ocall.cl:267"))
  
  _ = Core.F_attach_method(C_c_type.AddMethod(Signature(Language.C_Call_slot.Id(),C_type.Id()),0,MakeFunction1(E_c_type_Call_slot,"c_type_Call_slot")),MakeString("ocall.cl:269"))
  
  _ = Core.F_attach_method(C_c_type.AddMethod(Signature(Language.C_Call_table.Id(),C_type.Id()),0,MakeFunction1(E_c_type_Call_table,"c_type_Call_table")),MakeString("ocall.cl:270"))
  
  _ = Core.F_attach_method(C_c_type.AddMethod(Signature(Language.C_Call_array.Id(),C_type.Id()),0,MakeFunction1(E_c_type_Call_array,"c_type_Call_array")),MakeString("ocall.cl:271"))
  
  _ = Core.F_attach_method(C_Optimize_c_code_write.AddMethod(Signature(Language.C_Call.Id(),C_any.Id()),1,MakeFunction1(E_Optimize_c_code_write_Call,"Optimize_c_code_write_Call")),MakeString("ocall.cl:307"))
  
  _ = Core.F_attach_method(C_Optimize_c_code_hold.AddMethod(Signature(C_property.Id(),
    C_any.Id(),
    C_any.Id(),
    C_boolean.Id(),
    C_any.Id()),1,MakeFunction4(E_Optimize_c_code_hold_property,"Optimize_c_code_hold_property")),MakeString("ocall.cl:323"))
  
  _ = Core.F_attach_method(C_Optimize_c_code_add.AddMethod(Signature(Language.C_Call.Id(),C_any.Id()),1,MakeFunction1(E_Optimize_c_code_add_Call,"Optimize_c_code_add_Call")),MakeString("ocall.cl:350"))
  
  _ = Core.F_attach_method(C_Optimize_c_code_add_bag.AddMethod(Signature(Language.C_Call.Id(),C_any.Id()),1,MakeFunction1(E_Optimize_c_code_add_bag_Call,"Optimize_c_code_add_bag_Call")),MakeString("ocall.cl:366"))
  
  _ = Core.F_attach_method(C_Optimize_c_code_delete.AddMethod(Signature(Language.C_Call.Id(),C_any.Id()),1,MakeFunction1(E_Optimize_c_code_delete_Call,"Optimize_c_code_delete_Call")),MakeString("ocall.cl:381"))
  
  _ = Core.F_attach_method(C_Optimize_c_code_not.AddMethod(Signature(Language.C_Select.Id(),C_any.Id()),1,MakeFunction1(E_Optimize_c_code_not_Select,"Optimize_c_code_not_Select")),MakeString("ocall.cl:387"))
  
  _ = Core.F_attach_method(C_Optimize_c_code_belong.AddMethod(Signature(Language.C_Call.Id(),C_any.Id()),1,MakeFunction1(E_Optimize_c_code_belong_Call,"Optimize_c_code_belong_Call")),MakeString("ocall.cl:405"))
  
  _ = Core.F_attach_method(C_Optimize_c_code_nth.AddMethod(Signature(Language.C_Call.Id(),C_any.Id()),1,MakeFunction1(E_Optimize_c_code_nth_Call,"Optimize_c_code_nth_Call")),MakeString("ocall.cl:434"))
  
  _ = Core.F_attach_method(C_Optimize_c_code_table.AddMethod(Signature(Language.C_Call.Id(),C_any.Id()),1,MakeFunction1(E_Optimize_c_code_table_Call,"Optimize_c_code_table_Call")),MakeString("ocall.cl:460"))
  
  _ = Core.F_attach_method(C_Optimize_c_code_array.AddMethod(Signature(Language.C_Call.Id(),C_any.Id()),1,MakeFunction1(E_Optimize_c_code_array_Call,"Optimize_c_code_array_Call")),MakeString("ocall.cl:478"))
  
  _ = Core.F_attach_method(C_Optimize_Update_ask.AddMethod(Signature(C_relation.Id(),
    C_any.Id(),
    C_any.Id(),
    C_boolean.Id()),1,MakeFunction3(E_Optimize_Update_ask_relation1,"Optimize_Update?_relation1")),MakeString("ocall.cl:489"))
  
  _ = Core.F_attach_method(C_Optimize_Update_ask.AddMethod(Signature(C_relation.Id(),C_relation.Id(),C_boolean.Id()),0,MakeFunction2(E_Optimize_Update_ask_relation2,"Optimize_Update?_relation2")),MakeString("ocall.cl:493"))
  
  _ = Core.F_attach_method(C_c_type.AddMethod(Signature(Language.C_Update.Id(),C_type.Id()),0,MakeFunction1(E_c_type_Update,"c_type_Update")),MakeString("ocall.cl:496"))
  
  _ = Core.F_attach_method(C_Compile_update_write_ask.AddMethod(Signature(Language.C_Update.Id(),C_boolean.Id()),0,MakeFunction1(E_Compile_update_write_ask_Update,"Compile_update_write?_Update")),MakeString("ocall.cl:501"))
  
  _ = Core.F_attach_method(C_Optimize_c_code_method.AddMethod(Signature(C_method.Id(),
    C_list.Id(),
    C_list.Id(),
    C_any.Id()),1,MakeFunction3(E_Optimize_c_code_method_method1,"Optimize_c_code_method_method1")),MakeString("ocall.cl:510"))
  
  _ = Core.F_attach_method(C_Optimize_c_code_method.AddMethod(Signature(C_method.Id(),
    C_list.Id(),
    C_list.Id(),
    C_class.Id(),
    C_any.Id()),1,MakeFunction4(E_Optimize_c_code_method_method2,"Optimize_c_code_method_method2")),MakeString("ocall.cl:523"))
  
  _ = Core.F_attach_method(C_Optimize_Call_method_I.AddMethod(Signature(C_method.Id(),C_list.Id(),C_any.Id()),0,MakeFunction2(E_Optimize_Call_method_I_method,"Optimize_Call_method!_method")),MakeString("ocall.cl:531"))
  
  _ = Core.F_attach_method(C_c_type.AddMethod(Signature(Language.C_Call_method.Id(),C_type.Id()),1,MakeFunction1(E_c_type_Call_method,"c_type_Call_method")),MakeString("ocall.cl:537"))
  
  _ = Core.F_attach_method(C_c_code.AddMethod(Signature(Language.C_Call_method.Id(),C_any.Id()),1,MakeFunction1(E_c_code_Call_method,"c_code_Call_method")),MakeString("ocall.cl:543"))
  
  C_Compile_functional_I = MakeProperty("functional!",3,C_Compile)
  C_Compile_functional_I.Open = 3
  
  
  _ = Core.F_attach_method(C_Compile_functional_I.AddMethod(Signature(C_method.Id(),C_function.Id()),1,MakeFunction1(E_Compile_functional_I_method,"Compile_functional!_method")),MakeString("ocall.cl:551"))
  
  _ = Core.F_inlineok_ask_method(
  Core.F_attach_method(C_Optimize_nth_type_check.AddMethod(Signature(C_type.Id(),
    C_type.Id(),
    C_type.Id(),
    C_any.Id()),0,MakeFunction3(E_Optimize_nth_type_check_type,"Optimize_nth_type_check_type")),MakeString("ocall.cl:557")),
  MakeString("lambda[(tl:type,ti:type,tx:type),(if not(tx <= member(tl)) (Compile/warn(), trace(1, \"unsafe update on bag: type ~S into ~S [252]\\n\", tx, tl)) else false, tx)]"))
  
  ToMethod(Core.F__at_property1(C_nth_equal,C_list).Id()).Typing = C_Optimize_nth_type_check.Id()
  
  _ = Core.F_attach_method(C_Optimize_c_inline_ask.AddMethod(Signature(C_method.Id(),C_list.Id(),C_boolean.Id()),1,MakeFunction2(E_Optimize_c_inline_ask_method,"Optimize_c_inline?_method")),MakeString("ocall.cl:572"))
  
  _ = Core.F_attach_method(C_Optimize_inline_optimize_ask.AddMethod(Signature(Language.C_Call.Id(),C_any.Id()),1,MakeFunction1(E_Optimize_inline_optimize_ask_Call,"Optimize_inline_optimize?_Call")),MakeString("ocall.cl:581"))
  
  _ = Core.F_attach_method(C_c_type.AddMethod(Signature(Language.C_Assign.Id(),C_type.Id()),1,MakeFunction1(E_c_type_Assign,"c_type_Assign")),MakeString("ocontrol.cl:24"))
  
  _ = Core.F_attach_method(C_c_code.AddMethod(Signature(Language.C_Assign.Id(),C_any.Id()),1,MakeFunction1(E_c_code_Assign,"c_code_Assign")),MakeString("ocontrol.cl:34"))
  
  _ = Core.F_attach_method(C_c_type.AddMethod(Signature(Language.C_Gassign.Id(),C_type.Id()),1,MakeFunction1(E_c_type_Gassign,"c_type_Gassign")),MakeString("ocontrol.cl:38"))
  
  _ = Core.F_attach_method(C_c_code.AddMethod(Signature(Language.C_Gassign.Id(),C_any.Id()),1,MakeFunction1(E_c_code_Gassign,"c_code_Gassign")),MakeString("ocontrol.cl:47"))
  
  _ = Core.F_attach_method(C_c_type.AddMethod(Signature(Language.C_And.Id(),C_type.Id()),0,MakeFunction1(E_c_type_And,"c_type_And")),MakeString("ocontrol.cl:51"))
  
  _ = Core.F_attach_method(C_c_code.AddMethod(Signature(Language.C_And.Id(),C_any.Id()),1,MakeFunction1(E_c_code_And,"c_code_And")),MakeString("ocontrol.cl:56"))
  
  _ = Core.F_attach_method(C_c_type.AddMethod(Signature(Language.C_Or.Id(),C_type.Id()),0,MakeFunction1(E_c_type_Or,"c_type_Or")),MakeString("ocontrol.cl:58"))
  
  _ = Core.F_attach_method(C_c_code.AddMethod(Signature(Language.C_Or.Id(),C_any.Id()),1,MakeFunction1(E_c_code_Or,"c_code_Or")),MakeString("ocontrol.cl:63"))
  
  _ = Core.F_attach_method(C_c_type.AddMethod(Signature(Language.C_Quote.Id(),C_type.Id()),0,MakeFunction1(E_c_type_Quote,"c_type_Quote")),MakeString("ocontrol.cl:66"))
  
  _ = Core.F_attach_method(C_c_code.AddMethod(Signature(Language.C_Quote.Id(),CEMPTY.Id()),1,MakeFunction1(E_c_code_Quote,"c_code_Quote")),MakeString("ocontrol.cl:68"))
  
  _ = Core.F_attach_method(C_c_type.AddMethod(Signature(Language.C_Return.Id(),C_type.Id()),0,MakeFunction1(E_c_type_Return,"c_type_Return")),MakeString("ocontrol.cl:70"))
  
  _ = Core.F_attach_method(C_c_code.AddMethod(Signature(Language.C_Return.Id(),C_any.Id()),1,MakeFunction1(E_c_code_Return,"c_code_Return")),MakeString("ocontrol.cl:72"))
  
  _ = Core.F_attach_method(C_c_type.AddMethod(Signature(Language.C_Handle.Id(),C_type.Id()),1,MakeFunction1(E_c_type_Handle,"c_type_Handle")),MakeString("ocontrol.cl:75"))
  
  _ = Core.F_attach_method(C_c_code.AddMethod(Signature(Language.C_Handle.Id(),C_class.Id(),C_any.Id()),1,MakeFunction2(E_c_code_Handle,"c_code_Handle")),MakeString("ocontrol.cl:81"))
  
  _ = Core.F_attach_method(C_c_type.AddMethod(Signature(Language.C_Cast.Id(),C_type.Id()),0,MakeFunction1(E_c_type_Cast,"c_type_Cast")),MakeString("ocontrol.cl:90"))
  
  _ = Core.F_attach_method(C_c_code.AddMethod(Signature(Language.C_Cast.Id(),C_any.Id()),1,MakeFunction1(E_c_code_Cast,"c_code_Cast")),MakeString("ocontrol.cl:104"))
  
  _ = Core.F_attach_method(C_c_type.AddMethod(Signature(Language.C_Super.Id(),C_type.Id()),1,MakeFunction1(E_c_type_Super,"c_type_Super")),MakeString("ocontrol.cl:118"))
  
  _ = Core.F_attach_method(C_c_code.AddMethod(Signature(Language.C_Super.Id(),C_any.Id()),1,MakeFunction1(E_c_code_Super,"c_code_Super")),MakeString("ocontrol.cl:132"))
  
  C_Optimize_Call_function2 = NewClass("Call_function2",Language.C_Optimized_instruction,It)
  Core.F_close_slot(C_Optimize_Call_function2.AddSlot(C_arg,ToType(C_function.Id()),CNULL))
  Core.F_close_slot(C_Optimize_Call_function2.AddSlot(C_args,ToType(C_list.Id()),ToType(C_any.Id()).EmptyList().Id()))
  
  _ = Core.F_attach_method(C_self_print.AddMethod(Signature(C_Optimize_Call_function2.Id(),C_void.Id()),1,MakeFunction1(E_self_print_Call_function2,"self_print_Call_function2")),MakeString("ocontrol.cl:138"))
  
  _ = Core.F_attach_method(C_c_type.AddMethod(Signature(C_Optimize_Call_function2.Id(),C_type.Id()),0,MakeFunction1(E_c_type_Call_function2,"c_type_Call_function2")),MakeString("ocontrol.cl:140"))
  
  _ = Core.F_attach_method(C_c_code.AddMethod(Signature(C_Optimize_Call_function2.Id(),C_any.Id()),1,MakeFunction1(E_c_code_Call_function2,"c_code_Call_function2")),MakeString("ocontrol.cl:143"))
  
  _ = Core.F_attach_method(C_c_code.AddMethod(Signature(Language.C_Assert.Id(),C_any.Id()),1,MakeFunction1(E_c_code_Assert,"c_code_Assert")),MakeString("ocontrol.cl:154"))
  
  _ = Core.F_attach_method(C_c_code.AddMethod(Signature(Language.C_Trace.Id(),C_any.Id()),1,MakeFunction1(E_c_code_Trace,"c_code_Trace")),MakeString("ocontrol.cl:169"))
  
  _ = Core.F_attach_method(C_c_type.AddMethod(Signature(Language.C_Assert.Id(),C_type.Id()),0,MakeFunction1(E_c_type_Assert,"c_type_Assert")),MakeString("ocontrol.cl:171"))
  
  _ = Core.F_attach_method(C_c_type.AddMethod(Signature(Language.C_Trace.Id(),C_type.Id()),0,MakeFunction1(E_c_type_Trace,"c_type_Trace")),MakeString("ocontrol.cl:172"))
  
  _ = Core.F_attach_method(C_c_type.AddMethod(Signature(Language.C_Branch.Id(),C_type.Id()),0,MakeFunction1(E_c_type_Branch,"c_type_Branch")),MakeString("ocontrol.cl:173"))
  
  _ = Core.F_attach_method(C_c_code.AddMethod(Signature(Language.C_Branch.Id(),C_any.Id()),1,MakeFunction1(E_c_code_Branch,"c_code_Branch")),MakeString("ocontrol.cl:181"))
  
  _ = Core.F_attach_method(C_c_code.AddMethod(Signature(Language.C_Macro.Id(),C_class.Id(),C_any.Id()),1,MakeFunction2(E_c_code_Macro,"c_code_Macro")),MakeString("ocontrol.cl:183"))
  
  _ = Core.F_attach_method(C_c_type.AddMethod(Signature(Language.C_Macro.Id(),C_type.Id()),1,MakeFunction1(E_c_type_Macro,"c_type_Macro")),MakeString("ocontrol.cl:185"))
  
  _ = Core.F_attach_method(C_c_type.AddMethod(Signature(Language.C_Printf.Id(),C_type.Id()),0,MakeFunction1(E_c_type_Printf,"c_type_Printf")),MakeString("ocontrol.cl:188"))
  
  _ = Core.F_attach_method(C_c_code.AddMethod(Signature(Language.C_Printf.Id(),C_any.Id()),1,MakeFunction1(E_c_code_Printf,"c_code_Printf")),MakeString("ocontrol.cl:219"))
  
  _ = Core.F_attach_method(C_c_type.AddMethod(Signature(Language.C_Error.Id(),C_type.Id()),0,MakeFunction1(E_c_type_Error,"c_type_Error")),MakeString("ocontrol.cl:222"))
  
  _ = Core.F_attach_method(C_c_code.AddMethod(Signature(Language.C_Error.Id(),C_any.Id()),1,MakeFunction1(E_c_code_Error,"c_code_Error")),MakeString("ocontrol.cl:232"))
  
  _ = Core.F_attach_method(C_Optimize_extendedTest_ask.AddMethod(Signature(Language.C_If.Id(),C_type.Id()),0,MakeFunction1(E_Optimize_extendedTest_ask_If,"Optimize_extendedTest?_If")),MakeString("ocontrol.cl:248"))
  
  _ = Core.F_attach_method(C_c_type.AddMethod(Signature(Language.C_If.Id(),C_type.Id()),1,MakeFunction1(E_c_type_If,"c_type_If")),MakeString("ocontrol.cl:255"))
  
  { 
    var expr EID
    { var _CL_obj *Core.GlobalVariable
      C_PENIBLE = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("PENIBLE",C_claire)))
      
      _CL_obj = C_PENIBLE
      _CL_obj.Range = ToType(C_boolean.Id())
      _CL_obj.Value = CFALSE.Id()
      expr = Core.F_close_global_variable(_CL_obj)
      } 
    ErrorCheck(expr)} 
  
  _ = Core.F_attach_method(C_c_code.AddMethod(Signature(Language.C_If.Id(),C_class.Id(),C_any.Id()),1,MakeFunction2(E_c_code_If,"c_code_If")),MakeString("ocontrol.cl:268"))
  
  _ = Core.F_attach_method(C_c_type.AddMethod(Signature(Language.C_Case.Id(),C_type.Id()),1,MakeFunction1(E_c_type_Case,"c_type_Case")),MakeString("ocontrol.cl:295"))
  
  _ = Core.F_attach_method(C_Optimize_case_branch.AddMethod(Signature(C_any.Id(),
    C_any.Id(),
    C_type.Id(),
    C_any.Id()),0,MakeFunction3(E_Optimize_case_branch_any,"Optimize_case_branch_any")),MakeString("ocontrol.cl:305"))
  
  _ = Core.F_attach_method(C_Optimize_case_substitution.AddMethod(Signature(C_any.Id(),
    C_Variable.Id(),
    C_Variable.Id(),
    C_any.Id()),0,MakeFunction3(E_Optimize_case_substitution_any,"Optimize_case_substitution_any")),MakeString("ocontrol.cl:312"))
  
  _ = Core.F_attach_method(C_c_code.AddMethod(Signature(Language.C_Case.Id(),C_class.Id(),C_any.Id()),1,MakeFunction2(E_c_code_Case,"c_code_Case")),MakeString("ocontrol.cl:346"))
  
  _ = Core.F_attach_method(C_c_type.AddMethod(Signature(Language.C_Do.Id(),C_type.Id()),1,MakeFunction1(E_c_type_Do,"c_type_Do")),MakeString("ocontrol.cl:353"))
  
  _ = Core.F_attach_method(C_c_code.AddMethod(Signature(Language.C_Do.Id(),C_class.Id(),C_any.Id()),1,MakeFunction2(E_c_code_Do,"c_code_Do")),MakeString("ocontrol.cl:357"))
  
  _ = Core.F_attach_method(C_c_type.AddMethod(Signature(Language.C_Let.Id(),C_type.Id()),1,MakeFunction1(E_c_type_Let,"c_type_Let")),MakeString("ocontrol.cl:365"))
  
  _ = Core.F_attach_method(C_c_code.AddMethod(Signature(Language.C_Let.Id(),C_class.Id(),C_any.Id()),1,MakeFunction2(E_c_code_Let,"c_code_Let")),MakeString("ocontrol.cl:378"))
  
  _ = Core.F_attach_method(C_c_type.AddMethod(Signature(Language.C_When.Id(),C_type.Id()),1,MakeFunction1(E_c_type_When,"c_type_When")),MakeString("ocontrol.cl:388"))
  
  _ = Core.F_attach_method(C_c_code.AddMethod(Signature(Language.C_When.Id(),C_class.Id(),C_any.Id()),1,MakeFunction2(E_c_code_When,"c_code_When")),MakeString("ocontrol.cl:420"))
  
  _ = Core.F_attach_method(C_c_type.AddMethod(Signature(Language.C_For.Id(),C_type.Id()),1,MakeFunction1(E_c_type_For,"c_type_For")),MakeString("ocontrol.cl:427"))
  
  _ = Core.F_attach_method(C_Optimize_infers_from.AddMethod(Signature(C_type.Id(),C_any.Id(),C_type.Id()),0,MakeFunction2(E_Optimize_infers_from_type,"Optimize_infers_from_type")),MakeString("ocontrol.cl:434"))
  
  _ = Core.F_attach_method(C_c_code.AddMethod(Signature(Language.C_For.Id(),C_class.Id(),C_any.Id()),1,MakeFunction2(E_c_code_For,"c_code_For")),MakeString("ocontrol.cl:479"))
  
  _ = Core.F_attach_method(C_Optimize_c_code_multiple.AddMethod(Signature(Language.C_For.Id(),
    C_type.Id(),
    C_class.Id(),
    C_any.Id()),1,MakeFunction3(E_Optimize_c_code_multiple_For,"Optimize_c_code_multiple_For")),MakeString("ocontrol.cl:503"))
  
  _ = Core.F_attach_method(C_c_type.AddMethod(Signature(Language.C_Iteration.Id(),C_type.Id()),1,MakeFunction1(E_c_type_Iteration,"c_type_Iteration")),MakeString("ocontrol.cl:514"))
  
  _ = Core.F_attach_method(C_c_code.AddMethod(Signature(Language.C_Iteration.Id(),C_any.Id()),1,MakeFunction1(E_c_code_Iteration,"c_code_Iteration")),MakeString("ocontrol.cl:557"))
  
  _ = Core.F_attach_method(C_c_code.AddMethod(Signature(Language.C_Select.Id(),C_any.Id()),1,MakeFunction1(E_c_code_Select,"c_code_Select")),MakeString("ocontrol.cl:561"))
  
  _ = Core.F_attach_method(C_c_code.AddMethod(Signature(Language.C_Lselect.Id(),C_any.Id()),1,MakeFunction1(E_c_code_Lselect,"c_code_Lselect")),MakeString("ocontrol.cl:562"))
  
  _ = Core.F_attach_method(C_Optimize_c_code_select.AddMethod(Signature(Language.C_Iteration.Id(),C_class.Id(),C_any.Id()),1,MakeFunction2(E_Optimize_c_code_select_Iteration,"Optimize_c_code_select_Iteration")),MakeString("ocontrol.cl:588"))
  
  _ = Core.F_attach_method(C_Optimize_inner_select.AddMethod(Signature(Language.C_Iteration.Id(),
    C_any.Id(),
    C_any.Id(),
    C_any.Id(),
    C_any.Id()),1,MakeFunction4(E_Optimize_inner_select_Iteration,"Optimize_inner_select_Iteration")),MakeString("ocontrol.cl:601"))
  
  _ = Core.F_attach_method(C_c_type.AddMethod(Signature(Language.C_Exists.Id(),C_type.Id()),1,MakeFunction1(E_c_type_Exists,"c_type_Exists")),MakeString("ocontrol.cl:608"))
  
  _ = Core.F_attach_method(C_c_code.AddMethod(Signature(Language.C_Exists.Id(),C_class.Id(),C_any.Id()),1,MakeFunction2(E_c_code_Exists,"c_code_Exists")),MakeString("ocontrol.cl:635"))
  
  _ = Core.F_attach_method(C_c_type.AddMethod(Signature(Language.C_Image.Id(),C_type.Id()),1,MakeFunction1(E_c_type_Image,"c_type_Image")),MakeString("ocontrol.cl:642"))
  
  _ = Core.F_attach_method(C_c_type.AddMethod(Signature(Language.C_Select.Id(),C_type.Id()),1,MakeFunction1(E_c_type_Select,"c_type_Select")),MakeString("ocontrol.cl:648"))
  
  _ = Core.F_attach_method(C_c_type.AddMethod(Signature(Language.C_Lselect.Id(),C_type.Id()),1,MakeFunction1(E_c_type_Lselect,"c_type_Lselect")),MakeString("ocontrol.cl:655"))
  
  _ = Core.F_attach_method(C_c_type.AddMethod(Signature(Language.C_While.Id(),C_type.Id()),1,MakeFunction1(E_c_type_While,"c_type_While")),MakeString("ocontrol.cl:660"))
  
  _ = Core.F_attach_method(C_c_code.AddMethod(Signature(Language.C_While.Id(),C_class.Id(),C_any.Id()),1,MakeFunction2(E_c_code_While,"c_code_While")),MakeString("ocontrol.cl:665"))
  
  _ = Core.F_attach_method(C_Optimize_Iterate_I.AddMethod(Signature(Language.C_Iteration.Id(),C_any.Id()),0,MakeFunction1(E_Optimize_Iterate_I_Iteration,"Optimize_Iterate!_Iteration")),MakeString("ocontrol.cl:681"))
  
  _ = Core.F_inlineok_ask_method(
  Core.F_attach_method(Language.C_iterate.AddMethod(Signature(C_Interval.Id(),
    Core.F_nth_class2(C_Variable,MakeList(ToType(C_any.Id()),C_range.Id()),MakeConstantList(Core.F_nth_class1(C_type,ToType(C_integer.Id())).Id())).Id(),
    C_any.Id(),
    C_any.Id()),1,MakeFunction3(E_iterate_Interval,"iterate_Interval")),MakeString("ocontrol.cl:689")),
  MakeString("lambda[(x:Interval,v:Variable[range:(subtype[integer])],e:any),let v := eval(x.arg1, Interval),%max:integer := eval(x.arg2, Interval) in while (v <= %max) (e, v := v + 1)]"))
  
  _ = Core.F_inlineok_ask_method(
  Core.F_attach_method(Language.C_iterate.AddMethod(Signature(C_array.Id(),
    C_Variable.Id(),
    C_any.Id(),
    C_any.Id()),1,MakeFunction3(E_iterate_array,"iterate_array")),MakeString("ocontrol.cl:694")),
  MakeString("lambda[(x:array,v:Variable,e:any),let %i := 1,%a := x,%max := length(%a) in while (%i <= %max) let v := %a[%i] in (e, %i := %i + 1)]"))
  
  _ = Core.F_inlineok_ask_method(
  Core.F_attach_method(Language.C_Iterate.AddMethod(Signature(C_class.Id(),
    C_Variable.Id(),
    C_any.Id(),
    C_any.Id()),0,MakeFunction3(E_Iterate_class,"Iterate_class")),MakeString("ocontrol.cl:698")),
  MakeString("lambda[(x:class,v:Variable,e:any),for %v_1 in x.descendants let %v_2 := (for v in %v_1.instances e) in (if %v_2 break(%v_2) else false)]"))
  
  _ = Core.F_inlineok_ask_method(
  Core.F_attach_method(Language.C_Iterate.AddMethod(Signature(F_nth_property(ToProperty(C__dot_dot.Id()),MakeTuple(C_integer.Id(),C_integer.Id())).Id(),
    C_Variable.Id(),
    C_any.Id(),
    C_any.Id()),1,MakeFunction3(E_Iterate_any1,"Iterate_any1")),MakeString("ocontrol.cl:703")),
  MakeString("lambda[(x:any,v:Variable,e:any),let v := eval(x.args[1]),%max := eval(x.args[2]) in while (v <= %max) (e, v := v + 1)]"))
  
  _ = Core.F_inlineok_ask_method(
  Core.F_attach_method(Language.C_Iterate.AddMethod(Signature(Language.C_Lselect.Id(),
    C_Variable.Id(),
    C_any.Id(),
    C_any.Id()),1,MakeFunction3(E_Iterate_Lselect,"Iterate_Lselect")),MakeString("ocontrol.cl:706")),
  MakeString("lambda[(x:Lselect,v:Variable,e:any),for v in eval(x.iClaire/set_arg) (if eval(substitution(x.arg, x.var, v)) e else false)]"))
  
  _ = Core.F_inlineok_ask_method(
  Core.F_attach_method(Language.C_Iterate.AddMethod(Signature(Language.C_Select.Id(),
    C_Variable.Id(),
    C_any.Id(),
    C_any.Id()),1,MakeFunction3(E_Iterate_Select,"Iterate_Select")),MakeString("ocontrol.cl:709")),
  MakeString("lambda[(x:Select,v:Variable,e:any),for v in eval(x.iClaire/set_arg) (if eval(substitution(x.arg, x.var, v)) e else false)]"))
  
  _ = Core.F_inlineok_ask_method(
  Core.F_attach_method(Language.C_Iterate.AddMethod(Signature(Language.C_Collect.Id(),
    C_Variable.Id(),
    C_any.Id(),
    C_any.Id()),1,MakeFunction3(E_Iterate_Collect,"Iterate_Collect")),MakeString("ocontrol.cl:713")),
  MakeString("lambda[(x:Collect,v:Variable,e:any),for C%v in eval(x.iClaire/set_arg) let v := eval(substitution(x.arg, x.var, C%v)) in e]"))
  
  _ = Core.F_inlineok_ask_method(
  Core.F_attach_method(Language.C_Iterate.AddMethod(Signature(F_nth_property(ToProperty(Core.C_but.Id()),MakeTuple(C_any.Id(),C_any.Id())).Id(),
    C_Variable.Id(),
    C_any.Id(),
    C_any.Id()),1,MakeFunction3(E_Iterate_any2,"Iterate_any2")),MakeString("ocontrol.cl:716")),
  MakeString("lambda[(x:any,v:Variable,e:any),for v in eval(x.args[1]) (if (v != eval(x.args[2])) e else false)]"))
  
  _ = Core.F_inlineok_ask_method(
  Core.F_attach_method(Language.C_Iterate.AddMethod(Signature(F_nth_property(ToProperty(C__7_plus.Id()),MakeTuple(C_any.Id(),C_any.Id())).Id(),
    C_Variable.Id(),
    C_any.Id(),
    C_any.Id()),1,MakeFunction3(E_Iterate_any3,"Iterate_any3")),MakeString("ocontrol.cl:719")),
  MakeString("lambda[(x:any,v:Variable,e:any),(for v in eval(x.args[1]) e, for v in eval(x.args[2]) e)]"))
  
  _ = Core.F_attach_method(C_c_type.AddMethod(Signature(Language.C_List.Id(),C_type.Id()),1,MakeFunction1(E_c_type_List,"c_type_List")),MakeString("odefine.cl:28"))
  
  _ = Core.F_attach_method(C_c_code.AddMethod(Signature(Language.C_List.Id(),C_any.Id()),1,MakeFunction1(E_c_code_List,"c_code_List")),MakeString("odefine.cl:41"))
  
  _ = Core.F_attach_method(C_c_type.AddMethod(Signature(Language.C_Set.Id(),C_type.Id()),1,MakeFunction1(E_c_type_Set,"c_type_Set")),MakeString("odefine.cl:51"))
  
  _ = Core.F_attach_method(C_c_code.AddMethod(Signature(Language.C_Set.Id(),C_any.Id()),1,MakeFunction1(E_c_code_Set,"c_code_Set")),MakeString("odefine.cl:63"))
  
  _ = Core.F_attach_method(C_c_type.AddMethod(Signature(Language.C_Tuple.Id(),C_type.Id()),1,MakeFunction1(E_c_type_Tuple,"c_type_Tuple")),MakeString("odefine.cl:66"))
  
  _ = Core.F_attach_method(C_c_code.AddMethod(Signature(Language.C_Tuple.Id(),C_any.Id()),1,MakeFunction1(E_c_code_Tuple,"c_code_Tuple")),MakeString("odefine.cl:69"))
  
  _ = Core.F_attach_method(C_c_type.AddMethod(Signature(Language.C_Map.Id(),C_type.Id()),0,MakeFunction1(E_c_type_Map,"c_type_Map")),MakeString("odefine.cl:72"))
  
  _ = Core.F_attach_method(C_c_code.AddMethod(Signature(Language.C_Map.Id(),C_any.Id()),1,MakeFunction1(E_c_code_Map,"c_code_Map")),MakeString("odefine.cl:81"))
  
  _ = Core.F_attach_method(C_c_type.AddMethod(Signature(Language.C_Definition.Id(),C_type.Id()),0,MakeFunction1(E_c_type_Definition,"c_type_Definition")),MakeString("odefine.cl:88"))
  
  { 
    var expr EID
    { var _CL_obj *Core.GlobalVariable
      C_Compile__starname_star = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("*name*",C_Compile)))
      
      _CL_obj = C_Compile__starname_star
      _CL_obj.Range = ToType(C_symbol.Id())
      _CL_obj.Value = Core.F_symbol_I_string2(MakeString("_CL_obj")).Id()
      expr = Core.F_close_global_variable(_CL_obj)
      } 
    ErrorCheck(expr)} 
  
  _ = Core.F_attach_method(C_c_code.AddMethod(Signature(Language.C_Definition.Id(),C_class.Id(),C_any.Id()),1,MakeFunction2(E_c_code_Definition,"c_code_Definition")),MakeString("odefine.cl:103"))
  
  _ = Core.F_attach_method(C_Optimize_total_ask.AddMethod(Signature(C_class.Id(),C_list.Id(),C_any.Id()),1,MakeFunction2(E_Optimize_total_ask_class,"Optimize_total?_class")),MakeString("odefine.cl:120"))
  
  _ = Core.F_attach_method(C_Optimize_analyze_I.AddMethod(Signature(C_class.Id(),
    C_any.Id(),
    C_list.Id(),
    C_list.Id(),
    C_any.Id()),1,MakeFunction4(E_Optimize_analyze_I_class,"Optimize_analyze!_class")),MakeString("odefine.cl:151"))
  
  _ = Core.F_attach_method(C_c_code.AddMethod(Signature(Language.C_Defobj.Id(),C_class.Id(),C_any.Id()),1,MakeFunction2(E_c_code_Defobj,"c_code_Defobj")),MakeString("odefine.cl:173"))
  
  _ = Core.F_attach_method(C_c_code.AddMethod(Signature(Language.C_Defclass.Id(),C_class.Id(),C_any.Id()),1,MakeFunction2(E_c_code_Defclass,"c_code_Defclass")),MakeString("odefine.cl:192"))
  
  _ = Core.F_attach_method(C_c_type.AddMethod(Signature(Language.C_Defmethod.Id(),C_type.Id()),0,MakeFunction1(E_c_type_Defmethod,"c_type_Defmethod")),MakeString("odefine.cl:198"))
  
  _ = Core.F_attach_method(C_c_code.AddMethod(Signature(Language.C_Defmethod.Id(),C_any.Id()),1,MakeFunction1(E_c_code_Defmethod,"c_code_Defmethod")),MakeString("odefine.cl:237"))
  
  _ = Core.F_attach_method(C_Optimize_type_extension.AddMethod(Signature(C_string.Id(),C_string.Id()),0,MakeFunction1(E_Optimize_type_extension_string,"Optimize_type_extension_string")),MakeString("odefine.cl:243"))
  
  _ = Core.F_attach_method(C_Optimize_sort_pattern_ask.AddMethod(Signature(C_list.Id(),C_any.Id(),C_boolean.Id()),1,MakeFunction2(E_Optimize_sort_pattern_ask_list,"Optimize_sort_pattern?_list")),MakeString("odefine.cl:253"))
  
  _ = Core.F_attach_method(C_Optimize_sort_code.AddMethod(Signature(Language.C_Defmethod.Id(),C_list.Id(),C_any.Id()),1,MakeFunction2(E_Optimize_sort_code_Defmethod,"Optimize_sort_code_Defmethod")),MakeString("odefine.cl:298"))
  
  _ = Core.F_attach_method(C_add_method.AddMethod(Signature(C_property.Id(),
    C_list.Id(),
    C_type.Id(),
    C_integer.Id(),
    C_function.Id(),
    C_method.Id(),
    C_method.Id()),0,MakeFunction6(E_add_method_property2,"add_method_property2")),MakeString("odefine.cl:305"))
  
  { 
    var expr EID
    { var _CL_obj *Core.GlobalVariable
      C_Optimize__staradd_method2_star = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("*add_method2*",It)))
      
      _CL_obj = C_Optimize__staradd_method2_star
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = C_add_method.Restrictions.ValuesO()[1]
      expr = Core.F_close_global_variable(_CL_obj)
      } 
    ErrorCheck(expr)} 
  
  _ = Core.F_attach_method(C_Optimize_add_method_I.AddMethod(Signature(C_method.Id(),
    C_list.Id(),
    C_any.Id(),
    C_any.Id(),
    C_function.Id(),
    C_any.Id()),1,MakeFunction5(E_Optimize_add_method_I_method,"Optimize_add_method!_method")),MakeString("odefine.cl:320"))
  
  _ = Core.F_attach_method(C_Optimize_extract_signature_I.AddMethod(Signature(C_list.Id(),C_list.Id()),1,MakeFunction1(E_Optimize_extract_signature_I_list,"Optimize_extract_signature!_list")),MakeString("odefine.cl:336"))
  
  C_Optimize__equalsig_ask = ToOperation(new(ClaireOperation).IsNamed(C_operation,MakeSymbol("=sig?",It)))
  
  
  _ = Core.F_attach_method(C_Optimize__equalsig_ask.AddMethod(Signature(C_list.Id(),C_list.Id(),C_boolean.Id()),0,MakeFunction2(E_Optimize__equalsig_ask_list,"Optimize_=sig?_list")),MakeString("odefine.cl:340"))
  
  _ = Core.F_attach_method(C_Compile_function_name.AddMethod(Signature(C_property.Id(),
    C_list.Id(),
    C_any.Id(),
    C_string.Id()),0,MakeFunction3(E_Compile_function_name_property1,"Compile_function_name_property1")),MakeString("odefine.cl:359"))
  
  _ = Core.F_attach_method(C_Compile_compile_lambda.AddMethod(Signature(C_string.Id(),
    C_lambda.Id(),
    C_any.Id(),
    C_any.Id()),1,MakeFunction3(E_Compile_compile_lambda_string,"Compile_compile_lambda_string")),MakeString("odefine.cl:376"))
  
  _ = Core.F_attach_method(C_c_code.AddMethod(Signature(Language.C_Defarray.Id(),C_any.Id()),1,MakeFunction1(E_c_code_Defarray,"c_code_Defarray")),MakeString("odefine.cl:424"))
  
  _ = Core.F_attach_method(C_Compile_compute_if_write_inverse.AddMethod(Signature(C_relation.Id(),C_void.Id()),1,MakeFunction1(E_Compile_compute_if_write_inverse_relation,"Compile_compute_if_write_inverse_relation")),MakeString("odefine.cl:458"))
  
  _ = Core.F_attach_method(C_Compile_compute_set_write.AddMethod(Signature(C_relation.Id(),C_any.Id()),1,MakeFunction1(E_Compile_compute_set_write_relation,"Compile_compute_set_write_relation")),MakeString("odefine.cl:474"))
  
  _ = Core.F_attach_method(C_Optimize_Produce_put.AddMethod(Signature(C_property.Id(),
    C_Variable.Id(),
    C_any.Id(),
    C_any.Id()),1,MakeFunction3(E_Optimize_Produce_put_property,"Optimize_Produce_put_property")),MakeString("odefine.cl:486"))
  
  _ = Core.F_attach_method(C_Optimize_Produce_erase.AddMethod(Signature(C_property.Id(),C_Variable.Id(),C_any.Id()),1,MakeFunction2(E_Optimize_Produce_erase_property,"Optimize_Produce_erase_property")),MakeString("odefine.cl:499"))
  
  _ = Core.F_attach_method(C_Optimize_Produce_put.AddMethod(Signature(C_table.Id(),
    C_Variable.Id(),
    C_any.Id(),
    C_any.Id()),0,MakeFunction3(E_Optimize_Produce_put_table,"Optimize_Produce_put_table")),MakeString("odefine.cl:508"))
  
  _ = Core.F_attach_method(C_Optimize_Produce_get.AddMethod(Signature(C_relation.Id(),C_Variable.Id(),C_any.Id()),1,MakeFunction2(E_Optimize_Produce_get_relation,"Optimize_Produce_get_relation")),MakeString("odefine.cl:519"))
  
  _ = Core.F_attach_method(C_Optimize_Produce_remove.AddMethod(Signature(C_property.Id(),
    C_Variable.Id(),
    C_any.Id(),
    C_any.Id()),1,MakeFunction3(E_Optimize_Produce_remove_property,"Optimize_Produce_remove_property")),MakeString("odefine.cl:529"))
  
  _ = Core.F_attach_method(C_Optimize_Produce_remove.AddMethod(Signature(C_table.Id(),
    C_Variable.Id(),
    C_any.Id(),
    C_any.Id()),0,MakeFunction3(E_Optimize_Produce_remove_table,"Optimize_Produce_remove_table")),MakeString("odefine.cl:535"))
  
  _ = Core.F_attach_method(C_Optimize_Tighten.AddMethod(Signature(C_relation.Id(),C_void.Id()),0,MakeFunction1(E_Optimize_Tighten_relation,"Optimize_Tighten_relation")),MakeString("odefine.cl:549"))
  
  _ = Core.F_attach_method(C_Compile_Tighten_I.AddMethod(Signature(C_relation.Id(),C_void.Id()),0,MakeFunction1(E_Compile_Tighten_I_relation,"Compile_Tighten!_relation")),MakeString("odefine.cl:551"))
  
  _ = Core.F_attach_method(C_Compile_lexical_num.AddMethod(Signature(C_any.Id(),C_integer.Id(),C_void.Id()),1,MakeFunction2(E_Compile_lexical_num_any,"Compile_lexical_num_any")),MakeString("odefine.cl:566"))
  
  _ = Core.F_attach_method(C_c_type.AddMethod(Signature(Language.C_Defrule.Id(),C_type.Id()),0,MakeFunction1(E_c_type_Defrule,"c_type_Defrule")),MakeString("odefine.cl:571"))
  
  _ = Core.F_attach_method(C_c_code.AddMethod(Signature(Language.C_Defrule.Id(),C_class.Id(),C_any.Id()),1,MakeFunction2(E_c_code_Defrule,"c_code_Defrule")),MakeString("odefine.cl:591"))
  
  _ = Core.F_attach_method(C_Optimize_compile_if_write.AddMethod(Signature(C_relation.Id(),C_void.Id()),1,MakeFunction1(E_Optimize_compile_if_write_relation,"Optimize_compile_if_write_relation")),MakeString("odefine.cl:620"))
  
  _ = Core.F_attach_method(C_Optimize_demon_substitution.AddMethod(Signature(Language.C_Language_demon.Id(),C_list.Id(),C_any.Id()),0,MakeFunction2(E_Optimize_demon_substitution_demon,"Optimize_demon_substitution_demon")),MakeString("odefine.cl:625"))
  
  _ = Core.F_attach_method(C_Optimize_demon_function.AddMethod(Signature(C_string.Id(),C_function.Id()),0,MakeFunction1(E_Optimize_demon_function_string,"Optimize_demon_function_string")),MakeString("odefine.cl:631"))
  
  _ = Core.F_attach_method(C_Optimize_compileEventMethod.AddMethod(Signature(C_property.Id(),C_any.Id()),1,MakeFunction1(E_Optimize_compileEventMethod_property,"Optimize_compileEventMethod_property")),MakeString("odefine.cl:637"))
  
  } 

