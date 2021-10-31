/***** CLAIRE Compilation of file /Users/ycaseau/Dropbox/src/clairev4.0/src/compile/gogen.cl 
         [version 4.0.01 / safety 5] Saturday 10-30-2021 *****/

package Generate
import (_ "fmt"
	. "Kernel"
	"Core"
	"Language"
	"Reader"
	"Optimize"
)

//-------- dumb function to prevent import errors --------
func import_g0204() { 
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
/* {1} OPT.The go function for: new_block(tag:string) [] */
func F_Generate_new_block_string (tag *ClaireString )  { 
    // procedure body with s =  
PRINC("/* ")
    F_princ_string(tag)
    PRINC(":")
    F_princ_integer(Optimize.C_OPT.Level)
    PRINC(" */")
    F_Generate_new_block_void()
    } 
  
// The EID go function for: new_block @ string (throw: false) 
func E_Generate_new_block_string (tag EID) EID { 
    /*(sm for new_block @ string= void)*/ F_Generate_new_block_string(ToString(OBJ(tag)) )
    return EVOID} 
  
/* {1} OPT.The go function for: close_block(tag:string) [] */
func F_Generate_close_block_string (tag *ClaireString )  { 
    // procedure body with s =  
PRINC("/* ")
    F_princ_string(tag)
    PRINC("-")
    F_princ_integer((Optimize.C_OPT.Level-1))
    PRINC(" */")
    F_Generate_close_block_void()
    } 
  
// The EID go function for: close_block @ string (throw: false) 
func E_Generate_close_block_string (tag EID) EID { 
    /*(sm for close_block @ string= void)*/ F_Generate_close_block_string(ToString(OBJ(tag)) )
    return EVOID} 
  
/* {1} OPT.The go function for: finish_block(tag:string) [] */
func F_Generate_finish_block_string (tag *ClaireString )  { 
    // procedure body with s =  
PRINC("/* ")
    F_princ_string(tag)
    PRINC("!")
    F_princ_integer((Optimize.C_OPT.Level-1))
    PRINC(" */")
    F_Generate_finish_block_void()
    } 
  
// The EID go function for: finish_block @ string (throw: false) 
func E_Generate_finish_block_string (tag EID) EID { 
    /*(sm for finish_block @ string= void)*/ F_Generate_finish_block_string(ToString(OBJ(tag)) )
    return EVOID} 
  
// *******************************************************************
// *       Part 1: definition of the code producer               *
// *******************************************************************
// definition of the instance
// to do : update the reserved names progressively - note that classes do not need protection since
// ClaireX is added to X
// use this producer
// makes an ident (string) from a variable's name - we keep the ofuscated option (.naming = 2 )
/* {1} OPT.The go function for: c_string(c:go_producer,self:Variable) [] */
func F_Generate_c_string_go_producer1 (c *GenerateGoProducer ,self *ClaireVariable ) EID { 
    var Result EID 
    
    if (Optimize.C_compiler.Naming == 2) /* If:2 */{ 
      /* Let:3 */{ 
        var g0205UU *ClaireAny  
        /* noccur = 1 */
        var g0205UU_try02064 EID 
        /* Let:4 */{ 
          var g0207UU *ClaireAny  
          /* noccur = 1 */
          var g0207UU_try02085 EID 
          g0207UU_try02085 = Core.F_CALL(C_integer_I,ARGS(EID{self.Pname.Id(),0}))
          /* ERROR PROTECTION INSERTED (g0207UU-g0205UU_try02064) */
          if ErrorIn(g0207UU_try02085) {g0205UU_try02064 = g0207UU_try02085
          } else {
          g0207UU = ANY(g0207UU_try02085)
          g0205UU_try02064 = Core.F_CALL(C_string_I,ARGS(g0207UU.ToEID()))
          }
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (g0205UU-Result) */
        if ErrorIn(g0205UU_try02064) {Result = g0205UU_try02064
        } else {
        g0205UU = ANY(g0205UU_try02064)
        Result = EID{F_append_string(MakeString("v"),ToString(g0205UU)).Id(),0}
        }
        /* Let-3 */} 
      } else {
      Core.F_print_in_string_void()
      F_iClaire_ident_go_producer2(c,self.Pname)
      Result = Core.F_end_of_string_void()
      /* If-2 */} 
    return Result} 
  
// The EID go function for: c_string @ list<type_expression>(go_producer, Variable) (throw: true) 
func E_Generate_c_string_go_producer1 (c EID,self EID) EID { 
    return /*(sm for c_string @ list<type_expression>(go_producer, Variable)= EID)*/ F_Generate_c_string_go_producer1(ToGenerateGoProducer(OBJ(c)),To_Variable(OBJ(self)) )} 
  
/* {1} OPT.The go function for: c_string(c:go_producer,self:symbol) [] */
func F_Generate_c_string_go_producer2 (c *GenerateGoProducer ,self *ClaireSymbol ) EID { 
    var Result EID 
    Core.F_print_in_string_void()
    F_iClaire_ident_go_producer2(c,self)
    Result = Core.F_end_of_string_void()
    return Result} 
  
// The EID go function for: c_string @ list<type_expression>(go_producer, symbol) (throw: true) 
func E_Generate_c_string_go_producer2 (c EID,self EID) EID { 
    return /*(sm for c_string @ list<type_expression>(go_producer, symbol)= EID)*/ F_Generate_c_string_go_producer2(ToGenerateGoProducer(OBJ(c)),ToSymbol(OBJ(self)) )} 
  
// print a symbol for a variable
// two issues : replace with a dictionary some day (CLAIRE4) + why does c_string exist ?
// notice that ident should only exist for <strings> that will exist directly in Go code 
/* {1} OPT.The go function for: iClaire/ident(c:go_producer,v:Variable) [] */
func F_iClaire_ident_go_producer1 (c *GenerateGoProducer ,v *ClaireVariable ) EID { 
    var Result EID 
    if (Optimize.C_compiler.Naming == 2) /* If:2 */{ 
      /* Let:3 */{ 
        var g0209UU *ClaireString  
        /* noccur = 1 */
        var g0209UU_try02104 EID 
        /* Let:4 */{ 
          var g0211UU *ClaireAny  
          /* noccur = 1 */
          var g0211UU_try02125 EID 
          /* Let:5 */{ 
            var g0213UU *ClaireAny  
            /* noccur = 1 */
            var g0213UU_try02146 EID 
            g0213UU_try02146 = Core.F_CALL(C_integer_I,ARGS(EID{v.Pname.Id(),0}))
            /* ERROR PROTECTION INSERTED (g0213UU-g0211UU_try02125) */
            if ErrorIn(g0213UU_try02146) {g0211UU_try02125 = g0213UU_try02146
            } else {
            g0213UU = ANY(g0213UU_try02146)
            g0211UU_try02125 = Core.F_CALL(C_string_I,ARGS(g0213UU.ToEID()))
            }
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (g0211UU-g0209UU_try02104) */
          if ErrorIn(g0211UU_try02125) {g0209UU_try02104 = g0211UU_try02125
          } else {
          g0211UU = ANY(g0211UU_try02125)
          g0209UU_try02104 = EID{F_append_string(MakeString("v"),ToString(g0211UU)).Id(),0}
          }
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (g0209UU-Result) */
        if ErrorIn(g0209UU_try02104) {Result = g0209UU_try02104
        } else {
        g0209UU = ToString(OBJ(g0209UU_try02104))
        F_princ_string(g0209UU)
        Result = EVOID
        }
        /* Let-3 */} 
      } else {
      /* Let:3 */{ 
        var s *ClaireSymbol   = v.Pname
        /* noccur = 2 */
        /* Let:4 */{ 
          var n int  = F_index_list(c.BadNames,s.Id())
          /* noccur = 2 */
          if (n == 0) /* If:5 */{ 
            F_c_princ_string(s.String_I())
            Result = EVOID
            } else {
            ToSymbol(c.GoodNames.At(n-1)).CPrinc()
            Result = EVOID
            /* If-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* If-2 */} 
    return Result} 
  
// The EID go function for: iClaire/ident @ list<type_expression>(go_producer, Variable) (throw: true) 
func E_iClaire_ident_go_producer1 (c EID,v EID) EID { 
    return /*(sm for iClaire/ident @ list<type_expression>(go_producer, Variable)= EID)*/ F_iClaire_ident_go_producer1(ToGenerateGoProducer(OBJ(c)),To_Variable(OBJ(v)) )} 
  
// print a symbol for the structure definition  => use c_princ to get rid of special chars
/* {1} OPT.The go function for: iClaire/ident(c:go_producer,s:symbol) [] */
func F_iClaire_ident_go_producer2 (c *GenerateGoProducer ,s *ClaireSymbol )  { 
    // procedure body with s =  
/* Let:2 */{ 
      var n int  = F_index_list(c.BadNames,s.Id())
      /* noccur = 2 */
      if (n == 0) /* If:3 */{ 
        F_c_princ_string(s.String_I())
        } else {
        ToSymbol(c.GoodNames.At(n-1)).CPrinc()
        /* If-3 */} 
      /* Let-2 */} 
    } 
  
// The EID go function for: iClaire/ident @ list<type_expression>(go_producer, symbol) (throw: false) 
func E_iClaire_ident_go_producer2 (c EID,s EID) EID { 
    /*(sm for iClaire/ident @ list<type_expression>(go_producer, symbol)= void)*/ F_iClaire_ident_go_producer2(ToGenerateGoProducer(OBJ(c)),ToSymbol(OBJ(s)) )
    return EVOID} 
  
// new in claire4: printd the go identifier asociated with symbol s
// cap_ident(c,x) uses capitalization : used for Class and Method, required by Go for identifiers to be visible
// notice that we print explicitly s.module! (namespace) if not claire, to avoid c name conflicts
/* {1} OPT.The go function for: cap_ident(s:symbol) [] */
func F_Generate_cap_ident_symbol (s *ClaireSymbol )  { 
    // procedure body with s =  
F_Generate_capitalized_ident_symbol(s,s.Module_I())
    } 
  
// The EID go function for: cap_ident @ symbol (throw: false) 
func E_Generate_cap_ident_symbol (s EID) EID { 
    /*(sm for cap_ident @ symbol= void)*/ F_Generate_cap_ident_symbol(ToSymbol(OBJ(s)) )
    return EVOID} 
  
// this is the capitalized ident for s in namespace m
/* {1} OPT.The go function for: capitalized_ident(s:symbol,m:module) [] */
func F_Generate_capitalized_ident_symbol (s *ClaireSymbol ,m *ClaireModule )  { 
    // procedure body with s =  
/* Let:2 */{ 
      var n int  = F_index_list(ToGenerateGoProducer(Optimize.C_PRODUCER.Value).BadNames,s.Id())
      /* noccur = 2 */
      if (n == 0) /* If:3 */{ 
        if (m.Id() != C_claire.Id()) /* If:4 */{ 
          F_c_princ_string(F_Generate_capitalize_string(m.Name.String_I()))
          /* If-4 */} 
        F_c_princ_string(F_Generate_capitalize_string(s.String_I()))
        } else {
        ToSymbol(ToGenerateGoProducer(Optimize.C_PRODUCER.Value).GoodNames.At(n-1)).CPrinc()
        /* If-3 */} 
      /* Let-2 */} 
    } 
  
// The EID go function for: capitalized_ident @ symbol (throw: false) 
func E_Generate_capitalized_ident_symbol (s EID,m EID) EID { 
    /*(sm for capitalized_ident @ symbol= void)*/ F_Generate_capitalized_ident_symbol(ToSymbol(OBJ(s)),ToModule(OBJ(m)) )
    return EVOID} 
  
// short version (we do not care about the namespace) 
/* {1} OPT.The go function for: cap_short(s:symbol) [] */
func F_Generate_cap_short_symbol (s *ClaireSymbol )  { 
    // procedure body with s =  
F_Generate_capitalized_ident_symbol(s,C_claire)
    } 
  
// The EID go function for: cap_short @ symbol (throw: false) 
func E_Generate_cap_short_symbol (s EID) EID { 
    /*(sm for cap_short @ symbol= void)*/ F_Generate_cap_short_symbol(ToSymbol(OBJ(s)) )
    return EVOID} 
  
// CLAIRE 4 NEW ! a class name is printed with the module identifier
// go_class is the the go name ModuleClass 
// class_ident => thing_ident is the name of the global variable that contains the CLAIRE object 
/* {1} OPT.The go function for: go_class(self:class) [] */
func F_Generate_go_class_class (self *ClaireClass )  { 
    // procedure body with s =  
/* Let:2 */{ 
      var m *ClaireModule   = self.Name.Defined()
      /* noccur = 3 */
      if (m.Id() == C_Kernel.Id()) /* If:3 */{ 
        PRINC("Claire")
        /* If!3 */}  else if (m.Id() != ToGenerateGoProducer(Optimize.C_PRODUCER.Value).Current.Id()) /* If:3 */{ 
        F_Generate_cap_ident_symbol(m.Name)
        PRINC(".")
        /* If-3 */} 
      if ((self.Id() == C_array.Id()) || 
          (self.Id() == C_listargs.Id())) /* If:3 */{ 
        F_c_princ_string(MakeString("List"))
        } else {
        F_Generate_cap_ident_symbol(self.Name)
        /* If-3 */} 
      /* Let-2 */} 
    } 
  
// The EID go function for: go_class @ class (throw: false) 
func E_Generate_go_class_class (self EID) EID { 
    /*(sm for go_class @ class= void)*/ F_Generate_go_class_class(ToClass(OBJ(self)) )
    return EVOID} 
  
// small subtlety : in cast names (ToX) we ommit the "Claire" for simplicity
/* {1} OPT.The go function for: cast_class(self:class) [] */
func F_Generate_cast_class_class (self *ClaireClass )  { 
    // procedure body with s =  
/* Let:2 */{ 
      var m *ClaireModule   = self.Name.Defined()
      /* noccur = 3 */
      if ((m.Id() != C_Kernel.Id()) && 
          (m.Id() != ToGenerateGoProducer(Optimize.C_PRODUCER.Value).Current.Id())) /* If:3 */{ 
        F_Generate_cap_ident_symbol(m.Name)
        PRINC(".")
        /* If-3 */} 
      PRINC("To")
      F_Generate_addUnderscore_class(self)
      if (self.Id() == C_listargs.Id()) /* If:3 */{ 
        PRINC("List")
        } else {
        F_Generate_cap_ident_symbol(self.Name)
        /* If-3 */} 
      /* Let-2 */} 
    } 
  
// The EID go function for: cast_class @ class (throw: false) 
func E_Generate_cast_class_class (self EID) EID { 
    /*(sm for cast_class @ class= void)*/ F_Generate_cast_class_class(ToClass(OBJ(self)) )
    return EVOID} 
  
// class_ident(c) = C_c
/* {1} OPT.The go function for: class_ident(self:class) [] */
func F_Generate_class_ident_class (self *ClaireClass )  { 
    // procedure body with s =  
F_Generate_symbol_ident_symbol(self.Name)
    } 
  
// The EID go function for: class_ident @ class (throw: false) 
func E_Generate_class_ident_class (self EID) EID { 
    /*(sm for class_ident @ class= void)*/ F_Generate_class_ident_class(ToClass(OBJ(self)) )
    return EVOID} 
  
// same : remember that a class is  not a thing
/* {1} OPT.The go function for: thing_ident(self:thing) [] */
func F_Generate_thing_ident_thing (self *ClaireThing )  { 
    // procedure body with s =  
F_Generate_symbol_ident_symbol(self.Name)
    } 
  
// The EID go function for: thing_ident @ thing (throw: false) 
func E_Generate_thing_ident_thing (self EID) EID { 
    /*(sm for thing_ident @ thing= void)*/ F_Generate_thing_ident_thing(ToThing(OBJ(self)) )
    return EVOID} 
  
// how a named object is designated in go (through a global variable from the package = module). 
// CLAIRE v4: No prefix needed for current or Kernel
/* {1} OPT.The go function for: symbol_ident(s:symbol) [] */
func F_Generate_symbol_ident_symbol (s *ClaireSymbol )  { 
    // procedure body with s =  
/* Let:2 */{ 
      var m *ClaireModule   = s.Defined()
      /* noccur = 4 */
      if ((m.Id() != ToGenerateGoProducer(Optimize.C_PRODUCER.Value).Current.Id()) && 
          ((m.Id() != C_Kernel.Id()) && 
            (m.Id() != C_claire.Id()))) /* If:3 */{ 
        F_Generate_cap_short_symbol(m.Name)
        PRINC(".")
        /* If-3 */} 
      F_Generate_go_var_symbol(s)
      /* Let-2 */} 
    } 
  
// The EID go function for: symbol_ident @ symbol (throw: false) 
func E_Generate_symbol_ident_symbol (s EID) EID { 
    /*(sm for symbol_ident @ symbol= void)*/ F_Generate_symbol_ident_symbol(ToSymbol(OBJ(s)) )
    return EVOID} 
  
// this produced the C_s identifier which are go global variables, 
// all compiler code should use this (get rid of C_ in code)
/* {1} OPT.The go function for: go_var(s:symbol) [] */
func F_Generate_go_var_symbol (s *ClaireSymbol )  { 
    // procedure body with s =  
/* Let:2 */{ 
      var m *ClaireModule   = s.Module_I()
      /* noccur = 2 */
      PRINC("C_")
      if (m.Id() != C_claire.Id()) /* If:3 */{ 
        F_c_princ_string(m.Name.String_I())
        F_c_princ_string(MakeString("_"))
        /* If-3 */} 
      F_c_princ_string(s.String_I())
      /* Let-2 */} 
    } 
  
// The EID go function for: go_var @ symbol (throw: false) 
func E_Generate_go_var_symbol (s EID) EID { 
    /*(sm for go_var @ symbol= void)*/ F_Generate_go_var_symbol(ToSymbol(OBJ(s)) )
    return EVOID} 
  
// when we capitalize the name of class, we may create a conflict (list vs List)
/* {1} OPT.The go function for: addUnderscore(c:class) [] */
func F_Generate_addUnderscore_class (c *ClaireClass )  { 
    // procedure body with s =  
/* Let:2 */{ 
      var s *ClaireString   = c.Name.String_I()
      /* noccur = 2 */
      if ((F_integer_I_char(s.At(1)) >= 65) && 
          (F_integer_I_char(s.At(1)) <= 90)) /* If:3 */{ 
        PRINC("_")
        /* If-3 */} 
      /* Let-2 */} 
    } 
  
// The EID go function for: addUnderscore @ class (throw: false) 
func E_Generate_addUnderscore_class (c EID) EID { 
    /*(sm for addUnderscore @ class= void)*/ F_Generate_addUnderscore_class(ToClass(OBJ(c)) )
    return EVOID} 
  
// the Go code producer uses Capitalization as a strategy for name generation
//  capitalize(s)  => capitalize the first letter + search for _, remove and capitalize next letter
//  capitalize("foo_bar") = "FooBar"
/* {1} OPT.The go function for: capitalize(s:string) [] */
func F_Generate_capitalize_string (s *ClaireString ) *ClaireString  { 
    // procedure body with s =  
var Result *ClaireString  
    /* Let:2 */{ 
      var n int  = F_length_string(s)
      /* noccur = 1 */
      /* Let:3 */{ 
        var i int  = F_get_string(s,'_')
        /* noccur = 3 */
        if (i == 0) /* If:4 */{ 
          /* Let:5 */{ 
            var s2 *ClaireString   = F_copy_string(s)
            /* noccur = 2 */
            F_nth_set_string(s2,1,F_Generate_capitalize_char(s.At(1)))
            Result = s2
            /* Let-5 */} 
          } else {
          Result = F_append_string(F_Generate_capitalize_string(F_substring_string(s,1,(i-1))),F_Generate_capitalize_string(F_substring_string(s,(i+1),n)))
          /* If-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: capitalize @ string (throw: false) 
func E_Generate_capitalize_string (s EID) EID { 
    return EID{/*(sm for capitalize @ string= string)*/ F_Generate_capitalize_string(ToString(OBJ(s)) ).Id(),0}} 
  
/* {1} OPT.The go function for: capitalize(c:char) [] */
func F_Generate_capitalize_char (c rune) rune { 
    // procedure body with s =  
var Result rune 
    /* Let:2 */{ 
      var i int  = int(c)
      /* noccur = 3 */
      if ((i >= 97) && 
          (i <= 122)) /* If:3 */{ 
        Result = F_char_I_integer((i-32))
        } else {
        Result = c
        /* If-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: capitalize @ char (throw: false) 
func E_Generate_capitalize_char (c EID) EID { 
    return EID{C__CHAR,CVAL(/*(sm for capitalize @ char= char)*/ F_Generate_capitalize_char(CHAR(c) ))}} 
  
/* {1} OPT.The go function for: capitalize(s:symbol) [] */
func F_Generate_capitalize_symbol (s *ClaireSymbol ) *ClaireString  { 
    // use function body compiling 
return  F_Generate_capitalize_string(s.String_I())
    } 
  
// The EID go function for: capitalize @ symbol (throw: false) 
func E_Generate_capitalize_symbol (s EID) EID { 
    return EID{/*(sm for capitalize @ symbol= string)*/ F_Generate_capitalize_symbol(ToSymbol(OBJ(s)) ).Id(),0}} 
  
// v3.3 : new ! a global variable contains the native value
// range = {} for global constant
/* {1} OPT.The go function for: globalVar(c:go_producer,x:global_variable) [] */
func (c *GenerateGoProducer ) GlobalVar (x *Core.GlobalVariable )  { 
    // procedure body with s =  
F_Generate_thing_ident_thing(ToThing(x.Id()))
    PRINC(".Value")
    } 
  
// The EID go function for: globalVar @ go_producer (throw: false) 
func E_Generate_globalVar_go_producer (c EID,x EID) EID { 
    /*(sm for globalVar @ go_producer= void)*/ ToGenerateGoProducer(OBJ(c)).GlobalVar(Core.ToGlobalVariable(OBJ(x)) )
    return EVOID} 
  
// the go expression that represents a global variable, as a string (reused for Gassign)
// Five sorts in go : categories to distinguish between native, object, EID
//    x:object       x,    x,       EID(x.Id(),0)
//    x:int,float,char         x,    MakeX(x),      EID{C__C,xVAL(x)}
//    x:exception     x,  x,      EID{x,1}
//  notice that Boolean is a an object but it could be handled with a native form in the future
/* {1} OPT.The go function for: type_sort(x:type) [] */
func F_Generate_type_sort_type (x *ClaireType ) *ClaireClass  { 
    // procedure body with s =  
var Result *ClaireClass  
    /* Let:2 */{ 
      var c *ClaireClass   = x.Class_I()
      /* noccur = 5 */
      if ((c.Id() == C_float.Id()) || 
          ((c.Id() == C_integer.Id()) || 
            ((c.Id() == C_char.Id()) || 
              (c.Id() == Optimize.C_EID.Id())))) /* If:3 */{ 
        Result = c
        } else {
        Result = C_any
        /* If-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: type_sort @ type (throw: false) 
func E_Generate_type_sort_type (x EID) EID { 
    return EID{/*(sm for type_sort @ type= class)*/ F_Generate_type_sort_type(ToType(OBJ(x)) ).Id(),0}} 
  
// sorts in go are much simpler : int, float, any or EID
/* {1} OPT.The go function for: g_sort(x:any) [] */
func F_Generate_g_sort_any (x *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var g0215UU *ClaireClass  
      /* noccur = 1 */
      var g0215UU_try02163 EID 
      g0215UU_try02163 = Language.F_static_type_any(x)
      /* ERROR PROTECTION INSERTED (g0215UU-Result) */
      if ErrorIn(g0215UU_try02163) {Result = g0215UU_try02163
      } else {
      g0215UU = ToClass(OBJ(g0215UU_try02163))
      Result = EID{F_Generate_type_sort_type(ToType(g0215UU.Id())).Id(),0}
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: g_sort @ any (throw: true) 
func E_Generate_g_sort_any (x EID) EID { 
    return /*(sm for g_sort @ any= EID)*/ F_Generate_g_sort_any(ANY(x) )} 
  
// access the proper values slot for a list whose member type s is determined (not any)
/* {1} OPT.The go function for: valuesSlot(s:class) [] */
func F_Generate_valuesSlot_class (s *ClaireClass )  { 
    // procedure body with s =  
PRINC("Values")
    /* Let:2 */{ 
      var g0217UU *ClaireString  
      /* noccur = 1 */
      if (s.Id() == C_integer.Id()) /* If:3 */{ 
        g0217UU = MakeString("I")
        /* If!3 */}  else if (s.Id() == C_float.Id()) /* If:3 */{ 
        g0217UU = MakeString("F")
        } else {
        g0217UU = MakeString("O")
        /* If-3 */} 
      F_princ_string(g0217UU)
      /* Let-2 */} 
    PRINC("()")
    } 
  
// The EID go function for: valuesSlot @ class (throw: false) 
func E_Generate_valuesSlot_class (s EID) EID { 
    /*(sm for valuesSlot @ class= void)*/ F_Generate_valuesSlot_class(ToClass(OBJ(s)) )
    return EVOID} 
  
// *******************************************************************
// *       Part 2: utilities for file generation                     *
// *******************************************************************
// note that all code to produce interfaces is gone :)
// a module is simply the combination of (1) a Go package (2) a ClaireModule (3) a load function
// generate a namespace definition (Go package)
/* {1} OPT.The go function for: namespace!(c:go_producer,m:module) [] */
func (c *GenerateGoProducer ) Namespace_I (m *ClaireModule )  { 
    // procedure body with s =  
PRINC("package ")
    F_iClaire_ident_symbol(m.Name)
    PRINC("\n")
    } 
  
// The EID go function for: namespace! @ go_producer (throw: false) 
func E_Generate_namespace_I_go_producer (c EID,m EID) EID { 
    /*(sm for namespace! @ go_producer= void)*/ ToGenerateGoProducer(OBJ(c)).Namespace_I(ToModule(OBJ(m)) )
    return EVOID} 
  
// note : we have removed module!(c:go_producer,m:module) => nothing to add to the previous line
// define a new typed variable named v (called in go_stat.cl )
// short cut : var declaration without an initialisation + a breakline
// CRAZY: go compiler gets confused with some variables not being used .. the dump forces to issue a dumb
// statement to get rid of this
// mode : 0 : normal no newline, 1 : newline, 2: special
/* {1} OPT.The go function for: var_declaration(v:string,s:class,mode:integer) [] */
func F_Generate_var_declaration_string (v *ClaireString ,s *ClaireClass ,mode int)  { 
    // procedure body with s =  
PRINC("var ")
    F_c_princ_string(v)
    PRINC(" ")
    F_Generate_interface_I_class(s)
    PRINC(" ")
    if (mode > 0) /* If:2 */{ 
      F_Generate_breakline_void()
      /* If-2 */} 
    PRINC("")
    if (mode == 2) /* If:2 */{ 
      PRINC("_ = ")
      F_c_princ_string(v)
      F_Generate_breakline_void()
      PRINC("")
      /* If-2 */} 
    } 
  
// The EID go function for: var_declaration @ string (throw: false) 
func E_Generate_var_declaration_string (v EID,s EID,mode EID) EID { 
    /*(sm for var_declaration @ string= void)*/ F_Generate_var_declaration_string(ToString(OBJ(v)),ToClass(OBJ(s)),INT(mode) )
    return EVOID} 
  
// ! is a semantic marker for imported
/* {1} OPT.The go function for: imported_function?(f:any) [] */
func F_imported_function_ask_any (f *ClaireAny ) *ClaireBoolean  { 
    // procedure body with s =  
var Result *ClaireBoolean  
    if (C_function.Id() == f.Isa.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var g0218 *ClaireFunction   = ToFunction(f)
        /* noccur = 1 */
        Result = Equal(MakeChar(F_string_I_function(g0218).At(1)).Id(),MakeChar('#').Id())
        /* Let-3 */} 
      } else {
      Result = CFALSE
      /* If-2 */} 
    return Result} 
  
// The EID go function for: imported_function? @ any (throw: false) 
func E_imported_function_ask_any (f EID) EID { 
    return EID{/*(sm for imported_function? @ any= boolean)*/ F_imported_function_ask_any(ANY(f) ).Id(),0}} 
  
// tells if we can compile the CLAIRE method as a go method or if we shoud use a function
// we use the go method if the class is defined in the same 
// remember that Go does not support polymorphism on parameters : we can use a method only if there is one match 
// based on first argument - howver this restriction is package based (to be checked)
// we first check that the first char of the name is a proper letter
// also methods defined with #'#foo are forced to use foo :)
/* {1} OPT.The go function for: goMethod?(m:any) [] */
func F_Generate_goMethod_ask_any (m *ClaireAny ) *ClaireBoolean  { 
    // procedure body with s =  
var Result *ClaireBoolean  
    if (ToBoolean(ToGenerateGoProducer(Optimize.C_PRODUCER.Value).KernelMethods.Contain_ask(m).Id()) == CTRUE) /* If:2 */{ 
      Result = CTRUE
      /* If!2 */}  else if (C_method.Id() == m.Isa.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var g0220 *ClaireMethod   = ToMethod(m)
        /* noccur = 7 */
        /* Let:4 */{ 
          var firstc rune  = g0220.Selector.Name.String_I().At(1)
          /* noccur = 2 */
          /* Let:5 */{ 
            var _Zsig *ClaireList   = F_Generate_go_signature_method(g0220)
            /* noccur = 2 */
            /* Let:6 */{ 
              var c *ClaireAny   = _Zsig.ValuesO()[1-1]
              /* noccur = 5 */
              /* and:7 */{ 
                var v_and7 *ClaireBoolean  
                
                v_and7 = Core.F__inf_equal_char('A',firstc)
                if (v_and7 == CFALSE) {Result = CFALSE
                } else /* arg:8 */{ 
                  v_and7 = Core.F__inf_equal_char(firstc,'z')
                  if (v_and7 == CFALSE) {Result = CFALSE
                  } else /* arg:9 */{ 
                    v_and7 = MakeBoolean((ToType(c).Included(ToType(C_object.Id())) == CTRUE) || (c == C_port.Id()) || (c == C_environment.Id()))
                    if (v_and7 == CFALSE) {Result = CFALSE
                    } else /* arg:10 */{ 
                      v_and7 = Equal(g0220.Module_I.Id(),ToClass(c).Name.Defined().Id())
                      if (v_and7 == CFALSE) {Result = CFALSE
                      } else /* arg:11 */{ 
                        v_and7 = MakeBoolean((g0220.Functional.Id() == CNULL) || (F_imported_function_ask_any(g0220.Functional.Id()) != CTRUE))
                        if (v_and7 == CFALSE) {Result = CFALSE
                        } else /* arg:12 */{ 
                          /* Let:13 */{ 
                            var g0224UU *ClaireAny  
                            /* noccur = 1 */
                            /* For:14 */{ 
                              var m2 *ClaireAny  
                              _ = m2
                              g0224UU= CFALSE.Id()
                              for _,m2 = range(g0220.Selector.Restrictions.ValuesO())/* loop:15 */{ 
                                var g0225I *ClaireBoolean  
                                /* Let:16 */{ 
                                  var g0226UU *ClaireBoolean  
                                  /* noccur = 1 */
                                  if (C_method.Id() == m2.Isa.Id()) /* If:17 */{ 
                                    /* Let:18 */{ 
                                      var g0221 *ClaireMethod   = ToMethod(m2)
                                      /* noccur = 3 */
                                      if ((g0221.Module_I.Id() == g0220.Module_I.Id()) && 
                                          (Equal(Core.F__exp_type(ToType(c),ToType(g0221.Domain.ValuesO()[1-1])).Id(),CEMPTY.Id()) != CTRUE)) /* If:19 */{ 
                                        g0226UU = F_Generate_arg_match_list(F_Generate_go_signature_method(g0221),_Zsig)
                                        } else {
                                        g0226UU = CTRUE
                                        /* If-19 */} 
                                      /* Let-18 */} 
                                    } else {
                                    g0226UU = CTRUE
                                    /* If-17 */} 
                                  g0225I = g0226UU.Not
                                  /* Let-16 */} 
                                if (g0225I == CTRUE) /* If:16 */{ 
                                   /*v = g0224UU, s =any*/
g0224UU = CTRUE.Id()
                                  break
                                  /* If-16 */} 
                                /* loop-15 */} 
                              /* For-14 */} 
                            v_and7 = Core.F_not_any(g0224UU)
                            /* Let-13 */} 
                          if (v_and7 == CFALSE) {Result = CFALSE
                          } else /* arg:13 */{ 
                            Result = CTRUE/* arg-13 */} 
                          /* arg-12 */} 
                        /* arg-11 */} 
                      /* arg-10 */} 
                    /* arg-9 */} 
                  /* arg-8 */} 
                /* and-7 */} 
              /* Let-6 */} 
            /* Let-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      } else {
      Result = CFALSE
      /* If-2 */} 
    return Result} 
  
// The EID go function for: goMethod? @ any (throw: false) 
func E_Generate_goMethod_ask_any (m EID) EID { 
    return EID{/*(sm for goMethod? @ any= boolean)*/ F_Generate_goMethod_ask_any(ANY(m) ).Id(),0}} 
  
// useful for debug - notice that a CLAIRE method defined on a class which is NOT in the same module
// is always compiled as a function
/* {1} OPT.The go function for: dMethod?(m:any) [] */
func F_dMethod_ask_any (m *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var firstc rune  = ToSymbol(OBJ(Core.F_CALL(C_name,ARGS(Core.F_CALL(C_selector,ARGS(m.ToEID())))))).String_I().At(1)
      /* noccur = 2 */
      /* Let:3 */{ 
        var _Zsig *ClaireList   = F_Generate_go_signature_method(ToMethod(m))
        /* noccur = 2 */
        /* Let:4 */{ 
          var c *ClaireAny   = _Zsig.ValuesO()[1-1]
          /* noccur = 5 */
          PRINC("char -> ")
          Result = Core.F_print_any(MakeBoolean((Core.F__inf_equal_char('A',firstc) == CTRUE) && (Core.F__inf_equal_char(firstc,'z') == CTRUE)).Id())
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC("\n")
          Result = EVOID
          }
          {
          PRINC("hierarchy -> ")
          Result = Core.F_print_any(MakeBoolean((ToType(c).Included(ToType(C_object.Id())) == CTRUE) || (c == C_port.Id()) || (c == C_environment.Id())).Id())
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC("\n")
          Result = EVOID
          }
          {
          PRINC("module [")
          Result = Core.F_print_any(ANY(Core.F_CALL(C_module_I,ARGS(m.ToEID()))))
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC("] -> ")
          Result = Core.F_print_any(Equal(ANY(Core.F_CALL(C_module_I,ARGS(m.ToEID()))),ToClass(c).Name.Defined().Id()).Id())
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC("\n")
          Result = EVOID
          }}
          {
          PRINC("all m -> ")
          /* Let:5 */{ 
            var g0229UU *ClaireBoolean  
            /* noccur = 1 */
            var g0229UU_try02306 EID 
            /* Let:6 */{ 
              var g0231UU *ClaireAny  
              /* noccur = 1 */
              var g0231UU_try02327 EID 
              /* For:7 */{ 
                var m2 *ClaireAny  
                _ = m2
                g0231UU_try02327= EID{CFALSE.Id(),0}
                var m2_support *ClaireList  
                m2_support = ToProperty(OBJ(Core.F_CALL(C_selector,ARGS(m.ToEID())))).Restrictions
                for _,m2 = range(m2_support.ValuesO())/* loop2:8 */{ 
                  var void_try9 EID 
                  _ = void_try9
                  var g0233I *ClaireBoolean  
                  var g0233I_try02349 EID 
                  /* Let:9 */{ 
                    var g0235UU *ClaireBoolean  
                    /* noccur = 1 */
                    var g0235UU_try023610 EID 
                    PRINC("---- try m2 = ")
                    g0235UU_try023610 = Core.F_print_any(m2)
                    /* ERROR PROTECTION INSERTED (g0235UU_try023610-g0235UU_try023610) */
                    if !ErrorIn(g0235UU_try023610) {
                    PRINC(" in ")
                    g0235UU_try023610 = Core.F_print_any(ToRestriction(m2).Module_I.Id())
                    /* ERROR PROTECTION INSERTED (g0235UU_try023610-g0235UU_try023610) */
                    if !ErrorIn(g0235UU_try023610) {
                    PRINC("\n")
                    g0235UU_try023610 = EVOID
                    }}
                    /* ERROR PROTECTION INSERTED (g0235UU_try023610-g0235UU_try023610) */
                    if !ErrorIn(g0235UU_try023610) {
                    if (C_method.Id() == m2.Isa.Id()) /* If:10 */{ 
                      /* Let:11 */{ 
                        var g0227 *ClaireMethod   = ToMethod(m2)
                        /* noccur = 3 */
                        if ((g0227.Module_I.Id() == ANY(Core.F_CALL(C_module_I,ARGS(m.ToEID())))) && 
                            (Equal(Core.F__exp_type(ToType(c),ToType(g0227.Domain.ValuesO()[1-1])).Id(),CEMPTY.Id()) != CTRUE)) /* If:12 */{ 
                          g0235UU_try023610 = EID{F_Generate_arg_match_list(F_Generate_go_signature_method(g0227),_Zsig).Id(),0}
                          } else {
                          g0235UU_try023610 = EID{CTRUE.Id(),0}
                          /* If-12 */} 
                        /* Let-11 */} 
                      } else {
                      g0235UU_try023610 = EID{CTRUE.Id(),0}
                      /* If-10 */} 
                    }
                    /* ERROR PROTECTION INSERTED (g0235UU-g0233I_try02349) */
                    if ErrorIn(g0235UU_try023610) {g0233I_try02349 = g0235UU_try023610
                    } else {
                    g0235UU = ToBoolean(OBJ(g0235UU_try023610))
                    g0233I_try02349 = EID{g0235UU.Not.Id(),0}
                    }
                    /* Let-9 */} 
                  /* ERROR PROTECTION INSERTED (g0233I-void_try9) */
                  if ErrorIn(g0233I_try02349) {void_try9 = g0233I_try02349
                  } else {
                  g0233I = ToBoolean(OBJ(g0233I_try02349))
                  if (g0233I == CTRUE) /* If:9 */{ 
                     /*v = g0231UU_try02327, s =EID*/
g0231UU_try02327 = EID{CTRUE.Id(),0}
                    break
                    } else {
                    void_try9 = EID{CFALSE.Id(),0}
                    /* If-9 */} 
                  }
                  /* ERROR PROTECTION INSERTED (void_try9-g0231UU_try02327) */
                  if ErrorIn(void_try9) {g0231UU_try02327 = void_try9
                  g0231UU_try02327 = void_try9
                  break
                  } else {
                  }
                  /* loop-8 */} 
                /* For-7 */} 
              /* ERROR PROTECTION INSERTED (g0231UU-g0229UU_try02306) */
              if ErrorIn(g0231UU_try02327) {g0229UU_try02306 = g0231UU_try02327
              } else {
              g0231UU = ANY(g0231UU_try02327)
              g0229UU_try02306 = EID{Core.F_not_any(g0231UU).Id(),0}
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (g0229UU-Result) */
            if ErrorIn(g0229UU_try02306) {Result = g0229UU_try02306
            } else {
            g0229UU = ToBoolean(OBJ(g0229UU_try02306))
            Result = Core.F_print_any(g0229UU.Id())
            }
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC("\n")
          Result = EVOID
          }
          }}}
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: dMethod? @ any (throw: true) 
func E_dMethod_ask_any (m EID) EID { 
    return /*(sm for dMethod? @ any= EID)*/ F_dMethod_ask_any(ANY(m) )} 
  
// same argument types for all restrictions, excluding the range (that is included in go_signature)
/* {1} OPT.The go function for: arg_match(l1:list<class>,l2:list<class>) [] */
func F_Generate_arg_match_list (l1 *ClaireList ,l2 *ClaireList ) *ClaireBoolean  { 
    // procedure body with s =  
var Result *ClaireBoolean  
    /* Let:2 */{ 
      var n int  = l1.Length()
      /* noccur = 2 */
      /* and:3 */{ 
        var v_and3 *ClaireBoolean  
        
        v_and3 = Equal(MakeInteger(l2.Length()).Id(),MakeInteger(n).Id())
        if (v_and3 == CFALSE) {Result = CFALSE
        } else /* arg:4 */{ 
          /* Let:5 */{ 
            var g0238UU *ClaireAny  
            /* noccur = 1 */
            /* Let:6 */{ 
              var i int  = 2
              /* noccur = 5 */
              /* Let:7 */{ 
                var g0237 int  = (n-1)
                /* noccur = 1 */
                g0238UU= CFALSE.Id()
                for (i <= g0237) /* while:8 */{ 
                  if (Equal(l1.ValuesO()[i-1],l2.ValuesO()[i-1]) != CTRUE) /* If:9 */{ 
                     /*v = g0238UU, s =any*/
g0238UU = CTRUE.Id()
                    break
                    /* If-9 */} 
                  i = (i+1)
                  /* while-8 */} 
                /* Let-7 */} 
              /* Let-6 */} 
            v_and3 = Core.F_not_any(g0238UU)
            /* Let-5 */} 
          if (v_and3 == CFALSE) {Result = CFALSE
          } else /* arg:5 */{ 
            Result = CTRUE/* arg-5 */} 
          /* arg-4 */} 
        /* and-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: arg_match @ list (throw: false) 
func E_Generate_arg_match_list (l1 EID,l2 EID) EID { 
    return EID{/*(sm for arg_match @ list= boolean)*/ F_Generate_arg_match_list(ToList(OBJ(l1)),ToList(OBJ(l2)) ).Id(),0}} 
  
// create the function (a name) for a method with selector p and signature l
// the name of the module where p was defined is included (until claire => public)
/* {1} OPT.The go function for: Compile/function_name(p:property,l:list) [] */
func F_Compile_function_name_property2 (p *ClaireProperty ,l *ClaireList ) *ClaireString  { 
    // procedure body with s =  
var Result *ClaireString  
    /* Let:2 */{ 
      var n int  = 0
      /* noccur = 4 */
      /* Let:3 */{ 
        var m int  = 0
        /* noccur = 2 */
        /* Let:4 */{ 
          var md *ClaireModule   = p.Name.Module_I()
          /* noccur = 2 */
          /* Let:5 */{ 
            var c *ClaireClass   = ToTypeExpression(l.At(1-1)).Class_I()
            /* noccur = 2 */
            /* Let:6 */{ 
              var r *ClaireString   = F_append_string(F_append_string(p.Name.String_I(),MakeString("_")),c.Name.String_I())
              /* noccur = 4 */
              if ((p.Id() != Core.C_main.Id()) && 
                  (md.Id() != C_claire.Id())) /* If:7 */{ 
                r = F_append_string(F_append_string(md.Name.String_I(),MakeString("_")),r)
                /* If-7 */} 
              /* For:7 */{ 
                var r *ClaireAny  
                _ = r
                for _,r = range(p.Restrictions.ValuesO())/* loop:8 */{ 
                  if (c.Id() == Core.F_domain_I_restriction(ToRestriction(r)).Id()) /* If:9 */{ 
                    n = (n+1)
                    /* If-9 */} 
                  if (Optimize.F_Optimize__equalsig_ask_list(l,ToRestriction(r).Domain) == CTRUE) /* If:9 */{ 
                    m = n
                    /* If-9 */} 
                  /* loop-8 */} 
                /* For-7 */} 
              if (n <= 1) /* If:7 */{ 
                Result = r
                } else {
                Result = F_append_string(r,F_string_I_integer(m))
                /* If-7 */} 
              /* Let-6 */} 
            /* Let-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: Compile/function_name @ list<type_expression>(property, list) (throw: false) 
func E_Compile_function_name_property2 (p EID,l EID) EID { 
    return EID{/*(sm for Compile/function_name @ list<type_expression>(property, list)= string)*/ F_Compile_function_name_property2(ToProperty(OBJ(p)),ToList(OBJ(l)) ).Id(),0}} 
  
/* {1} OPT.The go function for: at(p:go_producer) [] */
func (p *GenerateGoProducer ) At ()  { 
    // procedure body with s =  
PRINC(".")
    } 
  
// The EID go function for: at @ go_producer (throw: false) 
func E_Generate_at_go_producer (p EID) EID { 
    /*(sm for at @ go_producer= void)*/ ToGenerateGoProducer(OBJ(p)).At( )
    return EVOID} 
  
// prints a list of arguments with types / replaces typed_args_list
/* {1} OPT.The go function for: goVariables(p:go_producer,self:list) [] */
func (p *GenerateGoProducer ) GoVariables (self *ClaireList ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var prems *ClaireBoolean   = CTRUE
      /* noccur = 2 */
      /* For:3 */{ 
        var x *ClaireAny  
        _ = x
        Result= EID{CFALSE.Id(),0}
        var x_support *ClaireList  
        x_support = self
        x_len := x_support.Length()
        for i_it := 0; i_it < x_len; i_it++ { 
          x = x_support.At(i_it)
          var void_try5 EID 
          _ = void_try5
          { 
          if (prems == CTRUE) /* If:5 */{ 
            prems = CFALSE
            } else {
            PRINC(",")
            /* If-5 */} 
          void_try5 = p.GoVariable(To_Variable(x))
          /* ERROR PROTECTION INSERTED (void_try5-void_try5) */
          if ErrorIn(void_try5) {Result = void_try5
          break
          } else {
          }
          }
          /* loop-4 */} 
        /* For-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: goVariables @ go_producer (throw: true) 
func E_Generate_goVariables_go_producer (p EID,self EID) EID { 
    return /*(sm for goVariables @ go_producer= EID)*/ ToGenerateGoProducer(OBJ(p)).GoVariables(ToList(OBJ(self)) )} 
  
// prints a variable declaration (inside an arg list
/* {1} OPT.The go function for: goVariable(p:go_producer,v:Variable) [] */
func (p *GenerateGoProducer ) GoVariable (v *ClaireVariable ) EID { 
    var Result EID 
    Result = F_iClaire_ident_go_producer1(p,v)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(" ")
    F_Generate_interface_I_class(v.Range.Class_I())
    PRINC("")
    Result = EVOID
    }
    return Result} 
  
// The EID go function for: goVariable @ go_producer (throw: true) 
func E_Generate_goVariable_go_producer (p EID,v EID) EID { 
    return /*(sm for goVariable @ go_producer= EID)*/ ToGenerateGoProducer(OBJ(p)).GoVariable(To_Variable(OBJ(v)) )} 
  
// prints the name of a method as a go method 
// Here we use the list of exceptions (kernel_methods) to force a "go method syntax" (with possibly a forced name)
// this is convenient when cross-compiling (when method move from one module/package to another)
/* {1} OPT.The go function for: goMethod(m:method) [] */
func F_Generate_goMethod_method (m *ClaireMethod )  { 
    // procedure body with s =  
/* Let:2 */{ 
      var lm *ClaireList   = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).KernelMethods
      /* noccur = 4 */
      /* Let:3 */{ 
        var i int  = F_index_list(lm,m.Id())
        /* noccur = 4 */
        if ((i > 0) && 
            ((lm.Length() > i) && 
              (C_string.Id() == lm.At((i+1)-1).Isa.Id()))) /* If:4 */{ 
          Core.F_CALL(C_c_princ,ARGS(lm.At((i+1)-1).ToEID()))
          } else {
          F_c_princ_string(F_Generate_capitalize_symbol(m.Selector.Name))
          /* If-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    } 
  
// The EID go function for: goMethod @ method (throw: false) 
func E_Generate_goMethod_method (m EID) EID { 
    /*(sm for goMethod @ method= void)*/ F_Generate_goMethod_method(ToMethod(OBJ(m)) )
    return EVOID} 
  
// prints the name of a function as a go function F_f
// NOTE : the link method <=> go function is not stored (the function is not known by CLAIRE)
// imported functions do not refer to the module/package
/* {1} OPT.The go function for: goFunction(m:method) [] */
func F_Generate_goFunction_method (m *ClaireMethod ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var md *ClaireModule   = m.Module_I
      /* noccur = 4 */
      if ((md.Id() != C_Kernel.Id()) && 
          ((md.Id() != ToGenerateGoProducer(Optimize.C_PRODUCER.Value).Current.Id()) && 
            ((md.Id() != C_claire.Id()) && 
              ((m.Functional.Id() == CNULL) || 
                  (F_imported_function_ask_any(m.Functional.Id()) != CTRUE))))) /* If:3 */{ 
        F_Generate_cap_short_symbol(md.Name)
        PRINC(".")
        /* If-3 */} 
      PRINC("F_")
      /* Let:3 */{ 
        var g0239UU *ClaireString  
        /* noccur = 1 */
        var g0239UU_try02404 EID 
        g0239UU_try02404 = F_Generate_getFunctionName_method(m)
        /* ERROR PROTECTION INSERTED (g0239UU-Result) */
        if ErrorIn(g0239UU_try02404) {Result = g0239UU_try02404
        } else {
        g0239UU = ToString(OBJ(g0239UU_try02404))
        F_Generate_import_princ_string(g0239UU)
        Result = EVOID
        }
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: goFunction @ method (throw: true) 
func E_Generate_goFunction_method (m EID) EID { 
    return /*(sm for goFunction @ method= EID)*/ F_Generate_goFunction_method(ToMethod(OBJ(m)) )} 
  
// specialized version for Core method
/* {1} OPT.The go function for: preCore?(_CL_obj:void) [] */
func F_Generate_preCore_ask_void ()  { 
    // procedure body with s =  
if (ToGenerateGoProducer(Optimize.C_PRODUCER.Value).Current.Id() != Core.It.Id()) /* If:2 */{ 
      PRINC("Core.")
      /* If-2 */} 
    } 
  
// The EID go function for: preCore? @ void (throw: false) 
func E_Generate_preCore_ask_void (_CL_obj EID) EID { 
    /*(sm for preCore? @ void= void)*/ F_Generate_preCore_ask_void( )
    return EVOID} 
  
// prints the name of the EID method that is linked by the compiler to the method
/* {1} OPT.The go function for: goEIDFunction(m:method) [] */
func F_Generate_goEIDFunction_method (m *ClaireMethod ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var s *ClaireString  
      /* noccur = 2 */
      var s_try02413 EID 
      s_try02413 = F_Generate_getFunctionName_method(m)
      /* ERROR PROTECTION INSERTED (s-Result) */
      if ErrorIn(s_try02413) {Result = s_try02413
      } else {
      s = ToString(OBJ(s_try02413))
      PRINC("MakeFunction")
      F_princ_integer(m.Domain.Length())
      PRINC("(E_")
      F_Generate_import_princ_string(s)
      PRINC(",")
      Result = Core.F_print_any((s).Id())
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(")")
      Result = EVOID
      }
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: goEIDFunction @ method (throw: true) 
func E_Generate_goEIDFunction_method (m EID) EID { 
    return /*(sm for goEIDFunction @ method= EID)*/ F_Generate_goEIDFunction_method(ToMethod(OBJ(m)) )} 
  
// prints the function MakeFunction(...) expression
/* {1} OPT.The go function for: goEIDFunctionName(m:method) [] */
func F_Generate_goEIDFunctionName_method (m *ClaireMethod ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var f *ClaireString  
      /* noccur = 1 */
      var f_try02423 EID 
      f_try02423 = F_Generate_getFunctionName_method(m)
      /* ERROR PROTECTION INSERTED (f-Result) */
      if ErrorIn(f_try02423) {Result = f_try02423
      } else {
      f = ToString(OBJ(f_try02423))
      PRINC("E_")
      F_c_princ_string(f)
      PRINC("")
      Result = EVOID
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: goEIDFunctionName @ method (throw: true) 
func E_Generate_goEIDFunctionName_method (m EID) EID { 
    return /*(sm for goEIDFunctionName @ method= EID)*/ F_Generate_goEIDFunctionName_method(ToMethod(OBJ(m)) )} 
  
// special function for self_eval of type  => added as an extra paramer of type eFunc
// cf. goexp: AddMethod -> AddEvalMethod     
/* {1} OPT.The go function for: goEvalFunction(m:method) [] */
func F_Generate_goEvalFunction_method (m *ClaireMethod )  { 
    // procedure body with s =  
/* Let:2 */{ 
      var c *ClaireClass   = Core.F_domain_I_restriction(ToRestriction(m.Id()))
      /* noccur = 1 */
      PRINC(",EVAL_")
      c.Name.CPrinc()
      PRINC("")
      /* Let-2 */} 
    } 
  
// The EID go function for: goEvalFunction @ method (throw: false) 
func E_Generate_goEvalFunction_method (m EID) EID { 
    /*(sm for goEvalFunction @ method= void)*/ F_Generate_goEvalFunction_method(ToMethod(OBJ(m)) )
    return EVOID} 
  
// exceptions
// get function name
/* {1} OPT.The go function for: getFunctionName(m:method) [] */
func F_Generate_getFunctionName_method (m *ClaireMethod ) EID { 
    var Result EID 
    if (m.Id() == C_Generate__starlength_string_star.Value) /* If:2 */{ 
      Result = EID{MakeString("length_string").Id(),0}
      /* If!2 */}  else if (m.Id() == C_Generate__starnth_list_star.Value) /* If:2 */{ 
      Result = EID{MakeString("nth_list").Id(),0}
      /* If!2 */}  else if (m.Id() == C_Generate__starset_I_list_star.Value) /* If:2 */{ 
      Result = EID{MakeString("set_I_list").Id(),0}
      /* If!2 */}  else if (m.Id() == C_Generate__starstack_apply_star.Value) /* If:2 */{ 
      Result = EID{MakeString("CALL").Id(),0}
      /* If!2 */}  else if (m.Id() == C_Generate__starsuper_apply_star.Value) /* If:2 */{ 
      Result = EID{MakeString("SUPER").Id(),0}
      /* If!2 */}  else if (m.Id() == C_Generate__starbelong_star.Value) /* If:2 */{ 
      Result = EID{MakeString("BELONG").Id(),0}
      /* If!2 */}  else if (m.Functional.Id() != CNULL) /* If:2 */{ 
      Result = EID{F_string_I_function(m.Functional).Id(),0}
      } else {
      Result = Core.F_CALL(Optimize.C_Compile_function_name,ARGS(EID{m.Selector.Id(),0},EID{m.Domain.Id(),0}))
      /* If-2 */} 
    return Result} 
  
// The EID go function for: getFunctionName @ method (throw: true) 
func E_Generate_getFunctionName_method (m EID) EID { 
    return /*(sm for getFunctionName @ method= EID)*/ F_Generate_getFunctionName_method(ToMethod(OBJ(m)) )} 
  
// ugly : reverse engineer a compiled definition into a method
// we need to do something better
/* {1} OPT.The go function for: retreive_method(p:any,lf:any) [] */
func F_Generate_retreive_method_any (p *ClaireAny ,lf *ClaireAny ) EID { 
    var Result EID 
    if (p.Isa.IsIn(C_property) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0243 *ClaireProperty   = ToProperty(p)
        /* noccur = 2 */
        /* Let:4 */{ 
          var m *ClaireAny  
          /* noccur = 2 */
          var m_try02535 EID 
          /* Let:5 */{ 
            var g0254UU *ClaireAny  
            /* noccur = 1 */
            var g0254UU_try02556 EID 
            g0254UU_try02556 = F_Generate_retreive_list_any(lf)
            /* ERROR PROTECTION INSERTED (g0254UU-m_try02535) */
            if ErrorIn(g0254UU_try02556) {m_try02535 = g0254UU_try02556
            } else {
            g0254UU = ANY(g0254UU_try02556)
            m_try02535 = Core.F_CALL(ToProperty(Core.C__at.Id()),ARGS(EID{g0243.Id(),0},g0254UU.ToEID()))
            }
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (m-Result) */
          if ErrorIn(m_try02535) {Result = m_try02535
          } else {
          m = ANY(m_try02535)
          if (C_method.Id() == m.Isa.Id()) /* If:5 */{ 
            /* Let:6 */{ 
              var g0244 *ClaireMethod   = ToMethod(m)
              /* noccur = 1 */
              Result = EID{g0244.Id(),0}
              /* Let-6 */} 
            } else {
            Result = ToException(Core.C_general_error.Make(MakeString("there is no method ~S @ ~S").Id(),MakeConstantList(g0243.Id(),lf).Id())).Close()
            /* If-5 */} 
          }
          /* Let-4 */} 
        /* Let-3 */} 
      } else {
      Result = ToException(Core.C_general_error.Make(MakeString("we have a problem to retreive ~S (not a property) at ~S").Id(),MakeConstantList(p,lf).Id())).Close()
      /* If-2 */} 
    return Result} 
  
// The EID go function for: retreive_method @ any (throw: true) 
func E_Generate_retreive_method_any (p EID,lf EID) EID { 
    return /*(sm for retreive_method @ any= EID)*/ F_Generate_retreive_method_any(ANY(p),ANY(lf) )} 
  
// constrained eval in disguise : returns a type or a list of types from CLAIRE expressions
/* {1} OPT.The go function for: retreive_list(x:any) [] */
func F_Generate_retreive_list_any (x *ClaireAny ) EID { 
    var Result EID 
    if (x.Isa.IsIn(C_type) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0256 *ClaireType   = ToType(x)
        /* noccur = 1 */
        Result = EID{g0256.Id(),0}
        /* Let-3 */} 
      /* If!2 */}  else if (x.Isa.IsIn(C_property) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0257 *ClaireProperty   = ToProperty(x)
        /* noccur = 1 */
        Result = EID{g0257.Id(),0}
        /* Let-3 */} 
      /* If!2 */}  else if (x.Isa.IsIn(Language.C_List) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0258 *Language.List   = Language.To_List(x)
        /* noccur = 1 */
        /* Iteration:4 */{ 
          var v_list4 *ClaireList  
          var y *ClaireAny  
          var v_local4 *ClaireAny  
          v_list4 = g0258.Args
          Result = EID{CreateList(ToType(CEMPTY.Id()),v_list4.Length()).Id(),0}
          for CLcount := 0; CLcount < v_list4.Length(); CLcount++{ 
            y = v_list4.At(CLcount)
            var v_local4_try02626 EID 
            v_local4_try02626 = F_Generate_retreive_list_any(y)
            /* ERROR PROTECTION INSERTED (v_local4-Result) */
            if ErrorIn(v_local4_try02626) {Result = v_local4_try02626
            Result = v_local4_try02626
            break
            } else {
            v_local4 = ANY(v_local4_try02626)
            ToList(OBJ(Result)).PutAt(CLcount,v_local4)
            } 
          }
          /* Iteration-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (x.Isa.IsIn(Language.C_Tuple) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0259 *Language.Tuple   = Language.To_Tuple(x)
        /* noccur = 1 */
        /* Let:4 */{ 
          var g0263UU *ClaireList  
          /* noccur = 1 */
          var g0263UU_try02645 EID 
          /* Iteration:5 */{ 
            var v_list5 *ClaireList  
            var y *ClaireAny  
            var v_local5 *ClaireAny  
            v_list5 = g0259.Args
            g0263UU_try02645 = EID{CreateList(ToType(CEMPTY.Id()),v_list5.Length()).Id(),0}
            for CLcount := 0; CLcount < v_list5.Length(); CLcount++{ 
              y = v_list5.At(CLcount)
              var v_local5_try02657 EID 
              v_local5_try02657 = F_Generate_retreive_list_any(y)
              /* ERROR PROTECTION INSERTED (v_local5-g0263UU_try02645) */
              if ErrorIn(v_local5_try02657) {g0263UU_try02645 = v_local5_try02657
              g0263UU_try02645 = v_local5_try02657
              break
              } else {
              v_local5 = ANY(v_local5_try02657)
              ToList(OBJ(g0263UU_try02645)).PutAt(CLcount,v_local5)
              } 
            }
            /* Iteration-5 */} 
          /* ERROR PROTECTION INSERTED (g0263UU-Result) */
          if ErrorIn(g0263UU_try02645) {Result = g0263UU_try02645
          } else {
          g0263UU = ToList(OBJ(g0263UU_try02645))
          Result = EID{g0263UU.Tuple_I().Id(),0}
          }
          /* Let-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (x.Isa.IsIn(Language.C_Call_method) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0260 *Language.CallMethod   = Language.To_CallMethod(x)
        /* noccur = 18 */
        if ((g0260.Arg.Selector.Id() == C_nth.Id()) && 
            (g0260.Args.Length() == 2)) /* If:4 */{ 
          /* Let:5 */{ 
            var g0266UU *ClaireAny  
            /* noccur = 1 */
            var g0266UU_try02686 EID 
            g0266UU_try02686 = F_Generate_retreive_list_any(g0260.Args.At(1-1))
            /* ERROR PROTECTION INSERTED (g0266UU-Result) */
            if ErrorIn(g0266UU_try02686) {Result = g0266UU_try02686
            } else {
            g0266UU = ANY(g0266UU_try02686)
            /* Let:6 */{ 
              var g0267UU *ClaireAny  
              /* noccur = 1 */
              var g0267UU_try02697 EID 
              g0267UU_try02697 = F_Generate_retreive_list_any(g0260.Args.At(2-1))
              /* ERROR PROTECTION INSERTED (g0267UU-Result) */
              if ErrorIn(g0267UU_try02697) {Result = g0267UU_try02697
              } else {
              g0267UU = ANY(g0267UU_try02697)
              Result = Core.F_CALL(C_nth,ARGS(g0266UU.ToEID(),g0267UU.ToEID()))
              }
              /* Let-6 */} 
            }
            /* Let-5 */} 
          /* If!4 */}  else if ((g0260.Arg.Selector.Id() == C_nth.Id()) && 
            (g0260.Args.Length() == 3)) /* If:4 */{ 
          /* Let:5 */{ 
            var g0270UU *ClaireAny  
            /* noccur = 1 */
            var g0270UU_try02736 EID 
            g0270UU_try02736 = F_Generate_retreive_list_any(g0260.Args.At(1-1))
            /* ERROR PROTECTION INSERTED (g0270UU-Result) */
            if ErrorIn(g0270UU_try02736) {Result = g0270UU_try02736
            } else {
            g0270UU = ANY(g0270UU_try02736)
            /* Let:6 */{ 
              var g0271UU *ClaireAny  
              /* noccur = 1 */
              var g0271UU_try02747 EID 
              g0271UU_try02747 = F_Generate_retreive_list_any(g0260.Args.At(2-1))
              /* ERROR PROTECTION INSERTED (g0271UU-Result) */
              if ErrorIn(g0271UU_try02747) {Result = g0271UU_try02747
              } else {
              g0271UU = ANY(g0271UU_try02747)
              /* Let:7 */{ 
                var g0272UU *ClaireAny  
                /* noccur = 1 */
                var g0272UU_try02758 EID 
                g0272UU_try02758 = F_Generate_retreive_list_any(g0260.Args.At(3-1))
                /* ERROR PROTECTION INSERTED (g0272UU-Result) */
                if ErrorIn(g0272UU_try02758) {Result = g0272UU_try02758
                } else {
                g0272UU = ANY(g0272UU_try02758)
                Result = Core.F_CALL(C_nth,ARGS(g0270UU.ToEID(),g0271UU.ToEID(),g0272UU.ToEID()))
                }
                /* Let-7 */} 
              }
              /* Let-6 */} 
            }
            /* Let-5 */} 
          /* If!4 */}  else if ((g0260.Arg.Selector.Id() == Core.C_Core_param_I.Id()) && 
            (g0260.Args.Length() == 2)) /* If:4 */{ 
          /* Let:5 */{ 
            var g0276UU *ClaireAny  
            /* noccur = 1 */
            var g0276UU_try02786 EID 
            g0276UU_try02786 = F_Generate_retreive_list_any(g0260.Args.At(1-1))
            /* ERROR PROTECTION INSERTED (g0276UU-Result) */
            if ErrorIn(g0276UU_try02786) {Result = g0276UU_try02786
            } else {
            g0276UU = ANY(g0276UU_try02786)
            /* Let:6 */{ 
              var g0277UU *ClaireAny  
              /* noccur = 1 */
              var g0277UU_try02797 EID 
              g0277UU_try02797 = F_Generate_retreive_list_any(g0260.Args.At(2-1))
              /* ERROR PROTECTION INSERTED (g0277UU-Result) */
              if ErrorIn(g0277UU_try02797) {Result = g0277UU_try02797
              } else {
              g0277UU = ANY(g0277UU_try02797)
              Result = EID{Core.F_param_I_class(ToClass(g0276UU),ToType(g0277UU)).Id(),0}
              }
              /* Let-6 */} 
            }
            /* Let-5 */} 
          /* If!4 */}  else if ((g0260.Arg.Selector.Id() == Core.C_U.Id()) && 
            (g0260.Args.Length() == 2)) /* If:4 */{ 
          /* Let:5 */{ 
            var g0280UU *ClaireAny  
            /* noccur = 1 */
            var g0280UU_try02826 EID 
            g0280UU_try02826 = F_Generate_retreive_list_any(g0260.Args.At(1-1))
            /* ERROR PROTECTION INSERTED (g0280UU-Result) */
            if ErrorIn(g0280UU_try02826) {Result = g0280UU_try02826
            } else {
            g0280UU = ANY(g0280UU_try02826)
            /* Let:6 */{ 
              var g0281UU *ClaireAny  
              /* noccur = 1 */
              var g0281UU_try02837 EID 
              g0281UU_try02837 = F_Generate_retreive_list_any(g0260.Args.At(2-1))
              /* ERROR PROTECTION INSERTED (g0281UU-Result) */
              if ErrorIn(g0281UU_try02837) {Result = g0281UU_try02837
              } else {
              g0281UU = ANY(g0281UU_try02837)
              Result = EID{Core.F_U_type(ToType(g0280UU),ToType(g0281UU)).Id(),0}
              }
              /* Let-6 */} 
            }
            /* Let-5 */} 
          } else {
          Result = ToException(Core.C_general_error.Make(MakeString("we need to extend retreive_list to handle a type call: ~S").Id(),MakeConstantList(g0260.Id()).Id())).Close()
          /* If-4 */} 
        /* Let-3 */} 
      } else {
      Result = ToException(Core.C_general_error.Make(MakeString("we need to extend retreive_list to handle ~S").Id(),MakeConstantList(x).Id())).Close()
      /* If-2 */} 
    return Result} 
  
// The EID go function for: retreive_list @ any (throw: true) 
func E_Generate_retreive_list_any (x EID) EID { 
    return /*(sm for retreive_list @ any= EID)*/ F_Generate_retreive_list_any(ANY(x) )} 
  
// *******************************************************************
// *       Part 3: interface declarations                            *
// *******************************************************************
// How to declare a sort in Go. The boolean tells if we are in an external
// mode , in which case we produce the C sort. Otherwise, we use OIDs.
// THERE are 5 sorts in go : int, float, char,  any (object) and EID
// there are 7 sorts in CLAIRE : int, float, char, object, string, function, any
/* {1} OPT.The go function for: interface!(self:class) [] */
func F_Generate_interface_I_class (self *ClaireClass )  { 
    // procedure body with s =  
if (self.Id() == C_void.Id()) /* If:2 */{ 
      PRINC("void ")
      /* If!2 */}  else if (self.Id() == C_integer.Id()) /* If:2 */{ 
      PRINC("int")
      /* If!2 */}  else if (self.Id() == C_float.Id()) /* If:2 */{ 
      PRINC("float64")
      /* If!2 */}  else if (self.Id() == C_char.Id()) /* If:2 */{ 
      PRINC("rune")
      /* If!2 */}  else if (self.Id() == Optimize.C_EID.Id()) /* If:2 */{ 
      PRINC("EID")
      } else {
      PRINC("*")
      F_Generate_go_class_class(self)
      PRINC(" ")
      /* If-2 */} 
    } 
  
// The EID go function for: interface! @ class (throw: false) 
func E_Generate_interface_I_class (self EID) EID { 
    /*(sm for interface! @ class= void)*/ F_Generate_interface_I_class(ToClass(OBJ(self)) )
    return EVOID} 
  
// general translation method: x is an expression that must be translated
// to a CLAIRE object (*ClaireX). x is known to be functional ! s is the sort for x.
/* {1} OPT.The go function for: to_cl(c:go_producer,x:any,s:class) [] */
func (c *GenerateGoProducer ) ToCl (x *ClaireAny ,s *ClaireClass ) EID { 
    var Result EID 
    if (s.Id() == C_void.Id()) /* If:2 */{ 
      PRINC("Void(")
      Result = Core.F_CALL(C_Generate_g_expression,ARGS(x.ToEID(),EID{C_any.Id(),0}))
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(")")
      Result = EVOID
      }
      /* If!2 */}  else if (s.Id() == C_integer.Id()) /* If:2 */{ 
      PRINC("MakeInteger(")
      Result = Core.F_CALL(C_Generate_g_expression,ARGS(x.ToEID(),EID{C_integer.Id(),0}))
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(")")
      Result = EVOID
      }
      /* If!2 */}  else if (s.Id() == C_float.Id()) /* If:2 */{ 
      PRINC("MakeFloat(")
      Result = Core.F_CALL(C_Generate_g_expression,ARGS(x.ToEID(),EID{C_float.Id(),0}))
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(")")
      Result = EVOID
      }
      /* If!2 */}  else if (s.Id() == C_char.Id()) /* If:2 */{ 
      PRINC("MakeChar(")
      Result = Core.F_CALL(C_Generate_g_expression,ARGS(x.ToEID(),EID{C_char.Id(),0}))
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(")")
      Result = EVOID
      }
      /* If!2 */}  else if ((s.IsIn(C_object) == CTRUE) || 
        ((s.Id() == C_any.Id()) || 
          (s.Id() == C_primitive.Id()))) /* If:2 */{ 
      Result = Core.F_CALL(C_Generate_g_expression,ARGS(x.ToEID(),EID{s.Id(),0}))
      } else {
      Result = ToException(Core.C_general_error.Make(MakeString("[internal] to_cl for a ~S is not implemented").Id(),MakeConstantList(s.Id()).Id())).Close()
      /* If-2 */} 
    return Result} 
  
// The EID go function for: to_cl @ go_producer (throw: true) 
func E_Generate_to_cl_go_producer (c EID,x EID,s EID) EID { 
    return /*(sm for to_cl @ go_producer= EID)*/ ToGenerateGoProducer(OBJ(c)).ToCl(ANY(x),ToClass(OBJ(s)) )} 
  
// reverse function : produce a native forme from a claire object
// quite simple with go since for object, OID is the object
/* {1} OPT.The go function for: to_c(c:go_producer,x:any,s:class) [] */
func (c *GenerateGoProducer ) ToC (x *ClaireAny ,s *ClaireClass )  { 
    // procedure body with s =  
if (x == CNULL) /* If:2 */{ 
      PRINC("CNULL")
      /* If!2 */}  else if ((s.Id() == C_integer.Id()) || 
        ((s.Id() == C_float.Id()) || 
          ((s.Id() == C_string.Id()) || 
            ((s.Id() == C_char.Id()) || 
              (s.Id() == C_function.Id()))))) /* If:2 */{ 
      Core.F_CALL(C_Generate_expression,ARGS(x.ToEID(),EID{s.Id(),0}))
      PRINC(".Value")
      } else {
      Core.F_CALL(C_Generate_c_expression,ARGS(x.ToEID(),EID{s.Id(),0}))
      /* If-2 */} 
    } 
  
// The EID go function for: to_c @ go_producer (throw: false) 
func E_Generate_to_c_go_producer (c EID,x EID,s EID) EID { 
    /*(sm for to_c @ go_producer= void)*/ ToGenerateGoProducer(OBJ(c)).ToC(ANY(x),ToClass(OBJ(s)) )
    return EVOID} 
  
// new for go: compile to an EID form (128 bit generic representation)
// s is the expected sort
/* {1} OPT.The go function for: to_eid(c:go_producer,x:any,s:class) [] */
func (c *GenerateGoProducer ) ToEid (x *ClaireAny ,s *ClaireClass ) EID { 
    var Result EID 
    if (s.Id() == C_void.Id()) /* If:2 */{ 
      PRINC("EVOID")
      Result = EVOID
      /* If!2 */}  else if (s.Id() == C_integer.Id()) /* If:2 */{ 
      PRINC("EID{C__INT,IVAL(")
      Result = Core.F_CALL(C_Generate_g_expression,ARGS(x.ToEID(),EID{C_integer.Id(),0}))
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(")}")
      Result = EVOID
      }
      /* If!2 */}  else if (s.Id() == C_float.Id()) /* If:2 */{ 
      PRINC("EID{C__FLOAT,FVAL(")
      Result = Core.F_CALL(C_Generate_g_expression,ARGS(x.ToEID(),EID{C_float.Id(),0}))
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(")}")
      Result = EVOID
      }
      /* If!2 */}  else if (s.Id() == C_char.Id()) /* If:2 */{ 
      PRINC("EID{C__CHAR,CVAL(")
      Result = Core.F_CALL(C_Generate_g_expression,ARGS(x.ToEID(),EID{C_char.Id(),0}))
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(")}")
      Result = EVOID
      }
      /* If!2 */}  else if ((s.Id() == C_string.Id()) || 
        ((s.Id() == C_function.Id()) || 
          (s.IsIn(C_object) == CTRUE))) /* If:2 */{ 
      PRINC("EID{")
      Result = Core.F_CALL(C_Generate_g_expression,ARGS(x.ToEID(),EID{C_any.Id(),0}))
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(",0}")
      Result = EVOID
      }
      /* If!2 */}  else if ((s.Id() == C_any.Id()) || 
        (s.Id() == C_primitive.Id())) /* If:2 */{ 
      Result = c.ToCl(x,s)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(".ToEID()")
      Result = EVOID
      }
      } else {
      Result = ToException(Core.C_general_error.Make(MakeString("[internal] to_eid for a ~S is not implemented").Id(),MakeConstantList(s.Id()).Id())).Close()
      /* If-2 */} 
    return Result} 
  
// The EID go function for: to_eid @ go_producer (throw: true) 
func E_Generate_to_eid_go_producer (c EID,x EID,s EID) EID { 
    return /*(sm for to_eid @ go_producer= EID)*/ ToGenerateGoProducer(OBJ(c)).ToEid(ANY(x),ToClass(OBJ(s)) )} 
  
// reciprocate with an expected class e / used for variables
/* {1} OPT.The go function for: from_eid(c:go_producer,x:string,e:class) [] */
func (c *GenerateGoProducer ) FromEid (x *ClaireString ,e *ClaireClass ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var s *ClaireClass   = e.Class_I()
      /* noccur = 2 */
      Result = F_Generate_eid_prefix_class(s)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      F_c_princ_string(x)
      F_Generate_eid_post_class(s)
      Result = EVOID
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: from_eid @ go_producer (throw: true) 
func E_Generate_from_eid_go_producer (c EID,x EID,e EID) EID { 
    return /*(sm for from_eid @ go_producer= EID)*/ ToGenerateGoProducer(OBJ(c)).FromEid(ToString(OBJ(x)),ToClass(OBJ(e)) )} 
  
// reciprocate : move from EID to a sort s (if s = any, do nothing )
/* {1} OPT.The go function for: eid_prefix(s:class) [] */
func F_Generate_eid_prefix_class (s *ClaireClass ) EID { 
    var Result EID 
    if ((s.Id() == Optimize.C_EID.Id()) || 
        (s.Id() == C_void.Id())) /* If:2 */{ 
      Result = EID{CNIL.Id(),0}
      /* If!2 */}  else if (s.Id() == C_integer.Id()) /* If:2 */{ 
      PRINC("INT(")
      Result = EVOID
      /* If!2 */}  else if (s.Id() == C_float.Id()) /* If:2 */{ 
      PRINC("FLOAT(")
      Result = EVOID
      /* If!2 */}  else if (s.Id() == C_char.Id()) /* If:2 */{ 
      PRINC("CHAR(")
      Result = EVOID
      /* If!2 */}  else if ((s.Id() == C_any.Id()) || 
        (s.Id() == C_primitive.Id())) /* If:2 */{ 
      PRINC("ANY(")
      Result = EVOID
      /* If!2 */}  else if ((ToType(s.Id()).Included(ToType(C_object.Id())) == CTRUE) || 
        ((s.Id() == C_array.Id()) || 
          ((s.Id() == C_string.Id()) || 
            ((s.Id() == C_port.Id()) || 
              (s.Id() == C_function.Id()))))) /* If:2 */{ 
      F_Generate_cast_class_class(s)
      PRINC("(OBJ(")
      Result = EVOID
      /* If!2 */}  else if (s.Id() != C_any.Id()) /* If:2 */{ 
      Result = ToException(Core.C_general_error.Make(MakeString("what the fuck: eid prefix for ~S").Id(),MakeConstantList(s.Id()).Id())).Close()
      } else {
      Result = EID{CFALSE.Id(),0}
      /* If-2 */} 
    return Result} 
  
// The EID go function for: eid_prefix @ class (throw: true) 
func E_Generate_eid_prefix_class (s EID) EID { 
    return /*(sm for eid_prefix @ class= EID)*/ F_Generate_eid_prefix_class(ToClass(OBJ(s)) )} 
  
/* {1} OPT.The go function for: eid_post(s:class) [] */
func F_Generate_eid_post_class (s *ClaireClass )  { 
    // procedure body with s =  
if ((s.Id() == Optimize.C_EID.Id()) || 
        (s.Id() == C_void.Id())) /* If:2 */{ 
      
      /* If!2 */}  else if ((s.Id() == C_char.Id()) || 
        ((s.Id() == C_any.Id()) || 
          (s.Id() == C_primitive.Id()))) /* If:2 */{ 
      PRINC(")")
      /* If!2 */}  else if ((ToType(s.Id()).Included(ToType(C_object.Id())) == CTRUE) || 
        ((s.Id() == C_array.Id()) || 
          ((s.Id() == C_string.Id()) || 
            ((s.Id() == C_port.Id()) || 
              (s.Id() == C_function.Id()))))) /* If:2 */{ 
      PRINC("))")
      /* If!2 */}  else if (s.Id() != C_any.Id()) /* If:2 */{ 
      PRINC(")")
      /* If-2 */} 
    } 
  
// The EID go function for: eid_post @ class (throw: false) 
func E_Generate_eid_post_class (s EID) EID { 
    /*(sm for eid_post @ class= void)*/ F_Generate_eid_post_class(ToClass(OBJ(s)) )
    return EVOID} 
  
// move from an integer to a EID or Object
/* {1} OPT.The go function for: integer_prefix(s:class) [] */
func F_Generate_integer_prefix_class (s *ClaireClass )  { 
    // procedure body with s =  
if (s.Id() == Optimize.C_EID.Id()) /* If:2 */{ 
      PRINC("EID{C__INT,IVAL(")
      /* If!2 */}  else if (s.Id() == C_any.Id()) /* If:2 */{ 
      PRINC("MakeInteger(")
      /* If-2 */} 
    } 
  
// The EID go function for: integer_prefix @ class (throw: false) 
func E_Generate_integer_prefix_class (s EID) EID { 
    /*(sm for integer_prefix @ class= void)*/ F_Generate_integer_prefix_class(ToClass(OBJ(s)) )
    return EVOID} 
  
// move from an integer to a EID or Object
/* {1} OPT.The go function for: float_prefix(s:class) [] */
func F_Generate_float_prefix_class (s *ClaireClass )  { 
    // procedure body with s =  
if (s.Id() == Optimize.C_EID.Id()) /* If:2 */{ 
      PRINC("EID{C__FLOAT,FVAL(")
      /* If!2 */}  else if (s.Id() == C_any.Id()) /* If:2 */{ 
      PRINC("MakeFloat(")
      /* If-2 */} 
    } 
  
// The EID go function for: float_prefix @ class (throw: false) 
func E_Generate_float_prefix_class (s EID) EID { 
    /*(sm for float_prefix @ class= void)*/ F_Generate_float_prefix_class(ToClass(OBJ(s)) )
    return EVOID} 
  
// move from an integer to a EID or Object
/* {1} OPT.The go function for: char_prefix(s:class) [] */
func F_Generate_char_prefix_class (s *ClaireClass )  { 
    // procedure body with s =  
if (s.Id() == Optimize.C_EID.Id()) /* If:2 */{ 
      PRINC("EID{C__CHAR,CVAL(")
      /* If!2 */}  else if (s.Id() == C_any.Id()) /* If:2 */{ 
      PRINC("MakeChar(")
      /* If-2 */} 
    } 
  
// The EID go function for: char_prefix @ class (throw: false) 
func E_Generate_char_prefix_class (s EID) EID { 
    /*(sm for char_prefix @ class= void)*/ F_Generate_char_prefix_class(ToClass(OBJ(s)) )
    return EVOID} 
  
// move from an integer to a EID or Object
/* {1} OPT.The go function for: string_prefix(s:class) [] */
func F_Generate_string_prefix_class (s *ClaireClass )  { 
    // procedure body with s =  
if (s.Id() == Optimize.C_EID.Id()) /* If:2 */{ 
      PRINC("EID{")
      /* If!2 */}  else if (s.Id() == C_any.Id()) /* If:2 */{ 
      PRINC("(")
      /* If-2 */} 
    } 
  
// The EID go function for: string_prefix @ class (throw: false) 
func E_Generate_string_prefix_class (s EID) EID { 
    /*(sm for string_prefix @ class= void)*/ F_Generate_string_prefix_class(ToClass(OBJ(s)) )
    return EVOID} 
  
/* {1} OPT.The go function for: string_post(s:class) [] */
func F_Generate_string_post_class (s *ClaireClass )  { 
    // procedure body with s =  
if (s.Id() == Optimize.C_EID.Id()) /* If:2 */{ 
      PRINC(".Id(),0}")
      /* If!2 */}  else if (s.Id() == C_any.Id()) /* If:2 */{ 
      PRINC(").Id()")
      /* If-2 */} 
    } 
  
// The EID go function for: string_post @ class (throw: false) 
func E_Generate_string_post_class (s EID) EID { 
    /*(sm for string_post @ class= void)*/ F_Generate_string_post_class(ToClass(OBJ(s)) )
    return EVOID} 
  
// works for integer, float, char
/* {1} OPT.The go function for: native_post(s:class) [] */
func F_Generate_native_post_class (s *ClaireClass )  { 
    // procedure body with s =  
if (s.Id() == Optimize.C_EID.Id()) /* If:2 */{ 
      PRINC(")}")
      /* If!2 */}  else if (s.Id() == C_any.Id()) /* If:2 */{ 
      PRINC(").Id()")
      /* If-2 */} 
    } 
  
// The EID go function for: native_post @ class (throw: false) 
func E_Generate_native_post_class (s EID) EID { 
    /*(sm for native_post @ class= void)*/ F_Generate_native_post_class(ToClass(OBJ(s)) )
    return EVOID} 
  
// move from ClaireId (inferred) to s (expected)
/* {1} OPT.The go function for: object_prefix(inferred:class,expected:class) [] */
func F_Generate_object_prefix_class (inferred *ClaireClass ,expected *ClaireClass )  { 
    // procedure body with s =  
if (expected.Id() == Optimize.C_EID.Id()) /* If:2 */{ 
      if (ToType(inferred.Id()).Included(ToType(C_object.Id())) == CTRUE) /* If:3 */{ 
        PRINC("EID{")
        /* If-3 */} 
      /* If!2 */}  else if (expected.Id() == inferred.Id()) /* If:2 */{ 
      
      /* If!2 */}  else if (expected.Id() == C_char.Id()) /* If:2 */{ 
      PRINC("ToChar(")
      /* If!2 */}  else if (ToType(expected.Id()).Included(ToType(C_primitive.Id())) == CTRUE) /* If:2 */{ 
      F_Generate_cast_class_class(expected)
      PRINC("(")
      /* If!2 */}  else if (ToType(expected.Id()).Included(ToType(C_object.Id())) == CTRUE) /* If:2 */{ 
      F_Generate_cast_class_class(expected)
      PRINC("(")
      /* If-2 */} 
    } 
  
// The EID go function for: object_prefix @ class (throw: false) 
func E_Generate_object_prefix_class (inferred EID,expected EID) EID { 
    /*(sm for object_prefix @ class= void)*/ F_Generate_object_prefix_class(ToClass(OBJ(inferred)),ToClass(OBJ(expected)) )
    return EVOID} 
  
// s: expected
/* {1} OPT.The go function for: object_post(inferred:class,s:class) [] */
func F_Generate_object_post_class (inferred *ClaireClass ,s *ClaireClass )  { 
    // procedure body with s =  
if (s.Id() == Optimize.C_EID.Id()) /* If:2 */{ 
      if (ToType(inferred.Id()).Included(ToType(C_object.Id())) != CTRUE) /* If:3 */{ 
        PRINC(".ToEID()")
        } else {
        PRINC(".Id(),0}")
        /* If-3 */} 
      /* If!2 */}  else if (s.Id() == inferred.Id()) /* If:2 */{ 
      
      /* If!2 */}  else if ((s.Id() == C_integer.Id()) || 
        ((s.Id() == C_float.Id()) || 
          (s.Id() == C_char.Id()))) /* If:2 */{ 
      if (inferred.Id() == C_any.Id()) /* If:3 */{ 
        PRINC(").Value")
        } else {
        PRINC(".Id()).Value")
        /* If-3 */} 
      /* If!2 */}  else if ((ToType(s.Id()).Included(ToType(C_object.Id())) == CTRUE) || 
        (ToType(s.Id()).Included(ToType(C_primitive.Id())) == CTRUE)) /* If:2 */{ 
      if (inferred.Id() == C_any.Id()) /* If:3 */{ 
        PRINC(")")
        } else {
        PRINC(".Id())")
        /* If-3 */} 
      /* If!2 */}  else if (s.Id() == C_any.Id()) /* If:2 */{ 
      PRINC(".Id()")
      /* If-2 */} 
    } 
  
// The EID go function for: object_post @ class (throw: false) 
func E_Generate_object_post_class (inferred EID,s EID) EID { 
    /*(sm for object_post @ class= void)*/ F_Generate_object_post_class(ToClass(OBJ(inferred)),ToClass(OBJ(s)) )
    return EVOID} 
  
// generic version that applies to everything (s1:infered) => *_prefix(s2:expected)
/* {1} OPT.The go function for: cast_prefix(s1:class,s2:class) [] */
func F_Generate_cast_prefix_class (s1 *ClaireClass ,s2 *ClaireClass ) EID { 
    var Result EID 
    if (s1.Id() == Optimize.C_EID.Id()) /* If:2 */{ 
      Result = F_Generate_eid_prefix_class(s2)
      /* If!2 */}  else if (s1.Id() == C_void.Id()) /* If:2 */{ 
      Result = EID{CNIL.Id(),0}
      /* If!2 */}  else if (s1.Id() == C_integer.Id()) /* If:2 */{ 
      F_Generate_integer_prefix_class(s2)
      Result = EVOID
      /* If!2 */}  else if (s1.Id() == C_float.Id()) /* If:2 */{ 
      F_Generate_float_prefix_class(s2)
      Result = EVOID
      /* If!2 */}  else if (s1.Id() == C_char.Id()) /* If:2 */{ 
      F_Generate_char_prefix_class(s2)
      Result = EVOID
      /* If!2 */}  else if (s1.Id() == C_string.Id()) /* If:2 */{ 
      F_Generate_string_prefix_class(s2)
      Result = EVOID
      } else {
      F_Generate_object_prefix_class(s1,s2)
      Result = EVOID
      /* If-2 */} 
    return Result} 
  
// The EID go function for: cast_prefix @ class (throw: true) 
func E_Generate_cast_prefix_class (s1 EID,s2 EID) EID { 
    return /*(sm for cast_prefix @ class= EID)*/ F_Generate_cast_prefix_class(ToClass(OBJ(s1)),ToClass(OBJ(s2)) )} 
  
// generic version that applies to everything (s1) => *_post(s2)
// s1 is the goType of the expression, s2 is the expected
/* {1} OPT.The go function for: cast_post(s1:class,s2:class) [] */
func F_Generate_cast_post_class (s1 *ClaireClass ,s2 *ClaireClass )  { 
    // procedure body with s =  
if (s1.Id() == Optimize.C_EID.Id()) /* If:2 */{ 
      F_Generate_eid_post_class(s2)
      /* If!2 */}  else if (s1.Id() == C_void.Id()) /* If:2 */{ 
      
      /* If!2 */}  else if ((s1.Id() == C_integer.Id()) || 
        ((s1.Id() == C_float.Id()) || 
          (s1.Id() == C_char.Id()))) /* If:2 */{ 
      F_Generate_native_post_class(s2)
      /* If!2 */}  else if (s1.Id() == C_string.Id()) /* If:2 */{ 
      F_Generate_string_post_class(s2)
      } else {
      F_Generate_object_post_class(s1,s2)
      /* If-2 */} 
    } 
  
// The EID go function for: cast_post @ class (throw: false) 
func E_Generate_cast_post_class (s1 EID,s2 EID) EID { 
    /*(sm for cast_post @ class= void)*/ F_Generate_cast_post_class(ToClass(OBJ(s1)),ToClass(OBJ(s2)) )
    return EVOID} 
  
// *******************************************************************
// *       Part 4: use of language dependent patterns (macros)       *
// *******************************************************************
// when we print an equality, we do not need to_CL !
// id is used to force the identifiability (use = vs equal)
/* {1} OPT.The go function for: equal_exp(c:go_producer,a1:any,pos?:boolean,a2:any,id?:any) [] */
func (c *GenerateGoProducer ) EqualExp (a1 *ClaireAny ,pos_ask *ClaireBoolean ,a2 *ClaireAny ,id_ask *ClaireAny ) EID { 
    var Result EID 
    var g0284I *ClaireBoolean  
    var g0284I_try02852 EID 
    /* and:2 */{ 
      var v_and2 *ClaireBoolean  
      
      var v_and2_try02863 EID 
      /* Let:3 */{ 
        var g0287UU *ClaireClass  
        /* noccur = 1 */
        var g0287UU_try02884 EID 
        g0287UU_try02884 = Language.F_static_type_any(a1)
        /* ERROR PROTECTION INSERTED (g0287UU-v_and2_try02863) */
        if ErrorIn(g0287UU_try02884) {v_and2_try02863 = g0287UU_try02884
        } else {
        g0287UU = ToClass(OBJ(g0287UU_try02884))
        v_and2_try02863 = EID{Equal(g0287UU.Id(),C_string.Id()).Id(),0}
        }
        /* Let-3 */} 
      /* ERROR PROTECTION INSERTED (v_and2-g0284I_try02852) */
      if ErrorIn(v_and2_try02863) {g0284I_try02852 = v_and2_try02863
      } else {
      v_and2 = ToBoolean(OBJ(v_and2_try02863))
      if (v_and2 == CFALSE) {g0284I_try02852 = EID{CFALSE.Id(),0}
      } else /* arg:3 */{ 
        var v_and2_try02894 EID 
        /* Let:4 */{ 
          var g0290UU *ClaireClass  
          /* noccur = 1 */
          var g0290UU_try02915 EID 
          g0290UU_try02915 = Language.F_static_type_any(a2)
          /* ERROR PROTECTION INSERTED (g0290UU-v_and2_try02894) */
          if ErrorIn(g0290UU_try02915) {v_and2_try02894 = g0290UU_try02915
          } else {
          g0290UU = ToClass(OBJ(g0290UU_try02915))
          v_and2_try02894 = EID{Equal(g0290UU.Id(),C_string.Id()).Id(),0}
          }
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (v_and2-g0284I_try02852) */
        if ErrorIn(v_and2_try02894) {g0284I_try02852 = v_and2_try02894
        } else {
        v_and2 = ToBoolean(OBJ(v_and2_try02894))
        if (v_and2 == CFALSE) {g0284I_try02852 = EID{CFALSE.Id(),0}
        } else /* arg:4 */{ 
          g0284I_try02852 = EID{CTRUE.Id(),0}/* arg-4 */} 
        /* arg-3 */} 
      }}
      /* and-2 */} 
    /* ERROR PROTECTION INSERTED (g0284I-Result) */
    if ErrorIn(g0284I_try02852) {Result = g0284I_try02852
    } else {
    g0284I = ToBoolean(OBJ(g0284I_try02852))
    if (g0284I == CTRUE) /* If:2 */{ 
      PRINC("(")
      Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_string.Id(),0}))
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(".Value ")
      F_Generate_sign_equal_boolean(pos_ask)
      PRINC(" ")
      Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{C_string.Id(),0}))
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(".Value)")
      Result = EVOID
      }}
      /* If!2 */}  else if ((c.CharExp_ask(a1) == CTRUE) || 
        (c.CharExp_ask(a2) == CTRUE)) /* If:2 */{ 
      PRINC("(")
      Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_char.Id(),0}))
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(" ")
      F_Generate_sign_equal_boolean(pos_ask)
      PRINC(" ")
      Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{C_char.Id(),0}))
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(")")
      Result = EVOID
      }}
      } else {
      var g0295I *ClaireBoolean  
      var g0295I_try02963 EID 
      /* and:3 */{ 
        var v_and3 *ClaireBoolean  
        
        var v_and3_try02974 EID 
        /* or:4 */{ 
          var v_or4 *ClaireBoolean  
          
          v_or4 = F_boolean_I_any(id_ask)
          if (v_or4 == CTRUE) {v_and3_try02974 = EID{CTRUE.Id(),0}
          } else /* or:5 */{ 
            var v_or4_try02986 EID 
            v_or4_try02986 = Optimize.F_Compile_identifiable_ask_any(a1)
            /* ERROR PROTECTION INSERTED (v_or4-v_and3_try02974) */
            if ErrorIn(v_or4_try02986) {v_and3_try02974 = v_or4_try02986
            } else {
            v_or4 = ToBoolean(OBJ(v_or4_try02986))
            if (v_or4 == CTRUE) {v_and3_try02974 = EID{CTRUE.Id(),0}
            } else /* or:6 */{ 
              var v_or4_try02997 EID 
              v_or4_try02997 = Optimize.F_Compile_identifiable_ask_any(a2)
              /* ERROR PROTECTION INSERTED (v_or4-v_and3_try02974) */
              if ErrorIn(v_or4_try02997) {v_and3_try02974 = v_or4_try02997
              } else {
              v_or4 = ToBoolean(OBJ(v_or4_try02997))
              if (v_or4 == CTRUE) {v_and3_try02974 = EID{CTRUE.Id(),0}
              } else /* or:7 */{ 
                var v_or4_try03008 EID 
                /* Let:8 */{ 
                  var g0301UU *ClaireClass  
                  /* noccur = 1 */
                  var g0301UU_try03029 EID 
                  g0301UU_try03029 = F_Generate_g_sort_any(a1)
                  /* ERROR PROTECTION INSERTED (g0301UU-v_or4_try03008) */
                  if ErrorIn(g0301UU_try03029) {v_or4_try03008 = g0301UU_try03029
                  } else {
                  g0301UU = ToClass(OBJ(g0301UU_try03029))
                  v_or4_try03008 = EID{Equal(g0301UU.Id(),C_float.Id()).Id(),0}
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (v_or4-v_and3_try02974) */
                if ErrorIn(v_or4_try03008) {v_and3_try02974 = v_or4_try03008
                } else {
                v_or4 = ToBoolean(OBJ(v_or4_try03008))
                if (v_or4 == CTRUE) {v_and3_try02974 = EID{CTRUE.Id(),0}
                } else /* or:8 */{ 
                  v_and3_try02974 = EID{CFALSE.Id(),0}/* org-8 */} 
                /* org-7 */} 
              /* org-6 */} 
            /* org-5 */} 
          }}}
          /* or-4 */} 
        /* ERROR PROTECTION INSERTED (v_and3-g0295I_try02963) */
        if ErrorIn(v_and3_try02974) {g0295I_try02963 = v_and3_try02974
        } else {
        v_and3 = ToBoolean(OBJ(v_and3_try02974))
        if (v_and3 == CFALSE) {g0295I_try02963 = EID{CFALSE.Id(),0}
        } else /* arg:4 */{ 
          var v_and3_try03035 EID 
          /* Let:5 */{ 
            var g0304UU *ClaireClass  
            /* noccur = 1 */
            var g0304UU_try03066 EID 
            g0304UU_try03066 = F_Generate_g_sort_any(a1)
            /* ERROR PROTECTION INSERTED (g0304UU-v_and3_try03035) */
            if ErrorIn(g0304UU_try03066) {v_and3_try03035 = g0304UU_try03066
            } else {
            g0304UU = ToClass(OBJ(g0304UU_try03066))
            /* Let:6 */{ 
              var g0305UU *ClaireClass  
              /* noccur = 1 */
              var g0305UU_try03077 EID 
              g0305UU_try03077 = F_Generate_g_sort_any(a2)
              /* ERROR PROTECTION INSERTED (g0305UU-v_and3_try03035) */
              if ErrorIn(g0305UU_try03077) {v_and3_try03035 = g0305UU_try03077
              } else {
              g0305UU = ToClass(OBJ(g0305UU_try03077))
              v_and3_try03035 = EID{Equal(g0304UU.Id(),g0305UU.Id()).Id(),0}
              }
              /* Let-6 */} 
            }
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (v_and3-g0295I_try02963) */
          if ErrorIn(v_and3_try03035) {g0295I_try02963 = v_and3_try03035
          } else {
          v_and3 = ToBoolean(OBJ(v_and3_try03035))
          if (v_and3 == CFALSE) {g0295I_try02963 = EID{CFALSE.Id(),0}
          } else /* arg:5 */{ 
            g0295I_try02963 = EID{CTRUE.Id(),0}/* arg-5 */} 
          /* arg-4 */} 
        }}
        /* and-3 */} 
      /* ERROR PROTECTION INSERTED (g0295I-Result) */
      if ErrorIn(g0295I_try02963) {Result = g0295I_try02963
      } else {
      g0295I = ToBoolean(OBJ(g0295I_try02963))
      if (g0295I == CTRUE) /* If:3 */{ 
        var g0308I *ClaireBoolean  
        var g0308I_try03094 EID 
        /* Let:4 */{ 
          var g0310UU *ClaireType  
          /* noccur = 1 */
          var g0310UU_try03115 EID 
          /* Let:5 */{ 
            var g0312UU *ClaireClass  
            /* noccur = 1 */
            var g0312UU_try03146 EID 
            g0312UU_try03146 = Optimize.F_Compile_stupid_t_any1(a1)
            /* ERROR PROTECTION INSERTED (g0312UU-g0310UU_try03115) */
            if ErrorIn(g0312UU_try03146) {g0310UU_try03115 = g0312UU_try03146
            } else {
            g0312UU = ToClass(OBJ(g0312UU_try03146))
            /* Let:6 */{ 
              var g0313UU *ClaireClass  
              /* noccur = 1 */
              var g0313UU_try03157 EID 
              g0313UU_try03157 = Optimize.F_Compile_stupid_t_any1(a2)
              /* ERROR PROTECTION INSERTED (g0313UU-g0310UU_try03115) */
              if ErrorIn(g0313UU_try03157) {g0310UU_try03115 = g0313UU_try03157
              } else {
              g0313UU = ToClass(OBJ(g0313UU_try03157))
              g0310UU_try03115 = EID{Core.F_glb_class(g0312UU,ToType(g0313UU.Id())).Id(),0}
              }
              /* Let-6 */} 
            }
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (g0310UU-g0308I_try03094) */
          if ErrorIn(g0310UU_try03115) {g0308I_try03094 = g0310UU_try03115
          } else {
          g0310UU = ToType(OBJ(g0310UU_try03115))
          g0308I_try03094 = EID{Equal(g0310UU.Id(),CEMPTY.Id()).Id(),0}
          }
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (g0308I-Result) */
        if ErrorIn(g0308I_try03094) {Result = g0308I_try03094
        } else {
        g0308I = ToBoolean(OBJ(g0308I_try03094))
        if (g0308I == CTRUE) /* If:4 */{ 
          Optimize.F_Compile_warn_void()
          Result = Core.F_tformat_string(MakeString("~S = ~S will fail ! [263]"),2,MakeConstantList(a1,a2))
          } else {
          Result = EID{CFALSE.Id(),0}
          /* If-4 */} 
        }
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC("(")
        /* Let:4 */{ 
          var g0316UU *ClaireClass  
          /* noccur = 1 */
          var g0316UU_try03175 EID 
          g0316UU_try03175 = F_Generate_g_sort_any(a1)
          /* ERROR PROTECTION INSERTED (g0316UU-Result) */
          if ErrorIn(g0316UU_try03175) {Result = g0316UU_try03175
          } else {
          g0316UU = ToClass(OBJ(g0316UU_try03175))
          Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{g0316UU.Id(),0}))
          }
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC(" ")
        F_Generate_sign_equal_boolean(pos_ask)
        PRINC(" ")
        /* Let:4 */{ 
          var g0318UU *ClaireClass  
          /* noccur = 1 */
          var g0318UU_try03195 EID 
          g0318UU_try03195 = F_Generate_g_sort_any(a1)
          /* ERROR PROTECTION INSERTED (g0318UU-Result) */
          if ErrorIn(g0318UU_try03195) {Result = g0318UU_try03195
          } else {
          g0318UU = ToClass(OBJ(g0318UU_try03195))
          Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{g0318UU.Id(),0}))
          }
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC(")")
        Result = EVOID
        }}
        }
        } else {
        PRINC("(Equal(")
        Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_any.Id(),0}))
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC(",")
        Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{C_any.Id(),0}))
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC(") ")
        F_Generate_sign_equal_boolean(pos_ask)
        PRINC(" CTRUE)")
        Result = EVOID
        }}
        /* If-3 */} 
      }
      /* If-2 */} 
    }
    return Result} 
  
// The EID go function for: equal_exp @ go_producer (throw: true) 
func E_Generate_equal_exp_go_producer (c EID,a1 EID,pos_ask EID,a2 EID,id_ask EID) EID { 
    return /*(sm for equal_exp @ go_producer= EID)*/ ToGenerateGoProducer(OBJ(c)).EqualExp(ANY(a1),
      ToBoolean(OBJ(pos_ask)),
      ANY(a2),
      ANY(id_ask) )} 
  
// new: special code for char
// CLAIRE 4 : removed char_exp => g_expression(x,char) should work
/* {1} OPT.The go function for: char_exp?(c:go_producer,x:any) [] */
func (c *GenerateGoProducer ) CharExp_ask (x *ClaireAny ) *ClaireBoolean  { 
    // procedure body with s =  
var Result *ClaireBoolean  
    if (C_char.Id() == x.Isa.Id()) /* If:2 */{ 
      Result = CTRUE
      /* If!2 */}  else if (x.Isa.IsIn(Language.C_Call_method) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0321 *Language.CallMethod   = Language.To_CallMethod(x)
        /* noccur = 1 */
        /* Let:4 */{ 
          var m *ClaireMethod   = g0321.Arg
          /* noccur = 2 */
          Result = MakeBoolean((m.Id() == C_Generate__starnth_1_string_star.Value) || ((m.Id() == C_Generate__starnth_string_star.Value) && 
              (Optimize.C_compiler.Safety >= 2)))
          /* Let-4 */} 
        /* Let-3 */} 
      } else {
      Result = CFALSE
      /* If-2 */} 
    return Result} 
  
// The EID go function for: char_exp? @ go_producer (throw: false) 
func E_Generate_char_exp_ask_go_producer (c EID,x EID) EID { 
    return EID{/*(sm for char_exp? @ go_producer= boolean)*/ ToGenerateGoProducer(OBJ(c)).CharExp_ask(ANY(x) ).Id(),0}} 
  
// reads the member x from an expression self of expected type s
/* {1} OPT.The go function for: c_member(c:go_producer,self:any,s:class,x:property) [] */
func (c *GenerateGoProducer ) CMember (self *ClaireAny ,s *ClaireClass ,x *ClaireProperty ) EID { 
    var Result EID 
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
    return /*(sm for c_member @ go_producer= EID)*/ ToGenerateGoProducer(OBJ(c)).CMember(ANY(self),
      ToClass(OBJ(s)),
      ToProperty(OBJ(x)) )} 
  
// generic for bags
/* {1} OPT.The go function for: bag_expression(c:go_producer,cl:class,l:list,t:type) [] */
func (c *GenerateGoProducer ) BagExpression (cl *ClaireClass ,l *ClaireList ,t *ClaireType ) EID { 
    var Result EID 
    if ((l.Length() == 0) && 
        (cl.Id() != C_tuple.Id())) /* If:2 */{ 
      /* Let:3 */{ 
        var g0323UU *ClaireAny  
        /* noccur = 1 */
        var g0323UU_try03244 EID 
        g0323UU_try03244 = Core.F_CALL(Optimize.C_c_code,ARGS(EID{t.Id(),0},EID{C_object.Id(),0}))
        /* ERROR PROTECTION INSERTED (g0323UU-Result) */
        if ErrorIn(g0323UU_try03244) {Result = g0323UU_try03244
        } else {
        g0323UU = ANY(g0323UU_try03244)
        Result = Core.F_CALL(C_Generate_g_expression,ARGS(g0323UU.ToEID(),EID{C_type.Id(),0}))
        }
        /* Let-3 */} 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(".Empty")
      F_Generate_cap_short_symbol(cl.Name)
      PRINC("()")
      Result = EVOID
      }
      /* If!2 */}  else if ((Equal(t.Id(),CEMPTY.Id()) == CTRUE) || 
        (t.Id() == C_void.Id())) /* If:2 */{ 
      PRINC("Make")
      /* Let:3 */{ 
        var g0325UU *ClaireString  
        /* noccur = 1 */
        if (cl.Id() == C_set.Id()) /* If:4 */{ 
          g0325UU = MakeString("ConstantSet")
          /* If!4 */}  else if (cl.Id() == C_list.Id()) /* If:4 */{ 
          g0325UU = MakeString("ConstantList")
          } else {
          g0325UU = MakeString("Tuple")
          /* If-4 */} 
        F_princ_string(g0325UU)
        /* Let-3 */} 
      PRINC("(")
      Result = F_Generate_args_list_list(l,C_any)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(")")
      Result = EVOID
      }
      } else {
      PRINC("Make")
      F_Generate_cap_short_symbol(cl.Name)
      PRINC("(")
      /* Let:3 */{ 
        var g0326UU *ClaireAny  
        /* noccur = 1 */
        var g0326UU_try03274 EID 
        g0326UU_try03274 = Core.F_CALL(Optimize.C_c_code,ARGS(EID{t.Id(),0},EID{C_object.Id(),0}))
        /* ERROR PROTECTION INSERTED (g0326UU-Result) */
        if ErrorIn(g0326UU_try03274) {Result = g0326UU_try03274
        } else {
        g0326UU = ANY(g0326UU_try03274)
        Result = Core.F_CALL(C_Generate_g_expression,ARGS(g0326UU.ToEID(),EID{C_type.Id(),0}))
        }
        /* Let-3 */} 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(",")
      Result = F_Generate_args_list_list(l,C_any)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(")")
      Result = EVOID
      }}
      /* If-2 */} 
    return Result} 
  
// The EID go function for: bag_expression @ go_producer (throw: true) 
func E_Generate_bag_expression_go_producer (c EID,cl EID,l EID,t EID) EID { 
    return /*(sm for bag_expression @ go_producer= EID)*/ ToGenerateGoProducer(OBJ(c)).BagExpression(ToClass(OBJ(cl)),
      ToList(OBJ(l)),
      ToType(OBJ(t)) )} 
  
// *******************************************************************
// *       Part 5: Utilities :                                       *
// *******************************************************************
// a constant can be evaluated with no cost in an IfThenElse(test,a,b)
/* {1} OPT.The go function for: constant?(self:any) [] */
func F_Generate_constant_ask_any (self *ClaireAny ) *ClaireBoolean  { 
    // use function body compiling 
return  MakeBoolean((self.Isa.IsIn(C_thing) == CTRUE) || 
    (C_boolean.Id() == self.Isa.Id()) || 
    (self.Isa.IsIn(C_Variable) == CTRUE) || 
    (C_string.Id() == self.Isa.Id()) || 
    (self == CNULL) || 
    (Equal(self,CNIL.Id()) == CTRUE) || 
    (Equal(self,CEMPTY.Id()) == CTRUE) || 
    (self.Isa.IsIn(Core.C_global_variable) == CTRUE))
    } 
  
// The EID go function for: constant? @ any (throw: false) 
func E_Generate_constant_ask_any (self EID) EID { 
    return EID{/*(sm for constant? @ any= boolean)*/ F_Generate_constant_ask_any(ANY(self) ).Id(),0}} 
  
// patch: remove protection and conversion layers
/* {1} OPT.The go function for: getC(x:any) [] */
func F_Generate_getC_any (x *ClaireAny ) *ClaireAny  { 
    // procedure body with s =  
var Result *ClaireAny  
    if (x.Isa.IsIn(Core.C_global_variable) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0328 *Core.GlobalVariable   = Core.ToGlobalVariable(x)
        /* noccur = 3 */
        if (Optimize.F_Compile_nativeVar_ask_global_variable(g0328) == CTRUE) /* If:4 */{ 
          Result = g0328.Id()
          } else {
          Result = ANY(Core.F_CALL(C_Generate_to_C,ARGS(EID{Equal(C_arg.Id(),g0328.Id()).Id(),0},EID{Equal(Language.C_iClaire_set_arg.Id(),C_type.Id()).Id(),0})))
          /* If-4 */} 
        /* Let-3 */} 
      } else {
      Result = x
      /* If-2 */} 
    return Result} 
  
// The EID go function for: getC @ any (throw: false) 
func E_Generate_getC_any (x EID) EID { 
    return /*(sm for getC @ any= any)*/ F_Generate_getC_any(ANY(x) ).ToEID()} 
  
// short cut for variable
/* {1} OPT.The go function for: go_range(v:Variable) [] */
func F_Generate_go_range_Variable (v *ClaireVariable ) *ClaireClass  { 
    // use function body compiling 
return  v.Range.Class_I()
    } 
  
// The EID go function for: go_range @ Variable (throw: false) 
func E_Generate_go_range_Variable (v EID) EID { 
    return EID{/*(sm for go_range @ Variable= class)*/ F_Generate_go_range_Variable(To_Variable(OBJ(v)) ).Id(),0}} 
  
// in claire 4, srange(m:method) is gone, replaced by signature => this is temporary method
/* {1} OPT.The go function for: go_signature(m:method) [] */
func F_Generate_go_signature_method (m *ClaireMethod ) *ClaireList  { 
    // procedure body with s =  
var Result *ClaireList  
    /* Let:2 */{ 
      var g0330UU *ClaireList  
      /* noccur = 1 */
      /* Iteration:3 */{ 
        var v_list3 *ClaireList  
        var t *ClaireAny  
        var v_local3 *ClaireAny  
        v_list3 = m.Domain
        g0330UU = CreateList(ToType(C_class.Id()),v_list3.Length())
        for CLcount := 0; CLcount < v_list3.Length(); CLcount++{ 
          t = v_list3.At(CLcount)
          v_local3 = ToTypeExpression(t).Class_I().Id()
          g0330UU.PutAt(CLcount,v_local3)
          } 
        /* Iteration-3 */} 
      Result = g0330UU.AddFast(m.Range.Class_I().Id())
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: go_signature @ method (throw: false) 
func E_Generate_go_signature_method (m EID) EID { 
    return EID{/*(sm for go_signature @ method= list)*/ F_Generate_go_signature_method(ToMethod(OBJ(m)) ).Id(),0}} 
  
// probably should exist elsewhere
/* {1} OPT.The go function for: full_signature(m:method) [] */
func F_Generate_full_signature_method (m *ClaireMethod ) *ClaireList  { 
    // procedure body with s =  
var Result *ClaireList  
    /* Let:2 */{ 
      var g0331UU *ClaireList  
      /* noccur = 1 */
      /* Iteration:3 */{ 
        var v_list3 *ClaireList  
        var t *ClaireAny  
        var v_local3 *ClaireAny  
        v_list3 = m.Domain
        g0331UU = CreateList(ToType(C_type.Id()),v_list3.Length())
        for CLcount := 0; CLcount < v_list3.Length(); CLcount++{ 
          t = v_list3.At(CLcount)
          v_local3 = t
          g0331UU.PutAt(CLcount,v_local3)
          } 
        /* Iteration-3 */} 
      Result = g0331UU.AddFast(m.Range.Id())
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: full_signature @ method (throw: false) 
func E_Generate_full_signature_method (m EID) EID { 
    return EID{/*(sm for full_signature @ method= list)*/ F_Generate_full_signature_method(ToMethod(OBJ(m)) ).Id(),0}} 
  
// print a signature in a AddMethod (goexp.cl)
/* {1} OPT.The go function for: signature!(c:go_producer,l:list<type>) [] */
func (c *GenerateGoProducer ) Signature_I (l *ClaireList ) EID { 
    var Result EID 
    PRINC("Signature(")
    /* Let:2 */{ 
      var g0332UU *ClaireList  
      /* noccur = 1 */
      var g0332UU_try03333 EID 
      /* Iteration:3 */{ 
        var v_list3 *ClaireList  
        var x *ClaireAny  
        var v_local3 *ClaireAny  
        v_list3 = l
        g0332UU_try03333 = EID{CreateList(ToType(CEMPTY.Id()),v_list3.Length()).Id(),0}
        for CLcount := 0; CLcount < v_list3.Length(); CLcount++{ 
          x = v_list3.At(CLcount)
          var v_local3_try03345 EID 
          v_local3_try03345 = Core.F_CALL(Optimize.C_c_code,ARGS(x.ToEID(),EID{C_type.Id(),0}))
          /* ERROR PROTECTION INSERTED (v_local3-g0332UU_try03333) */
          if ErrorIn(v_local3_try03345) {g0332UU_try03333 = v_local3_try03345
          g0332UU_try03333 = v_local3_try03345
          break
          } else {
          v_local3 = ANY(v_local3_try03345)
          ToList(OBJ(g0332UU_try03333)).PutAt(CLcount,v_local3)
          } 
        }
        /* Iteration-3 */} 
      /* ERROR PROTECTION INSERTED (g0332UU-Result) */
      if ErrorIn(g0332UU_try03333) {Result = g0332UU_try03333
      } else {
      g0332UU = ToList(OBJ(g0332UU_try03333))
      Result = F_Generate_args_list_list(g0332UU,C_any)
      }
      /* Let-2 */} 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(")")
    Result = EVOID
    }
    return Result} 
  
// The EID go function for: signature! @ go_producer (throw: true) 
func E_Generate_signature_I_go_producer (c EID,l EID) EID { 
    return /*(sm for signature! @ go_producer= EID)*/ ToGenerateGoProducer(OBJ(c)).Signature_I(ToList(OBJ(l)) )} 
  
// this is a specialized form for list expressions => see if Go should know if a ListObject, ListInt, ListFloat will be used versus generic List
/* {1} OPT.The go function for: g_member(x:any) [] */
func F_Generate_g_member_any (x *ClaireAny ) EID { 
    var Result EID 
    if ((x.Isa.IsIn(Language.C_Call_method) == CTRUE) || 
        ((x.Isa.IsIn(Language.C_Construct) == CTRUE) || 
          ((x.Isa.IsIn(C_Variable) == CTRUE) || 
            ((x.Isa.IsIn(Language.C_Call_slot) == CTRUE) || 
              (x.Isa.IsIn(Language.C_Cast) == CTRUE))))) /* If:2 */{ 
      /* Let:3 */{ 
        var t1 *ClaireType  
        /* noccur = 2 */
        var t1_try03354 EID 
        /* Let:4 */{ 
          var g0336UU *ClaireType  
          /* noccur = 1 */
          var g0336UU_try03375 EID 
          g0336UU_try03375 = Core.F_CALL(Optimize.C_c_type,ARGS(x.ToEID()))
          /* ERROR PROTECTION INSERTED (g0336UU-t1_try03354) */
          if ErrorIn(g0336UU_try03375) {t1_try03354 = g0336UU_try03375
          } else {
          g0336UU = ToType(OBJ(g0336UU_try03375))
          t1_try03354 = EID{g0336UU.At(C_of).Id(),0}
          }
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (t1-Result) */
        if ErrorIn(t1_try03354) {Result = t1_try03354
        } else {
        t1 = ToType(OBJ(t1_try03354))
        if (Core.F_unique_ask_type(t1) == CTRUE) /* If:4 */{ 
          Result = Core.F_the_type(t1)
          } else {
          Result = EID{C_any.Id(),0}
          /* If-4 */} 
        }
        /* Let-3 */} 
      } else {
      Result = EID{C_any.Id(),0}
      /* If-2 */} 
    return RangeCheck(ToType(C_class.Id()),Result)} 
  
// The EID go function for: g_member @ any (throw: true) 
func E_Generate_g_member_any (x EID) EID { 
    return /*(sm for g_member @ any= EID)*/ F_Generate_g_member_any(ANY(x) )} 
  
// associated prefix & post (for list only-> because Values is not defined at ClaireList level)
/* {1} OPT.The go function for: list_cast_values(sbag:class,gmem:class) [] */
func F_Generate_list_cast_values_class (sbag *ClaireClass ,gmem *ClaireClass )  { 
    // procedure body with s =  
if (sbag.Id() == C_list.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var short *ClaireString  
        /* noccur = 1 */
        if (gmem.Id() == C_integer.Id()) /* If:4 */{ 
          short = MakeString("I")
          /* If!4 */}  else if (gmem.Id() == C_float.Id()) /* If:4 */{ 
          short = MakeString("F")
          } else {
          short = MakeString("O")
          /* If-4 */} 
        PRINC(".Values")
        F_princ_string(short)
        PRINC("()")
        /* Let-3 */} 
      } else {
      PRINC(".Values")
      /* If-2 */} 
    } 
  
// The EID go function for: list_cast_values @ class (throw: false) 
func E_Generate_list_cast_values_class (sbag EID,gmem EID) EID { 
    /*(sm for list_cast_values @ class= void)*/ F_Generate_list_cast_values_class(ToClass(OBJ(sbag)),ToClass(OBJ(gmem)) )
    return EVOID} 
  
// regular for sets
//
// this method does nothing. It used to check if a name could create a naming conflict.
// we keep it until we have tested that it is safe to remove it
// we could use a stack of names that have been used (reset for each method)
/* {1} OPT.The go function for: check_var(self:string) [] */
func F_Generate_check_var_string (self *ClaireString ) *ClaireString  { 
    // use function body compiling 
return  F_append_string(self,F_string_I_integer(Optimize.C_OPT.Level))
    } 
  
// The EID go function for: check_var @ string (throw: false) 
func E_Generate_check_var_string (self EID) EID { 
    return EID{/*(sm for check_var @ string= string)*/ F_Generate_check_var_string(ToString(OBJ(self)) ).Id(),0}} 
  
/* {1} OPT.The go function for: build_Variable(s:string,t:any) [] */
func F_Generate_build_Variable_string (s *ClaireString ,t *ClaireAny ) *ClaireVariable  { 
    // use function body compiling 
return  Optimize.F_Compile_Variable_I_symbol(Core.F_symbol_I_string2(s),0,t)
    } 
  
// The EID go function for: build_Variable @ string (throw: false) 
func E_Generate_build_Variable_string (s EID,t EID) EID { 
    return EID{/*(sm for build_Variable @ string= Variable)*/ F_Generate_build_Variable_string(ToString(OBJ(s)),ANY(t) ).Id(),0}} 
  
// use a variable v with inferred type when expected : add the casts
/* {1} OPT.The go function for: use_variable(v:string,expected:class,inferred:class) [] */
func F_Generate_use_variable_string (v *ClaireString ,expected *ClaireClass ,inferred *ClaireClass ) EID { 
    var Result EID 
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
    return /*(sm for use_variable @ string= EID)*/ F_Generate_use_variable_string(ToString(OBJ(v)),ToClass(OBJ(expected)),ToClass(OBJ(inferred)) )} 
  