/***** CLAIRE Compilation of file /Users/ycaseau/Dropbox/src/clairev4.03/src/compile/ocall.cl 
         [version 4.0.03 / safety 5] Monday 12-27-2021 10:35:27 *****/

package Optimize
import (_ "fmt"
	. "Kernel"
	"Core"
	"Language"
	"Reader"
)

//-------- dumb function to prevent import errors --------
func import_g0164() { 
    _ = Core.It
    _ = Language.It
    _ = Reader.It
    } 
  
  
//+-------------------------------------------------------------+
//| CLAIRE                                                      |
//| ocall.cl                                                    |
//| Copyright (C) 1994 - 2021 Yves Caseau. All Rights Reserved  |
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
/* {1} The go function for: restriction!(self:property,l:list,mode:boolean) [status=0] */
func F_Optimize_restriction_I_property (self *ClaireProperty ,l *ClaireList ,mode *ClaireBoolean ) *ClaireAny  { 
    { var i int  = 1
      { var g0165 int  = l.Length()
        _ = g0165
        for (i <= g0165) { 
          /* While stat, v:"Unused" loop:false */
          ToArray(l.Id()).NthPut(i,F_Optimize_ptype_type(ToType(l.At(i-1))).Id())
          i = (i+1)
          /* try?:false, v2:"v_while4" loop will be:tuple("Unused", void) */
          } 
        } 
      } 
    return  F_Optimize_restriction_I_list(self.Definition,l,mode)
    } 
  
// The EID go function for: restriction! @ property (throw: false) 
func E_Optimize_restriction_I_property (self EID,l EID,mode EID) EID { 
    return F_Optimize_restriction_I_property(ToProperty(OBJ(self)),ToList(OBJ(l)),ToBoolean(OBJ(mode)) ).ToEID()} 
  
// finds a suitable restriction in lr. Returns a restriction for a match,
// list(r) for a possible match (unique), () for no match and ambiguous
// otherwise
// CLAIRE4 : we define "open required" based on the property.
/* {1} The go function for: restriction!(lr:list,l:list,mode:boolean) [status=0] */
func F_Optimize_restriction_I_list (lr *ClaireList ,l *ClaireList ,mode *ClaireBoolean ) *ClaireAny  { 
    // procedure body with s = any 
var Result *ClaireAny  
    
    { var open_required *ClaireBoolean   = MakeBoolean((lr.Length() > 0) && ANY(Core.F_CALL(C_selector,ARGS(lr.At(1-1).ToEID()))).IsInt(3))
      _ = open_required
      { var rep *ClaireAny   = CEMPTY.Id()
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
            if ((F_boolean_I_any(rep).Id() != CTRUE.Id()) && 
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
/* {1} The go function for: findr(p:property,l:list) [status=1] */
func F_findr_property (p *ClaireProperty ,l *ClaireList ) EID { 
    var Result EID 
    { var lr *ClaireList   = p.Definition
      _ = lr
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
          /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
          loop_1 = Core.F_print_any(r.Id())
          /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
          if ErrorIn(loop_1) {Result = loop_1
          break
          } else {
          PRINC(") with ")
          /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
          loop_1 = Core.F_print_any(l.Id())
          /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
          if ErrorIn(loop_1) {Result = loop_1
          break
          } else {
          PRINC(" -> ")
          /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
          loop_1 = Core.F_print_any(Core.F_tmatch_ask_list(l,r.Domain).Id())
          /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
          if ErrorIn(loop_1) {Result = loop_1
          break
          } else {
          PRINC(", intersection:")
          /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
          loop_1 = Core.F_print_any(Core.F__exp_list(r.Domain,l).Id())
          /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
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
  
// TO REMOVE DEBUG tmatch
/* {1} The go function for: dmatch?(l:list,l2:list) [status=0] */
func F_dmatch_ask_list (l *ClaireList ,l2 *ClaireList ) *ClaireBoolean  { 
    // procedure body with s = boolean 
var Result *ClaireBoolean  
    { var x int  = l2.Length()
      { var z int  = l.Length()
        if ((z != x) && 
            ((l2.At(x-1) != C_listargs.Id()) || 
                (z < (x-1)))) { 
          Result = CFALSE
          } else {
          { var arg_1 *ClaireAny  
            _ = arg_1
            { var i int  = 1
              { var g0166 int  = x
                _ = g0166
                arg_1= CFALSE.Id()
                for (i <= g0166) { 
                  /* While stat, v:"arg_1" loop:false */
                  if ((i == x) && 
                      (l2.At(i-1) == C_listargs.Id())) { 
                    arg_1 = CFALSE.Id()
                    break
                    }  else if (Core.F_tmatch_ask_any(l.At(i-1),l2.At(i-1),l) != CTRUE) { 
                    Core.F_tformat_string(MakeString("tmatch?(~S,~S) failed \n"),0,MakeConstantList(l.At(i-1),l2.At(i-1)))
                    arg_1 = CTRUE.Id()
                    break
                    } 
                  i = (i+1)
                  /* try?:false, v2:"v_while8" loop will be:tuple("arg_1", any) */
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
  
// special version for Super, which only looks at methods with domains
// bigger than c
/* {1} The go function for: restriction!(c:class,lr:list,l:list) [status=0] */
func F_Optimize_restriction_I_class (c *ClaireClass ,lr *ClaireList ,l *ClaireList ) *ClaireAny  { 
    // procedure body with s = any 
var Result *ClaireAny  
    if (C_compiler.Safety >= 2) { 
      ToArray(l.Id()).NthPut(1,Core.F__exp_type(ToType(c.Id()),ToType(l.At(1-1))).Id())
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
        if (ToType(c.Id()).Included(ToType(r.Domain.ValuesO()[1-1])) == CTRUE) { 
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
/* {1} The go function for: use_range(self:method,%l:list) [status=1] */
func F_Optimize_use_range_method (self *ClaireMethod ,_Zl *ClaireList ) EID { 
    var Result EID 
    if ((self.Inline_ask == CTRUE) && 
        (self.Typing == CNULL)) { 
      { var lv *ClaireList   = self.Formula.Vars
        { var _Zt *ClaireType   = ToType(C_any.Id())
          { var _Zl2 *ClaireList  
            _ = _Zl2
            { 
              var v_list6 *ClaireList  
              var v *ClaireAny  
              var v_local6 *ClaireAny  
              v_list6 = lv
              _Zl2 = CreateList(ToType(CEMPTY.Id()),v_list6.Length())
              for CLcount := 0; CLcount < v_list6.Length(); CLcount++{ 
                v = v_list6.At(CLcount)
                v_local6 = ANY(Core.F_CALL(C_range,ARGS(v.ToEID())))
                _Zl2.PutAt(CLcount,v_local6)
                } 
              } 
            /*g_try(v2:"Result",loop:true) */
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
                /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
                loop_1 = Core.F_put_property2(C_range,ToObject(v),_Zl.At((INT(Core.F_CALL(C_mClaire_index,ARGS(v.ToEID())))+1)-1))
                /* ERROR PROTECTION INSERTED (loop_1-Result) */
                if ErrorIn(loop_1) {Result = loop_1
                break
                } else {
                }
                } 
              } 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            var try_2 EID 
            /*g_try(v2:"try_2",loop:true) */
            try_2 = Core.F_CALL(C_c_type,ARGS(self.Formula.Body.ToEID()))
            /* ERROR PROTECTION INSERTED (_Zt-Result) */
            if ErrorIn(try_2) {Result = try_2
            } else {
            _Zt = ToType(OBJ(try_2))
            Result = EID{_Zt.Id(),0}
            /*g_try(v2:"Result",loop:true) */
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
                /*g_try(v2:"loop_3",loop:tuple("Result", EID)) */
                loop_3 = Core.F_put_property2(C_range,ToObject(v),_Zl2.At((INT(Core.F_CALL(C_mClaire_index,ARGS(v.ToEID())))+1)-1))
                /* ERROR PROTECTION INSERTED (loop_3-Result) */
                if ErrorIn(loop_3) {Result = loop_3
                break
                } else {
                }
                } 
              } 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            /*g_try(v2:"Result",loop:true) */
            if (self.Range.Isa.IsIn(C_type) == CTRUE) { 
              var try_4 EID 
              /*g_try(v2:"try_4",loop:false) */
              try_4 = Core.F_CALL(ToProperty(C__exp.Id()),ARGS(EID{_Zt.Id(),0},EID{self.Range.Id(),0}))
              /* ERROR PROTECTION INSERTED (_Zt-Result) */
              if ErrorIn(try_4) {Result = try_4
              } else {
              _Zt = ToType(OBJ(try_4))
              Result = EID{_Zt.Id(),0}
              }
              } else {
              Result = EID{CFALSE.Id(),0}
              } 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            /*g_try(v2:"Result",loop:true) */
            if (F_boolean_I_any(_Zt.Id()).Id() != CTRUE.Id()) { 
              { var arg_5 *ClaireType  
                _ = arg_5
                var try_6 EID 
                /*g_try(v2:"try_6",loop:false) */
                try_6 = Core.F_CALL(C_c_type,ARGS(self.Formula.Body.ToEID()))
                /* ERROR PROTECTION INSERTED (arg_5-Result) */
                if ErrorIn(try_6) {Result = try_6
                } else {
                arg_5 = ToType(OBJ(try_6))
                Result = Core.F_CALL(C_Compile_Cerror,ARGS(EID{MakeString("[207] inline ~S: range ~S is incompatible with ~S (inferred)").Id(),0},
                  EID{self.Id(),0},
                  EID{self.Range.Id(),0},
                  EID{arg_5.Id(),0}))
                }
                } 
              } else {
              Result = EID{CFALSE.Id(),0}
              } 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            Result = EID{_Zt.Id(),0}
            }}}}}
            } 
          } 
        } 
      } else {
      
      { var f *ClaireAny   = self.Typing
        { var _Zl2 *ClaireList  
          { 
            var v_list5 *ClaireList  
            var u *ClaireAny  
            var v_local5 *ClaireAny  
            v_list5 = _Zl
            _Zl2 = CreateList(ToType(C_type.Id()),v_list5.Length())
            for CLcount := 0; CLcount < v_list5.Length(); CLcount++{ 
              u = v_list5.At(CLcount)
              v_local5 = F_Optimize_ptype_type(ToType(u)).Id()
              _Zl2.PutAt(CLcount,v_local5)
              } 
            } 
          { var _Zt1 *ClaireType   = self.Range
            { var _Zt2 *ClaireType  
              _ = _Zt2
              { 
                var _Zt2_H EID 
                h_index := ClEnv.Index
                h_base := ClEnv.Base
                if (f.Isa.IsIn(C_lambda) == CTRUE) { 
                  { var g0167 *ClaireLambda   = ToLambda(f)
                    _ = g0167
                    _Zt2_H = Language.F_apply_lambda(g0167,_Zl2)
                    } 
                  }  else if (f.Isa.IsIn(C_property) == CTRUE) { 
                  { var g0168 *ClaireProperty   = ToProperty(f)
                    _ = g0168
                    _Zt2_H = Core.F_apply_property(g0168,_Zl2)
                    } 
                  }  else if (C_function.Id() == f.Isa.Id()) { 
                  { var g0169 *ClaireFunction   = ToFunction(f)
                    _ = g0169
                    _Zt2_H = F_apply_function(g0169,_Zl2)
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
// It follows the stucture of the evaluator (self_eval)
// optimize is the distributed compiling method equivalent to the
// evaluation "behave" method
/* {1} The go function for: c_type(self:Call) [status=1] */
func F_c_type_Call (self *Language.Call ) EID { 
    var Result EID 
    if (self.Selector.Id() == Language.C_function_I.Id()) { 
      Result = EID{C_function.Id(),0}
      } else {
      { var s *ClaireProperty   = self.Selector
        { var l *ClaireList   = self.Args
          { var _Ztype *ClaireList  
            var try_1 EID 
            /*g_try(v2:"try_1",loop:false) */
            { 
              var v_list6 *ClaireList  
              var x *ClaireAny  
              var v_local6 *ClaireAny  
              v_list6 = l
              try_1 = EID{CreateList(ToType(CEMPTY.Id()),v_list6.Length()).Id(),0}
              for CLcount := 0; CLcount < v_list6.Length(); CLcount++{ 
                x = v_list6.At(CLcount)
                var try_2 EID 
                /*g_try(v2:"try_2",loop:tuple("try_1", EID)) */
                try_2 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
                /* ERROR PROTECTION INSERTED (v_local6-try_1) */
                if ErrorIn(try_2) {try_1 = try_2
                break
                } else {
                v_local6 = ANY(try_2)
                ToList(OBJ(try_1)).PutAt(CLcount,v_local6)
                } 
              }
              } 
            /* ERROR PROTECTION INSERTED (_Ztype-Result) */
            if ErrorIn(try_1) {Result = try_1
            } else {
            _Ztype = ToList(OBJ(try_1))
            if (s.Id() == C_safe.Id()) { 
              Result = _Ztype.At(1-1).ToEID()
              }  else if ((s.Id() == Core.C_externC.Id()) && 
                ((l.Length() == 2) && 
                  (C_class.Id() == l.At(2-1).Isa.Id()))) { 
              Result = l.At(2-1).ToEID()
              }  else if ((s.Id() == C_new.Id()) && 
                (C_class.Id() == l.At(1-1).Isa.Id())) { 
              Result = l.At(1-1).ToEID()
              }  else if ((s.Id() == Core.C_check_in.Id()) && 
                (l.At(2-1).Isa.IsIn(C_type) == CTRUE)) { 
              Result = l.At(2-1).ToEID()
              }  else if ((s.Id() == C_nth.Id()) && 
                (ToType(_Ztype.At(1-1)).Included(ToType(C_array.Id())) == CTRUE)) { 
              if (Core.F_member_type(ToType(_Ztype.At(1-1))).Included(ToType(C_float.Id())) == CTRUE) { 
                Result = EID{C_float.Id(),0}
                } else {
                Result = EID{Core.F_member_type(ToType(_Ztype.At(1-1))).Id(),0}
                } 
              }  else if ((s.Id() == Core.C__at.Id()) && 
                (l.At(1-1).Isa.IsIn(C_property) == CTRUE)) { 
              { var p *ClaireProperty   = ToProperty(l.At(1-1))
                { var c *ClaireAny   = l.At(2-1)
                  if ((C_class.Id() == c.Isa.Id()) && 
                      (C_method.Id() == ANY(Core.F_CALL(ToProperty(Core.C__at.Id()),ARGS(EID{p.Id(),0},c.ToEID()))).Isa.Id())) { 
                    Result = EID{MakeConstantSet(ANY(Core.F_CALL(ToProperty(Core.C__at.Id()),ARGS(EID{p.Id(),0},c.ToEID())))).Id(),0}
                    } else {
                    Result = EID{C_any.Id(),0}
                    } 
                  } 
                } 
              }  else if ((s.Id() == C_get.Id()) && 
                (l.At(1-1).Isa.IsIn(C_relation) == CTRUE)) { 
              { var r *ClaireRelation   = ToRelation(l.At(1-1))
                if (r.Isa.IsIn(C_property) == CTRUE) { 
                  { var g0171 *ClaireProperty   = ToProperty(r.Id())
                    { var xs *ClaireObject   = Core.F__at_property1(g0171,ToTypeExpression(_Ztype.At(2-1)).Class_I())
                      if (C_slot.Id() == xs.Isa.Id()) { 
                        { var g0172 *ClaireSlot   = ToSlot(xs.Id())
                          if ((g0172.Range.Included(ToType(C_set.Id())) == CTRUE) && 
                              (C_compiler.Safety < 2)) { 
                            Result = EID{ToTypeExpression(OBJ(Core.F_CALL(C_range,ARGS(l.At(2-1).ToEID())))).Class_I().Id(),0}
                            }  else if (g0172.Range.Contains(g0172.Default) == CTRUE) { 
                            Result = EID{g0172.Range.Id(),0}
                            } else {
                            Result = EID{F_Optimize_extends_type(g0172.Range).Id(),0}
                            } 
                          } 
                        } else {
                        Result = EID{g0171.Range.Id(),0}
                        } 
                      } 
                    } 
                  }  else if (C_table.Id() == r.Isa.Id()) { 
                  { var g0174 *ClaireTable   = ToTable(r.Id())
                    if (g0174.Range.Contains(g0174.Default) == CTRUE) { 
                      Result = EID{g0174.Range.Id(),0}
                      } else {
                      Result = EID{F_Optimize_extends_type(g0174.Range).Id(),0}
                      } 
                    } 
                  } else {
                  Result = EID{CFALSE.Id(),0}
                  } 
                } 
              } else {
              { var r *ClaireAny   = F_Optimize_restriction_I_property(s,_Ztype,CTRUE)
                if (C_slot.Id() == r.Isa.Id()) { 
                  { var g0175 *ClaireSlot   = ToSlot(r)
                    _ = g0175
                    if ((s.Id() == C_instances.Id()) && 
                        (C_class.Id() == l.At(1-1).Isa.Id())) { 
                      { var _CL_obj *ClaireParam   = To_Param(new(ClaireParam).Is(C_Param))
                        _CL_obj.Arg = C_list
                        /*class->class*/_CL_obj.Params = MakeConstantList(C_of.Id())
                        /*list->list*/_CL_obj.Args = MakeConstantList(MakeConstantSet(l.At(1-1)).Id())
                        /*list->list*/Result = EID{_CL_obj.Id(),0}
                        } 
                      } else {
                      Result = EID{g0175.Range.Id(),0}
                      } 
                    } 
                  }  else if (C_method.Id() == r.Isa.Id()) { 
                  { var g0176 *ClaireMethod   = ToMethod(r)
                    _ = g0176
                    Result = F_Optimize_use_range_method(g0176,_Ztype)
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
  
// sort_abstract!(restriction!(s, %type, false) as type)))) ]
// this is the optimizer for messages : does not use the sort unless there is a macro
/* {1} The go function for: c_code(self:Call) [status=1] */
func F_c_code_Call (self *Language.Call ) EID { 
    var Result EID 
    Result = F_Optimize_c_code_call_Call(self,C_void)
    return Result} 
  
// The EID go function for: c_code @ Call (throw: true) 
func E_c_code_Call (self EID) EID { 
    return F_c_code_Call(Language.To_Call(OBJ(self)) )} 
  
/* {1} The go function for: c_code_call(self:Call,sx:class) [status=1] */
func F_Optimize_c_code_call_Call (self *Language.Call ,sx *ClaireClass ) EID { 
    var Result EID 
    
    { var s *ClaireProperty   = self.Selector
      { var l *ClaireList   = self.Args
        /*g_try(v2:"Result",loop:true) */
        var g0182I *ClaireBoolean  
        var try_1 EID 
        /*g_try(v2:"try_1",loop:false) */
        { 
          var v_and4 *ClaireBoolean  
          
          v_and4 = l.At(1-1).Isa.IsIn(Core.C_global_variable)
          if (v_and4 == CFALSE) {try_1 = EID{CFALSE.Id(),0}
          } else { 
            v_and4 = Equal(ANY(Core.F_CALL(C_range,ARGS(l.At(1-1).ToEID()))),CEMPTY.Id())
            if (v_and4 == CFALSE) {try_1 = EID{CFALSE.Id(),0}
            } else { 
              var try_2 EID 
              /*g_try(v2:"try_2",loop:false) */
              try_2 = F_Compile_designated_ask_any(ANY(Core.F_CALL(C_value,ARGS(l.At(1-1).ToEID()))))
              /* ERROR PROTECTION INSERTED (v_and4-try_1) */
              if ErrorIn(try_2) {try_1 = try_2
              } else {
              v_and4 = ToBoolean(OBJ(try_2))
              if (v_and4 == CFALSE) {try_1 = EID{CFALSE.Id(),0}
              } else { 
                try_1 = EID{CTRUE.Id(),0}} 
              } 
            } 
          }
          } 
        /* ERROR PROTECTION INSERTED (g0182I-Result) */
        if ErrorIn(try_1) {Result = try_1
        } else {
        g0182I = ToBoolean(OBJ(try_1))
        if (g0182I == CTRUE) { 
          Result = ToArray(l.Id()).NthPut(1,ANY(Core.F_CALL(C_value,ARGS(l.At(1-1).ToEID())))).ToEID()
          } else {
          Result = EID{CFALSE.Id(),0}
          } 
        }
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        { var m *ClaireAny  
          var try_3 EID 
          /*g_try(v2:"try_3",loop:false) */
          try_3 = F_Optimize_inline_optimize_ask_Call(self)
          /* ERROR PROTECTION INSERTED (m-Result) */
          if ErrorIn(try_3) {Result = try_3
          } else {
          m = ANY(try_3)
          { var b *ClaireBoolean   = l.At(1-1).Isa.IsIn(C_property)
            { var d *ClaireAny  
              var try_4 EID 
              /*g_try(v2:"try_4",loop:false) */
              try_4 = F_Optimize_daccess_any(self.Id(),Core.F__sup_integer(C_compiler.Safety,5))
              /* ERROR PROTECTION INSERTED (d-Result) */
              if ErrorIn(try_4) {Result = try_4
              } else {
              d = ANY(try_4)
              if ((b == CTRUE) && 
                  (((s.Id() == C_put.Id()) || 
                      (s.Id() == Core.C_write.Id())) && 
                    (l.Length() == 3))) { 
                Result = F_Optimize_c_code_write_Call(self)
                }  else if ((b == CTRUE) && 
                  ((s.Id() == Core.C_put_store.Id()) && 
                    ((l.Length() == 4) && 
                      (l.At(4-1) == CTRUE.Id())))) { 
                Result = F_Optimize_c_code_write_Call(self)
                }  else if ((b == CTRUE) && 
                  (s.Id() == Core.C_unknown_ask.Id())) { 
                Result = F_Optimize_c_code_hold_property(ToProperty(l.At(1-1)),l.At(2-1),CNULL,CTRUE)
                }  else if ((b == CTRUE) && 
                  (s.Id() == Core.C_known_ask.Id())) { 
                Result = F_Optimize_c_code_hold_property(ToProperty(l.At(1-1)),l.At(2-1),CNULL,CFALSE)
                }  else if ((b == CTRUE) && 
                  ((s.Id() == Core.C_erase.Id()) && 
                    (l.At(2-1).Isa.IsIn(C_Variable) == CTRUE))) { 
                { var arg_5 *ClaireAny  
                  _ = arg_5
                  var try_6 EID 
                  /*g_try(v2:"try_6",loop:false) */
                  try_6 = F_Optimize_Produce_erase_property(ToProperty(l.At(1-1)),To_Variable(l.At(2-1)))
                  /* ERROR PROTECTION INSERTED (arg_5-Result) */
                  if ErrorIn(try_6) {Result = try_6
                  } else {
                  arg_5 = ANY(try_6)
                  Result = Core.F_CALL(C_c_code,ARGS(arg_5.ToEID(),EID{sx.Id(),0}))
                  }
                  } 
                }  else if (s.Id() == C_safe.Id()) { 
                { var y int  = C_compiler.Safety
                  _ = y
                  { var b *ClaireBoolean   = C_compiler.Overflow_ask
                    _ = b
                    { var x *ClaireAny   = CNULL
                      _ = x
                      C_compiler.Safety = 1
                      /*integer->integer*/C_compiler.Overflow_ask = CTRUE
                      /*boolean->boolean*/var try_7 EID 
                      /*g_try(v2:"try_7",loop:true) */
                      { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                        _CL_obj.Selector = C_safe
                        /*property->property*//*g_try(v2:"try_7",loop:true) */
                        { 
                          var va_arg1 *Language.Call  
                          var va_arg2 *ClaireList  
                          va_arg1 = _CL_obj
                          var try_8 EID 
                          /*g_try(v2:"try_8",loop:false) */
                          { 
                            var v_bag_arg *ClaireAny  
                            try_8= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                            var try_9 EID 
                            /*g_try(v2:"try_9",loop:false) */
                            try_9 = Core.F_CALL(C_c_code,ARGS(l.At(1-1).ToEID(),EID{sx.Id(),0}))
                            /* ERROR PROTECTION INSERTED (v_bag_arg-try_8) */
                            if ErrorIn(try_9) {try_8 = try_9
                            } else {
                            v_bag_arg = ANY(try_9)
                            ToList(OBJ(try_8)).AddFast(v_bag_arg)}
                            } 
                          /* ERROR PROTECTION INSERTED (va_arg2-try_7) */
                          if ErrorIn(try_8) {try_7 = try_8
                          } else {
                          va_arg2 = ToList(OBJ(try_8))
                          va_arg1.Args = va_arg2
                          /*list->list*/try_7 = EID{va_arg2.Id(),0}
                          }
                          } 
                        /* ERROR PROTECTION INSERTED (try_7-try_7) */
                        if !ErrorIn(try_7) {
                        try_7 = EID{_CL_obj.Id(),0}
                        }
                        } 
                      /* ERROR PROTECTION INSERTED (x-Result) */
                      if ErrorIn(try_7) {Result = try_7
                      } else {
                      x = ANY(try_7)
                      Result = x.ToEID()
                      C_compiler.Safety = y
                      /*integer->integer*/C_compiler.Overflow_ask = b
                      /*boolean->boolean*/Result = x.ToEID()
                      }
                      } 
                    } 
                  } 
                }  else if (((s.Id() == C_add.Id()) || 
                    (s.Id() == C_add_I.Id())) && 
                  (b == CTRUE)) { 
                Result = F_Optimize_c_code_add_Call(self)
                } else {
                var g0183I *ClaireBoolean  
                var try_10 EID 
                /*g_try(v2:"try_10",loop:false) */
                { 
                  var v_and8 *ClaireBoolean  
                  
                  v_and8 = MakeBoolean((s.Id() == C_add.Id()) || (s.Id() == C_add_I.Id()))
                  if (v_and8 == CFALSE) {try_10 = EID{CFALSE.Id(),0}
                  } else { 
                    var try_11 EID 
                    /*g_try(v2:"try_11",loop:false) */
                    { var arg_12 *ClaireType  
                      _ = arg_12
                      var try_13 EID 
                      /*g_try(v2:"try_13",loop:false) */
                      try_13 = Core.F_CALL(C_c_type,ARGS(l.At(1-1).ToEID()))
                      /* ERROR PROTECTION INSERTED (arg_12-try_11) */
                      if ErrorIn(try_13) {try_11 = try_13
                      } else {
                      arg_12 = ToType(OBJ(try_13))
                      try_11 = EID{arg_12.Included(ToType(C_bag.Id())).Id(),0}
                      }
                      } 
                    /* ERROR PROTECTION INSERTED (v_and8-try_10) */
                    if ErrorIn(try_11) {try_10 = try_11
                    } else {
                    v_and8 = ToBoolean(OBJ(try_11))
                    if (v_and8 == CFALSE) {try_10 = EID{CFALSE.Id(),0}
                    } else { 
                      try_10 = EID{CTRUE.Id(),0}} 
                    } 
                  }
                  } 
                /* ERROR PROTECTION INSERTED (g0183I-Result) */
                if ErrorIn(try_10) {Result = try_10
                } else {
                g0183I = ToBoolean(OBJ(try_10))
                if (g0183I == CTRUE) { 
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
                    _ = arg_14
                    var try_15 EID 
                    /*g_try(v2:"try_15",loop:false) */
                    try_15 = F_Optimize_c_srange_method(ToMethod(m))
                    /* ERROR PROTECTION INSERTED (arg_14-Result) */
                    if ErrorIn(try_15) {Result = try_15
                    } else {
                    arg_14 = ToClass(OBJ(try_15))
                    Result = F_Optimize_c_inline_method1(ToMethod(m),l,arg_14)
                    }
                    } 
                  } else {
                  var g0184I *ClaireBoolean  
                  var try_16 EID 
                  /*g_try(v2:"try_16",loop:false) */
                  { 
                    var v_and9 *ClaireBoolean  
                    
                    v_and9 = MakeBoolean((s.Id() == Core.C__I_equal.Id()) || (s.Id() == C__equal.Id()))
                    if (v_and9 == CFALSE) {try_16 = EID{CFALSE.Id(),0}
                    } else { 
                      var try_17 EID 
                      /*g_try(v2:"try_17",loop:false) */
                      { var arg_18 *ClaireAny  
                        _ = arg_18
                        var try_19 EID 
                        /*g_try(v2:"try_19",loop:false) */
                        try_19 = F_Optimize_daccess_any(l.At(1-1),CTRUE)
                        /* ERROR PROTECTION INSERTED (arg_18-try_17) */
                        if ErrorIn(try_19) {try_17 = try_19
                        } else {
                        arg_18 = ANY(try_19)
                        try_17 = EID{Core.F_known_ask_any(arg_18).Id(),0}
                        }
                        } 
                      /* ERROR PROTECTION INSERTED (v_and9-try_16) */
                      if ErrorIn(try_17) {try_16 = try_17
                      } else {
                      v_and9 = ToBoolean(OBJ(try_17))
                      if (v_and9 == CFALSE) {try_16 = EID{CFALSE.Id(),0}
                      } else { 
                        try_16 = EID{CTRUE.Id(),0}} 
                      } 
                    }
                    } 
                  /* ERROR PROTECTION INSERTED (g0184I-Result) */
                  if ErrorIn(try_16) {Result = try_16
                  } else {
                  g0184I = ToBoolean(OBJ(try_16))
                  if (g0184I == CTRUE) { 
                    Result = F_Optimize_c_code_hold_property(ToProperty(ToList(OBJ(Core.F_CALL(C_args,ARGS(l.At(1-1).ToEID())))).At(1-1)),ToList(OBJ(Core.F_CALL(C_args,ARGS(l.At(1-1).ToEID())))).At(2-1),l.At(2-1),Equal(s.Id(),C__equal.Id()))
                    } else {
                    var g0185I *ClaireBoolean  
                    var try_20 EID 
                    /*g_try(v2:"try_20",loop:false) */
                    { 
                      var v_and10 *ClaireBoolean  
                      
                      v_and10 = MakeBoolean((s.Id() == Core.C__I_equal.Id()) || (s.Id() == C__equal.Id()))
                      if (v_and10 == CFALSE) {try_20 = EID{CFALSE.Id(),0}
                      } else { 
                        var try_21 EID 
                        /*g_try(v2:"try_21",loop:false) */
                        { var arg_22 *ClaireAny  
                          _ = arg_22
                          var try_23 EID 
                          /*g_try(v2:"try_23",loop:false) */
                          try_23 = F_Optimize_daccess_any(l.At(2-1),CTRUE)
                          /* ERROR PROTECTION INSERTED (arg_22-try_21) */
                          if ErrorIn(try_23) {try_21 = try_23
                          } else {
                          arg_22 = ANY(try_23)
                          try_21 = EID{Core.F_known_ask_any(arg_22).Id(),0}
                          }
                          } 
                        /* ERROR PROTECTION INSERTED (v_and10-try_20) */
                        if ErrorIn(try_21) {try_20 = try_21
                        } else {
                        v_and10 = ToBoolean(OBJ(try_21))
                        if (v_and10 == CFALSE) {try_20 = EID{CFALSE.Id(),0}
                        } else { 
                          try_20 = EID{CTRUE.Id(),0}} 
                        } 
                      }
                      } 
                    /* ERROR PROTECTION INSERTED (g0185I-Result) */
                    if ErrorIn(try_20) {Result = try_20
                    } else {
                    g0185I = ToBoolean(OBJ(try_20))
                    if (g0185I == CTRUE) { 
                      Result = F_Optimize_c_code_hold_property(ToProperty(ToList(OBJ(Core.F_CALL(C_args,ARGS(l.At(2-1).ToEID())))).At(1-1)),ToList(OBJ(Core.F_CALL(C_args,ARGS(l.At(2-1).ToEID())))).At(2-1),l.At(1-1),Equal(s.Id(),C__equal.Id()))
                      }  else if (((s.Id() == C_put.Id()) || 
                          (s.Id() == C_nth_equal.Id())) && 
                        ((C_table.Id() == l.At(1-1).Isa.Id()) && 
                          (l.Length() == 3))) { 
                      Result = F_Optimize_c_code_table_Call(self)
                      } else {
                      var g0186I *ClaireBoolean  
                      var try_24 EID 
                      /*g_try(v2:"try_24",loop:false) */
                      { 
                        var v_and11 *ClaireBoolean  
                        
                        v_and11 = MakeBoolean((s.Id() == C_nth_equal.Id()) || (s.Id() == C_nth_put.Id()))
                        if (v_and11 == CFALSE) {try_24 = EID{CFALSE.Id(),0}
                        } else { 
                          var try_25 EID 
                          /*g_try(v2:"try_25",loop:false) */
                          { var arg_26 *ClaireType  
                            _ = arg_26
                            var try_27 EID 
                            /*g_try(v2:"try_27",loop:false) */
                            try_27 = Core.F_CALL(C_c_type,ARGS(l.At(1-1).ToEID()))
                            /* ERROR PROTECTION INSERTED (arg_26-try_25) */
                            if ErrorIn(try_27) {try_25 = try_27
                            } else {
                            arg_26 = ToType(OBJ(try_27))
                            try_25 = EID{arg_26.Included(ToType(C_array.Id())).Id(),0}
                            }
                            } 
                          /* ERROR PROTECTION INSERTED (v_and11-try_24) */
                          if ErrorIn(try_25) {try_24 = try_25
                          } else {
                          v_and11 = ToBoolean(OBJ(try_25))
                          if (v_and11 == CFALSE) {try_24 = EID{CFALSE.Id(),0}
                          } else { 
                            v_and11 = Equal(MakeInteger(l.Length()).Id(),MakeInteger(3).Id())
                            if (v_and11 == CFALSE) {try_24 = EID{CFALSE.Id(),0}
                            } else { 
                              try_24 = EID{CTRUE.Id(),0}} 
                            } 
                          } 
                        }
                        } 
                      /* ERROR PROTECTION INSERTED (g0186I-Result) */
                      if ErrorIn(try_24) {Result = try_24
                      } else {
                      g0186I = ToBoolean(OBJ(try_24))
                      if (g0186I == CTRUE) { 
                        Result = F_Optimize_c_code_array_Call(self)
                        }  else if ((s.Id() == C_nth.Id()) || 
                          (((s.Id() == C_get.Id()) && 
                              (C_table.Id() == l.At(1-1).Isa.Id())) || 
                            ((s.Id() == C_nth_get.Id()) && 
                                (l.At(1-1).Isa.IsIn(C_array) == CTRUE)))) { 
                        Result = F_Optimize_c_code_nth_Call(self)
                        }  else if (s.Id() == C__Z.Id()) { 
                        Result = F_Optimize_c_code_belong_Call(self)
                        }  else if (s.Id() == Core.C_Id.Id()) { 
                        { var arg_28 *ClaireAny  
                          _ = arg_28
                          var try_29 EID 
                          /*g_try(v2:"try_29",loop:false) */
                          try_29 = EVAL(l.At(1-1))
                          /* ERROR PROTECTION INSERTED (arg_28-Result) */
                          if ErrorIn(try_29) {Result = try_29
                          } else {
                          arg_28 = ANY(try_29)
                          Result = Core.F_CALL(C_c_code,ARGS(arg_28.ToEID()))
                          }
                          } 
                        }  else if (s.Id() == Language.C_function_I.Id()) { 
                        { var arg_30 *ClaireString  
                          _ = arg_30
                          var try_31 EID 
                          /*g_try(v2:"try_31",loop:false) */
                          { var arg_32 *ClaireSymbol  
                            _ = arg_32
                            var try_33 EID 
                            /*g_try(v2:"try_33",loop:false) */
                            try_33 = Language.F_extract_symbol_any(l.At(1-1))
                            /* ERROR PROTECTION INSERTED (arg_32-try_31) */
                            if ErrorIn(try_33) {try_31 = try_33
                            } else {
                            arg_32 = ToSymbol(OBJ(try_33))
                            try_31 = EID{arg_32.String_I().Id(),0}
                            }
                            } 
                          /* ERROR PROTECTION INSERTED (arg_30-Result) */
                          if ErrorIn(try_31) {Result = try_31
                          } else {
                          arg_30 = ToString(OBJ(try_31))
                          Result = F_make_function_string(arg_30).ToEID()
                          }
                          } 
                        }  else if ((s.Id() == Core.C_not.Id()) && 
                          (l.At(1-1).Isa.IsIn(Language.C_Select) == CTRUE)) { 
                        Result = F_Optimize_c_code_not_Select(Language.To_Select(l.At(1-1)))
                        }  else if ((s.Id() == Core.C_call.Id()) && 
                          (l.At(1-1).Isa.IsIn(C_property) == CTRUE)) { 
                        { var arg_34 *Language.Call  
                          _ = arg_34
                          var try_35 EID 
                          /*g_try(v2:"try_35",loop:false) */
                          { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                            _CL_obj.Selector = ToProperty(l.At(1-1))
                            /*property->property*//*g_try(v2:"try_35",loop:false) */
                            { 
                              var va_arg1 *Language.Call  
                              var va_arg2 *ClaireList  
                              va_arg1 = _CL_obj
                              var try_36 EID 
                              /*g_try(v2:"try_36",loop:false) */
                              try_36 = l.Cdr()
                              /* ERROR PROTECTION INSERTED (va_arg2-try_35) */
                              if ErrorIn(try_36) {try_35 = try_36
                              } else {
                              va_arg2 = ToList(OBJ(try_36))
                              va_arg1.Args = va_arg2
                              /*list->list*/try_35 = EID{va_arg2.Id(),0}
                              }
                              } 
                            /* ERROR PROTECTION INSERTED (try_35-try_35) */
                            if !ErrorIn(try_35) {
                            try_35 = EID{_CL_obj.Id(),0}
                            }
                            } 
                          /* ERROR PROTECTION INSERTED (arg_34-Result) */
                          if ErrorIn(try_35) {Result = try_35
                          } else {
                          arg_34 = Language.To_Call(OBJ(try_35))
                          Result = Core.F_CALL(C_c_code,ARGS(EID{arg_34.Id(),0}))
                          }
                          } 
                        }  else if (s.Open == 3) { 
                        { var arg_37 *ClaireList  
                          _ = arg_37
                          var try_38 EID 
                          /*g_try(v2:"try_38",loop:false) */
                          { 
                            var v_list13 *ClaireList  
                            var x *ClaireAny  
                            var v_local13 *ClaireAny  
                            v_list13 = l
                            try_38 = EID{CreateList(ToType(CEMPTY.Id()),v_list13.Length()).Id(),0}
                            for CLcount := 0; CLcount < v_list13.Length(); CLcount++{ 
                              x = v_list13.At(CLcount)
                              var try_39 EID 
                              /*g_try(v2:"try_39",loop:tuple("try_38", EID)) */
                              try_39 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
                              /* ERROR PROTECTION INSERTED (v_local13-try_38) */
                              if ErrorIn(try_39) {try_38 = try_39
                              break
                              } else {
                              v_local13 = ANY(try_39)
                              ToList(OBJ(try_38)).PutAt(CLcount,v_local13)
                              } 
                            }
                            } 
                          /* ERROR PROTECTION INSERTED (arg_37-Result) */
                          if ErrorIn(try_38) {Result = try_38
                          } else {
                          arg_37 = ToList(OBJ(try_38))
                          Result = F_Optimize_c_warn_property(s,l,arg_37)
                          }
                          } 
                        } else {
                        var g0187I *ClaireBoolean  
                        { 
                          var v_and12 *ClaireBoolean  
                          
                          v_and12 = Equal(s.Id(),Language.C_bit_vector.Id())
                          if (v_and12 == CFALSE) {g0187I = CFALSE
                          } else { 
                            { var arg_40 *ClaireAny  
                              _ = arg_40
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
                              v_and12 = Core.F_not_any(arg_40)
                              } 
                            if (v_and12 == CFALSE) {g0187I = CFALSE
                            } else { 
                              g0187I = CTRUE} 
                            } 
                          } 
                        if (g0187I == CTRUE) { 
                          Result = EVAL(self.Id())
                          }  else if ((s.Id() == C_Compile_anyObject_I.Id()) || 
                            ((s.Id() == C_Compile_object_I.Id()) || 
                              ((s.Id() == C_add_method.Id()) && 
                                  (b == CTRUE)))) { 
                          Result = EID{self.Id(),0}
                          } else {
                          { var _Ztype *ClaireList  
                            var try_41 EID 
                            /*g_try(v2:"try_41",loop:false) */
                            { 
                              var v_list14 *ClaireList  
                              var x *ClaireAny  
                              var v_local14 *ClaireAny  
                              v_list14 = l
                              try_41 = EID{CreateList(ToType(CEMPTY.Id()),v_list14.Length()).Id(),0}
                              for CLcount := 0; CLcount < v_list14.Length(); CLcount++{ 
                                x = v_list14.At(CLcount)
                                var try_42 EID 
                                /*g_try(v2:"try_42",loop:tuple("try_41", EID)) */
                                try_42 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
                                /* ERROR PROTECTION INSERTED (v_local14-try_41) */
                                if ErrorIn(try_42) {try_41 = try_42
                                break
                                } else {
                                v_local14 = ANY(try_42)
                                ToList(OBJ(try_41)).PutAt(CLcount,v_local14)
                                } 
                              }
                              } 
                            /* ERROR PROTECTION INSERTED (_Ztype-Result) */
                            if ErrorIn(try_41) {Result = try_41
                            } else {
                            _Ztype = ToList(OBJ(try_41))
                            { var z *ClaireAny   = F_Optimize_restriction_I_property(s,_Ztype,CTRUE)
                              if (C_slot.Id() == z.Isa.Id()) { 
                                { var g0178 *ClaireSlot   = ToSlot(z)
                                  { var _Zunknown *ClaireBoolean   = MakeBoolean((g0178.Range.Contains(g0178.Default) != CTRUE) && (C_OPT.Knowns.Contain_ask(s.Id()) != CTRUE) && (C_compiler.Safety < 2))
                                    var g0188I *ClaireBoolean  
                                    var try_43 EID 
                                    /*g_try(v2:"try_43",loop:false) */
                                    { 
                                      /* Or stat: v="try_43", loop=false */
                                      var v_or18 *ClaireBoolean  
                                      
                                      /* Or stat: try not @ any(%unknown) with try:false, v="try_43", loop=false */
                                      v_or18 = _Zunknown.Not
                                      if (v_or18 == CTRUE) {try_43 = EID{CTRUE.Id(),0}
                                      } else { 
                                        /* Or stat: try Compile/designated? @ any(nth @ list(l,1)) with try:true, v="try_43", loop=false */
                                        var try_44 EID 
                                        /*g_try(v2:"try_44",loop:false) */
                                        try_44 = F_Compile_designated_ask_any(l.At(1-1))
                                        /* ERROR PROTECTION INSERTED (v_or18-try_43) */
                                        if ErrorIn(try_44) {try_43 = try_44
                                        } else {
                                        v_or18 = ToBoolean(OBJ(try_44))
                                        if (v_or18 == CTRUE) {try_43 = EID{CTRUE.Id(),0}
                                        } else { 
                                          try_43 = EID{CFALSE.Id(),0}} 
                                        } 
                                      }
                                      } 
                                    /* ERROR PROTECTION INSERTED (g0188I-Result) */
                                    if ErrorIn(try_43) {Result = try_43
                                    } else {
                                    g0188I = ToBoolean(OBJ(try_43))
                                    if (g0188I == CTRUE) { 
                                      { var _CL_obj *Language.CallSlot   = Language.To_CallSlot(new(Language.CallSlot).Is(Language.C_Call_slot))
                                        _CL_obj.Selector = g0178
                                        /*slot->slot*//*g_try(v2:"Result",loop:true) */
                                        { 
                                          var va_arg1 *Language.CallSlot  
                                          var va_arg2 *ClaireAny  
                                          va_arg1 = _CL_obj
                                          var try_45 EID 
                                          /*g_try(v2:"try_45",loop:false) */
                                          try_45 = Core.F_CALL(C_c_code,ARGS(l.At(1-1).ToEID(),EID{F_Compile_psort_any(Core.F_domain_I_restriction(ToRestriction(g0178.Id())).Id()).Id(),0}))
                                          /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                                          if ErrorIn(try_45) {Result = try_45
                                          } else {
                                          va_arg2 = ANY(try_45)
                                          va_arg1.Arg = va_arg2
                                          /*any->any*/Result = va_arg2.ToEID()
                                          }
                                          } 
                                        /* ERROR PROTECTION INSERTED (Result-Result) */
                                        if !ErrorIn(Result) {
                                        _CL_obj.Test = _Zunknown
                                        /*boolean->boolean*/Result = EID{_CL_obj.Id(),0}
                                        }
                                        } 
                                      } else {
                                      
                                      if (C_compiler.Optimize_ask == CTRUE) { 
                                        F_Compile_notice_void()
                                        
                                        } 
                                      Result = F_Optimize_c_warn_property(s,l,_Ztype)
                                      } 
                                    }
                                    } 
                                  } 
                                }  else if (C_method.Id() == z.Isa.Id()) { 
                                { var g0179 *ClaireMethod   = ToMethod(z)
                                  
                                  /*g_try(v2:"Result",loop:true) */
                                  if (_Ztype.Memq(C_void.Id()) == CTRUE) { 
                                    Result = Core.F_CALL(C_Compile_Cerror,ARGS(EID{MakeString("[205] call ~S uses a void argument [~S]").Id(),0},EID{self.Id(),0},EID{_Ztype.Id(),0}))
                                    } else {
                                    Result = EID{CFALSE.Id(),0}
                                    } 
                                  /* ERROR PROTECTION INSERTED (Result-Result) */
                                  if !ErrorIn(Result) {
                                  /*g_try(v2:"Result",loop:true) */
                                  if (((s.Id() == C_begin.Id()) || 
                                        (s.Id() == C_end.Id())) && 
                                      (l.At(1-1).Isa.IsIn(C_module) == CTRUE)) { 
                                    Result = EVAL(self.Id())
                                    } else {
                                    Result = EID{CFALSE.Id(),0}
                                    } 
                                  /* ERROR PROTECTION INSERTED (Result-Result) */
                                  if !ErrorIn(Result) {
                                  var g0189I *ClaireBoolean  
                                  var try_46 EID 
                                  /*g_try(v2:"try_46",loop:false) */
                                  { 
                                    /* Or stat: v="try_46", loop=false */
                                    var v_or17 *ClaireBoolean  
                                    
                                    /* Or stat: try = @ any(last @ list(domain @ restriction(g0179)),listargs) with try:true, v="try_46", loop=false */
                                    var try_47 EID 
                                    /*g_try(v2:"try_47",loop:false) */
                                    { var arg_48 *ClaireAny  
                                      _ = arg_48
                                      var try_49 EID 
                                      /*g_try(v2:"try_49",loop:false) */
                                      try_49 = Core.F_last_list(g0179.Domain)
                                      /* ERROR PROTECTION INSERTED (arg_48-try_47) */
                                      if ErrorIn(try_49) {try_47 = try_49
                                      } else {
                                      arg_48 = ANY(try_49)
                                      try_47 = EID{Equal(arg_48,C_listargs.Id()).Id(),0}
                                      }
                                      } 
                                    /* ERROR PROTECTION INSERTED (v_or17-try_46) */
                                    if ErrorIn(try_47) {try_46 = try_47
                                    } else {
                                    v_or17 = ToBoolean(OBJ(try_47))
                                    if (v_or17 == CTRUE) {try_46 = EID{CTRUE.Id(),0}
                                    } else { 
                                      /* Or stat: try ((= @ any(nth @ list(domain @ restriction(g0179),1),void)) & (!= @ any(nth @ list(l,1),<environment>))) with try:false, v="try_46", loop=false */
                                      v_or17 = MakeBoolean((g0179.Domain.ValuesO()[1-1] == C_void.Id()) && (l.At(1-1) != ClEnv.Id()))
                                      if (v_or17 == CTRUE) {try_46 = EID{CTRUE.Id(),0}
                                      } else { 
                                        try_46 = EID{CFALSE.Id(),0}} 
                                      } 
                                    }
                                    } 
                                  /* ERROR PROTECTION INSERTED (g0189I-Result) */
                                  if ErrorIn(try_46) {Result = try_46
                                  } else {
                                  g0189I = ToBoolean(OBJ(try_46))
                                  if (g0189I == CTRUE) { 
                                    Result = F_Optimize_open_message_property(s,l)
                                    } else {
                                    Result = F_Optimize_c_code_method_method2(g0179,l,_Ztype,sx)
                                    } 
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
// create the compiled message with necessary protections
/* {1} The go function for: open_message(self:property,l:list) [status=1] */
func F_Optimize_open_message_property (self *ClaireProperty ,l *ClaireList ) EID { 
    var Result EID 
    F_Optimize_selector_register_property(self)
    { var _Zarg *ClaireList  
      _ = _Zarg
      var try_1 EID 
      /*g_try(v2:"try_1",loop:false) */
      { 
        var v_list3 *ClaireList  
        var x *ClaireAny  
        var v_local3 *ClaireAny  
        v_list3 = l
        try_1 = EID{CreateList(ToType(CEMPTY.Id()),v_list3.Length()).Id(),0}
        for CLcount := 0; CLcount < v_list3.Length(); CLcount++{ 
          x = v_list3.At(CLcount)
          var try_2 EID 
          /*g_try(v2:"try_2",loop:tuple("try_1", EID)) */
          var g0191I *ClaireBoolean  
          var try_3 EID 
          /*g_try(v2:"try_3",loop:false) */
          { var arg_4 *ClaireType  
            _ = arg_4
            var try_5 EID 
            /*g_try(v2:"try_5",loop:false) */
            try_5 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
            /* ERROR PROTECTION INSERTED (arg_4-try_3) */
            if ErrorIn(try_5) {try_3 = try_5
            } else {
            arg_4 = ToType(OBJ(try_5))
            try_3 = EID{Core.F__I_equal_any(arg_4.Id(),C_void.Id()).Id(),0}
            }
            } 
          /* ERROR PROTECTION INSERTED (g0191I-try_2) */
          if ErrorIn(try_3) {try_2 = try_3
          } else {
          g0191I = ToBoolean(OBJ(try_3))
          if (g0191I == CTRUE) { 
            try_2 = Core.F_CALL(C_c_code,ARGS(x.ToEID(),EID{C_any.Id(),0}))
            } else {
            try_2 = Core.F_CALL(C_Compile_Cerror,ARGS(EID{MakeString("[206] use of void ~S in ~S~S").Id(),0},
              x.ToEID(),
              EID{self.Id(),0},
              EID{l.Id(),0}))
            } 
          }
          /* ERROR PROTECTION INSERTED (v_local3-try_1) */
          if ErrorIn(try_2) {try_1 = try_2
          break
          } else {
          v_local3 = ANY(try_2)
          ToList(OBJ(try_1)).PutAt(CLcount,v_local3)
          } 
        }
        } 
      /* ERROR PROTECTION INSERTED (_Zarg-Result) */
      if ErrorIn(try_1) {Result = try_1
      } else {
      _Zarg = ToList(OBJ(try_1))
      C_compiler.NDynamic = (C_compiler.NDynamic+1)
      /*integer->integer*/{ var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
        _CL_obj.Selector = self
        /*property->property*/_CL_obj.Args = _Zarg
        /*list->list*/Result = EID{_CL_obj.Id(),0}
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
/* {1} The go function for: daccess(self:any,b:boolean) [status=1] */
func F_Optimize_daccess_any (self *ClaireAny ,b *ClaireBoolean ) EID { 
    var Result EID 
    if (self.Isa.IsIn(Language.C_Call) == CTRUE) { 
      { var g0192 *Language.Call   = Language.To_Call(self)
        { var l *ClaireList   = g0192.Args
          { var xs *ClaireObject  
            var try_1 EID 
            /*g_try(v2:"try_1",loop:false) */
            if ((g0192.Selector.Id() == C_get.Id()) && 
                (l.At(1-1).Isa.IsIn(C_property) == CTRUE)) { 
              { var arg_2 *ClaireClass  
                _ = arg_2
                var try_3 EID 
                /*g_try(v2:"try_3",loop:false) */
                { var arg_4 *ClaireType  
                  _ = arg_4
                  var try_5 EID 
                  /*g_try(v2:"try_5",loop:false) */
                  try_5 = Core.F_CALL(C_c_type,ARGS(l.At(2-1).ToEID()))
                  /* ERROR PROTECTION INSERTED (arg_4-try_3) */
                  if ErrorIn(try_5) {try_3 = try_5
                  } else {
                  arg_4 = ToType(OBJ(try_5))
                  try_3 = EID{arg_4.Class_I().Id(),0}
                  }
                  } 
                /* ERROR PROTECTION INSERTED (arg_2-try_1) */
                if ErrorIn(try_3) {try_1 = try_3
                } else {
                arg_2 = ToClass(OBJ(try_3))
                try_1 = EID{Core.F__at_property1(ToProperty(l.At(1-1)),arg_2).Id(),0}
                }
                } 
              } else {
              try_1 = EID{CFALSE.Id(),0}
              } 
            /* ERROR PROTECTION INSERTED (xs-Result) */
            if ErrorIn(try_1) {Result = try_1
            } else {
            xs = ToObject(OBJ(try_1))
            var g0195I *ClaireBoolean  
            if (C_slot.Id() == xs.Isa.Id()) { 
              { var g0193 *ClaireSlot   = ToSlot(xs.Id())
                g0195I = MakeBoolean((b == CTRUE) || (g0193.Range.Contains(g0193.Default) == CTRUE) || (g0193.Srange.Id() == C_any.Id()) || (g0193.Srange.Id() == C_integer.Id()))
                } 
              } else {
              g0195I = CFALSE
              } 
            if (g0195I == CTRUE) { 
              { var _CL_obj *Language.CallSlot   = Language.To_CallSlot(new(Language.CallSlot).Is(Language.C_Call_slot))
                _CL_obj.Selector = ToSlot(xs.Id())
                /*slot->slot*//*g_try(v2:"Result",loop:true) */
                { 
                  var va_arg1 *Language.CallSlot  
                  var va_arg2 *ClaireAny  
                  va_arg1 = _CL_obj
                  var try_6 EID 
                  /*g_try(v2:"try_6",loop:false) */
                  try_6 = Core.F_CALL(C_c_code,ARGS(l.At(2-1).ToEID(),EID{F_Compile_psort_any(Core.F_domain_I_restriction(ToRestriction(xs.Id())).Id()).Id(),0}))
                  /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                  if ErrorIn(try_6) {Result = try_6
                  } else {
                  va_arg2 = ANY(try_6)
                  va_arg1.Arg = va_arg2
                  /*any->any*/Result = va_arg2.ToEID()
                  }
                  } 
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                _CL_obj.Test = CFALSE
                /*boolean->boolean*/Result = EID{_CL_obj.Id(),0}
                }
                } 
              } else {
              Result = EID{CNULL,0}
              } 
            }
            } 
          } 
        } 
      } else {
      Result = EID{CNULL,0}
      } 
    return Result} 
  
// The EID go function for: daccess @ any (throw: true) 
func E_Optimize_daccess_any (self EID,b EID) EID { 
    return F_Optimize_daccess_any(ANY(self),ToBoolean(OBJ(b)) )} 
  
/* {1} The go function for: c_type(self:Call_slot) [status=0] */
func F_c_type_Call_slot (self *Language.CallSlot ) *ClaireType  { 
    return  self.Selector.Range
    } 
  
// The EID go function for: c_type @ Call_slot (throw: false) 
func E_c_type_Call_slot (self EID) EID { 
    return EID{F_c_type_Call_slot(Language.To_CallSlot(OBJ(self)) ).Id(),0}} 
  
/* {1} The go function for: c_type(self:Call_table) [status=0] */
func F_c_type_Call_table (self *Language.CallTable ) *ClaireType  { 
    return  self.Selector.Range
    } 
  
// The EID go function for: c_type @ Call_table (throw: false) 
func E_c_type_Call_table (self EID) EID { 
    return EID{F_c_type_Call_table(Language.To_CallTable(OBJ(self)) ).Id(),0}} 
  
/* {1} The go function for: c_type(self:Call_array) [status=0] */
func F_c_type_Call_array (self *Language.CallArray ) *ClaireType  { 
    return  ToType(self.Test)
    } 
  
// The EID go function for: c_type @ Call_array (throw: false) 
func E_c_type_Call_array (self EID) EID { 
    return EID{F_c_type_Call_array(Language.To_CallArray(OBJ(self)) ).Id(),0}} 
  
// write optimization: ss is put, put_store or write
// note that a put(object,x that may be unknown) is hard to compile !
// v2.4.10 -> if x = unknown OK (o.r = NULL) otherwise use store
/* {1} The go function for: c_code_write(self:Call) [status=1] */
func F_Optimize_c_code_write_Call (self *Language.Call ) EID { 
    var Result EID 
    { var p *ClaireAny   = self.Args.At(1-1)
      { var x *ClaireAny   = self.Args.At(2-1)
        { var y *ClaireAny   = self.Args.At(3-1)
          { var yt *ClaireType  
            var try_1 EID 
            /*g_try(v2:"try_1",loop:false) */
            try_1 = Core.F_CALL(C_c_type,ARGS(y.ToEID()))
            /* ERROR PROTECTION INSERTED (yt-Result) */
            if ErrorIn(try_1) {Result = try_1
            } else {
            yt = ToType(OBJ(try_1))
            { var ss *ClaireProperty   = self.Selector
              { var s *ClaireAny  
                var try_2 EID 
                /*g_try(v2:"try_2",loop:false) */
                { var arg_3 *ClaireList  
                  _ = arg_3
                  var try_4 EID 
                  /*g_try(v2:"try_4",loop:false) */
                  { 
                    var v_bag_arg *ClaireAny  
                    try_4= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                    var try_5 EID 
                    /*g_try(v2:"try_5",loop:false) */
                    try_5 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
                    /* ERROR PROTECTION INSERTED (v_bag_arg-try_4) */
                    if ErrorIn(try_5) {try_4 = try_5
                    } else {
                    v_bag_arg = ANY(try_5)
                    ToList(OBJ(try_4)).AddFast(v_bag_arg)}
                    } 
                  /* ERROR PROTECTION INSERTED (arg_3-try_2) */
                  if ErrorIn(try_4) {try_2 = try_4
                  } else {
                  arg_3 = ToList(OBJ(try_4))
                  try_2 = Core.F_CALL(C_Optimize_restriction_I,ARGS(p.ToEID(),EID{arg_3.Id(),0},EID{CTRUE.Id(),0}))
                  }
                  } 
                /* ERROR PROTECTION INSERTED (s-Result) */
                if ErrorIn(try_2) {Result = try_2
                } else {
                s = ANY(try_2)
                
                if (C_OPT.ToRemove.Contain_ask(p) == CTRUE) { 
                  Result = EID{CNIL.Id(),0}
                  } else {
                  var g0199I *ClaireBoolean  
                  if (C_slot.Id() == s.Isa.Id()) { 
                    { var g0196 *ClaireSlot   = ToSlot(s)
                      _ = g0196
                      g0199I = MakeBoolean((yt.Included(g0196.Range) == CTRUE) || (C_compiler.Safety >= 2))
                      } 
                    } else {
                    g0199I = CFALSE
                    } 
                  if (g0199I == CTRUE) { 
                    /*g_try(v2:"Result",loop:true) */
                    var g0200I *ClaireBoolean  
                    var try_6 EID 
                    /*g_try(v2:"try_6",loop:false) */
                    { 
                      var v_and10 *ClaireBoolean  
                      
                      v_and10 = Core.F__I_equal_any(y,CNULL)
                      if (v_and10 == CFALSE) {try_6 = EID{CFALSE.Id(),0}
                      } else { 
                        var try_7 EID 
                        /*g_try(v2:"try_7",loop:false) */
                        { var arg_8 *ClaireBoolean  
                          _ = arg_8
                          var try_9 EID 
                          /*g_try(v2:"try_9",loop:false) */
                          { var arg_10 *ClaireAny  
                            _ = arg_10
                            var try_11 EID 
                            /*g_try(v2:"try_11",loop:false) */
                            try_11 = Core.F_CALL(ToProperty(C__exp.Id()),ARGS(EID{yt.Id(),0},Core.F_CALL(C_mClaire_srange,ARGS(s.ToEID()))))
                            /* ERROR PROTECTION INSERTED (arg_10-try_9) */
                            if ErrorIn(try_11) {try_9 = try_11
                            } else {
                            arg_10 = ANY(try_11)
                            try_9 = EID{F_boolean_I_any(arg_10).Id(),0}
                            }
                            } 
                          /* ERROR PROTECTION INSERTED (arg_8-try_7) */
                          if ErrorIn(try_9) {try_7 = try_9
                          } else {
                          arg_8 = ToBoolean(OBJ(try_9))
                          try_7 = EID{Core.F__I_equal_any(arg_8.Id(),CTRUE.Id()).Id(),0}
                          }
                          } 
                        /* ERROR PROTECTION INSERTED (v_and10-try_6) */
                        if ErrorIn(try_7) {try_6 = try_7
                        } else {
                        v_and10 = ToBoolean(OBJ(try_7))
                        if (v_and10 == CFALSE) {try_6 = EID{CFALSE.Id(),0}
                        } else { 
                          try_6 = EID{CTRUE.Id(),0}} 
                        } 
                      }
                      } 
                    /* ERROR PROTECTION INSERTED (g0200I-Result) */
                    if ErrorIn(try_6) {Result = try_6
                    } else {
                    g0200I = ToBoolean(OBJ(try_6))
                    if (g0200I == CTRUE) { 
                      F_Compile_warn_void()
                      Result = Core.F_tformat_string(MakeString("sort error in ~S: ~S is a ~S [253]\n"),1,MakeConstantList(self.Id(),y,yt.Id()))
                      } else {
                      Result = EID{CFALSE.Id(),0}
                      } 
                    }
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    var g0201I *ClaireBoolean  
                    var try_12 EID 
                    /*g_try(v2:"try_12",loop:false) */
                    { 
                      var v_and10 *ClaireBoolean  
                      
                      v_and10 = MakeBoolean((yt.Included(ToType(OBJ(Core.F_CALL(C_range,ARGS(s.ToEID()))))) == CTRUE) || (yt.Included(ToType(C_object.Id())) == CTRUE) || (ANY(Core.F_CALL(C_mClaire_srange,ARGS(s.ToEID()))) != C_object.Id()) || (y == CNULL))
                      if (v_and10 == CFALSE) {try_12 = EID{CFALSE.Id(),0}
                      } else { 
                        var try_13 EID 
                        /*g_try(v2:"try_13",loop:false) */
                        { 
                          /* Or stat: v="try_13", loop=false */
                          var v_or12 *ClaireBoolean  
                          
                          /* Or stat: try != @ any(ss,write) with try:false, v="try_13", loop=false */
                          v_or12 = Core.F__I_equal_any(ss.Id(),Core.C_write.Id())
                          if (v_or12 == CTRUE) {try_13 = EID{CTRUE.Id(),0}
                          } else { 
                            /* Or stat: try ((Update? @ list<type_expression>(relation, any, any)(<p:relation>,x,y)) & ((= @ any(multivalued? @ relation(p),false)) | (= @ any(get @ property(if_write,<<p:object>:object>),unknown)))) with try:true, v="try_13", loop=false */
                            var try_14 EID 
                            /*g_try(v2:"try_14",loop:false) */
                            { 
                              var v_and14 *ClaireBoolean  
                              
                              var try_15 EID 
                              /*g_try(v2:"try_15",loop:false) */
                              try_15 = F_Optimize_Update_ask_relation1(ToRelation(p),x,y)
                              /* ERROR PROTECTION INSERTED (v_and14-try_14) */
                              if ErrorIn(try_15) {try_14 = try_15
                              } else {
                              v_and14 = ToBoolean(OBJ(try_15))
                              if (v_and14 == CFALSE) {try_14 = EID{CFALSE.Id(),0}
                              } else { 
                                v_and14 = MakeBoolean((ToRelation(p).Multivalued_ask.Id() == CFALSE.Id()) || (Core.F_get_property(C_if_write,ToObject(p)) == CNULL))
                                if (v_and14 == CFALSE) {try_14 = EID{CFALSE.Id(),0}
                                } else { 
                                  try_14 = EID{CTRUE.Id(),0}} 
                                } 
                              }
                              } 
                            /* ERROR PROTECTION INSERTED (v_or12-try_13) */
                            if ErrorIn(try_14) {try_13 = try_14
                            } else {
                            v_or12 = ToBoolean(OBJ(try_14))
                            if (v_or12 == CTRUE) {try_13 = EID{CTRUE.Id(),0}
                            } else { 
                              try_13 = EID{CFALSE.Id(),0}} 
                            } 
                          }
                          } 
                        /* ERROR PROTECTION INSERTED (v_and10-try_12) */
                        if ErrorIn(try_13) {try_12 = try_13
                        } else {
                        v_and10 = ToBoolean(OBJ(try_13))
                        if (v_and10 == CFALSE) {try_12 = EID{CFALSE.Id(),0}
                        } else { 
                          try_12 = EID{CTRUE.Id(),0}} 
                        } 
                      }
                      } 
                    /* ERROR PROTECTION INSERTED (g0201I-Result) */
                    if ErrorIn(try_12) {Result = try_12
                    } else {
                    g0201I = ToBoolean(OBJ(try_12))
                    if (g0201I == CTRUE) { 
                      { var _Zx *ClaireAny  
                        _ = _Zx
                        var try_16 EID 
                        /*g_try(v2:"try_16",loop:false) */
                        try_16 = Core.F_CALL(C_c_code,ARGS(x.ToEID(),EID{F_Compile_psort_any(Core.F_domain_I_restriction(ToRestriction(s)).Id()).Id(),0}))
                        /* ERROR PROTECTION INSERTED (_Zx-Result) */
                        if ErrorIn(try_16) {Result = try_16
                        } else {
                        _Zx = ANY(try_16)
                        { var _Zy *ClaireAny  
                          _ = _Zy
                          var try_17 EID 
                          /*g_try(v2:"try_17",loop:false) */
                          try_17 = F_Compile_c_strict_code_any(y,F_Compile_psort_any(ANY(Core.F_CALL(C_range,ARGS(s.ToEID())))))
                          /* ERROR PROTECTION INSERTED (_Zy-Result) */
                          if ErrorIn(try_17) {Result = try_17
                          } else {
                          _Zy = ANY(try_17)
                          { var _CL_obj *Language.Update   = Language.To_Update(new(Language.Update).Is(Language.C_Update))
                            _CL_obj.Selector = p
                            /*any->any*/_CL_obj.Value = _Zy
                            /*any->any*//*g_try(v2:"Result",loop:true) */
                            { 
                              var va_arg1 *Language.Update  
                              var va_arg2 *ClaireAny  
                              va_arg1 = _CL_obj
                              var try_18 EID 
                              /*g_try(v2:"try_18",loop:false) */
                              if (ss.Id() != Core.C_write.Id()) { 
                                try_18 = EID{ss.Id(),0}
                                } else {
                                try_18 = Core.F_CALL(C_c_code,ARGS(x.ToEID(),EID{C_any.Id(),0}))
                                } 
                              /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                              if ErrorIn(try_18) {Result = try_18
                              } else {
                              va_arg2 = ANY(try_18)
                              va_arg1.Arg = va_arg2
                              /*any->any*/Result = va_arg2.ToEID()
                              }
                              } 
                            /* ERROR PROTECTION INSERTED (Result-Result) */
                            if !ErrorIn(Result) {
                            { 
                              var va_arg1 *Language.Update  
                              var va_arg2 *ClaireAny  
                              va_arg1 = _CL_obj
                              { var _CL_obj *Language.CallSlot   = Language.To_CallSlot(new(Language.CallSlot).Is(Language.C_Call_slot))
                                _CL_obj.Selector = ToSlot(s)
                                /*slot->slot*/_CL_obj.Arg = _Zx
                                /*any->any*/_CL_obj.Test = CFALSE
                                /*boolean->boolean*/va_arg2 = _CL_obj.Id()
                                } 
                              va_arg1.ClaireVar = va_arg2
                              /*any->any*/} 
                            Result = EID{_CL_obj.Id(),0}
                            }
                            } 
                          }
                          } 
                        }
                        } 
                      }  else if (ss.Id() == C_put.Id()) { 
                      { var arg_19 *Language.Call  
                        _ = arg_19
                        { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                          _CL_obj.Selector = C_store
                          /*property->property*/_CL_obj.Args = MakeConstantList(x,
                            ANY(Core.F_CALL(C_mClaire_index,ARGS(s.ToEID()))),
                            ANY(Core.F_CALL(C_mClaire_srange,ARGS(s.ToEID()))),
                            y,
                            ANY(Core.F_CALL(C_store_ask,ARGS(p.ToEID()))))
                          /*list->list*/arg_19 = _CL_obj
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
                        _ = arg_20
                        { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                          _CL_obj.Selector = Core.C_mClaire_update
                          /*property->property*/_CL_obj.Args = MakeConstantList(p,
                            x,
                            ANY(Core.F_CALL(C_mClaire_index,ARGS(s.ToEID()))),
                            ANY(Core.F_CALL(C_mClaire_srange,ARGS(s.ToEID()))),
                            y)
                          /*list->list*/arg_20 = _CL_obj
                          } 
                        Result = Core.F_CALL(C_c_code,ARGS(EID{arg_20.Id(),0}))
                        } 
                      } 
                    }
                    }
                    } else {
                    { var _Ztype *ClaireList  
                      var try_21 EID 
                      /*g_try(v2:"try_21",loop:false) */
                      { 
                        var v_list11 *ClaireList  
                        var x *ClaireAny  
                        var v_local11 *ClaireAny  
                        v_list11 = self.Args
                        try_21 = EID{CreateList(ToType(CEMPTY.Id()),v_list11.Length()).Id(),0}
                        for CLcount := 0; CLcount < v_list11.Length(); CLcount++{ 
                          x = v_list11.At(CLcount)
                          var try_22 EID 
                          /*g_try(v2:"try_22",loop:tuple("try_21", EID)) */
                          try_22 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
                          /* ERROR PROTECTION INSERTED (v_local11-try_21) */
                          if ErrorIn(try_22) {try_21 = try_22
                          break
                          } else {
                          v_local11 = ANY(try_22)
                          ToList(OBJ(try_21)).PutAt(CLcount,v_local11)
                          } 
                        }
                        } 
                      /* ERROR PROTECTION INSERTED (_Ztype-Result) */
                      if ErrorIn(try_21) {Result = try_21
                      } else {
                      _Ztype = ToList(OBJ(try_21))
                      { var z *ClaireAny   = F_Optimize_restriction_I_property(ss,_Ztype,CTRUE)
                        
                        if (C_method.Id() == z.Isa.Id()) { 
                          { var g0197 *ClaireMethod   = ToMethod(z)
                            _ = g0197
                            Result = F_Optimize_c_code_method_method1(g0197,self.Args,_Ztype)
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
    return Result} 
  
// The EID go function for: c_code_write @ Call (throw: true) 
func E_Optimize_c_code_write_Call (self EID) EID { 
    return F_Optimize_c_code_write_Call(Language.To_Call(OBJ(self)) )} 
  
// (get(p,x) =/= y) optimization. We try to use the smart form instead of the get
/* {1} The go function for: c_code_hold(p:property,x:any,y:any,b:boolean) [status=1] */
func F_Optimize_c_code_hold_property (p *ClaireProperty ,x *ClaireAny ,y *ClaireAny ,b *ClaireBoolean ) EID { 
    var Result EID 
    { var s *ClaireAny  
      var try_1 EID 
      /*g_try(v2:"try_1",loop:false) */
      { var arg_2 *ClaireList  
        _ = arg_2
        var try_3 EID 
        /*g_try(v2:"try_3",loop:false) */
        { 
          var v_bag_arg *ClaireAny  
          try_3= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
          var try_4 EID 
          /*g_try(v2:"try_4",loop:false) */
          try_4 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
          /* ERROR PROTECTION INSERTED (v_bag_arg-try_3) */
          if ErrorIn(try_4) {try_3 = try_4
          } else {
          v_bag_arg = ANY(try_4)
          ToList(OBJ(try_3)).AddFast(v_bag_arg)}
          } 
        /* ERROR PROTECTION INSERTED (arg_2-try_1) */
        if ErrorIn(try_3) {try_1 = try_3
        } else {
        arg_2 = ToList(OBJ(try_3))
        try_1 = F_Optimize_restriction_I_property(p,arg_2,CTRUE).ToEID()
        }
        } 
      /* ERROR PROTECTION INSERTED (s-Result) */
      if ErrorIn(try_1) {Result = try_1
      } else {
      s = ANY(try_1)
      var g0203I *ClaireBoolean  
      var try_5 EID 
      /*g_try(v2:"try_5",loop:false) */
      if (C_slot.Id() == s.Isa.Id()) { 
        { var g0202 *ClaireSlot   = ToSlot(s)
          _ = g0202
          { 
            /* Or stat: v="try_5", loop=false */
            var v_or5 *ClaireBoolean  
            
            /* Or stat: try = @ any(y,unknown) with try:false, v="try_5", loop=false */
            v_or5 = Equal(y,CNULL)
            if (v_or5 == CTRUE) {try_5 = EID{CTRUE.Id(),0}
            } else { 
              /* Or stat: try ((<= @ type_expression(c_type(y),mClaire/srange @ slot(g0202))) & (Compile/identifiable? @ any(y))) with try:true, v="try_5", loop=false */
              var try_6 EID 
              /*g_try(v2:"try_6",loop:false) */
              { 
                var v_and7 *ClaireBoolean  
                
                var try_7 EID 
                /*g_try(v2:"try_7",loop:false) */
                { var arg_8 *ClaireType  
                  _ = arg_8
                  var try_9 EID 
                  /*g_try(v2:"try_9",loop:false) */
                  try_9 = Core.F_CALL(C_c_type,ARGS(y.ToEID()))
                  /* ERROR PROTECTION INSERTED (arg_8-try_7) */
                  if ErrorIn(try_9) {try_7 = try_9
                  } else {
                  arg_8 = ToType(OBJ(try_9))
                  try_7 = EID{arg_8.Included(ToType(g0202.Srange.Id())).Id(),0}
                  }
                  } 
                /* ERROR PROTECTION INSERTED (v_and7-try_6) */
                if ErrorIn(try_7) {try_6 = try_7
                } else {
                v_and7 = ToBoolean(OBJ(try_7))
                if (v_and7 == CFALSE) {try_6 = EID{CFALSE.Id(),0}
                } else { 
                  var try_10 EID 
                  /*g_try(v2:"try_10",loop:false) */
                  try_10 = F_Compile_identifiable_ask_any(y)
                  /* ERROR PROTECTION INSERTED (v_and7-try_6) */
                  if ErrorIn(try_10) {try_6 = try_10
                  } else {
                  v_and7 = ToBoolean(OBJ(try_10))
                  if (v_and7 == CFALSE) {try_6 = EID{CFALSE.Id(),0}
                  } else { 
                    try_6 = EID{CTRUE.Id(),0}} 
                  } 
                }}
                } 
              /* ERROR PROTECTION INSERTED (v_or5-try_5) */
              if ErrorIn(try_6) {try_5 = try_6
              } else {
              v_or5 = ToBoolean(OBJ(try_6))
              if (v_or5 == CTRUE) {try_5 = EID{CTRUE.Id(),0}
              } else { 
                try_5 = EID{CFALSE.Id(),0}} 
              } 
            }
            } 
          } 
        } else {
        try_5 = EID{CFALSE.Id(),0}
        } 
      /* ERROR PROTECTION INSERTED (g0203I-Result) */
      if ErrorIn(try_5) {Result = try_5
      } else {
      g0203I = ToBoolean(OBJ(try_5))
      if (g0203I == CTRUE) { 
        { var cs *Language.CallSlot  
          _ = cs
          var try_11 EID 
          /*g_try(v2:"try_11",loop:false) */
          { var _CL_obj *Language.CallSlot   = Language.To_CallSlot(new(Language.CallSlot).Is(Language.C_Call_slot))
            _CL_obj.Selector = ToSlot(s)
            /*slot->slot*//*g_try(v2:"try_11",loop:false) */
            { 
              var va_arg1 *Language.CallSlot  
              var va_arg2 *ClaireAny  
              va_arg1 = _CL_obj
              var try_12 EID 
              /*g_try(v2:"try_12",loop:false) */
              try_12 = Core.F_CALL(C_c_code,ARGS(x.ToEID(),EID{F_Compile_psort_any(Core.F_domain_I_restriction(ToRestriction(s)).Id()).Id(),0}))
              /* ERROR PROTECTION INSERTED (va_arg2-try_11) */
              if ErrorIn(try_12) {try_11 = try_12
              } else {
              va_arg2 = ANY(try_12)
              va_arg1.Arg = va_arg2
              /*any->any*/try_11 = va_arg2.ToEID()
              }
              } 
            /* ERROR PROTECTION INSERTED (try_11-try_11) */
            if !ErrorIn(try_11) {
            _CL_obj.Test = CFALSE
            /*boolean->boolean*/try_11 = EID{_CL_obj.Id(),0}
            }
            } 
          /* ERROR PROTECTION INSERTED (cs-Result) */
          if ErrorIn(try_11) {Result = try_11
          } else {
          cs = Language.To_CallSlot(OBJ(try_11))
          { var cm *Language.CallMethod2  
            var try_13 EID 
            /*g_try(v2:"try_13",loop:false) */
            { var _CL_obj *Language.CallMethod2   = Language.To_CallMethod2(new(Language.CallMethod2).Is(Language.C_Call_method2))
              _CL_obj.Arg = ToMethod(Core.F__at_property1(Core.C_identical_ask,C_any).Id())
              /*method->method*//*g_try(v2:"try_13",loop:false) */
              { 
                var va_arg1 *Language.CallMethod  
                var va_arg2 *ClaireList  
                va_arg1 = Language.To_CallMethod(_CL_obj.Id())
                var try_14 EID 
                /*g_try(v2:"try_14",loop:false) */
                { 
                  var v_bag_arg *ClaireAny  
                  try_14= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                  ToList(OBJ(try_14)).AddFast(cs.Id())
                  var try_15 EID 
                  /*g_try(v2:"try_15",loop:false) */
                  try_15 = Core.F_CALL(C_c_code,ARGS(y.ToEID(),Core.F_CALL(C_mClaire_srange,ARGS(s.ToEID()))))
                  /* ERROR PROTECTION INSERTED (v_bag_arg-try_14) */
                  if ErrorIn(try_15) {try_14 = try_15
                  } else {
                  v_bag_arg = ANY(try_15)
                  ToList(OBJ(try_14)).AddFast(v_bag_arg)}
                  } 
                /* ERROR PROTECTION INSERTED (va_arg2-try_13) */
                if ErrorIn(try_14) {try_13 = try_14
                } else {
                va_arg2 = ToList(OBJ(try_14))
                va_arg1.Args = va_arg2
                /*list->list*/try_13 = EID{va_arg2.Id(),0}
                }
                } 
              /* ERROR PROTECTION INSERTED (try_13-try_13) */
              if !ErrorIn(try_13) {
              try_13 = EID{_CL_obj.Id(),0}
              }
              } 
            /* ERROR PROTECTION INSERTED (cm-Result) */
            if ErrorIn(try_13) {Result = try_13
            } else {
            cm = Language.To_CallMethod2(OBJ(try_13))
            if (b == CTRUE) { 
              Result = Core.F_CALL(C_c_code,ARGS(EID{cm.Id(),0}))
              } else {
              { var arg_16 *Language.Call  
                _ = arg_16
                { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                  _CL_obj.Selector = Core.C_not
                  /*property->property*/_CL_obj.Args = MakeConstantList(cm.Id())
                  /*list->list*/arg_16 = _CL_obj
                  } 
                Result = Core.F_CALL(C_c_code,ARGS(EID{arg_16.Id(),0}))
                } 
              } 
            }
            } 
          }
          } 
        } else {
        { var l *ClaireList   = MakeConstantList(C_any.Id(),C_any.Id())
          { var arg_17 *ClaireList  
            _ = arg_17
            var try_18 EID 
            /*g_try(v2:"try_18",loop:false) */
            { 
              var v_bag_arg *ClaireAny  
              try_18= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
              var try_19 EID 
              /*g_try(v2:"try_19",loop:false) */
              { var arg_20 *Language.Call  
                _ = arg_20
                { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                  _CL_obj.Selector = C_get
                  /*property->property*/_CL_obj.Args = MakeConstantList(p.Id(),x)
                  /*list->list*/arg_20 = _CL_obj
                  } 
                try_19 = Core.F_CALL(C_c_code,ARGS(EID{arg_20.Id(),0},EID{C_any.Id(),0}))
                } 
              /* ERROR PROTECTION INSERTED (v_bag_arg-try_18) */
              if ErrorIn(try_19) {try_18 = try_19
              } else {
              v_bag_arg = ANY(try_19)
              ToList(OBJ(try_18)).AddFast(v_bag_arg)
              var try_21 EID 
              /*g_try(v2:"try_21",loop:false) */
              try_21 = Core.F_CALL(C_c_code,ARGS(y.ToEID(),EID{C_any.Id(),0}))
              /* ERROR PROTECTION INSERTED (v_bag_arg-try_18) */
              if ErrorIn(try_21) {try_18 = try_21
              } else {
              v_bag_arg = ANY(try_21)
              ToList(OBJ(try_18)).AddFast(v_bag_arg)}}
              } 
            /* ERROR PROTECTION INSERTED (arg_17-Result) */
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
/* {1} The go function for: c_code_add(self:Call) [status=1] */
func F_Optimize_c_code_add_Call (self *Language.Call ) EID { 
    var Result EID 
    { var p *ClaireProperty   = ToProperty(self.Args.At(1-1))
      { var x *ClaireAny   = self.Args.At(2-1)
        { var y *ClaireAny   = self.Args.At(3-1)
          { var s *ClaireObject  
            var try_1 EID 
            /*g_try(v2:"try_1",loop:false) */
            { var arg_2 *ClaireClass  
              _ = arg_2
              var try_3 EID 
              /*g_try(v2:"try_3",loop:false) */
              { var arg_4 *ClaireType  
                _ = arg_4
                var try_5 EID 
                /*g_try(v2:"try_5",loop:false) */
                { var arg_6 *ClaireType  
                  _ = arg_6
                  var try_7 EID 
                  /*g_try(v2:"try_7",loop:false) */
                  try_7 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
                  /* ERROR PROTECTION INSERTED (arg_6-try_5) */
                  if ErrorIn(try_7) {try_5 = try_7
                  } else {
                  arg_6 = ToType(OBJ(try_7))
                  try_5 = EID{F_Optimize_ptype_type(arg_6).Id(),0}
                  }
                  } 
                /* ERROR PROTECTION INSERTED (arg_4-try_3) */
                if ErrorIn(try_5) {try_3 = try_5
                } else {
                arg_4 = ToType(OBJ(try_5))
                try_3 = EID{arg_4.Class_I().Id(),0}
                }
                } 
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
            var g0205I *ClaireBoolean  
            var try_8 EID 
            /*g_try(v2:"try_8",loop:false) */
            if (C_slot.Id() == s.Isa.Id()) { 
              { var g0204 *ClaireSlot   = ToSlot(s.Id())
                _ = g0204
                { 
                  /* Or stat: v="try_8", loop=false */
                  var v_or8 *ClaireBoolean  
                  
                  /* Or stat: try <= @ type_expression(c_type(y),member @ type(range @ restriction(g0204))) with try:true, v="try_8", loop=false */
                  var try_9 EID 
                  /*g_try(v2:"try_9",loop:false) */
                  { var arg_10 *ClaireType  
                    _ = arg_10
                    var try_11 EID 
                    /*g_try(v2:"try_11",loop:false) */
                    try_11 = Core.F_CALL(C_c_type,ARGS(y.ToEID()))
                    /* ERROR PROTECTION INSERTED (arg_10-try_9) */
                    if ErrorIn(try_11) {try_9 = try_11
                    } else {
                    arg_10 = ToType(OBJ(try_11))
                    try_9 = EID{arg_10.Included(Core.F_member_type(g0204.Range)).Id(),0}
                    }
                    } 
                  /* ERROR PROTECTION INSERTED (v_or8-try_8) */
                  if ErrorIn(try_9) {try_8 = try_9
                  } else {
                  v_or8 = ToBoolean(OBJ(try_9))
                  if (v_or8 == CTRUE) {try_8 = EID{CTRUE.Id(),0}
                  } else { 
                    /* Or stat: try >= @ integer(safety @ meta_compiler(compiler),2) with try:false, v="try_8", loop=false */
                    v_or8 = F__sup_equal_integer(C_compiler.Safety,2)
                    if (v_or8 == CTRUE) {try_8 = EID{CTRUE.Id(),0}
                    } else { 
                      try_8 = EID{CFALSE.Id(),0}} 
                    } 
                  }
                  } 
                } 
              } else {
              try_8 = EID{CFALSE.Id(),0}
              } 
            /* ERROR PROTECTION INSERTED (g0205I-Result) */
            if ErrorIn(try_8) {Result = try_8
            } else {
            g0205I = ToBoolean(OBJ(try_8))
            if (g0205I == CTRUE) { 
              if (F_Optimize_Update_ask_relation2(ToRelation(p.Id()),ToRelation(self.Selector.Id())) == CTRUE) { 
                { var x2 *ClaireAny  
                  _ = x2
                  var try_12 EID 
                  /*g_try(v2:"try_12",loop:false) */
                  try_12 = Core.F_CALL(C_c_code,ARGS(x.ToEID(),EID{F_Compile_psort_any(Core.F_domain_I_restriction(ToRestriction(s.Id())).Id()).Id(),0}))
                  /* ERROR PROTECTION INSERTED (x2-Result) */
                  if ErrorIn(try_12) {Result = try_12
                  } else {
                  x2 = ANY(try_12)
                  { var _CL_obj *Language.Update   = Language.To_Update(new(Language.Update).Is(Language.C_Update))
                    _CL_obj.Selector = p.Id()
                    /*any->any*/_CL_obj.Arg = C_add.Id()
                    /*any->any*/{ 
                      var va_arg1 *Language.Update  
                      var va_arg2 *ClaireAny  
                      va_arg1 = _CL_obj
                      { var _CL_obj *Language.CallSlot   = Language.To_CallSlot(new(Language.CallSlot).Is(Language.C_Call_slot))
                        _CL_obj.Selector = ToSlot(s.Id())
                        /*slot->slot*/_CL_obj.Arg = x2
                        /*any->any*/_CL_obj.Test = CFALSE
                        /*boolean->boolean*/va_arg2 = _CL_obj.Id()
                        } 
                      va_arg1.ClaireVar = va_arg2
                      /*any->any*/} 
                    /*g_try(v2:"Result",loop:true) */
                    { 
                      var va_arg1 *Language.Update  
                      var va_arg2 *ClaireAny  
                      va_arg1 = _CL_obj
                      var try_13 EID 
                      /*g_try(v2:"try_13",loop:false) */
                      try_13 = Core.F_CALL(C_c_code,ARGS(y.ToEID(),EID{F_Compile_psort_any(Core.F_member_type(ToType(OBJ(Core.F_CALL(C_range,ARGS(EID{s.Id(),0}))))).Id()).Id(),0}))
                      /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                      if ErrorIn(try_13) {Result = try_13
                      } else {
                      va_arg2 = ANY(try_13)
                      va_arg1.Value = va_arg2
                      /*any->any*/Result = va_arg2.ToEID()
                      }
                      } 
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    Result = EID{_CL_obj.Id(),0}
                    }
                    } 
                  }
                  } 
                } else {
                var g0206I *ClaireBoolean  
                var try_14 EID 
                /*g_try(v2:"try_14",loop:false) */
                { 
                  var v_and8 *ClaireBoolean  
                  
                  var try_15 EID 
                  /*g_try(v2:"try_15",loop:false) */
                  try_15 = F_Compile_designated_ask_any(x)
                  /* ERROR PROTECTION INSERTED (v_and8-try_14) */
                  if ErrorIn(try_15) {try_14 = try_15
                  } else {
                  v_and8 = ToBoolean(OBJ(try_15))
                  if (v_and8 == CFALSE) {try_14 = EID{CFALSE.Id(),0}
                  } else { 
                    v_and8 = MakeBoolean((p.Store_ask != CTRUE) && ((self.Selector.Id() == C_add_I.Id()) || 
                        (p.Inverse.Id() == CNULL)))
                    if (v_and8 == CFALSE) {try_14 = EID{CFALSE.Id(),0}
                    } else { 
                      try_14 = EID{CTRUE.Id(),0}} 
                    } 
                  }
                  } 
                /* ERROR PROTECTION INSERTED (g0206I-Result) */
                if ErrorIn(try_14) {Result = try_14
                } else {
                g0206I = ToBoolean(OBJ(try_14))
                if (g0206I == CTRUE) { 
                  { var x2 *ClaireAny  
                    _ = x2
                    var try_16 EID 
                    /*g_try(v2:"try_16",loop:false) */
                    try_16 = Core.F_CALL(C_c_code,ARGS(x.ToEID(),EID{F_Compile_psort_any(Core.F_domain_I_restriction(ToRestriction(s.Id())).Id()).Id(),0}))
                    /* ERROR PROTECTION INSERTED (x2-Result) */
                    if ErrorIn(try_16) {Result = try_16
                    } else {
                    x2 = ANY(try_16)
                    { var arg_17 *Language.Call  
                      _ = arg_17
                      { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                        _CL_obj.Selector = ToProperty(C_add.Id())
                        /*property->property*/{ 
                          var va_arg1 *Language.Call  
                          var va_arg2 *ClaireList  
                          va_arg1 = _CL_obj
                          { 
                            var v_bag_arg *ClaireAny  
                            va_arg2= ToType(CEMPTY.Id()).EmptyList()
                            { var _CL_obj *Language.CallSlot   = Language.To_CallSlot(new(Language.CallSlot).Is(Language.C_Call_slot))
                              _CL_obj.Selector = ToSlot(s.Id())
                              /*slot->slot*/_CL_obj.Arg = x2
                              /*any->any*/_CL_obj.Test = CFALSE
                              /*boolean->boolean*/v_bag_arg = _CL_obj.Id()
                              } 
                            va_arg2.AddFast(v_bag_arg)
                            va_arg2.AddFast(y)} 
                          va_arg1.Args = va_arg2
                          /*list->list*/} 
                        arg_17 = _CL_obj
                        } 
                      Result = Core.F_CALL(C_c_code,ARGS(EID{arg_17.Id(),0}))
                      } 
                    }
                    } 
                  } else {
                  if (C_compiler.Optimize_ask == CTRUE) { 
                    F_Compile_notice_void()
                    
                    } 
                  { var arg_18 *ClaireList  
                    _ = arg_18
                    var try_19 EID 
                    /*g_try(v2:"try_19",loop:false) */
                    { 
                      var v_bag_arg *ClaireAny  
                      try_19= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                      ToList(OBJ(try_19)).AddFast(C_property.Id())
                      var try_20 EID 
                      /*g_try(v2:"try_20",loop:false) */
                      try_20 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
                      /* ERROR PROTECTION INSERTED (v_bag_arg-try_19) */
                      if ErrorIn(try_20) {try_19 = try_20
                      } else {
                      v_bag_arg = ANY(try_20)
                      ToList(OBJ(try_19)).AddFast(v_bag_arg)
                      ToList(OBJ(try_19)).AddFast(C_integer.Id())
                      var try_21 EID 
                      /*g_try(v2:"try_21",loop:false) */
                      try_21 = Core.F_CALL(C_c_type,ARGS(y.ToEID()))
                      /* ERROR PROTECTION INSERTED (v_bag_arg-try_19) */
                      if ErrorIn(try_21) {try_19 = try_21
                      } else {
                      v_bag_arg = ANY(try_21)
                      ToList(OBJ(try_19)).AddFast(v_bag_arg)}}
                      } 
                    /* ERROR PROTECTION INSERTED (arg_18-Result) */
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
                _ = arg_22
                var try_23 EID 
                /*g_try(v2:"try_23",loop:false) */
                { 
                  var v_list8 *ClaireList  
                  var x *ClaireAny  
                  var v_local8 *ClaireAny  
                  v_list8 = self.Args
                  try_23 = EID{CreateList(ToType(CEMPTY.Id()),v_list8.Length()).Id(),0}
                  for CLcount := 0; CLcount < v_list8.Length(); CLcount++{ 
                    x = v_list8.At(CLcount)
                    var try_24 EID 
                    /*g_try(v2:"try_24",loop:tuple("try_23", EID)) */
                    try_24 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
                    /* ERROR PROTECTION INSERTED (v_local8-try_23) */
                    if ErrorIn(try_24) {try_23 = try_24
                    break
                    } else {
                    v_local8 = ANY(try_24)
                    ToList(OBJ(try_23)).PutAt(CLcount,v_local8)
                    } 
                  }
                  } 
                /* ERROR PROTECTION INSERTED (arg_22-Result) */
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
/* {1} The go function for: c_code_add_bag(self:Call) [status=1] */
func F_Optimize_c_code_add_bag_Call (self *Language.Call ) EID { 
    var Result EID 
    { var _Zt1 *ClaireType  
      var try_1 EID 
      /*g_try(v2:"try_1",loop:false) */
      try_1 = Core.F_CALL(C_c_type,ARGS(self.Args.At(1-1).ToEID()))
      /* ERROR PROTECTION INSERTED (_Zt1-Result) */
      if ErrorIn(try_1) {Result = try_1
      } else {
      _Zt1 = ToType(OBJ(try_1))
      { var _Zt2 *ClaireType  
        var try_2 EID 
        /*g_try(v2:"try_2",loop:false) */
        { var arg_3 *ClaireType  
          _ = arg_3
          var try_4 EID 
          /*g_try(v2:"try_4",loop:false) */
          try_4 = Core.F_CALL(C_c_type,ARGS(self.Args.At(2-1).ToEID()))
          /* ERROR PROTECTION INSERTED (arg_3-try_2) */
          if ErrorIn(try_4) {try_2 = try_4
          } else {
          arg_3 = ToType(OBJ(try_4))
          try_2 = EID{F_Optimize_ptype_type(arg_3).Id(),0}
          }
          } 
        /* ERROR PROTECTION INSERTED (_Zt2-Result) */
        if ErrorIn(try_2) {Result = try_2
        } else {
        _Zt2 = ToType(OBJ(try_2))
        { var _Zp *ClaireProperty  
          _ = _Zp
          if (((_Zt1.Isa.IsIn(C_Param) == CTRUE) && 
                (_Zt2.Included(Core.F_member_type(_Zt1)) == CTRUE)) || 
              (C_compiler.Safety >= 2)) { 
            _Zp = ToProperty(C_add_I.Id())
            } else {
            _Zp = self.Selector
            } 
          { var _Zltype *ClaireList   = MakeConstantList(_Zt1.Id(),_Zt2.Id())
            { var z *ClaireAny   = F_Optimize_restriction_I_property(_Zp,_Zltype,CTRUE)
              
              if ((_Zt2.Included(Core.F_member_type(_Zt1)) != CTRUE) && 
                  (self.Selector.Id() == C_add.Id())) { 
                F_Compile_warn_void()
                Core.F_tformat_string(MakeString("the bag addition ~S is poorly typed (~S) [251] \n"),1,MakeConstantList(self.Id(),Core.F_member_type(_Zt1).Id()))
                } 
              if (C_method.Id() == z.Isa.Id()) { 
                { var g0207 *ClaireMethod   = ToMethod(z)
                  _ = g0207
                  Result = F_Optimize_c_code_method_method1(g0207,self.Args,_Zltype)
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
    return Result} 
  
// The EID go function for: c_code_add_bag @ Call (throw: true) 
func E_Optimize_c_code_add_bag_Call (self EID) EID { 
    return F_Optimize_c_code_add_bag_Call(Language.To_Call(OBJ(self)) )} 
  
// delete optimization
// <yc> 7/98 new, also needed
/* {1} The go function for: c_code_delete(self:Call) [status=1] */
func F_Optimize_c_code_delete_Call (self *Language.Call ) EID { 
    var Result EID 
    { var p *ClaireAny   = self.Args.At(1-1)
      { var x *ClaireAny   = self.Args.At(2-1)
        { var y *ClaireAny   = self.Args.At(3-1)
          { var s *ClaireObject  
            var try_1 EID 
            /*g_try(v2:"try_1",loop:false) */
            { var arg_2 *ClaireClass  
              _ = arg_2
              var try_3 EID 
              /*g_try(v2:"try_3",loop:false) */
              { var arg_4 *ClaireType  
                _ = arg_4
                var try_5 EID 
                /*g_try(v2:"try_5",loop:false) */
                try_5 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
                /* ERROR PROTECTION INSERTED (arg_4-try_3) */
                if ErrorIn(try_5) {try_3 = try_5
                } else {
                arg_4 = ToType(OBJ(try_5))
                try_3 = EID{arg_4.Class_I().Id(),0}
                }
                } 
              /* ERROR PROTECTION INSERTED (arg_2-try_1) */
              if ErrorIn(try_3) {try_1 = try_3
              } else {
              arg_2 = ToClass(OBJ(try_3))
              try_1 = EID{Core.F__at_property1(ToProperty(p),arg_2).Id(),0}
              }
              } 
            /* ERROR PROTECTION INSERTED (s-Result) */
            if ErrorIn(try_1) {Result = try_1
            } else {
            s = ToObject(OBJ(try_1))
            var g0210I *ClaireBoolean  
            var try_6 EID 
            /*g_try(v2:"try_6",loop:false) */
            { 
              var v_and6 *ClaireBoolean  
              
              v_and6 = MakeBoolean((ToRelation(p).Inverse.Id() == CNULL))
              if (v_and6 == CFALSE) {try_6 = EID{CFALSE.Id(),0}
              } else { 
                var try_7 EID 
                /*g_try(v2:"try_7",loop:false) */
                try_7 = F_Compile_designated_ask_any(x)
                /* ERROR PROTECTION INSERTED (v_and6-try_6) */
                if ErrorIn(try_7) {try_6 = try_7
                } else {
                v_and6 = ToBoolean(OBJ(try_7))
                if (v_and6 == CFALSE) {try_6 = EID{CFALSE.Id(),0}
                } else { 
                  var try_8 EID 
                  /*g_try(v2:"try_8",loop:false) */
                  if (C_slot.Id() == s.Isa.Id()) { 
                    { var g0209 *ClaireSlot   = ToSlot(s.Id())
                      _ = g0209
                      { 
                        /* Or stat: v="try_8", loop=false */
                        var v_or11 *ClaireBoolean  
                        
                        /* Or stat: try <= @ type_expression(c_type(y),member @ type(range @ restriction(g0209))) with try:true, v="try_8", loop=false */
                        var try_9 EID 
                        /*g_try(v2:"try_9",loop:false) */
                        { var arg_10 *ClaireType  
                          _ = arg_10
                          var try_11 EID 
                          /*g_try(v2:"try_11",loop:false) */
                          try_11 = Core.F_CALL(C_c_type,ARGS(y.ToEID()))
                          /* ERROR PROTECTION INSERTED (arg_10-try_9) */
                          if ErrorIn(try_11) {try_9 = try_11
                          } else {
                          arg_10 = ToType(OBJ(try_11))
                          try_9 = EID{arg_10.Included(Core.F_member_type(g0209.Range)).Id(),0}
                          }
                          } 
                        /* ERROR PROTECTION INSERTED (v_or11-try_8) */
                        if ErrorIn(try_9) {try_8 = try_9
                        } else {
                        v_or11 = ToBoolean(OBJ(try_9))
                        if (v_or11 == CTRUE) {try_8 = EID{CTRUE.Id(),0}
                        } else { 
                          /* Or stat: try >= @ integer(safety @ meta_compiler(compiler),2) with try:false, v="try_8", loop=false */
                          v_or11 = F__sup_equal_integer(C_compiler.Safety,2)
                          if (v_or11 == CTRUE) {try_8 = EID{CTRUE.Id(),0}
                          } else { 
                            try_8 = EID{CFALSE.Id(),0}} 
                          } 
                        }
                        } 
                      } 
                    } else {
                    try_8 = EID{CFALSE.Id(),0}
                    } 
                  /* ERROR PROTECTION INSERTED (v_and6-try_6) */
                  if ErrorIn(try_8) {try_6 = try_8
                  } else {
                  v_and6 = ToBoolean(OBJ(try_8))
                  if (v_and6 == CFALSE) {try_6 = EID{CFALSE.Id(),0}
                  } else { 
                    try_6 = EID{CTRUE.Id(),0}} 
                  } 
                } 
              }}
              } 
            /* ERROR PROTECTION INSERTED (g0210I-Result) */
            if ErrorIn(try_6) {Result = try_6
            } else {
            g0210I = ToBoolean(OBJ(try_6))
            if (g0210I == CTRUE) { 
              { var x2 *ClaireAny  
                _ = x2
                var try_12 EID 
                /*g_try(v2:"try_12",loop:false) */
                try_12 = Core.F_CALL(C_c_code,ARGS(x.ToEID(),EID{F_Compile_psort_any(Core.F_domain_I_restriction(ToRestriction(s.Id())).Id()).Id(),0}))
                /* ERROR PROTECTION INSERTED (x2-Result) */
                if ErrorIn(try_12) {Result = try_12
                } else {
                x2 = ANY(try_12)
                { var arg_13 *Language.Call  
                  _ = arg_13
                  { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                    _CL_obj.Selector = ToProperty(C_delete.Id())
                    /*property->property*/{ 
                      var va_arg1 *Language.Call  
                      var va_arg2 *ClaireList  
                      va_arg1 = _CL_obj
                      { 
                        var v_bag_arg *ClaireAny  
                        va_arg2= ToType(CEMPTY.Id()).EmptyList()
                        { var _CL_obj *Language.CallSlot   = Language.To_CallSlot(new(Language.CallSlot).Is(Language.C_Call_slot))
                          _CL_obj.Selector = ToSlot(s.Id())
                          /*slot->slot*/_CL_obj.Arg = x2
                          /*any->any*/_CL_obj.Test = CFALSE
                          /*boolean->boolean*/v_bag_arg = _CL_obj.Id()
                          } 
                        va_arg2.AddFast(v_bag_arg)
                        va_arg2.AddFast(y)} 
                      va_arg1.Args = va_arg2
                      /*list->list*/} 
                    arg_13 = _CL_obj
                    } 
                  Result = Core.F_CALL(C_c_code,ARGS(EID{arg_13.Id(),0}))
                  } 
                }
                } 
              } else {
              { var arg_14 *ClaireList  
                _ = arg_14
                var try_15 EID 
                /*g_try(v2:"try_15",loop:false) */
                { 
                  var v_list8 *ClaireList  
                  var x *ClaireAny  
                  var v_local8 *ClaireAny  
                  v_list8 = self.Args
                  try_15 = EID{CreateList(ToType(CEMPTY.Id()),v_list8.Length()).Id(),0}
                  for CLcount := 0; CLcount < v_list8.Length(); CLcount++{ 
                    x = v_list8.At(CLcount)
                    var try_16 EID 
                    /*g_try(v2:"try_16",loop:tuple("try_15", EID)) */
                    try_16 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
                    /* ERROR PROTECTION INSERTED (v_local8-try_15) */
                    if ErrorIn(try_16) {try_15 = try_16
                    break
                    } else {
                    v_local8 = ANY(try_16)
                    ToList(OBJ(try_15)).PutAt(CLcount,v_local8)
                    } 
                  }
                  } 
                /* ERROR PROTECTION INSERTED (arg_14-Result) */
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
/* {1} The go function for: c_code_not(x:Select) [status=1] */
func F_Optimize_c_code_not_Select (x *Language.Select ) EID { 
    var Result EID 
    { var arg_1 *Language.Call  
      _ = arg_1
      { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
        _CL_obj.Selector = Core.C_not
        /*property->property*/{ 
          var va_arg1 *Language.Call  
          var va_arg2 *ClaireList  
          va_arg1 = _CL_obj
          { 
            var v_bag_arg *ClaireAny  
            va_arg2= ToType(CEMPTY.Id()).EmptyList()
            { var _CL_obj *Language.For   = Language.To_For(new(Language.For).Is(Language.C_For))
              _CL_obj.ClaireVar = x.ClaireVar
              /*Variable->Variable*/_CL_obj.SetArg = x.SetArg
              /*any->any*/_CL_obj.Arg = Language.C_If.Make(x.Arg,Language.C_Return.Make(CTRUE.Id()),CNULL)
              /*any->any*/v_bag_arg = _CL_obj.Id()
              } 
            va_arg2.AddFast(v_bag_arg)} 
          va_arg1.Args = va_arg2
          /*list->list*/} 
        arg_1 = _CL_obj
        } 
      Result = Core.F_CALL(C_c_code,ARGS(EID{arg_1.Id(),0}))
      } 
    return Result} 
  
// The EID go function for: c_code_not @ Select (throw: true) 
func E_Optimize_c_code_not_Select (x EID) EID { 
    return F_Optimize_c_code_not_Select(Language.To_Select(OBJ(x)) )} 
  
// old % optimization
/* {1} The go function for: c_code_belong(self:Call) [status=1] */
func F_Optimize_c_code_belong_Call (self *Language.Call ) EID { 
    var Result EID 
    { var x *ClaireAny   = self.Args.At(1-1)
      { var y *ClaireAny   = self.Args.At(2-1)
        { var _Ztype *ClaireList  
          var try_1 EID 
          /*g_try(v2:"try_1",loop:false) */
          { 
            var v_bag_arg *ClaireAny  
            try_1= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
            var try_2 EID 
            /*g_try(v2:"try_2",loop:false) */
            try_2 = Core.F_CALL(C_c_type,ARGS(y.ToEID()))
            /* ERROR PROTECTION INSERTED (v_bag_arg-try_1) */
            if ErrorIn(try_2) {try_1 = try_2
            } else {
            v_bag_arg = ANY(try_2)
            ToList(OBJ(try_1)).AddFast(v_bag_arg)
            var try_3 EID 
            /*g_try(v2:"try_3",loop:false) */
            try_3 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
            /* ERROR PROTECTION INSERTED (v_bag_arg-try_1) */
            if ErrorIn(try_3) {try_1 = try_3
            } else {
            v_bag_arg = ANY(try_3)
            ToList(OBJ(try_1)).AddFast(v_bag_arg)}}
            } 
          /* ERROR PROTECTION INSERTED (_Ztype-Result) */
          if ErrorIn(try_1) {Result = try_1
          } else {
          _Ztype = ToList(OBJ(try_1))
          if (C_set.Id() == y.Isa.Id()) { 
            { var _CL_obj *Language.Or   = Language.To_Or(new(Language.Or).Is(Language.C_Or))
              /*g_try(v2:"Result",loop:true) */
              { 
                var va_arg1 *Language.Or  
                var va_arg2 *ClaireList  
                va_arg1 = _CL_obj
                var try_4 EID 
                /*g_try(v2:"try_4",loop:false) */
                { var z_bag *ClaireList   = ToType(CEMPTY.Id()).EmptyList()
                  /*g_try(v2:"try_4",loop:false) */
                  { 
                    var z *ClaireAny  
                    _ = z
                    try_4= EID{CFALSE.Id(),0}
                    var z_support *ClaireList  
                    var try_5 EID 
                    /*g_try(v2:"try_5",loop:false) */
                    try_5 = Core.F_enumerate_any(y)
                    /* ERROR PROTECTION INSERTED (z_support-try_4) */
                    if ErrorIn(try_5) {try_4 = try_5
                    } else {
                    z_support = ToList(OBJ(try_5))
                    z_len := z_support.Length()
                    for i_it := 0; i_it < z_len; i_it++ { 
                      z = z_support.At(i_it)
                      var loop_6 EID 
                      _ = loop_6
                      /*g_try(v2:"loop_6",loop:tuple("try_4", EID)) */
                      { var arg_7 *ClaireAny  
                        _ = arg_7
                        var try_8 EID 
                        /*g_try(v2:"try_8",loop:false) */
                        { var arg_9 *Language.Call  
                          _ = arg_9
                          { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                            _CL_obj.Selector = ToProperty(C__equal.Id())
                            /*property->property*/_CL_obj.Args = MakeConstantList(x,z)
                            /*list->list*/arg_9 = _CL_obj
                            } 
                          try_8 = Core.F_CALL(C_c_code,ARGS(EID{arg_9.Id(),0},EID{C_any.Id(),0}))
                          } 
                        /* ERROR PROTECTION INSERTED (arg_7-loop_6) */
                        if ErrorIn(try_8) {loop_6 = try_8
                        } else {
                        arg_7 = ANY(try_8)
                        loop_6 = EID{z_bag.AddFast(arg_7).Id(),0}/*t=any,s=EID*/
                        }
                        } 
                      /* ERROR PROTECTION INSERTED (loop_6-try_4) */
                      if ErrorIn(loop_6) {try_4 = loop_6
                      break
                      } else {
                      }}
                      } 
                    } 
                  /* ERROR PROTECTION INSERTED (try_4-try_4) */
                  if !ErrorIn(try_4) {
                  try_4 = EID{z_bag.Id(),0}
                  }
                  } 
                /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                if ErrorIn(try_4) {Result = try_4
                } else {
                va_arg2 = ToList(OBJ(try_4))
                va_arg1.Args = va_arg2
                /*list->list*/Result = EID{va_arg2.Id(),0}
                }
                } 
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              Result = EID{_CL_obj.Id(),0}
              }
              } 
            }  else if (ToType(_Ztype.At(1-1)).Included(ToType(C_list.Id())) == CTRUE) { 
            Result = F_Optimize_c_code_method_method1(ToMethod(Core.F__at_property2(C_contain_ask,MakeConstantList(C_list.Id(),C_any.Id())).Id()),MakeConstantList(y,x),_Ztype)
            }  else if (ToType(_Ztype.At(1-1)).Included(ToType(C_set.Id())) == CTRUE) { 
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
/* {1} The go function for: c_code_nth(self:Call) [status=1] */
func F_Optimize_c_code_nth_Call (self *Language.Call ) EID { 
    var Result EID 
    { var l *ClaireList   = self.Args
      { var x *ClaireAny   = l.At(1-1)
        { var p *ClaireProperty   = self.Selector
          { var t *ClaireType  
            var try_1 EID 
            /*g_try(v2:"try_1",loop:false) */
            try_1 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
            /* ERROR PROTECTION INSERTED (t-Result) */
            if ErrorIn(try_1) {Result = try_1
            } else {
            t = ToType(OBJ(try_1))
            { var mt *ClaireType   = Core.F_member_type(t)
              { var r *ClaireAny  
                var try_2 EID 
                /*g_try(v2:"try_2",loop:false) */
                { var arg_3 *ClaireList  
                  _ = arg_3
                  var try_4 EID 
                  /*g_try(v2:"try_4",loop:false) */
                  { 
                    var v_list9 *ClaireList  
                    var u *ClaireAny  
                    var v_local9 *ClaireAny  
                    v_list9 = l
                    try_4 = EID{CreateList(ToType(CEMPTY.Id()),v_list9.Length()).Id(),0}
                    for CLcount := 0; CLcount < v_list9.Length(); CLcount++{ 
                      u = v_list9.At(CLcount)
                      var try_5 EID 
                      /*g_try(v2:"try_5",loop:tuple("try_4", EID)) */
                      try_5 = Core.F_CALL(C_c_type,ARGS(u.ToEID()))
                      /* ERROR PROTECTION INSERTED (v_local9-try_4) */
                      if ErrorIn(try_5) {try_4 = try_5
                      break
                      } else {
                      v_local9 = ANY(try_5)
                      ToList(OBJ(try_4)).PutAt(CLcount,v_local9)
                      } 
                    }
                    } 
                  /* ERROR PROTECTION INSERTED (arg_3-try_2) */
                  if ErrorIn(try_4) {try_2 = try_4
                  } else {
                  arg_3 = ToList(OBJ(try_4))
                  try_2 = F_Optimize_restriction_I_property(p,arg_3,CTRUE).ToEID()
                  }
                  } 
                /* ERROR PROTECTION INSERTED (r-Result) */
                if ErrorIn(try_2) {Result = try_2
                } else {
                r = ANY(try_2)
                if (C_OPT.ToRemove.Contain_ask(x) == CTRUE) { 
                  Result = EID{CNIL.Id(),0}
                  } else {
                  var g0213I *ClaireBoolean  
                  var try_6 EID 
                  /*g_try(v2:"try_6",loop:false) */
                  { 
                    var v_and9 *ClaireBoolean  
                    
                    if (C_table.Id() == x.Isa.Id()) { 
                      { var g0211 *ClaireTable   = ToTable(x)
                        _ = g0211
                        v_and9 = Equal(C_integer.Id(),g0211.Params.Isa.Id())
                        } 
                      } else {
                      v_and9 = CFALSE
                      } 
                    if (v_and9 == CFALSE) {try_6 = EID{CFALSE.Id(),0}
                    } else { 
                      var try_7 EID 
                      /*g_try(v2:"try_7",loop:false) */
                      { 
                        /* Or stat: v="try_7", loop=false */
                        var v_or11 *ClaireBoolean  
                        
                        /* Or stat: try <= @ type_expression(c_type((nth @ list(l,2))),<domain(x):type_expression>) with try:true, v="try_7", loop=false */
                        var try_8 EID 
                        /*g_try(v2:"try_8",loop:false) */
                        { var arg_9 *ClaireType  
                          _ = arg_9
                          var try_10 EID 
                          /*g_try(v2:"try_10",loop:false) */
                          try_10 = Core.F_CALL(C_c_type,ARGS(l.At(2-1).ToEID()))
                          /* ERROR PROTECTION INSERTED (arg_9-try_8) */
                          if ErrorIn(try_10) {try_8 = try_10
                          } else {
                          arg_9 = ToType(OBJ(try_10))
                          try_8 = EID{arg_9.Included(ToType(OBJ(Core.F_CALL(C_domain,ARGS(x.ToEID()))))).Id(),0}
                          }
                          } 
                        /* ERROR PROTECTION INSERTED (v_or11-try_7) */
                        if ErrorIn(try_8) {try_7 = try_8
                        } else {
                        v_or11 = ToBoolean(OBJ(try_8))
                        if (v_or11 == CTRUE) {try_7 = EID{CTRUE.Id(),0}
                        } else { 
                          /* Or stat: try ((= @ any(p,nth)) & (>= @ integer(safety @ meta_compiler(compiler),2))) with try:false, v="try_7", loop=false */
                          v_or11 = MakeBoolean((p.Id() == C_nth.Id()) && (C_compiler.Safety >= 2))
                          if (v_or11 == CTRUE) {try_7 = EID{CTRUE.Id(),0}
                          } else { 
                            try_7 = EID{CFALSE.Id(),0}} 
                          } 
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
                  /* ERROR PROTECTION INSERTED (g0213I-Result) */
                  if ErrorIn(try_6) {Result = try_6
                  } else {
                  g0213I = ToBoolean(OBJ(try_6))
                  if (g0213I == CTRUE) { 
                    { var _CL_obj *Language.CallTable   = Language.To_CallTable(new(Language.CallTable).Is(Language.C_Call_table))
                      _CL_obj.Selector = ToTable(x)
                      /*table->table*//*g_try(v2:"Result",loop:true) */
                      { 
                        var va_arg1 *Language.CallTable  
                        var va_arg2 *ClaireAny  
                        va_arg1 = _CL_obj
                        var try_11 EID 
                        /*g_try(v2:"try_11",loop:false) */
                        try_11 = Core.F_CALL(C_c_code,ARGS(l.At(2-1).ToEID(),EID{C_integer.Id(),0}))
                        /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                        if ErrorIn(try_11) {Result = try_11
                        } else {
                        va_arg2 = ANY(try_11)
                        va_arg1.Arg = va_arg2
                        /*any->any*/Result = va_arg2.ToEID()
                        }
                        } 
                      /* ERROR PROTECTION INSERTED (Result-Result) */
                      if !ErrorIn(Result) {
                      /*g_try(v2:"Result",loop:true) */
                      { 
                        var va_arg1 *Language.CallTable  
                        var va_arg2 *ClaireBoolean  
                        va_arg1 = _CL_obj
                        var try_12 EID 
                        /*g_try(v2:"try_12",loop:false) */
                        { var arg_13 *ClaireBoolean  
                          _ = arg_13
                          var try_14 EID 
                          /*g_try(v2:"try_14",loop:false) */
                          { 
                            /* Or stat: v="try_14", loop=false */
                            var v_or13 *ClaireBoolean  
                            
                            /* Or stat: try % @ list<type_expression>(any, any)(default(x),range(x)) with try:true, v="try_14", loop=false */
                            var try_15 EID 
                            /*g_try(v2:"try_15",loop:false) */
                            try_15 = Core.F_BELONG(ANY(Core.F_CALL(C_default,ARGS(x.ToEID()))),ANY(Core.F_CALL(C_range,ARGS(x.ToEID()))))
                            /* ERROR PROTECTION INSERTED (v_or13-try_14) */
                            if ErrorIn(try_15) {try_14 = try_15
                            } else {
                            v_or13 = ToBoolean(OBJ(try_15))
                            if (v_or13 == CTRUE) {try_14 = EID{CTRUE.Id(),0}
                            } else { 
                              /* Or stat: try = @ any(default(x),0) with try:false, v="try_14", loop=false */
                              v_or13 = Equal(ANY(Core.F_CALL(C_default,ARGS(x.ToEID()))),MakeInteger(0).Id())
                              if (v_or13 == CTRUE) {try_14 = EID{CTRUE.Id(),0}
                              } else { 
                                /* Or stat: try = @ any(p,get) with try:false, v="try_14", loop=false */
                                v_or13 = Equal(p.Id(),C_get.Id())
                                if (v_or13 == CTRUE) {try_14 = EID{CTRUE.Id(),0}
                                } else { 
                                  try_14 = EID{CFALSE.Id(),0}} 
                                } 
                              } 
                            }
                            } 
                          /* ERROR PROTECTION INSERTED (arg_13-try_12) */
                          if ErrorIn(try_14) {try_12 = try_14
                          } else {
                          arg_13 = ToBoolean(OBJ(try_14))
                          try_12 = EID{arg_13.Not.Id(),0}
                          }
                          } 
                        /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                        if ErrorIn(try_12) {Result = try_12
                        } else {
                        va_arg2 = ToBoolean(OBJ(try_12))
                        va_arg1.Test = va_arg2
                        /*boolean->boolean*/Result = EID{va_arg2.Id(),0}
                        }
                        } 
                      /* ERROR PROTECTION INSERTED (Result-Result) */
                      if !ErrorIn(Result) {
                      Result = EID{_CL_obj.Id(),0}
                      }}
                      } 
                    } else {
                    var g0214I *ClaireBoolean  
                    var try_16 EID 
                    /*g_try(v2:"try_16",loop:false) */
                    { 
                      var v_and10 *ClaireBoolean  
                      
                      if (C_table.Id() == x.Isa.Id()) { 
                        { var g0212 *ClaireTable   = ToTable(x)
                          _ = g0212
                          v_and10 = g0212.Params.Isa.IsIn(C_list)
                          } 
                        } else {
                        v_and10 = CFALSE
                        } 
                      if (v_and10 == CFALSE) {try_16 = EID{CFALSE.Id(),0}
                      } else { 
                        v_and10 = Equal(MakeInteger(l.Length()).Id(),MakeInteger(3).Id())
                        if (v_and10 == CFALSE) {try_16 = EID{CFALSE.Id(),0}
                        } else { 
                          var try_17 EID 
                          /*g_try(v2:"try_17",loop:false) */
                          { 
                            /* Or stat: v="try_17", loop=false */
                            var v_or13 *ClaireBoolean  
                            
                            /* Or stat: try <= @ type_expression(tuple! @ list(list(c_type((nth @ list(l,2))), c_type((nth @ list(l,3))))),<domain(x):type_expression>) with try:true, v="try_17", loop=false */
                            var try_18 EID 
                            /*g_try(v2:"try_18",loop:false) */
                            { var arg_19 *ClaireTuple  
                              _ = arg_19
                              var try_20 EID 
                              /*g_try(v2:"try_20",loop:false) */
                              { var arg_21 *ClaireList  
                                _ = arg_21
                                var try_22 EID 
                                /*g_try(v2:"try_22",loop:false) */
                                { 
                                  var v_bag_arg *ClaireAny  
                                  try_22= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                                  var try_23 EID 
                                  /*g_try(v2:"try_23",loop:false) */
                                  try_23 = Core.F_CALL(C_c_type,ARGS(l.At(2-1).ToEID()))
                                  /* ERROR PROTECTION INSERTED (v_bag_arg-try_22) */
                                  if ErrorIn(try_23) {try_22 = try_23
                                  } else {
                                  v_bag_arg = ANY(try_23)
                                  ToList(OBJ(try_22)).AddFast(v_bag_arg)
                                  var try_24 EID 
                                  /*g_try(v2:"try_24",loop:false) */
                                  try_24 = Core.F_CALL(C_c_type,ARGS(l.At(3-1).ToEID()))
                                  /* ERROR PROTECTION INSERTED (v_bag_arg-try_22) */
                                  if ErrorIn(try_24) {try_22 = try_24
                                  } else {
                                  v_bag_arg = ANY(try_24)
                                  ToList(OBJ(try_22)).AddFast(v_bag_arg)}}
                                  } 
                                /* ERROR PROTECTION INSERTED (arg_21-try_20) */
                                if ErrorIn(try_22) {try_20 = try_22
                                } else {
                                arg_21 = ToList(OBJ(try_22))
                                try_20 = EID{arg_21.Tuple_I().Id(),0}
                                }
                                } 
                              /* ERROR PROTECTION INSERTED (arg_19-try_18) */
                              if ErrorIn(try_20) {try_18 = try_20
                              } else {
                              arg_19 = ToTuple(OBJ(try_20))
                              try_18 = EID{ToType(arg_19.Id()).Included(ToType(OBJ(Core.F_CALL(C_domain,ARGS(x.ToEID()))))).Id(),0}
                              }
                              } 
                            /* ERROR PROTECTION INSERTED (v_or13-try_17) */
                            if ErrorIn(try_18) {try_17 = try_18
                            } else {
                            v_or13 = ToBoolean(OBJ(try_18))
                            if (v_or13 == CTRUE) {try_17 = EID{CTRUE.Id(),0}
                            } else { 
                              /* Or stat: try >= @ integer(safety @ meta_compiler(compiler),2) with try:false, v="try_17", loop=false */
                              v_or13 = F__sup_equal_integer(C_compiler.Safety,2)
                              if (v_or13 == CTRUE) {try_17 = EID{CTRUE.Id(),0}
                              } else { 
                                try_17 = EID{CFALSE.Id(),0}} 
                              } 
                            }
                            } 
                          /* ERROR PROTECTION INSERTED (v_and10-try_16) */
                          if ErrorIn(try_17) {try_16 = try_17
                          } else {
                          v_and10 = ToBoolean(OBJ(try_17))
                          if (v_and10 == CFALSE) {try_16 = EID{CFALSE.Id(),0}
                          } else { 
                            try_16 = EID{CTRUE.Id(),0}} 
                          } 
                        } 
                      }
                      } 
                    /* ERROR PROTECTION INSERTED (g0214I-Result) */
                    if ErrorIn(try_16) {Result = try_16
                    } else {
                    g0214I = ToBoolean(OBJ(try_16))
                    if (g0214I == CTRUE) { 
                      { var _CL_obj *Language.CallTable   = Language.To_CallTable(new(Language.CallTable).Is(Language.C_Call_table))
                        _CL_obj.Selector = ToTable(x)
                        /*table->table*//*g_try(v2:"Result",loop:true) */
                        { 
                          var va_arg1 *Language.CallTable  
                          var va_arg2 *ClaireAny  
                          va_arg1 = _CL_obj
                          var try_25 EID 
                          /*g_try(v2:"try_25",loop:false) */
                          { var _CL_obj *Language.List   = Language.To_List(new(Language.List).Is(Language.C_List))
                            /*g_try(v2:"try_25",loop:false) */
                            { 
                              var va_arg1 *Language.Construct  
                              var va_arg2 *ClaireList  
                              va_arg1 = Language.To_Construct(_CL_obj.Id())
                              var try_26 EID 
                              /*g_try(v2:"try_26",loop:false) */
                              { 
                                var v_bag_arg *ClaireAny  
                                try_26= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                                var try_27 EID 
                                /*g_try(v2:"try_27",loop:false) */
                                try_27 = Core.F_CALL(C_c_code,ARGS(l.At(2-1).ToEID(),EID{C_integer.Id(),0}))
                                /* ERROR PROTECTION INSERTED (v_bag_arg-try_26) */
                                if ErrorIn(try_27) {try_26 = try_27
                                } else {
                                v_bag_arg = ANY(try_27)
                                ToList(OBJ(try_26)).AddFast(v_bag_arg)
                                var try_28 EID 
                                /*g_try(v2:"try_28",loop:false) */
                                try_28 = Core.F_CALL(C_c_code,ARGS(l.At(3-1).ToEID(),EID{C_integer.Id(),0}))
                                /* ERROR PROTECTION INSERTED (v_bag_arg-try_26) */
                                if ErrorIn(try_28) {try_26 = try_28
                                } else {
                                v_bag_arg = ANY(try_28)
                                ToList(OBJ(try_26)).AddFast(v_bag_arg)}}
                                } 
                              /* ERROR PROTECTION INSERTED (va_arg2-try_25) */
                              if ErrorIn(try_26) {try_25 = try_26
                              } else {
                              va_arg2 = ToList(OBJ(try_26))
                              va_arg1.Args = va_arg2
                              /*list->list*/try_25 = EID{va_arg2.Id(),0}
                              }
                              } 
                            /* ERROR PROTECTION INSERTED (try_25-try_25) */
                            if !ErrorIn(try_25) {
                            try_25 = EID{_CL_obj.Id(),0}
                            }
                            } 
                          /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                          if ErrorIn(try_25) {Result = try_25
                          } else {
                          va_arg2 = ANY(try_25)
                          va_arg1.Arg = va_arg2
                          /*any->any*/Result = va_arg2.ToEID()
                          }
                          } 
                        /* ERROR PROTECTION INSERTED (Result-Result) */
                        if !ErrorIn(Result) {
                        /*g_try(v2:"Result",loop:true) */
                        { 
                          var va_arg1 *Language.CallTable  
                          var va_arg2 *ClaireBoolean  
                          va_arg1 = _CL_obj
                          var try_29 EID 
                          /*g_try(v2:"try_29",loop:false) */
                          { var arg_30 *ClaireBoolean  
                            _ = arg_30
                            var try_31 EID 
                            /*g_try(v2:"try_31",loop:false) */
                            { 
                              /* Or stat: v="try_31", loop=false */
                              var v_or14 *ClaireBoolean  
                              
                              /* Or stat: try % @ list<type_expression>(any, any)(default(x),range(x)) with try:true, v="try_31", loop=false */
                              var try_32 EID 
                              /*g_try(v2:"try_32",loop:false) */
                              try_32 = Core.F_BELONG(ANY(Core.F_CALL(C_default,ARGS(x.ToEID()))),ANY(Core.F_CALL(C_range,ARGS(x.ToEID()))))
                              /* ERROR PROTECTION INSERTED (v_or14-try_31) */
                              if ErrorIn(try_32) {try_31 = try_32
                              } else {
                              v_or14 = ToBoolean(OBJ(try_32))
                              if (v_or14 == CTRUE) {try_31 = EID{CTRUE.Id(),0}
                              } else { 
                                /* Or stat: try = @ any(default(x),0) with try:false, v="try_31", loop=false */
                                v_or14 = Equal(ANY(Core.F_CALL(C_default,ARGS(x.ToEID()))),MakeInteger(0).Id())
                                if (v_or14 == CTRUE) {try_31 = EID{CTRUE.Id(),0}
                                } else { 
                                  /* Or stat: try = @ any(p,get) with try:false, v="try_31", loop=false */
                                  v_or14 = Equal(p.Id(),C_get.Id())
                                  if (v_or14 == CTRUE) {try_31 = EID{CTRUE.Id(),0}
                                  } else { 
                                    try_31 = EID{CFALSE.Id(),0}} 
                                  } 
                                } 
                              }
                              } 
                            /* ERROR PROTECTION INSERTED (arg_30-try_29) */
                            if ErrorIn(try_31) {try_29 = try_31
                            } else {
                            arg_30 = ToBoolean(OBJ(try_31))
                            try_29 = EID{arg_30.Not.Id(),0}
                            }
                            } 
                          /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                          if ErrorIn(try_29) {Result = try_29
                          } else {
                          va_arg2 = ToBoolean(OBJ(try_29))
                          va_arg1.Test = va_arg2
                          /*boolean->boolean*/Result = EID{va_arg2.Id(),0}
                          }
                          } 
                        /* ERROR PROTECTION INSERTED (Result-Result) */
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
                        _ = arg_33
                        var try_35 EID 
                        /*g_try(v2:"try_35",loop:false) */
                        try_35 = Core.F_CALL(C_c_code,ARGS(x.ToEID(),EID{C_array.Id(),0}))
                        /* ERROR PROTECTION INSERTED (arg_33-Result) */
                        if ErrorIn(try_35) {Result = try_35
                        } else {
                        arg_33 = ANY(try_35)
                        { var arg_34 *ClaireAny  
                          _ = arg_34
                          var try_36 EID 
                          /*g_try(v2:"try_36",loop:false) */
                          try_36 = Core.F_CALL(C_c_code,ARGS(l.At(2-1).ToEID(),EID{C_integer.Id(),0}))
                          /* ERROR PROTECTION INSERTED (arg_34-Result) */
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
                        
                        } 
                      { var arg_37 *ClaireList  
                        _ = arg_37
                        var try_38 EID 
                        /*g_try(v2:"try_38",loop:false) */
                        { 
                          var v_list12 *ClaireList  
                          var x *ClaireAny  
                          var v_local12 *ClaireAny  
                          v_list12 = self.Args
                          try_38 = EID{CreateList(ToType(CEMPTY.Id()),v_list12.Length()).Id(),0}
                          for CLcount := 0; CLcount < v_list12.Length(); CLcount++{ 
                            x = v_list12.At(CLcount)
                            var try_39 EID 
                            /*g_try(v2:"try_39",loop:tuple("try_38", EID)) */
                            try_39 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
                            /* ERROR PROTECTION INSERTED (v_local12-try_38) */
                            if ErrorIn(try_39) {try_38 = try_39
                            break
                            } else {
                            v_local12 = ANY(try_39)
                            ToList(OBJ(try_38)).PutAt(CLcount,v_local12)
                            } 
                          }
                          } 
                        /* ERROR PROTECTION INSERTED (arg_37-Result) */
                        if ErrorIn(try_38) {Result = try_38
                        } else {
                        arg_37 = ToList(OBJ(try_38))
                        Result = F_Optimize_c_code_method_method1(ToMethod(r),self.Args,arg_37)
                        }
                        } 
                      } else {
                      { var arg_40 *ClaireList  
                        _ = arg_40
                        var try_41 EID 
                        /*g_try(v2:"try_41",loop:false) */
                        { 
                          var v_list12 *ClaireList  
                          var x *ClaireAny  
                          var v_local12 *ClaireAny  
                          v_list12 = self.Args
                          try_41 = EID{CreateList(ToType(CEMPTY.Id()),v_list12.Length()).Id(),0}
                          for CLcount := 0; CLcount < v_list12.Length(); CLcount++{ 
                            x = v_list12.At(CLcount)
                            var try_42 EID 
                            /*g_try(v2:"try_42",loop:tuple("try_41", EID)) */
                            try_42 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
                            /* ERROR PROTECTION INSERTED (v_local12-try_41) */
                            if ErrorIn(try_42) {try_41 = try_42
                            break
                            } else {
                            v_local12 = ANY(try_42)
                            ToList(OBJ(try_41)).PutAt(CLcount,v_local12)
                            } 
                          }
                          } 
                        /* ERROR PROTECTION INSERTED (arg_40-Result) */
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
/* {1} The go function for: c_code_table(self:Call) [status=1] */
func F_Optimize_c_code_table_Call (self *Language.Call ) EID { 
    var Result EID 
    { var sp *ClaireProperty   = self.Selector
      { var p *ClaireTable   = ToTable(self.Args.At(1-1))
        { var x *ClaireAny   = self.Args.At(2-1)
          { var y *ClaireAny   = self.Args.At(3-1)
            if (C_OPT.ToRemove.Contain_ask(p.Id()) == CTRUE) { 
              Result = EID{CNIL.Id(),0}
              } else {
              var g0215I *ClaireBoolean  
              var try_1 EID 
              /*g_try(v2:"try_1",loop:false) */
              { 
                /* Or stat: v="try_1", loop=false */
                var v_or7 *ClaireBoolean  
                
                /* Or stat: try = @ any(sp,put) with try:false, v="try_1", loop=false */
                v_or7 = Equal(sp.Id(),C_put.Id())
                if (v_or7 == CTRUE) {try_1 = EID{CTRUE.Id(),0}
                } else { 
                  /* Or stat: try (((<= @ type_expression(c_type(x),domain @ relation(p))) | (>= @ integer(safety @ meta_compiler(compiler),3))) & ((<= @ type_expression(c_type(y),range @ relation(p))) | (>= @ integer(safety @ meta_compiler(compiler),3)))) with try:true, v="try_1", loop=false */
                  var try_2 EID 
                  /*g_try(v2:"try_2",loop:false) */
                  { 
                    var v_and9 *ClaireBoolean  
                    
                    var try_3 EID 
                    /*g_try(v2:"try_3",loop:false) */
                    { 
                      /* Or stat: v="try_3", loop=false */
                      var v_or10 *ClaireBoolean  
                      
                      /* Or stat: try <= @ type_expression(c_type(x),domain @ relation(p)) with try:true, v="try_3", loop=false */
                      var try_4 EID 
                      /*g_try(v2:"try_4",loop:false) */
                      { var arg_5 *ClaireType  
                        _ = arg_5
                        var try_6 EID 
                        /*g_try(v2:"try_6",loop:false) */
                        try_6 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
                        /* ERROR PROTECTION INSERTED (arg_5-try_4) */
                        if ErrorIn(try_6) {try_4 = try_6
                        } else {
                        arg_5 = ToType(OBJ(try_6))
                        try_4 = EID{arg_5.Included(p.Domain).Id(),0}
                        }
                        } 
                      /* ERROR PROTECTION INSERTED (v_or10-try_3) */
                      if ErrorIn(try_4) {try_3 = try_4
                      } else {
                      v_or10 = ToBoolean(OBJ(try_4))
                      if (v_or10 == CTRUE) {try_3 = EID{CTRUE.Id(),0}
                      } else { 
                        /* Or stat: try >= @ integer(safety @ meta_compiler(compiler),3) with try:false, v="try_3", loop=false */
                        v_or10 = F__sup_equal_integer(C_compiler.Safety,3)
                        if (v_or10 == CTRUE) {try_3 = EID{CTRUE.Id(),0}
                        } else { 
                          try_3 = EID{CFALSE.Id(),0}} 
                        } 
                      }
                      } 
                    /* ERROR PROTECTION INSERTED (v_and9-try_2) */
                    if ErrorIn(try_3) {try_2 = try_3
                    } else {
                    v_and9 = ToBoolean(OBJ(try_3))
                    if (v_and9 == CFALSE) {try_2 = EID{CFALSE.Id(),0}
                    } else { 
                      var try_7 EID 
                      /*g_try(v2:"try_7",loop:false) */
                      { 
                        /* Or stat: v="try_7", loop=false */
                        var v_or11 *ClaireBoolean  
                        
                        /* Or stat: try <= @ type_expression(c_type(y),range @ relation(p)) with try:true, v="try_7", loop=false */
                        var try_8 EID 
                        /*g_try(v2:"try_8",loop:false) */
                        { var arg_9 *ClaireType  
                          _ = arg_9
                          var try_10 EID 
                          /*g_try(v2:"try_10",loop:false) */
                          try_10 = Core.F_CALL(C_c_type,ARGS(y.ToEID()))
                          /* ERROR PROTECTION INSERTED (arg_9-try_8) */
                          if ErrorIn(try_10) {try_8 = try_10
                          } else {
                          arg_9 = ToType(OBJ(try_10))
                          try_8 = EID{arg_9.Included(p.Range).Id(),0}
                          }
                          } 
                        /* ERROR PROTECTION INSERTED (v_or11-try_7) */
                        if ErrorIn(try_8) {try_7 = try_8
                        } else {
                        v_or11 = ToBoolean(OBJ(try_8))
                        if (v_or11 == CTRUE) {try_7 = EID{CTRUE.Id(),0}
                        } else { 
                          /* Or stat: try >= @ integer(safety @ meta_compiler(compiler),3) with try:false, v="try_7", loop=false */
                          v_or11 = F__sup_equal_integer(C_compiler.Safety,3)
                          if (v_or11 == CTRUE) {try_7 = EID{CTRUE.Id(),0}
                          } else { 
                            try_7 = EID{CFALSE.Id(),0}} 
                          } 
                        }
                        } 
                      /* ERROR PROTECTION INSERTED (v_and9-try_2) */
                      if ErrorIn(try_7) {try_2 = try_7
                      } else {
                      v_and9 = ToBoolean(OBJ(try_7))
                      if (v_and9 == CFALSE) {try_2 = EID{CFALSE.Id(),0}
                      } else { 
                        try_2 = EID{CTRUE.Id(),0}} 
                      } 
                    }}
                    } 
                  /* ERROR PROTECTION INSERTED (v_or7-try_1) */
                  if ErrorIn(try_2) {try_1 = try_2
                  } else {
                  v_or7 = ToBoolean(OBJ(try_2))
                  if (v_or7 == CTRUE) {try_1 = EID{CTRUE.Id(),0}
                  } else { 
                    try_1 = EID{CFALSE.Id(),0}} 
                  } 
                }
                } 
              /* ERROR PROTECTION INSERTED (g0215I-Result) */
              if ErrorIn(try_1) {Result = try_1
              } else {
              g0215I = ToBoolean(OBJ(try_1))
              if (g0215I == CTRUE) { 
                var g0216I *ClaireBoolean  
                var try_11 EID 
                /*g_try(v2:"try_11",loop:false) */
                { 
                  var v_and8 *ClaireBoolean  
                  
                  var try_12 EID 
                  /*g_try(v2:"try_12",loop:false) */
                  try_12 = F_Optimize_Update_ask_relation1(ToRelation(p.Id()),x,y)
                  /* ERROR PROTECTION INSERTED (v_and8-try_11) */
                  if ErrorIn(try_12) {try_11 = try_12
                  } else {
                  v_and8 = ToBoolean(OBJ(try_12))
                  if (v_and8 == CFALSE) {try_11 = EID{CFALSE.Id(),0}
                  } else { 
                    v_and8 = MakeBoolean((p.Params.Isa.IsIn(C_list) == CTRUE) || (C_integer.Id() == p.Params.Isa.Id()))
                    if (v_and8 == CFALSE) {try_11 = EID{CFALSE.Id(),0}
                    } else { 
                      v_and8 = F__sup_equal_integer(C_compiler.Safety,3)
                      if (v_and8 == CFALSE) {try_11 = EID{CFALSE.Id(),0}
                      } else { 
                        try_11 = EID{CTRUE.Id(),0}} 
                      } 
                    } 
                  }
                  } 
                /* ERROR PROTECTION INSERTED (g0216I-Result) */
                if ErrorIn(try_11) {Result = try_11
                } else {
                g0216I = ToBoolean(OBJ(try_11))
                if (g0216I == CTRUE) { 
                  { var _Zx *ClaireAny  
                    var try_13 EID 
                    /*g_try(v2:"try_13",loop:false) */
                    try_13 = Core.F_CALL(C_c_code,ARGS(x.ToEID(),EID{C_any.Id(),0}))
                    /* ERROR PROTECTION INSERTED (_Zx-Result) */
                    if ErrorIn(try_13) {Result = try_13
                    } else {
                    _Zx = ANY(try_13)
                    { var _Zy *ClaireAny  
                      _ = _Zy
                      var try_14 EID 
                      /*g_try(v2:"try_14",loop:false) */
                      try_14 = Core.F_CALL(C_c_code,ARGS(y.ToEID(),EID{C_any.Id(),0}))
                      /* ERROR PROTECTION INSERTED (_Zy-Result) */
                      if ErrorIn(try_14) {Result = try_14
                      } else {
                      _Zy = ANY(try_14)
                      { var _CL_obj *Language.Update   = Language.To_Update(new(Language.Update).Is(Language.C_Update))
                        _CL_obj.Selector = p.Id()
                        /*any->any*/_CL_obj.Value = _Zy
                        /*any->any*/_CL_obj.Arg = IfThenElse((sp.Id() == C_put.Id()),
                          C_put.Id(),
                          _Zx)
                        /*any->any*/{ 
                          var va_arg1 *Language.Update  
                          var va_arg2 *ClaireAny  
                          va_arg1 = _CL_obj
                          { var _CL_obj *Language.CallTable   = Language.To_CallTable(new(Language.CallTable).Is(Language.C_Call_table))
                            _CL_obj.Selector = p
                            /*table->table*/_CL_obj.Arg = _Zx
                            /*any->any*/_CL_obj.Test = CFALSE
                            /*boolean->boolean*/va_arg2 = _CL_obj.Id()
                            } 
                          va_arg1.ClaireVar = va_arg2
                          /*any->any*/} 
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
/* {1} The go function for: c_code_array(self:Call) [status=1] */
func F_Optimize_c_code_array_Call (self *Language.Call ) EID { 
    var Result EID 
    { var sp *ClaireProperty   = self.Selector
      { var p *ClaireAny   = self.Args.At(1-1)
        { var tp *ClaireType  
          var try_1 EID 
          /*g_try(v2:"try_1",loop:false) */
          try_1 = Core.F_CALL(C_c_type,ARGS(p.ToEID()))
          /* ERROR PROTECTION INSERTED (tp-Result) */
          if ErrorIn(try_1) {Result = try_1
          } else {
          tp = ToType(OBJ(try_1))
          { var mt *ClaireType   = Core.F_member_type(tp)
            { var x *ClaireAny   = self.Args.At(2-1)
              _ = x
              { var y *ClaireAny   = self.Args.At(3-1)
                { var typeok *ClaireBoolean  
                  var try_2 EID 
                  /*g_try(v2:"try_2",loop:false) */
                  { 
                    /* Or stat: v="try_2", loop=false */
                    var v_or9 *ClaireBoolean  
                    
                    /* Or stat: try <= @ type_expression(c_type(y),member @ type(tp)) with try:true, v="try_2", loop=false */
                    var try_3 EID 
                    /*g_try(v2:"try_3",loop:false) */
                    { var arg_4 *ClaireType  
                      _ = arg_4
                      var try_5 EID 
                      /*g_try(v2:"try_5",loop:false) */
                      try_5 = Core.F_CALL(C_c_type,ARGS(y.ToEID()))
                      /* ERROR PROTECTION INSERTED (arg_4-try_3) */
                      if ErrorIn(try_5) {try_3 = try_5
                      } else {
                      arg_4 = ToType(OBJ(try_5))
                      try_3 = EID{arg_4.Included(Core.F_member_type(tp)).Id(),0}
                      }
                      } 
                    /* ERROR PROTECTION INSERTED (v_or9-try_2) */
                    if ErrorIn(try_3) {try_2 = try_3
                    } else {
                    v_or9 = ToBoolean(OBJ(try_3))
                    if (v_or9 == CTRUE) {try_2 = EID{CTRUE.Id(),0}
                    } else { 
                      /* Or stat: try >= @ integer(safety @ meta_compiler(compiler),2) with try:false, v="try_2", loop=false */
                      v_or9 = F__sup_equal_integer(C_compiler.Safety,2)
                      if (v_or9 == CTRUE) {try_2 = EID{CTRUE.Id(),0}
                      } else { 
                        try_2 = EID{CFALSE.Id(),0}} 
                      } 
                    }
                    } 
                  /* ERROR PROTECTION INSERTED (typeok-Result) */
                  if ErrorIn(try_2) {Result = try_2
                  } else {
                  typeok = ToBoolean(OBJ(try_2))
                  { var _Zsel *ClaireAny  
                    var try_6 EID 
                    /*g_try(v2:"try_6",loop:false) */
                    try_6 = Core.F_CALL(C_c_code,ARGS(p.ToEID(),EID{C_array.Id(),0}))
                    /* ERROR PROTECTION INSERTED (_Zsel-Result) */
                    if ErrorIn(try_6) {Result = try_6
                    } else {
                    _Zsel = ANY(try_6)
                    var g0217I *ClaireBoolean  
                    var try_7 EID 
                    /*g_try(v2:"try_7",loop:false) */
                    { 
                      var v_and10 *ClaireBoolean  
                      
                      v_and10 = MakeBoolean((sp.Id() == C_nth_put.Id()) || (typeok == CTRUE))
                      if (v_and10 == CFALSE) {try_7 = EID{CFALSE.Id(),0}
                      } else { 
                        v_and10 = MakeBoolean((mt.Included(ToType(C_float.Id())) == CTRUE) || (Equal(Core.F__exp_type(mt,ToType(C_float.Id())).Id(),CEMPTY.Id()) == CTRUE))
                        if (v_and10 == CFALSE) {try_7 = EID{CFALSE.Id(),0}
                        } else { 
                          var try_8 EID 
                          /*g_try(v2:"try_8",loop:false) */
                          { var arg_9 *ClaireBoolean  
                            _ = arg_9
                            var try_10 EID 
                            /*g_try(v2:"try_10",loop:false) */
                            try_10 = F_Compile_g_throw_any(_Zsel)
                            /* ERROR PROTECTION INSERTED (arg_9-try_8) */
                            if ErrorIn(try_10) {try_8 = try_10
                            } else {
                            arg_9 = ToBoolean(OBJ(try_10))
                            try_8 = EID{arg_9.Not.Id(),0}
                            }
                            } 
                          /* ERROR PROTECTION INSERTED (v_and10-try_7) */
                          if ErrorIn(try_8) {try_7 = try_8
                          } else {
                          v_and10 = ToBoolean(OBJ(try_8))
                          if (v_and10 == CFALSE) {try_7 = EID{CFALSE.Id(),0}
                          } else { 
                            try_7 = EID{CTRUE.Id(),0}} 
                          } 
                        } 
                      }
                      } 
                    /* ERROR PROTECTION INSERTED (g0217I-Result) */
                    if ErrorIn(try_7) {Result = try_7
                    } else {
                    g0217I = ToBoolean(OBJ(try_7))
                    if (g0217I == CTRUE) { 
                      { var _Zx *ClaireAny  
                        _ = _Zx
                        var try_11 EID 
                        /*g_try(v2:"try_11",loop:false) */
                        try_11 = Core.F_CALL(C_c_code,ARGS(x.ToEID(),EID{C_integer.Id(),0}))
                        /* ERROR PROTECTION INSERTED (_Zx-Result) */
                        if ErrorIn(try_11) {Result = try_11
                        } else {
                        _Zx = ANY(try_11)
                        { var _Zy *ClaireAny  
                          _ = _Zy
                          var try_12 EID 
                          /*g_try(v2:"try_12",loop:false) */
                          { var arg_13 *ClaireClass  
                            _ = arg_13
                            if (mt.Included(ToType(C_float.Id())) == CTRUE) { 
                              arg_13 = C_float
                              } else {
                              arg_13 = C_any
                              } 
                            try_12 = Core.F_CALL(C_c_code,ARGS(y.ToEID(),EID{arg_13.Id(),0}))
                            } 
                          /* ERROR PROTECTION INSERTED (_Zy-Result) */
                          if ErrorIn(try_12) {Result = try_12
                          } else {
                          _Zy = ANY(try_12)
                          { var _CL_obj *Language.Update   = Language.To_Update(new(Language.Update).Is(Language.C_Update))
                            _CL_obj.Selector = p
                            /*any->any*/_CL_obj.Value = _Zy
                            /*any->any*/_CL_obj.Arg = C_put.Id()
                            /*any->any*/_CL_obj.ClaireVar = Language.C_Call_array.Make(_Zsel,_Zx,mt.Id())
                            /*any->any*/Result = EID{_CL_obj.Id(),0}
                            } 
                          }
                          } 
                        }
                        } 
                      } else {
                      if (C_compiler.Optimize_ask == CTRUE) { 
                        F_Compile_notice_void()
                        
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
/* {1} The go function for: Update?(p:relation,x:any,y:any) [status=1] */
func F_Optimize_Update_ask_relation1 (p *ClaireRelation ,x *ClaireAny ,y *ClaireAny ) EID { 
    var Result EID 
    { 
      var v_and2 *ClaireBoolean  
      
      v_and2 = Core.F__I_equal_any(p.Id(),C_inverse.Id())
      if (v_and2 == CFALSE) {Result = EID{CFALSE.Id(),0}
      } else { 
        var try_1 EID 
        /*g_try(v2:"try_1",loop:false) */
        { 
          /* Or stat: v="try_1", loop=false */
          var v_or4 *ClaireBoolean  
          
          /* Or stat: try ((not @ any(identical? @ any(if_write @ relation(p),unknown))) & (not @ any(inherit? @ class(owner @ any(if_write @ relation(p)),list)))) with try:false, v="try_1", loop=false */
          v_or4 = MakeBoolean((p.IfWrite != CNULL) && (p.IfWrite.Isa.IsIn(C_list) != CTRUE))
          if (v_or4 == CTRUE) {try_1 = EID{CTRUE.Id(),0}
          } else { 
            /* Or stat: try ((identical? @ any(inverse @ relation(p),unknown)) & (if (= @ any(table,isa @ any(p))) let g0218:table := (<p:table>) in = @ any(integer,owner @ any(params @ table(g0218))) else true) & (if (store? @ relation(p)) ((Compile/designated? @ any(x)) & (Compile/designated? @ any(y)) & (not @ any(multivalued? @ relation(p))) & ((Compile/identifiable? @ any(y)) | (<= @ type_expression(c_type(y),float)))) else true)) with try:true, v="try_1", loop=false */
            var try_2 EID 
            /*g_try(v2:"try_2",loop:false) */
            { 
              var v_and6 *ClaireBoolean  
              
              v_and6 = MakeBoolean((p.Inverse.Id() == CNULL))
              if (v_and6 == CFALSE) {try_2 = EID{CFALSE.Id(),0}
              } else { 
                if (C_table.Id() == p.Isa.Id()) { 
                  { var g0218 *ClaireTable   = ToTable(p.Id())
                    _ = g0218
                    v_and6 = Equal(C_integer.Id(),g0218.Params.Isa.Id())
                    } 
                  } else {
                  v_and6 = CTRUE
                  } 
                if (v_and6 == CFALSE) {try_2 = EID{CFALSE.Id(),0}
                } else { 
                  var try_3 EID 
                  /*g_try(v2:"try_3",loop:false) */
                  if (p.Store_ask == CTRUE) { 
                    { 
                      var v_and10 *ClaireBoolean  
                      
                      var try_4 EID 
                      /*g_try(v2:"try_4",loop:false) */
                      try_4 = F_Compile_designated_ask_any(x)
                      /* ERROR PROTECTION INSERTED (v_and10-try_3) */
                      if ErrorIn(try_4) {try_3 = try_4
                      } else {
                      v_and10 = ToBoolean(OBJ(try_4))
                      if (v_and10 == CFALSE) {try_3 = EID{CFALSE.Id(),0}
                      } else { 
                        var try_5 EID 
                        /*g_try(v2:"try_5",loop:false) */
                        try_5 = F_Compile_designated_ask_any(y)
                        /* ERROR PROTECTION INSERTED (v_and10-try_3) */
                        if ErrorIn(try_5) {try_3 = try_5
                        } else {
                        v_and10 = ToBoolean(OBJ(try_5))
                        if (v_and10 == CFALSE) {try_3 = EID{CFALSE.Id(),0}
                        } else { 
                          v_and10 = p.Multivalued_ask.Not
                          if (v_and10 == CFALSE) {try_3 = EID{CFALSE.Id(),0}
                          } else { 
                            var try_6 EID 
                            /*g_try(v2:"try_6",loop:false) */
                            { 
                              /* Or stat: v="try_6", loop=false */
                              var v_or14 *ClaireBoolean  
                              
                              /* Or stat: try Compile/identifiable? @ any(y) with try:true, v="try_6", loop=false */
                              var try_7 EID 
                              /*g_try(v2:"try_7",loop:false) */
                              try_7 = F_Compile_identifiable_ask_any(y)
                              /* ERROR PROTECTION INSERTED (v_or14-try_6) */
                              if ErrorIn(try_7) {try_6 = try_7
                              } else {
                              v_or14 = ToBoolean(OBJ(try_7))
                              if (v_or14 == CTRUE) {try_6 = EID{CTRUE.Id(),0}
                              } else { 
                                /* Or stat: try <= @ type_expression(c_type(y),float) with try:true, v="try_6", loop=false */
                                var try_8 EID 
                                /*g_try(v2:"try_8",loop:false) */
                                { var arg_9 *ClaireType  
                                  _ = arg_9
                                  var try_10 EID 
                                  /*g_try(v2:"try_10",loop:false) */
                                  try_10 = Core.F_CALL(C_c_type,ARGS(y.ToEID()))
                                  /* ERROR PROTECTION INSERTED (arg_9-try_8) */
                                  if ErrorIn(try_10) {try_8 = try_10
                                  } else {
                                  arg_9 = ToType(OBJ(try_10))
                                  try_8 = EID{arg_9.Included(ToType(C_float.Id())).Id(),0}
                                  }
                                  } 
                                /* ERROR PROTECTION INSERTED (v_or14-try_6) */
                                if ErrorIn(try_8) {try_6 = try_8
                                } else {
                                v_or14 = ToBoolean(OBJ(try_8))
                                if (v_or14 == CTRUE) {try_6 = EID{CTRUE.Id(),0}
                                } else { 
                                  try_6 = EID{CFALSE.Id(),0}} 
                                } 
                              }}
                              } 
                            /* ERROR PROTECTION INSERTED (v_and10-try_3) */
                            if ErrorIn(try_6) {try_3 = try_6
                            } else {
                            v_and10 = ToBoolean(OBJ(try_6))
                            if (v_and10 == CFALSE) {try_3 = EID{CFALSE.Id(),0}
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
                  /* ERROR PROTECTION INSERTED (v_and6-try_2) */
                  if ErrorIn(try_3) {try_2 = try_3
                  } else {
                  v_and6 = ToBoolean(OBJ(try_3))
                  if (v_and6 == CFALSE) {try_2 = EID{CFALSE.Id(),0}
                  } else { 
                    try_2 = EID{CTRUE.Id(),0}} 
                  } 
                } 
              }
              } 
            /* ERROR PROTECTION INSERTED (v_or4-try_1) */
            if ErrorIn(try_2) {try_1 = try_2
            } else {
            v_or4 = ToBoolean(OBJ(try_2))
            if (v_or4 == CTRUE) {try_1 = EID{CTRUE.Id(),0}
            } else { 
              try_1 = EID{CFALSE.Id(),0}} 
            } 
          }
          } 
        /* ERROR PROTECTION INSERTED (v_and2-Result) */
        if ErrorIn(try_1) {Result = try_1
        } else {
        v_and2 = ToBoolean(OBJ(try_1))
        if (v_and2 == CFALSE) {Result = EID{CFALSE.Id(),0}
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
/* {1} The go function for: Update?(p:relation,s:relation) [status=0] */
func F_Optimize_Update_ask_relation2 (p *ClaireRelation ,s *ClaireRelation ) *ClaireBoolean  { 
    if ((p.IfWrite != CNULL) && 
        ((p.IfWrite.Isa.IsIn(C_list) != CTRUE) && 
          (s.Id() == C_add.Id()))) {return CTRUE
    } else {return CFALSE}} 
  
// The EID go function for: Update? @ list<type_expression>(relation, relation) (throw: false) 
func E_Optimize_Update_ask_relation2 (p EID,s EID) EID { 
    return EID{F_Optimize_Update_ask_relation2(ToRelation(OBJ(p)),ToRelation(OBJ(s)) ).Id(),0}} 
  
// Update returns the value .. <yc:0.01 -> needed in CLAIRE 2.4 !!!>
/* {1} The go function for: c_type(self:Update) [status=0] */
func F_c_type_Update (self *Language.Update ) *ClaireType  { 
    return  ToType(C_void.Id())
    } 
  
// The EID go function for: c_type @ Update (throw: false) 
func E_c_type_Update (self EID) EID { 
    return EID{F_c_type_Update(Language.To_Update(OBJ(self)) ).Id(),0}} 
  
// in CLAIRE4 we isolate this case (call the if-write demon) because it may produce an error
/* {1} The go function for: Compile/update_write?(self:Update) [status=0] */
func F_Compile_update_write_ask_Update (self *Language.Update ) *ClaireBoolean  { 
    // procedure body with s = boolean 
var Result *ClaireBoolean  
    { var p *ClaireAny   = self.Selector
      { var a *ClaireAny   = self.Arg
        if (p.Isa.IsIn(C_relation) == CTRUE) { 
          { var g0220 *ClaireRelation   = ToRelation(p)
            _ = g0220
            Result = MakeBoolean((g0220.IfWrite != CNULL) && (a != C_put.Id()) && (a != Core.C_put_store.Id()))
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
/* {1} The go function for: c_code_method(self:method,l:list,%type:list) [status=1] */
func F_Optimize_c_code_method_method1 (self *ClaireMethod ,l *ClaireList ,_Ztype *ClaireList ) EID { 
    var Result EID 
    { var arg_1 *ClaireClass  
      _ = arg_1
      var try_2 EID 
      /*g_try(v2:"try_2",loop:false) */
      try_2 = F_Optimize_c_srange_method(self)
      /* ERROR PROTECTION INSERTED (arg_1-Result) */
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
  
/* {1} The go function for: c_code_method(self:method,l:list,%type:list,sx:class) [status=1] */
func F_Optimize_c_code_method_method2 (self *ClaireMethod ,l *ClaireList ,_Ztype *ClaireList ,sx *ClaireClass ) EID { 
    var Result EID 
    if ((self.Module_I.Id() != C_claire.Id()) || 
        ((C_compiler.Safety >= 2) || 
          (self.Functional.Id() != CNULL))) { 
      { var ld *ClaireList   = self.Domain
        { var n int  = ld.Length()
          if (n != l.Length()) { 
            { var arg_1 *ClaireList  
              _ = arg_1
              { var i_bag *ClaireList   = ToType(CEMPTY.Id()).EmptyList()
                { var i int  = 1
                  { var g0222 int  = (n-1)
                    _ = g0222
                    for (i <= g0222) { 
                      /* While stat, v:"arg_1" loop:true */
                      i_bag.AddFast(l.At(i-1))/*t=any,s=void*/
                      i = (i+1)
                      /* try?:false, v2:"v_while10" loop will be:tuple("arg_1", void) */
                      } 
                    } 
                  } 
                arg_1 = i_bag
                } 
              { var arg_2 *ClaireList  
                _ = arg_2
                { var i_bag *ClaireList   = ToType(CEMPTY.Id()).EmptyList()
                  { var i int  = n
                    { var g0223 int  = l.Length()
                      _ = g0223
                      for (i <= g0223) { 
                        /* While stat, v:"arg_2" loop:true */
                        i_bag.AddFast(l.At(i-1))/*t=any,s=void*/
                        i = (i+1)
                        /* try?:false, v2:"v_while11" loop will be:tuple("arg_2", void) */
                        } 
                      } 
                    } 
                  arg_2 = i_bag
                  } 
                l = arg_1.AddFast(arg_2.Id())/*t=any,s=list*/
                } 
              } 
            } 
          var g0225I *ClaireBoolean  
          var try_3 EID 
          /*g_try(v2:"try_3",loop:false) */
          { 
            var v_and5 *ClaireBoolean  
            
            v_and5 = self.Inline_ask
            if (v_and5 == CFALSE) {try_3 = EID{CFALSE.Id(),0}
            } else { 
              var try_4 EID 
              /*g_try(v2:"try_4",loop:false) */
              try_4 = F_Optimize_c_inline_ask_method(self,l)
              /* ERROR PROTECTION INSERTED (v_and5-try_3) */
              if ErrorIn(try_4) {try_3 = try_4
              } else {
              v_and5 = ToBoolean(OBJ(try_4))
              if (v_and5 == CFALSE) {try_3 = EID{CFALSE.Id(),0}
              } else { 
                try_3 = EID{CTRUE.Id(),0}} 
              } 
            }
            } 
          /* ERROR PROTECTION INSERTED (g0225I-Result) */
          if ErrorIn(try_3) {Result = try_3
          } else {
          g0225I = ToBoolean(OBJ(try_3))
          if (g0225I == CTRUE) { 
            Result = F_Optimize_c_inline_method1(self,l,sx)
            } else {
            { var arg_5 *ClaireList  
              _ = arg_5
              var try_6 EID 
              /*g_try(v2:"try_6",loop:false) */
              { var i_bag *ClaireList   = ToType(CEMPTY.Id()).EmptyList()
                /*g_try(v2:"try_6",loop:false) */
                { var i int  = 1
                  { var g0224 int  = n
                    _ = g0224
                    try_6= EID{CFALSE.Id(),0}
                    for (i <= g0224) { 
                      /* While stat, v:"try_6" loop:false */
                      var loop_7 EID 
                      _ = loop_7
                      { 
                      /*g_try(v2:"loop_7",loop:tuple("try_6", EID)) */
                      { var arg_8 *ClaireAny  
                        _ = arg_8
                        var try_9 EID 
                        /*g_try(v2:"try_9",loop:false) */
                        try_9 = F_Compile_c_strict_code_any(l.At(i-1),F_Compile_psort_any(ld.ValuesO()[i-1]))
                        /* ERROR PROTECTION INSERTED (arg_8-loop_7) */
                        if ErrorIn(try_9) {loop_7 = try_9
                        } else {
                        arg_8 = ANY(try_9)
                        loop_7 = EID{i_bag.AddFast(arg_8).Id(),0}/*t=any,s=EID*/
                        }
                        } 
                      /* ERROR PROTECTION INSERTED (loop_7-loop_7) */
                      if ErrorIn(loop_7) {try_6 = loop_7
                      break
                      } else {
                      i = (i+1)
                      }
                      /* try?:false, v2:"v_while10" loop will be:tuple("try_6", EID) */
                      } 
                    }
                    } 
                  } 
                /* ERROR PROTECTION INSERTED (try_6-try_6) */
                if !ErrorIn(try_6) {
                try_6 = EID{i_bag.Id(),0}
                }
                } 
              /* ERROR PROTECTION INSERTED (arg_5-Result) */
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
/* {1} The go function for: Call_method!(self:method,%code:list) [status=0] */
func F_Optimize_Call_method_I_method (self *ClaireMethod ,_Zcode *ClaireList ) *ClaireAny  { 
    // procedure body with s = any 
var Result *ClaireAny  
    if (F_Optimize_legal_ask_module(self.Module_I,self.Id()) != CTRUE) { 
      Core.F_tformat_string(MakeString("in call ~S~S\n"),0,MakeConstantList(self.Selector.Id(),_Zcode.Id()))
      } 
    if (_Zcode.Length() == 1) { 
      { var _CL_obj *Language.CallMethod1   = Language.To_CallMethod1(new(Language.CallMethod1).Is(Language.C_Call_method1))
        _CL_obj.Arg = self
        /*method->method*/_CL_obj.Args = _Zcode
        /*list->list*/Result = _CL_obj.Id()
        } 
      }  else if (_Zcode.Length() == 2) { 
      { var _CL_obj *Language.CallMethod2   = Language.To_CallMethod2(new(Language.CallMethod2).Is(Language.C_Call_method2))
        _CL_obj.Arg = self
        /*method->method*/_CL_obj.Args = _Zcode
        /*list->list*/Result = _CL_obj.Id()
        } 
      } else {
      { var _CL_obj *Language.CallMethod   = Language.To_CallMethod(new(Language.CallMethod).Is(Language.C_Call_method))
        _CL_obj.Arg = self
        /*method->method*/_CL_obj.Args = _Zcode
        /*list->list*/Result = _CL_obj.Id()
        } 
      } 
    return Result} 
  
// The EID go function for: Call_method! @ method (throw: false) 
func E_Optimize_Call_method_I_method (self EID,_Zcode EID) EID { 
    return F_Optimize_Call_method_I_method(ToMethod(OBJ(self)),ToList(OBJ(_Zcode)) ).ToEID()} 
  
// a call_method or a call external has an obvious type (we do not need to do
// better ?)
/* {1} The go function for: c_type(self:Call_method) [status=1] */
func F_c_type_Call_method (self *Language.CallMethod ) EID { 
    var Result EID 
    { var arg_1 *ClaireList  
      _ = arg_1
      var try_2 EID 
      /*g_try(v2:"try_2",loop:false) */
      { 
        var v_list3 *ClaireList  
        var x *ClaireAny  
        var v_local3 *ClaireAny  
        v_list3 = self.Args
        try_2 = EID{CreateList(ToType(CEMPTY.Id()),v_list3.Length()).Id(),0}
        for CLcount := 0; CLcount < v_list3.Length(); CLcount++{ 
          x = v_list3.At(CLcount)
          var try_3 EID 
          /*g_try(v2:"try_3",loop:tuple("try_2", EID)) */
          try_3 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
          /* ERROR PROTECTION INSERTED (v_local3-try_2) */
          if ErrorIn(try_3) {try_2 = try_3
          break
          } else {
          v_local3 = ANY(try_3)
          ToList(OBJ(try_2)).PutAt(CLcount,v_local3)
          } 
        }
        } 
      /* ERROR PROTECTION INSERTED (arg_1-Result) */
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
/* {1} The go function for: c_code(self:Call_method) [status=1] */
func F_c_code_Call_method (self *Language.CallMethod ) EID { 
    var Result EID 
    { var m *ClaireMethod   = self.Arg
      { var ld *ClaireList   = m.Domain
        { var n int  = Reader.F_min_integer(self.Args.Length(),ld.Length())
          _ = n
          { var arg_1 *ClaireList  
            _ = arg_1
            var try_2 EID 
            /*g_try(v2:"try_2",loop:false) */
            { var i_bag *ClaireList   = ToType(CEMPTY.Id()).EmptyList()
              /*g_try(v2:"try_2",loop:false) */
              { var i int  = 1
                { var g0226 int  = n
                  _ = g0226
                  try_2= EID{CFALSE.Id(),0}
                  for (i <= g0226) { 
                    /* While stat, v:"try_2" loop:false */
                    var loop_3 EID 
                    _ = loop_3
                    { 
                    /*g_try(v2:"loop_3",loop:tuple("try_2", EID)) */
                    { var arg_4 *ClaireAny  
                      _ = arg_4
                      var try_5 EID 
                      /*g_try(v2:"try_5",loop:false) */
                      try_5 = F_Compile_c_strict_code_any(self.Args.At(i-1),F_Compile_psort_any(ld.ValuesO()[i-1]))
                      /* ERROR PROTECTION INSERTED (arg_4-loop_3) */
                      if ErrorIn(try_5) {loop_3 = try_5
                      } else {
                      arg_4 = ANY(try_5)
                      loop_3 = EID{i_bag.AddFast(arg_4).Id(),0}/*t=any,s=EID*/
                      }
                      } 
                    /* ERROR PROTECTION INSERTED (loop_3-loop_3) */
                    if ErrorIn(loop_3) {try_2 = loop_3
                    break
                    } else {
                    i = (i+1)
                    }
                    /* try?:false, v2:"v_while9" loop will be:tuple("try_2", EID) */
                    } 
                  }
                  } 
                } 
              /* ERROR PROTECTION INSERTED (try_2-try_2) */
              if !ErrorIn(try_2) {
              try_2 = EID{i_bag.Id(),0}
              }
              } 
            /* ERROR PROTECTION INSERTED (arg_1-Result) */
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
/* {1} The go function for: Compile/functional!(self:method) [status=1] */
func F_Compile_functional_I_method (self *ClaireMethod ) EID { 
    var Result EID 
    { var f *ClaireAny   = Core.F_get_property(C_functional,ToObject(self.Id()))
      { var p *ClaireProperty   = self.Selector
        _ = p
        if (C_function.Id() == f.Isa.Id()) { 
          { var g0227 *ClaireFunction   = ToFunction(f)
            _ = g0227
            Result = EID{g0227.Id(),0}
            } 
          } else {
          { var arg_1 *ClaireAny  
            _ = arg_1
            var try_2 EID 
            /*g_try(v2:"try_2",loop:false) */
            try_2 = Core.F_CALL(C_Compile_function_name,ARGS(EID{p.Id(),0},EID{self.Domain.Id(),0},f.ToEID()))
            /* ERROR PROTECTION INSERTED (arg_1-Result) */
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
/* {1} The go function for: nth_type_check(tl:type,ti:type,tx:type) [status=0] */
func F_Optimize_nth_type_check_type (tl *ClaireType ,ti *ClaireType ,tx *ClaireType ) *ClaireAny  { 
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
/* {1} The go function for: c_inline?(self:method,l:list) [status=1] */
func F_Optimize_c_inline_ask_method (self *ClaireMethod ,l *ClaireList ) EID { 
    var Result EID 
    { var f *ClaireLambda   = self.Formula
      { var la *ClaireList   = f.Vars
        _ = la
        { var x *ClaireAny   = f.Body
          _ = x
          { var n int  = 1
            _ = n
            { var arg_1 *ClaireAny  
              _ = arg_1
              var try_2 EID 
              /*g_try(v2:"try_2",loop:false) */
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
                  /*g_try(v2:"loop_3",loop:tuple("try_2", EID)) */
                  var g0229I *ClaireBoolean  
                  var try_4 EID 
                  /*g_try(v2:"try_4",loop:false) */
                  { 
                    var v_and9 *ClaireBoolean  
                    
                    v_and9 = Core.F__sup_integer(Language.F_occurrence_any(x,To_Variable(v)),1)
                    if (v_and9 == CFALSE) {try_4 = EID{CFALSE.Id(),0}
                    } else { 
                      var try_5 EID 
                      /*g_try(v2:"try_5",loop:false) */
                      { var arg_6 *ClaireBoolean  
                        _ = arg_6
                        var try_7 EID 
                        /*g_try(v2:"try_7",loop:false) */
                        try_7 = F_Compile_designated_ask_any(l.At(n-1))
                        /* ERROR PROTECTION INSERTED (arg_6-try_5) */
                        if ErrorIn(try_7) {try_5 = try_7
                        } else {
                        arg_6 = ToBoolean(OBJ(try_7))
                        try_5 = EID{arg_6.Not.Id(),0}
                        }
                        } 
                      /* ERROR PROTECTION INSERTED (v_and9-try_4) */
                      if ErrorIn(try_5) {try_4 = try_5
                      } else {
                      v_and9 = ToBoolean(OBJ(try_5))
                      if (v_and9 == CFALSE) {try_4 = EID{CFALSE.Id(),0}
                      } else { 
                        v_and9 = ANY(Core.F_CALL(C_range,ARGS(v.ToEID()))).Isa.IsIn(C_Optimize_Pattern).Not
                        if (v_and9 == CFALSE) {try_4 = EID{CFALSE.Id(),0}
                        } else { 
                          try_4 = EID{CTRUE.Id(),0}} 
                        } 
                      } 
                    }
                    } 
                  /* ERROR PROTECTION INSERTED (g0229I-loop_3) */
                  if ErrorIn(try_4) {loop_3 = try_4
                  } else {
                  g0229I = ToBoolean(OBJ(try_4))
                  if (g0229I == CTRUE) { 
                    try_2 = EID{CTRUE.Id(),0}
                    break
                    } else {
                    n = (n+1)
                    loop_3 = EID{C__INT,IVAL(n)}
                    } 
                  }
                  /* ERROR PROTECTION INSERTED (loop_3-try_2) */
                  if ErrorIn(loop_3) {try_2 = loop_3
                  break
                  } else {
                  }
                  } 
                } 
              /* ERROR PROTECTION INSERTED (arg_1-Result) */
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
  
// checks if a special optization restriction can be used (with patterns)
/* {1} The go function for: inline_optimize?(self:Call) [status=1] */
func F_Optimize_inline_optimize_ask_Call (self *Language.Call ) EID { 
    var Result EID 
    { var l *ClaireList   = self.Args
      { var m *ClaireAny  
        { var arg_1 *ClaireList  
          _ = arg_1
          { 
            var v_list5 *ClaireList  
            var x *ClaireAny  
            var v_local5 *ClaireAny  
            v_list5 = l
            arg_1 = CreateList(ToType(CEMPTY.Id()),v_list5.Length())
            for CLcount := 0; CLcount < v_list5.Length(); CLcount++{ 
              x = v_list5.At(CLcount)
              v_local5 = MakeConstantSet(x).Id()
              arg_1.PutAt(CLcount,v_local5)
              } 
            } 
          m = F_Optimize_restriction_I_property(self.Selector,arg_1,CTRUE)
          } 
        if (C_method.Id() == m.Isa.Id()) { 
          { var g0230 *ClaireMethod   = ToMethod(m)
            var g0232I *ClaireBoolean  
            var try_2 EID 
            /*g_try(v2:"try_2",loop:false) */
            { 
              var v_and6 *ClaireBoolean  
              
              v_and6 = g0230.Inline_ask
              if (v_and6 == CFALSE) {try_2 = EID{CFALSE.Id(),0}
              } else { 
                { var arg_3 *ClaireAny  
                  _ = arg_3
                  { 
                    var s *ClaireTypeExpression  
                    _ = s
                    var s_iter *ClaireAny  
                    arg_3= CFALSE.Id()
                    for _,s_iter = range(g0230.Domain.ValuesO()){ 
                      s = ToTypeExpression(s_iter)
                      if (s.Isa.IsIn(C_Optimize_Pattern) == CTRUE) { 
                        arg_3 = CTRUE.Id()
                        break
                        } 
                      } 
                    } 
                  v_and6 = F_boolean_I_any(arg_3)
                  } 
                if (v_and6 == CFALSE) {try_2 = EID{CFALSE.Id(),0}
                } else { 
                  var try_4 EID 
                  /*g_try(v2:"try_4",loop:false) */
                  try_4 = F_Optimize_c_inline_ask_method(g0230,l)
                  /* ERROR PROTECTION INSERTED (v_and6-try_2) */
                  if ErrorIn(try_4) {try_2 = try_4
                  } else {
                  v_and6 = ToBoolean(OBJ(try_4))
                  if (v_and6 == CFALSE) {try_2 = EID{CFALSE.Id(),0}
                  } else { 
                    try_2 = EID{CTRUE.Id(),0}} 
                  } 
                } 
              }
              } 
            /* ERROR PROTECTION INSERTED (g0232I-Result) */
            if ErrorIn(try_2) {Result = try_2
            } else {
            g0232I = ToBoolean(OBJ(try_2))
            if (g0232I == CTRUE) { 
              Result = EID{g0230.Id(),0}
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