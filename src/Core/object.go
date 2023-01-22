/***** CLAIRE Compilation of file /Users/ycaseau/Dropbox/src/clairev4.07/src/meta/object.cl 
         [version 4.0.07 / safety 5] Sunday 12-25-2022 13:07:24 *****/

package Core
import (_ "fmt"
	. "Kernel"
)

//+-------------------------------------------------------------+
//| CLAIRE                                                      |
//| object.cl                                                   |
//| Copyright (C) 1994 - 2021 Yves Caseau. All Rights Reserved  |
//| cf. copyright info in about()                               |
//+-------------------------------------------------------------+
// ---------------------------------------------------------------------
// This file contains the definition of the objects that implement the
// core features of the microCLAIRE library: traceable & debug-able calls,
// tables, demons and exceptions
// ---------------------------------------------------------------------
// *********************************************************************
// *  Table of contents                                                *
// *   Part 1: Ask, debug & trace                                      *
// *   Part 2: Tables                                                  *
// *   Part 3: Demons & relations for the logic modules                *
// *   Part 4: Basics of Exceptions                                    *
// *********************************************************************
// release() should produce a version number
/* The go function for: release(_CL_obj:void) [status=1] */
func F_release_void () EID { 
    var Result EID
    { var arg_1 *ClaireString
      var try_2 EID
      try_2 = F_string_I_float(ClEnv.Version)
      if ErrorIn(try_2) {Result = try_2
      } else {
      arg_1 = ToString(OBJ(try_2))
      Result = EID{F_append_string(MakeString("4."),arg_1).Id(),0}
      }
      } 
    return Result} 
  
// The EID go function for: release @ void (throw: true) 
func E_release_void (_CL_obj EID) EID { 
    return F_release_void( )} 
  
// the about method produces the legal warning, according to the GNU software
// recommendation
/* The go function for: about(_CL_obj:void) [status=0] */
func F_about_void () *ClaireAny { 
    PRINC("CLAIRE v4.")
    F_princ_float(ClEnv.Version)
    PRINC(" Copyright (C) 1994-2021 Yves Caseau. All Rights Reserved.\n")
    PRINC("use and redistribution in source code or binary forms are permitted\n")
    PRINC("resale is not permitted without the explicit agreement of Yves Caseau\n")
    PRINC("THIS SOFTWARE IS PROVIDED AS IS AND WITHOUT ANY WARRANTY, INCLUDING,\n")
    PRINC("WITHOUT LIMITATION, THE IMPLIED WARRANTIES OF MERCHANTABILTY AND FITNESS\n")
    PRINC("FOR A PARTICULAR PURPOSE\n")
    return  CTRUE.Id()
    } 
  
// The EID go function for: about @ void (throw: false) 
func E_about_void (_CL_obj EID) EID { 
    return F_about_void( ).ToEID()} 
  
// properties that are defined through compiling (eval would entail a loop)
// added for upward compatibility reasons
// *********************************************************************
// *   Part 1: Ask, debug & trace                                      *
// *********************************************************************
// create the list of arguments if needed : allocate on the stack
/* The go function for: mClaire/get_args(i:integer) [status=0] */
func F_get_args_integer (i int) *ClaireList { 
    var Result *ClaireList
    { var liste *ClaireList = ToType(C_any.Id()).EmptyList()
      for (i < ClEnv.Index) { 
        liste = liste.AddFast(ANY(ClEnv.EvalStack[i]))
        i = (i+1)
        } 
      Result = liste
      } 
    return Result} 
  
// The EID go function for: mClaire/get_args @ integer (throw: false) 
func E_get_args_integer (i EID) EID { 
    return EID{F_get_args_integer(INT(i) ).Id(),0}} 
  
// a simple method for a direct call with no argument
/* The go function for: funcall(self:method,x:any) [status=1] */
func F_funcall_method1 (self *ClaireMethod,x *ClaireAny) EID { 
    var Result EID
    { var start int = ClEnv.Index
      ClEnv.Push(x.ToEID())
      Result = F_execute_method(self,start,CFALSE)
      } 
    return Result} 
  
// The EID go function for: funcall @ list<type_expression>(method, any) (throw: true) 
func E_funcall_method1 (self EID,x EID) EID { 
    return F_funcall_method1(ToMethod(OBJ(self)),ANY(x) )} 
  
// this is a simple method for calling directly a method with one argument
/* The go function for: funcall(self:method,x:any,y:any) [status=1] */
func F_funcall_method2 (self *ClaireMethod,x *ClaireAny,y *ClaireAny) EID { 
    var Result EID
    { var start int = ClEnv.Index
      ClEnv.Push(x.ToEID())
      ClEnv.Push(y.ToEID())
      Result = F_execute_method(self,start,CFALSE)
      } 
    return Result} 
  
// The EID go function for: funcall @ list<type_expression>(method, any, any) (throw: true) 
func E_funcall_method2 (self EID,x EID,y EID) EID { 
    return F_funcall_method2(ToMethod(OBJ(self)),ANY(x),ANY(y) )} 
  
// this is a simple method for calling directly a method with two arguments
/* The go function for: funcall(self:method,x:any,y:any,z:any) [status=1] */
func F_funcall_method3 (self *ClaireMethod,x *ClaireAny,y *ClaireAny,z *ClaireAny) EID { 
    var Result EID
    { var start int = ClEnv.Index
      ClEnv.Push(x.ToEID())
      ClEnv.Push(y.ToEID())
      ClEnv.Push(z.ToEID())
      Result = F_execute_method(self,start,CFALSE)
      } 
    return Result} 
  
// The EID go function for: funcall @ list<type_expression>(method, any, any, any) (throw: true) 
func E_funcall_method3 (self EID,x EID,y EID,z EID) EID { 
    return F_funcall_method3(ToMethod(OBJ(self)),
      ANY(x),
      ANY(y),
      ANY(z) )} 
  
// how to apply a property to a list  (the function case is handled in Kernel - primitive go code)
/* The go function for: call(p:property,l:listargs) [status=1] */
func F_call_property (p *ClaireProperty,l *ClaireList) EID { 
    var Result EID
    Result = F_apply_property(p,ToList(l.Id()))
    return Result} 
  
// The EID go function for: call @ property (throw: true) 
func E_call_property (p EID,l EID) EID { 
    return F_call_property(ToProperty(OBJ(p)),ToList(OBJ(l)) )} 
  
/* The go function for: apply(p:property,l:list) [status=1] */
func F_apply_property (p *ClaireProperty,l *ClaireList) EID { 
    var Result EID
    { var start int = ClEnv.Index
      { 
        var x *ClaireAny
        _ = x
        var x_support *ClaireList
        x_support = l
        x_len := x_support.Length()
        for i_it := 0; i_it < x_len; i_it++ { 
          x = x_support.At(i_it)
          ClEnv.Push(x.ToEID())
          } 
        } 
      Result = F_eval_message_property(p,F_find_which_property(p,start,l.At(0).Isa),start,CTRUE)
      } 
    return Result} 
  
// The EID go function for: apply @ property (throw: true) 
func E_apply_property (p EID,l EID) EID { 
    return F_apply_property(ToProperty(OBJ(p)),ToList(OBJ(l)) )} 
  
/* The go function for: apply(m:method,l:list) [status=1] */
func F_apply_method (m *ClaireMethod,l *ClaireList) EID { 
    var Result EID
    { var start int = ClEnv.Index
      { 
        var x *ClaireAny
        _ = x
        var x_support *ClaireList
        x_support = l
        x_len := x_support.Length()
        for i_it := 0; i_it < x_len; i_it++ { 
          x = x_support.At(i_it)
          ClEnv.Push(x.ToEID())
          } 
        } 
      Result = F_execute_method(m,start,CFALSE)
      } 
    return Result} 
  
// The EID go function for: apply @ method (throw: true) 
func E_apply_method (m EID,l EID) EID { 
    return F_apply_method(ToMethod(OBJ(m)),ToList(OBJ(l)) )} 
  
// defined in Reader but tested in Core
// push and pop debug info on the stack
// this method also does the tracing and the steppping
/* The go function for: push_debug(prop:property,arity:integer,start:integer) [status=1] */
func F_push_debug_property (prop *ClaireProperty,arity int,start int) EID { 
    var Result EID
    { var i int = ClEnv.Index
      { var n int = ClEnv.Trace_I
        if ((n > 0) && 
            ((prop.Trace_I+ClEnv.Verbose) > 4)) { 
          { var p *ClairePort = ClEnv.Ctrace.UseAsOutput()
            ClEnv.Trace_I = 0
            F_tr_indent_boolean(CFALSE,n)
            PRINC(" ")
            Result = F_print_any(prop.Id())
            if !ErrorIn(Result) {
            PRINC("(")
            Result = F_CALL(C_print,ARGS(ClEnv.EvalStack[start]))
            }
            if !ErrorIn(Result) {
            { var j int = (start+1)
              Result= EID{CFALSE.Id(),0}
              for (j < (start+arity)) { 
                var loop_1 EID
                _ = loop_1
                { 
                PRINC(",")
                loop_1 = F_CALL(C_print,ARGS(ClEnv.EvalStack[j]))
                if ErrorIn(loop_1) {Result = loop_1
                break
                } else {
                }
                if ErrorIn(loop_1) {Result = loop_1
                break
                } else {
                j = (j+1)
                }
                } 
              }
              } 
            if !ErrorIn(Result) {
            if (ClEnv.CountCall >= 0) { 
              ClEnv.CountCall = (ClEnv.CountCall+1)
              PRINC(" [")
              F_princ_integer(ClEnv.CountCall)
              PRINC("]")
              if (ClEnv.CountCall == ClEnv.CountLevel) { 
                if (ClEnv.CountTrigger == C_spy.Id()) { 
                  Result = F_update_property(C_spy_I,
                    ToObject(ClEnv.Id()),
                    17,
                    C_object,
                    F__at_property1(C_spy,C_void).Id())
                  } else {
                  { 
                    var va_arg1 *ClaireEnvironment
                    var va_arg2 int
                    va_arg1 = ClEnv
                    va_arg2 = ToInteger(ClEnv.CountTrigger).Value
                    va_arg1.Verbose = va_arg2
                    Result = EID{C__INT,IVAL(va_arg2)}
                    } 
                  } 
                } else {
                Result = EID{CFALSE.Id(),0}
                } 
              } else {
              Result = EID{CFALSE.Id(),0}
              } 
            if !ErrorIn(Result) {
            PRINC(")\n")
            ClEnv.Trace_I = (n+1)
            Result = p.UseAsOutput().ToEID()
            }}}
            } 
          } else {
          Result = EID{CFALSE.Id(),0}
          } 
        if !ErrorIn(Result) {
        if (F_get_table(C_Core_StopProperty,prop.Id()) != CNULL) { 
          var g0060I *ClaireBoolean
          var try_2 EID
          { 
            var v_or5 *ClaireBoolean
            
            v_or5 = Equal(F_get_table(C_Core_StopProperty,prop.Id()),CNIL.Id())
            if (v_or5 == CTRUE) {try_2 = EID{CTRUE.Id(),0}
            } else { 
              var try_3 EID
              { var arg_4 *ClaireAny
                var try_5 EID
                { 
                  var l2 *ClaireAny
                  _ = l2
                  try_5= EID{CFALSE.Id(),0}
                  var l2_support *ClaireList
                  l2_support = ToList(F_get_table(C_Core_StopProperty,prop.Id()))
                  l2_len := l2_support.Length()
                  for i_it := 0; i_it < l2_len; i_it++ { 
                    l2 = l2_support.At(i_it)
                    var loop_6 EID
                    _ = loop_6
                    var g0061I *ClaireBoolean
                    var try_7 EID
                    { var arg_8 *ClaireAny
                      var try_9 EID
                      { var j int = 1
                        { var g0059 int = INT(F_CALL(C_length,ARGS(l2.ToEID())))
                          try_9= EID{CFALSE.Id(),0}
                          for (j <= g0059) { 
                            var loop_10 EID
                            _ = loop_10
                            { 
                            var g0062I *ClaireBoolean
                            var try_11 EID
                            { var arg_12 *ClaireBoolean
                              var try_13 EID
                              { 
                                var v_and15 *ClaireBoolean
                                
                                v_and15 = F__inf_equal_integer((j+start),i)
                                if (v_and15 == CFALSE) {try_13 = EID{CFALSE.Id(),0}
                                } else { 
                                  var try_14 EID
                                  { var arg_15 *ClaireAny
                                    var try_16 EID
                                    try_16 = F_CALL(C_nth,ARGS(l2.ToEID(),EID{C__INT,IVAL(j)}))
                                    if ErrorIn(try_16) {try_14 = try_16
                                    } else {
                                    arg_15 = ANY(try_16)
                                    try_14 = EID{Equal(arg_15,ANY(ClEnv.EvalStack[((start+j)-1)])).Id(),0}
                                    }
                                    } 
                                  if ErrorIn(try_14) {try_13 = try_14
                                  } else {
                                  v_and15 = ToBoolean(OBJ(try_14))
                                  if (v_and15 == CFALSE) {try_13 = EID{CFALSE.Id(),0}
                                  } else { 
                                    try_13 = EID{CTRUE.Id(),0}} 
                                  } 
                                }
                                } 
                              if ErrorIn(try_13) {try_11 = try_13
                              } else {
                              arg_12 = ToBoolean(OBJ(try_13))
                              try_11 = EID{arg_12.Not.Id(),0}
                              }
                              } 
                            if ErrorIn(try_11) {loop_10 = try_11
                            } else {
                            g0062I = ToBoolean(OBJ(try_11))
                            if (g0062I == CTRUE) { 
                              try_9 = EID{CTRUE.Id(),0}
                              break
                              } else {
                              loop_10 = EID{CFALSE.Id(),0}
                              } 
                            }
                            if ErrorIn(loop_10) {try_9 = loop_10
                            break
                            } else {
                            j = (j+1)
                            }
                            } 
                          }
                          } 
                        } 
                      if ErrorIn(try_9) {try_7 = try_9
                      } else {
                      arg_8 = ANY(try_9)
                      try_7 = EID{F_not_any(arg_8).Id(),0}
                      }
                      } 
                    if ErrorIn(try_7) {loop_6 = try_7
                    } else {
                    g0061I = ToBoolean(OBJ(try_7))
                    if (g0061I == CTRUE) { 
                      try_5 = EID{CTRUE.Id(),0}
                      break
                      } else {
                      loop_6 = EID{CFALSE.Id(),0}
                      } 
                    }
                    if ErrorIn(loop_6) {try_5 = loop_6
                    break
                    } else {
                    }
                    } 
                  } 
                if ErrorIn(try_5) {try_3 = try_5
                } else {
                arg_4 = ANY(try_5)
                try_3 = EID{F_boolean_I_any(arg_4).Id(),0}
                }
                } 
              if ErrorIn(try_3) {try_2 = try_3
              } else {
              v_or5 = ToBoolean(OBJ(try_3))
              if (v_or5 == CTRUE) {try_2 = EID{CTRUE.Id(),0}
              } else { 
                try_2 = EID{CFALSE.Id(),0}} 
              } 
            }
            } 
          if ErrorIn(try_2) {Result = try_2
          } else {
          g0060I = ToBoolean(OBJ(try_2))
          if (g0060I == CTRUE) { 
            Result = ToException(C_general_error.Make(MakeString("stop as required in ~S(~A)").Id(),MakeConstantList(prop.Id(),F_get_args_integer(start).Id()).Id())).Close()
            } else {
            Result = EID{CFALSE.Id(),0}
            } 
          }
          } else {
          Result = EID{CFALSE.Id(),0}
          } 
        if !ErrorIn(Result) {
        ClEnv.Push(EID{C__INT,IVAL(ClEnv.Debug_I)})
        ClEnv.Push(EID{prop.Id(),0})
        ClEnv.Push(EID{C__INT,IVAL(arity)})
        ClEnv.Push(EID{C__INT,IVAL(start)})
        { 
          var va_arg1 *ClaireEnvironment
          var va_arg2 int
          va_arg1 = ClEnv
          va_arg2 = i
          va_arg1.Debug_I = va_arg2
          Result = EID{C__INT,IVAL(va_arg2)}
          } 
        }}
        } 
      } 
    return Result} 
  
// The EID go function for: push_debug @ property (throw: true) 
func E_push_debug_property (prop EID,arity EID,start EID) EID { 
    return F_push_debug_property(ToProperty(OBJ(prop)),INT(arity),INT(start) )} 
  
// value of the previous debug
// n is 0 for interpreted code and 1 for compiled code
/* The go function for: pop_debug(self:property,n:integer,val:any) [status=1] */
func F_pop_debug_property (self *ClaireProperty,n int,val *ClaireAny) EID { 
    var Result EID
    { var v int = ClEnv.Debug_I
      if (v > 0) { 
        if (n != 0) { 
          ClEnv.Index= INT(ClEnv.EvalStack[(v+3)])
          } 
        ClEnv.Debug_I = INT(ClEnv.EvalStack[ClEnv.Debug_I])
        if (self.IfWrite == CNULL) { 
          { var m *ClaireObject = ClEnv.Spy_I
            if (m.Id() != CNULL) { 
              ClEnv.Spy_I = ToObject(CNULL)
              Result = F_funcall_method1(ToMethod(m.Id()),ClEnv.Id())
              if !ErrorIn(Result) {
              { 
                var va_arg1 *ClaireEnvironment
                var va_arg2 *ClaireObject
                va_arg1 = ClEnv
                va_arg2 = m
                va_arg1.Spy_I = va_arg2
                Result = EID{va_arg2.Id(),0}
                } 
              }
              } else {
              Result = EID{CFALSE.Id(),0}
              } 
            } 
          } else {
          Result = EID{CFALSE.Id(),0}
          } 
        if !ErrorIn(Result) {
        if ((ClEnv.Trace_I > 1) && 
            ((self.Trace_I+ClEnv.Verbose) > 4)) { 
          { var i int = ClEnv.Trace_I
            ClEnv.Trace_I = 0
            if ((self.Trace_I+ClEnv.Verbose) > 4) { 
              { var p *ClairePort = ClEnv.Ctrace.UseAsOutput()
                F_tr_indent_boolean(CTRUE,(i-1))
                PRINC(" ")
                Result = F_CALL(C_print,ARGS(val.ToEID()))
                if !ErrorIn(Result) {
                PRINC("\n")
                Result = EVOID
                }
                if !ErrorIn(Result) {
                Result = p.UseAsOutput().ToEID()
                }
                } 
              } else {
              Result = EID{CFALSE.Id(),0}
              } 
            if !ErrorIn(Result) {
            { 
              var va_arg1 *ClaireEnvironment
              var va_arg2 int
              va_arg1 = ClEnv
              va_arg2 = (i-1)
              va_arg1.Trace_I = va_arg2
              Result = EID{C__INT,IVAL(va_arg2)}
              } 
            }
            } 
          } else {
          Result = EID{CFALSE.Id(),0}
          } 
        }
        } else {
        Result = EID{CFALSE.Id(),0}
        } 
      } 
    return Result} 
  
// The EID go function for: pop_debug @ property (throw: true) 
func E_pop_debug_property (self EID,n EID,val EID) EID { 
    return F_pop_debug_property(ToProperty(OBJ(self)),INT(n),ANY(val) )} 
  
// print a nice indented mark
/* The go function for: tr_indent(return?:boolean,n:integer) [status=0] */
func F_tr_indent_boolean (return_ask *ClaireBoolean,n int)  { 
    if (return_ask == CTRUE) { 
      PRINC("[")
      F_princ_integer(n)
      PRINC("]")
      } else {
      F_princ_integer(n)
      PRINC(":=")
      } 
    for (n > 9) { 
      PRINC("=")
      n = (n-10)
      } 
    for (n > 0) { 
      PRINC(">")
      n = (n-1)
      } 
    } 
  
// The EID go function for: tr_indent @ boolean (throw: false) 
func E_tr_indent_boolean (return_ask EID,n EID) EID { 
    F_tr_indent_boolean(ToBoolean(OBJ(return_ask)),INT(n) )
    return EVOID} 
  
// CLAIRE4 : because macros do not exist in go
// #define DB_BIND(m,p,n,l) if (m.it->status == 4) {l push_debug_property(p,n,ClEnv->index - n);}
/* The go function for: db_bind(m:module,p:property,n:integer) [status=1] */
func F_Core_db_bind_module (m *ClaireModule,p *ClaireProperty,n int) EID { 
    var Result EID
    if (m.Status == 4) { 
      Result = F_push_debug_property(p,n,(ClEnv.Index-n))
      } else {
      Result = EID{CFALSE.Id(),0}
      } 
    return Result} 
  
// The EID go function for: db_bind @ module (throw: true) 
func E_Core_db_bind_module (m EID,p EID,n EID) EID { 
    return F_Core_db_bind_module(ToModule(OBJ(m)),ToProperty(OBJ(p)),INT(n) )} 
  
// #define DB_UNBIND(m,p,v) if (m.it->status == 4) pop_debug_property(p,1,v)
/* The go function for: db_unbind(m:module,p:property,v:any) [status=1] */
func F_Core_db_unbind_module (m *ClaireModule,p *ClaireProperty,v *ClaireAny) EID { 
    var Result EID
    if (m.Status == 4) { 
      Result = F_pop_debug_property(p,1,v)
      } else {
      Result = EID{CFALSE.Id(),0}
      } 
    return Result} 
  
// The EID go function for: db_unbind @ module (throw: true) 
func E_Core_db_unbind_module (m EID,p EID,v EID) EID { 
    return F_Core_db_unbind_module(ToModule(OBJ(m)),ToProperty(OBJ(p)),ANY(v) )} 
  
// *********************************************************************
// *   Part 2: Tables                                                  *
// *********************************************************************
// finds if objects are identified - unclear if there is any need for this
/* The go function for: identified?(self:class) [status=0] */
func F_identified_ask_class (self *ClaireClass) *ClaireBoolean { 
    return  MakeBoolean((self.Id() == C_integer.Id()) || (self.IsIn(C_object) == CTRUE) || (self.Id() == C_symbol.Id()) || (self.Id() == C_boolean.Id()) || (self.Id() == C_char.Id()))
    } 
  
// The EID go function for: identified? @ class (throw: false) 
func E_identified_ask_class (self EID) EID { 
    return EID{F_identified_ask_class(ToClass(OBJ(self)) ).Id(),0}} 
  
// true pointer equality in go (used to be C++) => use externC form
/* The go function for: identical?(x:any,y:any) [status=0] */
func F_identical_ask_any (x *ClaireAny,y *ClaireAny) *ClaireBoolean { 
    return  ToBoolean(IfThenElse(x == y,CTRUE.Id(),CFALSE.Id()))
    } 
  
// The EID go function for: identical? @ any (throw: false) 
func E_identical_ask_any (x EID,y EID) EID { 
    return EID{F_identical_ask_any(ANY(x),ANY(y) ).Id(),0}} 
  
// writing a single value into a slot but does NOT trigger the rules !
// equivalent to is! of LAURE
// this definition should not be placed in the method.cl file
// (it requires some inheritance conflict processing)
/* The go function for: put(self:property,x:object,y:any) [status=1] */
func F_put_property2 (self *ClaireProperty,x *ClaireObject,y *ClaireAny) EID { 
    var Result EID
    { var s *ClaireObject = F__at_property1(self,x.Id().Isa)
      if (C_slot.Id() == s.Isa.Id()) { 
        { var g0063 *ClaireSlot = ToSlot(s.Id())
          Result = F_store_object(x,
            g0063.Index,
            g0063.Srange,
            y,
            self.Store_ask).ToEID()
          } 
        } else {
        Result = ToException(C_selector_error.Make(self.Id(),MakeConstantList(x.Id()).Id())).Close()
        } 
      } 
    return Result} 
  
// The EID go function for: put @ list<type_expression>(property, object, any) (throw: true) 
func E_put_property2 (self EID,x EID,y EID) EID { 
    return F_put_property2(ToProperty(OBJ(self)),ToObject(OBJ(x)),ANY(y) )} 
  
// v3.2 : same but multi valued
/* The go function for: add_value(self:property,x:object,y:any) [status=1] */
func F_add_value_property3 (self *ClaireProperty,x *ClaireObject,y *ClaireAny) EID { 
    var Result EID
    { var s *ClaireObject = F__at_property1(self,x.Id().Isa)
      if (F_boolean_I_any(s.Id()).Id() != CTRUE.Id()) { 
        Result = ToException(C_selector_error.Make(self.Id(),MakeConstantList(x.Id()).Id())).Close()
        }  else if (self.Multivalued_ask != CTRUE) { 
        Result = ToException(C_general_error.Make(MakeString("[134] Cannot apply add to ~S").Id(),MakeConstantList(self.Id()).Id())).Close()
        } else {
        { var n int = ToSlot(s.Id()).Index
          { var l1 *ClaireSet = ToSet(x.SlotGet(n,C_object))
            Result = EID{F_Core_add_value_I_property(self,
              x,
              n,
              l1,
              y).Id(),0}
            } 
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: add_value @ property (throw: true) 
func E_add_value_property3 (self EID,x EID,y EID) EID { 
    return F_add_value_property3(ToProperty(OBJ(self)),ToObject(OBJ(x)),ANY(y) )} 
  
// Claire4: add_value! is the internal form
// a table is implemented through its graph which is a list or a dictionary (a.params says which)
// graph_get(a:table,x:any) : any -> reads in a.graph
// graph_put(a:table,x:any,y:any) : void -> write in a.graph    
// access
// in Claire4 there is always a default hence the unknown check has disapeared
/* The go function for: nth(a:table,x:any) [status=1] */
func F_nth_table1 (a *ClaireTable,x *ClaireAny) EID { 
    var Result EID
    { var p *ClaireAny = a.Params
      if (a.Domain.Contains(x) != CTRUE) { 
        Result = ToException(C_general_error.Make(MakeString("[135] ~S does not belong to the domain of ~S").Id(),MakeConstantList(x,a.Id()).Id())).Close()
        } else {
        Result = EID{CFALSE.Id(),0}
        } 
      if !ErrorIn(Result) {
      if (C_integer.Id() == p.Isa.Id()) { 
        { var g0065 int = ToInteger(p).Value
          Result = ToList(a.Graph).At((ToInteger(x).Value-g0065)-1).ToEID()
          } 
        }  else if (p.Isa.IsIn(C_list) == CTRUE) { 
        Result = ToList(a.Graph).At(F_get_index_table2(a,ToInteger(ToList(x).At(0)).Value,ToInteger(ToList(x).At(1)).Value)-1).ToEID()
        } else {
        Result = a.GraphGet(x).ToEID()
        } 
      }
      } 
    return Result} 
  
// The EID go function for: nth @ list<type_expression>(table, any) (throw: true) 
func E_nth_table1 (a EID,x EID) EID { 
    return F_nth_table1(ToTable(OBJ(a)),ANY(x) )} 
  
/* The go function for: nth_table1_type */
func F_nth_table1_type (a *ClaireType,x *ClaireType) EID { 
    var Result EID
    if (F_unique_ask_type(a) == CTRUE) { 
      { var arg_1 *ClaireAny
        var try_2 EID
        try_2 = F_the_type(a)
        if ErrorIn(try_2) {Result = try_2
        } else {
        arg_1 = ANY(try_2)
        Result = F_CALL(C_range,ARGS(arg_1.ToEID()))
        }
        } 
      } else {
      Result = EID{C_any.Id(),0}
      } 
    return Result} 
  
  
// The dual EID go function for: "nth_table1_type" 
func E_nth_table1_type (a EID,x EID) EID { 
    return F_nth_table1_type(ToType(OBJ(a)),ToType(OBJ(x)))} 
  
// get is the same, with no error            
/* The go function for: get(a:table,x:any) [status=0] */
func F_get_table (a *ClaireTable,x *ClaireAny) *ClaireAny { 
    var Result *ClaireAny
    { var p *ClaireAny = a.Params
      if (C_integer.Id() == p.Isa.Id()) { 
        { var g0068 int = ToInteger(p).Value
          Result = ToList(a.Graph).At((ToInteger(x).Value-g0068)-1)
          } 
        }  else if (p.Isa.IsIn(C_list) == CTRUE) { 
        Result = ToList(a.Graph).At(F_get_index_table2(a,ToInteger(ToList(x).At(0)).Value,ToInteger(ToList(x).At(1)).Value)-1)
        } else {
        Result = a.GraphGet(x)
        } 
      } 
    return Result} 
  
// The EID go function for: get @ table (throw: false) 
func E_get_table (a EID,x EID) EID { 
    return F_get_table(ToTable(OBJ(a)),ANY(x) ).ToEID()} 
  
/* The go function for: get_table_type */
func F_get_table_type (a *ClaireType,x *ClaireType) EID { 
    var Result EID
    if (F_unique_ask_type(a) == CTRUE) { 
      { var arg_1 *ClaireAny
        var try_2 EID
        try_2 = F_the_type(a)
        if ErrorIn(try_2) {Result = try_2
        } else {
        arg_1 = ANY(try_2)
        Result = F_CALL(C_range,ARGS(arg_1.ToEID()))
        }
        } 
      } else {
      Result = EID{C_any.Id(),0}
      } 
    return Result} 
  
  
// The dual EID go function for: "get_table_type" 
func E_get_table_type (a EID,x EID) EID { 
    return F_get_table_type(ToType(OBJ(a)),ToType(OBJ(x)))} 
  
// interface update method for a[x] := y
/* The go function for: nth=(a:table,x:any,y:any) [status=1] */
func F_nth_equal_table1 (a *ClaireTable,x *ClaireAny,y *ClaireAny) EID { 
    var Result EID
    if (a.Domain.Contains(x) != CTRUE) { 
      Result = ToException(C_general_error.Make(MakeString("[135] ~S does not belong to the domain of ~S").Id(),MakeConstantList(x,a.Id()).Id())).Close()
      } else {
      Result = EID{CFALSE.Id(),0}
      } 
    if !ErrorIn(Result) {
    if (a.Range.Contains(y) != CTRUE) { 
      Result = ToException(C_range_error.Make(a.Id(),y,a.Range.Id())).Close()
      } else {
      Result = EID{CFALSE.Id(),0}
      } 
    if !ErrorIn(Result) {
    Result = F_nth_put_table(a,x,y)
    }}
    return Result} 
  
// The EID go function for: nth= @ list<type_expression>(table, any, any) (throw: true) 
func E_nth_equal_table1 (a EID,x EID,y EID) EID { 
    return F_nth_equal_table1(ToTable(OBJ(a)),ANY(x),ANY(y) )} 
  
// internal form without checks
// equivalent of update = put + put_inverse
/* The go function for: nth_put(a:table,x:any,y:any) [status=1] */
func F_nth_put_table (a *ClaireTable,x *ClaireAny,y *ClaireAny) EID { 
    var Result EID
    if ((a.IfWrite != CNULL) && 
        (a.Multivalued_ask != CTRUE)) { 
      Result = F_fastcall_relation2(ToRelation(a.Id()),x,y)
      }  else if (a.Multivalued_ask == CTRUE) { 
      { var r *ClaireAny = F_get_property(C_inverse,ToObject(a.Id()))
        { var old *ClaireSet = ToSet(F_get_table(a,x))
          if ((old.Id() != CNULL) && 
              (r != CNULL)) { 
            { 
              var z *ClaireAny
              _ = z
              var z_support *ClaireSet
              z_support = old
              for i_it := 0; i_it < z_support.Count; i_it++ { 
                z = z_support.At(i_it)
                F_update_dash_relation(ToRelation(r),z,x)
                } 
              } 
            } 
          F_put_table(a,x,y)
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
              loop_1 = F_update_plus_relation(ToRelation(a.Id()),x,z)
              if ErrorIn(loop_1) {Result = loop_1
              break
              } else {
              }
              } 
            } 
          } 
        } 
      } else {
      { var r *ClaireAny = F_get_property(C_inverse,ToObject(a.Id()))
        { var z *ClaireAny = F_get_table(a,x)
          if (Equal(z,y) != CTRUE) { 
            if (r != CNULL) { 
              { var z *ClaireAny = F_get_table(a,x)
                if ((z != CNULL) && 
                    ((r != a.Id()) || 
                        (Equal(x,z) != CTRUE))) { 
                  F_update_dash_relation(ToRelation(r),z,x)
                  } 
                } 
              } 
            F_put_table(a,x,y)
            Result = F_update_plus_relation(ToRelation(a.Id()),x,y)
            } else {
            Result = EID{CFALSE.Id(),0}
            } 
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: nth_put @ table (throw: true) 
func E_nth_put_table (a EID,x EID,y EID) EID { 
    return F_nth_put_table(ToTable(OBJ(a)),ANY(x),ANY(y) )} 
  
// update inverse 
// put does NOT update the inverse, but handles store ...
/* The go function for: put(a:table,x:any,y:any) [status=0] */
func F_put_table (a *ClaireTable,x *ClaireAny,y *ClaireAny)  { 
    { var p *ClaireAny = a.Params
      { var z *ClaireAny = F_get_table(a,x)
        if (Equal(z,y) != CTRUE) { 
          if (C_integer.Id() == p.Isa.Id()) { 
            { var g0071 int = ToInteger(p).Value
              F_store_list(ToList(a.Graph),(ToInteger(x).Value-g0071),y,a.Store_ask)
              } 
            }  else if (p.Isa.IsIn(C_list) == CTRUE) { 
            F_store_list(ToList(a.Graph),F_get_index_table2(a,ToInteger(ToList(x).At(0)).Value,ToInteger(ToList(x).At(1)).Value),y,a.Store_ask)
            } else {
            a.GraphPut(x,y)
            } 
          } 
        } 
      } 
    } 
  
// The EID go function for: put @ table (throw: false) 
func E_put_table (a EID,x EID,y EID) EID { 
    F_put_table(ToTable(OBJ(a)),ANY(x),ANY(y) )
    return EVOID} 
  
// takes care of the defeasible part :)
// adds a value to a multi-valued table: interface method
/* The go function for: add(a:table,x:any,y:any) [status=1] */
func F_add_table (a *ClaireTable,x *ClaireAny,y *ClaireAny) EID { 
    var Result EID
    if (a.Domain.Contains(x) != CTRUE) { 
      Result = ToException(C_general_error.Make(MakeString("[135] ~S does not belong to the domain of ~S").Id(),MakeConstantList(x,a.Id()).Id())).Close()
      } else {
      Result = EID{CFALSE.Id(),0}
      } 
    if !ErrorIn(Result) {
    if (F_member_type(a.Range).Contains(y) != CTRUE) { 
      Result = ToException(C_range_error.Make(a.Id(),y,a.Range.Id())).Close()
      } else {
      Result = EID{CFALSE.Id(),0}
      } 
    if !ErrorIn(Result) {
    Result = F_add_I_table(a,x,y)
    }}
    return Result} 
  
// The EID go function for: add @ table (throw: true) 
func E_add_table (a EID,x EID,y EID) EID { 
    return F_add_table(ToTable(OBJ(a)),ANY(x),ANY(y) )} 
  
// adds a value to a multi-valued table: internal version without type checks
/* The go function for: add!(a:table,x:any,y:any) [status=1] */
func F_add_I_table (a *ClaireTable,x *ClaireAny,y *ClaireAny) EID { 
    var Result EID
    if (a.IfWrite != CNULL) { 
      Result = F_fastcall_relation2(ToRelation(a.Id()),x,y)
      } else {
      { var old *ClaireSet = ToSet(F_get_table(a,x))
        if (F_Core_add_value_I_table(a,x,old,y) == CTRUE) { 
          Result = F_update_plus_relation(ToRelation(a.Id()),x,y)
          } else {
          Result = EID{CFALSE.Id(),0}
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: add! @ table (throw: true) 
func E_add_I_table (a EID,x EID,y EID) EID { 
    return F_add_I_table(ToTable(OBJ(a)),ANY(x),ANY(y) )} 
  
// this methods adds a value to a multi-valued table (used by the compiler)
// s1 is the current value in the table
/* The go function for: add_value!(self:table,x:any,s1:set,y:any) [status=0] */
func F_Core_add_value_I_table (self *ClaireTable,x *ClaireAny,s1 *ClaireSet,y *ClaireAny) *ClaireBoolean { 
    var Result *ClaireBoolean
    if (s1.Contain_ask(y) != CTRUE) { 
      { var s2 *ClaireSet
        { var arg_1 *ClaireSet
          if (self.Store_ask == CTRUE) { 
            arg_1 = s1.Copy()
            } else {
            arg_1 = s1
            } 
          s2 = arg_1.AddFast(y)
          } 
        F_put_table(self,x,s2.Id())
        Result = CTRUE
        } 
      } else {
      Result = CFALSE
      } 
    return Result} 
  
// The EID go function for: add_value! @ table (throw: false) 
func E_Core_add_value_I_table (self EID,x EID,s1 EID,y EID) EID { 
    return EID{F_Core_add_value_I_table(ToTable(OBJ(self)),
      ANY(x),
      ToSet(OBJ(s1)),
      ANY(y) ).Id(),0}} 
  
// a direct version (v3.2) that can be used in lieu of add!
/* The go function for: add_value(self:table,x:any,y:any) [status=0] */
func F_add_value_table3 (self *ClaireTable,x *ClaireAny,y *ClaireAny)  { 
    { var old *ClaireSet = ToSet(F_get_table(self,x))
      F_Core_add_value_I_table(self,x,old,y)
      } 
    } 
  
// The EID go function for: add_value @ table (throw: false) 
func E_add_value_table3 (self EID,x EID,y EID) EID { 
    F_add_value_table3(ToTable(OBJ(self)),ANY(x),ANY(y) )
    return EVOID} 
  
// removes a value from an table (multivalued only)
/* The go function for: delete(a:table,x:any,y:any) [status=0] */
func F_delete_table (a *ClaireTable,x *ClaireAny,y *ClaireAny) *ClaireAny { 
    var Result *ClaireAny
    { var old *ClaireSet = ToSet(F_get_table(a,x))
      if (old.Contain_ask(y) == CTRUE) { 
        { var s *ClaireSet
          { var arg_1 *ClaireSet
            if (a.Store_ask == CTRUE) { 
              arg_1 = old.Copy()
              } else {
              arg_1 = old
              } 
            s = arg_1.Delete(y)
            } 
          F_put_table(a,x,s.Id())
          { var r *ClaireRelation = a.Inverse
            if (r.Id() != CNULL) { 
              F_update_dash_relation(r,y,x)
              } 
            } 
          Result = s.Id()
          } 
        } else {
        Result = old.Id()
        } 
      } 
    return Result} 
  
// The EID go function for: delete @ table (throw: false) 
func E_delete_table (a EID,x EID,y EID) EID { 
    return F_delete_table(ToTable(OBJ(a)),ANY(x),ANY(y) ).ToEID()} 
  
// direct access to 2-dim tables
/* The go function for: nth(a:table,x:any,y:any) [status=1] */
func F_nth_table2 (a *ClaireTable,x *ClaireAny,y *ClaireAny) EID { 
    var Result EID
    { var p *ClaireAny = a.Params
      if (p.Isa.IsIn(C_list) == CTRUE) { 
        if ((ToType(ToList(a.Domain.Id()).At(0)).Contains(x) != CTRUE) || 
            (ToType(ToList(a.Domain.Id()).At(1)).Contains(y) != CTRUE)) { 
          Result = ToException(C_general_error.Make(MakeString("[135] ~S does not belong to the domain of ~S").Id(),MakeConstantList(x,a.Id()).Id())).Close()
          } else {
          Result = EID{CFALSE.Id(),0}
          } 
        if !ErrorIn(Result) {
        Result = F_CALL(C_nth,ARGS(a.Graph.ToEID(),EID{C__INT,IVAL(F_get_index_table2(a,ToInteger(x).Value,ToInteger(y).Value))}))
        }
        } else {
        Result = F_nth_table1(a,MakeTuple(x,y).Id())
        } 
      } 
    return Result} 
  
// The EID go function for: nth @ list<type_expression>(table, any, any) (throw: true) 
func E_nth_table2 (a EID,x EID,y EID) EID { 
    return F_nth_table2(ToTable(OBJ(a)),ANY(x),ANY(y) )} 
  
/* The go function for: nth_table2_type */
func F_nth_table2_type (a *ClaireType,x *ClaireType,y *ClaireType) EID { 
    var Result EID
    if (F_unique_ask_type(a) == CTRUE) { 
      { var arg_1 *ClaireAny
        var try_2 EID
        try_2 = F_the_type(a)
        if ErrorIn(try_2) {Result = try_2
        } else {
        arg_1 = ANY(try_2)
        Result = F_CALL(C_range,ARGS(arg_1.ToEID()))
        }
        } 
      } else {
      Result = EID{C_any.Id(),0}
      } 
    return Result} 
  
  
// The dual EID go function for: "nth_table2_type" 
func E_nth_table2_type (a EID,x EID,y EID) EID { 
    return F_nth_table2_type(ToType(OBJ(a)),ToType(OBJ(x)),ToType(OBJ(y)))} 
  
// sets a value in a 2-dim table
/* The go function for: nth=(a:table,x:any,y:any,z:any) [status=1] */
func F_nth_equal_table2 (a *ClaireTable,x *ClaireAny,y *ClaireAny,z *ClaireAny) EID { 
    var Result EID
    { var p *ClaireAny = a.Params
      if (p.Isa.IsIn(C_list) == CTRUE) { 
        if ((ToType(ToList(a.Domain.Id()).At(0)).Contains(x) != CTRUE) || 
            (ToType(ToList(a.Domain.Id()).At(1)).Contains(y) != CTRUE)) { 
          Result = ToException(C_general_error.Make(MakeString("[135] ~S does not belong to the domain of ~S").Id(),MakeConstantList(MakeConstantList(x,y).Id(),a.Id()).Id())).Close()
          } else {
          Result = EID{CFALSE.Id(),0}
          } 
        if !ErrorIn(Result) {
        if (a.Range.Contains(z) != CTRUE) { 
          Result = ToException(C_range_error.Make(a.Id(),z,a.Range.Id())).Close()
          } else {
          Result = EID{CFALSE.Id(),0}
          } 
        if !ErrorIn(Result) {
        if ((a.Inverse.Id() != CNULL) || 
            (a.IfWrite != CNULL)) { 
          Result = F_nth_put_table(a,MakeConstantList(x,y).Id(),z)
          } else {
          Result = F_store_list(ToList(a.Graph),F_get_index_table2(a,ToInteger(x).Value,ToInteger(y).Value),z,a.Store_ask).ToEID()
          } 
        }}
        } else {
        Result = F_nth_equal_table1(a,MakeTuple(x,y).Id(),z)
        } 
      } 
    return Result} 
  
// The EID go function for: nth= @ list<type_expression>(table, any, any, any) (throw: true) 
func E_nth_equal_table2 (a EID,x EID,y EID,z EID) EID { 
    return F_nth_equal_table2(ToTable(OBJ(a)),
      ANY(x),
      ANY(y),
      ANY(z) )} 
  
// v3.2.16 tuple(a,b) is not list(a,b) !
/* The go function for: get_index(a:table,x:any) [status=0] */
func F_get_index_table1 (a *ClaireTable,x *ClaireAny) int { 
    var Result int
    { var p *ClaireAny = a.Params
      if (C_integer.Id() == p.Isa.Id()) { 
        { var g0078 int = ToInteger(p).Value
          Result = (ToInteger(x).Value-g0078)
          } 
        }  else if (p.Isa.IsIn(C_list) == CTRUE) { 
        Result = F_get_index_table2(a,ToInteger(ToList(x).At(0)).Value,ToInteger(ToList(x).At(1)).Value)
        } else {
        Result = 1
        } 
      } 
    return Result} 
  
// The EID go function for: get_index @ list<type_expression>(table, any) (throw: false) 
func E_get_index_table1 (a EID,x EID) EID { 
    return EID{C__INT,IVAL(F_get_index_table1(ToTable(OBJ(a)),ANY(x) ))}} 
  
/* The go function for: get_index(a:table,x:integer,y:integer) [status=0] */
func F_get_index_table2 (a *ClaireTable,x int,y int) int { 
    var Result int
    { var p *ClaireList = ToList(a.Params)
      Result = (((p.ValuesI()[0]*x)+y)-p.ValuesI()[1])
      } 
    return Result} 
  
// The EID go function for: get_index @ list<type_expression>(table, integer, integer) (throw: false) 
func E_get_index_table2 (a EID,x EID,y EID) EID { 
    return EID{C__INT,IVAL(F_get_index_table2(ToTable(OBJ(a)),INT(x),INT(y) ))}} 
  
// erase an table means to clean its graph so that it becomes empty.
/* The go function for: erase(a:table) [status=1] */
func F_erase_table (a *ClaireTable) EID { 
    var Result EID
    { var p *ClaireAny = a.Params
      if (C_integer.Id() == p.Isa.Id()) { 
        { 
          var i *ClaireAny
          _ = i
          Result= EID{CFALSE.Id(),0}
          var i_support *ClaireList
          var try_1 EID
          try_1 = F_enumerate_any(a.Domain.Id())
          if ErrorIn(try_1) {Result = try_1
          } else {
          i_support = ToList(OBJ(try_1))
          i_len := i_support.Length()
          for i_it := 0; i_it < i_len; i_it++ { 
            i = i_support.At(i_it)
            var loop_2 EID
            _ = loop_2
            loop_2 = F_CALL(C_nth_equal,ARGS(a.Graph.ToEID(),EID{C__INT,IVAL(F_get_index_table1(a,i))},a.Default.ToEID()))
            if ErrorIn(loop_2) {Result = loop_2
            break
            } else {
            }}
            } 
          } 
        }  else if (p.Isa.IsIn(C_list) == CTRUE) { 
        { 
          var l *ClaireList
          _ = l
          var l_iter *ClaireAny
          Result= EID{CFALSE.Id(),0}
          var l_support *ClaireList
          var try_3 EID
          try_3 = F_enumerate_any(a.Domain.Id())
          if ErrorIn(try_3) {Result = try_3
          } else {
          l_support = ToList(OBJ(try_3))
          l_len := l_support.Length()
          for i_it := 0; i_it < l_len; i_it++ { 
            l_iter = l_support.At(i_it)
            l = ToList(l_iter)
            var loop_4 EID
            _ = loop_4
            loop_4 = F_CALL(C_nth_equal,ARGS(a.Graph.ToEID(),EID{C__INT,IVAL(F_get_index_table2(a,ToInteger(l.At(0)).Value,ToInteger(l.At(1)).Value))},a.Default.ToEID()))
            if ErrorIn(loop_4) {Result = loop_4
            break
            } else {
            }}
            } 
          } 
        } else {
        Result = EID{CNIL.Id(),0}
        } 
      } 
    return Result} 
  
// The EID go function for: erase @ table (throw: true) 
func E_erase_table (a EID) EID { 
    return F_erase_table(ToTable(OBJ(a)) )} 
  
// the general case is waiting for the dictionary method like erase(a.graph as dictionary)
// new in v3.2.50 a constructor for building a table dynamically
/* The go function for: make_table(%domain:type,%range:type,%default:any) [status=1] */
func F_make_table_type (_Zdomain *ClaireType,_Zrange *ClaireType,_Zdefault *ClaireAny) EID { 
    var Result EID
    { var t *ClaireTable = ToTable(new(ClaireTable).Is(C_table))
      t.Range = _Zrange
      C_table.Instances = C_table.Instances.AddFast(t.Id())
      t.Domain = _Zdomain
      t.Default = _Zdefault
      t.Params = C_any.Id()
      t.GraphInit()
      Result = EID{t.Id(),0}
      } 
    return Result} 
  
// The EID go function for: make_table @ type (throw: true) 
func E_make_table_type (_Zdomain EID,_Zrange EID,_Zdefault EID) EID { 
    return F_make_table_type(ToType(OBJ(_Zdomain)),ToType(OBJ(_Zrange)),ANY(_Zdefault) )} 
  
// Our first table: a debuging tool which stores a list of stopping values
// *********************************************************************
//   Part 3: Demons & relations for the logic modules                  *
// *********************************************************************
// applying a lambda to one argument
/* The go function for: funcall(self:lambda,x:any) [status=1] */
func F_funcall_lambda1 (self *ClaireLambda,x *ClaireAny) EID { 
    var Result EID
    { var start int = ClEnv.Index
      { var retour int = ClEnv.Base
        ClEnv.Base= start
        ClEnv.Push(x.ToEID())
        F_stack_add(self.Dimension)
        { 
          var val EID
          val = EVAL(self.Body)
          if ErrorIn(val) {Result = val
          } else {
          ClEnv.Base= retour
          ClEnv.Index= start
          Result = val}
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: funcall @ list<type_expression>(lambda, any) (throw: true) 
func E_funcall_lambda1 (self EID,x EID) EID { 
    return F_funcall_lambda1(ToLambda(OBJ(self)),ANY(x) )} 
  
// applying a lambda to two argument
/* The go function for: funcall(self:lambda,x:any,y:any) [status=1] */
func F_funcall_lambda2 (self *ClaireLambda,x *ClaireAny,y *ClaireAny) EID { 
    var Result EID
    { var start int = ClEnv.Index
      { var retour int = ClEnv.Base
        ClEnv.Base= start
        ClEnv.Push(x.ToEID())
        ClEnv.Push(y.ToEID())
        F_stack_add(self.Dimension)
        { 
          var val EID
          val = EVAL(self.Body)
          if ErrorIn(val) {Result = val
          } else {
          ClEnv.Base= retour
          ClEnv.Index= start
          Result = val}
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: funcall @ list<type_expression>(lambda, any, any) (throw: true) 
func E_funcall_lambda2 (self EID,x EID,y EID) EID { 
    return F_funcall_lambda2(ToLambda(OBJ(self)),ANY(x),ANY(y) )} 
  
// applying a lambda to two argument
/* The go function for: funcall(self:lambda,x:any,y:any,z:any) [status=1] */
func F_funcall_lambda3 (self *ClaireLambda,x *ClaireAny,y *ClaireAny,z *ClaireAny) EID { 
    var Result EID
    { var start int = ClEnv.Index
      { var retour int = ClEnv.Base
        ClEnv.Base= start
        ClEnv.Push(x.ToEID())
        ClEnv.Push(y.ToEID())
        ClEnv.Push(z.ToEID())
        F_stack_add(self.Dimension)
        { 
          var val EID
          val = EVAL(self.Body)
          if ErrorIn(val) {Result = val
          } else {
          ClEnv.Base= retour
          ClEnv.Index= start
          Result = val}
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: funcall @ list<type_expression>(lambda, any, any, any) (throw: true) 
func E_funcall_lambda3 (self EID,x EID,y EID,z EID) EID { 
    return F_funcall_lambda3(ToLambda(OBJ(self)),
      ANY(x),
      ANY(y),
      ANY(z) )} 
  
// dealing with inverse
/* The go function for: check_inverse(%r1:any,%r2:any) [status=1] */
func F_check_inverse_any (_Zr1 *ClaireAny,_Zr2 *ClaireAny) EID { 
    var Result EID
    { var r1 *ClaireRelation = ToRelation(_Zr1)
      { var r2 *ClaireRelation = ToRelation(_Zr2)
        r1.Inverse = r2
        r2.Inverse = r1
        F_final_relation(r1)
        F_final_relation(r2)
        var g0085I *ClaireBoolean
        { 
          var v_or4 *ClaireBoolean
          
          { var arg_1 *ClaireBoolean
            { var arg_2 *ClaireType
              if (r2.Multivalued_ask == CTRUE) { 
                arg_2 = F_member_type(r2.Range)
                } else {
                arg_2 = r2.Range
                } 
              arg_1 = r1.Domain.Included(arg_2)
              } 
            v_or4 = arg_1.Not
            } 
          if (v_or4 == CTRUE) {g0085I = CTRUE
          } else { 
            { var arg_3 *ClaireBoolean
              { var arg_4 *ClaireType
                if (r1.Multivalued_ask == CTRUE) { 
                  arg_4 = F_member_type(r1.Range)
                  } else {
                  arg_4 = r1.Range
                  } 
                arg_3 = r2.Domain.Included(arg_4)
                } 
              v_or4 = arg_3.Not
              } 
            if (v_or4 == CTRUE) {g0085I = CTRUE
            } else { 
              g0085I = CFALSE} 
            } 
          } 
        if (g0085I == CTRUE) { 
          Result = ToException(C_general_error.Make(MakeString("[137] ~S and ~S cannot be inverses for one another").Id(),MakeConstantList(r1.Id(),r2.Id()).Id())).Close()
          } else {
          Result = EID{CFALSE.Id(),0}
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: check_inverse @ any (throw: true) 
func E_check_inverse_any (_Zr1 EID,_Zr2 EID) EID { 
    return F_check_inverse_any(ANY(_Zr1),ANY(_Zr2) )} 
  
// very useful
/* The go function for: invert(r:relation,x:any) [status=1] */
func F_invert_relation (r *ClaireRelation,x *ClaireAny) EID { 
    var Result EID
    { var r2 *ClaireAny = F_get_property(C_inverse,ToObject(r.Id()))
      if (C_table.Id() == r2.Isa.Id()) { 
        { var g0086 *ClaireTable = ToTable(r2)
          { var v *ClaireAny
            var try_1 EID
            try_1 = F_nth_table1(g0086,x)
            if ErrorIn(try_1) {Result = try_1
            } else {
            v = ANY(try_1)
            if (g0086.Multivalued_ask.Id() != CFALSE.Id()) { 
              Result = v.ToEID()
              } else {
              Result = EID{MakeConstantSet(v).Id(),0}
              } 
            }
            } 
          } 
        }  else if (r2.Isa.IsIn(C_property) == CTRUE) { 
        { var g0087 *ClaireProperty = ToProperty(r2)
          { var v *ClaireAny = F_get_property(g0087,ToObject(x))
            if (g0087.Multivalued_ask.Id() != CFALSE.Id()) { 
              Result = v.ToEID()
              } else {
              Result = EID{MakeConstantSet(v).Id(),0}
              } 
            } 
          } 
        }  else if (r.Isa.IsIn(C_property) == CTRUE) { 
        { var g0089 *ClaireProperty = ToProperty(r.Id())
          if (g0089.Multivalued_ask.Id() != CFALSE.Id()) { 
            { var z_out *ClaireSet = ToType(CEMPTY.Id()).EmptySet()
              { 
                var z *ClaireAny
                _ = z
                Result= EID{CFALSE.Id(),0}
                var z_support *ClaireList
                var try_2 EID
                try_2 = F_enumerate_any(g0089.Domain.Id())
                if ErrorIn(try_2) {Result = try_2
                } else {
                z_support = ToList(OBJ(try_2))
                z_len := z_support.Length()
                for i_it := 0; i_it < z_len; i_it++ { 
                  z = z_support.At(i_it)
                  if (ToType(F_get_property(g0089,ToObject(z))).Contains(x) == CTRUE) { 
                    z_out.AddFast(z)
                    } 
                  }
                  } 
                } 
              if !ErrorIn(Result) {
              Result = EID{z_out.Id(),0}
              }
              } 
            } else {
            { var z_out *ClaireSet = ToType(CEMPTY.Id()).EmptySet()
              { 
                var z *ClaireAny
                _ = z
                Result= EID{CFALSE.Id(),0}
                var z_support *ClaireList
                var try_3 EID
                try_3 = F_enumerate_any(g0089.Domain.Id())
                if ErrorIn(try_3) {Result = try_3
                } else {
                z_support = ToList(OBJ(try_3))
                z_len := z_support.Length()
                for i_it := 0; i_it < z_len; i_it++ { 
                  z = z_support.At(i_it)
                  if (Equal(F_get_property(g0089,ToObject(z)),x) == CTRUE) { 
                    z_out.AddFast(z)
                    } 
                  }
                  } 
                } 
              if !ErrorIn(Result) {
              Result = EID{z_out.Id(),0}
              }
              } 
            } 
          } 
        }  else if (C_table.Id() == r.Isa.Id()) { 
        { var g0090 *ClaireTable = ToTable(r.Id())
          if (g0090.Multivalued_ask.Id() != CFALSE.Id()) { 
            { var z_out *ClaireSet = ToType(CEMPTY.Id()).EmptySet()
              { 
                var z *ClaireAny
                _ = z
                Result= EID{CFALSE.Id(),0}
                var z_support *ClaireList
                var try_4 EID
                try_4 = F_enumerate_any(g0090.Domain.Id())
                if ErrorIn(try_4) {Result = try_4
                } else {
                z_support = ToList(OBJ(try_4))
                z_len := z_support.Length()
                for i_it := 0; i_it < z_len; i_it++ { 
                  z = z_support.At(i_it)
                  var loop_5 EID
                  _ = loop_5
                  var g0091I *ClaireBoolean
                  var try_6 EID
                  { var arg_7 *ClaireAny
                    var try_8 EID
                    try_8 = F_nth_table1(g0090,z)
                    if ErrorIn(try_8) {try_6 = try_8
                    } else {
                    arg_7 = ANY(try_8)
                    try_6 = EID{ToType(arg_7).Contains(x).Id(),0}
                    }
                    } 
                  if ErrorIn(try_6) {loop_5 = try_6
                  } else {
                  g0091I = ToBoolean(OBJ(try_6))
                  if (g0091I == CTRUE) { 
                    loop_5 = EID{z_out.AddFast(z).Id(),0}
                    } else {
                    loop_5 = EID{CFALSE.Id(),0}
                    } 
                  }
                  if ErrorIn(loop_5) {Result = loop_5
                  break
                  } else {
                  }}
                  } 
                } 
              if !ErrorIn(Result) {
              Result = EID{z_out.Id(),0}
              }
              } 
            } else {
            { var z_out *ClaireSet = ToType(CEMPTY.Id()).EmptySet()
              { 
                var z *ClaireAny
                _ = z
                Result= EID{CFALSE.Id(),0}
                var z_support *ClaireList
                var try_9 EID
                try_9 = F_enumerate_any(g0090.Domain.Id())
                if ErrorIn(try_9) {Result = try_9
                } else {
                z_support = ToList(OBJ(try_9))
                z_len := z_support.Length()
                for i_it := 0; i_it < z_len; i_it++ { 
                  z = z_support.At(i_it)
                  var loop_10 EID
                  _ = loop_10
                  var g0092I *ClaireBoolean
                  var try_11 EID
                  { var arg_12 *ClaireAny
                    var try_13 EID
                    try_13 = F_nth_table1(g0090,z)
                    if ErrorIn(try_13) {try_11 = try_13
                    } else {
                    arg_12 = ANY(try_13)
                    try_11 = EID{Equal(arg_12,x).Id(),0}
                    }
                    } 
                  if ErrorIn(try_11) {loop_10 = try_11
                  } else {
                  g0092I = ToBoolean(OBJ(try_11))
                  if (g0092I == CTRUE) { 
                    loop_10 = EID{z_out.AddFast(z).Id(),0}
                    } else {
                    loop_10 = EID{CFALSE.Id(),0}
                    } 
                  }
                  if ErrorIn(loop_10) {Result = loop_10
                  break
                  } else {
                  }}
                  } 
                } 
              if !ErrorIn(Result) {
              Result = EID{z_out.Id(),0}
              }
              } 
            } 
          } 
        } else {
        Result = EID{CFALSE.Id(),0}
        } 
      } 
    return Result} 
  
// The EID go function for: invert @ relation (throw: true) 
func E_invert_relation (r EID,x EID) EID { 
    return F_invert_relation(ToRelation(OBJ(r)),ANY(x) )} 
  
// same: two useful methods that are used often
/* The go function for: domain!(x:restriction) [status=0] */
func F_domain_I_restriction (x *ClaireRestriction) *ClaireClass { 
    return  ToTypeExpression(x.Domain.ValuesO()[0]).Class_I()
    } 
  
// The EID go function for: domain! @ restriction (throw: false) 
func E_domain_I_restriction (x EID) EID { 
    return EID{F_domain_I_restriction(ToRestriction(OBJ(x)) ).Id(),0}} 
  
/* The go function for: methods(d:class,r:class) [status=0] */
func F_methods_class (d *ClaireClass,r *ClaireClass) *ClaireSet { 
    var Result *ClaireSet
    { var m_out *ClaireSet = ToType(CEMPTY.Id()).EmptySet()
      { 
        var m *ClaireMethod
        _ = m
        var m_iter *ClaireAny
        var m_support *ClaireList
        m_support = C_method.Instances
        m_len := m_support.Length()
        for i_it := 0; i_it < m_len; i_it++ { 
          m_iter = m_support.At(i_it)
          m = ToMethod(m_iter)
          if ((ToType(m.Domain.ValuesO()[0]).Included(ToType(d.Id())) == CTRUE) && 
              (m.Range.Included(ToType(r.Id())) == CTRUE)) { 
            m_out.AddFast(m.Id())
            } 
          } 
        } 
      Result = m_out
      } 
    return Result} 
  
// The EID go function for: methods @ class (throw: false) 
func E_methods_class (d EID,r EID) EID { 
    return EID{F_methods_class(ToClass(OBJ(d)),ToClass(OBJ(r)) ).Id(),0}} 
  
// sets the reified flag
/* The go function for: reify(l:listargs) [status=0] */
func F_reify_listargs (l *ClaireList)  { 
    { 
      var p *ClaireAny
      _ = p
      var p_support *ClaireList
      p_support = ToList(l.Id())
      p_len := p_support.Length()
      for i_it := 0; i_it < p_len; i_it++ { 
        p = p_support.At(i_it)
        if (p.Isa.IsIn(C_property) == CTRUE) { 
          { var g0093 *ClaireProperty = ToProperty(p)
            g0093.Reified = CTRUE
            } 
          } 
        } 
      } 
    } 
  
// The EID go function for: reify @ listargs (throw: false) 
func E_reify_listargs (l EID) EID { 
    F_reify_listargs(ToList(OBJ(l)) )
    return EVOID} 
  
// *********************************************************************
// *   Part 4: Basics of Exceptions                                    *
// *********************************************************************
// args :: property(open = 0)
// value :: property() - defined in kernel
// a generic error that is produced by the error(" ....") instruction
/* The go function for: self_print(self:general_error) [status=1] */
func (self *GeneralError) SelfPrint () EID { 
    var Result EID
    PRINC("**** An error has occurred.\n")
    Result = F_format_string(ToString(self.Cause),ToList(self.Arg))
    if !ErrorIn(Result) {
    PRINC("\n")
    Result = EVOID
    }
    return Result} 
  
// The EID go function for: self_print @ general_error (throw: true) 
func E_self_print_general_error_Core (self EID) EID { 
    return ToGeneralError(OBJ(self)).SelfPrint( )} 
  
// a read_slot error is produced when an unknown value is found
/* The go function for: self_print(self:read_slot_error) [status=1] */
func (self *ReadSlotError) SelfPrint () EID { 
    var Result EID
    PRINC("****[138] The value of ")
    Result = F_CALL(C_print,ARGS(self.Wrong.ToEID()))
    if !ErrorIn(Result) {
    PRINC("(")
    Result = F_CALL(C_print,ARGS(self.Arg.ToEID()))
    if !ErrorIn(Result) {
    PRINC(") is unknown")
    Result = EVOID
    }}
    return Result} 
  
// The EID go function for: self_print @ read_slot_error (throw: true) 
func E_self_print_read_slot_error_Core (self EID) EID { 
    return ToReadSlotError(OBJ(self)).SelfPrint( )} 
  
// range errors
/* The go function for: self_print(self:range_error) [status=1] */
func (self *RangeError) SelfPrint () EID { 
    var Result EID
    PRINC("****[139] ")
    Result = F_CALL(C_print,ARGS(self.Cause.ToEID()))
    if !ErrorIn(Result) {
    PRINC(": range error, ")
    Result = F_CALL(C_print,ARGS(self.Arg.ToEID()))
    if !ErrorIn(Result) {
    PRINC(" does not belong to ")
    Result = F_CALL(C_print,ARGS(self.Wrong.ToEID()))
    if !ErrorIn(Result) {
    PRINC(".\n")
    Result = EVOID
    }}}
    return Result} 
  
// The EID go function for: self_print @ range_error (throw: true) 
func E_self_print_range_error_Core (self EID) EID { 
    return ToRangeError(OBJ(self)).SelfPrint( )} 
  
// selector errors
/* The go function for: self_print(self:selector_error) [status=1] */
func (self *SelectorError) SelfPrint () EID { 
    var Result EID
    { var p *ClaireAny = self.Selector
      if (F_boolean_I_any(ToProperty(p).Restrictions.Id()).Id() != CTRUE.Id()) { 
        PRINC("[140] The property ")
        Result = F_CALL(C_print,ARGS(p.ToEID()))
        if !ErrorIn(Result) {
        PRINC(" is not defined (was applied to ")
        Result = F_CALL(C_print,ARGS(self.Arg.ToEID()))
        if !ErrorIn(Result) {
        PRINC(").\n")
        Result = EVOID
        }}
        } else {
        PRINC("****[141] ")
        Result = F_CALL(C_print,ARGS(self.Arg.ToEID()))
        if !ErrorIn(Result) {
        PRINC(" is a wrong arg list for ")
        Result = F_CALL(C_print,ARGS(p.ToEID()))
        if !ErrorIn(Result) {
        PRINC(".\n")
        Result = EVOID
        }}
        } 
      } 
    return Result} 
  
// The EID go function for: self_print @ selector_error (throw: true) 
func E_self_print_selector_error_Core (self EID) EID { 
    return ToSelectorError(OBJ(self)).SelfPrint( )} 
  
// produced by a return (usually trapped)
/* The go function for: self_print(self:return_error) [status=0] */
func (self *ReturnError) SelfPrint ()  { 
    PRINC("****[142] return called outside of a loop (for or while).")
    } 
  
// The EID go function for: self_print @ return_error (throw: false) 
func E_self_print_return_error_Core (self EID) EID { 
    ToReturnError(OBJ(self)).SelfPrint( )
    return EVOID} 
  
// interpretation of all the error codes
/* The go function for: self_print(self:system_error) [status=1] */
func F_self_print_system_error_Core (self *ClaireSystemError) EID { 
    var Result EID
    { var n int = self.Index
      PRINC("**** An internal error [")
      F_princ_integer(n)
      PRINC("] has occured:\n")
      { var arg_1 *ClaireString
        if (n == 1) { 
          arg_1 = MakeString("Slot Access : ~S(~S) is unknown")
          }  else if (n == 2) { 
          arg_1 = MakeString("dynamic allocation, too large for available memory (~S)")
          }  else if (n == 3) { 
          arg_1 = MakeString("object allocation, too large for available memory (~S)")
          }  else if (n == 5) { 
          arg_1 = MakeString("nth[~S] outside of scope for ~S")
          }  else if (n == 6) { 
          arg_1 = MakeString("invalid slot access: ~S does not have ~S")
          }  else if (n == 7) { 
          arg_1 = MakeString("Skip applied on ~S with a negative argument ~S")
          }  else if (n == 8) { 
          arg_1 = MakeString("List operation: cdr(()) is undefined")
          }  else if (n == 9) { 
          arg_1 = MakeString("String buffer is full: ~S")
          }  else if (n == 10) { 
          arg_1 = MakeString("Cannot create an imported entity from NULL reference")
          }  else if (n == 11) { 
          arg_1 = MakeString("nth_string[~S]: string too short~S")
          }  else if (n == 12) { 
          arg_1 = MakeString("Range check failed: ~S not in ~S")
          }  else if (n == 13) { 
          arg_1 = MakeString("Cannot create a subclass for ~S [~A]")
          }  else if (n == 16) { 
          arg_1 = MakeString("Temporary output string buffer too small")
          }  else if (n == 17) { 
          arg_1 = MakeString("Bag Type Error: ~S does not belong to type ~S")
          }  else if (n == 18) { 
          arg_1 = MakeString("definition of ~S is in conflict with an object from ~S")
          }  else if (n == 19) { 
          arg_1 = MakeString("Integer overflow")
          }  else if (n == 20) { 
          arg_1 = MakeString("Integer arithmetic: division/modulo of ~A by 0")
          }  else if (n == 21) { 
          arg_1 = MakeString("Integer to character: ~S is a wrong value")
          }  else if (n == 22) { 
          arg_1 = MakeString("Cannote create a string with negative length ~S")
          }  else if (n == 23) { 
          arg_1 = MakeString("Not enough memory to instal claire")
          }  else if (n == 24) { 
          arg_1 = MakeString("read unknown value with ~S")
          }  else if (n == 26) { 
          arg_1 = MakeString("Wrong usage of time counter [~A]")
          }  else if (n == 27) { 
          arg_1 = MakeString("internal garbage protection stack overflow")
          }  else if (n == 28) { 
          arg_1 = MakeString("the multivalued status of ~S is not compatible with ~S")
          }  else if (n == 29) { 
          arg_1 = MakeString("There is no module ~S")
          }  else if (n == 30) { 
          arg_1 = MakeString("Attempt to read a private symbol ~S")
          }  else if (n == 31) { 
          arg_1 = MakeString("External function not compiled yet")
          }  else if (n == 32) { 
          arg_1 = MakeString("Too many arguments (~S) for function ~S")
          }  else if (n == 33) { 
          arg_1 = MakeString("Exception handling: stack overflow")
          }  else if (n == 34) { 
          arg_1 = MakeString("User interrupt: EXECUTION ABORTED")
          }  else if (n == 35) { 
          arg_1 = MakeString("reading char '~S': wrong char: ~S")
          }  else if (n == 36) { 
          arg_1 = MakeString("cannot open file ~A")
          }  else if (n == 37) { 
          arg_1 = MakeString("world stack is full")
          }  else if (n == 38) { 
          arg_1 = MakeString("Undefined access to ~S")
          }  else if (n == 39) { 
          arg_1 = MakeString("cannot convert ~S to an integer")
          }  else if (n == 40) { 
          arg_1 = MakeString("integer multiplication overflow with ~S and ~S")
          }  else if (n == 41) { 
          arg_1 = MakeString("wrong NTH access on ~S and ~S")
          }  else if (n == 42) { 
          arg_1 = MakeString("Wrong array[~S] init value: ~S")
          }  else if (n == 43) { 
          arg_1 = MakeString("Defeasible addition on list ~S requires pre-allocation (size ~S)")
          }  else if (n == 50) { 
          arg_1 = MakeString("C++ imported error (~S) : ~S")
          }  else if (n == 300) { 
          arg_1 = MakeString("range error: write fail for ~S, not a ~S")
          } else {
          self.Value = MakeInteger(n).Id()
          arg_1 = MakeString("What the hell is this ! [code: ~S]")
          } 
        Result = F_format_string(arg_1,MakeConstantList(self.Value,self.Arg))
        } 
      } 
    return Result} 
  
// The EID go function for: self_print @ system_error (throw: true) 
func E_self_print_system_error_Core (self EID) EID { 
    return F_self_print_system_error_Core(ToSystemError(OBJ(self)) )} 
  
// contradictions are nice exceptions
/* The go function for: self_print(x:contradiction) [status=0] */
func (x *Contradiction) SelfPrint ()  { 
    PRINC("A contradiction has occured.")
    } 
  
// The EID go function for: self_print @ contradiction (throw: false) 
func E_self_print_contradiction_Core (x EID) EID { 
    ToContradiction(OBJ(x)).SelfPrint( )
    return EVOID} 
  
// the format method is used to print error messages (similar to a printf)
// Note: it would be nice to remove the duplication between format and self_eval@Print
/* The go function for: format(self:string,larg:list) [status=1] */
func F_format_string (self *ClaireString,larg *ClaireList) EID { 
    var Result EID
    { var s *ClaireString = self
      { var n int = F_get_string(s,'~')
        { var l *ClaireList = larg.Copy()
          Result= EID{CFALSE.Id(),0}
          for (n != 0) { 
            var loop_1 EID
            _ = loop_1
            { var m rune = s.At((n+1))
              if (n > 1) { 
                F_princ_string(F_substring_string(s,1,(n-1)))
                } 
              if ('A' == m) { 
                { var arg_2 *ClaireAny
                  var try_3 EID
                  try_3 = F_car_list(l)
                  if ErrorIn(try_3) {loop_1 = try_3
                  } else {
                  arg_2 = ANY(try_3)
                  loop_1 = F_CALL(C_princ,ARGS(arg_2.ToEID()))
                  }
                  } 
                }  else if ('S' == m) { 
                { var arg_4 *ClaireAny
                  var try_5 EID
                  try_5 = F_car_list(l)
                  if ErrorIn(try_5) {loop_1 = try_5
                  } else {
                  arg_4 = ANY(try_5)
                  loop_1 = F_CALL(C_print,ARGS(arg_4.ToEID()))
                  }
                  } 
                }  else if ('F' == m) { 
                { var fv *ClaireAny
                  var try_6 EID
                  try_6 = F_car_list(l)
                  if ErrorIn(try_6) {loop_1 = try_6
                  } else {
                  fv = ANY(try_6)
                  { var p_Z *ClaireBoolean = CFALSE
                    { var j int
                      var try_7 EID
                      { var arg_8 int
                        var try_9 EID
                        { var arg_10 rune
                          var try_11 EID
                          try_11 = F_nth_get_string(s,(n+2),(n+2))
                          if ErrorIn(try_11) {try_9 = try_11
                          } else {
                          arg_10 = CHAR(try_11)
                          try_9 = EID{C__INT,IVAL(int(arg_10))}
                          }
                          } 
                        if ErrorIn(try_9) {try_7 = try_9
                        } else {
                        arg_8 = INT(try_9)
                        try_7 = EID{C__INT,IVAL((arg_8-48))}
                        }
                        } 
                      if ErrorIn(try_7) {loop_1 = try_7
                      } else {
                      j = INT(try_7)
                      if ('%' == s.At((n+2))) { 
                        p_Z = CTRUE
                        j = 1
                        fv = ANY(F_CALL(ToProperty(C__star.Id()),ARGS(fv.ToEID(),EID{C__FLOAT,FVAL(100)})))
                        loop_1 = fv.ToEID()
                        }  else if ((j < 0) || 
                          (j > 9)) { 
                        loop_1 = ToException(C_general_error.Make(MakeString("[189] F requires a single digit integer in ~S").Id(),MakeConstantList((self).Id()).Id())).Close()
                        } else {
                        loop_1 = EID{CFALSE.Id(),0}
                        } 
                      if ErrorIn(loop_1) {Result = loop_1
                      break
                      } else {
                      if ((p_Z != CTRUE) && 
                          ('%' == s.At((n+3)))) { 
                        p_Z = CTRUE
                        fv = ANY(F_CALL(ToProperty(C__star.Id()),ARGS(fv.ToEID(),EID{C__FLOAT,FVAL(100)})))
                        n = (n+1)
                        } 
                      loop_1 = F_CALL(C_mClaire_printFDigit,ARGS(fv.ToEID(),EID{C__INT,IVAL(j)}))
                      if ErrorIn(loop_1) {Result = loop_1
                      break
                      } else {
                      if (p_Z == CTRUE) { 
                        PRINC("%")
                        } 
                      n = (n+1)
                      loop_1 = EID{C__INT,IVAL(n)}
                      }}
                      }
                      } 
                    } 
                  }
                  } 
                }  else if ('I' == m) { 
                loop_1 = ToException(C_general_error.Make(MakeString("[143] ~I not allowed in format").Id(),MakeConstantList(CNULL).Id())).Close()
                } else {
                loop_1 = EID{CFALSE.Id(),0}
                } 
              if ErrorIn(loop_1) {Result = loop_1
              break
              } else {
              l = l.Skip(1)
              s = F_substring_string(s,(n+2),1000)
              n = F_get_string(s,'~')
              loop_1 = EID{C__INT,IVAL(n)}
              }
              } 
            if ErrorIn(loop_1) {Result = loop_1
            break
            } else {
            } 
          }
          if !ErrorIn(Result) {
          if (F_length_string(s) > 0) { 
            F_princ_string(s)
            Result = EVOID
            } else {
            Result = EID{CFALSE.Id(),0}
            } 
          }
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: format @ string (throw: true) 
func E_format_string (self EID,larg EID) EID { 
    return F_format_string(ToString(OBJ(self)),ToList(OBJ(larg)) )} 
  
// special version that prints in the trace port
/* The go function for: tformat(self:string,i:integer,l:list) [status=1] */
func F_tformat_string (self *ClaireString,i int,l *ClaireList) EID { 
    var Result EID
    if (i <= ClEnv.Verbose) { 
      { var p *ClairePort = ClEnv.Ctrace.UseAsOutput()
        Result = F_format_string(self,l)
        if !ErrorIn(Result) {
        Result = p.UseAsOutput().ToEID()
        }
        } 
      } else {
      Result = EID{CFALSE.Id(),0}
      } 
    return Result} 
  
// The EID go function for: tformat @ string (throw: true) 
func E_tformat_string (self EID,i EID,l EID) EID { 
    return F_tformat_string(ToString(OBJ(self)),INT(i),ToList(OBJ(l)) )} 
  
// printing a bag without ( ) separate between sets and lists in CLAIRE4
/* The go function for: princ(s:list) [status=1] */
func F_princ_list (s *ClaireList) EID { 
    var Result EID
    { var f *ClaireBoolean = CTRUE
      { 
        var x *ClaireAny
        _ = x
        Result= EID{CFALSE.Id(),0}
        var x_support *ClaireList
        x_support = s
        x_len := x_support.Length()
        for i_it := 0; i_it < x_len; i_it++ { 
          x = x_support.At(i_it)
          var loop_1 EID
          _ = loop_1
          { 
          if (f == CTRUE) { 
            f = CFALSE
            } else {
            PRINC(",")
            } 
          loop_1 = F_CALL(C_print,ARGS(x.ToEID()))
          if ErrorIn(loop_1) {Result = loop_1
          break
          } else {
          }
          }
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: princ @ list (throw: true) 
func E_princ_list (s EID) EID { 
    return F_princ_list(ToList(OBJ(s)) )} 
  
/* The go function for: princ(s:set) [status=1] */
func F_princ_set (s *ClaireSet) EID { 
    var Result EID
    { var f *ClaireBoolean = CTRUE
      { 
        var x *ClaireAny
        _ = x
        Result= EID{CFALSE.Id(),0}
        var x_support *ClaireSet
        x_support = s
        for i_it := 0; i_it < x_support.Count; i_it++ { 
          x = x_support.At(i_it)
          var loop_1 EID
          _ = loop_1
          { 
          if (f == CTRUE) { 
            f = CFALSE
            } else {
            PRINC(",")
            } 
          loop_1 = F_CALL(C_print,ARGS(x.ToEID()))
          if ErrorIn(loop_1) {Result = loop_1
          break
          } else {
          }
          }
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: princ @ set (throw: true) 
func E_princ_set (s EID) EID { 
    return F_princ_set(ToSet(OBJ(s)) )} 
  
// a global variable is a named object with a special evaluation
// NOTE: we need to refine the scheme for global constants !
// GV are defeasible
/* The go function for: close(self:global_variable) [status=1] */
func (self *GlobalVariable) Close () EID { 
    var Result EID
    if ((self.Value != CNULL) && 
        ((ToType(C_set.Id()).Contains(self.Range.Id()) != CTRUE) && 
            (self.Range.Contains(self.Value) != CTRUE))) { 
      { var _CL_obj *RangeError = ToRangeError(new(RangeError).Is(C_range_error))
        _CL_obj.Arg = self.Value
        _CL_obj.Cause = self.Id()
        _CL_obj.Wrong = self.Range.Id()
        Result = _CL_obj.Close()
        } 
      } else {
      Result = EID{CFALSE.Id(),0}
      } 
    if !ErrorIn(Result) {
    Result = EID{self.Id(),0}
    }
    return Result} 
  
// The EID go function for: close @ global_variable (throw: true) 
func E_close_global_variable (self EID) EID { 
    return ToGlobalVariable(OBJ(self)).Close( )} 
  
/* The go function for: self_eval(self:global_variable) [status=0] */
func (self *GlobalVariable) SelfEval () EID { 
    var Result EID
    Result = self.Value.ToEID()
    return Result} 
  
// The EID go function for: self_eval @ global_variable (throw: true) 
func E_self_eval_global_variable (self EID) EID { 
    return ToGlobalVariable(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: global_variable 
func EVAL_global_variable (x *ClaireAny) EID { 
     return ToGlobalVariable(x).SelfEval()} 
  
// -> moved from pretty.cl
// we create a spcial contraidiction that we shall reuse
// how to use it
/* The go function for: contradiction!(_CL_obj:void) [status=1] */
func F_contradiction_I_void () EID { 
    var Result EID
    Result = ToContradiction(C_contradiction_occurs.Value).Close()
    return Result} 
  
// The EID go function for: contradiction! @ void (throw: true) 
func E_contradiction_I_void (_CL_obj EID) EID { 
    return F_contradiction_I_void( )} 
  
// v0.01
// end of file