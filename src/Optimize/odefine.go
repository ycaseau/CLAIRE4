/***** CLAIRE Compilation of file /Users/ycaseau/Dropbox/src/clairev4.12/src/compile/odefine.cl 
         [version 4.1.4 / safety 5] Friday 01-03-2025 16:21:04 *****/

package Optimize
import (_ "fmt"
	. "Kernel"
	"Core"
	"Language"
	"Reader"
)

//-------- dumb function to prevent import errors --------
func import_g0290() { 
_ = Core.It
_ = Language.It
_ = Reader.It
} 


//+-------------------------------------------------------------+
//| CLAIRE                                                      |
//| odefine.cl                                                  |
//| Copyright (C) 1994 - 2025 Yves Caseau. All Rights Reserved  |
//| cf. copyright info in file object.cl: about()               |
//+-------------------------------------------------------------+
// *********************************************************************
// *  Table of contents                                                *
// *     Part 1: Set, List and Tuple creation                          *
// *     Part 2: Object definition                                     *
// *     Part 3: Method instantiation                                  *
// *     Part 4: Inverse Management                                    *
// *********************************************************************
// */
// *********************************************************************
// *     Part 1: Set, List and Tuple creation                          *
// *********************************************************************
// type inference has changed in v3.2:
/* The go function for: c_type(self:List) [status=1] */
func F_c_type_List (self *Language.List) EID { 
var Result EID

if (self.Of.Id() != CNULL) { 
  Result = EID{Core.F_param_I_class(C_list,self.Of).Id(),0}
  } else {
  { var _Zres *ClaireAny = CEMPTY.Id()
    { 
      var _Zx *ClaireAny
      _ = _Zx
      Result= EID{CFALSE.Id(),0}
      var _Zx_support *ClaireList
      _Zx_support = self.Args
      _Zx_len := _Zx_support.Length()
      for i_it := 0; i_it < _Zx_len; i_it++ { 
        _Zx = _Zx_support.At(i_it)
        var loop_1 EID
        _ = loop_1
        if (F_boolean_I_any(_Zres) == CTRUE) { 
          var try_2 EID
          { var arg_3 *ClaireClass
            var try_4 EID
            { var arg_5 *ClaireType
              var try_6 EID
              { var arg_7 *ClaireType
                var try_8 EID
                try_8 = Core.F_CALL(C_c_type,ARGS(_Zx.ToEID()))
                if ErrorIn(try_8) {try_6 = try_8
                } else {
                arg_7 = ToType(OBJ(try_8))
                try_6 = EID{F_Optimize_ptype_type(arg_7).Id(),0}
                }
                } 
              if ErrorIn(try_6) {try_4 = try_6
              } else {
              arg_5 = ToType(OBJ(try_6))
              try_4 = EID{arg_5.Class_I().Id(),0}
              }
              } 
            if ErrorIn(try_4) {try_2 = try_4
            } else {
            arg_3 = ToClass(OBJ(try_4))
            try_2 = EID{Core.F_meet_class(ToClass(_Zres),arg_3).Id(),0}
            }
            } 
          if ErrorIn(try_2) {loop_1 = try_2
          } else {
          _Zres = ANY(try_2)
          loop_1 = _Zres.ToEID()
          }
          } else {
          var try_9 EID
          { var arg_10 *ClaireType
            var try_11 EID
            { var arg_12 *ClaireType
              var try_13 EID
              try_13 = Core.F_CALL(C_c_type,ARGS(_Zx.ToEID()))
              if ErrorIn(try_13) {try_11 = try_13
              } else {
              arg_12 = ToType(OBJ(try_13))
              try_11 = EID{F_Optimize_ptype_type(arg_12).Id(),0}
              }
              } 
            if ErrorIn(try_11) {try_9 = try_11
            } else {
            arg_10 = ToType(OBJ(try_11))
            try_9 = EID{arg_10.Class_I().Id(),0}
            }
            } 
          if ErrorIn(try_9) {loop_1 = try_9
          } else {
          _Zres = ANY(try_9)
          loop_1 = _Zres.ToEID()
          }
          } 
        if ErrorIn(loop_1) {Result = loop_1
        break
        } else {
        }
        } 
      } 
    if !ErrorIn(Result) {
    Result = EID{Core.F_nth_class1(C_list,ToType(_Zres)).Id(),0}
    }
    } 
  } 
return Result} 

// The EID go function for: c_type @ List (throw: true) 
func E_c_type_List (self EID) EID { 
return F_c_type_List(Language.To_List(OBJ(self)) )} 

// compile a List: take the of parameter into account !
/* The go function for: c_code(self:List) [status=1] */
func F_c_code_List (self *Language.List) EID { 
var Result EID
{ var x *Language.List
  var try_1 EID
  { var _CL_obj *Language.List = Language.To_List(new(Language.List).Is(Language.C_List))
    { 
      var va_arg1 *Language.Construct
      var va_arg2 *ClaireList
      va_arg1 = Language.To_Construct(_CL_obj.Id())
      var try_2 EID
      { 
        var v_list3 *ClaireList
        var _Zx *ClaireAny
        var v_local3 *ClaireAny
        v_list3 = self.Args
        try_2 = EID{CreateList(ToType(CEMPTY.Id()),v_list3.Length()).Id(),0}
        for CLcount := 0; CLcount < v_list3.Length(); CLcount++{ 
          _Zx = v_list3.At(CLcount)
          var try_3 EID
          try_3 = Core.F_CALL(C_c_code,ARGS(_Zx.ToEID(),EID{C_any.Id(),0}))
          if ErrorIn(try_3) {try_2 = try_3
          break
          } else {
          v_local3 = ANY(try_3)
          ToList(OBJ(try_2)).PutAt(CLcount,v_local3)
          } 
        }
        } 
      if ErrorIn(try_2) {try_1 = try_2
      } else {
      va_arg2 = ToList(OBJ(try_2))
      va_arg1.Args = va_arg2
      try_1 = EID{va_arg2.Id(),0}
      }
      } 
    if !ErrorIn(try_1) {
    try_1 = EID{_CL_obj.Id(),0}
    }
    } 
  if ErrorIn(try_1) {Result = try_1
  } else {
  x = Language.To_List(OBJ(try_1))
  if (self.Of.Id() != CNULL) { 
    var g0291I *ClaireBoolean
    var try_4 EID
    { 
      var v_or2 *ClaireBoolean
      
      v_or2 = Core.F__sup_integer(C_compiler.Safety,4)
      if (v_or2 == CTRUE) {try_4 = EID{CTRUE.Id(),0}
      } else { 
        v_or2 = Equal(self.Of.Id(),CEMPTY.Id())
        if (v_or2 == CTRUE) {try_4 = EID{CTRUE.Id(),0}
        } else { 
          var try_5 EID
          { var arg_6 *ClaireAny
            var try_7 EID
            { 
              var _Zx *ClaireAny
              _ = _Zx
              try_7= EID{CFALSE.Id(),0}
              var _Zx_support *ClaireList
              _Zx_support = self.Args
              _Zx_len := _Zx_support.Length()
              for i_it := 0; i_it < _Zx_len; i_it++ { 
                _Zx = _Zx_support.At(i_it)
                var loop_8 EID
                _ = loop_8
                var g0292I *ClaireBoolean
                var try_9 EID
                { var arg_10 *ClaireBoolean
                  var try_11 EID
                  { var arg_12 *ClaireType
                    var try_13 EID
                    try_13 = Core.F_CALL(C_c_type,ARGS(_Zx.ToEID()))
                    if ErrorIn(try_13) {try_11 = try_13
                    } else {
                    arg_12 = ToType(OBJ(try_13))
                    try_11 = EID{arg_12.Included(self.Of).Id(),0}
                    }
                    } 
                  if ErrorIn(try_11) {try_9 = try_11
                  } else {
                  arg_10 = ToBoolean(OBJ(try_11))
                  try_9 = EID{arg_10.Not.Id(),0}
                  }
                  } 
                if ErrorIn(try_9) {loop_8 = try_9
                } else {
                g0292I = ToBoolean(OBJ(try_9))
                if (g0292I == CTRUE) { 
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
            if ErrorIn(try_7) {try_5 = try_7
            } else {
            arg_6 = ANY(try_7)
            try_5 = EID{Core.F_not_any(arg_6).Id(),0}
            }
            } 
          if ErrorIn(try_5) {try_4 = try_5
          } else {
          v_or2 = ToBoolean(OBJ(try_5))
          if (v_or2 == CTRUE) {try_4 = EID{CTRUE.Id(),0}
          } else { 
            try_4 = EID{CFALSE.Id(),0}} 
          } 
        } 
      }
      } 
    if ErrorIn(try_4) {Result = try_4
    } else {
    g0291I = ToBoolean(OBJ(try_4))
    if (g0291I == CTRUE) { 
      x.Of = self.Of
      Result = EID{x.Id(),0}
      } else {
      F_Compile_warn_void()
      { var arg_14 *ClaireList
        var try_15 EID
        { 
          var v_bag_arg *ClaireAny
          try_15= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
          var try_16 EID
          { 
            var v_list5 *ClaireList
            var _Zx *ClaireAny
            var v_local5 *ClaireAny
            v_list5 = self.Args
            try_16 = EID{CreateList(ToType(CEMPTY.Id()),v_list5.Length()).Id(),0}
            for CLcount := 0; CLcount < v_list5.Length(); CLcount++{ 
              _Zx = v_list5.At(CLcount)
              var try_17 EID
              try_17 = Core.F_CALL(C_c_type,ARGS(_Zx.ToEID()))
              if ErrorIn(try_17) {try_16 = try_17
              break
              } else {
              v_local5 = ANY(try_17)
              ToList(OBJ(try_16)).PutAt(CLcount,v_local5)
              } 
            }
            } 
          if ErrorIn(try_16) {try_15 = try_16
          } else {
          v_bag_arg = ANY(try_16)
          ToList(OBJ(try_15)).AddFast(v_bag_arg)
          ToList(OBJ(try_15)).AddFast(self.Of.Id())}
          } 
        if ErrorIn(try_15) {Result = try_15
        } else {
        arg_14 = ToList(OBJ(try_15))
        Result = Core.F_tformat_string(MakeString("unsafe typed list: ~S not in ~S [262]\n"),2,arg_14)
        }
        } 
      if !ErrorIn(Result) {
      { var arg_18 *Language.Call
        { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
          _CL_obj.Selector = Core.C_check_in
          _CL_obj.Args = MakeConstantList(x.Id(),C_list.Id(),self.Of.Id())
          arg_18 = _CL_obj
          } 
        Result = Core.F_CALL(C_c_code,ARGS(EID{arg_18.Id(),0},EID{C_list.Id(),0}))
        } 
      }
      } 
    }
    } else {
    Result = EID{x.Id(),0}
    } 
  }
  } 
return Result} 

// The EID go function for: c_code @ List (throw: true) 
func E_c_code_List (self EID) EID { 
return F_c_code_List(Language.To_List(OBJ(self)) )} 

// new in v3.2: static list have type inference !         
/* The go function for: c_type(self:Set) [status=1] */
func F_c_type_Set (self *Language.Set) EID { 
var Result EID

if (self.Of.Id() != CNULL) { 
  Result = EID{Core.F_param_I_class(C_set,self.Of).Id(),0}
  } else {
  { var _Zres *ClaireAny = CEMPTY.Id()
    { 
      var _Zx *ClaireAny
      _ = _Zx
      Result= EID{CFALSE.Id(),0}
      var _Zx_support *ClaireList
      _Zx_support = self.Args
      _Zx_len := _Zx_support.Length()
      for i_it := 0; i_it < _Zx_len; i_it++ { 
        _Zx = _Zx_support.At(i_it)
        var loop_1 EID
        _ = loop_1
        if (F_boolean_I_any(_Zres) == CTRUE) { 
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
            try_2 = EID{Core.F_meet_class(ToClass(_Zres),arg_3).Id(),0}
            }
            } 
          if ErrorIn(try_2) {loop_1 = try_2
          } else {
          _Zres = ANY(try_2)
          loop_1 = _Zres.ToEID()
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
          _Zres = ANY(try_7)
          loop_1 = _Zres.ToEID()
          }
          } 
        if ErrorIn(loop_1) {Result = loop_1
        break
        } else {
        }
        } 
      } 
    if !ErrorIn(Result) {
    Result = EID{Core.F_nth_class1(C_set,ToType(_Zres)).Id(),0}
    }
    } 
  } 
return Result} 

// The EID go function for: c_type @ Set (throw: true) 
func E_c_type_Set (self EID) EID { 
return F_c_type_Set(Language.To_Set(OBJ(self)) )} 

/* The go function for: c_code(self:Set) [status=1] */
func F_c_code_Set (self *Language.Set) EID { 
var Result EID
{ var x *Language.Set
  var try_1 EID
  { var _CL_obj *Language.Set = Language.To_Set(new(Language.Set).Is(Language.C_Set))
    { 
      var va_arg1 *Language.Construct
      var va_arg2 *ClaireList
      va_arg1 = Language.To_Construct(_CL_obj.Id())
      var try_2 EID
      { 
        var v_list3 *ClaireList
        var _Zx *ClaireAny
        var v_local3 *ClaireAny
        v_list3 = self.Args
        try_2 = EID{CreateList(ToType(CEMPTY.Id()),v_list3.Length()).Id(),0}
        for CLcount := 0; CLcount < v_list3.Length(); CLcount++{ 
          _Zx = v_list3.At(CLcount)
          var try_3 EID
          try_3 = Core.F_CALL(C_c_code,ARGS(_Zx.ToEID(),EID{C_any.Id(),0}))
          if ErrorIn(try_3) {try_2 = try_3
          break
          } else {
          v_local3 = ANY(try_3)
          ToList(OBJ(try_2)).PutAt(CLcount,v_local3)
          } 
        }
        } 
      if ErrorIn(try_2) {try_1 = try_2
      } else {
      va_arg2 = ToList(OBJ(try_2))
      va_arg1.Args = va_arg2
      try_1 = EID{va_arg2.Id(),0}
      }
      } 
    if !ErrorIn(try_1) {
    try_1 = EID{_CL_obj.Id(),0}
    }
    } 
  if ErrorIn(try_1) {Result = try_1
  } else {
  x = Language.To_Set(OBJ(try_1))
  if (self.Of.Id() != CNULL) { 
    var g0293I *ClaireBoolean
    var try_4 EID
    { 
      var v_or2 *ClaireBoolean
      
      v_or2 = Core.F__sup_integer(C_compiler.Safety,4)
      if (v_or2 == CTRUE) {try_4 = EID{CTRUE.Id(),0}
      } else { 
        v_or2 = Equal(self.Of.Id(),CEMPTY.Id())
        if (v_or2 == CTRUE) {try_4 = EID{CTRUE.Id(),0}
        } else { 
          var try_5 EID
          { var arg_6 *ClaireAny
            var try_7 EID
            { 
              var _Zx *ClaireAny
              _ = _Zx
              try_7= EID{CFALSE.Id(),0}
              var _Zx_support *ClaireList
              _Zx_support = self.Args
              _Zx_len := _Zx_support.Length()
              for i_it := 0; i_it < _Zx_len; i_it++ { 
                _Zx = _Zx_support.At(i_it)
                var loop_8 EID
                _ = loop_8
                var g0294I *ClaireBoolean
                var try_9 EID
                { var arg_10 *ClaireBoolean
                  var try_11 EID
                  { var arg_12 *ClaireType
                    var try_13 EID
                    try_13 = Core.F_CALL(C_c_type,ARGS(_Zx.ToEID()))
                    if ErrorIn(try_13) {try_11 = try_13
                    } else {
                    arg_12 = ToType(OBJ(try_13))
                    try_11 = EID{arg_12.Included(self.Of).Id(),0}
                    }
                    } 
                  if ErrorIn(try_11) {try_9 = try_11
                  } else {
                  arg_10 = ToBoolean(OBJ(try_11))
                  try_9 = EID{arg_10.Not.Id(),0}
                  }
                  } 
                if ErrorIn(try_9) {loop_8 = try_9
                } else {
                g0294I = ToBoolean(OBJ(try_9))
                if (g0294I == CTRUE) { 
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
            if ErrorIn(try_7) {try_5 = try_7
            } else {
            arg_6 = ANY(try_7)
            try_5 = EID{Core.F_not_any(arg_6).Id(),0}
            }
            } 
          if ErrorIn(try_5) {try_4 = try_5
          } else {
          v_or2 = ToBoolean(OBJ(try_5))
          if (v_or2 == CTRUE) {try_4 = EID{CTRUE.Id(),0}
          } else { 
            try_4 = EID{CFALSE.Id(),0}} 
          } 
        } 
      }
      } 
    if ErrorIn(try_4) {Result = try_4
    } else {
    g0293I = ToBoolean(OBJ(try_4))
    if (g0293I == CTRUE) { 
      x.Of = self.Of
      Result = EID{x.Id(),0}
      } else {
      F_Compile_warn_void()
      { var arg_14 *ClaireList
        var try_15 EID
        { 
          var v_bag_arg *ClaireAny
          try_15= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
          var try_16 EID
          { 
            var v_list5 *ClaireList
            var _Zx *ClaireAny
            var v_local5 *ClaireAny
            v_list5 = self.Args
            try_16 = EID{CreateList(ToType(CEMPTY.Id()),v_list5.Length()).Id(),0}
            for CLcount := 0; CLcount < v_list5.Length(); CLcount++{ 
              _Zx = v_list5.At(CLcount)
              var try_17 EID
              try_17 = Core.F_CALL(C_c_type,ARGS(_Zx.ToEID()))
              if ErrorIn(try_17) {try_16 = try_17
              break
              } else {
              v_local5 = ANY(try_17)
              ToList(OBJ(try_16)).PutAt(CLcount,v_local5)
              } 
            }
            } 
          if ErrorIn(try_16) {try_15 = try_16
          } else {
          v_bag_arg = ANY(try_16)
          ToList(OBJ(try_15)).AddFast(v_bag_arg)
          ToList(OBJ(try_15)).AddFast(self.Of.Id())}
          } 
        if ErrorIn(try_15) {Result = try_15
        } else {
        arg_14 = ToList(OBJ(try_15))
        Result = Core.F_tformat_string(MakeString("unsafe typed set: ~S not in ~S [262]\n"),2,arg_14)
        }
        } 
      if !ErrorIn(Result) {
      { var arg_18 *Language.Call
        { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
          _CL_obj.Selector = Core.C_check_in
          _CL_obj.Args = MakeConstantList(x.Id(),C_set.Id(),self.Of.Id())
          arg_18 = _CL_obj
          } 
        Result = Core.F_CALL(C_c_code,ARGS(EID{arg_18.Id(),0},EID{C_set.Id(),0}))
        } 
      }
      } 
    }
    } else {
    Result = EID{x.Id(),0}
    } 
  }
  } 
return Result} 

// The EID go function for: c_code @ Set (throw: true) 
func E_c_code_Set (self EID) EID { 
return F_c_code_Set(Language.To_Set(OBJ(self)) )} 

/* The go function for: c_type(self:Tuple) [status=1] */
func F_c_type_Tuple (self *Language.Tuple) EID { 
var Result EID
{ var arg_1 *ClaireList
  var try_2 EID
  { 
    var v_list1 *ClaireList
    var x *ClaireAny
    var v_local1 *ClaireAny
    v_list1 = self.Args
    try_2 = EID{CreateList(ToType(CEMPTY.Id()),v_list1.Length()).Id(),0}
    for CLcount := 0; CLcount < v_list1.Length(); CLcount++{ 
      x = v_list1.At(CLcount)
      var try_3 EID
      try_3 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
      if ErrorIn(try_3) {try_2 = try_3
      break
      } else {
      v_local1 = ANY(try_3)
      ToList(OBJ(try_2)).PutAt(CLcount,v_local1)
      } 
    }
    } 
  if ErrorIn(try_2) {Result = try_2
  } else {
  arg_1 = ToList(OBJ(try_2))
  Result = EID{arg_1.Tuple_I().Id(),0}
  }
  } 
return Result} 

// The EID go function for: c_type @ Tuple (throw: true) 
func E_c_type_Tuple (self EID) EID { 
return F_c_type_Tuple(Language.To_Tuple(OBJ(self)) )} 

/* The go function for: c_code(self:Tuple) [status=1] */
func F_c_code_Tuple (self *Language.Tuple) EID { 
var Result EID
{ var _CL_obj *Language.Tuple = Language.To_Tuple(new(Language.Tuple).Is(Language.C_Tuple))
  { 
    var va_arg1 *Language.Construct
    var va_arg2 *ClaireList
    va_arg1 = Language.To_Construct(_CL_obj.Id())
    var try_1 EID
    { 
      var v_list2 *ClaireList
      var _Zx *ClaireAny
      var v_local2 *ClaireAny
      v_list2 = self.Args
      try_1 = EID{CreateList(ToType(CEMPTY.Id()),v_list2.Length()).Id(),0}
      for CLcount := 0; CLcount < v_list2.Length(); CLcount++{ 
        _Zx = v_list2.At(CLcount)
        var try_2 EID
        try_2 = Core.F_CALL(C_c_code,ARGS(_Zx.ToEID(),EID{C_any.Id(),0}))
        if ErrorIn(try_2) {try_1 = try_2
        break
        } else {
        v_local2 = ANY(try_2)
        ToList(OBJ(try_1)).PutAt(CLcount,v_local2)
        } 
      }
      } 
    if ErrorIn(try_1) {Result = try_1
    } else {
    va_arg2 = ToList(OBJ(try_1))
    va_arg1.Args = va_arg2
    Result = EID{va_arg2.Id(),0}
    }
    } 
  if !ErrorIn(Result) {
  Result = EID{_CL_obj.Id(),0}
  }
  } 
return Result} 

// The EID go function for: c_code @ Tuple (throw: true) 
func E_c_code_Tuple (self EID) EID { 
return F_c_code_Tuple(Language.To_Tuple(OBJ(self)) )} 

// CLAIRE 4: extended to maps
/* The go function for: c_type(self:Map) [status=0] */
func F_c_type_Map (self *Language.Map) *ClaireType { 
return  ToType(C_map_set.Id())
} 

// The EID go function for: c_type @ Map (throw: false) 
func E_c_type_Map (self EID) EID { 
return EID{F_c_type_Map(Language.To_Map(OBJ(self)) ).Id(),0}} 

// macroexpension of the 
/* The go function for: c_code(self:Map) [status=1] */
func F_c_code_Map (self *Language.Map) EID { 
var Result EID
{ var _Zv *ClaireVariable = F_Compile_Variable_I_symbol(ToSymbol(C_Compile__starname_star.Value),0,C_map_set.Id())
  { var arg_1 *Language.Let
    var try_2 EID
    { var _CL_obj *Language.Let = Language.To_Let(new(Language.Let).Is(Language.C_Let))
      _CL_obj.ClaireVar = _Zv
      { 
        var va_arg1 *Language.Let
        var va_arg2 *ClaireAny
        va_arg1 = _CL_obj
        var try_3 EID
        { var arg_4 *Language.Call
          { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
            _CL_obj.Selector = C_map_I
            _CL_obj.Args = MakeConstantList(self.Domain.Id(),self.Of.Id())
            arg_4 = _CL_obj
            } 
          try_3 = Core.F_CALL(C_c_code,ARGS(EID{arg_4.Id(),0}))
          } 
        if ErrorIn(try_3) {try_2 = try_3
        } else {
        va_arg2 = ANY(try_3)
        va_arg1.Value = va_arg2
        try_2 = va_arg2.ToEID()
        }
        } 
      if !ErrorIn(try_2) {
      { 
        var va_arg1 *Language.Let
        var va_arg2 *ClaireAny
        va_arg1 = _CL_obj
        { var _CL_obj *Language.Do = Language.To_Do(new(Language.Do).Is(Language.C_Do))
          { 
            var va_arg1 *Language.Do
            var va_arg2 *ClaireList
            va_arg1 = _CL_obj
            { var arg_5 *ClaireList
              { 
                var v_list7 *ClaireList
                var x *ClaireAny
                var v_local7 *ClaireAny
                v_list7 = self.Args
                arg_5 = CreateList(ToType(CEMPTY.Id()),v_list7.Length())
                for CLcount := 0; CLcount < v_list7.Length(); CLcount++{ 
                  x = v_list7.At(CLcount)
                  { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                    _CL_obj.Selector = C_put
                    _CL_obj.Args = MakeConstantList(_Zv.Id(),ToPair(x).First,ToPair(x).Second)
                    v_local7 = _CL_obj.Id()
                    } 
                  arg_5.PutAt(CLcount,v_local7)
                  } 
                } 
              va_arg2 = arg_5.AddFast(_Zv.Id())
              } 
            va_arg1.Args = va_arg2
            } 
          va_arg2 = _CL_obj.Id()
          } 
        va_arg1.Arg = va_arg2
        } 
      try_2 = EID{_CL_obj.Id(),0}
      }
      } 
    if ErrorIn(try_2) {Result = try_2
    } else {
    arg_1 = Language.To_Let(OBJ(try_2))
    Result = Core.F_CALL(C_c_code,ARGS(EID{arg_1.Id(),0},EID{C_map_set.Id(),0}))
    }
    } 
  } 
return Result} 

// The EID go function for: c_code @ Map (throw: true) 
func E_c_code_Map (self EID) EID { 
return F_c_code_Map(Language.To_Map(OBJ(self)) )} 

// ******************************************************************
// *      Part 2: Compiling Definitions                             *
// ******************************************************************
/* The go function for: c_type(self:Definition) [status=0] */
func F_c_type_Definition (self *Language.Definition) *ClaireType { 
if (ToType(self.Arg.Id()).Included(ToType(C_exception.Id())) == CTRUE) { 
  return  ToType(CEMPTY.Id())
  } else {
  return  ToType(self.Arg.Id())
  } 
} 

// The EID go function for: c_type @ Definition (throw: false) 
func E_c_type_Definition (self EID) EID { 
return EID{F_c_type_Definition(Language.To_Definition(OBJ(self)) ).Id(),0}} 

// creation of a new object
/* The go function for: c_code(self:Definition,s:class) [status=1] */
func F_c_code_Definition (self *Language.Definition,s *ClaireClass) EID { 
var Result EID
{ var _Zc *ClaireClass = self.Arg
  { var _Zv *ClaireVariable
    { var arg_1 int
      C_OPT.MaxVars = (C_OPT.MaxVars+1)
      arg_1 = 0
      _Zv = F_Compile_Variable_I_symbol(ToSymbol(C_Compile__starname_star.Value),arg_1,_Zc.Id())
      } 
    { var _Zx *ClaireAny
      var try_2 EID
      try_2 = F_Optimize_total_ask_class(_Zc,self.Args)
      if ErrorIn(try_2) {Result = try_2
      } else {
      _Zx = ANY(try_2)
      if (_Zc.Open <= 0) { 
        Result = ToException(Core.C_general_error.Make(MakeString("[105] cannot instantiate ~S").Id(),MakeConstantList(_Zc.Id()).Id())).Close()
        } else {
        Result = EID{CFALSE.Id(),0}
        } 
      if !ErrorIn(Result) {
      if (F_boolean_I_any(_Zx) == CTRUE) { 
        Result = Core.F_CALL(C_c_code,ARGS(_Zx.ToEID(),EID{s.Id(),0}))
        } else {
        { var arg_3 *Language.Let
          var try_4 EID
          { var _CL_obj *Language.Let = Language.To_Let(new(Language.Let).Is(Language.C_Let))
            _CL_obj.ClaireVar = _Zv
            { 
              var va_arg1 *Language.Let
              var va_arg2 *ClaireAny
              va_arg1 = _CL_obj
              var try_5 EID
              { var _CL_obj *Language.Cast = Language.To_Cast(new(Language.Cast).Is(Language.C_Cast))
                { 
                  var va_arg1 *Language.Cast
                  var va_arg2 *ClaireAny
                  va_arg1 = _CL_obj
                  var try_6 EID
                  { var arg_7 *Language.Call
                    { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                      _CL_obj.Selector = C_mClaire_new_I
                      _CL_obj.Args = MakeConstantList(_Zc.Id())
                      arg_7 = _CL_obj
                      } 
                    try_6 = Core.F_CALL(C_c_code,ARGS(EID{arg_7.Id(),0},EID{C_object.Id(),0}))
                    } 
                  if ErrorIn(try_6) {try_5 = try_6
                  } else {
                  va_arg2 = ANY(try_6)
                  va_arg1.Arg = va_arg2
                  try_5 = va_arg2.ToEID()
                  }
                  } 
                if !ErrorIn(try_5) {
                _CL_obj.SetArg = ToType(_Zc.Id())
                try_5 = EID{_CL_obj.Id(),0}
                }
                } 
              if ErrorIn(try_5) {try_4 = try_5
              } else {
              va_arg2 = ANY(try_5)
              va_arg1.Value = va_arg2
              try_4 = va_arg2.ToEID()
              }
              } 
            if !ErrorIn(try_4) {
            { 
              var va_arg1 *Language.Let
              var va_arg2 *ClaireAny
              va_arg1 = _CL_obj
              var try_8 EID
              { var arg_9 *ClaireAny
                var try_10 EID
                try_10 = F_Optimize_analyze_I_class(_Zc,_Zv.Id(),self.Args,ToType(CEMPTY.Id()).EmptyList())
                if ErrorIn(try_10) {try_8 = try_10
                } else {
                arg_9 = ANY(try_10)
                try_8 = F_Compile_Do_I_list(ToList(arg_9)).ToEID()
                }
                } 
              if ErrorIn(try_8) {try_4 = try_8
              } else {
              va_arg2 = ANY(try_8)
              va_arg1.Arg = va_arg2
              try_4 = va_arg2.ToEID()
              }
              } 
            if !ErrorIn(try_4) {
            try_4 = EID{_CL_obj.Id(),0}
            }}
            } 
          if ErrorIn(try_4) {Result = try_4
          } else {
          arg_3 = Language.To_Let(OBJ(try_4))
          Result = Core.F_CALL(C_c_code,ARGS(EID{arg_3.Id(),0},EID{s.Id(),0}))
          }
          } 
        } 
      }
      }
      } 
    } 
  } 
return Result} 

// The EID go function for: c_code @ Definition (throw: true) 
func E_c_code_Definition (self EID,s EID) EID { 
return F_c_code_Definition(Language.To_Definition(OBJ(self)),ToClass(OBJ(s)) )} 

// tells if a "total instantiation" is appropriate (for exceptions)
// we actually check that the srange is OID or integer for all slots
// CLAIRE4: check that the order is respected !
/* The go function for: total?(self:class,l:list) [status=1] */
func F_Optimize_total_ask_class (self *ClaireClass,l *ClaireList) EID { 
var Result EID
{ var lp *ClaireList
  var try_1 EID
  try_1 = Core.F_CALL(C_Compile_get_indexed,ARGS(EID{self.Id(),0}))
  if ErrorIn(try_1) {Result = try_1
  } else {
  lp = ToList(OBJ(try_1))
  { var n int = lp.Length()
    var g0298I *ClaireBoolean
    { 
      var v_and2 *ClaireBoolean
      
      v_and2 = Equal(MakeInteger(l.Length()).Id(),MakeInteger((n-1)).Id())
      if (v_and2 == CFALSE) {g0298I = CFALSE
      } else { 
        { var arg_2 *ClaireAny
          { var i int = 2
            { var g0296 int = n
              arg_2= CFALSE.Id()
              for (i <= g0296) { 
                if (Equal(ANY(Core.F_CALL(C_selector,ARGS(lp.At(i-1).ToEID()))),Language.To_Call(l.At((i-1)-1)).Args.At(0)) != CTRUE) { 
                  arg_2 = CTRUE.Id()
                  break
                  } 
                i = (i+1)
                } 
              } 
            } 
          v_and2 = Core.F_not_any(arg_2)
          } 
        if (v_and2 == CFALSE) {g0298I = CFALSE
        } else { 
          v_and2 = MakeBoolean((self.Open == ClEnv.Default) || (ToType(self.Id()).Included(ToType(C_exception.Id())) == CTRUE))
          if (v_and2 == CFALSE) {g0298I = CFALSE
          } else { 
            v_and2 = Core.F__inf_equal_integer(n,4)
            if (v_and2 == CFALSE) {g0298I = CFALSE
            } else { 
              { var arg_3 *ClaireAny
                { var i int = 2
                  { var g0297 int = n
                    arg_3= CFALSE.Id()
                    for (i <= g0297) { 
                      if ((ANY(Core.F_CALL(C_mClaire_srange,ARGS(lp.At(i-1).ToEID()))) != C_any.Id()) && 
                          (ANY(Core.F_CALL(C_mClaire_srange,ARGS(lp.At(i-1).ToEID()))) != C_integer.Id())) { 
                        arg_3 = CTRUE.Id()
                        break
                        } 
                      i = (i+1)
                      } 
                    } 
                  } 
                v_and2 = Core.F_not_any(arg_3)
                } 
              if (v_and2 == CFALSE) {g0298I = CFALSE
              } else { 
                g0298I = CTRUE} 
              } 
            } 
          } 
        } 
      } 
    if (g0298I == CTRUE) { 
      { var _Zc *ClaireAny
        var try_4 EID
        { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
          _CL_obj.Selector = ToProperty(IfThenElse((l.Length() == 0),
            C_mClaire_new_I.Id(),
            C_Compile_anyObject_I.Id()))
          { 
            var va_arg1 *Language.Call
            var va_arg2 *ClaireList
            va_arg1 = _CL_obj
            var try_5 EID
            { var arg_6 *ClaireList
              var try_7 EID
              { 
                var v_list7 *ClaireList
                var x *ClaireAny
                var v_local7 *ClaireAny
                v_list7 = l
                try_7 = EID{CreateList(ToType(CEMPTY.Id()),v_list7.Length()).Id(),0}
                for CLcount := 0; CLcount < v_list7.Length(); CLcount++{ 
                  x = v_list7.At(CLcount)
                  var try_8 EID
                  try_8 = Core.F_CALL(C_c_code,ARGS(ToList(OBJ(Core.F_CALL(C_args,ARGS(x.ToEID())))).At(1).ToEID(),EID{C_any.Id(),0}))
                  if ErrorIn(try_8) {try_7 = try_8
                  break
                  } else {
                  v_local7 = ANY(try_8)
                  ToList(OBJ(try_7)).PutAt(CLcount,v_local7)
                  } 
                }
                } 
              if ErrorIn(try_7) {try_5 = try_7
              } else {
              arg_6 = ToList(OBJ(try_7))
              try_5 = EID{F_cons_any(self.Id(),arg_6).Id(),0}
              }
              } 
            if ErrorIn(try_5) {try_4 = try_5
            } else {
            va_arg2 = ToList(OBJ(try_5))
            va_arg1.Args = va_arg2
            try_4 = EID{va_arg2.Id(),0}
            }
            } 
          if !ErrorIn(try_4) {
          try_4 = EID{_CL_obj.Id(),0}
          }
          } 
        if ErrorIn(try_4) {Result = try_4
        } else {
        _Zc = ANY(try_4)
        { var m *ClaireAny = Core.F__at_property1(C_close,self).Id()
          if (l.Length() == 0) { 
            var try_9 EID
            try_9 = Core.F_CALL(C_c_code,ARGS(_Zc.ToEID()))
            if ErrorIn(try_9) {Result = try_9
            } else {
            _Zc = ANY(try_9)
            Result = _Zc.ToEID()
            }
            } else {
            Result = EID{CFALSE.Id(),0}
            } 
          if !ErrorIn(Result) {
          if (F_boolean_I_any(m) == CTRUE) { 
            { var _CL_obj *Language.CallMethod1 = Language.To_CallMethod1(new(Language.CallMethod1).Is(Language.C_Call_method1))
              _CL_obj.Arg = ToMethod(m)
              _CL_obj.Args = MakeConstantList(_Zc)
              Result = EID{_CL_obj.Id(),0}
              } 
            } else {
            Result = _Zc.ToEID()
            } 
          }
          } 
        }
        } 
      } else {
      Result = EID{CFALSE.Id(),0}
      } 
    } 
  }
  } 
return Result} 

// The EID go function for: total? @ class (throw: true) 
func E_Optimize_total_ask_class (self EID,l EID) EID { 
return F_Optimize_total_ask_class(ToClass(OBJ(self)),ToList(OBJ(l)) )} 

// the instantiation body is a sequence of words from which the initialization
// of the object must be built. This method produces a list of CLAIRE instructions
// self is the object (if named) or a variable if unamed
// lp will become the list of properties with explicit value setup 
// in CLAIRE 4, we assume that instantiation will put all the default values
// so we need to add write(p,self,def) for all default of p not in lp with complex (inverse or rules) management 
/* The go function for: analyze!(c:class,self:any,%l:list,lp:list) [status=1] */
func F_Optimize_analyze_I_class (c *ClaireClass,self *ClaireAny,_Zl *ClaireList,lp *ClaireList) EID { 
var Result EID
{ var ins_ask *ClaireBoolean = MakeBoolean(((c.Open == 3) || 
      (c.Open == 1)) && (F_boolean_I_any(lp.Id()).Id() != CTRUE.Id()))
  _ = ins_ask
  { var r *ClaireList
    var try_1 EID
    { 
      var v_list2 *ClaireList
      var x *Language.Call
      var v_local2 *ClaireAny
      v_list2 = _Zl
      try_1 = EID{CreateList(ToType(C_any.Id()),v_list2.Length()).Id(),0}
      for CLcount := 0; CLcount < v_list2.Length(); CLcount++{ 
        x = Language.To_Call(v_list2.At(CLcount))
        var try_2 EID
        { var p *ClaireAny = x.Args.At(0)
          { var y *ClaireAny = x.Args.At(1)
            { var s *ClaireObject = Core.F__at_property1(ToProperty(p),c)
              { var special_ask *ClaireBoolean = MakeBoolean(ANY(Core.F_CALL(C_open,ARGS(p.ToEID()))).IsInt(0) && (C_slot.Id() == s.Isa.Id()))
                lp = lp.AddFast(p)
                { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                  _CL_obj.Selector = ToProperty(IfThenElse((special_ask == CTRUE),
                    C_put.Id(),
                    Core.C_write.Id()))
                  { 
                    var va_arg1 *Language.Call
                    var va_arg2 *ClaireList
                    va_arg1 = _CL_obj
                    var try_3 EID
                    { 
                      var v_bag_arg *ClaireAny
                      try_3= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                      ToList(OBJ(try_3)).AddFast(p)
                      ToList(OBJ(try_3)).AddFast(self)
                      var try_4 EID
                      var g0299I *ClaireBoolean
                      var try_5 EID
                      { 
                        var v_or11 *ClaireBoolean
                        
                        v_or11 = special_ask.Not
                        if (v_or11 == CTRUE) {try_5 = EID{CTRUE.Id(),0}
                        } else { 
                          var try_6 EID
                          { var arg_7 *ClaireType
                            var try_8 EID
                            try_8 = Core.F_CALL(C_c_type,ARGS(y.ToEID()))
                            if ErrorIn(try_8) {try_6 = try_8
                            } else {
                            arg_7 = ToType(OBJ(try_8))
                            try_6 = EID{arg_7.Included(ToType(OBJ(Core.F_CALL(C_range,ARGS(EID{s.Id(),0}))))).Id(),0}
                            }
                            } 
                          if ErrorIn(try_6) {try_5 = try_6
                          } else {
                          v_or11 = ToBoolean(OBJ(try_6))
                          if (v_or11 == CTRUE) {try_5 = EID{CTRUE.Id(),0}
                          } else { 
                            try_5 = EID{CFALSE.Id(),0}} 
                          } 
                        }
                        } 
                      if ErrorIn(try_5) {try_4 = try_5
                      } else {
                      g0299I = ToBoolean(OBJ(try_5))
                      if (g0299I == CTRUE) { 
                        try_4 = y.ToEID()
                        } else {
                        try_4 = Core.F_CALL(C_c_code,ARGS(y.ToEID(),EID{C_any.Id(),0}))
                        } 
                      }
                      if ErrorIn(try_4) {try_3 = try_4
                      } else {
                      v_bag_arg = ANY(try_4)
                      ToList(OBJ(try_3)).AddFast(v_bag_arg)}
                      } 
                    if ErrorIn(try_3) {try_2 = try_3
                    } else {
                    va_arg2 = ToList(OBJ(try_3))
                    va_arg1.Args = va_arg2
                    try_2 = EID{va_arg2.Id(),0}
                    }
                    } 
                  if ErrorIn(try_2) {try_1 = try_2
                  break
                  } else {
                  try_2 = EID{_CL_obj.Id(),0}
                  }
                  } 
                if ErrorIn(try_2) {try_1 = try_2
                break
                } else {
                }
                } 
              } 
            } 
          } 
        if ErrorIn(try_2) {try_1 = try_2
        break
        } else {
        v_local2 = ANY(try_2)
        ToList(OBJ(try_1)).PutAt(CLcount,v_local2)
        } 
      }
      } 
    if ErrorIn(try_1) {Result = try_1
    } else {
    r = ToList(OBJ(try_1))
    { 
      var s *ClaireSlot
      _ = s
      var s_iter *ClaireAny
      Result= EID{CFALSE.Id(),0}
      var s_support *ClaireList
      var try_9 EID
      try_9 = Core.F_CALL(C_Compile_get_indexed,ARGS(EID{c.Id(),0}))
      if ErrorIn(try_9) {Result = try_9
      } else {
      s_support = ToList(OBJ(try_9))
      s_len := s_support.Length()
      for i_it := 0; i_it < s_len; i_it++ { 
        s_iter = s_support.At(i_it)
        s = ToSlot(s_iter)
        var loop_10 EID
        _ = loop_10
        { var p *ClaireProperty = s.Selector
          { var v *ClaireAny = s.Default
            if ((v != CNULL) && 
                ((lp.Memq(p.Id()) != CTRUE) && 
                  ((p.Inverse.Id() != CNULL) || 
                      (p.IfWrite != CNULL)))) { 
              { var defExp *ClaireAny
                var try_11 EID
                var g0300I *ClaireBoolean
                var try_12 EID
                try_12 = F_Compile_designated_ask_any(v)
                if ErrorIn(try_12) {try_11 = try_12
                } else {
                g0300I = ToBoolean(OBJ(try_12))
                if (g0300I == CTRUE) { 
                  try_11 = v.ToEID()
                  } else {
                  { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                    _CL_obj.Selector = C_default
                    { 
                      var va_arg1 *Language.Call
                      var va_arg2 *ClaireList
                      va_arg1 = _CL_obj
                      { 
                        var v_bag_arg *ClaireAny
                        va_arg2= ToType(CEMPTY.Id()).EmptyList()
                        { var _CL_obj *Language.Cast = Language.To_Cast(new(Language.Cast).Is(Language.C_Cast))
                          { 
                            var va_arg1 *Language.Cast
                            var va_arg2 *ClaireAny
                            va_arg1 = _CL_obj
                            { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                              _CL_obj.Selector = ToProperty(Core.C__at.Id())
                              _CL_obj.Args = MakeConstantList(p.Id(),c.Id())
                              va_arg2 = _CL_obj.Id()
                              } 
                            va_arg1.Arg = va_arg2
                            } 
                          _CL_obj.SetArg = ToType(C_slot.Id())
                          v_bag_arg = _CL_obj.Id()
                          } 
                        va_arg2.AddFast(v_bag_arg)} 
                      va_arg1.Args = va_arg2
                      } 
                    try_11 = EID{_CL_obj.Id(),0}
                    } 
                  } 
                }
                if ErrorIn(try_11) {loop_10 = try_11
                } else {
                defExp = ANY(try_11)
                { var arg_13 *Language.Call
                  { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                    _CL_obj.Selector = Core.C_write
                    _CL_obj.Args = MakeConstantList(p.Id(),self,defExp)
                    arg_13 = _CL_obj
                    } 
                  r = r.AddFast(arg_13.Id())
                  } 
                loop_10 = EID{r.Id(),0}
                }
                } 
              } else {
              loop_10 = EID{CFALSE.Id(),0}
              } 
            } 
          } 
        if ErrorIn(loop_10) {Result = loop_10
        break
        } else {
        }}
        } 
      } 
    if !ErrorIn(Result) {
    { var m *ClaireAny = Core.F__at_property1(C_close,c).Id()
      { var arg_14 *ClaireAny
        if (F_boolean_I_any(m) == CTRUE) { 
          { var _CL_obj *Language.CallMethod1 = Language.To_CallMethod1(new(Language.CallMethod1).Is(Language.C_Call_method1))
            _CL_obj.Arg = ToMethod(m)
            _CL_obj.Args = MakeConstantList(self)
            arg_14 = _CL_obj.Id()
            } 
          } else {
          arg_14 = self
          } 
        r = r.AddFast(arg_14)
        } 
      } 
    Result = EID{r.Id(),0}
    }
    }
    } 
  } 
return Result} 

// The EID go function for: analyze! @ class (throw: true) 
func E_Optimize_analyze_I_class (c EID,self EID,_Zl EID,lp EID) EID { 
return F_Optimize_analyze_I_class(ToClass(OBJ(c)),
  ANY(self),
  ToList(OBJ(_Zl)),
  ToList(OBJ(lp)) )} 

// creation of a new named object
// CLAIRE4 : native variable need a specific 
/* The go function for: c_code(self:Defobj,s:class) [status=1] */
func F_c_code_Defobj (self *Language.Defobj,s *ClaireClass) EID { 
var Result EID
{ var _Zc *ClaireClass = self.Arg
  { var o *ClaireAny = self.Ident.Value()
    { var _Zx *ClaireAny = CNULL
      var g0305I *ClaireBoolean
      if (o.Isa.IsIn(Core.C_global_variable) == CTRUE) { 
        { var g0303 *Core.GlobalVariable = Core.ToGlobalVariable(o)
          g0305I = F_Compile_nativeVar_ask_global_variable(g0303)
          } 
        } else {
        g0305I = CFALSE
        } 
      if (g0305I == CTRUE) { 
        { var _Zval *ClaireAny = CNULL
          { 
            var c *ClaireAny
            _ = c
            var c_support *ClaireList
            c_support = self.Args
            c_len := c_support.Length()
            for i_it := 0; i_it < c_len; i_it++ { 
              c = c_support.At(i_it)
              var g0306I *ClaireBoolean
              if (c.Isa.IsIn(Language.C_Call) == CTRUE) { 
                { var g0304 *Language.Call = Language.To_Call(c)
                  g0306I = MakeBoolean((g0304.Selector.Id() == C__equal.Id()) && (g0304.Args.At(0) == C_value.Id()))
                  } 
                } else {
                g0306I = CFALSE
                } 
              if (g0306I == CTRUE) { 
                _Zval = ToList(OBJ(Core.F_CALL(C_args,ARGS(c.ToEID())))).At(1)
                } 
              } 
            } 
          var try_1 EID
          { var _CL_obj *Language.Gassign = Language.To_Gassign(new(Language.Gassign).Is(Language.C_Gassign))
            _CL_obj.ClaireVar = Core.ToGlobalVariable(o)
            { 
              var va_arg1 *Language.Gassign
              var va_arg2 *ClaireAny
              va_arg1 = _CL_obj
              var try_2 EID
              { var arg_3 *ClaireClass
                var try_4 EID
                { var arg_5 *ClaireType
                  var try_6 EID
                  try_6 = Core.F_CALL(C_c_type,ARGS(o.ToEID()))
                  if ErrorIn(try_6) {try_4 = try_6
                  } else {
                  arg_5 = ToType(OBJ(try_6))
                  try_4 = EID{arg_5.Class_I().Id(),0}
                  }
                  } 
                if ErrorIn(try_4) {try_2 = try_4
                } else {
                arg_3 = ToClass(OBJ(try_4))
                try_2 = Core.F_CALL(C_c_code,ARGS(_Zval.ToEID(),EID{arg_3.Id(),0}))
                }
                } 
              if ErrorIn(try_2) {try_1 = try_2
              } else {
              va_arg2 = ANY(try_2)
              va_arg1.Arg = va_arg2
              try_1 = va_arg2.ToEID()
              }
              } 
            if !ErrorIn(try_1) {
            try_1 = EID{_CL_obj.Id(),0}
            }
            } 
          if ErrorIn(try_1) {Result = try_1
          } else {
          _Zx = ANY(try_1)
          Result = _Zx.ToEID()
          }
          } 
        } else {
        { var _Zv *ClaireAny
          if ((o != CNULL) && 
              (o.Isa.IsIn(Core.C_global_variable) != CTRUE)) { 
            _Zv = o
            } else {
            { var arg_7 int
              C_OPT.MaxVars = (C_OPT.MaxVars+1)
              arg_7 = 0
              _Zv = F_Compile_Variable_I_symbol(ToSymbol(C_Compile__starname_star.Value),arg_7,_Zc.Id()).Id()
              } 
            } 
          { var _Zy1 *Language.Call
            { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
              _CL_obj.Selector = C_Compile_object_I
              _CL_obj.Args = MakeConstantList(self.Ident.Id(),_Zc.Id())
              _Zy1 = _CL_obj
              } 
            { var _Zy2 *ClaireAny
              var try_8 EID
              try_8 = F_Optimize_analyze_I_class(_Zc,_Zv,self.Args,MakeConstantList(C_name.Id()))
              if ErrorIn(try_8) {Result = try_8
              } else {
              _Zy2 = ANY(try_8)
              if (_Zv.Isa.IsIn(C_Variable) != CTRUE) { 
                { var _CL_obj *Language.Do = Language.To_Do(new(Language.Do).Is(Language.C_Do))
                  _CL_obj.Args = F_cons_any(_Zy1.Id(),ToList(_Zy2))
                  _Zx = _CL_obj.Id()
                  } 
                } else {
                { var _CL_obj *Language.Let = Language.To_Let(new(Language.Let).Is(Language.C_Let))
                  _CL_obj.ClaireVar = To_Variable(_Zv)
                  _CL_obj.Value = _Zy1.Id()
                  _CL_obj.Arg = F_Compile_Do_I_list(ToList(_Zy2))
                  _Zx = _CL_obj.Id()
                  } 
                } 
              var try_9 EID
              try_9 = Core.F_CALL(C_c_code,ARGS(_Zx.ToEID(),EID{s.Id(),0}))
              if ErrorIn(try_9) {Result = try_9
              } else {
              _Zx = ANY(try_9)
              Result = _Zx.ToEID()
              }
              }
              } 
            } 
          } 
        } 
      if !ErrorIn(Result) {
      
      if (_Zc.Open <= 0) { 
        Result = ToException(Core.C_general_error.Make(MakeString("[105] cannot instantiate ~S").Id(),MakeConstantList(_Zc.Id()).Id())).Close()
        } else {
        Result = EID{CFALSE.Id(),0}
        } 
      if !ErrorIn(Result) {
      if (o != CNULL) { 
        if (ToBoolean(C_OPT.Objects.Contain_ask(o).Id()) != CTRUE) { 
          C_OPT.Objects = C_OPT.Objects.AddFast(o)
          Core.F_CALL(C_Optimize_c_register,ARGS(o.ToEID()))
          } 
        } else {
        F_Compile_warn_void()
        Core.F_tformat_string(MakeString("~S is unknown [265]\n"),2,MakeConstantList(self.Ident.Id()))
        } 
      Result = _Zx.ToEID()
      }}
      } 
    } 
  } 
return Result} 

// The EID go function for: c_code @ Defobj (throw: true) 
func E_c_code_Defobj (self EID,s EID) EID { 
return F_c_code_Defobj(Language.To_Defobj(OBJ(self)),ToClass(OBJ(s)) )} 

// creation of a new named object
/* The go function for: c_code(self:Defclass,s:class) [status=1] */
func F_c_code_Defclass (self *Language.Defclass,s *ClaireClass) EID { 
var Result EID
{ var _Zname *ClaireSymbol = self.Ident
  { var o *ClaireAny = _Zname.Value()
    { var _Zcreate *Language.Call
      var try_1 EID
      if (o != CNULL) { 
        { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
          _CL_obj.Selector = C_class_I
          _CL_obj.Args = MakeConstantList(_Zname.Id(),self.Arg.Id())
          try_1 = EID{_CL_obj.Id(),0}
          } 
        } else {
        try_1 = ToException(Core.C_general_error.Make(MakeString("[internal] cannot compile unknown class ~S").Id(),MakeConstantList(_Zname.Id()).Id())).Close()
        } 
      if ErrorIn(try_1) {Result = try_1
      } else {
      _Zcreate = Language.To_Call(OBJ(try_1))
      { var _Zx *Language.Do
        var try_2 EID
        { var _CL_obj *Language.Do = Language.To_Do(new(Language.Do).Is(Language.C_Do))
          { 
            var va_arg1 *Language.Do
            var va_arg2 *ClaireList
            va_arg1 = _CL_obj
            var try_3 EID
            { var arg_4 *ClaireList
              var try_5 EID
              { var arg_6 *ClaireList
                var try_8 EID
                { 
                  var v_list8 *ClaireList
                  var x *ClaireAny
                  var v_local8 *ClaireAny
                  v_list8 = self.Args
                  try_8 = EID{CreateList(ToType(CEMPTY.Id()),v_list8.Length()).Id(),0}
                  for CLcount := 0; CLcount < v_list8.Length(); CLcount++{ 
                    x = v_list8.At(CLcount)
                    var try_9 EID
                    { var v *ClaireAny = CNULL
                      if (x.Isa.IsIn(Language.C_Call) == CTRUE) { 
                        { var g0308 *Language.Call = Language.To_Call(x)
                          v = g0308.Args.At(1)
                          g0308 = Language.To_Call(g0308.Args.At(0))
                          x = g0308.Id()
                          try_9 = x.ToEID()
                          } 
                        }  else if (x.Isa.IsIn(C_Variable) == CTRUE) { 
                        { var g0309 *ClaireVariable = To_Variable(x)
                          var try_10 EID
                          { var arg_11 *ClaireAny
                            var try_12 EID
                            try_12 = Language.F_extract_type_any(g0309.Range.Id())
                            if ErrorIn(try_12) {try_10 = try_12
                            } else {
                            arg_11 = ANY(try_12)
                            try_10 = Language.F_Language_getDefault_type(ToType(arg_11),v).ToEID()
                            }
                            } 
                          if ErrorIn(try_10) {try_9 = try_10
                          } else {
                          v = ANY(try_10)
                          try_9 = v.ToEID()
                          }
                          } 
                        } else {
                        try_9 = EID{CFALSE.Id(),0}
                        } 
                      if ErrorIn(try_9) {try_8 = try_9
                      break
                      } else {
                      { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                        _CL_obj.Selector = C_add_slot
                        { 
                          var va_arg1 *Language.Call
                          var va_arg2 *ClaireList
                          va_arg1 = _CL_obj
                          var try_13 EID
                          { 
                            var v_bag_arg *ClaireAny
                            try_13= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                            ToList(OBJ(try_13)).AddFast(o)
                            var try_14 EID
                            try_14 = Language.F_make_a_property_any(ANY(Core.F_CALL(C_mClaire_pname,ARGS(x.ToEID()))))
                            if ErrorIn(try_14) {try_13 = try_14
                            } else {
                            v_bag_arg = ANY(try_14)
                            ToList(OBJ(try_13)).AddFast(v_bag_arg)
                            ToList(OBJ(try_13)).AddFast(ANY(Core.F_CALL(C_range,ARGS(x.ToEID()))))
                            ToList(OBJ(try_13)).AddFast(v)}
                            } 
                          if ErrorIn(try_13) {try_9 = try_13
                          } else {
                          va_arg2 = ToList(OBJ(try_13))
                          va_arg1.Args = va_arg2
                          try_9 = EID{va_arg2.Id(),0}
                          }
                          } 
                        if ErrorIn(try_9) {try_8 = try_9
                        break
                        } else {
                        try_9 = EID{_CL_obj.Id(),0}
                        }
                        } 
                      if ErrorIn(try_9) {try_8 = try_9
                      break
                      } else {
                      }}
                      } 
                    if ErrorIn(try_9) {try_8 = try_9
                    break
                    } else {
                    v_local8 = ANY(try_9)
                    ToList(OBJ(try_8)).PutAt(CLcount,v_local8)
                    } 
                  }
                  } 
                if ErrorIn(try_8) {try_5 = try_8
                } else {
                arg_6 = ToList(OBJ(try_8))
                { var arg_7 *ClaireList
                  if (self.Params.Length() != 0) { 
                    { 
                      var v_bag_arg *ClaireAny
                      arg_7= ToType(CEMPTY.Id()).EmptyList()
                      { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                        _CL_obj.Selector = C_put
                        _CL_obj.Args = MakeConstantList(C_params.Id(),o,self.Params.Id())
                        v_bag_arg = _CL_obj.Id()
                        } 
                      arg_7.AddFast(v_bag_arg)} 
                    } else {
                    arg_7 = ToType(CEMPTY.Id()).EmptyList()
                    } 
                  try_5 = EID{arg_6.Append(arg_7).Id(),0}
                  } 
                }
                } 
              if ErrorIn(try_5) {try_3 = try_5
              } else {
              arg_4 = ToList(OBJ(try_5))
              try_3 = EID{F_cons_any(_Zcreate.Id(),arg_4).Id(),0}
              }
              } 
            if ErrorIn(try_3) {try_2 = try_3
            } else {
            va_arg2 = ToList(OBJ(try_3))
            va_arg1.Args = va_arg2
            try_2 = EID{va_arg2.Id(),0}
            }
            } 
          if !ErrorIn(try_2) {
          try_2 = EID{_CL_obj.Id(),0}
          }
          } 
        if ErrorIn(try_2) {Result = try_2
        } else {
        _Zx = Language.To_Do(OBJ(try_2))
        if (ToBoolean(C_OPT.Objects.Contain_ask(o).Id()) != CTRUE) { 
          C_OPT.Objects = C_OPT.Objects.AddFast(o)
          Core.F_CALL(C_Optimize_c_register,ARGS(o.ToEID()))
          } 
        Result = Core.F_CALL(C_c_code,ARGS(EID{_Zx.Id(),0},EID{s.Id(),0}))
        }
        } 
      }
      } 
    } 
  } 
return Result} 

// The EID go function for: c_code @ Defclass (throw: true) 
func E_c_code_Defclass (self EID,s EID) EID { 
return F_c_code_Defclass(Language.To_Defclass(OBJ(self)),ToClass(OBJ(s)) )} 

// method definition
// note (3.4): using the un-compiled code for c_status is weak, it would be much better to
// in CLAIRE4, we pass the method as an argument (needed by the code generator)
/* The go function for: c_type(self:Defmethod) [status=0] */
func F_c_type_Defmethod (self *Language.Defmethod) *ClaireType { 
return  ToType(C_any.Id())
} 

// The EID go function for: c_type @ Defmethod (throw: false) 
func E_c_type_Defmethod (self EID) EID { 
return EID{F_c_type_Defmethod(Language.To_Defmethod(OBJ(self)) ).Id(),0}} 

/* The go function for: c_code(self:Defmethod) [status=1] */
func F_c_code_Defmethod (self *Language.Defmethod) EID { 
var Result EID
{ var px *ClaireProperty = self.Arg.Selector
  { var l *ClaireList = self.Arg.Args
    { var lv *ClaireList
      if ((l.Length() == 1) && 
          (l.At(0) == ClEnv.Id())) { 
        lv = MakeConstantList(F_Compile_Variable_I_symbol(ToSymbol(C_Compile__starname_star.Value),0,C_void.Id()).Id())
        } else {
        lv = l
        } 
      { var ls *ClaireList
        var try_1 EID
        try_1 = F_Optimize_extract_signature_I_list(lv)
        if ErrorIn(try_1) {Result = try_1
        } else {
        ls = ToList(OBJ(try_1))
        { var lrange *ClaireList
          var try_2 EID
          { 
            var v_list5 *ClaireList
            var x *ClaireAny
            var v_local5 *ClaireAny
            var try_3 EID
            try_3 = Language.F_extract_range_any(self.SetArg,lv,ToList(Language.C_LDEF.Value))
            if ErrorIn(try_3) {try_2 = try_3
            } else {
            v_list5 = ToList(OBJ(try_3))
            try_2 = EID{CreateList(ToType(C_any.Id()),v_list5.Length()).Id(),0}
            for CLcount := 0; CLcount < v_list5.Length(); CLcount++{ 
              x = v_list5.At(CLcount)
              v_local5 = x
              ToList(OBJ(try_2)).PutAt(CLcount,v_local5)
              } 
            }
            } 
          if ErrorIn(try_2) {Result = try_2
          } else {
          lrange = ToList(OBJ(try_2))
          { var sdef *ClaireAny
            var try_4 EID
            if ((self.Inline_ask == CTRUE) && 
                (C_compiler.Inline_ask == CTRUE)) { 
              Core.F_print_in_string_void()
              PRINC("lambda[(")
              try_4 = Language.F_ppvariable_list(lv)
              if !ErrorIn(try_4) {
              PRINC("),")
              try_4 = Core.F_CALL(C_print,ARGS(self.Body.ToEID()))
              if !ErrorIn(try_4) {
              PRINC("]")
              try_4 = EVOID
              }}
              if !ErrorIn(try_4) {
              try_4 = Core.F_end_of_string_void()
              }
              } else {
              try_4 = EID{CFALSE.Id(),0}
              } 
            if ErrorIn(try_4) {Result = try_4
            } else {
            sdef = ANY(try_4)
            { var lbody *ClaireList
              var try_5 EID
              try_5 = Language.F_extract_status_any(self.Body)
              if ErrorIn(try_5) {Result = try_5
              } else {
              lbody = ToList(OBJ(try_5))
              { var getm *ClaireObject = ToObject(OBJ(Core.F_CALL(ToProperty(Core.C__at.Id()),ARGS(EID{px.Id(),0},ls.At(1).ToEID()))))
                { var m *ClaireMethod
                  var try_6 EID
                  if (C_method.Id() == getm.Isa.Id()) { 
                    { var g0310 *ClaireMethod = ToMethod(getm.Id())
                      try_6 = EID{g0310.Id(),0}
                      } 
                    } else {
                    try_6 = ToException(Core.C_general_error.Make(MakeString("[internal] the method ~S @ ~S is not known").Id(),MakeConstantList(px.Id(),ls.At(1)).Id())).Close()
                    } 
                  if ErrorIn(try_6) {Result = try_6
                  } else {
                  m = ToMethod(OBJ(try_6))
                  ToArray(lbody.Id()).NthPut(2,Core.F_get_property(C_functional,ToObject(m.Id())))
                  Core.F_put_table(C_Compile_FileOrigin,m.Id(),(F_append_string(F_append_string(ToCompileProducer(C_PRODUCER.Value).CurrentFile,MakeString(".cl:")),F_string_I_integer(ClEnv.NLine))).Id())
                  if ((C_compiler.Inline_ask != CTRUE) && 
                      ((px.Id() == Language.C_Iterate.Id()) || 
                          (px.Id() == Language.C_iterate.Id()))) { 
                    Result = EID{CNIL.Id(),0}
                    } else {
                    var g0312I *ClaireBoolean
                    var try_7 EID
                    { 
                      var v_and10 *ClaireBoolean
                      
                      v_and10 = Equal(lrange.At(0),C_void.Id())
                      if (v_and10 == CFALSE) {try_7 = EID{CFALSE.Id(),0}
                      } else { 
                        var try_8 EID
                        try_8 = F_Optimize_sort_pattern_ask_list(lv,self.Body)
                        if ErrorIn(try_8) {try_7 = try_8
                        } else {
                        v_and10 = ToBoolean(OBJ(try_8))
                        if (v_and10 == CFALSE) {try_7 = EID{CFALSE.Id(),0}
                        } else { 
                          try_7 = EID{CTRUE.Id(),0}} 
                        } 
                      }
                      } 
                    if ErrorIn(try_7) {Result = try_7
                    } else {
                    g0312I = ToBoolean(OBJ(try_7))
                    if (g0312I == CTRUE) { 
                      Result = F_Optimize_sort_code_Defmethod(self,lv)
                      } else {
                      
                      if (lbody.At(2) != C_body.Id()) { 
                        { var na *ClaireString
                          var try_9 EID
                          try_9 = Core.F_CALL(C_Compile_function_name,ARGS(EID{px.Id(),0},ls.At(1).ToEID(),lbody.At(1).ToEID()))
                          if ErrorIn(try_9) {Result = try_9
                          } else {
                          na = ToString(OBJ(try_9))
                          { var la *ClaireLambda
                            var try_10 EID
                            try_10 = Language.F_lambda_I_list(lv,lbody.At(2))
                            if ErrorIn(try_10) {Result = try_10
                            } else {
                            la = ToLambda(OBJ(try_10))
                            { var news int
                              var try_11 EID
                              var g0313I *ClaireBoolean
                              var try_12 EID
                              if (C_OPT.Recompute == CTRUE) { 
                                try_12 = F_Compile_g_throw_any(lbody.At(1))
                                } else {
                                try_12 = F_Compile_can_throw_ask_method(m)
                                } 
                              if ErrorIn(try_12) {try_11 = try_12
                              } else {
                              g0313I = ToBoolean(OBJ(try_12))
                              if (g0313I == CTRUE) { 
                                try_11 = EID{C__INT,IVAL(1)}
                                } else {
                                try_11 = EID{C__INT,IVAL(0)}
                                } 
                              }
                              if ErrorIn(try_11) {Result = try_11
                              } else {
                              news = INT(try_11)
                              Result = F_Compile_compile_lambda_string(na,la,m.Id())
                              if !ErrorIn(Result) {
                              if ((lbody.At(0) == CNULL) || 
                                  (C_OPT.Recompute == CTRUE)) { 
                                ToArray(lbody.Id()).NthPut(1,MakeInteger(news).Id())
                                } 
                              Result = ToArray(lbody.Id()).NthPut(2,F_make_function_string(na).Id()).ToEID()
                              }
                              }
                              } 
                            }
                            } 
                          }
                          } 
                        } else {
                        Result = EID{CFALSE.Id(),0}
                        } 
                      if !ErrorIn(Result) {
                      if (self.SetArg.Isa.IsIn(Core.C_global_variable) == CTRUE) { 
                        ToArray(lrange.Id()).NthPut(1,self.SetArg)
                        }  else if ((C_class.Id() == m.Range.Isa.Id()) && 
                          (C_class.Id() != lrange.At(0).Isa.Id())) { 
                        ToArray(lrange.Id()).NthPut(1,m.Range.Id())
                        } 
                      { var _Zm *ClaireAny
                        var try_13 EID
                        try_13 = F_Optimize_add_method_I_method(m,
                          ToList(ls.At(0)),
                          lrange.At(0),
                          lbody.At(0),
                          ToFunction(lbody.At(1)))
                        if ErrorIn(try_13) {Result = try_13
                        } else {
                        _Zm = ANY(try_13)
                        { var arg_14 *ClaireAny
                          var try_15 EID
                          if ((self.Inline_ask == CTRUE) && 
                              (C_compiler.Inline_ask == CTRUE)) { 
                            { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                              _CL_obj.Selector = Core.C_inlineok_ask
                              _CL_obj.Args = MakeConstantList(_Zm,sdef)
                              try_15 = EID{_CL_obj.Id(),0}
                              } 
                            }  else if (F_boolean_I_any(lrange.At(1)) == CTRUE) { 
                            { var na *ClaireString
                              var try_16 EID
                              { var arg_17 *ClaireAny
                                var try_18 EID
                                try_18 = Core.F_CALL(C_Compile_function_name,ARGS(EID{px.Id(),0},ls.At(1).ToEID(),lbody.At(1).ToEID()))
                                if ErrorIn(try_18) {try_16 = try_18
                                } else {
                                arg_17 = ANY(try_18)
                                try_16 = EID{F_Optimize_type_extension_string(ToString(arg_17)).Id(),0}
                                }
                                } 
                              if ErrorIn(try_16) {try_15 = try_16
                              } else {
                              na = ToString(OBJ(try_16))
                              { var _Zf *ClaireFunction = F_make_function_string(na)
                                try_15 = F_Compile_compile_lambda_string(na,ToLambda(lrange.At(1)),C_type.Id())
                                if !ErrorIn(try_15) {
                                F_set_arity_function(_Zf,m.Domain.Length())
                                { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                  _CL_obj.Selector = Core.C_write
                                  _CL_obj.Args = MakeConstantList(Language.C_iClaire_typing.Value,_Zm,_Zf.Id())
                                  try_15 = EID{_CL_obj.Id(),0}
                                  } 
                                }
                                } 
                              }
                              } 
                            } else {
                            try_15 = _Zm.ToEID()
                            } 
                          if ErrorIn(try_15) {Result = try_15
                          } else {
                          arg_14 = ANY(try_15)
                          Result = Core.F_CALL(C_c_code,ARGS(arg_14.ToEID()))
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
return Result} 

// The EID go function for: c_code @ Defmethod (throw: true) 
func E_c_code_Defmethod (self EID) EID { 
return F_c_code_Defmethod(Language.To_Defmethod(OBJ(self)) )} 

// create a type function name by adding a "_type" - will not be imported
/* The go function for: type_extension(s:string) [status=0] */
func F_Optimize_type_extension_string (s *ClaireString) *ClaireString { 
var Result *ClaireString
{ var n int = F_length_string(s)
  { var f *ClaireString
    if (s.At(1) == '#') { 
      f = F_substring_string(s,2,n)
      } else {
      f = s
      } 
    Result = F_append_string(f,MakeString("_type"))
    } 
  } 
return Result} 

// The EID go function for: type_extension @ string (throw: false) 
func E_Optimize_type_extension_string (s EID) EID { 
return EID{F_Optimize_type_extension_string(ToString(OBJ(s)) ).Id(),0}} 

// v3.3 : we optimize a single sort definition -----------------------------------------------
// [foo(x:list) : list -> sort(m,x) ]
/* The go function for: sort_pattern?(lv:list,%body:any) [status=1] */
func F_Optimize_sort_pattern_ask_list (lv *ClaireList,_Zbody *ClaireAny) EID { 
var Result EID
{ 
  var v_and0 *ClaireBoolean
  
  v_and0 = Equal(MakeInteger(lv.Length()).Id(),MakeInteger(1).Id())
  if (v_and0 == CFALSE) {Result = EID{CFALSE.Id(),0}
  } else { 
    var try_1 EID
    if (_Zbody.Isa.IsIn(Language.C_Call) == CTRUE) { 
      { var g0314 *Language.Call = Language.To_Call(_Zbody)
        { 
          var v_and4 *ClaireBoolean
          
          v_and4 = Equal(g0314.Selector.Id(),Core.C_sort.Id())
          if (v_and4 == CFALSE) {try_1 = EID{CFALSE.Id(),0}
          } else { 
            { var a1 *ClaireAny = g0314.Args.At(0)
              if (a1.Isa.IsIn(Language.C_Call) == CTRUE) { 
                { var g0315 *Language.Call = Language.To_Call(a1)
                  v_and4 = MakeBoolean((g0315.Selector.Id() == Core.C__at.Id()) && (g0315.Args.At(0).Isa.IsIn(C_property) == CTRUE))
                  } 
                } else {
                v_and4 = CFALSE
                } 
              } 
            if (v_and4 == CFALSE) {try_1 = EID{CFALSE.Id(),0}
            } else { 
              var try_2 EID
              { var arg_3 *ClaireAny
                var try_4 EID
                try_4 = Language.F_iClaire_lexical_index_any2(g0314.Args.At(1),lv,0,CFALSE)
                if ErrorIn(try_4) {try_2 = try_4
                } else {
                arg_3 = ANY(try_4)
                try_2 = EID{Equal(arg_3,lv.At(0)).Id(),0}
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
          } 
        } 
      } else {
      try_1 = EID{CFALSE.Id(),0}
      } 
    if ErrorIn(try_1) {Result = try_1
    } else {
    v_and0 = ToBoolean(OBJ(try_1))
    if (v_and0 == CFALSE) {Result = EID{CFALSE.Id(),0}
    } else { 
      Result = EID{CTRUE.Id(),0}} 
    } 
  }
  } 
return Result} 

// The EID go function for: sort_pattern? @ list (throw: true) 
func E_Optimize_sort_pattern_ask_list (lv EID,_Zbody EID) EID { 
return F_Optimize_sort_pattern_ask_list(ToList(OBJ(lv)),ANY(_Zbody) )} 

// this is the macroexpansion of the quick_sort which is difficult because of the dual recursion
// Thus, we generate two methods for one definition, and produce the explicit code for the specialized
// quicksort (v3.3)
/* The go function for: sort_code(self:Defmethod,lv:list) [status=1] */
func F_Optimize_sort_code_Defmethod (self *Language.Defmethod,lv *ClaireList) EID { 
var Result EID
{ var l *ClaireAny = lv.At(0)
  { var f *ClaireAny = ToList(OBJ(Core.F_CALL(C_args,ARGS(ToList(OBJ(Core.F_CALL(C_args,ARGS(self.Body.ToEID())))).At(0).ToEID())))).At(0)
    { var m *ClaireVariable = F_Compile_Variable_I_symbol(Core.F_symbol_I_string2(MakeString("m")),0,C_integer.Id())
      { var n *ClaireVariable = F_Compile_Variable_I_symbol(Core.F_symbol_I_string2(MakeString("n")),0,C_integer.Id())
        { var x *ClaireVariable = F_Compile_Variable_I_symbol(Core.F_symbol_I_string2(MakeString("x")),0,Core.F_member_type(ToType(OBJ(Core.F_CALL(C_range,ARGS(l.ToEID()))))).Id())
          { var p *ClaireVariable = F_Compile_Variable_I_symbol(Core.F_symbol_I_string2(MakeString("p")),0,C_integer.Id())
            { var q *ClaireVariable = F_Compile_Variable_I_symbol(Core.F_symbol_I_string2(MakeString("q")),0,C_integer.Id())
              { var def1 *Language.Defmethod
                { var _CL_obj *Language.Defmethod = Language.To_Defmethod(new(Language.Defmethod).Is(Language.C_Defmethod))
                  _CL_obj.Arg = self.Arg
                  _CL_obj.Inline_ask = CFALSE
                  _CL_obj.SetArg = self.SetArg
                  { 
                    var va_arg1 *Language.Defmethod
                    var va_arg2 *ClaireAny
                    va_arg1 = _CL_obj
                    { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                      _CL_obj.Selector = self.Arg.Selector
                      { 
                        var va_arg1 *Language.Call
                        var va_arg2 *ClaireList
                        va_arg1 = _CL_obj
                        { 
                          var v_bag_arg *ClaireAny
                          va_arg2= ToType(CEMPTY.Id()).EmptyList()
                          va_arg2.AddFast(MakeInteger(1).Id())
                          { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                            _CL_obj.Selector = C_length
                            _CL_obj.Args = MakeConstantList(lv.At(0))
                            v_bag_arg = _CL_obj.Id()
                            } 
                          va_arg2.AddFast(v_bag_arg)
                          va_arg2.AddFast(l)} 
                        va_arg1.Args = va_arg2
                        } 
                      va_arg2 = _CL_obj.Id()
                      } 
                    va_arg1.Body = va_arg2
                    } 
                  def1 = _CL_obj
                  } 
                { var _Zbd *Language.If
                  { var _CL_obj *Language.If = Language.To_If(new(Language.If).Is(Language.C_If))
                    { 
                      var va_arg1 *Language.If
                      var va_arg2 *ClaireAny
                      va_arg1 = _CL_obj
                      { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                        _CL_obj.Selector = ToProperty(C__sup.Id())
                        _CL_obj.Args = MakeConstantList(m.Id(),n.Id())
                        va_arg2 = _CL_obj.Id()
                        } 
                      va_arg1.Test = va_arg2
                      } 
                    { 
                      var va_arg1 *Language.If
                      var va_arg2 *ClaireAny
                      va_arg1 = _CL_obj
                      { var _CL_obj *Language.Let = Language.To_Let(new(Language.Let).Is(Language.C_Let))
                        _CL_obj.ClaireVar = x
                        { 
                          var va_arg1 *Language.Let
                          var va_arg2 *ClaireAny
                          va_arg1 = _CL_obj
                          { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                            _CL_obj.Selector = C_nth
                            _CL_obj.Args = MakeConstantList(l,n.Id())
                            va_arg2 = _CL_obj.Id()
                            } 
                          va_arg1.Value = va_arg2
                          } 
                        { 
                          var va_arg1 *Language.Let
                          var va_arg2 *ClaireAny
                          va_arg1 = _CL_obj
                          { var arg_1 *Language.Call
                            { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                              _CL_obj.Selector = ToProperty(C__equal.Id())
                              { 
                                var va_arg1 *Language.Call
                                var va_arg2 *ClaireList
                                va_arg1 = _CL_obj
                                { 
                                  var v_bag_arg *ClaireAny
                                  va_arg2= ToType(CEMPTY.Id()).EmptyList()
                                  va_arg2.AddFast(m.Id())
                                  { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                    _CL_obj.Selector = ToProperty(Core.C__plus.Id())
                                    _CL_obj.Args = MakeConstantList(n.Id(),MakeInteger(1).Id())
                                    v_bag_arg = _CL_obj.Id()
                                    } 
                                  va_arg2.AddFast(v_bag_arg)} 
                                va_arg1.Args = va_arg2
                                } 
                              arg_1 = _CL_obj
                              } 
                            { var arg_2 *Language.If
                              { var _CL_obj *Language.If = Language.To_If(new(Language.If).Is(Language.C_If))
                                { 
                                  var va_arg1 *Language.If
                                  var va_arg2 *ClaireAny
                                  va_arg1 = _CL_obj
                                  { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                    _CL_obj.Selector = ToProperty(f)
                                    { 
                                      var va_arg1 *Language.Call
                                      var va_arg2 *ClaireList
                                      va_arg1 = _CL_obj
                                      { 
                                        var v_bag_arg *ClaireAny
                                        va_arg2= ToType(CEMPTY.Id()).EmptyList()
                                        { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                          _CL_obj.Selector = C_nth
                                          _CL_obj.Args = MakeConstantList(l,m.Id())
                                          v_bag_arg = _CL_obj.Id()
                                          } 
                                        va_arg2.AddFast(v_bag_arg)
                                        va_arg2.AddFast(x.Id())} 
                                      va_arg1.Args = va_arg2
                                      } 
                                    va_arg2 = _CL_obj.Id()
                                    } 
                                  va_arg1.Test = va_arg2
                                  } 
                                { 
                                  var va_arg1 *Language.If
                                  var va_arg2 *ClaireAny
                                  va_arg1 = _CL_obj
                                  { var _CL_obj *Language.Do = Language.To_Do(new(Language.Do).Is(Language.C_Do))
                                    { 
                                      var va_arg1 *Language.Do
                                      var va_arg2 *ClaireList
                                      va_arg1 = _CL_obj
                                      { 
                                        var v_bag_arg *ClaireAny
                                        va_arg2= ToType(CEMPTY.Id()).EmptyList()
                                        { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                          _CL_obj.Selector = C_nth_equal
                                          { 
                                            var va_arg1 *Language.Call
                                            var va_arg2 *ClaireList
                                            va_arg1 = _CL_obj
                                            { 
                                              var v_bag_arg *ClaireAny
                                              va_arg2= ToType(CEMPTY.Id()).EmptyList()
                                              va_arg2.AddFast(l)
                                              va_arg2.AddFast(n.Id())
                                              { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                                _CL_obj.Selector = C_nth
                                                _CL_obj.Args = MakeConstantList(l,m.Id())
                                                v_bag_arg = _CL_obj.Id()
                                                } 
                                              va_arg2.AddFast(v_bag_arg)} 
                                            va_arg1.Args = va_arg2
                                            } 
                                          v_bag_arg = _CL_obj.Id()
                                          } 
                                        va_arg2.AddFast(v_bag_arg)
                                        { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                          _CL_obj.Selector = C_nth_equal
                                          _CL_obj.Args = MakeConstantList(l,m.Id(),x.Id())
                                          v_bag_arg = _CL_obj.Id()
                                          } 
                                        va_arg2.AddFast(v_bag_arg)} 
                                      va_arg1.Args = va_arg2
                                      } 
                                    va_arg2 = _CL_obj.Id()
                                    } 
                                  va_arg1.Arg = va_arg2
                                  } 
                                arg_2 = _CL_obj
                                } 
                              { var arg_3 *Language.Let
                                { var _CL_obj *Language.Let = Language.To_Let(new(Language.Let).Is(Language.C_Let))
                                  _CL_obj.ClaireVar = p
                                  { 
                                    var va_arg1 *Language.Let
                                    var va_arg2 *ClaireAny
                                    va_arg1 = _CL_obj
                                    { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                      _CL_obj.Selector = ToProperty(Core.C__sup_sup.Id())
                                      { 
                                        var va_arg1 *Language.Call
                                        var va_arg2 *ClaireList
                                        va_arg1 = _CL_obj
                                        { 
                                          var v_bag_arg *ClaireAny
                                          va_arg2= ToType(CEMPTY.Id()).EmptyList()
                                          { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                            _CL_obj.Selector = ToProperty(Core.C__plus.Id())
                                            _CL_obj.Args = MakeConstantList(n.Id(),m.Id())
                                            v_bag_arg = _CL_obj.Id()
                                            } 
                                          va_arg2.AddFast(v_bag_arg)
                                          va_arg2.AddFast(MakeInteger(1).Id())} 
                                        va_arg1.Args = va_arg2
                                        } 
                                      va_arg2 = _CL_obj.Id()
                                      } 
                                    va_arg1.Value = va_arg2
                                    } 
                                  { 
                                    var va_arg1 *Language.Let
                                    var va_arg2 *ClaireAny
                                    va_arg1 = _CL_obj
                                    { var _CL_obj *Language.Let = Language.To_Let(new(Language.Let).Is(Language.C_Let))
                                      _CL_obj.ClaireVar = q
                                      _CL_obj.Value = n.Id()
                                      { 
                                        var va_arg1 *Language.Let
                                        var va_arg2 *ClaireAny
                                        va_arg1 = _CL_obj
                                        { var _CL_obj *Language.Do = Language.To_Do(new(Language.Do).Is(Language.C_Do))
                                          { 
                                            var va_arg1 *Language.Do
                                            var va_arg2 *ClaireList
                                            va_arg1 = _CL_obj
                                            { 
                                              var v_bag_arg *ClaireAny
                                              va_arg2= ToType(CEMPTY.Id()).EmptyList()
                                              { var arg_4 *Language.Call
                                                { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                                  _CL_obj.Selector = C_nth
                                                  _CL_obj.Args = MakeConstantList(l,p.Id())
                                                  arg_4 = _CL_obj
                                                  } 
                                                v_bag_arg = Language.C_Assign.Make(x.Id(),arg_4.Id())
                                                } 
                                              va_arg2.AddFast(v_bag_arg)
                                              { var _CL_obj *Language.If = Language.To_If(new(Language.If).Is(Language.C_If))
                                                { 
                                                  var va_arg1 *Language.If
                                                  var va_arg2 *ClaireAny
                                                  va_arg1 = _CL_obj
                                                  { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                                    _CL_obj.Selector = ToProperty(Core.C__I_equal.Id())
                                                    _CL_obj.Args = MakeConstantList(p.Id(),n.Id())
                                                    va_arg2 = _CL_obj.Id()
                                                    } 
                                                  va_arg1.Test = va_arg2
                                                  } 
                                                { 
                                                  var va_arg1 *Language.If
                                                  var va_arg2 *ClaireAny
                                                  va_arg1 = _CL_obj
                                                  { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                                    _CL_obj.Selector = C_nth_equal
                                                    { 
                                                      var va_arg1 *Language.Call
                                                      var va_arg2 *ClaireList
                                                      va_arg1 = _CL_obj
                                                      { 
                                                        var v_bag_arg *ClaireAny
                                                        va_arg2= ToType(CEMPTY.Id()).EmptyList()
                                                        va_arg2.AddFast(l)
                                                        va_arg2.AddFast(p.Id())
                                                        { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                                          _CL_obj.Selector = C_nth
                                                          _CL_obj.Args = MakeConstantList(l,n.Id())
                                                          v_bag_arg = _CL_obj.Id()
                                                          } 
                                                        va_arg2.AddFast(v_bag_arg)} 
                                                      va_arg1.Args = va_arg2
                                                      } 
                                                    va_arg2 = _CL_obj.Id()
                                                    } 
                                                  va_arg1.Arg = va_arg2
                                                  } 
                                                v_bag_arg = _CL_obj.Id()
                                                } 
                                              va_arg2.AddFast(v_bag_arg)
                                              { var _CL_obj *Language.For = Language.To_For(new(Language.For).Is(Language.C_For))
                                                _CL_obj.ClaireVar = p
                                                { 
                                                  var va_arg1 *Language.Iteration
                                                  var va_arg2 *ClaireAny
                                                  va_arg1 = Language.To_Iteration(_CL_obj.Id())
                                                  { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                                    _CL_obj.Selector = ToProperty(C__dot_dot.Id())
                                                    { 
                                                      var va_arg1 *Language.Call
                                                      var va_arg2 *ClaireList
                                                      va_arg1 = _CL_obj
                                                      { 
                                                        var v_bag_arg *ClaireAny
                                                        va_arg2= ToType(CEMPTY.Id()).EmptyList()
                                                        { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                                          _CL_obj.Selector = ToProperty(Core.C__plus.Id())
                                                          _CL_obj.Args = MakeConstantList(n.Id(),MakeInteger(1).Id())
                                                          v_bag_arg = _CL_obj.Id()
                                                          } 
                                                        va_arg2.AddFast(v_bag_arg)
                                                        va_arg2.AddFast(m.Id())} 
                                                      va_arg1.Args = va_arg2
                                                      } 
                                                    va_arg2 = _CL_obj.Id()
                                                    } 
                                                  va_arg1.SetArg = va_arg2
                                                  } 
                                                { 
                                                  var va_arg1 *Language.Iteration
                                                  var va_arg2 *ClaireAny
                                                  va_arg1 = Language.To_Iteration(_CL_obj.Id())
                                                  { var _CL_obj *Language.If = Language.To_If(new(Language.If).Is(Language.C_If))
                                                    { 
                                                      var va_arg1 *Language.If
                                                      var va_arg2 *ClaireAny
                                                      va_arg1 = _CL_obj
                                                      { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                                        _CL_obj.Selector = ToProperty(f)
                                                        { 
                                                          var va_arg1 *Language.Call
                                                          var va_arg2 *ClaireList
                                                          va_arg1 = _CL_obj
                                                          { 
                                                            var v_bag_arg *ClaireAny
                                                            va_arg2= ToType(CEMPTY.Id()).EmptyList()
                                                            { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                                              _CL_obj.Selector = C_nth
                                                              _CL_obj.Args = MakeConstantList(l,p.Id())
                                                              v_bag_arg = _CL_obj.Id()
                                                              } 
                                                            va_arg2.AddFast(v_bag_arg)
                                                            va_arg2.AddFast(x.Id())} 
                                                          va_arg1.Args = va_arg2
                                                          } 
                                                        va_arg2 = _CL_obj.Id()
                                                        } 
                                                      va_arg1.Test = va_arg2
                                                      } 
                                                    { 
                                                      var va_arg1 *Language.If
                                                      var va_arg2 *ClaireAny
                                                      va_arg1 = _CL_obj
                                                      { var _CL_obj *Language.Do = Language.To_Do(new(Language.Do).Is(Language.C_Do))
                                                        { 
                                                          var va_arg1 *Language.Do
                                                          var va_arg2 *ClaireList
                                                          va_arg1 = _CL_obj
                                                          { 
                                                            var v_bag_arg *ClaireAny
                                                            va_arg2= ToType(CEMPTY.Id()).EmptyList()
                                                            { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                                              _CL_obj.Selector = C_nth_equal
                                                              { 
                                                                var va_arg1 *Language.Call
                                                                var va_arg2 *ClaireList
                                                                va_arg1 = _CL_obj
                                                                { 
                                                                  var v_bag_arg *ClaireAny
                                                                  va_arg2= ToType(CEMPTY.Id()).EmptyList()
                                                                  va_arg2.AddFast(l)
                                                                  va_arg2.AddFast(n.Id())
                                                                  { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                                                    _CL_obj.Selector = C_nth
                                                                    _CL_obj.Args = MakeConstantList(l,p.Id())
                                                                    v_bag_arg = _CL_obj.Id()
                                                                    } 
                                                                  va_arg2.AddFast(v_bag_arg)} 
                                                                va_arg1.Args = va_arg2
                                                                } 
                                                              v_bag_arg = _CL_obj.Id()
                                                              } 
                                                            va_arg2.AddFast(v_bag_arg)
                                                            { var arg_5 *Language.Call
                                                              { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                                                _CL_obj.Selector = ToProperty(Core.C__plus.Id())
                                                                _CL_obj.Args = MakeConstantList(n.Id(),MakeInteger(1).Id())
                                                                arg_5 = _CL_obj
                                                                } 
                                                              v_bag_arg = Language.C_Assign.Make(n.Id(),arg_5.Id())
                                                              } 
                                                            va_arg2.AddFast(v_bag_arg)
                                                            { var _CL_obj *Language.If = Language.To_If(new(Language.If).Is(Language.C_If))
                                                              { 
                                                                var va_arg1 *Language.If
                                                                var va_arg2 *ClaireAny
                                                                va_arg1 = _CL_obj
                                                                { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                                                  _CL_obj.Selector = ToProperty(C__sup.Id())
                                                                  _CL_obj.Args = MakeConstantList(p.Id(),n.Id())
                                                                  va_arg2 = _CL_obj.Id()
                                                                  } 
                                                                va_arg1.Test = va_arg2
                                                                } 
                                                              { 
                                                                var va_arg1 *Language.If
                                                                var va_arg2 *ClaireAny
                                                                va_arg1 = _CL_obj
                                                                { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                                                  _CL_obj.Selector = C_nth_equal
                                                                  { 
                                                                    var va_arg1 *Language.Call
                                                                    var va_arg2 *ClaireList
                                                                    va_arg1 = _CL_obj
                                                                    { 
                                                                      var v_bag_arg *ClaireAny
                                                                      va_arg2= ToType(CEMPTY.Id()).EmptyList()
                                                                      va_arg2.AddFast(l)
                                                                      va_arg2.AddFast(p.Id())
                                                                      { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                                                        _CL_obj.Selector = C_nth
                                                                        _CL_obj.Args = MakeConstantList(l,n.Id())
                                                                        v_bag_arg = _CL_obj.Id()
                                                                        } 
                                                                      va_arg2.AddFast(v_bag_arg)} 
                                                                    va_arg1.Args = va_arg2
                                                                    } 
                                                                  va_arg2 = _CL_obj.Id()
                                                                  } 
                                                                va_arg1.Arg = va_arg2
                                                                } 
                                                              v_bag_arg = _CL_obj.Id()
                                                              } 
                                                            va_arg2.AddFast(v_bag_arg)} 
                                                          va_arg1.Args = va_arg2
                                                          } 
                                                        va_arg2 = _CL_obj.Id()
                                                        } 
                                                      va_arg1.Arg = va_arg2
                                                      } 
                                                    va_arg2 = _CL_obj.Id()
                                                    } 
                                                  va_arg1.Arg = va_arg2
                                                  } 
                                                v_bag_arg = _CL_obj.Id()
                                                } 
                                              va_arg2.AddFast(v_bag_arg)
                                              { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                                _CL_obj.Selector = C_nth_equal
                                                _CL_obj.Args = MakeConstantList(l,n.Id(),x.Id())
                                                v_bag_arg = _CL_obj.Id()
                                                } 
                                              va_arg2.AddFast(v_bag_arg)
                                              { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                                _CL_obj.Selector = self.Arg.Selector
                                                { 
                                                  var va_arg1 *Language.Call
                                                  var va_arg2 *ClaireList
                                                  va_arg1 = _CL_obj
                                                  { 
                                                    var v_bag_arg *ClaireAny
                                                    va_arg2= ToType(CEMPTY.Id()).EmptyList()
                                                    va_arg2.AddFast(q.Id())
                                                    { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                                      _CL_obj.Selector = ToProperty(C__dash.Id())
                                                      _CL_obj.Args = MakeConstantList(n.Id(),MakeInteger(1).Id())
                                                      v_bag_arg = _CL_obj.Id()
                                                      } 
                                                    va_arg2.AddFast(v_bag_arg)
                                                    va_arg2.AddFast(l)} 
                                                  va_arg1.Args = va_arg2
                                                  } 
                                                v_bag_arg = _CL_obj.Id()
                                                } 
                                              va_arg2.AddFast(v_bag_arg)
                                              { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                                _CL_obj.Selector = self.Arg.Selector
                                                { 
                                                  var va_arg1 *Language.Call
                                                  var va_arg2 *ClaireList
                                                  va_arg1 = _CL_obj
                                                  { 
                                                    var v_bag_arg *ClaireAny
                                                    va_arg2= ToType(CEMPTY.Id()).EmptyList()
                                                    { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                                      _CL_obj.Selector = ToProperty(Core.C__plus.Id())
                                                      _CL_obj.Args = MakeConstantList(n.Id(),MakeInteger(1).Id())
                                                      v_bag_arg = _CL_obj.Id()
                                                      } 
                                                    va_arg2.AddFast(v_bag_arg)
                                                    va_arg2.AddFast(m.Id())
                                                    va_arg2.AddFast(l)} 
                                                  va_arg1.Args = va_arg2
                                                  } 
                                                v_bag_arg = _CL_obj.Id()
                                                } 
                                              va_arg2.AddFast(v_bag_arg)} 
                                            va_arg1.Args = va_arg2
                                            } 
                                          va_arg2 = _CL_obj.Id()
                                          } 
                                        va_arg1.Arg = va_arg2
                                        } 
                                      va_arg2 = _CL_obj.Id()
                                      } 
                                    va_arg1.Arg = va_arg2
                                    } 
                                  arg_3 = _CL_obj
                                  } 
                                va_arg2 = Language.C_If.Make(arg_1.Id(),arg_2.Id(),arg_3.Id())
                                } 
                              } 
                            } 
                          va_arg1.Arg = va_arg2
                          } 
                        va_arg2 = _CL_obj.Id()
                        } 
                      va_arg1.Arg = va_arg2
                      } 
                    _Zbd = _CL_obj
                    } 
                  { var def2 *Language.Defmethod
                    { var _CL_obj *Language.Defmethod = Language.To_Defmethod(new(Language.Defmethod).Is(Language.C_Defmethod))
                      { 
                        var va_arg1 *Language.Defmethod
                        var va_arg2 *Language.Call
                        va_arg1 = _CL_obj
                        { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                          _CL_obj.Selector = self.Arg.Selector
                          _CL_obj.Args = MakeConstantList(n.Id(),m.Id(),l)
                          va_arg2 = _CL_obj
                          } 
                        va_arg1.Arg = va_arg2
                        } 
                      _CL_obj.Inline_ask = CFALSE
                      _CL_obj.SetArg = self.SetArg
                      _CL_obj.Body = _Zbd.Id()
                      def2 = _CL_obj
                      } 
                    Core.F_tformat_string(MakeString("---- note: quick sort optimisation for ~S ---- \n"),2,MakeConstantList(self.Arg.Selector.Id()))
                    Result = EVAL(def2.Id())
                    if !ErrorIn(Result) {
                    { var _CL_obj *Language.Do = Language.To_Do(new(Language.Do).Is(Language.C_Do))
                      { 
                        var va_arg1 *Language.Do
                        var va_arg2 *ClaireList
                        va_arg1 = _CL_obj
                        var try_6 EID
                        { 
                          var v_bag_arg *ClaireAny
                          try_6= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                          var try_7 EID
                          try_7 = Core.F_CALL(C_c_code,ARGS(EID{def1.Id(),0}))
                          if ErrorIn(try_7) {try_6 = try_7
                          } else {
                          v_bag_arg = ANY(try_7)
                          ToList(OBJ(try_6)).AddFast(v_bag_arg)
                          var try_8 EID
                          try_8 = Core.F_CALL(C_c_code,ARGS(EID{def2.Id(),0}))
                          if ErrorIn(try_8) {try_6 = try_8
                          } else {
                          v_bag_arg = ANY(try_8)
                          ToList(OBJ(try_6)).AddFast(v_bag_arg)}}
                          } 
                        if ErrorIn(try_6) {Result = try_6
                        } else {
                        va_arg2 = ToList(OBJ(try_6))
                        va_arg1.Args = va_arg2
                        Result = EID{va_arg2.Id(),0}
                        }
                        } 
                      if !ErrorIn(Result) {
                      Result = EID{_CL_obj.Id(),0}
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
  } 
return Result} 

// The EID go function for: sort_code @ Defmethod (throw: true) 
func E_Optimize_sort_code_Defmethod (self EID,lv EID) EID { 
return F_Optimize_sort_code_Defmethod(Language.To_Defmethod(OBJ(self)),ToList(OBJ(lv)) )} 

// new: we deal with floats --------------------------------------
// create a restriction so that OPT is happy
// the last argument is the method that is passed to the code generator
/* The go function for: add_method(p:property,ls:list,rg:type,st:integer,f1:function,m:method) [status=0] */
func F_add_method_property2 (p *ClaireProperty,ls *ClaireList,rg *ClaireType,st int,f1 *ClaireFunction,m *ClaireMethod) *ClaireMethod { 
return  F_add_method_property(p,ls,rg,st,f1)
} 

// The EID go function for: add_method @ list<type_expression>(property, list, type, integer, function, method) (throw: false) 
func E_add_method_property2 (p EID,ls EID,rg EID,st EID,f1 EID,m EID) EID { 
return EID{F_add_method_property2(ToProperty(OBJ(p)),
  ToList(OBJ(ls)),
  ToType(OBJ(rg)),
  INT(st),
  ToFunction(OBJ(f1)),
  ToMethod(OBJ(m)) ).Id(),0}} 

/* The go function for: add_method!(m:method,ls:list,rg:any,stat:any,fu:function) [status=1] */
func F_Optimize_add_method_I_method (m *ClaireMethod,ls *ClaireList,rg *ClaireAny,stat *ClaireAny,fu *ClaireFunction) EID { 
var Result EID
{ 
  var _Zc EID
  { var _CL_obj *Language.CallMethod = Language.To_CallMethod(new(Language.CallMethod).Is(Language.C_Call_method))
    _CL_obj.Arg = ToMethod(C_Optimize__staradd_method2_star.Value)
    { 
      var va_arg1 *Language.CallMethod
      var va_arg2 *ClaireList
      va_arg1 = _CL_obj
      var try_1 EID
      { 
        var v_bag_arg *ClaireAny
        try_1= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
        var try_2 EID
        try_2 = Core.F_CALL(C_c_code,ARGS(EID{m.Selector.Id(),0},EID{C_property.Id(),0}))
        if ErrorIn(try_2) {try_1 = try_2
        } else {
        v_bag_arg = ANY(try_2)
        ToList(OBJ(try_1)).AddFast(v_bag_arg)
        var try_3 EID
        try_3 = Core.F_CALL(C_c_code,ARGS(EID{ls.Id(),0},EID{C_list.Id(),0}))
        if ErrorIn(try_3) {try_1 = try_3
        } else {
        v_bag_arg = ANY(try_3)
        ToList(OBJ(try_1)).AddFast(v_bag_arg)
        var try_4 EID
        try_4 = Core.F_CALL(C_c_code,ARGS(rg.ToEID(),EID{C_type.Id(),0}))
        if ErrorIn(try_4) {try_1 = try_4
        } else {
        v_bag_arg = ANY(try_4)
        ToList(OBJ(try_1)).AddFast(v_bag_arg)
        ToList(OBJ(try_1)).AddFast(stat)
        ToList(OBJ(try_1)).AddFast(fu.Id())
        ToList(OBJ(try_1)).AddFast(m.Id())}}}
        } 
      if ErrorIn(try_1) {_Zc = try_1
      } else {
      va_arg2 = ToList(OBJ(try_1))
      va_arg1.Args = va_arg2
      _Zc = EID{va_arg2.Id(),0}
      }
      } 
    if !ErrorIn(_Zc) {
    _Zc = EID{_CL_obj.Id(),0}
    }
    } 
  if ErrorIn(_Zc) {Result = _Zc
  } else {
  
  Result = _Zc}
  } 
return Result} 

// The EID go function for: add_method! @ method (throw: true) 
func E_Optimize_add_method_I_method (m EID,ls EID,rg EID,stat EID,fu EID) EID { 
return F_Optimize_add_method_I_method(ToMethod(OBJ(m)),
  ToList(OBJ(ls)),
  ANY(rg),
  ANY(stat),
  ToFunction(OBJ(fu)) )} 

// this signature extraction is more subtle since it also builds an external
// list. (l1 is the domain (may use global variables), l2 is the "pure"
// list of patterns)
/* The go function for: extract_signature!(l:list) [status=1] */
func F_Optimize_extract_signature_I_list (l *ClaireList) EID { 
var Result EID
Language.C_LDEF.Value = ToType(C_any.Id()).EmptyList().Id()
{ var n int = 0
  { var l1 *ClaireList = ToType(C_type_expression.Id()).EmptyList()
    { var l2 *ClaireList
      var try_1 EID
      { 
        var v_list3 *ClaireList
        var v *ClaireVariable
        var v_local3 *ClaireAny
        v_list3 = l
        try_1 = EID{CreateList(ToType(C_any.Id()),v_list3.Length()).Id(),0}
        for CLcount := 0; CLcount < v_list3.Length(); CLcount++{ 
          v = To_Variable(v_list3.At(CLcount))
          var try_2 EID
          { var p *ClaireAny
            var try_3 EID
            try_3 = Language.F_extract_pattern_any(v.Range.Id(),MakeConstantList(MakeInteger(n).Id()))
            if ErrorIn(try_3) {try_2 = try_3
            } else {
            p = ANY(try_3)
            n = (n+1)
            { var arg_4 *ClaireAny
              if (v.Range.Isa.IsIn(Core.C_global_variable) == CTRUE) { 
                arg_4 = v.Range.Id()
                } else {
                arg_4 = p
                } 
              l1 = l1.AddFast(arg_4)
              } 
            v.Range = Language.F_type_I_any(p)
            try_2 = p.ToEID()
            }
            } 
          if ErrorIn(try_2) {try_1 = try_2
          break
          } else {
          v_local3 = ANY(try_2)
          ToList(OBJ(try_1)).PutAt(CLcount,v_local3)
          } 
        }
        } 
      if ErrorIn(try_1) {Result = try_1
      } else {
      l2 = ToList(OBJ(try_1))
      Result = EID{MakeConstantList(l1.Id(),l2.Id()).Id(),0}
      }
      } 
    } 
  } 
return Result} 

// The EID go function for: extract_signature! @ list (throw: true) 
func E_Optimize_extract_signature_I_list (l EID) EID { 
return F_Optimize_extract_signature_I_list(ToList(OBJ(l)) )} 

// check signature equality
/* The go function for: =sig?(x:list,y:list) [status=0] */
func F_Optimize__equalsig_ask_list (x *ClaireList,y *ClaireList) *ClaireBoolean { 
return  MakeBoolean((Core.F_tmatch_ask_list(x,y) == CTRUE) && (Core.F_tmatch_ask_list(y,x) == CTRUE))
} 

// The EID go function for: =sig? @ list (throw: false) 
func E_Optimize__equalsig_ask_list (x EID,y EID) EID { 
return EID{F_Optimize__equalsig_ask_list(ToList(OBJ(x)),ToList(OBJ(y)) ).Id(),0}} 

// creates a name for a restriction from the full domain
// Note that we suppose that a new restriction is not allowed to be inserted
// in a list of restrictions when the property is closed.
//
/* The go function for: Compile/function_name(p:property,l:list,x:any) [status=0] */
func F_Compile_function_name_property1 (p *ClaireProperty,l *ClaireList,x *ClaireAny) *ClaireString { 
var Result *ClaireString
if (C_function.Id() == x.Isa.Id()) { 
  Result = F_string_I_function(ToFunction(x))
  } else {
  { var n int = 0
    { var m int = 0
      { var md *ClaireModule = p.Name.Module_I()
        { var c *ClaireClass = ToTypeExpression(l.At(0)).Class_I()
          { var r *ClaireString = F_append_string(F_append_string(p.Name.String_I(),MakeString("_")),c.Name.String_I())
            if (p.Id() != Core.C_main.Id()) { 
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
                if (F_Optimize__equalsig_ask_list(l,r.Domain) == CTRUE) { 
                  m = n
                  } 
                } 
              } 
            if (n <= 1) { 
              r = r
              } else {
              r = F_append_string(r,F_string_I_integer(m))
              } 
            if ((F_Optimize_stable_ask_relation(ToRelation(p.Id())) == CTRUE) || 
                (p.Id() == Core.C_main.Id())) { 
              Result = r
              } else {
              Result = F_append_string(F_append_string(r,MakeString("_")),ClEnv.Module_I.Name.String_I())
              } 
            } 
          } 
        } 
      } 
    } 
  } 
return Result} 

// The EID go function for: Compile/function_name @ list<type_expression>(property, list, any) (throw: false) 
func E_Compile_function_name_property1 (p EID,l EID,x EID) EID { 
return EID{F_Compile_function_name_property1(ToProperty(OBJ(p)),ToList(OBJ(l)),ANY(x) ).Id(),0}} 

// this compiles a lambda into a C method with name oself.
// the use_new flag will be raised if a new object is created inside the
// function.
// m is either the associated method,or the expected range
//
/* The go function for: Compile/compile_lambda(self:string,l:lambda,m:any) [status=1] */
func F_Compile_compile_lambda_string (self *ClaireString,l *ClaireLambda,m *ClaireAny) EID { 
var Result EID
{ var x int = C_compiler.Safety
  { var y *ClaireLambda = l
    _ = y
    Core.F_tformat_string(MakeString("---- Compiling ~A,\n"),3,MakeConstantList((self).Id()))
    if (C_method.Id() == m.Isa.Id()) { 
      { var g0317 *ClaireMethod = ToMethod(m)
        C_OPT.InMethod = g0317.Id()
        } 
      } 
    if (C_OPT.LoopIndex > 0) { 
      C_OPT.LoopIndex = 0
      } 
    C_OPT.MaxVars = 0
    if (ToBoolean(C_OPT.Unsure.Contain_ask(m).Id()) == CTRUE) { 
      C_compiler.Safety = 1
      } 
    Result = Core.F_CALL(C_Compile_make_c_function,ARGS(EID{l.Id(),0},EID{(self).Id(),0},m.ToEID()))
    if !ErrorIn(Result) {
    C_OPT.InMethod = CNULL
    C_compiler.Safety = x
    Result = EID{CTRUE.Id(),0}
    }
    } 
  } 
return Result} 

// The EID go function for: Compile/compile_lambda @ string (throw: true) 
func E_Compile_compile_lambda_string (self EID,l EID,m EID) EID { 
return F_Compile_compile_lambda_string(ToString(OBJ(self)),ToLambda(OBJ(l)),ANY(m) )} 

// how to compile an table definition
/* The go function for: c_code(self:Defarray) [status=1] */
func F_c_code_Defarray (self *Language.Defarray) EID { 
var Result EID
{ var a *ClaireList = self.Arg.Args
  { var _Za *ClaireAny
    var try_1 EID
    { var arg_2 *ClaireSymbol
      var try_3 EID
      try_3 = Language.F_extract_symbol_any(a.At(0))
      if ErrorIn(try_3) {try_1 = try_3
      } else {
      arg_2 = ToSymbol(OBJ(try_3))
      try_1 = arg_2.Value().ToEID()
      }
      } 
    if ErrorIn(try_1) {Result = try_1
    } else {
    _Za = ANY(try_1)
    { var _Zv *ClaireTable
      var try_4 EID
      if (C_table.Id() == _Za.Isa.Id()) { 
        { var g0319 *ClaireTable = ToTable(_Za)
          try_4 = EID{g0319.Id(),0}
          } 
        } else {
        try_4 = ToException(Core.C_general_error.Make(MakeString("[internal] the table ~S is unknown").Id(),MakeConstantList(a.At(0)).Id())).Close()
        } 
      if ErrorIn(try_4) {Result = try_4
      } else {
      _Zv = ToTable(OBJ(try_4))
      { var s *ClaireType = ToType(OBJ(Core.F_CALL(C_domain,ARGS(_Za.ToEID()))))
        { var e *ClaireAny
          var try_5 EID
          { var l *ClaireList
            var try_6 EID
            try_6 = a.Cdr()
            if ErrorIn(try_6) {try_5 = try_6
            } else {
            l = ToList(OBJ(try_6))
            { var b *ClaireAny
              var try_7 EID
              try_7 = Language.F_iClaire_lexical_index_any2(self.Body,l,0,CTRUE)
              if ErrorIn(try_7) {try_5 = try_7
              } else {
              b = ANY(try_7)
              var g0325I *ClaireBoolean
              { var arg_8 *ClaireAny
                { 
                  var va *ClaireAny
                  _ = va
                  arg_8= CFALSE.Id()
                  var va_support *ClaireList
                  va_support = l
                  va_len := va_support.Length()
                  for i_it := 0; i_it < va_len; i_it++ { 
                    va = va_support.At(i_it)
                    if (Language.F_occurrence_any(b,To_Variable(va)) > 0) { 
                      arg_8 = CTRUE.Id()
                      break
                      } 
                    } 
                  } 
                g0325I = F_boolean_I_any(arg_8)
                } 
              if (g0325I == CTRUE) { 
                try_5 = Language.F_lambda_I_list(l,b)
                } else {
                try_5 = self.Body.ToEID()
                } 
              }
              } 
            }
            } 
          if ErrorIn(try_5) {Result = try_5
          } else {
          e = ANY(try_5)
          { var d *ClaireAny
            _ = d
            if (e.Isa.IsIn(C_lambda) == CTRUE) { 
              d = CNULL
              } else {
              d = self.Body
              } 
            { var _Zl1 *ClaireList
              if (ToRelation(_Za).Multivalued_ask == CTRUE) { 
                { 
                  var v_bag_arg *ClaireAny
                  _Zl1= ToType(C_any.Id()).EmptyList()
                  { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                    _CL_obj.Selector = C_put
                    _CL_obj.Args = MakeConstantList(C_multivalued_ask.Id(),_Zv.Id(),ToRelation(_Za).Multivalued_ask.Id())
                    v_bag_arg = _CL_obj.Id()
                    } 
                  _Zl1.AddFast(v_bag_arg)} 
                } else {
                _Zl1 = ToType(C_any.Id()).EmptyList()
                } 
              { var _Zl2 *ClaireList
                { 
                  var v_bag_arg *ClaireAny
                  _Zl2= ToType(C_any.Id()).EmptyList()
                  { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                    _CL_obj.Selector = C_put
                    _CL_obj.Args = MakeConstantList(C_range.Id(),_Zv.Id(),ANY(Core.F_CALL(C_range,ARGS(_Za.ToEID()))))
                    v_bag_arg = _CL_obj.Id()
                    } 
                  _Zl2.AddFast(v_bag_arg)
                  { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                    _CL_obj.Selector = C_put
                    _CL_obj.Args = MakeConstantList(C_params.Id(),_Zv.Id(),ANY(Core.F_CALL(C_params,ARGS(_Za.ToEID()))))
                    v_bag_arg = _CL_obj.Id()
                    } 
                  _Zl2.AddFast(v_bag_arg)
                  { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                    _CL_obj.Selector = C_put
                    _CL_obj.Args = MakeConstantList(C_domain.Id(),_Zv.Id(),s.Id())
                    v_bag_arg = _CL_obj.Id()
                    } 
                  _Zl2.AddFast(v_bag_arg)} 
                { 
                  var va_arg1 *ClaireVariable
                  var va_arg2 *ClaireType
                  va_arg1 = To_Variable(a.At(1))
                  var try_9 EID
                  try_9 = Language.F_extract_type_any(To_Variable(a.At(1)).Range.Id())
                  if ErrorIn(try_9) {Result = try_9
                  } else {
                  va_arg2 = ToType(OBJ(try_9))
                  va_arg1.Range = va_arg2
                  Result = EID{va_arg2.Id(),0}
                  }
                  } 
                if !ErrorIn(Result) {
                if (a.Length() == 2) { 
                  var try_10 EID
                  { var arg_11 *Language.Call
                    var try_12 EID
                    if (s.Isa.IsIn(C_Interval) == CTRUE) { 
                      { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                        _CL_obj.Selector = C_put
                        { 
                          var va_arg1 *Language.Call
                          var va_arg2 *ClaireList
                          va_arg1 = _CL_obj
                          var try_13 EID
                          { 
                            var v_bag_arg *ClaireAny
                            try_13= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                            ToList(OBJ(try_13)).AddFast(C_mClaire_graph.Id())
                            ToList(OBJ(try_13)).AddFast(_Zv.Id())
                            var try_14 EID
                            { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                              _CL_obj.Selector = Core.C_typed_copy_list
                              { 
                                var va_arg1 *Language.Call
                                var va_arg2 *ClaireList
                                va_arg1 = _CL_obj
                                var try_15 EID
                                { 
                                  var v_bag_arg *ClaireAny
                                  try_15= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                                  ToList(OBJ(try_15)).AddFast(ANY(Core.F_CALL(C_range,ARGS(_Za.ToEID()))))
                                  var try_16 EID
                                  try_16 = Core.F_CALL(C_size,ARGS(EID{s.Id(),0}))
                                  if ErrorIn(try_16) {try_15 = try_16
                                  } else {
                                  v_bag_arg = ANY(try_16)
                                  ToList(OBJ(try_15)).AddFast(v_bag_arg)
                                  ToList(OBJ(try_15)).AddFast(ANY(Core.F_CALL(C_default,ARGS(_Za.ToEID()))))}
                                  } 
                                if ErrorIn(try_15) {try_14 = try_15
                                } else {
                                va_arg2 = ToList(OBJ(try_15))
                                va_arg1.Args = va_arg2
                                try_14 = EID{va_arg2.Id(),0}
                                }
                                } 
                              if !ErrorIn(try_14) {
                              try_14 = EID{_CL_obj.Id(),0}
                              }
                              } 
                            if ErrorIn(try_14) {try_13 = try_14
                            } else {
                            v_bag_arg = ANY(try_14)
                            ToList(OBJ(try_13)).AddFast(v_bag_arg)}
                            } 
                          if ErrorIn(try_13) {try_12 = try_13
                          } else {
                          va_arg2 = ToList(OBJ(try_13))
                          va_arg1.Args = va_arg2
                          try_12 = EID{va_arg2.Id(),0}
                          }
                          } 
                        if !ErrorIn(try_12) {
                        try_12 = EID{_CL_obj.Id(),0}
                        }
                        } 
                      } else {
                      { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                        _CL_obj.Selector = C_graph_init
                        _CL_obj.Args = MakeConstantList(_Zv.Id())
                        try_12 = EID{_CL_obj.Id(),0}
                        } 
                      } 
                    if ErrorIn(try_12) {try_10 = try_12
                    } else {
                    arg_11 = Language.To_Call(OBJ(try_12))
                    try_10 = EID{_Zl2.AddFast(arg_11.Id()).Id(),0}
                    }
                    } 
                  if ErrorIn(try_10) {Result = try_10
                  } else {
                  _Zl2 = ToList(OBJ(try_10))
                  Result = EID{_Zl2.Id(),0}
                  { var arg_17 *ClaireObject
                    if (e.Isa.IsIn(C_lambda) == CTRUE) { 
                      { var g0323 *ClaireLambda = ToLambda(e)
                        { var _CL_obj *Language.For = Language.To_For(new(Language.For).Is(Language.C_For))
                          _CL_obj.ClaireVar = To_Variable(a.At(1))
                          _CL_obj.SetArg = s.Id()
                          { 
                            var va_arg1 *Language.Iteration
                            var va_arg2 *ClaireAny
                            va_arg1 = Language.To_Iteration(_CL_obj.Id())
                            { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                              _CL_obj.Selector = C_nth_equal
                              _CL_obj.Args = MakeConstantList(_Zv.Id(),a.At(1),g0323.Body)
                              va_arg2 = _CL_obj.Id()
                              } 
                            va_arg1.Arg = va_arg2
                            } 
                          arg_17 = ToObject(_CL_obj.Id())
                          } 
                        } 
                      } else {
                      arg_17 = ToObject(CFALSE.Id())
                      } 
                    _Zl2 = _Zl2.AddFast(arg_17.Id())
                    } 
                  Result = EID{_Zl2.Id(),0}
                  }
                  } else {
                  { var s1 *ClaireTypeExpression
                    var try_18 EID
                    try_18 = Language.F_extract_type_any(To_Variable(a.At(1)).Range.Id())
                    if ErrorIn(try_18) {Result = try_18
                    } else {
                    s1 = ToTypeExpression(OBJ(try_18))
                    { var s2 *ClaireTypeExpression
                      var try_19 EID
                      try_19 = Language.F_extract_type_any(To_Variable(a.At(2)).Range.Id())
                      if ErrorIn(try_19) {Result = try_19
                      } else {
                      s2 = ToTypeExpression(OBJ(try_19))
                      To_Variable(a.At(2)).Range = ToType(s2.Id())
                      
                      { var arg_20 *Language.Call
                        if ((s1.Isa.IsIn(C_Interval) == CTRUE) && 
                            (s2.Isa.IsIn(C_Interval) == CTRUE)) { 
                          { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                            _CL_obj.Selector = C_put
                            { 
                              var va_arg1 *Language.Call
                              var va_arg2 *ClaireList
                              va_arg1 = _CL_obj
                              { 
                                var v_bag_arg *ClaireAny
                                va_arg2= ToType(CEMPTY.Id()).EmptyList()
                                va_arg2.AddFast(C_mClaire_graph.Id())
                                va_arg2.AddFast(_Zv.Id())
                                { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                  _CL_obj.Selector = Core.C_typed_copy_list
                                  _CL_obj.Args = MakeConstantList(ANY(Core.F_CALL(C_range,ARGS(_Za.ToEID()))),ANY(Core.F_CALL(C_length,ARGS(ToTable(_Za).Graph.ToEID()))),ANY(Core.F_CALL(C_default,ARGS(_Za.ToEID()))))
                                  v_bag_arg = _CL_obj.Id()
                                  } 
                                va_arg2.AddFast(v_bag_arg)} 
                              va_arg1.Args = va_arg2
                              } 
                            arg_20 = _CL_obj
                            } 
                          } else {
                          { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                            _CL_obj.Selector = C_graph_init
                            _CL_obj.Args = MakeConstantList(_Zv.Id())
                            arg_20 = _CL_obj
                            } 
                          } 
                        _Zl2 = _Zl2.AddFast(arg_20.Id())
                        } 
                      { var arg_21 *ClaireObject
                        if (e.Isa.IsIn(C_lambda) == CTRUE) { 
                          { var g0324 *ClaireLambda = ToLambda(e)
                            { var _CL_obj *Language.For = Language.To_For(new(Language.For).Is(Language.C_For))
                              _CL_obj.ClaireVar = To_Variable(a.At(1))
                              _CL_obj.SetArg = s1.Id()
                              { 
                                var va_arg1 *Language.Iteration
                                var va_arg2 *ClaireAny
                                va_arg1 = Language.To_Iteration(_CL_obj.Id())
                                { var _CL_obj *Language.For = Language.To_For(new(Language.For).Is(Language.C_For))
                                  _CL_obj.ClaireVar = To_Variable(a.At(2))
                                  _CL_obj.SetArg = s2.Id()
                                  { 
                                    var va_arg1 *Language.Iteration
                                    var va_arg2 *ClaireAny
                                    va_arg1 = Language.To_Iteration(_CL_obj.Id())
                                    { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                      _CL_obj.Selector = C_nth_equal
                                      _CL_obj.Args = MakeConstantList(_Zv.Id(),
                                        a.At(1),
                                        a.At(2),
                                        g0324.Body)
                                      va_arg2 = _CL_obj.Id()
                                      } 
                                    va_arg1.Arg = va_arg2
                                    } 
                                  va_arg2 = _CL_obj.Id()
                                  } 
                                va_arg1.Arg = va_arg2
                                } 
                              arg_21 = ToObject(_CL_obj.Id())
                              } 
                            } 
                          } else {
                          arg_21 = ToObject(CFALSE.Id())
                          } 
                        _Zl2 = _Zl2.AddFast(arg_21.Id())
                        } 
                      Result = EID{_Zl2.Id(),0}
                      }
                      } 
                    }
                    } 
                  } 
                if !ErrorIn(Result) {
                { var arg_22 *Language.Call
                  { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                    _CL_obj.Selector = C_put
                    _CL_obj.Args = MakeConstantList(C_default.Id(),_Zv.Id(),ANY(Core.F_CALL(C_default,ARGS(_Za.ToEID()))))
                    arg_22 = _CL_obj
                    } 
                  _Zl2 = _Zl2.AddFast(arg_22.Id())
                  } 
                C_OPT.Objects = C_OPT.Objects.AddFast(_Za)
                Core.F_CALL(C_Optimize_c_register,ARGS(_Za.ToEID()))
                { var arg_23 *Language.Do
                  var try_24 EID
                  { var _CL_obj *Language.Do = Language.To_Do(new(Language.Do).Is(Language.C_Do))
                    { 
                      var va_arg1 *Language.Do
                      var va_arg2 *ClaireList
                      va_arg1 = _CL_obj
                      var try_25 EID
                      { var arg_26 *Language.Call
                        { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                          _CL_obj.Selector = C_Compile_object_I
                          _CL_obj.Args = MakeConstantList(ANY(Core.F_CALL(C_name,ARGS(_Za.ToEID()))),C_table.Id())
                          arg_26 = _CL_obj
                          } 
                        { var arg_27 *ClaireList
                          var try_28 EID
                          try_28 = _Zl1.Add_star(_Zl2)
                          if ErrorIn(try_28) {try_25 = try_28
                          } else {
                          arg_27 = ToList(OBJ(try_28))
                          try_25 = EID{F_cons_any(arg_26.Id(),arg_27).Id(),0}
                          }
                          } 
                        } 
                      if ErrorIn(try_25) {try_24 = try_25
                      } else {
                      va_arg2 = ToList(OBJ(try_25))
                      va_arg1.Args = va_arg2
                      try_24 = EID{va_arg2.Id(),0}
                      }
                      } 
                    if !ErrorIn(try_24) {
                    try_24 = EID{_CL_obj.Id(),0}
                    }
                    } 
                  if ErrorIn(try_24) {Result = try_24
                  } else {
                  arg_23 = Language.To_Do(OBJ(try_24))
                  Result = Core.F_CALL(C_c_code,ARGS(EID{arg_23.Id(),0},EID{C_any.Id(),0}))
                  }
                  } 
                }}
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
return Result} 

// The EID go function for: c_code @ Defarray (throw: true) 
func E_c_code_Defarray (self EID) EID { 
return F_c_code_Defarray(Language.To_Defarray(OBJ(self)) )} 

// *********************************************************************
// *     Part 4: Inverse Management (new in v3.0.50)                   *
// *********************************************************************
// this method creates an if_write demon that takes care of the inverse
/* The go function for: Compile/compute_if_write_inverse(R:relation) [status=1] */
func F_Compile_compute_if_write_inverse_relation (R *ClaireRelation) EID { 
var Result EID
{ var x *ClaireVariable
  { var _CL_obj *ClaireVariable = To_Variable(new(ClaireVariable).Is(C_Variable))
    _CL_obj.Pname = Core.F_symbol_I_string2(MakeString("XX"))
    _CL_obj.Range = R.Domain
    x = _CL_obj
    } 
  { var y *ClaireVariable
    { var _CL_obj *ClaireVariable = To_Variable(new(ClaireVariable).Is(C_Variable))
      _CL_obj.Pname = Core.F_symbol_I_string2(MakeString("YY"))
      { 
        var va_arg1 *ClaireVariable
        var va_arg2 *ClaireType
        va_arg1 = _CL_obj
        if (R.Multivalued_ask == CTRUE) { 
          va_arg2 = Core.F_member_type(R.Range)
          } else {
          va_arg2 = R.Range
          } 
        va_arg1.Range = va_arg2
        } 
      y = _CL_obj
      } 
    { var z *ClaireVariable
      { var _CL_obj *ClaireVariable = To_Variable(new(ClaireVariable).Is(C_Variable))
        _CL_obj.Pname = Core.F_symbol_I_string2(MakeString("ZZ"))
        _CL_obj.Range = R.Range
        z = _CL_obj
        } 
      { var l1 *ClaireList = ToType(C_any.Id()).EmptyList()
        if (R.Multivalued_ask == CTRUE) { 
          var try_1 EID
          { 
            var v_bag_arg *ClaireAny
            try_1= EID{ToType(C_any.Id()).EmptyList().Id(),0}
            var try_2 EID
            try_2 = Core.F_CALL(C_Optimize_Produce_put,ARGS(EID{R.Id(),0},EID{x.Id(),0},EID{y.Id(),0}))
            if ErrorIn(try_2) {try_1 = try_2
            } else {
            v_bag_arg = ANY(try_2)
            ToList(OBJ(try_1)).AddFast(v_bag_arg)}
            } 
          if ErrorIn(try_1) {Result = try_1
          } else {
          l1 = ToList(OBJ(try_1))
          Result = EID{l1.Id(),0}
          if (R.Inverse.Id() != CNULL) { 
            var try_3 EID
            { var arg_4 *ClaireAny
              var try_5 EID
              try_5 = Core.F_CALL(C_Optimize_Produce_put,ARGS(EID{R.Inverse.Id(),0},EID{y.Id(),0},EID{x.Id(),0}))
              if ErrorIn(try_5) {try_3 = try_5
              } else {
              arg_4 = ANY(try_5)
              try_3 = EID{l1.AddFast(arg_4).Id(),0}
              }
              } 
            if ErrorIn(try_3) {Result = try_3
            } else {
            l1 = ToList(OBJ(try_3))
            Result = EID{l1.Id(),0}
            }
            } else {
            Result = EID{CFALSE.Id(),0}
            } 
          if !ErrorIn(Result) {
          { 
            var va_arg1 *ClaireRelation
            var va_arg2 *ClaireAny
            va_arg1 = R
            var try_6 EID
            { var arg_7 *Language.If
              var try_8 EID
              { var _CL_obj *Language.If = Language.To_If(new(Language.If).Is(Language.C_If))
                { 
                  var va_arg1 *Language.If
                  var va_arg2 *ClaireAny
                  va_arg1 = _CL_obj
                  var try_9 EID
                  { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                    _CL_obj.Selector = Core.C_not
                    { 
                      var va_arg1 *Language.Call
                      var va_arg2 *ClaireList
                      va_arg1 = _CL_obj
                      var try_10 EID
                      { 
                        var v_bag_arg *ClaireAny
                        try_10= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                        var try_11 EID
                        { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                          _CL_obj.Selector = ToProperty(C__Z.Id())
                          { 
                            var va_arg1 *Language.Call
                            var va_arg2 *ClaireList
                            va_arg1 = _CL_obj
                            var try_12 EID
                            { 
                              var v_bag_arg *ClaireAny
                              try_12= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                              ToList(OBJ(try_12)).AddFast(y.Id())
                              var try_13 EID
                              try_13 = F_Optimize_Produce_get_relation(R,x)
                              if ErrorIn(try_13) {try_12 = try_13
                              } else {
                              v_bag_arg = ANY(try_13)
                              ToList(OBJ(try_12)).AddFast(v_bag_arg)}
                              } 
                            if ErrorIn(try_12) {try_11 = try_12
                            } else {
                            va_arg2 = ToList(OBJ(try_12))
                            va_arg1.Args = va_arg2
                            try_11 = EID{va_arg2.Id(),0}
                            }
                            } 
                          if !ErrorIn(try_11) {
                          try_11 = EID{_CL_obj.Id(),0}
                          }
                          } 
                        if ErrorIn(try_11) {try_10 = try_11
                        } else {
                        v_bag_arg = ANY(try_11)
                        ToList(OBJ(try_10)).AddFast(v_bag_arg)}
                        } 
                      if ErrorIn(try_10) {try_9 = try_10
                      } else {
                      va_arg2 = ToList(OBJ(try_10))
                      va_arg1.Args = va_arg2
                      try_9 = EID{va_arg2.Id(),0}
                      }
                      } 
                    if !ErrorIn(try_9) {
                    try_9 = EID{_CL_obj.Id(),0}
                    }
                    } 
                  if ErrorIn(try_9) {try_8 = try_9
                  } else {
                  va_arg2 = ANY(try_9)
                  va_arg1.Test = va_arg2
                  try_8 = va_arg2.ToEID()
                  }
                  } 
                if !ErrorIn(try_8) {
                _CL_obj.Arg = F_Compile_Do_I_list(l1)
                try_8 = EID{_CL_obj.Id(),0}
                }
                } 
              if ErrorIn(try_8) {try_6 = try_8
              } else {
              arg_7 = Language.To_If(OBJ(try_8))
              try_6 = Language.F_lambda_I_list(MakeConstantList(x.Id(),y.Id()),arg_7.Id())
              }
              } 
            if ErrorIn(try_6) {Result = try_6
            } else {
            va_arg2 = ANY(try_6)
            va_arg1.IfWrite = va_arg2
            Result = va_arg2.ToEID()
            }
            } 
          }}
          } else {
          var try_14 EID
          { 
            var v_bag_arg *ClaireAny
            try_14= EID{ToType(C_any.Id()).EmptyList().Id(),0}
            var try_15 EID
            try_15 = Core.F_CALL(C_Optimize_Produce_put,ARGS(EID{R.Id(),0},EID{x.Id(),0},EID{y.Id(),0}))
            if ErrorIn(try_15) {try_14 = try_15
            } else {
            v_bag_arg = ANY(try_15)
            ToList(OBJ(try_14)).AddFast(v_bag_arg)}
            } 
          if ErrorIn(try_14) {Result = try_14
          } else {
          l1 = ToList(OBJ(try_14))
          Result = EID{l1.Id(),0}
          if (R.Inverse.Id() != CNULL) { 
            var try_16 EID
            { var arg_17 *Language.If
              var try_18 EID
              { var _CL_obj *Language.If = Language.To_If(new(Language.If).Is(Language.C_If))
                { 
                  var va_arg1 *Language.If
                  var va_arg2 *ClaireAny
                  va_arg1 = _CL_obj
                  { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                    _CL_obj.Selector = Core.C_known_ask
                    _CL_obj.Args = MakeConstantList(z.Id())
                    va_arg2 = _CL_obj.Id()
                    } 
                  va_arg1.Test = va_arg2
                  } 
                { 
                  var va_arg1 *Language.If
                  var va_arg2 *ClaireAny
                  va_arg1 = _CL_obj
                  var try_19 EID
                  try_19 = Core.F_CALL(C_Optimize_Produce_remove,ARGS(EID{R.Inverse.Id(),0},EID{z.Id(),0},EID{x.Id(),0}))
                  if ErrorIn(try_19) {try_18 = try_19
                  } else {
                  va_arg2 = ANY(try_19)
                  va_arg1.Arg = va_arg2
                  try_18 = va_arg2.ToEID()
                  }
                  } 
                if !ErrorIn(try_18) {
                try_18 = EID{_CL_obj.Id(),0}
                }
                } 
              if ErrorIn(try_18) {try_16 = try_18
              } else {
              arg_17 = Language.To_If(OBJ(try_18))
              try_16 = EID{l1.AddFast(arg_17.Id()).Id(),0}
              }
              } 
            if ErrorIn(try_16) {Result = try_16
            } else {
            l1 = ToList(OBJ(try_16))
            Result = EID{l1.Id(),0}
            var try_20 EID
            { var arg_21 *ClaireAny
              var try_22 EID
              try_22 = Core.F_CALL(C_Optimize_Produce_put,ARGS(EID{R.Inverse.Id(),0},EID{y.Id(),0},EID{x.Id(),0}))
              if ErrorIn(try_22) {try_20 = try_22
              } else {
              arg_21 = ANY(try_22)
              try_20 = EID{l1.AddFast(arg_21).Id(),0}
              }
              } 
            if ErrorIn(try_20) {Result = try_20
            } else {
            l1 = ToList(OBJ(try_20))
            Result = EID{l1.Id(),0}
            }
            }
            } else {
            Result = EID{CFALSE.Id(),0}
            } 
          if !ErrorIn(Result) {
          { 
            var va_arg1 *ClaireRelation
            var va_arg2 *ClaireAny
            va_arg1 = R
            var try_23 EID
            { var arg_24 *Language.Let
              var try_25 EID
              { var _CL_obj *Language.Let = Language.To_Let(new(Language.Let).Is(Language.C_Let))
                _CL_obj.ClaireVar = z
                { 
                  var va_arg1 *Language.Let
                  var va_arg2 *ClaireAny
                  va_arg1 = _CL_obj
                  var try_26 EID
                  try_26 = F_Optimize_Produce_get_relation(R,x)
                  if ErrorIn(try_26) {try_25 = try_26
                  } else {
                  va_arg2 = ANY(try_26)
                  va_arg1.Value = va_arg2
                  try_25 = va_arg2.ToEID()
                  }
                  } 
                if !ErrorIn(try_25) {
                { 
                  var va_arg1 *Language.Let
                  var va_arg2 *ClaireAny
                  va_arg1 = _CL_obj
                  { var _CL_obj *Language.If = Language.To_If(new(Language.If).Is(Language.C_If))
                    { 
                      var va_arg1 *Language.If
                      var va_arg2 *ClaireAny
                      va_arg1 = _CL_obj
                      { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                        _CL_obj.Selector = ToProperty(Core.C__I_equal.Id())
                        _CL_obj.Args = MakeConstantList(y.Id(),z.Id())
                        va_arg2 = _CL_obj.Id()
                        } 
                      va_arg1.Test = va_arg2
                      } 
                    _CL_obj.Arg = F_Compile_Do_I_list(l1)
                    va_arg2 = _CL_obj.Id()
                    } 
                  va_arg1.Arg = va_arg2
                  } 
                try_25 = EID{_CL_obj.Id(),0}
                }
                } 
              if ErrorIn(try_25) {try_23 = try_25
              } else {
              arg_24 = Language.To_Let(OBJ(try_25))
              try_23 = Language.F_lambda_I_list(MakeConstantList(x.Id(),y.Id()),arg_24.Id())
              }
              } 
            if ErrorIn(try_23) {Result = try_23
            } else {
            va_arg2 = ANY(try_23)
            va_arg1.IfWrite = va_arg2
            Result = va_arg2.ToEID()
            }
            } 
          }}
          } 
        if !ErrorIn(Result) {
        { var dn *ClaireString = F_append_string(R.Name.String_I(),MakeString("_write"))
          Result = F_Compile_compile_lambda_string(dn,ToLambda(R.IfWrite),C_void.Id())
          } 
        }
        } 
      } 
    } 
  } 
return Result} 

// The EID go function for: Compile/compute_if_write_inverse @ relation (throw: true) 
func E_Compile_compute_if_write_inverse_relation (R EID) EID { 
return F_Compile_compute_if_write_inverse_relation(ToRelation(OBJ(R)) )} 

// generate a demon to perform x.R := s (s is a set)
/* The go function for: Compile/compute_set_write(R:relation) [status=1] */
func F_Compile_compute_set_write_relation (R *ClaireRelation) EID { 
var Result EID
{ var x *ClaireVariable
  { var _CL_obj *ClaireVariable = To_Variable(new(ClaireVariable).Is(C_Variable))
    _CL_obj.Pname = Core.F_symbol_I_string2(MakeString("XX"))
    _CL_obj.Range = R.Domain
    x = _CL_obj
    } 
  { var y *ClaireVariable
    { var _CL_obj *ClaireVariable = To_Variable(new(ClaireVariable).Is(C_Variable))
      _CL_obj.Pname = Core.F_symbol_I_string2(MakeString("YY"))
      _CL_obj.Range = ToType(C_bag.Id())
      y = _CL_obj
      } 
    { var z *ClaireVariable
      { var _CL_obj *ClaireVariable = To_Variable(new(ClaireVariable).Is(C_Variable))
        _CL_obj.Pname = Core.F_symbol_I_string2(MakeString("ZZ"))
        _CL_obj.Range = Core.F_member_type(R.Range)
        z = _CL_obj
        } 
      { var l1 *ClaireList = ToType(C_any.Id()).EmptyList()
        Core.F_tformat_string(MakeString("compute set_write for ~S \n"),0,MakeConstantList(R.Id()))
        if (R.Inverse.Id() != CNULL) { 
          var try_1 EID
          { var arg_2 *Language.For
            var try_3 EID
            { var _CL_obj *Language.For = Language.To_For(new(Language.For).Is(Language.C_For))
              _CL_obj.ClaireVar = z
              { 
                var va_arg1 *Language.Iteration
                var va_arg2 *ClaireAny
                va_arg1 = Language.To_Iteration(_CL_obj.Id())
                var try_4 EID
                try_4 = F_Optimize_Produce_get_relation(R,x)
                if ErrorIn(try_4) {try_3 = try_4
                } else {
                va_arg2 = ANY(try_4)
                va_arg1.SetArg = va_arg2
                try_3 = va_arg2.ToEID()
                }
                } 
              if !ErrorIn(try_3) {
              { 
                var va_arg1 *Language.Iteration
                var va_arg2 *ClaireAny
                va_arg1 = Language.To_Iteration(_CL_obj.Id())
                var try_5 EID
                try_5 = Core.F_CALL(C_Optimize_Produce_remove,ARGS(EID{R.Inverse.Id(),0},EID{z.Id(),0},EID{x.Id(),0}))
                if ErrorIn(try_5) {try_3 = try_5
                } else {
                va_arg2 = ANY(try_5)
                va_arg1.Arg = va_arg2
                try_3 = va_arg2.ToEID()
                }
                } 
              if !ErrorIn(try_3) {
              try_3 = EID{_CL_obj.Id(),0}
              }}
              } 
            if ErrorIn(try_3) {try_1 = try_3
            } else {
            arg_2 = Language.To_For(OBJ(try_3))
            try_1 = EID{l1.AddFast(arg_2.Id()).Id(),0}
            }
            } 
          if ErrorIn(try_1) {Result = try_1
          } else {
          l1 = ToList(OBJ(try_1))
          Result = EID{l1.Id(),0}
          }
          } else {
          Result = EID{CFALSE.Id(),0}
          } 
        if !ErrorIn(Result) {
        var try_6 EID
        { var arg_7 *ClaireAny
          var try_8 EID
          try_8 = F_Optimize_Produce_erase_property(ToProperty(R.Id()),x)
          if ErrorIn(try_8) {try_6 = try_8
          } else {
          arg_7 = ANY(try_8)
          try_6 = EID{l1.AddFast(arg_7).Id(),0}
          }
          } 
        if ErrorIn(try_6) {Result = try_6
        } else {
        l1 = ToList(OBJ(try_6))
        Result = EID{l1.Id(),0}
        var try_9 EID
        { var arg_10 *Language.For
          var try_11 EID
          { var _CL_obj *Language.For = Language.To_For(new(Language.For).Is(Language.C_For))
            _CL_obj.ClaireVar = z
            _CL_obj.SetArg = y.Id()
            { 
              var va_arg1 *Language.Iteration
              var va_arg2 *ClaireAny
              va_arg1 = Language.To_Iteration(_CL_obj.Id())
              var try_12 EID
              try_12 = Core.F_CALL(C_Optimize_Produce_put,ARGS(EID{R.Id(),0},EID{x.Id(),0},EID{z.Id(),0}))
              if ErrorIn(try_12) {try_11 = try_12
              } else {
              va_arg2 = ANY(try_12)
              va_arg1.Arg = va_arg2
              try_11 = va_arg2.ToEID()
              }
              } 
            if !ErrorIn(try_11) {
            try_11 = EID{_CL_obj.Id(),0}
            }
            } 
          if ErrorIn(try_11) {try_9 = try_11
          } else {
          arg_10 = Language.To_For(OBJ(try_11))
          try_9 = EID{l1.AddFast(arg_10.Id()).Id(),0}
          }
          } 
        if ErrorIn(try_9) {Result = try_9
        } else {
        l1 = ToList(OBJ(try_9))
        Result = EID{l1.Id(),0}
        { var dn *ClaireString = F_append_string(R.Name.String_I(),MakeString("_set_write"))
          { var arg_13 *ClaireLambda
            var try_14 EID
            try_14 = Language.F_lambda_I_list(MakeConstantList(x.Id(),y.Id()),F_Compile_Do_I_list(l1))
            if ErrorIn(try_14) {Result = try_14
            } else {
            arg_13 = ToLambda(OBJ(try_14))
            Result = F_Compile_compile_lambda_string(dn,arg_13,C_void.Id())
            }
            } 
          } 
        }}}
        } 
      } 
    } 
  } 
return Result} 

// The EID go function for: Compile/compute_set_write @ relation (throw: true) 
func E_Compile_compute_set_write_relation (R EID) EID { 
return F_Compile_compute_set_write_relation(ToRelation(OBJ(R)) )} 

// generate a simple put for a property => generate a case to make sure
// that we get the fastest possible code
/* The go function for: Produce_put(r:property,x:Variable,y:any) [status=1] */
func F_Optimize_Produce_put_property (r *ClaireProperty,x *ClaireVariable,y *ClaireAny) EID { 
var Result EID
{ var l *ClaireList = ToType(C_any.Id()).EmptyList()
  { 
    var xs *ClaireRestriction
    _ = xs
    var xs_iter *ClaireAny
    Result= EID{CFALSE.Id(),0}
    for _,xs_iter = range(r.Restrictions.ValuesO()){ 
      xs = ToRestriction(xs_iter)
      var loop_1 EID
      _ = loop_1
      if ((C_slot.Id() == xs.Isa.Id()) && 
          (F_boolean_I_any(Core.F__exp_type(F_Optimize_ptype_type(x.Range),ToType(Core.F_domain_I_restriction(xs).Id())).Id()) == CTRUE)) { 
        var try_2 EID
        { var arg_3 *ClaireList
          { 
            var v_bag_arg *ClaireAny
            arg_3= ToType(CEMPTY.Id()).EmptyList()
            arg_3.AddFast(Core.F_domain_I_restriction(xs).Id())
            if (r.Multivalued_ask == CTRUE) { 
              { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                _CL_obj.Selector = ToProperty(C_add_I.Id())
                { 
                  var va_arg1 *Language.Call
                  var va_arg2 *ClaireList
                  va_arg1 = _CL_obj
                  { 
                    var v_bag_arg *ClaireAny
                    va_arg2= ToType(CEMPTY.Id()).EmptyList()
                    { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                      _CL_obj.Selector = r
                      { 
                        var va_arg1 *Language.Call
                        var va_arg2 *ClaireList
                        va_arg1 = _CL_obj
                        { 
                          var v_bag_arg *ClaireAny
                          va_arg2= ToType(CEMPTY.Id()).EmptyList()
                          { var _CL_obj *Language.Cast = Language.To_Cast(new(Language.Cast).Is(Language.C_Cast))
                            _CL_obj.Arg = x.Id()
                            _CL_obj.SetArg = ToType(Core.F_domain_I_restriction(xs).Id())
                            v_bag_arg = _CL_obj.Id()
                            } 
                          va_arg2.AddFast(v_bag_arg)} 
                        va_arg1.Args = va_arg2
                        } 
                      v_bag_arg = _CL_obj.Id()
                      } 
                    va_arg2.AddFast(v_bag_arg)
                    va_arg2.AddFast(y)} 
                  va_arg1.Args = va_arg2
                  } 
                v_bag_arg = _CL_obj.Id()
                } 
              } else {
              { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                _CL_obj.Selector = C_put
                { 
                  var va_arg1 *Language.Call
                  var va_arg2 *ClaireList
                  va_arg1 = _CL_obj
                  { 
                    var v_bag_arg *ClaireAny
                    va_arg2= ToType(CEMPTY.Id()).EmptyList()
                    va_arg2.AddFast(r.Id())
                    { var _CL_obj *Language.Cast = Language.To_Cast(new(Language.Cast).Is(Language.C_Cast))
                      _CL_obj.Arg = x.Id()
                      _CL_obj.SetArg = ToType(Core.F_domain_I_restriction(xs).Id())
                      v_bag_arg = _CL_obj.Id()
                      } 
                    va_arg2.AddFast(v_bag_arg)
                    va_arg2.AddFast(y)} 
                  va_arg1.Args = va_arg2
                  } 
                v_bag_arg = _CL_obj.Id()
                } 
              } 
            arg_3.AddFast(v_bag_arg)} 
          try_2 = l.Add_star(arg_3)
          } 
        if ErrorIn(try_2) {loop_1 = try_2
        } else {
        l = ToList(OBJ(try_2))
        loop_1 = EID{l.Id(),0}
        }
        } else {
        loop_1 = EID{CFALSE.Id(),0}
        } 
      if ErrorIn(loop_1) {Result = loop_1
      break
      } else {
      }
      } 
    } 
  if !ErrorIn(Result) {
  if (l.Length() == 2) { 
    Result = l.At(1).ToEID()
    } else {
    { var _CL_obj *Language.Case = Language.To_Case(new(Language.Case).Is(Language.C_Case))
      _CL_obj.ClaireVar = x.Id()
      _CL_obj.Args = l
      Result = EID{_CL_obj.Id(),0}
      } 
    } 
  }
  } 
return Result} 

// The EID go function for: Produce_put @ property (throw: true) 
func E_Optimize_Produce_put_property (r EID,x EID,y EID) EID { 
return F_Optimize_Produce_put_property(ToProperty(OBJ(r)),To_Variable(OBJ(x)),ANY(y) )} 

// generate a simple erase (the inverse management has been done)
// v3.2.50: use ptype(x.range) for variable whose type is t U any :-)
/* The go function for: Produce_erase(r:property,x:Variable) [status=1] */
func F_Optimize_Produce_erase_property (r *ClaireProperty,x *ClaireVariable) EID { 
var Result EID
{ var l *ClaireList = ToType(C_any.Id()).EmptyList()
  { var val *ClaireSet = ToType(C_any.Id()).EmptySet()
    val.Cast_I(Core.F_member_type(r.Range))
    { 
      var xs *ClaireRestriction
      _ = xs
      var xs_iter *ClaireAny
      Result= EID{CFALSE.Id(),0}
      for _,xs_iter = range(r.Restrictions.ValuesO()){ 
        xs = ToRestriction(xs_iter)
        var loop_1 EID
        _ = loop_1
        if ((C_slot.Id() == xs.Isa.Id()) && 
            (F_boolean_I_any(Core.F__exp_type(F_Optimize_ptype_type(x.Range),ToType(Core.F_domain_I_restriction(xs).Id())).Id()) == CTRUE)) { 
          var try_2 EID
          { var arg_3 *ClaireList
            { 
              var v_bag_arg *ClaireAny
              arg_3= ToType(CEMPTY.Id()).EmptyList()
              arg_3.AddFast(Core.F_domain_I_restriction(xs).Id())
              { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                _CL_obj.Selector = C_put
                { 
                  var va_arg1 *Language.Call
                  var va_arg2 *ClaireList
                  va_arg1 = _CL_obj
                  { 
                    var v_bag_arg *ClaireAny
                    va_arg2= ToType(CEMPTY.Id()).EmptyList()
                    va_arg2.AddFast(r.Id())
                    { var _CL_obj *Language.Cast = Language.To_Cast(new(Language.Cast).Is(Language.C_Cast))
                      _CL_obj.Arg = x.Id()
                      _CL_obj.SetArg = ToType(Core.F_domain_I_restriction(xs).Id())
                      v_bag_arg = _CL_obj.Id()
                      } 
                    va_arg2.AddFast(v_bag_arg)
                    if (r.Multivalued_ask == CTRUE) { 
                      v_bag_arg = val.Id()
                      } else {
                      v_bag_arg = ToSlot(xs.Id()).Default
                      } 
                    va_arg2.AddFast(v_bag_arg)} 
                  va_arg1.Args = va_arg2
                  } 
                v_bag_arg = _CL_obj.Id()
                } 
              arg_3.AddFast(v_bag_arg)} 
            try_2 = l.Add_star(arg_3)
            } 
          if ErrorIn(try_2) {loop_1 = try_2
          } else {
          l = ToList(OBJ(try_2))
          loop_1 = EID{l.Id(),0}
          }
          } else {
          loop_1 = EID{CFALSE.Id(),0}
          } 
        if ErrorIn(loop_1) {Result = loop_1
        break
        } else {
        }
        } 
      } 
    if !ErrorIn(Result) {
    if (l.Length() == 2) { 
      Result = l.At(1).ToEID()
      } else {
      { var _CL_obj *Language.Case = Language.To_Case(new(Language.Case).Is(Language.C_Case))
        _CL_obj.ClaireVar = x.Id()
        _CL_obj.Args = l
        Result = EID{_CL_obj.Id(),0}
        } 
      } 
    }
    } 
  } 
return Result} 

// The EID go function for: Produce_erase @ property (throw: true) 
func E_Optimize_Produce_erase_property (r EID,x EID) EID { 
return F_Optimize_Produce_erase_property(ToProperty(OBJ(r)),To_Variable(OBJ(x)) )} 

// note:  (a) Simpler because of v3.0 !! (siude-effects on lists or sets)
//        (b) if |l|= 1 domain!(r) = domain!(x) because of tighten
// same for a table
/* The go function for: Produce_put(r:table,x:Variable,y:any) [status=0] */
func F_Optimize_Produce_put_table (r *ClaireTable,x *ClaireVariable,y *ClaireAny) *ClaireAny { 
var Result *ClaireAny
{ var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
  _CL_obj.Selector = C_put
  { 
    var va_arg1 *Language.Call
    var va_arg2 *ClaireList
    va_arg1 = _CL_obj
    { 
      var v_bag_arg *ClaireAny
      va_arg2= ToType(CEMPTY.Id()).EmptyList()
      va_arg2.AddFast(r.Id())
      va_arg2.AddFast(x.Id())
      if (r.Multivalued_ask == CTRUE) { 
        { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
          _CL_obj.Selector = ToProperty(C_add.Id())
          _CL_obj.Args = MakeConstantList(MakeConstantList(C_nth.Id(),MakeConstantList(r.Id(),x.Id()).Id()).Id(),y)
          v_bag_arg = _CL_obj.Id()
          } 
        } else {
        v_bag_arg = y
        } 
      va_arg2.AddFast(v_bag_arg)} 
    va_arg1.Args = va_arg2
    } 
  Result = _CL_obj.Id()
  } 
return Result} 

// The EID go function for: Produce_put @ table (throw: false) 
func E_Optimize_Produce_put_table (r EID,x EID,y EID) EID { 
return F_Optimize_Produce_put_table(ToTable(OBJ(r)),To_Variable(OBJ(x)),ANY(y) ).ToEID()} 

/* The go function for: Produce_get(r:relation,x:Variable) [status=1] */
func F_Optimize_Produce_get_relation (r *ClaireRelation,x *ClaireVariable) EID { 
var Result EID
if (C_table.Id() == r.Isa.Id()) { 
  { var g0326 *ClaireTable = ToTable(r.Id())
    { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
      _CL_obj.Selector = C_nth
      _CL_obj.Args = MakeConstantList(g0326.Id(),x.Id())
      Result = EID{_CL_obj.Id(),0}
      } 
    } 
  }  else if (r.Isa.IsIn(C_property) == CTRUE) { 
  { var g0327 *ClaireProperty = ToProperty(r.Id())
    { var l *ClaireList = ToType(C_any.Id()).EmptyList()
      { 
        var xs *ClaireRestriction
        _ = xs
        var xs_iter *ClaireAny
        Result= EID{CFALSE.Id(),0}
        for _,xs_iter = range(g0327.Restrictions.ValuesO()){ 
          xs = ToRestriction(xs_iter)
          var loop_1 EID
          _ = loop_1
          if ((C_slot.Id() == xs.Isa.Id()) && 
              (F_boolean_I_any(Core.F__exp_type(F_Optimize_ptype_type(x.Range),ToType(Core.F_domain_I_restriction(xs).Id())).Id()) == CTRUE)) { 
            var try_2 EID
            { var arg_3 *ClaireList
              { 
                var v_bag_arg *ClaireAny
                arg_3= ToType(CEMPTY.Id()).EmptyList()
                arg_3.AddFast(Core.F_domain_I_restriction(xs).Id())
                { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                  _CL_obj.Selector = g0327
                  { 
                    var va_arg1 *Language.Call
                    var va_arg2 *ClaireList
                    va_arg1 = _CL_obj
                    { 
                      var v_bag_arg *ClaireAny
                      va_arg2= ToType(CEMPTY.Id()).EmptyList()
                      { var _CL_obj *Language.Cast = Language.To_Cast(new(Language.Cast).Is(Language.C_Cast))
                        _CL_obj.Arg = x.Id()
                        _CL_obj.SetArg = ToType(Core.F_domain_I_restriction(xs).Id())
                        v_bag_arg = _CL_obj.Id()
                        } 
                      va_arg2.AddFast(v_bag_arg)} 
                    va_arg1.Args = va_arg2
                    } 
                  v_bag_arg = _CL_obj.Id()
                  } 
                arg_3.AddFast(v_bag_arg)} 
              try_2 = l.Add_star(arg_3)
              } 
            if ErrorIn(try_2) {loop_1 = try_2
            } else {
            l = ToList(OBJ(try_2))
            loop_1 = EID{l.Id(),0}
            }
            } else {
            loop_1 = EID{CFALSE.Id(),0}
            } 
          if ErrorIn(loop_1) {Result = loop_1
          break
          } else {
          }
          } 
        } 
      if !ErrorIn(Result) {
      if (l.Length() == 2) { 
        Result = l.At(1).ToEID()
        } else {
        { var _CL_obj *Language.Case = Language.To_Case(new(Language.Case).Is(Language.C_Case))
          _CL_obj.ClaireVar = x.Id()
          _CL_obj.Args = l
          Result = EID{_CL_obj.Id(),0}
          } 
        } 
      }
      } 
    } 
  } else {
  Result = EID{CFALSE.Id(),0}
  } 
return Result} 

// The EID go function for: Produce_get @ relation (throw: true) 
func E_Optimize_Produce_get_relation (r EID,x EID) EID { 
return F_Optimize_Produce_get_relation(ToRelation(OBJ(r)),To_Variable(OBJ(x)) )} 

// generate a remove
/* The go function for: Produce_remove(r:property,x:Variable,y:any) [status=1] */
func F_Optimize_Produce_remove_property (r *ClaireProperty,x *ClaireVariable,y *ClaireAny) EID { 
var Result EID
{ var l *ClaireList = ToType(C_any.Id()).EmptyList()
  { 
    var xs *ClaireRestriction
    _ = xs
    var xs_iter *ClaireAny
    Result= EID{CFALSE.Id(),0}
    for _,xs_iter = range(r.Restrictions.ValuesO()){ 
      xs = ToRestriction(xs_iter)
      var loop_1 EID
      _ = loop_1
      if (C_slot.Id() == xs.Isa.Id()) { 
        var try_2 EID
        { var arg_3 *ClaireList
          { 
            var v_bag_arg *ClaireAny
            arg_3= ToType(CEMPTY.Id()).EmptyList()
            arg_3.AddFast(Core.F_domain_I_restriction(xs).Id())
            if (r.Multivalued_ask == CTRUE) { 
              { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                _CL_obj.Selector = ToProperty(C_delete.Id())
                { 
                  var va_arg1 *Language.Call
                  var va_arg2 *ClaireList
                  va_arg1 = _CL_obj
                  { 
                    var v_bag_arg *ClaireAny
                    va_arg2= ToType(CEMPTY.Id()).EmptyList()
                    { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                      _CL_obj.Selector = r
                      _CL_obj.Args = MakeConstantList(x.Id())
                      v_bag_arg = _CL_obj.Id()
                      } 
                    va_arg2.AddFast(v_bag_arg)
                    va_arg2.AddFast(y)} 
                  va_arg1.Args = va_arg2
                  } 
                v_bag_arg = _CL_obj.Id()
                } 
              } else {
              { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                _CL_obj.Selector = C_put
                _CL_obj.Args = MakeConstantList(r.Id(),x.Id(),CNULL)
                v_bag_arg = _CL_obj.Id()
                } 
              } 
            arg_3.AddFast(v_bag_arg)} 
          try_2 = l.Add_star(arg_3)
          } 
        if ErrorIn(try_2) {loop_1 = try_2
        } else {
        l = ToList(OBJ(try_2))
        loop_1 = EID{l.Id(),0}
        }
        } else {
        loop_1 = EID{CFALSE.Id(),0}
        } 
      if ErrorIn(loop_1) {Result = loop_1
      break
      } else {
      }
      } 
    } 
  if !ErrorIn(Result) {
  if (l.Length() == 2) { 
    Result = l.At(1).ToEID()
    } else {
    { var _CL_obj *Language.Case = Language.To_Case(new(Language.Case).Is(Language.C_Case))
      _CL_obj.ClaireVar = x.Id()
      _CL_obj.Args = l
      Result = EID{_CL_obj.Id(),0}
      } 
    } 
  }
  } 
return Result} 

// The EID go function for: Produce_remove @ property (throw: true) 
func E_Optimize_Produce_remove_property (r EID,x EID,y EID) EID { 
return F_Optimize_Produce_remove_property(ToProperty(OBJ(r)),To_Variable(OBJ(x)),ANY(y) )} 

// same for a table
/* The go function for: Produce_remove(r:table,x:Variable,y:any) [status=0] */
func F_Optimize_Produce_remove_table (r *ClaireTable,x *ClaireVariable,y *ClaireAny) *ClaireAny { 
var Result *ClaireAny
{ var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
  _CL_obj.Selector = C_put
  { 
    var va_arg1 *Language.Call
    var va_arg2 *ClaireList
    va_arg1 = _CL_obj
    { 
      var v_bag_arg *ClaireAny
      va_arg2= ToType(CEMPTY.Id()).EmptyList()
      va_arg2.AddFast(r.Id())
      va_arg2.AddFast(x.Id())
      if (r.Multivalued_ask == CTRUE) { 
        { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
          _CL_obj.Selector = ToProperty(C_delete.Id())
          _CL_obj.Args = MakeConstantList(MakeConstantList(C_nth.Id(),MakeConstantList(r.Id(),x.Id()).Id()).Id(),y)
          v_bag_arg = _CL_obj.Id()
          } 
        } else {
        v_bag_arg = CNULL
        } 
      va_arg2.AddFast(v_bag_arg)} 
    va_arg1.Args = va_arg2
    } 
  Result = _CL_obj.Id()
  } 
return Result} 

// The EID go function for: Produce_remove @ table (throw: false) 
func E_Optimize_Produce_remove_table (r EID,x EID,y EID) EID { 
return F_Optimize_Produce_remove_table(ToTable(OBJ(r)),To_Variable(OBJ(x)),ANY(y) ).ToEID()} 

// computes the best range and domain then sets r.open to 1
/* The go function for: Tighten(r:relation) [status=0] */
func F_Optimize_Tighten_relation (r *ClaireRelation)  { 
if (r.Isa.IsIn(C_property) == CTRUE) { 
  { var g0328 *ClaireProperty = ToProperty(r.Id())
    { var ad *ClaireType = ToType(CEMPTY.Id())
      { var ar *ClaireType = ToType(CEMPTY.Id())
        { 
          var s *ClaireRestriction
          _ = s
          var s_iter *ClaireAny
          for _,s_iter = range(g0328.Restrictions.ValuesO()){ 
            s = ToRestriction(s_iter)
            if (C_slot.Id() == s.Isa.Id()) { 
              ad = Core.F_U_type(ad,ToType(Core.F_domain_I_restriction(s).Id()))
              { var arg_1 *ClaireType
                if (g0328.Multivalued_ask == CTRUE) { 
                  arg_1 = Core.F_member_type(s.Range)
                  } else {
                  arg_1 = s.Range
                  } 
                ar = Core.F_U_type(ar,arg_1)
                } 
              } 
            } 
          } 
        g0328.Open = 1
        if (Equal(ad.Id(),CEMPTY.Id()) != CTRUE) { 
          g0328.Domain = ToType(ad.Class_I().Id())
          { 
            var va_arg1 *ClaireRelation
            var va_arg2 *ClaireType
            va_arg1 = ToRelation(g0328.Id())
            if (g0328.Multivalued_ask == CTRUE) { 
              va_arg2 = Core.F_param_I_class(C_set,ToType(ar.Class_I().Id()))
              } else {
              va_arg2 = ar
              } 
            va_arg1.Range = va_arg2
            } 
          } 
        
        } 
      } 
    } 
  } 
} 

// The EID go function for: Tighten @ relation (throw: false) 
func E_Optimize_Tighten_relation (r EID) EID { 
F_Optimize_Tighten_relation(ToRelation(OBJ(r)) )
return EVOID} 

/* The go function for: Compile/Tighten!(r:relation) [status=0] */
func F_Compile_Tighten_I_relation (r *ClaireRelation)  { 
if (r.Open > 0) { 
  F_Optimize_Tighten_relation(r)
  } 
} 

// The EID go function for: Compile/Tighten! @ relation (throw: false) 
func E_Compile_Tighten_I_relation (r EID) EID { 
F_Compile_Tighten_I_relation(ToRelation(OBJ(r)) )
return EVOID} 

// new: re-compute the numbering but without the side-effects of the interpreter version (v3.067)
/* The go function for: Compile/lexical_num(self:any,n:integer) [status=1] */
func F_Compile_lexical_num_any (self *ClaireAny,n int) EID { 
var Result EID
if (self.Isa.IsIn(Language.C_Call) == CTRUE) { 
  { var g0329 *Language.Call = Language.To_Call(self)
    Result = F_Compile_lexical_num_any(g0329.Args.Id(),n)
    } 
  }  else if (self.Isa.IsIn(C_Instruction) == CTRUE) { 
  { var g0330 *ClaireInstruction = To_Instruction(self)
    { var _Ztype *ClaireClass = g0330.Isa
      if (Language.C_Instruction_with_var.Descendants.Contain_ask(_Ztype.Id()) == CTRUE) { 
        Result = Core.F_put_property2(C_mClaire_index,ToObject(OBJ(Core.F_CALL(Language.C_var,ARGS(EID{g0330.Id(),0})))),MakeInteger(n).Id())
        if !ErrorIn(Result) {
        n = (n+1)
        if (n > ToInteger(Language.C__starvariable_index_star.Value).Value) { 
          var v_gassign1 *ClaireAny
          v_gassign1 = MakeInteger(n).Id()
          Language.C__starvariable_index_star.Value = v_gassign1
          Result = v_gassign1.ToEID()
          } else {
          Result = EID{CFALSE.Id(),0}
          } 
        }
        } else {
        Result = EID{CFALSE.Id(),0}
        } 
      if !ErrorIn(Result) {
      { 
        var s *ClaireSlot
        _ = s
        var s_iter *ClaireAny
        Result= EID{CFALSE.Id(),0}
        for _,s_iter = range(_Ztype.Slots.ValuesO()){ 
          s = ToSlot(s_iter)
          var loop_2 EID
          _ = loop_2
          loop_2 = F_Compile_lexical_num_any(Core.F_get_slot(s,ToObject(g0330.Id())),n)
          if ErrorIn(loop_2) {Result = loop_2
          break
          } else {
          }
          } 
        } 
      }
      } 
    } 
  }  else if (self.Isa.IsIn(C_bag) == CTRUE) { 
  { var g0331 *ClaireBag = ToBag(self)
    { 
      var x *ClaireAny
      _ = x
      Result= EID{CFALSE.Id(),0}
      var x_support *ClaireList
      var try_3 EID
      try_3 = Core.F_enumerate_any(g0331.Id())
      if ErrorIn(try_3) {Result = try_3
      } else {
      x_support = ToList(OBJ(try_3))
      x_len := x_support.Length()
      for i_it := 0; i_it < x_len; i_it++ { 
        x = x_support.At(i_it)
        var loop_4 EID
        _ = loop_4
        loop_4 = F_Compile_lexical_num_any(x,n)
        if ErrorIn(loop_4) {Result = loop_4
        break
        } else {
        }}
        } 
      } 
    } 
  } else {
  Result = EID{CNIL.Id(),0}
  } 
return Result} 

// The EID go function for: Compile/lexical_num @ any (throw: true) 
func E_Compile_lexical_num_any (self EID,n EID) EID { 
return F_Compile_lexical_num_any(ANY(self),INT(n) )} 

// v3.2 -----------------------------------------------------------------
/* The go function for: c_type(self:Defrule) [status=0] */
func F_c_type_Defrule (self *Language.Defrule) *ClaireType { 
return  ToType(C_any.Id())
} 

// The EID go function for: c_type @ Defrule (throw: false) 
func E_c_type_Defrule (self EID) EID { 
return EID{F_c_type_Defrule(Language.To_Defrule(OBJ(self)) ).Id(),0}} 

// compile a rule definition
/* The go function for: c_code(self:Defrule,s:class) [status=1] */
func F_c_code_Defrule (self *Language.Defrule,s *ClaireClass) EID { 
var Result EID
{ var ru *ClaireAny = self.Ident.Value()
  { var l *ClaireList = ToType(C_any.Id()).EmptyList()
    
    { 
      var r *ClaireAny
      _ = r
      var r_support *ClaireSet
      r_support = ToSet(Core.F_get_table(Language.C_Language_relations,ru))
      for i_it := 0; i_it < r_support.Count; i_it++ { 
        r = r_support.At(i_it)
        if (Language.F_eventMethod_ask_relation2(ToRelation(r)) != CTRUE) { 
          F_Optimize_Tighten_relation(ToRelation(r))
          } 
        } 
      } 
    { 
      var r *ClaireAny
      _ = r
      Result= EID{CFALSE.Id(),0}
      var r_support *ClaireSet
      r_support = ToSet(Core.F_get_table(Language.C_Language_relations,ru))
      for i_it := 0; i_it < r_support.Count; i_it++ { 
        r = r_support.At(i_it)
        var loop_1 EID
        _ = loop_1
        { 
        if (INT(Core.F_CALL(C_open,ARGS(r.ToEID()))) < 2) { 
          { var arg_2 *Language.Call
            { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
              _CL_obj.Selector = C_final
              _CL_obj.Args = MakeConstantList(r)
              arg_2 = _CL_obj
              } 
            l = l.AddFast(arg_2.Id())
            } 
          } 
        loop_1 = F_Optimize_compile_if_write_relation(ToRelation(r))
        if ErrorIn(loop_1) {Result = loop_1
        break
        } else {
        { var dn *ClaireAny = ANY(Core.F_CALL(ToProperty(C__7_plus.Id()),ARGS(Core.F_CALL(C_name,ARGS(r.ToEID())),EID{MakeString("_write").Id(),0})))
          { var s *ClaireString
            var try_3 EID
            try_3 = Core.F_CALL(C_string_I,ARGS(dn.ToEID()))
            if ErrorIn(try_3) {loop_1 = try_3
            } else {
            s = ToString(OBJ(try_3))
            { var lb *ClaireAny = ANY(Core.F_CALL(C_if_write,ARGS(r.ToEID())))
              
              loop_1 = F_Compile_compile_lambda_string(s,ToLambda(lb),C_void.Id())
              if ErrorIn(loop_1) {Result = loop_1
              break
              } else {
              { var arg_4 *Language.Call
                { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                  _CL_obj.Selector = C_put
                  _CL_obj.Args = MakeConstantList(C_if_write.Id(),r,F_Optimize_demon_function_string(s).Id())
                  arg_4 = _CL_obj
                  } 
                loop_1 = EID{l.AddFast(arg_4.Id()).Id(),0}
                } 
              }
              } 
            }
            } 
          } 
        if ErrorIn(loop_1) {Result = loop_1
        break
        } else {
        }}
        }
        } 
      } 
    if !ErrorIn(Result) {
    { 
      var r *ClaireAny
      _ = r
      Result= EID{CFALSE.Id(),0}
      var r_support *ClaireSet
      r_support = ToSet(Core.F_get_table(Language.C_Language_relations,ru))
      for i_it := 0; i_it < r_support.Count; i_it++ { 
        r = r_support.At(i_it)
        var loop_5 EID
        _ = loop_5
        if (Language.F_eventMethod_ask_relation2(ToRelation(r)) == CTRUE) { 
          var try_6 EID
          { var arg_7 *ClaireAny
            var try_8 EID
            try_8 = F_Optimize_compileEventMethod_property(ToProperty(r))
            if ErrorIn(try_8) {try_6 = try_8
            } else {
            arg_7 = ANY(try_8)
            try_6 = EID{l.AddFast(arg_7).Id(),0}
            }
            } 
          if ErrorIn(try_6) {loop_5 = try_6
          } else {
          l = ToList(OBJ(try_6))
          loop_5 = EID{l.Id(),0}
          }
          } else {
          loop_5 = EID{CFALSE.Id(),0}
          } 
        if ErrorIn(loop_5) {Result = loop_5
        break
        } else {
        }
        } 
      } 
    if !ErrorIn(Result) {
    Result = Core.F_CALL(C_c_code,ARGS(F_Compile_Do_I_list(l).ToEID(),EID{s.Id(),0}))
    }}
    } 
  } 
return Result} 

// The EID go function for: c_code @ Defrule (throw: true) 
func E_c_code_Defrule (self EID,s EID) EID { 
return F_c_code_Defrule(Language.To_Defrule(OBJ(self)),ToClass(OBJ(s)) )} 

// produce a beautiful if_write demon from all the claire demons created by each rule that applies to R
/* The go function for: compile_if_write(R:relation) [status=1] */
func F_Optimize_compile_if_write_relation (R *ClaireRelation) EID { 
var Result EID
{ var l *ClaireList = ToList(Core.F_get_table(Language.C_demons,R.Id()))
  { var lvar *ClaireList = Language.ToLanguageDemon(l.ValuesO()[0]).Formula.Vars
    { var l1 *ClaireList
      var try_1 EID
      { 
        var v_bag_arg *ClaireAny
        try_1= EID{ToType(C_any.Id()).EmptyList().Id(),0}
        var try_2 EID
        try_2 = Core.F_CALL(C_Optimize_Produce_put,ARGS(EID{R.Id(),0},lvar.At(0).ToEID(),lvar.At(1).ToEID()))
        if ErrorIn(try_2) {try_1 = try_2
        } else {
        v_bag_arg = ANY(try_2)
        ToList(OBJ(try_1)).AddFast(v_bag_arg)}
        } 
      if ErrorIn(try_1) {Result = try_1
      } else {
      l1 = ToList(OBJ(try_1))
      { var l2 *ClaireList
        { 
          var v_list4 *ClaireList
          var x *Language.LanguageDemon
          var v_local4 *ClaireAny
          v_list4 = l
          l2 = CreateList(ToType(C_any.Id()),v_list4.Length())
          for CLcount := 0; CLcount < v_list4.Length(); CLcount++{ 
            x = Language.ToLanguageDemon(v_list4.At(CLcount))
            v_local4 = F_Optimize_demon_substitution_demon(x,lvar)
            l2.PutAt(CLcount,v_local4)
            } 
          } 
        F_Optimize_Tighten_relation(R)
        Result = Core.F_put_property2(C_range,ToObject(lvar.At(0)),R.Domain.Id())
        if !ErrorIn(Result) {
        Result = Core.F_put_property2(C_range,ToObject(lvar.At(1)),R.Range.Id())
        if !ErrorIn(Result) {
        { 
          var v *ClaireAny
          _ = v
          Result= EID{CFALSE.Id(),0}
          var v_support *ClaireList
          v_support = lvar
          v_len := v_support.Length()
          for i_it := 0; i_it < v_len; i_it++ { 
            v = v_support.At(i_it)
            var loop_3 EID
            _ = loop_3
            loop_3 = Core.F_put_property2(C_range,ToObject(v),ToTypeExpression(OBJ(Core.F_CALL(C_range,ARGS(v.ToEID())))).Class_I().Id())
            if ErrorIn(loop_3) {Result = loop_3
            break
            } else {
            }
            } 
          } 
        if !ErrorIn(Result) {
        if ((l2.At(0).Isa.IsIn(Language.C_If) == CTRUE) && 
            (Language.F_eventMethod_ask_relation2(R) != CTRUE)) { 
          if (Core.F_owner_any(Language.To_If(l2.At(0)).Test).IsIn(Language.C_And) == CTRUE) { 
            { 
              var va_arg1 *Language.If
              var va_arg2 *ClaireAny
              va_arg1 = Language.To_If(l2.At(0))
              var try_4 EID
              { var _CL_obj *Language.And = Language.To_And(new(Language.And).Is(Language.C_And))
                { 
                  var va_arg1 *Language.And
                  var va_arg2 *ClaireList
                  va_arg1 = _CL_obj
                  var try_5 EID
                  try_5 = ToList(OBJ(Core.F_CALL(C_args,ARGS(Language.To_If(l2.At(0)).Test.ToEID())))).Cdr()
                  if ErrorIn(try_5) {try_4 = try_5
                  } else {
                  va_arg2 = ToList(OBJ(try_5))
                  va_arg1.Args = va_arg2
                  try_4 = EID{va_arg2.Id(),0}
                  }
                  } 
                if !ErrorIn(try_4) {
                try_4 = EID{_CL_obj.Id(),0}
                }
                } 
              if ErrorIn(try_4) {Result = try_4
              } else {
              va_arg2 = ANY(try_4)
              va_arg1.Test = va_arg2
              Result = va_arg2.ToEID()
              }
              } 
            } else {
            Result = ToArray(l2.Id()).NthPut(1,Language.To_If(l2.At(0)).Arg).ToEID()
            } 
          } else {
          Result = EID{CFALSE.Id(),0}
          } 
        if !ErrorIn(Result) {
        if (R.Inverse.Id() != CNULL) { 
          if (R.Multivalued_ask != CTRUE) { 
            var try_6 EID
            { var arg_7 *ClaireAny
              var try_8 EID
              try_8 = Core.F_CALL(C_Optimize_Produce_remove,ARGS(EID{R.Inverse.Id(),0},lvar.At(2).ToEID(),lvar.At(0).ToEID()))
              if ErrorIn(try_8) {try_6 = try_8
              } else {
              arg_7 = ANY(try_8)
              try_6 = EID{l1.AddFast(arg_7).Id(),0}
              }
              } 
            if ErrorIn(try_6) {Result = try_6
            } else {
            l1 = ToList(OBJ(try_6))
            Result = EID{l1.Id(),0}
            }
            } else {
            Result = EID{CFALSE.Id(),0}
            } 
          if !ErrorIn(Result) {
          var try_9 EID
          { var arg_10 *ClaireAny
            var try_11 EID
            try_11 = Core.F_CALL(C_Optimize_Produce_put,ARGS(EID{R.Inverse.Id(),0},lvar.At(1).ToEID(),lvar.At(0).ToEID()))
            if ErrorIn(try_11) {try_9 = try_11
            } else {
            arg_10 = ANY(try_11)
            try_9 = EID{l1.AddFast(arg_10).Id(),0}
            }
            } 
          if ErrorIn(try_9) {Result = try_9
          } else {
          l1 = ToList(OBJ(try_9))
          Result = EID{l1.Id(),0}
          }
          }
          } else {
          Result = EID{CFALSE.Id(),0}
          } 
        if !ErrorIn(Result) {
        { 
          var va_arg1 *ClaireRelation
          var va_arg2 *ClaireAny
          va_arg1 = R
          var try_12 EID
          { var arg_13 *ClaireAny
            if (Language.F_eventMethod_ask_relation2(R) == CTRUE) { 
              arg_13 = F_Compile_Do_I_list(l2)
              }  else if (R.Multivalued_ask == CTRUE) { 
              { var _CL_obj *Language.If = Language.To_If(new(Language.If).Is(Language.C_If))
                { 
                  var va_arg1 *Language.If
                  var va_arg2 *ClaireAny
                  va_arg1 = _CL_obj
                  { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                    _CL_obj.Selector = Core.C_not
                    { 
                      var va_arg1 *Language.Call
                      var va_arg2 *ClaireList
                      va_arg1 = _CL_obj
                      { 
                        var v_bag_arg *ClaireAny
                        va_arg2= ToType(CEMPTY.Id()).EmptyList()
                        { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                          _CL_obj.Selector = ToProperty(C__Z.Id())
                          _CL_obj.Args = MakeConstantList(lvar.At(1),Language.F_readCall_relation(R,lvar.At(0)).Id())
                          v_bag_arg = _CL_obj.Id()
                          } 
                        va_arg2.AddFast(v_bag_arg)} 
                      va_arg1.Args = va_arg2
                      } 
                    va_arg2 = _CL_obj.Id()
                    } 
                  va_arg1.Test = va_arg2
                  } 
                { 
                  var va_arg1 *Language.If
                  var va_arg2 *ClaireAny
                  va_arg1 = _CL_obj
                  { var _CL_obj *Language.Do = Language.To_Do(new(Language.Do).Is(Language.C_Do))
                    _CL_obj.Args = l1.Append(l2)
                    va_arg2 = _CL_obj.Id()
                    } 
                  va_arg1.Arg = va_arg2
                  } 
                arg_13 = _CL_obj.Id()
                } 
              } else {
              { var _CL_obj *Language.Let = Language.To_Let(new(Language.Let).Is(Language.C_Let))
                _CL_obj.ClaireVar = To_Variable(lvar.At(2))
                _CL_obj.Value = Language.F_readCall_relation(R,lvar.At(0)).Id()
                { 
                  var va_arg1 *Language.Let
                  var va_arg2 *ClaireAny
                  va_arg1 = _CL_obj
                  { var _CL_obj *Language.If = Language.To_If(new(Language.If).Is(Language.C_If))
                    { 
                      var va_arg1 *Language.If
                      var va_arg2 *ClaireAny
                      va_arg1 = _CL_obj
                      { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                        _CL_obj.Selector = ToProperty(Core.C__I_equal.Id())
                        _CL_obj.Args = MakeConstantList(lvar.At(1),lvar.At(2))
                        va_arg2 = _CL_obj.Id()
                        } 
                      va_arg1.Test = va_arg2
                      } 
                    { 
                      var va_arg1 *Language.If
                      var va_arg2 *ClaireAny
                      va_arg1 = _CL_obj
                      { var _CL_obj *Language.Do = Language.To_Do(new(Language.Do).Is(Language.C_Do))
                        _CL_obj.Args = l1.Append(l2)
                        va_arg2 = _CL_obj.Id()
                        } 
                      va_arg1.Arg = va_arg2
                      } 
                    va_arg2 = _CL_obj.Id()
                    } 
                  va_arg1.Arg = va_arg2
                  } 
                arg_13 = _CL_obj.Id()
                } 
              } 
            try_12 = Language.F_lambda_I_list(MakeConstantList(lvar.At(0),lvar.At(1)),arg_13)
            } 
          if ErrorIn(try_12) {Result = try_12
          } else {
          va_arg2 = ANY(try_12)
          va_arg1.IfWrite = va_arg2
          Result = va_arg2.ToEID()
          }
          } 
        }}}}}
        } 
      }
      } 
    } 
  } 
return Result} 

// The EID go function for: compile_if_write @ relation (throw: true) 
func E_Optimize_compile_if_write_relation (R EID) EID { 
return F_Optimize_compile_if_write_relation(ToRelation(OBJ(R)) )} 

// substitute 2 or 3 variables depending on lvar (CLAIRE4 : fixed an old CLAIRE3 bug that did not show)
/* The go function for: demon_substitution(x:Language/demon,lvar:list) [status=0] */
func F_Optimize_demon_substitution_demon (x *Language.LanguageDemon,lvar *ClaireList) *ClaireAny { 
var Result *ClaireAny
{ var y *ClaireAny
  if (lvar.Length() > 2) { 
    y = Language.F_substitution_any(x.Formula.Body,To_Variable(x.Formula.Vars.At(2)),lvar.At(2))
    } else {
    y = x.Formula.Body
    } 
  Result = Language.F_substitution_any(Language.F_substitution_any(y,To_Variable(x.Formula.Vars.At(0)),lvar.At(0)),To_Variable(x.Formula.Vars.At(1)),lvar.At(1))
  } 
return Result} 

// The EID go function for: demon_substitution @ Language/demon (throw: false) 
func E_Optimize_demon_substitution_demon (x EID,lvar EID) EID { 
return F_Optimize_demon_substitution_demon(Language.ToLanguageDemon(OBJ(x)),ToList(OBJ(lvar)) ).ToEID()} 

// create an arity 2 function  for demons
/* The go function for: demon_function(%name:string) [status=0] */
func F_Optimize_demon_function_string (_Zname *ClaireString) *ClaireFunction { 
var Result *ClaireFunction
{ var _Zf *ClaireFunction = F_make_function_string(_Zname)
  F_set_arity_function(_Zf,2)
  Result = _Zf
  } 
return Result} 

// The EID go function for: demon_function @ string (throw: false) 
func E_Optimize_demon_function_string (_Zname EID) EID { 
return F_Optimize_demon_function_string(ToString(OBJ(_Zname)) ).ToEID()} 

// create a simple method that will trigger the event
/* The go function for: compileEventMethod(p:property) [status=1] */
func F_Optimize_compileEventMethod_property (p *ClaireProperty) EID { 
var Result EID
{ var m *ClaireMethod = ToMethod(p.Restrictions.ValuesO()[0])
  { var na *ClaireString = F_append_string(p.Name.String_I(),MakeString("_write"))
    Result = F_Optimize_add_method_I_method(m,
      MakeConstantList(p.Domain.Id(),p.Range.Id()),
      C_void.Id(),
      MakeInteger(0).Id(),
      F_Optimize_demon_function_string(na))
    } 
  } 
return Result} 

// The EID go function for: compileEventMethod @ property (throw: true) 
func E_Optimize_compileEventMethod_property (p EID) EID { 
return F_Optimize_compileEventMethod_property(ToProperty(OBJ(p)) )} 
