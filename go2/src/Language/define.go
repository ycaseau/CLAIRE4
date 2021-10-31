/***** CLAIRE Compilation of file /Users/ycaseau/Dropbox/src/clairev4.0/src/meta/define.cl 
         [version 4.0.01 / safety 5] Saturday 10-30-2021 *****/

package Language
import (_ "fmt"
	. "Kernel"
	"Core"
)

//-------- dumb function to prevent import errors --------
func import_g0440() { 
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
              var g0441 int  = n
              /* noccur = 1 */
              Result= EID{CFALSE.Id(),0}
              for (i <= g0441) /* while:7 */{ 
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
      var g0446UU *ClaireList  
      /* noccur = 1 */
      var g0446UU_try04473 EID 
      g0446UU_try04473 = self.Arg.Args.Cdr()
      /* ERROR PROTECTION INSERTED (g0446UU-Result) */
      if ErrorIn(g0446UU_try04473) {Result = g0446UU_try04473
      } else {
      g0446UU = ToList(OBJ(g0446UU_try04473))
      Result = F_ppvariable_list(g0446UU)
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
        var _Zo_try04494 EID 
        if (_Zc.Open <= 1) /* If:4 */{ 
          _Zo_try04494 = ToException(Core.C_general_error.Make(MakeString("[105] cannot instantiate ~S").Id(),MakeConstantList(_Zc.Id()).Id())).Close()
          } else {
          _Zo_try04494 = EID{CFALSE.Id(),0}
          /* If-4 */} 
        /* ERROR PROTECTION INSERTED (_Zo_try04494-_Zo_try04494) */
        if !ErrorIn(_Zo_try04494) {
        _Zo_try04494 = EID{F_new_object_class(_Zc).Id(),0}
        }
        /* ERROR PROTECTION INSERTED (_Zo-Result) */
        if ErrorIn(_Zo_try04494) {Result = _Zo_try04494
        } else {
        _Zo = ToObject(OBJ(_Zo_try04494))
        /* Let:4 */{ 
          var g0450UU *ClaireList  
          /* noccur = 1 */
          var g0450UU_try04515 EID 
          g0450UU_try04515 = F_Language_new_writes_object(_Zo,self.Args)
          /* ERROR PROTECTION INSERTED (g0450UU-Result) */
          if ErrorIn(g0450UU_try04515) {Result = g0450UU_try04515
          } else {
          g0450UU = ToList(OBJ(g0450UU_try04515))
          Result = Core.F_Core_new_defaults_object(_Zo,g0450UU)
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
            var g0452UU *ClaireAny  
            /* noccur = 1 */
            /* For:6 */{ 
              var s *ClaireAny  
              _ = s
              g0452UU= CFALSE.Id()
              for _,s = range(c.Slots.ValuesO())/* loop:7 */{ 
                if ((ToRestriction(s).Selector.Inverse.Id() != CNULL) || 
                    ((ToRestriction(s).Selector.Store_ask == CTRUE) || 
                      (ToRestriction(s).Selector.IfWrite != CNULL))) /* If:8 */{ 
                   /*v = g0452UU, s =any*/
g0452UU = CTRUE.Id()
                  break
                  /* If-8 */} 
                /* loop-7 */} 
              /* For-6 */} 
            v_and2 = Core.F_not_any(g0452UU)
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
            var p_try04536 EID 
            p_try04536 = F_make_a_property_any(To_Call(x).Args.At(1-1))
            /* ERROR PROTECTION INSERTED (p-void_try5) */
            if ErrorIn(p_try04536) {void_try5 = p_try04536
            } else {
            p = ToProperty(OBJ(p_try04536))
            /* LetEID:6 */{ 
              var g0454UU EID 
              g0454UU = EVAL(To_Call(x).Args.At(2-1))
              /* ERROR PROTECTION INSERTED (g0454UU-void_try5) */
              if ErrorIn(g0454UU) {void_try5 = g0454UU
              } else {
              void_try5 = p.WriteEID(_Zo,g0454UU)}
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
            var p_try04576 EID 
            p_try04576 = F_make_a_property_any(To_Call(x).Args.At(1-1))
            /* ERROR PROTECTION INSERTED (p-void_try5) */
            if ErrorIn(p_try04576) {void_try5 = p_try04576
            } else {
            p = ToProperty(OBJ(p_try04576))
            /* Let:6 */{ 
              var y *ClaireAny  
              /* noccur = 4 */
              var y_try04587 EID 
              y_try04587 = EVAL(To_Call(x).Args.At(2-1))
              /* ERROR PROTECTION INSERTED (y-void_try5) */
              if ErrorIn(y_try04587) {void_try5 = y_try04587
              } else {
              y = ANY(y_try04587)
              /* Let:7 */{ 
                var s *ClaireObject   = Core.F__at_property1(p,self.Isa)
                /* noccur = 2 */
                if (C_slot.Id() == s.Isa.Id()) /* If:8 */{ 
                  /* Let:9 */{ 
                    var g0455 *ClaireSlot   = ToSlot(s.Id())
                    /* noccur = 4 */
                    if (y == CNULL) /* If:10 */{ 
                      lp = lp.AddFast(p.Id())
                      /* If-10 */} 
                    if (g0455.Range.Contains(y) != CTRUE) /* If:10 */{ 
                      void_try5 = Core.F_range_is_wrong_slot(g0455,y)
                      } else {
                      void_try5 = Core.F_update_property(p,
                        self,
                        g0455.Index,
                        g0455.Srange,
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
          var _Zo_try04605 EID 
          _Zo_try04605 = F_new_thing_class(_Zc,self.Ident)
          /* ERROR PROTECTION INSERTED (_Zo-Result) */
          if ErrorIn(_Zo_try04605) {Result = _Zo_try04605
          } else {
          _Zo = ToObject(OBJ(_Zo_try04605))
          Result = EID{_Zo.Id(),0}
          if (_Zo.Isa.IsIn(C_property) == CTRUE) /* If:5 */{ 
            /* Let:6 */{ 
              var g0459 *ClaireProperty   = ToProperty(_Zo.Id())
              /* noccur = 2 */
              if (g0459.Restrictions.Length() > 0) /* If:7 */{ 
                Result = ToException(Core.C_general_error.Make(MakeString("[188] the property ~S is already defined").Id(),MakeConstantList(g0459.Id()).Id())).Close()
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
          var g0461UU *ClaireList  
          /* noccur = 1 */
          var g0461UU_try04625 EID 
          g0461UU_try04625 = F_Language_new_writes_object(_Zo,self.Args)
          /* ERROR PROTECTION INSERTED (g0461UU-Result) */
          if ErrorIn(g0461UU_try04625) {Result = g0461UU_try04625
          } else {
          g0461UU = ToList(OBJ(g0461UU_try04625))
          Result = Core.F_Core_new_defaults_object(_Zo,g0461UU)
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
        var _Zo_try04644 EID 
        _Zo_try04644 = self.Ident.Class_I(self.Arg)
        /* ERROR PROTECTION INSERTED (_Zo-Result) */
        if ErrorIn(_Zo_try04644) {Result = _Zo_try04644
        } else {
        _Zo = ToClass(OBJ(_Zo_try04644))
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
                  var g0463 *Call   = To_Call(x)
                  /* noccur = 4 */
                  var v_try04659 EID 
                  v_try04659 = EVAL(g0463.Args.At(2-1))
                  /* ERROR PROTECTION INSERTED (v-void_try6) */
                  if ErrorIn(v_try04659) {void_try6 = v_try04659
                  Result = v_try04659
                  break
                  } else {
                  v = ANY(v_try04659)
                  void_try6 = v.ToEID()
                  g0463 = To_Call(g0463.Args.At(1-1))
                  void_try6 = EID{g0463.Id(),0}
                  }
                  {
                  x = g0463.Id()
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
                var rt_try04668 EID 
                rt_try04668 = F_extract_type_any(ANY(Core.F_CALL(C_range,ARGS(x.ToEID()))))
                /* ERROR PROTECTION INSERTED (rt-void_try6) */
                if ErrorIn(rt_try04668) {void_try6 = rt_try04668
                } else {
                rt = ToTypeExpression(OBJ(rt_try04668))
                /* Let:8 */{ 
                  var p *ClaireProperty  
                  /* noccur = 4 */
                  var p_try04679 EID 
                  p_try04679 = F_make_a_property_any(ANY(Core.F_CALL(C_mClaire_pname,ARGS(x.ToEID()))))
                  /* ERROR PROTECTION INSERTED (p-void_try6) */
                  if ErrorIn(p_try04679) {void_try6 = p_try04679
                  } else {
                  p = ToProperty(OBJ(p_try04679))
                  var g0468I *ClaireBoolean  
                  var g0468I_try04699 EID 
                  /* and:9 */{ 
                    var v_and9 *ClaireBoolean  
                    
                    v_and9 = Core.F_known_ask_any(v)
                    if (v_and9 == CFALSE) {g0468I_try04699 = EID{CFALSE.Id(),0}
                    } else /* arg:10 */{ 
                      var v_and9_try047011 EID 
                      /* Let:11 */{ 
                        var g0471UU *ClaireBoolean  
                        /* noccur = 1 */
                        var g0471UU_try047212 EID 
                        g0471UU_try047212 = Core.F_BELONG(v,rt.Id())
                        /* ERROR PROTECTION INSERTED (g0471UU-v_and9_try047011) */
                        if ErrorIn(g0471UU_try047212) {v_and9_try047011 = g0471UU_try047212
                        } else {
                        g0471UU = ToBoolean(OBJ(g0471UU_try047212))
                        v_and9_try047011 = EID{g0471UU.Not.Id(),0}
                        }
                        /* Let-11 */} 
                      /* ERROR PROTECTION INSERTED (v_and9-g0468I_try04699) */
                      if ErrorIn(v_and9_try047011) {g0468I_try04699 = v_and9_try047011
                      } else {
                      v_and9 = ToBoolean(OBJ(v_and9_try047011))
                      if (v_and9 == CFALSE) {g0468I_try04699 = EID{CFALSE.Id(),0}
                      } else /* arg:11 */{ 
                        g0468I_try04699 = EID{CTRUE.Id(),0}/* arg-11 */} 
                      /* arg-10 */} 
                    }
                    /* and-9 */} 
                  /* ERROR PROTECTION INSERTED (g0468I-void_try6) */
                  if ErrorIn(g0468I_try04699) {void_try6 = g0468I_try04699
                  } else {
                  g0468I = ToBoolean(OBJ(g0468I_try04699))
                  if (g0468I == CTRUE) /* If:9 */{ 
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
      var p_try04733 EID 
      p_try04733 = F_make_a_property_any(self.Arg.Selector.Id())
      /* ERROR PROTECTION INSERTED (p-Result) */
      if ErrorIn(p_try04733) {Result = p_try04733
      } else {
      p = ToProperty(OBJ(p_try04733))
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
            var lp_try04746 EID 
            lp_try04746 = F_extract_signature_list(lv)
            /* ERROR PROTECTION INSERTED (lp-Result) */
            if ErrorIn(lp_try04746) {Result = lp_try04746
            } else {
            lp = ToList(OBJ(lp_try04746))
            /* Let:6 */{ 
              var lrange *ClaireList  
              /* noccur = 3 */
              var lrange_try04757 EID 
              lrange_try04757 = F_extract_range_any(self.SetArg,lv,ToList(C_LDEF.Value))
              /* ERROR PROTECTION INSERTED (lrange-Result) */
              if ErrorIn(lrange_try04757) {Result = lrange_try04757
              } else {
              lrange = ToList(OBJ(lrange_try04757))
              /* Let:7 */{ 
                var lbody *ClaireList  
                /* noccur = 4 */
                var lbody_try04768 EID 
                lbody_try04768 = F_extract_status_any(self.Body)
                /* ERROR PROTECTION INSERTED (lbody-Result) */
                if ErrorIn(lbody_try04768) {Result = lbody_try04768
                } else {
                lbody = ToList(OBJ(lbody_try04768))
                /* Let:8 */{ 
                  var m *ClaireMethod   = F_add_method_property(p,lp,ToType(lrange.At(1-1)),ToInteger(lbody.At(1-1)).Value,ToFunction(lbody.At(2-1)))
                  /* noccur = 14 */
                  if ((p.Open > 0) && 
                      (p.Open <= 1)) /* If:9 */{ 
                    /* Let:10 */{ 
                      var r *ClaireAny  
                      /* noccur = 2 */
                      var r_try047711 EID 
                      /* Let:11 */{ 
                        var r_some *ClaireAny   = CNULL
                        /* noccur = 2 */
                        /* For:12 */{ 
                          var r *ClaireAny  
                          _ = r
                          r_try047711= EID{CFALSE.Id(),0}
                          for _,r = range(p.Restrictions.ValuesO())/* loop:13 */{ 
                            var void_try14 EID 
                            _ = void_try14
                            if (r != m.Id()) /* If:14 */{ 
                              var g0478I *ClaireBoolean  
                              var g0478I_try047915 EID 
                              /* Let:15 */{ 
                                var g0480UU *ClaireAny  
                                /* noccur = 1 */
                                var g0480UU_try048116 EID 
                                g0480UU_try048116 = Core.F_CALL(ToProperty(C__exp.Id()),ARGS(Core.F_CALL(C_domain,ARGS(r.ToEID())),EID{m.Domain.Id(),0}))
                                /* ERROR PROTECTION INSERTED (g0480UU-g0478I_try047915) */
                                if ErrorIn(g0480UU_try048116) {g0478I_try047915 = g0480UU_try048116
                                } else {
                                g0480UU = ANY(g0480UU_try048116)
                                g0478I_try047915 = EID{F_boolean_I_any(g0480UU).Id(),0}
                                }
                                /* Let-15 */} 
                              /* ERROR PROTECTION INSERTED (g0478I-void_try14) */
                              if ErrorIn(g0478I_try047915) {void_try14 = g0478I_try047915
                              } else {
                              g0478I = ToBoolean(OBJ(g0478I_try047915))
                              if (g0478I == CTRUE) /* If:15 */{ 
                                 /*v = r_try047711, s =EID*/
r_some = r
                                r_try047711 = r_some.ToEID()
                                break
                                } else {
                                void_try14 = EID{CFALSE.Id(),0}
                                /* If-15 */} 
                              }
                              } else {
                              void_try14 = EID{CFALSE.Id(),0}
                              /* If-14 */} 
                            /* ERROR PROTECTION INSERTED (void_try14-r_try047711) */
                            if ErrorIn(void_try14) {r_try047711 = void_try14
                            r_try047711 = void_try14
                            break
                            } else {
                            }
                            /* loop-13 */} 
                          /* For-12 */} 
                        /* ERROR PROTECTION INSERTED (r_try047711-r_try047711) */
                        if !ErrorIn(r_try047711) {
                        r_try047711 = r_some.ToEID()
                        }
                        /* Let-11 */} 
                      /* ERROR PROTECTION INSERTED (r-Result) */
                      if ErrorIn(r_try047711) {Result = r_try047711
                      } else {
                      r = ANY(r_try047711)
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
                      var va_arg2_try048211 EID 
                      /* Let:11 */{ 
                        var g0483UU *ClaireLambda  
                        /* noccur = 1 */
                        var g0483UU_try048412 EID 
                        g0483UU_try048412 = F_lambda_I_list(lv,lbody.At(3-1))
                        /* ERROR PROTECTION INSERTED (g0483UU-va_arg2_try048211) */
                        if ErrorIn(g0483UU_try048412) {va_arg2_try048211 = g0483UU_try048412
                        } else {
                        g0483UU = ToLambda(OBJ(g0483UU_try048412))
                        va_arg2_try048211 = F_Language_jito_any(g0483UU.Id())
                        }
                        /* Let-11 */} 
                      /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                      if ErrorIn(va_arg2_try048211) {Result = va_arg2_try048211
                      } else {
                      va_arg2 = ToLambda(OBJ(va_arg2_try048211))
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
          var v_local3_try04855 EID 
          if (v.Isa.IsIn(C_Variable) != CTRUE) /* If:5 */{ 
            v_local3_try04855 = ToException(Core.C_general_error.Make(MakeString("[111] wrong typed argument ~S").Id(),MakeConstantList(v).Id())).Close()
            } else {
            /* Let:6 */{ 
              var p *ClaireAny  
              /* noccur = 3 */
              var p_try04867 EID 
              p_try04867 = F_extract_pattern_any(To_Variable(v).Range.Id(),MakeConstantList(MakeInteger(n).Id()))
              /* ERROR PROTECTION INSERTED (p-v_local3_try04855) */
              if ErrorIn(p_try04867) {v_local3_try04855 = p_try04867
              } else {
              p = ANY(p_try04867)
              n = (n+1)
              if (p == CNULL) /* If:7 */{ 
                v_local3_try04855 = ToException(Core.C_general_error.Make(MakeString("[111] wrong typed argument ~S (~S)").Id(),MakeConstantList(v,To_Variable(v).Range.Id()).Id())).Close()
                } else {
                v_local3_try04855 = EID{CFALSE.Id(),0}
                /* If-7 */} 
              /* ERROR PROTECTION INSERTED (v_local3_try04855-v_local3_try04855) */
              if ErrorIn(v_local3_try04855) {Result = v_local3_try04855
              break
              } else {
              To_Variable(v).Range = F_type_I_any(p)
              v_local3_try04855 = p.ToEID()
              }
              }
              /* Let-6 */} 
            /* If-5 */} 
          /* ERROR PROTECTION INSERTED (v_local3-Result) */
          if ErrorIn(v_local3_try04855) {Result = v_local3_try04855
          Result = v_local3_try04855
          break
          } else {
          v_local3 = ANY(v_local3_try04855)
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
        var g0487 *ClaireClass   = ToClass(x)
        /* noccur = 1 */
        Result = EID{g0487.Id(),0}
        /* Let-3 */} 
      /* If!2 */}  else if (C_set.Id() == x.Isa.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var g0488 *ClaireSet   = ToSet(x)
        /* noccur = 3 */
        /* Let:4 */{ 
          var z *ClaireAny  
          /* noccur = 2 */
          var z_try04985 EID 
          if (Equal(ANY(Core.F_CALL(C_length,ARGS(EID{g0488.Id(),0}))),MakeInteger(1).Id()) == CTRUE) /* If:5 */{ 
            /* Let:6 */{ 
              var g0499UU *ClaireAny  
              /* noccur = 1 */
              var g0499UU_try05007 EID 
              g0499UU_try05007 = Core.F_CALL(C_nth,ARGS(EID{g0488.Id(),0},EID{C__INT,IVAL(1)}))
              /* ERROR PROTECTION INSERTED (g0499UU-z_try04985) */
              if ErrorIn(g0499UU_try05007) {z_try04985 = g0499UU_try05007
              } else {
              g0499UU = ANY(g0499UU_try05007)
              z_try04985 = F_extract_pattern_any(g0499UU,CNIL)
              }
              /* Let-6 */} 
            } else {
            z_try04985 = EID{CFALSE.Id(),0}
            /* If-5 */} 
          /* ERROR PROTECTION INSERTED (z-Result) */
          if ErrorIn(z_try04985) {Result = z_try04985
          } else {
          z = ANY(z_try04985)
          if (z.Isa.IsIn(C_Reference) == CTRUE) /* If:5 */{ 
            /* Let:6 */{ 
              var g0489 *ClaireReference   = To_Reference(z)
              /* noccur = 1 */
              /* Let:7 */{ 
                var w *ClaireReference   = To_Reference(g0489.Copy().Id())
                /* noccur = 3 */
                w.Arg = CTRUE
                Result = EID{w.Id(),0}
                /* Let-7 */} 
              /* Let-6 */} 
            } else {
            Result = EID{g0488.Id(),0}
            /* If-5 */} 
          }
          /* Let-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (x.Isa.IsIn(C_Tuple) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0491 *Tuple   = To_Tuple(x)
        /* noccur = 1 */
        /* Let:4 */{ 
          var ltp *ClaireList  
          /* noccur = 2 */
          var ltp_try05015 EID 
          /* Iteration:5 */{ 
            var v_list5 *ClaireList  
            var z *ClaireAny  
            var v_local5 *ClaireAny  
            v_list5 = g0491.Args
            ltp_try05015 = EID{CreateList(ToType(CEMPTY.Id()),v_list5.Length()).Id(),0}
            for CLcount := 0; CLcount < v_list5.Length(); CLcount++{ 
              z = v_list5.At(CLcount)
              var v_local5_try05027 EID 
              v_local5_try05027 = F_extract_pattern_any(z,path)
              /* ERROR PROTECTION INSERTED (v_local5-ltp_try05015) */
              if ErrorIn(v_local5_try05027) {ltp_try05015 = v_local5_try05027
              ltp_try05015 = v_local5_try05027
              break
              } else {
              v_local5 = ANY(v_local5_try05027)
              ToList(OBJ(ltp_try05015)).PutAt(CLcount,v_local5)
              } 
            }
            /* Iteration-5 */} 
          /* ERROR PROTECTION INSERTED (ltp-Result) */
          if ErrorIn(ltp_try05015) {Result = ltp_try05015
          } else {
          ltp = ToList(OBJ(ltp_try05015))
          var g0503I *ClaireBoolean  
          /* Let:5 */{ 
            var g0504UU *ClaireAny  
            /* noccur = 1 */
            /* For:6 */{ 
              var y *ClaireAny  
              _ = y
              g0504UU= CFALSE.Id()
              var y_support *ClaireList  
              y_support = ltp
              y_len := y_support.Length()
              for i_it := 0; i_it < y_len; i_it++ { 
                y = y_support.At(i_it)
                if (y == CNULL) /* If:8 */{ 
                   /*v = g0504UU, s =any*/
g0504UU = CTRUE.Id()
                  break
                  /* If-8 */} 
                /* loop-7 */} 
              /* For-6 */} 
            g0503I = F_boolean_I_any(g0504UU)
            /* Let-5 */} 
          if (g0503I == CTRUE) /* If:5 */{ 
            Result = EID{CNULL,0}
            } else {
            Result = EID{ltp.Tuple_I().Id(),0}
            /* If-5 */} 
          }
          /* Let-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (x.Isa.IsIn(Core.C_global_variable) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0492 *Core.GlobalVariable   = Core.ToGlobalVariable(x)
        /* noccur = 1 */
        Result = F_extract_pattern_any(g0492.Value,path)
        /* Let-3 */} 
      /* If!2 */}  else if (x.Isa.IsIn(C_Call) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0493 *Call   = To_Call(x)
        /* noccur = 9 */
        /* Let:4 */{ 
          var p *ClaireProperty   = g0493.Selector
          /* noccur = 5 */
          if (p.Id() == Core.C_U.Id()) /* If:5 */{ 
            /* Let:6 */{ 
              var x1 *ClaireAny  
              /* noccur = 2 */
              var x1_try05057 EID 
              x1_try05057 = F_extract_pattern_any(g0493.Args.At(1-1),CNIL)
              /* ERROR PROTECTION INSERTED (x1-Result) */
              if ErrorIn(x1_try05057) {Result = x1_try05057
              } else {
              x1 = ANY(x1_try05057)
              /* Let:7 */{ 
                var x2 *ClaireAny  
                /* noccur = 2 */
                var x2_try05068 EID 
                x2_try05068 = F_extract_pattern_any(g0493.Args.At(2-1),CNIL)
                /* ERROR PROTECTION INSERTED (x2-Result) */
                if ErrorIn(x2_try05068) {Result = x2_try05068
                } else {
                x2 = ANY(x2_try05068)
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
              var g0507UU *ClaireAny  
              /* noccur = 1 */
              var g0507UU_try05097 EID 
              g0507UU_try05097 = F_extract_pattern_any(g0493.Args.At(1-1),CNIL)
              /* ERROR PROTECTION INSERTED (g0507UU-Result) */
              if ErrorIn(g0507UU_try05097) {Result = g0507UU_try05097
              } else {
              g0507UU = ANY(g0507UU_try05097)
              /* Let:7 */{ 
                var g0508UU *ClaireAny  
                /* noccur = 1 */
                var g0508UU_try05108 EID 
                g0508UU_try05108 = F_extract_pattern_any(g0493.Args.At(2-1),CNIL)
                /* ERROR PROTECTION INSERTED (g0508UU-Result) */
                if ErrorIn(g0508UU_try05108) {Result = g0508UU_try05108
                } else {
                g0508UU = ANY(g0508UU_try05108)
                Result = Core.F_CALL(ToProperty(C__exp.Id()),ARGS(g0507UU.ToEID(),g0508UU.ToEID()))
                }
                /* Let-7 */} 
              }
              /* Let-6 */} 
            /* If!5 */}  else if (p.Id() == C__dot_dot.Id()) /* If:5 */{ 
            /* Let:6 */{ 
              var v1 *ClaireAny   = F_extract_item_any(g0493.Args.At(1-1),CNIL.Id())
              /* noccur = 2 */
              /* Let:7 */{ 
                var v2 *ClaireAny   = F_extract_item_any(g0493.Args.At(2-1),CNIL.Id())
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
            Result = F_extract_pattern_nth_list(g0493.Args,path)
            /* If!5 */}  else if (p.Id() == C__star.Id()) /* If:5 */{ 
            /* Let:6 */{ 
              var z *ClaireAny  
              /* noccur = 2 */
              var z_try05117 EID 
              z_try05117 = F_extract_pattern_any(g0493.Args.At(1-1),path)
              /* ERROR PROTECTION INSERTED (z-Result) */
              if ErrorIn(z_try05117) {Result = z_try05117
              } else {
              z = ANY(z_try05117)
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
        var g0494 *ClaireType   = ToType(x)
        /* noccur = 1 */
        Result = EID{g0494.Id(),0}
        /* Let-3 */} 
      /* If!2 */}  else if (x.Isa.IsIn(C_unbound_symbol) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0495 *ClaireUnboundSymbol   = ToUnboundSymbol(x)
        /* noccur = 1 */
        /* Let:4 */{ 
          var s *ClaireSymbol  
          /* noccur = 3 */
          var s_try05125 EID 
          s_try05125 = F_extract_symbol_any(g0495.Id())
          /* ERROR PROTECTION INSERTED (s-Result) */
          if ErrorIn(s_try05125) {Result = s_try05125
          } else {
          s = ToSymbol(OBJ(s_try05125))
          /* Let:5 */{ 
            var v *ClaireAny  
            /* noccur = 2 */
            var v_try05136 EID 
            /* Let:6 */{ 
              var z_some *ClaireAny   = CNULL
              /* noccur = 2 */
              /* For:7 */{ 
                var z *ClaireAny  
                _ = z
                v_try05136= EID{CFALSE.Id(),0}
                var z_support *ClaireList  
                var z_support_try05148 EID 
                z_support_try05148 = Core.F_enumerate_any(C_LDEF.Value)
                /* ERROR PROTECTION INSERTED (z_support-v_try05136) */
                if ErrorIn(z_support_try05148) {v_try05136 = z_support_try05148
                } else {
                z_support = ToList(OBJ(z_support_try05148))
                z_len := z_support.Length()
                for i_it := 0; i_it < z_len; i_it++ { 
                  z = z_support.At(i_it)
                  if (ANY(Core.F_CALL(C_mClaire_pname,ARGS(z.ToEID()))) == s.Id()) /* If:9 */{ 
                     /*v = v_try05136, s =EID*/
z_some = z
                    v_try05136 = z_some.ToEID()
                    break
                    /* If-9 */} 
                  }
                  /* loop-8 */} 
                /* For-7 */} 
              /* ERROR PROTECTION INSERTED (v_try05136-v_try05136) */
              if !ErrorIn(v_try05136) {
              v_try05136 = z_some.ToEID()
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (v-Result) */
            if ErrorIn(v_try05136) {Result = v_try05136
            } else {
            v = ANY(v_try05136)
            if (v != CNULL) /* If:6 */{ 
              Result = Core.F_CALL(C_range,ARGS(v.ToEID()))
              } else {
              var g0515I *ClaireBoolean  
              if (path.Isa.IsIn(C_list) == CTRUE) /* If:7 */{ 
                g0515I = Core.F__sup_integer(path.Length(),1)
                } else {
                g0515I = CFALSE
                /* If-7 */} 
              if (g0515I == CTRUE) /* If:7 */{ 
                /* Let:8 */{ 
                  var y *ClaireReference  
                  /* noccur = 2 */
                  var y_try05169 EID 
                  /* Let:9 */{ 
                    var _CL_obj *ClaireReference   = To_Reference(new(ClaireReference).Is(C_Reference))
                    /* noccur = 5 */
                    _CL_obj.Index = ToInteger(path.At(1-1)).Value
                    /* update:10 */{ 
                      var va_arg1 *ClaireReference  
                      var va_arg2 *ClaireList  
                      va_arg1 = _CL_obj
                      var va_arg2_try051711 EID 
                      va_arg2_try051711 = path.Cdr()
                      /* ERROR PROTECTION INSERTED (va_arg2-y_try05169) */
                      if ErrorIn(va_arg2_try051711) {y_try05169 = va_arg2_try051711
                      } else {
                      va_arg2 = ToList(OBJ(va_arg2_try051711))
                      /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                      va_arg1.Args = va_arg2
                      y_try05169 = EID{va_arg2.Id(),0}
                      }
                      /* update-10 */} 
                    /* ERROR PROTECTION INSERTED (y_try05169-y_try05169) */
                    if !ErrorIn(y_try05169) {
                    y_try05169 = EID{_CL_obj.Id(),0}
                    }
                    /* Let-9 */} 
                  /* ERROR PROTECTION INSERTED (y-Result) */
                  if ErrorIn(y_try05169) {Result = y_try05169
                  } else {
                  y = To_Reference(OBJ(y_try05169))
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
                    var v_gassign10_try051810 EID 
                    v_gassign10_try051810 = Core.F_CALL(ToProperty(C_add.Id()),ARGS(EID{C_LDEF.Value,0},EID{v.Id(),0}))
                    /* ERROR PROTECTION INSERTED (v_gassign10-Result) */
                    if ErrorIn(v_gassign10_try051810) {Result = v_gassign10_try051810
                    } else {
                    v_gassign10 = ANY(v_gassign10_try051810)
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
    var r_try05192 EID 
    r_try05192 = F_extract_pattern_any(x,CNIL)
    /* ERROR PROTECTION INSERTED (r-Result) */
    if ErrorIn(r_try05192) {Result = r_try05192
    } else {
    r = ANY(r_try05192)
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
          var y_try05265 EID 
          y_try05265 = F_extract_pattern_any(l.At(1-1),CNIL)
          /* ERROR PROTECTION INSERTED (y-Result) */
          if ErrorIn(y_try05265) {Result = y_try05265
          } else {
          y = ANY(y_try05265)
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
            var y_try05276 EID 
            y_try05276 = F_extract_pattern_any(l.At(2-1),CNIL)
            /* ERROR PROTECTION INSERTED (y-Result) */
            if ErrorIn(y_try05276) {Result = y_try05276
            } else {
            y = ANY(y_try05276)
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
                  var g0520 int  = INT(Core.F_CALL(C_length,ARGS(l1.ToEID())))
                  /* noccur = 1 */
                  Result= EID{CFALSE.Id(),0}
                  for (n <= g0520) /* while:9 */{ 
                    var void_try10 EID 
                    _ = void_try10
                    { 
                    /* Let:10 */{ 
                      var y *ClaireAny   = ToList(l2).At(n-1)
                      /* noccur = 3 */
                      var l3_try052811 EID 
                      /* Let:11 */{ 
                        var g0529UU *ClaireAny  
                        /* noccur = 1 */
                        var g0529UU_try053012 EID 
                        if (y.Isa.IsIn(C_Set) == CTRUE) /* If:12 */{ 
                          /* Let:13 */{ 
                            var g0521 *Set   = To_Set(y)
                            /* noccur = 2 */
                            /* Let:14 */{ 
                              var v *ClaireAny  
                              /* noccur = 5 */
                              var v_try053115 EID 
                              /* Let:15 */{ 
                                var g0532UU *ClaireList  
                                /* noccur = 1 */
                                var g0532UU_try053316 EID 
                                /* Let:16 */{ 
                                  var g0534UU *ClaireAny  
                                  /* noccur = 1 */
                                  var g0534UU_try053517 EID 
                                  g0534UU_try053517 = Core.F_CALL(C_nth,ARGS(l1.ToEID(),EID{C__INT,IVAL(n)}))
                                  /* ERROR PROTECTION INSERTED (g0534UU-g0532UU_try053316) */
                                  if ErrorIn(g0534UU_try053517) {g0532UU_try053316 = g0534UU_try053517
                                  } else {
                                  g0534UU = ANY(g0534UU_try053517)
                                  g0532UU_try053316 = EID{path.Copy().AddFast(g0534UU).Id(),0}
                                  }
                                  /* Let-16 */} 
                                /* ERROR PROTECTION INSERTED (g0532UU-v_try053115) */
                                if ErrorIn(g0532UU_try053316) {v_try053115 = g0532UU_try053316
                                } else {
                                g0532UU = ToList(OBJ(g0532UU_try053316))
                                v_try053115 = F_extract_pattern_any(g0521.Args.At(1-1),g0532UU)
                                }
                                /* Let-15 */} 
                              /* ERROR PROTECTION INSERTED (v-g0529UU_try053012) */
                              if ErrorIn(v_try053115) {g0529UU_try053012 = v_try053115
                              } else {
                              v = ANY(v_try053115)
                              if (v == C_void.Id()) /* If:15 */{ 
                                g0529UU_try053012 = EID{C_any.Id(),0}
                                /* If!15 */}  else if (v.Isa.IsIn(C_Reference) == CTRUE) /* If:15 */{ 
                                /* Let:16 */{ 
                                  var g0523 *ClaireReference   = To_Reference(v)
                                  /* noccur = 1 */
                                  /* Let:17 */{ 
                                    var z *ClaireReference   = To_Reference(g0523.Copy().Id())
                                    /* noccur = 2 */
                                    z.Arg = CTRUE
                                    g0529UU_try053012 = EID{z.Id(),0}
                                    /* Let-17 */} 
                                  /* Let-16 */} 
                                } else {
                                /* Construct:16 */{ 
                                  var v_bag_arg *ClaireAny  
                                  g0529UU_try053012= EID{ToType(CEMPTY.Id()).EmptySet().Id(),0}
                                  var v_bag_arg_try053617 EID 
                                  if (v != CNULL) /* If:17 */{ 
                                    v_bag_arg_try053617 = v.ToEID()
                                    } else {
                                    v_bag_arg_try053617 = EVAL(g0521.Args.At(1-1))
                                    /* If-17 */} 
                                  /* ERROR PROTECTION INSERTED (v_bag_arg-g0529UU_try053012) */
                                  if ErrorIn(v_bag_arg_try053617) {g0529UU_try053012 = v_bag_arg_try053617
                                  } else {
                                  v_bag_arg = ANY(v_bag_arg_try053617)
                                  ToSet(OBJ(g0529UU_try053012)).AddFast(v_bag_arg)}
                                  /* Construct-16 */} 
                                /* If-15 */} 
                              }
                              /* Let-14 */} 
                            /* Let-13 */} 
                          } else {
                          /* Let:13 */{ 
                            var g0537UU *ClaireAny  
                            /* noccur = 1 */
                            var g0537UU_try053814 EID 
                            if (path.Length() != 0) /* If:14 */{ 
                              /* Let:15 */{ 
                                var g0539UU *ClaireAny  
                                /* noccur = 1 */
                                var g0539UU_try054016 EID 
                                g0539UU_try054016 = Core.F_CALL(C_nth,ARGS(l1.ToEID(),EID{C__INT,IVAL(n)}))
                                /* ERROR PROTECTION INSERTED (g0539UU-g0537UU_try053814) */
                                if ErrorIn(g0539UU_try054016) {g0537UU_try053814 = g0539UU_try054016
                                } else {
                                g0539UU = ANY(g0539UU_try054016)
                                g0537UU_try053814 = EID{path.AddFast(g0539UU).Id(),0}
                                }
                                /* Let-15 */} 
                              } else {
                              g0537UU_try053814 = EID{CFALSE.Id(),0}
                              /* If-14 */} 
                            /* ERROR PROTECTION INSERTED (g0537UU-g0529UU_try053012) */
                            if ErrorIn(g0537UU_try053814) {g0529UU_try053012 = g0537UU_try053814
                            } else {
                            g0537UU = ANY(g0537UU_try053814)
                            g0529UU_try053012 = F_extract_pattern_any(y,ToList(g0537UU))
                            }
                            /* Let-13 */} 
                          /* If-12 */} 
                        /* ERROR PROTECTION INSERTED (g0529UU-l3_try052811) */
                        if ErrorIn(g0529UU_try053012) {l3_try052811 = g0529UU_try053012
                        } else {
                        g0529UU = ANY(g0529UU_try053012)
                        l3_try052811 = EID{l3.AddFast(g0529UU).Id(),0}
                        }
                        /* Let-11 */} 
                      /* ERROR PROTECTION INSERTED (l3-void_try10) */
                      if ErrorIn(l3_try052811) {void_try10 = l3_try052811
                      } else {
                      l3 = ToList(OBJ(l3_try052811))
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
  var g0548I *ClaireBoolean  
  var g0548I_try05491 EID 
  /* and:1 */{ 
    var v_and1 *ClaireBoolean  
    
    v_and1 = MakeBoolean((self.Id() == C_subtype.Id()) || (self.Id() == C_list.Id()) || (self.Id() == C_set.Id()))
    if (v_and1 == CFALSE) {g0548I_try05491 = EID{CFALSE.Id(),0}
    } else /* arg:2 */{ 
      v_and1 = Equal(MakeInteger(l.Length()).Id(),MakeInteger(1).Id())
      if (v_and1 == CFALSE) {g0548I_try05491 = EID{CFALSE.Id(),0}
      } else /* arg:3 */{ 
        var v_and1_try05504 EID 
        /* Let:4 */{ 
          var y *ClaireAny   = l.At(1-1)
          /* noccur = 7 */
          /* Let:5 */{ 
            var z *ClaireAny  
            /* noccur = 1 */
            var z_try05516 EID 
            z_try05516 = F_extract_pattern_any(y,CNIL)
            /* ERROR PROTECTION INSERTED (z-v_and1_try05504) */
            if ErrorIn(z_try05516) {v_and1_try05504 = z_try05516
            } else {
            z = ANY(z_try05516)
            if (y.Isa.IsIn(Core.C_global_variable) == CTRUE) /* If:6 */{ 
              /* Let:7 */{ 
                var g0541 *Core.GlobalVariable   = Core.ToGlobalVariable(y)
                /* noccur = 2 */
                g0541 = Core.ToGlobalVariable(OBJ(Core.F_CALL(C_value,ARGS(l.At(1-1).ToEID()))))
                y = g0541.Id()
                /* Let-7 */} 
              /* If-6 */} 
            /* or:6 */{ 
              var v_or6 *ClaireBoolean  
              
              v_or6 = z.Isa.IsIn(C_type)
              if (v_or6 == CTRUE) {v_and1_try05504 = EID{CTRUE.Id(),0}
              } else /* or:7 */{ 
                v_or6 = Equal(self.Id(),C_subtype.Id())
                if (v_or6 == CTRUE) {v_and1_try05504 = EID{CTRUE.Id(),0}
                } else /* or:8 */{ 
                  if (y.Isa.IsIn(C_Call) == CTRUE) /* If:9 */{ 
                    /* Let:10 */{ 
                      var g0542 *Call   = To_Call(y)
                      /* noccur = 2 */
                      v_or6 = MakeBoolean((g0542.Selector.Id() != C__equal.Id()) || (g0542.Args.Length() != 2))
                      /* Let-10 */} 
                    /* If!9 */}  else if (y.Isa.IsIn(C_Tuple) == CTRUE) /* If:9 */{ 
                    v_or6 = CTRUE
                    } else {
                    v_or6 = CFALSE
                    /* If-9 */} 
                  if (v_or6 == CTRUE) {v_and1_try05504 = EID{CTRUE.Id(),0}
                  } else /* or:9 */{ 
                    v_and1_try05504 = EID{CFALSE.Id(),0}/* org-9 */} 
                  /* org-8 */} 
                /* org-7 */} 
              /* or-6 */} 
            }
            /* Let-5 */} 
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (v_and1-g0548I_try05491) */
        if ErrorIn(v_and1_try05504) {g0548I_try05491 = v_and1_try05504
        } else {
        v_and1 = ToBoolean(OBJ(v_and1_try05504))
        if (v_and1 == CFALSE) {g0548I_try05491 = EID{CFALSE.Id(),0}
        } else /* arg:4 */{ 
          g0548I_try05491 = EID{CTRUE.Id(),0}/* arg-4 */} 
        /* arg-3 */} 
      /* arg-2 */} 
    }
    /* and-1 */} 
  /* ERROR PROTECTION INSERTED (g0548I-Result) */
  if ErrorIn(g0548I_try05491) {Result = g0548I_try05491
  } else {
  g0548I = ToBoolean(OBJ(g0548I_try05491))
  if (g0548I == CTRUE) /* If:1 */{ 
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
              var g0544 int  = m
              /* noccur = 1 */
              Result= EID{CFALSE.Id(),0}
              for (n <= g0544) /* while:7 */{ 
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
                          var g0545 *Call   = To_Call(y)
                          /* noccur = 5 */
                          if ((g0545.Selector.Id() != C__equal.Id()) || 
                              (g0545.Args.Length() != 2)) /* If:13 */{ 
                            void_try8 = ToException(Core.C_general_error.Make(MakeString("[114] Wrong parametrization ~S").Id(),MakeConstantList(g0545.Id()).Id())).Close()
                            } else {
                            void_try8 = EID{CFALSE.Id(),0}
                            /* If-13 */} 
                          /* ERROR PROTECTION INSERTED (void_try8-void_try8) */
                          if ErrorIn(void_try8) {Result = void_try8
                          break
                          } else {
                          var p_try055213 EID 
                          p_try055213 = F_make_a_property_any(g0545.Args.At(1-1))
                          /* ERROR PROTECTION INSERTED (p-void_try8) */
                          if ErrorIn(p_try055213) {void_try8 = p_try055213
                          Result = p_try055213
                          break
                          } else {
                          p = ANY(p_try055213)
                          void_try8 = p.ToEID()
                          /* Let:13 */{ 
                            var _CL_obj *Set   = To_Set(new(Set).Is(C_Set))
                            /* noccur = 3 */
                            _CL_obj.Args = MakeConstantList(g0545.Args.At(2-1))
                            v = _CL_obj.Id()
                            /* Let-13 */} 
                          void_try8 = v.ToEID()
                          }}
                          /* Let-12 */} 
                        /* If!11 */}  else if (y.Isa.IsIn(C_Vardef) == CTRUE) /* If:11 */{ 
                        /* Let:12 */{ 
                          var g0546 *Vardef   = To_Vardef(y)
                          /* noccur = 2 */
                          var p_try055313 EID 
                          p_try055313 = F_make_a_property_any(g0546.Pname.Id())
                          /* ERROR PROTECTION INSERTED (p-void_try8) */
                          if ErrorIn(p_try055313) {void_try8 = p_try055313
                          Result = p_try055313
                          break
                          } else {
                          p = ANY(p_try055313)
                          void_try8 = p.ToEID()
                          v = g0546.Range.Id()
                          void_try8 = v.ToEID()
                          }
                          /* Let-12 */} 
                        } else {
                        var p_try055412 EID 
                        p_try055412 = F_make_a_property_any(y)
                        /* ERROR PROTECTION INSERTED (p-void_try8) */
                        if ErrorIn(p_try055412) {void_try8 = p_try055412
                        Result = p_try055412
                        break
                        } else {
                        p = ANY(p_try055412)
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
                var g0555UU *ClaireList  
                /* noccur = 1 */
                /* Construct:8 */{ 
                  var v_bag_arg *ClaireAny  
                  g0555UU= ToType(CEMPTY.Id()).EmptyList()
                  g0555UU.AddFast(l1.Id())
                  /* Let:9 */{ 
                    var _CL_obj *List   = To_List(new(List).Is(C_List))
                    /* noccur = 3 */
                    _CL_obj.Args = l2
                    v_bag_arg = _CL_obj.Id()
                    /* Let-9 */} 
                  g0555UU.AddFast(v_bag_arg)/* Construct-8 */} 
                va_arg2 = F_cons_any(self.Id(),g0555UU)
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
  var g0558I *ClaireBoolean  
  /* Let:1 */{ 
    var g0559UU *ClaireBoolean  
    /* noccur = 1 */
    if (x.Isa.IsIn(C_Call) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0556 *Call   = To_Call(x)
        /* noccur = 2 */
        g0559UU = MakeBoolean((g0556.Selector.Id() == C_nth.Id()) && (g0556.Args.At(1-1) == C_type.Id()))
        /* Let-3 */} 
      } else {
      g0559UU = CFALSE
      /* If-2 */} 
    g0558I = g0559UU.Not
    /* Let-1 */} 
  if (g0558I == CTRUE) /* If:1 */{ 
    /* Construct:2 */{ 
      var v_bag_arg *ClaireAny  
      Result= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
      var v_bag_arg_try05603 EID 
      v_bag_arg_try05603 = F_extract_type_any(x)
      /* ERROR PROTECTION INSERTED (v_bag_arg-Result) */
      if ErrorIn(v_bag_arg_try05603) {Result = v_bag_arg_try05603
      } else {
      v_bag_arg = ANY(v_bag_arg_try05603)
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
                    var g0557 int  = n
                    /* noccur = 1 */
                    for (i <= g0557) /* while:10 */{ 
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
                  var g0561UU *Call  
                  /* noccur = 1 */
                  /* Let:9 */{ 
                    var _CL_obj *Call   = To_Call(new(Call).Is(C_Call))
                    /* noccur = 5 */
                    _CL_obj.Selector = Core.C_member
                    _CL_obj.Args = MakeConstantList(y)
                    g0561UU = _CL_obj
                    /* Let-9 */} 
                  x = F_substitution_any(x,To_Variable(v),g0561UU.Id())
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
        var lb_try05624 EID 
        lb_try05624 = F_lambda_I_list(lv2,ToList(OBJ(Core.F_CALL(C_args,ARGS(x.ToEID())))).At(2-1))
        /* ERROR PROTECTION INSERTED (lb-Result) */
        if ErrorIn(lb_try05624) {Result = lb_try05624
        } else {
        lb = ToLambda(OBJ(lb_try05624))
        /* Let:4 */{ 
          var ur *ClaireAny   = CNULL
          /* noccur = 4 */
          /* Let:5 */{ 
            var g0563UU *ClaireList  
            /* noccur = 1 */
            /* Construct:6 */{ 
              var v_bag_arg *ClaireAny  
              g0563UU= ToType(CEMPTY.Id()).EmptyList()
              g0563UU.AddFast(lb.Id())
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
              g0563UU.AddFast(v_bag_arg)/* Construct-6 */} 
            Core.F_tformat_string(MakeString("extract range applies type lambda ~S to arg list ~S \n"),0,g0563UU)
            /* Let-5 */} 
          h_index := ClEnv.Index /* Handle */
          h_base := ClEnv.Base
          var ur_try05645 EID 
          /* Let:5 */{ 
            var g0565UU *ClaireList  
            /* noccur = 1 */
            /* Iteration:6 */{ 
              var v_list6 *ClaireList  
              var v *ClaireAny  
              var v_local6 *ClaireAny  
              v_list6 = lvar
              g0565UU = CreateList(ToType(CEMPTY.Id()),v_list6.Length())
              for CLcount := 0; CLcount < v_list6.Length(); CLcount++{ 
                v = v_list6.At(CLcount)
                v_local6 = ANY(Core.F_CALL(C_range,ARGS(v.ToEID())))
                g0565UU.PutAt(CLcount,v_local6)
                } 
              /* Iteration-6 */} 
            ur_try05645 = F_apply_lambda(lb,g0565UU)
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (ur-Result) */
          if ErrorIn(ur_try05645) {Result = ur_try05645
          } else {
          ur = ANY(ur_try05645)
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
              var g0566UU *ClaireList  
              /* noccur = 1 */
              /* Iteration:7 */{ 
                var v_list7 *ClaireList  
                var v *ClaireAny  
                var v_local7 *ClaireAny  
                v_list7 = lvar
                g0566UU = CreateList(ToType(CEMPTY.Id()),v_list7.Length())
                for CLcount := 0; CLcount < v_list7.Length(); CLcount++{ 
                  v = v_list7.At(CLcount)
                  v_local7 = ANY(Core.F_CALL(C_range,ARGS(v.ToEID())))
                  g0566UU.PutAt(CLcount,v_local7)
                  } 
                /* Iteration-7 */} 
              Result = Core.F_print_any(g0566UU.Id())
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
        var d_try05674 EID 
        /* Let:4 */{ 
          var g0568UU int 
          /* noccur = 1 */
          var g0568UU_try05695 EID 
          g0568UU_try05695 = F__exp2_integer(ToInteger(x).Value)
          /* ERROR PROTECTION INSERTED (g0568UU-d_try05674) */
          if ErrorIn(g0568UU_try05695) {d_try05674 = g0568UU_try05695
          } else {
          g0568UU = INT(g0568UU_try05695)
          d_try05674 = EID{C__INT,IVAL((d+g0568UU))}
          }
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (d-Result) */
        if ErrorIn(d_try05674) {Result = d_try05674
        Result = d_try05674
        break
        } else {
        d = INT(d_try05674)
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
      var g0575I *ClaireBoolean  
      if (x.Isa.IsIn(C_Call) == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var g0570 *Call   = To_Call(x)
          /* noccur = 1 */
          g0575I = Equal(g0570.Selector.Id(),C_function_I.Id())
          /* Let-4 */} 
        } else {
        g0575I = CFALSE
        /* If-3 */} 
      if (g0575I == CTRUE) /* If:3 */{ 
        f = x
        } else {
        f = CNULL
        /* If-3 */} 
      if (x.Isa.IsIn(C_And) == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var g0571 *And   = To_And(x)
          /* noccur = 4 */
          /* Let:5 */{ 
            var y *ClaireAny   = g0571.Args.At(1-1)
            /* noccur = 3 */
            var g0576I *ClaireBoolean  
            if (y.Isa.IsIn(C_Call) == CTRUE) /* If:6 */{ 
              /* Let:7 */{ 
                var g0572 *Call   = To_Call(y)
                /* noccur = 1 */
                g0576I = Equal(g0572.Selector.Id(),C_function_I.Id())
                /* Let-7 */} 
              } else {
              g0576I = CFALSE
              /* If-6 */} 
            if (g0576I == CTRUE) /* If:6 */{ 
              f = y
              g0571 = To_And(g0571.Args.At(2-1))
              /* If-6 */} 
            /* Let-5 */} 
          x = g0571.Id()
          /* Let-4 */} 
        /* If!3 */}  else if (x.Isa.IsIn(C_Call) == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var g0573 *Call   = To_Call(x)
          /* noccur = 3 */
          if (g0573.Selector.Id() == C_function_I.Id()) /* If:5 */{ 
            g0573 = To_Call(C_body.Id())
            /* If-5 */} 
          x = g0573.Id()
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
        var f_try05774 EID 
        /* Let:4 */{ 
          var g0578UU *ClaireString  
          /* noccur = 1 */
          var g0578UU_try05795 EID 
          /* Let:5 */{ 
            var g0580UU *ClaireSymbol  
            /* noccur = 1 */
            var g0580UU_try05816 EID 
            g0580UU_try05816 = F_extract_symbol_any(ToList(OBJ(Core.F_CALL(C_args,ARGS(f.ToEID())))).At(1-1))
            /* ERROR PROTECTION INSERTED (g0580UU-g0578UU_try05795) */
            if ErrorIn(g0580UU_try05816) {g0578UU_try05795 = g0580UU_try05816
            } else {
            g0580UU = ToSymbol(OBJ(g0580UU_try05816))
            g0578UU_try05795 = EID{g0580UU.String_I().Id(),0}
            }
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (g0578UU-f_try05774) */
          if ErrorIn(g0578UU_try05795) {f_try05774 = g0578UU_try05795
          } else {
          g0578UU = ToString(OBJ(g0578UU_try05795))
          f_try05774 = F_imported_function_string(g0578UU).ToEID()
          }
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (f-Result) */
        if ErrorIn(f_try05774) {Result = f_try05774
        } else {
        f = ANY(f_try05774)
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
      var g0582 *ClaireList   = ToList(x)
      /* noccur = 1 */
      /* Iteration:3 */{ 
        var v_list3 *ClaireList  
        var y *ClaireAny  
        var v_local3 *ClaireAny  
        v_list3 = g0582
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
      var g0583 *ClaireParam   = To_Param(x)
      /* noccur = 3 */
      /* Let:3 */{ 
        var _CL_obj *ClaireParam   = To_Param(new(ClaireParam).Is(C_Param))
        /* noccur = 7 */
        _CL_obj.Arg = g0583.Arg
        _CL_obj.Params = g0583.Params
        /* update:4 */{ 
          var va_arg1 *ClaireParam  
          var va_arg2 *ClaireList  
          va_arg1 = _CL_obj
          /* Iteration:5 */{ 
            var v_list5 *ClaireList  
            var y *ClaireAny  
            var v_local5 *ClaireAny  
            v_list5 = g0583.Args
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
      var g0585 *ClaireType   = ToType(x)
      /* noccur = 1 */
      Result = g0585
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
      var ar_try05953 EID 
      /* Let:3 */{ 
        var g0596UU *ClaireSymbol  
        /* noccur = 1 */
        var g0596UU_try05974 EID 
        g0596UU_try05974 = F_extract_symbol_any(a.At(1-1))
        /* ERROR PROTECTION INSERTED (g0596UU-ar_try05953) */
        if ErrorIn(g0596UU_try05974) {ar_try05953 = g0596UU_try05974
        } else {
        g0596UU = ToSymbol(OBJ(g0596UU_try05974))
        ar_try05953 = new(ClaireTable).IsNamed(C_table,g0596UU).ToEID()
        }
        /* Let-3 */} 
      /* ERROR PROTECTION INSERTED (ar-Result) */
      if ErrorIn(ar_try05953) {Result = ar_try05953
      } else {
      ar = ToTable(OBJ(ar_try05953))
      /* Let:3 */{ 
        var v *ClaireVariable   = To_Variable(a.At(2-1))
        /* noccur = 2 */
        /* Let:4 */{ 
          var s *ClaireTypeExpression  
          /* noccur = 11 */
          var s_try05985 EID 
          s_try05985 = F_extract_type_any(v.Range.Id())
          /* ERROR PROTECTION INSERTED (s-Result) */
          if ErrorIn(s_try05985) {Result = s_try05985
          } else {
          s = ToTypeExpression(OBJ(s_try05985))
          /* Let:5 */{ 
            var e *ClaireAny  
            /* noccur = 5 */
            var e_try05996 EID 
            /* Let:6 */{ 
              var l *ClaireList  
              /* noccur = 3 */
              var l_try06007 EID 
              l_try06007 = a.Cdr()
              /* ERROR PROTECTION INSERTED (l-e_try05996) */
              if ErrorIn(l_try06007) {e_try05996 = l_try06007
              } else {
              l = ToList(OBJ(l_try06007))
              /* Let:7 */{ 
                var b *ClaireAny  
                /* noccur = 2 */
                var b_try06018 EID 
                b_try06018 = F_lexical_build_any(self.Body,l,0)
                /* ERROR PROTECTION INSERTED (b-e_try05996) */
                if ErrorIn(b_try06018) {e_try05996 = b_try06018
                } else {
                b = ANY(b_try06018)
                var g0602I *ClaireBoolean  
                /* Let:8 */{ 
                  var g0603UU *ClaireAny  
                  /* noccur = 1 */
                  /* For:9 */{ 
                    var va *ClaireAny  
                    _ = va
                    g0603UU= CFALSE.Id()
                    var va_support *ClaireList  
                    va_support = l
                    va_len := va_support.Length()
                    for i_it := 0; i_it < va_len; i_it++ { 
                      va = va_support.At(i_it)
                      if (F_occurrence_any(b,To_Variable(va)) > 0) /* If:11 */{ 
                         /*v = g0603UU, s =any*/
g0603UU = CTRUE.Id()
                        break
                        /* If-11 */} 
                      /* loop-10 */} 
                    /* For-9 */} 
                  g0602I = F_boolean_I_any(g0603UU)
                  /* Let-8 */} 
                if (g0602I == CTRUE) /* If:8 */{ 
                  e_try05996 = F_lambda_I_list(l,b)
                  } else {
                  e_try05996 = self.Body.ToEID()
                  /* If-8 */} 
                }
                /* Let-7 */} 
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (e-Result) */
            if ErrorIn(e_try05996) {Result = e_try05996
            } else {
            e = ANY(e_try05996)
            /* Let:6 */{ 
              var d *ClaireAny  
              /* noccur = 9 */
              var d_try06047 EID 
              if (e.Isa.IsIn(C_lambda) == CTRUE) /* If:7 */{ 
                d_try06047 = EID{CNULL,0}
                } else {
                d_try06047 = EVAL(self.Body)
                /* If-7 */} 
              /* ERROR PROTECTION INSERTED (d-Result) */
              if ErrorIn(d_try06047) {Result = d_try06047
              } else {
              d = ANY(d_try06047)
              /* update:7 */{ 
                var va_arg1 *ClaireRelation  
                var va_arg2 *ClaireType  
                va_arg1 = ToRelation(ar.Id())
                var va_arg2_try06058 EID 
                va_arg2_try06058 = F_extract_pattern_any(self.SetArg,CNIL)
                /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                if ErrorIn(va_arg2_try06058) {Result = va_arg2_try06058
                } else {
                va_arg2 = ToType(OBJ(va_arg2_try06058))
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
                    var g0589 *ClaireInterval   = To_Interval(s.Id())
                    /* noccur = 2 */
                    ar.Params = MakeInteger((g0589.Arg1-1)).Id()
                    ar.Graph = Core.F_make_copy_list_integer(Core.F_size_Interval(g0589),d).Id()
                    /* Let-9 */} 
                  } else {
                  ar.Params = C_any.Id()
                  ar.GraphInit()
                  /* If-8 */} 
                if (e.Isa.IsIn(C_lambda) == CTRUE) /* If:8 */{ 
                  /* Let:9 */{ 
                    var g0591 *ClaireLambda   = ToLambda(e)
                    /* noccur = 1 */
                    /* For:10 */{ 
                      var y *ClaireAny  
                      _ = y
                      Result= EID{CFALSE.Id(),0}
                      var y_support *ClaireList  
                      var y_support_try060611 EID 
                      y_support_try060611 = Core.F_enumerate_any(ar.Domain.Id())
                      /* ERROR PROTECTION INSERTED (y_support-Result) */
                      if ErrorIn(y_support_try060611) {Result = y_support_try060611
                      } else {
                      y_support = ToList(OBJ(y_support_try060611))
                      y_len := y_support.Length()
                      for i_it := 0; i_it < y_len; i_it++ { 
                        y = y_support.At(i_it)
                        var void_try12 EID 
                        _ = void_try12
                        /* Let:12 */{ 
                          var g0607UU *ClaireAny  
                          /* noccur = 1 */
                          var g0607UU_try060813 EID 
                          g0607UU_try060813 = Core.F_funcall_lambda1(g0591,y)
                          /* ERROR PROTECTION INSERTED (g0607UU-void_try12) */
                          if ErrorIn(g0607UU_try060813) {void_try12 = g0607UU_try060813
                          } else {
                          g0607UU = ANY(g0607UU_try060813)
                          void_try12 = Core.F_nth_equal_table1(ar,y,g0607UU)
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
                  var s2_try06099 EID 
                  s2_try06099 = F_extract_type_any(To_Variable(a.At(3-1)).Range.Id())
                  /* ERROR PROTECTION INSERTED (s2-Result) */
                  if ErrorIn(s2_try06099) {Result = s2_try06099
                  } else {
                  s2 = ToTypeExpression(OBJ(s2_try06099))
                  ar.Domain = ToType(MakeConstantList(s.Id(),s2.Id()).Tuple_I().Id())
                  To_Variable(a.At(3-1)).Range = ToType(s2.Id())
                  if ((s.Isa.IsIn(C_Interval) == CTRUE) && 
                      (s2.Isa.IsIn(C_Interval) == CTRUE)) /* If:9 */{ 
                    
                    /* update:10 */{ 
                      var va_arg1 *ClaireTable  
                      var va_arg2 *ClaireAny  
                      va_arg1 = ar
                      var va_arg2_try061011 EID 
                      /* Construct:11 */{ 
                        var v_bag_arg *ClaireAny  
                        va_arg2_try061011= EID{ToType(C_integer.Id()).EmptyList().Id(),0}
                        var v_bag_arg_try061112 EID 
                        v_bag_arg_try061112 = Core.F_CALL(C_size,ARGS(EID{s2.Id(),0}))
                        /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try061011) */
                        if ErrorIn(v_bag_arg_try061112) {va_arg2_try061011 = v_bag_arg_try061112
                        } else {
                        v_bag_arg = ANY(v_bag_arg_try061112)
                        ToList(OBJ(va_arg2_try061011)).AddFast(v_bag_arg)
                        var v_bag_arg_try061212 EID 
                        /* Let:12 */{ 
                          var g0613UU int 
                          /* noccur = 1 */
                          var g0613UU_try061413 EID 
                          /* Let:13 */{ 
                            var g0615UU int 
                            /* noccur = 1 */
                            var g0615UU_try061614 EID 
                            /* Let:14 */{ 
                              var g0617UU *ClaireAny  
                              /* noccur = 1 */
                              var g0617UU_try061815 EID 
                              g0617UU_try061815 = Core.F_CALL(C_size,ARGS(EID{s2.Id(),0}))
                              /* ERROR PROTECTION INSERTED (g0617UU-g0615UU_try061614) */
                              if ErrorIn(g0617UU_try061815) {g0615UU_try061614 = g0617UU_try061815
                              } else {
                              g0617UU = ANY(g0617UU_try061815)
                              g0615UU_try061614 = EID{C__INT,IVAL((To_Interval(s.Id()).Arg1*ToInteger(g0617UU).Value))}
                              }
                              /* Let-14 */} 
                            /* ERROR PROTECTION INSERTED (g0615UU-g0613UU_try061413) */
                            if ErrorIn(g0615UU_try061614) {g0613UU_try061413 = g0615UU_try061614
                            } else {
                            g0615UU = INT(g0615UU_try061614)
                            g0613UU_try061413 = EID{C__INT,IVAL((g0615UU+To_Interval(s2.Id()).Arg1))}
                            }
                            /* Let-13 */} 
                          /* ERROR PROTECTION INSERTED (g0613UU-v_bag_arg_try061212) */
                          if ErrorIn(g0613UU_try061413) {v_bag_arg_try061212 = g0613UU_try061413
                          } else {
                          g0613UU = INT(g0613UU_try061413)
                          v_bag_arg_try061212 = EID{C__INT,IVAL((g0613UU-1))}
                          }
                          /* Let-12 */} 
                        /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try061011) */
                        if ErrorIn(v_bag_arg_try061212) {va_arg2_try061011 = v_bag_arg_try061212
                        } else {
                        v_bag_arg = ANY(v_bag_arg_try061212)
                        ToList(OBJ(va_arg2_try061011)).AddFast(v_bag_arg)}}
                        /* Construct-11 */} 
                      /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                      if ErrorIn(va_arg2_try061011) {Result = va_arg2_try061011
                      } else {
                      va_arg2 = ANY(va_arg2_try061011)
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
                      var va_arg2_try061911 EID 
                      /* Let:11 */{ 
                        var g0620UU int 
                        /* noccur = 1 */
                        var g0620UU_try062112 EID 
                        /* Let:12 */{ 
                          var g0622UU *ClaireAny  
                          /* noccur = 1 */
                          var g0622UU_try062413 EID 
                          g0622UU_try062413 = Core.F_CALL(C_size,ARGS(EID{s.Id(),0}))
                          /* ERROR PROTECTION INSERTED (g0622UU-g0620UU_try062112) */
                          if ErrorIn(g0622UU_try062413) {g0620UU_try062112 = g0622UU_try062413
                          } else {
                          g0622UU = ANY(g0622UU_try062413)
                          /* Let:13 */{ 
                            var g0623UU *ClaireAny  
                            /* noccur = 1 */
                            var g0623UU_try062514 EID 
                            g0623UU_try062514 = Core.F_CALL(C_size,ARGS(EID{s2.Id(),0}))
                            /* ERROR PROTECTION INSERTED (g0623UU-g0620UU_try062112) */
                            if ErrorIn(g0623UU_try062514) {g0620UU_try062112 = g0623UU_try062514
                            } else {
                            g0623UU = ANY(g0623UU_try062514)
                            g0620UU_try062112 = F_times_integer(ToInteger(g0622UU).Value,ToInteger(g0623UU).Value)
                            }
                            /* Let-13 */} 
                          }
                          /* Let-12 */} 
                        /* ERROR PROTECTION INSERTED (g0620UU-va_arg2_try061911) */
                        if ErrorIn(g0620UU_try062112) {va_arg2_try061911 = g0620UU_try062112
                        } else {
                        g0620UU = INT(g0620UU_try062112)
                        va_arg2_try061911 = EID{Core.F_make_copy_list_integer(g0620UU,d).Id(),0}
                        }
                        /* Let-11 */} 
                      /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                      if ErrorIn(va_arg2_try061911) {Result = va_arg2_try061911
                      } else {
                      va_arg2 = ANY(va_arg2_try061911)
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
                      var g0593 *ClaireLambda   = ToLambda(e)
                      /* noccur = 1 */
                      /* For:11 */{ 
                        var y1 *ClaireAny  
                        _ = y1
                        Result= EID{CFALSE.Id(),0}
                        var y1_support *ClaireList  
                        var y1_support_try062612 EID 
                        y1_support_try062612 = Core.F_enumerate_any(s.Id())
                        /* ERROR PROTECTION INSERTED (y1_support-Result) */
                        if ErrorIn(y1_support_try062612) {Result = y1_support_try062612
                        } else {
                        y1_support = ToList(OBJ(y1_support_try062612))
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
                            var y2_support_try062714 EID 
                            y2_support_try062714 = Core.F_enumerate_any(s2.Id())
                            /* ERROR PROTECTION INSERTED (y2_support-void_try13) */
                            if ErrorIn(y2_support_try062714) {void_try13 = y2_support_try062714
                            } else {
                            y2_support = ToList(OBJ(y2_support_try062714))
                            y2_len := y2_support.Length()
                            for i_it := 0; i_it < y2_len; i_it++ { 
                              y2 = y2_support.At(i_it)
                              var void_try15 EID 
                              _ = void_try15
                              /* Let:15 */{ 
                                var g0628UU *ClaireAny  
                                /* noccur = 1 */
                                var g0628UU_try062916 EID 
                                g0628UU_try062916 = Core.F_CALL(C_funcall,ARGS(EID{g0593.Id(),0},y1.ToEID(),y2.ToEID()))
                                /* ERROR PROTECTION INSERTED (g0628UU-void_try15) */
                                if ErrorIn(g0628UU_try062916) {void_try15 = g0628UU_try062916
                                } else {
                                g0628UU = ANY(g0628UU_try062916)
                                void_try15 = Core.F_nth_equal_table2(ar,y1,y2,g0628UU)
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
          var g0630 *ClaireTuple  
          /* noccur = 2 */
          var g0630_try06335 EID 
          g0630_try06335 = F_make_filter_any(_Zcondition)
          /* ERROR PROTECTION INSERTED (g0630-Result) */
          if ErrorIn(g0630_try06335) {Result = g0630_try06335
          } else {
          g0630 = ToTuple(OBJ(g0630_try06335))
          /* Let:5 */{ 
            var R *ClaireAny   = ToList(g0630.Id()).At(1-1)
            /* noccur = 12 */
            /* Let:6 */{ 
              var lvar *ClaireAny   = ToList(g0630.Id()).At(2-1)
              /* noccur = 2 */
              /* Let:7 */{ 
                var d *LanguageDemon  
                /* noccur = 2 */
                var d_try06348 EID 
                /* Let:8 */{ 
                  var g0635UU *ClaireAny  
                  /* noccur = 1 */
                  var g0635UU_try06369 EID 
                  g0635UU_try06369 = F_lexical_build_any(self.Body,ToList(lvar),0)
                  /* ERROR PROTECTION INSERTED (g0635UU-d_try06348) */
                  if ErrorIn(g0635UU_try06369) {d_try06348 = g0635UU_try06369
                  } else {
                  g0635UU = ANY(g0635UU_try06369)
                  d_try06348 = F_make_demon_relation(ToRelation(R),
                    ToSymbol(OBJ(Core.F_CALL(C_name,ARGS(ru.ToEID())))),
                    ToList(lvar),
                    _Zcondition,
                    g0635UU)
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (d-Result) */
                if ErrorIn(d_try06348) {Result = d_try06348
                } else {
                d = ToLanguageDemon(OBJ(d_try06348))
                if (C_function.Id() == Core.F_owner_any(ANY(Core.F_CALL(C_if_write,ARGS(R.ToEID())))).Id()) /* If:8 */{ 
                  Result = ToException(Core.C_general_error.Make(MakeString("cannot define a new rule on ~S which is closed").Id(),MakeConstantList(R).Id())).Close()
                  } else {
                  Result = EID{CFALSE.Id(),0}
                  /* If-8 */} 
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                Core.F_tformat_string(MakeString("we have defined a demon ~S for ~S \n"),0,MakeConstantList(d.Id(),R))
                /* Let:8 */{ 
                  var g0637UU *ClaireList  
                  /* noccur = 1 */
                  var g0637UU_try06389 EID 
                  /* Let:9 */{ 
                    var g0639UU *ClaireAny  
                    /* noccur = 1 */
                    var g0639UU_try064010 EID 
                    g0639UU_try064010 = Core.F_nth_table1(C_demons,R)
                    /* ERROR PROTECTION INSERTED (g0639UU-g0637UU_try06389) */
                    if ErrorIn(g0639UU_try064010) {g0637UU_try06389 = g0639UU_try064010
                    } else {
                    g0639UU = ANY(g0639UU_try064010)
                    g0637UU_try06389 = ToList(g0639UU).Add(d.Id())
                    }
                    /* Let-9 */} 
                  /* ERROR PROTECTION INSERTED (g0637UU-Result) */
                  if ErrorIn(g0637UU_try06389) {Result = g0637UU_try06389
                  } else {
                  g0637UU = ToList(OBJ(g0637UU_try06389))
                  Core.F_put_table(C_demons,R,g0637UU.Id())
                  Result = EVOID
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                Result = Core.F_nth_put_table(C_Language_last_rule,R,ru)
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                var g0641I *ClaireBoolean  
                var g0641I_try06428 EID 
                /* Let:8 */{ 
                  var g0643UU int 
                  /* noccur = 1 */
                  var g0643UU_try06449 EID 
                  /* Let:9 */{ 
                    var g0645UU *ClaireAny  
                    /* noccur = 1 */
                    var g0645UU_try064610 EID 
                    g0645UU_try064610 = Core.F_nth_table1(C_demons,R)
                    /* ERROR PROTECTION INSERTED (g0645UU-g0643UU_try06449) */
                    if ErrorIn(g0645UU_try064610) {g0643UU_try06449 = g0645UU_try064610
                    } else {
                    g0645UU = ANY(g0645UU_try064610)
                    g0643UU_try06449 = EID{C__INT,IVAL(ToList(g0645UU).Length())}
                    }
                    /* Let-9 */} 
                  /* ERROR PROTECTION INSERTED (g0643UU-g0641I_try06428) */
                  if ErrorIn(g0643UU_try06449) {g0641I_try06428 = g0643UU_try06449
                  } else {
                  g0643UU = INT(g0643UU_try06449)
                  g0641I_try06428 = EID{Equal(MakeInteger(g0643UU).Id(),MakeInteger(1).Id()).Id(),0}
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (g0641I-Result) */
                if ErrorIn(g0641I_try06428) {Result = g0641I_try06428
                } else {
                g0641I = ToBoolean(OBJ(g0641I_try06428))
                if (g0641I == CTRUE) /* If:8 */{ 
                  Result = F_eval_if_write_relation(ToRelation(R))
                  } else {
                  Result = EID{CFALSE.Id(),0}
                  /* If-8 */} 
                }
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                var g0647I *ClaireBoolean  
                if (R.Isa.IsIn(C_property) == CTRUE) /* If:8 */{ 
                  /* Let:9 */{ 
                    var g0632 *ClaireProperty   = ToProperty(R)
                    /* noccur = 1 */
                    g0647I = Equal(MakeInteger(g0632.Restrictions.Length()).Id(),MakeInteger(0).Id())
                    /* Let-9 */} 
                  } else {
                  g0647I = CFALSE
                  /* If-8 */} 
                if (g0647I == CTRUE) /* If:8 */{ 
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
      var g0648 *ClaireProperty   = ToProperty(r.Id())
      /* noccur = 1 */
      /* Let:3 */{ 
        var g0649UU *ClaireAny  
        /* noccur = 1 */
        /* For:4 */{ 
          var x *ClaireAny  
          _ = x
          g0649UU= CFALSE.Id()
          for _,x = range(g0648.Restrictions.ValuesO())/* loop:5 */{ 
            if (C_slot.Id() == x.Isa.Id()) /* If:6 */{ 
               /*v = g0649UU, s =any*/
g0649UU = CTRUE.Id()
              break
              /* If-6 */} 
            /* loop-5 */} 
          /* For-4 */} 
        Result = Core.F_not_any(g0649UU)
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
        var g0650 *And   = To_And(cond)
        /* noccur = 1 */
        c = g0650.Args.At(1-1)
        /* Let-3 */} 
      } else {
      c = cond
      /* If-2 */} 
    
    var g0656I *ClaireBoolean  
    if (c.Isa.IsIn(C_Call) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0652 *Call   = To_Call(c)
        /* noccur = 3 */
        g0656I = MakeBoolean(((g0652.Selector.Id() == Core.C_write.Id()) || 
            (g0652.Selector.Id() == C_nth_equal.Id())) && (g0652.Args.At(1-1).Isa.IsIn(C_relation) == CTRUE))
        /* Let-3 */} 
      } else {
      g0656I = CFALSE
      /* If-2 */} 
    if (g0656I == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var R *ClaireRelation   = ToRelation(ToList(OBJ(Core.F_CALL(C_args,ARGS(c.ToEID())))).At(1-1))
        /* noccur = 9 */
        /* Let:4 */{ 
          var x *ClaireVariable  
          /* noccur = 2 */
          var x_try06575 EID 
          /* Let:5 */{ 
            var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
            /* noccur = 5 */
            /* update:6 */{ 
              var va_arg1 *ClaireVariable  
              var va_arg2 *ClaireSymbol  
              va_arg1 = _CL_obj
              var va_arg2_try06587 EID 
              va_arg2_try06587 = F_extract_symbol_any(ToList(OBJ(Core.F_CALL(C_args,ARGS(c.ToEID())))).At(2-1))
              /* ERROR PROTECTION INSERTED (va_arg2-x_try06575) */
              if ErrorIn(va_arg2_try06587) {x_try06575 = va_arg2_try06587
              } else {
              va_arg2 = ToSymbol(OBJ(va_arg2_try06587))
              /* ---------- now we compile update mClaire/pname(va_arg1) := va_arg2 ------- */
              va_arg1.Pname = va_arg2
              x_try06575 = EID{va_arg2.Id(),0}
              }
              /* update-6 */} 
            /* ERROR PROTECTION INSERTED (x_try06575-x_try06575) */
            if !ErrorIn(x_try06575) {
            _CL_obj.Range = R.Domain
            x_try06575 = EID{_CL_obj.Id(),0}
            }
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (x-Result) */
          if ErrorIn(x_try06575) {Result = x_try06575
          } else {
          x = To_Variable(OBJ(x_try06575))
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
            var g0659I *ClaireBoolean  
            if (y1.Isa.IsIn(C_Call) == CTRUE) /* If:6 */{ 
              /* Let:7 */{ 
                var g0653 *Call   = To_Call(y1)
                /* noccur = 1 */
                g0659I = Equal(g0653.Selector.Id(),C__inf_dash.Id())
                /* Let-7 */} 
              } else {
              g0659I = CFALSE
              /* If-6 */} 
            if (g0659I == CTRUE) /* If:6 */{ 
              /* Construct:7 */{ 
                var v_bag_arg *ClaireAny  
                Result= EID{MakeEmptyTuple().Id(),0}
                ToTuple(OBJ(Result)).AddFast(R.Id())
                var v_bag_arg_try06608 EID 
                /* Construct:8 */{ 
                  var v_bag_arg *ClaireAny  
                  v_bag_arg_try06608= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                  ToList(OBJ(v_bag_arg_try06608)).AddFast(x.Id())
                  var v_bag_arg_try06619 EID 
                  /* Let:9 */{ 
                    var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
                    /* noccur = 5 */
                    /* update:10 */{ 
                      var va_arg1 *ClaireVariable  
                      var va_arg2 *ClaireSymbol  
                      va_arg1 = _CL_obj
                      var va_arg2_try066211 EID 
                      va_arg2_try066211 = F_extract_symbol_any(ToList(OBJ(Core.F_CALL(C_args,ARGS(y1.ToEID())))).At(1-1))
                      /* ERROR PROTECTION INSERTED (va_arg2-v_bag_arg_try06619) */
                      if ErrorIn(va_arg2_try066211) {v_bag_arg_try06619 = va_arg2_try066211
                      } else {
                      va_arg2 = ToSymbol(OBJ(va_arg2_try066211))
                      /* ---------- now we compile update mClaire/pname(va_arg1) := va_arg2 ------- */
                      va_arg1.Pname = va_arg2
                      v_bag_arg_try06619 = EID{va_arg2.Id(),0}
                      }
                      /* update-10 */} 
                    /* ERROR PROTECTION INSERTED (v_bag_arg_try06619-v_bag_arg_try06619) */
                    if !ErrorIn(v_bag_arg_try06619) {
                    _CL_obj.Range = R.Range
                    v_bag_arg_try06619 = EID{_CL_obj.Id(),0}
                    }
                    /* Let-9 */} 
                  /* ERROR PROTECTION INSERTED (v_bag_arg-v_bag_arg_try06608) */
                  if ErrorIn(v_bag_arg_try06619) {v_bag_arg_try06608 = v_bag_arg_try06619
                  } else {
                  v_bag_arg = ANY(v_bag_arg_try06619)
                  ToList(OBJ(v_bag_arg_try06608)).AddFast(v_bag_arg)
                  var v_bag_arg_try06639 EID 
                  /* Let:9 */{ 
                    var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
                    /* noccur = 5 */
                    /* update:10 */{ 
                      var va_arg1 *ClaireVariable  
                      var va_arg2 *ClaireSymbol  
                      va_arg1 = _CL_obj
                      var va_arg2_try066411 EID 
                      va_arg2_try066411 = F_extract_symbol_any(ToList(OBJ(Core.F_CALL(C_args,ARGS(y1.ToEID())))).At(2-1))
                      /* ERROR PROTECTION INSERTED (va_arg2-v_bag_arg_try06639) */
                      if ErrorIn(va_arg2_try066411) {v_bag_arg_try06639 = va_arg2_try066411
                      } else {
                      va_arg2 = ToSymbol(OBJ(va_arg2_try066411))
                      /* ---------- now we compile update mClaire/pname(va_arg1) := va_arg2 ------- */
                      va_arg1.Pname = va_arg2
                      v_bag_arg_try06639 = EID{va_arg2.Id(),0}
                      }
                      /* update-10 */} 
                    /* ERROR PROTECTION INSERTED (v_bag_arg_try06639-v_bag_arg_try06639) */
                    if !ErrorIn(v_bag_arg_try06639) {
                    _CL_obj.Range = R.Range
                    v_bag_arg_try06639 = EID{_CL_obj.Id(),0}
                    }
                    /* Let-9 */} 
                  /* ERROR PROTECTION INSERTED (v_bag_arg-v_bag_arg_try06608) */
                  if ErrorIn(v_bag_arg_try06639) {v_bag_arg_try06608 = v_bag_arg_try06639
                  } else {
                  v_bag_arg = ANY(v_bag_arg_try06639)
                  ToList(OBJ(v_bag_arg_try06608)).AddFast(v_bag_arg)}}
                  /* Construct-8 */} 
                /* ERROR PROTECTION INSERTED (v_bag_arg-Result) */
                if ErrorIn(v_bag_arg_try06608) {Result = v_bag_arg_try06608
                } else {
                v_bag_arg = ANY(v_bag_arg_try06608)
                ToTuple(OBJ(Result)).AddFast(v_bag_arg)}
                /* Construct-7 */} 
              } else {
              /* Construct:7 */{ 
                var v_bag_arg *ClaireAny  
                Result= EID{MakeEmptyTuple().Id(),0}
                ToTuple(OBJ(Result)).AddFast(R.Id())
                var v_bag_arg_try06658 EID 
                /* Construct:8 */{ 
                  var v_bag_arg *ClaireAny  
                  v_bag_arg_try06658= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                  ToList(OBJ(v_bag_arg_try06658)).AddFast(x.Id())
                  var v_bag_arg_try06669 EID 
                  /* Let:9 */{ 
                    var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
                    /* noccur = 5 */
                    /* update:10 */{ 
                      var va_arg1 *ClaireVariable  
                      var va_arg2 *ClaireSymbol  
                      va_arg1 = _CL_obj
                      var va_arg2_try066711 EID 
                      va_arg2_try066711 = F_extract_symbol_any(y1)
                      /* ERROR PROTECTION INSERTED (va_arg2-v_bag_arg_try06669) */
                      if ErrorIn(va_arg2_try066711) {v_bag_arg_try06669 = va_arg2_try066711
                      } else {
                      va_arg2 = ToSymbol(OBJ(va_arg2_try066711))
                      /* ---------- now we compile update mClaire/pname(va_arg1) := va_arg2 ------- */
                      va_arg1.Pname = va_arg2
                      v_bag_arg_try06669 = EID{va_arg2.Id(),0}
                      }
                      /* update-10 */} 
                    /* ERROR PROTECTION INSERTED (v_bag_arg_try06669-v_bag_arg_try06669) */
                    if !ErrorIn(v_bag_arg_try06669) {
                    _CL_obj.Range = F_safeRange_relation(R)
                    v_bag_arg_try06669 = EID{_CL_obj.Id(),0}
                    }
                    /* Let-9 */} 
                  /* ERROR PROTECTION INSERTED (v_bag_arg-v_bag_arg_try06658) */
                  if ErrorIn(v_bag_arg_try06669) {v_bag_arg_try06658 = v_bag_arg_try06669
                  } else {
                  v_bag_arg = ANY(v_bag_arg_try06669)
                  ToList(OBJ(v_bag_arg_try06658)).AddFast(v_bag_arg)
                  /* Let:9 */{ 
                    var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
                    /* noccur = 5 */
                    _CL_obj.Pname = Core.F_gensym_void()
                    _CL_obj.Range = F_safeRange_relation(R)
                    v_bag_arg = _CL_obj.Id()
                    /* Let-9 */} 
                  ToList(OBJ(v_bag_arg_try06658)).AddFast(v_bag_arg)}
                  /* Construct-8 */} 
                /* ERROR PROTECTION INSERTED (v_bag_arg-Result) */
                if ErrorIn(v_bag_arg_try06658) {Result = v_bag_arg_try06658
                } else {
                v_bag_arg = ANY(v_bag_arg_try06658)
                ToTuple(OBJ(Result)).AddFast(v_bag_arg)}
                /* Construct-7 */} 
              /* If-6 */} 
            }
            /* Let-5 */} 
          }
          /* Let-4 */} 
        /* Let-3 */} 
      } else {
      var g0668I *ClaireBoolean  
      if (c.Isa.IsIn(C_Call) == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var g0654 *Call   = To_Call(c)
          /* noccur = 2 */
          g0668I = MakeBoolean((g0654.Selector.Id() == C_add.Id()) && (g0654.Args.At(1-1).Isa.IsIn(C_relation) == CTRUE))
          /* Let-4 */} 
        } else {
        g0668I = CFALSE
        /* If-3 */} 
      if (g0668I == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var R *ClaireRelation   = ToRelation(ToList(OBJ(Core.F_CALL(C_args,ARGS(c.ToEID())))).At(1-1))
          /* noccur = 3 */
          /* Let:5 */{ 
            var x *ClaireVariable  
            /* noccur = 1 */
            var x_try06696 EID 
            /* Let:6 */{ 
              var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
              /* noccur = 5 */
              /* update:7 */{ 
                var va_arg1 *ClaireVariable  
                var va_arg2 *ClaireSymbol  
                va_arg1 = _CL_obj
                var va_arg2_try06708 EID 
                va_arg2_try06708 = F_extract_symbol_any(ToList(OBJ(Core.F_CALL(C_args,ARGS(c.ToEID())))).At(2-1))
                /* ERROR PROTECTION INSERTED (va_arg2-x_try06696) */
                if ErrorIn(va_arg2_try06708) {x_try06696 = va_arg2_try06708
                } else {
                va_arg2 = ToSymbol(OBJ(va_arg2_try06708))
                /* ---------- now we compile update mClaire/pname(va_arg1) := va_arg2 ------- */
                va_arg1.Pname = va_arg2
                x_try06696 = EID{va_arg2.Id(),0}
                }
                /* update-7 */} 
              /* ERROR PROTECTION INSERTED (x_try06696-x_try06696) */
              if !ErrorIn(x_try06696) {
              _CL_obj.Range = R.Domain
              x_try06696 = EID{_CL_obj.Id(),0}
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (x-Result) */
            if ErrorIn(x_try06696) {Result = x_try06696
            } else {
            x = To_Variable(OBJ(x_try06696))
            /* Let:6 */{ 
              var y *ClaireVariable  
              /* noccur = 1 */
              var y_try06717 EID 
              /* Let:7 */{ 
                var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
                /* noccur = 5 */
                /* update:8 */{ 
                  var va_arg1 *ClaireVariable  
                  var va_arg2 *ClaireSymbol  
                  va_arg1 = _CL_obj
                  var va_arg2_try06729 EID 
                  va_arg2_try06729 = F_extract_symbol_any(ToList(OBJ(Core.F_CALL(C_args,ARGS(c.ToEID())))).At(3-1))
                  /* ERROR PROTECTION INSERTED (va_arg2-y_try06717) */
                  if ErrorIn(va_arg2_try06729) {y_try06717 = va_arg2_try06729
                  } else {
                  va_arg2 = ToSymbol(OBJ(va_arg2_try06729))
                  /* ---------- now we compile update mClaire/pname(va_arg1) := va_arg2 ------- */
                  va_arg1.Pname = va_arg2
                  y_try06717 = EID{va_arg2.Id(),0}
                  }
                  /* update-8 */} 
                /* ERROR PROTECTION INSERTED (y_try06717-y_try06717) */
                if !ErrorIn(y_try06717) {
                _CL_obj.Range = R.Range
                y_try06717 = EID{_CL_obj.Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (y-Result) */
              if ErrorIn(y_try06717) {Result = y_try06717
              } else {
              y = To_Variable(OBJ(y_try06717))
              Result = EID{MakeTuple(R.Id(),MakeConstantList(x.Id(),y.Id()).Id()).Id(),0}
              }
              /* Let-6 */} 
            }
            /* Let-5 */} 
          /* Let-4 */} 
        } else {
        var g0673I *ClaireBoolean  
        if (c.Isa.IsIn(C_Call) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var g0655 *Call   = To_Call(c)
            /* noccur = 1 */
            g0673I = Equal(MakeInteger(g0655.Args.Length()).Id(),MakeInteger(2).Id())
            /* Let-5 */} 
          } else {
          g0673I = CFALSE
          /* If-4 */} 
        if (g0673I == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var R *ClaireProperty   = To_Call(c).Selector
            /* noccur = 3 */
            /* Let:6 */{ 
              var x *ClaireVariable  
              /* noccur = 1 */
              var x_try06747 EID 
              /* Let:7 */{ 
                var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
                /* noccur = 5 */
                /* update:8 */{ 
                  var va_arg1 *ClaireVariable  
                  var va_arg2 *ClaireSymbol  
                  va_arg1 = _CL_obj
                  var va_arg2_try06759 EID 
                  va_arg2_try06759 = F_extract_symbol_any(ToList(OBJ(Core.F_CALL(C_args,ARGS(c.ToEID())))).At(1-1))
                  /* ERROR PROTECTION INSERTED (va_arg2-x_try06747) */
                  if ErrorIn(va_arg2_try06759) {x_try06747 = va_arg2_try06759
                  } else {
                  va_arg2 = ToSymbol(OBJ(va_arg2_try06759))
                  /* ---------- now we compile update mClaire/pname(va_arg1) := va_arg2 ------- */
                  va_arg1.Pname = va_arg2
                  x_try06747 = EID{va_arg2.Id(),0}
                  }
                  /* update-8 */} 
                /* ERROR PROTECTION INSERTED (x_try06747-x_try06747) */
                if !ErrorIn(x_try06747) {
                _CL_obj.Range = R.Domain
                x_try06747 = EID{_CL_obj.Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (x-Result) */
              if ErrorIn(x_try06747) {Result = x_try06747
              } else {
              x = To_Variable(OBJ(x_try06747))
              /* Let:7 */{ 
                var y *ClaireVariable  
                /* noccur = 1 */
                var y_try06768 EID 
                /* Let:8 */{ 
                  var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
                  /* noccur = 5 */
                  /* update:9 */{ 
                    var va_arg1 *ClaireVariable  
                    var va_arg2 *ClaireSymbol  
                    va_arg1 = _CL_obj
                    var va_arg2_try067710 EID 
                    va_arg2_try067710 = F_extract_symbol_any(ToList(OBJ(Core.F_CALL(C_args,ARGS(c.ToEID())))).At(2-1))
                    /* ERROR PROTECTION INSERTED (va_arg2-y_try06768) */
                    if ErrorIn(va_arg2_try067710) {y_try06768 = va_arg2_try067710
                    } else {
                    va_arg2 = ToSymbol(OBJ(va_arg2_try067710))
                    /* ---------- now we compile update mClaire/pname(va_arg1) := va_arg2 ------- */
                    va_arg1.Pname = va_arg2
                    y_try06768 = EID{va_arg2.Id(),0}
                    }
                    /* update-9 */} 
                  /* ERROR PROTECTION INSERTED (y_try06768-y_try06768) */
                  if !ErrorIn(y_try06768) {
                  _CL_obj.Range = R.Range
                  y_try06768 = EID{_CL_obj.Id(),0}
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (y-Result) */
                if ErrorIn(y_try06768) {Result = y_try06768
                } else {
                y = To_Variable(OBJ(y_try06768))
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
                var g0678 *And   = To_And(cond)
                /* noccur = 3 */
                var _Ztest_try06828 EID 
                if (g0678.Args.Length() > 2) /* If:8 */{ 
                  /* Let:9 */{ 
                    var _CL_obj *And   = To_And(new(And).Is(C_And))
                    /* noccur = 3 */
                    /* update:10 */{ 
                      var va_arg1 *And  
                      var va_arg2 *ClaireList  
                      va_arg1 = _CL_obj
                      var va_arg2_try068311 EID 
                      va_arg2_try068311 = g0678.Args.Cdr()
                      /* ERROR PROTECTION INSERTED (va_arg2-_Ztest_try06828) */
                      if ErrorIn(va_arg2_try068311) {_Ztest_try06828 = va_arg2_try068311
                      } else {
                      va_arg2 = ToList(OBJ(va_arg2_try068311))
                      /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                      va_arg1.Args = va_arg2
                      _Ztest_try06828 = EID{va_arg2.Id(),0}
                      }
                      /* update-10 */} 
                    /* ERROR PROTECTION INSERTED (_Ztest_try06828-_Ztest_try06828) */
                    if !ErrorIn(_Ztest_try06828) {
                    _Ztest_try06828 = EID{_CL_obj.Id(),0}
                    }
                    /* Let-9 */} 
                  } else {
                  _Ztest_try06828 = g0678.Args.At(2-1).ToEID()
                  /* If-8 */} 
                /* ERROR PROTECTION INSERTED (_Ztest-Result) */
                if ErrorIn(_Ztest_try06828) {Result = _Ztest_try06828
                } else {
                _Ztest = ANY(_Ztest_try06828)
                Result = _Ztest.ToEID()
                }
                /* Let-7 */} 
              } else {
              _Zbody = conc
              Result = _Zbody.ToEID()
              /* If-6 */} 
            /* If!5 */}  else if (cond.Isa.IsIn(C_And) == CTRUE) /* If:5 */{ 
            /* Let:6 */{ 
              var g0680 *And   = To_And(cond)
              /* noccur = 1 */
              var _Ztest_try06847 EID 
              /* Let:7 */{ 
                var _CL_obj *And   = To_And(new(And).Is(C_And))
                /* noccur = 3 */
                /* update:8 */{ 
                  var va_arg1 *And  
                  var va_arg2 *ClaireList  
                  va_arg1 = _CL_obj
                  var va_arg2_try06859 EID 
                  /* Let:9 */{ 
                    var g0686UU *ClaireList  
                    /* noccur = 1 */
                    var g0686UU_try068710 EID 
                    g0686UU_try068710 = g0680.Args.Cdr()
                    /* ERROR PROTECTION INSERTED (g0686UU-va_arg2_try06859) */
                    if ErrorIn(g0686UU_try068710) {va_arg2_try06859 = g0686UU_try068710
                    } else {
                    g0686UU = ToList(OBJ(g0686UU_try068710))
                    va_arg2_try06859 = EID{MakeConstantList(_Ztest).Append(g0686UU).Id(),0}
                    }
                    /* Let-9 */} 
                  /* ERROR PROTECTION INSERTED (va_arg2-_Ztest_try06847) */
                  if ErrorIn(va_arg2_try06859) {_Ztest_try06847 = va_arg2_try06859
                  } else {
                  va_arg2 = ToList(OBJ(va_arg2_try06859))
                  /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                  va_arg1.Args = va_arg2
                  _Ztest_try06847 = EID{va_arg2.Id(),0}
                  }
                  /* update-8 */} 
                /* ERROR PROTECTION INSERTED (_Ztest_try06847-_Ztest_try06847) */
                if !ErrorIn(_Ztest_try06847) {
                _Ztest_try06847 = EID{_CL_obj.Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (_Ztest-Result) */
              if ErrorIn(_Ztest_try06847) {Result = _Ztest_try06847
              } else {
              _Ztest = ANY(_Ztest_try06847)
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
              var g0681 *If   = To_If(_Zbody)
              /* noccur = 2 */
              g0681.Test = _Ztest
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
              var va_arg2_try06887 EID 
              va_arg2_try06887 = F_lambda_I_list(lvar,_Zbody)
              /* ERROR PROTECTION INSERTED (va_arg2-Result) */
              if ErrorIn(va_arg2_try06887) {Result = va_arg2_try06887
              } else {
              va_arg2 = ToLambda(OBJ(va_arg2_try06887))
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
      var g0689 *ClaireProperty   = ToProperty(x.Id())
      /* noccur = 2 */
      var g0693I *ClaireBoolean  
      /* Let:3 */{ 
        var g0694UU *ClaireAny  
        /* noccur = 1 */
        /* For:4 */{ 
          var s *ClaireAny  
          _ = s
          g0694UU= CFALSE.Id()
          for _,s = range(g0689.Restrictions.ValuesO())/* loop:5 */{ 
            var g0695I *ClaireBoolean  
            /* Let:6 */{ 
              var g0696UU *ClaireBoolean  
              /* noccur = 1 */
              /* Let:7 */{ 
                var g0697UU *ClaireBoolean  
                /* noccur = 1 */
                if (C_slot.Id() == s.Isa.Id()) /* If:8 */{ 
                  /* Let:9 */{ 
                    var g0690 *ClaireSlot   = ToSlot(s)
                    /* noccur = 2 */
                    g0697UU = g0690.Range.Contains(g0690.Default)
                    /* Let-9 */} 
                  } else {
                  g0697UU = CFALSE
                  /* If-8 */} 
                g0696UU = F_boolean_I_any(g0697UU.Id())
                /* Let-7 */} 
              g0695I = Core.F__I_equal_any(g0696UU.Id(),CTRUE.Id())
              /* Let-6 */} 
            if (g0695I == CTRUE) /* If:6 */{ 
               /*v = g0694UU, s =any*/
g0694UU = CTRUE.Id()
              break
              /* If-6 */} 
            /* loop-5 */} 
          /* For-4 */} 
        g0693I = Core.F_not_any(g0694UU)
        /* Let-3 */} 
      if (g0693I == CTRUE) /* If:3 */{ 
        Result = g0689.Range
        } else {
        Result = ToType(C_any.Id())
        /* If-3 */} 
      /* Let-2 */} 
    /* If!1 */}  else if (C_table.Id() == x.Isa.Id()) /* If:1 */{ 
    /* Let:2 */{ 
      var g0691 *ClaireTable   = ToTable(x.Id())
      /* noccur = 3 */
      if (g0691.Range.Contains(g0691.Default) == CTRUE) /* If:3 */{ 
        Result = g0691.Range
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
    var l_try06982 EID 
    l_try06982 = Core.F_nth_table1(C_demons,R.Id())
    /* ERROR PROTECTION INSERTED (l-Result) */
    if ErrorIn(l_try06982) {Result = l_try06982
    } else {
    l = ANY(l_try06982)
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
                  var g0699UU *Call  
                  /* noccur = 1 */
                  /* Let:9 */{ 
                    var _CL_obj *Call   = To_Call(new(Call).Is(C_Call))
                    /* noccur = 5 */
                    _CL_obj.Selector = Core.C_Core_update_dash
                    _CL_obj.Args = MakeConstantList(R.Inverse.Id(),lvar.At(3-1),lvar.At(1-1))
                    g0699UU = _CL_obj
                    /* Let-9 */} 
                  l1 = l1.AddFast(g0699UU.Id())
                  /* Let-8 */} 
                /* If-7 */} 
              l1 = l1.AddFast(F_putCall_relation2(R.Inverse,lvar.At(2-1),lvar.At(1-1)).Id())
              /* If-6 */} 
            /* update:6 */{ 
              var va_arg1 *ClaireRelation  
              var va_arg2 *ClaireAny  
              va_arg1 = R
              var va_arg2_try07007 EID 
              /* Let:7 */{ 
                var g0701UU *ComplexInstruction  
                /* noccur = 1 */
                if (F_eventMethod_ask_relation2(R) == CTRUE) /* If:8 */{ 
                  /* Let:9 */{ 
                    var _CL_obj *Do   = To_Do(new(Do).Is(C_Do))
                    /* noccur = 3 */
                    _CL_obj.Args = l2
                    g0701UU = To_ComplexInstruction(_CL_obj.Id())
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
                    g0701UU = To_ComplexInstruction(_CL_obj.Id())
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
                    g0701UU = To_ComplexInstruction(_CL_obj.Id())
                    /* Let-9 */} 
                  /* If-8 */} 
                va_arg2_try07007 = F_lambda_I_list(MakeConstantList(lvar.At(1-1),lvar.At(2-1)),g0701UU.Id())
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (va_arg2-Result) */
              if ErrorIn(va_arg2_try07007) {Result = va_arg2_try07007
              } else {
              va_arg2 = ANY(va_arg2_try07007)
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
      var g0702 *ClaireList   = ToList(self)
      /* noccur = 1 */
      /* For:3 */{ 
        var x *ClaireAny  
        _ = x
        Result= EID{CFALSE.Id(),0}
        var x_support *ClaireList  
        x_support = g0702
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
      var g0703 *Vardef   = To_Vardef(self)
      /* noccur = 2 */
      g0703.Isa = C_Variable
      Result = EID{g0703.Id(),0}
      /* Let-2 */} 
    /* If!1 */}  else if (self.Isa.IsIn(C_lambda) == CTRUE) /* If:1 */{ 
    /* Let:2 */{ 
      var g0704 *ClaireLambda   = ToLambda(self)
      /* noccur = 2 */
      Result = F_Language_jito_any(g0704.Body)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = EID{g0704.Id(),0}
      }
      /* Let-2 */} 
    /* If!1 */}  else if (self.Isa.IsIn(C_And) == CTRUE) /* If:1 */{ 
    /* Let:2 */{ 
      var g0705 *And   = To_And(self)
      /* noccur = 1 */
      Result = F_Language_jito_any(g0705.Args.Id())
      /* Let-2 */} 
    /* If!1 */}  else if (self.Isa.IsIn(C_Or) == CTRUE) /* If:1 */{ 
    /* Let:2 */{ 
      var g0706 *Or   = To_Or(self)
      /* noccur = 1 */
      Result = F_Language_jito_any(g0706.Args.Id())
      /* Let-2 */} 
    /* If!1 */}  else if (self.Isa.IsIn(C_Call) == CTRUE) /* If:1 */{ 
    /* Let:2 */{ 
      var g0707 *Call   = To_Call(self)
      /* noccur = 1 */
      Result = g0707.MakeJito()
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = EID{CTRUE.Id(),0}
      }
      /* Let-2 */} 
    /* If!1 */}  else if (self.Isa.IsIn(C_Let) == CTRUE) /* If:1 */{ 
    /* Let:2 */{ 
      var g0708 *Let   = To_Let(self)
      /* noccur = 1 */
      Result = g0708.LetJito()
      /* Let-2 */} 
    /* If!1 */}  else if (self.Isa.IsIn(C_Assign) == CTRUE) /* If:1 */{ 
    /* Let:2 */{ 
      var g0709 *Assign   = To_Assign(self)
      /* noccur = 3 */
      if (g0709.ClaireVar.Isa.IsIn(C_Variable) != CTRUE) /* If:3 */{ 
        Result = ToException(Core.C_general_error.Make(MakeString("[101] ~S is not a variable").Id(),MakeConstantList(g0709.ClaireVar).Id())).Close()
        } else {
        Result = EID{CFALSE.Id(),0}
        /* If-3 */} 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = F_Language_jito_any(g0709.Arg)
      }
      /* Let-2 */} 
    /* If!1 */}  else if (self.Isa.IsIn(C_Gassign) == CTRUE) /* If:1 */{ 
    /* Let:2 */{ 
      var g0710 *Gassign   = To_Gassign(self)
      /* noccur = 3 */
      if (g0710.ClaireVar.Range.Contains(g0710.ClaireVar.Value) == CTRUE) /* If:3 */{ 
        Result = F_Language_jito_any(g0710.Arg)
        } else {
        Result = EID{CFALSE.Id(),0}
        /* If-3 */} 
      /* Let-2 */} 
    /* If!1 */}  else if (self.Isa.IsIn(C_Do) == CTRUE) /* If:1 */{ 
    /* Let:2 */{ 
      var g0711 *Do   = To_Do(self)
      /* noccur = 1 */
      Result = F_Language_jito_any(g0711.Args.Id())
      /* Let-2 */} 
    /* If!1 */}  else if (self.Isa.IsIn(C_If) == CTRUE) /* If:1 */{ 
    /* Let:2 */{ 
      var g0712 *If   = To_If(self)
      /* noccur = 3 */
      Result = F_Language_jito_any(g0712.Arg)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = F_Language_jito_any(g0712.Test)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = F_Language_jito_any(g0712.Other)
      }}
      /* Let-2 */} 
    /* If!1 */}  else if (self.Isa.IsIn(C_Iteration) == CTRUE) /* If:1 */{ 
    /* Let:2 */{ 
      var g0713 *Iteration   = To_Iteration(self)
      /* noccur = 5 */
      /* Let:3 */{ 
        var v *ClaireVariable   = g0713.ClaireVar
        /* noccur = 4 */
        /* Let:4 */{ 
          var s *ClaireAny   = g0713.SetArg
          /* noccur = 3 */
          /* Let:5 */{ 
            var o_ask *ClaireBoolean  
            /* noccur = 2 */
            /* and:6 */{ 
              var v_and6 *ClaireBoolean  
              
              if (s.Isa.IsIn(C_Call) == CTRUE) /* If:7 */{ 
                /* Let:8 */{ 
                  var g0714 *Call   = To_Call(s)
                  /* noccur = 1 */
                  v_and6 = Equal(g0714.Selector.Id(),C__dot_dot.Id())
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
            /* Let:6 */{ 
              var g0720UU *ClaireList  
              /* noccur = 1 */
              var g0720UU_try07217 EID 
              /* Construct:7 */{ 
                var v_bag_arg *ClaireAny  
                g0720UU_try07217= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                ToList(OBJ(g0720UU_try07217)).AddFast(g0713.Id())
                var v_bag_arg_try07228 EID 
                v_bag_arg_try07228 = F_static_type_any(g0713.SetArg)
                /* ERROR PROTECTION INSERTED (v_bag_arg-g0720UU_try07217) */
                if ErrorIn(v_bag_arg_try07228) {g0720UU_try07217 = v_bag_arg_try07228
                } else {
                v_bag_arg = ANY(v_bag_arg_try07228)
                ToList(OBJ(g0720UU_try07217)).AddFast(v_bag_arg)}
                /* Construct-7 */} 
              /* ERROR PROTECTION INSERTED (g0720UU-Result) */
              if ErrorIn(g0720UU_try07217) {Result = g0720UU_try07217
              } else {
              g0720UU = ToList(OBJ(g0720UU_try07217))
              Result = Core.F_tformat_string(MakeString("-- Iteration jito: ~S (~S)\n"),3,g0720UU)
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            if (o_ask == CTRUE) /* If:6 */{ 
              v.Range = ToType(C_integer.Id())
              Core.F_tformat_string(MakeString("-- jito:put range ~S as integer\n"),3,MakeConstantList(v.Id()))
              /* If-6 */} 
            Core.F_CALL(C_Language_ofto,ARGS(s.ToEID()))
            Result = F_Language_jito_any(g0713.Arg)
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
            }}
            /* Let-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    /* If!1 */}  else if (self.Isa.IsIn(C_Construct) == CTRUE) /* If:1 */{ 
    /* Let:2 */{ 
      var g0715 *Construct   = To_Construct(self)
      /* noccur = 2 */
      Core.F_tformat_string(MakeString("-- Construct jito: ~S\n"),3,MakeConstantList(g0715.Id()))
      Result = F_Language_jito_any(g0715.Args.Id())
      /* Let-2 */} 
    /* If!1 */}  else if (self.Isa.IsIn(C_Exists) == CTRUE) /* If:1 */{ 
    /* Let:2 */{ 
      var g0716 *Exists   = To_Exists(self)
      /* noccur = 3 */
      Result = F_Language_jito_any(g0716.SetArg)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = F_Language_jito_any(g0716.Arg)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = F_Language_jito_any(g0716.Other)
      }}
      /* Let-2 */} 
    /* If!1 */}  else if (self.Isa.IsIn(C_Handle) == CTRUE) /* If:1 */{ 
    /* Let:2 */{ 
      var g0717 *ClaireHandle   = To_ClaireHandle(self)
      /* noccur = 4 */
      if (C_class.Id() != g0717.Test.Isa.Id()) /* If:3 */{ 
        Result = ToException(Core.C_general_error.Make(MakeString("syntax: [try %S] must use a class").Id(),MakeConstantList(g0717.Test).Id())).Close()
        } else {
        Result = EID{CFALSE.Id(),0}
        /* If-3 */} 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = F_Language_jito_any(g0717.Arg)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = F_Language_jito_any(g0717.Other)
      }}
      /* Let-2 */} 
    /* If!1 */}  else if (self.Isa.IsIn(C_Definition) == CTRUE) /* If:1 */{ 
    /* Let:2 */{ 
      var g0718 *Definition   = To_Definition(self)
      /* noccur = 2 */
      if (F_boolean_I_any(ANY(Core.F_CALL(C_Language_fast_definition,ARGS(EID{g0718.Arg.Id(),0})))) == CTRUE) /* If:3 */{ 
        /* update:4 */{ 
          var va_arg1 *ClaireAny  
          var va_arg2 *ClaireClass  
          va_arg1 = g0718.Id()
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
    /* noccur = 7 */
    /* Let:2 */{ 
      var x *ClaireAny   = self.Value
      /* noccur = 5 */
      /* Let:3 */{ 
        var untyped *ClaireBoolean   = MakeBoolean((v.Range.Id() == CNULL))
        /* noccur = 3 */
        Core.F_tformat_string(MakeString("Let Jito with var ~S => ~S\n"),3,MakeConstantList(v.Id(),untyped.Id()))
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
              var va_arg2_try07237 EID 
              va_arg2_try07237 = F_static_type_any(x)
              /* ERROR PROTECTION INSERTED (va_arg2-Result) */
              if ErrorIn(va_arg2_try07237) {Result = va_arg2_try07237
              } else {
              va_arg2 = ToType(OBJ(va_arg2_try07237))
              /* ---------- now we compile update range(va_arg1) := va_arg2 ------- */
              va_arg1.Range = va_arg2
              Result = EID{va_arg2.Id(),0}
              }
              /* update-6 */} 
            /* If-5 */} 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          Result = Core.F_tformat_string(MakeString("--- let Jito ~S:~S (~S)\n"),3,MakeConstantList(v.Id(),v.Range.Id(),x))
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
          var g0728I *ClaireBoolean  
          /* and:5 */{ 
            var v_and5 *ClaireBoolean  
            
            v_and5 = Equal(p.Id(),Core.C_write.Id())
            if (v_and5 == CFALSE) {g0728I = CFALSE
            } else /* arg:6 */{ 
              /* Let:7 */{ 
                var p2 *ClaireAny   = self.Args.At(1-1)
                /* noccur = 2 */
                if (p2.Isa.IsIn(C_property) == CTRUE) /* If:8 */{ 
                  /* Let:9 */{ 
                    var g0724 *ClaireProperty   = ToProperty(p2)
                    /* noccur = 3 */
                    v_and5 = MakeBoolean((g0724.Inverse.Id() == CNULL) && (g0724.Store_ask != CTRUE) && (g0724.IfWrite == CNULL))
                    /* Let-9 */} 
                  } else {
                  v_and5 = CFALSE
                  /* If-8 */} 
                /* Let-7 */} 
              if (v_and5 == CFALSE) {g0728I = CFALSE
              } else /* arg:7 */{ 
                g0728I = CTRUE/* arg-7 */} 
              /* arg-6 */} 
            /* and-5 */} 
          if (g0728I == CTRUE) /* If:5 */{ 
            p = C_write_fast
            self.Selector = C_write_fast
            /* If-5 */} 
          var g0729I *ClaireBoolean  
          /* and:5 */{ 
            var v_and5 *ClaireBoolean  
            
            v_and5 = Core.F__inf_equal_integer(p.Open,1)
            if (v_and5 == CFALSE) {g0729I = CFALSE
            } else /* arg:6 */{ 
              v_and5 = Core.F__inf_equal_integer(p.Restrictions.Length(),10)
              if (v_and5 == CFALSE) {g0729I = CFALSE
              } else /* arg:7 */{ 
                /* Let:8 */{ 
                  var g0730UU *ClaireAny  
                  /* noccur = 1 */
                  /* For:9 */{ 
                    var x *ClaireAny  
                    _ = x
                    g0730UU= CFALSE.Id()
                    for _,x = range(p.Restrictions.ValuesO())/* loop:10 */{ 
                      var g0731I *ClaireBoolean  
                      /* Let:11 */{ 
                        var g0732UU *ClaireBoolean  
                        /* noccur = 1 */
                        /* Let:12 */{ 
                          var g0733UU *ClaireAny  
                          /* noccur = 1 */
                          /* For:13 */{ 
                            var t *ClaireAny  
                            _ = t
                            g0733UU= CFALSE.Id()
                            for _,t = range(ToRestriction(x).Domain.ValuesO())/* loop:14 */{ 
                              if (C_class.Id() != t.Isa.Id()) /* If:15 */{ 
                                 /*v = g0733UU, s =any*/
g0733UU = CTRUE.Id()
                                break
                                /* If-15 */} 
                              /* loop-14 */} 
                            /* For-13 */} 
                          g0732UU = Core.F_not_any(g0733UU)
                          /* Let-12 */} 
                        g0731I = g0732UU.Not
                        /* Let-11 */} 
                      if (g0731I == CTRUE) /* If:11 */{ 
                         /*v = g0730UU, s =any*/
g0730UU = CTRUE.Id()
                        break
                        /* If-11 */} 
                      /* loop-10 */} 
                    /* For-9 */} 
                  v_and5 = Core.F_not_any(g0730UU)
                  /* Let-8 */} 
                if (v_and5 == CFALSE) {g0729I = CFALSE
                } else /* arg:8 */{ 
                  g0729I = CTRUE/* arg-8 */} 
                /* arg-7 */} 
              /* arg-6 */} 
            /* and-5 */} 
          if (g0729I == CTRUE) /* If:5 */{ 
            /* Let:6 */{ 
              var lt *ClaireList  
              /* noccur = 2 */
              var lt_try07347 EID 
              /* Iteration:7 */{ 
                var v_list7 *ClaireList  
                var x *ClaireAny  
                var v_local7 *ClaireAny  
                v_list7 = larg
                lt_try07347 = EID{CreateList(ToType(CEMPTY.Id()),v_list7.Length()).Id(),0}
                for CLcount := 0; CLcount < v_list7.Length(); CLcount++{ 
                  x = v_list7.At(CLcount)
                  var v_local7_try07359 EID 
                  v_local7_try07359 = F_static_type_any(x)
                  /* ERROR PROTECTION INSERTED (v_local7-lt_try07347) */
                  if ErrorIn(v_local7_try07359) {lt_try07347 = v_local7_try07359
                  lt_try07347 = v_local7_try07359
                  break
                  } else {
                  v_local7 = ANY(v_local7_try07359)
                  ToList(OBJ(lt_try07347)).PutAt(CLcount,v_local7)
                  } 
                }
                /* Iteration-7 */} 
              /* ERROR PROTECTION INSERTED (lt-Result) */
              if ErrorIn(lt_try07347) {Result = lt_try07347
              } else {
              lt = ToList(OBJ(lt_try07347))
              Core.F_tformat_string(MakeString("-- call jito: ~S : ~S\n"),3,MakeConstantList(self.Id(),lt.Id()))
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
          var g0736I *ClaireBoolean  
          if (C_method.Id() == m.Isa.Id()) /* If:5 */{ 
            /* Let:6 */{ 
              var g0726 *ClaireMethod   = ToMethod(m)
              /* noccur = 1 */
              g0736I = MakeBoolean((g0726.Functional.Id() == CNULL)).Not
              /* Let-6 */} 
            } else {
            g0736I = CFALSE
            /* If-5 */} 
          if (g0736I == CTRUE) /* If:5 */{ 
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
            var g0738UU *ClaireAny  
            /* noccur = 1 */
            /* Let:6 */{ 
              var i int  = 1
              /* noccur = 5 */
              /* Let:7 */{ 
                var g0737 int  = n
                /* noccur = 1 */
                g0738UU= CFALSE.Id()
                for (i <= g0737) /* while:8 */{ 
                  if (ToType(lt.At(i-1)).Included(ToType(ld.ValuesO()[i-1])) != CTRUE) /* If:9 */{ 
                     /*v = g0738UU, s =any*/
g0738UU = CTRUE.Id()
                    break
                    /* If-9 */} 
                  i = (i+1)
                  /* while-8 */} 
                /* Let-7 */} 
              /* Let-6 */} 
            v_and3 = Core.F_not_any(g0738UU)
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