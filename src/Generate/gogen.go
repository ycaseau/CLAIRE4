/***** CLAIRE Compilation of file /Users/ycaseau/Dropbox/src/clairev4.03/src/compile/gogen.cl 
         [version 4.0.04 / safety 5] Sunday 12-26-2021 17:16:12 *****/

package Generate
import (_ "fmt"
	. "Kernel"
	"Core"
	"Language"
	"Reader"
	"Optimize"
)

//-------- dumb function to prevent import errors --------
func import_g0041() { 
    _ = Core.It
    _ = Language.It
    _ = Reader.It
    _ = Optimize.It
    } 
  
  
//+-------------------------------------------------------------+
//| CLAIRE                                                      |
//| gogen.cl                                                    |
//| Copyright (C) 2020 - 2021 Yves Caseau. All Rights Reserved  |
//| cf. copyright info in file object.cl: about()               |
//+-------------------------------------------------------------+
// *******************************************************************
// * Contents                                                        *
// *     Part 1: definition of the code producer                     *
// *     Part 2: utilities for file generation                       *
// *     Part 3: interface declarations                              *
// *     Part 4: use of language dependent patterns (macros)         *
// *     Part 5: Utilities                                           *
// *******************************************************************
// renaming philosophy:
// keyword => become allcaps and if not good enough, add _CL_
// class => add Claire
// debug
/* {1} The go function for: new_block(tag:string) [status=0] */
func F_Generate_new_block_string (tag *ClaireString )  { 
    // procedure body with s = void 
if (ToGenerateGoProducer(Optimize.C_PRODUCER.Value).Debug_ask == CTRUE) { 
      PRINC("/* ")
      F_princ_string(tag)
      PRINC(":")
      F_princ_integer(Optimize.C_OPT.Level)
      PRINC(" */")
      } 
    F_Generate_new_block_void()
    } 
  
// The EID go function for: new_block @ string (throw: false) 
func E_Generate_new_block_string (tag EID) EID { 
    F_Generate_new_block_string(ToString(OBJ(tag)) )
    return EVOID} 
  
/* {1} The go function for: close_block(tag:string) [status=0] */
func F_Generate_close_block_string (tag *ClaireString )  { 
    // procedure body with s = void 
if (ToGenerateGoProducer(Optimize.C_PRODUCER.Value).Debug_ask == CTRUE) { 
      PRINC("/* ")
      F_princ_string(tag)
      PRINC("-")
      F_princ_integer((Optimize.C_OPT.Level-1))
      PRINC(" */")
      } 
    F_Generate_close_block_void()
    } 
  
// The EID go function for: close_block @ string (throw: false) 
func E_Generate_close_block_string (tag EID) EID { 
    F_Generate_close_block_string(ToString(OBJ(tag)) )
    return EVOID} 
  
/* {1} The go function for: finish_block(tag:string) [status=0] */
func F_Generate_finish_block_string (tag *ClaireString )  { 
    // procedure body with s = void 
if (ToGenerateGoProducer(Optimize.C_PRODUCER.Value).Debug_ask == CTRUE) { 
      PRINC("/* ")
      F_princ_string(tag)
      PRINC("!")
      F_princ_integer((Optimize.C_OPT.Level-1))
      PRINC(" */")
      } 
    F_Generate_finish_block_void()
    } 
  
// The EID go function for: finish_block @ string (throw: false) 
func E_Generate_finish_block_string (tag EID) EID { 
    F_Generate_finish_block_string(ToString(OBJ(tag)) )
    return EVOID} 
  
// adds a distinct ID to a variable name that may be reused
/* {1} The go function for: genvar(v:string) [status=0] */
func F_Generate_genvar_string (v *ClaireString ) *ClaireString  { 
    ToGenerateGoProducer(Optimize.C_PRODUCER.Value).Varsym = (ToGenerateGoProducer(Optimize.C_PRODUCER.Value).Varsym+1)
    /*integer->integer*/return  F_append_string(v,F_string_I_integer(ToGenerateGoProducer(Optimize.C_PRODUCER.Value).Varsym))
    } 
  
// The EID go function for: genvar @ string (throw: false) 
func E_Generate_genvar_string (v EID) EID { 
    return EID{F_Generate_genvar_string(ToString(OBJ(v)) ).Id(),0}} 
  
// *******************************************************************
// *       Part 1: definition of the code producer               *
// *******************************************************************
// definition of the instance
// to do : update the reserved names progressively - note that classes do not need protection since
// ClaireX is added to X
// use this producer
// makes an ident (string) from a variable's name - in CLAIRE4 we got rid of .naming option
/* {1} The go function for: c_string(c:go_producer,self:Variable) [status=1] */
func F_Generate_c_string_go_producer1 (c *GenerateGoProducer ,self *ClaireVariable ) EID { 
    var Result EID 
    
    Core.F_print_in_string_void()
    F_iClaire_ident_go_producer2(c,self.Pname)
    Result = Core.F_end_of_string_void()
    return Result} 
  
// The EID go function for: c_string @ list<type_expression>(go_producer, Variable) (throw: true) 
func E_Generate_c_string_go_producer1 (c EID,self EID) EID { 
    return F_Generate_c_string_go_producer1(ToGenerateGoProducer(OBJ(c)),To_Variable(OBJ(self)) )} 
  
/* {1} The go function for: c_string(c:go_producer,self:symbol) [status=1] */
func F_Generate_c_string_go_producer2 (c *GenerateGoProducer ,self *ClaireSymbol ) EID { 
    var Result EID 
    Core.F_print_in_string_void()
    F_iClaire_ident_go_producer2(c,self)
    Result = Core.F_end_of_string_void()
    return Result} 
  
// The EID go function for: c_string @ list<type_expression>(go_producer, symbol) (throw: true) 
func E_Generate_c_string_go_producer2 (c EID,self EID) EID { 
    return F_Generate_c_string_go_producer2(ToGenerateGoProducer(OBJ(c)),ToSymbol(OBJ(self)) )} 
  
// print a symbol for a variable
// two issues : replace with a dictionary some day (CLAIRE4) + why does c_string exist ?
// notice that ident should only exist for <strings> that will exist directly in Go code 
/* {1} The go function for: iClaire/ident(c:go_producer,v:Variable) [status=0] */
func F_iClaire_ident_go_producer1 (c *GenerateGoProducer ,v *ClaireVariable )  { 
    // procedure body with s = void 
{ var s *ClaireSymbol   = v.Pname
      { var n int  = F_index_list(c.BadNames,s.Id())
        if (n == 0) { 
          F_c_princ_string(s.String_I())
          } else {
          ToSymbol(c.GoodNames.At(n-1)).CPrinc()
          } 
        } 
      } 
    } 
  
// The EID go function for: iClaire/ident @ list<type_expression>(go_producer, Variable) (throw: false) 
func E_iClaire_ident_go_producer1 (c EID,v EID) EID { 
    F_iClaire_ident_go_producer1(ToGenerateGoProducer(OBJ(c)),To_Variable(OBJ(v)) )
    return EVOID} 
  
// print a symbol for the structure definition  => use c_princ to get rid of special chars
/* {1} The go function for: iClaire/ident(c:go_producer,s:symbol) [status=0] */
func F_iClaire_ident_go_producer2 (c *GenerateGoProducer ,s *ClaireSymbol )  { 
    // procedure body with s = void 
{ var n int  = F_index_list(c.BadNames,s.Id())
      if (n == 0) { 
        F_c_princ_string(s.String_I())
        } else {
        ToSymbol(c.GoodNames.At(n-1)).CPrinc()
        } 
      } 
    } 
  
// The EID go function for: iClaire/ident @ list<type_expression>(go_producer, symbol) (throw: false) 
func E_iClaire_ident_go_producer2 (c EID,s EID) EID { 
    F_iClaire_ident_go_producer2(ToGenerateGoProducer(OBJ(c)),ToSymbol(OBJ(s)) )
    return EVOID} 
  
// new in claire4: printd the go identifier asociated with symbol s
// cap_ident(c,x) uses capitalization : used for Class and Method, required by Go for identifiers to be visible
// notice that we print explicitly s.module! (namespace) if not claire, to avoid c name conflicts
/* {1} The go function for: cap_ident(s:symbol) [status=0] */
func F_Generate_cap_ident_symbol (s *ClaireSymbol )  { 
    // procedure body with s = void 
F_Generate_capitalized_ident_symbol(s,s.Module_I())
    } 
  
// The EID go function for: cap_ident @ symbol (throw: false) 
func E_Generate_cap_ident_symbol (s EID) EID { 
    F_Generate_cap_ident_symbol(ToSymbol(OBJ(s)) )
    return EVOID} 
  
// this is the capitalized ident for s in namespace m
/* {1} The go function for: capitalized_ident(s:symbol,m:module) [status=0] */
func F_Generate_capitalized_ident_symbol (s *ClaireSymbol ,m *ClaireModule )  { 
    // procedure body with s = void 
{ var n int  = F_index_list(ToGenerateGoProducer(Optimize.C_PRODUCER.Value).BadNames,s.Id())
      if (n == 0) { 
        if (m.Id() != C_claire.Id()) { 
          F_c_princ_string(F_Generate_capitalize_string(m.Name.String_I()))
          F_Generate_add_underscore_symbol(s)
          } 
        F_c_princ_string(F_Generate_capitalize_string(s.String_I()))
        } else {
        ToSymbol(ToGenerateGoProducer(Optimize.C_PRODUCER.Value).GoodNames.At(n-1)).CPrinc()
        } 
      } 
    } 
  
// The EID go function for: capitalized_ident @ symbol (throw: false) 
func E_Generate_capitalized_ident_symbol (s EID,m EID) EID { 
    F_Generate_capitalized_ident_symbol(ToSymbol(OBJ(s)),ToModule(OBJ(m)) )
    return EVOID} 
  
// short version (we do not care about the namespace) 
/* {1} The go function for: cap_short(s:symbol) [status=0] */
func F_Generate_cap_short_symbol (s *ClaireSymbol )  { 
    // procedure body with s = void 
F_Generate_capitalized_ident_symbol(s,C_claire)
    } 
  
// The EID go function for: cap_short @ symbol (throw: false) 
func E_Generate_cap_short_symbol (s EID) EID { 
    F_Generate_cap_short_symbol(ToSymbol(OBJ(s)) )
    return EVOID} 
  
// CLAIRE 4 NEW ! a class name is printed with the module identifier
// go_class is the the go name ModuleClass 
// class_ident => thing_ident is the name of the global variable that contains the CLAIRE object 
/* {1} The go function for: go_class(self:class) [status=0] */
func F_Generate_go_class_class (self *ClaireClass )  { 
    // procedure body with s = void 
{ var m *ClaireModule   = self.Name.Defined()
      if (m.Id() == C_Kernel.Id()) { 
        PRINC("Claire")
        }  else if (m.Id() != ToGenerateGoProducer(Optimize.C_PRODUCER.Value).Current.Id()) { 
        F_Generate_cap_ident_symbol(m.Name)
        PRINC(".")
        } 
      if ((self.Id() == C_array.Id()) || 
          (self.Id() == C_listargs.Id())) { 
        F_c_princ_string(MakeString("List"))
        } else {
        F_Generate_cap_ident_symbol(self.Name)
        } 
      } 
    } 
  
// The EID go function for: go_class @ class (throw: false) 
func E_Generate_go_class_class (self EID) EID { 
    F_Generate_go_class_class(ToClass(OBJ(self)) )
    return EVOID} 
  
// small subtlety : in cast names (ToX) we ommit the "Claire" for simplicity
/* {1} The go function for: cast_class(self:class) [status=0] */
func F_Generate_cast_class_class (self *ClaireClass )  { 
    // procedure body with s = void 
{ var m *ClaireModule   = self.Name.Defined()
      if ((m.Id() != C_Kernel.Id()) && 
          (m.Id() != ToGenerateGoProducer(Optimize.C_PRODUCER.Value).Current.Id())) { 
        F_Generate_cap_ident_symbol(m.Name)
        PRINC(".")
        } 
      PRINC("To")
      F_Generate_add_underscore_symbol(self.Name)
      if (self.Id() == C_listargs.Id()) { 
        PRINC("List")
        } else {
        F_Generate_cap_ident_symbol(self.Name)
        } 
      } 
    } 
  
// The EID go function for: cast_class @ class (throw: false) 
func E_Generate_cast_class_class (self EID) EID { 
    F_Generate_cast_class_class(ToClass(OBJ(self)) )
    return EVOID} 
  
// class_ident(c) = C_c
/* {1} The go function for: class_ident(self:class) [status=0] */
func F_Generate_class_ident_class (self *ClaireClass )  { 
    // procedure body with s = void 
F_Generate_symbol_ident_symbol(self.Name)
    } 
  
// The EID go function for: class_ident @ class (throw: false) 
func E_Generate_class_ident_class (self EID) EID { 
    F_Generate_class_ident_class(ToClass(OBJ(self)) )
    return EVOID} 
  
// same : remember that a class is  not a thing
/* {1} The go function for: thing_ident(self:thing) [status=0] */
func F_Generate_thing_ident_thing (self *ClaireThing )  { 
    // procedure body with s = void 
F_Generate_symbol_ident_symbol(self.Name)
    } 
  
// The EID go function for: thing_ident @ thing (throw: false) 
func E_Generate_thing_ident_thing (self EID) EID { 
    F_Generate_thing_ident_thing(ToThing(OBJ(self)) )
    return EVOID} 
  
// how a named object is designated in go (through a global variable from the package = module). 
// CLAIRE v4: No prefix needed for current or Kernel
/* {1} The go function for: symbol_ident(s:symbol) [status=0] */
func F_Generate_symbol_ident_symbol (s *ClaireSymbol )  { 
    // procedure body with s = void 
{ var m *ClaireModule   = s.Defined()
      if ((m.Id() != ToGenerateGoProducer(Optimize.C_PRODUCER.Value).Current.Id()) && 
          ((m.Id() != C_Kernel.Id()) && 
            (m.Id() != C_claire.Id()))) { 
        F_Generate_cap_short_symbol(m.Name)
        PRINC(".")
        } 
      F_Generate_go_var_symbol(s)
      } 
    } 
  
// The EID go function for: symbol_ident @ symbol (throw: false) 
func E_Generate_symbol_ident_symbol (s EID) EID { 
    F_Generate_symbol_ident_symbol(ToSymbol(OBJ(s)) )
    return EVOID} 
  
// this produced the C_s identifier which are go global variables, 
// all compiler code should use this (get rid of C_ in code)
/* {1} The go function for: go_var(s:symbol) [status=0] */
func F_Generate_go_var_symbol (s *ClaireSymbol )  { 
    // procedure body with s = void 
{ var m *ClaireModule   = s.Module_I()
      PRINC("C_")
      if (m.Id() != C_claire.Id()) { 
        F_c_princ_string(m.Name.String_I())
        F_c_princ_string(MakeString("_"))
        } 
      F_c_princ_string(s.String_I())
      } 
    } 
  
// The EID go function for: go_var @ symbol (throw: false) 
func E_Generate_go_var_symbol (s EID) EID { 
    F_Generate_go_var_symbol(ToSymbol(OBJ(s)) )
    return EVOID} 
  
// when we capitalize the name of class, we may create a conflict (list vs List)
/* {1} The go function for: add_underscore(name:symbol) [status=0] */
func F_Generate_add_underscore_symbol (name *ClaireSymbol )  { 
    // procedure body with s = void 
{ var s *ClaireString   = name.String_I()
      if ((F_integer_I_char(s.At(1)) >= 65) && 
          (F_integer_I_char(s.At(1)) <= 90)) { 
        PRINC("_")
        } 
      } 
    } 
  
// The EID go function for: add_underscore @ symbol (throw: false) 
func E_Generate_add_underscore_symbol (name EID) EID { 
    F_Generate_add_underscore_symbol(ToSymbol(OBJ(name)) )
    return EVOID} 
  
// the Go code producer uses Capitalization as a strategy for name generation
//  capitalize(s)  => capitalize the first letter + search for _, remove and capitalize next letter
//  capitalize("foo_bar") = "FooBar"
/* {1} The go function for: capitalize(s:string) [status=0] */
func F_Generate_capitalize_string (s *ClaireString ) *ClaireString  { 
    // procedure body with s = string 
var Result *ClaireString  
    { var n int  = F_length_string(s)
      _ = n
      { var i int  = F_get_string(s,'_')
        if (i == 0) { 
          { var s2 *ClaireString   = F_copy_string(s)
            F_nth_set_string(s2,1,F_Generate_capitalize_char(s.At(1)))
            Result = s2
            } 
          } else {
          Result = F_append_string(F_Generate_capitalize_string(F_substring_string(s,1,(i-1))),F_Generate_capitalize_string(F_substring_string(s,(i+1),n)))
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: capitalize @ string (throw: false) 
func E_Generate_capitalize_string (s EID) EID { 
    return EID{F_Generate_capitalize_string(ToString(OBJ(s)) ).Id(),0}} 
  
/* {1} The go function for: capitalize(c:char) [status=0] */
func F_Generate_capitalize_char (c rune) rune { 
    // procedure body with s = char 
var Result rune 
    { var i int  = int(c)
      if ((i >= 97) && 
          (i <= 122)) { 
        Result = F_char_I_integer((i-32))
        } else {
        Result = c
        } 
      } 
    return Result} 
  
// The EID go function for: capitalize @ char (throw: false) 
func E_Generate_capitalize_char (c EID) EID { 
    return EID{C__CHAR,CVAL(F_Generate_capitalize_char(CHAR(c) ))}} 
  
/* {1} The go function for: capitalize(s:symbol) [status=0] */
func F_Generate_capitalize_symbol (s *ClaireSymbol ) *ClaireString  { 
    return  F_Generate_capitalize_string(s.String_I())
    } 
  
// The EID go function for: capitalize @ symbol (throw: false) 
func E_Generate_capitalize_symbol (s EID) EID { 
    return EID{F_Generate_capitalize_symbol(ToSymbol(OBJ(s)) ).Id(),0}} 
  
// v3.3 : new ! a global variable contains the native value
// range = {} for global constant
/* {1} The go function for: globalVar(c:go_producer,x:global_variable) [status=0] */
func (c *GenerateGoProducer ) GlobalVar (x *Core.GlobalVariable )  { 
    // procedure body with s = void 
F_Generate_thing_ident_thing(ToThing(x.Id()))
    PRINC(".Value")
    } 
  
// The EID go function for: globalVar @ go_producer (throw: false) 
func E_Generate_globalVar_go_producer (c EID,x EID) EID { 
    ToGenerateGoProducer(OBJ(c)).GlobalVar(Core.ToGlobalVariable(OBJ(x)) )
    return EVOID} 
  
// the go expression that represents a global variable, as a string (reused for Gassign)
// Five sorts in go : categories to distinguish between native, object, EID
//    x:object       x,    x,       EID(x.Id(),0)
//    x:int,float,char         x,    MakeX(x),      EID{C__C,xVAL(x)}
//    x:exception     x,  x,      EID{x,1}
//  notice that Boolean is a an object but it could be handled with a native form in the future
/* {1} The go function for: type_sort(x:type) [status=0] */
func F_Generate_type_sort_type (x *ClaireType ) *ClaireClass  { 
    // procedure body with s = class 
var Result *ClaireClass  
    { var c *ClaireClass   = x.Class_I()
      if ((c.Id() == C_float.Id()) || 
          ((c.Id() == C_integer.Id()) || 
            ((c.Id() == C_char.Id()) || 
              (c.Id() == Optimize.C_EID.Id())))) { 
        Result = c
        } else {
        Result = C_any
        } 
      } 
    return Result} 
  
// The EID go function for: type_sort @ type (throw: false) 
func E_Generate_type_sort_type (x EID) EID { 
    return EID{F_Generate_type_sort_type(ToType(OBJ(x)) ).Id(),0}} 
  
// sorts in go are much simpler : int, float, any or EID
/* {1} The go function for: g_sort(x:any) [status=1] */
func F_Generate_g_sort_any (x *ClaireAny ) EID { 
    var Result EID 
    { var arg_1 *ClaireClass  
      _ = arg_1
      var try_2 EID 
      /*g_try(v2:"try_2",loop:false) */
      try_2 = Language.F_static_type_any(x)
      /* ERROR PROTECTION INSERTED (arg_1-Result) */
      if ErrorIn(try_2) {Result = try_2
      } else {
      arg_1 = ToClass(OBJ(try_2))
      Result = EID{F_Generate_type_sort_type(ToType(arg_1.Id())).Id(),0}
      }
      } 
    return Result} 
  
// The EID go function for: g_sort @ any (throw: true) 
func E_Generate_g_sort_any (x EID) EID { 
    return F_Generate_g_sort_any(ANY(x) )} 
  
// access the proper values slot for a list whose member type s is determined (not any)
/* {1} The go function for: valuesSlot(s:class) [status=0] */
func F_Generate_valuesSlot_class (s *ClaireClass )  { 
    // procedure body with s = void 
PRINC("Values")
    { var arg_1 *ClaireString  
      _ = arg_1
      if (s.Id() == C_integer.Id()) { 
        arg_1 = MakeString("I")
        }  else if (s.Id() == C_float.Id()) { 
        arg_1 = MakeString("F")
        } else {
        arg_1 = MakeString("O")
        } 
      F_princ_string(arg_1)
      } 
    PRINC("()")
    } 
  
// The EID go function for: valuesSlot @ class (throw: false) 
func E_Generate_valuesSlot_class (s EID) EID { 
    F_Generate_valuesSlot_class(ToClass(OBJ(s)) )
    return EVOID} 
  
// *******************************************************************
// *       Part 2: utilities for file generation                     *
// *******************************************************************
// note that all code to produce interfaces is gone :)
// a module is simply the combination of (1) a Go package (2) a ClaireModule (3) a load function
// generate a namespace definition (Go package)
/* {1} The go function for: namespace!(c:go_producer,m:module) [status=0] */
func (c *GenerateGoProducer ) Namespace_I (m *ClaireModule )  { 
    // procedure body with s = void 
PRINC("package ")
    F_iClaire_ident_symbol(m.Name)
    PRINC("\n")
    } 
  
// The EID go function for: namespace! @ go_producer (throw: false) 
func E_Generate_namespace_I_go_producer (c EID,m EID) EID { 
    ToGenerateGoProducer(OBJ(c)).Namespace_I(ToModule(OBJ(m)) )
    return EVOID} 
  
// note : we have removed module!(c:go_producer,m:module) => nothing to add to the previous line
// define a new typed variable named v (called in go_stat.cl )
// short cut : var declaration without an initialisation + a breakline
// CRAZY: go compiler gets confused with some variables not being used .. the dump forces to issue a dumb
// statement to get rid of this
// mode : 0 : normal no newline, 1 : newline, 2: special
/* {1} The go function for: var_declaration(v:string,s:class,mode:integer) [status=0] */
func F_Generate_var_declaration_string (v *ClaireString ,s *ClaireClass ,mode int)  { 
    // procedure body with s = void 
PRINC("var ")
    F_c_princ_string(v)
    PRINC(" ")
    F_Generate_interface_I_class(s)
    PRINC(" ")
    if (mode > 0) { 
      F_Generate_breakline_void()
      } 
    PRINC("")
    if (mode == 2) { 
      PRINC("_ = ")
      F_c_princ_string(v)
      F_Generate_breakline_void()
      PRINC("")
      } 
    } 
  
// The EID go function for: var_declaration @ string (throw: false) 
func E_Generate_var_declaration_string (v EID,s EID,mode EID) EID { 
    F_Generate_var_declaration_string(ToString(OBJ(v)),ToClass(OBJ(s)),INT(mode) )
    return EVOID} 
  
// ! is a semantic marker for imported
/* {1} The go function for: imported_function?(f:any) [status=0] */
func F_imported_function_ask_any (f *ClaireAny ) *ClaireBoolean  { 
    // procedure body with s = boolean 
var Result *ClaireBoolean  
    if (C_function.Id() == f.Isa.Id()) { 
      { var g0043 *ClaireFunction   = ToFunction(f)
        _ = g0043
        Result = Equal(MakeChar(F_string_I_function(g0043).At(1)).Id(),MakeChar('#').Id())
        } 
      } else {
      Result = CFALSE
      } 
    return Result} 
  
// The EID go function for: imported_function? @ any (throw: false) 
func E_imported_function_ask_any (f EID) EID { 
    return EID{F_imported_function_ask_any(ANY(f) ).Id(),0}} 
  
// tells if we can compile the CLAIRE method as a go method or if we shoud use a function
// we use the go method if the class is defined in the same 
// remember that Go does not support polymorphism on parameters : we can use a method only if there is one match 
// based on first argument - howver this restriction is package based (to be checked)
// we first check that the first char of the name is a proper letter
// also methods defined with #'#foo are forced to use foo :)
/* {1} The go function for: goMethod?(m:any) [status=0] */
func F_Generate_goMethod_ask_any (m *ClaireAny ) *ClaireBoolean  { 
    // procedure body with s = boolean 
var Result *ClaireBoolean  
    if (ToBoolean(ToGenerateGoProducer(Optimize.C_PRODUCER.Value).KernelMethods.Contain_ask(m).Id()) == CTRUE) { 
      Result = CTRUE
      }  else if (C_method.Id() == m.Isa.Id()) { 
      { var g0045 *ClaireMethod   = ToMethod(m)
        { var firstc rune  = g0045.Selector.Name.String_I().At(1)
          { var _Zsig *ClaireList   = F_Generate_go_signature_method(g0045)
            { var c *ClaireClass   = ToClass(_Zsig.ValuesO()[1-1])
              { 
                var v_and7 *ClaireBoolean  
                
                v_and7 = Core.F__inf_equal_char('A',firstc)
                if (v_and7 == CFALSE) {Result = CFALSE
                } else { 
                  v_and7 = Core.F__inf_equal_char(firstc,'z')
                  if (v_and7 == CFALSE) {Result = CFALSE
                  } else { 
                    v_and7 = Core.F__I_equal_any(MakeChar(firstc).Id(),MakeChar('^').Id())
                    if (v_and7 == CFALSE) {Result = CFALSE
                    } else { 
                      v_and7 = MakeBoolean((ToType(c.Id()).Included(ToType(C_object.Id())) == CTRUE) || (c.Id() == C_port.Id()) || (c.Id() == C_environment.Id()))
                      if (v_and7 == CFALSE) {Result = CFALSE
                      } else { 
                        v_and7 = Equal(g0045.Module_I.Id(),c.Name.Defined().Id())
                        if (v_and7 == CFALSE) {Result = CFALSE
                        } else { 
                          v_and7 = MakeBoolean((g0045.Selector.IfWrite == CNULL))
                          if (v_and7 == CFALSE) {Result = CFALSE
                          } else { 
                            v_and7 = MakeBoolean((g0045.Functional.Id() == CNULL) || (F_imported_function_ask_any(g0045.Functional.Id()) != CTRUE))
                            if (v_and7 == CFALSE) {Result = CFALSE
                            } else { 
                              { var arg_1 *ClaireAny  
                                _ = arg_1
                                { 
                                  var m2 *ClaireRestriction  
                                  _ = m2
                                  var m2_iter *ClaireAny  
                                  arg_1= CFALSE.Id()
                                  for _,m2_iter = range(g0045.Selector.Restrictions.ValuesO()){ 
                                    m2 = ToRestriction(m2_iter)
                                    var g0049I *ClaireBoolean  
                                    { var arg_2 *ClaireBoolean  
                                      _ = arg_2
                                      if (C_method.Id() == m2.Isa.Id()) { 
                                        { var g0046 *ClaireMethod   = ToMethod(m2.Id())
                                          if (C_class.Id() != g0046.Domain.ValuesO()[1-1].Isa.Id()) { 
                                            arg_2 = CFALSE
                                            }  else if ((g0046.Module_I.Id() == g0045.Module_I.Id()) && 
                                              (Equal(Core.F__exp_type(ToType(c.Id()),ToType(g0046.Domain.ValuesO()[1-1])).Id(),CEMPTY.Id()) != CTRUE)) { 
                                            arg_2 = F_Generate_arg_match_list(F_Generate_go_signature_method(g0046),_Zsig)
                                            } else {
                                            arg_2 = CTRUE
                                            } 
                                          } 
                                        } else {
                                        arg_2 = CTRUE
                                        } 
                                      g0049I = arg_2.Not
                                      } 
                                    if (g0049I == CTRUE) { 
                                      arg_1 = CTRUE.Id()
                                      break
                                      } 
                                    } 
                                  } 
                                v_and7 = Core.F_not_any(arg_1)
                                } 
                              if (v_and7 == CFALSE) {Result = CFALSE
                              } else { 
                                Result = CTRUE} 
                              } 
                            } 
                          } 
                        } 
                      } 
                    } 
                  } 
                } 
              } 
            } 
          } 
        } 
      } else {
      Result = CFALSE
      } 
    return Result} 
  
// The EID go function for: goMethod? @ any (throw: false) 
func E_Generate_goMethod_ask_any (m EID) EID { 
    return EID{F_Generate_goMethod_ask_any(ANY(m) ).Id(),0}} 
  
// useful for debug - notice that a CLAIRE method defined on a class which is NOT in the same module
// is always compiled as a function
/* {1} The go function for: dMethod?(m:any) [status=1] */
func F_dMethod_ask_any (m *ClaireAny ) EID { 
    var Result EID 
    { var firstc rune 
      var try_1 EID 
      /*g_try(v2:"try_1",loop:false) */
      { var arg_2 *ClaireAny  
        _ = arg_2
        var try_3 EID 
        /*g_try(v2:"try_3",loop:false) */
        try_3 = Core.F_CALL(C_string_I,ARGS(Core.F_CALL(C_name,ARGS(Core.F_CALL(C_selector,ARGS(m.ToEID()))))))
        /* ERROR PROTECTION INSERTED (arg_2-try_1) */
        if ErrorIn(try_3) {try_1 = try_3
        } else {
        arg_2 = ANY(try_3)
        try_1 = EID{C__CHAR,CVAL(ToString(arg_2).At(1))}
        }
        } 
      /* ERROR PROTECTION INSERTED (firstc-Result) */
      if ErrorIn(try_1) {Result = try_1
      } else {
      firstc = CHAR(try_1)
      { var _Zsig *ClaireList   = F_Generate_go_signature_method(ToMethod(m))
        { var c *ClaireClass   = ToClass(_Zsig.ValuesO()[1-1])
          /*g_try(v2:"Result",loop:true) */
          PRINC("char -> ")
          /*g_try(v2:"Result",loop:true) */
          Result = Core.F_print_any(MakeBoolean((Core.F__inf_equal_char('A',firstc) == CTRUE) && (Core.F__inf_equal_char(firstc,'z') == CTRUE) && (firstc != '^')).Id())
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC("\n")
          Result = EVOID
          }
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          /*g_try(v2:"Result",loop:true) */
          PRINC("hierarchy -> ")
          /*g_try(v2:"Result",loop:true) */
          Result = Core.F_print_any(MakeBoolean((ToType(c.Id()).Included(ToType(C_object.Id())) == CTRUE) || (c.Id() == C_port.Id()) || (c.Id() == C_environment.Id())).Id())
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC("\n")
          Result = EVOID
          }
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          /*g_try(v2:"Result",loop:true) */
          PRINC("module [")
          /*g_try(v2:"Result",loop:true) */
          Result = Core.F_print_any(ANY(Core.F_CALL(C_module_I,ARGS(m.ToEID()))))
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC("] -> ")
          /*g_try(v2:"Result",loop:true) */
          Result = Core.F_print_any(Equal(ANY(Core.F_CALL(C_module_I,ARGS(m.ToEID()))),c.Name.Defined().Id()).Id())
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC("\n")
          Result = EVOID
          }}
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC("all m -> ")
          /*g_try(v2:"Result",loop:true) */
          { var arg_4 *ClaireBoolean  
            _ = arg_4
            var try_5 EID 
            /*g_try(v2:"try_5",loop:false) */
            { var arg_6 *ClaireAny  
              _ = arg_6
              var try_7 EID 
              /*g_try(v2:"try_7",loop:false) */
              { 
                var m2 *ClaireRestriction  
                _ = m2
                var m2_iter *ClaireAny  
                try_7= EID{CFALSE.Id(),0}
                var m2_support *ClaireList  
                m2_support = ToProperty(OBJ(Core.F_CALL(C_selector,ARGS(m.ToEID())))).Restrictions
                for _,m2_iter = range(m2_support.ValuesO()){ 
                  m2 = ToRestriction(m2_iter)
                  var loop_8 EID 
                  _ = loop_8
                  /*g_try(v2:"loop_8",loop:tuple("try_7", EID)) */
                  var g0052I *ClaireBoolean  
                  var try_9 EID 
                  /*g_try(v2:"try_9",loop:false) */
                  { var arg_10 *ClaireBoolean  
                    _ = arg_10
                    var try_11 EID 
                    /*g_try(v2:"try_11",loop:false) */
                    /*g_try(v2:"try_11",loop:false) */
                    PRINC("---- try m2 = ")
                    /*g_try(v2:"try_11",loop:false) */
                    try_11 = Core.F_print_any(m2.Id())
                    /* ERROR PROTECTION INSERTED (try_11-try_11) */
                    if !ErrorIn(try_11) {
                    PRINC(" in ")
                    /*g_try(v2:"try_11",loop:false) */
                    try_11 = Core.F_print_any(m2.Module_I.Id())
                    /* ERROR PROTECTION INSERTED (try_11-try_11) */
                    if !ErrorIn(try_11) {
                    PRINC("\n")
                    try_11 = EVOID
                    }}
                    /* ERROR PROTECTION INSERTED (try_11-try_11) */
                    if !ErrorIn(try_11) {
                    if (C_method.Id() == m2.Isa.Id()) { 
                      { var g0050 *ClaireMethod   = ToMethod(m2.Id())
                        if ((g0050.Module_I.Id() == ANY(Core.F_CALL(C_module_I,ARGS(m.ToEID())))) && 
                            (Equal(Core.F__exp_type(ToType(c.Id()),ToType(g0050.Domain.ValuesO()[1-1])).Id(),CEMPTY.Id()) != CTRUE)) { 
                          try_11 = EID{F_Generate_arg_match_list(F_Generate_go_signature_method(g0050),_Zsig).Id(),0}
                          } else {
                          try_11 = EID{CTRUE.Id(),0}
                          } 
                        } 
                      } else {
                      try_11 = EID{CTRUE.Id(),0}
                      } 
                    }
                    /* ERROR PROTECTION INSERTED (arg_10-try_9) */
                    if ErrorIn(try_11) {try_9 = try_11
                    } else {
                    arg_10 = ToBoolean(OBJ(try_11))
                    try_9 = EID{arg_10.Not.Id(),0}
                    }
                    } 
                  /* ERROR PROTECTION INSERTED (g0052I-loop_8) */
                  if ErrorIn(try_9) {loop_8 = try_9
                  } else {
                  g0052I = ToBoolean(OBJ(try_9))
                  if (g0052I == CTRUE) { 
                    try_7 = EID{CTRUE.Id(),0}
                    break
                    } else {
                    loop_8 = EID{CFALSE.Id(),0}
                    } 
                  }
                  /* ERROR PROTECTION INSERTED (loop_8-try_7) */
                  if ErrorIn(loop_8) {try_7 = loop_8
                  break
                  } else {
                  }
                  } 
                } 
              /* ERROR PROTECTION INSERTED (arg_6-try_5) */
              if ErrorIn(try_7) {try_5 = try_7
              } else {
              arg_6 = ANY(try_7)
              try_5 = EID{Core.F_not_any(arg_6).Id(),0}
              }
              } 
            /* ERROR PROTECTION INSERTED (arg_4-Result) */
            if ErrorIn(try_5) {Result = try_5
            } else {
            arg_4 = ToBoolean(OBJ(try_5))
            Result = Core.F_print_any(arg_4.Id())
            }
            } 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC("\n")
          Result = EVOID
          }
          }}}
          } 
        } 
      }
      } 
    return Result} 
  
// The EID go function for: dMethod? @ any (throw: true) 
func E_dMethod_ask_any (m EID) EID { 
    return F_dMethod_ask_any(ANY(m) )} 
  
// same argument types for all restrictions, excluding the range (that is included in go_signature)
/* {1} The go function for: arg_match(l1:list<class>,l2:list<class>) [status=0] */
func F_Generate_arg_match_list (l1 *ClaireList ,l2 *ClaireList ) *ClaireBoolean  { 
    // procedure body with s = boolean 
var Result *ClaireBoolean  
    { var n int  = l1.Length()
      { 
        var v_and3 *ClaireBoolean  
        
        v_and3 = Equal(MakeInteger(l2.Length()).Id(),MakeInteger(n).Id())
        if (v_and3 == CFALSE) {Result = CFALSE
        } else { 
          { var arg_1 *ClaireAny  
            _ = arg_1
            { var i int  = 2
              { var g0053 int  = (n-1)
                _ = g0053
                arg_1= CFALSE.Id()
                for (i <= g0053) { 
                  /* While stat, v:"arg_1" loop:false */
                  if (l1.ValuesO()[i-1] != l2.ValuesO()[i-1]) { 
                    arg_1 = CTRUE.Id()
                    break
                    } 
                  i = (i+1)
                  /* try?:false, v2:"v_while8" loop will be:tuple("arg_1", any) */
                  } 
                } 
              } 
            v_and3 = Core.F_not_any(arg_1)
            } 
          if (v_and3 == CFALSE) {Result = CFALSE
          } else { 
            Result = CTRUE} 
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: arg_match @ list (throw: false) 
func E_Generate_arg_match_list (l1 EID,l2 EID) EID { 
    return EID{F_Generate_arg_match_list(ToList(OBJ(l1)),ToList(OBJ(l2)) ).Id(),0}} 
  
// create the function (a name) for a method with selector p and signature l
// the name of the module where p was defined is included (until claire => public)
/* {1} The go function for: Compile/function_name(p:property,l:list) [status=0] */
func F_Compile_function_name_property2 (p *ClaireProperty ,l *ClaireList ) *ClaireString  { 
    // procedure body with s = string 
var Result *ClaireString  
    { var n int  = 0
      _ = n
      { var m int  = 0
        _ = m
        { var md *ClaireModule   = p.Name.Module_I()
          _ = md
          { var c *ClaireClass   = ToTypeExpression(l.At(1-1)).Class_I()
            { var r *ClaireString   = F_append_string(F_append_string(p.Name.String_I(),MakeString("_")),c.Name.String_I())
              if ((p.Id() != Core.C_main.Id()) && 
                  (md.Id() != C_claire.Id())) { 
                r = F_append_string(F_append_string(md.Name.String_I(),MakeString("_")),r)
                } 
              { 
                var r *ClaireRestriction  
                _ = r
                var r_iter *ClaireAny  
                for _,r_iter = range(p.Restrictions.ValuesO()){ 
                  r = ToRestriction(r_iter)
                  if (c.Id() == Core.F_domain_I_restriction(r).Id()) { 
                    n = (n+1)
                    } 
                  if (Optimize.F_Optimize__equalsig_ask_list(l,r.Domain) == CTRUE) { 
                    m = n
                    } 
                  } 
                } 
              if (n <= 1) { 
                Result = r
                } else {
                Result = F_append_string(r,F_string_I_integer(m))
                } 
              } 
            } 
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: Compile/function_name @ list<type_expression>(property, list) (throw: false) 
func E_Compile_function_name_property2 (p EID,l EID) EID { 
    return EID{F_Compile_function_name_property2(ToProperty(OBJ(p)),ToList(OBJ(l)) ).Id(),0}} 
  
/* {1} The go function for: at(p:go_producer) [status=0] */
func (p *GenerateGoProducer ) At ()  { 
    // procedure body with s = void 
PRINC(".")
    } 
  
// The EID go function for: at @ go_producer (throw: false) 
func E_Generate_at_go_producer (p EID) EID { 
    ToGenerateGoProducer(OBJ(p)).At( )
    return EVOID} 
  
// prints a list of arguments with types / replaces typed_args_list
/* {1} The go function for: goVariables(p:go_producer,self:list) [status=0] */
func (p *GenerateGoProducer ) GoVariables (self *ClaireList ) *ClaireAny  { 
    // procedure body with s = any 
var Result *ClaireAny  
    { var prems *ClaireBoolean   = CTRUE
      _ = prems
      { 
        var x *ClaireVariable  
        _ = x
        var x_iter *ClaireAny  
        Result= CFALSE.Id()
        var x_support *ClaireList  
        x_support = self
        x_len := x_support.Length()
        for i_it := 0; i_it < x_len; i_it++ { 
          x_iter = x_support.At(i_it)
          x = To_Variable(x_iter)
          if (prems == CTRUE) { 
            prems = CFALSE
            } else {
            PRINC(",")
            } 
          p.GoVariable(x)
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: goVariables @ go_producer (throw: false) 
func E_Generate_goVariables_go_producer (p EID,self EID) EID { 
    return ToGenerateGoProducer(OBJ(p)).GoVariables(ToList(OBJ(self)) ).ToEID()} 
  
// prints a variable declaration (inside an arg list
/* {1} The go function for: goVariable(p:go_producer,v:Variable) [status=0] */
func (p *GenerateGoProducer ) GoVariable (v *ClaireVariable )  { 
    // procedure body with s = void 
F_iClaire_ident_go_producer1(p,v)
    PRINC(" ")
    F_Generate_interface_I_class(v.Range.Class_I())
    PRINC("")
    } 
  
// The EID go function for: goVariable @ go_producer (throw: false) 
func E_Generate_goVariable_go_producer (p EID,v EID) EID { 
    ToGenerateGoProducer(OBJ(p)).GoVariable(To_Variable(OBJ(v)) )
    return EVOID} 
  
// prints the name of a method as a go method 
// Here we use the list of exceptions (kernel_methods) to force a "go method syntax" (with possibly a forced name)
// this is convenient when cross-compiling (when method move from one module/package to another)
/* {1} The go function for: goMethod(m:method) [status=0] */
func F_Generate_goMethod_method (m *ClaireMethod )  { 
    // procedure body with s = void 
{ var lm *ClaireList   = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).KernelMethods
      { var i int  = F_index_list(lm,m.Id())
        if ((i > 0) && 
            ((lm.Length() > i) && 
              (C_string.Id() == lm.At((i+1)-1).Isa.Id()))) { 
          Core.F_CALL(C_c_princ,ARGS(lm.At((i+1)-1).ToEID()))
          } else {
          F_c_princ_string(F_Generate_capitalize_symbol(m.Selector.Name))
          } 
        } 
      } 
    } 
  
// The EID go function for: goMethod @ method (throw: false) 
func E_Generate_goMethod_method (m EID) EID { 
    F_Generate_goMethod_method(ToMethod(OBJ(m)) )
    return EVOID} 
  
// prints the name of a function as a go function F_f
// NOTE : the link method <=> go function is not stored (the function is not known by CLAIRE)
// imported functions do not refer to the module/package
/* {1} The go function for: goFunction(m:method) [status=1] */
func F_Generate_goFunction_method (m *ClaireMethod ) EID { 
    var Result EID 
    { var md *ClaireModule   = m.Module_I
      if ((md.Id() != C_Kernel.Id()) && 
          ((md.Id() != ToGenerateGoProducer(Optimize.C_PRODUCER.Value).Current.Id()) && 
            ((md.Id() != C_claire.Id()) && 
              ((m.Functional.Id() == CNULL) || 
                  (F_imported_function_ask_any(m.Functional.Id()) != CTRUE))))) { 
        F_Generate_cap_short_symbol(md.Name)
        PRINC(".")
        } 
      PRINC("F_")
      { var arg_1 *ClaireString  
        _ = arg_1
        var try_2 EID 
        /*g_try(v2:"try_2",loop:false) */
        try_2 = F_Generate_getFunctionName_method(m)
        /* ERROR PROTECTION INSERTED (arg_1-Result) */
        if ErrorIn(try_2) {Result = try_2
        } else {
        arg_1 = ToString(OBJ(try_2))
        F_Generate_import_princ_string(arg_1)
        Result = EVOID
        }
        } 
      } 
    return Result} 
  
// The EID go function for: goFunction @ method (throw: true) 
func E_Generate_goFunction_method (m EID) EID { 
    return F_Generate_goFunction_method(ToMethod(OBJ(m)) )} 
  
// specialized version for Core method
/* {1} The go function for: preCore?(_CL_obj:void) [status=0] */
func F_Generate_preCore_ask_void ()  { 
    // procedure body with s = void 
if (ToGenerateGoProducer(Optimize.C_PRODUCER.Value).Current.Id() != Core.It.Id()) { 
      PRINC("Core.")
      } 
    } 
  
// The EID go function for: preCore? @ void (throw: false) 
func E_Generate_preCore_ask_void (_CL_obj EID) EID { 
    F_Generate_preCore_ask_void( )
    return EVOID} 
  
// prints the name of the EID method that is linked by the compiler to the method
/* {1} The go function for: goEIDFunction(m:method) [status=1] */
func F_Generate_goEIDFunction_method (m *ClaireMethod ) EID { 
    var Result EID 
    { var s *ClaireString  
      var try_1 EID 
      /*g_try(v2:"try_1",loop:false) */
      try_1 = F_Generate_getFunctionName_method(m)
      /* ERROR PROTECTION INSERTED (s-Result) */
      if ErrorIn(try_1) {Result = try_1
      } else {
      s = ToString(OBJ(try_1))
      PRINC("MakeFunction")
      F_princ_integer(m.Domain.Length())
      PRINC("(E_")
      F_Generate_import_princ_string(s)
      PRINC(",")
      /*g_try(v2:"Result",loop:true) */
      Result = Core.F_print_any((s).Id())
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(")")
      Result = EVOID
      }
      }
      } 
    return Result} 
  
// The EID go function for: goEIDFunction @ method (throw: true) 
func E_Generate_goEIDFunction_method (m EID) EID { 
    return F_Generate_goEIDFunction_method(ToMethod(OBJ(m)) )} 
  
// prints the function MakeFunction(...) expression
/* {1} The go function for: goEIDFunctionName(m:method) [status=1] */
func F_Generate_goEIDFunctionName_method (m *ClaireMethod ) EID { 
    var Result EID 
    { var f *ClaireString  
      _ = f
      var try_1 EID 
      /*g_try(v2:"try_1",loop:false) */
      try_1 = F_Generate_getFunctionName_method(m)
      /* ERROR PROTECTION INSERTED (f-Result) */
      if ErrorIn(try_1) {Result = try_1
      } else {
      f = ToString(OBJ(try_1))
      PRINC("E_")
      F_c_princ_string(f)
      PRINC("")
      Result = EVOID
      }
      } 
    return Result} 
  
// The EID go function for: goEIDFunctionName @ method (throw: true) 
func E_Generate_goEIDFunctionName_method (m EID) EID { 
    return F_Generate_goEIDFunctionName_method(ToMethod(OBJ(m)) )} 
  
// special function for self_eval of type  => added as an extra paramer of type eFunc
// cf. goexp: AddMethod -> AddEvalMethod     
/* {1} The go function for: goEvalFunction(m:method) [status=0] */
func F_Generate_goEvalFunction_method (m *ClaireMethod )  { 
    // procedure body with s = void 
{ var c *ClaireClass   = Core.F_domain_I_restriction(ToRestriction(m.Id()))
      _ = c
      PRINC(",EVAL_")
      c.Name.CPrinc()
      PRINC("")
      } 
    } 
  
// The EID go function for: goEvalFunction @ method (throw: false) 
func E_Generate_goEvalFunction_method (m EID) EID { 
    F_Generate_goEvalFunction_method(ToMethod(OBJ(m)) )
    return EVOID} 
  
// exceptions
// get function name
/* {1} The go function for: getFunctionName(m:method) [status=1] */
func F_Generate_getFunctionName_method (m *ClaireMethod ) EID { 
    var Result EID 
    if (m.Id() == C_Generate__starlength_string_star.Value) { 
      Result = EID{MakeString("length_string").Id(),0}
      }  else if (m.Id() == C_Generate__starnth_list_star.Value) { 
      Result = EID{MakeString("nth_list").Id(),0}
      }  else if (m.Id() == C_Generate__starset_I_list_star.Value) { 
      Result = EID{MakeString("set_I_list").Id(),0}
      }  else if (m.Id() == C_Generate__starstack_apply_star.Value) { 
      Result = EID{MakeString("CALL").Id(),0}
      }  else if (m.Id() == C_Generate__starsuper_apply_star.Value) { 
      Result = EID{MakeString("SUPER").Id(),0}
      }  else if (m.Id() == C_Generate__starbelong_star.Value) { 
      Result = EID{MakeString("BELONG").Id(),0}
      }  else if (m.Functional.Id() != CNULL) { 
      Result = EID{F_string_I_function(m.Functional).Id(),0}
      } else {
      Result = Core.F_CALL(Optimize.C_Compile_function_name,ARGS(EID{m.Selector.Id(),0},EID{m.Domain.Id(),0}))
      } 
    return Result} 
  
// The EID go function for: getFunctionName @ method (throw: true) 
func E_Generate_getFunctionName_method (m EID) EID { 
    return F_Generate_getFunctionName_method(ToMethod(OBJ(m)) )} 
  
// ugly : reverse engineer a compiled definition into a method
// we need to do something better
/* {1} The go function for: retreive_method(p:any,lf:any) [status=1] */
func F_Generate_retreive_method_any (p *ClaireAny ,lf *ClaireAny ) EID { 
    var Result EID 
    if (p.Isa.IsIn(C_property) == CTRUE) { 
      { var g0054 *ClaireProperty   = ToProperty(p)
        { var m *ClaireObject  
          var try_1 EID 
          /*g_try(v2:"try_1",loop:false) */
          { var arg_2 *ClaireAny  
            _ = arg_2
            var try_3 EID 
            /*g_try(v2:"try_3",loop:false) */
            try_3 = F_Generate_retreive_list_any(lf)
            /* ERROR PROTECTION INSERTED (arg_2-try_1) */
            if ErrorIn(try_3) {try_1 = try_3
            } else {
            arg_2 = ANY(try_3)
            try_1 = Core.F_CALL(ToProperty(Core.C__at.Id()),ARGS(EID{g0054.Id(),0},arg_2.ToEID()))
            }
            } 
          /* ERROR PROTECTION INSERTED (m-Result) */
          if ErrorIn(try_1) {Result = try_1
          } else {
          m = ToObject(OBJ(try_1))
          if (C_method.Id() == m.Isa.Id()) { 
            { var g0055 *ClaireMethod   = ToMethod(m.Id())
              _ = g0055
              Result = EID{g0055.Id(),0}
              } 
            } else {
            Result = ToException(Core.C_general_error.Make(MakeString("there is no method ~S @ ~S").Id(),MakeConstantList(g0054.Id(),lf).Id())).Close()
            } 
          }
          } 
        } 
      } else {
      Result = ToException(Core.C_general_error.Make(MakeString("we have a problem to retreive ~S (not a property) at ~S").Id(),MakeConstantList(p,lf).Id())).Close()
      } 
    return Result} 
  
// The EID go function for: retreive_method @ any (throw: true) 
func E_Generate_retreive_method_any (p EID,lf EID) EID { 
    return F_Generate_retreive_method_any(ANY(p),ANY(lf) )} 
  
// constrained eval in disguise : returns a type or a list of types from CLAIRE expressions
/* {1} The go function for: retreive_list(x:any) [status=1] */
func F_Generate_retreive_list_any (x *ClaireAny ) EID { 
    var Result EID 
    if (x.Isa.IsIn(C_type) == CTRUE) { 
      { var g0058 *ClaireType   = ToType(x)
        _ = g0058
        Result = EID{g0058.Id(),0}
        } 
      }  else if (C_integer.Id() == x.Isa.Id()) { 
      { var g0059 int  = ToInteger(x).Value
        _ = g0059
        Result = EID{C__INT,IVAL(g0059)}
        } 
      }  else if (x.Isa.IsIn(C_property) == CTRUE) { 
      { var g0060 *ClaireProperty   = ToProperty(x)
        _ = g0060
        Result = EID{g0060.Id(),0}
        } 
      }  else if (x.Isa.IsIn(Core.C_global_variable) == CTRUE) { 
      { var g0061 *Core.GlobalVariable   = Core.ToGlobalVariable(x)
        if (Equal(g0061.Range.Id(),CEMPTY.Id()) == CTRUE) { 
          Result = F_Generate_retreive_list_any(g0061.Value)
          } else {
          Result = ToException(Core.C_general_error.Make(MakeString("we cannot retreive a type from a variable ~S").Id(),MakeConstantList(g0061.Id()).Id())).Close()
          } 
        } 
      }  else if (x.Isa.IsIn(Language.C_List) == CTRUE) { 
      { var g0062 *Language.List   = Language.To_List(x)
        _ = g0062
        { 
          var v_list4 *ClaireList  
          var y *ClaireAny  
          var v_local4 *ClaireAny  
          v_list4 = g0062.Args
          Result = EID{CreateList(ToType(CEMPTY.Id()),v_list4.Length()).Id(),0}
          for CLcount := 0; CLcount < v_list4.Length(); CLcount++{ 
            y = v_list4.At(CLcount)
            var try_1 EID 
            /*g_try(v2:"try_1",loop:tuple("Result", EID)) */
            try_1 = F_Generate_retreive_list_any(y)
            /* ERROR PROTECTION INSERTED (v_local4-Result) */
            if ErrorIn(try_1) {Result = try_1
            break
            } else {
            v_local4 = ANY(try_1)
            ToList(OBJ(Result)).PutAt(CLcount,v_local4)
            } 
          }
          } 
        } 
      }  else if (x.Isa.IsIn(Language.C_Tuple) == CTRUE) { 
      { var g0063 *Language.Tuple   = Language.To_Tuple(x)
        _ = g0063
        { var arg_2 *ClaireList  
          _ = arg_2
          var try_3 EID 
          /*g_try(v2:"try_3",loop:false) */
          { 
            var v_list5 *ClaireList  
            var y *ClaireAny  
            var v_local5 *ClaireAny  
            v_list5 = g0063.Args
            try_3 = EID{CreateList(ToType(CEMPTY.Id()),v_list5.Length()).Id(),0}
            for CLcount := 0; CLcount < v_list5.Length(); CLcount++{ 
              y = v_list5.At(CLcount)
              var try_4 EID 
              /*g_try(v2:"try_4",loop:tuple("try_3", EID)) */
              try_4 = F_Generate_retreive_list_any(y)
              /* ERROR PROTECTION INSERTED (v_local5-try_3) */
              if ErrorIn(try_4) {try_3 = try_4
              break
              } else {
              v_local5 = ANY(try_4)
              ToList(OBJ(try_3)).PutAt(CLcount,v_local5)
              } 
            }
            } 
          /* ERROR PROTECTION INSERTED (arg_2-Result) */
          if ErrorIn(try_3) {Result = try_3
          } else {
          arg_2 = ToList(OBJ(try_3))
          Result = EID{arg_2.Tuple_I().Id(),0}
          }
          } 
        } 
      }  else if (x.Isa.IsIn(Language.C_Call_method) == CTRUE) { 
      { var g0064 *Language.CallMethod   = Language.To_CallMethod(x)
        if ((g0064.Arg.Selector.Id() == C_nth.Id()) && 
            (g0064.Args.Length() == 2)) { 
          { var arg_5 *ClaireAny  
            _ = arg_5
            var try_7 EID 
            /*g_try(v2:"try_7",loop:false) */
            try_7 = F_Generate_retreive_list_any(g0064.Args.At(1-1))
            /* ERROR PROTECTION INSERTED (arg_5-Result) */
            if ErrorIn(try_7) {Result = try_7
            } else {
            arg_5 = ANY(try_7)
            { var arg_6 *ClaireAny  
              _ = arg_6
              var try_8 EID 
              /*g_try(v2:"try_8",loop:false) */
              try_8 = F_Generate_retreive_list_any(g0064.Args.At(2-1))
              /* ERROR PROTECTION INSERTED (arg_6-Result) */
              if ErrorIn(try_8) {Result = try_8
              } else {
              arg_6 = ANY(try_8)
              Result = Core.F_CALL(C_nth,ARGS(arg_5.ToEID(),arg_6.ToEID()))
              }
              } 
            }
            } 
          }  else if ((g0064.Arg.Selector.Id() == C_nth.Id()) && 
            (g0064.Args.Length() == 3)) { 
          { var arg_9 *ClaireAny  
            _ = arg_9
            var try_12 EID 
            /*g_try(v2:"try_12",loop:false) */
            try_12 = F_Generate_retreive_list_any(g0064.Args.At(1-1))
            /* ERROR PROTECTION INSERTED (arg_9-Result) */
            if ErrorIn(try_12) {Result = try_12
            } else {
            arg_9 = ANY(try_12)
            { var arg_10 *ClaireAny  
              _ = arg_10
              var try_13 EID 
              /*g_try(v2:"try_13",loop:false) */
              try_13 = F_Generate_retreive_list_any(g0064.Args.At(2-1))
              /* ERROR PROTECTION INSERTED (arg_10-Result) */
              if ErrorIn(try_13) {Result = try_13
              } else {
              arg_10 = ANY(try_13)
              { var arg_11 *ClaireAny  
                _ = arg_11
                var try_14 EID 
                /*g_try(v2:"try_14",loop:false) */
                try_14 = F_Generate_retreive_list_any(g0064.Args.At(3-1))
                /* ERROR PROTECTION INSERTED (arg_11-Result) */
                if ErrorIn(try_14) {Result = try_14
                } else {
                arg_11 = ANY(try_14)
                Result = Core.F_CALL(C_nth,ARGS(arg_9.ToEID(),arg_10.ToEID(),arg_11.ToEID()))
                }
                } 
              }
              } 
            }
            } 
          }  else if ((g0064.Arg.Selector.Id() == Core.C_Core_param_I.Id()) && 
            (g0064.Args.Length() == 2)) { 
          { var arg_15 *ClaireAny  
            _ = arg_15
            var try_17 EID 
            /*g_try(v2:"try_17",loop:false) */
            try_17 = F_Generate_retreive_list_any(g0064.Args.At(1-1))
            /* ERROR PROTECTION INSERTED (arg_15-Result) */
            if ErrorIn(try_17) {Result = try_17
            } else {
            arg_15 = ANY(try_17)
            { var arg_16 *ClaireAny  
              _ = arg_16
              var try_18 EID 
              /*g_try(v2:"try_18",loop:false) */
              try_18 = F_Generate_retreive_list_any(g0064.Args.At(2-1))
              /* ERROR PROTECTION INSERTED (arg_16-Result) */
              if ErrorIn(try_18) {Result = try_18
              } else {
              arg_16 = ANY(try_18)
              Result = EID{Core.F_param_I_class(ToClass(arg_15),ToType(arg_16)).Id(),0}
              }
              } 
            }
            } 
          }  else if ((g0064.Arg.Selector.Id() == Core.C_U.Id()) && 
            (g0064.Args.Length() == 2)) { 
          { var arg_19 *ClaireAny  
            _ = arg_19
            var try_21 EID 
            /*g_try(v2:"try_21",loop:false) */
            try_21 = F_Generate_retreive_list_any(g0064.Args.At(1-1))
            /* ERROR PROTECTION INSERTED (arg_19-Result) */
            if ErrorIn(try_21) {Result = try_21
            } else {
            arg_19 = ANY(try_21)
            { var arg_20 *ClaireAny  
              _ = arg_20
              var try_22 EID 
              /*g_try(v2:"try_22",loop:false) */
              try_22 = F_Generate_retreive_list_any(g0064.Args.At(2-1))
              /* ERROR PROTECTION INSERTED (arg_20-Result) */
              if ErrorIn(try_22) {Result = try_22
              } else {
              arg_20 = ANY(try_22)
              Result = EID{Core.F_U_type(ToType(arg_19),ToType(arg_20)).Id(),0}
              }
              } 
            }
            } 
          }  else if ((g0064.Arg.Selector.Id() == C__dot_dot.Id()) && 
            (g0064.Args.Length() == 2)) { 
          { var arg_23 *ClaireAny  
            _ = arg_23
            var try_25 EID 
            /*g_try(v2:"try_25",loop:false) */
            try_25 = F_Generate_retreive_list_any(g0064.Args.At(1-1))
            /* ERROR PROTECTION INSERTED (arg_23-Result) */
            if ErrorIn(try_25) {Result = try_25
            } else {
            arg_23 = ANY(try_25)
            { var arg_24 *ClaireAny  
              _ = arg_24
              var try_26 EID 
              /*g_try(v2:"try_26",loop:false) */
              try_26 = F_Generate_retreive_list_any(g0064.Args.At(2-1))
              /* ERROR PROTECTION INSERTED (arg_24-Result) */
              if ErrorIn(try_26) {Result = try_26
              } else {
              arg_24 = ANY(try_26)
              Result = EID{Core.F__dot_dot_integer(ToInteger(arg_23).Value,ToInteger(arg_24).Value).Id(),0}
              }
              } 
            }
            } 
          } else {
          Result = ToException(Core.C_general_error.Make(MakeString("we need to extend retreive_list to handle a type call: ~S").Id(),MakeConstantList(g0064.Id()).Id())).Close()
          } 
        } 
      } else {
      Result = ToException(Core.C_general_error.Make(MakeString("we need to extend retreive_list to handle ~S").Id(),MakeConstantList(x).Id())).Close()
      } 
    return Result} 
  
// The EID go function for: retreive_list @ any (throw: true) 
func E_Generate_retreive_list_any (x EID) EID { 
    return F_Generate_retreive_list_any(ANY(x) )} 
  
// *******************************************************************
// *       Part 3: interface declarations                            *
// *******************************************************************
// How to declare a sort in Go. The boolean tells if we are in an external
// mode , in which case we produce the C sort. Otherwise, we use OIDs.
// THERE are 5 sorts in go : int, float, char,  any (object) and EID
// there are 7 sorts in CLAIRE : int, float, char, object, string, function, any
/* {1} The go function for: interface!(self:class) [status=0] */
func F_Generate_interface_I_class (self *ClaireClass )  { 
    // procedure body with s = void 
if (self.Id() == C_void.Id()) { 
      PRINC("void ")
      }  else if (self.Id() == C_integer.Id()) { 
      PRINC("int")
      }  else if (self.Id() == C_float.Id()) { 
      PRINC("float64")
      }  else if (self.Id() == C_char.Id()) { 
      PRINC("rune")
      }  else if (self.Id() == Optimize.C_EID.Id()) { 
      PRINC("EID")
      } else {
      PRINC("*")
      F_Generate_go_class_class(self)
      PRINC(" ")
      } 
    } 
  
// The EID go function for: interface! @ class (throw: false) 
func E_Generate_interface_I_class (self EID) EID { 
    F_Generate_interface_I_class(ToClass(OBJ(self)) )
    return EVOID} 
  
// general translation method: x is an expression that must be translated
// to a CLAIRE object (*ClaireX). x is known to be functional ! s is the sort for x.
/* {1} The go function for: to_cl(c:go_producer,x:any,s:class) [status=1] */
func (c *GenerateGoProducer ) ToCl (x *ClaireAny ,s *ClaireClass ) EID { 
    var Result EID 
    if (s.Id() == C_void.Id()) { 
      PRINC("Void(")
      /*g_try(v2:"Result",loop:true) */
      Result = Core.F_CALL(C_Generate_g_expression,ARGS(x.ToEID(),EID{C_any.Id(),0}))
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(")")
      Result = EVOID
      }
      }  else if (s.Id() == C_integer.Id()) { 
      PRINC("MakeInteger(")
      /*g_try(v2:"Result",loop:true) */
      Result = Core.F_CALL(C_Generate_g_expression,ARGS(x.ToEID(),EID{C_integer.Id(),0}))
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(")")
      Result = EVOID
      }
      }  else if (s.Id() == C_float.Id()) { 
      PRINC("MakeFloat(")
      /*g_try(v2:"Result",loop:true) */
      Result = Core.F_CALL(C_Generate_g_expression,ARGS(x.ToEID(),EID{C_float.Id(),0}))
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(")")
      Result = EVOID
      }
      }  else if (s.Id() == C_char.Id()) { 
      PRINC("MakeChar(")
      /*g_try(v2:"Result",loop:true) */
      Result = Core.F_CALL(C_Generate_g_expression,ARGS(x.ToEID(),EID{C_char.Id(),0}))
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(")")
      Result = EVOID
      }
      }  else if ((s.IsIn(C_object) == CTRUE) || 
        ((s.Id() == C_any.Id()) || 
          (s.Id() == C_primitive.Id()))) { 
      Result = Core.F_CALL(C_Generate_g_expression,ARGS(x.ToEID(),EID{s.Id(),0}))
      } else {
      Result = ToException(Core.C_general_error.Make(MakeString("[internal] to_cl for a ~S is not implemented").Id(),MakeConstantList(s.Id()).Id())).Close()
      } 
    return Result} 
  
// The EID go function for: to_cl @ go_producer (throw: true) 
func E_Generate_to_cl_go_producer (c EID,x EID,s EID) EID { 
    return ToGenerateGoProducer(OBJ(c)).ToCl(ANY(x),ToClass(OBJ(s)) )} 
  
// reverse function : produce a native forme from a claire object
// quite simple with go since for object, OID is the object
/* {1} The go function for: to_c(c:go_producer,x:any,s:class) [status=0] */
func (c *GenerateGoProducer ) ToC (x *ClaireAny ,s *ClaireClass )  { 
    // procedure body with s = void 
if (x == CNULL) { 
      PRINC("CNULL")
      }  else if ((s.Id() == C_integer.Id()) || 
        ((s.Id() == C_float.Id()) || 
          ((s.Id() == C_string.Id()) || 
            ((s.Id() == C_char.Id()) || 
              (s.Id() == C_function.Id()))))) { 
      Core.F_CALL(C_Generate_expression,ARGS(x.ToEID(),EID{s.Id(),0}))
      PRINC(".Value")
      } else {
      Core.F_CALL(C_Generate_c_expression,ARGS(x.ToEID(),EID{s.Id(),0}))
      } 
    } 
  
// The EID go function for: to_c @ go_producer (throw: false) 
func E_Generate_to_c_go_producer (c EID,x EID,s EID) EID { 
    ToGenerateGoProducer(OBJ(c)).ToC(ANY(x),ToClass(OBJ(s)) )
    return EVOID} 
  
// new for go: compile to an EID form (128 bit generic representation)
// s is the expected sort
/* {1} The go function for: to_eid(c:go_producer,x:any,s:class) [status=1] */
func (c *GenerateGoProducer ) ToEid (x *ClaireAny ,s *ClaireClass ) EID { 
    var Result EID 
    if (s.Id() == C_void.Id()) { 
      PRINC("EVOID")
      Result = EVOID
      }  else if (s.Id() == C_integer.Id()) { 
      PRINC("EID{C__INT,IVAL(")
      /*g_try(v2:"Result",loop:true) */
      Result = Core.F_CALL(C_Generate_g_expression,ARGS(x.ToEID(),EID{C_integer.Id(),0}))
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(")}")
      Result = EVOID
      }
      }  else if (s.Id() == C_float.Id()) { 
      PRINC("EID{C__FLOAT,FVAL(")
      /*g_try(v2:"Result",loop:true) */
      Result = Core.F_CALL(C_Generate_g_expression,ARGS(x.ToEID(),EID{C_float.Id(),0}))
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(")}")
      Result = EVOID
      }
      }  else if (s.Id() == C_char.Id()) { 
      PRINC("EID{C__CHAR,CVAL(")
      /*g_try(v2:"Result",loop:true) */
      Result = Core.F_CALL(C_Generate_g_expression,ARGS(x.ToEID(),EID{C_char.Id(),0}))
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(")}")
      Result = EVOID
      }
      }  else if ((s.Id() == C_string.Id()) || 
        ((s.Id() == C_function.Id()) || 
          (s.IsIn(C_object) == CTRUE))) { 
      PRINC("EID{")
      /*g_try(v2:"Result",loop:true) */
      Result = Core.F_CALL(C_Generate_g_expression,ARGS(x.ToEID(),EID{C_any.Id(),0}))
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(",0}")
      Result = EVOID
      }
      }  else if ((s.Id() == C_any.Id()) || 
        (s.Id() == C_primitive.Id())) { 
      /*g_try(v2:"Result",loop:true) */
      Result = c.ToCl(x,s)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(".ToEID()")
      Result = EVOID
      }
      } else {
      Result = ToException(Core.C_general_error.Make(MakeString("[internal] to_eid for a ~S is not implemented").Id(),MakeConstantList(s.Id()).Id())).Close()
      } 
    return Result} 
  
// The EID go function for: to_eid @ go_producer (throw: true) 
func E_Generate_to_eid_go_producer (c EID,x EID,s EID) EID { 
    return ToGenerateGoProducer(OBJ(c)).ToEid(ANY(x),ToClass(OBJ(s)) )} 
  
// reciprocate with an expected class e / used for variables
/* {1} The go function for: from_eid(c:go_producer,x:string,e:class) [status=1] */
func (c *GenerateGoProducer ) FromEid (x *ClaireString ,e *ClaireClass ) EID { 
    var Result EID 
    { var s *ClaireClass   = e.Class_I()
      /*g_try(v2:"Result",loop:true) */
      Result = F_Generate_eid_prefix_class(s)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      F_c_princ_string(x)
      F_Generate_eid_post_class(s)
      Result = EVOID
      }
      } 
    return Result} 
  
// The EID go function for: from_eid @ go_producer (throw: true) 
func E_Generate_from_eid_go_producer (c EID,x EID,e EID) EID { 
    return ToGenerateGoProducer(OBJ(c)).FromEid(ToString(OBJ(x)),ToClass(OBJ(e)) )} 
  
// reciprocate : move from EID to a sort s (if s = any, do nothing )
/* {1} The go function for: eid_prefix(s:class) [status=1] */
func F_Generate_eid_prefix_class (s *ClaireClass ) EID { 
    var Result EID 
    if ((s.Id() == Optimize.C_EID.Id()) || 
        (s.Id() == C_void.Id())) { 
      Result = EID{CNIL.Id(),0}
      }  else if (s.Id() == C_integer.Id()) { 
      PRINC("INT(")
      Result = EVOID
      }  else if (s.Id() == C_float.Id()) { 
      PRINC("FLOAT(")
      Result = EVOID
      }  else if (s.Id() == C_char.Id()) { 
      PRINC("CHAR(")
      Result = EVOID
      }  else if ((s.Id() == C_any.Id()) || 
        (s.Id() == C_primitive.Id())) { 
      PRINC("ANY(")
      Result = EVOID
      }  else if ((ToType(s.Id()).Included(ToType(C_object.Id())) == CTRUE) || 
        ((s.Id() == C_array.Id()) || 
          ((s.Id() == C_string.Id()) || 
            ((s.Id() == C_port.Id()) || 
              (s.Id() == C_function.Id()))))) { 
      F_Generate_cast_class_class(s)
      PRINC("(OBJ(")
      Result = EVOID
      }  else if (s.Id() != C_any.Id()) { 
      Result = ToException(Core.C_general_error.Make(MakeString("what the fuck: eid prefix for ~S").Id(),MakeConstantList(s.Id()).Id())).Close()
      } else {
      Result = EID{CFALSE.Id(),0}
      } 
    return Result} 
  
// The EID go function for: eid_prefix @ class (throw: true) 
func E_Generate_eid_prefix_class (s EID) EID { 
    return F_Generate_eid_prefix_class(ToClass(OBJ(s)) )} 
  
/* {1} The go function for: eid_post(s:class) [status=0] */
func F_Generate_eid_post_class (s *ClaireClass )  { 
    // procedure body with s = void 
if ((s.Id() == Optimize.C_EID.Id()) || 
        (s.Id() == C_void.Id())) { 
      
      }  else if ((s.Id() == C_char.Id()) || 
        ((s.Id() == C_any.Id()) || 
          (s.Id() == C_primitive.Id()))) { 
      PRINC(")")
      }  else if ((ToType(s.Id()).Included(ToType(C_object.Id())) == CTRUE) || 
        ((s.Id() == C_array.Id()) || 
          ((s.Id() == C_string.Id()) || 
            ((s.Id() == C_port.Id()) || 
              (s.Id() == C_function.Id()))))) { 
      PRINC("))")
      }  else if (s.Id() != C_any.Id()) { 
      PRINC(")")
      } 
    } 
  
// The EID go function for: eid_post @ class (throw: false) 
func E_Generate_eid_post_class (s EID) EID { 
    F_Generate_eid_post_class(ToClass(OBJ(s)) )
    return EVOID} 
  
// move from an integer to a EID or Object
/* {1} The go function for: integer_prefix(s:class) [status=0] */
func F_Generate_integer_prefix_class (s *ClaireClass )  { 
    // procedure body with s = void 
if (s.Id() == Optimize.C_EID.Id()) { 
      PRINC("EID{C__INT,IVAL(")
      }  else if (s.Id() == C_any.Id()) { 
      PRINC("MakeInteger(")
      } 
    } 
  
// The EID go function for: integer_prefix @ class (throw: false) 
func E_Generate_integer_prefix_class (s EID) EID { 
    F_Generate_integer_prefix_class(ToClass(OBJ(s)) )
    return EVOID} 
  
// move from an integer to a EID or Object
/* {1} The go function for: float_prefix(s:class) [status=0] */
func F_Generate_float_prefix_class (s *ClaireClass )  { 
    // procedure body with s = void 
if (s.Id() == Optimize.C_EID.Id()) { 
      PRINC("EID{C__FLOAT,FVAL(")
      }  else if (s.Id() == C_any.Id()) { 
      PRINC("MakeFloat(")
      } 
    } 
  
// The EID go function for: float_prefix @ class (throw: false) 
func E_Generate_float_prefix_class (s EID) EID { 
    F_Generate_float_prefix_class(ToClass(OBJ(s)) )
    return EVOID} 
  
// move from an integer to a EID or Object
/* {1} The go function for: char_prefix(s:class) [status=0] */
func F_Generate_char_prefix_class (s *ClaireClass )  { 
    // procedure body with s = void 
if (s.Id() == Optimize.C_EID.Id()) { 
      PRINC("EID{C__CHAR,CVAL(")
      }  else if (s.Id() == C_any.Id()) { 
      PRINC("MakeChar(")
      } 
    } 
  
// The EID go function for: char_prefix @ class (throw: false) 
func E_Generate_char_prefix_class (s EID) EID { 
    F_Generate_char_prefix_class(ToClass(OBJ(s)) )
    return EVOID} 
  
// move from an integer to a EID or Object
/* {1} The go function for: string_prefix(s:class) [status=0] */
func F_Generate_string_prefix_class (s *ClaireClass )  { 
    // procedure body with s = void 
if (s.Id() == Optimize.C_EID.Id()) { 
      PRINC("EID{")
      }  else if (s.Id() == C_any.Id()) { 
      PRINC("(")
      } 
    } 
  
// The EID go function for: string_prefix @ class (throw: false) 
func E_Generate_string_prefix_class (s EID) EID { 
    F_Generate_string_prefix_class(ToClass(OBJ(s)) )
    return EVOID} 
  
/* {1} The go function for: string_post(s:class) [status=0] */
func F_Generate_string_post_class (s *ClaireClass )  { 
    // procedure body with s = void 
if (s.Id() == Optimize.C_EID.Id()) { 
      PRINC(".Id(),0}")
      }  else if (s.Id() == C_any.Id()) { 
      PRINC(").Id()")
      } 
    } 
  
// The EID go function for: string_post @ class (throw: false) 
func E_Generate_string_post_class (s EID) EID { 
    F_Generate_string_post_class(ToClass(OBJ(s)) )
    return EVOID} 
  
// works for integer, float, char
/* {1} The go function for: native_post(s:class) [status=0] */
func F_Generate_native_post_class (s *ClaireClass )  { 
    // procedure body with s = void 
if (s.Id() == Optimize.C_EID.Id()) { 
      PRINC(")}")
      }  else if (s.Id() == C_any.Id()) { 
      PRINC(").Id()")
      } 
    } 
  
// The EID go function for: native_post @ class (throw: false) 
func E_Generate_native_post_class (s EID) EID { 
    F_Generate_native_post_class(ToClass(OBJ(s)) )
    return EVOID} 
  
// move from ClaireId (inferred) to s (expected)
/* {1} The go function for: object_prefix(inferred:class,expected:class) [status=0] */
func F_Generate_object_prefix_class (inferred *ClaireClass ,expected *ClaireClass )  { 
    // procedure body with s = void 
if (expected.Id() == Optimize.C_EID.Id()) { 
      if (ToType(inferred.Id()).Included(ToType(C_object.Id())) == CTRUE) { 
        PRINC("EID{")
        } 
      }  else if (expected.Id() == inferred.Id()) { 
      
      }  else if (expected.Id() == C_char.Id()) { 
      PRINC("ToChar(")
      }  else if (ToType(expected.Id()).Included(ToType(C_primitive.Id())) == CTRUE) { 
      F_Generate_cast_class_class(expected)
      PRINC("(")
      }  else if (ToType(expected.Id()).Included(ToType(C_object.Id())) == CTRUE) { 
      F_Generate_cast_class_class(expected)
      PRINC("(")
      } 
    } 
  
// The EID go function for: object_prefix @ class (throw: false) 
func E_Generate_object_prefix_class (inferred EID,expected EID) EID { 
    F_Generate_object_prefix_class(ToClass(OBJ(inferred)),ToClass(OBJ(expected)) )
    return EVOID} 
  
// s: expected
/* {1} The go function for: object_post(inferred:class,s:class) [status=0] */
func F_Generate_object_post_class (inferred *ClaireClass ,s *ClaireClass )  { 
    // procedure body with s = void 
if (s.Id() == Optimize.C_EID.Id()) { 
      if (ToType(inferred.Id()).Included(ToType(C_object.Id())) != CTRUE) { 
        PRINC(".ToEID()")
        } else {
        PRINC(".Id(),0}")
        } 
      }  else if (s.Id() == inferred.Id()) { 
      
      }  else if ((s.Id() == C_integer.Id()) || 
        ((s.Id() == C_float.Id()) || 
          (s.Id() == C_char.Id()))) { 
      if (inferred.Id() == C_any.Id()) { 
        PRINC(").Value")
        } else {
        PRINC(".Id()).Value")
        } 
      }  else if ((ToType(s.Id()).Included(ToType(C_object.Id())) == CTRUE) || 
        (ToType(s.Id()).Included(ToType(C_primitive.Id())) == CTRUE)) { 
      if (inferred.Id() == C_any.Id()) { 
        PRINC(")")
        } else {
        PRINC(".Id())")
        } 
      }  else if (s.Id() == C_any.Id()) { 
      PRINC(".Id()")
      } 
    } 
  
// The EID go function for: object_post @ class (throw: false) 
func E_Generate_object_post_class (inferred EID,s EID) EID { 
    F_Generate_object_post_class(ToClass(OBJ(inferred)),ToClass(OBJ(s)) )
    return EVOID} 
  
// generic version that applies to everything (s1:infered) => *_prefix(s2:expected)
/* {1} The go function for: cast_prefix(s1:class,s2:class) [status=1] */
func F_Generate_cast_prefix_class (s1 *ClaireClass ,s2 *ClaireClass ) EID { 
    var Result EID 
    if (s1.Id() == Optimize.C_EID.Id()) { 
      Result = F_Generate_eid_prefix_class(s2)
      }  else if (s1.Id() == C_void.Id()) { 
      Result = EID{CNIL.Id(),0}
      }  else if (s1.Id() == C_integer.Id()) { 
      F_Generate_integer_prefix_class(s2)
      Result = EVOID
      }  else if (s1.Id() == C_float.Id()) { 
      F_Generate_float_prefix_class(s2)
      Result = EVOID
      }  else if (s1.Id() == C_char.Id()) { 
      F_Generate_char_prefix_class(s2)
      Result = EVOID
      }  else if (s1.Id() == C_string.Id()) { 
      F_Generate_string_prefix_class(s2)
      Result = EVOID
      } else {
      F_Generate_object_prefix_class(s1,s2)
      Result = EVOID
      } 
    return Result} 
  
// The EID go function for: cast_prefix @ class (throw: true) 
func E_Generate_cast_prefix_class (s1 EID,s2 EID) EID { 
    return F_Generate_cast_prefix_class(ToClass(OBJ(s1)),ToClass(OBJ(s2)) )} 
  
// generic version that applies to everything (s1) => *_post(s2)
// s1 is the goType of the expression, s2 is the expected
/* {1} The go function for: cast_post(s1:class,s2:class) [status=0] */
func F_Generate_cast_post_class (s1 *ClaireClass ,s2 *ClaireClass )  { 
    // procedure body with s = void 
if (s1.Id() == Optimize.C_EID.Id()) { 
      F_Generate_eid_post_class(s2)
      }  else if (s1.Id() == C_void.Id()) { 
      
      }  else if ((s1.Id() == C_integer.Id()) || 
        ((s1.Id() == C_float.Id()) || 
          (s1.Id() == C_char.Id()))) { 
      F_Generate_native_post_class(s2)
      }  else if (s1.Id() == C_string.Id()) { 
      F_Generate_string_post_class(s2)
      } else {
      F_Generate_object_post_class(s1,s2)
      } 
    } 
  
// The EID go function for: cast_post @ class (throw: false) 
func E_Generate_cast_post_class (s1 EID,s2 EID) EID { 
    F_Generate_cast_post_class(ToClass(OBJ(s1)),ToClass(OBJ(s2)) )
    return EVOID} 
  
// *******************************************************************
// *       Part 4: use of language dependent patterns (macros)       *
// *******************************************************************
// when we print an equality, we do not need to_CL !
// id is used to force the identifiability (use = vs equal)
/* {1} The go function for: equal_exp(c:go_producer,a1:any,pos?:boolean,a2:any,id?:any) [status=1] */
func (c *GenerateGoProducer ) EqualExp (a1 *ClaireAny ,pos_ask *ClaireBoolean ,a2 *ClaireAny ,id_ask *ClaireAny ) EID { 
    var Result EID 
    var g0066I *ClaireBoolean  
    var try_1 EID 
    /*g_try(v2:"try_1",loop:false) */
    { 
      var v_and2 *ClaireBoolean  
      
      var try_2 EID 
      /*g_try(v2:"try_2",loop:false) */
      { var arg_3 *ClaireClass  
        _ = arg_3
        var try_4 EID 
        /*g_try(v2:"try_4",loop:false) */
        try_4 = Language.F_static_type_any(a1)
        /* ERROR PROTECTION INSERTED (arg_3-try_2) */
        if ErrorIn(try_4) {try_2 = try_4
        } else {
        arg_3 = ToClass(OBJ(try_4))
        try_2 = EID{Equal(arg_3.Id(),C_string.Id()).Id(),0}
        }
        } 
      /* ERROR PROTECTION INSERTED (v_and2-try_1) */
      if ErrorIn(try_2) {try_1 = try_2
      } else {
      v_and2 = ToBoolean(OBJ(try_2))
      if (v_and2 == CFALSE) {try_1 = EID{CFALSE.Id(),0}
      } else { 
        var try_5 EID 
        /*g_try(v2:"try_5",loop:false) */
        { var arg_6 *ClaireClass  
          _ = arg_6
          var try_7 EID 
          /*g_try(v2:"try_7",loop:false) */
          try_7 = Language.F_static_type_any(a2)
          /* ERROR PROTECTION INSERTED (arg_6-try_5) */
          if ErrorIn(try_7) {try_5 = try_7
          } else {
          arg_6 = ToClass(OBJ(try_7))
          try_5 = EID{Equal(arg_6.Id(),C_string.Id()).Id(),0}
          }
          } 
        /* ERROR PROTECTION INSERTED (v_and2-try_1) */
        if ErrorIn(try_5) {try_1 = try_5
        } else {
        v_and2 = ToBoolean(OBJ(try_5))
        if (v_and2 == CFALSE) {try_1 = EID{CFALSE.Id(),0}
        } else { 
          try_1 = EID{CTRUE.Id(),0}} 
        } 
      }}
      } 
    /* ERROR PROTECTION INSERTED (g0066I-Result) */
    if ErrorIn(try_1) {Result = try_1
    } else {
    g0066I = ToBoolean(OBJ(try_1))
    if (g0066I == CTRUE) { 
      PRINC("(")
      /*g_try(v2:"Result",loop:true) */
      Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_string.Id(),0}))
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(".Value ")
      F_Generate_sign_equal_boolean(pos_ask)
      PRINC(" ")
      /*g_try(v2:"Result",loop:true) */
      Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{C_string.Id(),0}))
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(".Value)")
      Result = EVOID
      }}
      }  else if ((c.CharExp_ask(a1) == CTRUE) || 
        (c.CharExp_ask(a2) == CTRUE)) { 
      PRINC("(")
      /*g_try(v2:"Result",loop:true) */
      Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_char.Id(),0}))
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(" ")
      F_Generate_sign_equal_boolean(pos_ask)
      PRINC(" ")
      /*g_try(v2:"Result",loop:true) */
      Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{C_char.Id(),0}))
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(")")
      Result = EVOID
      }}
      } else {
      var g0067I *ClaireBoolean  
      var try_8 EID 
      /*g_try(v2:"try_8",loop:false) */
      { 
        var v_and3 *ClaireBoolean  
        
        var try_9 EID 
        /*g_try(v2:"try_9",loop:false) */
        { 
          /* Or stat: v="try_9", loop=false */
          var v_or4 *ClaireBoolean  
          
          /* Or stat: try boolean! @ any(id?) with try:false, v="try_9", loop=false */
          v_or4 = F_boolean_I_any(id_ask)
          if (v_or4 == CTRUE) {try_9 = EID{CTRUE.Id(),0}
          } else { 
            /* Or stat: try Compile/identifiable? @ any(a1) with try:true, v="try_9", loop=false */
            var try_10 EID 
            /*g_try(v2:"try_10",loop:false) */
            try_10 = Optimize.F_Compile_identifiable_ask_any(a1)
            /* ERROR PROTECTION INSERTED (v_or4-try_9) */
            if ErrorIn(try_10) {try_9 = try_10
            } else {
            v_or4 = ToBoolean(OBJ(try_10))
            if (v_or4 == CTRUE) {try_9 = EID{CTRUE.Id(),0}
            } else { 
              /* Or stat: try Compile/identifiable? @ any(a2) with try:true, v="try_9", loop=false */
              var try_11 EID 
              /*g_try(v2:"try_11",loop:false) */
              try_11 = Optimize.F_Compile_identifiable_ask_any(a2)
              /* ERROR PROTECTION INSERTED (v_or4-try_9) */
              if ErrorIn(try_11) {try_9 = try_11
              } else {
              v_or4 = ToBoolean(OBJ(try_11))
              if (v_or4 == CTRUE) {try_9 = EID{CTRUE.Id(),0}
              } else { 
                /* Or stat: try = @ any(g_sort @ any(a1),float) with try:true, v="try_9", loop=false */
                var try_12 EID 
                /*g_try(v2:"try_12",loop:false) */
                { var arg_13 *ClaireClass  
                  _ = arg_13
                  var try_14 EID 
                  /*g_try(v2:"try_14",loop:false) */
                  try_14 = F_Generate_g_sort_any(a1)
                  /* ERROR PROTECTION INSERTED (arg_13-try_12) */
                  if ErrorIn(try_14) {try_12 = try_14
                  } else {
                  arg_13 = ToClass(OBJ(try_14))
                  try_12 = EID{Equal(arg_13.Id(),C_float.Id()).Id(),0}
                  }
                  } 
                /* ERROR PROTECTION INSERTED (v_or4-try_9) */
                if ErrorIn(try_12) {try_9 = try_12
                } else {
                v_or4 = ToBoolean(OBJ(try_12))
                if (v_or4 == CTRUE) {try_9 = EID{CTRUE.Id(),0}
                } else { 
                  try_9 = EID{CFALSE.Id(),0}} 
                } 
              } 
            } 
          }}}
          } 
        /* ERROR PROTECTION INSERTED (v_and3-try_8) */
        if ErrorIn(try_9) {try_8 = try_9
        } else {
        v_and3 = ToBoolean(OBJ(try_9))
        if (v_and3 == CFALSE) {try_8 = EID{CFALSE.Id(),0}
        } else { 
          var try_15 EID 
          /*g_try(v2:"try_15",loop:false) */
          { var arg_16 *ClaireClass  
            _ = arg_16
            var try_18 EID 
            /*g_try(v2:"try_18",loop:false) */
            try_18 = F_Generate_g_sort_any(a1)
            /* ERROR PROTECTION INSERTED (arg_16-try_15) */
            if ErrorIn(try_18) {try_15 = try_18
            } else {
            arg_16 = ToClass(OBJ(try_18))
            { var arg_17 *ClaireClass  
              _ = arg_17
              var try_19 EID 
              /*g_try(v2:"try_19",loop:false) */
              try_19 = F_Generate_g_sort_any(a2)
              /* ERROR PROTECTION INSERTED (arg_17-try_15) */
              if ErrorIn(try_19) {try_15 = try_19
              } else {
              arg_17 = ToClass(OBJ(try_19))
              try_15 = EID{Equal(arg_16.Id(),arg_17.Id()).Id(),0}
              }
              } 
            }
            } 
          /* ERROR PROTECTION INSERTED (v_and3-try_8) */
          if ErrorIn(try_15) {try_8 = try_15
          } else {
          v_and3 = ToBoolean(OBJ(try_15))
          if (v_and3 == CFALSE) {try_8 = EID{CFALSE.Id(),0}
          } else { 
            try_8 = EID{CTRUE.Id(),0}} 
          } 
        }}
        } 
      /* ERROR PROTECTION INSERTED (g0067I-Result) */
      if ErrorIn(try_8) {Result = try_8
      } else {
      g0067I = ToBoolean(OBJ(try_8))
      if (g0067I == CTRUE) { 
        /*g_try(v2:"Result",loop:true) */
        var g0068I *ClaireBoolean  
        var try_20 EID 
        /*g_try(v2:"try_20",loop:false) */
        { var arg_21 *ClaireType  
          _ = arg_21
          var try_22 EID 
          /*g_try(v2:"try_22",loop:false) */
          { var arg_23 *ClaireClass  
            _ = arg_23
            var try_25 EID 
            /*g_try(v2:"try_25",loop:false) */
            try_25 = Optimize.F_Compile_stupid_t_any1(a1)
            /* ERROR PROTECTION INSERTED (arg_23-try_22) */
            if ErrorIn(try_25) {try_22 = try_25
            } else {
            arg_23 = ToClass(OBJ(try_25))
            { var arg_24 *ClaireClass  
              _ = arg_24
              var try_26 EID 
              /*g_try(v2:"try_26",loop:false) */
              try_26 = Optimize.F_Compile_stupid_t_any1(a2)
              /* ERROR PROTECTION INSERTED (arg_24-try_22) */
              if ErrorIn(try_26) {try_22 = try_26
              } else {
              arg_24 = ToClass(OBJ(try_26))
              try_22 = EID{Core.F_glb_class(arg_23,ToType(arg_24.Id())).Id(),0}
              }
              } 
            }
            } 
          /* ERROR PROTECTION INSERTED (arg_21-try_20) */
          if ErrorIn(try_22) {try_20 = try_22
          } else {
          arg_21 = ToType(OBJ(try_22))
          try_20 = EID{Equal(arg_21.Id(),CEMPTY.Id()).Id(),0}
          }
          } 
        /* ERROR PROTECTION INSERTED (g0068I-Result) */
        if ErrorIn(try_20) {Result = try_20
        } else {
        g0068I = ToBoolean(OBJ(try_20))
        if (g0068I == CTRUE) { 
          Optimize.F_Compile_warn_void()
          Result = Core.F_tformat_string(MakeString("~S = ~S will fail ! [263]"),2,MakeConstantList(a1,a2))
          } else {
          Result = EID{CFALSE.Id(),0}
          } 
        }
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC("(")
        /*g_try(v2:"Result",loop:true) */
        { var arg_27 *ClaireClass  
          _ = arg_27
          var try_28 EID 
          /*g_try(v2:"try_28",loop:false) */
          try_28 = F_Generate_g_sort_any(a1)
          /* ERROR PROTECTION INSERTED (arg_27-Result) */
          if ErrorIn(try_28) {Result = try_28
          } else {
          arg_27 = ToClass(OBJ(try_28))
          Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{arg_27.Id(),0}))
          }
          } 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC(" ")
        F_Generate_sign_equal_boolean(pos_ask)
        PRINC(" ")
        /*g_try(v2:"Result",loop:true) */
        { var arg_29 *ClaireClass  
          _ = arg_29
          var try_30 EID 
          /*g_try(v2:"try_30",loop:false) */
          try_30 = F_Generate_g_sort_any(a1)
          /* ERROR PROTECTION INSERTED (arg_29-Result) */
          if ErrorIn(try_30) {Result = try_30
          } else {
          arg_29 = ToClass(OBJ(try_30))
          Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{arg_29.Id(),0}))
          }
          } 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC(")")
        Result = EVOID
        }}
        }
        } else {
        var g0069I *ClaireBoolean  
        var try_31 EID 
        /*g_try(v2:"try_31",loop:false) */
        { var arg_32 *ClaireClass  
          _ = arg_32
          var try_33 EID 
          /*g_try(v2:"try_33",loop:false) */
          try_33 = Optimize.F_Compile_stupid_t_any1(a2)
          /* ERROR PROTECTION INSERTED (arg_32-try_31) */
          if ErrorIn(try_33) {try_31 = try_33
          } else {
          arg_32 = ToClass(OBJ(try_33))
          try_31 = EID{Equal(arg_32.Id(),C_integer.Id()).Id(),0}
          }
          } 
        /* ERROR PROTECTION INSERTED (g0069I-Result) */
        if ErrorIn(try_31) {Result = try_31
        } else {
        g0069I = ToBoolean(OBJ(try_31))
        if (g0069I == CTRUE) { 
          if (pos_ask != CTRUE) { 
            PRINC("!")
            } 
          /*g_try(v2:"Result",loop:true) */
          Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_any.Id(),0}))
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(".IsInt(")
          /*g_try(v2:"Result",loop:true) */
          Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{C_integer.Id(),0}))
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(")")
          Result = EVOID
          }}
          } else {
          PRINC("(Equal(")
          /*g_try(v2:"Result",loop:true) */
          Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_any.Id(),0}))
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(",")
          /*g_try(v2:"Result",loop:true) */
          Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{C_any.Id(),0}))
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(") ")
          F_Generate_sign_equal_boolean(pos_ask)
          PRINC(" CTRUE)")
          Result = EVOID
          }}
          } 
        }
        } 
      }
      } 
    }
    return Result} 
  
// The EID go function for: equal_exp @ go_producer (throw: true) 
func E_Generate_equal_exp_go_producer (c EID,a1 EID,pos_ask EID,a2 EID,id_ask EID) EID { 
    return ToGenerateGoProducer(OBJ(c)).EqualExp(ANY(a1),
      ToBoolean(OBJ(pos_ask)),
      ANY(a2),
      ANY(id_ask) )} 
  
// new: special code for char
// CLAIRE 4 : removed char_exp => g_expression(x,char) should work
/* {1} The go function for: char_exp?(c:go_producer,x:any) [status=0] */
func (c *GenerateGoProducer ) CharExp_ask (x *ClaireAny ) *ClaireBoolean  { 
    // procedure body with s = boolean 
var Result *ClaireBoolean  
    if (C_char.Id() == x.Isa.Id()) { 
      Result = CTRUE
      }  else if (x.Isa.IsIn(Language.C_Call_method) == CTRUE) { 
      { var g0071 *Language.CallMethod   = Language.To_CallMethod(x)
        _ = g0071
        { var m *ClaireMethod   = g0071.Arg
          Result = MakeBoolean((m.Id() == C_Generate__starnth_1_string_star.Value) || ((m.Id() == C_Generate__starnth_string_star.Value) && 
              (Optimize.C_compiler.Safety >= 2)))
          } 
        } 
      } else {
      Result = CFALSE
      } 
    return Result} 
  
// The EID go function for: char_exp? @ go_producer (throw: false) 
func E_Generate_char_exp_ask_go_producer (c EID,x EID) EID { 
    return EID{ToGenerateGoProducer(OBJ(c)).CharExp_ask(ANY(x) ).Id(),0}} 
  
// reads the member x from an expression self of expected type s
/* {1} The go function for: c_member(c:go_producer,self:any,s:class,x:property) [status=1] */
func (c *GenerateGoProducer ) CMember (self *ClaireAny ,s *ClaireClass ,x *ClaireProperty ) EID { 
    var Result EID 
    /*g_try(v2:"Result",loop:true) */
    Result = Core.F_CALL(C_Generate_g_expression,ARGS(self.ToEID(),EID{s.Id(),0}))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(".")
    F_Generate_cap_short_symbol(x.Name)
    PRINC("")
    Result = EVOID
    }
    return Result} 
  
// The EID go function for: c_member @ go_producer (throw: true) 
func E_Generate_c_member_go_producer (c EID,self EID,s EID,x EID) EID { 
    return ToGenerateGoProducer(OBJ(c)).CMember(ANY(self),
      ToClass(OBJ(s)),
      ToProperty(OBJ(x)) )} 
  
// generic for bags
/* {1} The go function for: bag_expression(c:go_producer,cl:class,l:list,t:type) [status=1] */
func (c *GenerateGoProducer ) BagExpression (cl *ClaireClass ,l *ClaireList ,t *ClaireType ) EID { 
    var Result EID 
    if ((l.Length() == 0) && 
        (cl.Id() != C_tuple.Id())) { 
      /*g_try(v2:"Result",loop:true) */
      { var arg_1 *ClaireAny  
        _ = arg_1
        var try_2 EID 
        /*g_try(v2:"try_2",loop:false) */
        try_2 = Core.F_CALL(Optimize.C_c_code,ARGS(EID{t.Id(),0},EID{C_object.Id(),0}))
        /* ERROR PROTECTION INSERTED (arg_1-Result) */
        if ErrorIn(try_2) {Result = try_2
        } else {
        arg_1 = ANY(try_2)
        Result = Core.F_CALL(C_Generate_g_expression,ARGS(arg_1.ToEID(),EID{C_type.Id(),0}))
        }
        } 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(".Empty")
      F_Generate_cap_short_symbol(cl.Name)
      PRINC("()")
      Result = EVOID
      }
      }  else if ((Equal(t.Id(),CEMPTY.Id()) == CTRUE) || 
        (t.Id() == C_void.Id())) { 
      PRINC("Make")
      { var arg_3 *ClaireString  
        _ = arg_3
        if (cl.Id() == C_set.Id()) { 
          arg_3 = MakeString("ConstantSet")
          }  else if (cl.Id() == C_list.Id()) { 
          arg_3 = MakeString("ConstantList")
          } else {
          arg_3 = MakeString("Tuple")
          } 
        F_princ_string(arg_3)
        } 
      PRINC("(")
      /*g_try(v2:"Result",loop:true) */
      Result = F_Generate_args_list_list(l,C_any)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(")")
      Result = EVOID
      }
      }  else if ((cl.Id() == C_list.Id()) && 
        (t.Id() == C_integer.Id())) { 
      PRINC("MakeListInteger(")
      /*g_try(v2:"Result",loop:true) */
      Result = F_Generate_args_list_list(l,C_integer)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(")")
      Result = EVOID
      }
      } else {
      PRINC("Make")
      F_Generate_cap_short_symbol(cl.Name)
      PRINC("(")
      /*g_try(v2:"Result",loop:true) */
      { var arg_4 *ClaireAny  
        _ = arg_4
        var try_5 EID 
        /*g_try(v2:"try_5",loop:false) */
        try_5 = Core.F_CALL(Optimize.C_c_code,ARGS(EID{t.Id(),0},EID{C_object.Id(),0}))
        /* ERROR PROTECTION INSERTED (arg_4-Result) */
        if ErrorIn(try_5) {Result = try_5
        } else {
        arg_4 = ANY(try_5)
        Result = Core.F_CALL(C_Generate_g_expression,ARGS(arg_4.ToEID(),EID{C_type.Id(),0}))
        }
        } 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(",")
      /*g_try(v2:"Result",loop:true) */
      Result = F_Generate_args_list_list(l,C_any)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(")")
      Result = EVOID
      }}
      } 
    return Result} 
  
// The EID go function for: bag_expression @ go_producer (throw: true) 
func E_Generate_bag_expression_go_producer (c EID,cl EID,l EID,t EID) EID { 
    return ToGenerateGoProducer(OBJ(c)).BagExpression(ToClass(OBJ(cl)),
      ToList(OBJ(l)),
      ToType(OBJ(t)) )} 
  
// *******************************************************************
// *       Part 5: Utilities :                                       *
// *******************************************************************
// a constant can be evaluated with no cost in an IfThenElse(test,a,b)
/* {1} The go function for: constant?(self:any) [status=0] */
func F_Generate_constant_ask_any (self *ClaireAny ) *ClaireBoolean  { 
    if ((self.Isa.IsIn(C_thing) == CTRUE) || 
        ((C_boolean.Id() == self.Isa.Id()) || 
          ((self.Isa.IsIn(C_Variable) == CTRUE) || 
            ((C_string.Id() == self.Isa.Id()) || 
              ((self == CNULL) || 
                ((Equal(self,CNIL.Id()) == CTRUE) || 
                  ((Equal(self,CEMPTY.Id()) == CTRUE) || 
                    (self.Isa.IsIn(Core.C_global_variable) == CTRUE)))))))) {return CTRUE
    } else {return CFALSE}} 
  
// The EID go function for: constant? @ any (throw: false) 
func E_Generate_constant_ask_any (self EID) EID { 
    return EID{F_Generate_constant_ask_any(ANY(self) ).Id(),0}} 
  
// patch: remove protection and conversion layers
/* {1} The go function for: getC(x:any) [status=0] */
func F_Generate_getC_any (x *ClaireAny ) *ClaireAny  { 
    // procedure body with s = any 
var Result *ClaireAny  
    if (x.Isa.IsIn(Core.C_global_variable) == CTRUE) { 
      { var g0073 *Core.GlobalVariable   = Core.ToGlobalVariable(x)
        if (Optimize.F_Compile_nativeVar_ask_global_variable(g0073) == CTRUE) { 
          Result = g0073.Id()
          } else {
          Result = ANY(Core.F_CALL(C_Generate_to_C,ARGS(EID{Equal(C_arg.Id(),g0073.Id()).Id(),0},EID{Equal(Language.C_iClaire_set_arg.Id(),C_type.Id()).Id(),0})))
          } 
        } 
      } else {
      Result = x
      } 
    return Result} 
  
// The EID go function for: getC @ any (throw: false) 
func E_Generate_getC_any (x EID) EID { 
    return F_Generate_getC_any(ANY(x) ).ToEID()} 
  
// short cut for variable
/* {1} The go function for: go_range(v:Variable) [status=0] */
func F_Generate_go_range_Variable (v *ClaireVariable ) *ClaireClass  { 
    return  v.Range.Class_I()
    } 
  
// The EID go function for: go_range @ Variable (throw: false) 
func E_Generate_go_range_Variable (v EID) EID { 
    return EID{F_Generate_go_range_Variable(To_Variable(OBJ(v)) ).Id(),0}} 
  
// in claire 4, srange(m:method) is gone, replaced by signature => this is temporary method
/* {1} The go function for: go_signature(m:method) [status=0] */
func F_Generate_go_signature_method (m *ClaireMethod ) *ClaireList  { 
    // procedure body with s = list 
var Result *ClaireList  
    { var arg_1 *ClaireList  
      _ = arg_1
      { 
        var v_list3 *ClaireList  
        var t *ClaireTypeExpression  
        var v_local3 *ClaireAny  
        v_list3 = m.Domain
        arg_1 = CreateList(ToType(C_class.Id()),v_list3.Length())
        for CLcount := 0; CLcount < v_list3.Length(); CLcount++{ 
          t = ToTypeExpression(v_list3.At(CLcount))
          v_local3 = t.Class_I().Id()
          arg_1.PutAt(CLcount,v_local3)
          } 
        } 
      Result = arg_1.AddFast(m.Range.Class_I().Id())/*t=any,s=list*/
      } 
    return Result} 
  
// The EID go function for: go_signature @ method (throw: false) 
func E_Generate_go_signature_method (m EID) EID { 
    return EID{F_Generate_go_signature_method(ToMethod(OBJ(m)) ).Id(),0}} 
  
// probably should exist elsewhere
/* {1} The go function for: full_signature(m:method) [status=0] */
func F_Generate_full_signature_method (m *ClaireMethod ) *ClaireList  { 
    // procedure body with s = list 
var Result *ClaireList  
    { var arg_1 *ClaireList  
      _ = arg_1
      { 
        var v_list3 *ClaireList  
        var t *ClaireTypeExpression  
        var v_local3 *ClaireAny  
        v_list3 = m.Domain
        arg_1 = CreateList(ToType(C_type.Id()),v_list3.Length())
        for CLcount := 0; CLcount < v_list3.Length(); CLcount++{ 
          t = ToTypeExpression(v_list3.At(CLcount))
          v_local3 = t.Id()
          arg_1.PutAt(CLcount,v_local3)
          } 
        } 
      Result = arg_1.AddFast(m.Range.Id())/*t=any,s=list*/
      } 
    return Result} 
  
// The EID go function for: full_signature @ method (throw: false) 
func E_Generate_full_signature_method (m EID) EID { 
    return EID{F_Generate_full_signature_method(ToMethod(OBJ(m)) ).Id(),0}} 
  
// print a signature in a AddMethod (goexp.cl)
/* {1} The go function for: signature!(c:go_producer,l:list<type>) [status=1] */
func (c *GenerateGoProducer ) Signature_I (l *ClaireList ) EID { 
    var Result EID 
    PRINC("Signature(")
    /*g_try(v2:"Result",loop:true) */
    { var arg_1 *ClaireList  
      _ = arg_1
      var try_2 EID 
      /*g_try(v2:"try_2",loop:false) */
      { 
        var v_list3 *ClaireList  
        var x *ClaireType  
        var v_local3 *ClaireAny  
        v_list3 = l
        try_2 = EID{CreateList(ToType(CEMPTY.Id()),v_list3.Length()).Id(),0}
        for CLcount := 0; CLcount < v_list3.Length(); CLcount++{ 
          x = ToType(v_list3.At(CLcount))
          var try_3 EID 
          /*g_try(v2:"try_3",loop:tuple("try_2", EID)) */
          try_3 = Core.F_CALL(Optimize.C_c_code,ARGS(EID{x.Id(),0},EID{C_type.Id(),0}))
          /* ERROR PROTECTION INSERTED (v_local3-try_2) */
          if ErrorIn(try_3) {try_2 = try_3
          break
          } else {
          v_local3 = ANY(try_3)
          ToList(OBJ(try_2)).PutAt(CLcount,v_local3)
          } 
        }
        } 
      /* ERROR PROTECTION INSERTED (arg_1-Result) */
      if ErrorIn(try_2) {Result = try_2
      } else {
      arg_1 = ToList(OBJ(try_2))
      Result = F_Generate_args_list_list(arg_1,C_any)
      }
      } 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(")")
    Result = EVOID
    }
    return Result} 
  
// The EID go function for: signature! @ go_producer (throw: true) 
func E_Generate_signature_I_go_producer (c EID,l EID) EID { 
    return ToGenerateGoProducer(OBJ(c)).Signature_I(ToList(OBJ(l)) )} 
  
// this is a specialized form for list expressions => see if Go should know if a ListObject, ListInt, ListFloat will be used versus generic List
/* {1} The go function for: g_member(x:any) [status=1] */
func F_Generate_g_member_any (x *ClaireAny ) EID { 
    var Result EID 
    if ((x.Isa.IsIn(Language.C_Call_method) == CTRUE) || 
        ((x.Isa.IsIn(Language.C_Construct) == CTRUE) || 
          ((x.Isa.IsIn(C_Variable) == CTRUE) || 
            ((x.Isa.IsIn(Language.C_Call_slot) == CTRUE) || 
              ((x.Isa.IsIn(Language.C_Cast) == CTRUE) || 
                (x.Isa.IsIn(Core.C_global_variable) == CTRUE)))))) { 
      { var t1 *ClaireType  
        var try_1 EID 
        /*g_try(v2:"try_1",loop:false) */
        { var arg_2 *ClaireType  
          _ = arg_2
          var try_3 EID 
          /*g_try(v2:"try_3",loop:false) */
          try_3 = Core.F_CALL(Optimize.C_c_type,ARGS(x.ToEID()))
          /* ERROR PROTECTION INSERTED (arg_2-try_1) */
          if ErrorIn(try_3) {try_1 = try_3
          } else {
          arg_2 = ToType(OBJ(try_3))
          try_1 = EID{arg_2.At(C_of).Id(),0}
          }
          } 
        /* ERROR PROTECTION INSERTED (t1-Result) */
        if ErrorIn(try_1) {Result = try_1
        } else {
        t1 = ToType(OBJ(try_1))
        if (Core.F_unique_ask_type(t1) == CTRUE) { 
          Result = Core.F_the_type(t1)
          } else {
          Result = EID{C_any.Id(),0}
          } 
        }
        } 
      } else {
      Result = EID{C_any.Id(),0}
      } 
    return Result} 
  
// The EID go function for: g_member @ any (throw: true) 
func E_Generate_g_member_any (x EID) EID { 
    return F_Generate_g_member_any(ANY(x) )} 
  
// this is a way to access the low-level native slices (for list and sets)
/* {1} The go function for: cast_Values(sbag:class,gmem:class) [status=0] */
func F_Generate_cast_Values_class (sbag *ClaireClass ,gmem *ClaireClass )  { 
    // procedure body with s = void 
{ var short *ClaireString  
      _ = short
      if (gmem.Id() == C_integer.Id()) { 
        short = MakeString("I")
        }  else if (gmem.Id() == C_float.Id()) { 
        short = MakeString("F")
        } else {
        short = MakeString("O")
        } 
      PRINC(".Values")
      F_princ_string(short)
      PRINC("()")
      } 
    } 
  
// The EID go function for: cast_Values @ class (throw: false) 
func E_Generate_cast_Values_class (sbag EID,gmem EID) EID { 
    F_Generate_cast_Values_class(ToClass(OBJ(sbag)),ToClass(OBJ(gmem)) )
    return EVOID} 
  
//
// this method does nothing. It used to check if a name could create a naming conflict.
// we keep it until we have tested that it is safe to remove it
// we could use a stack of names that have been used (reset for each method)
/* {1} The go function for: check_var(self:string) [status=0] */
func F_Generate_check_var_string (self *ClaireString ) *ClaireString  { 
    return  F_append_string(self,F_string_I_integer(Optimize.C_OPT.Level))
    } 
  
// The EID go function for: check_var @ string (throw: false) 
func E_Generate_check_var_string (self EID) EID { 
    return EID{F_Generate_check_var_string(ToString(OBJ(self)) ).Id(),0}} 
  
/* {1} The go function for: build_Variable(s:string,t:any) [status=0] */
func F_Generate_build_Variable_string (s *ClaireString ,t *ClaireAny ) *ClaireVariable  { 
    return  Optimize.F_Compile_Variable_I_symbol(Core.F_symbol_I_string2(s),0,t)
    } 
  
// The EID go function for: build_Variable @ string (throw: false) 
func E_Generate_build_Variable_string (s EID,t EID) EID { 
    return EID{F_Generate_build_Variable_string(ToString(OBJ(s)),ANY(t) ).Id(),0}} 
  
// use a variable v with inferred type when expected : add the casts
/* {1} The go function for: use_variable(v:string,expected:class,inferred:class) [status=1] */
func F_Generate_use_variable_string (v *ClaireString ,expected *ClaireClass ,inferred *ClaireClass ) EID { 
    var Result EID 
    /*g_try(v2:"Result",loop:true) */
    Result = F_Generate_cast_prefix_class(inferred,expected)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    F_c_princ_string(v)
    F_Generate_cast_post_class(inferred,expected)
    Result = EVOID
    }
    return Result} 
  
// The EID go function for: use_variable @ string (throw: true) 
func E_Generate_use_variable_string (v EID,expected EID,inferred EID) EID { 
    return F_Generate_use_variable_string(ToString(OBJ(v)),ToClass(OBJ(expected)),ToClass(OBJ(inferred)) )} 
  
// a clean expression is both a functional expression and one that does not throw an error
/* {1} The go function for: g_clean(x:any) [status=1] */
func F_Generate_g_clean_any (x *ClaireAny ) EID { 
    var Result EID 
    { 
      var v_and2 *ClaireBoolean  
      
      var try_1 EID 
      /*g_try(v2:"try_1",loop:false) */
      try_1 = F_Generate_g_func_any(x)
      /* ERROR PROTECTION INSERTED (v_and2-Result) */
      if ErrorIn(try_1) {Result = try_1
      } else {
      v_and2 = ToBoolean(OBJ(try_1))
      if (v_and2 == CFALSE) {Result = EID{CFALSE.Id(),0}
      } else { 
        var try_2 EID 
        /*g_try(v2:"try_2",loop:false) */
        { var arg_3 *ClaireBoolean  
          _ = arg_3
          var try_4 EID 
          /*g_try(v2:"try_4",loop:false) */
          try_4 = Optimize.F_Compile_g_throw_any(x)
          /* ERROR PROTECTION INSERTED (arg_3-try_2) */
          if ErrorIn(try_4) {try_2 = try_4
          } else {
          arg_3 = ToBoolean(OBJ(try_4))
          try_2 = EID{arg_3.Not.Id(),0}
          }
          } 
        /* ERROR PROTECTION INSERTED (v_and2-Result) */
        if ErrorIn(try_2) {Result = try_2
        } else {
        v_and2 = ToBoolean(OBJ(try_2))
        if (v_and2 == CFALSE) {Result = EID{CFALSE.Id(),0}
        } else { 
          Result = EID{CTRUE.Id(),0}} 
        } 
      }}
      } 
    return Result} 
  
// The EID go function for: g_clean @ any (throw: true) 
func E_Generate_g_clean_any (x EID) EID { 
    return F_Generate_g_clean_any(ANY(x) )} 
  
// a simple func expression that should not be left in go code
/* {1} The go function for: simple_func?(x:any) [status=1] */
func F_Generate_simple_func_ask_any (x *ClaireAny ) EID { 
    var Result EID 
    var g0075I *ClaireBoolean  
    var try_1 EID 
    /*g_try(v2:"try_1",loop:false) */
    { 
      var v_and2 *ClaireBoolean  
      
      var try_2 EID 
      /*g_try(v2:"try_2",loop:false) */
      try_2 = F_Generate_g_clean_any(x)
      /* ERROR PROTECTION INSERTED (v_and2-try_1) */
      if ErrorIn(try_2) {try_1 = try_2
      } else {
      v_and2 = ToBoolean(OBJ(try_2))
      if (v_and2 == CFALSE) {try_1 = EID{CFALSE.Id(),0}
      } else { 
        var try_3 EID 
        /*g_try(v2:"try_3",loop:false) */
        { var arg_4 *ClaireType  
          _ = arg_4
          var try_5 EID 
          /*g_try(v2:"try_5",loop:false) */
          try_5 = Core.F_CALL(Optimize.C_c_type,ARGS(x.ToEID()))
          /* ERROR PROTECTION INSERTED (arg_4-try_3) */
          if ErrorIn(try_5) {try_3 = try_5
          } else {
          arg_4 = ToType(OBJ(try_5))
          try_3 = EID{Core.F__I_equal_any(arg_4.Id(),C_void.Id()).Id(),0}
          }
          } 
        /* ERROR PROTECTION INSERTED (v_and2-try_1) */
        if ErrorIn(try_3) {try_1 = try_3
        } else {
        v_and2 = ToBoolean(OBJ(try_3))
        if (v_and2 == CFALSE) {try_1 = EID{CFALSE.Id(),0}
        } else { 
          try_1 = EID{CTRUE.Id(),0}} 
        } 
      }}
      } 
    /* ERROR PROTECTION INSERTED (g0075I-Result) */
    if ErrorIn(try_1) {Result = try_1
    } else {
    g0075I = ToBoolean(OBJ(try_1))
    if (g0075I == CTRUE) { 
      Result = EID{CTRUE.Id(),0}
      } else {
      Result = EID{CFALSE.Id(),0}
      } 
    }
    return Result} 
  
// The EID go function for: simple_func? @ any (throw: true) 
func E_Generate_simple_func_ask_any (x EID) EID { 
    return F_Generate_simple_func_ask_any(ANY(x) )} 
  