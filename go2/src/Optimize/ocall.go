/***** CLAIRE Compilation of file /Users/ycaseau/Dropbox/src/clairev4.0/src/compile/ocall.cl 
         [version 4.0.01 / safety 5] Saturday 10-30-2021 *****/

package Optimize
import (_ "fmt"
	. "Kernel"
	"Core"
	"Language"
	"Reader"
)

//-------- dumb function to prevent import errors --------
func import_g0478() { 
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
/* {1} OPT.The go function for: restriction!(self:property,l:list,mode:boolean) [] */
func F_Optimize_restriction_I_property (self *ClaireProperty ,l *ClaireList ,mode *ClaireBoolean ) *ClaireAny  { 
    // use function body compiling 
/* Let:2 */{ 
      var i int  = 1
      /* noccur = 5 */
      /* Let:3 */{ 
        var g0479 int  = l.Length()
        /* noccur = 1 */
        for (i <= g0479) /* while:4 */{ 
          ToArray(l.Id()).NthPut(i,F_Optimize_ptype_type(ToType(l.At(i-1))).Id())
          i = (i+1)
          /* while-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return  F_Optimize_restriction_I_list(self.Definition,l,mode)
    } 
  
// The EID go function for: restriction! @ property (throw: false) 
func E_Optimize_restriction_I_property (self EID,l EID,mode EID) EID { 
    return /*(sm for restriction! @ property= any)*/ F_Optimize_restriction_I_property(ToProperty(OBJ(self)),ToList(OBJ(l)),ToBoolean(OBJ(mode)) ).ToEID()} 
  
// finds a suitable restriction in lr. Returns a restriction for a match,
// list(r) for a possible match (unique), () for no match and ambiguous
// otherwise
// CLAIRE4 : we define "open required" based on the property.
/* {1} OPT.The go function for: restriction!(lr:list,l:list,mode:boolean) [] */
func F_Optimize_restriction_I_list (lr *ClaireList ,l *ClaireList ,mode *ClaireBoolean ) *ClaireAny  { 
    // procedure body with s =  
var Result *ClaireAny  
    
    /* Let:2 */{ 
      var open_required *ClaireBoolean   = MakeBoolean((lr.Length() > 0) && (Equal(ANY(Core.F_CALL(C_selector,ARGS(lr.At(1-1).ToEID()))),MakeInteger(3).Id()) == CTRUE))
      /* noccur = 1 */
      /* Let:3 */{ 
        var rep *ClaireAny   = CEMPTY.Id()
        /* noccur = 9 */
        /* For:4 */{ 
          var r *ClaireAny  
          _ = r
          var r_support *ClaireList  
          r_support = lr
          r_len := r_support.Length()
          for i_it := 0; i_it < r_len; i_it++ { 
            r = r_support.At(i_it)
            if ((F_boolean_I_any(rep).Id() != CTRUE.Id()) && 
                (Core.F_tmatch_ask_list(l,ToRestriction(r).Domain) == CTRUE)) /* If:6 */{ 
              if (mode == CTRUE) /* If:7 */{ 
                rep = r
                } else {
                rep = ToRestriction(r).Range.Id()
                /* If-7 */} 
               /*v = Result, s =void*/

              break
              /* If!6 */}  else if (Core.F__exp_list(ToRestriction(r).Domain,l).Length() != 0) /* If:6 */{ 
              if (mode != CTRUE) /* If:7 */{ 
                
                rep = Core.F_U_type(ToType(rep),ToRestriction(r).Range).Id()
                /* If!7 */}  else if ((C_compiler.Safety <= 3) || 
                  ((Equal(rep,CEMPTY.Id()) != CTRUE) || 
                    (open_required == CTRUE))) /* If:7 */{ 
                rep = C_Optimize_ambiguous.Id()
                 /*v = Result, s =void*/

                break
                } else {
                rep = r
                /* If-7 */} 
              /* If-6 */} 
            /* loop-5 */} 
          /* For-4 */} 
        
        Result = rep
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: restriction! @ list (throw: false) 
func E_Optimize_restriction_I_list (lr EID,l EID,mode EID) EID { 
    return /*(sm for restriction! @ list= any)*/ F_Optimize_restriction_I_list(ToList(OBJ(lr)),ToList(OBJ(l)),ToBoolean(OBJ(mode)) ).ToEID()} 
  
// special version for Super, which only looks at methods with domains
// bigger than c
/* {1} OPT.The go function for: restriction!(c:class,lr:list,l:list) [] */
func F_Optimize_restriction_I_class (c *ClaireClass ,lr *ClaireList ,l *ClaireList ) *ClaireAny  { 
    // procedure body with s =  
var Result *ClaireAny  
    if (C_compiler.Safety > 3) /* If:2 */{ 
      ToArray(l.Id()).NthPut(1,Core.F__exp_type(ToType(c.Id()),ToType(l.At(1-1))).Id())
      /* If-2 */} 
    /* For:2 */{ 
      var r *ClaireAny  
      _ = r
      Result= CFALSE.Id()
      var r_support *ClaireList  
      r_support = lr
      r_len := r_support.Length()
      for i_it := 0; i_it < r_len; i_it++ { 
        r = r_support.At(i_it)
        if (ToType(c.Id()).Included(ToType(ToRestriction(r).Domain.ValuesO()[1-1])) == CTRUE) /* If:4 */{ 
          if (Core.F_tmatch_ask_list(l,ToRestriction(r).Domain) == CTRUE) /* If:5 */{ 
             /*v = Result, s =any*/
Result = r
            break
            /* If!5 */}  else if (Core.F__exp_list(ToRestriction(r).Domain,l).Length() != 0) /* If:5 */{ 
             /*v = Result, s =any*/
Result = C_Optimize_ambiguous.Id()
            break
            /* If-5 */} 
          /* If-4 */} 
        /* loop-3 */} 
      /* For-2 */} 
    return Result} 
  
// The EID go function for: restriction! @ class (throw: false) 
func E_Optimize_restriction_I_class (c EID,lr EID,l EID) EID { 
    return /*(sm for restriction! @ class= any)*/ F_Optimize_restriction_I_class(ToClass(OBJ(c)),ToList(OBJ(lr)),ToList(OBJ(l)) ).ToEID()} 
  
// uses a second order type {property + function}
/* {1} OPT.The go function for: use_range(self:method,%l:list) [] */
func F_Optimize_use_range_method (self *ClaireMethod ,_Zl *ClaireList ) EID { 
    var Result EID 
    if ((self.Inline_ask == CTRUE) && 
        (self.Typing == CNULL)) /* If:2 */{ 
      /* Let:3 */{ 
        var lv *ClaireList   = self.Formula.Vars
        /* noccur = 3 */
        /* Let:4 */{ 
          var _Zt *ClaireType   = ToType(C_any.Id())
          /* noccur = 5 */
          /* Let:5 */{ 
            var _Zl2 *ClaireList  
            /* noccur = 1 */
            /* Iteration:6 */{ 
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
              /* Iteration-6 */} 
            /* For:6 */{ 
              var v *ClaireAny  
              _ = v
              Result= EID{CFALSE.Id(),0}
              var v_support *ClaireList  
              v_support = lv
              v_len := v_support.Length()
              for i_it := 0; i_it < v_len; i_it++ { 
                v = v_support.At(i_it)
                var void_try8 EID 
                _ = void_try8
                void_try8 = Core.F_put_property2(C_range,ToObject(v),_Zl.At((INT(Core.F_CALL(C_mClaire_index,ARGS(v.ToEID())))+1)-1))
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
            var _Zt_try04846 EID 
            _Zt_try04846 = Core.F_CALL(C_c_type,ARGS(self.Formula.Body.ToEID()))
            /* ERROR PROTECTION INSERTED (_Zt-Result) */
            if ErrorIn(_Zt_try04846) {Result = _Zt_try04846
            } else {
            _Zt = ToType(OBJ(_Zt_try04846))
            Result = EID{_Zt.Id(),0}
            /* For:6 */{ 
              var v *ClaireAny  
              _ = v
              Result= EID{CFALSE.Id(),0}
              var v_support *ClaireList  
              v_support = lv
              v_len := v_support.Length()
              for i_it := 0; i_it < v_len; i_it++ { 
                v = v_support.At(i_it)
                var void_try8 EID 
                _ = void_try8
                void_try8 = Core.F_put_property2(C_range,ToObject(v),_Zl2.At((INT(Core.F_CALL(C_mClaire_index,ARGS(v.ToEID())))+1)-1))
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
            if (self.Range.Isa.IsIn(C_type) == CTRUE) /* If:6 */{ 
              var _Zt_try04857 EID 
              _Zt_try04857 = Core.F_CALL(ToProperty(C__exp.Id()),ARGS(EID{_Zt.Id(),0},EID{self.Range.Id(),0}))
              /* ERROR PROTECTION INSERTED (_Zt-Result) */
              if ErrorIn(_Zt_try04857) {Result = _Zt_try04857
              } else {
              _Zt = ToType(OBJ(_Zt_try04857))
              Result = EID{_Zt.Id(),0}
              }
              } else {
              Result = EID{CFALSE.Id(),0}
              /* If-6 */} 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            if (F_boolean_I_any(_Zt.Id()).Id() != CTRUE.Id()) /* If:6 */{ 
              /* Let:7 */{ 
                var g0486UU *ClaireType  
                /* noccur = 1 */
                var g0486UU_try04878 EID 
                g0486UU_try04878 = Core.F_CALL(C_c_type,ARGS(self.Formula.Body.ToEID()))
                /* ERROR PROTECTION INSERTED (g0486UU-Result) */
                if ErrorIn(g0486UU_try04878) {Result = g0486UU_try04878
                } else {
                g0486UU = ToType(OBJ(g0486UU_try04878))
                Result = Core.F_CALL(C_Compile_Cerror,ARGS(EID{MakeString("[207] inline ~S: range ~S is incompatible with ~S (inferred)").Id(),0},
                  EID{self.Id(),0},
                  EID{self.Range.Id(),0},
                  EID{g0486UU.Id(),0}))
                }
                /* Let-7 */} 
              } else {
              Result = EID{CFALSE.Id(),0}
              /* If-6 */} 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            Result = EID{_Zt.Id(),0}
            }}}}}
            /* Let-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      } else {
      
      /* Let:3 */{ 
        var f *ClaireAny   = self.Typing
        /* noccur = 6 */
        /* Let:4 */{ 
          var _Zl2 *ClaireList  
          /* noccur = 3 */
          /* Iteration:5 */{ 
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
            /* Iteration-5 */} 
          /* Let:5 */{ 
            var _Zt1 *ClaireType   = self.Range
            /* noccur = 5 */
            /* Let:6 */{ 
              var _Zt2 *ClaireType  
              /* noccur = 3 */
              var _Zt2_try7 EID 
              h_index := ClEnv.Index /* Handle */
              h_base := ClEnv.Base
              if (f.Isa.IsIn(C_lambda) == CTRUE) /* If:7 */{ 
                /* Let:8 */{ 
                  var g0480 *ClaireLambda   = ToLambda(f)
                  /* noccur = 1 */
                  _Zt2_try7 = Language.F_apply_lambda(g0480,_Zl2)
                  /* Let-8 */} 
                /* If!7 */}  else if (f.Isa.IsIn(C_property) == CTRUE) /* If:7 */{ 
                /* Let:8 */{ 
                  var g0481 *ClaireProperty   = ToProperty(f)
                  /* noccur = 1 */
                  _Zt2_try7 = Core.F_apply_property(g0481,_Zl2)
                  /* Let-8 */} 
                /* If!7 */}  else if (C_function.Id() == f.Isa.Id()) /* If:7 */{ 
                /* Let:8 */{ 
                  var g0482 *ClaireFunction   = ToFunction(f)
                  /* noccur = 1 */
                  _Zt2_try7 = F_apply_function(g0482,_Zl2)
                  /* Let-8 */} 
                } else {
                _Zt2_try7 = EID{_Zt1.Id(),0}
                /* If-7 */} 
              if ErrorIn(_Zt2_try7){ 
                /* s=type */ClEnv.Index = h_index
                ClEnv.Base = h_base
                Core.F_tformat_string(MakeString("~S's 2nd-order type failed on ~S \n"),0,MakeConstantList(self.Id(),_Zl.Id()))
                _Zt2 = _Zt1
                } else {
                _Zt2 = ToType(OBJ(_Zt2_try7))
                } 
              if ((F_boolean_I_any(F_Compile_sort_equal_class(F_Compile_osort_any(_Zt1.Id()),F_Compile_osort_any(_Zt2.Id()))) == CTRUE) || 
                  (self.Selector.Id() == Core.C_externC.Id())) /* If:7 */{ 
                Result = EID{_Zt2.Id(),0}
                /* If!7 */}  else if (F_boolean_I_any(F_Compile_sort_equal_class(C_any,F_Compile_osort_any(_Zt1.Id()))) == CTRUE) /* If:7 */{ 
                /* Let:8 */{ 
                  var _CL_obj *ClaireUnion   = To_Union(new(ClaireUnion).Is(C_Union))
                  /* noccur = 5 */
                  _CL_obj.T1 = ToType(C_any.Id())
                  _CL_obj.T2 = _Zt2
                  Result = EID{_CL_obj.Id(),0}
                  /* Let-8 */} 
                } else {
                Result = EID{_Zt1.Id(),0}
                /* If-7 */} 
              /* Let-6 */} 
            /* Let-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* If-2 */} 
    return Result} 
  
// The EID go function for: use_range @ method (throw: true) 
func E_Optimize_use_range_method (self EID,_Zl EID) EID { 
    return /*(sm for use_range @ method= EID)*/ F_Optimize_use_range_method(ToMethod(OBJ(self)),ToList(OBJ(_Zl)) )} 
  
// ******************************************************************
// *    Part 2: Generic c_type & c_code                             *
// ******************************************************************
// this is the optimizer for messages
// It follows the stucture of the evaluator (self_eval)
// optimize is the distributed compiling method equivalent to the
// evaluation "behave" method
/* {1} OPT.The go function for: c_type(self:Call) [] */
func F_c_type_Call (self *Language.Call ) EID { 
    var Result EID 
    if (self.Selector.Id() == Language.C_function_I.Id()) /* If:2 */{ 
      Result = EID{C_function.Id(),0}
      } else {
      /* Let:3 */{ 
        var s *ClaireProperty   = self.Selector
        /* noccur = 13 */
        /* Let:4 */{ 
          var l *ClaireList   = self.Args
          /* noccur = 18 */
          /* Let:5 */{ 
            var _Ztype *ClaireList  
            /* noccur = 8 */
            var _Ztype_try04956 EID 
            /* Iteration:6 */{ 
              var v_list6 *ClaireList  
              var x *ClaireAny  
              var v_local6 *ClaireAny  
              v_list6 = l
              _Ztype_try04956 = EID{CreateList(ToType(CEMPTY.Id()),v_list6.Length()).Id(),0}
              for CLcount := 0; CLcount < v_list6.Length(); CLcount++{ 
                x = v_list6.At(CLcount)
                var v_local6_try04968 EID 
                v_local6_try04968 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
                /* ERROR PROTECTION INSERTED (v_local6-_Ztype_try04956) */
                if ErrorIn(v_local6_try04968) {_Ztype_try04956 = v_local6_try04968
                _Ztype_try04956 = v_local6_try04968
                break
                } else {
                v_local6 = ANY(v_local6_try04968)
                ToList(OBJ(_Ztype_try04956)).PutAt(CLcount,v_local6)
                } 
              }
              /* Iteration-6 */} 
            /* ERROR PROTECTION INSERTED (_Ztype-Result) */
            if ErrorIn(_Ztype_try04956) {Result = _Ztype_try04956
            } else {
            _Ztype = ToList(OBJ(_Ztype_try04956))
            if (s.Id() == C_safe.Id()) /* If:6 */{ 
              Result = _Ztype.At(1-1).ToEID()
              /* If!6 */}  else if ((s.Id() == Core.C_externC.Id()) && 
                ((l.Length() == 2) && 
                  (C_class.Id() == l.At(2-1).Isa.Id()))) /* If:6 */{ 
              Result = l.At(2-1).ToEID()
              /* If!6 */}  else if ((s.Id() == C_new.Id()) && 
                (C_class.Id() == l.At(1-1).Isa.Id())) /* If:6 */{ 
              Result = l.At(1-1).ToEID()
              /* If!6 */}  else if ((s.Id() == Core.C_check_in.Id()) && 
                (l.At(2-1).Isa.IsIn(C_type) == CTRUE)) /* If:6 */{ 
              if (l.Length() == 2) /* If:7 */{ 
                Result = EID{F_Optimize_sort_abstract_I_type(ToType(l.At(2-1))).Id(),0}
                } else {
                Result = l.At(2-1).ToEID()
                /* If-7 */} 
              /* If!6 */}  else if ((s.Id() == C_nth.Id()) && 
                (ToType(_Ztype.At(1-1)).Included(ToType(C_array.Id())) == CTRUE)) /* If:6 */{ 
              if (Core.F_member_type(ToType(_Ztype.At(1-1))).Included(ToType(C_float.Id())) == CTRUE) /* If:7 */{ 
                Result = EID{C_float.Id(),0}
                } else {
                Result = EID{Core.F_member_type(ToType(_Ztype.At(1-1))).Id(),0}
                /* If-7 */} 
              /* If!6 */}  else if ((s.Id() == Core.C__at.Id()) && 
                (l.At(1-1).Isa.IsIn(C_property) == CTRUE)) /* If:6 */{ 
              /* Let:7 */{ 
                var p *ClaireProperty   = ToProperty(l.At(1-1))
                /* noccur = 2 */
                /* Let:8 */{ 
                  var c *ClaireAny   = l.At(2-1)
                  /* noccur = 3 */
                  if ((C_class.Id() == c.Isa.Id()) && 
                      (C_method.Id() == Core.F_owner_any(ANY(Core.F_CALL(ToProperty(Core.C__at.Id()),ARGS(EID{p.Id(),0},c.ToEID())))).Id())) /* If:9 */{ 
                    Result = EID{MakeConstantSet(ANY(Core.F_CALL(ToProperty(Core.C__at.Id()),ARGS(EID{p.Id(),0},c.ToEID())))).Id(),0}
                    } else {
                    Result = EID{C_any.Id(),0}
                    /* If-9 */} 
                  /* Let-8 */} 
                /* Let-7 */} 
              /* If!6 */}  else if ((s.Id() == C_get.Id()) && 
                (l.At(1-1).Isa.IsIn(C_relation) == CTRUE)) /* If:6 */{ 
              /* Let:7 */{ 
                var r *ClaireRelation   = ToRelation(l.At(1-1))
                /* noccur = 4 */
                if (r.Isa.IsIn(C_property) == CTRUE) /* If:8 */{ 
                  /* Let:9 */{ 
                    var g0488 *ClaireProperty   = ToProperty(r.Id())
                    /* noccur = 2 */
                    /* Let:10 */{ 
                      var xs *ClaireObject   = Core.F__at_property1(g0488,ToTypeExpression(_Ztype.At(2-1)).Class_I())
                      /* noccur = 2 */
                      if (C_slot.Id() == xs.Isa.Id()) /* If:11 */{ 
                        /* Let:12 */{ 
                          var g0489 *ClaireSlot   = ToSlot(xs.Id())
                          /* noccur = 5 */
                          if ((g0489.Range.Included(ToType(C_set.Id())) == CTRUE) && 
                              (C_compiler.Safety < 3)) /* If:13 */{ 
                            Result = EID{ToTypeExpression(OBJ(Core.F_CALL(C_range,ARGS(l.At(2-1).ToEID())))).Class_I().Id(),0}
                            /* If!13 */}  else if (g0489.Range.Contains(g0489.Default) == CTRUE) /* If:13 */{ 
                            Result = EID{g0489.Range.Id(),0}
                            } else {
                            Result = EID{F_Optimize_extends_type(g0489.Range).Id(),0}
                            /* If-13 */} 
                          /* Let-12 */} 
                        } else {
                        Result = EID{g0488.Range.Id(),0}
                        /* If-11 */} 
                      /* Let-10 */} 
                    /* Let-9 */} 
                  /* If!8 */}  else if (C_table.Id() == r.Isa.Id()) /* If:8 */{ 
                  /* Let:9 */{ 
                    var g0491 *ClaireTable   = ToTable(r.Id())
                    /* noccur = 4 */
                    if (g0491.Range.Contains(g0491.Default) == CTRUE) /* If:10 */{ 
                      Result = EID{g0491.Range.Id(),0}
                      } else {
                      Result = EID{F_Optimize_extends_type(g0491.Range).Id(),0}
                      /* If-10 */} 
                    /* Let-9 */} 
                  } else {
                  Result = EID{CFALSE.Id(),0}
                  /* If-8 */} 
                /* Let-7 */} 
              } else {
              /* Let:7 */{ 
                var r *ClaireAny   = F_Optimize_restriction_I_property(s,_Ztype,CTRUE)
                /* noccur = 5 */
                if (C_slot.Id() == r.Isa.Id()) /* If:8 */{ 
                  /* Let:9 */{ 
                    var g0492 *ClaireSlot   = ToSlot(r)
                    /* noccur = 1 */
                    if ((s.Id() == C_instances.Id()) && 
                        (C_class.Id() == l.At(1-1).Isa.Id())) /* If:10 */{ 
                      /* Let:11 */{ 
                        var _CL_obj *ClaireParam   = To_Param(new(ClaireParam).Is(C_Param))
                        /* noccur = 7 */
                        _CL_obj.Arg = C_list
                        _CL_obj.Params = MakeConstantList(C_of.Id())
                        _CL_obj.Args = MakeConstantList(MakeConstantSet(l.At(1-1)).Id())
                        Result = EID{_CL_obj.Id(),0}
                        /* Let-11 */} 
                      } else {
                      Result = EID{g0492.Range.Id(),0}
                      /* If-10 */} 
                    /* Let-9 */} 
                  /* If!8 */}  else if (C_method.Id() == r.Isa.Id()) /* If:8 */{ 
                  /* Let:9 */{ 
                    var g0493 *ClaireMethod   = ToMethod(r)
                    /* noccur = 1 */
                    Result = F_Optimize_use_range_method(g0493,_Ztype)
                    /* Let-9 */} 
                  /* If!8 */}  else if (F_boolean_I_any(s.Restrictions.Id()).Id() != CTRUE.Id()) /* If:8 */{ 
                  Result = EID{F_Optimize_selector_psort_Call(self).Id(),0}
                  /* If!8 */}  else if ((s.Open == 3) || 
                    (r != C_Optimize_ambiguous.Id())) /* If:8 */{ 
                  Result = EID{F_Optimize_sort_abstract_I_type(s.Range).Id(),0}
                  } else {
                  Result = EID{F_Optimize_sort_abstract_I_type(ToType(F_Optimize_restriction_I_property(s,_Ztype,CFALSE))).Id(),0}
                  /* If-8 */} 
                /* Let-7 */} 
              /* If-6 */} 
            }
            /* Let-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* If-2 */} 
    return RangeCheck(ToType(C_type.Id()),Result)} 
  
// The EID go function for: c_type @ Call (throw: true) 
func E_c_type_Call (self EID) EID { 
    return /*(sm for c_type @ Call= EID)*/ F_c_type_Call(Language.To_Call(OBJ(self)) )} 
  
// this is the optimizer for messages : does not use the sort unless there is a macro
/* {1} OPT.The go function for: c_code(self:Call) [] */
func F_c_code_Call (self *Language.Call ) EID { 
    var Result EID 
    Result = F_Optimize_c_code_call_Call(self,C_void)
    return Result} 
  
// The EID go function for: c_code @ Call (throw: true) 
func E_c_code_Call (self EID) EID { 
    return /*(sm for c_code @ Call= EID)*/ F_c_code_Call(Language.To_Call(OBJ(self)) )} 
  
/* {1} OPT.The go function for: c_code_call(self:Call,sx:class) [] */
func F_Optimize_c_code_call_Call (self *Language.Call ,sx *ClaireClass ) EID { 
    var Result EID 
    
    /* Let:2 */{ 
      var s *ClaireProperty   = self.Selector
      /* noccur = 46 */
      /* Let:3 */{ 
        var l *ClaireList   = self.Args
        /* noccur = 52 */
        var g0505I *ClaireBoolean  
        var g0505I_try05064 EID 
        /* and:4 */{ 
          var v_and4 *ClaireBoolean  
          
          v_and4 = l.At(1-1).Isa.IsIn(Core.C_global_variable)
          if (v_and4 == CFALSE) {g0505I_try05064 = EID{CFALSE.Id(),0}
          } else /* arg:5 */{ 
            v_and4 = Equal(ANY(Core.F_CALL(C_range,ARGS(l.At(1-1).ToEID()))),CEMPTY.Id())
            if (v_and4 == CFALSE) {g0505I_try05064 = EID{CFALSE.Id(),0}
            } else /* arg:6 */{ 
              var v_and4_try05077 EID 
              v_and4_try05077 = F_Compile_designated_ask_any(ANY(Core.F_CALL(C_value,ARGS(l.At(1-1).ToEID()))))
              /* ERROR PROTECTION INSERTED (v_and4-g0505I_try05064) */
              if ErrorIn(v_and4_try05077) {g0505I_try05064 = v_and4_try05077
              } else {
              v_and4 = ToBoolean(OBJ(v_and4_try05077))
              if (v_and4 == CFALSE) {g0505I_try05064 = EID{CFALSE.Id(),0}
              } else /* arg:7 */{ 
                g0505I_try05064 = EID{CTRUE.Id(),0}/* arg-7 */} 
              /* arg-6 */} 
            /* arg-5 */} 
          }
          /* and-4 */} 
        /* ERROR PROTECTION INSERTED (g0505I-Result) */
        if ErrorIn(g0505I_try05064) {Result = g0505I_try05064
        } else {
        g0505I = ToBoolean(OBJ(g0505I_try05064))
        if (g0505I == CTRUE) /* If:4 */{ 
          Result = ToArray(l.Id()).NthPut(1,ANY(Core.F_CALL(C_value,ARGS(l.At(1-1).ToEID())))).ToEID()
          } else {
          Result = EID{CFALSE.Id(),0}
          /* If-4 */} 
        }
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        /* Let:4 */{ 
          var m *ClaireAny  
          /* noccur = 3 */
          var m_try05105 EID 
          m_try05105 = F_Optimize_inline_optimize_ask_Call(self)
          /* ERROR PROTECTION INSERTED (m-Result) */
          if ErrorIn(m_try05105) {Result = m_try05105
          } else {
          m = ANY(m_try05105)
          /* Let:5 */{ 
            var b *ClaireBoolean   = l.At(1-1).Isa.IsIn(C_property)
            /* noccur = 8 */
            /* Let:6 */{ 
              var d *ClaireAny  
              /* noccur = 2 */
              var d_try05147 EID 
              d_try05147 = F_Optimize_daccess_any(self.Id(),Core.F__sup_integer(C_compiler.Safety,5))
              /* ERROR PROTECTION INSERTED (d-Result) */
              if ErrorIn(d_try05147) {Result = d_try05147
              } else {
              d = ANY(d_try05147)
              if ((b == CTRUE) && 
                  (((s.Id() == Core.C_write.Id()) || 
                      (s.Id() == C_put.Id())) && 
                    (l.Length() == 3))) /* If:7 */{ 
                Result = F_Optimize_c_code_write_Call(self)
                /* If!7 */}  else if ((b == CTRUE) && 
                  ((s.Id() == Core.C_put_store.Id()) && 
                    ((l.Length() == 4) && 
                      (l.At(4-1) == CTRUE.Id())))) /* If:7 */{ 
                Result = F_Optimize_c_code_write_Call(self)
                /* If!7 */}  else if ((b == CTRUE) && 
                  (s.Id() == Core.C_unknown_ask.Id())) /* If:7 */{ 
                Result = F_Optimize_c_code_hold_property(ToProperty(l.At(1-1)),l.At(2-1),CNULL,CTRUE)
                /* If!7 */}  else if ((b == CTRUE) && 
                  (s.Id() == Core.C_known_ask.Id())) /* If:7 */{ 
                Result = F_Optimize_c_code_hold_property(ToProperty(l.At(1-1)),l.At(2-1),CNULL,CFALSE)
                /* If!7 */}  else if ((b == CTRUE) && 
                  ((s.Id() == Core.C_erase.Id()) && 
                    (l.At(2-1).Isa.IsIn(C_Variable) == CTRUE))) /* If:7 */{ 
                /* Let:8 */{ 
                  var g0519UU *ClaireAny  
                  /* noccur = 1 */
                  var g0519UU_try05209 EID 
                  g0519UU_try05209 = F_Optimize_Produce_erase_property(ToProperty(l.At(1-1)),To_Variable(l.At(2-1)))
                  /* ERROR PROTECTION INSERTED (g0519UU-Result) */
                  if ErrorIn(g0519UU_try05209) {Result = g0519UU_try05209
                  } else {
                  g0519UU = ANY(g0519UU_try05209)
                  Result = Core.F_CALL(C_c_code,ARGS(g0519UU.ToEID(),EID{sx.Id(),0}))
                  }
                  /* Let-8 */} 
                /* If!7 */}  else if (s.Id() == C_safe.Id()) /* If:7 */{ 
                /* Let:8 */{ 
                  var y int  = C_compiler.Safety
                  /* noccur = 1 */
                  /* Let:9 */{ 
                    var b *ClaireBoolean   = C_compiler.Overflow_ask
                    /* noccur = 1 */
                    /* Let:10 */{ 
                      var x *ClaireAny   = CNULL
                      /* noccur = 2 */
                      C_compiler.Safety = 1
                      C_compiler.Overflow_ask = CTRUE
                      var x_try052111 EID 
                      /* Let:11 */{ 
                        var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                        /* noccur = 5 */
                        _CL_obj.Selector = C_safe
                        /* update:12 */{ 
                          var va_arg1 *Language.Call  
                          var va_arg2 *ClaireList  
                          va_arg1 = _CL_obj
                          var va_arg2_try052213 EID 
                          /* Construct:13 */{ 
                            var v_bag_arg *ClaireAny  
                            va_arg2_try052213= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                            var v_bag_arg_try052314 EID 
                            v_bag_arg_try052314 = Core.F_CALL(C_c_code,ARGS(l.At(1-1).ToEID(),EID{sx.Id(),0}))
                            /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try052213) */
                            if ErrorIn(v_bag_arg_try052314) {va_arg2_try052213 = v_bag_arg_try052314
                            } else {
                            v_bag_arg = ANY(v_bag_arg_try052314)
                            ToList(OBJ(va_arg2_try052213)).AddFast(v_bag_arg)}
                            /* Construct-13 */} 
                          /* ERROR PROTECTION INSERTED (va_arg2-x_try052111) */
                          if ErrorIn(va_arg2_try052213) {x_try052111 = va_arg2_try052213
                          } else {
                          va_arg2 = ToList(OBJ(va_arg2_try052213))
                          /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                          va_arg1.Args = va_arg2
                          x_try052111 = EID{va_arg2.Id(),0}
                          }
                          /* update-12 */} 
                        /* ERROR PROTECTION INSERTED (x_try052111-x_try052111) */
                        if !ErrorIn(x_try052111) {
                        x_try052111 = EID{_CL_obj.Id(),0}
                        }
                        /* Let-11 */} 
                      /* ERROR PROTECTION INSERTED (x-Result) */
                      if ErrorIn(x_try052111) {Result = x_try052111
                      } else {
                      x = ANY(x_try052111)
                      Result = x.ToEID()
                      C_compiler.Safety = y
                      C_compiler.Overflow_ask = b
                      Result = x.ToEID()
                      }
                      /* Let-10 */} 
                    /* Let-9 */} 
                  /* Let-8 */} 
                /* If!7 */}  else if (((s.Id() == C_add.Id()) || 
                    (s.Id() == C_add_I.Id())) && 
                  (b == CTRUE)) /* If:7 */{ 
                Result = F_Optimize_c_code_add_Call(self)
                } else {
                var g0525I *ClaireBoolean  
                var g0525I_try05268 EID 
                /* and:8 */{ 
                  var v_and8 *ClaireBoolean  
                  
                  v_and8 = MakeBoolean((s.Id() == C_add.Id()) || (s.Id() == C_add_I.Id()))
                  if (v_and8 == CFALSE) {g0525I_try05268 = EID{CFALSE.Id(),0}
                  } else /* arg:9 */{ 
                    var v_and8_try052710 EID 
                    /* Let:10 */{ 
                      var g0528UU *ClaireType  
                      /* noccur = 1 */
                      var g0528UU_try052911 EID 
                      g0528UU_try052911 = Core.F_CALL(C_c_type,ARGS(l.At(1-1).ToEID()))
                      /* ERROR PROTECTION INSERTED (g0528UU-v_and8_try052710) */
                      if ErrorIn(g0528UU_try052911) {v_and8_try052710 = g0528UU_try052911
                      } else {
                      g0528UU = ToType(OBJ(g0528UU_try052911))
                      v_and8_try052710 = EID{g0528UU.Included(ToType(C_bag.Id())).Id(),0}
                      }
                      /* Let-10 */} 
                    /* ERROR PROTECTION INSERTED (v_and8-g0525I_try05268) */
                    if ErrorIn(v_and8_try052710) {g0525I_try05268 = v_and8_try052710
                    } else {
                    v_and8 = ToBoolean(OBJ(v_and8_try052710))
                    if (v_and8 == CFALSE) {g0525I_try05268 = EID{CFALSE.Id(),0}
                    } else /* arg:10 */{ 
                      g0525I_try05268 = EID{CTRUE.Id(),0}/* arg-10 */} 
                    /* arg-9 */} 
                  }
                  /* and-8 */} 
                /* ERROR PROTECTION INSERTED (g0525I-Result) */
                if ErrorIn(g0525I_try05268) {Result = g0525I_try05268
                } else {
                g0525I = ToBoolean(OBJ(g0525I_try05268))
                if (g0525I == CTRUE) /* If:8 */{ 
                  Result = F_Optimize_c_code_add_bag_Call(self)
                  /* If!8 */}  else if ((b == CTRUE) && 
                    (s.Id() == C_delete.Id())) /* If:8 */{ 
                  Result = F_Optimize_c_code_delete_Call(self)
                  /* If!8 */}  else if ((C_OPT.ToRemove.Contain_ask(s.Id()) == CTRUE) || 
                    ((s.Id() == C_c_interface.Id()) && 
                        ((l.Length() == 2) && 
                          (C_OPT.LegalModules.Contain_ask(F_Compile_get_module_object(ToObject(s.Id()))) != CTRUE)))) /* If:8 */{ 
                  Result = EID{CNIL.Id(),0}
                  /* If!8 */}  else if (d != CNULL) /* If:8 */{ 
                  Result = d.ToEID()
                  /* If!8 */}  else if (C_method.Id() == m.Isa.Id()) /* If:8 */{ 
                  /* Let:9 */{ 
                    var g0533UU *ClaireClass  
                    /* noccur = 1 */
                    var g0533UU_try053410 EID 
                    g0533UU_try053410 = F_Optimize_c_srange_method(ToMethod(m))
                    /* ERROR PROTECTION INSERTED (g0533UU-Result) */
                    if ErrorIn(g0533UU_try053410) {Result = g0533UU_try053410
                    } else {
                    g0533UU = ToClass(OBJ(g0533UU_try053410))
                    Result = F_Optimize_c_inline_method1(ToMethod(m),l,g0533UU)
                    }
                    /* Let-9 */} 
                  } else {
                  var g0535I *ClaireBoolean  
                  var g0535I_try05369 EID 
                  /* and:9 */{ 
                    var v_and9 *ClaireBoolean  
                    
                    v_and9 = MakeBoolean((s.Id() == C__equal.Id()) || (s.Id() == Core.C__I_equal.Id()))
                    if (v_and9 == CFALSE) {g0535I_try05369 = EID{CFALSE.Id(),0}
                    } else /* arg:10 */{ 
                      var v_and9_try053711 EID 
                      /* Let:11 */{ 
                        var g0538UU *ClaireAny  
                        /* noccur = 1 */
                        var g0538UU_try053912 EID 
                        g0538UU_try053912 = F_Optimize_daccess_any(l.At(1-1),CTRUE)
                        /* ERROR PROTECTION INSERTED (g0538UU-v_and9_try053711) */
                        if ErrorIn(g0538UU_try053912) {v_and9_try053711 = g0538UU_try053912
                        } else {
                        g0538UU = ANY(g0538UU_try053912)
                        v_and9_try053711 = EID{Core.F_known_ask_any(g0538UU).Id(),0}
                        }
                        /* Let-11 */} 
                      /* ERROR PROTECTION INSERTED (v_and9-g0535I_try05369) */
                      if ErrorIn(v_and9_try053711) {g0535I_try05369 = v_and9_try053711
                      } else {
                      v_and9 = ToBoolean(OBJ(v_and9_try053711))
                      if (v_and9 == CFALSE) {g0535I_try05369 = EID{CFALSE.Id(),0}
                      } else /* arg:11 */{ 
                        g0535I_try05369 = EID{CTRUE.Id(),0}/* arg-11 */} 
                      /* arg-10 */} 
                    }
                    /* and-9 */} 
                  /* ERROR PROTECTION INSERTED (g0535I-Result) */
                  if ErrorIn(g0535I_try05369) {Result = g0535I_try05369
                  } else {
                  g0535I = ToBoolean(OBJ(g0535I_try05369))
                  if (g0535I == CTRUE) /* If:9 */{ 
                    Result = F_Optimize_c_code_hold_property(ToProperty(ToList(OBJ(Core.F_CALL(C_args,ARGS(l.At(1-1).ToEID())))).At(1-1)),ToList(OBJ(Core.F_CALL(C_args,ARGS(l.At(1-1).ToEID())))).At(2-1),l.At(2-1),Equal(s.Id(),C__equal.Id()))
                    } else {
                    var g0540I *ClaireBoolean  
                    var g0540I_try054110 EID 
                    /* and:10 */{ 
                      var v_and10 *ClaireBoolean  
                      
                      v_and10 = MakeBoolean((s.Id() == C__equal.Id()) || (s.Id() == Core.C__I_equal.Id()))
                      if (v_and10 == CFALSE) {g0540I_try054110 = EID{CFALSE.Id(),0}
                      } else /* arg:11 */{ 
                        var v_and10_try054212 EID 
                        /* Let:12 */{ 
                          var g0543UU *ClaireAny  
                          /* noccur = 1 */
                          var g0543UU_try054413 EID 
                          g0543UU_try054413 = F_Optimize_daccess_any(l.At(2-1),CTRUE)
                          /* ERROR PROTECTION INSERTED (g0543UU-v_and10_try054212) */
                          if ErrorIn(g0543UU_try054413) {v_and10_try054212 = g0543UU_try054413
                          } else {
                          g0543UU = ANY(g0543UU_try054413)
                          v_and10_try054212 = EID{Core.F_known_ask_any(g0543UU).Id(),0}
                          }
                          /* Let-12 */} 
                        /* ERROR PROTECTION INSERTED (v_and10-g0540I_try054110) */
                        if ErrorIn(v_and10_try054212) {g0540I_try054110 = v_and10_try054212
                        } else {
                        v_and10 = ToBoolean(OBJ(v_and10_try054212))
                        if (v_and10 == CFALSE) {g0540I_try054110 = EID{CFALSE.Id(),0}
                        } else /* arg:12 */{ 
                          g0540I_try054110 = EID{CTRUE.Id(),0}/* arg-12 */} 
                        /* arg-11 */} 
                      }
                      /* and-10 */} 
                    /* ERROR PROTECTION INSERTED (g0540I-Result) */
                    if ErrorIn(g0540I_try054110) {Result = g0540I_try054110
                    } else {
                    g0540I = ToBoolean(OBJ(g0540I_try054110))
                    if (g0540I == CTRUE) /* If:10 */{ 
                      Result = F_Optimize_c_code_hold_property(ToProperty(ToList(OBJ(Core.F_CALL(C_args,ARGS(l.At(2-1).ToEID())))).At(1-1)),ToList(OBJ(Core.F_CALL(C_args,ARGS(l.At(2-1).ToEID())))).At(2-1),l.At(1-1),Equal(s.Id(),C__equal.Id()))
                      /* If!10 */}  else if (((s.Id() == C_put.Id()) || 
                          (s.Id() == C_nth_equal.Id())) && 
                        ((C_table.Id() == l.At(1-1).Isa.Id()) && 
                          (l.Length() == 3))) /* If:10 */{ 
                      Result = F_Optimize_c_code_table_Call(self)
                      } else {
                      var g0545I *ClaireBoolean  
                      var g0545I_try054611 EID 
                      /* and:11 */{ 
                        var v_and11 *ClaireBoolean  
                        
                        v_and11 = MakeBoolean((s.Id() == C_nth_put.Id()) || (s.Id() == C_nth_equal.Id()))
                        if (v_and11 == CFALSE) {g0545I_try054611 = EID{CFALSE.Id(),0}
                        } else /* arg:12 */{ 
                          var v_and11_try054713 EID 
                          /* Let:13 */{ 
                            var g0548UU *ClaireType  
                            /* noccur = 1 */
                            var g0548UU_try054914 EID 
                            g0548UU_try054914 = Core.F_CALL(C_c_type,ARGS(l.At(1-1).ToEID()))
                            /* ERROR PROTECTION INSERTED (g0548UU-v_and11_try054713) */
                            if ErrorIn(g0548UU_try054914) {v_and11_try054713 = g0548UU_try054914
                            } else {
                            g0548UU = ToType(OBJ(g0548UU_try054914))
                            v_and11_try054713 = EID{g0548UU.Included(ToType(C_array.Id())).Id(),0}
                            }
                            /* Let-13 */} 
                          /* ERROR PROTECTION INSERTED (v_and11-g0545I_try054611) */
                          if ErrorIn(v_and11_try054713) {g0545I_try054611 = v_and11_try054713
                          } else {
                          v_and11 = ToBoolean(OBJ(v_and11_try054713))
                          if (v_and11 == CFALSE) {g0545I_try054611 = EID{CFALSE.Id(),0}
                          } else /* arg:13 */{ 
                            v_and11 = Equal(MakeInteger(l.Length()).Id(),MakeInteger(3).Id())
                            if (v_and11 == CFALSE) {g0545I_try054611 = EID{CFALSE.Id(),0}
                            } else /* arg:14 */{ 
                              g0545I_try054611 = EID{CTRUE.Id(),0}/* arg-14 */} 
                            /* arg-13 */} 
                          /* arg-12 */} 
                        }
                        /* and-11 */} 
                      /* ERROR PROTECTION INSERTED (g0545I-Result) */
                      if ErrorIn(g0545I_try054611) {Result = g0545I_try054611
                      } else {
                      g0545I = ToBoolean(OBJ(g0545I_try054611))
                      if (g0545I == CTRUE) /* If:11 */{ 
                        Result = F_Optimize_c_code_array_Call(self)
                        /* If!11 */}  else if ((s.Id() == C_nth.Id()) || 
                          (((s.Id() == C_get.Id()) && 
                              (C_table.Id() == l.At(1-1).Isa.Id())) || 
                            ((s.Id() == C_nth_get.Id()) && 
                                (l.At(1-1).Isa.IsIn(C_array) == CTRUE)))) /* If:11 */{ 
                        Result = F_Optimize_c_code_nth_Call(self)
                        /* If!11 */}  else if (s.Id() == C__Z.Id()) /* If:11 */{ 
                        Result = F_Optimize_c_code_belong_Call(self)
                        /* If!11 */}  else if (s.Id() == Core.C_Id.Id()) /* If:11 */{ 
                        /* Let:12 */{ 
                          var g0552UU *ClaireAny  
                          /* noccur = 1 */
                          var g0552UU_try055313 EID 
                          g0552UU_try055313 = EVAL(l.At(1-1))
                          /* ERROR PROTECTION INSERTED (g0552UU-Result) */
                          if ErrorIn(g0552UU_try055313) {Result = g0552UU_try055313
                          } else {
                          g0552UU = ANY(g0552UU_try055313)
                          Result = Core.F_CALL(C_c_code,ARGS(g0552UU.ToEID()))
                          }
                          /* Let-12 */} 
                        /* If!11 */}  else if (s.Id() == Language.C_function_I.Id()) /* If:11 */{ 
                        /* Let:12 */{ 
                          var g0554UU *ClaireString  
                          /* noccur = 1 */
                          var g0554UU_try055513 EID 
                          /* Let:13 */{ 
                            var g0556UU *ClaireSymbol  
                            /* noccur = 1 */
                            var g0556UU_try055714 EID 
                            g0556UU_try055714 = Language.F_extract_symbol_any(l.At(1-1))
                            /* ERROR PROTECTION INSERTED (g0556UU-g0554UU_try055513) */
                            if ErrorIn(g0556UU_try055714) {g0554UU_try055513 = g0556UU_try055714
                            } else {
                            g0556UU = ToSymbol(OBJ(g0556UU_try055714))
                            g0554UU_try055513 = EID{g0556UU.String_I().Id(),0}
                            }
                            /* Let-13 */} 
                          /* ERROR PROTECTION INSERTED (g0554UU-Result) */
                          if ErrorIn(g0554UU_try055513) {Result = g0554UU_try055513
                          } else {
                          g0554UU = ToString(OBJ(g0554UU_try055513))
                          Result = F_make_function_string(g0554UU).ToEID()
                          }
                          /* Let-12 */} 
                        /* If!11 */}  else if ((s.Id() == Core.C_not.Id()) && 
                          (l.At(1-1).Isa.IsIn(Language.C_Select) == CTRUE)) /* If:11 */{ 
                        Result = F_Optimize_c_code_not_Select(Language.To_Select(l.At(1-1)))
                        /* If!11 */}  else if ((s.Id() == Core.C_call.Id()) && 
                          (l.At(1-1).Isa.IsIn(C_property) == CTRUE)) /* If:11 */{ 
                        /* Let:12 */{ 
                          var g0558UU *Language.Call  
                          /* noccur = 1 */
                          var g0558UU_try055913 EID 
                          /* Let:13 */{ 
                            var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                            /* noccur = 5 */
                            _CL_obj.Selector = ToProperty(l.At(1-1))
                            /* update:14 */{ 
                              var va_arg1 *Language.Call  
                              var va_arg2 *ClaireList  
                              va_arg1 = _CL_obj
                              var va_arg2_try056015 EID 
                              va_arg2_try056015 = l.Cdr()
                              /* ERROR PROTECTION INSERTED (va_arg2-g0558UU_try055913) */
                              if ErrorIn(va_arg2_try056015) {g0558UU_try055913 = va_arg2_try056015
                              } else {
                              va_arg2 = ToList(OBJ(va_arg2_try056015))
                              /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                              va_arg1.Args = va_arg2
                              g0558UU_try055913 = EID{va_arg2.Id(),0}
                              }
                              /* update-14 */} 
                            /* ERROR PROTECTION INSERTED (g0558UU_try055913-g0558UU_try055913) */
                            if !ErrorIn(g0558UU_try055913) {
                            g0558UU_try055913 = EID{_CL_obj.Id(),0}
                            }
                            /* Let-13 */} 
                          /* ERROR PROTECTION INSERTED (g0558UU-Result) */
                          if ErrorIn(g0558UU_try055913) {Result = g0558UU_try055913
                          } else {
                          g0558UU = Language.To_Call(OBJ(g0558UU_try055913))
                          Result = Core.F_CALL(C_c_code,ARGS(EID{g0558UU.Id(),0}))
                          }
                          /* Let-12 */} 
                        /* If!11 */}  else if (s.Open == 3) /* If:11 */{ 
                        /* Let:12 */{ 
                          var g0561UU *ClaireList  
                          /* noccur = 1 */
                          var g0561UU_try056213 EID 
                          /* Iteration:13 */{ 
                            var v_list13 *ClaireList  
                            var x *ClaireAny  
                            var v_local13 *ClaireAny  
                            v_list13 = l
                            g0561UU_try056213 = EID{CreateList(ToType(CEMPTY.Id()),v_list13.Length()).Id(),0}
                            for CLcount := 0; CLcount < v_list13.Length(); CLcount++{ 
                              x = v_list13.At(CLcount)
                              var v_local13_try056315 EID 
                              v_local13_try056315 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
                              /* ERROR PROTECTION INSERTED (v_local13-g0561UU_try056213) */
                              if ErrorIn(v_local13_try056315) {g0561UU_try056213 = v_local13_try056315
                              g0561UU_try056213 = v_local13_try056315
                              break
                              } else {
                              v_local13 = ANY(v_local13_try056315)
                              ToList(OBJ(g0561UU_try056213)).PutAt(CLcount,v_local13)
                              } 
                            }
                            /* Iteration-13 */} 
                          /* ERROR PROTECTION INSERTED (g0561UU-Result) */
                          if ErrorIn(g0561UU_try056213) {Result = g0561UU_try056213
                          } else {
                          g0561UU = ToList(OBJ(g0561UU_try056213))
                          Result = F_Optimize_c_warn_property(s,l,g0561UU)
                          }
                          /* Let-12 */} 
                        } else {
                        var g0564I *ClaireBoolean  
                        /* and:12 */{ 
                          var v_and12 *ClaireBoolean  
                          
                          v_and12 = Equal(s.Id(),Language.C_bit_vector.Id())
                          if (v_and12 == CFALSE) {g0564I = CFALSE
                          } else /* arg:13 */{ 
                            /* Let:14 */{ 
                              var g0565UU *ClaireAny  
                              /* noccur = 1 */
                              /* For:15 */{ 
                                var y *ClaireAny  
                                _ = y
                                g0565UU= CFALSE.Id()
                                var y_support *ClaireList  
                                y_support = self.Args
                                y_len := y_support.Length()
                                for i_it := 0; i_it < y_len; i_it++ { 
                                  y = y_support.At(i_it)
                                  if (C_integer.Id() != y.Isa.Id()) /* If:17 */{ 
                                     /*v = g0565UU, s =any*/
g0565UU = CTRUE.Id()
                                    break
                                    /* If-17 */} 
                                  /* loop-16 */} 
                                /* For-15 */} 
                              v_and12 = Core.F_not_any(g0565UU)
                              /* Let-14 */} 
                            if (v_and12 == CFALSE) {g0564I = CFALSE
                            } else /* arg:14 */{ 
                              g0564I = CTRUE/* arg-14 */} 
                            /* arg-13 */} 
                          /* and-12 */} 
                        if (g0564I == CTRUE) /* If:12 */{ 
                          Result = EVAL(self.Id())
                          /* If!12 */}  else if ((s.Id() == C_Compile_anyObject_I.Id()) || 
                            ((s.Id() == C_Compile_object_I.Id()) || 
                              ((s.Id() == C_add_method.Id()) && 
                                  (b == CTRUE)))) /* If:12 */{ 
                          Result = EID{self.Id(),0}
                          } else {
                          /* Let:13 */{ 
                            var _Ztype *ClaireList  
                            /* noccur = 7 */
                            var _Ztype_try056614 EID 
                            /* Iteration:14 */{ 
                              var v_list14 *ClaireList  
                              var x *ClaireAny  
                              var v_local14 *ClaireAny  
                              v_list14 = l
                              _Ztype_try056614 = EID{CreateList(ToType(CEMPTY.Id()),v_list14.Length()).Id(),0}
                              for CLcount := 0; CLcount < v_list14.Length(); CLcount++{ 
                                x = v_list14.At(CLcount)
                                var v_local14_try056716 EID 
                                v_local14_try056716 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
                                /* ERROR PROTECTION INSERTED (v_local14-_Ztype_try056614) */
                                if ErrorIn(v_local14_try056716) {_Ztype_try056614 = v_local14_try056716
                                _Ztype_try056614 = v_local14_try056716
                                break
                                } else {
                                v_local14 = ANY(v_local14_try056716)
                                ToList(OBJ(_Ztype_try056614)).PutAt(CLcount,v_local14)
                                } 
                              }
                              /* Iteration-14 */} 
                            /* ERROR PROTECTION INSERTED (_Ztype-Result) */
                            if ErrorIn(_Ztype_try056614) {Result = _Ztype_try056614
                            } else {
                            _Ztype = ToList(OBJ(_Ztype_try056614))
                            /* Let:14 */{ 
                              var z *ClaireAny   = F_Optimize_restriction_I_property(s,_Ztype,CTRUE)
                              /* noccur = 5 */
                              if (C_slot.Id() == z.Isa.Id()) /* If:15 */{ 
                                /* Let:16 */{ 
                                  var g0501 *ClaireSlot   = ToSlot(z)
                                  /* noccur = 4 */
                                  /* Let:17 */{ 
                                    var _Zunknown *ClaireBoolean   = MakeBoolean((g0501.Range.Contains(g0501.Default) != CTRUE) && (C_OPT.Knowns.Contain_ask(s.Id()) != CTRUE) && (C_compiler.Safety < 5))
                                    /* noccur = 2 */
                                    var g0568I *ClaireBoolean  
                                    var g0568I_try056918 EID 
                                    /* or:18 */{ 
                                      var v_or18 *ClaireBoolean  
                                      
                                      v_or18 = _Zunknown.Not
                                      if (v_or18 == CTRUE) {g0568I_try056918 = EID{CTRUE.Id(),0}
                                      } else /* or:19 */{ 
                                        var v_or18_try057020 EID 
                                        v_or18_try057020 = F_Compile_designated_ask_any(l.At(1-1))
                                        /* ERROR PROTECTION INSERTED (v_or18-g0568I_try056918) */
                                        if ErrorIn(v_or18_try057020) {g0568I_try056918 = v_or18_try057020
                                        } else {
                                        v_or18 = ToBoolean(OBJ(v_or18_try057020))
                                        if (v_or18 == CTRUE) {g0568I_try056918 = EID{CTRUE.Id(),0}
                                        } else /* or:20 */{ 
                                          g0568I_try056918 = EID{CFALSE.Id(),0}/* org-20 */} 
                                        /* org-19 */} 
                                      }
                                      /* or-18 */} 
                                    /* ERROR PROTECTION INSERTED (g0568I-Result) */
                                    if ErrorIn(g0568I_try056918) {Result = g0568I_try056918
                                    } else {
                                    g0568I = ToBoolean(OBJ(g0568I_try056918))
                                    if (g0568I == CTRUE) /* If:18 */{ 
                                      /* Let:19 */{ 
                                        var _CL_obj *Language.CallSlot   = Language.To_CallSlot(new(Language.CallSlot).Is(Language.C_Call_slot))
                                        /* noccur = 7 */
                                        _CL_obj.Selector = g0501
                                        /* update:20 */{ 
                                          var va_arg1 *Language.CallSlot  
                                          var va_arg2 *ClaireAny  
                                          va_arg1 = _CL_obj
                                          var va_arg2_try057121 EID 
                                          va_arg2_try057121 = Core.F_CALL(C_c_code,ARGS(l.At(1-1).ToEID(),EID{F_Compile_psort_any(Core.F_domain_I_restriction(ToRestriction(g0501.Id())).Id()).Id(),0}))
                                          /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                                          if ErrorIn(va_arg2_try057121) {Result = va_arg2_try057121
                                          } else {
                                          va_arg2 = ANY(va_arg2_try057121)
                                          /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                                          va_arg1.Arg = va_arg2
                                          Result = va_arg2.ToEID()
                                          }
                                          /* update-20 */} 
                                        /* ERROR PROTECTION INSERTED (Result-Result) */
                                        if !ErrorIn(Result) {
                                        _CL_obj.Test = _Zunknown
                                        Result = EID{_CL_obj.Id(),0}
                                        }
                                        /* Let-19 */} 
                                      } else {
                                      Core.F_tformat_string(MakeString("---- note: ... unsafe access to unknown slot: ~S \n"),3,MakeConstantList(self.Id()))
                                      if (C_compiler.Optimize_ask == CTRUE) /* If:19 */{ 
                                        F_Compile_notice_void()
                                        Core.F_tformat_string(MakeString("poorly optimized slot access: ~S\n"),3,MakeConstantList(self.Id()))
                                        /* If-19 */} 
                                      Result = F_Optimize_c_warn_property(s,l,_Ztype)
                                      /* If-18 */} 
                                    }
                                    /* Let-17 */} 
                                  /* Let-16 */} 
                                /* If!15 */}  else if (C_method.Id() == z.Isa.Id()) /* If:15 */{ 
                                /* Let:16 */{ 
                                  var g0502 *ClaireMethod   = ToMethod(z)
                                  /* noccur = 3 */
                                  
                                  if (_Ztype.Memq(C_void.Id()) == CTRUE) /* If:17 */{ 
                                    Result = Core.F_CALL(C_Compile_Cerror,ARGS(EID{MakeString("[205] call ~S uses a void argument [~S]").Id(),0},EID{self.Id(),0},EID{_Ztype.Id(),0}))
                                    } else {
                                    Result = EID{CFALSE.Id(),0}
                                    /* If-17 */} 
                                  /* ERROR PROTECTION INSERTED (Result-Result) */
                                  if !ErrorIn(Result) {
                                  if (((s.Id() == C_begin.Id()) || 
                                        (s.Id() == C_end.Id())) && 
                                      (l.At(1-1).Isa.IsIn(C_module) == CTRUE)) /* If:17 */{ 
                                    Result = EVAL(self.Id())
                                    } else {
                                    Result = EID{CFALSE.Id(),0}
                                    /* If-17 */} 
                                  /* ERROR PROTECTION INSERTED (Result-Result) */
                                  if !ErrorIn(Result) {
                                  var g0572I *ClaireBoolean  
                                  var g0572I_try057317 EID 
                                  /* or:17 */{ 
                                    var v_or17 *ClaireBoolean  
                                    
                                    var v_or17_try057418 EID 
                                    /* Let:18 */{ 
                                      var g0575UU *ClaireAny  
                                      /* noccur = 1 */
                                      var g0575UU_try057619 EID 
                                      g0575UU_try057619 = Core.F_last_list(g0502.Domain)
                                      /* ERROR PROTECTION INSERTED (g0575UU-v_or17_try057418) */
                                      if ErrorIn(g0575UU_try057619) {v_or17_try057418 = g0575UU_try057619
                                      } else {
                                      g0575UU = ANY(g0575UU_try057619)
                                      v_or17_try057418 = EID{Equal(g0575UU,C_listargs.Id()).Id(),0}
                                      }
                                      /* Let-18 */} 
                                    /* ERROR PROTECTION INSERTED (v_or17-g0572I_try057317) */
                                    if ErrorIn(v_or17_try057418) {g0572I_try057317 = v_or17_try057418
                                    } else {
                                    v_or17 = ToBoolean(OBJ(v_or17_try057418))
                                    if (v_or17 == CTRUE) {g0572I_try057317 = EID{CTRUE.Id(),0}
                                    } else /* or:18 */{ 
                                      v_or17 = MakeBoolean((g0502.Domain.ValuesO()[1-1] == C_void.Id()) && (l.At(1-1) != ClEnv.Id()))
                                      if (v_or17 == CTRUE) {g0572I_try057317 = EID{CTRUE.Id(),0}
                                      } else /* or:19 */{ 
                                        g0572I_try057317 = EID{CFALSE.Id(),0}/* org-19 */} 
                                      /* org-18 */} 
                                    }
                                    /* or-17 */} 
                                  /* ERROR PROTECTION INSERTED (g0572I-Result) */
                                  if ErrorIn(g0572I_try057317) {Result = g0572I_try057317
                                  } else {
                                  g0572I = ToBoolean(OBJ(g0572I_try057317))
                                  if (g0572I == CTRUE) /* If:17 */{ 
                                    Result = F_Optimize_open_message_property(s,l)
                                    } else {
                                    Result = F_Optimize_c_code_method_method2(g0502,l,_Ztype,sx)
                                    /* If-17 */} 
                                  }
                                  }}
                                  /* Let-16 */} 
                                /* If!15 */}  else if (z.Isa.IsIn(C_keyword) == CTRUE) /* If:15 */{ 
                                Result = F_Optimize_c_warn_property(s,l,_Ztype)
                                } else {
                                Result = F_Optimize_c_warn_Call(self,_Ztype.Id())
                                /* If-15 */} 
                              /* Let-14 */} 
                            }
                            /* Let-13 */} 
                          /* If-12 */} 
                        /* If-11 */} 
                      }
                      /* If-10 */} 
                    }
                    /* If-9 */} 
                  }
                  /* If-8 */} 
                }
                /* If-7 */} 
              }
              /* Let-6 */} 
            /* Let-5 */} 
          }
          /* Let-4 */} 
        }
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_code_call @ Call (throw: true) 
func E_Optimize_c_code_call_Call (self EID,sx EID) EID { 
    return /*(sm for c_code_call @ Call= EID)*/ F_Optimize_c_code_call_Call(Language.To_Call(OBJ(self)),ToClass(OBJ(sx)) )} 
  
// true error
// create the compiled message with necessary protections
/* {1} OPT.The go function for: open_message(self:property,l:list) [] */
func F_Optimize_open_message_property (self *ClaireProperty ,l *ClaireList ) EID { 
    var Result EID 
    F_Optimize_selector_register_property(self)
    /* Let:2 */{ 
      var _Zarg *ClaireList  
      /* noccur = 1 */
      var _Zarg_try05813 EID 
      /* Iteration:3 */{ 
        var v_list3 *ClaireList  
        var x *ClaireAny  
        var v_local3 *ClaireAny  
        v_list3 = l
        _Zarg_try05813 = EID{CreateList(ToType(CEMPTY.Id()),v_list3.Length()).Id(),0}
        for CLcount := 0; CLcount < v_list3.Length(); CLcount++{ 
          x = v_list3.At(CLcount)
          var v_local3_try05825 EID 
          var g0583I *ClaireBoolean  
          var g0583I_try05845 EID 
          /* Let:5 */{ 
            var g0585UU *ClaireType  
            /* noccur = 1 */
            var g0585UU_try05866 EID 
            g0585UU_try05866 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
            /* ERROR PROTECTION INSERTED (g0585UU-g0583I_try05845) */
            if ErrorIn(g0585UU_try05866) {g0583I_try05845 = g0585UU_try05866
            } else {
            g0585UU = ToType(OBJ(g0585UU_try05866))
            g0583I_try05845 = EID{Core.F__I_equal_any(g0585UU.Id(),C_void.Id()).Id(),0}
            }
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (g0583I-v_local3_try05825) */
          if ErrorIn(g0583I_try05845) {v_local3_try05825 = g0583I_try05845
          } else {
          g0583I = ToBoolean(OBJ(g0583I_try05845))
          if (g0583I == CTRUE) /* If:5 */{ 
            v_local3_try05825 = Core.F_CALL(C_c_code,ARGS(x.ToEID(),EID{C_any.Id(),0}))
            } else {
            v_local3_try05825 = Core.F_CALL(C_Compile_Cerror,ARGS(EID{MakeString("[206] use of void ~S in ~S~S").Id(),0},
              x.ToEID(),
              EID{self.Id(),0},
              EID{l.Id(),0}))
            /* If-5 */} 
          }
          /* ERROR PROTECTION INSERTED (v_local3-_Zarg_try05813) */
          if ErrorIn(v_local3_try05825) {_Zarg_try05813 = v_local3_try05825
          _Zarg_try05813 = v_local3_try05825
          break
          } else {
          v_local3 = ANY(v_local3_try05825)
          ToList(OBJ(_Zarg_try05813)).PutAt(CLcount,v_local3)
          } 
        }
        /* Iteration-3 */} 
      /* ERROR PROTECTION INSERTED (_Zarg-Result) */
      if ErrorIn(_Zarg_try05813) {Result = _Zarg_try05813
      } else {
      _Zarg = ToList(OBJ(_Zarg_try05813))
      var g0587I *ClaireBoolean  
      /* and:3 */{ 
        var v_and3 *ClaireBoolean  
        
        v_and3 = C_compiler.Diet_ask
        if (v_and3 == CFALSE) {g0587I = CFALSE
        } else /* arg:4 */{ 
          /* Let:5 */{ 
            var g0588UU *ClaireAny  
            /* noccur = 1 */
            /* For:6 */{ 
              var x *ClaireAny  
              _ = x
              g0588UU= CFALSE.Id()
              var x_support *ClaireList  
              x_support = l
              x_len := x_support.Length()
              for i_it := 0; i_it < x_len; i_it++ { 
                x = x_support.At(i_it)
                if ((C_class.Id() == x.Isa.Id()) || 
                    (x.Isa.IsIn(C_property) == CTRUE)) /* If:8 */{ 
                   /*v = g0588UU, s =any*/
g0588UU = CTRUE.Id()
                  break
                  /* If-8 */} 
                /* loop-7 */} 
              /* For-6 */} 
            v_and3 = F_boolean_I_any(g0588UU)
            /* Let-5 */} 
          if (v_and3 == CFALSE) {g0587I = CFALSE
          } else /* arg:5 */{ 
            g0587I = CTRUE/* arg-5 */} 
          /* arg-4 */} 
        /* and-3 */} 
      if (g0587I == CTRUE) /* If:3 */{ 
        F_Compile_warn_void()
        Core.F_tformat_string(MakeString("Non diet call ~S(~A) [254]\n"),2,MakeConstantList(self.Id(),l.Id()))
        /* If-3 */} 
      /* Let:3 */{ 
        var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
        /* noccur = 5 */
        _CL_obj.Selector = self
        _CL_obj.Args = _Zarg
        Result = EID{_CL_obj.Id(),0}
        /* Let-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: open_message @ property (throw: true) 
func E_Optimize_open_message_property (self EID,l EID) EID { 
    return /*(sm for open_message @ property= EID)*/ F_Optimize_open_message_property(ToProperty(OBJ(self)),ToList(OBJ(l)) )} 
  
// ******************************************************************
// *    Part 3: specialized c_code                                  *
// ******************************************************************
// a get message is special since it represent a direct access. The boolean
// tells if we accept a special form of the unknown value
/* {1} OPT.The go function for: daccess(self:any,b:boolean) [] */
func F_Optimize_daccess_any (self *ClaireAny ,b *ClaireBoolean ) EID { 
    var Result EID 
    if (self.Isa.IsIn(Language.C_Call) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0589 *Language.Call   = Language.To_Call(self)
        /* noccur = 2 */
        /* Let:4 */{ 
          var l *ClaireList   = g0589.Args
          /* noccur = 4 */
          /* Let:5 */{ 
            var xs *ClaireObject  
            /* noccur = 4 */
            var xs_try05926 EID 
            if ((g0589.Selector.Id() == C_get.Id()) && 
                (l.At(1-1).Isa.IsIn(C_property) == CTRUE)) /* If:6 */{ 
              /* Let:7 */{ 
                var g0593UU *ClaireClass  
                /* noccur = 1 */
                var g0593UU_try05948 EID 
                /* Let:8 */{ 
                  var g0595UU *ClaireType  
                  /* noccur = 1 */
                  var g0595UU_try05969 EID 
                  g0595UU_try05969 = Core.F_CALL(C_c_type,ARGS(l.At(2-1).ToEID()))
                  /* ERROR PROTECTION INSERTED (g0595UU-g0593UU_try05948) */
                  if ErrorIn(g0595UU_try05969) {g0593UU_try05948 = g0595UU_try05969
                  } else {
                  g0595UU = ToType(OBJ(g0595UU_try05969))
                  g0593UU_try05948 = EID{g0595UU.Class_I().Id(),0}
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (g0593UU-xs_try05926) */
                if ErrorIn(g0593UU_try05948) {xs_try05926 = g0593UU_try05948
                } else {
                g0593UU = ToClass(OBJ(g0593UU_try05948))
                xs_try05926 = EID{Core.F__at_property1(ToProperty(l.At(1-1)),g0593UU).Id(),0}
                }
                /* Let-7 */} 
              } else {
              xs_try05926 = EID{CFALSE.Id(),0}
              /* If-6 */} 
            /* ERROR PROTECTION INSERTED (xs-Result) */
            if ErrorIn(xs_try05926) {Result = xs_try05926
            } else {
            xs = ToObject(OBJ(xs_try05926))
            var g0597I *ClaireBoolean  
            if (C_slot.Id() == xs.Isa.Id()) /* If:6 */{ 
              /* Let:7 */{ 
                var g0590 *ClaireSlot   = ToSlot(xs.Id())
                /* noccur = 4 */
                g0597I = MakeBoolean((b == CTRUE) || (g0590.Range.Contains(g0590.Default) == CTRUE) || (g0590.Srange.Id() == C_any.Id()) || (g0590.Srange.Id() == C_integer.Id()))
                /* Let-7 */} 
              } else {
              g0597I = CFALSE
              /* If-6 */} 
            if (g0597I == CTRUE) /* If:6 */{ 
              /* Let:7 */{ 
                var _CL_obj *Language.CallSlot   = Language.To_CallSlot(new(Language.CallSlot).Is(Language.C_Call_slot))
                /* noccur = 7 */
                _CL_obj.Selector = ToSlot(xs.Id())
                /* update:8 */{ 
                  var va_arg1 *Language.CallSlot  
                  var va_arg2 *ClaireAny  
                  va_arg1 = _CL_obj
                  var va_arg2_try05989 EID 
                  va_arg2_try05989 = Core.F_CALL(C_c_code,ARGS(l.At(2-1).ToEID(),EID{F_Compile_psort_any(Core.F_domain_I_restriction(ToRestriction(xs.Id())).Id()).Id(),0}))
                  /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                  if ErrorIn(va_arg2_try05989) {Result = va_arg2_try05989
                  } else {
                  va_arg2 = ANY(va_arg2_try05989)
                  /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                  va_arg1.Arg = va_arg2
                  Result = va_arg2.ToEID()
                  }
                  /* update-8 */} 
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                _CL_obj.Test = CFALSE
                Result = EID{_CL_obj.Id(),0}
                }
                /* Let-7 */} 
              } else {
              Result = EID{CNULL,0}
              /* If-6 */} 
            }
            /* Let-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      } else {
      Result = EID{CNULL,0}
      /* If-2 */} 
    return Result} 
  
// The EID go function for: daccess @ any (throw: true) 
func E_Optimize_daccess_any (self EID,b EID) EID { 
    return /*(sm for daccess @ any= EID)*/ F_Optimize_daccess_any(ANY(self),ToBoolean(OBJ(b)) )} 
  
/* {1} OPT.The go function for: c_type(self:Call_slot) [] */
func F_c_type_Call_slot (self *Language.CallSlot ) *ClaireType  { 
    // use function body compiling 
return  self.Selector.Range
    } 
  
// The EID go function for: c_type @ Call_slot (throw: false) 
func E_c_type_Call_slot (self EID) EID { 
    return EID{/*(sm for c_type @ Call_slot= type)*/ F_c_type_Call_slot(Language.To_CallSlot(OBJ(self)) ).Id(),0}} 
  
/* {1} OPT.The go function for: c_type(self:Call_table) [] */
func F_c_type_Call_table (self *Language.CallTable ) *ClaireType  { 
    // use function body compiling 
return  self.Selector.Range
    } 
  
// The EID go function for: c_type @ Call_table (throw: false) 
func E_c_type_Call_table (self EID) EID { 
    return EID{/*(sm for c_type @ Call_table= type)*/ F_c_type_Call_table(Language.To_CallTable(OBJ(self)) ).Id(),0}} 
  
/* {1} OPT.The go function for: c_type(self:Call_array) [] */
func F_c_type_Call_array (self *Language.CallArray ) *ClaireType  { 
    // use function body compiling 
return  ToType(self.Test)
    } 
  
// The EID go function for: c_type @ Call_array (throw: false) 
func E_c_type_Call_array (self EID) EID { 
    return EID{/*(sm for c_type @ Call_array= type)*/ F_c_type_Call_array(Language.To_CallArray(OBJ(self)) ).Id(),0}} 
  
// write optimization: ss is put, put_store or write
// note that a put(object,x that may be unknown) is hard to compile !
// v2.4.10 -> if x = unknown OK (o.r = NULL) otherwise use store
/* {1} OPT.The go function for: c_code_write(self:Call) [] */
func F_Optimize_c_code_write_Call (self *Language.Call ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var p *ClaireAny   = self.Args.At(1-1)
      /* noccur = 9 */
      /* Let:3 */{ 
        var x *ClaireAny   = self.Args.At(2-1)
        /* noccur = 7 */
        /* Let:4 */{ 
          var y *ClaireAny   = self.Args.At(3-1)
          /* noccur = 8 */
          /* Let:5 */{ 
            var yt *ClaireAny  
            /* noccur = 6 */
            var yt_try06026 EID 
            yt_try06026 = Core.F_CALL(C_c_type,ARGS(y.ToEID()))
            /* ERROR PROTECTION INSERTED (yt-Result) */
            if ErrorIn(yt_try06026) {Result = yt_try06026
            } else {
            yt = ANY(yt_try06026)
            /* Let:6 */{ 
              var ss *ClaireProperty   = self.Selector
              /* noccur = 5 */
              /* Let:7 */{ 
                var s *ClaireAny  
                /* noccur = 12 */
                var s_try06038 EID 
                /* Let:8 */{ 
                  var g0604UU *ClaireList  
                  /* noccur = 1 */
                  var g0604UU_try06059 EID 
                  /* Construct:9 */{ 
                    var v_bag_arg *ClaireAny  
                    g0604UU_try06059= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                    var v_bag_arg_try060610 EID 
                    v_bag_arg_try060610 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
                    /* ERROR PROTECTION INSERTED (v_bag_arg-g0604UU_try06059) */
                    if ErrorIn(v_bag_arg_try060610) {g0604UU_try06059 = v_bag_arg_try060610
                    } else {
                    v_bag_arg = ANY(v_bag_arg_try060610)
                    ToList(OBJ(g0604UU_try06059)).AddFast(v_bag_arg)}
                    /* Construct-9 */} 
                  /* ERROR PROTECTION INSERTED (g0604UU-s_try06038) */
                  if ErrorIn(g0604UU_try06059) {s_try06038 = g0604UU_try06059
                  } else {
                  g0604UU = ToList(OBJ(g0604UU_try06059))
                  s_try06038 = Core.F_CALL(C_Optimize_restriction_I,ARGS(p.ToEID(),EID{g0604UU.Id(),0},EID{CTRUE.Id(),0}))
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (s-Result) */
                if ErrorIn(s_try06038) {Result = s_try06038
                } else {
                s = ANY(s_try06038)
                
                if (C_OPT.ToRemove.Contain_ask(p) == CTRUE) /* If:8 */{ 
                  Result = EID{CNIL.Id(),0}
                  } else {
                  var g0607I *ClaireBoolean  
                  if (C_slot.Id() == s.Isa.Id()) /* If:9 */{ 
                    /* Let:10 */{ 
                      var g0599 *ClaireSlot   = ToSlot(s)
                      /* noccur = 1 */
                      g0607I = MakeBoolean((ToType(yt).Included(g0599.Range) == CTRUE) || (C_compiler.Safety >= 4))
                      /* Let-10 */} 
                    } else {
                    g0607I = CFALSE
                    /* If-9 */} 
                  if (g0607I == CTRUE) /* If:9 */{ 
                    var g0608I *ClaireBoolean  
                    var g0608I_try060910 EID 
                    /* and:10 */{ 
                      var v_and10 *ClaireBoolean  
                      
                      v_and10 = Core.F__I_equal_any(y,CNULL)
                      if (v_and10 == CFALSE) {g0608I_try060910 = EID{CFALSE.Id(),0}
                      } else /* arg:11 */{ 
                        var v_and10_try061012 EID 
                        /* Let:12 */{ 
                          var g0611UU *ClaireBoolean  
                          /* noccur = 1 */
                          var g0611UU_try061213 EID 
                          /* Let:13 */{ 
                            var g0613UU *ClaireAny  
                            /* noccur = 1 */
                            var g0613UU_try061414 EID 
                            g0613UU_try061414 = Core.F_CALL(ToProperty(C__exp.Id()),ARGS(yt.ToEID(),Core.F_CALL(C_mClaire_srange,ARGS(s.ToEID()))))
                            /* ERROR PROTECTION INSERTED (g0613UU-g0611UU_try061213) */
                            if ErrorIn(g0613UU_try061414) {g0611UU_try061213 = g0613UU_try061414
                            } else {
                            g0613UU = ANY(g0613UU_try061414)
                            g0611UU_try061213 = EID{F_boolean_I_any(g0613UU).Id(),0}
                            }
                            /* Let-13 */} 
                          /* ERROR PROTECTION INSERTED (g0611UU-v_and10_try061012) */
                          if ErrorIn(g0611UU_try061213) {v_and10_try061012 = g0611UU_try061213
                          } else {
                          g0611UU = ToBoolean(OBJ(g0611UU_try061213))
                          v_and10_try061012 = EID{Core.F__I_equal_any(g0611UU.Id(),CTRUE.Id()).Id(),0}
                          }
                          /* Let-12 */} 
                        /* ERROR PROTECTION INSERTED (v_and10-g0608I_try060910) */
                        if ErrorIn(v_and10_try061012) {g0608I_try060910 = v_and10_try061012
                        } else {
                        v_and10 = ToBoolean(OBJ(v_and10_try061012))
                        if (v_and10 == CFALSE) {g0608I_try060910 = EID{CFALSE.Id(),0}
                        } else /* arg:12 */{ 
                          g0608I_try060910 = EID{CTRUE.Id(),0}/* arg-12 */} 
                        /* arg-11 */} 
                      }
                      /* and-10 */} 
                    /* ERROR PROTECTION INSERTED (g0608I-Result) */
                    if ErrorIn(g0608I_try060910) {Result = g0608I_try060910
                    } else {
                    g0608I = ToBoolean(OBJ(g0608I_try060910))
                    if (g0608I == CTRUE) /* If:10 */{ 
                      
                      F_Compile_warn_void()
                      Result = Core.F_tformat_string(MakeString("sort error in ~S: ~S is a ~S [253]\n"),2,MakeConstantList(self.Id(),y,yt))
                      } else {
                      Result = EID{CFALSE.Id(),0}
                      /* If-10 */} 
                    }
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    var g0617I *ClaireBoolean  
                    var g0617I_try061810 EID 
                    /* and:10 */{ 
                      var v_and10 *ClaireBoolean  
                      
                      v_and10 = MakeBoolean((ToType(yt).Included(ToType(OBJ(Core.F_CALL(C_range,ARGS(s.ToEID()))))) == CTRUE) || (ToType(yt).Included(ToType(C_object.Id())) == CTRUE) || (ANY(Core.F_CALL(C_mClaire_srange,ARGS(s.ToEID()))) != C_object.Id()) || (y == CNULL))
                      if (v_and10 == CFALSE) {g0617I_try061810 = EID{CFALSE.Id(),0}
                      } else /* arg:11 */{ 
                        var v_and10_try061912 EID 
                        /* or:12 */{ 
                          var v_or12 *ClaireBoolean  
                          
                          v_or12 = Core.F__I_equal_any(ss.Id(),Core.C_write.Id())
                          if (v_or12 == CTRUE) {v_and10_try061912 = EID{CTRUE.Id(),0}
                          } else /* or:13 */{ 
                            var v_or12_try062014 EID 
                            /* and:14 */{ 
                              var v_and14 *ClaireBoolean  
                              
                              var v_and14_try062115 EID 
                              v_and14_try062115 = F_Optimize_Update_ask_relation1(ToRelation(p),x,y)
                              /* ERROR PROTECTION INSERTED (v_and14-v_or12_try062014) */
                              if ErrorIn(v_and14_try062115) {v_or12_try062014 = v_and14_try062115
                              } else {
                              v_and14 = ToBoolean(OBJ(v_and14_try062115))
                              if (v_and14 == CFALSE) {v_or12_try062014 = EID{CFALSE.Id(),0}
                              } else /* arg:15 */{ 
                                v_and14 = MakeBoolean((ToRelation(p).Multivalued_ask.Id() == CFALSE.Id()) || (Core.F_get_property(C_if_write,ToObject(p)) == CNULL))
                                if (v_and14 == CFALSE) {v_or12_try062014 = EID{CFALSE.Id(),0}
                                } else /* arg:16 */{ 
                                  v_or12_try062014 = EID{CTRUE.Id(),0}/* arg-16 */} 
                                /* arg-15 */} 
                              }
                              /* and-14 */} 
                            /* ERROR PROTECTION INSERTED (v_or12-v_and10_try061912) */
                            if ErrorIn(v_or12_try062014) {v_and10_try061912 = v_or12_try062014
                            } else {
                            v_or12 = ToBoolean(OBJ(v_or12_try062014))
                            if (v_or12 == CTRUE) {v_and10_try061912 = EID{CTRUE.Id(),0}
                            } else /* or:14 */{ 
                              v_and10_try061912 = EID{CFALSE.Id(),0}/* org-14 */} 
                            /* org-13 */} 
                          }
                          /* or-12 */} 
                        /* ERROR PROTECTION INSERTED (v_and10-g0617I_try061810) */
                        if ErrorIn(v_and10_try061912) {g0617I_try061810 = v_and10_try061912
                        } else {
                        v_and10 = ToBoolean(OBJ(v_and10_try061912))
                        if (v_and10 == CFALSE) {g0617I_try061810 = EID{CFALSE.Id(),0}
                        } else /* arg:12 */{ 
                          g0617I_try061810 = EID{CTRUE.Id(),0}/* arg-12 */} 
                        /* arg-11 */} 
                      }
                      /* and-10 */} 
                    /* ERROR PROTECTION INSERTED (g0617I-Result) */
                    if ErrorIn(g0617I_try061810) {Result = g0617I_try061810
                    } else {
                    g0617I = ToBoolean(OBJ(g0617I_try061810))
                    if (g0617I == CTRUE) /* If:10 */{ 
                      /* Let:11 */{ 
                        var _Zx *ClaireAny  
                        /* noccur = 1 */
                        var _Zx_try062212 EID 
                        _Zx_try062212 = Core.F_CALL(C_c_code,ARGS(x.ToEID(),EID{F_Compile_psort_any(Core.F_domain_I_restriction(ToRestriction(s)).Id()).Id(),0}))
                        /* ERROR PROTECTION INSERTED (_Zx-Result) */
                        if ErrorIn(_Zx_try062212) {Result = _Zx_try062212
                        } else {
                        _Zx = ANY(_Zx_try062212)
                        /* Let:12 */{ 
                          var _Zy *ClaireAny  
                          /* noccur = 1 */
                          var _Zy_try062313 EID 
                          _Zy_try062313 = F_Compile_c_strict_code_any(y,F_Compile_psort_any(ANY(Core.F_CALL(C_range,ARGS(s.ToEID())))))
                          /* ERROR PROTECTION INSERTED (_Zy-Result) */
                          if ErrorIn(_Zy_try062313) {Result = _Zy_try062313
                          } else {
                          _Zy = ANY(_Zy_try062313)
                          /* Let:13 */{ 
                            var _CL_obj *Language.Update   = Language.To_Update(new(Language.Update).Is(Language.C_Update))
                            /* noccur = 16 */
                            _CL_obj.Selector = p
                            _CL_obj.Value = _Zy
                            /* update:14 */{ 
                              var va_arg1 *Language.Update  
                              var va_arg2 *ClaireAny  
                              va_arg1 = _CL_obj
                              var va_arg2_try062415 EID 
                              if (ss.Id() != Core.C_write.Id()) /* If:15 */{ 
                                va_arg2_try062415 = EID{ss.Id(),0}
                                } else {
                                va_arg2_try062415 = Core.F_CALL(C_c_code,ARGS(x.ToEID(),EID{C_any.Id(),0}))
                                /* If-15 */} 
                              /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                              if ErrorIn(va_arg2_try062415) {Result = va_arg2_try062415
                              } else {
                              va_arg2 = ANY(va_arg2_try062415)
                              /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                              va_arg1.Arg = va_arg2
                              Result = va_arg2.ToEID()
                              }
                              /* update-14 */} 
                            /* ERROR PROTECTION INSERTED (Result-Result) */
                            if !ErrorIn(Result) {
                            /* update:14 */{ 
                              var va_arg1 *Language.Update  
                              var va_arg2 *ClaireAny  
                              va_arg1 = _CL_obj
                              /* Let:15 */{ 
                                var _CL_obj *Language.CallSlot   = Language.To_CallSlot(new(Language.CallSlot).Is(Language.C_Call_slot))
                                /* noccur = 7 */
                                _CL_obj.Selector = ToSlot(s)
                                _CL_obj.Arg = _Zx
                                _CL_obj.Test = CFALSE
                                va_arg2 = _CL_obj.Id()
                                /* Let-15 */} 
                              /* ---------- now we compile update var(va_arg1) := va_arg2 ------- */
                              va_arg1.ClaireVar = va_arg2
                              /* update-14 */} 
                            Result = EID{_CL_obj.Id(),0}
                            }
                            /* Let-13 */} 
                          }
                          /* Let-12 */} 
                        }
                        /* Let-11 */} 
                      /* If!10 */}  else if (ss.Id() == C_put.Id()) /* If:10 */{ 
                      /* Let:11 */{ 
                        var g0625UU *Language.Call  
                        /* noccur = 1 */
                        /* Let:12 */{ 
                          var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                          /* noccur = 5 */
                          _CL_obj.Selector = C_store
                          _CL_obj.Args = MakeConstantList(x,
                            ANY(Core.F_CALL(C_mClaire_index,ARGS(s.ToEID()))),
                            ANY(Core.F_CALL(C_mClaire_srange,ARGS(s.ToEID()))),
                            y,
                            ANY(Core.F_CALL(C_store_ask,ARGS(p.ToEID()))))
                          g0625UU = _CL_obj
                          /* Let-12 */} 
                        Result = Core.F_CALL(C_c_code,ARGS(EID{g0625UU.Id(),0}))
                        /* Let-11 */} 
                      } else {
                      if (C_compiler.Diet_ask == CTRUE) /* If:11 */{ 
                        F_Compile_warn_void()
                        Core.F_tformat_string(MakeString("~S is not a diet call [254]"),2,MakeConstantList(self.Id()))
                        /* If-11 */} 
                      if ((C_compiler.Optimize_ask == CTRUE) && 
                          (p != C_instances.Id())) /* If:11 */{ 
                        F_Compile_notice_void()
                        Core.F_tformat_string(MakeString("poorly typed update: ~S\n"),3,MakeConstantList(self.Id()))
                        /* If-11 */} 
                      /* Let:11 */{ 
                        var g0626UU *Language.Call  
                        /* noccur = 1 */
                        /* Let:12 */{ 
                          var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                          /* noccur = 5 */
                          _CL_obj.Selector = Core.C_mClaire_update
                          _CL_obj.Args = MakeConstantList(p,
                            x,
                            ANY(Core.F_CALL(C_mClaire_index,ARGS(s.ToEID()))),
                            ANY(Core.F_CALL(C_mClaire_srange,ARGS(s.ToEID()))),
                            y)
                          g0626UU = _CL_obj
                          /* Let-12 */} 
                        Result = Core.F_CALL(C_c_code,ARGS(EID{g0626UU.Id(),0}))
                        /* Let-11 */} 
                      /* If-10 */} 
                    }
                    }
                    } else {
                    /* Let:10 */{ 
                      var _Ztype *ClaireList  
                      /* noccur = 3 */
                      var _Ztype_try062711 EID 
                      /* Iteration:11 */{ 
                        var v_list11 *ClaireList  
                        var x *ClaireAny  
                        var v_local11 *ClaireAny  
                        v_list11 = self.Args
                        _Ztype_try062711 = EID{CreateList(ToType(CEMPTY.Id()),v_list11.Length()).Id(),0}
                        for CLcount := 0; CLcount < v_list11.Length(); CLcount++{ 
                          x = v_list11.At(CLcount)
                          var v_local11_try062813 EID 
                          v_local11_try062813 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
                          /* ERROR PROTECTION INSERTED (v_local11-_Ztype_try062711) */
                          if ErrorIn(v_local11_try062813) {_Ztype_try062711 = v_local11_try062813
                          _Ztype_try062711 = v_local11_try062813
                          break
                          } else {
                          v_local11 = ANY(v_local11_try062813)
                          ToList(OBJ(_Ztype_try062711)).PutAt(CLcount,v_local11)
                          } 
                        }
                        /* Iteration-11 */} 
                      /* ERROR PROTECTION INSERTED (_Ztype-Result) */
                      if ErrorIn(_Ztype_try062711) {Result = _Ztype_try062711
                      } else {
                      _Ztype = ToList(OBJ(_Ztype_try062711))
                      /* Let:11 */{ 
                        var z *ClaireAny   = F_Optimize_restriction_I_property(ss,_Ztype,CTRUE)
                        /* noccur = 2 */
                        /* Let:12 */{ 
                          var g0629UU *ClaireList  
                          /* noccur = 1 */
                          var g0629UU_try063013 EID 
                          /* Construct:13 */{ 
                            var v_bag_arg *ClaireAny  
                            g0629UU_try063013= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                            ToList(OBJ(g0629UU_try063013)).AddFast(self.Id())
                            var v_bag_arg_try063114 EID 
                            v_bag_arg_try063114 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
                            /* ERROR PROTECTION INSERTED (v_bag_arg-g0629UU_try063013) */
                            if ErrorIn(v_bag_arg_try063114) {g0629UU_try063013 = v_bag_arg_try063114
                            } else {
                            v_bag_arg = ANY(v_bag_arg_try063114)
                            ToList(OBJ(g0629UU_try063013)).AddFast(v_bag_arg)
                            ToList(OBJ(g0629UU_try063013)).AddFast(yt)}
                            /* Construct-13 */} 
                          /* ERROR PROTECTION INSERTED (g0629UU-Result) */
                          if ErrorIn(g0629UU_try063013) {Result = g0629UU_try063013
                          } else {
                          g0629UU = ToList(OBJ(g0629UU_try063013))
                          Result = Core.F_tformat_string(MakeString("---- note: ~S is poorly typed (~S,~S) \n"),3,g0629UU)
                          }
                          /* Let-12 */} 
                        /* ERROR PROTECTION INSERTED (Result-Result) */
                        if !ErrorIn(Result) {
                        if (C_method.Id() == z.Isa.Id()) /* If:12 */{ 
                          /* Let:13 */{ 
                            var g0600 *ClaireMethod   = ToMethod(z)
                            /* noccur = 1 */
                            Result = F_Optimize_c_code_method_method1(g0600,self.Args,_Ztype)
                            /* Let-13 */} 
                          } else {
                          Result = F_Optimize_c_warn_Call(self,_Ztype.Id())
                          /* If-12 */} 
                        }
                        /* Let-11 */} 
                      }
                      /* Let-10 */} 
                    /* If-9 */} 
                  /* If-8 */} 
                }
                /* Let-7 */} 
              /* Let-6 */} 
            }
            /* Let-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_code_write @ Call (throw: true) 
func E_Optimize_c_code_write_Call (self EID) EID { 
    return /*(sm for c_code_write @ Call= EID)*/ F_Optimize_c_code_write_Call(Language.To_Call(OBJ(self)) )} 
  
// (get(p,x) =/= y) optimization. We try to use the smart form instead of the get
/* {1} OPT.The go function for: c_code_hold(p:property,x:any,y:any,b:boolean) [] */
func F_Optimize_c_code_hold_property (p *ClaireProperty ,x *ClaireAny ,y *ClaireAny ,b *ClaireBoolean ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var s *ClaireAny  
      /* noccur = 5 */
      var s_try06333 EID 
      /* Let:3 */{ 
        var g0634UU *ClaireList  
        /* noccur = 1 */
        var g0634UU_try06354 EID 
        /* Construct:4 */{ 
          var v_bag_arg *ClaireAny  
          g0634UU_try06354= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
          var v_bag_arg_try06365 EID 
          v_bag_arg_try06365 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
          /* ERROR PROTECTION INSERTED (v_bag_arg-g0634UU_try06354) */
          if ErrorIn(v_bag_arg_try06365) {g0634UU_try06354 = v_bag_arg_try06365
          } else {
          v_bag_arg = ANY(v_bag_arg_try06365)
          ToList(OBJ(g0634UU_try06354)).AddFast(v_bag_arg)}
          /* Construct-4 */} 
        /* ERROR PROTECTION INSERTED (g0634UU-s_try06333) */
        if ErrorIn(g0634UU_try06354) {s_try06333 = g0634UU_try06354
        } else {
        g0634UU = ToList(OBJ(g0634UU_try06354))
        s_try06333 = F_Optimize_restriction_I_property(p,g0634UU,CTRUE).ToEID()
        }
        /* Let-3 */} 
      /* ERROR PROTECTION INSERTED (s-Result) */
      if ErrorIn(s_try06333) {Result = s_try06333
      } else {
      s = ANY(s_try06333)
      var g0637I *ClaireBoolean  
      var g0637I_try06383 EID 
      if (C_slot.Id() == s.Isa.Id()) /* If:3 */{ 
        /* Let:4 */{ 
          var g0632 *ClaireSlot   = ToSlot(s)
          /* noccur = 1 */
          /* or:5 */{ 
            var v_or5 *ClaireBoolean  
            
            v_or5 = Equal(y,CNULL)
            if (v_or5 == CTRUE) {g0637I_try06383 = EID{CTRUE.Id(),0}
            } else /* or:6 */{ 
              var v_or5_try06397 EID 
              /* and:7 */{ 
                var v_and7 *ClaireBoolean  
                
                var v_and7_try06408 EID 
                /* Let:8 */{ 
                  var g0641UU *ClaireType  
                  /* noccur = 1 */
                  var g0641UU_try06429 EID 
                  g0641UU_try06429 = Core.F_CALL(C_c_type,ARGS(y.ToEID()))
                  /* ERROR PROTECTION INSERTED (g0641UU-v_and7_try06408) */
                  if ErrorIn(g0641UU_try06429) {v_and7_try06408 = g0641UU_try06429
                  } else {
                  g0641UU = ToType(OBJ(g0641UU_try06429))
                  v_and7_try06408 = EID{g0641UU.Included(ToType(g0632.Srange.Id())).Id(),0}
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (v_and7-v_or5_try06397) */
                if ErrorIn(v_and7_try06408) {v_or5_try06397 = v_and7_try06408
                } else {
                v_and7 = ToBoolean(OBJ(v_and7_try06408))
                if (v_and7 == CFALSE) {v_or5_try06397 = EID{CFALSE.Id(),0}
                } else /* arg:8 */{ 
                  var v_and7_try06439 EID 
                  v_and7_try06439 = F_Compile_identifiable_ask_any(y)
                  /* ERROR PROTECTION INSERTED (v_and7-v_or5_try06397) */
                  if ErrorIn(v_and7_try06439) {v_or5_try06397 = v_and7_try06439
                  } else {
                  v_and7 = ToBoolean(OBJ(v_and7_try06439))
                  if (v_and7 == CFALSE) {v_or5_try06397 = EID{CFALSE.Id(),0}
                  } else /* arg:9 */{ 
                    v_or5_try06397 = EID{CTRUE.Id(),0}/* arg-9 */} 
                  /* arg-8 */} 
                }}
                /* and-7 */} 
              /* ERROR PROTECTION INSERTED (v_or5-g0637I_try06383) */
              if ErrorIn(v_or5_try06397) {g0637I_try06383 = v_or5_try06397
              } else {
              v_or5 = ToBoolean(OBJ(v_or5_try06397))
              if (v_or5 == CTRUE) {g0637I_try06383 = EID{CTRUE.Id(),0}
              } else /* or:7 */{ 
                g0637I_try06383 = EID{CFALSE.Id(),0}/* org-7 */} 
              /* org-6 */} 
            }
            /* or-5 */} 
          /* Let-4 */} 
        } else {
        g0637I_try06383 = EID{CFALSE.Id(),0}
        /* If-3 */} 
      /* ERROR PROTECTION INSERTED (g0637I-Result) */
      if ErrorIn(g0637I_try06383) {Result = g0637I_try06383
      } else {
      g0637I = ToBoolean(OBJ(g0637I_try06383))
      if (g0637I == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var cs *Language.CallSlot  
          /* noccur = 1 */
          var cs_try06445 EID 
          /* Let:5 */{ 
            var _CL_obj *Language.CallSlot   = Language.To_CallSlot(new(Language.CallSlot).Is(Language.C_Call_slot))
            /* noccur = 7 */
            _CL_obj.Selector = ToSlot(s)
            /* update:6 */{ 
              var va_arg1 *Language.CallSlot  
              var va_arg2 *ClaireAny  
              va_arg1 = _CL_obj
              var va_arg2_try06457 EID 
              va_arg2_try06457 = Core.F_CALL(C_c_code,ARGS(x.ToEID(),EID{F_Compile_psort_any(Core.F_domain_I_restriction(ToRestriction(s)).Id()).Id(),0}))
              /* ERROR PROTECTION INSERTED (va_arg2-cs_try06445) */
              if ErrorIn(va_arg2_try06457) {cs_try06445 = va_arg2_try06457
              } else {
              va_arg2 = ANY(va_arg2_try06457)
              /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
              va_arg1.Arg = va_arg2
              cs_try06445 = va_arg2.ToEID()
              }
              /* update-6 */} 
            /* ERROR PROTECTION INSERTED (cs_try06445-cs_try06445) */
            if !ErrorIn(cs_try06445) {
            _CL_obj.Test = CFALSE
            cs_try06445 = EID{_CL_obj.Id(),0}
            }
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (cs-Result) */
          if ErrorIn(cs_try06445) {Result = cs_try06445
          } else {
          cs = Language.To_CallSlot(OBJ(cs_try06445))
          /* Let:5 */{ 
            var cm *Language.CallMethod2  
            /* noccur = 2 */
            var cm_try06466 EID 
            /* Let:6 */{ 
              var _CL_obj *Language.CallMethod2   = Language.To_CallMethod2(new(Language.CallMethod2).Is(Language.C_Call_method2))
              /* noccur = 5 */
              _CL_obj.Arg = ToMethod(Core.F__at_property1(Core.C_identical_ask,C_any).Id())
              /* update:7 */{ 
                var va_arg1 *Language.CallMethod  
                var va_arg2 *ClaireList  
                va_arg1 = Language.To_CallMethod(_CL_obj.Id())
                var va_arg2_try06478 EID 
                /* Construct:8 */{ 
                  var v_bag_arg *ClaireAny  
                  va_arg2_try06478= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                  ToList(OBJ(va_arg2_try06478)).AddFast(cs.Id())
                  var v_bag_arg_try06489 EID 
                  v_bag_arg_try06489 = Core.F_CALL(C_c_code,ARGS(y.ToEID(),Core.F_CALL(C_mClaire_srange,ARGS(s.ToEID()))))
                  /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try06478) */
                  if ErrorIn(v_bag_arg_try06489) {va_arg2_try06478 = v_bag_arg_try06489
                  } else {
                  v_bag_arg = ANY(v_bag_arg_try06489)
                  ToList(OBJ(va_arg2_try06478)).AddFast(v_bag_arg)}
                  /* Construct-8 */} 
                /* ERROR PROTECTION INSERTED (va_arg2-cm_try06466) */
                if ErrorIn(va_arg2_try06478) {cm_try06466 = va_arg2_try06478
                } else {
                va_arg2 = ToList(OBJ(va_arg2_try06478))
                /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                va_arg1.Args = va_arg2
                cm_try06466 = EID{va_arg2.Id(),0}
                }
                /* update-7 */} 
              /* ERROR PROTECTION INSERTED (cm_try06466-cm_try06466) */
              if !ErrorIn(cm_try06466) {
              cm_try06466 = EID{_CL_obj.Id(),0}
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (cm-Result) */
            if ErrorIn(cm_try06466) {Result = cm_try06466
            } else {
            cm = Language.To_CallMethod2(OBJ(cm_try06466))
            if (b == CTRUE) /* If:6 */{ 
              Result = Core.F_CALL(C_c_code,ARGS(EID{cm.Id(),0}))
              } else {
              /* Let:7 */{ 
                var g0649UU *Language.Call  
                /* noccur = 1 */
                /* Let:8 */{ 
                  var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                  /* noccur = 5 */
                  _CL_obj.Selector = Core.C_not
                  _CL_obj.Args = MakeConstantList(cm.Id())
                  g0649UU = _CL_obj
                  /* Let-8 */} 
                Result = Core.F_CALL(C_c_code,ARGS(EID{g0649UU.Id(),0}))
                /* Let-7 */} 
              /* If-6 */} 
            }
            /* Let-5 */} 
          }
          /* Let-4 */} 
        } else {
        /* Let:4 */{ 
          var l *ClaireList   = MakeConstantList(C_any.Id(),C_any.Id())
          /* noccur = 2 */
          /* Let:5 */{ 
            var g0650UU *ClaireList  
            /* noccur = 1 */
            var g0650UU_try06516 EID 
            /* Construct:6 */{ 
              var v_bag_arg *ClaireAny  
              g0650UU_try06516= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
              var v_bag_arg_try06527 EID 
              /* Let:7 */{ 
                var g0653UU *Language.Call  
                /* noccur = 1 */
                /* Let:8 */{ 
                  var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                  /* noccur = 5 */
                  _CL_obj.Selector = C_get
                  _CL_obj.Args = MakeConstantList(p.Id(),x)
                  g0653UU = _CL_obj
                  /* Let-8 */} 
                v_bag_arg_try06527 = Core.F_CALL(C_c_code,ARGS(EID{g0653UU.Id(),0},EID{C_any.Id(),0}))
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (v_bag_arg-g0650UU_try06516) */
              if ErrorIn(v_bag_arg_try06527) {g0650UU_try06516 = v_bag_arg_try06527
              } else {
              v_bag_arg = ANY(v_bag_arg_try06527)
              ToList(OBJ(g0650UU_try06516)).AddFast(v_bag_arg)
              var v_bag_arg_try06547 EID 
              v_bag_arg_try06547 = Core.F_CALL(C_c_code,ARGS(y.ToEID(),EID{C_any.Id(),0}))
              /* ERROR PROTECTION INSERTED (v_bag_arg-g0650UU_try06516) */
              if ErrorIn(v_bag_arg_try06547) {g0650UU_try06516 = v_bag_arg_try06547
              } else {
              v_bag_arg = ANY(v_bag_arg_try06547)
              ToList(OBJ(g0650UU_try06516)).AddFast(v_bag_arg)}}
              /* Construct-6 */} 
            /* ERROR PROTECTION INSERTED (g0650UU-Result) */
            if ErrorIn(g0650UU_try06516) {Result = g0650UU_try06516
            } else {
            g0650UU = ToList(OBJ(g0650UU_try06516))
            Result = F_Optimize_c_code_method_method1(ToMethod(Core.F__at_property2(ToProperty(IfThenElse((b == CTRUE),
              C__equal.Id(),
              Core.C__I_equal.Id())),l).Id()),g0650UU,l)
            }
            /* Let-5 */} 
          /* Let-4 */} 
        /* If-3 */} 
      }
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_code_hold @ property (throw: true) 
func E_Optimize_c_code_hold_property (p EID,x EID,y EID,b EID) EID { 
    return /*(sm for c_code_hold @ property= EID)*/ F_Optimize_c_code_hold_property(ToProperty(OBJ(p)),
      ANY(x),
      ANY(y),
      ToBoolean(OBJ(b)) )} 
  
// add optimization
/* {1} OPT.The go function for: c_code_add(self:Call) [] */
func F_Optimize_c_code_add_Call (self *Language.Call ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var p *ClaireProperty   = ToProperty(self.Args.At(1-1))
      /* noccur = 6 */
      /* Let:3 */{ 
        var x *ClaireAny   = self.Args.At(2-1)
        /* noccur = 6 */
        /* Let:4 */{ 
          var y *ClaireAny   = self.Args.At(3-1)
          /* noccur = 5 */
          /* Let:5 */{ 
            var s *ClaireObject  
            /* noccur = 8 */
            var s_try06566 EID 
            /* Let:6 */{ 
              var g0657UU *ClaireClass  
              /* noccur = 1 */
              var g0657UU_try06587 EID 
              /* Let:7 */{ 
                var g0659UU *ClaireType  
                /* noccur = 1 */
                var g0659UU_try06608 EID 
                /* Let:8 */{ 
                  var g0661UU *ClaireType  
                  /* noccur = 1 */
                  var g0661UU_try06629 EID 
                  g0661UU_try06629 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
                  /* ERROR PROTECTION INSERTED (g0661UU-g0659UU_try06608) */
                  if ErrorIn(g0661UU_try06629) {g0659UU_try06608 = g0661UU_try06629
                  } else {
                  g0661UU = ToType(OBJ(g0661UU_try06629))
                  g0659UU_try06608 = EID{F_Optimize_ptype_type(g0661UU).Id(),0}
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (g0659UU-g0657UU_try06587) */
                if ErrorIn(g0659UU_try06608) {g0657UU_try06587 = g0659UU_try06608
                } else {
                g0659UU = ToType(OBJ(g0659UU_try06608))
                g0657UU_try06587 = EID{g0659UU.Class_I().Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (g0657UU-s_try06566) */
              if ErrorIn(g0657UU_try06587) {s_try06566 = g0657UU_try06587
              } else {
              g0657UU = ToClass(OBJ(g0657UU_try06587))
              s_try06566 = EID{Core.F__at_property1(p,g0657UU).Id(),0}
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (s-Result) */
            if ErrorIn(s_try06566) {Result = s_try06566
            } else {
            s = ToObject(OBJ(s_try06566))
            var g0663I *ClaireBoolean  
            var g0663I_try06646 EID 
            if (C_slot.Id() == s.Isa.Id()) /* If:6 */{ 
              /* Let:7 */{ 
                var g0655 *ClaireSlot   = ToSlot(s.Id())
                /* noccur = 1 */
                /* or:8 */{ 
                  var v_or8 *ClaireBoolean  
                  
                  var v_or8_try06659 EID 
                  /* Let:9 */{ 
                    var g0666UU *ClaireType  
                    /* noccur = 1 */
                    var g0666UU_try066710 EID 
                    g0666UU_try066710 = Core.F_CALL(C_c_type,ARGS(y.ToEID()))
                    /* ERROR PROTECTION INSERTED (g0666UU-v_or8_try06659) */
                    if ErrorIn(g0666UU_try066710) {v_or8_try06659 = g0666UU_try066710
                    } else {
                    g0666UU = ToType(OBJ(g0666UU_try066710))
                    v_or8_try06659 = EID{g0666UU.Included(Core.F_member_type(g0655.Range)).Id(),0}
                    }
                    /* Let-9 */} 
                  /* ERROR PROTECTION INSERTED (v_or8-g0663I_try06646) */
                  if ErrorIn(v_or8_try06659) {g0663I_try06646 = v_or8_try06659
                  } else {
                  v_or8 = ToBoolean(OBJ(v_or8_try06659))
                  if (v_or8 == CTRUE) {g0663I_try06646 = EID{CTRUE.Id(),0}
                  } else /* or:9 */{ 
                    v_or8 = F__sup_equal_integer(C_compiler.Safety,4)
                    if (v_or8 == CTRUE) {g0663I_try06646 = EID{CTRUE.Id(),0}
                    } else /* or:10 */{ 
                      g0663I_try06646 = EID{CFALSE.Id(),0}/* org-10 */} 
                    /* org-9 */} 
                  }
                  /* or-8 */} 
                /* Let-7 */} 
              } else {
              g0663I_try06646 = EID{CFALSE.Id(),0}
              /* If-6 */} 
            /* ERROR PROTECTION INSERTED (g0663I-Result) */
            if ErrorIn(g0663I_try06646) {Result = g0663I_try06646
            } else {
            g0663I = ToBoolean(OBJ(g0663I_try06646))
            if (g0663I == CTRUE) /* If:6 */{ 
              if (F_Optimize_Update_ask_relation2(ToRelation(p.Id()),ToRelation(self.Selector.Id())) == CTRUE) /* If:7 */{ 
                /* Let:8 */{ 
                  var x2 *ClaireAny  
                  /* noccur = 1 */
                  var x2_try06689 EID 
                  x2_try06689 = Core.F_CALL(C_c_code,ARGS(x.ToEID(),EID{F_Compile_psort_any(Core.F_domain_I_restriction(ToRestriction(s.Id())).Id()).Id(),0}))
                  /* ERROR PROTECTION INSERTED (x2-Result) */
                  if ErrorIn(x2_try06689) {Result = x2_try06689
                  } else {
                  x2 = ANY(x2_try06689)
                  /* Let:9 */{ 
                    var _CL_obj *Language.Update   = Language.To_Update(new(Language.Update).Is(Language.C_Update))
                    /* noccur = 16 */
                    _CL_obj.Selector = p.Id()
                    _CL_obj.Arg = C_add.Id()
                    /* update:10 */{ 
                      var va_arg1 *Language.Update  
                      var va_arg2 *ClaireAny  
                      va_arg1 = _CL_obj
                      /* Let:11 */{ 
                        var _CL_obj *Language.CallSlot   = Language.To_CallSlot(new(Language.CallSlot).Is(Language.C_Call_slot))
                        /* noccur = 7 */
                        _CL_obj.Selector = ToSlot(s.Id())
                        _CL_obj.Arg = x2
                        _CL_obj.Test = CFALSE
                        va_arg2 = _CL_obj.Id()
                        /* Let-11 */} 
                      /* ---------- now we compile update var(va_arg1) := va_arg2 ------- */
                      va_arg1.ClaireVar = va_arg2
                      /* update-10 */} 
                    /* update:10 */{ 
                      var va_arg1 *Language.Update  
                      var va_arg2 *ClaireAny  
                      va_arg1 = _CL_obj
                      var va_arg2_try066911 EID 
                      va_arg2_try066911 = Core.F_CALL(C_c_code,ARGS(y.ToEID(),EID{F_Compile_psort_any(Core.F_member_type(ToType(OBJ(Core.F_CALL(C_range,ARGS(EID{s.Id(),0}))))).Id()).Id(),0}))
                      /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                      if ErrorIn(va_arg2_try066911) {Result = va_arg2_try066911
                      } else {
                      va_arg2 = ANY(va_arg2_try066911)
                      /* ---------- now we compile update value(va_arg1) := va_arg2 ------- */
                      va_arg1.Value = va_arg2
                      Result = va_arg2.ToEID()
                      }
                      /* update-10 */} 
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    Result = EID{_CL_obj.Id(),0}
                    }
                    /* Let-9 */} 
                  }
                  /* Let-8 */} 
                } else {
                var g0670I *ClaireBoolean  
                var g0670I_try06718 EID 
                /* and:8 */{ 
                  var v_and8 *ClaireBoolean  
                  
                  var v_and8_try06729 EID 
                  v_and8_try06729 = F_Compile_designated_ask_any(x)
                  /* ERROR PROTECTION INSERTED (v_and8-g0670I_try06718) */
                  if ErrorIn(v_and8_try06729) {g0670I_try06718 = v_and8_try06729
                  } else {
                  v_and8 = ToBoolean(OBJ(v_and8_try06729))
                  if (v_and8 == CFALSE) {g0670I_try06718 = EID{CFALSE.Id(),0}
                  } else /* arg:9 */{ 
                    v_and8 = MakeBoolean((p.Store_ask != CTRUE) && ((self.Selector.Id() == C_add_I.Id()) || 
                        (p.Inverse.Id() == CNULL)))
                    if (v_and8 == CFALSE) {g0670I_try06718 = EID{CFALSE.Id(),0}
                    } else /* arg:10 */{ 
                      g0670I_try06718 = EID{CTRUE.Id(),0}/* arg-10 */} 
                    /* arg-9 */} 
                  }
                  /* and-8 */} 
                /* ERROR PROTECTION INSERTED (g0670I-Result) */
                if ErrorIn(g0670I_try06718) {Result = g0670I_try06718
                } else {
                g0670I = ToBoolean(OBJ(g0670I_try06718))
                if (g0670I == CTRUE) /* If:8 */{ 
                  /* Let:9 */{ 
                    var x2 *ClaireAny  
                    /* noccur = 1 */
                    var x2_try067310 EID 
                    x2_try067310 = Core.F_CALL(C_c_code,ARGS(x.ToEID(),EID{F_Compile_psort_any(Core.F_domain_I_restriction(ToRestriction(s.Id())).Id()).Id(),0}))
                    /* ERROR PROTECTION INSERTED (x2-Result) */
                    if ErrorIn(x2_try067310) {Result = x2_try067310
                    } else {
                    x2 = ANY(x2_try067310)
                    /* Let:10 */{ 
                      var g0674UU *Language.Call  
                      /* noccur = 1 */
                      /* Let:11 */{ 
                        var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                        /* noccur = 13 */
                        _CL_obj.Selector = ToProperty(C_add.Id())
                        /* update:12 */{ 
                          var va_arg1 *Language.Call  
                          var va_arg2 *ClaireList  
                          va_arg1 = _CL_obj
                          /* Construct:13 */{ 
                            var v_bag_arg *ClaireAny  
                            va_arg2= ToType(CEMPTY.Id()).EmptyList()
                            /* Let:14 */{ 
                              var _CL_obj *Language.CallSlot   = Language.To_CallSlot(new(Language.CallSlot).Is(Language.C_Call_slot))
                              /* noccur = 7 */
                              _CL_obj.Selector = ToSlot(s.Id())
                              _CL_obj.Arg = x2
                              _CL_obj.Test = CFALSE
                              v_bag_arg = _CL_obj.Id()
                              /* Let-14 */} 
                            va_arg2.AddFast(v_bag_arg)
                            va_arg2.AddFast(y)/* Construct-13 */} 
                          /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                          va_arg1.Args = va_arg2
                          /* update-12 */} 
                        g0674UU = _CL_obj
                        /* Let-11 */} 
                      Result = Core.F_CALL(C_c_code,ARGS(EID{g0674UU.Id(),0}))
                      /* Let-10 */} 
                    }
                    /* Let-9 */} 
                  } else {
                  if (C_compiler.Optimize_ask == CTRUE) /* If:9 */{ 
                    F_Compile_notice_void()
                    Core.F_tformat_string(MakeString("poorly typed update: ~S\n"),3,MakeConstantList(self.Id()))
                    /* If-9 */} 
                  /* Let:9 */{ 
                    var g0675UU *ClaireList  
                    /* noccur = 1 */
                    var g0675UU_try067610 EID 
                    /* Construct:10 */{ 
                      var v_bag_arg *ClaireAny  
                      g0675UU_try067610= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                      ToList(OBJ(g0675UU_try067610)).AddFast(C_property.Id())
                      var v_bag_arg_try067711 EID 
                      v_bag_arg_try067711 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
                      /* ERROR PROTECTION INSERTED (v_bag_arg-g0675UU_try067610) */
                      if ErrorIn(v_bag_arg_try067711) {g0675UU_try067610 = v_bag_arg_try067711
                      } else {
                      v_bag_arg = ANY(v_bag_arg_try067711)
                      ToList(OBJ(g0675UU_try067610)).AddFast(v_bag_arg)
                      ToList(OBJ(g0675UU_try067610)).AddFast(C_integer.Id())
                      var v_bag_arg_try067811 EID 
                      v_bag_arg_try067811 = Core.F_CALL(C_c_type,ARGS(y.ToEID()))
                      /* ERROR PROTECTION INSERTED (v_bag_arg-g0675UU_try067610) */
                      if ErrorIn(v_bag_arg_try067811) {g0675UU_try067610 = v_bag_arg_try067811
                      } else {
                      v_bag_arg = ANY(v_bag_arg_try067811)
                      ToList(OBJ(g0675UU_try067610)).AddFast(v_bag_arg)}}
                      /* Construct-10 */} 
                    /* ERROR PROTECTION INSERTED (g0675UU-Result) */
                    if ErrorIn(g0675UU_try067610) {Result = g0675UU_try067610
                    } else {
                    g0675UU = ToList(OBJ(g0675UU_try067610))
                    Result = F_Optimize_c_code_method_method1(ToMethod(Core.F__at_property1(ToProperty(C_add_I.Id()),C_property).Id()),MakeConstantList(p.Id(),
                      x,
                      ANY(Core.F_CALL(C_mClaire_index,ARGS(EID{s.Id(),0}))),
                      y),g0675UU)
                    }
                    /* Let-9 */} 
                  /* If-8 */} 
                }
                /* If-7 */} 
              } else {
              /* Let:7 */{ 
                var g0679UU *ClaireList  
                /* noccur = 1 */
                var g0679UU_try06808 EID 
                /* Iteration:8 */{ 
                  var v_list8 *ClaireList  
                  var x *ClaireAny  
                  var v_local8 *ClaireAny  
                  v_list8 = self.Args
                  g0679UU_try06808 = EID{CreateList(ToType(CEMPTY.Id()),v_list8.Length()).Id(),0}
                  for CLcount := 0; CLcount < v_list8.Length(); CLcount++{ 
                    x = v_list8.At(CLcount)
                    var v_local8_try068110 EID 
                    v_local8_try068110 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
                    /* ERROR PROTECTION INSERTED (v_local8-g0679UU_try06808) */
                    if ErrorIn(v_local8_try068110) {g0679UU_try06808 = v_local8_try068110
                    g0679UU_try06808 = v_local8_try068110
                    break
                    } else {
                    v_local8 = ANY(v_local8_try068110)
                    ToList(OBJ(g0679UU_try06808)).PutAt(CLcount,v_local8)
                    } 
                  }
                  /* Iteration-8 */} 
                /* ERROR PROTECTION INSERTED (g0679UU-Result) */
                if ErrorIn(g0679UU_try06808) {Result = g0679UU_try06808
                } else {
                g0679UU = ToList(OBJ(g0679UU_try06808))
                Result = F_Optimize_c_code_method_method1(ToMethod(Core.F__at_property1(ToProperty(C_add.Id()),C_property).Id()),self.Args,g0679UU)
                }
                /* Let-7 */} 
              /* If-6 */} 
            }
            }
            /* Let-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_code_add @ Call (throw: true) 
func E_Optimize_c_code_add_Call (self EID) EID { 
    return /*(sm for c_code_add @ Call= EID)*/ F_Optimize_c_code_add_Call(Language.To_Call(OBJ(self)) )} 
  
// new in v3.0.59
/* {1} OPT.The go function for: c_code_add_bag(self:Call) [] */
func F_Optimize_c_code_add_bag_Call (self *Language.Call ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var _Zt1 *ClaireAny  
      /* noccur = 5 */
      var _Zt1_try06843 EID 
      _Zt1_try06843 = Core.F_CALL(C_c_type,ARGS(self.Args.At(1-1).ToEID()))
      /* ERROR PROTECTION INSERTED (_Zt1-Result) */
      if ErrorIn(_Zt1_try06843) {Result = _Zt1_try06843
      } else {
      _Zt1 = ANY(_Zt1_try06843)
      /* Let:3 */{ 
        var _Zt2 *ClaireType  
        /* noccur = 3 */
        var _Zt2_try06854 EID 
        /* Let:4 */{ 
          var g0686UU *ClaireType  
          /* noccur = 1 */
          var g0686UU_try06875 EID 
          g0686UU_try06875 = Core.F_CALL(C_c_type,ARGS(self.Args.At(2-1).ToEID()))
          /* ERROR PROTECTION INSERTED (g0686UU-_Zt2_try06854) */
          if ErrorIn(g0686UU_try06875) {_Zt2_try06854 = g0686UU_try06875
          } else {
          g0686UU = ToType(OBJ(g0686UU_try06875))
          _Zt2_try06854 = EID{F_Optimize_ptype_type(g0686UU).Id(),0}
          }
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (_Zt2-Result) */
        if ErrorIn(_Zt2_try06854) {Result = _Zt2_try06854
        } else {
        _Zt2 = ToType(OBJ(_Zt2_try06854))
        /* Let:4 */{ 
          var _Zp *ClaireProperty  
          /* noccur = 1 */
          if (((_Zt1.Isa.IsIn(C_Param) == CTRUE) && 
                (_Zt2.Included(Core.F_member_type(ToType(_Zt1))) == CTRUE)) || 
              (C_compiler.Safety >= 4)) /* If:5 */{ 
            _Zp = ToProperty(C_add_I.Id())
            } else {
            _Zp = self.Selector
            /* If-5 */} 
          /* Let:5 */{ 
            var _Zltype *ClaireList   = MakeConstantList(_Zt1,_Zt2.Id())
            /* noccur = 3 */
            /* Let:6 */{ 
              var z *ClaireAny   = F_Optimize_restriction_I_property(_Zp,_Zltype,CTRUE)
              /* noccur = 2 */
              
              if ((_Zt2.Included(Core.F_member_type(ToType(_Zt1))) != CTRUE) && 
                  (self.Selector.Id() == C_add.Id())) /* If:7 */{ 
                F_Compile_warn_void()
                Core.F_tformat_string(MakeString("the bag addition ~S is poorly typed (~S) [251] \n"),2,MakeConstantList(self.Id(),Core.F_member_type(ToType(_Zt1)).Id()))
                /* If-7 */} 
              if (C_method.Id() == z.Isa.Id()) /* If:7 */{ 
                /* Let:8 */{ 
                  var g0682 *ClaireMethod   = ToMethod(z)
                  /* noccur = 1 */
                  Result = F_Optimize_c_code_method_method1(g0682,self.Args,_Zltype)
                  /* Let-8 */} 
                } else {
                Result = F_Optimize_c_warn_Call(self,_Zltype.Id())
                /* If-7 */} 
              /* Let-6 */} 
            /* Let-5 */} 
          /* Let-4 */} 
        }
        /* Let-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_code_add_bag @ Call (throw: true) 
func E_Optimize_c_code_add_bag_Call (self EID) EID { 
    return /*(sm for c_code_add_bag @ Call= EID)*/ F_Optimize_c_code_add_bag_Call(Language.To_Call(OBJ(self)) )} 
  
// delete optimization
// <yc> 7/98 new, also needed
/* {1} OPT.The go function for: c_code_delete(self:Call) [] */
func F_Optimize_c_code_delete_Call (self *Language.Call ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var p *ClaireAny   = self.Args.At(1-1)
      /* noccur = 2 */
      /* Let:3 */{ 
        var x *ClaireAny   = self.Args.At(2-1)
        /* noccur = 3 */
        /* Let:4 */{ 
          var y *ClaireAny   = self.Args.At(3-1)
          /* noccur = 2 */
          /* Let:5 */{ 
            var s *ClaireObject  
            /* noccur = 4 */
            var s_try06896 EID 
            /* Let:6 */{ 
              var g0690UU *ClaireClass  
              /* noccur = 1 */
              var g0690UU_try06917 EID 
              /* Let:7 */{ 
                var g0692UU *ClaireType  
                /* noccur = 1 */
                var g0692UU_try06938 EID 
                g0692UU_try06938 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
                /* ERROR PROTECTION INSERTED (g0692UU-g0690UU_try06917) */
                if ErrorIn(g0692UU_try06938) {g0690UU_try06917 = g0692UU_try06938
                } else {
                g0692UU = ToType(OBJ(g0692UU_try06938))
                g0690UU_try06917 = EID{g0692UU.Class_I().Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (g0690UU-s_try06896) */
              if ErrorIn(g0690UU_try06917) {s_try06896 = g0690UU_try06917
              } else {
              g0690UU = ToClass(OBJ(g0690UU_try06917))
              s_try06896 = EID{Core.F__at_property1(ToProperty(p),g0690UU).Id(),0}
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (s-Result) */
            if ErrorIn(s_try06896) {Result = s_try06896
            } else {
            s = ToObject(OBJ(s_try06896))
            var g0694I *ClaireBoolean  
            var g0694I_try06956 EID 
            /* and:6 */{ 
              var v_and6 *ClaireBoolean  
              
              v_and6 = MakeBoolean((ToRelation(p).Inverse.Id() == CNULL))
              if (v_and6 == CFALSE) {g0694I_try06956 = EID{CFALSE.Id(),0}
              } else /* arg:7 */{ 
                var v_and6_try06968 EID 
                v_and6_try06968 = F_Compile_designated_ask_any(x)
                /* ERROR PROTECTION INSERTED (v_and6-g0694I_try06956) */
                if ErrorIn(v_and6_try06968) {g0694I_try06956 = v_and6_try06968
                } else {
                v_and6 = ToBoolean(OBJ(v_and6_try06968))
                if (v_and6 == CFALSE) {g0694I_try06956 = EID{CFALSE.Id(),0}
                } else /* arg:8 */{ 
                  var v_and6_try06979 EID 
                  if (C_slot.Id() == s.Isa.Id()) /* If:9 */{ 
                    /* Let:10 */{ 
                      var g0688 *ClaireSlot   = ToSlot(s.Id())
                      /* noccur = 1 */
                      /* or:11 */{ 
                        var v_or11 *ClaireBoolean  
                        
                        var v_or11_try069812 EID 
                        /* Let:12 */{ 
                          var g0699UU *ClaireType  
                          /* noccur = 1 */
                          var g0699UU_try070013 EID 
                          g0699UU_try070013 = Core.F_CALL(C_c_type,ARGS(y.ToEID()))
                          /* ERROR PROTECTION INSERTED (g0699UU-v_or11_try069812) */
                          if ErrorIn(g0699UU_try070013) {v_or11_try069812 = g0699UU_try070013
                          } else {
                          g0699UU = ToType(OBJ(g0699UU_try070013))
                          v_or11_try069812 = EID{g0699UU.Included(Core.F_member_type(g0688.Range)).Id(),0}
                          }
                          /* Let-12 */} 
                        /* ERROR PROTECTION INSERTED (v_or11-v_and6_try06979) */
                        if ErrorIn(v_or11_try069812) {v_and6_try06979 = v_or11_try069812
                        } else {
                        v_or11 = ToBoolean(OBJ(v_or11_try069812))
                        if (v_or11 == CTRUE) {v_and6_try06979 = EID{CTRUE.Id(),0}
                        } else /* or:12 */{ 
                          v_or11 = F__sup_equal_integer(C_compiler.Safety,4)
                          if (v_or11 == CTRUE) {v_and6_try06979 = EID{CTRUE.Id(),0}
                          } else /* or:13 */{ 
                            v_and6_try06979 = EID{CFALSE.Id(),0}/* org-13 */} 
                          /* org-12 */} 
                        }
                        /* or-11 */} 
                      /* Let-10 */} 
                    } else {
                    v_and6_try06979 = EID{CFALSE.Id(),0}
                    /* If-9 */} 
                  /* ERROR PROTECTION INSERTED (v_and6-g0694I_try06956) */
                  if ErrorIn(v_and6_try06979) {g0694I_try06956 = v_and6_try06979
                  } else {
                  v_and6 = ToBoolean(OBJ(v_and6_try06979))
                  if (v_and6 == CFALSE) {g0694I_try06956 = EID{CFALSE.Id(),0}
                  } else /* arg:9 */{ 
                    g0694I_try06956 = EID{CTRUE.Id(),0}/* arg-9 */} 
                  /* arg-8 */} 
                /* arg-7 */} 
              }}
              /* and-6 */} 
            /* ERROR PROTECTION INSERTED (g0694I-Result) */
            if ErrorIn(g0694I_try06956) {Result = g0694I_try06956
            } else {
            g0694I = ToBoolean(OBJ(g0694I_try06956))
            if (g0694I == CTRUE) /* If:6 */{ 
              /* Let:7 */{ 
                var x2 *ClaireAny  
                /* noccur = 1 */
                var x2_try07018 EID 
                x2_try07018 = Core.F_CALL(C_c_code,ARGS(x.ToEID(),EID{F_Compile_psort_any(Core.F_domain_I_restriction(ToRestriction(s.Id())).Id()).Id(),0}))
                /* ERROR PROTECTION INSERTED (x2-Result) */
                if ErrorIn(x2_try07018) {Result = x2_try07018
                } else {
                x2 = ANY(x2_try07018)
                /* Let:8 */{ 
                  var g0702UU *Language.Call  
                  /* noccur = 1 */
                  /* Let:9 */{ 
                    var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                    /* noccur = 13 */
                    _CL_obj.Selector = ToProperty(C_delete.Id())
                    /* update:10 */{ 
                      var va_arg1 *Language.Call  
                      var va_arg2 *ClaireList  
                      va_arg1 = _CL_obj
                      /* Construct:11 */{ 
                        var v_bag_arg *ClaireAny  
                        va_arg2= ToType(CEMPTY.Id()).EmptyList()
                        /* Let:12 */{ 
                          var _CL_obj *Language.CallSlot   = Language.To_CallSlot(new(Language.CallSlot).Is(Language.C_Call_slot))
                          /* noccur = 7 */
                          _CL_obj.Selector = ToSlot(s.Id())
                          _CL_obj.Arg = x2
                          _CL_obj.Test = CFALSE
                          v_bag_arg = _CL_obj.Id()
                          /* Let-12 */} 
                        va_arg2.AddFast(v_bag_arg)
                        va_arg2.AddFast(y)/* Construct-11 */} 
                      /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                      va_arg1.Args = va_arg2
                      /* update-10 */} 
                    g0702UU = _CL_obj
                    /* Let-9 */} 
                  Result = Core.F_CALL(C_c_code,ARGS(EID{g0702UU.Id(),0}))
                  /* Let-8 */} 
                }
                /* Let-7 */} 
              } else {
              /* Let:7 */{ 
                var g0703UU *ClaireList  
                /* noccur = 1 */
                var g0703UU_try07048 EID 
                /* Iteration:8 */{ 
                  var v_list8 *ClaireList  
                  var x *ClaireAny  
                  var v_local8 *ClaireAny  
                  v_list8 = self.Args
                  g0703UU_try07048 = EID{CreateList(ToType(CEMPTY.Id()),v_list8.Length()).Id(),0}
                  for CLcount := 0; CLcount < v_list8.Length(); CLcount++{ 
                    x = v_list8.At(CLcount)
                    var v_local8_try070510 EID 
                    v_local8_try070510 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
                    /* ERROR PROTECTION INSERTED (v_local8-g0703UU_try07048) */
                    if ErrorIn(v_local8_try070510) {g0703UU_try07048 = v_local8_try070510
                    g0703UU_try07048 = v_local8_try070510
                    break
                    } else {
                    v_local8 = ANY(v_local8_try070510)
                    ToList(OBJ(g0703UU_try07048)).PutAt(CLcount,v_local8)
                    } 
                  }
                  /* Iteration-8 */} 
                /* ERROR PROTECTION INSERTED (g0703UU-Result) */
                if ErrorIn(g0703UU_try07048) {Result = g0703UU_try07048
                } else {
                g0703UU = ToList(OBJ(g0703UU_try07048))
                Result = F_Optimize_c_code_method_method1(ToMethod(Core.F__at_property1(ToProperty(C_delete.Id()),C_property).Id()),self.Args,g0703UU)
                }
                /* Let-7 */} 
              /* If-6 */} 
            }
            }
            /* Let-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_code_delete @ Call (throw: true) 
func E_Optimize_c_code_delete_Call (self EID) EID { 
    return /*(sm for c_code_delete @ Call= EID)*/ F_Optimize_c_code_delete_Call(Language.To_Call(OBJ(self)) )} 
  
// cute optimization
/* {1} OPT.The go function for: c_code_not(x:Select) [] */
func F_Optimize_c_code_not_Select (x *Language.Select ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var g0706UU *Language.Call  
      /* noccur = 1 */
      /* Let:3 */{ 
        var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
        /* noccur = 11 */
        _CL_obj.Selector = Core.C_not
        /* update:4 */{ 
          var va_arg1 *Language.Call  
          var va_arg2 *ClaireList  
          va_arg1 = _CL_obj
          /* Construct:5 */{ 
            var v_bag_arg *ClaireAny  
            va_arg2= ToType(CEMPTY.Id()).EmptyList()
            /* Let:6 */{ 
              var _CL_obj *Language.For   = Language.To_For(new(Language.For).Is(Language.C_For))
              /* noccur = 5 */
              _CL_obj.ClaireVar = x.ClaireVar
              _CL_obj.SetArg = x.SetArg
              _CL_obj.Arg = Language.C_If.Make(x.Arg,Language.C_Return.Make(CTRUE.Id()),CNULL)
              v_bag_arg = _CL_obj.Id()
              /* Let-6 */} 
            va_arg2.AddFast(v_bag_arg)/* Construct-5 */} 
          /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
          va_arg1.Args = va_arg2
          /* update-4 */} 
        g0706UU = _CL_obj
        /* Let-3 */} 
      Result = Core.F_CALL(C_c_code,ARGS(EID{g0706UU.Id(),0}))
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_code_not @ Select (throw: true) 
func E_Optimize_c_code_not_Select (x EID) EID { 
    return /*(sm for c_code_not @ Select= EID)*/ F_Optimize_c_code_not_Select(Language.To_Select(OBJ(x)) )} 
  
// old % optimization
/* {1} OPT.The go function for: c_code_belong(self:Call) [] */
func F_Optimize_c_code_belong_Call (self *Language.Call ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var x *ClaireAny   = self.Args.At(1-1)
      /* noccur = 6 */
      /* Let:3 */{ 
        var y *ClaireAny   = self.Args.At(2-1)
        /* noccur = 8 */
        /* Let:4 */{ 
          var _Ztype *ClaireList  
          /* noccur = 4 */
          var _Ztype_try07075 EID 
          /* Construct:5 */{ 
            var v_bag_arg *ClaireAny  
            _Ztype_try07075= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
            var v_bag_arg_try07086 EID 
            v_bag_arg_try07086 = Core.F_CALL(C_c_type,ARGS(y.ToEID()))
            /* ERROR PROTECTION INSERTED (v_bag_arg-_Ztype_try07075) */
            if ErrorIn(v_bag_arg_try07086) {_Ztype_try07075 = v_bag_arg_try07086
            } else {
            v_bag_arg = ANY(v_bag_arg_try07086)
            ToList(OBJ(_Ztype_try07075)).AddFast(v_bag_arg)
            var v_bag_arg_try07096 EID 
            v_bag_arg_try07096 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
            /* ERROR PROTECTION INSERTED (v_bag_arg-_Ztype_try07075) */
            if ErrorIn(v_bag_arg_try07096) {_Ztype_try07075 = v_bag_arg_try07096
            } else {
            v_bag_arg = ANY(v_bag_arg_try07096)
            ToList(OBJ(_Ztype_try07075)).AddFast(v_bag_arg)}}
            /* Construct-5 */} 
          /* ERROR PROTECTION INSERTED (_Ztype-Result) */
          if ErrorIn(_Ztype_try07075) {Result = _Ztype_try07075
          } else {
          _Ztype = ToList(OBJ(_Ztype_try07075))
          if (C_set.Id() == y.Isa.Id()) /* If:5 */{ 
            /* Let:6 */{ 
              var _CL_obj *Language.Or   = Language.To_Or(new(Language.Or).Is(Language.C_Or))
              /* noccur = 9 */
              /* update:7 */{ 
                var va_arg1 *Language.Or  
                var va_arg2 *ClaireList  
                va_arg1 = _CL_obj
                var va_arg2_try07108 EID 
                /* Let:8 */{ 
                  var z_bag *ClaireList   = ToType(CEMPTY.Id()).EmptyList()
                  /* noccur = 2 */
                  /* For:9 */{ 
                    var z *ClaireAny  
                    _ = z
                    va_arg2_try07108= EID{CFALSE.Id(),0}
                    var z_support *ClaireList  
                    var z_support_try071110 EID 
                    z_support_try071110 = Core.F_enumerate_any(y)
                    /* ERROR PROTECTION INSERTED (z_support-va_arg2_try07108) */
                    if ErrorIn(z_support_try071110) {va_arg2_try07108 = z_support_try071110
                    } else {
                    z_support = ToList(OBJ(z_support_try071110))
                    z_len := z_support.Length()
                    for i_it := 0; i_it < z_len; i_it++ { 
                      z = z_support.At(i_it)
                      var void_try11 EID 
                      _ = void_try11
                      /* Let:11 */{ 
                        var g0712UU *ClaireAny  
                        /* noccur = 1 */
                        var g0712UU_try071312 EID 
                        /* Let:12 */{ 
                          var g0714UU *Language.Call  
                          /* noccur = 1 */
                          /* Let:13 */{ 
                            var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                            /* noccur = 5 */
                            _CL_obj.Selector = ToProperty(C__equal.Id())
                            _CL_obj.Args = MakeConstantList(x,z)
                            g0714UU = _CL_obj
                            /* Let-13 */} 
                          g0712UU_try071312 = Core.F_CALL(C_c_code,ARGS(EID{g0714UU.Id(),0},EID{C_any.Id(),0}))
                          /* Let-12 */} 
                        /* ERROR PROTECTION INSERTED (g0712UU-void_try11) */
                        if ErrorIn(g0712UU_try071312) {void_try11 = g0712UU_try071312
                        } else {
                        g0712UU = ANY(g0712UU_try071312)
                        void_try11 = EID{z_bag.AddFast(g0712UU).Id(),0}
                        }
                        /* Let-11 */} 
                      /* ERROR PROTECTION INSERTED (void_try11-va_arg2_try07108) */
                      if ErrorIn(void_try11) {va_arg2_try07108 = void_try11
                      va_arg2_try07108 = void_try11
                      break
                      } else {
                      }}
                      /* loop-10 */} 
                    /* For-9 */} 
                  /* ERROR PROTECTION INSERTED (va_arg2_try07108-va_arg2_try07108) */
                  if !ErrorIn(va_arg2_try07108) {
                  va_arg2_try07108 = EID{z_bag.Id(),0}
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                if ErrorIn(va_arg2_try07108) {Result = va_arg2_try07108
                } else {
                va_arg2 = ToList(OBJ(va_arg2_try07108))
                /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                va_arg1.Args = va_arg2
                Result = EID{va_arg2.Id(),0}
                }
                /* update-7 */} 
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              Result = EID{_CL_obj.Id(),0}
              }
              /* Let-6 */} 
            /* If!5 */}  else if (ToType(_Ztype.At(1-1)).Included(ToType(C_list.Id())) == CTRUE) /* If:5 */{ 
            Result = F_Optimize_c_code_method_method1(ToMethod(Core.F__at_property2(C_contain_ask,MakeConstantList(C_list.Id(),C_any.Id())).Id()),MakeConstantList(y,x),_Ztype)
            /* If!5 */}  else if (ToType(_Ztype.At(1-1)).Included(ToType(C_set.Id())) == CTRUE) /* If:5 */{ 
            Result = F_Optimize_c_code_method_method1(ToMethod(Core.F__at_property2(C_contain_ask,MakeConstantList(C_set.Id(),C_any.Id())).Id()),MakeConstantList(y,x),_Ztype)
            /* If!5 */}  else if (y == C_object.Id()) /* If:5 */{ 
            Result = F_Optimize_c_code_method_method1(ToMethod(Core.F__at_property2(ToProperty(C__Z.Id()),MakeConstantList(C_any.Id(),C_class.Id())).Id()),MakeConstantList(x,y),MakeConstantList(C_any.Id(),C_class.Id()))
            } else {
            Result = Core.F_CALL(C_Optimize_member_code,ARGS(y.ToEID(),x.ToEID()))
            /* If-5 */} 
          }
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_code_belong @ Call (throw: true) 
func E_Optimize_c_code_belong_Call (self EID) EID { 
    return /*(sm for c_code_belong @ Call= EID)*/ F_Optimize_c_code_belong_Call(Language.To_Call(OBJ(self)) )} 
  
// nth optimization for arrays (the selector may also be get)
/* {1} OPT.The go function for: c_code_nth(self:Call) [] */
func F_Optimize_c_code_nth_Call (self *Language.Call ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var l *ClaireList   = self.Args
      /* noccur = 10 */
      /* Let:3 */{ 
        var x *ClaireAny   = l.At(1-1)
        /* noccur = 17 */
        /* Let:4 */{ 
          var p *ClaireProperty   = self.Selector
          /* noccur = 6 */
          /* Let:5 */{ 
            var t *ClaireAny  
            /* noccur = 4 */
            var t_try07176 EID 
            t_try07176 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
            /* ERROR PROTECTION INSERTED (t-Result) */
            if ErrorIn(t_try07176) {Result = t_try07176
            } else {
            t = ANY(t_try07176)
            /* Let:6 */{ 
              var mt *ClaireType   = Core.F_member_type(ToType(t))
              /* noccur = 3 */
              /* Let:7 */{ 
                var r *ClaireAny  
                /* noccur = 2 */
                var r_try07188 EID 
                /* Let:8 */{ 
                  var g0719UU *ClaireList  
                  /* noccur = 1 */
                  var g0719UU_try07209 EID 
                  /* Iteration:9 */{ 
                    var v_list9 *ClaireList  
                    var u *ClaireAny  
                    var v_local9 *ClaireAny  
                    v_list9 = l
                    g0719UU_try07209 = EID{CreateList(ToType(CEMPTY.Id()),v_list9.Length()).Id(),0}
                    for CLcount := 0; CLcount < v_list9.Length(); CLcount++{ 
                      u = v_list9.At(CLcount)
                      var v_local9_try072111 EID 
                      v_local9_try072111 = Core.F_CALL(C_c_type,ARGS(u.ToEID()))
                      /* ERROR PROTECTION INSERTED (v_local9-g0719UU_try07209) */
                      if ErrorIn(v_local9_try072111) {g0719UU_try07209 = v_local9_try072111
                      g0719UU_try07209 = v_local9_try072111
                      break
                      } else {
                      v_local9 = ANY(v_local9_try072111)
                      ToList(OBJ(g0719UU_try07209)).PutAt(CLcount,v_local9)
                      } 
                    }
                    /* Iteration-9 */} 
                  /* ERROR PROTECTION INSERTED (g0719UU-r_try07188) */
                  if ErrorIn(g0719UU_try07209) {r_try07188 = g0719UU_try07209
                  } else {
                  g0719UU = ToList(OBJ(g0719UU_try07209))
                  r_try07188 = F_Optimize_restriction_I_property(p,g0719UU,CTRUE).ToEID()
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (r-Result) */
                if ErrorIn(r_try07188) {Result = r_try07188
                } else {
                r = ANY(r_try07188)
                if (C_OPT.ToRemove.Contain_ask(x) == CTRUE) /* If:8 */{ 
                  Result = EID{CNIL.Id(),0}
                  } else {
                  var g0722I *ClaireBoolean  
                  var g0722I_try07239 EID 
                  /* and:9 */{ 
                    var v_and9 *ClaireBoolean  
                    
                    if (C_table.Id() == x.Isa.Id()) /* If:10 */{ 
                      /* Let:11 */{ 
                        var g0715 *ClaireTable   = ToTable(x)
                        /* noccur = 1 */
                        v_and9 = Equal(C_integer.Id(),g0715.Params.Isa.Id())
                        /* Let-11 */} 
                      } else {
                      v_and9 = CFALSE
                      /* If-10 */} 
                    if (v_and9 == CFALSE) {g0722I_try07239 = EID{CFALSE.Id(),0}
                    } else /* arg:10 */{ 
                      var v_and9_try072411 EID 
                      /* or:11 */{ 
                        var v_or11 *ClaireBoolean  
                        
                        var v_or11_try072512 EID 
                        /* Let:12 */{ 
                          var g0726UU *ClaireType  
                          /* noccur = 1 */
                          var g0726UU_try072713 EID 
                          g0726UU_try072713 = Core.F_CALL(C_c_type,ARGS(l.At(2-1).ToEID()))
                          /* ERROR PROTECTION INSERTED (g0726UU-v_or11_try072512) */
                          if ErrorIn(g0726UU_try072713) {v_or11_try072512 = g0726UU_try072713
                          } else {
                          g0726UU = ToType(OBJ(g0726UU_try072713))
                          v_or11_try072512 = EID{g0726UU.Included(ToType(OBJ(Core.F_CALL(C_domain,ARGS(x.ToEID()))))).Id(),0}
                          }
                          /* Let-12 */} 
                        /* ERROR PROTECTION INSERTED (v_or11-v_and9_try072411) */
                        if ErrorIn(v_or11_try072512) {v_and9_try072411 = v_or11_try072512
                        } else {
                        v_or11 = ToBoolean(OBJ(v_or11_try072512))
                        if (v_or11 == CTRUE) {v_and9_try072411 = EID{CTRUE.Id(),0}
                        } else /* or:12 */{ 
                          v_or11 = MakeBoolean((p.Id() == C_nth.Id()) && (C_compiler.Safety > 2))
                          if (v_or11 == CTRUE) {v_and9_try072411 = EID{CTRUE.Id(),0}
                          } else /* or:13 */{ 
                            v_and9_try072411 = EID{CFALSE.Id(),0}/* org-13 */} 
                          /* org-12 */} 
                        }
                        /* or-11 */} 
                      /* ERROR PROTECTION INSERTED (v_and9-g0722I_try07239) */
                      if ErrorIn(v_and9_try072411) {g0722I_try07239 = v_and9_try072411
                      } else {
                      v_and9 = ToBoolean(OBJ(v_and9_try072411))
                      if (v_and9 == CFALSE) {g0722I_try07239 = EID{CFALSE.Id(),0}
                      } else /* arg:11 */{ 
                        g0722I_try07239 = EID{CTRUE.Id(),0}/* arg-11 */} 
                      /* arg-10 */} 
                    }
                    /* and-9 */} 
                  /* ERROR PROTECTION INSERTED (g0722I-Result) */
                  if ErrorIn(g0722I_try07239) {Result = g0722I_try07239
                  } else {
                  g0722I = ToBoolean(OBJ(g0722I_try07239))
                  if (g0722I == CTRUE) /* If:9 */{ 
                    /* Let:10 */{ 
                      var _CL_obj *Language.CallTable   = Language.To_CallTable(new(Language.CallTable).Is(Language.C_Call_table))
                      /* noccur = 7 */
                      _CL_obj.Selector = ToTable(x)
                      /* update:11 */{ 
                        var va_arg1 *Language.CallTable  
                        var va_arg2 *ClaireAny  
                        va_arg1 = _CL_obj
                        var va_arg2_try072812 EID 
                        va_arg2_try072812 = Core.F_CALL(C_c_code,ARGS(l.At(2-1).ToEID(),EID{C_integer.Id(),0}))
                        /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                        if ErrorIn(va_arg2_try072812) {Result = va_arg2_try072812
                        } else {
                        va_arg2 = ANY(va_arg2_try072812)
                        /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                        va_arg1.Arg = va_arg2
                        Result = va_arg2.ToEID()
                        }
                        /* update-11 */} 
                      /* ERROR PROTECTION INSERTED (Result-Result) */
                      if !ErrorIn(Result) {
                      /* update:11 */{ 
                        var va_arg1 *Language.CallTable  
                        var va_arg2 *ClaireBoolean  
                        va_arg1 = _CL_obj
                        var va_arg2_try072912 EID 
                        /* Let:12 */{ 
                          var g0730UU *ClaireBoolean  
                          /* noccur = 1 */
                          var g0730UU_try073113 EID 
                          /* or:13 */{ 
                            var v_or13 *ClaireBoolean  
                            
                            var v_or13_try073214 EID 
                            v_or13_try073214 = Core.F_BELONG(ANY(Core.F_CALL(C_default,ARGS(x.ToEID()))),ANY(Core.F_CALL(C_range,ARGS(x.ToEID()))))
                            /* ERROR PROTECTION INSERTED (v_or13-g0730UU_try073113) */
                            if ErrorIn(v_or13_try073214) {g0730UU_try073113 = v_or13_try073214
                            } else {
                            v_or13 = ToBoolean(OBJ(v_or13_try073214))
                            if (v_or13 == CTRUE) {g0730UU_try073113 = EID{CTRUE.Id(),0}
                            } else /* or:14 */{ 
                              v_or13 = Equal(ANY(Core.F_CALL(C_default,ARGS(x.ToEID()))),MakeInteger(0).Id())
                              if (v_or13 == CTRUE) {g0730UU_try073113 = EID{CTRUE.Id(),0}
                              } else /* or:15 */{ 
                                v_or13 = Equal(p.Id(),C_get.Id())
                                if (v_or13 == CTRUE) {g0730UU_try073113 = EID{CTRUE.Id(),0}
                                } else /* or:16 */{ 
                                  g0730UU_try073113 = EID{CFALSE.Id(),0}/* org-16 */} 
                                /* org-15 */} 
                              /* org-14 */} 
                            }
                            /* or-13 */} 
                          /* ERROR PROTECTION INSERTED (g0730UU-va_arg2_try072912) */
                          if ErrorIn(g0730UU_try073113) {va_arg2_try072912 = g0730UU_try073113
                          } else {
                          g0730UU = ToBoolean(OBJ(g0730UU_try073113))
                          va_arg2_try072912 = EID{g0730UU.Not.Id(),0}
                          }
                          /* Let-12 */} 
                        /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                        if ErrorIn(va_arg2_try072912) {Result = va_arg2_try072912
                        } else {
                        va_arg2 = ToBoolean(OBJ(va_arg2_try072912))
                        /* ---------- now we compile update iClaire/test(va_arg1) := va_arg2 ------- */
                        va_arg1.Test = va_arg2
                        Result = EID{va_arg2.Id(),0}
                        }
                        /* update-11 */} 
                      /* ERROR PROTECTION INSERTED (Result-Result) */
                      if !ErrorIn(Result) {
                      Result = EID{_CL_obj.Id(),0}
                      }}
                      /* Let-10 */} 
                    } else {
                    var g0733I *ClaireBoolean  
                    var g0733I_try073410 EID 
                    /* and:10 */{ 
                      var v_and10 *ClaireBoolean  
                      
                      if (C_table.Id() == x.Isa.Id()) /* If:11 */{ 
                        /* Let:12 */{ 
                          var g0716 *ClaireTable   = ToTable(x)
                          /* noccur = 1 */
                          v_and10 = g0716.Params.Isa.IsIn(C_list)
                          /* Let-12 */} 
                        } else {
                        v_and10 = CFALSE
                        /* If-11 */} 
                      if (v_and10 == CFALSE) {g0733I_try073410 = EID{CFALSE.Id(),0}
                      } else /* arg:11 */{ 
                        v_and10 = Equal(MakeInteger(l.Length()).Id(),MakeInteger(3).Id())
                        if (v_and10 == CFALSE) {g0733I_try073410 = EID{CFALSE.Id(),0}
                        } else /* arg:12 */{ 
                          var v_and10_try073513 EID 
                          /* or:13 */{ 
                            var v_or13 *ClaireBoolean  
                            
                            var v_or13_try073614 EID 
                            /* Let:14 */{ 
                              var g0737UU *ClaireTuple  
                              /* noccur = 1 */
                              var g0737UU_try073815 EID 
                              /* Let:15 */{ 
                                var g0739UU *ClaireList  
                                /* noccur = 1 */
                                var g0739UU_try074016 EID 
                                /* Construct:16 */{ 
                                  var v_bag_arg *ClaireAny  
                                  g0739UU_try074016= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                                  var v_bag_arg_try074117 EID 
                                  v_bag_arg_try074117 = Core.F_CALL(C_c_type,ARGS(l.At(2-1).ToEID()))
                                  /* ERROR PROTECTION INSERTED (v_bag_arg-g0739UU_try074016) */
                                  if ErrorIn(v_bag_arg_try074117) {g0739UU_try074016 = v_bag_arg_try074117
                                  } else {
                                  v_bag_arg = ANY(v_bag_arg_try074117)
                                  ToList(OBJ(g0739UU_try074016)).AddFast(v_bag_arg)
                                  var v_bag_arg_try074217 EID 
                                  v_bag_arg_try074217 = Core.F_CALL(C_c_type,ARGS(l.At(3-1).ToEID()))
                                  /* ERROR PROTECTION INSERTED (v_bag_arg-g0739UU_try074016) */
                                  if ErrorIn(v_bag_arg_try074217) {g0739UU_try074016 = v_bag_arg_try074217
                                  } else {
                                  v_bag_arg = ANY(v_bag_arg_try074217)
                                  ToList(OBJ(g0739UU_try074016)).AddFast(v_bag_arg)}}
                                  /* Construct-16 */} 
                                /* ERROR PROTECTION INSERTED (g0739UU-g0737UU_try073815) */
                                if ErrorIn(g0739UU_try074016) {g0737UU_try073815 = g0739UU_try074016
                                } else {
                                g0739UU = ToList(OBJ(g0739UU_try074016))
                                g0737UU_try073815 = EID{g0739UU.Tuple_I().Id(),0}
                                }
                                /* Let-15 */} 
                              /* ERROR PROTECTION INSERTED (g0737UU-v_or13_try073614) */
                              if ErrorIn(g0737UU_try073815) {v_or13_try073614 = g0737UU_try073815
                              } else {
                              g0737UU = ToTuple(OBJ(g0737UU_try073815))
                              v_or13_try073614 = EID{ToType(g0737UU.Id()).Included(ToType(OBJ(Core.F_CALL(C_domain,ARGS(x.ToEID()))))).Id(),0}
                              }
                              /* Let-14 */} 
                            /* ERROR PROTECTION INSERTED (v_or13-v_and10_try073513) */
                            if ErrorIn(v_or13_try073614) {v_and10_try073513 = v_or13_try073614
                            } else {
                            v_or13 = ToBoolean(OBJ(v_or13_try073614))
                            if (v_or13 == CTRUE) {v_and10_try073513 = EID{CTRUE.Id(),0}
                            } else /* or:14 */{ 
                              v_or13 = Core.F__sup_integer(C_compiler.Safety,2)
                              if (v_or13 == CTRUE) {v_and10_try073513 = EID{CTRUE.Id(),0}
                              } else /* or:15 */{ 
                                v_and10_try073513 = EID{CFALSE.Id(),0}/* org-15 */} 
                              /* org-14 */} 
                            }
                            /* or-13 */} 
                          /* ERROR PROTECTION INSERTED (v_and10-g0733I_try073410) */
                          if ErrorIn(v_and10_try073513) {g0733I_try073410 = v_and10_try073513
                          } else {
                          v_and10 = ToBoolean(OBJ(v_and10_try073513))
                          if (v_and10 == CFALSE) {g0733I_try073410 = EID{CFALSE.Id(),0}
                          } else /* arg:13 */{ 
                            g0733I_try073410 = EID{CTRUE.Id(),0}/* arg-13 */} 
                          /* arg-12 */} 
                        /* arg-11 */} 
                      }
                      /* and-10 */} 
                    /* ERROR PROTECTION INSERTED (g0733I-Result) */
                    if ErrorIn(g0733I_try073410) {Result = g0733I_try073410
                    } else {
                    g0733I = ToBoolean(OBJ(g0733I_try073410))
                    if (g0733I == CTRUE) /* If:10 */{ 
                      /* Let:11 */{ 
                        var _CL_obj *Language.CallTable   = Language.To_CallTable(new(Language.CallTable).Is(Language.C_Call_table))
                        /* noccur = 11 */
                        _CL_obj.Selector = ToTable(x)
                        /* update:12 */{ 
                          var va_arg1 *Language.CallTable  
                          var va_arg2 *ClaireAny  
                          va_arg1 = _CL_obj
                          var va_arg2_try074313 EID 
                          /* Let:13 */{ 
                            var _CL_obj *Language.List   = Language.To_List(new(Language.List).Is(Language.C_List))
                            /* noccur = 3 */
                            /* update:14 */{ 
                              var va_arg1 *Language.Construct  
                              var va_arg2 *ClaireList  
                              va_arg1 = Language.To_Construct(_CL_obj.Id())
                              var va_arg2_try074415 EID 
                              /* Construct:15 */{ 
                                var v_bag_arg *ClaireAny  
                                va_arg2_try074415= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                                var v_bag_arg_try074516 EID 
                                v_bag_arg_try074516 = Core.F_CALL(C_c_code,ARGS(l.At(2-1).ToEID(),EID{C_integer.Id(),0}))
                                /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try074415) */
                                if ErrorIn(v_bag_arg_try074516) {va_arg2_try074415 = v_bag_arg_try074516
                                } else {
                                v_bag_arg = ANY(v_bag_arg_try074516)
                                ToList(OBJ(va_arg2_try074415)).AddFast(v_bag_arg)
                                var v_bag_arg_try074616 EID 
                                v_bag_arg_try074616 = Core.F_CALL(C_c_code,ARGS(l.At(3-1).ToEID(),EID{C_integer.Id(),0}))
                                /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try074415) */
                                if ErrorIn(v_bag_arg_try074616) {va_arg2_try074415 = v_bag_arg_try074616
                                } else {
                                v_bag_arg = ANY(v_bag_arg_try074616)
                                ToList(OBJ(va_arg2_try074415)).AddFast(v_bag_arg)}}
                                /* Construct-15 */} 
                              /* ERROR PROTECTION INSERTED (va_arg2-va_arg2_try074313) */
                              if ErrorIn(va_arg2_try074415) {va_arg2_try074313 = va_arg2_try074415
                              } else {
                              va_arg2 = ToList(OBJ(va_arg2_try074415))
                              /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                              va_arg1.Args = va_arg2
                              va_arg2_try074313 = EID{va_arg2.Id(),0}
                              }
                              /* update-14 */} 
                            /* ERROR PROTECTION INSERTED (va_arg2_try074313-va_arg2_try074313) */
                            if !ErrorIn(va_arg2_try074313) {
                            va_arg2_try074313 = EID{_CL_obj.Id(),0}
                            }
                            /* Let-13 */} 
                          /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                          if ErrorIn(va_arg2_try074313) {Result = va_arg2_try074313
                          } else {
                          va_arg2 = ANY(va_arg2_try074313)
                          /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                          va_arg1.Arg = va_arg2
                          Result = va_arg2.ToEID()
                          }
                          /* update-12 */} 
                        /* ERROR PROTECTION INSERTED (Result-Result) */
                        if !ErrorIn(Result) {
                        /* update:12 */{ 
                          var va_arg1 *Language.CallTable  
                          var va_arg2 *ClaireBoolean  
                          va_arg1 = _CL_obj
                          var va_arg2_try074713 EID 
                          /* Let:13 */{ 
                            var g0748UU *ClaireBoolean  
                            /* noccur = 1 */
                            var g0748UU_try074914 EID 
                            /* or:14 */{ 
                              var v_or14 *ClaireBoolean  
                              
                              var v_or14_try075015 EID 
                              v_or14_try075015 = Core.F_BELONG(ANY(Core.F_CALL(C_default,ARGS(x.ToEID()))),ANY(Core.F_CALL(C_range,ARGS(x.ToEID()))))
                              /* ERROR PROTECTION INSERTED (v_or14-g0748UU_try074914) */
                              if ErrorIn(v_or14_try075015) {g0748UU_try074914 = v_or14_try075015
                              } else {
                              v_or14 = ToBoolean(OBJ(v_or14_try075015))
                              if (v_or14 == CTRUE) {g0748UU_try074914 = EID{CTRUE.Id(),0}
                              } else /* or:15 */{ 
                                v_or14 = Equal(ANY(Core.F_CALL(C_default,ARGS(x.ToEID()))),MakeInteger(0).Id())
                                if (v_or14 == CTRUE) {g0748UU_try074914 = EID{CTRUE.Id(),0}
                                } else /* or:16 */{ 
                                  v_or14 = Equal(p.Id(),C_get.Id())
                                  if (v_or14 == CTRUE) {g0748UU_try074914 = EID{CTRUE.Id(),0}
                                  } else /* or:17 */{ 
                                    g0748UU_try074914 = EID{CFALSE.Id(),0}/* org-17 */} 
                                  /* org-16 */} 
                                /* org-15 */} 
                              }
                              /* or-14 */} 
                            /* ERROR PROTECTION INSERTED (g0748UU-va_arg2_try074713) */
                            if ErrorIn(g0748UU_try074914) {va_arg2_try074713 = g0748UU_try074914
                            } else {
                            g0748UU = ToBoolean(OBJ(g0748UU_try074914))
                            va_arg2_try074713 = EID{g0748UU.Not.Id(),0}
                            }
                            /* Let-13 */} 
                          /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                          if ErrorIn(va_arg2_try074713) {Result = va_arg2_try074713
                          } else {
                          va_arg2 = ToBoolean(OBJ(va_arg2_try074713))
                          /* ---------- now we compile update iClaire/test(va_arg1) := va_arg2 ------- */
                          va_arg1.Test = va_arg2
                          Result = EID{va_arg2.Id(),0}
                          }
                          /* update-12 */} 
                        /* ERROR PROTECTION INSERTED (Result-Result) */
                        if !ErrorIn(Result) {
                        Result = EID{_CL_obj.Id(),0}
                        }}
                        /* Let-11 */} 
                      /* If!10 */}  else if ((ToType(t).Included(ToType(C_array.Id())) == CTRUE) && 
                        (((p.Id() == C_nth_get.Id()) || 
                            (C_compiler.Safety > 2)) && 
                          ((mt.Included(ToType(C_float.Id())) == CTRUE) || 
                              (Equal(Core.F__exp_type(mt,ToType(C_float.Id())).Id(),CEMPTY.Id()) == CTRUE)))) /* If:10 */{ 
                      /* Let:11 */{ 
                        var g0751UU *ClaireAny  
                        /* noccur = 1 */
                        var g0751UU_try075312 EID 
                        g0751UU_try075312 = Core.F_CALL(C_c_code,ARGS(x.ToEID(),EID{C_array.Id(),0}))
                        /* ERROR PROTECTION INSERTED (g0751UU-Result) */
                        if ErrorIn(g0751UU_try075312) {Result = g0751UU_try075312
                        } else {
                        g0751UU = ANY(g0751UU_try075312)
                        /* Let:12 */{ 
                          var g0752UU *ClaireAny  
                          /* noccur = 1 */
                          var g0752UU_try075413 EID 
                          g0752UU_try075413 = Core.F_CALL(C_c_code,ARGS(l.At(2-1).ToEID(),EID{C_integer.Id(),0}))
                          /* ERROR PROTECTION INSERTED (g0752UU-Result) */
                          if ErrorIn(g0752UU_try075413) {Result = g0752UU_try075413
                          } else {
                          g0752UU = ANY(g0752UU_try075413)
                          Result = Language.C_Call_array.Make(g0751UU,g0752UU,mt.Id()).ToEID()
                          }
                          /* Let-12 */} 
                        }
                        /* Let-11 */} 
                      /* If!10 */}  else if (C_method.Id() == r.Isa.Id()) /* If:10 */{ 
                      if ((C_compiler.Optimize_ask == CTRUE) && 
                          ((ToType(t).Included(ToType(C_array.Id())) == CTRUE) || 
                              (ToType(t).Included(ToType(C_table.Id())) == CTRUE))) /* If:11 */{ 
                        F_Compile_notice_void()
                        Core.F_tformat_string(MakeString("poorly typed call: ~S\n"),3,MakeConstantList(self.Id()))
                        /* If-11 */} 
                      /* Let:11 */{ 
                        var g0755UU *ClaireList  
                        /* noccur = 1 */
                        var g0755UU_try075612 EID 
                        /* Iteration:12 */{ 
                          var v_list12 *ClaireList  
                          var x *ClaireAny  
                          var v_local12 *ClaireAny  
                          v_list12 = self.Args
                          g0755UU_try075612 = EID{CreateList(ToType(CEMPTY.Id()),v_list12.Length()).Id(),0}
                          for CLcount := 0; CLcount < v_list12.Length(); CLcount++{ 
                            x = v_list12.At(CLcount)
                            var v_local12_try075714 EID 
                            v_local12_try075714 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
                            /* ERROR PROTECTION INSERTED (v_local12-g0755UU_try075612) */
                            if ErrorIn(v_local12_try075714) {g0755UU_try075612 = v_local12_try075714
                            g0755UU_try075612 = v_local12_try075714
                            break
                            } else {
                            v_local12 = ANY(v_local12_try075714)
                            ToList(OBJ(g0755UU_try075612)).PutAt(CLcount,v_local12)
                            } 
                          }
                          /* Iteration-12 */} 
                        /* ERROR PROTECTION INSERTED (g0755UU-Result) */
                        if ErrorIn(g0755UU_try075612) {Result = g0755UU_try075612
                        } else {
                        g0755UU = ToList(OBJ(g0755UU_try075612))
                        Result = F_Optimize_c_code_method_method1(ToMethod(r),self.Args,g0755UU)
                        }
                        /* Let-11 */} 
                      } else {
                      /* Let:11 */{ 
                        var g0758UU *ClaireList  
                        /* noccur = 1 */
                        var g0758UU_try075912 EID 
                        /* Iteration:12 */{ 
                          var v_list12 *ClaireList  
                          var x *ClaireAny  
                          var v_local12 *ClaireAny  
                          v_list12 = self.Args
                          g0758UU_try075912 = EID{CreateList(ToType(CEMPTY.Id()),v_list12.Length()).Id(),0}
                          for CLcount := 0; CLcount < v_list12.Length(); CLcount++{ 
                            x = v_list12.At(CLcount)
                            var v_local12_try076014 EID 
                            v_local12_try076014 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
                            /* ERROR PROTECTION INSERTED (v_local12-g0758UU_try075912) */
                            if ErrorIn(v_local12_try076014) {g0758UU_try075912 = v_local12_try076014
                            g0758UU_try075912 = v_local12_try076014
                            break
                            } else {
                            v_local12 = ANY(v_local12_try076014)
                            ToList(OBJ(g0758UU_try075912)).PutAt(CLcount,v_local12)
                            } 
                          }
                          /* Iteration-12 */} 
                        /* ERROR PROTECTION INSERTED (g0758UU-Result) */
                        if ErrorIn(g0758UU_try075912) {Result = g0758UU_try075912
                        } else {
                        g0758UU = ToList(OBJ(g0758UU_try075912))
                        Result = F_Optimize_c_warn_property(p,self.Args,g0758UU)
                        }
                        /* Let-11 */} 
                      /* If-10 */} 
                    }
                    /* If-9 */} 
                  }
                  /* If-8 */} 
                }
                /* Let-7 */} 
              /* Let-6 */} 
            }
            /* Let-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_code_nth @ Call (throw: true) 
func E_Optimize_c_code_nth_Call (self EID) EID { 
    return /*(sm for c_code_nth @ Call= EID)*/ F_Optimize_c_code_nth_Call(Language.To_Call(OBJ(self)) )} 
  
// nth= optimization for tables
/* {1} OPT.The go function for: c_code_table(self:Call) [] */
func F_Optimize_c_code_table_Call (self *Language.Call ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var sp *ClaireProperty   = self.Selector
      /* noccur = 3 */
      /* Let:3 */{ 
        var p *ClaireTable   = ToTable(self.Args.At(1-1))
        /* noccur = 10 */
        /* Let:4 */{ 
          var x *ClaireAny   = self.Args.At(2-1)
          /* noccur = 3 */
          /* Let:5 */{ 
            var y *ClaireAny   = self.Args.At(3-1)
            /* noccur = 3 */
            if (C_OPT.ToRemove.Contain_ask(p.Id()) == CTRUE) /* If:6 */{ 
              Result = EID{CNIL.Id(),0}
              } else {
              var g0761I *ClaireBoolean  
              var g0761I_try07627 EID 
              /* or:7 */{ 
                var v_or7 *ClaireBoolean  
                
                v_or7 = Equal(sp.Id(),C_put.Id())
                if (v_or7 == CTRUE) {g0761I_try07627 = EID{CTRUE.Id(),0}
                } else /* or:8 */{ 
                  var v_or7_try07639 EID 
                  /* and:9 */{ 
                    var v_and9 *ClaireBoolean  
                    
                    var v_and9_try076410 EID 
                    /* or:10 */{ 
                      var v_or10 *ClaireBoolean  
                      
                      var v_or10_try076511 EID 
                      /* Let:11 */{ 
                        var g0766UU *ClaireType  
                        /* noccur = 1 */
                        var g0766UU_try076712 EID 
                        g0766UU_try076712 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
                        /* ERROR PROTECTION INSERTED (g0766UU-v_or10_try076511) */
                        if ErrorIn(g0766UU_try076712) {v_or10_try076511 = g0766UU_try076712
                        } else {
                        g0766UU = ToType(OBJ(g0766UU_try076712))
                        v_or10_try076511 = EID{g0766UU.Included(p.Domain).Id(),0}
                        }
                        /* Let-11 */} 
                      /* ERROR PROTECTION INSERTED (v_or10-v_and9_try076410) */
                      if ErrorIn(v_or10_try076511) {v_and9_try076410 = v_or10_try076511
                      } else {
                      v_or10 = ToBoolean(OBJ(v_or10_try076511))
                      if (v_or10 == CTRUE) {v_and9_try076410 = EID{CTRUE.Id(),0}
                      } else /* or:11 */{ 
                        v_or10 = F__sup_equal_integer(C_compiler.Safety,5)
                        if (v_or10 == CTRUE) {v_and9_try076410 = EID{CTRUE.Id(),0}
                        } else /* or:12 */{ 
                          v_and9_try076410 = EID{CFALSE.Id(),0}/* org-12 */} 
                        /* org-11 */} 
                      }
                      /* or-10 */} 
                    /* ERROR PROTECTION INSERTED (v_and9-v_or7_try07639) */
                    if ErrorIn(v_and9_try076410) {v_or7_try07639 = v_and9_try076410
                    } else {
                    v_and9 = ToBoolean(OBJ(v_and9_try076410))
                    if (v_and9 == CFALSE) {v_or7_try07639 = EID{CFALSE.Id(),0}
                    } else /* arg:10 */{ 
                      var v_and9_try076811 EID 
                      /* or:11 */{ 
                        var v_or11 *ClaireBoolean  
                        
                        var v_or11_try076912 EID 
                        /* Let:12 */{ 
                          var g0770UU *ClaireType  
                          /* noccur = 1 */
                          var g0770UU_try077113 EID 
                          g0770UU_try077113 = Core.F_CALL(C_c_type,ARGS(y.ToEID()))
                          /* ERROR PROTECTION INSERTED (g0770UU-v_or11_try076912) */
                          if ErrorIn(g0770UU_try077113) {v_or11_try076912 = g0770UU_try077113
                          } else {
                          g0770UU = ToType(OBJ(g0770UU_try077113))
                          v_or11_try076912 = EID{g0770UU.Included(p.Range).Id(),0}
                          }
                          /* Let-12 */} 
                        /* ERROR PROTECTION INSERTED (v_or11-v_and9_try076811) */
                        if ErrorIn(v_or11_try076912) {v_and9_try076811 = v_or11_try076912
                        } else {
                        v_or11 = ToBoolean(OBJ(v_or11_try076912))
                        if (v_or11 == CTRUE) {v_and9_try076811 = EID{CTRUE.Id(),0}
                        } else /* or:12 */{ 
                          v_or11 = F__sup_equal_integer(C_compiler.Safety,4)
                          if (v_or11 == CTRUE) {v_and9_try076811 = EID{CTRUE.Id(),0}
                          } else /* or:13 */{ 
                            v_and9_try076811 = EID{CFALSE.Id(),0}/* org-13 */} 
                          /* org-12 */} 
                        }
                        /* or-11 */} 
                      /* ERROR PROTECTION INSERTED (v_and9-v_or7_try07639) */
                      if ErrorIn(v_and9_try076811) {v_or7_try07639 = v_and9_try076811
                      } else {
                      v_and9 = ToBoolean(OBJ(v_and9_try076811))
                      if (v_and9 == CFALSE) {v_or7_try07639 = EID{CFALSE.Id(),0}
                      } else /* arg:11 */{ 
                        v_or7_try07639 = EID{CTRUE.Id(),0}/* arg-11 */} 
                      /* arg-10 */} 
                    }}
                    /* and-9 */} 
                  /* ERROR PROTECTION INSERTED (v_or7-g0761I_try07627) */
                  if ErrorIn(v_or7_try07639) {g0761I_try07627 = v_or7_try07639
                  } else {
                  v_or7 = ToBoolean(OBJ(v_or7_try07639))
                  if (v_or7 == CTRUE) {g0761I_try07627 = EID{CTRUE.Id(),0}
                  } else /* or:9 */{ 
                    g0761I_try07627 = EID{CFALSE.Id(),0}/* org-9 */} 
                  /* org-8 */} 
                }
                /* or-7 */} 
              /* ERROR PROTECTION INSERTED (g0761I-Result) */
              if ErrorIn(g0761I_try07627) {Result = g0761I_try07627
              } else {
              g0761I = ToBoolean(OBJ(g0761I_try07627))
              if (g0761I == CTRUE) /* If:7 */{ 
                var g0772I *ClaireBoolean  
                var g0772I_try07738 EID 
                /* and:8 */{ 
                  var v_and8 *ClaireBoolean  
                  
                  var v_and8_try07749 EID 
                  v_and8_try07749 = F_Optimize_Update_ask_relation1(ToRelation(p.Id()),x,y)
                  /* ERROR PROTECTION INSERTED (v_and8-g0772I_try07738) */
                  if ErrorIn(v_and8_try07749) {g0772I_try07738 = v_and8_try07749
                  } else {
                  v_and8 = ToBoolean(OBJ(v_and8_try07749))
                  if (v_and8 == CFALSE) {g0772I_try07738 = EID{CFALSE.Id(),0}
                  } else /* arg:9 */{ 
                    v_and8 = MakeBoolean((p.Params.Isa.IsIn(C_list) == CTRUE) || (C_integer.Id() == p.Params.Isa.Id()))
                    if (v_and8 == CFALSE) {g0772I_try07738 = EID{CFALSE.Id(),0}
                    } else /* arg:10 */{ 
                      g0772I_try07738 = EID{CTRUE.Id(),0}/* arg-10 */} 
                    /* arg-9 */} 
                  }
                  /* and-8 */} 
                /* ERROR PROTECTION INSERTED (g0772I-Result) */
                if ErrorIn(g0772I_try07738) {Result = g0772I_try07738
                } else {
                g0772I = ToBoolean(OBJ(g0772I_try07738))
                if (g0772I == CTRUE) /* If:8 */{ 
                  /* Let:9 */{ 
                    var _Zx *ClaireAny  
                    /* noccur = 2 */
                    var _Zx_try077510 EID 
                    _Zx_try077510 = Core.F_CALL(C_c_code,ARGS(x.ToEID(),EID{C_any.Id(),0}))
                    /* ERROR PROTECTION INSERTED (_Zx-Result) */
                    if ErrorIn(_Zx_try077510) {Result = _Zx_try077510
                    } else {
                    _Zx = ANY(_Zx_try077510)
                    /* Let:10 */{ 
                      var _Zy *ClaireAny  
                      /* noccur = 1 */
                      var _Zy_try077611 EID 
                      _Zy_try077611 = Core.F_CALL(C_c_code,ARGS(y.ToEID(),EID{C_any.Id(),0}))
                      /* ERROR PROTECTION INSERTED (_Zy-Result) */
                      if ErrorIn(_Zy_try077611) {Result = _Zy_try077611
                      } else {
                      _Zy = ANY(_Zy_try077611)
                      /* Let:11 */{ 
                        var _CL_obj *Language.Update   = Language.To_Update(new(Language.Update).Is(Language.C_Update))
                        /* noccur = 16 */
                        _CL_obj.Selector = p.Id()
                        _CL_obj.Value = _Zy
                        _CL_obj.Arg = IfThenElse((sp.Id() == C_put.Id()),
                          C_put.Id(),
                          _Zx)
                        /* update:12 */{ 
                          var va_arg1 *Language.Update  
                          var va_arg2 *ClaireAny  
                          va_arg1 = _CL_obj
                          /* Let:13 */{ 
                            var _CL_obj *Language.CallTable   = Language.To_CallTable(new(Language.CallTable).Is(Language.C_Call_table))
                            /* noccur = 7 */
                            _CL_obj.Selector = p
                            _CL_obj.Arg = _Zx
                            _CL_obj.Test = CFALSE
                            va_arg2 = _CL_obj.Id()
                            /* Let-13 */} 
                          /* ---------- now we compile update var(va_arg1) := va_arg2 ------- */
                          va_arg1.ClaireVar = va_arg2
                          /* update-12 */} 
                        Result = EID{_CL_obj.Id(),0}
                        /* Let-11 */} 
                      }
                      /* Let-10 */} 
                    }
                    /* Let-9 */} 
                  /* If!8 */}  else if ((sp.Id() == C_put.Id()) || 
                    ((p.Inverse.Id() == CNULL) && 
                        (p.IfWrite == CNULL))) /* If:8 */{ 
                  if (C_compiler.Optimize_ask == CTRUE) /* If:9 */{ 
                    F_Compile_notice_void()
                    Core.F_tformat_string(MakeString("poorly typed update: ~S\n"),3,MakeConstantList(self.Id()))
                    /* If-9 */} 
                  Result = F_Optimize_c_code_method_method1(ToMethod(Core.F__at_property1(C_put,C_table).Id()),self.Args,MakeConstantList(C_table.Id(),C_any.Id(),C_any.Id()))
                  } else {
                  Result = F_Optimize_c_code_method_method1(ToMethod(Core.F__at_property1(C_nth_put,C_table).Id()),self.Args,MakeConstantList(C_table.Id(),C_any.Id(),C_any.Id()))
                  /* If-8 */} 
                }
                } else {
                Result = F_Optimize_c_code_method_method1(ToMethod(Core.F__at_property1(C_nth_equal,C_table).Id()),self.Args,MakeConstantList(C_table.Id(),C_any.Id(),C_any.Id()))
                /* If-7 */} 
              }
              /* If-6 */} 
            /* Let-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_code_table @ Call (throw: true) 
func E_Optimize_c_code_table_Call (self EID) EID { 
    return /*(sm for c_code_table @ Call= EID)*/ F_Optimize_c_code_table_Call(Language.To_Call(OBJ(self)) )} 
  
// version for arrays
/* {1} OPT.The go function for: c_code_array(self:Call) [] */
func F_Optimize_c_code_array_Call (self *Language.Call ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var sp *ClaireProperty   = self.Selector
      /* noccur = 2 */
      /* Let:3 */{ 
        var p *ClaireAny   = self.Args.At(1-1)
        /* noccur = 3 */
        /* Let:4 */{ 
          var tp *ClaireAny  
          /* noccur = 3 */
          var tp_try07775 EID 
          tp_try07775 = Core.F_CALL(C_c_type,ARGS(p.ToEID()))
          /* ERROR PROTECTION INSERTED (tp-Result) */
          if ErrorIn(tp_try07775) {Result = tp_try07775
          } else {
          tp = ANY(tp_try07775)
          /* Let:5 */{ 
            var mt *ClaireType   = Core.F_member_type(ToType(tp))
            /* noccur = 4 */
            /* Let:6 */{ 
              var x *ClaireAny   = self.Args.At(2-1)
              /* noccur = 1 */
              /* Let:7 */{ 
                var y *ClaireAny   = self.Args.At(3-1)
                /* noccur = 2 */
                /* Let:8 */{ 
                  var typeok *ClaireBoolean  
                  /* noccur = 2 */
                  var typeok_try07789 EID 
                  /* or:9 */{ 
                    var v_or9 *ClaireBoolean  
                    
                    var v_or9_try077910 EID 
                    /* Let:10 */{ 
                      var g0780UU *ClaireType  
                      /* noccur = 1 */
                      var g0780UU_try078111 EID 
                      g0780UU_try078111 = Core.F_CALL(C_c_type,ARGS(y.ToEID()))
                      /* ERROR PROTECTION INSERTED (g0780UU-v_or9_try077910) */
                      if ErrorIn(g0780UU_try078111) {v_or9_try077910 = g0780UU_try078111
                      } else {
                      g0780UU = ToType(OBJ(g0780UU_try078111))
                      v_or9_try077910 = EID{g0780UU.Included(Core.F_member_type(ToType(tp))).Id(),0}
                      }
                      /* Let-10 */} 
                    /* ERROR PROTECTION INSERTED (v_or9-typeok_try07789) */
                    if ErrorIn(v_or9_try077910) {typeok_try07789 = v_or9_try077910
                    } else {
                    v_or9 = ToBoolean(OBJ(v_or9_try077910))
                    if (v_or9 == CTRUE) {typeok_try07789 = EID{CTRUE.Id(),0}
                    } else /* or:10 */{ 
                      v_or9 = F__sup_equal_integer(C_compiler.Safety,4)
                      if (v_or9 == CTRUE) {typeok_try07789 = EID{CTRUE.Id(),0}
                      } else /* or:11 */{ 
                        typeok_try07789 = EID{CFALSE.Id(),0}/* org-11 */} 
                      /* org-10 */} 
                    }
                    /* or-9 */} 
                  /* ERROR PROTECTION INSERTED (typeok-Result) */
                  if ErrorIn(typeok_try07789) {Result = typeok_try07789
                  } else {
                  typeok = ToBoolean(OBJ(typeok_try07789))
                  if (((sp.Id() == C_nth_put.Id()) || 
                        (typeok == CTRUE)) && 
                      ((mt.Included(ToType(C_float.Id())) == CTRUE) || 
                          (Equal(Core.F__exp_type(mt,ToType(C_float.Id())).Id(),CEMPTY.Id()) == CTRUE))) /* If:9 */{ 
                    /* Let:10 */{ 
                      var _Zx *ClaireAny  
                      /* noccur = 1 */
                      var _Zx_try078211 EID 
                      _Zx_try078211 = Core.F_CALL(C_c_code,ARGS(x.ToEID(),EID{C_integer.Id(),0}))
                      /* ERROR PROTECTION INSERTED (_Zx-Result) */
                      if ErrorIn(_Zx_try078211) {Result = _Zx_try078211
                      } else {
                      _Zx = ANY(_Zx_try078211)
                      /* Let:11 */{ 
                        var _Zy *ClaireAny  
                        /* noccur = 1 */
                        var _Zy_try078312 EID 
                        /* Let:12 */{ 
                          var g0784UU *ClaireClass  
                          /* noccur = 1 */
                          if (mt.Included(ToType(C_float.Id())) == CTRUE) /* If:13 */{ 
                            g0784UU = C_float
                            } else {
                            g0784UU = C_any
                            /* If-13 */} 
                          _Zy_try078312 = Core.F_CALL(C_c_code,ARGS(y.ToEID(),EID{g0784UU.Id(),0}))
                          /* Let-12 */} 
                        /* ERROR PROTECTION INSERTED (_Zy-Result) */
                        if ErrorIn(_Zy_try078312) {Result = _Zy_try078312
                        } else {
                        _Zy = ANY(_Zy_try078312)
                        /* Let:12 */{ 
                          var _CL_obj *Language.Update   = Language.To_Update(new(Language.Update).Is(Language.C_Update))
                          /* noccur = 8 */
                          _CL_obj.Selector = p
                          _CL_obj.Value = _Zy
                          _CL_obj.Arg = C_put.Id()
                          /* update:13 */{ 
                            var va_arg1 *Language.Update  
                            var va_arg2 *ClaireAny  
                            va_arg1 = _CL_obj
                            var va_arg2_try078514 EID 
                            /* Let:14 */{ 
                              var g0786UU *ClaireAny  
                              /* noccur = 1 */
                              var g0786UU_try078715 EID 
                              g0786UU_try078715 = Core.F_CALL(C_c_code,ARGS(p.ToEID(),EID{C_array.Id(),0}))
                              /* ERROR PROTECTION INSERTED (g0786UU-va_arg2_try078514) */
                              if ErrorIn(g0786UU_try078715) {va_arg2_try078514 = g0786UU_try078715
                              } else {
                              g0786UU = ANY(g0786UU_try078715)
                              va_arg2_try078514 = Language.C_Call_array.Make(g0786UU,_Zx,mt.Id()).ToEID()
                              }
                              /* Let-14 */} 
                            /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                            if ErrorIn(va_arg2_try078514) {Result = va_arg2_try078514
                            } else {
                            va_arg2 = ANY(va_arg2_try078514)
                            /* ---------- now we compile update var(va_arg1) := va_arg2 ------- */
                            va_arg1.ClaireVar = va_arg2
                            Result = va_arg2.ToEID()
                            }
                            /* update-13 */} 
                          /* ERROR PROTECTION INSERTED (Result-Result) */
                          if !ErrorIn(Result) {
                          Result = EID{_CL_obj.Id(),0}
                          }
                          /* Let-12 */} 
                        }
                        /* Let-11 */} 
                      }
                      /* Let-10 */} 
                    } else {
                    if (C_compiler.Optimize_ask == CTRUE) /* If:10 */{ 
                      F_Compile_notice_void()
                      Core.F_tformat_string(MakeString("poorly typed update: ~S\n"),3,MakeConstantList(self.Id()))
                      /* If-10 */} 
                    Result = F_Optimize_c_code_method_method1(ToMethod(Core.F__at_property1(ToProperty(IfThenElse((typeok == CTRUE),
                      C_nth_put.Id(),
                      sp.Id())),C_array).Id()),self.Args,MakeConstantList(tp,C_any.Id(),C_any.Id()))
                    /* If-9 */} 
                  }
                  /* Let-8 */} 
                /* Let-7 */} 
              /* Let-6 */} 
            /* Let-5 */} 
          }
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_code_array @ Call (throw: true) 
func E_Optimize_c_code_array_Call (self EID) EID { 
    return /*(sm for c_code_array @ Call= EID)*/ F_Optimize_c_code_array_Call(Language.To_Call(OBJ(self)) )} 
  
// can we use the special UDATE form ?nth
/* {1} OPT.The go function for: Update?(p:relation,x:any,y:any) [] */
func F_Optimize_Update_ask_relation1 (p *ClaireRelation ,x *ClaireAny ,y *ClaireAny ) EID { 
    var Result EID 
    /* and:2 */{ 
      var v_and2 *ClaireBoolean  
      
      v_and2 = Core.F__I_equal_any(p.Id(),C_inverse.Id())
      if (v_and2 == CFALSE) {Result = EID{CFALSE.Id(),0}
      } else /* arg:3 */{ 
        var v_and2_try07904 EID 
        /* or:4 */{ 
          var v_or4 *ClaireBoolean  
          
          v_or4 = MakeBoolean((p.IfWrite != CNULL) && (p.IfWrite.Isa.IsIn(C_list) != CTRUE))
          if (v_or4 == CTRUE) {v_and2_try07904 = EID{CTRUE.Id(),0}
          } else /* or:5 */{ 
            var v_or4_try07916 EID 
            /* and:6 */{ 
              var v_and6 *ClaireBoolean  
              
              v_and6 = MakeBoolean((p.Inverse.Id() == CNULL))
              if (v_and6 == CFALSE) {v_or4_try07916 = EID{CFALSE.Id(),0}
              } else /* arg:7 */{ 
                if (C_table.Id() == p.Isa.Id()) /* If:8 */{ 
                  /* Let:9 */{ 
                    var g0788 *ClaireTable   = ToTable(p.Id())
                    /* noccur = 1 */
                    v_and6 = Equal(C_integer.Id(),g0788.Params.Isa.Id())
                    /* Let-9 */} 
                  } else {
                  v_and6 = CTRUE
                  /* If-8 */} 
                if (v_and6 == CFALSE) {v_or4_try07916 = EID{CFALSE.Id(),0}
                } else /* arg:8 */{ 
                  var v_and6_try07929 EID 
                  if (p.Store_ask == CTRUE) /* If:9 */{ 
                    /* and:10 */{ 
                      var v_and10 *ClaireBoolean  
                      
                      var v_and10_try079311 EID 
                      v_and10_try079311 = F_Compile_designated_ask_any(x)
                      /* ERROR PROTECTION INSERTED (v_and10-v_and6_try07929) */
                      if ErrorIn(v_and10_try079311) {v_and6_try07929 = v_and10_try079311
                      } else {
                      v_and10 = ToBoolean(OBJ(v_and10_try079311))
                      if (v_and10 == CFALSE) {v_and6_try07929 = EID{CFALSE.Id(),0}
                      } else /* arg:11 */{ 
                        var v_and10_try079412 EID 
                        v_and10_try079412 = F_Compile_designated_ask_any(y)
                        /* ERROR PROTECTION INSERTED (v_and10-v_and6_try07929) */
                        if ErrorIn(v_and10_try079412) {v_and6_try07929 = v_and10_try079412
                        } else {
                        v_and10 = ToBoolean(OBJ(v_and10_try079412))
                        if (v_and10 == CFALSE) {v_and6_try07929 = EID{CFALSE.Id(),0}
                        } else /* arg:12 */{ 
                          v_and10 = Core.F__I_equal_any(F_boolean_I_any(ANY(Core.F_CALL(C_Optimize_multi_ask,ARGS(EID{p.Id(),0})))).Id(),CTRUE.Id())
                          if (v_and10 == CFALSE) {v_and6_try07929 = EID{CFALSE.Id(),0}
                          } else /* arg:13 */{ 
                            var v_and10_try079514 EID 
                            /* or:14 */{ 
                              var v_or14 *ClaireBoolean  
                              
                              var v_or14_try079615 EID 
                              v_or14_try079615 = F_Compile_identifiable_ask_any(y)
                              /* ERROR PROTECTION INSERTED (v_or14-v_and10_try079514) */
                              if ErrorIn(v_or14_try079615) {v_and10_try079514 = v_or14_try079615
                              } else {
                              v_or14 = ToBoolean(OBJ(v_or14_try079615))
                              if (v_or14 == CTRUE) {v_and10_try079514 = EID{CTRUE.Id(),0}
                              } else /* or:15 */{ 
                                var v_or14_try079716 EID 
                                /* Let:16 */{ 
                                  var g0798UU *ClaireType  
                                  /* noccur = 1 */
                                  var g0798UU_try079917 EID 
                                  g0798UU_try079917 = Core.F_CALL(C_c_type,ARGS(y.ToEID()))
                                  /* ERROR PROTECTION INSERTED (g0798UU-v_or14_try079716) */
                                  if ErrorIn(g0798UU_try079917) {v_or14_try079716 = g0798UU_try079917
                                  } else {
                                  g0798UU = ToType(OBJ(g0798UU_try079917))
                                  v_or14_try079716 = EID{g0798UU.Included(ToType(C_float.Id())).Id(),0}
                                  }
                                  /* Let-16 */} 
                                /* ERROR PROTECTION INSERTED (v_or14-v_and10_try079514) */
                                if ErrorIn(v_or14_try079716) {v_and10_try079514 = v_or14_try079716
                                } else {
                                v_or14 = ToBoolean(OBJ(v_or14_try079716))
                                if (v_or14 == CTRUE) {v_and10_try079514 = EID{CTRUE.Id(),0}
                                } else /* or:16 */{ 
                                  v_and10_try079514 = EID{CFALSE.Id(),0}/* org-16 */} 
                                /* org-15 */} 
                              }}
                              /* or-14 */} 
                            /* ERROR PROTECTION INSERTED (v_and10-v_and6_try07929) */
                            if ErrorIn(v_and10_try079514) {v_and6_try07929 = v_and10_try079514
                            } else {
                            v_and10 = ToBoolean(OBJ(v_and10_try079514))
                            if (v_and10 == CFALSE) {v_and6_try07929 = EID{CFALSE.Id(),0}
                            } else /* arg:14 */{ 
                              v_and6_try07929 = EID{CTRUE.Id(),0}/* arg-14 */} 
                            /* arg-13 */} 
                          /* arg-12 */} 
                        /* arg-11 */} 
                      }}}
                      /* and-10 */} 
                    } else {
                    v_and6_try07929 = EID{CTRUE.Id(),0}
                    /* If-9 */} 
                  /* ERROR PROTECTION INSERTED (v_and6-v_or4_try07916) */
                  if ErrorIn(v_and6_try07929) {v_or4_try07916 = v_and6_try07929
                  } else {
                  v_and6 = ToBoolean(OBJ(v_and6_try07929))
                  if (v_and6 == CFALSE) {v_or4_try07916 = EID{CFALSE.Id(),0}
                  } else /* arg:9 */{ 
                    v_or4_try07916 = EID{CTRUE.Id(),0}/* arg-9 */} 
                  /* arg-8 */} 
                /* arg-7 */} 
              }
              /* and-6 */} 
            /* ERROR PROTECTION INSERTED (v_or4-v_and2_try07904) */
            if ErrorIn(v_or4_try07916) {v_and2_try07904 = v_or4_try07916
            } else {
            v_or4 = ToBoolean(OBJ(v_or4_try07916))
            if (v_or4 == CTRUE) {v_and2_try07904 = EID{CTRUE.Id(),0}
            } else /* or:6 */{ 
              v_and2_try07904 = EID{CFALSE.Id(),0}/* org-6 */} 
            /* org-5 */} 
          }
          /* or-4 */} 
        /* ERROR PROTECTION INSERTED (v_and2-Result) */
        if ErrorIn(v_and2_try07904) {Result = v_and2_try07904
        } else {
        v_and2 = ToBoolean(OBJ(v_and2_try07904))
        if (v_and2 == CFALSE) {Result = EID{CFALSE.Id(),0}
        } else /* arg:4 */{ 
          Result = EID{CTRUE.Id(),0}/* arg-4 */} 
        /* arg-3 */} 
      }
      /* and-2 */} 
    return Result} 
  
// The EID go function for: Update? @ list<type_expression>(relation, any, any) (throw: true) 
func E_Optimize_Update_ask_relation1 (p EID,x EID,y EID) EID { 
    return /*(sm for Update? @ list<type_expression>(relation, any, any)= EID)*/ F_Optimize_Update_ask_relation1(ToRelation(OBJ(p)),ANY(x),ANY(y) )} 
  
// we do not use an Update form for add
/* {1} OPT.The go function for: Update?(p:relation,s:relation) [] */
func F_Optimize_Update_ask_relation2 (p *ClaireRelation ,s *ClaireRelation ) *ClaireBoolean  { 
    // use function body compiling 
return  MakeBoolean((p.IfWrite != CNULL) && (p.IfWrite.Isa.IsIn(C_list) != CTRUE) && (s.Id() == C_add.Id()))
    } 
  
// The EID go function for: Update? @ list<type_expression>(relation, relation) (throw: false) 
func E_Optimize_Update_ask_relation2 (p EID,s EID) EID { 
    return EID{/*(sm for Update? @ list<type_expression>(relation, relation)= boolean)*/ F_Optimize_Update_ask_relation2(ToRelation(OBJ(p)),ToRelation(OBJ(s)) ).Id(),0}} 
  
// Update returns the value .. <yc:0.01 -> needed in CLAIRE 2.4 !!!>
/* {1} OPT.The go function for: c_type(self:Update) [] */
func F_c_type_Update (self *Language.Update ) *ClaireType  { 
    // use function body compiling 
return  ToType(C_void.Id())
    } 
  
// The EID go function for: c_type @ Update (throw: false) 
func E_c_type_Update (self EID) EID { 
    return EID{/*(sm for c_type @ Update= type)*/ F_c_type_Update(Language.To_Update(OBJ(self)) ).Id(),0}} 
  
// ******************************************************************
// *    Part 3: Method optimizer                                    *
// ******************************************************************
// a basic method is assumed to be compiled if it is in the right module
/* {1} OPT.The go function for: c_code_method(self:method,l:list,%type:list) [] */
func F_Optimize_c_code_method_method1 (self *ClaireMethod ,l *ClaireList ,_Ztype *ClaireList ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var g0800UU *ClaireClass  
      /* noccur = 1 */
      var g0800UU_try08013 EID 
      g0800UU_try08013 = F_Optimize_c_srange_method(self)
      /* ERROR PROTECTION INSERTED (g0800UU-Result) */
      if ErrorIn(g0800UU_try08013) {Result = g0800UU_try08013
      } else {
      g0800UU = ToClass(OBJ(g0800UU_try08013))
      Result = F_Optimize_c_code_method_method2(self,l,_Ztype,g0800UU)
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_code_method @ list<type_expression>(method, list, list) (throw: true) 
func E_Optimize_c_code_method_method1 (self EID,l EID,_Ztype EID) EID { 
    return /*(sm for c_code_method @ list<type_expression>(method, list, list)= EID)*/ F_Optimize_c_code_method_method1(ToMethod(OBJ(self)),ToList(OBJ(l)),ToList(OBJ(_Ztype)) )} 
  
/* {1} OPT.The go function for: c_code_method(self:method,l:list,%type:list,sx:class) [] */
func F_Optimize_c_code_method_method2 (self *ClaireMethod ,l *ClaireList ,_Ztype *ClaireList ,sx *ClaireClass ) EID { 
    var Result EID 
    if ((self.Module_I.Id() != C_claire.Id()) || 
        ((C_compiler.Safety > 4) || 
          (self.Functional.Id() != CNULL))) /* If:2 */{ 
      /* Let:3 */{ 
        var ld *ClaireList   = self.Domain
        /* noccur = 2 */
        /* Let:4 */{ 
          var n int  = ld.Length()
          /* noccur = 4 */
          if (n != l.Length()) /* If:5 */{ 
            /* Let:6 */{ 
              var g0805UU *ClaireList  
              /* noccur = 1 */
              /* Let:7 */{ 
                var i_bag *ClaireList   = ToType(CEMPTY.Id()).EmptyList()
                /* noccur = 2 */
                /* Let:8 */{ 
                  var i int  = 1
                  /* noccur = 4 */
                  /* Let:9 */{ 
                    var g0802 int  = (n-1)
                    /* noccur = 1 */
                    for (i <= g0802) /* while:10 */{ 
                      i_bag.AddFast(l.At(i-1))
                      i = (i+1)
                      /* while-10 */} 
                    /* Let-9 */} 
                  /* Let-8 */} 
                g0805UU = i_bag
                /* Let-7 */} 
              /* Let:7 */{ 
                var g0806UU *ClaireList  
                /* noccur = 1 */
                /* Let:8 */{ 
                  var i_bag *ClaireList   = ToType(CEMPTY.Id()).EmptyList()
                  /* noccur = 2 */
                  /* Let:9 */{ 
                    var i int  = n
                    /* noccur = 4 */
                    /* Let:10 */{ 
                      var g0803 int  = l.Length()
                      /* noccur = 1 */
                      for (i <= g0803) /* while:11 */{ 
                        i_bag.AddFast(l.At(i-1))
                        i = (i+1)
                        /* while-11 */} 
                      /* Let-10 */} 
                    /* Let-9 */} 
                  g0806UU = i_bag
                  /* Let-8 */} 
                l = g0805UU.AddFast(g0806UU.Id())
                /* Let-7 */} 
              /* Let-6 */} 
            /* If-5 */} 
          var g0807I *ClaireBoolean  
          var g0807I_try08085 EID 
          /* and:5 */{ 
            var v_and5 *ClaireBoolean  
            
            v_and5 = self.Inline_ask
            if (v_and5 == CFALSE) {g0807I_try08085 = EID{CFALSE.Id(),0}
            } else /* arg:6 */{ 
              var v_and5_try08097 EID 
              v_and5_try08097 = F_Optimize_c_inline_ask_method(self,l)
              /* ERROR PROTECTION INSERTED (v_and5-g0807I_try08085) */
              if ErrorIn(v_and5_try08097) {g0807I_try08085 = v_and5_try08097
              } else {
              v_and5 = ToBoolean(OBJ(v_and5_try08097))
              if (v_and5 == CFALSE) {g0807I_try08085 = EID{CFALSE.Id(),0}
              } else /* arg:7 */{ 
                g0807I_try08085 = EID{CTRUE.Id(),0}/* arg-7 */} 
              /* arg-6 */} 
            }
            /* and-5 */} 
          /* ERROR PROTECTION INSERTED (g0807I-Result) */
          if ErrorIn(g0807I_try08085) {Result = g0807I_try08085
          } else {
          g0807I = ToBoolean(OBJ(g0807I_try08085))
          if (g0807I == CTRUE) /* If:5 */{ 
            Result = F_Optimize_c_inline_method1(self,l,sx)
            } else {
            /* Let:6 */{ 
              var g0810UU *ClaireList  
              /* noccur = 1 */
              var g0810UU_try08117 EID 
              /* Let:7 */{ 
                var i_bag *ClaireList   = ToType(CEMPTY.Id()).EmptyList()
                /* noccur = 2 */
                /* Let:8 */{ 
                  var i int  = 1
                  /* noccur = 5 */
                  /* Let:9 */{ 
                    var g0804 int  = n
                    /* noccur = 1 */
                    g0810UU_try08117= EID{CFALSE.Id(),0}
                    for (i <= g0804) /* while:10 */{ 
                      var void_try11 EID 
                      _ = void_try11
                      { 
                      /* Let:11 */{ 
                        var g0812UU *ClaireAny  
                        /* noccur = 1 */
                        var g0812UU_try081312 EID 
                        g0812UU_try081312 = F_Compile_c_strict_code_any(l.At(i-1),F_Compile_psort_any(ld.ValuesO()[i-1]))
                        /* ERROR PROTECTION INSERTED (g0812UU-void_try11) */
                        if ErrorIn(g0812UU_try081312) {void_try11 = g0812UU_try081312
                        } else {
                        g0812UU = ANY(g0812UU_try081312)
                        void_try11 = EID{i_bag.AddFast(g0812UU).Id(),0}
                        }
                        /* Let-11 */} 
                      /* ERROR PROTECTION INSERTED (void_try11-void_try11) */
                      if ErrorIn(void_try11) {g0810UU_try08117 = void_try11
                      break
                      } else {
                      i = (i+1)
                      }
                      /* while-10 */} 
                    }
                    /* Let-9 */} 
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (g0810UU_try08117-g0810UU_try08117) */
                if !ErrorIn(g0810UU_try08117) {
                g0810UU_try08117 = EID{i_bag.Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (g0810UU-Result) */
              if ErrorIn(g0810UU_try08117) {Result = g0810UU_try08117
              } else {
              g0810UU = ToList(OBJ(g0810UU_try08117))
              Result = F_Optimize_Call_method_I_method(self,g0810UU).ToEID()
              }
              /* Let-6 */} 
            /* If-5 */} 
          }
          /* Let-4 */} 
        /* Let-3 */} 
      } else {
      if (C_compiler.Optimize_ask == CTRUE) /* If:3 */{ 
        F_Compile_notice_void()
        Core.F_tformat_string(MakeString("poorly typed update: ~S\n"),3,MakeConstantList(self.Id()))
        /* If-3 */} 
      Result = F_Optimize_open_message_property(self.Selector,l)
      /* If-2 */} 
    return Result} 
  
// The EID go function for: c_code_method @ list<type_expression>(method, list, list, class) (throw: true) 
func E_Optimize_c_code_method_method2 (self EID,l EID,_Ztype EID,sx EID) EID { 
    return /*(sm for c_code_method @ list<type_expression>(method, list, list, class)= EID)*/ F_Optimize_c_code_method_method2(ToMethod(OBJ(self)),
      ToList(OBJ(l)),
      ToList(OBJ(_Ztype)),
      ToClass(OBJ(sx)) )} 
  
// the code to be produced for a method
/* {1} OPT.The go function for: Call_method!(self:method,%code:list) [] */
func F_Optimize_Call_method_I_method (self *ClaireMethod ,_Zcode *ClaireList ) *ClaireAny  { 
    // procedure body with s =  
var Result *ClaireAny  
    if (F_Optimize_legal_ask_module(self.Module_I,self.Id()) != CTRUE) /* If:2 */{ 
      Core.F_tformat_string(MakeString("in call ~S~S\n"),0,MakeConstantList(self.Selector.Id(),_Zcode.Id()))
      /* If-2 */} 
    if (_Zcode.Length() == 1) /* If:2 */{ 
      /* Let:3 */{ 
        var _CL_obj *Language.CallMethod1   = Language.To_CallMethod1(new(Language.CallMethod1).Is(Language.C_Call_method1))
        /* noccur = 5 */
        _CL_obj.Arg = self
        _CL_obj.Args = _Zcode
        Result = _CL_obj.Id()
        /* Let-3 */} 
      /* If!2 */}  else if (_Zcode.Length() == 2) /* If:2 */{ 
      /* Let:3 */{ 
        var _CL_obj *Language.CallMethod2   = Language.To_CallMethod2(new(Language.CallMethod2).Is(Language.C_Call_method2))
        /* noccur = 5 */
        _CL_obj.Arg = self
        _CL_obj.Args = _Zcode
        Result = _CL_obj.Id()
        /* Let-3 */} 
      } else {
      /* Let:3 */{ 
        var _CL_obj *Language.CallMethod   = Language.To_CallMethod(new(Language.CallMethod).Is(Language.C_Call_method))
        /* noccur = 5 */
        _CL_obj.Arg = self
        _CL_obj.Args = _Zcode
        Result = _CL_obj.Id()
        /* Let-3 */} 
      /* If-2 */} 
    return Result} 
  
// The EID go function for: Call_method! @ method (throw: false) 
func E_Optimize_Call_method_I_method (self EID,_Zcode EID) EID { 
    return /*(sm for Call_method! @ method= any)*/ F_Optimize_Call_method_I_method(ToMethod(OBJ(self)),ToList(OBJ(_Zcode)) ).ToEID()} 
  
// a call_method or a call external has an obvious type (we do not need to do
// better ?)
/* {1} OPT.The go function for: c_type(self:Call_method) [] */
func F_c_type_Call_method (self *Language.CallMethod ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var g0814UU *ClaireList  
      /* noccur = 1 */
      var g0814UU_try08153 EID 
      /* Iteration:3 */{ 
        var v_list3 *ClaireList  
        var x *ClaireAny  
        var v_local3 *ClaireAny  
        v_list3 = self.Args
        g0814UU_try08153 = EID{CreateList(ToType(CEMPTY.Id()),v_list3.Length()).Id(),0}
        for CLcount := 0; CLcount < v_list3.Length(); CLcount++{ 
          x = v_list3.At(CLcount)
          var v_local3_try08165 EID 
          v_local3_try08165 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
          /* ERROR PROTECTION INSERTED (v_local3-g0814UU_try08153) */
          if ErrorIn(v_local3_try08165) {g0814UU_try08153 = v_local3_try08165
          g0814UU_try08153 = v_local3_try08165
          break
          } else {
          v_local3 = ANY(v_local3_try08165)
          ToList(OBJ(g0814UU_try08153)).PutAt(CLcount,v_local3)
          } 
        }
        /* Iteration-3 */} 
      /* ERROR PROTECTION INSERTED (g0814UU-Result) */
      if ErrorIn(g0814UU_try08153) {Result = g0814UU_try08153
      } else {
      g0814UU = ToList(OBJ(g0814UU_try08153))
      Result = F_Optimize_use_range_method(self.Arg,g0814UU)
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_type @ Call_method (throw: true) 
func E_c_type_Call_method (self EID) EID { 
    return /*(sm for c_type @ Call_method= EID)*/ F_c_type_Call_method(Language.To_CallMethod(OBJ(self)) )} 
  
// a call_method is already compiled
/* {1} OPT.The go function for: c_code(self:Call_method) [] */
func F_c_code_Call_method (self *Language.CallMethod ) *ClaireAny  { 
    // use function body compiling 
return  self.Id()
    } 
  
// The EID go function for: c_code @ Call_method (throw: false) 
func E_c_code_Call_method (self EID) EID { 
    return /*(sm for c_code @ Call_method= any)*/ F_c_code_Call_method(Language.To_CallMethod(OBJ(self)) ).ToEID()} 
  
// gets the associated function if it exists and create one otherwise
/* {1} OPT.The go function for: Compile/functional!(self:method) [] */
func F_Compile_functional_I_method (self *ClaireMethod ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var f *ClaireAny   = Core.F_get_property(C_functional,ToObject(self.Id()))
      /* noccur = 3 */
      /* Let:3 */{ 
        var p *ClaireProperty   = self.Selector
        /* noccur = 1 */
        if (C_function.Id() == f.Isa.Id()) /* If:4 */{ 
          /* Let:5 */{ 
            var g0817 *ClaireFunction   = ToFunction(f)
            /* noccur = 1 */
            Result = EID{g0817.Id(),0}
            /* Let-5 */} 
          } else {
          /* Let:5 */{ 
            var g0819UU *ClaireAny  
            /* noccur = 1 */
            var g0819UU_try08206 EID 
            g0819UU_try08206 = Core.F_CALL(C_Compile_function_name,ARGS(EID{p.Id(),0},EID{self.Domain.Id(),0},f.ToEID()))
            /* ERROR PROTECTION INSERTED (g0819UU-Result) */
            if ErrorIn(g0819UU_try08206) {Result = g0819UU_try08206
            } else {
            g0819UU = ANY(g0819UU_try08206)
            Result = F_make_function_string(ToString(g0819UU)).ToEID()
            }
            /* Let-5 */} 
          /* If-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: Compile/functional! @ method (throw: true) 
func E_Compile_functional_I_method (self EID) EID { 
    return /*(sm for Compile/functional! @ method= EID)*/ F_Compile_functional_I_method(ToMethod(OBJ(self)) )} 
  
// second-order types for better safety or optimization -------------------------------
/* {1} OPT.The go function for: nth_type_check(tl:type,ti:type,tx:type) [] */
func F_Optimize_nth_type_check_type (tl *ClaireType ,ti *ClaireType ,tx *ClaireType ) *ClaireAny  { 
    // use function body compiling 
if (tx.Included(Core.F_member_type(tl)) != CTRUE) /* If:2 */{ 
      F_Compile_warn_void()
      Core.F_tformat_string(MakeString("unsafe update on bag: type ~S into ~S [252]\n"),2,MakeConstantList(tx.Id(),tl.Id()))
      /* If-2 */} 
    return  tx.Id()
    } 
  
// The EID go function for: nth_type_check @ type (throw: false) 
func E_Optimize_nth_type_check_type (tl EID,ti EID,tx EID) EID { 
    return /*(sm for nth_type_check @ type= any)*/ F_Optimize_nth_type_check_type(ToType(OBJ(tl)),ToType(OBJ(ti)),ToType(OBJ(tx)) ).ToEID()} 
  
// ******************************************************************
// *    Part 5: inline methods                                      *
// ******************************************************************
// macro expansion for inline method ?
// we check that it is a good idea
/* {1} OPT.The go function for: c_inline?(self:method,l:list) [] */
func F_Optimize_c_inline_ask_method (self *ClaireMethod ,l *ClaireList ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var f *ClaireLambda   = self.Formula
      /* noccur = 3 */
      /* Let:3 */{ 
        var la *ClaireList   = f.Vars
        /* noccur = 0 */
        _ = la
        /* Let:4 */{ 
          var x *ClaireAny   = f.Body
          /* noccur = 1 */
          /* Let:5 */{ 
            var n int  = 1
            /* noccur = 3 */
            /* Let:6 */{ 
              var g0821UU *ClaireAny  
              /* noccur = 1 */
              var g0821UU_try08227 EID 
              /* For:7 */{ 
                var v *ClaireAny  
                _ = v
                g0821UU_try08227= EID{CFALSE.Id(),0}
                var v_support *ClaireList  
                v_support = f.Vars
                v_len := v_support.Length()
                for i_it := 0; i_it < v_len; i_it++ { 
                  v = v_support.At(i_it)
                  var void_try9 EID 
                  _ = void_try9
                  var g0823I *ClaireBoolean  
                  var g0823I_try08249 EID 
                  /* and:9 */{ 
                    var v_and9 *ClaireBoolean  
                    
                    v_and9 = Core.F__sup_integer(Language.F_occurrence_any(x,To_Variable(v)),1)
                    if (v_and9 == CFALSE) {g0823I_try08249 = EID{CFALSE.Id(),0}
                    } else /* arg:10 */{ 
                      var v_and9_try082511 EID 
                      /* Let:11 */{ 
                        var g0826UU *ClaireBoolean  
                        /* noccur = 1 */
                        var g0826UU_try082712 EID 
                        g0826UU_try082712 = F_Compile_designated_ask_any(l.At(n-1))
                        /* ERROR PROTECTION INSERTED (g0826UU-v_and9_try082511) */
                        if ErrorIn(g0826UU_try082712) {v_and9_try082511 = g0826UU_try082712
                        } else {
                        g0826UU = ToBoolean(OBJ(g0826UU_try082712))
                        v_and9_try082511 = EID{g0826UU.Not.Id(),0}
                        }
                        /* Let-11 */} 
                      /* ERROR PROTECTION INSERTED (v_and9-g0823I_try08249) */
                      if ErrorIn(v_and9_try082511) {g0823I_try08249 = v_and9_try082511
                      } else {
                      v_and9 = ToBoolean(OBJ(v_and9_try082511))
                      if (v_and9 == CFALSE) {g0823I_try08249 = EID{CFALSE.Id(),0}
                      } else /* arg:11 */{ 
                        v_and9 = Core.F_owner_any(ANY(Core.F_CALL(C_range,ARGS(v.ToEID())))).IsIn(C_Optimize_Pattern).Not
                        if (v_and9 == CFALSE) {g0823I_try08249 = EID{CFALSE.Id(),0}
                        } else /* arg:12 */{ 
                          g0823I_try08249 = EID{CTRUE.Id(),0}/* arg-12 */} 
                        /* arg-11 */} 
                      /* arg-10 */} 
                    }
                    /* and-9 */} 
                  /* ERROR PROTECTION INSERTED (g0823I-void_try9) */
                  if ErrorIn(g0823I_try08249) {void_try9 = g0823I_try08249
                  } else {
                  g0823I = ToBoolean(OBJ(g0823I_try08249))
                  if (g0823I == CTRUE) /* If:9 */{ 
                     /*v = g0821UU_try08227, s =EID*/
g0821UU_try08227 = EID{CTRUE.Id(),0}
                    break
                    } else {
                    n = (n+1)
                    void_try9 = EID{C__INT,IVAL(n)}
                    /* If-9 */} 
                  }
                  /* ERROR PROTECTION INSERTED (void_try9-g0821UU_try08227) */
                  if ErrorIn(void_try9) {g0821UU_try08227 = void_try9
                  g0821UU_try08227 = void_try9
                  break
                  } else {
                  }
                  /* loop-8 */} 
                /* For-7 */} 
              /* ERROR PROTECTION INSERTED (g0821UU-Result) */
              if ErrorIn(g0821UU_try08227) {Result = g0821UU_try08227
              } else {
              g0821UU = ANY(g0821UU_try08227)
              Result = EID{Core.F_not_any(g0821UU).Id(),0}
              }
              /* Let-6 */} 
            /* Let-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: c_inline? @ method (throw: true) 
func E_Optimize_c_inline_ask_method (self EID,l EID) EID { 
    return /*(sm for c_inline? @ method= EID)*/ F_Optimize_c_inline_ask_method(ToMethod(OBJ(self)),ToList(OBJ(l)) )} 
  
// checks if a special optization restriction can be used (with patterns)
/* {1} OPT.The go function for: inline_optimize?(self:Call) [] */
func F_Optimize_inline_optimize_ask_Call (self *Language.Call ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var l *ClaireList   = self.Args
      /* noccur = 2 */
      /* Let:3 */{ 
        var m *ClaireAny  
        /* noccur = 2 */
        /* Let:4 */{ 
          var g0830UU *ClaireList  
          /* noccur = 1 */
          /* Iteration:5 */{ 
            var v_list5 *ClaireList  
            var x *ClaireAny  
            var v_local5 *ClaireAny  
            v_list5 = l
            g0830UU = CreateList(ToType(CEMPTY.Id()),v_list5.Length())
            for CLcount := 0; CLcount < v_list5.Length(); CLcount++{ 
              x = v_list5.At(CLcount)
              v_local5 = MakeConstantSet(x).Id()
              g0830UU.PutAt(CLcount,v_local5)
              } 
            /* Iteration-5 */} 
          m = F_Optimize_restriction_I_property(self.Selector,g0830UU,CTRUE)
          /* Let-4 */} 
        if (C_method.Id() == m.Isa.Id()) /* If:4 */{ 
          /* Let:5 */{ 
            var g0828 *ClaireMethod   = ToMethod(m)
            /* noccur = 4 */
            var g0831I *ClaireBoolean  
            var g0831I_try08326 EID 
            /* and:6 */{ 
              var v_and6 *ClaireBoolean  
              
              v_and6 = g0828.Inline_ask
              if (v_and6 == CFALSE) {g0831I_try08326 = EID{CFALSE.Id(),0}
              } else /* arg:7 */{ 
                /* Let:8 */{ 
                  var g0833UU *ClaireAny  
                  /* noccur = 1 */
                  /* For:9 */{ 
                    var s *ClaireAny  
                    _ = s
                    g0833UU= CFALSE.Id()
                    for _,s = range(g0828.Domain.ValuesO())/* loop:10 */{ 
                      if (s.Isa.IsIn(C_Optimize_Pattern) == CTRUE) /* If:11 */{ 
                         /*v = g0833UU, s =any*/
g0833UU = CTRUE.Id()
                        break
                        /* If-11 */} 
                      /* loop-10 */} 
                    /* For-9 */} 
                  v_and6 = F_boolean_I_any(g0833UU)
                  /* Let-8 */} 
                if (v_and6 == CFALSE) {g0831I_try08326 = EID{CFALSE.Id(),0}
                } else /* arg:8 */{ 
                  var v_and6_try08349 EID 
                  v_and6_try08349 = F_Optimize_c_inline_ask_method(g0828,l)
                  /* ERROR PROTECTION INSERTED (v_and6-g0831I_try08326) */
                  if ErrorIn(v_and6_try08349) {g0831I_try08326 = v_and6_try08349
                  } else {
                  v_and6 = ToBoolean(OBJ(v_and6_try08349))
                  if (v_and6 == CFALSE) {g0831I_try08326 = EID{CFALSE.Id(),0}
                  } else /* arg:9 */{ 
                    g0831I_try08326 = EID{CTRUE.Id(),0}/* arg-9 */} 
                  /* arg-8 */} 
                /* arg-7 */} 
              }
              /* and-6 */} 
            /* ERROR PROTECTION INSERTED (g0831I-Result) */
            if ErrorIn(g0831I_try08326) {Result = g0831I_try08326
            } else {
            g0831I = ToBoolean(OBJ(g0831I_try08326))
            if (g0831I == CTRUE) /* If:6 */{ 
              Result = EID{g0828.Id(),0}
              } else {
              Result = EID{CFALSE.Id(),0}
              /* If-6 */} 
            }
            /* Let-5 */} 
          } else {
          Result = EID{CFALSE.Id(),0}
          /* If-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: inline_optimize? @ Call (throw: true) 
func E_Optimize_inline_optimize_ask_Call (self EID) EID { 
    return /*(sm for inline_optimize? @ Call= EID)*/ F_Optimize_inline_optimize_ask_Call(Language.To_Call(OBJ(self)) )} 
  
// eof