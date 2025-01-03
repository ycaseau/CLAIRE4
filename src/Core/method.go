/***** CLAIRE Compilation of file /Users/ycaseau/Dropbox/src/clairev4.12/src/meta/method.cl 
         [version 4.1.4 / safety 5] Friday 01-03-2025 16:21:02 *****/

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
/* The go function for: close(r:slot) [status=0] */
func F_close_slot (r *ClaireSlot) *ClaireSlot { 
  F_insert_definition_property(r.Selector,ToRestriction(r.Id()))
  return  r
  } 

// The EID go function for: close @ slot (throw: false) 
func E_close_slot (r EID) EID { 
  return EID{F_close_slot(ToSlot(OBJ(r)) ).Id(),0}} 

/* The go function for: close(r:method) [status=0] */
func F_close_method (r *ClaireMethod) *ClaireMethod { 
  F_insert_definition_property(r.Selector,ToRestriction(r.Id()))
  return  r
  } 

// The EID go function for: close @ method (throw: false) 
func E_close_method (r EID) EID { 
  return EID{F_close_method(ToMethod(OBJ(r)) ).Id(),0}} 

// Claire 4: introduce the capacity to set the comment automatically at compile time
/* The go function for: attach(r:method,s:string) [status=0] */
func F_attach_method (r *ClaireMethod,s *ClaireString) *ClaireMethod { 
  r.Comment = F_append_string(MakeString("defined in file "),s)
  return  F_close_method(r)
  } 

// The EID go function for: attach @ method (throw: false) 
func E_attach_method (r EID,s EID) EID { 
  return EID{F_attach_method(ToMethod(OBJ(r)),ToString(OBJ(s)) ).Id(),0}} 

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
/* The go function for: eval_message(self:property,r:object,start:integer,int?:boolean) [status=1] */
func F_eval_message_property (self *ClaireProperty,r *ClaireObject,start int,int_ask *ClaireBoolean) EID { 
  var Result EID
  if (r.Isa.Id() == C_method.Id()) { 
    if (ClEnv.Debug_I != -1) { 
      Result = F_execute_method(ToMethod(r.Id()),start,int_ask)
      } else {
      { var m *ClaireMethod = ToMethod(r.Id())
        if (m.Formula.Id() != CNULL) { 
          { var retour int = ClEnv.Base
            ClEnv.Base= start
            F_stack_add(m.Formula.Dimension)
            { 
              var val EID
              val = EVAL(m.Formula.Body)
              if ErrorIn(val) {Result = val
              } else {
              ClEnv.Base= retour
              ClEnv.Index= start
              Result = val}
              } 
            } 
          } else {
          Result = F_stack_apply_function(m.Functional,start,ClEnv.Index)
          } 
        } 
      } 
    }  else if ((r.Id().Isa.Id() == C_slot.Id()) && 
      (ClEnv.Index == (start+1))) { 
    { var val *ClaireAny = F_get_slot(ToSlot(r.Id()),ToObject(OBJ(ClEnv.EvalStack[start])))
      ClEnv.Index= start
      { var n int = ClEnv.Trace_I
        if ((n > 0) && 
            ((self.Trace_I+ClEnv.Verbose) > 4)) { 
          ClEnv.Trace_I = 0
          PRINC("read: ")
          Result = F_print_any(self.Id())
          if !ErrorIn(Result) {
          PRINC("(")
          Result = F_CALL(C_print,ARGS(ClEnv.EvalStack[start]))
          if !ErrorIn(Result) {
          PRINC(") = ")
          Result = F_CALL(C_print,ARGS(val.ToEID()))
          if !ErrorIn(Result) {
          PRINC("\n")
          Result = EVOID
          }}}
          if !ErrorIn(Result) {
          { 
            var va_arg1 *ClaireEnvironment
            var va_arg2 int
            va_arg1 = ClEnv
            va_arg2 = n
            va_arg1.Trace_I = va_arg2
            Result = EID{C__INT,IVAL(va_arg2)}
            } 
          }
          } else {
          Result = EID{CFALSE.Id(),0}
          } 
        } 
      if !ErrorIn(Result) {
      Result = val.ToEID()
      }
      } 
    } else {
    Result = F_noeval_message_property2(self,start)
    } 
  return Result} 

// The EID go function for: eval_message @ property (throw: true) 
func E_eval_message_property (self EID,r EID,start EID,int_ask EID) EID { 
  return F_eval_message_property(ToProperty(OBJ(self)),
    ToObject(OBJ(r)),
    INT(start),
    ToBoolean(OBJ(int_ask)) )} 

/* The go function for: noeval_message(self:property,start:integer) [status=1] */
func F_noeval_message_property2 (self *ClaireProperty,start int) EID { 
  var Result EID
  { var l *ClaireList = F_get_args_integer(start)
    if (ClEnv.Debug_I != -1) { 
      Result = F_push_debug_property(self,(ClEnv.Index-start),start)
      } else {
      Result = EID{CFALSE.Id(),0}
      } 
    if !ErrorIn(Result) {
    Result = ToException(C_selector_error.Make(self.Id(),l.Id())).Close()
    if !ErrorIn(Result) {
    Result = EID{CNIL.Id(),0}
    }}
    } 
  return Result} 

// The EID go function for: noeval_message @ property (throw: true) 
func E_noeval_message_property2 (self EID,start EID) EID { 
  return F_noeval_message_property2(ToProperty(OBJ(self)),INT(start) )} 

// a generic method : same as previously but (1) can be called by other methods
// and (2) takes care of the debugging piece, which implies a slower run (GC)
/* The go function for: execute(self:method,start:integer,int?:boolean) [status=1] */
func F_execute_method (self *ClaireMethod,start int,int_ask *ClaireBoolean) EID { 
  var Result EID
  { var n int = self.Domain.Length()
    if (self.Formula.Id() != CNULL) { 
      { var retour int = ClEnv.Base
        { var st_ask *ClaireBoolean = MakeBoolean((ClEnv.Debug_I != -1) && ((int_ask == CTRUE) || 
              (self.Module_I.Status != 4)))
          ClEnv.Base= start
          F_stack_add(self.Formula.Dimension)
          if (st_ask == CTRUE) { 
            Result = F_push_debug_property(self.Selector,n,start)
            } else {
            Result = EID{CFALSE.Id(),0}
            } 
          if !ErrorIn(Result) {
          { var val *ClaireAny
            var try_1 EID
            try_1 = EVAL(self.Formula.Body)
            if ErrorIn(try_1) {Result = try_1
            } else {
            val = ANY(try_1)
            if (st_ask == CTRUE) { 
              Result = F_pop_debug_property(self.Selector,0,val)
              } else {
              Result = EID{CFALSE.Id(),0}
              } 
            if !ErrorIn(Result) {
            ClEnv.Base= retour
            ClEnv.Index= start
            if ((ClEnv.Debug_I != -1) && 
                (self.Range.Contains(val) != CTRUE)) { 
              Result = ToException(C_range_error.Make(self.Id(),val,self.Range.Id())).Close()
              } else {
              Result = EID{CFALSE.Id(),0}
              } 
            if !ErrorIn(Result) {
            Result = val.ToEID()
            }}
            }
            } 
          }
          } 
        } 
      } else {
      { var st_ask *ClaireBoolean = MakeBoolean((ClEnv.Debug_I != -1) && ((int_ask == CTRUE) || 
            (self.Module_I.Status != 3)) && (self.Selector.Id() != C_debug.Id()))
        { var i int = ClEnv.Index
          if (st_ask == CTRUE) { 
            Result = F_push_debug_property(self.Selector,n,start)
            } else {
            Result = EID{CFALSE.Id(),0}
            } 
          if !ErrorIn(Result) {
          { var val *ClaireAny
            var try_2 EID
            try_2 = F_stack_apply_function(self.Functional,start,i)
            if ErrorIn(try_2) {Result = try_2
            } else {
            val = ANY(try_2)
            if (st_ask == CTRUE) { 
              Result = F_pop_debug_property(self.Selector,0,val)
              } else {
              Result = EID{CFALSE.Id(),0}
              } 
            if !ErrorIn(Result) {
            if ((ClEnv.Debug_I != -1) && 
                (self.Range.Contains(val) != CTRUE)) { 
              Result = ToException(C_range_error.Make(self.Id(),val,self.Range.Id())).Close()
              } else {
              Result = EID{CFALSE.Id(),0}
              } 
            if !ErrorIn(Result) {
            Result = val.ToEID()
            }}
            }
            } 
          }
          } 
        } 
      } 
    } 
  return Result} 

// The EID go function for: execute @ method (throw: true) 
func E_execute_method (self EID,start EID,int_ask EID) EID { 
  return F_execute_method(ToMethod(OBJ(self)),INT(start),ToBoolean(OBJ(int_ask)) )} 

// the evaluator is open coded
/* The go function for: eval(self:any) [status=1] */
func F_eval_any (self *ClaireAny) EID { 
  var Result EID
  Result = EVAL(self)
  return Result} 

// The EID go function for: eval @ list<type_expression>(any) (throw: true) 
func E_eval_any (self EID) EID { 
  return F_eval_any(ANY(self) )} 

// this is the standard evaluation
// self_eval(self:object) : any -> self
// reads an inline definition for a method
// notice that it does not return an error
/* The go function for: inlineok?(self:method,s:string) [status=0] */
func F_inlineok_ask_method (self *ClaireMethod,s *ClaireString) *ClaireMethod { 
  { 
    var Unused_H EID
    h_index := ClEnv.Index
    h_base := ClEnv.Base
    { var p *ClaireProperty = C_read
      { var l *ClaireAny
        var try_1 EID
        try_1 = F_call_property(p,ToList(MakeConstantList((s).Id()).Id()))
        if ErrorIn(try_1) {Unused_H = try_1
        } else {
        l = ANY(try_1)
        self.Inline_ask = CTRUE
        { 
          var va_arg1 *ClaireMethod
          var va_arg2 *ClaireLambda
          va_arg1 = self
          va_arg2 = ToLambda(l)
          va_arg1.Formula = va_arg2
          Unused_H = EID{va_arg2.Id(),0}
          } 
        }
        } 
      } 
    if ErrorIn(Unused_H){ 
      ClEnv.Index = h_index
      ClEnv.Base = h_base
      F_tformat_string(MakeString("---- WARNING: inline definition of ~S is wrong\n"),0,MakeConstantList(self.Id()))
      } 
    } 
  return  self
  } 

// The EID go function for: inlineok? @ method (throw: false) 
func E_inlineok_ask_method (self EID,s EID) EID { 
  return EID{F_inlineok_ask_method(ToMethod(OBJ(self)),ToString(OBJ(s)) ).Id(),0}} 

// reads a lambda - may return an error
/* The go function for: read_lambda(s:string) [status=1] */
func F_read_lambda_string (s *ClaireString) EID { 
  var Result EID
  { 
    h_index := ClEnv.Index
    h_base := ClEnv.Base
    { var p *ClaireProperty = C_read
      { var l *ClaireAny
        var try_1 EID
        try_1 = F_call_property(p,ToList(MakeConstantList((s).Id()).Id()))
        if ErrorIn(try_1) {Result = try_1
        } else {
        l = ANY(try_1)
        if (l.Isa.IsIn(C_lambda) == CTRUE) { 
          { var g0000 *ClaireLambda = ToLambda(l)
            Result = EID{g0000.Id(),0}
            } 
          } else {
          Result = ToException(C_general_error.Make(MakeString("compiled lambda error with ~S (not a lambda!)").Id(),MakeConstantList((s).Id()).Id())).Close()
          } 
        }
        } 
      } 
    if ErrorIn(Result){ 
      ClEnv.Index = h_index
      ClEnv.Base = h_base
      Result = ToException(C_general_error.Make(MakeString("compiled lambda parse error with ~S").Id(),MakeConstantList((s).Id()).Id())).Close()
      } 
    } 
  return Result} 

// The EID go function for: read_lambda @ string (throw: true) 
func E_read_lambda_string (s EID) EID { 
  return F_read_lambda_string(ToString(OBJ(s)) )} 

// ****************************************************************
// *    Part 2: Update methods                                    *
// ****************************************************************
//get/put for a slot: should be inline
/* The go function for: get(s:slot,x:object) [status=0] */
func F_get_slot (s *ClaireSlot,x *ClaireObject) *ClaireAny { 
  return  x.SlotGet(s.Index,s.Srange)
  } 

// The EID go function for: get @ slot (throw: false) 
func E_get_slot (s EID,x EID) EID { 
  return F_get_slot(ToSlot(OBJ(s)),ToObject(OBJ(x)) ).ToEID()} 

/* The go function for: put(s:slot,x:object,y:any) [status=0] */
func F_put_slot (s *ClaireSlot,x *ClaireObject,y *ClaireAny) *ClaireAny { 
  return  F_store_object(x,
    s.Index,
    s.Srange,
    y,
    s.Selector.Store_ask)
  } 

// The EID go function for: put @ slot (throw: false) 
func E_put_slot (s EID,x EID,y EID) EID { 
  return F_put_slot(ToSlot(OBJ(s)),ToObject(OBJ(x)),ANY(y) ).ToEID()} 

// reading a value from a property (unknown is allowed)
// when unknown is not allowed, we use read which is defined in Kernel
/* The go function for: get(self:property,x:object) [status=0] */
func F_get_property (self *ClaireProperty,x *ClaireObject) *ClaireAny { 
  var Result *ClaireAny
  { var s *ClaireObject = F__at_property1(self,x.Id().Isa)
    if (C_slot.Id() == s.Isa.Id()) { 
      { var g0002 *ClaireSlot = ToSlot(s.Id())
        Result = x.SlotGet(g0002.Index,g0002.Srange)
        } 
      } else {
      Result = CNULL
      } 
    } 
  return Result} 

// The EID go function for: get @ property (throw: false) 
func E_get_property (self EID,x EID) EID { 
  return F_get_property(ToProperty(OBJ(self)),ToObject(OBJ(x)) ).ToEID()} 

// a more general value that is useful for types
/* The go function for: funcall(self:property,x:any) [status=1] */
func F_funcall_property (self *ClaireProperty,x *ClaireAny) EID { 
  var Result EID
  { var s *ClaireObject = F__at_property1(self,x.Isa)
    if (C_slot.Id() == s.Isa.Id()) { 
      { var g0004 *ClaireSlot = ToSlot(s.Id())
        Result = ToObject(x).SlotGet(g0004.Index,g0004.Srange).ToEID()
        } 
      }  else if (C_method.Id() == s.Isa.Id()) { 
      { var g0005 *ClaireMethod = ToMethod(s.Id())
        Result = F_funcall_method1(g0005,x)
        } 
      } else {
      Result = EID{CNULL,0}
      } 
    } 
  return Result} 

// The EID go function for: funcall @ property (throw: true) 
func E_funcall_property (self EID,x EID) EID { 
  return F_funcall_property(ToProperty(OBJ(self)),ANY(x) )} 

// verifying
/* The go function for: hold?(self:property,x:object,y:any) [status=0] */
func F_hold_ask_property (self *ClaireProperty,x *ClaireObject,y *ClaireAny) *ClaireBoolean { 
  var Result *ClaireBoolean
  { var s *ClaireObject = F__at_property1(self,x.Id().Isa)
    if (C_slot.Id() == s.Isa.Id()) { 
      { var g0007 *ClaireSlot = ToSlot(s.Id())
        { var z *ClaireAny = x.SlotGet(g0007.Index,g0007.Srange)
          if (C_set.Id() == z.Isa.Id()) { 
            { var g0008 *ClaireSet = ToSet(z)
              Result = g0008.Contain_ask(y)
              } 
            } else {
            Result = Equal(y,z)
            } 
          } 
        } 
      } else {
      Result = CFALSE
      } 
    } 
  return Result} 

// The EID go function for: hold? @ property (throw: false) 
func E_hold_ask_property (self EID,x EID,y EID) EID { 
  return EID{F_hold_ask_property(ToProperty(OBJ(self)),ToObject(OBJ(x)),ANY(y) ).Id(),0}} 

// writing a single value into a slot & trigger rules
//  write = check + put + put_inverse + propagate
//  if_write = put + put_inverse + propagate  (propagate => if_write)
//  update = put + put_inverse
// note in CLAIRE 4: with no inverse/store write_fast, defined in Kernel, works better
/* The go function for: write(self:property,x:object,y:any) [status=1] */
func F_write_property (self *ClaireProperty,x *ClaireObject,y *ClaireAny) EID { 
  var Result EID
  { var s *ClaireObject = F__at_property1(self,x.Id().Isa)
    if (C_slot.Id() == s.Isa.Id()) { 
      { var g0011 *ClaireSlot = ToSlot(s.Id())
        if (g0011.Range.Contains(y) != CTRUE) { 
          Result = F_range_is_wrong_slot(g0011,y)
          }  else if ((self.Open < 1) && 
            (x.SlotGet(g0011.Index,g0011.Srange) != CNULL)) { 
          Result = ToException(C_general_error.Make(MakeString("[132] Cannot change ~S(~S)").Id(),MakeConstantList(self.Id(),x.Id()).Id())).Close()
          }  else if ((self.IfWrite != CNULL) && 
            (self.Multivalued_ask != CTRUE)) { 
          Result = F_fastcall_relation2(ToRelation(self.Id()),x.Id(),y)
          } else {
          Result = F_update_property(self,
            x,
            g0011.Index,
            g0011.Srange,
            y)
          } 
        } 
      } else {
      Result = ToException(C_selector_error.Make(self.Id(),MakeConstantList(x.Id()).Id())).Close()
      } 
    } 
  if !ErrorIn(Result) {
  Result = y.ToEID()
  }
  return Result} 

// The EID go function for: write @ property (throw: true) 
func E_write_property (self EID,x EID,y EID) EID { 
  return F_write_property(ToProperty(OBJ(self)),ToObject(OBJ(x)),ANY(y) )} 

// the value does not belong to the range: error!
/* The go function for: range_is_wrong(self:slot,y:any) [status=1] */
func F_range_is_wrong_slot (self *ClaireSlot,y *ClaireAny) EID { 
  var Result EID
  Result = ToException(C_range_error.Make(self.Id(),y,self.Range.Id())).Close()
  return Result} 

// The EID go function for: range_is_wrong @ slot (throw: true) 
func E_range_is_wrong_slot (self EID,y EID) EID { 
  return F_range_is_wrong_slot(ToSlot(OBJ(self)),ANY(y) )} 

// to remove
/* The go function for: put(p:property,x:object,n:integer,s:class,y:any) [status=1] */
func F_put_property1 (p *ClaireProperty,x *ClaireObject,n int,s *ClaireClass,y *ClaireAny) EID { 
  var Result EID
  Result = F_update_property(p,
    x,
    n,
    s,
    y)
  return Result} 

// The EID go function for: put @ list<type_expression>(property, object, integer, class, any) (throw: true) 
func E_put_property1 (p EID,x EID,n EID,s EID,y EID) EID { 
  return F_put_property1(ToProperty(OBJ(p)),
    ToObject(OBJ(x)),
    INT(n),
    ToClass(OBJ(s)),
    ANY(y) )} 

// update (method called by the compiler)     // v3.0.20: renamed from put !
// update = put + put_inverse  (complex links) .. it does not trigger the rules (if_write)
// update uses two satellite methods: update+ and update-
// CLAIRE 4: inverse management only applies with set multivalued properties
/* The go function for: mClaire/update(p:property,x:object,n:integer,s:class,y:any) [status=1] */
func F_update_property (p *ClaireProperty,x *ClaireObject,n int,s *ClaireClass,y *ClaireAny) EID { 
  var Result EID
  { var old *ClaireAny = x.SlotGet(n,s)
    if (ClEnv.Verbose == 8) { 
      F_tformat_string(MakeString("update ~S(~S) old = ~S\n"),0,MakeConstantList(p.Id(),x.Id(),old))
      } 
    if (p.Multivalued_ask.Id() == CTRUE.Id()) { 
      if (INT(F_CALL(C_length,ARGS(old.ToEID()))) > 0) { 
        { var v *ClaireSet = ToType(CEMPTY.Id()).EmptySet()
          if (ANY(F_CALL(C_of,ARGS(old.ToEID()))) != C_void.Id()) { 
            v.Cast_I(ToType(OBJ(F_CALL(C_of,ARGS(old.ToEID())))))
            } 
          F_store_object(x,
            n,
            s,
            v.Id(),
            p.Store_ask)
          } 
        } 
      { var r *ClaireRelation = p.Inverse
        if (r.Id() == CNULL) { 
          
          } else {
          { 
            var z *ClaireAny
            _ = z
            var z_support *ClaireSet
            z_support = ToSet(old)
            for i_it := 0; i_it < z_support.Count; i_it++ { 
              z = z_support.At(i_it)
              F_update_dash_relation(r,z,x.Id())
              } 
            } 
          } 
        } 
      { 
        var z *ClaireAny
        _ = z
        Result= EID{CFALSE.Id(),0}
        var z_support *ClaireSet
        z_support = ToSet(y)
        for i_it := 0; i_it < z_support.Count; i_it++ { 
          z = z_support.At(i_it)
          var loop_1 EID
          _ = loop_1
          loop_1 = F_add_I_property(p,x,n,z)
          if ErrorIn(loop_1) {Result = loop_1
          break
          } else {
          }
          } 
        } 
      }  else if (Equal(old,y) != CTRUE) { 
      { var r *ClaireRelation = p.Inverse
        if (r.Id() == CNULL) { 
          
          }  else if ((old != CNULL) && 
            ((r.Id() != p.Id()) || 
                (Equal(x.Id(),old) != CTRUE))) { 
          F_update_dash_relation(r,old,x.Id())
          } 
        } 
      F_store_object(x,
        n,
        s,
        y,
        p.Store_ask)
      Result = F_update_plus_relation(ToRelation(p.Id()),x.Id(),y)
      } else {
      Result = EID{CFALSE.Id(),0}
      } 
    if !ErrorIn(Result) {
    Result = y.ToEID()
    }
    } 
  return Result} 

// The EID go function for: mClaire/update @ property (throw: true) 
func E_update_property (p EID,x EID,n EID,s EID,y EID) EID { 
  return F_update_property(ToProperty(OBJ(p)),
    ToObject(OBJ(x)),
    INT(n),
    ToClass(OBJ(s)),
    ANY(y) )} 

// this method checks the correctness of the inverse from a global view.
/* The go function for: update+(self:relation,x:any,y:any) [status=1] */
func F_update_plus_relation (self *ClaireRelation,x *ClaireAny,y *ClaireAny) EID { 
  var Result EID
  { var r *ClaireRelation = self.Inverse
    if ((r.Id() != CNULL) && 
        ((r.Id() != self.Id()) || 
            (Equal(x,y) != CTRUE))) { 
      if (r.Isa.IsIn(C_property) == CTRUE) { 
        { var g0013 *ClaireProperty = ToProperty(r.Id())
          { var s *ClaireObject = F__at_property1(g0013,y.Isa)
            if (C_slot.Id() == s.Isa.Id()) { 
              { var g0014 *ClaireSlot = ToSlot(s.Id())
                { var old_y *ClaireAny = F_get_slot(g0014,ToObject(y))
                  if (g0013.Multivalued_ask.Id() != CFALSE.Id()) { 
                    Result = EID{F_Core_add_value_I_property(g0013,
                      ToObject(y),
                      g0014.Index,
                      ToSet(old_y),
                      x).Id(),0}
                    } else {
                    Result = F_CALL(C_store,ARGS(y.ToEID(),
                      EID{C__INT,IVAL(g0014.Index)},
                      EID{g0014.Srange.Id(),0},
                      x.ToEID(),
                      EID{g0013.Store_ask.Id(),0}))
                    } 
                  } 
                } 
              } else {
              Result = ToException(C_general_error.Make(MakeString("[133] Inversion of ~S(~S,~S) impossible").Id(),MakeConstantList(self.Id(),x,y).Id())).Close()
              } 
            } 
          } 
        }  else if (C_table.Id() == r.Isa.Id()) { 
        { var g0016 *ClaireTable = ToTable(r.Id())
          { var old_v *ClaireAny = F_get_table(g0016,y)
            if (g0016.Multivalued_ask.Id() != CFALSE.Id()) { 
              Result = EID{F_Core_add_value_I_table(g0016,y,ToSet(old_v),x).Id(),0}
              } else {
              if (old_v != CNULL) { 
                F_update_dash_relation(self,old_v,y)
                } 
              F_put_table(g0016,y,x)
              Result = EVOID
              } 
            } 
          } 
        } else {
        Result = EID{CFALSE.Id(),0}
        } 
      } else {
      Result = EID{CFALSE.Id(),0}
      } 
    } 
  return Result} 

// The EID go function for: update+ @ relation (throw: true) 
func E_update_plus_relation (self EID,x EID,y EID) EID { 
  return F_update_plus_relation(ToRelation(OBJ(self)),ANY(x),ANY(y) )} 

// this methods deletes a value in the inverse of a global_relation
/* The go function for: update-(r:relation,x:any,y:any) [status=0] */
func F_update_dash_relation (r *ClaireRelation,x *ClaireAny,y *ClaireAny)  { 
  if (r.Isa.IsIn(C_property) == CTRUE) { 
    { var g0017 *ClaireProperty = ToProperty(r.Id())
      { var s *ClaireObject = F__at_property1(g0017,x.Isa)
        if (C_slot.Id() == s.Isa.Id()) { 
          { var g0018 *ClaireSlot = ToSlot(s.Id())
            { var l *ClaireAny = F_get_slot(g0018,ToObject(x))
              { var v *ClaireAny
                if (C_set.Id() == l.Isa.Id()) { 
                  { var g0019 *ClaireSet = ToSet(l)
                    { var arg_1 *ClaireSet
                      if (g0017.Store_ask == CTRUE) { 
                        arg_1 = g0019.Copy()
                        } else {
                        arg_1 = g0019
                        } 
                      v = arg_1.Delete(y).Id()
                      } 
                    } 
                  } else {
                  v = CNULL
                  } 
                F_put_slot(g0018,ToObject(x),v)
                } 
              } 
            } 
          } 
        } 
      } 
    }  else if (C_table.Id() == r.Isa.Id()) { 
    { var g0021 *ClaireTable = ToTable(r.Id())
      { var l *ClaireAny = F_get_table(g0021,x)
        { var v *ClaireAny
          if (C_set.Id() == l.Isa.Id()) { 
            { var g0022 *ClaireSet = ToSet(l)
              { var arg_2 *ClaireSet
                if (g0021.Store_ask == CTRUE) { 
                  arg_2 = g0022.Copy()
                  } else {
                  arg_2 = g0022
                  } 
                v = arg_2.Delete(y).Id()
                } 
              } 
            } else {
            v = CNULL
            } 
          F_put_table(g0021,x,v)
          } 
        } 
      } 
    } 
  } 

// The EID go function for: update- @ relation (throw: false) 
func E_update_dash_relation (r EID,x EID,y EID) EID { 
  F_update_dash_relation(ToRelation(OBJ(r)),ANY(x),ANY(y) )
  return EVOID} 

// this methods adds a value to a multi-slot (used by the compiler)
// this is the multi-valued equivalent of update - we know self to be multivalued (hence a set in Claire 4)
/* The go function for: add!(self:property,x:object,n:integer,y:any) [status=1] */
func F_add_I_property (self *ClaireProperty,x *ClaireObject,n int,y *ClaireAny) EID { 
  var Result EID
  if (self.IfWrite != CNULL) { 
    Result = F_fastcall_relation2(ToRelation(self.Id()),x.Id(),y)
    } else {
    { var s1 *ClaireSet = ToSet(x.SlotGet(n,C_object))
      if (F_Core_add_value_I_property(self,
        x,
        n,
        s1,
        y) == CTRUE) { 
        Result = F_update_plus_relation(ToRelation(self.Id()),x.Id(),y)
        } else {
        Result = EID{CFALSE.Id(),0}
        } 
      } 
    } 
  return Result} 

// The EID go function for: add! @ property (throw: true) 
func E_add_I_property (self EID,x EID,n EID,y EID) EID { 
  return F_add_I_property(ToProperty(OBJ(self)),
    ToObject(OBJ(x)),
    INT(n),
    ANY(y) )} 

// this methods adds a value to a multi-slot (internal form)
// this is the multi-valued equivalent of put
// return true if the set is actually changed (y added to s)
/* The go function for: add_value!(self:property,x:object,n:integer,s1:set,y:any) [status=0] */
func F_Core_add_value_I_property (self *ClaireProperty,x *ClaireObject,n int,s1 *ClaireSet,y *ClaireAny) *ClaireBoolean { 
  var Result *ClaireBoolean
  if (s1.Contain_ask(y) != CTRUE) { 
    { var s2 *ClaireSet
      { var arg_1 *ClaireAny
        if (self.Store_ask == CTRUE) { 
          arg_1 = s1.Copy().Id()
          } else {
          arg_1 = s1.Id()
          } 
        s2 = ToSet(arg_1).AddFast(y)
        } 
      F_store_object(x,
        n,
        C_object,
        s2.Id(),
        self.Store_ask)
      Result = CTRUE
      } 
    } else {
    Result = CFALSE
    } 
  return Result} 

// The EID go function for: add_value! @ property (throw: false) 
func E_Core_add_value_I_property (self EID,x EID,n EID,s1 EID,y EID) EID { 
  return EID{F_Core_add_value_I_property(ToProperty(OBJ(self)),
    ToObject(OBJ(x)),
    INT(n),
    ToSet(OBJ(s1)),
    ANY(y) ).Id(),0}} 

// same method with error checking
/* The go function for: add(self:property,x:object,y:any) [status=1] */
func F_add_property (self *ClaireProperty,x *ClaireObject,y *ClaireAny) EID { 
  var Result EID
  { var s *ClaireObject = F__at_property1(self,x.Id().Isa)
    if (F_boolean_I_any(s.Id()).Id() != CTRUE.Id()) { 
      Result = ToException(C_selector_error.Make(self.Id(),MakeConstantList(x.Id()).Id())).Close()
      }  else if (self.Multivalued_ask != CTRUE) { 
      Result = ToException(C_general_error.Make(MakeString("[134] Cannot apply add to ~S").Id(),MakeConstantList(self.Id()).Id())).Close()
      }  else if (F_member_type(ToRestriction(s.Id()).Range).Contains(y) == CTRUE) { 
      if (self.IfWrite != CNULL) { 
        Result = F_fastcall_relation2(ToRelation(self.Id()),x.Id(),y)
        } else {
        Result = F_add_I_property(self,x,ToSlot(s.Id()).Index,y)
        } 
      } else {
      Result = F_range_is_wrong_slot(ToSlot(s.Id()),y)
      } 
    } 
  if !ErrorIn(Result) {
  Result = y.ToEID()
  }
  return Result} 

// The EID go function for: add @ property (throw: true) 
func E_add_property (self EID,x EID,y EID) EID { 
  return F_add_property(ToProperty(OBJ(self)),ToObject(OBJ(x)),ANY(y) )} 

// known ?
/* The go function for: known?(self:property,x:object) [status=0] */
func F_known_ask_property (self *ClaireProperty,x *ClaireObject) *ClaireBoolean { 
  var Result *ClaireBoolean
  { var s *ClaireObject = F__at_property1(self,x.Id().Isa)
    if (C_slot.Id() == s.Isa.Id()) { 
      { var g0024 *ClaireSlot = ToSlot(s.Id())
        Result = F__I_equal_any(x.SlotGet(g0024.Index,g0024.Srange),CNULL)
        } 
      } else {
      Result = CFALSE
      } 
    } 
  return Result} 

// The EID go function for: known? @ property (throw: false) 
func E_known_ask_property (self EID,x EID) EID { 
  return EID{F_known_ask_property(ToProperty(OBJ(self)),ToObject(OBJ(x)) ).Id(),0}} 

/* The go function for: unknown?(self:property,x:object) [status=0] */
func F_unknown_ask_property (self *ClaireProperty,x *ClaireObject) *ClaireBoolean { 
  var Result *ClaireBoolean
  { var s *ClaireObject = F__at_property1(self,x.Id().Isa)
    if (C_slot.Id() == s.Isa.Id()) { 
      { var g0026 *ClaireSlot = ToSlot(s.Id())
        Result = Equal(x.SlotGet(g0026.Index,g0026.Srange),CNULL)
        } 
      } else {
      Result = CTRUE
      } 
    } 
  return Result} 

// The EID go function for: unknown? @ property (throw: false) 
func E_unknown_ask_property (self EID,x EID) EID { 
  return EID{F_unknown_ask_property(ToProperty(OBJ(self)),ToObject(OBJ(x)) ).Id(),0}} 

// delete takes care of the inverse also
// assumes that self is multivalued -> should check !
/* The go function for: delete(self:property,x:object,y:any) [status=1] */
func F_delete_property (self *ClaireProperty,x *ClaireObject,y *ClaireAny) EID { 
  var Result EID
  { var s *ClaireObject = F__at_property1(self,x.Id().Isa)
    if (F_boolean_I_any(s.Id()).Id() != CTRUE.Id()) { 
      Result = ToException(C_selector_error.Make(self.Id(),MakeConstantList(x.Id()).Id())).Close()
      }  else if (self.Multivalued_ask.Id() == CTRUE.Id()) { 
      { var l1 *ClaireSet = ToSet(x.SlotGet(ToSlot(s.Id()).Index,C_object))
        { var l *ClaireSet
          { var arg_1 *ClaireSet
            if (self.Store_ask == CTRUE) { 
              arg_1 = l1.Copy()
              } else {
              arg_1 = l1
              } 
            l = arg_1.Delete(y)
            } 
          F_store_object(x,
            ToSlot(s.Id()).Index,
            C_object,
            l.Id(),
            self.Store_ask)
          { var r *ClaireRelation = self.Inverse
            if (r.Id() != CNULL) { 
              F_update_dash_relation(r,y,x.Id())
              } 
            } 
          Result = EID{l.Id(),0}
          } 
        } 
      } else {
      Result = EID{CFALSE.Id(),0}
      } 
    } 
  return Result} 

// The EID go function for: delete @ property (throw: true) 
func E_delete_property (self EID,x EID,y EID) EID { 
  return F_delete_property(ToProperty(OBJ(self)),ToObject(OBJ(x)),ANY(y) )} 

// erase is similar for mono-valued properties takes care of the inverse also
// v3.2.22: take care of multi-valued slot as well
/* The go function for: erase(self:property,x:object) [status=1] */
func F_erase_property (self *ClaireProperty,x *ClaireObject) EID { 
  var Result EID
  { var s *ClaireObject = F__at_property1(self,x.Id().Isa)
    if (F_boolean_I_any(s.Id()).Id() != CTRUE.Id()) { 
      Result = ToException(C_selector_error.Make(self.Id(),MakeConstantList(x.Id()).Id())).Close()
      } else {
      { var y *ClaireAny = x.SlotGet(ToSlot(s.Id()).Index,ToClass(OBJ(F_CALL(C_mClaire_srange,ARGS(EID{s.Id(),0})))))
        if (self.Multivalued_ask.Id() == CTRUE.Id()) { 
          { var r *ClaireRelation = self.Inverse
            if (r.Id() == CNULL) { 
              
              } else {
              { 
                var y1 *ClaireAny
                _ = y1
                var y1_support *ClaireSet
                y1_support = ToSet(y)
                for i_it := 0; i_it < y1_support.Count; i_it++ { 
                  y1 = y1_support.At(i_it)
                  F_update_dash_relation(r,y1,x.Id())
                  } 
                } 
              } 
            } 
          { var l *ClaireSet = ToSet(y).Empty()
            F_store_object(x,
              ToSlot(s.Id()).Index,
              C_object,
              l.Id(),
              self.Store_ask)
            Result = EID{l.Id(),0}
            } 
          } else {
          Result = F_CALL(C_store,ARGS(EID{x.Id(),0},
            EID{C__INT,IVAL(ToSlot(s.Id()).Index)},
            F_CALL(C_mClaire_srange,ARGS(EID{s.Id(),0})),
            F_CALL(C_default,ARGS(EID{s.Id(),0})),
            EID{self.Store_ask.Id(),0}))
          if !ErrorIn(Result) {
          { var r *ClaireRelation = self.Inverse
            if ((r.Id() != CNULL) && 
                (y != CNULL)) { 
              F_update_dash_relation(r,y,x.Id())
              } 
            } 
          Result = F_CALL(C_default,ARGS(EID{s.Id(),0}))
          }
          } 
        } 
      } 
    } 
  return Result} 

// The EID go function for: erase @ property (throw: true) 
func E_erase_property (self EID,x EID) EID { 
  return F_erase_property(ToProperty(OBJ(self)),ToObject(OBJ(x)) )} 

/* The go function for: set_range(p:property,c:class,r:type) [status=0] */
func F_set_range_property (p *ClaireProperty,c *ClaireClass,r *ClaireType)  { 
  { var s *ClaireSlot = ToSlot(F__at_property1(p,c).Id())
    s.Range = r
    s.Srange = ToTypeExpression(r.Id()).Class_I()
    } 
  } 

// The EID go function for: set_range @ property (throw: false) 
func E_set_range_property (p EID,c EID,r EID) EID { 
  F_set_range_property(ToProperty(OBJ(p)),ToClass(OBJ(c)),ToType(OBJ(r)) )
  return EVOID} 

// no longer needed because changing the range is not changing the prototype ?
// we should rather generate an error if the condition for dealing with
// defaults changes (TODO)
//        if (s.srange != any & s.srange != integer)
//           c.prototype[s.index] := 0))
// this method allows to bypass the storage mechanism - to be optimized ..
/* The go function for: put_store(self:property,x:object,y:any,b:boolean) [status=1] */
func F_put_store_property2 (self *ClaireProperty,x *ClaireObject,y *ClaireAny,b *ClaireBoolean) EID { 
  var Result EID
  { var s *ClaireObject = F__at_property1(self,x.Id().Isa)
    if (C_slot.Id() == s.Isa.Id()) { 
      { var g0028 *ClaireSlot = ToSlot(s.Id())
        { var z *ClaireAny = x.SlotGet(g0028.Index,g0028.Srange)
          if (Equal(z,y) != CTRUE) { 
            Result = F_store_object(x,
              g0028.Index,
              g0028.Srange,
              y,
              b).ToEID()
            } else {
            Result = EID{CFALSE.Id(),0}
            } 
          } 
        } 
      } else {
      Result = ToException(C_selector_error.Make(self.Id(),MakeConstantList(x.Id()).Id())).Close()
      } 
    } 
  return Result} 

// The EID go function for: put_store @ property (throw: true) 
func E_put_store_property2 (self EID,x EID,y EID,b EID) EID { 
  return F_put_store_property2(ToProperty(OBJ(self)),
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
/* The go function for: fastcall(r:relation,x:any,y:any) [status=1] */
func F_fastcall_relation2 (r *ClaireRelation,x *ClaireAny,y *ClaireAny) EID { 
  var Result EID
  { var f *ClaireAny = r.IfWrite
    if (f.Isa.IsIn(C_lambda) == CTRUE) { 
      Result = F_CALL(C_funcall,ARGS(f.ToEID(),x.ToEID(),y.ToEID()))
      } else {
      Result = F_funcall2(ToFunction(f),x,y)
      } 
    } 
  return Result} 

// The EID go function for: fastcall @ relation (throw: true) 
func E_fastcall_relation2 (r EID,x EID,y EID) EID { 
  return F_fastcall_relation2(ToRelation(OBJ(r)),ANY(x),ANY(y) )} 

// *********************************************************************
// *   Part 3: Management of definition(p)                             *
// *********************************************************************
// the dictionary slot
// insertion in the definition tree
/* The go function for: insert_definition(p:property,r:restriction) [status=0] */
func F_insert_definition_property (p *ClaireProperty,r *ClaireRestriction)  { 
  p.Definition = F_initialize_restriction1(r,ToTypeExpression(r.Domain.ValuesO()[0]).Class_I(),p.Definition)
  } 

// The EID go function for: insert_definition @ property (throw: false) 
func E_insert_definition_property (p EID,r EID) EID { 
  F_insert_definition_property(ToProperty(OBJ(p)),ToRestriction(OBJ(r)) )
  return EVOID} 

// insert a restriction with class-domain d into a property p
// claire4 : get rid of dispatcher
/* The go function for: initialize(x:restriction,d:class,l:list) [status=0] */
func F_initialize_restriction1 (x *ClaireRestriction,d *ClaireClass,l *ClaireList) *ClaireList { 
  var Result *ClaireList
  { var p *ClaireProperty = x.Selector
    if ((p.Restrictions.Length() == 5) && 
        (F_uniform_property(p) == CTRUE)) { 
      { 
        var r *ClaireRestriction
        _ = r
        var r_iter *ClaireAny
        for _,r_iter = range(p.Restrictions.ValuesO()){ 
          r = ToRestriction(r_iter)
          F_hashinsert_restriction(r)
          } 
        } 
      p.Dictionary = CTRUE
      } 
    if (p.Dictionary == CTRUE) { 
      if (F_uniform_restriction(x) == CTRUE) { 
        F_hashinsert_restriction(x)
        } else {
        p.Dictionary = CFALSE
        } 
      } 
    Result = F_initialize_restriction2(x,l)
    } 
  return Result} 

// The EID go function for: initialize @ list<type_expression>(restriction, class, list) (throw: false) 
func E_initialize_restriction1 (x EID,d EID,l EID) EID { 
  return EID{F_initialize_restriction1(ToRestriction(OBJ(x)),ToClass(OBJ(d)),ToList(OBJ(l)) ).Id(),0}} 

// only uniform properties can use the dictionary representation
/* The go function for: uniform(x:restriction) [status=0] */
func F_uniform_restriction (x *ClaireRestriction) *ClaireBoolean { 
  var Result *ClaireBoolean
  { var l *ClaireList = x.Domain
    { var n int = l.Length()
      { var arg_1 *ClaireAny
        { 
          var r *ClaireRestriction
          _ = r
          var r_iter *ClaireAny
          arg_1= CFALSE.Id()
          for _,r_iter = range(x.Selector.Restrictions.ValuesO()){ 
            r = ToRestriction(r_iter)
            var g0031I *ClaireBoolean
            { var arg_2 *ClaireBoolean
              { var l2 *ClaireList = r.Domain
                { 
                  var v_and8 *ClaireBoolean
                  
                  v_and8 = Equal(C_class.Id(),ToTypeExpression(l2.ValuesO()[0]).Isa.Id())
                  if (v_and8 == CFALSE) {arg_2 = CFALSE
                  } else { 
                    v_and8 = Equal(MakeInteger(l2.Length()).Id(),MakeInteger(n).Id())
                    if (v_and8 == CFALSE) {arg_2 = CFALSE
                    } else { 
                      v_and8 = F__I_equal_any(l2.ValuesO()[0],C_listargs.Id())
                      if (v_and8 == CFALSE) {arg_2 = CFALSE
                      } else { 
                        { var arg_3 *ClaireAny
                          { var i int = 2
                            { var g0030 int = n
                              arg_3= CFALSE.Id()
                              for (i <= g0030) { 
                                if ((Equal(l.ValuesO()[i-1],l2.ValuesO()[i-1]) != CTRUE) && 
                                    ((l.ValuesO()[i-1].Isa.Id() == C_class.Id()) || 
                                        ((l.ValuesO()[i-1].Isa.Id() != l2.ValuesO()[i-1].Isa.Id()) || 
                                          (F__equaltype_ask_any(ToType(l.ValuesO()[i-1]),ToType(l2.ValuesO()[i-1])) != CTRUE)))) { 
                                  arg_3 = CTRUE.Id()
                                  break
                                  } 
                                i = (i+1)
                                } 
                              } 
                            } 
                          v_and8 = F_not_any(arg_3)
                          } 
                        if (v_and8 == CFALSE) {arg_2 = CFALSE
                        } else { 
                          arg_2 = CTRUE} 
                        } 
                      } 
                    } 
                  } 
                } 
              g0031I = arg_2.Not
              } 
            if (g0031I == CTRUE) { 
              arg_1 = CTRUE.Id()
              break
              } 
            } 
          } 
        Result = F_not_any(arg_1)
        } 
      } 
    } 
  return Result} 

// The EID go function for: uniform @ restriction (throw: false) 
func E_uniform_restriction (x EID) EID { 
  return EID{F_uniform_restriction(ToRestriction(OBJ(x)) ).Id(),0}} 

// v3.3.36      
// v3.0.54 check that a uniform property only uses methods !
/* The go function for: uniform(p:property) [status=0] */
func F_uniform_property (p *ClaireProperty) *ClaireBoolean { 
  var Result *ClaireBoolean
  { 
    var v_and1 *ClaireBoolean
    
    { var arg_1 *ClaireAny
      { 
        var x *ClaireRestriction
        _ = x
        var x_iter *ClaireAny
        arg_1= CFALSE.Id()
        for _,x_iter = range(p.Restrictions.ValuesO()){ 
          x = ToRestriction(x_iter)
          if (C_method.Id() != x.Isa.Id()) { 
            arg_1 = CTRUE.Id()
            break
            } 
          } 
        } 
      v_and1 = F_not_any(arg_1)
      } 
    if (v_and1 == CFALSE) {Result = CFALSE
    } else { 
      v_and1 = F_uniform_restriction(ToRestriction(p.Restrictions.ValuesO()[0]))
      if (v_and1 == CFALSE) {Result = CFALSE
      } else { 
        Result = CTRUE} 
      } 
    } 
  return Result} 

// The EID go function for: uniform @ property (throw: false) 
func E_uniform_property (p EID) EID { 
  return EID{F_uniform_property(ToProperty(OBJ(p)) ).Id(),0}} 

// insert a restriction in a list with the good order
/* The go function for: initialize(x:restriction,l:list) [status=0] */
func F_initialize_restriction2 (x *ClaireRestriction,l *ClaireList) *ClaireList { 
  var Result *ClaireList
  { var l1 *ClaireList = CNIL
    { var i int = 1
      { var g0032 int = l.Length()
        for (i <= g0032) { 
          { var l2 *ClaireList = ToRestriction(l.At(i-1)).Domain
            if (F_tmatch_ask_list(x.Domain,l2) == CTRUE) { 
              if (F_tmatch_ask_list(l2,x.Domain) == CTRUE) { 
                ToArray(l.Id()).NthPut(i,x.Id())
                l1 = l
                
                break
                } else {
                l1 = ToList(ANY(l.Nth_plus(i,x.Id())))
                
                break
                } 
              }  else if ((F_tmatch_ask_list(l2,x.Domain) != CTRUE) && 
                ((F_join_list(x.Domain,l2) == CTRUE) && 
                  (x.Selector.Open <= 1))) { 
              F_tformat_string(MakeString("Note: ~S and ~S are conflicting\n"),2,MakeConstantList(l.At(0),x.Id()))
              } 
            } 
          i = (i+1)
          } 
        } 
      } 
    if (l1.Length() != 0) { 
      Result = l1
      } else {
      Result = l.AddFast(x.Id())
      } 
    } 
  return Result} 

// The EID go function for: initialize @ list<type_expression>(restriction, list) (throw: false) 
func E_initialize_restriction2 (x EID,l EID) EID { 
  return EID{F_initialize_restriction2(ToRestriction(OBJ(x)),ToList(OBJ(l)) ).Id(),0}} 

// definition of dictionary: standard hash-table
/* The go function for: hashinsert(m:restriction) [status=0] */
func F_hashinsert_restriction (m *ClaireRestriction) *ClaireAny { 
  var Result *ClaireAny
  { var c *ClaireClass = F_domain_I_restriction(m)
    { 
      var c2 *ClaireClass
      _ = c2
      var c2_iter *ClaireAny
      Result= CFALSE.Id()
      var c2_support *ClaireSet
      c2_support = c.Descendants
      for i_it := 0; i_it < c2_support.Count; i_it++ { 
        c2_iter = c2_support.At(i_it)
        c2 = ToClass(c2_iter)
        F_hashinsert_class(c2,ToMethod(m.Id()))
        } 
      } 
    } 
  return Result} 

// The EID go function for: hashinsert @ restriction (throw: false) 
func E_hashinsert_restriction (m EID) EID { 
  return F_hashinsert_restriction(ToRestriction(OBJ(m)) ).ToEID()} 

// insert into the hash table - since the order is not guaranteed when we build the dictionary, we
// need to check that m is more suited than anything that could be there
/* The go function for: hashinsert(c:class,m:method) [status=0] */
func F_hashinsert_class (c *ClaireClass,m *ClaireMethod) *ClaireAny { 
  if (c.Dictionary.Id() == CNULL) { 
    c.Dictionary = ToType(C_property.Id()).Map_I(ToType(C_method.Id()))
    } 
  { var m1 *ClaireAny = F_dict_get_any(c.Dictionary.Id(),m.Selector.Id())
    if ((m1 == CNULL) || 
        (F_domain_I_restriction(ToRestriction(m.Id())).IsIn(F_domain_I_restriction(ToRestriction(m1))) == CTRUE)) { 
      F_dict_put_any(c.Dictionary.Id(),m.Selector.Id(),m.Id())
      } 
    } 
  return  c.Dictionary.Id()
  } 

// The EID go function for: hashinsert @ class (throw: false) 
func E_hashinsert_class (c EID,m EID) EID { 
  return F_hashinsert_class(ToClass(OBJ(c)),ToMethod(OBJ(m)) ).ToEID()} 

// read the value in the directory (a method or unknown)
/* The go function for: hashget(c:class,p:property) [status=0] */
func F_hashget_class (c *ClaireClass,p *ClaireProperty) *ClaireObject { 
  return  ToObject(F_dict_get_any(c.Dictionary.Id(),p.Id()))
  } 

// The EID go function for: hashget @ class (throw: false) 
func E_hashget_class (c EID,p EID) EID { 
  return EID{F_hashget_class(ToClass(OBJ(c)),ToProperty(OBJ(p)) ).Id(),0}} 

// UGLY CAST to remove
// look if two signature have a non-empty intersection
// note that the first case with classes is necessary for bootstrapping
/* The go function for: join(x:list,y:list) [status=0] */
func F_join_list (x *ClaireList,y *ClaireList) *ClaireBoolean { 
  var Result *ClaireBoolean
  { var n int = x.Length()
    { 
      var v_and2 *ClaireBoolean
      
      v_and2 = Equal(MakeInteger(n).Id(),MakeInteger(y.Length()).Id())
      if (v_and2 == CFALSE) {Result = CFALSE
      } else { 
        { var arg_1 *ClaireAny
          { var i int = 1
            { var g0033 int = n
              arg_1= CFALSE.Id()
              for (i <= g0033) { 
                if (F_boolean_I_any(F_join_class(ToTypeExpression(x.At(i-1)).Class_I(),ToTypeExpression(y.At(i-1)).Class_I()).Id()).Id() != CTRUE.Id()) { 
                  arg_1 = CTRUE.Id()
                  break
                  } 
                i = (i+1)
                } 
              } 
            } 
          v_and2 = F_not_any(arg_1)
          } 
        if (v_and2 == CFALSE) {Result = CFALSE
        } else { 
          { var arg_2 *ClaireAny
            { var i int = 1
              { var g0034 int = n
                arg_2= CFALSE.Id()
                for (i <= g0034) { 
                  if (F_boolean_I_any(ANY(F_CALL(ToProperty(C_glb.Id()),ARGS(x.At(i-1).ToEID(),y.At(i-1).ToEID())))).Id() != CTRUE.Id()) { 
                    arg_2 = CTRUE.Id()
                    break
                    } 
                  i = (i+1)
                  } 
                } 
              } 
            v_and2 = F_not_any(arg_2)
            } 
          if (v_and2 == CFALSE) {Result = CFALSE
          } else { 
            Result = CTRUE} 
          } 
        } 
      } 
    } 
  return Result} 

// The EID go function for: join @ list (throw: false) 
func E_join_list (x EID,y EID) EID { 
  return EID{F_join_list(ToList(OBJ(x)),ToList(OBJ(y)) ).Id(),0}} 

// *********************************************************************
// *      Part 4: Matching Methods                                     *
// *********************************************************************
// Key Axiom : this code is not using dynamic calls because we use the two closed forms %type and <=t 
// which are defined in Kernel as functions (Contains and Included)
// this is the method that matches the compilation pattern 
// n is the number of args that have been pushed in the stack
/* The go function for: stack_apply(p:property,n:integer) [status=1] */
func F_CALL (p *ClaireProperty,n int) EID { 
  var Result EID
  { var i int = (ClEnv.Index-n)
    Result = F_eval_message_property(p,F_find_which_property(p,i,OWNER(ClEnv.EvalStack[i])),i,CFALSE)
    } 
  return Result} 

// The EID go function for: stack_apply @ property (throw: true) 
func E_CALL (p EID,n EID) EID { 
  return F_CALL(ToProperty(OBJ(p)),INT(n) )} 

// version where the class of first argument is forced (super)       
/* The go function for: super_apply(p:property,c:class,n:integer) [status=1] */
func F_SUPER (p *ClaireProperty,c *ClaireClass,n int) EID { 
  var Result EID
  { var top int = ClEnv.Index
    { var i int = (top-n)
      Result = F_eval_message_property(p,F_find_which_class(c,p.Definition,i,top),i,CFALSE)
      } 
    } 
  return Result} 

// The EID go function for: super_apply @ property (throw: true) 
func E_SUPER (p EID,c EID,n EID) EID { 
  return F_SUPER(ToProperty(OBJ(p)),ToClass(OBJ(c)),INT(n) )} 

// find the correct restrictions to be applied on a given set
// This is also optimized because it is very useful (it returns false if none is found)
/* The go function for: @(self:property,x:class) [status=0] */
func F__at_property1 (self *ClaireProperty,x *ClaireClass) *ClaireObject { 
  var Result *ClaireObject
  if (self.Dictionary == CTRUE) { 
    { var rx *ClaireObject = F_hashget_class(x,self)
      Result = ToObject(IfThenElse((rx.Id() != CNULL),
        rx.Id(),
        CFALSE.Id()))
      } 
    } else {
    { var rx *ClaireAny
      { var r_some *ClaireAny = CNULL
        { 
          var r *ClaireRestriction
          _ = r
          var r_iter *ClaireAny
          for _,r_iter = range(self.Definition.ValuesO()){ 
            r = ToRestriction(r_iter)
            if (x.IsIn(ToTypeExpression(r.Domain.ValuesO()[0]).Class_I()) == CTRUE) { 
              r_some = r.Id()
              break
              } 
            } 
          } 
        rx = r_some
        } 
      Result = ToObject(IfThenElse((rx != CNULL),
        rx,
        CFALSE.Id()))
      } 
    } 
  return Result} 

// The EID go function for: @ @ list<type_expression>(property, class) (throw: false) 
func E__at_property1 (self EID,x EID) EID { 
  return EID{F__at_property1(ToProperty(OBJ(self)),ToClass(OBJ(x)) ).Id(),0}} 

// finds a property through its full domain
/* The go function for: @(self:property,lt:list) [status=0] */
func F__at_property2 (self *ClaireProperty,lt *ClaireList) *ClaireObject { 
  var Result *ClaireObject
  { var rx *ClaireAny
    { var r_some *ClaireAny = CNULL
      { 
        var r *ClaireRestriction
        _ = r
        var r_iter *ClaireAny
        for _,r_iter = range(self.Definition.ValuesO()){ 
          r = ToRestriction(r_iter)
          if (F_tmatch_ask_list(lt,r.Domain) == CTRUE) { 
            r_some = r.Id()
            break
            } 
          } 
        } 
      rx = r_some
      } 
    Result = ToObject(IfThenElse((rx != CNULL),
      rx,
      CFALSE.Id()))
    } 
  return Result} 

// The EID go function for: @ @ list<type_expression>(property, list) (throw: false) 
func E__at_property2 (self EID,lt EID) EID { 
  return EID{F__at_property2(ToProperty(OBJ(self)),ToList(OBJ(lt)) ).Id(),0}} 

// method's pattern matching : l is non nil, hence last(l) is safe  {called in find_which}
// we match a list of args in the stack [n ... m] to the list of type_expressions l
/* The go function for: matching?(l:list,n:integer,m:integer) [status=0] */
func F_matching_ask_list (l *ClaireList,n int,m int) *ClaireBoolean { 
  var Result *ClaireBoolean
  { var x int = (m-n)
    { var z int = l.Length()
      if ((z == x) && 
          (l.ValuesO()[x-1].Id() != C_listargs.Id())) { 
        { var arg_1 *ClaireAny
          { var i int = 1
            { var g0035 int = x
              arg_1= CFALSE.Id()
              for (i <= g0035) { 
                { var y int = ((n-1)+i)
                  { var u *ClaireAny = l.ValuesO()[i-1].Id()
                    var g0037I *ClaireBoolean
                    if (u.Isa.Id() == C_class.Id()) { 
                      g0037I = OWNER(ClEnv.EvalStack[y]).IsIn(ToClass(u)).Not
                      } else {
                      g0037I = F_vmatch_ask_any(u,ANY(ClEnv.EvalStack[y]),n).Not
                      } 
                    if (g0037I == CTRUE) { 
                      arg_1 = CTRUE.Id()
                      break
                      } 
                    } 
                  } 
                i = (i+1)
                } 
              } 
            } 
          Result = F_not_any(arg_1)
          } 
        }  else if ((ANY(F_last_list(l)) == C_listargs.Id()) && 
          (x >= (z-1))) { 
        { var arg_2 *ClaireAny
          { var i int = 1
            { var g0036 int = z
              arg_2= CFALSE.Id()
              for (i <= g0036) { 
                { var y int = ((n-1)+i)
                  if (l.At(i-1) == C_listargs.Id()) { 
                    ClEnv.EvalStack[y]=EID{F_get_args_integer(y).Id(),0}
                    ClEnv.Index= (y+1)
                    arg_2 = CFALSE.Id()
                    break
                    }  else if (F_vmatch_ask_any(l.At(i-1),ANY(ClEnv.EvalStack[y]),n) != CTRUE) { 
                    arg_2 = CTRUE.Id()
                    break
                    } 
                  } 
                i = (i+1)
                } 
              } 
            } 
          Result = F_not_any(arg_2)
          } 
        } else {
        Result = CFALSE
        } 
      } 
    } 
  return Result} 

// The EID go function for: matching? @ list (throw: false) 
func E_matching_ask_list (l EID,n EID,m EID) EID { 
  return EID{F_matching_ask_list(ToList(OBJ(l)),INT(n),INT(m) ).Id(),0}} 

// type's pattern matching - almost like % but accepts patterns such as Reference (extended in Optimizer)
// this is why we pass n (index in stack) as an argument
// t is the type expression and x is the value
/* The go function for: vmatch?(t:any,x:any,n:integer) [status=0] */
func F_vmatch_ask_any (t *ClaireAny,x *ClaireAny,n int) *ClaireBoolean { 
  var Result *ClaireBoolean
  if (C_class.Id() == t.Isa.Id()) { 
    { var g0038 *ClaireClass = ToClass(t)
      Result = x.Isa.IsIn(g0038)
      } 
    }  else if (C_set.Id() == t.Isa.Id()) { 
    { var g0039 *ClaireSet = ToSet(t)
      Result = g0039.Contain_ask(x)
      } 
    }  else if (t.Isa.IsIn(C_subtype) == CTRUE) { 
    { var g0040 *ClaireSubtype = ToSubtype(t)
      { 
        var v_and3 *ClaireBoolean
        
        if (g0040.Arg.Id() == C_subtype.Id()) { 
          v_and3 = x.Isa.IsIn(C_type)
          } else {
          v_and3 = F__Z_any1(x,g0040.Arg)
          } 
        if (v_and3 == CFALSE) {Result = CFALSE
        } else { 
          v_and3 = ToType(x).Included(g0040.T1)
          if (v_and3 == CFALSE) {Result = CFALSE
          } else { 
            Result = CTRUE} 
          } 
        } 
      } 
    }  else if (t.Isa.IsIn(C_Param) == CTRUE) { 
    { var g0041 *ClaireParam = To_Param(t)
      { 
        var v_and3 *ClaireBoolean
        
        v_and3 = F_vmatch_ask_any(g0041.Arg.Id(),x,n)
        if (v_and3 == CFALSE) {Result = CFALSE
        } else { 
          { var arg_1 *ClaireAny
            { var i int = 1
              { var g0042 int = g0041.Params.Length()
                arg_1= CFALSE.Id()
                for (i <= g0042) { 
                  var g0049I *ClaireBoolean
                  { var arg_2 *ClaireBoolean
                    { var _Zt *ClaireAny = g0041.Args.At(i-1)
                      { var _Zv *ClaireAny = ANY(F_funcall_property(ToProperty(g0041.Params.At(i-1)),x))
                        if ((C_set.Id() == _Zt.Isa.Id()) && 
                            (_Zv.Isa.IsIn(C_type) == CTRUE)) { 
                          { var arg_3 *ClaireAny
                            { 
                              var z *ClaireAny
                              _ = z
                              arg_3= CFALSE.Id()
                              var z_support *ClaireSet
                              z_support = ToSet(_Zt)
                              for i_it := 0; i_it < z_support.Count; i_it++ { 
                                z = z_support.At(i_it)
                                if (F__equaltype_ask_any(ToType(_Zv),ToType(z)) == CTRUE) { 
                                  arg_3 = CTRUE.Id()
                                  break
                                  } 
                                } 
                              } 
                            arg_2 = F_boolean_I_any(arg_3)
                            } 
                          } else {
                          arg_2 = F_vmatch_ask_any(_Zt,_Zv,n)
                          } 
                        } 
                      } 
                    g0049I = arg_2.Not
                    } 
                  if (g0049I == CTRUE) { 
                    arg_1 = CTRUE.Id()
                    break
                    } 
                  i = (i+1)
                  } 
                } 
              } 
            v_and3 = F_not_any(arg_1)
            } 
          if (v_and3 == CFALSE) {Result = CFALSE
          } else { 
            Result = CTRUE} 
          } 
        } 
      } 
    }  else if (t.Isa.IsIn(C_Reference) == CTRUE) { 
    { var g0043 *ClaireReference = To_Reference(t)
      { var v *ClaireAny = F_get_Reference(g0043,ANY(ClEnv.EvalStack[(n+g0043.Index)]))
        if (g0043.Arg == CTRUE) { 
          Result = Equal(x,v)
          } else {
          Result = ToType(v).Contains(x)
          } 
        } 
      } 
    }  else if (C_tuple.Id() == t.Isa.Id()) { 
    { var g0044 *ClaireTuple = ToTuple(t)
      if (C_tuple.Id() == x.Isa.Id()) { 
        { var g0045 *ClaireTuple = ToTuple(x)
          { 
            var v_and5 *ClaireBoolean
            
            v_and5 = Equal(MakeInteger(g0044.Length()).Id(),MakeInteger(g0045.Length()).Id())
            if (v_and5 == CFALSE) {Result = CFALSE
            } else { 
              { var arg_4 *ClaireAny
                { var i int = 1
                  { var g0046 int = g0045.Length()
                    arg_4= CFALSE.Id()
                    for (i <= g0046) { 
                      if (F_vmatch_ask_any(ToList(g0044.Id()).At(i-1),ToList(g0045.Id()).At(i-1),n) != CTRUE) { 
                        arg_4 = CTRUE.Id()
                        break
                        } 
                      i = (i+1)
                      } 
                    } 
                  } 
                v_and5 = F_not_any(arg_4)
                } 
              if (v_and5 == CFALSE) {Result = CFALSE
              } else { 
                Result = CTRUE} 
              } 
            } 
          } 
        } else {
        Result = CFALSE
        } 
      } 
    } else {
    Result = ToBoolean(ANY(F_BELONG(x,t)))
    } 
  return Result} 

// The EID go function for: vmatch? @ any (throw: false) 
func E_vmatch_ask_any (t EID,x EID,n EID) EID { 
  return EID{F_vmatch_ask_any(ANY(t),ANY(x),INT(n) ).Id(),0}} 

// extensibility for type_expressions
// method's pattern matching based on type expressions (i.e. l2 is another list of type expressions).
// this is an extension of <=t to   all type expressions
/* The go function for: tmatch?(l:list,l2:list) [status=0] */
func F_tmatch_ask_list (l *ClaireList,l2 *ClaireList) *ClaireBoolean { 
  var Result *ClaireBoolean
  { var x int = l2.Length()
    { var z int = l.Length()
      if ((z != x) && 
          ((l2.At(x-1) != C_listargs.Id()) || 
              (z < (x-1)))) { 
        Result = CFALSE
        } else {
        { var arg_1 *ClaireAny
          { var i int = 1
            { var g0050 int = x
              arg_1= CFALSE.Id()
              for (i <= g0050) { 
                if ((i == x) && 
                    (l2.At(i-1) == C_listargs.Id())) { 
                  arg_1 = CFALSE.Id()
                  break
                  }  else if (F_tmatch_ask_any(l.At(i-1),l2.At(i-1),l) != CTRUE) { 
                  arg_1 = CTRUE.Id()
                  break
                  } 
                i = (i+1)
                } 
              } 
            } 
          Result = F_not_any(arg_1)
          } 
        } 
      } 
    } 
  return Result} 

// The EID go function for: tmatch? @ list (throw: false) 
func E_tmatch_ask_list (l EID,l2 EID) EID { 
  return EID{F_tmatch_ask_list(ToList(OBJ(l)),ToList(OBJ(l2)) ).Id(),0}} 

// type_expression pattern matching (t is the variable and t2 the pattern)
// this is an extension of <=t for the pattern Reference
/* The go function for: tmatch?(t:any,mClaire/t2:any,l:list) [status=0] */
func F_tmatch_ask_any (t *ClaireAny,t2 *ClaireAny,l *ClaireList) *ClaireBoolean { 
  var Result *ClaireBoolean
  if (t2.Isa.IsIn(C_Reference) == CTRUE) { 
    { var g0051 *ClaireReference = To_Reference(t2)
      if (g0051.Arg == CTRUE) { 
        Result = CFALSE
        }  else if (t.Isa.IsIn(C_Reference) == CTRUE) { 
        { var g0052 *ClaireReference = To_Reference(t)
          Result = MakeBoolean((g0052.Index == g0051.Index) && (Equal(g0052.Args.Id(),g0051.Args.Id()) == CTRUE))
          } 
        }  else if (t.Isa.IsIn(C_type) == CTRUE) { 
        { var g0053 *ClaireType = ToType(t)
          { var tref *ClaireType = F_member_type(ToType(F__at_Reference(g0051,g0051.Args,l.At((g0051.Index+1)-1))))
            
            Result = g0053.Included(tref)
            } 
          } 
        } else {
        Result = CFALSE
        } 
      } 
    }  else if (t2.Isa.IsIn(C_type) == CTRUE) { 
    { var g0054 *ClaireType = ToType(t2)
      if (t.Isa.IsIn(C_type) == CTRUE) { 
        { var g0055 *ClaireType = ToType(t)
          Result = g0055.Included(g0054)
          } 
        } else {
        Result = ToBoolean(OBJ(F_CALL(ToProperty(C_less_ask.Id()),ARGS(t.ToEID(),EID{g0054.Id(),0}))))
        } 
      } 
    } else {
    Result = ToBoolean(OBJ(F_CALL(ToProperty(C_less_ask.Id()),ARGS(t.ToEID(),t2.ToEID()))))
    } 
  return Result} 

// The EID go function for: tmatch? @ any (throw: false) 
func E_tmatch_ask_any (t EID,t2 EID,l EID) EID { 
  return EID{F_tmatch_ask_any(ANY(t),ANY(t2),ToList(OBJ(l)) ).Id(),0}} 

// find the restriction (n is the position of the arglist start)
/* The go function for: find_which(p:property,n:integer,c:class) [status=0] */
func F_find_which_property (p *ClaireProperty,n int,c *ClaireClass) *ClaireObject { 
  var Result *ClaireObject
  if (p.Dictionary == CTRUE) { 
    Result = F_hashget_class(c,p)
    } else {
    { 
      var r *ClaireRestriction
      _ = r
      var r_iter *ClaireAny
      Result= ToObject(CFALSE.Id())
      for _,r_iter = range(p.Definition.ValuesO()){ 
        r = ToRestriction(r_iter)
        if (F_matching_ask_list(r.Domain,n,ClEnv.Index) == CTRUE) { 
          Result = ToObject(r.Id())
          break
          } 
        } 
      } 
    } 
  return Result} 

// The EID go function for: find_which @ property (throw: false) 
func E_find_which_property (p EID,n EID,c EID) EID { 
  return EID{F_find_which_property(ToProperty(OBJ(p)),INT(n),ToClass(OBJ(c)) ).Id(),0}} 

// used by inspect.cl
/* The go function for: find_which(l:list,c:class,n:integer,m:integer) [status=0] */
func F_find_which_list (l *ClaireList,c *ClaireClass,n int,m int) *ClaireObject { 
  var Result *ClaireObject
  { 
    var r *ClaireRestriction
    _ = r
    var r_iter *ClaireAny
    Result= ToObject(CFALSE.Id())
    var r_support *ClaireList
    r_support = l
    r_len := r_support.Length()
    for i_it := 0; i_it < r_len; i_it++ { 
      r_iter = r_support.At(i_it)
      r = ToRestriction(r_iter)
      if (F_matching_ask_list(r.Domain,n,m) == CTRUE) { 
        Result = ToObject(r.Id())
        break
        } 
      } 
    } 
  return Result} 

// The EID go function for: find_which @ list (throw: false) 
func E_find_which_list (l EID,c EID,n EID,m EID) EID { 
  return EID{F_find_which_list(ToList(OBJ(l)),
    ToClass(OBJ(c)),
    INT(n),
    INT(m) ).Id(),0}} 

// special version for super, where we give (n,m) -> position of arglist in the stack
/* The go function for: find_which(c:class,l:list,n:integer,m:integer) [status=0] */
func F_find_which_class (c *ClaireClass,l *ClaireList,n int,m int) *ClaireObject { 
  var Result *ClaireObject
  { 
    var r *ClaireRestriction
    _ = r
    var r_iter *ClaireAny
    Result= ToObject(CFALSE.Id())
    var r_support *ClaireList
    r_support = l
    r_len := r_support.Length()
    for i_it := 0; i_it < r_len; i_it++ { 
      r_iter = r_support.At(i_it)
      r = ToRestriction(r_iter)
      if ((ToType(c.Id()).Included(ToType(r.Domain.ValuesO()[0])) == CTRUE) && 
          (F_matching_ask_list(r.Domain,n,m) == CTRUE)) { 
        Result = ToObject(r.Id())
        break
        } 
      } 
    } 
  return Result} 

// The EID go function for: find_which @ class (throw: false) 
func E_find_which_class (c EID,l EID,n EID,m EID) EID { 
  return EID{F_find_which_class(ToClass(OBJ(c)),
    ToList(OBJ(l)),
    INT(n),
    INT(m) ).Id(),0}} 
