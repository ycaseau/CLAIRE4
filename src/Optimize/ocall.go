/***** CLAIRE Compilation of file /Users/ycaseau/Dropbox/src/clairev4.0/src/compile/ocall.cl 
         [version 4.0.02 / safety 5] Saturday 10-30-2021 *****/

package Optimize
import (_ "fmt"
	. "Kernel"
	"Core"
	"Language"
	"Reader"
)

//-------- dumb function to prevent import errors --------
func import_g0377() { 
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
        var g0378 int  = l.Length()
        /* noccur = 1 */
        for (i <= g0378) /* while:4 */{ 
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
            var _Zt_try03836 EID 
            _Zt_try03836 = Core.F_CALL(C_c_type,ARGS(self.Formula.Body.ToEID()))
            /* ERROR PROTECTION INSERTED (_Zt-Result) */
            if ErrorIn(_Zt_try03836) {Result = _Zt_try03836
            } else {
            _Zt = ToType(OBJ(_Zt_try03836))
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
              var _Zt_try03847 EID 
              _Zt_try03847 = Core.F_CALL(ToProperty(C__exp.Id()),ARGS(EID{_Zt.Id(),0},EID{self.Range.Id(),0}))
              /* ERROR PROTECTION INSERTED (_Zt-Result) */
              if ErrorIn(_Zt_try03847) {Result = _Zt_try03847
              } else {
              _Zt = ToType(OBJ(_Zt_try03847))
              Result = EID{_Zt.Id(),0}
              }
              } else {
              Result = EID{CFALSE.Id(),0}
              /* If-6 */} 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            if (F_boolean_I_any(_Zt.Id()).Id() != CTRUE.Id()) /* If:6 */{ 
              /* Let:7 */{ 
                var g0385UU *ClaireType  
                /* noccur = 1 */
                var g0385UU_try03868 EID 
                g0385UU_try03868 = Core.F_CALL(C_c_type,ARGS(self.Formula.Body.ToEID()))
                /* ERROR PROTECTION INSERTED (g0385UU-Result) */
                if ErrorIn(g0385UU_try03868) {Result = g0385UU_try03868
                } else {
                g0385UU = ToType(OBJ(g0385UU_try03868))
                Result = Core.F_CALL(C_Compile_Cerror,ARGS(EID{MakeString("[207] inline ~S: range ~S is incompatible with ~S (inferred)").Id(),0},
                  EID{self.Id(),0},
                  EID{self.Range.Id(),0},
                  EID{g0385UU.Id(),0}))
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
                  var g0379 *ClaireLambda   = ToLambda(f)
                  /* noccur = 1 */
                  _Zt2_try7 = Language.F_apply_lambda(g0379,_Zl2)
                  /* Let-8 */} 
                /* If!7 */}  else if (f.Isa.IsIn(C_property) == CTRUE) /* If:7 */{ 
                /* Let:8 */{ 
                  var g0380 *ClaireProperty   = ToProperty(f)
                  /* noccur = 1 */
                  _Zt2_try7 = Core.F_apply_property(g0380,_Zl2)
                  /* Let-8 */} 
                /* If!7 */}  else if (C_function.Id() == f.Isa.Id()) /* If:7 */{ 
                /* Let:8 */{ 
                  var g0381 *ClaireFunction   = ToFunction(f)
                  /* noccur = 1 */
                  _Zt2_try7 = F_apply_function(g0381,_Zl2)
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
            var _Ztype_try03946 EID 
            /* Iteration:6 */{ 
              var v_list6 *ClaireList  
              var x *ClaireAny  
              var v_local6 *ClaireAny  
              v_list6 = l
              _Ztype_try03946 = EID{CreateList(ToType(CEMPTY.Id()),v_list6.Length()).Id(),0}
              for CLcount := 0; CLcount < v_list6.Length(); CLcount++{ 
                x = v_list6.At(CLcount)
                var v_local6_try03958 EID 
                v_local6_try03958 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
                /* ERROR PROTECTION INSERTED (v_local6-_Ztype_try03946) */
                if ErrorIn(v_local6_try03958) {_Ztype_try03946 = v_local6_try03958
                _Ztype_try03946 = v_local6_try03958
                break
                } else {
                v_local6 = ANY(v_local6_try03958)
                ToList(OBJ(_Ztype_try03946)).PutAt(CLcount,v_local6)
                } 
              }
              /* Iteration-6 */} 
            /* ERROR PROTECTION INSERTED (_Ztype-Result) */
            if ErrorIn(_Ztype_try03946) {Result = _Ztype_try03946
            } else {
            _Ztype = ToList(OBJ(_Ztype_try03946))
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
                    var g0387 *ClaireProperty   = ToProperty(r.Id())
                    /* noccur = 2 */
                    /* Let:10 */{ 
                      var xs *ClaireObject   = Core.F__at_property1(g0387,ToTypeExpression(_Ztype.At(2-1)).Class_I())
                      /* noccur = 2 */
                      if (C_slot.Id() == xs.Isa.Id()) /* If:11 */{ 
                        /* Let:12 */{ 
                          var g0388 *ClaireSlot   = ToSlot(xs.Id())
                          /* noccur = 5 */
                          if ((g0388.Range.Included(ToType(C_set.Id())) == CTRUE) && 
                              (C_compiler.Safety < 3)) /* If:13 */{ 
                            Result = EID{ToTypeExpression(OBJ(Core.F_CALL(C_range,ARGS(l.At(2-1).ToEID())))).Class_I().Id(),0}
                            /* If!13 */}  else if (g0388.Range.Contains(g0388.Default) == CTRUE) /* If:13 */{ 
                            Result = EID{g0388.Range.Id(),0}
                            } else {
                            Result = EID{F_Optimize_extends_type(g0388.Range).Id(),0}
                            /* If-13 */} 
                          /* Let-12 */} 
                        } else {
                        Result = EID{g0387.Range.Id(),0}
                        /* If-11 */} 
                      /* Let-10 */} 
                    /* Let-9 */} 
                  /* If!8 */}  else if (C_table.Id() == r.Isa.Id()) /* If:8 */{ 
                  /* Let:9 */{ 
                    var g0390 *ClaireTable   = ToTable(r.Id())
                    /* noccur = 4 */
                    if (g0390.Range.Contains(g0390.Default) == CTRUE) /* If:10 */{ 
                      Result = EID{g0390.Range.Id(),0}
                      } else {
                      Result = EID{F_Optimize_extends_type(g0390.Range).Id(),0}
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
                    var g0391 *ClaireSlot   = ToSlot(r)
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
                      Result = EID{g0391.Range.Id(),0}
                      /* If-10 */} 
                    /* Let-9 */} 
                  /* If!8 */}  else if (C_method.Id() == r.Isa.Id()) /* If:8 */{ 
                  /* Let:9 */{ 
                    var g0392 *ClaireMethod   = ToMethod(r)
                    /* noccur = 1 */
                    Result = F_Optimize_use_range_method(g0392,_Ztype)
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
        var g0400I *ClaireBoolean  
        var g0400I_try04014 EID 
        /* and:4 */{ 
          var v_and4 *ClaireBoolean  
          
          v_and4 = l.At(1-1).Isa.IsIn(Core.C_global_variable)
          if (v_and4 == CFALSE) {g0400I_try04014 = EID{CFALSE.Id(),0}
          } else /* arg:5 */{ 
            v_and4 = Equal(ANY(Core.F_CALL(C_range,ARGS(l.At(1-1).ToEID()))),CEMPTY.Id())
            if (v_and4 == CFALSE) {g0400I_try04014 = EID{CFALSE.Id(),0}
            } else /* arg:6 */{ 
              var v_and4_try04027 EID 
              v_and4_try04027 = F_Compile_designated_ask_any(ANY(Core.F_CALL(C_value,ARGS(l.At(1-1).ToEID()))))
              /* ERROR PROTECTION INSERTED (v_and4-g0400I_try04014) */
              if ErrorIn(v_and4_try04027) {g0400I_try04014 = v_and4_try04027
              } else {
              v_and4 = ToBoolean(OBJ(v_and4_try04027))
              if (v_and4 == CFALSE) {g0400I_try04014 = EID{CFALSE.Id(),0}
              } else /* arg:7 */{ 
                g0400I_try04014 = EID{CTRUE.Id(),0}/* arg-7 */} 
              /* arg-6 */} 
            /* arg-5 */} 
          }
          /* and-4 */} 
        /* ERROR PROTECTION INSERTED (g0400I-Result) */
        if ErrorIn(g0400I_try04014) {Result = g0400I_try04014
        } else {
        g0400I = ToBoolean(OBJ(g0400I_try04014))
        if (g0400I == CTRUE) /* If:4 */{ 
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
          var m_try04035 EID 
          m_try04035 = F_Optimize_inline_optimize_ask_Call(self)
          /* ERROR PROTECTION INSERTED (m-Result) */
          if ErrorIn(m_try04035) {Result = m_try04035
          } else {
          m = ANY(m_try04035)
          /* Let:5 */{ 
            var b *ClaireBoolean   = l.At(1-1).Isa.IsIn(C_property)
            /* noccur = 8 */
            /* Let:6 */{ 
              var d *ClaireAny  
              /* noccur = 2 */
              var d_try04047 EID 
              d_try04047 = F_Optimize_daccess_any(self.Id(),Core.F__sup_integer(C_compiler.Safety,5))
              /* ERROR PROTECTION INSERTED (d-Result) */
              if ErrorIn(d_try04047) {Result = d_try04047
              } else {
              d = ANY(d_try04047)
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
                  var g0405UU *ClaireAny  
                  /* noccur = 1 */
                  var g0405UU_try04069 EID 
                  g0405UU_try04069 = F_Optimize_Produce_erase_property(ToProperty(l.At(1-1)),To_Variable(l.At(2-1)))
                  /* ERROR PROTECTION INSERTED (g0405UU-Result) */
                  if ErrorIn(g0405UU_try04069) {Result = g0405UU_try04069
                  } else {
                  g0405UU = ANY(g0405UU_try04069)
                  Result = Core.F_CALL(C_c_code,ARGS(g0405UU.ToEID(),EID{sx.Id(),0}))
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
                      var x_try040711 EID 
                      /* Let:11 */{ 
                        var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                        /* noccur = 5 */
                        _CL_obj.Selector = C_safe
                        /* update:12 */{ 
                          var va_arg1 *Language.Call  
                          var va_arg2 *ClaireList  
                          va_arg1 = _CL_obj
                          var va_arg2_try040813 EID 
                          /* Construct:13 */{ 
                            var v_bag_arg *ClaireAny  
                            va_arg2_try040813= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                            var v_bag_arg_try040914 EID 
                            v_bag_arg_try040914 = Core.F_CALL(C_c_code,ARGS(l.At(1-1).ToEID(),EID{sx.Id(),0}))
                            /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try040813) */
                            if ErrorIn(v_bag_arg_try040914) {va_arg2_try040813 = v_bag_arg_try040914
                            } else {
                            v_bag_arg = ANY(v_bag_arg_try040914)
                            ToList(OBJ(va_arg2_try040813)).AddFast(v_bag_arg)}
                            /* Construct-13 */} 
                          /* ERROR PROTECTION INSERTED (va_arg2-x_try040711) */
                          if ErrorIn(va_arg2_try040813) {x_try040711 = va_arg2_try040813
                          } else {
                          va_arg2 = ToList(OBJ(va_arg2_try040813))
                          /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                          va_arg1.Args = va_arg2
                          x_try040711 = EID{va_arg2.Id(),0}
                          }
                          /* update-12 */} 
                        /* ERROR PROTECTION INSERTED (x_try040711-x_try040711) */
                        if !ErrorIn(x_try040711) {
                        x_try040711 = EID{_CL_obj.Id(),0}
                        }
                        /* Let-11 */} 
                      /* ERROR PROTECTION INSERTED (x-Result) */
                      if ErrorIn(x_try040711) {Result = x_try040711
                      } else {
                      x = ANY(x_try040711)
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
                var g0410I *ClaireBoolean  
                var g0410I_try04118 EID 
                /* and:8 */{ 
                  var v_and8 *ClaireBoolean  
                  
                  v_and8 = MakeBoolean((s.Id() == C_add.Id()) || (s.Id() == C_add_I.Id()))
                  if (v_and8 == CFALSE) {g0410I_try04118 = EID{CFALSE.Id(),0}
                  } else /* arg:9 */{ 
                    var v_and8_try041210 EID 
                    /* Let:10 */{ 
                      var g0413UU *ClaireType  
                      /* noccur = 1 */
                      var g0413UU_try041411 EID 
                      g0413UU_try041411 = Core.F_CALL(C_c_type,ARGS(l.At(1-1).ToEID()))
                      /* ERROR PROTECTION INSERTED (g0413UU-v_and8_try041210) */
                      if ErrorIn(g0413UU_try041411) {v_and8_try041210 = g0413UU_try041411
                      } else {
                      g0413UU = ToType(OBJ(g0413UU_try041411))
                      v_and8_try041210 = EID{g0413UU.Included(ToType(C_bag.Id())).Id(),0}
                      }
                      /* Let-10 */} 
                    /* ERROR PROTECTION INSERTED (v_and8-g0410I_try04118) */
                    if ErrorIn(v_and8_try041210) {g0410I_try04118 = v_and8_try041210
                    } else {
                    v_and8 = ToBoolean(OBJ(v_and8_try041210))
                    if (v_and8 == CFALSE) {g0410I_try04118 = EID{CFALSE.Id(),0}
                    } else /* arg:10 */{ 
                      g0410I_try04118 = EID{CTRUE.Id(),0}/* arg-10 */} 
                    /* arg-9 */} 
                  }
                  /* and-8 */} 
                /* ERROR PROTECTION INSERTED (g0410I-Result) */
                if ErrorIn(g0410I_try04118) {Result = g0410I_try04118
                } else {
                g0410I = ToBoolean(OBJ(g0410I_try04118))
                if (g0410I == CTRUE) /* If:8 */{ 
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
                    var g0415UU *ClaireClass  
                    /* noccur = 1 */
                    var g0415UU_try041610 EID 
                    g0415UU_try041610 = F_Optimize_c_srange_method(ToMethod(m))
                    /* ERROR PROTECTION INSERTED (g0415UU-Result) */
                    if ErrorIn(g0415UU_try041610) {Result = g0415UU_try041610
                    } else {
                    g0415UU = ToClass(OBJ(g0415UU_try041610))
                    Result = F_Optimize_c_inline_method1(ToMethod(m),l,g0415UU)
                    }
                    /* Let-9 */} 
                  } else {
                  var g0417I *ClaireBoolean  
                  var g0417I_try04189 EID 
                  /* and:9 */{ 
                    var v_and9 *ClaireBoolean  
                    
                    v_and9 = MakeBoolean((s.Id() == C__equal.Id()) || (s.Id() == Core.C__I_equal.Id()))
                    if (v_and9 == CFALSE) {g0417I_try04189 = EID{CFALSE.Id(),0}
                    } else /* arg:10 */{ 
                      var v_and9_try041911 EID 
                      /* Let:11 */{ 
                        var g0420UU *ClaireAny  
                        /* noccur = 1 */
                        var g0420UU_try042112 EID 
                        g0420UU_try042112 = F_Optimize_daccess_any(l.At(1-1),CTRUE)
                        /* ERROR PROTECTION INSERTED (g0420UU-v_and9_try041911) */
                        if ErrorIn(g0420UU_try042112) {v_and9_try041911 = g0420UU_try042112
                        } else {
                        g0420UU = ANY(g0420UU_try042112)
                        v_and9_try041911 = EID{Core.F_known_ask_any(g0420UU).Id(),0}
                        }
                        /* Let-11 */} 
                      /* ERROR PROTECTION INSERTED (v_and9-g0417I_try04189) */
                      if ErrorIn(v_and9_try041911) {g0417I_try04189 = v_and9_try041911
                      } else {
                      v_and9 = ToBoolean(OBJ(v_and9_try041911))
                      if (v_and9 == CFALSE) {g0417I_try04189 = EID{CFALSE.Id(),0}
                      } else /* arg:11 */{ 
                        g0417I_try04189 = EID{CTRUE.Id(),0}/* arg-11 */} 
                      /* arg-10 */} 
                    }
                    /* and-9 */} 
                  /* ERROR PROTECTION INSERTED (g0417I-Result) */
                  if ErrorIn(g0417I_try04189) {Result = g0417I_try04189
                  } else {
                  g0417I = ToBoolean(OBJ(g0417I_try04189))
                  if (g0417I == CTRUE) /* If:9 */{ 
                    Result = F_Optimize_c_code_hold_property(ToProperty(ToList(OBJ(Core.F_CALL(C_args,ARGS(l.At(1-1).ToEID())))).At(1-1)),ToList(OBJ(Core.F_CALL(C_args,ARGS(l.At(1-1).ToEID())))).At(2-1),l.At(2-1),Equal(s.Id(),C__equal.Id()))
                    } else {
                    var g0422I *ClaireBoolean  
                    var g0422I_try042310 EID 
                    /* and:10 */{ 
                      var v_and10 *ClaireBoolean  
                      
                      v_and10 = MakeBoolean((s.Id() == C__equal.Id()) || (s.Id() == Core.C__I_equal.Id()))
                      if (v_and10 == CFALSE) {g0422I_try042310 = EID{CFALSE.Id(),0}
                      } else /* arg:11 */{ 
                        var v_and10_try042412 EID 
                        /* Let:12 */{ 
                          var g0425UU *ClaireAny  
                          /* noccur = 1 */
                          var g0425UU_try042613 EID 
                          g0425UU_try042613 = F_Optimize_daccess_any(l.At(2-1),CTRUE)
                          /* ERROR PROTECTION INSERTED (g0425UU-v_and10_try042412) */
                          if ErrorIn(g0425UU_try042613) {v_and10_try042412 = g0425UU_try042613
                          } else {
                          g0425UU = ANY(g0425UU_try042613)
                          v_and10_try042412 = EID{Core.F_known_ask_any(g0425UU).Id(),0}
                          }
                          /* Let-12 */} 
                        /* ERROR PROTECTION INSERTED (v_and10-g0422I_try042310) */
                        if ErrorIn(v_and10_try042412) {g0422I_try042310 = v_and10_try042412
                        } else {
                        v_and10 = ToBoolean(OBJ(v_and10_try042412))
                        if (v_and10 == CFALSE) {g0422I_try042310 = EID{CFALSE.Id(),0}
                        } else /* arg:12 */{ 
                          g0422I_try042310 = EID{CTRUE.Id(),0}/* arg-12 */} 
                        /* arg-11 */} 
                      }
                      /* and-10 */} 
                    /* ERROR PROTECTION INSERTED (g0422I-Result) */
                    if ErrorIn(g0422I_try042310) {Result = g0422I_try042310
                    } else {
                    g0422I = ToBoolean(OBJ(g0422I_try042310))
                    if (g0422I == CTRUE) /* If:10 */{ 
                      Result = F_Optimize_c_code_hold_property(ToProperty(ToList(OBJ(Core.F_CALL(C_args,ARGS(l.At(2-1).ToEID())))).At(1-1)),ToList(OBJ(Core.F_CALL(C_args,ARGS(l.At(2-1).ToEID())))).At(2-1),l.At(1-1),Equal(s.Id(),C__equal.Id()))
                      /* If!10 */}  else if (((s.Id() == C_put.Id()) || 
                          (s.Id() == C_nth_equal.Id())) && 
                        ((C_table.Id() == l.At(1-1).Isa.Id()) && 
                          (l.Length() == 3))) /* If:10 */{ 
                      Result = F_Optimize_c_code_table_Call(self)
                      } else {
                      var g0427I *ClaireBoolean  
                      var g0427I_try042811 EID 
                      /* and:11 */{ 
                        var v_and11 *ClaireBoolean  
                        
                        v_and11 = MakeBoolean((s.Id() == C_nth_put.Id()) || (s.Id() == C_nth_equal.Id()))
                        if (v_and11 == CFALSE) {g0427I_try042811 = EID{CFALSE.Id(),0}
                        } else /* arg:12 */{ 
                          var v_and11_try042913 EID 
                          /* Let:13 */{ 
                            var g0430UU *ClaireType  
                            /* noccur = 1 */
                            var g0430UU_try043114 EID 
                            g0430UU_try043114 = Core.F_CALL(C_c_type,ARGS(l.At(1-1).ToEID()))
                            /* ERROR PROTECTION INSERTED (g0430UU-v_and11_try042913) */
                            if ErrorIn(g0430UU_try043114) {v_and11_try042913 = g0430UU_try043114
                            } else {
                            g0430UU = ToType(OBJ(g0430UU_try043114))
                            v_and11_try042913 = EID{g0430UU.Included(ToType(C_array.Id())).Id(),0}
                            }
                            /* Let-13 */} 
                          /* ERROR PROTECTION INSERTED (v_and11-g0427I_try042811) */
                          if ErrorIn(v_and11_try042913) {g0427I_try042811 = v_and11_try042913
                          } else {
                          v_and11 = ToBoolean(OBJ(v_and11_try042913))
                          if (v_and11 == CFALSE) {g0427I_try042811 = EID{CFALSE.Id(),0}
                          } else /* arg:13 */{ 
                            v_and11 = Equal(MakeInteger(l.Length()).Id(),MakeInteger(3).Id())
                            if (v_and11 == CFALSE) {g0427I_try042811 = EID{CFALSE.Id(),0}
                            } else /* arg:14 */{ 
                              g0427I_try042811 = EID{CTRUE.Id(),0}/* arg-14 */} 
                            /* arg-13 */} 
                          /* arg-12 */} 
                        }
                        /* and-11 */} 
                      /* ERROR PROTECTION INSERTED (g0427I-Result) */
                      if ErrorIn(g0427I_try042811) {Result = g0427I_try042811
                      } else {
                      g0427I = ToBoolean(OBJ(g0427I_try042811))
                      if (g0427I == CTRUE) /* If:11 */{ 
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
                          var g0432UU *ClaireAny  
                          /* noccur = 1 */
                          var g0432UU_try043313 EID 
                          g0432UU_try043313 = EVAL(l.At(1-1))
                          /* ERROR PROTECTION INSERTED (g0432UU-Result) */
                          if ErrorIn(g0432UU_try043313) {Result = g0432UU_try043313
                          } else {
                          g0432UU = ANY(g0432UU_try043313)
                          Result = Core.F_CALL(C_c_code,ARGS(g0432UU.ToEID()))
                          }
                          /* Let-12 */} 
                        /* If!11 */}  else if (s.Id() == Language.C_function_I.Id()) /* If:11 */{ 
                        /* Let:12 */{ 
                          var g0434UU *ClaireString  
                          /* noccur = 1 */
                          var g0434UU_try043513 EID 
                          /* Let:13 */{ 
                            var g0436UU *ClaireSymbol  
                            /* noccur = 1 */
                            var g0436UU_try043714 EID 
                            g0436UU_try043714 = Language.F_extract_symbol_any(l.At(1-1))
                            /* ERROR PROTECTION INSERTED (g0436UU-g0434UU_try043513) */
                            if ErrorIn(g0436UU_try043714) {g0434UU_try043513 = g0436UU_try043714
                            } else {
                            g0436UU = ToSymbol(OBJ(g0436UU_try043714))
                            g0434UU_try043513 = EID{g0436UU.String_I().Id(),0}
                            }
                            /* Let-13 */} 
                          /* ERROR PROTECTION INSERTED (g0434UU-Result) */
                          if ErrorIn(g0434UU_try043513) {Result = g0434UU_try043513
                          } else {
                          g0434UU = ToString(OBJ(g0434UU_try043513))
                          Result = F_make_function_string(g0434UU).ToEID()
                          }
                          /* Let-12 */} 
                        /* If!11 */}  else if ((s.Id() == Core.C_not.Id()) && 
                          (l.At(1-1).Isa.IsIn(Language.C_Select) == CTRUE)) /* If:11 */{ 
                        Result = F_Optimize_c_code_not_Select(Language.To_Select(l.At(1-1)))
                        /* If!11 */}  else if ((s.Id() == Core.C_call.Id()) && 
                          (l.At(1-1).Isa.IsIn(C_property) == CTRUE)) /* If:11 */{ 
                        /* Let:12 */{ 
                          var g0438UU *Language.Call  
                          /* noccur = 1 */
                          var g0438UU_try043913 EID 
                          /* Let:13 */{ 
                            var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                            /* noccur = 5 */
                            _CL_obj.Selector = ToProperty(l.At(1-1))
                            /* update:14 */{ 
                              var va_arg1 *Language.Call  
                              var va_arg2 *ClaireList  
                              va_arg1 = _CL_obj
                              var va_arg2_try044015 EID 
                              va_arg2_try044015 = l.Cdr()
                              /* ERROR PROTECTION INSERTED (va_arg2-g0438UU_try043913) */
                              if ErrorIn(va_arg2_try044015) {g0438UU_try043913 = va_arg2_try044015
                              } else {
                              va_arg2 = ToList(OBJ(va_arg2_try044015))
                              /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                              va_arg1.Args = va_arg2
                              g0438UU_try043913 = EID{va_arg2.Id(),0}
                              }
                              /* update-14 */} 
                            /* ERROR PROTECTION INSERTED (g0438UU_try043913-g0438UU_try043913) */
                            if !ErrorIn(g0438UU_try043913) {
                            g0438UU_try043913 = EID{_CL_obj.Id(),0}
                            }
                            /* Let-13 */} 
                          /* ERROR PROTECTION INSERTED (g0438UU-Result) */
                          if ErrorIn(g0438UU_try043913) {Result = g0438UU_try043913
                          } else {
                          g0438UU = Language.To_Call(OBJ(g0438UU_try043913))
                          Result = Core.F_CALL(C_c_code,ARGS(EID{g0438UU.Id(),0}))
                          }
                          /* Let-12 */} 
                        /* If!11 */}  else if (s.Open == 3) /* If:11 */{ 
                        /* Let:12 */{ 
                          var g0441UU *ClaireList  
                          /* noccur = 1 */
                          var g0441UU_try044213 EID 
                          /* Iteration:13 */{ 
                            var v_list13 *ClaireList  
                            var x *ClaireAny  
                            var v_local13 *ClaireAny  
                            v_list13 = l
                            g0441UU_try044213 = EID{CreateList(ToType(CEMPTY.Id()),v_list13.Length()).Id(),0}
                            for CLcount := 0; CLcount < v_list13.Length(); CLcount++{ 
                              x = v_list13.At(CLcount)
                              var v_local13_try044315 EID 
                              v_local13_try044315 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
                              /* ERROR PROTECTION INSERTED (v_local13-g0441UU_try044213) */
                              if ErrorIn(v_local13_try044315) {g0441UU_try044213 = v_local13_try044315
                              g0441UU_try044213 = v_local13_try044315
                              break
                              } else {
                              v_local13 = ANY(v_local13_try044315)
                              ToList(OBJ(g0441UU_try044213)).PutAt(CLcount,v_local13)
                              } 
                            }
                            /* Iteration-13 */} 
                          /* ERROR PROTECTION INSERTED (g0441UU-Result) */
                          if ErrorIn(g0441UU_try044213) {Result = g0441UU_try044213
                          } else {
                          g0441UU = ToList(OBJ(g0441UU_try044213))
                          Result = F_Optimize_c_warn_property(s,l,g0441UU)
                          }
                          /* Let-12 */} 
                        } else {
                        var g0444I *ClaireBoolean  
                        /* and:12 */{ 
                          var v_and12 *ClaireBoolean  
                          
                          v_and12 = Equal(s.Id(),Language.C_bit_vector.Id())
                          if (v_and12 == CFALSE) {g0444I = CFALSE
                          } else /* arg:13 */{ 
                            /* Let:14 */{ 
                              var g0445UU *ClaireAny  
                              /* noccur = 1 */
                              /* For:15 */{ 
                                var y *ClaireAny  
                                _ = y
                                g0445UU= CFALSE.Id()
                                var y_support *ClaireList  
                                y_support = self.Args
                                y_len := y_support.Length()
                                for i_it := 0; i_it < y_len; i_it++ { 
                                  y = y_support.At(i_it)
                                  if (C_integer.Id() != y.Isa.Id()) /* If:17 */{ 
                                     /*v = g0445UU, s =any*/
g0445UU = CTRUE.Id()
                                    break
                                    /* If-17 */} 
                                  /* loop-16 */} 
                                /* For-15 */} 
                              v_and12 = Core.F_not_any(g0445UU)
                              /* Let-14 */} 
                            if (v_and12 == CFALSE) {g0444I = CFALSE
                            } else /* arg:14 */{ 
                              g0444I = CTRUE/* arg-14 */} 
                            /* arg-13 */} 
                          /* and-12 */} 
                        if (g0444I == CTRUE) /* If:12 */{ 
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
                            var _Ztype_try044614 EID 
                            /* Iteration:14 */{ 
                              var v_list14 *ClaireList  
                              var x *ClaireAny  
                              var v_local14 *ClaireAny  
                              v_list14 = l
                              _Ztype_try044614 = EID{CreateList(ToType(CEMPTY.Id()),v_list14.Length()).Id(),0}
                              for CLcount := 0; CLcount < v_list14.Length(); CLcount++{ 
                                x = v_list14.At(CLcount)
                                var v_local14_try044716 EID 
                                v_local14_try044716 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
                                /* ERROR PROTECTION INSERTED (v_local14-_Ztype_try044614) */
                                if ErrorIn(v_local14_try044716) {_Ztype_try044614 = v_local14_try044716
                                _Ztype_try044614 = v_local14_try044716
                                break
                                } else {
                                v_local14 = ANY(v_local14_try044716)
                                ToList(OBJ(_Ztype_try044614)).PutAt(CLcount,v_local14)
                                } 
                              }
                              /* Iteration-14 */} 
                            /* ERROR PROTECTION INSERTED (_Ztype-Result) */
                            if ErrorIn(_Ztype_try044614) {Result = _Ztype_try044614
                            } else {
                            _Ztype = ToList(OBJ(_Ztype_try044614))
                            /* Let:14 */{ 
                              var z *ClaireAny   = F_Optimize_restriction_I_property(s,_Ztype,CTRUE)
                              /* noccur = 5 */
                              if (C_slot.Id() == z.Isa.Id()) /* If:15 */{ 
                                /* Let:16 */{ 
                                  var g0396 *ClaireSlot   = ToSlot(z)
                                  /* noccur = 4 */
                                  /* Let:17 */{ 
                                    var _Zunknown *ClaireBoolean   = MakeBoolean((g0396.Range.Contains(g0396.Default) != CTRUE) && (C_OPT.Knowns.Contain_ask(s.Id()) != CTRUE) && (C_compiler.Safety < 5))
                                    /* noccur = 2 */
                                    var g0448I *ClaireBoolean  
                                    var g0448I_try044918 EID 
                                    /* or:18 */{ 
                                      var v_or18 *ClaireBoolean  
                                      
                                      v_or18 = _Zunknown.Not
                                      if (v_or18 == CTRUE) {g0448I_try044918 = EID{CTRUE.Id(),0}
                                      } else /* or:19 */{ 
                                        var v_or18_try045020 EID 
                                        v_or18_try045020 = F_Compile_designated_ask_any(l.At(1-1))
                                        /* ERROR PROTECTION INSERTED (v_or18-g0448I_try044918) */
                                        if ErrorIn(v_or18_try045020) {g0448I_try044918 = v_or18_try045020
                                        } else {
                                        v_or18 = ToBoolean(OBJ(v_or18_try045020))
                                        if (v_or18 == CTRUE) {g0448I_try044918 = EID{CTRUE.Id(),0}
                                        } else /* or:20 */{ 
                                          g0448I_try044918 = EID{CFALSE.Id(),0}/* org-20 */} 
                                        /* org-19 */} 
                                      }
                                      /* or-18 */} 
                                    /* ERROR PROTECTION INSERTED (g0448I-Result) */
                                    if ErrorIn(g0448I_try044918) {Result = g0448I_try044918
                                    } else {
                                    g0448I = ToBoolean(OBJ(g0448I_try044918))
                                    if (g0448I == CTRUE) /* If:18 */{ 
                                      /* Let:19 */{ 
                                        var _CL_obj *Language.CallSlot   = Language.To_CallSlot(new(Language.CallSlot).Is(Language.C_Call_slot))
                                        /* noccur = 7 */
                                        _CL_obj.Selector = g0396
                                        /* update:20 */{ 
                                          var va_arg1 *Language.CallSlot  
                                          var va_arg2 *ClaireAny  
                                          va_arg1 = _CL_obj
                                          var va_arg2_try045121 EID 
                                          va_arg2_try045121 = Core.F_CALL(C_c_code,ARGS(l.At(1-1).ToEID(),EID{F_Compile_psort_any(Core.F_domain_I_restriction(ToRestriction(g0396.Id())).Id()).Id(),0}))
                                          /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                                          if ErrorIn(va_arg2_try045121) {Result = va_arg2_try045121
                                          } else {
                                          va_arg2 = ANY(va_arg2_try045121)
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
                                      
                                      if (C_compiler.Optimize_ask == CTRUE) /* If:19 */{ 
                                        F_Compile_notice_void()
                                        
                                        /* If-19 */} 
                                      Result = F_Optimize_c_warn_property(s,l,_Ztype)
                                      /* If-18 */} 
                                    }
                                    /* Let-17 */} 
                                  /* Let-16 */} 
                                /* If!15 */}  else if (C_method.Id() == z.Isa.Id()) /* If:15 */{ 
                                /* Let:16 */{ 
                                  var g0397 *ClaireMethod   = ToMethod(z)
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
                                  var g0452I *ClaireBoolean  
                                  var g0452I_try045317 EID 
                                  /* or:17 */{ 
                                    var v_or17 *ClaireBoolean  
                                    
                                    var v_or17_try045418 EID 
                                    /* Let:18 */{ 
                                      var g0455UU *ClaireAny  
                                      /* noccur = 1 */
                                      var g0455UU_try045619 EID 
                                      g0455UU_try045619 = Core.F_last_list(g0397.Domain)
                                      /* ERROR PROTECTION INSERTED (g0455UU-v_or17_try045418) */
                                      if ErrorIn(g0455UU_try045619) {v_or17_try045418 = g0455UU_try045619
                                      } else {
                                      g0455UU = ANY(g0455UU_try045619)
                                      v_or17_try045418 = EID{Equal(g0455UU,C_listargs.Id()).Id(),0}
                                      }
                                      /* Let-18 */} 
                                    /* ERROR PROTECTION INSERTED (v_or17-g0452I_try045317) */
                                    if ErrorIn(v_or17_try045418) {g0452I_try045317 = v_or17_try045418
                                    } else {
                                    v_or17 = ToBoolean(OBJ(v_or17_try045418))
                                    if (v_or17 == CTRUE) {g0452I_try045317 = EID{CTRUE.Id(),0}
                                    } else /* or:18 */{ 
                                      v_or17 = MakeBoolean((g0397.Domain.ValuesO()[1-1] == C_void.Id()) && (l.At(1-1) != ClEnv.Id()))
                                      if (v_or17 == CTRUE) {g0452I_try045317 = EID{CTRUE.Id(),0}
                                      } else /* or:19 */{ 
                                        g0452I_try045317 = EID{CFALSE.Id(),0}/* org-19 */} 
                                      /* org-18 */} 
                                    }
                                    /* or-17 */} 
                                  /* ERROR PROTECTION INSERTED (g0452I-Result) */
                                  if ErrorIn(g0452I_try045317) {Result = g0452I_try045317
                                  } else {
                                  g0452I = ToBoolean(OBJ(g0452I_try045317))
                                  if (g0452I == CTRUE) /* If:17 */{ 
                                    Result = F_Optimize_open_message_property(s,l)
                                    } else {
                                    Result = F_Optimize_c_code_method_method2(g0397,l,_Ztype,sx)
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
      var _Zarg_try04573 EID 
      /* Iteration:3 */{ 
        var v_list3 *ClaireList  
        var x *ClaireAny  
        var v_local3 *ClaireAny  
        v_list3 = l
        _Zarg_try04573 = EID{CreateList(ToType(CEMPTY.Id()),v_list3.Length()).Id(),0}
        for CLcount := 0; CLcount < v_list3.Length(); CLcount++{ 
          x = v_list3.At(CLcount)
          var v_local3_try04585 EID 
          var g0459I *ClaireBoolean  
          var g0459I_try04605 EID 
          /* Let:5 */{ 
            var g0461UU *ClaireType  
            /* noccur = 1 */
            var g0461UU_try04626 EID 
            g0461UU_try04626 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
            /* ERROR PROTECTION INSERTED (g0461UU-g0459I_try04605) */
            if ErrorIn(g0461UU_try04626) {g0459I_try04605 = g0461UU_try04626
            } else {
            g0461UU = ToType(OBJ(g0461UU_try04626))
            g0459I_try04605 = EID{Core.F__I_equal_any(g0461UU.Id(),C_void.Id()).Id(),0}
            }
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (g0459I-v_local3_try04585) */
          if ErrorIn(g0459I_try04605) {v_local3_try04585 = g0459I_try04605
          } else {
          g0459I = ToBoolean(OBJ(g0459I_try04605))
          if (g0459I == CTRUE) /* If:5 */{ 
            v_local3_try04585 = Core.F_CALL(C_c_code,ARGS(x.ToEID(),EID{C_any.Id(),0}))
            } else {
            v_local3_try04585 = Core.F_CALL(C_Compile_Cerror,ARGS(EID{MakeString("[206] use of void ~S in ~S~S").Id(),0},
              x.ToEID(),
              EID{self.Id(),0},
              EID{l.Id(),0}))
            /* If-5 */} 
          }
          /* ERROR PROTECTION INSERTED (v_local3-_Zarg_try04573) */
          if ErrorIn(v_local3_try04585) {_Zarg_try04573 = v_local3_try04585
          _Zarg_try04573 = v_local3_try04585
          break
          } else {
          v_local3 = ANY(v_local3_try04585)
          ToList(OBJ(_Zarg_try04573)).PutAt(CLcount,v_local3)
          } 
        }
        /* Iteration-3 */} 
      /* ERROR PROTECTION INSERTED (_Zarg-Result) */
      if ErrorIn(_Zarg_try04573) {Result = _Zarg_try04573
      } else {
      _Zarg = ToList(OBJ(_Zarg_try04573))
      var g0463I *ClaireBoolean  
      /* and:3 */{ 
        var v_and3 *ClaireBoolean  
        
        v_and3 = C_compiler.Diet_ask
        if (v_and3 == CFALSE) {g0463I = CFALSE
        } else /* arg:4 */{ 
          /* Let:5 */{ 
            var g0464UU *ClaireAny  
            /* noccur = 1 */
            /* For:6 */{ 
              var x *ClaireAny  
              _ = x
              g0464UU= CFALSE.Id()
              var x_support *ClaireList  
              x_support = l
              x_len := x_support.Length()
              for i_it := 0; i_it < x_len; i_it++ { 
                x = x_support.At(i_it)
                if ((C_class.Id() == x.Isa.Id()) || 
                    (x.Isa.IsIn(C_property) == CTRUE)) /* If:8 */{ 
                   /*v = g0464UU, s =any*/
g0464UU = CTRUE.Id()
                  break
                  /* If-8 */} 
                /* loop-7 */} 
              /* For-6 */} 
            v_and3 = F_boolean_I_any(g0464UU)
            /* Let-5 */} 
          if (v_and3 == CFALSE) {g0463I = CFALSE
          } else /* arg:5 */{ 
            g0463I = CTRUE/* arg-5 */} 
          /* arg-4 */} 
        /* and-3 */} 
      if (g0463I == CTRUE) /* If:3 */{ 
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
        var g0465 *Language.Call   = Language.To_Call(self)
        /* noccur = 2 */
        /* Let:4 */{ 
          var l *ClaireList   = g0465.Args
          /* noccur = 4 */
          /* Let:5 */{ 
            var xs *ClaireObject  
            /* noccur = 4 */
            var xs_try04686 EID 
            if ((g0465.Selector.Id() == C_get.Id()) && 
                (l.At(1-1).Isa.IsIn(C_property) == CTRUE)) /* If:6 */{ 
              /* Let:7 */{ 
                var g0469UU *ClaireClass  
                /* noccur = 1 */
                var g0469UU_try04708 EID 
                /* Let:8 */{ 
                  var g0471UU *ClaireType  
                  /* noccur = 1 */
                  var g0471UU_try04729 EID 
                  g0471UU_try04729 = Core.F_CALL(C_c_type,ARGS(l.At(2-1).ToEID()))
                  /* ERROR PROTECTION INSERTED (g0471UU-g0469UU_try04708) */
                  if ErrorIn(g0471UU_try04729) {g0469UU_try04708 = g0471UU_try04729
                  } else {
                  g0471UU = ToType(OBJ(g0471UU_try04729))
                  g0469UU_try04708 = EID{g0471UU.Class_I().Id(),0}
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (g0469UU-xs_try04686) */
                if ErrorIn(g0469UU_try04708) {xs_try04686 = g0469UU_try04708
                } else {
                g0469UU = ToClass(OBJ(g0469UU_try04708))
                xs_try04686 = EID{Core.F__at_property1(ToProperty(l.At(1-1)),g0469UU).Id(),0}
                }
                /* Let-7 */} 
              } else {
              xs_try04686 = EID{CFALSE.Id(),0}
              /* If-6 */} 
            /* ERROR PROTECTION INSERTED (xs-Result) */
            if ErrorIn(xs_try04686) {Result = xs_try04686
            } else {
            xs = ToObject(OBJ(xs_try04686))
            var g0473I *ClaireBoolean  
            if (C_slot.Id() == xs.Isa.Id()) /* If:6 */{ 
              /* Let:7 */{ 
                var g0466 *ClaireSlot   = ToSlot(xs.Id())
                /* noccur = 4 */
                g0473I = MakeBoolean((b == CTRUE) || (g0466.Range.Contains(g0466.Default) == CTRUE) || (g0466.Srange.Id() == C_any.Id()) || (g0466.Srange.Id() == C_integer.Id()))
                /* Let-7 */} 
              } else {
              g0473I = CFALSE
              /* If-6 */} 
            if (g0473I == CTRUE) /* If:6 */{ 
              /* Let:7 */{ 
                var _CL_obj *Language.CallSlot   = Language.To_CallSlot(new(Language.CallSlot).Is(Language.C_Call_slot))
                /* noccur = 7 */
                _CL_obj.Selector = ToSlot(xs.Id())
                /* update:8 */{ 
                  var va_arg1 *Language.CallSlot  
                  var va_arg2 *ClaireAny  
                  va_arg1 = _CL_obj
                  var va_arg2_try04749 EID 
                  va_arg2_try04749 = Core.F_CALL(C_c_code,ARGS(l.At(2-1).ToEID(),EID{F_Compile_psort_any(Core.F_domain_I_restriction(ToRestriction(xs.Id())).Id()).Id(),0}))
                  /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                  if ErrorIn(va_arg2_try04749) {Result = va_arg2_try04749
                  } else {
                  va_arg2 = ANY(va_arg2_try04749)
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
        /* noccur = 6 */
        /* Let:4 */{ 
          var y *ClaireAny   = self.Args.At(3-1)
          /* noccur = 8 */
          /* Let:5 */{ 
            var yt *ClaireAny  
            /* noccur = 5 */
            var yt_try04786 EID 
            yt_try04786 = Core.F_CALL(C_c_type,ARGS(y.ToEID()))
            /* ERROR PROTECTION INSERTED (yt-Result) */
            if ErrorIn(yt_try04786) {Result = yt_try04786
            } else {
            yt = ANY(yt_try04786)
            /* Let:6 */{ 
              var ss *ClaireProperty   = self.Selector
              /* noccur = 5 */
              /* Let:7 */{ 
                var s *ClaireAny  
                /* noccur = 12 */
                var s_try04798 EID 
                /* Let:8 */{ 
                  var g0480UU *ClaireList  
                  /* noccur = 1 */
                  var g0480UU_try04819 EID 
                  /* Construct:9 */{ 
                    var v_bag_arg *ClaireAny  
                    g0480UU_try04819= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                    var v_bag_arg_try048210 EID 
                    v_bag_arg_try048210 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
                    /* ERROR PROTECTION INSERTED (v_bag_arg-g0480UU_try04819) */
                    if ErrorIn(v_bag_arg_try048210) {g0480UU_try04819 = v_bag_arg_try048210
                    } else {
                    v_bag_arg = ANY(v_bag_arg_try048210)
                    ToList(OBJ(g0480UU_try04819)).AddFast(v_bag_arg)}
                    /* Construct-9 */} 
                  /* ERROR PROTECTION INSERTED (g0480UU-s_try04798) */
                  if ErrorIn(g0480UU_try04819) {s_try04798 = g0480UU_try04819
                  } else {
                  g0480UU = ToList(OBJ(g0480UU_try04819))
                  s_try04798 = Core.F_CALL(C_Optimize_restriction_I,ARGS(p.ToEID(),EID{g0480UU.Id(),0},EID{CTRUE.Id(),0}))
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (s-Result) */
                if ErrorIn(s_try04798) {Result = s_try04798
                } else {
                s = ANY(s_try04798)
                
                if (C_OPT.ToRemove.Contain_ask(p) == CTRUE) /* If:8 */{ 
                  Result = EID{CNIL.Id(),0}
                  } else {
                  var g0483I *ClaireBoolean  
                  if (C_slot.Id() == s.Isa.Id()) /* If:9 */{ 
                    /* Let:10 */{ 
                      var g0475 *ClaireSlot   = ToSlot(s)
                      /* noccur = 1 */
                      g0483I = MakeBoolean((ToType(yt).Included(g0475.Range) == CTRUE) || (C_compiler.Safety >= 4))
                      /* Let-10 */} 
                    } else {
                    g0483I = CFALSE
                    /* If-9 */} 
                  if (g0483I == CTRUE) /* If:9 */{ 
                    var g0484I *ClaireBoolean  
                    var g0484I_try048510 EID 
                    /* and:10 */{ 
                      var v_and10 *ClaireBoolean  
                      
                      v_and10 = Core.F__I_equal_any(y,CNULL)
                      if (v_and10 == CFALSE) {g0484I_try048510 = EID{CFALSE.Id(),0}
                      } else /* arg:11 */{ 
                        var v_and10_try048612 EID 
                        /* Let:12 */{ 
                          var g0487UU *ClaireBoolean  
                          /* noccur = 1 */
                          var g0487UU_try048813 EID 
                          /* Let:13 */{ 
                            var g0489UU *ClaireAny  
                            /* noccur = 1 */
                            var g0489UU_try049014 EID 
                            g0489UU_try049014 = Core.F_CALL(ToProperty(C__exp.Id()),ARGS(yt.ToEID(),Core.F_CALL(C_mClaire_srange,ARGS(s.ToEID()))))
                            /* ERROR PROTECTION INSERTED (g0489UU-g0487UU_try048813) */
                            if ErrorIn(g0489UU_try049014) {g0487UU_try048813 = g0489UU_try049014
                            } else {
                            g0489UU = ANY(g0489UU_try049014)
                            g0487UU_try048813 = EID{F_boolean_I_any(g0489UU).Id(),0}
                            }
                            /* Let-13 */} 
                          /* ERROR PROTECTION INSERTED (g0487UU-v_and10_try048612) */
                          if ErrorIn(g0487UU_try048813) {v_and10_try048612 = g0487UU_try048813
                          } else {
                          g0487UU = ToBoolean(OBJ(g0487UU_try048813))
                          v_and10_try048612 = EID{Core.F__I_equal_any(g0487UU.Id(),CTRUE.Id()).Id(),0}
                          }
                          /* Let-12 */} 
                        /* ERROR PROTECTION INSERTED (v_and10-g0484I_try048510) */
                        if ErrorIn(v_and10_try048612) {g0484I_try048510 = v_and10_try048612
                        } else {
                        v_and10 = ToBoolean(OBJ(v_and10_try048612))
                        if (v_and10 == CFALSE) {g0484I_try048510 = EID{CFALSE.Id(),0}
                        } else /* arg:12 */{ 
                          g0484I_try048510 = EID{CTRUE.Id(),0}/* arg-12 */} 
                        /* arg-11 */} 
                      }
                      /* and-10 */} 
                    /* ERROR PROTECTION INSERTED (g0484I-Result) */
                    if ErrorIn(g0484I_try048510) {Result = g0484I_try048510
                    } else {
                    g0484I = ToBoolean(OBJ(g0484I_try048510))
                    if (g0484I == CTRUE) /* If:10 */{ 
                      
                      F_Compile_warn_void()
                      Result = Core.F_tformat_string(MakeString("sort error in ~S: ~S is a ~S [253]\n"),2,MakeConstantList(self.Id(),y,yt))
                      } else {
                      Result = EID{CFALSE.Id(),0}
                      /* If-10 */} 
                    }
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    var g0491I *ClaireBoolean  
                    var g0491I_try049210 EID 
                    /* and:10 */{ 
                      var v_and10 *ClaireBoolean  
                      
                      v_and10 = MakeBoolean((ToType(yt).Included(ToType(OBJ(Core.F_CALL(C_range,ARGS(s.ToEID()))))) == CTRUE) || (ToType(yt).Included(ToType(C_object.Id())) == CTRUE) || (ANY(Core.F_CALL(C_mClaire_srange,ARGS(s.ToEID()))) != C_object.Id()) || (y == CNULL))
                      if (v_and10 == CFALSE) {g0491I_try049210 = EID{CFALSE.Id(),0}
                      } else /* arg:11 */{ 
                        var v_and10_try049312 EID 
                        /* or:12 */{ 
                          var v_or12 *ClaireBoolean  
                          
                          v_or12 = Core.F__I_equal_any(ss.Id(),Core.C_write.Id())
                          if (v_or12 == CTRUE) {v_and10_try049312 = EID{CTRUE.Id(),0}
                          } else /* or:13 */{ 
                            var v_or12_try049414 EID 
                            /* and:14 */{ 
                              var v_and14 *ClaireBoolean  
                              
                              var v_and14_try049515 EID 
                              v_and14_try049515 = F_Optimize_Update_ask_relation1(ToRelation(p),x,y)
                              /* ERROR PROTECTION INSERTED (v_and14-v_or12_try049414) */
                              if ErrorIn(v_and14_try049515) {v_or12_try049414 = v_and14_try049515
                              } else {
                              v_and14 = ToBoolean(OBJ(v_and14_try049515))
                              if (v_and14 == CFALSE) {v_or12_try049414 = EID{CFALSE.Id(),0}
                              } else /* arg:15 */{ 
                                v_and14 = MakeBoolean((ToRelation(p).Multivalued_ask.Id() == CFALSE.Id()) || (Core.F_get_property(C_if_write,ToObject(p)) == CNULL))
                                if (v_and14 == CFALSE) {v_or12_try049414 = EID{CFALSE.Id(),0}
                                } else /* arg:16 */{ 
                                  v_or12_try049414 = EID{CTRUE.Id(),0}/* arg-16 */} 
                                /* arg-15 */} 
                              }
                              /* and-14 */} 
                            /* ERROR PROTECTION INSERTED (v_or12-v_and10_try049312) */
                            if ErrorIn(v_or12_try049414) {v_and10_try049312 = v_or12_try049414
                            } else {
                            v_or12 = ToBoolean(OBJ(v_or12_try049414))
                            if (v_or12 == CTRUE) {v_and10_try049312 = EID{CTRUE.Id(),0}
                            } else /* or:14 */{ 
                              v_and10_try049312 = EID{CFALSE.Id(),0}/* org-14 */} 
                            /* org-13 */} 
                          }
                          /* or-12 */} 
                        /* ERROR PROTECTION INSERTED (v_and10-g0491I_try049210) */
                        if ErrorIn(v_and10_try049312) {g0491I_try049210 = v_and10_try049312
                        } else {
                        v_and10 = ToBoolean(OBJ(v_and10_try049312))
                        if (v_and10 == CFALSE) {g0491I_try049210 = EID{CFALSE.Id(),0}
                        } else /* arg:12 */{ 
                          g0491I_try049210 = EID{CTRUE.Id(),0}/* arg-12 */} 
                        /* arg-11 */} 
                      }
                      /* and-10 */} 
                    /* ERROR PROTECTION INSERTED (g0491I-Result) */
                    if ErrorIn(g0491I_try049210) {Result = g0491I_try049210
                    } else {
                    g0491I = ToBoolean(OBJ(g0491I_try049210))
                    if (g0491I == CTRUE) /* If:10 */{ 
                      /* Let:11 */{ 
                        var _Zx *ClaireAny  
                        /* noccur = 1 */
                        var _Zx_try049612 EID 
                        _Zx_try049612 = Core.F_CALL(C_c_code,ARGS(x.ToEID(),EID{F_Compile_psort_any(Core.F_domain_I_restriction(ToRestriction(s)).Id()).Id(),0}))
                        /* ERROR PROTECTION INSERTED (_Zx-Result) */
                        if ErrorIn(_Zx_try049612) {Result = _Zx_try049612
                        } else {
                        _Zx = ANY(_Zx_try049612)
                        /* Let:12 */{ 
                          var _Zy *ClaireAny  
                          /* noccur = 1 */
                          var _Zy_try049713 EID 
                          _Zy_try049713 = F_Compile_c_strict_code_any(y,F_Compile_psort_any(ANY(Core.F_CALL(C_range,ARGS(s.ToEID())))))
                          /* ERROR PROTECTION INSERTED (_Zy-Result) */
                          if ErrorIn(_Zy_try049713) {Result = _Zy_try049713
                          } else {
                          _Zy = ANY(_Zy_try049713)
                          /* Let:13 */{ 
                            var _CL_obj *Language.Update   = Language.To_Update(new(Language.Update).Is(Language.C_Update))
                            /* noccur = 16 */
                            _CL_obj.Selector = p
                            _CL_obj.Value = _Zy
                            /* update:14 */{ 
                              var va_arg1 *Language.Update  
                              var va_arg2 *ClaireAny  
                              va_arg1 = _CL_obj
                              var va_arg2_try049815 EID 
                              if (ss.Id() != Core.C_write.Id()) /* If:15 */{ 
                                va_arg2_try049815 = EID{ss.Id(),0}
                                } else {
                                va_arg2_try049815 = Core.F_CALL(C_c_code,ARGS(x.ToEID(),EID{C_any.Id(),0}))
                                /* If-15 */} 
                              /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                              if ErrorIn(va_arg2_try049815) {Result = va_arg2_try049815
                              } else {
                              va_arg2 = ANY(va_arg2_try049815)
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
                        var g0499UU *Language.Call  
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
                          g0499UU = _CL_obj
                          /* Let-12 */} 
                        Result = Core.F_CALL(C_c_code,ARGS(EID{g0499UU.Id(),0}))
                        /* Let-11 */} 
                      } else {
                      if (C_compiler.Diet_ask == CTRUE) /* If:11 */{ 
                        F_Compile_warn_void()
                        Core.F_tformat_string(MakeString("~S is not a diet call [254]"),2,MakeConstantList(self.Id()))
                        /* If-11 */} 
                      if ((C_compiler.Optimize_ask == CTRUE) && 
                          (p != C_instances.Id())) /* If:11 */{ 
                        F_Compile_notice_void()
                        
                        /* If-11 */} 
                      /* Let:11 */{ 
                        var g0500UU *Language.Call  
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
                          g0500UU = _CL_obj
                          /* Let-12 */} 
                        Result = Core.F_CALL(C_c_code,ARGS(EID{g0500UU.Id(),0}))
                        /* Let-11 */} 
                      /* If-10 */} 
                    }
                    }
                    } else {
                    /* Let:10 */{ 
                      var _Ztype *ClaireList  
                      /* noccur = 3 */
                      var _Ztype_try050111 EID 
                      /* Iteration:11 */{ 
                        var v_list11 *ClaireList  
                        var x *ClaireAny  
                        var v_local11 *ClaireAny  
                        v_list11 = self.Args
                        _Ztype_try050111 = EID{CreateList(ToType(CEMPTY.Id()),v_list11.Length()).Id(),0}
                        for CLcount := 0; CLcount < v_list11.Length(); CLcount++{ 
                          x = v_list11.At(CLcount)
                          var v_local11_try050213 EID 
                          v_local11_try050213 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
                          /* ERROR PROTECTION INSERTED (v_local11-_Ztype_try050111) */
                          if ErrorIn(v_local11_try050213) {_Ztype_try050111 = v_local11_try050213
                          _Ztype_try050111 = v_local11_try050213
                          break
                          } else {
                          v_local11 = ANY(v_local11_try050213)
                          ToList(OBJ(_Ztype_try050111)).PutAt(CLcount,v_local11)
                          } 
                        }
                        /* Iteration-11 */} 
                      /* ERROR PROTECTION INSERTED (_Ztype-Result) */
                      if ErrorIn(_Ztype_try050111) {Result = _Ztype_try050111
                      } else {
                      _Ztype = ToList(OBJ(_Ztype_try050111))
                      /* Let:11 */{ 
                        var z *ClaireAny   = F_Optimize_restriction_I_property(ss,_Ztype,CTRUE)
                        /* noccur = 2 */
                        
                        if (C_method.Id() == z.Isa.Id()) /* If:12 */{ 
                          /* Let:13 */{ 
                            var g0476 *ClaireMethod   = ToMethod(z)
                            /* noccur = 1 */
                            Result = F_Optimize_c_code_method_method1(g0476,self.Args,_Ztype)
                            /* Let-13 */} 
                          } else {
                          Result = F_Optimize_c_warn_Call(self,_Ztype.Id())
                          /* If-12 */} 
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
      var s_try05043 EID 
      /* Let:3 */{ 
        var g0505UU *ClaireList  
        /* noccur = 1 */
        var g0505UU_try05064 EID 
        /* Construct:4 */{ 
          var v_bag_arg *ClaireAny  
          g0505UU_try05064= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
          var v_bag_arg_try05075 EID 
          v_bag_arg_try05075 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
          /* ERROR PROTECTION INSERTED (v_bag_arg-g0505UU_try05064) */
          if ErrorIn(v_bag_arg_try05075) {g0505UU_try05064 = v_bag_arg_try05075
          } else {
          v_bag_arg = ANY(v_bag_arg_try05075)
          ToList(OBJ(g0505UU_try05064)).AddFast(v_bag_arg)}
          /* Construct-4 */} 
        /* ERROR PROTECTION INSERTED (g0505UU-s_try05043) */
        if ErrorIn(g0505UU_try05064) {s_try05043 = g0505UU_try05064
        } else {
        g0505UU = ToList(OBJ(g0505UU_try05064))
        s_try05043 = F_Optimize_restriction_I_property(p,g0505UU,CTRUE).ToEID()
        }
        /* Let-3 */} 
      /* ERROR PROTECTION INSERTED (s-Result) */
      if ErrorIn(s_try05043) {Result = s_try05043
      } else {
      s = ANY(s_try05043)
      var g0508I *ClaireBoolean  
      var g0508I_try05093 EID 
      if (C_slot.Id() == s.Isa.Id()) /* If:3 */{ 
        /* Let:4 */{ 
          var g0503 *ClaireSlot   = ToSlot(s)
          /* noccur = 1 */
          /* or:5 */{ 
            var v_or5 *ClaireBoolean  
            
            v_or5 = Equal(y,CNULL)
            if (v_or5 == CTRUE) {g0508I_try05093 = EID{CTRUE.Id(),0}
            } else /* or:6 */{ 
              var v_or5_try05107 EID 
              /* and:7 */{ 
                var v_and7 *ClaireBoolean  
                
                var v_and7_try05118 EID 
                /* Let:8 */{ 
                  var g0512UU *ClaireType  
                  /* noccur = 1 */
                  var g0512UU_try05139 EID 
                  g0512UU_try05139 = Core.F_CALL(C_c_type,ARGS(y.ToEID()))
                  /* ERROR PROTECTION INSERTED (g0512UU-v_and7_try05118) */
                  if ErrorIn(g0512UU_try05139) {v_and7_try05118 = g0512UU_try05139
                  } else {
                  g0512UU = ToType(OBJ(g0512UU_try05139))
                  v_and7_try05118 = EID{g0512UU.Included(ToType(g0503.Srange.Id())).Id(),0}
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (v_and7-v_or5_try05107) */
                if ErrorIn(v_and7_try05118) {v_or5_try05107 = v_and7_try05118
                } else {
                v_and7 = ToBoolean(OBJ(v_and7_try05118))
                if (v_and7 == CFALSE) {v_or5_try05107 = EID{CFALSE.Id(),0}
                } else /* arg:8 */{ 
                  var v_and7_try05149 EID 
                  v_and7_try05149 = F_Compile_identifiable_ask_any(y)
                  /* ERROR PROTECTION INSERTED (v_and7-v_or5_try05107) */
                  if ErrorIn(v_and7_try05149) {v_or5_try05107 = v_and7_try05149
                  } else {
                  v_and7 = ToBoolean(OBJ(v_and7_try05149))
                  if (v_and7 == CFALSE) {v_or5_try05107 = EID{CFALSE.Id(),0}
                  } else /* arg:9 */{ 
                    v_or5_try05107 = EID{CTRUE.Id(),0}/* arg-9 */} 
                  /* arg-8 */} 
                }}
                /* and-7 */} 
              /* ERROR PROTECTION INSERTED (v_or5-g0508I_try05093) */
              if ErrorIn(v_or5_try05107) {g0508I_try05093 = v_or5_try05107
              } else {
              v_or5 = ToBoolean(OBJ(v_or5_try05107))
              if (v_or5 == CTRUE) {g0508I_try05093 = EID{CTRUE.Id(),0}
              } else /* or:7 */{ 
                g0508I_try05093 = EID{CFALSE.Id(),0}/* org-7 */} 
              /* org-6 */} 
            }
            /* or-5 */} 
          /* Let-4 */} 
        } else {
        g0508I_try05093 = EID{CFALSE.Id(),0}
        /* If-3 */} 
      /* ERROR PROTECTION INSERTED (g0508I-Result) */
      if ErrorIn(g0508I_try05093) {Result = g0508I_try05093
      } else {
      g0508I = ToBoolean(OBJ(g0508I_try05093))
      if (g0508I == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var cs *Language.CallSlot  
          /* noccur = 1 */
          var cs_try05155 EID 
          /* Let:5 */{ 
            var _CL_obj *Language.CallSlot   = Language.To_CallSlot(new(Language.CallSlot).Is(Language.C_Call_slot))
            /* noccur = 7 */
            _CL_obj.Selector = ToSlot(s)
            /* update:6 */{ 
              var va_arg1 *Language.CallSlot  
              var va_arg2 *ClaireAny  
              va_arg1 = _CL_obj
              var va_arg2_try05167 EID 
              va_arg2_try05167 = Core.F_CALL(C_c_code,ARGS(x.ToEID(),EID{F_Compile_psort_any(Core.F_domain_I_restriction(ToRestriction(s)).Id()).Id(),0}))
              /* ERROR PROTECTION INSERTED (va_arg2-cs_try05155) */
              if ErrorIn(va_arg2_try05167) {cs_try05155 = va_arg2_try05167
              } else {
              va_arg2 = ANY(va_arg2_try05167)
              /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
              va_arg1.Arg = va_arg2
              cs_try05155 = va_arg2.ToEID()
              }
              /* update-6 */} 
            /* ERROR PROTECTION INSERTED (cs_try05155-cs_try05155) */
            if !ErrorIn(cs_try05155) {
            _CL_obj.Test = CFALSE
            cs_try05155 = EID{_CL_obj.Id(),0}
            }
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (cs-Result) */
          if ErrorIn(cs_try05155) {Result = cs_try05155
          } else {
          cs = Language.To_CallSlot(OBJ(cs_try05155))
          /* Let:5 */{ 
            var cm *Language.CallMethod2  
            /* noccur = 2 */
            var cm_try05176 EID 
            /* Let:6 */{ 
              var _CL_obj *Language.CallMethod2   = Language.To_CallMethod2(new(Language.CallMethod2).Is(Language.C_Call_method2))
              /* noccur = 5 */
              _CL_obj.Arg = ToMethod(Core.F__at_property1(Core.C_identical_ask,C_any).Id())
              /* update:7 */{ 
                var va_arg1 *Language.CallMethod  
                var va_arg2 *ClaireList  
                va_arg1 = Language.To_CallMethod(_CL_obj.Id())
                var va_arg2_try05188 EID 
                /* Construct:8 */{ 
                  var v_bag_arg *ClaireAny  
                  va_arg2_try05188= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                  ToList(OBJ(va_arg2_try05188)).AddFast(cs.Id())
                  var v_bag_arg_try05199 EID 
                  v_bag_arg_try05199 = Core.F_CALL(C_c_code,ARGS(y.ToEID(),Core.F_CALL(C_mClaire_srange,ARGS(s.ToEID()))))
                  /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try05188) */
                  if ErrorIn(v_bag_arg_try05199) {va_arg2_try05188 = v_bag_arg_try05199
                  } else {
                  v_bag_arg = ANY(v_bag_arg_try05199)
                  ToList(OBJ(va_arg2_try05188)).AddFast(v_bag_arg)}
                  /* Construct-8 */} 
                /* ERROR PROTECTION INSERTED (va_arg2-cm_try05176) */
                if ErrorIn(va_arg2_try05188) {cm_try05176 = va_arg2_try05188
                } else {
                va_arg2 = ToList(OBJ(va_arg2_try05188))
                /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                va_arg1.Args = va_arg2
                cm_try05176 = EID{va_arg2.Id(),0}
                }
                /* update-7 */} 
              /* ERROR PROTECTION INSERTED (cm_try05176-cm_try05176) */
              if !ErrorIn(cm_try05176) {
              cm_try05176 = EID{_CL_obj.Id(),0}
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (cm-Result) */
            if ErrorIn(cm_try05176) {Result = cm_try05176
            } else {
            cm = Language.To_CallMethod2(OBJ(cm_try05176))
            if (b == CTRUE) /* If:6 */{ 
              Result = Core.F_CALL(C_c_code,ARGS(EID{cm.Id(),0}))
              } else {
              /* Let:7 */{ 
                var g0520UU *Language.Call  
                /* noccur = 1 */
                /* Let:8 */{ 
                  var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                  /* noccur = 5 */
                  _CL_obj.Selector = Core.C_not
                  _CL_obj.Args = MakeConstantList(cm.Id())
                  g0520UU = _CL_obj
                  /* Let-8 */} 
                Result = Core.F_CALL(C_c_code,ARGS(EID{g0520UU.Id(),0}))
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
            var g0521UU *ClaireList  
            /* noccur = 1 */
            var g0521UU_try05226 EID 
            /* Construct:6 */{ 
              var v_bag_arg *ClaireAny  
              g0521UU_try05226= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
              var v_bag_arg_try05237 EID 
              /* Let:7 */{ 
                var g0524UU *Language.Call  
                /* noccur = 1 */
                /* Let:8 */{ 
                  var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                  /* noccur = 5 */
                  _CL_obj.Selector = C_get
                  _CL_obj.Args = MakeConstantList(p.Id(),x)
                  g0524UU = _CL_obj
                  /* Let-8 */} 
                v_bag_arg_try05237 = Core.F_CALL(C_c_code,ARGS(EID{g0524UU.Id(),0},EID{C_any.Id(),0}))
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (v_bag_arg-g0521UU_try05226) */
              if ErrorIn(v_bag_arg_try05237) {g0521UU_try05226 = v_bag_arg_try05237
              } else {
              v_bag_arg = ANY(v_bag_arg_try05237)
              ToList(OBJ(g0521UU_try05226)).AddFast(v_bag_arg)
              var v_bag_arg_try05257 EID 
              v_bag_arg_try05257 = Core.F_CALL(C_c_code,ARGS(y.ToEID(),EID{C_any.Id(),0}))
              /* ERROR PROTECTION INSERTED (v_bag_arg-g0521UU_try05226) */
              if ErrorIn(v_bag_arg_try05257) {g0521UU_try05226 = v_bag_arg_try05257
              } else {
              v_bag_arg = ANY(v_bag_arg_try05257)
              ToList(OBJ(g0521UU_try05226)).AddFast(v_bag_arg)}}
              /* Construct-6 */} 
            /* ERROR PROTECTION INSERTED (g0521UU-Result) */
            if ErrorIn(g0521UU_try05226) {Result = g0521UU_try05226
            } else {
            g0521UU = ToList(OBJ(g0521UU_try05226))
            Result = F_Optimize_c_code_method_method1(ToMethod(Core.F__at_property2(ToProperty(IfThenElse((b == CTRUE),
              C__equal.Id(),
              Core.C__I_equal.Id())),l).Id()),g0521UU,l)
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
            var s_try05276 EID 
            /* Let:6 */{ 
              var g0528UU *ClaireClass  
              /* noccur = 1 */
              var g0528UU_try05297 EID 
              /* Let:7 */{ 
                var g0530UU *ClaireType  
                /* noccur = 1 */
                var g0530UU_try05318 EID 
                /* Let:8 */{ 
                  var g0532UU *ClaireType  
                  /* noccur = 1 */
                  var g0532UU_try05339 EID 
                  g0532UU_try05339 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
                  /* ERROR PROTECTION INSERTED (g0532UU-g0530UU_try05318) */
                  if ErrorIn(g0532UU_try05339) {g0530UU_try05318 = g0532UU_try05339
                  } else {
                  g0532UU = ToType(OBJ(g0532UU_try05339))
                  g0530UU_try05318 = EID{F_Optimize_ptype_type(g0532UU).Id(),0}
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (g0530UU-g0528UU_try05297) */
                if ErrorIn(g0530UU_try05318) {g0528UU_try05297 = g0530UU_try05318
                } else {
                g0530UU = ToType(OBJ(g0530UU_try05318))
                g0528UU_try05297 = EID{g0530UU.Class_I().Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (g0528UU-s_try05276) */
              if ErrorIn(g0528UU_try05297) {s_try05276 = g0528UU_try05297
              } else {
              g0528UU = ToClass(OBJ(g0528UU_try05297))
              s_try05276 = EID{Core.F__at_property1(p,g0528UU).Id(),0}
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (s-Result) */
            if ErrorIn(s_try05276) {Result = s_try05276
            } else {
            s = ToObject(OBJ(s_try05276))
            var g0534I *ClaireBoolean  
            var g0534I_try05356 EID 
            if (C_slot.Id() == s.Isa.Id()) /* If:6 */{ 
              /* Let:7 */{ 
                var g0526 *ClaireSlot   = ToSlot(s.Id())
                /* noccur = 1 */
                /* or:8 */{ 
                  var v_or8 *ClaireBoolean  
                  
                  var v_or8_try05369 EID 
                  /* Let:9 */{ 
                    var g0537UU *ClaireType  
                    /* noccur = 1 */
                    var g0537UU_try053810 EID 
                    g0537UU_try053810 = Core.F_CALL(C_c_type,ARGS(y.ToEID()))
                    /* ERROR PROTECTION INSERTED (g0537UU-v_or8_try05369) */
                    if ErrorIn(g0537UU_try053810) {v_or8_try05369 = g0537UU_try053810
                    } else {
                    g0537UU = ToType(OBJ(g0537UU_try053810))
                    v_or8_try05369 = EID{g0537UU.Included(Core.F_member_type(g0526.Range)).Id(),0}
                    }
                    /* Let-9 */} 
                  /* ERROR PROTECTION INSERTED (v_or8-g0534I_try05356) */
                  if ErrorIn(v_or8_try05369) {g0534I_try05356 = v_or8_try05369
                  } else {
                  v_or8 = ToBoolean(OBJ(v_or8_try05369))
                  if (v_or8 == CTRUE) {g0534I_try05356 = EID{CTRUE.Id(),0}
                  } else /* or:9 */{ 
                    v_or8 = F__sup_equal_integer(C_compiler.Safety,4)
                    if (v_or8 == CTRUE) {g0534I_try05356 = EID{CTRUE.Id(),0}
                    } else /* or:10 */{ 
                      g0534I_try05356 = EID{CFALSE.Id(),0}/* org-10 */} 
                    /* org-9 */} 
                  }
                  /* or-8 */} 
                /* Let-7 */} 
              } else {
              g0534I_try05356 = EID{CFALSE.Id(),0}
              /* If-6 */} 
            /* ERROR PROTECTION INSERTED (g0534I-Result) */
            if ErrorIn(g0534I_try05356) {Result = g0534I_try05356
            } else {
            g0534I = ToBoolean(OBJ(g0534I_try05356))
            if (g0534I == CTRUE) /* If:6 */{ 
              if (F_Optimize_Update_ask_relation2(ToRelation(p.Id()),ToRelation(self.Selector.Id())) == CTRUE) /* If:7 */{ 
                /* Let:8 */{ 
                  var x2 *ClaireAny  
                  /* noccur = 1 */
                  var x2_try05399 EID 
                  x2_try05399 = Core.F_CALL(C_c_code,ARGS(x.ToEID(),EID{F_Compile_psort_any(Core.F_domain_I_restriction(ToRestriction(s.Id())).Id()).Id(),0}))
                  /* ERROR PROTECTION INSERTED (x2-Result) */
                  if ErrorIn(x2_try05399) {Result = x2_try05399
                  } else {
                  x2 = ANY(x2_try05399)
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
                      var va_arg2_try054011 EID 
                      va_arg2_try054011 = Core.F_CALL(C_c_code,ARGS(y.ToEID(),EID{F_Compile_psort_any(Core.F_member_type(ToType(OBJ(Core.F_CALL(C_range,ARGS(EID{s.Id(),0}))))).Id()).Id(),0}))
                      /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                      if ErrorIn(va_arg2_try054011) {Result = va_arg2_try054011
                      } else {
                      va_arg2 = ANY(va_arg2_try054011)
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
                var g0541I *ClaireBoolean  
                var g0541I_try05428 EID 
                /* and:8 */{ 
                  var v_and8 *ClaireBoolean  
                  
                  var v_and8_try05439 EID 
                  v_and8_try05439 = F_Compile_designated_ask_any(x)
                  /* ERROR PROTECTION INSERTED (v_and8-g0541I_try05428) */
                  if ErrorIn(v_and8_try05439) {g0541I_try05428 = v_and8_try05439
                  } else {
                  v_and8 = ToBoolean(OBJ(v_and8_try05439))
                  if (v_and8 == CFALSE) {g0541I_try05428 = EID{CFALSE.Id(),0}
                  } else /* arg:9 */{ 
                    v_and8 = MakeBoolean((p.Store_ask != CTRUE) && ((self.Selector.Id() == C_add_I.Id()) || 
                        (p.Inverse.Id() == CNULL)))
                    if (v_and8 == CFALSE) {g0541I_try05428 = EID{CFALSE.Id(),0}
                    } else /* arg:10 */{ 
                      g0541I_try05428 = EID{CTRUE.Id(),0}/* arg-10 */} 
                    /* arg-9 */} 
                  }
                  /* and-8 */} 
                /* ERROR PROTECTION INSERTED (g0541I-Result) */
                if ErrorIn(g0541I_try05428) {Result = g0541I_try05428
                } else {
                g0541I = ToBoolean(OBJ(g0541I_try05428))
                if (g0541I == CTRUE) /* If:8 */{ 
                  /* Let:9 */{ 
                    var x2 *ClaireAny  
                    /* noccur = 1 */
                    var x2_try054410 EID 
                    x2_try054410 = Core.F_CALL(C_c_code,ARGS(x.ToEID(),EID{F_Compile_psort_any(Core.F_domain_I_restriction(ToRestriction(s.Id())).Id()).Id(),0}))
                    /* ERROR PROTECTION INSERTED (x2-Result) */
                    if ErrorIn(x2_try054410) {Result = x2_try054410
                    } else {
                    x2 = ANY(x2_try054410)
                    /* Let:10 */{ 
                      var g0545UU *Language.Call  
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
                        g0545UU = _CL_obj
                        /* Let-11 */} 
                      Result = Core.F_CALL(C_c_code,ARGS(EID{g0545UU.Id(),0}))
                      /* Let-10 */} 
                    }
                    /* Let-9 */} 
                  } else {
                  if (C_compiler.Optimize_ask == CTRUE) /* If:9 */{ 
                    F_Compile_notice_void()
                    
                    /* If-9 */} 
                  /* Let:9 */{ 
                    var g0546UU *ClaireList  
                    /* noccur = 1 */
                    var g0546UU_try054710 EID 
                    /* Construct:10 */{ 
                      var v_bag_arg *ClaireAny  
                      g0546UU_try054710= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                      ToList(OBJ(g0546UU_try054710)).AddFast(C_property.Id())
                      var v_bag_arg_try054811 EID 
                      v_bag_arg_try054811 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
                      /* ERROR PROTECTION INSERTED (v_bag_arg-g0546UU_try054710) */
                      if ErrorIn(v_bag_arg_try054811) {g0546UU_try054710 = v_bag_arg_try054811
                      } else {
                      v_bag_arg = ANY(v_bag_arg_try054811)
                      ToList(OBJ(g0546UU_try054710)).AddFast(v_bag_arg)
                      ToList(OBJ(g0546UU_try054710)).AddFast(C_integer.Id())
                      var v_bag_arg_try054911 EID 
                      v_bag_arg_try054911 = Core.F_CALL(C_c_type,ARGS(y.ToEID()))
                      /* ERROR PROTECTION INSERTED (v_bag_arg-g0546UU_try054710) */
                      if ErrorIn(v_bag_arg_try054911) {g0546UU_try054710 = v_bag_arg_try054911
                      } else {
                      v_bag_arg = ANY(v_bag_arg_try054911)
                      ToList(OBJ(g0546UU_try054710)).AddFast(v_bag_arg)}}
                      /* Construct-10 */} 
                    /* ERROR PROTECTION INSERTED (g0546UU-Result) */
                    if ErrorIn(g0546UU_try054710) {Result = g0546UU_try054710
                    } else {
                    g0546UU = ToList(OBJ(g0546UU_try054710))
                    Result = F_Optimize_c_code_method_method1(ToMethod(Core.F__at_property1(ToProperty(C_add_I.Id()),C_property).Id()),MakeConstantList(p.Id(),
                      x,
                      ANY(Core.F_CALL(C_mClaire_index,ARGS(EID{s.Id(),0}))),
                      y),g0546UU)
                    }
                    /* Let-9 */} 
                  /* If-8 */} 
                }
                /* If-7 */} 
              } else {
              /* Let:7 */{ 
                var g0550UU *ClaireList  
                /* noccur = 1 */
                var g0550UU_try05518 EID 
                /* Iteration:8 */{ 
                  var v_list8 *ClaireList  
                  var x *ClaireAny  
                  var v_local8 *ClaireAny  
                  v_list8 = self.Args
                  g0550UU_try05518 = EID{CreateList(ToType(CEMPTY.Id()),v_list8.Length()).Id(),0}
                  for CLcount := 0; CLcount < v_list8.Length(); CLcount++{ 
                    x = v_list8.At(CLcount)
                    var v_local8_try055210 EID 
                    v_local8_try055210 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
                    /* ERROR PROTECTION INSERTED (v_local8-g0550UU_try05518) */
                    if ErrorIn(v_local8_try055210) {g0550UU_try05518 = v_local8_try055210
                    g0550UU_try05518 = v_local8_try055210
                    break
                    } else {
                    v_local8 = ANY(v_local8_try055210)
                    ToList(OBJ(g0550UU_try05518)).PutAt(CLcount,v_local8)
                    } 
                  }
                  /* Iteration-8 */} 
                /* ERROR PROTECTION INSERTED (g0550UU-Result) */
                if ErrorIn(g0550UU_try05518) {Result = g0550UU_try05518
                } else {
                g0550UU = ToList(OBJ(g0550UU_try05518))
                Result = F_Optimize_c_code_method_method1(ToMethod(Core.F__at_property1(ToProperty(C_add.Id()),C_property).Id()),self.Args,g0550UU)
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
      var _Zt1_try05553 EID 
      _Zt1_try05553 = Core.F_CALL(C_c_type,ARGS(self.Args.At(1-1).ToEID()))
      /* ERROR PROTECTION INSERTED (_Zt1-Result) */
      if ErrorIn(_Zt1_try05553) {Result = _Zt1_try05553
      } else {
      _Zt1 = ANY(_Zt1_try05553)
      /* Let:3 */{ 
        var _Zt2 *ClaireType  
        /* noccur = 3 */
        var _Zt2_try05564 EID 
        /* Let:4 */{ 
          var g0557UU *ClaireType  
          /* noccur = 1 */
          var g0557UU_try05585 EID 
          g0557UU_try05585 = Core.F_CALL(C_c_type,ARGS(self.Args.At(2-1).ToEID()))
          /* ERROR PROTECTION INSERTED (g0557UU-_Zt2_try05564) */
          if ErrorIn(g0557UU_try05585) {_Zt2_try05564 = g0557UU_try05585
          } else {
          g0557UU = ToType(OBJ(g0557UU_try05585))
          _Zt2_try05564 = EID{F_Optimize_ptype_type(g0557UU).Id(),0}
          }
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (_Zt2-Result) */
        if ErrorIn(_Zt2_try05564) {Result = _Zt2_try05564
        } else {
        _Zt2 = ToType(OBJ(_Zt2_try05564))
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
                  var g0553 *ClaireMethod   = ToMethod(z)
                  /* noccur = 1 */
                  Result = F_Optimize_c_code_method_method1(g0553,self.Args,_Zltype)
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
            var s_try05606 EID 
            /* Let:6 */{ 
              var g0561UU *ClaireClass  
              /* noccur = 1 */
              var g0561UU_try05627 EID 
              /* Let:7 */{ 
                var g0563UU *ClaireType  
                /* noccur = 1 */
                var g0563UU_try05648 EID 
                g0563UU_try05648 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
                /* ERROR PROTECTION INSERTED (g0563UU-g0561UU_try05627) */
                if ErrorIn(g0563UU_try05648) {g0561UU_try05627 = g0563UU_try05648
                } else {
                g0563UU = ToType(OBJ(g0563UU_try05648))
                g0561UU_try05627 = EID{g0563UU.Class_I().Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (g0561UU-s_try05606) */
              if ErrorIn(g0561UU_try05627) {s_try05606 = g0561UU_try05627
              } else {
              g0561UU = ToClass(OBJ(g0561UU_try05627))
              s_try05606 = EID{Core.F__at_property1(ToProperty(p),g0561UU).Id(),0}
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (s-Result) */
            if ErrorIn(s_try05606) {Result = s_try05606
            } else {
            s = ToObject(OBJ(s_try05606))
            var g0565I *ClaireBoolean  
            var g0565I_try05666 EID 
            /* and:6 */{ 
              var v_and6 *ClaireBoolean  
              
              v_and6 = MakeBoolean((ToRelation(p).Inverse.Id() == CNULL))
              if (v_and6 == CFALSE) {g0565I_try05666 = EID{CFALSE.Id(),0}
              } else /* arg:7 */{ 
                var v_and6_try05678 EID 
                v_and6_try05678 = F_Compile_designated_ask_any(x)
                /* ERROR PROTECTION INSERTED (v_and6-g0565I_try05666) */
                if ErrorIn(v_and6_try05678) {g0565I_try05666 = v_and6_try05678
                } else {
                v_and6 = ToBoolean(OBJ(v_and6_try05678))
                if (v_and6 == CFALSE) {g0565I_try05666 = EID{CFALSE.Id(),0}
                } else /* arg:8 */{ 
                  var v_and6_try05689 EID 
                  if (C_slot.Id() == s.Isa.Id()) /* If:9 */{ 
                    /* Let:10 */{ 
                      var g0559 *ClaireSlot   = ToSlot(s.Id())
                      /* noccur = 1 */
                      /* or:11 */{ 
                        var v_or11 *ClaireBoolean  
                        
                        var v_or11_try056912 EID 
                        /* Let:12 */{ 
                          var g0570UU *ClaireType  
                          /* noccur = 1 */
                          var g0570UU_try057113 EID 
                          g0570UU_try057113 = Core.F_CALL(C_c_type,ARGS(y.ToEID()))
                          /* ERROR PROTECTION INSERTED (g0570UU-v_or11_try056912) */
                          if ErrorIn(g0570UU_try057113) {v_or11_try056912 = g0570UU_try057113
                          } else {
                          g0570UU = ToType(OBJ(g0570UU_try057113))
                          v_or11_try056912 = EID{g0570UU.Included(Core.F_member_type(g0559.Range)).Id(),0}
                          }
                          /* Let-12 */} 
                        /* ERROR PROTECTION INSERTED (v_or11-v_and6_try05689) */
                        if ErrorIn(v_or11_try056912) {v_and6_try05689 = v_or11_try056912
                        } else {
                        v_or11 = ToBoolean(OBJ(v_or11_try056912))
                        if (v_or11 == CTRUE) {v_and6_try05689 = EID{CTRUE.Id(),0}
                        } else /* or:12 */{ 
                          v_or11 = F__sup_equal_integer(C_compiler.Safety,4)
                          if (v_or11 == CTRUE) {v_and6_try05689 = EID{CTRUE.Id(),0}
                          } else /* or:13 */{ 
                            v_and6_try05689 = EID{CFALSE.Id(),0}/* org-13 */} 
                          /* org-12 */} 
                        }
                        /* or-11 */} 
                      /* Let-10 */} 
                    } else {
                    v_and6_try05689 = EID{CFALSE.Id(),0}
                    /* If-9 */} 
                  /* ERROR PROTECTION INSERTED (v_and6-g0565I_try05666) */
                  if ErrorIn(v_and6_try05689) {g0565I_try05666 = v_and6_try05689
                  } else {
                  v_and6 = ToBoolean(OBJ(v_and6_try05689))
                  if (v_and6 == CFALSE) {g0565I_try05666 = EID{CFALSE.Id(),0}
                  } else /* arg:9 */{ 
                    g0565I_try05666 = EID{CTRUE.Id(),0}/* arg-9 */} 
                  /* arg-8 */} 
                /* arg-7 */} 
              }}
              /* and-6 */} 
            /* ERROR PROTECTION INSERTED (g0565I-Result) */
            if ErrorIn(g0565I_try05666) {Result = g0565I_try05666
            } else {
            g0565I = ToBoolean(OBJ(g0565I_try05666))
            if (g0565I == CTRUE) /* If:6 */{ 
              /* Let:7 */{ 
                var x2 *ClaireAny  
                /* noccur = 1 */
                var x2_try05728 EID 
                x2_try05728 = Core.F_CALL(C_c_code,ARGS(x.ToEID(),EID{F_Compile_psort_any(Core.F_domain_I_restriction(ToRestriction(s.Id())).Id()).Id(),0}))
                /* ERROR PROTECTION INSERTED (x2-Result) */
                if ErrorIn(x2_try05728) {Result = x2_try05728
                } else {
                x2 = ANY(x2_try05728)
                /* Let:8 */{ 
                  var g0573UU *Language.Call  
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
                    g0573UU = _CL_obj
                    /* Let-9 */} 
                  Result = Core.F_CALL(C_c_code,ARGS(EID{g0573UU.Id(),0}))
                  /* Let-8 */} 
                }
                /* Let-7 */} 
              } else {
              /* Let:7 */{ 
                var g0574UU *ClaireList  
                /* noccur = 1 */
                var g0574UU_try05758 EID 
                /* Iteration:8 */{ 
                  var v_list8 *ClaireList  
                  var x *ClaireAny  
                  var v_local8 *ClaireAny  
                  v_list8 = self.Args
                  g0574UU_try05758 = EID{CreateList(ToType(CEMPTY.Id()),v_list8.Length()).Id(),0}
                  for CLcount := 0; CLcount < v_list8.Length(); CLcount++{ 
                    x = v_list8.At(CLcount)
                    var v_local8_try057610 EID 
                    v_local8_try057610 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
                    /* ERROR PROTECTION INSERTED (v_local8-g0574UU_try05758) */
                    if ErrorIn(v_local8_try057610) {g0574UU_try05758 = v_local8_try057610
                    g0574UU_try05758 = v_local8_try057610
                    break
                    } else {
                    v_local8 = ANY(v_local8_try057610)
                    ToList(OBJ(g0574UU_try05758)).PutAt(CLcount,v_local8)
                    } 
                  }
                  /* Iteration-8 */} 
                /* ERROR PROTECTION INSERTED (g0574UU-Result) */
                if ErrorIn(g0574UU_try05758) {Result = g0574UU_try05758
                } else {
                g0574UU = ToList(OBJ(g0574UU_try05758))
                Result = F_Optimize_c_code_method_method1(ToMethod(Core.F__at_property1(ToProperty(C_delete.Id()),C_property).Id()),self.Args,g0574UU)
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
      var g0577UU *Language.Call  
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
        g0577UU = _CL_obj
        /* Let-3 */} 
      Result = Core.F_CALL(C_c_code,ARGS(EID{g0577UU.Id(),0}))
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
          var _Ztype_try05785 EID 
          /* Construct:5 */{ 
            var v_bag_arg *ClaireAny  
            _Ztype_try05785= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
            var v_bag_arg_try05796 EID 
            v_bag_arg_try05796 = Core.F_CALL(C_c_type,ARGS(y.ToEID()))
            /* ERROR PROTECTION INSERTED (v_bag_arg-_Ztype_try05785) */
            if ErrorIn(v_bag_arg_try05796) {_Ztype_try05785 = v_bag_arg_try05796
            } else {
            v_bag_arg = ANY(v_bag_arg_try05796)
            ToList(OBJ(_Ztype_try05785)).AddFast(v_bag_arg)
            var v_bag_arg_try05806 EID 
            v_bag_arg_try05806 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
            /* ERROR PROTECTION INSERTED (v_bag_arg-_Ztype_try05785) */
            if ErrorIn(v_bag_arg_try05806) {_Ztype_try05785 = v_bag_arg_try05806
            } else {
            v_bag_arg = ANY(v_bag_arg_try05806)
            ToList(OBJ(_Ztype_try05785)).AddFast(v_bag_arg)}}
            /* Construct-5 */} 
          /* ERROR PROTECTION INSERTED (_Ztype-Result) */
          if ErrorIn(_Ztype_try05785) {Result = _Ztype_try05785
          } else {
          _Ztype = ToList(OBJ(_Ztype_try05785))
          if (C_set.Id() == y.Isa.Id()) /* If:5 */{ 
            /* Let:6 */{ 
              var _CL_obj *Language.Or   = Language.To_Or(new(Language.Or).Is(Language.C_Or))
              /* noccur = 9 */
              /* update:7 */{ 
                var va_arg1 *Language.Or  
                var va_arg2 *ClaireList  
                va_arg1 = _CL_obj
                var va_arg2_try05818 EID 
                /* Let:8 */{ 
                  var z_bag *ClaireList   = ToType(CEMPTY.Id()).EmptyList()
                  /* noccur = 2 */
                  /* For:9 */{ 
                    var z *ClaireAny  
                    _ = z
                    va_arg2_try05818= EID{CFALSE.Id(),0}
                    var z_support *ClaireList  
                    var z_support_try058210 EID 
                    z_support_try058210 = Core.F_enumerate_any(y)
                    /* ERROR PROTECTION INSERTED (z_support-va_arg2_try05818) */
                    if ErrorIn(z_support_try058210) {va_arg2_try05818 = z_support_try058210
                    } else {
                    z_support = ToList(OBJ(z_support_try058210))
                    z_len := z_support.Length()
                    for i_it := 0; i_it < z_len; i_it++ { 
                      z = z_support.At(i_it)
                      var void_try11 EID 
                      _ = void_try11
                      /* Let:11 */{ 
                        var g0583UU *ClaireAny  
                        /* noccur = 1 */
                        var g0583UU_try058412 EID 
                        /* Let:12 */{ 
                          var g0585UU *Language.Call  
                          /* noccur = 1 */
                          /* Let:13 */{ 
                            var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                            /* noccur = 5 */
                            _CL_obj.Selector = ToProperty(C__equal.Id())
                            _CL_obj.Args = MakeConstantList(x,z)
                            g0585UU = _CL_obj
                            /* Let-13 */} 
                          g0583UU_try058412 = Core.F_CALL(C_c_code,ARGS(EID{g0585UU.Id(),0},EID{C_any.Id(),0}))
                          /* Let-12 */} 
                        /* ERROR PROTECTION INSERTED (g0583UU-void_try11) */
                        if ErrorIn(g0583UU_try058412) {void_try11 = g0583UU_try058412
                        } else {
                        g0583UU = ANY(g0583UU_try058412)
                        void_try11 = EID{z_bag.AddFast(g0583UU).Id(),0}
                        }
                        /* Let-11 */} 
                      /* ERROR PROTECTION INSERTED (void_try11-va_arg2_try05818) */
                      if ErrorIn(void_try11) {va_arg2_try05818 = void_try11
                      va_arg2_try05818 = void_try11
                      break
                      } else {
                      }}
                      /* loop-10 */} 
                    /* For-9 */} 
                  /* ERROR PROTECTION INSERTED (va_arg2_try05818-va_arg2_try05818) */
                  if !ErrorIn(va_arg2_try05818) {
                  va_arg2_try05818 = EID{z_bag.Id(),0}
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                if ErrorIn(va_arg2_try05818) {Result = va_arg2_try05818
                } else {
                va_arg2 = ToList(OBJ(va_arg2_try05818))
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
            var t_try05886 EID 
            t_try05886 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
            /* ERROR PROTECTION INSERTED (t-Result) */
            if ErrorIn(t_try05886) {Result = t_try05886
            } else {
            t = ANY(t_try05886)
            /* Let:6 */{ 
              var mt *ClaireType   = Core.F_member_type(ToType(t))
              /* noccur = 3 */
              /* Let:7 */{ 
                var r *ClaireAny  
                /* noccur = 2 */
                var r_try05898 EID 
                /* Let:8 */{ 
                  var g0590UU *ClaireList  
                  /* noccur = 1 */
                  var g0590UU_try05919 EID 
                  /* Iteration:9 */{ 
                    var v_list9 *ClaireList  
                    var u *ClaireAny  
                    var v_local9 *ClaireAny  
                    v_list9 = l
                    g0590UU_try05919 = EID{CreateList(ToType(CEMPTY.Id()),v_list9.Length()).Id(),0}
                    for CLcount := 0; CLcount < v_list9.Length(); CLcount++{ 
                      u = v_list9.At(CLcount)
                      var v_local9_try059211 EID 
                      v_local9_try059211 = Core.F_CALL(C_c_type,ARGS(u.ToEID()))
                      /* ERROR PROTECTION INSERTED (v_local9-g0590UU_try05919) */
                      if ErrorIn(v_local9_try059211) {g0590UU_try05919 = v_local9_try059211
                      g0590UU_try05919 = v_local9_try059211
                      break
                      } else {
                      v_local9 = ANY(v_local9_try059211)
                      ToList(OBJ(g0590UU_try05919)).PutAt(CLcount,v_local9)
                      } 
                    }
                    /* Iteration-9 */} 
                  /* ERROR PROTECTION INSERTED (g0590UU-r_try05898) */
                  if ErrorIn(g0590UU_try05919) {r_try05898 = g0590UU_try05919
                  } else {
                  g0590UU = ToList(OBJ(g0590UU_try05919))
                  r_try05898 = F_Optimize_restriction_I_property(p,g0590UU,CTRUE).ToEID()
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (r-Result) */
                if ErrorIn(r_try05898) {Result = r_try05898
                } else {
                r = ANY(r_try05898)
                if (C_OPT.ToRemove.Contain_ask(x) == CTRUE) /* If:8 */{ 
                  Result = EID{CNIL.Id(),0}
                  } else {
                  var g0593I *ClaireBoolean  
                  var g0593I_try05949 EID 
                  /* and:9 */{ 
                    var v_and9 *ClaireBoolean  
                    
                    if (C_table.Id() == x.Isa.Id()) /* If:10 */{ 
                      /* Let:11 */{ 
                        var g0586 *ClaireTable   = ToTable(x)
                        /* noccur = 1 */
                        v_and9 = Equal(C_integer.Id(),g0586.Params.Isa.Id())
                        /* Let-11 */} 
                      } else {
                      v_and9 = CFALSE
                      /* If-10 */} 
                    if (v_and9 == CFALSE) {g0593I_try05949 = EID{CFALSE.Id(),0}
                    } else /* arg:10 */{ 
                      var v_and9_try059511 EID 
                      /* or:11 */{ 
                        var v_or11 *ClaireBoolean  
                        
                        var v_or11_try059612 EID 
                        /* Let:12 */{ 
                          var g0597UU *ClaireType  
                          /* noccur = 1 */
                          var g0597UU_try059813 EID 
                          g0597UU_try059813 = Core.F_CALL(C_c_type,ARGS(l.At(2-1).ToEID()))
                          /* ERROR PROTECTION INSERTED (g0597UU-v_or11_try059612) */
                          if ErrorIn(g0597UU_try059813) {v_or11_try059612 = g0597UU_try059813
                          } else {
                          g0597UU = ToType(OBJ(g0597UU_try059813))
                          v_or11_try059612 = EID{g0597UU.Included(ToType(OBJ(Core.F_CALL(C_domain,ARGS(x.ToEID()))))).Id(),0}
                          }
                          /* Let-12 */} 
                        /* ERROR PROTECTION INSERTED (v_or11-v_and9_try059511) */
                        if ErrorIn(v_or11_try059612) {v_and9_try059511 = v_or11_try059612
                        } else {
                        v_or11 = ToBoolean(OBJ(v_or11_try059612))
                        if (v_or11 == CTRUE) {v_and9_try059511 = EID{CTRUE.Id(),0}
                        } else /* or:12 */{ 
                          v_or11 = MakeBoolean((p.Id() == C_nth.Id()) && (C_compiler.Safety > 2))
                          if (v_or11 == CTRUE) {v_and9_try059511 = EID{CTRUE.Id(),0}
                          } else /* or:13 */{ 
                            v_and9_try059511 = EID{CFALSE.Id(),0}/* org-13 */} 
                          /* org-12 */} 
                        }
                        /* or-11 */} 
                      /* ERROR PROTECTION INSERTED (v_and9-g0593I_try05949) */
                      if ErrorIn(v_and9_try059511) {g0593I_try05949 = v_and9_try059511
                      } else {
                      v_and9 = ToBoolean(OBJ(v_and9_try059511))
                      if (v_and9 == CFALSE) {g0593I_try05949 = EID{CFALSE.Id(),0}
                      } else /* arg:11 */{ 
                        g0593I_try05949 = EID{CTRUE.Id(),0}/* arg-11 */} 
                      /* arg-10 */} 
                    }
                    /* and-9 */} 
                  /* ERROR PROTECTION INSERTED (g0593I-Result) */
                  if ErrorIn(g0593I_try05949) {Result = g0593I_try05949
                  } else {
                  g0593I = ToBoolean(OBJ(g0593I_try05949))
                  if (g0593I == CTRUE) /* If:9 */{ 
                    /* Let:10 */{ 
                      var _CL_obj *Language.CallTable   = Language.To_CallTable(new(Language.CallTable).Is(Language.C_Call_table))
                      /* noccur = 7 */
                      _CL_obj.Selector = ToTable(x)
                      /* update:11 */{ 
                        var va_arg1 *Language.CallTable  
                        var va_arg2 *ClaireAny  
                        va_arg1 = _CL_obj
                        var va_arg2_try059912 EID 
                        va_arg2_try059912 = Core.F_CALL(C_c_code,ARGS(l.At(2-1).ToEID(),EID{C_integer.Id(),0}))
                        /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                        if ErrorIn(va_arg2_try059912) {Result = va_arg2_try059912
                        } else {
                        va_arg2 = ANY(va_arg2_try059912)
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
                        var va_arg2_try060012 EID 
                        /* Let:12 */{ 
                          var g0601UU *ClaireBoolean  
                          /* noccur = 1 */
                          var g0601UU_try060213 EID 
                          /* or:13 */{ 
                            var v_or13 *ClaireBoolean  
                            
                            var v_or13_try060314 EID 
                            v_or13_try060314 = Core.F_BELONG(ANY(Core.F_CALL(C_default,ARGS(x.ToEID()))),ANY(Core.F_CALL(C_range,ARGS(x.ToEID()))))
                            /* ERROR PROTECTION INSERTED (v_or13-g0601UU_try060213) */
                            if ErrorIn(v_or13_try060314) {g0601UU_try060213 = v_or13_try060314
                            } else {
                            v_or13 = ToBoolean(OBJ(v_or13_try060314))
                            if (v_or13 == CTRUE) {g0601UU_try060213 = EID{CTRUE.Id(),0}
                            } else /* or:14 */{ 
                              v_or13 = Equal(ANY(Core.F_CALL(C_default,ARGS(x.ToEID()))),MakeInteger(0).Id())
                              if (v_or13 == CTRUE) {g0601UU_try060213 = EID{CTRUE.Id(),0}
                              } else /* or:15 */{ 
                                v_or13 = Equal(p.Id(),C_get.Id())
                                if (v_or13 == CTRUE) {g0601UU_try060213 = EID{CTRUE.Id(),0}
                                } else /* or:16 */{ 
                                  g0601UU_try060213 = EID{CFALSE.Id(),0}/* org-16 */} 
                                /* org-15 */} 
                              /* org-14 */} 
                            }
                            /* or-13 */} 
                          /* ERROR PROTECTION INSERTED (g0601UU-va_arg2_try060012) */
                          if ErrorIn(g0601UU_try060213) {va_arg2_try060012 = g0601UU_try060213
                          } else {
                          g0601UU = ToBoolean(OBJ(g0601UU_try060213))
                          va_arg2_try060012 = EID{g0601UU.Not.Id(),0}
                          }
                          /* Let-12 */} 
                        /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                        if ErrorIn(va_arg2_try060012) {Result = va_arg2_try060012
                        } else {
                        va_arg2 = ToBoolean(OBJ(va_arg2_try060012))
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
                    var g0604I *ClaireBoolean  
                    var g0604I_try060510 EID 
                    /* and:10 */{ 
                      var v_and10 *ClaireBoolean  
                      
                      if (C_table.Id() == x.Isa.Id()) /* If:11 */{ 
                        /* Let:12 */{ 
                          var g0587 *ClaireTable   = ToTable(x)
                          /* noccur = 1 */
                          v_and10 = g0587.Params.Isa.IsIn(C_list)
                          /* Let-12 */} 
                        } else {
                        v_and10 = CFALSE
                        /* If-11 */} 
                      if (v_and10 == CFALSE) {g0604I_try060510 = EID{CFALSE.Id(),0}
                      } else /* arg:11 */{ 
                        v_and10 = Equal(MakeInteger(l.Length()).Id(),MakeInteger(3).Id())
                        if (v_and10 == CFALSE) {g0604I_try060510 = EID{CFALSE.Id(),0}
                        } else /* arg:12 */{ 
                          var v_and10_try060613 EID 
                          /* or:13 */{ 
                            var v_or13 *ClaireBoolean  
                            
                            var v_or13_try060714 EID 
                            /* Let:14 */{ 
                              var g0608UU *ClaireTuple  
                              /* noccur = 1 */
                              var g0608UU_try060915 EID 
                              /* Let:15 */{ 
                                var g0610UU *ClaireList  
                                /* noccur = 1 */
                                var g0610UU_try061116 EID 
                                /* Construct:16 */{ 
                                  var v_bag_arg *ClaireAny  
                                  g0610UU_try061116= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                                  var v_bag_arg_try061217 EID 
                                  v_bag_arg_try061217 = Core.F_CALL(C_c_type,ARGS(l.At(2-1).ToEID()))
                                  /* ERROR PROTECTION INSERTED (v_bag_arg-g0610UU_try061116) */
                                  if ErrorIn(v_bag_arg_try061217) {g0610UU_try061116 = v_bag_arg_try061217
                                  } else {
                                  v_bag_arg = ANY(v_bag_arg_try061217)
                                  ToList(OBJ(g0610UU_try061116)).AddFast(v_bag_arg)
                                  var v_bag_arg_try061317 EID 
                                  v_bag_arg_try061317 = Core.F_CALL(C_c_type,ARGS(l.At(3-1).ToEID()))
                                  /* ERROR PROTECTION INSERTED (v_bag_arg-g0610UU_try061116) */
                                  if ErrorIn(v_bag_arg_try061317) {g0610UU_try061116 = v_bag_arg_try061317
                                  } else {
                                  v_bag_arg = ANY(v_bag_arg_try061317)
                                  ToList(OBJ(g0610UU_try061116)).AddFast(v_bag_arg)}}
                                  /* Construct-16 */} 
                                /* ERROR PROTECTION INSERTED (g0610UU-g0608UU_try060915) */
                                if ErrorIn(g0610UU_try061116) {g0608UU_try060915 = g0610UU_try061116
                                } else {
                                g0610UU = ToList(OBJ(g0610UU_try061116))
                                g0608UU_try060915 = EID{g0610UU.Tuple_I().Id(),0}
                                }
                                /* Let-15 */} 
                              /* ERROR PROTECTION INSERTED (g0608UU-v_or13_try060714) */
                              if ErrorIn(g0608UU_try060915) {v_or13_try060714 = g0608UU_try060915
                              } else {
                              g0608UU = ToTuple(OBJ(g0608UU_try060915))
                              v_or13_try060714 = EID{ToType(g0608UU.Id()).Included(ToType(OBJ(Core.F_CALL(C_domain,ARGS(x.ToEID()))))).Id(),0}
                              }
                              /* Let-14 */} 
                            /* ERROR PROTECTION INSERTED (v_or13-v_and10_try060613) */
                            if ErrorIn(v_or13_try060714) {v_and10_try060613 = v_or13_try060714
                            } else {
                            v_or13 = ToBoolean(OBJ(v_or13_try060714))
                            if (v_or13 == CTRUE) {v_and10_try060613 = EID{CTRUE.Id(),0}
                            } else /* or:14 */{ 
                              v_or13 = Core.F__sup_integer(C_compiler.Safety,2)
                              if (v_or13 == CTRUE) {v_and10_try060613 = EID{CTRUE.Id(),0}
                              } else /* or:15 */{ 
                                v_and10_try060613 = EID{CFALSE.Id(),0}/* org-15 */} 
                              /* org-14 */} 
                            }
                            /* or-13 */} 
                          /* ERROR PROTECTION INSERTED (v_and10-g0604I_try060510) */
                          if ErrorIn(v_and10_try060613) {g0604I_try060510 = v_and10_try060613
                          } else {
                          v_and10 = ToBoolean(OBJ(v_and10_try060613))
                          if (v_and10 == CFALSE) {g0604I_try060510 = EID{CFALSE.Id(),0}
                          } else /* arg:13 */{ 
                            g0604I_try060510 = EID{CTRUE.Id(),0}/* arg-13 */} 
                          /* arg-12 */} 
                        /* arg-11 */} 
                      }
                      /* and-10 */} 
                    /* ERROR PROTECTION INSERTED (g0604I-Result) */
                    if ErrorIn(g0604I_try060510) {Result = g0604I_try060510
                    } else {
                    g0604I = ToBoolean(OBJ(g0604I_try060510))
                    if (g0604I == CTRUE) /* If:10 */{ 
                      /* Let:11 */{ 
                        var _CL_obj *Language.CallTable   = Language.To_CallTable(new(Language.CallTable).Is(Language.C_Call_table))
                        /* noccur = 11 */
                        _CL_obj.Selector = ToTable(x)
                        /* update:12 */{ 
                          var va_arg1 *Language.CallTable  
                          var va_arg2 *ClaireAny  
                          va_arg1 = _CL_obj
                          var va_arg2_try061413 EID 
                          /* Let:13 */{ 
                            var _CL_obj *Language.List   = Language.To_List(new(Language.List).Is(Language.C_List))
                            /* noccur = 3 */
                            /* update:14 */{ 
                              var va_arg1 *Language.Construct  
                              var va_arg2 *ClaireList  
                              va_arg1 = Language.To_Construct(_CL_obj.Id())
                              var va_arg2_try061515 EID 
                              /* Construct:15 */{ 
                                var v_bag_arg *ClaireAny  
                                va_arg2_try061515= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                                var v_bag_arg_try061616 EID 
                                v_bag_arg_try061616 = Core.F_CALL(C_c_code,ARGS(l.At(2-1).ToEID(),EID{C_integer.Id(),0}))
                                /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try061515) */
                                if ErrorIn(v_bag_arg_try061616) {va_arg2_try061515 = v_bag_arg_try061616
                                } else {
                                v_bag_arg = ANY(v_bag_arg_try061616)
                                ToList(OBJ(va_arg2_try061515)).AddFast(v_bag_arg)
                                var v_bag_arg_try061716 EID 
                                v_bag_arg_try061716 = Core.F_CALL(C_c_code,ARGS(l.At(3-1).ToEID(),EID{C_integer.Id(),0}))
                                /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try061515) */
                                if ErrorIn(v_bag_arg_try061716) {va_arg2_try061515 = v_bag_arg_try061716
                                } else {
                                v_bag_arg = ANY(v_bag_arg_try061716)
                                ToList(OBJ(va_arg2_try061515)).AddFast(v_bag_arg)}}
                                /* Construct-15 */} 
                              /* ERROR PROTECTION INSERTED (va_arg2-va_arg2_try061413) */
                              if ErrorIn(va_arg2_try061515) {va_arg2_try061413 = va_arg2_try061515
                              } else {
                              va_arg2 = ToList(OBJ(va_arg2_try061515))
                              /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                              va_arg1.Args = va_arg2
                              va_arg2_try061413 = EID{va_arg2.Id(),0}
                              }
                              /* update-14 */} 
                            /* ERROR PROTECTION INSERTED (va_arg2_try061413-va_arg2_try061413) */
                            if !ErrorIn(va_arg2_try061413) {
                            va_arg2_try061413 = EID{_CL_obj.Id(),0}
                            }
                            /* Let-13 */} 
                          /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                          if ErrorIn(va_arg2_try061413) {Result = va_arg2_try061413
                          } else {
                          va_arg2 = ANY(va_arg2_try061413)
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
                          var va_arg2_try061813 EID 
                          /* Let:13 */{ 
                            var g0619UU *ClaireBoolean  
                            /* noccur = 1 */
                            var g0619UU_try062014 EID 
                            /* or:14 */{ 
                              var v_or14 *ClaireBoolean  
                              
                              var v_or14_try062115 EID 
                              v_or14_try062115 = Core.F_BELONG(ANY(Core.F_CALL(C_default,ARGS(x.ToEID()))),ANY(Core.F_CALL(C_range,ARGS(x.ToEID()))))
                              /* ERROR PROTECTION INSERTED (v_or14-g0619UU_try062014) */
                              if ErrorIn(v_or14_try062115) {g0619UU_try062014 = v_or14_try062115
                              } else {
                              v_or14 = ToBoolean(OBJ(v_or14_try062115))
                              if (v_or14 == CTRUE) {g0619UU_try062014 = EID{CTRUE.Id(),0}
                              } else /* or:15 */{ 
                                v_or14 = Equal(ANY(Core.F_CALL(C_default,ARGS(x.ToEID()))),MakeInteger(0).Id())
                                if (v_or14 == CTRUE) {g0619UU_try062014 = EID{CTRUE.Id(),0}
                                } else /* or:16 */{ 
                                  v_or14 = Equal(p.Id(),C_get.Id())
                                  if (v_or14 == CTRUE) {g0619UU_try062014 = EID{CTRUE.Id(),0}
                                  } else /* or:17 */{ 
                                    g0619UU_try062014 = EID{CFALSE.Id(),0}/* org-17 */} 
                                  /* org-16 */} 
                                /* org-15 */} 
                              }
                              /* or-14 */} 
                            /* ERROR PROTECTION INSERTED (g0619UU-va_arg2_try061813) */
                            if ErrorIn(g0619UU_try062014) {va_arg2_try061813 = g0619UU_try062014
                            } else {
                            g0619UU = ToBoolean(OBJ(g0619UU_try062014))
                            va_arg2_try061813 = EID{g0619UU.Not.Id(),0}
                            }
                            /* Let-13 */} 
                          /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                          if ErrorIn(va_arg2_try061813) {Result = va_arg2_try061813
                          } else {
                          va_arg2 = ToBoolean(OBJ(va_arg2_try061813))
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
                        var g0622UU *ClaireAny  
                        /* noccur = 1 */
                        var g0622UU_try062412 EID 
                        g0622UU_try062412 = Core.F_CALL(C_c_code,ARGS(x.ToEID(),EID{C_array.Id(),0}))
                        /* ERROR PROTECTION INSERTED (g0622UU-Result) */
                        if ErrorIn(g0622UU_try062412) {Result = g0622UU_try062412
                        } else {
                        g0622UU = ANY(g0622UU_try062412)
                        /* Let:12 */{ 
                          var g0623UU *ClaireAny  
                          /* noccur = 1 */
                          var g0623UU_try062513 EID 
                          g0623UU_try062513 = Core.F_CALL(C_c_code,ARGS(l.At(2-1).ToEID(),EID{C_integer.Id(),0}))
                          /* ERROR PROTECTION INSERTED (g0623UU-Result) */
                          if ErrorIn(g0623UU_try062513) {Result = g0623UU_try062513
                          } else {
                          g0623UU = ANY(g0623UU_try062513)
                          Result = Language.C_Call_array.Make(g0622UU,g0623UU,mt.Id()).ToEID()
                          }
                          /* Let-12 */} 
                        }
                        /* Let-11 */} 
                      /* If!10 */}  else if (C_method.Id() == r.Isa.Id()) /* If:10 */{ 
                      if ((C_compiler.Optimize_ask == CTRUE) && 
                          ((ToType(t).Included(ToType(C_array.Id())) == CTRUE) || 
                              (ToType(t).Included(ToType(C_table.Id())) == CTRUE))) /* If:11 */{ 
                        F_Compile_notice_void()
                        
                        /* If-11 */} 
                      /* Let:11 */{ 
                        var g0626UU *ClaireList  
                        /* noccur = 1 */
                        var g0626UU_try062712 EID 
                        /* Iteration:12 */{ 
                          var v_list12 *ClaireList  
                          var x *ClaireAny  
                          var v_local12 *ClaireAny  
                          v_list12 = self.Args
                          g0626UU_try062712 = EID{CreateList(ToType(CEMPTY.Id()),v_list12.Length()).Id(),0}
                          for CLcount := 0; CLcount < v_list12.Length(); CLcount++{ 
                            x = v_list12.At(CLcount)
                            var v_local12_try062814 EID 
                            v_local12_try062814 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
                            /* ERROR PROTECTION INSERTED (v_local12-g0626UU_try062712) */
                            if ErrorIn(v_local12_try062814) {g0626UU_try062712 = v_local12_try062814
                            g0626UU_try062712 = v_local12_try062814
                            break
                            } else {
                            v_local12 = ANY(v_local12_try062814)
                            ToList(OBJ(g0626UU_try062712)).PutAt(CLcount,v_local12)
                            } 
                          }
                          /* Iteration-12 */} 
                        /* ERROR PROTECTION INSERTED (g0626UU-Result) */
                        if ErrorIn(g0626UU_try062712) {Result = g0626UU_try062712
                        } else {
                        g0626UU = ToList(OBJ(g0626UU_try062712))
                        Result = F_Optimize_c_code_method_method1(ToMethod(r),self.Args,g0626UU)
                        }
                        /* Let-11 */} 
                      } else {
                      /* Let:11 */{ 
                        var g0629UU *ClaireList  
                        /* noccur = 1 */
                        var g0629UU_try063012 EID 
                        /* Iteration:12 */{ 
                          var v_list12 *ClaireList  
                          var x *ClaireAny  
                          var v_local12 *ClaireAny  
                          v_list12 = self.Args
                          g0629UU_try063012 = EID{CreateList(ToType(CEMPTY.Id()),v_list12.Length()).Id(),0}
                          for CLcount := 0; CLcount < v_list12.Length(); CLcount++{ 
                            x = v_list12.At(CLcount)
                            var v_local12_try063114 EID 
                            v_local12_try063114 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
                            /* ERROR PROTECTION INSERTED (v_local12-g0629UU_try063012) */
                            if ErrorIn(v_local12_try063114) {g0629UU_try063012 = v_local12_try063114
                            g0629UU_try063012 = v_local12_try063114
                            break
                            } else {
                            v_local12 = ANY(v_local12_try063114)
                            ToList(OBJ(g0629UU_try063012)).PutAt(CLcount,v_local12)
                            } 
                          }
                          /* Iteration-12 */} 
                        /* ERROR PROTECTION INSERTED (g0629UU-Result) */
                        if ErrorIn(g0629UU_try063012) {Result = g0629UU_try063012
                        } else {
                        g0629UU = ToList(OBJ(g0629UU_try063012))
                        Result = F_Optimize_c_warn_property(p,self.Args,g0629UU)
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
              var g0632I *ClaireBoolean  
              var g0632I_try06337 EID 
              /* or:7 */{ 
                var v_or7 *ClaireBoolean  
                
                v_or7 = Equal(sp.Id(),C_put.Id())
                if (v_or7 == CTRUE) {g0632I_try06337 = EID{CTRUE.Id(),0}
                } else /* or:8 */{ 
                  var v_or7_try06349 EID 
                  /* and:9 */{ 
                    var v_and9 *ClaireBoolean  
                    
                    var v_and9_try063510 EID 
                    /* or:10 */{ 
                      var v_or10 *ClaireBoolean  
                      
                      var v_or10_try063611 EID 
                      /* Let:11 */{ 
                        var g0637UU *ClaireType  
                        /* noccur = 1 */
                        var g0637UU_try063812 EID 
                        g0637UU_try063812 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
                        /* ERROR PROTECTION INSERTED (g0637UU-v_or10_try063611) */
                        if ErrorIn(g0637UU_try063812) {v_or10_try063611 = g0637UU_try063812
                        } else {
                        g0637UU = ToType(OBJ(g0637UU_try063812))
                        v_or10_try063611 = EID{g0637UU.Included(p.Domain).Id(),0}
                        }
                        /* Let-11 */} 
                      /* ERROR PROTECTION INSERTED (v_or10-v_and9_try063510) */
                      if ErrorIn(v_or10_try063611) {v_and9_try063510 = v_or10_try063611
                      } else {
                      v_or10 = ToBoolean(OBJ(v_or10_try063611))
                      if (v_or10 == CTRUE) {v_and9_try063510 = EID{CTRUE.Id(),0}
                      } else /* or:11 */{ 
                        v_or10 = F__sup_equal_integer(C_compiler.Safety,5)
                        if (v_or10 == CTRUE) {v_and9_try063510 = EID{CTRUE.Id(),0}
                        } else /* or:12 */{ 
                          v_and9_try063510 = EID{CFALSE.Id(),0}/* org-12 */} 
                        /* org-11 */} 
                      }
                      /* or-10 */} 
                    /* ERROR PROTECTION INSERTED (v_and9-v_or7_try06349) */
                    if ErrorIn(v_and9_try063510) {v_or7_try06349 = v_and9_try063510
                    } else {
                    v_and9 = ToBoolean(OBJ(v_and9_try063510))
                    if (v_and9 == CFALSE) {v_or7_try06349 = EID{CFALSE.Id(),0}
                    } else /* arg:10 */{ 
                      var v_and9_try063911 EID 
                      /* or:11 */{ 
                        var v_or11 *ClaireBoolean  
                        
                        var v_or11_try064012 EID 
                        /* Let:12 */{ 
                          var g0641UU *ClaireType  
                          /* noccur = 1 */
                          var g0641UU_try064213 EID 
                          g0641UU_try064213 = Core.F_CALL(C_c_type,ARGS(y.ToEID()))
                          /* ERROR PROTECTION INSERTED (g0641UU-v_or11_try064012) */
                          if ErrorIn(g0641UU_try064213) {v_or11_try064012 = g0641UU_try064213
                          } else {
                          g0641UU = ToType(OBJ(g0641UU_try064213))
                          v_or11_try064012 = EID{g0641UU.Included(p.Range).Id(),0}
                          }
                          /* Let-12 */} 
                        /* ERROR PROTECTION INSERTED (v_or11-v_and9_try063911) */
                        if ErrorIn(v_or11_try064012) {v_and9_try063911 = v_or11_try064012
                        } else {
                        v_or11 = ToBoolean(OBJ(v_or11_try064012))
                        if (v_or11 == CTRUE) {v_and9_try063911 = EID{CTRUE.Id(),0}
                        } else /* or:12 */{ 
                          v_or11 = F__sup_equal_integer(C_compiler.Safety,4)
                          if (v_or11 == CTRUE) {v_and9_try063911 = EID{CTRUE.Id(),0}
                          } else /* or:13 */{ 
                            v_and9_try063911 = EID{CFALSE.Id(),0}/* org-13 */} 
                          /* org-12 */} 
                        }
                        /* or-11 */} 
                      /* ERROR PROTECTION INSERTED (v_and9-v_or7_try06349) */
                      if ErrorIn(v_and9_try063911) {v_or7_try06349 = v_and9_try063911
                      } else {
                      v_and9 = ToBoolean(OBJ(v_and9_try063911))
                      if (v_and9 == CFALSE) {v_or7_try06349 = EID{CFALSE.Id(),0}
                      } else /* arg:11 */{ 
                        v_or7_try06349 = EID{CTRUE.Id(),0}/* arg-11 */} 
                      /* arg-10 */} 
                    }}
                    /* and-9 */} 
                  /* ERROR PROTECTION INSERTED (v_or7-g0632I_try06337) */
                  if ErrorIn(v_or7_try06349) {g0632I_try06337 = v_or7_try06349
                  } else {
                  v_or7 = ToBoolean(OBJ(v_or7_try06349))
                  if (v_or7 == CTRUE) {g0632I_try06337 = EID{CTRUE.Id(),0}
                  } else /* or:9 */{ 
                    g0632I_try06337 = EID{CFALSE.Id(),0}/* org-9 */} 
                  /* org-8 */} 
                }
                /* or-7 */} 
              /* ERROR PROTECTION INSERTED (g0632I-Result) */
              if ErrorIn(g0632I_try06337) {Result = g0632I_try06337
              } else {
              g0632I = ToBoolean(OBJ(g0632I_try06337))
              if (g0632I == CTRUE) /* If:7 */{ 
                var g0643I *ClaireBoolean  
                var g0643I_try06448 EID 
                /* and:8 */{ 
                  var v_and8 *ClaireBoolean  
                  
                  var v_and8_try06459 EID 
                  v_and8_try06459 = F_Optimize_Update_ask_relation1(ToRelation(p.Id()),x,y)
                  /* ERROR PROTECTION INSERTED (v_and8-g0643I_try06448) */
                  if ErrorIn(v_and8_try06459) {g0643I_try06448 = v_and8_try06459
                  } else {
                  v_and8 = ToBoolean(OBJ(v_and8_try06459))
                  if (v_and8 == CFALSE) {g0643I_try06448 = EID{CFALSE.Id(),0}
                  } else /* arg:9 */{ 
                    v_and8 = MakeBoolean((p.Params.Isa.IsIn(C_list) == CTRUE) || (C_integer.Id() == p.Params.Isa.Id()))
                    if (v_and8 == CFALSE) {g0643I_try06448 = EID{CFALSE.Id(),0}
                    } else /* arg:10 */{ 
                      g0643I_try06448 = EID{CTRUE.Id(),0}/* arg-10 */} 
                    /* arg-9 */} 
                  }
                  /* and-8 */} 
                /* ERROR PROTECTION INSERTED (g0643I-Result) */
                if ErrorIn(g0643I_try06448) {Result = g0643I_try06448
                } else {
                g0643I = ToBoolean(OBJ(g0643I_try06448))
                if (g0643I == CTRUE) /* If:8 */{ 
                  /* Let:9 */{ 
                    var _Zx *ClaireAny  
                    /* noccur = 2 */
                    var _Zx_try064610 EID 
                    _Zx_try064610 = Core.F_CALL(C_c_code,ARGS(x.ToEID(),EID{C_any.Id(),0}))
                    /* ERROR PROTECTION INSERTED (_Zx-Result) */
                    if ErrorIn(_Zx_try064610) {Result = _Zx_try064610
                    } else {
                    _Zx = ANY(_Zx_try064610)
                    /* Let:10 */{ 
                      var _Zy *ClaireAny  
                      /* noccur = 1 */
                      var _Zy_try064711 EID 
                      _Zy_try064711 = Core.F_CALL(C_c_code,ARGS(y.ToEID(),EID{C_any.Id(),0}))
                      /* ERROR PROTECTION INSERTED (_Zy-Result) */
                      if ErrorIn(_Zy_try064711) {Result = _Zy_try064711
                      } else {
                      _Zy = ANY(_Zy_try064711)
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
          var tp_try06485 EID 
          tp_try06485 = Core.F_CALL(C_c_type,ARGS(p.ToEID()))
          /* ERROR PROTECTION INSERTED (tp-Result) */
          if ErrorIn(tp_try06485) {Result = tp_try06485
          } else {
          tp = ANY(tp_try06485)
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
                  var typeok_try06499 EID 
                  /* or:9 */{ 
                    var v_or9 *ClaireBoolean  
                    
                    var v_or9_try065010 EID 
                    /* Let:10 */{ 
                      var g0651UU *ClaireType  
                      /* noccur = 1 */
                      var g0651UU_try065211 EID 
                      g0651UU_try065211 = Core.F_CALL(C_c_type,ARGS(y.ToEID()))
                      /* ERROR PROTECTION INSERTED (g0651UU-v_or9_try065010) */
                      if ErrorIn(g0651UU_try065211) {v_or9_try065010 = g0651UU_try065211
                      } else {
                      g0651UU = ToType(OBJ(g0651UU_try065211))
                      v_or9_try065010 = EID{g0651UU.Included(Core.F_member_type(ToType(tp))).Id(),0}
                      }
                      /* Let-10 */} 
                    /* ERROR PROTECTION INSERTED (v_or9-typeok_try06499) */
                    if ErrorIn(v_or9_try065010) {typeok_try06499 = v_or9_try065010
                    } else {
                    v_or9 = ToBoolean(OBJ(v_or9_try065010))
                    if (v_or9 == CTRUE) {typeok_try06499 = EID{CTRUE.Id(),0}
                    } else /* or:10 */{ 
                      v_or9 = F__sup_equal_integer(C_compiler.Safety,4)
                      if (v_or9 == CTRUE) {typeok_try06499 = EID{CTRUE.Id(),0}
                      } else /* or:11 */{ 
                        typeok_try06499 = EID{CFALSE.Id(),0}/* org-11 */} 
                      /* org-10 */} 
                    }
                    /* or-9 */} 
                  /* ERROR PROTECTION INSERTED (typeok-Result) */
                  if ErrorIn(typeok_try06499) {Result = typeok_try06499
                  } else {
                  typeok = ToBoolean(OBJ(typeok_try06499))
                  if (((sp.Id() == C_nth_put.Id()) || 
                        (typeok == CTRUE)) && 
                      ((mt.Included(ToType(C_float.Id())) == CTRUE) || 
                          (Equal(Core.F__exp_type(mt,ToType(C_float.Id())).Id(),CEMPTY.Id()) == CTRUE))) /* If:9 */{ 
                    /* Let:10 */{ 
                      var _Zx *ClaireAny  
                      /* noccur = 1 */
                      var _Zx_try065311 EID 
                      _Zx_try065311 = Core.F_CALL(C_c_code,ARGS(x.ToEID(),EID{C_integer.Id(),0}))
                      /* ERROR PROTECTION INSERTED (_Zx-Result) */
                      if ErrorIn(_Zx_try065311) {Result = _Zx_try065311
                      } else {
                      _Zx = ANY(_Zx_try065311)
                      /* Let:11 */{ 
                        var _Zy *ClaireAny  
                        /* noccur = 1 */
                        var _Zy_try065412 EID 
                        /* Let:12 */{ 
                          var g0655UU *ClaireClass  
                          /* noccur = 1 */
                          if (mt.Included(ToType(C_float.Id())) == CTRUE) /* If:13 */{ 
                            g0655UU = C_float
                            } else {
                            g0655UU = C_any
                            /* If-13 */} 
                          _Zy_try065412 = Core.F_CALL(C_c_code,ARGS(y.ToEID(),EID{g0655UU.Id(),0}))
                          /* Let-12 */} 
                        /* ERROR PROTECTION INSERTED (_Zy-Result) */
                        if ErrorIn(_Zy_try065412) {Result = _Zy_try065412
                        } else {
                        _Zy = ANY(_Zy_try065412)
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
                            var va_arg2_try065614 EID 
                            /* Let:14 */{ 
                              var g0657UU *ClaireAny  
                              /* noccur = 1 */
                              var g0657UU_try065815 EID 
                              g0657UU_try065815 = Core.F_CALL(C_c_code,ARGS(p.ToEID(),EID{C_array.Id(),0}))
                              /* ERROR PROTECTION INSERTED (g0657UU-va_arg2_try065614) */
                              if ErrorIn(g0657UU_try065815) {va_arg2_try065614 = g0657UU_try065815
                              } else {
                              g0657UU = ANY(g0657UU_try065815)
                              va_arg2_try065614 = Language.C_Call_array.Make(g0657UU,_Zx,mt.Id()).ToEID()
                              }
                              /* Let-14 */} 
                            /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                            if ErrorIn(va_arg2_try065614) {Result = va_arg2_try065614
                            } else {
                            va_arg2 = ANY(va_arg2_try065614)
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
        var v_and2_try06614 EID 
        /* or:4 */{ 
          var v_or4 *ClaireBoolean  
          
          v_or4 = MakeBoolean((p.IfWrite != CNULL) && (p.IfWrite.Isa.IsIn(C_list) != CTRUE))
          if (v_or4 == CTRUE) {v_and2_try06614 = EID{CTRUE.Id(),0}
          } else /* or:5 */{ 
            var v_or4_try06626 EID 
            /* and:6 */{ 
              var v_and6 *ClaireBoolean  
              
              v_and6 = MakeBoolean((p.Inverse.Id() == CNULL))
              if (v_and6 == CFALSE) {v_or4_try06626 = EID{CFALSE.Id(),0}
              } else /* arg:7 */{ 
                if (C_table.Id() == p.Isa.Id()) /* If:8 */{ 
                  /* Let:9 */{ 
                    var g0659 *ClaireTable   = ToTable(p.Id())
                    /* noccur = 1 */
                    v_and6 = Equal(C_integer.Id(),g0659.Params.Isa.Id())
                    /* Let-9 */} 
                  } else {
                  v_and6 = CTRUE
                  /* If-8 */} 
                if (v_and6 == CFALSE) {v_or4_try06626 = EID{CFALSE.Id(),0}
                } else /* arg:8 */{ 
                  var v_and6_try06639 EID 
                  if (p.Store_ask == CTRUE) /* If:9 */{ 
                    /* and:10 */{ 
                      var v_and10 *ClaireBoolean  
                      
                      var v_and10_try066411 EID 
                      v_and10_try066411 = F_Compile_designated_ask_any(x)
                      /* ERROR PROTECTION INSERTED (v_and10-v_and6_try06639) */
                      if ErrorIn(v_and10_try066411) {v_and6_try06639 = v_and10_try066411
                      } else {
                      v_and10 = ToBoolean(OBJ(v_and10_try066411))
                      if (v_and10 == CFALSE) {v_and6_try06639 = EID{CFALSE.Id(),0}
                      } else /* arg:11 */{ 
                        var v_and10_try066512 EID 
                        v_and10_try066512 = F_Compile_designated_ask_any(y)
                        /* ERROR PROTECTION INSERTED (v_and10-v_and6_try06639) */
                        if ErrorIn(v_and10_try066512) {v_and6_try06639 = v_and10_try066512
                        } else {
                        v_and10 = ToBoolean(OBJ(v_and10_try066512))
                        if (v_and10 == CFALSE) {v_and6_try06639 = EID{CFALSE.Id(),0}
                        } else /* arg:12 */{ 
                          v_and10 = Core.F__I_equal_any(F_boolean_I_any(ANY(Core.F_CALL(C_Optimize_multi_ask,ARGS(EID{p.Id(),0})))).Id(),CTRUE.Id())
                          if (v_and10 == CFALSE) {v_and6_try06639 = EID{CFALSE.Id(),0}
                          } else /* arg:13 */{ 
                            var v_and10_try066614 EID 
                            /* or:14 */{ 
                              var v_or14 *ClaireBoolean  
                              
                              var v_or14_try066715 EID 
                              v_or14_try066715 = F_Compile_identifiable_ask_any(y)
                              /* ERROR PROTECTION INSERTED (v_or14-v_and10_try066614) */
                              if ErrorIn(v_or14_try066715) {v_and10_try066614 = v_or14_try066715
                              } else {
                              v_or14 = ToBoolean(OBJ(v_or14_try066715))
                              if (v_or14 == CTRUE) {v_and10_try066614 = EID{CTRUE.Id(),0}
                              } else /* or:15 */{ 
                                var v_or14_try066816 EID 
                                /* Let:16 */{ 
                                  var g0669UU *ClaireType  
                                  /* noccur = 1 */
                                  var g0669UU_try067017 EID 
                                  g0669UU_try067017 = Core.F_CALL(C_c_type,ARGS(y.ToEID()))
                                  /* ERROR PROTECTION INSERTED (g0669UU-v_or14_try066816) */
                                  if ErrorIn(g0669UU_try067017) {v_or14_try066816 = g0669UU_try067017
                                  } else {
                                  g0669UU = ToType(OBJ(g0669UU_try067017))
                                  v_or14_try066816 = EID{g0669UU.Included(ToType(C_float.Id())).Id(),0}
                                  }
                                  /* Let-16 */} 
                                /* ERROR PROTECTION INSERTED (v_or14-v_and10_try066614) */
                                if ErrorIn(v_or14_try066816) {v_and10_try066614 = v_or14_try066816
                                } else {
                                v_or14 = ToBoolean(OBJ(v_or14_try066816))
                                if (v_or14 == CTRUE) {v_and10_try066614 = EID{CTRUE.Id(),0}
                                } else /* or:16 */{ 
                                  v_and10_try066614 = EID{CFALSE.Id(),0}/* org-16 */} 
                                /* org-15 */} 
                              }}
                              /* or-14 */} 
                            /* ERROR PROTECTION INSERTED (v_and10-v_and6_try06639) */
                            if ErrorIn(v_and10_try066614) {v_and6_try06639 = v_and10_try066614
                            } else {
                            v_and10 = ToBoolean(OBJ(v_and10_try066614))
                            if (v_and10 == CFALSE) {v_and6_try06639 = EID{CFALSE.Id(),0}
                            } else /* arg:14 */{ 
                              v_and6_try06639 = EID{CTRUE.Id(),0}/* arg-14 */} 
                            /* arg-13 */} 
                          /* arg-12 */} 
                        /* arg-11 */} 
                      }}}
                      /* and-10 */} 
                    } else {
                    v_and6_try06639 = EID{CTRUE.Id(),0}
                    /* If-9 */} 
                  /* ERROR PROTECTION INSERTED (v_and6-v_or4_try06626) */
                  if ErrorIn(v_and6_try06639) {v_or4_try06626 = v_and6_try06639
                  } else {
                  v_and6 = ToBoolean(OBJ(v_and6_try06639))
                  if (v_and6 == CFALSE) {v_or4_try06626 = EID{CFALSE.Id(),0}
                  } else /* arg:9 */{ 
                    v_or4_try06626 = EID{CTRUE.Id(),0}/* arg-9 */} 
                  /* arg-8 */} 
                /* arg-7 */} 
              }
              /* and-6 */} 
            /* ERROR PROTECTION INSERTED (v_or4-v_and2_try06614) */
            if ErrorIn(v_or4_try06626) {v_and2_try06614 = v_or4_try06626
            } else {
            v_or4 = ToBoolean(OBJ(v_or4_try06626))
            if (v_or4 == CTRUE) {v_and2_try06614 = EID{CTRUE.Id(),0}
            } else /* or:6 */{ 
              v_and2_try06614 = EID{CFALSE.Id(),0}/* org-6 */} 
            /* org-5 */} 
          }
          /* or-4 */} 
        /* ERROR PROTECTION INSERTED (v_and2-Result) */
        if ErrorIn(v_and2_try06614) {Result = v_and2_try06614
        } else {
        v_and2 = ToBoolean(OBJ(v_and2_try06614))
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
      var g0671UU *ClaireClass  
      /* noccur = 1 */
      var g0671UU_try06723 EID 
      g0671UU_try06723 = F_Optimize_c_srange_method(self)
      /* ERROR PROTECTION INSERTED (g0671UU-Result) */
      if ErrorIn(g0671UU_try06723) {Result = g0671UU_try06723
      } else {
      g0671UU = ToClass(OBJ(g0671UU_try06723))
      Result = F_Optimize_c_code_method_method2(self,l,_Ztype,g0671UU)
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
              var g0676UU *ClaireList  
              /* noccur = 1 */
              /* Let:7 */{ 
                var i_bag *ClaireList   = ToType(CEMPTY.Id()).EmptyList()
                /* noccur = 2 */
                /* Let:8 */{ 
                  var i int  = 1
                  /* noccur = 4 */
                  /* Let:9 */{ 
                    var g0673 int  = (n-1)
                    /* noccur = 1 */
                    for (i <= g0673) /* while:10 */{ 
                      i_bag.AddFast(l.At(i-1))
                      i = (i+1)
                      /* while-10 */} 
                    /* Let-9 */} 
                  /* Let-8 */} 
                g0676UU = i_bag
                /* Let-7 */} 
              /* Let:7 */{ 
                var g0677UU *ClaireList  
                /* noccur = 1 */
                /* Let:8 */{ 
                  var i_bag *ClaireList   = ToType(CEMPTY.Id()).EmptyList()
                  /* noccur = 2 */
                  /* Let:9 */{ 
                    var i int  = n
                    /* noccur = 4 */
                    /* Let:10 */{ 
                      var g0674 int  = l.Length()
                      /* noccur = 1 */
                      for (i <= g0674) /* while:11 */{ 
                        i_bag.AddFast(l.At(i-1))
                        i = (i+1)
                        /* while-11 */} 
                      /* Let-10 */} 
                    /* Let-9 */} 
                  g0677UU = i_bag
                  /* Let-8 */} 
                l = g0676UU.AddFast(g0677UU.Id())
                /* Let-7 */} 
              /* Let-6 */} 
            /* If-5 */} 
          var g0678I *ClaireBoolean  
          var g0678I_try06795 EID 
          /* and:5 */{ 
            var v_and5 *ClaireBoolean  
            
            v_and5 = self.Inline_ask
            if (v_and5 == CFALSE) {g0678I_try06795 = EID{CFALSE.Id(),0}
            } else /* arg:6 */{ 
              var v_and5_try06807 EID 
              v_and5_try06807 = F_Optimize_c_inline_ask_method(self,l)
              /* ERROR PROTECTION INSERTED (v_and5-g0678I_try06795) */
              if ErrorIn(v_and5_try06807) {g0678I_try06795 = v_and5_try06807
              } else {
              v_and5 = ToBoolean(OBJ(v_and5_try06807))
              if (v_and5 == CFALSE) {g0678I_try06795 = EID{CFALSE.Id(),0}
              } else /* arg:7 */{ 
                g0678I_try06795 = EID{CTRUE.Id(),0}/* arg-7 */} 
              /* arg-6 */} 
            }
            /* and-5 */} 
          /* ERROR PROTECTION INSERTED (g0678I-Result) */
          if ErrorIn(g0678I_try06795) {Result = g0678I_try06795
          } else {
          g0678I = ToBoolean(OBJ(g0678I_try06795))
          if (g0678I == CTRUE) /* If:5 */{ 
            Result = F_Optimize_c_inline_method1(self,l,sx)
            } else {
            /* Let:6 */{ 
              var g0681UU *ClaireList  
              /* noccur = 1 */
              var g0681UU_try06827 EID 
              /* Let:7 */{ 
                var i_bag *ClaireList   = ToType(CEMPTY.Id()).EmptyList()
                /* noccur = 2 */
                /* Let:8 */{ 
                  var i int  = 1
                  /* noccur = 5 */
                  /* Let:9 */{ 
                    var g0675 int  = n
                    /* noccur = 1 */
                    g0681UU_try06827= EID{CFALSE.Id(),0}
                    for (i <= g0675) /* while:10 */{ 
                      var void_try11 EID 
                      _ = void_try11
                      { 
                      /* Let:11 */{ 
                        var g0683UU *ClaireAny  
                        /* noccur = 1 */
                        var g0683UU_try068412 EID 
                        g0683UU_try068412 = F_Compile_c_strict_code_any(l.At(i-1),F_Compile_psort_any(ld.ValuesO()[i-1]))
                        /* ERROR PROTECTION INSERTED (g0683UU-void_try11) */
                        if ErrorIn(g0683UU_try068412) {void_try11 = g0683UU_try068412
                        } else {
                        g0683UU = ANY(g0683UU_try068412)
                        void_try11 = EID{i_bag.AddFast(g0683UU).Id(),0}
                        }
                        /* Let-11 */} 
                      /* ERROR PROTECTION INSERTED (void_try11-void_try11) */
                      if ErrorIn(void_try11) {g0681UU_try06827 = void_try11
                      break
                      } else {
                      i = (i+1)
                      }
                      /* while-10 */} 
                    }
                    /* Let-9 */} 
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (g0681UU_try06827-g0681UU_try06827) */
                if !ErrorIn(g0681UU_try06827) {
                g0681UU_try06827 = EID{i_bag.Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (g0681UU-Result) */
              if ErrorIn(g0681UU_try06827) {Result = g0681UU_try06827
              } else {
              g0681UU = ToList(OBJ(g0681UU_try06827))
              Result = F_Optimize_Call_method_I_method(self,g0681UU).ToEID()
              }
              /* Let-6 */} 
            /* If-5 */} 
          }
          /* Let-4 */} 
        /* Let-3 */} 
      } else {
      if (C_compiler.Optimize_ask == CTRUE) /* If:3 */{ 
        F_Compile_notice_void()
        
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
      var g0685UU *ClaireList  
      /* noccur = 1 */
      var g0685UU_try06863 EID 
      /* Iteration:3 */{ 
        var v_list3 *ClaireList  
        var x *ClaireAny  
        var v_local3 *ClaireAny  
        v_list3 = self.Args
        g0685UU_try06863 = EID{CreateList(ToType(CEMPTY.Id()),v_list3.Length()).Id(),0}
        for CLcount := 0; CLcount < v_list3.Length(); CLcount++{ 
          x = v_list3.At(CLcount)
          var v_local3_try06875 EID 
          v_local3_try06875 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
          /* ERROR PROTECTION INSERTED (v_local3-g0685UU_try06863) */
          if ErrorIn(v_local3_try06875) {g0685UU_try06863 = v_local3_try06875
          g0685UU_try06863 = v_local3_try06875
          break
          } else {
          v_local3 = ANY(v_local3_try06875)
          ToList(OBJ(g0685UU_try06863)).PutAt(CLcount,v_local3)
          } 
        }
        /* Iteration-3 */} 
      /* ERROR PROTECTION INSERTED (g0685UU-Result) */
      if ErrorIn(g0685UU_try06863) {Result = g0685UU_try06863
      } else {
      g0685UU = ToList(OBJ(g0685UU_try06863))
      Result = F_Optimize_use_range_method(self.Arg,g0685UU)
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
            var g0688 *ClaireFunction   = ToFunction(f)
            /* noccur = 1 */
            Result = EID{g0688.Id(),0}
            /* Let-5 */} 
          } else {
          /* Let:5 */{ 
            var g0690UU *ClaireAny  
            /* noccur = 1 */
            var g0690UU_try06916 EID 
            g0690UU_try06916 = Core.F_CALL(C_Compile_function_name,ARGS(EID{p.Id(),0},EID{self.Domain.Id(),0},f.ToEID()))
            /* ERROR PROTECTION INSERTED (g0690UU-Result) */
            if ErrorIn(g0690UU_try06916) {Result = g0690UU_try06916
            } else {
            g0690UU = ANY(g0690UU_try06916)
            Result = F_make_function_string(ToString(g0690UU)).ToEID()
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
              var g0692UU *ClaireAny  
              /* noccur = 1 */
              var g0692UU_try06937 EID 
              /* For:7 */{ 
                var v *ClaireAny  
                _ = v
                g0692UU_try06937= EID{CFALSE.Id(),0}
                var v_support *ClaireList  
                v_support = f.Vars
                v_len := v_support.Length()
                for i_it := 0; i_it < v_len; i_it++ { 
                  v = v_support.At(i_it)
                  var void_try9 EID 
                  _ = void_try9
                  var g0694I *ClaireBoolean  
                  var g0694I_try06959 EID 
                  /* and:9 */{ 
                    var v_and9 *ClaireBoolean  
                    
                    v_and9 = Core.F__sup_integer(Language.F_occurrence_any(x,To_Variable(v)),1)
                    if (v_and9 == CFALSE) {g0694I_try06959 = EID{CFALSE.Id(),0}
                    } else /* arg:10 */{ 
                      var v_and9_try069611 EID 
                      /* Let:11 */{ 
                        var g0697UU *ClaireBoolean  
                        /* noccur = 1 */
                        var g0697UU_try069812 EID 
                        g0697UU_try069812 = F_Compile_designated_ask_any(l.At(n-1))
                        /* ERROR PROTECTION INSERTED (g0697UU-v_and9_try069611) */
                        if ErrorIn(g0697UU_try069812) {v_and9_try069611 = g0697UU_try069812
                        } else {
                        g0697UU = ToBoolean(OBJ(g0697UU_try069812))
                        v_and9_try069611 = EID{g0697UU.Not.Id(),0}
                        }
                        /* Let-11 */} 
                      /* ERROR PROTECTION INSERTED (v_and9-g0694I_try06959) */
                      if ErrorIn(v_and9_try069611) {g0694I_try06959 = v_and9_try069611
                      } else {
                      v_and9 = ToBoolean(OBJ(v_and9_try069611))
                      if (v_and9 == CFALSE) {g0694I_try06959 = EID{CFALSE.Id(),0}
                      } else /* arg:11 */{ 
                        v_and9 = Core.F_owner_any(ANY(Core.F_CALL(C_range,ARGS(v.ToEID())))).IsIn(C_Optimize_Pattern).Not
                        if (v_and9 == CFALSE) {g0694I_try06959 = EID{CFALSE.Id(),0}
                        } else /* arg:12 */{ 
                          g0694I_try06959 = EID{CTRUE.Id(),0}/* arg-12 */} 
                        /* arg-11 */} 
                      /* arg-10 */} 
                    }
                    /* and-9 */} 
                  /* ERROR PROTECTION INSERTED (g0694I-void_try9) */
                  if ErrorIn(g0694I_try06959) {void_try9 = g0694I_try06959
                  } else {
                  g0694I = ToBoolean(OBJ(g0694I_try06959))
                  if (g0694I == CTRUE) /* If:9 */{ 
                     /*v = g0692UU_try06937, s =EID*/
g0692UU_try06937 = EID{CTRUE.Id(),0}
                    break
                    } else {
                    n = (n+1)
                    void_try9 = EID{C__INT,IVAL(n)}
                    /* If-9 */} 
                  }
                  /* ERROR PROTECTION INSERTED (void_try9-g0692UU_try06937) */
                  if ErrorIn(void_try9) {g0692UU_try06937 = void_try9
                  g0692UU_try06937 = void_try9
                  break
                  } else {
                  }
                  /* loop-8 */} 
                /* For-7 */} 
              /* ERROR PROTECTION INSERTED (g0692UU-Result) */
              if ErrorIn(g0692UU_try06937) {Result = g0692UU_try06937
              } else {
              g0692UU = ANY(g0692UU_try06937)
              Result = EID{Core.F_not_any(g0692UU).Id(),0}
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
          var g0701UU *ClaireList  
          /* noccur = 1 */
          /* Iteration:5 */{ 
            var v_list5 *ClaireList  
            var x *ClaireAny  
            var v_local5 *ClaireAny  
            v_list5 = l
            g0701UU = CreateList(ToType(CEMPTY.Id()),v_list5.Length())
            for CLcount := 0; CLcount < v_list5.Length(); CLcount++{ 
              x = v_list5.At(CLcount)
              v_local5 = MakeConstantSet(x).Id()
              g0701UU.PutAt(CLcount,v_local5)
              } 
            /* Iteration-5 */} 
          m = F_Optimize_restriction_I_property(self.Selector,g0701UU,CTRUE)
          /* Let-4 */} 
        if (C_method.Id() == m.Isa.Id()) /* If:4 */{ 
          /* Let:5 */{ 
            var g0699 *ClaireMethod   = ToMethod(m)
            /* noccur = 4 */
            var g0702I *ClaireBoolean  
            var g0702I_try07036 EID 
            /* and:6 */{ 
              var v_and6 *ClaireBoolean  
              
              v_and6 = g0699.Inline_ask
              if (v_and6 == CFALSE) {g0702I_try07036 = EID{CFALSE.Id(),0}
              } else /* arg:7 */{ 
                /* Let:8 */{ 
                  var g0704UU *ClaireAny  
                  /* noccur = 1 */
                  /* For:9 */{ 
                    var s *ClaireAny  
                    _ = s
                    g0704UU= CFALSE.Id()
                    for _,s = range(g0699.Domain.ValuesO())/* loop:10 */{ 
                      if (s.Isa.IsIn(C_Optimize_Pattern) == CTRUE) /* If:11 */{ 
                         /*v = g0704UU, s =any*/
g0704UU = CTRUE.Id()
                        break
                        /* If-11 */} 
                      /* loop-10 */} 
                    /* For-9 */} 
                  v_and6 = F_boolean_I_any(g0704UU)
                  /* Let-8 */} 
                if (v_and6 == CFALSE) {g0702I_try07036 = EID{CFALSE.Id(),0}
                } else /* arg:8 */{ 
                  var v_and6_try07059 EID 
                  v_and6_try07059 = F_Optimize_c_inline_ask_method(g0699,l)
                  /* ERROR PROTECTION INSERTED (v_and6-g0702I_try07036) */
                  if ErrorIn(v_and6_try07059) {g0702I_try07036 = v_and6_try07059
                  } else {
                  v_and6 = ToBoolean(OBJ(v_and6_try07059))
                  if (v_and6 == CFALSE) {g0702I_try07036 = EID{CFALSE.Id(),0}
                  } else /* arg:9 */{ 
                    g0702I_try07036 = EID{CTRUE.Id(),0}/* arg-9 */} 
                  /* arg-8 */} 
                /* arg-7 */} 
              }
              /* and-6 */} 
            /* ERROR PROTECTION INSERTED (g0702I-Result) */
            if ErrorIn(g0702I_try07036) {Result = g0702I_try07036
            } else {
            g0702I = ToBoolean(OBJ(g0702I_try07036))
            if (g0702I == CTRUE) /* If:6 */{ 
              Result = EID{g0699.Id(),0}
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