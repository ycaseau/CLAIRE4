/***** CLAIRE Compilation of file /Users/ycaseau/Dropbox/src/clairev4.03/src/meta/function.cl 
         [version 4.0.03 / safety 5] Wednesday 12-29-2021 08:34:14 *****/

package Core
import (_ "fmt"
	. "Kernel"
)

//+-------------------------------------------------------------+
//| CLAIRE                                                      |
//| function.cl                                                 |
//| Copyright (C) 1994 - 2021 Yves Caseau. All Rights Reserved  |
//| cf. copyright info in file object.cl: about()               |
//+-------------------------------------------------------------+
// --------------------------------------------------------------------
// This file hold the methods that are defined by an external function
// and those that apply to a primitive type
// --------------------------------------------------------------------
// *********************************************************************
// *  Table of contents                                                *
// *   Part 1: Basics of pretty printing                               *
// *   Part 2: Methods for CLAIRE objects                              *
// *   Part 3: System Methods                                          *
// *   Part 4: Methods for Native entities                             *
// *********************************************************************
// we find here what is necessary for the minimal kernel of CLAIRE
// ==============================================================
// min :: operation(precedence = 20)
// max :: operation(precedence = 20)
// *********************************************************************
// *   Part 1: Basics of pretty printing                               *
// *********************************************************************
// we use a nice object
// support reccursive print-in-string 
// print :: property() - defined in Kernel
// buffered print
// new in v3.3.26: unbounded recursion is supported :-)
/* {1} The go function for: print_in_string(_CL_obj:void) [status=0] */
func F_print_in_string_void ()  { 
    // procedure body with s = void
    { var n int  = (C_pretty.Cprevious+1)
      { var p1 *ClairePort  
        if (n < C_pretty.Cpstack.Length()) { 
          p1 = ToPort(C_pretty.Cpstack.At((n+1)-1))
          } else {
          p1 = F_port_I_void()
          } 
        { var p2 *ClairePort   = p1.UseAsOutput()
          C_pretty.Cprevious = n
          C_pretty.Cpretty = p1
          if (Equal(C_pretty.Cpstack.Id(),CNIL.Id()) == CTRUE) { 
            C_pretty.Cpstack = MakeList(ToType(C_port.Id()),p2.Id(),p1.Id())
            } else {
            ToArray(C_pretty.Cpstack.Id()).NthPut(n,p2.Id())
            if (n == C_pretty.Cpstack.Length()) { 
              C_pretty.Cpstack = C_pretty.Cpstack.AddFast(p1.Id())
              } 
            } 
          } 
        } 
      } 
    } 
  
// The EID go function for: print_in_string @ void (throw: false) 
func E_print_in_string_void (_CL_obj EID) EID { 
    F_print_in_string_void( )
    return EVOID} 
  
/* {1} The go function for: end_of_string(_CL_obj:void) [status=1] */
func F_end_of_string_void () EID { 
    // eid body s = string
    var Result EID 
    if (C_pretty.Cprevious == 0) { 
      Result = ToException(C_general_error.Make(MakeString("[123] unbalanced use of print-in-string").Id(),CNIL.Id())).Close()
      } else {
      Result = EID{CFALSE.Id(),0}
      } 
    if !ErrorIn(Result) {
    { var n int  = C_pretty.Cprevious
      { var s *ClaireString   = C_pretty.Cpretty.String_I()
        { var p *ClairePort   = ToPort(C_pretty.Cpstack.At(n-1))
          C_pretty.Cpretty.SetLength(0)
          p.UseAsOutput()
          C_pretty.Cpretty = p
          C_pretty.Cprevious = (C_pretty.Cprevious-1)
          Result = EID{(s).Id(),0}
          } 
        } 
      } 
    }
    return Result} 
  
// The EID go function for: end_of_string @ void (throw: true) 
func E_end_of_string_void (_CL_obj EID) EID { 
    return F_end_of_string_void( )} 
  
/* {1} The go function for: mClaire/buffer_length(_CL_obj:void) [status=0] */
func F_buffer_length_void () int { 
    return  C_pretty.Cpretty.Length()
    } 
  
// The EID go function for: mClaire/buffer_length @ void (throw: false) 
func E_buffer_length_void (_CL_obj EID) EID { 
    return EID{C__INT,IVAL(F_buffer_length_void( ))}} 
  
/* {1} The go function for: mClaire/buffer_set_length(i:integer) [status=0] */
func F_buffer_set_length_integer (i int)  { 
    // procedure body with s = void
    C_pretty.Cpretty.SetLength(i)
    } 
  
// The EID go function for: mClaire/buffer_set_length @ integer (throw: false) 
func E_buffer_set_length_integer (i EID) EID { 
    F_buffer_set_length_integer(INT(i) )
    return EVOID} 
  
// a method for calling the printer without issuing a message
// here we assume that self_print is always defined as a function
/* {1} The go function for: apply_self_print(self:any) [status=1] */
func F_apply_self_print_any (self *ClaireAny ) EID { 
    // eid body s = void
    var Result EID 
    if (self == CNULL) { 
      PRINC("unknown")
      Result = EVOID
      }  else if (self.Isa.IsIn(C_thing) == CTRUE) { 
      { var g0096 *ClaireThing   = ToThing(self)
        g0096.Name.Princ()
        Result = EVOID
        } 
      }  else if (C_class.Id() == self.Isa.Id()) { 
      { var g0097 *ClaireClass   = ToClass(self)
        g0097.Name.Princ()
        Result = EVOID
        } 
      }  else if (C_integer.Id() == self.Isa.Id()) { 
      { var g0098 int  = ToInteger(self).Value
        F_princ_integer(g0098)
        Result = EVOID
        } 
      }  else if (C_string.Id() == self.Isa.Id()) { 
      { var g0099 *ClaireString   = ToString(self)
        Result = F_CALL(C_self_print,ARGS(EID{(g0099).Id(),0}))
        } 
      } else {
      { var _Zprop *ClaireMethod   = ToMethod(F__at_property1(C_self_print,self.Isa).Id())
        if ((F_boolean_I_any(_Zprop.Id()) == CTRUE) && 
            (_Zprop.Functional.Id() != CNULL)) { 
          Result = F_funcall1(_Zprop.Functional,self)
          } else {
          { 
            h_index := ClEnv.Index
            h_base := ClEnv.Base
            Result = F_CALL(C_self_print,ARGS(self.ToEID()))
            if ErrorIn(Result){ 
              ClEnv.Index = h_index
              ClEnv.Base = h_base
              PRINC("<unprintable:")
              Result = F_print_any(self.Isa.Id())
              if !ErrorIn(Result) {
              PRINC(">")
              Result = EVOID
              }
              } 
            } 
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: apply_self_print @ any (throw: true) 
func E_apply_self_print_any (self EID) EID { 
    return F_apply_self_print_any(ANY(self) )} 
  
// some basic definitions
/* {1} The go function for: self_print(self:any) [status=1] */
func F_self_print_any_Core (self *ClaireAny ) EID { 
    // eid body s = void
    var Result EID 
    if (self == CNULL) { 
      PRINC("unknown")
      Result = EVOID
      } else {
      { var c *ClaireClass   = self.Isa
        { var n int  = c.Params.Length()
          if (n > 0) { 
            Result = F_print_any(c.Id())
            if !ErrorIn(Result) {
            PRINC("(")
            { var i int  = 1
              { var g0101 int  = n
                Result= EID{CFALSE.Id(),0}
                for (i <= g0101) { 
                  var loop_1 EID 
                  _ = loop_1
                  { 
                  { var arg_2 *ClaireAny  
                    var try_3 EID 
                    try_3 = F_CALL(C_get,ARGS(c.Params.At(i-1).ToEID(),self.ToEID()))
                    if ErrorIn(try_3) {loop_1 = try_3
                    } else {
                    arg_2 = ANY(try_3)
                    loop_1 = F_CALL(C_print,ARGS(arg_2.ToEID()))
                    }
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
            Result = EVOID
            }}
            } else {
            PRINC("<")
            Result = F_print_any(c.Id())
            if !ErrorIn(Result) {
            PRINC(">")
            Result = EVOID
            }
            } 
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: self_print @ any (throw: true) 
func E_self_print_any_Core (self EID) EID { 
    return F_self_print_any_Core(ANY(self) )} 
  
/* {1} The go function for: self_print(self:boolean) [status=0] */
func F_self_print_boolean_Core (self *ClaireBoolean )  { 
    // procedure body with s = void
    if (self == CTRUE) { 
      PRINC("true")
      } else {
      PRINC("false")
      } 
    } 
  
// The EID go function for: self_print @ boolean (throw: false) 
func E_self_print_boolean_Core (self EID) EID { 
    F_self_print_boolean_Core(ToBoolean(OBJ(self)) )
    return EVOID} 
  
/* {1} The go function for: self_print(self:function) [status=0] */
func F_self_print_function_Core (self *ClaireFunction )  { 
    // procedure body with s = void
    PRINC("#'")
    F_princ_string(F_string_I_function(self))
    PRINC("")
    } 
  
// The EID go function for: self_print @ function (throw: false) 
func E_self_print_function_Core (self EID) EID { 
    F_self_print_function_Core(ToFunction(OBJ(self)) )
    return EVOID} 
  
// prints the name of a restriction. If we have a close property and if a
// short-cut is possible, we use it.
/* {1} The go function for: self_print(self:restriction) [status=1] */
func F_self_print_restriction_Core (self *ClaireRestriction ) EID { 
    // eid body s = void
    var Result EID 
    if ((self.Selector.Id() == CNULL) || 
        (self.Domain.Id() == CNULL)) { 
      PRINC("<")
      Result = F_print_any(self.Id().Isa.Id())
      if !ErrorIn(Result) {
      PRINC(">")
      Result = EVOID
      }
      } else {
      { var p *ClaireProperty   = self.Selector
        { var n int  = 0
          { var c *ClaireClass   = F_domain_I_restriction(self)
            { 
              var r *ClaireRestriction  
              _ = r
              var r_iter *ClaireAny  
              for _,r_iter = range(p.Restrictions.ValuesO()){ 
                r = ToRestriction(r_iter)
                if (F_domain_I_restriction(r).Id() == c.Id()) { 
                  n = (n+1)
                  } 
                } 
              } 
            p.Name.Princ()
            PRINC(" @ ")
            { var arg_1 *ClaireType  
              if (n == 1) { 
                arg_1 = ToType(c.Id())
                } else {
                arg_1 = ToType(self.Domain.Id())
                } 
              Result = F_print_any(arg_1.Id())
              } 
            if !ErrorIn(Result) {
            PRINC("")
            Result = EVOID
            }
            } 
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: self_print @ restriction (throw: true) 
func E_self_print_restriction_Core (self EID) EID { 
    return F_self_print_restriction_Core(ToRestriction(OBJ(self)) )} 
  
// we are too far
/* {1} The go function for: print(x:any) [status=1] */
func F_print_any (x *ClaireAny ) EID { 
    // eid body s = void
    var Result EID 
    { var _Zl int  = C_pretty.Index
      if ((C_pretty.Pbreak == CTRUE) && 
          (C_pretty.Pprint == CTRUE)) { 
        { var b_index int  = F_buffer_length_void()
          { var missed *ClaireBoolean   = CFALSE
            if (F_short_enough_integer((b_index+10)) != CTRUE) { 
              C_pretty.Pprint = CFALSE
              C_pretty.Pbreak = CFALSE
              Result = F_CALL(C_print,ARGS(x.ToEID()))
              if !ErrorIn(Result) {
              C_pretty.Pbreak = CTRUE
              { 
                var va_arg1 *PrettyPrinter  
                var va_arg2 *ClaireBoolean  
                va_arg1 = C_pretty
                va_arg2 = CTRUE
                va_arg1.Pprint = va_arg2
                Result = EID{va_arg2.Id(),0}
                } 
              }
              } else {
              { 
                h_index := ClEnv.Index
                h_base := ClEnv.Base
                C_pretty.Pbreak = CFALSE
                Result = F_apply_self_print_any(x)
                if !ErrorIn(Result) {
                { 
                  var va_arg1 *PrettyPrinter  
                  var va_arg2 *ClaireBoolean  
                  va_arg1 = C_pretty
                  va_arg2 = CTRUE
                  va_arg1.Pbreak = va_arg2
                  Result = EID{va_arg2.Id(),0}
                  } 
                }
                if ErrorIn(Result) && ToType(C_much_too_far.Id()).Contains(ANY(Result)) == CTRUE { 
                  ClEnv.Index = h_index
                  ClEnv.Base = h_base
                  missed = CTRUE
                  Result = EID{missed.Id(),0}
                  } 
                } 
              if !ErrorIn(Result) {
              if (missed == CTRUE) { 
                C_pretty.Pprint = CTRUE
                C_pretty.Pbreak = CTRUE
                F_buffer_set_length_integer(b_index)
                C_pretty.Index = _Zl
                Result = F_apply_self_print_any(x)
                } else {
                Result = EID{CFALSE.Id(),0}
                } 
              }
              } 
            } 
          } 
        } else {
        Result = F_apply_self_print_any(x)
        } 
      if !ErrorIn(Result) {
      C_pretty.Index = _Zl
      Result = EID{CNULL,0}
      }
      } 
    return Result} 
  
// The EID go function for: print @ any (throw: true) 
func E_print_any (x EID) EID { 
    return F_print_any(ANY(x) )} 
  
// short_enough = we expect that what we want to print is short enough (more that 10 chars to the width)
/* {1} The go function for: short_enough(self:integer) [status=0] */
func F_short_enough_integer (self int) *ClaireBoolean  { 
    if (self < C_pretty.Width) {return CTRUE
    } else {return CFALSE}} 
  
// The EID go function for: short_enough @ integer (throw: false) 
func E_short_enough_integer (self EID) EID { 
    return EID{F_short_enough_integer(INT(self) ).Id(),0}} 
  
// *********************************************************************
// *   Part 2: Methods for CLAIRE objects                              *
// *********************************************************************
// the instantiation body is a sequence of words from which the initialization
// of the object must be built.
//  copied_def = object (for object) + float (for float) + integer (for all)
//               + NULL for objects
// v3.2.12: use a condition that is coherent with ClReflect.cl : a slot defaut value must be placed
// unless it is a copied_def
// CLAIRE4 (fix OLD bug): lp is list of p such that o.p := unknown is requested !
/* {1} The go function for: new_defaults(self:object,lp:list) [status=1] */
func F_Core_new_defaults_object (self *ClaireObject ,lp *ClaireList ) EID { 
    // eid body s = object
    var Result EID 
    { 
      var s *ClaireSlot  
      _ = s
      var s_iter *ClaireAny  
      Result= EID{CFALSE.Id(),0}
      for _,s_iter = range(self.Isa.Slots.ValuesO()){ 
        s = ToSlot(s_iter)
        var loop_1 EID 
        _ = loop_1
        { var p *ClaireProperty   = s.Selector
          { var s2 *ClaireClass   = s.Srange
            { var d *ClaireAny   = s.Default
              if (d != CNULL) { 
                { var v *ClaireAny   = self.SlotGet(s.Index,s.Srange)
                  if ((v == CNULL) && 
                      (lp.Memq(p.Id()) == CTRUE)) { 
                    loop_1 = EID{CNIL.Id(),0}
                    }  else if ((v == CNULL) && 
                      ((s2.Id() != C_object.Id()) && 
                          ((C_integer.Id() != d.Isa.Id()) && 
                            (s2.Id() != C_float.Id())))) { 
                    loop_1 = F_update_property(p,
                      self,
                      s.Index,
                      s.Srange,
                      d)
                    }  else if (Equal(d,v) == CTRUE) { 
                    if (p.Multivalued_ask == CTRUE) { 
                      { 
                        var y *ClaireAny  
                        _ = y
                        loop_1= EID{CFALSE.Id(),0}
                        var y_support *ClaireList  
                        var try_2 EID 
                        try_2 = F_enumerate_any(d)
                        if ErrorIn(try_2) {loop_1 = try_2
                        } else {
                        y_support = ToList(OBJ(try_2))
                        y_len := y_support.Length()
                        for i_it := 0; i_it < y_len; i_it++ { 
                          y = y_support.At(i_it)
                          var loop_3 EID 
                          _ = loop_3
                          loop_3 = F_update_plus_relation(ToRelation(p.Id()),self.Id(),y)
                          if ErrorIn(loop_3) {loop_1 = loop_3
                          break
                          } else {
                          }}
                          } 
                        } 
                      } else {
                      loop_1 = F_update_plus_relation(ToRelation(p.Id()),self.Id(),d)
                      } 
                    } else {
                    loop_1 = EID{CFALSE.Id(),0}
                    } 
                  } 
                } else {
                loop_1 = EID{CFALSE.Id(),0}
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
    { var m *ClaireAny   = F__at_property1(C_close,self.Id().Isa).Id()
      if (C_method.Id() == m.Isa.Id()) { 
        { var g0102 *ClaireMethod   = ToMethod(m)
          Result = F_funcall_method1(g0102,self.Id())
          } 
        } else {
        Result = EID{self.Id(),0}
        } 
      } 
    }
    return Result} 
  
// The EID go function for: new_defaults @ object (throw: true) 
func E_Core_new_defaults_object (self EID,lp EID) EID { 
    return F_Core_new_defaults_object(ToObject(OBJ(self)),ToList(OBJ(lp)) )} 
  
// v3.0.41  obviously
//-------------------------- ENTITY   --------------------------------------
/* {1} The go function for: not(self:any) [status=0] */
func F_not_any (self *ClaireAny ) *ClaireBoolean  { 
    if (self == CTRUE.Id()) { 
      if (CFALSE == CTRUE) {return CTRUE
      } else {return CFALSE}}  else if (self == CFALSE.Id()) { 
      if (CTRUE == CTRUE) {return CTRUE
      } else {return CFALSE}}  else if (F_boolean_I_any(self).Id() != CTRUE.Id()) { 
      if (CTRUE == CTRUE) {return CTRUE
      } else {return CFALSE}} else {
      if (CFALSE == CTRUE) {return CTRUE
      } else {return CFALSE}} 
    } 
  
// The EID go function for: not @ any (throw: false) 
func E_not_any (self EID) EID { 
    return EID{F_not_any(ANY(self) ).Id(),0}} 
  
/* {1} The go function for: !=(self:any,x:any) [status=0] */
func F__I_equal_any (self *ClaireAny ,x *ClaireAny ) *ClaireBoolean  { 
    if (Equal(self,x) == CTRUE) { 
      if (CFALSE == CTRUE) {return CTRUE
      } else {return CFALSE}} else {
      if (CTRUE == CTRUE) {return CTRUE
      } else {return CFALSE}} 
    } 
  
// The EID go function for: != @ any (throw: false) 
func E__I_equal_any (self EID,x EID) EID { 
    return EID{F__I_equal_any(ANY(self),ANY(x) ).Id(),0}} 
  
// gives the type of any object. This is open_coded.
/* {1} The go function for: owner(self:any) [status=0] */
func F_owner_any (self *ClaireAny ) *ClaireClass  { 
    return  self.Isa
    } 
  
// The EID go function for: owner @ any (throw: false) 
func E_owner_any (self EID) EID { 
    return EID{F_owner_any(ANY(self) ).Id(),0}} 
  
// some useful methods
/* {1} The go function for: known?(self:any) [status=0] */
func F_known_ask_any (self *ClaireAny ) *ClaireBoolean  { 
    if (CNULL != self) {return CTRUE
    } else {return CFALSE}} 
  
// The EID go function for: known? @ any (throw: false) 
func E_known_ask_any (self EID) EID { 
    return EID{F_known_ask_any(ANY(self) ).Id(),0}} 
  
/* {1} The go function for: unknown?(self:any) [status=0] */
func F_unknown_ask_any (self *ClaireAny ) *ClaireBoolean  { 
    if (CNULL == self) {return CTRUE
    } else {return CFALSE}} 
  
// The EID go function for: unknown? @ any (throw: false) 
func E_unknown_ask_any (self EID) EID { 
    return EID{F_unknown_ask_any(ANY(self) ).Id(),0}} 
  
// needed by the compiled code for casts (inserted by ocontrol for dynamic type checks)
// Claire 4: TODO - add a second order type
/* {1} The go function for: check_in(self:any,y:type) [status=1] */
func F_check_in_any (self *ClaireAny ,y *ClaireType ) EID { 
    // eid body s = any
    var Result EID 
    if (y.Contains(self) == CTRUE) { 
      Result = self.ToEID()
      } else {
      Result = ToException(C_general_error.Make(MakeString("[124] the value ~S does not belong to the range ~S").Id(),MakeConstantList(self,y.Id()).Id())).Close()
      } 
    return Result} 
  
// The EID go function for: check_in @ any (throw: true) 
func E_check_in_any (self EID,y EID) EID { 
    return F_check_in_any(ANY(self),ToType(OBJ(y)) )} 
  
// used to cast dynamically a non-mutable bag to a typed mutable bag
// claire 4 : aplied to list and sets (bags)
/* {1} The go function for: check_in(self:bag,c:class,y:type) [status=1] */
func F_check_in_bag (self *ClaireBag ,c *ClaireClass ,y *ClaireType ) EID { 
    // eid body s = bag
    var Result EID 
    { var OK *ClaireBoolean  
      if (self.Isa.IsIn(C_list) == CTRUE) { 
        { var g0104 *ClaireList   = ToList(self.Id())
          { var arg_1 *ClaireAny  
            { 
              var z *ClaireAny  
              _ = z
              arg_1= CFALSE.Id()
              var z_support *ClaireList  
              z_support = g0104
              z_len := z_support.Length()
              for i_it := 0; i_it < z_len; i_it++ { 
                z = z_support.At(i_it)
                if (y.Contains(z) != CTRUE) { 
                  arg_1 = CTRUE.Id()
                  break
                  } 
                } 
              } 
            OK = F_not_any(arg_1)
            } 
          } 
        } else {
        { var arg_2 *ClaireAny  
          { 
            var z *ClaireAny  
            _ = z
            arg_2= CFALSE.Id()
            var z_support *ClaireSet  
            z_support = ToSet(self.Id())
            for i_it := 0; i_it < z_support.Count; i_it++ { 
              z = z_support.At(i_it)
              if (y.Contains(z) != CTRUE) { 
                arg_2 = CTRUE.Id()
                break
                } 
              } 
            } 
          OK = F_not_any(arg_2)
          } 
        } 
      if (OK == CTRUE) { 
        Result = EID{self.Cast_I(y).Id(),0}
        } else {
        Result = ToException(C_general_error.Make(MakeString("[124] the value ~S does not belong to subtype[~S]").Id(),MakeConstantList(self.Id(),y.Id()).Id())).Close()
        } 
      } 
    return Result} 
  
// The EID go function for: check_in @ bag (throw: true) 
func E_check_in_bag (self EID,c EID,y EID) EID { 
    return F_check_in_bag(ToBag(OBJ(self)),ToClass(OBJ(c)),ToType(OBJ(y)) )} 
  
// new in v3.00.48
/* {1} The go function for: <(self:any,x:any) [status=1] */
func F__inf_any (self *ClaireAny ,x *ClaireAny ) EID { 
    // eid body s = boolean
    var Result EID 
    if (Equal(self,x) == CTRUE) { 
      Result = EID{CFALSE.Id(),0}
      } else {
      Result = F_CALL(ToProperty(C__inf_equal.Id()),ARGS(self.ToEID(),x.ToEID()))
      } 
    return Result} 
  
// The EID go function for: < @ any (throw: true) 
func E__inf_any (self EID,x EID) EID { 
    return F__inf_any(ANY(self),ANY(x) )} 
  
/* {1} The go function for: >(self:any,x:any) [status=1] */
func F__sup_any (self *ClaireAny ,x *ClaireAny ) EID { 
    // eid body s = boolean
    var Result EID 
    if (Equal(self,x) == CTRUE) { 
      Result = EID{CFALSE.Id(),0}
      } else {
      Result = F_CALL(ToProperty(C__inf_equal.Id()),ARGS(x.ToEID(),self.ToEID()))
      } 
    return Result} 
  
// The EID go function for: > @ any (throw: true) 
func E__sup_any (self EID,x EID) EID { 
    return F__sup_any(ANY(self),ANY(x) )} 
  
// >= is defined as a macro in file.cl
// unsafe is a pragma : it tells the compiler not to catch a possible error 
// to add in the documentation
/* {1} The go function for: unsafe(x:any) [status=0] */
func F_unsafe_any (x *ClaireAny ) *ClaireAny  { 
    return  x
    } 
  
// The EID go function for: unsafe @ any (throw: false) 
func E_unsafe_any (x EID) EID { 
    return F_unsafe_any(ANY(x) ).ToEID()} 
  
// ----------------------- CLASS ---------------------------------------------
// classes have a "open" status with three states
// forward : -2          - use open as a marker for forward definition
// close : -1           ????
// abstract :  open = 0  - cannot create an instance
// final : open = 1  - no more subclasses or instance
// default : open : 2 - default for ephemeral - no instance
// open :  open = 3  - default for things (named things can be queried)
// declares a class as ephemeral: the member set is not maintained
// v3.2.14 recusively applies to subclasses
/* {1} The go function for: ephemeral(self:class) [status=1] */
func F_ephemeral_class (self *ClaireClass ) EID { 
    // eid body s = any
    var Result EID 
    { 
      var c *ClaireClass  
      _ = c
      var c_iter *ClaireAny  
      Result= EID{CFALSE.Id(),0}
      var c_support *ClaireSet  
      c_support = self.Descendents
      for i_it := 0; i_it < c_support.Count; i_it++ { 
        c_iter = c_support.At(i_it)
        c = ToClass(c_iter)
        var loop_1 EID 
        _ = loop_1
        if (c.Instances.Length() != 0) { 
          loop_1 = ToException(C_general_error.Make(MakeString("[187] cannot declare ~S as ephemeral because of ~S has instances").Id(),MakeConstantList(self.Id(),c.Id()).Id())).Close()
          } else {
          { 
            var va_arg1 *ClaireClass  
            var va_arg2 int 
            va_arg1 = c
            va_arg2 = ClEnv.Default
            va_arg1.Open = va_arg2
            loop_1 = EID{C__INT,IVAL(va_arg2)}
            } 
          } 
        if ErrorIn(loop_1) {Result = loop_1
        break
        } else {
        }
        } 
      } 
    return Result} 
  
// The EID go function for: ephemeral @ class (throw: true) 
func E_ephemeral_class (self EID) EID { 
    return F_ephemeral_class(ToClass(OBJ(self)) )} 
  
// claire4 : opposite of ephemeral(c)
/* {1} The go function for: instanced(c:class) [status=1] */
func F_instanced_class (c *ClaireClass ) EID { 
    // eid body s = any
    var Result EID 
    { var n int  = c.Open
      if ((n == ClEnv.Default) || 
          (n == ClEnv.Open)) { 
        { 
          var va_arg1 *ClaireClass  
          var va_arg2 int 
          va_arg1 = c
          va_arg2 = ClEnv.Open
          va_arg1.Open = va_arg2
          Result = EID{C__INT,IVAL(va_arg2)}
          } 
        } else {
        Result = ToException(C_general_error.Make(MakeString("[125] abstract classes cannot be instanced").Id(),CNIL.Id())).Close()
        } 
      } 
    return Result} 
  
// The EID go function for: instanced @ class (throw: true) 
func E_instanced_class (c EID) EID { 
    return F_instanced_class(ToClass(OBJ(c)) )} 
  
// declares a class as an abtract class (without instances)
/* {1} The go function for: abstract(c:class) [status=1] */
func F_abstract_class (c *ClaireClass ) EID { 
    // eid body s = any
    var Result EID 
    { var n int  = c.Open
      if (n == ClEnv.Close) { 
        Result = ToException(C_general_error.Make(MakeString("[125] closed classes cannot be abstract").Id(),CNIL.Id())).Close()
        }  else if (c.Instances.Length() != 0) { 
        Result = ToException(C_general_error.Make(MakeString("[125] instanced classes cannot be abstract").Id(),CNIL.Id())).Close()
        } else {
        { 
          var va_arg1 *ClaireClass  
          var va_arg2 int 
          va_arg1 = c
          va_arg2 = ClEnv.ABSTRACT
          va_arg1.Open = va_arg2
          Result = EID{C__INT,IVAL(va_arg2)}
          } 
        } 
      if !ErrorIn(Result) {
      Result = EID{c.Id(),0}
      }
      } 
    return Result} 
  
// The EID go function for: abstract @ class (throw: true) 
func E_abstract_class (c EID) EID { 
    return F_abstract_class(ToClass(OBJ(c)) )} 
  
// declares a class with no subclasses (apply to things)
/* {1} The go function for: final(c:class) [status=1] */
func F_final_class (c *ClaireClass ) EID { 
    // eid body s = any
    var Result EID 
    { var n int  = c.Open
      if (n == ClEnv.Default) { 
        Result = ToException(C_general_error.Make(MakeString("[125] ephemetral classes cannot be final").Id(),CNIL.Id())).Close()
        }  else if (c.Subclass.Size() != 0) { 
        Result = ToException(C_general_error.Make(MakeString("[125] a class with subclasses cannot be final").Id(),CNIL.Id())).Close()
        } else {
        { 
          var va_arg1 *ClaireClass  
          var va_arg2 int 
          va_arg1 = c
          va_arg2 = ClEnv.Final
          va_arg1.Open = va_arg2
          Result = EID{C__INT,IVAL(va_arg2)}
          } 
        } 
      if !ErrorIn(Result) {
      Result = EID{c.Id(),0}
      }
      } 
    return Result} 
  
// The EID go function for: final @ class (throw: true) 
func E_final_class (c EID) EID { 
    return F_final_class(ToClass(OBJ(c)) )} 
  
// instantiation with and without a name
// new! is a method-less property that is managed by the compiler
/* {1} The go function for: new(self:class) [status=1] */
func F_new_class1 (self *ClaireClass ) EID { 
    // eid body s = object
    var Result EID 
    { var o *ClaireObject  
      var try_1 EID 
      if (self.Open <= 0) { 
        try_1 = ToException(C_general_error.Make(MakeString("[105] cannot instantiate ~S").Id(),MakeConstantList(self.Id()).Id())).Close()
        } else {
        try_1 = EID{CFALSE.Id(),0}
        } 
      if !ErrorIn(try_1) {
      try_1 = EID{F_new_object_class(self).Id(),0}
      }
      if ErrorIn(try_1) {Result = try_1
      } else {
      o = ToObject(OBJ(try_1))
      Result = F_Core_new_defaults_object(o,CNIL)
      }
      } 
    return Result} 
  
// The EID go function for: new @ list<type_expression>(class) (throw: true) 
func E_new_class1 (self EID) EID { 
    return F_new_class1(ToClass(OBJ(self)) )} 
  
/* {1} The go function for: new_class1_type */
func F_new_class1_type (self *ClaireType ) EID { 
    var Result EID 
    Result = EID{F_glb_class(C_object,F_member_type(self)).Id(),0}
    return Result} 
  
  
// The dual EID go function for: "new_class1_type" 
func E_new_class1_type (self EID) EID { 
    return F_new_class1_type(ToType(OBJ(self)))} 
  
// v3.2.26
/* {1} The go function for: new(self:class,%nom:symbol) [status=1] */
func F_new_class2 (self *ClaireClass ,_Znom *ClaireSymbol ) EID { 
    // eid body s = thing
    var Result EID 
    { var o *ClaireThing  
      var try_1 EID 
      if (self.Open <= 0) { 
        try_1 = ToException(C_general_error.Make(MakeString("[105] cannot instantiate ~S").Id(),MakeConstantList(self.Id()).Id())).Close()
        } else {
        try_1 = EID{CFALSE.Id(),0}
        } 
      if !ErrorIn(try_1) {
      try_1 = F_new_thing_class(self,_Znom)
      }
      if ErrorIn(try_1) {Result = try_1
      } else {
      o = ToThing(OBJ(try_1))
      Result = F_Core_new_defaults_object(ToObject(o.Id()),CNIL)
      }
      } 
    return Result} 
  
// The EID go function for: new @ list<type_expression>(class, symbol) (throw: true) 
func E_new_class2 (self EID,_Znom EID) EID { 
    return F_new_class2(ToClass(OBJ(self)),ToSymbol(OBJ(_Znom)) )} 
  
/* {1} The go function for: new_class2_type */
func F_new_class2_type (self *ClaireType ,_Znom *ClaireType ) EID { 
    var Result EID 
    Result = EID{F_glb_class(C_thing,F_member_type(self)).Id(),0}
    return Result} 
  
  
// The dual EID go function for: "new_class2_type" 
func E_new_class2_type (self EID,_Znom EID) EID { 
    return F_new_class2_type(ToType(OBJ(self)),ToType(OBJ(_Znom)))} 
  
// v3.2.26
// internal version => moved to Kernel
// mClaire/new!(self:class) :  type[object glb member(self)]
//  -> function!(new_object_class)
// mClaire/new!(self:class,%nom:symbol) :  type[thing glb member(self)]
//   -> function!(new_thing_class)
// the smallest super_set of two sets
// there is always any, so it always returns a class
/* {1} The go function for: meet(self:class,ens:class) [status=0] */
func F_meet_class (self *ClaireClass ,ens *ClaireClass ) *ClaireClass  { 
    // procedure body with s = class
    var Result *ClaireClass  
    { var l1 *ClaireList   = self.Ancestors
      { var l2 *ClaireList   = ens.Ancestors
        { var m int 
          if (l1.Length() < l2.Length()) { 
            m = l1.Length()
            } else {
            m = l2.Length()
            } 
          for (l1.ValuesO()[m-1] != l2.ValuesO()[m-1]) { 
            m = (m-1)
            } 
          Result = ToClass(l1.ValuesO()[m-1])
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: meet @ class (throw: false) 
func E_meet_class (self EID,ens EID) EID { 
    return EID{F_meet_class(ToClass(OBJ(self)),ToClass(OBJ(ens)) ).Id(),0}} 
  
// fast inclusion method for lattice_sets (lattice order). The argument is
// either a lattice_set or {}
/* {1} The go function for: inherit?(self:class,ens:class) [status=0] */
func F_inherit_ask_class (self *ClaireClass ,ens *ClaireClass ) *ClaireBoolean  { 
    // procedure body with s = boolean
    var Result *ClaireBoolean  
    { var l *ClaireList   = self.Ancestors
      { var n int  = ens.Ancestors.Length()
        Result = MakeBoolean((n <= l.Length()) && (l.ValuesO()[n-1] == ens.Id()))
        } 
      } 
    return Result} 
  
// The EID go function for: inherit? @ class (throw: false) 
func E_inherit_ask_class (self EID,ens EID) EID { 
    return EID{F_inherit_ask_class(ToClass(OBJ(self)),ToClass(OBJ(ens)) ).Id(),0}} 
  
//------------- PROPERTY ---------------------------------------------------
// the two methods to access open(r)
// an abstract property is extensible and can receive new restrictions
/* {1} The go function for: abstract(p:property) [status=1] */
func F_abstract_property (p *ClaireProperty ) EID { 
    // eid body s = any
    var Result EID 
    { var n int  = p.Open
      if (n < 2) { 
        Result = ToException(C_general_error.Make(MakeString("[127] ~S can no longer become abstract").Id(),MakeConstantList(p.Id()).Id())).Close()
        } else {
        { 
          var va_arg1 *ClaireRelation  
          var va_arg2 int 
          va_arg1 = ToRelation(p.Id())
          va_arg2 = 3
          va_arg1.Open = va_arg2
          Result = EID{C__INT,IVAL(va_arg2)}
          } 
        } 
      if !ErrorIn(Result) {
      Result = EID{p.Id(),0}
      }
      } 
    return Result} 
  
// The EID go function for: abstract @ property (throw: true) 
func E_abstract_property (p EID) EID { 
    return F_abstract_property(ToProperty(OBJ(p)) )} 
  
// a final property is completely defined and cannot receive a new restriction
// v3.2.04: the new value 4 will be used to represent (compiled but open)
/* {1} The go function for: final(r:relation) [status=0] */
func F_final_relation (r *ClaireRelation )  { 
    // procedure body with s = void
    if (r.Isa.IsIn(C_property) == CTRUE) { 
      { var g0106 *ClaireProperty   = ToProperty(r.Id())
        if (g0106.Open <= 2) { 
          g0106.Open = 1
          { 
            var va_arg1 *ClaireRelation  
            var va_arg2 *ClaireType  
            va_arg1 = ToRelation(g0106.Id())
            { var arg_1 *ClaireList  
              { 
                var v_list7 *ClaireList  
                var x *ClaireRestriction  
                var v_local7 *ClaireAny  
                v_list7 = g0106.Restrictions
                arg_1 = CreateList(ToType(CEMPTY.Id()),v_list7.Length())
                for CLcount := 0; CLcount < v_list7.Length(); CLcount++{ 
                  x = ToRestriction(v_list7.At(CLcount))
                  v_local7 = x.Domain.ValuesO()[1-1]
                  arg_1.PutAt(CLcount,v_local7)
                  } 
                } 
              va_arg2 = F_Uall_list(arg_1)
              } 
            va_arg1.Domain = va_arg2
            } 
          { 
            var va_arg1 *ClaireRelation  
            var va_arg2 *ClaireType  
            va_arg1 = ToRelation(g0106.Id())
            { var arg_2 *ClaireList  
              { 
                var v_list7 *ClaireList  
                var x *ClaireRestriction  
                var v_local7 *ClaireAny  
                v_list7 = g0106.Restrictions
                arg_2 = CreateList(ToType(CEMPTY.Id()),v_list7.Length())
                for CLcount := 0; CLcount < v_list7.Length(); CLcount++{ 
                  x = ToRestriction(v_list7.At(CLcount))
                  v_local7 = x.Range.Id()
                  arg_2.PutAt(CLcount,v_local7)
                  } 
                } 
              va_arg2 = F_Uall_list(arg_2)
              } 
            va_arg1.Range = va_arg2
            } 
          } 
        } 
      } 
    } 
  
// The EID go function for: final @ relation (throw: false) 
func E_final_relation (r EID) EID { 
    F_final_relation(ToRelation(OBJ(r)) )
    return EVOID} 
  
//------------- MODULES   --------------------------------------------------
// book-keeping for a module : based on parts/part_of hierarchy
// propagates uses declaration + register a new associated namespace.
/* {1} The go function for: close(self:module) [status=0] */
func F_close_module (self *ClaireModule ) *ClaireModule  { 
    if (self.Id() != C_claire.Id()) { 
      if (self.PartOf.Id() != CNULL) { 
        { var sup *ClaireModule   = self.PartOf
          sup.Parts = sup.Parts.AddFast(self.Id())
          { 
            var x *ClaireAny  
            _ = x
            var x_support *ClaireList  
            x_support = sup.Uses
            x_len := x_support.Length()
            for i_it := 0; i_it < x_len; i_it++ { 
              x = x_support.At(i_it)
              if ((ToBoolean(self.Uses.Contain_ask(x).Id()) != CTRUE) && 
                  (x.Isa.IsIn(C_module) == CTRUE)) { 
                self.Uses = self.Uses.AddFast(x)
                } 
              } 
            } 
          } 
        } 
      } 
    self.Namespace()
    return  self
    } 
  
// The EID go function for: close @ module (throw: false) 
func E_close_module (self EID) EID { 
    return EID{F_close_module(ToModule(OBJ(self)) ).Id(),0}} 
  
// note: dynamic modules are no longer supported
/* {1} The go function for: get_symbol(self:string) [status=0] */
func F_get_symbol_string (self *ClaireString ) *ClaireAny  { 
    return  F_get_symbol_module(C_claire,self)
    } 
  
// The EID go function for: get_symbol @ string (throw: false) 
func E_get_symbol_string (self EID) EID { 
    return F_get_symbol_string(ToString(OBJ(self)) ).ToEID()} 
  
// *********************************************************************
// *   Part 3: System Methods                                          *
// *********************************************************************
// all these methods will be open-coded by the compiler
//get_stack(self:integer) : any -> get_stack(self)
//put_stack(self:integer,x:any) : any -> put_stack(self, x)
//push!(self:meta_system,x:any) : void -> push!(self, x)
//stat() : void -> function!(claire_stat)
/* {1} The go function for: gensym(self:void) [status=0] */
func F_gensym_void () *ClaireSymbol  { 
    return  F_gensym_string(MakeString("g"))
    } 
  
// The EID go function for: gensym @ void (throw: false) 
func E_gensym_void (self EID) EID { 
    return EID{F_gensym_void( ).Id(),0}} 
  
// world management
/* {1} The go function for: store(l:list,n:integer,y:any) [status=1] */
func F_store_list4 (l *ClaireList ,n int,y *ClaireAny ) EID { 
    // eid body s = any
    var Result EID 
    if ((n < 1) || 
        (n > l.Length())) { 
      Result = ToException(C_general_error.Make(MakeString("store @ list: ~S out of bounds for ~S").Id(),MakeConstantList(MakeInteger(n).Id(),l.Id()).Id())).Close()
      } else {
      Result = F_store_list(l,n,y,CTRUE).ToEID()
      } 
    return Result} 
  
// The EID go function for: store @ list<type_expression>(list, integer, any) (throw: true) 
func E_store_list4 (l EID,n EID,y EID) EID { 
    return F_store_list4(ToList(OBJ(l)),INT(n),ANY(y) )} 
  
/* {1} The go function for: store(l:array,n:integer,y:any) [status=1] */
func F_store_array1 (l *ClaireList ,n int,y *ClaireAny ) EID { 
    // eid body s = any
    var Result EID 
    if ((n < 1) || 
        (n > l.Length())) { 
      Result = ToException(C_general_error.Make(MakeString("store @ array: ~S out of bounds for ~S").Id(),MakeConstantList(MakeInteger(n).Id(),l.Id()).Id())).Close()
      } else {
      Result = F_store_list(l,n,y,CTRUE).ToEID()
      } 
    return Result} 
  
// The EID go function for: store @ list<type_expression>(array, integer, any) (throw: true) 
func E_store_array1 (l EID,n EID,y EID) EID { 
    return F_store_array1(ToArray(OBJ(l)),INT(n),ANY(y) )} 
  
/* {1} The go function for: commit(n:integer) [status=0] */
func F_commit_integer (n int)  { 
    // procedure body with s = void
    for (n < F_world_number()) { 
      F_world_remove()
      } 
    } 
  
// The EID go function for: commit @ integer (throw: false) 
func E_commit_integer (n EID) EID { 
    F_commit_integer(INT(n) )
    return EVOID} 
  
/* {1} The go function for: backtrack(n:integer) [status=0] */
func F_backtrack_integer (n int)  { 
    // procedure body with s = void
    for (n < F_world_number()) { 
      F_world_pop()
      } 
    } 
  
// The EID go function for: backtrack @ integer (throw: false) 
func E_backtrack_integer (n EID) EID { 
    F_backtrack_integer(INT(n) )
    return EVOID} 
  
// allows to change the storage class
/* {1} The go function for: store(l:listargs) [status=0] */
func F_store_listargs (l *ClaireList ) *ClaireAny  { 
    // procedure body with s = any
    var Result *ClaireAny  
    { 
      var r *ClaireAny  
      _ = r
      Result= CFALSE.Id()
      var r_support *ClaireList  
      r_support = ToList(l.Id())
      r_len := r_support.Length()
      for i_it := 0; i_it < r_len; i_it++ { 
        r = r_support.At(i_it)
        if (r.Isa.IsIn(C_relation) == CTRUE) { 
          { var g0109 *ClaireRelation   = ToRelation(r)
            g0109.Store_ask = CTRUE
            } 
          }  else if (C_string.Id() == r.Isa.Id()) { 
          { var g0110 *ClaireString   = ToString(r)
            { var v *ClaireAny   = F_value_string(g0110)
              if (v.Isa.IsIn(C_global_variable) == CTRUE) { 
                { var g0111 *GlobalVariable   = ToGlobalVariable(v)
                  g0111.Store_ask = CTRUE
                  } 
                } 
              } 
            } 
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: store @ listargs (throw: false) 
func E_store_listargs (l EID) EID { 
    return F_store_listargs(ToList(OBJ(l)) ).ToEID()} 
  
// *********************************************************************
// *   Part 4: Methods for Native entities                             *
// *********************************************************************
//------------------- STRING -----------------------------------------------
// length(self:string) : integer -> function!(length_string)  -> in Kernel
// make_function(self:string) : function -> function!(make_function_string)
/* {1} The go function for: symbol!(self:string) [status=0] */
func F_symbol_I_string2 (self *ClaireString ) *ClaireSymbol  { 
    return  F_symbol_I_string(self,C_claire)
    } 
  
// The EID go function for: symbol! @ list<type_expression>(string) (throw: false) 
func E_symbol_I_string2 (self EID) EID { 
    return EID{F_symbol_I_string2(ToString(OBJ(self)) ).Id(),0}} 
  
// deprecated in claire 4 - do not use a string as a byte buffer
/* {1} The go function for: nth_get(s:string,n:integer,max:integer) [status=1] */
func F_nth_get_string (s *ClaireString ,n int,max int) EID { 
    // eid body s = char
    var Result EID 
    if (n <= max) { 
      Result = EID{C__CHAR,CVAL(s.At(n))}
      } else {
      Result = ToException(C_general_error.Make(MakeString("Buffer string access").Id(),CNIL.Id())).Close()
      } 
    return Result} 
  
// The EID go function for: nth_get @ string (throw: true) 
func E_nth_get_string (s EID,n EID,max EID) EID { 
    return F_nth_get_string(ToString(OBJ(s)),INT(n),INT(max) )} 
  
/* {1} The go function for: nth_put(s:string,n:integer,c:char,max:integer) [status=1] */
func F_nth_put_string (s *ClaireString ,n int,c rune,max int) EID { 
    // eid body s = void
    var Result EID 
    if (n <= max) { 
      F_nth_set_string(s,n,c)
      Result = EVOID
      } else {
      Result = ToException(C_general_error.Make(MakeString("Buffer string access").Id(),CNIL.Id())).Close()
      } 
    return Result} 
  
// The EID go function for: nth_put @ string (throw: true) 
func E_nth_put_string (s EID,n EID,c EID,max EID) EID { 
    return F_nth_put_string(ToString(OBJ(s)),
      INT(n),
      CHAR(c),
      INT(max) )} 
  
// shell(self:string) : void -> function!(claire_shell)
// value @ string no longer supported in CLAIRE 4
//  v3.2.14
// we keep the externC method name even if it now support go code
/* {1} The go function for: externC(s:string) [status=1] */
func F_externC_string (s *ClaireString ) EID { 
    // eid body s = void
    var Result EID 
    Result = ToException(C_general_error.Make(MakeString("cannot execute Go code: ~A").Id(),MakeConstantList((s).Id()).Id())).Close()
    return Result} 
  
// The EID go function for: externC @ list<type_expression>(string) (throw: true) 
func E_externC_string (s EID) EID { 
    return F_externC_string(ToString(OBJ(s)) )} 
  
/* {1} The go function for: externC(s:string,c:class) [status=1] */
func F_externC_string2 (s *ClaireString ,c *ClaireClass ) EID { 
    // eid body s = any
    var Result EID 
    Result = ToException(C_general_error.Make(MakeString("cannot execute ~A").Id(),MakeConstantList((s).Id()).Id())).Close()
    if !ErrorIn(Result) {
    Result = EID{CNULL,0}
    }
    return Result} 
  
// The EID go function for: externC @ list<type_expression>(string, class) (throw: true) 
func E_externC_string2 (s EID,c EID) EID { 
    return F_externC_string2(ToString(OBJ(s)),ToClass(OBJ(c)) )} 
  
/* {1} The go function for: externC_string2_type */
func F_externC_string2_type (s *ClaireType ,c *ClaireType ) EID { 
    var Result EID 
    Result = EID{F_member_type(c).Id(),0}
    return Result} 
  
  
// The dual EID go function for: "externC_string2_type" 
func E_externC_string2_type (s EID,c EID) EID { 
    return F_externC_string2_type(ToType(OBJ(s)),ToType(OBJ(c)))} 
  
//------------------- SYMBOL -----------------------------------------------
/* {1} The go function for: make_string(self:symbol) [status=1] */
func F_make_string_symbol (self *ClaireSymbol ) EID { 
    // eid body s = string
    var Result EID 
    F_print_in_string_void()
    self.Princ()
    Result = F_end_of_string_void()
    return Result} 
  
// The EID go function for: make_string @ symbol (throw: true) 
func E_make_string_symbol (self EID) EID { 
    return F_make_string_symbol(ToSymbol(OBJ(self)) )} 
  
//princ(self:symbol) : any -> function!(princ_symbol)
/* {1} The go function for: self_print(self:symbol) [status=1] */
func F_self_print_symbol_Core (self *ClaireSymbol ) EID { 
    // eid body s = void
    var Result EID 
    self.Module_I().Name.Princ()
    PRINC("/")
    Result = F_print_any((self.String_I()).Id())
    if !ErrorIn(Result) {
    PRINC("")
    Result = EVOID
    }
    return Result} 
  
// The EID go function for: self_print @ symbol (throw: true) 
func E_self_print_symbol_Core (self EID) EID { 
    return F_self_print_symbol_Core(ToSymbol(OBJ(self)) )} 
  
//c_princ(self:symbol) : any -> function!(c_princ_symbol)
//gensym(self:string) : symbol -> function!(gensym_string, NEW_ALLOC)
//--------------------- INTEGER -----------------------------------------
/* {1} The go function for: +(self:integer,x:integer) [status=0] */
func F__plus_integer (self int,x int) int { 
    return  (self+x)
    } 
  
// The EID go function for: + @ list<type_expression>(integer, integer) (throw: false) 
func E__plus_integer (self EID,x EID) EID { 
    return EID{C__INT,IVAL(F__plus_integer(INT(self),INT(x) ))}} 
  
/* {1} The go function for: _plus_integer_type */
func F__plus_integer_type (self *ClaireType ,x *ClaireType ) EID { 
    var Result EID 
    Result = EID{F_abstract_type_operation(C__plus,self,x).Id(),0}
    return Result} 
  
  
// The dual EID go function for: "_plus_integer_type" 
func E__plus_integer_type (self EID,x EID) EID { 
    return F__plus_integer_type(ToType(OBJ(self)),ToType(OBJ(x)))} 
  
/* {1} The go function for: -(self:integer,x:integer) [status=0] */
func F__dash_integer1 (self int,x int) int { 
    return  (self-x)
    } 
  
// The EID go function for: - @ list<type_expression>(integer, integer) (throw: false) 
func E__dash_integer1 (self EID,x EID) EID { 
    return EID{C__INT,IVAL(F__dash_integer1(INT(self),INT(x) ))}} 
  
/* {1} The go function for: _dash_integer1_type */
func F__dash_integer1_type (self *ClaireType ,x *ClaireType ) EID { 
    var Result EID 
    Result = EID{F_abstract_type_operation(C__dash,self,x).Id(),0}
    return Result} 
  
  
// The dual EID go function for: "_dash_integer1_type" 
func E__dash_integer1_type (self EID,x EID) EID { 
    return F__dash_integer1_type(ToType(OBJ(self)),ToType(OBJ(x)))} 
  
//-(self:integer) : integer -> function!(ch_sign)
//float!(self:integer) : float -> function!(to_float)
//mod(self:integer,x:integer) : integer -> function!(mod_integer)
//less_code(n:integer,i:integer) : boolean -> function!(less_code_integer)
/* {1} The go function for: <<(x:integer,y:integer) [status=0] */
func F__inf_inf_integer (x int,y int) int { 
    return  (x << y)
    } 
  
// The EID go function for: << @ integer (throw: false) 
func E__inf_inf_integer (x EID,y EID) EID { 
    return EID{C__INT,IVAL(F__inf_inf_integer(INT(x),INT(y) ))}} 
  
/* {1} The go function for: >>(x:integer,y:integer) [status=0] */
func F__sup_sup_integer (x int,y int) int { 
    return  (x >> y)
    } 
  
// The EID go function for: >> @ integer (throw: false) 
func E__sup_sup_integer (x EID,y EID) EID { 
    return EID{C__INT,IVAL(F__sup_sup_integer(INT(x),INT(y) ))}} 
  
/* {1} The go function for: and(x:integer,y:integer) [status=0] */
func F_and_integer (x int,y int) int { 
    return  (x & y)
    } 
  
// The EID go function for: and @ integer (throw: false) 
func E_and_integer (x EID,y EID) EID { 
    return EID{C__INT,IVAL(F_and_integer(INT(x),INT(y) ))}} 
  
/* {1} The go function for: or(x:integer,y:integer) [status=0] */
func F_or_integer (x int,y int) int { 
    return  (x | y)
    } 
  
// The EID go function for: or @ integer (throw: false) 
func E_or_integer (x EID,y EID) EID { 
    return EID{C__INT,IVAL(F_or_integer(INT(x),INT(y) ))}} 
  
// open-coded
/* {1} The go function for: <(self:integer,x:integer) [status=0] */
func F__inf_integer (self int,x int) *ClaireBoolean  { 
    if (self < x) { 
      if (CTRUE == CTRUE) {return CTRUE
      } else {return CFALSE}} else {
      if (CFALSE == CTRUE) {return CTRUE
      } else {return CFALSE}} 
    } 
  
// The EID go function for: < @ integer (throw: false) 
func E__inf_integer (self EID,x EID) EID { 
    return EID{F__inf_integer(INT(self),INT(x) ).Id(),0}} 
  
/* {1} The go function for: <=(self:integer,x:integer) [status=0] */
func F__inf_equal_integer (self int,x int) *ClaireBoolean  { 
    if (self <= x) { 
      if (CTRUE == CTRUE) {return CTRUE
      } else {return CFALSE}} else {
      if (CFALSE == CTRUE) {return CTRUE
      } else {return CFALSE}} 
    } 
  
// The EID go function for: <= @ integer (throw: false) 
func E__inf_equal_integer (self EID,x EID) EID { 
    return EID{F__inf_equal_integer(INT(self),INT(x) ).Id(),0}} 
  
/* {1} The go function for: >(self:integer,x:integer) [status=0] */
func F__sup_integer (self int,x int) *ClaireBoolean  { 
    if (self > x) { 
      if (CTRUE == CTRUE) {return CTRUE
      } else {return CFALSE}} else {
      if (CFALSE == CTRUE) {return CTRUE
      } else {return CFALSE}} 
    } 
  
// The EID go function for: > @ integer (throw: false) 
func E__sup_integer (self EID,x EID) EID { 
    return EID{F__sup_integer(INT(self),INT(x) ).Id(),0}} 
  
/* {1} The go function for: nth(self:integer,y:integer) [status=0] */
func F_nth_integer (self int,y int) *ClaireBoolean  { 
    if (BitVectorContains(self,y) == CTRUE) { 
      if (CTRUE == CTRUE) {return CTRUE
      } else {return CFALSE}} else {
      if (CFALSE == CTRUE) {return CTRUE
      } else {return CFALSE}} 
    } 
  
// The EID go function for: nth @ integer (throw: false) 
func E_nth_integer (self EID,y EID) EID { 
    return EID{F_nth_integer(INT(self),INT(y) ).Id(),0}} 
  
/* {1} The go function for: abs(x:integer) [status=0] */
func F_abs_integer (x int) int { 
    if (x >= 0) { 
      return  x
      } else {
      return  (-x)
      } 
    } 
  
// The EID go function for: abs @ integer (throw: false) 
func E_abs_integer (x EID) EID { 
    return EID{C__INT,IVAL(F_abs_integer(INT(x) ))}} 
  
/* {1} The go function for: random(a:integer,b:integer) [status=0] */
func F_random_integer2 (a int,b int) int { 
    return  (a+F_random_integer(((b+1)-a)))
    } 
  
// The EID go function for: random @ list<type_expression>(integer, integer) (throw: false) 
func E_random_integer2 (a EID,b EID) EID { 
    return EID{C__INT,IVAL(F_random_integer2(INT(a),INT(b) ))}} 
  
// used by the logic
/* {1} The go function for: factor?(x:integer,y:integer) [status=1] */
func F_factor_ask_integer (x int,y int) EID { 
    // eid body s = boolean
    var Result EID 
    { var arg_1 int 
      var try_2 EID 
      try_2 = EID{C__INT,IVAL((x%y))}
      if ErrorIn(try_2) {Result = try_2
      } else {
      arg_1 = INT(try_2)
      Result = EID{Equal(MakeInteger(arg_1).Id(),MakeInteger(0).Id()).Id(),0}
      }
      } 
    return Result} 
  
// The EID go function for: factor? @ integer (throw: true) 
func E_factor_ask_integer (x EID,y EID) EID { 
    return F_factor_ask_integer(INT(x),INT(y) )} 
  
/* {1} The go function for: divide?(x:integer,y:integer) [status=1] */
func F_divide_ask_integer (x int,y int) EID { 
    // eid body s = boolean
    var Result EID 
    { var arg_1 int 
      var try_2 EID 
      try_2 = EID{C__INT,IVAL((y%x))}
      if ErrorIn(try_2) {Result = try_2
      } else {
      arg_1 = INT(try_2)
      Result = EID{Equal(MakeInteger(arg_1).Id(),MakeInteger(0).Id()).Id(),0}
      }
      } 
    return Result} 
  
// The EID go function for: divide? @ integer (throw: true) 
func E_divide_ask_integer (x EID,y EID) EID { 
    return F_divide_ask_integer(INT(x),INT(y) )} 
  
/* {1} The go function for: Id(x:any) [status=0] */
func F_Id_any (x *ClaireAny ) *ClaireAny  { 
    return  x
    } 
  
// The EID go function for: Id @ any (throw: false) 
func E_Id_any (x EID) EID { 
    return F_Id_any(ANY(x) ).ToEID()} 
  
/* {1} The go function for: Id_any_type */
func F_Id_any_type (x *ClaireType ) EID { 
    var Result EID 
    Result = EID{x.Id(),0}
    return Result} 
  
  
// The dual EID go function for: "Id_any_type" 
func E_Id_any_type (x EID) EID { 
    return F_Id_any_type(ToType(OBJ(x)))} 
  
//------------------------ FLOAT ---------------------------------------------
/* {1} The go function for: +(self:float,x:float) [status=0] */
func F__plus_float (self float64,x float64) float64 { 
    return  (self+x)
    } 
  
// The EID go function for: + @ list<type_expression>(float, float) (throw: false) 
func E__plus_float (self EID,x EID) EID { 
    return EID{C__FLOAT,FVAL(F__plus_float(FLOAT(self),FLOAT(x) ))}} 
  
/* {1} The go function for: -(self:float,x:float) [status=0] */
func F__dash_float (self float64,x float64) float64 { 
    return  (self-x)
    } 
  
// The EID go function for: - @ list<type_expression>(float, float) (throw: false) 
func E__dash_float (self EID,x EID) EID { 
    return EID{C__FLOAT,FVAL(F__dash_float(FLOAT(self),FLOAT(x) ))}} 
  
/* {1} The go function for: *(self:float,x:float) [status=0] */
func F__star_float (self float64,x float64) float64 { 
    return  (self*x)
    } 
  
// The EID go function for: * @ list<type_expression>(float, float) (throw: false) 
func E__star_float (self EID,x EID) EID { 
    return EID{C__FLOAT,FVAL(F__star_float(FLOAT(self),FLOAT(x) ))}} 
  
/* {1} The go function for: /(self:float,x:float) [status=0] */
func F__7_float (self float64,x float64) float64 { 
    return  (self/x)
    } 
  
// The EID go function for: / @ list<type_expression>(float, float) (throw: false) 
func E__7_float (self EID,x EID) EID { 
    return EID{C__FLOAT,FVAL(F__7_float(FLOAT(self),FLOAT(x) ))}} 
  
// old junk
// +(self:float,x:float) : float -> (let y:float := (self + x) in y)
// -(self:float,x:float) : float -> (let y:float := (self - x) in y)
// *(self:float,x:float) : float -> (let y:float := (self * x) in y)
// /(self:float,x:float) : float -> (let y:float := (self / x) in y)
/* {1} The go function for: -(self:float) [status=0] */
func F__dash_float2 (self float64) float64 { 
    return  ((-1)*self)
    } 
  
// The EID go function for: - @ list<type_expression>(float) (throw: false) 
func E__dash_float2 (self EID) EID { 
    return EID{C__FLOAT,FVAL(F__dash_float2(FLOAT(self) ))}} 
  
/* {1} The go function for: string!(self:float) [status=1] */
func F_string_I_float (self float64) EID { 
    // eid body s = string
    var Result EID 
    F_print_in_string_void()
    F_princ_float(self)
    Result = F_end_of_string_void()
    return Result} 
  
// The EID go function for: string! @ float (throw: true) 
func E_string_I_float (self EID) EID { 
    return F_string_I_float(FLOAT(self) )} 
  
// v3.3.42
/* {1} The go function for: abs(x:float) [status=0] */
func F_abs_float (x float64) float64 { 
    if (x >= 0) { 
      return  x
      } else {
      return  (-x)
      } 
    } 
  
// The EID go function for: abs @ float (throw: false) 
func E_abs_float (x EID) EID { 
    return EID{C__FLOAT,FVAL(F_abs_float(FLOAT(x) ))}} 
  
// the pF is my ugly duckling :) -------------------------------------------
// float print is now standard in v3.4.42 (princ(float_integer)  but this is still a cuter print ...
/* {1} The go function for: mClaire/printFDigit(x:float,i:integer) [status=1] */
func F_printFDigit_float (x float64,i int) EID { 
    // eid body s = void
    var Result EID 
    if (x < 0) { 
      PRINC("-")
      Result = F_printFDigit_float((-x),i)
      } else {
      { var frac float64 
        var try_1 EID 
        { var arg_2 float64 
          var try_3 EID 
          { var arg_4 float64 
            var try_5 EID 
            { var arg_6 int 
              var try_7 EID 
              try_7 = F_integer_I_float((x+1e-10))
              if ErrorIn(try_7) {try_5 = try_7
              } else {
              arg_6 = INT(try_7)
              try_5 = EID{C__FLOAT,FVAL(F_to_float(arg_6))}
              }
              } 
            if ErrorIn(try_5) {try_3 = try_5
            } else {
            arg_4 = FLOAT(try_5)
            try_3 = EID{C__FLOAT,FVAL((x-arg_4))}
            }
            } 
          if ErrorIn(try_3) {try_1 = try_3
          } else {
          arg_2 = FLOAT(try_3)
          try_1 = EID{C__FLOAT,FVAL((arg_2+1e-10))}
          }
          } 
        if ErrorIn(try_1) {Result = try_1
        } else {
        frac = FLOAT(try_1)
        { var arg_8 int 
          var try_9 EID 
          try_9 = F_integer_I_float((x+1e-10))
          if ErrorIn(try_9) {Result = try_9
          } else {
          arg_8 = INT(try_9)
          F_princ_integer(arg_8)
          Result = EVOID
          }
          } 
        if !ErrorIn(Result) {
        PRINC(".")
        { var arg_10 int 
          var try_11 EID 
          try_11 = F_integer_I_float((frac*F__exp_float(10,F_to_float(i))))
          if ErrorIn(try_11) {Result = try_11
          } else {
          arg_10 = INT(try_11)
          Result = F_printFDigit_integer(arg_10,i)
          }
          } 
        if !ErrorIn(Result) {
        PRINC("")
        Result = EVOID
        }}
        }
        } 
      } 
    return Result} 
  
// The EID go function for: mClaire/printFDigit @ float (throw: true) 
func E_printFDigit_float (x EID,i EID) EID { 
    return F_printFDigit_float(FLOAT(x),INT(i) )} 
  
// print the first i digits of an integer
/* {1} The go function for: mClaire/printFDigit(x:integer,i:integer) [status=1] */
func F_printFDigit_integer (x int,i int) EID { 
    // eid body s = void
    var Result EID 
    if (i > 0) { 
      { var f int 
        var try_1 EID 
        try_1 = F__exp_integer(10,(i-1))
        if ErrorIn(try_1) {Result = try_1
        } else {
        f = INT(try_1)
        { var d int 
          var try_2 EID 
          try_2 = EID{C__INT,IVAL((x/f))}
          if ErrorIn(try_2) {Result = try_2
          } else {
          d = INT(try_2)
          F_princ_integer(d)
          if (i > 1) { 
            { var arg_3 int 
              var try_4 EID 
              try_4 = EID{C__INT,IVAL((x%f))}
              if ErrorIn(try_4) {Result = try_4
              } else {
              arg_3 = INT(try_4)
              Result = F_printFDigit_integer(arg_3,(i-1))
              }
              } 
            } else {
            Result = EID{CFALSE.Id(),0}
            } 
          }
          } 
        }
        } 
      } else {
      Result = EID{CFALSE.Id(),0}
      } 
    return Result} 
  
// The EID go function for: mClaire/printFDigit @ integer (throw: true) 
func E_printFDigit_integer (x EID,i EID) EID { 
    return F_printFDigit_integer(INT(x),INT(i) )} 
  
//--------- BAG --------------------------------------------------------
// in CLAIRE 4, bag is no longer a concrete data type, hence bag methods must be divided between sets and lists
// length(self:bag ) : integer -> length(self)  => becomes part of go
// in CLAIRE 4, we can tell the compiler that the list is a list<object> (optimization purpose)
/* {1} The go function for: mClaire/nth_object(self:list,n:integer) [status=0] */
func F_mClaire_nth_object_list (self *ClaireList ,n int) *ClaireAny  { 
    return  self.ValuesO()[n-1].Id()
    } 
  
// The EID go function for: mClaire/nth_object @ list (throw: false) 
func E_mClaire_nth_object_list (self EID,n EID) EID { 
    return F_mClaire_nth_object_list(ToList(OBJ(self)),INT(n) ).ToEID()} 
  
//
// nth_get(self:list,x:integer) : any -> nth_get(self, x)
// new in claire 4: tells the compiler that range check is required + EID optimized
/* {1} The go function for: nth_write(self:list,i:integer,v:any) [status=1] */
func F_nth_write_list (self *ClaireList ,i int,v *ClaireAny ) EID { 
    // eid body s = any
    var Result EID 
    if (self.Of().Contains(v) == CTRUE) { 
      Result = ToArray(self.Id()).NthPut(i,v).ToEID()
      } else {
      { var _CL_obj *ClaireSystemError   = ToSystemError(new(ClaireSystemError).Is(C_system_error))
        _CL_obj.Index = 17
        _CL_obj.Arg = v
        _CL_obj.Value = self.Of().Id()
        Result = _CL_obj.Close()
        } 
      } 
    return Result} 
  
// The EID go function for: nth_write @ list (throw: true) 
func E_nth_write_list (self EID,i EID,v EID) EID { 
    return F_nth_write_list(ToList(OBJ(self)),INT(i),ANY(v) )} 
  
// CLAIRE 4 duplication: define min/max for sets first
/* {1} The go function for: min(f:method,self:set) [status=1] */
func F_min_method2 (f *ClaireMethod ,self *ClaireSet ) EID { 
    // eid body s = any
    var Result EID 
    if (self.Size() != 0) { 
      { var x *ClaireAny   = CNULL
        { 
          var y *ClaireAny  
          _ = y
          Result= EID{CFALSE.Id(),0}
          var y_support *ClaireSet  
          y_support = self
          for i_it := 0; i_it < y_support.Count; i_it++ { 
            y = y_support.At(i_it)
            var loop_1 EID 
            _ = loop_1
            var g0112I *ClaireBoolean  
            var try_2 EID 
            { 
              var v_or6 *ClaireBoolean  
              
              v_or6 = Equal(x,CNULL)
              if (v_or6 == CTRUE) {try_2 = EID{CTRUE.Id(),0}
              } else { 
                var try_3 EID 
                try_3 = F_funcall_method2(f,y,x)
                if ErrorIn(try_3) {try_2 = try_3
                } else {
                v_or6 = ToBoolean(OBJ(try_3))
                if (v_or6 == CTRUE) {try_2 = EID{CTRUE.Id(),0}
                } else { 
                  try_2 = EID{CFALSE.Id(),0}} 
                } 
              }
              } 
            if ErrorIn(try_2) {loop_1 = try_2
            } else {
            g0112I = ToBoolean(OBJ(try_2))
            if (g0112I == CTRUE) { 
              x = y
              loop_1 = x.ToEID()
              } else {
              loop_1 = EID{CFALSE.Id(),0}
              } 
            }
            if ErrorIn(loop_1) {Result = loop_1
            break
            } else {
            }
            } 
          } 
        if !ErrorIn(Result) {
        Result = x.ToEID()
        }
        } 
      } else {
      Result = ToException(C_general_error.Make(MakeString("[183] min of empty set is undefined").Id(),CNIL.Id())).Close()
      } 
    return Result} 
  
// The EID go function for: min @ list<type_expression>(method, set) (throw: true) 
func E_min_method2 (f EID,self EID) EID { 
    return F_min_method2(ToMethod(OBJ(f)),ToSet(OBJ(self)) )} 
  
/* {1} The go function for: min_method2_type */
func F_min_method2_type (f *ClaireType ,self *ClaireType ) EID { 
    var Result EID 
    Result = EID{F_member_type(self).Id(),0}
    return Result} 
  
  
// The dual EID go function for: "min_method2_type" 
func E_min_method2_type (f EID,self EID) EID { 
    return F_min_method2_type(ToType(OBJ(f)),ToType(OBJ(self)))} 
  
/* {1} The go function for: max(f:method,self:set) [status=1] */
func F_max_method2 (f *ClaireMethod ,self *ClaireSet ) EID { 
    // eid body s = any
    var Result EID 
    if (self.Size() != 0) { 
      { var x *ClaireAny   = CNULL
        { 
          var y *ClaireAny  
          _ = y
          Result= EID{CFALSE.Id(),0}
          var y_support *ClaireSet  
          y_support = self
          for i_it := 0; i_it < y_support.Count; i_it++ { 
            y = y_support.At(i_it)
            var loop_1 EID 
            _ = loop_1
            var g0113I *ClaireBoolean  
            var try_2 EID 
            { 
              var v_or6 *ClaireBoolean  
              
              v_or6 = Equal(x,CNULL)
              if (v_or6 == CTRUE) {try_2 = EID{CTRUE.Id(),0}
              } else { 
                var try_3 EID 
                { var arg_4 *ClaireBoolean  
                  var try_5 EID 
                  { var arg_6 *ClaireAny  
                    var try_7 EID 
                    try_7 = F_funcall_method2(f,y,x)
                    if ErrorIn(try_7) {try_5 = try_7
                    } else {
                    arg_6 = ANY(try_7)
                    try_5 = EID{F_boolean_I_any(arg_6).Id(),0}
                    }
                    } 
                  if ErrorIn(try_5) {try_3 = try_5
                  } else {
                  arg_4 = ToBoolean(OBJ(try_5))
                  try_3 = EID{F__I_equal_any(arg_4.Id(),CTRUE.Id()).Id(),0}
                  }
                  } 
                if ErrorIn(try_3) {try_2 = try_3
                } else {
                v_or6 = ToBoolean(OBJ(try_3))
                if (v_or6 == CTRUE) {try_2 = EID{CTRUE.Id(),0}
                } else { 
                  try_2 = EID{CFALSE.Id(),0}} 
                } 
              }
              } 
            if ErrorIn(try_2) {loop_1 = try_2
            } else {
            g0113I = ToBoolean(OBJ(try_2))
            if (g0113I == CTRUE) { 
              x = y
              loop_1 = x.ToEID()
              } else {
              loop_1 = EID{CFALSE.Id(),0}
              } 
            }
            if ErrorIn(loop_1) {Result = loop_1
            break
            } else {
            }
            } 
          } 
        if !ErrorIn(Result) {
        Result = x.ToEID()
        }
        } 
      } else {
      Result = ToException(C_general_error.Make(MakeString("[183] max of empty set is undefined").Id(),CNIL.Id())).Close()
      } 
    return Result} 
  
// The EID go function for: max @ list<type_expression>(method, set) (throw: true) 
func E_max_method2 (f EID,self EID) EID { 
    return F_max_method2(ToMethod(OBJ(f)),ToSet(OBJ(self)) )} 
  
/* {1} The go function for: max_method2_type */
func F_max_method2_type (f *ClaireType ,self *ClaireType ) EID { 
    var Result EID 
    Result = EID{F_member_type(self).Id(),0}
    return Result} 
  
  
// The dual EID go function for: "max_method2_type" 
func E_max_method2_type (f EID,self EID) EID { 
    return F_max_method2_type(ToType(OBJ(f)),ToType(OBJ(self)))} 
  
// CLAIRE 4 : optimize for lists
/* {1} The go function for: min(f:method,self:list) [status=1] */
func F_min_method3 (f *ClaireMethod ,self *ClaireList ) EID { 
    // eid body s = any
    var Result EID 
    { var n int  = self.Length()
      if (n != 0) { 
        { var x *ClaireAny   = self.At(1-1)
          { var i int  = 2
            { var g0114 int  = n
              Result= EID{CFALSE.Id(),0}
              for (i <= g0114) { 
                var loop_1 EID 
                _ = loop_1
                { 
                var g0115I *ClaireBoolean  
                var try_2 EID 
                try_2 = F_funcall_method2(f,self.At(i-1),x)
                if ErrorIn(try_2) {loop_1 = try_2
                } else {
                g0115I = ToBoolean(OBJ(try_2))
                if (g0115I == CTRUE) { 
                  x = self.At(i-1)
                  loop_1 = x.ToEID()
                  } else {
                  loop_1 = EID{CFALSE.Id(),0}
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
          Result = x.ToEID()
          }
          } 
        } else {
        Result = ToException(C_general_error.Make(MakeString("[183] min of empty list is undefined").Id(),CNIL.Id())).Close()
        } 
      } 
    return Result} 
  
// The EID go function for: min @ list<type_expression>(method, list) (throw: true) 
func E_min_method3 (f EID,self EID) EID { 
    return F_min_method3(ToMethod(OBJ(f)),ToList(OBJ(self)) )} 
  
/* {1} The go function for: min_method3_type */
func F_min_method3_type (f *ClaireType ,self *ClaireType ) EID { 
    var Result EID 
    Result = EID{F_member_type(self).Id(),0}
    return Result} 
  
  
// The dual EID go function for: "min_method3_type" 
func E_min_method3_type (f EID,self EID) EID { 
    return F_min_method3_type(ToType(OBJ(f)),ToType(OBJ(self)))} 
  
/* {1} The go function for: max(f:method,self:list) [status=1] */
func F_max_method3 (f *ClaireMethod ,self *ClaireList ) EID { 
    // eid body s = any
    var Result EID 
    { var n int  = self.Length()
      if (n != 0) { 
        { var x *ClaireAny   = self.At(1-1)
          { var i int  = 2
            { var g0116 int  = n
              Result= EID{CFALSE.Id(),0}
              for (i <= g0116) { 
                var loop_1 EID 
                _ = loop_1
                { 
                var g0117I *ClaireBoolean  
                var try_2 EID 
                { var arg_3 *ClaireBoolean  
                  var try_4 EID 
                  { var arg_5 *ClaireAny  
                    var try_6 EID 
                    try_6 = F_funcall_method2(f,self.At(i-1),x)
                    if ErrorIn(try_6) {try_4 = try_6
                    } else {
                    arg_5 = ANY(try_6)
                    try_4 = EID{F_boolean_I_any(arg_5).Id(),0}
                    }
                    } 
                  if ErrorIn(try_4) {try_2 = try_4
                  } else {
                  arg_3 = ToBoolean(OBJ(try_4))
                  try_2 = EID{F__I_equal_any(arg_3.Id(),CTRUE.Id()).Id(),0}
                  }
                  } 
                if ErrorIn(try_2) {loop_1 = try_2
                } else {
                g0117I = ToBoolean(OBJ(try_2))
                if (g0117I == CTRUE) { 
                  x = self.At(i-1)
                  loop_1 = x.ToEID()
                  } else {
                  loop_1 = EID{CFALSE.Id(),0}
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
          Result = x.ToEID()
          }
          } 
        } else {
        Result = ToException(C_general_error.Make(MakeString("[183] max of empty list is undefined").Id(),CNIL.Id())).Close()
        } 
      } 
    return Result} 
  
// The EID go function for: max @ list<type_expression>(method, list) (throw: true) 
func E_max_method3 (f EID,self EID) EID { 
    return F_max_method3(ToMethod(OBJ(f)),ToList(OBJ(self)) )} 
  
/* {1} The go function for: max_method3_type */
func F_max_method3_type (f *ClaireType ,self *ClaireType ) EID { 
    var Result EID 
    Result = EID{F_member_type(self).Id(),0}
    return Result} 
  
  
// The dual EID go function for: "max_method3_type" 
func E_max_method3_type (f EID,self EID) EID { 
    return F_max_method3_type(ToType(OBJ(f)),ToType(OBJ(self)))} 
  
// CLAIRE4 : /+ is native for list
// new for claire 3.4
/* {1} The go function for: random(self:list) [status=0] */
func F_random_list (self *ClaireList ) *ClaireAny  { 
    // procedure body with s = any
    var Result *ClaireAny  
    { var n int  = self.Length()
      Result = self.At((1+F_random_integer(n))-1)
      } 
    return Result} 
  
// The EID go function for: random @ list (throw: false) 
func E_random_list (self EID) EID { 
    return F_random_list(ToList(OBJ(self)) ).ToEID()} 
  
/* {1} The go function for: random_list_type */
func F_random_list_type (self *ClaireType ) EID { 
    var Result EID 
    Result = EID{F_member_type(self).Id(),0}
    return Result} 
  
  
// The dual EID go function for: "random_list_type" 
func E_random_list_type (self EID) EID { 
    return F_random_list_type(ToType(OBJ(self)))} 
  
//--------- LIST --------------------------------------------------------
// last element of a list
/* {1} The go function for: last(self:list) [status=1] */
func F_last_list (self *ClaireList ) EID { 
    // eid body s = any
    var Result EID 
    if (self.Length() > 0) { 
      Result = self.At(self.Length()-1).ToEID()
      } else {
      Result = ToException(C_general_error.Make(MakeString("[41] car(nil) is undefined").Id(),CNIL.Id())).Close()
      } 
    return Result} 
  
// The EID go function for: last @ list (throw: true) 
func E_last_list (self EID) EID { 
    return F_last_list(ToList(OBJ(self)) )} 
  
/* {1} The go function for: last_list_type */
func F_last_list_type (self *ClaireType ) EID { 
    var Result EID 
    Result = EID{F_member_type(self).Id(),0}
    return Result} 
  
  
// The dual EID go function for: "last_list_type" 
func E_last_list_type (self EID) EID { 
    return F_last_list_type(ToType(OBJ(self)))} 
  
// remove the last element
/* {1} The go function for: rmlast(self:list) [status=1] */
func F_rmlast_list (self *ClaireList ) EID { 
    // eid body s = list
    var Result EID 
    Result = self.Nth_dash(self.Length())
    if !ErrorIn(Result) {
    Result = EID{self.Id(),0}
    }
    return Result} 
  
// The EID go function for: rmlast @ list (throw: true) 
func E_rmlast_list (self EID) EID { 
    return F_rmlast_list(ToList(OBJ(self)) )} 
  
// the old LISP method
/* {1} The go function for: car(self:list) [status=1] */
func F_car_list (self *ClaireList ) EID { 
    // eid body s = any
    var Result EID 
    if (self.Length() > 0) { 
      Result = self.At(1-1).ToEID()
      } else {
      Result = ToException(C_general_error.Make(MakeString("[41] car(nil) is undefined").Id(),CNIL.Id())).Close()
      } 
    return Result} 
  
// The EID go function for: car @ list (throw: true) 
func E_car_list (self EID) EID { 
    return F_car_list(ToList(OBJ(self)) )} 
  
// this method sorts a list according to an order
/* {1} The go function for: sort(f:method,self:list) [status=1] */
func F_sort_method (f *ClaireMethod ,self *ClaireList ) EID { 
    // eid body s = list
    var Result EID 
    Result = F_quicksort_list(self,f,1,self.Length())
    if !ErrorIn(Result) {
    Result = EID{self.Id(),0}
    }
    return Result} 
  
// The EID go function for: sort @ method (throw: true) 
func E_sort_method (f EID,self EID) EID { 
    return F_sort_method(ToMethod(OBJ(f)),ToList(OBJ(self)) )} 
  
// v3.0.38: upgrade the quicksort algorithm with a better pivot selection cf.bag.cpp
// this is also proposed as a macro: cf. file.cl
/* {1} The go function for: quicksort(self:list,f:method,n:integer,m:integer) [status=1] */
func F_quicksort_list (self *ClaireList ,f *ClaireMethod ,n int,m int) EID { 
    // eid body s = void
    var Result EID 
    if (m > n) { 
      { var x *ClaireAny   = self.At(n-1)
        if (m == (n+1)) { 
          var g0119I *ClaireBoolean  
          var try_1 EID 
          try_1 = F_funcall_method2(f,self.At(m-1),x)
          if ErrorIn(try_1) {Result = try_1
          } else {
          g0119I = ToBoolean(OBJ(try_1))
          if (g0119I == CTRUE) { 
            ToArray(self.Id()).NthPut(n,self.At(m-1))
            Result = ToArray(self.Id()).NthPut(m,x).ToEID()
            } else {
            Result = EID{CFALSE.Id(),0}
            } 
          }
          } else {
          { var p int  = ((m+n)>>1)
            { var q int  = n
              x = self.At(p-1)
              if (p != n) { 
                ToArray(self.Id()).NthPut(p,self.At(n-1))
                } 
              { var p int  = (n+1)
                { var g0118 int  = m
                  Result= EID{CFALSE.Id(),0}
                  for (p <= g0118) { 
                    var loop_2 EID 
                    _ = loop_2
                    { 
                    var g0120I *ClaireBoolean  
                    var try_3 EID 
                    try_3 = F_funcall_method2(f,self.At(p-1),x)
                    if ErrorIn(try_3) {loop_2 = try_3
                    } else {
                    g0120I = ToBoolean(OBJ(try_3))
                    if (g0120I == CTRUE) { 
                      ToArray(self.Id()).NthPut(n,self.At(p-1))
                      n = (n+1)
                      if (p > n) { 
                        loop_2 = ToArray(self.Id()).NthPut(p,self.At(n-1)).ToEID()
                        } else {
                        loop_2 = EID{CFALSE.Id(),0}
                        } 
                      } else {
                      loop_2 = EID{CFALSE.Id(),0}
                      } 
                    }
                    if ErrorIn(loop_2) {Result = loop_2
                    break
                    } else {
                    p = (p+1)
                    }
                    } 
                  }
                  } 
                } 
              if !ErrorIn(Result) {
              ToArray(self.Id()).NthPut(n,x)
              Result = F_quicksort_list(self,f,q,(n-1))
              if !ErrorIn(Result) {
              Result = F_quicksort_list(self,f,(n+1),m)
              }}
              } 
            } 
          } 
        } 
      } else {
      Result = EID{CFALSE.Id(),0}
      } 
    return Result} 
  
// The EID go function for: quicksort @ list (throw: true) 
func E_quicksort_list (self EID,f EID,n EID,m EID) EID { 
    return F_quicksort_list(ToList(OBJ(self)),
      ToMethod(OBJ(f)),
      INT(n),
      INT(m) )} 
  
// destructive method that build the powerset
/* {1} The go function for: build_powerset(self:list) [status=0] */
func F_build_powerset_list (self *ClaireList ) *ClaireSet  { 
    // procedure body with s = set
    var Result *ClaireSet  
    if (self.Length() != 0) { 
      { var x *ClaireAny   = self.At(1-1)
        { var l1 *ClaireSet   = F_build_powerset_list(self.Skip(1))
          { var l2 *ClaireSet   = l1
            { 
              var y *ClaireSet  
              _ = y
              var y_iter *ClaireAny  
              var y_support *ClaireSet  
              y_support = l1
              for i_it := 0; i_it < y_support.Count; i_it++ { 
                y_iter = y_support.At(i_it)
                y = ToSet(y_iter)
                l2 = l2.AddFast(F_append_set(MakeConstantSet(x),y).Id())
                } 
              } 
            Result = l2
            } 
          } 
        } 
      } else {
      Result = MakeConstantSet(CEMPTY.Id())
      } 
    return Result} 
  
// The EID go function for: build_powerset @ list (throw: false) 
func E_build_powerset_list (self EID) EID { 
    return EID{F_build_powerset_list(ToList(OBJ(self)) ).Id(),0}} 
  
// skip 
// new and useful (v3.1.06) - create a list with n replication of the default value - deprecated
/* {1} The go function for: make_copy_list(n:integer,d:any) [status=0] */
func F_make_copy_list_integer (n int,d *ClaireAny ) *ClaireList  { 
    // procedure body with s = list
    var Result *ClaireList  
    { var l *ClaireList   = F_make_list_integer(n,d)
      if (d.Isa.IsIn(C_list) == CTRUE) { 
        { var g0121 *ClaireList   = ToList(d)
          { var i int  = 1
            { var g0122 int  = n
              for (i <= g0122) { 
                ToArray(l.Id()).NthPut(i,g0121.Copy().Id())
                i = (i+1)
                } 
              } 
            } 
          } 
        } 
      Result = l
      } 
    return Result} 
  
// The EID go function for: make_copy_list @ integer (throw: false) 
func E_make_copy_list_integer (n EID,d EID) EID { 
    return EID{F_make_copy_list_integer(INT(n),ANY(d) ).Id(),0}} 
  
// new version : create a typed list for integer or floats
/* {1} The go function for: typed_copy_list(t:type,n:integer,d:any) [status=0] */
func F_typed_copy_list_type (t *ClaireType ,n int,d *ClaireAny ) *ClaireList  { 
    // procedure body with s = list
    var Result *ClaireList  
    if (t.Included(ToType(C_integer.Id())) == CTRUE) { 
      Result = F_make_list_integer2(n,ToType(C_integer.Id()),d)
      }  else if (t.Included(ToType(C_float.Id())) == CTRUE) { 
      Result = F_make_list_integer2(n,ToType(C_float.Id()),d)
      } else {
      { var l *ClaireList   = F_make_list_integer(n,d)
        if (d.Isa.IsIn(C_bag) == CTRUE) { 
          { var g0123 *ClaireBag   = ToBag(d)
            { var i int  = 1
              { var g0124 int  = n
                for (i <= g0124) { 
                  ToArray(l.Id()).NthPut(i,ANY(F_CALL(C_copy,ARGS(EID{g0123.Id(),0}))))
                  i = (i+1)
                  } 
                } 
              } 
            } 
          } 
        Result = l
        } 
      } 
    return Result} 
  
// The EID go function for: typed_copy_list @ type (throw: false) 
func E_typed_copy_list_type (t EID,n EID,d EID) EID { 
    return EID{F_typed_copy_list_type(ToType(OBJ(t)),INT(n),ANY(d) ).Id(),0}} 
  
//----------------------  SET  ---------------------------------------------
/* {1} The go function for: difference(self:set,x:set) [status=0] */
func F_difference_set (self *ClaireSet ,x *ClaireSet ) *ClaireSet  { 
    // procedure body with s = set
    var Result *ClaireSet  
    { var y_in *ClaireSet   = self
      { var y_out *ClaireSet   = y_in.Empty()
        { 
          var y *ClaireAny  
          _ = y
          var y_support *ClaireSet  
          y_support = y_in
          for i_it := 0; i_it < y_support.Count; i_it++ { 
            y = y_support.At(i_it)
            if (x.Contain_ask(y) != CTRUE) { 
              y_out.AddFast(y)
              } 
            } 
          } 
        Result = y_out
        } 
      } 
    return Result} 
  
// The EID go function for: difference @ set (throw: false) 
func E_difference_set (self EID,x EID) EID { 
    return EID{F_difference_set(ToSet(OBJ(self)),ToSet(OBJ(x)) ).Id(),0}} 
  
//----------------------  TYPE ---------------------------------------------
//--------- ARRAY --------------------------------------------------------
/* {1} The go function for: nth=(self:array,x:integer,y:any) [status=1] */
func F_nth_equal_array (self *ClaireList ,x int,y *ClaireAny ) EID { 
    // eid body s = void
    var Result EID 
    if (ToList(self.Id()).Of().Contains(y) != CTRUE) { 
      Result = ToException(C_general_error.Make(MakeString("type mismatch for array update ~S, ~S").Id(),MakeConstantList(y,self.Id()).Id())).Close()
      }  else if ((x > 0) && 
        (x <= self.Length())) { 
      Result = self.NthPut(x,y).ToEID()
      } else {
      Result = ToException(C_general_error.Make(MakeString("nth[~S] out of scope for ~S").Id(),MakeConstantList(MakeInteger(x).Id(),self.Id()).Id())).Close()
      } 
    return Result} 
  
// The EID go function for: nth= @ array (throw: true) 
func E_nth_equal_array (self EID,x EID,y EID) EID { 
    return F_nth_equal_array(ToArray(OBJ(self)),INT(x),ANY(y) )} 
  
/* {1} The go function for: self_print(self:array) [status=1] */
func F_self_print_array_Core (self *ClaireList ) EID { 
    // eid body s = void
    var Result EID 
    PRINC("array<")
    Result = F_print_any(ToList(self.Id()).Of().Id())
    if !ErrorIn(Result) {
    PRINC(">[")
    F_princ_integer(self.Length())
    PRINC("]")
    Result = EVOID
    }
    return Result} 
  
// The EID go function for: self_print @ array (throw: true) 
func E_self_print_array_Core (self EID) EID { 
    return F_self_print_array_Core(ToArray(OBJ(self)) )} 
  
//---------------------- CHAR --------------------------------------------
/* {1} The go function for: self_print(self:char) [status=0] */
func F_self_print_char_Core (self rune)  { 
    // procedure body with s = void
    PRINC("'")
    F_princ_char(self)
    PRINC("'")
    } 
  
// The EID go function for: self_print @ char (throw: false) 
func E_self_print_char_Core (self EID) EID { 
    F_self_print_char_Core(CHAR(self) )
    return EVOID} 
  
/* {1} The go function for: <=(c1:char,c2:char) [status=0] */
func F__inf_equal_char (c1 rune,c2 rune) *ClaireBoolean  { 
    if (int(c1) <= int(c2)) {return CTRUE
    } else {return CFALSE}} 
  
// The EID go function for: <= @ char (throw: false) 
func E__inf_equal_char (c1 EID,c2 EID) EID { 
    return EID{F__inf_equal_char(CHAR(c1),CHAR(c2) ).Id(),0}} 
  
// --------------------- BOOL -----------------------------------------------
/* {1} The go function for: random(b:boolean) [status=0] */
func F_random_boolean (b *ClaireBoolean ) *ClaireBoolean  { 
    if (b == CTRUE) { 
      if (F_random_integer(10000) >= 5000) {return CTRUE
      } else {return CFALSE}} else {
      if (CFALSE == CTRUE) {return CTRUE
      } else {return CFALSE}} 
    } 
  
// The EID go function for: random @ boolean (throw: false) 
func E_random_boolean (b EID) EID { 
    return EID{F_random_boolean(ToBoolean(OBJ(b)) ).Id(),0}} 
  