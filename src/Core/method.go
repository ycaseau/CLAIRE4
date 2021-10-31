/***** CLAIRE Compilation of file /Users/ycaseau/Dropbox/src/clairev4.0/src/meta/method.cl 
         [version 4.0.02 / safety 5] Saturday 10-30-2021 *****/

package Core
import (_ "fmt"
	. "Kernel"
)

//+-------------------------------------------------------------+
//| CLAIRE                                                      |
//| method.cl                                                   |
//| Copyright (C) 1994 - 2021 Yves Caseau. All Rights Reserved  |
//| cf. copyright info in file  object.cl: about()              |
//+-------------------------------------------------------------+
// ---------------------------------------------------------------------
// This file contains the reflective description of the most primitive
// CLAIRE kernel: the embryo of the class hierarchy and a set of methods
// to read/write objects and to evaluate messages.
// ---------------------------------------------------------------------
// CLAIRE4 revision : more objects are defined in Kernel
// *********************************************************************
// *  Table of contents                                                *
// *      Part 1: Lambda & Methods Evaluation                          *
// *      Part 2: Update methods                                       *
// *      Part 3: Management of definition(p)                          *
// *      Part 4: Matching Methods                                     *
// *********************************************************************
// catch what was missed in Kernel
// complete instanciation
/* {1} OPT.The go function for: close(r:slot) [] */
func F_close_slot (r *ClaireSlot ) *ClaireSlot  { 
    // use function body compiling 
F_insert_definition_property(r.Selector,ToRestriction(r.Id()))
    return  r
    } 
  
// The EID go function for: close @ slot (throw: false) 
func E_close_slot (r EID) EID { 
    return EID{/*(sm for close @ slot= slot)*/ F_close_slot(ToSlot(OBJ(r)) ).Id(),0}} 
  
/* {1} OPT.The go function for: close(r:method) [] */
func F_close_method (r *ClaireMethod ) *ClaireMethod  { 
    // use function body compiling 
F_insert_definition_property(r.Selector,ToRestriction(r.Id()))
    return  r
    } 
  
// The EID go function for: close @ method (throw: false) 
func E_close_method (r EID) EID { 
    return EID{/*(sm for close @ method= method)*/ F_close_method(ToMethod(OBJ(r)) ).Id(),0}} 
  
// Claire 4: introduce the capacity to set the comment automatically at compile time
/* {1} OPT.The go function for: attach(r:method,s:string) [] */
func F_attach_method (r *ClaireMethod ,s *ClaireString ) *ClaireMethod  { 
    // use function body compiling 
r.Comment = F_append_string(MakeString("defined in file "),s)
    return  F_close_method(r)
    } 
  
// The EID go function for: attach @ method (throw: false) 
func E_attach_method (r EID,s EID) EID { 
    return EID{/*(sm for attach @ method= method)*/ F_attach_method(ToMethod(OBJ(r)),ToString(OBJ(s)) ).Id(),0}} 
  
// *********************************************************************
// *      Part 1: Lambda & Methods Evaluation                          *
// *********************************************************************
// Lambda is now defined in Kernel
// explicit definition of the functions that are used in method [to avoid out_of-place implicit definitions]
// read :: property() moved to Kernel
// the eval_message is a method that tells how to handle the message.
// it used to be distributed in CLAIRE (so that it was extensible) and each
// definition was called the behavior of a kind of restriction
// int? tells us if this is an interpreted message
/* {1} OPT.The go function for: eval_message(self:property,r:object,start:integer,int?:boolean) [] */
func F_eval_message_property (self *ClaireProperty ,r *ClaireObject ,start int,int_ask *ClaireBoolean ) EID { 
    var Result EID 
    if (r.Isa.Id() == C_method.Id()) /* If:2 */{ 
      if (ClEnv.Debug_I != -1) /* If:3 */{ 
        Result = F_execute_method(ToMethod(r.Id()),start,int_ask)
        } else {
        /* Let:4 */{ 
          var m *ClaireMethod   = ToMethod(r.Id())
          /* noccur = 4 */
          if (m.Formula.Id() != CNULL) /* If:5 */{ 
            /* Let:6 */{ 
              var retour int  = ClEnv.Base
              /* noccur = 1 */
              ClEnv.Base= start
              F_stack_add(m.Formula.Dimension)
              /* LetE:7 */{ 
                var val EID 
                val = EVAL(m.Formula.Body)
                /* ERROR PROTECTION INSERTED (val-Result) */
                if ErrorIn(val) {Result = val
                } else {
                ClEnv.Base= retour
                ClEnv.Index= start
                Result = val}
                /* LetE-7 */} 
              /* Let-6 */} 
            } else {
            Result = F_stack_apply_function(m.Functional,start,ClEnv.Index)
            /* If-5 */} 
          /* Let-4 */} 
        /* If-3 */} 
      /* If!2 */}  else if ((r.Id().Isa.Id() == C_slot.Id()) && 
        (ClEnv.Index == (start+1))) /* If:2 */{ 
      /* Let:3 */{ 
        var val *ClaireAny   = F_get_slot(ToSlot(r.Id()),ToObject(OBJ(ClEnv.EvalStack[start])))
        /* noccur = 4 */
        if ((val == CNULL) && 
            (ToRestriction(r.Id()).Range.Contains(val) != CTRUE)) /* If:4 */{ 
          Result = ToException(C_read_slot_error.Make(ANY(ClEnv.EvalStack[start]),self.Id())).Close()
          } else {
          Result = EID{CFALSE.Id(),0}
          /* If-4 */} 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        ClEnv.Index= start
        /* Let:4 */{ 
          var n int  = ClEnv.Trace_I
          /* noccur = 2 */
          if ((n > 0) && 
              ((self.Trace_I+ClEnv.Verbose) > 4)) /* If:5 */{ 
            ClEnv.Trace_I = 0
            PRINC("read: ")
            Result = F_print_any(self.Id())
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC("(")
            Result = F_CALL(C_print,ARGS(ClEnv.EvalStack[start]))
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC(") = ")
            Result = F_CALL(C_print,ARGS(val.ToEID()))
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC("\n")
            Result = EVOID
            }}}
            {
            /* update:6 */{ 
              var va_arg1 *ClaireEnvironment  
              var va_arg2 int 
              va_arg1 = ClEnv
              va_arg2 = n
              /* ---------- now we compile update trace!(va_arg1) := va_arg2 ------- */
              va_arg1.Trace_I = va_arg2
              Result = EID{C__INT,IVAL(va_arg2)}
              /* update-6 */} 
            }
            } else {
            Result = EID{CFALSE.Id(),0}
            /* If-5 */} 
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        Result = val.ToEID()
        }}
        /* Let-3 */} 
      } else {
      Result = F_noeval_message_property2(self,start)
      /* If-2 */} 
    return Result} 
  
// The EID go function for: eval_message @ property (throw: true) 
func E_eval_message_property (self EID,r EID,start EID,int_ask EID) EID { 
    return /*(sm for eval_message @ property= EID)*/ F_eval_message_property(ToProperty(OBJ(self)),
      ToObject(OBJ(r)),
      INT(start),
      ToBoolean(OBJ(int_ask)) )} 
  
/* {1} OPT.The go function for: noeval_message(self:property,start:integer) [] */
func F_noeval_message_property2 (self *ClaireProperty ,start int) EID { 
    var Result EID 
    /* Let:2 */{ 
      var l *ClaireList   = F_get_args_integer(start)
      /* noccur = 1 */
      if (ClEnv.Debug_I != -1) /* If:3 */{ 
        Result = F_push_debug_property(self,(ClEnv.Index-start),start)
        } else {
        Result = EID{CFALSE.Id(),0}
        /* If-3 */} 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = ToException(C_selector_error.Make(self.Id(),l.Id())).Close()
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = EID{CNIL.Id(),0}
      }}
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: noeval_message @ property (throw: true) 
func E_noeval_message_property2 (self EID,start EID) EID { 
    return /*(sm for noeval_message @ property= EID)*/ F_noeval_message_property2(ToProperty(OBJ(self)),INT(start) )} 
  
// a generic method : same as previously but (1) can be called by other methods
// and (2) takes care of the debugging piece, which implies a slower run (GC)
/* {1} OPT.The go function for: execute(self:method,start:integer,int?:boolean) [] */
func F_execute_method (self *ClaireMethod ,start int,int_ask *ClaireBoolean ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var n int  = self.Domain.Length()
      /* noccur = 2 */
      if (self.Formula.Id() != CNULL) /* If:3 */{ 
        /* Let:4 */{ 
          var retour int  = ClEnv.Base
          /* noccur = 1 */
          /* Let:5 */{ 
            var st_ask *ClaireBoolean   = MakeBoolean((ClEnv.Debug_I != -1) && ((int_ask == CTRUE) || 
                (self.Module_I.Status != 4)))
            /* noccur = 2 */
            ClEnv.Base= start
            F_stack_add(self.Formula.Dimension)
            if (st_ask == CTRUE) /* If:6 */{ 
              Result = F_push_debug_property(self.Selector,n,start)
              } else {
              Result = EID{CFALSE.Id(),0}
              /* If-6 */} 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            /* Let:6 */{ 
              var val *ClaireAny  
              /* noccur = 4 */
              var val_try00007 EID 
              val_try00007 = EVAL(self.Formula.Body)
              /* ERROR PROTECTION INSERTED (val-Result) */
              if ErrorIn(val_try00007) {Result = val_try00007
              } else {
              val = ANY(val_try00007)
              if (st_ask == CTRUE) /* If:7 */{ 
                Result = F_pop_debug_property(self.Selector,0,val)
                } else {
                Result = EID{CFALSE.Id(),0}
                /* If-7 */} 
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              ClEnv.Base= retour
              ClEnv.Index= start
              if ((ClEnv.Debug_I != -1) && 
                  (self.Range.Contains(val) != CTRUE)) /* If:7 */{ 
                Result = ToException(C_range_error.Make(self.Id(),val,self.Range.Id())).Close()
                } else {
                Result = EID{CFALSE.Id(),0}
                /* If-7 */} 
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              Result = val.ToEID()
              }}
              }
              /* Let-6 */} 
            }
            /* Let-5 */} 
          /* Let-4 */} 
        } else {
        /* Let:4 */{ 
          var st_ask *ClaireBoolean   = MakeBoolean((ClEnv.Debug_I != -1) && ((int_ask == CTRUE) || 
              (self.Module_I.Status != 3)) && (self.Selector.Id() != C_debug.Id()))
          /* noccur = 2 */
          /* Let:5 */{ 
            var i int  = ClEnv.Index
            /* noccur = 1 */
            if (st_ask == CTRUE) /* If:6 */{ 
              Result = F_push_debug_property(self.Selector,n,start)
              } else {
              Result = EID{CFALSE.Id(),0}
              /* If-6 */} 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            /* Let:6 */{ 
              var val *ClaireAny  
              /* noccur = 4 */
              var val_try00017 EID 
              val_try00017 = F_stack_apply_function(self.Functional,start,i)
              /* ERROR PROTECTION INSERTED (val-Result) */
              if ErrorIn(val_try00017) {Result = val_try00017
              } else {
              val = ANY(val_try00017)
              if (st_ask == CTRUE) /* If:7 */{ 
                Result = F_pop_debug_property(self.Selector,0,val)
                } else {
                Result = EID{CFALSE.Id(),0}
                /* If-7 */} 
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              if ((ClEnv.Debug_I != -1) && 
                  (self.Range.Contains(val) != CTRUE)) /* If:7 */{ 
                Result = ToException(C_range_error.Make(self.Id(),val,self.Range.Id())).Close()
                } else {
                Result = EID{CFALSE.Id(),0}
                /* If-7 */} 
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              Result = val.ToEID()
              }}
              }
              /* Let-6 */} 
            }
            /* Let-5 */} 
          /* Let-4 */} 
        /* If-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: execute @ method (throw: true) 
func E_execute_method (self EID,start EID,int_ask EID) EID { 
    return /*(sm for execute @ method= EID)*/ F_execute_method(ToMethod(OBJ(self)),INT(start),ToBoolean(OBJ(int_ask)) )} 
  
// the evaluator is open coded
/* {1} OPT.The go function for: eval(self:any) [] */
func F_eval_any (self *ClaireAny ) EID { 
    var Result EID 
    Result = EVAL(self)
    return Result} 
  
// The EID go function for: eval @ list<type_expression>(any) (throw: true) 
func E_eval_any (self EID) EID { 
    return /*(sm for eval @ list<type_expression>(any)= EID)*/ F_eval_any(ANY(self) )} 
  
// this is the standard evaluation
// self_eval(self:object) : any -> self
// reads an inline definition for a method
// notice that it does not return an error
/* {1} OPT.The go function for: inlineok?(self:method,s:string) [] */
func F_inlineok_ask_method (self *ClaireMethod ,s *ClaireString ) *ClaireMethod  { 
    // use function body compiling 
var Unused_try2 EID 
    h_index := ClEnv.Index /* Handle */
    h_base := ClEnv.Base
    /* Let:2 */{ 
      var p *ClaireProperty   = C_read
      /* noccur = 1 */
      /* Let:3 */{ 
        var l *ClaireAny  
        /* noccur = 1 */
        var l_try00024 EID 
        l_try00024 = F_CALL(C_call,ARGS(EID{p.Id(),0},EID{(s).Id(),0}))
        /* ERROR PROTECTION INSERTED (l-Unused_try2) */
        if ErrorIn(l_try00024) {Unused_try2 = l_try00024
        } else {
        l = ANY(l_try00024)
        self.Inline_ask = CTRUE
        /* update:4 */{ 
          var va_arg1 *ClaireMethod  
          var va_arg2 *ClaireLambda  
          va_arg1 = self
          va_arg2 = ToLambda(l)
          /* ---------- now we compile update formula(va_arg1) := va_arg2 ------- */
          va_arg1.Formula = va_arg2
          Unused_try2 = EID{va_arg2.Id(),0}
          /* update-4 */} 
        }
        /* Let-3 */} 
      /* Let-2 */} 
    if ErrorIn(Unused_try2){ 
      /* s=void */ClEnv.Index = h_index
      ClEnv.Base = h_base
      F_tformat_string(MakeString("---- WARNING: inline definition of ~S is wrong\n"),0,MakeConstantList(self.Id()))
      } 
    return  self
    } 
  
// The EID go function for: inlineok? @ method (throw: false) 
func E_inlineok_ask_method (self EID,s EID) EID { 
    return EID{/*(sm for inlineok? @ method= method)*/ F_inlineok_ask_method(ToMethod(OBJ(self)),ToString(OBJ(s)) ).Id(),0}} 
  
// ****************************************************************
// *    Part 2: Update methods                                    *
// ****************************************************************
//get/put for a slot: should be inline
/* {1} OPT.The go function for: get(s:slot,x:object) [] */
func F_get_slot (s *ClaireSlot ,x *ClaireObject ) *ClaireAny  { 
    // use function body compiling 
return  x.SlotGet(s.Index,s.Srange)
    } 
  
// The EID go function for: get @ slot (throw: false) 
func E_get_slot (s EID,x EID) EID { 
    return /*(sm for get @ slot= any)*/ F_get_slot(ToSlot(OBJ(s)),ToObject(OBJ(x)) ).ToEID()} 
  
/* {1} OPT.The go function for: put(s:slot,x:object,y:any) [] */
func F_put_slot (s *ClaireSlot ,x *ClaireObject ,y *ClaireAny ) *ClaireAny  { 
    // use function body compiling 
return  F_store_object(x,
      s.Index,
      s.Srange,
      y,
      s.Selector.Store_ask)
    } 
  
// The EID go function for: put @ slot (throw: false) 
func E_put_slot (s EID,x EID,y EID) EID { 
    return /*(sm for put @ slot= any)*/ F_put_slot(ToSlot(OBJ(s)),ToObject(OBJ(x)),ANY(y) ).ToEID()} 
  
// reading a value from a property (unknown is allowed)
// when unknown is not allowed, we use read which is defined in Kernel
/* {1} OPT.The go function for: get(self:property,x:object) [] */
func F_get_property (self *ClaireProperty ,x *ClaireObject ) *ClaireAny  { 
    // procedure body with s =  
var Result *ClaireAny  
    /* Let:2 */{ 
      var s *ClaireObject   = F__at_property1(self,x.Id().Isa)
      /* noccur = 2 */
      if (C_slot.Id() == s.Isa.Id()) /* If:3 */{ 
        /* Let:4 */{ 
          var g0003 *ClaireSlot   = ToSlot(s.Id())
          /* noccur = 2 */
          Result = x.SlotGet(g0003.Index,g0003.Srange)
          /* Let-4 */} 
        } else {
        Result = CNULL
        /* If-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: get @ property (throw: false) 
func E_get_property (self EID,x EID) EID { 
    return /*(sm for get @ property= any)*/ F_get_property(ToProperty(OBJ(self)),ToObject(OBJ(x)) ).ToEID()} 
  
// a more general value that is useful for types
/* {1} OPT.The go function for: funcall(self:property,x:any) [] */
func F_funcall_property (self *ClaireProperty ,x *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var s *ClaireObject   = F__at_property1(self,x.Isa)
      /* noccur = 4 */
      if (C_slot.Id() == s.Isa.Id()) /* If:3 */{ 
        /* Let:4 */{ 
          var g0005 *ClaireSlot   = ToSlot(s.Id())
          /* noccur = 2 */
          Result = ToObject(x).SlotGet(g0005.Index,g0005.Srange).ToEID()
          /* Let-4 */} 
        /* If!3 */}  else if (C_method.Id() == s.Isa.Id()) /* If:3 */{ 
        /* Let:4 */{ 
          var g0006 *ClaireMethod   = ToMethod(s.Id())
          /* noccur = 1 */
          Result = F_funcall_method1(g0006,x)
          /* Let-4 */} 
        } else {
        Result = EID{CNULL,0}
        /* If-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: funcall @ property (throw: true) 
func E_funcall_property (self EID,x EID) EID { 
    return /*(sm for funcall @ property= EID)*/ F_funcall_property(ToProperty(OBJ(self)),ANY(x) )} 
  
// verifying
/* {1} OPT.The go function for: hold?(self:property,x:object,y:any) [] */
func F_hold_ask_property (self *ClaireProperty ,x *ClaireObject ,y *ClaireAny ) *ClaireBoolean  { 
    // procedure body with s =  
var Result *ClaireBoolean  
    /* Let:2 */{ 
      var s *ClaireObject   = F__at_property1(self,x.Id().Isa)
      /* noccur = 2 */
      if (C_slot.Id() == s.Isa.Id()) /* If:3 */{ 
        /* Let:4 */{ 
          var g0008 *ClaireSlot   = ToSlot(s.Id())
          /* noccur = 2 */
          /* Let:5 */{ 
            var z *ClaireAny   = x.SlotGet(g0008.Index,g0008.Srange)
            /* noccur = 3 */
            if (C_set.Id() == z.Isa.Id()) /* If:6 */{ 
              /* Let:7 */{ 
                var g0009 *ClaireSet   = ToSet(z)
                /* noccur = 1 */
                Result = g0009.Contain_ask(y)
                /* Let-7 */} 
              } else {
              Result = Equal(y,z)
              /* If-6 */} 
            /* Let-5 */} 
          /* Let-4 */} 
        } else {
        Result = CFALSE
        /* If-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: hold? @ property (throw: false) 
func E_hold_ask_property (self EID,x EID,y EID) EID { 
    return EID{/*(sm for hold? @ property= boolean)*/ F_hold_ask_property(ToProperty(OBJ(self)),ToObject(OBJ(x)),ANY(y) ).Id(),0}} 
  
// writing a single value into a slot & trigger rules
//  write = check + put + put_inverse + propagate
//  if_write = put + put_inverse + propagate  (propagate => if_write)
//  update = put + put_inverse
// note in CLAIRE 4: with no inverse/store write_fast, defined in Kernel, works better
/* {1} OPT.The go function for: write(self:property,x:object,y:any) [] */
func F_write_property (self *ClaireProperty ,x *ClaireObject ,y *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var s *ClaireObject   = F__at_property1(self,x.Id().Isa)
      /* noccur = 2 */
      if (C_slot.Id() == s.Isa.Id()) /* If:3 */{ 
        /* Let:4 */{ 
          var g0012 *ClaireSlot   = ToSlot(s.Id())
          /* noccur = 6 */
          if (g0012.Range.Contains(y) != CTRUE) /* If:5 */{ 
            Result = F_range_is_wrong_slot(g0012,y)
            /* If!5 */}  else if ((self.Open < 1) && 
              (x.SlotGet(g0012.Index,g0012.Srange) != CNULL)) /* If:5 */{ 
            Result = ToException(C_general_error.Make(MakeString("[132] Cannot change ~S(~S)").Id(),MakeConstantList(self.Id(),x.Id()).Id())).Close()
            /* If!5 */}  else if ((self.IfWrite != CNULL) && 
              (self.Multivalued_ask != CTRUE)) /* If:5 */{ 
            Result = F_fastcall_relation2(ToRelation(self.Id()),x.Id(),y)
            } else {
            Result = F_update_property(self,
              x,
              g0012.Index,
              g0012.Srange,
              y)
            /* If-5 */} 
          /* Let-4 */} 
        } else {
        Result = ToException(C_selector_error.Make(self.Id(),MakeConstantList(x.Id()).Id())).Close()
        /* If-3 */} 
      /* Let-2 */} 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    Result = y.ToEID()
    }
    return Result} 
  
// The EID go function for: write @ property (throw: true) 
func E_write_property (self EID,x EID,y EID) EID { 
    return /*(sm for write @ property= EID)*/ F_write_property(ToProperty(OBJ(self)),ToObject(OBJ(x)),ANY(y) )} 
  
// the value does not belong to the range: error!
/* {1} OPT.The go function for: range_is_wrong(self:slot,y:any) [] */
func F_range_is_wrong_slot (self *ClaireSlot ,y *ClaireAny ) EID { 
    var Result EID 
    Result = ToException(C_range_error.Make(self.Id(),y,self.Range.Id())).Close()
    return Result} 
  
// The EID go function for: range_is_wrong @ slot (throw: true) 
func E_range_is_wrong_slot (self EID,y EID) EID { 
    return /*(sm for range_is_wrong @ slot= EID)*/ F_range_is_wrong_slot(ToSlot(OBJ(self)),ANY(y) )} 
  
// to remove
/* {1} OPT.The go function for: put(p:property,x:object,n:integer,s:class,y:any) [] */
func F_put_property1 (p *ClaireProperty ,x *ClaireObject ,n int,s *ClaireClass ,y *ClaireAny ) EID { 
    var Result EID 
    Result = F_update_property(p,
      x,
      n,
      s,
      y)
    return Result} 
  
// The EID go function for: put @ list<type_expression>(property, object, integer, class, any) (throw: true) 
func E_put_property1 (p EID,x EID,n EID,s EID,y EID) EID { 
    return /*(sm for put @ list<type_expression>(property, object, integer, class, any)= EID)*/ F_put_property1(ToProperty(OBJ(p)),
      ToObject(OBJ(x)),
      INT(n),
      ToClass(OBJ(s)),
      ANY(y) )} 
  
// update (method called by the compiler)     // v3.0.20: renamed from put !
// update = put + put_inverse  (complex links)
// update uses two satellite methods: update+ and update-
// CLAIRE 4: inverse management only applies with set multivalued properties
/* {1} OPT.The go function for: mClaire/update(p:property,x:object,n:integer,s:class,y:any) [] */
func F_update_property (p *ClaireProperty ,x *ClaireObject ,n int,s *ClaireClass ,y *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var old *ClaireAny   = x.SlotGet(n,s)
      /* noccur = 9 */
      if (ClEnv.Verbose == 8) /* If:3 */{ 
        F_tformat_string(MakeString("update ~S(~S) old = ~S\n"),0,MakeConstantList(p.Id(),x.Id(),old))
        /* If-3 */} 
      if (p.Multivalued_ask.Id() == CTRUE.Id()) /* If:3 */{ 
        if (INT(F_CALL(C_length,ARGS(old.ToEID()))) > 0) /* If:4 */{ 
          /* Let:5 */{ 
            var v *ClaireSet   = ToType(CEMPTY.Id()).EmptySet()
            /* noccur = 2 */
            if (ANY(F_CALL(C_of,ARGS(old.ToEID()))) != C_void.Id()) /* If:6 */{ 
              v.Cast_I(ToType(OBJ(F_CALL(C_of,ARGS(old.ToEID())))))
              /* If-6 */} 
            F_store_object(x,
              n,
              s,
              v.Id(),
              p.Store_ask)
            /* Let-5 */} 
          /* If-4 */} 
        /* Let:4 */{ 
          var r *ClaireRelation   = p.Inverse
          /* noccur = 2 */
          if (r.Id() == CNULL) /* If:5 */{ 
            
            } else {
            /* For:6 */{ 
              var z *ClaireAny  
              _ = z
              var z_support *ClaireSet  
              z_support = ToSet(old)
              for _,z = range(z_support.Values)/* loop2:7 */{ 
                F_update_dash_relation(r,z,x.Id())
                /* loop-7 */} 
              /* For-6 */} 
            /* If-5 */} 
          /* Let-4 */} 
        /* For:4 */{ 
          var z *ClaireAny  
          _ = z
          Result= EID{CFALSE.Id(),0}
          var z_support *ClaireSet  
          z_support = ToSet(y)
          for _,z = range(z_support.Values)/* loop2:5 */{ 
            var void_try6 EID 
            _ = void_try6
            void_try6 = F_add_I_property(p,x,n,z)
            /* ERROR PROTECTION INSERTED (void_try6-Result) */
            if ErrorIn(void_try6) {Result = void_try6
            Result = void_try6
            break
            } else {
            }
            /* loop-5 */} 
          /* For-4 */} 
        /* If!3 */}  else if (Equal(old,y) != CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var r *ClaireRelation   = p.Inverse
          /* noccur = 3 */
          if (r.Id() == CNULL) /* If:5 */{ 
            
            /* If!5 */}  else if ((old != CNULL) && 
              ((r.Id() != p.Id()) || 
                  (Equal(x.Id(),old) != CTRUE))) /* If:5 */{ 
            F_update_dash_relation(r,old,x.Id())
            /* If-5 */} 
          /* Let-4 */} 
        F_store_object(x,
          n,
          s,
          y,
          p.Store_ask)
        Result = F_update_plus_relation(ToRelation(p.Id()),x.Id(),y)
        } else {
        Result = EID{CFALSE.Id(),0}
        /* If-3 */} 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = y.ToEID()
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: mClaire/update @ property (throw: true) 
func E_update_property (p EID,x EID,n EID,s EID,y EID) EID { 
    return /*(sm for mClaire/update @ property= EID)*/ F_update_property(ToProperty(OBJ(p)),
      ToObject(OBJ(x)),
      INT(n),
      ToClass(OBJ(s)),
      ANY(y) )} 
  
// this method checks the correctness of the inverse from a global view.
/* {1} OPT.The go function for: update+(self:relation,x:any,y:any) [] */
func F_update_plus_relation (self *ClaireRelation ,x *ClaireAny ,y *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var r *ClaireRelation   = self.Inverse
      /* noccur = 6 */
      if ((r.Id() != CNULL) && 
          ((r.Id() != self.Id()) || 
              (Equal(x,y) != CTRUE))) /* If:3 */{ 
        if (r.Isa.IsIn(C_property) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var g0014 *ClaireProperty   = ToProperty(r.Id())
            /* noccur = 4 */
            /* Let:6 */{ 
              var s *ClaireObject   = F__at_property1(g0014,y.Isa)
              /* noccur = 2 */
              if (C_slot.Id() == s.Isa.Id()) /* If:7 */{ 
                /* Let:8 */{ 
                  var g0015 *ClaireSlot   = ToSlot(s.Id())
                  /* noccur = 4 */
                  /* Let:9 */{ 
                    var old_y *ClaireAny   = F_get_slot(g0015,ToObject(y))
                    /* noccur = 1 */
                    if (g0014.Multivalued_ask.Id() != CFALSE.Id()) /* If:10 */{ 
                      Result = EID{F_Core_add_value_I_property(g0014,
                        ToObject(y),
                        g0015.Index,
                        ToSet(old_y),
                        x).Id(),0}
                      } else {
                      Result = F_store_object(ToObject(y),
                        g0015.Index,
                        g0015.Srange,
                        x,
                        g0014.Store_ask).ToEID()
                      /* If-10 */} 
                    /* Let-9 */} 
                  /* Let-8 */} 
                } else {
                Result = ToException(C_general_error.Make(MakeString("[133] Inversion of ~S(~S,~S) impossible").Id(),MakeConstantList(self.Id(),x,y).Id())).Close()
                /* If-7 */} 
              /* Let-6 */} 
            /* Let-5 */} 
          /* If!4 */}  else if (C_table.Id() == r.Isa.Id()) /* If:4 */{ 
          /* Let:5 */{ 
            var g0017 *ClaireTable   = ToTable(r.Id())
            /* noccur = 4 */
            /* Let:6 */{ 
              var old_v *ClaireAny   = F_get_table(g0017,y)
              /* noccur = 3 */
              if (g0017.Multivalued_ask.Id() != CFALSE.Id()) /* If:7 */{ 
                Result = EID{F_Core_add_value_I_table(g0017,y,ToSet(old_v),x).Id(),0}
                } else {
                if (old_v != CNULL) /* If:8 */{ 
                  F_update_dash_relation(self,old_v,y)
                  /* If-8 */} 
                F_put_table(g0017,y,x)
                Result = EVOID
                /* If-7 */} 
              /* Let-6 */} 
            /* Let-5 */} 
          } else {
          Result = EID{CFALSE.Id(),0}
          /* If-4 */} 
        } else {
        Result = EID{CFALSE.Id(),0}
        /* If-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: update+ @ relation (throw: true) 
func E_update_plus_relation (self EID,x EID,y EID) EID { 
    return /*(sm for update+ @ relation= EID)*/ F_update_plus_relation(ToRelation(OBJ(self)),ANY(x),ANY(y) )} 
  
// this methods deletes a value in the inverse of a global_relation
/* {1} OPT.The go function for: update-(r:relation,x:any,y:any) [] */
func F_update_dash_relation (r *ClaireRelation ,x *ClaireAny ,y *ClaireAny )  { 
    // procedure body with s =  
if (r.Isa.IsIn(C_property) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0018 *ClaireProperty   = ToProperty(r.Id())
        /* noccur = 2 */
        /* Let:4 */{ 
          var s *ClaireObject   = F__at_property1(g0018,x.Isa)
          /* noccur = 2 */
          if (C_slot.Id() == s.Isa.Id()) /* If:5 */{ 
            /* Let:6 */{ 
              var g0019 *ClaireSlot   = ToSlot(s.Id())
              /* noccur = 2 */
              /* Let:7 */{ 
                var l *ClaireAny   = F_get_slot(g0019,ToObject(x))
                /* noccur = 2 */
                /* Let:8 */{ 
                  var v *ClaireAny  
                  /* noccur = 1 */
                  if (C_set.Id() == l.Isa.Id()) /* If:9 */{ 
                    /* Let:10 */{ 
                      var g0020 *ClaireSet   = ToSet(l)
                      /* noccur = 2 */
                      /* Let:11 */{ 
                        var g0025UU *ClaireSet  
                        /* noccur = 1 */
                        if (g0018.Store_ask == CTRUE) /* If:12 */{ 
                          g0025UU = g0020.Copy()
                          } else {
                          g0025UU = g0020
                          /* If-12 */} 
                        v = g0025UU.Delete(y).Id()
                        /* Let-11 */} 
                      /* Let-10 */} 
                    } else {
                    v = CNULL
                    /* If-9 */} 
                  F_put_slot(g0019,ToObject(x),v)
                  /* Let-8 */} 
                /* Let-7 */} 
              /* Let-6 */} 
            /* If-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (C_table.Id() == r.Isa.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var g0022 *ClaireTable   = ToTable(r.Id())
        /* noccur = 3 */
        /* Let:4 */{ 
          var l *ClaireAny   = F_get_table(g0022,x)
          /* noccur = 2 */
          /* Let:5 */{ 
            var v *ClaireAny  
            /* noccur = 1 */
            if (C_set.Id() == l.Isa.Id()) /* If:6 */{ 
              /* Let:7 */{ 
                var g0023 *ClaireSet   = ToSet(l)
                /* noccur = 2 */
                /* Let:8 */{ 
                  var g0026UU *ClaireSet  
                  /* noccur = 1 */
                  if (g0022.Store_ask == CTRUE) /* If:9 */{ 
                    g0026UU = g0023.Copy()
                    } else {
                    g0026UU = g0023
                    /* If-9 */} 
                  v = g0026UU.Delete(y).Id()
                  /* Let-8 */} 
                /* Let-7 */} 
              } else {
              v = CNULL
              /* If-6 */} 
            F_put_table(g0022,x,v)
            /* Let-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* If-2 */} 
    } 
  
// The EID go function for: update- @ relation (throw: false) 
func E_update_dash_relation (r EID,x EID,y EID) EID { 
    /*(sm for update- @ relation= void)*/ F_update_dash_relation(ToRelation(OBJ(r)),ANY(x),ANY(y) )
    return EVOID} 
  
// this methods adds a value to a multi-slot (used by the compiler)
// this is the multi-valued equivalent of update - we know self to be multivalued (hence a set in Claire 4)
/* {1} OPT.The go function for: add!(self:property,x:object,n:integer,y:any) [] */
func F_add_I_property (self *ClaireProperty ,x *ClaireObject ,n int,y *ClaireAny ) EID { 
    var Result EID 
    if (self.IfWrite != CNULL) /* If:2 */{ 
      Result = F_fastcall_relation2(ToRelation(self.Id()),x.Id(),y)
      } else {
      /* Let:3 */{ 
        var s1 *ClaireSet   = ToSet(x.SlotGet(n,C_object))
        /* noccur = 1 */
        if (F_Core_add_value_I_property(self,
          x,
          n,
          s1,
          y) == CTRUE) /* If:4 */{ 
          Result = F_update_plus_relation(ToRelation(self.Id()),x.Id(),y)
          } else {
          Result = EID{CFALSE.Id(),0}
          /* If-4 */} 
        /* Let-3 */} 
      /* If-2 */} 
    return Result} 
  
// The EID go function for: add! @ property (throw: true) 
func E_add_I_property (self EID,x EID,n EID,y EID) EID { 
    return /*(sm for add! @ property= EID)*/ F_add_I_property(ToProperty(OBJ(self)),
      ToObject(OBJ(x)),
      INT(n),
      ANY(y) )} 
  
// this methods adds a value to a multi-slot (internal form)
// this is the multi-valued equivalent of put
// return true if the set is actually changed (y added to s)
/* {1} OPT.The go function for: add_value!(self:property,x:object,n:integer,s1:set,y:any) [] */
func F_Core_add_value_I_property (self *ClaireProperty ,x *ClaireObject ,n int,s1 *ClaireSet ,y *ClaireAny ) *ClaireBoolean  { 
    // procedure body with s =  
var Result *ClaireBoolean  
    if (s1.Contain_ask(y) != CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var s2 *ClaireSet  
        /* noccur = 1 */
        /* Let:4 */{ 
          var g0027UU *ClaireSet  
          /* noccur = 1 */
          if (self.Store_ask == CTRUE) /* If:5 */{ 
            g0027UU = s1.Copy()
            } else {
            g0027UU = s1
            /* If-5 */} 
          s2 = g0027UU.AddFast(y)
          /* Let-4 */} 
        F_store_object(x,
          n,
          C_object,
          s2.Id(),
          self.Store_ask)
        Result = CTRUE
        /* Let-3 */} 
      } else {
      Result = CFALSE
      /* If-2 */} 
    return Result} 
  
// The EID go function for: add_value! @ property (throw: false) 
func E_Core_add_value_I_property (self EID,x EID,n EID,s1 EID,y EID) EID { 
    return EID{/*(sm for add_value! @ property= boolean)*/ F_Core_add_value_I_property(ToProperty(OBJ(self)),
      ToObject(OBJ(x)),
      INT(n),
      ToSet(OBJ(s1)),
      ANY(y) ).Id(),0}} 
  
// same method with error checking
/* {1} OPT.The go function for: add(self:property,x:object,y:any) [] */
func F_add_property (self *ClaireProperty ,x *ClaireObject ,y *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var s *ClaireObject   = F__at_property1(self,x.Id().Isa)
      /* noccur = 4 */
      if (F_boolean_I_any(s.Id()).Id() != CTRUE.Id()) /* If:3 */{ 
        Result = ToException(C_selector_error.Make(self.Id(),MakeConstantList(x.Id()).Id())).Close()
        /* If!3 */}  else if (self.Multivalued_ask != CTRUE) /* If:3 */{ 
        Result = ToException(C_general_error.Make(MakeString("[134] Cannot apply add to ~S").Id(),MakeConstantList(self.Id()).Id())).Close()
        /* If!3 */}  else if (F_member_type(ToRestriction(s.Id()).Range).Contains(y) == CTRUE) /* If:3 */{ 
        if (self.IfWrite != CNULL) /* If:4 */{ 
          Result = F_fastcall_relation2(ToRelation(self.Id()),x.Id(),y)
          } else {
          Result = F_add_I_property(self,x,ToSlot(s.Id()).Index,y)
          /* If-4 */} 
        } else {
        Result = F_range_is_wrong_slot(ToSlot(s.Id()),y)
        /* If-3 */} 
      /* Let-2 */} 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    Result = y.ToEID()
    }
    return Result} 
  
// The EID go function for: add @ property (throw: true) 
func E_add_property (self EID,x EID,y EID) EID { 
    return /*(sm for add @ property= EID)*/ F_add_property(ToProperty(OBJ(self)),ToObject(OBJ(x)),ANY(y) )} 
  
// known ?
/* {1} OPT.The go function for: known?(self:property,x:object) [] */
func F_known_ask_property (self *ClaireProperty ,x *ClaireObject ) *ClaireBoolean  { 
    // procedure body with s =  
var Result *ClaireBoolean  
    /* Let:2 */{ 
      var s *ClaireObject   = F__at_property1(self,x.Id().Isa)
      /* noccur = 2 */
      if (C_slot.Id() == s.Isa.Id()) /* If:3 */{ 
        /* Let:4 */{ 
          var g0028 *ClaireSlot   = ToSlot(s.Id())
          /* noccur = 2 */
          Result = F__I_equal_any(x.SlotGet(g0028.Index,g0028.Srange),CNULL)
          /* Let-4 */} 
        } else {
        Result = CFALSE
        /* If-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: known? @ property (throw: false) 
func E_known_ask_property (self EID,x EID) EID { 
    return EID{/*(sm for known? @ property= boolean)*/ F_known_ask_property(ToProperty(OBJ(self)),ToObject(OBJ(x)) ).Id(),0}} 
  
/* {1} OPT.The go function for: unknown?(self:property,x:object) [] */
func F_unknown_ask_property (self *ClaireProperty ,x *ClaireObject ) *ClaireBoolean  { 
    // procedure body with s =  
var Result *ClaireBoolean  
    /* Let:2 */{ 
      var s *ClaireObject   = F__at_property1(self,x.Id().Isa)
      /* noccur = 2 */
      if (C_slot.Id() == s.Isa.Id()) /* If:3 */{ 
        /* Let:4 */{ 
          var g0030 *ClaireSlot   = ToSlot(s.Id())
          /* noccur = 2 */
          Result = Equal(x.SlotGet(g0030.Index,g0030.Srange),CNULL)
          /* Let-4 */} 
        } else {
        Result = CTRUE
        /* If-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: unknown? @ property (throw: false) 
func E_unknown_ask_property (self EID,x EID) EID { 
    return EID{/*(sm for unknown? @ property= boolean)*/ F_unknown_ask_property(ToProperty(OBJ(self)),ToObject(OBJ(x)) ).Id(),0}} 
  
// delete takes care of the inverse also
// assumes that self is multivalued -> should check !
/* {1} OPT.The go function for: delete(self:property,x:object,y:any) [] */
func F_delete_property (self *ClaireProperty ,x *ClaireObject ,y *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var s *ClaireObject   = F__at_property1(self,x.Id().Isa)
      /* noccur = 3 */
      if (F_boolean_I_any(s.Id()).Id() != CTRUE.Id()) /* If:3 */{ 
        Result = ToException(C_selector_error.Make(self.Id(),MakeConstantList(x.Id()).Id())).Close()
        /* If!3 */}  else if (self.Multivalued_ask.Id() == CTRUE.Id()) /* If:3 */{ 
        /* Let:4 */{ 
          var l1 *ClaireSet   = ToSet(x.SlotGet(ToSlot(s.Id()).Index,C_object))
          /* noccur = 2 */
          /* Let:5 */{ 
            var l *ClaireSet  
            /* noccur = 2 */
            /* Let:6 */{ 
              var g0032UU *ClaireSet  
              /* noccur = 1 */
              if (self.Store_ask == CTRUE) /* If:7 */{ 
                g0032UU = l1.Copy()
                } else {
                g0032UU = l1
                /* If-7 */} 
              l = g0032UU.Delete(y)
              /* Let-6 */} 
            F_store_object(x,
              ToSlot(s.Id()).Index,
              C_object,
              l.Id(),
              self.Store_ask)
            /* Let:6 */{ 
              var r *ClaireRelation   = self.Inverse
              /* noccur = 2 */
              if (r.Id() != CNULL) /* If:7 */{ 
                F_update_dash_relation(r,y,x.Id())
                /* If-7 */} 
              /* Let-6 */} 
            Result = EID{l.Id(),0}
            /* Let-5 */} 
          /* Let-4 */} 
        } else {
        Result = EID{CFALSE.Id(),0}
        /* If-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: delete @ property (throw: true) 
func E_delete_property (self EID,x EID,y EID) EID { 
    return /*(sm for delete @ property= EID)*/ F_delete_property(ToProperty(OBJ(self)),ToObject(OBJ(x)),ANY(y) )} 
  
// erase is similar for mono-valued properties takes care of the inverse also
// v3.2.22: take care of multi-valued slot as well
/* {1} OPT.The go function for: erase(self:property,x:object) [] */
func F_erase_property (self *ClaireProperty ,x *ClaireObject ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var s *ClaireObject   = F__at_property1(self,x.Id().Isa)
      /* noccur = 8 */
      if (F_boolean_I_any(s.Id()).Id() != CTRUE.Id()) /* If:3 */{ 
        Result = ToException(C_selector_error.Make(self.Id(),MakeConstantList(x.Id()).Id())).Close()
        } else {
        /* Let:4 */{ 
          var y *ClaireAny   = x.SlotGet(ToSlot(s.Id()).Index,ToClass(OBJ(F_CALL(C_mClaire_srange,ARGS(EID{s.Id(),0})))))
          /* noccur = 4 */
          if (self.Multivalued_ask.Id() == CTRUE.Id()) /* If:5 */{ 
            /* Let:6 */{ 
              var r *ClaireRelation   = self.Inverse
              /* noccur = 2 */
              if (r.Id() == CNULL) /* If:7 */{ 
                
                } else {
                /* For:8 */{ 
                  var y1 *ClaireAny  
                  _ = y1
                  var y1_support *ClaireSet  
                  y1_support = ToSet(y)
                  for _,y1 = range(y1_support.Values)/* loop2:9 */{ 
                    F_update_dash_relation(r,y1,x.Id())
                    /* loop-9 */} 
                  /* For-8 */} 
                /* If-7 */} 
              /* Let-6 */} 
            /* Let:6 */{ 
              var l *ClaireSet   = ToSet(y).Empty()
              /* noccur = 2 */
              F_store_object(x,
                ToSlot(s.Id()).Index,
                C_object,
                l.Id(),
                self.Store_ask)
              Result = EID{l.Id(),0}
              /* Let-6 */} 
            } else {
            F_store_object(x,
              ToSlot(s.Id()).Index,
              ToClass(OBJ(F_CALL(C_mClaire_srange,ARGS(EID{s.Id(),0})))),
              ANY(F_CALL(C_default,ARGS(EID{s.Id(),0}))),
              self.Store_ask)
            /* Let:6 */{ 
              var r *ClaireRelation   = self.Inverse
              /* noccur = 2 */
              if ((r.Id() != CNULL) && 
                  (y != CNULL)) /* If:7 */{ 
                F_update_dash_relation(r,y,x.Id())
                /* If-7 */} 
              /* Let-6 */} 
            Result = F_CALL(C_default,ARGS(EID{s.Id(),0}))
            /* If-5 */} 
          /* Let-4 */} 
        /* If-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: erase @ property (throw: true) 
func E_erase_property (self EID,x EID) EID { 
    return /*(sm for erase @ property= EID)*/ F_erase_property(ToProperty(OBJ(self)),ToObject(OBJ(x)) )} 
  
/* {1} OPT.The go function for: set_range(p:property,c:class,r:type) [] */
func F_set_range_property (p *ClaireProperty ,c *ClaireClass ,r *ClaireType )  { 
    // procedure body with s =  
/* Let:2 */{ 
      var s *ClaireSlot   = ToSlot(F__at_property1(p,c).Id())
      /* noccur = 4 */
      s.Range = r
      s.Srange = r.Class_I()
      /* Let-2 */} 
    } 
  
// The EID go function for: set_range @ property (throw: false) 
func E_set_range_property (p EID,c EID,r EID) EID { 
    /*(sm for set_range @ property= void)*/ F_set_range_property(ToProperty(OBJ(p)),ToClass(OBJ(c)),ToType(OBJ(r)) )
    return EVOID} 
  
// no longer needed because changing the range is not changing the prototype ?
// we should rather generate an error if the condition for dealing with
// defaults changes (TODO)
//        if (s.srange != any & s.srange != integer)
//           c.prototype[s.index] := 0))
// this method allows to bypass the storage mechanism - to be optimized ..
/* {1} OPT.The go function for: put_store(self:property,x:object,y:any,b:boolean) [] */
func F_put_store_property2 (self *ClaireProperty ,x *ClaireObject ,y *ClaireAny ,b *ClaireBoolean ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var s *ClaireObject   = F__at_property1(self,x.Id().Isa)
      /* noccur = 2 */
      if (C_slot.Id() == s.Isa.Id()) /* If:3 */{ 
        /* Let:4 */{ 
          var g0033 *ClaireSlot   = ToSlot(s.Id())
          /* noccur = 4 */
          /* Let:5 */{ 
            var z *ClaireAny   = x.SlotGet(g0033.Index,g0033.Srange)
            /* noccur = 1 */
            if (Equal(z,y) != CTRUE) /* If:6 */{ 
              Result = F_store_object(x,
                g0033.Index,
                g0033.Srange,
                y,
                b).ToEID()
              } else {
              Result = EID{CFALSE.Id(),0}
              /* If-6 */} 
            /* Let-5 */} 
          /* Let-4 */} 
        } else {
        Result = ToException(C_selector_error.Make(self.Id(),MakeConstantList(x.Id()).Id())).Close()
        /* If-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: put_store @ property (throw: true) 
func E_put_store_property2 (self EID,x EID,y EID,b EID) EID { 
    return /*(sm for put_store @ property= EID)*/ F_put_store_property2(ToProperty(OBJ(self)),
      ToObject(OBJ(x)),
      ANY(y),
      ToBoolean(OBJ(b)) )} 
  
// tells if we have a multivalued relation nolonger used in CLAIRE 4
// multi? :: property()
// [multi?(x:any) : boolean
//   -> case x (relation (x.multivalued? != false), any false) ]
// new: (v3.0) we have a simpler management of demons thus fastcall can be
// written in CLAIRE. A demon is either a lambda or a function, which
// arguments have precise sorts that match the relation sort
// the demon does everything: put + inverse + propagation
// thus write(R,x,y) <=> fastcall(R,x,y) <=> f(x,y)
/* {1} OPT.The go function for: fastcall(r:relation,x:any,y:any) [] */
func F_fastcall_relation2 (r *ClaireRelation ,x *ClaireAny ,y *ClaireAny ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var f *ClaireAny   = r.IfWrite
      /* noccur = 3 */
      if (f.Isa.IsIn(C_lambda) == CTRUE) /* If:3 */{ 
        Result = F_CALL(C_funcall,ARGS(f.ToEID(),x.ToEID(),y.ToEID()))
        } else {
        Result = F_funcall2(ToFunction(f),x,y)
        /* If-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: fastcall @ relation (throw: true) 
func E_fastcall_relation2 (r EID,x EID,y EID) EID { 
    return /*(sm for fastcall @ relation= EID)*/ F_fastcall_relation2(ToRelation(OBJ(r)),ANY(x),ANY(y) )} 
  
// *********************************************************************
// *   Part 3: Management of definition(p)                             *
// *********************************************************************
// the dictionarty slot
// insertion in the definition tree
/* {1} OPT.The go function for: insert_definition(p:property,r:restriction) [] */
func F_insert_definition_property (p *ClaireProperty ,r *ClaireRestriction )  { 
    // procedure body with s =  
p.Definition = F_initialize_restriction1(r,ToTypeExpression(r.Domain.ValuesO()[1-1]).Class_I(),p.Definition)
    } 
  
// The EID go function for: insert_definition @ property (throw: false) 
func E_insert_definition_property (p EID,r EID) EID { 
    /*(sm for insert_definition @ property= void)*/ F_insert_definition_property(ToProperty(OBJ(p)),ToRestriction(OBJ(r)) )
    return EVOID} 
  
// insert a restriction with class-domain d into a property p
// claire4 : get rid of dispatcher
/* {1} OPT.The go function for: initialize(x:restriction,d:class,l:list) [] */
func F_initialize_restriction1 (x *ClaireRestriction ,d *ClaireClass ,l *ClaireList ) *ClaireList  { 
    // procedure body with s =  
var Result *ClaireList  
    /* Let:2 */{ 
      var p *ClaireProperty   = x.Selector
      /* noccur = 8 */
      if ((p.Restrictions.Length() == 5) && 
          (F_uniform_property(p) == CTRUE)) /* If:3 */{ 
        /* For:4 */{ 
          var r *ClaireAny  
          _ = r
          for _,r = range(p.Restrictions.ValuesO())/* loop:5 */{ 
            F_hashinsert_restriction(ToRestriction(r))
            /* loop-5 */} 
          /* For-4 */} 
        p.Dictionary = CTRUE
        /* If-3 */} 
      if (p.Dictionary == CTRUE) /* If:3 */{ 
        if (F_uniform_restriction(x) == CTRUE) /* If:4 */{ 
          F_hashinsert_restriction(x)
          } else {
          p.Dictionary = CFALSE
          /* If-4 */} 
        /* If-3 */} 
      Result = F_initialize_restriction2(x,l)
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: initialize @ list<type_expression>(restriction, class, list) (throw: false) 
func E_initialize_restriction1 (x EID,d EID,l EID) EID { 
    return EID{/*(sm for initialize @ list<type_expression>(restriction, class, list)= list)*/ F_initialize_restriction1(ToRestriction(OBJ(x)),ToClass(OBJ(d)),ToList(OBJ(l)) ).Id(),0}} 
  
// only uniform properties can use the dictionary representation
/* {1} OPT.The go function for: uniform(x:restriction) [] */
func F_uniform_restriction (x *ClaireRestriction ) *ClaireBoolean  { 
    // procedure body with s =  
var Result *ClaireBoolean  
    /* Let:2 */{ 
      var l *ClaireList   = x.Domain
      /* noccur = 5 */
      /* Let:3 */{ 
        var n int  = l.Length()
        /* noccur = 2 */
        /* Let:4 */{ 
          var g0036UU *ClaireAny  
          /* noccur = 1 */
          /* For:5 */{ 
            var r *ClaireAny  
            _ = r
            g0036UU= CFALSE.Id()
            for _,r = range(x.Selector.Restrictions.ValuesO())/* loop:6 */{ 
              var g0037I *ClaireBoolean  
              /* Let:7 */{ 
                var g0038UU *ClaireBoolean  
                /* noccur = 1 */
                /* Let:8 */{ 
                  var l2 *ClaireList   = ToRestriction(r).Domain
                  /* noccur = 6 */
                  /* and:9 */{ 
                    var v_and9 *ClaireBoolean  
                    
                    v_and9 = Equal(C_class.Id(),l2.ValuesO()[1-1].Isa.Id())
                    if (v_and9 == CFALSE) {g0038UU = CFALSE
                    } else /* arg:10 */{ 
                      v_and9 = Equal(MakeInteger(l2.Length()).Id(),MakeInteger(n).Id())
                      if (v_and9 == CFALSE) {g0038UU = CFALSE
                      } else /* arg:11 */{ 
                        v_and9 = F__I_equal_any(l2.ValuesO()[1-1],C_listargs.Id())
                        if (v_and9 == CFALSE) {g0038UU = CFALSE
                        } else /* arg:12 */{ 
                          /* Let:13 */{ 
                            var g0039UU *ClaireAny  
                            /* noccur = 1 */
                            /* Let:14 */{ 
                              var i int  = 2
                              /* noccur = 10 */
                              /* Let:15 */{ 
                                var g0035 int  = n
                                /* noccur = 1 */
                                g0039UU= CFALSE.Id()
                                for (i <= g0035) /* while:16 */{ 
                                  if ((Equal(l.ValuesO()[i-1],l2.ValuesO()[i-1]) != CTRUE) && 
                                      ((l.ValuesO()[i-1].Isa.Id() == C_class.Id()) || 
                                          ((l.ValuesO()[i-1].Isa.Id() != l2.ValuesO()[i-1].Isa.Id()) || 
                                            (F__equaltype_ask_any(ToType(l.ValuesO()[i-1]),ToType(l2.ValuesO()[i-1])) != CTRUE)))) /* If:17 */{ 
                                     /*v = g0039UU, s =any*/
g0039UU = CTRUE.Id()
                                    break
                                    /* If-17 */} 
                                  i = (i+1)
                                  /* while-16 */} 
                                /* Let-15 */} 
                              /* Let-14 */} 
                            v_and9 = F_not_any(g0039UU)
                            /* Let-13 */} 
                          if (v_and9 == CFALSE) {g0038UU = CFALSE
                          } else /* arg:13 */{ 
                            g0038UU = CTRUE/* arg-13 */} 
                          /* arg-12 */} 
                        /* arg-11 */} 
                      /* arg-10 */} 
                    /* and-9 */} 
                  /* Let-8 */} 
                g0037I = g0038UU.Not
                /* Let-7 */} 
              if (g0037I == CTRUE) /* If:7 */{ 
                 /*v = g0036UU, s =any*/
g0036UU = CTRUE.Id()
                break
                /* If-7 */} 
              /* loop-6 */} 
            /* For-5 */} 
          Result = F_not_any(g0036UU)
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: uniform @ restriction (throw: false) 
func E_uniform_restriction (x EID) EID { 
    return EID{/*(sm for uniform @ restriction= boolean)*/ F_uniform_restriction(ToRestriction(OBJ(x)) ).Id(),0}} 
  
// v3.3.36      
// v3.0.54 check that a uniform property only uses methods !
/* {1} OPT.The go function for: uniform(p:property) [] */
func F_uniform_property (p *ClaireProperty ) *ClaireBoolean  { 
    // procedure body with s =  
var Result *ClaireBoolean  
    /* and:2 */{ 
      var v_and2 *ClaireBoolean  
      
      /* Let:3 */{ 
        var g0040UU *ClaireAny  
        /* noccur = 1 */
        /* For:4 */{ 
          var x *ClaireAny  
          _ = x
          g0040UU= CFALSE.Id()
          for _,x = range(p.Restrictions.ValuesO())/* loop:5 */{ 
            if (C_method.Id() != x.Isa.Id()) /* If:6 */{ 
               /*v = g0040UU, s =any*/
g0040UU = CTRUE.Id()
              break
              /* If-6 */} 
            /* loop-5 */} 
          /* For-4 */} 
        v_and2 = F_not_any(g0040UU)
        /* Let-3 */} 
      if (v_and2 == CFALSE) {Result = CFALSE
      } else /* arg:3 */{ 
        v_and2 = F_uniform_restriction(ToRestriction(p.Restrictions.ValuesO()[1-1]))
        if (v_and2 == CFALSE) {Result = CFALSE
        } else /* arg:4 */{ 
          Result = CTRUE/* arg-4 */} 
        /* arg-3 */} 
      /* and-2 */} 
    return Result} 
  
// The EID go function for: uniform @ property (throw: false) 
func E_uniform_property (p EID) EID { 
    return EID{/*(sm for uniform @ property= boolean)*/ F_uniform_property(ToProperty(OBJ(p)) ).Id(),0}} 
  
// insert a restriction in a list with the good order
/* {1} OPT.The go function for: initialize(x:restriction,l:list) [] */
func F_initialize_restriction2 (x *ClaireRestriction ,l *ClaireList ) *ClaireList  { 
    // procedure body with s =  
var Result *ClaireList  
    /* Let:2 */{ 
      var l1 *ClaireList   = CNIL
      /* noccur = 4 */
      /* Let:3 */{ 
        var i int  = 1
        /* noccur = 6 */
        /* Let:4 */{ 
          var g0041 int  = l.Length()
          /* noccur = 1 */
          for (i <= g0041) /* while:5 */{ 
            /* Let:6 */{ 
              var l2 *ClaireList   = ToRestriction(l.At(i-1)).Domain
              /* noccur = 4 */
              if (F_tmatch_ask_list(x.Domain,l2) == CTRUE) /* If:7 */{ 
                if (F_tmatch_ask_list(l2,x.Domain) == CTRUE) /* If:8 */{ 
                  ToArray(l.Id()).NthPut(i,x.Id())
                  l1 = l
                   /*v = Result, s =void*/

                  break
                  } else {
                  l1 = ToList(ANY(l.Nth_plus(i,x.Id())))
                   /*v = Result, s =void*/

                  break
                  /* If-8 */} 
                /* If!7 */}  else if ((F_tmatch_ask_list(l2,x.Domain) != CTRUE) && 
                  (F_join_list(x.Domain,l2) == CTRUE)) /* If:7 */{ 
                F_tformat_string(MakeString("~S and ~S are conflicting"),2,MakeConstantList(l.At(1-1),x.Id()))
                /* If-7 */} 
              /* Let-6 */} 
            i = (i+1)
            /* while-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      if (l1.Length() != 0) /* If:3 */{ 
        Result = l1
        } else {
        Result = l.AddFast(x.Id())
        /* If-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: initialize @ list<type_expression>(restriction, list) (throw: false) 
func E_initialize_restriction2 (x EID,l EID) EID { 
    return EID{/*(sm for initialize @ list<type_expression>(restriction, list)= list)*/ F_initialize_restriction2(ToRestriction(OBJ(x)),ToList(OBJ(l)) ).Id(),0}} 
  
// definition of dictionary: standart hash-table
/* {1} OPT.The go function for: hashinsert(m:restriction) [] */
func F_hashinsert_restriction (m *ClaireRestriction ) *ClaireAny  { 
    // procedure body with s =  
var Result *ClaireAny  
    /* Let:2 */{ 
      var c *ClaireClass   = F_domain_I_restriction(m)
      /* noccur = 1 */
      /* For:3 */{ 
        var c2 *ClaireAny  
        _ = c2
        Result= CFALSE.Id()
        for _,c2 = range(c.Descendents.Values)/* loop:4 */{ 
          F_hashinsert_class(ToClass(c2),ToMethod(m.Id()))
          /* loop-4 */} 
        /* For-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: hashinsert @ restriction (throw: false) 
func E_hashinsert_restriction (m EID) EID { 
    return /*(sm for hashinsert @ restriction= any)*/ F_hashinsert_restriction(ToRestriction(OBJ(m)) ).ToEID()} 
  
// insert into the hash table - since the order is not garanteed when we build the dictionary, we
// need to check that m is more suited than anything that could be there
/* {1} OPT.The go function for: hashinsert(c:class,m:method) [] */
func F_hashinsert_class (c *ClaireClass ,m *ClaireMethod ) *ClaireAny  { 
    // use function body compiling 
if (c.Dictionary.Id() == CNULL) /* If:2 */{ 
      c.Dictionary = ToType(C_property.Id()).Map_I(ToType(C_method.Id()))
      /* If-2 */} 
    /* Let:2 */{ 
      var m1 *ClaireAny   = F_dict_get_any(c.Dictionary.Id(),m.Selector.Id())
      /* noccur = 2 */
      if ((m1 == CNULL) || 
          (F_domain_I_restriction(ToRestriction(m.Id())).IsIn(F_domain_I_restriction(ToRestriction(m1))) == CTRUE)) /* If:3 */{ 
        F_dict_put_any(c.Dictionary.Id(),m.Selector.Id(),m.Id())
        /* If-3 */} 
      /* Let-2 */} 
    return  c.Dictionary.Id()
    } 
  
// The EID go function for: hashinsert @ class (throw: false) 
func E_hashinsert_class (c EID,m EID) EID { 
    return /*(sm for hashinsert @ class= any)*/ F_hashinsert_class(ToClass(OBJ(c)),ToMethod(OBJ(m)) ).ToEID()} 
  
// read the value in the directory (a method or unknown)
/* {1} OPT.The go function for: hashget(c:class,p:property) [] */
func F_hashget_class (c *ClaireClass ,p *ClaireProperty ) *ClaireObject  { 
    // use function body compiling 
return  ToObject(F_dict_get_any(c.Dictionary.Id(),p.Id()))
    } 
  
// The EID go function for: hashget @ class (throw: false) 
func E_hashget_class (c EID,p EID) EID { 
    return EID{/*(sm for hashget @ class= object)*/ F_hashget_class(ToClass(OBJ(c)),ToProperty(OBJ(p)) ).Id(),0}} 
  
// UGLY CAST to remove
// look if two signature have a non-empty intersection
// note that the first case with classes is necessary for bootstraping
/* {1} OPT.The go function for: join(x:list,y:list) [] */
func F_join_list (x *ClaireList ,y *ClaireList ) *ClaireBoolean  { 
    // procedure body with s =  
var Result *ClaireBoolean  
    /* Let:2 */{ 
      var n int  = x.Length()
      /* noccur = 3 */
      /* and:3 */{ 
        var v_and3 *ClaireBoolean  
        
        v_and3 = Equal(MakeInteger(n).Id(),MakeInteger(y.Length()).Id())
        if (v_and3 == CFALSE) {Result = CFALSE
        } else /* arg:4 */{ 
          /* Let:5 */{ 
            var g0044UU *ClaireAny  
            /* noccur = 1 */
            /* Let:6 */{ 
              var i int  = 1
              /* noccur = 5 */
              /* Let:7 */{ 
                var g0042 int  = n
                /* noccur = 1 */
                g0044UU= CFALSE.Id()
                for (i <= g0042) /* while:8 */{ 
                  if (F_boolean_I_any(F_join_class(ToTypeExpression(x.At(i-1)).Class_I(),ToTypeExpression(y.At(i-1)).Class_I()).Id()).Id() != CTRUE.Id()) /* If:9 */{ 
                     /*v = g0044UU, s =any*/
g0044UU = CTRUE.Id()
                    break
                    /* If-9 */} 
                  i = (i+1)
                  /* while-8 */} 
                /* Let-7 */} 
              /* Let-6 */} 
            v_and3 = F_not_any(g0044UU)
            /* Let-5 */} 
          if (v_and3 == CFALSE) {Result = CFALSE
          } else /* arg:5 */{ 
            /* Let:6 */{ 
              var g0045UU *ClaireAny  
              /* noccur = 1 */
              /* Let:7 */{ 
                var i int  = 1
                /* noccur = 5 */
                /* Let:8 */{ 
                  var g0043 int  = n
                  /* noccur = 1 */
                  g0045UU= CFALSE.Id()
                  for (i <= g0043) /* while:9 */{ 
                    if (F_boolean_I_any(ANY(F_CALL(ToProperty(C_glb.Id()),ARGS(x.At(i-1).ToEID(),y.At(i-1).ToEID())))).Id() != CTRUE.Id()) /* If:10 */{ 
                       /*v = g0045UU, s =any*/
g0045UU = CTRUE.Id()
                      break
                      /* If-10 */} 
                    i = (i+1)
                    /* while-9 */} 
                  /* Let-8 */} 
                /* Let-7 */} 
              v_and3 = F_not_any(g0045UU)
              /* Let-6 */} 
            if (v_and3 == CFALSE) {Result = CFALSE
            } else /* arg:6 */{ 
              Result = CTRUE/* arg-6 */} 
            /* arg-5 */} 
          /* arg-4 */} 
        /* and-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: join @ list (throw: false) 
func E_join_list (x EID,y EID) EID { 
    return EID{/*(sm for join @ list= boolean)*/ F_join_list(ToList(OBJ(x)),ToList(OBJ(y)) ).Id(),0}} 
  
// *********************************************************************
// *      Part 4: Matching Methods                                     *
// *********************************************************************
// Key Axiom : this code is not using dynamic calls because we use the two closed forms %type and <=t 
// which are defined in Kernel as functions (Contains and Included)
// this is the method that matches the compilation pattern 
// n is the number of args that have been pushed in the stack
/* {1} OPT.The go function for: stack_apply(p:property,n:integer) [] */
func F_CALL (p *ClaireProperty ,n int) EID { 
    var Result EID 
    /* Let:2 */{ 
      var i int  = (ClEnv.Index-n)
      /* noccur = 3 */
      Result = F_eval_message_property(p,F_find_which_property(p,i,OWNER(ClEnv.EvalStack[i])),i,CFALSE)
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: stack_apply @ property (throw: true) 
func E_CALL (p EID,n EID) EID { 
    return /*(sm for stack_apply @ property= EID)*/ F_CALL(ToProperty(OBJ(p)),INT(n) )} 
  
// version where the class of first argument is forced (super)       
/* {1} OPT.The go function for: super_apply(p:property,c:class,n:integer) [] */
func F_SUPER (p *ClaireProperty ,c *ClaireClass ,n int) EID { 
    var Result EID 
    /* Let:2 */{ 
      var i int  = (ClEnv.Index-n)
      /* noccur = 2 */
      Result = F_eval_message_property(p,F_find_which_property(p,i,c),i,CFALSE)
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: super_apply @ property (throw: true) 
func E_SUPER (p EID,c EID,n EID) EID { 
    return /*(sm for super_apply @ property= EID)*/ F_SUPER(ToProperty(OBJ(p)),ToClass(OBJ(c)),INT(n) )} 
  
// find the correct restrictions to be applied on a given set
// This is also optimized because it is very useful (it returns false if none is found)
/* {1} OPT.The go function for: @(self:property,x:class) [] */
func F__at_property1 (self *ClaireProperty ,x *ClaireClass ) *ClaireObject  { 
    // procedure body with s =  
var Result *ClaireObject  
    if (self.Dictionary == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var rx *ClaireObject   = F_hashget_class(x,self)
        /* noccur = 2 */
        Result = ToObject(IfThenElse((rx.Id() != CNULL),
          rx.Id(),
          CFALSE.Id()))
        /* Let-3 */} 
      } else {
      /* Let:3 */{ 
        var rx *ClaireAny  
        /* noccur = 2 */
        /* Let:4 */{ 
          var r_some *ClaireAny   = CNULL
          /* noccur = 2 */
          /* For:5 */{ 
            var r *ClaireAny  
            _ = r
            for _,r = range(self.Definition.ValuesO())/* loop:6 */{ 
              if (x.IsIn(ToTypeExpression(ToRestriction(r).Domain.ValuesO()[1-1]).Class_I()) == CTRUE) /* If:7 */{ 
                 /*v = rx, s =void*/
r_some = r
                break
                /* If-7 */} 
              /* loop-6 */} 
            /* For-5 */} 
          rx = r_some
          /* Let-4 */} 
        Result = ToObject(IfThenElse((rx != CNULL),
          rx,
          CFALSE.Id()))
        /* Let-3 */} 
      /* If-2 */} 
    return Result} 
  
// The EID go function for: @ @ list<type_expression>(property, class) (throw: false) 
func E__at_property1 (self EID,x EID) EID { 
    return EID{/*(sm for @ @ list<type_expression>(property, class)= object)*/ F__at_property1(ToProperty(OBJ(self)),ToClass(OBJ(x)) ).Id(),0}} 
  
// finds a property through its full domain
/* {1} OPT.The go function for: @(self:property,lt:list) [] */
func F__at_property2 (self *ClaireProperty ,lt *ClaireList ) *ClaireObject  { 
    // procedure body with s =  
var Result *ClaireObject  
    /* Let:2 */{ 
      var rx *ClaireAny  
      /* noccur = 2 */
      /* Let:3 */{ 
        var r_some *ClaireAny   = CNULL
        /* noccur = 2 */
        /* For:4 */{ 
          var r *ClaireAny  
          _ = r
          for _,r = range(self.Definition.ValuesO())/* loop:5 */{ 
            if (F_tmatch_ask_list(lt,ToRestriction(r).Domain) == CTRUE) /* If:6 */{ 
               /*v = rx, s =void*/
r_some = r
              break
              /* If-6 */} 
            /* loop-5 */} 
          /* For-4 */} 
        rx = r_some
        /* Let-3 */} 
      Result = ToObject(IfThenElse((rx != CNULL),
        rx,
        CFALSE.Id()))
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: @ @ list<type_expression>(property, list) (throw: false) 
func E__at_property2 (self EID,lt EID) EID { 
    return EID{/*(sm for @ @ list<type_expression>(property, list)= object)*/ F__at_property2(ToProperty(OBJ(self)),ToList(OBJ(lt)) ).Id(),0}} 
  
// method's pattern matching : l is non nil, hence last(l) is safe  {called in find_which}
// we match a list of args in the stack [n ... m] to the list of type_expressions l
/* {1} OPT.The go function for: matching?(l:list,n:integer,m:integer) [] */
func F_matching_ask_list (l *ClaireList ,n int,m int) *ClaireBoolean  { 
    // procedure body with s =  
var Result *ClaireBoolean  
    /* Let:2 */{ 
      var x int  = (m-n)
      /* noccur = 4 */
      /* Let:3 */{ 
        var z int  = l.Length()
        /* noccur = 3 */
        if ((z == x) && 
            (l.ValuesO()[x-1].Id() != C_listargs.Id())) /* If:4 */{ 
          /* Let:5 */{ 
            var g0048UU *ClaireAny  
            /* noccur = 1 */
            /* Let:6 */{ 
              var i int  = 1
              /* noccur = 5 */
              /* Let:7 */{ 
                var g0046 int  = x
                /* noccur = 1 */
                g0048UU= CFALSE.Id()
                for (i <= g0046) /* while:8 */{ 
                  /* Let:9 */{ 
                    var y int  = ((n-1)+i)
                    /* noccur = 2 */
                    /* Let:10 */{ 
                      var u *ClaireAny   = l.ValuesO()[i-1].Id()
                      /* noccur = 3 */
                      var g0049I *ClaireBoolean  
                      if (u.Isa.Id() == C_class.Id()) /* If:11 */{ 
                        g0049I = OWNER(ClEnv.EvalStack[y]).IsIn(ToClass(u)).Not
                        } else {
                        g0049I = F_vmatch_ask_any(u,ANY(ClEnv.EvalStack[y]),n).Not
                        /* If-11 */} 
                      if (g0049I == CTRUE) /* If:11 */{ 
                         /*v = g0048UU, s =any*/
g0048UU = CTRUE.Id()
                        break
                        /* If-11 */} 
                      /* Let-10 */} 
                    /* Let-9 */} 
                  i = (i+1)
                  /* while-8 */} 
                /* Let-7 */} 
              /* Let-6 */} 
            Result = F_not_any(g0048UU)
            /* Let-5 */} 
          /* If!4 */}  else if ((ANY(F_last_list(l)) == C_listargs.Id()) && 
            (x >= (z-1))) /* If:4 */{ 
          /* Let:5 */{ 
            var g0050UU *ClaireAny  
            /* noccur = 1 */
            /* Let:6 */{ 
              var i int  = 1
              /* noccur = 6 */
              /* Let:7 */{ 
                var g0047 int  = z
                /* noccur = 1 */
                g0050UU= CFALSE.Id()
                for (i <= g0047) /* while:8 */{ 
                  /* Let:9 */{ 
                    var y int  = ((n-1)+i)
                    /* noccur = 4 */
                    if (l.At(i-1) == C_listargs.Id()) /* If:10 */{ 
                      ClEnv.EvalStack[y]=EID{F_get_args_integer(y).Id(),0}
                      ClEnv.Index= (y+1)
                       /*v = g0050UU, s =any*/
g0050UU = CFALSE.Id()
                      break
                      /* If!10 */}  else if (F_vmatch_ask_any(l.At(i-1),ANY(ClEnv.EvalStack[y]),n) != CTRUE) /* If:10 */{ 
                       /*v = g0050UU, s =any*/
g0050UU = CTRUE.Id()
                      break
                      /* If-10 */} 
                    /* Let-9 */} 
                  i = (i+1)
                  /* while-8 */} 
                /* Let-7 */} 
              /* Let-6 */} 
            Result = F_not_any(g0050UU)
            /* Let-5 */} 
          } else {
          Result = CFALSE
          /* If-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: matching? @ list (throw: false) 
func E_matching_ask_list (l EID,n EID,m EID) EID { 
    return EID{/*(sm for matching? @ list= boolean)*/ F_matching_ask_list(ToList(OBJ(l)),INT(n),INT(m) ).Id(),0}} 
  
// type's pattern matching - almost like % but accepts patterns such as Reference (extended in Optimizer)
// this is why we pass n (index in stack) as an argument
// t is the type expression and x is the value
/* {1} OPT.The go function for: vmatch?(t:any,x:any,n:integer) [] */
func F_vmatch_ask_any (t *ClaireAny ,x *ClaireAny ,n int) *ClaireBoolean  { 
    // procedure body with s =  
var Result *ClaireBoolean  
    if (C_class.Id() == t.Isa.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var g0051 *ClaireClass   = ToClass(t)
        /* noccur = 1 */
        Result = x.Isa.IsIn(g0051)
        /* Let-3 */} 
      /* If!2 */}  else if (C_set.Id() == t.Isa.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var g0052 *ClaireSet   = ToSet(t)
        /* noccur = 1 */
        Result = g0052.Contain_ask(x)
        /* Let-3 */} 
      /* If!2 */}  else if (t.Isa.IsIn(C_subtype) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0053 *ClaireSubtype   = ToSubtype(t)
        /* noccur = 3 */
        /* and:4 */{ 
          var v_and4 *ClaireBoolean  
          
          if (g0053.Arg.Id() == C_subtype.Id()) /* If:5 */{ 
            v_and4 = x.Isa.IsIn(C_type)
            } else {
            v_and4 = F__Z_any1(x,g0053.Arg)
            /* If-5 */} 
          if (v_and4 == CFALSE) {Result = CFALSE
          } else /* arg:5 */{ 
            v_and4 = ToType(x).Included(g0053.T1)
            if (v_and4 == CFALSE) {Result = CFALSE
            } else /* arg:6 */{ 
              Result = CTRUE/* arg-6 */} 
            /* arg-5 */} 
          /* and-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (t.Isa.IsIn(C_Param) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0054 *ClaireParam   = To_Param(t)
        /* noccur = 4 */
        /* and:4 */{ 
          var v_and4 *ClaireBoolean  
          
          v_and4 = F_vmatch_ask_any(g0054.Arg.Id(),x,n)
          if (v_and4 == CFALSE) {Result = CFALSE
          } else /* arg:5 */{ 
            /* Let:6 */{ 
              var g0062UU *ClaireAny  
              /* noccur = 1 */
              /* Let:7 */{ 
                var i int  = 1
                /* noccur = 5 */
                /* Let:8 */{ 
                  var g0055 int  = g0054.Params.Length()
                  /* noccur = 1 */
                  g0062UU= CFALSE.Id()
                  for (i <= g0055) /* while:9 */{ 
                    var g0063I *ClaireBoolean  
                    /* Let:10 */{ 
                      var g0064UU *ClaireBoolean  
                      /* noccur = 1 */
                      /* Let:11 */{ 
                        var _Zt *ClaireAny   = g0054.Args.At(i-1)
                        /* noccur = 3 */
                        /* Let:12 */{ 
                          var _Zv *ClaireAny   = ANY(F_funcall_property(ToProperty(g0054.Params.At(i-1)),x))
                          /* noccur = 3 */
                          if ((C_set.Id() == _Zt.Isa.Id()) && 
                              (_Zv.Isa.IsIn(C_type) == CTRUE)) /* If:13 */{ 
                            /* Let:14 */{ 
                              var g0065UU *ClaireAny  
                              /* noccur = 1 */
                              /* For:15 */{ 
                                var z *ClaireAny  
                                _ = z
                                g0065UU= CFALSE.Id()
                                var z_support *ClaireSet  
                                z_support = ToSet(_Zt)
                                for _,z = range(z_support.Values)/* loop2:16 */{ 
                                  if (F__equaltype_ask_any(ToType(_Zv),ToType(z)) == CTRUE) /* If:17 */{ 
                                     /*v = g0065UU, s =any*/
g0065UU = CTRUE.Id()
                                    break
                                    /* If-17 */} 
                                  /* loop-16 */} 
                                /* For-15 */} 
                              g0064UU = F_boolean_I_any(g0065UU)
                              /* Let-14 */} 
                            } else {
                            g0064UU = F_vmatch_ask_any(_Zt,_Zv,n)
                            /* If-13 */} 
                          /* Let-12 */} 
                        /* Let-11 */} 
                      g0063I = g0064UU.Not
                      /* Let-10 */} 
                    if (g0063I == CTRUE) /* If:10 */{ 
                       /*v = g0062UU, s =any*/
g0062UU = CTRUE.Id()
                      break
                      /* If-10 */} 
                    i = (i+1)
                    /* while-9 */} 
                  /* Let-8 */} 
                /* Let-7 */} 
              v_and4 = F_not_any(g0062UU)
              /* Let-6 */} 
            if (v_and4 == CFALSE) {Result = CFALSE
            } else /* arg:6 */{ 
              Result = CTRUE/* arg-6 */} 
            /* arg-5 */} 
          /* and-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (t.Isa.IsIn(C_Reference) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0056 *ClaireReference   = To_Reference(t)
        /* noccur = 3 */
        /* Let:4 */{ 
          var v *ClaireAny   = F_get_Reference(g0056,ANY(ClEnv.EvalStack[(n+g0056.Index)]))
          /* noccur = 2 */
          if (g0056.Arg == CTRUE) /* If:5 */{ 
            Result = Equal(x,v)
            } else {
            Result = ToType(v).Contains(x)
            /* If-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (C_tuple.Id() == t.Isa.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var g0057 *ClaireTuple   = ToTuple(t)
        /* noccur = 2 */
        if (C_tuple.Id() == x.Isa.Id()) /* If:4 */{ 
          /* Let:5 */{ 
            var g0058 *ClaireTuple   = ToTuple(x)
            /* noccur = 3 */
            /* and:6 */{ 
              var v_and6 *ClaireBoolean  
              
              v_and6 = Equal(MakeInteger(g0057.Length()).Id(),MakeInteger(g0058.Length()).Id())
              if (v_and6 == CFALSE) {Result = CFALSE
              } else /* arg:7 */{ 
                /* Let:8 */{ 
                  var g0066UU *ClaireAny  
                  /* noccur = 1 */
                  /* Let:9 */{ 
                    var i int  = 1
                    /* noccur = 5 */
                    /* Let:10 */{ 
                      var g0059 int  = g0058.Length()
                      /* noccur = 1 */
                      g0066UU= CFALSE.Id()
                      for (i <= g0059) /* while:11 */{ 
                        if (F_vmatch_ask_any(ToList(g0057.Id()).At(i-1),ToList(g0058.Id()).At(i-1),n) != CTRUE) /* If:12 */{ 
                           /*v = g0066UU, s =any*/
g0066UU = CTRUE.Id()
                          break
                          /* If-12 */} 
                        i = (i+1)
                        /* while-11 */} 
                      /* Let-10 */} 
                    /* Let-9 */} 
                  v_and6 = F_not_any(g0066UU)
                  /* Let-8 */} 
                if (v_and6 == CFALSE) {Result = CFALSE
                } else /* arg:8 */{ 
                  Result = CTRUE/* arg-8 */} 
                /* arg-7 */} 
              /* and-6 */} 
            /* Let-5 */} 
          } else {
          Result = CFALSE
          /* If-4 */} 
        /* Let-3 */} 
      } else {
      Result = ToBoolean(ANY(F_BELONG(x,t)))
      /* If-2 */} 
    return Result} 
  
// The EID go function for: vmatch? @ any (throw: false) 
func E_vmatch_ask_any (t EID,x EID,n EID) EID { 
    return EID{/*(sm for vmatch? @ any= boolean)*/ F_vmatch_ask_any(ANY(t),ANY(x),INT(n) ).Id(),0}} 
  
// extensibility for type_expressions
// method's pattern matching based on type expressions (i.e. l2 is another list of type expressions).
// this is an extension of <=t to   all type expressions
/* {1} OPT.The go function for: tmatch?(l:list,l2:list) [] */
func F_tmatch_ask_list (l *ClaireList ,l2 *ClaireList ) *ClaireBoolean  { 
    // procedure body with s =  
var Result *ClaireBoolean  
    /* Let:2 */{ 
      var x int  = l2.Length()
      /* noccur = 5 */
      /* Let:3 */{ 
        var z int  = l.Length()
        /* noccur = 1 */
        if ((l.Length() != x) && 
            ((l2.At(x-1) != C_listargs.Id()) || 
                (z < (x-1)))) /* If:4 */{ 
          Result = CFALSE
          } else {
          /* Let:5 */{ 
            var g0068UU *ClaireAny  
            /* noccur = 1 */
            /* Let:6 */{ 
              var i int  = 1
              /* noccur = 7 */
              /* Let:7 */{ 
                var g0067 int  = x
                /* noccur = 1 */
                g0068UU= CFALSE.Id()
                for (i <= g0067) /* while:8 */{ 
                  if ((i == x) && 
                      (l2.At(i-1) == C_listargs.Id())) /* If:9 */{ 
                     /*v = g0068UU, s =any*/
g0068UU = CFALSE.Id()
                    break
                    /* If!9 */}  else if (F_tmatch_ask_any(l.At(i-1),l2.At(i-1),l) != CTRUE) /* If:9 */{ 
                     /*v = g0068UU, s =any*/
g0068UU = CTRUE.Id()
                    break
                    /* If-9 */} 
                  i = (i+1)
                  /* while-8 */} 
                /* Let-7 */} 
              /* Let-6 */} 
            Result = F_not_any(g0068UU)
            /* Let-5 */} 
          /* If-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: tmatch? @ list (throw: false) 
func E_tmatch_ask_list (l EID,l2 EID) EID { 
    return EID{/*(sm for tmatch? @ list= boolean)*/ F_tmatch_ask_list(ToList(OBJ(l)),ToList(OBJ(l2)) ).Id(),0}} 
  
// type_expression pattern matching (t is the variable and t2 the pattern)
// this is an extension of <=t for the pattern Reference
/* {1} OPT.The go function for: tmatch?(t:any,mClaire/t2:any,l:list) [] */
func F_tmatch_ask_any (t *ClaireAny ,t2 *ClaireAny ,l *ClaireList ) *ClaireBoolean  { 
    // procedure body with s =  
var Result *ClaireBoolean  
    if (t2.Isa.IsIn(C_Reference) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0069 *ClaireReference   = To_Reference(t2)
        /* noccur = 4 */
        if (g0069.Arg == CTRUE) /* If:4 */{ 
          Result = CFALSE
          } else {
          Result = ToType(t).Included(ToType(F__at_Reference(g0069,g0069.Args,l.At((g0069.Index+1)-1))))
          /* If-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (t2.Isa.IsIn(C_type) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0070 *ClaireType   = ToType(t2)
        /* noccur = 2 */
        if (t.Isa.IsIn(C_type) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var g0071 *ClaireType   = ToType(t)
            /* noccur = 1 */
            Result = g0071.Included(g0070)
            /* Let-5 */} 
          } else {
          Result = ToBoolean(OBJ(F_CALL(ToProperty(C_less_ask.Id()),ARGS(t.ToEID(),EID{g0070.Id(),0}))))
          /* If-4 */} 
        /* Let-3 */} 
      } else {
      Result = ToBoolean(OBJ(F_CALL(ToProperty(C_less_ask.Id()),ARGS(t.ToEID(),t2.ToEID()))))
      /* If-2 */} 
    return Result} 
  
// The EID go function for: tmatch? @ any (throw: false) 
func E_tmatch_ask_any (t EID,t2 EID,l EID) EID { 
    return EID{/*(sm for tmatch? @ any= boolean)*/ F_tmatch_ask_any(ANY(t),ANY(t2),ToList(OBJ(l)) ).Id(),0}} 
  
// find the restriction
/* {1} OPT.The go function for: find_which(p:property,n:integer,c:class) [] */
func F_find_which_property (p *ClaireProperty ,n int,c *ClaireClass ) *ClaireObject  { 
    // procedure body with s =  
var Result *ClaireObject  
    if (p.Dictionary == CTRUE) /* If:2 */{ 
      Result = F_hashget_class(c,p)
      } else {
      /* For:3 */{ 
        var r *ClaireAny  
        _ = r
        Result= ToObject(CFALSE.Id())
        for _,r = range(p.Definition.ValuesO())/* loop:4 */{ 
          if (F_matching_ask_list(ToRestriction(r).Domain,n,ClEnv.Index) == CTRUE) /* If:5 */{ 
             /*v = Result, s =object*/
Result = ToObject(r)
            break
            /* If-5 */} 
          /* loop-4 */} 
        /* For-3 */} 
      /* If-2 */} 
    return Result} 
  
// The EID go function for: find_which @ property (throw: false) 
func E_find_which_property (p EID,n EID,c EID) EID { 
    return EID{/*(sm for find_which @ property= object)*/ F_find_which_property(ToProperty(OBJ(p)),INT(n),ToClass(OBJ(c)) ).Id(),0}} 
  
/* {1} OPT.The go function for: find_which(l:list,c:class,n:integer,m:integer) [] */
func F_find_which_list (l *ClaireList ,c *ClaireClass ,n int,m int) *ClaireObject  { 
    // procedure body with s =  
var Result *ClaireObject  
    /* For:2 */{ 
      var r *ClaireAny  
      _ = r
      Result= ToObject(CFALSE.Id())
      var r_support *ClaireList  
      r_support = l
      r_len := r_support.Length()
      for i_it := 0; i_it < r_len; i_it++ { 
        r = r_support.At(i_it)
        if (F_matching_ask_list(ToRestriction(r).Domain,n,m) == CTRUE) /* If:4 */{ 
           /*v = Result, s =object*/
Result = ToObject(r)
          break
          /* If-4 */} 
        /* loop-3 */} 
      /* For-2 */} 
    return Result} 
  
// The EID go function for: find_which @ list (throw: false) 
func E_find_which_list (l EID,c EID,n EID,m EID) EID { 
    return EID{/*(sm for find_which @ list= object)*/ F_find_which_list(ToList(OBJ(l)),
      ToClass(OBJ(c)),
      INT(n),
      INT(m) ).Id(),0}} 
  
// special version for super
/* {1} OPT.The go function for: find_which(c:class,l:list,n:integer,m:integer) [] */
func F_find_which_class (c *ClaireClass ,l *ClaireList ,n int,m int) *ClaireObject  { 
    // procedure body with s =  
var Result *ClaireObject  
    /* For:2 */{ 
      var r *ClaireAny  
      _ = r
      Result= ToObject(CFALSE.Id())
      var r_support *ClaireList  
      r_support = l
      r_len := r_support.Length()
      for i_it := 0; i_it < r_len; i_it++ { 
        r = r_support.At(i_it)
        if ((ToType(c.Id()).Included(ToType(ToRestriction(r).Domain.ValuesO()[1-1])) == CTRUE) && 
            (F_matching_ask_list(ToRestriction(r).Domain,n,m) == CTRUE)) /* If:4 */{ 
           /*v = Result, s =object*/
Result = ToObject(r)
          break
          /* If-4 */} 
        /* loop-3 */} 
      /* For-2 */} 
    return Result} 
  
// The EID go function for: find_which @ class (throw: false) 
func E_find_which_class (c EID,l EID,n EID,m EID) EID { 
    return EID{/*(sm for find_which @ class= object)*/ F_find_which_class(ToClass(OBJ(c)),
      ToList(OBJ(l)),
      INT(n),
      INT(m) ).Id(),0}} 
  