/***** CLAIRE Compilation of file /Users/ycaseau/Dropbox/src/clairev4.03/src/compile/otool.cl 
         [version 4.0.04 / safety 5] Saturday 01-01-2022 16:47:21 *****/

package Optimize
import (_ "fmt"
	. "Kernel"
	"Core"
	"Language"
	"Reader"
)

//-------- dumb function to prevent import errors --------
func import_g0096() { 
    _ = Core.It
    _ = Language.It
    _ = Reader.It
    } 
  
  
//+-------------------------------------------------------------+
//| CLAIRE                                                      |
//| otool.cl                                                    |
//| Copyright (C) 1994 - 2021 Yves Caseau. All Rights Reserved  |
//| cf. copyright info in file object.cl: about()               |
//+-------------------------------------------------------------+
//-------------------------------------------------------------------
// this file contains the auxiliairy methods for the source optimizer
//-----------------------------------------------------------------
// ******************************************************************
// *  Table of contents                                             *
// *    Part 1: New Instructions & associated stuff                 *
// *    Part 2: Optimizer Warnings                                  *
// *    Part 3: Type Handling                                       *
// *    Part 4: Miscellaneous                                       *
// ******************************************************************
// ******************************************************************
// *    Part 1: New Instructions & associated stuff                 *
// ******************************************************************
//
// this is a same-sort (object) casting from one class to another because of the
// stupidity of the target type system
// its use is linked to stupid_t(x)
/* The go function for: self_print(self:Compile/C_cast) [status=1] */
func (self *Compile_CCast ) SelfPrint () EID { 
    // eid body s = void
    var Result EID 
    PRINC("<")
    Result = Core.F_CALL(C_print,ARGS(self.Arg.ToEID()))
    if !ErrorIn(Result) {
    PRINC(":")
    Result = Core.F_print_any(self.SetArg.Id())
    if !ErrorIn(Result) {
    PRINC(">")
    Result = EVOID
    }}
    return Result} 
  
// The EID go function for: self_print @ Compile/C_cast (throw: true) 
func E_self_print_C_cast (self EID) EID { 
    return To_Compile_CCast(OBJ(self)).SelfPrint( )} 
  
/* The go function for: Compile/c_gc?(self:Compile/C_cast) [status=0] */
func (self *Compile_CCast ) CGc_ask () *ClaireBoolean  { 
    if (To_Compile_CCast(self.Arg).CGc_ask() == CTRUE) {return CTRUE
    } else {return CFALSE}} 
  
// The EID go function for: Compile/c_gc? @ Compile/C_cast (throw: false) 
func E_Compile_c_gc_ask_C_cast (self EID) EID { 
    return EID{To_Compile_CCast(OBJ(self)).CGc_ask( ).Id(),0}} 
  
/* The go function for: c_type(self:Compile/C_cast) [status=0] */
func (self *Compile_CCast ) CType () *ClaireType  { 
    return  ToType(self.SetArg.Id())
    } 
  
// The EID go function for: c_type @ Compile/C_cast (throw: false) 
func E_c_type_C_cast (self EID) EID { 
    return EID{To_Compile_CCast(OBJ(self)).CType( ).Id(),0}} 
  
// v3.0 : better safe
/* The go function for: c_code(self:Compile/C_cast,s:class) [status=1] */
func F_c_code_C_cast (self *Compile_CCast ,s *ClaireClass ) EID { 
    // eid body s = any
    var Result EID 
    if (s.IsIn(C_object) == CTRUE) { 
      { var _CL_obj *Compile_CCast   = To_Compile_CCast(new(Compile_CCast).Is(C_Compile_C_cast))
        { 
          var va_arg1 *Compile_CCast  
          var va_arg2 *ClaireAny  
          va_arg1 = _CL_obj
          var try_1 EID 
          try_1 = Core.F_CALL(C_c_code,ARGS(self.Arg.ToEID(),EID{s.Id(),0}))
          if ErrorIn(try_1) {Result = try_1
          } else {
          va_arg2 = ANY(try_1)
          va_arg1.Arg = va_arg2
          Result = va_arg2.ToEID()
          }
          } 
        if !ErrorIn(Result) {
        _CL_obj.SetArg = self.SetArg
        Result = EID{_CL_obj.Id(),0}
        }
        } 
      } else {
      Result = Core.F_CALL(C_c_code,ARGS(self.Arg.ToEID(),EID{s.Id(),0}))
      } 
    return Result} 
  
// The EID go function for: c_code @ Compile/C_cast (throw: true) 
func E_c_code_C_cast (self EID,s EID) EID { 
    return F_c_code_C_cast(To_Compile_CCast(OBJ(self)),ToClass(OBJ(s)) )} 
  
// we need a new type to express powerful Iterate rules
// Note: Patterns require the compiler !
/* The go function for: self_print(self:Pattern) [status=1] */
func (self *ClairePattern ) SelfPrint () EID { 
    // eid body s = void
    var Result EID 
    Result = Core.F_print_any(self.Selector.Id())
    if !ErrorIn(Result) {
    PRINC("[tuple(")
    Result = Core.F_princ_list(self.Arg)
    if !ErrorIn(Result) {
    PRINC(")]")
    Result = EVOID
    }}
    return Result} 
  
// The EID go function for: self_print @ Pattern (throw: true) 
func E_self_print_Pattern (self EID) EID { 
    return To_ClairePattern(OBJ(self)).SelfPrint( )} 
  
/* The go function for: %(x:any,y:Pattern) [status=1] */
func F__Z_any3 (x *ClaireAny ,y *ClairePattern ) EID { 
    // eid body s = boolean
    var Result EID 
    if (x.Isa.IsIn(Language.C_Call) == CTRUE) { 
      { var g0097 *Language.Call   = Language.To_Call(x)
        { 
          var v_and4 *ClaireBoolean  
          
          v_and4 = Equal(g0097.Selector.Id(),y.Selector.Id())
          if (v_and4 == CFALSE) {Result = EID{CFALSE.Id(),0}
          } else { 
            var try_1 EID 
            { var arg_2 *ClaireList  
              var try_3 EID 
              { 
                var v_list7 *ClaireList  
                var z *ClaireAny  
                var v_local7 *ClaireAny  
                v_list7 = g0097.Args
                try_3 = EID{CreateList(ToType(CEMPTY.Id()),v_list7.Length()).Id(),0}
                for CLcount := 0; CLcount < v_list7.Length(); CLcount++{ 
                  z = v_list7.At(CLcount)
                  var try_4 EID 
                  try_4 = Core.F_CALL(C_c_type,ARGS(z.ToEID()))
                  if ErrorIn(try_4) {try_3 = try_4
                  break
                  } else {
                  v_local7 = ANY(try_4)
                  ToList(OBJ(try_3)).PutAt(CLcount,v_local7)
                  } 
                }
                } 
              if ErrorIn(try_3) {try_1 = try_3
              } else {
              arg_2 = ToList(OBJ(try_3))
              try_1 = EID{Core.F_tmatch_ask_list(arg_2,y.Arg).Id(),0}
              }
              } 
            if ErrorIn(try_1) {Result = try_1
            } else {
            v_and4 = ToBoolean(OBJ(try_1))
            if (v_and4 == CFALSE) {Result = EID{CFALSE.Id(),0}
            } else { 
              Result = EID{CTRUE.Id(),0}} 
            } 
          }
          } 
        } 
      } else {
      Result = EID{CFALSE.Id(),0}
      } 
    return Result} 
  
// The EID go function for: % @ list<type_expression>(any, Pattern) (throw: true) 
func E__Z_any3 (x EID,y EID) EID { 
    return F__Z_any3(ANY(x),To_ClairePattern(OBJ(y)) )} 
  
// this is very lazy, we could do better
/* The go function for: glb(x:Pattern,y:type_expression) [status=1] */
func (x *ClairePattern ) Glb (y *ClaireTypeExpression ) *ClaireTypeExpression  { 
    // procedure body with s = type_expression
    var Result *ClaireTypeExpression  
    if (y.Isa.IsIn(C_Optimize_Pattern) == CTRUE) { 
      { var g0099 *ClairePattern   = To_ClairePattern(y.Id())
        if (x.Selector.Id() == g0099.Selector.Id()) { 
          { var _CL_obj *ClairePattern   = To_ClairePattern(new(ClairePattern).Is(C_Optimize_Pattern))
            _CL_obj.Selector = x.Selector
            _CL_obj.Arg = ToList(OBJ(Core.F_CALL(ToProperty(Core.C_glb.Id()),ARGS(EID{x.Arg.Id(),0},EID{g0099.Arg.Id(),0}))))
            Result = ToTypeExpression(_CL_obj.Id())
            } 
          } else {
          Result = ToTypeExpression(CEMPTY.Id())
          } 
        } 
      } else {
      Result = ToTypeExpression(CEMPTY.Id())
      } 
    return Result} 
  
// The EID go function for: glb @ Pattern (throw: false) 
func E_glb_Pattern (x EID,y EID) EID { 
    return EID{To_ClairePattern(OBJ(x)).Glb(ToTypeExpression(OBJ(y)) ).Id(),0}} 
  
// extension of <= for Patterns
/* The go function for: less?(x:Pattern,y:type_expression) [status=0] */
func F_less_ask_Pattern (x *ClairePattern ,y *ClaireTypeExpression ) *ClaireBoolean  { 
    // procedure body with s = boolean
    var Result *ClaireBoolean  
    if (y.Isa.IsIn(C_Optimize_Pattern) == CTRUE) { 
      { var g0101 *ClairePattern   = To_ClairePattern(y.Id())
        { 
          var v_and4 *ClaireBoolean  
          
          v_and4 = Equal(x.Selector.Id(),g0101.Selector.Id())
          if (v_and4 == CFALSE) {Result = CFALSE
          } else { 
            v_and4 = Equal(MakeInteger(x.Arg.Length()).Id(),MakeInteger(g0101.Arg.Length()).Id())
            if (v_and4 == CFALSE) {Result = CFALSE
            } else { 
              { var arg_1 *ClaireAny  
                { var i int  = 1
                  { var g0102 int  = x.Arg.Length()
                    arg_1= CFALSE.Id()
                    for (i <= g0102) { 
                      if (Core.F__equaltype_ask_any(ToType(x.Arg.At(i-1)),ToType(g0101.Arg.At(i-1))) != CTRUE) { 
                        arg_1 = CTRUE.Id()
                        break
                        } 
                      i = (i+1)
                      } 
                    } 
                  } 
                v_and4 = Core.F_not_any(arg_1)
                } 
              if (v_and4 == CFALSE) {Result = CFALSE
              } else { 
                Result = CTRUE} 
              } 
            } 
          } 
        } 
      } else {
      Result = ToType(Language.C_Call.Id()).Included(ToType(y.Id()))
      } 
    return Result} 
  
// The EID go function for: less? @ Pattern (throw: false) 
func E_less_ask_Pattern (x EID,y EID) EID { 
    return EID{F_less_ask_Pattern(To_ClairePattern(OBJ(x)),ToTypeExpression(OBJ(y)) ).Id(),0}} 
  
/* The go function for: less?(x:type_expression,y:Pattern) [status=1] */
func F_less_ask_type_expression2 (x *ClaireTypeExpression ,y *ClairePattern ) EID { 
    // eid body s = boolean
    var Result EID 
    if (C_set.Id() == x.Isa.Id()) { 
      { var g0104 *ClaireSet   = ToSet(x.Id())
        { var arg_1 *ClaireAny  
          var try_2 EID 
          { 
            var z *ClaireAny  
            _ = z
            try_2= EID{CFALSE.Id(),0}
            var z_support *ClaireSet  
            z_support = g0104
            for i_it := 0; i_it < z_support.Count; i_it++ { 
              z = z_support.At(i_it)
              var loop_3 EID 
              _ = loop_3
              var g0108I *ClaireBoolean  
              var try_4 EID 
              { var arg_5 *ClaireBoolean  
                var try_6 EID 
                try_6 = F__Z_any3(z,y)
                if ErrorIn(try_6) {try_4 = try_6
                } else {
                arg_5 = ToBoolean(OBJ(try_6))
                try_4 = EID{arg_5.Not.Id(),0}
                }
                } 
              if ErrorIn(try_4) {loop_3 = try_4
              } else {
              g0108I = ToBoolean(OBJ(try_4))
              if (g0108I == CTRUE) { 
                try_2 = EID{CTRUE.Id(),0}
                break
                } else {
                loop_3 = EID{CFALSE.Id(),0}
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
      }  else if (x.Isa.IsIn(C_Optimize_Pattern) == CTRUE) { 
      { var g0105 *ClairePattern   = To_ClairePattern(x.Id())
        { 
          var v_and4 *ClaireBoolean  
          
          v_and4 = Equal(g0105.Selector.Id(),y.Selector.Id())
          if (v_and4 == CFALSE) {Result = EID{CFALSE.Id(),0}
          } else { 
            v_and4 = Equal(MakeInteger(g0105.Arg.Length()).Id(),MakeInteger(y.Arg.Length()).Id())
            if (v_and4 == CFALSE) {Result = EID{CFALSE.Id(),0}
            } else { 
              { var arg_7 *ClaireAny  
                { var i int  = 1
                  { var g0106 int  = g0105.Arg.Length()
                    arg_7= CFALSE.Id()
                    for (i <= g0106) { 
                      if (Core.F__equaltype_ask_any(ToType(g0105.Arg.At(i-1)),ToType(y.Arg.At(i-1))) != CTRUE) { 
                        arg_7 = CTRUE.Id()
                        break
                        } 
                      i = (i+1)
                      } 
                    } 
                  } 
                v_and4 = Core.F_not_any(arg_7)
                } 
              if (v_and4 == CFALSE) {Result = EID{CFALSE.Id(),0}
              } else { 
                Result = EID{CTRUE.Id(),0}} 
              } 
            } 
          } 
        } 
      } else {
      Result = EID{CFALSE.Id(),0}
      } 
    return Result} 
  
// The EID go function for: less? @ list<type_expression>(type_expression, Pattern) (throw: true) 
func E_less_ask_type_expression2 (x EID,y EID) EID { 
    return F_less_ask_type_expression2(ToTypeExpression(OBJ(x)),To_ClairePattern(OBJ(y)) )} 
  
// v0.03 must return a type
/* The go function for: nth(p:property,x:tuple) [status=0] */
func F_nth_property (p *ClaireProperty ,x *ClaireTuple ) *ClairePattern  { 
    // procedure body with s = Pattern
    var Result *ClairePattern  
    { var _CL_obj *ClairePattern   = To_ClairePattern(new(ClairePattern).Is(C_Optimize_Pattern))
      _CL_obj.Selector = p
      _CL_obj.Arg = x.List_I()
      Result = _CL_obj
      } 
    return Result} 
  
// The EID go function for: nth @ property (throw: false) 
func E_nth_property (p EID,x EID) EID { 
    return EID{F_nth_property(ToProperty(OBJ(p)),ToTuple(OBJ(x)) ).Id(),0}} 
  
// ******************************************************************
// *    Part 2: Optimizer Warnings                                  *
// ******************************************************************
// unified warning
/* The go function for: Compile/warn(_CL_obj:void) [status=0] */
func F_Compile_warn_void ()  { 
    // procedure body with s = void
    C_compiler.NWarnings = (C_compiler.NWarnings+1)
    if (C_OPT.InMethod != CNULL) { 
      Core.F_tformat_string(MakeString("---- WARNING[in ~S, line ~A]: "),1,MakeConstantList(C_OPT.InMethod,MakeInteger(ClEnv.NLine).Id()))
      } else {
      Core.F_tformat_string(MakeString("---- WARNING[lien ~A]: "),1,MakeConstantList(MakeInteger(ClEnv.NLine).Id()))
      } 
    } 
  
// The EID go function for: Compile/warn @ void (throw: false) 
func E_Compile_warn_void (_CL_obj EID) EID { 
    F_Compile_warn_void( )
    return EVOID} 
  
/* The go function for: Compile/Cerror(s:string,l:listargs) [status=1] */
func F_Compile_Cerror_string (s *ClaireString ,l *ClaireList ) EID { 
    // eid body s = void
    var Result EID 
    PRINC("---- Compiler Error[in ")
    Result = Core.F_CALL(C_print,ARGS(C_OPT.InMethod.ToEID()))
    if !ErrorIn(Result) {
    PRINC("]:\n")
    Result = EVOID
    }
    if !ErrorIn(Result) {
    PRINC("---- file read up to line ")
    F_princ_integer(ClEnv.NLine)
    PRINC("\n")
    Result = ToException(Core.C_general_error.Make((s).Id(),l.Id())).Close()
    }
    return Result} 
  
// The EID go function for: Compile/Cerror @ string (throw: true) 
func E_Compile_Cerror_string (s EID,l EID) EID { 
    return F_Compile_Cerror_string(ToString(OBJ(s)),ToList(OBJ(l)) )} 
  
// a note
/* The go function for: Compile/notice(_CL_obj:void) [status=0] */
func F_Compile_notice_void ()  { 
    // procedure body with s = void
    C_compiler.NNotes = (C_compiler.NNotes+1)
    if (C_OPT.InMethod != CNULL) { 
      Core.F_tformat_string(MakeString("---- note[in ~S]: "),2,MakeConstantList(C_OPT.InMethod))
      } else {
      Core.F_tformat_string(MakeString("---- note: "),2,ToType(CEMPTY.Id()).EmptyList())
      } 
    } 
  
// The EID go function for: Compile/notice @ void (throw: false) 
func E_Compile_notice_void (_CL_obj EID) EID { 
    F_Compile_notice_void( )
    return EVOID} 
  
// Warning : compiling is impossible, wrong selector
/* The go function for: c_warn(self:Call,%type:any) [status=1] */
func F_Optimize_c_warn_Call (self *Language.Call ,_Ztype *ClaireAny ) EID { 
    // eid body s = any
    var Result EID 
    { var s *ClaireProperty   = self.Selector
      if (_Ztype == C_void.Id()) { 
        Result = Core.F_CALL(C_Compile_Cerror,ARGS(EID{MakeString("[205] message ~S sent to void object").Id(),0},EID{self.Id(),0}))
        }  else if ((F_boolean_I_any(s.Restrictions.Id()).Id() != CTRUE.Id()) && 
          (C_OPT.Ignore.Contain_ask(s.Id()) != CTRUE)) { 
        F_Compile_warn_void()
        Result = Core.F_tformat_string(MakeString("the property ~S is undefined [255]\n"),1,MakeConstantList(s.Id()))
        } else {
        var g0112I *ClaireBoolean  
        { 
          var v_and4 *ClaireBoolean  
          
          v_and4 = C_OPT.Ignore.Contain_ask(s.Id()).Not
          if (v_and4 == CFALSE) {g0112I = CFALSE
          } else { 
            v_and4 = MakeBoolean((s.Open <= 1) || (s.Open == 4))
            if (v_and4 == CFALSE) {g0112I = CFALSE
            } else { 
              if (_Ztype.Isa.IsIn(C_list) == CTRUE) { 
                { var g0111 *ClaireList   = ToList(_Ztype)
                  v_and4 = Core.F__I_equal_any(MakeInteger(ToTypeExpression(g0111.At(1-1)).Class_I().Open).Id(),MakeInteger(3).Id())
                  } 
                } else {
                v_and4 = CFALSE
                } 
              if (v_and4 == CFALSE) {g0112I = CFALSE
              } else { 
                g0112I = CTRUE} 
              } 
            } 
          } 
        if (g0112I == CTRUE) { 
          F_Compile_warn_void()
          Result = Core.F_tformat_string(MakeString("wrongly typed message ~S (~S) [256]\n"),1,MakeConstantList(self.Id(),_Ztype))
          }  else if (C_compiler.Optimize_ask == CTRUE) { 
          F_Compile_notice_void()
          Result = EID{CFALSE.Id(),0}
          } else {
          Result = EID{CFALSE.Id(),0}
          } 
        } 
      if !ErrorIn(Result) {
      Result = F_Optimize_open_message_property(self.Selector,self.Args)
      }
      } 
    return Result} 
  
// The EID go function for: c_warn @ Call (throw: true) 
func E_Optimize_c_warn_Call (self EID,_Ztype EID) EID { 
    return F_Optimize_c_warn_Call(Language.To_Call(OBJ(self)),ANY(_Ztype) )} 
  
/* The go function for: c_warn(self:Super,%type:any) [status=1] */
func F_Optimize_c_warn_Super (self *Language.Super ,_Ztype *ClaireAny ) EID { 
    // eid body s = any
    var Result EID 
    { var s *ClaireProperty   = self.Selector
      if (_Ztype == C_void.Id()) { 
        Result = Core.F_CALL(C_Compile_Cerror,ARGS(EID{MakeString("[205] message ~S sent to void object").Id(),0},EID{self.Id(),0}))
        }  else if (F_boolean_I_any(s.Restrictions.Id()).Id() != CTRUE.Id()) { 
        F_Compile_warn_void()
        Result = Core.F_tformat_string(MakeString("the property ~S is undefined [255]\n"),1,MakeConstantList(s.Id()))
        }  else if ((C_OPT.Ignore.Contain_ask(s.Id()) != CTRUE) && 
          (s.Open <= 1)) { 
        Result = EID{CFALSE.Id(),0}
        } else {
        Result = EID{CFALSE.Id(),0}
        } 
      if !ErrorIn(Result) {
      { var m *Language.Call  
        var try_1 EID 
        try_1 = F_Optimize_open_message_property(self.Selector,self.Args)
        if ErrorIn(try_1) {Result = try_1
        } else {
        m = Language.To_Call(OBJ(try_1))
        { var _CL_obj *Language.Super   = Language.To_Super(new(Language.Super).Is(Language.C_Super))
          _CL_obj.Selector = m.Selector
          _CL_obj.CastTo = self.CastTo
          _CL_obj.Args = m.Args
          Result = EID{_CL_obj.Id(),0}
          } 
        }
        } 
      }
      } 
    return Result} 
  
// The EID go function for: c_warn @ Super (throw: true) 
func E_Optimize_c_warn_Super (self EID,_Ztype EID) EID { 
    return F_Optimize_c_warn_Super(Language.To_Super(OBJ(self)),ANY(_Ztype) )} 
  
// a message cannot be compiled into efficient code
// here the property does not allow the compilation and we want to see it
/* The go function for: c_warn(self:property,l:list,%type:list) [status=1] */
func F_Optimize_c_warn_property (self *ClaireProperty ,l *ClaireList ,_Ztype *ClaireList ) EID { 
    // eid body s = any
    var Result EID 
    if ((self.Open <= 1) && 
        ((C_OPT.Ignore.Contain_ask(self.Id()) != CTRUE) && 
          (C_compiler.Safety >= 2))) { 
      
      } 
    Result = F_Optimize_open_message_property(self,l)
    return Result} 
  
// The EID go function for: c_warn @ property (throw: true) 
func E_Optimize_c_warn_property (self EID,l EID,_Ztype EID) EID { 
    return F_Optimize_c_warn_property(ToProperty(OBJ(self)),ToList(OBJ(l)),ToList(OBJ(_Ztype)) )} 
  
// a variable should not be abused ! Either it is a true error or it is
// simply dangerous. The result is the value to be used (either x or
// ckeck_in(x,range(oself))
/* The go function for: c_warn(self:Variable,x:any,y:type) [status=1] */
func F_Optimize_c_warn_Variable (self *ClaireVariable ,x *ClaireAny ,y *ClaireType ) EID { 
    // eid body s = any
    var Result EID 
    if (self.Index == -1) { 
      Result = x.ToEID()
      } else {
      var g0113I *ClaireBoolean  
      var try_1 EID 
      { var arg_2 *ClaireBoolean  
        var try_3 EID 
        { var arg_4 *ClaireAny  
          var try_5 EID 
          try_5 = Core.F_CALL(ToProperty(C__exp.Id()),ARGS(EID{y.Id(),0},EID{self.Range.Id(),0}))
          if ErrorIn(try_5) {try_3 = try_5
          } else {
          arg_4 = ANY(try_5)
          try_3 = EID{F_boolean_I_any(arg_4).Id(),0}
          }
          } 
        if ErrorIn(try_3) {try_1 = try_3
        } else {
        arg_2 = ToBoolean(OBJ(try_3))
        try_1 = EID{Core.F__I_equal_any(arg_2.Id(),CTRUE.Id()).Id(),0}
        }
        } 
      if ErrorIn(try_1) {Result = try_1
      } else {
      g0113I = ToBoolean(OBJ(try_1))
      if (g0113I == CTRUE) { 
        if (C_compiler.Safety >= 2) { 
          F_Compile_warn_void()
          Result = Core.F_tformat_string(MakeString("~S of type ~S is put in the variable ~S:~S [257a]\n"),1,MakeConstantList(x,
            y.Id(),
            self.Id(),
            self.Range.Id()))
          } else {
          Result = Core.F_CALL(C_Compile_Cerror,ARGS(EID{MakeString("[212] the value ~S of type ~S cannot be placed in the variable ~S:~S").Id(),0},
            x.ToEID(),
            EID{y.Id(),0},
            EID{self.Id(),0},
            EID{self.Range.Id(),0}))
          } 
        }  else if ((C_compiler.Safety <= 1) || 
          (F_boolean_I_any(F_Compile_sort_equal_class(F_Compile_osort_any(self.Range.Id()),F_Compile_osort_any(y.Id()))).Id() != CTRUE.Id())) { 
        F_Compile_warn_void()
        Result = Core.F_tformat_string(MakeString("~S of type ~S is put in the variable ~S:~S (~A) [257b]\n"),1,MakeConstantList(x,
          y.Id(),
          self.Id(),
          self.Range.Id(),
          MakeInteger(self.Index).Id()))
        } else {
        Result = EID{CFALSE.Id(),0}
        } 
      }
      } 
    if !ErrorIn(Result) {
    if ((C_compiler.Safety <= 1) && 
        (y.Included(self.Range) != CTRUE)) { 
      Result = Core.F_CALL(C_c_code,ARGS(x.ToEID(),EID{C_any.Id(),0}))
      } else {
      Result = x.ToEID()
      } 
    }
    return Result} 
  
// The EID go function for: c_warn @ Variable (throw: true) 
func E_Optimize_c_warn_Variable (self EID,x EID,y EID) EID { 
    return F_Optimize_c_warn_Variable(To_Variable(OBJ(self)),ANY(x),ToType(OBJ(y)) )} 
  
// ******************************************************************
// *    Part 3: Type Handling                                       *
// ******************************************************************
// we use  {any U type} to represent the change of sort  (to any)
//         {} U (c U t) to represent a change of psort   (to c)
// e.g.: (any U class) = class stored as an OID
// tests if two sorts are similar
// the compiler.overflow? test
/* The go function for: Compile/sort=(c:class,c2:class) [status=0] */
func F_Compile_sort_equal_class (c *ClaireClass ,c2 *ClaireClass ) *ClaireAny  { 
    if (c.IsIn(C_object) == CTRUE) { 
      return  c2.IsIn(C_object).Id()
      } else {
      return  MakeBoolean((c.Id() == c2.Id()) || (((C_compiler.Overflow_ask != CTRUE) && 
            ((c.Id() == C_any.Id()) && 
                (c2.Id() == C_integer.Id()))) || 
          ((c.Id() == C_integer.Id()) && 
              (c2.Id() == C_any.Id())))).Id()
      } 
    } 
  
// The EID go function for: Compile/sort= @ class (throw: false) 
func E_Compile_sort_equal_class (c EID,c2 EID) EID { 
    return F_Compile_sort_equal_class(ToClass(OBJ(c)),ToClass(OBJ(c2)) ).ToEID()} 
  
// give the "precise sort", i.e., a class under object is a sort
/* The go function for: Compile/psort(x:any) [status=0] */
func F_Compile_psort_any (x *ClaireAny ) *ClaireClass  { 
    // procedure body with s = class
    var Result *ClaireClass  
    { var c *ClaireClass   = ToTypeExpression(x).Class_I()
      if (c.IsIn(C_object) == CTRUE) { 
        Result = c
        } else {
        Result = c.Sort_I()
        } 
      } 
    return Result} 
  
// The EID go function for: Compile/psort @ any (throw: false) 
func E_Compile_psort_any (x EID) EID { 
    return EID{F_Compile_psort_any(ANY(x) ).Id(),0}} 
  
// gives the "optimizer sort", which is one of
// any, object, float, X <= import,
/* The go function for: Compile/osort(x:any) [status=0] */
func F_Compile_osort_any (x *ClaireAny ) *ClaireClass  { 
    return  ToTypeExpression(x).Class_I().Sort_I()
    } 
  
// The EID go function for: Compile/osort @ any (throw: false) 
func E_Compile_osort_any (x EID) EID { 
    return EID{F_Compile_osort_any(ANY(x) ).Id(),0}} 
  
/* The go function for: sort(x:Variable) [status=0] */
func F_sort_Variable (x *ClaireVariable ) *ClaireClass  { 
    // procedure body with s = class
    var Result *ClaireClass  
    { var r *ClaireType   = x.Range
      var g0115I *ClaireBoolean  
      if (r.Isa.IsIn(C_Union) == CTRUE) { 
        { var g0114 *ClaireUnion   = To_Union(r.Id())
          g0115I = Equal(g0114.T1.Id(),CEMPTY.Id())
          } 
        } else {
        g0115I = CFALSE
        } 
      if (g0115I == CTRUE) { 
        Result = F_Compile_psort_any(To_Union(To_Union(r.Id()).T2.Id()).T2.Id())
        } else {
        Result = F_Compile_psort_any(r.Id())
        } 
      } 
    return Result} 
  
// The EID go function for: sort @ Variable (throw: false) 
func E_sort_Variable (x EID) EID { 
    return EID{F_sort_Variable(To_Variable(OBJ(x)) ).Id(),0}} 
  
// this is a very stupid type inference that mimicks the go compiler - defined in pretty.cl with CLAIRE4
// it returns a class
/* The go function for: Compile/stupid_t(self:any) [status=1] */
func F_Compile_stupid_t_any1 (self *ClaireAny ) EID { 
    // eid body s = class
    var Result EID 
    Result = Language.F_static_type_any(self)
    return Result} 
  
// The EID go function for: Compile/stupid_t @ list<type_expression>(any) (throw: true) 
func E_Compile_stupid_t_any1 (self EID) EID { 
    return F_Compile_stupid_t_any1(ANY(self) )} 
  
// comparison
/* The go function for: Compile/stupid_t(self:any,x:any) [status=1] */
func F_Compile_stupid_t_any2 (self *ClaireAny ,x *ClaireAny ) EID { 
    // eid body s = boolean
    var Result EID 
    { var c1 *ClaireClass  
      var try_1 EID 
      try_1 = F_Compile_stupid_t_any1(self)
      if ErrorIn(try_1) {Result = try_1
      } else {
      c1 = ToClass(OBJ(try_1))
      { var c2 *ClaireClass  
        var try_2 EID 
        try_2 = F_Compile_stupid_t_any1(x)
        if ErrorIn(try_2) {Result = try_2
        } else {
        c2 = ToClass(OBJ(try_2))
        Result = EID{MakeBoolean((c1.Id() != C_any.Id()) && (c1.Id() == c2.Id())).Id(),0}
        }
        } 
      }
      } 
    return Result} 
  
// The EID go function for: Compile/stupid_t @ list<type_expression>(any, any) (throw: true) 
func E_Compile_stupid_t_any2 (self EID,x EID) EID { 
    return F_Compile_stupid_t_any2(ANY(self),ANY(x) )} 
  
// an extended type is of the kind (t U {unknown})
// CLAIRE4: got rid of optUnion
/* The go function for: extended?(self:type) [status=1] */
func F_Optimize_extended_ask_type (self *ClaireType ) EID { 
    // eid body s = boolean
    var Result EID 
    if (self.Isa.IsIn(C_Union) == CTRUE) { 
      { var g0116 *ClaireUnion   = To_Union(self.Id())
        { 
          var v_and4 *ClaireBoolean  
          
          v_and4 = Equal(C_set.Id(),g0116.T2.Isa.Id())
          if (v_and4 == CFALSE) {Result = EID{CFALSE.Id(),0}
          } else { 
            var try_1 EID 
            { var arg_2 *ClaireAny  
              var try_3 EID 
              try_3 = Core.F_CALL(C_size,ARGS(EID{g0116.T2.Id(),0}))
              if ErrorIn(try_3) {try_1 = try_3
              } else {
              arg_2 = ANY(try_3)
              try_1 = EID{Equal(arg_2,MakeInteger(1).Id()).Id(),0}
              }
              } 
            if ErrorIn(try_1) {Result = try_1
            } else {
            v_and4 = ToBoolean(OBJ(try_1))
            if (v_and4 == CFALSE) {Result = EID{CFALSE.Id(),0}
            } else { 
              var try_4 EID 
              { var arg_5 *ClaireAny  
                var try_6 EID 
                try_6 = Core.F_the_type(g0116.T2)
                if ErrorIn(try_6) {try_4 = try_6
                } else {
                arg_5 = ANY(try_6)
                try_4 = EID{Equal(arg_5,CNULL).Id(),0}
                }
                } 
              if ErrorIn(try_4) {Result = try_4
              } else {
              v_and4 = ToBoolean(OBJ(try_4))
              if (v_and4 == CFALSE) {Result = EID{CFALSE.Id(),0}
              } else { 
                Result = EID{CTRUE.Id(),0}} 
              } 
            } 
          }}
          } 
        } 
      } else {
      Result = EID{CFALSE.Id(),0}
      } 
    return Result} 
  
// The EID go function for: extended? @ type (throw: true) 
func E_Optimize_extended_ask_type (self EID) EID { 
    return F_Optimize_extended_ask_type(ToType(OBJ(self)) )} 
  
// creates an extended type (v0.02) that can be checked easily 
// CLAIRE4: we removed the syntactic marker optUnion for (X U {unknown})
// used in ocall and ocontrol
/* The go function for: extends(x:type) [status=0] */
func F_Optimize_extends_type (x *ClaireType ) *ClaireType  { 
    // procedure body with s = type
    var Result *ClaireType  
    { var _CL_obj *ClaireUnion   = To_Union(new(ClaireUnion).Is(C_Union))
      _CL_obj.T1 = x
      _CL_obj.T2 = ToType(MakeConstantSet(CNULL).Id())
      Result = ToType(_CL_obj.Id())
      } 
    return Result} 
  
// The EID go function for: extends @ type (throw: false) 
func E_Optimize_extends_type (x EID) EID { 
    return EID{F_Optimize_extends_type(ToType(OBJ(x)) ).Id(),0}} 
  
// a sort abstraction is the special union any U t, which is known to represent t by
// the type system (used for variables only) but tells the compiler that the sort is any
/* The go function for: sort_abstract!(x:type) [status=0] */
func F_Optimize_sort_abstract_I_type (x *ClaireType ) *ClaireType  { 
    // procedure body with s = type
    var Result *ClaireType  
    if ((ANY(Core.F_CALL(C_sort_I,ARGS(EID{x.Id(),0}))) != C_any.Id()) && 
        ((ANY(Core.F_CALL(C_sort_I,ARGS(EID{x.Id(),0}))) != C_integer.Id()) && 
          (ANY(Core.F_CALL(C_sort_I,ARGS(EID{x.Id(),0}))) != C_float.Id()))) { 
      { var _CL_obj *ClaireUnion   = To_Union(new(ClaireUnion).Is(C_Union))
        _CL_obj.T1 = ToType(C_any.Id())
        _CL_obj.T2 = x
        Result = ToType(_CL_obj.Id())
        } 
      } else {
      Result = x
      } 
    return Result} 
  
// The EID go function for: sort_abstract! @ type (throw: false) 
func E_Optimize_sort_abstract_I_type (x EID) EID { 
    return EID{F_Optimize_sort_abstract_I_type(ToType(OBJ(x)) ).Id(),0}} 
  
// v3.00.05
/* The go function for: sort_abstract?(x:type) [status=0] */
func F_Optimize_sort_abstract_ask_type (x *ClaireType ) *ClaireBoolean  { 
    // procedure body with s = boolean
    var Result *ClaireBoolean  
    if (x.Isa.IsIn(C_Union) == CTRUE) { 
      { var g0118 *ClaireUnion   = To_Union(x.Id())
        Result = Equal(g0118.T1.Id(),C_any.Id())
        } 
      } else {
      Result = CFALSE
      } 
    return Result} 
  
// The EID go function for: sort_abstract? @ type (throw: false) 
func E_Optimize_sort_abstract_ask_type (x EID) EID { 
    return EID{F_Optimize_sort_abstract_ask_type(ToType(OBJ(x)) ).Id(),0}} 
  
// since we introduce some fuzziness with types (any U t), we need a way to get
// the precise type t back
/* The go function for: ptype(x:type) [status=0] */
func F_Optimize_ptype_type (x *ClaireType ) *ClaireType  { 
    // procedure body with s = type
    var Result *ClaireType  
    if (x.Isa.IsIn(C_Union) == CTRUE) { 
      { var g0120 *ClaireUnion   = To_Union(x.Id())
        if (g0120.T1.Id() == C_any.Id()) { 
          Result = g0120.T2
          } else {
          Result = ToType(g0120.Id())
          } 
        } 
      } else {
      Result = x
      } 
    return Result} 
  
// The EID go function for: ptype @ type (throw: false) 
func E_Optimize_ptype_type (x EID) EID { 
    return EID{F_Optimize_ptype_type(ToType(OBJ(x)) ).Id(),0}} 
  
// v3.1.06: member -> always apply to a ptype
/* The go function for: pmember(x:type) [status=0] */
func F_Optimize_pmember_type (x *ClaireType ) *ClaireType  { 
    return  Core.F_member_type(F_Optimize_ptype_type(x))
    } 
  
// The EID go function for: pmember @ type (throw: false) 
func E_Optimize_pmember_type (x EID) EID { 
    return EID{F_Optimize_pmember_type(ToType(OBJ(x)) ).Id(),0}} 
  
// transform an instruction representing a set into an instruction
// representing an enumeration
/* The go function for: enumerate_code(self:any,%t:type) [status=1] */
func F_Optimize_enumerate_code_any (self *ClaireAny ,_Zt *ClaireType ) EID { 
    // eid body s = any
    var Result EID 
    if ((F_Optimize_ptype_type(_Zt).Included(ToType(C_list.Id())) == CTRUE) || 
        ((F_Optimize_ptype_type(_Zt).Included(ToType(C_set.Id())) == CTRUE) || 
          (F_Optimize_ptype_type(_Zt).Included(ToType(C_tuple.Id())) == CTRUE))) { 
      Result = F_Compile_c_strict_code_any(self,F_Optimize_ptype_type(_Zt).Class_I())
      } else {
      if (C_compiler.Optimize_ask == CTRUE) { 
        F_Compile_notice_void()
        
        } 
      Result = F_Optimize_c_code_method_method1(ToMethod(Core.F__at_property1(Core.C_Core_enumerate,C_any).Id()),MakeConstantList(self),MakeConstantList(_Zt.Id()))
      } 
    return Result} 
  
// The EID go function for: enumerate_code @ any (throw: true) 
func E_Optimize_enumerate_code_any (self EID,_Zt EID) EID { 
    return F_Optimize_enumerate_code_any(ANY(self),ToType(OBJ(_Zt)) )} 
  
// range inference for a "for" structure: y is the new type and ts is the type of
// the collection structure. Note that except for the case of float arrays, the
// sort of the collection is assumed to be any or integer (thus we "correct" the
// type inference with sort_abstract)
/* The go function for: range_infers_for(self:Variable,y:type,ts:type) [status=0] */
func F_Optimize_range_infers_for_Variable (self *ClaireVariable ,y *ClaireType ,ts *ClaireType ) *ClaireAny  { 
    // procedure body with s = any
    var Result *ClaireAny  
    if (self.Range.Id() == CNULL) { 
      
      if (y.Isa.IsIn(C_Interval) == CTRUE) { 
        y = ToType(C_integer.Id())
        } 
      { 
        var va_arg1 *ClaireVariable  
        var va_arg2 *ClaireType  
        va_arg1 = self
        va_arg2 = y
        va_arg1.Range = va_arg2
        Result = va_arg2.Id()
        } 
      }  else if ((y.Included(self.Range) != CTRUE) && 
        (C_compiler.Safety <= 1)) { 
      if ((F_boolean_I_any(y.Id()) != CTRUE) || 
          (F_boolean_I_any(self.Range.Id()) != CTRUE)) { 
        F_Compile_warn_void()
        Result = ANY(Core.F_tformat_string(MakeString("range of variable in ~S is wrong [258]\n"),1,MakeConstantList(self.Id())))
        } else {
        Result = CFALSE.Id()
        } 
      } else {
      Result = CFALSE.Id()
      } 
    return Result} 
  
// The EID go function for: range_infers_for @ Variable (throw: false) 
func E_Optimize_range_infers_for_Variable (self EID,y EID,ts EID) EID { 
    return F_Optimize_range_infers_for_Variable(To_Variable(OBJ(self)),ToType(OBJ(y)),ToType(OBJ(ts)) ).ToEID()} 
  
// v3.1.06: remove complains because it traps the compiler's own inferences
// to reintroduce, we need to distinguish between user and compiler
// types for iteration variables !
// if (sort(self) != any & (sort(self) != integer | compiler.overflow?) &
//    not(ts <= array & y <= float))               // iteration of float array is a special case
//  (//[5] protect original sort with ~S // sort_abstract!(self.range),
//   put(range, self, sort_abstract!(self.range))) ]
// variable range inference, how to guess a type from the value ...
/* The go function for: range_infers(self:Variable,y:type) [status=1] */
func F_Optimize_range_infers_Variable (self *ClaireVariable ,y *ClaireType ) EID { 
    // eid body s = any
    var Result EID 
    var g0122I *ClaireBoolean  
    var try_1 EID 
    { 
      var v_or2 *ClaireBoolean  
      
      v_or2 = MakeBoolean((self.Range.Id() == CNULL))
      if (v_or2 == CTRUE) {try_1 = EID{CTRUE.Id(),0}
      } else { 
        var try_2 EID 
        try_2 = F_Optimize_extended_ask_type(self.Range)
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
    g0122I = ToBoolean(OBJ(try_1))
    if (g0122I == CTRUE) { 
      if (C_set.Id() == y.Isa.Id()) { 
        { 
          var va_arg1 *ClaireVariable  
          var va_arg2 *ClaireType  
          va_arg1 = self
          va_arg2 = ToType(y.Class_I().Id())
          va_arg1.Range = va_arg2
          Result = EID{va_arg2.Id(),0}
          } 
        } else {
        { 
          var va_arg1 *ClaireVariable  
          var va_arg2 *ClaireType  
          va_arg1 = self
          va_arg2 = y
          va_arg1.Range = va_arg2
          Result = EID{va_arg2.Id(),0}
          } 
        } 
      } else {
      Result = EID{CFALSE.Id(),0}
      } 
    }
    return Result} 
  
// The EID go function for: range_infers @ Variable (throw: true) 
func E_Optimize_range_infers_Variable (self EID,y EID) EID { 
    return F_Optimize_range_infers_Variable(To_Variable(OBJ(self)),ToType(OBJ(y)) )} 
  
// temporary range inference for case, which may use a special form:
// {any U type} to represent the change of sort
// {} U (c U t) to represent a change of psort
/* The go function for: range_sets(self:any,y:type) [status=0] */
func F_Optimize_range_sets_any (self *ClaireAny ,y *ClaireType )  { 
    // procedure body with s = void
    if (self.Isa.IsIn(C_Variable) == CTRUE) { 
      { var g0123 *ClaireVariable   = To_Variable(self)
        g0123.Range = y
        } 
      } 
    } 
  
// The EID go function for: range_sets @ any (throw: false) 
func E_Optimize_range_sets_any (self EID,y EID) EID { 
    F_Optimize_range_sets_any(ANY(self),ToType(OBJ(y)) )
    return EVOID} 
  
//
// the srange of a method = class!(range)
/* The go function for: c_srange(m:method) [status=1] */
func F_Optimize_c_srange_method (m *ClaireMethod ) EID { 
    // eid body s = class
    var Result EID 
    Result = Core.F_last_list(m.Srange)
    return Result} 
  
// The EID go function for: c_srange @ method (throw: true) 
func E_Optimize_c_srange_method (m EID) EID { 
    return F_Optimize_c_srange_method(ToMethod(OBJ(m)) )} 
  
// v3.3 some of the global variables are compiled with a native var approach
// we require the range to be safe, no backtrack & local global var
/* The go function for: Compile/nativeVar?(x:global_variable) [status=0] */
func F_Compile_nativeVar_ask_global_variable (x *Core.GlobalVariable ) *ClaireBoolean  { 
    if ((C_compiler.Optimize_ask == CTRUE) && 
        ((x.Store_ask.Id() == CFALSE.Id()) && 
          (x.Name.Module_I().Id() == x.Name.Defined().Id()))) {return CTRUE
    } else {return CFALSE}} 
  
// The EID go function for: Compile/nativeVar? @ global_variable (throw: false) 
func E_Compile_nativeVar_ask_global_variable (x EID) EID { 
    return EID{F_Compile_nativeVar_ask_global_variable(Core.ToGlobalVariable(OBJ(x)) ).Id(),0}} 
  
// v3.3 finds the possible return type of a block (within a loop)
// it returns a class for the time being ...
/* The go function for: Compile/return_type(self:any) [status=1] */
func F_Compile_return_type_any (self *ClaireAny ) EID { 
    // eid body s = type
    var Result EID 
    if (self.Isa.IsIn(Language.C_Let) == CTRUE) { 
      { var g0124 *Language.Let   = Language.To_Let(self)
        Result = F_Compile_return_type_any(g0124.Arg)
        } 
      }  else if (self.Isa.IsIn(Language.C_Do) == CTRUE) { 
      { var g0125 *Language.Do   = Language.To_Do(self)
        { var x *ClaireSet   = CEMPTY
          { 
            var y *ClaireAny  
            _ = y
            Result= EID{CFALSE.Id(),0}
            var y_support *ClaireList  
            y_support = g0125.Args
            y_len := y_support.Length()
            for i_it := 0; i_it < y_len; i_it++ { 
              y = y_support.At(i_it)
              var loop_1 EID 
              _ = loop_1
              var try_2 EID 
              { var arg_3 *ClaireType  
                var try_4 EID 
                try_4 = F_Compile_return_type_any(y)
                if ErrorIn(try_4) {try_2 = try_4
                } else {
                arg_3 = ToType(OBJ(try_4))
                try_2 = Core.F_CALL(ToProperty(C__exp.Id()),ARGS(EID{x.Id(),0},EID{arg_3.Id(),0}))
                }
                } 
              if ErrorIn(try_2) {Result = try_2
              break
              } else {
              x = ToSet(OBJ(try_2))
              loop_1 = EID{x.Id(),0}
              }
              } 
            } 
          if !ErrorIn(Result) {
          Result = EID{x.Id(),0}
          }
          } 
        } 
      }  else if (self.Isa.IsIn(Language.C_If) == CTRUE) { 
      { var g0126 *Language.If   = Language.To_If(self)
        { var arg_5 *ClaireType  
          var try_7 EID 
          try_7 = F_Compile_return_type_any(g0126.Arg)
          if ErrorIn(try_7) {Result = try_7
          } else {
          arg_5 = ToType(OBJ(try_7))
          { var arg_6 *ClaireType  
            var try_8 EID 
            try_8 = F_Compile_return_type_any(g0126.Other)
            if ErrorIn(try_8) {Result = try_8
            } else {
            arg_6 = ToType(OBJ(try_8))
            Result = Core.F_CALL(ToProperty(C__exp.Id()),ARGS(EID{arg_5.Id(),0},EID{arg_6.Id(),0}))
            }
            } 
          }
          } 
        } 
      }  else if (self.Isa.IsIn(Language.C_Return) == CTRUE) { 
      { var g0127 *Language.Return   = Language.To_Return(self)
        Result = Core.F_CALL(C_c_type,ARGS(g0127.Arg.ToEID()))
        } 
      }  else if (self.Isa.IsIn(Language.C_Case) == CTRUE) { 
      { var g0128 *Language.Case   = Language.To_Case(self)
        { var x *ClaireSet   = CEMPTY
          { 
            var y *ClaireAny  
            _ = y
            Result= EID{CFALSE.Id(),0}
            var y_support *ClaireList  
            y_support = g0128.Args
            y_len := y_support.Length()
            for i_it := 0; i_it < y_len; i_it++ { 
              y = y_support.At(i_it)
              var loop_9 EID 
              _ = loop_9
              var try_10 EID 
              { var arg_11 *ClaireType  
                var try_12 EID 
                try_12 = F_Compile_return_type_any(y)
                if ErrorIn(try_12) {try_10 = try_12
                } else {
                arg_11 = ToType(OBJ(try_12))
                try_10 = Core.F_CALL(ToProperty(C__exp.Id()),ARGS(EID{x.Id(),0},EID{arg_11.Id(),0}))
                }
                } 
              if ErrorIn(try_10) {Result = try_10
              break
              } else {
              x = ToSet(OBJ(try_10))
              loop_9 = EID{x.Id(),0}
              }
              } 
            } 
          if !ErrorIn(Result) {
          Result = EID{x.Id(),0}
          }
          } 
        } 
      }  else if (self.Isa.IsIn(Language.C_Handle) == CTRUE) { 
      { var g0129 *Language.ClaireHandle   = Language.To_ClaireHandle(self)
        Result = F_Compile_return_type_any(g0129.Arg)
        } 
      } else {
      Result = EID{CEMPTY.Id(),0}
      } 
    return Result} 
  
// The EID go function for: Compile/return_type @ any (throw: true) 
func E_Compile_return_type_any (self EID) EID { 
    return F_Compile_return_type_any(ANY(self) )} 
  
// compiling a type expression --------------------------------------------
//
// creates the functional code that produce the code by evaluation
// note this is expensive -> we should encourage the use of global variables
/* The go function for: c_code(self:((type_operator U Reference) U Pattern),s:class) [status=1] */
func F_c_code_type_expression (self *ClaireTypeExpression ,s *ClaireClass ) EID { 
    // eid body s = any
    var Result EID 
    { var arg_1 *ClaireAny  
      var try_2 EID 
      try_2 = Core.F_CALL(C_Compile_self_code,ARGS(EID{self.Id(),0}))
      if ErrorIn(try_2) {Result = try_2
      } else {
      arg_1 = ANY(try_2)
      Result = Core.F_CALL(C_c_code,ARGS(arg_1.ToEID(),EID{s.Id(),0}))
      }
      } 
    return Result} 
  
// The EID go function for: c_code @ type_expression (throw: true) 
func E_c_code_type_expression (self EID,s EID) EID { 
    return F_c_code_type_expression(ToTypeExpression(OBJ(self)),ToClass(OBJ(s)) )} 
  
// to check - seems OK for 3.2 !
/* The go function for: Compile/self_code(self:subtype) [status=1] */
func F_Compile_self_code_subtype (self *ClaireSubtype ) EID { 
    // eid body s = any
    var Result EID 
    { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
      _CL_obj.Selector = C_nth
      { 
        var va_arg1 *Language.Call  
        var va_arg2 *ClaireList  
        va_arg1 = _CL_obj
        var try_1 EID 
        { 
          var v_bag_arg *ClaireAny  
          try_1= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
          ToList(OBJ(try_1)).AddFast(self.Arg.Id())
          var try_2 EID 
          try_2 = Core.F_CALL(C_c_code,ARGS(EID{self.T1.Id(),0},EID{C_type.Id(),0}))
          if ErrorIn(try_2) {try_1 = try_2
          } else {
          v_bag_arg = ANY(try_2)
          ToList(OBJ(try_1)).AddFast(v_bag_arg)}
          } 
        if ErrorIn(try_1) {Result = try_1
        } else {
        va_arg2 = ToList(OBJ(try_1))
        va_arg1.Args = va_arg2
        Result = EID{va_arg2.Id(),0}
        }
        } 
      if !ErrorIn(Result) {
      Result = EID{_CL_obj.Id(),0}
      }
      } 
    return Result} 
  
// The EID go function for: Compile/self_code @ subtype (throw: true) 
func E_Compile_self_code_subtype (self EID) EID { 
    return F_Compile_self_code_subtype(ToSubtype(OBJ(self)) )} 
  
// create a Param. Optimized in v3.2.28 for list<X>
/* The go function for: Compile/self_code(self:Param) [status=1] */
func F_Compile_self_code_Param (self *ClaireParam ) EID { 
    // eid body s = any
    var Result EID 
    if ((self.Params.Length() == 1) && 
        ((self.Params.At(1-1) == C_of.Id()) && 
          (C_set.Id() == self.Args.At(1-1).Isa.Id()))) { 
      { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
        _CL_obj.Selector = Core.C_Core_param_I
        { 
          var va_arg1 *Language.Call  
          var va_arg2 *ClaireList  
          va_arg1 = _CL_obj
          var try_1 EID 
          { 
            var v_bag_arg *ClaireAny  
            try_1= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
            ToList(OBJ(try_1)).AddFast(self.Arg.Id())
            var try_2 EID 
            { var arg_3 *ClaireAny  
              var try_4 EID 
              try_4 = Core.F_the_type(ToType(self.Args.At(1-1)))
              if ErrorIn(try_4) {try_2 = try_4
              } else {
              arg_3 = ANY(try_4)
              try_2 = Core.F_CALL(C_c_code,ARGS(arg_3.ToEID(),EID{C_type.Id(),0}))
              }
              } 
            if ErrorIn(try_2) {try_1 = try_2
            } else {
            v_bag_arg = ANY(try_2)
            ToList(OBJ(try_1)).AddFast(v_bag_arg)}
            } 
          if ErrorIn(try_1) {Result = try_1
          } else {
          va_arg2 = ToList(OBJ(try_1))
          va_arg1.Args = va_arg2
          Result = EID{va_arg2.Id(),0}
          }
          } 
        if !ErrorIn(Result) {
        Result = EID{_CL_obj.Id(),0}
        }
        } 
      } else {
      { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
        _CL_obj.Selector = C_nth
        { 
          var va_arg1 *Language.Call  
          var va_arg2 *ClaireList  
          va_arg1 = _CL_obj
          var try_5 EID 
          { 
            var v_bag_arg *ClaireAny  
            try_5= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
            ToList(OBJ(try_5)).AddFast(self.Arg.Id())
            ToList(OBJ(try_5)).AddFast(self.Params.Id())
            var try_6 EID 
            { 
              var v_list6 *ClaireList  
              var y *ClaireAny  
              var v_local6 *ClaireAny  
              v_list6 = self.Args
              try_6 = EID{CreateList(ToType(CEMPTY.Id()),v_list6.Length()).Id(),0}
              for CLcount := 0; CLcount < v_list6.Length(); CLcount++{ 
                y = v_list6.At(CLcount)
                var try_7 EID 
                try_7 = Core.F_CALL(C_c_code,ARGS(y.ToEID(),EID{C_type.Id(),0}))
                if ErrorIn(try_7) {try_6 = try_7
                break
                } else {
                v_local6 = ANY(try_7)
                ToList(OBJ(try_6)).PutAt(CLcount,v_local6)
                } 
              }
              } 
            if ErrorIn(try_6) {try_5 = try_6
            } else {
            v_bag_arg = ANY(try_6)
            ToList(OBJ(try_5)).AddFast(v_bag_arg)}
            } 
          if ErrorIn(try_5) {Result = try_5
          } else {
          va_arg2 = ToList(OBJ(try_5))
          va_arg1.Args = va_arg2
          Result = EID{va_arg2.Id(),0}
          }
          } 
        if !ErrorIn(Result) {
        Result = EID{_CL_obj.Id(),0}
        }
        } 
      } 
    return Result} 
  
// The EID go function for: Compile/self_code @ Param (throw: true) 
func E_Compile_self_code_Param (self EID) EID { 
    return F_Compile_self_code_Param(To_Param(OBJ(self)) )} 
  
/* The go function for: Compile/self_code(self:Union) [status=1] */
func F_Compile_self_code_Union (self *ClaireUnion ) EID { 
    // eid body s = any
    var Result EID 
    { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
      _CL_obj.Selector = ToProperty(Core.C_U.Id())
      { 
        var va_arg1 *Language.Call  
        var va_arg2 *ClaireList  
        va_arg1 = _CL_obj
        var try_1 EID 
        { 
          var v_bag_arg *ClaireAny  
          try_1= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
          var try_2 EID 
          try_2 = Core.F_CALL(C_c_code,ARGS(EID{self.T1.Id(),0},EID{C_type.Id(),0}))
          if ErrorIn(try_2) {try_1 = try_2
          } else {
          v_bag_arg = ANY(try_2)
          ToList(OBJ(try_1)).AddFast(v_bag_arg)
          var try_3 EID 
          try_3 = Core.F_CALL(C_c_code,ARGS(EID{self.T2.Id(),0},EID{C_type.Id(),0}))
          if ErrorIn(try_3) {try_1 = try_3
          } else {
          v_bag_arg = ANY(try_3)
          ToList(OBJ(try_1)).AddFast(v_bag_arg)}}
          } 
        if ErrorIn(try_1) {Result = try_1
        } else {
        va_arg2 = ToList(OBJ(try_1))
        va_arg1.Args = va_arg2
        Result = EID{va_arg2.Id(),0}
        }
        } 
      if !ErrorIn(Result) {
      Result = EID{_CL_obj.Id(),0}
      }
      } 
    return Result} 
  
// The EID go function for: Compile/self_code @ Union (throw: true) 
func E_Compile_self_code_Union (self EID) EID { 
    return F_Compile_self_code_Union(To_Union(OBJ(self)) )} 
  
/* The go function for: Compile/self_code(self:Interval) [status=0] */
func F_Compile_self_code_Interval (self *ClaireInterval ) *ClaireAny  { 
    // procedure body with s = any
    var Result *ClaireAny  
    { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
      _CL_obj.Selector = ToProperty(C__dot_dot.Id())
      _CL_obj.Args = MakeConstantList(MakeInteger(self.Arg1).Id(),MakeInteger(self.Arg2).Id())
      Result = _CL_obj.Id()
      } 
    return Result} 
  
// The EID go function for: Compile/self_code @ Interval (throw: false) 
func E_Compile_self_code_Interval (self EID) EID { 
    return F_Compile_self_code_Interval(To_Interval(OBJ(self)) ).ToEID()} 
  
/* The go function for: Compile/self_code(self:Reference) [status=0] */
func F_Compile_self_code_Reference (self *ClaireReference ) *ClaireAny  { 
    // procedure body with s = any
    var Result *ClaireAny  
    { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
      _CL_obj.Selector = Core.C_Reference_I
      _CL_obj.Args = MakeConstantList(self.Args.Id(),MakeInteger(self.Index).Id())
      Result = _CL_obj.Id()
      } 
    return Result} 
  
// The EID go function for: Compile/self_code @ Reference (throw: false) 
func E_Compile_self_code_Reference (self EID) EID { 
    return F_Compile_self_code_Reference(To_Reference(OBJ(self)) ).ToEID()} 
  
// compilation of a Pattern
/* The go function for: Compile/self_code(self:Pattern) [status=0] */
func (self *ClairePattern ) SelfCode () *ClaireAny  { 
    // procedure body with s = any
    var Result *ClaireAny  
    if (C_compiler.Inline_ask == CTRUE) { 
      { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
        _CL_obj.Selector = C_nth
        { 
          var va_arg1 *Language.Call  
          var va_arg2 *ClaireList  
          va_arg1 = _CL_obj
          { 
            var v_bag_arg *ClaireAny  
            va_arg2= ToType(CEMPTY.Id()).EmptyList()
            va_arg2.AddFast(self.Selector.Id())
            { var _CL_obj *Language.Tuple   = Language.To_Tuple(new(Language.Tuple).Is(Language.C_Tuple))
              _CL_obj.Args = self.Arg
              v_bag_arg = _CL_obj.Id()
              } 
            va_arg2.AddFast(v_bag_arg)} 
          va_arg1.Args = va_arg2
          } 
        Result = _CL_obj.Id()
        } 
      } else {
      Result = Language.C_Call.Id()
      } 
    return Result} 
  
// The EID go function for: Compile/self_code @ Pattern (throw: false) 
func E_Compile_self_code_Pattern (self EID) EID { 
    return To_ClairePattern(OBJ(self)).SelfCode( ).ToEID()} 
  
//-------------- membership compiling -------------------------------
// membership to a class : for final and closed classes => x.isa = c !
/* The go function for: member_code(self:class,x:any) [status=1] */
func F_Optimize_member_code_class (self *ClaireClass ,x *ClaireAny ) EID { 
    // eid body s = any
    var Result EID 
    { var _Zxt *Language.Call  
      var try_1 EID 
      { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
        { 
          var va_arg1 *Language.Call  
          var va_arg2 *ClaireProperty  
          va_arg1 = _CL_obj
          var try_2 EID 
          var g0131I *ClaireBoolean  
          var try_3 EID 
          { var arg_4 *ClaireType  
            var try_5 EID 
            try_5 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
            if ErrorIn(try_5) {try_3 = try_5
            } else {
            arg_4 = ToType(OBJ(try_5))
            try_3 = EID{arg_4.Included(ToType(C_object.Id())).Id(),0}
            }
            } 
          if ErrorIn(try_3) {try_2 = try_3
          } else {
          g0131I = ToBoolean(OBJ(try_3))
          if (g0131I == CTRUE) { 
            try_2 = EID{C_isa.Id(),0}
            } else {
            try_2 = EID{Core.C_owner.Id(),0}
            } 
          }
          if ErrorIn(try_2) {try_1 = try_2
          } else {
          va_arg2 = ToProperty(OBJ(try_2))
          va_arg1.Selector = va_arg2
          try_1 = EID{va_arg2.Id(),0}
          }
          } 
        if !ErrorIn(try_1) {
        _CL_obj.Args = MakeConstantList(x)
        try_1 = EID{_CL_obj.Id(),0}
        }
        } 
      if ErrorIn(try_1) {Result = try_1
      } else {
      _Zxt = Language.To_Call(OBJ(try_1))
      if (((self.Open <= -1) || 
            (self.Open == 1)) && 
          (F_boolean_I_any(self.Subclass.Id()).Id() != CTRUE.Id())) { 
        { var arg_6 *Language.Call  
          { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
            _CL_obj.Selector = ToProperty(C__equal.Id())
            _CL_obj.Args = MakeConstantList(self.Id(),_Zxt.Id())
            arg_6 = _CL_obj
            } 
          Result = Core.F_CALL(C_c_code,ARGS(EID{arg_6.Id(),0}))
          } 
        } else {
        { var arg_7 *Language.Call  
          { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
            _CL_obj.Selector = ToProperty(Core.C_inherit_ask.Id())
            _CL_obj.Args = MakeConstantList(_Zxt.Id(),self.Id())
            arg_7 = _CL_obj
            } 
          Result = Core.F_CALL(C_c_code,ARGS(EID{arg_7.Id(),0}))
          } 
        } 
      }
      } 
    return Result} 
  
// The EID go function for: member_code @ class (throw: true) 
func E_Optimize_member_code_class (self EID,x EID) EID { 
    return F_Optimize_member_code_class(ToClass(OBJ(self)),ANY(x) )} 
  
/* The go function for: member_code(self:type_operator,x:any) [status=1] */
func F_Optimize_member_code_type_operator (self *ClaireTypeOperator ,x *ClaireAny ) EID { 
    // eid body s = any
    var Result EID 
    { var _CL_obj *Language.CallMethod2   = Language.To_CallMethod2(new(Language.CallMethod2).Is(Language.C_Call_method2))
      _CL_obj.Arg = ToMethod(Core.F__at_property2(ToProperty(C__Z.Id()),MakeConstantList(C_any.Id(),C_any.Id())).Id())
      { 
        var va_arg1 *Language.CallMethod  
        var va_arg2 *ClaireList  
        va_arg1 = Language.To_CallMethod(_CL_obj.Id())
        var try_1 EID 
        { 
          var v_bag_arg *ClaireAny  
          try_1= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
          var try_2 EID 
          try_2 = Core.F_CALL(C_c_code,ARGS(x.ToEID(),EID{C_any.Id(),0}))
          if ErrorIn(try_2) {try_1 = try_2
          } else {
          v_bag_arg = ANY(try_2)
          ToList(OBJ(try_1)).AddFast(v_bag_arg)
          var try_3 EID 
          try_3 = Core.F_CALL(C_c_code,ARGS(EID{self.Id(),0},EID{C_any.Id(),0}))
          if ErrorIn(try_3) {try_1 = try_3
          } else {
          v_bag_arg = ANY(try_3)
          ToList(OBJ(try_1)).AddFast(v_bag_arg)}}
          } 
        if ErrorIn(try_1) {Result = try_1
        } else {
        va_arg2 = ToList(OBJ(try_1))
        va_arg1.Args = va_arg2
        Result = EID{va_arg2.Id(),0}
        }
        } 
      if !ErrorIn(Result) {
      Result = EID{_CL_obj.Id(),0}
      }
      } 
    return Result} 
  
// The EID go function for: member_code @ type_operator (throw: true) 
func E_Optimize_member_code_type_operator (self EID,x EID) EID { 
    return F_Optimize_member_code_type_operator(ToTypeOperator(OBJ(self)),ANY(x) )} 
  
/* The go function for: member_code(self:Union,x:any) [status=1] */
func F_Optimize_member_code_Union (self *ClaireUnion ,x *ClaireAny ) EID { 
    // eid body s = any
    var Result EID 
    { var _CL_obj *Language.Or   = Language.To_Or(new(Language.Or).Is(Language.C_Or))
      { 
        var va_arg1 *Language.Or  
        var va_arg2 *ClaireList  
        va_arg1 = _CL_obj
        var try_1 EID 
        { 
          var v_bag_arg *ClaireAny  
          try_1= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
          var try_2 EID 
          try_2 = Core.F_CALL(C_Optimize_member_code,ARGS(EID{self.T1.Id(),0},x.ToEID()))
          if ErrorIn(try_2) {try_1 = try_2
          } else {
          v_bag_arg = ANY(try_2)
          ToList(OBJ(try_1)).AddFast(v_bag_arg)
          var try_3 EID 
          try_3 = Core.F_CALL(C_Optimize_member_code,ARGS(EID{self.T2.Id(),0},x.ToEID()))
          if ErrorIn(try_3) {try_1 = try_3
          } else {
          v_bag_arg = ANY(try_3)
          ToList(OBJ(try_1)).AddFast(v_bag_arg)}}
          } 
        if ErrorIn(try_1) {Result = try_1
        } else {
        va_arg2 = ToList(OBJ(try_1))
        va_arg1.Args = va_arg2
        Result = EID{va_arg2.Id(),0}
        }
        } 
      if !ErrorIn(Result) {
      Result = EID{_CL_obj.Id(),0}
      }
      } 
    return Result} 
  
// The EID go function for: member_code @ Union (throw: true) 
func E_Optimize_member_code_Union (self EID,x EID) EID { 
    return F_Optimize_member_code_Union(To_Union(OBJ(self)),ANY(x) )} 
  
/* The go function for: member_code(self:Interval,x:any) [status=1] */
func F_Optimize_member_code_Interval (self *ClaireInterval ,x *ClaireAny ) EID { 
    // eid body s = any
    var Result EID 
    { var arg_1 *Language.And  
      { var _CL_obj *Language.And   = Language.To_And(new(Language.And).Is(Language.C_And))
        { 
          var va_arg1 *Language.And  
          var va_arg2 *ClaireList  
          va_arg1 = _CL_obj
          { 
            var v_bag_arg *ClaireAny  
            va_arg2= ToType(CEMPTY.Id()).EmptyList()
            { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
              _CL_obj.Selector = ToProperty(C__sup_equal.Id())
              _CL_obj.Args = MakeConstantList(x,MakeInteger(self.Arg1).Id())
              v_bag_arg = _CL_obj.Id()
              } 
            va_arg2.AddFast(v_bag_arg)
            { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
              _CL_obj.Selector = ToProperty(C__inf_equal.Id())
              _CL_obj.Args = MakeConstantList(x,MakeInteger(self.Arg2).Id())
              v_bag_arg = _CL_obj.Id()
              } 
            va_arg2.AddFast(v_bag_arg)} 
          va_arg1.Args = va_arg2
          } 
        arg_1 = _CL_obj
        } 
      Result = Core.F_CALL(C_c_code,ARGS(EID{arg_1.Id(),0},EID{C_any.Id(),0}))
      } 
    return Result} 
  
// The EID go function for: member_code @ Interval (throw: true) 
func E_Optimize_member_code_Interval (self EID,x EID) EID { 
    return F_Optimize_member_code_Interval(To_Interval(OBJ(self)),ANY(x) )} 
  
/* The go function for: member_code(self:Param,x:any) [status=1] */
func F_Optimize_member_code_Param (self *ClaireParam ,x *ClaireAny ) EID { 
    // eid body s = any
    var Result EID 
    { var arg_1 *Language.And  
      { var _CL_obj *Language.And   = Language.To_And(new(Language.And).Is(Language.C_And))
        { 
          var va_arg1 *Language.And  
          var va_arg2 *ClaireList  
          va_arg1 = _CL_obj
          { var arg_2 *ClaireList  
            { 
              var v_bag_arg *ClaireAny  
              arg_2= ToType(CEMPTY.Id()).EmptyList()
              { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                _CL_obj.Selector = ToProperty(C__Z.Id())
                _CL_obj.Args = MakeConstantList(x,self.Arg.Id())
                v_bag_arg = _CL_obj.Id()
                } 
              arg_2.AddFast(v_bag_arg)} 
            { var arg_3 *ClaireList  
              { var i_bag *ClaireList   = ToType(CEMPTY.Id()).EmptyList()
                { var i int  = 1
                  { var g0132 int  = self.Params.Length()
                    for (i <= g0132) { 
                      { var arg_4 *Language.Call  
                        { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                          _CL_obj.Selector = ToProperty(C__Z.Id())
                          { 
                            var va_arg1 *Language.Call  
                            var va_arg2 *ClaireList  
                            va_arg1 = _CL_obj
                            { 
                              var v_bag_arg *ClaireAny  
                              va_arg2= ToType(CEMPTY.Id()).EmptyList()
                              { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                                _CL_obj.Selector = ToProperty(self.Params.At(i-1))
                                _CL_obj.Args = MakeConstantList(x)
                                v_bag_arg = _CL_obj.Id()
                                } 
                              va_arg2.AddFast(v_bag_arg)
                              va_arg2.AddFast(self.Args.At(i-1))} 
                            va_arg1.Args = va_arg2
                            } 
                          arg_4 = _CL_obj
                          } 
                        i_bag.AddFast(arg_4.Id())
                        } 
                      i = (i+1)
                      } 
                    } 
                  } 
                arg_3 = i_bag
                } 
              va_arg2 = arg_2.Append(arg_3)
              } 
            } 
          va_arg1.Args = va_arg2
          } 
        arg_1 = _CL_obj
        } 
      Result = Core.F_CALL(C_c_code,ARGS(EID{arg_1.Id(),0},EID{C_any.Id(),0}))
      } 
    return Result} 
  
// The EID go function for: member_code @ Param (throw: true) 
func E_Optimize_member_code_Param (self EID,x EID) EID { 
    return F_Optimize_member_code_Param(To_Param(OBJ(self)),ANY(x) )} 
  
// v3.3.14: specialized code for tuple
/* The go function for: member_code(self:tuple,x:any) [status=1] */
func F_Optimize_member_code_tuple (self *ClaireTuple ,x *ClaireAny ) EID { 
    // eid body s = any
    var Result EID 
    if (x.Isa.IsIn(Language.C_Tuple) == CTRUE) { 
      if (ToList(OBJ(Core.F_CALL(C_args,ARGS(x.ToEID())))).Length() != self.Length()) { 
        Result = EID{CFALSE.Id(),0}
        } else {
        { var arg_1 *Language.And  
          { var _CL_obj *Language.And   = Language.To_And(new(Language.And).Is(Language.C_And))
            { 
              var va_arg1 *Language.And  
              var va_arg2 *ClaireList  
              va_arg1 = _CL_obj
              { var i_bag *ClaireList   = ToType(CEMPTY.Id()).EmptyList()
                { var i int  = 1
                  { var g0133 int  = self.Length()
                    for (i <= g0133) { 
                      { var arg_2 *Language.Call  
                        { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                          _CL_obj.Selector = ToProperty(C__Z.Id())
                          _CL_obj.Args = MakeConstantList(ToList(OBJ(Core.F_CALL(C_args,ARGS(x.ToEID())))).At(i-1),ToList(self.Id()).At(i-1))
                          arg_2 = _CL_obj
                          } 
                        i_bag.AddFast(arg_2.Id())
                        } 
                      i = (i+1)
                      } 
                    } 
                  } 
                va_arg2 = i_bag
                } 
              va_arg1.Args = va_arg2
              } 
            arg_1 = _CL_obj
            } 
          Result = Core.F_CALL(C_c_code,ARGS(EID{arg_1.Id(),0},EID{C_any.Id(),0}))
          } 
        } 
      } else {
      Result = F_Optimize_c_code_method_method1(ToMethod(Core.F__at_property2(ToProperty(C__Z.Id()),MakeConstantList(C_any.Id(),C_any.Id())).Id()),MakeConstantList(x,self.Id()),MakeConstantList(C_any.Id(),C_any.Id()))
      } 
    return Result} 
  
// The EID go function for: member_code @ tuple (throw: true) 
func E_Optimize_member_code_tuple (self EID,x EID) EID { 
    return F_Optimize_member_code_tuple(ToTuple(OBJ(self)),ANY(x) )} 
  
/* The go function for: member_code(self:any,x:any) [status=1] */
func F_Optimize_member_code_any (self *ClaireAny ,x *ClaireAny ) EID { 
    // eid body s = any
    var Result EID 
    Language.C_LDEF.Value = CNIL.Id()
    { var _Ztype *ClaireList  
      var try_1 EID 
      { 
        var v_bag_arg *ClaireAny  
        try_1= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
        var try_2 EID 
        try_2 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
        if ErrorIn(try_2) {try_1 = try_2
        } else {
        v_bag_arg = ANY(try_2)
        ToList(OBJ(try_1)).AddFast(v_bag_arg)
        var try_3 EID 
        try_3 = Core.F_CALL(C_c_type,ARGS(self.ToEID()))
        if ErrorIn(try_3) {try_1 = try_3
        } else {
        v_bag_arg = ANY(try_3)
        ToList(OBJ(try_1)).AddFast(v_bag_arg)}}
        } 
      if ErrorIn(try_1) {Result = try_1
      } else {
      _Ztype = ToList(OBJ(try_1))
      { var r *ClaireAny  
        var try_4 EID 
        try_4 = Language.F_extract_pattern_any(self,CNIL)
        if ErrorIn(try_4) {Result = try_4
        } else {
        r = ANY(try_4)
        var g0135I *ClaireBoolean  
        { 
          var v_or4 *ClaireBoolean  
          
          v_or4 = Equal(r,CNULL)
          if (v_or4 == CTRUE) {g0135I = CTRUE
          } else { 
            v_or4 = Equal(self,C_object.Id())
            if (v_or4 == CTRUE) {g0135I = CTRUE
            } else { 
              { var arg_5 *ClaireObject  
                if (self.Isa.IsIn(Core.C_global_variable) == CTRUE) { 
                  { var g0134 *Core.GlobalVariable   = Core.ToGlobalVariable(self)
                    arg_5 = ToObject(g0134.Range.Id())
                    } 
                  } else {
                  arg_5 = ToObject(CFALSE.Id())
                  } 
                v_or4 = F_boolean_I_any(arg_5.Id())
                } 
              if (v_or4 == CTRUE) {g0135I = CTRUE
              } else { 
                g0135I = CFALSE} 
              } 
            } 
          } 
        if (g0135I == CTRUE) { 
          Result = F_Optimize_c_code_method_method1(ToMethod(Core.F__at_property2(ToProperty(C__Z.Id()),_Ztype).Id()),MakeConstantList(x,self),_Ztype)
          } else {
          Result = Core.F_CALL(C_Optimize_member_code,ARGS(r.ToEID(),x.ToEID()))
          } 
        }
        } 
      }
      } 
    return Result} 
  
// The EID go function for: member_code @ any (throw: true) 
func E_Optimize_member_code_any (self EID,x EID) EID { 
    return F_Optimize_member_code_any(ANY(self),ANY(x) )} 
  
// membership optimization though inline definition of %
/* The go function for: %(x:any,y:any) [status=1] */
func F__Z_any4 (x *ClaireAny ,y *ClaireAny ) EID { 
    // eid body s = boolean
    var Result EID 
    { 
      var v_and2 *ClaireBoolean  
      
      var try_1 EID 
      { var arg_2 *ClaireAny  
        var try_3 EID 
        try_3 = EVAL(ToList(OBJ(Core.F_CALL(C_args,ARGS(y.ToEID())))).At(2-1))
        if ErrorIn(try_3) {try_1 = try_3
        } else {
        arg_2 = ANY(try_3)
        try_1 = Core.F_CALL(ToProperty(C__inf_equal.Id()),ARGS(x.ToEID(),arg_2.ToEID()))
        }
        } 
      if ErrorIn(try_1) {Result = try_1
      } else {
      v_and2 = ToBoolean(OBJ(try_1))
      if (v_and2 == CFALSE) {Result = EID{CFALSE.Id(),0}
      } else { 
        var try_4 EID 
        { var arg_5 *ClaireAny  
          var try_6 EID 
          try_6 = EVAL(ToList(OBJ(Core.F_CALL(C_args,ARGS(y.ToEID())))).At(1-1))
          if ErrorIn(try_6) {try_4 = try_6
          } else {
          arg_5 = ANY(try_6)
          try_4 = Core.F_CALL(ToProperty(C__inf_equal.Id()),ARGS(arg_5.ToEID(),x.ToEID()))
          }
          } 
        if ErrorIn(try_4) {Result = try_4
        } else {
        v_and2 = ToBoolean(OBJ(try_4))
        if (v_and2 == CFALSE) {Result = EID{CFALSE.Id(),0}
        } else { 
          Result = EID{CTRUE.Id(),0}} 
        } 
      }}
      } 
    return Result} 
  
// The EID go function for: % @ list<type_expression>(any, ..[tuple(any,any)]) (throw: true) 
func E__Z_any4 (x EID,y EID) EID { 
    return F__Z_any4(ANY(x),ANY(y) )} 
  
/* The go function for: %(x:any,y:any) [status=1] */
func F__Z_any5 (x *ClaireAny ,y *ClaireAny ) EID { 
    // eid body s = boolean
    var Result EID 
    { 
      var v_and2 *ClaireBoolean  
      
      var try_1 EID 
      { var arg_2 *ClaireAny  
        var try_3 EID 
        try_3 = EVAL(ToList(OBJ(Core.F_CALL(C_args,ARGS(y.ToEID())))).At(1-1))
        if ErrorIn(try_3) {try_1 = try_3
        } else {
        arg_2 = ANY(try_3)
        try_1 = Core.F_BELONG(x,arg_2)
        }
        } 
      if ErrorIn(try_1) {Result = try_1
      } else {
      v_and2 = ToBoolean(OBJ(try_1))
      if (v_and2 == CFALSE) {Result = EID{CFALSE.Id(),0}
      } else { 
        var try_4 EID 
        { var arg_5 *ClaireAny  
          var try_6 EID 
          try_6 = EVAL(ToList(OBJ(Core.F_CALL(C_args,ARGS(y.ToEID())))).At(2-1))
          if ErrorIn(try_6) {try_4 = try_6
          } else {
          arg_5 = ANY(try_6)
          try_4 = EID{Core.F__I_equal_any(x,arg_5).Id(),0}
          }
          } 
        if ErrorIn(try_4) {Result = try_4
        } else {
        v_and2 = ToBoolean(OBJ(try_4))
        if (v_and2 == CFALSE) {Result = EID{CFALSE.Id(),0}
        } else { 
          Result = EID{CTRUE.Id(),0}} 
        } 
      }}
      } 
    return Result} 
  
// The EID go function for: % @ list<type_expression>(any, but[tuple(any,any)]) (throw: true) 
func E__Z_any5 (x EID,y EID) EID { 
    return F__Z_any5(ANY(x),ANY(y) )} 
  
// ******************************************************************
// *    Part 4: Miscellaneous                                       *
// ******************************************************************
// ------- variables ------------------------------------------------
/* The go function for: Compile/Variable!(s:symbol,n:integer,t:any) [status=0] */
func F_Compile_Variable_I_symbol (s *ClaireSymbol ,n int,t *ClaireAny ) *ClaireVariable  { 
    // procedure body with s = Variable
    var Result *ClaireVariable  
    if (t.Isa.IsIn(C_type) == CTRUE) { 
      { var g0136 *ClaireType   = ToType(t)
        { var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
          _CL_obj.Pname = s
          _CL_obj.Index = n
          _CL_obj.Range = g0136
          Result = _CL_obj
          } 
        } 
      } else {
      { var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
        _CL_obj.Pname = s
        _CL_obj.Index = n
        Result = _CL_obj
        } 
      } 
    return Result} 
  
// The EID go function for: Compile/Variable! @ symbol (throw: false) 
func E_Compile_Variable_I_symbol (s EID,n EID,t EID) EID { 
    return EID{F_Compile_Variable_I_symbol(ToSymbol(OBJ(s)),INT(n),ANY(t) ).Id(),0}} 
  
/* The go function for: Compile/get_indexed(c:class) [status=0] */
func F_Compile_get_indexed_class (c *ClaireClass ) *ClaireList  { 
    return  c.Slots
    } 
  
// The EID go function for: Compile/get_indexed @ class (throw: false) 
func E_Compile_get_indexed_class (c EID) EID { 
    return EID{F_Compile_get_indexed_class(ToClass(OBJ(c)) ).Id(),0}} 
  
// simple C operations that can be duplicated at no cost {+, -, /, *}
// tells if an expression is a go simply designated object
/* The go function for: Compile/designated?(self:any) [status=1] */
func F_Compile_designated_ask_any (self *ClaireAny ) EID { 
    // eid body s = boolean
    var Result EID 
    { 
      var v_or2 *ClaireBoolean  
      
      v_or2 = self.Isa.IsIn(C_thing)
      if (v_or2 == CTRUE) {Result = EID{CTRUE.Id(),0}
      } else { 
        v_or2 = self.Isa.IsIn(C_Variable)
        if (v_or2 == CTRUE) {Result = EID{CTRUE.Id(),0}
        } else { 
          v_or2 = Equal(C_integer.Id(),self.Isa.Id())
          if (v_or2 == CTRUE) {Result = EID{CTRUE.Id(),0}
          } else { 
            v_or2 = Equal(C_boolean.Id(),self.Isa.Id())
            if (v_or2 == CTRUE) {Result = EID{CTRUE.Id(),0}
            } else { 
              v_or2 = Equal(self,CNIL.Id())
              if (v_or2 == CTRUE) {Result = EID{CTRUE.Id(),0}
              } else { 
                v_or2 = Equal(self,CEMPTY.Id())
                if (v_or2 == CTRUE) {Result = EID{CTRUE.Id(),0}
                } else { 
                  v_or2 = Equal(self,CNULL)
                  if (v_or2 == CTRUE) {Result = EID{CTRUE.Id(),0}
                  } else { 
                    v_or2 = Equal(C_float.Id(),self.Isa.Id())
                    if (v_or2 == CTRUE) {Result = EID{CTRUE.Id(),0}
                    } else { 
                      var try_1 EID 
                      if (self.Isa.IsIn(Language.C_Call) == CTRUE) { 
                        { var g0138 *Language.Call   = Language.To_Call(self)
                          { var x *ClaireAny  
                            var try_2 EID 
                            try_2 = Core.F_CALL(C_c_code,ARGS(EID{g0138.Id(),0}))
                            if ErrorIn(try_2) {try_1 = try_2
                            } else {
                            x = ANY(try_2)
                            { 
                              var v_or14 *ClaireBoolean  
                              
                              var try_3 EID 
                              { 
                                var v_and15 *ClaireBoolean  
                                
                                v_and15 = x.Isa.IsIn(Language.C_Call).Not
                                if (v_and15 == CFALSE) {try_3 = EID{CFALSE.Id(),0}
                                } else { 
                                  var try_4 EID 
                                  try_4 = F_Compile_designated_ask_any(x)
                                  if ErrorIn(try_4) {try_3 = try_4
                                  } else {
                                  v_and15 = ToBoolean(OBJ(try_4))
                                  if (v_and15 == CFALSE) {try_3 = EID{CFALSE.Id(),0}
                                  } else { 
                                    try_3 = EID{CTRUE.Id(),0}} 
                                  } 
                                }
                                } 
                              if ErrorIn(try_3) {try_1 = try_3
                              } else {
                              v_or14 = ToBoolean(OBJ(try_3))
                              if (v_or14 == CTRUE) {try_1 = EID{CTRUE.Id(),0}
                              } else { 
                                v_or14 = Equal(g0138.Selector.Id(),Core.C_mClaire_get_stack.Id())
                                if (v_or14 == CTRUE) {try_1 = EID{CTRUE.Id(),0}
                                } else { 
                                  try_1 = EID{CFALSE.Id(),0}} 
                                } 
                              }
                              } 
                            }
                            } 
                          } 
                        }  else if (self.Isa.IsIn(Language.C_Call_slot) == CTRUE) { 
                        { var g0139 *Language.CallSlot   = Language.To_CallSlot(self)
                          try_1 = F_Compile_designated_ask_any(g0139.Arg)
                          } 
                        }  else if (self.Isa.IsIn(Language.C_Call_table) == CTRUE) { 
                        { var g0140 *Language.CallTable   = Language.To_CallTable(self)
                          try_1 = F_Compile_designated_ask_any(g0140.Arg)
                          } 
                        }  else if (self.Isa.IsIn(Language.C_Call_array) == CTRUE) { 
                        { var g0141 *Language.CallArray   = Language.To_CallArray(self)
                          try_1 = F_Compile_designated_ask_any(g0141.Arg)
                          } 
                        }  else if (self.Isa.IsIn(Language.C_Call_method) == CTRUE) { 
                        { var g0142 *Language.CallMethod   = Language.To_CallMethod(self)
                          { 
                            var v_and13 *ClaireBoolean  
                            
                            v_and13 = MakeBoolean((C_OPT.SimpleOperations.Contain_ask(g0142.Arg.Selector.Id()) == CTRUE) || (g0142.Arg.Id() == Core.F__at_property1(Core.C_unsafe,C_any).Id()) || (g0142.Arg.Id() == Core.F__at_property1(C_nth,C_list).Id()))
                            if (v_and13 == CFALSE) {try_1 = EID{CFALSE.Id(),0}
                            } else { 
                              var try_5 EID 
                              { var arg_6 *ClaireAny  
                                var try_7 EID 
                                { 
                                  var y *ClaireAny  
                                  _ = y
                                  try_7= EID{CFALSE.Id(),0}
                                  var y_support *ClaireList  
                                  y_support = g0142.Args
                                  y_len := y_support.Length()
                                  for i_it := 0; i_it < y_len; i_it++ { 
                                    y = y_support.At(i_it)
                                    var loop_8 EID 
                                    _ = loop_8
                                    var g0144I *ClaireBoolean  
                                    var try_9 EID 
                                    { var arg_10 *ClaireBoolean  
                                      var try_11 EID 
                                      try_11 = F_Compile_designated_ask_any(y)
                                      if ErrorIn(try_11) {try_9 = try_11
                                      } else {
                                      arg_10 = ToBoolean(OBJ(try_11))
                                      try_9 = EID{arg_10.Not.Id(),0}
                                      }
                                      } 
                                    if ErrorIn(try_9) {loop_8 = try_9
                                    } else {
                                    g0144I = ToBoolean(OBJ(try_9))
                                    if (g0144I == CTRUE) { 
                                      try_7 = EID{CTRUE.Id(),0}
                                      break
                                      } else {
                                      loop_8 = EID{CFALSE.Id(),0}
                                      } 
                                    }
                                    if ErrorIn(loop_8) {try_7 = loop_8
                                    break
                                    } else {
                                    }
                                    } 
                                  } 
                                if ErrorIn(try_7) {try_5 = try_7
                                } else {
                                arg_6 = ANY(try_7)
                                try_5 = EID{Core.F_not_any(arg_6).Id(),0}
                                }
                                } 
                              if ErrorIn(try_5) {try_1 = try_5
                              } else {
                              v_and13 = ToBoolean(OBJ(try_5))
                              if (v_and13 == CFALSE) {try_1 = EID{CFALSE.Id(),0}
                              } else { 
                                try_1 = EID{CTRUE.Id(),0}} 
                              } 
                            }
                            } 
                          } 
                        } else {
                        try_1 = EID{CFALSE.Id(),0}
                        } 
                      if ErrorIn(try_1) {Result = try_1
                      } else {
                      v_or2 = ToBoolean(OBJ(try_1))
                      if (v_or2 == CTRUE) {Result = EID{CTRUE.Id(),0}
                      } else { 
                        Result = EID{CFALSE.Id(),0}} 
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
  
// The EID go function for: Compile/designated? @ any (throw: true) 
func E_Compile_designated_ask_any (self EID) EID { 
    return F_Compile_designated_ask_any(ANY(self) )} 
  
// OPT.non_identifiable_set: those sets who are identifiable (closure)
// set<class>{c in class | exists(c2 in c.descendants | c2.ident? = false)})
// equality is identity?
/* The go function for: Compile/identifiable?(self:any) [status=1] */
func F_Compile_identifiable_ask_any (self *ClaireAny ) EID { 
    // eid body s = boolean
    var Result EID 
    { 
      var v_or2 *ClaireBoolean  
      
      v_or2 = Equal(self,CNULL)
      if (v_or2 == CTRUE) {Result = EID{CTRUE.Id(),0}
      } else { 
        var try_1 EID 
        { var t *ClaireClass  
          var try_2 EID 
          { var arg_3 *ClaireType  
            var try_4 EID 
            try_4 = Core.F_CALL(C_c_type,ARGS(self.ToEID()))
            if ErrorIn(try_4) {try_2 = try_4
            } else {
            arg_3 = ToType(OBJ(try_4))
            try_2 = EID{arg_3.Class_I().Id(),0}
            }
            } 
          if ErrorIn(try_2) {try_1 = try_2
          } else {
          t = ToClass(OBJ(try_2))
          try_1 = EID{C_OPT.NonIdentifiableSet.Contain_ask(t.Id()).Not.Id(),0}
          }
          } 
        if ErrorIn(try_1) {Result = try_1
        } else {
        v_or2 = ToBoolean(OBJ(try_1))
        if (v_or2 == CTRUE) {Result = EID{CTRUE.Id(),0}
        } else { 
          Result = EID{CFALSE.Id(),0}} 
        } 
      }
      } 
    return Result} 
  
// The EID go function for: Compile/identifiable? @ any (throw: true) 
func E_Compile_identifiable_ask_any (self EID) EID { 
    return F_Compile_identifiable_ask_any(ANY(self) )} 
  
// inlinning ---------------------------------------------------------
// macro expansion of method self with argument list l
/* The go function for: c_inline(self:method,l:list,s:class) [status=1] */
func F_Optimize_c_inline_method1 (self *ClaireMethod ,l *ClaireList ,s *ClaireClass ) EID { 
    // eid body s = any
    var Result EID 
    
    { var arg_1 *ClaireAny  
      var try_2 EID 
      try_2 = F_Optimize_c_inline_method2(self,l)
      if ErrorIn(try_2) {Result = try_2
      } else {
      arg_1 = ANY(try_2)
      Result = Core.F_CALL(C_c_code,ARGS(arg_1.ToEID(),EID{s.Id(),0}))
      }
      } 
    return Result} 
  
// The EID go function for: c_inline @ list<type_expression>(method, list, class) (throw: true) 
func E_Optimize_c_inline_method1 (self EID,l EID,s EID) EID { 
    return F_Optimize_c_inline_method1(ToMethod(OBJ(self)),ToList(OBJ(l)),ToClass(OBJ(s)) )} 
  
// apply the body of a macro definition
// notice that the name of the inner variables is changed except the second variable
// of iterate macros    
/* The go function for: c_inline(self:method,l:list) [status=1] */
func F_Optimize_c_inline_method2 (self *ClaireMethod ,l *ClaireList ) EID { 
    // eid body s = any
    var Result EID 
    { var f *ClaireLambda   = self.Formula
      { var x *ClaireAny   = f.Body
        { var lbv *ClaireList  
          var try_1 EID 
          try_1 = F_Optimize_bound_variables_any(x)
          if ErrorIn(try_1) {Result = try_1
          } else {
          lbv = ToList(OBJ(try_1))
          { var pv0 *ClaireSymbol  
            if ((self.Selector.Id() == Language.C_iterate.Id()) || 
                (self.Selector.Id() == Language.C_Iterate.Id())) { 
              pv0 = ToSymbol(OBJ(Core.F_CALL(C_mClaire_pname,ARGS(f.Vars.At(2-1).ToEID()))))
              } else {
              pv0 = C_class.Name
              } 
            x = Language.F_instruction_copy_any(x)
            
            { 
              var v *ClaireAny  
              _ = v
              var v_support *ClaireList  
              v_support = lbv
              v_len := v_support.Length()
              for i_it := 0; i_it < v_len; i_it++ { 
                v = v_support.At(i_it)
                { var v2 *ClaireVariable  
                  { var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
                    { 
                      var va_arg1 *ClaireVariable  
                      var va_arg2 *ClaireSymbol  
                      va_arg1 = _CL_obj
                      if (ANY(Core.F_CALL(C_mClaire_pname,ARGS(v.ToEID()))) == pv0.Id()) { 
                        va_arg2 = pv0
                        } else {
                        va_arg2 = Core.F_gensym_void()
                        } 
                      va_arg1.Pname = va_arg2
                      } 
                    _CL_obj.Index = 1000
                    v2 = _CL_obj
                    } 
                  v2.Range = ToType(Core.F_get_property(C_range,ToObject(v)))
                  x = Language.F_substitution_any(x,To_Variable(v),v2.Id())
                  } 
                } 
              } 
            C_OPT.MaxVars = (C_OPT.MaxVars+lbv.Length())
            
            Result = F_Optimize_c_substitution_any(x,f.Vars,l,CFALSE)
            } 
          }
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: c_inline @ list<type_expression>(method, list) (throw: true) 
func E_Optimize_c_inline_method2 (self EID,l EID) EID { 
    return F_Optimize_c_inline_method2(ToMethod(OBJ(self)),ToList(OBJ(l)) )} 
  
// returns the macro expanded code if a macro is involved and nil otherwise
/* The go function for: c_inline_arg?(self:any) [status=1] */
func F_Optimize_c_inline_arg_ask_any (self *ClaireAny ) EID { 
    // eid body s = any
    var Result EID 
    if (self.Isa.IsIn(Language.C_Call) == CTRUE) { 
      { var g0146 *Language.Call   = Language.To_Call(self)
        { var l *ClaireList   = g0146.Args
          { var m *ClaireAny  
            var try_1 EID 
            { var arg_2 *ClaireList  
              var try_3 EID 
              { 
                var v_list7 *ClaireList  
                var x *ClaireAny  
                var v_local7 *ClaireAny  
                v_list7 = l
                try_3 = EID{CreateList(ToType(CEMPTY.Id()),v_list7.Length()).Id(),0}
                for CLcount := 0; CLcount < v_list7.Length(); CLcount++{ 
                  x = v_list7.At(CLcount)
                  var try_4 EID 
                  try_4 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
                  if ErrorIn(try_4) {try_3 = try_4
                  break
                  } else {
                  v_local7 = ANY(try_4)
                  ToList(OBJ(try_3)).PutAt(CLcount,v_local7)
                  } 
                }
                } 
              if ErrorIn(try_3) {try_1 = try_3
              } else {
              arg_2 = ToList(OBJ(try_3))
              try_1 = F_Optimize_restriction_I_property(g0146.Selector,arg_2,CTRUE).ToEID()
              }
              } 
            if ErrorIn(try_1) {Result = try_1
            } else {
            m = ANY(try_1)
            if (C_method.Id() == m.Isa.Id()) { 
              { var g0147 *ClaireMethod   = ToMethod(m)
                var g0150I *ClaireBoolean  
                var try_5 EID 
                { 
                  var v_and8 *ClaireBoolean  
                  
                  v_and8 = g0147.Inline_ask
                  if (v_and8 == CFALSE) {try_5 = EID{CFALSE.Id(),0}
                  } else { 
                    var try_6 EID 
                    try_6 = F_Optimize_c_inline_ask_method(g0147,l)
                    if ErrorIn(try_6) {try_5 = try_6
                    } else {
                    v_and8 = ToBoolean(OBJ(try_6))
                    if (v_and8 == CFALSE) {try_5 = EID{CFALSE.Id(),0}
                    } else { 
                      try_5 = EID{CTRUE.Id(),0}} 
                    } 
                  }
                  } 
                if ErrorIn(try_5) {Result = try_5
                } else {
                g0150I = ToBoolean(OBJ(try_5))
                if (g0150I == CTRUE) { 
                  Result = F_Optimize_c_inline_method2(g0147,l)
                  } else {
                  Result = EID{CFALSE.Id(),0}
                  } 
                }
                } 
              } else {
              Result = EID{CNIL.Id(),0}
              } 
            }
            } 
          } 
        } 
      } else {
      { var arg_7 *Language.Call  
        { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
          _CL_obj.Selector = C_set_I
          _CL_obj.Args = MakeConstantList(self)
          arg_7 = _CL_obj
          } 
        Result = F_Optimize_c_inline_arg_ask_any(arg_7.Id())
        } 
      } 
    return Result} 
  
// The EID go function for: c_inline_arg? @ any (throw: true) 
func E_Optimize_c_inline_arg_ask_any (self EID) EID { 
    return F_Optimize_c_inline_arg_ask_any(ANY(self) )} 
  
// substitute any variable with same name as x with the value val. val is an expression
// when the special form eval() is found, it is "evaluated"
// NEW: in v3.0.5 -> eval(x,C) evals only if x is actually a C
/* The go function for: c_substitution(self:any,lx:list[Variable],val:list,eval?:boolean) [status=1] */
func F_Optimize_c_substitution_any (self *ClaireAny ,lx *ClaireList ,val *ClaireList ,eval_ask *ClaireBoolean ) EID { 
    // eid body s = any
    var Result EID 
    if (self.Isa.IsIn(C_Variable) == CTRUE) { 
      { var g0151 *ClaireVariable   = To_Variable(self)
        { var i *ClaireAny  
          { var j_some *ClaireAny   = CNULL
            { var j int  = 1
              { var g0152 int  = lx.Length()
                for (j <= g0152) { 
                  if (g0151.Pname.Id() == To_Variable(lx.At(j-1)).Pname.Id()) { 
                    j_some = MakeInteger(j).Id()
                    break
                    } 
                  j = (j+1)
                  } 
                } 
              } 
            i = j_some
            } 
          if (i != CNULL) { 
            Result = val.At(ToInteger(i).Value-1).ToEID()
            } else {
            Result = EID{g0151.Id(),0}
            } 
          } 
        } 
      }  else if (self.Isa.IsIn(C_list) == CTRUE) { 
      { var g0153 *ClaireList   = ToList(self)
        { var i int  = 1
          { var g0154 int  = g0153.Length()
            Result= EID{CFALSE.Id(),0}
            for (i <= g0154) { 
              var loop_1 EID 
              _ = loop_1
              { 
              { var arg_2 *ClaireAny  
                var try_3 EID 
                try_3 = F_Optimize_c_substitution_any(g0153.At(i-1),lx,val,eval_ask)
                if ErrorIn(try_3) {loop_1 = try_3
                } else {
                arg_2 = ANY(try_3)
                loop_1 = ToArray(g0153.Id()).NthPut(i,arg_2).ToEID()
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
        Result = EID{g0153.Id(),0}
        }
        } 
      }  else if (self.Isa.IsIn(Language.C_Call) == CTRUE) { 
      { var g0155 *Language.Call   = Language.To_Call(self)
        if (g0155.Selector.Id() == Core.C_eval.Id()) { 
          { var arg_4 *ClaireBoolean  
            var try_5 EID 
            { 
              var v_or6 *ClaireBoolean  
              
              v_or6 = Equal(MakeInteger(g0155.Args.Length()).Id(),MakeInteger(1).Id())
              if (v_or6 == CTRUE) {try_5 = EID{CTRUE.Id(),0}
              } else { 
                var try_6 EID 
                { 
                  var v_and8 *ClaireBoolean  
                  
                  v_and8 = Equal(MakeInteger(g0155.Args.Length()).Id(),MakeInteger(2).Id())
                  if (v_and8 == CFALSE) {try_6 = EID{CFALSE.Id(),0}
                  } else { 
                    var try_7 EID 
                    try_7 = Core.F_BELONG(val.At(1-1),g0155.Args.At(2-1))
                    if ErrorIn(try_7) {try_6 = try_7
                    } else {
                    v_and8 = ToBoolean(OBJ(try_7))
                    if (v_and8 == CFALSE) {try_6 = EID{CFALSE.Id(),0}
                    } else { 
                      try_6 = EID{CTRUE.Id(),0}} 
                    } 
                  }
                  } 
                if ErrorIn(try_6) {try_5 = try_6
                } else {
                v_or6 = ToBoolean(OBJ(try_6))
                if (v_or6 == CTRUE) {try_5 = EID{CTRUE.Id(),0}
                } else { 
                  try_5 = EID{CFALSE.Id(),0}} 
                } 
              }
              } 
            if ErrorIn(try_5) {Result = try_5
            } else {
            arg_4 = ToBoolean(OBJ(try_5))
            Result = F_Optimize_c_substitution_any(g0155.Args.At(1-1),lx,val,arg_4)
            }
            } 
          }  else if (eval_ask == CTRUE) { 
          { 
            h_index := ClEnv.Index
            h_base := ClEnv.Base
            { var arg_8 *ClaireList  
              var try_9 EID 
              { 
                var v_list7 *ClaireList  
                var y *ClaireAny  
                var v_local7 *ClaireAny  
                v_list7 = g0155.Args
                try_9 = EID{CreateList(ToType(CEMPTY.Id()),v_list7.Length()).Id(),0}
                for CLcount := 0; CLcount < v_list7.Length(); CLcount++{ 
                  y = v_list7.At(CLcount)
                  var try_10 EID 
                  try_10 = F_Optimize_c_substitution_any(y,lx,val,CTRUE)
                  if ErrorIn(try_10) {try_9 = try_10
                  break
                  } else {
                  v_local7 = ANY(try_10)
                  ToList(OBJ(try_9)).PutAt(CLcount,v_local7)
                  } 
                }
                } 
              if ErrorIn(try_9) {Result = try_9
              } else {
              arg_8 = ToList(OBJ(try_9))
              Result = Core.F_apply_property(g0155.Selector,arg_8)
              }
              } 
            if ErrorIn(Result){ 
              ClEnv.Index = h_index
              ClEnv.Base = h_base
              Core.F_tformat_string(MakeString("a strange problem happens ~A \n"),0,MakeConstantList(MakeInteger(ClEnv.Verbose).Id()))
              F_Compile_warn_void()
              Core.F_tformat_string(MakeString("failed substitution: ~S"),1,MakeConstantList(ClEnv.Exception_I.Id()))
              Result = F_Optimize_c_substitution_any(g0155.Args.Id(),lx,val,CFALSE)
              if !ErrorIn(Result) {
              Result = EID{g0155.Id(),0}
              }
              } 
            } 
          } else {
          Result = F_Optimize_c_substitution_any(g0155.Args.Id(),lx,val,CFALSE)
          if !ErrorIn(Result) {
          Result = EID{g0155.Id(),0}
          }
          } 
        } 
      }  else if (self.Isa.IsIn(C_Instruction) == CTRUE) { 
      { var g0156 *ClaireInstruction   = To_Instruction(self)
        { 
          var s *ClaireSlot  
          _ = s
          var s_iter *ClaireAny  
          Result= EID{CFALSE.Id(),0}
          var s_support *ClaireList  
          s_support = g0156.Id().Isa.Slots
          for _,s_iter = range(s_support.ValuesO()){ 
            s = ToSlot(s_iter)
            var loop_11 EID 
            _ = loop_11
            { var y *ClaireAny   = Core.F_get_slot(s,ToObject(g0156.Id()))
              { var arg_12 *ClaireAny  
                var try_13 EID 
                try_13 = F_Optimize_c_substitution_any(y,lx,val,eval_ask)
                if ErrorIn(try_13) {loop_11 = try_13
                } else {
                arg_12 = ANY(try_13)
                loop_11 = Core.F_put_slot(s,ToObject(g0156.Id()),arg_12).ToEID()
                }
                } 
              } 
            if ErrorIn(loop_11) {Result = loop_11
            break
            } else {
            }
            } 
          } 
        if !ErrorIn(Result) {
        Result = EID{g0156.Id(),0}
        }
        } 
      } else {
      Result = self.ToEID()
      } 
    return Result} 
  
// The EID go function for: c_substitution @ any (throw: true) 
func E_Optimize_c_substitution_any (self EID,lx EID,val EID,eval_ask EID) EID { 
    return F_Optimize_c_substitution_any(ANY(self),
      ToList(OBJ(lx)),
      ToList(OBJ(val)),
      ToBoolean(OBJ(eval_ask)) )} 
  
// needed
/* The go function for: eval(x:any,y:class) [status=1] */
func F_eval_any2 (x *ClaireAny ,y *ClaireClass ) EID { 
    // eid body s = any
    var Result EID 
    Result = EVAL(x)
    return Result} 
  
// The EID go function for: eval @ list<type_expression>(any, class) (throw: true) 
func E_eval_any2 (x EID,y EID) EID { 
    return F_eval_any2(ANY(x),ToClass(OBJ(y)) )} 
  
// returns the list of bound variables in a piece of code
/* The go function for: bound_variables(self:any) [status=1] */
func F_Optimize_bound_variables_any (self *ClaireAny ) EID { 
    // eid body s = list
    var Result EID 
    { var l *ClaireList   = ToType(C_any.Id()).EmptyList()
      if (self.Isa.IsIn(Language.C_Instruction_with_var) == CTRUE) { 
        { var g0158 *Language.InstructionWithVar   = Language.To_InstructionWithVar(self)
          l = MakeList(ToType(C_any.Id()),g0158.ClaireVar.Id())
          } 
        } 
      if (self.Isa.IsIn(C_Variable) == CTRUE) { 
        Result = EID{CNIL.Id(),0}
        }  else if (self.Isa.IsIn(C_Instruction) == CTRUE) { 
        { var g0160 *ClaireInstruction   = To_Instruction(self)
          { 
            var s *ClaireSlot  
            _ = s
            var s_iter *ClaireAny  
            Result= EID{CFALSE.Id(),0}
            for _,s_iter = range(g0160.Isa.Slots.ValuesO()){ 
              s = ToSlot(s_iter)
              var loop_1 EID 
              _ = loop_1
              var try_2 EID 
              { var arg_3 *ClaireList  
                var try_4 EID 
                try_4 = F_Optimize_bound_variables_any(Core.F_get_slot(s,ToObject(g0160.Id())))
                if ErrorIn(try_4) {try_2 = try_4
                } else {
                arg_3 = ToList(OBJ(try_4))
                try_2 = l.Add_star(arg_3)
                }
                } 
              if ErrorIn(try_2) {Result = try_2
              break
              } else {
              l = ToList(OBJ(try_2))
              loop_1 = EID{l.Id(),0}
              }
              } 
            } 
          } 
        }  else if (self.Isa.IsIn(C_list) == CTRUE) { 
        { var g0161 *ClaireList   = ToList(self)
          { 
            var x *ClaireAny  
            _ = x
            Result= EID{CFALSE.Id(),0}
            var x_support *ClaireList  
            x_support = g0161
            x_len := x_support.Length()
            for i_it := 0; i_it < x_len; i_it++ { 
              x = x_support.At(i_it)
              var loop_5 EID 
              _ = loop_5
              var try_6 EID 
              { var arg_7 *ClaireList  
                var try_8 EID 
                try_8 = F_Optimize_bound_variables_any(x)
                if ErrorIn(try_8) {try_6 = try_8
                } else {
                arg_7 = ToList(OBJ(try_8))
                try_6 = l.Add_star(arg_7)
                }
                } 
              if ErrorIn(try_6) {Result = try_6
              break
              } else {
              l = ToList(OBJ(try_6))
              loop_5 = EID{l.Id(),0}
              }
              } 
            } 
          } 
        } else {
        Result = EID{CFALSE.Id(),0}
        } 
      if !ErrorIn(Result) {
      Result = EID{l.Id(),0}
      }
      } 
    return Result} 
  
// The EID go function for: bound_variables @ any (throw: true) 
func E_Optimize_bound_variables_any (self EID) EID { 
    return F_Optimize_bound_variables_any(ANY(self) )} 
  
// we must recognize true boolean ! coercion
/* The go function for: c_boolean(x:any) [status=1] */
func F_Optimize_c_boolean_any (x *ClaireAny ) EID { 
    // eid body s = any
    var Result EID 
    { var tx *ClaireType  
      var try_1 EID 
      try_1 = Core.F_CALL(C_c_type,ARGS(x.ToEID()))
      if ErrorIn(try_1) {Result = try_1
      } else {
      tx = ToType(OBJ(try_1))
      { var ptx *ClaireType   = F_Optimize_ptype_type(tx)
        if (ptx.Included(ToType(C_boolean.Id())) == CTRUE) { 
          if (x.Isa.IsIn(Language.C_Call) == CTRUE) { 
            { var g0162 *Language.Call   = Language.To_Call(x)
              var g0163I *ClaireBoolean  
              var try_2 EID 
              { 
                var v_and7 *ClaireBoolean  
                
                v_and7 = Equal(g0162.Selector.Id(),Core.C_not.Id())
                if (v_and7 == CFALSE) {try_2 = EID{CFALSE.Id(),0}
                } else { 
                  var try_3 EID 
                  { var arg_4 *ClaireType  
                    var try_5 EID 
                    { var arg_6 *ClaireType  
                      var try_7 EID 
                      try_7 = Core.F_CALL(C_c_type,ARGS(g0162.Args.At(1-1).ToEID()))
                      if ErrorIn(try_7) {try_5 = try_7
                      } else {
                      arg_6 = ToType(OBJ(try_7))
                      try_5 = EID{F_Optimize_ptype_type(arg_6).Id(),0}
                      }
                      } 
                    if ErrorIn(try_5) {try_3 = try_5
                    } else {
                    arg_4 = ToType(OBJ(try_5))
                    try_3 = EID{Core.F__I_equal_any(arg_4.Id(),C_boolean.Id()).Id(),0}
                    }
                    } 
                  if ErrorIn(try_3) {try_2 = try_3
                  } else {
                  v_and7 = ToBoolean(OBJ(try_3))
                  if (v_and7 == CFALSE) {try_2 = EID{CFALSE.Id(),0}
                  } else { 
                    try_2 = EID{CTRUE.Id(),0}} 
                  } 
                }
                } 
              if ErrorIn(try_2) {Result = try_2
              } else {
              g0163I = ToBoolean(OBJ(try_2))
              if (g0163I == CTRUE) { 
                { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                  _CL_obj.Selector = ToProperty(Core.C__I_equal.Id())
                  { 
                    var va_arg1 *Language.Call  
                    var va_arg2 *ClaireList  
                    va_arg1 = _CL_obj
                    { 
                      var v_bag_arg *ClaireAny  
                      va_arg2= ToType(CEMPTY.Id()).EmptyList()
                      { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                        _CL_obj.Selector = C_boolean_I
                        _CL_obj.Args = MakeConstantList(g0162.Args.At(1-1))
                        v_bag_arg = _CL_obj.Id()
                        } 
                      va_arg2.AddFast(v_bag_arg)
                      va_arg2.AddFast(CTRUE.Id())} 
                    va_arg1.Args = va_arg2
                    } 
                  g0162 = _CL_obj
                  } 
                Result = EID{g0162.Id(),0}
                } else {
                Result = EID{CFALSE.Id(),0}
                } 
              }
              if !ErrorIn(Result) {
              x = g0162.Id()
              Result = x.ToEID()
              }
              } 
            } else {
            Result = EID{CFALSE.Id(),0}
            } 
          if !ErrorIn(Result) {
          if (tx.Included(ToType(C_boolean.Id())) == CTRUE) { 
            Result = F_Compile_c_strict_code_any(x,C_boolean)
            } else {
            Result = Core.F_CALL(C_c_code,ARGS(x.ToEID(),EID{C_boolean.Id(),0}))
            } 
          }
          }  else if (tx.Included(ToType(C_list.Id())) == CTRUE) { 
          { var arg_8 *Language.Call  
            { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
              _CL_obj.Selector = ToProperty(Core.C__I_equal.Id())
              { 
                var va_arg1 *Language.Call  
                var va_arg2 *ClaireList  
                va_arg1 = _CL_obj
                { 
                  var v_bag_arg *ClaireAny  
                  va_arg2= ToType(CEMPTY.Id()).EmptyList()
                  { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                    _CL_obj.Selector = C_length
                    _CL_obj.Args = MakeConstantList(x)
                    v_bag_arg = _CL_obj.Id()
                    } 
                  va_arg2.AddFast(v_bag_arg)
                  va_arg2.AddFast(MakeInteger(0).Id())} 
                va_arg1.Args = va_arg2
                } 
              arg_8 = _CL_obj
              } 
            Result = Core.F_CALL(C_c_code,ARGS(EID{arg_8.Id(),0}))
            } 
          } else {
          { var arg_9 *Language.Call  
            { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
              _CL_obj.Selector = C_boolean_I
              _CL_obj.Args = MakeConstantList(x)
              arg_9 = _CL_obj
              } 
            Result = Core.F_CALL(C_c_code,ARGS(EID{arg_9.Id(),0}))
            } 
          } 
        } 
      }
      } 
    return Result} 
  
// The EID go function for: c_boolean @ any (throw: true) 
func E_Optimize_c_boolean_any (x EID) EID { 
    return F_Optimize_c_boolean_any(ANY(x) )} 
  