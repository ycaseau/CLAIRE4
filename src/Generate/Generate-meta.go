/***** CLAIRE Compilation of module Generate.cl 
         [version 4.0.04 / safety 5] Sunday 03-13-2022 07:28:45 *****/

package Generate
import (_ "fmt"
	"unsafe"
	. "Kernel"
	"Core"
	"Language"
	"Reader"
	"Optimize"
)

//-------- dumb function to prevent import errors --------
func import_g0209() { 
    _ = Core.It
    _ = Language.It
    _ = Reader.It
    _ = Optimize.It
    } 
  
  
// class file for Generate/code_producer in module Generate 
type GenerateCodeProducer struct { 
   Optimize.CompileProducer
   OpenComparators *ClaireList 
  OpenOperators *ClaireList 
  DivOperators *ClaireList 
  Body *ClaireAny 
  Extension *ClaireString 
  Comment *ClaireString 
  Interfaces *ClaireList 
  Stat int
  } 

// automatic cast function
func ToGenerateCodeProducer(x *ClaireAny) *GenerateCodeProducer {return (*GenerateCodeProducer)(unsafe.Pointer(x))}

// class file for Generate/go_producer in module Generate 
type GenerateGoProducer struct { 
   GenerateCodeProducer
   Current *ClaireModule 
  BadNames *ClaireList 
  GoodNames *ClaireList 
  KernelMethods *ClaireList 
  Source *ClaireString 
  Debug_ask *ClaireBoolean 
  Varsym int
  } 

// automatic cast function
func ToGenerateGoProducer(x *ClaireAny) *GenerateGoProducer {return (*GenerateGoProducer)(unsafe.Pointer(x))}

var C_Generate__star_ask_interval_star *Core.GlobalVariable 
var C_Generate__star_dash_dash_integer_star *Core.GlobalVariable 
var C_Generate__star_plus_integer_star *Core.GlobalVariable 
var C_Generate__starnth_integer_star *Core.GlobalVariable 
var C_Generate__starnth_list_star *Core.GlobalVariable 
var C_Generate__starnth_tuple_star *Core.GlobalVariable 
var C_Generate__starnth_1_list_star *Core.GlobalVariable 
var C_Generate__starnth_1_tuple_star *Core.GlobalVariable 
var C_Generate__starnth_1_array_star *Core.GlobalVariable 
var C_Generate__starnth_string_star *Core.GlobalVariable 
var C_Generate__starnth_1_string_star *Core.GlobalVariable 
var C_Generate__starnth_equal_list_star *Core.GlobalVariable 
var C_Generate__starnth_put_list_star *Core.GlobalVariable 
var C_Generate__starmake_list_star *Core.GlobalVariable 
var C_Generate__starnot_star *Core.GlobalVariable 
var C_Generate__starknown_star *Core.GlobalVariable 
var C_Generate__starunknown_star *Core.GlobalVariable 
var C_Generate__starnot_equal_star *Core.GlobalVariable 
var C_Generate__starequal_star *Core.GlobalVariable 
var C_Generate__starbelong_star *Core.GlobalVariable 
var C_Generate__starcontain_list_star *Core.GlobalVariable 
var C_Generate__starcontain_set_star *Core.GlobalVariable 
var C_Generate__starlength_array_star *Core.GlobalVariable 
var C_Generate__starlength_bag_star *Core.GlobalVariable 
var C_Generate__starclose_exception_star *Core.GlobalVariable 
var C_Generate__starnew_class1_star *Core.GlobalVariable 
var C_Generate__starnew_class2_star *Core.GlobalVariable 
var C_Generate__starslot_get_star *Core.GlobalVariable 
var C_Generate__starmap_star *Core.GlobalVariable 
var C_Generate__starof_bag_star *Core.GlobalVariable 
var C_Generate__starof_array_star *Core.GlobalVariable 
var C_Generate__starcopy_list_star *Core.GlobalVariable 
var C_Generate__starcopy_set_star *Core.GlobalVariable 
var C_Generate__starempty_set_star *Core.GlobalVariable 
var C_Generate__starnth_put_array_star *Core.GlobalVariable 
var C_Generate__star_Zt_star *Core.GlobalVariable 
var C_Generate__starlesst_star *Core.GlobalVariable 
var C_Generate__starincluded_star *Core.GlobalVariable 
var C_Generate__starstack_apply_star *Core.GlobalVariable 
var C_Generate__starsuper_apply_star *Core.GlobalVariable 
var C_Generate__starprinc_string_star *Core.GlobalVariable 
var C_Generate__starinherit_star *Core.GlobalVariable 
var C_Generate__starwrite_value_star *Core.GlobalVariable 
var C_Generate__starread_property_star *Core.GlobalVariable 
var C_Generate_code_producer *ClaireClass  /*obj*/
var C_Generate_go_producer *ClaireClass  /*obj*/
var C_Generate_EIDSET *Core.GlobalVariable 
var C_BadMethods *Core.GlobalVariable 
var C_ABODY *Core.GlobalVariable 
var C_Generate_GO_PRODUCER *GenerateGoProducer  /*obj*/
var C_Generate__starlength_string_star *Core.GlobalVariable 
var C_Generate__starset_I_list_star *Core.GlobalVariable 
var C_Generate_open_comparators *ClaireProperty  // Generate/"open_comparators"
var C_Generate_open_operators *ClaireProperty  // Generate/"open_operators"
var C_Generate_div_operators *ClaireProperty  // Generate/"div_operators"
var C_Generate_extension *ClaireProperty  // Generate/"extension"
var C_Generate_interfaces *ClaireProperty  // Generate/"interfaces"
var C_Generate_current *ClaireProperty  // Generate/"current"
var C_Generate_bad_names *ClaireProperty  // Generate/"bad_names"
var C_Generate_good_names *ClaireProperty  // Generate/"good_names"
var C_Generate_kernel_methods *ClaireProperty  // Generate/"kernel_methods"
var C_Generate_varsym *ClaireProperty  // Generate/"varsym"
var C_Generate_indent_c *ClaireProperty  // Generate/"indent_c"
var C_Generate_breakline *ClaireProperty  // Generate/"breakline"
var C_Generate_new_block *ClaireProperty  // Generate/"new_block"
var C_Generate_let_block *ClaireProperty  // Generate/"let_block"
var C_Generate_close_block *ClaireProperty  // Generate/"close_block"
var C_Generate_finish_block *ClaireProperty  // Generate/"finish_block"
var C_g_test *ClaireProperty  // claire/"g_test"
var C_Generate_g_func *ClaireProperty  // Generate/"g_func"
var C_Generate_g_expression *ClaireProperty  // Generate/"g_expression"
var C_Generate_statement *ClaireProperty  // Generate/"statement"
var C_Generate_gtop *ClaireProperty  // Generate/"gtop"
var C_Generate_make_go_function *ClaireProperty  // Generate/"make_go_function"
var C_compile *ClaireProperty  // claire/"compile"
var C_Generate_parents *ClaireProperty  // Generate/"parents"
var C_Generate_gen_files *ClaireProperty  // Generate/"gen_files"
var C_Generate_gen_mod_file *ClaireProperty  // Generate/"gen_mod_file"
var C_Generate_gen_file *ClaireProperty  // Generate/"gen_file"
var C_Generate_start_file *ClaireProperty  // Generate/"start_file"
var C_Generate_gen_classes *ClaireProperty  // Generate/"gen_classes"
var C_Generate_gen_objects *ClaireProperty  // Generate/"gen_objects"
var C_Generate_gen_meta_load *ClaireProperty  // Generate/"gen_meta_load"
var C_Generate_namespace_I *ClaireProperty  // Generate/"namespace!"
var C_Generate_import_declaration *ClaireProperty  // Generate/"import_declaration"
var C_Generate_dumb_import *ClaireProperty  // Generate/"dumb_import"
var C_Generate_needed_modules *ClaireProperty  // Generate/"needed_modules"
var C_Generate_cap_short *ClaireProperty  // Generate/"cap_short"
var C_Generate_representative *ClaireProperty  // Generate/"representative"
var C_Generate_clean_duplicates *ClaireProperty  // Generate/"clean_duplicates"
var C_Generate_gen_class_def *ClaireProperty  // Generate/"gen_class_def"
var C_Generate_gen_cast_function *ClaireProperty  // Generate/"gen_cast_function"
var C_Generate_construct_class_ask *ClaireProperty  // Generate/"construct_class?"
var C_Generate_gen_construct *ClaireProperty  // Generate/"gen_construct"
var C_Generate_rootSlot *ClaireProperty  // Generate/"rootSlot"
var C_Generate_go_class *ClaireProperty  // Generate/"go_class"
var C_Generate_interface_I *ClaireProperty  // Generate/"interface!"
var C_Generate_cast_class *ClaireProperty  // Generate/"cast_class"
var C_Generate_add_underscore *ClaireProperty  // Generate/"add_underscore"
var C_Generate_class_ident *ClaireProperty  // Generate/"class_ident"
var C_Generate_cast_prefix *ClaireProperty  // Generate/"cast_prefix"
var C_Generate_cast_post *ClaireProperty  // Generate/"cast_post"
var C_Generate_go_var *ClaireProperty  // Generate/"go_var"
var C_Generate_getRange *ClaireProperty  // Generate/"getRange"
var C_Generate_thing_ident *ClaireProperty  // Generate/"thing_ident"
var C_Generate_gen_module *ClaireProperty  // Generate/"gen_module"
var C_Generate_declare *ClaireProperty  // Generate/"declare"
var C_Generate_var_declaration *ClaireProperty  // Generate/"var_declaration"
var C_Generate_g_statement *ClaireProperty  // Generate/"g_statement"
var C_Generate_simple_func_ask *ClaireProperty  // Generate/"simple_func?"
var C_Generate_get_made *ClaireProperty  // Generate/"get_made"
var C_Generate_fileName *ClaireProperty  // Generate/"fileName"
var C_Generate_make_lambda_function *ClaireProperty  // Generate/"make_lambda_function"
var C_Generate_generate_function_start *ClaireProperty  // Generate/"generate_function_start"
var C_Generate_eid_body *ClaireProperty  // Generate/"eid_body"
var C_Generate_generate_eid_dual *ClaireProperty  // Generate/"generate_eid_dual"
var C_Generate_goMethod_ask *ClaireProperty  // Generate/"goMethod?"
var C_Generate_goVariable *ClaireProperty  // Generate/"goVariable"
var C_Generate_goMethod *ClaireProperty  // Generate/"goMethod"
var C_Generate_goVariables *ClaireProperty  // Generate/"goVariables"
var C_Generate_goFunction *ClaireProperty  // Generate/"goFunction"
var C_Generate_check_range *ClaireProperty  // Generate/"check_range"
var C_Generate_simple_body_ask *ClaireProperty  // Generate/"simple_body?"
var C_Generate_need_debug_ask *ClaireProperty  // Generate/"need_debug?"
var C_Generate_procedure_body *ClaireProperty  // Generate/"procedure_body"
var C_Generate_function_body *ClaireProperty  // Generate/"function_body"
var C_Generate_generate_eid_function *ClaireProperty  // Generate/"generate_eid_function"
var C_Generate_generate_eval_function *ClaireProperty  // Generate/"generate_eval_function"
var C_Generate_debug_intro *ClaireProperty  // Generate/"debug_intro"
var C_Generate_return_result *ClaireProperty  // Generate/"return_result"
var C_Generate_go_signature *ClaireProperty  // Generate/"go_signature"
var C_Generate_goEIDFunctionName *ClaireProperty  // Generate/"goEIDFunctionName"
var C_Generate_goEIDVariables *ClaireProperty  // Generate/"goEIDVariables"
var C_Generate_print_EID_call *ClaireProperty  // Generate/"print_EID_call"
var C_Generate_external_EID_arg *ClaireProperty  // Generate/"external_EID_arg"
var C_Generate_eid_prefix *ClaireProperty  // Generate/"eid_prefix"
var C_Generate_eid_post *ClaireProperty  // Generate/"eid_post"
var C_Generate_to_eid *ClaireProperty  // Generate/"to_eid"
var C_Generate_build_Variable *ClaireProperty  // Generate/"build_Variable"
var C_Generate_import_princ *ClaireProperty  // Generate/"import_princ"
var C_Generate_genvar *ClaireProperty  // Generate/"genvar"
var C_Generate_c_string *ClaireProperty  // Generate/"c_string"
var C_Generate_cap_ident *ClaireProperty  // Generate/"cap_ident"
var C_Generate_capitalized_ident *ClaireProperty  // Generate/"capitalized_ident"
var C_Generate_capitalize *ClaireProperty  // Generate/"capitalize"
var C_Generate_symbol_ident *ClaireProperty  // Generate/"symbol_ident"
var C_Generate_globalVar *ClaireProperty  // Generate/"globalVar"
var C_Generate_type_sort *ClaireProperty  // Generate/"type_sort"
var C_Generate_g_sort *ClaireProperty  // Generate/"g_sort"
var C_Generate_valuesSlot *ClaireProperty  // Generate/"valuesSlot"
var C_imported_function_ask *ClaireProperty  // claire/"imported_function?"
var C_Generate_arg_match *ClaireProperty  // Generate/"arg_match"
var C_dMethod_ask *ClaireProperty  // claire/"dMethod?"
var C_Generate_at *ClaireProperty  // Generate/"at"
var C_Generate_getFunctionName *ClaireProperty  // Generate/"getFunctionName"
var C_Generate_preCore_ask *ClaireProperty  // Generate/"preCore?"
var C_Generate_goEIDFunction *ClaireProperty  // Generate/"goEIDFunction"
var C_Generate_goEvalFunction *ClaireProperty  // Generate/"goEvalFunction"
var C_Generate_retreive_method *ClaireProperty  // Generate/"retreive_method"
var C_Generate_retreive_list *ClaireProperty  // Generate/"retreive_list"
var C_Generate_to_cl *ClaireProperty  // Generate/"to_cl"
var C_Generate_from_eid *ClaireProperty  // Generate/"from_eid"
var C_Generate_integer_prefix *ClaireProperty  // Generate/"integer_prefix"
var C_Generate_float_prefix *ClaireProperty  // Generate/"float_prefix"
var C_Generate_char_prefix *ClaireProperty  // Generate/"char_prefix"
var C_Generate_string_prefix *ClaireProperty  // Generate/"string_prefix"
var C_Generate_string_post *ClaireProperty  // Generate/"string_post"
var C_Generate_native_post *ClaireProperty  // Generate/"native_post"
var C_Generate_object_prefix *ClaireProperty  // Generate/"object_prefix"
var C_Generate_object_post *ClaireProperty  // Generate/"object_post"
var C_Generate_equal_exp *ClaireProperty  // Generate/"equal_exp"
var C_Generate_sign_equal *ClaireProperty  // Generate/"sign_equal"
var C_Generate_char_exp_ask *ClaireProperty  // Generate/"char_exp?"
var C_Generate_c_member *ClaireProperty  // Generate/"c_member"
var C_Generate_bag_expression *ClaireProperty  // Generate/"bag_expression"
var C_Generate_args_list *ClaireProperty  // Generate/"args_list"
var C_Generate_constant_ask *ClaireProperty  // Generate/"constant?"
var C_Generate_go_range *ClaireProperty  // Generate/"go_range"
var C_Generate_full_signature *ClaireProperty  // Generate/"full_signature"
var C_Generate_signature_I *ClaireProperty  // Generate/"signature!"
var C_Generate_g_member *ClaireProperty  // Generate/"g_member"
var C_Generate_cast_Values *ClaireProperty  // Generate/"cast_Values"
var C_Generate_check_var *ClaireProperty  // Generate/"check_var"
var C_Generate_use_variable *ClaireProperty  // Generate/"use_variable"
var C_Generate_g_clean *ClaireProperty  // Generate/"g_clean"
var C_Generate_inline_exp *ClaireProperty  // Generate/"inline_exp"
var C_Generate_bounded_expression *ClaireProperty  // Generate/"bounded_expression"
var C_Generate_eid_provide_ask *ClaireProperty  // Generate/"eid_provide?"
var C_Generate_print_external_call *ClaireProperty  // Generate/"print_external_call"
var C_Generate_belong_exp *ClaireProperty  // Generate/"belong_exp"
var C_Generate_external_casted_arg *ClaireProperty  // Generate/"external_casted_arg"
var C_Generate_g_table_index *ClaireProperty  // Generate/"g_table_index"
var C_Generate_sign_or *ClaireProperty  // Generate/"sign_or"
var C_Generate_unfold_args *ClaireProperty  // Generate/"unfold_args"
var C_Generate_unfold_arg *ClaireProperty  // Generate/"unfold_arg"
var C_Generate_unfold_use *ClaireProperty  // Generate/"unfold_use"
var C_Generate_eid_require_ask *ClaireProperty  // Generate/"eid_require?"
var C_Generate_unfold_eid *ClaireProperty  // Generate/"unfold_eid"
var C_Generate_g_try *ClaireProperty  // Generate/"g_try"
var C_Generate_close_try *ClaireProperty  // Generate/"close_try"
var C_Generate_error_wrap *ClaireProperty  // Generate/"error_wrap"
var C_Generate_g_try_void *ClaireProperty  // Generate/"g_try_void"
var C_Generate_do_statement *ClaireProperty  // Generate/"do_statement"
var C_Generate_eid_expression *ClaireProperty  // Generate/"eid_expression"
var C_Generate_stat_exp *ClaireProperty  // Generate/"stat_exp"
var C_Generate_let_eid_ask *ClaireProperty  // Generate/"let_eid?"
var C_Generate_g_eid_stat *ClaireProperty  // Generate/"g_eid_stat"
var C_Generate_bag_class *ClaireProperty  // Generate/"bag_class"
var C_Generate_iteration_statement *ClaireProperty  // Generate/"iteration_statement"
var C_Generate_inline_stat *ClaireProperty  // Generate/"inline_stat"
var C_Generate_update_statement *ClaireProperty  // Generate/"update_statement"
var C_Generate_need_shortcut *ClaireProperty  // Generate/"need_shortcut"
var C_Generate_external_I *ClaireProperty  // Generate/"external!"
var C_Generate_string2module *ClaireProperty  // Generate/"string2module"
var C_Generate_printHelp *ClaireProperty  // Generate/"printHelp"
var C_Generate_complex_main *ClaireProperty  // Generate/"complex_main"
var C_system_file *ClaireProperty  // claire/"system_file"
var C_Generate_compile_dir *ClaireProperty  // Generate/"compile_dir"
var C_Generate_compile_exe *ClaireProperty  // Generate/"compile_exe"
var C_Generate_system_imports *ClaireProperty  // Generate/"system_imports"
var C_Generate_load_function *ClaireProperty  // Generate/"load_function"
var C_Generate_main_function *ClaireProperty  // Generate/"main_function"
var It *ClaireModule

// definition of the meta-model for module Generate 
func MetaLoad() { 
  
  It = MakeModule("Generate",Optimize.C_Compile)
  It.Comment = MakeString("Compiled on Sunday 03-13-2022 07:28:45(v4.0.04), lines:3550, warnings:2,safety:5")
  ClEnv.Module_I = It
  
  // definition of the properties
  C_Generate_open_comparators = MakeProperty("open_comparators",2,It)
  C_Generate_open_operators = MakeProperty("open_operators",2,It)
  C_Generate_div_operators = MakeProperty("div_operators",2,It)
  C_Generate_extension = MakeProperty("extension",2,It)
  C_Generate_interfaces = MakeProperty("interfaces",2,It)
  C_Generate_current = MakeProperty("current",2,It)
  C_Generate_bad_names = MakeProperty("bad_names",2,It)
  C_Generate_good_names = MakeProperty("good_names",2,It)
  C_Generate_kernel_methods = MakeProperty("kernel_methods",2,It)
  C_Generate_varsym = MakeProperty("varsym",2,It)
  C_Generate_indent_c = MakeProperty("indent_c",1,It)
  C_Generate_breakline = MakeProperty("breakline",1,It)
  C_Generate_new_block = MakeProperty("new_block",1,It)
  C_Generate_let_block = MakeProperty("let_block",1,It)
  C_Generate_close_block = MakeProperty("close_block",1,It)
  C_Generate_finish_block = MakeProperty("finish_block",1,It)
  C_g_test = MakeProperty("g_test",1,C_claire)
  C_Generate_g_func = MakeProperty("g_func",1,It)
  C_Generate_g_expression = MakeProperty("g_expression",1,It)
  C_Generate_statement = MakeProperty("statement",1,It)
  C_Generate_gtop = MakeProperty("gtop",1,It)
  C_Generate_make_go_function = MakeProperty("make_go_function",1,It)
  C_compile = MakeProperty("compile",1,C_claire)
  C_Generate_parents = MakeProperty("parents",1,It)
  C_Generate_gen_files = MakeProperty("gen_files",1,It)
  C_Generate_gen_mod_file = MakeProperty("gen_mod_file",1,It)
  C_Generate_gen_file = MakeProperty("gen_file",1,It)
  C_Generate_start_file = MakeProperty("start_file",1,It)
  C_Generate_gen_classes = MakeProperty("gen_classes",1,It)
  C_Generate_gen_objects = MakeProperty("gen_objects",1,It)
  C_Generate_gen_meta_load = MakeProperty("gen_meta_load",1,It)
  C_Generate_namespace_I = MakeProperty("namespace!",1,It)
  C_Generate_import_declaration = MakeProperty("import_declaration",1,It)
  C_Generate_dumb_import = MakeProperty("dumb_import",1,It)
  C_Generate_needed_modules = MakeProperty("needed_modules",1,It)
  C_Generate_cap_short = MakeProperty("cap_short",1,It)
  C_Generate_representative = MakeProperty("representative",1,It)
  C_Generate_clean_duplicates = MakeProperty("clean_duplicates",1,It)
  C_Generate_gen_class_def = MakeProperty("gen_class_def",1,It)
  C_Generate_gen_cast_function = MakeProperty("gen_cast_function",1,It)
  C_Generate_construct_class_ask = MakeProperty("construct_class?",1,It)
  C_Generate_gen_construct = MakeProperty("gen_construct",1,It)
  C_Generate_rootSlot = MakeProperty("rootSlot",1,It)
  C_Generate_go_class = MakeProperty("go_class",1,It)
  C_Generate_interface_I = MakeProperty("interface!",1,It)
  C_Generate_cast_class = MakeProperty("cast_class",1,It)
  C_Generate_add_underscore = MakeProperty("add_underscore",1,It)
  C_Generate_class_ident = MakeProperty("class_ident",1,It)
  C_Generate_cast_prefix = MakeProperty("cast_prefix",1,It)
  C_Generate_cast_post = MakeProperty("cast_post",1,It)
  C_Generate_go_var = MakeProperty("go_var",1,It)
  C_Generate_getRange = MakeProperty("getRange",1,It)
  C_Generate_thing_ident = MakeProperty("thing_ident",1,It)
  C_Generate_gen_module = MakeProperty("gen_module",1,It)
  C_Generate_declare = MakeProperty("declare",1,It)
  C_Generate_var_declaration = MakeProperty("var_declaration",1,It)
  C_Generate_g_statement = MakeProperty("g_statement",1,It)
  C_Generate_simple_func_ask = MakeProperty("simple_func?",1,It)
  C_Generate_get_made = MakeProperty("get_made",1,It)
  C_Generate_fileName = MakeProperty("fileName",1,It)
  C_Generate_make_lambda_function = MakeProperty("make_lambda_function",1,It)
  C_Generate_generate_function_start = MakeProperty("generate_function_start",1,It)
  C_Generate_eid_body = MakeProperty("eid_body",1,It)
  C_Generate_generate_eid_dual = MakeProperty("generate_eid_dual",1,It)
  C_Generate_goMethod_ask = MakeProperty("goMethod?",1,It)
  C_Generate_goVariable = MakeProperty("goVariable",1,It)
  C_Generate_goMethod = MakeProperty("goMethod",1,It)
  C_Generate_goVariables = MakeProperty("goVariables",1,It)
  C_Generate_goFunction = MakeProperty("goFunction",1,It)
  C_Generate_check_range = MakeProperty("check_range",1,It)
  C_Generate_simple_body_ask = MakeProperty("simple_body?",1,It)
  C_Generate_need_debug_ask = MakeProperty("need_debug?",1,It)
  C_Generate_procedure_body = MakeProperty("procedure_body",1,It)
  C_Generate_function_body = MakeProperty("function_body",1,It)
  C_Generate_generate_eid_function = MakeProperty("generate_eid_function",1,It)
  C_Generate_generate_eval_function = MakeProperty("generate_eval_function",1,It)
  C_Generate_debug_intro = MakeProperty("debug_intro",1,It)
  C_Generate_return_result = MakeProperty("return_result",1,It)
  C_Generate_go_signature = MakeProperty("go_signature",1,It)
  C_Generate_goEIDFunctionName = MakeProperty("goEIDFunctionName",1,It)
  C_Generate_goEIDVariables = MakeProperty("goEIDVariables",1,It)
  C_Generate_print_EID_call = MakeProperty("print_EID_call",1,It)
  C_Generate_external_EID_arg = MakeProperty("external_EID_arg",1,It)
  C_Generate_eid_prefix = MakeProperty("eid_prefix",1,It)
  C_Generate_eid_post = MakeProperty("eid_post",1,It)
  C_Generate_to_eid = MakeProperty("to_eid",1,It)
  C_Generate_build_Variable = MakeProperty("build_Variable",1,It)
  C_Generate_import_princ = MakeProperty("import_princ",1,It)
  C_Generate_genvar = MakeProperty("genvar",1,It)
  C_Generate_c_string = MakeProperty("c_string",1,It)
  C_Generate_cap_ident = MakeProperty("cap_ident",1,It)
  C_Generate_capitalized_ident = MakeProperty("capitalized_ident",1,It)
  C_Generate_capitalize = MakeProperty("capitalize",1,It)
  C_Generate_symbol_ident = MakeProperty("symbol_ident",1,It)
  C_Generate_globalVar = MakeProperty("globalVar",1,It)
  C_Generate_type_sort = MakeProperty("type_sort",1,It)
  C_Generate_g_sort = MakeProperty("g_sort",1,It)
  C_Generate_valuesSlot = MakeProperty("valuesSlot",1,It)
  C_imported_function_ask = MakeProperty("imported_function?",1,C_claire)
  C_Generate_arg_match = MakeProperty("arg_match",1,It)
  C_dMethod_ask = MakeProperty("dMethod?",1,C_claire)
  C_Generate_at = MakeProperty("at",1,It)
  C_Generate_getFunctionName = MakeProperty("getFunctionName",1,It)
  C_Generate_preCore_ask = MakeProperty("preCore?",1,It)
  C_Generate_goEIDFunction = MakeProperty("goEIDFunction",1,It)
  C_Generate_goEvalFunction = MakeProperty("goEvalFunction",1,It)
  C_Generate_retreive_method = MakeProperty("retreive_method",1,It)
  C_Generate_retreive_list = MakeProperty("retreive_list",1,It)
  C_Generate_to_cl = MakeProperty("to_cl",1,It)
  C_Generate_from_eid = MakeProperty("from_eid",1,It)
  C_Generate_integer_prefix = MakeProperty("integer_prefix",1,It)
  C_Generate_float_prefix = MakeProperty("float_prefix",1,It)
  C_Generate_char_prefix = MakeProperty("char_prefix",1,It)
  C_Generate_string_prefix = MakeProperty("string_prefix",1,It)
  C_Generate_string_post = MakeProperty("string_post",1,It)
  C_Generate_native_post = MakeProperty("native_post",1,It)
  C_Generate_object_prefix = MakeProperty("object_prefix",1,It)
  C_Generate_object_post = MakeProperty("object_post",1,It)
  C_Generate_equal_exp = MakeProperty("equal_exp",1,It)
  C_Generate_sign_equal = MakeProperty("sign_equal",1,It)
  C_Generate_char_exp_ask = MakeProperty("char_exp?",1,It)
  C_Generate_c_member = MakeProperty("c_member",1,It)
  C_Generate_bag_expression = MakeProperty("bag_expression",1,It)
  C_Generate_args_list = MakeProperty("args_list",1,It)
  C_Generate_constant_ask = MakeProperty("constant?",1,It)
  C_Generate_go_range = MakeProperty("go_range",1,It)
  C_Generate_full_signature = MakeProperty("full_signature",1,It)
  C_Generate_signature_I = MakeProperty("signature!",1,It)
  C_Generate_g_member = MakeProperty("g_member",1,It)
  C_Generate_cast_Values = MakeProperty("cast_Values",1,It)
  C_Generate_check_var = MakeProperty("check_var",1,It)
  C_Generate_use_variable = MakeProperty("use_variable",1,It)
  C_Generate_g_clean = MakeProperty("g_clean",1,It)
  C_Generate_inline_exp = MakeProperty("inline_exp",1,It)
  C_Generate_bounded_expression = MakeProperty("bounded_expression",1,It)
  C_Generate_eid_provide_ask = MakeProperty("eid_provide?",1,It)
  C_Generate_print_external_call = MakeProperty("print_external_call",1,It)
  C_Generate_belong_exp = MakeProperty("belong_exp",1,It)
  C_Generate_external_casted_arg = MakeProperty("external_casted_arg",1,It)
  C_Generate_g_table_index = MakeProperty("g_table_index",1,It)
  C_Generate_sign_or = MakeProperty("sign_or",1,It)
  C_Generate_unfold_args = MakeProperty("unfold_args",1,It)
  C_Generate_unfold_arg = MakeProperty("unfold_arg",1,It)
  C_Generate_unfold_use = MakeProperty("unfold_use",1,It)
  C_Generate_eid_require_ask = MakeProperty("eid_require?",1,It)
  C_Generate_unfold_eid = MakeProperty("unfold_eid",1,It)
  C_Generate_g_try = MakeProperty("g_try",1,It)
  C_Generate_close_try = MakeProperty("close_try",1,It)
  C_Generate_error_wrap = MakeProperty("error_wrap",1,It)
  C_Generate_g_try_void = MakeProperty("g_try_void",1,It)
  C_Generate_do_statement = MakeProperty("do_statement",1,It)
  C_Generate_eid_expression = MakeProperty("eid_expression",1,It)
  C_Generate_stat_exp = MakeProperty("stat_exp",1,It)
  C_Generate_let_eid_ask = MakeProperty("let_eid?",1,It)
  C_Generate_g_eid_stat = MakeProperty("g_eid_stat",1,It)
  C_Generate_bag_class = MakeProperty("bag_class",1,It)
  C_Generate_iteration_statement = MakeProperty("iteration_statement",1,It)
  C_Generate_inline_stat = MakeProperty("inline_stat",1,It)
  C_Generate_update_statement = MakeProperty("update_statement",1,It)
  C_Generate_need_shortcut = MakeProperty("need_shortcut",1,It)
  C_Generate_external_I = MakeProperty("external!",1,It)
  C_Generate_string2module = MakeProperty("string2module",1,It)
  C_Generate_printHelp = MakeProperty("printHelp",1,It)
  C_Generate_complex_main = MakeProperty("complex_main",1,It)
  C_system_file = MakeProperty("system_file",1,C_claire)
  C_Generate_compile_dir = MakeProperty("compile_dir",1,It)
  C_Generate_compile_exe = MakeProperty("compile_exe",1,It)
  C_Generate_system_imports = MakeProperty("system_imports",1,It)
  C_Generate_load_function = MakeProperty("load_function",1,It)
  C_Generate_main_function = MakeProperty("main_function",1,It)
  
  // instructions from module sources
  { 
    var expr EID 
    { var _CL_obj *Core.GlobalVariable  
      C_Generate__star_ask_interval_star = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("*?_interval*",It)))
      
      _CL_obj = C_Generate__star_ask_interval_star
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = Core.F__at_property1(C_set_I,C_Interval).Id()
      expr = _CL_obj.Close()
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID 
    { var _CL_obj *Core.GlobalVariable  
      C_Generate__star_dash_dash_integer_star = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("*--_integer*",It)))
      
      _CL_obj = C_Generate__star_dash_dash_integer_star
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = Core.F__at_property1(ToProperty(C__dot_dot.Id()),C_integer).Id()
      expr = _CL_obj.Close()
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID 
    { var _CL_obj *Core.GlobalVariable  
      C_Generate__star_plus_integer_star = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("*+_integer*",It)))
      
      _CL_obj = C_Generate__star_plus_integer_star
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = Core.F__at_property1(ToProperty(Core.C__plus.Id()),C_integer).Id()
      expr = _CL_obj.Close()
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID 
    { var _CL_obj *Core.GlobalVariable  
      C_Generate__starnth_integer_star = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("*nth_integer*",It)))
      
      _CL_obj = C_Generate__starnth_integer_star
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = Core.F__at_property1(C_nth,C_integer).Id()
      expr = _CL_obj.Close()
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID 
    { var _CL_obj *Core.GlobalVariable  
      C_Generate__starnth_list_star = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("*nth_list*",It)))
      
      _CL_obj = C_Generate__starnth_list_star
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = Core.F__at_property1(C_nth,C_list).Id()
      expr = _CL_obj.Close()
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID 
    { var _CL_obj *Core.GlobalVariable  
      C_Generate__starnth_tuple_star = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("*nth_tuple*",It)))
      
      _CL_obj = C_Generate__starnth_tuple_star
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = Core.F__at_property1(C_nth,C_tuple).Id()
      expr = _CL_obj.Close()
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID 
    { var _CL_obj *Core.GlobalVariable  
      C_Generate__starnth_1_list_star = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("*nth_1_list*",It)))
      
      _CL_obj = C_Generate__starnth_1_list_star
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = Core.F__at_property1(C_nth_get,C_list).Id()
      expr = _CL_obj.Close()
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID 
    { var _CL_obj *Core.GlobalVariable  
      C_Generate__starnth_1_tuple_star = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("*nth_1_tuple*",It)))
      
      _CL_obj = C_Generate__starnth_1_tuple_star
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = Core.F__at_property1(C_nth_get,C_tuple).Id()
      expr = _CL_obj.Close()
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID 
    { var _CL_obj *Core.GlobalVariable  
      C_Generate__starnth_1_array_star = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("*nth_1_array*",It)))
      
      _CL_obj = C_Generate__starnth_1_array_star
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = Core.F__at_property1(C_nth_get,C_array).Id()
      expr = _CL_obj.Close()
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID 
    { var _CL_obj *Core.GlobalVariable  
      C_Generate__starnth_string_star = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("*nth_string*",It)))
      
      _CL_obj = C_Generate__starnth_string_star
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = Core.F__at_property1(C_nth,C_string).Id()
      expr = _CL_obj.Close()
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID 
    { var _CL_obj *Core.GlobalVariable  
      C_Generate__starnth_1_string_star = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("*nth_1_string*",It)))
      
      _CL_obj = C_Generate__starnth_1_string_star
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = Core.F__at_property1(C_nth_get,C_string).Id()
      expr = _CL_obj.Close()
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID 
    { var _CL_obj *Core.GlobalVariable  
      C_Generate__starnth_equal_list_star = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("*nth=_list*",It)))
      
      _CL_obj = C_Generate__starnth_equal_list_star
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = Core.F__at_property1(C_nth_equal,C_list).Id()
      expr = _CL_obj.Close()
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID 
    { var _CL_obj *Core.GlobalVariable  
      C_Generate__starnth_put_list_star = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("*nth_put_list*",It)))
      
      _CL_obj = C_Generate__starnth_put_list_star
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = Core.F__at_property1(C_nth_put,C_list).Id()
      expr = _CL_obj.Close()
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID 
    { var _CL_obj *Core.GlobalVariable  
      C_Generate__starmake_list_star = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("*make_list*",It)))
      
      _CL_obj = C_Generate__starmake_list_star
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = Core.F__at_property2(C_make_list,MakeConstantList(C_integer.Id(),C_type.Id(),C_any.Id())).Id()
      expr = _CL_obj.Close()
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID 
    { var _CL_obj *Core.GlobalVariable  
      C_Generate__starnot_star = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("*not*",It)))
      
      _CL_obj = C_Generate__starnot_star
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = Core.F__at_property1(Core.C_not,C_any).Id()
      expr = _CL_obj.Close()
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID 
    { var _CL_obj *Core.GlobalVariable  
      C_Generate__starknown_star = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("*known*",It)))
      
      _CL_obj = C_Generate__starknown_star
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = Core.F__at_property1(Core.C_known_ask,C_any).Id()
      expr = _CL_obj.Close()
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID 
    { var _CL_obj *Core.GlobalVariable  
      C_Generate__starunknown_star = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("*unknown*",It)))
      
      _CL_obj = C_Generate__starunknown_star
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = Core.F__at_property1(Core.C_unknown_ask,C_any).Id()
      expr = _CL_obj.Close()
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID 
    { var _CL_obj *Core.GlobalVariable  
      C_Generate__starnot_equal_star = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("*not_equal*",It)))
      
      _CL_obj = C_Generate__starnot_equal_star
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = Core.F__at_property1(ToProperty(Core.C__I_equal.Id()),C_any).Id()
      expr = _CL_obj.Close()
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID 
    { var _CL_obj *Core.GlobalVariable  
      C_Generate__starequal_star = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("*equal*",It)))
      
      _CL_obj = C_Generate__starequal_star
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = Core.F__at_property1(ToProperty(C__equal.Id()),C_any).Id()
      expr = _CL_obj.Close()
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID 
    { var _CL_obj *Core.GlobalVariable  
      C_Generate__starbelong_star = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("*belong*",It)))
      
      _CL_obj = C_Generate__starbelong_star
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = Core.F__at_property1(Core.C_Core_belong,C_any).Id()
      expr = _CL_obj.Close()
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID 
    { var _CL_obj *Core.GlobalVariable  
      C_Generate__starcontain_list_star = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("*contain_list*",It)))
      
      _CL_obj = C_Generate__starcontain_list_star
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = Core.F__at_property1(C_contain_ask,C_list).Id()
      expr = _CL_obj.Close()
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID 
    { var _CL_obj *Core.GlobalVariable  
      C_Generate__starcontain_set_star = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("*contain_set*",It)))
      
      _CL_obj = C_Generate__starcontain_set_star
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = Core.F__at_property1(C_contain_ask,C_set).Id()
      expr = _CL_obj.Close()
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID 
    { var _CL_obj *Core.GlobalVariable  
      C_Generate__starlength_array_star = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("*length_array*",It)))
      
      _CL_obj = C_Generate__starlength_array_star
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = Core.F__at_property1(C_length,C_array).Id()
      expr = _CL_obj.Close()
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID 
    { var _CL_obj *Core.GlobalVariable  
      C_Generate__starlength_bag_star = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("*length_bag*",It)))
      
      _CL_obj = C_Generate__starlength_bag_star
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = Core.F__at_property1(C_length,C_bag).Id()
      expr = _CL_obj.Close()
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID 
    { var _CL_obj *Core.GlobalVariable  
      C_Generate__starclose_exception_star = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("*close_exception*",It)))
      
      _CL_obj = C_Generate__starclose_exception_star
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = Core.F__at_property1(C_close,C_exception).Id()
      expr = _CL_obj.Close()
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID 
    { var _CL_obj *Core.GlobalVariable  
      C_Generate__starnew_class1_star = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("*new_class1*",It)))
      
      _CL_obj = C_Generate__starnew_class1_star
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = Core.F__at_property2(C_mClaire_new_I,MakeConstantList(C_class.Id())).Id()
      expr = _CL_obj.Close()
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID 
    { var _CL_obj *Core.GlobalVariable  
      C_Generate__starnew_class2_star = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("*new_class2*",It)))
      
      _CL_obj = C_Generate__starnew_class2_star
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = Core.F__at_property2(C_mClaire_new_I,MakeConstantList(C_class.Id(),C_symbol.Id())).Id()
      expr = _CL_obj.Close()
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID 
    { var _CL_obj *Core.GlobalVariable  
      C_Generate__starslot_get_star = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("*slot_get*",It)))
      
      _CL_obj = C_Generate__starslot_get_star
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = Core.F__at_property1(C_slot_get,C_object).Id()
      expr = _CL_obj.Close()
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID 
    { var _CL_obj *Core.GlobalVariable  
      C_Generate__starmap_star = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("*map*",It)))
      
      _CL_obj = C_Generate__starmap_star
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = Core.F__at_property1(C_map_I,C_type).Id()
      expr = _CL_obj.Close()
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID 
    { var _CL_obj *Core.GlobalVariable  
      C_Generate__starof_bag_star = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("*of_bag*",It)))
      
      _CL_obj = C_Generate__starof_bag_star
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = Core.F__at_property1(C_of,C_bag).Id()
      expr = _CL_obj.Close()
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID 
    { var _CL_obj *Core.GlobalVariable  
      C_Generate__starof_array_star = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("*of_array*",It)))
      
      _CL_obj = C_Generate__starof_array_star
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = Core.F__at_property1(C_of,C_array).Id()
      expr = _CL_obj.Close()
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID 
    { var _CL_obj *Core.GlobalVariable  
      C_Generate__starcopy_list_star = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("*copy_list*",It)))
      
      _CL_obj = C_Generate__starcopy_list_star
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = Core.F__at_property1(C_copy,C_list).Id()
      expr = _CL_obj.Close()
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID 
    { var _CL_obj *Core.GlobalVariable  
      C_Generate__starcopy_set_star = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("*copy_set*",It)))
      
      _CL_obj = C_Generate__starcopy_set_star
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = Core.F__at_property1(C_copy,C_set).Id()
      expr = _CL_obj.Close()
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID 
    { var _CL_obj *Core.GlobalVariable  
      C_Generate__starempty_set_star = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("*empty_set*",It)))
      
      _CL_obj = C_Generate__starempty_set_star
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = Core.F__at_property1(C_empty,C_set).Id()
      expr = _CL_obj.Close()
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID 
    { var _CL_obj *Core.GlobalVariable  
      C_Generate__starnth_put_array_star = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("*nth_put_array*",It)))
      
      _CL_obj = C_Generate__starnth_put_array_star
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = Core.F__at_property1(C_nth_put,C_array).Id()
      expr = _CL_obj.Close()
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID 
    { var _CL_obj *Core.GlobalVariable  
      C_Generate__star_Zt_star = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("*%t*",It)))
      
      _CL_obj = C_Generate__star_Zt_star
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = Core.F__at_property1(ToProperty(Core.C_Core__Zt.Id()),C_any).Id()
      expr = _CL_obj.Close()
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID 
    { var _CL_obj *Core.GlobalVariable  
      C_Generate__starlesst_star = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("*lesst*",It)))
      
      _CL_obj = C_Generate__starlesst_star
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = Core.F__at_property1(ToProperty(Core.C_Core__inf_equalt.Id()),C_type).Id()
      expr = _CL_obj.Close()
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID 
    { var _CL_obj *Core.GlobalVariable  
      C_Generate__starincluded_star = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("*included*",It)))
      
      _CL_obj = C_Generate__starincluded_star
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = Core.F__at_property1(ToProperty(C__inf_equal.Id()),C_type).Id()
      expr = _CL_obj.Close()
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID 
    { var _CL_obj *Core.GlobalVariable  
      C_Generate__starstack_apply_star = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("*stack_apply*",It)))
      
      _CL_obj = C_Generate__starstack_apply_star
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = Core.F__at_property1(C_stack_apply,C_property).Id()
      expr = _CL_obj.Close()
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID 
    { var _CL_obj *Core.GlobalVariable  
      C_Generate__starsuper_apply_star = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("*super_apply*",It)))
      
      _CL_obj = C_Generate__starsuper_apply_star
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = Core.F__at_property1(Core.C_Core_super_apply,C_property).Id()
      expr = _CL_obj.Close()
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID 
    { var _CL_obj *Core.GlobalVariable  
      C_Generate__starprinc_string_star = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("*princ_string*",It)))
      
      _CL_obj = C_Generate__starprinc_string_star
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = Core.F__at_property1(C_princ,C_string).Id()
      expr = _CL_obj.Close()
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID 
    { var _CL_obj *Core.GlobalVariable  
      C_Generate__starinherit_star = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("*inherit*",It)))
      
      _CL_obj = C_Generate__starinherit_star
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = Core.F__at_property1(ToProperty(Core.C_inherit_ask.Id()),C_class).Id()
      expr = _CL_obj.Close()
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID 
    { var _CL_obj *Core.GlobalVariable  
      C_Generate__starwrite_value_star = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("*write_value*",It)))
      
      _CL_obj = C_Generate__starwrite_value_star
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = Core.F__at_property1(Language.C_Language_write_value,C_Variable).Id()
      expr = _CL_obj.Close()
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID 
    { var _CL_obj *Core.GlobalVariable  
      C_Generate__starread_property_star = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("*read_property*",It)))
      
      _CL_obj = C_Generate__starread_property_star
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = Core.F__at_property1(C_read,C_property).Id()
      expr = _CL_obj.Close()
      } 
    ErrorCheck(expr)} 
  
  C_Generate_code_producer = MakeClass("code_producer",Optimize.C_Compile_producer,It)
  Core.F_close_slot(C_Generate_code_producer.AddSlot(C_Generate_open_comparators,Core.F_nth_class1(C_list,ToType(C_operation.Id())),ToType(C_operation.Id()).EmptyList().Id()))
  Core.F_close_slot(C_Generate_code_producer.AddSlot(C_Generate_open_operators,Core.F_nth_class1(C_list,ToType(C_operation.Id())),ToType(C_operation.Id()).EmptyList().Id()))
  Core.F_close_slot(C_Generate_code_producer.AddSlot(C_Generate_div_operators,Core.F_nth_class1(C_list,ToType(C_operation.Id())),ToType(C_operation.Id()).EmptyList().Id()))
  Core.F_close_slot(C_Generate_code_producer.AddSlot(C_body,ToType(C_any.Id()),MakeInteger(0).Id()))
  Core.F_close_slot(C_Generate_code_producer.AddSlot(C_Generate_extension,ToType(C_string.Id()),CNULL))
  Core.F_close_slot(C_Generate_code_producer.AddSlot(C_comment,ToType(C_string.Id()),CNULL))
  Core.F_close_slot(C_Generate_code_producer.AddSlot(C_Generate_interfaces,ToType(C_list.Id()),ToType(C_any.Id()).EmptyList().Id()))
  Core.F_close_slot(C_Generate_code_producer.AddSlot(C_stat,ToType(C_integer.Id()),MakeInteger(0).Id()))
  
  C_Generate_go_producer = MakeClass("go_producer",C_Generate_code_producer,It)
  Core.F_close_slot(C_Generate_go_producer.AddSlot(C_Generate_current,ToType(C_module.Id()),CNULL))
  Core.F_close_slot(C_Generate_go_producer.AddSlot(C_Generate_bad_names,Core.F_nth_class1(C_list,ToType(C_symbol.Id())),ToType(C_symbol.Id()).EmptyList().Id()))
  Core.F_close_slot(C_Generate_go_producer.AddSlot(C_Generate_good_names,Core.F_nth_class1(C_list,ToType(C_symbol.Id())),ToType(C_symbol.Id()).EmptyList().Id()))
  Core.F_close_slot(C_Generate_go_producer.AddSlot(C_Generate_kernel_methods,ToType(C_list.Id()),ToType(C_any.Id()).EmptyList().Id()))
  Core.F_close_slot(C_Generate_go_producer.AddSlot(C_source,ToType(C_string.Id()),CNULL))
  Core.F_close_slot(C_Generate_go_producer.AddSlot(Optimize.C_debug_ask,ToType(C_boolean.Id()),CFALSE.Id()))
  Core.F_close_slot(C_Generate_go_producer.AddSlot(C_Generate_varsym,ToType(C_integer.Id()),MakeInteger(0).Id()))
  
  { 
    var expr EID 
    { var _CL_obj *Core.GlobalVariable  
      C_Generate_EIDSET = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("EIDSET",It)))
      
      _CL_obj = C_Generate_EIDSET
      _CL_obj.Range = Core.F_param_I_class(C_set,ToType(C_any.Id()))
      _CL_obj.Value = MakeSet(ToType(C_any.Id()),Core.F__at_property1(C_nth,C_list).Id()).Id()
      expr = _CL_obj.Close()
      } 
    ErrorCheck(expr)} 
  
  _ = Core.F_attach_method(Language.C_iClaire_ident.AddMethod(Signature(C_symbol.Id(),C_void.Id()),0,MakeFunction1(E_iClaire_ident_symbol,"iClaire_ident_symbol")),MakeString("gosystem.cl:120"))
  
  _ = Core.F_attach_method(Language.C_iClaire_ident.AddMethod(Signature(C_thing.Id(),C_void.Id()),0,MakeFunction1(E_iClaire_ident_thing,"iClaire_ident_thing")),MakeString("gosystem.cl:121"))
  
  _ = Core.F_attach_method(Language.C_iClaire_ident.AddMethod(Signature(C_class.Id(),C_void.Id()),0,MakeFunction1(E_iClaire_ident_class,"iClaire_ident_class")),MakeString("gosystem.cl:122"))
  
  _ = Core.F_attach_method(C_Generate_indent_c.AddMethod(Signature(C_void.Id(),C_any.Id()),0,MakeFunction1(E_Generate_indent_c_void,"Generate_indent_c_void")),MakeString("gosystem.cl:126"))
  
  _ = Core.F_attach_method(C_Generate_breakline.AddMethod(Signature(C_void.Id(),C_any.Id()),0,MakeFunction1(E_Generate_breakline_void,"Generate_breakline_void")),MakeString("gosystem.cl:128"))
  
  _ = Core.F_attach_method(C_Generate_new_block.AddMethod(Signature(C_void.Id(),C_void.Id()),0,MakeFunction1(E_Generate_new_block_void,"Generate_new_block_void")),MakeString("gosystem.cl:132"))
  
  _ = Core.F_attach_method(C_Generate_let_block.AddMethod(Signature(C_void.Id(),C_void.Id()),0,MakeFunction1(E_Generate_let_block_void,"Generate_let_block_void")),MakeString("gosystem.cl:136"))
  
  _ = Core.F_attach_method(C_Generate_close_block.AddMethod(Signature(C_void.Id(),C_void.Id()),0,MakeFunction1(E_Generate_close_block_void,"Generate_close_block_void")),MakeString("gosystem.cl:140"))
  
  _ = Core.F_attach_method(C_Generate_finish_block.AddMethod(Signature(C_void.Id(),C_void.Id()),0,MakeFunction1(E_Generate_finish_block_void,"Generate_finish_block_void")),MakeString("gosystem.cl:144"))
  
  _ = Core.F_attach_method(C_g_test.AddMethod(Signature(C_any.Id(),C_void.Id()),1,MakeFunction1(E_g_test_any,"g_test_any")),MakeString("gosystem.cl:153"))
  
  _ = Core.F_attach_method(C_g_test.AddMethod(Signature(C_module.Id(),C_any.Id(),C_void.Id()),1,MakeFunction2(E_g_test_module,"g_test_module")),MakeString("gosystem.cl:166"))
  
  _ = Core.F_attach_method(C_Generate_gtop.AddMethod(Signature(C_void.Id(),C_void.Id()),1,MakeFunction1(E_Generate_gtop_void,"Generate_gtop_void")),MakeString("gosystem.cl:173"))
  
  _ = Core.F_attach_method(C_g_test.AddMethod(Signature(C_method.Id(),C_void.Id()),1,MakeFunction1(E_g_test_method,"g_test_method")),MakeString("gosystem.cl:193"))
  
  { 
    var expr EID 
    { var _CL_obj *Core.GlobalVariable  
      C_BadMethods = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("BadMethods",C_claire)))
      
      _CL_obj = C_BadMethods
      _CL_obj.Range = Core.F_param_I_class(C_list,ToType(C_method.Id()))
      _CL_obj.Value = CNULL
      expr = _CL_obj.Close()
      } 
    ErrorCheck(expr)} 
  
  _ = Core.F_attach_method(C_compile.AddMethod(Signature(C_module.Id(),C_void.Id()),1,MakeFunction1(E_compile_module,"compile_module")),MakeString("gosystem.cl:200"))
  
  _ = Core.F_attach_method(C_compile.AddMethod(Signature(C_Generate_go_producer.Id(),C_module.Id(),C_void.Id()),1,MakeFunction2(E_compile_go_producer,"compile_go_producer")),MakeString("gosystem.cl:221"))
  
  _ = Core.F_attach_method(C_Generate_gen_files.AddMethod(Signature(C_Generate_go_producer.Id(),C_module.Id(),C_void.Id()),1,MakeFunction2(E_Generate_gen_files_go_producer,"Generate_gen_files_go_producer")),MakeString("gosystem.cl:239"))
  
  _ = Core.F_attach_method(C_Generate_gen_mod_file.AddMethod(Signature(C_Generate_go_producer.Id(),C_module.Id(),C_void.Id()),1,MakeFunction2(E_Generate_gen_mod_file_go_producer,"Generate_gen_mod_file_go_producer")),MakeString("gosystem.cl:262"))
  
  _ = Core.F_attach_method(C_Generate_start_file.AddMethod(Signature(C_Generate_go_producer.Id(),
    C_string.Id(),
    C_module.Id(),
    C_boolean.Id(),
    C_any.Id()),1,MakeFunction4(E_Generate_start_file_go_producer,"Generate_start_file_go_producer")),MakeString("gosystem.cl:281"))
  
  _ = Core.F_attach_method(C_Generate_import_declaration.AddMethod(Signature(C_module.Id(),C_void.Id()),1,MakeFunction1(E_Generate_import_declaration_module,"Generate_import_declaration_module")),MakeString("gosystem.cl:288"))
  
  _ = Core.F_attach_method(C_Generate_needed_modules.AddMethod(Signature(C_module.Id(),C_list.Id()),0,MakeFunction1(E_Generate_needed_modules_module,"Generate_needed_modules_module")),MakeString("gosystem.cl:296"))
  
  _ = Core.F_attach_method(C_Generate_dumb_import.AddMethod(Signature(C_module.Id(),C_void.Id()),0,MakeFunction1(E_Generate_dumb_import_module,"Generate_dumb_import_module")),MakeString("gosystem.cl:309"))
  
  _ = Core.F_attach_method(C_Generate_representative.AddMethod(Signature(C_module.Id(),C_any.Id()),1,MakeFunction1(E_Generate_representative_module,"Generate_representative_module")),MakeString("gosystem.cl:315"))
  
  _ = Core.F_attach_method(C_Generate_clean_duplicates.AddMethod(Signature(C_list.Id(),C_list.Id()),0,MakeFunction1(E_Generate_clean_duplicates_list,"Generate_clean_duplicates_list")),MakeString("gosystem.cl:324"))
  
  _ = Core.F_attach_method(C_Generate_gen_classes.AddMethod(Signature(C_Generate_go_producer.Id(),C_module.Id(),C_void.Id()),1,MakeFunction2(E_Generate_gen_classes_go_producer,"Generate_gen_classes_go_producer")),MakeString("gosystem.cl:338"))
  
  _ = Core.F_attach_method(C_Generate_rootSlot.AddMethod(Signature(C_slot.Id(),C_slot.Id()),0,MakeFunction1(E_Generate_rootSlot_slot,"Generate_rootSlot_slot")),MakeString("gosystem.cl:348"))
  
  _ = Core.F_attach_method(C_Generate_gen_class_def.AddMethod(Signature(C_Generate_go_producer.Id(),C_class.Id(),C_void.Id()),1,MakeFunction2(E_Generate_gen_class_def_go_producer,"Generate_gen_class_def_go_producer")),MakeString("gosystem.cl:361"))
  
  _ = Core.F_attach_method(C_Generate_gen_cast_function.AddMethod(Signature(C_Generate_go_producer.Id(),C_class.Id(),C_void.Id()),0,MakeFunction2(E_Generate_gen_cast_function_go_producer,"Generate_gen_cast_function_go_producer")),MakeString("gosystem.cl:370"))
  
  _ = Core.F_attach_method(C_Generate_construct_class_ask.AddMethod(Signature(C_class.Id(),C_boolean.Id()),0,MakeFunction1(E_Generate_construct_class_ask_class,"Generate_construct_class?_class")),MakeString("gosystem.cl:375"))
  
  _ = Core.F_attach_method(C_Generate_gen_construct.AddMethod(Signature(C_Generate_go_producer.Id(),C_class.Id(),C_void.Id()),1,MakeFunction2(E_Generate_gen_construct_go_producer,"Generate_gen_construct_go_producer")),MakeString("gosystem.cl:398"))
  
  _ = Core.F_attach_method(C_Generate_gen_objects.AddMethod(Signature(C_Generate_go_producer.Id(),C_module.Id(),C_void.Id()),1,MakeFunction2(E_Generate_gen_objects_go_producer,"Generate_gen_objects_go_producer")),MakeString("gosystem.cl:428"))
  
  _ = Core.F_attach_method(C_Generate_getRange.AddMethod(Signature(Core.C_global_variable.Id(),C_class.Id()),0,MakeFunction1(E_Generate_getRange_global_variable,"Generate_getRange_global_variable")),MakeString("gosystem.cl:432"))
  
  _ = Core.F_attach_method(C_Generate_gen_meta_load.AddMethod(Signature(C_Generate_go_producer.Id(),C_module.Id(),C_void.Id()),1,MakeFunction2(E_Generate_gen_meta_load_go_producer,"Generate_gen_meta_load_go_producer")),MakeString("gosystem.cl:464"))
  
  _ = Core.F_attach_method(C_Generate_gen_module.AddMethod(Signature(C_Generate_go_producer.Id(),
    C_module.Id(),
    C_module.Id(),
    C_void.Id()),1,MakeFunction3(E_Generate_gen_module_go_producer,"Generate_gen_module_go_producer")),MakeString("gosystem.cl:483"))
  
  _ = Core.F_attach_method(C_Generate_get_made.AddMethod(Signature(C_module.Id(),C_module.Id()),0,MakeFunction1(E_Generate_get_made_module,"Generate_get_made_module")),MakeString("gosystem.cl:488"))
  
  _ = Core.F_attach_method(C_Generate_declare.AddMethod(Signature(C_Generate_go_producer.Id(),C_property.Id(),C_void.Id()),1,MakeFunction2(E_Generate_declare_go_producer,"Generate_declare_go_producer")),MakeString("gosystem.cl:498"))
  
  _ = Core.F_attach_method(C_Generate_parents.AddMethod(Signature(C_module.Id(),C_list.Id(),C_list.Id()),0,MakeFunction2(E_Generate_parents_module,"Generate_parents_module")),MakeString("gosystem.cl:505"))
  
  _ = Core.F_attach_method(C_Generate_parents.AddMethod(Signature(C_list.Id(),C_list.Id()),0,MakeFunction1(E_Generate_parents_list,"Generate_parents_list")),MakeString("gosystem.cl:511"))
  
  _ = Core.F_attach_method(C_get.AddMethod(Signature(C_module.Id(),C_void.Id()),1,MakeFunction1(E_get_module2,"get_module2")),MakeString("gosystem.cl:515"))
  
  _ = Core.F_attach_method(C_Generate_gen_file.AddMethod(Signature(C_Generate_go_producer.Id(),
    C_string.Id(),
    C_string.Id(),
    C_void.Id()),1,MakeFunction3(E_Generate_gen_file_go_producer,"Generate_gen_file_go_producer")),MakeString("gosystem.cl:549"))
  
  _ = Core.F_attach_method(C_Generate_fileName.AddMethod(Signature(C_string.Id(),C_string.Id()),0,MakeFunction1(E_Generate_fileName_string,"Generate_fileName_string")),MakeString("gosystem.cl:554"))
  
  _ = Core.F_attach_method(Optimize.C_Compile_make_c_function.AddMethod(Signature(C_lambda.Id(),
    C_string.Id(),
    C_any.Id(),
    C_void.Id()),1,MakeFunction3(E_Compile_make_c_function_lambda,"Compile_make_c_function_lambda")),MakeString("gosystem.cl:565"))
  
  _ = Core.F_attach_method(C_Generate_make_lambda_function.AddMethod(Signature(C_Generate_go_producer.Id(),
    C_lambda.Id(),
    C_string.Id(),
    C_void.Id()),1,MakeFunction3(E_Generate_make_lambda_function_go_producer,"Generate_make_lambda_function_go_producer")),MakeString("gosystem.cl:584"))
  
  _ = Core.F_attach_method(C_Generate_generate_function_start.AddMethod(Signature(C_Generate_go_producer.Id(),
    C_lambda.Id(),
    C_class.Id(),
    C_any.Id(),
    C_string.Id(),
    C_void.Id()),1,MakeFunction5(E_Generate_generate_function_start_go_producer,"Generate_generate_function_start_go_producer")),MakeString("gosystem.cl:604"))
  
  _ = Core.F_attach_method(C_Generate_make_go_function.AddMethod(Signature(C_Generate_go_producer.Id(),
    C_lambda.Id(),
    C_string.Id(),
    C_method.Id(),
    C_void.Id()),1,MakeFunction4(E_Generate_make_go_function_go_producer,"Generate_make_go_function_go_producer")),MakeString("gosystem.cl:646"))
  
  _ = Core.F_attach_method(C_Generate_simple_body_ask.AddMethod(Signature(C_any.Id(),C_boolean.Id()),1,MakeFunction1(E_Generate_simple_body_ask_any,"Generate_simple_body?_any")),MakeString("gosystem.cl:654"))
  
  _ = Core.F_attach_method(C_Generate_function_body.AddMethod(Signature(C_any.Id(),C_class.Id(),C_void.Id()),1,MakeFunction2(E_Generate_function_body_any,"Generate_function_body_any")),MakeString("gosystem.cl:669"))
  
  _ = Core.F_attach_method(C_Generate_function_body.AddMethod(Signature(Language.C_If.Id(),C_class.Id(),C_void.Id()),1,MakeFunction2(E_Generate_function_body_If,"Generate_function_body_If")),MakeString("gosystem.cl:684"))
  
  _ = Core.F_attach_method(C_Generate_function_body.AddMethod(Signature(Language.C_Do.Id(),C_class.Id(),C_void.Id()),1,MakeFunction2(E_Generate_function_body_Do,"Generate_function_body_Do")),MakeString("gosystem.cl:693"))
  
  _ = Core.F_attach_method(C_Generate_procedure_body.AddMethod(Signature(C_method.Id(),
    C_lambda.Id(),
    C_any.Id(),
    C_class.Id(),
    C_void.Id()),1,MakeFunction4(E_Generate_procedure_body_method,"Generate_procedure_body_method")),MakeString("gosystem.cl:703"))
  
  _ = Core.F_attach_method(C_Generate_eid_body.AddMethod(Signature(C_any.Id(),
    C_boolean.Id(),
    C_class.Id(),
    C_void.Id()),1,MakeFunction3(E_Generate_eid_body_any,"Generate_eid_body_any")),MakeString("gosystem.cl:709"))
  
  _ = Core.F_attach_method(C_Generate_eid_body.AddMethod(Signature(C_method.Id(),
    C_any.Id(),
    C_boolean.Id(),
    C_class.Id(),
    C_void.Id()),1,MakeFunction4(E_Generate_eid_body_method,"Generate_eid_body_method")),MakeString("gosystem.cl:720"))
  
  _ = Core.F_attach_method(C_Generate_generate_eid_function.AddMethod(Signature(C_lambda.Id(),
    C_method.Id(),
    C_boolean.Id(),
    C_void.Id()),1,MakeFunction3(E_Generate_generate_eid_function_lambda,"Generate_generate_eid_function_lambda")),MakeString("gosystem.cl:731"))
  
  _ = Core.F_attach_method(C_Generate_generate_eid_dual.AddMethod(Signature(C_lambda.Id(),C_string.Id(),C_void.Id()),1,MakeFunction2(E_Generate_generate_eid_dual_lambda,"Generate_generate_eid_dual_lambda")),MakeString("gosystem.cl:744"))
  
  _ = Core.F_attach_method(C_Generate_print_EID_call.AddMethod(Signature(C_method.Id(),
    C_list.Id(),
    Core.F_param_I_class(C_list,ToType(C_class.Id())).Id(),
    C_boolean.Id(),
    C_void.Id()),1,MakeFunction4(E_Generate_print_EID_call_method,"Generate_print_EID_call_method")),MakeString("gosystem.cl:767"))
  
  _ = Core.F_attach_method(C_Generate_external_EID_arg.AddMethod(Signature(C_Variable.Id(),
    C_class.Id(),
    C_integer.Id(),
    C_boolean.Id(),
    C_void.Id()),1,MakeFunction4(E_Generate_external_EID_arg_Variable,"Generate_external_EID_arg_Variable")),MakeString("gosystem.cl:776"))
  
  _ = Core.F_attach_method(C_Generate_goEIDVariables.AddMethod(Signature(C_Generate_go_producer.Id(),C_list.Id(),C_any.Id()),0,MakeFunction2(E_Generate_goEIDVariables_go_producer,"Generate_goEIDVariables_go_producer")),MakeString("gosystem.cl:783"))
  
  { 
    var expr EID 
    { var _CL_obj *Core.GlobalVariable  
      C_ABODY = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("ABODY",C_claire)))
      
      _CL_obj = C_ABODY
      _CL_obj.Range = ToType(C_any.Id())
      _CL_obj.Value = CNULL
      expr = _CL_obj.Close()
      } 
    ErrorCheck(expr)} 
  
  _ = Core.F_attach_method(C_Generate_check_range.AddMethod(Signature(C_method.Id(),C_any.Id(),C_any.Id()),1,MakeFunction2(E_Generate_check_range_method,"Generate_check_range_method")),MakeString("gosystem.cl:800"))
  
  _ = Core.F_attach_method(C_Generate_generate_eval_function.AddMethod(Signature(C_lambda.Id(),C_method.Id(),C_void.Id()),1,MakeFunction2(E_Generate_generate_eval_function_lambda,"Generate_generate_eval_function_lambda")),MakeString("gosystem.cl:812"))
  
  _ = Core.F_attach_method(C_Generate_need_debug_ask.AddMethod(Signature(C_any.Id(),C_boolean.Id()),0,MakeFunction1(E_Generate_need_debug_ask_any,"Generate_need_debug?_any")),MakeString("gosystem.cl:824"))
  
  _ = Core.F_attach_method(C_Generate_debug_intro.AddMethod(Signature(C_Generate_go_producer.Id(),
    C_lambda.Id(),
    C_method.Id(),
    C_void.Id()),1,MakeFunction3(E_Generate_debug_intro_go_producer,"Generate_debug_intro_go_producer")),MakeString("gosystem.cl:840"))
  
  _ = Core.F_attach_method(C_Generate_return_result.AddMethod(Signature(C_Generate_go_producer.Id(),
    C_class.Id(),
    C_method.Id(),
    C_string.Id(),
    C_void.Id()),1,MakeFunction4(E_Generate_return_result_go_producer,"Generate_return_result_go_producer")),MakeString("gosystem.cl:853"))
  
  _ = Core.F_attach_method(C_c_princ.AddMethod(Signature(C_function.Id(),C_void.Id()),0,MakeFunction1(E_c_princ_function,"c_princ_function")),MakeString("gosystem.cl:858"))
  
  _ = Core.F_attach_method(C_Generate_import_princ.AddMethod(Signature(C_string.Id(),C_void.Id()),0,MakeFunction1(E_Generate_import_princ_string,"Generate_import_princ_string")),MakeString("gosystem.cl:861"))
  
  _ = Core.F_attach_method(C_Generate_new_block.AddMethod(Signature(C_string.Id(),C_void.Id()),0,MakeFunction1(E_Generate_new_block_string,"Generate_new_block_string")),MakeString("gogen.cl:25"))
  
  _ = Core.F_attach_method(C_Generate_close_block.AddMethod(Signature(C_string.Id(),C_void.Id()),0,MakeFunction1(E_Generate_close_block_string,"Generate_close_block_string")),MakeString("gogen.cl:29"))
  
  _ = Core.F_attach_method(C_Generate_finish_block.AddMethod(Signature(C_string.Id(),C_void.Id()),0,MakeFunction1(E_Generate_finish_block_string,"Generate_finish_block_string")),MakeString("gogen.cl:33"))
  
  _ = Core.F_attach_method(C_Generate_genvar.AddMethod(Signature(C_string.Id(),C_string.Id()),0,MakeFunction1(E_Generate_genvar_string,"Generate_genvar_string")),MakeString("gogen.cl:38"))
  
  C_Generate_GO_PRODUCER = ToGenerateGoProducer(new(GenerateGoProducer).IsNamed(C_Generate_go_producer,MakeSymbol("GO_PRODUCER",It)))
  C_Generate_GO_PRODUCER.OpenComparators = MakeConstantList(C__inf.Id(),
    C__sup.Id(),
    C__sup_equal.Id(),
    C__inf_equal.Id())
  C_Generate_GO_PRODUCER.OpenOperators = MakeConstantList(Core.C__plus.Id(),
    C__dash.Id(),
    C__star.Id(),
    Core.C__sup_sup.Id())
  C_Generate_GO_PRODUCER.DivOperators = MakeConstantList(C__7.Id(),C_mod.Id())
  C_Generate_GO_PRODUCER.Extension = MakeString(".go")
  C_Generate_GO_PRODUCER.Comment = MakeString("Go")
  { 
    var va_arg1 *GenerateGoProducer  
    var va_arg2 *ClaireList  
    va_arg1 = C_Generate_GO_PRODUCER
    { 
      va_arg2= ToType(CEMPTY.Id()).EmptyList()
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("do")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("if")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("and")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("or")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("not")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("printf")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("void")).Id())
      va_arg2.AddFast(Optimize.C_Optimize_Pattern.Name.Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("return")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("new")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("default")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("private")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("operator")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("EID")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("Handle")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("import")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("var")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("catch")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("stdout")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("stdin")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("break")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("char")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("interface")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("EOF")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("System")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("delete")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("package")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("abstract")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("final")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("system_object")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("range")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("register")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("template")).Id())} 
    va_arg1.BadNames = va_arg2
    } 
  { 
    var va_arg1 *GenerateGoProducer  
    var va_arg2 *ClaireList  
    va_arg1 = C_Generate_GO_PRODUCER
    { 
      va_arg2= ToType(CEMPTY.Id()).EmptyList()
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("DO")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("IF")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("ClaireAnd")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("ClaireOr")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("NOT")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("PRINTF")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("ClaireVoid")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("ClairePattern")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("RETURN")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("NEW")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("Default")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("PRIVATE")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("ClaireOperator")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("ClaireEID")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("ClaireHandle")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("ClaireImport")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("ClaireVar")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("CATCH")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("STDOUT")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("STDIN")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("BREAK")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("ClaireChar")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("ClaireInterface")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("_eof")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("Core")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("Delete")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("ClairePackage")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("ABSTRACT")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("Final")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("SystemObject")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("Range")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("ClaireRegister")).Id())
      va_arg2.AddFast(Core.F_symbol_I_string2(MakeString("ClaireTemplate")).Id())} 
    va_arg1.GoodNames = va_arg2
    } 
  C_Generate_GO_PRODUCER.Interfaces = MakeConstantList(C_integer.Id(),
    MakeString("int").Id(),
    C_char.Id(),
    MakeString("rune").Id(),
    C_string.Id(),
    MakeString("string").Id(),
    C_float.Id(),
    MakeString("float64 ").Id())
  C_Generate_GO_PRODUCER.KernelMethods = MakeList(ToType(C_any.Id()),Core.F__at_property1(C_nth,C_list).Id(),
    Core.F__at_property1(C_nth,C_tuple).Id(),
    Core.F__at_property1(ToProperty(Core.C__at.Id()),C_type).Id(),
    MakeString("At").Id(),
    Core.F__at_property1(C_array_I,C_list).Id(),
    Core.F__at_property1(C_list_I,C_set).Id(),
    Core.F__at_property1(C_set_I,C_list).Id(),
    Core.F__at_property1(C_tuple_I,C_list).Id(),
    Core.F__at_property1(C_list_I,C_tuple).Id(),
    Core.F__at_property1(ToProperty(C__7_plus.Id()),C_list).Id(),
    MakeString("Append").Id(),
    Core.F__at_property1(ToProperty(Core.C__inf_inf.Id()),C_list).Id(),
    MakeString("Skip").Id())
  
  
  Optimize.C_PRODUCER.Value = C_Generate_GO_PRODUCER.Id()
  
  _ = Core.F_attach_method(C_Generate_c_string.AddMethod(Signature(C_Generate_go_producer.Id(),C_Variable.Id(),C_string.Id()),1,MakeFunction2(E_Generate_c_string_go_producer1,"Generate_c_string_go_producer1")),MakeString("gogen.cl:90"))
  
  _ = Core.F_attach_method(C_Generate_c_string.AddMethod(Signature(C_Generate_go_producer.Id(),C_symbol.Id(),C_string.Id()),1,MakeFunction2(E_Generate_c_string_go_producer2,"Generate_c_string_go_producer2")),MakeString("gogen.cl:93"))
  
  _ = Core.F_attach_method(Language.C_iClaire_ident.AddMethod(Signature(C_Generate_go_producer.Id(),C_Variable.Id(),C_void.Id()),0,MakeFunction2(E_iClaire_ident_go_producer1,"iClaire_ident_go_producer1")),MakeString("gogen.cl:101"))
  
  _ = Core.F_attach_method(Language.C_iClaire_ident.AddMethod(Signature(C_Generate_go_producer.Id(),C_symbol.Id(),C_void.Id()),0,MakeFunction2(E_iClaire_ident_go_producer2,"iClaire_ident_go_producer2")),MakeString("gogen.cl:108"))
  
  _ = Core.F_attach_method(C_Generate_cap_ident.AddMethod(Signature(C_symbol.Id(),C_void.Id()),0,MakeFunction1(E_Generate_cap_ident_symbol,"Generate_cap_ident_symbol")),MakeString("gogen.cl:114"))
  
  _ = Core.F_attach_method(C_Generate_capitalized_ident.AddMethod(Signature(C_symbol.Id(),C_module.Id(),C_void.Id()),0,MakeFunction2(E_Generate_capitalized_ident_symbol,"Generate_capitalized_ident_symbol")),MakeString("gogen.cl:124"))
  
  _ = Core.F_attach_method(C_Generate_cap_short.AddMethod(Signature(C_symbol.Id(),C_void.Id()),0,MakeFunction1(E_Generate_cap_short_symbol,"Generate_cap_short_symbol")),MakeString("gogen.cl:128"))
  
  _ = Core.F_attach_method(C_Generate_go_class.AddMethod(Signature(C_class.Id(),C_void.Id()),0,MakeFunction1(E_Generate_go_class_class,"Generate_go_class_class")),MakeString("gogen.cl:139"))
  
  _ = Core.F_attach_method(C_Generate_cast_class.AddMethod(Signature(C_class.Id(),C_void.Id()),0,MakeFunction1(E_Generate_cast_class_class,"Generate_cast_class_class")),MakeString("gogen.cl:149"))
  
  _ = Core.F_attach_method(C_Generate_class_ident.AddMethod(Signature(C_class.Id(),C_void.Id()),0,MakeFunction1(E_Generate_class_ident_class,"Generate_class_ident_class")),MakeString("gogen.cl:153"))
  
  _ = Core.F_attach_method(C_Generate_thing_ident.AddMethod(Signature(C_thing.Id(),C_void.Id()),0,MakeFunction1(E_Generate_thing_ident_thing,"Generate_thing_ident_thing")),MakeString("gogen.cl:157"))
  
  _ = Core.F_attach_method(C_Generate_symbol_ident.AddMethod(Signature(C_symbol.Id(),C_void.Id()),0,MakeFunction1(E_Generate_symbol_ident_symbol,"Generate_symbol_ident_symbol")),MakeString("gogen.cl:164"))
  
  _ = Core.F_attach_method(C_Generate_go_var.AddMethod(Signature(C_symbol.Id(),C_void.Id()),0,MakeFunction1(E_Generate_go_var_symbol,"Generate_go_var_symbol")),MakeString("gogen.cl:172"))
  
  _ = Core.F_attach_method(C_Generate_add_underscore.AddMethod(Signature(C_symbol.Id(),C_void.Id()),0,MakeFunction1(E_Generate_add_underscore_symbol,"Generate_add_underscore_symbol")),MakeString("gogen.cl:178"))
  
  _ = Core.F_attach_method(C_Generate_capitalize.AddMethod(Signature(C_string.Id(),C_string.Id()),0,MakeFunction1(E_Generate_capitalize_string,"Generate_capitalize_string")),MakeString("gogen.cl:189"))
  
  _ = Core.F_attach_method(C_Generate_capitalize.AddMethod(Signature(C_char.Id(),C_char.Id()),0,MakeFunction1(E_Generate_capitalize_char,"Generate_capitalize_char")),MakeString("gogen.cl:193"))
  
  _ = Core.F_attach_method(C_Generate_capitalize.AddMethod(Signature(C_symbol.Id(),C_string.Id()),0,MakeFunction1(E_Generate_capitalize_symbol,"Generate_capitalize_symbol")),MakeString("gogen.cl:195"))
  
  _ = Core.F_attach_method(C_Generate_globalVar.AddMethod(Signature(C_Generate_go_producer.Id(),Core.C_global_variable.Id(),C_void.Id()),0,MakeFunction2(E_Generate_globalVar_go_producer,"Generate_globalVar_go_producer")),MakeString("gogen.cl:202"))
  
  _ = Core.F_attach_method(C_Generate_type_sort.AddMethod(Signature(C_type.Id(),C_class.Id()),0,MakeFunction1(E_Generate_type_sort_type,"Generate_type_sort_type")),MakeString("gogen.cl:214"))
  
  _ = Core.F_attach_method(C_Generate_g_sort.AddMethod(Signature(C_any.Id(),C_class.Id()),1,MakeFunction1(E_Generate_g_sort_any,"Generate_g_sort_any")),MakeString("gogen.cl:218"))
  
  _ = Core.F_attach_method(C_Generate_valuesSlot.AddMethod(Signature(C_class.Id(),C_void.Id()),0,MakeFunction1(E_Generate_valuesSlot_class,"Generate_valuesSlot_class")),MakeString("gogen.cl:222"))
  
  _ = Core.F_attach_method(C_Generate_namespace_I.AddMethod(Signature(C_Generate_go_producer.Id(),C_module.Id(),C_void.Id()),0,MakeFunction2(E_Generate_namespace_I_go_producer,"Generate_namespace!_go_producer")),MakeString("gogen.cl:234"))
  
  _ = Core.F_attach_method(C_Generate_var_declaration.AddMethod(Signature(C_string.Id(),
    C_class.Id(),
    C_integer.Id(),
    C_void.Id()),0,MakeFunction3(E_Generate_var_declaration_string,"Generate_var_declaration_string")),MakeString("gogen.cl:245"))
  
  _ = Core.F_attach_method(C_imported_function_ask.AddMethod(Signature(C_any.Id(),C_boolean.Id()),0,MakeFunction1(E_imported_function_ask_any,"imported_function?_any")),MakeString("gogen.cl:249"))
  
  _ = Core.F_attach_method(C_Generate_goMethod_ask.AddMethod(Signature(C_any.Id(),C_boolean.Id()),0,MakeFunction1(E_Generate_goMethod_ask_any,"Generate_goMethod?_any")),MakeString("gogen.cl:275"))
  
  _ = Core.F_attach_method(C_dMethod_ask.AddMethod(Signature(C_any.Id(),C_void.Id()),1,MakeFunction1(E_dMethod_ask_any,"dMethod?_any")),MakeString("gogen.cl:292"))
  
  _ = Core.F_attach_method(C_Generate_arg_match.AddMethod(Signature(Core.F_param_I_class(C_list,ToType(C_class.Id())).Id(),Core.F_param_I_class(C_list,ToType(C_class.Id())).Id(),C_boolean.Id()),0,MakeFunction2(E_Generate_arg_match_list,"Generate_arg_match_list")),MakeString("gogen.cl:298"))
  
  _ = Core.F_attach_method(Optimize.C_Compile_function_name.AddMethod(Signature(C_property.Id(),C_list.Id(),C_string.Id()),0,MakeFunction2(E_Compile_function_name_property2,"Compile_function_name_property2")),MakeString("gogen.cl:312"))
  
  _ = Core.F_attach_method(C_Generate_at.AddMethod(Signature(C_Generate_go_producer.Id(),C_void.Id()),0,MakeFunction1(E_Generate_at_go_producer,"Generate_at_go_producer")),MakeString("gogen.cl:316"))
  
  _ = Core.F_attach_method(C_Generate_goVariables.AddMethod(Signature(C_Generate_go_producer.Id(),C_list.Id(),C_any.Id()),0,MakeFunction2(E_Generate_goVariables_go_producer,"Generate_goVariables_go_producer")),MakeString("gogen.cl:323"))
  
  _ = Core.F_attach_method(C_Generate_goVariable.AddMethod(Signature(C_Generate_go_producer.Id(),C_Variable.Id(),C_void.Id()),0,MakeFunction2(E_Generate_goVariable_go_producer,"Generate_goVariable_go_producer")),MakeString("gogen.cl:327"))
  
  _ = Core.F_attach_method(C_Generate_goMethod.AddMethod(Signature(C_method.Id(),C_void.Id()),0,MakeFunction1(E_Generate_goMethod_method,"Generate_goMethod_method")),MakeString("gogen.cl:336"))
  
  _ = Core.F_attach_method(C_Generate_goFunction.AddMethod(Signature(C_method.Id(),C_void.Id()),1,MakeFunction1(E_Generate_goFunction_method,"Generate_goFunction_method")),MakeString("gogen.cl:347"))
  
  _ = Core.F_attach_method(C_Generate_preCore_ask.AddMethod(Signature(C_void.Id(),C_void.Id()),0,MakeFunction1(E_Generate_preCore_ask_void,"Generate_preCore?_void")),MakeString("gogen.cl:352"))
  
  _ = Core.F_attach_method(C_Generate_goEIDFunction.AddMethod(Signature(C_method.Id(),C_void.Id()),1,MakeFunction1(E_Generate_goEIDFunction_method,"Generate_goEIDFunction_method")),MakeString("gogen.cl:358"))
  
  _ = Core.F_attach_method(C_Generate_goEIDFunctionName.AddMethod(Signature(C_method.Id(),C_void.Id()),1,MakeFunction1(E_Generate_goEIDFunctionName_method,"Generate_goEIDFunctionName_method")),MakeString("gogen.cl:363"))
  
  _ = Core.F_attach_method(C_Generate_goEvalFunction.AddMethod(Signature(C_method.Id(),C_void.Id()),0,MakeFunction1(E_Generate_goEvalFunction_method,"Generate_goEvalFunction_method")),MakeString("gogen.cl:370"))
  
  { 
    var expr EID 
    { var _CL_obj *Core.GlobalVariable  
      C_Generate__starlength_string_star = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("*length_string*",It)))
      
      _CL_obj = C_Generate__starlength_string_star
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = Core.F__at_property1(C_length,C_string).Id()
      expr = _CL_obj.Close()
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID 
    { var _CL_obj *Core.GlobalVariable  
      C_Generate__starset_I_list_star = Core.ToGlobalVariable(new(Core.GlobalVariable).IsNamed(Core.C_global_variable,MakeSymbol("*set!_list*",It)))
      
      _CL_obj = C_Generate__starset_I_list_star
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = Core.F__at_property1(C_set_I,C_list).Id()
      expr = _CL_obj.Close()
      } 
    ErrorCheck(expr)} 
  
  _ = Core.F_attach_method(C_Generate_getFunctionName.AddMethod(Signature(C_method.Id(),C_string.Id()),1,MakeFunction1(E_Generate_getFunctionName_method,"Generate_getFunctionName_method")),MakeString("gogen.cl:385"))
  
  _ = Core.F_attach_method(C_Generate_retreive_method.AddMethod(Signature(C_any.Id(),C_any.Id(),C_method.Id()),1,MakeFunction2(E_Generate_retreive_method_any,"Generate_retreive_method_any")),MakeString("gogen.cl:394"))
  
  _ = Core.F_attach_method(C_Generate_retreive_list.AddMethod(Signature(C_any.Id(),C_any.Id()),1,MakeFunction1(E_Generate_retreive_list_any,"Generate_retreive_list_any")),MakeString("gogen.cl:419"))
  
  _ = Core.F_attach_method(C_Generate_interface_I.AddMethod(Signature(C_class.Id(),C_void.Id()),0,MakeFunction1(E_Generate_interface_I_class,"Generate_interface!_class")),MakeString("gogen.cl:435"))
  
  _ = Core.F_attach_method(C_Generate_to_cl.AddMethod(Signature(C_Generate_go_producer.Id(),
    C_any.Id(),
    C_class.Id(),
    C_void.Id()),1,MakeFunction3(E_Generate_to_cl_go_producer,"Generate_to_cl_go_producer")),MakeString("gogen.cl:446"))
  
  _ = Core.F_attach_method(C_Generate_to_eid.AddMethod(Signature(C_Generate_go_producer.Id(),
    C_any.Id(),
    C_class.Id(),
    C_void.Id()),1,MakeFunction3(E_Generate_to_eid_go_producer,"Generate_to_eid_go_producer")),MakeString("gogen.cl:466"))
  
  _ = Core.F_attach_method(C_Generate_from_eid.AddMethod(Signature(C_Generate_go_producer.Id(),
    C_string.Id(),
    C_class.Id(),
    C_void.Id()),1,MakeFunction3(E_Generate_from_eid_go_producer,"Generate_from_eid_go_producer")),MakeString("gogen.cl:473"))
  
  _ = Core.F_attach_method(C_Generate_eid_prefix.AddMethod(Signature(C_class.Id(),C_void.Id()),1,MakeFunction1(E_Generate_eid_prefix_class,"Generate_eid_prefix_class")),MakeString("gogen.cl:484"))
  
  _ = Core.F_attach_method(C_Generate_eid_post.AddMethod(Signature(C_class.Id(),C_void.Id()),0,MakeFunction1(E_Generate_eid_post_class,"Generate_eid_post_class")),MakeString("gogen.cl:490"))
  
  _ = Core.F_attach_method(C_Generate_integer_prefix.AddMethod(Signature(C_class.Id(),C_void.Id()),0,MakeFunction1(E_Generate_integer_prefix_class,"Generate_integer_prefix_class")),MakeString("gogen.cl:495"))
  
  _ = Core.F_attach_method(C_Generate_float_prefix.AddMethod(Signature(C_class.Id(),C_void.Id()),0,MakeFunction1(E_Generate_float_prefix_class,"Generate_float_prefix_class")),MakeString("gogen.cl:500"))
  
  _ = Core.F_attach_method(C_Generate_char_prefix.AddMethod(Signature(C_class.Id(),C_void.Id()),0,MakeFunction1(E_Generate_char_prefix_class,"Generate_char_prefix_class")),MakeString("gogen.cl:505"))
  
  _ = Core.F_attach_method(C_Generate_string_prefix.AddMethod(Signature(C_class.Id(),C_void.Id()),0,MakeFunction1(E_Generate_string_prefix_class,"Generate_string_prefix_class")),MakeString("gogen.cl:510"))
  
  _ = Core.F_attach_method(C_Generate_string_post.AddMethod(Signature(C_class.Id(),C_void.Id()),0,MakeFunction1(E_Generate_string_post_class,"Generate_string_post_class")),MakeString("gogen.cl:514"))
  
  _ = Core.F_attach_method(C_Generate_native_post.AddMethod(Signature(C_class.Id(),C_void.Id()),0,MakeFunction1(E_Generate_native_post_class,"Generate_native_post_class")),MakeString("gogen.cl:519"))
  
  _ = Core.F_attach_method(C_Generate_object_prefix.AddMethod(Signature(C_class.Id(),C_class.Id(),C_void.Id()),0,MakeFunction2(E_Generate_object_prefix_class,"Generate_object_prefix_class")),MakeString("gogen.cl:528"))
  
  _ = Core.F_attach_method(C_Generate_object_post.AddMethod(Signature(C_class.Id(),C_class.Id(),C_void.Id()),0,MakeFunction2(E_Generate_object_post_class,"Generate_object_post_class")),MakeString("gogen.cl:541"))
  
  _ = Core.F_attach_method(C_Generate_cast_prefix.AddMethod(Signature(C_class.Id(),C_class.Id(),C_void.Id()),1,MakeFunction2(E_Generate_cast_prefix_class,"Generate_cast_prefix_class")),MakeString("gogen.cl:551"))
  
  _ = Core.F_attach_method(C_Generate_cast_post.AddMethod(Signature(C_class.Id(),C_class.Id(),C_void.Id()),0,MakeFunction2(E_Generate_cast_post_class,"Generate_cast_post_class")),MakeString("gogen.cl:560"))
  
  _ = Core.F_attach_method(C_Generate_equal_exp.AddMethod(Signature(C_Generate_go_producer.Id(),
    C_any.Id(),
    C_boolean.Id(),
    C_any.Id(),
    C_any.Id(),
    C_void.Id()),1,MakeFunction5(E_Generate_equal_exp_go_producer,"Generate_equal_exp_go_producer")),MakeString("gogen.cl:589"))
  
  _ = Core.F_attach_method(C_Generate_char_exp_ask.AddMethod(Signature(C_Generate_go_producer.Id(),C_any.Id(),C_boolean.Id()),0,MakeFunction2(E_Generate_char_exp_ask_go_producer,"Generate_char_exp?_go_producer")),MakeString("gogen.cl:598"))
  
  _ = Core.F_attach_method(C_Generate_c_member.AddMethod(Signature(C_Generate_go_producer.Id(),
    C_any.Id(),
    C_class.Id(),
    C_property.Id(),
    C_void.Id()),1,MakeFunction4(E_Generate_c_member_go_producer,"Generate_c_member_go_producer")),MakeString("gogen.cl:602"))
  
  _ = Core.F_attach_method(C_Generate_bag_expression.AddMethod(Signature(C_Generate_go_producer.Id(),
    C_class.Id(),
    C_list.Id(),
    C_type.Id(),
    C_void.Id()),1,MakeFunction4(E_Generate_bag_expression_go_producer,"Generate_bag_expression_go_producer")),MakeString("gogen.cl:614"))
  
  _ = Core.F_attach_method(C_Generate_constant_ask.AddMethod(Signature(C_any.Id(),C_boolean.Id()),0,MakeFunction1(E_Generate_constant_ask_any,"Generate_constant?_any")),MakeString("gogen.cl:626"))
  
  _ = Core.F_attach_method(C_Generate_go_range.AddMethod(Signature(C_Variable.Id(),C_class.Id()),0,MakeFunction1(E_Generate_go_range_Variable,"Generate_go_range_Variable")),MakeString("gogen.cl:630"))
  
  _ = Core.F_attach_method(C_Generate_go_signature.AddMethod(Signature(C_method.Id(),Core.F_param_I_class(C_list,ToType(C_class.Id())).Id()),0,MakeFunction1(E_Generate_go_signature_method,"Generate_go_signature_method")),MakeString("gogen.cl:634"))
  
  _ = Core.F_attach_method(C_Generate_full_signature.AddMethod(Signature(C_method.Id(),Core.F_param_I_class(C_list,ToType(C_type.Id())).Id()),0,MakeFunction1(E_Generate_full_signature_method,"Generate_full_signature_method")),MakeString("gogen.cl:638"))
  
  _ = Core.F_attach_method(C_Generate_signature_I.AddMethod(Signature(C_Generate_go_producer.Id(),Core.F_param_I_class(C_list,ToType(C_type.Id())).Id(),C_void.Id()),1,MakeFunction2(E_Generate_signature_I_go_producer,"Generate_signature!_go_producer")),MakeString("gogen.cl:642"))
  
  _ = Core.F_attach_method(C_Generate_g_member.AddMethod(Signature(C_any.Id(),C_class.Id()),1,MakeFunction1(E_Generate_g_member_any,"Generate_g_member_any")),MakeString("gogen.cl:651"))
  
  _ = Core.F_attach_method(C_Generate_cast_Values.AddMethod(Signature(C_class.Id(),C_class.Id(),C_void.Id()),0,MakeFunction2(E_Generate_cast_Values_class,"Generate_cast_Values_class")),MakeString("gogen.cl:656"))
  
  _ = Core.F_attach_method(C_Generate_check_var.AddMethod(Signature(C_string.Id(),C_string.Id()),0,MakeFunction1(E_Generate_check_var_string,"Generate_check_var_string")),MakeString("gogen.cl:662"))
  
  _ = Core.F_attach_method(C_Generate_build_Variable.AddMethod(Signature(C_string.Id(),C_any.Id(),C_Variable.Id()),0,MakeFunction2(E_Generate_build_Variable_string,"Generate_build_Variable_string")),MakeString("gogen.cl:665"))
  
  _ = Core.F_attach_method(C_Generate_use_variable.AddMethod(Signature(C_string.Id(),
    C_class.Id(),
    C_class.Id(),
    C_void.Id()),1,MakeFunction3(E_Generate_use_variable_string,"Generate_use_variable_string")),MakeString("gogen.cl:671"))
  
  _ = Core.F_attach_method(C_Generate_g_clean.AddMethod(Signature(C_any.Id(),C_boolean.Id()),1,MakeFunction1(E_Generate_g_clean_any,"Generate_g_clean_any")),MakeString("gogen.cl:675"))
  
  _ = Core.F_attach_method(C_Generate_simple_func_ask.AddMethod(Signature(C_any.Id(),C_boolean.Id()),1,MakeFunction1(E_Generate_simple_func_ask_any,"Generate_simple_func?_any")),MakeString("gogen.cl:680"))
  
  _ = Core.F_attach_method(C_Generate_g_func.AddMethod(Signature(C_any.Id(),C_boolean.Id()),1,MakeFunction1(E_Generate_g_func_any,"Generate_g_func_any")),MakeString("goexp.cl:69"))
  
  _ = Core.F_attach_method(C_Generate_g_expression.AddMethod(Signature(C_any.Id(),C_class.Id(),C_void.Id()),1,MakeFunction2(E_Generate_g_expression_any,"Generate_g_expression_any")),MakeString("goexp.cl:76"))
  
  _ = Core.F_attach_method(C_Generate_g_expression.AddMethod(Signature(C_thing.Id(),C_class.Id(),C_void.Id()),1,MakeFunction2(E_Generate_g_expression_thing,"Generate_g_expression_thing")),MakeString("goexp.cl:83"))
  
  _ = Core.F_attach_method(C_Generate_g_expression.AddMethod(Signature(C_module.Id(),C_class.Id(),C_void.Id()),1,MakeFunction2(E_Generate_g_expression_module,"Generate_g_expression_module")),MakeString("goexp.cl:98"))
  
  _ = Core.F_attach_method(C_Generate_g_expression.AddMethod(Signature(C_class.Id(),C_class.Id(),C_void.Id()),1,MakeFunction2(E_Generate_g_expression_class,"Generate_g_expression_class")),MakeString("goexp.cl:104"))
  
  _ = Core.F_attach_method(C_Generate_g_expression.AddMethod(Signature(C_boolean.Id(),C_class.Id(),C_void.Id()),1,MakeFunction2(E_Generate_g_expression_boolean,"Generate_g_expression_boolean")),MakeString("goexp.cl:109"))
  
  _ = Core.F_attach_method(C_Generate_g_expression.AddMethod(Signature(C_integer.Id(),C_class.Id(),C_void.Id()),1,MakeFunction2(E_Generate_g_expression_integer,"Generate_g_expression_integer")),MakeString("goexp.cl:118"))
  
  _ = Core.F_attach_method(C_Generate_g_expression.AddMethod(Signature(C_float.Id(),C_class.Id(),C_void.Id()),1,MakeFunction2(E_Generate_g_expression_float,"Generate_g_expression_float")),MakeString("goexp.cl:125"))
  
  _ = Core.F_attach_method(C_Generate_g_expression.AddMethod(Signature(C_char.Id(),C_class.Id(),C_void.Id()),1,MakeFunction2(E_Generate_g_expression_char,"Generate_g_expression_char")),MakeString("goexp.cl:132"))
  
  _ = Core.F_attach_method(C_Generate_g_expression.AddMethod(Signature(C_string.Id(),C_class.Id(),C_void.Id()),1,MakeFunction2(E_Generate_g_expression_string,"Generate_g_expression_string")),MakeString("goexp.cl:137"))
  
  _ = Core.F_attach_method(C_Generate_g_expression.AddMethod(Signature(C_symbol.Id(),C_class.Id(),C_void.Id()),1,MakeFunction2(E_Generate_g_expression_symbol,"Generate_g_expression_symbol")),MakeString("goexp.cl:145"))
  
  _ = Core.F_attach_method(C_Generate_g_expression.AddMethod(Signature(C_environment.Id(),C_class.Id(),C_void.Id()),1,MakeFunction2(E_Generate_g_expression_environment,"Generate_g_expression_environment")),MakeString("goexp.cl:149"))
  
  _ = Core.F_attach_method(C_Generate_g_expression.AddMethod(Signature(C_function.Id(),C_class.Id(),C_void.Id()),1,MakeFunction2(E_Generate_g_expression_function,"Generate_g_expression_function")),MakeString("goexp.cl:157"))
  
  _ = Core.F_attach_method(C_Generate_g_expression.AddMethod(Signature(C_Variable.Id(),C_class.Id(),C_void.Id()),1,MakeFunction2(E_Generate_g_expression_Variable,"Generate_g_expression_Variable")),MakeString("goexp.cl:170"))
  
  _ = Core.F_attach_method(C_Generate_g_expression.AddMethod(Signature(Core.C_global_variable.Id(),C_class.Id(),C_void.Id()),1,MakeFunction2(E_Generate_g_expression_global_variable,"Generate_g_expression_global_variable")),MakeString("goexp.cl:181"))
  
  _ = Core.F_attach_method(C_Generate_g_expression.AddMethod(Signature(Language.C_Set.Id(),C_class.Id(),C_void.Id()),1,MakeFunction2(E_Generate_g_expression_Set,"Generate_g_expression_Set")),MakeString("goexp.cl:188"))
  
  _ = Core.F_attach_method(C_Generate_g_expression.AddMethod(Signature(C_set.Id(),C_class.Id(),C_void.Id()),1,MakeFunction2(E_Generate_g_expression_set,"Generate_g_expression_set")),MakeString("goexp.cl:196"))
  
  _ = Core.F_attach_method(C_Generate_g_expression.AddMethod(Signature(Language.C_Tuple.Id(),C_class.Id(),C_void.Id()),1,MakeFunction2(E_Generate_g_expression_Tuple,"Generate_g_expression_Tuple")),MakeString("goexp.cl:203"))
  
  _ = Core.F_attach_method(C_Generate_g_expression.AddMethod(Signature(C_tuple.Id(),C_class.Id(),C_void.Id()),1,MakeFunction2(E_Generate_g_expression_tuple,"Generate_g_expression_tuple")),MakeString("goexp.cl:209"))
  
  _ = Core.F_attach_method(C_Generate_g_expression.AddMethod(Signature(Language.C_List.Id(),C_class.Id(),C_void.Id()),1,MakeFunction2(E_Generate_g_expression_List,"Generate_g_expression_List")),MakeString("goexp.cl:216"))
  
  _ = Core.F_attach_method(C_Generate_g_expression.AddMethod(Signature(C_list.Id(),C_class.Id(),C_void.Id()),1,MakeFunction2(E_Generate_g_expression_list,"Generate_g_expression_list")),MakeString("goexp.cl:224"))
  
  _ = Core.F_attach_method(C_Generate_g_expression.AddMethod(Signature(C_lambda.Id(),C_class.Id(),C_void.Id()),1,MakeFunction2(E_Generate_g_expression_lambda,"Generate_g_expression_lambda")),MakeString("goexp.cl:233"))
  
  _ = Core.F_attach_method(C_Generate_g_expression.AddMethod(Signature(Language.C_Call.Id(),C_class.Id(),C_void.Id()),1,MakeFunction2(E_Generate_g_expression_Call,"Generate_g_expression_Call")),MakeString("goexp.cl:242"))
  
  _ = Core.F_attach_method(C_Generate_g_expression.AddMethod(Signature(Language.C_Call_method1.Id(),C_class.Id(),C_void.Id()),1,MakeFunction2(E_Generate_g_expression_Call_method1,"Generate_g_expression_Call_method1")),MakeString("goexp.cl:245"))
  
  _ = Core.F_attach_method(C_Generate_g_expression.AddMethod(Signature(Language.C_Call_method2.Id(),C_class.Id(),C_void.Id()),1,MakeFunction2(E_Generate_g_expression_Call_method2,"Generate_g_expression_Call_method2")),MakeString("goexp.cl:246"))
  
  _ = Core.F_attach_method(C_Generate_g_expression.AddMethod(Signature(Language.C_Call_method.Id(),C_class.Id(),C_void.Id()),1,MakeFunction2(E_Generate_g_expression_Call_method,"Generate_g_expression_Call_method")),MakeString("goexp.cl:247"))
  
  _ = Core.F_attach_method(C_Generate_inline_exp.AddMethod(Signature(C_Generate_go_producer.Id(),
    Language.C_Call.Id(),
    C_class.Id(),
    C_void.Id()),1,MakeFunction3(E_Generate_inline_exp_go_producer1,"Generate_inline_exp_go_producer1")),MakeString("goexp.cl:314"))
  
  _ = Core.F_attach_method(C_Generate_args_list.AddMethod(Signature(C_list.Id(),C_class.Id(),C_void.Id()),1,MakeFunction2(E_Generate_args_list_list,"Generate_args_list_list")),MakeString("goexp.cl:324"))
  
  _ = Core.F_attach_method(C_Generate_g_expression.AddMethod(Signature(Language.C_Super.Id(),C_class.Id(),C_void.Id()),1,MakeFunction2(E_Generate_g_expression_Super,"Generate_g_expression_Super")),MakeString("goexp.cl:335"))
  
  _ = Core.F_attach_method(C_Generate_inline_exp.AddMethod(Signature(C_Generate_go_producer.Id(),
    Language.C_Call_method1.Id(),
    C_class.Id(),
    C_void.Id()),1,MakeFunction3(E_Generate_inline_exp_go_producer2,"Generate_inline_exp_go_producer2")),MakeString("goexp.cl:385"))
  
  _ = Core.F_attach_method(C_Generate_inline_exp.AddMethod(Signature(C_Generate_go_producer.Id(),
    Language.C_Call_method2.Id(),
    C_class.Id(),
    C_void.Id()),1,MakeFunction3(E_Generate_inline_exp_go_producer3,"Generate_inline_exp_go_producer3")),MakeString("goexp.cl:480"))
  
  _ = Core.F_attach_method(C_Generate_inline_exp.AddMethod(Signature(C_Generate_go_producer.Id(),
    Language.C_Call_method.Id(),
    C_class.Id(),
    C_void.Id()),1,MakeFunction3(E_Generate_inline_exp_go_producer4,"Generate_inline_exp_go_producer4")),MakeString("goexp.cl:523"))
  
  _ = Core.F_attach_method(C_Generate_print_external_call.AddMethod(Signature(C_Generate_go_producer.Id(),
    Language.C_Call_method.Id(),
    C_class.Id(),
    C_void.Id()),1,MakeFunction3(E_Generate_print_external_call_go_producer,"Generate_print_external_call_go_producer")),MakeString("goexp.cl:545"))
  
  _ = Core.F_attach_method(C_Generate_external_casted_arg.AddMethod(Signature(C_any.Id(),
    C_class.Id(),
    C_integer.Id(),
    C_boolean.Id(),
    C_void.Id()),1,MakeFunction4(E_Generate_external_casted_arg_any,"Generate_external_casted_arg_any")),MakeString("goexp.cl:554"))
  
  _ = Core.F_attach_method(C_Generate_bounded_expression.AddMethod(Signature(C_any.Id(),C_class.Id(),C_void.Id()),1,MakeFunction2(E_Generate_bounded_expression_any,"Generate_bounded_expression_any")),MakeString("goexp.cl:571"))
  
  _ = Core.F_attach_method(C_Generate_g_expression.AddMethod(Signature(Language.C_If.Id(),C_class.Id(),C_void.Id()),1,MakeFunction2(E_Generate_g_expression_If,"Generate_g_expression_If")),MakeString("goexp.cl:583"))
  
  _ = Core.F_attach_method(C_Generate_g_expression.AddMethod(Signature(Language.C_And.Id(),C_class.Id(),C_void.Id()),1,MakeFunction2(E_Generate_g_expression_And,"Generate_g_expression_And")),MakeString("goexp.cl:596"))
  
  _ = Core.F_attach_method(C_Generate_g_expression.AddMethod(Signature(Language.C_Or.Id(),C_class.Id(),C_void.Id()),1,MakeFunction2(E_Generate_g_expression_Or,"Generate_g_expression_Or")),MakeString("goexp.cl:608"))
  
  _ = Core.F_attach_method(C_Generate_g_expression.AddMethod(Signature(Optimize.C_Compile_C_cast.Id(),C_class.Id(),C_void.Id()),1,MakeFunction2(E_Generate_g_expression_C_cast,"Generate_g_expression_C_cast")),MakeString("goexp.cl:621"))
  
  _ = Core.F_attach_method(C_Generate_g_expression.AddMethod(Signature(Language.C_Call_slot.Id(),C_class.Id(),C_void.Id()),1,MakeFunction2(E_Generate_g_expression_Call_slot,"Generate_g_expression_Call_slot")),MakeString("goexp.cl:636"))
  
  _ = Core.F_attach_method(C_Generate_g_expression.AddMethod(Signature(Language.C_Call_table.Id(),C_class.Id(),C_void.Id()),1,MakeFunction2(E_Generate_g_expression_Call_table,"Generate_g_expression_Call_table")),MakeString("goexp.cl:656"))
  
  _ = Core.F_attach_method(C_Generate_g_table_index.AddMethod(Signature(C_table.Id(),C_any.Id(),C_void.Id()),1,MakeFunction2(E_Generate_g_table_index_table,"Generate_g_table_index_table")),MakeString("goexp.cl:667"))
  
  _ = Core.F_attach_method(C_Generate_g_expression.AddMethod(Signature(Language.C_Call_array.Id(),C_class.Id(),C_void.Id()),1,MakeFunction2(E_Generate_g_expression_Call_array,"Generate_g_expression_Call_array")),MakeString("goexp.cl:678"))
  
  _ = Core.F_attach_method(C_Generate_sign_equal.AddMethod(Signature(C_boolean.Id(),C_void.Id()),0,MakeFunction1(E_Generate_sign_equal_boolean,"Generate_sign_equal_boolean")),MakeString("goexp.cl:698"))
  
  _ = Core.F_attach_method(C_Generate_sign_or.AddMethod(Signature(C_boolean.Id(),C_void.Id()),0,MakeFunction1(E_Generate_sign_or_boolean,"Generate_sign_or_boolean")),MakeString("goexp.cl:701"))
  
  _ = Core.F_attach_method(Optimize.C_Compile_bool_exp.AddMethod(Signature(C_any.Id(),C_boolean.Id(),C_void.Id()),1,MakeFunction2(E_Compile_bool_exp_any,"Compile_bool_exp_any")),MakeString("goexp.cl:705"))
  
  _ = Core.F_attach_method(Optimize.C_Compile_bool_exp.AddMethod(Signature(Optimize.C_Compile_C_cast.Id(),C_boolean.Id(),C_void.Id()),1,MakeFunction2(E_Compile_bool_exp_C_cast,"Compile_bool_exp_C_cast")),MakeString("goexp.cl:709"))
  
  _ = Core.F_attach_method(Optimize.C_Compile_bool_exp.AddMethod(Signature(Language.C_If.Id(),C_boolean.Id(),C_void.Id()),1,MakeFunction2(E_Compile_bool_exp_If,"Compile_bool_exp_If")),MakeString("goexp.cl:723"))
  
  _ = Core.F_attach_method(Optimize.C_Compile_bool_exp.AddMethod(Signature(Language.C_And.Id(),C_boolean.Id(),C_void.Id()),1,MakeFunction2(E_Compile_bool_exp_And,"Compile_bool_exp_And")),MakeString("goexp.cl:736"))
  
  _ = Core.F_attach_method(Optimize.C_Compile_bool_exp.AddMethod(Signature(Language.C_Or.Id(),C_boolean.Id(),C_void.Id()),1,MakeFunction2(E_Compile_bool_exp_Or,"Compile_bool_exp_Or")),MakeString("goexp.cl:749"))
  
  _ = Core.F_attach_method(Optimize.C_Compile_bool_exp.AddMethod(Signature(Language.C_Call.Id(),C_boolean.Id(),C_void.Id()),1,MakeFunction2(E_Compile_bool_exp_Call,"Compile_bool_exp_Call")),MakeString("goexp.cl:756"))
  
  _ = Core.F_attach_method(C_Generate_belong_exp.AddMethod(Signature(C_any.Id(),
    C_any.Id(),
    C_class.Id(),
    C_void.Id()),1,MakeFunction3(E_Generate_belong_exp_any,"Generate_belong_exp_any")),MakeString("goexp.cl:776"))
  
  _ = Core.F_attach_method(Optimize.C_Compile_bool_exp.AddMethod(Signature(Language.C_Call_method1.Id(),C_boolean.Id(),C_void.Id()),1,MakeFunction2(E_Compile_bool_exp_Call_method1,"Compile_bool_exp_Call_method1")),MakeString("goexp.cl:787"))
  
  _ = Core.F_attach_method(Optimize.C_Compile_bool_exp.AddMethod(Signature(Language.C_Call_method2.Id(),C_boolean.Id(),C_void.Id()),1,MakeFunction2(E_Compile_bool_exp_Call_method2,"Compile_bool_exp_Call_method2")),MakeString("goexp.cl:811"))
  
  _ = Core.F_attach_method(C_Generate_unfold_args.AddMethod(Signature(C_list.Id(),C_list.Id()),1,MakeFunction1(E_Generate_unfold_args_list,"Generate_unfold_args_list")),MakeString("gostat.cl:57"))
  
  _ = Core.F_attach_method(C_Generate_unfold_arg.AddMethod(Signature(C_list.Id(),
    C_list.Id(),
    C_any.Id(),
    C_any.Id()),1,MakeFunction3(E_Generate_unfold_arg_list,"Generate_unfold_arg_list")),MakeString("gostat.cl:69"))
  
  _ = Core.F_attach_method(C_Generate_unfold_use.AddMethod(Signature(C_list.Id(),
    C_any.Id(),
    C_class.Id(),
    C_string.Id(),
    C_boolean.Id(),
    C_any.Id(),
    C_void.Id()),1,MakeFunction6(E_Generate_unfold_use_list,"Generate_unfold_use_list")),MakeString("gostat.cl:83"))
  
  _ = Core.F_attach_method(C_Generate_g_try.AddMethod(Signature(C_any.Id(),
    C_string.Id(),
    C_class.Id(),
    C_string.Id(),
    C_any.Id(),
    C_void.Id()),1,MakeFunction5(E_Generate_g_try_any,"Generate_g_try_any")),MakeString("gostat.cl:116"))
  
  _ = Core.F_attach_method(C_Generate_g_try.AddMethod(Signature(Language.C_Assign.Id(),
    C_string.Id(),
    C_class.Id(),
    C_string.Id(),
    C_any.Id(),
    C_void.Id()),1,MakeFunction5(E_Generate_g_try_Assign,"Generate_g_try_Assign")),MakeString("gostat.cl:123"))
  
  _ = Core.F_attach_method(C_Generate_close_try.AddMethod(Signature(C_integer.Id(),C_void.Id()),0,MakeFunction1(E_Generate_close_try_integer,"Generate_close_try_integer")),MakeString("gostat.cl:129"))
  
  _ = Core.F_attach_method(C_Generate_error_wrap.AddMethod(Signature(C_any.Id(),
    C_class.Id(),
    C_string.Id(),
    C_void.Id()),1,MakeFunction3(E_Generate_error_wrap_any,"Generate_error_wrap_any")),MakeString("gostat.cl:136"))
  
  _ = Core.F_attach_method(C_Generate_g_try_void.AddMethod(Signature(C_any.Id(),
    C_string.Id(),
    C_any.Id(),
    C_void.Id()),1,MakeFunction3(E_Generate_g_try_void_any,"Generate_g_try_void_any")),MakeString("gostat.cl:153"))
  
  _ = Core.F_attach_method(C_Generate_eid_require_ask.AddMethod(Signature(C_any.Id(),C_boolean.Id()),0,MakeFunction1(E_Generate_eid_require_ask_any,"Generate_eid_require?_any")),MakeString("gostat.cl:162"))
  
  _ = Core.F_attach_method(C_Generate_eid_provide_ask.AddMethod(Signature(C_any.Id(),C_boolean.Id()),0,MakeFunction1(E_Generate_eid_provide_ask_any,"Generate_eid_provide?_any")),MakeString("gostat.cl:171"))
  
  _ = Core.F_attach_method(C_Generate_unfold_eid.AddMethod(Signature(C_list.Id(),
    C_any.Id(),
    C_class.Id(),
    C_any.Id(),
    C_boolean.Id(),
    C_any.Id(),
    C_void.Id()),1,MakeFunction6(E_Generate_unfold_eid_list,"Generate_unfold_eid_list")),MakeString("gostat.cl:192"))
  
  _ = Core.F_attach_method(C_Generate_eid_expression.AddMethod(Signature(C_any.Id(),
    C_class.Id(),
    Core.F_param_I_class(C_list,ToType(C_Variable.Id())).Id(),
    C_void.Id()),1,MakeFunction3(E_Generate_eid_expression_any,"Generate_eid_expression_any")),MakeString("gostat.cl:220"))
  
  _ = Core.F_attach_method(C_Generate_statement.AddMethod(Signature(C_any.Id(),
    C_class.Id(),
    C_string.Id(),
    C_any.Id(),
    C_void.Id()),1,MakeFunction4(E_Generate_statement_any,"Generate_statement_any")),MakeString("gostat.cl:242"))
  
  _ = Core.F_attach_method(C_Generate_stat_exp.AddMethod(Signature(C_Generate_go_producer.Id(),
    C_any.Id(),
    C_class.Id(),
    C_void.Id()),1,MakeFunction3(E_Generate_stat_exp_go_producer,"Generate_stat_exp_go_producer")),MakeString("gostat.cl:252"))
  
  _ = Core.F_attach_method(C_Generate_g_statement.AddMethod(Signature(Language.C_Do.Id(),
    C_class.Id(),
    C_string.Id(),
    C_boolean.Id(),
    C_any.Id(),
    C_void.Id()),1,MakeFunction5(E_Generate_g_statement_Do,"Generate_g_statement_Do")),MakeString("gostat.cl:260"))
  
  _ = Core.F_attach_method(C_Generate_do_statement.AddMethod(Signature(Language.C_Do.Id(),
    C_class.Id(),
    C_string.Id(),
    C_boolean.Id(),
    C_any.Id(),
    C_boolean.Id(),
    C_void.Id()),1,MakeFunction6(E_Generate_do_statement_Do,"Generate_do_statement_Do")),MakeString("gostat.cl:278"))
  
  _ = Core.F_attach_method(C_Generate_g_statement.AddMethod(Signature(Language.C_Let.Id(),
    C_class.Id(),
    C_string.Id(),
    C_boolean.Id(),
    C_any.Id(),
    C_void.Id()),1,MakeFunction5(E_Generate_g_statement_Let,"Generate_g_statement_Let")),MakeString("gostat.cl:305"))
  
  _ = Core.F_attach_method(C_Generate_let_eid_ask.AddMethod(Signature(Language.C_Let.Id(),C_boolean.Id()),1,MakeFunction1(E_Generate_let_eid_ask_Let,"Generate_let_eid?_Let")),MakeString("gostat.cl:315"))
  
  _ = Core.F_attach_method(C_Generate_g_eid_stat.AddMethod(Signature(Language.C_Let.Id(),
    C_class.Id(),
    C_string.Id(),
    C_boolean.Id(),
    C_any.Id(),
    C_void.Id()),1,MakeFunction5(E_Generate_g_eid_stat_Let,"Generate_g_eid_stat_Let")),MakeString("gostat.cl:329"))
  
  _ = Core.F_attach_method(C_Generate_g_statement.AddMethod(Signature(Language.C_Construct.Id(),
    C_class.Id(),
    C_string.Id(),
    C_boolean.Id(),
    C_any.Id(),
    C_void.Id()),1,MakeFunction5(E_Generate_g_statement_Construct,"Generate_g_statement_Construct")),MakeString("gostat.cl:360"))
  
  _ = Core.F_attach_method(C_Generate_g_statement.AddMethod(Signature(Language.C_If.Id(),
    C_class.Id(),
    C_string.Id(),
    C_boolean.Id(),
    C_any.Id(),
    C_void.Id()),1,MakeFunction5(E_Generate_g_statement_If,"Generate_g_statement_If")),MakeString("gostat.cl:384"))
  
  _ = Core.F_attach_method(C_Generate_g_statement.AddMethod(Signature(Language.C_And.Id(),
    C_class.Id(),
    C_string.Id(),
    C_boolean.Id(),
    C_any.Id(),
    C_void.Id()),1,MakeFunction5(E_Generate_g_statement_And,"Generate_g_statement_And")),MakeString("gostat.cl:414"))
  
  _ = Core.F_attach_method(C_Generate_g_statement.AddMethod(Signature(Language.C_Or.Id(),
    C_class.Id(),
    C_string.Id(),
    C_boolean.Id(),
    C_any.Id(),
    C_void.Id()),1,MakeFunction5(E_Generate_g_statement_Or,"Generate_g_statement_Or")),MakeString("gostat.cl:438"))
  
  _ = Core.F_attach_method(C_Generate_g_statement.AddMethod(Signature(Language.C_Assign.Id(),
    C_class.Id(),
    C_string.Id(),
    C_boolean.Id(),
    C_any.Id(),
    C_void.Id()),1,MakeFunction5(E_Generate_g_statement_Assign,"Generate_g_statement_Assign")),MakeString("gostat.cl:453"))
  
  _ = Core.F_attach_method(C_Generate_g_statement.AddMethod(Signature(Language.C_Gassign.Id(),
    C_class.Id(),
    C_string.Id(),
    C_boolean.Id(),
    C_any.Id(),
    C_void.Id()),1,MakeFunction5(E_Generate_g_statement_Gassign,"Generate_g_statement_Gassign")),MakeString("gostat.cl:473"))
  
  _ = Core.F_attach_method(C_Generate_bag_class.AddMethod(Signature(C_any.Id(),C_class.Id()),1,MakeFunction1(E_Generate_bag_class_any,"Generate_bag_class_any")),MakeString("gostat.cl:487"))
  
  _ = Core.F_attach_method(C_Generate_g_statement.AddMethod(Signature(Language.C_For.Id(),
    C_class.Id(),
    C_string.Id(),
    C_boolean.Id(),
    C_any.Id(),
    C_void.Id()),1,MakeFunction5(E_Generate_g_statement_For,"Generate_g_statement_For")),MakeString("gostat.cl:517"))
  
  _ = Core.F_attach_method(C_Generate_iteration_statement.AddMethod(Signature(Language.C_For.Id(),
    C_any.Id(),
    C_class.Id(),
    C_class.Id(),
    C_string.Id(),
    C_string.Id(),
    C_string.Id(),
    C_integer.Id()),1,MakeFunction7(E_Generate_iteration_statement_For,"Generate_iteration_statement_For")),MakeString("gostat.cl:547"))
  
  _ = Core.F_attach_method(C_Generate_g_statement.AddMethod(Signature(Language.C_Iteration.Id(),
    C_class.Id(),
    C_string.Id(),
    C_boolean.Id(),
    C_any.Id(),
    C_void.Id()),1,MakeFunction5(E_Generate_g_statement_Iteration,"Generate_g_statement_Iteration")),MakeString("gostat.cl:590"))
  
  _ = Core.F_attach_method(C_Generate_g_statement.AddMethod(Signature(Language.C_While.Id(),
    C_class.Id(),
    C_string.Id(),
    C_boolean.Id(),
    C_any.Id(),
    C_void.Id()),1,MakeFunction5(E_Generate_g_statement_While,"Generate_g_statement_While")),MakeString("gostat.cl:622"))
  
  _ = Core.F_attach_method(C_Generate_g_statement.AddMethod(Signature(Language.C_Return.Id(),
    C_class.Id(),
    C_string.Id(),
    C_boolean.Id(),
    C_any.Id(),
    C_void.Id()),1,MakeFunction5(E_Generate_g_statement_Return,"Generate_g_statement_Return")),MakeString("gostat.cl:634"))
  
  _ = Core.F_attach_method(C_Generate_g_statement.AddMethod(Signature(Language.C_Call.Id(),
    C_class.Id(),
    C_string.Id(),
    C_boolean.Id(),
    C_any.Id(),
    C_void.Id()),1,MakeFunction5(E_Generate_g_statement_Call,"Generate_g_statement_Call")),MakeString("gostat.cl:652"))
  
  _ = Core.F_attach_method(C_Generate_inline_stat.AddMethod(Signature(Language.C_Call.Id(),
    C_class.Id(),
    C_string.Id(),
    C_void.Id()),1,MakeFunction3(E_Generate_inline_stat_Call,"Generate_inline_stat_Call")),MakeString("gostat.cl:673"))
  
  _ = Core.F_attach_method(C_Generate_g_statement.AddMethod(Signature(Language.C_Call_method.Id(),
    C_class.Id(),
    C_string.Id(),
    C_boolean.Id(),
    C_any.Id(),
    C_void.Id()),1,MakeFunction5(E_Generate_g_statement_Call_method,"Generate_g_statement_Call_method")),MakeString("gostat.cl:682"))
  
  _ = Core.F_attach_method(C_Generate_g_statement.AddMethod(Signature(Language.C_Call_method1.Id(),
    C_class.Id(),
    C_string.Id(),
    C_boolean.Id(),
    C_any.Id(),
    C_void.Id()),1,MakeFunction5(E_Generate_g_statement_Call_method1,"Generate_g_statement_Call_method1")),MakeString("gostat.cl:690"))
  
  _ = Core.F_attach_method(C_Generate_g_statement.AddMethod(Signature(Language.C_Call_method2.Id(),
    C_class.Id(),
    C_string.Id(),
    C_boolean.Id(),
    C_any.Id(),
    C_void.Id()),1,MakeFunction5(E_Generate_g_statement_Call_method2,"Generate_g_statement_Call_method2")),MakeString("gostat.cl:697"))
  
  _ = Core.F_attach_method(C_Generate_g_statement.AddMethod(Signature(Language.C_Super.Id(),
    C_class.Id(),
    C_string.Id(),
    C_boolean.Id(),
    C_any.Id(),
    C_void.Id()),1,MakeFunction5(E_Generate_g_statement_Super,"Generate_g_statement_Super")),MakeString("gostat.cl:708"))
  
  _ = Core.F_attach_method(C_Generate_g_statement.AddMethod(Signature(Language.C_Cast.Id(),
    C_class.Id(),
    C_string.Id(),
    C_boolean.Id(),
    C_any.Id(),
    C_void.Id()),1,MakeFunction5(E_Generate_g_statement_Cast,"Generate_g_statement_Cast")),MakeString("gostat.cl:713"))
  
  _ = Core.F_attach_method(C_Generate_g_statement.AddMethod(Signature(Language.C_Handle.Id(),
    C_class.Id(),
    C_string.Id(),
    C_boolean.Id(),
    C_any.Id(),
    C_void.Id()),1,MakeFunction5(E_Generate_g_statement_Handle,"Generate_g_statement_Handle")),MakeString("gostat.cl:741"))
  
  _ = Core.F_attach_method(C_Generate_g_statement.AddMethod(Signature(Optimize.C_Compile_C_cast.Id(),
    C_class.Id(),
    C_string.Id(),
    C_boolean.Id(),
    C_any.Id(),
    C_void.Id()),1,MakeFunction5(E_Generate_g_statement_C_cast,"Generate_g_statement_C_cast")),MakeString("gostat.cl:753"))
  
  _ = Core.F_attach_method(C_Generate_g_statement.AddMethod(Signature(Language.C_Call_slot.Id(),
    C_class.Id(),
    C_string.Id(),
    C_boolean.Id(),
    C_any.Id(),
    C_void.Id()),1,MakeFunction5(E_Generate_g_statement_Call_slot,"Generate_g_statement_Call_slot")),MakeString("gostat.cl:767"))
  
  _ = Core.F_attach_method(C_Generate_g_statement.AddMethod(Signature(Language.C_Call_table.Id(),
    C_class.Id(),
    C_string.Id(),
    C_boolean.Id(),
    C_any.Id(),
    C_void.Id()),1,MakeFunction5(E_Generate_g_statement_Call_table,"Generate_g_statement_Call_table")),MakeString("gostat.cl:776"))
  
  _ = Core.F_attach_method(C_Generate_g_statement.AddMethod(Signature(Language.C_Call_array.Id(),
    C_class.Id(),
    C_string.Id(),
    C_boolean.Id(),
    C_any.Id(),
    C_void.Id()),1,MakeFunction5(E_Generate_g_statement_Call_array,"Generate_g_statement_Call_array")),MakeString("gostat.cl:785"))
  
  _ = Core.F_attach_method(C_Generate_g_statement.AddMethod(Signature(Language.C_Update.Id(),
    C_class.Id(),
    C_string.Id(),
    C_boolean.Id(),
    C_any.Id(),
    C_void.Id()),1,MakeFunction5(E_Generate_g_statement_Update,"Generate_g_statement_Update")),MakeString("gostat.cl:834"))
  
  _ = Core.F_attach_method(C_Generate_update_statement.AddMethod(Signature(Language.C_Update.Id(),C_class.Id(),C_void.Id()),1,MakeFunction2(E_Generate_update_statement_Update,"Generate_update_statement_Update")),MakeString("gostat.cl:888"))
  
  _ = Core.F_attach_method(C_Generate_need_shortcut.AddMethod(Signature(C_any.Id(),C_boolean.Id()),0,MakeFunction1(E_Generate_need_shortcut_any,"Generate_need_shortcut_any")),MakeString("gostat.cl:896"))
  
  _ = Core.F_attach_method(C_Generate_external_I.AddMethod(Signature(C_module.Id(),C_string.Id()),0,MakeFunction1(E_Generate_external_I_module,"Generate_external!_module")),MakeString("gomain.cl:31"))
  
  _ = Core.F_attach_method(C_Generate_string2module.AddMethod(Signature(C_string.Id(),C_module.Id()),1,MakeFunction1(E_Generate_string2module_string,"Generate_string2module_string")),MakeString("gomain.cl:36"))
  
  _ = Core.F_attach_method(C_Generate_printHelp.AddMethod(Signature(C_void.Id(),C_void.Id()),0,MakeFunction1(E_Generate_printHelp_void,"Generate_printHelp_void")),MakeString("gomain.cl:65"))
  
  _ = Core.F_attach_method(C_Generate_complex_main.AddMethod(Signature(C_void.Id(),C_void.Id()),1,MakeFunction1(E_Generate_complex_main_void,"Generate_complex_main_void")),MakeString("gomain.cl:158"))
  
  _ = Core.F_attach_method(C_system_file.AddMethod(Signature(C_module.Id(),
    C_string.Id(),
    C_boolean.Id(),
    C_void.Id()),1,MakeFunction3(E_system_file_module,"system_file_module")),MakeString("gomain.cl:188"))
  
  _ = Core.F_attach_method(C_Generate_system_imports.AddMethod(Signature(C_module.Id(),C_void.Id()),1,MakeFunction1(E_Generate_system_imports_module,"Generate_system_imports_module")),MakeString("gomain.cl:197"))
  
  _ = Core.F_attach_method(C_Generate_load_function.AddMethod(Signature(C_module.Id(),C_list.Id(),C_void.Id()),1,MakeFunction2(E_Generate_load_function_module,"Generate_load_function_module")),MakeString("gomain.cl:231"))
  
  _ = Core.F_attach_method(C_Generate_main_function.AddMethod(Signature(C_module.Id(),
    Core.F_nth_class1(C_list,ToType(C_module.Id())).Id(),
    C_boolean.Id(),
    C_void.Id()),0,MakeFunction3(E_Generate_main_function_module,"Generate_main_function_module")),MakeString("gomain.cl:260"))
  
  _ = Core.F_attach_method(C_Generate_compile_dir.AddMethod(Signature(C_module.Id(),C_void.Id()),0,MakeFunction1(E_Generate_compile_dir_module,"Generate_compile_dir_module")),MakeString("gomain.cl:271"))
  
  _ = Core.F_attach_method(C_Generate_compile_exe.AddMethod(Signature(C_string.Id(),C_void.Id()),0,MakeFunction1(E_Generate_compile_exe_string,"Generate_compile_exe_string")),MakeString("gomain.cl:277"))
  
  } 

