/***** CLAIRE Compilation of file /Users/ycaseau/Dropbox/src/clairev4.03/src/meta/pretty.cl 
         [version 4.0.03 / safety 5] Monday 12-27-2021 10:35:23 *****/

package Language
import (_ "fmt"
	. "Kernel"
	"Core"
)

//-------- dumb function to prevent import errors --------
func import_g0000() { 
    _ = Core.It
    } 
  
  
//+-------------------------------------------------------------+
//| CLAIRE                                                      |
//| pretty.cl                                                   |
//| Copyright (C) 1994 - 2021 Yves Caseau. All Rights Reserved  |
//| cf. copyright info in file object.cl: about()               |
//+-------------------------------------------------------------+
// *********************************************************************
// *  Contents:                                                        *
// *  Part 1: unbound_symbol and variables                             *
// *  Part 2: lambdas                                                  *
// *  Part 3: close methods for lattice_set instantiation              *
// *  Part 4: Pretty printing                                          *
// *********************************************************************
// Instruction <: system_object()  : in Kernel (CLAIRE 4)
/* {1} The go function for: no_eval(self:Instruction) [status=1] */
func F_no_eval_Instruction (self *ClaireInstruction ) EID { 
    var Result EID 
    Result = ToException(Core.C_general_error.Make(MakeString("[144] evaluate(~S) is not defined").Id(),MakeConstantList(self.Id().Isa.Id()).Id())).Close()
    return Result} 
  
// The EID go function for: no_eval @ Instruction (throw: true) 
func E_no_eval_Instruction (self EID) EID { 
    return F_no_eval_Instruction(To_Instruction(OBJ(self)) )} 
  
// import => cannot work in CLAIRE4
// *********************************************************************
// *   Part 1: unbound_symbol and variables                            *
// *********************************************************************
// An unbound_symbol is created by the reader when a symbol is not bound
//
//unbound_symbol <: Basic_instruction(identifier:symbol)
/* {1} The go function for: self_print(self:unbound_symbol) [status=0] */
func F_self_print_unbound_symbol_Language (self *ClaireUnboundSymbol )  { 
    // procedure body with s = void 
self.Name.Princ()
    PRINC("")
    } 
  
// The EID go function for: self_print @ unbound_symbol (throw: false) 
func E_self_print_unbound_symbol_Language (self EID) EID { 
    F_self_print_unbound_symbol_Language(ToUnboundSymbol(OBJ(self)) )
    return EVOID} 
  
/* {1} The go function for: self_eval(self:unbound_symbol) [status=1] */
func F_self_eval_unbound_symbol (self *ClaireUnboundSymbol ) EID { 
    var Result EID 
    if (Core.F_owner_any(self.Name.Get()).IsIn(C_thing) == CTRUE) { 
      Result = EVAL(self.Name.Get())
      } else {
      Result = ToException(Core.C_general_error.Make(MakeString("[145] the symbol ~A is unbound").Id(),MakeConstantList(self.Name.Id()).Id())).Close()
      } 
    return Result} 
  
// The EID go function for: self_eval @ unbound_symbol (throw: true) 
func E_self_eval_unbound_symbol (self EID) EID { 
    return F_self_eval_unbound_symbol(ToUnboundSymbol(OBJ(self)) )} 
  
// The EVAL go function for: unbound_symbol 
func EVAL_unbound_symbol (x *ClaireAny) EID { 
     return F_self_eval_unbound_symbol(ToUnboundSymbol(x))} 
  
// A lexical variable is defined by a "Let" or inside a method's definition
// Lexical variables --------------------------------------------------
//
//
/* {1} The go function for: self_print(self:Variable) [status=0] */
func F_self_print_Variable_Language (self *ClaireVariable )  { 
    // procedure body with s = void 
{ var s *ClaireSymbol   = self.Pname
      if (s.Id() == CNULL) { 
        PRINC("V?")
        } else {
        s.Princ()
        } 
      } 
    } 
  
// The EID go function for: self_print @ Variable (throw: false) 
func E_self_print_Variable_Language (self EID) EID { 
    F_self_print_Variable_Language(To_Variable(OBJ(self)) )
    return EVOID} 
  
/* {1} The go function for: ppvariable(self:Variable) [status=1] */
func F_ppvariable_Variable (self *ClaireVariable ) EID { 
    var Result EID 
    if (self.Range.Id() != CNULL) { 
      self.Pname.Princ()
      PRINC(":")
      /*g_try(v2:"Result",loop:true) */
      Result = F_printexp_any(self.Range.Id(),CFALSE)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC("")
      Result = EVOID
      }
      } else {
      self.Pname.Princ()
      Result = EVOID
      } 
    return Result} 
  
// The EID go function for: ppvariable @ Variable (throw: true) 
func E_ppvariable_Variable (self EID) EID { 
    return F_ppvariable_Variable(To_Variable(OBJ(self)) )} 
  
/* {1} The go function for: ppvariable(self:list) [status=1] */
func F_ppvariable_list (self *ClaireList ) EID { 
    var Result EID 
    { var f *ClaireBoolean   = CTRUE
      _ = f
      { 
        var v *ClaireAny  
        _ = v
        Result= EID{CFALSE.Id(),0}
        var v_support *ClaireList  
        v_support = self
        v_len := v_support.Length()
        for i_it := 0; i_it < v_len; i_it++ { 
          v = v_support.At(i_it)
          var loop_1 EID 
          _ = loop_1
          { 
          if (f == CTRUE) { 
            f = CFALSE
            } else {
            PRINC(",")
            } 
          /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
          if (v.Isa.IsIn(C_Variable) == CTRUE) { 
            { var g0001 *ClaireVariable   = To_Variable(v)
              _ = g0001
              loop_1 = F_ppvariable_Variable(g0001)
              } 
            } else {
            loop_1 = Core.F_CALL(C_print,ARGS(v.ToEID()))
            } 
          /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
          if ErrorIn(loop_1) {Result = loop_1
          break
          } else {
          }
          }
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: ppvariable @ list (throw: true) 
func E_ppvariable_list (self EID) EID { 
    return F_ppvariable_list(ToList(OBJ(self)) )} 
  
/* {1} The go function for: self_eval(self:Variable) [status=0] */
func F_self_eval_Variable (self *ClaireVariable ) EID { 
    var Result EID 
    Result = ClEnv.EvalStack[(ClEnv.Base+self.Index)]
    return Result} 
  
// The EID go function for: self_eval @ Variable (throw: true) 
func E_self_eval_Variable (self EID) EID { 
    return F_self_eval_Variable(To_Variable(OBJ(self)) )} 
  
/* {1} The go function for: write_value(self:Variable,val:any) [status=1] */
func F_write_value_Variable (self *ClaireVariable ,val *ClaireAny ) EID { 
    var Result EID 
    if ((self.Range.Id() == CNULL) || 
        (self.Range.Contains(val) == CTRUE)) { 
      ClEnv.EvalStack[(ClEnv.Base+self.Index)]=val.ToEID()
      Result = val.ToEID()
      } else {
      { var _CL_obj *Core.RangeError   = Core.ToRangeError(new(Core.RangeError).Is(Core.C_range_error))
        _CL_obj.Arg = self.Id()
        /*any->any*/_CL_obj.Cause = val
        /*any->any*//*g_try(v2:"Result",loop:true) */
        Result = Core.F_write_property(C_Language_wrong,ToObject(_CL_obj.Id()),self.Range.Id())
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        Result = _CL_obj.Close()
        }
        } 
      } 
    return Result} 
  
// The EID go function for: write_value @ Variable (throw: true) 
func E_write_value_Variable (self EID,val EID) EID { 
    return F_write_value_Variable(To_Variable(OBJ(self)),ANY(val) )} 
  
// this is the definition of a typed variable / Vardef is a syntactic marker
// in CLAIRE 4, Vardef are transformed in Var at run time
// this is strange and should be fixed  or understood
/* {1} The go function for: self_eval(self:Vardef) [status=1] */
func (self *Vardef ) SelfEval () EID { 
    var Result EID 
    { var i *ClaireAny   = MakeInteger(self.Index).Id()
      if (ToBoolean(OBJ(Core.F_CALL(ToProperty(C__sup_equal.Id()),ARGS(i.ToEID(),EID{C__INT,IVAL(0)})))) == CTRUE) { 
        Result = ClEnv.EvalStack[INT(Core.F_CALL(ToProperty(Core.C__plus.Id()),ARGS(EID{C__INT,IVAL(ClEnv.Base)},i.ToEID())))]
        } else {
        Result = ToException(Core.C_general_error.Make(MakeString("[146] The variable ~S is not defined").Id(),MakeConstantList(self.Id()).Id())).Close()
        } 
      } 
    return Result} 
  
// The EID go function for: self_eval @ Vardef (throw: true) 
func E_self_eval_Vardef (self EID) EID { 
    return To_Vardef(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Vardef 
func EVAL_Vardef (x *ClaireAny) EID { 
     return To_Vardef(x).SelfEval()} 
  
//   [self_print(self:Vardef) : any -> ppvariable(self) ]
// global_variables are defined in exception ? ---------------------------
// a global variable is a named object with a special evaluation
//
// self_eval(self:global_variable) : any -> self.value  -> moved to object.cl
/* {1} The go function for: write_value(self:global_variable,val:any) [status=1] */
func F_write_value_global_variable (self *Core.GlobalVariable ,val *ClaireAny ) EID { 
    var Result EID 
    if (self.Range.Contains(val) == CTRUE) { 
      /*g_try(v2:"Result",loop:true) */
      Result = Core.F_put_store_property2(C_value,ToObject(self.Id()),val,self.Store_ask)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = val.ToEID()
      }
      } else {
      { var _CL_obj *Core.RangeError   = Core.ToRangeError(new(Core.RangeError).Is(Core.C_range_error))
        _CL_obj.Cause = self.Id()
        /*any->any*/_CL_obj.Arg = val
        /*any->any*//*g_try(v2:"Result",loop:true) */
        Result = Core.F_write_property(C_Language_wrong,ToObject(_CL_obj.Id()),self.Range.Id())
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        Result = _CL_obj.Close()
        }
        } 
      } 
    return Result} 
  
// The EID go function for: write_value @ global_variable (throw: true) 
func E_write_value_global_variable (self EID,val EID) EID { 
    return F_write_value_global_variable(Core.ToGlobalVariable(OBJ(self)),ANY(val) )} 
  
// v0.01
// same as C  (used externC("((int) EOF",integer))
// v3.2.52
// v3.4
// *********************************************************************
// *   Part 2: CLAIRE Lambdas                                           *
// *********************************************************************
// CLAIRE lambdas are the basic functional objects, defined by a filter
// and a piece of code. Lambda is defined in the "method" file.
// applying a lambda to a list of arguments
//
/* {1} The go function for: apply(self:lambda,%l:list) [status=1] */
func F_apply_lambda (self *ClaireLambda ,_Zl *ClaireList ) EID { 
    var Result EID 
    { var start int  = ClEnv.Index
      { var retour int  = ClEnv.Base
        _ = retour
        ClEnv.Base= start
        { 
          var _Zx *ClaireAny  
          _ = _Zx
          var _Zx_support *ClaireList  
          _Zx_support = _Zl
          _Zx_len := _Zx_support.Length()
          for i_it := 0; i_it < _Zx_len; i_it++ { 
            _Zx = _Zx_support.At(i_it)
            ClEnv.Push(_Zx.ToEID())
            } 
          } 
        F_stack_add(self.Dimension)
        { 
          var val EID 
          /*g_try(v2:"val",loop:false) */
          val = EVAL(self.Body)
          /* ERROR PROTECTION INSERTED (val-Result) */
          if ErrorIn(val) {Result = val
          } else {
          ClEnv.Base= retour
          ClEnv.Index= start
          Result = val}
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: apply @ lambda (throw: true) 
func E_apply_lambda (self EID,_Zl EID) EID { 
    return F_apply_lambda(ToLambda(OBJ(self)),ToList(OBJ(_Zl)) )} 
  
/* {1} The go function for: call(self:lambda,l:listargs) [status=1] */
func F_call_lambda2 (self *ClaireLambda ,l *ClaireList ) EID { 
    var Result EID 
    Result = F_apply_lambda(self,ToList(l.Id()))
    return Result} 
  
// The EID go function for: call @ lambda (throw: true) 
func E_call_lambda2 (self EID,l EID) EID { 
    return F_call_lambda2(ToLambda(OBJ(self)),ToList(OBJ(l)) )} 
  
// printing a lambda
/* {1} The go function for: self_print(self:lambda) [status=1] */
func F_self_print_lambda_Language (self *ClaireLambda ) EID { 
    var Result EID 
    PRINC("lambda[(")
    /*g_try(v2:"Result",loop:true) */
    Result = F_ppvariable_list(self.Vars)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("),")
    /*g_try(v2:"Result",loop:true) */
    Result = F_lbreak_integer(1)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    /*g_try(v2:"Result",loop:true) */
    Result = Core.F_CALL(C_print,ARGS(self.Body.ToEID()))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    Core.C_pretty.Index = (Core.C_pretty.Index-1)
    /*integer->integer*/PRINC("]")
    Result = EVOID
    }}}
    return Result} 
  
// The EID go function for: self_print @ lambda (throw: true) 
func E_self_print_lambda_Language (self EID) EID { 
    return F_self_print_lambda_Language(ToLambda(OBJ(self)) )} 
  
// map is the most famous function on a lambda
/* {1} The go function for: map(self:lambda,%l:bag) [status=1] */
func F_map_lambda (self *ClaireLambda ,_Zl *ClaireBag ) EID { 
    var Result EID 
    if (C_set.Id() == _Zl.Isa.Id()) { 
      { var g0004 *ClaireSet   = ToSet(_Zl.Id())
        _ = g0004
        { var x_bag *ClaireSet   = ToType(CEMPTY.Id()).EmptySet()
          /*g_try(v2:"Result",loop:true) */
          { 
            var x *ClaireAny  
            _ = x
            Result= EID{CFALSE.Id(),0}
            var x_support *ClaireSet  
            x_support = g0004
            for i_it := 0; i_it < x_support.Count; i_it++ { 
              x = x_support.At(i_it)
              var loop_1 EID 
              _ = loop_1
              /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
              { var arg_2 *ClaireAny  
                _ = arg_2
                var try_3 EID 
                /*g_try(v2:"try_3",loop:false) */
                try_3 = Core.F_funcall_lambda1(self,x)
                /* ERROR PROTECTION INSERTED (arg_2-loop_1) */
                if ErrorIn(try_3) {loop_1 = try_3
                } else {
                arg_2 = ANY(try_3)
                loop_1 = EID{x_bag.AddFast(arg_2).Id(),0}/*t=any,s=EID*/
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
          Result = EID{x_bag.Id(),0}
          }
          } 
        } 
      } else {
      { 
        var v_list3 *ClaireList  
        var x *ClaireAny  
        var v_local3 *ClaireAny  
        v_list3 = ToList(_Zl.Id())
        Result = EID{CreateList(ToType(CEMPTY.Id()),v_list3.Length()).Id(),0}
        for CLcount := 0; CLcount < v_list3.Length(); CLcount++{ 
          x = v_list3.At(CLcount)
          var try_4 EID 
          /*g_try(v2:"try_4",loop:tuple("Result", EID)) */
          try_4 = Core.F_funcall_lambda1(self,x)
          /* ERROR PROTECTION INSERTED (v_local3-Result) */
          if ErrorIn(try_4) {Result = try_4
          break
          } else {
          v_local3 = ANY(try_4)
          ToList(OBJ(Result)).PutAt(CLcount,v_local3)
          } 
        }
        } 
      } 
    return Result} 
  
// The EID go function for: map @ lambda (throw: true) 
func E_map_lambda (self EID,_Zl EID) EID { 
    return F_map_lambda(ToLambda(OBJ(self)),ToBag(OBJ(_Zl)) )} 
  
// lambda! and flexical_build communicate via a global_variable, which
// however is only used in this file (and also by cfile :-) ):
//
// creating a lambda from an instruction and a list of variables
/* {1} The go function for: iClaire/lambda!(lvar:list,self:any) [status=1] */
func F_lambda_I_list (lvar *ClaireList ,self *ClaireAny ) EID { 
    var Result EID 
    C__starvariable_index_star.Value = MakeInteger(0).Id()
    { 
      var v *ClaireVariable  
      _ = v
      var v_iter *ClaireAny  
      var v_support *ClaireList  
      v_support = lvar
      v_len := v_support.Length()
      for i_it := 0; i_it < v_len; i_it++ { 
        v_iter = v_support.At(i_it)
        v = To_Variable(v_iter)
        v.Index = ToInteger(C__starvariable_index_star.Value).Value
        /*integer->integer*/v.Isa = C_Variable
        /*class->class*/C__starvariable_index_star.Value = MakeInteger((ToInteger(C__starvariable_index_star.Value).Value+1)).Id()
        } 
      } 
    { var corps *ClaireAny  
      _ = corps
      var try_1 EID 
      /*g_try(v2:"try_1",loop:false) */
      try_1 = F_lexical_build_any(self,lvar,ToInteger(C__starvariable_index_star.Value).Value)
      /* ERROR PROTECTION INSERTED (corps-Result) */
      if ErrorIn(try_1) {Result = try_1
      } else {
      corps = ANY(try_1)
      { var resultat *ClaireLambda   = ToLambda(new(ClaireLambda).Is(C_lambda))
        resultat.Vars = lvar
        /*list->list*/resultat.Body = corps
        /*any->any*/resultat.Dimension = ToInteger(C__starvariable_index_star.Value).Value
        /*integer->integer*/Result = EID{resultat.Id(),0}
        } 
      }
      } 
    return Result} 
  
// The EID go function for: iClaire/lambda! @ list (throw: true) 
func E_lambda_I_list (lvar EID,self EID) EID { 
    return F_lambda_I_list(ToList(OBJ(lvar)),ANY(self) )} 
  
// Give to each lexical variable its right position in the stack.
// We look for a named object or an unbound symbol to replace by a lexical
// variable.
// The number of variables is kept in the global_variable *variable_index*.
// On entry, n need not be equal to size(lvar) (see [case ...instruction]).
//
/* {1} The go function for: iClaire/lexical_build(self:any,lvar:list,n:integer) [status=1] */
func F_lexical_build_any (self *ClaireAny ,lvar *ClaireList ,n int) EID { 
    var Result EID 
    if ((self.Isa.IsIn(C_thing) == CTRUE) || 
        (self.Isa.IsIn(C_unbound_symbol) == CTRUE)) { 
      Result = F_lexical_change_any(self,lvar)
      } else {
      /*g_try(v2:"Result",loop:true) */
      if (self.Isa.IsIn(C_Variable) == CTRUE) { 
        { var g0006 *ClaireVariable   = To_Variable(self)
          /*g_try(v2:"Result",loop:true) */
          if (Equal(MakeInteger(g0006.Index).Id(),CNULL) == CTRUE) { 
            Result = ToException(Core.C_general_error.Make(MakeString("[145] the symbol ~A is unbound").Id(),MakeConstantList(g0006.Pname.Id()).Id())).Close()
            } else {
            Result = EID{CFALSE.Id(),0}
            } 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          Result = EID{g0006.Id(),0}
          }
          } 
        }  else if (self.Isa.IsIn(C_Call) == CTRUE) { 
        { var g0007 *Call   = To_Call(self)
          { var s *ClaireAny  
            var try_1 EID 
            /*g_try(v2:"try_1",loop:false) */
            try_1 = F_lexical_change_any(g0007.Selector.Id(),lvar)
            /* ERROR PROTECTION INSERTED (s-Result) */
            if ErrorIn(try_1) {Result = try_1
            } else {
            s = ANY(try_1)
            /*g_try(v2:"Result",loop:true) */
            Result = F_lexical_build_any(g0007.Args.Id(),lvar,n)
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            if (g0007.Selector.Id() != s) { 
              g0007.Selector = Core.C_call
              /*property->property*/{ 
                var va_arg1 *Call  
                var va_arg2 *ClaireList  
                va_arg1 = g0007
                va_arg2 = F_cons_any(s,g0007.Args)
                va_arg1.Args = va_arg2
                /*list->list*/Result = EID{va_arg2.Id(),0}
                } 
              } else {
              Result = EID{CFALSE.Id(),0}
              } 
            }
            }
            } 
          } 
        }  else if (self.Isa.IsIn(C_Instruction) == CTRUE) { 
        { var g0008 *ClaireInstruction   = To_Instruction(self)
          { var _Ztype *ClaireClass   = g0008.Isa
            /*g_try(v2:"Result",loop:true) */
            if (C_Instruction_with_var.Descendents.Contain_ask(_Ztype.Id()) == CTRUE) { 
              /*g_try(v2:"Result",loop:true) */
              Result = Core.F_put_property2(C_mClaire_index,ToObject(OBJ(Core.F_CALL(C_var,ARGS(EID{g0008.Id(),0})))),MakeInteger(n).Id())
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              n = (n+1)
              if (n > ToInteger(C__starvariable_index_star.Value).Value) { 
                var v_gassign2 *ClaireAny  
                v_gassign2 = MakeInteger(n).Id()
                C__starvariable_index_star.Value = v_gassign2
                Result = v_gassign2.ToEID()
                } else {
                Result = EID{CFALSE.Id(),0}
                } 
              }
              } else {
              Result = EID{CFALSE.Id(),0}
              } 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            { 
              var s *ClaireSlot  
              _ = s
              var s_iter *ClaireAny  
              Result= EID{CFALSE.Id(),0}
              for _,s_iter = range(_Ztype.Slots.ValuesO()){ 
                s = ToSlot(s_iter)
                var loop_3 EID 
                _ = loop_3
                /*g_try(v2:"loop_3",loop:tuple("Result", EID)) */
                { var x *ClaireAny   = Core.F_get_slot(s,ToObject(g0008.Id()))
                  if (((x.Isa.IsIn(C_thing) == CTRUE) || 
                        (x.Isa.IsIn(C_unbound_symbol) == CTRUE)) && 
                      (s.Range.Id() == C_any.Id())) { 
                    { var arg_4 *ClaireAny  
                      _ = arg_4
                      var try_5 EID 
                      /*g_try(v2:"try_5",loop:false) */
                      try_5 = F_lexical_change_any(x,lvar)
                      /* ERROR PROTECTION INSERTED (arg_4-loop_3) */
                      if ErrorIn(try_5) {loop_3 = try_5
                      } else {
                      arg_4 = ANY(try_5)
                      loop_3 = Core.F_put_slot(s,ToObject(g0008.Id()),arg_4).ToEID()
                      }
                      } 
                    } else {
                    loop_3 = F_lexical_build_any(x,lvar,n)
                    } 
                  } 
                /* ERROR PROTECTION INSERTED (loop_3-Result) */
                if ErrorIn(loop_3) {Result = loop_3
                break
                } else {
                }
                } 
              } 
            }
            } 
          } 
        }  else if (self.Isa.IsIn(C_list) == CTRUE) { 
        { var g0009 *ClaireList   = ToList(self)
          { var _Zn int  = g0009.Length()
            Result= EID{CFALSE.Id(),0}
            for (_Zn > 0) { 
              /* While stat, v:"Result" loop:true */
              var loop_6 EID 
              _ = loop_6
              { 
              /*g_try(v2:"loop_6",loop:tuple("Result", EID)) */
              { var x *ClaireAny   = g0009.At(_Zn-1)
                if ((x.Isa.IsIn(C_thing) == CTRUE) || 
                    (x.Isa.IsIn(C_unbound_symbol) == CTRUE)) { 
                  { var arg_7 *ClaireAny  
                    _ = arg_7
                    var try_8 EID 
                    /*g_try(v2:"try_8",loop:false) */
                    try_8 = F_lexical_change_any(x,lvar)
                    /* ERROR PROTECTION INSERTED (arg_7-loop_6) */
                    if ErrorIn(try_8) {loop_6 = try_8
                    } else {
                    arg_7 = ANY(try_8)
                    loop_6 = ToArray(g0009.Id()).NthPut(_Zn,arg_7).ToEID()
                    }
                    } 
                  } else {
                  loop_6 = F_lexical_build_any(x,lvar,n)
                  } 
                } 
              /* ERROR PROTECTION INSERTED (loop_6-loop_6) */
              if ErrorIn(loop_6) {Result = loop_6
              break
              } else {
              _Zn = (_Zn-1)
              }
              /* try?:false, v2:"v_while6" loop will be:tuple("Result", EID) */
              } 
            }
            } 
          } 
        } else {
        Result = EID{CNIL.Id(),0}
        } 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = self.ToEID()
      }
      } 
    return Result} 
  
// The EID go function for: iClaire/lexical_build @ any (throw: true) 
func E_lexical_build_any (self EID,lvar EID,n EID) EID { 
    return F_lexical_build_any(ANY(self),ToList(OBJ(lvar)),INT(n) )} 
  
/* {1} The go function for: iClaire/lexical_change(self:any,lvar:list) [status=1] */
func F_lexical_change_any (self *ClaireAny ,lvar *ClaireList ) EID { 
    var Result EID 
    { var rep *ClaireAny   = self
      _ = rep
      { var _Zname *ClaireSymbol  
        _ = _Zname
        var try_1 EID 
        /*g_try(v2:"try_1",loop:false) */
        if (self.Isa.IsIn(C_Variable) == CTRUE) { 
          { var g0011 *ClaireVariable   = To_Variable(self)
            _ = g0011
            try_1 = EID{g0011.Pname.Id(),0}
            } 
          } else {
          try_1 = F_extract_symbol_any(self)
          } 
        /* ERROR PROTECTION INSERTED (_Zname-Result) */
        if ErrorIn(try_1) {Result = try_1
        } else {
        _Zname = ToSymbol(OBJ(try_1))
        { 
          var x *ClaireVariable  
          _ = x
          var x_iter *ClaireAny  
          var x_support *ClaireList  
          x_support = lvar
          x_len := x_support.Length()
          for i_it := 0; i_it < x_len; i_it++ { 
            x_iter = x_support.At(i_it)
            x = To_Variable(x_iter)
            if (x.Pname.Id() == _Zname.Id()) { 
              rep = x.Id()
              } 
            } 
          } 
        Result = rep.ToEID()
        }
        } 
      } 
    return Result} 
  
// The EID go function for: iClaire/lexical_change @ any (throw: true) 
func E_lexical_change_any (self EID,lvar EID) EID { 
    return F_lexical_change_any(ANY(self),ToList(OBJ(lvar)) )} 
  
// *******************************************************************
// *       Part 3: functions for lattice_set instantiation           *
// *******************************************************************
// close is the basic method called by an instantiation.
// Once the indexed list is built, we never call it again.
//
/* {1} The go function for: close(self:class) [status=0] */
func F_close_class (self *ClaireClass ) *ClaireClass  { 
    return  self
    } 
  
// The EID go function for: close @ class (throw: false) 
func E_close_class (self EID) EID { 
    return EID{F_close_class(ToClass(OBJ(self)) ).Id(),0}} 
  
// Extract the symbol associated with self.
// This is useful e.g. when using read() (read@port, read@string).
//
/* {1} The go function for: iClaire/extract_symbol(self:any) [status=1] */
func F_extract_symbol_any (self *ClaireAny ) EID { 
    var Result EID 
    if (self.Isa.IsIn(C_unbound_symbol) == CTRUE) { 
      { var g0013 *ClaireUnboundSymbol   = ToUnboundSymbol(self)
        _ = g0013
        Result = EID{g0013.Name.Id(),0}
        } 
      }  else if (self.Isa.IsIn(C_thing) == CTRUE) { 
      { var g0014 *ClaireThing   = ToThing(self)
        _ = g0014
        Result = EID{g0014.Name.Id(),0}
        } 
      }  else if (C_class.Id() == self.Isa.Id()) { 
      { var g0015 *ClaireClass   = ToClass(self)
        _ = g0015
        Result = EID{g0015.Name.Id(),0}
        } 
      }  else if (self.Isa.IsIn(C_symbol) == CTRUE) { 
      { var g0016 *ClaireSymbol   = ToSymbol(self)
        _ = g0016
        Result = EID{g0016.Id(),0}
        } 
      }  else if (self.Isa.IsIn(C_Variable) == CTRUE) { 
      { var g0017 *ClaireVariable   = To_Variable(self)
        _ = g0017
        Result = EID{g0017.Pname.Id(),0}
        } 
      }  else if (C_boolean.Id() == self.Isa.Id()) { 
      { var g0018 *ClaireBoolean   = ToBoolean(self)
        _ = g0018
        if (g0018 == CTRUE) { 
          Result = EID{Core.F_symbol_I_string2(MakeString("true")).Id(),0}
          } else {
          Result = EID{Core.F_symbol_I_string2(MakeString("nil")).Id(),0}
          } 
        } 
      } else {
      Result = ToException(Core.C_general_error.Make(MakeString("[147] a name cannot be made from ~S").Id(),MakeConstantList(self).Id())).Close()
      } 
    return Result} 
  
// The EID go function for: iClaire/extract_symbol @ any (throw: true) 
func E_extract_symbol_any (self EID) EID { 
    return F_extract_symbol_any(ANY(self) )} 
  
// we must be sure that the selector (in a has statement or in a message)
// is a property.
//
/* {1} The go function for: iClaire/make_a_property(self:any) [status=1] */
func F_make_a_property_any (self *ClaireAny ) EID { 
    var Result EID 
    if (self.Isa.IsIn(Core.C_global_variable) == CTRUE) { 
      { var g0020 *Core.GlobalVariable   = Core.ToGlobalVariable(self)
        _ = g0020
        Result = F_make_a_property_any(g0020.Value)
        } 
      }  else if (self.Isa.IsIn(C_property) == CTRUE) { 
      { var g0021 *ClaireProperty   = ToProperty(self)
        _ = g0021
        Result = EID{g0021.Id(),0}
        } 
      }  else if (self.Isa.IsIn(C_symbol) == CTRUE) { 
      { var g0022 *ClaireSymbol   = ToSymbol(self)
        { var x *ClaireAny   = g0022.Get()
          if (x.Isa.IsIn(C_property) == CTRUE) { 
            { var g0023 *ClaireProperty   = ToProperty(x)
              _ = g0023
              Result = F_make_a_property_any(g0023.Id())
              } 
            }  else if (x.Isa.IsIn(Core.C_global_variable) == CTRUE) { 
            { var g0024 *Core.GlobalVariable   = Core.ToGlobalVariable(x)
              _ = g0024
              Result = F_make_a_property_any(g0024.Value)
              } 
            } else {
            { var p *ClaireProperty  
              var try_1 EID 
              /*g_try(v2:"try_1",loop:false) */
              try_1 = new(ClaireProperty).IsNamed(C_property,g0022).ToEID()
              /* ERROR PROTECTION INSERTED (p-Result) */
              if ErrorIn(try_1) {Result = try_1
              } else {
              p = ToProperty(OBJ(try_1))
              p.Comment = g0022.String_I()
              /*string->string*/p.Domain = ToType(C_any.Id())
              /*type->type*/p.Range = ToType(C_any.Id())
              /*type->type*/Result = EID{p.Id(),0}
              }
              } 
            } 
          } 
        } 
      }  else if (self.Isa.IsIn(C_unbound_symbol) == CTRUE) { 
      { var g0026 *ClaireUnboundSymbol   = ToUnboundSymbol(self)
        _ = g0026
        Result = F_make_a_property_any(g0026.Name.Id())
        } 
      } else {
      Result = ToException(Core.C_general_error.Make(MakeString("[148] Wrong selector: ~S, cannot make a property\n").Id(),MakeConstantList(self).Id())).Close()
      } 
    return Result} 
  
// The EID go function for: iClaire/make_a_property @ any (throw: true) 
func E_make_a_property_any (self EID) EID { 
    return F_make_a_property_any(ANY(self) )} 
  
// *********************************************************************
// *  Part 4: Pretty printing                                          *
// *********************************************************************
// debug
// create a line break
// if the pretty mode is here ... 
//    (1) pbreak = true means that we create a new line (whatever the length)
//    (2) break = false => we generate a much too far exception
/* {1} The go function for: lbreak(_CL_obj:void) [status=1] */
func F_lbreak_void () EID { 
    var Result EID 
    if (Core.C_pretty.Pprint == CTRUE) { 
      if (Core.C_pretty.Pbreak == CTRUE) { 
        PRINC("\n")
        /*g_try(v2:"Result",loop:true) */
        Result = F_put_buffer_void()
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        Result = F_indent_integer(Core.C_pretty.Index).ToEID()
        }
        }  else if (Core.F_buffer_length_void() > Core.C_pretty.Width) { 
        Result = ToException(new(Core.MuchTooFar).Is(Core.C_much_too_far)).Close()
        } else {
        Result = EID{CFALSE.Id(),0}
        } 
      } else {
      Result = EID{CFALSE.Id(),0}
      } 
    return Result} 
  
// The EID go function for: lbreak @ void (throw: true) 
func E_lbreak_void (_CL_obj EID) EID { 
    return F_lbreak_void( )} 
  
/* {1} The go function for: put_buffer(_CL_obj:void) [status=1] */
func F_put_buffer_void () EID { 
    var Result EID 
    { var buffer *ClaireString  
      _ = buffer
      var try_1 EID 
      /*g_try(v2:"try_1",loop:false) */
      try_1 = Core.F_end_of_string_void()
      /* ERROR PROTECTION INSERTED (buffer-Result) */
      if ErrorIn(try_1) {Result = try_1
      } else {
      buffer = ToString(OBJ(try_1))
      F_princ_string(buffer)
      Core.F_print_in_string_void()
      Result = EID{CEMPTY.Id(),0}
      }
      } 
    return Result} 
  
// The EID go function for: put_buffer @ void (throw: true) 
func E_put_buffer_void (_CL_obj EID) EID { 
    return F_put_buffer_void( )} 
  
/* {1} The go function for: checkfar(_CL_obj:void) [status=1] */
func F_checkfar_void () EID { 
    var Result EID 
    if ((Core.C_pretty.Pprint == CTRUE) && 
        ((Core.C_pretty.Pbreak != CTRUE) && 
          (Core.F_buffer_length_void() > Core.C_pretty.Width))) { 
      Result = ToException(new(Core.MuchTooFar).Is(Core.C_much_too_far)).Close()
      } else {
      Result = EID{CFALSE.Id(),0}
      } 
    return Result} 
  
// The EID go function for: checkfar @ void (throw: true) 
func E_checkfar_void (_CL_obj EID) EID { 
    return F_checkfar_void( )} 
  
/* {1} The go function for: lbreak(n:integer) [status=1] */
func F_lbreak_integer (n int) EID { 
    var Result EID 
    Core.C_pretty.Index = (Core.C_pretty.Index+n)
    /*integer->integer*/Result = F_lbreak_void()
    return Result} 
  
// The EID go function for: lbreak @ integer (throw: true) 
func E_lbreak_integer (n EID) EID { 
    return F_lbreak_integer(INT(n) )} 
  
// indentation
//
/* {1} The go function for: indent(limit:integer) [status=0] */
func F_indent_integer (limit int) *ClaireAny  { 
    // procedure body with s = any 
var Result *ClaireAny  
    { var x int  = Core.F_buffer_length_void()
      _ = x
      Result= CFALSE.Id()
      for (x < limit) { 
        /* While stat, v:"Result" loop:false */
        PRINC(" ")
        x = (x+1)
        /* try?:false, v2:"v_while3" loop will be:tuple("Result", any) */
        } 
      } 
    return Result} 
  
// The EID go function for: indent @ integer (throw: false) 
func E_indent_integer (limit EID) EID { 
    return F_indent_integer(INT(limit) ).ToEID()} 
  
// sets the current_level
/* {1} The go function for: set_level(_CL_obj:void) [status=0] */
func F_set_level_void ()  { 
    // procedure body with s = void 
Core.C_pretty.Index = (Core.F_buffer_length_void()-1)
    /*integer->integer*/} 
  
// The EID go function for: set_level @ void (throw: false) 
func E_set_level_void (_CL_obj EID) EID { 
    F_set_level_void( )
    return EVOID} 
  
/* {1} The go function for: set_level(n:integer) [status=0] */
func F_set_level_integer (n int)  { 
    // procedure body with s = void 
F_set_level_void()
    Core.C_pretty.Index = (Core.C_pretty.Index+n)
    /*integer->integer*/} 
  
// The EID go function for: set_level @ integer (throw: false) 
func E_set_level_integer (n EID) EID { 
    F_set_level_integer(INT(n) )
    return EVOID} 
  
// prints a list as a box in character zone [start, finish], s is the separator (",")
// pbreak = true means that we will print step by step; false => try to add to current place
// if impossible of if pbreak = false, we will switch to printl
// the tricky part is that this method can generate a too far error
/* {1} The go function for: printbox(self:list,start:integer,finish:integer,s:string) [status=1] */
func F_Language_printbox_list1 (self *ClaireList ,start int,finish int,s *ClaireString ) EID { 
    var Result EID 
    { var i int  = 1
      { var startline *ClaireBoolean   = CTRUE
        _ = startline
        { var n int  = self.Length()
          { var _Zl int  = Core.C_pretty.Index
            Core.C_pretty.Index = start
            /*integer->integer*//*g_try(v2:"Result",loop:true) */
            if ((Core.C_pretty.Pprint != CTRUE) || 
                ((Core.F_short_enough_integer((start+10)) != CTRUE) && 
                    (Core.C_pretty.Pbreak == CTRUE))) { 
              Result = F_printl_list(self,s)
              }  else if (Core.C_pretty.Pbreak != CTRUE) { 
              Result = F_printl_list(self,s)
              } else {
              Result= EID{CFALSE.Id(),0}
              for (i <= n) { 
                /* While stat, v:"Result" loop:true */
                var loop_1 EID 
                _ = loop_1
                { 
                for (Core.F_buffer_length_void() < start) { 
                  /* While stat, v:"loop_1" loop:tuple("Result", EID) */
                  PRINC(" ")
                  /* try?:false, v2:"v_while8" loop will be:tuple("loop_1", void) */
                  } 
                /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
                { var idx int  = Core.F_buffer_length_void()
                  _ = idx
                  /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
                  { 
                    h_index := ClEnv.Index
                    h_base := ClEnv.Base
                    Core.C_pretty.Pbreak = CFALSE
                    /*boolean->boolean*//*g_try(v2:"loop_1",loop:false) */
                    loop_1 = F_printexp_any(self.At(i-1),CTRUE)
                    /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
                    if !ErrorIn(loop_1) {
                    { 
                      var va_arg1 *Core.PrettyPrinter  
                      var va_arg2 *ClaireBoolean  
                      va_arg1 = Core.C_pretty
                      va_arg2 = CTRUE
                      va_arg1.Pbreak = va_arg2
                      /*boolean->boolean*/loop_1 = EID{va_arg2.Id(),0}
                      } 
                    }
                    if ErrorIn(loop_1) && ToType(Core.C_much_too_far.Id()).Contains(ANY(loop_1)) == CTRUE { 
                      ClEnv.Index = h_index
                      ClEnv.Base = h_base
                      Core.C_pretty.Pbreak = CTRUE
                      /*boolean->boolean*/{ 
                        var va_arg1 *Core.PrettyPrinter  
                        var va_arg2 int 
                        va_arg1 = Core.C_pretty
                        va_arg2 = start
                        va_arg1.Index = va_arg2
                        /*integer->integer*/loop_1 = EID{C__INT,IVAL(va_arg2)}
                        } 
                      } 
                    } 
                  /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
                  if ErrorIn(loop_1) {Result = loop_1
                  break
                  } else {
                  if (i != n) { 
                    F_princ_string(s)
                    } 
                  /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
                  if (Core.F_buffer_length_void() < finish) { 
                    i = (i+1)
                    startline = CFALSE
                    loop_1 = EID{startline.Id(),0}
                    } else {
                    Core.F_buffer_set_length_integer(idx)
                    /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
                    if (startline != CTRUE) { 
                      /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
                      loop_1 = F_lbreak_void()
                      /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
                      if ErrorIn(loop_1) {Result = loop_1
                      break
                      } else {
                      startline = CTRUE
                      loop_1 = EID{startline.Id(),0}
                      }
                      } else {
                      F_set_level_void()
                      Core.C_pretty.Index = (Core.C_pretty.Index+1)
                      /*integer->integer*//*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
                      loop_1 = F_printexp_any(self.At(i-1),CTRUE)
                      /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
                      if ErrorIn(loop_1) {Result = loop_1
                      break
                      } else {
                      Core.C_pretty.Index = _Zl
                      /*integer->integer*//*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
                      if (i != n) { 
                        F_princ_string(s)
                        /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
                        loop_1 = F_lbreak_void()
                        /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
                        if ErrorIn(loop_1) {Result = loop_1
                        break
                        } else {
                        }
                        } else {
                        loop_1 = EID{CFALSE.Id(),0}
                        } 
                      /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
                      if ErrorIn(loop_1) {Result = loop_1
                      break
                      } else {
                      i = (i+1)
                      loop_1 = EID{C__INT,IVAL(i)}
                      }}
                      } 
                    /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
                    if ErrorIn(loop_1) {Result = loop_1
                    break
                    } else {
                    }
                    } 
                  /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
                  if ErrorIn(loop_1) {Result = loop_1
                  break
                  } else {
                  }}
                  } 
                /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
                if ErrorIn(loop_1) {Result = loop_1
                break
                } else {
                }
                /* try?:false, v2:"v_while7" loop will be:tuple("Result", EID) */
                } 
              }
              } 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            Core.C_pretty.Index = _Zl
            /*integer->integer*/Result = EID{CNULL,0}
            }
            } 
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: printbox @ list<type_expression>(list, integer, integer, string) (throw: true) 
func E_Language_printbox_list1 (self EID,start EID,finish EID,s EID) EID { 
    return F_Language_printbox_list1(ToList(OBJ(self)),
      INT(start),
      INT(finish),
      ToString(OBJ(s)) )} 
  
// default value of arguments
//
/* {1} The go function for: printbox(self:list) [status=1] */
func F_Language_printbox_list2 (self *ClaireList ) EID { 
    var Result EID 
    Result = F_Language_printbox_list1(self,Core.F_buffer_length_void(),Core.C_pretty.Width,MakeString(", "))
    return Result} 
  
// The EID go function for: printbox @ list<type_expression>(list) (throw: true) 
func E_Language_printbox_list2 (self EID) EID { 
    return F_Language_printbox_list2(ToList(OBJ(self)) )} 
  
/* {1} The go function for: printbox(self:list,s:string) [status=1] */
func F_Language_printbox_list3 (self *ClaireList ,s *ClaireString ) EID { 
    var Result EID 
    Result = F_Language_printbox_list1(self,Core.F_buffer_length_void(),Core.C_pretty.Width,s)
    return Result} 
  
// The EID go function for: printbox @ list<type_expression>(list, string) (throw: true) 
func E_Language_printbox_list3 (self EID,s EID) EID { 
    return F_Language_printbox_list3(ToList(OBJ(self)),ToString(OBJ(s)) )} 
  
// this is a tricky method : first try to print without pretty (box) 
/* {1} The go function for: printl(self:list,s:string) [status=1] */
func F_printl_list (self *ClaireList ,s *ClaireString ) EID { 
    var Result EID 
    { var f *ClaireBoolean   = CTRUE
      _ = f
      { var b *ClaireBoolean   = Core.C_pretty.Pprint
        Core.C_pretty.Pprint = CFALSE
        /*boolean->boolean*//*g_try(v2:"Result",loop:true) */
        { 
          h_index := ClEnv.Index
          h_base := ClEnv.Base
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
              { 
              if (f == CTRUE) { 
                f = CFALSE
                } else {
                F_princ_string(s)
                } 
              /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
              loop_1 = F_printexp_any(x,CTRUE)
              /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
              if ErrorIn(loop_1) {Result = loop_1
              break
              } else {
              /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
              if ((b == CTRUE) && 
                  ((Core.C_pretty.Pbreak != CTRUE) && 
                    (Core.F_buffer_length_void() > Core.C_pretty.Width))) { 
                Core.C_pretty.Pprint = b
                /*boolean->boolean*//*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
                loop_1 = ToException(new(Core.MuchTooFar).Is(Core.C_much_too_far)).Close()
                /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
                if ErrorIn(loop_1) {Result = loop_1
                break
                } else {
                }
                } else {
                loop_1 = EID{CFALSE.Id(),0}
                } 
              /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
              if ErrorIn(loop_1) {Result = loop_1
              break
              } else {
              }}
              }
              } 
            } 
          if ErrorIn(Result) && ToType(C_system_error.Id()).Contains(ANY(Result)) == CTRUE { 
            ClEnv.Index = h_index
            ClEnv.Base = h_base
            { var x *ClaireException   = ClEnv.Exception_I
              if ((b == CTRUE) && 
                  (ToSystemError(x.Id()).Index == 16)) { 
                Core.C_pretty.Pprint = b
                /*boolean->boolean*/Result = ToException(new(Core.MuchTooFar).Is(Core.C_much_too_far)).Close()
                } else {
                Result = x.Close()
                } 
              } 
            } 
          } 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        { 
          var va_arg1 *Core.PrettyPrinter  
          var va_arg2 *ClaireBoolean  
          va_arg1 = Core.C_pretty
          va_arg2 = b
          va_arg1.Pprint = va_arg2
          /*boolean->boolean*/Result = EID{va_arg2.Id(),0}
          } 
        }
        } 
      } 
    return Result} 
  
// The EID go function for: printl @ list (throw: true) 
func E_printl_list (self EID,s EID) EID { 
    return F_printl_list(ToList(OBJ(self)),ToString(OBJ(s)) )} 
  
// print bounded prints a bounded expression using ( and )
/* {1} The go function for: printexp(self:any,comp:boolean) [status=1] */
func F_printexp_any (self *ClaireAny ,comp *ClaireBoolean ) EID { 
    var Result EID 
    var g0032I *ClaireBoolean  
    { 
      /* Or stat: v="g0032I", loop=true */
      var v_or2 *ClaireBoolean  
      
      /* Or stat: try (if (inherit? @ class(owner @ any(self),Call)) let g0031:Call := (<self:Call>) in not @ any(((inherit? @ class(isa @ any(selector @ Call(g0031)),operation)) & (not @ any(comp)) & (= @ any(length @ list(args @ Call(g0031)),2)))) else false) with try:false, v="g0032I", loop=true */
      if (self.Isa.IsIn(C_Call) == CTRUE) { 
        { var g0031 *Call   = To_Call(self)
          v_or2 = MakeBoolean((g0031.Selector.Isa.IsIn(C_operation) == CTRUE) && (comp != CTRUE) && (g0031.Args.Length() == 2)).Not
          } 
        } else {
        v_or2 = CFALSE
        } 
      if (v_or2 == CTRUE) {g0032I = CTRUE
      } else { 
        /* Or stat: try inherit? @ class(owner @ any(self),Collect) with try:false, v="g0032I", loop=true */
        v_or2 = self.Isa.IsIn(C_Collect)
        if (v_or2 == CTRUE) {g0032I = CTRUE
        } else { 
          /* Or stat: try inherit? @ class(owner @ any(self),Select) with try:false, v="g0032I", loop=true */
          v_or2 = self.Isa.IsIn(C_Select)
          if (v_or2 == CTRUE) {g0032I = CTRUE
          } else { 
            /* Or stat: try inherit? @ class(owner @ any(self),Definition) with try:false, v="g0032I", loop=true */
            v_or2 = self.Isa.IsIn(C_Definition)
            if (v_or2 == CTRUE) {g0032I = CTRUE
            } else { 
              /* Or stat: try inherit? @ class(owner @ any(self),Construct) with try:false, v="g0032I", loop=true */
              v_or2 = self.Isa.IsIn(C_Construct)
              if (v_or2 == CTRUE) {g0032I = CTRUE
              } else { 
                /* Or stat: try inherit? @ class(owner @ any(self),Do) with try:false, v="g0032I", loop=true */
                v_or2 = self.Isa.IsIn(C_Do)
                if (v_or2 == CTRUE) {g0032I = CTRUE
                } else { 
                  /* Or stat: try = @ any(self,unknown) with try:false, v="g0032I", loop=true */
                  v_or2 = Equal(self,CNULL)
                  if (v_or2 == CTRUE) {g0032I = CTRUE
                  } else { 
                    /* Or stat: try inherit? @ class(owner @ any(self),And) with try:false, v="g0032I", loop=true */
                    v_or2 = self.Isa.IsIn(C_And)
                    if (v_or2 == CTRUE) {g0032I = CTRUE
                    } else { 
                      /* Or stat: try inherit? @ class(owner @ any(self),primitive) with try:false, v="g0032I", loop=true */
                      v_or2 = self.Isa.IsIn(C_primitive)
                      if (v_or2 == CTRUE) {g0032I = CTRUE
                      } else { 
                        /* Or stat: try inherit? @ class(owner @ any(self),Or) with try:false, v="g0032I", loop=true */
                        v_or2 = self.Isa.IsIn(C_Or)
                        if (v_or2 == CTRUE) {g0032I = CTRUE
                        } else { 
                          /* Or stat: try inherit? @ class(owner @ any(self),If) with try:false, v="g0032I", loop=true */
                          v_or2 = self.Isa.IsIn(C_If)
                          if (v_or2 == CTRUE) {g0032I = CTRUE
                          } else { 
                            /* Or stat: try inherit? @ class(owner @ any(self),restriction) with try:false, v="g0032I", loop=true */
                            v_or2 = self.Isa.IsIn(C_restriction)
                            if (v_or2 == CTRUE) {g0032I = CTRUE
                            } else { 
                              /* Or stat: try inherit? @ class(owner @ any(self),unbound_symbol) with try:false, v="g0032I", loop=true */
                              v_or2 = self.Isa.IsIn(C_unbound_symbol)
                              if (v_or2 == CTRUE) {g0032I = CTRUE
                              } else { 
                                /* Or stat: try inherit? @ class(owner @ any(self),Variable) with try:false, v="g0032I", loop=true */
                                v_or2 = self.Isa.IsIn(C_Variable)
                                if (v_or2 == CTRUE) {g0032I = CTRUE
                                } else { 
                                  /* Or stat: try not @ any(inherit? @ class(owner @ any(self),Instruction)) with try:false, v="g0032I", loop=true */
                                  v_or2 = self.Isa.IsIn(C_Instruction).Not
                                  if (v_or2 == CTRUE) {g0032I = CTRUE
                                  } else { 
                                    g0032I = CFALSE} 
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
    if (g0032I == CTRUE) { 
      Result = Core.F_CALL(C_print,ARGS(self.ToEID()))
      } else {
      { var _Zl int  = Core.C_pretty.Index
        _ = _Zl
        /*g_try(v2:"Result",loop:true) */
        PRINC("(")
        F_set_level_integer(1)
        /*g_try(v2:"Result",loop:true) */
        Result = Core.F_CALL(C_print,ARGS(self.ToEID()))
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC(")")
        Result = EVOID
        }
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        { 
          var va_arg1 *Core.PrettyPrinter  
          var va_arg2 int 
          va_arg1 = Core.C_pretty
          va_arg2 = _Zl
          va_arg1.Index = va_arg2
          /*integer->integer*/Result = EID{C__INT,IVAL(va_arg2)}
          } 
        }
        } 
      } 
    return Result} 
  
// The EID go function for: printexp @ any (throw: true) 
func E_printexp_any (self EID,comp EID) EID { 
    return F_printexp_any(ANY(self),ToBoolean(OBJ(comp)) )} 
  
// pretty print is using the buffered print (into a string)
/* {1} The go function for: pretty_print(self:any) [status=1] */
func F_pretty_print_any (self *ClaireAny ) EID { 
    var Result EID 
    Core.F_print_in_string_void()
    Core.C_pretty.Pprint = CTRUE
    /*boolean->boolean*/Core.C_pretty.Pbreak = CTRUE
    /*boolean->boolean*/Core.C_pretty.Index = 0
    /*integer->integer*//*g_try(v2:"Result",loop:true) */
    Result = Core.F_CALL(C_print,ARGS(self.ToEID()))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    Core.C_pretty.Pprint = CFALSE
    /*boolean->boolean*/{ var arg_1 *ClaireString  
      _ = arg_1
      var try_2 EID 
      /*g_try(v2:"try_2",loop:false) */
      try_2 = Core.F_end_of_string_void()
      /* ERROR PROTECTION INSERTED (arg_1-Result) */
      if ErrorIn(try_2) {Result = try_2
      } else {
      arg_1 = ToString(OBJ(try_2))
      F_princ_string(arg_1)
      Result = EVOID
      }
      } 
    }
    return Result} 
  
// The EID go function for: pretty_print @ any (throw: true) 
func E_pretty_print_any (self EID) EID { 
    return F_pretty_print_any(ANY(self) )} 
  
// self_print uses the default boxing
/* {1} The go function for: self_print(self:list) [status=1] */
func F_self_print_list_Language (self *ClaireList ) EID { 
    var Result EID 
    /*g_try(v2:"Result",loop:true) */
    if (Equal(self.Of().Id(),CEMPTY.Id()) != CTRUE) { 
      PRINC("list<")
      /*g_try(v2:"Result",loop:true) */
      Result = Core.F_print_any(self.Of().Id())
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(">")
      Result = EVOID
      }
      } else {
      PRINC("list")
      Result = EVOID
      } 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("(")
    /*g_try(v2:"Result",loop:true) */
    Result = F_Language_printbox_list2(self)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(")")
    Result = EVOID
    }
    }
    return Result} 
  
// The EID go function for: self_print @ list (throw: true) 
func E_self_print_list_Language (self EID) EID { 
    return F_self_print_list_Language(ToList(OBJ(self)) )} 
  
/* {1} The go function for: self_print(self:set) [status=1] */
func F_self_print_set_Language (self *ClaireSet ) EID { 
    var Result EID 
    if (Equal(ToList(self.Id()).Of().Id(),CEMPTY.Id()) == CTRUE) { 
      PRINC("{")
      /*g_try(v2:"Result",loop:true) */
      Result = F_Language_printbox_list2(self.List_I())
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC("}")
      Result = EVOID
      }
      } else {
      /*g_try(v2:"Result",loop:true) */
      PRINC("set<")
      /*g_try(v2:"Result",loop:true) */
      Result = Core.F_print_any(ToList(self.Id()).Of().Id())
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(">")
      Result = EVOID
      }
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC("(")
      /*g_try(v2:"Result",loop:true) */
      Result = F_Language_printbox_list2(self.List_I())
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(")")
      Result = EVOID
      }
      }
      } 
    return Result} 
  
// The EID go function for: self_print @ set (throw: true) 
func E_self_print_set_Language (self EID) EID { 
    return F_self_print_set_Language(ToSet(OBJ(self)) )} 
  
// to remove !
/* {1} The go function for: self_print(self:tuple) [status=1] */
func F_self_print_tuple_Language (self *ClaireTuple ) EID { 
    var Result EID 
    PRINC("tuple(")
    /*g_try(v2:"Result",loop:true) */
    Result = F_Language_printbox_list2(ToList(self.Id()))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(")")
    Result = EVOID
    }
    return Result} 
  
// The EID go function for: self_print @ tuple (throw: true) 
func E_self_print_tuple_Language (self EID) EID { 
    return F_self_print_tuple_Language(ToTuple(OBJ(self)) )} 
  
// a map_set 
/* {1} The go function for: self_print(self:map_set) [status=1] */
func F_self_print_map_set (self *ClaireMapSet ) EID { 
    var Result EID 
    PRINC("map<")
    /*g_try(v2:"Result",loop:true) */
    Result = Core.F_print_any(self.Domain().Id())
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(",")
    /*g_try(v2:"Result",loop:true) */
    Result = Core.F_print_any(self.Range().Id())
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(">")
    Result = EVOID
    }}
    return Result} 
  
// The EID go function for: self_print @ map_set (throw: true) 
func E_self_print_map_set (self EID) EID { 
    return F_self_print_map_set(ToMapSet(OBJ(self)) )} 
  
// a pair
/* {1} The go function for: self_print(x:pair) [status=1] */
func F_self_print_pair (x *ClairePair ) EID { 
    var Result EID 
    /*g_try(v2:"Result",loop:true) */
    Result = F_printexp_any(x.First,CFALSE)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(":")
    Result = F_printexp_any(x.Second,CFALSE)
    }
    return Result} 
  
// The EID go function for: self_print @ pair (throw: true) 
func E_self_print_pair (x EID) EID { 
    return F_self_print_pair(ToPair(OBJ(x)) )} 
  
// *********************************************************************
// *  Part 5: simple type inference  (class based)                     *
// *********************************************************************
// this is a simple, self-contained, type inference method that mimicks what GO is bound to know
// it is used to check the type safety of the gerenated code in the Generate module and it is used
// in call.cl to produce OFTO (on-the-fly optimization) => see readcall
// s_type =  static type, or stupid_type  (we should remove stupid_t)
/* {1} The go function for: static_type(self:any) [status=1] */
func F_static_type_any (self *ClaireAny ) EID { 
    var Result EID 
    if (self.Isa.IsIn(C_Variable) == CTRUE) { 
      { var g0033 *ClaireVariable   = To_Variable(self)
        _ = g0033
        Result = EID{g0033.Range.Class_I().Id(),0}
        } 
      }  else if (self.Isa.IsIn(Core.C_global_variable) == CTRUE) { 
      { var g0034 *Core.GlobalVariable   = Core.ToGlobalVariable(self)
        { var r *ClaireType   = g0034.Range
          if (F_boolean_I_any(r.Id()) == CTRUE) { 
            Result = EID{r.Class_I().Id(),0}
            } else {
            Result = EID{g0034.Value.Isa.Id(),0}
            } 
          } 
        } 
      }  else if (self.Isa.IsIn(C_And) == CTRUE) { 
      Result = EID{C_boolean.Id(),0}
      }  else if (self.Isa.IsIn(C_Or) == CTRUE) { 
      Result = EID{C_boolean.Id(),0}
      }  else if (self.Isa.IsIn(C_environment) == CTRUE) { 
      Result = EID{C_environment.Id(),0}
      }  else if (self.Isa.IsIn(C_Call_plus) == CTRUE) { 
      { var g0038 *Call_plus   = To_Call_plus(self)
        { var p *ClaireProperty   = g0038.Selector
          { var s *ClaireObject  
            var try_1 EID 
            /*g_try(v2:"try_1",loop:false) */
            { var arg_2 *ClaireClass  
              _ = arg_2
              var try_3 EID 
              /*g_try(v2:"try_3",loop:false) */
              try_3 = F_static_type_any(g0038.Args.At(1-1))
              /* ERROR PROTECTION INSERTED (arg_2-try_1) */
              if ErrorIn(try_3) {try_1 = try_3
              } else {
              arg_2 = ToClass(OBJ(try_3))
              try_1 = EID{Core.F__at_property1(p,arg_2).Id(),0}
              }
              } 
            /* ERROR PROTECTION INSERTED (s-Result) */
            if ErrorIn(try_1) {Result = try_1
            } else {
            s = ToObject(OBJ(try_1))
            if (C_slot.Id() == s.Isa.Id()) { 
              { var g0039 *ClaireSlot   = ToSlot(s.Id())
                _ = g0039
                Result = EID{g0039.Range.Id(),0}
                } 
              } else {
              Result = EID{p.Range.Id(),0}
              } 
            }
            } 
          } 
        } 
      }  else if (self.Isa.IsIn(C_Call_slot) == CTRUE) { 
      { var g0041 *CallSlot   = To_CallSlot(self)
        _ = g0041
        { var s *ClaireSlot   = g0041.Selector
          { var p *ClaireProperty   = s.Selector
            _ = p
            { 
              var s2 *ClaireRestriction  
              _ = s2
              var s2_iter *ClaireAny  
              for _,s2_iter = range(p.Definition.ValuesO()){ 
                s2 = ToRestriction(s2_iter)
                if (C_slot.Id() == s2.Isa.Id()) { 
                  { var g0042 *ClaireSlot   = ToSlot(s2.Id())
                    _ = g0042
                    if (ToType(Core.F_domain_I_restriction(ToRestriction(s.Id())).Id()).Included(ToType(Core.F_domain_I_restriction(ToRestriction(g0042.Id())).Id())) == CTRUE) { 
                      s = g0042
                      } 
                    } 
                  } 
                } 
              } 
            Result = EID{s.Range.Class_I().Id(),0}
            } 
          } 
        } 
      }  else if (self.Isa.IsIn(C_Call_method) == CTRUE) { 
      { var g0043 *CallMethod   = To_CallMethod(self)
        { var p *ClaireProperty   = g0043.Arg.Selector
          _ = p
          if (p.Id() == C_nth.Id()) { 
            Result = F_Language_static_type_nth_any(g0043.Args.At(1-1))
            } else {
            Result = EID{g0043.Arg.Range.Class_I().Id(),0}
            } 
          } 
        } 
      }  else if (self.Isa.IsIn(C_Call) == CTRUE) { 
      { var g0044 *Call   = To_Call(self)
        { var p *ClaireProperty   = g0044.Selector
          if (p.Id() == C_nth.Id()) { 
            Result = F_Language_static_type_nth_any(g0044.Args.At(1-1))
            } else {
            Result = EID{p.Range.Class_I().Id(),0}
            } 
          } 
        } 
      }  else if (self.Isa.IsIn(C_Assign) == CTRUE) { 
      { var g0045 *Assign   = To_Assign(self)
        _ = g0045
        Result = F_static_type_any(g0045.Arg)
        } 
      }  else if (self.Isa.IsIn(C_Let) == CTRUE) { 
      { var g0046 *Let   = To_Let(self)
        _ = g0046
        Result = F_static_type_any(g0046.Arg)
        } 
      }  else if (self.Isa.IsIn(C_Do) == CTRUE) { 
      { var g0047 *Do   = To_Do(self)
        _ = g0047
        { var arg_4 *ClaireAny  
          _ = arg_4
          var try_5 EID 
          /*g_try(v2:"try_5",loop:false) */
          try_5 = Core.F_last_list(g0047.Args)
          /* ERROR PROTECTION INSERTED (arg_4-Result) */
          if ErrorIn(try_5) {Result = try_5
          } else {
          arg_4 = ANY(try_5)
          Result = F_static_type_any(arg_4)
          }
          } 
        } 
      }  else if (self.Isa.IsIn(C_If) == CTRUE) { 
      { var g0048 *If   = To_If(self)
        { var arg_6 *ClaireClass  
          _ = arg_6
          var try_8 EID 
          /*g_try(v2:"try_8",loop:false) */
          try_8 = F_static_type_any(g0048.Arg)
          /* ERROR PROTECTION INSERTED (arg_6-Result) */
          if ErrorIn(try_8) {Result = try_8
          } else {
          arg_6 = ToClass(OBJ(try_8))
          { var arg_7 *ClaireClass  
            _ = arg_7
            var try_9 EID 
            /*g_try(v2:"try_9",loop:false) */
            try_9 = F_static_type_any(g0048.Other)
            /* ERROR PROTECTION INSERTED (arg_7-Result) */
            if ErrorIn(try_9) {Result = try_9
            } else {
            arg_7 = ToClass(OBJ(try_9))
            Result = EID{Core.F_meet_class(arg_6,arg_7).Id(),0}
            }
            } 
          }
          } 
        } 
      }  else if (self.Isa.IsIn(C_Collect) == CTRUE) { 
      Result = EID{C_list.Id(),0}
      }  else if (self.Isa.IsIn(C_Image) == CTRUE) { 
      Result = EID{C_set.Id(),0}
      }  else if (self.Isa.IsIn(C_Select) == CTRUE) { 
      Result = EID{C_set.Id(),0}
      }  else if (self.Isa.IsIn(C_Lselect) == CTRUE) { 
      Result = EID{C_list.Id(),0}
      }  else if (self.Isa.IsIn(C_List) == CTRUE) { 
      Result = EID{C_list.Id(),0}
      }  else if (self.Isa.IsIn(C_Set) == CTRUE) { 
      Result = EID{C_set.Id(),0}
      }  else if (self.Isa.IsIn(C_Tuple) == CTRUE) { 
      Result = EID{C_tuple.Id(),0}
      }  else if (self.Isa.IsIn(C_Exists) == CTRUE) { 
      { var g0056 *Exists   = To_Exists(self)
        _ = g0056
        if (g0056.Other == CNULL) { 
          Result = EID{C_any.Id(),0}
          } else {
          Result = EID{C_boolean.Id(),0}
          } 
        } 
      }  else if (self.Isa.IsIn(C_Definition) == CTRUE) { 
      { var g0057 *Definition   = To_Definition(self)
        _ = g0057
        Result = EID{g0057.Arg.Id(),0}
        } 
      }  else if (self.Isa.IsIn(C_Instruction) == CTRUE) { 
      Result = EID{C_any.Id(),0}
      } else {
      Result = EID{self.Isa.Id(),0}
      } 
    return Result} 
  
// The EID go function for: static_type @ any (throw: true) 
func E_static_type_any (self EID) EID { 
    return F_static_type_any(ANY(self) )} 
  
// second order pattern for a very common case l[i] where l:list<X>
/* {1} The go function for: static_type_nth(x:any) [status=1] */
func F_Language_static_type_nth_any (x *ClaireAny ) EID { 
    var Result EID 
    if (x.Isa.IsIn(C_Variable) == CTRUE) { 
      { var g0060 *ClaireVariable   = To_Variable(x)
        _ = g0060
        { var s *ClaireType   = g0060.Range
          if (s.Isa.IsIn(C_Param) == CTRUE) { 
            { var g0061 *ClaireParam   = To_Param(s.Id())
              if (g0061.Params.At(1-1) == C_of.Id()) { 
                { var arg_1 *ClaireAny  
                  _ = arg_1
                  var try_2 EID 
                  /*g_try(v2:"try_2",loop:false) */
                  try_2 = Core.F_the_type(ToType(g0061.Args.At(1-1)))
                  /* ERROR PROTECTION INSERTED (arg_1-Result) */
                  if ErrorIn(try_2) {Result = try_2
                  } else {
                  arg_1 = ANY(try_2)
                  Result = EID{ToTypeExpression(arg_1).Class_I().Id(),0}
                  }
                  } 
                } else {
                Result = EID{C_any.Id(),0}
                } 
              } 
            } else {
            Result = EID{C_any.Id(),0}
            } 
          } 
        } 
      } else {
      Result = EID{C_any.Id(),0}
      } 
    return Result} 
  
// The EID go function for: static_type_nth @ any (throw: true) 
func E_Language_static_type_nth_any (x EID) EID { 
    return F_Language_static_type_nth_any(ANY(x) )} 
  
// end of file