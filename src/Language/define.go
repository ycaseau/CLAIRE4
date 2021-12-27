/***** CLAIRE Compilation of file /Users/ycaseau/Dropbox/src/clairev4.03/src/meta/define.cl 
         [version 4.0.03 / safety 5] Monday 12-27-2021 10:35:23 *****/

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
/* {1} The go function for: self_print(self:Definition) [status=1] */
func (self *Definition ) SelfPrint () EID { 
    var Result EID 
    /*g_try(v2:"Result",loop:true) */
    Result = Core.F_print_any(self.Arg.Id())
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("(")
    /*g_try(v2:"Result",loop:true) */
    Result = F_Language_printbox_list2(self.Args)
    /* ERROR PROTECTION INSERTED (Result-Result) */
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
/* {1} The go function for: self_print(self:Defobj) [status=1] */
func (self *Defobj ) SelfPrint () EID { 
    var Result EID 
    if (self.Arg.Id() == Core.C_global_variable.Id()) { 
      { var r *ClaireAny   = C_any.Id()
        { var v *ClaireAny   = CNULL
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
              if (x.Args.At(1-1) == C_value.Id()) { 
                v = x.Args.At(2-1)
                }  else if (x.Args.At(1-1) == C_range.Id()) { 
                r = x.Args.At(2-1)
                } 
              } 
            } 
          if (F_boolean_I_any(r) == CTRUE) { 
            self.Ident.Princ()
            PRINC(":")
            /*g_try(v2:"Result",loop:true) */
            Result = Core.F_CALL(C_print,ARGS(r.ToEID()))
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC(" := ")
            /*g_try(v2:"Result",loop:true) */
            Result = F_printexp_any(v,CFALSE)
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC("")
            Result = EVOID
            }}
            } else {
            self.Ident.Princ()
            PRINC(" :: ")
            /*g_try(v2:"Result",loop:true) */
            Result = F_printexp_any(v,CFALSE)
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC("")
            Result = EVOID
            }
            } 
          } 
        } 
      } else {
      self.Ident.Princ()
      PRINC(" :: ")
      /*g_try(v2:"Result",loop:true) */
      Result = Core.F_CALL(C_print,ARGS(EID{self.Arg.Id(),0}))
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC("(")
      /*g_try(v2:"Result",loop:true) */
      Result = F_Language_printbox_list2(self.Args)
      /* ERROR PROTECTION INSERTED (Result-Result) */
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
/* {1} The go function for: self_print(self:Defclass) [status=1] */
func (self *Defclass ) SelfPrint () EID { 
    var Result EID 
    if (self.Ident.Id() == CNULL) { 
      Result = Core.F_print_any(MakeString("<Defclass>").Id())
      } else {
      self.Ident.Princ()
      /*g_try(v2:"Result",loop:true) */
      if (self.Params.Length() != 0) { 
        PRINC("[")
        /*g_try(v2:"Result",loop:true) */
        Result = Core.F_princ_list(self.Params)
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC("]")
        Result = EVOID
        }
        } else {
        Result = EID{CFALSE.Id(),0}
        } 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(" <: ")
      /*g_try(v2:"Result",loop:true) */
      Result = Core.F_print_any(self.Arg.Id())
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC("(")
      /*g_try(v2:"Result",loop:true) */
      { var l *ClaireList   = self.Args
        { var n int  = l.Length()
          { var i int  = 1
            { var g0213 int  = n
              _ = g0213
              Result= EID{CFALSE.Id(),0}
              for (i <= g0213) { 
                /* While stat, v:"Result" loop:true */
                var loop_1 EID 
                _ = loop_1
                { 
                /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
                /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
                if (i == 1) { 
                  F_set_level_void()
                  loop_1 = EVOID
                  } else {
                  loop_1 = F_lbreak_void()
                  } 
                /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
                if ErrorIn(loop_1) {Result = loop_1
                break
                } else {
                /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
                if (l.At(i-1).Isa.IsIn(C_Vardef) == CTRUE) { 
                  loop_1 = Core.F_CALL(C_Language_ppvariable,ARGS(l.At(i-1).ToEID()))
                  } else {
                  /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
                  loop_1 = Core.F_CALL(C_Language_ppvariable,ARGS(ToList(OBJ(Core.F_CALL(C_args,ARGS(l.At(i-1).ToEID())))).At(1-1).ToEID()))
                  /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
                  if ErrorIn(loop_1) {Result = loop_1
                  break
                  } else {
                  PRINC(" = ")
                  /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
                  loop_1 = Core.F_CALL(C_print,ARGS(ToList(OBJ(Core.F_CALL(C_args,ARGS(l.At(i-1).ToEID())))).At(2-1).ToEID()))
                  /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
                  if ErrorIn(loop_1) {Result = loop_1
                  break
                  } else {
                  PRINC("")
                  loop_1 = EVOID
                  }}
                  } 
                /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
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
                /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
                if ErrorIn(loop_1) {Result = loop_1
                break
                } else {
                i = (i+1)
                }
                /* try?:false, v2:"v_while7" loop will be:tuple("Result", EID) */
                } 
              }
              } 
            } 
          } 
        } 
      /* ERROR PROTECTION INSERTED (Result-Result) */
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
/* {1} The go function for: self_print(self:Defmethod) [status=1] */
func (self *Defmethod ) SelfPrint () EID { 
    var Result EID 
    /*g_try(v2:"Result",loop:true) */
    /*g_try(v2:"Result",loop:true) */
    Result = Core.F_print_any(self.Arg.Selector.Id())
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("(")
    /*g_try(v2:"Result",loop:true) */
    if (self.Arg.Args.Id() != CNULL) { 
      Result = F_ppvariable_list(self.Arg.Args)
      } else {
      Result = EID{CFALSE.Id(),0}
      } 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(") : ")
    /*g_try(v2:"Result",loop:true) */
    Result = F_printexp_any(self.SetArg,CFALSE)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    /*g_try(v2:"Result",loop:true) */
    /*g_try(v2:"Result",loop:true) */
    Result = F_lbreak_void()
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    { 
      var va_arg1 *Core.PrettyPrinter  
      var va_arg2 int 
      va_arg1 = Core.C_pretty
      va_arg2 = (Core.C_pretty.Index+4)
      va_arg1.Index = va_arg2
      /*integer->integer*/Result = EID{C__INT,IVAL(va_arg2)}
      } 
    }
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(" ")
    F_princ_string(ToString(IfThenElse((self.Inline_ask == CTRUE),
      MakeString("=>").Id(),
      MakeString("->").Id())))
    PRINC(" ")
    /*g_try(v2:"Result",loop:true) */
    Result = F_printexp_any(self.Body,CFALSE)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(" ")
    Result = EVOID
    }}}}}
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    { 
      var va_arg1 *Core.PrettyPrinter  
      var va_arg2 int 
      va_arg1 = Core.C_pretty
      va_arg2 = (Core.C_pretty.Index-4)
      va_arg1.Index = va_arg2
      /*integer->integer*/Result = EID{C__INT,IVAL(va_arg2)}
      } 
    }
    return Result} 
  
// The EID go function for: self_print @ Defmethod (throw: true) 
func E_self_print_Defmethod_Language (self EID) EID { 
    return To_Defmethod(OBJ(self)).SelfPrint( )} 
  
// -------------- array definition -----------------------------------
/* {1} The go function for: self_print(self:Defarray) [status=1] */
func (self *Defarray ) SelfPrint () EID { 
    var Result EID 
    /*g_try(v2:"Result",loop:true) */
    /*g_try(v2:"Result",loop:true) */
    Result = Core.F_CALL(C_print,ARGS(self.Arg.Args.At(1-1).ToEID()))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("[")
    /*g_try(v2:"Result",loop:true) */
    { var arg_1 *ClaireList  
      _ = arg_1
      var try_2 EID 
      /*g_try(v2:"try_2",loop:false) */
      try_2 = self.Arg.Args.Cdr()
      /* ERROR PROTECTION INSERTED (arg_1-Result) */
      if ErrorIn(try_2) {Result = try_2
      } else {
      arg_1 = ToList(OBJ(try_2))
      Result = F_ppvariable_list(arg_1)
      }
      } 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("] : ")
    /*g_try(v2:"Result",loop:true) */
    Result = Core.F_CALL(C_print,ARGS(self.SetArg.ToEID()))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    /*g_try(v2:"Result",loop:true) */
    /*g_try(v2:"Result",loop:true) */
    Result = F_lbreak_void()
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    { 
      var va_arg1 *Core.PrettyPrinter  
      var va_arg2 int 
      va_arg1 = Core.C_pretty
      va_arg2 = (Core.C_pretty.Index+4)
      va_arg1.Index = va_arg2
      /*integer->integer*/Result = EID{C__INT,IVAL(va_arg2)}
      } 
    }
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(" := ")
    /*g_try(v2:"Result",loop:true) */
    Result = F_printexp_any(self.Body,CFALSE)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(" ")
    Result = EVOID
    }}}}}
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    { 
      var va_arg1 *Core.PrettyPrinter  
      var va_arg2 int 
      va_arg1 = Core.C_pretty
      va_arg2 = (Core.C_pretty.Index-4)
      va_arg1.Index = va_arg2
      /*integer->integer*/Result = EID{C__INT,IVAL(va_arg2)}
      } 
    }
    return Result} 
  
// The EID go function for: self_print @ Defarray (throw: true) 
func E_self_print_Defarray_Language (self EID) EID { 
    return To_Defarray(OBJ(self)).SelfPrint( )} 
  
// -------------- rule definition ------------------------------------
/* {1} The go function for: self_print(self:Defrule) [status=1] */
func (self *Defrule ) SelfPrint () EID { 
    var Result EID 
    /*g_try(v2:"Result",loop:true) */
    self.Ident.Princ()
    PRINC("(")
    /*g_try(v2:"Result",loop:true) */
    Result = F_ppvariable_list(self.Args)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(") :: rule(")
    /*g_try(v2:"Result",loop:true) */
    Result = F_lbreak_integer(4)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(" ")
    /*g_try(v2:"Result",loop:true) */
    Result = Core.F_CALL(C_print,ARGS(self.Arg.ToEID()))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(" ")
    /*g_try(v2:"Result",loop:true) */
    Result = F_lbreak_integer(4)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("=> ")
    /*g_try(v2:"Result",loop:true) */
    Result = Core.F_CALL(C_print,ARGS(self.Body.ToEID()))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(")")
    Result = EVOID
    }}}}}
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    { 
      var va_arg1 *Core.PrettyPrinter  
      var va_arg2 int 
      va_arg1 = Core.C_pretty
      va_arg2 = (Core.C_pretty.Index-4)
      va_arg1.Index = va_arg2
      /*integer->integer*/Result = EID{C__INT,IVAL(va_arg2)}
      } 
    }
    return Result} 
  
// The EID go function for: self_print @ Defrule (throw: true) 
func E_self_print_Defrule_Language (self EID) EID { 
    return To_Defrule(OBJ(self)).SelfPrint( )} 
  
/* {1} The go function for: self_print(self:Defvar) [status=1] */
func (self *Defvar ) SelfPrint () EID { 
    var Result EID 
    /*g_try(v2:"Result",loop:true) */
    Result = F_ppvariable_Variable(self.Ident)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(" := ")
    /*g_try(v2:"Result",loop:true) */
    Result = F_printexp_any(self.Arg,CFALSE)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("")
    Result = EVOID
    }}
    return Result} 
  
// The EID go function for: self_print @ Defvar (throw: true) 
func E_self_print_Defvar_Language (self EID) EID { 
    return To_Defvar(OBJ(self)).SelfPrint( )} 
  
// *********************************************************************
// *     Part 2: the general instantiation macro                       *
// *********************************************************************
// creation of a new object
//
/* {1} The go function for: self_eval(self:Definition) [status=1] */
func (self *Definition ) SelfEval () EID { 
    var Result EID 
    { var _Zc *ClaireClass   = self.Arg
      { var _Zo *ClaireObject  
        var try_1 EID 
        /*g_try(v2:"try_1",loop:false) */
        /*g_try(v2:"try_1",loop:false) */
        if (_Zc.Open <= 1) { 
          try_1 = ToException(Core.C_general_error.Make(MakeString("[105] cannot instantiate ~S").Id(),MakeConstantList(_Zc.Id()).Id())).Close()
          } else {
          try_1 = EID{CFALSE.Id(),0}
          } 
        /* ERROR PROTECTION INSERTED (try_1-try_1) */
        if !ErrorIn(try_1) {
        try_1 = EID{F_new_object_class(_Zc).Id(),0}
        }
        /* ERROR PROTECTION INSERTED (_Zo-Result) */
        if ErrorIn(try_1) {Result = try_1
        } else {
        _Zo = ToObject(OBJ(try_1))
        { var arg_2 *ClaireList  
          _ = arg_2
          var try_3 EID 
          /*g_try(v2:"try_3",loop:false) */
          try_3 = F_Language_new_writes_object(_Zo,self.Args)
          /* ERROR PROTECTION INSERTED (arg_2-Result) */
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
/* {1} The go function for: fast_definition?(c:class) [status=0] */
func F_Language_fast_definition_ask_class (c *ClaireClass ) *ClaireBoolean  { 
    // procedure body with s = boolean 
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
            _ = arg_1
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
/* {1} The go function for: self_eval(self:DefFast) [status=1] */
func (self *Language_DefFast ) SelfEval () EID { 
    var Result EID 
    { var _Zo *ClaireObject   = F_new_object_class(self.Arg)
      /*g_try(v2:"Result",loop:true) */
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
          /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
          { var p *ClaireProperty  
            _ = p
            var try_2 EID 
            /*g_try(v2:"try_2",loop:false) */
            try_2 = F_make_a_property_any(x.Args.At(1-1))
            /* ERROR PROTECTION INSERTED (p-loop_1) */
            if ErrorIn(try_2) {loop_1 = try_2
            } else {
            p = ToProperty(OBJ(try_2))
            { 
              var arg_3 EID 
              /*g_try(v2:"arg_3",loop:false) */
              arg_3 = EVAL(x.Args.At(2-1))
              /* ERROR PROTECTION INSERTED (arg_3-loop_1) */
              if ErrorIn(arg_3) {loop_1 = arg_3
              } else {
              loop_1 = p.WriteEID(_Zo,arg_3)}
              } 
            }
            } 
          /* ERROR PROTECTION INSERTED (loop_1-Result) */
          if ErrorIn(loop_1) {Result = loop_1
          break
          } else {
          }
          } 
        } 
      /* ERROR PROTECTION INSERTED (Result-Result) */
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
/* {1} The go function for: new_writes(self:object,%l:list) [status=1] */
func F_Language_new_writes_object (self *ClaireObject ,_Zl *ClaireList ) EID { 
    var Result EID 
    { var lp *ClaireList   = ToType(CEMPTY.Id()).EmptyList()
      _ = lp
      /*g_try(v2:"Result",loop:true) */
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
          /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
          { var p *ClaireProperty  
            var try_2 EID 
            /*g_try(v2:"try_2",loop:false) */
            try_2 = F_make_a_property_any(x.Args.At(1-1))
            /* ERROR PROTECTION INSERTED (p-loop_1) */
            if ErrorIn(try_2) {loop_1 = try_2
            } else {
            p = ToProperty(OBJ(try_2))
            { var y *ClaireAny  
              var try_3 EID 
              /*g_try(v2:"try_3",loop:false) */
              try_3 = EVAL(x.Args.At(2-1))
              /* ERROR PROTECTION INSERTED (y-loop_1) */
              if ErrorIn(try_3) {loop_1 = try_3
              } else {
              y = ANY(try_3)
              { var s *ClaireObject   = Core.F__at_property1(p,self.Isa)
                if (C_slot.Id() == s.Isa.Id()) { 
                  { var g0219 *ClaireSlot   = ToSlot(s.Id())
                    if (y == CNULL) { 
                      lp = lp.AddFast(p.Id())/*t=any,s=list*/
                      } 
                    /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
                    if (g0219.Range.Contains(y) != CTRUE) { 
                      loop_1 = Core.F_range_is_wrong_slot(g0219,y)
                      } else {
                      loop_1 = Core.F_update_property(p,
                        self,
                        g0219.Index,
                        g0219.Srange,
                        y)
                      } 
                    /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
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
          /* ERROR PROTECTION INSERTED (loop_1-Result) */
          if ErrorIn(loop_1) {Result = loop_1
          break
          } else {
          }
          } 
        } 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = EID{lp.Id(),0}
      }
      } 
    return Result} 
  
// The EID go function for: new_writes @ object (throw: true) 
func E_Language_new_writes_object (self EID,_Zl EID) EID { 
    return F_Language_new_writes_object(ToObject(OBJ(self)),ToList(OBJ(_Zl)) )} 
  
// creation of a new named object
/* {1} The go function for: self_eval(self:Defobj) [status=1] */
func (self *Defobj ) SelfEval () EID { 
    var Result EID 
    { var _Zc *ClaireClass   = self.Arg
      { var _Zo *ClaireObject   = ToObject(CNULL)
        /*g_try(v2:"Result",loop:true) */
        if (_Zc.Open <= 1) { 
          Result = ToException(Core.C_general_error.Make(MakeString("[105] cannot instantiate ~S").Id(),MakeConstantList(_Zc.Id()).Id())).Close()
          } else {
          Result = EID{CFALSE.Id(),0}
          } 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        /*g_try(v2:"Result",loop:true) */
        if (_Zc.IsIn(C_thing) == CTRUE) { 
          var try_1 EID 
          /*g_try(v2:"try_1",loop:true) */
          try_1 = F_new_thing_class(_Zc,self.Ident)
          /* ERROR PROTECTION INSERTED (_Zo-Result) */
          if ErrorIn(try_1) {Result = try_1
          } else {
          _Zo = ToObject(OBJ(try_1))
          Result = EID{_Zo.Id(),0}
          if (_Zo.Isa.IsIn(C_property) == CTRUE) { 
            { var g0221 *ClaireProperty   = ToProperty(_Zo.Id())
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
            _Zc.Instances.AddFast(_Zo.Id())/*t=any,s=void*/
            } 
          Result = self.Ident.Put(_Zo.Id()).ToEID()
          } 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        { var arg_2 *ClaireList  
          _ = arg_2
          var try_3 EID 
          /*g_try(v2:"try_3",loop:false) */
          try_3 = F_Language_new_writes_object(_Zo,self.Args)
          /* ERROR PROTECTION INSERTED (arg_2-Result) */
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
/* {1} The go function for: self_eval(self:Defclass) [status=1] */
func (self *Defclass ) SelfEval () EID { 
    var Result EID 
    if ((C_class.Id() == Core.F_owner_any(self.Ident.Get()).Id()) && 
        ((ToClass(self.Ident.Get()).Open != ClEnv.Final) || 
            (self.Arg.Id() != ToClass(self.Ident.Get()).Superclass.Id()))) { 
      Result = ToException(Core.C_general_error.Make(MakeString("[107] class re-definition is not valid: ~S").Id(),MakeConstantList(self.Id()).Id())).Close()
      } else {
      { var _Zo *ClaireClass  
        var try_1 EID 
        /*g_try(v2:"try_1",loop:false) */
        try_1 = self.Ident.Class_I(self.Arg)
        /* ERROR PROTECTION INSERTED (_Zo-Result) */
        if ErrorIn(try_1) {Result = try_1
        } else {
        _Zo = ToClass(OBJ(try_1))
        /*g_try(v2:"Result",loop:true) */
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
            /*g_try(v2:"loop_2",loop:tuple("Result", EID)) */
            { var v *ClaireAny   = CNULL
              /*g_try(v2:"loop_2",loop:tuple("Result", EID)) */
              if (x.Isa.IsIn(C_Call) == CTRUE) { 
                { var g0222 *Call   = To_Call(x)
                  _ = g0222
                  /*g_try(v2:"loop_2",loop:tuple("Result", EID)) */
                  var try_3 EID 
                  /*g_try(v2:"try_3",loop:tuple("Result", EID)) */
                  try_3 = EVAL(g0222.Args.At(2-1))
                  /* ERROR PROTECTION INSERTED (v-loop_2) */
                  if ErrorIn(try_3) {loop_2 = try_3
                  Result = try_3
                  break
                  } else {
                  v = ANY(try_3)
                  loop_2 = v.ToEID()
                  g0222 = To_Call(g0222.Args.At(1-1))
                  loop_2 = EID{g0222.Id(),0}
                  }
                  /* ERROR PROTECTION INSERTED (loop_2-loop_2) */
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
              /* ERROR PROTECTION INSERTED (loop_2-loop_2) */
              if ErrorIn(loop_2) {Result = loop_2
              break
              } else {
              /*g_try(v2:"loop_2",loop:tuple("Result", EID)) */
              { var rt *ClaireTypeExpression  
                var try_4 EID 
                /*g_try(v2:"try_4",loop:false) */
                try_4 = F_extract_type_any(ANY(Core.F_CALL(C_range,ARGS(x.ToEID()))))
                /* ERROR PROTECTION INSERTED (rt-loop_2) */
                if ErrorIn(try_4) {loop_2 = try_4
                } else {
                rt = ToTypeExpression(OBJ(try_4))
                { var p *ClaireProperty  
                  var try_5 EID 
                  /*g_try(v2:"try_5",loop:false) */
                  try_5 = F_make_a_property_any(ANY(Core.F_CALL(C_mClaire_pname,ARGS(x.ToEID()))))
                  /* ERROR PROTECTION INSERTED (p-loop_2) */
                  if ErrorIn(try_5) {loop_2 = try_5
                  } else {
                  p = ToProperty(OBJ(try_5))
                  /*g_try(v2:"loop_2",loop:tuple("Result", EID)) */
                  var g0223I *ClaireBoolean  
                  var try_6 EID 
                  /*g_try(v2:"try_6",loop:false) */
                  { 
                    var v_and9 *ClaireBoolean  
                    
                    v_and9 = Core.F_known_ask_any(v)
                    if (v_and9 == CFALSE) {try_6 = EID{CFALSE.Id(),0}
                    } else { 
                      var try_7 EID 
                      /*g_try(v2:"try_7",loop:false) */
                      { var arg_8 *ClaireBoolean  
                        _ = arg_8
                        var try_9 EID 
                        /*g_try(v2:"try_9",loop:false) */
                        try_9 = Core.F_BELONG(v,rt.Id())
                        /* ERROR PROTECTION INSERTED (arg_8-try_7) */
                        if ErrorIn(try_9) {try_7 = try_9
                        } else {
                        arg_8 = ToBoolean(OBJ(try_9))
                        try_7 = EID{arg_8.Not.Id(),0}
                        }
                        } 
                      /* ERROR PROTECTION INSERTED (v_and9-try_6) */
                      if ErrorIn(try_7) {try_6 = try_7
                      } else {
                      v_and9 = ToBoolean(OBJ(try_7))
                      if (v_and9 == CFALSE) {try_6 = EID{CFALSE.Id(),0}
                      } else { 
                        try_6 = EID{CTRUE.Id(),0}} 
                      } 
                    }
                    } 
                  /* ERROR PROTECTION INSERTED (g0223I-loop_2) */
                  if ErrorIn(try_6) {loop_2 = try_6
                  } else {
                  g0223I = ToBoolean(OBJ(try_6))
                  if (g0223I == CTRUE) { 
                    loop_2 = ToException(Core.C_general_error.Make(MakeString("[108] default(~S) = ~S does not belong to ~S").Id(),MakeConstantList(x,v,rt.Id()).Id())).Close()
                    } else {
                    loop_2 = EID{CFALSE.Id(),0}
                    } 
                  }
                  /* ERROR PROTECTION INSERTED (loop_2-loop_2) */
                  if ErrorIn(loop_2) {Result = loop_2
                  break
                  } else {
                  /*g_try(v2:"loop_2",loop:tuple("Result", EID)) */
                  { var s2test *ClaireAny  
                    { var sx_some *ClaireAny   = CNULL
                      _ = sx_some
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
                      { var s2 *ClaireSlot   = ToSlot(s2test)
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
                  /* ERROR PROTECTION INSERTED (loop_2-loop_2) */
                  if ErrorIn(loop_2) {Result = loop_2
                  break
                  } else {
                  loop_2 = EID{Core.F_close_slot(_Zo.AddSlot(p,ToType(rt.Id()),F_Language_getDefault_type(ToType(rt.Id()),v))).Id(),0}
                  }}
                  }
                  } 
                }
                } 
              /* ERROR PROTECTION INSERTED (loop_2-loop_2) */
              if ErrorIn(loop_2) {Result = loop_2
              break
              } else {
              }}
              } 
            /* ERROR PROTECTION INSERTED (loop_2-Result) */
            if ErrorIn(loop_2) {Result = loop_2
            break
            } else {
            }
            } 
          } 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        F_close_class(_Zo)
        if (self.Forward_ask == CTRUE) { 
          _Zo.Open = ClEnv.Final
          /*integer->integer*/}  else if (_Zo.Open == ClEnv.Final) { 
          _Zo.Open = self.Arg.Open
          /*integer->integer*/} 
        if (ToType(_Zo.Id()).Included(ToType(C_primitive.Id())) == CTRUE) { 
          _Zo.Open = -1
          /*integer->integer*/} 
        _Zo.Params = self.Params
        /*list->list*/{ 
          var p *ClaireAny  
          _ = p
          var p_support *ClaireList  
          p_support = self.Params
          p_len := p_support.Length()
          for i_it := 0; i_it < p_len; i_it++ { 
            p = p_support.At(i_it)
            ToRelation(p).Open = 0
            /*integer->integer*/} 
          } 
        /*g_try(v2:"Result",loop:true) */
        Result = F_attach_comment_any(_Zo.Id())
        /* ERROR PROTECTION INSERTED (Result-Result) */
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
/* {1} The go function for: getDefault(rt:type,v:any) [status=0] */
func F_Language_getDefault_type (rt *ClaireType ,v *ClaireAny ) *ClaireAny  { 
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
/* {1} The go function for: self_eval(self:Defmethod) [status=1] */
func (self *Defmethod ) SelfEval () EID { 
    var Result EID 
    /*g_try(v2:"Result",loop:true) */
    if (self.Arg.Isa.IsIn(C_Call) != CTRUE) { 
      Result = ToException(Core.C_general_error.Make(MakeString("[110] wrong signature definition ~S").Id(),MakeConstantList(self.Arg.Id()).Id())).Close()
      } else {
      Result = EID{CFALSE.Id(),0}
      } 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    { var p *ClaireProperty  
      var try_1 EID 
      /*g_try(v2:"try_1",loop:false) */
      try_1 = F_make_a_property_any(self.Arg.Selector.Id())
      /* ERROR PROTECTION INSERTED (p-Result) */
      if ErrorIn(try_1) {Result = try_1
      } else {
      p = ToProperty(OBJ(try_1))
      { var l *ClaireList   = self.Arg.Args
        { var lv *ClaireList  
          if ((l.Length() == 1) && 
              (l.At(1-1) == ClEnv.Id())) { 
            { 
              var v_bag_arg *ClaireAny  
              lv= ToType(CEMPTY.Id()).EmptyList()
              { var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
                _CL_obj.Pname = Core.F_symbol_I_string2(MakeString("XfakeParameter"))
                /*symbol->symbol*/_CL_obj.Range = ToType(C_void.Id())
                /*type->type*/v_bag_arg = _CL_obj.Id()
                } 
              lv.AddFast(v_bag_arg)} 
            } else {
            lv = l
            } 
          { var lp *ClaireList  
            _ = lp
            var try_2 EID 
            /*g_try(v2:"try_2",loop:false) */
            try_2 = F_extract_signature_list(lv)
            /* ERROR PROTECTION INSERTED (lp-Result) */
            if ErrorIn(try_2) {Result = try_2
            } else {
            lp = ToList(OBJ(try_2))
            { var lrange *ClaireList  
              var try_3 EID 
              /*g_try(v2:"try_3",loop:false) */
              try_3 = F_extract_range_any(self.SetArg,lv,ToList(C_LDEF.Value))
              /* ERROR PROTECTION INSERTED (lrange-Result) */
              if ErrorIn(try_3) {Result = try_3
              } else {
              lrange = ToList(OBJ(try_3))
              { var lbody *ClaireList  
                var try_4 EID 
                /*g_try(v2:"try_4",loop:false) */
                try_4 = F_extract_status_any(self.Body)
                /* ERROR PROTECTION INSERTED (lbody-Result) */
                if ErrorIn(try_4) {Result = try_4
                } else {
                lbody = ToList(OBJ(try_4))
                { var m *ClaireMethod   = F_add_method_property(p,lp,ToType(lrange.At(1-1)),ToInteger(lbody.At(1-1)).Value,ToFunction(lbody.At(2-1)))
                  /*g_try(v2:"Result",loop:true) */
                  if ((p.Open > 0) && 
                      (p.Open <= 1)) { 
                    { var r *ClaireAny  
                      var try_5 EID 
                      /*g_try(v2:"try_5",loop:false) */
                      { var r_some *ClaireAny   = CNULL
                        _ = r_some
                        /*g_try(v2:"try_5",loop:false) */
                        { 
                          var r *ClaireAny  
                          _ = r
                          try_5= EID{CFALSE.Id(),0}
                          for _,r = range(p.Restrictions.ValuesO()){ 
                            var loop_6 EID 
                            _ = loop_6
                            /*g_try(v2:"loop_6",loop:tuple("try_5", EID)) */
                            if (r != m.Id()) { 
                              var g0224I *ClaireBoolean  
                              var try_7 EID 
                              /*g_try(v2:"try_7",loop:false) */
                              { var arg_8 *ClaireAny  
                                _ = arg_8
                                var try_9 EID 
                                /*g_try(v2:"try_9",loop:false) */
                                try_9 = Core.F_CALL(ToProperty(C__exp.Id()),ARGS(Core.F_CALL(C_domain,ARGS(r.ToEID())),EID{m.Domain.Id(),0}))
                                /* ERROR PROTECTION INSERTED (arg_8-try_7) */
                                if ErrorIn(try_9) {try_7 = try_9
                                } else {
                                arg_8 = ANY(try_9)
                                try_7 = EID{F_boolean_I_any(arg_8).Id(),0}
                                }
                                } 
                              /* ERROR PROTECTION INSERTED (g0224I-loop_6) */
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
                            /* ERROR PROTECTION INSERTED (loop_6-try_5) */
                            if ErrorIn(loop_6) {try_5 = loop_6
                            break
                            } else {
                            }
                            } 
                          } 
                        /* ERROR PROTECTION INSERTED (try_5-try_5) */
                        if !ErrorIn(try_5) {
                        try_5 = r_some.ToEID()
                        }
                        } 
                      /* ERROR PROTECTION INSERTED (r-Result) */
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
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  C_LDEF.Value = ToType(C_any.Id()).EmptyList().Id()
                  /*g_try(v2:"Result",loop:true) */
                  if (lbody.At(3-1) != C_body.Id()) { 
                    Core.F_tformat_string(MakeString("---- jito for ~S\n"),2,MakeConstantList(m.Id()))
                    { 
                      var va_arg1 *ClaireMethod  
                      var va_arg2 *ClaireLambda  
                      va_arg1 = m
                      var try_10 EID 
                      /*g_try(v2:"try_10",loop:false) */
                      { var arg_11 *ClaireLambda  
                        _ = arg_11
                        var try_12 EID 
                        /*g_try(v2:"try_12",loop:false) */
                        try_12 = F_lambda_I_list(lv,lbody.At(3-1))
                        /* ERROR PROTECTION INSERTED (arg_11-try_10) */
                        if ErrorIn(try_12) {try_10 = try_12
                        } else {
                        arg_11 = ToLambda(OBJ(try_12))
                        try_10 = F_Language_jito_any(arg_11.Id())
                        }
                        } 
                      /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                      if ErrorIn(try_10) {Result = try_10
                      } else {
                      va_arg2 = ToLambda(OBJ(try_10))
                      va_arg1.Formula = va_arg2
                      /*lambda->lambda*/Result = EID{va_arg2.Id(),0}
                      }
                      } 
                    } else {
                    Result = EID{CFALSE.Id(),0}
                    } 
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  if (lrange.Length() > 1) { 
                    m.Typing = lrange.At(2-1)
                    /*any->any*/} 
                  m.Inline_ask = self.Inline_ask
                  /*boolean->boolean*//*g_try(v2:"Result",loop:true) */
                  Result = F_attach_comment_any(m.Id())
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  Core.F_close_method(m)
                  /*g_try(v2:"Result",loop:true) */
                  if ((p.Id() == C_close.Id()) && 
                      (m.Range.Included(ToType(Core.F_domain_I_restriction(ToRestriction(m.Id())).Id())) != CTRUE)) { 
                    Result = ToException(Core.C_general_error.Make(MakeString("[184] the close method ~S has a wrong range").Id(),MakeConstantList(m.Id()).Id())).Close()
                    } else {
                    Result = EID{CFALSE.Id(),0}
                    } 
                  /* ERROR PROTECTION INSERTED (Result-Result) */
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
/* {1} The go function for: attach_comment(x:any) [status=1] */
func F_attach_comment_any (x *ClaireAny ) EID { 
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
/* {1} The go function for: iClaire/extract_signature(l:list) [status=1] */
func F_extract_signature_list (l *ClaireList ) EID { 
    var Result EID 
    C_LDEF.Value = ToType(C_any.Id()).EmptyList().Id()
    { var n int  = 0
      _ = n
      { 
        var v_list3 *ClaireList  
        var v *ClaireVariable  
        var v_local3 *ClaireAny  
        v_list3 = l
        Result = EID{CreateList(ToType(CEMPTY.Id()),v_list3.Length()).Id(),0}
        for CLcount := 0; CLcount < v_list3.Length(); CLcount++{ 
          v = To_Variable(v_list3.At(CLcount))
          var try_1 EID 
          /*g_try(v2:"try_1",loop:tuple("Result", EID)) */
          if (v.Isa.IsIn(C_Variable) != CTRUE) { 
            try_1 = ToException(Core.C_general_error.Make(MakeString("[111] wrong typed argument ~S").Id(),MakeConstantList(v.Id()).Id())).Close()
            } else {
            { var p *ClaireAny  
              var try_2 EID 
              /*g_try(v2:"try_2",loop:false) */
              try_2 = F_extract_pattern_any(v.Range.Id(),MakeConstantList(MakeInteger(n).Id()))
              /* ERROR PROTECTION INSERTED (p-try_1) */
              if ErrorIn(try_2) {try_1 = try_2
              } else {
              p = ANY(try_2)
              n = (n+1)
              /*g_try(v2:"try_1",loop:tuple("Result", EID)) */
              if (p == CNULL) { 
                try_1 = ToException(Core.C_general_error.Make(MakeString("[111] wrong typed argument ~S (~S)").Id(),MakeConstantList(v.Id(),v.Range.Id()).Id())).Close()
                } else {
                try_1 = EID{CFALSE.Id(),0}
                } 
              /* ERROR PROTECTION INSERTED (try_1-try_1) */
              if ErrorIn(try_1) {Result = try_1
              break
              } else {
              v.Range = F_type_I_any(p)
              /*type->type*/try_1 = p.ToEID()
              }
              }
              } 
            } 
          /* ERROR PROTECTION INSERTED (v_local3-Result) */
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
/* {1} The go function for: iClaire/extract_pattern(x:any,path:list) [status=1] */
func F_extract_pattern_any (x *ClaireAny ,path *ClaireList ) EID { 
    var Result EID 
    if (C_class.Id() == x.Isa.Id()) { 
      { var g0225 *ClaireClass   = ToClass(x)
        _ = g0225
        Result = EID{g0225.Id(),0}
        } 
      }  else if (C_set.Id() == x.Isa.Id()) { 
      { var g0226 *ClaireSet   = ToSet(x)
        { var z *ClaireAny  
          var try_1 EID 
          /*g_try(v2:"try_1",loop:false) */
          if ANY(Core.F_CALL(C_length,ARGS(EID{g0226.Id(),0}))).IsInt(1) { 
            { var arg_2 *ClaireAny  
              _ = arg_2
              var try_3 EID 
              /*g_try(v2:"try_3",loop:false) */
              try_3 = Core.F_CALL(C_nth,ARGS(EID{g0226.Id(),0},EID{C__INT,IVAL(1)}))
              /* ERROR PROTECTION INSERTED (arg_2-try_1) */
              if ErrorIn(try_3) {try_1 = try_3
              } else {
              arg_2 = ANY(try_3)
              try_1 = F_extract_pattern_any(arg_2,CNIL)
              }
              } 
            } else {
            try_1 = EID{CFALSE.Id(),0}
            } 
          /* ERROR PROTECTION INSERTED (z-Result) */
          if ErrorIn(try_1) {Result = try_1
          } else {
          z = ANY(try_1)
          if (z.Isa.IsIn(C_Reference) == CTRUE) { 
            { var g0227 *ClaireReference   = To_Reference(z)
              _ = g0227
              { var w *ClaireReference   = To_Reference(g0227.Copy().Id())
                w.Arg = CTRUE
                /*boolean->boolean*/Result = EID{w.Id(),0}
                } 
              } 
            } else {
            Result = EID{g0226.Id(),0}
            } 
          }
          } 
        } 
      }  else if (x.Isa.IsIn(C_Tuple) == CTRUE) { 
      { var g0229 *Tuple   = To_Tuple(x)
        _ = g0229
        { var ltp *ClaireList  
          var try_4 EID 
          /*g_try(v2:"try_4",loop:false) */
          { 
            var v_list5 *ClaireList  
            var z *ClaireAny  
            var v_local5 *ClaireAny  
            v_list5 = g0229.Args
            try_4 = EID{CreateList(ToType(CEMPTY.Id()),v_list5.Length()).Id(),0}
            for CLcount := 0; CLcount < v_list5.Length(); CLcount++{ 
              z = v_list5.At(CLcount)
              var try_5 EID 
              /*g_try(v2:"try_5",loop:tuple("try_4", EID)) */
              try_5 = F_extract_pattern_any(z,path)
              /* ERROR PROTECTION INSERTED (v_local5-try_4) */
              if ErrorIn(try_5) {try_4 = try_5
              break
              } else {
              v_local5 = ANY(try_5)
              ToList(OBJ(try_4)).PutAt(CLcount,v_local5)
              } 
            }
            } 
          /* ERROR PROTECTION INSERTED (ltp-Result) */
          if ErrorIn(try_4) {Result = try_4
          } else {
          ltp = ToList(OBJ(try_4))
          var g0236I *ClaireBoolean  
          { var arg_6 *ClaireAny  
            _ = arg_6
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
      { var g0230 *Core.GlobalVariable   = Core.ToGlobalVariable(x)
        _ = g0230
        Result = F_extract_pattern_any(g0230.Value,path)
        } 
      }  else if (x.Isa.IsIn(C_Call) == CTRUE) { 
      { var g0231 *Call   = To_Call(x)
        { var p *ClaireProperty   = g0231.Selector
          if (p.Id() == Core.C_U.Id()) { 
            { var x1 *ClaireAny  
              var try_7 EID 
              /*g_try(v2:"try_7",loop:false) */
              try_7 = F_extract_pattern_any(g0231.Args.At(1-1),CNIL)
              /* ERROR PROTECTION INSERTED (x1-Result) */
              if ErrorIn(try_7) {Result = try_7
              } else {
              x1 = ANY(try_7)
              { var x2 *ClaireAny  
                var try_8 EID 
                /*g_try(v2:"try_8",loop:false) */
                try_8 = F_extract_pattern_any(g0231.Args.At(2-1),CNIL)
                /* ERROR PROTECTION INSERTED (x2-Result) */
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
              _ = arg_9
              var try_11 EID 
              /*g_try(v2:"try_11",loop:false) */
              try_11 = F_extract_pattern_any(g0231.Args.At(1-1),CNIL)
              /* ERROR PROTECTION INSERTED (arg_9-Result) */
              if ErrorIn(try_11) {Result = try_11
              } else {
              arg_9 = ANY(try_11)
              { var arg_10 *ClaireAny  
                _ = arg_10
                var try_12 EID 
                /*g_try(v2:"try_12",loop:false) */
                try_12 = F_extract_pattern_any(g0231.Args.At(2-1),CNIL)
                /* ERROR PROTECTION INSERTED (arg_10-Result) */
                if ErrorIn(try_12) {Result = try_12
                } else {
                arg_10 = ANY(try_12)
                Result = Core.F_CALL(ToProperty(C__exp.Id()),ARGS(arg_9.ToEID(),arg_10.ToEID()))
                }
                } 
              }
              } 
            }  else if (p.Id() == C__dot_dot.Id()) { 
            { var v1 *ClaireAny   = F_extract_item_any(g0231.Args.At(1-1),CNIL.Id())
              { var v2 *ClaireAny   = F_extract_item_any(g0231.Args.At(2-1),CNIL.Id())
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
              /*g_try(v2:"try_13",loop:false) */
              try_13 = F_extract_pattern_any(g0231.Args.At(1-1),path)
              /* ERROR PROTECTION INSERTED (z-Result) */
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
      { var g0232 *ClaireType   = ToType(x)
        _ = g0232
        Result = EID{g0232.Id(),0}
        } 
      }  else if (x.Isa.IsIn(C_unbound_symbol) == CTRUE) { 
      { var g0233 *ClaireUnboundSymbol   = ToUnboundSymbol(x)
        _ = g0233
        { var s *ClaireSymbol  
          var try_14 EID 
          /*g_try(v2:"try_14",loop:false) */
          try_14 = F_extract_symbol_any(g0233.Id())
          /* ERROR PROTECTION INSERTED (s-Result) */
          if ErrorIn(try_14) {Result = try_14
          } else {
          s = ToSymbol(OBJ(try_14))
          { var v *ClaireAny  
            var try_15 EID 
            /*g_try(v2:"try_15",loop:false) */
            { var z_some *ClaireAny   = CNULL
              _ = z_some
              /*g_try(v2:"try_15",loop:false) */
              { 
                var z *ClaireAny  
                _ = z
                try_15= EID{CFALSE.Id(),0}
                var z_support *ClaireList  
                var try_16 EID 
                /*g_try(v2:"try_16",loop:false) */
                try_16 = Core.F_enumerate_any(C_LDEF.Value)
                /* ERROR PROTECTION INSERTED (z_support-try_15) */
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
              /* ERROR PROTECTION INSERTED (try_15-try_15) */
              if !ErrorIn(try_15) {
              try_15 = z_some.ToEID()
              }
              } 
            /* ERROR PROTECTION INSERTED (v-Result) */
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
                  _ = y
                  var try_17 EID 
                  /*g_try(v2:"try_17",loop:false) */
                  { var arg_18 *ClaireList  
                    _ = arg_18
                    var try_19 EID 
                    /*g_try(v2:"try_19",loop:false) */
                    try_19 = path.Cdr()
                    /* ERROR PROTECTION INSERTED (arg_18-try_17) */
                    if ErrorIn(try_19) {try_17 = try_19
                    } else {
                    arg_18 = ToList(OBJ(try_19))
                    try_17 = EID{Core.F_Reference_I_list(arg_18,ToInteger(path.At(1-1)).Value).Id(),0}
                    }
                    } 
                  /* ERROR PROTECTION INSERTED (y-Result) */
                  if ErrorIn(try_17) {Result = try_17
                  } else {
                  y = To_Reference(OBJ(try_17))
                  { var v *ClaireVariable  
                    _ = v
                    { var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
                      _CL_obj.Pname = s
                      /*symbol->symbol*/_CL_obj.Range = ToType(y.Id())
                      /*type->type*/v = _CL_obj
                      } 
                    
                    /*g_try(v2:"Result",loop:true) */
                    var v_gassign20 *ClaireAny  
                    var try_21 EID 
                    /*g_try(v2:"try_21",loop:false) */
                    try_21 = Core.F_CALL(ToProperty(C_add.Id()),ARGS(EID{C_LDEF.Value,0},EID{v.Id(),0}))
                    /* ERROR PROTECTION INSERTED (v_gassign20-Result) */
                    if ErrorIn(try_21) {Result = try_21
                    } else {
                    v_gassign20 = ANY(try_21)
                    C_LDEF.Value = v_gassign20
                    Result = v_gassign20.ToEID()
                    } 
                  /* ERROR PROTECTION INSERTED (Result-Result) */
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
/* {0} The go function for: iClaire/extract_type(x:any) [status=1] */
func F_extract_type_any (x *ClaireAny ) EID { 
  var Result EID 
  C_LDEF.Value = ToType(C_any.Id()).EmptyList().Id()
  { var r *ClaireAny  
    var try_1 EID 
    /*g_try(v2:"try_1",loop:false) */
    try_1 = F_extract_pattern_any(x,CNIL)
    /* ERROR PROTECTION INSERTED (r-Result) */
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
/* {0} The go function for: extract_item(x:any,y:any) [status=0] */
func F_extract_item_any (x *ClaireAny ,y *ClaireAny ) *ClaireAny  { 
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
/* {0} The go function for: extract_pattern_nth(l:list,path:list) [status=1] */
func F_extract_pattern_nth_list (l *ClaireList ,path *ClaireList ) EID { 
  var Result EID 
  { var m int  = l.Length()
    { var x *ClaireAny   = l.At(1-1)
      if (m == 1) { 
        { var y *ClaireAny  
          var try_1 EID 
          /*g_try(v2:"try_1",loop:false) */
          try_1 = F_extract_pattern_any(l.At(1-1),CNIL)
          /* ERROR PROTECTION INSERTED (y-Result) */
          if ErrorIn(try_1) {Result = try_1
          } else {
          y = ANY(try_1)
          if (y == CNULL) { 
            Result = EID{CNULL,0}
            } else {
            { var _CL_obj *ClaireParam   = To_Param(new(ClaireParam).Is(C_Param))
              _CL_obj.Arg = C_array
              /*class->class*/_CL_obj.Params = MakeConstantList(C_of.Id())
              /*list->list*/_CL_obj.Args = MakeConstantList(MakeConstantSet(y).Id())
              /*list->list*/Result = EID{_CL_obj.Id(),0}
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
            /*g_try(v2:"try_2",loop:false) */
            try_2 = F_extract_pattern_any(l.At(2-1),CNIL)
            /* ERROR PROTECTION INSERTED (y-Result) */
            if ErrorIn(try_2) {Result = try_2
            } else {
            y = ANY(try_2)
            { 
              h_index := ClEnv.Index
              h_base := ClEnv.Base
              if (y != CNULL) { 
                Result = Core.F_CALL(C_nth,ARGS(l.At(1-1).ToEID(),y.ToEID()))
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
        { var l1 *ClaireAny   = l.At(2-1)
          { var l2 *ClaireList   = ToList(OBJ(Core.F_CALL(C_args,ARGS(l.At(3-1).ToEID()))))
            _ = l2
            { var l3 *ClaireList   = ToType(C_any.Id()).EmptyList()
              /*g_try(v2:"Result",loop:true) */
              { var n int  = 1
                { var g0238 int  = INT(Core.F_CALL(C_length,ARGS(l1.ToEID())))
                  _ = g0238
                  Result= EID{CFALSE.Id(),0}
                  for (n <= g0238) { 
                    /* While stat, v:"Result" loop:true */
                    var loop_3 EID 
                    _ = loop_3
                    { 
                    /*g_try(v2:"loop_3",loop:tuple("Result", EID)) */
                    { var y *ClaireAny   = l2.At(n-1)
                      _ = y
                      var try_4 EID 
                      /*g_try(v2:"try_4",loop:false) */
                      { var arg_5 *ClaireAny  
                        _ = arg_5
                        var try_6 EID 
                        /*g_try(v2:"try_6",loop:false) */
                        if (y.Isa.IsIn(C_Set) == CTRUE) { 
                          { var g0239 *Set   = To_Set(y)
                            { var v *ClaireAny  
                              var try_7 EID 
                              /*g_try(v2:"try_7",loop:false) */
                              { var arg_8 *ClaireList  
                                _ = arg_8
                                var try_9 EID 
                                /*g_try(v2:"try_9",loop:false) */
                                { var arg_10 *ClaireAny  
                                  _ = arg_10
                                  var try_11 EID 
                                  /*g_try(v2:"try_11",loop:false) */
                                  try_11 = Core.F_CALL(C_nth,ARGS(l1.ToEID(),EID{C__INT,IVAL(n)}))
                                  /* ERROR PROTECTION INSERTED (arg_10-try_9) */
                                  if ErrorIn(try_11) {try_9 = try_11
                                  } else {
                                  arg_10 = ANY(try_11)
                                  try_9 = EID{path.Copy().AddFast(arg_10).Id(),0}/*t=any,s=EID*/
                                  }
                                  } 
                                /* ERROR PROTECTION INSERTED (arg_8-try_7) */
                                if ErrorIn(try_9) {try_7 = try_9
                                } else {
                                arg_8 = ToList(OBJ(try_9))
                                try_7 = F_extract_pattern_any(g0239.Args.At(1-1),arg_8)
                                }
                                } 
                              /* ERROR PROTECTION INSERTED (v-try_6) */
                              if ErrorIn(try_7) {try_6 = try_7
                              } else {
                              v = ANY(try_7)
                              if (v == C_void.Id()) { 
                                try_6 = EID{C_any.Id(),0}
                                }  else if (v.Isa.IsIn(C_Reference) == CTRUE) { 
                                { var g0241 *ClaireReference   = To_Reference(v)
                                  _ = g0241
                                  { var z *ClaireReference   = To_Reference(g0241.Copy().Id())
                                    z.Arg = CTRUE
                                    /*boolean->boolean*/try_6 = EID{z.Id(),0}
                                    } 
                                  } 
                                } else {
                                { 
                                  var v_bag_arg *ClaireAny  
                                  try_6= EID{ToType(CEMPTY.Id()).EmptySet().Id(),0}
                                  var try_12 EID 
                                  /*g_try(v2:"try_12",loop:false) */
                                  if (v != CNULL) { 
                                    try_12 = v.ToEID()
                                    } else {
                                    try_12 = EVAL(g0239.Args.At(1-1))
                                    } 
                                  /* ERROR PROTECTION INSERTED (v_bag_arg-try_6) */
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
                            _ = arg_13
                            var try_14 EID 
                            /*g_try(v2:"try_14",loop:false) */
                            if (path.Length() != 0) { 
                              { var arg_15 *ClaireAny  
                                _ = arg_15
                                var try_16 EID 
                                /*g_try(v2:"try_16",loop:false) */
                                try_16 = Core.F_CALL(C_nth,ARGS(l1.ToEID(),EID{C__INT,IVAL(n)}))
                                /* ERROR PROTECTION INSERTED (arg_15-try_14) */
                                if ErrorIn(try_16) {try_14 = try_16
                                } else {
                                arg_15 = ANY(try_16)
                                try_14 = EID{path.AddFast(arg_15).Id(),0}/*t=any,s=EID*/
                                }
                                } 
                              } else {
                              try_14 = EID{CFALSE.Id(),0}
                              } 
                            /* ERROR PROTECTION INSERTED (arg_13-try_6) */
                            if ErrorIn(try_14) {try_6 = try_14
                            } else {
                            arg_13 = ANY(try_14)
                            try_6 = F_extract_pattern_any(y,ToList(arg_13))
                            }
                            } 
                          } 
                        /* ERROR PROTECTION INSERTED (arg_5-try_4) */
                        if ErrorIn(try_6) {try_4 = try_6
                        } else {
                        arg_5 = ANY(try_6)
                        try_4 = EID{l3.AddFast(arg_5).Id(),0}/*t=any,s=EID*/
                        }
                        } 
                      /* ERROR PROTECTION INSERTED (l3-loop_3) */
                      if ErrorIn(try_4) {loop_3 = try_4
                      } else {
                      l3 = ToList(OBJ(try_4))
                      loop_3 = EID{l3.Id(),0}
                      }
                      } 
                    /* ERROR PROTECTION INSERTED (loop_3-loop_3) */
                    if ErrorIn(loop_3) {Result = loop_3
                    break
                    } else {
                    n = (n+1)
                    }
                    /* try?:false, v2:"v_while9" loop will be:tuple("Result", EID) */
                    } 
                  }
                  } 
                } 
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              if (l3.Memq(CNULL) == CTRUE) { 
                Result = EID{CNULL,0}
                } else {
                { var _CL_obj *ClaireParam   = To_Param(new(ClaireParam).Is(C_Param))
                  _CL_obj.Arg = ToClass(x)
                  /*class->class*/_CL_obj.Params = ToList(l1)
                  /*list->list*/_CL_obj.Args = l3
                  /*list->list*/Result = EID{_CL_obj.Id(),0}
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
/* {0} The go function for: iClaire/extract_class_call(self:class,l:list) [status=1] */
func F_extract_class_call_class (self *ClaireClass ,l *ClaireList ) EID { 
  var Result EID 
  var g0251I *ClaireBoolean  
  var try_1 EID 
  /*g_try(v2:"try_1",loop:false) */
  { 
    var v_and1 *ClaireBoolean  
    
    v_and1 = MakeBoolean((self.Id() == C_list.Id()) || (self.Id() == C_set.Id()) || (self.Id() == C_subtype.Id()))
    if (v_and1 == CFALSE) {try_1 = EID{CFALSE.Id(),0}
    } else { 
      v_and1 = Equal(MakeInteger(l.Length()).Id(),MakeInteger(1).Id())
      if (v_and1 == CFALSE) {try_1 = EID{CFALSE.Id(),0}
      } else { 
        var try_2 EID 
        /*g_try(v2:"try_2",loop:false) */
        { var y *ClaireAny   = l.At(1-1)
          { var z *ClaireAny  
            _ = z
            var try_3 EID 
            /*g_try(v2:"try_3",loop:false) */
            try_3 = F_extract_pattern_any(y,CNIL)
            /* ERROR PROTECTION INSERTED (z-try_2) */
            if ErrorIn(try_3) {try_2 = try_3
            } else {
            z = ANY(try_3)
            if (y.Isa.IsIn(Core.C_global_variable) == CTRUE) { 
              { var g0244 *Core.GlobalVariable   = Core.ToGlobalVariable(y)
                _ = g0244
                g0244 = Core.ToGlobalVariable(OBJ(Core.F_CALL(C_value,ARGS(l.At(1-1).ToEID()))))
                y = g0244.Id()
                } 
              } 
            { 
              /* Or stat: v="try_2", loop=false */
              var v_or6 *ClaireBoolean  
              
              /* Or stat: try inherit? @ class(owner @ any(z),type) with try:false, v="try_2", loop=false */
              v_or6 = z.Isa.IsIn(C_type)
              if (v_or6 == CTRUE) {try_2 = EID{CTRUE.Id(),0}
              } else { 
                /* Or stat: try = @ any(self,subtype) with try:false, v="try_2", loop=false */
                v_or6 = Equal(self.Id(),C_subtype.Id())
                if (v_or6 == CTRUE) {try_2 = EID{CTRUE.Id(),0}
                } else { 
                  /* Or stat: try (if (inherit? @ class(owner @ any(y),Call)) let g0245:Call := (<y:Call>) in ((!= @ any(selector @ Call(g0245),=)) | (!= @ any(length @ list(args @ Call(g0245)),2))) else if (inherit? @ class(owner @ any(y),Tuple)) true else false) with try:false, v="try_2", loop=false */
                  if (y.Isa.IsIn(C_Call) == CTRUE) { 
                    { var g0245 *Call   = To_Call(y)
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
        /* ERROR PROTECTION INSERTED (v_and1-try_1) */
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
  /* ERROR PROTECTION INSERTED (g0251I-Result) */
  if ErrorIn(try_1) {Result = try_1
  } else {
  g0251I = ToBoolean(OBJ(try_1))
  if (g0251I == CTRUE) { 
    { var _CL_obj *Call   = To_Call(new(Call).Is(C_Call))
      _CL_obj.Selector = C_nth
      /*property->property*/_CL_obj.Args = F_cons_any(self.Id(),l)
      /*list->list*/Result = EID{_CL_obj.Id(),0}
      } 
    }  else if (self.Id() == C_lambda.Id()) { 
    if ((l.Length() == 2) && 
        ((l.At(1-1).Isa.IsIn(C_Do) == CTRUE) || 
            (l.At(1-1).Isa.IsIn(C_Variable) == CTRUE))) { 
      { var lv *ClaireList  
        if (l.At(1-1).Isa.IsIn(C_Do) == CTRUE) { 
          { var v_in *ClaireList   = ToList(OBJ(Core.F_CALL(C_args,ARGS(l.At(1-1).ToEID()))))
            { var v_out *ClaireList   = v_in.Empty()
              { 
                var v *ClaireAny  
                _ = v
                var v_support *ClaireList  
                v_support = v_in
                v_len := v_support.Length()
                for i_it := 0; i_it < v_len; i_it++ { 
                  v = v_support.At(i_it)
                  if (v.Isa.IsIn(C_Variable) == CTRUE) { 
                    v_out.AddFast(v)/*t=any,s=void*/
                    } 
                  } 
                } 
              lv = v_out
              } 
            } 
          } else {
          lv = MakeConstantList(l.At(1-1))
          } 
        /*g_try(v2:"Result",loop:true) */
        Result = F_extract_signature_list(lv)
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        Result = F_lambda_I_list(lv,l.At(2-1))
        }
        } 
      } else {
      Result = ToException(Core.C_general_error.Make(MakeString("[113] Wrong lambda definition lambda[~S]").Id(),MakeConstantList(l.Id()).Id())).Close()
      } 
    } else {
    { var l1 *ClaireList   = ToType(C_any.Id()).EmptyList()
      _ = l1
      { var l2 *ClaireList   = ToType(C_any.Id()).EmptyList()
        _ = l2
        { var m int  = l.Length()
          _ = m
          /*g_try(v2:"Result",loop:true) */
          { var n int  = 1
            { var g0247 int  = m
              _ = g0247
              Result= EID{CFALSE.Id(),0}
              for (n <= g0247) { 
                /* While stat, v:"Result" loop:true */
                var loop_4 EID 
                _ = loop_4
                { 
                /*g_try(v2:"loop_4",loop:tuple("Result", EID)) */
                { var y *ClaireAny   = l.At(n-1)
                  { var p *ClaireAny   = CNULL
                    _ = p
                    { var v *ClaireAny   = CNULL
                      _ = v
                      /*g_try(v2:"loop_4",loop:tuple("Result", EID)) */
                      if (y.Isa.IsIn(C_Call) == CTRUE) { 
                        { var g0248 *Call   = To_Call(y)
                          /*g_try(v2:"loop_4",loop:tuple("Result", EID)) */
                          if ((g0248.Selector.Id() != C__equal.Id()) || 
                              (g0248.Args.Length() != 2)) { 
                            loop_4 = ToException(Core.C_general_error.Make(MakeString("[114] Wrong parametrization ~S").Id(),MakeConstantList(g0248.Id()).Id())).Close()
                            } else {
                            loop_4 = EID{CFALSE.Id(),0}
                            } 
                          /* ERROR PROTECTION INSERTED (loop_4-loop_4) */
                          if ErrorIn(loop_4) {Result = loop_4
                          break
                          } else {
                          var try_5 EID 
                          /*g_try(v2:"try_5",loop:tuple("Result", EID)) */
                          try_5 = F_make_a_property_any(g0248.Args.At(1-1))
                          /* ERROR PROTECTION INSERTED (p-loop_4) */
                          if ErrorIn(try_5) {loop_4 = try_5
                          Result = try_5
                          break
                          } else {
                          p = ANY(try_5)
                          loop_4 = p.ToEID()
                          { var _CL_obj *Set   = To_Set(new(Set).Is(C_Set))
                            _CL_obj.Args = MakeConstantList(g0248.Args.At(2-1))
                            /*list->list*/v = _CL_obj.Id()
                            } 
                          loop_4 = v.ToEID()
                          }}
                          } 
                        }  else if (y.Isa.IsIn(C_Vardef) == CTRUE) { 
                        { var g0249 *Vardef   = To_Vardef(y)
                          _ = g0249
                          var try_6 EID 
                          /*g_try(v2:"try_6",loop:tuple("Result", EID)) */
                          try_6 = F_make_a_property_any(g0249.Pname.Id())
                          /* ERROR PROTECTION INSERTED (p-loop_4) */
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
                        /*g_try(v2:"try_7",loop:tuple("Result", EID)) */
                        try_7 = F_make_a_property_any(y)
                        /* ERROR PROTECTION INSERTED (p-loop_4) */
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
                      /* ERROR PROTECTION INSERTED (loop_4-loop_4) */
                      if ErrorIn(loop_4) {Result = loop_4
                      break
                      } else {
                      l1 = l1.AddFast(p)/*t=any,s=list*/
                      l2 = l2.AddFast(v)/*t=any,s=list*/
                      loop_4 = EID{l2.Id(),0}
                      }
                      } 
                    } 
                  } 
                /* ERROR PROTECTION INSERTED (loop_4-loop_4) */
                if ErrorIn(loop_4) {Result = loop_4
                break
                } else {
                n = (n+1)
                }
                /* try?:false, v2:"v_while7" loop will be:tuple("Result", EID) */
                } 
              }
              } 
            } 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          { var _CL_obj *Call   = To_Call(new(Call).Is(C_Call))
            _CL_obj.Selector = C_nth
            /*property->property*/{ 
              var va_arg1 *Call  
              var va_arg2 *ClaireList  
              va_arg1 = _CL_obj
              { var arg_8 *ClaireList  
                _ = arg_8
                { 
                  var v_bag_arg *ClaireAny  
                  arg_8= ToType(CEMPTY.Id()).EmptyList()
                  arg_8.AddFast(l1.Id())
                  { var _CL_obj *List   = To_List(new(List).Is(C_List))
                    _CL_obj.Args = l2
                    /*list->list*/v_bag_arg = _CL_obj.Id()
                    } 
                  arg_8.AddFast(v_bag_arg)} 
                va_arg2 = F_cons_any(self.Id(),arg_8)
                } 
              va_arg1.Args = va_arg2
              /*list->list*/} 
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
/* {0} The go function for: iClaire/extract_range(x:any,lvar:list,ldef:list) [status=1] */
func F_extract_range_any (x *ClaireAny ,lvar *ClaireList ,ldef *ClaireList ) EID { 
  var Result EID 
  var g0254I *ClaireBoolean  
  { var arg_1 *ClaireBoolean  
    _ = arg_1
    if (x.Isa.IsIn(C_Call) == CTRUE) { 
      { var g0252 *Call   = To_Call(x)
        arg_1 = MakeBoolean((g0252.Selector.Id() == C_nth.Id()) && (g0252.Args.At(1-1) == C_type.Id()))
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
      /*g_try(v2:"try_2",loop:false) */
      try_2 = F_extract_type_any(x)
      /* ERROR PROTECTION INSERTED (v_bag_arg-Result) */
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
        { var r *ClaireReference   = To_Reference(OBJ(Core.F_CALL(C_range,ARGS(v.ToEID()))))
          { var path *ClaireList   = r.Args
            _ = path
            { var n int  = path.Length()
              _ = n
              { var y *ClaireAny   = lvar.At((r.Index+1)-1)
                _ = y
                { var i int  = 1
                  _ = i
                  { var g0253 int  = n
                    _ = g0253
                    for (i <= g0253) { 
                      /* While stat, v:"Result" loop:tuple("Result", void) */
                      { var _CL_obj *Call   = To_Call(new(Call).Is(C_Call))
                        _CL_obj.Selector = ToProperty(Core.C__at.Id())
                        /*property->property*/_CL_obj.Args = MakeConstantList(y,path.At(i-1))
                        /*list->list*/y = _CL_obj.Id()
                        } 
                      i = (i+1)
                      /* try?:false, v2:"v_while10" loop will be:tuple("Result", void) */
                      } 
                    } 
                  } 
                { var arg_3 *Call  
                  _ = arg_3
                  { var _CL_obj *Call   = To_Call(new(Call).Is(C_Call))
                    _CL_obj.Selector = Core.C_member
                    /*property->property*/_CL_obj.Args = MakeConstantList(y)
                    /*list->list*/arg_3 = _CL_obj
                    } 
                  x = F_substitution_any(x,To_Variable(v),arg_3.Id())
                  } 
                } 
              } 
            } 
          } 
        } 
      } 
    { var lv2 *ClaireList   = ToType(C_any.Id()).EmptyList()
      _ = lv2
      { 
        var v *ClaireAny  
        _ = v
        var v_support *ClaireList  
        v_support = lvar
        v_len := v_support.Length()
        for i_it := 0; i_it < v_len; i_it++ { 
          v = v_support.At(i_it)
          { var v2 *ClaireVariable  
            _ = v2
            { var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
              _CL_obj.Pname = ToSymbol(OBJ(Core.F_CALL(C_mClaire_pname,ARGS(v.ToEID()))))
              /*symbol->symbol*/_CL_obj.Range = ToType(C_type.Id())
              /*type->type*/v2 = _CL_obj
              } 
            lv2 = lv2.AddFast(v2.Id())/*t=any,s=list*/
            x = F_substitution_any(x,To_Variable(v),v2.Id())
            } 
          } 
        } 
      { var lb *ClaireLambda  
        var try_4 EID 
        /*g_try(v2:"try_4",loop:false) */
        try_4 = F_lambda_I_list(lv2,ToList(OBJ(Core.F_CALL(C_args,ARGS(x.ToEID())))).At(2-1))
        /* ERROR PROTECTION INSERTED (lb-Result) */
        if ErrorIn(try_4) {Result = try_4
        } else {
        lb = ToLambda(OBJ(try_4))
        { var ur *ClaireAny   = CNULL
          
          /*g_try(v2:"Result",loop:true) */
          { 
            h_index := ClEnv.Index
            h_base := ClEnv.Base
            var try_5 EID 
            /*g_try(v2:"try_5",loop:false) */
            { var arg_6 *ClaireList  
              _ = arg_6
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
            /* ERROR PROTECTION INSERTED (ur-Result) */
            if ErrorIn(try_5) {Result = try_5
            } else {
            ur = ANY(try_5)
            Result = ur.ToEID()
            }
            if ErrorIn(Result){ 
              ClEnv.Index = h_index
              ClEnv.Base = h_base
              /*g_try(v2:"Result",loop:true) */
              PRINC("The type expression ")
              /*g_try(v2:"Result",loop:true) */
              Result = Core.F_CALL(C_print,ARGS(x.ToEID()))
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              PRINC(" is not valid ... \n")
              Result = EVOID
              }
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              /*g_try(v2:"Result",loop:true) */
              PRINC("context: lambda = ")
              /*g_try(v2:"Result",loop:true) */
              Result = Core.F_print_any(lb.Id())
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              PRINC(", lvars = ")
              /*g_try(v2:"Result",loop:true) */
              { var arg_7 *ClaireList  
                _ = arg_7
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
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              PRINC("\n")
              Result = EVOID
              }}
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              Result = ClEnv.Exception_I.Close()
              }}
              } 
            } 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          /*g_try(v2:"Result",loop:true) */
          if (ur.Isa.IsIn(C_type) != CTRUE) { 
            Result = ToException(Core.C_general_error.Make(MakeString("[115] the (resulting) range ~S is not a type").Id(),MakeConstantList(ur).Id())).Close()
            } else {
            Result = EID{CFALSE.Id(),0}
            } 
          /* ERROR PROTECTION INSERTED (Result-Result) */
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
/* {0} The go function for: bit_vector(l:listargs) [status=1] */
func F_bit_vector_listargs2 (l *ClaireList ) EID { 
  var Result EID 
  { var d int  = 0
    _ = d
    /*g_try(v2:"Result",loop:true) */
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
        /*g_try(v2:"try_2",loop:tuple("Result", EID)) */
        { var arg_3 int 
          _ = arg_3
          var try_4 EID 
          /*g_try(v2:"try_4",loop:false) */
          try_4 = F__exp2_integer(ToInteger(x).Value)
          /* ERROR PROTECTION INSERTED (arg_3-try_2) */
          if ErrorIn(try_4) {try_2 = try_4
          } else {
          arg_3 = INT(try_4)
          try_2 = EID{C__INT,IVAL((d+arg_3))}
          }
          } 
        /* ERROR PROTECTION INSERTED (d-Result) */
        if ErrorIn(try_2) {Result = try_2
        break
        } else {
        d = INT(try_2)
        loop_1 = EID{C__INT,IVAL(d)}
        }
        } 
      } 
    /* ERROR PROTECTION INSERTED (Result-Result) */
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
/* {0} The go function for: iClaire/extract_status(x:any) [status=1] */
func F_extract_status_any (x *ClaireAny ) EID { 
  var Result EID 
  { var s int  = -1
    _ = s
    { var f *ClaireAny  
      var g0260I *ClaireBoolean  
      if (x.Isa.IsIn(C_Call) == CTRUE) { 
        { var g0255 *Call   = To_Call(x)
          _ = g0255
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
        { var g0256 *And   = To_And(x)
          _ = g0256
          { var y *ClaireAny   = g0256.Args.At(1-1)
            var g0261I *ClaireBoolean  
            if (y.Isa.IsIn(C_Call) == CTRUE) { 
              { var g0257 *Call   = To_Call(y)
                _ = g0257
                g0261I = Equal(g0257.Selector.Id(),C_function_I.Id())
                } 
              } else {
              g0261I = CFALSE
              } 
            if (g0261I == CTRUE) { 
              f = y
              g0256 = To_And(g0256.Args.At(2-1))
              } 
            } 
          x = g0256.Id()
          } 
        }  else if (x.Isa.IsIn(C_Call) == CTRUE) { 
        { var g0258 *Call   = To_Call(x)
          _ = g0258
          if (g0258.Selector.Id() == C_function_I.Id()) { 
            g0258 = To_Call(C_body.Id())
            } 
          x = g0258.Id()
          } 
        } else {
        
        } 
      /*g_try(v2:"Result",loop:true) */
      if (f != CNULL) { 
        x = C_body.Id()
        if (ToList(OBJ(Core.F_CALL(C_args,ARGS(f.ToEID())))).Length() > 1) { 
          s = 1
          } else {
          s = 0
          } 
        var try_1 EID 
        /*g_try(v2:"try_1",loop:false) */
        { var arg_2 *ClaireString  
          _ = arg_2
          var try_3 EID 
          /*g_try(v2:"try_3",loop:false) */
          { var arg_4 *ClaireSymbol  
            _ = arg_4
            var try_5 EID 
            /*g_try(v2:"try_5",loop:false) */
            try_5 = F_extract_symbol_any(ToList(OBJ(Core.F_CALL(C_args,ARGS(f.ToEID())))).At(1-1))
            /* ERROR PROTECTION INSERTED (arg_4-try_3) */
            if ErrorIn(try_5) {try_3 = try_5
            } else {
            arg_4 = ToSymbol(OBJ(try_5))
            try_3 = EID{arg_4.String_I().Id(),0}
            }
            } 
          /* ERROR PROTECTION INSERTED (arg_2-try_1) */
          if ErrorIn(try_3) {try_1 = try_3
          } else {
          arg_2 = ToString(OBJ(try_3))
          try_1 = F_imported_function_string(arg_2).ToEID()
          }
          } 
        /* ERROR PROTECTION INSERTED (f-Result) */
        if ErrorIn(try_1) {Result = try_1
        } else {
        f = ANY(try_1)
        Result = f.ToEID()
        }
        } else {
        Result = EID{CFALSE.Id(),0}
        } 
      /* ERROR PROTECTION INSERTED (Result-Result) */
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
/* {0} The go function for: imported_function(s:string) [status=0] */
func F_imported_function_string (s *ClaireString ) *ClaireFunction  { 
  return  F_make_function_string(F_append_string(MakeString("#"),s))
  } 

// The EID go function for: imported_function @ string (throw: false) 
func E_imported_function_string (s EID) EID { 
  return F_imported_function_string(ToString(OBJ(s)) ).ToEID()} 

// cleans a pattern into a type
/* {0} The go function for: iClaire/type!(x:any) [status=0] */
func F_type_I_any (x *ClaireAny ) *ClaireType  { 
  // procedure body with s = type 
var Result *ClaireType  
  if (x.Isa.IsIn(C_list) == CTRUE) { 
    { var g0262 *ClaireList   = ToList(x)
      _ = g0262
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
    { var g0263 *ClaireParam   = To_Param(x)
      { var _CL_obj *ClaireParam   = To_Param(new(ClaireParam).Is(C_Param))
        _CL_obj.Arg = g0263.Arg
        /*class->class*/_CL_obj.Params = g0263.Params
        /*list->list*/{ 
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
          /*list->list*/} 
        Result = ToType(_CL_obj.Id())
        } 
      } 
    }  else if (x.Isa.IsIn(C_Reference) == CTRUE) { 
    Result = ToType(C_any.Id())
    }  else if (x.Isa.IsIn(C_type) == CTRUE) { 
    { var g0265 *ClaireType   = ToType(x)
      _ = g0265
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
/* {0} The go function for: self_eval(self:Defarray) [status=1] */
func (self *Defarray ) SelfEval () EID { 
  var Result EID 
  { var a *ClaireList   = self.Arg.Args
    { var ar *ClaireTable  
      var try_1 EID 
      /*g_try(v2:"try_1",loop:false) */
      { var arg_2 *ClaireSymbol  
        _ = arg_2
        var try_3 EID 
        /*g_try(v2:"try_3",loop:false) */
        try_3 = F_extract_symbol_any(a.At(1-1))
        /* ERROR PROTECTION INSERTED (arg_2-try_1) */
        if ErrorIn(try_3) {try_1 = try_3
        } else {
        arg_2 = ToSymbol(OBJ(try_3))
        try_1 = new(ClaireTable).IsNamed(C_table,arg_2).ToEID()
        }
        } 
      /* ERROR PROTECTION INSERTED (ar-Result) */
      if ErrorIn(try_1) {Result = try_1
      } else {
      ar = ToTable(OBJ(try_1))
      { var v *ClaireVariable   = To_Variable(a.At(2-1))
        { var s *ClaireTypeExpression  
          var try_4 EID 
          /*g_try(v2:"try_4",loop:false) */
          try_4 = F_extract_type_any(v.Range.Id())
          /* ERROR PROTECTION INSERTED (s-Result) */
          if ErrorIn(try_4) {Result = try_4
          } else {
          s = ToTypeExpression(OBJ(try_4))
          { var e *ClaireAny  
            var try_5 EID 
            /*g_try(v2:"try_5",loop:false) */
            { var l *ClaireList  
              var try_6 EID 
              /*g_try(v2:"try_6",loop:false) */
              try_6 = a.Cdr()
              /* ERROR PROTECTION INSERTED (l-try_5) */
              if ErrorIn(try_6) {try_5 = try_6
              } else {
              l = ToList(OBJ(try_6))
              { var b *ClaireAny  
                var try_7 EID 
                /*g_try(v2:"try_7",loop:false) */
                try_7 = F_lexical_build_any(self.Body,l,0)
                /* ERROR PROTECTION INSERTED (b-try_5) */
                if ErrorIn(try_7) {try_5 = try_7
                } else {
                b = ANY(try_7)
                var g0275I *ClaireBoolean  
                { var arg_8 *ClaireAny  
                  _ = arg_8
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
            /* ERROR PROTECTION INSERTED (e-Result) */
            if ErrorIn(try_5) {Result = try_5
            } else {
            e = ANY(try_5)
            { var d *ClaireAny  
              var try_9 EID 
              /*g_try(v2:"try_9",loop:false) */
              if (e.Isa.IsIn(C_lambda) == CTRUE) { 
                try_9 = EID{CNULL,0}
                } else {
                try_9 = EVAL(self.Body)
                } 
              /* ERROR PROTECTION INSERTED (d-Result) */
              if ErrorIn(try_9) {Result = try_9
              } else {
              d = ANY(try_9)
              /*g_try(v2:"Result",loop:true) */
              { 
                var va_arg1 *ClaireRelation  
                var va_arg2 *ClaireType  
                va_arg1 = ToRelation(ar.Id())
                var try_10 EID 
                /*g_try(v2:"try_10",loop:false) */
                try_10 = F_extract_pattern_any(self.SetArg,CNIL)
                /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                if ErrorIn(try_10) {Result = try_10
                } else {
                va_arg2 = ToType(OBJ(try_10))
                va_arg1.Range = va_arg2
                /*type->type*/Result = EID{va_arg2.Id(),0}
                }
                } 
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              /*g_try(v2:"Result",loop:true) */
              if (ar.Range.Id() == CNULL) { 
                { var _CL_obj *Core.RangeError   = Core.ToRangeError(new(Core.RangeError).Is(Core.C_range_error))
                  _CL_obj.Cause = C_table.Id()
                  /*any->any*/_CL_obj.Arg = self.SetArg
                  /*any->any*//*g_try(v2:"Result",loop:true) */
                  Result = Core.F_write_property(C_Language_wrong,ToObject(_CL_obj.Id()),C_type.Id())
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  Result = _CL_obj.Close()
                  }
                  } 
                } else {
                Result = EID{CFALSE.Id(),0}
                } 
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              /*g_try(v2:"Result",loop:true) */
              if (d != CNULL) { 
                if (ar.Range.Contains(d) != CTRUE) { 
                  { var _CL_obj *Core.RangeError   = Core.ToRangeError(new(Core.RangeError).Is(Core.C_range_error))
                    _CL_obj.Cause = ar.Id()
                    /*any->any*/_CL_obj.Arg = d
                    /*any->any*//*g_try(v2:"Result",loop:true) */
                    Result = Core.F_write_property(C_Language_wrong,ToObject(_CL_obj.Id()),ar.Range.Id())
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    Result = _CL_obj.Close()
                    }
                    } 
                  } else {
                  Result = EID{CFALSE.Id(),0}
                  } 
                }  else if (ToType(s.Id()).Included(ToType(C_integer.Id())) == CTRUE) { 
                d = MakeInteger(0).Id()
                Result = d.ToEID()
                }  else if (ToType(s.Id()).Included(ToType(C_float.Id())) == CTRUE) { 
                d = MakeFloat(0).Id()
                Result = d.ToEID()
                } else {
                Result = EID{CFALSE.Id(),0}
                } 
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              v.Range = ToType(s.Id())
              /*type->type*//*g_try(v2:"Result",loop:true) */
              Result = F_attach_comment_any(ar.Id())
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              if (ar.Range.Class_I().IsIn(C_set) == CTRUE) { 
                ar.Multivalued_ask = CTRUE
                /*boolean->boolean*/} 
              /*g_try(v2:"Result",loop:true) */
              if (a.Length() == 2) { 
                ar.Domain = ToType(s.Id())
                /*type->type*/if (s.Isa.IsIn(C_Interval) == CTRUE) { 
                  { var g0269 *ClaireInterval   = To_Interval(s.Id())
                    ar.Params = MakeInteger((g0269.Arg1-1)).Id()
                    /*any->any*/ar.Graph = Core.F_typed_copy_list_type(ar.Range,Core.F_size_Interval(g0269),d).Id()
                    /*any->any*/} 
                  } else {
                  ar.Params = C_any.Id()
                  /*any->any*/ar.GraphInit()
                  } 
                if (e.Isa.IsIn(C_lambda) == CTRUE) { 
                  { var g0271 *ClaireLambda   = ToLambda(e)
                    _ = g0271
                    { 
                      var y *ClaireAny  
                      _ = y
                      Result= EID{CFALSE.Id(),0}
                      var y_support *ClaireList  
                      var try_11 EID 
                      /*g_try(v2:"try_11",loop:false) */
                      try_11 = Core.F_enumerate_any(ar.Domain.Id())
                      /* ERROR PROTECTION INSERTED (y_support-Result) */
                      if ErrorIn(try_11) {Result = try_11
                      } else {
                      y_support = ToList(OBJ(try_11))
                      y_len := y_support.Length()
                      for i_it := 0; i_it < y_len; i_it++ { 
                        y = y_support.At(i_it)
                        var loop_12 EID 
                        _ = loop_12
                        /*g_try(v2:"loop_12",loop:tuple("Result", EID)) */
                        { var arg_13 *ClaireAny  
                          _ = arg_13
                          var try_14 EID 
                          /*g_try(v2:"try_14",loop:false) */
                          try_14 = Core.F_funcall_lambda1(g0271,y)
                          /* ERROR PROTECTION INSERTED (arg_13-loop_12) */
                          if ErrorIn(try_14) {loop_12 = try_14
                          } else {
                          arg_13 = ANY(try_14)
                          loop_12 = Core.F_nth_equal_table1(ar,y,arg_13)
                          }
                          } 
                        /* ERROR PROTECTION INSERTED (loop_12-Result) */
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
                    /*any->any*/Result = va_arg2.ToEID()
                    } 
                  } 
                } else {
                { var s2 *ClaireTypeExpression  
                  var try_15 EID 
                  /*g_try(v2:"try_15",loop:false) */
                  try_15 = F_extract_type_any(To_Variable(a.At(3-1)).Range.Id())
                  /* ERROR PROTECTION INSERTED (s2-Result) */
                  if ErrorIn(try_15) {Result = try_15
                  } else {
                  s2 = ToTypeExpression(OBJ(try_15))
                  ar.Domain = ToType(MakeConstantList(s.Id(),s2.Id()).Tuple_I().Id())
                  /*type->type*/To_Variable(a.At(3-1)).Range = ToType(s2.Id())
                  /*type->type*//*g_try(v2:"Result",loop:true) */
                  if ((s.Isa.IsIn(C_Interval) == CTRUE) && 
                      (s2.Isa.IsIn(C_Interval) == CTRUE)) { 
                    
                    /*g_try(v2:"Result",loop:true) */
                    { 
                      var va_arg1 *ClaireTable  
                      var va_arg2 *ClaireAny  
                      va_arg1 = ar
                      var try_16 EID 
                      /*g_try(v2:"try_16",loop:false) */
                      { 
                        var v_bag_arg *ClaireAny  
                        try_16= EID{ToType(C_integer.Id()).EmptyList().Id(),0}
                        var try_17 EID 
                        /*g_try(v2:"try_17",loop:false) */
                        try_17 = Core.F_CALL(C_size,ARGS(EID{s2.Id(),0}))
                        /* ERROR PROTECTION INSERTED (v_bag_arg-try_16) */
                        if ErrorIn(try_17) {try_16 = try_17
                        } else {
                        v_bag_arg = ANY(try_17)
                        ToList(OBJ(try_16)).AddFast(v_bag_arg)
                        var try_18 EID 
                        /*g_try(v2:"try_18",loop:false) */
                        { var arg_19 int 
                          _ = arg_19
                          var try_20 EID 
                          /*g_try(v2:"try_20",loop:false) */
                          { var arg_21 int 
                            _ = arg_21
                            var try_22 EID 
                            /*g_try(v2:"try_22",loop:false) */
                            { var arg_23 *ClaireAny  
                              _ = arg_23
                              var try_24 EID 
                              /*g_try(v2:"try_24",loop:false) */
                              try_24 = Core.F_CALL(C_size,ARGS(EID{s2.Id(),0}))
                              /* ERROR PROTECTION INSERTED (arg_23-try_22) */
                              if ErrorIn(try_24) {try_22 = try_24
                              } else {
                              arg_23 = ANY(try_24)
                              try_22 = EID{C__INT,IVAL((To_Interval(s.Id()).Arg1*ToInteger(arg_23).Value))}
                              }
                              } 
                            /* ERROR PROTECTION INSERTED (arg_21-try_20) */
                            if ErrorIn(try_22) {try_20 = try_22
                            } else {
                            arg_21 = INT(try_22)
                            try_20 = EID{C__INT,IVAL((arg_21+To_Interval(s2.Id()).Arg1))}
                            }
                            } 
                          /* ERROR PROTECTION INSERTED (arg_19-try_18) */
                          if ErrorIn(try_20) {try_18 = try_20
                          } else {
                          arg_19 = INT(try_20)
                          try_18 = EID{C__INT,IVAL((arg_19-1))}
                          }
                          } 
                        /* ERROR PROTECTION INSERTED (v_bag_arg-try_16) */
                        if ErrorIn(try_18) {try_16 = try_18
                        } else {
                        v_bag_arg = ANY(try_18)
                        ToList(OBJ(try_16)).AddFast(v_bag_arg)}}
                        } 
                      /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                      if ErrorIn(try_16) {Result = try_16
                      } else {
                      va_arg2 = ANY(try_16)
                      va_arg1.Params = va_arg2
                      /*any->any*/Result = va_arg2.ToEID()
                      }
                      } 
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    { 
                      var va_arg1 *ClaireTable  
                      var va_arg2 *ClaireAny  
                      va_arg1 = ar
                      var try_25 EID 
                      /*g_try(v2:"try_25",loop:false) */
                      { var arg_26 int 
                        _ = arg_26
                        var try_27 EID 
                        /*g_try(v2:"try_27",loop:false) */
                        { var arg_28 *ClaireAny  
                          _ = arg_28
                          var try_30 EID 
                          /*g_try(v2:"try_30",loop:false) */
                          try_30 = Core.F_CALL(C_size,ARGS(EID{s.Id(),0}))
                          /* ERROR PROTECTION INSERTED (arg_28-try_27) */
                          if ErrorIn(try_30) {try_27 = try_30
                          } else {
                          arg_28 = ANY(try_30)
                          { var arg_29 *ClaireAny  
                            _ = arg_29
                            var try_31 EID 
                            /*g_try(v2:"try_31",loop:false) */
                            try_31 = Core.F_CALL(C_size,ARGS(EID{s2.Id(),0}))
                            /* ERROR PROTECTION INSERTED (arg_29-try_27) */
                            if ErrorIn(try_31) {try_27 = try_31
                            } else {
                            arg_29 = ANY(try_31)
                            try_27 = F_times_integer(ToInteger(arg_28).Value,ToInteger(arg_29).Value)
                            }
                            } 
                          }
                          } 
                        /* ERROR PROTECTION INSERTED (arg_26-try_25) */
                        if ErrorIn(try_27) {try_25 = try_27
                        } else {
                        arg_26 = INT(try_27)
                        try_25 = EID{Core.F_typed_copy_list_type(ar.Range,arg_26,d).Id(),0}
                        }
                        } 
                      /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                      if ErrorIn(try_25) {Result = try_25
                      } else {
                      va_arg2 = ANY(try_25)
                      va_arg1.Graph = va_arg2
                      /*any->any*/Result = va_arg2.ToEID()
                      }
                      } 
                    }
                    } else {
                    
                    ar.Params = C_any.Id()
                    /*any->any*/ar.GraphInit()
                    Result = EVOID
                    } 
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  
                  if (e.Isa.IsIn(C_lambda) == CTRUE) { 
                    { var g0273 *ClaireLambda   = ToLambda(e)
                      _ = g0273
                      { 
                        var y1 *ClaireAny  
                        _ = y1
                        Result= EID{CFALSE.Id(),0}
                        var y1_support *ClaireList  
                        var try_32 EID 
                        /*g_try(v2:"try_32",loop:false) */
                        try_32 = Core.F_enumerate_any(s.Id())
                        /* ERROR PROTECTION INSERTED (y1_support-Result) */
                        if ErrorIn(try_32) {Result = try_32
                        } else {
                        y1_support = ToList(OBJ(try_32))
                        y1_len := y1_support.Length()
                        for i_it := 0; i_it < y1_len; i_it++ { 
                          y1 = y1_support.At(i_it)
                          var loop_33 EID 
                          _ = loop_33
                          /*g_try(v2:"loop_33",loop:tuple("Result", EID)) */
                          { 
                            var y2 *ClaireAny  
                            _ = y2
                            loop_33= EID{CFALSE.Id(),0}
                            var y2_support *ClaireList  
                            var try_34 EID 
                            /*g_try(v2:"try_34",loop:false) */
                            try_34 = Core.F_enumerate_any(s2.Id())
                            /* ERROR PROTECTION INSERTED (y2_support-loop_33) */
                            if ErrorIn(try_34) {loop_33 = try_34
                            } else {
                            y2_support = ToList(OBJ(try_34))
                            y2_len := y2_support.Length()
                            for i_it := 0; i_it < y2_len; i_it++ { 
                              y2 = y2_support.At(i_it)
                              var loop_35 EID 
                              _ = loop_35
                              /*g_try(v2:"loop_35",loop:tuple("loop_33", EID)) */
                              { var arg_36 *ClaireAny  
                                _ = arg_36
                                var try_37 EID 
                                /*g_try(v2:"try_37",loop:false) */
                                try_37 = Core.F_CALL(C_funcall,ARGS(EID{g0273.Id(),0},y1.ToEID(),y2.ToEID()))
                                /* ERROR PROTECTION INSERTED (arg_36-loop_35) */
                                if ErrorIn(try_37) {loop_35 = try_37
                                } else {
                                arg_36 = ANY(try_37)
                                loop_35 = Core.F_nth_equal_table2(ar,y1,y2,arg_36)
                                }
                                } 
                              /* ERROR PROTECTION INSERTED (loop_35-loop_33) */
                              if ErrorIn(loop_35) {loop_33 = loop_35
                              break
                              } else {
                              }}
                              } 
                            } 
                          /* ERROR PROTECTION INSERTED (loop_33-Result) */
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
                      /*any->any*/Result = va_arg2.ToEID()
                      } 
                    } 
                  }
                  }
                  } 
                } 
              /* ERROR PROTECTION INSERTED (Result-Result) */
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
/* {0} The go function for: self_print(self:demon) [status=1] */
func (self *LanguageDemon ) SelfPrint () EID { 
  var Result EID 
  self.Pname.Princ()
  Result = EVOID
  return Result} 

// The EID go function for: self_print @ demon (throw: true) 
func E_self_print_demon (self EID) EID { 
  return ToLanguageDemon(OBJ(self)).SelfPrint( )} 

/* {0} The go function for: funcall(self:demon,x:any,y:any) [status=1] */
func F_funcall_demon1 (self *LanguageDemon ,x *ClaireAny ,y *ClaireAny ) EID { 
  var Result EID 
  Result = Core.F_CALL(C_funcall,ARGS(EID{self.Formula.Id(),0},x.ToEID(),y.ToEID()))
  return Result} 

// The EID go function for: funcall @ list<type_expression>(demon, any, any) (throw: true) 
func E_funcall_demon1 (self EID,x EID,y EID) EID { 
  return F_funcall_demon1(ToLanguageDemon(OBJ(self)),ANY(x),ANY(y) )} 

/* {0} The go function for: funcall(self:demon,x:any,y:any,z:any) [status=1] */
func F_funcall_demon2 (self *LanguageDemon ,x *ClaireAny ,y *ClaireAny ,z *ClaireAny ) EID { 
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
/* {0} The go function for: self_eval(self:Defrule) [status=1] */
func (self *Defrule ) SelfEval () EID { 
  var Result EID 
  if (self.Args.At(1-1) != ClEnv.Id()) { 
    Result = Core.F_CALL(C_eval_rule,ARGS(EID{self.Id(),0}))
    } else {
    { var _Zcondition *ClaireAny   = self.Arg
      { var ru *ClaireAny   = self.Ident.Get()
        ru.Isa = C_Language_rule_object
        /*class->class*/C_Language_rule_object.Instances.AddFast(ru)/*t=any,s=void*/
        { var g0276 *ClaireTuple  
          var try_1 EID 
          /*g_try(v2:"try_1",loop:false) */
          try_1 = F_make_filter_any(_Zcondition)
          /* ERROR PROTECTION INSERTED (g0276-Result) */
          if ErrorIn(try_1) {Result = try_1
          } else {
          g0276 = ToTuple(OBJ(try_1))
          { var R *ClaireAny   = ToList(g0276.Id()).At(1-1)
            { var lvar *ClaireAny   = ToList(g0276.Id()).At(2-1)
              { var d *LanguageDemon  
                var try_2 EID 
                /*g_try(v2:"try_2",loop:false) */
                { var arg_3 *ClaireAny  
                  _ = arg_3
                  var try_4 EID 
                  /*g_try(v2:"try_4",loop:false) */
                  try_4 = F_lexical_build_any(self.Body,ToList(lvar),0)
                  /* ERROR PROTECTION INSERTED (arg_3-try_2) */
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
                /* ERROR PROTECTION INSERTED (d-Result) */
                if ErrorIn(try_2) {Result = try_2
                } else {
                d = ToLanguageDemon(OBJ(try_2))
                /*g_try(v2:"Result",loop:true) */
                if (C_function.Id() == Core.F_owner_any(ANY(Core.F_CALL(C_if_write,ARGS(R.ToEID())))).Id()) { 
                  Result = ToException(Core.C_general_error.Make(MakeString("cannot define a new rule on ~S which is closed").Id(),MakeConstantList(R).Id())).Close()
                  } else {
                  Result = EID{CFALSE.Id(),0}
                  } 
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                Core.F_tformat_string(MakeString("we have defined a demon ~S for ~S \n"),0,MakeConstantList(d.Id(),R))
                Core.F_put_table(C_demons,R,ToList(Core.F_get_table(C_demons,R)).AddFast(d.Id()).Id()/*t=any,s=any*/)
                /*g_try(v2:"Result",loop:true) */
                Result = Core.F_nth_put_table(C_Language_last_rule,R,ru)
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                /*g_try(v2:"Result",loop:true) */
                if (ToList(Core.F_get_table(C_demons,R)).Length() == 1) { 
                  Result = F_eval_if_write_relation(ToRelation(R))
                  } else {
                  Result = EID{CFALSE.Id(),0}
                  } 
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                var g0279I *ClaireBoolean  
                if (R.Isa.IsIn(C_property) == CTRUE) { 
                  { var g0278 *ClaireProperty   = ToProperty(R)
                    _ = g0278
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
/* {0} The go function for: eventMethod?(r:relation) [status=0] */
func F_eventMethod_ask_relation2 (r *ClaireRelation ) *ClaireBoolean  { 
  // procedure body with s = boolean 
var Result *ClaireBoolean  
  if (r.Isa.IsIn(C_property) == CTRUE) { 
    { var g0280 *ClaireProperty   = ToProperty(r.Id())
      _ = g0280
      { var arg_1 *ClaireAny  
        _ = arg_1
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
/* {0} The go function for: make_filter(cond:any) [status=1] */
func F_make_filter_any (cond *ClaireAny ) EID { 
  var Result EID 
  { var c *ClaireAny  
    if (cond.Isa.IsIn(C_And) == CTRUE) { 
      { var g0281 *And   = To_And(cond)
        _ = g0281
        c = g0281.Args.At(1-1)
        } 
      } else {
      c = cond
      } 
    
    var g0287I *ClaireBoolean  
    if (c.Isa.IsIn(C_Call) == CTRUE) { 
      { var g0283 *Call   = To_Call(c)
        g0287I = MakeBoolean(((g0283.Selector.Id() == Core.C_write.Id()) || 
            (g0283.Selector.Id() == C_nth_equal.Id())) && (g0283.Args.At(1-1).Isa.IsIn(C_relation) == CTRUE))
        } 
      } else {
      g0287I = CFALSE
      } 
    if (g0287I == CTRUE) { 
      { var R *ClaireRelation   = ToRelation(ToList(OBJ(Core.F_CALL(C_args,ARGS(c.ToEID())))).At(1-1))
        { var x *ClaireVariable  
          var try_1 EID 
          /*g_try(v2:"try_1",loop:false) */
          { var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
            /*g_try(v2:"try_1",loop:false) */
            { 
              var va_arg1 *ClaireVariable  
              var va_arg2 *ClaireSymbol  
              va_arg1 = _CL_obj
              var try_2 EID 
              /*g_try(v2:"try_2",loop:false) */
              try_2 = F_extract_symbol_any(ToList(OBJ(Core.F_CALL(C_args,ARGS(c.ToEID())))).At(2-1))
              /* ERROR PROTECTION INSERTED (va_arg2-try_1) */
              if ErrorIn(try_2) {try_1 = try_2
              } else {
              va_arg2 = ToSymbol(OBJ(try_2))
              va_arg1.Pname = va_arg2
              /*symbol->symbol*/try_1 = EID{va_arg2.Id(),0}
              }
              } 
            /* ERROR PROTECTION INSERTED (try_1-try_1) */
            if !ErrorIn(try_1) {
            _CL_obj.Range = R.Domain
            /*type->type*/try_1 = EID{_CL_obj.Id(),0}
            }
            } 
          /* ERROR PROTECTION INSERTED (x-Result) */
          if ErrorIn(try_1) {Result = try_1
          } else {
          x = To_Variable(OBJ(try_1))
          { var y1 *ClaireAny   = ToList(OBJ(Core.F_CALL(C_args,ARGS(c.ToEID())))).At(3-1)
            /*g_try(v2:"Result",loop:true) */
            if (R.Multivalued_ask == CTRUE) { 
              Result = ToException(Core.C_general_error.Make(MakeString("[188] wrong event filter ~S for multi-valued relation").Id(),MakeConstantList(c,R.Id()).Id())).Close()
              } else {
              Result = EID{CFALSE.Id(),0}
              } 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            var g0288I *ClaireBoolean  
            if (y1.Isa.IsIn(C_Call) == CTRUE) { 
              { var g0284 *Call   = To_Call(y1)
                _ = g0284
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
                /*g_try(v2:"try_3",loop:false) */
                { 
                  var v_bag_arg *ClaireAny  
                  try_3= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                  ToList(OBJ(try_3)).AddFast(x.Id())
                  var try_4 EID 
                  /*g_try(v2:"try_4",loop:false) */
                  { var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
                    /*g_try(v2:"try_4",loop:false) */
                    { 
                      var va_arg1 *ClaireVariable  
                      var va_arg2 *ClaireSymbol  
                      va_arg1 = _CL_obj
                      var try_5 EID 
                      /*g_try(v2:"try_5",loop:false) */
                      try_5 = F_extract_symbol_any(ToList(OBJ(Core.F_CALL(C_args,ARGS(y1.ToEID())))).At(1-1))
                      /* ERROR PROTECTION INSERTED (va_arg2-try_4) */
                      if ErrorIn(try_5) {try_4 = try_5
                      } else {
                      va_arg2 = ToSymbol(OBJ(try_5))
                      va_arg1.Pname = va_arg2
                      /*symbol->symbol*/try_4 = EID{va_arg2.Id(),0}
                      }
                      } 
                    /* ERROR PROTECTION INSERTED (try_4-try_4) */
                    if !ErrorIn(try_4) {
                    _CL_obj.Range = R.Range
                    /*type->type*/try_4 = EID{_CL_obj.Id(),0}
                    }
                    } 
                  /* ERROR PROTECTION INSERTED (v_bag_arg-try_3) */
                  if ErrorIn(try_4) {try_3 = try_4
                  } else {
                  v_bag_arg = ANY(try_4)
                  ToList(OBJ(try_3)).AddFast(v_bag_arg)
                  var try_6 EID 
                  /*g_try(v2:"try_6",loop:false) */
                  { var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
                    /*g_try(v2:"try_6",loop:false) */
                    { 
                      var va_arg1 *ClaireVariable  
                      var va_arg2 *ClaireSymbol  
                      va_arg1 = _CL_obj
                      var try_7 EID 
                      /*g_try(v2:"try_7",loop:false) */
                      try_7 = F_extract_symbol_any(ToList(OBJ(Core.F_CALL(C_args,ARGS(y1.ToEID())))).At(2-1))
                      /* ERROR PROTECTION INSERTED (va_arg2-try_6) */
                      if ErrorIn(try_7) {try_6 = try_7
                      } else {
                      va_arg2 = ToSymbol(OBJ(try_7))
                      va_arg1.Pname = va_arg2
                      /*symbol->symbol*/try_6 = EID{va_arg2.Id(),0}
                      }
                      } 
                    /* ERROR PROTECTION INSERTED (try_6-try_6) */
                    if !ErrorIn(try_6) {
                    _CL_obj.Range = R.Range
                    /*type->type*/try_6 = EID{_CL_obj.Id(),0}
                    }
                    } 
                  /* ERROR PROTECTION INSERTED (v_bag_arg-try_3) */
                  if ErrorIn(try_6) {try_3 = try_6
                  } else {
                  v_bag_arg = ANY(try_6)
                  ToList(OBJ(try_3)).AddFast(v_bag_arg)}}
                  } 
                /* ERROR PROTECTION INSERTED (v_bag_arg-Result) */
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
                /*g_try(v2:"try_8",loop:false) */
                { 
                  var v_bag_arg *ClaireAny  
                  try_8= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                  ToList(OBJ(try_8)).AddFast(x.Id())
                  var try_9 EID 
                  /*g_try(v2:"try_9",loop:false) */
                  { var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
                    /*g_try(v2:"try_9",loop:false) */
                    { 
                      var va_arg1 *ClaireVariable  
                      var va_arg2 *ClaireSymbol  
                      va_arg1 = _CL_obj
                      var try_10 EID 
                      /*g_try(v2:"try_10",loop:false) */
                      try_10 = F_extract_symbol_any(y1)
                      /* ERROR PROTECTION INSERTED (va_arg2-try_9) */
                      if ErrorIn(try_10) {try_9 = try_10
                      } else {
                      va_arg2 = ToSymbol(OBJ(try_10))
                      va_arg1.Pname = va_arg2
                      /*symbol->symbol*/try_9 = EID{va_arg2.Id(),0}
                      }
                      } 
                    /* ERROR PROTECTION INSERTED (try_9-try_9) */
                    if !ErrorIn(try_9) {
                    _CL_obj.Range = F_safeRange_relation(R)
                    /*type->type*/try_9 = EID{_CL_obj.Id(),0}
                    }
                    } 
                  /* ERROR PROTECTION INSERTED (v_bag_arg-try_8) */
                  if ErrorIn(try_9) {try_8 = try_9
                  } else {
                  v_bag_arg = ANY(try_9)
                  ToList(OBJ(try_8)).AddFast(v_bag_arg)
                  { var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
                    _CL_obj.Pname = Core.F_gensym_void()
                    /*symbol->symbol*/_CL_obj.Range = F_safeRange_relation(R)
                    /*type->type*/v_bag_arg = _CL_obj.Id()
                    } 
                  ToList(OBJ(try_8)).AddFast(v_bag_arg)}
                  } 
                /* ERROR PROTECTION INSERTED (v_bag_arg-Result) */
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
        { var g0285 *Call   = To_Call(c)
          g0289I = MakeBoolean((g0285.Selector.Id() == C_add.Id()) && (g0285.Args.At(1-1).Isa.IsIn(C_relation) == CTRUE))
          } 
        } else {
        g0289I = CFALSE
        } 
      if (g0289I == CTRUE) { 
        { var R *ClaireRelation   = ToRelation(ToList(OBJ(Core.F_CALL(C_args,ARGS(c.ToEID())))).At(1-1))
          { var x *ClaireVariable  
            _ = x
            var try_11 EID 
            /*g_try(v2:"try_11",loop:false) */
            { var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
              /*g_try(v2:"try_11",loop:false) */
              { 
                var va_arg1 *ClaireVariable  
                var va_arg2 *ClaireSymbol  
                va_arg1 = _CL_obj
                var try_12 EID 
                /*g_try(v2:"try_12",loop:false) */
                try_12 = F_extract_symbol_any(ToList(OBJ(Core.F_CALL(C_args,ARGS(c.ToEID())))).At(2-1))
                /* ERROR PROTECTION INSERTED (va_arg2-try_11) */
                if ErrorIn(try_12) {try_11 = try_12
                } else {
                va_arg2 = ToSymbol(OBJ(try_12))
                va_arg1.Pname = va_arg2
                /*symbol->symbol*/try_11 = EID{va_arg2.Id(),0}
                }
                } 
              /* ERROR PROTECTION INSERTED (try_11-try_11) */
              if !ErrorIn(try_11) {
              _CL_obj.Range = R.Domain
              /*type->type*/try_11 = EID{_CL_obj.Id(),0}
              }
              } 
            /* ERROR PROTECTION INSERTED (x-Result) */
            if ErrorIn(try_11) {Result = try_11
            } else {
            x = To_Variable(OBJ(try_11))
            { var y *ClaireVariable  
              _ = y
              var try_13 EID 
              /*g_try(v2:"try_13",loop:false) */
              { var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
                /*g_try(v2:"try_13",loop:false) */
                { 
                  var va_arg1 *ClaireVariable  
                  var va_arg2 *ClaireSymbol  
                  va_arg1 = _CL_obj
                  var try_14 EID 
                  /*g_try(v2:"try_14",loop:false) */
                  try_14 = F_extract_symbol_any(ToList(OBJ(Core.F_CALL(C_args,ARGS(c.ToEID())))).At(3-1))
                  /* ERROR PROTECTION INSERTED (va_arg2-try_13) */
                  if ErrorIn(try_14) {try_13 = try_14
                  } else {
                  va_arg2 = ToSymbol(OBJ(try_14))
                  va_arg1.Pname = va_arg2
                  /*symbol->symbol*/try_13 = EID{va_arg2.Id(),0}
                  }
                  } 
                /* ERROR PROTECTION INSERTED (try_13-try_13) */
                if !ErrorIn(try_13) {
                _CL_obj.Range = R.Range
                /*type->type*/try_13 = EID{_CL_obj.Id(),0}
                }
                } 
              /* ERROR PROTECTION INSERTED (y-Result) */
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
          { var g0286 *Call   = To_Call(c)
            _ = g0286
            g0290I = Equal(MakeInteger(g0286.Args.Length()).Id(),MakeInteger(2).Id())
            } 
          } else {
          g0290I = CFALSE
          } 
        if (g0290I == CTRUE) { 
          { var R *ClaireProperty   = To_Call(c).Selector
            { var x *ClaireVariable  
              _ = x
              var try_15 EID 
              /*g_try(v2:"try_15",loop:false) */
              { var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
                /*g_try(v2:"try_15",loop:false) */
                { 
                  var va_arg1 *ClaireVariable  
                  var va_arg2 *ClaireSymbol  
                  va_arg1 = _CL_obj
                  var try_16 EID 
                  /*g_try(v2:"try_16",loop:false) */
                  try_16 = F_extract_symbol_any(ToList(OBJ(Core.F_CALL(C_args,ARGS(c.ToEID())))).At(1-1))
                  /* ERROR PROTECTION INSERTED (va_arg2-try_15) */
                  if ErrorIn(try_16) {try_15 = try_16
                  } else {
                  va_arg2 = ToSymbol(OBJ(try_16))
                  va_arg1.Pname = va_arg2
                  /*symbol->symbol*/try_15 = EID{va_arg2.Id(),0}
                  }
                  } 
                /* ERROR PROTECTION INSERTED (try_15-try_15) */
                if !ErrorIn(try_15) {
                _CL_obj.Range = R.Domain
                /*type->type*/try_15 = EID{_CL_obj.Id(),0}
                }
                } 
              /* ERROR PROTECTION INSERTED (x-Result) */
              if ErrorIn(try_15) {Result = try_15
              } else {
              x = To_Variable(OBJ(try_15))
              { var y *ClaireVariable  
                _ = y
                var try_17 EID 
                /*g_try(v2:"try_17",loop:false) */
                { var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
                  /*g_try(v2:"try_17",loop:false) */
                  { 
                    var va_arg1 *ClaireVariable  
                    var va_arg2 *ClaireSymbol  
                    va_arg1 = _CL_obj
                    var try_18 EID 
                    /*g_try(v2:"try_18",loop:false) */
                    try_18 = F_extract_symbol_any(ToList(OBJ(Core.F_CALL(C_args,ARGS(c.ToEID())))).At(2-1))
                    /* ERROR PROTECTION INSERTED (va_arg2-try_17) */
                    if ErrorIn(try_18) {try_17 = try_18
                    } else {
                    va_arg2 = ToSymbol(OBJ(try_18))
                    va_arg1.Pname = va_arg2
                    /*symbol->symbol*/try_17 = EID{va_arg2.Id(),0}
                    }
                    } 
                  /* ERROR PROTECTION INSERTED (try_17-try_17) */
                  if !ErrorIn(try_17) {
                  _CL_obj.Range = R.Range
                  /*type->type*/try_17 = EID{_CL_obj.Id(),0}
                  }
                  } 
                /* ERROR PROTECTION INSERTED (y-Result) */
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
/* {0} The go function for: make_demon(R:relation,n:symbol,lvar:list[Variable],cond:any,conc:any) [status=1] */
func F_make_demon_relation (R *ClaireRelation ,n *ClaireSymbol ,lvar *ClaireList ,cond *ClaireAny ,conc *ClaireAny ) EID { 
  var Result EID 
  { var x *ClaireVariable   = To_Variable(lvar.At(1-1))
    _ = x
    { var y *ClaireVariable   = To_Variable(lvar.At(2-1))
      _ = y
      { var _Ztest *ClaireAny  
        _ = _Ztest
        { var _CL_obj *Call   = To_Call(new(Call).Is(C_Call))
          _CL_obj.Selector = ToProperty(IfThenElse((R.Multivalued_ask == CTRUE),
            C__Z.Id(),
            C__equal.Id()))
          /*property->property*/_CL_obj.Args = MakeConstantList(y.Id(),F_readCall_relation(R,x.Id()).Id())
          /*list->list*/_Ztest = _CL_obj.Id()
          } 
        { var _Zbody *ClaireAny   = conc
          Core.F_tformat_string(MakeString("make a demon for ~S from ~S => ~S (name = ~S) \n"),0,MakeConstantList(R.Id(),
            cond,
            conc,
            n.Id()))
          if (C_if_write.Trace_I > ClEnv.Verbose) { 
            { var _CL_obj *Do   = To_Do(new(Do).Is(C_Do))
              { 
                var va_arg1 *Do  
                var va_arg2 *ClaireList  
                va_arg1 = _CL_obj
                { 
                  var v_bag_arg *ClaireAny  
                  va_arg2= ToType(CEMPTY.Id()).EmptyList()
                  { var _CL_obj *Call   = To_Call(new(Call).Is(C_Call))
                    _CL_obj.Selector = Core.C_format
                    /*property->property*/{ 
                      var va_arg1 *Call  
                      var va_arg2 *ClaireList  
                      va_arg1 = _CL_obj
                      { 
                        var v_bag_arg *ClaireAny  
                        va_arg2= ToType(CEMPTY.Id()).EmptyList()
                        va_arg2.AddFast(MakeString("--- trigger ~A(~S,~S)\n").Id())
                        { var _CL_obj *List   = To_List(new(List).Is(C_List))
                          _CL_obj.Args = MakeConstantList((n.String_I()).Id(),x.Id(),y.Id())
                          /*list->list*/v_bag_arg = _CL_obj.Id()
                          } 
                        va_arg2.AddFast(v_bag_arg)} 
                      va_arg1.Args = va_arg2
                      /*list->list*/} 
                    v_bag_arg = _CL_obj.Id()
                    } 
                  va_arg2.AddFast(v_bag_arg)
                  va_arg2.AddFast(conc)} 
                va_arg1.Args = va_arg2
                /*list->list*/} 
              conc = _CL_obj.Id()
              } 
            } 
          { var _CL_obj *If   = To_If(new(If).Is(C_If))
            _CL_obj.Arg = conc
            /*any->any*/_Zbody = _CL_obj.Id()
            } 
          /*g_try(v2:"Result",loop:true) */
          if (F_eventMethod_ask_relation2(R) == CTRUE) { 
            if (cond.Isa.IsIn(C_And) == CTRUE) { 
              { var g0291 *And   = To_And(cond)
                _ = g0291
                var try_1 EID 
                /*g_try(v2:"try_1",loop:false) */
                if (g0291.Args.Length() > 2) { 
                  { var _CL_obj *And   = To_And(new(And).Is(C_And))
                    /*g_try(v2:"try_1",loop:false) */
                    { 
                      var va_arg1 *And  
                      var va_arg2 *ClaireList  
                      va_arg1 = _CL_obj
                      var try_2 EID 
                      /*g_try(v2:"try_2",loop:false) */
                      try_2 = g0291.Args.Cdr()
                      /* ERROR PROTECTION INSERTED (va_arg2-try_1) */
                      if ErrorIn(try_2) {try_1 = try_2
                      } else {
                      va_arg2 = ToList(OBJ(try_2))
                      va_arg1.Args = va_arg2
                      /*list->list*/try_1 = EID{va_arg2.Id(),0}
                      }
                      } 
                    /* ERROR PROTECTION INSERTED (try_1-try_1) */
                    if !ErrorIn(try_1) {
                    try_1 = EID{_CL_obj.Id(),0}
                    }
                    } 
                  } else {
                  try_1 = g0291.Args.At(2-1).ToEID()
                  } 
                /* ERROR PROTECTION INSERTED (_Ztest-Result) */
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
            { var g0293 *And   = To_And(cond)
              _ = g0293
              var try_3 EID 
              /*g_try(v2:"try_3",loop:false) */
              { var _CL_obj *And   = To_And(new(And).Is(C_And))
                /*g_try(v2:"try_3",loop:false) */
                { 
                  var va_arg1 *And  
                  var va_arg2 *ClaireList  
                  va_arg1 = _CL_obj
                  var try_4 EID 
                  /*g_try(v2:"try_4",loop:false) */
                  { var arg_5 *ClaireList  
                    _ = arg_5
                    var try_6 EID 
                    /*g_try(v2:"try_6",loop:false) */
                    try_6 = g0293.Args.Cdr()
                    /* ERROR PROTECTION INSERTED (arg_5-try_4) */
                    if ErrorIn(try_6) {try_4 = try_6
                    } else {
                    arg_5 = ToList(OBJ(try_6))
                    try_4 = EID{MakeConstantList(_Ztest).Append(arg_5).Id(),0}
                    }
                    } 
                  /* ERROR PROTECTION INSERTED (va_arg2-try_3) */
                  if ErrorIn(try_4) {try_3 = try_4
                  } else {
                  va_arg2 = ToList(OBJ(try_4))
                  va_arg1.Args = va_arg2
                  /*list->list*/try_3 = EID{va_arg2.Id(),0}
                  }
                  } 
                /* ERROR PROTECTION INSERTED (try_3-try_3) */
                if !ErrorIn(try_3) {
                try_3 = EID{_CL_obj.Id(),0}
                }
                } 
              /* ERROR PROTECTION INSERTED (_Ztest-Result) */
              if ErrorIn(try_3) {Result = try_3
              } else {
              _Ztest = ANY(try_3)
              Result = _Ztest.ToEID()
              }
              } 
            } else {
            Result = EID{CFALSE.Id(),0}
            } 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          if (_Zbody.Isa.IsIn(C_If) == CTRUE) { 
            { var g0294 *If   = To_If(_Zbody)
              g0294.Test = _Ztest
              /*any->any*/} 
            } 
          Core.F_tformat_string(MakeString("create a demon with name ~S \n"),0,MakeConstantList(n.Id()))
          { var _CL_obj *LanguageDemon   = ToLanguageDemon(new(LanguageDemon).Is(C_Language_demon))
            _CL_obj.Pname = n
            /*symbol->symbol*//*g_try(v2:"Result",loop:true) */
            { 
              var va_arg1 *LanguageDemon  
              var va_arg2 *ClaireLambda  
              va_arg1 = _CL_obj
              var try_7 EID 
              /*g_try(v2:"try_7",loop:false) */
              try_7 = F_lambda_I_list(lvar,_Zbody)
              /* ERROR PROTECTION INSERTED (va_arg2-Result) */
              if ErrorIn(try_7) {Result = try_7
              } else {
              va_arg2 = ToLambda(OBJ(try_7))
              va_arg1.Formula = va_arg2
              /*lambda->lambda*/Result = EID{va_arg2.Id(),0}
              }
              } 
            /* ERROR PROTECTION INSERTED (Result-Result) */
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
/* {0} The go function for: readCall(R:relation,x:any) [status=0] */
func F_readCall_relation (R *ClaireRelation ,x *ClaireAny ) *Call  { 
  // procedure body with s = Call 
var Result *Call  
  if (C_table.Id() == R.Isa.Id()) { 
    { var _CL_obj *Call   = To_Call(new(Call).Is(C_Call))
      _CL_obj.Selector = C_get
      /*property->property*/_CL_obj.Args = MakeConstantList(R.Id(),x)
      /*list->list*/Result = _CL_obj
      } 
    } else {
    { var _CL_obj *Call_plus   = To_Call_plus(new(Call_plus).Is(C_Call_plus))
      _CL_obj.Selector = ToProperty(R.Id())
      /*property->property*/_CL_obj.Args = MakeConstantList(x)
      /*list->list*/Result = To_Call(_CL_obj.Id())
      } 
    } 
  return Result} 

// The EID go function for: readCall @ relation (throw: false) 
func E_readCall_relation (R EID,x EID) EID { 
  return EID{F_readCall_relation(ToRelation(OBJ(R)),ANY(x) ).Id(),0}} 

// a small brother
/* {0} The go function for: putCall(R:relation,x:any,y:any) [status=0] */
func F_putCall_relation2 (R *ClaireRelation ,x *ClaireAny ,y *ClaireAny ) *Call  { 
  // procedure body with s = Call 
var Result *Call  
  if (R.Multivalued_ask == CTRUE) { 
    { var _CL_obj *Call   = To_Call(new(Call).Is(C_Call))
      _CL_obj.Selector = Core.C_add_value
      /*property->property*/_CL_obj.Args = MakeConstantList(R.Id(),x,y)
      /*list->list*/Result = _CL_obj
      } 
    } else {
    { var _CL_obj *Call   = To_Call(new(Call).Is(C_Call))
      _CL_obj.Selector = C_put
      /*property->property*/_CL_obj.Args = MakeConstantList(R.Id(),x,y)
      /*list->list*/Result = _CL_obj
      } 
    } 
  return Result} 

// The EID go function for: putCall @ relation (throw: false) 
func E_putCall_relation2 (R EID,x EID,y EID) EID { 
  return EID{F_putCall_relation2(ToRelation(OBJ(R)),ANY(x),ANY(y) ).Id(),0}} 

// v3.3 : find the range when we read the current value     
/* {0} The go function for: safeRange(x:relation) [status=0] */
func F_safeRange_relation (x *ClaireRelation ) *ClaireType  { 
  // procedure body with s = type 
var Result *ClaireType  
  if (x.Isa.IsIn(C_property) == CTRUE) { 
    { var g0295 *ClaireProperty   = ToProperty(x.Id())
      var g0299I *ClaireBoolean  
      { var arg_1 *ClaireAny  
        _ = arg_1
        { 
          var s *ClaireRestriction  
          _ = s
          var s_iter *ClaireAny  
          arg_1= CFALSE.Id()
          for _,s_iter = range(g0295.Restrictions.ValuesO()){ 
            s = ToRestriction(s_iter)
            var g0300I *ClaireBoolean  
            { var arg_2 *ClaireBoolean  
              _ = arg_2
              if (C_slot.Id() == s.Isa.Id()) { 
                { var g0296 *ClaireSlot   = ToSlot(s.Id())
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
    { var g0297 *ClaireTable   = ToTable(x.Id())
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
/* {0} The go function for: eval_if_write(R:relation) [status=1] */
func F_eval_if_write_relation (R *ClaireRelation ) EID { 
  var Result EID 
  { var l *ClaireList   = ToList(Core.F_get_table(C_demons,R.Id()))
    _ = l
    { var lvar *ClaireList   = ToLanguageDemon(l.ValuesO()[1-1]).Formula.Vars
      { var dv *ClaireVariable  
        { var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
          _CL_obj.Pname = Core.F_gensym_void()
          /*symbol->symbol*/_CL_obj.Range = ToType(C_Language_demon.Id())
          /*type->type*/dv = _CL_obj
          } 
        { var l1 *ClaireList   = MakeList(ToType(C_any.Id()),F_putCall_relation2(R,lvar.At(1-1),lvar.At(2-1)).Id())
          { var l2 *ClaireList  
            { 
              var v_bag_arg *ClaireAny  
              l2= ToType(C_any.Id()).EmptyList()
              { var _CL_obj *For   = To_For(new(For).Is(C_For))
                _CL_obj.ClaireVar = dv
                /*Variable->Variable*/{ 
                  var va_arg1 *Iteration  
                  var va_arg2 *ClaireAny  
                  va_arg1 = To_Iteration(_CL_obj.Id())
                  { var _CL_obj *Call   = To_Call(new(Call).Is(C_Call))
                    _CL_obj.Selector = C_nth
                    /*property->property*/_CL_obj.Args = MakeConstantList(C_demons.Id(),R.Id())
                    /*list->list*/va_arg2 = _CL_obj.Id()
                    } 
                  va_arg1.SetArg = va_arg2
                  /*any->any*/} 
                { 
                  var va_arg1 *Iteration  
                  var va_arg2 *ClaireAny  
                  va_arg1 = To_Iteration(_CL_obj.Id())
                  { var _CL_obj *Call   = To_Call(new(Call).Is(C_Call))
                    _CL_obj.Selector = C_funcall
                    /*property->property*/_CL_obj.Args = MakeConstantList(dv.Id()).Append(lvar)
                    /*list->list*/va_arg2 = _CL_obj.Id()
                    } 
                  va_arg1.Arg = va_arg2
                  /*any->any*/} 
                v_bag_arg = _CL_obj.Id()
                } 
              l2.AddFast(v_bag_arg)} 
            Core.F_tformat_string(MakeString("generate a if_write demon for ~S \n"),0,MakeConstantList(R.Id()))
            /*g_try(v2:"Result",loop:true) */
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
                /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
                loop_1 = Core.F_put_property2(C_range,ToObject(v),ToTypeExpression(OBJ(Core.F_CALL(C_range,ARGS(v.ToEID())))).Class_I().Id())
                /* ERROR PROTECTION INSERTED (loop_1-Result) */
                if ErrorIn(loop_1) {Result = loop_1
                break
                } else {
                }
                } 
              } 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            if (R.Inverse.Id() != CNULL) { 
              if (R.Multivalued_ask != CTRUE) { 
                { var arg_2 *Call  
                  _ = arg_2
                  { var _CL_obj *Call   = To_Call(new(Call).Is(C_Call))
                    _CL_obj.Selector = Core.C_Core_update_dash
                    /*property->property*/_CL_obj.Args = MakeConstantList(R.Inverse.Id(),lvar.At(3-1),lvar.At(1-1))
                    /*list->list*/arg_2 = _CL_obj
                    } 
                  l1 = l1.AddFast(arg_2.Id())/*t=any,s=list*/
                  } 
                } 
              l1 = l1.AddFast(F_putCall_relation2(R.Inverse,lvar.At(2-1),lvar.At(1-1)).Id())/*t=any,s=list*/
              } 
            { 
              var va_arg1 *ClaireRelation  
              var va_arg2 *ClaireAny  
              va_arg1 = R
              var try_3 EID 
              /*g_try(v2:"try_3",loop:false) */
              { var arg_4 *ComplexInstruction  
                _ = arg_4
                if (F_eventMethod_ask_relation2(R) == CTRUE) { 
                  { var _CL_obj *Do   = To_Do(new(Do).Is(C_Do))
                    _CL_obj.Args = l2
                    /*list->list*/arg_4 = To_ComplexInstruction(_CL_obj.Id())
                    } 
                  }  else if (R.Multivalued_ask == CTRUE) { 
                  { var _CL_obj *If   = To_If(new(If).Is(C_If))
                    { 
                      var va_arg1 *If  
                      var va_arg2 *ClaireAny  
                      va_arg1 = _CL_obj
                      { var _CL_obj *Call   = To_Call(new(Call).Is(C_Call))
                        _CL_obj.Selector = Core.C_not
                        /*property->property*/{ 
                          var va_arg1 *Call  
                          var va_arg2 *ClaireList  
                          va_arg1 = _CL_obj
                          { 
                            var v_bag_arg *ClaireAny  
                            va_arg2= ToType(CEMPTY.Id()).EmptyList()
                            { var _CL_obj *Call   = To_Call(new(Call).Is(C_Call))
                              _CL_obj.Selector = ToProperty(C__Z.Id())
                              /*property->property*/_CL_obj.Args = MakeConstantList(lvar.At(2-1),F_readCall_relation(R,lvar.At(1-1)).Id())
                              /*list->list*/v_bag_arg = _CL_obj.Id()
                              } 
                            va_arg2.AddFast(v_bag_arg)} 
                          va_arg1.Args = va_arg2
                          /*list->list*/} 
                        va_arg2 = _CL_obj.Id()
                        } 
                      va_arg1.Test = va_arg2
                      /*any->any*/} 
                    { 
                      var va_arg1 *If  
                      var va_arg2 *ClaireAny  
                      va_arg1 = _CL_obj
                      { var _CL_obj *Do   = To_Do(new(Do).Is(C_Do))
                        _CL_obj.Args = l1.Append(l2)
                        /*list->list*/va_arg2 = _CL_obj.Id()
                        } 
                      va_arg1.Arg = va_arg2
                      /*any->any*/} 
                    arg_4 = To_ComplexInstruction(_CL_obj.Id())
                    } 
                  } else {
                  { var _CL_obj *Let   = To_Let(new(Let).Is(C_Let))
                    _CL_obj.ClaireVar = To_Variable(lvar.At(3-1))
                    /*Variable->Variable*/_CL_obj.Value = F_readCall_relation(R,lvar.At(1-1)).Id()
                    /*any->any*/{ 
                      var va_arg1 *Let  
                      var va_arg2 *ClaireAny  
                      va_arg1 = _CL_obj
                      { var _CL_obj *If   = To_If(new(If).Is(C_If))
                        { 
                          var va_arg1 *If  
                          var va_arg2 *ClaireAny  
                          va_arg1 = _CL_obj
                          { var _CL_obj *Call   = To_Call(new(Call).Is(C_Call))
                            _CL_obj.Selector = ToProperty(Core.C__I_equal.Id())
                            /*property->property*/_CL_obj.Args = MakeConstantList(lvar.At(2-1),lvar.At(3-1))
                            /*list->list*/va_arg2 = _CL_obj.Id()
                            } 
                          va_arg1.Test = va_arg2
                          /*any->any*/} 
                        { 
                          var va_arg1 *If  
                          var va_arg2 *ClaireAny  
                          va_arg1 = _CL_obj
                          { var _CL_obj *Do   = To_Do(new(Do).Is(C_Do))
                            _CL_obj.Args = l1.Append(l2)
                            /*list->list*/va_arg2 = _CL_obj.Id()
                            } 
                          va_arg1.Arg = va_arg2
                          /*any->any*/} 
                        va_arg2 = _CL_obj.Id()
                        } 
                      va_arg1.Arg = va_arg2
                      /*any->any*/} 
                    arg_4 = To_ComplexInstruction(_CL_obj.Id())
                    } 
                  } 
                try_3 = F_lambda_I_list(MakeConstantList(lvar.At(1-1),lvar.At(2-1)),arg_4.Id())
                } 
              /* ERROR PROTECTION INSERTED (va_arg2-Result) */
              if ErrorIn(try_3) {Result = try_3
              } else {
              va_arg2 = ANY(try_3)
              va_arg1.IfWrite = va_arg2
              /*any->any*/Result = va_arg2.ToEID()
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
/* {0} The go function for: eventMethod(p:property) [status=0] */
func F_eventMethod_property (p *ClaireProperty )  { 
  // procedure body with s = void 
{ var m *ClaireMethod   = F_add_method_property(p,MakeConstantList(p.Domain.Id(),p.Range.Id()),ToType(C_void.Id()),0,ToFunction(CNULL))
    { var _Zf *ClaireFunction   = F_make_function_string(F_append_string(p.Name.String_I(),MakeString("_write")))
      m.Formula = ToLambda(p.IfWrite)
      /*lambda->lambda*/Core.F_close_method(m)
      F_set_arity_function(_Zf,2)
      m.Functional = _Zf
      /*function->function*/} 
    } 
  } 

// The EID go function for: eventMethod @ property (throw: false) 
func E_eventMethod_property (p EID) EID { 
  F_eventMethod_property(ToProperty(OBJ(p)) )
  return EVOID} 

// when we compile -> directly call the demon 
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
/* {0} The go function for: jito(self:any) [status=1] */
func F_Language_jito_any (self *ClaireAny ) EID { 
  var Result EID 
  if ((ClEnv.Jito_ask != CTRUE) || 
      (ClEnv.Debug_I >= 0)) { 
    Result = self.ToEID()
    }  else if (self.Isa.IsIn(C_list) == CTRUE) { 
    { var g0301 *ClaireList   = ToList(self)
      _ = g0301
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
          /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
          loop_1 = F_Language_jito_any(x)
          /* ERROR PROTECTION INSERTED (loop_1-Result) */
          if ErrorIn(loop_1) {Result = loop_1
          break
          } else {
          }
          } 
        } 
      } 
    }  else if (self.Isa.IsIn(C_Vardef) == CTRUE) { 
    { var g0302 *Vardef   = To_Vardef(self)
      g0302.Isa = C_Variable
      /*class->class*/Result = EID{g0302.Id(),0}
      } 
    }  else if (self.Isa.IsIn(C_lambda) == CTRUE) { 
    { var g0303 *ClaireLambda   = ToLambda(self)
      /*g_try(v2:"Result",loop:true) */
      Result = F_Language_jito_any(g0303.Body)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = EID{g0303.Id(),0}
      }
      } 
    }  else if (self.Isa.IsIn(C_And) == CTRUE) { 
    { var g0304 *And   = To_And(self)
      _ = g0304
      Result = F_Language_jito_any(g0304.Args.Id())
      } 
    }  else if (self.Isa.IsIn(C_Or) == CTRUE) { 
    { var g0305 *Or   = To_Or(self)
      _ = g0305
      Result = F_Language_jito_any(g0305.Args.Id())
      } 
    }  else if (self.Isa.IsIn(C_Call) == CTRUE) { 
    { var g0306 *Call   = To_Call(self)
      _ = g0306
      /*g_try(v2:"Result",loop:true) */
      Result = g0306.MakeJito()
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = EID{CTRUE.Id(),0}
      }
      } 
    }  else if (self.Isa.IsIn(C_Let) == CTRUE) { 
    { var g0307 *Let   = To_Let(self)
      _ = g0307
      Result = g0307.LetJito()
      } 
    }  else if (self.Isa.IsIn(C_Assign) == CTRUE) { 
    { var g0308 *Assign   = To_Assign(self)
      /*g_try(v2:"Result",loop:true) */
      if (g0308.ClaireVar.Isa.IsIn(C_Variable) != CTRUE) { 
        Result = ToException(Core.C_general_error.Make(MakeString("[101] ~S is not a variable").Id(),MakeConstantList(g0308.ClaireVar).Id())).Close()
        } else {
        Result = EID{CFALSE.Id(),0}
        } 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = F_Language_jito_any(g0308.Arg)
      }
      } 
    }  else if (self.Isa.IsIn(C_Gassign) == CTRUE) { 
    { var g0309 *Gassign   = To_Gassign(self)
      if (g0309.ClaireVar.Range.Contains(g0309.ClaireVar.Value) == CTRUE) { 
        Result = F_Language_jito_any(g0309.Arg)
        } else {
        Result = EID{CFALSE.Id(),0}
        } 
      } 
    }  else if (self.Isa.IsIn(C_Do) == CTRUE) { 
    { var g0310 *Do   = To_Do(self)
      _ = g0310
      Result = F_Language_jito_any(g0310.Args.Id())
      } 
    }  else if (self.Isa.IsIn(C_If) == CTRUE) { 
    { var g0311 *If   = To_If(self)
      /*g_try(v2:"Result",loop:true) */
      Result = F_Language_jito_any(g0311.Arg)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      /*g_try(v2:"Result",loop:true) */
      Result = F_Language_jito_any(g0311.Test)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = F_Language_jito_any(g0311.Other)
      }}
      } 
    }  else if (self.Isa.IsIn(C_Iteration) == CTRUE) { 
    { var g0312 *Iteration   = To_Iteration(self)
      { var v *ClaireVariable   = g0312.ClaireVar
        { var s *ClaireAny   = g0312.SetArg
          { var o_ask *ClaireBoolean  
            { 
              var v_and6 *ClaireBoolean  
              
              if (s.Isa.IsIn(C_Call) == CTRUE) { 
                { var g0313 *Call   = To_Call(s)
                  _ = g0313
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
            
            if (o_ask == CTRUE) { 
              v.Range = ToType(C_integer.Id())
              /*type->type*/
              } 
            /*g_try(v2:"Result",loop:true) */
            Result = F_Language_jito_any(s)
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            /*g_try(v2:"Result",loop:true) */
            Result = F_Language_jito_any(g0312.Arg)
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            if (o_ask == CTRUE) { 
              { 
                var va_arg1 *ClaireVariable  
                var va_arg2 *ClaireType  
                va_arg1 = v
                va_arg2 = ToType(CNULL)
                va_arg1.Range = va_arg2
                /*type->type*/Result = EID{va_arg2.Id(),0}
                } 
              } else {
              Result = EID{CFALSE.Id(),0}
              } 
            }}
            } 
          } 
        } 
      } 
    }  else if (self.Isa.IsIn(C_Construct) == CTRUE) { 
    { var g0314 *Construct   = To_Construct(self)
      _ = g0314
      
      Result = F_Language_jito_any(g0314.Args.Id())
      } 
    }  else if (self.Isa.IsIn(C_Exists) == CTRUE) { 
    { var g0315 *Exists   = To_Exists(self)
      /*g_try(v2:"Result",loop:true) */
      Result = F_Language_jito_any(g0315.SetArg)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      /*g_try(v2:"Result",loop:true) */
      Result = F_Language_jito_any(g0315.Arg)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = F_Language_jito_any(g0315.Other)
      }}
      } 
    }  else if (self.Isa.IsIn(C_Handle) == CTRUE) { 
    { var g0316 *ClaireHandle   = To_ClaireHandle(self)
      /*g_try(v2:"Result",loop:true) */
      if (C_class.Id() != g0316.Test.Isa.Id()) { 
        Result = ToException(Core.C_general_error.Make(MakeString("syntax: [try %S] must use a class").Id(),MakeConstantList(g0316.Test).Id())).Close()
        } else {
        Result = EID{CFALSE.Id(),0}
        } 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      /*g_try(v2:"Result",loop:true) */
      Result = F_Language_jito_any(g0316.Arg)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = F_Language_jito_any(g0316.Other)
      }}
      } 
    }  else if (self.Isa.IsIn(C_Definition) == CTRUE) { 
    { var g0317 *Definition   = To_Definition(self)
      if (F_Language_fast_definition_ask_class(g0317.Arg) == CTRUE) { 
        { 
          var va_arg1 *ClaireAny  
          var va_arg2 *ClaireClass  
          va_arg1 = g0317.Id()
          va_arg2 = C_Language_DefFast
          va_arg1.Isa = va_arg2
          /*class->class*/Result = EID{va_arg2.Id(),0}
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

// claire/VARIANT:boolean := true         // debug to remove, replace by jito?()
// Let is special in CLAIRE4 : we implement the implicit typing found in the compiler = to infer
// the type  from the value (when no range is given)
// Note : this is doubtful ... 
/* {0} The go function for: letJito(self:Let) [status=1] */
func (self *Let ) LetJito () EID { 
  var Result EID 
  { var v *ClaireVariable   = self.ClaireVar
    { var x *ClaireAny   = self.Value
      { var untyped *ClaireBoolean   = MakeBoolean((v.Range.Id() == CNULL))
        
        /*g_try(v2:"Result",loop:true) */
        if (untyped == CTRUE) { 
          /*g_try(v2:"Result",loop:true) */
          if (x.Isa.IsIn(C_List) == CTRUE) { 
            { var t *ClaireType   = ToType(OBJ(Core.F_CALL(C_of,ARGS(x.ToEID()))))
              if (Equal(t.Id(),CEMPTY.Id()) != CTRUE) { 
                { 
                  var va_arg1 *ClaireVariable  
                  var va_arg2 *ClaireType  
                  va_arg1 = v
                  va_arg2 = Core.F_param_I_class(C_list,t)
                  va_arg1.Range = va_arg2
                  /*type->type*/Result = EID{va_arg2.Id(),0}
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
              /*g_try(v2:"try_1",loop:false) */
              try_1 = F_static_type_any(x)
              /* ERROR PROTECTION INSERTED (va_arg2-Result) */
              if ErrorIn(try_1) {Result = try_1
              } else {
              va_arg2 = ToType(OBJ(try_1))
              va_arg1.Range = va_arg2
              /*type->type*/Result = EID{va_arg2.Id(),0}
              }
              } 
            } 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          Result = EID{CFALSE.Id(),0}
          }
          } else {
          Result = EID{CFALSE.Id(),0}
          } 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        /*g_try(v2:"Result",loop:true) */
        Result = F_Language_jito_any(x)
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        /*g_try(v2:"Result",loop:true) */
        Result = F_Language_jito_any(self.Arg)
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        if (untyped == CTRUE) { 
          { 
            var va_arg1 *ClaireVariable  
            var va_arg2 *ClaireType  
            va_arg1 = v
            va_arg2 = ToType(CNULL)
            va_arg1.Range = va_arg2
            /*type->type*/Result = EID{va_arg2.Id(),0}
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
/* {0} The go function for: makeJito(self:Call) [status=1] */
func (self *Call ) MakeJito () EID { 
  var Result EID 
  /*g_try(v2:"Result",loop:true) */
  Result = F_Language_jito_any(self.Args.Id())
  /* ERROR PROTECTION INSERTED (Result-Result) */
  if !ErrorIn(Result) {
  { var p *ClaireProperty   = self.Selector
    { var larg *ClaireList   = self.Args
      { var n int  = larg.Length()
        { var m *ClaireAny   = CNULL
          var g0323I *ClaireBoolean  
          { 
            var v_and5 *ClaireBoolean  
            
            v_and5 = Equal(p.Id(),Core.C_write.Id())
            if (v_and5 == CFALSE) {g0323I = CFALSE
            } else { 
              { var p2 *ClaireAny   = self.Args.At(1-1)
                if (p2.Isa.IsIn(C_property) == CTRUE) { 
                  { var g0319 *ClaireProperty   = ToProperty(p2)
                    v_and5 = MakeBoolean((g0319.Inverse.Id() == CNULL) && (g0319.Store_ask != CTRUE) && (g0319.IfWrite == CNULL))
                    } 
                  } else {
                  v_and5 = CFALSE
                  } 
                } 
              if (v_and5 == CFALSE) {g0323I = CFALSE
              } else { 
                g0323I = CTRUE} 
              } 
            } 
          if (g0323I == CTRUE) { 
            p = C_write_fast
            self.Selector = C_write_fast
            /*property->property*/} 
          /*g_try(v2:"Result",loop:true) */
          var g0324I *ClaireBoolean  
          { 
            var v_and5 *ClaireBoolean  
            
            v_and5 = Core.F__inf_equal_integer(p.Open,1)
            if (v_and5 == CFALSE) {g0324I = CFALSE
            } else { 
              v_and5 = Core.F__inf_equal_integer(p.Restrictions.Length(),12)
              if (v_and5 == CFALSE) {g0324I = CFALSE
              } else { 
                { var arg_1 *ClaireAny  
                  _ = arg_1
                  { 
                    var x *ClaireRestriction  
                    _ = x
                    var x_iter *ClaireAny  
                    arg_1= CFALSE.Id()
                    for _,x_iter = range(p.Restrictions.ValuesO()){ 
                      x = ToRestriction(x_iter)
                      var g0325I *ClaireBoolean  
                      { var arg_2 *ClaireBoolean  
                        _ = arg_2
                        { var arg_3 *ClaireAny  
                          _ = arg_3
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
                        g0325I = arg_2.Not
                        } 
                      if (g0325I == CTRUE) { 
                        arg_1 = CTRUE.Id()
                        break
                        } 
                      } 
                    } 
                  v_and5 = Core.F_not_any(arg_1)
                  } 
                if (v_and5 == CFALSE) {g0324I = CFALSE
                } else { 
                  g0324I = CTRUE} 
                } 
              } 
            } 
          if (g0324I == CTRUE) { 
            { var lt *ClaireList  
              _ = lt
              var try_4 EID 
              /*g_try(v2:"try_4",loop:false) */
              { 
                var v_list7 *ClaireList  
                var x *ClaireAny  
                var v_local7 *ClaireAny  
                v_list7 = larg
                try_4 = EID{CreateList(ToType(CEMPTY.Id()),v_list7.Length()).Id(),0}
                for CLcount := 0; CLcount < v_list7.Length(); CLcount++{ 
                  x = v_list7.At(CLcount)
                  var try_5 EID 
                  /*g_try(v2:"try_5",loop:tuple("try_4", EID)) */
                  try_5 = F_static_type_any(x)
                  /* ERROR PROTECTION INSERTED (v_local7-try_4) */
                  if ErrorIn(try_5) {try_4 = try_5
                  break
                  } else {
                  v_local7 = ANY(try_5)
                  ToList(OBJ(try_4)).PutAt(CLcount,v_local7)
                  } 
                }
                } 
              /* ERROR PROTECTION INSERTED (lt-Result) */
              if ErrorIn(try_4) {Result = try_4
              } else {
              lt = ToList(OBJ(try_4))
              
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
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          var g0326I *ClaireBoolean  
          if (C_method.Id() == m.Isa.Id()) { 
            { var g0321 *ClaireMethod   = ToMethod(m)
              _ = g0321
              g0326I = MakeBoolean((g0321.Functional.Id() == CNULL)).Not
              } 
            } else {
            g0326I = CFALSE
            } 
          if (g0326I == CTRUE) { 
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
              /*class->class*/} 
            { 
              var va_arg1 *CallMethod  
              var va_arg2 *ClaireMethod  
              va_arg1 = To_CallMethod(self.Id())
              va_arg2 = ToMethod(m)
              va_arg1.Arg = va_arg2
              /*method->method*/Result = EID{va_arg2.Id(),0}
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
/* {0} The go function for: makeCallMatch(x:restriction,lt:list) [status=0] */
func F_Language_makeCallMatch_restriction (x *ClaireRestriction ,lt *ClaireList ) *ClaireBoolean  { 
  // procedure body with s = boolean 
var Result *ClaireBoolean  
  { var n int  = lt.Length()
    { var ld *ClaireList   = x.Domain
      { 
        var v_and3 *ClaireBoolean  
        
        v_and3 = Equal(MakeInteger(ld.Length()).Id(),MakeInteger(n).Id())
        if (v_and3 == CFALSE) {Result = CFALSE
        } else { 
          { var arg_1 *ClaireAny  
            _ = arg_1
            { var i int  = 1
              { var g0327 int  = n
                _ = g0327
                arg_1= CFALSE.Id()
                for (i <= g0327) { 
                  /* While stat, v:"arg_1" loop:false */
                  if (ToType(lt.At(i-1)).Included(ToType(ld.ValuesO()[i-1])) != CTRUE) { 
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
    } 
  return Result} 

// The EID go function for: makeCallMatch @ restriction (throw: false) 
func E_Language_makeCallMatch_restriction (x EID,lt EID) EID { 
  return EID{F_Language_makeCallMatch_restriction(ToRestriction(OBJ(x)),ToList(OBJ(lt)) ).Id(),0}} 

// close some classes : final => no subclasses,  default() => ephemeral
// CLAIRE 4 : make sure that open statement for class are all here
// instuctions are ephemeral