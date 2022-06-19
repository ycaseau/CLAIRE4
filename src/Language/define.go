/***** CLAIRE Compilation of file /Users/ycaseau/Dropbox/src/clairev4.05/src/meta/define.cl 
         [version 4.0.06 / safety 5] Monday 06-06-2022 08:31:22 *****/

package Language
import (_ "fmt"
	. "Kernel"
	"Core"
)

//-------- dumb function to prevent import errors --------
func import_g0212() { 
    _ = Core.It
    } 
  
  
//+-------------------------------------------------------------+
//| CLAIRE                                                      |
//| define.cl                                                   |
//| Copyright (C) 1994 - 2021 Yves Caseau. All Rights Reserved  |
//| cf. copyright info in file object.cl: about()               |
//+-------------------------------------------------------------+
// --------------------------------------------------------------
// this file contains all definition & instanciation instructions
//---------------------------------------------------------------
// **************************************************************************
// * Contents:                                                              *
// *     Part 1: Definition instructions (Defobj, Defclass, Defmethod ...)  *
// *     Part 2: the instantiation macros                                   *
// *     Part 3: the useful stuff                                           *
// *     Part 4: the other macros                                           *
// *     Part 5: OFTO for methods                                           *
// **************************************************************************
// *********************************************************************
// *     Part 1: Definition                                            *
// *********************************************************************
// this is the basic class instantiation
//
/* The go function for: self_print(self:Definition) [status=1] */
func (self *Definition) SelfPrint () EID { 
    var Result EID
    Result = Core.F_print_any(self.Arg.Id())
    if !ErrorIn(Result) {
    PRINC("(")
    Result = F_Language_printbox_list2(self.Args)
    if !ErrorIn(Result) {
    PRINC(")")
    Result = EVOID
    }}
    return Result} 
  
// The EID go function for: self_print @ Definition (throw: true) 
func E_self_print_Definition_Language (self EID) EID { 
    return To_Definition(OBJ(self)).SelfPrint( )} 
  
// CLAIRE 4: fast definition when no close nor facy slots
// ------------- named object definition ------------------------------
//
/* The go function for: self_print(self:Defobj) [status=1] */
func (self *Defobj) SelfPrint () EID { 
    var Result EID
    if (self.Arg.Id() == Core.C_global_variable.Id()) { 
      { var r *ClaireAny = C_any.Id()
        { var v *ClaireAny = CNULL
          { 
            var x *Call
            _ = x
            var x_iter *ClaireAny
            var x_support *ClaireList
            x_support = self.Args
            x_len := x_support.Length()
            for i_it := 0; i_it < x_len; i_it++ { 
              x_iter = x_support.At(i_it)
              x = To_Call(x_iter)
              if (x.Args.At(0) == C_value.Id()) { 
                v = x.Args.At(1)
                }  else if (x.Args.At(0) == C_range.Id()) { 
                r = x.Args.At(1)
                } 
              } 
            } 
          if (F_boolean_I_any(r) == CTRUE) { 
            self.Ident.Princ()
            PRINC(":")
            Result = Core.F_CALL(C_print,ARGS(r.ToEID()))
            if !ErrorIn(Result) {
            PRINC(" := ")
            Result = F_printexp_any(v,CFALSE)
            }
            } else {
            self.Ident.Princ()
            PRINC(" :: ")
            Result = F_printexp_any(v,CFALSE)
            } 
          } 
        } 
      } else {
      self.Ident.Princ()
      PRINC(" :: ")
      Result = Core.F_CALL(C_print,ARGS(EID{self.Arg.Id(),0}))
      if !ErrorIn(Result) {
      PRINC("(")
      Result = F_Language_printbox_list2(self.Args)
      if !ErrorIn(Result) {
      PRINC(")")
      Result = EVOID
      }}
      } 
    return Result} 
  
// The EID go function for: self_print @ Defobj (throw: true) 
func E_self_print_Defobj_Language (self EID) EID { 
    return To_Defobj(OBJ(self)).SelfPrint( )} 
  
// ------------- class definition ------------------------------------
//
/* The go function for: self_print(self:Defclass) [status=1] */
func (self *Defclass) SelfPrint () EID { 
    var Result EID
    if (self.Ident.Id() == CNULL) { 
      Result = Core.F_print_any(MakeString("<Defclass>").Id())
      } else {
      self.Ident.Princ()
      if (self.Params.Length() != 0) { 
        PRINC("[")
        Result = Core.F_princ_list(self.Params)
        if !ErrorIn(Result) {
        PRINC("]")
        Result = EVOID
        }
        } else {
        Result = EID{CFALSE.Id(),0}
        } 
      if !ErrorIn(Result) {
      PRINC(" <: ")
      Result = Core.F_print_any(self.Arg.Id())
      if !ErrorIn(Result) {
      PRINC("(")
      { var l *ClaireList = self.Args
        { var n int = l.Length()
          { var i int = 1
            { var g0213 int = n
              Result= EID{CFALSE.Id(),0}
              for (i <= g0213) { 
                var loop_1 EID
                _ = loop_1
                { 
                if (i == 1) { 
                  F_set_level_void()
                  loop_1 = EVOID
                  } else {
                  loop_1 = F_lbreak_void()
                  } 
                if ErrorIn(loop_1) {Result = loop_1
                break
                } else {
                if (l.At(i-1).Isa.IsIn(C_Vardef) == CTRUE) { 
                  loop_1 = Core.F_CALL(C_Language_ppvariable,ARGS(l.At(i-1).ToEID()))
                  } else {
                  loop_1 = Core.F_CALL(C_Language_ppvariable,ARGS(ToList(OBJ(Core.F_CALL(C_args,ARGS(l.At(i-1).ToEID())))).At(0).ToEID()))
                  if ErrorIn(loop_1) {Result = loop_1
                  break
                  } else {
                  PRINC(" = ")
                  loop_1 = Core.F_CALL(C_print,ARGS(ToList(OBJ(Core.F_CALL(C_args,ARGS(l.At(i-1).ToEID())))).At(1).ToEID()))
                  if ErrorIn(loop_1) {Result = loop_1
                  break
                  } else {
                  }}
                  } 
                if ErrorIn(loop_1) {Result = loop_1
                break
                } else {
                if (i < n) { 
                  PRINC(",")
                  loop_1 = EVOID
                  } else {
                  loop_1 = EID{CFALSE.Id(),0}
                  } 
                }}
                if ErrorIn(loop_1) {Result = loop_1
                break
                } else {
                i = (i+1)
                }
                } 
              }
              } 
            } 
          } 
        } 
      if !ErrorIn(Result) {
      PRINC(")")
      Result = EVOID
      }}}
      } 
    return Result} 
  
// The EID go function for: self_print @ Defclass (throw: true) 
func E_self_print_Defclass_Language (self EID) EID { 
    return To_Defclass(OBJ(self)).SelfPrint( )} 
  
// -------------- method definition ----------------------------------
//
/* The go function for: self_print(self:Defmethod) [status=1] */
func (self *Defmethod) SelfPrint () EID { 
    var Result EID
    Result = Core.F_print_any(self.Arg.Selector.Id())
    if !ErrorIn(Result) {
    PRINC("(")
    if (self.Arg.Args.Id() != CNULL) { 
      Result = F_ppvariable_list(self.Arg.Args)
      } else {
      Result = EID{CFALSE.Id(),0}
      } 
    if !ErrorIn(Result) {
    PRINC(") : ")
    Result = F_printexp_any(self.SetArg,CFALSE)
    if !ErrorIn(Result) {
    Result = F_lbreak_void()
    if !ErrorIn(Result) {
    { 
      var va_arg1 *Core.PrettyPrinter
      var va_arg2 int
      va_arg1 = Core.C_pretty
      va_arg2 = (Core.C_pretty.Index+4)
      va_arg1.Index = va_arg2
      Result = EID{C__INT,IVAL(va_arg2)}
      } 
    }
    if !ErrorIn(Result) {
    PRINC(" ")
    F_princ_string(ToString(IfThenElse((self.Inline_ask == CTRUE),
      MakeString("=>").Id(),
      MakeString("->").Id())))
    PRINC(" ")
    Result = F_printexp_any(self.Body,CFALSE)
    if !ErrorIn(Result) {
    PRINC(" ")
    Result = EVOID
    }}}}}
    if !ErrorIn(Result) {
    { 
      var va_arg1 *Core.PrettyPrinter
      var va_arg2 int
      va_arg1 = Core.C_pretty
      va_arg2 = (Core.C_pretty.Index-4)
      va_arg1.Index = va_arg2
      Result = EID{C__INT,IVAL(va_arg2)}
      } 
    }
    return Result} 
  
// The EID go function for: self_print @ Defmethod (throw: true) 
func E_self_print_Defmethod_Language (self EID) EID { 
    return To_Defmethod(OBJ(self)).SelfPrint( )} 
  
// -------------- array definition -----------------------------------
/* The go function for: self_print(self:Defarray) [status=1] */
func (self *Defarray) SelfPrint () EID { 
    var Result EID
    Result = Core.F_CALL(C_print,ARGS(self.Arg.Args.At(0).ToEID()))
    if !ErrorIn(Result) {
    PRINC("[")
    { var arg_1 *ClaireList
      var try_2 EID
      try_2 = self.Arg.Args.Cdr()
      if ErrorIn(try_2) {Result = try_2
      } else {
      arg_1 = ToList(OBJ(try_2))
      Result = F_ppvariable_list(arg_1)
      }
      } 
    if !ErrorIn(Result) {
    PRINC("] : ")
    Result = Core.F_CALL(C_print,ARGS(self.SetArg.ToEID()))
    if !ErrorIn(Result) {
    Result = F_lbreak_void()
    if !ErrorIn(Result) {
    { 
      var va_arg1 *Core.PrettyPrinter
      var va_arg2 int
      va_arg1 = Core.C_pretty
      va_arg2 = (Core.C_pretty.Index+4)
      va_arg1.Index = va_arg2
      Result = EID{C__INT,IVAL(va_arg2)}
      } 
    }
    if !ErrorIn(Result) {
    PRINC(" := ")
    Result = F_printexp_any(self.Body,CFALSE)
    if !ErrorIn(Result) {
    PRINC(" ")
    Result = EVOID
    }}}}}
    if !ErrorIn(Result) {
    { 
      var va_arg1 *Core.PrettyPrinter
      var va_arg2 int
      va_arg1 = Core.C_pretty
      va_arg2 = (Core.C_pretty.Index-4)
      va_arg1.Index = va_arg2
      Result = EID{C__INT,IVAL(va_arg2)}
      } 
    }
    return Result} 
  
// The EID go function for: self_print @ Defarray (throw: true) 
func E_self_print_Defarray_Language (self EID) EID { 
    return To_Defarray(OBJ(self)).SelfPrint( )} 
  
// -------------- rule definition ------------------------------------
/* The go function for: self_print(self:Defrule) [status=1] */
func (self *Defrule) SelfPrint () EID { 
    var Result EID
    self.Ident.Princ()
    PRINC("(")
    Result = F_ppvariable_list(self.Args)
    if !ErrorIn(Result) {
    PRINC(") :: rule(")
    Result = F_lbreak_integer(4)
    if !ErrorIn(Result) {
    PRINC(" ")
    Result = Core.F_CALL(C_print,ARGS(self.Arg.ToEID()))
    if !ErrorIn(Result) {
    PRINC(" ")
    Result = F_lbreak_integer(4)
    if !ErrorIn(Result) {
    PRINC("=> ")
    Result = Core.F_CALL(C_print,ARGS(self.Body.ToEID()))
    if !ErrorIn(Result) {
    PRINC(")")
    Result = EVOID
    }}}}}
    if !ErrorIn(Result) {
    { 
      var va_arg1 *Core.PrettyPrinter
      var va_arg2 int
      va_arg1 = Core.C_pretty
      va_arg2 = (Core.C_pretty.Index-4)
      va_arg1.Index = va_arg2
      Result = EID{C__INT,IVAL(va_arg2)}
      } 
    }
    return Result} 
  
// The EID go function for: self_print @ Defrule (throw: true) 
func E_self_print_Defrule_Language (self EID) EID { 
    return To_Defrule(OBJ(self)).SelfPrint( )} 
  
/* The go function for: self_print(self:Defvar) [status=1] */
func (self *Defvar) SelfPrint () EID { 
    var Result EID
    Result = F_ppvariable_Variable(self.Ident)
    if !ErrorIn(Result) {
    PRINC(" := ")
    Result = F_printexp_any(self.Arg,CFALSE)
    }
    return Result} 
  
// The EID go function for: self_print @ Defvar (throw: true) 
func E_self_print_Defvar_Language (self EID) EID { 
    return To_Defvar(OBJ(self)).SelfPrint( )} 
  
// *********************************************************************
// *     Part 2: the general instantiation macro                       *
// *********************************************************************
// creation of a new object
//
/* The go function for: self_eval(self:Definition) [status=1] */
func (self *Definition) SelfEval () EID { 
    var Result EID
    { var _Zc *ClaireClass = self.Arg
      { var _Zo *ClaireObject
        var try_1 EID
        if (_Zc.Open <= 0) { 
          try_1 = ToException(Core.C_general_error.Make(MakeString("[105] cannot instantiate ~S").Id(),MakeConstantList(_Zc.Id()).Id())).Close()
          } else {
          try_1 = EID{CFALSE.Id(),0}
          } 
        if !ErrorIn(try_1) {
        try_1 = EID{F_new_object_class(_Zc).Id(),0}
        }
        if ErrorIn(try_1) {Result = try_1
        } else {
        _Zo = ToObject(OBJ(try_1))
        { var arg_2 *ClaireList
          var try_3 EID
          try_3 = F_Language_new_writes_object(_Zo,self.Args)
          if ErrorIn(try_3) {Result = try_3
          } else {
          arg_2 = ToList(OBJ(try_3))
          Result = Core.F_Core_new_defaults_object(_Zo,arg_2)
          }
          } 
        }
        } 
      } 
    return Result} 
  
// The EID go function for: self_eval @ Definition (throw: true) 
func E_self_eval_Definition (self EID) EID { 
    return To_Definition(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Definition 
func EVAL_Definition (x *ClaireAny) EID { 
     return To_Definition(x).SelfEval()} 
  
// fast definition : no inverse management + no "close" method
/* The go function for: fast_definition?(c:class) [status=0] */
func F_Language_fast_definition_ask_class (c *ClaireClass) *ClaireBoolean { 
    var Result *ClaireBoolean
    { 
      var v_and2 *ClaireBoolean
      
      v_and2 = Core.F__sup_integer(c.Open,1)
      if (v_and2 == CFALSE) {Result = CFALSE
      } else { 
        v_and2 = Equal(C_method.Id(),Core.F_owner_any(Core.F__at_property1(C_close,c).Id()).Id()).Not
        if (v_and2 == CFALSE) {Result = CFALSE
        } else { 
          { var arg_1 *ClaireAny
            { 
              var s *ClaireSlot
              _ = s
              var s_iter *ClaireAny
              arg_1= CFALSE.Id()
              for _,s_iter = range(c.Slots.ValuesO()){ 
                s = ToSlot(s_iter)
                if ((s.Selector.Inverse.Id() != CNULL) || 
                    ((s.Selector.Store_ask == CTRUE) || 
                      (s.Selector.IfWrite != CNULL))) { 
                  arg_1 = CTRUE.Id()
                  break
                  } 
                } 
              } 
            v_and2 = Core.F_not_any(arg_1)
            } 
          if (v_and2 == CFALSE) {Result = CFALSE
          } else { 
            Result = CTRUE} 
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: fast_definition? @ class (throw: false) 
func E_Language_fast_definition_ask_class (c EID) EID { 
    return EID{F_Language_fast_definition_ask_class(ToClass(OBJ(c)) ).Id(),0}} 
  
// then the evaluation is simpler ! write_fast does the range checking (may return an error) 
/* The go function for: self_eval(self:DefFast) [status=1] */
func (self *Language_DefFast) SelfEval () EID { 
    var Result EID
    { var _Zo *ClaireObject = F_new_object_class(self.Arg)
      { 
        var x *Call
        _ = x
        var x_iter *ClaireAny
        Result= EID{CFALSE.Id(),0}
        var x_support *ClaireList
        x_support = self.Args
        x_len := x_support.Length()
        for i_it := 0; i_it < x_len; i_it++ { 
          x_iter = x_support.At(i_it)
          x = To_Call(x_iter)
          var loop_1 EID
          _ = loop_1
          { var p *ClaireProperty
            var try_2 EID
            try_2 = F_make_a_property_any(x.Args.At(0))
            if ErrorIn(try_2) {loop_1 = try_2
            } else {
            p = ToProperty(OBJ(try_2))
            { 
              var arg_3 EID
              arg_3 = EVAL(x.Args.At(1))
              if ErrorIn(arg_3) {loop_1 = arg_3
              } else {
              loop_1 = p.WriteEID(_Zo,arg_3)}
              } 
            }
            } 
          if ErrorIn(loop_1) {Result = loop_1
          break
          } else {
          }
          } 
        } 
      if !ErrorIn(Result) {
      Result = EID{_Zo.Id(),0}
      }
      } 
    return Result} 
  
// The EID go function for: self_eval @ DefFast (throw: true) 
func E_self_eval_DefFast (self EID) EID { 
    return To_Language_DefFast(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: DefFast 
func EVAL_Language_DefFast (x *ClaireAny) EID { 
     return To_Language_DefFast(x).SelfEval()} 
  
// for a fast_definition, simpler eval
// the instantiation body is a sequence of words from which the initialization
// of the object must be built.
// CLAIRE4 : renamed complete(self:object,%l:list) to new_writes()
/* The go function for: new_writes(self:object,%l:list) [status=1] */
func F_Language_new_writes_object (self *ClaireObject,_Zl *ClaireList) EID { 
    var Result EID
    { var lp *ClaireList = ToType(CEMPTY.Id()).EmptyList()
      { 
        var x *Call
        _ = x
        var x_iter *ClaireAny
        Result= EID{CFALSE.Id(),0}
        var x_support *ClaireList
        x_support = _Zl
        x_len := x_support.Length()
        for i_it := 0; i_it < x_len; i_it++ { 
          x_iter = x_support.At(i_it)
          x = To_Call(x_iter)
          var loop_1 EID
          _ = loop_1
          { var p *ClaireProperty
            var try_2 EID
            try_2 = F_make_a_property_any(x.Args.At(0))
            if ErrorIn(try_2) {loop_1 = try_2
            } else {
            p = ToProperty(OBJ(try_2))
            { var y *ClaireAny
              var try_3 EID
              try_3 = EVAL(x.Args.At(1))
              if ErrorIn(try_3) {loop_1 = try_3
              } else {
              y = ANY(try_3)
              { var s *ClaireObject = Core.F__at_property1(p,self.Isa)
                if (C_slot.Id() == s.Isa.Id()) { 
                  { var g0219 *ClaireSlot = ToSlot(s.Id())
                    if (y == CNULL) { 
                      lp = lp.AddFast(p.Id())
                      } 
                    if (g0219.Range.Contains(y) != CTRUE) { 
                      loop_1 = Core.F_range_is_wrong_slot(g0219,y)
                      } else {
                      loop_1 = Core.F_update_property(p,
                        self,
                        g0219.Index,
                        g0219.Srange,
                        y)
                      } 
                    if ErrorIn(loop_1) {Result = loop_1
                    break
                    } else {
                    }
                    } 
                  } else {
                  loop_1 = ToException(Core.C_general_error.Make(MakeString("[106] the object ~S does not understand ~S").Id(),MakeConstantList(self.Id(),p.Id()).Id())).Close()
                  } 
                } 
              }
              } 
            }
            } 
          if ErrorIn(loop_1) {Result = loop_1
          break
          } else {
          }
          } 
        } 
      if !ErrorIn(Result) {
      Result = EID{lp.Id(),0}
      }
      } 
    return Result} 
  
// The EID go function for: new_writes @ object (throw: true) 
func E_Language_new_writes_object (self EID,_Zl EID) EID { 
    return F_Language_new_writes_object(ToObject(OBJ(self)),ToList(OBJ(_Zl)) )} 
  
// creation of a new named object
/* The go function for: self_eval(self:Defobj) [status=1] */
func (self *Defobj) SelfEval () EID { 
    var Result EID
    { var _Zc *ClaireClass = self.Arg
      { var _Zo *ClaireObject = ToObject(CNULL)
        if (_Zc.Open <= 0) { 
          Result = ToException(Core.C_general_error.Make(MakeString("[105] cannot instantiate ~S").Id(),MakeConstantList(_Zc.Id()).Id())).Close()
          } else {
          Result = EID{CFALSE.Id(),0}
          } 
        if !ErrorIn(Result) {
        if (_Zc.IsIn(C_thing) == CTRUE) { 
          var try_1 EID
          try_1 = F_new_thing_class(_Zc,self.Ident)
          if ErrorIn(try_1) {Result = try_1
          } else {
          _Zo = ToObject(OBJ(try_1))
          Result = EID{_Zo.Id(),0}
          if (_Zo.Isa.IsIn(C_property) == CTRUE) { 
            { var g0221 *ClaireProperty = ToProperty(_Zo.Id())
              if (g0221.Restrictions.Length() > 0) { 
                Result = ToException(Core.C_general_error.Make(MakeString("[188] the property ~S is already defined").Id(),MakeConstantList(g0221.Id()).Id())).Close()
                } else {
                Result = EID{CFALSE.Id(),0}
                } 
              } 
            } else {
            Result = EID{CFALSE.Id(),0}
            } 
          }
          } else {
          _Zo = F_new_object_class(_Zc)
          if (_Zc.Open == ClEnv.Open) { 
            _Zc.Instances.AddFast(_Zo.Id())
            } 
          Result = self.Ident.Put(_Zo.Id()).ToEID()
          } 
        if !ErrorIn(Result) {
        { var arg_2 *ClaireList
          var try_3 EID
          try_3 = F_Language_new_writes_object(_Zo,self.Args)
          if ErrorIn(try_3) {Result = try_3
          } else {
          arg_2 = ToList(OBJ(try_3))
          Result = Core.F_Core_new_defaults_object(_Zo,arg_2)
          }
          } 
        }}
        } 
      } 
    return Result} 
  
// The EID go function for: self_eval @ Defobj (throw: true) 
func E_self_eval_Defobj (self EID) EID { 
    return To_Defobj(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Defobj 
func EVAL_Defobj (x *ClaireAny) EID { 
     return To_Defobj(x).SelfEval()} 
  
// creation of a new named object
// note that final() is the marker of a forward definition in CLAIRE4
/* The go function for: self_eval(self:Defclass) [status=1] */
func (self *Defclass) SelfEval () EID { 
    var Result EID
    if ((C_class.Id() == Core.F_owner_any(self.Ident.Value()).Id()) && 
        ((ToClass(self.Ident.Value()).Open != ClEnv.Final) || 
            (self.Arg.Id() != ToClass(self.Ident.Value()).Superclass.Id()))) { 
      Result = ToException(Core.C_general_error.Make(MakeString("[107] class re-definition is not valid: ~S").Id(),MakeConstantList(self.Id()).Id())).Close()
      } else {
      { var _Zo *ClaireClass
        var try_1 EID
        try_1 = self.Ident.Class_I(self.Arg)
        if ErrorIn(try_1) {Result = try_1
        } else {
        _Zo = ToClass(OBJ(try_1))
        { 
          var x *ClaireAny
          _ = x
          Result= EID{CFALSE.Id(),0}
          var x_support *ClaireList
          x_support = self.Args
          x_len := x_support.Length()
          for i_it := 0; i_it < x_len; i_it++ { 
            x = x_support.At(i_it)
            var loop_2 EID
            _ = loop_2
            { var v *ClaireAny = CNULL
              if (x.Isa.IsIn(C_Call) == CTRUE) { 
                { var g0222 *Call = To_Call(x)
                  var try_3 EID
                  try_3 = EVAL(g0222.Args.At(1))
                  if ErrorIn(try_3) {loop_2 = try_3
                  Result = try_3
                  break
                  } else {
                  v = ANY(try_3)
                  loop_2 = v.ToEID()
                  g0222 = To_Call(g0222.Args.At(0))
                  loop_2 = EID{g0222.Id(),0}
                  }
                  if ErrorIn(loop_2) {Result = loop_2
                  break
                  } else {
                  x = g0222.Id()
                  loop_2 = x.ToEID()
                  }
                  } 
                } else {
                loop_2 = EID{CFALSE.Id(),0}
                } 
              if ErrorIn(loop_2) {Result = loop_2
              break
              } else {
              { var rt *ClaireTypeExpression
                var try_4 EID
                try_4 = F_extract_type_any(ANY(Core.F_CALL(C_range,ARGS(x.ToEID()))))
                if ErrorIn(try_4) {loop_2 = try_4
                } else {
                rt = ToTypeExpression(OBJ(try_4))
                { var p *ClaireProperty
                  var try_5 EID
                  try_5 = F_make_a_property_any(ANY(Core.F_CALL(C_mClaire_pname,ARGS(x.ToEID()))))
                  if ErrorIn(try_5) {loop_2 = try_5
                  } else {
                  p = ToProperty(OBJ(try_5))
                  var g0223I *ClaireBoolean
                  var try_6 EID
                  { 
                    var v_and9 *ClaireBoolean
                    
                    v_and9 = Core.F_known_ask_any(v)
                    if (v_and9 == CFALSE) {try_6 = EID{CFALSE.Id(),0}
                    } else { 
                      var try_7 EID
                      { var arg_8 *ClaireBoolean
                        var try_9 EID
                        try_9 = Core.F_BELONG(v,rt.Id())
                        if ErrorIn(try_9) {try_7 = try_9
                        } else {
                        arg_8 = ToBoolean(OBJ(try_9))
                        try_7 = EID{arg_8.Not.Id(),0}
                        }
                        } 
                      if ErrorIn(try_7) {try_6 = try_7
                      } else {
                      v_and9 = ToBoolean(OBJ(try_7))
                      if (v_and9 == CFALSE) {try_6 = EID{CFALSE.Id(),0}
                      } else { 
                        try_6 = EID{CTRUE.Id(),0}} 
                      } 
                    }
                    } 
                  if ErrorIn(try_6) {loop_2 = try_6
                  } else {
                  g0223I = ToBoolean(OBJ(try_6))
                  if (g0223I == CTRUE) { 
                    loop_2 = ToException(Core.C_general_error.Make(MakeString("[108] default(~S) = ~S does not belong to ~S").Id(),MakeConstantList(x,v,rt.Id()).Id())).Close()
                    } else {
                    loop_2 = EID{CFALSE.Id(),0}
                    } 
                  }
                  if ErrorIn(loop_2) {Result = loop_2
                  break
                  } else {
                  { var s2test *ClaireAny
                    { var sx_some *ClaireAny = CNULL
                      { 
                        var sx *ClaireSlot
                        _ = sx
                        var sx_iter *ClaireAny
                        for _,sx_iter = range(self.Arg.Slots.ValuesO()){ 
                          sx = ToSlot(sx_iter)
                          if (sx.Selector.Id() == p.Id()) { 
                            sx_some = sx.Id()
                            break
                            } 
                          } 
                        } 
                      s2test = sx_some
                      } 
                    if (s2test != CNULL) { 
                      { var s2 *ClaireSlot = ToSlot(s2test)
                        if (p.Open <= 0) { 
                          loop_2 = ToException(Core.C_general_error.Make(MakeString("[181] cannot overide a slot for a closed property ~S").Id(),MakeConstantList(p.Id()).Id())).Close()
                          }  else if (ToType(rt.Id()).Included(s2.Range) != CTRUE) { 
                          loop_2 = ToException(Core.C_general_error.Make(MakeString("[XXX] slot redefinition of ~S must be covariant, ~S is not a subtype").Id(),MakeConstantList(s2.Id(),rt.Id()).Id())).Close()
                          } else {
                          loop_2 = EID{CFALSE.Id(),0}
                          } 
                        } 
                      } else {
                      loop_2 = EID{CNULL,0}
                      } 
                    } 
                  if ErrorIn(loop_2) {Result = loop_2
                  break
                  } else {
                  loop_2 = EID{Core.F_close_slot(_Zo.AddSlot(p,ToType(rt.Id()),F_Language_getDefault_type(ToType(rt.Id()),v))).Id(),0}
                  }}
                  }
                  } 
                }
                } 
              if ErrorIn(loop_2) {Result = loop_2
              break
              } else {
              }}
              } 
            if ErrorIn(loop_2) {Result = loop_2
            break
            } else {
            }
            } 
          } 
        if !ErrorIn(Result) {
        F_close_class(_Zo)
        if (self.Forward_ask == CTRUE) { 
          _Zo.Open = ClEnv.Final
          }  else if (_Zo.Open == ClEnv.Final) { 
          _Zo.Open = self.Arg.Open
          } 
        if (ToType(_Zo.Id()).Included(ToType(C_primitive.Id())) == CTRUE) { 
          _Zo.Open = -1
          } 
        _Zo.Params = self.Params
        { 
          var p *ClaireAny
          _ = p
          var p_support *ClaireList
          p_support = self.Params
          p_len := p_support.Length()
          for i_it := 0; i_it < p_len; i_it++ { 
            p = p_support.At(i_it)
            ToRelation(p).Open = 0
            } 
          } 
        Result = F_attach_comment_any(_Zo.Id())
        if !ErrorIn(Result) {
        Result = EID{_Zo.Id(),0}
        }}
        }
        } 
      } 
    return Result} 
  
// The EID go function for: self_eval @ Defclass (throw: true) 
func E_self_eval_Defclass (self EID) EID { 
    return To_Defclass(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Defclass 
func EVAL_Defclass (x *ClaireAny) EID { 
     return To_Defclass(x).SelfEval()} 
  
// we compute the proper default value (reused by compiler) - for int, float, sets and lists
/* The go function for: getDefault(rt:type,v:any) [status=0] */
func F_Language_getDefault_type (rt *ClaireType,v *ClaireAny) *ClaireAny { 
    if (v == CNULL) { 
      if (rt.Included(ToType(C_integer.Id())) == CTRUE) { 
        return  MakeInteger(0).Id()
        }  else if (rt.Included(ToType(C_float.Id())) == CTRUE) { 
        return  MakeFloat(0).Id()
        }  else if (rt.Included(ToType(C_set.Id())) == CTRUE) { 
        return  Core.F_of_extract_type(rt).EmptySet().Id()
        }  else if (rt.Included(ToType(C_list.Id())) == CTRUE) { 
        return  Core.F_of_extract_type(rt).EmptyList().Id()
        } else {
        return  CNULL
        } 
      } else {
      return  v
      } 
    } 
  
// The EID go function for: getDefault @ type (throw: false) 
func E_Language_getDefault_type (rt EID,v EID) EID { 
    return F_Language_getDefault_type(ToType(OBJ(rt)),ANY(v) ).ToEID()} 
  
// method definition
// v0.01
/* The go function for: self_eval(self:Defmethod) [status=1] */
func (self *Defmethod) SelfEval () EID { 
    var Result EID
    if (self.Arg.Isa.IsIn(C_Call) != CTRUE) { 
      Result = ToException(Core.C_general_error.Make(MakeString("[110] wrong signature definition ~S").Id(),MakeConstantList(self.Arg.Id()).Id())).Close()
      } else {
      Result = EID{CFALSE.Id(),0}
      } 
    if !ErrorIn(Result) {
    { var p *ClaireProperty
      var try_1 EID
      try_1 = F_make_a_property_any(self.Arg.Selector.Id())
      if ErrorIn(try_1) {Result = try_1
      } else {
      p = ToProperty(OBJ(try_1))
      { var l *ClaireList = self.Arg.Args
        { var lv *ClaireList
          if ((l.Length() == 1) && 
              (l.At(0) == ClEnv.Id())) { 
            { 
              var v_bag_arg *ClaireAny
              lv= ToType(CEMPTY.Id()).EmptyList()
              { var _CL_obj *ClaireVariable = To_Variable(new(ClaireVariable).Is(C_Variable))
                _CL_obj.Pname = Core.F_symbol_I_string2(MakeString("XfakeParameter"))
                _CL_obj.Range = ToType(C_void.Id())
                v_bag_arg = _CL_obj.Id()
                } 
              lv.AddFast(v_bag_arg)} 
            } else {
            lv = l
            } 
          { var lp *ClaireList
            var try_2 EID
            try_2 = F_extract_signature_list(lv)
            if ErrorIn(try_2) {Result = try_2
            } else {
            lp = ToList(OBJ(try_2))
            { var lrange *ClaireList
              var try_3 EID
              try_3 = F_extract_range_any(self.SetArg,lv,ToList(C_LDEF.Value))
              if ErrorIn(try_3) {Result = try_3
              } else {
              lrange = ToList(OBJ(try_3))
              { var lbody *ClaireList
                var try_4 EID
                try_4 = F_extract_status_any(self.Body)
                if ErrorIn(try_4) {Result = try_4
                } else {
                lbody = ToList(OBJ(try_4))
                { var m *ClaireMethod = F_add_method_property(p,lp,ToType(lrange.At(0)),ToInteger(lbody.At(0)).Value,ToFunction(lbody.At(1)))
                  if ((p.Open > 0) && 
                      (p.Open <= 1)) { 
                    { var r *ClaireAny
                      var try_5 EID
                      { var r_some *ClaireAny = CNULL
                        { 
                          var r *ClaireAny
                          _ = r
                          try_5= EID{CFALSE.Id(),0}
                          for _,r = range(p.Restrictions.ValuesO()){ 
                            var loop_6 EID
                            _ = loop_6
                            if (r != m.Id()) { 
                              var g0224I *ClaireBoolean
                              var try_7 EID
                              { var arg_8 *ClaireAny
                                var try_9 EID
                                try_9 = Core.F_CALL(ToProperty(C__exp.Id()),ARGS(Core.F_CALL(C_domain,ARGS(r.ToEID())),EID{m.Domain.Id(),0}))
                                if ErrorIn(try_9) {try_7 = try_9
                                } else {
                                arg_8 = ANY(try_9)
                                try_7 = EID{F_boolean_I_any(arg_8).Id(),0}
                                }
                                } 
                              if ErrorIn(try_7) {loop_6 = try_7
                              } else {
                              g0224I = ToBoolean(OBJ(try_7))
                              if (g0224I == CTRUE) { 
                                r_some = r
                                try_5 = r_some.ToEID()
                                break
                                } else {
                                loop_6 = EID{CFALSE.Id(),0}
                                } 
                              }
                              } else {
                              loop_6 = EID{CFALSE.Id(),0}
                              } 
                            if ErrorIn(loop_6) {try_5 = loop_6
                            break
                            } else {
                            }
                            } 
                          } 
                        if !ErrorIn(try_5) {
                        try_5 = r_some.ToEID()
                        }
                        } 
                      if ErrorIn(try_5) {Result = try_5
                      } else {
                      r = ANY(try_5)
                      if (r != CNULL) { 
                        Result = Core.F_tformat_string(MakeString("--- WARNING ! [186] conflict between ~S and ~S is dangerous since ~S is closed\n"),1,MakeConstantList(m.Id(),r,p.Id()))
                        } else {
                        Result = EID{CNULL,0}
                        } 
                      }
                      } 
                    } else {
                    Result = EID{CFALSE.Id(),0}
                    } 
                  if !ErrorIn(Result) {
                  C_LDEF.Value = ToType(C_any.Id()).EmptyList().Id()
                  if (lbody.At(2) != C_body.Id()) { 
                    if (ClEnv.Jito_ask == CTRUE) { 
                      Core.F_tformat_string(MakeString("---- jito for ~S\n"),3,MakeConstantList(m.Id()))
                      } 
                    { 
                      var va_arg1 *ClaireMethod
                      var va_arg2 *ClaireLambda
                      va_arg1 = m
                      var try_10 EID
                      { var arg_11 *ClaireLambda
                        var try_12 EID
                        try_12 = F_lambda_I_list(lv,lbody.At(2))
                        if ErrorIn(try_12) {try_10 = try_12
                        } else {
                        arg_11 = ToLambda(OBJ(try_12))
                        try_10 = F_Language_jito_any(arg_11.Id())
                        }
                        } 
                      if ErrorIn(try_10) {Result = try_10
                      } else {
                      va_arg2 = ToLambda(OBJ(try_10))
                      va_arg1.Formula = va_arg2
                      Result = EID{va_arg2.Id(),0}
                      }
                      } 
                    } else {
                    Result = EID{CFALSE.Id(),0}
                    } 
                  if !ErrorIn(Result) {
                  if (lrange.Length() > 1) { 
                    m.Typing = lrange.At(1)
                    } 
                  m.Inline_ask = self.Inline_ask
                  Result = F_attach_comment_any(m.Id())
                  if !ErrorIn(Result) {
                  Core.F_close_method(m)
                  if ((p.Id() == C_close.Id()) && 
                      (m.Range.Included(ToType(Core.F_domain_I_restriction(ToRestriction(m.Id())).Id())) != CTRUE)) { 
                    Result = ToException(Core.C_general_error.Make(MakeString("[184] the close method ~S has a wrong range").Id(),MakeConstantList(m.Id()).Id())).Close()
                    } else {
                    Result = EID{CFALSE.Id(),0}
                    } 
                  if !ErrorIn(Result) {
                  Result = EID{m.Id(),0}
                  }}}}
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
  
// The EID go function for: self_eval @ Defmethod (throw: true) 
func E_self_eval_Defmethod (self EID) EID { 
    return To_Defmethod(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Defmethod 
func EVAL_Defmethod (x *ClaireAny) EID { 
     return To_Defmethod(x).SelfEval()} 
  
// v3.2.24 : -1 : final
// attach a cute comment if needed ... to a defclass or a defmethod
/* The go function for: attach_comment(x:any) [status=1] */
func F_attach_comment_any (x *ClaireAny) EID { 
    var Result EID
    if ((ToBoolean(C_NeedComment.Value) == CTRUE) && 
        (C_iClaire_LastComment.Value != CNULL)) { 
      Result = Core.F_write_property(C_comment,ToObject(x),C_iClaire_LastComment.Value)
      } else {
      Result = EID{CFALSE.Id(),0}
      } 
    return Result} 
  
// The EID go function for: attach_comment @ any (throw: true) 
func E_attach_comment_any (x EID) EID { 
    return F_attach_comment_any(ANY(x) )} 
  
// returns the list of types AND modifies LDEF
/* The go function for: iClaire/extract_signature(l:list) [status=1] */
func F_extract_signature_list (l *ClaireList) EID { 
    var Result EID
    C_LDEF.Value = ToType(C_any.Id()).EmptyList().Id()
    { var n int = 0
      { 
        var v_list3 *ClaireList
        var v *ClaireVariable
        var v_local3 *ClaireAny
        v_list3 = l
        Result = EID{CreateList(ToType(CEMPTY.Id()),v_list3.Length()).Id(),0}
        for CLcount := 0; CLcount < v_list3.Length(); CLcount++{ 
          v = To_Variable(v_list3.At(CLcount))
          var try_1 EID
          if (v.Isa.IsIn(C_Variable) != CTRUE) { 
            try_1 = ToException(Core.C_general_error.Make(MakeString("[111] wrong typed argument ~S").Id(),MakeConstantList(v.Id()).Id())).Close()
            } else {
            { var p *ClaireAny
              var try_2 EID
              try_2 = F_extract_pattern_any(v.Range.Id(),MakeConstantList(MakeInteger(n).Id()))
              if ErrorIn(try_2) {try_1 = try_2
              } else {
              p = ANY(try_2)
              n = (n+1)
              if (p == CNULL) { 
                try_1 = ToException(Core.C_general_error.Make(MakeString("[111] wrong typed argument ~S (~S)").Id(),MakeConstantList(v.Id(),v.Range.Id()).Id())).Close()
                } else {
                try_1 = EID{CFALSE.Id(),0}
                } 
              if ErrorIn(try_1) {Result = try_1
              break
              } else {
              v.Range = F_type_I_any(p)
              try_1 = p.ToEID()
              }
              }
              } 
            } 
          if ErrorIn(try_1) {Result = try_1
          break
          } else {
          v_local3 = ANY(try_1)
          ToList(OBJ(Result)).PutAt(CLcount,v_local3)
          } 
        }
        } 
      } 
    return Result} 
  
// The EID go function for: iClaire/extract_signature @ list (throw: true) 
func E_extract_signature_list (l EID) EID { 
    return F_extract_signature_list(ToList(OBJ(l)) )} 
  
// takes an <exp> that must belong to <type> and returns the CLAIRE type
// if LDEF is non-empty, it is used as a list of type variable and patterns
// may be returned. In addition, if the path list is non empty, new type
// variables may be defined. a syntax error will produce the unknown value
//
/* The go function for: iClaire/extract_pattern(x:any,path:list) [status=1] */
func F_extract_pattern_any (x *ClaireAny,path *ClaireList) EID { 
    var Result EID
    if (C_class.Id() == x.Isa.Id()) { 
      { var g0225 *ClaireClass = ToClass(x)
        Result = EID{g0225.Id(),0}
        } 
      }  else if (C_set.Id() == x.Isa.Id()) { 
      { var g0226 *ClaireSet = ToSet(x)
        { var z *ClaireAny
          var try_1 EID
          if (g0226.Size() == 1) { 
            { var arg_2 *ClaireAny
              var try_3 EID
              try_3 = Core.F_the_type(ToType(g0226.Id()))
              if ErrorIn(try_3) {try_1 = try_3
              } else {
              arg_2 = ANY(try_3)
              try_1 = F_extract_pattern_any(arg_2,CNIL)
              }
              } 
            } else {
            try_1 = EID{CFALSE.Id(),0}
            } 
          if ErrorIn(try_1) {Result = try_1
          } else {
          z = ANY(try_1)
          if (z.Isa.IsIn(C_Reference) == CTRUE) { 
            { var g0227 *ClaireReference = To_Reference(z)
              { var w *ClaireReference = To_Reference(g0227.Copy().Id())
                w.Arg = CTRUE
                Result = EID{w.Id(),0}
                } 
              } 
            } else {
            Result = EID{g0226.Id(),0}
            } 
          }
          } 
        } 
      }  else if (x.Isa.IsIn(C_Tuple) == CTRUE) { 
      { var g0229 *Tuple = To_Tuple(x)
        { var ltp *ClaireList
          var try_4 EID
          { 
            var v_list5 *ClaireList
            var z *ClaireAny
            var v_local5 *ClaireAny
            v_list5 = g0229.Args
            try_4 = EID{CreateList(ToType(CEMPTY.Id()),v_list5.Length()).Id(),0}
            for CLcount := 0; CLcount < v_list5.Length(); CLcount++{ 
              z = v_list5.At(CLcount)
              var try_5 EID
              try_5 = F_extract_pattern_any(z,path)
              if ErrorIn(try_5) {try_4 = try_5
              break
              } else {
              v_local5 = ANY(try_5)
              ToList(OBJ(try_4)).PutAt(CLcount,v_local5)
              } 
            }
            } 
          if ErrorIn(try_4) {Result = try_4
          } else {
          ltp = ToList(OBJ(try_4))
          var g0236I *ClaireBoolean
          { var arg_6 *ClaireAny
            { 
              var y *ClaireAny
              _ = y
              arg_6= CFALSE.Id()
              var y_support *ClaireList
              y_support = ltp
              y_len := y_support.Length()
              for i_it := 0; i_it < y_len; i_it++ { 
                y = y_support.At(i_it)
                if (y == CNULL) { 
                  arg_6 = CTRUE.Id()
                  break
                  } 
                } 
              } 
            g0236I = F_boolean_I_any(arg_6)
            } 
          if (g0236I == CTRUE) { 
            Result = EID{CNULL,0}
            } else {
            Result = EID{ltp.Tuple_I().Id(),0}
            } 
          }
          } 
        } 
      }  else if (x.Isa.IsIn(Core.C_global_variable) == CTRUE) { 
      { var g0230 *Core.GlobalVariable = Core.ToGlobalVariable(x)
        Result = F_extract_pattern_any(g0230.Value,path)
        } 
      }  else if (x.Isa.IsIn(C_Call) == CTRUE) { 
      { var g0231 *Call = To_Call(x)
        { var p *ClaireProperty = g0231.Selector
          if (p.Id() == Core.C_U.Id()) { 
            { var x1 *ClaireAny
              var try_7 EID
              try_7 = F_extract_pattern_any(g0231.Args.At(0),CNIL)
              if ErrorIn(try_7) {Result = try_7
              } else {
              x1 = ANY(try_7)
              { var x2 *ClaireAny
                var try_8 EID
                try_8 = F_extract_pattern_any(g0231.Args.At(1),CNIL)
                if ErrorIn(try_8) {Result = try_8
                } else {
                x2 = ANY(try_8)
                if ((x1 == CNULL) || 
                    (x2 == CNULL)) { 
                  Result = EID{CNULL,0}
                  } else {
                  Result = EID{Core.F_U_type(ToType(x1),ToType(x2)).Id(),0}
                  } 
                }
                } 
              }
              } 
            }  else if (p.Id() == C__exp.Id()) { 
            { var arg_9 *ClaireAny
              var try_11 EID
              try_11 = F_extract_pattern_any(g0231.Args.At(0),CNIL)
              if ErrorIn(try_11) {Result = try_11
              } else {
              arg_9 = ANY(try_11)
              { var arg_10 *ClaireAny
                var try_12 EID
                try_12 = F_extract_pattern_any(g0231.Args.At(1),CNIL)
                if ErrorIn(try_12) {Result = try_12
                } else {
                arg_10 = ANY(try_12)
                Result = Core.F_CALL(ToProperty(C__exp.Id()),ARGS(arg_9.ToEID(),arg_10.ToEID()))
                }
                } 
              }
              } 
            }  else if (p.Id() == C__dot_dot.Id()) { 
            { var v1 *ClaireAny = F_extract_item_any(g0231.Args.At(0),CNIL.Id())
              { var v2 *ClaireAny = F_extract_item_any(g0231.Args.At(1),CNIL.Id())
                if ((C_integer.Id() == v1.Isa.Id()) && 
                    (C_integer.Id() == v2.Isa.Id())) { 
                  Result = EID{Core.F__dot_dot_integer(ToInteger(v1).Value,ToInteger(v2).Value).Id(),0}
                  } else {
                  Result = EID{CNULL,0}
                  } 
                } 
              } 
            }  else if (p.Id() == C_nth.Id()) { 
            Result = F_extract_pattern_nth_list(g0231.Args,path)
            }  else if (p.Id() == C__star.Id()) { 
            { var z *ClaireAny
              var try_13 EID
              try_13 = F_extract_pattern_any(g0231.Args.At(0),path)
              if ErrorIn(try_13) {Result = try_13
              } else {
              z = ANY(try_13)
              if (z != CNULL) { 
                Result = EID{Core.F_U_type(ToType(z),ToType(MakeConstantSet(CNULL).Id())).Id(),0}
                } else {
                Result = EID{CNULL,0}
                } 
              }
              } 
            } else {
            Result = EID{CNULL,0}
            } 
          } 
        } 
      }  else if (x.Isa.IsIn(C_type) == CTRUE) { 
      { var g0232 *ClaireType = ToType(x)
        Result = EID{g0232.Id(),0}
        } 
      }  else if (x.Isa.IsIn(C_unbound_symbol) == CTRUE) { 
      { var g0233 *ClaireUnboundSymbol = ToUnboundSymbol(x)
        { var s *ClaireSymbol
          var try_14 EID
          try_14 = F_extract_symbol_any(g0233.Id())
          if ErrorIn(try_14) {Result = try_14
          } else {
          s = ToSymbol(OBJ(try_14))
          { var v *ClaireAny
            var try_15 EID
            { var z_some *ClaireAny = CNULL
              { 
                var z *ClaireAny
                _ = z
                try_15= EID{CFALSE.Id(),0}
                var z_support *ClaireList
                var try_16 EID
                try_16 = Core.F_enumerate_any(C_LDEF.Value)
                if ErrorIn(try_16) {try_15 = try_16
                } else {
                z_support = ToList(OBJ(try_16))
                z_len := z_support.Length()
                for i_it := 0; i_it < z_len; i_it++ { 
                  z = z_support.At(i_it)
                  if (ANY(Core.F_CALL(C_mClaire_pname,ARGS(z.ToEID()))) == s.Id()) { 
                    z_some = z
                    try_15 = z_some.ToEID()
                    break
                    } 
                  }
                  } 
                } 
              if !ErrorIn(try_15) {
              try_15 = z_some.ToEID()
              }
              } 
            if ErrorIn(try_15) {Result = try_15
            } else {
            v = ANY(try_15)
            if (v != CNULL) { 
              Result = Core.F_CALL(C_range,ARGS(v.ToEID()))
              } else {
              var g0237I *ClaireBoolean
              if (path.Isa.IsIn(C_list) == CTRUE) { 
                g0237I = Core.F__sup_integer(path.Length(),1)
                } else {
                g0237I = CFALSE
                } 
              if (g0237I == CTRUE) { 
                { var y *ClaireReference
                  var try_17 EID
                  { var arg_18 *ClaireList
                    var try_19 EID
                    try_19 = path.Cdr()
                    if ErrorIn(try_19) {try_17 = try_19
                    } else {
                    arg_18 = ToList(OBJ(try_19))
                    try_17 = EID{Core.F_Reference_I_list(arg_18,ToInteger(path.At(0)).Value).Id(),0}
                    }
                    } 
                  if ErrorIn(try_17) {Result = try_17
                  } else {
                  y = To_Reference(OBJ(try_17))
                  { var v *ClaireVariable
                    { var _CL_obj *ClaireVariable = To_Variable(new(ClaireVariable).Is(C_Variable))
                      _CL_obj.Pname = s
                      _CL_obj.Range = ToType(y.Id())
                      v = _CL_obj
                      } 
                    
                    var v_gassign20 *ClaireAny
                    var try_21 EID
                    try_21 = Core.F_CALL(ToProperty(C_add.Id()),ARGS(EID{C_LDEF.Value,0},EID{v.Id(),0}))
                    if ErrorIn(try_21) {Result = try_21
                    } else {
                    v_gassign20 = ANY(try_21)
                    C_LDEF.Value = v_gassign20
                    Result = v_gassign20.ToEID()
                    } 
                  if !ErrorIn(Result) {
                  Result = EID{C_void.Id(),0}
                  }
                  } 
                }
                } 
              } else {
              Result = EID{CNULL,0}
              } 
            } 
          }
          } 
        }
        } 
      } 
    } else {
    Result = EID{CNULL,0}
    } 
  return Result} 

// The EID go function for: iClaire/extract_pattern @ any (throw: true) 
func E_extract_pattern_any (x EID,path EID) EID { 
  return F_extract_pattern_any(ANY(x),ToList(OBJ(path)) )} 

// takes an <exp> that must belong to <type> and returns the CLAIRE type
/* The go function for: iClaire/extract_type(x:any) [status=1] */
func F_extract_type_any (x *ClaireAny) EID { 
  var Result EID
  C_LDEF.Value = ToType(C_any.Id()).EmptyList().Id()
  { var r *ClaireAny
    var try_1 EID
    try_1 = F_extract_pattern_any(x,CNIL)
    if ErrorIn(try_1) {Result = try_1
    } else {
    r = ANY(try_1)
    if (r == CNULL) { 
      Result = ToException(Core.C_general_error.Make(MakeString("[112] wrong type expression ~S").Id(),MakeConstantList(x).Id())).Close()
      } else {
      Result = r.ToEID()
      } 
    }
    } 
  return Result} 

// The EID go function for: iClaire/extract_type @ any (throw: true) 
func E_extract_type_any (x EID) EID { 
  return F_extract_type_any(ANY(x) )} 

// an item is an integer, a float, a symbol, a string or a type
/* The go function for: extract_item(x:any,y:any) [status=0] */
func F_extract_item_any (x *ClaireAny,y *ClaireAny) *ClaireAny { 
  if (((((C_integer.Id() == x.Isa.Id()) || 
            (C_float.Id() == x.Isa.Id())) || 
          (x.Isa.IsIn(C_symbol) == CTRUE)) || 
        (C_string.Id() == x.Isa.Id())) || 
      (x.Isa.IsIn(C_type) == CTRUE)) { 
    return  x
    }  else if (x.Isa.IsIn(Core.C_global_variable) == CTRUE) { 
    return  F_extract_item_any(ANY(Core.F_CALL(C_value,ARGS(x.ToEID()))),y)
    } else {
    return  CNULL
    } 
  } 

// The EID go function for: extract_item @ any (throw: false) 
func E_extract_item_any (x EID,y EID) EID { 
  return F_extract_item_any(ANY(x),ANY(y) ).ToEID()} 

// version for X[...] which is the most complex case - note the extensibility
// patch.
/* The go function for: extract_pattern_nth(l:list,path:list) [status=1] */
func F_extract_pattern_nth_list (l *ClaireList,path *ClaireList) EID { 
  var Result EID
  { var m int = l.Length()
    { var x *ClaireAny = l.At(0)
      if (m == 1) { 
        { var y *ClaireAny
          var try_1 EID
          try_1 = F_extract_pattern_any(l.At(0),CNIL)
          if ErrorIn(try_1) {Result = try_1
          } else {
          y = ANY(try_1)
          if (y == CNULL) { 
            Result = EID{CNULL,0}
            } else {
            { var _CL_obj *ClaireParam = To_Param(new(ClaireParam).Is(C_Param))
              _CL_obj.Arg = C_array
              _CL_obj.Params = MakeConstantList(C_of.Id())
              _CL_obj.Args = MakeConstantList(MakeConstantSet(y).Id())
              Result = EID{_CL_obj.Id(),0}
              } 
            } 
          }
          } 
        }  else if (m == 2) { 
        if (((x == C_list.Id()) || 
              ((x == C_set.Id()) || 
                (x == C_subtype.Id()))) || 
            (C_class.Id() != x.Isa.Id())) { 
          { var y *ClaireAny
            var try_2 EID
            try_2 = F_extract_pattern_any(l.At(1),CNIL)
            if ErrorIn(try_2) {Result = try_2
            } else {
            y = ANY(try_2)
            { 
              h_index := ClEnv.Index
              h_base := ClEnv.Base
              if (y != CNULL) { 
                Result = Core.F_CALL(C_nth,ARGS(l.At(0).ToEID(),y.ToEID()))
                } else {
                Result = EID{CNULL,0}
                } 
              if ErrorIn(Result){ 
                ClEnv.Index = h_index
                ClEnv.Base = h_base
                Result = EID{CNULL,0}
                } 
              } 
            }
            } 
          } else {
          Result = EID{CNULL,0}
          } 
        } else {
        { var l1 *ClaireAny = l.At(1)
          { var l2 *ClaireList = ToList(OBJ(Core.F_CALL(C_args,ARGS(l.At(2).ToEID()))))
            { var l3 *ClaireList = ToType(C_any.Id()).EmptyList()
              { var n int = 1
                { var g0238 int = INT(Core.F_CALL(C_length,ARGS(l1.ToEID())))
                  Result= EID{CFALSE.Id(),0}
                  for (n <= g0238) { 
                    var loop_3 EID
                    _ = loop_3
                    { 
                    { var y *ClaireAny = l2.At(n-1)
                      var try_4 EID
                      { var arg_5 *ClaireAny
                        var try_6 EID
                        if (y.Isa.IsIn(C_Set) == CTRUE) { 
                          { var g0239 *Set = To_Set(y)
                            { var v *ClaireAny
                              var try_7 EID
                              { var arg_8 *ClaireList
                                var try_9 EID
                                { var arg_10 *ClaireAny
                                  var try_11 EID
                                  try_11 = Core.F_CALL(C_nth,ARGS(l1.ToEID(),EID{C__INT,IVAL(n)}))
                                  if ErrorIn(try_11) {try_9 = try_11
                                  } else {
                                  arg_10 = ANY(try_11)
                                  try_9 = EID{path.Copy().AddFast(arg_10).Id(),0}
                                  }
                                  } 
                                if ErrorIn(try_9) {try_7 = try_9
                                } else {
                                arg_8 = ToList(OBJ(try_9))
                                try_7 = F_extract_pattern_any(g0239.Args.At(0),arg_8)
                                }
                                } 
                              if ErrorIn(try_7) {try_6 = try_7
                              } else {
                              v = ANY(try_7)
                              if (v == C_void.Id()) { 
                                try_6 = EID{C_any.Id(),0}
                                }  else if (v.Isa.IsIn(C_Reference) == CTRUE) { 
                                { var g0241 *ClaireReference = To_Reference(v)
                                  { var z *ClaireReference = To_Reference(g0241.Copy().Id())
                                    z.Arg = CTRUE
                                    try_6 = EID{z.Id(),0}
                                    } 
                                  } 
                                } else {
                                { 
                                  var v_bag_arg *ClaireAny
                                  try_6= EID{ToType(CEMPTY.Id()).EmptySet().Id(),0}
                                  var try_12 EID
                                  if (v != CNULL) { 
                                    try_12 = v.ToEID()
                                    } else {
                                    try_12 = EVAL(g0239.Args.At(0))
                                    } 
                                  if ErrorIn(try_12) {try_6 = try_12
                                  } else {
                                  v_bag_arg = ANY(try_12)
                                  ToSet(OBJ(try_6)).AddFast(v_bag_arg)}
                                  } 
                                } 
                              }
                              } 
                            } 
                          } else {
                          { var arg_13 *ClaireAny
                            var try_14 EID
                            if (path.Length() != 0) { 
                              { var arg_15 *ClaireAny
                                var try_16 EID
                                try_16 = Core.F_CALL(C_nth,ARGS(l1.ToEID(),EID{C__INT,IVAL(n)}))
                                if ErrorIn(try_16) {try_14 = try_16
                                } else {
                                arg_15 = ANY(try_16)
                                try_14 = EID{path.AddFast(arg_15).Id(),0}
                                }
                                } 
                              } else {
                              try_14 = EID{CFALSE.Id(),0}
                              } 
                            if ErrorIn(try_14) {try_6 = try_14
                            } else {
                            arg_13 = ANY(try_14)
                            try_6 = F_extract_pattern_any(y,ToList(arg_13))
                            }
                            } 
                          } 
                        if ErrorIn(try_6) {try_4 = try_6
                        } else {
                        arg_5 = ANY(try_6)
                        try_4 = EID{l3.AddFast(arg_5).Id(),0}
                        }
                        } 
                      if ErrorIn(try_4) {loop_3 = try_4
                      } else {
                      l3 = ToList(OBJ(try_4))
                      loop_3 = EID{l3.Id(),0}
                      }
                      } 
                    if ErrorIn(loop_3) {Result = loop_3
                    break
                    } else {
                    n = (n+1)
                    }
                    } 
                  }
                  } 
                } 
              if !ErrorIn(Result) {
              if (l3.Memq(CNULL) == CTRUE) { 
                Result = EID{CNULL,0}
                } else {
                { var _CL_obj *ClaireParam = To_Param(new(ClaireParam).Is(C_Param))
                  _CL_obj.Arg = ToClass(x)
                  _CL_obj.Params = ToList(l1)
                  _CL_obj.Args = l3
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
  return Result} 

// The EID go function for: extract_pattern_nth @ list (throw: true) 
func E_extract_pattern_nth_list (l EID,path EID) EID { 
  return F_extract_pattern_nth_list(ToList(OBJ(l)),ToList(OBJ(path)) )} 

// we perform some pre-processing on x[l] at reading time to make evaluation easier
/* The go function for: iClaire/extract_class_call(self:class,l:list) [status=1] */
func F_extract_class_call_class (self *ClaireClass,l *ClaireList) EID { 
  var Result EID
  var g0251I *ClaireBoolean
  var try_1 EID
  { 
    var v_and1 *ClaireBoolean
    
    v_and1 = MakeBoolean((self.Id() == C_list.Id()) || (self.Id() == C_set.Id()) || (self.Id() == C_subtype.Id()))
    if (v_and1 == CFALSE) {try_1 = EID{CFALSE.Id(),0}
    } else { 
      v_and1 = Equal(MakeInteger(l.Length()).Id(),MakeInteger(1).Id())
      if (v_and1 == CFALSE) {try_1 = EID{CFALSE.Id(),0}
      } else { 
        var try_2 EID
        { var y *ClaireAny = l.At(0)
          { var z *ClaireAny
            var try_3 EID
            try_3 = F_extract_pattern_any(y,CNIL)
            if ErrorIn(try_3) {try_2 = try_3
            } else {
            z = ANY(try_3)
            if (y.Isa.IsIn(Core.C_global_variable) == CTRUE) { 
              { var g0244 *Core.GlobalVariable = Core.ToGlobalVariable(y)
                g0244 = Core.ToGlobalVariable(OBJ(Core.F_CALL(C_value,ARGS(l.At(0).ToEID()))))
                y = g0244.Id()
                } 
              } 
            { 
              var v_or6 *ClaireBoolean
              
              v_or6 = z.Isa.IsIn(C_type)
              if (v_or6 == CTRUE) {try_2 = EID{CTRUE.Id(),0}
              } else { 
                v_or6 = Equal(self.Id(),C_subtype.Id())
                if (v_or6 == CTRUE) {try_2 = EID{CTRUE.Id(),0}
                } else { 
                  if (y.Isa.IsIn(C_Call) == CTRUE) { 
                    { var g0245 *Call = To_Call(y)
                      v_or6 = MakeBoolean((g0245.Selector.Id() != C__equal.Id()) || (g0245.Args.Length() != 2))
                      } 
                    }  else if (y.Isa.IsIn(C_Tuple) == CTRUE) { 
                    v_or6 = CTRUE
                    } else {
                    v_or6 = CFALSE
                    } 
                  if (v_or6 == CTRUE) {try_2 = EID{CTRUE.Id(),0}
                  } else { 
                    try_2 = EID{CFALSE.Id(),0}} 
                  } 
                } 
              } 
            }
            } 
          } 
        if ErrorIn(try_2) {try_1 = try_2
        } else {
        v_and1 = ToBoolean(OBJ(try_2))
        if (v_and1 == CFALSE) {try_1 = EID{CFALSE.Id(),0}
        } else { 
          try_1 = EID{CTRUE.Id(),0}} 
        } 
      } 
    }
    } 
  if ErrorIn(try_1) {Result = try_1
  } else {
  g0251I = ToBoolean(OBJ(try_1))
  if (g0251I == CTRUE) { 
    { var _CL_obj *Call = To_Call(new(Call).Is(C_Call))
      _CL_obj.Selector = C_nth
      _CL_obj.Args = F_cons_any(self.Id(),l)
      Result = EID{_CL_obj.Id(),0}
      } 
    }  else if (self.Id() == C_lambda.Id()) { 
    if ((l.Length() == 2) && 
        ((l.At(0).Isa.IsIn(C_Do) == CTRUE) || 
            (l.At(0).Isa.IsIn(C_Variable) == CTRUE))) { 
      { var lv *ClaireList
        if (l.At(0).Isa.IsIn(C_Do) == CTRUE) { 
          { var v_in *ClaireList = ToList(OBJ(Core.F_CALL(C_args,ARGS(l.At(0).ToEID()))))
            { var v_out *ClaireList = v_in.Empty()
              { 
                var v *ClaireAny
                _ = v
                var v_support *ClaireList
                v_support = v_in
                v_len := v_support.Length()
                for i_it := 0; i_it < v_len; i_it++ { 
                  v = v_support.At(i_it)
                  if (v.Isa.IsIn(C_Variable) == CTRUE) { 
                    v_out.AddFast(v)
                    } 
                  } 
                } 
              lv = v_out
              } 
            } 
          } else {
          lv = MakeConstantList(l.At(0))
          } 
        Result = F_extract_signature_list(lv)
        if !ErrorIn(Result) {
        Result = F_lambda_I_list(lv,l.At(1))
        }
        } 
      } else {
      Result = ToException(Core.C_general_error.Make(MakeString("[113] Wrong lambda definition lambda[~S]").Id(),MakeConstantList(l.Id()).Id())).Close()
      } 
    } else {
    { var l1 *ClaireList = ToType(C_any.Id()).EmptyList()
      { var l2 *ClaireList = ToType(C_any.Id()).EmptyList()
        { var m int = l.Length()
          { var n int = 1
            { var g0247 int = m
              Result= EID{CFALSE.Id(),0}
              for (n <= g0247) { 
                var loop_4 EID
                _ = loop_4
                { 
                { var y *ClaireAny = l.At(n-1)
                  { var p *ClaireAny = CNULL
                    { var v *ClaireAny = CNULL
                      if (y.Isa.IsIn(C_Call) == CTRUE) { 
                        { var g0248 *Call = To_Call(y)
                          if ((g0248.Selector.Id() != C__equal.Id()) || 
                              (g0248.Args.Length() != 2)) { 
                            loop_4 = ToException(Core.C_general_error.Make(MakeString("[114] Wrong parametrization ~S").Id(),MakeConstantList(g0248.Id()).Id())).Close()
                            } else {
                            loop_4 = EID{CFALSE.Id(),0}
                            } 
                          if ErrorIn(loop_4) {Result = loop_4
                          break
                          } else {
                          var try_5 EID
                          try_5 = F_make_a_property_any(g0248.Args.At(0))
                          if ErrorIn(try_5) {loop_4 = try_5
                          Result = try_5
                          break
                          } else {
                          p = ANY(try_5)
                          loop_4 = p.ToEID()
                          { var _CL_obj *Set = To_Set(new(Set).Is(C_Set))
                            _CL_obj.Args = MakeConstantList(g0248.Args.At(1))
                            v = _CL_obj.Id()
                            } 
                          loop_4 = v.ToEID()
                          }}
                          } 
                        }  else if (y.Isa.IsIn(C_Vardef) == CTRUE) { 
                        { var g0249 *Vardef = To_Vardef(y)
                          var try_6 EID
                          try_6 = F_make_a_property_any(g0249.Pname.Id())
                          if ErrorIn(try_6) {loop_4 = try_6
                          Result = try_6
                          break
                          } else {
                          p = ANY(try_6)
                          loop_4 = p.ToEID()
                          v = g0249.Range.Id()
                          loop_4 = v.ToEID()
                          }
                          } 
                        } else {
                        var try_7 EID
                        try_7 = F_make_a_property_any(y)
                        if ErrorIn(try_7) {loop_4 = try_7
                        Result = try_7
                        break
                        } else {
                        p = ANY(try_7)
                        loop_4 = p.ToEID()
                        v = CEMPTY.Id()
                        loop_4 = v.ToEID()
                        }
                        } 
                      if ErrorIn(loop_4) {Result = loop_4
                      break
                      } else {
                      l1 = l1.AddFast(p)
                      l2 = l2.AddFast(v)
                      loop_4 = EID{l2.Id(),0}
                      }
                      } 
                    } 
                  } 
                if ErrorIn(loop_4) {Result = loop_4
                break
                } else {
                n = (n+1)
                }
                } 
              }
              } 
            } 
          if !ErrorIn(Result) {
          { var _CL_obj *Call = To_Call(new(Call).Is(C_Call))
            _CL_obj.Selector = C_nth
            { 
              var va_arg1 *Call
              var va_arg2 *ClaireList
              va_arg1 = _CL_obj
              { var arg_8 *ClaireList
                { 
                  var v_bag_arg *ClaireAny
                  arg_8= ToType(CEMPTY.Id()).EmptyList()
                  arg_8.AddFast(l1.Id())
                  { var _CL_obj *List = To_List(new(List).Is(C_List))
                    _CL_obj.Args = l2
                    v_bag_arg = _CL_obj.Id()
                    } 
                  arg_8.AddFast(v_bag_arg)} 
                va_arg2 = F_cons_any(self.Id(),arg_8)
                } 
              va_arg1.Args = va_arg2
              } 
            Result = EID{_CL_obj.Id(),0}
            } 
          }
          } 
        } 
      } 
    } 
  }
  return Result} 

// The EID go function for: iClaire/extract_class_call @ class (throw: true) 
func E_extract_class_call_class (self EID,l EID) EID { 
  return F_extract_class_call_class(ToClass(OBJ(self)),ToList(OBJ(l)) )} 

// extract the range (type and/or second-order function)
// lvar is the list of arguments that will serve as second-o. args
// ldef is the list of extra type variables that are defined in the sig.
/* The go function for: iClaire/extract_range(x:any,lvar:list,ldef:list) [status=1] */
func F_extract_range_any (x *ClaireAny,lvar *ClaireList,ldef *ClaireList) EID { 
  var Result EID
  var g0254I *ClaireBoolean
  { var arg_1 *ClaireBoolean
    if (x.Isa.IsIn(C_Call) == CTRUE) { 
      { var g0252 *Call = To_Call(x)
        arg_1 = MakeBoolean((g0252.Selector.Id() == C_nth.Id()) && (g0252.Args.At(0) == C_type.Id()))
        } 
      } else {
      arg_1 = CFALSE
      } 
    g0254I = arg_1.Not
    } 
  if (g0254I == CTRUE) { 
    { 
      var v_bag_arg *ClaireAny
      Result= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
      var try_2 EID
      try_2 = F_extract_type_any(x)
      if ErrorIn(try_2) {Result = try_2
      } else {
      v_bag_arg = ANY(try_2)
      ToList(OBJ(Result)).AddFast(v_bag_arg)
      ToList(OBJ(Result)).AddFast(CEMPTY.Id())}
      } 
    } else {
    
    { 
      var v *ClaireAny
      _ = v
      var v_support *ClaireList
      v_support = ldef
      v_len := v_support.Length()
      for i_it := 0; i_it < v_len; i_it++ { 
        v = v_support.At(i_it)
        { var r *ClaireReference = To_Reference(OBJ(Core.F_CALL(C_range,ARGS(v.ToEID()))))
          { var path *ClaireList = r.Args
            { var n int = path.Length()
              { var y *ClaireAny = lvar.At((r.Index+1)-1)
                { var i int = 1
                  { var g0253 int = n
                    for (i <= g0253) { 
                      { var _CL_obj *Call = To_Call(new(Call).Is(C_Call))
                        _CL_obj.Selector = ToProperty(Core.C__at.Id())
                        _CL_obj.Args = MakeConstantList(y,path.At(i-1))
                        y = _CL_obj.Id()
                        } 
                      i = (i+1)
                      } 
                    } 
                  } 
                { var arg_3 *Call
                  { var _CL_obj *Call = To_Call(new(Call).Is(C_Call))
                    _CL_obj.Selector = Core.C_member
                    _CL_obj.Args = MakeConstantList(y)
                    arg_3 = _CL_obj
                    } 
                  x = F_substitution_any(x,To_Variable(v),arg_3.Id())
                  } 
                } 
              } 
            } 
          } 
        } 
      } 
    { var lv2 *ClaireList = ToType(C_any.Id()).EmptyList()
      { 
        var v *ClaireAny
        _ = v
        var v_support *ClaireList
        v_support = lvar
        v_len := v_support.Length()
        for i_it := 0; i_it < v_len; i_it++ { 
          v = v_support.At(i_it)
          { var v2 *ClaireVariable
            { var _CL_obj *ClaireVariable = To_Variable(new(ClaireVariable).Is(C_Variable))
              _CL_obj.Pname = ToSymbol(OBJ(Core.F_CALL(C_mClaire_pname,ARGS(v.ToEID()))))
              _CL_obj.Range = ToType(C_type.Id())
              v2 = _CL_obj
              } 
            lv2 = lv2.AddFast(v2.Id())
            x = F_substitution_any(x,To_Variable(v),v2.Id())
            } 
          } 
        } 
      { var lb *ClaireLambda
        var try_4 EID
        try_4 = F_lambda_I_list(lv2,ToList(OBJ(Core.F_CALL(C_args,ARGS(x.ToEID())))).At(1))
        if ErrorIn(try_4) {Result = try_4
        } else {
        lb = ToLambda(OBJ(try_4))
        { var ur *ClaireAny = CNULL
          
          { 
            h_index := ClEnv.Index
            h_base := ClEnv.Base
            var try_5 EID
            { var arg_6 *ClaireList
              { 
                var v_list7 *ClaireList
                var v *ClaireAny
                var v_local7 *ClaireAny
                v_list7 = lvar
                arg_6 = CreateList(ToType(CEMPTY.Id()),v_list7.Length())
                for CLcount := 0; CLcount < v_list7.Length(); CLcount++{ 
                  v = v_list7.At(CLcount)
                  v_local7 = ANY(Core.F_CALL(C_range,ARGS(v.ToEID())))
                  arg_6.PutAt(CLcount,v_local7)
                  } 
                } 
              try_5 = F_apply_lambda(lb,arg_6)
              } 
            if ErrorIn(try_5) {Result = try_5
            } else {
            ur = ANY(try_5)
            Result = ur.ToEID()
            }
            if ErrorIn(Result){ 
              ClEnv.Index = h_index
              ClEnv.Base = h_base
              PRINC("The type expression ")
              Result = Core.F_CALL(C_print,ARGS(x.ToEID()))
              if !ErrorIn(Result) {
              PRINC(" is not valid ... \n")
              Result = EVOID
              }
              if !ErrorIn(Result) {
              PRINC("context: lambda = ")
              Result = Core.F_print_any(lb.Id())
              if !ErrorIn(Result) {
              PRINC(", lvars = ")
              { var arg_7 *ClaireList
                { 
                  var v_list8 *ClaireList
                  var v *ClaireAny
                  var v_local8 *ClaireAny
                  v_list8 = lvar
                  arg_7 = CreateList(ToType(CEMPTY.Id()),v_list8.Length())
                  for CLcount := 0; CLcount < v_list8.Length(); CLcount++{ 
                    v = v_list8.At(CLcount)
                    v_local8 = ANY(Core.F_CALL(C_range,ARGS(v.ToEID())))
                    arg_7.PutAt(CLcount,v_local8)
                    } 
                  } 
                Result = Core.F_print_any(arg_7.Id())
                } 
              if !ErrorIn(Result) {
              PRINC("\n")
              Result = EVOID
              }}
              if !ErrorIn(Result) {
              Result = ClEnv.Exception_I.Close()
              }}
              } 
            } 
          if !ErrorIn(Result) {
          if (ur.Isa.IsIn(C_type) != CTRUE) { 
            Result = ToException(Core.C_general_error.Make(MakeString("[115] the (resulting) range ~S is not a type").Id(),MakeConstantList(ur).Id())).Close()
            } else {
            Result = EID{CFALSE.Id(),0}
            } 
          if !ErrorIn(Result) {
          Result = EID{MakeConstantList(ur,lb.Id()).Id(),0}
          }}
          } 
        }
        } 
      } 
    } 
  return Result} 

// The EID go function for: iClaire/extract_range @ any (throw: true) 
func E_extract_range_any (x EID,lvar EID,ldef EID) EID { 
  return F_extract_range_any(ANY(x),ToList(OBJ(lvar)),ToList(OBJ(ldef)) )} 

// create a bitvector from a list of flags
/* The go function for: bit_vector(l:listargs) [status=1] */
func F_bit_vector_listargs2 (l *ClaireList) EID { 
  var Result EID
  { var d int = 0
    { 
      var x *ClaireAny
      _ = x
      Result= EID{CFALSE.Id(),0}
      var x_support *ClaireList
      x_support = ToList(l.Id())
      x_len := x_support.Length()
      for i_it := 0; i_it < x_len; i_it++ { 
        x = x_support.At(i_it)
        var loop_1 EID
        _ = loop_1
        var try_2 EID
        { var arg_3 int
          var try_4 EID
          try_4 = F__exp2_integer(ToInteger(x).Value)
          if ErrorIn(try_4) {try_2 = try_4
          } else {
          arg_3 = INT(try_4)
          try_2 = EID{C__INT,IVAL((d+arg_3))}
          }
          } 
        if ErrorIn(try_2) {Result = try_2
        break
        } else {
        d = INT(try_2)
        loop_1 = EID{C__INT,IVAL(d)}
        }
        } 
      } 
    if !ErrorIn(Result) {
    Result = EID{C__INT,IVAL(d)}
    }
    } 
  return Result} 

// The EID go function for: bit_vector @ listargs (throw: true) 
func E_bit_vector_listargs2 (l EID) EID { 
  return F_bit_vector_listargs2(ToList(OBJ(l)) )} 

// parse the body and return (status, functional, body)
// the input is  body | (function!(f) | function!(f,s)) < | body> opt
// CLAIRE4: status is -1 : unknown, 0: no error, 1: an error may be thrown
//
/* The go function for: iClaire/extract_status(x:any) [status=1] */
func F_extract_status_any (x *ClaireAny) EID { 
  var Result EID
  { var s int = -1
    { var f *ClaireAny
      var g0260I *ClaireBoolean
      if (x.Isa.IsIn(C_Call) == CTRUE) { 
        { var g0255 *Call = To_Call(x)
          g0260I = Equal(g0255.Selector.Id(),C_function_I.Id())
          } 
        } else {
        g0260I = CFALSE
        } 
      if (g0260I == CTRUE) { 
        f = x
        } else {
        f = CNULL
        } 
      if (x.Isa.IsIn(C_And) == CTRUE) { 
        { var g0256 *And = To_And(x)
          { var y *ClaireAny = g0256.Args.At(0)
            var g0261I *ClaireBoolean
            if (y.Isa.IsIn(C_Call) == CTRUE) { 
              { var g0257 *Call = To_Call(y)
                g0261I = Equal(g0257.Selector.Id(),C_function_I.Id())
                } 
              } else {
              g0261I = CFALSE
              } 
            if (g0261I == CTRUE) { 
              f = y
              g0256 = To_And(g0256.Args.At(1))
              } 
            } 
          x = g0256.Id()
          } 
        }  else if (x.Isa.IsIn(C_Call) == CTRUE) { 
        { var g0258 *Call = To_Call(x)
          if (g0258.Selector.Id() == C_function_I.Id()) { 
            g0258 = To_Call(C_body.Id())
            } 
          x = g0258.Id()
          } 
        } else {
        
        } 
      if (f != CNULL) { 
        x = C_body.Id()
        if (ToList(OBJ(Core.F_CALL(C_args,ARGS(f.ToEID())))).Length() > 1) { 
          s = 1
          } else {
          s = 0
          } 
        var try_1 EID
        { var arg_2 *ClaireString
          var try_3 EID
          { var arg_4 *ClaireSymbol
            var try_5 EID
            try_5 = F_extract_symbol_any(ToList(OBJ(Core.F_CALL(C_args,ARGS(f.ToEID())))).At(0))
            if ErrorIn(try_5) {try_3 = try_5
            } else {
            arg_4 = ToSymbol(OBJ(try_5))
            try_3 = EID{arg_4.String_I().Id(),0}
            }
            } 
          if ErrorIn(try_3) {try_1 = try_3
          } else {
          arg_2 = ToString(OBJ(try_3))
          try_1 = F_imported_function_string(arg_2).ToEID()
          }
          } 
        if ErrorIn(try_1) {Result = try_1
        } else {
        f = ANY(try_1)
        Result = f.ToEID()
        }
        } else {
        Result = EID{CFALSE.Id(),0}
        } 
      if !ErrorIn(Result) {
      Result = EID{MakeConstantList(MakeInteger(s).Id(),f,x).Id(),0}
      }
      } 
    } 
  return Result} 

// The EID go function for: iClaire/extract_status @ any (throw: true) 
func E_extract_status_any (x EID) EID { 
  return F_extract_status_any(ANY(x) )} 

// new in CLAIRE4 : create a function with a syntactic marker ! for imported
/* The go function for: imported_function(s:string) [status=0] */
func F_imported_function_string (s *ClaireString) *ClaireFunction { 
  return  F_make_function_string(F_append_string(MakeString("#"),s))
  } 

// The EID go function for: imported_function @ string (throw: false) 
func E_imported_function_string (s EID) EID { 
  return F_imported_function_string(ToString(OBJ(s)) ).ToEID()} 

// cleans a pattern into a type
/* The go function for: iClaire/type!(x:any) [status=0] */
func F_type_I_any (x *ClaireAny) *ClaireType { 
  var Result *ClaireType
  if (x.Isa.IsIn(C_list) == CTRUE) { 
    { var g0262 *ClaireList = ToList(x)
      { 
        var v_list3 *ClaireList
        var y *ClaireAny
        var v_local3 *ClaireAny
        v_list3 = g0262
        Result = ToType(CreateList(ToType(CEMPTY.Id()),v_list3.Length()).Id())
        for CLcount := 0; CLcount < v_list3.Length(); CLcount++{ 
          y = v_list3.At(CLcount)
          v_local3 = F_type_I_any(y).Id()
          ToList(Result.Id()).PutAt(CLcount,v_local3)
          } 
        } 
      } 
    }  else if (x.Isa.IsIn(C_Param) == CTRUE) { 
    { var g0263 *ClaireParam = To_Param(x)
      { var _CL_obj *ClaireParam = To_Param(new(ClaireParam).Is(C_Param))
        _CL_obj.Arg = g0263.Arg
        _CL_obj.Params = g0263.Params
        { 
          var va_arg1 *ClaireParam
          var va_arg2 *ClaireList
          va_arg1 = _CL_obj
          { 
            var v_list5 *ClaireList
            var y *ClaireAny
            var v_local5 *ClaireAny
            v_list5 = g0263.Args
            va_arg2 = CreateList(ToType(CEMPTY.Id()),v_list5.Length())
            for CLcount := 0; CLcount < v_list5.Length(); CLcount++{ 
              y = v_list5.At(CLcount)
              v_local5 = F_type_I_any(y).Id()
              va_arg2.PutAt(CLcount,v_local5)
              } 
            } 
          va_arg1.Args = va_arg2
          } 
        Result = ToType(_CL_obj.Id())
        } 
      } 
    }  else if (x.Isa.IsIn(C_Reference) == CTRUE) { 
    Result = ToType(C_any.Id())
    }  else if (x.Isa.IsIn(C_type) == CTRUE) { 
    { var g0265 *ClaireType = ToType(x)
      Result = g0265
      } 
    } else {
    Result = ToType(C_any.Id())
    } 
  return Result} 

// The EID go function for: iClaire/type! @ any (throw: false) 
func E_type_I_any (x EID) EID { 
  return EID{F_type_I_any(ANY(x) ).Id(),0}} 

// for instance patterns
// creates a table
// to do in later versions: use an array if direct indexed access
// in the meanwhile, arrays of float should be used with care (indexed arrays)
//
/* The go function for: self_eval(self:Defarray) [status=1] */
func (self *Defarray) SelfEval () EID { 
  var Result EID
  { var a *ClaireList = self.Arg.Args
    { var ar *ClaireTable
      var try_1 EID
      { var arg_2 *ClaireSymbol
        var try_3 EID
        try_3 = F_extract_symbol_any(a.At(0))
        if ErrorIn(try_3) {try_1 = try_3
        } else {
        arg_2 = ToSymbol(OBJ(try_3))
        try_1 = new(ClaireTable).IsNamed(C_table,arg_2).ToEID()
        }
        } 
      if ErrorIn(try_1) {Result = try_1
      } else {
      ar = ToTable(OBJ(try_1))
      { var v *ClaireVariable = To_Variable(a.At(1))
        { var s *ClaireTypeExpression
          var try_4 EID
          try_4 = F_extract_type_any(v.Range.Id())
          if ErrorIn(try_4) {Result = try_4
          } else {
          s = ToTypeExpression(OBJ(try_4))
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
                try_7 = F_iClaire_lexical_index_any2(self.Body,l,0,CTRUE)
                if ErrorIn(try_7) {try_5 = try_7
                } else {
                b = ANY(try_7)
                var g0275I *ClaireBoolean
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
                      if (F_occurrence_any(b,To_Variable(va)) > 0) { 
                        arg_8 = CTRUE.Id()
                        break
                        } 
                      } 
                    } 
                  g0275I = F_boolean_I_any(arg_8)
                  } 
                if (g0275I == CTRUE) { 
                  try_5 = F_lambda_I_list(l,b)
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
              var try_9 EID
              if (e.Isa.IsIn(C_lambda) == CTRUE) { 
                try_9 = EID{CNULL,0}
                } else {
                try_9 = EVAL(self.Body)
                } 
              if ErrorIn(try_9) {Result = try_9
              } else {
              d = ANY(try_9)
              { 
                var va_arg1 *ClaireRelation
                var va_arg2 *ClaireType
                va_arg1 = ToRelation(ar.Id())
                var try_10 EID
                try_10 = F_extract_pattern_any(self.SetArg,CNIL)
                if ErrorIn(try_10) {Result = try_10
                } else {
                va_arg2 = ToType(OBJ(try_10))
                va_arg1.Range = va_arg2
                Result = EID{va_arg2.Id(),0}
                }
                } 
              if !ErrorIn(Result) {
              if (ar.Range.Id() == CNULL) { 
                { var _CL_obj *Core.RangeError = Core.ToRangeError(new(Core.RangeError).Is(Core.C_range_error))
                  _CL_obj.Cause = C_table.Id()
                  _CL_obj.Arg = self.SetArg
                  Result = Core.F_write_property(C_Language_wrong,ToObject(_CL_obj.Id()),C_type.Id())
                  if !ErrorIn(Result) {
                  Result = _CL_obj.Close()
                  }
                  } 
                } else {
                Result = EID{CFALSE.Id(),0}
                } 
              if !ErrorIn(Result) {
              if ((d == CNULL) && 
                  ((ar.Range.Included(ToType(C_integer.Id())) == CTRUE) || 
                      (ar.Range.Included(ToType(C_float.Id())) == CTRUE))) { 
                Core.F_tformat_string(MakeString("=== CLAIRE4 Warning: unknown not allowed as a default for table with range ~S\n "),0,MakeConstantList(ar.Range.Id()))
                } 
              if (d != CNULL) { 
                if (ar.Range.Contains(d) != CTRUE) { 
                  { var _CL_obj *Core.RangeError = Core.ToRangeError(new(Core.RangeError).Is(Core.C_range_error))
                    _CL_obj.Cause = ar.Id()
                    _CL_obj.Arg = d
                    Result = Core.F_write_property(C_Language_wrong,ToObject(_CL_obj.Id()),ar.Range.Id())
                    if !ErrorIn(Result) {
                    Result = _CL_obj.Close()
                    }
                    } 
                  } else {
                  Result = EID{CFALSE.Id(),0}
                  } 
                }  else if (ar.Range.Included(ToType(C_integer.Id())) == CTRUE) { 
                d = MakeInteger(0).Id()
                Result = d.ToEID()
                }  else if (ar.Range.Included(ToType(C_float.Id())) == CTRUE) { 
                d = MakeFloat(0).Id()
                Result = d.ToEID()
                } else {
                Result = EID{CFALSE.Id(),0}
                } 
              if !ErrorIn(Result) {
              v.Range = ToType(s.Id())
              Result = F_attach_comment_any(ar.Id())
              if !ErrorIn(Result) {
              if (ar.Range.Class_I().IsIn(C_set) == CTRUE) { 
                ar.Multivalued_ask = CTRUE
                } 
              if (a.Length() == 2) { 
                ar.Domain = ToType(s.Id())
                if (s.Isa.IsIn(C_Interval) == CTRUE) { 
                  { var g0269 *ClaireInterval = To_Interval(s.Id())
                    ar.Params = MakeInteger((g0269.Arg1-1)).Id()
                    ar.Graph = Core.F_typed_copy_list_type(ar.Range,Core.F_size_Interval(g0269),d).Id()
                    } 
                  } else {
                  ar.Params = C_any.Id()
                  ar.GraphInit()
                  } 
                if (e.Isa.IsIn(C_lambda) == CTRUE) { 
                  { var g0271 *ClaireLambda = ToLambda(e)
                    { 
                      var y *ClaireAny
                      _ = y
                      Result= EID{CFALSE.Id(),0}
                      var y_support *ClaireList
                      var try_11 EID
                      try_11 = Core.F_enumerate_any(ar.Domain.Id())
                      if ErrorIn(try_11) {Result = try_11
                      } else {
                      y_support = ToList(OBJ(try_11))
                      y_len := y_support.Length()
                      for i_it := 0; i_it < y_len; i_it++ { 
                        y = y_support.At(i_it)
                        var loop_12 EID
                        _ = loop_12
                        { var arg_13 *ClaireAny
                          var try_14 EID
                          try_14 = Core.F_funcall_lambda1(g0271,y)
                          if ErrorIn(try_14) {loop_12 = try_14
                          } else {
                          arg_13 = ANY(try_14)
                          loop_12 = Core.F_nth_equal_table1(ar,y,arg_13)
                          }
                          } 
                        if ErrorIn(loop_12) {Result = loop_12
                        break
                        } else {
                        }}
                        } 
                      } 
                    } 
                  } else {
                  { 
                    var va_arg1 *ClaireTable
                    var va_arg2 *ClaireAny
                    va_arg1 = ar
                    va_arg2 = d
                    va_arg1.Default = va_arg2
                    Result = va_arg2.ToEID()
                    } 
                  } 
                } else {
                { var s2 *ClaireTypeExpression
                  var try_15 EID
                  try_15 = F_extract_type_any(To_Variable(a.At(2)).Range.Id())
                  if ErrorIn(try_15) {Result = try_15
                  } else {
                  s2 = ToTypeExpression(OBJ(try_15))
                  ar.Domain = ToType(MakeConstantList(s.Id(),s2.Id()).Tuple_I().Id())
                  To_Variable(a.At(2)).Range = ToType(s2.Id())
                  if ((s.Isa.IsIn(C_Interval) == CTRUE) && 
                      (s2.Isa.IsIn(C_Interval) == CTRUE)) { 
                    
                    { 
                      var va_arg1 *ClaireTable
                      var va_arg2 *ClaireAny
                      va_arg1 = ar
                      var try_16 EID
                      { 
                        var v_bag_arg *ClaireAny
                        try_16= EID{ToType(C_integer.Id()).EmptyList().Id(),0}
                        var try_17 EID
                        try_17 = Core.F_CALL(C_size,ARGS(EID{s2.Id(),0}))
                        if ErrorIn(try_17) {try_16 = try_17
                        } else {
                        v_bag_arg = ANY(try_17)
                        ToList(OBJ(try_16)).AddFast(v_bag_arg)
                        var try_18 EID
                        { var arg_19 int
                          var try_20 EID
                          { var arg_21 int
                            var try_22 EID
                            { var arg_23 *ClaireAny
                              var try_24 EID
                              try_24 = Core.F_CALL(C_size,ARGS(EID{s2.Id(),0}))
                              if ErrorIn(try_24) {try_22 = try_24
                              } else {
                              arg_23 = ANY(try_24)
                              try_22 = EID{C__INT,IVAL((To_Interval(s.Id()).Arg1*ToInteger(arg_23).Value))}
                              }
                              } 
                            if ErrorIn(try_22) {try_20 = try_22
                            } else {
                            arg_21 = INT(try_22)
                            try_20 = EID{C__INT,IVAL((arg_21+To_Interval(s2.Id()).Arg1))}
                            }
                            } 
                          if ErrorIn(try_20) {try_18 = try_20
                          } else {
                          arg_19 = INT(try_20)
                          try_18 = EID{C__INT,IVAL((arg_19-1))}
                          }
                          } 
                        if ErrorIn(try_18) {try_16 = try_18
                        } else {
                        v_bag_arg = ANY(try_18)
                        ToList(OBJ(try_16)).AddFast(v_bag_arg)}}
                        } 
                      if ErrorIn(try_16) {Result = try_16
                      } else {
                      va_arg2 = ANY(try_16)
                      va_arg1.Params = va_arg2
                      Result = va_arg2.ToEID()
                      }
                      } 
                    if !ErrorIn(Result) {
                    { 
                      var va_arg1 *ClaireTable
                      var va_arg2 *ClaireAny
                      va_arg1 = ar
                      var try_25 EID
                      { var arg_26 int
                        var try_27 EID
                        { var arg_28 *ClaireAny
                          var try_30 EID
                          try_30 = Core.F_CALL(C_size,ARGS(EID{s.Id(),0}))
                          if ErrorIn(try_30) {try_27 = try_30
                          } else {
                          arg_28 = ANY(try_30)
                          { var arg_29 *ClaireAny
                            var try_31 EID
                            try_31 = Core.F_CALL(C_size,ARGS(EID{s2.Id(),0}))
                            if ErrorIn(try_31) {try_27 = try_31
                            } else {
                            arg_29 = ANY(try_31)
                            try_27 = F_times_integer(ToInteger(arg_28).Value,ToInteger(arg_29).Value)
                            }
                            } 
                          }
                          } 
                        if ErrorIn(try_27) {try_25 = try_27
                        } else {
                        arg_26 = INT(try_27)
                        try_25 = EID{Core.F_typed_copy_list_type(ar.Range,arg_26,d).Id(),0}
                        }
                        } 
                      if ErrorIn(try_25) {Result = try_25
                      } else {
                      va_arg2 = ANY(try_25)
                      va_arg1.Graph = va_arg2
                      Result = va_arg2.ToEID()
                      }
                      } 
                    }
                    } else {
                    
                    ar.Params = C_any.Id()
                    ar.GraphInit()
                    Result = EVOID
                    } 
                  if !ErrorIn(Result) {
                  
                  if (e.Isa.IsIn(C_lambda) == CTRUE) { 
                    { var g0273 *ClaireLambda = ToLambda(e)
                      { 
                        var y1 *ClaireAny
                        _ = y1
                        Result= EID{CFALSE.Id(),0}
                        var y1_support *ClaireList
                        var try_32 EID
                        try_32 = Core.F_enumerate_any(s.Id())
                        if ErrorIn(try_32) {Result = try_32
                        } else {
                        y1_support = ToList(OBJ(try_32))
                        y1_len := y1_support.Length()
                        for i_it := 0; i_it < y1_len; i_it++ { 
                          y1 = y1_support.At(i_it)
                          var loop_33 EID
                          _ = loop_33
                          { 
                            var y2 *ClaireAny
                            _ = y2
                            loop_33= EID{CFALSE.Id(),0}
                            var y2_support *ClaireList
                            var try_34 EID
                            try_34 = Core.F_enumerate_any(s2.Id())
                            if ErrorIn(try_34) {loop_33 = try_34
                            } else {
                            y2_support = ToList(OBJ(try_34))
                            y2_len := y2_support.Length()
                            for i_it := 0; i_it < y2_len; i_it++ { 
                              y2 = y2_support.At(i_it)
                              var loop_35 EID
                              _ = loop_35
                              { var arg_36 *ClaireAny
                                var try_37 EID
                                try_37 = Core.F_CALL(C_funcall,ARGS(EID{g0273.Id(),0},y1.ToEID(),y2.ToEID()))
                                if ErrorIn(try_37) {loop_35 = try_37
                                } else {
                                arg_36 = ANY(try_37)
                                loop_35 = Core.F_nth_equal_table2(ar,y1,y2,arg_36)
                                }
                                } 
                              if ErrorIn(loop_35) {loop_33 = loop_35
                              break
                              } else {
                              }}
                              } 
                            } 
                          if ErrorIn(loop_33) {Result = loop_33
                          break
                          } else {
                          }}
                          } 
                        } 
                      } 
                    } else {
                    { 
                      var va_arg1 *ClaireTable
                      var va_arg2 *ClaireAny
                      va_arg1 = ar
                      va_arg2 = d
                      va_arg1.Default = va_arg2
                      Result = va_arg2.ToEID()
                      } 
                    } 
                  }
                  }
                  } 
                } 
              if !ErrorIn(Result) {
              Result = EID{ar.Id(),0}
              }}}}}
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

// The EID go function for: self_eval @ Defarray (throw: true) 
func E_self_eval_Defarray (self EID) EID { 
  return To_Defarray(OBJ(self)).SelfEval( )} 

// The EVAL go function for: Defarray 
func EVAL_Defarray (x *ClaireAny) EID { 
   return To_Defarray(x).SelfEval()} 

// ------------------ NEW in v3.2 : definition of rules -----------------------
//
// a demon is a lambda with a name and a priority
/* The go function for: self_print(self:demon) [status=1] */
func (self *LanguageDemon) SelfPrint () EID { 
  var Result EID
  self.Pname.Princ()
  Result = EVOID
  return Result} 

// The EID go function for: self_print @ demon (throw: true) 
func E_self_print_demon (self EID) EID { 
  return ToLanguageDemon(OBJ(self)).SelfPrint( )} 

/* The go function for: funcall(self:demon,x:any,y:any) [status=1] */
func F_funcall_demon1 (self *LanguageDemon,x *ClaireAny,y *ClaireAny) EID { 
  var Result EID
  Result = Core.F_CALL(C_funcall,ARGS(EID{self.Formula.Id(),0},x.ToEID(),y.ToEID()))
  return Result} 

// The EID go function for: funcall @ list<type_expression>(demon, any, any) (throw: true) 
func E_funcall_demon1 (self EID,x EID,y EID) EID { 
  return F_funcall_demon1(ToLanguageDemon(OBJ(self)),ANY(x),ANY(y) )} 

/* The go function for: funcall(self:demon,x:any,y:any,z:any) [status=1] */
func F_funcall_demon2 (self *LanguageDemon,x *ClaireAny,y *ClaireAny,z *ClaireAny) EID { 
  var Result EID
  Result = Core.F_CALL(C_funcall,ARGS(EID{self.Formula.Id(),0},
    x.ToEID(),
    y.ToEID(),
    z.ToEID()))
  return Result} 

// The EID go function for: funcall @ list<type_expression>(demon, any, any, any) (throw: true) 
func E_funcall_demon2 (self EID,x EID,y EID,z EID) EID { 
  return F_funcall_demon2(ToLanguageDemon(OBJ(self)),
    ANY(x),
    ANY(y),
    ANY(z) )} 

// in the interpreted mode we store the list of demons using a table
// list of relevant demons
// the last rule/axiom that was defined on each relation
// this is used to find when the relation may be compiled
// list of involved relations
// compile(ru) => may compile(r)
// evaluate a rule definition: create a new demon and, if needed, the if_write 
// function
/* The go function for: self_eval(self:Defrule) [status=1] */
func (self *Defrule) SelfEval () EID { 
  var Result EID
  if (self.Args.At(0) != ClEnv.Id()) { 
    Result = Core.F_CALL(C_eval_rule,ARGS(EID{self.Id(),0}))
    } else {
    { var _Zcondition *ClaireAny = self.Arg
      { var ru *ClaireAny = self.Ident.Value()
        ru.Isa = C_Language_rule_object
        C_Language_rule_object.Instances.AddFast(ru)
        { var g0276 *ClaireTuple
          var try_1 EID
          try_1 = F_make_filter_any(_Zcondition)
          if ErrorIn(try_1) {Result = try_1
          } else {
          g0276 = ToTuple(OBJ(try_1))
          { var R *ClaireAny = ToList(g0276.Id()).At(0)
            { var lvar *ClaireAny = ToList(g0276.Id()).At(1)
              { var d *LanguageDemon
                var try_2 EID
                { var arg_3 *ClaireAny
                  var try_4 EID
                  try_4 = F_iClaire_lexical_index_any2(self.Body,ToList(lvar),0,CTRUE)
                  if ErrorIn(try_4) {try_2 = try_4
                  } else {
                  arg_3 = ANY(try_4)
                  try_2 = F_make_demon_relation(ToRelation(R),
                    ToSymbol(OBJ(Core.F_CALL(C_name,ARGS(ru.ToEID())))),
                    ToList(lvar),
                    _Zcondition,
                    arg_3)
                  }
                  } 
                if ErrorIn(try_2) {Result = try_2
                } else {
                d = ToLanguageDemon(OBJ(try_2))
                if (C_function.Id() == Core.F_owner_any(ANY(Core.F_CALL(C_if_write,ARGS(R.ToEID())))).Id()) { 
                  Result = ToException(Core.C_general_error.Make(MakeString("cannot define a new rule on ~S which is closed").Id(),MakeConstantList(R).Id())).Close()
                  } else {
                  Result = EID{CFALSE.Id(),0}
                  } 
                if !ErrorIn(Result) {
                
                Core.F_put_table(C_demons,R,ToList(Core.F_get_table(C_demons,R)).AddFast(d.Id()).Id())
                Result = Core.F_nth_put_table(C_Language_last_rule,R,ru)
                if !ErrorIn(Result) {
                if (ToList(Core.F_get_table(C_demons,R)).Length() == 1) { 
                  Result = F_eval_if_write_relation(ToRelation(R))
                  } else {
                  Result = EID{CFALSE.Id(),0}
                  } 
                if !ErrorIn(Result) {
                var g0279I *ClaireBoolean
                if (R.Isa.IsIn(C_property) == CTRUE) { 
                  { var g0278 *ClaireProperty = ToProperty(R)
                    g0279I = Equal(MakeInteger(g0278.Restrictions.Length()).Id(),MakeInteger(0).Id())
                    } 
                  } else {
                  g0279I = CFALSE
                  } 
                if (g0279I == CTRUE) { 
                  F_eventMethod_property(ToProperty(R))
                  } 
                Result = ru.ToEID()
                }}}
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

// The EID go function for: self_eval @ Defrule (throw: true) 
func E_self_eval_Defrule (self EID) EID { 
  return To_Defrule(OBJ(self)).SelfEval( )} 

// The EVAL go function for: Defrule 
func EVAL_Defrule (x *ClaireAny) EID { 
   return To_Defrule(x).SelfEval()} 

// an eventMethod is a property whose unique (?) restriction is a method
/* The go function for: eventMethod?(r:relation) [status=0] */
func F_eventMethod_ask_relation2 (r *ClaireRelation) *ClaireBoolean { 
  var Result *ClaireBoolean
  if (r.Isa.IsIn(C_property) == CTRUE) { 
    { var g0280 *ClaireProperty = ToProperty(r.Id())
      { var arg_1 *ClaireAny
        { 
          var x *ClaireRestriction
          _ = x
          var x_iter *ClaireAny
          arg_1= CFALSE.Id()
          for _,x_iter = range(g0280.Restrictions.ValuesO()){ 
            x = ToRestriction(x_iter)
            if (C_slot.Id() == x.Isa.Id()) { 
              arg_1 = CTRUE.Id()
              break
              } 
            } 
          } 
        Result = Core.F_not_any(arg_1)
        } 
      } 
    } else {
    Result = CFALSE
    } 
  return Result} 

// The EID go function for: eventMethod? @ relation (throw: false) 
func E_eventMethod_ask_relation2 (r EID) EID { 
  return EID{F_eventMethod_ask_relation2(ToRelation(OBJ(r)) ).Id(),0}} 

// check that condition is either a filter or the conjunction of a filter and a 
// condition
// a filter is R(x) := y | R(x) := (y <- z) | R(x) :add y | P(x,y)
// R(x) is x.r or A[x]
// the list of variable is of length 3 when R is mono-valued, whether we use a <- filter or a regular := 
/* The go function for: make_filter(cond:any) [status=1] */
func F_make_filter_any (cond *ClaireAny) EID { 
  var Result EID
  { var c *ClaireAny
    if (cond.Isa.IsIn(C_And) == CTRUE) { 
      { var g0281 *And = To_And(cond)
        c = g0281.Args.At(0)
        } 
      } else {
      c = cond
      } 
    
    var g0287I *ClaireBoolean
    if (c.Isa.IsIn(C_Call) == CTRUE) { 
      { var g0283 *Call = To_Call(c)
        g0287I = MakeBoolean(((g0283.Selector.Id() == Core.C_write.Id()) || 
            (g0283.Selector.Id() == C_nth_equal.Id())) && (g0283.Args.At(0).Isa.IsIn(C_relation) == CTRUE))
        } 
      } else {
      g0287I = CFALSE
      } 
    if (g0287I == CTRUE) { 
      { var R *ClaireRelation = ToRelation(ToList(OBJ(Core.F_CALL(C_args,ARGS(c.ToEID())))).At(0))
        { var x *ClaireVariable
          var try_1 EID
          { var _CL_obj *ClaireVariable = To_Variable(new(ClaireVariable).Is(C_Variable))
            { 
              var va_arg1 *ClaireVariable
              var va_arg2 *ClaireSymbol
              va_arg1 = _CL_obj
              var try_2 EID
              try_2 = F_extract_symbol_any(ToList(OBJ(Core.F_CALL(C_args,ARGS(c.ToEID())))).At(1))
              if ErrorIn(try_2) {try_1 = try_2
              } else {
              va_arg2 = ToSymbol(OBJ(try_2))
              va_arg1.Pname = va_arg2
              try_1 = EID{va_arg2.Id(),0}
              }
              } 
            if !ErrorIn(try_1) {
            _CL_obj.Range = R.Domain
            try_1 = EID{_CL_obj.Id(),0}
            }
            } 
          if ErrorIn(try_1) {Result = try_1
          } else {
          x = To_Variable(OBJ(try_1))
          { var y1 *ClaireAny = ToList(OBJ(Core.F_CALL(C_args,ARGS(c.ToEID())))).At(2)
            if (R.Multivalued_ask == CTRUE) { 
              Result = ToException(Core.C_general_error.Make(MakeString("[188] wrong event filter ~S for multi-valued relation").Id(),MakeConstantList(c,R.Id()).Id())).Close()
              } else {
              Result = EID{CFALSE.Id(),0}
              } 
            if !ErrorIn(Result) {
            var g0288I *ClaireBoolean
            if (y1.Isa.IsIn(C_Call) == CTRUE) { 
              { var g0284 *Call = To_Call(y1)
                g0288I = Equal(g0284.Selector.Id(),C__inf_dash.Id())
                } 
              } else {
              g0288I = CFALSE
              } 
            if (g0288I == CTRUE) { 
              { 
                var v_bag_arg *ClaireAny
                Result= EID{MakeEmptyTuple().Id(),0}
                ToTuple(OBJ(Result)).AddFast(R.Id())
                var try_3 EID
                { 
                  var v_bag_arg *ClaireAny
                  try_3= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                  ToList(OBJ(try_3)).AddFast(x.Id())
                  var try_4 EID
                  { var _CL_obj *ClaireVariable = To_Variable(new(ClaireVariable).Is(C_Variable))
                    { 
                      var va_arg1 *ClaireVariable
                      var va_arg2 *ClaireSymbol
                      va_arg1 = _CL_obj
                      var try_5 EID
                      try_5 = F_extract_symbol_any(ToList(OBJ(Core.F_CALL(C_args,ARGS(y1.ToEID())))).At(0))
                      if ErrorIn(try_5) {try_4 = try_5
                      } else {
                      va_arg2 = ToSymbol(OBJ(try_5))
                      va_arg1.Pname = va_arg2
                      try_4 = EID{va_arg2.Id(),0}
                      }
                      } 
                    if !ErrorIn(try_4) {
                    _CL_obj.Range = R.Range
                    try_4 = EID{_CL_obj.Id(),0}
                    }
                    } 
                  if ErrorIn(try_4) {try_3 = try_4
                  } else {
                  v_bag_arg = ANY(try_4)
                  ToList(OBJ(try_3)).AddFast(v_bag_arg)
                  var try_6 EID
                  { var _CL_obj *ClaireVariable = To_Variable(new(ClaireVariable).Is(C_Variable))
                    { 
                      var va_arg1 *ClaireVariable
                      var va_arg2 *ClaireSymbol
                      va_arg1 = _CL_obj
                      var try_7 EID
                      try_7 = F_extract_symbol_any(ToList(OBJ(Core.F_CALL(C_args,ARGS(y1.ToEID())))).At(1))
                      if ErrorIn(try_7) {try_6 = try_7
                      } else {
                      va_arg2 = ToSymbol(OBJ(try_7))
                      va_arg1.Pname = va_arg2
                      try_6 = EID{va_arg2.Id(),0}
                      }
                      } 
                    if !ErrorIn(try_6) {
                    _CL_obj.Range = R.Range
                    try_6 = EID{_CL_obj.Id(),0}
                    }
                    } 
                  if ErrorIn(try_6) {try_3 = try_6
                  } else {
                  v_bag_arg = ANY(try_6)
                  ToList(OBJ(try_3)).AddFast(v_bag_arg)}}
                  } 
                if ErrorIn(try_3) {Result = try_3
                } else {
                v_bag_arg = ANY(try_3)
                ToTuple(OBJ(Result)).AddFast(v_bag_arg)}
                } 
              } else {
              { 
                var v_bag_arg *ClaireAny
                Result= EID{MakeEmptyTuple().Id(),0}
                ToTuple(OBJ(Result)).AddFast(R.Id())
                var try_8 EID
                { 
                  var v_bag_arg *ClaireAny
                  try_8= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                  ToList(OBJ(try_8)).AddFast(x.Id())
                  var try_9 EID
                  { var _CL_obj *ClaireVariable = To_Variable(new(ClaireVariable).Is(C_Variable))
                    { 
                      var va_arg1 *ClaireVariable
                      var va_arg2 *ClaireSymbol
                      va_arg1 = _CL_obj
                      var try_10 EID
                      try_10 = F_extract_symbol_any(y1)
                      if ErrorIn(try_10) {try_9 = try_10
                      } else {
                      va_arg2 = ToSymbol(OBJ(try_10))
                      va_arg1.Pname = va_arg2
                      try_9 = EID{va_arg2.Id(),0}
                      }
                      } 
                    if !ErrorIn(try_9) {
                    _CL_obj.Range = F_safeRange_relation(R)
                    try_9 = EID{_CL_obj.Id(),0}
                    }
                    } 
                  if ErrorIn(try_9) {try_8 = try_9
                  } else {
                  v_bag_arg = ANY(try_9)
                  ToList(OBJ(try_8)).AddFast(v_bag_arg)
                  { var _CL_obj *ClaireVariable = To_Variable(new(ClaireVariable).Is(C_Variable))
                    _CL_obj.Pname = Core.F_gensym_void()
                    _CL_obj.Range = F_safeRange_relation(R)
                    v_bag_arg = _CL_obj.Id()
                    } 
                  ToList(OBJ(try_8)).AddFast(v_bag_arg)}
                  } 
                if ErrorIn(try_8) {Result = try_8
                } else {
                v_bag_arg = ANY(try_8)
                ToTuple(OBJ(Result)).AddFast(v_bag_arg)}
                } 
              } 
            }
            } 
          }
          } 
        } 
      } else {
      var g0289I *ClaireBoolean
      if (c.Isa.IsIn(C_Call) == CTRUE) { 
        { var g0285 *Call = To_Call(c)
          g0289I = MakeBoolean((g0285.Selector.Id() == C_add.Id()) && (g0285.Args.At(0).Isa.IsIn(C_relation) == CTRUE))
          } 
        } else {
        g0289I = CFALSE
        } 
      if (g0289I == CTRUE) { 
        { var R *ClaireRelation = ToRelation(ToList(OBJ(Core.F_CALL(C_args,ARGS(c.ToEID())))).At(0))
          { var x *ClaireVariable
            var try_11 EID
            { var _CL_obj *ClaireVariable = To_Variable(new(ClaireVariable).Is(C_Variable))
              { 
                var va_arg1 *ClaireVariable
                var va_arg2 *ClaireSymbol
                va_arg1 = _CL_obj
                var try_12 EID
                try_12 = F_extract_symbol_any(ToList(OBJ(Core.F_CALL(C_args,ARGS(c.ToEID())))).At(1))
                if ErrorIn(try_12) {try_11 = try_12
                } else {
                va_arg2 = ToSymbol(OBJ(try_12))
                va_arg1.Pname = va_arg2
                try_11 = EID{va_arg2.Id(),0}
                }
                } 
              if !ErrorIn(try_11) {
              _CL_obj.Range = R.Domain
              try_11 = EID{_CL_obj.Id(),0}
              }
              } 
            if ErrorIn(try_11) {Result = try_11
            } else {
            x = To_Variable(OBJ(try_11))
            { var y *ClaireVariable
              var try_13 EID
              { var _CL_obj *ClaireVariable = To_Variable(new(ClaireVariable).Is(C_Variable))
                { 
                  var va_arg1 *ClaireVariable
                  var va_arg2 *ClaireSymbol
                  va_arg1 = _CL_obj
                  var try_14 EID
                  try_14 = F_extract_symbol_any(ToList(OBJ(Core.F_CALL(C_args,ARGS(c.ToEID())))).At(2))
                  if ErrorIn(try_14) {try_13 = try_14
                  } else {
                  va_arg2 = ToSymbol(OBJ(try_14))
                  va_arg1.Pname = va_arg2
                  try_13 = EID{va_arg2.Id(),0}
                  }
                  } 
                if !ErrorIn(try_13) {
                _CL_obj.Range = R.Range
                try_13 = EID{_CL_obj.Id(),0}
                }
                } 
              if ErrorIn(try_13) {Result = try_13
              } else {
              y = To_Variable(OBJ(try_13))
              Result = EID{MakeTuple(R.Id(),MakeConstantList(x.Id(),y.Id()).Id()).Id(),0}
              }
              } 
            }
            } 
          } 
        } else {
        var g0290I *ClaireBoolean
        if (c.Isa.IsIn(C_Call) == CTRUE) { 
          { var g0286 *Call = To_Call(c)
            g0290I = Equal(MakeInteger(g0286.Args.Length()).Id(),MakeInteger(2).Id())
            } 
          } else {
          g0290I = CFALSE
          } 
        if (g0290I == CTRUE) { 
          { var R *ClaireProperty = To_Call(c).Selector
            { var x *ClaireVariable
              var try_15 EID
              { var _CL_obj *ClaireVariable = To_Variable(new(ClaireVariable).Is(C_Variable))
                { 
                  var va_arg1 *ClaireVariable
                  var va_arg2 *ClaireSymbol
                  va_arg1 = _CL_obj
                  var try_16 EID
                  try_16 = F_extract_symbol_any(ToList(OBJ(Core.F_CALL(C_args,ARGS(c.ToEID())))).At(0))
                  if ErrorIn(try_16) {try_15 = try_16
                  } else {
                  va_arg2 = ToSymbol(OBJ(try_16))
                  va_arg1.Pname = va_arg2
                  try_15 = EID{va_arg2.Id(),0}
                  }
                  } 
                if !ErrorIn(try_15) {
                _CL_obj.Range = R.Domain
                try_15 = EID{_CL_obj.Id(),0}
                }
                } 
              if ErrorIn(try_15) {Result = try_15
              } else {
              x = To_Variable(OBJ(try_15))
              { var y *ClaireVariable
                var try_17 EID
                { var _CL_obj *ClaireVariable = To_Variable(new(ClaireVariable).Is(C_Variable))
                  { 
                    var va_arg1 *ClaireVariable
                    var va_arg2 *ClaireSymbol
                    va_arg1 = _CL_obj
                    var try_18 EID
                    try_18 = F_extract_symbol_any(ToList(OBJ(Core.F_CALL(C_args,ARGS(c.ToEID())))).At(1))
                    if ErrorIn(try_18) {try_17 = try_18
                    } else {
                    va_arg2 = ToSymbol(OBJ(try_18))
                    va_arg1.Pname = va_arg2
                    try_17 = EID{va_arg2.Id(),0}
                    }
                    } 
                  if !ErrorIn(try_17) {
                  _CL_obj.Range = R.Range
                  try_17 = EID{_CL_obj.Id(),0}
                  }
                  } 
                if ErrorIn(try_17) {Result = try_17
                } else {
                y = To_Variable(OBJ(try_17))
                Result = EID{MakeTuple(R.Id(),MakeConstantList(x.Id(),y.Id()).Id()).Id(),0}
                }
                } 
              }
              } 
            } 
          } else {
          Result = ToException(Core.C_general_error.Make(MakeString("[188] wrong event filter: ~S").Id(),MakeConstantList(c).Id())).Close()
          } 
        } 
      } 
    } 
  return Result} 

// The EID go function for: make_filter @ any (throw: true) 
func E_make_filter_any (cond EID) EID { 
  return F_make_filter_any(ANY(cond) )} 

// create a demon with lvar as list of variables
// notice that a demon may have 3 args if R is monovalued 
/* The go function for: make_demon(R:relation,n:symbol,lvar:list[Variable],cond:any,conc:any) [status=1] */
func F_make_demon_relation (R *ClaireRelation,n *ClaireSymbol,lvar *ClaireList,cond *ClaireAny,conc *ClaireAny) EID { 
  var Result EID
  { var x *ClaireVariable = To_Variable(lvar.At(0))
    { var y *ClaireVariable = To_Variable(lvar.At(1))
      { var _Ztest *ClaireAny
        { var _CL_obj *Call = To_Call(new(Call).Is(C_Call))
          _CL_obj.Selector = ToProperty(IfThenElse((R.Multivalued_ask == CTRUE),
            C__Z.Id(),
            C__equal.Id()))
          _CL_obj.Args = MakeConstantList(y.Id(),F_readCall_relation(R,x.Id()).Id())
          _Ztest = _CL_obj.Id()
          } 
        { var _Zbody *ClaireAny = conc
          
          if (C_if_write.Trace_I > ClEnv.Verbose) { 
            { var _CL_obj *Do = To_Do(new(Do).Is(C_Do))
              { 
                var va_arg1 *Do
                var va_arg2 *ClaireList
                va_arg1 = _CL_obj
                { 
                  var v_bag_arg *ClaireAny
                  va_arg2= ToType(CEMPTY.Id()).EmptyList()
                  { var _CL_obj *Call = To_Call(new(Call).Is(C_Call))
                    _CL_obj.Selector = Core.C_format
                    { 
                      var va_arg1 *Call
                      var va_arg2 *ClaireList
                      va_arg1 = _CL_obj
                      { 
                        var v_bag_arg *ClaireAny
                        va_arg2= ToType(CEMPTY.Id()).EmptyList()
                        va_arg2.AddFast(MakeString("--- trigger ~A(~S,~S)\n").Id())
                        { var _CL_obj *List = To_List(new(List).Is(C_List))
                          _CL_obj.Args = MakeConstantList((n.String_I()).Id(),x.Id(),y.Id())
                          v_bag_arg = _CL_obj.Id()
                          } 
                        va_arg2.AddFast(v_bag_arg)} 
                      va_arg1.Args = va_arg2
                      } 
                    v_bag_arg = _CL_obj.Id()
                    } 
                  va_arg2.AddFast(v_bag_arg)
                  va_arg2.AddFast(conc)} 
                va_arg1.Args = va_arg2
                } 
              conc = _CL_obj.Id()
              } 
            } 
          { var _CL_obj *If = To_If(new(If).Is(C_If))
            _CL_obj.Arg = conc
            _Zbody = _CL_obj.Id()
            } 
          if (F_eventMethod_ask_relation2(R) == CTRUE) { 
            if (cond.Isa.IsIn(C_And) == CTRUE) { 
              { var g0291 *And = To_And(cond)
                var try_1 EID
                if (g0291.Args.Length() > 2) { 
                  { var _CL_obj *And = To_And(new(And).Is(C_And))
                    { 
                      var va_arg1 *And
                      var va_arg2 *ClaireList
                      va_arg1 = _CL_obj
                      var try_2 EID
                      try_2 = g0291.Args.Cdr()
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
                  } else {
                  try_1 = g0291.Args.At(1).ToEID()
                  } 
                if ErrorIn(try_1) {Result = try_1
                } else {
                _Ztest = ANY(try_1)
                Result = _Ztest.ToEID()
                }
                } 
              } else {
              _Zbody = conc
              Result = _Zbody.ToEID()
              } 
            }  else if (cond.Isa.IsIn(C_And) == CTRUE) { 
            { var g0293 *And = To_And(cond)
              var try_3 EID
              { var _CL_obj *And = To_And(new(And).Is(C_And))
                { 
                  var va_arg1 *And
                  var va_arg2 *ClaireList
                  va_arg1 = _CL_obj
                  var try_4 EID
                  { var arg_5 *ClaireList
                    var try_6 EID
                    try_6 = g0293.Args.Cdr()
                    if ErrorIn(try_6) {try_4 = try_6
                    } else {
                    arg_5 = ToList(OBJ(try_6))
                    try_4 = EID{MakeConstantList(_Ztest).Append(arg_5).Id(),0}
                    }
                    } 
                  if ErrorIn(try_4) {try_3 = try_4
                  } else {
                  va_arg2 = ToList(OBJ(try_4))
                  va_arg1.Args = va_arg2
                  try_3 = EID{va_arg2.Id(),0}
                  }
                  } 
                if !ErrorIn(try_3) {
                try_3 = EID{_CL_obj.Id(),0}
                }
                } 
              if ErrorIn(try_3) {Result = try_3
              } else {
              _Ztest = ANY(try_3)
              Result = _Ztest.ToEID()
              }
              } 
            } else {
            Result = EID{CFALSE.Id(),0}
            } 
          if !ErrorIn(Result) {
          if (_Zbody.Isa.IsIn(C_If) == CTRUE) { 
            { var g0294 *If = To_If(_Zbody)
              g0294.Test = _Ztest
              } 
            } 
          { var _CL_obj *LanguageDemon = ToLanguageDemon(new(LanguageDemon).Is(C_Language_demon))
            _CL_obj.Pname = n
            { 
              var va_arg1 *LanguageDemon
              var va_arg2 *ClaireLambda
              va_arg1 = _CL_obj
              var try_7 EID
              try_7 = F_lambda_I_list(lvar,_Zbody)
              if ErrorIn(try_7) {Result = try_7
              } else {
              va_arg2 = ToLambda(OBJ(try_7))
              va_arg1.Formula = va_arg2
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
  return Result} 

// The EID go function for: make_demon @ relation (throw: true) 
func E_make_demon_relation (R EID,n EID,lvar EID,cond EID,conc EID) EID { 
  return F_make_demon_relation(ToRelation(OBJ(R)),
    ToSymbol(OBJ(n)),
    ToList(OBJ(lvar)),
    ANY(cond),
    ANY(conc) )} 

// cute litle guy : create the read instruction both for a table and a property
/* The go function for: readCall(R:relation,x:any) [status=0] */
func F_readCall_relation (R *ClaireRelation,x *ClaireAny) *Call { 
  var Result *Call
  if (C_table.Id() == R.Isa.Id()) { 
    { var _CL_obj *Call = To_Call(new(Call).Is(C_Call))
      _CL_obj.Selector = C_get
      _CL_obj.Args = MakeConstantList(R.Id(),x)
      Result = _CL_obj
      } 
    } else {
    { var _CL_obj *Call_plus = To_Call_plus(new(Call_plus).Is(C_Call_plus))
      _CL_obj.Selector = ToProperty(R.Id())
      _CL_obj.Args = MakeConstantList(x)
      Result = To_Call(_CL_obj.Id())
      } 
    } 
  return Result} 

// The EID go function for: readCall @ relation (throw: false) 
func E_readCall_relation (R EID,x EID) EID { 
  return EID{F_readCall_relation(ToRelation(OBJ(R)),ANY(x) ).Id(),0}} 

// a small brother
/* The go function for: putCall(R:relation,x:any,y:any) [status=0] */
func F_putCall_relation2 (R *ClaireRelation,x *ClaireAny,y *ClaireAny) *Call { 
  var Result *Call
  if (R.Multivalued_ask == CTRUE) { 
    { var _CL_obj *Call = To_Call(new(Call).Is(C_Call))
      _CL_obj.Selector = Core.C_add_value
      _CL_obj.Args = MakeConstantList(R.Id(),x,y)
      Result = _CL_obj
      } 
    } else {
    { var _CL_obj *Call = To_Call(new(Call).Is(C_Call))
      _CL_obj.Selector = C_put
      _CL_obj.Args = MakeConstantList(R.Id(),x,y)
      Result = _CL_obj
      } 
    } 
  return Result} 

// The EID go function for: putCall @ relation (throw: false) 
func E_putCall_relation2 (R EID,x EID,y EID) EID { 
  return EID{F_putCall_relation2(ToRelation(OBJ(R)),ANY(x),ANY(y) ).Id(),0}} 

// v3.3 : find the range when we read the current value     
/* The go function for: safeRange(x:relation) [status=0] */
func F_safeRange_relation (x *ClaireRelation) *ClaireType { 
  var Result *ClaireType
  if (x.Isa.IsIn(C_property) == CTRUE) { 
    { var g0295 *ClaireProperty = ToProperty(x.Id())
      var g0299I *ClaireBoolean
      { var arg_1 *ClaireAny
        { 
          var s *ClaireRestriction
          _ = s
          var s_iter *ClaireAny
          arg_1= CFALSE.Id()
          for _,s_iter = range(g0295.Restrictions.ValuesO()){ 
            s = ToRestriction(s_iter)
            var g0300I *ClaireBoolean
            { var arg_2 *ClaireBoolean
              if (C_slot.Id() == s.Isa.Id()) { 
                { var g0296 *ClaireSlot = ToSlot(s.Id())
                  arg_2 = g0296.Range.Contains(g0296.Default)
                  } 
                } else {
                arg_2 = CFALSE
                } 
              g0300I = arg_2.Not
              } 
            if (g0300I == CTRUE) { 
              arg_1 = CTRUE.Id()
              break
              } 
            } 
          } 
        g0299I = Core.F_not_any(arg_1)
        } 
      if (g0299I == CTRUE) { 
        Result = g0295.Range
        } else {
        Result = ToType(C_any.Id())
        } 
      } 
    }  else if (C_table.Id() == x.Isa.Id()) { 
    { var g0297 *ClaireTable = ToTable(x.Id())
      if (g0297.Range.Contains(g0297.Default) == CTRUE) { 
        Result = g0297.Range
        } else {
        Result = ToType(C_any.Id())
        } 
      } 
    } else {
    Result = ToType(C_any.Id())
    } 
  return Result} 

// The EID go function for: safeRange @ relation (throw: false) 
func E_safeRange_relation (x EID) EID { 
  return EID{F_safeRange_relation(ToRelation(OBJ(x)) ).Id(),0}} 

// generate an if_write "daemon", only the first time, which uses
// the list in demons[R]
// the first step is to make the update (with inverse management)
/* The go function for: eval_if_write(R:relation) [status=1] */
func F_eval_if_write_relation (R *ClaireRelation) EID { 
  var Result EID
  { var l *ClaireList = ToList(Core.F_get_table(C_demons,R.Id()))
    { var lvar *ClaireList = ToLanguageDemon(l.ValuesO()[0]).Formula.Vars
      { var dv *ClaireVariable
        { var _CL_obj *ClaireVariable = To_Variable(new(ClaireVariable).Is(C_Variable))
          _CL_obj.Pname = Core.F_gensym_void()
          _CL_obj.Range = ToType(C_Language_demon.Id())
          dv = _CL_obj
          } 
        { var l1 *ClaireList = MakeList(ToType(C_any.Id()),F_putCall_relation2(R,lvar.At(0),lvar.At(1)).Id())
          { var l2 *ClaireList
            { 
              var v_bag_arg *ClaireAny
              l2= ToType(C_any.Id()).EmptyList()
              { var _CL_obj *For = To_For(new(For).Is(C_For))
                _CL_obj.ClaireVar = dv
                { 
                  var va_arg1 *Iteration
                  var va_arg2 *ClaireAny
                  va_arg1 = To_Iteration(_CL_obj.Id())
                  { var _CL_obj *Call = To_Call(new(Call).Is(C_Call))
                    _CL_obj.Selector = C_nth
                    _CL_obj.Args = MakeConstantList(C_demons.Id(),R.Id())
                    va_arg2 = _CL_obj.Id()
                    } 
                  va_arg1.SetArg = va_arg2
                  } 
                { 
                  var va_arg1 *Iteration
                  var va_arg2 *ClaireAny
                  va_arg1 = To_Iteration(_CL_obj.Id())
                  { var _CL_obj *Call = To_Call(new(Call).Is(C_Call))
                    _CL_obj.Selector = C_funcall
                    _CL_obj.Args = MakeConstantList(dv.Id()).Append(lvar)
                    va_arg2 = _CL_obj.Id()
                    } 
                  va_arg1.Arg = va_arg2
                  } 
                v_bag_arg = _CL_obj.Id()
                } 
              l2.AddFast(v_bag_arg)} 
            
            { 
              var v *ClaireAny
              _ = v
              Result= EID{CFALSE.Id(),0}
              var v_support *ClaireList
              v_support = lvar
              v_len := v_support.Length()
              for i_it := 0; i_it < v_len; i_it++ { 
                v = v_support.At(i_it)
                var loop_1 EID
                _ = loop_1
                loop_1 = Core.F_put_property2(C_range,ToObject(v),ToTypeExpression(OBJ(Core.F_CALL(C_range,ARGS(v.ToEID())))).Class_I().Id())
                if ErrorIn(loop_1) {Result = loop_1
                break
                } else {
                }
                } 
              } 
            if !ErrorIn(Result) {
            if (R.Inverse.Id() != CNULL) { 
              if (R.Multivalued_ask != CTRUE) { 
                { var arg_2 *Call
                  { var _CL_obj *Call = To_Call(new(Call).Is(C_Call))
                    _CL_obj.Selector = Core.C_Core_update_dash
                    _CL_obj.Args = MakeConstantList(R.Inverse.Id(),lvar.At(2),lvar.At(0))
                    arg_2 = _CL_obj
                    } 
                  l1 = l1.AddFast(arg_2.Id())
                  } 
                } 
              l1 = l1.AddFast(F_putCall_relation2(R.Inverse,lvar.At(1),lvar.At(0)).Id())
              } 
            { 
              var va_arg1 *ClaireRelation
              var va_arg2 *ClaireAny
              va_arg1 = R
              var try_3 EID
              { var arg_4 *ComplexInstruction
                if (F_eventMethod_ask_relation2(R) == CTRUE) { 
                  { var _CL_obj *Do = To_Do(new(Do).Is(C_Do))
                    _CL_obj.Args = l2
                    arg_4 = To_ComplexInstruction(_CL_obj.Id())
                    } 
                  }  else if (R.Multivalued_ask == CTRUE) { 
                  { var _CL_obj *If = To_If(new(If).Is(C_If))
                    { 
                      var va_arg1 *If
                      var va_arg2 *ClaireAny
                      va_arg1 = _CL_obj
                      { var _CL_obj *Call = To_Call(new(Call).Is(C_Call))
                        _CL_obj.Selector = Core.C_not
                        { 
                          var va_arg1 *Call
                          var va_arg2 *ClaireList
                          va_arg1 = _CL_obj
                          { 
                            var v_bag_arg *ClaireAny
                            va_arg2= ToType(CEMPTY.Id()).EmptyList()
                            { var _CL_obj *Call = To_Call(new(Call).Is(C_Call))
                              _CL_obj.Selector = ToProperty(C__Z.Id())
                              _CL_obj.Args = MakeConstantList(lvar.At(1),F_readCall_relation(R,lvar.At(0)).Id())
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
                      var va_arg1 *If
                      var va_arg2 *ClaireAny
                      va_arg1 = _CL_obj
                      { var _CL_obj *Do = To_Do(new(Do).Is(C_Do))
                        _CL_obj.Args = l1.Append(l2)
                        va_arg2 = _CL_obj.Id()
                        } 
                      va_arg1.Arg = va_arg2
                      } 
                    arg_4 = To_ComplexInstruction(_CL_obj.Id())
                    } 
                  } else {
                  { var _CL_obj *Let = To_Let(new(Let).Is(C_Let))
                    _CL_obj.ClaireVar = To_Variable(lvar.At(2))
                    _CL_obj.Value = F_readCall_relation(R,lvar.At(0)).Id()
                    { 
                      var va_arg1 *Let
                      var va_arg2 *ClaireAny
                      va_arg1 = _CL_obj
                      { var _CL_obj *If = To_If(new(If).Is(C_If))
                        { 
                          var va_arg1 *If
                          var va_arg2 *ClaireAny
                          va_arg1 = _CL_obj
                          { var _CL_obj *Call = To_Call(new(Call).Is(C_Call))
                            _CL_obj.Selector = ToProperty(Core.C__I_equal.Id())
                            _CL_obj.Args = MakeConstantList(lvar.At(1),lvar.At(2))
                            va_arg2 = _CL_obj.Id()
                            } 
                          va_arg1.Test = va_arg2
                          } 
                        { 
                          var va_arg1 *If
                          var va_arg2 *ClaireAny
                          va_arg1 = _CL_obj
                          { var _CL_obj *Do = To_Do(new(Do).Is(C_Do))
                            _CL_obj.Args = l1.Append(l2)
                            va_arg2 = _CL_obj.Id()
                            } 
                          va_arg1.Arg = va_arg2
                          } 
                        va_arg2 = _CL_obj.Id()
                        } 
                      va_arg1.Arg = va_arg2
                      } 
                    arg_4 = To_ComplexInstruction(_CL_obj.Id())
                    } 
                  } 
                try_3 = F_lambda_I_list(MakeConstantList(lvar.At(0),lvar.At(1)),arg_4.Id())
                } 
              if ErrorIn(try_3) {Result = try_3
              } else {
              va_arg2 = ANY(try_3)
              va_arg1.IfWrite = va_arg2
              Result = va_arg2.ToEID()
              }
              } 
            }
            } 
          } 
        } 
      } 
    } 
  return Result} 

// The EID go function for: eval_if_write @ relation (throw: true) 
func E_eval_if_write_relation (R EID) EID { 
  return F_eval_if_write_relation(ToRelation(OBJ(R)) )} 

// create a restriction (method) that will trigger an event
/* The go function for: eventMethod(p:property) [status=0] */
func F_eventMethod_property (p *ClaireProperty)  { 
  { var m *ClaireMethod = F_add_method_property(p,MakeConstantList(p.Domain.Id(),p.Range.Id()),ToType(C_void.Id()),0,ToFunction(CNULL))
    { var _Zf *ClaireFunction = F_make_function_string(F_append_string(p.Name.String_I(),MakeString("_write")))
      m.Formula = ToLambda(p.IfWrite)
      Core.F_close_method(m)
      F_set_arity_function(_Zf,2)
      m.Functional = _Zf
      } 
    } 
  } 

// The EID go function for: eventMethod @ property (throw: false) 
func E_eventMethod_property (p EID) EID { 
  F_eventMethod_property(ToProperty(OBJ(p)) )
  return EVOID} 

// when we compile -> directly call the demon 
// **************************************************************************
// *     Part 5: JITO for methods                                           *
// **************************************************************************
// CLAIRE 4 reintroduced JITO : Just-In-Time Optimization
// we perform an on-the-fly optimization of lambdas through substitution (static calls)
// Jito(l:lambda) -> apply makeJito to the body (in place substitution)
/* The go function for: jito(self:any) [status=1] */
func F_Language_jito_any (self *ClaireAny) EID { 
  var Result EID
  if (ClEnv.Jito_ask != CTRUE) { 
    Result = self.ToEID()
    }  else if (self.Isa.IsIn(C_list) == CTRUE) { 
    { var g0301 *ClaireList = ToList(self)
      { 
        var x *ClaireAny
        _ = x
        Result= EID{CFALSE.Id(),0}
        var x_support *ClaireList
        x_support = g0301
        x_len := x_support.Length()
        for i_it := 0; i_it < x_len; i_it++ { 
          x = x_support.At(i_it)
          var loop_1 EID
          _ = loop_1
          loop_1 = F_Language_jito_any(x)
          if ErrorIn(loop_1) {Result = loop_1
          break
          } else {
          }
          } 
        } 
      } 
    }  else if (self.Isa.IsIn(C_Vardef) == CTRUE) { 
    { var g0302 *Vardef = To_Vardef(self)
      g0302.Isa = C_Variable
      Result = EID{g0302.Id(),0}
      } 
    }  else if (self.Isa.IsIn(C_lambda) == CTRUE) { 
    { var g0303 *ClaireLambda = ToLambda(self)
      Result = F_Language_jito_any(g0303.Body)
      if !ErrorIn(Result) {
      Result = EID{g0303.Id(),0}
      }
      } 
    }  else if (self.Isa.IsIn(C_And) == CTRUE) { 
    { var g0304 *And = To_And(self)
      Result = F_Language_jito_any(g0304.Args.Id())
      } 
    }  else if (self.Isa.IsIn(C_Or) == CTRUE) { 
    { var g0305 *Or = To_Or(self)
      Result = F_Language_jito_any(g0305.Args.Id())
      } 
    }  else if (self.Isa.IsIn(C_Call) == CTRUE) { 
    { var g0306 *Call = To_Call(self)
      Result = g0306.MakeJito()
      if !ErrorIn(Result) {
      Result = EID{CTRUE.Id(),0}
      }
      } 
    }  else if (self.Isa.IsIn(C_Let) == CTRUE) { 
    { var g0307 *Let = To_Let(self)
      Result = g0307.LetJito()
      } 
    }  else if (self.Isa.IsIn(C_Assign) == CTRUE) { 
    { var g0308 *Assign = To_Assign(self)
      if (g0308.ClaireVar.Isa.IsIn(C_Variable) != CTRUE) { 
        Result = ToException(Core.C_general_error.Make(MakeString("[101] ~S is not a variable but a ~S").Id(),MakeConstantList(g0308.ClaireVar,g0308.ClaireVar.Isa.Id()).Id())).Close()
        } else {
        Result = EID{CFALSE.Id(),0}
        } 
      if !ErrorIn(Result) {
      Result = F_Language_jito_any(g0308.Arg)
      }
      } 
    }  else if (self.Isa.IsIn(C_Gassign) == CTRUE) { 
    { var g0309 *Gassign = To_Gassign(self)
      if (g0309.ClaireVar.Range.Contains(g0309.ClaireVar.Value) == CTRUE) { 
        Result = F_Language_jito_any(g0309.Arg)
        } else {
        Result = EID{CFALSE.Id(),0}
        } 
      } 
    }  else if (self.Isa.IsIn(C_Do) == CTRUE) { 
    { var g0310 *Do = To_Do(self)
      Result = F_Language_jito_any(g0310.Args.Id())
      } 
    }  else if (self.Isa.IsIn(C_If) == CTRUE) { 
    { var g0311 *If = To_If(self)
      Result = F_Language_jito_any(g0311.Arg)
      if !ErrorIn(Result) {
      Result = F_Language_jito_any(g0311.Test)
      if !ErrorIn(Result) {
      Result = F_Language_jito_any(g0311.Other)
      }}
      } 
    }  else if (self.Isa.IsIn(C_Iteration) == CTRUE) { 
    { var g0312 *Iteration = To_Iteration(self)
      { var v *ClaireVariable = g0312.ClaireVar
        { var s *ClaireAny = g0312.SetArg
          { var o_ask *ClaireBoolean
            { 
              var v_and6 *ClaireBoolean
              
              if (s.Isa.IsIn(C_Call) == CTRUE) { 
                { var g0313 *Call = To_Call(s)
                  v_and6 = Equal(g0313.Selector.Id(),C__dot_dot.Id())
                  } 
                } else {
                v_and6 = CFALSE
                } 
              if (v_and6 == CFALSE) {o_ask = CFALSE
              } else { 
                v_and6 = MakeBoolean((v.Range.Id() == CNULL))
                if (v_and6 == CFALSE) {o_ask = CFALSE
                } else { 
                  o_ask = CTRUE} 
                } 
              } 
            { var arg_2 *ClaireList
              var try_3 EID
              { 
                var v_bag_arg *ClaireAny
                try_3= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                ToList(OBJ(try_3)).AddFast(g0312.Id())
                var try_4 EID
                try_4 = F_static_type_any(g0312.SetArg)
                if ErrorIn(try_4) {try_3 = try_4
                } else {
                v_bag_arg = ANY(try_4)
                ToList(OBJ(try_3)).AddFast(v_bag_arg)}
                } 
              if ErrorIn(try_3) {Result = try_3
              } else {
              arg_2 = ToList(OBJ(try_3))
              Result = Core.F_tformat_string(MakeString("-- Iteration jito: ~S (~S)\n"),3,arg_2)
              }
              } 
            if !ErrorIn(Result) {
            if (o_ask == CTRUE) { 
              v.Range = ToType(C_integer.Id())
              Core.F_tformat_string(MakeString("-- jito:put range ~S as integer\n"),3,MakeConstantList(v.Id()))
              } 
            Result = F_Language_jito_any(s)
            if !ErrorIn(Result) {
            Result = F_Language_jito_any(g0312.Arg)
            if !ErrorIn(Result) {
            if (o_ask == CTRUE) { 
              { 
                var va_arg1 *ClaireVariable
                var va_arg2 *ClaireType
                va_arg1 = v
                va_arg2 = ToType(CNULL)
                va_arg1.Range = va_arg2
                Result = EID{va_arg2.Id(),0}
                } 
              } else {
              Result = EID{CFALSE.Id(),0}
              } 
            }}}
            } 
          } 
        } 
      } 
    }  else if (self.Isa.IsIn(C_While) == CTRUE) { 
    { var g0314 *While = To_While(self)
      Result = F_Language_jito_any(g0314.Test)
      if !ErrorIn(Result) {
      Result = F_Language_jito_any(g0314.Arg)
      }
      } 
    }  else if (self.Isa.IsIn(C_Construct) == CTRUE) { 
    { var g0315 *Construct = To_Construct(self)
      Core.F_tformat_string(MakeString("-- Construct jito: ~S\n"),3,MakeConstantList(g0315.Id()))
      Result = F_Language_jito_any(g0315.Args.Id())
      } 
    }  else if (self.Isa.IsIn(C_Exists) == CTRUE) { 
    { var g0316 *Exists = To_Exists(self)
      Result = F_Language_jito_any(g0316.SetArg)
      if !ErrorIn(Result) {
      Result = F_Language_jito_any(g0316.Arg)
      if !ErrorIn(Result) {
      Result = F_Language_jito_any(g0316.Other)
      }}
      } 
    }  else if (self.Isa.IsIn(C_Handle) == CTRUE) { 
    { var g0317 *ClaireHandle = To_ClaireHandle(self)
      if (C_class.Id() != g0317.Test.Isa.Id()) { 
        Result = ToException(Core.C_general_error.Make(MakeString("syntax: [try %S] must use a class").Id(),MakeConstantList(g0317.Test).Id())).Close()
        } else {
        Result = EID{CFALSE.Id(),0}
        } 
      if !ErrorIn(Result) {
      Result = F_Language_jito_any(g0317.Arg)
      if !ErrorIn(Result) {
      Result = F_Language_jito_any(g0317.Other)
      }}
      } 
    }  else if (self.Isa.IsIn(C_Definition) == CTRUE) { 
    { var g0318 *Definition = To_Definition(self)
      if (F_Language_fast_definition_ask_class(g0318.Arg) == CTRUE) { 
        { 
          var va_arg1 *ClaireAny
          var va_arg2 *ClaireClass
          va_arg1 = g0318.Id()
          va_arg2 = C_Language_DefFast
          va_arg1.Isa = va_arg2
          Result = EID{va_arg2.Id(),0}
          } 
        } else {
        Result = EID{CFALSE.Id(),0}
        } 
      } 
    } else {
    Result = EID{CFALSE.Id(),0}
    } 
  return Result} 

// The EID go function for: jito @ any (throw: true) 
func E_Language_jito_any (self EID) EID { 
  return F_Language_jito_any(ANY(self) )} 

// Let is special in CLAIRE4 : we implement the implicit typing found in the compiler = to infer
// the type  from the value (when no range is given)
// Note : this is doubtful ... 
/* The go function for: letJito(self:Let) [status=1] */
func (self *Let) LetJito () EID { 
  var Result EID
  { var v *ClaireVariable = self.ClaireVar
    { var x *ClaireAny = self.Value
      { var untyped *ClaireBoolean = MakeBoolean((v.Range.Id() == CNULL))
        Core.F_tformat_string(MakeString("Let Jito with var ~S => ~S\n"),3,MakeConstantList(v.Id(),untyped.Id()))
        if (untyped == CTRUE) { 
          if (x.Isa.IsIn(C_List) == CTRUE) { 
            { var t *ClaireType = ToType(OBJ(Core.F_CALL(C_of,ARGS(x.ToEID()))))
              if (Equal(t.Id(),CEMPTY.Id()) != CTRUE) { 
                { 
                  var va_arg1 *ClaireVariable
                  var va_arg2 *ClaireType
                  va_arg1 = v
                  va_arg2 = Core.F_param_I_class(C_list,t)
                  va_arg1.Range = va_arg2
                  Result = EID{va_arg2.Id(),0}
                  } 
                } else {
                Result = EID{CFALSE.Id(),0}
                } 
              } 
            } else {
            { 
              var va_arg1 *ClaireVariable
              var va_arg2 *ClaireType
              va_arg1 = v
              var try_1 EID
              try_1 = F_static_type_any(x)
              if ErrorIn(try_1) {Result = try_1
              } else {
              va_arg2 = ToType(OBJ(try_1))
              va_arg1.Range = va_arg2
              Result = EID{va_arg2.Id(),0}
              }
              } 
            } 
          if !ErrorIn(Result) {
          Result = Core.F_tformat_string(MakeString("--- let Jito ~S:~S (~S)\n"),3,MakeConstantList(v.Id(),v.Range.Id(),x))
          }
          } else {
          Result = EID{CFALSE.Id(),0}
          } 
        if !ErrorIn(Result) {
        Result = F_Language_jito_any(x)
        if !ErrorIn(Result) {
        Result = F_Language_jito_any(self.Arg)
        if !ErrorIn(Result) {
        if (untyped == CTRUE) { 
          { 
            var va_arg1 *ClaireVariable
            var va_arg2 *ClaireType
            va_arg1 = v
            va_arg2 = ToType(CNULL)
            va_arg1.Range = va_arg2
            Result = EID{va_arg2.Id(),0}
            } 
          } else {
          Result = EID{CFALSE.Id(),0}
          } 
        }}}
        } 
      } 
    } 
  return Result} 

// The EID go function for: letJito @ Let (throw: true) 
func E_Language_letJito_Let (self EID) EID { 
  return To_Let(OBJ(self)).LetJito( )} 

// we optimize statically (Call(p) -> Call_method(m)) when
//   - only one restriction match 
//   - all domains are classes => class match
//   - the only one match is a compiled method
//   - the property is static (open = 1, vs extensible) and not too many restrictions
// note: the 12 hard limit is to avoid spending too much time with self_print or equivalent methods ... it is arbitrary
/* The go function for: makeJito(self:Call) [status=1] */
func (self *Call) MakeJito () EID { 
  var Result EID
  Result = F_Language_jito_any(self.Args.Id())
  if !ErrorIn(Result) {
  { var p *ClaireProperty = self.Selector
    { var larg *ClaireList = self.Args
      { var n int = larg.Length()
        { var m *ClaireAny = CNULL
          var g0324I *ClaireBoolean
          { 
            var v_and5 *ClaireBoolean
            
            v_and5 = Equal(p.Id(),Core.C_write.Id())
            if (v_and5 == CFALSE) {g0324I = CFALSE
            } else { 
              { var p2 *ClaireAny = self.Args.At(0)
                if (p2.Isa.IsIn(C_property) == CTRUE) { 
                  { var g0320 *ClaireProperty = ToProperty(p2)
                    v_and5 = MakeBoolean((g0320.Inverse.Id() == CNULL) && (g0320.Store_ask != CTRUE) && (g0320.IfWrite == CNULL))
                    } 
                  } else {
                  v_and5 = CFALSE
                  } 
                } 
              if (v_and5 == CFALSE) {g0324I = CFALSE
              } else { 
                g0324I = CTRUE} 
              } 
            } 
          if (g0324I == CTRUE) { 
            p = C_write_fast
            self.Selector = C_write_fast
            } 
          var g0325I *ClaireBoolean
          { 
            var v_and5 *ClaireBoolean
            
            v_and5 = Core.F__inf_equal_integer(p.Open,1)
            if (v_and5 == CFALSE) {g0325I = CFALSE
            } else { 
              v_and5 = Core.F__inf_equal_integer(p.Restrictions.Length(),12)
              if (v_and5 == CFALSE) {g0325I = CFALSE
              } else { 
                { var arg_1 *ClaireAny
                  { 
                    var x *ClaireRestriction
                    _ = x
                    var x_iter *ClaireAny
                    arg_1= CFALSE.Id()
                    for _,x_iter = range(p.Restrictions.ValuesO()){ 
                      x = ToRestriction(x_iter)
                      var g0326I *ClaireBoolean
                      { var arg_2 *ClaireBoolean
                        { var arg_3 *ClaireAny
                          { 
                            var t *ClaireTypeExpression
                            _ = t
                            var t_iter *ClaireAny
                            arg_3= CFALSE.Id()
                            for _,t_iter = range(x.Domain.ValuesO()){ 
                              t = ToTypeExpression(t_iter)
                              if (C_class.Id() != t.Isa.Id()) { 
                                arg_3 = CTRUE.Id()
                                break
                                } 
                              } 
                            } 
                          arg_2 = Core.F_not_any(arg_3)
                          } 
                        g0326I = arg_2.Not
                        } 
                      if (g0326I == CTRUE) { 
                        arg_1 = CTRUE.Id()
                        break
                        } 
                      } 
                    } 
                  v_and5 = Core.F_not_any(arg_1)
                  } 
                if (v_and5 == CFALSE) {g0325I = CFALSE
                } else { 
                  g0325I = CTRUE} 
                } 
              } 
            } 
          if (g0325I == CTRUE) { 
            { var lt *ClaireList
              var try_4 EID
              { 
                var v_list7 *ClaireList
                var x *ClaireAny
                var v_local7 *ClaireAny
                v_list7 = larg
                try_4 = EID{CreateList(ToType(CEMPTY.Id()),v_list7.Length()).Id(),0}
                for CLcount := 0; CLcount < v_list7.Length(); CLcount++{ 
                  x = v_list7.At(CLcount)
                  var try_5 EID
                  try_5 = F_static_type_any(x)
                  if ErrorIn(try_5) {try_4 = try_5
                  break
                  } else {
                  v_local7 = ANY(try_5)
                  ToList(OBJ(try_4)).PutAt(CLcount,v_local7)
                  } 
                }
                } 
              if ErrorIn(try_4) {Result = try_4
              } else {
              lt = ToList(OBJ(try_4))
              Core.F_tformat_string(MakeString("-- call jito: ~S : ~S\n"),3,MakeConstantList(self.Id(),lt.Id()))
              { 
                var x *ClaireRestriction
                _ = x
                var x_iter *ClaireAny
                Result= EID{CFALSE.Id(),0}
                for _,x_iter = range(p.Definition.ValuesO()){ 
                  x = ToRestriction(x_iter)
                  if (F_Language_makeCallMatch_restriction(x,lt) == CTRUE) { 
                    m = x.Id()
                    Result = EID{CTRUE.Id(),0}
                    break
                    } 
                  } 
                } 
              }
              } 
            } else {
            Result = EID{CFALSE.Id(),0}
            } 
          if !ErrorIn(Result) {
          var g0327I *ClaireBoolean
          if (C_method.Id() == m.Isa.Id()) { 
            { var g0322 *ClaireMethod = ToMethod(m)
              g0327I = MakeBoolean((g0322.Functional.Id() == CNULL)).Not
              } 
            } else {
            g0327I = CFALSE
            } 
          if (g0327I == CTRUE) { 
            { 
              var va_arg1 *ClaireAny
              var va_arg2 *ClaireClass
              va_arg1 = self.Id()
              if (n == 1) { 
                va_arg2 = C_Call_method1
                }  else if (n == 2) { 
                va_arg2 = C_Call_method2
                }  else if (n == 3) { 
                va_arg2 = C_Language_Call_method3
                } else {
                va_arg2 = C_Call_method
                } 
              va_arg1.Isa = va_arg2
              } 
            { 
              var va_arg1 *CallMethod
              var va_arg2 *ClaireMethod
              va_arg1 = To_CallMethod(self.Id())
              va_arg2 = ToMethod(m)
              va_arg1.Arg = va_arg2
              Result = EID{va_arg2.Id(),0}
              } 
            } else {
            Result = EID{CFALSE.Id(),0}
            } 
          }
          } 
        } 
      } 
    } 
  }
  return Result} 

// The EID go function for: makeJito @ Call (throw: true) 
func E_Language_makeJito_Call (self EID) EID { 
  return To_Call(OBJ(self)).MakeJito( )} 

// tells if the restriction matches the type list lt : we know that the domain is made of classes
// only use for a compiled method, to help with debug
/* The go function for: makeCallMatch(x:restriction,lt:list) [status=0] */
func F_Language_makeCallMatch_restriction (x *ClaireRestriction,lt *ClaireList) *ClaireBoolean { 
  var Result *ClaireBoolean
  { var n int = lt.Length()
    { var ld *ClaireList = x.Domain
      { 
        var v_and3 *ClaireBoolean
        
        v_and3 = Equal(MakeInteger(ld.Length()).Id(),MakeInteger(n).Id())
        if (v_and3 == CFALSE) {Result = CFALSE
        } else { 
          { var arg_1 *ClaireAny
            { var i int = 1
              { var g0328 int = n
                arg_1= CFALSE.Id()
                for (i <= g0328) { 
                  if (ToType(lt.At(i-1)).Included(ToType(ld.ValuesO()[i-1])) != CTRUE) { 
                    arg_1 = CTRUE.Id()
                    break
                    } 
                  i = (i+1)
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
    } 
  return Result} 

// The EID go function for: makeCallMatch @ restriction (throw: false) 
func E_Language_makeCallMatch_restriction (x EID,lt EID) EID { 
  return EID{F_Language_makeCallMatch_restriction(ToRestriction(OBJ(x)),ToList(OBJ(lt)) ).Id(),0}} 

// close some classes : final => no subclasses,  default() => ephemeral
// CLAIRE 4 : make sure that open statement for class are all here
// instuctions are ephemeral