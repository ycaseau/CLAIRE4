/***** CLAIRE Compilation of file /Users/ycaseau/Dropbox/src/clairev4.03/src/compile/osystem.cl 
         [version 4.0.04 / safety 5] Saturday 01-01-2022 16:47:21 *****/

package Optimize
import (_ "fmt"
	. "Kernel"
	"Core"
	"Language"
	"Reader"
)

//-------- dumb function to prevent import errors --------
func import_g0000() { 
    _ = Core.It
    _ = Language.It
    _ = Reader.It
    } 
  
  
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
//
// import
// Compile/index :: Kernel/index   (1) should not be needed (inherited by iClaire)
// Compile/typing :: Kernel/typing
// where to find the CLAIRE libraries
/* The go function for: home(_CL_obj:void) [status=0] */
func F_home_void () *ClaireString  { 
    return  F_getenv_string(MakeString("CLAIRE_HOME"))
    } 
  
// The EID go function for: home @ void (throw: false) 
func E_home_void (_CL_obj EID) EID { 
    return EID{F_home_void( ).Id(),0}} 
  
// TO CHANGE -> simply read claire_home
// CLAIRE4 uses EID where CLAIRE used C++ OID (integer representation)
// used as a marker for form EID
// ******************************************************************
// *    Part 1: General Global Variables and Properties             *
// ******************************************************************
// we use an optimizer object with all the necessary resources
// they are all private.
// update on strings   v3.3.46
// The meta_compiler contains the definition of the compiler flags and slots
// that are important for the user. Other stuff is hidden in OPT
// number of methods compiled with EID (error handling)
// code producer are defined in Generate
// but the stub is define in Optimize to have access to current_file
// name of the file being compiled
// we use a global variable to hide the indirection through the producer
// this is kept in CLAIRE 4.0 so that the C++ compiler could be re-introduced
// new in CLAIRE4: create an automated comment
// the three variables that are used in the main files
// safety:
//       0  -> super-safe (keep assertion)
//       1  -> safe (regular)
//       2  -> we trus typing
//       3  -> no overflow checking (integer & arrays)
// re-definable items for bootstrap modifications
// Compile/make_float_function :: property(Core/open = 3)
// Compile/c_expression :: property(Core/open = 3)
// other useful properties shared between Optimize & Generate
// Optimizer version of sorts
// code with strict (stupid) type
// new: allow future overload !!
// compiler instantiation
// fast instantiation if all any slots are known
// how to compile a type expression
// these are the classes defined especially for this module
// Compile/to_CL <: Optimized_instruction(arg:any,set_arg:class)
// Compile/to_C <: Optimized_instruction(arg:any,set_arg:class)
// was to_C()
// Patterns are calls p(X) that are seen as a type expression
// the tuple is made into a list
// OPT contains all the parameters for the optimizer
// pragma for the compiler  => MOVED TO LANGUAGE in CLAIRE 4
// this pragma tells to compile with full safety (include arithmetic checks)
/* The go function for: safe(x:any) [status=0] */
func F_safe_any (x *ClaireAny ) *ClaireAny  { 
    return  x
    } 
  
// The EID go function for: safe @ any (throw: false) 
func E_safe_any (x EID) EID { 
    return F_safe_any(ANY(x) ).ToEID()} 
  
/* The go function for: safe_any_type */
func F_safe_any_type (x *ClaireType ) EID { 
    var Result EID 
    Result = EID{x.Id(),0}
    return Result} 
  
  
// The dual EID go function for: "safe_any_type" 
func E_safe_any_type (x EID) EID { 
    return F_safe_any_type(ToType(OBJ(x)))} 
  
// ******************************************************************
// *    Part 2: The defaults for c_type, c_code and c_sort          *
// ******************************************************************
// basic type inference
/* The go function for: c_type(self:any) [status=1] */
func F_c_type_any (self *ClaireAny ) EID { 
    // eid body s = type
    var Result EID 
    if (self.Isa.IsIn(C_Variable) == CTRUE) { 
      { var g0001 *ClaireVariable   = To_Variable(self)
        { var r *ClaireAny   = Core.F_get_property(C_range,ToObject(g0001.Id()))
          if ((r == CNULL) || 
              (r == C_EID.Id())) { 
            Result = EID{C_any.Id(),0}
            } else {
            var g0012I *ClaireBoolean  
            if (r.Isa.IsIn(C_Union) == CTRUE) { 
              { var g0002 *ClaireUnion   = To_Union(r)
                g0012I = Equal(g0002.T1.Id(),CEMPTY.Id())
                } 
              } else {
              g0012I = CFALSE
              } 
            if (g0012I == CTRUE) { 
              Result = EID{To_Union(To_Union(r).T2.Id()).T2.Id(),0}
              } else {
              Result = EID{F_Optimize_ptype_type(ToType(r)).Id(),0}
              } 
            } 
          } 
        } 
      }  else if (self.Isa.IsIn(Core.C_global_variable) == CTRUE) { 
      { var g0003 *Core.GlobalVariable   = Core.ToGlobalVariable(self)
        { var r *ClaireType   = g0003.Range
          if (F_boolean_I_any(r.Id()) == CTRUE) { 
            Result = EID{r.Id(),0}
            } else {
            Result = EID{MakeConstantSet(g0003.Value).Id(),0}
            } 
          } 
        } 
      }  else if (self.Isa.IsIn(C_unbound_symbol) == CTRUE) { 
      { var g0004 *ClaireUnboundSymbol   = ToUnboundSymbol(self)
        Result = Core.F_CALL(C_Compile_Cerror,ARGS(EID{MakeString("[215] the symbol ~A is unbound_symbol").Id(),0},EID{g0004.Name.Id(),0}))
        } 
      }  else if (self.Isa.IsIn(C_error) == CTRUE) { 
      Result = EID{CEMPTY.Id(),0}
      }  else if (self.Isa.IsIn(Language.C_Update) == CTRUE) { 
      { var g0006 *Language.Update   = Language.To_Update(self)
        Result = Core.F_CALL(C_c_type,ARGS(g0006.Value.ToEID()))
        } 
      }  else if (self.Isa.IsIn(Language.C_Construct) == CTRUE) { 
      { var g0007 *Language.Construct   = Language.To_Construct(self)
        if ((g0007.Isa.IsIn(Language.C_List) != CTRUE) && 
            (g0007.Isa.IsIn(Language.C_Set) != CTRUE)) { 
          Result = EID{C_any.Id(),0}
          } else {
          { var _Zres *ClaireType   = ToType(CEMPTY.Id())
            { 
              var _Zx *ClaireAny  
              _ = _Zx
              Result= EID{CFALSE.Id(),0}
              var _Zx_support *ClaireList  
              _Zx_support = g0007.Args
              _Zx_len := _Zx_support.Length()
              for i_it := 0; i_it < _Zx_len; i_it++ { 
                _Zx = _Zx_support.At(i_it)
                var loop_1 EID 
                _ = loop_1
                if (F_boolean_I_any(_Zres.Id()) == CTRUE) { 
                  var try_2 EID 
                  { var arg_3 *ClaireClass  
                    var try_4 EID 
                    { var arg_5 *ClaireType  
                      var try_6 EID 
                      try_6 = Core.F_CALL(C_c_type,ARGS(_Zx.ToEID()))
                      if ErrorIn(try_6) {try_4 = try_6
                      } else {
                      arg_5 = ToType(OBJ(try_6))
                      try_4 = EID{arg_5.Class_I().Id(),0}
                      }
                      } 
                    if ErrorIn(try_4) {try_2 = try_4
                    } else {
                    arg_3 = ToClass(OBJ(try_4))
                    try_2 = EID{Core.F_meet_class(ToClass(_Zres.Id()),arg_3).Id(),0}
                    }
                    } 
                  if ErrorIn(try_2) {loop_1 = try_2
                  } else {
                  _Zres = ToType(OBJ(try_2))
                  loop_1 = EID{_Zres.Id(),0}
                  }
                  } else {
                  var try_7 EID 
                  { var arg_8 *ClaireType  
                    var try_9 EID 
                    try_9 = Core.F_CALL(C_c_type,ARGS(_Zx.ToEID()))
                    if ErrorIn(try_9) {try_7 = try_9
                    } else {
                    arg_8 = ToType(OBJ(try_9))
                    try_7 = EID{arg_8.Class_I().Id(),0}
                    }
                    } 
                  if ErrorIn(try_7) {loop_1 = try_7
                  } else {
                  _Zres = ToType(OBJ(try_7))
                  loop_1 = EID{_Zres.Id(),0}
                  }
                  } 
                if ErrorIn(loop_1) {Result = loop_1
                break
                } else {
                }
                } 
              } 
            if !ErrorIn(Result) {
            { var arg_10 *ClaireClass  
              if (g0007.Isa.IsIn(Language.C_Set) == CTRUE) { 
                arg_10 = C_set
                } else {
                arg_10 = C_list
                } 
              Result = EID{Core.F_nth_class1(arg_10,_Zres).Id(),0}
              } 
            }
            } 
          } 
        } 
      }  else if (self.Isa.IsIn(C_Instruction) == CTRUE) { 
      { var g0010 *ClaireInstruction   = To_Instruction(self)
        Result = ToException(Core.C_general_error.Make(MakeString("c_type of ~S is not defined").Id(),MakeConstantList(g0010.Id().Isa.Id()).Id())).Close()
        } 
      } else {
      Result = EID{MakeConstantSet(self).Id(),0}
      } 
    return Result} 
  
// The EID go function for: c_type @ any (throw: true) 
func E_c_type_any (self EID) EID { 
    return F_c_type_any(ANY(self) )} 
  
// compile into a sort and checks strict type matching (naive/stupid)
/* The go function for: Compile/c_strict_code(x:any,s:class) [status=1] */
func F_Compile_c_strict_code_any (x *ClaireAny ,s *ClaireClass ) EID { 
    // eid body s = any
    var Result EID 
    { var arg_1 *ClaireAny  
      var try_2 EID 
      try_2 = Core.F_CALL(C_c_code,ARGS(x.ToEID(),EID{s.Id(),0}))
      if ErrorIn(try_2) {Result = try_2
      } else {
      arg_1 = ANY(try_2)
      Result = Core.F_CALL(C_Compile_c_strict_check,ARGS(arg_1.ToEID(),EID{s.Id(),0}))
      }
      } 
    return Result} 
  
// The EID go function for: Compile/c_strict_code @ any (throw: true) 
func E_Compile_c_strict_code_any (x EID,s EID) EID { 
    return F_Compile_c_strict_code_any(ANY(x),ToClass(OBJ(s)) )} 
  
// CLAIRE 4: introduce C_cast so that psort(x) is what is expected (s)
/* The go function for: Compile/c_strict_check(x:any,s:class) [status=1] */
func F_Compile_c_strict_check_any (x *ClaireAny ,s *ClaireClass ) EID { 
    // eid body s = any
    var Result EID 
    var g0013I *ClaireBoolean  
    var try_1 EID 
    { 
      var v_and2 *ClaireBoolean  
      
      v_and2 = s.IsIn(C_object)
      if (v_and2 == CFALSE) {try_1 = EID{CFALSE.Id(),0}
      } else { 
        var try_2 EID 
        { var arg_3 *ClaireBoolean  
          var try_4 EID 
          { var arg_5 *ClaireClass  
            var try_6 EID 
            try_6 = Language.F_static_type_any(x)
            if ErrorIn(try_6) {try_4 = try_6
            } else {
            arg_5 = ToClass(OBJ(try_6))
            try_4 = EID{arg_5.IsIn(s).Id(),0}
            }
            } 
          if ErrorIn(try_4) {try_2 = try_4
          } else {
          arg_3 = ToBoolean(OBJ(try_4))
          try_2 = EID{arg_3.Not.Id(),0}
          }
          } 
        if ErrorIn(try_2) {try_1 = try_2
        } else {
        v_and2 = ToBoolean(OBJ(try_2))
        if (v_and2 == CFALSE) {try_1 = EID{CFALSE.Id(),0}
        } else { 
          try_1 = EID{CTRUE.Id(),0}} 
        } 
      }
      } 
    if ErrorIn(try_1) {Result = try_1
    } else {
    g0013I = ToBoolean(OBJ(try_1))
    if (g0013I == CTRUE) { 
      { var _CL_obj *Compile_CCast   = To_Compile_CCast(new(Compile_CCast).Is(C_Compile_C_cast))
        _CL_obj.Arg = x
        _CL_obj.SetArg = s
        Result = EID{_CL_obj.Id(),0}
        } 
      } else {
      Result = x.ToEID()
      } 
    }
    return Result} 
  
// The EID go function for: Compile/c_strict_check @ any (throw: true) 
func E_Compile_c_strict_check_any (x EID,s EID) EID { 
    return F_Compile_c_strict_check_any(ANY(x),ToClass(OBJ(s)) )} 
  
// using conversions. s is a sort or void (we do not need the value).
// note: we need s to be the precise sort for C++
// the is the default version that uses c_code(x)/ c_sort(x)
// in CLAIRE 4, we do not generate conversion at optim time
/* The go function for: c_code(x:any,s:class) [status=1] */
func F_c_code_any1 (x *ClaireAny ,s *ClaireClass ) EID { 
    // eid body s = any
    var Result EID 
    { var y *ClaireAny  
      var try_1 EID 
      if (x.Isa.IsIn(Language.C_Call) == CTRUE) { 
        { var g0014 *Language.Call   = Language.To_Call(x)
          try_1 = F_Optimize_c_code_call_Call(g0014,s)
          } 
        } else {
        try_1 = Core.F_CALL(C_c_code,ARGS(x.ToEID()))
        } 
      if ErrorIn(try_1) {Result = try_1
      } else {
      y = ANY(try_1)
      { var z *ClaireClass  
        var try_2 EID 
        try_2 = Core.F_CALL(C_Compile_c_sort,ARGS(y.ToEID()))
        if ErrorIn(try_2) {Result = try_2
        } else {
        z = ToClass(OBJ(try_2))
        if ((s.Id() == C_void.Id()) || 
            (z.Id() == s.Id())) { 
          var g0017I *ClaireBoolean  
          { 
            var v_and5 *ClaireBoolean  
            
            v_and5 = Equal(s.Id(),C_void.Id())
            if (v_and5 == CFALSE) {g0017I = CFALSE
            } else { 
              if (x.Isa.IsIn(Language.C_Call) == CTRUE) { 
                { var g0016 *Language.Call   = Language.To_Call(x)
                  v_and5 = Equal(g0016.Selector.Id(),C__equal.Id())
                  } 
                } else {
                v_and5 = CFALSE
                } 
              if (v_and5 == CFALSE) {g0017I = CFALSE
              } else { 
                g0017I = CTRUE} 
              } 
            } 
          if (g0017I == CTRUE) { 
            F_Compile_warn_void()
            Core.F_tformat_string(MakeString("-- Equality meant as an assignment: ~S [264]\n"),2,MakeConstantList(x))
            } 
          Result = y.ToEID()
          } else {
          Result = y.ToEID()
          } 
        }
        } 
      }
      } 
    return Result} 
  
// The EID go function for: c_code @ list<type_expression>(any, class) (throw: true) 
func E_c_code_any1 (x EID,s EID) EID { 
    return F_c_code_any1(ANY(x),ToClass(OBJ(s)) )} 
  
// basic code generation
// c_code without a sort parameter means that we do not care about the resulting sort,
// which will be checked later on using c_sort
/* The go function for: c_code(self:any) [status=1] */
func F_c_code_any2 (self *ClaireAny ) EID { 
    // eid body s = any
    var Result EID 
    if (self.Isa.IsIn(C_unbound_symbol) == CTRUE) { 
      { var g0018 *ClaireUnboundSymbol   = ToUnboundSymbol(self)
        Result = Core.F_CALL(C_Compile_Cerror,ARGS(EID{MakeString("[215] the symbol ~A is unbound_symbol").Id(),0},EID{g0018.Name.Id(),0}))
        } 
      }  else if (self.Isa.IsIn(C_Variable) == CTRUE) { 
      { var g0019 *ClaireVariable   = To_Variable(self)
        Result = EID{g0019.Id(),0}
        } 
      }  else if (self.Isa.IsIn(Core.C_global_variable) == CTRUE) { 
      { var g0020 *Core.GlobalVariable   = Core.ToGlobalVariable(self)
        F_Optimize_c_register_object(ToObject(g0020.Id()))
        Result = EID{g0020.Id(),0}
        } 
      }  else if (self.Isa.IsIn(Language.C_Optimized_instruction) == CTRUE) { 
      { var g0021 *Language.OptimizedInstruction   = Language.To_OptimizedInstruction(self)
        Result = EID{g0021.Id(),0}
        } 
      }  else if (self.Isa.IsIn(C_Instruction) == CTRUE) { 
      { var g0022 *ClaireInstruction   = To_Instruction(self)
        Result = Core.F_CALL(C_Compile_Cerror,ARGS(EID{MakeString("[internal] c_code(~S) should have a parameter").Id(),0},EID{g0022.Id(),0}))
        } 
      }  else if (C_set.Id() == self.Isa.Id()) { 
      { var g0023 *ClaireSet   = ToSet(self)
        if (F_boolean_I_any(g0023.Id()) == CTRUE) { 
          { var x *Language.Set  
            { var _CL_obj *Language.Set   = Language.To_Set(new(Language.Set).Is(Language.C_Set))
              _CL_obj.Args = g0023.List_I()
              x = _CL_obj
              } 
            if (ToList(g0023.Id()).Of().Id() != C_void.Id()) { 
              x.Of = ToList(g0023.Id()).Of()
              } 
            Result = Core.F_CALL(C_c_code,ARGS(EID{x.Id(),0}))
            } 
          } else {
          Result = EID{g0023.Id(),0}
          } 
        } 
      }  else if (self.Isa.IsIn(C_list) == CTRUE) { 
      { var g0024 *ClaireList   = ToList(self)
        if (g0024.Length() != 0) { 
          { var x *Language.List  
            { var _CL_obj *Language.List   = Language.To_List(new(Language.List).Is(Language.C_List))
              _CL_obj.Args = g0024
              x = _CL_obj
              } 
            if (g0024.Of().Id() != C_void.Id()) { 
              x.Of = g0024.Of()
              } 
            Result = Core.F_CALL(C_c_code,ARGS(EID{x.Id(),0}))
            } 
          } else {
          Result = EID{g0024.Id(),0}
          } 
        } 
      }  else if (C_tuple.Id() == self.Isa.Id()) { 
      { var g0025 *ClaireTuple   = ToTuple(self)
        { var arg_1 *Language.Tuple  
          { var _CL_obj *Language.Tuple   = Language.To_Tuple(new(Language.Tuple).Is(Language.C_Tuple))
            _CL_obj.Args = g0025.List_I()
            arg_1 = _CL_obj
            } 
          Result = Core.F_CALL(C_c_code,ARGS(EID{arg_1.Id(),0}))
          } 
        } 
      } else {
      if (self.Isa.IsIn(C_thing) == CTRUE) { 
        Core.F_CALL(C_Optimize_c_register,ARGS(self.ToEID()))
        } 
      Result = self.ToEID()
      } 
    return Result} 
  
// The EID go function for: c_code @ list<type_expression>(any) (throw: true) 
func E_c_code_any2 (self EID) EID { 
    return F_c_code_any2(ANY(self) )} 
  
// suggestion for claire4 : get rid of c_sort
/* The go function for: get_sort(self:any) [status=1] */
func F_Optimize_get_sort_any (self *ClaireAny ) EID { 
    // eid body s = class
    var Result EID 
    Result = Language.F_static_type_any(self)
    return Result} 
  
// The EID go function for: get_sort @ any (throw: true) 
func E_Optimize_get_sort_any (self EID) EID { 
    return F_Optimize_get_sort_any(ANY(self) )} 
  
// gives the sort of a compiled expression (does not apply to instructions that
// have a direct c_code(x,s)
// v2.4.9: special type => special sorts !!!
/* The go function for: Compile/c_sort(self:any) [status=1] */
func F_Compile_c_sort_any (self *ClaireAny ) EID { 
    // eid body s = class
    var Result EID 
    if (self.Isa.IsIn(Core.C_global_variable) == CTRUE) { 
      { var g0027 *Core.GlobalVariable   = Core.ToGlobalVariable(self)
        if (F_Compile_nativeVar_ask_global_variable(g0027) == CTRUE) { 
          if (Equal(g0027.Range.Id(),CEMPTY.Id()) == CTRUE) { 
            Result = EID{F_Compile_osort_any(g0027.Value.Isa.Id()).Id(),0}
            } else {
            Result = EID{F_Compile_osort_any(g0027.Range.Id()).Id(),0}
            } 
          } else {
          Result = EID{C_any.Id(),0}
          } 
        } 
      }  else if (self.Isa.IsIn(C_Instruction) == CTRUE) { 
      { var g0028 *ClaireInstruction   = To_Instruction(self)
        if (g0028.Isa.IsIn(C_Variable) == CTRUE) { 
          { var g0029 *ClaireVariable   = To_Variable(g0028.Id())
            Result = EID{F_sort_Variable(g0029).Id(),0}
            } 
          }  else if (g0028.Isa.IsIn(Language.C_Assign) == CTRUE) { 
          { var g0030 *Language.Assign   = Language.To_Assign(g0028.Id())
            Result = EID{F_sort_Variable(To_Variable(g0030.ClaireVar)).Id(),0}
            } 
          }  else if (g0028.Isa.IsIn(Language.C_Call) == CTRUE) { 
          { var g0031 *Language.Call   = Language.To_Call(g0028.Id())
            Result = EID{F_Compile_osort_any(F_Optimize_selector_psort_Call(g0031).Id()).Id(),0}
            } 
          }  else if (g0028.Isa.IsIn(Language.C_Call_method) == CTRUE) { 
          { var g0032 *Language.CallMethod   = Language.To_CallMethod(g0028.Id())
            if ((g0032.Arg.Selector.Id() == Core.C_externC.Id()) && 
                (g0032.Args.Length() == 2)) { 
              Result = EID{F_Compile_psort_any(g0032.Args.At(2-1)).Id(),0}
              } else {
              Result = F_Optimize_c_srange_method(g0032.Arg)
              } 
            } 
          }  else if (g0028.Isa.IsIn(Language.C_Call_slot) == CTRUE) { 
          { var g0033 *Language.CallSlot   = Language.To_CallSlot(g0028.Id())
            Result = EID{g0033.Selector.Srange.Id(),0}
            } 
          }  else if (g0028.Isa.IsIn(Language.C_Call_table) == CTRUE) { 
          Result = EID{C_any.Id(),0}
          }  else if (g0028.Isa.IsIn(Language.C_Call_array) == CTRUE) { 
          { var g0035 *Language.CallArray   = Language.To_CallArray(g0028.Id())
            if (g0035.Test == C_float.Id()) { 
              Result = EID{C_float.Id(),0}
              } else {
              Result = EID{C_any.Id(),0}
              } 
            } 
          }  else if (g0028.Isa.IsIn(Language.C_Definition) == CTRUE) { 
          Result = EID{C_object.Id(),0}
          }  else if (g0028.Isa.IsIn(C_Compile_C_cast) == CTRUE) { 
          { var g0037 *Compile_CCast   = To_Compile_CCast(g0028.Id())
            Result = EID{g0037.SetArg.Id(),0}
            } 
          }  else if (g0028.Isa.IsIn(Language.C_Update) == CTRUE) { 
          { var g0038 *Language.Update   = Language.To_Update(g0028.Id())
            Result = Core.F_CALL(C_Compile_c_sort,ARGS(g0038.Value.ToEID()))
            } 
          }  else if (g0028.Isa.IsIn(Language.C_If) == CTRUE) { 
          { var g0039 *Language.If   = Language.To_If(g0028.Id())
            { var arg_1 *ClaireClass  
              var try_2 EID 
              { var arg_3 *ClaireAny  
                var try_5 EID 
                try_5 = Core.F_CALL(C_Compile_c_sort,ARGS(g0039.Arg.ToEID()))
                if ErrorIn(try_5) {try_2 = try_5
                } else {
                arg_3 = ANY(try_5)
                { var arg_4 *ClaireAny  
                  var try_6 EID 
                  try_6 = Core.F_CALL(C_Compile_c_sort,ARGS(g0039.Other.ToEID()))
                  if ErrorIn(try_6) {try_2 = try_6
                  } else {
                  arg_4 = ANY(try_6)
                  try_2 = EID{Core.F_meet_class(ToClass(arg_3),ToClass(arg_4)).Id(),0}
                  }
                  } 
                }
                } 
              if ErrorIn(try_2) {Result = try_2
              } else {
              arg_1 = ToClass(OBJ(try_2))
              Result = EID{F_Compile_psort_any(arg_1.Id()).Id(),0}
              }
              } 
            } 
          }  else if (g0028.Isa.IsIn(Language.C_Handle) == CTRUE) { 
          { var g0040 *Language.ClaireHandle   = Language.To_ClaireHandle(g0028.Id())
            { var arg_7 *ClaireClass  
              var try_8 EID 
              { var arg_9 *ClaireAny  
                var try_11 EID 
                try_11 = Core.F_CALL(C_Compile_c_sort,ARGS(g0040.Arg.ToEID()))
                if ErrorIn(try_11) {try_8 = try_11
                } else {
                arg_9 = ANY(try_11)
                { var arg_10 *ClaireAny  
                  var try_12 EID 
                  try_12 = Core.F_CALL(C_Compile_c_sort,ARGS(g0040.Other.ToEID()))
                  if ErrorIn(try_12) {try_8 = try_12
                  } else {
                  arg_10 = ANY(try_12)
                  try_8 = EID{Core.F_meet_class(ToClass(arg_9),ToClass(arg_10)).Id(),0}
                  }
                  } 
                }
                } 
              if ErrorIn(try_8) {Result = try_8
              } else {
              arg_7 = ToClass(OBJ(try_8))
              Result = EID{F_Compile_psort_any(arg_7.Id()).Id(),0}
              }
              } 
            } 
          }  else if (g0028.Isa.IsIn(Language.C_Let) == CTRUE) { 
          { var g0041 *Language.Let   = Language.To_Let(g0028.Id())
            Result = Core.F_CALL(C_Compile_c_sort,ARGS(g0041.Arg.ToEID()))
            } 
          }  else if (g0028.Isa.IsIn(Language.C_Return) == CTRUE) { 
          Result = EID{C_any.Id(),0}
          }  else if (g0028.Isa.IsIn(Language.C_List) == CTRUE) { 
          Result = EID{C_object.Id(),0}
          }  else if (g0028.Isa.IsIn(Language.C_Set) == CTRUE) { 
          Result = EID{C_object.Id(),0}
          }  else if (g0028.Isa.IsIn(Language.C_Tuple) == CTRUE) { 
          Result = EID{C_object.Id(),0}
          }  else if (g0028.Isa.IsIn(Language.C_Construct) == CTRUE) { 
          Result = EID{C_any.Id(),0}
          }  else if (g0028.Isa.IsIn(Language.C_Gassign) == CTRUE) { 
          Result = EID{C_any.Id(),0}
          }  else if (g0028.Isa.IsIn(Language.C_Super) == CTRUE) { 
          Result = EID{C_any.Id(),0}
          }  else if (g0028.Isa.IsIn(Language.C_For) == CTRUE) { 
          Result = EID{C_any.Id(),0}
          }  else if (g0028.Isa.IsIn(Language.C_Exists) == CTRUE) { 
          { var g0050 *Language.Exists   = Language.To_Exists(g0028.Id())
            if (g0050.Other == CNULL) { 
              Result = EID{C_any.Id(),0}
              } else {
              Result = EID{C_object.Id(),0}
              } 
            } 
          }  else if (g0028.Isa.IsIn(Language.C_Iteration) == CTRUE) { 
          Result = EID{C_object.Id(),0}
          }  else if (g0028.Isa.IsIn(Language.C_And) == CTRUE) { 
          Result = EID{C_boolean.Id(),0}
          }  else if (g0028.Isa.IsIn(Language.C_Or) == CTRUE) { 
          Result = EID{C_boolean.Id(),0}
          }  else if (g0028.Isa.IsIn(Language.C_While) == CTRUE) { 
          Result = EID{C_any.Id(),0}
          }  else if (g0028.Isa.IsIn(Language.C_Do) == CTRUE) { 
          { var g0055 *Language.Do   = Language.To_Do(g0028.Id())
            { var arg_13 *ClaireAny  
              var try_14 EID 
              try_14 = Core.F_last_list(g0055.Args)
              if ErrorIn(try_14) {Result = try_14
              } else {
              arg_13 = ANY(try_14)
              Result = Core.F_CALL(C_Compile_c_sort,ARGS(arg_13.ToEID()))
              }
              } 
            } 
          } else {
          Result = ToException(Core.C_general_error.Make(MakeString("[internal] c_sort is not implemented for ~S").Id(),MakeConstantList(g0028.Id().Isa.Id()).Id())).Close()
          } 
        } 
      }  else if (C_float.Id() == self.Isa.Id()) { 
      Result = EID{C_float.Id(),0}
      } else {
      { var arg_15 *ClaireType  
        var try_16 EID 
        try_16 = Core.F_CALL(C_c_type,ARGS(self.ToEID()))
        if ErrorIn(try_16) {Result = try_16
        } else {
        arg_15 = ToType(OBJ(try_16))
        Result = EID{F_Compile_psort_any(arg_15.Id()).Id(),0}
        }
        } 
      } 
    return Result} 
  
// The EID go function for: Compile/c_sort @ any (throw: true) 
func E_Compile_c_sort_any (self EID) EID { 
    return F_Compile_c_sort_any(ANY(self) )} 
  
// for the special compiler properties, we need to tell the sort of the optimized
// form
/* The go function for: selector_psort(self:Call) [status=0] */
func F_Optimize_selector_psort_Call (self *Language.Call ) *ClaireClass  { 
    // procedure body with s = class
    var Result *ClaireClass  
    { var p *ClaireProperty   = self.Selector
      if ((p.Id() == Core.C_mClaire_base_I.Id()) || 
          (p.Id() == Core.C_mClaire_index_I.Id())) { 
        Result = C_integer
        }  else if (p.Id() == C_Compile_anyObject_I.Id()) { 
        Result = ToClass(self.Args.At(1-1))
        }  else if (p.Id() == C_Compile_object_I.Id()) { 
        Result = ToClass(self.Args.At(2-1))
        } else {
        Result = C_any
        } 
      } 
    return Result} 
  
// The EID go function for: selector_psort @ Call (throw: false) 
func E_Optimize_selector_psort_Call (self EID) EID { 
    return EID{F_Optimize_selector_psort_Call(Language.To_Call(OBJ(self)) ).Id(),0}} 
  
// ******************************************************************
// *    Part 3: g_throw and status(m:method)                        *
// ******************************************************************
// NEW in claire4 : optimization when compiler.safety is high may prevent throwing exceptions
// these two variabler are used for cross-compiling, when the status changes from the existing(compiled) version to the
// new one being compiled
// NEW in claire 4, because error handling is mananaged by the compiler
// tells if an expression can throw an exception, based on can_throw?(p or m)
// debug loop
/* The go function for: Compile/g_throw(self:any) [status=1] */
func F_Compile_g_throw_any (self *ClaireAny ) EID { 
    // eid body s = boolean
    var Result EID 
    { var v *ClaireBoolean  
      var try_1 EID 
      try_1 = F_Optimize_g_throw1_any(self)
      if ErrorIn(try_1) {Result = try_1
      } else {
      v = ToBoolean(OBJ(try_1))
      if (ToBoolean(C_DSHOW.Value) == CTRUE) { 
        Core.F_tformat_string(MakeString("-> g_throw(~S)=~S\n"),0,MakeConstantList(self,v.Id()))
        } 
      Result = EID{v.Id(),0}
      }
      } 
    return Result} 
  
// The EID go function for: Compile/g_throw @ any (throw: true) 
func E_Compile_g_throw_any (self EID) EID { 
    return F_Compile_g_throw_any(ANY(self) )} 
  
/* The go function for: g_throw1(self:any) [status=1] */
func F_Optimize_g_throw1_any (self *ClaireAny ) EID { 
    // eid body s = boolean
    var Result EID 
    if (self.Isa.IsIn(C_bag) == CTRUE) { 
      { var g0059 *ClaireBag   = ToBag(self)
        { var arg_1 *ClaireAny  
          var try_2 EID 
          { 
            var x *ClaireAny  
            _ = x
            try_2= EID{CFALSE.Id(),0}
            var x_support *ClaireList  
            var try_3 EID 
            try_3 = Core.F_enumerate_any(g0059.Id())
            if ErrorIn(try_3) {try_2 = try_3
            } else {
            x_support = ToList(OBJ(try_3))
            x_len := x_support.Length()
            for i_it := 0; i_it < x_len; i_it++ { 
              x = x_support.At(i_it)
              var loop_4 EID 
              _ = loop_4
              var g0083I *ClaireBoolean  
              var try_5 EID 
              try_5 = F_Compile_g_throw_any(x)
              if ErrorIn(try_5) {loop_4 = try_5
              } else {
              g0083I = ToBoolean(OBJ(try_5))
              if (g0083I == CTRUE) { 
                try_2 = EID{CTRUE.Id(),0}
                break
                } else {
                loop_4 = EID{CFALSE.Id(),0}
                } 
              }
              if ErrorIn(loop_4) {try_2 = loop_4
              break
              } else {
              }}
              } 
            } 
          if ErrorIn(try_2) {Result = try_2
          } else {
          arg_1 = ANY(try_2)
          Result = EID{F_boolean_I_any(arg_1).Id(),0}
          }
          } 
        } 
      }  else if (self.Isa.IsIn(Language.C_Construct) == CTRUE) { 
      { var g0060 *Language.Construct   = Language.To_Construct(self)
        { var arg_6 *ClaireAny  
          var try_7 EID 
          { 
            var x *ClaireAny  
            _ = x
            try_7= EID{CFALSE.Id(),0}
            var x_support *ClaireList  
            x_support = g0060.Args
            x_len := x_support.Length()
            for i_it := 0; i_it < x_len; i_it++ { 
              x = x_support.At(i_it)
              var loop_8 EID 
              _ = loop_8
              var g0084I *ClaireBoolean  
              var try_9 EID 
              try_9 = F_Compile_g_throw_any(x)
              if ErrorIn(try_9) {loop_8 = try_9
              } else {
              g0084I = ToBoolean(OBJ(try_9))
              if (g0084I == CTRUE) { 
                try_7 = EID{CTRUE.Id(),0}
                break
                } else {
                loop_8 = EID{CFALSE.Id(),0}
                } 
              }
              if ErrorIn(loop_8) {try_7 = loop_8
              break
              } else {
              }
              } 
            } 
          if ErrorIn(try_7) {Result = try_7
          } else {
          arg_6 = ANY(try_7)
          Result = EID{F_boolean_I_any(arg_6).Id(),0}
          }
          } 
        } 
      }  else if (self.Isa.IsIn(Language.C_Assign) == CTRUE) { 
      { var g0061 *Language.Assign   = Language.To_Assign(self)
        Result = F_Compile_g_throw_any(g0061.Arg)
        } 
      }  else if (self.Isa.IsIn(Language.C_Gassign) == CTRUE) { 
      { var g0062 *Language.Gassign   = Language.To_Gassign(self)
        Result = F_Compile_g_throw_any(g0062.Arg)
        } 
      }  else if (self.Isa.IsIn(Language.C_And) == CTRUE) { 
      { var g0063 *Language.And   = Language.To_And(self)
        Result = F_Compile_g_throw_any(g0063.Args.Id())
        } 
      }  else if (self.Isa.IsIn(Language.C_Or) == CTRUE) { 
      { var g0064 *Language.Or   = Language.To_Or(self)
        Result = F_Compile_g_throw_any(g0064.Args.Id())
        } 
      }  else if (self.Isa.IsIn(Language.C_Call) == CTRUE) { 
      { var g0065 *Language.Call   = Language.To_Call(self)
        { 
          var v_and4 *ClaireBoolean  
          
          v_and4 = Core.F__I_equal_any(g0065.Selector.Id(),Core.C_unsafe.Id())
          if (v_and4 == CFALSE) {Result = EID{CFALSE.Id(),0}
          } else { 
            var try_10 EID 
            { 
              var v_or6 *ClaireBoolean  
              
              var try_11 EID 
              try_11 = F_Compile_g_throw_any(g0065.Args.Id())
              if ErrorIn(try_11) {try_10 = try_11
              } else {
              v_or6 = ToBoolean(OBJ(try_11))
              if (v_or6 == CTRUE) {try_10 = EID{CTRUE.Id(),0}
              } else { 
                var try_12 EID 
                try_12 = F_Compile_can_throw_ask_property(g0065.Selector)
                if ErrorIn(try_12) {try_10 = try_12
                } else {
                v_or6 = ToBoolean(OBJ(try_12))
                if (v_or6 == CTRUE) {try_10 = EID{CTRUE.Id(),0}
                } else { 
                  try_10 = EID{CFALSE.Id(),0}} 
                } 
              }}
              } 
            if ErrorIn(try_10) {Result = try_10
            } else {
            v_and4 = ToBoolean(OBJ(try_10))
            if (v_and4 == CFALSE) {Result = EID{CFALSE.Id(),0}
            } else { 
              Result = EID{CTRUE.Id(),0}} 
            } 
          }
          } 
        } 
      }  else if (self.Isa.IsIn(Language.C_Call_method) == CTRUE) { 
      { var g0066 *Language.CallMethod   = Language.To_CallMethod(self)
        { 
          var v_and4 *ClaireBoolean  
          
          v_and4 = Core.F__I_equal_any(g0066.Arg.Id(),C_Compile_m_unsafe.Value)
          if (v_and4 == CFALSE) {Result = EID{CFALSE.Id(),0}
          } else { 
            var try_13 EID 
            try_13 = F_Optimize_notOpt_Call_method(g0066)
            if ErrorIn(try_13) {Result = try_13
            } else {
            v_and4 = ToBoolean(OBJ(try_13))
            if (v_and4 == CFALSE) {Result = EID{CFALSE.Id(),0}
            } else { 
              v_and4 = Core.F__I_equal_any(g0066.Arg.Selector.Id(),Core.C_externC.Id())
              if (v_and4 == CFALSE) {Result = EID{CFALSE.Id(),0}
              } else { 
                var try_14 EID 
                { 
                  var v_or8 *ClaireBoolean  
                  
                  var try_15 EID 
                  try_15 = F_Compile_g_throw_any(g0066.Args.Id())
                  if ErrorIn(try_15) {try_14 = try_15
                  } else {
                  v_or8 = ToBoolean(OBJ(try_15))
                  if (v_or8 == CTRUE) {try_14 = EID{CTRUE.Id(),0}
                  } else { 
                    var try_16 EID 
                    try_16 = F_Compile_can_throw_ask_method(g0066.Arg)
                    if ErrorIn(try_16) {try_14 = try_16
                    } else {
                    v_or8 = ToBoolean(OBJ(try_16))
                    if (v_or8 == CTRUE) {try_14 = EID{CTRUE.Id(),0}
                    } else { 
                      try_14 = EID{CFALSE.Id(),0}} 
                    } 
                  }}
                  } 
                if ErrorIn(try_14) {Result = try_14
                } else {
                v_and4 = ToBoolean(OBJ(try_14))
                if (v_and4 == CFALSE) {Result = EID{CFALSE.Id(),0}
                } else { 
                  Result = EID{CTRUE.Id(),0}} 
                } 
              } 
            } 
          }}
          } 
        } 
      }  else if (self.Isa.IsIn(Language.C_Call_slot) == CTRUE) { 
      { var g0067 *Language.CallSlot   = Language.To_CallSlot(self)
        { 
          var v_or4 *ClaireBoolean  
          
          var try_17 EID 
          try_17 = F_Compile_g_throw_any(g0067.Arg)
          if ErrorIn(try_17) {Result = try_17
          } else {
          v_or4 = ToBoolean(OBJ(try_17))
          if (v_or4 == CTRUE) {Result = EID{CTRUE.Id(),0}
          } else { 
            v_or4 = MakeBoolean((g0067.Test.Id() != CNULL) && (g0067.Test == CTRUE))
            if (v_or4 == CTRUE) {Result = EID{CTRUE.Id(),0}
            } else { 
              Result = EID{CFALSE.Id(),0}} 
            } 
          }
          } 
        } 
      }  else if (self.Isa.IsIn(Language.C_Call_table) == CTRUE) { 
      { var g0068 *Language.CallTable   = Language.To_CallTable(self)
        { 
          var v_or4 *ClaireBoolean  
          
          var try_18 EID 
          try_18 = F_Compile_g_throw_any(g0068.Arg)
          if ErrorIn(try_18) {Result = try_18
          } else {
          v_or4 = ToBoolean(OBJ(try_18))
          if (v_or4 == CTRUE) {Result = EID{CTRUE.Id(),0}
          } else { 
            v_or4 = MakeBoolean((g0068.Test.Id() != CNULL) && (g0068.Test == CTRUE))
            if (v_or4 == CTRUE) {Result = EID{CTRUE.Id(),0}
            } else { 
              Result = EID{CFALSE.Id(),0}} 
            } 
          }
          } 
        } 
      }  else if (self.Isa.IsIn(Language.C_Call_array) == CTRUE) { 
      { var g0069 *Language.CallArray   = Language.To_CallArray(self)
        { 
          var v_or4 *ClaireBoolean  
          
          var try_19 EID 
          try_19 = F_Compile_g_throw_any(g0069.Selector)
          if ErrorIn(try_19) {Result = try_19
          } else {
          v_or4 = ToBoolean(OBJ(try_19))
          if (v_or4 == CTRUE) {Result = EID{CTRUE.Id(),0}
          } else { 
            var try_20 EID 
            try_20 = F_Compile_g_throw_any(g0069.Arg)
            if ErrorIn(try_20) {Result = try_20
            } else {
            v_or4 = ToBoolean(OBJ(try_20))
            if (v_or4 == CTRUE) {Result = EID{CTRUE.Id(),0}
            } else { 
              Result = EID{CFALSE.Id(),0}} 
            } 
          }}
          } 
        } 
      }  else if (self.Isa.IsIn(Language.C_Super) == CTRUE) { 
      { var g0070 *Language.Super   = Language.To_Super(self)
        { 
          var v_or4 *ClaireBoolean  
          
          var try_21 EID 
          try_21 = F_Compile_g_throw_any(g0070.Args.Id())
          if ErrorIn(try_21) {Result = try_21
          } else {
          v_or4 = ToBoolean(OBJ(try_21))
          if (v_or4 == CTRUE) {Result = EID{CTRUE.Id(),0}
          } else { 
            var try_22 EID 
            try_22 = F_Compile_can_throw_ask_property(g0070.Selector)
            if ErrorIn(try_22) {Result = try_22
            } else {
            v_or4 = ToBoolean(OBJ(try_22))
            if (v_or4 == CTRUE) {Result = EID{CTRUE.Id(),0}
            } else { 
              Result = EID{CFALSE.Id(),0}} 
            } 
          }}
          } 
        } 
      }  else if (self.Isa.IsIn(Language.C_Update) == CTRUE) { 
      { var g0071 *Language.Update   = Language.To_Update(self)
        { 
          var v_or4 *ClaireBoolean  
          
          var try_23 EID 
          try_23 = F_Compile_g_throw_any(g0071.Value)
          if ErrorIn(try_23) {Result = try_23
          } else {
          v_or4 = ToBoolean(OBJ(try_23))
          if (v_or4 == CTRUE) {Result = EID{CTRUE.Id(),0}
          } else { 
            var try_24 EID 
            try_24 = F_Compile_g_throw_any(g0071.ClaireVar)
            if ErrorIn(try_24) {Result = try_24
            } else {
            v_or4 = ToBoolean(OBJ(try_24))
            if (v_or4 == CTRUE) {Result = EID{CTRUE.Id(),0}
            } else { 
              v_or4 = F_Compile_update_write_ask_Update(g0071)
              if (v_or4 == CTRUE) {Result = EID{CTRUE.Id(),0}
              } else { 
                Result = EID{CFALSE.Id(),0}} 
              } 
            } 
          }}
          } 
        } 
      }  else if (self.Isa.IsIn(Language.C_Cast) == CTRUE) { 
      { var g0072 *Language.Cast   = Language.To_Cast(self)
        Result = F_Compile_g_throw_any(g0072.Arg)
        } 
      }  else if (self.Isa.IsIn(C_Compile_C_cast) == CTRUE) { 
      { var g0073 *Compile_CCast   = To_Compile_CCast(self)
        Result = F_Compile_g_throw_any(g0073.Arg)
        } 
      }  else if (self.Isa.IsIn(Language.C_Let) == CTRUE) { 
      { var g0074 *Language.Let   = Language.To_Let(self)
        { 
          var v_or4 *ClaireBoolean  
          
          var try_25 EID 
          try_25 = F_Compile_g_throw_any(g0074.Value)
          if ErrorIn(try_25) {Result = try_25
          } else {
          v_or4 = ToBoolean(OBJ(try_25))
          if (v_or4 == CTRUE) {Result = EID{CTRUE.Id(),0}
          } else { 
            var try_26 EID 
            try_26 = F_Compile_g_throw_any(g0074.Arg)
            if ErrorIn(try_26) {Result = try_26
            } else {
            v_or4 = ToBoolean(OBJ(try_26))
            if (v_or4 == CTRUE) {Result = EID{CTRUE.Id(),0}
            } else { 
              Result = EID{CFALSE.Id(),0}} 
            } 
          }}
          } 
        } 
      }  else if (self.Isa.IsIn(Language.C_Do) == CTRUE) { 
      { var g0075 *Language.Do   = Language.To_Do(self)
        Result = F_Compile_g_throw_any(g0075.Args.Id())
        } 
      }  else if (self.Isa.IsIn(Language.C_While) == CTRUE) { 
      { var g0076 *Language.While   = Language.To_While(self)
        { 
          var v_or4 *ClaireBoolean  
          
          var try_27 EID 
          try_27 = F_Compile_g_throw_any(g0076.Test)
          if ErrorIn(try_27) {Result = try_27
          } else {
          v_or4 = ToBoolean(OBJ(try_27))
          if (v_or4 == CTRUE) {Result = EID{CTRUE.Id(),0}
          } else { 
            var try_28 EID 
            try_28 = F_Compile_g_throw_any(g0076.Arg)
            if ErrorIn(try_28) {Result = try_28
            } else {
            v_or4 = ToBoolean(OBJ(try_28))
            if (v_or4 == CTRUE) {Result = EID{CTRUE.Id(),0}
            } else { 
              Result = EID{CFALSE.Id(),0}} 
            } 
          }}
          } 
        } 
      }  else if (self.Isa.IsIn(Language.C_Construct) == CTRUE) { 
      { var g0077 *Language.Construct   = Language.To_Construct(self)
        Result = F_Compile_g_throw_any(g0077.Args.Id())
        } 
      }  else if (self.Isa.IsIn(Language.C_If) == CTRUE) { 
      { var g0078 *Language.If   = Language.To_If(self)
        { 
          var v_or4 *ClaireBoolean  
          
          var try_29 EID 
          try_29 = F_Compile_g_throw_any(g0078.Test)
          if ErrorIn(try_29) {Result = try_29
          } else {
          v_or4 = ToBoolean(OBJ(try_29))
          if (v_or4 == CTRUE) {Result = EID{CTRUE.Id(),0}
          } else { 
            var try_30 EID 
            try_30 = F_Compile_g_throw_any(g0078.Arg)
            if ErrorIn(try_30) {Result = try_30
            } else {
            v_or4 = ToBoolean(OBJ(try_30))
            if (v_or4 == CTRUE) {Result = EID{CTRUE.Id(),0}
            } else { 
              var try_31 EID 
              try_31 = F_Compile_g_throw_any(g0078.Other)
              if ErrorIn(try_31) {Result = try_31
              } else {
              v_or4 = ToBoolean(OBJ(try_31))
              if (v_or4 == CTRUE) {Result = EID{CTRUE.Id(),0}
              } else { 
                Result = EID{CFALSE.Id(),0}} 
              } 
            } 
          }}}
          } 
        } 
      }  else if (self.Isa.IsIn(Language.C_For) == CTRUE) { 
      { var g0079 *Language.For   = Language.To_For(self)
        { 
          var v_or4 *ClaireBoolean  
          
          var try_32 EID 
          try_32 = F_Compile_g_throw_any(g0079.SetArg)
          if ErrorIn(try_32) {Result = try_32
          } else {
          v_or4 = ToBoolean(OBJ(try_32))
          if (v_or4 == CTRUE) {Result = EID{CTRUE.Id(),0}
          } else { 
            var try_33 EID 
            try_33 = F_Compile_g_throw_any(g0079.Arg)
            if ErrorIn(try_33) {Result = try_33
            } else {
            v_or4 = ToBoolean(OBJ(try_33))
            if (v_or4 == CTRUE) {Result = EID{CTRUE.Id(),0}
            } else { 
              Result = EID{CFALSE.Id(),0}} 
            } 
          }}
          } 
        } 
      }  else if (self.Isa.IsIn(Language.C_Iteration) == CTRUE) { 
      { var g0080 *Language.Iteration   = Language.To_Iteration(self)
        { 
          var v_or4 *ClaireBoolean  
          
          var try_34 EID 
          try_34 = F_Compile_g_throw_any(g0080.SetArg)
          if ErrorIn(try_34) {Result = try_34
          } else {
          v_or4 = ToBoolean(OBJ(try_34))
          if (v_or4 == CTRUE) {Result = EID{CTRUE.Id(),0}
          } else { 
            var try_35 EID 
            try_35 = F_Compile_g_throw_any(g0080.Arg)
            if ErrorIn(try_35) {Result = try_35
            } else {
            v_or4 = ToBoolean(OBJ(try_35))
            if (v_or4 == CTRUE) {Result = EID{CTRUE.Id(),0}
            } else { 
              Result = EID{CFALSE.Id(),0}} 
            } 
          }}
          } 
        } 
      }  else if (self.Isa.IsIn(Language.C_Handle) == CTRUE) { 
      { var g0081 *Language.ClaireHandle   = Language.To_ClaireHandle(self)
        { 
          var v_or4 *ClaireBoolean  
          
          v_or4 = Core.F__I_equal_any(g0081.Test,C_any.Id())
          if (v_or4 == CTRUE) {Result = EID{CTRUE.Id(),0}
          } else { 
            var try_36 EID 
            try_36 = F_Compile_g_throw_any(g0081.Other)
            if ErrorIn(try_36) {Result = try_36
            } else {
            v_or4 = ToBoolean(OBJ(try_36))
            if (v_or4 == CTRUE) {Result = EID{CTRUE.Id(),0}
            } else { 
              Result = EID{CFALSE.Id(),0}} 
            } 
          }
          } 
        } 
      } else {
      Result = EID{CFALSE.Id(),0}
      } 
    return Result} 
  
// The EID go function for: g_throw1 @ any (throw: true) 
func E_Optimize_g_throw1_any (self EID) EID { 
    return F_Optimize_g_throw1_any(ANY(self) )} 
  
// return true in regular case, false if the optimization means that no error will occur.
//  this is ugly, so use sparingly for truly critical code optimization:  
//    - (x % y) can raise an error in the generic case (using F_belong) but not in the  optimized case
//    - class!(...) can raise an error in interpreted mode nut not at compile time
//    - division by non-zero integer constant is OK
//    - etc (extensible)   ... hopefully go will support exceptions one day so I can get rid of this junk :)
/* The go function for: notOpt(self:Call_method) [status=1] */
func F_Optimize_notOpt_Call_method (self *Language.CallMethod ) EID { 
    // eid body s = boolean
    var Result EID 
    if (self.Arg.Id() == C_Compile_m_member.Value) { 
      { var t2 *ClaireClass  
        var try_1 EID 
        try_1 = Language.F_static_type_any(self.Args.At(2-1))
        if ErrorIn(try_1) {Result = try_1
        } else {
        t2 = ToClass(OBJ(try_1))
        Result = EID{MakeBoolean((ToType(t2.Id()).Included(ToType(C_type.Id())) == CTRUE) || (ToType(t2.Id()).Included(ToType(C_list.Id())) == CTRUE) || (ToType(t2.Id()).Included(ToType(C_integer.Id())) == CTRUE) || (ToType(t2.Id()).Included(ToType(C_array.Id())) == CTRUE)).Not.Id(),0}
        }
        } 
      }  else if (self.Arg.Selector.Id() == C_class_I.Id()) { 
      Result = EID{self.Args.At(1-1).Isa.IsIn(C_symbol).Not.Id(),0}
      }  else if ((self.Arg.Selector.Id() == C__7.Id()) || 
        (self.Arg.Selector.Id() == C_mod.Id())) { 
      Result = EID{MakeBoolean((C_integer.Id() != self.Args.At(2-1).Isa.Id()) || self.Args.At(2-1).IsInt(0)).Id(),0}
      }  else if ((C_table.Id() == self.Args.At(1-1).Isa.Id()) && 
        (self.Arg.Selector.Id() == C_nth.Id())) { 
      { 
        var v_and3 *ClaireBoolean  
        
        var try_2 EID 
        { var arg_3 *ClaireBoolean  
          var try_4 EID 
          { var arg_5 *ClaireClass  
            var try_6 EID 
            try_6 = Language.F_static_type_any(self.Args.At(2-1))
            if ErrorIn(try_6) {try_4 = try_6
            } else {
            arg_5 = ToClass(OBJ(try_6))
            try_4 = EID{ToType(arg_5.Id()).Included(ToType(OBJ(Core.F_CALL(C_domain,ARGS(self.Args.At(1-1).ToEID()))))).Id(),0}
            }
            } 
          if ErrorIn(try_4) {try_2 = try_4
          } else {
          arg_3 = ToBoolean(OBJ(try_4))
          try_2 = EID{arg_3.Not.Id(),0}
          }
          } 
        if ErrorIn(try_2) {Result = try_2
        } else {
        v_and3 = ToBoolean(OBJ(try_2))
        if (v_and3 == CFALSE) {Result = EID{CFALSE.Id(),0}
        } else { 
          v_and3 = Core.F__inf_integer(C_compiler.Safety,2)
          if (v_and3 == CFALSE) {Result = EID{CFALSE.Id(),0}
          } else { 
            Result = EID{CTRUE.Id(),0}} 
          } 
        }
        } 
      } else {
      Result = EID{CTRUE.Id(),0}
      } 
    return Result} 
  
// The EID go function for: notOpt @ Call_method (throw: true) 
func E_Optimize_notOpt_Call_method (self EID) EID { 
    return F_Optimize_notOpt_Call_method(Language.To_CallMethod(OBJ(self)) )} 
  
//regular case !
// can_throw is based on restrictions analysis ... unless it is open => could always return an error
/* The go function for: Compile/can_throw?(p:property) [status=1] */
func F_Compile_can_throw_ask_property (p *ClaireProperty ) EID { 
    // eid body s = boolean
    var Result EID 
    { 
      var v_or2 *ClaireBoolean  
      
      v_or2 = Equal(MakeInteger(p.Open).Id(),MakeInteger(3).Id())
      if (v_or2 == CTRUE) {Result = EID{CTRUE.Id(),0}
      } else { 
        var try_1 EID 
        { 
          var v_and4 *ClaireBoolean  
          
          v_and4 = Core.F_not_any(ToList(C_Compile_NoErrorOptimize.Value).Memq(p.Id()).Id())
          if (v_and4 == CFALSE) {try_1 = EID{CFALSE.Id(),0}
          } else { 
            var try_2 EID 
            { var arg_3 *ClaireAny  
              var try_4 EID 
              { 
                var m *ClaireRestriction  
                _ = m
                var m_iter *ClaireAny  
                try_4= EID{CFALSE.Id(),0}
                for _,m_iter = range(p.Restrictions.ValuesO()){ 
                  m = ToRestriction(m_iter)
                  var loop_5 EID 
                  _ = loop_5
                  var g0087I *ClaireBoolean  
                  var try_6 EID 
                  if (C_method.Id() == m.Isa.Id()) { 
                    { var g0085 *ClaireMethod   = ToMethod(m.Id())
                      try_6 = F_Compile_can_throw_ask_method(g0085)
                      } 
                    } else {
                    try_6 = EID{CFALSE.Id(),0}
                    } 
                  if ErrorIn(try_6) {loop_5 = try_6
                  } else {
                  g0087I = ToBoolean(OBJ(try_6))
                  if (g0087I == CTRUE) { 
                    try_4 = EID{CTRUE.Id(),0}
                    break
                    } else {
                    loop_5 = EID{CFALSE.Id(),0}
                    } 
                  }
                  if ErrorIn(loop_5) {try_4 = loop_5
                  break
                  } else {
                  }
                  } 
                } 
              if ErrorIn(try_4) {try_2 = try_4
              } else {
              arg_3 = ANY(try_4)
              try_2 = EID{F_boolean_I_any(arg_3).Id(),0}
              }
              } 
            if ErrorIn(try_2) {try_1 = try_2
            } else {
            v_and4 = ToBoolean(OBJ(try_2))
            if (v_and4 == CFALSE) {try_1 = EID{CFALSE.Id(),0}
            } else { 
              try_1 = EID{CTRUE.Id(),0}} 
            } 
          }
          } 
        if ErrorIn(try_1) {Result = try_1
        } else {
        v_or2 = ToBoolean(OBJ(try_1))
        if (v_or2 == CTRUE) {Result = EID{CTRUE.Id(),0}
        } else { 
          Result = EID{CFALSE.Id(),0}} 
        } 
      }
      } 
    return Result} 
  
// The EID go function for: Compile/can_throw? @ property (throw: true) 
func E_Compile_can_throw_ask_property (p EID) EID { 
    return F_Compile_can_throw_ask_property(ToProperty(OBJ(p)) )} 
  
// access to status ... -1 means that it was never computed 
// Force*Throw is used to adjust for cross-compiling with a status change
// the reference to safety is here to ensure cross-compilation mode
/* The go function for: Compile/can_throw?(m:method) [status=1] */
func F_Compile_can_throw_ask_method (m *ClaireMethod ) EID { 
    // eid body s = boolean
    var Result EID 
    if (((C_compiler.Safety >= 2) && 
          ((ToList(C_Compile_NoErrorOptimize.Value).Memq(m.Id()) == CTRUE) || 
              (ToList(C_Compile_NoErrorOptimize.Value).Memq(m.Selector.Id()) == CTRUE))) || 
        (ToType(C_method.Id()).EmptyList().Memq(m.Id()) == CTRUE)) { 
      Result = EID{CFALSE.Id(),0}
      }  else if (ToType(C_method.Id()).EmptyList().Memq(m.Id()) == CTRUE) { 
      Result = EID{CTRUE.Id(),0}
      }  else if ((m.Status != -1) || 
        (m.Formula.Id() == CNULL)) { 
      Result = EID{Core.F__I_equal_any(MakeInteger(m.Status).Id(),MakeInteger(0).Id()).Id(),0}
      } else {
      Result = F_Compile_can_throw_I_method(m)
      } 
    return Result} 
  
// The EID go function for: Compile/can_throw? @ method (throw: true) 
func E_Compile_can_throw_ask_method (m EID) EID { 
    return F_Compile_can_throw_ask_method(ToMethod(OBJ(m)) )} 
  
// debug handle
// here we recursively call g_throw on the body => forced re-compute of status(m) (status!(m) in CLAIRE3)
/* The go function for: Compile/can_throw!(m:method) [status=1] */
func F_Compile_can_throw_I_method (m *ClaireMethod ) EID { 
    // eid body s = boolean
    var Result EID 
    m.Status = 0
    var v_gassign1 *ClaireAny  
    var try_2 EID 
    try_2 = Core.F_CALL(C_c_code,ARGS(m.Formula.Body.ToEID(),EID{m.Range.Class_I().Id(),0}))
    if ErrorIn(try_2) {Result = try_2
    } else {
    v_gassign1 = ANY(try_2)
    C_DTHROW.Value = v_gassign1
    Result = v_gassign1.ToEID()
    } 
  if !ErrorIn(Result) {
  var g0088I *ClaireBoolean  
  var try_3 EID 
  { var arg_4 *ClaireAny  
    var try_5 EID 
    try_5 = Core.F_CALL(C_c_code,ARGS(m.Formula.Body.ToEID(),EID{m.Range.Class_I().Id(),0}))
    if ErrorIn(try_5) {try_3 = try_5
    } else {
    arg_4 = ANY(try_5)
    try_3 = F_Compile_g_throw_any(arg_4)
    }
    } 
  if ErrorIn(try_3) {Result = try_3
  } else {
  g0088I = ToBoolean(OBJ(try_3))
  if (g0088I == CTRUE) { 
    m.Status = 1
    Result = EID{CTRUE.Id(),0}
    } else {
    Result = EID{CFALSE.Id(),0}
    } 
  }
  }
  return Result} 

// The EID go function for: Compile/can_throw! @ method (throw: true) 
func E_Compile_can_throw_I_method (m EID) EID { 
  return F_Compile_can_throw_I_method(ToMethod(OBJ(m)) )} 

// read can_throw from the status, not influenced by exceptions (for code generation)
// however, for a new method, compute the status
/* The go function for: Compile/can_throw_status(m:method) [status=1] */
func F_Compile_can_throw_status_method (m *ClaireMethod ) EID { 
  // eid body s = boolean
  var Result EID 
  if (m.Status == -1) { 
    Result = F_Compile_can_throw_I_method(m)
    } else {
    Result = EID{CFALSE.Id(),0}
    } 
  if !ErrorIn(Result) {
  Result = EID{Core.F__I_equal_any(MakeInteger(m.Status).Id(),MakeInteger(0).Id()).Id(),0}
  }
  return Result} 

// The EID go function for: Compile/can_throw_status @ method (throw: true) 
func E_Compile_can_throw_status_method (m EID) EID { 
  return F_Compile_can_throw_status_method(ToMethod(OBJ(m)) )} 

// useful #2: provoke a recomputation of status
/* The go function for: s_throw(m:method) [status=1] */
func F_s_throw_method (m *ClaireMethod ) EID { 
  // eid body s = void
  var Result EID 
  { var la *ClaireLambda   = m.Formula
    { var news *ClaireBoolean  
      var try_1 EID 
      try_1 = F_Compile_g_throw_any(la.Body)
      if ErrorIn(try_1) {Result = try_1
      } else {
      news = ToBoolean(OBJ(try_1))
      Core.F_tformat_string(MakeString("status(~S) := ~S \n"),0,MakeConstantList(m.Id(),news.Id()))
      { 
        var va_arg1 *ClaireMethod  
        var va_arg2 int 
        va_arg1 = m
        if (news == CTRUE) { 
          va_arg2 = 1
          } else {
          va_arg2 = 0
          } 
        va_arg1.Status = va_arg2
        Result = EID{C__INT,IVAL(va_arg2)}
        } 
      }
      } 
    } 
  return Result} 

// The EID go function for: s_throw @ method (throw: true) 
func E_s_throw_method (m EID) EID { 
  return F_s_throw_method(ToMethod(OBJ(m)) )} 

// ******************************************************************
// *    Part 4: Names & identifiers management                      *
// ******************************************************************
// check that the module is allowed and otherwise complain because of x;
// this should raise an error, it simply returns false if there is a problem
/* The go function for: legal?(self:module,x:any) [status=0] */
func F_Optimize_legal_ask_module (self *ClaireModule ,x *ClaireAny ) *ClaireBoolean  { 
  // procedure body with s = boolean
  var Result *ClaireBoolean  
  if ((x == C_Compile_object_I.Id()) || 
      (x == C_Compile_anyObject_I.Id())) { 
    Result = CTRUE
    }  else if (F_boolean_I_any(C_OPT.LegalModules.Id()) == CTRUE) { 
    var g0091I *ClaireBoolean  
    { 
      var v_and2 *ClaireBoolean  
      
      v_and2 = C_OPT.LegalModules.Contain_ask(self.Id()).Not
      if (v_and2 == CFALSE) {g0091I = CFALSE
      } else { 
        if (C_method.Id() == x.Isa.Id()) { 
          { var g0090 *ClaireMethod   = ToMethod(x)
            v_and2 = MakeBoolean((g0090.Selector.Id() != C_add_method.Id()) && ((g0090.Inline_ask.Id() == CFALSE.Id()) || 
                (C_compiler.Inline_ask != CTRUE)))
            } 
          } else {
          v_and2 = CFALSE
          } 
        if (v_and2 == CFALSE) {g0091I = CFALSE
        } else { 
          g0091I = CTRUE} 
        } 
      } 
    if (g0091I == CTRUE) { 
      Core.F_tformat_string(MakeString("legal_modules = ~S\n"),0,MakeConstantList(C_OPT.LegalModules.Id()))
      Core.F_tformat_string(MakeString("---- ERROR: ~S implies using ~S !\n\n"),0,MakeConstantList(x,self.Id()))
      Result = CFALSE
      } else {
      Result = CTRUE
      } 
    } else {
    C_OPT.NeedModules.AddFast(self.Id())
    Result = CTRUE
    } 
  return Result} 

// The EID go function for: legal? @ module (throw: false) 
func E_Optimize_legal_ask_module (self EID,x EID) EID { 
  return EID{F_Optimize_legal_ask_module(ToModule(OBJ(self)),ANY(x) ).Id(),0}} 

/* The go function for: legal?(self:environment,x:any) [status=0] */
func F_Optimize_legal_ask_environment (self *ClaireEnvironment ,x *ClaireAny ) *ClaireAny  { 
  return  CTRUE.Id()
  } 

// The EID go function for: legal? @ environment (throw: false) 
func E_Optimize_legal_ask_environment (self EID,x EID) EID { 
  return F_Optimize_legal_ask_environment(ToEnvironment(OBJ(self)),ANY(x) ).ToEID()} 

// A named object is used, thus it must be declared if it belongs to the
// current module - returns true if OK
/* The go function for: c_register(self:(thing U class)) [status=0] */
func F_Optimize_c_register_object (self *ClaireObject ) *ClaireAny  { 
  // procedure body with s = any
  var Result *ClaireAny  
  { var x *ClaireAny   = F_Compile_get_module_object(self)
    if (x != ClEnv.Id()) { 
      Result = ANY(Core.F_CALL(C_Optimize_legal_ask,ARGS(x.ToEID(),EID{self.Id(),0})))
      } else {
      Result = CTRUE.Id()
      } 
    } 
  return Result} 

// The EID go function for: c_register @ object (throw: false) 
func E_Optimize_c_register_object (self EID) EID { 
  return F_Optimize_c_register_object(ToObject(OBJ(self)) ).ToEID()} 

// looks if a property may be implicit and then add it in the right list
/* The go function for: c_register(self:property) [status=0] */
func F_Optimize_c_register_property (self *ClaireProperty ) *ClaireAny  { 
  // procedure body with s = any
  var Result *ClaireAny  
  { var m *ClaireModule   = ClEnv.Module_I
    { var m2 *ClaireAny   = F_Compile_get_module_object(ToObject(self.Id()))
      if (((m2 == C_claire.Id()) || 
            (m2 == m.Id())) && 
          (C_OPT.Objects.Memq(self.Id()) != CTRUE)) { 
        C_OPT.Properties.AddFast(self.Id())
        } 
      Result = F_Optimize_c_register_object(ToObject(self.Id()))
      } 
    } 
  return Result} 

// The EID go function for: c_register @ property (throw: false) 
func E_Optimize_c_register_property (self EID) EID { 
  return F_Optimize_c_register_property(ToProperty(OBJ(self)) ).ToEID()} 

// declare the property as used and check if a property may allocate
/* The go function for: selector_register(self:property) [status=0] */
func F_Optimize_selector_register_property (self *ClaireProperty ) *ClaireAny  { 
  F_Optimize_c_register_property(self)
  return  self.Id()
  } 

// The EID go function for: selector_register @ property (throw: false) 
func E_Optimize_selector_register_property (self EID) EID { 
  return F_Optimize_selector_register_property(ToProperty(OBJ(self)) ).ToEID()} 

// this method looks if the open slot is less than 1 or can be set to 1
// v3.3.48 note - weaken the open semantic to get a better c_status
/* The go function for: stable?(self:relation) [status=0] */
func F_Optimize_stable_ask_relation (self *ClaireRelation ) *ClaireBoolean  { 
  { var m *ClaireAny   = F_Compile_get_module_object(ToObject(self.Id()))
    _ = m
    if (self.Open == 2) { 
      self.Open = 1
      } 
    } 
  if ((self.Open <= 1) || 
      (self.Open == 4)) {return CTRUE
  } else {return CFALSE}} 

// The EID go function for: stable? @ relation (throw: false) 
func E_Optimize_stable_ask_relation (self EID) EID { 
  return EID{F_Optimize_stable_ask_relation(ToRelation(OBJ(self)) ).Id(),0}} 

// v3.2.04
// returns the module (i.e. the compilation unit, not the namespace) in which self is
// defined
/* The go function for: Compile/get_module(self:(thing U class)) [status=0] */
func F_Compile_get_module_object (self *ClaireObject ) *ClaireAny  { 
  return  ToSymbol(OBJ(Core.F_CALL(C_name,ARGS(EID{self.Id(),0})))).Defined().Id()
  } 

// The EID go function for: Compile/get_module @ object (throw: false) 
func E_Compile_get_module_object (self EID) EID { 
  return F_Compile_get_module_object(ToObject(OBJ(self)) ).ToEID()} 

//      (while (m.loaded = 0) m := m.part_of, m) ]
// allows to optimize the access
/* The go function for: known!(l:listargs) [status=0] */
func F_known_I_listargs (l *ClaireList ) *ClaireAny  { 
  // procedure body with s = any
  var Result *ClaireAny  
  C_OPT.ToRemove.AddFast(Reader.C_known_I.Id())
  { 
    var r *ClaireAny  
    _ = r
    Result= CFALSE.Id()
    var r_support *ClaireList  
    r_support = ToList(l.Id())
    r_len := r_support.Length()
    for i_it := 0; i_it < r_len; i_it++ { 
      r = r_support.At(i_it)
      if (r.Isa.IsIn(C_property) == CTRUE) { 
        { var g0095 *ClaireProperty   = ToProperty(r)
          C_OPT.Knowns.AddFast(g0095.Id())
          } 
        } 
      } 
    } 
  return Result} 

// The EID go function for: known! @ listargs (throw: false) 
func E_known_I_listargs (l EID) EID { 
  return F_known_I_listargs(ToList(OBJ(l)) ).ToEID()} 
