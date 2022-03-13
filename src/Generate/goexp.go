/***** CLAIRE Compilation of file /Users/ycaseau/Dropbox/src/clairev4.03/src/compile/goexp.cl 
         [version 4.0.04 / safety 5] Sunday 03-13-2022 07:28:45 *****/

package Generate
import (_ "fmt"
	. "Kernel"
	"Core"
	"Language"
	"Reader"
	"Optimize"
)

//-------- dumb function to prevent import errors --------
func import_g0072() { 
    _ = Core.It
    _ = Language.It
    _ = Reader.It
    _ = Optimize.It
    } 
  
  
//+-------------------------------------------------------------+
//| CLAIRE                                                      |
//| goexp.cl                                                    |
//| Copyright (C) 2020 - 2021 Yves Caseau. All Rights Reserved  |
//| cf. copyright info in file object.cl: about()               |
//+-------------------------------------------------------------+
// ---------------------------------------------------------------------
// Compiling is based upon three methods:
//  - g_func? tests if the CLAIRE form can be represented by a C/ expression.
//    In this case,
//  - g_expression transforms it into an equivalent go expression.
//    otherwise,
//  - gstatement takes also a variable as an argument, and transforms a CLAIRE
//    expression into a C statement that assigns the value of the expression
//    into the variable;
//
// A special case occurs when the expression represent a boolean value and is
// functional, we can use bool_exp that returns a C boolean
// ---------------------------------------------------------------------
// *********************************************************************
// *  Contents                                                         *
// *  Part 1: g_func & expression for objects                          *
// *  Part 2: expression for messages                                  *
// *  Part 3: the inline coding of function calls                      *
// *  Part 4: expression for structures                                *
// *  Part 5: boolean optimization                                     *
// *********************************************************************
// g_expression(x:any,s:class) produces a go expression based on expected go type
//     s = EID                            => produce an EID
//     s = any, object, c                 => produces a *ClaireAny  representation (default case)
//     s = integer, char, float, string   => produced a native representation
//**********************************************************************
//*          Part 1: g_func & expression for objects                   *
//**********************************************************************
// this methods tells if a CLAIRE instruction can be compiled as an expression,as opposed to a statement.
// CHANGE in CLAIRE 4 : everything that may throw an exception needs a statement (because of go limitation)
// HOWEVER : if a call produces the possible error, it should simply be compiled in EID mode
/* The go function for: g_func(self:any) [status=1] */
func F_Generate_g_func_any (self *ClaireAny ) EID { 
    // eid body s = boolean
    var Result EID 
    if (self.Isa.IsIn(C_bag) == CTRUE) { 
      { var g0073 *ClaireBag   = ToBag(self)
        { var arg_1 *ClaireAny  
          var try_2 EID 
          { 
            var x *ClaireAny  
            _ = x
            try_2= EID{CFALSE.Id(),0}
            var x_support *ClaireList  
            var try_3 EID 
            try_3 = Core.F_enumerate_any(g0073.Id())
            if ErrorIn(try_3) {try_2 = try_3
            } else {
            x_support = ToList(OBJ(try_3))
            x_len := x_support.Length()
            for i_it := 0; i_it < x_len; i_it++ { 
              x = x_support.At(i_it)
              var loop_4 EID 
              _ = loop_4
              var g0087I *ClaireBoolean  
              var try_5 EID 
              { var arg_6 *ClaireBoolean  
                var try_7 EID 
                try_7 = F_Generate_g_func_any(x)
                if ErrorIn(try_7) {try_5 = try_7
                } else {
                arg_6 = ToBoolean(OBJ(try_7))
                try_5 = EID{arg_6.Not.Id(),0}
                }
                } 
              if ErrorIn(try_5) {loop_4 = try_5
              } else {
              g0087I = ToBoolean(OBJ(try_5))
              if (g0087I == CTRUE) { 
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
          Result = EID{Core.F_not_any(arg_1).Id(),0}
          }
          } 
        } 
      }  else if (self.Isa.IsIn(Language.C_Construct) == CTRUE) { 
      { var g0074 *Language.Construct   = Language.To_Construct(self)
        if (((g0074.Isa.IsIn(Language.C_Set) == CTRUE) || 
              (g0074.Isa.IsIn(Language.C_List) == CTRUE)) || 
            (g0074.Isa.IsIn(Language.C_Tuple) == CTRUE)) { 
          { 
            var v_and5 *ClaireBoolean  
            
            v_and5 = Core.F__inf_integer(g0074.Args.Length(),15)
            if (v_and5 == CFALSE) {Result = EID{CFALSE.Id(),0}
            } else { 
              var try_8 EID 
              { var arg_9 *ClaireAny  
                var try_10 EID 
                { 
                  var x *ClaireAny  
                  _ = x
                  try_10= EID{CFALSE.Id(),0}
                  var x_support *ClaireList  
                  x_support = g0074.Args
                  x_len := x_support.Length()
                  for i_it := 0; i_it < x_len; i_it++ { 
                    x = x_support.At(i_it)
                    var loop_11 EID 
                    _ = loop_11
                    var g0088I *ClaireBoolean  
                    var try_12 EID 
                    { var arg_13 *ClaireBoolean  
                      var try_14 EID 
                      try_14 = F_Generate_g_func_any(x)
                      if ErrorIn(try_14) {try_12 = try_14
                      } else {
                      arg_13 = ToBoolean(OBJ(try_14))
                      try_12 = EID{arg_13.Not.Id(),0}
                      }
                      } 
                    if ErrorIn(try_12) {loop_11 = try_12
                    } else {
                    g0088I = ToBoolean(OBJ(try_12))
                    if (g0088I == CTRUE) { 
                      try_10 = EID{CTRUE.Id(),0}
                      break
                      } else {
                      loop_11 = EID{CFALSE.Id(),0}
                      } 
                    }
                    if ErrorIn(loop_11) {try_10 = loop_11
                    break
                    } else {
                    }
                    } 
                  } 
                if ErrorIn(try_10) {try_8 = try_10
                } else {
                arg_9 = ANY(try_10)
                try_8 = EID{Core.F_not_any(arg_9).Id(),0}
                }
                } 
              if ErrorIn(try_8) {Result = try_8
              } else {
              v_and5 = ToBoolean(OBJ(try_8))
              if (v_and5 == CFALSE) {Result = EID{CFALSE.Id(),0}
              } else { 
                Result = EID{CTRUE.Id(),0}} 
              } 
            }
            } 
          } else {
          Result = EID{CFALSE.Id(),0}
          } 
        } 
      }  else if (self.Isa.IsIn(Language.C_If) == CTRUE) { 
      { var g0075 *Language.If   = Language.To_If(self)
        { 
          var v_and4 *ClaireBoolean  
          
          var try_15 EID 
          try_15 = F_Generate_g_func_any(g0075.Test)
          if ErrorIn(try_15) {Result = try_15
          } else {
          v_and4 = ToBoolean(OBJ(try_15))
          if (v_and4 == CFALSE) {Result = EID{CFALSE.Id(),0}
          } else { 
            v_and4 = F_Generate_constant_ask_any(g0075.Arg)
            if (v_and4 == CFALSE) {Result = EID{CFALSE.Id(),0}
            } else { 
              v_and4 = F_Generate_constant_ask_any(g0075.Other)
              if (v_and4 == CFALSE) {Result = EID{CFALSE.Id(),0}
              } else { 
                Result = EID{CTRUE.Id(),0}} 
              } 
            } 
          }
          } 
        } 
      }  else if (self.Isa.IsIn(Language.C_And) == CTRUE) { 
      { var g0076 *Language.And   = Language.To_And(self)
        Result = F_Generate_g_func_any(g0076.Args.Id())
        } 
      }  else if (self.Isa.IsIn(Language.C_Or) == CTRUE) { 
      { var g0077 *Language.Or   = Language.To_Or(self)
        Result = F_Generate_g_func_any(g0077.Args.Id())
        } 
      }  else if (self.Isa.IsIn(Language.C_Call) == CTRUE) { 
      { var g0078 *Language.Call   = Language.To_Call(self)
        { 
          var v_and4 *ClaireBoolean  
          
          var try_16 EID 
          try_16 = F_Generate_g_func_any(g0078.Args.Id())
          if ErrorIn(try_16) {Result = try_16
          } else {
          v_and4 = ToBoolean(OBJ(try_16))
          if (v_and4 == CFALSE) {Result = EID{CFALSE.Id(),0}
          } else { 
            v_and4 = Core.F__I_equal_any(g0078.Selector.Id(),Optimize.C_Compile_object_I.Id())
            if (v_and4 == CFALSE) {Result = EID{CFALSE.Id(),0}
            } else { 
              var try_17 EID 
              { var arg_18 *ClaireAny  
                var try_19 EID 
                { 
                  var x *ClaireAny  
                  _ = x
                  try_19= EID{CFALSE.Id(),0}
                  var x_support *ClaireList  
                  x_support = g0078.Args
                  x_len := x_support.Length()
                  for i_it := 0; i_it < x_len; i_it++ { 
                    x = x_support.At(i_it)
                    var loop_20 EID 
                    _ = loop_20
                    var g0089I *ClaireBoolean  
                    var try_21 EID 
                    { var arg_22 *ClaireBoolean  
                      var try_23 EID 
                      { var arg_24 *ClaireBoolean  
                        var try_25 EID 
                        try_25 = Optimize.F_Compile_g_throw_any(x)
                        if ErrorIn(try_25) {try_23 = try_25
                        } else {
                        arg_24 = ToBoolean(OBJ(try_25))
                        try_23 = EID{arg_24.Not.Id(),0}
                        }
                        } 
                      if ErrorIn(try_23) {try_21 = try_23
                      } else {
                      arg_22 = ToBoolean(OBJ(try_23))
                      try_21 = EID{arg_22.Not.Id(),0}
                      }
                      } 
                    if ErrorIn(try_21) {loop_20 = try_21
                    } else {
                    g0089I = ToBoolean(OBJ(try_21))
                    if (g0089I == CTRUE) { 
                      try_19 = EID{CTRUE.Id(),0}
                      break
                      } else {
                      loop_20 = EID{CFALSE.Id(),0}
                      } 
                    }
                    if ErrorIn(loop_20) {try_19 = loop_20
                    break
                    } else {
                    }
                    } 
                  } 
                if ErrorIn(try_19) {try_17 = try_19
                } else {
                arg_18 = ANY(try_19)
                try_17 = EID{Core.F_not_any(arg_18).Id(),0}
                }
                } 
              if ErrorIn(try_17) {Result = try_17
              } else {
              v_and4 = ToBoolean(OBJ(try_17))
              if (v_and4 == CFALSE) {Result = EID{CFALSE.Id(),0}
              } else { 
                Result = EID{CTRUE.Id(),0}} 
              } 
            } 
          }}
          } 
        } 
      }  else if (self.Isa.IsIn(Language.C_Super) == CTRUE) { 
      { var g0079 *Language.Super   = Language.To_Super(self)
        Result = F_Generate_g_func_any(g0079.Args.Id())
        } 
      }  else if (self.Isa.IsIn(Language.C_Call_method) == CTRUE) { 
      { var g0080 *Language.CallMethod   = Language.To_CallMethod(self)
        { 
          var v_and4 *ClaireBoolean  
          
          var try_26 EID 
          try_26 = F_Generate_g_func_any(g0080.Args.Id())
          if ErrorIn(try_26) {Result = try_26
          } else {
          v_and4 = ToBoolean(OBJ(try_26))
          if (v_and4 == CFALSE) {Result = EID{CFALSE.Id(),0}
          } else { 
            var try_27 EID 
            { 
              var v_or6 *ClaireBoolean  
              
              v_or6 = Equal(g0080.Arg.Id(),Optimize.C_Compile_m_unsafe.Value)
              if (v_or6 == CTRUE) {try_27 = EID{CTRUE.Id(),0}
              } else { 
                var try_28 EID 
                { var arg_29 *ClaireAny  
                  var try_30 EID 
                  { 
                    var x *ClaireAny  
                    _ = x
                    try_30= EID{CFALSE.Id(),0}
                    var x_support *ClaireList  
                    x_support = g0080.Args
                    x_len := x_support.Length()
                    for i_it := 0; i_it < x_len; i_it++ { 
                      x = x_support.At(i_it)
                      var loop_31 EID 
                      _ = loop_31
                      var g0090I *ClaireBoolean  
                      var try_32 EID 
                      { var arg_33 *ClaireBoolean  
                        var try_34 EID 
                        { var arg_35 *ClaireBoolean  
                          var try_36 EID 
                          try_36 = Optimize.F_Compile_g_throw_any(x)
                          if ErrorIn(try_36) {try_34 = try_36
                          } else {
                          arg_35 = ToBoolean(OBJ(try_36))
                          try_34 = EID{arg_35.Not.Id(),0}
                          }
                          } 
                        if ErrorIn(try_34) {try_32 = try_34
                        } else {
                        arg_33 = ToBoolean(OBJ(try_34))
                        try_32 = EID{arg_33.Not.Id(),0}
                        }
                        } 
                      if ErrorIn(try_32) {loop_31 = try_32
                      } else {
                      g0090I = ToBoolean(OBJ(try_32))
                      if (g0090I == CTRUE) { 
                        try_30 = EID{CTRUE.Id(),0}
                        break
                        } else {
                        loop_31 = EID{CFALSE.Id(),0}
                        } 
                      }
                      if ErrorIn(loop_31) {try_30 = loop_31
                      break
                      } else {
                      }
                      } 
                    } 
                  if ErrorIn(try_30) {try_28 = try_30
                  } else {
                  arg_29 = ANY(try_30)
                  try_28 = EID{Core.F_not_any(arg_29).Id(),0}
                  }
                  } 
                if ErrorIn(try_28) {try_27 = try_28
                } else {
                v_or6 = ToBoolean(OBJ(try_28))
                if (v_or6 == CTRUE) {try_27 = EID{CTRUE.Id(),0}
                } else { 
                  try_27 = EID{CFALSE.Id(),0}} 
                } 
              }
              } 
            if ErrorIn(try_27) {Result = try_27
            } else {
            v_and4 = ToBoolean(OBJ(try_27))
            if (v_and4 == CFALSE) {Result = EID{CFALSE.Id(),0}
            } else { 
              Result = EID{CTRUE.Id(),0}} 
            } 
          }}
          } 
        } 
      }  else if (self.Isa.IsIn(Language.C_Call_slot) == CTRUE) { 
      { var g0081 *Language.CallSlot   = Language.To_CallSlot(self)
        Result = F_Generate_g_func_any(g0081.Arg)
        } 
      }  else if (self.Isa.IsIn(Language.C_Call_table) == CTRUE) { 
      { var g0082 *Language.CallTable   = Language.To_CallTable(self)
        Result = F_Generate_g_func_any(g0082.Arg)
        } 
      }  else if (self.Isa.IsIn(Language.C_Call_array) == CTRUE) { 
      { var g0083 *Language.CallArray   = Language.To_CallArray(self)
        Result = F_Generate_g_func_any(g0083.Arg)
        } 
      }  else if (self.Isa.IsIn(Language.C_Cast) == CTRUE) { 
      { var g0084 *Language.Cast   = Language.To_Cast(self)
        Result = F_Generate_g_func_any(g0084.Arg)
        } 
      }  else if (self.Isa.IsIn(Optimize.C_Compile_C_cast) == CTRUE) { 
      { var g0085 *Optimize.Compile_CCast   = Optimize.To_Compile_CCast(self)
        Result = F_Generate_g_func_any(g0085.Arg)
        } 
      } else {
      Result = EID{MakeBoolean((self.Isa.IsIn(C_thing) == CTRUE) || 
      (C_integer.Id() == self.Isa.Id()) || 
      (C_string.Id() == self.Isa.Id()) || 
      (C_char.Id() == self.Isa.Id()) || 
      (self.Isa.IsIn(C_lambda) == CTRUE) || 
      (C_float.Id() == self.Isa.Id()) || 
      (self.Isa.IsIn(C_Variable) == CTRUE) || 
      (self.Isa.IsIn(Core.C_global_variable) == CTRUE) || 
      (C_function.Id() == self.Isa.Id()) || 
      (self.Isa.IsIn(C_symbol) == CTRUE) || 
      (self == CNULL) || 
      (C_method.Id() == self.Isa.Id()) || 
      (C_boolean.Id() == self.Isa.Id()) || 
      (C_class.Id() == self.Isa.Id()) || 
      (self.Isa.IsIn(C_environment) == CTRUE)).Id(),0}
      } 
    return Result} 
  
// The EID go function for: g_func @ any (throw: true) 
func E_Generate_g_func_any (self EID) EID { 
    return F_Generate_g_func_any(ANY(self) )} 
  
// manages unknown + catch-all 
/* The go function for: g_expression(self:any,s:class) [status=1] */
func F_Generate_g_expression_any (self *ClaireAny ,s *ClaireClass ) EID { 
    // eid body s = void
    var Result EID 
    if (self != CNULL) { 
      Result = ToException(Core.C_general_error.Make(MakeString("/!\\ design error: g_expression(~S: ~S) unknown").Id(),MakeConstantList(self,self.Isa.Id()).Id())).Close()
      }  else if (s.Id() == Optimize.C_EID.Id()) { 
      Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).ToEid(self,C_object)
      }  else if (s.Id() == C_any.Id()) { 
      PRINC("CNULL")
      Result = EVOID
      } else {
      F_Generate_object_prefix_class(C_any,s)
      PRINC("CNULL")
      F_Generate_object_post_class(C_any,s)
      Result = EVOID
      } 
    return Result} 
  
// The EID go function for: g_expression @ any (throw: true) 
func E_Generate_g_expression_any (self EID,s EID) EID { 
    return F_Generate_g_expression_any(ANY(self),ToClass(OBJ(s)) )} 
  
// Things are represented by global variables in the associated go package
/* The go function for: g_expression(self:thing,s:class) [status=1] */
func F_Generate_g_expression_thing (self *ClaireThing ,s *ClaireClass ) EID { 
    // eid body s = void
    var Result EID 
    if (s.Id() == Optimize.C_EID.Id()) { 
      Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).ToEid(self.Id(),C_object)
      } else {
      F_Generate_object_prefix_class(self.Id().Isa,s)
      F_Generate_thing_ident_thing(self)
      F_Generate_object_post_class(self.Id().Isa,s)
      Result = EVOID
      } 
    return Result} 
  
// The EID go function for: g_expression @ thing (throw: true) 
func E_Generate_g_expression_thing (self EID,s EID) EID { 
    return F_Generate_g_expression_thing(ToThing(OBJ(self)),ToClass(OBJ(s)) )} 
  
// note that there are two kinds of modules
//    - packages (when m.made_of != nil)  -> defined in their first members (iClaire in Language)
//    - node modules (abstractions) => need to be attached to packages
/* The go function for: g_expression(self:module,s:class) [status=1] */
func F_Generate_g_expression_module (self *ClaireModule ,s *ClaireClass ) EID { 
    // eid body s = void
    var Result EID 
    if (s.Id() == Optimize.C_EID.Id()) { 
      Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).ToEid(self.Id(),C_object)
      } else {
      F_Generate_object_prefix_class(self.Id().Isa,s)
      if (self.Id() == ToGenerateGoProducer(Optimize.C_PRODUCER.Value).Current.Id()) { 
        PRINC("It")
        }  else if (self.Id() == C_Kernel.Id()) { 
        PRINC("C_Kernel")
        }  else if (Equal(self.MadeOf.Id(),CNIL.Id()) == CTRUE) { 
        { var m *ClaireModule   = F_Generate_get_made_module(self)
          if ((m.Id() != C_Kernel.Id()) && 
              (m.Id() != ToGenerateGoProducer(Optimize.C_PRODUCER.Value).Current.Id())) { 
            F_iClaire_ident_symbol(m.Name)
            PRINC(".")
            } 
          F_Generate_go_var_symbol(self.Name)
          } 
        } else {
        F_iClaire_ident_symbol(self.Name)
        PRINC(".It")
        } 
      F_Generate_object_post_class(self.Id().Isa,s)
      Result = EVOID
      } 
    return Result} 
  
// The EID go function for: g_expression @ module (throw: true) 
func E_Generate_g_expression_module (self EID,s EID) EID { 
    return F_Generate_g_expression_module(ToModule(OBJ(self)),ToClass(OBJ(s)) )} 
  
// A class is similar to a thing
/* The go function for: g_expression(self:class,s:class) [status=1] */
func F_Generate_g_expression_class (self *ClaireClass ,s *ClaireClass ) EID { 
    // eid body s = void
    var Result EID 
    if (s.Id() == Optimize.C_EID.Id()) { 
      Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).ToEid(self.Id(),C_object)
      } else {
      F_Generate_object_prefix_class(C_class,s)
      F_Generate_class_ident_class(self)
      F_Generate_object_post_class(C_class,s)
      Result = EVOID
      } 
    return Result} 
  
// The EID go function for: g_expression @ class (throw: true) 
func E_Generate_g_expression_class (self EID,s EID) EID { 
    return F_Generate_g_expression_class(ToClass(OBJ(self)),ToClass(OBJ(s)) )} 
  
// A named object is designed by a C identifier !
/* The go function for: g_expression(self:boolean,s:class) [status=1] */
func F_Generate_g_expression_boolean (self *ClaireBoolean ,s *ClaireClass ) EID { 
    // eid body s = void
    var Result EID 
    if (s.Id() == Optimize.C_EID.Id()) { 
      Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).ToEid(self.Id(),C_object)
      } else {
      F_Generate_object_prefix_class(C_boolean,s)
      F_princ_string(ToString(IfThenElse((self == CTRUE),
        MakeString("CTRUE").Id(),
        MakeString("CFALSE").Id())))
      F_Generate_object_post_class(C_boolean,s)
      Result = EVOID
      } 
    return Result} 
  
// The EID go function for: g_expression @ boolean (throw: true) 
func E_Generate_g_expression_boolean (self EID,s EID) EID { 
    return F_Generate_g_expression_boolean(ToBoolean(OBJ(self)),ToClass(OBJ(s)) )} 
  
// Primitive types rely on the producer to generate code that uses their specific implementation
// this is done on purpose: supports the customization through another producer
/* The go function for: g_expression(self:integer,s:class) [status=1] */
func F_Generate_g_expression_integer (self int,s *ClaireClass ) EID { 
    // eid body s = void
    var Result EID 
    if (s.Id() == Optimize.C_EID.Id()) { 
      Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).ToEid(MakeInteger(self).Id(),C_integer)
      }  else if (s.Id() == C_integer.Id()) { 
      F_princ_integer(self)
      Result = EVOID
      } else {
      F_Generate_object_prefix_class(C_integer,s)
      Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).ToCl(MakeInteger(self).Id(),C_integer)
      if !ErrorIn(Result) {
      F_Generate_object_post_class(C_integer,s)
      Result = EVOID
      }
      } 
    return Result} 
  
// The EID go function for: g_expression @ integer (throw: true) 
func E_Generate_g_expression_integer (self EID,s EID) EID { 
    return F_Generate_g_expression_integer(INT(self),ToClass(OBJ(s)) )} 
  
/* The go function for: g_expression(self:float,s:class) [status=1] */
func F_Generate_g_expression_float (self float64,s *ClaireClass ) EID { 
    // eid body s = void
    var Result EID 
    if (s.Id() == Optimize.C_EID.Id()) { 
      Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).ToEid(MakeFloat(self).Id(),C_float)
      }  else if (s.Id() == C_float.Id()) { 
      F_princ_float(self)
      Result = EVOID
      } else {
      F_Generate_object_prefix_class(C_float,s)
      Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).ToCl(MakeFloat(self).Id(),C_float)
      if !ErrorIn(Result) {
      F_Generate_object_post_class(C_float,s)
      Result = EVOID
      }
      } 
    return Result} 
  
// The EID go function for: g_expression @ float (throw: true) 
func E_Generate_g_expression_float (self EID,s EID) EID { 
    return F_Generate_g_expression_float(FLOAT(self),ToClass(OBJ(s)) )} 
  
/* The go function for: g_expression(self:char,s:class) [status=1] */
func F_Generate_g_expression_char (self rune,s *ClaireClass ) EID { 
    // eid body s = void
    var Result EID 
    if (s.Id() == Optimize.C_EID.Id()) { 
      Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).ToEid(MakeChar(self).Id(),C_char)
      }  else if (s.Id() == C_char.Id()) { 
      Result = Core.F_print_any(MakeChar(self).Id())
      } else {
      F_Generate_object_prefix_class(C_char,s)
      Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).ToCl(MakeChar(self).Id(),C_char)
      if !ErrorIn(Result) {
      F_Generate_object_post_class(C_char,s)
      Result = EVOID
      }
      } 
    return Result} 
  
// The EID go function for: g_expression @ char (throw: true) 
func E_Generate_g_expression_char (self EID,s EID) EID { 
    return F_Generate_g_expression_char(CHAR(self),ToClass(OBJ(s)) )} 
  
// strings are primitive objects, same as function
/* The go function for: g_expression(self:string,s:class) [status=1] */
func F_Generate_g_expression_string (self *ClaireString ,s *ClaireClass ) EID { 
    // eid body s = void
    var Result EID 
    if (s.Id() == Optimize.C_EID.Id()) { 
      Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).ToEid((self).Id(),C_string)
      } else {
      F_Generate_object_prefix_class(C_string,s)
      PRINC("MakeString(")
      Result = Core.F_print_any((self).Id())
      if !ErrorIn(Result) {
      PRINC(")")
      F_Generate_object_post_class(C_string,s)
      Result = EVOID
      }
      } 
    return Result} 
  
// The EID go function for: g_expression @ string (throw: true) 
func E_Generate_g_expression_string (self EID,s EID) EID { 
    return F_Generate_g_expression_string(ToString(OBJ(self)),ToClass(OBJ(s)) )} 
  
// symboles are primitive objects, same as function
/* The go function for: g_expression(self:symbol,s:class) [status=1] */
func F_Generate_g_expression_symbol (self *ClaireSymbol ,s *ClaireClass ) EID { 
    // eid body s = void
    var Result EID 
    if (s.Id() == Optimize.C_EID.Id()) { 
      Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).ToEid(self.Id(),C_object)
      } else {
      F_Generate_object_prefix_class(C_object,s)
      PRINC("MakeSymbol(")
      Result = Core.F_print_any((self.String_I()).Id())
      if !ErrorIn(Result) {
      PRINC(",")
      Result = F_Generate_g_expression_module(self.Module_I(),C_module)
      if !ErrorIn(Result) {
      PRINC(")")
      F_Generate_object_post_class(C_object,s)
      Result = EVOID
      }}
      } 
    return Result} 
  
// The EID go function for: g_expression @ symbol (throw: true) 
func E_Generate_g_expression_symbol (self EID,s EID) EID { 
    return F_Generate_g_expression_symbol(ToSymbol(OBJ(self)),ToClass(OBJ(s)) )} 
  
/* The go function for: g_expression(self:environment,s:class) [status=1] */
func F_Generate_g_expression_environment (self *ClaireEnvironment ,s *ClaireClass ) EID { 
    // eid body s = void
    var Result EID 
    if (s.Id() == Optimize.C_EID.Id()) { 
      Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).ToEid(self.Id(),C_object)
      } else {
      F_Generate_object_prefix_class(C_environment,s)
      PRINC("ClEnv")
      F_Generate_object_post_class(C_environment,s)
      Result = EVOID
      } 
    return Result} 
  
// The EID go function for: g_expression @ environment (throw: true) 
func E_Generate_g_expression_environment (self EID,s EID) EID { 
    return F_Generate_g_expression_environment(ToEnvironment(OBJ(self)),ToClass(OBJ(s)) )} 
  
/* The go function for: g_expression(self:function,s:class) [status=1] */
func F_Generate_g_expression_function (self *ClaireFunction ,s *ClaireClass ) EID { 
    // eid body s = void
    var Result EID 
    if (s.Id() == Optimize.C_EID.Id()) { 
      Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).ToEid(self.Id(),C_object)
      } else {
      F_Generate_object_prefix_class(C_function,s)
      PRINC("MakeFunction")
      F_princ_integer(F_arity_function(self))
      PRINC("(E_")
      F_c_princ_function(self)
      PRINC(",")
      Result = Core.F_print_any((F_string_I_function(self)).Id())
      if !ErrorIn(Result) {
      PRINC(")")
      F_Generate_object_post_class(C_function,s)
      Result = EVOID
      }
      } 
    return Result} 
  
// The EID go function for: g_expression @ function (throw: true) 
func E_Generate_g_expression_function (self EID,s EID) EID { 
    return F_Generate_g_expression_function(ToFunction(OBJ(self)),ToClass(OBJ(s)) )} 
  
// lexical variables are represented by C variables
// notice that we may need native to object conversion
/* The go function for: g_expression(self:Variable,s:class) [status=1] */
func F_Generate_g_expression_Variable (self *ClaireVariable ,s *ClaireClass ) EID { 
    // eid body s = void
    var Result EID 
    
    { var s2 *ClaireClass   = self.Range.Class_I()
      if (s.Id() == Optimize.C_EID.Id()) { 
        if (s2.Id() == Optimize.C_EID.Id()) { 
          F_iClaire_ident_go_producer1(ToGenerateGoProducer(Optimize.C_PRODUCER.Value),self)
          Result = EVOID
          } else {
          Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).ToEid(self.Id(),s2)
          } 
        } else {
        Result = F_Generate_cast_prefix_class(s2,s)
        if !ErrorIn(Result) {
        F_iClaire_ident_go_producer1(ToGenerateGoProducer(Optimize.C_PRODUCER.Value),self)
        F_Generate_cast_post_class(s2,s)
        Result = EVOID
        }
        } 
      } 
    return Result} 
  
// The EID go function for: g_expression @ Variable (throw: true) 
func E_Generate_g_expression_Variable (self EID,s EID) EID { 
    return F_Generate_g_expression_Variable(To_Variable(OBJ(self)),ToClass(OBJ(s)) )} 
  
// global_variables are CLAIRE objects
// v4.0.4 : handle optimized variables (nativeVarG)
/* The go function for: g_expression(self:global_variable,s:class) [status=1] */
func F_Generate_g_expression_global_variable (self *Core.GlobalVariable ,s *ClaireClass ) EID { 
    // eid body s = void
    var Result EID 
    if (s.Id() == Optimize.C_EID.Id()) { 
      Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).ToEid(self.Id(),C_object)
      }  else if ((Equal(self.Range.Id(),CEMPTY.Id()) == CTRUE) && 
        ((C_integer.Id() == self.Value.Isa.Id()) || 
            ((C_float.Id() == self.Value.Isa.Id()) || 
              (Equal(self.Value,CNIL.Id()) == CTRUE)))) { 
      Result = Core.F_CALL(C_Generate_g_expression,ARGS(self.Value.ToEID(),EID{s.Id(),0}))
      } else {
      { var s2 *ClaireClass  
        if (Optimize.F_Compile_nativeVar_ask_global_variable(self) == CTRUE) { 
          s2 = F_Generate_getRange_global_variable(self)
          } else {
          s2 = C_any
          } 
        Result = F_Generate_cast_prefix_class(s2,s)
        if !ErrorIn(Result) {
        ToGenerateGoProducer(Optimize.C_PRODUCER.Value).GlobalVar(self)
        F_Generate_cast_post_class(s2,s)
        Result = EVOID
        }
        } 
      } 
    return Result} 
  
// The EID go function for: g_expression @ global_variable (throw: true) 
func E_Generate_g_expression_global_variable (self EID,s EID) EID { 
    return F_Generate_g_expression_global_variable(Core.ToGlobalVariable(OBJ(self)),ToClass(OBJ(s)) )} 
  
// builds a set
/* The go function for: g_expression(self:Set,s:class) [status=1] */
func F_Generate_g_expression_Set (self *Language.Set ,s *ClaireClass ) EID { 
    // eid body s = void
    var Result EID 
    if (s.Id() == Optimize.C_EID.Id()) { 
      Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).ToEid(self.Id(),C_object)
      } else {
      Result = F_Generate_cast_prefix_class(C_set,s)
      if !ErrorIn(Result) {
      { var arg_1 *ClaireType  
        if (self.Of.Id() != CNULL) { 
          arg_1 = self.Of
          } else {
          arg_1 = ToType(CEMPTY.Id())
          } 
        Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).BagExpression(C_set,self.Args,arg_1)
        } 
      if !ErrorIn(Result) {
      F_Generate_cast_post_class(C_set,s)
      Result = EVOID
      }}
      } 
    return Result} 
  
// The EID go function for: g_expression @ Set (throw: true) 
func E_Generate_g_expression_Set (self EID,s EID) EID { 
    return F_Generate_g_expression_Set(Language.To_Set(OBJ(self)),ToClass(OBJ(s)) )} 
  
/* The go function for: g_expression(self:set,s:class) [status=1] */
func F_Generate_g_expression_set (self *ClaireSet ,s *ClaireClass ) EID { 
    // eid body s = void
    var Result EID 
    if (s.Id() == Optimize.C_EID.Id()) { 
      Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).ToEid(self.Id(),C_object)
      }  else if ((self.Size() == 0) && 
        (Equal(ToList(self.Id()).Of().Id(),CEMPTY.Id()) == CTRUE)) { 
      Result = F_Generate_cast_prefix_class(C_set,s)
      if !ErrorIn(Result) {
      PRINC("CEMPTY")
      F_Generate_cast_post_class(C_set,s)
      Result = EVOID
      }
      } else {
      Result = F_Generate_cast_prefix_class(C_set,s)
      if !ErrorIn(Result) {
      Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).BagExpression(C_set,self.List_I(),ToList(self.Id()).Of())
      if !ErrorIn(Result) {
      F_Generate_cast_post_class(C_set,s)
      Result = EVOID
      }}
      } 
    return Result} 
  
// The EID go function for: g_expression @ set (throw: true) 
func E_Generate_g_expression_set (self EID,s EID) EID { 
    return F_Generate_g_expression_set(ToSet(OBJ(self)),ToClass(OBJ(s)) )} 
  
// builds a tuple
/* The go function for: g_expression(self:Tuple,s:class) [status=1] */
func F_Generate_g_expression_Tuple (self *Language.Tuple ,s *ClaireClass ) EID { 
    // eid body s = void
    var Result EID 
    if (s.Id() == Optimize.C_EID.Id()) { 
      Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).ToEid(self.Id(),C_object)
      } else {
      Result = F_Generate_cast_prefix_class(C_tuple,s)
      if !ErrorIn(Result) {
      Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).BagExpression(C_tuple,self.Args,ToType(CEMPTY.Id()))
      if !ErrorIn(Result) {
      F_Generate_cast_post_class(C_tuple,s)
      Result = EVOID
      }}
      } 
    return Result} 
  
// The EID go function for: g_expression @ Tuple (throw: true) 
func E_Generate_g_expression_Tuple (self EID,s EID) EID { 
    return F_Generate_g_expression_Tuple(Language.To_Tuple(OBJ(self)),ToClass(OBJ(s)) )} 
  
/* The go function for: g_expression(self:tuple,s:class) [status=1] */
func F_Generate_g_expression_tuple (self *ClaireTuple ,s *ClaireClass ) EID { 
    // eid body s = void
    var Result EID 
    if (s.Id() == Optimize.C_EID.Id()) { 
      Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).ToEid(self.Id(),C_object)
      } else {
      Result = F_Generate_cast_prefix_class(C_tuple,s)
      if !ErrorIn(Result) {
      Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).BagExpression(C_tuple,self.List_I(),ToType(CEMPTY.Id()))
      if !ErrorIn(Result) {
      F_Generate_cast_post_class(C_tuple,s)
      Result = EVOID
      }}
      } 
    return Result} 
  
// The EID go function for: g_expression @ tuple (throw: true) 
func E_Generate_g_expression_tuple (self EID,s EID) EID { 
    return F_Generate_g_expression_tuple(ToTuple(OBJ(self)),ToClass(OBJ(s)) )} 
  
// builds a list
/* The go function for: g_expression(self:List,s:class) [status=1] */
func F_Generate_g_expression_List (self *Language.List ,s *ClaireClass ) EID { 
    // eid body s = void
    var Result EID 
    if (s.Id() == Optimize.C_EID.Id()) { 
      Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).ToEid(self.Id(),C_object)
      } else {
      Result = F_Generate_cast_prefix_class(C_list,s)
      if !ErrorIn(Result) {
      { var arg_1 *ClaireType  
        if (self.Of.Id() != CNULL) { 
          arg_1 = self.Of
          } else {
          arg_1 = ToType(CEMPTY.Id())
          } 
        Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).BagExpression(C_list,self.Args,arg_1)
        } 
      if !ErrorIn(Result) {
      F_Generate_cast_post_class(C_list,s)
      Result = EVOID
      }}
      } 
    return Result} 
  
// The EID go function for: g_expression @ List (throw: true) 
func E_Generate_g_expression_List (self EID,s EID) EID { 
    return F_Generate_g_expression_List(Language.To_List(OBJ(self)),ToClass(OBJ(s)) )} 
  
/* The go function for: g_expression(self:list,s:class) [status=1] */
func F_Generate_g_expression_list (self *ClaireList ,s *ClaireClass ) EID { 
    // eid body s = void
    var Result EID 
    if (s.Id() == Optimize.C_EID.Id()) { 
      Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).ToEid(self.Id(),C_object)
      }  else if ((self.Length() == 0) && 
        (Equal(self.Of().Id(),CEMPTY.Id()) == CTRUE)) { 
      Result = F_Generate_cast_prefix_class(C_list,s)
      if !ErrorIn(Result) {
      PRINC("CNIL")
      F_Generate_cast_post_class(C_list,s)
      Result = EVOID
      }
      } else {
      Result = F_Generate_cast_prefix_class(C_list,s)
      if !ErrorIn(Result) {
      Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).BagExpression(C_list,self,self.Of())
      if !ErrorIn(Result) {
      F_Generate_cast_post_class(C_list,s)
      Result = EVOID
      }}
      } 
    return Result} 
  
// The EID go function for: g_expression @ list (throw: true) 
func E_Generate_g_expression_list (self EID,s EID) EID { 
    return F_Generate_g_expression_list(ToList(OBJ(self)),ToClass(OBJ(s)) )} 
  
// new in CLAIRE 4 !! compilation of lambda is OK but requires the reader (similar to macros)
/* The go function for: g_expression(self:lambda,s:class) [status=1] */
func F_Generate_g_expression_lambda (self *ClaireLambda ,s *ClaireClass ) EID { 
    // eid body s = void
    var Result EID 
    Optimize.F_Optimize_legal_ask_module(Reader.It,self.Id())
    Result = F_Generate_eid_prefix_class(s)
    if !ErrorIn(Result) {
    PRINC("Core.F_read_lambda_string(MakeString(\"lambda[(")
    Result = Language.F_ppvariable_list(self.Vars)
    if !ErrorIn(Result) {
    PRINC("),")
    Result = Core.F_CALL(C_print,ARGS(self.Body.ToEID()))
    if !ErrorIn(Result) {
    PRINC("]\"))")
    F_Generate_eid_post_class(s)
    Result = EVOID
    }}}
    return Result} 
  
// The EID go function for: g_expression @ lambda (throw: true) 
func E_Generate_g_expression_lambda (self EID,s EID) EID { 
    return F_Generate_g_expression_lambda(ToLambda(OBJ(self)),ToClass(OBJ(s)) )} 
  
//**********************************************************************
//*          Part 2: expression for messages                         *
//**********************************************************************
// message compiling is tricky in go : Calls produce EID but for inline, Call_method produce native forms
// calls are expected to produce an EID
/* The go function for: g_expression(self:Call,s:class) [status=1] */
func F_Generate_g_expression_Call (self *Language.Call ,s *ClaireClass ) EID { 
    // eid body s = void
    var Result EID 
    Result = F_Generate_inline_exp_go_producer1(ToGenerateGoProducer(Optimize.C_PRODUCER.Value),self,s)
    return Result} 
  
// The EID go function for: g_expression @ Call (throw: true) 
func E_Generate_g_expression_Call (self EID,s EID) EID { 
    return F_Generate_g_expression_Call(Language.To_Call(OBJ(self)),ToClass(OBJ(s)) )} 
  
// the other cases will be taken care in the optimization part
/* The go function for: g_expression(self:Call_method1,s:class) [status=1] */
func F_Generate_g_expression_Call_method1 (self *Language.CallMethod1 ,s *ClaireClass ) EID { 
    // eid body s = void
    var Result EID 
    Result = Core.F_CALL(C_Generate_inline_exp,ARGS(EID{Optimize.C_PRODUCER.Value,0},EID{self.Id(),0},EID{s.Id(),0}))
    return Result} 
  
// The EID go function for: g_expression @ Call_method1 (throw: true) 
func E_Generate_g_expression_Call_method1 (self EID,s EID) EID { 
    return F_Generate_g_expression_Call_method1(Language.To_CallMethod1(OBJ(self)),ToClass(OBJ(s)) )} 
  
/* The go function for: g_expression(self:Call_method2,s:class) [status=1] */
func F_Generate_g_expression_Call_method2 (self *Language.CallMethod2 ,s *ClaireClass ) EID { 
    // eid body s = void
    var Result EID 
    Result = Core.F_CALL(C_Generate_inline_exp,ARGS(EID{Optimize.C_PRODUCER.Value,0},EID{self.Id(),0},EID{s.Id(),0}))
    return Result} 
  
// The EID go function for: g_expression @ Call_method2 (throw: true) 
func E_Generate_g_expression_Call_method2 (self EID,s EID) EID { 
    return F_Generate_g_expression_Call_method2(Language.To_CallMethod2(OBJ(self)),ToClass(OBJ(s)) )} 
  
/* The go function for: g_expression(self:Call_method,s:class) [status=1] */
func F_Generate_g_expression_Call_method (self *Language.CallMethod ,s *ClaireClass ) EID { 
    // eid body s = void
    var Result EID 
    Result = Core.F_CALL(C_Generate_inline_exp,ARGS(EID{Optimize.C_PRODUCER.Value,0},EID{self.Id(),0},EID{s.Id(),0}))
    return Result} 
  
// The EID go function for: g_expression @ Call_method (throw: true) 
func E_Generate_g_expression_Call_method (self EID,s EID) EID { 
    return F_Generate_g_expression_Call_method(Language.To_CallMethod(OBJ(self)),ToClass(OBJ(s)) )} 
  
// ---------------------------------------- dynamic call -------------------------------------------------------------------
// new in 3.0 : really low level method are virtual and only rely on inline compiling
// note the *_prefix(s) ... *_postfix(s) that add a conversion from * to exprected type s
// WARNING : we can use assignment (x = y) only when s = void (we do not care for the result)
/* The go function for: inline_exp(c:go_producer,self:Call,s:class) [status=1] */
func F_Generate_inline_exp_go_producer1 (c *GenerateGoProducer ,self *Language.Call ,s *ClaireClass ) EID { 
    // eid body s = void
    var Result EID 
    { var p *ClaireProperty   = self.Selector
      { var a1 *ClaireAny  
        var try_1 EID 
        try_1 = Core.F_car_list(self.Args)
        if ErrorIn(try_1) {Result = try_1
        } else {
        a1 = ANY(try_1)
        { var n int  = self.Args.Length()
          if (p.Id() == Core.C_mClaire_get_stack.Id()) { 
            Result = F_Generate_eid_prefix_class(s)
            if !ErrorIn(Result) {
            PRINC("ClEnv.EvalStack[")
            Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_integer.Id(),0}))
            if !ErrorIn(Result) {
            PRINC("]")
            F_Generate_eid_post_class(s)
            Result = EVOID
            }}
            }  else if (p.Id() == Optimize.C_safe.Id()) { 
            { var y int  = Optimize.C_compiler.Safety
              Optimize.C_compiler.Safety = 1
              Result = Core.F_CALL(C_Generate_g_expression,ARGS(self.Args.At(1-1).ToEID(),EID{s.Id(),0}))
              if !ErrorIn(Result) {
              { 
                var va_arg1 *Optimize.OptimizeMetaCompiler  
                var va_arg2 int 
                va_arg1 = Optimize.C_compiler
                va_arg2 = y
                va_arg1.Safety = va_arg2
                Result = EID{C__INT,IVAL(va_arg2)}
                } 
              }
              } 
            }  else if (p.Id() == Core.C_mClaire_base_I.Id()) { 
            F_Generate_integer_prefix_class(s)
            PRINC("ClEnv.Base")
            F_Generate_native_post_class(s)
            Result = EVOID
            }  else if (p.Id() == Core.C_Core__inf_equalt.Id()) { 
            F_Generate_object_prefix_class(C_boolean,s)
            Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_type.Id(),0}))
            if !ErrorIn(Result) {
            PRINC(".Included(")
            Result = Core.F_CALL(C_Generate_g_expression,ARGS(self.Args.At(2-1).ToEID(),EID{C_type.Id(),0}))
            if !ErrorIn(Result) {
            PRINC(")")
            F_Generate_object_post_class(C_boolean,s)
            Result = EVOID
            }}
            }  else if ((p.Id() == Core.C_mClaire_index_I.Id()) && 
              (n == 1)) { 
            F_Generate_integer_prefix_class(s)
            PRINC("ClEnv.Index")
            F_Generate_native_post_class(s)
            Result = EVOID
            }  else if ((p.Id() == Core.C_mClaire_push_I.Id()) && 
              (n == 1)) { 
            PRINC("ClEnv.Push(")
            Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{Optimize.C_EID.Id(),0}))
            if !ErrorIn(Result) {
            PRINC(")")
            Result = EVOID
            }
            }  else if (p.Id() == Core.C_mClaire_put_stack.Id()) { 
            if (s.Id() != C_void.Id()) { 
              Optimize.F_Compile_warn_void()
              Core.F_tformat_string(MakeString("use ~S in non void context\n"),1,MakeConstantList(self.Id()))
              } 
            PRINC("ClEnv.EvalStack[")
            Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_integer.Id(),0}))
            if !ErrorIn(Result) {
            PRINC("]=")
            Result = Core.F_CALL(C_Generate_g_expression,ARGS(self.Args.At(2-1).ToEID(),EID{Optimize.C_EID.Id(),0}))
            }
            }  else if ((p.Id() == Core.C_mClaire_set_base.Id()) && 
              (s.Id() == C_void.Id())) { 
            PRINC("ClEnv.Base= ")
            Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_integer.Id(),0}))
            }  else if ((p.Id() == Core.C_mClaire_set_index.Id()) && 
              (s.Id() == C_void.Id())) { 
            PRINC("ClEnv.Index= ")
            Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_integer.Id(),0}))
            }  else if (p.Id() == Optimize.C_Compile_anyObject_I.Id()) { 
            if (a1 == C_Interval.Id()) { 
              F_Generate_object_prefix_class(C_any,s)
              F_Generate_class_ident_class(ToClass(a1))
              PRINC(".MakeInts(")
              { var arg_2 *ClaireList  
                var try_3 EID 
                try_3 = self.Args.Cdr()
                if ErrorIn(try_3) {Result = try_3
                } else {
                arg_2 = ToList(OBJ(try_3))
                Result = F_Generate_args_list_list(arg_2,C_integer)
                }
                } 
              if !ErrorIn(Result) {
              PRINC(")")
              F_Generate_object_post_class(C_any,s)
              Result = EVOID
              }
              } else {
              F_Generate_object_prefix_class(C_any,s)
              F_Generate_class_ident_class(ToClass(a1))
              PRINC(".Make(")
              { var arg_4 *ClaireList  
                var try_5 EID 
                try_5 = self.Args.Cdr()
                if ErrorIn(try_5) {Result = try_5
                } else {
                arg_4 = ToList(OBJ(try_5))
                Result = F_Generate_args_list_list(arg_4,C_any)
                }
                } 
              if !ErrorIn(Result) {
              PRINC(")")
              F_Generate_object_post_class(C_any,s)
              Result = EVOID
              }
              } 
            }  else if (p.Id() == C_add_slot.Id()) { 
            Result = F_Generate_cast_prefix_class(C_slot,s)
            if !ErrorIn(Result) {
            Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_class.Id(),0}))
            if !ErrorIn(Result) {
            PRINC(".AddSlot(")
            Result = Core.F_CALL(C_Generate_g_expression,ARGS(self.Args.At(2-1).ToEID(),EID{C_property.Id(),0}))
            if !ErrorIn(Result) {
            PRINC(",")
            Result = Core.F_CALL(C_Generate_g_expression,ARGS(self.Args.At(3-1).ToEID(),EID{C_type.Id(),0}))
            if !ErrorIn(Result) {
            PRINC(",")
            Result = Core.F_CALL(C_Generate_g_expression,ARGS(self.Args.At(4-1).ToEID(),EID{C_any.Id(),0}))
            if !ErrorIn(Result) {
            PRINC(")")
            F_Generate_cast_post_class(C_slot,s)
            Result = EVOID
            }}}}}
            } else {
            Result = F_Generate_eid_prefix_class(s)
            if !ErrorIn(Result) {
            F_Generate_preCore_ask_void()
            PRINC("F_CALL(")
            Result = F_Generate_g_expression_thing(ToThing(self.Selector.Id()),C_property)
            if !ErrorIn(Result) {
            PRINC(",ARGS(")
            Result = F_Generate_args_list_list(self.Args,Optimize.C_EID)
            if !ErrorIn(Result) {
            PRINC("))")
            F_Generate_eid_post_class(s)
            Result = EVOID
            }}}
            } 
          } 
        }
        } 
      } 
    return Result} 
  
// The EID go function for: inline_exp @ list<type_expression>(go_producer, Call, class) (throw: true) 
func E_Generate_inline_exp_go_producer1 (c EID,self EID,s EID) EID { 
    return F_Generate_inline_exp_go_producer1(ToGenerateGoProducer(OBJ(c)),Language.To_Call(OBJ(self)),ToClass(OBJ(s)) )} 
  
// produces a list of C expressions, separated by commas
/* The go function for: args_list(self:list,s:class) [status=1] */
func F_Generate_args_list_list (self *ClaireList ,s *ClaireClass ) EID { 
    // eid body s = void
    var Result EID 
    { var _Zfirst *ClaireBoolean   = CTRUE
      { var bk_ask *ClaireBoolean   = Core.F__sup_integer(self.Length(),3)
        if (bk_ask == CTRUE) { 
          Optimize.C_OPT.Level = (Optimize.C_OPT.Level+1)
          } 
        { 
          var x *ClaireAny  
          _ = x
          Result= EID{CFALSE.Id(),0}
          var x_support *ClaireList  
          x_support = self
          x_len := x_support.Length()
          for i_it := 0; i_it < x_len; i_it++ { 
            x = x_support.At(i_it)
            var loop_1 EID 
            _ = loop_1
            if (_Zfirst == CTRUE) { 
              loop_1 = Core.F_CALL(C_Generate_g_expression,ARGS(x.ToEID(),EID{s.Id(),0}))
              if ErrorIn(loop_1) {Result = loop_1
              break
              } else {
              _Zfirst = CFALSE
              loop_1 = EID{_Zfirst.Id(),0}
              }
              } else {
              PRINC(",")
              if (bk_ask == CTRUE) { 
                F_Generate_breakline_void()
                } 
              loop_1 = Core.F_CALL(C_Generate_g_expression,ARGS(x.ToEID(),EID{s.Id(),0}))
              if ErrorIn(loop_1) {Result = loop_1
              break
              } else {
              }
              } 
            if ErrorIn(loop_1) {Result = loop_1
            break
            } else {
            }
            } 
          } 
        if !ErrorIn(Result) {
        if (bk_ask == CTRUE) { 
          { 
            var va_arg1 *Optimize.OptimizeMetaOPT  
            var va_arg2 int 
            va_arg1 = Optimize.C_OPT
            va_arg2 = (Optimize.C_OPT.Level-1)
            va_arg1.Level = va_arg2
            Result = EID{C__INT,IVAL(va_arg2)}
            } 
          } else {
          Result = EID{CFALSE.Id(),0}
          } 
        }
        } 
      } 
    return Result} 
  
// The EID go function for: args_list @ list (throw: true) 
func E_Generate_args_list_list (self EID,s EID) EID { 
    return F_Generate_args_list_list(ToList(OBJ(self)),ToClass(OBJ(s)) )} 
  
// CLAIRE4 : get rid of fast dispatch (fcall + dispatcher)
// Super is like a call
/* The go function for: g_expression(self:Super,s:class) [status=1] */
func F_Generate_g_expression_Super (self *Language.Super ,s *ClaireClass ) EID { 
    // eid body s = void
    var Result EID 
    Result = F_Generate_eid_prefix_class(s)
    if !ErrorIn(Result) {
    F_Generate_preCore_ask_void()
    PRINC("F_SUPER(")
    Result = F_Generate_g_expression_thing(ToThing(self.Selector.Id()),C_property)
    if !ErrorIn(Result) {
    PRINC(", ")
    Result = Core.F_CALL(C_Generate_g_expression,ARGS(EID{self.CastTo.Id(),0},EID{C_class.Id(),0}))
    if !ErrorIn(Result) {
    PRINC(", ARGS(")
    Result = F_Generate_args_list_list(self.Args,Optimize.C_EID)
    if !ErrorIn(Result) {
    PRINC("))")
    F_Generate_eid_post_class(s)
    Result = EVOID
    }}}}
    return Result} 
  
// The EID go function for: g_expression @ Super (throw: true) 
func E_Generate_g_expression_Super (self EID,s EID) EID { 
    return F_Generate_g_expression_Super(Language.To_Super(OBJ(self)),ToClass(OBJ(s)) )} 
  
// *******************************************************************
// *       Part 3: the inline coding of function calls               *
// *******************************************************************
// CLAIRE4 Note : all inline optimization assume than can_throw?(m) = false
// these methods are important since they contain the open-coding optimisations. Some of the method calls are be replaced
// directly by  expressions. We always expect the native form (the sort s is passed as a parameter)
// functions with one argument
// note that we need the *_prefix / *_post 
/* The go function for: inline_exp(c:go_producer,self:Call_method1,s:class) [status=1] */
func F_Generate_inline_exp_go_producer2 (c *GenerateGoProducer ,self *Language.CallMethod1 ,s *ClaireClass ) EID { 
    // eid body s = void
    var Result EID 
    { var m *ClaireMethod   = self.Arg
      { var p *ClaireProperty   = m.Selector
        { var a1 *ClaireAny  
          var try_1 EID 
          try_1 = Core.F_car_list(self.Args)
          if ErrorIn(try_1) {Result = try_1
          } else {
          a1 = ANY(try_1)
          { var dm *ClaireClass   = Core.F_domain_I_restriction(ToRestriction(m.Id()))
            if ((p.Id() == C__dash.Id()) && 
                ((dm.Id() == C_integer.Id()) || 
                    (dm.Id() == C_float.Id()))) { 
              Result = F_Generate_cast_prefix_class(dm,s)
              if !ErrorIn(Result) {
              PRINC("(-")
              Result = F_Generate_bounded_expression_any(a1,s)
              if !ErrorIn(Result) {
              PRINC(")")
              F_Generate_cast_post_class(dm,s)
              Result = EVOID
              }}
              }  else if ((p.Id() == Core.C_owner.Id()) && 
                (F_Generate_eid_provide_ask_any(a1) == CTRUE)) { 
              F_Generate_object_prefix_class(C_class,s)
              PRINC("OWNER(")
              Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{Optimize.C_EID.Id(),0}))
              if !ErrorIn(Result) {
              PRINC(")")
              F_Generate_object_post_class(C_class,s)
              Result = EVOID
              }
              } else {
              var g0093I *ClaireBoolean  
              var try_2 EID 
              { 
                var v_and7 *ClaireBoolean  
                
                v_and7 = Equal(p.Id(),Core.C_owner.Id())
                if (v_and7 == CFALSE) {try_2 = EID{CFALSE.Id(),0}
                } else { 
                  var try_3 EID 
                  try_3 = Optimize.F_Compile_designated_ask_any(a1)
                  if ErrorIn(try_3) {try_2 = try_3
                  } else {
                  v_and7 = ToBoolean(OBJ(try_3))
                  if (v_and7 == CFALSE) {try_2 = EID{CFALSE.Id(),0}
                  } else { 
                    try_2 = EID{CTRUE.Id(),0}} 
                  } 
                }
                } 
              if ErrorIn(try_2) {Result = try_2
              } else {
              g0093I = ToBoolean(OBJ(try_2))
              if (g0093I == CTRUE) { 
                F_Generate_object_prefix_class(C_class,s)
                Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_any.Id(),0}))
                if !ErrorIn(Result) {
                PRINC(".Isa")
                F_Generate_object_post_class(C_class,s)
                Result = EVOID
                }
                }  else if (p.Id() == Core.C_eval.Id()) { 
                Result = F_Generate_eid_prefix_class(s)
                if !ErrorIn(Result) {
                PRINC("EVAL(")
                Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_any.Id(),0}))
                if !ErrorIn(Result) {
                PRINC(")")
                F_Generate_eid_post_class(s)
                Result = EVOID
                }}
                }  else if (m.Selector.Id() == Core.C_externC.Id()) { 
                Result = Core.F_CALL(C_princ,ARGS(a1.ToEID()))
                } else {
                var g0094I *ClaireBoolean  
                var try_4 EID 
                { 
                  var v_and8 *ClaireBoolean  
                  
                  v_and8 = Equal(m.Id(),C_Generate__starlength_bag_star.Value)
                  if (v_and8 == CFALSE) {try_4 = EID{CFALSE.Id(),0}
                  } else { 
                    var try_5 EID 
                    try_5 = Optimize.F_Compile_designated_ask_any(a1)
                    if ErrorIn(try_5) {try_4 = try_5
                    } else {
                    v_and8 = ToBoolean(OBJ(try_5))
                    if (v_and8 == CFALSE) {try_4 = EID{CFALSE.Id(),0}
                    } else { 
                      try_4 = EID{CTRUE.Id(),0}} 
                    } 
                  }
                  } 
                if ErrorIn(try_4) {Result = try_4
                } else {
                g0094I = ToBoolean(OBJ(try_4))
                if (g0094I == CTRUE) { 
                  F_Generate_integer_prefix_class(s)
                  Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_list.Id(),0}))
                  if !ErrorIn(Result) {
                  PRINC(".Length()")
                  F_Generate_native_post_class(s)
                  Result = EVOID
                  }
                  } else {
                  var g0095I *ClaireBoolean  
                  var try_6 EID 
                  { 
                    var v_and9 *ClaireBoolean  
                    
                    v_and9 = Equal(p.Id(),C_integer_I.Id())
                    if (v_and9 == CFALSE) {try_6 = EID{CFALSE.Id(),0}
                    } else { 
                      v_and9 = Equal(Core.F_domain_I_restriction(ToRestriction(m.Id())).Id(),C_char.Id())
                      if (v_and9 == CFALSE) {try_6 = EID{CFALSE.Id(),0}
                      } else { 
                        var try_7 EID 
                        try_7 = Optimize.F_Compile_designated_ask_any(a1)
                        if ErrorIn(try_7) {try_6 = try_7
                        } else {
                        v_and9 = ToBoolean(OBJ(try_7))
                        if (v_and9 == CFALSE) {try_6 = EID{CFALSE.Id(),0}
                        } else { 
                          try_6 = EID{CTRUE.Id(),0}} 
                        } 
                      } 
                    }
                    } 
                  if ErrorIn(try_6) {Result = try_6
                  } else {
                  g0095I = ToBoolean(OBJ(try_6))
                  if (g0095I == CTRUE) { 
                    F_Generate_integer_prefix_class(s)
                    PRINC("int(")
                    Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_char.Id(),0}))
                    if !ErrorIn(Result) {
                    PRINC(")")
                    F_Generate_native_post_class(s)
                    Result = EVOID
                    }
                    }  else if ((m.Id() == C_Generate__starof_bag_star.Value) || 
                      (m.Id() == C_Generate__starof_array_star.Value)) { 
                    Result = F_Generate_cast_prefix_class(C_type,s)
                    if !ErrorIn(Result) {
                    Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_list.Id(),0}))
                    if !ErrorIn(Result) {
                    PRINC(".Of()")
                    F_Generate_cast_post_class(C_type,s)
                    Result = EVOID
                    }}
                    }  else if (m.Id() == Optimize.C_Compile_m_unsafe.Value) { 
                    if (s.Id() == Optimize.C_EID.Id()) { 
                      Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{Optimize.C_EID.Id(),0}))
                      } else {
                      var g0096I *ClaireBoolean  
                      var try_8 EID 
                      { var arg_9 *ClaireBoolean  
                        var try_10 EID 
                        try_10 = Optimize.F_Compile_g_throw_any(a1)
                        if ErrorIn(try_10) {try_8 = try_10
                        } else {
                        arg_9 = ToBoolean(OBJ(try_10))
                        try_8 = EID{arg_9.Not.Id(),0}
                        }
                        } 
                      if ErrorIn(try_8) {Result = try_8
                      } else {
                      g0096I = ToBoolean(OBJ(try_8))
                      if (g0096I == CTRUE) { 
                        Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{s.Id(),0}))
                        } else {
                        Result = F_Generate_cast_prefix_class(C_any,s)
                        if !ErrorIn(Result) {
                        PRINC("ANY(")
                        Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{Optimize.C_EID.Id(),0}))
                        if !ErrorIn(Result) {
                        PRINC(")")
                        F_Generate_cast_post_class(C_any,s)
                        Result = EVOID
                        }}
                        } 
                      }
                      } 
                    }  else if ((m.Id() == C_Generate__starprinc_string_star.Value) && 
                      (C_string.Id() == a1.Isa.Id())) { 
                    PRINC("PRINC(")
                    Result = Core.F_CALL(C_print,ARGS(a1.ToEID()))
                    if !ErrorIn(Result) {
                    PRINC(")")
                    Result = EVOID
                    }
                    }  else if (m.Id() == C_Generate__starcopy_list_star.Value) { 
                    Result = F_Generate_cast_prefix_class(C_list,s)
                    if !ErrorIn(Result) {
                    Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_list.Id(),0}))
                    if !ErrorIn(Result) {
                    PRINC(".Copy()")
                    F_Generate_cast_post_class(C_list,s)
                    Result = EVOID
                    }}
                    }  else if (m.Id() == C_Generate__starlength_array_star.Value) { 
                    F_Generate_integer_prefix_class(s)
                    Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_array.Id(),0}))
                    if !ErrorIn(Result) {
                    PRINC(".Length()")
                    F_Generate_native_post_class(s)
                    Result = EVOID
                    }
                    } else {
                    var g0097I *ClaireBoolean  
                    var try_11 EID 
                    { 
                      var v_and10 *ClaireBoolean  
                      
                      v_and10 = Equal(m.Id(),C_Generate__starnot_star.Value)
                      if (v_and10 == CFALSE) {try_11 = EID{CFALSE.Id(),0}
                      } else { 
                        var try_12 EID 
                        { var arg_13 *ClaireClass  
                          var try_14 EID 
                          try_14 = Language.F_static_type_any(a1)
                          if ErrorIn(try_14) {try_12 = try_14
                          } else {
                          arg_13 = ToClass(OBJ(try_14))
                          try_12 = EID{ToType(arg_13.Id()).Included(ToType(C_boolean.Id())).Id(),0}
                          }
                          } 
                        if ErrorIn(try_12) {try_11 = try_12
                        } else {
                        v_and10 = ToBoolean(OBJ(try_12))
                        if (v_and10 == CFALSE) {try_11 = EID{CFALSE.Id(),0}
                        } else { 
                          try_11 = EID{CTRUE.Id(),0}} 
                        } 
                      }
                      } 
                    if ErrorIn(try_11) {Result = try_11
                    } else {
                    g0097I = ToBoolean(OBJ(try_11))
                    if (g0097I == CTRUE) { 
                      F_Generate_object_prefix_class(C_boolean,s)
                      Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_boolean.Id(),0}))
                      if !ErrorIn(Result) {
                      PRINC(".Not")
                      F_Generate_object_post_class(C_boolean,s)
                      Result = EVOID
                      }
                      }  else if ((m.Id() == C_Generate__starnew_class1_star.Value) && 
                        (C_class.Id() == a1.Isa.Id())) { 
                      F_Generate_object_prefix_class(C_any,s)
                      PRINC("new(")
                      F_Generate_go_class_class(ToClass(a1))
                      PRINC(").Is(")
                      Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_class.Id(),0}))
                      if !ErrorIn(Result) {
                      PRINC(")")
                      F_Generate_object_post_class(C_any,s)
                      Result = EVOID
                      }
                      } else {
                      Result = c.PrintExternalCall(Language.To_CallMethod(self.Id()),s)
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
  
// The EID go function for: inline_exp @ list<type_expression>(go_producer, Call_method1, class) (throw: true) 
func E_Generate_inline_exp_go_producer2 (c EID,self EID,s EID) EID { 
    return F_Generate_inline_exp_go_producer2(ToGenerateGoProducer(OBJ(c)),Language.To_CallMethod1(OBJ(self)),ToClass(OBJ(s)) )} 
  
// ===  functions with two arguments ===
/* The go function for: inline_exp(c:go_producer,self:Call_method2,s:class) [status=1] */
func F_Generate_inline_exp_go_producer3 (c *GenerateGoProducer ,self *Language.CallMethod2 ,s *ClaireClass ) EID { 
    // eid body s = void
    var Result EID 
    { var m *ClaireMethod   = self.Arg
      { var p *ClaireProperty   = m.Selector
        { var a1 *ClaireAny   = self.Args.At(1-1)
          { var a2 *ClaireAny   = self.Args.At(2-1)
            { var s1 *ClaireClass  
              var try_1 EID 
              { var arg_2 *ClaireType  
                var try_3 EID 
                try_3 = Core.F_CALL(Optimize.C_c_type,ARGS(a1.ToEID()))
                if ErrorIn(try_3) {try_1 = try_3
                } else {
                arg_2 = ToType(OBJ(try_3))
                try_1 = EID{arg_2.Class_I().Id(),0}
                }
                } 
              if ErrorIn(try_1) {Result = try_1
              } else {
              s1 = ToClass(OBJ(try_1))
              if ((p.Id() == C_class_I.Id()) && 
                  (a1.Isa.IsIn(C_symbol) == CTRUE)) { 
                F_Generate_symbol_ident_symbol(ToSymbol(a1))
                PRINC(" = MakeClass(")
                { var arg_4 *ClaireAny  
                  var try_5 EID 
                  try_5 = Core.F_CALL(C_string_I,ARGS(a1.ToEID()))
                  if ErrorIn(try_5) {Result = try_5
                  } else {
                  arg_4 = ANY(try_5)
                  Result = Core.F_print_any(arg_4)
                  }
                  } 
                if !ErrorIn(Result) {
                PRINC(",")
                Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{C_class.Id(),0}))
                if !ErrorIn(Result) {
                PRINC(",")
                Result = F_Generate_g_expression_module(ToModule(OBJ(Core.F_CALL(C_module_I,ARGS(a1.ToEID())))),C_module)
                if !ErrorIn(Result) {
                PRINC(")")
                Result = EVOID
                }}}
                } else {
                var g0101I *ClaireBoolean  
                { 
                  var v_and8 *ClaireBoolean  
                  
                  v_and8 = Equal(m.Domain.ValuesO()[1-1],m.Domain.ValuesO()[2-1])
                  if (v_and8 == CFALSE) {g0101I = CFALSE
                  } else { 
                    v_and8 = MakeBoolean((s1.Id() == C_integer.Id()) || (s1.Id() == C_float.Id()))
                    if (v_and8 == CFALSE) {g0101I = CFALSE
                    } else { 
                      { 
                        var v_or11 *ClaireBoolean  
                        
                        v_or11 = c.OpenOperators.Memq(p.Id())
                        if (v_or11 == CTRUE) {v_and8 = CTRUE
                        } else { 
                          { 
                            var v_and13 *ClaireBoolean  
                            
                            v_and13 = c.DivOperators.Memq(p.Id())
                            if (v_and13 == CFALSE) {v_or11 = CFALSE
                            } else { 
                              if (C_integer.Id() == a2.Isa.Id()) { 
                                { var g0098 int  = ToInteger(a2).Value
                                  v_and13 = Core.F__I_equal_any(MakeInteger(g0098).Id(),MakeInteger(0).Id())
                                  } 
                                }  else if (C_float.Id() == a2.Isa.Id()) { 
                                { var g0099 float64  = ToFloat(a2).Value
                                  v_and13 = Core.F__I_equal_any(MakeFloat(g0099).Id(),MakeFloat(0).Id())
                                  } 
                                } else {
                                v_and13 = F__sup_equal_integer(Optimize.C_compiler.Safety,3)
                                } 
                              if (v_and13 == CFALSE) {v_or11 = CFALSE
                              } else { 
                                v_or11 = CTRUE} 
                              } 
                            } 
                          if (v_or11 == CTRUE) {v_and8 = CTRUE
                          } else { 
                            v_and8 = CFALSE} 
                          } 
                        } 
                      if (v_and8 == CFALSE) {g0101I = CFALSE
                      } else { 
                        g0101I = CTRUE} 
                      } 
                    } 
                  } 
                if (g0101I == CTRUE) { 
                  Result = F_Generate_cast_prefix_class(s1,s)
                  if !ErrorIn(Result) {
                  PRINC("(")
                  Result = F_Generate_bounded_expression_any(a1,s1)
                  if !ErrorIn(Result) {
                  { var arg_6 *ClaireString  
                    if (p.Id() == C_mod.Id()) { 
                      arg_6 = MakeString("%")
                      } else {
                      arg_6 = p.Name.String_I()
                      } 
                    F_princ_string(arg_6)
                    } 
                  Result = F_Generate_bounded_expression_any(a2,s1)
                  if !ErrorIn(Result) {
                  PRINC(")")
                  F_Generate_cast_post_class(s1,s)
                  Result = EVOID
                  }}}
                  } else {
                  var g0102I *ClaireBoolean  
                  var try_7 EID 
                  { 
                    var v_and9 *ClaireBoolean  
                    
                    v_and9 = Equal(m.Id(),C_Generate__starcontain_list_star.Value)
                    if (v_and9 == CFALSE) {try_7 = EID{CFALSE.Id(),0}
                    } else { 
                      var try_8 EID 
                      try_8 = Optimize.F_Compile_identifiable_ask_any(a2)
                      if ErrorIn(try_8) {try_7 = try_8
                      } else {
                      v_and9 = ToBoolean(OBJ(try_8))
                      if (v_and9 == CFALSE) {try_7 = EID{CFALSE.Id(),0}
                      } else { 
                        try_7 = EID{CTRUE.Id(),0}} 
                      } 
                    }
                    } 
                  if ErrorIn(try_7) {Result = try_7
                  } else {
                  g0102I = ToBoolean(OBJ(try_7))
                  if (g0102I == CTRUE) { 
                    F_Generate_object_prefix_class(C_boolean,s)
                    Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_list.Id(),0}))
                    if !ErrorIn(Result) {
                    PRINC(".Memq(")
                    Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{C_any.Id(),0}))
                    if !ErrorIn(Result) {
                    PRINC(")")
                    F_Generate_object_post_class(C_boolean,s)
                    Result = EVOID
                    }}
                    }  else if (m.Id() == C_Generate__starcontain_set_star.Value) { 
                    { var sm *ClaireClass  
                      var try_9 EID 
                      try_9 = F_Generate_g_member_any(a1)
                      if ErrorIn(try_9) {Result = try_9
                      } else {
                      sm = ToClass(OBJ(try_9))
                      F_Generate_object_prefix_class(C_boolean,s)
                      Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_set.Id(),0}))
                      if !ErrorIn(Result) {
                      PRINC(".Contain")
                      if (sm.Id() == C_integer.Id()) { 
                        PRINC("SetInteger")
                        }  else if (sm.Id() == C_float.Id()) { 
                        PRINC("SetFloat")
                        } else {
                        PRINC("_ask")
                        } 
                      PRINC("(")
                      { var arg_10 *ClaireClass  
                        if (sm.Id() == C_integer.Id()) { 
                          arg_10 = C_integer
                          }  else if (sm.Id() == C_float.Id()) { 
                          arg_10 = C_float
                          } else {
                          arg_10 = C_any
                          } 
                        Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{arg_10.Id(),0}))
                        } 
                      if !ErrorIn(Result) {
                      PRINC(")")
                      F_Generate_object_post_class(C_boolean,s)
                      Result = EVOID
                      }}
                      }
                      } 
                    }  else if (m.Selector.Id() == Core.C_externC.Id()) { 
                    Result = Core.F_CALL(C_princ,ARGS(a1.ToEID()))
                    }  else if (m.Id() == Optimize.C_Compile_m_member.Value) { 
                    Result = F_Generate_belong_exp_any(a1,a2,s)
                    }  else if ((m.Id() == C_Generate__starwrite_value_star.Value) && 
                      (F_Generate_eid_provide_ask_any(a2) == CTRUE)) { 
                    Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_Variable.Id(),0}))
                    if !ErrorIn(Result) {
                    PRINC(".WriteEID(")
                    Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{Optimize.C_EID.Id(),0}))
                    if !ErrorIn(Result) {
                    PRINC(")")
                    Result = EVOID
                    }}
                    }  else if (m.Id() == C_Generate__starinherit_star.Value) { 
                    F_Generate_object_prefix_class(C_boolean,s)
                    Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_class.Id(),0}))
                    if !ErrorIn(Result) {
                    PRINC(".IsIn(")
                    Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{C_class.Id(),0}))
                    if !ErrorIn(Result) {
                    PRINC(")")
                    F_Generate_object_post_class(C_boolean,s)
                    Result = EVOID
                    }}
                    }  else if (m.Id() == C_Generate__starequal_star.Value) { 
                    F_Generate_object_prefix_class(C_boolean,s)
                    PRINC("Equal(")
                    Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_any.Id(),0}))
                    if !ErrorIn(Result) {
                    PRINC(",")
                    Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{C_any.Id(),0}))
                    if !ErrorIn(Result) {
                    PRINC(")")
                    F_Generate_object_post_class(C_boolean,s)
                    Result = EVOID
                    }}
                    }  else if (m.Id() == C_Generate__starmap_star.Value) { 
                    Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_type.Id(),0}))
                    if !ErrorIn(Result) {
                    PRINC(".Map_I(")
                    Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{C_type.Id(),0}))
                    if !ErrorIn(Result) {
                    PRINC(")")
                    Result = EVOID
                    }}
                    }  else if (m.Id() == C_Generate__star_Zt_star.Value) { 
                    F_Generate_object_prefix_class(C_boolean,s)
                    Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{C_type.Id(),0}))
                    if !ErrorIn(Result) {
                    PRINC(".Contains(")
                    Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_any.Id(),0}))
                    if !ErrorIn(Result) {
                    PRINC(")")
                    F_Generate_object_post_class(C_boolean,s)
                    Result = EVOID
                    }}
                    }  else if ((p.Id() == Core.C_Core__inf_equalt.Id()) || 
                      (m.Id() == C_Generate__starincluded_star.Value)) { 
                    F_Generate_object_prefix_class(C_boolean,s)
                    Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_type.Id(),0}))
                    if !ErrorIn(Result) {
                    PRINC(".Included(")
                    Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{C_type.Id(),0}))
                    if !ErrorIn(Result) {
                    PRINC(")")
                    F_Generate_object_post_class(C_boolean,s)
                    Result = EVOID
                    }}
                    } else {
                    var g0103I *ClaireBoolean  
                    var try_11 EID 
                    { 
                      var v_or10 *ClaireBoolean  
                      
                      var try_12 EID 
                      { 
                        var v_and11 *ClaireBoolean  
                        
                        v_and11 = MakeBoolean((((m.Id() == C_Generate__starnth_list_star.Value) || 
                              (m.Id() == C_Generate__starnth_tuple_star.Value)) && 
                            (Optimize.C_compiler.Safety >= 3)) || (m.Id() == C_Generate__starnth_1_list_star.Value) || (m.Id() == C_Generate__starnth_1_tuple_star.Value) || (m.Id() == C_Generate__starnth_1_array_star.Value))
                        if (v_and11 == CFALSE) {try_12 = EID{CFALSE.Id(),0}
                        } else { 
                          var try_13 EID 
                          { var arg_14 *ClaireClass  
                            var try_15 EID 
                            try_15 = F_Generate_g_member_any(a1)
                            if ErrorIn(try_15) {try_13 = try_15
                            } else {
                            arg_14 = ToClass(OBJ(try_15))
                            try_13 = EID{Core.F__I_equal_any(arg_14.Id(),C_any.Id()).Id(),0}
                            }
                            } 
                          if ErrorIn(try_13) {try_12 = try_13
                          } else {
                          v_and11 = ToBoolean(OBJ(try_13))
                          if (v_and11 == CFALSE) {try_12 = EID{CFALSE.Id(),0}
                          } else { 
                            try_12 = EID{CTRUE.Id(),0}} 
                          } 
                        }
                        } 
                      if ErrorIn(try_12) {try_11 = try_12
                      } else {
                      v_or10 = ToBoolean(OBJ(try_12))
                      if (v_or10 == CTRUE) {try_11 = EID{CTRUE.Id(),0}
                      } else { 
                        v_or10 = Equal(m.Selector.Id(),Core.C_mClaire_nth_object.Id())
                        if (v_or10 == CTRUE) {try_11 = EID{CTRUE.Id(),0}
                        } else { 
                          try_11 = EID{CFALSE.Id(),0}} 
                        } 
                      }
                      } 
                    if ErrorIn(try_11) {Result = try_11
                    } else {
                    g0103I = ToBoolean(OBJ(try_11))
                    if (g0103I == CTRUE) { 
                      { var s1 *ClaireClass  
                        var try_16 EID 
                        if (m.Selector.Id() == Core.C_mClaire_nth_object.Id()) { 
                          try_16 = EID{C_object.Id(),0}
                          } else {
                          { var arg_17 *ClaireClass  
                            var try_18 EID 
                            try_18 = F_Generate_g_member_any(a1)
                            if ErrorIn(try_18) {try_16 = try_18
                            } else {
                            arg_17 = ToClass(OBJ(try_18))
                            try_16 = EID{F_Generate_type_sort_type(ToType(arg_17.Id())).Id(),0}
                            }
                            } 
                          } 
                        if ErrorIn(try_16) {Result = try_16
                        } else {
                        s1 = ToClass(OBJ(try_16))
                        Result = F_Generate_cast_prefix_class(s1,s)
                        if !ErrorIn(Result) {
                        Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_list.Id(),0}))
                        if !ErrorIn(Result) {
                        PRINC(".")
                        { var arg_19 *ClaireClass  
                          var try_20 EID 
                          try_20 = F_Generate_g_member_any(a1)
                          if ErrorIn(try_20) {Result = try_20
                          } else {
                          arg_19 = ToClass(OBJ(try_20))
                          F_Generate_valuesSlot_class(arg_19)
                          Result = EVOID
                          }
                          } 
                        if !ErrorIn(Result) {
                        PRINC("[")
                        Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{C_integer.Id(),0}))
                        if !ErrorIn(Result) {
                        PRINC("-1]")
                        F_Generate_cast_post_class(s1,s)
                        Result = EVOID
                        }}}}
                        }
                        } 
                      }  else if ((((m.Id() == C_Generate__starnth_list_star.Value) || 
                            (m.Id() == C_Generate__starnth_tuple_star.Value)) && 
                          (Optimize.C_compiler.Safety >= 3)) || 
                        ((m.Id() == C_Generate__starnth_1_list_star.Value) || 
                            ((m.Id() == C_Generate__starnth_1_tuple_star.Value) || 
                              (m.Id() == C_Generate__starnth_1_array_star.Value)))) { 
                      Result = F_Generate_cast_prefix_class(C_any,s)
                      if !ErrorIn(Result) {
                      Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_list.Id(),0}))
                      if !ErrorIn(Result) {
                      PRINC(".At(")
                      Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{C_integer.Id(),0}))
                      if !ErrorIn(Result) {
                      PRINC("-1)")
                      F_Generate_cast_post_class(C_any,s)
                      Result = EVOID
                      }}}
                      }  else if ((p.Id() == C_add_I.Id()) && 
                        (ToType(Core.F_domain_I_restriction(ToRestriction(m.Id())).Id()).Included(ToType(C_bag.Id())) == CTRUE)) { 
                      { var sbag *ClaireClass  
                        if (Core.F_domain_I_restriction(ToRestriction(m.Id())).Id() == C_set.Id()) { 
                          sbag = C_set
                          } else {
                          sbag = C_list
                          } 
                        { var _Ztype *ClaireClass  
                          var try_21 EID 
                          try_21 = F_Generate_g_member_any(a1)
                          if ErrorIn(try_21) {Result = try_21
                          } else {
                          _Ztype = ToClass(OBJ(try_21))
                          if ((sbag.Id() == C_list.Id()) && 
                              ((_Ztype.Id() == C_integer.Id()) && 
                                (s.Id() == C_void.Id()))) { 
                            Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_list.Id(),0}))
                            if !ErrorIn(Result) {
                            PRINC(".AddFastInteger(")
                            Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{C_integer.Id(),0}))
                            if !ErrorIn(Result) {
                            PRINC(")")
                            Result = EVOID
                            }}
                            }  else if ((sbag.Id() == C_set.Id()) && 
                              (_Ztype.Id() == C_integer.Id())) { 
                            Result = F_Generate_cast_prefix_class(sbag,s)
                            if !ErrorIn(Result) {
                            Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_set.Id(),0}))
                            if !ErrorIn(Result) {
                            PRINC(".AddSetInteger(")
                            Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{C_integer.Id(),0}))
                            if !ErrorIn(Result) {
                            PRINC(")")
                            F_Generate_cast_post_class(sbag,s)
                            Result = EVOID
                            }}}
                            } else {
                            Result = F_Generate_cast_prefix_class(sbag,s)
                            if !ErrorIn(Result) {
                            Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{Core.F_domain_I_restriction(ToRestriction(m.Id())).Id(),0}))
                            if !ErrorIn(Result) {
                            PRINC(".AddFast(")
                            Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{C_any.Id(),0}))
                            if !ErrorIn(Result) {
                            PRINC(")")
                            F_Generate_cast_post_class(sbag,s)
                            Result = EVOID
                            }}}
                            } 
                          }
                          } 
                        } 
                      }  else if ((m.Id() == C_Generate__starnth_1_string_star.Value) || 
                        ((m.Id() == C_Generate__starnth_string_star.Value) && 
                            (Optimize.C_compiler.Safety >= 3))) { 
                      F_Generate_char_prefix_class(s)
                      Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_string.Id(),0}))
                      if !ErrorIn(Result) {
                      PRINC(".At(")
                      Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{C_integer.Id(),0}))
                      if !ErrorIn(Result) {
                      PRINC(")")
                      F_Generate_native_post_class(s)
                      Result = EVOID
                      }}
                      } else {
                      var g0104I *ClaireBoolean  
                      var try_22 EID 
                      { 
                        var v_and11 *ClaireBoolean  
                        
                        v_and11 = Equal(C_table.Id(),a1.Isa.Id())
                        if (v_and11 == CFALSE) {try_22 = EID{CFALSE.Id(),0}
                        } else { 
                          v_and11 = Equal(p.Id(),C_nth.Id())
                          if (v_and11 == CFALSE) {try_22 = EID{CFALSE.Id(),0}
                          } else { 
                            var try_23 EID 
                            { 
                              var v_or14 *ClaireBoolean  
                              
                              var try_24 EID 
                              { var arg_25 *ClaireType  
                                var try_26 EID 
                                try_26 = Core.F_CALL(Optimize.C_c_type,ARGS(a2.ToEID()))
                                if ErrorIn(try_26) {try_24 = try_26
                                } else {
                                arg_25 = ToType(OBJ(try_26))
                                try_24 = EID{arg_25.Included(ToRelation(a1).Domain).Id(),0}
                                }
                                } 
                              if ErrorIn(try_24) {try_23 = try_24
                              } else {
                              v_or14 = ToBoolean(OBJ(try_24))
                              if (v_or14 == CTRUE) {try_23 = EID{CTRUE.Id(),0}
                              } else { 
                                v_or14 = F__sup_equal_integer(Optimize.C_compiler.Safety,2)
                                if (v_or14 == CTRUE) {try_23 = EID{CTRUE.Id(),0}
                                } else { 
                                  try_23 = EID{CFALSE.Id(),0}} 
                                } 
                              }
                              } 
                            if ErrorIn(try_23) {try_22 = try_23
                            } else {
                            v_and11 = ToBoolean(OBJ(try_23))
                            if (v_and11 == CFALSE) {try_22 = EID{CFALSE.Id(),0}
                            } else { 
                              try_22 = EID{CTRUE.Id(),0}} 
                            } 
                          } 
                        }
                        } 
                      if ErrorIn(try_22) {Result = try_22
                      } else {
                      g0104I = ToBoolean(OBJ(try_22))
                      if (g0104I == CTRUE) { 
                        Result = F_Generate_cast_prefix_class(C_any,s)
                        if !ErrorIn(Result) {
                        F_Generate_preCore_ask_void()
                        PRINC("F_get_table(")
                        Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_table.Id(),0}))
                        if !ErrorIn(Result) {
                        PRINC(",")
                        Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{C_any.Id(),0}))
                        if !ErrorIn(Result) {
                        PRINC(")")
                        F_Generate_cast_post_class(C_any,s)
                        Result = EVOID
                        }}}
                        }  else if (m.Selector.Id() == Core.C_identical_ask.Id()) { 
                        Result = F_Generate_cast_prefix_class(C_boolean,s)
                        if !ErrorIn(Result) {
                        PRINC("MakeBoolean(")
                        Result = Core.F_CALL(Optimize.C_Compile_bool_exp,ARGS(EID{self.Id(),0},EID{CTRUE.Id(),0}))
                        if !ErrorIn(Result) {
                        PRINC(")")
                        F_Generate_cast_post_class(C_boolean,s)
                        Result = EVOID
                        }}
                        }  else if ((p.Id() == Core.C_inlineok_ask.Id()) && 
                          (C_string.Id() == a2.Isa.Id())) { 
                        F_Generate_preCore_ask_void()
                        PRINC("F_inlineok_ask_method(")
                        F_Generate_breakline_void()
                        Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_property.Id(),0}))
                        if !ErrorIn(Result) {
                        PRINC(",")
                        F_Generate_breakline_void()
                        PRINC("MakeString(")
                        Result = Core.F_CALL(C_print,ARGS(a2.ToEID()))
                        if !ErrorIn(Result) {
                        PRINC("))")
                        Result = EVOID
                        }}
                        }  else if ((m.Id() == C_Generate__starnew_class2_star.Value) && 
                          ((C_class.Id() == a1.Isa.Id()) && 
                            (Optimize.C_compiler.Safety >= 2))) { 
                        F_Generate_object_prefix_class(C_any,s)
                        PRINC("new(")
                        F_Generate_go_class_class(ToClass(a1))
                        PRINC(").IsNamed(")
                        Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_class.Id(),0}))
                        if !ErrorIn(Result) {
                        PRINC(",")
                        Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{C_symbol.Id(),0}))
                        if !ErrorIn(Result) {
                        PRINC(")")
                        F_Generate_object_post_class(C_any,s)
                        Result = EVOID
                        }}
                        } else {
                        Result = c.PrintExternalCall(Language.To_CallMethod(self.Id()),s)
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
  
// The EID go function for: inline_exp @ list<type_expression>(go_producer, Call_method2, class) (throw: true) 
func E_Generate_inline_exp_go_producer3 (c EID,self EID,s EID) EID { 
    return F_Generate_inline_exp_go_producer3(ToGenerateGoProducer(OBJ(c)),Language.To_CallMethod2(OBJ(self)),ToClass(OBJ(s)) )} 
  
// === functions with three arguments or more
/* The go function for: inline_exp(c:go_producer,self:Call_method,s:class) [status=1] */
func F_Generate_inline_exp_go_producer4 (c *GenerateGoProducer ,self *Language.CallMethod ,s *ClaireClass ) EID { 
    // eid body s = void
    var Result EID 
    { var m *ClaireMethod   = self.Arg
      { var a1 *ClaireAny   = self.Args.At(1-1)
        { var a2 *ClaireAny   = self.Args.At(2-1)
          { var a3 *ClaireAny   = self.Args.At(3-1)
            var g0105I *ClaireBoolean  
            var try_1 EID 
            { 
              var v_and6 *ClaireBoolean  
              
              v_and6 = Equal(m.Id(),C_Generate__starnth_equal_list_star.Value)
              if (v_and6 == CFALSE) {try_1 = EID{CFALSE.Id(),0}
              } else { 
                v_and6 = F__sup_equal_integer(Optimize.C_compiler.Safety,3)
                if (v_and6 == CFALSE) {try_1 = EID{CFALSE.Id(),0}
                } else { 
                  var try_2 EID 
                  { var arg_3 *ClaireClass  
                    var try_4 EID 
                    try_4 = F_Generate_g_member_any(a1)
                    if ErrorIn(try_4) {try_2 = try_4
                    } else {
                    arg_3 = ToClass(OBJ(try_4))
                    try_2 = EID{Core.F__I_equal_any(arg_3.Id(),C_any.Id()).Id(),0}
                    }
                    } 
                  if ErrorIn(try_2) {try_1 = try_2
                  } else {
                  v_and6 = ToBoolean(OBJ(try_2))
                  if (v_and6 == CFALSE) {try_1 = EID{CFALSE.Id(),0}
                  } else { 
                    v_and6 = Equal(s.Id(),C_void.Id())
                    if (v_and6 == CFALSE) {try_1 = EID{CFALSE.Id(),0}
                    } else { 
                      try_1 = EID{CTRUE.Id(),0}} 
                    } 
                  } 
                } 
              }
              } 
            if ErrorIn(try_1) {Result = try_1
            } else {
            g0105I = ToBoolean(OBJ(try_1))
            if (g0105I == CTRUE) { 
              Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_list.Id(),0}))
              if !ErrorIn(Result) {
              PRINC(".")
              { var arg_5 *ClaireClass  
                var try_6 EID 
                try_6 = F_Generate_g_member_any(a1)
                if ErrorIn(try_6) {Result = try_6
                } else {
                arg_5 = ToClass(OBJ(try_6))
                F_Generate_valuesSlot_class(arg_5)
                Result = EVOID
                }
                } 
              if !ErrorIn(Result) {
              PRINC("[")
              Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{C_integer.Id(),0}))
              if !ErrorIn(Result) {
              PRINC("-1]=")
              { var arg_7 *ClaireClass  
                var try_8 EID 
                try_8 = F_Generate_g_member_any(a1)
                if ErrorIn(try_8) {Result = try_8
                } else {
                arg_7 = ToClass(OBJ(try_8))
                Result = Core.F_CALL(C_Generate_g_expression,ARGS(a3.ToEID(),EID{arg_7.Id(),0}))
                }
                } 
              }}}
              }  else if ((m.Id() == C_Generate__starnth_put_list_star.Value) || 
                ((m.Id() == C_Generate__starnth_put_array_star.Value) || 
                  ((Optimize.C_compiler.Safety >= 3) && 
                      (m.Id() == C_Generate__starnth_equal_list_star.Value)))) { 
              Result = F_Generate_cast_prefix_class(C_any,s)
              if !ErrorIn(Result) {
              Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_array.Id(),0}))
              if !ErrorIn(Result) {
              PRINC(".NthPut(")
              Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{C_integer.Id(),0}))
              if !ErrorIn(Result) {
              PRINC(",")
              Result = Core.F_CALL(C_Generate_g_expression,ARGS(a3.ToEID(),EID{C_any.Id(),0}))
              if !ErrorIn(Result) {
              PRINC(")")
              F_Generate_cast_post_class(C_any,s)
              Result = EVOID
              }}}}
              }  else if ((m.Id() == C_Generate__starmake_list_star.Value) && 
                (a3 == C_void.Id())) { 
              Result = F_Generate_cast_prefix_class(C_list,s)
              if !ErrorIn(Result) {
              PRINC("CreateList(")
              Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{C_type.Id(),0}))
              if !ErrorIn(Result) {
              PRINC(",")
              Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_integer.Id(),0}))
              if !ErrorIn(Result) {
              PRINC(")")
              F_Generate_cast_post_class(C_list,s)
              Result = EVOID
              }}}
              }  else if ((m.Selector.Id() == C_add_slot.Id()) && 
                (C_class.Id() == a1.Isa.Id())) { 
              F_Generate_preCore_ask_void()
              PRINC("F_close_slot(")
              Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_class.Id(),0}))
              if !ErrorIn(Result) {
              PRINC(".AddSlot(")
              Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{C_property.Id(),0}))
              if !ErrorIn(Result) {
              PRINC(",")
              Result = Core.F_CALL(C_Generate_g_expression,ARGS(a3.ToEID(),EID{C_type.Id(),0}))
              if !ErrorIn(Result) {
              PRINC(",")
              Result = Core.F_CALL(C_Generate_g_expression,ARGS(self.Args.At(4-1).ToEID(),EID{C_any.Id(),0}))
              if !ErrorIn(Result) {
              PRINC("))")
              Result = EVOID
              }}}}
              }  else if (m.Selector.Id() == C_add_method.Id()) { 
              if (a1.Isa.IsIn(C_property) == CTRUE) { 
                { var m *ClaireAny   = self.Args.At(6-1)
                  F_Generate_preCore_ask_void()
                  PRINC("F_attach_method(")
                  Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{a1.Isa.Id(),0}))
                  if !ErrorIn(Result) {
                  PRINC(".")
                  F_princ_string(ToString(IfThenElse((a1 == Core.C_self_eval.Id()),
                    MakeString("AddEvalMethod").Id(),
                    MakeString("AddMethod").Id())))
                  PRINC("(")
                  Result = c.Signature_I(F_Generate_full_signature_method(ToMethod(m)))
                  if !ErrorIn(Result) {
                  PRINC(",")
                  { var arg_9 int 
                    var try_10 EID 
                    var g0106I *ClaireBoolean  
                    var try_11 EID 
                    try_11 = Optimize.F_Compile_can_throw_status_method(ToMethod(m))
                    if ErrorIn(try_11) {try_10 = try_11
                    } else {
                    g0106I = ToBoolean(OBJ(try_11))
                    if (g0106I == CTRUE) { 
                      try_10 = EID{C__INT,IVAL(1)}
                      } else {
                      try_10 = EID{C__INT,IVAL(0)}
                      } 
                    }
                    if ErrorIn(try_10) {Result = try_10
                    } else {
                    arg_9 = INT(try_10)
                    F_princ_integer(arg_9)
                    Result = EVOID
                    }
                    } 
                  if !ErrorIn(Result) {
                  PRINC(",")
                  Result = F_Generate_goEIDFunction_method(ToMethod(m))
                  if !ErrorIn(Result) {
                  if (a1 == Core.C_self_eval.Id()) { 
                    F_Generate_goEvalFunction_method(ToMethod(m))
                    } 
                  PRINC("),MakeString(")
                  Result = Core.F_print_any(Core.F_get_table(Optimize.C_Compile_FileOrigin,m))
                  if !ErrorIn(Result) {
                  PRINC("))")
                  Result = EVOID
                  }}}}}
                  } 
                } else {
                PRINC("F_add_method_property(")
                Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_property.Id(),0}))
                if !ErrorIn(Result) {
                PRINC(",")
                Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{C_list.Id(),0}))
                if !ErrorIn(Result) {
                PRINC(",")
                Result = Core.F_CALL(C_Generate_g_expression,ARGS(a3.ToEID(),EID{C_type.Id(),0}))
                if !ErrorIn(Result) {
                PRINC(",")
                Result = Core.F_CALL(C_Generate_g_expression,ARGS(self.Args.At(4-1).ToEID(),EID{C_integer.Id(),0}))
                if !ErrorIn(Result) {
                PRINC(",")
                Result = Core.F_CALL(C_Generate_g_expression,ARGS(self.Args.At(5-1).ToEID(),EID{C_function.Id(),0}))
                if !ErrorIn(Result) {
                PRINC(")")
                Result = EVOID
                }}}}}
                } 
              } else {
              Result = c.PrintExternalCall(self,s)
              } 
            }
            } 
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: inline_exp @ list<type_expression>(go_producer, Call_method, class) (throw: true) 
func E_Generate_inline_exp_go_producer4 (c EID,self EID,s EID) EID { 
    return F_Generate_inline_exp_go_producer4(ToGenerateGoProducer(OBJ(c)),Language.To_CallMethod(OBJ(self)),ToClass(OBJ(s)) )} 
  
// THIS IS ONE OF THE KEY PATTERNS: calls a method through its compiled function
// the arguments and the result are expected in native format
/* The go function for: print_external_call(c:go_producer,self:Call_method,s:class) [status=1] */
func (c *GenerateGoProducer ) PrintExternalCall (self *Language.CallMethod ,s *ClaireClass ) EID { 
    // eid body s = void
    var Result EID 
    { var m *ClaireMethod   = self.Arg
      { var l *ClaireList   = self.Args
        { var n int  = 1
          _ = n
          { var _Zsig *ClaireList   = F_Generate_go_signature_method(m)
            { var sm *ClaireClass  
              var try_1 EID 
              try_1 = Core.F_last_list(_Zsig)
              if ErrorIn(try_1) {Result = try_1
              } else {
              sm = ToClass(OBJ(try_1))
              if (l.Length() > 4) { 
                Optimize.C_OPT.Level = (Optimize.C_OPT.Level+1)
                } 
              var g0111I *ClaireBoolean  
              var try_2 EID 
              try_2 = Optimize.F_Compile_can_throw_status_method(m)
              if ErrorIn(try_2) {Result = try_2
              } else {
              g0111I = ToBoolean(OBJ(try_2))
              if (g0111I == CTRUE) { 
                sm = Optimize.C_EID
                Result = EID{sm.Id(),0}
                } else {
                Result = EID{CFALSE.Id(),0}
                } 
              }
              if !ErrorIn(Result) {
              
              Result = F_Generate_cast_prefix_class(sm,s)
              if !ErrorIn(Result) {
              if (F_Generate_goMethod_ask_any(m.Id()) == CTRUE) { 
                Result = F_Generate_external_casted_arg_any(l.At(1-1),ToClass(_Zsig.ValuesO()[1-1]),0,Core.F__sup_integer(l.Length(),4))
                if !ErrorIn(Result) {
                PRINC(".")
                F_Generate_goMethod_method(m)
                PRINC("(")
                { var n int  = 2
                  { var g0109 int  = l.Length()
                    Result= EID{CFALSE.Id(),0}
                    for (n <= g0109) { 
                      var loop_3 EID 
                      _ = loop_3
                      { 
                      loop_3 = F_Generate_external_casted_arg_any(l.At(n-1),ToClass(_Zsig.ValuesO()[n-1]),(n-1),Core.F__sup_integer(l.Length(),4))
                      if ErrorIn(loop_3) {Result = loop_3
                      break
                      } else {
                      n = (n+1)
                      }
                      } 
                    }
                    } 
                  } 
                }
                } else {
                Result = F_Generate_goFunction_method(m)
                if !ErrorIn(Result) {
                PRINC("(")
                Result = EVOID
                }
                if !ErrorIn(Result) {
                if ((l.Length() == 1) && 
                    (Core.F_domain_I_restriction(ToRestriction(m.Id())).Id() == C_void.Id())) { 
                  l = CNIL
                  } 
                { var n int  = 1
                  { var g0110 int  = l.Length()
                    Result= EID{CFALSE.Id(),0}
                    for (n <= g0110) { 
                      var loop_4 EID 
                      _ = loop_4
                      { 
                      loop_4 = F_Generate_external_casted_arg_any(l.At(n-1),ToClass(_Zsig.ValuesO()[n-1]),n,Core.F__sup_integer(l.Length(),4))
                      if ErrorIn(loop_4) {Result = loop_4
                      break
                      } else {
                      n = (n+1)
                      }
                      } 
                    }
                    } 
                  } 
                }
                } 
              if !ErrorIn(Result) {
              PRINC(")")
              if (l.Length() > 4) { 
                Optimize.C_OPT.Level = (Optimize.C_OPT.Level-1)
                } 
              F_Generate_cast_post_class(sm,s)
              Result = EVOID
              }}}
              }
              } 
            } 
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: print_external_call @ go_producer (throw: true) 
func E_Generate_print_external_call_go_producer (c EID,self EID,s EID) EID { 
    return ToGenerateGoProducer(OBJ(c)).PrintExternalCall(Language.To_CallMethod(OBJ(self)),ToClass(OBJ(s)) )} 
  
// prints the n-th arg with a possible cast if necessary since we expect the type t (hence the class class!(t))
// n=0 is a special marker when the arg the receiver x in x.f(....)
// in that case we can do with the static_type because of Go polymorphism
/* The go function for: external_casted_arg(x:any,s:class,n:integer,nl?:boolean) [status=1] */
func F_Generate_external_casted_arg_any (x *ClaireAny ,s *ClaireClass ,n int,nl_ask *ClaireBoolean ) EID { 
    // eid body s = void
    var Result EID 
    { var st *ClaireClass  
      var try_1 EID 
      try_1 = Language.F_static_type_any(x)
      if ErrorIn(try_1) {Result = try_1
      } else {
      st = ToClass(OBJ(try_1))
      if (n > 1) { 
        PRINC(",")
        if (nl_ask == CTRUE) { 
          F_Generate_breakline_void()
          } 
        } 
      if ((n == 0) && 
          (ToType(st.Id()).Included(ToType(s.Id())) == CTRUE)) { 
        Result = Core.F_CALL(C_Generate_g_expression,ARGS(x.ToEID(),EID{st.Id(),0}))
        } else {
        Result = Core.F_CALL(C_Generate_g_expression,ARGS(x.ToEID(),EID{s.Id(),0}))
        } 
      }
      } 
    return Result} 
  
// The EID go function for: external_casted_arg @ any (throw: true) 
func E_Generate_external_casted_arg_any (x EID,s EID,n EID,nl_ask EID) EID { 
    return F_Generate_external_casted_arg_any(ANY(x),
      ToClass(OBJ(s)),
      INT(n),
      ToBoolean(OBJ(nl_ask)) )} 
  
//**********************************************************************
//*          Part 4: expression for structures                       *
//**********************************************************************
// this is an attempt to get rid of useless parenthesis without creating ambuiguous situations
// bounded_expression(x,loop) adds wrapping ( ) if needed     ==     bounded expression :)
// here we assume that native is needed
/* The go function for: bounded_expression(self:any,s:class) [status=1] */
func F_Generate_bounded_expression_any (self *ClaireAny ,s *ClaireClass ) EID { 
    // eid body s = void
    var Result EID 
    if (self.Isa.IsIn(Language.C_Assign) == CTRUE) { 
      { var g0112 *Language.Assign   = Language.To_Assign(self)
        PRINC("(")
        Result = F_Generate_g_expression_any(g0112.Id(),s)
        if !ErrorIn(Result) {
        PRINC(")")
        Result = EVOID
        }
        } 
      }  else if (C_integer.Id() == self.Isa.Id()) { 
      { var g0113 int  = ToInteger(self).Value
        if (g0113 < 0) { 
          PRINC("(")
          Result = F_Generate_g_expression_integer(g0113,s)
          if !ErrorIn(Result) {
          PRINC(")")
          Result = EVOID
          }
          } else {
          Result = F_Generate_g_expression_integer(g0113,s)
          } 
        } 
      }  else if (C_float.Id() == self.Isa.Id()) { 
      { var g0114 float64  = ToFloat(self).Value
        if (g0114 < 0) { 
          PRINC("(")
          Result = F_Generate_g_expression_float(g0114,s)
          if !ErrorIn(Result) {
          PRINC(")")
          Result = EVOID
          }
          } else {
          Result = F_Generate_g_expression_float(g0114,s)
          } 
        } 
      } else {
      Result = Core.F_CALL(C_Generate_g_expression,ARGS(self.ToEID(),EID{s.Id(),0}))
      } 
    return Result} 
  
// The EID go function for: bounded_expression @ any (throw: true) 
func E_Generate_bounded_expression_any (self EID,s EID) EID { 
    return F_Generate_bounded_expression_any(ANY(self),ToClass(OBJ(s)) )} 
  
// if can be represented by an expression if the two arguments are constants (evaluation does not cost)
/* The go function for: g_expression(self:If,s:class) [status=1] */
func F_Generate_g_expression_If (self *Language.If ,s *ClaireClass ) EID { 
    // eid body s = void
    var Result EID 
    F_Generate_object_prefix_class(C_any,s)
    PRINC("IfThenElse(")
    Result = Core.F_CALL(Optimize.C_Compile_bool_exp,ARGS(self.Test.ToEID(),EID{CTRUE.Id(),0}))
    if !ErrorIn(Result) {
    PRINC(",")
    Result = EVOID
    }
    if !ErrorIn(Result) {
    Optimize.C_OPT.Level = (Optimize.C_OPT.Level+1)
    F_Generate_breakline_void()
    Result = Core.F_CALL(C_Generate_g_expression,ARGS(self.Arg.ToEID(),EID{C_any.Id(),0}))
    if !ErrorIn(Result) {
    PRINC(",")
    Result = EVOID
    }
    if !ErrorIn(Result) {
    F_Generate_breakline_void()
    Result = Core.F_CALL(C_Generate_g_expression,ARGS(self.Other.ToEID(),EID{C_any.Id(),0}))
    if !ErrorIn(Result) {
    PRINC(")")
    Result = EVOID
    }
    if !ErrorIn(Result) {
    F_Generate_object_post_class(C_any,s)
    { 
      var va_arg1 *Optimize.OptimizeMetaOPT  
      var va_arg2 int 
      va_arg1 = Optimize.C_OPT
      va_arg2 = (Optimize.C_OPT.Level-1)
      va_arg1.Level = va_arg2
      Result = EID{C__INT,IVAL(va_arg2)}
      } 
    }}}
    return Result} 
  
// The EID go function for: g_expression @ If (throw: true) 
func E_Generate_g_expression_If (self EID,s EID) EID { 
    return F_Generate_g_expression_If(Language.To_If(OBJ(self)),ToClass(OBJ(s)) )} 
  
// a conjunction is also a C expression
// note that go requires && before the line break hence the more complex code
/* The go function for: g_expression(self:And,s:class) [status=1] */
func F_Generate_g_expression_And (self *Language.And ,s *ClaireClass ) EID { 
    // eid body s = void
    var Result EID 
    { var b *ClaireBoolean   = Core.F__sup_integer(self.Args.Length(),5)
      { var n int  = self.Args.Length()
        F_Generate_object_prefix_class(C_boolean,s)
        PRINC("MakeBoolean(")
        { var i int  = 1
          { var g0118 int  = n
            Result= EID{CFALSE.Id(),0}
            for (i <= g0118) { 
              var loop_1 EID 
              _ = loop_1
              { 
              { var x *ClaireAny   = self.Args.At(i-1)
                loop_1 = Core.F_CALL(Optimize.C_Compile_bool_exp,ARGS(x.ToEID(),EID{CTRUE.Id(),0}))
                if ErrorIn(loop_1) {Result = loop_1
                break
                } else {
                if (i < n) { 
                  PRINC(" && ")
                  if (b == CTRUE) { 
                    loop_1 = F_Generate_breakline_void().ToEID()
                    } else {
                    loop_1 = EID{CFALSE.Id(),0}
                    } 
                  } else {
                  loop_1 = EID{CFALSE.Id(),0}
                  } 
                }
                } 
              if ErrorIn(loop_1) {Result = loop_1
              break
              } else {
              i = (i+1)
              }
              } 
            }
            } 
          } 
        if !ErrorIn(Result) {
        PRINC(")")
        F_Generate_object_post_class(C_boolean,s)
        Result = EVOID
        }
        } 
      } 
    return Result} 
  
// The EID go function for: g_expression @ And (throw: true) 
func E_Generate_g_expression_And (self EID,s EID) EID { 
    return F_Generate_g_expression_And(Language.To_And(OBJ(self)),ToClass(OBJ(s)) )} 
  
// same thing for a disjunction
/* The go function for: g_expression(self:Or,s:class) [status=1] */
func F_Generate_g_expression_Or (self *Language.Or ,s *ClaireClass ) EID { 
    // eid body s = void
    var Result EID 
    { var b *ClaireBoolean   = Core.F__sup_integer(self.Args.Length(),5)
      { var n int  = self.Args.Length()
        F_Generate_object_prefix_class(C_boolean,s)
        PRINC("MakeBoolean(")
        { var i int  = 1
          { var g0119 int  = n
            Result= EID{CFALSE.Id(),0}
            for (i <= g0119) { 
              var loop_1 EID 
              _ = loop_1
              { 
              { var x *ClaireAny   = self.Args.At(i-1)
                loop_1 = Core.F_CALL(Optimize.C_Compile_bool_exp,ARGS(x.ToEID(),EID{CTRUE.Id(),0}))
                if ErrorIn(loop_1) {Result = loop_1
                break
                } else {
                if (i < n) { 
                  PRINC(" || ")
                  if (b == CTRUE) { 
                    loop_1 = F_Generate_breakline_void().ToEID()
                    } else {
                    loop_1 = EID{CFALSE.Id(),0}
                    } 
                  } else {
                  loop_1 = EID{CFALSE.Id(),0}
                  } 
                }
                } 
              if ErrorIn(loop_1) {Result = loop_1
              break
              } else {
              i = (i+1)
              }
              } 
            }
            } 
          } 
        if !ErrorIn(Result) {
        PRINC(")")
        F_Generate_object_post_class(C_boolean,s)
        Result = EVOID
        }
        } 
      } 
    return Result} 
  
// The EID go function for: g_expression @ Or (throw: true) 
func E_Generate_g_expression_Or (self EID,s EID) EID { 
    return F_Generate_g_expression_Or(Language.To_Or(OBJ(self)),ToClass(OBJ(s)) )} 
  
// to_CL(x) produces a CLAIRE id from an external representation
// [g_expression(self:Generate/to_CL,s:class) : void
//  -> //[5] toCL -> ~S:~S // self.arg, owner(self.arg),
//    g_expression(self.arg, s)]
// to_C(x) produces an external representation from a CLAIRE id
// g_expression(self:Generate/to_C,s:class) : void
// -> g_expression(self.arg, s)
// C_cast(x) produces a cast for go  => unclear if it is still needed
/* The go function for: g_expression(self:Compile/C_cast,s:class) [status=1] */
func F_Generate_g_expression_C_cast (self *Optimize.Compile_CCast ,s *ClaireClass ) EID { 
    // eid body s = void
    var Result EID 
    Result = Core.F_CALL(C_Generate_g_expression,ARGS(self.Arg.ToEID(),EID{s.Id(),0}))
    return Result} 
  
// The EID go function for: g_expression @ Compile/C_cast (throw: true) 
func E_Generate_g_expression_C_cast (self EID,s EID) EID { 
    return F_Generate_g_expression_C_cast(Optimize.To_Compile_CCast(OBJ(self)),ToClass(OBJ(s)) )} 
  
// reads a slot : more complex that it looks
// when the test is on, we produce x.p.KNOWN(p) To transform CNULL into an error 
// because slots can be native, we need the generic pre/post to convert to the proper slot
/* The go function for: g_expression(self:Call_slot,s:class) [status=1] */
func F_Generate_g_expression_Call_slot (self *Language.CallSlot ,s *ClaireClass ) EID { 
    // eid body s = void
    var Result EID 
    { var sc *ClaireClass   = F_Generate_rootSlot_slot(self.Selector).Range.Class_I()
      { var dc *ClaireClass  
        var try_1 EID 
        try_1 = Language.F_static_type_any(self.Arg)
        if ErrorIn(try_1) {Result = try_1
        } else {
        dc = ToClass(OBJ(try_1))
        { var s2 *ClaireClass  
          if (ToType(dc.Id()).Included(ToType(Core.F_domain_I_restriction(ToRestriction(self.Selector.Id())).Id())) == CTRUE) { 
            s2 = dc
            } else {
            s2 = Core.F_domain_I_restriction(ToRestriction(self.Selector.Id())).Class_I()
            } 
          { var kt_ask *ClaireBoolean   = MakeBoolean((self.Test.Id() != CNULL) && (self.Test == CTRUE))
            if (kt_ask != CTRUE) { 
              Result = F_Generate_cast_prefix_class(sc,s)
              } else {
              Result = EID{CFALSE.Id(),0}
              } 
            if !ErrorIn(Result) {
            Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).CMember(self.Arg,s2,self.Selector.Selector)
            if !ErrorIn(Result) {
            if (kt_ask == CTRUE) { 
              PRINC(".KNOWN(")
              Result = F_Generate_g_expression_thing(ToThing(self.Selector.Selector.Id()),C_any)
              if !ErrorIn(Result) {
              PRINC(")")
              Result = EVOID
              }
              } else {
              F_Generate_cast_post_class(sc,s)
              Result = EVOID
              } 
            }}
            } 
          } 
        }
        } 
      } 
    return Result} 
  
// The EID go function for: g_expression @ Call_slot (throw: true) 
func E_Generate_g_expression_Call_slot (self EID,s EID) EID { 
    return F_Generate_g_expression_Call_slot(Language.To_CallSlot(OBJ(self)),ToClass(OBJ(s)) )} 
  
// reads an (integer) table  = WARNING - this will change in the future when tables are implemented with dictionaries
// here we  assume that the table uses a list ....
/* The go function for: g_expression(self:Call_table,s:class) [status=1] */
func F_Generate_g_expression_Call_table (self *Language.CallTable ,s *ClaireClass ) EID { 
    // eid body s = void
    var Result EID 
    { var a *ClaireTable   = self.Selector
      { var p *ClaireAny   = a.Params
        _ = p
        { var l *ClaireAny   = self.Arg
          if (a.Range.Included(ToType(C_integer.Id())) == CTRUE) { 
            Result = F_Generate_cast_prefix_class(C_integer,s)
            if !ErrorIn(Result) {
            PRINC("ToList(")
            Result = F_Generate_g_expression_thing(ToThing(a.Id()),C_table)
            if !ErrorIn(Result) {
            PRINC(".Graph).ValuesI()[")
            Result = F_Generate_g_table_index_table(a,l)
            if !ErrorIn(Result) {
            PRINC("-1]")
            Result = EVOID
            }}
            if !ErrorIn(Result) {
            F_Generate_cast_post_class(C_integer,s)
            Result = EVOID
            }}
            }  else if (a.Range.Id() == C_float.Id()) { 
            Result = F_Generate_cast_prefix_class(C_float,s)
            if !ErrorIn(Result) {
            PRINC("ToList(")
            Result = F_Generate_g_expression_thing(ToThing(a.Id()),C_table)
            if !ErrorIn(Result) {
            PRINC(".Graph).ValuesF()[")
            Result = F_Generate_g_table_index_table(a,l)
            if !ErrorIn(Result) {
            PRINC("-1]")
            Result = EVOID
            }}
            if !ErrorIn(Result) {
            F_Generate_cast_post_class(C_float,s)
            Result = EVOID
            }}
            } else {
            F_Generate_object_prefix_class(C_any,s)
            PRINC("ToList(")
            Result = F_Generate_g_expression_thing(ToThing(a.Id()),C_table)
            if !ErrorIn(Result) {
            PRINC(".Graph).At(")
            Result = F_Generate_g_table_index_table(a,l)
            if !ErrorIn(Result) {
            PRINC("-1)")
            Result = EVOID
            }}
            if !ErrorIn(Result) {
            if (self.Test == CTRUE) { 
              PRINC(".KNOWN(")
              Result = F_Generate_g_expression_thing(ToThing(a.Id()),C_any)
              if !ErrorIn(Result) {
              PRINC(")")
              Result = EVOID
              }
              } else {
              F_Generate_object_post_class(C_any,s)
              Result = EVOID
              } 
            }
            } 
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: g_expression @ Call_table (throw: true) 
func E_Generate_g_expression_Call_table (self EID,s EID) EID { 
    return F_Generate_g_expression_Call_table(Language.To_CallTable(OBJ(self)),ToClass(OBJ(s)) )} 
  
// printf the code to access the index 
/* The go function for: g_table_index(a:table,l:any) [status=1] */
func F_Generate_g_table_index_table (a *ClaireTable ,l *ClaireAny ) EID { 
    // eid body s = void
    var Result EID 
    { var p *ClaireAny   = a.Params
      if (C_integer.Id() == p.Isa.Id()) { 
        { var g0120 int  = ToInteger(p).Value
          Result = Core.F_CALL(C_Generate_g_expression,ARGS(l.ToEID(),EID{C_integer.Id(),0}))
          if !ErrorIn(Result) {
          PRINC(" - ")
          F_princ_integer(g0120)
          Result = EVOID
          }
          } 
        }  else if (p.Isa.IsIn(C_list) == CTRUE) { 
        { var g0121 *ClaireList   = ToList(p)
          if (l.Isa.IsIn(Language.C_List) != CTRUE) { 
            Result = ToException(Core.C_general_error.Make(MakeString("shit with call_table ~S[~S]").Id(),MakeConstantList(a.Id(),l).Id())).Close()
            } else {
            Result = EID{CFALSE.Id(),0}
            } 
          if !ErrorIn(Result) {
          Result = Core.F_CALL(C_Generate_g_expression,ARGS(ToList(OBJ(Core.F_CALL(C_args,ARGS(l.ToEID())))).At(1-1).ToEID(),EID{C_integer.Id(),0}))
          if !ErrorIn(Result) {
          PRINC(" * ")
          Result = Core.F_CALL(C_princ,ARGS(g0121.At(1-1).ToEID()))
          if !ErrorIn(Result) {
          PRINC(" + ")
          Result = Core.F_CALL(C_Generate_g_expression,ARGS(ToList(OBJ(Core.F_CALL(C_args,ARGS(l.ToEID())))).At(2-1).ToEID(),EID{C_integer.Id(),0}))
          if !ErrorIn(Result) {
          PRINC(" - ")
          Result = Core.F_CALL(C_princ,ARGS(g0121.At(2-1).ToEID()))
          }}}
          }
          } 
        } else {
        Result = EID{CFALSE.Id(),0}
        } 
      } 
    return Result} 
  
// The EID go function for: g_table_index @ table (throw: true) 
func E_Generate_g_table_index_table (a EID,l EID) EID { 
    return F_Generate_g_table_index_table(ToTable(OBJ(a)),ANY(l) )} 
  
// reads an array - remember that in CLAIRE 4, arrays are nothing but fixed size lists (with 3 sorts)
/* The go function for: g_expression(self:Call_array,s:class) [status=1] */
func F_Generate_g_expression_Call_array (self *Language.CallArray ,s *ClaireClass ) EID { 
    // eid body s = void
    var Result EID 
    { var sa *ClaireClass  
      var try_1 EID 
      { var arg_2 *ClaireType  
        var try_3 EID 
        { var arg_4 *ClaireType  
          var try_5 EID 
          try_5 = Core.F_CALL(Optimize.C_c_type,ARGS(self.Selector.ToEID()))
          if ErrorIn(try_5) {try_3 = try_5
          } else {
          arg_4 = ToType(OBJ(try_5))
          try_3 = EID{Core.F_member_type(arg_4).Id(),0}
          }
          } 
        if ErrorIn(try_3) {try_1 = try_3
        } else {
        arg_2 = ToType(OBJ(try_3))
        try_1 = EID{F_Generate_type_sort_type(arg_2).Id(),0}
        }
        } 
      if ErrorIn(try_1) {Result = try_1
      } else {
      sa = ToClass(OBJ(try_1))
      { var sm *ClaireClass  
        var try_6 EID 
        try_6 = F_Generate_g_member_any(self.Selector)
        if ErrorIn(try_6) {Result = try_6
        } else {
        sm = ToClass(OBJ(try_6))
        Result = F_Generate_cast_prefix_class(sa,s)
        if !ErrorIn(Result) {
        if (sm.Id() != C_any.Id()) { 
          Result = Core.F_CALL(C_Generate_g_expression,ARGS(self.Selector.ToEID(),EID{C_list.Id(),0}))
          if !ErrorIn(Result) {
          PRINC(".")
          F_Generate_valuesSlot_class(sm)
          PRINC("[")
          Result = Core.F_CALL(C_Generate_g_expression,ARGS(self.Arg.ToEID(),EID{C_integer.Id(),0}))
          if !ErrorIn(Result) {
          PRINC(" - 1]")
          Result = EVOID
          }}
          } else {
          Result = Core.F_CALL(C_Generate_g_expression,ARGS(self.Selector.ToEID(),EID{C_list.Id(),0}))
          if !ErrorIn(Result) {
          PRINC(".At(")
          Result = Core.F_CALL(C_Generate_g_expression,ARGS(self.Arg.ToEID(),EID{C_integer.Id(),0}))
          if !ErrorIn(Result) {
          PRINC(" - 1)")
          Result = EVOID
          }}
          } 
        if !ErrorIn(Result) {
        F_Generate_cast_post_class(sa,s)
        Result = EVOID
        }}
        }
        } 
      }
      } 
    return Result} 
  
// The EID go function for: g_expression @ Call_array (throw: true) 
func E_Generate_g_expression_Call_array (self EID,s EID) EID { 
    return F_Generate_g_expression_Call_array(Language.To_CallArray(OBJ(self)),ToClass(OBJ(s)) )} 
  
//**********************************************************************
//*          Part 5: the logical expression compilation                *
//**********************************************************************
// bool_exp(x,pos?) returns a native boolean go expression, assumes that g_func(x) !
// bool_expression(x) could be g_expression(x,boolean)
// however, boolean are not native in CLAIRE4 () to avoid conversions
// note : we drop bool_exp? and bool_exp!
// this is the boolean compiler. An automatic computation of negation is
// included. The flag pos? tells if the assertion is positive. When a
// negation occurs, we simply change the flag. At the end of compiling,
// the flag is used to generate == or != according to this method:
// generate the = or /=
/* The go function for: sign_equal(self:boolean) [status=0] */
func F_Generate_sign_equal_boolean (self *ClaireBoolean )  { 
    // procedure body with s = void
    if (self == CTRUE) { 
      PRINC("==")
      } else {
      PRINC("!=")
      } 
    } 
  
// The EID go function for: sign_equal @ boolean (throw: false) 
func E_Generate_sign_equal_boolean (self EID) EID { 
    F_Generate_sign_equal_boolean(ToBoolean(OBJ(self)) )
    return EVOID} 
  
// generate a conjunction/disjunction
/* The go function for: sign_or(self:boolean) [status=0] */
func F_Generate_sign_or_boolean (self *ClaireBoolean )  { 
    // procedure body with s = void
    if (self == CTRUE) { 
      PRINC("||")
      } else {
      PRINC("&&")
      } 
    } 
  
// The EID go function for: sign_or @ boolean (throw: false) 
func E_Generate_sign_or_boolean (self EID) EID { 
    F_Generate_sign_or_boolean(ToBoolean(OBJ(self)) )
    return EVOID} 
  
// default solution
/* The go function for: Compile/bool_exp(self:any,pos?:boolean) [status=1] */
func F_Compile_bool_exp_any (self *ClaireAny ,pos_ask *ClaireBoolean ) EID { 
    // eid body s = void
    var Result EID 
    PRINC("(")
    Result = Core.F_CALL(C_Generate_g_expression,ARGS(self.ToEID(),EID{C_boolean.Id(),0}))
    if !ErrorIn(Result) {
    PRINC(" ")
    F_Generate_sign_equal_boolean(pos_ask)
    PRINC(" CTRUE)")
    Result = EVOID
    }
    return Result} 
  
// The EID go function for: Compile/bool_exp @ any (throw: true) 
func E_Compile_bool_exp_any (self EID,pos_ask EID) EID { 
    return F_Compile_bool_exp_any(ANY(self),ToBoolean(OBJ(pos_ask)) )} 
  
// strange : not clear why we should see a C_cast here
/* The go function for: Compile/bool_exp(self:Compile/C_cast,pos?:boolean) [status=1] */
func F_Compile_bool_exp_C_cast (self *Optimize.Compile_CCast ,pos_ask *ClaireBoolean ) EID { 
    // eid body s = void
    var Result EID 
    Result = Core.F_CALL(Optimize.C_Compile_bool_exp,ARGS(self.Arg.ToEID(),EID{pos_ask.Id(),0}))
    return Result} 
  
// The EID go function for: Compile/bool_exp @ Compile/C_cast (throw: true) 
func E_Compile_bool_exp_C_cast (self EID,pos_ask EID) EID { 
    return F_Compile_bool_exp_C_cast(Optimize.To_Compile_CCast(OBJ(self)),ToBoolean(OBJ(pos_ask)) )} 
  
// if we have a CL, we know that the self.arg is of type boolean
// [bool_exp(self:Generate/to_CL,pos?:boolean) : void
//  -> bool_exp(self.arg,pos?) ]
// If is supported with IfThenElse (means that all terms will be evaluated),
/* The go function for: Compile/bool_exp(self:If,pos?:boolean) [status=1] */
func F_Compile_bool_exp_If (self *Language.If ,pos_ask *ClaireBoolean ) EID { 
    // eid body s = void
    var Result EID 
    if (F_boolean_I_any(self.Other) == CTRUE) { 
      PRINC("(")
      Result = Core.F_CALL(Optimize.C_Compile_bool_exp,ARGS(self.Test.ToEID(),EID{CTRUE.Id(),0}))
      if !ErrorIn(Result) {
      PRINC(" ? ")
      Result = Core.F_CALL(Optimize.C_Compile_bool_exp,ARGS(self.Arg.ToEID(),EID{pos_ask.Id(),0}))
      if !ErrorIn(Result) {
      PRINC(" : ")
      Result = Core.F_CALL(Optimize.C_Compile_bool_exp,ARGS(self.Other.ToEID(),EID{pos_ask.Id(),0}))
      if !ErrorIn(Result) {
      PRINC(")")
      Result = EVOID
      }}}
      } else {
      PRINC("(")
      Result = Core.F_CALL(Optimize.C_Compile_bool_exp,ARGS(self.Test.ToEID(),EID{pos_ask.Id(),0}))
      if !ErrorIn(Result) {
      PRINC(" ")
      F_Generate_sign_or_boolean(pos_ask.Not)
      PRINC(" ")
      Result = Core.F_CALL(Optimize.C_Compile_bool_exp,ARGS(self.Arg.ToEID(),EID{pos_ask.Id(),0}))
      if !ErrorIn(Result) {
      PRINC(")")
      Result = EVOID
      }}
      } 
    return Result} 
  
// The EID go function for: Compile/bool_exp @ If (throw: true) 
func E_Compile_bool_exp_If (self EID,pos_ask EID) EID { 
    return F_Compile_bool_exp_If(Language.To_If(OBJ(self)),ToBoolean(OBJ(pos_ask)) )} 
  
// for a AND, we can used the && C operation
/* The go function for: Compile/bool_exp(self:And,pos?:boolean) [status=1] */
func F_Compile_bool_exp_And (self *Language.And ,pos_ask *ClaireBoolean ) EID { 
    // eid body s = void
    var Result EID 
    { var l *ClaireList   = self.Args
      { var m int  = l.Length()
        { var n int  = 0
          { var _Zl int  = Optimize.C_OPT.Level
            Optimize.C_OPT.Level = (Optimize.C_OPT.Level+1)
            { 
              var x *ClaireAny  
              _ = x
              Result= EID{CFALSE.Id(),0}
              var x_support *ClaireList  
              x_support = l
              x_len := x_support.Length()
              for i_it := 0; i_it < x_len; i_it++ { 
                x = x_support.At(i_it)
                var loop_1 EID 
                _ = loop_1
                { 
                n = (n+1)
                if (n == m) { 
                  loop_1 = Core.F_CALL(Optimize.C_Compile_bool_exp,ARGS(x.ToEID(),EID{pos_ask.Id(),0}))
                  } else {
                  PRINC("(")
                  loop_1 = Core.F_CALL(Optimize.C_Compile_bool_exp,ARGS(x.ToEID(),EID{pos_ask.Id(),0}))
                  if ErrorIn(loop_1) {Result = loop_1
                  break
                  } else {
                  PRINC(" ")
                  F_Generate_sign_or_boolean(pos_ask.Not)
                  PRINC(" ")
                  loop_1 = EVOID
                  }
                  if ErrorIn(loop_1) {Result = loop_1
                  break
                  } else {
                  Optimize.C_OPT.Level = (Optimize.C_OPT.Level+1)
                  loop_1 = F_Generate_breakline_void().ToEID()
                  }
                  } 
                if ErrorIn(loop_1) {Result = loop_1
                break
                } else {
                }
                }
                } 
              } 
            if !ErrorIn(Result) {
            { var x int  = 2
              { var g0124 int  = m
                for (x <= g0124) { 
                  PRINC(")")
                  x = (x+1)
                  } 
                } 
              } 
            { 
              var va_arg1 *Optimize.OptimizeMetaOPT  
              var va_arg2 int 
              va_arg1 = Optimize.C_OPT
              va_arg2 = _Zl
              va_arg1.Level = va_arg2
              Result = EID{C__INT,IVAL(va_arg2)}
              } 
            }
            } 
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: Compile/bool_exp @ And (throw: true) 
func E_Compile_bool_exp_And (self EID,pos_ask EID) EID { 
    return F_Compile_bool_exp_And(Language.To_And(OBJ(self)),ToBoolean(OBJ(pos_ask)) )} 
  
// idem for OR: we use ||
/* The go function for: Compile/bool_exp(self:Or,pos?:boolean) [status=1] */
func F_Compile_bool_exp_Or (self *Language.Or ,pos_ask *ClaireBoolean ) EID { 
    // eid body s = void
    var Result EID 
    { var l *ClaireList   = self.Args
      { var m int  = l.Length()
        { var n int  = 0
          { var _Zl int  = Optimize.C_OPT.Level
            Optimize.C_OPT.Level = (Optimize.C_OPT.Level+1)
            { 
              var x *ClaireAny  
              _ = x
              Result= EID{CFALSE.Id(),0}
              var x_support *ClaireList  
              x_support = l
              x_len := x_support.Length()
              for i_it := 0; i_it < x_len; i_it++ { 
                x = x_support.At(i_it)
                var loop_1 EID 
                _ = loop_1
                { 
                n = (n+1)
                if (n == m) { 
                  loop_1 = Core.F_CALL(Optimize.C_Compile_bool_exp,ARGS(x.ToEID(),EID{pos_ask.Id(),0}))
                  } else {
                  PRINC("(")
                  loop_1 = Core.F_CALL(Optimize.C_Compile_bool_exp,ARGS(x.ToEID(),EID{pos_ask.Id(),0}))
                  if ErrorIn(loop_1) {Result = loop_1
                  break
                  } else {
                  PRINC(" ")
                  F_Generate_sign_or_boolean(pos_ask)
                  PRINC(" ")
                  loop_1 = EVOID
                  }
                  if ErrorIn(loop_1) {Result = loop_1
                  break
                  } else {
                  Optimize.C_OPT.Level = (Optimize.C_OPT.Level+1)
                  loop_1 = F_Generate_breakline_void().ToEID()
                  }
                  } 
                if ErrorIn(loop_1) {Result = loop_1
                break
                } else {
                }
                }
                } 
              } 
            if !ErrorIn(Result) {
            { var x int  = 2
              { var g0127 int  = m
                for (x <= g0127) { 
                  PRINC(")")
                  x = (x+1)
                  } 
                } 
              } 
            { 
              var va_arg1 *Optimize.OptimizeMetaOPT  
              var va_arg2 int 
              va_arg1 = Optimize.C_OPT
              va_arg2 = _Zl
              va_arg1.Level = va_arg2
              Result = EID{C__INT,IVAL(va_arg2)}
              } 
            }
            } 
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: Compile/bool_exp @ Or (throw: true) 
func E_Compile_bool_exp_Or (self EID,pos_ask EID) EID { 
    return F_Compile_bool_exp_Or(Language.To_Or(OBJ(self)),ToBoolean(OBJ(pos_ask)) )} 
  
// membership
/* The go function for: Compile/bool_exp(self:Call,pos?:boolean) [status=1] */
func F_Compile_bool_exp_Call (self *Language.Call ,pos_ask *ClaireBoolean ) EID { 
    // eid body s = void
    var Result EID 
    { var p *ClaireProperty   = self.Selector
      if (p.Id() == C__Z.Id()) { 
        PRINC("(")
        Result = F_Generate_belong_exp_any(self.Args.At(1-1),self.Args.At(2-1),C_boolean)
        if !ErrorIn(Result) {
        PRINC(" ")
        F_Generate_sign_equal_boolean(pos_ask)
        PRINC(" CTRUE)")
        Result = EVOID
        }
        } else {
        Result = Core.F_SUPER(Optimize.C_Compile_bool_exp, C_any, ARGS(EID{self.Id(),0},EID{pos_ask.Id(),0}))
        } 
      } 
    return Result} 
  
// The EID go function for: Compile/bool_exp @ Call (throw: true) 
func E_Compile_bool_exp_Call (self EID,pos_ask EID) EID { 
    return F_Compile_bool_exp_Call(Language.To_Call(OBJ(self)),ToBoolean(OBJ(pos_ask)) )} 
  
// compile (a % ..), s is always a boolean but for EID mode
// the notOpt() test in gostat.cl ensures that the first three cases are seen as not-throw (not EID)
// however this fragment may be called to return an EID hence the global wrap with prefix/post
/* The go function for: belong_exp(a1:any,a2:any,s:class) [status=1] */
func F_Generate_belong_exp_any (a1 *ClaireAny ,a2 *ClaireAny ,s *ClaireClass ) EID { 
    // eid body s = void
    var Result EID 
    var g0128I *ClaireBoolean  
    var try_1 EID 
    { var arg_2 *ClaireClass  
      var try_3 EID 
      try_3 = Language.F_static_type_any(a2)
      if ErrorIn(try_3) {try_1 = try_3
      } else {
      arg_2 = ToClass(OBJ(try_3))
      try_1 = EID{ToType(arg_2.Id()).Included(ToType(C_type.Id())).Id(),0}
      }
      } 
    if ErrorIn(try_1) {Result = try_1
    } else {
    g0128I = ToBoolean(OBJ(try_1))
    if (g0128I == CTRUE) { 
      Result = F_Generate_cast_prefix_class(C_boolean,s)
      if !ErrorIn(Result) {
      Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{C_type.Id(),0}))
      if !ErrorIn(Result) {
      PRINC(".Contains(")
      Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_any.Id(),0}))
      if !ErrorIn(Result) {
      PRINC(")")
      F_Generate_cast_post_class(C_boolean,s)
      Result = EVOID
      }}}
      } else {
      var g0129I *ClaireBoolean  
      var try_4 EID 
      { 
        var v_and3 *ClaireBoolean  
        
        var try_5 EID 
        { var arg_6 *ClaireClass  
          var try_7 EID 
          try_7 = Language.F_static_type_any(a2)
          if ErrorIn(try_7) {try_5 = try_7
          } else {
          arg_6 = ToClass(OBJ(try_7))
          try_5 = EID{ToType(arg_6.Id()).Included(ToType(C_integer.Id())).Id(),0}
          }
          } 
        if ErrorIn(try_5) {try_4 = try_5
        } else {
        v_and3 = ToBoolean(OBJ(try_5))
        if (v_and3 == CFALSE) {try_4 = EID{CFALSE.Id(),0}
        } else { 
          var try_8 EID 
          { var arg_9 *ClaireClass  
            var try_10 EID 
            try_10 = Language.F_static_type_any(a1)
            if ErrorIn(try_10) {try_8 = try_10
            } else {
            arg_9 = ToClass(OBJ(try_10))
            try_8 = EID{ToType(arg_9.Id()).Included(ToType(C_integer.Id())).Id(),0}
            }
            } 
          if ErrorIn(try_8) {try_4 = try_8
          } else {
          v_and3 = ToBoolean(OBJ(try_8))
          if (v_and3 == CFALSE) {try_4 = EID{CFALSE.Id(),0}
          } else { 
            try_4 = EID{CTRUE.Id(),0}} 
          } 
        }}
        } 
      if ErrorIn(try_4) {Result = try_4
      } else {
      g0129I = ToBoolean(OBJ(try_4))
      if (g0129I == CTRUE) { 
        Result = F_Generate_cast_prefix_class(C_boolean,s)
        if !ErrorIn(Result) {
        PRINC("BitVectorContains(")
        Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{C_integer.Id(),0}))
        if !ErrorIn(Result) {
        PRINC(",")
        Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_integer.Id(),0}))
        if !ErrorIn(Result) {
        PRINC(")")
        F_Generate_cast_post_class(C_boolean,s)
        Result = EVOID
        }}}
        } else {
        var g0130I *ClaireBoolean  
        var try_11 EID 
        { 
          var v_or4 *ClaireBoolean  
          
          var try_12 EID 
          { var arg_13 *ClaireClass  
            var try_14 EID 
            try_14 = Language.F_static_type_any(a2)
            if ErrorIn(try_14) {try_12 = try_14
            } else {
            arg_13 = ToClass(OBJ(try_14))
            try_12 = EID{ToType(arg_13.Id()).Included(ToType(C_list.Id())).Id(),0}
            }
            } 
          if ErrorIn(try_12) {try_11 = try_12
          } else {
          v_or4 = ToBoolean(OBJ(try_12))
          if (v_or4 == CTRUE) {try_11 = EID{CTRUE.Id(),0}
          } else { 
            var try_15 EID 
            { var arg_16 *ClaireClass  
              var try_17 EID 
              try_17 = Language.F_static_type_any(a2)
              if ErrorIn(try_17) {try_15 = try_17
              } else {
              arg_16 = ToClass(OBJ(try_17))
              try_15 = EID{ToType(arg_16.Id()).Included(ToType(C_array.Id())).Id(),0}
              }
              } 
            if ErrorIn(try_15) {try_11 = try_15
            } else {
            v_or4 = ToBoolean(OBJ(try_15))
            if (v_or4 == CTRUE) {try_11 = EID{CTRUE.Id(),0}
            } else { 
              try_11 = EID{CFALSE.Id(),0}} 
            } 
          }}
          } 
        if ErrorIn(try_11) {Result = try_11
        } else {
        g0130I = ToBoolean(OBJ(try_11))
        if (g0130I == CTRUE) { 
          Result = F_Generate_cast_prefix_class(C_boolean,s)
          if !ErrorIn(Result) {
          Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{C_list.Id(),0}))
          if !ErrorIn(Result) {
          PRINC(".Contain_ask(")
          Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_any.Id(),0}))
          if !ErrorIn(Result) {
          PRINC(")")
          F_Generate_cast_post_class(C_boolean,s)
          Result = EVOID
          }}}
          } else {
          Result = F_Generate_cast_prefix_class(Optimize.C_EID,s)
          if !ErrorIn(Result) {
          F_Generate_preCore_ask_void()
          PRINC("F_BELONG(")
          Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_any.Id(),0}))
          if !ErrorIn(Result) {
          PRINC(",")
          Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{C_any.Id(),0}))
          if !ErrorIn(Result) {
          PRINC(")")
          F_Generate_cast_post_class(Optimize.C_EID,s)
          Result = EVOID
          }}}
          } 
        }
        } 
      }
      } 
    }
    return Result} 
  
// The EID go function for: belong_exp @ any (throw: true) 
func E_Generate_belong_exp_any (a1 EID,a2 EID,s EID) EID { 
    return F_Generate_belong_exp_any(ANY(a1),ANY(a2),ToClass(OBJ(s)) )} 
  
// some special functions are open coded when used in a logical test
/* The go function for: Compile/bool_exp(self:Call_method1,pos?:boolean) [status=1] */
func F_Compile_bool_exp_Call_method1 (self *Language.CallMethod1 ,pos_ask *ClaireBoolean ) EID { 
    // eid body s = void
    var Result EID 
    { var m *ClaireMethod   = self.Arg
      { var a1 *ClaireAny   = self.Args.At(1-1)
        if (m.Id() == C_Generate__starnot_star.Value) { 
          Result = Core.F_CALL(Optimize.C_Compile_bool_exp,ARGS(a1.ToEID(),EID{pos_ask.Not.Id(),0}))
          }  else if (m.Id() == C_Generate__starknown_star.Value) { 
          Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).EqualExp(a1,
            pos_ask.Not,
            CNULL,
            CTRUE.Id())
          }  else if (m.Id() == C_Generate__starunknown_star.Value) { 
          Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).EqualExp(a1,
            pos_ask,
            CNULL,
            CTRUE.Id())
          }  else if (m.Range.Included(ToType(C_boolean.Id())) == CTRUE) { 
          PRINC("(")
          Result = F_Generate_g_expression_Call_method1(self,C_boolean)
          if !ErrorIn(Result) {
          PRINC(" ")
          F_Generate_sign_equal_boolean(pos_ask)
          PRINC(" CTRUE)")
          Result = EVOID
          }
          } else {
          Result = Core.F_SUPER(Optimize.C_Compile_bool_exp, C_any, ARGS(EID{self.Id(),0},EID{pos_ask.Id(),0}))
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: Compile/bool_exp @ Call_method1 (throw: true) 
func E_Compile_bool_exp_Call_method1 (self EID,pos_ask EID) EID { 
    return F_Compile_bool_exp_Call_method1(Language.To_CallMethod1(OBJ(self)),ToBoolean(OBJ(pos_ask)) )} 
  
// same thing for two arguments functions
// equal_exp is in gogen.cl
/* The go function for: Compile/bool_exp(self:Call_method2,pos?:boolean) [status=1] */
func F_Compile_bool_exp_Call_method2 (self *Language.CallMethod2 ,pos_ask *ClaireBoolean ) EID { 
    // eid body s = void
    var Result EID 
    { var m *ClaireMethod   = self.Arg
      { var p *ClaireProperty   = m.Selector
        { var lop *ClaireList   = ToGenerateCodeProducer(Optimize.C_PRODUCER.Value).OpenComparators
          { var a1 *ClaireAny   = self.Args.At(1-1)
            { var a2 *ClaireAny   = self.Args.At(2-1)
              if (p.Id() == Core.C__I_equal.Id()) { 
                Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).EqualExp(a1,
                  pos_ask.Not,
                  a2,
                  CFALSE.Id())
                }  else if (p.Id() == Core.C_identical_ask.Id()) { 
                Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).EqualExp(a1,
                  pos_ask,
                  a2,
                  CTRUE.Id())
                }  else if (p.Id() == C__equal.Id()) { 
                Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).EqualExp(a1,
                  pos_ask,
                  a2,
                  CFALSE.Id())
                }  else if (m.Id() == Optimize.C_Compile_m_member.Value) { 
                PRINC("(")
                Result = F_Generate_belong_exp_any(a1,a2,C_boolean)
                if !ErrorIn(Result) {
                PRINC(" ")
                F_Generate_sign_equal_boolean(pos_ask)
                PRINC(" CTRUE)")
                Result = EVOID
                }
                }  else if ((lop.Memq(p.Id()) == CTRUE) && 
                  ((Core.F_domain_I_restriction(ToRestriction(m.Id())).Id() == C_integer.Id()) || 
                      (Core.F_domain_I_restriction(ToRestriction(m.Id())).Id() == C_float.Id()))) { 
                PRINC("(")
                Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{Core.F_domain_I_restriction(ToRestriction(m.Id())).Id(),0}))
                if !ErrorIn(Result) {
                PRINC(" ")
                if (pos_ask == CTRUE) { 
                  Result = Core.F_print_any(p.Id())
                  } else {
                  Result = Core.F_print_any(lop.At((((F_index_list(lop,p.Id())+1)%4)+1)-1))
                  } 
                if !ErrorIn(Result) {
                PRINC(" ")
                Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{Core.F_domain_I_restriction(ToRestriction(m.Id())).Id(),0}))
                if !ErrorIn(Result) {
                PRINC(")")
                Result = EVOID
                }}}
                }  else if (m.Id() == C_Generate__starnth_integer_star.Value) { 
                PRINC("(BitVectorContains(")
                Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_integer.Id(),0}))
                if !ErrorIn(Result) {
                PRINC(",")
                Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{C_integer.Id(),0}))
                if !ErrorIn(Result) {
                PRINC(") ")
                F_Generate_sign_equal_boolean(pos_ask)
                PRINC(" CTRUE)")
                Result = EVOID
                }}
                }  else if ((p.Id() == Core.C_inherit_ask.Id()) && 
                  (Core.F_domain_I_restriction(ToRestriction(m.Id())).Id() == C_class.Id())) { 
                PRINC("(")
                Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_class.Id(),0}))
                if !ErrorIn(Result) {
                PRINC(".IsIn(")
                Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{C_class.Id(),0}))
                if !ErrorIn(Result) {
                PRINC(") ")
                F_Generate_sign_equal_boolean(pos_ask)
                PRINC(" CTRUE)")
                Result = EVOID
                }}
                }  else if (m.Range.Included(ToType(C_boolean.Id())) == CTRUE) { 
                PRINC("(")
                Result = F_Generate_g_expression_Call_method2(self,C_boolean)
                if !ErrorIn(Result) {
                PRINC(" ")
                F_Generate_sign_equal_boolean(pos_ask)
                PRINC(" CTRUE)")
                Result = EVOID
                }
                } else {
                Result = Core.F_SUPER(Optimize.C_Compile_bool_exp, C_any, ARGS(EID{self.Id(),0},EID{pos_ask.Id(),0}))
                } 
              } 
            } 
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: Compile/bool_exp @ Call_method2 (throw: true) 
func E_Compile_bool_exp_Call_method2 (self EID,pos_ask EID) EID { 
    return F_Compile_bool_exp_Call_method2(Language.To_CallMethod2(OBJ(self)),ToBoolean(OBJ(pos_ask)) )} 
  