/***** CLAIRE Compilation of module Core.cl 
         [version 4.1.4 / safety 5] Wednesday 01-01-2025 17:30:04 *****/

package Core
import (_ "fmt"
	"unsafe"
	. "Kernel"
)

// class file for ephemeral_object in module Core 
type EphemeralObject struct { 
   ClaireObject
   } 

// automatic cast function
func ToEphemeralObject(x *ClaireAny) *EphemeralObject {return (*EphemeralObject)(unsafe.Pointer(x))}

// automatic constructor function
func MakeEphemeralObject() *EphemeralObject { 
  var o *EphemeralObject = new(EphemeralObject)
  o.Isa = C_ephemeral_object
  return o 
  } 

// class file for general_error in module Core 
type GeneralError struct { 
   ClaireError
   Cause *ClaireAny
  Arg *ClaireAny
  } 

// automatic cast function
func ToGeneralError(x *ClaireAny) *GeneralError {return (*GeneralError)(unsafe.Pointer(x))}

// automatic constructor function
func MakeGeneralError(cause *ClaireAny,arg *ClaireAny) *GeneralError { 
  var o *GeneralError = new(GeneralError)
  o.Isa = C_general_error
  o.Cause = cause
  o.Arg = arg
  return o 
  } 

// class file for read_slot_error in module Core 
type ReadSlotError struct { 
   ClaireError
   Arg *ClaireAny
  Wrong *ClaireAny
  } 

// automatic cast function
func ToReadSlotError(x *ClaireAny) *ReadSlotError {return (*ReadSlotError)(unsafe.Pointer(x))}

// automatic constructor function
func MakeReadSlotError(arg *ClaireAny,wrong *ClaireAny) *ReadSlotError { 
  var o *ReadSlotError = new(ReadSlotError)
  o.Isa = C_read_slot_error
  o.Arg = arg
  o.Wrong = wrong
  return o 
  } 

// class file for range_error in module Core 
type RangeError struct { 
   ClaireError
   Cause *ClaireAny
  Arg *ClaireAny
  Wrong *ClaireAny
  } 

// automatic cast function
func ToRangeError(x *ClaireAny) *RangeError {return (*RangeError)(unsafe.Pointer(x))}

// automatic constructor function
func MakeRangeError(cause *ClaireAny,arg *ClaireAny,wrong *ClaireAny) *RangeError { 
  var o *RangeError = new(RangeError)
  o.Isa = C_range_error
  o.Cause = cause
  o.Arg = arg
  o.Wrong = wrong
  return o 
  } 

// class file for selector_error in module Core 
type SelectorError struct { 
   ClaireError
   Selector *ClaireAny
  Arg *ClaireAny
  } 

// automatic cast function
func ToSelectorError(x *ClaireAny) *SelectorError {return (*SelectorError)(unsafe.Pointer(x))}

// automatic constructor function
func MakeSelectorError(selector *ClaireAny,arg *ClaireAny) *SelectorError { 
  var o *SelectorError = new(SelectorError)
  o.Isa = C_selector_error
  o.Selector = selector
  o.Arg = arg
  return o 
  } 

// class file for return_error in module Core 
type ReturnError struct { 
   ClaireError
   Arg *ClaireAny
  } 

// automatic cast function
func ToReturnError(x *ClaireAny) *ReturnError {return (*ReturnError)(unsafe.Pointer(x))}

// automatic constructor function
func MakeReturnError(arg *ClaireAny) *ReturnError { 
  var o *ReturnError = new(ReturnError)
  o.Isa = C_return_error
  o.Arg = arg
  return o 
  } 

// class file for contradiction in module Core 
type Contradiction struct { 
   ClaireException
   } 

// automatic cast function
func ToContradiction(x *ClaireAny) *Contradiction {return (*Contradiction)(unsafe.Pointer(x))}

// automatic constructor function
func MakeContradiction() *Contradiction { 
  var o *Contradiction = new(Contradiction)
  o.Isa = C_contradiction
  return o 
  } 

// class file for global_variable in module Core 
type GlobalVariable struct { 
   ClaireSystemThing
   Value *ClaireAny
  Range *ClaireType
  Store_ask *ClaireBoolean
  } 

// automatic cast function
func ToGlobalVariable(x *ClaireAny) *GlobalVariable {return (*GlobalVariable)(unsafe.Pointer(x))}

// automatic constructor function
func MakeGlobalVariable(name *ClaireSymbol,value *ClaireAny,Range *ClaireType,store_ask *ClaireBoolean) *GlobalVariable { 
  var o *GlobalVariable = new(GlobalVariable)
  o.Isa = C_global_variable
  o.Name = name
  o.Value = value
  o.Range = Range
  o.Store_ask = store_ask
  return o 
  } 

// class file for pretty_printer in module Core 
type PrettyPrinter struct { 
   ClaireThing
   Cpretty *ClairePort
  Cprevious int
  Index int
  Width int
  Pprint *ClaireBoolean
  Pbreak *ClaireBoolean
  Cpstack *ClaireList
  } 

// automatic cast function
func ToPrettyPrinter(x *ClaireAny) *PrettyPrinter {return (*PrettyPrinter)(unsafe.Pointer(x))}

// class file for much_too_far in module Core 
type MuchTooFar struct { 
   ClaireError
   } 

// automatic cast function
func ToMuchTooFar(x *ClaireAny) *MuchTooFar {return (*MuchTooFar)(unsafe.Pointer(x))}

// automatic constructor function
func MakeMuchTooFar() *MuchTooFar { 
  var o *MuchTooFar = new(MuchTooFar)
  o.Isa = C_much_too_far
  return o 
  } 

var C_execute *ClaireProperty /*obj*/
var C_debug *ClaireProperty /*obj*/
var C_eval_message *ClaireProperty /*obj*/
var C_Core_noeval_message *ClaireProperty /*obj*/
var C_eval *ClaireProperty /*obj*/
var C_call *ClaireProperty /*obj*/
var C_self_eval *ClaireProperty /*obj*/
var C_inlineok_ask *ClaireProperty /*obj*/
var C_hold_ask *ClaireProperty /*obj*/
var C_write *ClaireProperty /*obj*/
var C_range_is_wrong *ClaireProperty /*obj*/
var C_Core_update_plus *ClaireProperty /*obj*/
var C_Core_update_dash *ClaireProperty /*obj*/
var C_add_value *ClaireProperty /*obj*/
var C_known_ask *ClaireProperty /*obj*/
var C_unknown_ask *ClaireProperty /*obj*/
var C_erase *ClaireProperty /*obj*/
var C_set_range *ClaireProperty /*obj*/
var C_put_store *ClaireProperty /*obj*/
var C_Core_matching_ask *ClaireProperty /*obj*/
var C_Core_vmatch_ask *ClaireProperty /*obj*/
var C_Core_tmatch_ask *ClaireProperty /*obj*/
var C_Core_find_which *ClaireProperty /*obj*/
var C_main *ClaireProperty /*obj*/
var C_join *ClaireOperation /*obj*/
var C_mClaire_index_I *ClaireProperty /*obj*/
var C_mClaire_base_I *ClaireProperty /*obj*/
var C_ephemeral_object *ClaireClass /*obj*/
var C_spy *ClaireProperty /*obj*/
var C_Core_StopProperty *ClaireTable /*obj*/
var C_reify *ClaireProperty /*obj*/
var C_Core__star_stararg *ClaireProperty /*obj*/
var C_general_error *ClaireClass /*obj*/
var C_read_slot_error *ClaireClass /*obj*/
var C_range_error *ClaireClass /*obj*/
var C_selector_error *ClaireClass /*obj*/
var C_return_error *ClaireClass /*obj*/
var C_contradiction *ClaireClass /*obj*/
var C_global_variable *ClaireClass /*obj*/
var C__inf_equal2 *ClaireOperation /*obj*/
var C_contradiction_occurs *GlobalVariable
var C_nil *GlobalVariable
var C_claire_date *GlobalVariable
var C__I_equal *ClaireOperation /*obj*/
var C__inf_inf *ClaireOperation /*obj*/
var C__sup_sup *ClaireOperation /*obj*/
var C_and *ClaireOperation /*obj*/
var C_or *ClaireOperation /*obj*/
var C_U *ClaireOperation /*obj*/
var C__and *ClaireOperation /*obj*/
var C_meet *ClaireOperation /*obj*/
var C_inherit_ask *ClaireOperation /*obj*/
var C__dash_dash_ask *ClaireOperation /*obj*/
var C__dash_dash_I *ClaireOperation /*obj*/
var C_Core_cpstack *ClaireProperty /*obj*/
var C_pretty_printer *ClaireClass /*obj*/
var C_pretty *PrettyPrinter /*obj*/
var C_apply_self_print *ClaireProperty /*obj*/
var C_short_enough *ClaireProperty /*obj*/
var C_much_too_far *ClaireClass /*obj*/
var C_world_plus *GlobalVariable
var C_world_dash *GlobalVariable
var C_world_dash_I *GlobalVariable
var C__dash_dash *ClaireOperation /*obj*/
var C__equaltype_ask *ClaireOperation /*obj*/
var C_but *ClaireOperation /*obj*/
var C__backslash *ClaireOperation /*obj*/
var C_glb *ClaireOperation /*obj*/
var C_Core__inf_equalt *ClaireOperation /*obj*/
var C_Core__Zt *ClaireOperation /*obj*/
var C_Core_initialize *ClaireProperty // Core/"initialize"
var C_Core_uniform *ClaireProperty // Core/"uniform"
var C_Core_hashinsert *ClaireProperty // Core/"hashinsert"
var C_Core_hashget *ClaireProperty // Core/"hashget"
var C_Core_param_I *ClaireProperty // Core/"param!"
var C_end_of_string *ClaireProperty // claire/"end_of_string"
var C_finite_ask *ClaireProperty // claire/"finite?"
var C_release *ClaireProperty // claire/"release"
var C__at *ClaireOperation // claire/"@"
var C_about *ClaireProperty // claire/"about"
var C_mClaire_get_args *ClaireProperty // mClaire/"get_args"
var C_Core_push_debug *ClaireProperty // Core/"push_debug"
var C_Core_pop_debug *ClaireProperty // Core/"pop_debug"
var C_Core_tr_indent *ClaireProperty // Core/"tr_indent"
var C__plus *ClaireOperation // claire/"+"
var C_Core_identified_ask *ClaireProperty // Core/"identified?"
var C_identical_ask *ClaireProperty // claire/"identical?"
var C_Core_get_index *ClaireProperty // Core/"get_index"
var C_factor_ask *ClaireProperty // claire/"factor?"
var C_divide_ask *ClaireProperty // claire/"divide?"
var C_Id *ClaireProperty // claire/"Id"
var C_Core_check_inverse *ClaireProperty // Core/"check_inverse"
var C_invert *ClaireProperty // claire/"invert"
var C_domain_I *ClaireProperty // claire/"domain!"
var C_methods *ClaireProperty // claire/"methods"
var C_mClaire_cause *ClaireProperty // mClaire/"cause"
var C_Core_wrong *ClaireProperty // Core/"wrong"
var C_format *ClaireProperty // claire/"format"
var C_Core_tformat *ClaireProperty // Core/"tformat"
var C_contradiction_I *ClaireProperty // claire/"contradiction!"
var C_mClaire_get_stack *ClaireProperty // mClaire/"get_stack"
var C_mClaire_put_stack *ClaireProperty // mClaire/"put_stack"
var C_mClaire_push_I *ClaireProperty // mClaire/"push!"
var C_time_get *ClaireProperty // claire/"time_get"
var C_time_set *ClaireProperty // claire/"time_set"
var C_time_show *ClaireProperty // claire/"time_show"
var C_print_in_string *ClaireProperty // claire/"print_in_string"
var C_mClaire_buffer_length *ClaireProperty // mClaire/"buffer_length"
var C_mClaire_buffer_set_length *ClaireProperty // mClaire/"buffer_set_length"
var C_not *ClaireProperty // claire/"not"
var C_externC *ClaireProperty // claire/"externC"
var C_getenv *ClaireProperty // claire/"getenv"
var C_last *ClaireProperty // claire/"last"
var C_rmlast *ClaireProperty // claire/"rmlast"
var C_car *ClaireProperty // claire/"car"
var C_sort *ClaireProperty // claire/"sort"
var C_Core_quicksort *ClaireProperty // Core/"quicksort"
var C_Core_build_powerset *ClaireProperty // Core/"build_powerset"
var C_difference *ClaireProperty // claire/"difference"
var C_Core_of_extract *ClaireProperty // Core/"of_extract"
var C_member *ClaireProperty // claire/"member"
var C_get_value *ClaireProperty // claire/"get_value"
var C_Core_enumerate *ClaireProperty // Core/"enumerate"
var C_Uall *ClaireProperty // claire/"Uall"
var C_unique_ask *ClaireProperty // claire/"unique?"
var C_the *ClaireProperty // claire/"the"
var C_Core_abstract_type *ClaireProperty // Core/"abstract_type"
var C_Core_insert_definition *ClaireProperty // Core/"insert_definition"
var C_mClaire_cpretty *ClaireProperty // mClaire/"cpretty"
var C_mClaire_cprevious *ClaireProperty // mClaire/"cprevious"
var C_mClaire_width *ClaireProperty // mClaire/"width"
var C_mClaire_pprint *ClaireProperty // mClaire/"pprint"
var C_mClaire_pbreak *ClaireProperty // mClaire/"pbreak"
var C_mClaire_set_base *ClaireProperty // mClaire/"set_base"
var C_mClaire_set_index *ClaireProperty // mClaire/"set_index"
var C_mClaire_update *ClaireProperty // mClaire/"update"
var C_make_set *ClaireProperty // claire/"make_set"
var C_get_symbol *ClaireProperty // claire/"get_symbol"
var C_time_read *ClaireProperty // claire/"time_read"
var C_Core_first_arg_type *ClaireProperty // Core/"first_arg_type"
var C_Core_second_arg_type *ClaireProperty // Core/"second_arg_type"
var C_Core_meet_arg_types *ClaireProperty // Core/"meet_arg_types"
var C_make_copy_list *ClaireProperty // claire/"make_copy_list"
var C_make_table *ClaireProperty // claire/"make_table"
var C_Core_first_member_type *ClaireProperty // Core/"first_member_type"
var C_mClaire_printFDigit *ClaireProperty // mClaire/"printFDigit"
var C_abs *ClaireProperty // claire/"abs"
var C_less_ask *ClaireOperation // claire/"less?"
var C_unsafe *ClaireProperty // claire/"unsafe"
var C_attach *ClaireProperty // claire/"attach"
var C_Core_nth_arg_type *ClaireProperty // Core/"nth_arg_type"
var C_instanced *ClaireProperty // claire/"instanced"
var C_typed_copy_list *ClaireProperty // claire/"typed_copy_list"
var C_nth_write *ClaireProperty // claire/"nth_write"
var C_Core_belong *ClaireProperty // Core/"belong"
var C_Core_super_apply *ClaireProperty // Core/"super_apply"
var C_Core_add_value_I *ClaireProperty // Core/"add_value!"
var C_mClaire_nth_object *ClaireProperty // mClaire/"nth_object"
var C_Core_thing_type_class *ClaireProperty // Core/"thing_type_class"
var C_Core_object_type_class *ClaireProperty // Core/"object_type_class"
var C_Core_new_defaults *ClaireProperty // Core/"new_defaults"
var C_Reference_I *ClaireProperty // claire/"Reference!"
var C_read_lambda *ClaireProperty // claire/"read_lambda"
var C_Core_db_bind *ClaireProperty // Core/"db_bind"
var C_Core_db_unbind *ClaireProperty // Core/"db_unbind"
var C_owner *ClaireProperty // claire/"owner"
var C_check_in *ClaireProperty // claire/"check_in"
var It *ClaireModule

// definition of the meta-model for module Core 
func MetaLoad() { 
  
  It = MakeModule("Core",C_mClaire)
  It.Comment = MakeString("Compiled on Wednesday 01-01-2025 17:30:04(v4.1.4), lines:2428, warnings:7,safety:5")
  ClEnv.Module_I = It
  
  // definition of the properties
  C_Core_initialize = MakeProperty("initialize",1,It)
  C_Core_uniform = MakeProperty("uniform",1,It)
  C_Core_hashinsert = MakeProperty("hashinsert",1,It)
  C_Core_hashget = MakeProperty("hashget",1,It)
  C_Core_param_I = MakeProperty("param!",1,It)
  C_end_of_string = MakeProperty("end_of_string",1,C_claire)
  C_finite_ask = MakeProperty("finite?",1,C_claire)
  C_release = MakeProperty("release",1,C_claire)
  C__at = MakeOperation("@",1,C_claire,10)
  C_about = MakeProperty("about",1,C_claire)
  C_mClaire_get_args = MakeProperty("get_args",1,C_mClaire)
  C_Core_push_debug = MakeProperty("push_debug",1,It)
  C_Core_pop_debug = MakeProperty("pop_debug",1,It)
  C_Core_tr_indent = MakeProperty("tr_indent",1,It)
  C__plus = MakeOperation("+",1,C_claire,20)
  C_Core_identified_ask = MakeProperty("identified?",1,It)
  C_identical_ask = MakeProperty("identical?",1,C_claire)
  C_Core_get_index = MakeProperty("get_index",1,It)
  C_factor_ask = MakeProperty("factor?",1,C_claire)
  C_divide_ask = MakeProperty("divide?",1,C_claire)
  C_Id = MakeProperty("Id",1,C_claire)
  C_Core_check_inverse = MakeProperty("check_inverse",1,It)
  C_invert = MakeProperty("invert",1,C_claire)
  C_domain_I = MakeProperty("domain!",1,C_claire)
  C_methods = MakeProperty("methods",1,C_claire)
  C_mClaire_cause = MakeProperty("cause",1,C_mClaire)
  C_Core_wrong = MakeProperty("wrong",1,It)
  C_format = MakeProperty("format",1,C_claire)
  C_Core_tformat = MakeProperty("tformat",1,It)
  C_contradiction_I = MakeProperty("contradiction!",1,C_claire)
  C_mClaire_get_stack = MakeProperty("get_stack",1,C_mClaire)
  C_mClaire_put_stack = MakeProperty("put_stack",1,C_mClaire)
  C_mClaire_push_I = MakeProperty("push!",1,C_mClaire)
  C_time_get = MakeProperty("time_get",1,C_claire)
  C_time_set = MakeProperty("time_set",1,C_claire)
  C_time_show = MakeProperty("time_show",2,C_claire)
  C_print_in_string = MakeProperty("print_in_string",1,C_claire)
  C_mClaire_buffer_length = MakeProperty("buffer_length",1,C_mClaire)
  C_mClaire_buffer_set_length = MakeProperty("buffer_set_length",1,C_mClaire)
  C_not = MakeProperty("not",1,C_claire)
  C_externC = MakeProperty("externC",1,C_claire)
  C_getenv = MakeProperty("getenv",2,C_claire)
  C_last = MakeProperty("last",1,C_claire)
  C_rmlast = MakeProperty("rmlast",1,C_claire)
  C_car = MakeProperty("car",1,C_claire)
  C_sort = MakeProperty("sort",1,C_claire)
  C_Core_quicksort = MakeProperty("quicksort",1,It)
  C_Core_build_powerset = MakeProperty("build_powerset",1,It)
  C_difference = MakeProperty("difference",1,C_claire)
  C_Core_of_extract = MakeProperty("of_extract",1,It)
  C_member = MakeProperty("member",1,C_claire)
  C_get_value = MakeProperty("get_value",1,C_claire)
  C_Core_enumerate = MakeProperty("enumerate",1,It)
  C_Uall = MakeProperty("Uall",1,C_claire)
  C_unique_ask = MakeProperty("unique?",1,C_claire)
  C_the = MakeProperty("the",1,C_claire)
  C_Core_abstract_type = MakeProperty("abstract_type",1,It)
  C_Core_insert_definition = MakeProperty("insert_definition",1,It)
  C_mClaire_cpretty = MakeProperty("cpretty",1,C_mClaire)
  C_mClaire_cprevious = MakeProperty("cprevious",1,C_mClaire)
  C_mClaire_width = MakeProperty("width",1,C_mClaire)
  C_mClaire_pprint = MakeProperty("pprint",1,C_mClaire)
  C_mClaire_pbreak = MakeProperty("pbreak",1,C_mClaire)
  C_mClaire_set_base = MakeProperty("set_base",1,C_mClaire)
  C_mClaire_set_index = MakeProperty("set_index",1,C_mClaire)
  C_mClaire_update = MakeProperty("update",1,C_mClaire)
  C_make_set = MakeProperty("make_set",1,C_claire)
  C_get_symbol = MakeProperty("get_symbol",1,C_claire)
  C_time_read = MakeProperty("time_read",2,C_claire)
  C_Core_first_arg_type = MakeProperty("first_arg_type",1,It)
  C_Core_second_arg_type = MakeProperty("second_arg_type",1,It)
  C_Core_meet_arg_types = MakeProperty("meet_arg_types",1,It)
  C_make_copy_list = MakeProperty("make_copy_list",1,C_claire)
  C_make_table = MakeProperty("make_table",1,C_claire)
  C_Core_first_member_type = MakeProperty("first_member_type",1,It)
  C_mClaire_printFDigit = MakeProperty("printFDigit",1,C_mClaire)
  C_abs = MakeProperty("abs",1,C_claire)
  C_less_ask = MakeOperation("less?",1,C_claire,60)
  C_unsafe = MakeProperty("unsafe",1,C_claire)
  C_attach = MakeProperty("attach",1,C_claire)
  C_Core_nth_arg_type = MakeProperty("nth_arg_type",1,It)
  C_instanced = MakeProperty("instanced",1,C_claire)
  C_typed_copy_list = MakeProperty("typed_copy_list",1,C_claire)
  C_nth_write = MakeProperty("nth_write",1,C_claire)
  C_Core_belong = MakeProperty("belong",1,It)
  C_Core_super_apply = MakeProperty("super_apply",1,It)
  C_Core_add_value_I = MakeProperty("add_value!",1,It)
  C_mClaire_nth_object = MakeProperty("nth_object",1,C_mClaire)
  C_Core_thing_type_class = MakeProperty("thing_type_class",1,It)
  C_Core_object_type_class = MakeProperty("object_type_class",1,It)
  C_Core_new_defaults = MakeProperty("new_defaults",1,It)
  C_Reference_I = MakeProperty("Reference!",1,C_claire)
  C_read_lambda = MakeProperty("read_lambda",1,C_claire)
  C_Core_db_bind = MakeProperty("db_bind",1,It)
  C_Core_db_unbind = MakeProperty("db_unbind",1,It)
  C_owner = MakeProperty("owner",1,C_claire)
  C_check_in = MakeProperty("check_in",2,C_claire)
  
  // instructions from module sources
  { 
    var s *ClaireSlot
    _ = s
    var s_iter *ClaireAny
    var s_support *ClaireList
    s_support = C_slot.Instances
    s_len := s_support.Length()
    for i_it := 0; i_it < s_len; i_it++ { 
      s_iter = s_support.At(i_it)
      s = ToSlot(s_iter)
      F_close_slot(s)
      } 
    } 
  { 
    var m *ClaireMethod
    _ = m
    var m_iter *ClaireAny
    var m_support *ClaireList
    m_support = C_method.Instances
    m_len := m_support.Length()
    for i_it := 0; i_it < m_len; i_it++ { 
      m_iter = m_support.At(i_it)
      m = ToMethod(m_iter)
      F_close_method(m)
      } 
    } 
  
  _ = F_attach_method(C_close.AddMethod(Signature(C_slot.Id(),C_slot.Id()),0,MakeFunction1(E_close_slot,"close_slot")),MakeString("method.cl:31"))
  
  _ = F_attach_method(C_close.AddMethod(Signature(C_method.Id(),C_method.Id()),0,MakeFunction1(E_close_method,"close_method")),MakeString("method.cl:35"))
  
  _ = F_attach_method(C_attach.AddMethod(Signature(C_method.Id(),C_string.Id(),C_method.Id()),0,MakeFunction2(E_attach_method,"attach_method")),MakeString("method.cl:40"))
  
  C_execute = MakeProperty("execute",1,C_claire)
  
  
  C_debug = MakeProperty("debug",2,C_claire)
  
  
  C_eval_message = MakeProperty("eval_message",1,C_claire)
  
  
  C_Core_noeval_message = MakeProperty("noeval_message",1,It)
  
  
  C_eval = MakeProperty("eval",1,C_claire)
  
  
  C_call = MakeProperty("call",1,C_claire)
  
  
  C_self_eval = MakeProperty("self_eval",2,C_claire)
  
  
  C_inlineok_ask = MakeProperty("inlineok?",2,C_claire)
  
  
  C_hold_ask = MakeProperty("hold?",2,C_claire)
  
  
  C_write = MakeProperty("write",1,C_claire)
  
  
  C_range_is_wrong = MakeProperty("range_is_wrong",1,C_claire)
  
  
  C_Core_update_plus = MakeProperty("update+",1,It)
  
  
  C_Core_update_dash = MakeProperty("update-",1,It)
  
  
  C_add_value = MakeProperty("add_value",1,C_claire)
  
  
  C_known_ask = MakeProperty("known?",1,C_claire)
  
  
  C_unknown_ask = MakeProperty("unknown?",1,C_claire)
  
  
  C_erase = MakeProperty("erase",2,C_claire)
  
  
  C_set_range = MakeProperty("set_range",2,C_claire)
  
  
  C_put_store = MakeProperty("put_store",2,C_claire)
  
  
  C_Core_matching_ask = MakeProperty("matching?",1,It)
  
  
  C_Core_vmatch_ask = MakeProperty("vmatch?",1,It)
  
  
  C_Core_tmatch_ask = MakeProperty("tmatch?",1,It)
  
  
  C_Core_find_which = MakeProperty("find_which",1,It)
  
  
  C_main = MakeProperty("main",2,C_claire)
  
  
  _ = F_attach_method(C_eval_message.AddMethod(Signature(C_property.Id(),
    C_object.Id(),
    C_integer.Id(),
    C_boolean.Id(),
    C_any.Id()),1,MakeFunction4(E_eval_message_property,"eval_message_property")),MakeString("method.cl:101"))
  
  _ = F_attach_method(C_Core_noeval_message.AddMethod(Signature(C_property.Id(),C_integer.Id(),C_any.Id()),1,MakeFunction2(E_noeval_message_property2,"noeval_message_property2")),MakeString("method.cl:108"))
  
  _ = F_attach_method(C_execute.AddMethod(Signature(C_method.Id(),
    C_integer.Id(),
    C_boolean.Id(),
    C_any.Id()),1,MakeFunction3(E_execute_method,"execute_method")),MakeString("method.cl:137"))
  
  _ = F_attach_method(C_eval.AddMethod(Signature(C_any.Id(),C_any.Id()),1,MakeFunction1(E_eval_any,"eval_any")),MakeString("method.cl:140"))
  
  _ = F_attach_method(C_inlineok_ask.AddMethod(Signature(C_method.Id(),C_string.Id(),C_method.Id()),0,MakeFunction2(E_inlineok_ask_method,"inlineok_ask_method")),MakeString("method.cl:151"))
  
  _ = F_attach_method(C_read_lambda.AddMethod(Signature(C_string.Id(),C_lambda.Id()),1,MakeFunction1(E_read_lambda_string,"read_lambda_string")),MakeString("method.cl:157"))
  
  _ = F_attach_method(C_get.AddMethod(Signature(C_slot.Id(),C_object.Id(),C_any.Id()),0,MakeFunction2(E_get_slot,"get_slot")),MakeString("method.cl:164"))
  
  _ = F_attach_method(C_put.AddMethod(Signature(C_slot.Id(),
    C_object.Id(),
    C_any.Id(),
    C_any.Id()),0,MakeFunction3(E_put_slot,"put_slot")),MakeString("method.cl:166"))
  
  _ = F_attach_method(C_get.AddMethod(Signature(C_property.Id(),C_object.Id(),C_any.Id()),0,MakeFunction2(E_get_property,"get_property")),MakeString("method.cl:172"))
  
  _ = F_attach_method(C_funcall.AddMethod(Signature(C_property.Id(),C_any.Id(),C_any.Id()),1,MakeFunction2(E_funcall_property,"funcall_property")),MakeString("method.cl:179"))
  
  _ = F_attach_method(C_hold_ask.AddMethod(Signature(C_property.Id(),
    C_object.Id(),
    C_any.Id(),
    C_boolean.Id()),0,MakeFunction3(E_hold_ask_property,"hold_ask_property")),MakeString("method.cl:187"))
  
  _ = F_attach_method(C_write.AddMethod(Signature(C_property.Id(),
    C_object.Id(),
    C_any.Id(),
    C_void.Id()),1,MakeFunction3(E_write_property,"write_property")),MakeString("method.cl:205"))
  
  _ = F_attach_method(C_range_is_wrong.AddMethod(Signature(C_slot.Id(),C_any.Id(),C_void.Id()),1,MakeFunction2(E_range_is_wrong_slot,"range_is_wrong_slot")),MakeString("method.cl:209"))
  
  _ = F_attach_method(C_put.AddMethod(Signature(C_property.Id(),
    C_object.Id(),
    C_integer.Id(),
    C_class.Id(),
    C_any.Id(),
    C_void.Id()),1,MakeFunction5(E_put_property1,"put_property1")),MakeString("method.cl:214"))
  
  _ = F_attach_method(C_mClaire_update.AddMethod(Signature(C_property.Id(),
    C_object.Id(),
    C_integer.Id(),
    C_class.Id(),
    C_any.Id(),
    C_void.Id()),1,MakeFunction5(E_update_property,"update_property")),MakeString("method.cl:235"))
  
  _ = F_attach_method(C_Core_update_plus.AddMethod(Signature(C_relation.Id(),
    C_any.Id(),
    C_any.Id(),
    C_void.Id()),1,MakeFunction3(E_update_plus_relation,"update_plus_relation")),MakeString("method.cl:253"))
  
  _ = F_attach_method(C_Core_update_dash.AddMethod(Signature(C_relation.Id(),
    C_any.Id(),
    C_any.Id(),
    C_void.Id()),0,MakeFunction3(E_update_dash_relation,"update_dash_relation")),MakeString("method.cl:268"))
  
  _ = F_attach_method(C_add_I.AddMethod(Signature(C_property.Id(),
    C_object.Id(),
    C_integer.Id(),
    C_any.Id(),
    C_void.Id()),1,MakeFunction4(E_add_I_property,"add_I_property")),MakeString("method.cl:275"))
  
  _ = F_attach_method(C_Core_add_value_I.AddMethod(Signature(C_property.Id(),
    C_object.Id(),
    C_integer.Id(),
    C_set.Id(),
    C_any.Id(),
    C_boolean.Id()),0,MakeFunction5(E_Core_add_value_I_property,"Core_add_value_I_property")),MakeString("method.cl:284"))
  
  _ = F_attach_method(C_add.AddMethod(Signature(C_property.Id(),
    C_object.Id(),
    C_any.Id(),
    C_void.Id()),1,MakeFunction3(E_add_property,"add_property")),MakeString("method.cl:296"))
  
  _ = F_attach_method(C_known_ask.AddMethod(Signature(C_property.Id(),C_object.Id(),C_boolean.Id()),0,MakeFunction2(E_known_ask_property,"known_ask_property")),MakeString("method.cl:301"))
  
  _ = F_attach_method(C_unknown_ask.AddMethod(Signature(C_property.Id(),C_object.Id(),C_boolean.Id()),0,MakeFunction2(E_unknown_ask_property,"unknown_ask_property")),MakeString("method.cl:304"))
  
  _ = F_attach_method(C_delete.AddMethod(Signature(C_property.Id(),
    C_object.Id(),
    C_any.Id(),
    C_any.Id()),1,MakeFunction3(E_delete_property,"delete_property")),MakeString("method.cl:316"))
  
  _ = F_attach_method(C_erase.AddMethod(Signature(C_property.Id(),C_object.Id(),C_any.Id()),1,MakeFunction2(E_erase_property,"erase_property")),MakeString("method.cl:331"))
  
  _ = F_attach_method(C_set_range.AddMethod(Signature(C_property.Id(),
    C_class.Id(),
    C_type.Id(),
    C_void.Id()),0,MakeFunction3(E_set_range_property,"set_range_property")),MakeString("method.cl:338"))
  
  _ = F_attach_method(C_put_store.AddMethod(Signature(C_property.Id(),
    C_object.Id(),
    C_any.Id(),
    C_boolean.Id(),
    C_void.Id()),1,MakeFunction4(E_put_store_property2,"put_store_property2")),MakeString("method.cl:352"))
  
  _ = F_attach_method(C_fastcall.AddMethod(Signature(C_relation.Id(),
    C_any.Id(),
    C_any.Id(),
    C_void.Id()),1,MakeFunction3(E_fastcall_relation2,"fastcall_relation2")),MakeString("method.cl:368"))
  
  C_join = ToOperation(new(ClaireOperation).IsNamed(C_operation,MakeSymbol("join",C_claire)))
  
  
  _ = F_attach_method(C_Core_insert_definition.AddMethod(Signature(C_property.Id(),C_restriction.Id(),C_void.Id()),0,MakeFunction2(E_insert_definition_property,"insert_definition_property")),MakeString("method.cl:380"))
  
  _ = F_attach_method(C_Core_initialize.AddMethod(Signature(C_restriction.Id(),
    C_class.Id(),
    C_list.Id(),
    C_list.Id()),0,MakeFunction3(E_initialize_restriction1,"initialize_restriction1")),MakeString("method.cl:394"))
  
  _ = F_attach_method(C_Core_uniform.AddMethod(Signature(C_restriction.Id(),C_boolean.Id()),0,MakeFunction1(E_uniform_restriction,"uniform_restriction")),MakeString("method.cl:406"))
  
  _ = F_attach_method(C_Core_uniform.AddMethod(Signature(C_property.Id(),C_boolean.Id()),0,MakeFunction1(E_uniform_property,"uniform_property")),MakeString("method.cl:410"))
  
  _ = F_attach_method(C_Core_initialize.AddMethod(Signature(C_restriction.Id(),C_list.Id(),C_list.Id()),0,MakeFunction2(E_initialize_restriction2,"initialize_restriction2")),MakeString("method.cl:424"))
  
  _ = F_attach_method(C_Core_hashinsert.AddMethod(Signature(C_restriction.Id(),C_any.Id()),0,MakeFunction1(E_hashinsert_restriction,"hashinsert_restriction")),MakeString("method.cl:430"))
  
  _ = F_attach_method(C_Core_hashinsert.AddMethod(Signature(C_class.Id(),C_method.Id(),C_any.Id()),0,MakeFunction2(E_hashinsert_class,"hashinsert_class")),MakeString("method.cl:439"))
  
  _ = F_attach_method(C_Core_hashget.AddMethod(Signature(C_class.Id(),C_property.Id(),C_object.Id()),0,MakeFunction2(E_hashget_class,"hashget_class")),MakeString("method.cl:443"))
  
  _ = F_attach_method(C_join.AddMethod(Signature(C_list.Id(),C_list.Id(),C_boolean.Id()),0,MakeFunction2(E_join_list,"join_list")),MakeString("method.cl:451"))
  
  _ = F_attach_method(C_stack_apply.AddMethod(Signature(C_property.Id(),C_integer.Id(),C_any.Id()),1,MakeFunction2(E_CALL,"CALL")),MakeString("method.cl:464"))
  
  _ = F_attach_method(C_Core_super_apply.AddMethod(Signature(C_property.Id(),
    C_class.Id(),
    C_integer.Id(),
    C_any.Id()),1,MakeFunction3(E_SUPER,"SUPER")),MakeString("method.cl:470"))
  
  _ = F_attach_method(C__at.AddMethod(Signature(C_property.Id(),C_class.Id(),C_object.Id()),0,MakeFunction2(E__at_property1,"_at_property1")),MakeString("method.cl:480"))
  
  _ = F_attach_method(C__at.AddMethod(Signature(C_property.Id(),C_list.Id(),C_object.Id()),0,MakeFunction2(E__at_property2,"_at_property2")),MakeString("method.cl:486"))
  
  _ = F_attach_method(C_Core_matching_ask.AddMethod(Signature(C_list.Id(),
    C_integer.Id(),
    C_integer.Id(),
    C_boolean.Id()),0,MakeFunction3(E_matching_ask_list,"matching_ask_list")),MakeString("method.cl:508"))
  
  _ = F_attach_method(C_Core_vmatch_ask.AddMethod(Signature(C_any.Id(),
    C_any.Id(),
    C_integer.Id(),
    C_boolean.Id()),0,MakeFunction3(E_vmatch_ask_any,"vmatch_ask_any")),MakeString("method.cl:531"))
  
  _ = F_attach_method(C_Core_tmatch_ask.AddMethod(Signature(C_list.Id(),C_list.Id(),C_boolean.Id()),0,MakeFunction2(E_tmatch_ask_list,"tmatch_ask_list")),MakeString("method.cl:540"))
  
  _ = F_attach_method(C_Core_tmatch_ask.AddMethod(Signature(C_any.Id(),
    C_any.Id(),
    C_list.Id(),
    C_boolean.Id()),0,MakeFunction3(E_tmatch_ask_any,"tmatch_ask_any")),MakeString("method.cl:554"))
  
  _ = F_attach_method(C_Core_find_which.AddMethod(Signature(C_property.Id(),
    C_integer.Id(),
    C_class.Id(),
    C_object.Id()),0,MakeFunction3(E_find_which_property,"find_which_property")),MakeString("method.cl:560"))
  
  _ = F_attach_method(C_Core_find_which.AddMethod(Signature(C_list.Id(),
    C_class.Id(),
    C_integer.Id(),
    C_integer.Id(),
    C_object.Id()),0,MakeFunction4(E_find_which_list,"find_which_list")),MakeString("method.cl:564"))
  
  _ = F_attach_method(C_Core_find_which.AddMethod(Signature(C_class.Id(),
    C_list.Id(),
    C_integer.Id(),
    C_integer.Id(),
    C_object.Id()),0,MakeFunction4(E_find_which_class,"find_which_class")),MakeString("method.cl:569"))
  
  ClEnv.Version = 1.4
  PRINC("-- CLAIRE run-time library v 4.")
  F_princ_float(1.4)
  PRINC(" [os: ")
  PRINC("macos")
  PRINC(", compiler:")
  PRINC("go")
  PRINC(" ] --\n")
  
  _ = F_attach_method(C_release.AddMethod(Signature(C_void.Id(),C_any.Id()),1,MakeFunction1(E_release_void,"release_void")),MakeString("object.cl:29"))
  
  _ = F_attach_method(C_about.AddMethod(Signature(C_void.Id(),C_any.Id()),0,MakeFunction1(E_about_void,"about_void")),MakeString("object.cl:41"))
  
  C_mClaire_index_I = MakeProperty("index!",1,C_mClaire)
  C_mClaire_index_I.Range = ToType(C_integer.Id())
  
  
  C_mClaire_base_I = MakeProperty("base!",1,C_mClaire)
  C_mClaire_base_I.Range = ToType(C_integer.Id())
  
  
  C_ephemeral_object = NewClass("ephemeral_object",C_object,C_claire)
  
  _ = F_attach_method(C_mClaire_get_args.AddMethod(Signature(C_integer.Id(),C_list.Id()),0,MakeFunction1(E_get_args_integer,"get_args_integer")),MakeString("object.cl:57"))
  
  _ = F_attach_method(C_funcall.AddMethod(Signature(C_method.Id(),C_any.Id(),C_any.Id()),1,MakeFunction2(E_funcall_method1,"funcall_method1")),MakeString("object.cl:62"))
  
  _ = F_attach_method(C_funcall.AddMethod(Signature(C_method.Id(),
    C_any.Id(),
    C_any.Id(),
    C_any.Id()),1,MakeFunction3(E_funcall_method2,"funcall_method2")),MakeString("object.cl:67"))
  
  _ = F_attach_method(C_funcall.AddMethod(Signature(C_method.Id(),
    C_any.Id(),
    C_any.Id(),
    C_any.Id(),
    C_any.Id()),1,MakeFunction4(E_funcall_method3,"funcall_method3")),MakeString("object.cl:75"))
  
  _ = F_attach_method(C_call.AddMethod(Signature(C_property.Id(),C_listargs.Id(),C_any.Id()),1,MakeFunction2(E_call_property,"call_property")),MakeString("object.cl:78"))
  
  _ = F_attach_method(C_apply.AddMethod(Signature(C_property.Id(),C_list.Id(),C_any.Id()),1,MakeFunction2(E_apply_property,"apply_property")),MakeString("object.cl:84"))
  
  _ = F_attach_method(C_apply.AddMethod(Signature(C_method.Id(),C_list.Id(),C_any.Id()),1,MakeFunction2(E_apply_method,"apply_method")),MakeString("object.cl:88"))
  
  C_spy = MakeProperty("spy",3,C_claire)
  
  
  _ = F_attach_method(C_Core_push_debug.AddMethod(Signature(C_property.Id(),
    C_integer.Id(),
    C_integer.Id(),
    C_void.Id()),1,MakeFunction3(E_push_debug_property,"push_debug_property")),MakeString("object.cl:122"))
  
  _ = F_attach_method(C_Core_pop_debug.AddMethod(Signature(C_property.Id(),
    C_integer.Id(),
    C_any.Id(),
    C_void.Id()),1,MakeFunction3(E_pop_debug_property,"pop_debug_property")),MakeString("object.cl:144"))
  
  _ = F_attach_method(C_Core_tr_indent.AddMethod(Signature(C_boolean.Id(),C_integer.Id(),C_void.Id()),0,MakeFunction2(E_tr_indent_boolean,"tr_indent_boolean")),MakeString("object.cl:151"))
  
  _ = F_attach_method(C_Core_db_bind.AddMethod(Signature(C_module.Id(),
    C_property.Id(),
    C_integer.Id(),
    C_void.Id()),1,MakeFunction3(E_Core_db_bind_module,"Core_db_bind_module")),MakeString("object.cl:156"))
  
  _ = F_attach_method(C_Core_db_unbind.AddMethod(Signature(C_module.Id(),
    C_property.Id(),
    C_any.Id(),
    C_void.Id()),1,MakeFunction3(E_Core_db_unbind_module,"Core_db_unbind_module")),MakeString("object.cl:160"))
  
  _ = F_attach_method(C_Core_identified_ask.AddMethod(Signature(C_class.Id(),C_boolean.Id()),0,MakeFunction1(E_identified_ask_class,"identified_ask_class")),MakeString("object.cl:170"))
  
  _ = F_attach_method(C_identical_ask.AddMethod(Signature(C_any.Id(),C_any.Id(),C_boolean.Id()),0,MakeFunction2(E_identical_ask_any,"identical_ask_any")),MakeString("object.cl:174"))
  
  _ = F_attach_method(C_put.AddMethod(Signature(C_property.Id(),
    C_object.Id(),
    C_any.Id(),
    C_any.Id()),1,MakeFunction3(E_put_property2,"put_property2")),MakeString("object.cl:184"))
  
  _ = F_attach_method(C_add_value.AddMethod(Signature(C_property.Id(),
    C_object.Id(),
    C_any.Id(),
    C_void.Id()),1,MakeFunction3(E_add_value_property3,"add_value_property3")),MakeString("object.cl:193"))
  
  F_attach_method(C_nth.AddMethod(Signature(C_table.Id(),C_any.Id(),C_any.Id()),1,MakeFunction2(E_nth_table1,"nth_table1")),MakeString("object.cl:209")).Typing = MakeFunction2(E_nth_table1_type,"nth_table1_type").Id()
  
  F_attach_method(C_get.AddMethod(Signature(C_table.Id(),C_any.Id(),C_any.Id()),0,MakeFunction2(E_get_table,"get_table")),MakeString("object.cl:218")).Typing = MakeFunction2(E_get_table_type,"get_table_type").Id()
  
  _ = F_attach_method(C_nth_equal.AddMethod(Signature(C_table.Id(),
    C_any.Id(),
    C_any.Id(),
    C_void.Id()),1,MakeFunction3(E_nth_equal_table1,"nth_equal_table1")),MakeString("object.cl:225"))
  
  _ = F_attach_method(C_nth_put.AddMethod(Signature(C_table.Id(),
    C_any.Id(),
    C_any.Id(),
    C_void.Id()),1,MakeFunction3(E_nth_put_table,"nth_put_table")),MakeString("object.cl:244"))
  
  _ = F_attach_method(C_put.AddMethod(Signature(C_table.Id(),
    C_any.Id(),
    C_any.Id(),
    C_void.Id()),0,MakeFunction3(E_put_table,"put_table")),MakeString("object.cl:255"))
  
  _ = F_attach_method(C_add.AddMethod(Signature(C_table.Id(),
    C_any.Id(),
    C_any.Id(),
    C_void.Id()),1,MakeFunction3(E_add_table,"add_table")),MakeString("object.cl:262"))
  
  _ = F_attach_method(C_add_I.AddMethod(Signature(C_table.Id(),
    C_any.Id(),
    C_any.Id(),
    C_void.Id()),1,MakeFunction3(E_add_I_table,"add_I_table")),MakeString("object.cl:267"))
  
  _ = F_attach_method(C_Core_add_value_I.AddMethod(Signature(C_table.Id(),
    C_any.Id(),
    C_set.Id(),
    C_any.Id(),
    C_boolean.Id()),0,MakeFunction4(E_Core_add_value_I_table,"Core_add_value_I_table")),MakeString("object.cl:275"))
  
  _ = F_attach_method(C_add_value.AddMethod(Signature(C_table.Id(),
    C_any.Id(),
    C_any.Id(),
    C_void.Id()),0,MakeFunction3(E_add_value_table3,"add_value_table3")),MakeString("object.cl:280"))
  
  _ = F_attach_method(C_delete.AddMethod(Signature(C_table.Id(),
    C_any.Id(),
    C_any.Id(),
    C_any.Id()),0,MakeFunction3(E_delete_table,"delete_table")),MakeString("object.cl:290"))
  
  F_attach_method(C_nth.AddMethod(Signature(C_table.Id(),
    C_any.Id(),
    C_any.Id(),
    C_any.Id()),1,MakeFunction3(E_nth_table2,"nth_table2")),MakeString("object.cl:300")).Typing = MakeFunction3(E_nth_table2_type,"nth_table2_type").Id()
  
  _ = F_attach_method(C_nth_equal.AddMethod(Signature(C_table.Id(),
    C_any.Id(),
    C_any.Id(),
    C_any.Id(),
    C_void.Id()),1,MakeFunction4(E_nth_equal_table2,"nth_equal_table2")),MakeString("object.cl:314"))
  
  _ = F_attach_method(C_Core_get_index.AddMethod(Signature(C_table.Id(),C_any.Id(),C_integer.Id()),0,MakeFunction2(E_get_index_table1,"get_index_table1")),MakeString("object.cl:321"))
  
  _ = F_attach_method(C_Core_get_index.AddMethod(Signature(C_table.Id(),
    C_integer.Id(),
    C_integer.Id(),
    C_integer.Id()),0,MakeFunction3(E_get_index_table2,"get_index_table2")),MakeString("object.cl:324"))
  
  _ = F_attach_method(C_erase.AddMethod(Signature(C_table.Id(),C_void.Id()),1,MakeFunction1(E_erase_table,"erase_table")),MakeString("object.cl:333"))
  
  _ = F_attach_method(C_make_table.AddMethod(Signature(C_type.Id(),
    C_type.Id(),
    C_any.Id(),
    C_table.Id()),1,MakeFunction3(E_make_table_type,"make_table_type")),MakeString("object.cl:345"))
  
  C_Core_StopProperty = ToTable(new(ClaireTable).IsNamed(C_table,MakeSymbol("StopProperty",It)))
  C_Core_StopProperty.Range = ToType(C_list.Id())
  C_Core_StopProperty.Params = C_any.Id()
  C_Core_StopProperty.Domain = ToType(C_property.Id())
  C_Core_StopProperty.GraphInit()
  
  C_Core_StopProperty.Default = CNULL
  
  _ = F_attach_method(C_funcall.AddMethod(Signature(C_lambda.Id(),C_any.Id(),C_any.Id()),1,MakeFunction2(E_funcall_lambda1,"funcall_lambda1")),MakeString("object.cl:363"))
  
  _ = F_attach_method(C_funcall.AddMethod(Signature(C_lambda.Id(),
    C_any.Id(),
    C_any.Id(),
    C_any.Id()),1,MakeFunction3(E_funcall_lambda2,"funcall_lambda2")),MakeString("object.cl:376"))
  
  _ = F_attach_method(C_funcall.AddMethod(Signature(C_lambda.Id(),
    C_any.Id(),
    C_any.Id(),
    C_any.Id(),
    C_any.Id()),1,MakeFunction4(E_funcall_lambda3,"funcall_lambda3")),MakeString("object.cl:390"))
  
  _ = F_attach_method(C_Core_check_inverse.AddMethod(Signature(C_any.Id(),C_any.Id(),C_void.Id()),1,MakeFunction2(E_check_inverse_any,"check_inverse_any")),MakeString("object.cl:402"))
  
  { 
    var expr EID
    C_relation.Open = 0
    expr = F_update_property(C_inverse,
      ToObject(C_inverse.Id()),
      8,
      C_relation,
      C_inverse.Id())
    if !ErrorIn(expr) {
    { 
      var va_arg1 *ClaireRelation
      var va_arg2 *ClaireAny
      va_arg1 = ToRelation(C_inverse.Id())
      va_arg2 = ToMethod(F__at_property1(C_Core_check_inverse,C_any).Id()).Functional.Id()
      va_arg1.IfWrite = va_arg2
      expr = va_arg2.ToEID()
      } 
    }
    ErrorCheck(expr)} 
  
  _ = F_attach_method(C_invert.AddMethod(Signature(C_relation.Id(),C_any.Id(),C_set.Id()),1,MakeFunction2(E_invert_relation,"invert_relation")),MakeString("object.cl:420"))
  
  _ = F_attach_method(C_domain_I.AddMethod(Signature(C_restriction.Id(),C_class.Id()),0,MakeFunction1(E_domain_I_restriction,"domain_I_restriction")),MakeString("object.cl:423"))
  
  _ = F_attach_method(C_methods.AddMethod(Signature(C_class.Id(),C_class.Id(),C_set.Id()),0,MakeFunction2(E_methods_class,"methods_class")),MakeString("object.cl:425"))
  
  C_reify = MakeProperty("reify",1,C_claire)
  
  
  _ = F_attach_method(C_reify.AddMethod(Signature(C_listargs.Id(),C_void.Id()),0,MakeFunction1(E_reify_listargs,"reify_listargs")),MakeString("object.cl:431"))
  
  C_Core__star_stararg = MakeProperty("**arg",0,It)
  C_Core__star_stararg.Open = 0
  
  
  C_general_error = NewClass("general_error",C_error,C_claire)
  F_close_slot(C_general_error.AddSlot(C_mClaire_cause,ToType(C_any.Id()),CNULL))
  F_close_slot(C_general_error.AddSlot(C_arg,ToType(C_any.Id()),CNULL))
  
  _ = F_attach_method(C_self_print.AddMethod(Signature(C_general_error.Id(),C_void.Id()),1,MakeFunction1(E_self_print_general_error_Core,"self_print_general_error_Core")),MakeString("object.cl:444"))
  
  C_read_slot_error = NewClass("read_slot_error",C_error,C_claire)
  F_close_slot(C_read_slot_error.AddSlot(C_arg,ToType(C_any.Id()),CNULL))
  F_close_slot(C_read_slot_error.AddSlot(C_Core_wrong,ToType(C_any.Id()),CNULL))
  
  _ = F_attach_method(C_self_print.AddMethod(Signature(C_read_slot_error.Id(),C_void.Id()),1,MakeFunction1(E_self_print_read_slot_error_Core,"self_print_read_slot_error_Core")),MakeString("object.cl:449"))
  
  C_range_error = NewClass("range_error",C_error,C_claire)
  F_close_slot(C_range_error.AddSlot(C_mClaire_cause,ToType(C_any.Id()),CNULL))
  F_close_slot(C_range_error.AddSlot(C_arg,ToType(C_any.Id()),CNULL))
  F_close_slot(C_range_error.AddSlot(C_Core_wrong,ToType(C_any.Id()),CNULL))
  
  _ = F_attach_method(C_self_print.AddMethod(Signature(C_range_error.Id(),C_void.Id()),1,MakeFunction1(E_self_print_range_error_Core,"self_print_range_error_Core")),MakeString("object.cl:455"))
  
  C_selector_error = NewClass("selector_error",C_error,C_claire)
  F_close_slot(C_selector_error.AddSlot(C_selector,ToType(C_any.Id()),CNULL))
  F_close_slot(C_selector_error.AddSlot(C_arg,ToType(C_any.Id()),CNULL))
  
  _ = F_attach_method(C_self_print.AddMethod(Signature(C_selector_error.Id(),C_void.Id()),1,MakeFunction1(E_self_print_selector_error_Core,"self_print_selector_error_Core")),MakeString("object.cl:463"))
  
  C_return_error = NewClass("return_error",C_error,C_claire)
  F_close_slot(C_return_error.AddSlot(C_arg,ToType(C_any.Id()),CNULL))
  
  _ = F_attach_method(C_self_print.AddMethod(Signature(C_return_error.Id(),C_void.Id()),0,MakeFunction1(E_self_print_return_error_Core,"self_print_return_error_Core")),MakeString("object.cl:468"))
  
  _ = F_attach_method(C_self_print.AddMethod(Signature(C_system_error.Id(),C_void.Id()),1,MakeFunction1(E_self_print_system_error_Core,"self_print_system_error_Core")),MakeString("object.cl:517"))
  
  C_contradiction = NewClass("contradiction",C_exception,C_claire)
  
  _ = F_attach_method(C_self_print.AddMethod(Signature(C_contradiction.Id(),C_void.Id()),0,MakeFunction1(E_self_print_contradiction_Core,"self_print_contradiction_Core")),MakeString("object.cl:522"))
  
  _ = F_attach_method(C_format.AddMethod(Signature(C_string.Id(),C_list.Id(),C_void.Id()),1,MakeFunction2(E_format_string,"format_string")),MakeString("object.cl:549"))
  
  _ = F_attach_method(C_Core_tformat.AddMethod(Signature(C_string.Id(),
    C_integer.Id(),
    C_list.Id(),
    C_any.Id()),1,MakeFunction3(E_tformat_string,"tformat_string")),MakeString("object.cl:555"))
  
  _ = F_attach_method(C_princ.AddMethod(Signature(C_list.Id(),C_void.Id()),1,MakeFunction1(E_princ_list,"princ_list")),MakeString("object.cl:560"))
  
  _ = F_attach_method(C_princ.AddMethod(Signature(C_set.Id(),C_void.Id()),1,MakeFunction1(E_princ_set,"princ_set")),MakeString("object.cl:564"))
  
  C_global_variable = NewClass("global_variable",C_system_thing,C_claire)
  F_close_slot(C_global_variable.AddSlot(C_value,ToType(C_any.Id()),CNULL))
  F_close_slot(C_global_variable.AddSlot(C_range,ToType(C_type.Id()),C_any.Id()))
  F_close_slot(C_global_variable.AddSlot(C_store_ask,ToType(C_boolean.Id()),CFALSE.Id()))
  
  _ = F_attach_method(C_close.AddMethod(Signature(C_global_variable.Id(),C_global_variable.Id()),1,MakeFunction1(E_close_global_variable,"close_global_variable")),MakeString("object.cl:577"))
  
  _ = F_attach_method(C_self_eval.AddEvalMethod(Signature(C_global_variable.Id(),C_any.Id()),0,MakeFunction1(E_self_eval_global_variable,"self_eval_global_variable"),EVAL_global_variable),MakeString("object.cl:579"))
  
  C__inf_equal2 = ToOperation(new(ClaireOperation).IsNamed(C_operation,MakeSymbol("<=2",C_claire)))
  
  
  { 
    var expr EID
    { var _CL_obj *GlobalVariable
      C_contradiction_occurs = ToGlobalVariable(new(GlobalVariable).IsNamed(C_global_variable,MakeSymbol("contradiction_occurs",C_claire)))
      
      _CL_obj = C_contradiction_occurs
      _CL_obj.Range = ToType(C_contradiction.Id())
      _CL_obj.Value = new(Contradiction).Is(C_contradiction)
      expr = F_close_global_variable(_CL_obj)
      } 
    ErrorCheck(expr)} 
  
  _ = F_attach_method(C_contradiction_I.AddMethod(Signature(C_void.Id(),C_void.Id()),1,MakeFunction1(E_contradiction_I_void,"contradiction_I_void")),MakeString("object.cl:588"))
  
  { 
    var expr EID
    { var _CL_obj *GlobalVariable
      C_nil = ToGlobalVariable(new(GlobalVariable).IsNamed(C_global_variable,MakeSymbol("nil",C_claire)))
      
      _CL_obj = C_nil
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = CNIL.Id()
      expr = F_close_global_variable(_CL_obj)
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID
    { var _CL_obj *GlobalVariable
      C_claire_date = ToGlobalVariable(new(GlobalVariable).IsNamed(C_global_variable,MakeSymbol("claire_date",C_claire)))
      
      _CL_obj = C_claire_date
      _CL_obj.Range = ToType(C_string.Id())
      _CL_obj.Value = MakeString("Wednesday 01-01-2025").Id()
      expr = F_close_global_variable(_CL_obj)
      } 
    ErrorCheck(expr)} 
  
  C__I_equal = ToOperation(new(ClaireOperation).IsNamed(C_operation,MakeSymbol("!=",C_claire)))
  C__I_equal.Precedence = 60
  
  
  C__inf_inf = ToOperation(new(ClaireOperation).IsNamed(C_operation,MakeSymbol("<<",C_claire)))
  C__inf_inf.Precedence = 10
  
  
  C__sup_sup = ToOperation(new(ClaireOperation).IsNamed(C_operation,MakeSymbol(">>",C_claire)))
  C__sup_sup.Precedence = 10
  
  
  C_and = ToOperation(new(ClaireOperation).IsNamed(C_operation,MakeSymbol("and",C_claire)))
  
  
  C_or = ToOperation(new(ClaireOperation).IsNamed(C_operation,MakeSymbol("or",C_claire)))
  
  
  C_U = ToOperation(new(ClaireOperation).IsNamed(C_operation,MakeSymbol("U",C_claire)))
  C_U.Precedence = 50
  
  
  C__and = ToOperation(new(ClaireOperation).IsNamed(C_operation,MakeSymbol("&",C_claire)))
  
  
  C_meet = ToOperation(new(ClaireOperation).IsNamed(C_operation,MakeSymbol("meet",C_claire)))
  
  
  C_inherit_ask = ToOperation(new(ClaireOperation).IsNamed(C_operation,MakeSymbol("inherit?",C_claire)))
  
  
  C__dash_dash_ask = ToOperation(new(ClaireOperation).IsNamed(C_operation,MakeSymbol("--?",C_claire)))
  C__dash_dash_ask.Precedence = 30
  
  
  C__dash_dash_I = ToOperation(new(ClaireOperation).IsNamed(C_operation,MakeSymbol("--!",C_claire)))
  C__dash_dash_I.Precedence = 30
  
  
  C_Core_cpstack = MakeProperty("cpstack",1,It)
  
  
  C_pretty_printer = NewClass("pretty_printer",C_thing,C_claire)
  F_close_slot(C_pretty_printer.AddSlot(C_mClaire_cpretty,ToType(C_port.Id()),CNULL))
  F_close_slot(C_pretty_printer.AddSlot(C_mClaire_cprevious,ToType(C_integer.Id()),MakeInteger(0).Id()))
  F_close_slot(C_pretty_printer.AddSlot(C_mClaire_index,ToType(C_integer.Id()),MakeInteger(0).Id()))
  F_close_slot(C_pretty_printer.AddSlot(C_mClaire_width,ToType(C_integer.Id()),MakeInteger(75).Id()))
  F_close_slot(C_pretty_printer.AddSlot(C_mClaire_pprint,ToType(C_boolean.Id()),CFALSE.Id()))
  F_close_slot(C_pretty_printer.AddSlot(C_mClaire_pbreak,ToType(C_boolean.Id()),CFALSE.Id()))
  F_close_slot(C_pretty_printer.AddSlot(C_Core_cpstack,ToType(C_list.Id()),ToType(C_any.Id()).EmptyList().Id()))
  
  C_pretty = ToPrettyPrinter(new(PrettyPrinter).IsNamed(C_pretty_printer,MakeSymbol("pretty",C_claire)))
  C_pretty.Cpretty = F_port_I_void()
  C_pretty.Cpstack = CNIL
  
  
  C_apply_self_print = MakeProperty("apply_self_print",1,C_claire)
  
  
  C_short_enough = MakeProperty("short_enough",1,C_claire)
  
  
  _ = F_attach_method(C_print_in_string.AddMethod(Signature(C_void.Id(),C_void.Id()),0,MakeFunction1(E_print_in_string_void,"print_in_string_void")),MakeString("function.cl:70"))
  
  _ = F_attach_method(C_end_of_string.AddMethod(Signature(C_void.Id(),C_string.Id()),1,MakeFunction1(E_end_of_string_void,"end_of_string_void")),MakeString("function.cl:82"))
  
  _ = F_attach_method(C_mClaire_buffer_length.AddMethod(Signature(C_void.Id(),C_integer.Id()),0,MakeFunction1(E_buffer_length_void,"buffer_length_void")),MakeString("function.cl:85"))
  
  _ = F_attach_method(C_mClaire_buffer_set_length.AddMethod(Signature(C_integer.Id(),C_void.Id()),0,MakeFunction1(E_buffer_set_length_integer,"buffer_set_length_integer")),MakeString("function.cl:89"))
  
  _ = F_attach_method(C_apply_self_print.AddMethod(Signature(C_any.Id(),C_void.Id()),1,MakeFunction1(E_apply_self_print_any,"apply_self_print_any")),MakeString("function.cl:103"))
  
  _ = F_attach_method(C_self_print.AddMethod(Signature(C_any.Id(),C_void.Id()),1,MakeFunction1(E_self_print_any_Core,"self_print_any_Core")),MakeString("function.cl:112"))
  
  _ = F_attach_method(C_self_print.AddMethod(Signature(C_boolean.Id(),C_void.Id()),0,MakeFunction1(E_self_print_boolean_Core,"self_print_boolean_Core")),MakeString("function.cl:115"))
  
  _ = F_attach_method(C_self_print.AddMethod(Signature(C_function.Id(),C_void.Id()),0,MakeFunction1(E_self_print_function_Core,"self_print_function_Core")),MakeString("function.cl:118"))
  
  _ = F_attach_method(C_self_print.AddMethod(Signature(C_restriction.Id(),C_void.Id()),1,MakeFunction1(E_self_print_restriction_Core,"self_print_restriction_Core")),MakeString("function.cl:132"))
  
  C_much_too_far = NewClass("much_too_far",C_error,C_claire)
  
  _ = F_attach_method(C_print.AddMethod(Signature(C_any.Id(),C_void.Id()),1,MakeFunction1(E_print_any,"print_any")),MakeString("function.cl:156"))
  
  _ = F_attach_method(C_short_enough.AddMethod(Signature(C_integer.Id(),C_boolean.Id()),0,MakeFunction1(E_short_enough_integer,"short_enough_integer")),MakeString("function.cl:159"))
  
  _ = F_attach_method(C_Core_new_defaults.AddMethod(Signature(C_object.Id(),C_list.Id(),C_object.Id()),1,MakeFunction2(E_Core_new_defaults_object,"Core_new_defaults_object")),MakeString("function.cl:188"))
  
  _ = F_attach_method(C_not.AddMethod(Signature(C_any.Id(),C_boolean.Id()),0,MakeFunction1(E_not_any,"not_any")),MakeString("function.cl:198"))
  
  _ = F_attach_method(C__I_equal.AddMethod(Signature(C_any.Id(),C_any.Id(),C_boolean.Id()),0,MakeFunction2(E__I_equal_any,"_I_equal_any")),MakeString("function.cl:200"))
  
  _ = F_attach_method(C_owner.AddMethod(Signature(C_any.Id(),C_class.Id()),0,MakeFunction1(E_owner_any,"owner_any")),MakeString("function.cl:203"))
  
  _ = F_attach_method(C_known_ask.AddMethod(Signature(C_any.Id(),C_boolean.Id()),0,MakeFunction1(E_known_ask_any,"known_ask_any")),MakeString("function.cl:206"))
  
  _ = F_attach_method(C_unknown_ask.AddMethod(Signature(C_any.Id(),C_boolean.Id()),0,MakeFunction1(E_unknown_ask_any,"unknown_ask_any")),MakeString("function.cl:207"))
  
  _ = F_attach_method(C_check_in.AddMethod(Signature(C_any.Id(),C_type.Id(),C_any.Id()),1,MakeFunction2(E_check_in_any,"check_in_any")),MakeString("function.cl:213"))
  
  _ = F_attach_method(C_check_in.AddMethod(Signature(C_bag.Id(),
    C_class.Id(),
    C_type.Id(),
    C_bag.Id()),1,MakeFunction3(E_check_in_bag,"check_in_bag")),MakeString("function.cl:221"))
  
  _ = F_attach_method(C__inf.AddMethod(Signature(C_any.Id(),C_any.Id(),C_boolean.Id()),1,MakeFunction2(E__inf_any,"_inf_any")),MakeString("function.cl:224"))
  
  _ = F_attach_method(C__sup.AddMethod(Signature(C_any.Id(),C_any.Id(),C_boolean.Id()),1,MakeFunction2(E__sup_any,"_sup_any")),MakeString("function.cl:225"))
  
  _ = F_attach_method(C_unsafe.AddMethod(Signature(C_any.Id(),C_any.Id()),0,MakeFunction1(E_unsafe_any,"unsafe_any")),MakeString("function.cl:230"))
  
  _ = F_attach_method(C_ephemeral.AddMethod(Signature(C_class.Id(),C_any.Id()),1,MakeFunction1(E_ephemeral_class,"ephemeral_class")),MakeString("function.cl:248"))
  
  _ = F_attach_method(C_instanced.AddMethod(Signature(C_class.Id(),C_any.Id()),1,MakeFunction1(E_instanced_class,"instanced_class")),MakeString("function.cl:254"))
  
  _ = F_attach_method(C_abstract.AddMethod(Signature(C_class.Id(),C_any.Id()),1,MakeFunction1(E_abstract_class,"abstract_class")),MakeString("function.cl:264"))
  
  _ = F_attach_method(C_final.AddMethod(Signature(C_class.Id(),C_any.Id()),1,MakeFunction1(E_final_class,"final_class")),MakeString("function.cl:273"))
  
  F_attach_method(C_new.AddMethod(Signature(C_class.Id(),C_object.Id()),1,MakeFunction1(E_new_class1,"new_class1")),MakeString("function.cl:281")).Typing = MakeFunction1(E_new_class1_type,"new_class1_type").Id()
  
  F_attach_method(C_new.AddMethod(Signature(C_class.Id(),C_symbol.Id(),C_thing.Id()),1,MakeFunction2(E_new_class2,"new_class2")),MakeString("function.cl:286")).Typing = MakeFunction2(E_new_class2_type,"new_class2_type").Id()
  
  _ = F_attach_method(C_meet.AddMethod(Signature(C_class.Id(),C_class.Id(),C_class.Id()),0,MakeFunction2(E_meet_class,"meet_class")),MakeString("function.cl:300"))
  
  _ = F_attach_method(C_inherit_ask.AddMethod(Signature(C_class.Id(),C_class.Id(),C_boolean.Id()),0,MakeFunction2(E_inherit_ask_class,"inherit_ask_class")),MakeString("function.cl:307"))
  
  _ = F_attach_method(C_abstract.AddMethod(Signature(C_property.Id(),C_any.Id()),1,MakeFunction1(E_abstract_property,"abstract_property")),MakeString("function.cl:318"))
  
  _ = F_attach_method(C_final.AddMethod(Signature(C_relation.Id(),C_void.Id()),0,MakeFunction1(E_final_relation,"final_relation")),MakeString("function.cl:329"))
  
  _ = F_attach_method(C_close.AddMethod(Signature(C_module.Id(),C_module.Id()),0,MakeFunction1(E_close_module,"close_module")),MakeString("function.cl:345"))
  
  _ = F_attach_method(C_get_symbol.AddMethod(Signature(C_module.Id(),C_string.Id(),C_any.Id()),0,MakeFunction2(E_get_symbol_module,"#get_symbol_module")),MakeString("function.cl:348"))
  
  _ = F_attach_method(C_get_symbol.AddMethod(Signature(C_string.Id(),C_any.Id()),0,MakeFunction1(E_get_symbol_string,"get_symbol_string")),MakeString("function.cl:349"))
  
  _ = F_attach_method(C_time_get.AddMethod(Signature(C_void.Id(),C_integer.Id()),0,MakeFunction1(E_time_get_void,"#time_get_void")),MakeString("function.cl:360"))
  
  _ = F_attach_method(C_time_read.AddMethod(Signature(C_void.Id(),C_integer.Id()),0,MakeFunction1(E_time_read_void,"#time_read_void")),MakeString("function.cl:361"))
  
  _ = F_attach_method(C_time_set.AddMethod(Signature(C_void.Id(),C_void.Id()),0,MakeFunction1(E_time_set_void,"#time_set_void")),MakeString("function.cl:362"))
  
  _ = F_attach_method(C_time_show.AddMethod(Signature(C_void.Id(),C_void.Id()),0,MakeFunction1(E_time_show_void,"#time_show_void")),MakeString("function.cl:363"))
  
  _ = F_attach_method(C_gensym.AddMethod(Signature(C_void.Id(),C_symbol.Id()),0,MakeFunction1(E_gensym_void,"gensym_void")),MakeString("function.cl:365"))
  
  _ = F_attach_method(C_store.AddMethod(Signature(C_list.Id(),
    C_integer.Id(),
    C_any.Id(),
    C_any.Id()),1,MakeFunction3(E_store_list4,"store_list4")),MakeString("function.cl:370"))
  
  _ = F_attach_method(C_store.AddMethod(Signature(C_array.Id(),
    C_integer.Id(),
    C_any.Id(),
    C_any.Id()),1,MakeFunction3(E_store_array1,"store_array1")),MakeString("function.cl:374"))
  
  _ = F_attach_method(C_commit.AddMethod(Signature(C_integer.Id(),C_void.Id()),0,MakeFunction1(E_commit_integer,"commit_integer")),MakeString("function.cl:376"))
  
  _ = F_attach_method(C_backtrack.AddMethod(Signature(C_integer.Id(),C_void.Id()),0,MakeFunction1(E_backtrack_integer,"backtrack_integer")),MakeString("function.cl:377"))
  
  { 
    var expr EID
    { var _CL_obj *GlobalVariable
      C_world_plus = ToGlobalVariable(new(GlobalVariable).IsNamed(C_global_variable,MakeSymbol("world+",C_claire)))
      
      _CL_obj = C_world_plus
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = C_choice.Id()
      expr = F_close_global_variable(_CL_obj)
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID
    { var _CL_obj *GlobalVariable
      C_world_dash = ToGlobalVariable(new(GlobalVariable).IsNamed(C_global_variable,MakeSymbol("world-",C_claire)))
      
      _CL_obj = C_world_dash
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = C_backtrack.Id()
      expr = F_close_global_variable(_CL_obj)
      } 
    ErrorCheck(expr)} 
  
  { 
    var expr EID
    { var _CL_obj *GlobalVariable
      C_world_dash_I = ToGlobalVariable(new(GlobalVariable).IsNamed(C_global_variable,MakeSymbol("world-!",C_claire)))
      
      _CL_obj = C_world_dash_I
      _CL_obj.Range = ToType(CEMPTY.Id())
      _CL_obj.Value = C_commit.Id()
      expr = F_close_global_variable(_CL_obj)
      } 
    ErrorCheck(expr)} 
  
  _ = F_attach_method(C_store.AddMethod(Signature(C_listargs.Id(),C_any.Id()),0,MakeFunction1(E_store_listargs,"store_listargs")),MakeString("function.cl:387"))
  
  _ = F_attach_method(C_symbol_I.AddMethod(Signature(C_string.Id(),C_symbol.Id()),0,MakeFunction1(E_symbol_I_string2,"symbol_I_string2")),MakeString("function.cl:398"))
  
  _ = F_attach_method(C_nth_get.AddMethod(Signature(C_string.Id(),
    C_integer.Id(),
    C_integer.Id(),
    C_char.Id()),1,MakeFunction3(E_nth_get_string,"nth_get_string")),MakeString("function.cl:402"))
  
  _ = F_attach_method(C_nth_put.AddMethod(Signature(C_string.Id(),
    C_integer.Id(),
    C_char.Id(),
    C_integer.Id(),
    C_void.Id()),1,MakeFunction4(E_nth_put_string,"nth_put_string")),MakeString("function.cl:404"))
  
  _ = F_attach_method(C_getenv.AddMethod(Signature(C_string.Id(),C_string.Id()),0,MakeFunction1(E_getenv_string,"#getenv_string")),MakeString("function.cl:407"))
  
  _ = F_attach_method(C_get_value.AddMethod(Signature(C_string.Id(),C_any.Id()),0,MakeFunction1(E_value_string,"#value_string")),MakeString("function.cl:408"))
  
  _ = F_attach_method(C_get_value.AddMethod(Signature(C_module.Id(),C_string.Id(),C_any.Id()),0,MakeFunction2(E_value_module,"#value_module")),MakeString("function.cl:409"))
  
  _ = F_attach_method(C_externC.AddMethod(Signature(C_string.Id(),C_void.Id()),1,MakeFunction1(E_externC_string,"externC_string")),MakeString("function.cl:412"))
  
  F_attach_method(C_externC.AddMethod(Signature(C_string.Id(),C_class.Id(),C_any.Id()),1,MakeFunction2(E_externC_string2,"externC_string2")),MakeString("function.cl:413")).Typing = MakeFunction2(E_externC_string2_type,"externC_string2_type").Id()
  
  _ = F_attach_method(C_princ.AddMethod(Signature(C_string.Id(),C_integer.Id(),C_void.Id()),0,MakeFunction2(E_princ_string3,"princ_string3")),MakeString("function.cl:420"))
  
  _ = F_attach_method(C_make_string.AddMethod(Signature(C_symbol.Id(),C_string.Id()),1,MakeFunction1(E_make_string_symbol,"make_string_symbol")),MakeString("function.cl:424"))
  
  _ = F_attach_method(C_self_print.AddMethod(Signature(C_symbol.Id(),C_void.Id()),1,MakeFunction1(E_self_print_symbol_Core,"self_print_symbol_Core")),MakeString("function.cl:426"))
  
  F_attach_method(C__plus.AddMethod(Signature(C_integer.Id(),C_integer.Id(),C_integer.Id()),0,MakeFunction2(E__plus_integer,"_plus_integer")),MakeString("function.cl:431")).Typing = MakeFunction2(E__plus_integer_type,"_plus_integer_type").Id()
  
  F_attach_method(C__dash.AddMethod(Signature(C_integer.Id(),C_integer.Id(),C_integer.Id()),0,MakeFunction2(E__dash_integer1,"_dash_integer1")),MakeString("function.cl:432")).Typing = MakeFunction2(E__dash_integer1_type,"_dash_integer1_type").Id()
  
  _ = F_attach_method(C__dash_dash_ask.AddMethod(Signature(C_integer.Id(),C_integer.Id(),C_collection.Id()),0,MakeFunction2(E_sequence_integer,"#sequence_integer")),MakeString("function.cl:437"))
  
  _ = F_attach_method(C__dash_dash_I.AddMethod(Signature(C_integer.Id(),C_integer.Id(),C_list.Id()),0,MakeFunction2(E_list_integer,"#list_integer")),MakeString("function.cl:438"))
  
  _ = F_attach_method(C__inf_inf.AddMethod(Signature(C_integer.Id(),C_integer.Id(),C_integer.Id()),0,MakeFunction2(E__inf_inf_integer,"_inf_inf_integer")),MakeString("function.cl:441"))
  
  _ = F_attach_method(C__sup_sup.AddMethod(Signature(C_integer.Id(),C_integer.Id(),C_integer.Id()),0,MakeFunction2(E__sup_sup_integer,"_sup_sup_integer")),MakeString("function.cl:442"))
  
  _ = F_attach_method(C_and.AddMethod(Signature(C_integer.Id(),C_integer.Id(),C_integer.Id()),0,MakeFunction2(E_and_integer,"and_integer")),MakeString("function.cl:443"))
  
  _ = F_attach_method(C_or.AddMethod(Signature(C_integer.Id(),C_integer.Id(),C_integer.Id()),0,MakeFunction2(E_or_integer,"or_integer")),MakeString("function.cl:444"))
  
  _ = F_attach_method(C__inf.AddMethod(Signature(C_integer.Id(),C_integer.Id(),C_boolean.Id()),0,MakeFunction2(E__inf_integer,"_inf_integer")),MakeString("function.cl:447"))
  
  _ = F_attach_method(C__inf_equal.AddMethod(Signature(C_integer.Id(),C_integer.Id(),C_boolean.Id()),0,MakeFunction2(E__inf_equal_integer,"_inf_equal_integer")),MakeString("function.cl:448"))
  
  _ = F_attach_method(C__sup.AddMethod(Signature(C_integer.Id(),C_integer.Id(),C_boolean.Id()),0,MakeFunction2(E__sup_integer,"_sup_integer")),MakeString("function.cl:449"))
  
  _ = F_attach_method(C_nth.AddMethod(Signature(C_integer.Id(),C_integer.Id(),C_boolean.Id()),0,MakeFunction2(E_nth_integer,"nth_integer")),MakeString("function.cl:450"))
  
  _ = F_attach_method(C_abs.AddMethod(Signature(C_integer.Id(),C_integer.Id()),0,MakeFunction1(E_abs_integer,"abs_integer")),MakeString("function.cl:452"))
  
  _ = F_attach_method(C_random.AddMethod(Signature(C_integer.Id(),C_integer.Id(),C_integer.Id()),0,MakeFunction2(E_random_integer2,"random_integer2")),MakeString("function.cl:453"))
  
  _ = F_attach_method(C_factor_ask.AddMethod(Signature(C_integer.Id(),C_integer.Id(),C_boolean.Id()),1,MakeFunction2(E_factor_ask_integer,"factor_ask_integer")),MakeString("function.cl:456"))
  
  _ = F_attach_method(C_divide_ask.AddMethod(Signature(C_integer.Id(),C_integer.Id(),C_boolean.Id()),1,MakeFunction2(E_divide_ask_integer,"divide_ask_integer")),MakeString("function.cl:457"))
  
  F_attach_method(C_Id.AddMethod(Signature(C_any.Id(),C_any.Id()),0,MakeFunction1(E_Id_any,"Id_any")),MakeString("function.cl:458")).Typing = MakeFunction1(E_Id_any_type,"Id_any_type").Id()
  
  _ = F_attach_method(C_self_print.AddMethod(Signature(C_float.Id(),C_void.Id()),0,MakeFunction1(E_print_float,"#print_float")),MakeString("function.cl:461"))
  
  _ = F_attach_method(C__plus.AddMethod(Signature(C_float.Id(),C_float.Id(),C_float.Id()),0,MakeFunction2(E__plus_float,"_plus_float")),MakeString("function.cl:462"))
  
  _ = F_attach_method(C__dash.AddMethod(Signature(C_float.Id(),C_float.Id(),C_float.Id()),0,MakeFunction2(E__dash_float,"_dash_float")),MakeString("function.cl:463"))
  
  _ = F_attach_method(C__star.AddMethod(Signature(C_float.Id(),C_float.Id(),C_float.Id()),0,MakeFunction2(E__star_float,"_star_float")),MakeString("function.cl:464"))
  
  _ = F_attach_method(C__7.AddMethod(Signature(C_float.Id(),C_float.Id(),C_float.Id()),0,MakeFunction2(E__7_float,"_7_float")),MakeString("function.cl:465"))
  
  _ = F_attach_method(C__dash.AddMethod(Signature(C_float.Id(),C_float.Id()),0,MakeFunction1(E__dash_float2,"_dash_float2")),MakeString("function.cl:468"))
  
  _ = F_attach_method(C_string_I.AddMethod(Signature(C_float.Id(),C_string.Id()),1,MakeFunction1(E_string_I_float,"string_I_float")),MakeString("function.cl:470"))
  
  _ = F_attach_method(C_princ.AddMethod(Signature(C_float.Id(),C_integer.Id(),C_void.Id()),0,MakeFunction2(E_print_format_float,"#print_format_float")),MakeString("function.cl:471"))
  
  _ = F_attach_method(C_abs.AddMethod(Signature(C_float.Id(),C_float.Id()),0,MakeFunction1(E_abs_float,"abs_float")),MakeString("function.cl:473"))
  
  _ = F_attach_method(C_mClaire_printFDigit.AddMethod(Signature(C_float.Id(),C_integer.Id(),C_void.Id()),1,MakeFunction2(E_printFDigit_float,"printFDigit_float")),MakeString("function.cl:481"))
  
  _ = F_attach_method(C_mClaire_printFDigit.AddMethod(Signature(C_integer.Id(),C_integer.Id(),C_void.Id()),1,MakeFunction2(E_printFDigit_integer,"printFDigit_integer")),MakeString("function.cl:486"))
  
  _ = F_attach_method(C_mClaire_nth_object.AddMethod(Signature(C_list.Id(),C_integer.Id(),C_any.Id()),0,MakeFunction2(E_mClaire_nth_object_list,"mClaire_nth_object_list")),MakeString("function.cl:494"))
  
  _ = F_attach_method(C_nth_write.AddMethod(Signature(C_list.Id(),
    C_integer.Id(),
    C_any.Id(),
    C_any.Id()),1,MakeFunction3(E_nth_write_list,"nth_write_list")),MakeString("function.cl:507"))
  
  F_attach_method(C_min.AddMethod(Signature(C_method.Id(),C_set.Id(),C_any.Id()),1,MakeFunction2(E_min_method2,"min_method2")),MakeString("function.cl:514")).Typing = MakeFunction2(E_min_method2_type,"min_method2_type").Id()
  
  F_attach_method(C_max.AddMethod(Signature(C_method.Id(),C_set.Id(),C_any.Id()),1,MakeFunction2(E_max_method2,"max_method2")),MakeString("function.cl:520")).Typing = MakeFunction2(E_max_method2_type,"max_method2_type").Id()
  
  F_attach_method(C_min.AddMethod(Signature(C_method.Id(),C_list.Id(),C_any.Id()),1,MakeFunction2(E_min_method3,"min_method3")),MakeString("function.cl:528")).Typing = MakeFunction2(E_min_method3_type,"min_method3_type").Id()
  
  F_attach_method(C_max.AddMethod(Signature(C_method.Id(),C_list.Id(),C_any.Id()),1,MakeFunction2(E_max_method3,"max_method3")),MakeString("function.cl:535")).Typing = MakeFunction2(E_max_method3_type,"max_method3_type").Id()
  
  F_attach_method(C_random.AddMethod(Signature(C_list.Id(),C_any.Id()),0,MakeFunction1(E_random_list,"random_list")),MakeString("function.cl:540")).Typing = MakeFunction1(E_random_list_type,"random_list_type").Id()
  
  F_attach_method(C_last.AddMethod(Signature(C_list.Id(),C_any.Id()),1,MakeFunction1(E_last_list,"last_list")),MakeString("function.cl:545")).Typing = MakeFunction1(E_last_list_type,"last_list_type").Id()
  
  _ = F_attach_method(C_rmlast.AddMethod(Signature(C_list.Id(),C_list.Id()),1,MakeFunction1(E_rmlast_list,"rmlast_list")),MakeString("function.cl:548"))
  
  _ = F_attach_method(C_car.AddMethod(Signature(C_list.Id(),C_any.Id()),1,MakeFunction1(E_car_list,"car_list")),MakeString("function.cl:552"))
  
  _ = F_attach_method(C_sort.AddMethod(Signature(C_method.Id(),C_list.Id(),C_list.Id()),1,MakeFunction2(E_sort_method,"sort_method")),MakeString("function.cl:556"))
  
  _ = F_attach_method(C_Core_quicksort.AddMethod(Signature(C_list.Id(),
    C_method.Id(),
    C_integer.Id(),
    C_integer.Id(),
    C_void.Id()),1,MakeFunction4(E_quicksort_list,"quicksort_list")),MakeString("function.cl:576"))
  
  _ = F_attach_method(C_Core_build_powerset.AddMethod(Signature(C_list.Id(),C_set.Id()),0,MakeFunction1(E_build_powerset_list,"build_powerset_list")),MakeString("function.cl:585"))
  
  _ = F_attach_method(C__inf_inf.AddMethod(Signature(C_list.Id(),C_integer.Id(),C_list.Id()),0,MakeFunction2(E_skip_list,"#skip_list")),MakeString("function.cl:588"))
  
  _ = F_attach_method(C_make_copy_list.AddMethod(Signature(C_integer.Id(),C_any.Id(),C_list.Id()),0,MakeFunction2(E_make_copy_list_integer,"make_copy_list_integer")),MakeString("function.cl:594"))
  
  _ = F_attach_method(C_typed_copy_list.AddMethod(Signature(C_type.Id(),
    C_integer.Id(),
    C_any.Id(),
    C_list.Id()),0,MakeFunction3(E_typed_copy_list_type,"typed_copy_list_type")),MakeString("function.cl:602"))
  
  _ = F_attach_method(C_difference.AddMethod(Signature(C_set.Id(),C_set.Id(),C_set.Id()),0,MakeFunction2(E_difference_set,"difference_set")),MakeString("function.cl:605"))
  
  _ = F_attach_method(C__at.AddMethod(Signature(C_type.Id(),C_property.Id(),C_type.Id()),0,MakeFunction2(E__at_type,"#_at_type")),MakeString("function.cl:608"))
  
  _ = F_attach_method(C_self_print.AddMethod(Signature(C_array.Id(),C_void.Id()),1,MakeFunction1(E_self_print_array_Core,"self_print_array_Core")),MakeString("function.cl:618"))
  
  _ = F_attach_method(C_self_print.AddMethod(Signature(C_char.Id(),C_void.Id()),0,MakeFunction1(E_self_print_char_Core,"self_print_char_Core")),MakeString("function.cl:621"))
  
  _ = F_attach_method(C__inf_equal.AddMethod(Signature(C_char.Id(),C_char.Id(),C_boolean.Id()),0,MakeFunction2(E__inf_equal_char,"_inf_equal_char")),MakeString("function.cl:622"))
  
  _ = F_attach_method(C_random.AddMethod(Signature(C_boolean.Id(),C_boolean.Id()),0,MakeFunction1(E_random_boolean,"random_boolean")),MakeString("function.cl:626"))
  
  C__dash_dash = ToOperation(new(ClaireOperation).IsNamed(C_operation,MakeSymbol("--",C_claire)))
  C__dash_dash.Precedence = C__dot_dot.Precedence
  
  
  _ = F_attach_method(C_finite_ask.AddMethod(Signature(C_type.Id(),C_boolean.Id()),0,MakeFunction1(E_finite_ask_type,"finite_ask_type")),MakeString("types.cl:36"))
  
  _ = F_attach_method(C_Core_enumerate.AddMethod(Signature(C_any.Id(),C_list.Id()),1,MakeFunction1(E_enumerate_any,"enumerate_any")),MakeString("types.cl:51"))
  
  C__equaltype_ask = ToOperation(new(ClaireOperation).IsNamed(C_operation,MakeSymbol("=type?",C_claire)))
  
  
  _ = F_attach_method(C__equaltype_ask.AddMethod(Signature(C_type.Id(),C_type.Id(),C_boolean.Id()),0,MakeFunction2(E__equaltype_ask_any,"_equaltype_ask_any")),MakeString("types.cl:55"))
  
  _ = F_attach_method(C_sort_I.AddMethod(Signature(C_type.Id(),C_class.Id()),0,MakeFunction1(E_sort_I_type,"sort_I_type")),MakeString("types.cl:61"))
  
  _ = F_attach_method(C__Z.AddMethod(Signature(C_any.Id(),C_class.Id(),C_boolean.Id()),0,MakeFunction2(E__Z_any1,"_Z_any1")),MakeString("types.cl:65"))
  
  _ = F_attach_method(C_Core_belong.AddMethod(Signature(C_any.Id(),C_any.Id(),C_boolean.Id()),1,MakeFunction2(E_BELONG,"BELONG")),MakeString("types.cl:93"))
  
  _ = F_attach_method(C__Z.AddMethod(Signature(C_any.Id(),C_any.Id(),C_boolean.Id()),1,MakeFunction2(E_belong_to,"belong_to")),MakeString("types.cl:97"))
  
  _ = F_attach_method(C_self_print.AddMethod(Signature(C_Union.Id(),C_void.Id()),1,MakeFunction1(E_self_print_Union_Core,"self_print_Union_Core")),MakeString("types.cl:116"))
  
  _ = F_attach_method(C_finite_ask.AddMethod(Signature(C_Union.Id(),C_boolean.Id()),0,MakeFunction1(E_finite_ask_Union,"finite_ask_Union")),MakeString("types.cl:117"))
  
  _ = F_attach_method(C_self_print.AddMethod(Signature(C_Interval.Id(),C_void.Id()),1,MakeFunction1(E_self_print_Interval_Core,"self_print_Interval_Core")),MakeString("types.cl:121"))
  
  _ = F_attach_method(C_finite_ask.AddMethod(Signature(C_Interval.Id(),C_boolean.Id()),0,MakeFunction1(E_finite_ask_Interval,"finite_ask_Interval")),MakeString("types.cl:123"))
  
  _ = F_attach_method(C__dash_dash.AddMethod(Signature(C_integer.Id(),C_integer.Id(),C_Interval.Id()),1,MakeFunction2(E__dash_dash_integer,"_dash_dash_integer")),MakeString("types.cl:128"))
  
  _ = F_attach_method(C_self_print.AddMethod(Signature(C_Param.Id(),C_void.Id()),1,MakeFunction1(E_self_print_Param_Core,"self_print_Param_Core")),MakeString("types.cl:138"))
  
  _ = F_attach_method(C_finite_ask.AddMethod(Signature(C_Param.Id(),C_boolean.Id()),0,MakeFunction1(E_finite_ask_Param,"finite_ask_Param")),MakeString("types.cl:140"))
  
  _ = F_attach_method(C_self_print.AddMethod(Signature(C_subtype.Id(),C_void.Id()),1,MakeFunction1(E_self_print_subtype_Core,"self_print_subtype_Core")),MakeString("types.cl:149"))
  
  _ = F_attach_method(C_finite_ask.AddMethod(Signature(C_subtype.Id(),C_boolean.Id()),0,MakeFunction1(E_finite_ask_subtype,"finite_ask_subtype")),MakeString("types.cl:151"))
  
  _ = F_attach_method(C_nth.AddMethod(Signature(C_class.Id(),C_type.Id(),C_type.Id()),0,MakeFunction2(E_nth_class1,"nth_class1")),MakeString("types.cl:161"))
  
  _ = F_attach_method(C_nth.AddMethod(Signature(C_class.Id(),
    C_list.Id(),
    C_list.Id(),
    C_type.Id()),0,MakeFunction3(E_nth_class2,"nth_class2")),MakeString("types.cl:172"))
  
  _ = F_attach_method(C_Core_param_I.AddMethod(Signature(C_class.Id(),C_type.Id(),C_type.Id()),0,MakeFunction2(E_param_I_class,"param_I_class")),MakeString("types.cl:176"))
  
  _ = F_attach_method(C_nth.AddMethod(Signature(C_type.Id(),C_type.Id()),0,MakeFunction1(E_nth_type,"nth_type")),MakeString("types.cl:180"))
  
  _ = F_attach_method(C_finite_ask.AddMethod(Signature(C_tuple.Id(),C_boolean.Id()),0,MakeFunction1(E_finite_ask_tuple,"finite_ask_tuple")),MakeString("types.cl:183"))
  
  _ = F_attach_method(C_self_print.AddMethod(Signature(C_Reference.Id(),C_void.Id()),1,MakeFunction1(E_self_print_Reference_Core,"self_print_Reference_Core")),MakeString("types.cl:191"))
  
  _ = F_attach_method(C_get.AddMethod(Signature(C_Reference.Id(),C_any.Id(),C_any.Id()),0,MakeFunction2(E_get_Reference,"get_Reference")),MakeString("types.cl:195"))
  
  _ = F_attach_method(C_Reference_I.AddMethod(Signature(C_list.Id(),C_integer.Id(),C_Reference.Id()),0,MakeFunction2(E_Reference_I_list,"Reference_I_list")),MakeString("types.cl:199"))
  
  _ = F_attach_method(C__at.AddMethod(Signature(C_Reference.Id(),
    C_list.Id(),
    C_any.Id(),
    C_any.Id()),0,MakeFunction3(E__at_Reference,"_at_Reference")),MakeString("types.cl:204"))
  
  _ = F_attach_method(C_set_I.AddMethod(Signature(C_collection.Id(),C_set.Id()),1,MakeFunction1(E_set_I_collection,"set_I_collection")),MakeString("types.cl:216"))
  
  _ = F_attach_method(C_size.AddMethod(Signature(C_collection.Id(),C_integer.Id()),1,MakeFunction1(E_size_collection,"size_collection")),MakeString("types.cl:221"))
  
  _ = F_attach_method(C_set_I.AddMethod(Signature(C_set.Id(),C_set.Id()),0,MakeFunction1(E_set_I_set,"set_I_set")),MakeString("types.cl:226"))
  
  _ = F_attach_method(C_size.AddMethod(Signature(C_list.Id(),C_integer.Id()),0,MakeFunction1(E_size_list2_Core,"size_list2_Core")),MakeString("types.cl:229"))
  
  _ = F_attach_method(C_set_I.AddMethod(Signature(C_class.Id(),C_set.Id()),1,MakeFunction1(E_set_I_class,"set_I_class")),MakeString("types.cl:238"))
  
  _ = F_attach_method(C_size.AddMethod(Signature(C_class.Id(),C_integer.Id()),0,MakeFunction1(E_size_class,"size_class")),MakeString("types.cl:242"))
  
  _ = F_attach_method(C_set_I.AddMethod(Signature(C_Union.Id(),C_set.Id()),1,MakeFunction1(E_set_I_Union,"set_I_Union")),MakeString("types.cl:246"))
  
  _ = F_attach_method(C_size.AddMethod(Signature(C_Union.Id(),C_integer.Id()),1,MakeFunction1(E_size_Union,"size_Union")),MakeString("types.cl:250"))
  
  _ = F_attach_method(C_set_I.AddMethod(Signature(C_Interval.Id(),C_set.Id()),0,MakeFunction1(E_set_I_Interval,"set_I_Interval")),MakeString("types.cl:254"))
  
  _ = F_attach_method(C_size.AddMethod(Signature(C_Interval.Id(),C_integer.Id()),0,MakeFunction1(E_size_Interval,"size_Interval")),MakeString("types.cl:257"))
  
  _ = F_attach_method(C_set_I.AddMethod(Signature(C_Param.Id(),C_set.Id()),1,MakeFunction1(E_set_I_Param,"set_I_Param")),MakeString("types.cl:261"))
  
  _ = F_attach_method(C_size.AddMethod(Signature(C_Param.Id(),C_integer.Id()),1,MakeFunction1(E_size_Param,"size_Param")),MakeString("types.cl:262"))
  
  _ = F_attach_method(C_set_I.AddMethod(Signature(C_subtype.Id(),C_set.Id()),1,MakeFunction1(E_set_I_subtype,"set_I_subtype")),MakeString("types.cl:267"))
  
  _ = F_attach_method(C_size.AddMethod(Signature(C_subtype.Id(),C_integer.Id()),1,MakeFunction1(E_size_subtype,"size_subtype")),MakeString("types.cl:270"))
  
  _ = F_attach_method(C_set_I.AddMethod(Signature(C_tuple.Id(),C_set.Id()),1,MakeFunction1(E_set_I_tuple,"set_I_tuple")),MakeString("types.cl:282"))
  
  _ = F_attach_method(C_size.AddMethod(Signature(C_tuple.Id(),C_integer.Id()),1,MakeFunction1(E_size_tuple,"size_tuple")),MakeString("types.cl:288"))
  
  { 
    var expr EID
    expr = F_ephemeral_class(C_Union)
    if !ErrorIn(expr) {
    expr = F_ephemeral_class(C_Param)
    if !ErrorIn(expr) {
    expr = F_ephemeral_class(C_Interval)
    if !ErrorIn(expr) {
    expr = F_ephemeral_class(C_subtype)
    }}}
    ErrorCheck(expr)} 
  
  _ = F_attach_method(C_U.AddMethod(Signature(C_type.Id(),C_type.Id(),C_type.Id()),0,MakeFunction2(E_U_type,"U_type")),MakeString("types.cl:331"))
  
  F_attach_method(C__dot_dot.AddMethod(Signature(C_integer.Id(),C_integer.Id(),F_nth_class1(C_type,ToType(C_integer.Id())).Id()),0,MakeFunction2(E__dot_dot_integer,"_dot_dot_integer")),MakeString("types.cl:338")).Typing = MakeFunction2(E__dot_dot_integer_type,"_dot_dot_integer_type").Id()
  
  C_but = ToOperation(new(ClaireOperation).IsNamed(C_operation,MakeSymbol("but",C_claire)))
  
  
  F_attach_method(C_but.AddMethod(Signature(C_any.Id(),C_any.Id(),C_any.Id()),1,MakeFunction2(E_but_any,"but_any")),MakeString("types.cl:345")).Typing = MakeFunction2(E_but_any_type,"but_any_type").Id()
  
  C__backslash = ToOperation(new(ClaireOperation).IsNamed(C_operation,MakeSymbol("\\",C_claire)))
  C__backslash.Precedence = C_U.Precedence
  
  
  _ = F_attach_method(C__backslash.AddMethod(Signature(C_type.Id(),C_type.Id(),C_set.Id()),1,MakeFunction2(E__backslash_type,"_backslash_type")),MakeString("types.cl:349"))
  
  C_glb = ToOperation(new(ClaireOperation).IsNamed(C_operation,MakeSymbol("glb",C_claire)))
  C_glb.Precedence = C__exp.Precedence
  C_glb.Domain = ToType(C_type_expression.Id())
  C_glb.Range = ToType(C_type.Id())
  
  
  _ = F_attach_method(C_glb.AddMethod(Signature(C_set.Id(),C_type.Id(),C_set.Id()),0,MakeFunction2(E_glb_set,"glb_set")),MakeString("types.cl:362"))
  
  _ = F_attach_method(C_glb.AddMethod(Signature(C_Union.Id(),C_type.Id(),C_type.Id()),0,MakeFunction2(E_glb_Union,"glb_Union")),MakeString("types.cl:364"))
  
  _ = F_attach_method(C_glb.AddMethod(Signature(C_Interval.Id(),C_type.Id(),C_type.Id()),0,MakeFunction2(E_glb_Interval,"glb_Interval")),MakeString("types.cl:377"))
  
  _ = F_attach_method(C_glb.AddMethod(Signature(C_class.Id(),C_type.Id(),C_type.Id()),0,MakeFunction2(E_glb_class,"glb_class")),MakeString("types.cl:383"))
  
  _ = F_attach_method(C_glb.AddMethod(Signature(C_Param.Id(),C_type.Id(),C_type.Id()),0,MakeFunction2(E_glb_Param,"glb_Param")),MakeString("types.cl:398"))
  
  _ = F_attach_method(C_glb.AddMethod(Signature(C_subtype.Id(),C_type.Id(),C_type.Id()),0,MakeFunction2(E_glb_subtype,"glb_subtype")),MakeString("types.cl:413"))
  
  _ = F_attach_method(C_glb.AddMethod(Signature(C_tuple.Id(),C_type.Id(),C_type.Id()),0,MakeFunction2(E_glb_tuple,"glb_tuple")),MakeString("types.cl:423"))
  
  _ = F_attach_method(C_glb.AddMethod(Signature(C_Reference.Id(),C_type.Id(),C_type.Id()),0,MakeFunction2(E_glb_Reference,"glb_Reference")),MakeString("types.cl:427"))
  
  _ = F_attach_method(C__exp.AddMethod(Signature(C_type.Id(),C_type.Id(),C_type.Id()),0,MakeFunction2(E__exp_type,"_exp_type")),MakeString("types.cl:430"))
  
  _ = F_attach_method(C_join.AddMethod(Signature(C_class.Id(),C_class.Id(),C_type.Id()),0,MakeFunction2(E_join_class,"join_class")),MakeString("types.cl:439"))
  
  _ = F_attach_method(C__exp.AddMethod(Signature(C_list.Id(),C_list.Id(),C_list.Id()),0,MakeFunction2(E__exp_list,"_exp_list")),MakeString("types.cl:449"))
  
  _ = F_attach_method(C_Uall.AddMethod(Signature(C_list.Id(),C_type.Id()),0,MakeFunction1(E_Uall_list,"Uall_list")),MakeString("types.cl:453"))
  
  C_Core__inf_equalt = ToOperation(new(ClaireOperation).IsNamed(C_operation,MakeSymbol("<=t",It)))
  C_Core__inf_equalt.Precedence = C__inf_equal.Precedence
  C_Core__inf_equalt.Domain = ToType(C_type.Id())
  C_Core__inf_equalt.Range = ToType(C_boolean.Id())
  
  
  _ = F_attach_method(C_Core__inf_equalt.AddMethod(Signature(C_type.Id(),C_type.Id(),C_boolean.Id()),0,MakeFunction2(E__inf_equalt_type,"_inf_equalt_type")),MakeString("types.cl:462"))
  
  _ = F_attach_method(C__inf_equal.AddMethod(Signature(C_type_expression.Id(),C_type_expression.Id(),C_boolean.Id()),1,MakeFunction2(E__inf_equal_type_expression,"_inf_equal_type_expression")),MakeString("types.cl:469"))
  
  C_Core__Zt = ToOperation(new(ClaireOperation).IsNamed(C_operation,MakeSymbol("%t",It)))
  C_Core__Zt.Precedence = C__inf_equal.Precedence
  C_Core__Zt.Domain = ToType(C_type.Id())
  C_Core__Zt.Range = ToType(C_boolean.Id())
  
  
  _ = F_attach_method(C_Core__Zt.AddMethod(Signature(C_any.Id(),C_type.Id(),C_boolean.Id()),0,MakeFunction2(E_Core__Zt_any,"Core__Zt_any")),MakeString("types.cl:475"))
  
  _ = F_attach_method(C_less_ask.AddMethod(Signature(C_type_expression.Id(),C_type_expression.Id(),C_boolean.Id()),0,MakeFunction2(E_less_ask_type_expression,"less_ask_type_expression")),MakeString("types.cl:479"))
  
  _ = F_attach_method(C_member.AddMethod(Signature(C_type.Id(),C_type.Id()),0,MakeFunction1(E_member_type,"member_type")),MakeString("types.cl:498"))
  
  _ = F_attach_method(C_Core_of_extract.AddMethod(Signature(C_type.Id(),C_type.Id()),0,MakeFunction1(E_of_extract_type,"of_extract_type")),MakeString("types.cl:512"))
  
  _ = F_attach_method(C_unique_ask.AddMethod(Signature(C_type.Id(),C_boolean.Id()),0,MakeFunction1(E_unique_ask_type,"unique_ask_type")),MakeString("types.cl:519"))
  
  _ = F_attach_method(C_the.AddMethod(Signature(C_type.Id(),C_any.Id()),1,MakeFunction1(E_the_type,"the_type")),MakeString("types.cl:522"))
  
  _ = F_attach_method(C_integer_I.AddMethod(Signature(F_nth_class1(C_set,ToType(C_integer.Id())).Id(),C_integer.Id()),1,MakeFunction1(E_integer_I_set,"integer_I_set")),MakeString("types.cl:528"))
  
  _ = F_attach_method(C_make_set.AddMethod(Signature(C_integer.Id(),C_set.Id()),0,MakeFunction1(E_make_set_integer,"make_set_integer")),MakeString("types.cl:530"))
  
  _ = F_attach_method(C_Core_abstract_type.AddMethod(Signature(C_set.Id(),C_type.Id()),0,MakeFunction1(E_abstract_type_set,"abstract_type_set")),MakeString("types.cl:542"))
  
  _ = F_attach_method(C_Core_abstract_type.AddMethod(Signature(C_operation.Id(),
    C_type.Id(),
    C_type.Id(),
    C_type.Id()),0,MakeFunction3(E_abstract_type_operation,"abstract_type_operation")),MakeString("types.cl:560"))
  
  F_set_range_property(C_subclass,C_class,F_nth_class2(C_set,MakeList(ToType(C_any.Id()),C_of.Id()),MakeConstantList(MakeConstantSet(C_class.Id()).Id())))
  F_set_range_property(C_ancestors,C_class,F_nth_class2(C_list,MakeList(ToType(C_any.Id()),C_of.Id()),MakeConstantList(MakeConstantSet(C_class.Id()).Id())))
  F_set_range_property(C_descendants,C_class,F_nth_class2(C_set,MakeList(ToType(C_any.Id()),C_of.Id()),MakeConstantList(MakeConstantSet(C_class.Id()).Id())))
  F_set_range_property(C_mClaire_definition,C_property,F_nth_class2(C_list,MakeList(ToType(C_any.Id()),C_of.Id()),MakeConstantList(MakeConstantSet(C_restriction.Id()).Id())))
  F_set_range_property(C_restrictions,C_property,F_nth_class2(C_list,MakeList(ToType(C_any.Id()),C_of.Id()),MakeConstantList(MakeConstantSet(C_restriction.Id()).Id())))
  F_set_range_property(C_domain,C_restriction,F_nth_class2(C_list,MakeList(ToType(C_any.Id()),C_of.Id()),MakeConstantList(MakeConstantSet(C_type_expression.Id()).Id())))
  F_set_range_property(C_slots,C_class,F_nth_class2(C_list,MakeList(ToType(C_any.Id()),C_of.Id()),MakeConstantList(MakeConstantSet(C_slot.Id()).Id())))
  
  _ = F_attach_method(C_Core_first_arg_type.AddMethod(Signature(C_type.Id(),C_type.Id(),C_type.Id()),0,MakeFunction2(E_first_arg_type_type,"first_arg_type_type")),MakeString("types.cl:572"))
  
  _ = F_attach_method(C_Core_first_arg_type.AddMethod(Signature(C_type.Id(),
    C_type.Id(),
    C_type.Id(),
    C_type.Id()),0,MakeFunction3(E_first_arg_type_type2,"first_arg_type_type2")),MakeString("types.cl:573"))
  
  _ = F_attach_method(C_Core_second_arg_type.AddMethod(Signature(C_type.Id(),C_type.Id(),C_type.Id()),0,MakeFunction2(E_second_arg_type_type,"second_arg_type_type")),MakeString("types.cl:574"))
  
  _ = F_attach_method(C_Core_meet_arg_types.AddMethod(Signature(C_type.Id(),C_type.Id(),C_type.Id()),0,MakeFunction2(E_meet_arg_types_type,"meet_arg_types_type")),MakeString("types.cl:575"))
  
  _ = F_attach_method(C_Core_first_member_type.AddMethod(Signature(C_type.Id(),C_type.Id(),C_type.Id()),0,MakeFunction2(E_first_member_type_type,"first_member_type_type")),MakeString("types.cl:576"))
  
  _ = F_attach_method(C_Core_nth_arg_type.AddMethod(Signature(C_type.Id(),C_type.Id(),C_type.Id()),1,MakeFunction2(E_Core_nth_arg_type_type,"Core_nth_arg_type_type")),MakeString("types.cl:579"))
  
  F_attach_method(C_nth_get.AddMethod(Signature(C_array.Id(),C_integer.Id(),C_any.Id()),0,MakeFunction2(E_nth_get_array,"nth_get_array")),MakeString("types.cl:582")).Typing = MakeFunction2(E_nth_get_array_type,"nth_get_array_type").Id()
  
  F_attach_method(C_make_array.AddMethod(Signature(C_integer.Id(),
    C_type.Id(),
    C_any.Id(),
    C_array.Id()),0,MakeFunction3(E_make_array_integer,"#make_array_integer")),MakeString("types.cl:588")).Typing = MakeFunction3(E_make_array_integer_type,"make_array_integer_type").Id()
  
  F_attach_method(C_make_list.AddMethod(Signature(C_integer.Id(),
    C_type.Id(),
    C_any.Id(),
    C_list.Id()),0,MakeFunction3(E_make_list_integer2,"make_list_integer2")),MakeString("types.cl:591")).Typing = MakeFunction3(E_make_list_integer2_type,"make_list_integer2_type").Id()
  
  F_attach_method(C_make_set.AddMethod(Signature(F_nth_class2(C_array,MakeList(ToType(C_any.Id()),C_of.Id()),MakeConstantList(C_any.Id())).Id(),C_set.Id()),0,MakeFunction1(E_make_set_array,"make_set_array")),MakeString("types.cl:594")).Typing = MakeFunction1(E_make_set_array_type,"make_set_array_type").Id()
  
  F_attach_method(C_list_I.AddMethod(Signature(F_nth_class2(C_array,MakeList(ToType(C_any.Id()),C_of.Id()),MakeConstantList(C_any.Id())).Id(),C_list.Id()),0,MakeFunction1(E_list_I_array,"#list_I_array")),MakeString("types.cl:599")).Typing = MakeFunction1(E_list_I_array_type,"list_I_array_type").Id()
  
  F_attach_method(C_array_I.AddMethod(Signature(F_nth_class2(C_list,MakeList(ToType(C_any.Id()),C_of.Id()),MakeConstantList(C_any.Id())).Id(),C_array.Id()),0,MakeFunction1(E_array_I_list,"#array_I_list")),MakeString("types.cl:601")).Typing = MakeFunction1(E_array_I_list_type,"array_I_list_type").Id()
  
  F_attach_method(C_set_I.AddMethod(Signature(F_nth_class2(C_list,MakeList(ToType(C_any.Id()),C_of.Id()),MakeConstantList(C_any.Id())).Id(),C_set.Id()),0,MakeFunction1(E_set_I_list,"set_I_list")),MakeString("types.cl:604")).Typing = MakeFunction1(E_set_I_list_type,"set_I_list_type").Id()
  
  F_attach_method(C_list_I.AddMethod(Signature(F_nth_class2(C_set,MakeList(ToType(C_any.Id()),C_of.Id()),MakeConstantList(C_any.Id())).Id(),C_list.Id()),0,MakeFunction1(E_list_I_set,"#list_I_set")),MakeString("types.cl:606")).Typing = MakeFunction1(E_list_I_set_type,"list_I_set_type").Id()
  
  _ = F_attach_method(C_Core_thing_type_class.AddMethod(Signature(C_type.Id(),C_type.Id(),C_type.Id()),0,MakeFunction2(E_Core_thing_type_class_type2,"Core_thing_type_class_type2")),MakeString("types.cl:609"))
  
  _ = F_attach_method(C_Core_object_type_class.AddMethod(Signature(C_type.Id(),C_type.Id()),0,MakeFunction1(E_Core_object_type_class_type,"Core_object_type_class_type")),MakeString("types.cl:610"))
  
  { 
    var r *ClaireRestriction
    _ = r
    var r_iter *ClaireAny
    for _,r_iter = range(C_copy.Restrictions.ValuesO()){ 
      r = ToRestriction(r_iter)
      ToMethod(r.Id()).Typing = C_Id.Id()
      } 
    } 
  { 
    var r *ClaireRestriction
    _ = r
    var r_iter *ClaireAny
    for _,r_iter = range(C_empty.Restrictions.ValuesO()){ 
      r = ToRestriction(r_iter)
      ToMethod(r.Id()).Typing = C_Id.Id()
      } 
    } 
  { 
    var r *ClaireRestriction
    _ = r
    var r_iter *ClaireAny
    for _,r_iter = range(C_sort.Restrictions.ValuesO()){ 
      r = ToRestriction(r_iter)
      ToMethod(r.Id()).Typing = C_Core_second_arg_type.Id()
      } 
    } 
  { 
    var r *ClaireRestriction
    _ = r
    var r_iter *ClaireAny
    for _,r_iter = range(C__7_plus.Restrictions.ValuesO()){ 
      r = ToRestriction(r_iter)
      ToMethod(r.Id()).Typing = C_Core_meet_arg_types.Id()
      } 
    } 
  ToMethod(F__at_property2(C_mClaire_new_I,MakeConstantList(C_class.Id(),C_symbol.Id())).Id()).Typing = C_Core_thing_type_class.Id()
  ToMethod(F__at_property2(C_mClaire_new_I,MakeConstantList(C_class.Id())).Id()).Typing = C_Core_object_type_class.Id()
  ToMethod(F__at_property1(C_nth_get,C_array).Id()).Typing = C_Core_first_member_type.Id()
  ToMethod(F__at_property1(C_nth,C_list).Id()).Typing = C_Core_nth_arg_type.Id()
  ToMethod(F__at_property1(C_nth,C_array).Id()).Typing = C_Core_nth_arg_type.Id()
  ToMethod(F__at_property1(C_nth,C_set).Id()).Typing = C_Core_nth_arg_type.Id()
  { 
    var r *ClaireRestriction
    _ = r
    var r_iter *ClaireAny
    for _,r_iter = range(C_nth_plus.Restrictions.ValuesO()){ 
      r = ToRestriction(r_iter)
      ToMethod(r.Id()).Typing = C_Core_first_arg_type.Id()
      } 
    } 
  { 
    var r *ClaireRestriction
    _ = r
    var r_iter *ClaireAny
    for _,r_iter = range(C_add.Restrictions.ValuesO()){ 
      r = ToRestriction(r_iter)
      if (r.Domain.Length() == 2) { 
        ToMethod(r.Id()).Typing = C_Core_first_arg_type.Id()
        } 
      } 
    } 
  { 
    var r *ClaireRestriction
    _ = r
    var r_iter *ClaireAny
    for _,r_iter = range(C_delete.Restrictions.ValuesO()){ 
      r = ToRestriction(r_iter)
      if (r.Domain.Length() == 2) { 
        ToMethod(r.Id()).Typing = C_Core_first_arg_type.Id()
        } 
      } 
    } 
  
  } 

