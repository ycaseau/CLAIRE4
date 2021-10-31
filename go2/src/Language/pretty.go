/***** CLAIRE Compilation of file /Users/ycaseau/Dropbox/src/clairev4.0/src/meta/pretty.cl 
         [version 4.0.01 / safety 5] Saturday 10-30-2021 *****/

package Language
import (_ "fmt"
	. "Kernel"
	"Core"
)

//-------- dumb function to prevent import errors --------
func import_g0035() { 
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
/* {1} OPT.The go function for: no_eval(self:Instruction) [] */
func F_no_eval_Instruction (self *ClaireInstruction ) EID { 
    var Result EID 
    Result = ToException(Core.C_general_error.Make(MakeString("[144] evaluate(~S) is not defined").Id(),MakeConstantList(self.Id().Isa.Id()).Id())).Close()
    return Result} 
  
// The EID go function for: no_eval @ Instruction (throw: true) 
func E_no_eval_Instruction (self EID) EID { 
    return /*(sm for no_eval @ Instruction= EID)*/ F_no_eval_Instruction(To_Instruction(OBJ(self)) )} 
  
// import => cannot work in CLAIRE4
// *********************************************************************
// *   Part 1: unbound_symbol and variables                            *
// *********************************************************************
// An unbound_symbol is created by the reader when a symbol is not bound
//
//unbound_symbol <: Basic_instruction(identifier:symbol)
/* {1} OPT.The go function for: self_print(self:unbound_symbol) [] */
func F_self_print_unbound_symbol_Language (self *ClaireUnboundSymbol )  { 
    // procedure body with s =  
self.Name.Princ()
    PRINC("")
    } 
  
// The EID go function for: self_print @ unbound_symbol (throw: false) 
func E_self_print_unbound_symbol_Language (self EID) EID { 
    /*(sm for self_print @ unbound_symbol= void)*/ F_self_print_unbound_symbol_Language(ToUnboundSymbol(OBJ(self)) )
    return EVOID} 
  
/* {1} OPT.The go function for: self_eval(self:unbound_symbol) [] */
func F_self_eval_unbound_symbol (self *ClaireUnboundSymbol ) EID { 
    var Result EID 
    if (Core.F_owner_any(self.Name.Get()).IsIn(C_thing) == CTRUE) /* If:2 */{ 
      Result = EVAL(self.Name.Get())
      } else {
      Result = ToException(Core.C_general_error.Make(MakeString("[145] the symbol ~A is unbound").Id(),MakeConstantList(self.Name.Id()).Id())).Close()
      /* If-2 */} 
    return Result} 
  
// The EID go function for: self_eval @ unbound_symbol (throw: true) 
func E_self_eval_unbound_symbol (self EID) EID { 
    return /*(sm for self_eval @ unbound_symbol= EID)*/ F_self_eval_unbound_symbol(ToUnboundSymbol(OBJ(self)) )} 
  
// The EVAL go function for: unbound_symbol 
func EVAL_unbound_symbol (x *ClaireAny) EID { 
     return F_self_eval_unbound_symbol(ToUnboundSymbol(x))} 
  
// A lexical variable is defined by a "Let" or inside a method's definition
// Lexical variables --------------------------------------------------
//
//
/* {1} OPT.The go function for: self_print(self:Variable) [] */
func F_self_print_Variable_Language (self *ClaireVariable )  { 
    // procedure body with s =  
/* Let:2 */{ 
      var s *ClaireSymbol   = self.Pname
      /* noccur = 2 */
      if (s.Id() == CNULL) /* If:3 */{ 
        PRINC("V?")
        } else {
        s.Princ()
        /* If-3 */} 
      /* Let-2 */} 
    } 
  
// The EID go function for: self_print @ Variable (throw: false) 
func E_self_print_Variable_Language (self EID) EID { 
    /*(sm for self_print @ Variable= void)*/ F_self_print_Variable_Language(To_Variable(OBJ(self)) )
    return EVOID} 
  
/* {1} OPT.The go function for: ppvariable(self:Variable) [] */
func F_ppvariable_Variable (self *ClaireVariable ) EID { 
    var Result EID 
    if (self.Range.Id() != CNULL) /* If:2 */{ 
      self.Pname.Princ()
      PRINC(":")
      Result = F_printexp_any(self.Range.Id(),CFALSE)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC("")
      Result = EVOID
      }
      } else {
      self.Pname.Princ()
      Result = EVOID
      /* If-2 */} 
    return Result} 
  
// The EID go function for: ppvariable @ Variable (throw: true) 
func E_ppvariable_Variable (self EID) EID { 
    return /*(sm for ppvariable @ Variable= EID)*/ F_ppvariable_Variable(To_Variable(OBJ(self)) )} 
  
/* {1} OPT.The go function for: ppvariable(self:list) [] */
func F_ppvariable_list (self *ClaireList ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var f *ClaireBoolean   = CTRUE
      /* noccur = 2 */
      /* For:3 */{ 
        var v *ClaireAny  
        _ = v
        Result= EID{CFALSE.Id(),0}
        var v_support *ClaireList  
        v_support = self
        v_len := v_support.Length()
        for i_it := 0; i_it < v_len; i_it++ { 
          v = v_support.At(i_it)
          var void_try5 EID 
          _ = void_try5
          { 
          if (f == CTRUE) /* If:5 */{ 
            f = CFALSE
            } else {
            PRINC(",")
            /* If-5 */} 
          if (v.Isa.IsIn(C_Variable) == CTRUE) /* If:5 */{ 
            /* Let:6 */{ 
              var g0036 *ClaireVariable   = To_Variable(v)
              /* noccur = 1 */
              void_try5 = F_ppvariable_Variable(g0036)
              /* Let-6 */} 
            } else {
            void_try5 = Core.F_CALL(C_print,ARGS(v.ToEID()))
            /* If-5 */} 
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
  
// The EID go function for: ppvariable @ list (throw: true) 
func E_ppvariable_list (self EID) EID { 
    return /*(sm for ppvariable @ list= EID)*/ F_ppvariable_list(ToList(OBJ(self)) )} 
  
/* {1} OPT.The go function for: self_eval(self:Variable) [] */
func F_self_eval_Variable (self *ClaireVariable ) EID { 
    var Result EID 
    Result = ClEnv.EvalStack[(ClEnv.Base+self.Index)]
    return Result} 
  
// The EID go function for: self_eval @ Variable (throw: true) 
func E_self_eval_Variable (self EID) EID { 
    return /*(sm for self_eval @ Variable= EID)*/ F_self_eval_Variable(To_Variable(OBJ(self)) )} 
  
/* {1} OPT.The go function for: write_value(self:Variable,val:any) [] */
func F_write_value_Variable (self *ClaireVariable ,val *ClaireAny ) EID { 
    var Result EID 
    if ((self.Range.Id() == CNULL) || 
        (self.Range.Contains(val) == CTRUE)) /* If:2 */{ 
      ClEnv.EvalStack[(ClEnv.Base+self.Index)]=val.ToEID()
      Result = val.ToEID()
      } else {
      Result = ToException(Core.C_range_error.Make(self.Id(),val,self.Range.Id())).Close()
      /* If-2 */} 
    return Result} 
  
// The EID go function for: write_value @ Variable (throw: true) 
func E_write_value_Variable (self EID,val EID) EID { 
    return /*(sm for write_value @ Variable= EID)*/ F_write_value_Variable(To_Variable(OBJ(self)),ANY(val) )} 
  
// this is the definition of a typed variable / Vardef is a syntactic marker
// in CLAIRE 4, Vardef are transformed in Var at run time
// this is strange and should be fixed  or understood
/* {1} OPT.The go function for: self_eval(self:Vardef) [] */
func (self *Vardef ) SelfEval () EID { 
    var Result EID 
    /* Let:2 */{ 
      var i *ClaireAny   = MakeInteger(self.Index).Id()
      /* noccur = 2 */
      if (ToBoolean(OBJ(Core.F_CALL(ToProperty(C__sup_equal.Id()),ARGS(i.ToEID(),EID{C__INT,IVAL(0)})))) == CTRUE) /* If:3 */{ 
        Result = ClEnv.EvalStack[INT(Core.F_CALL(ToProperty(Core.C__plus.Id()),ARGS(EID{C__INT,IVAL(ClEnv.Base)},i.ToEID())))]
        } else {
        Result = ToException(Core.C_general_error.Make(MakeString("[146] The variable ~S is not defined").Id(),MakeConstantList(self.Id()).Id())).Close()
        /* If-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: self_eval @ Vardef (throw: true) 
func E_self_eval_Vardef (self EID) EID { 
    return /*(sm for self_eval @ Vardef= EID)*/ To_Vardef(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: Vardef 
func EVAL_Vardef (x *ClaireAny) EID { 
     return To_Vardef(x).SelfEval()} 
  
//   [self_print(self:Vardef) : any -> ppvariable(self) ]
// global_variables are defined in exception ? ---------------------------
// a global variable is a named object with a special evaluation
//
// self_eval(self:global_variable) : any -> self.value  -> moved to object.cl
/* {1} OPT.The go function for: write_value(self:global_variable,val:any) [] */
func F_write_value_global_variable (self *Core.GlobalVariable ,val *ClaireAny ) EID { 
    var Result EID 
    if (self.Range.Contains(val) == CTRUE) /* If:2 */{ 
      Result = Core.F_put_store_property2(C_value,ToObject(self.Id()),val,self.Store_ask)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = val.ToEID()
      }
      } else {
      Result = ToException(Core.C_range_error.Make(self.Id(),val,self.Range.Id())).Close()
      /* If-2 */} 
    return Result} 
  
// The EID go function for: write_value @ global_variable (throw: true) 
func E_write_value_global_variable (self EID,val EID) EID { 
    return /*(sm for write_value @ global_variable= EID)*/ F_write_value_global_variable(Core.ToGlobalVariable(OBJ(self)),ANY(val) )} 
  
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
/* {1} OPT.The go function for: apply(self:lambda,%l:list) [] */
func F_apply_lambda (self *ClaireLambda ,_Zl *ClaireList ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var start int  = ClEnv.Index
      /* noccur = 2 */
      /* Let:3 */{ 
        var retour int  = ClEnv.Base
        /* noccur = 1 */
        ClEnv.Base= start
        /* For:4 */{ 
          var _Zx *ClaireAny  
          _ = _Zx
          var _Zx_support *ClaireList  
          _Zx_support = _Zl
          _Zx_len := _Zx_support.Length()
          for i_it := 0; i_it < _Zx_len; i_it++ { 
            _Zx = _Zx_support.At(i_it)
            ClEnv.Push(_Zx.ToEID())
            /* loop-5 */} 
          /* For-4 */} 
        F_stack_add(self.Dimension)
        /* LetE:4 */{ 
          var val EID 
          val = EVAL(self.Body)
          /* ERROR PROTECTION INSERTED (val-Result) */
          if ErrorIn(val) {Result = val
          } else {
          ClEnv.Base= retour
          ClEnv.Index= start
          Result = val}
          /* LetE-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: apply @ lambda (throw: true) 
func E_apply_lambda (self EID,_Zl EID) EID { 
    return /*(sm for apply @ lambda= EID)*/ F_apply_lambda(ToLambda(OBJ(self)),ToList(OBJ(_Zl)) )} 
  
/* {1} OPT.The go function for: call(self:lambda,l:listargs) [] */
func F_call_lambda2 (self *ClaireLambda ,l *ClaireList ) EID { 
    var Result EID 
    Result = F_apply_lambda(self,ToList(l.Id()))
    return Result} 
  
// The EID go function for: call @ lambda (throw: true) 
func E_call_lambda2 (self EID,l EID) EID { 
    return /*(sm for call @ lambda= EID)*/ F_call_lambda2(ToLambda(OBJ(self)),ToList(OBJ(l)) )} 
  
// printing a lambda
//
/* {1} OPT.The go function for: self_print(self:lambda) [] */
func F_self_print_lambda_Language (self *ClaireLambda ) EID { 
    var Result EID 
    PRINC("lambda[(")
    Result = F_ppvariable_list(self.Vars)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("),")
    Result = F_lbreak_integer(1)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    Result = Core.F_CALL(C_print,ARGS(self.Body.ToEID()))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    Core.C_pretty.Index = (Core.C_pretty.Index-1)
    PRINC("]")
    Result = EVOID
    }}}
    return Result} 
  
// The EID go function for: self_print @ lambda (throw: true) 
func E_self_print_lambda_Language (self EID) EID { 
    return /*(sm for self_print @ lambda= EID)*/ F_self_print_lambda_Language(ToLambda(OBJ(self)) )} 
  
// lambda! and flexical_build communicate via a global_variable, which
// however is only used in this file (and also by cfile :-) ):
//
// creating a lambda from an instruction and a list of variables
/* {1} OPT.The go function for: iClaire/lambda!(lvar:list,self:any) [] */
func F_lambda_I_list (lvar *ClaireList ,self *ClaireAny ) EID { 
    var Result EID 
    C__starvariable_index_star.Value = MakeInteger(0).Id()
    /* For:2 */{ 
      var v *ClaireAny  
      _ = v
      var v_support *ClaireList  
      v_support = lvar
      v_len := v_support.Length()
      for i_it := 0; i_it < v_len; i_it++ { 
        v = v_support.At(i_it)
        To_Variable(v).Index = ToInteger(C__starvariable_index_star.Value).Value
        v.Isa = C_Variable
        C__starvariable_index_star.Value = MakeInteger((ToInteger(C__starvariable_index_star.Value).Value+1)).Id()
        /* loop-3 */} 
      /* For-2 */} 
    /* Let:2 */{ 
      var corps *ClaireAny  
      /* noccur = 1 */
      var corps_try00393 EID 
      corps_try00393 = F_lexical_build_any(self,lvar,ToInteger(C__starvariable_index_star.Value).Value)
      /* ERROR PROTECTION INSERTED (corps-Result) */
      if ErrorIn(corps_try00393) {Result = corps_try00393
      } else {
      corps = ANY(corps_try00393)
      /* Let:3 */{ 
        var resultat *ClaireLambda   = ToLambda(new(ClaireLambda).Is(C_lambda))
        /* noccur = 4 */
        resultat.Vars = lvar
        resultat.Body = corps
        resultat.Dimension = ToInteger(C__starvariable_index_star.Value).Value
        Result = EID{resultat.Id(),0}
        /* Let-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: iClaire/lambda! @ list (throw: true) 
func E_lambda_I_list (lvar EID,self EID) EID { 
    return /*(sm for iClaire/lambda! @ list= EID)*/ F_lambda_I_list(ToList(OBJ(lvar)),ANY(self) )} 
  
// Give to each lexical variable its right position in the stack.
// We look for a named object or an unbound symbol to replace by a lexical
// variable.
// The number of variables is kept in the global_variable *variable_index*.
// On entry, n need not be equal to size(lvar) (see [case ...instruction]).
//
/* {1} OPT.The go function for: iClaire/lexical_build(self:any,lvar:list,n:integer) [] */
func F_lexical_build_any (self *ClaireAny ,lvar *ClaireList ,n int) EID { 
    var Result EID 
    if ((self.Isa.IsIn(C_thing) == CTRUE) || 
        (self.Isa.IsIn(C_unbound_symbol) == CTRUE)) /* If:2 */{ 
      Result = F_lexical_change_any(self,lvar)
      } else {
      if (self.Isa.IsIn(C_Variable) == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var g0040 *ClaireVariable   = To_Variable(self)
          /* noccur = 3 */
          if (Equal(MakeInteger(g0040.Index).Id(),CNULL) == CTRUE) /* If:5 */{ 
            Result = ToException(Core.C_general_error.Make(MakeString("[145] the symbol ~A is unbound").Id(),MakeConstantList(g0040.Pname.Id()).Id())).Close()
            } else {
            Result = EID{CFALSE.Id(),0}
            /* If-5 */} 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          Result = EID{g0040.Id(),0}
          }
          /* Let-4 */} 
        /* If!3 */}  else if (self.Isa.IsIn(C_Call) == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var g0041 *Call   = To_Call(self)
          /* noccur = 6 */
          /* Let:5 */{ 
            var s *ClaireAny  
            /* noccur = 2 */
            var s_try00456 EID 
            s_try00456 = F_lexical_change_any(g0041.Selector.Id(),lvar)
            /* ERROR PROTECTION INSERTED (s-Result) */
            if ErrorIn(s_try00456) {Result = s_try00456
            } else {
            s = ANY(s_try00456)
            Result = F_lexical_build_any(g0041.Args.Id(),lvar,n)
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            if (g0041.Selector.Id() != s) /* If:6 */{ 
              g0041.Selector = Core.C_call
              /* update:7 */{ 
                var va_arg1 *Call  
                var va_arg2 *ClaireList  
                va_arg1 = g0041
                va_arg2 = F_cons_any(s,g0041.Args)
                /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                va_arg1.Args = va_arg2
                Result = EID{va_arg2.Id(),0}
                /* update-7 */} 
              } else {
              Result = EID{CFALSE.Id(),0}
              /* If-6 */} 
            }
            }
            /* Let-5 */} 
          /* Let-4 */} 
        /* If!3 */}  else if (self.Isa.IsIn(C_Instruction) == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var g0042 *ClaireInstruction   = To_Instruction(self)
          /* noccur = 4 */
          /* Let:5 */{ 
            var _Ztype *ClaireClass   = g0042.Isa
            /* noccur = 2 */
            if (C_Instruction_with_var.Descendents.Contain_ask(_Ztype.Id()) == CTRUE) /* If:6 */{ 
              Result = Core.F_put_property2(C_mClaire_index,ToObject(OBJ(Core.F_CALL(C_var,ARGS(EID{g0042.Id(),0})))),MakeInteger(n).Id())
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              n = (n+1)
              if (n > ToInteger(C__starvariable_index_star.Value).Value) /* If:7 */{ 
                var v_gassign8 *ClaireAny  
                v_gassign8 = MakeInteger(n).Id()
                C__starvariable_index_star.Value = v_gassign8
                Result = v_gassign8.ToEID()
                } else {
                Result = EID{CFALSE.Id(),0}
                /* If-7 */} 
              }
              } else {
              Result = EID{CFALSE.Id(),0}
              /* If-6 */} 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            /* For:6 */{ 
              var s *ClaireAny  
              _ = s
              Result= EID{CFALSE.Id(),0}
              for _,s = range(_Ztype.Slots.ValuesO())/* loop:7 */{ 
                var void_try8 EID 
                _ = void_try8
                /* Let:8 */{ 
                  var x *ClaireAny   = Core.F_get_slot(ToSlot(s),ToObject(g0042.Id()))
                  /* noccur = 4 */
                  if (((x.Isa.IsIn(C_thing) == CTRUE) || 
                        (x.Isa.IsIn(C_unbound_symbol) == CTRUE)) && 
                      (ToRestriction(s).Range.Id() == C_any.Id())) /* If:9 */{ 
                    /* Let:10 */{ 
                      var g0046UU *ClaireAny  
                      /* noccur = 1 */
                      var g0046UU_try004711 EID 
                      g0046UU_try004711 = F_lexical_change_any(x,lvar)
                      /* ERROR PROTECTION INSERTED (g0046UU-void_try8) */
                      if ErrorIn(g0046UU_try004711) {void_try8 = g0046UU_try004711
                      } else {
                      g0046UU = ANY(g0046UU_try004711)
                      void_try8 = Core.F_put_slot(ToSlot(s),ToObject(g0042.Id()),g0046UU).ToEID()
                      }
                      /* Let-10 */} 
                    } else {
                    void_try8 = F_lexical_build_any(x,lvar,n)
                    /* If-9 */} 
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (void_try8-Result) */
                if ErrorIn(void_try8) {Result = void_try8
                Result = void_try8
                break
                } else {
                }
                /* loop-7 */} 
              /* For-6 */} 
            }
            /* Let-5 */} 
          /* Let-4 */} 
        /* If!3 */}  else if (self.Isa.IsIn(C_list) == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var g0043 *ClaireList   = ToList(self)
          /* noccur = 3 */
          /* Let:5 */{ 
            var _Zn int  = g0043.Length()
            /* noccur = 5 */
            Result= EID{CFALSE.Id(),0}
            for (_Zn > 0) /* while:6 */{ 
              var void_try7 EID 
              _ = void_try7
              { 
              /* Let:7 */{ 
                var x *ClaireAny   = g0043.At(_Zn-1)
                /* noccur = 4 */
                if ((x.Isa.IsIn(C_thing) == CTRUE) || 
                    (x.Isa.IsIn(C_unbound_symbol) == CTRUE)) /* If:8 */{ 
                  /* Let:9 */{ 
                    var g0048UU *ClaireAny  
                    /* noccur = 1 */
                    var g0048UU_try004910 EID 
                    g0048UU_try004910 = F_lexical_change_any(x,lvar)
                    /* ERROR PROTECTION INSERTED (g0048UU-void_try7) */
                    if ErrorIn(g0048UU_try004910) {void_try7 = g0048UU_try004910
                    } else {
                    g0048UU = ANY(g0048UU_try004910)
                    void_try7 = ToArray(g0043.Id()).NthPut(_Zn,g0048UU).ToEID()
                    }
                    /* Let-9 */} 
                  } else {
                  void_try7 = F_lexical_build_any(x,lvar,n)
                  /* If-8 */} 
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (void_try7-void_try7) */
              if ErrorIn(void_try7) {Result = void_try7
              break
              } else {
              _Zn = (_Zn-1)
              }
              /* while-6 */} 
            }
            /* Let-5 */} 
          /* Let-4 */} 
        } else {
        Result = EID{CNIL.Id(),0}
        /* If-3 */} 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = self.ToEID()
      }
      /* If-2 */} 
    return Result} 
  
// The EID go function for: iClaire/lexical_build @ any (throw: true) 
func E_lexical_build_any (self EID,lvar EID,n EID) EID { 
    return /*(sm for iClaire/lexical_build @ any= EID)*/ F_lexical_build_any(ANY(self),ToList(OBJ(lvar)),INT(n) )} 
  
/* {1} OPT.The go function for: iClaire/lexical_change(self:any,lvar:list) [] */
func F_lexical_change_any (self *ClaireAny ,lvar *ClaireList ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var rep *ClaireAny   = self
      /* noccur = 2 */
      /* Let:3 */{ 
        var _Zname *ClaireSymbol  
        /* noccur = 1 */
        var _Zname_try00524 EID 
        if (self.Isa.IsIn(C_Variable) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var g0050 *ClaireVariable   = To_Variable(self)
            /* noccur = 1 */
            _Zname_try00524 = EID{g0050.Pname.Id(),0}
            /* Let-5 */} 
          } else {
          _Zname_try00524 = F_extract_symbol_any(self)
          /* If-4 */} 
        /* ERROR PROTECTION INSERTED (_Zname-Result) */
        if ErrorIn(_Zname_try00524) {Result = _Zname_try00524
        } else {
        _Zname = ToSymbol(OBJ(_Zname_try00524))
        /* For:4 */{ 
          var x *ClaireAny  
          _ = x
          var x_support *ClaireList  
          x_support = lvar
          x_len := x_support.Length()
          for i_it := 0; i_it < x_len; i_it++ { 
            x = x_support.At(i_it)
            if (To_Variable(x).Pname.Id() == _Zname.Id()) /* If:6 */{ 
              rep = x
              /* If-6 */} 
            /* loop-5 */} 
          /* For-4 */} 
        Result = rep.ToEID()
        }
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: iClaire/lexical_change @ any (throw: true) 
func E_lexical_change_any (self EID,lvar EID) EID { 
    return /*(sm for iClaire/lexical_change @ any= EID)*/ F_lexical_change_any(ANY(self),ToList(OBJ(lvar)) )} 
  
// *******************************************************************
// *       Part 3: functions for lattice_set instantiation           *
// *******************************************************************
// close is the basic method called by an instantiation.
// Once the indexed list is built, we never call it again.
//
/* {1} OPT.The go function for: close(self:class) [] */
func F_close_class (self *ClaireClass ) *ClaireClass  { 
    // use function body compiling 
return  self
    } 
  
// The EID go function for: close @ class (throw: false) 
func E_close_class (self EID) EID { 
    return EID{/*(sm for close @ class= class)*/ F_close_class(ToClass(OBJ(self)) ).Id(),0}} 
  
// Extract the symbol associated with self.
// This is useful e.g. when using read() (read@port, read@string).
//
/* {1} OPT.The go function for: iClaire/extract_symbol(self:any) [] */
func F_extract_symbol_any (self *ClaireAny ) EID { 
    var Result EID 
    if (self.Isa.IsIn(C_unbound_symbol) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0053 *ClaireUnboundSymbol   = ToUnboundSymbol(self)
        /* noccur = 1 */
        Result = EID{g0053.Name.Id(),0}
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(C_thing) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0054 *ClaireThing   = ToThing(self)
        /* noccur = 1 */
        Result = EID{g0054.Name.Id(),0}
        /* Let-3 */} 
      /* If!2 */}  else if (C_class.Id() == self.Isa.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var g0055 *ClaireClass   = ToClass(self)
        /* noccur = 1 */
        Result = EID{g0055.Name.Id(),0}
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(C_symbol) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0056 *ClaireSymbol   = ToSymbol(self)
        /* noccur = 1 */
        Result = EID{g0056.Id(),0}
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(C_Variable) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0057 *ClaireVariable   = To_Variable(self)
        /* noccur = 1 */
        Result = EID{g0057.Pname.Id(),0}
        /* Let-3 */} 
      /* If!2 */}  else if (C_boolean.Id() == self.Isa.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var g0058 *ClaireBoolean   = ToBoolean(self)
        /* noccur = 1 */
        if (g0058 == CTRUE) /* If:4 */{ 
          Result = EID{Core.F_symbol_I_string2(MakeString("true")).Id(),0}
          } else {
          Result = EID{Core.F_symbol_I_string2(MakeString("nil")).Id(),0}
          /* If-4 */} 
        /* Let-3 */} 
      } else {
      Result = ToException(Core.C_general_error.Make(MakeString("[147] a name cannot be made from ~S").Id(),MakeConstantList(self).Id())).Close()
      /* If-2 */} 
    return Result} 
  
// The EID go function for: iClaire/extract_symbol @ any (throw: true) 
func E_extract_symbol_any (self EID) EID { 
    return /*(sm for iClaire/extract_symbol @ any= EID)*/ F_extract_symbol_any(ANY(self) )} 
  
// we must be sure that the selector (in a has statement or in a message)
// is a property.
//
/* {1} OPT.The go function for: iClaire/make_a_property(self:any) [] */
func F_make_a_property_any (self *ClaireAny ) EID { 
    var Result EID 
    if (self.Isa.IsIn(Core.C_global_variable) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0060 *Core.GlobalVariable   = Core.ToGlobalVariable(self)
        /* noccur = 1 */
        Result = F_make_a_property_any(g0060.Value)
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(C_property) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0061 *ClaireProperty   = ToProperty(self)
        /* noccur = 1 */
        Result = EID{g0061.Id(),0}
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(C_symbol) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0062 *ClaireSymbol   = ToSymbol(self)
        /* noccur = 3 */
        /* Let:4 */{ 
          var x *ClaireAny   = g0062.Get()
          /* noccur = 4 */
          if (x.Isa.IsIn(C_property) == CTRUE) /* If:5 */{ 
            /* Let:6 */{ 
              var g0063 *ClaireProperty   = ToProperty(x)
              /* noccur = 1 */
              Result = F_make_a_property_any(g0063.Id())
              /* Let-6 */} 
            /* If!5 */}  else if (x.Isa.IsIn(Core.C_global_variable) == CTRUE) /* If:5 */{ 
            /* Let:6 */{ 
              var g0064 *Core.GlobalVariable   = Core.ToGlobalVariable(x)
              /* noccur = 1 */
              Result = F_make_a_property_any(g0064.Value)
              /* Let-6 */} 
            } else {
            /* Let:6 */{ 
              var p *ClaireProperty  
              /* noccur = 5 */
              var p_try00687 EID 
              p_try00687 = new(ClaireProperty).IsNamed(C_property,g0062).ToEID()
              /* ERROR PROTECTION INSERTED (p-Result) */
              if ErrorIn(p_try00687) {Result = p_try00687
              } else {
              p = ToProperty(OBJ(p_try00687))
              p.Comment = g0062.String_I()
              p.Domain = ToType(C_any.Id())
              p.Range = ToType(C_any.Id())
              Result = EID{p.Id(),0}
              }
              /* Let-6 */} 
            /* If-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(C_unbound_symbol) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0066 *ClaireUnboundSymbol   = ToUnboundSymbol(self)
        /* noccur = 1 */
        Result = F_make_a_property_any(g0066.Name.Id())
        /* Let-3 */} 
      } else {
      Result = ToException(Core.C_general_error.Make(MakeString("[148] Wrong selector: ~S, cannot make a property\n").Id(),MakeConstantList(self).Id())).Close()
      /* If-2 */} 
    return Result} 
  
// The EID go function for: iClaire/make_a_property @ any (throw: true) 
func E_make_a_property_any (self EID) EID { 
    return /*(sm for iClaire/make_a_property @ any= EID)*/ F_make_a_property_any(ANY(self) )} 
  
// *********************************************************************
// *  Part 4: Pretty printing                                          *
// *********************************************************************
// debug
// create a line break
// if the pretty mode is here ... 
//    (1) pbreak = true means that we create a new line (whatever the length)
//    (2) break = false => we generate a much too far exception
/* {1} OPT.The go function for: lbreak(_CL_obj:void) [] */
func F_lbreak_void () EID { 
    var Result EID 
    if (Core.C_pretty.Pprint == CTRUE) /* If:2 */{ 
      if (Core.C_pretty.Pbreak == CTRUE) /* If:3 */{ 
        PRINC("\n")
        Result = F_put_buffer_void()
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        Result = F_indent_integer(Core.C_pretty.Index).ToEID()
        }
        /* If!3 */}  else if (Core.F_buffer_length_void() > Core.C_pretty.Width) /* If:3 */{ 
        Result = ToException(new(Core.MuchTooFar).Is(Core.C_much_too_far)).Close()
        } else {
        Result = EID{CFALSE.Id(),0}
        /* If-3 */} 
      } else {
      Result = EID{CFALSE.Id(),0}
      /* If-2 */} 
    return Result} 
  
// The EID go function for: lbreak @ void (throw: true) 
func E_lbreak_void (_CL_obj EID) EID { 
    return /*(sm for lbreak @ void= EID)*/ F_lbreak_void( )} 
  
/* {1} OPT.The go function for: put_buffer(_CL_obj:void) [] */
func F_put_buffer_void () EID { 
    var Result EID 
    /* Let:2 */{ 
      var buffer *ClaireString  
      /* noccur = 1 */
      var buffer_try00693 EID 
      buffer_try00693 = Core.F_end_of_string_void()
      /* ERROR PROTECTION INSERTED (buffer-Result) */
      if ErrorIn(buffer_try00693) {Result = buffer_try00693
      } else {
      buffer = ToString(OBJ(buffer_try00693))
      F_princ_string(buffer)
      Core.F_print_in_string_void()
      Result = EID{CEMPTY.Id(),0}
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: put_buffer @ void (throw: true) 
func E_put_buffer_void (_CL_obj EID) EID { 
    return /*(sm for put_buffer @ void= EID)*/ F_put_buffer_void( )} 
  
/* {1} OPT.The go function for: checkfar(_CL_obj:void) [] */
func F_checkfar_void () EID { 
    var Result EID 
    if ((Core.C_pretty.Pprint == CTRUE) && 
        ((Core.C_pretty.Pbreak != CTRUE) && 
          (Core.F_buffer_length_void() > Core.C_pretty.Width))) /* If:2 */{ 
      Result = ToException(new(Core.MuchTooFar).Is(Core.C_much_too_far)).Close()
      } else {
      Result = EID{CFALSE.Id(),0}
      /* If-2 */} 
    return Result} 
  
// The EID go function for: checkfar @ void (throw: true) 
func E_checkfar_void (_CL_obj EID) EID { 
    return /*(sm for checkfar @ void= EID)*/ F_checkfar_void( )} 
  
/* {1} OPT.The go function for: lbreak(n:integer) [] */
func F_lbreak_integer (n int) EID { 
    var Result EID 
    Core.C_pretty.Index = (Core.C_pretty.Index+n)
    Result = F_lbreak_void()
    return Result} 
  
// The EID go function for: lbreak @ integer (throw: true) 
func E_lbreak_integer (n EID) EID { 
    return /*(sm for lbreak @ integer= EID)*/ F_lbreak_integer(INT(n) )} 
  
// indentation
//
/* {1} OPT.The go function for: indent(limit:integer) [] */
func F_indent_integer (limit int) *ClaireAny  { 
    // procedure body with s =  
var Result *ClaireAny  
    /* Let:2 */{ 
      var x int  = Core.F_buffer_length_void()
      /* noccur = 3 */
      Result= CFALSE.Id()
      for (x < limit) /* while:3 */{ 
        PRINC(" ")
        x = (x+1)
        /* while-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: indent @ integer (throw: false) 
func E_indent_integer (limit EID) EID { 
    return /*(sm for indent @ integer= any)*/ F_indent_integer(INT(limit) ).ToEID()} 
  
// sets the current_level
/* {1} OPT.The go function for: set_level(_CL_obj:void) [] */
func F_set_level_void ()  { 
    // procedure body with s =  
Core.C_pretty.Index = (Core.F_buffer_length_void()-1)
    } 
  
// The EID go function for: set_level @ void (throw: false) 
func E_set_level_void (_CL_obj EID) EID { 
    /*(sm for set_level @ void= void)*/ F_set_level_void( )
    return EVOID} 
  
/* {1} OPT.The go function for: set_level(n:integer) [] */
func F_set_level_integer (n int)  { 
    // procedure body with s =  
F_set_level_void()
    Core.C_pretty.Index = (Core.C_pretty.Index+n)
    } 
  
// The EID go function for: set_level @ integer (throw: false) 
func E_set_level_integer (n EID) EID { 
    /*(sm for set_level @ integer= void)*/ F_set_level_integer(INT(n) )
    return EVOID} 
  
// prints a list as a box in character zone [start, finish], s is the separator (",")
// pbreak = true means that we will print step by step; false => try to add to current place
// if impossible of if pbreak = false, we will switch to printl
// the tricky part is that this method can generate a too far error
/* {1} OPT.The go function for: printbox(self:list,start:integer,finish:integer,s:string) [] */
func F_Language_printbox_list1 (self *ClaireList ,start int,finish int,s *ClaireString ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var i int  = 1
      /* noccur = 9 */
      /* Let:3 */{ 
        var startline *ClaireBoolean   = CTRUE
        /* noccur = 3 */
        /* Let:4 */{ 
          var n int  = self.Length()
          /* noccur = 3 */
          /* Let:5 */{ 
            var _Zl int  = Core.C_pretty.Index
            /* noccur = 2 */
            Core.C_pretty.Index = start
            if ((Core.C_pretty.Pprint != CTRUE) || 
                ((Core.F_short_enough_integer((start+10)) != CTRUE) && 
                    (Core.C_pretty.Pbreak == CTRUE))) /* If:6 */{ 
              Result = F_printl_list(self,s)
              /* If!6 */}  else if (Core.C_pretty.Pbreak != CTRUE) /* If:6 */{ 
              Result = F_printl_list(self,s)
              } else {
              Result= EID{CFALSE.Id(),0}
              for (i <= n) /* while:7 */{ 
                var void_try8 EID 
                _ = void_try8
                { 
                for (Core.F_buffer_length_void() < start) /* while:8 */{ 
                  PRINC(" ")
                  /* while-8 */} 
                /* Let:8 */{ 
                  var idx int  = Core.F_buffer_length_void()
                  /* noccur = 1 */
                  h_index := ClEnv.Index /* Handle */
                  h_base := ClEnv.Base
                  Core.C_pretty.Pbreak = CFALSE
                  void_try8 = F_printexp_any(self.At(i-1),CTRUE)
                  /* ERROR PROTECTION INSERTED (void_try8-void_try8) */
                  if !ErrorIn(void_try8) {
                  /* update:9 */{ 
                    var va_arg1 *Core.PrettyPrinter  
                    var va_arg2 *ClaireBoolean  
                    va_arg1 = Core.C_pretty
                    va_arg2 = CTRUE
                    /* ---------- now we compile update mClaire/pbreak(va_arg1) := va_arg2 ------- */
                    va_arg1.Pbreak = va_arg2
                    void_try8 = EID{va_arg2.Id(),0}
                    /* update-9 */} 
                  }
                  if ErrorIn(void_try8) && ToType(Core.C_much_too_far.Id()).Contains(ANY(void_try8)) == CTRUE { 
                    /* s=EID */ClEnv.Index = h_index
                    ClEnv.Base = h_base
                    Core.C_pretty.Pbreak = CTRUE
                    /* update:10 */{ 
                      var va_arg1 *Core.PrettyPrinter  
                      var va_arg2 int 
                      va_arg1 = Core.C_pretty
                      va_arg2 = start
                      /* ---------- now we compile update mClaire/index(va_arg1) := va_arg2 ------- */
                      va_arg1.Index = va_arg2
                      void_try8 = EID{C__INT,IVAL(va_arg2)}
                      /* update-10 */} 
                    } 
                  /* ERROR PROTECTION INSERTED (void_try8-void_try8) */
                  if ErrorIn(void_try8) {Result = void_try8
                  break
                  } else {
                  if (i != n) /* If:9 */{ 
                    F_princ_string(s)
                    /* If-9 */} 
                  if (Core.F_buffer_length_void() < finish) /* If:9 */{ 
                    i = (i+1)
                    startline = CFALSE
                    void_try8 = EID{startline.Id(),0}
                    } else {
                    Core.F_buffer_set_length_integer(idx)
                    if (startline != CTRUE) /* If:10 */{ 
                      void_try8 = F_lbreak_void()
                      /* ERROR PROTECTION INSERTED (void_try8-void_try8) */
                      if ErrorIn(void_try8) {Result = void_try8
                      break
                      } else {
                      startline = CTRUE
                      void_try8 = EID{startline.Id(),0}
                      }
                      } else {
                      F_set_level_void()
                      Core.C_pretty.Index = (Core.C_pretty.Index+1)
                      void_try8 = F_printexp_any(self.At(i-1),CTRUE)
                      /* ERROR PROTECTION INSERTED (void_try8-void_try8) */
                      if ErrorIn(void_try8) {Result = void_try8
                      break
                      } else {
                      Core.C_pretty.Index = _Zl
                      if (i != n) /* If:11 */{ 
                        F_princ_string(s)
                        void_try8 = F_lbreak_void()
                        /* ERROR PROTECTION INSERTED (void_try8-void_try8) */
                        if ErrorIn(void_try8) {Result = void_try8
                        break
                        } else {
                        }
                        } else {
                        void_try8 = EID{CFALSE.Id(),0}
                        /* If-11 */} 
                      /* ERROR PROTECTION INSERTED (void_try8-void_try8) */
                      if ErrorIn(void_try8) {Result = void_try8
                      break
                      } else {
                      i = (i+1)
                      void_try8 = EID{C__INT,IVAL(i)}
                      }}
                      /* If-10 */} 
                    /* ERROR PROTECTION INSERTED (void_try8-void_try8) */
                    if ErrorIn(void_try8) {Result = void_try8
                    break
                    } else {
                    }
                    /* If-9 */} 
                  /* ERROR PROTECTION INSERTED (void_try8-void_try8) */
                  if ErrorIn(void_try8) {Result = void_try8
                  break
                  } else {
                  }}
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (void_try8-void_try8) */
                if ErrorIn(void_try8) {Result = void_try8
                break
                } else {
                }
                /* while-7 */} 
              }
              /* If-6 */} 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            Core.C_pretty.Index = _Zl
            Result = EID{CNULL,0}
            }
            /* Let-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: printbox @ list<type_expression>(list, integer, integer, string) (throw: true) 
func E_Language_printbox_list1 (self EID,start EID,finish EID,s EID) EID { 
    return /*(sm for printbox @ list<type_expression>(list, integer, integer, string)= EID)*/ F_Language_printbox_list1(ToList(OBJ(self)),
      INT(start),
      INT(finish),
      ToString(OBJ(s)) )} 
  
// default value of arguments
//
/* {1} OPT.The go function for: printbox(self:list) [] */
func F_Language_printbox_list2 (self *ClaireList ) EID { 
    var Result EID 
    Result = F_Language_printbox_list1(self,Core.F_buffer_length_void(),Core.C_pretty.Width,MakeString(", "))
    return Result} 
  
// The EID go function for: printbox @ list<type_expression>(list) (throw: true) 
func E_Language_printbox_list2 (self EID) EID { 
    return /*(sm for printbox @ list<type_expression>(list)= EID)*/ F_Language_printbox_list2(ToList(OBJ(self)) )} 
  
/* {1} OPT.The go function for: printbox(self:list,s:string) [] */
func F_Language_printbox_list3 (self *ClaireList ,s *ClaireString ) EID { 
    var Result EID 
    Result = F_Language_printbox_list1(self,Core.F_buffer_length_void(),Core.C_pretty.Width,s)
    return Result} 
  
// The EID go function for: printbox @ list<type_expression>(list, string) (throw: true) 
func E_Language_printbox_list3 (self EID,s EID) EID { 
    return /*(sm for printbox @ list<type_expression>(list, string)= EID)*/ F_Language_printbox_list3(ToList(OBJ(self)),ToString(OBJ(s)) )} 
  
// this is a tricky method : first try to print without pretty (box) 
/* {1} OPT.The go function for: printl(self:list,s:string) [] */
func F_printl_list (self *ClaireList ,s *ClaireString ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var f *ClaireBoolean   = CTRUE
      /* noccur = 2 */
      /* Let:3 */{ 
        var b *ClaireBoolean   = Core.C_pretty.Pprint
        /* noccur = 5 */
        Core.C_pretty.Pprint = CFALSE
        h_index := ClEnv.Index /* Handle */
        h_base := ClEnv.Base
        /* For:4 */{ 
          var x *ClaireAny  
          _ = x
          Result= EID{CFALSE.Id(),0}
          var x_support *ClaireList  
          x_support = self
          x_len := x_support.Length()
          for i_it := 0; i_it < x_len; i_it++ { 
            x = x_support.At(i_it)
            var void_try6 EID 
            _ = void_try6
            { 
            if (f == CTRUE) /* If:6 */{ 
              f = CFALSE
              } else {
              F_princ_string(s)
              /* If-6 */} 
            void_try6 = F_printexp_any(x,CTRUE)
            /* ERROR PROTECTION INSERTED (void_try6-void_try6) */
            if ErrorIn(void_try6) {Result = void_try6
            break
            } else {
            if ((b == CTRUE) && 
                ((Core.C_pretty.Pbreak != CTRUE) && 
                  (Core.F_buffer_length_void() > Core.C_pretty.Width))) /* If:6 */{ 
              Core.C_pretty.Pprint = b
              void_try6 = ToException(new(Core.MuchTooFar).Is(Core.C_much_too_far)).Close()
              /* ERROR PROTECTION INSERTED (void_try6-void_try6) */
              if ErrorIn(void_try6) {Result = void_try6
              break
              } else {
              }
              } else {
              void_try6 = EID{CFALSE.Id(),0}
              /* If-6 */} 
            /* ERROR PROTECTION INSERTED (void_try6-void_try6) */
            if ErrorIn(void_try6) {Result = void_try6
            break
            } else {
            }}
            }
            /* loop-5 */} 
          /* For-4 */} 
        if ErrorIn(Result) && ToType(C_system_error.Id()).Contains(ANY(Result)) == CTRUE { 
          /* s=EID */ClEnv.Index = h_index
          ClEnv.Base = h_base
          /* Let:5 */{ 
            var x *ClaireException   = ClEnv.Exception_I
            /* noccur = 2 */
            if ((b == CTRUE) && 
                (ToSystemError(x.Id()).Index == 16)) /* If:6 */{ 
              Core.C_pretty.Pprint = b
              Result = ToException(new(Core.MuchTooFar).Is(Core.C_much_too_far)).Close()
              } else {
              Result = x.Close()
              /* If-6 */} 
            /* Let-5 */} 
          } 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        /* update:4 */{ 
          var va_arg1 *Core.PrettyPrinter  
          var va_arg2 *ClaireBoolean  
          va_arg1 = Core.C_pretty
          va_arg2 = b
          /* ---------- now we compile update mClaire/pprint(va_arg1) := va_arg2 ------- */
          va_arg1.Pprint = va_arg2
          Result = EID{va_arg2.Id(),0}
          /* update-4 */} 
        }
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: printl @ list (throw: true) 
func E_printl_list (self EID,s EID) EID { 
    return /*(sm for printl @ list= EID)*/ F_printl_list(ToList(OBJ(self)),ToString(OBJ(s)) )} 
  
// print bounded prints a bounded expression using ( and )
/* {1} OPT.The go function for: printexp(self:any,comp:boolean) [] */
func F_printexp_any (self *ClaireAny ,comp *ClaireBoolean ) EID { 
    var Result EID 
    var g0074I *ClaireBoolean  
    /* or:2 */{ 
      var v_or2 *ClaireBoolean  
      
      if (self.Isa.IsIn(C_Call) == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var g0073 *Call   = To_Call(self)
          /* noccur = 2 */
          v_or2 = MakeBoolean((g0073.Selector.Isa.IsIn(C_operation) == CTRUE) && (comp != CTRUE) && (g0073.Args.Length() == 2)).Not
          /* Let-4 */} 
        } else {
        v_or2 = CFALSE
        /* If-3 */} 
      if (v_or2 == CTRUE) {g0074I = CTRUE
      } else /* or:3 */{ 
        v_or2 = self.Isa.IsIn(C_Collect)
        if (v_or2 == CTRUE) {g0074I = CTRUE
        } else /* or:4 */{ 
          v_or2 = self.Isa.IsIn(C_Select)
          if (v_or2 == CTRUE) {g0074I = CTRUE
          } else /* or:5 */{ 
            v_or2 = self.Isa.IsIn(C_Definition)
            if (v_or2 == CTRUE) {g0074I = CTRUE
            } else /* or:6 */{ 
              v_or2 = self.Isa.IsIn(C_Construct)
              if (v_or2 == CTRUE) {g0074I = CTRUE
              } else /* or:7 */{ 
                v_or2 = self.Isa.IsIn(C_Do)
                if (v_or2 == CTRUE) {g0074I = CTRUE
                } else /* or:8 */{ 
                  v_or2 = Equal(self,CNULL)
                  if (v_or2 == CTRUE) {g0074I = CTRUE
                  } else /* or:9 */{ 
                    v_or2 = self.Isa.IsIn(C_And)
                    if (v_or2 == CTRUE) {g0074I = CTRUE
                    } else /* or:10 */{ 
                      v_or2 = self.Isa.IsIn(C_primitive)
                      if (v_or2 == CTRUE) {g0074I = CTRUE
                      } else /* or:11 */{ 
                        v_or2 = self.Isa.IsIn(C_Or)
                        if (v_or2 == CTRUE) {g0074I = CTRUE
                        } else /* or:12 */{ 
                          v_or2 = self.Isa.IsIn(C_If)
                          if (v_or2 == CTRUE) {g0074I = CTRUE
                          } else /* or:13 */{ 
                            v_or2 = self.Isa.IsIn(C_restriction)
                            if (v_or2 == CTRUE) {g0074I = CTRUE
                            } else /* or:14 */{ 
                              v_or2 = self.Isa.IsIn(C_unbound_symbol)
                              if (v_or2 == CTRUE) {g0074I = CTRUE
                              } else /* or:15 */{ 
                                v_or2 = self.Isa.IsIn(C_Variable)
                                if (v_or2 == CTRUE) {g0074I = CTRUE
                                } else /* or:16 */{ 
                                  v_or2 = self.Isa.IsIn(C_Instruction).Not
                                  if (v_or2 == CTRUE) {g0074I = CTRUE
                                  } else /* or:17 */{ 
                                    g0074I = CFALSE/* org-17 */} 
                                  /* org-16 */} 
                                /* org-15 */} 
                              /* org-14 */} 
                            /* org-13 */} 
                          /* org-12 */} 
                        /* org-11 */} 
                      /* org-10 */} 
                    /* org-9 */} 
                  /* org-8 */} 
                /* org-7 */} 
              /* org-6 */} 
            /* org-5 */} 
          /* org-4 */} 
        /* org-3 */} 
      /* or-2 */} 
    if (g0074I == CTRUE) /* If:2 */{ 
      Result = Core.F_CALL(C_print,ARGS(self.ToEID()))
      } else {
      /* Let:3 */{ 
        var _Zl int  = Core.C_pretty.Index
        /* noccur = 1 */
        PRINC("(")
        F_set_level_integer(1)
        Result = Core.F_CALL(C_print,ARGS(self.ToEID()))
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC(")")
        Result = EVOID
        }
        {
        /* update:4 */{ 
          var va_arg1 *Core.PrettyPrinter  
          var va_arg2 int 
          va_arg1 = Core.C_pretty
          va_arg2 = _Zl
          /* ---------- now we compile update mClaire/index(va_arg1) := va_arg2 ------- */
          va_arg1.Index = va_arg2
          Result = EID{C__INT,IVAL(va_arg2)}
          /* update-4 */} 
        }
        /* Let-3 */} 
      /* If-2 */} 
    return Result} 
  
// The EID go function for: printexp @ any (throw: true) 
func E_printexp_any (self EID,comp EID) EID { 
    return /*(sm for printexp @ any= EID)*/ F_printexp_any(ANY(self),ToBoolean(OBJ(comp)) )} 
  
// pretty print is using the buffered print (into a string)
/* {1} OPT.The go function for: pretty_print(self:any) [] */
func F_pretty_print_any (self *ClaireAny ) EID { 
    var Result EID 
    Core.F_print_in_string_void()
    Core.C_pretty.Pprint = CTRUE
    Core.C_pretty.Pbreak = CTRUE
    Core.C_pretty.Index = 0
    Result = Core.F_CALL(C_print,ARGS(self.ToEID()))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    Core.C_pretty.Pprint = CFALSE
    /* Let:2 */{ 
      var g0075UU *ClaireString  
      /* noccur = 1 */
      var g0075UU_try00763 EID 
      g0075UU_try00763 = Core.F_end_of_string_void()
      /* ERROR PROTECTION INSERTED (g0075UU-Result) */
      if ErrorIn(g0075UU_try00763) {Result = g0075UU_try00763
      } else {
      g0075UU = ToString(OBJ(g0075UU_try00763))
      F_princ_string(g0075UU)
      Result = EVOID
      }
      /* Let-2 */} 
    }
    return Result} 
  
// The EID go function for: pretty_print @ any (throw: true) 
func E_pretty_print_any (self EID) EID { 
    return /*(sm for pretty_print @ any= EID)*/ F_pretty_print_any(ANY(self) )} 
  
// self_print uses the default boxing
/* {1} OPT.The go function for: self_print(self:list) [] */
func F_self_print_list_Language (self *ClaireList ) EID { 
    var Result EID 
    if (Equal(self.Of().Id(),CEMPTY.Id()) != CTRUE) /* If:2 */{ 
      PRINC("list<")
      Result = Core.F_print_any(self.Of().Id())
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(">")
      Result = EVOID
      }
      } else {
      PRINC("list")
      Result = EVOID
      /* If-2 */} 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("(")
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
    return /*(sm for self_print @ list= EID)*/ F_self_print_list_Language(ToList(OBJ(self)) )} 
  
/* {1} OPT.The go function for: self_print(self:set) [] */
func F_self_print_set_Language (self *ClaireSet ) EID { 
    var Result EID 
    if (Equal(ToList(self.Id()).Of().Id(),CEMPTY.Id()) == CTRUE) /* If:2 */{ 
      PRINC("{")
      Result = F_Language_printbox_list2(self.List_I())
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC("}")
      Result = EVOID
      }
      } else {
      PRINC("set<")
      Result = Core.F_print_any(ToList(self.Id()).Of().Id())
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(">")
      Result = EVOID
      }
      {
      PRINC("(")
      Result = F_Language_printbox_list2(self.List_I())
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(")")
      Result = EVOID
      }
      }
      /* If-2 */} 
    return Result} 
  
// The EID go function for: self_print @ set (throw: true) 
func E_self_print_set_Language (self EID) EID { 
    return /*(sm for self_print @ set= EID)*/ F_self_print_set_Language(ToSet(OBJ(self)) )} 
  
// to remove !
/* {1} OPT.The go function for: self_print(self:tuple) [] */
func F_self_print_tuple_Language (self *ClaireTuple ) EID { 
    var Result EID 
    PRINC("tuple(")
    Result = F_Language_printbox_list2(ToList(self.Id()))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(")")
    Result = EVOID
    }
    return Result} 
  
// The EID go function for: self_print @ tuple (throw: true) 
func E_self_print_tuple_Language (self EID) EID { 
    return /*(sm for self_print @ tuple= EID)*/ F_self_print_tuple_Language(ToTuple(OBJ(self)) )} 
  
// *********************************************************************
// *  Part 5: simple type inference  (class based)                     *
// *********************************************************************
// this is a simple, self-contained, type inference method that mimicks what GO is bound to know
// it is used to check the type safety of the gerenated code in the Generate module and it is used
// in call.cl to produce OFTO (on-the-fly optimization) => see readcall
// s_type =  static type, or stupid_type  (we should remove stupid_t)
/* {1} OPT.The go function for: static_type(self:any) [] */
func F_static_type_any (self *ClaireAny ) EID { 
    var Result EID 
    if (self.Isa.IsIn(C_Variable) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0077 *ClaireVariable   = To_Variable(self)
        /* noccur = 1 */
        Result = EID{g0077.Range.Class_I().Id(),0}
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(Core.C_global_variable) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0078 *Core.GlobalVariable   = Core.ToGlobalVariable(self)
        /* noccur = 2 */
        /* Let:4 */{ 
          var r *ClaireType   = g0078.Range
          /* noccur = 2 */
          if (F_boolean_I_any(r.Id()) == CTRUE) /* If:5 */{ 
            Result = EID{r.Class_I().Id(),0}
            } else {
            Result = EID{g0078.Value.Isa.Id(),0}
            /* If-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(C_And) == CTRUE) /* If:2 */{ 
      Result = EID{C_boolean.Id(),0}
      /* If!2 */}  else if (self.Isa.IsIn(C_Or) == CTRUE) /* If:2 */{ 
      Result = EID{C_boolean.Id(),0}
      /* If!2 */}  else if (self.Isa.IsIn(C_environment) == CTRUE) /* If:2 */{ 
      Result = EID{C_environment.Id(),0}
      /* If!2 */}  else if (self.Isa.IsIn(C_Call_plus) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0082 *Call_plus   = To_Call_plus(self)
        /* noccur = 2 */
        /* Let:4 */{ 
          var p *ClaireProperty   = g0082.Selector
          /* noccur = 2 */
          /* Let:5 */{ 
            var s *ClaireObject  
            /* noccur = 2 */
            var s_try01046 EID 
            /* Let:6 */{ 
              var g0105UU *ClaireClass  
              /* noccur = 1 */
              var g0105UU_try01067 EID 
              g0105UU_try01067 = F_static_type_any(g0082.Args.At(1-1))
              /* ERROR PROTECTION INSERTED (g0105UU-s_try01046) */
              if ErrorIn(g0105UU_try01067) {s_try01046 = g0105UU_try01067
              } else {
              g0105UU = ToClass(OBJ(g0105UU_try01067))
              s_try01046 = EID{Core.F__at_property1(p,g0105UU).Id(),0}
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (s-Result) */
            if ErrorIn(s_try01046) {Result = s_try01046
            } else {
            s = ToObject(OBJ(s_try01046))
            if (C_slot.Id() == s.Isa.Id()) /* If:6 */{ 
              /* Let:7 */{ 
                var g0083 *ClaireSlot   = ToSlot(s.Id())
                /* noccur = 1 */
                Result = EID{g0083.Range.Id(),0}
                /* Let-7 */} 
              } else {
              Result = EID{p.Range.Id(),0}
              /* If-6 */} 
            }
            /* Let-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(C_Call_slot) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0085 *CallSlot   = To_CallSlot(self)
        /* noccur = 1 */
        /* Let:4 */{ 
          var s *ClaireSlot   = g0085.Selector
          /* noccur = 4 */
          /* Let:5 */{ 
            var p *ClaireProperty   = s.Selector
            /* noccur = 1 */
            /* For:6 */{ 
              var s2 *ClaireAny  
              _ = s2
              for _,s2 = range(p.Definition.ValuesO())/* loop:7 */{ 
                if (C_slot.Id() == s2.Isa.Id()) /* If:8 */{ 
                  /* Let:9 */{ 
                    var g0086 *ClaireSlot   = ToSlot(s2)
                    /* noccur = 2 */
                    if (ToType(Core.F_domain_I_restriction(ToRestriction(s.Id())).Id()).Included(ToType(Core.F_domain_I_restriction(ToRestriction(g0086.Id())).Id())) == CTRUE) /* If:10 */{ 
                      s = g0086
                      /* If-10 */} 
                    /* Let-9 */} 
                  /* If-8 */} 
                /* loop-7 */} 
              /* For-6 */} 
            Result = EID{s.Range.Class_I().Id(),0}
            /* Let-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(C_Call_method) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0087 *CallMethod   = To_CallMethod(self)
        /* noccur = 3 */
        /* Let:4 */{ 
          var p *ClaireProperty   = g0087.Arg.Selector
          /* noccur = 1 */
          if (p.Id() == C_nth.Id()) /* If:5 */{ 
            Result = F_Language_static_type_nth_any(g0087.Args.At(1-1))
            } else {
            Result = EID{g0087.Arg.Range.Class_I().Id(),0}
            /* If-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(C_Call) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0088 *Call   = To_Call(self)
        /* noccur = 2 */
        /* Let:4 */{ 
          var p *ClaireProperty   = g0088.Selector
          /* noccur = 2 */
          if (p.Id() == C_nth.Id()) /* If:5 */{ 
            Result = F_Language_static_type_nth_any(g0088.Args.At(1-1))
            } else {
            Result = EID{p.Range.Class_I().Id(),0}
            /* If-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(C_Assign) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0089 *Assign   = To_Assign(self)
        /* noccur = 1 */
        Result = Core.F_CALL(C_Language_s_type,ARGS(g0089.Arg.ToEID()))
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(C_Let) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0090 *Let   = To_Let(self)
        /* noccur = 1 */
        Result = F_static_type_any(g0090.Arg)
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(C_Do) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0091 *Do   = To_Do(self)
        /* noccur = 1 */
        /* Let:4 */{ 
          var g0107UU *ClaireAny  
          /* noccur = 1 */
          var g0107UU_try01085 EID 
          g0107UU_try01085 = Core.F_last_list(g0091.Args)
          /* ERROR PROTECTION INSERTED (g0107UU-Result) */
          if ErrorIn(g0107UU_try01085) {Result = g0107UU_try01085
          } else {
          g0107UU = ANY(g0107UU_try01085)
          Result = F_static_type_any(g0107UU)
          }
          /* Let-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(C_If) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0092 *If   = To_If(self)
        /* noccur = 2 */
        /* Let:4 */{ 
          var g0109UU *ClaireClass  
          /* noccur = 1 */
          var g0109UU_try01115 EID 
          g0109UU_try01115 = F_static_type_any(g0092.Arg)
          /* ERROR PROTECTION INSERTED (g0109UU-Result) */
          if ErrorIn(g0109UU_try01115) {Result = g0109UU_try01115
          } else {
          g0109UU = ToClass(OBJ(g0109UU_try01115))
          /* Let:5 */{ 
            var g0110UU *ClaireClass  
            /* noccur = 1 */
            var g0110UU_try01126 EID 
            g0110UU_try01126 = F_static_type_any(g0092.Other)
            /* ERROR PROTECTION INSERTED (g0110UU-Result) */
            if ErrorIn(g0110UU_try01126) {Result = g0110UU_try01126
            } else {
            g0110UU = ToClass(OBJ(g0110UU_try01126))
            Result = EID{Core.F_meet_class(g0109UU,g0110UU).Id(),0}
            }
            /* Let-5 */} 
          }
          /* Let-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(C_Collect) == CTRUE) /* If:2 */{ 
      Result = EID{C_list.Id(),0}
      /* If!2 */}  else if (self.Isa.IsIn(C_Image) == CTRUE) /* If:2 */{ 
      Result = EID{C_set.Id(),0}
      /* If!2 */}  else if (self.Isa.IsIn(C_Select) == CTRUE) /* If:2 */{ 
      Result = EID{C_set.Id(),0}
      /* If!2 */}  else if (self.Isa.IsIn(C_Lselect) == CTRUE) /* If:2 */{ 
      Result = EID{C_list.Id(),0}
      /* If!2 */}  else if (self.Isa.IsIn(C_List) == CTRUE) /* If:2 */{ 
      Result = EID{C_list.Id(),0}
      /* If!2 */}  else if (self.Isa.IsIn(C_Set) == CTRUE) /* If:2 */{ 
      Result = EID{C_set.Id(),0}
      /* If!2 */}  else if (self.Isa.IsIn(C_Tuple) == CTRUE) /* If:2 */{ 
      Result = EID{C_tuple.Id(),0}
      /* If!2 */}  else if (self.Isa.IsIn(C_Exists) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0100 *Exists   = To_Exists(self)
        /* noccur = 1 */
        if (g0100.Other == CNULL) /* If:4 */{ 
          Result = EID{C_any.Id(),0}
          } else {
          Result = EID{C_boolean.Id(),0}
          /* If-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(C_Definition) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0101 *Definition   = To_Definition(self)
        /* noccur = 1 */
        Result = EID{g0101.Arg.Id(),0}
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(C_Instruction) == CTRUE) /* If:2 */{ 
      Result = EID{C_any.Id(),0}
      } else {
      Result = EID{self.Isa.Id(),0}
      /* If-2 */} 
    return RangeCheck(ToType(C_class.Id()),Result)} 
  
// The EID go function for: static_type @ any (throw: true) 
func E_static_type_any (self EID) EID { 
    return /*(sm for static_type @ any= EID)*/ F_static_type_any(ANY(self) )} 
  
// second order pattern for a very common case l[i] where l:list<X>
/* {1} OPT.The go function for: static_type_nth(x:any) [] */
func F_Language_static_type_nth_any (x *ClaireAny ) EID { 
    var Result EID 
    if (x.Isa.IsIn(C_Variable) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0113 *ClaireVariable   = To_Variable(x)
        /* noccur = 1 */
        /* Let:4 */{ 
          var s *ClaireType   = g0113.Range
          /* noccur = 2 */
          if (s.Isa.IsIn(C_Param) == CTRUE) /* If:5 */{ 
            /* Let:6 */{ 
              var g0114 *ClaireParam   = To_Param(s.Id())
              /* noccur = 2 */
              if (g0114.Params.At(1-1) == C_of.Id()) /* If:7 */{ 
                /* Let:8 */{ 
                  var g0117UU *ClaireAny  
                  /* noccur = 1 */
                  var g0117UU_try01189 EID 
                  g0117UU_try01189 = Core.F_the_type(ToType(g0114.Args.At(1-1)))
                  /* ERROR PROTECTION INSERTED (g0117UU-Result) */
                  if ErrorIn(g0117UU_try01189) {Result = g0117UU_try01189
                  } else {
                  g0117UU = ANY(g0117UU_try01189)
                  Result = EID{ToTypeExpression(g0117UU).Class_I().Id(),0}
                  }
                  /* Let-8 */} 
                } else {
                Result = EID{C_any.Id(),0}
                /* If-7 */} 
              /* Let-6 */} 
            } else {
            Result = EID{C_any.Id(),0}
            /* If-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      } else {
      Result = EID{C_any.Id(),0}
      /* If-2 */} 
    return Result} 
  
// The EID go function for: static_type_nth @ any (throw: true) 
func E_Language_static_type_nth_any (x EID) EID { 
    return /*(sm for static_type_nth @ any= EID)*/ F_Language_static_type_nth_any(ANY(x) )} 
  
// end of file