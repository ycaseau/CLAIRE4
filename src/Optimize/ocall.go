/***** CLAIRE Compilation of file /Users/ycaseau/Dropbox/src/clairev4.12/src/compile/ocall.cl 
         [version 4.1.6 / safety 5] Saturday 08-09-2025 06:51:14 *****/

package Optimize
import (_ "fmt"
	. "Kernel"
	"Core"
	"Language"
	"Reader"
)

//-------- dumb function to prevent import errors --------
func import_g0167() { 
_ = Core.It
_ = Language.It
_ = Reader.It
} 


//+-------------------------------------------------------------+
//| CLAIRE                                                      |
//| ocall.cl                                                    |
//| Copyright (C) 1994 - 2025 Yves Caseau. All Rights Reserved  |
//| cf. copyright info in file object.cl: about()               |
//+-------------------------------------------------------------+
// ------------------------------------------------------------------
// this is the heart of the CLAIRE optimizer : message to function calls
// ------------------------------------------------------------------
// ******************************************************************
// *  Table of contents                                             *
// *    Part 1: Restruction Binding                                 *
// *    Part 2: Generic c_type & c_code                             *
// *    Part 3: specialized c_code                                  *
// *    Part 4: Method optimizing                                   *
// *    Part 5: inline methods                                      *
// ******************************************************************
// ******************************************************************
// *    Part 1: Restruction Binding                                 *
// ******************************************************************
// if mode = true This method finds the unique property that can be used, if any;
// returns () if no restriction exist, and "ambiguous" if it cannot
// answer.
// if mode = false, we return the union of the matching ranges
/* The go function for: restriction!(self:property,l:list,mode:boolean) [status=0] */
func F_Optimize_restriction_I_property (self *ClaireProperty,l *ClaireList,mode *ClaireBoolean) *ClaireAny { 
{ var i int = 1
  { var g0168 int = l.Length()
    for (i <= g0168) { 
      ToArray(l.Id()).NthPut(i,F_Optimize_ptype_type(ToType(l.At(i-1))).Id())
      i = (i+1)
      } 
    } 
  } 
return  F_Optimize_restriction_I_list(self.Definition,l,mode)
} 

// The EID go function for: restriction! @ property (throw: false) 
func E_Optimize_restriction_I_property (self EID,l EID,mode EID) EID { 
return F_Optimize_restriction_I_property(ToProperty(OBJ(self)),ToList(OBJ(l)),ToBoolean(OBJ(mode)) ).ToEID()} 

// new: extended match for listargs (last member of l2) - note that we test this precondition :)
/* The go function for: dmatch?(l:list,l2:list) [status=0] */
func F_dmatch_ask_list (l *ClaireList,l2 *ClaireList) *ClaireBoolean { 
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
          { var g0169 int = x
            arg_1= CFALSE.Id()
            for (i <= g0169) { 
              if ((i == x) && 
                  (l2.At(i-1) == C_listargs.Id())) { 
                arg_1 = CFALSE.Id()
                break
                }  else if (Core.F_tmatch_ask_any(l.At(i-1),l2.At(i-1),l) != CTRUE) { 
                
                arg_1 = CTRUE.Id()
                break
                } 
              i = (i+1)
              } 
            } 
          } 
        Result = Core.F_not_any(arg_1)
        } 
      } 
    } 
  } 
return Result} 

// The EID go function for: dmatch? @ list (throw: false) 
func E_dmatch_ask_list (l EID,l2 EID) EID { 
return EID{F_dmatch_ask_list(ToList(OBJ(l)),ToList(OBJ(l2)) ).Id(),0}} 

// finds a suitable restriction in lr. Returns a restriction for a match,
// list(r) for a possible match (unique), () for no match and ambiguous
// otherwise
// CLAIRE4 : we define "open required" based on the property.
/* The go function for: restriction!(lr:list,l:list,mode:boolean) [status=0] */
func F_Optimize_restriction_I_list (lr *ClaireList,l *ClaireList,mode *ClaireBoolean) *ClaireAny { 
var Result *ClaireAny

{ var open_required *ClaireBoolean = MakeBoolean((lr.Length() > 0) && ANY(Core.F_CALL(C_selector,ARGS(lr.At(0).ToEID()))).IsInt(3))
  { var rep *ClaireAny = CEMPTY.Id()
    { 
      var r *ClaireRestriction
      _ = r
      var r_iter *ClaireAny
      var r_support *ClaireList
      r_support = lr
      r_len := r_support.Length()
      for i_it := 0; i_it < r_len; i_it++ { 
        r_iter = r_support.At(i_it)
        r = ToRestriction(r_iter)
        if (r.Domain.ValuesO()[r.Domain.Length()-1] == C_listargs.Id()) { 
          if (F_dmatch_ask_list(l,r.Domain) == CTRUE) { 
            if (F_boolean_I_any(rep).Id() != CTRUE.Id()) { 
              if (mode == CTRUE) { 
                rep = r.Id()
                } else {
                rep = r.Range.Id()
                } 
              
              break
              }  else if (mode == CTRUE) { 
              rep = C_Optimize_ambiguous.Id()
              
              break
              } else {
              rep = Core.F_U_type(ToType(rep),r.Range).Id()
              } 
            } 
          }  else if ((F_boolean_I_any(rep).Id() != CTRUE.Id()) && 
            (Core.F_tmatch_ask_list(l,r.Domain) == CTRUE)) { 
          if (mode == CTRUE) { 
            rep = r.Id()
            } else {
            rep = r.Range.Id()
            } 
          
          break
          }  else if (Core.F__exp_list(r.Domain,l).Length() != 0) { 
          if (mode != CTRUE) { 
            
            rep = Core.F_U_type(ToType(rep),r.Range).Id()
            }  else if ((C_compiler.Safety <= 1) || 
              ((Equal(rep,CEMPTY.Id()) != CTRUE) || 
                (open_required == CTRUE))) { 
            rep = C_Optimize_ambiguous.Id()
            
            break
            } else {
            rep = r.Id()
            } 
          } 
        } 
      } 
    
    Result = rep
    } 
  } 
return Result} 

// The EID go function for: restriction! @ list (throw: false) 
func E_Optimize_restriction_I_list (lr EID,l EID,mode EID) EID { 
return F_Optimize_restriction_I_list(ToList(OBJ(lr)),ToList(OBJ(l)),ToBoolean(OBJ(mode)) ).ToEID()} 

// we need a debug mode : shows the tmatch of all restrictions
/* The go function for: findr(p:property,l:list) [status=1] */
func F_findr_property (p *ClaireProperty,l *ClaireList) EID { 
var Result EID
{ var lr *ClaireList = p.Definition
  { 
    var r *ClaireRestriction
    _ = r
    var r_iter *ClaireAny
    Result= EID{CFALSE.Id(),0}
    for _,r_iter = range(lr.ValuesO()){ 
      r = ToRestriction(r_iter)
      var loop_1 EID
      _ = loop_1
      { 
      PRINC("tmatch(")
      loop_1 = Core.F_print_any(r.Id())
      if ErrorIn(loop_1) {Result = loop_1
      break
      } else {
      PRINC(") with ")
      loop_1 = Core.F_print_any(l.Id())
      if ErrorIn(loop_1) {Result = loop_1
      break
      } else {
      PRINC(" -> ")
      loop_1 = Core.F_print_any(Core.F_tmatch_ask_list(l,r.Domain).Id())
      if ErrorIn(loop_1) {Result = loop_1
      break
      } else {
      PRINC(", intersection:")
      loop_1 = Core.F_print_any(Core.F__exp_list(r.Domain,l).Id())
      if ErrorIn(loop_1) {Result = loop_1
      break
      } else {
      PRINC("\n")
      }}}}
      }
      } 
    } 
  } 
return Result} 

// The EID go function for: findr @ property (throw: true) 
func E_findr_property (p EID,l EID) EID { 
return F_findr_property(ToProperty(OBJ(p)),ToList(OBJ(l)) )} 

// special version for Super, which only looks at methods with domains
// bigger than c
/* The go function for: restriction!(c:class,lr:list,l:list) [status=0] */
func F_Optimize_restriction_I_class (c *ClaireClass,lr *ClaireList,l *ClaireList) *ClaireAny { 
var Result *ClaireAny
if (C_compiler.Safety >= 2) { 
  ToArray(l.Id()).NthPut(1,Core.F__exp_type(ToType(c.Id()),ToType(l.At(0))).Id())
  } 
{ 
  var r *ClaireRestriction
  _ = r
  var r_iter *ClaireAny
  Result= CFALSE.Id()
  var r_support *ClaireList
  r_support = lr
  r_len := r_support.Length()
  for i_it := 0; i_it < r_len; i_it++ { 
    r_iter = r_support.At(i_it)
    r = ToRestriction(r_iter)
    if (ToType(c.Id()).Included(ToType(r.Domain.ValuesO()[0])) == CTRUE) { 
      if (Core.F_tmatch_ask_list(l,r.Domain) == CTRUE) { 
        Result = r.Id()
        break
        }  else if (Core.F__exp_list(r.Domain,l).Length() != 0) { 
        Result = C_Optimize_ambiguous.Id()
        break
        } 
      } 
    } 
  } 
return Result} 

// The EID go function for: restriction! @ class (throw: false) 
func E_Optimize_restriction_I_class (c EID,lr EID,l EID) EID { 
return F_Optimize_restriction_I_class(ToClass(OBJ(c)),ToList(OBJ(lr)),ToList(OBJ(l)) ).ToEID()} 

// uses a second order type {property + function}
/* The go function for: use_range(self:method,%l:list) [status=1] */
func F_Optimize_use_range_method (self *ClaireMethod,_Zl *ClaireList) EID { 
var Result EID
if ((self.Inline_ask == CTRUE) && 
    (self.Typing == CNULL)) { 
  { var lv *ClaireList = self.Formula.Vars
    { var _Zt *ClaireType = ToType(C_any.Id())
      { var _Zl2 *ClaireList
        { 
          var v_list4 *ClaireList
          var v *ClaireAny
          var v_local4 *ClaireAny
          v_list4 = lv
          _Zl2 = CreateList(ToType(CEMPTY.Id()),v_list4.Length())
          for CLcount := 0; CLcount < v_list4.Length(); CLcount++{ 
            v = v_list4.At(CLcount)
            v_local4 = ANY(Core.F_CALL(C_range,ARGS(v.ToEID())))
            _Zl2.PutAt(CLcount,v_local4)
            } 
          } 
        { 
          var v *ClaireAny
          _ = v
          Result= EID{CFALSE.Id(),0}
          var v_support *ClaireList
          v_support = lv
          v_len := v_support.Length()
          for i_it := 0; i_it < v_len; i_it++ { 
            v = v_support.At(i_it)
            var loop_1 EID
            _ = loop_1
            loop_1 = Core.F_put_property2(C_range,ToObject(v),_Zl.At((INT(Core.F_CALL(C_mClaire_index,ARGS(v.ToEID())))+1)-1))
            if ErrorIn(loop_1) {Result = loop_1
            break
            } else {
            }
            } 
          } 
        if !ErrorIn(Result) {
        var try_2 EID
        try_2 = Core.F_CALL(C_c_type,ARGS(self.Formula.Body.ToEID()))
        if ErrorIn(try_2) {Result = try_2
        } else {
        _Zt = ToType(OBJ(try_2))
        Result = EID{_Zt.Id(),0}
        { 
          var v *ClaireAny
          _ = v
          Result= EID{CFALSE.Id(),0}
          var v_support *ClaireList
          v_support = lv
          v_len := v_support.Length()
          for i_it := 0; i_it < v_len; i_it++ { 
            v = v_support.At(i_it)
            var loop_3 EID
            _ = loop_3
            loop_3 = Core.F_put_property2(C_range,ToObject(v),_Zl2.At((INT(Core.F_CALL(C_mClaire_index,ARGS(v.ToEID())))+1)-1))
            if ErrorIn(loop_3) {Result = loop_3
            break
            } else {
            }
            } 
          } 
        if !ErrorIn(Result) {
        if (self.Range.Isa.IsIn(C_type) == CTRUE) { 
          var try_4 EID
          try_4 = Core.F_CALL(ToProperty(C__exp.Id()),ARGS(EID{_Zt.Id(),0},EID{self.Range.Id(),0}))
          if ErrorIn(try_4) {Result = try_4
          } else {
          _Zt = ToType(OBJ(try_4))
          Result = EID{_Zt.Id(),0}
          }
          } else {
          Result = EID{CFALSE.Id(),0}
          } 
        if !ErrorIn(Result) {
        if (F_boolean_I_any(_Zt.Id()).Id() != CTRUE.Id()) { 
          { var arg_5 *ClaireAny
            var try_6 EID
            { 
              var v_bag_arg *ClaireAny
              try_6= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
              ToList(OBJ(try_6)).AddFast(self.Id())
              ToList(OBJ(try_6)).AddFast(self.Range.Id())
              var try_7 EID
              try_7 = Core.F_CALL(C_c_type,ARGS(self.Formula.Body.ToEID()))
              if ErrorIn(try_7) {try_6 = try_7
              } else {
              v_bag_arg = ANY(try_7)
              ToList(OBJ(try_6)).AddFast(v_bag_arg)}
              } 
            if ErrorIn(try_6) {Result = try_6
            } else {
            arg_5 = ANY(try_6)
            Result = F_Compile_Cerror_string(MakeString("[207] inline ~S: range ~S is incompatible with ~S (inferred)"),ToList(arg_5))
            }
            } 
          } else {
          Result = EID{CFALSE.Id(),0}
          } 
        if !ErrorIn(Result) {
        Result = EID{_Zt.Id(),0}
        }}}}}
        } 
      } 
    } 
  } else {
  
  { var f *ClaireAny = self.Typing
    { var _Zl2 *ClaireList
      { 
        var v_list3 *ClaireList
        var u *ClaireAny
        var v_local3 *ClaireAny
        v_list3 = _Zl
        _Zl2 = CreateList(ToType(C_type.Id()),v_list3.Length())
        for CLcount := 0; CLcount < v_list3.Length(); CLcount++{ 
          u = v_list3.At(CLcount)
          v_local3 = F_Optimize_ptype_type(ToType(u)).Id()
          _Zl2.PutAt(CLcount,v_local3)
          } 
        } 
      { var _Zt1 *ClaireType = self.Range
        { var _Zt2 *ClaireType
          { 
            var _Zt2_H EID
            h_index := ClEnv.Index
            h_base := ClEnv.Base
            if (f.Isa.IsIn(C_lambda) == CTRUE) { 
              { var g0170 *ClaireLambda = ToLambda(f)
                _Zt2_H = Language.F_apply_lambda(g0170,_Zl2)
                } 
              }  else if (f.Isa.IsIn(C_property) == CTRUE) { 
              { var g0171 *ClaireProperty = ToProperty(f)
                _Zt2_H = Core.F_apply_property(g0171,_Zl2)
                } 
              }  else if (C_function.Id() == f.Isa.Id()) { 
              { var g0172 *ClaireFunction = ToFunction(f)
                _Zt2_H = F_apply_function(g0172,_Zl2)
                } 
              } else {
              _Zt2_H = EID{_Zt1.Id(),0}
              } 
            if ErrorIn(_Zt2_H){ 
              ClEnv.Index = h_index
              ClEnv.Base = h_base
              F_Compile_warn_void()
              Core.F_tformat_string(MakeString(" ~S's 2nd-order type failed on ~S\n"),1,MakeConstantList(self.Id(),_Zl.Id()))
              Reader.F_print_exception_void()
              _Zt2 = _Zt1
              } else {
              _Zt2 = ToType(OBJ(_Zt2_H))
              } 
            } 
          Result = EID{_Zt2.Id(),0}
          } 
        } 
      } 
    } 
  } 
return Result} 

// The EID go function for: use_range @ method (throw: true) 
func E_Optimize_use_range_method (self EID,_Zl EID) EID { 
return F_Optimize_use_range_method(ToMethod(OBJ(self)),ToList(OBJ(_Zl)) )} 

//      (if (sort=(osort(%t1), osort(%t2)) | self.selector = externC) %t2
//       else if sort=(any, osort(%t1))
//          Union(Kernel/t1 = any, Kernel/t2 = %t2)   // forces the sort and preserves the type
//       else %t1)) ]
// ******************************************************************
// *    Part 2: Generic c_type & c_code                             *
// ******************************************************************
// this is the optimizer for messages
// It follows the structure of the evaluator (self_eval)
// optimize is the distributed compiling method equivalent to the
// evaluation "behave" method
/* The go function for: c_type(self:Call) [status=1] */
func F_c_type_Call (self *Language.Call) EID { 
var Result EID
if (self.Selector.Id() == Language.C_function_I.Id()) { 
  Result = EID{C_function.Id(),0}
  } else {
  { var s *ClaireProperty = self.Selector
    { var l *ClaireList = self.Args
      { var _Ztype *ClaireList
        var try_1 EID
        { 
          var v_list4 *ClaireList
          var x *ClaireAny
          var v_local4 *ClaireAny
          v_list4 = l
          try_1 = EID{CreateList(ToType(CEMPTY.Id()),v_list4.Length()).Id(),0}
          for CLcount := 0; CLcount < v_list4.Length(); CLcount++{ 
            x = v_list4.At(CLcount)
            var try_2 EID
            try_2 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
            if ErrorIn(try_2) {try_1 = try_2
            break
            } else {
            v_local4 = ANY(try_2)
            ToList(OBJ(try_1)).PutAt(CLcount,v_local4)
            } 
          }
          } 
        if ErrorIn(try_1) {Result = try_1
        } else {
        _Ztype = ToList(OBJ(try_1))
        if (s.Id() == C_safe.Id()) { 
          Result = _Ztype.At(0).ToEID()
          }  else if ((s.Id() == Core.C_externC.Id()) && 
            ((l.Length() == 2) && 
              (C_class.Id() == l.At(1).Isa.Id()))) { 
          Result = l.At(1).ToEID()
          }  else if ((s.Id() == C_new.Id()) && 
            (C_class.Id() == l.At(0).Isa.Id())) { 
          Result = l.At(0).ToEID()
          }  else if ((s.Id() == Core.C_check_in.Id()) && 
            (l.At(1).Isa.IsIn(C_type) == CTRUE)) { 
          Result = l.At(1).ToEID()
          }  else if ((s.Id() == C_nth.Id()) && 
            (ToType(_Ztype.At(0)).Included(ToType(C_array.Id())) == CTRUE)) { 
          if (Core.F_member_type(ToType(_Ztype.At(0))).Included(ToType(C_float.Id())) == CTRUE) { 
            Result = EID{C_float.Id(),0}
            } else {
            Result = EID{Core.F_member_type(ToType(_Ztype.At(0))).Id(),0}
            } 
          }  else if ((s.Id() == Core.C__at.Id()) && 
            (l.At(0).Isa.IsIn(C_property) == CTRUE)) { 
          { var p *ClaireProperty = ToProperty(l.At(0))
            { var c *ClaireAny = l.At(1)
              if ((C_class.Id() == c.Isa.Id()) && 
                  (C_method.Id() == ANY(Core.F_CALL(ToProperty(Core.C__at.Id()),ARGS(EID{p.Id(),0},c.ToEID()))).Isa.Id())) { 
                Result = EID{MakeConstantSet(ANY(Core.F_CALL(ToProperty(Core.C__at.Id()),ARGS(EID{p.Id(),0},c.ToEID())))).Id(),0}
                } else {
                Result = EID{C_any.Id(),0}
                } 
              } 
            } 
          }  else if ((s.Id() == C_get.Id()) && 
            (l.At(0).Isa.IsIn(C_relation) == CTRUE)) { 
          { var r *ClaireRelation = ToRelation(l.At(0))
            if (r.Isa.IsIn(C_property) == CTRUE) { 
              { var g0174 *ClaireProperty = ToProperty(r.Id())
                { var xs *ClaireObject = Core.F__at_property1(g0174,ToTypeExpression(_Ztype.At(1)).Class_I())
                  if (C_slot.Id() == xs.Isa.Id()) { 
                    { var g0175 *ClaireSlot = ToSlot(xs.Id())
                      if ((g0175.Range.Included(ToType(C_set.Id())) == CTRUE) && 
                          (C_compiler.Safety < 2)) { 
                        Result = EID{ToTypeExpression(OBJ(Core.F_CALL(C_range,ARGS(l.At(1).ToEID())))).Class_I().Id(),0}
                        }  else if (g0175.Range.Contains(g0175.Default) == CTRUE) { 
                        Result = EID{g0175.Range.Id(),0}
                        } else {
                        Result = EID{F_Optimize_extends_type(g0175.Range).Id(),0}
                        } 
                      } 
                    } else {
                    Result = EID{g0174.Range.Id(),0}
                    } 
                  } 
                } 
              }  else if (C_table.Id() == r.Isa.Id()) { 
              { var g0177 *ClaireTable = ToTable(r.Id())
                if (g0177.Range.Contains(g0177.Default) == CTRUE) { 
                  Result = EID{g0177.Range.Id(),0}
                  } else {
                  Result = EID{F_Optimize_extends_type(g0177.Range).Id(),0}
                  } 
                } 
              } else {
              Result = EID{CFALSE.Id(),0}
              } 
            } 
          } else {
          { var r *ClaireAny = F_Optimize_restriction_I_property(s,_Ztype,CTRUE)
            if (C_slot.Id() == r.Isa.Id()) { 
              { var g0178 *ClaireSlot = ToSlot(r)
                if ((s.Id() == C_instances.Id()) && 
                    (C_class.Id() == l.At(0).Isa.Id())) { 
                  { var _CL_obj *ClaireParam = To_Param(new(ClaireParam).Is(C_Param))
                    _CL_obj.Arg = C_list
                    _CL_obj.Params = MakeConstantList(C_of.Id())
                    _CL_obj.Args = MakeConstantList(MakeConstantSet(l.At(0)).Id())
                    Result = EID{_CL_obj.Id(),0}
                    } 
                  } else {
                  Result = EID{g0178.Range.Id(),0}
                  } 
                } 
              }  else if (C_method.Id() == r.Isa.Id()) { 
              { var g0179 *ClaireMethod = ToMethod(r)
                Result = F_Optimize_use_range_method(g0179,_Ztype)
                } 
              }  else if (F_boolean_I_any(s.Restrictions.Id()).Id() != CTRUE.Id()) { 
              Result = EID{F_Optimize_selector_psort_Call(self).Id(),0}
              }  else if ((s.Open == 3) || 
                (r != C_Optimize_ambiguous.Id())) { 
              Result = EID{s.Range.Class_I().Id(),0}
              } else {
              Result = EID{ToTypeExpression(F_Optimize_restriction_I_property(s,_Ztype,CFALSE)).Class_I().Id(),0}
              } 
            } 
          } 
        }
        } 
      } 
    } 
  } 
return Result} 

// The EID go function for: c_type @ Call (throw: true) 
func E_c_type_Call (self EID) EID { 
return F_c_type_Call(Language.To_Call(OBJ(self)) )} 

// this is the optimizer for messages : does not use the sort unless there is a macro
/* The go function for: c_code(self:Call) [status=1] */
func F_c_code_Call (self *Language.Call) EID { 
var Result EID
Result = F_Optimize_c_code_call_Call(self,C_void)
return Result} 

// The EID go function for: c_code @ Call (throw: true) 
func E_c_code_Call (self EID) EID { 
return F_c_code_Call(Language.To_Call(OBJ(self)) )} 

/* The go function for: c_code_call(self:Call,sx:class) [status=1] */
func F_Optimize_c_code_call_Call (self *Language.Call,sx *ClaireClass) EID { 
var Result EID

{ var s *ClaireProperty = self.Selector
  { var l *ClaireList = self.Args
    var g0185I *ClaireBoolean
    var try_1 EID
    { 
      var v_and2 *ClaireBoolean
      
      v_and2 = l.At(0).Isa.IsIn(Core.C_global_variable)
      if (v_and2 == CFALSE) {try_1 = EID{CFALSE.Id(),0}
      } else { 
        v_and2 = Equal(ANY(Core.F_CALL(C_range,ARGS(l.At(0).ToEID()))),CEMPTY.Id())
        if (v_and2 == CFALSE) {try_1 = EID{CFALSE.Id(),0}
        } else { 
          var try_2 EID
          try_2 = F_Compile_designated_ask_any(ANY(Core.F_CALL(C_value,ARGS(l.At(0).ToEID()))))
          if ErrorIn(try_2) {try_1 = try_2
          } else {
          v_and2 = ToBoolean(OBJ(try_2))
          if (v_and2 == CFALSE) {try_1 = EID{CFALSE.Id(),0}
          } else { 
            try_1 = EID{CTRUE.Id(),0}} 
          } 
        } 
      }
      } 
    if ErrorIn(try_1) {Result = try_1
    } else {
    g0185I = ToBoolean(OBJ(try_1))
    if (g0185I == CTRUE) { 
      Result = ToArray(l.Id()).NthPut(1,ANY(Core.F_CALL(C_value,ARGS(l.At(0).ToEID())))).ToEID()
      } else {
      Result = EID{CFALSE.Id(),0}
      } 
    }
    if !ErrorIn(Result) {
    { var m *ClaireAny
      var try_3 EID
      try_3 = F_Optimize_inline_optimize_ask_Call(self)
      if ErrorIn(try_3) {Result = try_3
      } else {
      m = ANY(try_3)
      { var b *ClaireBoolean = l.At(0).Isa.IsIn(C_property)
        { var d *ClaireAny
          var try_4 EID
          try_4 = F_Optimize_daccess_any(self.Id(),Core.F__sup_integer(C_compiler.Safety,5))
          if ErrorIn(try_4) {Result = try_4
          } else {
          d = ANY(try_4)
          if ((b == CTRUE) && 
              (((s.Id() == Core.C_write.Id()) || 
                  (s.Id() == C_put.Id())) && 
                (l.Length() == 3))) { 
            Result = F_Optimize_c_code_write_Call(self)
            }  else if ((b == CTRUE) && 
              ((s.Id() == Core.C_put_store.Id()) && 
                ((l.Length() == 4) && 
                  (l.At(3) == CTRUE.Id())))) { 
            Result = F_Optimize_c_code_write_Call(self)
            }  else if ((b == CTRUE) && 
              (s.Id() == Core.C_unknown_ask.Id())) { 
            Result = F_Optimize_c_code_hold_property(ToProperty(l.At(0)),l.At(1),CNULL,CTRUE)
            }  else if ((b == CTRUE) && 
              (s.Id() == Core.C_known_ask.Id())) { 
            Result = F_Optimize_c_code_hold_property(ToProperty(l.At(0)),l.At(1),CNULL,CFALSE)
            }  else if ((b == CTRUE) && 
              ((s.Id() == Core.C_erase.Id()) && 
                (l.At(1).Isa.IsIn(C_Variable) == CTRUE))) { 
            { var arg_5 *ClaireAny
              var try_6 EID
              try_6 = F_Optimize_Produce_erase_property(ToProperty(l.At(0)),To_Variable(l.At(1)))
              if ErrorIn(try_6) {Result = try_6
              } else {
              arg_5 = ANY(try_6)
              Result = Core.F_CALL(C_c_code,ARGS(arg_5.ToEID(),EID{sx.Id(),0}))
              }
              } 
            }  else if (s.Id() == C_safe.Id()) { 
            { var y int = C_compiler.Safety
              { var b *ClaireBoolean = C_compiler.Overflow_ask
                { var x *ClaireAny = CNULL
                  C_compiler.Safety = 1
                  C_compiler.Overflow_ask = CTRUE
                  var try_7 EID
                  { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                    _CL_obj.Selector = C_safe
                    { 
                      var va_arg1 *Language.Call
                      var va_arg2 *ClaireList
                      va_arg1 = _CL_obj
                      var try_8 EID
                      { 
                        var v_bag_arg *ClaireAny
                        try_8= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                        var try_9 EID
                        try_9 = Core.F_CALL(C_c_code,ARGS(l.At(0).ToEID(),EID{sx.Id(),0}))
                        if ErrorIn(try_9) {try_8 = try_9
                        } else {
                        v_bag_arg = ANY(try_9)
                        ToList(OBJ(try_8)).AddFast(v_bag_arg)}
                        } 
                      if ErrorIn(try_8) {try_7 = try_8
                      } else {
                      va_arg2 = ToList(OBJ(try_8))
                      va_arg1.Args = va_arg2
                      try_7 = EID{va_arg2.Id(),0}
                      }
                      } 
                    if !ErrorIn(try_7) {
                    try_7 = EID{_CL_obj.Id(),0}
                    }
                    } 
                  if ErrorIn(try_7) {Result = try_7
                  } else {
                  x = ANY(try_7)
                  Result = x.ToEID()
                  C_compiler.Safety = y
                  C_compiler.Overflow_ask = b
                  Result = x.ToEID()
                  }
                  } 
                } 
              } 
            }  else if (((s.Id() == C_add.Id()) || 
                (s.Id() == C_add_I.Id())) && 
              (b == CTRUE)) { 
            Result = F_Optimize_c_code_add_Call(self)
            } else {
            var g0186I *ClaireBoolean
            var try_10 EID
            { 
              var v_and6 *ClaireBoolean
              
              v_and6 = MakeBoolean((s.Id() == C_add.Id()) || (s.Id() == C_add_I.Id()))
              if (v_and6 == CFALSE) {try_10 = EID{CFALSE.Id(),0}
              } else { 
                var try_11 EID
                { var arg_12 *ClaireType
                  var try_13 EID
                  try_13 = Core.F_CALL(C_c_type,ARGS(l.At(0).ToEID()))
                  if ErrorIn(try_13) {try_11 = try_13
                  } else {
                  arg_12 = ToType(OBJ(try_13))
                  try_11 = EID{arg_12.Included(ToType(C_bag.Id())).Id(),0}
                  }
                  } 
                if ErrorIn(try_11) {try_10 = try_11
                } else {
                v_and6 = ToBoolean(OBJ(try_11))
                if (v_and6 == CFALSE) {try_10 = EID{CFALSE.Id(),0}
                } else { 
                  try_10 = EID{CTRUE.Id(),0}} 
                } 
              }
              } 
            if ErrorIn(try_10) {Result = try_10
            } else {
            g0186I = ToBoolean(OBJ(try_10))
            if (g0186I == CTRUE) { 
              Result = F_Optimize_c_code_add_bag_Call(self)
              }  else if ((b == CTRUE) && 
                (s.Id() == C_delete.Id())) { 
              Result = F_Optimize_c_code_delete_Call(self)
              }  else if ((C_OPT.ToRemove.Contain_ask(s.Id()) == CTRUE) || 
                ((s.Id() == C_c_interface.Id()) && 
                    ((l.Length() == 2) && 
                      (C_OPT.LegalModules.Contain_ask(F_Compile_get_module_object(ToObject(s.Id()))) != CTRUE)))) { 
              Result = EID{CNIL.Id(),0}
              }  else if (d != CNULL) { 
              Result = d.ToEID()
              }  else if (C_method.Id() == m.Isa.Id()) { 
              { var arg_14 *ClaireClass
                var try_15 EID
                try_15 = F_Optimize_c_srange_method(ToMethod(m))
                if ErrorIn(try_15) {Result = try_15
                } else {
                arg_14 = ToClass(OBJ(try_15))
                Result = F_Optimize_c_inline_method1(ToMethod(m),l,arg_14)
                }
                } 
              } else {
              var g0187I *ClaireBoolean
              var try_16 EID
              { 
                var v_and7 *ClaireBoolean
                
                v_and7 = MakeBoolean((s.Id() == Core.C__I_equal.Id()) || (s.Id() == C__equal.Id()))
                if (v_and7 == CFALSE) {try_16 = EID{CFALSE.Id(),0}
                } else { 
                  var try_17 EID
                  { var arg_18 *ClaireAny
                    var try_19 EID
                    try_19 = F_Optimize_daccess_any(l.At(0),CTRUE)
                    if ErrorIn(try_19) {try_17 = try_19
                    } else {
                    arg_18 = ANY(try_19)
                    try_17 = EID{Core.F_known_ask_any(arg_18).Id(),0}
                    }
                    } 
                  if ErrorIn(try_17) {try_16 = try_17
                  } else {
                  v_and7 = ToBoolean(OBJ(try_17))
                  if (v_and7 == CFALSE) {try_16 = EID{CFALSE.Id(),0}
                  } else { 
                    try_16 = EID{CTRUE.Id(),0}} 
                  } 
                }
                } 
              if ErrorIn(try_16) {Result = try_16
              } else {
              g0187I = ToBoolean(OBJ(try_16))
              if (g0187I == CTRUE) { 
                Result = F_Optimize_c_code_hold_property(ToProperty(ToList(OBJ(Core.F_CALL(C_args,ARGS(l.At(0).ToEID())))).At(0)),ToList(OBJ(Core.F_CALL(C_args,ARGS(l.At(0).ToEID())))).At(1),l.At(1),Equal(s.Id(),C__equal.Id()))
                } else {
                var g0188I *ClaireBoolean
                var try_20 EID
                { 
                  var v_and8 *ClaireBoolean
                  
                  v_and8 = MakeBoolean((s.Id() == Core.C__I_equal.Id()) || (s.Id() == C__equal.Id()))
                  if (v_and8 == CFALSE) {try_20 = EID{CFALSE.Id(),0}
                  } else { 
                    var try_21 EID
                    { var arg_22 *ClaireAny
                      var try_23 EID
                      try_23 = F_Optimize_daccess_any(l.At(1),CTRUE)
                      if ErrorIn(try_23) {try_21 = try_23
                      } else {
                      arg_22 = ANY(try_23)
                      try_21 = EID{Core.F_known_ask_any(arg_22).Id(),0}
                      }
                      } 
                    if ErrorIn(try_21) {try_20 = try_21
                    } else {
                    v_and8 = ToBoolean(OBJ(try_21))
                    if (v_and8 == CFALSE) {try_20 = EID{CFALSE.Id(),0}
                    } else { 
                      try_20 = EID{CTRUE.Id(),0}} 
                    } 
                  }
                  } 
                if ErrorIn(try_20) {Result = try_20
                } else {
                g0188I = ToBoolean(OBJ(try_20))
                if (g0188I == CTRUE) { 
                  Result = F_Optimize_c_code_hold_property(ToProperty(ToList(OBJ(Core.F_CALL(C_args,ARGS(l.At(1).ToEID())))).At(0)),ToList(OBJ(Core.F_CALL(C_args,ARGS(l.At(1).ToEID())))).At(1),l.At(0),Equal(s.Id(),C__equal.Id()))
                  }  else if (((s.Id() == C_put.Id()) || 
                      (s.Id() == C_nth_equal.Id())) && 
                    ((C_table.Id() == l.At(0).Isa.Id()) && 
                      (l.Length() == 3))) { 
                  Result = F_Optimize_c_code_table_Call(self)
                  } else {
                  var g0189I *ClaireBoolean
                  var try_24 EID
                  { 
                    var v_and9 *ClaireBoolean
                    
                    v_and9 = MakeBoolean((s.Id() == C_nth_equal.Id()) || (s.Id() == C_nth_put.Id()))
                    if (v_and9 == CFALSE) {try_24 = EID{CFALSE.Id(),0}
                    } else { 
                      var try_25 EID
                      { var arg_26 *ClaireType
                        var try_27 EID
                        try_27 = Core.F_CALL(C_c_type,ARGS(l.At(0).ToEID()))
                        if ErrorIn(try_27) {try_25 = try_27
                        } else {
                        arg_26 = ToType(OBJ(try_27))
                        try_25 = EID{arg_26.Included(ToType(C_array.Id())).Id(),0}
                        }
                        } 
                      if ErrorIn(try_25) {try_24 = try_25
                      } else {
                      v_and9 = ToBoolean(OBJ(try_25))
                      if (v_and9 == CFALSE) {try_24 = EID{CFALSE.Id(),0}
                      } else { 
                        v_and9 = Equal(MakeInteger(l.Length()).Id(),MakeInteger(3).Id())
                        if (v_and9 == CFALSE) {try_24 = EID{CFALSE.Id(),0}
                        } else { 
                          try_24 = EID{CTRUE.Id(),0}} 
                        } 
                      } 
                    }
                    } 
                  if ErrorIn(try_24) {Result = try_24
                  } else {
                  g0189I = ToBoolean(OBJ(try_24))
                  if (g0189I == CTRUE) { 
                    Result = F_Optimize_c_code_array_Call(self)
                    }  else if ((s.Id() == C_nth.Id()) || 
                      (((s.Id() == C_get.Id()) && 
                          (C_table.Id() == l.At(0).Isa.Id())) || 
                        ((s.Id() == C_nth_get.Id()) && 
                            (l.At(0).Isa.IsIn(C_array) == CTRUE)))) { 
                    Result = F_Optimize_c_code_nth_Call(self)
                    }  else if (s.Id() == C__Z.Id()) { 
                    Result = F_Optimize_c_code_belong_Call(self)
                    }  else if (s.Id() == Core.C_Id.Id()) { 
                    { var arg_28 *ClaireAny
                      var try_29 EID
                      try_29 = EVAL(l.At(0))
                      if ErrorIn(try_29) {Result = try_29
                      } else {
                      arg_28 = ANY(try_29)
                      Result = Core.F_CALL(C_c_code,ARGS(arg_28.ToEID()))
                      }
                      } 
                    }  else if (s.Id() == Language.C_function_I.Id()) { 
                    { var arg_30 *ClaireString
                      var try_31 EID
                      { var arg_32 *ClaireSymbol
                        var try_33 EID
                        try_33 = Language.F_extract_symbol_any(l.At(0))
                        if ErrorIn(try_33) {try_31 = try_33
                        } else {
                        arg_32 = ToSymbol(OBJ(try_33))
                        try_31 = EID{arg_32.String_I().Id(),0}
                        }
                        } 
                      if ErrorIn(try_31) {Result = try_31
                      } else {
                      arg_30 = ToString(OBJ(try_31))
                      Result = F_make_function_string(arg_30).ToEID()
                      }
                      } 
                    }  else if ((s.Id() == Core.C_not.Id()) && 
                      (l.At(0).Isa.IsIn(Language.C_Select) == CTRUE)) { 
                    Result = F_Optimize_c_code_not_Select(Language.To_Select(l.At(0)))
                    }  else if ((s.Id() == Core.C_call.Id()) && 
                      (l.At(0).Isa.IsIn(C_property) == CTRUE)) { 
                    { var arg_34 *Language.Call
                      var try_35 EID
                      { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                        _CL_obj.Selector = ToProperty(l.At(0))
                        { 
                          var va_arg1 *Language.Call
                          var va_arg2 *ClaireList
                          va_arg1 = _CL_obj
                          var try_36 EID
                          try_36 = l.Cdr()
                          if ErrorIn(try_36) {try_35 = try_36
                          } else {
                          va_arg2 = ToList(OBJ(try_36))
                          va_arg1.Args = va_arg2
                          try_35 = EID{va_arg2.Id(),0}
                          }
                          } 
                        if !ErrorIn(try_35) {
                        try_35 = EID{_CL_obj.Id(),0}
                        }
                        } 
                      if ErrorIn(try_35) {Result = try_35
                      } else {
                      arg_34 = Language.To_Call(OBJ(try_35))
                      Result = Core.F_CALL(C_c_code,ARGS(EID{arg_34.Id(),0}))
                      }
                      } 
                    }  else if (s.Open == 3) { 
                    { var arg_37 *ClaireList
                      var try_38 EID
                      { 
                        var v_list11 *ClaireList
                        var x *ClaireAny
                        var v_local11 *ClaireAny
                        v_list11 = l
                        try_38 = EID{CreateList(ToType(CEMPTY.Id()),v_list11.Length()).Id(),0}
                        for CLcount := 0; CLcount < v_list11.Length(); CLcount++{ 
                          x = v_list11.At(CLcount)
                          var try_39 EID
                          try_39 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
                          if ErrorIn(try_39) {try_38 = try_39
                          break
                          } else {
                          v_local11 = ANY(try_39)
                          ToList(OBJ(try_38)).PutAt(CLcount,v_local11)
                          } 
                        }
                        } 
                      if ErrorIn(try_38) {Result = try_38
                      } else {
                      arg_37 = ToList(OBJ(try_38))
                      Result = F_Optimize_c_warn_property(s,l,arg_37)
                      }
                      } 
                    } else {
                    var g0190I *ClaireBoolean
                    { 
                      var v_and10 *ClaireBoolean
                      
                      v_and10 = Equal(s.Id(),Language.C_bit_vector.Id())
                      if (v_and10 == CFALSE) {g0190I = CFALSE
                      } else { 
                        { var arg_40 *ClaireAny
                          { 
                            var y *ClaireAny
                            _ = y
                            arg_40= CFALSE.Id()
                            var y_support *ClaireList
                            y_support = self.Args
                            y_len := y_support.Length()
                            for i_it := 0; i_it < y_len; i_it++ { 
                              y = y_support.At(i_it)
                              if (C_integer.Id() != y.Isa.Id()) { 
                                arg_40 = CTRUE.Id()
                                break
                                } 
                              } 
                            } 
                          v_and10 = Core.F_not_any(arg_40)
                          } 
                        if (v_and10 == CFALSE) {g0190I = CFALSE
                        } else { 
                          g0190I = CTRUE} 
                        } 
                      } 
                    if (g0190I == CTRUE) { 
                      Result = EVAL(self.Id())
                      }  else if ((s.Id() == C_Compile_anyObject_I.Id()) || 
                        ((s.Id() == C_Compile_object_I.Id()) || 
                          ((s.Id() == C_add_method.Id()) && 
                              (b == CTRUE)))) { 
                      Result = EID{self.Id(),0}
                      } else {
                      { var _Ztype *ClaireList
                        var try_41 EID
                        { 
                          var v_list12 *ClaireList
                          var x *ClaireAny
                          var v_local12 *ClaireAny
                          v_list12 = l
                          try_41 = EID{CreateList(ToType(CEMPTY.Id()),v_list12.Length()).Id(),0}
                          for CLcount := 0; CLcount < v_list12.Length(); CLcount++{ 
                            x = v_list12.At(CLcount)
                            var try_42 EID
                            try_42 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
                            if ErrorIn(try_42) {try_41 = try_42
                            break
                            } else {
                            v_local12 = ANY(try_42)
                            ToList(OBJ(try_41)).PutAt(CLcount,v_local12)
                            } 
                          }
                          } 
                        if ErrorIn(try_41) {Result = try_41
                        } else {
                        _Ztype = ToList(OBJ(try_41))
                        { var z *ClaireAny = F_Optimize_restriction_I_property(s,_Ztype,CTRUE)
                          if (C_slot.Id() == z.Isa.Id()) { 
                            { var g0181 *ClaireSlot = ToSlot(z)
                              { var _Zunknown *ClaireBoolean = MakeBoolean((g0181.Range.Contains(g0181.Default) != CTRUE) && (C_OPT.Knowns.Contain_ask(s.Id()) != CTRUE) && (C_compiler.Safety < 2))
                                var g0191I *ClaireBoolean
                                var try_43 EID
                                { 
                                  var v_or16 *ClaireBoolean
                                  
                                  v_or16 = _Zunknown.Not
                                  if (v_or16 == CTRUE) {try_43 = EID{CTRUE.Id(),0}
                                  } else { 
                                    var try_44 EID
                                    try_44 = F_Compile_designated_ask_any(l.At(0))
                                    if ErrorIn(try_44) {try_43 = try_44
                                    } else {
                                    v_or16 = ToBoolean(OBJ(try_44))
                                    if (v_or16 == CTRUE) {try_43 = EID{CTRUE.Id(),0}
                                    } else { 
                                      try_43 = EID{CFALSE.Id(),0}} 
                                    } 
                                  }
                                  } 
                                if ErrorIn(try_43) {Result = try_43
                                } else {
                                g0191I = ToBoolean(OBJ(try_43))
                                if (g0191I == CTRUE) { 
                                  { var _CL_obj *Language.CallSlot = Language.To_CallSlot(new(Language.CallSlot).Is(Language.C_Call_slot))
                                    _CL_obj.Selector = g0181
                                    { 
                                      var va_arg1 *Language.CallSlot
                                      var va_arg2 *ClaireAny
                                      va_arg1 = _CL_obj
                                      var try_45 EID
                                      try_45 = Core.F_CALL(C_c_code,ARGS(l.At(0).ToEID(),EID{F_Compile_psort_any(Core.F_domain_I_restriction(ToRestriction(g0181.Id())).Id()).Id(),0}))
                                      if ErrorIn(try_45) {Result = try_45
                                      } else {
                                      va_arg2 = ANY(try_45)
                                      va_arg1.Arg = va_arg2
                                      Result = va_arg2.ToEID()
                                      }
                                      } 
                                    if !ErrorIn(Result) {
                                    _CL_obj.Test = _Zunknown
                                    Result = EID{_CL_obj.Id(),0}
                                    }
                                    } 
                                  } else {
                                  Core.F_tformat_string(MakeString("---- note: ... unsafe access to unknown slot: ~S \n"),3,MakeConstantList(self.Id()))
                                  if (C_compiler.Optimize_ask == CTRUE) { 
                                    F_Compile_notice_void()
                                    Core.F_tformat_string(MakeString("poorly optimized slot access: ~S\n"),3,MakeConstantList(self.Id()))
                                    } 
                                  Result = F_Optimize_c_warn_property(s,l,_Ztype)
                                  } 
                                }
                                } 
                              } 
                            }  else if (C_method.Id() == z.Isa.Id()) { 
                            { var g0182 *ClaireMethod = ToMethod(z)
                              
                              if (_Ztype.Memq(C_void.Id()) == CTRUE) { 
                                Result = F_Compile_Cerror_string(MakeString("[205] call ~S uses a void argument [~S]"),ToList(MakeConstantList(self.Id(),_Ztype.Id()).Id()))
                                } else {
                                Result = EID{CFALSE.Id(),0}
                                } 
                              if !ErrorIn(Result) {
                              if (((s.Id() == C_begin.Id()) || 
                                    (s.Id() == C_end.Id())) && 
                                  (l.At(0).Isa.IsIn(C_module) == CTRUE)) { 
                                Result = EVAL(self.Id())
                                } else {
                                Result = EID{CFALSE.Id(),0}
                                } 
                              if !ErrorIn(Result) {
                              if (g0182.Domain.ValuesO()[g0182.Domain.Length()-1] == C_listargs.Id()) { 
                                { var arg_46 *ClaireList
                                  var try_47 EID
                                  try_47 = F_listargsFormat_list2(l,g0182.Domain.Length())
                                  if ErrorIn(try_47) {Result = try_47
                                  } else {
                                  arg_46 = ToList(OBJ(try_47))
                                  Result = F_Optimize_c_code_method_method2(g0182,arg_46,_Ztype,sx)
                                  }
                                  } 
                                }  else if ((g0182.Domain.ValuesO()[0] == C_void.Id()) && 
                                  (l.At(0) != ClEnv.Id())) { 
                                Result = F_Optimize_open_message_property(s,l)
                                } else {
                                Result = F_Optimize_c_code_method_method2(g0182,l,_Ztype,sx)
                                } 
                              }}
                              } 
                            }  else if (z.Isa.IsIn(C_keyword) == CTRUE) { 
                            Result = F_Optimize_c_warn_property(s,l,_Ztype)
                            } else {
                            Result = F_Optimize_c_warn_Call(self,_Ztype.Id())
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
          } 
        } 
      }
      } 
    }
    } 
  } 
return Result} 

// The EID go function for: c_code_call @ Call (throw: true) 
func E_Optimize_c_code_call_Call (self EID,sx EID) EID { 
return F_Optimize_c_code_call_Call(Language.To_Call(OBJ(self)),ToClass(OBJ(sx)) )} 

// true error
// new in v4.12 : keep the n-1 args and group the rest in a List
/* The go function for: listargsFormat(l:list,n:integer) [status=1] */
func F_listargsFormat_list2 (l *ClaireList,n int) EID { 
var Result EID
{ var i_bag *ClaireList = ToType(CEMPTY.Id()).EmptyList()
  { var i int = 1
    { var g0192 int = n
      for (i <= g0192) { 
        { var arg_1 *ClaireAny
          if (i < n) { 
            arg_1 = l.At(i-1)
            } else {
            { var _CL_obj *Language.List = Language.To_List(new(Language.List).Is(Language.C_List))
              { 
                var va_arg1 *Language.Construct
                var va_arg2 *ClaireList
                va_arg1 = Language.To_Construct(_CL_obj.Id())
                { var j_bag *ClaireList = ToType(CEMPTY.Id()).EmptyList()
                  { var j int = n
                    { var g0193 int = l.Length()
                      for (j <= g0193) { 
                        j_bag.AddFast(l.At(j-1))
                        j = (j+1)
                        } 
                      } 
                    } 
                  va_arg2 = j_bag
                  } 
                va_arg1.Args = va_arg2
                } 
              arg_1 = _CL_obj.Id()
              } 
            } 
          i_bag.AddFast(arg_1)
          } 
        i = (i+1)
        } 
      } 
    } 
  Result = EID{i_bag.Id(),0}
  } 
return Result} 

// The EID go function for: listargsFormat @ list (throw: true) 
func E_listargsFormat_list2 (l EID,n EID) EID { 
return F_listargsFormat_list2(ToList(OBJ(l)),INT(n) )} 

// create the compiled message with necessary protections
/* The go function for: open_message(self:property,l:list) [status=1] */
func F_Optimize_open_message_property (self *ClaireProperty,l *ClaireList) EID { 
var Result EID
F_Optimize_selector_register_property(self)
{ var _Zarg *ClaireList
  var try_1 EID
  { 
    var v_list1 *ClaireList
    var x *ClaireAny
    var v_local1 *ClaireAny
    v_list1 = l
    try_1 = EID{CreateList(ToType(CEMPTY.Id()),v_list1.Length()).Id(),0}
    for CLcount := 0; CLcount < v_list1.Length(); CLcount++{ 
      x = v_list1.At(CLcount)
      var try_2 EID
      var g0195I *ClaireBoolean
      var try_3 EID
      { var arg_4 *ClaireType
        var try_5 EID
        try_5 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
        if ErrorIn(try_5) {try_3 = try_5
        } else {
        arg_4 = ToType(OBJ(try_5))
        try_3 = EID{Core.F__I_equal_any(arg_4.Id(),C_void.Id()).Id(),0}
        }
        } 
      if ErrorIn(try_3) {try_2 = try_3
      } else {
      g0195I = ToBoolean(OBJ(try_3))
      if (g0195I == CTRUE) { 
        try_2 = Core.F_CALL(C_c_code,ARGS(x.ToEID(),EID{C_any.Id(),0}))
        } else {
        try_2 = F_Compile_Cerror_string(MakeString("[206] use of void ~S in ~S~S"),ToList(MakeConstantList(x,self.Id(),l.Id()).Id()))
        } 
      }
      if ErrorIn(try_2) {try_1 = try_2
      break
      } else {
      v_local1 = ANY(try_2)
      ToList(OBJ(try_1)).PutAt(CLcount,v_local1)
      } 
    }
    } 
  if ErrorIn(try_1) {Result = try_1
  } else {
  _Zarg = ToList(OBJ(try_1))
  C_compiler.NDynamic = (C_compiler.NDynamic+1)
  { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
    _CL_obj.Selector = self
    _CL_obj.Args = _Zarg
    Result = EID{_CL_obj.Id(),0}
    } 
  }
  } 
return Result} 

// The EID go function for: open_message @ property (throw: true) 
func E_Optimize_open_message_property (self EID,l EID) EID { 
return F_Optimize_open_message_property(ToProperty(OBJ(self)),ToList(OBJ(l)) )} 

// ******************************************************************
// *    Part 3: specialized c_code                                  *
// ******************************************************************
// a get message is special since it represent a direct access. The boolean
// tells if we accept a special form of the unknown value
/* The go function for: daccess(self:any,b:boolean) [status=1] */
func F_Optimize_daccess_any (self *ClaireAny,b *ClaireBoolean) EID { 
var Result EID
if (self.Isa.IsIn(Language.C_Call) == CTRUE) { 
  { var g0196 *Language.Call = Language.To_Call(self)
    { var l *ClaireList = g0196.Args
      { var xs *ClaireObject
        var try_1 EID
        if ((g0196.Selector.Id() == C_get.Id()) && 
            (l.At(0).Isa.IsIn(C_property) == CTRUE)) { 
          { var arg_2 *ClaireClass
            var try_3 EID
            { var arg_4 *ClaireType
              var try_5 EID
              try_5 = Core.F_CALL(C_c_type,ARGS(l.At(1).ToEID()))
              if ErrorIn(try_5) {try_3 = try_5
              } else {
              arg_4 = ToType(OBJ(try_5))
              try_3 = EID{arg_4.Class_I().Id(),0}
              }
              } 
            if ErrorIn(try_3) {try_1 = try_3
            } else {
            arg_2 = ToClass(OBJ(try_3))
            try_1 = EID{Core.F__at_property1(ToProperty(l.At(0)),arg_2).Id(),0}
            }
            } 
          } else {
          try_1 = EID{CFALSE.Id(),0}
          } 
        if ErrorIn(try_1) {Result = try_1
        } else {
        xs = ToObject(OBJ(try_1))
        var g0200I *ClaireBoolean
        if (C_slot.Id() == xs.Isa.Id()) { 
          { var g0197 *ClaireSlot = ToSlot(xs.Id())
            g0200I = MakeBoolean((b == CTRUE) || (g0197.Range.Contains(g0197.Default) == CTRUE) || (g0197.Srange.Id() == C_any.Id()) || (g0197.Srange.Id() == C_integer.Id()))
            } 
          } else {
          g0200I = CFALSE
          } 
        if (g0200I == CTRUE) { 
          { var _CL_obj *Language.CallSlot = Language.To_CallSlot(new(Language.CallSlot).Is(Language.C_Call_slot))
            _CL_obj.Selector = ToSlot(xs.Id())
            { 
              var va_arg1 *Language.CallSlot
              var va_arg2 *ClaireAny
              va_arg1 = _CL_obj
              var try_6 EID
              try_6 = Core.F_CALL(C_c_code,ARGS(l.At(1).ToEID(),EID{F_Compile_psort_any(Core.F_domain_I_restriction(ToRestriction(xs.Id())).Id()).Id(),0}))
              if ErrorIn(try_6) {Result = try_6
              } else {
              va_arg2 = ANY(try_6)
              va_arg1.Arg = va_arg2
              Result = va_arg2.ToEID()
              }
              } 
            if !ErrorIn(Result) {
            _CL_obj.Test = CFALSE
            Result = EID{_CL_obj.Id(),0}
            }
            } 
          } else {
          Result = EID{CNULL,0}
          } 
        }
        } 
      } 
    } 
  }  else if (self.Isa.IsIn(Language.C_Call_method2) == CTRUE) { 
  { var g0198 *Language.CallMethod2 = Language.To_CallMethod2(self)
    if (g0198.Arg.Selector.Id() == C_get.Id()) { 
      { var arg_7 *Language.Call
        { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
          _CL_obj.Selector = C_get
          _CL_obj.Args = g0198.Args
          arg_7 = _CL_obj
          } 
        Result = F_Optimize_daccess_any(arg_7.Id(),b)
        } 
      } else {
      Result = EID{CNULL,0}
      } 
    } 
  } else {
  Result = EID{CNULL,0}
  } 
return Result} 

// The EID go function for: daccess @ any (throw: true) 
func E_Optimize_daccess_any (self EID,b EID) EID { 
return F_Optimize_daccess_any(ANY(self),ToBoolean(OBJ(b)) )} 

/* The go function for: c_type(self:Call_slot) [status=0] */
func F_c_type_Call_slot (self *Language.CallSlot) *ClaireType { 
return  self.Selector.Range
} 

// The EID go function for: c_type @ Call_slot (throw: false) 
func E_c_type_Call_slot (self EID) EID { 
return EID{F_c_type_Call_slot(Language.To_CallSlot(OBJ(self)) ).Id(),0}} 

/* The go function for: c_type(self:Call_table) [status=0] */
func F_c_type_Call_table (self *Language.CallTable) *ClaireType { 
return  self.Selector.Range
} 

// The EID go function for: c_type @ Call_table (throw: false) 
func E_c_type_Call_table (self EID) EID { 
return EID{F_c_type_Call_table(Language.To_CallTable(OBJ(self)) ).Id(),0}} 

/* The go function for: c_type(self:Call_array) [status=0] */
func F_c_type_Call_array (self *Language.CallArray) *ClaireType { 
return  ToType(self.Test)
} 

// The EID go function for: c_type @ Call_array (throw: false) 
func E_c_type_Call_array (self EID) EID { 
return EID{F_c_type_Call_array(Language.To_CallArray(OBJ(self)) ).Id(),0}} 

// write optimization: ss is put, put_store or write
// note that a put(object,x that may be unknown) is hard to compile !
// v2.4.10 -> if x = unknown OK (o.r = NULL) otherwise use store
/* The go function for: c_code_write(self:Call) [status=1] */
func F_Optimize_c_code_write_Call (self *Language.Call) EID { 
var Result EID
{ var p *ClaireAny = self.Args.At(0)
  { var x *ClaireAny = self.Args.At(1)
    { var y *ClaireAny = self.Args.At(2)
      { var yt *ClaireType
        var try_1 EID
        try_1 = Core.F_CALL(C_c_type,ARGS(y.ToEID()))
        if ErrorIn(try_1) {Result = try_1
        } else {
        yt = ToType(OBJ(try_1))
        { var ss *ClaireProperty = self.Selector
          { var s *ClaireAny
            var try_2 EID
            { var arg_3 *ClaireList
              var try_4 EID
              { 
                var v_bag_arg *ClaireAny
                try_4= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                var try_5 EID
                try_5 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
                if ErrorIn(try_5) {try_4 = try_5
                } else {
                v_bag_arg = ANY(try_5)
                ToList(OBJ(try_4)).AddFast(v_bag_arg)}
                } 
              if ErrorIn(try_4) {try_2 = try_4
              } else {
              arg_3 = ToList(OBJ(try_4))
              try_2 = Core.F_CALL(C_Optimize_restriction_I,ARGS(p.ToEID(),EID{arg_3.Id(),0},EID{CTRUE.Id(),0}))
              }
              } 
            if ErrorIn(try_2) {Result = try_2
            } else {
            s = ANY(try_2)
            
            if (C_OPT.ToRemove.Contain_ask(p) == CTRUE) { 
              Result = EID{CNIL.Id(),0}
              } else {
              var g0204I *ClaireBoolean
              if (C_slot.Id() == s.Isa.Id()) { 
                { var g0201 *ClaireSlot = ToSlot(s)
                  g0204I = MakeBoolean((yt.Included(g0201.Range) == CTRUE) || (C_compiler.Safety >= 2))
                  } 
                } else {
                g0204I = CFALSE
                } 
              if (g0204I == CTRUE) { 
                var g0205I *ClaireBoolean
                var try_6 EID
                { 
                  var v_and8 *ClaireBoolean
                  
                  v_and8 = Core.F__I_equal_any(y,CNULL)
                  if (v_and8 == CFALSE) {try_6 = EID{CFALSE.Id(),0}
                  } else { 
                    var try_7 EID
                    { var arg_8 *ClaireBoolean
                      var try_9 EID
                      { var arg_10 *ClaireAny
                        var try_11 EID
                        try_11 = Core.F_CALL(ToProperty(C__exp.Id()),ARGS(EID{yt.Id(),0},Core.F_CALL(C_mClaire_srange,ARGS(s.ToEID()))))
                        if ErrorIn(try_11) {try_9 = try_11
                        } else {
                        arg_10 = ANY(try_11)
                        try_9 = EID{F_boolean_I_any(arg_10).Id(),0}
                        }
                        } 
                      if ErrorIn(try_9) {try_7 = try_9
                      } else {
                      arg_8 = ToBoolean(OBJ(try_9))
                      try_7 = EID{Core.F__I_equal_any(arg_8.Id(),CTRUE.Id()).Id(),0}
                      }
                      } 
                    if ErrorIn(try_7) {try_6 = try_7
                    } else {
                    v_and8 = ToBoolean(OBJ(try_7))
                    if (v_and8 == CFALSE) {try_6 = EID{CFALSE.Id(),0}
                    } else { 
                      try_6 = EID{CTRUE.Id(),0}} 
                    } 
                  }
                  } 
                if ErrorIn(try_6) {Result = try_6
                } else {
                g0205I = ToBoolean(OBJ(try_6))
                if (g0205I == CTRUE) { 
                  F_Compile_warn_void()
                  Result = Core.F_tformat_string(MakeString("sort error in ~S: ~S is a ~S [253]\n"),1,MakeConstantList(self.Id(),y,yt.Id()))
                  } else {
                  Result = EID{CFALSE.Id(),0}
                  } 
                }
                if !ErrorIn(Result) {
                var g0206I *ClaireBoolean
                var try_12 EID
                { 
                  var v_and8 *ClaireBoolean
                  
                  v_and8 = MakeBoolean((yt.Included(ToType(OBJ(Core.F_CALL(C_range,ARGS(s.ToEID()))))) == CTRUE) || (yt.Included(ToType(C_object.Id())) == CTRUE) || (ANY(Core.F_CALL(C_mClaire_srange,ARGS(s.ToEID()))) != C_object.Id()) || (y == CNULL))
                  if (v_and8 == CFALSE) {try_12 = EID{CFALSE.Id(),0}
                  } else { 
                    var try_13 EID
                    { 
                      var v_or10 *ClaireBoolean
                      
                      v_or10 = Core.F__I_equal_any(ss.Id(),Core.C_write.Id())
                      if (v_or10 == CTRUE) {try_13 = EID{CTRUE.Id(),0}
                      } else { 
                        var try_14 EID
                        { 
                          var v_and12 *ClaireBoolean
                          
                          var try_15 EID
                          try_15 = F_Optimize_Update_ask_relation1(ToRelation(p),x,y)
                          if ErrorIn(try_15) {try_14 = try_15
                          } else {
                          v_and12 = ToBoolean(OBJ(try_15))
                          if (v_and12 == CFALSE) {try_14 = EID{CFALSE.Id(),0}
                          } else { 
                            v_and12 = MakeBoolean((ToRelation(p).Multivalued_ask.Id() == CFALSE.Id()) || (Core.F_get_property(C_if_write,ToObject(p)) == CNULL))
                            if (v_and12 == CFALSE) {try_14 = EID{CFALSE.Id(),0}
                            } else { 
                              try_14 = EID{CTRUE.Id(),0}} 
                            } 
                          }
                          } 
                        if ErrorIn(try_14) {try_13 = try_14
                        } else {
                        v_or10 = ToBoolean(OBJ(try_14))
                        if (v_or10 == CTRUE) {try_13 = EID{CTRUE.Id(),0}
                        } else { 
                          try_13 = EID{CFALSE.Id(),0}} 
                        } 
                      }
                      } 
                    if ErrorIn(try_13) {try_12 = try_13
                    } else {
                    v_and8 = ToBoolean(OBJ(try_13))
                    if (v_and8 == CFALSE) {try_12 = EID{CFALSE.Id(),0}
                    } else { 
                      try_12 = EID{CTRUE.Id(),0}} 
                    } 
                  }
                  } 
                if ErrorIn(try_12) {Result = try_12
                } else {
                g0206I = ToBoolean(OBJ(try_12))
                if (g0206I == CTRUE) { 
                  { var _Zx *ClaireAny
                    var try_16 EID
                    try_16 = Core.F_CALL(C_c_code,ARGS(x.ToEID(),EID{F_Compile_psort_any(Core.F_domain_I_restriction(ToRestriction(s)).Id()).Id(),0}))
                    if ErrorIn(try_16) {Result = try_16
                    } else {
                    _Zx = ANY(try_16)
                    { var _Zy *ClaireAny
                      var try_17 EID
                      try_17 = F_Compile_c_strict_code_any(y,F_Compile_psort_any(ANY(Core.F_CALL(C_range,ARGS(s.ToEID())))))
                      if ErrorIn(try_17) {Result = try_17
                      } else {
                      _Zy = ANY(try_17)
                      { var _CL_obj *Language.Update = Language.To_Update(new(Language.Update).Is(Language.C_Update))
                        _CL_obj.Selector = p
                        _CL_obj.Value = _Zy
                        { 
                          var va_arg1 *Language.Update
                          var va_arg2 *ClaireAny
                          va_arg1 = _CL_obj
                          var try_18 EID
                          if (ss.Id() != Core.C_write.Id()) { 
                            try_18 = EID{ss.Id(),0}
                            } else {
                            try_18 = Core.F_CALL(C_c_code,ARGS(x.ToEID(),EID{C_any.Id(),0}))
                            } 
                          if ErrorIn(try_18) {Result = try_18
                          } else {
                          va_arg2 = ANY(try_18)
                          va_arg1.Arg = va_arg2
                          Result = va_arg2.ToEID()
                          }
                          } 
                        if !ErrorIn(Result) {
                        { 
                          var va_arg1 *Language.Update
                          var va_arg2 *ClaireAny
                          va_arg1 = _CL_obj
                          { var _CL_obj *Language.CallSlot = Language.To_CallSlot(new(Language.CallSlot).Is(Language.C_Call_slot))
                            _CL_obj.Selector = ToSlot(s)
                            _CL_obj.Arg = _Zx
                            _CL_obj.Test = CFALSE
                            va_arg2 = _CL_obj.Id()
                            } 
                          va_arg1.ClaireVar = va_arg2
                          } 
                        Result = EID{_CL_obj.Id(),0}
                        }
                        } 
                      }
                      } 
                    }
                    } 
                  }  else if (ss.Id() == C_put.Id()) { 
                  { var arg_19 *Language.Call
                    { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                      _CL_obj.Selector = C_store
                      _CL_obj.Args = MakeConstantList(x,
                        ANY(Core.F_CALL(C_mClaire_index,ARGS(s.ToEID()))),
                        ANY(Core.F_CALL(C_mClaire_srange,ARGS(s.ToEID()))),
                        y,
                        ANY(Core.F_CALL(C_store_ask,ARGS(p.ToEID()))))
                      arg_19 = _CL_obj
                      } 
                    Result = Core.F_CALL(C_c_code,ARGS(EID{arg_19.Id(),0}))
                    } 
                  } else {
                  
                  if ((C_compiler.Optimize_ask == CTRUE) && 
                      (p != C_instances.Id())) { 
                    F_Compile_notice_void()
                    Core.F_tformat_string(MakeString("poorly typed update: ~S\n"),2,MakeConstantList(self.Id()))
                    } 
                  { var arg_20 *Language.Call
                    { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                      _CL_obj.Selector = Core.C_mClaire_update
                      _CL_obj.Args = MakeConstantList(p,
                        x,
                        ANY(Core.F_CALL(C_mClaire_index,ARGS(s.ToEID()))),
                        ANY(Core.F_CALL(C_mClaire_srange,ARGS(s.ToEID()))),
                        y)
                      arg_20 = _CL_obj
                      } 
                    Result = Core.F_CALL(C_c_code,ARGS(EID{arg_20.Id(),0}))
                    } 
                  } 
                }
                }
                } else {
                { var _Ztype *ClaireList
                  var try_21 EID
                  { 
                    var v_list9 *ClaireList
                    var x *ClaireAny
                    var v_local9 *ClaireAny
                    v_list9 = self.Args
                    try_21 = EID{CreateList(ToType(CEMPTY.Id()),v_list9.Length()).Id(),0}
                    for CLcount := 0; CLcount < v_list9.Length(); CLcount++{ 
                      x = v_list9.At(CLcount)
                      var try_22 EID
                      try_22 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
                      if ErrorIn(try_22) {try_21 = try_22
                      break
                      } else {
                      v_local9 = ANY(try_22)
                      ToList(OBJ(try_21)).PutAt(CLcount,v_local9)
                      } 
                    }
                    } 
                  if ErrorIn(try_21) {Result = try_21
                  } else {
                  _Ztype = ToList(OBJ(try_21))
                  { var z *ClaireAny = F_Optimize_restriction_I_property(ss,_Ztype,CTRUE)
                    { var arg_23 *ClaireList
                      var try_24 EID
                      { 
                        var v_bag_arg *ClaireAny
                        try_24= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                        ToList(OBJ(try_24)).AddFast(self.Id())
                        var try_25 EID
                        try_25 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
                        if ErrorIn(try_25) {try_24 = try_25
                        } else {
                        v_bag_arg = ANY(try_25)
                        ToList(OBJ(try_24)).AddFast(v_bag_arg)
                        ToList(OBJ(try_24)).AddFast(yt.Id())}
                        } 
                      if ErrorIn(try_24) {Result = try_24
                      } else {
                      arg_23 = ToList(OBJ(try_24))
                      Result = Core.F_tformat_string(MakeString("---- note: ~S is poorly typed (~S,~S) \n"),3,arg_23)
                      }
                      } 
                    if !ErrorIn(Result) {
                    if (C_method.Id() == z.Isa.Id()) { 
                      { var g0202 *ClaireMethod = ToMethod(z)
                        Result = F_Optimize_c_code_method_method1(g0202,self.Args,_Ztype)
                        } 
                      } else {
                      Result = F_Optimize_c_warn_Call(self,_Ztype.Id())
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
return Result} 

// The EID go function for: c_code_write @ Call (throw: true) 
func E_Optimize_c_code_write_Call (self EID) EID { 
return F_Optimize_c_code_write_Call(Language.To_Call(OBJ(self)) )} 

// (get(p,x) =/= y) optimization. We try to use the smart form instead of the get
/* The go function for: c_code_hold(p:property,x:any,y:any,b:boolean) [status=1] */
func F_Optimize_c_code_hold_property (p *ClaireProperty,x *ClaireAny,y *ClaireAny,b *ClaireBoolean) EID { 
var Result EID
{ var s *ClaireAny
  var try_1 EID
  { var arg_2 *ClaireList
    var try_3 EID
    { 
      var v_bag_arg *ClaireAny
      try_3= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
      var try_4 EID
      try_4 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
      if ErrorIn(try_4) {try_3 = try_4
      } else {
      v_bag_arg = ANY(try_4)
      ToList(OBJ(try_3)).AddFast(v_bag_arg)}
      } 
    if ErrorIn(try_3) {try_1 = try_3
    } else {
    arg_2 = ToList(OBJ(try_3))
    try_1 = F_Optimize_restriction_I_property(p,arg_2,CTRUE).ToEID()
    }
    } 
  if ErrorIn(try_1) {Result = try_1
  } else {
  s = ANY(try_1)
  var g0208I *ClaireBoolean
  var try_5 EID
  if (C_slot.Id() == s.Isa.Id()) { 
    { var g0207 *ClaireSlot = ToSlot(s)
      { 
        var v_or3 *ClaireBoolean
        
        v_or3 = Equal(y,CNULL)
        if (v_or3 == CTRUE) {try_5 = EID{CTRUE.Id(),0}
        } else { 
          var try_6 EID
          { 
            var v_and5 *ClaireBoolean
            
            var try_7 EID
            { var arg_8 *ClaireType
              var try_9 EID
              try_9 = Core.F_CALL(C_c_type,ARGS(y.ToEID()))
              if ErrorIn(try_9) {try_7 = try_9
              } else {
              arg_8 = ToType(OBJ(try_9))
              try_7 = EID{arg_8.Included(ToType(g0207.Srange.Id())).Id(),0}
              }
              } 
            if ErrorIn(try_7) {try_6 = try_7
            } else {
            v_and5 = ToBoolean(OBJ(try_7))
            if (v_and5 == CFALSE) {try_6 = EID{CFALSE.Id(),0}
            } else { 
              var try_10 EID
              try_10 = F_Compile_identifiable_ask_any(y)
              if ErrorIn(try_10) {try_6 = try_10
              } else {
              v_and5 = ToBoolean(OBJ(try_10))
              if (v_and5 == CFALSE) {try_6 = EID{CFALSE.Id(),0}
              } else { 
                try_6 = EID{CTRUE.Id(),0}} 
              } 
            }}
            } 
          if ErrorIn(try_6) {try_5 = try_6
          } else {
          v_or3 = ToBoolean(OBJ(try_6))
          if (v_or3 == CTRUE) {try_5 = EID{CTRUE.Id(),0}
          } else { 
            try_5 = EID{CFALSE.Id(),0}} 
          } 
        }
        } 
      } 
    } else {
    try_5 = EID{CFALSE.Id(),0}
    } 
  if ErrorIn(try_5) {Result = try_5
  } else {
  g0208I = ToBoolean(OBJ(try_5))
  if (g0208I == CTRUE) { 
    { var cs *Language.CallSlot
      var try_11 EID
      { var _CL_obj *Language.CallSlot = Language.To_CallSlot(new(Language.CallSlot).Is(Language.C_Call_slot))
        _CL_obj.Selector = ToSlot(s)
        { 
          var va_arg1 *Language.CallSlot
          var va_arg2 *ClaireAny
          va_arg1 = _CL_obj
          var try_12 EID
          try_12 = Core.F_CALL(C_c_code,ARGS(x.ToEID(),EID{F_Compile_psort_any(Core.F_domain_I_restriction(ToRestriction(s)).Id()).Id(),0}))
          if ErrorIn(try_12) {try_11 = try_12
          } else {
          va_arg2 = ANY(try_12)
          va_arg1.Arg = va_arg2
          try_11 = va_arg2.ToEID()
          }
          } 
        if !ErrorIn(try_11) {
        _CL_obj.Test = CFALSE
        try_11 = EID{_CL_obj.Id(),0}
        }
        } 
      if ErrorIn(try_11) {Result = try_11
      } else {
      cs = Language.To_CallSlot(OBJ(try_11))
      { var cm *Language.CallMethod2
        var try_13 EID
        { var _CL_obj *Language.CallMethod2 = Language.To_CallMethod2(new(Language.CallMethod2).Is(Language.C_Call_method2))
          _CL_obj.Arg = ToMethod(Core.F__at_property1(Core.C_identical_ask,C_any).Id())
          { 
            var va_arg1 *Language.CallMethod
            var va_arg2 *ClaireList
            va_arg1 = Language.To_CallMethod(_CL_obj.Id())
            var try_14 EID
            { 
              var v_bag_arg *ClaireAny
              try_14= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
              ToList(OBJ(try_14)).AddFast(cs.Id())
              var try_15 EID
              try_15 = Core.F_CALL(C_c_code,ARGS(y.ToEID(),Core.F_CALL(C_mClaire_srange,ARGS(s.ToEID()))))
              if ErrorIn(try_15) {try_14 = try_15
              } else {
              v_bag_arg = ANY(try_15)
              ToList(OBJ(try_14)).AddFast(v_bag_arg)}
              } 
            if ErrorIn(try_14) {try_13 = try_14
            } else {
            va_arg2 = ToList(OBJ(try_14))
            va_arg1.Args = va_arg2
            try_13 = EID{va_arg2.Id(),0}
            }
            } 
          if !ErrorIn(try_13) {
          try_13 = EID{_CL_obj.Id(),0}
          }
          } 
        if ErrorIn(try_13) {Result = try_13
        } else {
        cm = Language.To_CallMethod2(OBJ(try_13))
        if (b == CTRUE) { 
          Result = Core.F_CALL(C_c_code,ARGS(EID{cm.Id(),0}))
          } else {
          { var arg_16 *Language.Call
            { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
              _CL_obj.Selector = Core.C_not
              _CL_obj.Args = MakeConstantList(cm.Id())
              arg_16 = _CL_obj
              } 
            Result = Core.F_CALL(C_c_code,ARGS(EID{arg_16.Id(),0}))
            } 
          } 
        }
        } 
      }
      } 
    } else {
    { var l *ClaireList = MakeConstantList(C_any.Id(),C_any.Id())
      { var arg_17 *ClaireList
        var try_18 EID
        { 
          var v_bag_arg *ClaireAny
          try_18= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
          var try_19 EID
          { var arg_20 *Language.Call
            { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
              _CL_obj.Selector = C_get
              _CL_obj.Args = MakeConstantList(p.Id(),x)
              arg_20 = _CL_obj
              } 
            try_19 = Core.F_CALL(C_c_code,ARGS(EID{arg_20.Id(),0},EID{C_any.Id(),0}))
            } 
          if ErrorIn(try_19) {try_18 = try_19
          } else {
          v_bag_arg = ANY(try_19)
          ToList(OBJ(try_18)).AddFast(v_bag_arg)
          var try_21 EID
          try_21 = Core.F_CALL(C_c_code,ARGS(y.ToEID(),EID{C_any.Id(),0}))
          if ErrorIn(try_21) {try_18 = try_21
          } else {
          v_bag_arg = ANY(try_21)
          ToList(OBJ(try_18)).AddFast(v_bag_arg)}}
          } 
        if ErrorIn(try_18) {Result = try_18
        } else {
        arg_17 = ToList(OBJ(try_18))
        Result = F_Optimize_c_code_method_method1(ToMethod(Core.F__at_property2(ToProperty(IfThenElse((b == CTRUE),
          C__equal.Id(),
          Core.C__I_equal.Id())),l).Id()),arg_17,l)
        }
        } 
      } 
    } 
  }
  }
  } 
return Result} 

// The EID go function for: c_code_hold @ property (throw: true) 
func E_Optimize_c_code_hold_property (p EID,x EID,y EID,b EID) EID { 
return F_Optimize_c_code_hold_property(ToProperty(OBJ(p)),
  ANY(x),
  ANY(y),
  ToBoolean(OBJ(b)) )} 

// add optimization
/* The go function for: c_code_add(self:Call) [status=1] */
func F_Optimize_c_code_add_Call (self *Language.Call) EID { 
var Result EID
{ var p *ClaireProperty = ToProperty(self.Args.At(0))
  { var x *ClaireAny = self.Args.At(1)
    { var y *ClaireAny = self.Args.At(2)
      { var s *ClaireObject
        var try_1 EID
        { var arg_2 *ClaireClass
          var try_3 EID
          { var arg_4 *ClaireType
            var try_5 EID
            { var arg_6 *ClaireType
              var try_7 EID
              try_7 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
              if ErrorIn(try_7) {try_5 = try_7
              } else {
              arg_6 = ToType(OBJ(try_7))
              try_5 = EID{F_Optimize_ptype_type(arg_6).Id(),0}
              }
              } 
            if ErrorIn(try_5) {try_3 = try_5
            } else {
            arg_4 = ToType(OBJ(try_5))
            try_3 = EID{arg_4.Class_I().Id(),0}
            }
            } 
          if ErrorIn(try_3) {try_1 = try_3
          } else {
          arg_2 = ToClass(OBJ(try_3))
          try_1 = EID{Core.F__at_property1(p,arg_2).Id(),0}
          }
          } 
        if ErrorIn(try_1) {Result = try_1
        } else {
        s = ToObject(OBJ(try_1))
        var g0210I *ClaireBoolean
        var try_8 EID
        if (C_slot.Id() == s.Isa.Id()) { 
          { var g0209 *ClaireSlot = ToSlot(s.Id())
            { 
              var v_or6 *ClaireBoolean
              
              var try_9 EID
              { var arg_10 *ClaireType
                var try_11 EID
                try_11 = Core.F_CALL(C_c_type,ARGS(y.ToEID()))
                if ErrorIn(try_11) {try_9 = try_11
                } else {
                arg_10 = ToType(OBJ(try_11))
                try_9 = EID{arg_10.Included(Core.F_member_type(g0209.Range)).Id(),0}
                }
                } 
              if ErrorIn(try_9) {try_8 = try_9
              } else {
              v_or6 = ToBoolean(OBJ(try_9))
              if (v_or6 == CTRUE) {try_8 = EID{CTRUE.Id(),0}
              } else { 
                v_or6 = F__sup_equal_integer(C_compiler.Safety,2)
                if (v_or6 == CTRUE) {try_8 = EID{CTRUE.Id(),0}
                } else { 
                  try_8 = EID{CFALSE.Id(),0}} 
                } 
              }
              } 
            } 
          } else {
          try_8 = EID{CFALSE.Id(),0}
          } 
        if ErrorIn(try_8) {Result = try_8
        } else {
        g0210I = ToBoolean(OBJ(try_8))
        if (g0210I == CTRUE) { 
          if (F_Optimize_Update_ask_relation2(ToRelation(p.Id()),ToRelation(self.Selector.Id())) == CTRUE) { 
            { var x2 *ClaireAny
              var try_12 EID
              try_12 = Core.F_CALL(C_c_code,ARGS(x.ToEID(),EID{F_Compile_psort_any(Core.F_domain_I_restriction(ToRestriction(s.Id())).Id()).Id(),0}))
              if ErrorIn(try_12) {Result = try_12
              } else {
              x2 = ANY(try_12)
              { var _CL_obj *Language.Update = Language.To_Update(new(Language.Update).Is(Language.C_Update))
                _CL_obj.Selector = p.Id()
                _CL_obj.Arg = C_add.Id()
                { 
                  var va_arg1 *Language.Update
                  var va_arg2 *ClaireAny
                  va_arg1 = _CL_obj
                  { var _CL_obj *Language.CallSlot = Language.To_CallSlot(new(Language.CallSlot).Is(Language.C_Call_slot))
                    _CL_obj.Selector = ToSlot(s.Id())
                    _CL_obj.Arg = x2
                    _CL_obj.Test = CFALSE
                    va_arg2 = _CL_obj.Id()
                    } 
                  va_arg1.ClaireVar = va_arg2
                  } 
                { 
                  var va_arg1 *Language.Update
                  var va_arg2 *ClaireAny
                  va_arg1 = _CL_obj
                  var try_13 EID
                  try_13 = Core.F_CALL(C_c_code,ARGS(y.ToEID(),EID{F_Compile_psort_any(Core.F_member_type(ToType(OBJ(Core.F_CALL(C_range,ARGS(EID{s.Id(),0}))))).Id()).Id(),0}))
                  if ErrorIn(try_13) {Result = try_13
                  } else {
                  va_arg2 = ANY(try_13)
                  va_arg1.Value = va_arg2
                  Result = va_arg2.ToEID()
                  }
                  } 
                if !ErrorIn(Result) {
                Result = EID{_CL_obj.Id(),0}
                }
                } 
              }
              } 
            } else {
            var g0211I *ClaireBoolean
            var try_14 EID
            { 
              var v_and6 *ClaireBoolean
              
              var try_15 EID
              try_15 = F_Compile_designated_ask_any(x)
              if ErrorIn(try_15) {try_14 = try_15
              } else {
              v_and6 = ToBoolean(OBJ(try_15))
              if (v_and6 == CFALSE) {try_14 = EID{CFALSE.Id(),0}
              } else { 
                v_and6 = MakeBoolean((p.Store_ask != CTRUE) && ((self.Selector.Id() == C_add_I.Id()) || 
                    (p.Inverse.Id() == CNULL)))
                if (v_and6 == CFALSE) {try_14 = EID{CFALSE.Id(),0}
                } else { 
                  try_14 = EID{CTRUE.Id(),0}} 
                } 
              }
              } 
            if ErrorIn(try_14) {Result = try_14
            } else {
            g0211I = ToBoolean(OBJ(try_14))
            if (g0211I == CTRUE) { 
              { var x2 *ClaireAny
                var try_16 EID
                try_16 = Core.F_CALL(C_c_code,ARGS(x.ToEID(),EID{F_Compile_psort_any(Core.F_domain_I_restriction(ToRestriction(s.Id())).Id()).Id(),0}))
                if ErrorIn(try_16) {Result = try_16
                } else {
                x2 = ANY(try_16)
                { var arg_17 *Language.Call
                  { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                    _CL_obj.Selector = ToProperty(C_add.Id())
                    { 
                      var va_arg1 *Language.Call
                      var va_arg2 *ClaireList
                      va_arg1 = _CL_obj
                      { 
                        var v_bag_arg *ClaireAny
                        va_arg2= ToType(CEMPTY.Id()).EmptyList()
                        { var _CL_obj *Language.CallSlot = Language.To_CallSlot(new(Language.CallSlot).Is(Language.C_Call_slot))
                          _CL_obj.Selector = ToSlot(s.Id())
                          _CL_obj.Arg = x2
                          _CL_obj.Test = CFALSE
                          v_bag_arg = _CL_obj.Id()
                          } 
                        va_arg2.AddFast(v_bag_arg)
                        va_arg2.AddFast(y)} 
                      va_arg1.Args = va_arg2
                      } 
                    arg_17 = _CL_obj
                    } 
                  Result = Core.F_CALL(C_c_code,ARGS(EID{arg_17.Id(),0}))
                  } 
                }
                } 
              } else {
              if (C_compiler.Optimize_ask == CTRUE) { 
                F_Compile_notice_void()
                Core.F_tformat_string(MakeString("poorly typed update: ~S\n"),3,MakeConstantList(self.Id()))
                } 
              { var arg_18 *ClaireList
                var try_19 EID
                { 
                  var v_bag_arg *ClaireAny
                  try_19= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                  ToList(OBJ(try_19)).AddFast(C_property.Id())
                  var try_20 EID
                  try_20 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
                  if ErrorIn(try_20) {try_19 = try_20
                  } else {
                  v_bag_arg = ANY(try_20)
                  ToList(OBJ(try_19)).AddFast(v_bag_arg)
                  ToList(OBJ(try_19)).AddFast(C_integer.Id())
                  var try_21 EID
                  try_21 = Core.F_CALL(C_c_type,ARGS(y.ToEID()))
                  if ErrorIn(try_21) {try_19 = try_21
                  } else {
                  v_bag_arg = ANY(try_21)
                  ToList(OBJ(try_19)).AddFast(v_bag_arg)}}
                  } 
                if ErrorIn(try_19) {Result = try_19
                } else {
                arg_18 = ToList(OBJ(try_19))
                Result = F_Optimize_c_code_method_method1(ToMethod(Core.F__at_property1(ToProperty(C_add_I.Id()),C_property).Id()),MakeConstantList(p.Id(),
                  x,
                  ANY(Core.F_CALL(C_mClaire_index,ARGS(EID{s.Id(),0}))),
                  y),arg_18)
                }
                } 
              } 
            }
            } 
          } else {
          { var arg_22 *ClaireList
            var try_23 EID
            { 
              var v_list6 *ClaireList
              var x *ClaireAny
              var v_local6 *ClaireAny
              v_list6 = self.Args
              try_23 = EID{CreateList(ToType(CEMPTY.Id()),v_list6.Length()).Id(),0}
              for CLcount := 0; CLcount < v_list6.Length(); CLcount++{ 
                x = v_list6.At(CLcount)
                var try_24 EID
                try_24 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
                if ErrorIn(try_24) {try_23 = try_24
                break
                } else {
                v_local6 = ANY(try_24)
                ToList(OBJ(try_23)).PutAt(CLcount,v_local6)
                } 
              }
              } 
            if ErrorIn(try_23) {Result = try_23
            } else {
            arg_22 = ToList(OBJ(try_23))
            Result = F_Optimize_c_code_method_method1(ToMethod(Core.F__at_property1(ToProperty(C_add.Id()),C_property).Id()),self.Args,arg_22)
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

// The EID go function for: c_code_add @ Call (throw: true) 
func E_Optimize_c_code_add_Call (self EID) EID { 
return F_Optimize_c_code_add_Call(Language.To_Call(OBJ(self)) )} 

// new in v3.0.59
/* The go function for: c_code_add_bag(self:Call) [status=1] */
func F_Optimize_c_code_add_bag_Call (self *Language.Call) EID { 
var Result EID
{ var _Zt1 *ClaireType
  var try_1 EID
  try_1 = Core.F_CALL(C_c_type,ARGS(self.Args.At(0).ToEID()))
  if ErrorIn(try_1) {Result = try_1
  } else {
  _Zt1 = ToType(OBJ(try_1))
  { var _Zt2 *ClaireType
    var try_2 EID
    { var arg_3 *ClaireType
      var try_4 EID
      try_4 = Core.F_CALL(C_c_type,ARGS(self.Args.At(1).ToEID()))
      if ErrorIn(try_4) {try_2 = try_4
      } else {
      arg_3 = ToType(OBJ(try_4))
      try_2 = EID{F_Optimize_ptype_type(arg_3).Id(),0}
      }
      } 
    if ErrorIn(try_2) {Result = try_2
    } else {
    _Zt2 = ToType(OBJ(try_2))
    { var _Zp *ClaireProperty
      if (((_Zt1.Isa.IsIn(C_Param) == CTRUE) && 
            (_Zt2.Included(Core.F_member_type(_Zt1)) == CTRUE)) || 
          (C_compiler.Safety >= 2)) { 
        _Zp = ToProperty(C_add_I.Id())
        } else {
        _Zp = self.Selector
        } 
      { var _Zltype *ClaireList = MakeConstantList(_Zt1.Id(),_Zt2.Id())
        { var z *ClaireAny = F_Optimize_restriction_I_property(_Zp,_Zltype,CTRUE)
          
          if ((_Zt2.Included(Core.F_member_type(_Zt1)) != CTRUE) && 
              (self.Selector.Id() == C_add.Id())) { 
            F_Compile_warn_void()
            Core.F_tformat_string(MakeString("the bag addition ~S is poorly typed (~S not in ~S) [251] \n"),1,MakeConstantList(self.Id(),_Zt2.Id(),Core.F_member_type(_Zt1).Id()))
            } 
          if (_Zt2.Id() == C_void.Id()) { 
            Result = F_Compile_Cerror_string(MakeString("[206] use of void ~S in ~S"),ToList(MakeConstantList(self.Args.At(1),self.Id()).Id()))
            } else {
            Result = EID{CFALSE.Id(),0}
            } 
          if !ErrorIn(Result) {
          if (C_method.Id() == z.Isa.Id()) { 
            { var g0212 *ClaireMethod = ToMethod(z)
              Result = F_Optimize_c_code_method_method1(g0212,self.Args,_Zltype)
              } 
            } else {
            Result = F_Optimize_c_warn_Call(self,_Zltype.Id())
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

// The EID go function for: c_code_add_bag @ Call (throw: true) 
func E_Optimize_c_code_add_bag_Call (self EID) EID { 
return F_Optimize_c_code_add_bag_Call(Language.To_Call(OBJ(self)) )} 

// delete optimization
// <yc> 7/98 new, also needed
/* The go function for: c_code_delete(self:Call) [status=1] */
func F_Optimize_c_code_delete_Call (self *Language.Call) EID { 
var Result EID
{ var p *ClaireAny = self.Args.At(0)
  { var x *ClaireAny = self.Args.At(1)
    { var y *ClaireAny = self.Args.At(2)
      { var s *ClaireObject
        var try_1 EID
        { var arg_2 *ClaireClass
          var try_3 EID
          { var arg_4 *ClaireType
            var try_5 EID
            try_5 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
            if ErrorIn(try_5) {try_3 = try_5
            } else {
            arg_4 = ToType(OBJ(try_5))
            try_3 = EID{arg_4.Class_I().Id(),0}
            }
            } 
          if ErrorIn(try_3) {try_1 = try_3
          } else {
          arg_2 = ToClass(OBJ(try_3))
          try_1 = EID{Core.F__at_property1(ToProperty(p),arg_2).Id(),0}
          }
          } 
        if ErrorIn(try_1) {Result = try_1
        } else {
        s = ToObject(OBJ(try_1))
        var g0215I *ClaireBoolean
        var try_6 EID
        { 
          var v_and4 *ClaireBoolean
          
          v_and4 = MakeBoolean((ToRelation(p).Inverse.Id() == CNULL))
          if (v_and4 == CFALSE) {try_6 = EID{CFALSE.Id(),0}
          } else { 
            var try_7 EID
            try_7 = F_Compile_designated_ask_any(x)
            if ErrorIn(try_7) {try_6 = try_7
            } else {
            v_and4 = ToBoolean(OBJ(try_7))
            if (v_and4 == CFALSE) {try_6 = EID{CFALSE.Id(),0}
            } else { 
              var try_8 EID
              if (C_slot.Id() == s.Isa.Id()) { 
                { var g0214 *ClaireSlot = ToSlot(s.Id())
                  { 
                    var v_or9 *ClaireBoolean
                    
                    var try_9 EID
                    { var arg_10 *ClaireType
                      var try_11 EID
                      try_11 = Core.F_CALL(C_c_type,ARGS(y.ToEID()))
                      if ErrorIn(try_11) {try_9 = try_11
                      } else {
                      arg_10 = ToType(OBJ(try_11))
                      try_9 = EID{arg_10.Included(Core.F_member_type(g0214.Range)).Id(),0}
                      }
                      } 
                    if ErrorIn(try_9) {try_8 = try_9
                    } else {
                    v_or9 = ToBoolean(OBJ(try_9))
                    if (v_or9 == CTRUE) {try_8 = EID{CTRUE.Id(),0}
                    } else { 
                      v_or9 = F__sup_equal_integer(C_compiler.Safety,2)
                      if (v_or9 == CTRUE) {try_8 = EID{CTRUE.Id(),0}
                      } else { 
                        try_8 = EID{CFALSE.Id(),0}} 
                      } 
                    }
                    } 
                  } 
                } else {
                try_8 = EID{CFALSE.Id(),0}
                } 
              if ErrorIn(try_8) {try_6 = try_8
              } else {
              v_and4 = ToBoolean(OBJ(try_8))
              if (v_and4 == CFALSE) {try_6 = EID{CFALSE.Id(),0}
              } else { 
                try_6 = EID{CTRUE.Id(),0}} 
              } 
            } 
          }}
          } 
        if ErrorIn(try_6) {Result = try_6
        } else {
        g0215I = ToBoolean(OBJ(try_6))
        if (g0215I == CTRUE) { 
          { var x2 *ClaireAny
            var try_12 EID
            try_12 = Core.F_CALL(C_c_code,ARGS(x.ToEID(),EID{F_Compile_psort_any(Core.F_domain_I_restriction(ToRestriction(s.Id())).Id()).Id(),0}))
            if ErrorIn(try_12) {Result = try_12
            } else {
            x2 = ANY(try_12)
            { var arg_13 *Language.Call
              { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                _CL_obj.Selector = ToProperty(C_delete.Id())
                { 
                  var va_arg1 *Language.Call
                  var va_arg2 *ClaireList
                  va_arg1 = _CL_obj
                  { 
                    var v_bag_arg *ClaireAny
                    va_arg2= ToType(CEMPTY.Id()).EmptyList()
                    { var _CL_obj *Language.CallSlot = Language.To_CallSlot(new(Language.CallSlot).Is(Language.C_Call_slot))
                      _CL_obj.Selector = ToSlot(s.Id())
                      _CL_obj.Arg = x2
                      _CL_obj.Test = CFALSE
                      v_bag_arg = _CL_obj.Id()
                      } 
                    va_arg2.AddFast(v_bag_arg)
                    va_arg2.AddFast(y)} 
                  va_arg1.Args = va_arg2
                  } 
                arg_13 = _CL_obj
                } 
              Result = Core.F_CALL(C_c_code,ARGS(EID{arg_13.Id(),0}))
              } 
            }
            } 
          } else {
          { var arg_14 *ClaireList
            var try_15 EID
            { 
              var v_list6 *ClaireList
              var x *ClaireAny
              var v_local6 *ClaireAny
              v_list6 = self.Args
              try_15 = EID{CreateList(ToType(CEMPTY.Id()),v_list6.Length()).Id(),0}
              for CLcount := 0; CLcount < v_list6.Length(); CLcount++{ 
                x = v_list6.At(CLcount)
                var try_16 EID
                try_16 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
                if ErrorIn(try_16) {try_15 = try_16
                break
                } else {
                v_local6 = ANY(try_16)
                ToList(OBJ(try_15)).PutAt(CLcount,v_local6)
                } 
              }
              } 
            if ErrorIn(try_15) {Result = try_15
            } else {
            arg_14 = ToList(OBJ(try_15))
            Result = F_Optimize_c_code_method_method1(ToMethod(Core.F__at_property1(ToProperty(C_delete.Id()),C_property).Id()),self.Args,arg_14)
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

// The EID go function for: c_code_delete @ Call (throw: true) 
func E_Optimize_c_code_delete_Call (self EID) EID { 
return F_Optimize_c_code_delete_Call(Language.To_Call(OBJ(self)) )} 

// cute optimization
/* The go function for: c_code_not(x:Select) [status=1] */
func F_Optimize_c_code_not_Select (x *Language.Select) EID { 
var Result EID
{ var arg_1 *Language.Call
  { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
    _CL_obj.Selector = Core.C_not
    { 
      var va_arg1 *Language.Call
      var va_arg2 *ClaireList
      va_arg1 = _CL_obj
      { 
        var v_bag_arg *ClaireAny
        va_arg2= ToType(CEMPTY.Id()).EmptyList()
        { var _CL_obj *Language.For = Language.To_For(new(Language.For).Is(Language.C_For))
          _CL_obj.ClaireVar = x.ClaireVar
          _CL_obj.SetArg = x.SetArg
          _CL_obj.Arg = Language.C_If.Make(x.Arg,Language.C_Return.Make(CTRUE.Id()),CNULL)
          v_bag_arg = _CL_obj.Id()
          } 
        va_arg2.AddFast(v_bag_arg)} 
      va_arg1.Args = va_arg2
      } 
    arg_1 = _CL_obj
    } 
  Result = Core.F_CALL(C_c_code,ARGS(EID{arg_1.Id(),0}))
  } 
return Result} 

// The EID go function for: c_code_not @ Select (throw: true) 
func E_Optimize_c_code_not_Select (x EID) EID { 
return F_Optimize_c_code_not_Select(Language.To_Select(OBJ(x)) )} 

// old % optimization
/* The go function for: c_code_belong(self:Call) [status=1] */
func F_Optimize_c_code_belong_Call (self *Language.Call) EID { 
var Result EID
{ var x *ClaireAny = self.Args.At(0)
  { var y *ClaireAny = self.Args.At(1)
    { var _Ztype *ClaireList
      var try_1 EID
      { 
        var v_bag_arg *ClaireAny
        try_1= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
        var try_2 EID
        try_2 = Core.F_CALL(C_c_type,ARGS(y.ToEID()))
        if ErrorIn(try_2) {try_1 = try_2
        } else {
        v_bag_arg = ANY(try_2)
        ToList(OBJ(try_1)).AddFast(v_bag_arg)
        var try_3 EID
        try_3 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
        if ErrorIn(try_3) {try_1 = try_3
        } else {
        v_bag_arg = ANY(try_3)
        ToList(OBJ(try_1)).AddFast(v_bag_arg)}}
        } 
      if ErrorIn(try_1) {Result = try_1
      } else {
      _Ztype = ToList(OBJ(try_1))
      if (C_set.Id() == y.Isa.Id()) { 
        { var _CL_obj *Language.Or = Language.To_Or(new(Language.Or).Is(Language.C_Or))
          { 
            var va_arg1 *Language.Or
            var va_arg2 *ClaireList
            va_arg1 = _CL_obj
            var try_4 EID
            { var z_bag *ClaireList = ToType(CEMPTY.Id()).EmptyList()
              { 
                var z *ClaireAny
                _ = z
                try_4= EID{CFALSE.Id(),0}
                var z_support *ClaireList
                var try_5 EID
                try_5 = Core.F_enumerate_any(y)
                if ErrorIn(try_5) {try_4 = try_5
                } else {
                z_support = ToList(OBJ(try_5))
                z_len := z_support.Length()
                for i_it := 0; i_it < z_len; i_it++ { 
                  z = z_support.At(i_it)
                  var loop_6 EID
                  _ = loop_6
                  { var arg_7 *ClaireAny
                    var try_8 EID
                    { var arg_9 *Language.Call
                      { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                        _CL_obj.Selector = ToProperty(C__equal.Id())
                        _CL_obj.Args = MakeConstantList(x,z)
                        arg_9 = _CL_obj
                        } 
                      try_8 = Core.F_CALL(C_c_code,ARGS(EID{arg_9.Id(),0},EID{C_any.Id(),0}))
                      } 
                    if ErrorIn(try_8) {loop_6 = try_8
                    } else {
                    arg_7 = ANY(try_8)
                    loop_6 = EID{z_bag.AddFast(arg_7).Id(),0}
                    }
                    } 
                  if ErrorIn(loop_6) {try_4 = loop_6
                  break
                  } else {
                  }}
                  } 
                } 
              if !ErrorIn(try_4) {
              try_4 = EID{z_bag.Id(),0}
              }
              } 
            if ErrorIn(try_4) {Result = try_4
            } else {
            va_arg2 = ToList(OBJ(try_4))
            va_arg1.Args = va_arg2
            Result = EID{va_arg2.Id(),0}
            }
            } 
          if !ErrorIn(Result) {
          Result = EID{_CL_obj.Id(),0}
          }
          } 
        }  else if (ToType(_Ztype.At(0)).Included(ToType(C_list.Id())) == CTRUE) { 
        Result = F_Optimize_c_code_method_method1(ToMethod(Core.F__at_property2(C_contain_ask,MakeConstantList(C_list.Id(),C_any.Id())).Id()),MakeConstantList(y,x),_Ztype)
        }  else if (ToType(_Ztype.At(0)).Included(ToType(C_set.Id())) == CTRUE) { 
        Result = F_Optimize_c_code_method_method1(ToMethod(Core.F__at_property2(C_contain_ask,MakeConstantList(C_set.Id(),C_any.Id())).Id()),MakeConstantList(y,x),_Ztype)
        }  else if (y == C_object.Id()) { 
        Result = F_Optimize_c_code_method_method1(ToMethod(Core.F__at_property2(ToProperty(C__Z.Id()),MakeConstantList(C_any.Id(),C_class.Id())).Id()),MakeConstantList(x,y),MakeConstantList(C_any.Id(),C_class.Id()))
        } else {
        Result = Core.F_CALL(C_Optimize_member_code,ARGS(y.ToEID(),x.ToEID()))
        } 
      }
      } 
    } 
  } 
return Result} 

// The EID go function for: c_code_belong @ Call (throw: true) 
func E_Optimize_c_code_belong_Call (self EID) EID { 
return F_Optimize_c_code_belong_Call(Language.To_Call(OBJ(self)) )} 

// nth optimization for arrays (the selector may also be get)
/* The go function for: c_code_nth(self:Call) [status=1] */
func F_Optimize_c_code_nth_Call (self *Language.Call) EID { 
var Result EID
{ var l *ClaireList = self.Args
  { var x *ClaireAny = l.At(0)
    { var p *ClaireProperty = self.Selector
      { var t *ClaireType
        var try_1 EID
        try_1 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
        if ErrorIn(try_1) {Result = try_1
        } else {
        t = ToType(OBJ(try_1))
        { var mt *ClaireType = Core.F_member_type(t)
          { var r *ClaireAny
            var try_2 EID
            { var arg_3 *ClaireList
              var try_4 EID
              { 
                var v_list7 *ClaireList
                var u *ClaireAny
                var v_local7 *ClaireAny
                v_list7 = l
                try_4 = EID{CreateList(ToType(CEMPTY.Id()),v_list7.Length()).Id(),0}
                for CLcount := 0; CLcount < v_list7.Length(); CLcount++{ 
                  u = v_list7.At(CLcount)
                  var try_5 EID
                  try_5 = Core.F_CALL(C_c_type,ARGS(u.ToEID()))
                  if ErrorIn(try_5) {try_4 = try_5
                  break
                  } else {
                  v_local7 = ANY(try_5)
                  ToList(OBJ(try_4)).PutAt(CLcount,v_local7)
                  } 
                }
                } 
              if ErrorIn(try_4) {try_2 = try_4
              } else {
              arg_3 = ToList(OBJ(try_4))
              try_2 = F_Optimize_restriction_I_property(p,arg_3,CTRUE).ToEID()
              }
              } 
            if ErrorIn(try_2) {Result = try_2
            } else {
            r = ANY(try_2)
            if (C_OPT.ToRemove.Contain_ask(x) == CTRUE) { 
              Result = EID{CNIL.Id(),0}
              } else {
              var g0218I *ClaireBoolean
              var try_6 EID
              { 
                var v_and7 *ClaireBoolean
                
                if (C_table.Id() == x.Isa.Id()) { 
                  { var g0216 *ClaireTable = ToTable(x)
                    v_and7 = Equal(C_integer.Id(),g0216.Params.Isa.Id())
                    } 
                  } else {
                  v_and7 = CFALSE
                  } 
                if (v_and7 == CFALSE) {try_6 = EID{CFALSE.Id(),0}
                } else { 
                  var try_7 EID
                  { 
                    var v_or9 *ClaireBoolean
                    
                    var try_8 EID
                    { var arg_9 *ClaireType
                      var try_10 EID
                      try_10 = Core.F_CALL(C_c_type,ARGS(l.At(1).ToEID()))
                      if ErrorIn(try_10) {try_8 = try_10
                      } else {
                      arg_9 = ToType(OBJ(try_10))
                      try_8 = EID{arg_9.Included(ToType(OBJ(Core.F_CALL(C_domain,ARGS(x.ToEID()))))).Id(),0}
                      }
                      } 
                    if ErrorIn(try_8) {try_7 = try_8
                    } else {
                    v_or9 = ToBoolean(OBJ(try_8))
                    if (v_or9 == CTRUE) {try_7 = EID{CTRUE.Id(),0}
                    } else { 
                      v_or9 = MakeBoolean((p.Id() == C_nth.Id()) && (C_compiler.Safety >= 2))
                      if (v_or9 == CTRUE) {try_7 = EID{CTRUE.Id(),0}
                      } else { 
                        try_7 = EID{CFALSE.Id(),0}} 
                      } 
                    }
                    } 
                  if ErrorIn(try_7) {try_6 = try_7
                  } else {
                  v_and7 = ToBoolean(OBJ(try_7))
                  if (v_and7 == CFALSE) {try_6 = EID{CFALSE.Id(),0}
                  } else { 
                    try_6 = EID{CTRUE.Id(),0}} 
                  } 
                }
                } 
              if ErrorIn(try_6) {Result = try_6
              } else {
              g0218I = ToBoolean(OBJ(try_6))
              if (g0218I == CTRUE) { 
                { var _CL_obj *Language.CallTable = Language.To_CallTable(new(Language.CallTable).Is(Language.C_Call_table))
                  _CL_obj.Selector = ToTable(x)
                  { 
                    var va_arg1 *Language.CallTable
                    var va_arg2 *ClaireAny
                    va_arg1 = _CL_obj
                    var try_11 EID
                    try_11 = Core.F_CALL(C_c_code,ARGS(l.At(1).ToEID(),EID{C_integer.Id(),0}))
                    if ErrorIn(try_11) {Result = try_11
                    } else {
                    va_arg2 = ANY(try_11)
                    va_arg1.Arg = va_arg2
                    Result = va_arg2.ToEID()
                    }
                    } 
                  if !ErrorIn(Result) {
                  { 
                    var va_arg1 *Language.CallTable
                    var va_arg2 *ClaireBoolean
                    va_arg1 = _CL_obj
                    var try_12 EID
                    { var arg_13 *ClaireBoolean
                      var try_14 EID
                      { 
                        var v_or11 *ClaireBoolean
                        
                        var try_15 EID
                        try_15 = Core.F_BELONG(ANY(Core.F_CALL(C_default,ARGS(x.ToEID()))),ANY(Core.F_CALL(C_range,ARGS(x.ToEID()))))
                        if ErrorIn(try_15) {try_14 = try_15
                        } else {
                        v_or11 = ToBoolean(OBJ(try_15))
                        if (v_or11 == CTRUE) {try_14 = EID{CTRUE.Id(),0}
                        } else { 
                          v_or11 = Equal(ANY(Core.F_CALL(C_default,ARGS(x.ToEID()))),MakeInteger(0).Id())
                          if (v_or11 == CTRUE) {try_14 = EID{CTRUE.Id(),0}
                          } else { 
                            v_or11 = Equal(p.Id(),C_get.Id())
                            if (v_or11 == CTRUE) {try_14 = EID{CTRUE.Id(),0}
                            } else { 
                              try_14 = EID{CFALSE.Id(),0}} 
                            } 
                          } 
                        }
                        } 
                      if ErrorIn(try_14) {try_12 = try_14
                      } else {
                      arg_13 = ToBoolean(OBJ(try_14))
                      try_12 = EID{arg_13.Not.Id(),0}
                      }
                      } 
                    if ErrorIn(try_12) {Result = try_12
                    } else {
                    va_arg2 = ToBoolean(OBJ(try_12))
                    va_arg1.Test = va_arg2
                    Result = EID{va_arg2.Id(),0}
                    }
                    } 
                  if !ErrorIn(Result) {
                  Result = EID{_CL_obj.Id(),0}
                  }}
                  } 
                } else {
                var g0219I *ClaireBoolean
                var try_16 EID
                { 
                  var v_and8 *ClaireBoolean
                  
                  if (C_table.Id() == x.Isa.Id()) { 
                    { var g0217 *ClaireTable = ToTable(x)
                      v_and8 = g0217.Params.Isa.IsIn(C_list)
                      } 
                    } else {
                    v_and8 = CFALSE
                    } 
                  if (v_and8 == CFALSE) {try_16 = EID{CFALSE.Id(),0}
                  } else { 
                    v_and8 = Equal(MakeInteger(l.Length()).Id(),MakeInteger(3).Id())
                    if (v_and8 == CFALSE) {try_16 = EID{CFALSE.Id(),0}
                    } else { 
                      var try_17 EID
                      { 
                        var v_or11 *ClaireBoolean
                        
                        var try_18 EID
                        { var arg_19 *ClaireTuple
                          var try_20 EID
                          { var arg_21 *ClaireList
                            var try_22 EID
                            { 
                              var v_bag_arg *ClaireAny
                              try_22= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                              var try_23 EID
                              try_23 = Core.F_CALL(C_c_type,ARGS(l.At(1).ToEID()))
                              if ErrorIn(try_23) {try_22 = try_23
                              } else {
                              v_bag_arg = ANY(try_23)
                              ToList(OBJ(try_22)).AddFast(v_bag_arg)
                              var try_24 EID
                              try_24 = Core.F_CALL(C_c_type,ARGS(l.At(2).ToEID()))
                              if ErrorIn(try_24) {try_22 = try_24
                              } else {
                              v_bag_arg = ANY(try_24)
                              ToList(OBJ(try_22)).AddFast(v_bag_arg)}}
                              } 
                            if ErrorIn(try_22) {try_20 = try_22
                            } else {
                            arg_21 = ToList(OBJ(try_22))
                            try_20 = EID{arg_21.Tuple_I().Id(),0}
                            }
                            } 
                          if ErrorIn(try_20) {try_18 = try_20
                          } else {
                          arg_19 = ToTuple(OBJ(try_20))
                          try_18 = EID{ToType(arg_19.Id()).Included(ToType(OBJ(Core.F_CALL(C_domain,ARGS(x.ToEID()))))).Id(),0}
                          }
                          } 
                        if ErrorIn(try_18) {try_17 = try_18
                        } else {
                        v_or11 = ToBoolean(OBJ(try_18))
                        if (v_or11 == CTRUE) {try_17 = EID{CTRUE.Id(),0}
                        } else { 
                          v_or11 = F__sup_equal_integer(C_compiler.Safety,2)
                          if (v_or11 == CTRUE) {try_17 = EID{CTRUE.Id(),0}
                          } else { 
                            try_17 = EID{CFALSE.Id(),0}} 
                          } 
                        }
                        } 
                      if ErrorIn(try_17) {try_16 = try_17
                      } else {
                      v_and8 = ToBoolean(OBJ(try_17))
                      if (v_and8 == CFALSE) {try_16 = EID{CFALSE.Id(),0}
                      } else { 
                        try_16 = EID{CTRUE.Id(),0}} 
                      } 
                    } 
                  }
                  } 
                if ErrorIn(try_16) {Result = try_16
                } else {
                g0219I = ToBoolean(OBJ(try_16))
                if (g0219I == CTRUE) { 
                  { var _CL_obj *Language.CallTable = Language.To_CallTable(new(Language.CallTable).Is(Language.C_Call_table))
                    _CL_obj.Selector = ToTable(x)
                    { 
                      var va_arg1 *Language.CallTable
                      var va_arg2 *ClaireAny
                      va_arg1 = _CL_obj
                      var try_25 EID
                      { var _CL_obj *Language.List = Language.To_List(new(Language.List).Is(Language.C_List))
                        { 
                          var va_arg1 *Language.Construct
                          var va_arg2 *ClaireList
                          va_arg1 = Language.To_Construct(_CL_obj.Id())
                          var try_26 EID
                          { 
                            var v_bag_arg *ClaireAny
                            try_26= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                            var try_27 EID
                            try_27 = Core.F_CALL(C_c_code,ARGS(l.At(1).ToEID(),EID{C_integer.Id(),0}))
                            if ErrorIn(try_27) {try_26 = try_27
                            } else {
                            v_bag_arg = ANY(try_27)
                            ToList(OBJ(try_26)).AddFast(v_bag_arg)
                            var try_28 EID
                            try_28 = Core.F_CALL(C_c_code,ARGS(l.At(2).ToEID(),EID{C_integer.Id(),0}))
                            if ErrorIn(try_28) {try_26 = try_28
                            } else {
                            v_bag_arg = ANY(try_28)
                            ToList(OBJ(try_26)).AddFast(v_bag_arg)}}
                            } 
                          if ErrorIn(try_26) {try_25 = try_26
                          } else {
                          va_arg2 = ToList(OBJ(try_26))
                          va_arg1.Args = va_arg2
                          try_25 = EID{va_arg2.Id(),0}
                          }
                          } 
                        if !ErrorIn(try_25) {
                        try_25 = EID{_CL_obj.Id(),0}
                        }
                        } 
                      if ErrorIn(try_25) {Result = try_25
                      } else {
                      va_arg2 = ANY(try_25)
                      va_arg1.Arg = va_arg2
                      Result = va_arg2.ToEID()
                      }
                      } 
                    if !ErrorIn(Result) {
                    { 
                      var va_arg1 *Language.CallTable
                      var va_arg2 *ClaireBoolean
                      va_arg1 = _CL_obj
                      var try_29 EID
                      { var arg_30 *ClaireBoolean
                        var try_31 EID
                        { 
                          var v_or12 *ClaireBoolean
                          
                          var try_32 EID
                          try_32 = Core.F_BELONG(ANY(Core.F_CALL(C_default,ARGS(x.ToEID()))),ANY(Core.F_CALL(C_range,ARGS(x.ToEID()))))
                          if ErrorIn(try_32) {try_31 = try_32
                          } else {
                          v_or12 = ToBoolean(OBJ(try_32))
                          if (v_or12 == CTRUE) {try_31 = EID{CTRUE.Id(),0}
                          } else { 
                            v_or12 = Equal(ANY(Core.F_CALL(C_default,ARGS(x.ToEID()))),MakeInteger(0).Id())
                            if (v_or12 == CTRUE) {try_31 = EID{CTRUE.Id(),0}
                            } else { 
                              v_or12 = Equal(p.Id(),C_get.Id())
                              if (v_or12 == CTRUE) {try_31 = EID{CTRUE.Id(),0}
                              } else { 
                                try_31 = EID{CFALSE.Id(),0}} 
                              } 
                            } 
                          }
                          } 
                        if ErrorIn(try_31) {try_29 = try_31
                        } else {
                        arg_30 = ToBoolean(OBJ(try_31))
                        try_29 = EID{arg_30.Not.Id(),0}
                        }
                        } 
                      if ErrorIn(try_29) {Result = try_29
                      } else {
                      va_arg2 = ToBoolean(OBJ(try_29))
                      va_arg1.Test = va_arg2
                      Result = EID{va_arg2.Id(),0}
                      }
                      } 
                    if !ErrorIn(Result) {
                    Result = EID{_CL_obj.Id(),0}
                    }}
                    } 
                  }  else if ((t.Included(ToType(C_array.Id())) == CTRUE) && 
                    (((p.Id() == C_nth_get.Id()) || 
                        (C_compiler.Safety >= 2)) && 
                      ((mt.Included(ToType(C_float.Id())) == CTRUE) || 
                          (Equal(Core.F__exp_type(mt,ToType(C_float.Id())).Id(),CEMPTY.Id()) == CTRUE)))) { 
                  { var arg_33 *ClaireAny
                    var try_35 EID
                    try_35 = Core.F_CALL(C_c_code,ARGS(x.ToEID(),EID{C_array.Id(),0}))
                    if ErrorIn(try_35) {Result = try_35
                    } else {
                    arg_33 = ANY(try_35)
                    { var arg_34 *ClaireAny
                      var try_36 EID
                      try_36 = Core.F_CALL(C_c_code,ARGS(l.At(1).ToEID(),EID{C_integer.Id(),0}))
                      if ErrorIn(try_36) {Result = try_36
                      } else {
                      arg_34 = ANY(try_36)
                      Result = Language.C_Call_array.Make(arg_33,arg_34,mt.Id()).ToEID()
                      }
                      } 
                    }
                    } 
                  }  else if (C_method.Id() == r.Isa.Id()) { 
                  if ((C_compiler.Optimize_ask == CTRUE) && 
                      ((t.Included(ToType(C_array.Id())) == CTRUE) || 
                          (t.Included(ToType(C_table.Id())) == CTRUE))) { 
                    F_Compile_notice_void()
                    Core.F_tformat_string(MakeString("poorly typed call: ~S\n"),3,MakeConstantList(self.Id()))
                    } 
                  { var arg_37 *ClaireList
                    var try_38 EID
                    { 
                      var v_list10 *ClaireList
                      var x *ClaireAny
                      var v_local10 *ClaireAny
                      v_list10 = self.Args
                      try_38 = EID{CreateList(ToType(CEMPTY.Id()),v_list10.Length()).Id(),0}
                      for CLcount := 0; CLcount < v_list10.Length(); CLcount++{ 
                        x = v_list10.At(CLcount)
                        var try_39 EID
                        try_39 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
                        if ErrorIn(try_39) {try_38 = try_39
                        break
                        } else {
                        v_local10 = ANY(try_39)
                        ToList(OBJ(try_38)).PutAt(CLcount,v_local10)
                        } 
                      }
                      } 
                    if ErrorIn(try_38) {Result = try_38
                    } else {
                    arg_37 = ToList(OBJ(try_38))
                    Result = F_Optimize_c_code_method_method1(ToMethod(r),self.Args,arg_37)
                    }
                    } 
                  } else {
                  { var arg_40 *ClaireList
                    var try_41 EID
                    { 
                      var v_list10 *ClaireList
                      var x *ClaireAny
                      var v_local10 *ClaireAny
                      v_list10 = self.Args
                      try_41 = EID{CreateList(ToType(CEMPTY.Id()),v_list10.Length()).Id(),0}
                      for CLcount := 0; CLcount < v_list10.Length(); CLcount++{ 
                        x = v_list10.At(CLcount)
                        var try_42 EID
                        try_42 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
                        if ErrorIn(try_42) {try_41 = try_42
                        break
                        } else {
                        v_local10 = ANY(try_42)
                        ToList(OBJ(try_41)).PutAt(CLcount,v_local10)
                        } 
                      }
                      } 
                    if ErrorIn(try_41) {Result = try_41
                    } else {
                    arg_40 = ToList(OBJ(try_41))
                    Result = F_Optimize_c_warn_property(p,self.Args,arg_40)
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
return Result} 

// The EID go function for: c_code_nth @ Call (throw: true) 
func E_Optimize_c_code_nth_Call (self EID) EID { 
return F_Optimize_c_code_nth_Call(Language.To_Call(OBJ(self)) )} 

// nth= optimization for tables
// notice that we generate updates ONLY IF the table is implemented by a list (one or two dimensions)
/* The go function for: c_code_table(self:Call) [status=1] */
func F_Optimize_c_code_table_Call (self *Language.Call) EID { 
var Result EID
{ var sp *ClaireProperty = self.Selector
  { var p *ClaireTable = ToTable(self.Args.At(0))
    { var x *ClaireAny = self.Args.At(1)
      { var y *ClaireAny = self.Args.At(2)
        if (C_OPT.ToRemove.Contain_ask(p.Id()) == CTRUE) { 
          Result = EID{CNIL.Id(),0}
          } else {
          var g0220I *ClaireBoolean
          var try_1 EID
          { 
            var v_or5 *ClaireBoolean
            
            v_or5 = Equal(sp.Id(),C_put.Id())
            if (v_or5 == CTRUE) {try_1 = EID{CTRUE.Id(),0}
            } else { 
              var try_2 EID
              { 
                var v_and7 *ClaireBoolean
                
                var try_3 EID
                { 
                  var v_or8 *ClaireBoolean
                  
                  var try_4 EID
                  { var arg_5 *ClaireType
                    var try_6 EID
                    try_6 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
                    if ErrorIn(try_6) {try_4 = try_6
                    } else {
                    arg_5 = ToType(OBJ(try_6))
                    try_4 = EID{arg_5.Included(p.Domain).Id(),0}
                    }
                    } 
                  if ErrorIn(try_4) {try_3 = try_4
                  } else {
                  v_or8 = ToBoolean(OBJ(try_4))
                  if (v_or8 == CTRUE) {try_3 = EID{CTRUE.Id(),0}
                  } else { 
                    v_or8 = F__sup_equal_integer(C_compiler.Safety,3)
                    if (v_or8 == CTRUE) {try_3 = EID{CTRUE.Id(),0}
                    } else { 
                      try_3 = EID{CFALSE.Id(),0}} 
                    } 
                  }
                  } 
                if ErrorIn(try_3) {try_2 = try_3
                } else {
                v_and7 = ToBoolean(OBJ(try_3))
                if (v_and7 == CFALSE) {try_2 = EID{CFALSE.Id(),0}
                } else { 
                  var try_7 EID
                  { 
                    var v_or9 *ClaireBoolean
                    
                    var try_8 EID
                    { var arg_9 *ClaireType
                      var try_10 EID
                      try_10 = Core.F_CALL(C_c_type,ARGS(y.ToEID()))
                      if ErrorIn(try_10) {try_8 = try_10
                      } else {
                      arg_9 = ToType(OBJ(try_10))
                      try_8 = EID{arg_9.Included(p.Range).Id(),0}
                      }
                      } 
                    if ErrorIn(try_8) {try_7 = try_8
                    } else {
                    v_or9 = ToBoolean(OBJ(try_8))
                    if (v_or9 == CTRUE) {try_7 = EID{CTRUE.Id(),0}
                    } else { 
                      v_or9 = F__sup_equal_integer(C_compiler.Safety,3)
                      if (v_or9 == CTRUE) {try_7 = EID{CTRUE.Id(),0}
                      } else { 
                        try_7 = EID{CFALSE.Id(),0}} 
                      } 
                    }
                    } 
                  if ErrorIn(try_7) {try_2 = try_7
                  } else {
                  v_and7 = ToBoolean(OBJ(try_7))
                  if (v_and7 == CFALSE) {try_2 = EID{CFALSE.Id(),0}
                  } else { 
                    try_2 = EID{CTRUE.Id(),0}} 
                  } 
                }}
                } 
              if ErrorIn(try_2) {try_1 = try_2
              } else {
              v_or5 = ToBoolean(OBJ(try_2))
              if (v_or5 == CTRUE) {try_1 = EID{CTRUE.Id(),0}
              } else { 
                try_1 = EID{CFALSE.Id(),0}} 
              } 
            }
            } 
          if ErrorIn(try_1) {Result = try_1
          } else {
          g0220I = ToBoolean(OBJ(try_1))
          if (g0220I == CTRUE) { 
            var g0221I *ClaireBoolean
            var try_11 EID
            { 
              var v_and6 *ClaireBoolean
              
              var try_12 EID
              try_12 = F_Optimize_Update_ask_relation1(ToRelation(p.Id()),x,y)
              if ErrorIn(try_12) {try_11 = try_12
              } else {
              v_and6 = ToBoolean(OBJ(try_12))
              if (v_and6 == CFALSE) {try_11 = EID{CFALSE.Id(),0}
              } else { 
                v_and6 = MakeBoolean((p.Params.Isa.IsIn(C_list) == CTRUE) || (C_integer.Id() == p.Params.Isa.Id()))
                if (v_and6 == CFALSE) {try_11 = EID{CFALSE.Id(),0}
                } else { 
                  v_and6 = F__sup_equal_integer(C_compiler.Safety,3)
                  if (v_and6 == CFALSE) {try_11 = EID{CFALSE.Id(),0}
                  } else { 
                    try_11 = EID{CTRUE.Id(),0}} 
                  } 
                } 
              }
              } 
            if ErrorIn(try_11) {Result = try_11
            } else {
            g0221I = ToBoolean(OBJ(try_11))
            if (g0221I == CTRUE) { 
              { var _Zx *ClaireAny
                var try_13 EID
                try_13 = Core.F_CALL(C_c_code,ARGS(x.ToEID(),EID{C_any.Id(),0}))
                if ErrorIn(try_13) {Result = try_13
                } else {
                _Zx = ANY(try_13)
                { var _Zy *ClaireAny
                  var try_14 EID
                  try_14 = Core.F_CALL(C_c_code,ARGS(y.ToEID(),EID{C_any.Id(),0}))
                  if ErrorIn(try_14) {Result = try_14
                  } else {
                  _Zy = ANY(try_14)
                  { var _CL_obj *Language.Update = Language.To_Update(new(Language.Update).Is(Language.C_Update))
                    _CL_obj.Selector = p.Id()
                    _CL_obj.Value = _Zy
                    _CL_obj.Arg = IfThenElse((sp.Id() == C_put.Id()),
                      C_put.Id(),
                      _Zx)
                    { 
                      var va_arg1 *Language.Update
                      var va_arg2 *ClaireAny
                      va_arg1 = _CL_obj
                      { var _CL_obj *Language.CallTable = Language.To_CallTable(new(Language.CallTable).Is(Language.C_Call_table))
                        _CL_obj.Selector = p
                        _CL_obj.Arg = _Zx
                        _CL_obj.Test = CFALSE
                        va_arg2 = _CL_obj.Id()
                        } 
                      va_arg1.ClaireVar = va_arg2
                      } 
                    Result = EID{_CL_obj.Id(),0}
                    } 
                  }
                  } 
                }
                } 
              }  else if ((sp.Id() == C_put.Id()) || 
                ((p.Inverse.Id() == CNULL) && 
                    (p.IfWrite == CNULL))) { 
              if (C_compiler.Optimize_ask == CTRUE) { 
                F_Compile_notice_void()
                Core.F_tformat_string(MakeString("poorly typed update: ~S\n"),3,MakeConstantList(self.Id()))
                } 
              Result = F_Optimize_c_code_method_method1(ToMethod(Core.F__at_property1(C_put,C_table).Id()),self.Args,MakeConstantList(C_table.Id(),C_any.Id(),C_any.Id()))
              } else {
              Result = F_Optimize_c_code_method_method1(ToMethod(Core.F__at_property1(C_nth_put,C_table).Id()),self.Args,MakeConstantList(C_table.Id(),C_any.Id(),C_any.Id()))
              } 
            }
            } else {
            Result = F_Optimize_c_code_method_method1(ToMethod(Core.F__at_property1(C_nth_equal,C_table).Id()),self.Args,MakeConstantList(C_table.Id(),C_any.Id(),C_any.Id()))
            } 
          }
          } 
        } 
      } 
    } 
  } 
return Result} 

// The EID go function for: c_code_table @ Call (throw: true) 
func E_Optimize_c_code_table_Call (self EID) EID { 
return F_Optimize_c_code_table_Call(Language.To_Call(OBJ(self)) )} 

// version for arrays (manage a nth= for array)
// we can use an Update only if no gthrow is involved in the selector (CLAIRE4)
/* The go function for: c_code_array(self:Call) [status=1] */
func F_Optimize_c_code_array_Call (self *Language.Call) EID { 
var Result EID
{ var sp *ClaireProperty = self.Selector
  { var p *ClaireAny = self.Args.At(0)
    { var tp *ClaireType
      var try_1 EID
      try_1 = Core.F_CALL(C_c_type,ARGS(p.ToEID()))
      if ErrorIn(try_1) {Result = try_1
      } else {
      tp = ToType(OBJ(try_1))
      { var mt *ClaireType = Core.F_member_type(tp)
        { var x *ClaireAny = self.Args.At(1)
          { var y *ClaireAny = self.Args.At(2)
            { var typeok *ClaireBoolean
              var try_2 EID
              { 
                var v_or7 *ClaireBoolean
                
                var try_3 EID
                { var arg_4 *ClaireType
                  var try_5 EID
                  try_5 = Core.F_CALL(C_c_type,ARGS(y.ToEID()))
                  if ErrorIn(try_5) {try_3 = try_5
                  } else {
                  arg_4 = ToType(OBJ(try_5))
                  try_3 = EID{arg_4.Included(Core.F_member_type(tp)).Id(),0}
                  }
                  } 
                if ErrorIn(try_3) {try_2 = try_3
                } else {
                v_or7 = ToBoolean(OBJ(try_3))
                if (v_or7 == CTRUE) {try_2 = EID{CTRUE.Id(),0}
                } else { 
                  v_or7 = F__sup_equal_integer(C_compiler.Safety,2)
                  if (v_or7 == CTRUE) {try_2 = EID{CTRUE.Id(),0}
                  } else { 
                    try_2 = EID{CFALSE.Id(),0}} 
                  } 
                }
                } 
              if ErrorIn(try_2) {Result = try_2
              } else {
              typeok = ToBoolean(OBJ(try_2))
              { var _Zsel *ClaireAny
                var try_6 EID
                try_6 = Core.F_CALL(C_c_code,ARGS(p.ToEID(),EID{C_array.Id(),0}))
                if ErrorIn(try_6) {Result = try_6
                } else {
                _Zsel = ANY(try_6)
                var g0222I *ClaireBoolean
                var try_7 EID
                { 
                  var v_and8 *ClaireBoolean
                  
                  v_and8 = MakeBoolean((sp.Id() == C_nth_put.Id()) || (typeok == CTRUE))
                  if (v_and8 == CFALSE) {try_7 = EID{CFALSE.Id(),0}
                  } else { 
                    v_and8 = MakeBoolean((mt.Included(ToType(C_float.Id())) == CTRUE) || (Equal(Core.F__exp_type(mt,ToType(C_float.Id())).Id(),CEMPTY.Id()) == CTRUE))
                    if (v_and8 == CFALSE) {try_7 = EID{CFALSE.Id(),0}
                    } else { 
                      var try_8 EID
                      { var arg_9 *ClaireBoolean
                        var try_10 EID
                        try_10 = F_Compile_g_throw_any(_Zsel)
                        if ErrorIn(try_10) {try_8 = try_10
                        } else {
                        arg_9 = ToBoolean(OBJ(try_10))
                        try_8 = EID{arg_9.Not.Id(),0}
                        }
                        } 
                      if ErrorIn(try_8) {try_7 = try_8
                      } else {
                      v_and8 = ToBoolean(OBJ(try_8))
                      if (v_and8 == CFALSE) {try_7 = EID{CFALSE.Id(),0}
                      } else { 
                        try_7 = EID{CTRUE.Id(),0}} 
                      } 
                    } 
                  }
                  } 
                if ErrorIn(try_7) {Result = try_7
                } else {
                g0222I = ToBoolean(OBJ(try_7))
                if (g0222I == CTRUE) { 
                  { var _Zx *ClaireAny
                    var try_11 EID
                    try_11 = Core.F_CALL(C_c_code,ARGS(x.ToEID(),EID{C_integer.Id(),0}))
                    if ErrorIn(try_11) {Result = try_11
                    } else {
                    _Zx = ANY(try_11)
                    { var _Zy *ClaireAny
                      var try_12 EID
                      { var arg_13 *ClaireClass
                        if (mt.Included(ToType(C_float.Id())) == CTRUE) { 
                          arg_13 = C_float
                          } else {
                          arg_13 = C_any
                          } 
                        try_12 = Core.F_CALL(C_c_code,ARGS(y.ToEID(),EID{arg_13.Id(),0}))
                        } 
                      if ErrorIn(try_12) {Result = try_12
                      } else {
                      _Zy = ANY(try_12)
                      { var _CL_obj *Language.Update = Language.To_Update(new(Language.Update).Is(Language.C_Update))
                        _CL_obj.Selector = p
                        _CL_obj.Value = _Zy
                        _CL_obj.Arg = C_put.Id()
                        _CL_obj.ClaireVar = Language.C_Call_array.Make(_Zsel,_Zx,mt.Id())
                        Result = EID{_CL_obj.Id(),0}
                        } 
                      }
                      } 
                    }
                    } 
                  } else {
                  if (C_compiler.Optimize_ask == CTRUE) { 
                    F_Compile_notice_void()
                    Core.F_tformat_string(MakeString("poorly typed update: ~S\n"),3,MakeConstantList(self.Id()))
                    } 
                  Result = F_Optimize_c_code_method_method1(ToMethod(Core.F__at_property1(ToProperty(IfThenElse((typeok == CTRUE),
                    C_nth_put.Id(),
                    sp.Id())),C_array).Id()),self.Args,MakeConstantList(tp.Id(),C_any.Id(),C_any.Id()))
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
return Result} 

// The EID go function for: c_code_array @ Call (throw: true) 
func E_Optimize_c_code_array_Call (self EID) EID { 
return F_Optimize_c_code_array_Call(Language.To_Call(OBJ(self)) )} 

// can we use the special UDATE form ?nth
/* The go function for: Update?(p:relation,x:any,y:any) [status=1] */
func F_Optimize_Update_ask_relation1 (p *ClaireRelation,x *ClaireAny,y *ClaireAny) EID { 
var Result EID
{ 
  var v_and0 *ClaireBoolean
  
  v_and0 = Core.F__I_equal_any(p.Id(),C_inverse.Id())
  if (v_and0 == CFALSE) {Result = EID{CFALSE.Id(),0}
  } else { 
    var try_1 EID
    { 
      var v_or2 *ClaireBoolean
      
      v_or2 = MakeBoolean((p.IfWrite != CNULL) && (p.IfWrite.Isa.IsIn(C_list) != CTRUE))
      if (v_or2 == CTRUE) {try_1 = EID{CTRUE.Id(),0}
      } else { 
        var try_2 EID
        { 
          var v_and4 *ClaireBoolean
          
          v_and4 = MakeBoolean((p.Inverse.Id() == CNULL))
          if (v_and4 == CFALSE) {try_2 = EID{CFALSE.Id(),0}
          } else { 
            if (C_table.Id() == p.Isa.Id()) { 
              { var g0223 *ClaireTable = ToTable(p.Id())
                v_and4 = Equal(C_integer.Id(),g0223.Params.Isa.Id())
                } 
              } else {
              v_and4 = CTRUE
              } 
            if (v_and4 == CFALSE) {try_2 = EID{CFALSE.Id(),0}
            } else { 
              var try_3 EID
              if (p.Store_ask == CTRUE) { 
                { 
                  var v_and8 *ClaireBoolean
                  
                  var try_4 EID
                  try_4 = F_Compile_designated_ask_any(x)
                  if ErrorIn(try_4) {try_3 = try_4
                  } else {
                  v_and8 = ToBoolean(OBJ(try_4))
                  if (v_and8 == CFALSE) {try_3 = EID{CFALSE.Id(),0}
                  } else { 
                    var try_5 EID
                    try_5 = F_Compile_designated_ask_any(y)
                    if ErrorIn(try_5) {try_3 = try_5
                    } else {
                    v_and8 = ToBoolean(OBJ(try_5))
                    if (v_and8 == CFALSE) {try_3 = EID{CFALSE.Id(),0}
                    } else { 
                      v_and8 = p.Multivalued_ask.Not
                      if (v_and8 == CFALSE) {try_3 = EID{CFALSE.Id(),0}
                      } else { 
                        var try_6 EID
                        { 
                          var v_or12 *ClaireBoolean
                          
                          var try_7 EID
                          try_7 = F_Compile_identifiable_ask_any(y)
                          if ErrorIn(try_7) {try_6 = try_7
                          } else {
                          v_or12 = ToBoolean(OBJ(try_7))
                          if (v_or12 == CTRUE) {try_6 = EID{CTRUE.Id(),0}
                          } else { 
                            var try_8 EID
                            { var arg_9 *ClaireType
                              var try_10 EID
                              try_10 = Core.F_CALL(C_c_type,ARGS(y.ToEID()))
                              if ErrorIn(try_10) {try_8 = try_10
                              } else {
                              arg_9 = ToType(OBJ(try_10))
                              try_8 = EID{arg_9.Included(ToType(C_float.Id())).Id(),0}
                              }
                              } 
                            if ErrorIn(try_8) {try_6 = try_8
                            } else {
                            v_or12 = ToBoolean(OBJ(try_8))
                            if (v_or12 == CTRUE) {try_6 = EID{CTRUE.Id(),0}
                            } else { 
                              try_6 = EID{CFALSE.Id(),0}} 
                            } 
                          }}
                          } 
                        if ErrorIn(try_6) {try_3 = try_6
                        } else {
                        v_and8 = ToBoolean(OBJ(try_6))
                        if (v_and8 == CFALSE) {try_3 = EID{CFALSE.Id(),0}
                        } else { 
                          try_3 = EID{CTRUE.Id(),0}} 
                        } 
                      } 
                    } 
                  }}}
                  } 
                } else {
                try_3 = EID{CTRUE.Id(),0}
                } 
              if ErrorIn(try_3) {try_2 = try_3
              } else {
              v_and4 = ToBoolean(OBJ(try_3))
              if (v_and4 == CFALSE) {try_2 = EID{CFALSE.Id(),0}
              } else { 
                try_2 = EID{CTRUE.Id(),0}} 
              } 
            } 
          }
          } 
        if ErrorIn(try_2) {try_1 = try_2
        } else {
        v_or2 = ToBoolean(OBJ(try_2))
        if (v_or2 == CTRUE) {try_1 = EID{CTRUE.Id(),0}
        } else { 
          try_1 = EID{CFALSE.Id(),0}} 
        } 
      }
      } 
    if ErrorIn(try_1) {Result = try_1
    } else {
    v_and0 = ToBoolean(OBJ(try_1))
    if (v_and0 == CFALSE) {Result = EID{CFALSE.Id(),0}
    } else { 
      Result = EID{CTRUE.Id(),0}} 
    } 
  }
  } 
return Result} 

// The EID go function for: Update? @ list<type_expression>(relation, any, any) (throw: true) 
func E_Optimize_Update_ask_relation1 (p EID,x EID,y EID) EID { 
return F_Optimize_Update_ask_relation1(ToRelation(OBJ(p)),ANY(x),ANY(y) )} 

// we do not use an Update form for add
/* The go function for: Update?(p:relation,s:relation) [status=0] */
func F_Optimize_Update_ask_relation2 (p *ClaireRelation,s *ClaireRelation) *ClaireBoolean { 
return  MakeBoolean((p.IfWrite != CNULL) && (p.IfWrite.Isa.IsIn(C_list) != CTRUE) && (s.Id() == C_add.Id()))
} 

// The EID go function for: Update? @ list<type_expression>(relation, relation) (throw: false) 
func E_Optimize_Update_ask_relation2 (p EID,s EID) EID { 
return EID{F_Optimize_Update_ask_relation2(ToRelation(OBJ(p)),ToRelation(OBJ(s)) ).Id(),0}} 

// Update returns the value .. <yc:0.01 -> needed in CLAIRE 2.4 !!!>
/* The go function for: c_type(self:Update) [status=0] */
func F_c_type_Update (self *Language.Update) *ClaireType { 
return  ToType(C_void.Id())
} 

// The EID go function for: c_type @ Update (throw: false) 
func E_c_type_Update (self EID) EID { 
return EID{F_c_type_Update(Language.To_Update(OBJ(self)) ).Id(),0}} 

// in CLAIRE4 we isolate this case (call the if-write demon) because it may produce an error
/* The go function for: Compile/update_write?(self:Update) [status=0] */
func F_Compile_update_write_ask_Update (self *Language.Update) *ClaireBoolean { 
var Result *ClaireBoolean
{ var p *ClaireAny = self.Selector
  { var a *ClaireAny = self.Arg
    if (p.Isa.IsIn(C_relation) == CTRUE) { 
      { var g0225 *ClaireRelation = ToRelation(p)
        Result = MakeBoolean((g0225.IfWrite != CNULL) && (a != C_put.Id()) && (a != Core.C_put_store.Id()))
        } 
      } else {
      Result = CFALSE
      } 
    } 
  } 
return Result} 

// The EID go function for: Compile/update_write? @ Update (throw: false) 
func E_Compile_update_write_ask_Update (self EID) EID { 
return EID{F_Compile_update_write_ask_Update(Language.To_Update(OBJ(self)) ).Id(),0}} 

// ******************************************************************
// *    Part 3: Method optimizer                                    *
// ******************************************************************
// a basic method is assumed to be compiled if it is in the right module
/* The go function for: c_code_method(self:method,l:list,%type:list) [status=1] */
func F_Optimize_c_code_method_method1 (self *ClaireMethod,l *ClaireList,_Ztype *ClaireList) EID { 
var Result EID
{ var arg_1 *ClaireClass
  var try_2 EID
  try_2 = F_Optimize_c_srange_method(self)
  if ErrorIn(try_2) {Result = try_2
  } else {
  arg_1 = ToClass(OBJ(try_2))
  Result = F_Optimize_c_code_method_method2(self,l,_Ztype,arg_1)
  }
  } 
return Result} 

// The EID go function for: c_code_method @ list<type_expression>(method, list, list) (throw: true) 
func E_Optimize_c_code_method_method1 (self EID,l EID,_Ztype EID) EID { 
return F_Optimize_c_code_method_method1(ToMethod(OBJ(self)),ToList(OBJ(l)),ToList(OBJ(_Ztype)) )} 

/* The go function for: c_code_method(self:method,l:list,%type:list,sx:class) [status=1] */
func F_Optimize_c_code_method_method2 (self *ClaireMethod,l *ClaireList,_Ztype *ClaireList,sx *ClaireClass) EID { 
var Result EID
if ((self.Module_I.Id() != C_claire.Id()) || 
    ((C_compiler.Safety >= 2) || 
      (self.Functional.Id() != CNULL))) { 
  { var ld *ClaireList = self.Domain
    { var n int = ld.Length()
      if (n != l.Length()) { 
        { var arg_1 *ClaireList
          { var i_bag *ClaireList = ToType(CEMPTY.Id()).EmptyList()
            { var i int = 1
              { var g0227 int = (n-1)
                for (i <= g0227) { 
                  i_bag.AddFast(l.At(i-1))
                  i = (i+1)
                  } 
                } 
              } 
            arg_1 = i_bag
            } 
          { var arg_2 *ClaireList
            { var i_bag *ClaireList = ToType(CEMPTY.Id()).EmptyList()
              { var i int = n
                { var g0228 int = l.Length()
                  for (i <= g0228) { 
                    i_bag.AddFast(l.At(i-1))
                    i = (i+1)
                    } 
                  } 
                } 
              arg_2 = i_bag
              } 
            l = arg_1.AddFast(arg_2.Id())
            } 
          } 
        } 
      var g0230I *ClaireBoolean
      var try_3 EID
      { 
        var v_and3 *ClaireBoolean
        
        v_and3 = self.Inline_ask
        if (v_and3 == CFALSE) {try_3 = EID{CFALSE.Id(),0}
        } else { 
          var try_4 EID
          try_4 = F_Optimize_c_inline_ask_method(self,l)
          if ErrorIn(try_4) {try_3 = try_4
          } else {
          v_and3 = ToBoolean(OBJ(try_4))
          if (v_and3 == CFALSE) {try_3 = EID{CFALSE.Id(),0}
          } else { 
            try_3 = EID{CTRUE.Id(),0}} 
          } 
        }
        } 
      if ErrorIn(try_3) {Result = try_3
      } else {
      g0230I = ToBoolean(OBJ(try_3))
      if (g0230I == CTRUE) { 
        Result = F_Optimize_c_inline_method1(self,l,sx)
        } else {
        { var arg_5 *ClaireList
          var try_6 EID
          { var i_bag *ClaireList = ToType(CEMPTY.Id()).EmptyList()
            { var i int = 1
              { var g0229 int = n
                try_6= EID{CFALSE.Id(),0}
                for (i <= g0229) { 
                  var loop_7 EID
                  _ = loop_7
                  { 
                  { var arg_8 *ClaireAny
                    var try_9 EID
                    try_9 = F_Compile_c_strict_code_any(l.At(i-1),F_Compile_psort_any(ld.ValuesO()[i-1]))
                    if ErrorIn(try_9) {loop_7 = try_9
                    } else {
                    arg_8 = ANY(try_9)
                    loop_7 = EID{i_bag.AddFast(arg_8).Id(),0}
                    }
                    } 
                  if ErrorIn(loop_7) {try_6 = loop_7
                  break
                  } else {
                  i = (i+1)
                  }
                  } 
                }
                } 
              } 
            if !ErrorIn(try_6) {
            try_6 = EID{i_bag.Id(),0}
            }
            } 
          if ErrorIn(try_6) {Result = try_6
          } else {
          arg_5 = ToList(OBJ(try_6))
          Result = F_Optimize_Call_method_I_method(self,arg_5).ToEID()
          }
          } 
        } 
      }
      } 
    } 
  } else {
  if (C_compiler.Optimize_ask == CTRUE) { 
    F_Compile_notice_void()
    Core.F_tformat_string(MakeString("poorly typed update: ~S\n"),3,MakeConstantList(self.Id()))
    } 
  Result = F_Optimize_open_message_property(self.Selector,l)
  } 
return Result} 

// The EID go function for: c_code_method @ list<type_expression>(method, list, list, class) (throw: true) 
func E_Optimize_c_code_method_method2 (self EID,l EID,_Ztype EID,sx EID) EID { 
return F_Optimize_c_code_method_method2(ToMethod(OBJ(self)),
  ToList(OBJ(l)),
  ToList(OBJ(_Ztype)),
  ToClass(OBJ(sx)) )} 

// the code to be produced for a method
/* The go function for: Call_method!(self:method,%code:list) [status=0] */
func F_Optimize_Call_method_I_method (self *ClaireMethod,_Zcode *ClaireList) *ClaireAny { 
var Result *ClaireAny
if (F_Optimize_legal_ask_module(self.Module_I,self.Id()) != CTRUE) { 
  Core.F_tformat_string(MakeString("in call ~S~S\n"),0,MakeConstantList(self.Selector.Id(),_Zcode.Id()))
  } 
if (_Zcode.Length() == 1) { 
  { var _CL_obj *Language.CallMethod1 = Language.To_CallMethod1(new(Language.CallMethod1).Is(Language.C_Call_method1))
    _CL_obj.Arg = self
    _CL_obj.Args = _Zcode
    Result = _CL_obj.Id()
    } 
  }  else if (_Zcode.Length() == 2) { 
  { var _CL_obj *Language.CallMethod2 = Language.To_CallMethod2(new(Language.CallMethod2).Is(Language.C_Call_method2))
    _CL_obj.Arg = self
    _CL_obj.Args = _Zcode
    Result = _CL_obj.Id()
    } 
  } else {
  { var _CL_obj *Language.CallMethod = Language.To_CallMethod(new(Language.CallMethod).Is(Language.C_Call_method))
    _CL_obj.Arg = self
    _CL_obj.Args = _Zcode
    Result = _CL_obj.Id()
    } 
  } 
return Result} 

// The EID go function for: Call_method! @ method (throw: false) 
func E_Optimize_Call_method_I_method (self EID,_Zcode EID) EID { 
return F_Optimize_Call_method_I_method(ToMethod(OBJ(self)),ToList(OBJ(_Zcode)) ).ToEID()} 

// a call_method or a call external has an obvious type (we do not need to do
// better ?)
/* The go function for: c_type(self:Call_method) [status=1] */
func F_c_type_Call_method (self *Language.CallMethod) EID { 
var Result EID
{ var arg_1 *ClaireList
  var try_2 EID
  { 
    var v_list1 *ClaireList
    var x *ClaireAny
    var v_local1 *ClaireAny
    v_list1 = self.Args
    try_2 = EID{CreateList(ToType(CEMPTY.Id()),v_list1.Length()).Id(),0}
    for CLcount := 0; CLcount < v_list1.Length(); CLcount++{ 
      x = v_list1.At(CLcount)
      var try_3 EID
      try_3 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
      if ErrorIn(try_3) {try_2 = try_3
      break
      } else {
      v_local1 = ANY(try_3)
      ToList(OBJ(try_2)).PutAt(CLcount,v_local1)
      } 
    }
    } 
  if ErrorIn(try_2) {Result = try_2
  } else {
  arg_1 = ToList(OBJ(try_2))
  Result = F_Optimize_use_range_method(self.Arg,arg_1)
  }
  } 
return Result} 

// The EID go function for: c_type @ Call_method (throw: true) 
func E_c_type_Call_method (self EID) EID { 
return F_c_type_Call_method(Language.To_CallMethod(OBJ(self)) )} 

// a call_method is already compiled, but in CLAIRE4 it may be obtained with JITO so we optimize the 
/* The go function for: c_code(self:Call_method) [status=1] */
func F_c_code_Call_method (self *Language.CallMethod) EID { 
var Result EID
{ var m *ClaireMethod = self.Arg
  { var ld *ClaireList = m.Domain
    { var n int = Reader.F_min_integer(self.Args.Length(),ld.Length())
      { var arg_1 *ClaireList
        var try_2 EID
        { var i_bag *ClaireList = ToType(CEMPTY.Id()).EmptyList()
          { var i int = 1
            { var g0231 int = n
              try_2= EID{CFALSE.Id(),0}
              for (i <= g0231) { 
                var loop_3 EID
                _ = loop_3
                { 
                { var arg_4 *ClaireAny
                  var try_5 EID
                  try_5 = F_Compile_c_strict_code_any(self.Args.At(i-1),F_Compile_psort_any(ld.ValuesO()[i-1]))
                  if ErrorIn(try_5) {loop_3 = try_5
                  } else {
                  arg_4 = ANY(try_5)
                  loop_3 = EID{i_bag.AddFast(arg_4).Id(),0}
                  }
                  } 
                if ErrorIn(loop_3) {try_2 = loop_3
                break
                } else {
                i = (i+1)
                }
                } 
              }
              } 
            } 
          if !ErrorIn(try_2) {
          try_2 = EID{i_bag.Id(),0}
          }
          } 
        if ErrorIn(try_2) {Result = try_2
        } else {
        arg_1 = ToList(OBJ(try_2))
        Result = F_Optimize_Call_method_I_method(m,arg_1).ToEID()
        }
        } 
      } 
    } 
  } 
return Result} 

// The EID go function for: c_code @ Call_method (throw: true) 
func E_c_code_Call_method (self EID) EID { 
return F_c_code_Call_method(Language.To_CallMethod(OBJ(self)) )} 

// gets the associated function if it exists and create one otherwise
/* The go function for: Compile/functional!(self:method) [status=1] */
func F_Compile_functional_I_method (self *ClaireMethod) EID { 
var Result EID
{ var f *ClaireAny = Core.F_get_property(C_functional,ToObject(self.Id()))
  { var p *ClaireProperty = self.Selector
    if (C_function.Id() == f.Isa.Id()) { 
      { var g0232 *ClaireFunction = ToFunction(f)
        Result = EID{g0232.Id(),0}
        } 
      } else {
      { var arg_1 *ClaireAny
        var try_2 EID
        try_2 = Core.F_CALL(C_Compile_function_name,ARGS(EID{p.Id(),0},EID{self.Domain.Id(),0},f.ToEID()))
        if ErrorIn(try_2) {Result = try_2
        } else {
        arg_1 = ANY(try_2)
        Result = F_make_function_string(ToString(arg_1)).ToEID()
        }
        } 
      } 
    } 
  } 
return Result} 

// The EID go function for: Compile/functional! @ method (throw: true) 
func E_Compile_functional_I_method (self EID) EID { 
return F_Compile_functional_I_method(ToMethod(OBJ(self)) )} 

// second-order types for better safety or optimization -------------------------------
/* The go function for: nth_type_check(tl:type,ti:type,tx:type) [status=0] */
func F_Optimize_nth_type_check_type (tl *ClaireType,ti *ClaireType,tx *ClaireType) *ClaireAny { 
if (tx.Included(Core.F_member_type(tl)) != CTRUE) { 
  F_Compile_warn_void()
  Core.F_tformat_string(MakeString("unsafe update on bag: type ~S into ~S [252]\n"),1,MakeConstantList(tx.Id(),tl.Id()))
  } 
return  tx.Id()
} 

// The EID go function for: nth_type_check @ type (throw: false) 
func E_Optimize_nth_type_check_type (tl EID,ti EID,tx EID) EID { 
return F_Optimize_nth_type_check_type(ToType(OBJ(tl)),ToType(OBJ(ti)),ToType(OBJ(tx)) ).ToEID()} 

// ******************************************************************
// *    Part 5: inline methods                                      *
// ******************************************************************
// macro expansion for inline method ?
// we check that it is a good idea
/* The go function for: c_inline?(self:method,l:list) [status=1] */
func F_Optimize_c_inline_ask_method (self *ClaireMethod,l *ClaireList) EID { 
var Result EID
{ var f *ClaireLambda = self.Formula
  { var la *ClaireList = f.Vars
    _ = la
    { var x *ClaireAny = f.Body
      { var n int = 1
        { var arg_1 *ClaireAny
          var try_2 EID
          { 
            var v *ClaireAny
            _ = v
            try_2= EID{CFALSE.Id(),0}
            var v_support *ClaireList
            v_support = f.Vars
            v_len := v_support.Length()
            for i_it := 0; i_it < v_len; i_it++ { 
              v = v_support.At(i_it)
              var loop_3 EID
              _ = loop_3
              var g0234I *ClaireBoolean
              var try_4 EID
              { 
                var v_and7 *ClaireBoolean
                
                v_and7 = Core.F__sup_integer(Language.F_occurrence_any(x,To_Variable(v)),1)
                if (v_and7 == CFALSE) {try_4 = EID{CFALSE.Id(),0}
                } else { 
                  var try_5 EID
                  { var arg_6 *ClaireBoolean
                    var try_7 EID
                    try_7 = F_Compile_designated_ask_any(l.At(n-1))
                    if ErrorIn(try_7) {try_5 = try_7
                    } else {
                    arg_6 = ToBoolean(OBJ(try_7))
                    try_5 = EID{arg_6.Not.Id(),0}
                    }
                    } 
                  if ErrorIn(try_5) {try_4 = try_5
                  } else {
                  v_and7 = ToBoolean(OBJ(try_5))
                  if (v_and7 == CFALSE) {try_4 = EID{CFALSE.Id(),0}
                  } else { 
                    v_and7 = ANY(Core.F_CALL(C_range,ARGS(v.ToEID()))).Isa.IsIn(C_Optimize_Pattern).Not
                    if (v_and7 == CFALSE) {try_4 = EID{CFALSE.Id(),0}
                    } else { 
                      try_4 = EID{CTRUE.Id(),0}} 
                    } 
                  } 
                }
                } 
              if ErrorIn(try_4) {loop_3 = try_4
              } else {
              g0234I = ToBoolean(OBJ(try_4))
              if (g0234I == CTRUE) { 
                try_2 = EID{CTRUE.Id(),0}
                break
                } else {
                n = (n+1)
                loop_3 = EID{C__INT,IVAL(n)}
                } 
              }
              if ErrorIn(loop_3) {try_2 = loop_3
              break
              } else {
              }
              } 
            } 
          if ErrorIn(try_2) {Result = try_2
          } else {
          arg_1 = ANY(try_2)
          Result = EID{Core.F_not_any(arg_1).Id(),0}
          }
          } 
        } 
      } 
    } 
  } 
return Result} 

// The EID go function for: c_inline? @ method (throw: true) 
func E_Optimize_c_inline_ask_method (self EID,l EID) EID { 
return F_Optimize_c_inline_ask_method(ToMethod(OBJ(self)),ToList(OBJ(l)) )} 

// checks if a special optimization restriction can be used (with patterns)
/* The go function for: inline_optimize?(self:Call) [status=1] */
func F_Optimize_inline_optimize_ask_Call (self *Language.Call) EID { 
var Result EID
{ var l *ClaireList = self.Args
  { var m *ClaireAny
    { var arg_1 *ClaireList
      { 
        var v_list3 *ClaireList
        var x *ClaireAny
        var v_local3 *ClaireAny
        v_list3 = l
        arg_1 = CreateList(ToType(CEMPTY.Id()),v_list3.Length())
        for CLcount := 0; CLcount < v_list3.Length(); CLcount++{ 
          x = v_list3.At(CLcount)
          v_local3 = MakeConstantSet(x).Id()
          arg_1.PutAt(CLcount,v_local3)
          } 
        } 
      m = F_Optimize_restriction_I_property(self.Selector,arg_1,CTRUE)
      } 
    if (C_method.Id() == m.Isa.Id()) { 
      { var g0235 *ClaireMethod = ToMethod(m)
        var g0237I *ClaireBoolean
        var try_2 EID
        { 
          var v_and4 *ClaireBoolean
          
          v_and4 = g0235.Inline_ask
          if (v_and4 == CFALSE) {try_2 = EID{CFALSE.Id(),0}
          } else { 
            { var arg_3 *ClaireAny
              { 
                var s *ClaireTypeExpression
                _ = s
                var s_iter *ClaireAny
                arg_3= CFALSE.Id()
                for _,s_iter = range(g0235.Domain.ValuesO()){ 
                  s = ToTypeExpression(s_iter)
                  if (s.Isa.IsIn(C_Optimize_Pattern) == CTRUE) { 
                    arg_3 = CTRUE.Id()
                    break
                    } 
                  } 
                } 
              v_and4 = F_boolean_I_any(arg_3)
              } 
            if (v_and4 == CFALSE) {try_2 = EID{CFALSE.Id(),0}
            } else { 
              var try_4 EID
              try_4 = F_Optimize_c_inline_ask_method(g0235,l)
              if ErrorIn(try_4) {try_2 = try_4
              } else {
              v_and4 = ToBoolean(OBJ(try_4))
              if (v_and4 == CFALSE) {try_2 = EID{CFALSE.Id(),0}
              } else { 
                try_2 = EID{CTRUE.Id(),0}} 
              } 
            } 
          }
          } 
        if ErrorIn(try_2) {Result = try_2
        } else {
        g0237I = ToBoolean(OBJ(try_2))
        if (g0237I == CTRUE) { 
          Result = EID{g0235.Id(),0}
          } else {
          Result = EID{CFALSE.Id(),0}
          } 
        }
        } 
      } else {
      Result = EID{CFALSE.Id(),0}
      } 
    } 
  } 
return Result} 

// The EID go function for: inline_optimize? @ Call (throw: true) 
func E_Optimize_inline_optimize_ask_Call (self EID) EID { 
return F_Optimize_inline_optimize_ask_Call(Language.To_Call(OBJ(self)) )} 

// eof