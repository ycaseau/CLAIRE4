/***** CLAIRE Compilation of file /Users/ycaseau/Dropbox/src/clairev4.0/src/meta/define.cl 
         [version 4.0.02 / safety 5] Saturday 10-30-2021 *****/

package Language
import (_ "fmt"
	. "Kernel"
	"Core"
)

//-------- dumb function to prevent import errors --------
func import_g0405() { 
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
/* {1} OPT.The go function for: self_print(self:Definition) [] */
func (self *Definition ) SelfPrint () EID { 
    var Result EID 
    Result = Core.F_print_any(self.Arg.Id())
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("(")
    Result = F_Language_printbox_list2(self.Args)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(")")
    Result = EVOID
    }}
    return Result} 
  
// The EID go function for: self_print @ Definition (throw: true) 
func E_self_print_Definition_Language (self EID) EID { 
    return /*(sm for self_print @ Definition= EID)*/ To_Definition(OBJ(self)).SelfPrint( )} 
  
// CLAIRE 4: fast definition when no close nor facy slots
// ------------- named object definition ------------------------------
//
/* {1} OPT.The go function for: self_print(self:Defobj) [] */
func (self *Defobj ) SelfPrint () EID { 
    var Result EID 
    if (self.Arg.Id() == Core.C_global_variable.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var r *ClaireAny   = C_any.Id()
        /* noccur = 3 */
        /* Let:4 */{ 
          var v *ClaireAny   = CNULL
          /* noccur = 3 */
          /* For:5 */{ 
            var x *ClaireAny  
            _ = x
            var x_support *ClaireList  
            x_support = self.Args
            x_len := x_support.Length()
            for i_it := 0; i_it < x_len; i_it++ { 
              x = x_support.At(i_it)
              if (To_Call(x).Args.At(1-1) == C_value.Id()) /* If:7 */{ 
                v = To_Call(x).Args.At(2-1)
                /* If!7 */}  else if (To_Call(x).Args.At(1-1) == C_range.Id()) /* If:7 */{ 
                r = To_Call(x).Args.At(2-1)
                /* If-7 */} 
              /* loop-6 */} 
            /* For-5 */} 
          if (F_boolean_I_any(r) == CTRUE) /* If:5 */{ 
            self.Ident.Princ()
            PRINC(":")
            Result = Core.F_CALL(C_print,ARGS(r.ToEID()))
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC(" := ")
            Result = F_printexp_any(v,CFALSE)
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC("")
            Result = EVOID
            }}
            } else {
            self.Ident.Princ()
            PRINC(" :: ")
            Result = F_printexp_any(v,CFALSE)
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC("")
            Result = EVOID
            }
            /* If-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      } else {
      self.Ident.Princ()
      PRINC(" :: ")
      Result = Core.F_CALL(C_print,ARGS(EID{self.Arg.Id(),0}))
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC("(")
      Result = F_Language_printbox_list2(self.Args)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(")")
      Result = EVOID
      }}
      /* If-2 */} 
    return Result} 
  
// The EID go function for: self_print @ Defobj (throw: true) 
func E_self_print_Defobj_Language (self EID) EID { 
    return /*(sm for self_print @ Defobj= EID)*/ To_Defobj(OBJ(self)).SelfPrint( )} 
  
// ------------- class definition ------------------------------------
//
/* {1} OPT.The go function for: self_print(self:Defclass) [] */
func (self *Defclass ) SelfPrint () EID { 
    var Result EID 
    if (self.Ident.Id() == CNULL) /* If:2 */{ 
      Result = Core.F_print_any(MakeString("<Defclass>").Id())
      } else {
      self.Ident.Princ()
      if (self.Params.Length() != 0) /* If:3 */{ 
        PRINC("[")
        Result = Core.F_princ_list(self.Params)
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC("]")
        Result = EVOID
        }
        } else {
        Result = EID{CFALSE.Id(),0}
        /* If-3 */} 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(" <: ")
      Result = Core.F_print_any(self.Arg.Id())
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC("(")
      /* Let:3 */{ 
        var l *ClaireList   = self.Args
        /* noccur = 5 */
        /* Let:4 */{ 
          var n int  = l.Length()
          /* noccur = 2 */
          /* Let:5 */{ 
            var i int  = 1
            /* noccur = 9 */
            /* Let:6 */{ 
              var g0406 int  = n
              /* noccur = 1 */
              Result= EID{CFALSE.Id(),0}
              for (i <= g0406) /* while:7 */{ 
                var void_try8 EID 
                _ = void_try8
                { 
                if (i == 1) /* If:8 */{ 
                  F_set_level_void()
                  void_try8 = EVOID
                  } else {
                  void_try8 = F_lbreak_void()
                  /* If-8 */} 
                /* ERROR PROTECTION INSERTED (void_try8-void_try8) */
                if ErrorIn(void_try8) {Result = void_try8
                break
                } else {
                if (l.At(i-1).Isa.IsIn(C_Vardef) == CTRUE) /* If:8 */{ 
                  void_try8 = Core.F_CALL(C_Language_ppvariable,ARGS(l.At(i-1).ToEID()))
                  } else {
                  void_try8 = Core.F_CALL(C_Language_ppvariable,ARGS(ToList(OBJ(Core.F_CALL(C_args,ARGS(l.At(i-1).ToEID())))).At(1-1).ToEID()))
                  /* ERROR PROTECTION INSERTED (void_try8-void_try8) */
                  if ErrorIn(void_try8) {Result = void_try8
                  break
                  } else {
                  PRINC(" = ")
                  void_try8 = Core.F_CALL(C_print,ARGS(ToList(OBJ(Core.F_CALL(C_args,ARGS(l.At(i-1).ToEID())))).At(2-1).ToEID()))
                  /* ERROR PROTECTION INSERTED (void_try8-void_try8) */
                  if ErrorIn(void_try8) {Result = void_try8
                  break
                  } else {
                  PRINC("")
                  void_try8 = EVOID
                  }}
                  /* If-8 */} 
                /* ERROR PROTECTION INSERTED (void_try8-void_try8) */
                if ErrorIn(void_try8) {Result = void_try8
                break
                } else {
                if (i < n) /* If:8 */{ 
                  PRINC(",")
                  void_try8 = EVOID
                  } else {
                  void_try8 = EID{CFALSE.Id(),0}
                  /* If-8 */} 
                }}
                {
                i = (i+1)
                }
                /* while-7 */} 
              }
              /* Let-6 */} 
            /* Let-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(")")
      Result = EVOID
      }}}
      /* If-2 */} 
    return Result} 
  
// The EID go function for: self_print @ Defclass (throw: true) 
func E_self_print_Defclass_Language (self EID) EID { 
    return /*(sm for self_print @ Defclass= EID)*/ To_Defclass(OBJ(self)).SelfPrint( )} 
  
// -------------- method definition ----------------------------------
//
/* {1} OPT.The go function for: self_print(self:Defmethod) [] */
func (self *Defmethod ) SelfPrint () EID { 
    var Result EID 
    Result = Core.F_print_any(self.Arg.Selector.Id())
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("(")
    if (self.Arg.Args.Id() != CNULL) /* If:2 */{ 
      Result = F_ppvariable_list(self.Arg.Args)
      } else {
      Result = EID{CFALSE.Id(),0}
      /* If-2 */} 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(") : ")
    Result = F_printexp_any(self.SetArg,CFALSE)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    Result = F_lbreak_void()
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    /* update:2 */{ 
      var va_arg1 *Core.PrettyPrinter  
      var va_arg2 int 
      va_arg1 = Core.C_pretty
      va_arg2 = (Core.C_pretty.Index+4)
      /* ---------- now we compile update mClaire/index(va_arg1) := va_arg2 ------- */
      va_arg1.Index = va_arg2
      Result = EID{C__INT,IVAL(va_arg2)}
      /* update-2 */} 
    }
    {
    PRINC(" ")
    F_princ_string(ToString(IfThenElse((self.Inline_ask == CTRUE),
      MakeString("=>").Id(),
      MakeString("->").Id())))
    PRINC(" ")
    Result = F_printexp_any(self.Body,CFALSE)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(" ")
    Result = EVOID
    }}}}}
    {
    /* update:2 */{ 
      var va_arg1 *Core.PrettyPrinter  
      var va_arg2 int 
      va_arg1 = Core.C_pretty
      va_arg2 = (Core.C_pretty.Index-4)
      /* ---------- now we compile update mClaire/index(va_arg1) := va_arg2 ------- */
      va_arg1.Index = va_arg2
      Result = EID{C__INT,IVAL(va_arg2)}
      /* update-2 */} 
    }
    return Result} 
  
// The EID go function for: self_print @ Defmethod (throw: true) 
func E_self_print_Defmethod_Language (self EID) EID { 
    return /*(sm for self_print @ Defmethod= EID)*/ To_Defmethod(OBJ(self)).SelfPrint( )} 
  
// -------------- array definition -----------------------------------
/* {1} OPT.The go function for: self_print(self:Defarray) [] */
func (self *Defarray ) SelfPrint () EID { 
    var Result EID 
    Result = Core.F_CALL(C_print,ARGS(self.Arg.Args.At(1-1).ToEID()))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("[")
    /* Let:2 */{ 
      var g0411UU *ClaireList  
      /* noccur = 1 */
      var g0411UU_try04123 EID 
      g0411UU_try04123 = self.Arg.Args.Cdr()
      /* ERROR PROTECTION INSERTED (g0411UU-Result) */
      if ErrorIn(g0411UU_try04123) {Result = g0411UU_try04123
      } else {
      g0411UU = ToList(OBJ(g0411UU_try04123))
      Result = F_ppvariable_list(g0411UU)
      }
      /* Let-2 */} 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("] : ")
    Result = Core.F_CALL(C_print,ARGS(self.SetArg.ToEID()))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    Result = F_lbreak_void()
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    /* update:2 */{ 
      var va_arg1 *Core.PrettyPrinter  
      var va_arg2 int 
      va_arg1 = Core.C_pretty
      va_arg2 = (Core.C_pretty.Index+4)
      /* ---------- now we compile update mClaire/index(va_arg1) := va_arg2 ------- */
      va_arg1.Index = va_arg2
      Result = EID{C__INT,IVAL(va_arg2)}
      /* update-2 */} 
    }
    {
    PRINC(" := ")
    Result = F_printexp_any(self.Body,CFALSE)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(" ")
    Result = EVOID
    }}}}}
    {
    /* update:2 */{ 
      var va_arg1 *Core.PrettyPrinter  
      var va_arg2 int 
      va_arg1 = Core.C_pretty
      va_arg2 = (Core.C_pretty.Index-4)
      /* ---------- now we compile update mClaire/index(va_arg1) := va_arg2 ------- */
      va_arg1.Index = va_arg2
      Result = EID{C__INT,IVAL(va_arg2)}
      /* update-2 */} 
    }
    return Result} 
  
// The EID go function for: self_print @ Defarray (throw: true) 
func E_self_print_Defarray_Language (self EID) EID { 
    return /*(sm for self_print @ Defarray= EID)*/ To_Defarray(OBJ(self)).SelfPrint( )} 
  
// -------------- rule definition ------------------------------------
/* {1} OPT.The go function for: self_print(self:Defrule) [] */
func (self *Defrule ) SelfPrint () EID { 
    var Result EID 
    self.Ident.Princ()
    PRINC("(")
    Result = F_ppvariable_list(self.Args)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(") :: rule(")
    Result = F_lbreak_integer(4)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(" ")
    Result = Core.F_CALL(C_print,ARGS(self.Arg.ToEID()))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(" ")
    Result = F_lbreak_integer(4)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("=> ")
    Result = Core.F_CALL(C_print,ARGS(self.Body.ToEID()))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(")")
    Result = EVOID
    }}}}}
    {
    /* update:2 */{ 
      var va_arg1 *Core.PrettyPrinter  
      var va_arg2 int 
      va_arg1 = Core.C_pretty
      va_arg2 = (Core.C_pretty.Index-4)
      /* ---------- now we compile update mClaire/index(va_arg1) := va_arg2 ------- */
      va_arg1.Index = va_arg2
      Result = EID{C__INT,IVAL(va_arg2)}
      /* update-2 */} 
    }
    return Result} 
  
// The EID go function for: self_print @ Defrule (throw: true) 
func E_self_print_Defrule_Language (self EID) EID { 
    return /*(sm for self_print @ Defrule= EID)*/ To_Defrule(OBJ(self)).SelfPrint( )} 
  
/* {1} OPT.The go function for: self_print(self:Defvar) [] */
func (self *Defvar ) SelfPrint () EID { 
    var Result EID 
    Result = F_ppvariable_Variable(self.Ident)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(" := ")
    Result = F_printexp_any(self.Arg,CFALSE)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("")
    Result = EVOID
    }}
    return Result} 
  
// The EID go function for: self_print @ Defvar (throw: true) 
func E_self_print_Defvar_Language (self EID) EID { 
    return /*(sm for self_print @ Defvar= EID)*/ To_Defvar(OBJ(self)).SelfPrint( )} 
  
// *********************************************************************
// *     Part 2: the general instantiation macro                       *
// *********************************************************************
// creation of a new object
//
/* {1} OPT.The go function for: self_eval(self:Definition) [] */
func (self *Definition ) SelfEval () EID { 
    var Result EID 
    /* Let:2 */{ 
      var _Zc *ClaireClass   = self.Arg
      /* noccur = 3 */
      /* Let:3 */{ 
        var _Zo *ClaireObject  
        /* noccur = 2 */
        var _Zo_try04144 EID 
        if (_Zc.Open <= 1) /* If:4 */{ 
          _Zo_try04144 = ToException(Core.C_general_error.Make(MakeString("[105] cannot instantiate ~S").Id(),MakeConstantList(_Zc.Id()).Id())).Close()
          } else {
          _Zo_try04144 = EID{CFALSE.Id(),0}
          /* If-4 */} 
        /* ERROR PROTECTION INSERTED (_Zo_try04144-_Zo_try04144) */
        if !ErrorIn(_Zo_try04144) {
        _Zo_try04144 = EID{F_new_object_class(_Zc).Id(),0}
        }
        /* ERROR PROTECTION INSERTED (_Zo-Result) */
        if ErrorIn(_Zo_try04144) {Result = _Zo_try04144
        } else {
        _Zo = ToObject(OBJ(_Zo_try04144))
        /* Let:4 */{ 
          var g0415UU *ClaireList  
          /* noccur = 1 */
          var g0415UU_try04165 EID 
          g0415UU_try04165 = F_Language_new_writes_object(_Zo,self.Args)
          /* ERROR PROTECTION INSERTED (g0415UU-Result) */
          if ErrorIn(g0415UU_try04165) {Result = g0415UU_try04165
          } else {
          g0415UU = ToList(OBJ(g0415UU_try04165))
          Result = Core.F_Core_new_defaults_object(_Zo,g0415UU)
          }
          /* Let-4 */} 
        }
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: self_eval @ Definition (throw: true) 
func E_self_eval_Definition (self EID) EID { 
    return /*(sm for self_eval @ Definition= EID)*/ To_Definition(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Definition 
func EVAL_Definition (x *ClaireAny) EID { 
     return To_Definition(x).SelfEval()} 
  
// fast definition : no inverse management + no "close" method
/* {1} OPT.The go function for: fast_definition?(c:class) [] */
func F_Language_fast_definition_ask_class (c *ClaireClass ) *ClaireBoolean  { 
    // procedure body with s =  
var Result *ClaireBoolean  
    /* and:2 */{ 
      var v_and2 *ClaireBoolean  
      
      v_and2 = Core.F__sup_integer(c.Open,1)
      if (v_and2 == CFALSE) {Result = CFALSE
      } else /* arg:3 */{ 
        v_and2 = Equal(C_method.Id(),Core.F_owner_any(Core.F__at_property1(C_close,c).Id()).Id()).Not
        if (v_and2 == CFALSE) {Result = CFALSE
        } else /* arg:4 */{ 
          /* Let:5 */{ 
            var g0417UU *ClaireAny  
            /* noccur = 1 */
            /* For:6 */{ 
              var s *ClaireAny  
              _ = s
              g0417UU= CFALSE.Id()
              for _,s = range(c.Slots.ValuesO())/* loop:7 */{ 
                if ((ToRestriction(s).Selector.Inverse.Id() != CNULL) || 
                    ((ToRestriction(s).Selector.Store_ask == CTRUE) || 
                      (ToRestriction(s).Selector.IfWrite != CNULL))) /* If:8 */{ 
                   /*v = g0417UU, s =any*/
g0417UU = CTRUE.Id()
                  break
                  /* If-8 */} 
                /* loop-7 */} 
              /* For-6 */} 
            v_and2 = Core.F_not_any(g0417UU)
            /* Let-5 */} 
          if (v_and2 == CFALSE) {Result = CFALSE
          } else /* arg:5 */{ 
            Result = CTRUE/* arg-5 */} 
          /* arg-4 */} 
        /* arg-3 */} 
      /* and-2 */} 
    return Result} 
  
// The EID go function for: fast_definition? @ class (throw: false) 
func E_Language_fast_definition_ask_class (c EID) EID { 
    return EID{/*(sm for fast_definition? @ class= boolean)*/ F_Language_fast_definition_ask_class(ToClass(OBJ(c)) ).Id(),0}} 
  
// then the evaluation is simpler ! write_fast does the range checking (may return an error) 
/* {1} OPT.The go function for: self_eval(self:DefFast) [] */
func (self *LanguageDefFast ) SelfEval () EID { 
    var Result EID 
    /* Let:2 */{ 
      var _Zo *ClaireObject   = F_new_object_class(self.Arg)
      /* noccur = 2 */
      /* For:3 */{ 
        var x *ClaireAny  
        _ = x
        Result= EID{CFALSE.Id(),0}
        var x_support *ClaireList  
        x_support = self.Args
        x_len := x_support.Length()
        for i_it := 0; i_it < x_len; i_it++ { 
          x = x_support.At(i_it)
          var void_try5 EID 
          _ = void_try5
          /* Let:5 */{ 
            var p *ClaireProperty  
            /* noccur = 1 */
            var p_try04186 EID 
            p_try04186 = F_make_a_property_any(To_Call(x).Args.At(1-1))
            /* ERROR PROTECTION INSERTED (p-void_try5) */
            if ErrorIn(p_try04186) {void_try5 = p_try04186
            } else {
            p = ToProperty(OBJ(p_try04186))
            /* LetEID:6 */{ 
              var g0419UU EID 
              g0419UU = EVAL(To_Call(x).Args.At(2-1))
              /* ERROR PROTECTION INSERTED (g0419UU-void_try5) */
              if ErrorIn(g0419UU) {void_try5 = g0419UU
              } else {
              void_try5 = p.WriteEID(_Zo,g0419UU)}
              /* LetEID-6 */} 
            }
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (void_try5-Result) */
          if ErrorIn(void_try5) {Result = void_try5
          Result = void_try5
          break
          } else {
          }
          /* loop-4 */} 
        /* For-3 */} 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = EID{_Zo.Id(),0}
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: self_eval @ DefFast (throw: true) 
func E_self_eval_DefFast (self EID) EID { 
    return /*(sm for self_eval @ DefFast= EID)*/ To_LanguageDefFast(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: DefFast 
func EVAL_Language_DefFast (x *ClaireAny) EID { 
     return To_LanguageDefFast(x).SelfEval()} 
  
// for a fast_definition, simpler eval
// the instantiation body is a sequence of words from which the initialization
// of the object must be built.
// CLAIRE4 : renamed complete(self:object,%l:list) to new_writes()
/* {1} OPT.The go function for: new_writes(self:object,%l:list) [] */
func F_Language_new_writes_object (self *ClaireObject ,_Zl *ClaireList ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var lp *ClaireList   = ToType(CEMPTY.Id()).EmptyList()
      /* noccur = 3 */
      /* For:3 */{ 
        var x *ClaireAny  
        _ = x
        Result= EID{CFALSE.Id(),0}
        var x_support *ClaireList  
        x_support = _Zl
        x_len := x_support.Length()
        for i_it := 0; i_it < x_len; i_it++ { 
          x = x_support.At(i_it)
          var void_try5 EID 
          _ = void_try5
          /* Let:5 */{ 
            var p *ClaireProperty  
            /* noccur = 4 */
            var p_try04226 EID 
            p_try04226 = F_make_a_property_any(To_Call(x).Args.At(1-1))
            /* ERROR PROTECTION INSERTED (p-void_try5) */
            if ErrorIn(p_try04226) {void_try5 = p_try04226
            } else {
            p = ToProperty(OBJ(p_try04226))
            /* Let:6 */{ 
              var y *ClaireAny  
              /* noccur = 4 */
              var y_try04237 EID 
              y_try04237 = EVAL(To_Call(x).Args.At(2-1))
              /* ERROR PROTECTION INSERTED (y-void_try5) */
              if ErrorIn(y_try04237) {void_try5 = y_try04237
              } else {
              y = ANY(y_try04237)
              /* Let:7 */{ 
                var s *ClaireObject   = Core.F__at_property1(p,self.Isa)
                /* noccur = 2 */
                if (C_slot.Id() == s.Isa.Id()) /* If:8 */{ 
                  /* Let:9 */{ 
                    var g0420 *ClaireSlot   = ToSlot(s.Id())
                    /* noccur = 4 */
                    if (y == CNULL) /* If:10 */{ 
                      lp = lp.AddFast(p.Id())
                      /* If-10 */} 
                    if (g0420.Range.Contains(y) != CTRUE) /* If:10 */{ 
                      void_try5 = Core.F_range_is_wrong_slot(g0420,y)
                      } else {
                      void_try5 = Core.F_update_property(p,
                        self,
                        g0420.Index,
                        g0420.Srange,
                        y)
                      /* If-10 */} 
                    /* ERROR PROTECTION INSERTED (void_try5-void_try5) */
                    if ErrorIn(void_try5) {Result = void_try5
                    break
                    } else {
                    }
                    /* Let-9 */} 
                  } else {
                  void_try5 = ToException(Core.C_general_error.Make(MakeString("[106] the object ~S does not understand ~S").Id(),MakeConstantList(self.Id(),p.Id()).Id())).Close()
                  /* If-8 */} 
                /* Let-7 */} 
              }
              /* Let-6 */} 
            }
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (void_try5-Result) */
          if ErrorIn(void_try5) {Result = void_try5
          Result = void_try5
          break
          } else {
          }
          /* loop-4 */} 
        /* For-3 */} 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = EID{lp.Id(),0}
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: new_writes @ object (throw: true) 
func E_Language_new_writes_object (self EID,_Zl EID) EID { 
    return /*(sm for new_writes @ object= EID)*/ F_Language_new_writes_object(ToObject(OBJ(self)),ToList(OBJ(_Zl)) )} 
  
// creation of a new named object
/* {1} OPT.The go function for: self_eval(self:Defobj) [] */
func (self *Defobj ) SelfEval () EID { 
    var Result EID 
    /* Let:2 */{ 
      var _Zc *ClaireClass   = self.Arg
      /* noccur = 7 */
      /* Let:3 */{ 
        var _Zo *ClaireObject   = ToObject(CNULL)
        /* noccur = 8 */
        if (_Zc.Open <= 1) /* If:4 */{ 
          Result = ToException(Core.C_general_error.Make(MakeString("[105] cannot instantiate ~S").Id(),MakeConstantList(_Zc.Id()).Id())).Close()
          } else {
          Result = EID{CFALSE.Id(),0}
          /* If-4 */} 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        if (_Zc.IsIn(C_thing) == CTRUE) /* If:4 */{ 
          var _Zo_try04255 EID 
          _Zo_try04255 = F_new_thing_class(_Zc,self.Ident)
          /* ERROR PROTECTION INSERTED (_Zo-Result) */
          if ErrorIn(_Zo_try04255) {Result = _Zo_try04255
          } else {
          _Zo = ToObject(OBJ(_Zo_try04255))
          Result = EID{_Zo.Id(),0}
          if (_Zo.Isa.IsIn(C_property) == CTRUE) /* If:5 */{ 
            /* Let:6 */{ 
              var g0424 *ClaireProperty   = ToProperty(_Zo.Id())
              /* noccur = 2 */
              if (g0424.Restrictions.Length() > 0) /* If:7 */{ 
                Result = ToException(Core.C_general_error.Make(MakeString("[188] the property ~S is already defined").Id(),MakeConstantList(g0424.Id()).Id())).Close()
                } else {
                Result = EID{CFALSE.Id(),0}
                /* If-7 */} 
              /* Let-6 */} 
            } else {
            Result = EID{CFALSE.Id(),0}
            /* If-5 */} 
          }
          } else {
          _Zo = F_new_object_class(_Zc)
          if (_Zc.Open == ClEnv.Open) /* If:5 */{ 
            _Zc.Instances.AddFast(_Zo.Id())
            /* If-5 */} 
          Result = self.Ident.Put(_Zo.Id()).ToEID()
          /* If-4 */} 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        /* Let:4 */{ 
          var g0426UU *ClaireList  
          /* noccur = 1 */
          var g0426UU_try04275 EID 
          g0426UU_try04275 = F_Language_new_writes_object(_Zo,self.Args)
          /* ERROR PROTECTION INSERTED (g0426UU-Result) */
          if ErrorIn(g0426UU_try04275) {Result = g0426UU_try04275
          } else {
          g0426UU = ToList(OBJ(g0426UU_try04275))
          Result = Core.F_Core_new_defaults_object(_Zo,g0426UU)
          }
          /* Let-4 */} 
        }}
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: self_eval @ Defobj (throw: true) 
func E_self_eval_Defobj (self EID) EID { 
    return /*(sm for self_eval @ Defobj= EID)*/ To_Defobj(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Defobj 
func EVAL_Defobj (x *ClaireAny) EID { 
     return To_Defobj(x).SelfEval()} 
  
// creation of a new named object
// note that final() is the marker of a forward definition in CLAIRE4
/* {1} OPT.The go function for: self_eval(self:Defclass) [] */
func (self *Defclass ) SelfEval () EID { 
    var Result EID 
    if ((C_class.Id() == Core.F_owner_any(self.Ident.Get()).Id()) && 
        ((ToClass(self.Ident.Get()).Open != ClEnv.Final) || 
            (self.Arg.Id() != ToClass(self.Ident.Get()).Superclass.Id()))) /* If:2 */{ 
      Result = ToException(Core.C_general_error.Make(MakeString("[107] class re-definition is not valid: ~S").Id(),MakeConstantList(self.Id()).Id())).Close()
      } else {
      /* Let:3 */{ 
        var _Zo *ClaireClass  
        /* noccur = 13 */
        var _Zo_try04294 EID 
        _Zo_try04294 = self.Ident.Class_I(self.Arg)
        /* ERROR PROTECTION INSERTED (_Zo-Result) */
        if ErrorIn(_Zo_try04294) {Result = _Zo_try04294
        } else {
        _Zo = ToClass(OBJ(_Zo_try04294))
        /* For:4 */{ 
          var x *ClaireAny  
          _ = x
          Result= EID{CFALSE.Id(),0}
          var x_support *ClaireList  
          x_support = self.Args
          x_len := x_support.Length()
          for i_it := 0; i_it < x_len; i_it++ { 
            x = x_support.At(i_it)
            var void_try6 EID 
            _ = void_try6
            /* Let:6 */{ 
              var v *ClaireAny   = CNULL
              /* noccur = 5 */
              if (x.Isa.IsIn(C_Call) == CTRUE) /* If:7 */{ 
                /* Let:8 */{ 
                  var g0428 *Call   = To_Call(x)
                  /* noccur = 4 */
                  var v_try04309 EID 
                  v_try04309 = EVAL(g0428.Args.At(2-1))
                  /* ERROR PROTECTION INSERTED (v-void_try6) */
                  if ErrorIn(v_try04309) {void_try6 = v_try04309
                  Result = v_try04309
                  break
                  } else {
                  v = ANY(v_try04309)
                  void_try6 = v.ToEID()
                  g0428 = To_Call(g0428.Args.At(1-1))
                  void_try6 = EID{g0428.Id(),0}
                  }
                  {
                  x = g0428.Id()
                  void_try6 = x.ToEID()
                  }
                  /* Let-8 */} 
                } else {
                void_try6 = EID{CFALSE.Id(),0}
                /* If-7 */} 
              /* ERROR PROTECTION INSERTED (void_try6-void_try6) */
              if ErrorIn(void_try6) {Result = void_try6
              break
              } else {
              /* Let:7 */{ 
                var rt *ClaireTypeExpression  
                /* noccur = 6 */
                var rt_try04318 EID 
                rt_try04318 = F_extract_type_any(ANY(Core.F_CALL(C_range,ARGS(x.ToEID()))))
                /* ERROR PROTECTION INSERTED (rt-void_try6) */
                if ErrorIn(rt_try04318) {void_try6 = rt_try04318
                } else {
                rt = ToTypeExpression(OBJ(rt_try04318))
                /* Let:8 */{ 
                  var p *ClaireProperty  
                  /* noccur = 4 */
                  var p_try04329 EID 
                  p_try04329 = F_make_a_property_any(ANY(Core.F_CALL(C_mClaire_pname,ARGS(x.ToEID()))))
                  /* ERROR PROTECTION INSERTED (p-void_try6) */
                  if ErrorIn(p_try04329) {void_try6 = p_try04329
                  } else {
                  p = ToProperty(OBJ(p_try04329))
                  var g0433I *ClaireBoolean  
                  var g0433I_try04349 EID 
                  /* and:9 */{ 
                    var v_and9 *ClaireBoolean  
                    
                    v_and9 = Core.F_known_ask_any(v)
                    if (v_and9 == CFALSE) {g0433I_try04349 = EID{CFALSE.Id(),0}
                    } else /* arg:10 */{ 
                      var v_and9_try043511 EID 
                      /* Let:11 */{ 
                        var g0436UU *ClaireBoolean  
                        /* noccur = 1 */
                        var g0436UU_try043712 EID 
                        g0436UU_try043712 = Core.F_BELONG(v,rt.Id())
                        /* ERROR PROTECTION INSERTED (g0436UU-v_and9_try043511) */
                        if ErrorIn(g0436UU_try043712) {v_and9_try043511 = g0436UU_try043712
                        } else {
                        g0436UU = ToBoolean(OBJ(g0436UU_try043712))
                        v_and9_try043511 = EID{g0436UU.Not.Id(),0}
                        }
                        /* Let-11 */} 
                      /* ERROR PROTECTION INSERTED (v_and9-g0433I_try04349) */
                      if ErrorIn(v_and9_try043511) {g0433I_try04349 = v_and9_try043511
                      } else {
                      v_and9 = ToBoolean(OBJ(v_and9_try043511))
                      if (v_and9 == CFALSE) {g0433I_try04349 = EID{CFALSE.Id(),0}
                      } else /* arg:11 */{ 
                        g0433I_try04349 = EID{CTRUE.Id(),0}/* arg-11 */} 
                      /* arg-10 */} 
                    }
                    /* and-9 */} 
                  /* ERROR PROTECTION INSERTED (g0433I-void_try6) */
                  if ErrorIn(g0433I_try04349) {void_try6 = g0433I_try04349
                  } else {
                  g0433I = ToBoolean(OBJ(g0433I_try04349))
                  if (g0433I == CTRUE) /* If:9 */{ 
                    void_try6 = ToException(Core.C_general_error.Make(MakeString("[108] default(~S) = ~S does not belong to ~S").Id(),MakeConstantList(x,v,rt.Id()).Id())).Close()
                    } else {
                    void_try6 = EID{CFALSE.Id(),0}
                    /* If-9 */} 
                  }
                  /* ERROR PROTECTION INSERTED (void_try6-void_try6) */
                  if ErrorIn(void_try6) {Result = void_try6
                  break
                  } else {
                  /* Let:9 */{ 
                    var s2test *ClaireAny  
                    /* noccur = 2 */
                    /* Let:10 */{ 
                      var sx_some *ClaireAny   = CNULL
                      /* noccur = 2 */
                      /* For:11 */{ 
                        var sx *ClaireAny  
                        _ = sx
                        for _,sx = range(self.Arg.Slots.ValuesO())/* loop:12 */{ 
                          if (ToRestriction(sx).Selector.Id() == p.Id()) /* If:13 */{ 
                             /*v = s2test, s =void*/
sx_some = sx
                            break
                            /* If-13 */} 
                          /* loop-12 */} 
                        /* For-11 */} 
                      s2test = sx_some
                      /* Let-10 */} 
                    if (s2test != CNULL) /* If:10 */{ 
                      /* Let:11 */{ 
                        var s2 *ClaireSlot   = ToSlot(s2test)
                        /* noccur = 2 */
                        if (p.Open <= 0) /* If:12 */{ 
                          void_try6 = ToException(Core.C_general_error.Make(MakeString("[181] cannot overide a slot for a closed property ~S").Id(),MakeConstantList(p.Id()).Id())).Close()
                          /* If!12 */}  else if (ToType(rt.Id()).Included(s2.Range) != CTRUE) /* If:12 */{ 
                          void_try6 = ToException(Core.C_general_error.Make(MakeString("[XXX] slot redefinition of ~S must be covariant, ~S is not a subtype").Id(),MakeConstantList(s2.Id(),rt.Id()).Id())).Close()
                          } else {
                          void_try6 = EID{CFALSE.Id(),0}
                          /* If-12 */} 
                        /* Let-11 */} 
                      } else {
                      void_try6 = EID{CNULL,0}
                      /* If-10 */} 
                    /* Let-9 */} 
                  /* ERROR PROTECTION INSERTED (void_try6-void_try6) */
                  if ErrorIn(void_try6) {Result = void_try6
                  break
                  } else {
                  void_try6 = EID{Core.F_close_slot(_Zo.AddSlot(p,ToType(rt.Id()),F_Language_getDefault_type(ToType(rt.Id()),v))).Id(),0}
                  }}
                  }
                  /* Let-8 */} 
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (void_try6-void_try6) */
              if ErrorIn(void_try6) {Result = void_try6
              break
              } else {
              }}
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (void_try6-Result) */
            if ErrorIn(void_try6) {Result = void_try6
            Result = void_try6
            break
            } else {
            }
            /* loop-5 */} 
          /* For-4 */} 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        F_close_class(_Zo)
        if (self.Forward_ask == CTRUE) /* If:4 */{ 
          _Zo.Open = ClEnv.Final
          /* If!4 */}  else if (_Zo.Open == ClEnv.Final) /* If:4 */{ 
          _Zo.Open = self.Arg.Open
          /* If-4 */} 
        if (ToType(_Zo.Id()).Included(ToType(C_primitive.Id())) == CTRUE) /* If:4 */{ 
          _Zo.Open = -1
          /* If-4 */} 
        _Zo.Params = self.Params
        /* For:4 */{ 
          var p *ClaireAny  
          _ = p
          var p_support *ClaireList  
          p_support = self.Params
          p_len := p_support.Length()
          for i_it := 0; i_it < p_len; i_it++ { 
            p = p_support.At(i_it)
            ToRelation(p).Open = 0
            /* loop-5 */} 
          /* For-4 */} 
        Result = F_attach_comment_any(_Zo.Id())
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        Result = EID{_Zo.Id(),0}
        }}
        }
        /* Let-3 */} 
      /* If-2 */} 
    return Result} 
  
// The EID go function for: self_eval @ Defclass (throw: true) 
func E_self_eval_Defclass (self EID) EID { 
    return /*(sm for self_eval @ Defclass= EID)*/ To_Defclass(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Defclass 
func EVAL_Defclass (x *ClaireAny) EID { 
     return To_Defclass(x).SelfEval()} 
  
// we compute the proper default value (reused by compiler) - for int, float, sets and lists
/* {1} OPT.The go function for: getDefault(rt:type,v:any) [] */
func F_Language_getDefault_type (rt *ClaireType ,v *ClaireAny ) *ClaireAny  { 
    // use function body compiling 
if (v == CNULL) /* body If:2 */{ 
      if (rt.Included(ToType(C_integer.Id())) == CTRUE) /* body If:3 */{ 
        return  MakeInteger(0).Id()
        }  else if (rt.Included(ToType(C_float.Id())) == CTRUE) /* body If:3 */{ 
        return  MakeFloat(0).Id()
        }  else if (rt.Included(ToType(C_set.Id())) == CTRUE) /* body If:3 */{ 
        return  Core.F_of_extract_type(rt).EmptySet().Id()
        }  else if (rt.Included(ToType(C_list.Id())) == CTRUE) /* body If:3 */{ 
        return  Core.F_of_extract_type(rt).EmptyList().Id()
        } else {
        return  CNULL
        /* body If-3 */} 
      } else {
      return  v
      /* body If-2 */} 
    } 
  
// The EID go function for: getDefault @ type (throw: false) 
func E_Language_getDefault_type (rt EID,v EID) EID { 
    return /*(sm for getDefault @ type= any)*/ F_Language_getDefault_type(ToType(OBJ(rt)),ANY(v) ).ToEID()} 
  
// method definition
// v0.01
/* {1} OPT.The go function for: self_eval(self:Defmethod) [] */
func (self *Defmethod ) SelfEval () EID { 
    var Result EID 
    if (self.Arg.Isa.IsIn(C_Call) != CTRUE) /* If:2 */{ 
      Result = ToException(Core.C_general_error.Make(MakeString("[110] wrong signature definition ~S").Id(),MakeConstantList(self.Arg.Id()).Id())).Close()
      } else {
      Result = EID{CFALSE.Id(),0}
      /* If-2 */} 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    /* Let:2 */{ 
      var p *ClaireProperty  
      /* noccur = 6 */
      var p_try04383 EID 
      p_try04383 = F_make_a_property_any(self.Arg.Selector.Id())
      /* ERROR PROTECTION INSERTED (p-Result) */
      if ErrorIn(p_try04383) {Result = p_try04383
      } else {
      p = ToProperty(OBJ(p_try04383))
      /* Let:3 */{ 
        var l *ClaireList   = self.Arg.Args
        /* noccur = 3 */
        /* Let:4 */{ 
          var lv *ClaireList  
          /* noccur = 3 */
          if ((l.Length() == 1) && 
              (l.At(1-1) == ClEnv.Id())) /* If:5 */{ 
            /* Construct:6 */{ 
              var v_bag_arg *ClaireAny  
              lv= ToType(CEMPTY.Id()).EmptyList()
              /* Let:7 */{ 
                var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
                /* noccur = 5 */
                _CL_obj.Pname = Core.F_symbol_I_string2(MakeString("XfakeParameter"))
                _CL_obj.Range = ToType(C_void.Id())
                v_bag_arg = _CL_obj.Id()
                /* Let-7 */} 
              lv.AddFast(v_bag_arg)/* Construct-6 */} 
            } else {
            lv = l
            /* If-5 */} 
          /* Let:5 */{ 
            var lp *ClaireList  
            /* noccur = 1 */
            var lp_try04396 EID 
            lp_try04396 = F_extract_signature_list(lv)
            /* ERROR PROTECTION INSERTED (lp-Result) */
            if ErrorIn(lp_try04396) {Result = lp_try04396
            } else {
            lp = ToList(OBJ(lp_try04396))
            /* Let:6 */{ 
              var lrange *ClaireList  
              /* noccur = 3 */
              var lrange_try04407 EID 
              lrange_try04407 = F_extract_range_any(self.SetArg,lv,ToList(C_LDEF.Value))
              /* ERROR PROTECTION INSERTED (lrange-Result) */
              if ErrorIn(lrange_try04407) {Result = lrange_try04407
              } else {
              lrange = ToList(OBJ(lrange_try04407))
              /* Let:7 */{ 
                var lbody *ClaireList  
                /* noccur = 4 */
                var lbody_try04418 EID 
                lbody_try04418 = F_extract_status_any(self.Body)
                /* ERROR PROTECTION INSERTED (lbody-Result) */
                if ErrorIn(lbody_try04418) {Result = lbody_try04418
                } else {
                lbody = ToList(OBJ(lbody_try04418))
                /* Let:8 */{ 
                  var m *ClaireMethod   = F_add_method_property(p,lp,ToType(lrange.At(1-1)),ToInteger(lbody.At(1-1)).Value,ToFunction(lbody.At(2-1)))
                  /* noccur = 14 */
                  if ((p.Open > 0) && 
                      (p.Open <= 1)) /* If:9 */{ 
                    /* Let:10 */{ 
                      var r *ClaireAny  
                      /* noccur = 2 */
                      var r_try044211 EID 
                      /* Let:11 */{ 
                        var r_some *ClaireAny   = CNULL
                        /* noccur = 2 */
                        /* For:12 */{ 
                          var r *ClaireAny  
                          _ = r
                          r_try044211= EID{CFALSE.Id(),0}
                          for _,r = range(p.Restrictions.ValuesO())/* loop:13 */{ 
                            var void_try14 EID 
                            _ = void_try14
                            if (r != m.Id()) /* If:14 */{ 
                              var g0443I *ClaireBoolean  
                              var g0443I_try044415 EID 
                              /* Let:15 */{ 
                                var g0445UU *ClaireAny  
                                /* noccur = 1 */
                                var g0445UU_try044616 EID 
                                g0445UU_try044616 = Core.F_CALL(ToProperty(C__exp.Id()),ARGS(Core.F_CALL(C_domain,ARGS(r.ToEID())),EID{m.Domain.Id(),0}))
                                /* ERROR PROTECTION INSERTED (g0445UU-g0443I_try044415) */
                                if ErrorIn(g0445UU_try044616) {g0443I_try044415 = g0445UU_try044616
                                } else {
                                g0445UU = ANY(g0445UU_try044616)
                                g0443I_try044415 = EID{F_boolean_I_any(g0445UU).Id(),0}
                                }
                                /* Let-15 */} 
                              /* ERROR PROTECTION INSERTED (g0443I-void_try14) */
                              if ErrorIn(g0443I_try044415) {void_try14 = g0443I_try044415
                              } else {
                              g0443I = ToBoolean(OBJ(g0443I_try044415))
                              if (g0443I == CTRUE) /* If:15 */{ 
                                 /*v = r_try044211, s =EID*/
r_some = r
                                r_try044211 = r_some.ToEID()
                                break
                                } else {
                                void_try14 = EID{CFALSE.Id(),0}
                                /* If-15 */} 
                              }
                              } else {
                              void_try14 = EID{CFALSE.Id(),0}
                              /* If-14 */} 
                            /* ERROR PROTECTION INSERTED (void_try14-r_try044211) */
                            if ErrorIn(void_try14) {r_try044211 = void_try14
                            r_try044211 = void_try14
                            break
                            } else {
                            }
                            /* loop-13 */} 
                          /* For-12 */} 
                        /* ERROR PROTECTION INSERTED (r_try044211-r_try044211) */
                        if !ErrorIn(r_try044211) {
                        r_try044211 = r_some.ToEID()
                        }
                        /* Let-11 */} 
                      /* ERROR PROTECTION INSERTED (r-Result) */
                      if ErrorIn(r_try044211) {Result = r_try044211
                      } else {
                      r = ANY(r_try044211)
                      if (r != CNULL) /* If:11 */{ 
                        Result = Core.F_tformat_string(MakeString("--- WARNING ! [186] conflict between ~S and ~S is dangerous since ~S is closed\n"),1,MakeConstantList(m.Id(),r,p.Id()))
                        } else {
                        Result = EID{CNULL,0}
                        /* If-11 */} 
                      }
                      /* Let-10 */} 
                    } else {
                    Result = EID{CFALSE.Id(),0}
                    /* If-9 */} 
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  C_LDEF.Value = ToType(C_any.Id()).EmptyList().Id()
                  if (lbody.At(3-1) != C_body.Id()) /* If:9 */{ 
                    Core.F_tformat_string(MakeString("---- jito for ~S\n"),2,MakeConstantList(m.Id()))
                    /* update:10 */{ 
                      var va_arg1 *ClaireMethod  
                      var va_arg2 *ClaireLambda  
                      va_arg1 = m
                      var va_arg2_try044711 EID 
                      /* Let:11 */{ 
                        var g0448UU *ClaireLambda  
                        /* noccur = 1 */
                        var g0448UU_try044912 EID 
                        g0448UU_try044912 = F_lambda_I_list(lv,lbody.At(3-1))
                        /* ERROR PROTECTION INSERTED (g0448UU-va_arg2_try044711) */
                        if ErrorIn(g0448UU_try044912) {va_arg2_try044711 = g0448UU_try044912
                        } else {
                        g0448UU = ToLambda(OBJ(g0448UU_try044912))
                        va_arg2_try044711 = F_Language_jito_any(g0448UU.Id())
                        }
                        /* Let-11 */} 
                      /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                      if ErrorIn(va_arg2_try044711) {Result = va_arg2_try044711
                      } else {
                      va_arg2 = ToLambda(OBJ(va_arg2_try044711))
                      /* ---------- now we compile update formula(va_arg1) := va_arg2 ------- */
                      va_arg1.Formula = va_arg2
                      Result = EID{va_arg2.Id(),0}
                      }
                      /* update-10 */} 
                    } else {
                    Result = EID{CFALSE.Id(),0}
                    /* If-9 */} 
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  if (lrange.Length() > 1) /* If:9 */{ 
                    m.Typing = lrange.At(2-1)
                    /* If-9 */} 
                  m.Inline_ask = self.Inline_ask
                  Result = F_attach_comment_any(m.Id())
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  Core.F_close_method(m)
                  if ((p.Id() == C_close.Id()) && 
                      (m.Range.Included(ToType(Core.F_domain_I_restriction(ToRestriction(m.Id())).Id())) != CTRUE)) /* If:9 */{ 
                    Result = ToException(Core.C_general_error.Make(MakeString("[184] the close method ~S has a wrong range").Id(),MakeConstantList(m.Id()).Id())).Close()
                    } else {
                    Result = EID{CFALSE.Id(),0}
                    /* If-9 */} 
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  Result = EID{m.Id(),0}
                  }}}}
                  /* Let-8 */} 
                }
                /* Let-7 */} 
              }
              /* Let-6 */} 
            }
            /* Let-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      }
      /* Let-2 */} 
    }
    return Result} 
  
// The EID go function for: self_eval @ Defmethod (throw: true) 
func E_self_eval_Defmethod (self EID) EID { 
    return /*(sm for self_eval @ Defmethod= EID)*/ To_Defmethod(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Defmethod 
func EVAL_Defmethod (x *ClaireAny) EID { 
     return To_Defmethod(x).SelfEval()} 
  
// v3.2.24 : -1 : final
// attach a cute comment if needed ... to a defclass or a defmethod
/* {1} OPT.The go function for: attach_comment(x:any) [] */
func F_attach_comment_any (x *ClaireAny ) EID { 
    var Result EID 
    if ((ToBoolean(C_NeedComment.Value) == CTRUE) && 
        (C_iClaire_LastComment.Value != CNULL)) /* If:2 */{ 
      Result = Core.F_write_property(C_comment,ToObject(x),C_iClaire_LastComment.Value)
      } else {
      Result = EID{CFALSE.Id(),0}
      /* If-2 */} 
    return Result} 
  
// The EID go function for: attach_comment @ any (throw: true) 
func E_attach_comment_any (x EID) EID { 
    return /*(sm for attach_comment @ any= EID)*/ F_attach_comment_any(ANY(x) )} 
  
// returns the list of types AND modifies LDEF
/* {1} OPT.The go function for: iClaire/extract_signature(l:list) [] */
func F_extract_signature_list (l *ClaireList ) EID { 
    var Result EID 
    C_LDEF.Value = ToType(C_any.Id()).EmptyList().Id()
    /* Let:2 */{ 
      var n int  = 0
      /* noccur = 3 */
      /* Iteration:3 */{ 
        var v_list3 *ClaireList  
        var v *ClaireAny  
        var v_local3 *ClaireAny  
        v_list3 = l
        Result = EID{CreateList(ToType(CEMPTY.Id()),v_list3.Length()).Id(),0}
        for CLcount := 0; CLcount < v_list3.Length(); CLcount++{ 
          v = v_list3.At(CLcount)
          var v_local3_try04505 EID 
          if (v.Isa.IsIn(C_Variable) != CTRUE) /* If:5 */{ 
            v_local3_try04505 = ToException(Core.C_general_error.Make(MakeString("[111] wrong typed argument ~S").Id(),MakeConstantList(v).Id())).Close()
            } else {
            /* Let:6 */{ 
              var p *ClaireAny  
              /* noccur = 3 */
              var p_try04517 EID 
              p_try04517 = F_extract_pattern_any(To_Variable(v).Range.Id(),MakeConstantList(MakeInteger(n).Id()))
              /* ERROR PROTECTION INSERTED (p-v_local3_try04505) */
              if ErrorIn(p_try04517) {v_local3_try04505 = p_try04517
              } else {
              p = ANY(p_try04517)
              n = (n+1)
              if (p == CNULL) /* If:7 */{ 
                v_local3_try04505 = ToException(Core.C_general_error.Make(MakeString("[111] wrong typed argument ~S (~S)").Id(),MakeConstantList(v,To_Variable(v).Range.Id()).Id())).Close()
                } else {
                v_local3_try04505 = EID{CFALSE.Id(),0}
                /* If-7 */} 
              /* ERROR PROTECTION INSERTED (v_local3_try04505-v_local3_try04505) */
              if ErrorIn(v_local3_try04505) {Result = v_local3_try04505
              break
              } else {
              To_Variable(v).Range = F_type_I_any(p)
              v_local3_try04505 = p.ToEID()
              }
              }
              /* Let-6 */} 
            /* If-5 */} 
          /* ERROR PROTECTION INSERTED (v_local3-Result) */
          if ErrorIn(v_local3_try04505) {Result = v_local3_try04505
          Result = v_local3_try04505
          break
          } else {
          v_local3 = ANY(v_local3_try04505)
          ToList(OBJ(Result)).PutAt(CLcount,v_local3)
          } 
        }
        /* Iteration-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: iClaire/extract_signature @ list (throw: true) 
func E_extract_signature_list (l EID) EID { 
    return /*(sm for iClaire/extract_signature @ list= EID)*/ F_extract_signature_list(ToList(OBJ(l)) )} 
  
// takes an <exp> that must belong to <type> and returns the CLAIRE type
// if LDEF is non-empty, it is used as a list of type variable and patterns
// may be returned. In addition, if the path list is non empty, new type
// variables may be defined. a syntax error will produce the unknown value
//
/* {1} OPT.The go function for: iClaire/extract_pattern(x:any,path:list) [] */
func F_extract_pattern_any (x *ClaireAny ,path *ClaireList ) EID { 
    var Result EID 
    if (C_class.Id() == x.Isa.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var g0452 *ClaireClass   = ToClass(x)
        /* noccur = 1 */
        Result = EID{g0452.Id(),0}
        /* Let-3 */} 
      /* If!2 */}  else if (C_set.Id() == x.Isa.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var g0453 *ClaireSet   = ToSet(x)
        /* noccur = 3 */
        /* Let:4 */{ 
          var z *ClaireAny  
          /* noccur = 2 */
          var z_try04635 EID 
          if (Equal(ANY(Core.F_CALL(C_length,ARGS(EID{g0453.Id(),0}))),MakeInteger(1).Id()) == CTRUE) /* If:5 */{ 
            /* Let:6 */{ 
              var g0464UU *ClaireAny  
              /* noccur = 1 */
              var g0464UU_try04657 EID 
              g0464UU_try04657 = Core.F_CALL(C_nth,ARGS(EID{g0453.Id(),0},EID{C__INT,IVAL(1)}))
              /* ERROR PROTECTION INSERTED (g0464UU-z_try04635) */
              if ErrorIn(g0464UU_try04657) {z_try04635 = g0464UU_try04657
              } else {
              g0464UU = ANY(g0464UU_try04657)
              z_try04635 = F_extract_pattern_any(g0464UU,CNIL)
              }
              /* Let-6 */} 
            } else {
            z_try04635 = EID{CFALSE.Id(),0}
            /* If-5 */} 
          /* ERROR PROTECTION INSERTED (z-Result) */
          if ErrorIn(z_try04635) {Result = z_try04635
          } else {
          z = ANY(z_try04635)
          if (z.Isa.IsIn(C_Reference) == CTRUE) /* If:5 */{ 
            /* Let:6 */{ 
              var g0454 *ClaireReference   = To_Reference(z)
              /* noccur = 1 */
              /* Let:7 */{ 
                var w *ClaireReference   = To_Reference(g0454.Copy().Id())
                /* noccur = 3 */
                w.Arg = CTRUE
                Result = EID{w.Id(),0}
                /* Let-7 */} 
              /* Let-6 */} 
            } else {
            Result = EID{g0453.Id(),0}
            /* If-5 */} 
          }
          /* Let-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (x.Isa.IsIn(C_Tuple) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0456 *Tuple   = To_Tuple(x)
        /* noccur = 1 */
        /* Let:4 */{ 
          var ltp *ClaireList  
          /* noccur = 2 */
          var ltp_try04665 EID 
          /* Iteration:5 */{ 
            var v_list5 *ClaireList  
            var z *ClaireAny  
            var v_local5 *ClaireAny  
            v_list5 = g0456.Args
            ltp_try04665 = EID{CreateList(ToType(CEMPTY.Id()),v_list5.Length()).Id(),0}
            for CLcount := 0; CLcount < v_list5.Length(); CLcount++{ 
              z = v_list5.At(CLcount)
              var v_local5_try04677 EID 
              v_local5_try04677 = F_extract_pattern_any(z,path)
              /* ERROR PROTECTION INSERTED (v_local5-ltp_try04665) */
              if ErrorIn(v_local5_try04677) {ltp_try04665 = v_local5_try04677
              ltp_try04665 = v_local5_try04677
              break
              } else {
              v_local5 = ANY(v_local5_try04677)
              ToList(OBJ(ltp_try04665)).PutAt(CLcount,v_local5)
              } 
            }
            /* Iteration-5 */} 
          /* ERROR PROTECTION INSERTED (ltp-Result) */
          if ErrorIn(ltp_try04665) {Result = ltp_try04665
          } else {
          ltp = ToList(OBJ(ltp_try04665))
          var g0468I *ClaireBoolean  
          /* Let:5 */{ 
            var g0469UU *ClaireAny  
            /* noccur = 1 */
            /* For:6 */{ 
              var y *ClaireAny  
              _ = y
              g0469UU= CFALSE.Id()
              var y_support *ClaireList  
              y_support = ltp
              y_len := y_support.Length()
              for i_it := 0; i_it < y_len; i_it++ { 
                y = y_support.At(i_it)
                if (y == CNULL) /* If:8 */{ 
                   /*v = g0469UU, s =any*/
g0469UU = CTRUE.Id()
                  break
                  /* If-8 */} 
                /* loop-7 */} 
              /* For-6 */} 
            g0468I = F_boolean_I_any(g0469UU)
            /* Let-5 */} 
          if (g0468I == CTRUE) /* If:5 */{ 
            Result = EID{CNULL,0}
            } else {
            Result = EID{ltp.Tuple_I().Id(),0}
            /* If-5 */} 
          }
          /* Let-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (x.Isa.IsIn(Core.C_global_variable) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0457 *Core.GlobalVariable   = Core.ToGlobalVariable(x)
        /* noccur = 1 */
        Result = F_extract_pattern_any(g0457.Value,path)
        /* Let-3 */} 
      /* If!2 */}  else if (x.Isa.IsIn(C_Call) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0458 *Call   = To_Call(x)
        /* noccur = 9 */
        /* Let:4 */{ 
          var p *ClaireProperty   = g0458.Selector
          /* noccur = 5 */
          if (p.Id() == Core.C_U.Id()) /* If:5 */{ 
            /* Let:6 */{ 
              var x1 *ClaireAny  
              /* noccur = 2 */
              var x1_try04707 EID 
              x1_try04707 = F_extract_pattern_any(g0458.Args.At(1-1),CNIL)
              /* ERROR PROTECTION INSERTED (x1-Result) */
              if ErrorIn(x1_try04707) {Result = x1_try04707
              } else {
              x1 = ANY(x1_try04707)
              /* Let:7 */{ 
                var x2 *ClaireAny  
                /* noccur = 2 */
                var x2_try04718 EID 
                x2_try04718 = F_extract_pattern_any(g0458.Args.At(2-1),CNIL)
                /* ERROR PROTECTION INSERTED (x2-Result) */
                if ErrorIn(x2_try04718) {Result = x2_try04718
                } else {
                x2 = ANY(x2_try04718)
                if ((x1 == CNULL) || 
                    (x2 == CNULL)) /* If:8 */{ 
                  Result = EID{CNULL,0}
                  } else {
                  Result = EID{Core.F_U_type(ToType(x1),ToType(x2)).Id(),0}
                  /* If-8 */} 
                }
                /* Let-7 */} 
              }
              /* Let-6 */} 
            /* If!5 */}  else if (p.Id() == C__exp.Id()) /* If:5 */{ 
            /* Let:6 */{ 
              var g0472UU *ClaireAny  
              /* noccur = 1 */
              var g0472UU_try04747 EID 
              g0472UU_try04747 = F_extract_pattern_any(g0458.Args.At(1-1),CNIL)
              /* ERROR PROTECTION INSERTED (g0472UU-Result) */
              if ErrorIn(g0472UU_try04747) {Result = g0472UU_try04747
              } else {
              g0472UU = ANY(g0472UU_try04747)
              /* Let:7 */{ 
                var g0473UU *ClaireAny  
                /* noccur = 1 */
                var g0473UU_try04758 EID 
                g0473UU_try04758 = F_extract_pattern_any(g0458.Args.At(2-1),CNIL)
                /* ERROR PROTECTION INSERTED (g0473UU-Result) */
                if ErrorIn(g0473UU_try04758) {Result = g0473UU_try04758
                } else {
                g0473UU = ANY(g0473UU_try04758)
                Result = Core.F_CALL(ToProperty(C__exp.Id()),ARGS(g0472UU.ToEID(),g0473UU.ToEID()))
                }
                /* Let-7 */} 
              }
              /* Let-6 */} 
            /* If!5 */}  else if (p.Id() == C__dot_dot.Id()) /* If:5 */{ 
            /* Let:6 */{ 
              var v1 *ClaireAny   = F_extract_item_any(g0458.Args.At(1-1),CNIL.Id())
              /* noccur = 2 */
              /* Let:7 */{ 
                var v2 *ClaireAny   = F_extract_item_any(g0458.Args.At(2-1),CNIL.Id())
                /* noccur = 2 */
                if ((C_integer.Id() == v1.Isa.Id()) && 
                    (C_integer.Id() == v2.Isa.Id())) /* If:8 */{ 
                  Result = EID{Core.F__dot_dot_integer(ToInteger(v1).Value,ToInteger(v2).Value).Id(),0}
                  } else {
                  Result = EID{CNULL,0}
                  /* If-8 */} 
                /* Let-7 */} 
              /* Let-6 */} 
            /* If!5 */}  else if (p.Id() == C_nth.Id()) /* If:5 */{ 
            Result = F_extract_pattern_nth_list(g0458.Args,path)
            /* If!5 */}  else if (p.Id() == C__star.Id()) /* If:5 */{ 
            /* Let:6 */{ 
              var z *ClaireAny  
              /* noccur = 2 */
              var z_try04767 EID 
              z_try04767 = F_extract_pattern_any(g0458.Args.At(1-1),path)
              /* ERROR PROTECTION INSERTED (z-Result) */
              if ErrorIn(z_try04767) {Result = z_try04767
              } else {
              z = ANY(z_try04767)
              if (z != CNULL) /* If:7 */{ 
                Result = EID{Core.F_U_type(ToType(z),ToType(MakeConstantSet(CNULL).Id())).Id(),0}
                } else {
                Result = EID{CNULL,0}
                /* If-7 */} 
              }
              /* Let-6 */} 
            } else {
            Result = EID{CNULL,0}
            /* If-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (x.Isa.IsIn(C_type) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0459 *ClaireType   = ToType(x)
        /* noccur = 1 */
        Result = EID{g0459.Id(),0}
        /* Let-3 */} 
      /* If!2 */}  else if (x.Isa.IsIn(C_unbound_symbol) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0460 *ClaireUnboundSymbol   = ToUnboundSymbol(x)
        /* noccur = 1 */
        /* Let:4 */{ 
          var s *ClaireSymbol  
          /* noccur = 3 */
          var s_try04775 EID 
          s_try04775 = F_extract_symbol_any(g0460.Id())
          /* ERROR PROTECTION INSERTED (s-Result) */
          if ErrorIn(s_try04775) {Result = s_try04775
          } else {
          s = ToSymbol(OBJ(s_try04775))
          /* Let:5 */{ 
            var v *ClaireAny  
            /* noccur = 2 */
            var v_try04786 EID 
            /* Let:6 */{ 
              var z_some *ClaireAny   = CNULL
              /* noccur = 2 */
              /* For:7 */{ 
                var z *ClaireAny  
                _ = z
                v_try04786= EID{CFALSE.Id(),0}
                var z_support *ClaireList  
                var z_support_try04798 EID 
                z_support_try04798 = Core.F_enumerate_any(C_LDEF.Value)
                /* ERROR PROTECTION INSERTED (z_support-v_try04786) */
                if ErrorIn(z_support_try04798) {v_try04786 = z_support_try04798
                } else {
                z_support = ToList(OBJ(z_support_try04798))
                z_len := z_support.Length()
                for i_it := 0; i_it < z_len; i_it++ { 
                  z = z_support.At(i_it)
                  if (ANY(Core.F_CALL(C_mClaire_pname,ARGS(z.ToEID()))) == s.Id()) /* If:9 */{ 
                     /*v = v_try04786, s =EID*/
z_some = z
                    v_try04786 = z_some.ToEID()
                    break
                    /* If-9 */} 
                  }
                  /* loop-8 */} 
                /* For-7 */} 
              /* ERROR PROTECTION INSERTED (v_try04786-v_try04786) */
              if !ErrorIn(v_try04786) {
              v_try04786 = z_some.ToEID()
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (v-Result) */
            if ErrorIn(v_try04786) {Result = v_try04786
            } else {
            v = ANY(v_try04786)
            if (v != CNULL) /* If:6 */{ 
              Result = Core.F_CALL(C_range,ARGS(v.ToEID()))
              } else {
              var g0480I *ClaireBoolean  
              if (path.Isa.IsIn(C_list) == CTRUE) /* If:7 */{ 
                g0480I = Core.F__sup_integer(path.Length(),1)
                } else {
                g0480I = CFALSE
                /* If-7 */} 
              if (g0480I == CTRUE) /* If:7 */{ 
                /* Let:8 */{ 
                  var y *ClaireReference  
                  /* noccur = 2 */
                  var y_try04819 EID 
                  /* Let:9 */{ 
                    var _CL_obj *ClaireReference   = To_Reference(new(ClaireReference).Is(C_Reference))
                    /* noccur = 5 */
                    _CL_obj.Index = ToInteger(path.At(1-1)).Value
                    /* update:10 */{ 
                      var va_arg1 *ClaireReference  
                      var va_arg2 *ClaireList  
                      va_arg1 = _CL_obj
                      var va_arg2_try048211 EID 
                      va_arg2_try048211 = path.Cdr()
                      /* ERROR PROTECTION INSERTED (va_arg2-y_try04819) */
                      if ErrorIn(va_arg2_try048211) {y_try04819 = va_arg2_try048211
                      } else {
                      va_arg2 = ToList(OBJ(va_arg2_try048211))
                      /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                      va_arg1.Args = va_arg2
                      y_try04819 = EID{va_arg2.Id(),0}
                      }
                      /* update-10 */} 
                    /* ERROR PROTECTION INSERTED (y_try04819-y_try04819) */
                    if !ErrorIn(y_try04819) {
                    y_try04819 = EID{_CL_obj.Id(),0}
                    }
                    /* Let-9 */} 
                  /* ERROR PROTECTION INSERTED (y-Result) */
                  if ErrorIn(y_try04819) {Result = y_try04819
                  } else {
                  y = To_Reference(OBJ(y_try04819))
                  /* Let:9 */{ 
                    var v *ClaireVariable  
                    /* noccur = 1 */
                    /* Let:10 */{ 
                      var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
                      /* noccur = 5 */
                      _CL_obj.Pname = s
                      _CL_obj.Range = ToType(y.Id())
                      v = _CL_obj
                      /* Let-10 */} 
                    Core.F_tformat_string(MakeString("create a reference for ~S args=~S \n"),0,MakeConstantList(s.Id(),y.Args.Id()))
                    var v_gassign10 *ClaireAny  
                    var v_gassign10_try048310 EID 
                    v_gassign10_try048310 = Core.F_CALL(ToProperty(C_add.Id()),ARGS(EID{C_LDEF.Value,0},EID{v.Id(),0}))
                    /* ERROR PROTECTION INSERTED (v_gassign10-Result) */
                    if ErrorIn(v_gassign10_try048310) {Result = v_gassign10_try048310
                    } else {
                    v_gassign10 = ANY(v_gassign10_try048310)
                    C_LDEF.Value = v_gassign10
                    Result = v_gassign10.ToEID()
                    } 
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  Result = EID{C_void.Id(),0}
                  }
                  /* Let-8 */} 
                }
                /* Let-7 */} 
              } else {
              Result = EID{CNULL,0}
              /* If-6 */} 
            /* If-5 */} 
          }
          /* Let-4 */} 
        }
        /* Let-3 */} 
      /* Let-2 */} 
    } else {
    Result = EID{CNULL,0}
    /* If-1 */} 
  return Result} 

// The EID go function for: iClaire/extract_pattern @ any (throw: true) 
func E_extract_pattern_any (x EID,path EID) EID { 
  return /*(sm for iClaire/extract_pattern @ any= EID)*/ F_extract_pattern_any(ANY(x),ToList(OBJ(path)) )} 

// takes an <exp> that must belong to <type> and returns the CLAIRE type
/* {0} OPT.The go function for: iClaire/extract_type(x:any) [] */
func F_extract_type_any (x *ClaireAny ) EID { 
  var Result EID 
  C_LDEF.Value = ToType(C_any.Id()).EmptyList().Id()
  /* Let:1 */{ 
    var r *ClaireAny  
    /* noccur = 2 */
    var r_try04842 EID 
    r_try04842 = F_extract_pattern_any(x,CNIL)
    /* ERROR PROTECTION INSERTED (r-Result) */
    if ErrorIn(r_try04842) {Result = r_try04842
    } else {
    r = ANY(r_try04842)
    if (r == CNULL) /* If:2 */{ 
      Result = ToException(Core.C_general_error.Make(MakeString("[112] wrong type expression ~S").Id(),MakeConstantList(x).Id())).Close()
      } else {
      Result = r.ToEID()
      /* If-2 */} 
    }
    /* Let-1 */} 
  return Result} 

// The EID go function for: iClaire/extract_type @ any (throw: true) 
func E_extract_type_any (x EID) EID { 
  return /*(sm for iClaire/extract_type @ any= EID)*/ F_extract_type_any(ANY(x) )} 

// an item is an integer, a float, a symbol, a string or a type
/* {0} OPT.The go function for: extract_item(x:any,y:any) [] */
func F_extract_item_any (x *ClaireAny ,y *ClaireAny ) *ClaireAny  { 
  // use function body compiling 
if (((((C_integer.Id() == x.Isa.Id()) || 
            (C_float.Id() == x.Isa.Id())) || 
          (x.Isa.IsIn(C_symbol) == CTRUE)) || 
        (C_string.Id() == x.Isa.Id())) || 
      (x.Isa.IsIn(C_type) == CTRUE)) /* body If:1 */{ 
    return  x
    }  else if (x.Isa.IsIn(Core.C_global_variable) == CTRUE) /* body If:1 */{ 
    return  F_extract_item_any(ANY(Core.F_CALL(C_value,ARGS(x.ToEID()))),y)
    } else {
    return  CNULL
    /* body If-1 */} 
  } 

// The EID go function for: extract_item @ any (throw: false) 
func E_extract_item_any (x EID,y EID) EID { 
  return /*(sm for extract_item @ any= any)*/ F_extract_item_any(ANY(x),ANY(y) ).ToEID()} 

// version for X[...] which is the most complex case - note the extensibility
// patch.
/* {0} OPT.The go function for: extract_pattern_nth(l:list,path:list) [] */
func F_extract_pattern_nth_list (l *ClaireList ,path *ClaireList ) EID { 
  var Result EID 
  /* Let:1 */{ 
    var m int  = l.Length()
    /* noccur = 2 */
    /* Let:2 */{ 
      var x *ClaireAny   = l.At(1-1)
      /* noccur = 5 */
      if (m == 1) /* If:3 */{ 
        /* Let:4 */{ 
          var y *ClaireAny  
          /* noccur = 2 */
          var y_try04915 EID 
          y_try04915 = F_extract_pattern_any(l.At(1-1),CNIL)
          /* ERROR PROTECTION INSERTED (y-Result) */
          if ErrorIn(y_try04915) {Result = y_try04915
          } else {
          y = ANY(y_try04915)
          if (y == CNULL) /* If:5 */{ 
            Result = EID{CNULL,0}
            } else {
            /* Let:6 */{ 
              var _CL_obj *ClaireParam   = To_Param(new(ClaireParam).Is(C_Param))
              /* noccur = 7 */
              _CL_obj.Arg = C_array
              _CL_obj.Params = MakeConstantList(C_of.Id())
              _CL_obj.Args = MakeConstantList(MakeConstantSet(y).Id())
              Result = EID{_CL_obj.Id(),0}
              /* Let-6 */} 
            /* If-5 */} 
          }
          /* Let-4 */} 
        /* If!3 */}  else if (m == 2) /* If:3 */{ 
        if (((x == C_list.Id()) || 
              ((x == C_set.Id()) || 
                (x == C_subtype.Id()))) || 
            (C_class.Id() != x.Isa.Id())) /* If:4 */{ 
          /* Let:5 */{ 
            var y *ClaireAny  
            /* noccur = 2 */
            var y_try04926 EID 
            y_try04926 = F_extract_pattern_any(l.At(2-1),CNIL)
            /* ERROR PROTECTION INSERTED (y-Result) */
            if ErrorIn(y_try04926) {Result = y_try04926
            } else {
            y = ANY(y_try04926)
            h_index := ClEnv.Index /* Handle */
            h_base := ClEnv.Base
            if (y != CNULL) /* If:6 */{ 
              Result = Core.F_CALL(C_nth,ARGS(l.At(1-1).ToEID(),y.ToEID()))
              } else {
              Result = EID{CNULL,0}
              /* If-6 */} 
            if ErrorIn(Result){ 
              /* s=EID */ClEnv.Index = h_index
              ClEnv.Base = h_base
              Result = EID{CNULL,0}
              } 
            }
            /* Let-5 */} 
          } else {
          Result = EID{CNULL,0}
          /* If-4 */} 
        } else {
        /* Let:4 */{ 
          var l1 *ClaireAny   = l.At(2-1)
          /* noccur = 4 */
          /* Let:5 */{ 
            var l2 *ClaireAny   = ANY(Core.F_CALL(C_args,ARGS(l.At(3-1).ToEID())))
            /* noccur = 1 */
            /* Let:6 */{ 
              var l3 *ClaireList   = ToType(C_any.Id()).EmptyList()
              /* noccur = 4 */
              /* Let:7 */{ 
                var n int  = 1
                /* noccur = 6 */
                /* Let:8 */{ 
                  var g0485 int  = INT(Core.F_CALL(C_length,ARGS(l1.ToEID())))
                  /* noccur = 1 */
                  Result= EID{CFALSE.Id(),0}
                  for (n <= g0485) /* while:9 */{ 
                    var void_try10 EID 
                    _ = void_try10
                    { 
                    /* Let:10 */{ 
                      var y *ClaireAny   = ToList(l2).At(n-1)
                      /* noccur = 3 */
                      var l3_try049311 EID 
                      /* Let:11 */{ 
                        var g0494UU *ClaireAny  
                        /* noccur = 1 */
                        var g0494UU_try049512 EID 
                        if (y.Isa.IsIn(C_Set) == CTRUE) /* If:12 */{ 
                          /* Let:13 */{ 
                            var g0486 *Set   = To_Set(y)
                            /* noccur = 2 */
                            /* Let:14 */{ 
                              var v *ClaireAny  
                              /* noccur = 5 */
                              var v_try049615 EID 
                              /* Let:15 */{ 
                                var g0497UU *ClaireList  
                                /* noccur = 1 */
                                var g0497UU_try049816 EID 
                                /* Let:16 */{ 
                                  var g0499UU *ClaireAny  
                                  /* noccur = 1 */
                                  var g0499UU_try050017 EID 
                                  g0499UU_try050017 = Core.F_CALL(C_nth,ARGS(l1.ToEID(),EID{C__INT,IVAL(n)}))
                                  /* ERROR PROTECTION INSERTED (g0499UU-g0497UU_try049816) */
                                  if ErrorIn(g0499UU_try050017) {g0497UU_try049816 = g0499UU_try050017
                                  } else {
                                  g0499UU = ANY(g0499UU_try050017)
                                  g0497UU_try049816 = EID{path.Copy().AddFast(g0499UU).Id(),0}
                                  }
                                  /* Let-16 */} 
                                /* ERROR PROTECTION INSERTED (g0497UU-v_try049615) */
                                if ErrorIn(g0497UU_try049816) {v_try049615 = g0497UU_try049816
                                } else {
                                g0497UU = ToList(OBJ(g0497UU_try049816))
                                v_try049615 = F_extract_pattern_any(g0486.Args.At(1-1),g0497UU)
                                }
                                /* Let-15 */} 
                              /* ERROR PROTECTION INSERTED (v-g0494UU_try049512) */
                              if ErrorIn(v_try049615) {g0494UU_try049512 = v_try049615
                              } else {
                              v = ANY(v_try049615)
                              if (v == C_void.Id()) /* If:15 */{ 
                                g0494UU_try049512 = EID{C_any.Id(),0}
                                /* If!15 */}  else if (v.Isa.IsIn(C_Reference) == CTRUE) /* If:15 */{ 
                                /* Let:16 */{ 
                                  var g0488 *ClaireReference   = To_Reference(v)
                                  /* noccur = 1 */
                                  /* Let:17 */{ 
                                    var z *ClaireReference   = To_Reference(g0488.Copy().Id())
                                    /* noccur = 2 */
                                    z.Arg = CTRUE
                                    g0494UU_try049512 = EID{z.Id(),0}
                                    /* Let-17 */} 
                                  /* Let-16 */} 
                                } else {
                                /* Construct:16 */{ 
                                  var v_bag_arg *ClaireAny  
                                  g0494UU_try049512= EID{ToType(CEMPTY.Id()).EmptySet().Id(),0}
                                  var v_bag_arg_try050117 EID 
                                  if (v != CNULL) /* If:17 */{ 
                                    v_bag_arg_try050117 = v.ToEID()
                                    } else {
                                    v_bag_arg_try050117 = EVAL(g0486.Args.At(1-1))
                                    /* If-17 */} 
                                  /* ERROR PROTECTION INSERTED (v_bag_arg-g0494UU_try049512) */
                                  if ErrorIn(v_bag_arg_try050117) {g0494UU_try049512 = v_bag_arg_try050117
                                  } else {
                                  v_bag_arg = ANY(v_bag_arg_try050117)
                                  ToSet(OBJ(g0494UU_try049512)).AddFast(v_bag_arg)}
                                  /* Construct-16 */} 
                                /* If-15 */} 
                              }
                              /* Let-14 */} 
                            /* Let-13 */} 
                          } else {
                          /* Let:13 */{ 
                            var g0502UU *ClaireAny  
                            /* noccur = 1 */
                            var g0502UU_try050314 EID 
                            if (path.Length() != 0) /* If:14 */{ 
                              /* Let:15 */{ 
                                var g0504UU *ClaireAny  
                                /* noccur = 1 */
                                var g0504UU_try050516 EID 
                                g0504UU_try050516 = Core.F_CALL(C_nth,ARGS(l1.ToEID(),EID{C__INT,IVAL(n)}))
                                /* ERROR PROTECTION INSERTED (g0504UU-g0502UU_try050314) */
                                if ErrorIn(g0504UU_try050516) {g0502UU_try050314 = g0504UU_try050516
                                } else {
                                g0504UU = ANY(g0504UU_try050516)
                                g0502UU_try050314 = EID{path.AddFast(g0504UU).Id(),0}
                                }
                                /* Let-15 */} 
                              } else {
                              g0502UU_try050314 = EID{CFALSE.Id(),0}
                              /* If-14 */} 
                            /* ERROR PROTECTION INSERTED (g0502UU-g0494UU_try049512) */
                            if ErrorIn(g0502UU_try050314) {g0494UU_try049512 = g0502UU_try050314
                            } else {
                            g0502UU = ANY(g0502UU_try050314)
                            g0494UU_try049512 = F_extract_pattern_any(y,ToList(g0502UU))
                            }
                            /* Let-13 */} 
                          /* If-12 */} 
                        /* ERROR PROTECTION INSERTED (g0494UU-l3_try049311) */
                        if ErrorIn(g0494UU_try049512) {l3_try049311 = g0494UU_try049512
                        } else {
                        g0494UU = ANY(g0494UU_try049512)
                        l3_try049311 = EID{l3.AddFast(g0494UU).Id(),0}
                        }
                        /* Let-11 */} 
                      /* ERROR PROTECTION INSERTED (l3-void_try10) */
                      if ErrorIn(l3_try049311) {void_try10 = l3_try049311
                      } else {
                      l3 = ToList(OBJ(l3_try049311))
                      void_try10 = EID{l3.Id(),0}
                      }
                      /* Let-10 */} 
                    /* ERROR PROTECTION INSERTED (void_try10-void_try10) */
                    if ErrorIn(void_try10) {Result = void_try10
                    break
                    } else {
                    n = (n+1)
                    }
                    /* while-9 */} 
                  }
                  /* Let-8 */} 
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              if (l3.Memq(CNULL) == CTRUE) /* If:7 */{ 
                Result = EID{CNULL,0}
                } else {
                /* Let:8 */{ 
                  var _CL_obj *ClaireParam   = To_Param(new(ClaireParam).Is(C_Param))
                  /* noccur = 7 */
                  _CL_obj.Arg = ToClass(x)
                  _CL_obj.Params = ToList(l1)
                  _CL_obj.Args = l3
                  Result = EID{_CL_obj.Id(),0}
                  /* Let-8 */} 
                /* If-7 */} 
              }
              /* Let-6 */} 
            /* Let-5 */} 
          /* Let-4 */} 
        /* If-3 */} 
      /* Let-2 */} 
    /* Let-1 */} 
  return Result} 

// The EID go function for: extract_pattern_nth @ list (throw: true) 
func E_extract_pattern_nth_list (l EID,path EID) EID { 
  return /*(sm for extract_pattern_nth @ list= EID)*/ F_extract_pattern_nth_list(ToList(OBJ(l)),ToList(OBJ(path)) )} 

// we perform some pre-processing on x[l] at reading time to make evaluation easier
/* {0} OPT.The go function for: iClaire/extract_class_call(self:class,l:list) [] */
func F_extract_class_call_class (self *ClaireClass ,l *ClaireList ) EID { 
  var Result EID 
  var g0513I *ClaireBoolean  
  var g0513I_try05141 EID 
  /* and:1 */{ 
    var v_and1 *ClaireBoolean  
    
    v_and1 = MakeBoolean((self.Id() == C_list.Id()) || (self.Id() == C_set.Id()) || (self.Id() == C_subtype.Id()))
    if (v_and1 == CFALSE) {g0513I_try05141 = EID{CFALSE.Id(),0}
    } else /* arg:2 */{ 
      v_and1 = Equal(MakeInteger(l.Length()).Id(),MakeInteger(1).Id())
      if (v_and1 == CFALSE) {g0513I_try05141 = EID{CFALSE.Id(),0}
      } else /* arg:3 */{ 
        var v_and1_try05154 EID 
        /* Let:4 */{ 
          var y *ClaireAny   = l.At(1-1)
          /* noccur = 7 */
          /* Let:5 */{ 
            var z *ClaireAny  
            /* noccur = 1 */
            var z_try05166 EID 
            z_try05166 = F_extract_pattern_any(y,CNIL)
            /* ERROR PROTECTION INSERTED (z-v_and1_try05154) */
            if ErrorIn(z_try05166) {v_and1_try05154 = z_try05166
            } else {
            z = ANY(z_try05166)
            if (y.Isa.IsIn(Core.C_global_variable) == CTRUE) /* If:6 */{ 
              /* Let:7 */{ 
                var g0506 *Core.GlobalVariable   = Core.ToGlobalVariable(y)
                /* noccur = 2 */
                g0506 = Core.ToGlobalVariable(OBJ(Core.F_CALL(C_value,ARGS(l.At(1-1).ToEID()))))
                y = g0506.Id()
                /* Let-7 */} 
              /* If-6 */} 
            /* or:6 */{ 
              var v_or6 *ClaireBoolean  
              
              v_or6 = z.Isa.IsIn(C_type)
              if (v_or6 == CTRUE) {v_and1_try05154 = EID{CTRUE.Id(),0}
              } else /* or:7 */{ 
                v_or6 = Equal(self.Id(),C_subtype.Id())
                if (v_or6 == CTRUE) {v_and1_try05154 = EID{CTRUE.Id(),0}
                } else /* or:8 */{ 
                  if (y.Isa.IsIn(C_Call) == CTRUE) /* If:9 */{ 
                    /* Let:10 */{ 
                      var g0507 *Call   = To_Call(y)
                      /* noccur = 2 */
                      v_or6 = MakeBoolean((g0507.Selector.Id() != C__equal.Id()) || (g0507.Args.Length() != 2))
                      /* Let-10 */} 
                    /* If!9 */}  else if (y.Isa.IsIn(C_Tuple) == CTRUE) /* If:9 */{ 
                    v_or6 = CTRUE
                    } else {
                    v_or6 = CFALSE
                    /* If-9 */} 
                  if (v_or6 == CTRUE) {v_and1_try05154 = EID{CTRUE.Id(),0}
                  } else /* or:9 */{ 
                    v_and1_try05154 = EID{CFALSE.Id(),0}/* org-9 */} 
                  /* org-8 */} 
                /* org-7 */} 
              /* or-6 */} 
            }
            /* Let-5 */} 
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (v_and1-g0513I_try05141) */
        if ErrorIn(v_and1_try05154) {g0513I_try05141 = v_and1_try05154
        } else {
        v_and1 = ToBoolean(OBJ(v_and1_try05154))
        if (v_and1 == CFALSE) {g0513I_try05141 = EID{CFALSE.Id(),0}
        } else /* arg:4 */{ 
          g0513I_try05141 = EID{CTRUE.Id(),0}/* arg-4 */} 
        /* arg-3 */} 
      /* arg-2 */} 
    }
    /* and-1 */} 
  /* ERROR PROTECTION INSERTED (g0513I-Result) */
  if ErrorIn(g0513I_try05141) {Result = g0513I_try05141
  } else {
  g0513I = ToBoolean(OBJ(g0513I_try05141))
  if (g0513I == CTRUE) /* If:1 */{ 
    /* Let:2 */{ 
      var _CL_obj *Call   = To_Call(new(Call).Is(C_Call))
      /* noccur = 5 */
      _CL_obj.Selector = C_nth
      _CL_obj.Args = F_cons_any(self.Id(),l)
      Result = EID{_CL_obj.Id(),0}
      /* Let-2 */} 
    /* If!1 */}  else if (self.Id() == C_lambda.Id()) /* If:1 */{ 
    if ((l.Length() == 2) && 
        ((l.At(1-1).Isa.IsIn(C_Do) == CTRUE) || 
            (l.At(1-1).Isa.IsIn(C_Variable) == CTRUE))) /* If:2 */{ 
      /* Let:3 */{ 
        var lv *ClaireList  
        /* noccur = 2 */
        if (l.At(1-1).Isa.IsIn(C_Do) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var v_out *ClaireList   = ToType(CEMPTY.Id()).EmptyList()
            /* noccur = 2 */
            /* For:6 */{ 
              var v *ClaireAny  
              _ = v
              var v_support *ClaireList  
              v_support = ToList(OBJ(Core.F_CALL(C_args,ARGS(l.At(1-1).ToEID()))))
              v_len := v_support.Length()
              for i_it := 0; i_it < v_len; i_it++ { 
                v = v_support.At(i_it)
                if (v.Isa.IsIn(C_Variable) == CTRUE) /* If:8 */{ 
                  v_out.AddFast(v)
                  /* If-8 */} 
                /* loop-7 */} 
              /* For-6 */} 
            lv = v_out
            /* Let-5 */} 
          } else {
          lv = MakeConstantList(l.At(1-1))
          /* If-4 */} 
        Result = F_extract_signature_list(lv)
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        Result = F_lambda_I_list(lv,l.At(2-1))
        }
        /* Let-3 */} 
      } else {
      Result = ToException(Core.C_general_error.Make(MakeString("[113] Wrong lambda definition lambda[~S]").Id(),MakeConstantList(l.Id()).Id())).Close()
      /* If-2 */} 
    } else {
    /* Let:2 */{ 
      var l1 *ClaireList   = ToType(C_any.Id()).EmptyList()
      /* noccur = 3 */
      /* Let:3 */{ 
        var l2 *ClaireList   = ToType(C_any.Id()).EmptyList()
        /* noccur = 3 */
        /* Let:4 */{ 
          var m int  = l.Length()
          /* noccur = 1 */
          /* Let:5 */{ 
            var n int  = 1
            /* noccur = 4 */
            /* Let:6 */{ 
              var g0509 int  = m
              /* noccur = 1 */
              Result= EID{CFALSE.Id(),0}
              for (n <= g0509) /* while:7 */{ 
                var void_try8 EID 
                _ = void_try8
                { 
                /* Let:8 */{ 
                  var y *ClaireAny   = l.At(n-1)
                  /* noccur = 5 */
                  /* Let:9 */{ 
                    var p *ClaireAny   = CNULL
                    /* noccur = 4 */
                    /* Let:10 */{ 
                      var v *ClaireAny   = CNULL
                      /* noccur = 4 */
                      if (y.Isa.IsIn(C_Call) == CTRUE) /* If:11 */{ 
                        /* Let:12 */{ 
                          var g0510 *Call   = To_Call(y)
                          /* noccur = 5 */
                          if ((g0510.Selector.Id() != C__equal.Id()) || 
                              (g0510.Args.Length() != 2)) /* If:13 */{ 
                            void_try8 = ToException(Core.C_general_error.Make(MakeString("[114] Wrong parametrization ~S").Id(),MakeConstantList(g0510.Id()).Id())).Close()
                            } else {
                            void_try8 = EID{CFALSE.Id(),0}
                            /* If-13 */} 
                          /* ERROR PROTECTION INSERTED (void_try8-void_try8) */
                          if ErrorIn(void_try8) {Result = void_try8
                          break
                          } else {
                          var p_try051713 EID 
                          p_try051713 = F_make_a_property_any(g0510.Args.At(1-1))
                          /* ERROR PROTECTION INSERTED (p-void_try8) */
                          if ErrorIn(p_try051713) {void_try8 = p_try051713
                          Result = p_try051713
                          break
                          } else {
                          p = ANY(p_try051713)
                          void_try8 = p.ToEID()
                          /* Let:13 */{ 
                            var _CL_obj *Set   = To_Set(new(Set).Is(C_Set))
                            /* noccur = 3 */
                            _CL_obj.Args = MakeConstantList(g0510.Args.At(2-1))
                            v = _CL_obj.Id()
                            /* Let-13 */} 
                          void_try8 = v.ToEID()
                          }}
                          /* Let-12 */} 
                        /* If!11 */}  else if (y.Isa.IsIn(C_Vardef) == CTRUE) /* If:11 */{ 
                        /* Let:12 */{ 
                          var g0511 *Vardef   = To_Vardef(y)
                          /* noccur = 2 */
                          var p_try051813 EID 
                          p_try051813 = F_make_a_property_any(g0511.Pname.Id())
                          /* ERROR PROTECTION INSERTED (p-void_try8) */
                          if ErrorIn(p_try051813) {void_try8 = p_try051813
                          Result = p_try051813
                          break
                          } else {
                          p = ANY(p_try051813)
                          void_try8 = p.ToEID()
                          v = g0511.Range.Id()
                          void_try8 = v.ToEID()
                          }
                          /* Let-12 */} 
                        } else {
                        var p_try051912 EID 
                        p_try051912 = F_make_a_property_any(y)
                        /* ERROR PROTECTION INSERTED (p-void_try8) */
                        if ErrorIn(p_try051912) {void_try8 = p_try051912
                        Result = p_try051912
                        break
                        } else {
                        p = ANY(p_try051912)
                        void_try8 = p.ToEID()
                        v = CEMPTY.Id()
                        void_try8 = v.ToEID()
                        }
                        /* If-11 */} 
                      /* ERROR PROTECTION INSERTED (void_try8-void_try8) */
                      if ErrorIn(void_try8) {Result = void_try8
                      break
                      } else {
                      l1 = l1.AddFast(p)
                      l2 = l2.AddFast(v)
                      void_try8 = EID{l2.Id(),0}
                      }
                      /* Let-10 */} 
                    /* Let-9 */} 
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (void_try8-void_try8) */
                if ErrorIn(void_try8) {Result = void_try8
                break
                } else {
                n = (n+1)
                }
                /* while-7 */} 
              }
              /* Let-6 */} 
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          /* Let:5 */{ 
            var _CL_obj *Call   = To_Call(new(Call).Is(C_Call))
            /* noccur = 9 */
            _CL_obj.Selector = C_nth
            /* update:6 */{ 
              var va_arg1 *Call  
              var va_arg2 *ClaireList  
              va_arg1 = _CL_obj
              /* Let:7 */{ 
                var g0520UU *ClaireList  
                /* noccur = 1 */
                /* Construct:8 */{ 
                  var v_bag_arg *ClaireAny  
                  g0520UU= ToType(CEMPTY.Id()).EmptyList()
                  g0520UU.AddFast(l1.Id())
                  /* Let:9 */{ 
                    var _CL_obj *List   = To_List(new(List).Is(C_List))
                    /* noccur = 3 */
                    _CL_obj.Args = l2
                    v_bag_arg = _CL_obj.Id()
                    /* Let-9 */} 
                  g0520UU.AddFast(v_bag_arg)/* Construct-8 */} 
                va_arg2 = F_cons_any(self.Id(),g0520UU)
                /* Let-7 */} 
              /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
              va_arg1.Args = va_arg2
              /* update-6 */} 
            Result = EID{_CL_obj.Id(),0}
            /* Let-5 */} 
          }
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    /* If-1 */} 
  }
  return Result} 

// The EID go function for: iClaire/extract_class_call @ class (throw: true) 
func E_extract_class_call_class (self EID,l EID) EID { 
  return /*(sm for iClaire/extract_class_call @ class= EID)*/ F_extract_class_call_class(ToClass(OBJ(self)),ToList(OBJ(l)) )} 

// extract the range (type and/or second-order function)
// lvar is the list of arguments that will serve as second-o. args
// ldef is the list of extra type variables that are defined in the sig.
/* {0} OPT.The go function for: iClaire/extract_range(x:any,lvar:list,ldef:list) [] */
func F_extract_range_any (x *ClaireAny ,lvar *ClaireList ,ldef *ClaireList ) EID { 
  var Result EID 
  var g0523I *ClaireBoolean  
  /* Let:1 */{ 
    var g0524UU *ClaireBoolean  
    /* noccur = 1 */
    if (x.Isa.IsIn(C_Call) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0521 *Call   = To_Call(x)
        /* noccur = 2 */
        g0524UU = MakeBoolean((g0521.Selector.Id() == C_nth.Id()) && (g0521.Args.At(1-1) == C_type.Id()))
        /* Let-3 */} 
      } else {
      g0524UU = CFALSE
      /* If-2 */} 
    g0523I = g0524UU.Not
    /* Let-1 */} 
  if (g0523I == CTRUE) /* If:1 */{ 
    /* Construct:2 */{ 
      var v_bag_arg *ClaireAny  
      Result= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
      var v_bag_arg_try05253 EID 
      v_bag_arg_try05253 = F_extract_type_any(x)
      /* ERROR PROTECTION INSERTED (v_bag_arg-Result) */
      if ErrorIn(v_bag_arg_try05253) {Result = v_bag_arg_try05253
      } else {
      v_bag_arg = ANY(v_bag_arg_try05253)
      ToList(OBJ(Result)).AddFast(v_bag_arg)
      ToList(OBJ(Result)).AddFast(CEMPTY.Id())}
      /* Construct-2 */} 
    } else {
    Core.F_tformat_string(MakeString("extract the range from ~S with lval = ~S and ldedf = ~S \n"),0,MakeConstantList(x,lvar.Id(),ldef.Id()))
    /* For:2 */{ 
      var v *ClaireAny  
      _ = v
      var v_support *ClaireList  
      v_support = ldef
      v_len := v_support.Length()
      for i_it := 0; i_it < v_len; i_it++ { 
        v = v_support.At(i_it)
        /* Let:4 */{ 
          var r *ClaireReference   = To_Reference(OBJ(Core.F_CALL(C_range,ARGS(v.ToEID()))))
          /* noccur = 2 */
          /* Let:5 */{ 
            var path *ClaireList   = r.Args
            /* noccur = 2 */
            /* Let:6 */{ 
              var n int  = path.Length()
              /* noccur = 1 */
              /* Let:7 */{ 
                var y *ClaireAny   = lvar.At((r.Index+1)-1)
                /* noccur = 3 */
                /* Let:8 */{ 
                  var i int  = 1
                  /* noccur = 4 */
                  /* Let:9 */{ 
                    var g0522 int  = n
                    /* noccur = 1 */
                    for (i <= g0522) /* while:10 */{ 
                      /* Let:11 */{ 
                        var _CL_obj *Call   = To_Call(new(Call).Is(C_Call))
                        /* noccur = 5 */
                        _CL_obj.Selector = ToProperty(Core.C__at.Id())
                        _CL_obj.Args = MakeConstantList(y,path.At(i-1))
                        y = _CL_obj.Id()
                        /* Let-11 */} 
                      i = (i+1)
                      /* while-10 */} 
                    /* Let-9 */} 
                  /* Let-8 */} 
                /* Let:8 */{ 
                  var g0526UU *Call  
                  /* noccur = 1 */
                  /* Let:9 */{ 
                    var _CL_obj *Call   = To_Call(new(Call).Is(C_Call))
                    /* noccur = 5 */
                    _CL_obj.Selector = Core.C_member
                    _CL_obj.Args = MakeConstantList(y)
                    g0526UU = _CL_obj
                    /* Let-9 */} 
                  x = F_substitution_any(x,To_Variable(v),g0526UU.Id())
                  /* Let-8 */} 
                /* Let-7 */} 
              /* Let-6 */} 
            /* Let-5 */} 
          /* Let-4 */} 
        /* loop-3 */} 
      /* For-2 */} 
    /* Let:2 */{ 
      var lv2 *ClaireList   = ToType(C_any.Id()).EmptyList()
      /* noccur = 3 */
      /* For:3 */{ 
        var v *ClaireAny  
        _ = v
        var v_support *ClaireList  
        v_support = lvar
        v_len := v_support.Length()
        for i_it := 0; i_it < v_len; i_it++ { 
          v = v_support.At(i_it)
          /* Let:5 */{ 
            var v2 *ClaireVariable  
            /* noccur = 2 */
            /* Let:6 */{ 
              var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
              /* noccur = 5 */
              _CL_obj.Pname = ToSymbol(OBJ(Core.F_CALL(C_mClaire_pname,ARGS(v.ToEID()))))
              _CL_obj.Range = ToType(C_type.Id())
              v2 = _CL_obj
              /* Let-6 */} 
            lv2 = lv2.AddFast(v2.Id())
            x = F_substitution_any(x,To_Variable(v),v2.Id())
            /* Let-5 */} 
          /* loop-4 */} 
        /* For-3 */} 
      /* Let:3 */{ 
        var lb *ClaireLambda  
        /* noccur = 4 */
        var lb_try05274 EID 
        lb_try05274 = F_lambda_I_list(lv2,ToList(OBJ(Core.F_CALL(C_args,ARGS(x.ToEID())))).At(2-1))
        /* ERROR PROTECTION INSERTED (lb-Result) */
        if ErrorIn(lb_try05274) {Result = lb_try05274
        } else {
        lb = ToLambda(OBJ(lb_try05274))
        /* Let:4 */{ 
          var ur *ClaireAny   = CNULL
          /* noccur = 4 */
          /* Let:5 */{ 
            var g0528UU *ClaireList  
            /* noccur = 1 */
            /* Construct:6 */{ 
              var v_bag_arg *ClaireAny  
              g0528UU= ToType(CEMPTY.Id()).EmptyList()
              g0528UU.AddFast(lb.Id())
              /* Iteration:7 */{ 
                var v_list7 *ClaireList  
                var v *ClaireAny  
                var v_local7 *ClaireAny  
                v_list7 = lvar
                v_bag_arg = CreateList(ToType(CEMPTY.Id()),v_list7.Length()).Id()
                for CLcount := 0; CLcount < v_list7.Length(); CLcount++{ 
                  v = v_list7.At(CLcount)
                  v_local7 = ANY(Core.F_CALL(C_range,ARGS(v.ToEID())))
                  ToList(v_bag_arg).PutAt(CLcount,v_local7)
                  } 
                /* Iteration-7 */} 
              g0528UU.AddFast(v_bag_arg)/* Construct-6 */} 
            Core.F_tformat_string(MakeString("extract range applies type lambda ~S to arg list ~S \n"),0,g0528UU)
            /* Let-5 */} 
          h_index := ClEnv.Index /* Handle */
          h_base := ClEnv.Base
          var ur_try05295 EID 
          /* Let:5 */{ 
            var g0530UU *ClaireList  
            /* noccur = 1 */
            /* Iteration:6 */{ 
              var v_list6 *ClaireList  
              var v *ClaireAny  
              var v_local6 *ClaireAny  
              v_list6 = lvar
              g0530UU = CreateList(ToType(CEMPTY.Id()),v_list6.Length())
              for CLcount := 0; CLcount < v_list6.Length(); CLcount++{ 
                v = v_list6.At(CLcount)
                v_local6 = ANY(Core.F_CALL(C_range,ARGS(v.ToEID())))
                g0530UU.PutAt(CLcount,v_local6)
                } 
              /* Iteration-6 */} 
            ur_try05295 = F_apply_lambda(lb,g0530UU)
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (ur-Result) */
          if ErrorIn(ur_try05295) {Result = ur_try05295
          } else {
          ur = ANY(ur_try05295)
          Result = ur.ToEID()
          }
          if ErrorIn(Result){ 
            /* s=EID */ClEnv.Index = h_index
            ClEnv.Base = h_base
            PRINC("The type expression ")
            Result = Core.F_CALL(C_print,ARGS(x.ToEID()))
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC(" is not valid ... \n")
            Result = EVOID
            }
            {
            PRINC("context: lambda = ")
            Result = Core.F_print_any(lb.Id())
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC(", lvars = ")
            /* Let:6 */{ 
              var g0531UU *ClaireList  
              /* noccur = 1 */
              /* Iteration:7 */{ 
                var v_list7 *ClaireList  
                var v *ClaireAny  
                var v_local7 *ClaireAny  
                v_list7 = lvar
                g0531UU = CreateList(ToType(CEMPTY.Id()),v_list7.Length())
                for CLcount := 0; CLcount < v_list7.Length(); CLcount++{ 
                  v = v_list7.At(CLcount)
                  v_local7 = ANY(Core.F_CALL(C_range,ARGS(v.ToEID())))
                  g0531UU.PutAt(CLcount,v_local7)
                  } 
                /* Iteration-7 */} 
              Result = Core.F_print_any(g0531UU.Id())
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC("\n")
            Result = EVOID
            }}
            {
            Result = ClEnv.Exception_I.Close()
            }}
            } 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          if (ur.Isa.IsIn(C_type) != CTRUE) /* If:5 */{ 
            Result = ToException(Core.C_general_error.Make(MakeString("[115] the (resulting) range ~S is not a type").Id(),MakeConstantList(ur).Id())).Close()
            } else {
            Result = EID{CFALSE.Id(),0}
            /* If-5 */} 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          Result = EID{MakeConstantList(ur,lb.Id()).Id(),0}
          }}
          /* Let-4 */} 
        }
        /* Let-3 */} 
      /* Let-2 */} 
    /* If-1 */} 
  return Result} 

// The EID go function for: iClaire/extract_range @ any (throw: true) 
func E_extract_range_any (x EID,lvar EID,ldef EID) EID { 
  return /*(sm for iClaire/extract_range @ any= EID)*/ F_extract_range_any(ANY(x),ToList(OBJ(lvar)),ToList(OBJ(ldef)) )} 

// create a bitvector from a list of flags
/* {0} OPT.The go function for: bit_vector(l:listargs) [] */
func F_bit_vector_listargs2 (l *ClaireList ) EID { 
  var Result EID 
  /* Let:1 */{ 
    var d int  = 0
    /* noccur = 3 */
    /* For:2 */{ 
      var x *ClaireAny  
      _ = x
      Result= EID{CFALSE.Id(),0}
      var x_support *ClaireList  
      x_support = ToList(l.Id())
      x_len := x_support.Length()
      for i_it := 0; i_it < x_len; i_it++ { 
        x = x_support.At(i_it)
        var void_try4 EID 
        _ = void_try4
        var d_try05324 EID 
        /* Let:4 */{ 
          var g0533UU int 
          /* noccur = 1 */
          var g0533UU_try05345 EID 
          g0533UU_try05345 = F__exp2_integer(ToInteger(x).Value)
          /* ERROR PROTECTION INSERTED (g0533UU-d_try05324) */
          if ErrorIn(g0533UU_try05345) {d_try05324 = g0533UU_try05345
          } else {
          g0533UU = INT(g0533UU_try05345)
          d_try05324 = EID{C__INT,IVAL((d+g0533UU))}
          }
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (d-Result) */
        if ErrorIn(d_try05324) {Result = d_try05324
        Result = d_try05324
        break
        } else {
        d = INT(d_try05324)
        void_try4 = EID{C__INT,IVAL(d)}
        }
        /* loop-3 */} 
      /* For-2 */} 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    Result = EID{C__INT,IVAL(d)}
    }
    /* Let-1 */} 
  return Result} 

// The EID go function for: bit_vector @ listargs (throw: true) 
func E_bit_vector_listargs2 (l EID) EID { 
  return /*(sm for bit_vector @ listargs= EID)*/ F_bit_vector_listargs2(ToList(OBJ(l)) )} 

// parse the body and return (status, functional, body)
// the input is  body | (function!(f) | function!(f,s)) < | body> opt
// CLAIRE4: status is -1 : unknown, 0: no error, 1: an error may be thrown
//
/* {0} OPT.The go function for: iClaire/extract_status(x:any) [] */
func F_extract_status_any (x *ClaireAny ) EID { 
  var Result EID 
  /* Let:1 */{ 
    var s int  = -1
    /* noccur = 3 */
    /* Let:2 */{ 
      var f *ClaireAny  
      /* noccur = 6 */
      var g0540I *ClaireBoolean  
      if (x.Isa.IsIn(C_Call) == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var g0535 *Call   = To_Call(x)
          /* noccur = 1 */
          g0540I = Equal(g0535.Selector.Id(),C_function_I.Id())
          /* Let-4 */} 
        } else {
        g0540I = CFALSE
        /* If-3 */} 
      if (g0540I == CTRUE) /* If:3 */{ 
        f = x
        } else {
        f = CNULL
        /* If-3 */} 
      if (x.Isa.IsIn(C_And) == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var g0536 *And   = To_And(x)
          /* noccur = 4 */
          /* Let:5 */{ 
            var y *ClaireAny   = g0536.Args.At(1-1)
            /* noccur = 3 */
            var g0541I *ClaireBoolean  
            if (y.Isa.IsIn(C_Call) == CTRUE) /* If:6 */{ 
              /* Let:7 */{ 
                var g0537 *Call   = To_Call(y)
                /* noccur = 1 */
                g0541I = Equal(g0537.Selector.Id(),C_function_I.Id())
                /* Let-7 */} 
              } else {
              g0541I = CFALSE
              /* If-6 */} 
            if (g0541I == CTRUE) /* If:6 */{ 
              f = y
              g0536 = To_And(g0536.Args.At(2-1))
              /* If-6 */} 
            /* Let-5 */} 
          x = g0536.Id()
          /* Let-4 */} 
        /* If!3 */}  else if (x.Isa.IsIn(C_Call) == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var g0538 *Call   = To_Call(x)
          /* noccur = 3 */
          if (g0538.Selector.Id() == C_function_I.Id()) /* If:5 */{ 
            g0538 = To_Call(C_body.Id())
            /* If-5 */} 
          x = g0538.Id()
          /* Let-4 */} 
        } else {
        
        /* If-3 */} 
      if (f != CNULL) /* If:3 */{ 
        x = C_body.Id()
        if (ToList(OBJ(Core.F_CALL(C_args,ARGS(f.ToEID())))).Length() > 1) /* If:4 */{ 
          s = 1
          } else {
          s = 0
          /* If-4 */} 
        var f_try05424 EID 
        /* Let:4 */{ 
          var g0543UU *ClaireString  
          /* noccur = 1 */
          var g0543UU_try05445 EID 
          /* Let:5 */{ 
            var g0545UU *ClaireSymbol  
            /* noccur = 1 */
            var g0545UU_try05466 EID 
            g0545UU_try05466 = F_extract_symbol_any(ToList(OBJ(Core.F_CALL(C_args,ARGS(f.ToEID())))).At(1-1))
            /* ERROR PROTECTION INSERTED (g0545UU-g0543UU_try05445) */
            if ErrorIn(g0545UU_try05466) {g0543UU_try05445 = g0545UU_try05466
            } else {
            g0545UU = ToSymbol(OBJ(g0545UU_try05466))
            g0543UU_try05445 = EID{g0545UU.String_I().Id(),0}
            }
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (g0543UU-f_try05424) */
          if ErrorIn(g0543UU_try05445) {f_try05424 = g0543UU_try05445
          } else {
          g0543UU = ToString(OBJ(g0543UU_try05445))
          f_try05424 = F_imported_function_string(g0543UU).ToEID()
          }
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (f-Result) */
        if ErrorIn(f_try05424) {Result = f_try05424
        } else {
        f = ANY(f_try05424)
        Result = f.ToEID()
        }
        } else {
        Result = EID{CFALSE.Id(),0}
        /* If-3 */} 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = EID{MakeConstantList(MakeInteger(s).Id(),f,x).Id(),0}
      }
      /* Let-2 */} 
    /* Let-1 */} 
  return Result} 

// The EID go function for: iClaire/extract_status @ any (throw: true) 
func E_extract_status_any (x EID) EID { 
  return /*(sm for iClaire/extract_status @ any= EID)*/ F_extract_status_any(ANY(x) )} 

// new in CLAIRE4 : create a function with a syntactic marker ! for imported
/* {0} OPT.The go function for: imported_function(s:string) [] */
func F_imported_function_string (s *ClaireString ) *ClaireFunction  { 
  // use function body compiling 
return  F_make_function_string(F_append_string(MakeString("#"),s))
  } 

// The EID go function for: imported_function @ string (throw: false) 
func E_imported_function_string (s EID) EID { 
  return /*(sm for imported_function @ string= function)*/ F_imported_function_string(ToString(OBJ(s)) ).ToEID()} 

// cleans a pattern into a type
/* {0} OPT.The go function for: iClaire/type!(x:any) [] */
func F_type_I_any (x *ClaireAny ) *ClaireType  { 
  // procedure body with s =  
var Result *ClaireType  
  if (x.Isa.IsIn(C_list) == CTRUE) /* If:1 */{ 
    /* Let:2 */{ 
      var g0547 *ClaireList   = ToList(x)
      /* noccur = 1 */
      /* Iteration:3 */{ 
        var v_list3 *ClaireList  
        var y *ClaireAny  
        var v_local3 *ClaireAny  
        v_list3 = g0547
        Result = ToType(CreateList(ToType(CEMPTY.Id()),v_list3.Length()).Id())
        for CLcount := 0; CLcount < v_list3.Length(); CLcount++{ 
          y = v_list3.At(CLcount)
          v_local3 = F_type_I_any(y).Id()
          ToList(Result.Id()).PutAt(CLcount,v_local3)
          } 
        /* Iteration-3 */} 
      /* Let-2 */} 
    /* If!1 */}  else if (x.Isa.IsIn(C_Param) == CTRUE) /* If:1 */{ 
    /* Let:2 */{ 
      var g0548 *ClaireParam   = To_Param(x)
      /* noccur = 3 */
      /* Let:3 */{ 
        var _CL_obj *ClaireParam   = To_Param(new(ClaireParam).Is(C_Param))
        /* noccur = 7 */
        _CL_obj.Arg = g0548.Arg
        _CL_obj.Params = g0548.Params
        /* update:4 */{ 
          var va_arg1 *ClaireParam  
          var va_arg2 *ClaireList  
          va_arg1 = _CL_obj
          /* Iteration:5 */{ 
            var v_list5 *ClaireList  
            var y *ClaireAny  
            var v_local5 *ClaireAny  
            v_list5 = g0548.Args
            va_arg2 = CreateList(ToType(CEMPTY.Id()),v_list5.Length())
            for CLcount := 0; CLcount < v_list5.Length(); CLcount++{ 
              y = v_list5.At(CLcount)
              v_local5 = F_type_I_any(y).Id()
              va_arg2.PutAt(CLcount,v_local5)
              } 
            /* Iteration-5 */} 
          /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
          va_arg1.Args = va_arg2
          /* update-4 */} 
        Result = ToType(_CL_obj.Id())
        /* Let-3 */} 
      /* Let-2 */} 
    /* If!1 */}  else if (x.Isa.IsIn(C_Reference) == CTRUE) /* If:1 */{ 
    Result = ToType(C_any.Id())
    /* If!1 */}  else if (x.Isa.IsIn(C_type) == CTRUE) /* If:1 */{ 
    /* Let:2 */{ 
      var g0550 *ClaireType   = ToType(x)
      /* noccur = 1 */
      Result = g0550
      /* Let-2 */} 
    } else {
    Result = ToType(C_any.Id())
    /* If-1 */} 
  return Result} 

// The EID go function for: iClaire/type! @ any (throw: false) 
func E_type_I_any (x EID) EID { 
  return EID{/*(sm for iClaire/type! @ any= type)*/ F_type_I_any(ANY(x) ).Id(),0}} 

// for instance patterns
// creates a table
// to do in later versions: use an array if direct indexed access
// in the meanwhile, arrays of float should be used with care (indexed arrays)
//
/* {0} OPT.The go function for: self_eval(self:Defarray) [] */
func (self *Defarray ) SelfEval () EID { 
  var Result EID 
  /* Let:1 */{ 
    var a *ClaireList   = self.Arg.Args
    /* noccur = 6 */
    /* Let:2 */{ 
      var ar *ClaireTable  
      /* noccur = 36 */
      var ar_try05603 EID 
      /* Let:3 */{ 
        var g0561UU *ClaireSymbol  
        /* noccur = 1 */
        var g0561UU_try05624 EID 
        g0561UU_try05624 = F_extract_symbol_any(a.At(1-1))
        /* ERROR PROTECTION INSERTED (g0561UU-ar_try05603) */
        if ErrorIn(g0561UU_try05624) {ar_try05603 = g0561UU_try05624
        } else {
        g0561UU = ToSymbol(OBJ(g0561UU_try05624))
        ar_try05603 = new(ClaireTable).IsNamed(C_table,g0561UU).ToEID()
        }
        /* Let-3 */} 
      /* ERROR PROTECTION INSERTED (ar-Result) */
      if ErrorIn(ar_try05603) {Result = ar_try05603
      } else {
      ar = ToTable(OBJ(ar_try05603))
      /* Let:3 */{ 
        var v *ClaireVariable   = To_Variable(a.At(2-1))
        /* noccur = 2 */
        /* Let:4 */{ 
          var s *ClaireTypeExpression  
          /* noccur = 11 */
          var s_try05635 EID 
          s_try05635 = F_extract_type_any(v.Range.Id())
          /* ERROR PROTECTION INSERTED (s-Result) */
          if ErrorIn(s_try05635) {Result = s_try05635
          } else {
          s = ToTypeExpression(OBJ(s_try05635))
          /* Let:5 */{ 
            var e *ClaireAny  
            /* noccur = 5 */
            var e_try05646 EID 
            /* Let:6 */{ 
              var l *ClaireList  
              /* noccur = 3 */
              var l_try05657 EID 
              l_try05657 = a.Cdr()
              /* ERROR PROTECTION INSERTED (l-e_try05646) */
              if ErrorIn(l_try05657) {e_try05646 = l_try05657
              } else {
              l = ToList(OBJ(l_try05657))
              /* Let:7 */{ 
                var b *ClaireAny  
                /* noccur = 2 */
                var b_try05668 EID 
                b_try05668 = F_lexical_build_any(self.Body,l,0)
                /* ERROR PROTECTION INSERTED (b-e_try05646) */
                if ErrorIn(b_try05668) {e_try05646 = b_try05668
                } else {
                b = ANY(b_try05668)
                var g0567I *ClaireBoolean  
                /* Let:8 */{ 
                  var g0568UU *ClaireAny  
                  /* noccur = 1 */
                  /* For:9 */{ 
                    var va *ClaireAny  
                    _ = va
                    g0568UU= CFALSE.Id()
                    var va_support *ClaireList  
                    va_support = l
                    va_len := va_support.Length()
                    for i_it := 0; i_it < va_len; i_it++ { 
                      va = va_support.At(i_it)
                      if (F_occurrence_any(b,To_Variable(va)) > 0) /* If:11 */{ 
                         /*v = g0568UU, s =any*/
g0568UU = CTRUE.Id()
                        break
                        /* If-11 */} 
                      /* loop-10 */} 
                    /* For-9 */} 
                  g0567I = F_boolean_I_any(g0568UU)
                  /* Let-8 */} 
                if (g0567I == CTRUE) /* If:8 */{ 
                  e_try05646 = F_lambda_I_list(l,b)
                  } else {
                  e_try05646 = self.Body.ToEID()
                  /* If-8 */} 
                }
                /* Let-7 */} 
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (e-Result) */
            if ErrorIn(e_try05646) {Result = e_try05646
            } else {
            e = ANY(e_try05646)
            /* Let:6 */{ 
              var d *ClaireAny  
              /* noccur = 9 */
              var d_try05697 EID 
              if (e.Isa.IsIn(C_lambda) == CTRUE) /* If:7 */{ 
                d_try05697 = EID{CNULL,0}
                } else {
                d_try05697 = EVAL(self.Body)
                /* If-7 */} 
              /* ERROR PROTECTION INSERTED (d-Result) */
              if ErrorIn(d_try05697) {Result = d_try05697
              } else {
              d = ANY(d_try05697)
              /* update:7 */{ 
                var va_arg1 *ClaireRelation  
                var va_arg2 *ClaireType  
                va_arg1 = ToRelation(ar.Id())
                var va_arg2_try05708 EID 
                va_arg2_try05708 = F_extract_pattern_any(self.SetArg,CNIL)
                /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                if ErrorIn(va_arg2_try05708) {Result = va_arg2_try05708
                } else {
                va_arg2 = ToType(OBJ(va_arg2_try05708))
                /* ---------- now we compile update range(va_arg1) := va_arg2 ------- */
                va_arg1.Range = va_arg2
                Result = EID{va_arg2.Id(),0}
                }
                /* update-7 */} 
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              if (ar.Range.Id() == CNULL) /* If:7 */{ 
                Result = ToException(Core.C_range_error.Make(C_table.Id(),self.SetArg,C_type.Id())).Close()
                } else {
                Result = EID{CFALSE.Id(),0}
                /* If-7 */} 
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              if (d != CNULL) /* If:7 */{ 
                if (ar.Range.Contains(d) != CTRUE) /* If:8 */{ 
                  Result = ToException(Core.C_range_error.Make(ar.Id(),d,ar.Range.Id())).Close()
                  } else {
                  Result = EID{CFALSE.Id(),0}
                  /* If-8 */} 
                /* If!7 */}  else if (ToType(s.Id()).Included(ToType(C_integer.Id())) == CTRUE) /* If:7 */{ 
                d = MakeInteger(0).Id()
                Result = d.ToEID()
                /* If!7 */}  else if (ToType(s.Id()).Included(ToType(C_float.Id())) == CTRUE) /* If:7 */{ 
                d = MakeFloat(0).Id()
                Result = d.ToEID()
                } else {
                Result = EID{CFALSE.Id(),0}
                /* If-7 */} 
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              v.Range = ToType(s.Id())
              Result = F_attach_comment_any(ar.Id())
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              if (ar.Range.Class_I().IsIn(C_set) == CTRUE) /* If:7 */{ 
                ar.Multivalued_ask = CTRUE
                /* If-7 */} 
              if (a.Length() == 2) /* If:7 */{ 
                ar.Domain = ToType(s.Id())
                if (s.Isa.IsIn(C_Interval) == CTRUE) /* If:8 */{ 
                  /* Let:9 */{ 
                    var g0554 *ClaireInterval   = To_Interval(s.Id())
                    /* noccur = 2 */
                    ar.Params = MakeInteger((g0554.Arg1-1)).Id()
                    ar.Graph = Core.F_make_copy_list_integer(Core.F_size_Interval(g0554),d).Id()
                    /* Let-9 */} 
                  } else {
                  ar.Params = C_any.Id()
                  ar.GraphInit()
                  /* If-8 */} 
                if (e.Isa.IsIn(C_lambda) == CTRUE) /* If:8 */{ 
                  /* Let:9 */{ 
                    var g0556 *ClaireLambda   = ToLambda(e)
                    /* noccur = 1 */
                    /* For:10 */{ 
                      var y *ClaireAny  
                      _ = y
                      Result= EID{CFALSE.Id(),0}
                      var y_support *ClaireList  
                      var y_support_try057111 EID 
                      y_support_try057111 = Core.F_enumerate_any(ar.Domain.Id())
                      /* ERROR PROTECTION INSERTED (y_support-Result) */
                      if ErrorIn(y_support_try057111) {Result = y_support_try057111
                      } else {
                      y_support = ToList(OBJ(y_support_try057111))
                      y_len := y_support.Length()
                      for i_it := 0; i_it < y_len; i_it++ { 
                        y = y_support.At(i_it)
                        var void_try12 EID 
                        _ = void_try12
                        /* Let:12 */{ 
                          var g0572UU *ClaireAny  
                          /* noccur = 1 */
                          var g0572UU_try057313 EID 
                          g0572UU_try057313 = Core.F_funcall_lambda1(g0556,y)
                          /* ERROR PROTECTION INSERTED (g0572UU-void_try12) */
                          if ErrorIn(g0572UU_try057313) {void_try12 = g0572UU_try057313
                          } else {
                          g0572UU = ANY(g0572UU_try057313)
                          void_try12 = Core.F_nth_equal_table1(ar,y,g0572UU)
                          }
                          /* Let-12 */} 
                        /* ERROR PROTECTION INSERTED (void_try12-Result) */
                        if ErrorIn(void_try12) {Result = void_try12
                        Result = void_try12
                        break
                        } else {
                        }}
                        /* loop-11 */} 
                      /* For-10 */} 
                    /* Let-9 */} 
                  } else {
                  /* update:9 */{ 
                    var va_arg1 *ClaireTable  
                    var va_arg2 *ClaireAny  
                    va_arg1 = ar
                    va_arg2 = d
                    /* ---------- now we compile update default(va_arg1) := va_arg2 ------- */
                    va_arg1.Default = va_arg2
                    Result = va_arg2.ToEID()
                    /* update-9 */} 
                  /* If-8 */} 
                } else {
                /* Let:8 */{ 
                  var s2 *ClaireTypeExpression  
                  /* noccur = 8 */
                  var s2_try05749 EID 
                  s2_try05749 = F_extract_type_any(To_Variable(a.At(3-1)).Range.Id())
                  /* ERROR PROTECTION INSERTED (s2-Result) */
                  if ErrorIn(s2_try05749) {Result = s2_try05749
                  } else {
                  s2 = ToTypeExpression(OBJ(s2_try05749))
                  ar.Domain = ToType(MakeConstantList(s.Id(),s2.Id()).Tuple_I().Id())
                  To_Variable(a.At(3-1)).Range = ToType(s2.Id())
                  if ((s.Isa.IsIn(C_Interval) == CTRUE) && 
                      (s2.Isa.IsIn(C_Interval) == CTRUE)) /* If:9 */{ 
                    
                    /* update:10 */{ 
                      var va_arg1 *ClaireTable  
                      var va_arg2 *ClaireAny  
                      va_arg1 = ar
                      var va_arg2_try057511 EID 
                      /* Construct:11 */{ 
                        var v_bag_arg *ClaireAny  
                        va_arg2_try057511= EID{ToType(C_integer.Id()).EmptyList().Id(),0}
                        var v_bag_arg_try057612 EID 
                        v_bag_arg_try057612 = Core.F_CALL(C_size,ARGS(EID{s2.Id(),0}))
                        /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try057511) */
                        if ErrorIn(v_bag_arg_try057612) {va_arg2_try057511 = v_bag_arg_try057612
                        } else {
                        v_bag_arg = ANY(v_bag_arg_try057612)
                        ToList(OBJ(va_arg2_try057511)).AddFast(v_bag_arg)
                        var v_bag_arg_try057712 EID 
                        /* Let:12 */{ 
                          var g0578UU int 
                          /* noccur = 1 */
                          var g0578UU_try057913 EID 
                          /* Let:13 */{ 
                            var g0580UU int 
                            /* noccur = 1 */
                            var g0580UU_try058114 EID 
                            /* Let:14 */{ 
                              var g0582UU *ClaireAny  
                              /* noccur = 1 */
                              var g0582UU_try058315 EID 
                              g0582UU_try058315 = Core.F_CALL(C_size,ARGS(EID{s2.Id(),0}))
                              /* ERROR PROTECTION INSERTED (g0582UU-g0580UU_try058114) */
                              if ErrorIn(g0582UU_try058315) {g0580UU_try058114 = g0582UU_try058315
                              } else {
                              g0582UU = ANY(g0582UU_try058315)
                              g0580UU_try058114 = EID{C__INT,IVAL((To_Interval(s.Id()).Arg1*ToInteger(g0582UU).Value))}
                              }
                              /* Let-14 */} 
                            /* ERROR PROTECTION INSERTED (g0580UU-g0578UU_try057913) */
                            if ErrorIn(g0580UU_try058114) {g0578UU_try057913 = g0580UU_try058114
                            } else {
                            g0580UU = INT(g0580UU_try058114)
                            g0578UU_try057913 = EID{C__INT,IVAL((g0580UU+To_Interval(s2.Id()).Arg1))}
                            }
                            /* Let-13 */} 
                          /* ERROR PROTECTION INSERTED (g0578UU-v_bag_arg_try057712) */
                          if ErrorIn(g0578UU_try057913) {v_bag_arg_try057712 = g0578UU_try057913
                          } else {
                          g0578UU = INT(g0578UU_try057913)
                          v_bag_arg_try057712 = EID{C__INT,IVAL((g0578UU-1))}
                          }
                          /* Let-12 */} 
                        /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try057511) */
                        if ErrorIn(v_bag_arg_try057712) {va_arg2_try057511 = v_bag_arg_try057712
                        } else {
                        v_bag_arg = ANY(v_bag_arg_try057712)
                        ToList(OBJ(va_arg2_try057511)).AddFast(v_bag_arg)}}
                        /* Construct-11 */} 
                      /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                      if ErrorIn(va_arg2_try057511) {Result = va_arg2_try057511
                      } else {
                      va_arg2 = ANY(va_arg2_try057511)
                      /* ---------- now we compile update params(va_arg1) := va_arg2 ------- */
                      va_arg1.Params = va_arg2
                      Result = va_arg2.ToEID()
                      }
                      /* update-10 */} 
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    /* update:10 */{ 
                      var va_arg1 *ClaireTable  
                      var va_arg2 *ClaireAny  
                      va_arg1 = ar
                      var va_arg2_try058411 EID 
                      /* Let:11 */{ 
                        var g0585UU int 
                        /* noccur = 1 */
                        var g0585UU_try058612 EID 
                        /* Let:12 */{ 
                          var g0587UU *ClaireAny  
                          /* noccur = 1 */
                          var g0587UU_try058913 EID 
                          g0587UU_try058913 = Core.F_CALL(C_size,ARGS(EID{s.Id(),0}))
                          /* ERROR PROTECTION INSERTED (g0587UU-g0585UU_try058612) */
                          if ErrorIn(g0587UU_try058913) {g0585UU_try058612 = g0587UU_try058913
                          } else {
                          g0587UU = ANY(g0587UU_try058913)
                          /* Let:13 */{ 
                            var g0588UU *ClaireAny  
                            /* noccur = 1 */
                            var g0588UU_try059014 EID 
                            g0588UU_try059014 = Core.F_CALL(C_size,ARGS(EID{s2.Id(),0}))
                            /* ERROR PROTECTION INSERTED (g0588UU-g0585UU_try058612) */
                            if ErrorIn(g0588UU_try059014) {g0585UU_try058612 = g0588UU_try059014
                            } else {
                            g0588UU = ANY(g0588UU_try059014)
                            g0585UU_try058612 = F_times_integer(ToInteger(g0587UU).Value,ToInteger(g0588UU).Value)
                            }
                            /* Let-13 */} 
                          }
                          /* Let-12 */} 
                        /* ERROR PROTECTION INSERTED (g0585UU-va_arg2_try058411) */
                        if ErrorIn(g0585UU_try058612) {va_arg2_try058411 = g0585UU_try058612
                        } else {
                        g0585UU = INT(g0585UU_try058612)
                        va_arg2_try058411 = EID{Core.F_make_copy_list_integer(g0585UU,d).Id(),0}
                        }
                        /* Let-11 */} 
                      /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                      if ErrorIn(va_arg2_try058411) {Result = va_arg2_try058411
                      } else {
                      va_arg2 = ANY(va_arg2_try058411)
                      /* ---------- now we compile update mClaire/graph(va_arg1) := va_arg2 ------- */
                      va_arg1.Graph = va_arg2
                      Result = va_arg2.ToEID()
                      }
                      /* update-10 */} 
                    }
                    } else {
                    
                    ar.Params = C_any.Id()
                    ar.GraphInit()
                    Result = EVOID
                    /* If-9 */} 
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  
                  if (e.Isa.IsIn(C_lambda) == CTRUE) /* If:9 */{ 
                    /* Let:10 */{ 
                      var g0558 *ClaireLambda   = ToLambda(e)
                      /* noccur = 1 */
                      /* For:11 */{ 
                        var y1 *ClaireAny  
                        _ = y1
                        Result= EID{CFALSE.Id(),0}
                        var y1_support *ClaireList  
                        var y1_support_try059112 EID 
                        y1_support_try059112 = Core.F_enumerate_any(s.Id())
                        /* ERROR PROTECTION INSERTED (y1_support-Result) */
                        if ErrorIn(y1_support_try059112) {Result = y1_support_try059112
                        } else {
                        y1_support = ToList(OBJ(y1_support_try059112))
                        y1_len := y1_support.Length()
                        for i_it := 0; i_it < y1_len; i_it++ { 
                          y1 = y1_support.At(i_it)
                          var void_try13 EID 
                          _ = void_try13
                          /* For:13 */{ 
                            var y2 *ClaireAny  
                            _ = y2
                            void_try13= EID{CFALSE.Id(),0}
                            var y2_support *ClaireList  
                            var y2_support_try059214 EID 
                            y2_support_try059214 = Core.F_enumerate_any(s2.Id())
                            /* ERROR PROTECTION INSERTED (y2_support-void_try13) */
                            if ErrorIn(y2_support_try059214) {void_try13 = y2_support_try059214
                            } else {
                            y2_support = ToList(OBJ(y2_support_try059214))
                            y2_len := y2_support.Length()
                            for i_it := 0; i_it < y2_len; i_it++ { 
                              y2 = y2_support.At(i_it)
                              var void_try15 EID 
                              _ = void_try15
                              /* Let:15 */{ 
                                var g0593UU *ClaireAny  
                                /* noccur = 1 */
                                var g0593UU_try059416 EID 
                                g0593UU_try059416 = Core.F_CALL(C_funcall,ARGS(EID{g0558.Id(),0},y1.ToEID(),y2.ToEID()))
                                /* ERROR PROTECTION INSERTED (g0593UU-void_try15) */
                                if ErrorIn(g0593UU_try059416) {void_try15 = g0593UU_try059416
                                } else {
                                g0593UU = ANY(g0593UU_try059416)
                                void_try15 = Core.F_nth_equal_table2(ar,y1,y2,g0593UU)
                                }
                                /* Let-15 */} 
                              /* ERROR PROTECTION INSERTED (void_try15-void_try13) */
                              if ErrorIn(void_try15) {void_try13 = void_try15
                              void_try13 = void_try15
                              break
                              } else {
                              }}
                              /* loop-14 */} 
                            /* For-13 */} 
                          /* ERROR PROTECTION INSERTED (void_try13-Result) */
                          if ErrorIn(void_try13) {Result = void_try13
                          Result = void_try13
                          break
                          } else {
                          }}
                          /* loop-12 */} 
                        /* For-11 */} 
                      /* Let-10 */} 
                    } else {
                    /* update:10 */{ 
                      var va_arg1 *ClaireTable  
                      var va_arg2 *ClaireAny  
                      va_arg1 = ar
                      va_arg2 = d
                      /* ---------- now we compile update default(va_arg1) := va_arg2 ------- */
                      va_arg1.Default = va_arg2
                      Result = va_arg2.ToEID()
                      /* update-10 */} 
                    /* If-9 */} 
                  }
                  }
                  /* Let-8 */} 
                /* If-7 */} 
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              Result = EID{ar.Id(),0}
              }}}}}
              }
              /* Let-6 */} 
            }
            /* Let-5 */} 
          }
          /* Let-4 */} 
        /* Let-3 */} 
      }
      /* Let-2 */} 
    /* Let-1 */} 
  return Result} 

// The EID go function for: self_eval @ Defarray (throw: true) 
func E_self_eval_Defarray (self EID) EID { 
  return /*(sm for self_eval @ Defarray= EID)*/ To_Defarray(OBJ(self)).SelfEval( )} 

// The EVAL go function for: Defarray 
func EVAL_Defarray (x *ClaireAny) EID { 
   return To_Defarray(x).SelfEval()} 

// ------------------ NEW in v3.2 : definition of rules -----------------------
//
// a demon is a lambda with a name and a priority
/* {0} OPT.The go function for: self_print(self:demon) [] */
func (self *LanguageDemon ) SelfPrint ()  { 
  // procedure body with s =  
self.Pname.Princ()
  } 

// The EID go function for: self_print @ demon (throw: false) 
func E_self_print_demon (self EID) EID { 
  /*(sm for self_print @ demon= void)*/ ToLanguageDemon(OBJ(self)).SelfPrint( )
  return EVOID} 

/* {0} OPT.The go function for: funcall(self:demon,x:any,y:any) [] */
func F_funcall_demon1 (self *LanguageDemon ,x *ClaireAny ,y *ClaireAny ) EID { 
  var Result EID 
  Result = Core.F_CALL(C_funcall,ARGS(EID{self.Formula.Id(),0},x.ToEID(),y.ToEID()))
  return Result} 

// The EID go function for: funcall @ list<type_expression>(demon, any, any) (throw: true) 
func E_funcall_demon1 (self EID,x EID,y EID) EID { 
  return /*(sm for funcall @ list<type_expression>(demon, any, any)= EID)*/ F_funcall_demon1(ToLanguageDemon(OBJ(self)),ANY(x),ANY(y) )} 

/* {0} OPT.The go function for: funcall(self:demon,x:any,y:any,z:any) [] */
func F_funcall_demon2 (self *LanguageDemon ,x *ClaireAny ,y *ClaireAny ,z *ClaireAny ) EID { 
  var Result EID 
  Result = Core.F_CALL(C_funcall,ARGS(EID{self.Formula.Id(),0},
    x.ToEID(),
    y.ToEID(),
    z.ToEID()))
  return Result} 

// The EID go function for: funcall @ list<type_expression>(demon, any, any, any) (throw: true) 
func E_funcall_demon2 (self EID,x EID,y EID,z EID) EID { 
  return /*(sm for funcall @ list<type_expression>(demon, any, any, any)= EID)*/ F_funcall_demon2(ToLanguageDemon(OBJ(self)),
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
/* {0} OPT.The go function for: self_eval(self:Defrule) [] */
func (self *Defrule ) SelfEval () EID { 
  var Result EID 
  if (self.Args.At(1-1) != ClEnv.Id()) /* If:1 */{ 
    Result = Core.F_CALL(C_eval_rule,ARGS(EID{self.Id(),0}))
    } else {
    /* Let:2 */{ 
      var _Zcondition *ClaireAny   = self.Arg
      /* noccur = 2 */
      /* Let:3 */{ 
        var ru *ClaireAny   = self.Ident.Get()
        /* noccur = 5 */
        ru.Isa = C_Language_rule_object
        C_Language_rule_object.Instances.AddFast(ru)
        /* Let:4 */{ 
          var g0595 *ClaireTuple  
          /* noccur = 2 */
          var g0595_try05985 EID 
          g0595_try05985 = F_make_filter_any(_Zcondition)
          /* ERROR PROTECTION INSERTED (g0595-Result) */
          if ErrorIn(g0595_try05985) {Result = g0595_try05985
          } else {
          g0595 = ToTuple(OBJ(g0595_try05985))
          /* Let:5 */{ 
            var R *ClaireAny   = ToList(g0595.Id()).At(1-1)
            /* noccur = 12 */
            /* Let:6 */{ 
              var lvar *ClaireAny   = ToList(g0595.Id()).At(2-1)
              /* noccur = 2 */
              /* Let:7 */{ 
                var d *LanguageDemon  
                /* noccur = 2 */
                var d_try05998 EID 
                /* Let:8 */{ 
                  var g0600UU *ClaireAny  
                  /* noccur = 1 */
                  var g0600UU_try06019 EID 
                  g0600UU_try06019 = F_lexical_build_any(self.Body,ToList(lvar),0)
                  /* ERROR PROTECTION INSERTED (g0600UU-d_try05998) */
                  if ErrorIn(g0600UU_try06019) {d_try05998 = g0600UU_try06019
                  } else {
                  g0600UU = ANY(g0600UU_try06019)
                  d_try05998 = F_make_demon_relation(ToRelation(R),
                    ToSymbol(OBJ(Core.F_CALL(C_name,ARGS(ru.ToEID())))),
                    ToList(lvar),
                    _Zcondition,
                    g0600UU)
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (d-Result) */
                if ErrorIn(d_try05998) {Result = d_try05998
                } else {
                d = ToLanguageDemon(OBJ(d_try05998))
                if (C_function.Id() == Core.F_owner_any(ANY(Core.F_CALL(C_if_write,ARGS(R.ToEID())))).Id()) /* If:8 */{ 
                  Result = ToException(Core.C_general_error.Make(MakeString("cannot define a new rule on ~S which is closed").Id(),MakeConstantList(R).Id())).Close()
                  } else {
                  Result = EID{CFALSE.Id(),0}
                  /* If-8 */} 
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                Core.F_tformat_string(MakeString("we have defined a demon ~S for ~S \n"),0,MakeConstantList(d.Id(),R))
                /* Let:8 */{ 
                  var g0602UU *ClaireList  
                  /* noccur = 1 */
                  var g0602UU_try06039 EID 
                  /* Let:9 */{ 
                    var g0604UU *ClaireAny  
                    /* noccur = 1 */
                    var g0604UU_try060510 EID 
                    g0604UU_try060510 = Core.F_nth_table1(C_demons,R)
                    /* ERROR PROTECTION INSERTED (g0604UU-g0602UU_try06039) */
                    if ErrorIn(g0604UU_try060510) {g0602UU_try06039 = g0604UU_try060510
                    } else {
                    g0604UU = ANY(g0604UU_try060510)
                    g0602UU_try06039 = ToList(g0604UU).Add(d.Id())
                    }
                    /* Let-9 */} 
                  /* ERROR PROTECTION INSERTED (g0602UU-Result) */
                  if ErrorIn(g0602UU_try06039) {Result = g0602UU_try06039
                  } else {
                  g0602UU = ToList(OBJ(g0602UU_try06039))
                  Core.F_put_table(C_demons,R,g0602UU.Id())
                  Result = EVOID
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                Result = Core.F_nth_put_table(C_Language_last_rule,R,ru)
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                var g0606I *ClaireBoolean  
                var g0606I_try06078 EID 
                /* Let:8 */{ 
                  var g0608UU int 
                  /* noccur = 1 */
                  var g0608UU_try06099 EID 
                  /* Let:9 */{ 
                    var g0610UU *ClaireAny  
                    /* noccur = 1 */
                    var g0610UU_try061110 EID 
                    g0610UU_try061110 = Core.F_nth_table1(C_demons,R)
                    /* ERROR PROTECTION INSERTED (g0610UU-g0608UU_try06099) */
                    if ErrorIn(g0610UU_try061110) {g0608UU_try06099 = g0610UU_try061110
                    } else {
                    g0610UU = ANY(g0610UU_try061110)
                    g0608UU_try06099 = EID{C__INT,IVAL(ToList(g0610UU).Length())}
                    }
                    /* Let-9 */} 
                  /* ERROR PROTECTION INSERTED (g0608UU-g0606I_try06078) */
                  if ErrorIn(g0608UU_try06099) {g0606I_try06078 = g0608UU_try06099
                  } else {
                  g0608UU = INT(g0608UU_try06099)
                  g0606I_try06078 = EID{Equal(MakeInteger(g0608UU).Id(),MakeInteger(1).Id()).Id(),0}
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (g0606I-Result) */
                if ErrorIn(g0606I_try06078) {Result = g0606I_try06078
                } else {
                g0606I = ToBoolean(OBJ(g0606I_try06078))
                if (g0606I == CTRUE) /* If:8 */{ 
                  Result = F_eval_if_write_relation(ToRelation(R))
                  } else {
                  Result = EID{CFALSE.Id(),0}
                  /* If-8 */} 
                }
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                var g0612I *ClaireBoolean  
                if (R.Isa.IsIn(C_property) == CTRUE) /* If:8 */{ 
                  /* Let:9 */{ 
                    var g0597 *ClaireProperty   = ToProperty(R)
                    /* noccur = 1 */
                    g0612I = Equal(MakeInteger(g0597.Restrictions.Length()).Id(),MakeInteger(0).Id())
                    /* Let-9 */} 
                  } else {
                  g0612I = CFALSE
                  /* If-8 */} 
                if (g0612I == CTRUE) /* If:8 */{ 
                  F_eventMethod_property(ToProperty(R))
                  /* If-8 */} 
                Result = ru.ToEID()
                }}}}
                }
                /* Let-7 */} 
              /* Let-6 */} 
            /* Let-5 */} 
          }
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    /* If-1 */} 
  return Result} 

// The EID go function for: self_eval @ Defrule (throw: true) 
func E_self_eval_Defrule (self EID) EID { 
  return /*(sm for self_eval @ Defrule= EID)*/ To_Defrule(OBJ(self)).SelfEval( )} 

// The EVAL go function for: Defrule 
func EVAL_Defrule (x *ClaireAny) EID { 
   return To_Defrule(x).SelfEval()} 

// an eventMethod is a property whose unique (?) restriction is a method
/* {0} OPT.The go function for: eventMethod?(r:relation) [] */
func F_eventMethod_ask_relation2 (r *ClaireRelation ) *ClaireBoolean  { 
  // procedure body with s =  
var Result *ClaireBoolean  
  if (r.Isa.IsIn(C_property) == CTRUE) /* If:1 */{ 
    /* Let:2 */{ 
      var g0613 *ClaireProperty   = ToProperty(r.Id())
      /* noccur = 1 */
      /* Let:3 */{ 
        var g0614UU *ClaireAny  
        /* noccur = 1 */
        /* For:4 */{ 
          var x *ClaireAny  
          _ = x
          g0614UU= CFALSE.Id()
          for _,x = range(g0613.Restrictions.ValuesO())/* loop:5 */{ 
            if (C_slot.Id() == x.Isa.Id()) /* If:6 */{ 
               /*v = g0614UU, s =any*/
g0614UU = CTRUE.Id()
              break
              /* If-6 */} 
            /* loop-5 */} 
          /* For-4 */} 
        Result = Core.F_not_any(g0614UU)
        /* Let-3 */} 
      /* Let-2 */} 
    } else {
    Result = CFALSE
    /* If-1 */} 
  return Result} 

// The EID go function for: eventMethod? @ relation (throw: false) 
func E_eventMethod_ask_relation2 (r EID) EID { 
  return EID{/*(sm for eventMethod? @ relation= boolean)*/ F_eventMethod_ask_relation2(ToRelation(OBJ(r)) ).Id(),0}} 

// check that condition is either a filter or the conjunction of a filter and a 
// condition
// a filter is R(x) := y | R(x) := (y <- z) | R(x) :add y | P(x,y)
// R(x) is x.r or A[x]
// the list of variable is of length 3 if R is mono-valued
/* {0} OPT.The go function for: make_filter(cond:any) [] */
func F_make_filter_any (cond *ClaireAny ) EID { 
  var Result EID 
  /* Let:1 */{ 
    var c *ClaireAny  
    /* noccur = 17 */
    if (cond.Isa.IsIn(C_And) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0615 *And   = To_And(cond)
        /* noccur = 1 */
        c = g0615.Args.At(1-1)
        /* Let-3 */} 
      } else {
      c = cond
      /* If-2 */} 
    
    var g0621I *ClaireBoolean  
    if (c.Isa.IsIn(C_Call) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0617 *Call   = To_Call(c)
        /* noccur = 3 */
        g0621I = MakeBoolean(((g0617.Selector.Id() == Core.C_write.Id()) || 
            (g0617.Selector.Id() == C_nth_equal.Id())) && (g0617.Args.At(1-1).Isa.IsIn(C_relation) == CTRUE))
        /* Let-3 */} 
      } else {
      g0621I = CFALSE
      /* If-2 */} 
    if (g0621I == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var R *ClaireRelation   = ToRelation(ToList(OBJ(Core.F_CALL(C_args,ARGS(c.ToEID())))).At(1-1))
        /* noccur = 9 */
        /* Let:4 */{ 
          var x *ClaireVariable  
          /* noccur = 2 */
          var x_try06225 EID 
          /* Let:5 */{ 
            var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
            /* noccur = 5 */
            /* update:6 */{ 
              var va_arg1 *ClaireVariable  
              var va_arg2 *ClaireSymbol  
              va_arg1 = _CL_obj
              var va_arg2_try06237 EID 
              va_arg2_try06237 = F_extract_symbol_any(ToList(OBJ(Core.F_CALL(C_args,ARGS(c.ToEID())))).At(2-1))
              /* ERROR PROTECTION INSERTED (va_arg2-x_try06225) */
              if ErrorIn(va_arg2_try06237) {x_try06225 = va_arg2_try06237
              } else {
              va_arg2 = ToSymbol(OBJ(va_arg2_try06237))
              /* ---------- now we compile update mClaire/pname(va_arg1) := va_arg2 ------- */
              va_arg1.Pname = va_arg2
              x_try06225 = EID{va_arg2.Id(),0}
              }
              /* update-6 */} 
            /* ERROR PROTECTION INSERTED (x_try06225-x_try06225) */
            if !ErrorIn(x_try06225) {
            _CL_obj.Range = R.Domain
            x_try06225 = EID{_CL_obj.Id(),0}
            }
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (x-Result) */
          if ErrorIn(x_try06225) {Result = x_try06225
          } else {
          x = To_Variable(OBJ(x_try06225))
          /* Let:5 */{ 
            var y1 *ClaireAny   = ToList(OBJ(Core.F_CALL(C_args,ARGS(c.ToEID())))).At(3-1)
            /* noccur = 5 */
            if (R.Multivalued_ask == CTRUE) /* If:6 */{ 
              Result = ToException(Core.C_general_error.Make(MakeString("[188] wrong event filter ~S for multi-valued relation").Id(),MakeConstantList(c,R.Id()).Id())).Close()
              } else {
              Result = EID{CFALSE.Id(),0}
              /* If-6 */} 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            var g0624I *ClaireBoolean  
            if (y1.Isa.IsIn(C_Call) == CTRUE) /* If:6 */{ 
              /* Let:7 */{ 
                var g0618 *Call   = To_Call(y1)
                /* noccur = 1 */
                g0624I = Equal(g0618.Selector.Id(),C__inf_dash.Id())
                /* Let-7 */} 
              } else {
              g0624I = CFALSE
              /* If-6 */} 
            if (g0624I == CTRUE) /* If:6 */{ 
              /* Construct:7 */{ 
                var v_bag_arg *ClaireAny  
                Result= EID{MakeEmptyTuple().Id(),0}
                ToTuple(OBJ(Result)).AddFast(R.Id())
                var v_bag_arg_try06258 EID 
                /* Construct:8 */{ 
                  var v_bag_arg *ClaireAny  
                  v_bag_arg_try06258= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                  ToList(OBJ(v_bag_arg_try06258)).AddFast(x.Id())
                  var v_bag_arg_try06269 EID 
                  /* Let:9 */{ 
                    var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
                    /* noccur = 5 */
                    /* update:10 */{ 
                      var va_arg1 *ClaireVariable  
                      var va_arg2 *ClaireSymbol  
                      va_arg1 = _CL_obj
                      var va_arg2_try062711 EID 
                      va_arg2_try062711 = F_extract_symbol_any(ToList(OBJ(Core.F_CALL(C_args,ARGS(y1.ToEID())))).At(1-1))
                      /* ERROR PROTECTION INSERTED (va_arg2-v_bag_arg_try06269) */
                      if ErrorIn(va_arg2_try062711) {v_bag_arg_try06269 = va_arg2_try062711
                      } else {
                      va_arg2 = ToSymbol(OBJ(va_arg2_try062711))
                      /* ---------- now we compile update mClaire/pname(va_arg1) := va_arg2 ------- */
                      va_arg1.Pname = va_arg2
                      v_bag_arg_try06269 = EID{va_arg2.Id(),0}
                      }
                      /* update-10 */} 
                    /* ERROR PROTECTION INSERTED (v_bag_arg_try06269-v_bag_arg_try06269) */
                    if !ErrorIn(v_bag_arg_try06269) {
                    _CL_obj.Range = R.Range
                    v_bag_arg_try06269 = EID{_CL_obj.Id(),0}
                    }
                    /* Let-9 */} 
                  /* ERROR PROTECTION INSERTED (v_bag_arg-v_bag_arg_try06258) */
                  if ErrorIn(v_bag_arg_try06269) {v_bag_arg_try06258 = v_bag_arg_try06269
                  } else {
                  v_bag_arg = ANY(v_bag_arg_try06269)
                  ToList(OBJ(v_bag_arg_try06258)).AddFast(v_bag_arg)
                  var v_bag_arg_try06289 EID 
                  /* Let:9 */{ 
                    var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
                    /* noccur = 5 */
                    /* update:10 */{ 
                      var va_arg1 *ClaireVariable  
                      var va_arg2 *ClaireSymbol  
                      va_arg1 = _CL_obj
                      var va_arg2_try062911 EID 
                      va_arg2_try062911 = F_extract_symbol_any(ToList(OBJ(Core.F_CALL(C_args,ARGS(y1.ToEID())))).At(2-1))
                      /* ERROR PROTECTION INSERTED (va_arg2-v_bag_arg_try06289) */
                      if ErrorIn(va_arg2_try062911) {v_bag_arg_try06289 = va_arg2_try062911
                      } else {
                      va_arg2 = ToSymbol(OBJ(va_arg2_try062911))
                      /* ---------- now we compile update mClaire/pname(va_arg1) := va_arg2 ------- */
                      va_arg1.Pname = va_arg2
                      v_bag_arg_try06289 = EID{va_arg2.Id(),0}
                      }
                      /* update-10 */} 
                    /* ERROR PROTECTION INSERTED (v_bag_arg_try06289-v_bag_arg_try06289) */
                    if !ErrorIn(v_bag_arg_try06289) {
                    _CL_obj.Range = R.Range
                    v_bag_arg_try06289 = EID{_CL_obj.Id(),0}
                    }
                    /* Let-9 */} 
                  /* ERROR PROTECTION INSERTED (v_bag_arg-v_bag_arg_try06258) */
                  if ErrorIn(v_bag_arg_try06289) {v_bag_arg_try06258 = v_bag_arg_try06289
                  } else {
                  v_bag_arg = ANY(v_bag_arg_try06289)
                  ToList(OBJ(v_bag_arg_try06258)).AddFast(v_bag_arg)}}
                  /* Construct-8 */} 
                /* ERROR PROTECTION INSERTED (v_bag_arg-Result) */
                if ErrorIn(v_bag_arg_try06258) {Result = v_bag_arg_try06258
                } else {
                v_bag_arg = ANY(v_bag_arg_try06258)
                ToTuple(OBJ(Result)).AddFast(v_bag_arg)}
                /* Construct-7 */} 
              } else {
              /* Construct:7 */{ 
                var v_bag_arg *ClaireAny  
                Result= EID{MakeEmptyTuple().Id(),0}
                ToTuple(OBJ(Result)).AddFast(R.Id())
                var v_bag_arg_try06308 EID 
                /* Construct:8 */{ 
                  var v_bag_arg *ClaireAny  
                  v_bag_arg_try06308= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                  ToList(OBJ(v_bag_arg_try06308)).AddFast(x.Id())
                  var v_bag_arg_try06319 EID 
                  /* Let:9 */{ 
                    var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
                    /* noccur = 5 */
                    /* update:10 */{ 
                      var va_arg1 *ClaireVariable  
                      var va_arg2 *ClaireSymbol  
                      va_arg1 = _CL_obj
                      var va_arg2_try063211 EID 
                      va_arg2_try063211 = F_extract_symbol_any(y1)
                      /* ERROR PROTECTION INSERTED (va_arg2-v_bag_arg_try06319) */
                      if ErrorIn(va_arg2_try063211) {v_bag_arg_try06319 = va_arg2_try063211
                      } else {
                      va_arg2 = ToSymbol(OBJ(va_arg2_try063211))
                      /* ---------- now we compile update mClaire/pname(va_arg1) := va_arg2 ------- */
                      va_arg1.Pname = va_arg2
                      v_bag_arg_try06319 = EID{va_arg2.Id(),0}
                      }
                      /* update-10 */} 
                    /* ERROR PROTECTION INSERTED (v_bag_arg_try06319-v_bag_arg_try06319) */
                    if !ErrorIn(v_bag_arg_try06319) {
                    _CL_obj.Range = F_safeRange_relation(R)
                    v_bag_arg_try06319 = EID{_CL_obj.Id(),0}
                    }
                    /* Let-9 */} 
                  /* ERROR PROTECTION INSERTED (v_bag_arg-v_bag_arg_try06308) */
                  if ErrorIn(v_bag_arg_try06319) {v_bag_arg_try06308 = v_bag_arg_try06319
                  } else {
                  v_bag_arg = ANY(v_bag_arg_try06319)
                  ToList(OBJ(v_bag_arg_try06308)).AddFast(v_bag_arg)
                  /* Let:9 */{ 
                    var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
                    /* noccur = 5 */
                    _CL_obj.Pname = Core.F_gensym_void()
                    _CL_obj.Range = F_safeRange_relation(R)
                    v_bag_arg = _CL_obj.Id()
                    /* Let-9 */} 
                  ToList(OBJ(v_bag_arg_try06308)).AddFast(v_bag_arg)}
                  /* Construct-8 */} 
                /* ERROR PROTECTION INSERTED (v_bag_arg-Result) */
                if ErrorIn(v_bag_arg_try06308) {Result = v_bag_arg_try06308
                } else {
                v_bag_arg = ANY(v_bag_arg_try06308)
                ToTuple(OBJ(Result)).AddFast(v_bag_arg)}
                /* Construct-7 */} 
              /* If-6 */} 
            }
            /* Let-5 */} 
          }
          /* Let-4 */} 
        /* Let-3 */} 
      } else {
      var g0633I *ClaireBoolean  
      if (c.Isa.IsIn(C_Call) == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var g0619 *Call   = To_Call(c)
          /* noccur = 2 */
          g0633I = MakeBoolean((g0619.Selector.Id() == C_add.Id()) && (g0619.Args.At(1-1).Isa.IsIn(C_relation) == CTRUE))
          /* Let-4 */} 
        } else {
        g0633I = CFALSE
        /* If-3 */} 
      if (g0633I == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var R *ClaireRelation   = ToRelation(ToList(OBJ(Core.F_CALL(C_args,ARGS(c.ToEID())))).At(1-1))
          /* noccur = 3 */
          /* Let:5 */{ 
            var x *ClaireVariable  
            /* noccur = 1 */
            var x_try06346 EID 
            /* Let:6 */{ 
              var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
              /* noccur = 5 */
              /* update:7 */{ 
                var va_arg1 *ClaireVariable  
                var va_arg2 *ClaireSymbol  
                va_arg1 = _CL_obj
                var va_arg2_try06358 EID 
                va_arg2_try06358 = F_extract_symbol_any(ToList(OBJ(Core.F_CALL(C_args,ARGS(c.ToEID())))).At(2-1))
                /* ERROR PROTECTION INSERTED (va_arg2-x_try06346) */
                if ErrorIn(va_arg2_try06358) {x_try06346 = va_arg2_try06358
                } else {
                va_arg2 = ToSymbol(OBJ(va_arg2_try06358))
                /* ---------- now we compile update mClaire/pname(va_arg1) := va_arg2 ------- */
                va_arg1.Pname = va_arg2
                x_try06346 = EID{va_arg2.Id(),0}
                }
                /* update-7 */} 
              /* ERROR PROTECTION INSERTED (x_try06346-x_try06346) */
              if !ErrorIn(x_try06346) {
              _CL_obj.Range = R.Domain
              x_try06346 = EID{_CL_obj.Id(),0}
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (x-Result) */
            if ErrorIn(x_try06346) {Result = x_try06346
            } else {
            x = To_Variable(OBJ(x_try06346))
            /* Let:6 */{ 
              var y *ClaireVariable  
              /* noccur = 1 */
              var y_try06367 EID 
              /* Let:7 */{ 
                var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
                /* noccur = 5 */
                /* update:8 */{ 
                  var va_arg1 *ClaireVariable  
                  var va_arg2 *ClaireSymbol  
                  va_arg1 = _CL_obj
                  var va_arg2_try06379 EID 
                  va_arg2_try06379 = F_extract_symbol_any(ToList(OBJ(Core.F_CALL(C_args,ARGS(c.ToEID())))).At(3-1))
                  /* ERROR PROTECTION INSERTED (va_arg2-y_try06367) */
                  if ErrorIn(va_arg2_try06379) {y_try06367 = va_arg2_try06379
                  } else {
                  va_arg2 = ToSymbol(OBJ(va_arg2_try06379))
                  /* ---------- now we compile update mClaire/pname(va_arg1) := va_arg2 ------- */
                  va_arg1.Pname = va_arg2
                  y_try06367 = EID{va_arg2.Id(),0}
                  }
                  /* update-8 */} 
                /* ERROR PROTECTION INSERTED (y_try06367-y_try06367) */
                if !ErrorIn(y_try06367) {
                _CL_obj.Range = R.Range
                y_try06367 = EID{_CL_obj.Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (y-Result) */
              if ErrorIn(y_try06367) {Result = y_try06367
              } else {
              y = To_Variable(OBJ(y_try06367))
              Result = EID{MakeTuple(R.Id(),MakeConstantList(x.Id(),y.Id()).Id()).Id(),0}
              }
              /* Let-6 */} 
            }
            /* Let-5 */} 
          /* Let-4 */} 
        } else {
        var g0638I *ClaireBoolean  
        if (c.Isa.IsIn(C_Call) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var g0620 *Call   = To_Call(c)
            /* noccur = 1 */
            g0638I = Equal(MakeInteger(g0620.Args.Length()).Id(),MakeInteger(2).Id())
            /* Let-5 */} 
          } else {
          g0638I = CFALSE
          /* If-4 */} 
        if (g0638I == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var R *ClaireProperty   = To_Call(c).Selector
            /* noccur = 3 */
            /* Let:6 */{ 
              var x *ClaireVariable  
              /* noccur = 1 */
              var x_try06397 EID 
              /* Let:7 */{ 
                var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
                /* noccur = 5 */
                /* update:8 */{ 
                  var va_arg1 *ClaireVariable  
                  var va_arg2 *ClaireSymbol  
                  va_arg1 = _CL_obj
                  var va_arg2_try06409 EID 
                  va_arg2_try06409 = F_extract_symbol_any(ToList(OBJ(Core.F_CALL(C_args,ARGS(c.ToEID())))).At(1-1))
                  /* ERROR PROTECTION INSERTED (va_arg2-x_try06397) */
                  if ErrorIn(va_arg2_try06409) {x_try06397 = va_arg2_try06409
                  } else {
                  va_arg2 = ToSymbol(OBJ(va_arg2_try06409))
                  /* ---------- now we compile update mClaire/pname(va_arg1) := va_arg2 ------- */
                  va_arg1.Pname = va_arg2
                  x_try06397 = EID{va_arg2.Id(),0}
                  }
                  /* update-8 */} 
                /* ERROR PROTECTION INSERTED (x_try06397-x_try06397) */
                if !ErrorIn(x_try06397) {
                _CL_obj.Range = R.Domain
                x_try06397 = EID{_CL_obj.Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (x-Result) */
              if ErrorIn(x_try06397) {Result = x_try06397
              } else {
              x = To_Variable(OBJ(x_try06397))
              /* Let:7 */{ 
                var y *ClaireVariable  
                /* noccur = 1 */
                var y_try06418 EID 
                /* Let:8 */{ 
                  var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
                  /* noccur = 5 */
                  /* update:9 */{ 
                    var va_arg1 *ClaireVariable  
                    var va_arg2 *ClaireSymbol  
                    va_arg1 = _CL_obj
                    var va_arg2_try064210 EID 
                    va_arg2_try064210 = F_extract_symbol_any(ToList(OBJ(Core.F_CALL(C_args,ARGS(c.ToEID())))).At(2-1))
                    /* ERROR PROTECTION INSERTED (va_arg2-y_try06418) */
                    if ErrorIn(va_arg2_try064210) {y_try06418 = va_arg2_try064210
                    } else {
                    va_arg2 = ToSymbol(OBJ(va_arg2_try064210))
                    /* ---------- now we compile update mClaire/pname(va_arg1) := va_arg2 ------- */
                    va_arg1.Pname = va_arg2
                    y_try06418 = EID{va_arg2.Id(),0}
                    }
                    /* update-9 */} 
                  /* ERROR PROTECTION INSERTED (y_try06418-y_try06418) */
                  if !ErrorIn(y_try06418) {
                  _CL_obj.Range = R.Range
                  y_try06418 = EID{_CL_obj.Id(),0}
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (y-Result) */
                if ErrorIn(y_try06418) {Result = y_try06418
                } else {
                y = To_Variable(OBJ(y_try06418))
                Result = EID{MakeTuple(R.Id(),MakeConstantList(x.Id(),y.Id()).Id()).Id(),0}
                }
                /* Let-7 */} 
              }
              /* Let-6 */} 
            /* Let-5 */} 
          } else {
          Result = ToException(Core.C_general_error.Make(MakeString("[188] wrong event filter: ~S").Id(),MakeConstantList(c).Id())).Close()
          /* If-4 */} 
        /* If-3 */} 
      /* If-2 */} 
    /* Let-1 */} 
  return Result} 

// The EID go function for: make_filter @ any (throw: true) 
func E_make_filter_any (cond EID) EID { 
  return /*(sm for make_filter @ any= EID)*/ F_make_filter_any(ANY(cond) )} 

// create a demon
// notice that a demon has 3 args if R is monovalued 
/* {0} OPT.The go function for: make_demon(R:relation,n:symbol,lvar:list[Variable],cond:any,conc:any) [] */
func F_make_demon_relation (R *ClaireRelation ,n *ClaireSymbol ,lvar *ClaireList ,cond *ClaireAny ,conc *ClaireAny ) EID { 
  var Result EID 
  /* Let:1 */{ 
    var x *ClaireAny   = lvar.At(1-1)
    /* noccur = 2 */
    /* Let:2 */{ 
      var y *ClaireAny   = lvar.At(2-1)
      /* noccur = 2 */
      /* Let:3 */{ 
        var _Ztest *ClaireAny  
        /* noccur = 4 */
        /* Let:4 */{ 
          var _CL_obj *Call   = To_Call(new(Call).Is(C_Call))
          /* noccur = 5 */
          _CL_obj.Selector = ToProperty(IfThenElse((R.Multivalued_ask == CTRUE),
            C__Z.Id(),
            C__equal.Id()))
          _CL_obj.Args = MakeConstantList(y,F_readCall_relation(R,x).Id())
          _Ztest = _CL_obj.Id()
          /* Let-4 */} 
        /* Let:4 */{ 
          var _Zbody *ClaireAny   = conc
          /* noccur = 5 */
          Core.F_tformat_string(MakeString("make a demon for ~S from ~S => ~S (name = ~S) \n"),0,MakeConstantList(R.Id(),
            cond,
            conc,
            n.Id()))
          if (C_if_write.Trace_I > ClEnv.Verbose) /* If:5 */{ 
            /* Let:6 */{ 
              var _CL_obj *Do   = To_Do(new(Do).Is(C_Do))
              /* noccur = 13 */
              /* update:7 */{ 
                var va_arg1 *Do  
                var va_arg2 *ClaireList  
                va_arg1 = _CL_obj
                /* Construct:8 */{ 
                  var v_bag_arg *ClaireAny  
                  va_arg2= ToType(CEMPTY.Id()).EmptyList()
                  /* Let:9 */{ 
                    var _CL_obj *Call   = To_Call(new(Call).Is(C_Call))
                    /* noccur = 9 */
                    _CL_obj.Selector = Core.C_format
                    /* update:10 */{ 
                      var va_arg1 *Call  
                      var va_arg2 *ClaireList  
                      va_arg1 = _CL_obj
                      /* Construct:11 */{ 
                        var v_bag_arg *ClaireAny  
                        va_arg2= ToType(CEMPTY.Id()).EmptyList()
                        va_arg2.AddFast(MakeString("--- trigger ~A(~S,~S)\n").Id())
                        /* Let:12 */{ 
                          var _CL_obj *List   = To_List(new(List).Is(C_List))
                          /* noccur = 3 */
                          _CL_obj.Args = MakeConstantList((n.String_I()).Id(),x,y)
                          v_bag_arg = _CL_obj.Id()
                          /* Let-12 */} 
                        va_arg2.AddFast(v_bag_arg)/* Construct-11 */} 
                      /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                      va_arg1.Args = va_arg2
                      /* update-10 */} 
                    v_bag_arg = _CL_obj.Id()
                    /* Let-9 */} 
                  va_arg2.AddFast(v_bag_arg)
                  va_arg2.AddFast(conc)/* Construct-8 */} 
                /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                va_arg1.Args = va_arg2
                /* update-7 */} 
              conc = _CL_obj.Id()
              /* Let-6 */} 
            /* If-5 */} 
          /* Let:5 */{ 
            var _CL_obj *If   = To_If(new(If).Is(C_If))
            /* noccur = 3 */
            _CL_obj.Arg = conc
            _Zbody = _CL_obj.Id()
            /* Let-5 */} 
          if (F_eventMethod_ask_relation2(R) == CTRUE) /* If:5 */{ 
            if (cond.Isa.IsIn(C_And) == CTRUE) /* If:6 */{ 
              /* Let:7 */{ 
                var g0643 *And   = To_And(cond)
                /* noccur = 3 */
                var _Ztest_try06478 EID 
                if (g0643.Args.Length() > 2) /* If:8 */{ 
                  /* Let:9 */{ 
                    var _CL_obj *And   = To_And(new(And).Is(C_And))
                    /* noccur = 3 */
                    /* update:10 */{ 
                      var va_arg1 *And  
                      var va_arg2 *ClaireList  
                      va_arg1 = _CL_obj
                      var va_arg2_try064811 EID 
                      va_arg2_try064811 = g0643.Args.Cdr()
                      /* ERROR PROTECTION INSERTED (va_arg2-_Ztest_try06478) */
                      if ErrorIn(va_arg2_try064811) {_Ztest_try06478 = va_arg2_try064811
                      } else {
                      va_arg2 = ToList(OBJ(va_arg2_try064811))
                      /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                      va_arg1.Args = va_arg2
                      _Ztest_try06478 = EID{va_arg2.Id(),0}
                      }
                      /* update-10 */} 
                    /* ERROR PROTECTION INSERTED (_Ztest_try06478-_Ztest_try06478) */
                    if !ErrorIn(_Ztest_try06478) {
                    _Ztest_try06478 = EID{_CL_obj.Id(),0}
                    }
                    /* Let-9 */} 
                  } else {
                  _Ztest_try06478 = g0643.Args.At(2-1).ToEID()
                  /* If-8 */} 
                /* ERROR PROTECTION INSERTED (_Ztest-Result) */
                if ErrorIn(_Ztest_try06478) {Result = _Ztest_try06478
                } else {
                _Ztest = ANY(_Ztest_try06478)
                Result = _Ztest.ToEID()
                }
                /* Let-7 */} 
              } else {
              _Zbody = conc
              Result = _Zbody.ToEID()
              /* If-6 */} 
            /* If!5 */}  else if (cond.Isa.IsIn(C_And) == CTRUE) /* If:5 */{ 
            /* Let:6 */{ 
              var g0645 *And   = To_And(cond)
              /* noccur = 1 */
              var _Ztest_try06497 EID 
              /* Let:7 */{ 
                var _CL_obj *And   = To_And(new(And).Is(C_And))
                /* noccur = 3 */
                /* update:8 */{ 
                  var va_arg1 *And  
                  var va_arg2 *ClaireList  
                  va_arg1 = _CL_obj
                  var va_arg2_try06509 EID 
                  /* Let:9 */{ 
                    var g0651UU *ClaireList  
                    /* noccur = 1 */
                    var g0651UU_try065210 EID 
                    g0651UU_try065210 = g0645.Args.Cdr()
                    /* ERROR PROTECTION INSERTED (g0651UU-va_arg2_try06509) */
                    if ErrorIn(g0651UU_try065210) {va_arg2_try06509 = g0651UU_try065210
                    } else {
                    g0651UU = ToList(OBJ(g0651UU_try065210))
                    va_arg2_try06509 = EID{MakeConstantList(_Ztest).Append(g0651UU).Id(),0}
                    }
                    /* Let-9 */} 
                  /* ERROR PROTECTION INSERTED (va_arg2-_Ztest_try06497) */
                  if ErrorIn(va_arg2_try06509) {_Ztest_try06497 = va_arg2_try06509
                  } else {
                  va_arg2 = ToList(OBJ(va_arg2_try06509))
                  /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                  va_arg1.Args = va_arg2
                  _Ztest_try06497 = EID{va_arg2.Id(),0}
                  }
                  /* update-8 */} 
                /* ERROR PROTECTION INSERTED (_Ztest_try06497-_Ztest_try06497) */
                if !ErrorIn(_Ztest_try06497) {
                _Ztest_try06497 = EID{_CL_obj.Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (_Ztest-Result) */
              if ErrorIn(_Ztest_try06497) {Result = _Ztest_try06497
              } else {
              _Ztest = ANY(_Ztest_try06497)
              Result = _Ztest.ToEID()
              }
              /* Let-6 */} 
            } else {
            Result = EID{CFALSE.Id(),0}
            /* If-5 */} 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          if (_Zbody.Isa.IsIn(C_If) == CTRUE) /* If:5 */{ 
            /* Let:6 */{ 
              var g0646 *If   = To_If(_Zbody)
              /* noccur = 2 */
              g0646.Test = _Ztest
              /* Let-6 */} 
            /* If-5 */} 
          Core.F_tformat_string(MakeString("create a demon with name ~S \n"),0,MakeConstantList(n.Id()))
          /* Let:5 */{ 
            var _CL_obj *LanguageDemon   = ToLanguageDemon(new(LanguageDemon).Is(C_Language_demon))
            /* noccur = 5 */
            _CL_obj.Pname = n
            /* update:6 */{ 
              var va_arg1 *LanguageDemon  
              var va_arg2 *ClaireLambda  
              va_arg1 = _CL_obj
              var va_arg2_try06537 EID 
              va_arg2_try06537 = F_lambda_I_list(lvar,_Zbody)
              /* ERROR PROTECTION INSERTED (va_arg2-Result) */
              if ErrorIn(va_arg2_try06537) {Result = va_arg2_try06537
              } else {
              va_arg2 = ToLambda(OBJ(va_arg2_try06537))
              /* ---------- now we compile update formula(va_arg1) := va_arg2 ------- */
              va_arg1.Formula = va_arg2
              Result = EID{va_arg2.Id(),0}
              }
              /* update-6 */} 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            Result = EID{_CL_obj.Id(),0}
            }
            /* Let-5 */} 
          }
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    /* Let-1 */} 
  return Result} 

// The EID go function for: make_demon @ relation (throw: true) 
func E_make_demon_relation (R EID,n EID,lvar EID,cond EID,conc EID) EID { 
  return /*(sm for make_demon @ relation= EID)*/ F_make_demon_relation(ToRelation(OBJ(R)),
    ToSymbol(OBJ(n)),
    ToList(OBJ(lvar)),
    ANY(cond),
    ANY(conc) )} 

// cute litle guy
/* {0} OPT.The go function for: readCall(R:relation,x:any) [] */
func F_readCall_relation (R *ClaireRelation ,x *ClaireAny ) *Call  { 
  // procedure body with s =  
var Result *Call  
  if (C_table.Id() == R.Isa.Id()) /* If:1 */{ 
    /* Let:2 */{ 
      var _CL_obj *Call   = To_Call(new(Call).Is(C_Call))
      /* noccur = 5 */
      _CL_obj.Selector = C_get
      _CL_obj.Args = MakeConstantList(R.Id(),x)
      Result = _CL_obj
      /* Let-2 */} 
    } else {
    /* Let:2 */{ 
      var _CL_obj *Call_plus   = To_Call_plus(new(Call_plus).Is(C_Call_plus))
      /* noccur = 5 */
      _CL_obj.Selector = ToProperty(R.Id())
      _CL_obj.Args = MakeConstantList(x)
      Result = To_Call(_CL_obj.Id())
      /* Let-2 */} 
    /* If-1 */} 
  return Result} 

// The EID go function for: readCall @ relation (throw: false) 
func E_readCall_relation (R EID,x EID) EID { 
  return EID{/*(sm for readCall @ relation= Call)*/ F_readCall_relation(ToRelation(OBJ(R)),ANY(x) ).Id(),0}} 

// a small brother
/* {0} OPT.The go function for: putCall(R:relation,x:any,y:any) [] */
func F_putCall_relation2 (R *ClaireRelation ,x *ClaireAny ,y *ClaireAny ) *Call  { 
  // procedure body with s =  
var Result *Call  
  if (R.Multivalued_ask == CTRUE) /* If:1 */{ 
    /* Let:2 */{ 
      var _CL_obj *Call   = To_Call(new(Call).Is(C_Call))
      /* noccur = 5 */
      _CL_obj.Selector = Core.C_add_value
      _CL_obj.Args = MakeConstantList(R.Id(),x,y)
      Result = _CL_obj
      /* Let-2 */} 
    } else {
    /* Let:2 */{ 
      var _CL_obj *Call   = To_Call(new(Call).Is(C_Call))
      /* noccur = 5 */
      _CL_obj.Selector = C_put
      _CL_obj.Args = MakeConstantList(R.Id(),x,y)
      Result = _CL_obj
      /* Let-2 */} 
    /* If-1 */} 
  return Result} 

// The EID go function for: putCall @ relation (throw: false) 
func E_putCall_relation2 (R EID,x EID,y EID) EID { 
  return EID{/*(sm for putCall @ relation= Call)*/ F_putCall_relation2(ToRelation(OBJ(R)),ANY(x),ANY(y) ).Id(),0}} 

// v3.3 : find the range when we read the current value     
/* {0} OPT.The go function for: safeRange(x:relation) [] */
func F_safeRange_relation (x *ClaireRelation ) *ClaireType  { 
  // procedure body with s =  
var Result *ClaireType  
  if (x.Isa.IsIn(C_property) == CTRUE) /* If:1 */{ 
    /* Let:2 */{ 
      var g0654 *ClaireProperty   = ToProperty(x.Id())
      /* noccur = 2 */
      var g0658I *ClaireBoolean  
      /* Let:3 */{ 
        var g0659UU *ClaireAny  
        /* noccur = 1 */
        /* For:4 */{ 
          var s *ClaireAny  
          _ = s
          g0659UU= CFALSE.Id()
          for _,s = range(g0654.Restrictions.ValuesO())/* loop:5 */{ 
            var g0660I *ClaireBoolean  
            /* Let:6 */{ 
              var g0661UU *ClaireBoolean  
              /* noccur = 1 */
              /* Let:7 */{ 
                var g0662UU *ClaireBoolean  
                /* noccur = 1 */
                if (C_slot.Id() == s.Isa.Id()) /* If:8 */{ 
                  /* Let:9 */{ 
                    var g0655 *ClaireSlot   = ToSlot(s)
                    /* noccur = 2 */
                    g0662UU = g0655.Range.Contains(g0655.Default)
                    /* Let-9 */} 
                  } else {
                  g0662UU = CFALSE
                  /* If-8 */} 
                g0661UU = F_boolean_I_any(g0662UU.Id())
                /* Let-7 */} 
              g0660I = Core.F__I_equal_any(g0661UU.Id(),CTRUE.Id())
              /* Let-6 */} 
            if (g0660I == CTRUE) /* If:6 */{ 
               /*v = g0659UU, s =any*/
g0659UU = CTRUE.Id()
              break
              /* If-6 */} 
            /* loop-5 */} 
          /* For-4 */} 
        g0658I = Core.F_not_any(g0659UU)
        /* Let-3 */} 
      if (g0658I == CTRUE) /* If:3 */{ 
        Result = g0654.Range
        } else {
        Result = ToType(C_any.Id())
        /* If-3 */} 
      /* Let-2 */} 
    /* If!1 */}  else if (C_table.Id() == x.Isa.Id()) /* If:1 */{ 
    /* Let:2 */{ 
      var g0656 *ClaireTable   = ToTable(x.Id())
      /* noccur = 3 */
      if (g0656.Range.Contains(g0656.Default) == CTRUE) /* If:3 */{ 
        Result = g0656.Range
        } else {
        Result = ToType(C_any.Id())
        /* If-3 */} 
      /* Let-2 */} 
    } else {
    Result = ToType(C_any.Id())
    /* If-1 */} 
  return Result} 

// The EID go function for: safeRange @ relation (throw: false) 
func E_safeRange_relation (x EID) EID { 
  return EID{/*(sm for safeRange @ relation= type)*/ F_safeRange_relation(ToRelation(OBJ(x)) ).Id(),0}} 

// generate an if_write "daemon", only the first time, which uses
// the list in demons[R]
// the first step is to make the update (with inverse management)
/* {0} OPT.The go function for: eval_if_write(R:relation) [] */
func F_eval_if_write_relation (R *ClaireRelation ) EID { 
  var Result EID 
  /* Let:1 */{ 
    var l *ClaireAny  
    /* noccur = 1 */
    var l_try06632 EID 
    l_try06632 = Core.F_nth_table1(C_demons,R.Id())
    /* ERROR PROTECTION INSERTED (l-Result) */
    if ErrorIn(l_try06632) {Result = l_try06632
    } else {
    l = ANY(l_try06632)
    /* Let:2 */{ 
      var lvar *ClaireList   = ToLanguageDemon(ToList(l).At(1-1)).Formula.Vars
      /* noccur = 16 */
      /* Let:3 */{ 
        var dv *ClaireVariable  
        /* noccur = 2 */
        /* Let:4 */{ 
          var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
          /* noccur = 5 */
          _CL_obj.Pname = Core.F_gensym_void()
          _CL_obj.Range = ToType(C_Language_demon.Id())
          dv = _CL_obj
          /* Let-4 */} 
        /* Let:4 */{ 
          var l1 *ClaireList   = MakeList(ToType(C_any.Id()),F_putCall_relation2(R,lvar.At(1-1),lvar.At(2-1)).Id())
          /* noccur = 6 */
          /* Let:5 */{ 
            var l2 *ClaireList  
            /* noccur = 3 */
            /* Construct:6 */{ 
              var v_bag_arg *ClaireAny  
              l2= ToType(C_any.Id()).EmptyList()
              /* Let:7 */{ 
                var _CL_obj *For   = To_For(new(For).Is(C_For))
                /* noccur = 17 */
                _CL_obj.ClaireVar = dv
                /* update:8 */{ 
                  var va_arg1 *Iteration  
                  var va_arg2 *ClaireAny  
                  va_arg1 = To_Iteration(_CL_obj.Id())
                  /* Let:9 */{ 
                    var _CL_obj *Call   = To_Call(new(Call).Is(C_Call))
                    /* noccur = 5 */
                    _CL_obj.Selector = C_nth
                    _CL_obj.Args = MakeConstantList(C_demons.Id(),R.Id())
                    va_arg2 = _CL_obj.Id()
                    /* Let-9 */} 
                  /* ---------- now we compile update iClaire/set_arg(va_arg1) := va_arg2 ------- */
                  va_arg1.SetArg = va_arg2
                  /* update-8 */} 
                /* update:8 */{ 
                  var va_arg1 *Iteration  
                  var va_arg2 *ClaireAny  
                  va_arg1 = To_Iteration(_CL_obj.Id())
                  /* Let:9 */{ 
                    var _CL_obj *Call   = To_Call(new(Call).Is(C_Call))
                    /* noccur = 5 */
                    _CL_obj.Selector = C_funcall
                    _CL_obj.Args = MakeConstantList(dv.Id()).Append(lvar)
                    va_arg2 = _CL_obj.Id()
                    /* Let-9 */} 
                  /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                  va_arg1.Arg = va_arg2
                  /* update-8 */} 
                v_bag_arg = _CL_obj.Id()
                /* Let-7 */} 
              l2.AddFast(v_bag_arg)/* Construct-6 */} 
            Core.F_tformat_string(MakeString("generate a if_write demon for ~S \n"),0,MakeConstantList(R.Id()))
            /* For:6 */{ 
              var v *ClaireAny  
              _ = v
              Result= EID{CFALSE.Id(),0}
              var v_support *ClaireList  
              v_support = lvar
              v_len := v_support.Length()
              for i_it := 0; i_it < v_len; i_it++ { 
                v = v_support.At(i_it)
                var void_try8 EID 
                _ = void_try8
                void_try8 = Core.F_put_property2(C_range,ToObject(v),ToTypeExpression(OBJ(Core.F_CALL(C_range,ARGS(v.ToEID())))).Class_I().Id())
                /* ERROR PROTECTION INSERTED (void_try8-Result) */
                if ErrorIn(void_try8) {Result = void_try8
                Result = void_try8
                break
                } else {
                }
                /* loop-7 */} 
              /* For-6 */} 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            if (R.Inverse.Id() != CNULL) /* If:6 */{ 
              if (R.Multivalued_ask != CTRUE) /* If:7 */{ 
                /* Let:8 */{ 
                  var g0664UU *Call  
                  /* noccur = 1 */
                  /* Let:9 */{ 
                    var _CL_obj *Call   = To_Call(new(Call).Is(C_Call))
                    /* noccur = 5 */
                    _CL_obj.Selector = Core.C_Core_update_dash
                    _CL_obj.Args = MakeConstantList(R.Inverse.Id(),lvar.At(3-1),lvar.At(1-1))
                    g0664UU = _CL_obj
                    /* Let-9 */} 
                  l1 = l1.AddFast(g0664UU.Id())
                  /* Let-8 */} 
                /* If-7 */} 
              l1 = l1.AddFast(F_putCall_relation2(R.Inverse,lvar.At(2-1),lvar.At(1-1)).Id())
              /* If-6 */} 
            /* update:6 */{ 
              var va_arg1 *ClaireRelation  
              var va_arg2 *ClaireAny  
              va_arg1 = R
              var va_arg2_try06657 EID 
              /* Let:7 */{ 
                var g0666UU *ComplexInstruction  
                /* noccur = 1 */
                if (F_eventMethod_ask_relation2(R) == CTRUE) /* If:8 */{ 
                  /* Let:9 */{ 
                    var _CL_obj *Do   = To_Do(new(Do).Is(C_Do))
                    /* noccur = 3 */
                    _CL_obj.Args = l2
                    g0666UU = To_ComplexInstruction(_CL_obj.Id())
                    /* Let-9 */} 
                  /* If!8 */}  else if (R.Multivalued_ask == CTRUE) /* If:8 */{ 
                  /* Let:9 */{ 
                    var _CL_obj *If   = To_If(new(If).Is(C_If))
                    /* noccur = 21 */
                    /* update:10 */{ 
                      var va_arg1 *If  
                      var va_arg2 *ClaireAny  
                      va_arg1 = _CL_obj
                      /* Let:11 */{ 
                        var _CL_obj *Call   = To_Call(new(Call).Is(C_Call))
                        /* noccur = 11 */
                        _CL_obj.Selector = Core.C_not
                        /* update:12 */{ 
                          var va_arg1 *Call  
                          var va_arg2 *ClaireList  
                          va_arg1 = _CL_obj
                          /* Construct:13 */{ 
                            var v_bag_arg *ClaireAny  
                            va_arg2= ToType(CEMPTY.Id()).EmptyList()
                            /* Let:14 */{ 
                              var _CL_obj *Call   = To_Call(new(Call).Is(C_Call))
                              /* noccur = 5 */
                              _CL_obj.Selector = ToProperty(C__Z.Id())
                              _CL_obj.Args = MakeConstantList(lvar.At(2-1),F_readCall_relation(R,lvar.At(1-1)).Id())
                              v_bag_arg = _CL_obj.Id()
                              /* Let-14 */} 
                            va_arg2.AddFast(v_bag_arg)/* Construct-13 */} 
                          /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                          va_arg1.Args = va_arg2
                          /* update-12 */} 
                        va_arg2 = _CL_obj.Id()
                        /* Let-11 */} 
                      /* ---------- now we compile update iClaire/test(va_arg1) := va_arg2 ------- */
                      va_arg1.Test = va_arg2
                      /* update-10 */} 
                    /* update:10 */{ 
                      var va_arg1 *If  
                      var va_arg2 *ClaireAny  
                      va_arg1 = _CL_obj
                      /* Let:11 */{ 
                        var _CL_obj *Do   = To_Do(new(Do).Is(C_Do))
                        /* noccur = 3 */
                        _CL_obj.Args = l1.Append(l2)
                        va_arg2 = _CL_obj.Id()
                        /* Let-11 */} 
                      /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                      va_arg1.Arg = va_arg2
                      /* update-10 */} 
                    g0666UU = To_ComplexInstruction(_CL_obj.Id())
                    /* Let-9 */} 
                  } else {
                  /* Let:9 */{ 
                    var _CL_obj *Let   = To_Let(new(Let).Is(C_Let))
                    /* noccur = 22 */
                    _CL_obj.ClaireVar = To_Variable(lvar.At(3-1))
                    _CL_obj.Value = F_readCall_relation(R,lvar.At(1-1)).Id()
                    /* update:10 */{ 
                      var va_arg1 *Let  
                      var va_arg2 *ClaireAny  
                      va_arg1 = _CL_obj
                      /* Let:11 */{ 
                        var _CL_obj *If   = To_If(new(If).Is(C_If))
                        /* noccur = 15 */
                        /* update:12 */{ 
                          var va_arg1 *If  
                          var va_arg2 *ClaireAny  
                          va_arg1 = _CL_obj
                          /* Let:13 */{ 
                            var _CL_obj *Call   = To_Call(new(Call).Is(C_Call))
                            /* noccur = 5 */
                            _CL_obj.Selector = ToProperty(Core.C__I_equal.Id())
                            _CL_obj.Args = MakeConstantList(lvar.At(2-1),lvar.At(3-1))
                            va_arg2 = _CL_obj.Id()
                            /* Let-13 */} 
                          /* ---------- now we compile update iClaire/test(va_arg1) := va_arg2 ------- */
                          va_arg1.Test = va_arg2
                          /* update-12 */} 
                        /* update:12 */{ 
                          var va_arg1 *If  
                          var va_arg2 *ClaireAny  
                          va_arg1 = _CL_obj
                          /* Let:13 */{ 
                            var _CL_obj *Do   = To_Do(new(Do).Is(C_Do))
                            /* noccur = 3 */
                            _CL_obj.Args = l1.Append(l2)
                            va_arg2 = _CL_obj.Id()
                            /* Let-13 */} 
                          /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                          va_arg1.Arg = va_arg2
                          /* update-12 */} 
                        va_arg2 = _CL_obj.Id()
                        /* Let-11 */} 
                      /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                      va_arg1.Arg = va_arg2
                      /* update-10 */} 
                    g0666UU = To_ComplexInstruction(_CL_obj.Id())
                    /* Let-9 */} 
                  /* If-8 */} 
                va_arg2_try06657 = F_lambda_I_list(MakeConstantList(lvar.At(1-1),lvar.At(2-1)),g0666UU.Id())
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (va_arg2-Result) */
              if ErrorIn(va_arg2_try06657) {Result = va_arg2_try06657
              } else {
              va_arg2 = ANY(va_arg2_try06657)
              /* ---------- now we compile update if_write(va_arg1) := va_arg2 ------- */
              va_arg1.IfWrite = va_arg2
              Result = va_arg2.ToEID()
              }
              /* update-6 */} 
            }
            /* Let-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    }
    /* Let-1 */} 
  return Result} 

// The EID go function for: eval_if_write @ relation (throw: true) 
func E_eval_if_write_relation (R EID) EID { 
  return /*(sm for eval_if_write @ relation= EID)*/ F_eval_if_write_relation(ToRelation(OBJ(R)) )} 

// create a restriction (method) that will trigger an event
/* {0} OPT.The go function for: eventMethod(p:property) [] */
func F_eventMethod_property (p *ClaireProperty )  { 
  // procedure body with s =  
/* Let:1 */{ 
    var m *ClaireMethod   = F_add_method_property(p,MakeConstantList(p.Domain.Id(),p.Range.Id()),ToType(C_void.Id()),0,ToFunction(CNULL))
    /* noccur = 3 */
    m.Formula = ToLambda(p.IfWrite)
    Core.F_close_method(m)
    m.Functional = F_make_function_string(F_append_string(p.Name.String_I(),MakeString("_write")))
    /* Let-1 */} 
  } 

// The EID go function for: eventMethod @ property (throw: false) 
func E_eventMethod_property (p EID) EID { 
  /*(sm for eventMethod @ property= void)*/ F_eventMethod_property(ToProperty(OBJ(p)) )
  return EVOID} 

//
// safe pragma  - used to be defined in optimize
// this pragma tells to compile with full safety (include arithmetic checks) 
// claire/safe(x:any) : type[x] -> x
// **************************************************************************
// *     Part 5: JITO for methods                                           *
// **************************************************************************
// CLAIRE 4 reintroduced JITO : Just-In-Time Optimization
// we perform an on-the-fly optimization of lambdas through substitution (static calls)
// Jito(l:lambda) -> apply makeJito to the body (in place substitution)
/* {0} OPT.The go function for: jito(self:any) [] */
func F_Language_jito_any (self *ClaireAny ) EID { 
  var Result EID 
  if ((ToBoolean(C_VARIANT.Value) != CTRUE) || 
      (ClEnv.Debug_I >= 0)) /* If:1 */{ 
    Result = self.ToEID()
    /* If!1 */}  else if (self.Isa.IsIn(C_list) == CTRUE) /* If:1 */{ 
    /* Let:2 */{ 
      var g0667 *ClaireList   = ToList(self)
      /* noccur = 1 */
      /* For:3 */{ 
        var x *ClaireAny  
        _ = x
        Result= EID{CFALSE.Id(),0}
        var x_support *ClaireList  
        x_support = g0667
        x_len := x_support.Length()
        for i_it := 0; i_it < x_len; i_it++ { 
          x = x_support.At(i_it)
          var void_try5 EID 
          _ = void_try5
          void_try5 = F_Language_jito_any(x)
          /* ERROR PROTECTION INSERTED (void_try5-Result) */
          if ErrorIn(void_try5) {Result = void_try5
          Result = void_try5
          break
          } else {
          }
          /* loop-4 */} 
        /* For-3 */} 
      /* Let-2 */} 
    /* If!1 */}  else if (self.Isa.IsIn(C_Vardef) == CTRUE) /* If:1 */{ 
    /* Let:2 */{ 
      var g0668 *Vardef   = To_Vardef(self)
      /* noccur = 2 */
      g0668.Isa = C_Variable
      Result = EID{g0668.Id(),0}
      /* Let-2 */} 
    /* If!1 */}  else if (self.Isa.IsIn(C_lambda) == CTRUE) /* If:1 */{ 
    /* Let:2 */{ 
      var g0669 *ClaireLambda   = ToLambda(self)
      /* noccur = 2 */
      Result = F_Language_jito_any(g0669.Body)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = EID{g0669.Id(),0}
      }
      /* Let-2 */} 
    /* If!1 */}  else if (self.Isa.IsIn(C_And) == CTRUE) /* If:1 */{ 
    /* Let:2 */{ 
      var g0670 *And   = To_And(self)
      /* noccur = 1 */
      Result = F_Language_jito_any(g0670.Args.Id())
      /* Let-2 */} 
    /* If!1 */}  else if (self.Isa.IsIn(C_Or) == CTRUE) /* If:1 */{ 
    /* Let:2 */{ 
      var g0671 *Or   = To_Or(self)
      /* noccur = 1 */
      Result = F_Language_jito_any(g0671.Args.Id())
      /* Let-2 */} 
    /* If!1 */}  else if (self.Isa.IsIn(C_Call) == CTRUE) /* If:1 */{ 
    /* Let:2 */{ 
      var g0672 *Call   = To_Call(self)
      /* noccur = 1 */
      Result = g0672.MakeJito()
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = EID{CTRUE.Id(),0}
      }
      /* Let-2 */} 
    /* If!1 */}  else if (self.Isa.IsIn(C_Let) == CTRUE) /* If:1 */{ 
    /* Let:2 */{ 
      var g0673 *Let   = To_Let(self)
      /* noccur = 1 */
      Result = g0673.LetJito()
      /* Let-2 */} 
    /* If!1 */}  else if (self.Isa.IsIn(C_Assign) == CTRUE) /* If:1 */{ 
    /* Let:2 */{ 
      var g0674 *Assign   = To_Assign(self)
      /* noccur = 3 */
      if (g0674.ClaireVar.Isa.IsIn(C_Variable) != CTRUE) /* If:3 */{ 
        Result = ToException(Core.C_general_error.Make(MakeString("[101] ~S is not a variable").Id(),MakeConstantList(g0674.ClaireVar).Id())).Close()
        } else {
        Result = EID{CFALSE.Id(),0}
        /* If-3 */} 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = F_Language_jito_any(g0674.Arg)
      }
      /* Let-2 */} 
    /* If!1 */}  else if (self.Isa.IsIn(C_Gassign) == CTRUE) /* If:1 */{ 
    /* Let:2 */{ 
      var g0675 *Gassign   = To_Gassign(self)
      /* noccur = 3 */
      if (g0675.ClaireVar.Range.Contains(g0675.ClaireVar.Value) == CTRUE) /* If:3 */{ 
        Result = F_Language_jito_any(g0675.Arg)
        } else {
        Result = EID{CFALSE.Id(),0}
        /* If-3 */} 
      /* Let-2 */} 
    /* If!1 */}  else if (self.Isa.IsIn(C_Do) == CTRUE) /* If:1 */{ 
    /* Let:2 */{ 
      var g0676 *Do   = To_Do(self)
      /* noccur = 1 */
      Result = F_Language_jito_any(g0676.Args.Id())
      /* Let-2 */} 
    /* If!1 */}  else if (self.Isa.IsIn(C_If) == CTRUE) /* If:1 */{ 
    /* Let:2 */{ 
      var g0677 *If   = To_If(self)
      /* noccur = 3 */
      Result = F_Language_jito_any(g0677.Arg)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = F_Language_jito_any(g0677.Test)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = F_Language_jito_any(g0677.Other)
      }}
      /* Let-2 */} 
    /* If!1 */}  else if (self.Isa.IsIn(C_Iteration) == CTRUE) /* If:1 */{ 
    /* Let:2 */{ 
      var g0678 *Iteration   = To_Iteration(self)
      /* noccur = 3 */
      /* Let:3 */{ 
        var v *ClaireVariable   = g0678.ClaireVar
        /* noccur = 3 */
        /* Let:4 */{ 
          var s *ClaireAny   = g0678.SetArg
          /* noccur = 3 */
          /* Let:5 */{ 
            var o_ask *ClaireBoolean  
            /* noccur = 2 */
            /* and:6 */{ 
              var v_and6 *ClaireBoolean  
              
              if (s.Isa.IsIn(C_Call) == CTRUE) /* If:7 */{ 
                /* Let:8 */{ 
                  var g0679 *Call   = To_Call(s)
                  /* noccur = 1 */
                  v_and6 = Equal(g0679.Selector.Id(),C__dot_dot.Id())
                  /* Let-8 */} 
                } else {
                v_and6 = CFALSE
                /* If-7 */} 
              if (v_and6 == CFALSE) {o_ask = CFALSE
              } else /* arg:7 */{ 
                v_and6 = MakeBoolean((v.Range.Id() == CNULL))
                if (v_and6 == CFALSE) {o_ask = CFALSE
                } else /* arg:8 */{ 
                  o_ask = CTRUE/* arg-8 */} 
                /* arg-7 */} 
              /* and-6 */} 
            
            if (o_ask == CTRUE) /* If:6 */{ 
              v.Range = ToType(C_integer.Id())
              
              /* If-6 */} 
            Core.F_CALL(C_Language_ofto,ARGS(s.ToEID()))
            Result = F_Language_jito_any(g0678.Arg)
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            if (o_ask == CTRUE) /* If:6 */{ 
              /* update:7 */{ 
                var va_arg1 *ClaireVariable  
                var va_arg2 *ClaireType  
                va_arg1 = v
                va_arg2 = ToType(CNULL)
                /* ---------- now we compile update range(va_arg1) := va_arg2 ------- */
                va_arg1.Range = va_arg2
                Result = EID{va_arg2.Id(),0}
                /* update-7 */} 
              } else {
              Result = EID{CFALSE.Id(),0}
              /* If-6 */} 
            }
            /* Let-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    /* If!1 */}  else if (self.Isa.IsIn(C_Construct) == CTRUE) /* If:1 */{ 
    /* Let:2 */{ 
      var g0680 *Construct   = To_Construct(self)
      /* noccur = 1 */
      
      Result = F_Language_jito_any(g0680.Args.Id())
      /* Let-2 */} 
    /* If!1 */}  else if (self.Isa.IsIn(C_Exists) == CTRUE) /* If:1 */{ 
    /* Let:2 */{ 
      var g0681 *Exists   = To_Exists(self)
      /* noccur = 3 */
      Result = F_Language_jito_any(g0681.SetArg)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = F_Language_jito_any(g0681.Arg)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = F_Language_jito_any(g0681.Other)
      }}
      /* Let-2 */} 
    /* If!1 */}  else if (self.Isa.IsIn(C_Handle) == CTRUE) /* If:1 */{ 
    /* Let:2 */{ 
      var g0682 *ClaireHandle   = To_ClaireHandle(self)
      /* noccur = 4 */
      if (C_class.Id() != g0682.Test.Isa.Id()) /* If:3 */{ 
        Result = ToException(Core.C_general_error.Make(MakeString("syntax: [try %S] must use a class").Id(),MakeConstantList(g0682.Test).Id())).Close()
        } else {
        Result = EID{CFALSE.Id(),0}
        /* If-3 */} 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = F_Language_jito_any(g0682.Arg)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = F_Language_jito_any(g0682.Other)
      }}
      /* Let-2 */} 
    /* If!1 */}  else if (self.Isa.IsIn(C_Definition) == CTRUE) /* If:1 */{ 
    /* Let:2 */{ 
      var g0683 *Definition   = To_Definition(self)
      /* noccur = 2 */
      if (F_boolean_I_any(ANY(Core.F_CALL(C_Language_fast_definition,ARGS(EID{g0683.Arg.Id(),0})))) == CTRUE) /* If:3 */{ 
        /* update:4 */{ 
          var va_arg1 *ClaireAny  
          var va_arg2 *ClaireClass  
          va_arg1 = g0683.Id()
          va_arg2 = C_Language_DefFast
          /* ---------- now we compile update isa(va_arg1) := va_arg2 ------- */
          va_arg1.Isa = va_arg2
          Result = EID{va_arg2.Id(),0}
          /* update-4 */} 
        } else {
        Result = EID{CFALSE.Id(),0}
        /* If-3 */} 
      /* Let-2 */} 
    } else {
    Result = EID{CFALSE.Id(),0}
    /* If-1 */} 
  return Result} 

// The EID go function for: jito @ any (throw: true) 
func E_Language_jito_any (self EID) EID { 
  return /*(sm for jito @ any= EID)*/ F_Language_jito_any(ANY(self) )} 

// debug to remove, replace by jito?()
// Let is special in CLAIRE4 : we implement the implicit typing found in the compiler = to infer
// the type  from the value (when no range is given)
// Note : this is doubtful ... 
/* {0} OPT.The go function for: letJito(self:Let) [] */
func (self *Let ) LetJito () EID { 
  var Result EID 
  /* Let:1 */{ 
    var v *ClaireVariable   = self.ClaireVar
    /* noccur = 4 */
    /* Let:2 */{ 
      var x *ClaireAny   = self.Value
      /* noccur = 4 */
      /* Let:3 */{ 
        var untyped *ClaireBoolean   = MakeBoolean((v.Range.Id() == CNULL))
        /* noccur = 2 */
        
        if (untyped == CTRUE) /* If:4 */{ 
          if (x.Isa.IsIn(C_List) == CTRUE) /* If:5 */{ 
            /* Let:6 */{ 
              var t *ClaireType   = ToType(OBJ(Core.F_CALL(C_of,ARGS(x.ToEID()))))
              /* noccur = 2 */
              if (Equal(t.Id(),CEMPTY.Id()) != CTRUE) /* If:7 */{ 
                /* update:8 */{ 
                  var va_arg1 *ClaireVariable  
                  var va_arg2 *ClaireType  
                  va_arg1 = v
                  va_arg2 = Core.F_param_I_class(C_list,t)
                  /* ---------- now we compile update range(va_arg1) := va_arg2 ------- */
                  va_arg1.Range = va_arg2
                  Result = EID{va_arg2.Id(),0}
                  /* update-8 */} 
                } else {
                Result = EID{CFALSE.Id(),0}
                /* If-7 */} 
              /* Let-6 */} 
            } else {
            /* update:6 */{ 
              var va_arg1 *ClaireVariable  
              var va_arg2 *ClaireType  
              va_arg1 = v
              var va_arg2_try06857 EID 
              va_arg2_try06857 = F_static_type_any(x)
              /* ERROR PROTECTION INSERTED (va_arg2-Result) */
              if ErrorIn(va_arg2_try06857) {Result = va_arg2_try06857
              } else {
              va_arg2 = ToType(OBJ(va_arg2_try06857))
              /* ---------- now we compile update range(va_arg1) := va_arg2 ------- */
              va_arg1.Range = va_arg2
              Result = EID{va_arg2.Id(),0}
              }
              /* update-6 */} 
            /* If-5 */} 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          Result = EID{CFALSE.Id(),0}
          }
          } else {
          Result = EID{CFALSE.Id(),0}
          /* If-4 */} 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        Result = F_Language_jito_any(x)
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        Result = F_Language_jito_any(self.Arg)
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        if (untyped == CTRUE) /* If:4 */{ 
          /* update:5 */{ 
            var va_arg1 *ClaireVariable  
            var va_arg2 *ClaireType  
            va_arg1 = v
            va_arg2 = ToType(CNULL)
            /* ---------- now we compile update range(va_arg1) := va_arg2 ------- */
            va_arg1.Range = va_arg2
            Result = EID{va_arg2.Id(),0}
            /* update-5 */} 
          } else {
          Result = EID{CFALSE.Id(),0}
          /* If-4 */} 
        }}}
        /* Let-3 */} 
      /* Let-2 */} 
    /* Let-1 */} 
  return Result} 

// The EID go function for: letJito @ Let (throw: true) 
func E_Language_letJito_Let (self EID) EID { 
  return /*(sm for letJito @ Let= EID)*/ To_Let(OBJ(self)).LetJito( )} 

// we optimize statically (Call(p) -> Call_method(m)) when
//   - only one restriction match 
//   - all domains are classes => class match
//   - the only one match is a compiled method
//   - the property is static (open = 1, vs extensible) and not too many restrictions
/* {0} OPT.The go function for: makeJito(self:Call) [] */
func (self *Call ) MakeJito () EID { 
  var Result EID 
  Result = F_Language_jito_any(self.Args.Id())
  /* ERROR PROTECTION INSERTED (Result-Result) */
  if !ErrorIn(Result) {
  /* Let:1 */{ 
    var p *ClaireProperty   = self.Selector
    /* noccur = 6 */
    /* Let:2 */{ 
      var larg *ClaireList   = self.Args
      /* noccur = 2 */
      /* Let:3 */{ 
        var n int  = larg.Length()
        /* noccur = 3 */
        /* Let:4 */{ 
          var m *ClaireAny   = CNULL
          /* noccur = 4 */
          var g0690I *ClaireBoolean  
          /* and:5 */{ 
            var v_and5 *ClaireBoolean  
            
            v_and5 = Equal(p.Id(),Core.C_write.Id())
            if (v_and5 == CFALSE) {g0690I = CFALSE
            } else /* arg:6 */{ 
              /* Let:7 */{ 
                var p2 *ClaireAny   = self.Args.At(1-1)
                /* noccur = 2 */
                if (p2.Isa.IsIn(C_property) == CTRUE) /* If:8 */{ 
                  /* Let:9 */{ 
                    var g0686 *ClaireProperty   = ToProperty(p2)
                    /* noccur = 3 */
                    v_and5 = MakeBoolean((g0686.Inverse.Id() == CNULL) && (g0686.Store_ask != CTRUE) && (g0686.IfWrite == CNULL))
                    /* Let-9 */} 
                  } else {
                  v_and5 = CFALSE
                  /* If-8 */} 
                /* Let-7 */} 
              if (v_and5 == CFALSE) {g0690I = CFALSE
              } else /* arg:7 */{ 
                g0690I = CTRUE/* arg-7 */} 
              /* arg-6 */} 
            /* and-5 */} 
          if (g0690I == CTRUE) /* If:5 */{ 
            p = C_write_fast
            self.Selector = C_write_fast
            /* If-5 */} 
          var g0691I *ClaireBoolean  
          /* and:5 */{ 
            var v_and5 *ClaireBoolean  
            
            v_and5 = Core.F__inf_equal_integer(p.Open,1)
            if (v_and5 == CFALSE) {g0691I = CFALSE
            } else /* arg:6 */{ 
              v_and5 = Core.F__inf_equal_integer(p.Restrictions.Length(),10)
              if (v_and5 == CFALSE) {g0691I = CFALSE
              } else /* arg:7 */{ 
                /* Let:8 */{ 
                  var g0692UU *ClaireAny  
                  /* noccur = 1 */
                  /* For:9 */{ 
                    var x *ClaireAny  
                    _ = x
                    g0692UU= CFALSE.Id()
                    for _,x = range(p.Restrictions.ValuesO())/* loop:10 */{ 
                      var g0693I *ClaireBoolean  
                      /* Let:11 */{ 
                        var g0694UU *ClaireBoolean  
                        /* noccur = 1 */
                        /* Let:12 */{ 
                          var g0695UU *ClaireAny  
                          /* noccur = 1 */
                          /* For:13 */{ 
                            var t *ClaireAny  
                            _ = t
                            g0695UU= CFALSE.Id()
                            for _,t = range(ToRestriction(x).Domain.ValuesO())/* loop:14 */{ 
                              if (C_class.Id() != t.Isa.Id()) /* If:15 */{ 
                                 /*v = g0695UU, s =any*/
g0695UU = CTRUE.Id()
                                break
                                /* If-15 */} 
                              /* loop-14 */} 
                            /* For-13 */} 
                          g0694UU = Core.F_not_any(g0695UU)
                          /* Let-12 */} 
                        g0693I = g0694UU.Not
                        /* Let-11 */} 
                      if (g0693I == CTRUE) /* If:11 */{ 
                         /*v = g0692UU, s =any*/
g0692UU = CTRUE.Id()
                        break
                        /* If-11 */} 
                      /* loop-10 */} 
                    /* For-9 */} 
                  v_and5 = Core.F_not_any(g0692UU)
                  /* Let-8 */} 
                if (v_and5 == CFALSE) {g0691I = CFALSE
                } else /* arg:8 */{ 
                  g0691I = CTRUE/* arg-8 */} 
                /* arg-7 */} 
              /* arg-6 */} 
            /* and-5 */} 
          if (g0691I == CTRUE) /* If:5 */{ 
            /* Let:6 */{ 
              var lt *ClaireList  
              /* noccur = 1 */
              var lt_try06967 EID 
              /* Iteration:7 */{ 
                var v_list7 *ClaireList  
                var x *ClaireAny  
                var v_local7 *ClaireAny  
                v_list7 = larg
                lt_try06967 = EID{CreateList(ToType(CEMPTY.Id()),v_list7.Length()).Id(),0}
                for CLcount := 0; CLcount < v_list7.Length(); CLcount++{ 
                  x = v_list7.At(CLcount)
                  var v_local7_try06979 EID 
                  v_local7_try06979 = F_static_type_any(x)
                  /* ERROR PROTECTION INSERTED (v_local7-lt_try06967) */
                  if ErrorIn(v_local7_try06979) {lt_try06967 = v_local7_try06979
                  lt_try06967 = v_local7_try06979
                  break
                  } else {
                  v_local7 = ANY(v_local7_try06979)
                  ToList(OBJ(lt_try06967)).PutAt(CLcount,v_local7)
                  } 
                }
                /* Iteration-7 */} 
              /* ERROR PROTECTION INSERTED (lt-Result) */
              if ErrorIn(lt_try06967) {Result = lt_try06967
              } else {
              lt = ToList(OBJ(lt_try06967))
              
              /* For:7 */{ 
                var x *ClaireAny  
                _ = x
                Result= EID{CFALSE.Id(),0}
                for _,x = range(p.Definition.ValuesO())/* loop:8 */{ 
                  if (F_Language_makeCallMatch_restriction(ToRestriction(x),lt) == CTRUE) /* If:9 */{ 
                    m = x
                     /*v = Result, s =EID*/
Result = EID{CTRUE.Id(),0}
                    break
                    /* If-9 */} 
                  /* loop-8 */} 
                /* For-7 */} 
              }
              /* Let-6 */} 
            } else {
            Result = EID{CFALSE.Id(),0}
            /* If-5 */} 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          var g0698I *ClaireBoolean  
          if (C_method.Id() == m.Isa.Id()) /* If:5 */{ 
            /* Let:6 */{ 
              var g0688 *ClaireMethod   = ToMethod(m)
              /* noccur = 1 */
              g0698I = MakeBoolean((g0688.Functional.Id() == CNULL)).Not
              /* Let-6 */} 
            } else {
            g0698I = CFALSE
            /* If-5 */} 
          if (g0698I == CTRUE) /* If:5 */{ 
            /* update:6 */{ 
              var va_arg1 *ClaireAny  
              var va_arg2 *ClaireClass  
              va_arg1 = self.Id()
              if (n == 1) /* If:7 */{ 
                va_arg2 = C_Call_method1
                /* If!7 */}  else if (n == 2) /* If:7 */{ 
                va_arg2 = C_Call_method2
                /* If!7 */}  else if (n == 3) /* If:7 */{ 
                va_arg2 = C_Language_Call_method3
                } else {
                va_arg2 = C_Call_method
                /* If-7 */} 
              /* ---------- now we compile update isa(va_arg1) := va_arg2 ------- */
              va_arg1.Isa = va_arg2
              /* update-6 */} 
            /* update:6 */{ 
              var va_arg1 *CallMethod  
              var va_arg2 *ClaireMethod  
              va_arg1 = To_CallMethod(self.Id())
              va_arg2 = ToMethod(m)
              /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
              va_arg1.Arg = va_arg2
              Result = EID{va_arg2.Id(),0}
              /* update-6 */} 
            } else {
            Result = EID{CFALSE.Id(),0}
            /* If-5 */} 
          }
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    /* Let-1 */} 
  }
  return Result} 

// The EID go function for: makeJito @ Call (throw: true) 
func E_Language_makeJito_Call (self EID) EID { 
  return /*(sm for makeJito @ Call= EID)*/ To_Call(OBJ(self)).MakeJito( )} 

// tells if the restriction matches the type list lt : we know that the domain is made of classes
/* {0} OPT.The go function for: makeCallMatch(x:restriction,lt:list) [] */
func F_Language_makeCallMatch_restriction (x *ClaireRestriction ,lt *ClaireList ) *ClaireBoolean  { 
  // procedure body with s =  
var Result *ClaireBoolean  
  /* Let:1 */{ 
    var n int  = lt.Length()
    /* noccur = 2 */
    /* Let:2 */{ 
      var ld *ClaireList   = x.Domain
      /* noccur = 2 */
      /* and:3 */{ 
        var v_and3 *ClaireBoolean  
        
        v_and3 = Equal(MakeInteger(ld.Length()).Id(),MakeInteger(n).Id())
        if (v_and3 == CFALSE) {Result = CFALSE
        } else /* arg:4 */{ 
          /* Let:5 */{ 
            var g0700UU *ClaireAny  
            /* noccur = 1 */
            /* Let:6 */{ 
              var i int  = 1
              /* noccur = 5 */
              /* Let:7 */{ 
                var g0699 int  = n
                /* noccur = 1 */
                g0700UU= CFALSE.Id()
                for (i <= g0699) /* while:8 */{ 
                  if (ToType(lt.At(i-1)).Included(ToType(ld.ValuesO()[i-1])) != CTRUE) /* If:9 */{ 
                     /*v = g0700UU, s =any*/
g0700UU = CTRUE.Id()
                    break
                    /* If-9 */} 
                  i = (i+1)
                  /* while-8 */} 
                /* Let-7 */} 
              /* Let-6 */} 
            v_and3 = Core.F_not_any(g0700UU)
            /* Let-5 */} 
          if (v_and3 == CFALSE) {Result = CFALSE
          } else /* arg:5 */{ 
            Result = CTRUE/* arg-5 */} 
          /* arg-4 */} 
        /* and-3 */} 
      /* Let-2 */} 
    /* Let-1 */} 
  return Result} 

// The EID go function for: makeCallMatch @ restriction (throw: false) 
func E_Language_makeCallMatch_restriction (x EID,lt EID) EID { 
  return EID{/*(sm for makeCallMatch @ restriction= boolean)*/ F_Language_makeCallMatch_restriction(ToRestriction(OBJ(x)),ToList(OBJ(lt)) ).Id(),0}} 

// close some classes : final => no subclasses,  default() => ephemeral
// CLAIRE 4 : make sure that open statement for class are all here
// instuctions are ephemeral