/***** CLAIRE Compilation of file /Users/ycaseau/Dropbox/src/clairev4.0/src/compile/gostat.cl 
         [version 4.0.02 / safety 5] Friday 12-24-2021 *****/

package Generate
import (_ "fmt"
	. "Kernel"
	"Core"
	"Language"
	"Reader"
	"Optimize"
)

//-------- dumb function to prevent import errors --------
func import_g0136() { 
    _ = Core.It
    _ = Language.It
    _ = Reader.It
    _ = Optimize.It
    } 
  
  
//+-------------------------------------------------------------+
//| CLAIRE                                                      |
//| gostat.cl                                                   |
//| Copyright (C) 2020-2021   Yves Caseau. All Rights Reserved  |
//| cf. copyright info in file object.cl: about()               |
//+-------------------------------------------------------------+
// statement is implemented as a general method that calls a restriction
//        g_statement(self:any,e:class,v:string,err:boolean,loop:any)
// (1) e is the goType that the variable v must receive (HENCE goCast must be inserted)
//     a proper goType is a class, or EID, or void
// (2) The argument v is the named of the C variable in which the
//     result of the evaluation must be placed.
// (3) err tells if an error is possible, which forces to create a chain an not a block (see Do for example)
//     Note : if err = true, s is expected to be EID to (a) force a chain (b) place the error value in v
// (4) loop is either false (not within a loop) or a tuple(v,s) inside the compiling of While/For
//     This tuple describes the vreturn Variable in case a break(v) is encoutered
// there are two possible outputs: blocks (lines of code without {}, used to be call inner_statement)
// and chains  (we use chains to denote long nested ifs that manage error handling)
// indentation : 
//    we call statement(s) at the proper current indentation level => it produices n lines with the indentation
//    and stop after a break line, at the proper identation level
//**********************************************************************
//*  Table of contents:                                                *
//*          Part 1: Unfolding of complex expressions                  *
//*          Part 2: Error Management and EID Unfolding                *
//*          Part 3: Basic control structures                          *
//*          Part 4: iteration                                         *
//*          Part 5: CLAIRE-specific structures                        *
//**********************************************************************
//*************************************************************************
//*          Part 1: Unfolding of complex expressions                     *
//*************************************************************************
// when local CLAIRE expressions are not go expression, we need to unfold the global expression into a big Let
// HOWEVER, if only works for list of arguments whose evaluation order is not specified ! (because we move some of the evaluations earlier)
// this reintrant compiling (calling g_statement on a expanded Let) works because Let checks if g_expression can be used
// the same pattern is used for call_slot/call_table
// this function is used to unfold complex expressions that should be compiled as
// expressions and not statements. It takes a list of arguments l and returns the
// embedded Lets that defines the necessary variable or nil (nothing is needed)
// this list is of the form  (a1 .. am) where
//     m is the number of statement args in l
//     ai is a Let that defines the i-th variable corresponding to the i-th bad guy
// CLAIRE 4: we unfold args that are not functional or args that can throw error
/* {1} The go function for: unfold_args(l:list) [status=1] */
func F_Generate_unfold_args_list (l *ClaireList ) EID { 
    var Result EID 
    { var lbad *ClaireList  
      _ = lbad
      var try_1 EID 
      /*g_try(v2:"try_1",loop:false) */
      { var i_out *ClaireList   = ToType(CEMPTY.Id()).EmptyList()
        /*g_try(v2:"try_1",loop:false) */
        { var i int  = 1
          { var g0137 int  = l.Length()
            _ = g0137
            try_1= EID{CFALSE.Id(),0}
            for (i <= g0137) { 
              /* While stat, v:"try_1" loop:false */
              var loop_2 EID 
              _ = loop_2
              { 
              /*g_try(v2:"loop_2",loop:tuple("try_1", EID)) */
              var g0138I *ClaireBoolean  
              var try_3 EID 
              /*g_try(v2:"try_3",loop:false) */
              { var arg_4 *ClaireBoolean  
                _ = arg_4
                var try_5 EID 
                /*g_try(v2:"try_5",loop:false) */
                try_5 = F_Generate_g_clean_any(l.At(i-1))
                /* ERROR PROTECTION INSERTED (arg_4-try_3) */
                if ErrorIn(try_5) {try_3 = try_5
                } else {
                arg_4 = ToBoolean(OBJ(try_5))
                try_3 = EID{arg_4.Not.Id(),0}
                }
                } 
              /* ERROR PROTECTION INSERTED (g0138I-loop_2) */
              if ErrorIn(try_3) {loop_2 = try_3
              } else {
              g0138I = ToBoolean(OBJ(try_3))
              if (g0138I == CTRUE) { 
                loop_2 = EID{i_out.AddFast(MakeInteger(i).Id()).Id(),0}/*t=any,s=EID*/
                } else {
                loop_2 = EID{CFALSE.Id(),0}
                } 
              }
              /* ERROR PROTECTION INSERTED (loop_2-loop_2) */
              if ErrorIn(loop_2) {try_1 = loop_2
              break
              } else {
              i = (i+1)
              }
              /* try?:false, v2:"v_while6" loop will be:tuple("try_1", EID) */
              } 
            }
            } 
          } 
        /* ERROR PROTECTION INSERTED (try_1-try_1) */
        if !ErrorIn(try_1) {
        try_1 = EID{i_out.Id(),0}
        }
        } 
      /* ERROR PROTECTION INSERTED (lbad-Result) */
      if ErrorIn(try_1) {Result = try_1
      } else {
      lbad = ToList(OBJ(try_1))
      
      { 
        var v_list3 *ClaireList  
        var i int 
        var v_local3 *ClaireAny  
        v_list3 = lbad
        Result = EID{CreateList(ToType(CEMPTY.Id()),v_list3.Length()).Id(),0}
        for CLcount := 0; CLcount < v_list3.Length(); CLcount++{ 
          i = ToInteger(v_list3.At(CLcount)).Value
          var try_6 EID 
          /*g_try(v2:"try_6",loop:tuple("Result", EID)) */
          { var _CL_obj *Language.Let   = Language.To_Let(new(Language.Let).Is(Language.C_Let))
            /*g_try(v2:"try_6",loop:tuple("Result", EID)) */
            { 
              var va_arg1 *Language.InstructionWithVar  
              var va_arg2 *ClaireVariable  
              va_arg1 = Language.To_InstructionWithVar(_CL_obj.Id())
              var try_7 EID 
              /*g_try(v2:"try_7",loop:false) */
              { var arg_8 *ClaireClass  
                _ = arg_8
                var try_9 EID 
                /*g_try(v2:"try_9",loop:false) */
                try_9 = Language.F_static_type_any(l.At(i-1))
                /* ERROR PROTECTION INSERTED (arg_8-try_7) */
                if ErrorIn(try_9) {try_7 = try_9
                } else {
                arg_8 = ToClass(OBJ(try_9))
                try_7 = EID{F_Generate_build_Variable_string(F_Generate_genvar_string(MakeString("arg_")),arg_8.Id()).Id(),0}
                }
                } 
              /* ERROR PROTECTION INSERTED (va_arg2-try_6) */
              if ErrorIn(try_7) {try_6 = try_7
              } else {
              va_arg2 = To_Variable(OBJ(try_7))
              va_arg1.ClaireVar = va_arg2
              /*Variable->Variable*/try_6 = EID{va_arg2.Id(),0}
              }
              } 
            /* ERROR PROTECTION INSERTED (try_6-try_6) */
            if ErrorIn(try_6) {Result = try_6
            break
            } else {
            _CL_obj.Value = l.At(i-1)
            /*any->any*/try_6 = EID{_CL_obj.Id(),0}
            }
            } 
          /* ERROR PROTECTION INSERTED (v_local3-Result) */
          if ErrorIn(try_6) {Result = try_6
          break
          } else {
          v_local3 = ANY(try_6)
          ToList(OBJ(Result)).PutAt(CLcount,v_local3)
          } 
        }
        } 
      }
      } 
    return Result} 
  
// The EID go function for: unfold_args @ list (throw: true) 
func E_Generate_unfold_args_list (l EID) EID { 
    return F_Generate_unfold_args_list(ToList(OBJ(l)) )} 
  
// uses the previous list to use the variable instead of the Fold.
// l is the list of arguments, ld is the previously build unfold_args(l)
/* {1} The go function for: unfold_arg(l:list,ld:list,x:any) [status=1] */
func F_Generate_unfold_arg_list (l *ClaireList ,ld *ClaireList ,x *ClaireAny ) EID { 
    var Result EID 
    { var i int  = 1
      { var j int  = 0
        _ = j
        { var m int  = l.Length()
          _ = m
          var g0139I *ClaireBoolean  
          var try_1 EID 
          /*g_try(v2:"try_1",loop:false) */
          try_1= EID{CFALSE.Id(),0}
          for (i <= m) { 
            /* While stat, v:"try_1" loop:false */
            var loop_2 EID 
            _ = loop_2
            { 
            /*g_try(v2:"loop_2",loop:tuple("try_1", EID)) */
            var g0140I *ClaireBoolean  
            var try_3 EID 
            /*g_try(v2:"try_3",loop:false) */
            { var arg_4 *ClaireBoolean  
              _ = arg_4
              var try_5 EID 
              /*g_try(v2:"try_5",loop:false) */
              try_5 = F_Generate_g_clean_any(l.At(i-1))
              /* ERROR PROTECTION INSERTED (arg_4-try_3) */
              if ErrorIn(try_5) {try_3 = try_5
              } else {
              arg_4 = ToBoolean(OBJ(try_5))
              try_3 = EID{arg_4.Not.Id(),0}
              }
              } 
            /* ERROR PROTECTION INSERTED (g0140I-loop_2) */
            if ErrorIn(try_3) {loop_2 = try_3
            } else {
            g0140I = ToBoolean(OBJ(try_3))
            if (g0140I == CTRUE) { 
              j = (j+1)
              if (Equal(l.At(i-1),x) == CTRUE) { 
                try_1 = EID{CTRUE.Id(),0}
                break
                } else {
                loop_2 = EID{CFALSE.Id(),0}
                } 
              }  else if (Equal(l.At(i-1),x) == CTRUE) { 
              try_1 = EID{CFALSE.Id(),0}
              break
              } else {
              loop_2 = EID{CFALSE.Id(),0}
              } 
            }
            /* ERROR PROTECTION INSERTED (loop_2-loop_2) */
            if ErrorIn(loop_2) {try_1 = loop_2
            break
            } else {
            i = (i+1)
            }
            /* try?:false, v2:"v_while5" loop will be:tuple("try_1", EID) */
            } 
          }
          /* ERROR PROTECTION INSERTED (g0139I-Result) */
          if ErrorIn(try_1) {Result = try_1
          } else {
          g0139I = ToBoolean(OBJ(try_1))
          if (g0139I == CTRUE) { 
            Result = Core.F_CALL(Language.C_var,ARGS(ld.At(j-1).ToEID()))
            } else {
            Result = l.At(i-1).ToEID()
            } 
          }
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: unfold_arg @ list (throw: true) 
func E_Generate_unfold_arg_list (l EID,ld EID,x EID) EID { 
    return F_Generate_unfold_arg_list(ToList(OBJ(l)),ToList(OBJ(ld)),ANY(x) )} 
  
// creates the Let from the ldef definition and places the statement x in the body
// note that the error handling is done in the Let (with g_statement)
// x is the call form where the variable has been replaced if needed
/* {1} The go function for: unfold_use(ldef:list,x:any,s:class,v:string,err:boolean,loop:any) [status=1] */
func F_Generate_unfold_use_list (ldef *ClaireList ,x *ClaireAny ,s *ClaireClass ,v *ClaireString ,err *ClaireBoolean ,loop *ClaireAny ) EID { 
    var Result EID 
    if (F_Generate_eid_require_ask_any(x) == CTRUE) { 
      Result = F_Generate_unfold_eid_list(ldef,
        x,
        s,
        (v).Id(),
        err,
        loop)
      } else {
      { var n int  = ldef.Length()
        { var vb int  = ClEnv.Verbose
          _ = vb
          ClEnv.Verbose = 0
          /*integer->integer*//*g_try(v2:"Result",loop:true) */
          if (F_boolean_I_any(ldef.Id()).Id() != CTRUE.Id()) { 
            Result = ToException(Core.C_general_error.Make(MakeString("[internal] design bug g_func(~S) should be true").Id(),MakeConstantList(x).Id())).Close()
            } else {
            Result = EID{CFALSE.Id(),0}
            } 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          /*g_try(v2:"Result",loop:true) */
          { var i int  = 1
            { var g0141 int  = (n-1)
              _ = g0141
              Result= EID{CFALSE.Id(),0}
              for (i <= g0141) { 
                /* While stat, v:"Result" loop:true */
                var loop_1 EID 
                _ = loop_1
                { 
                /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
                loop_1 = Core.F_write_property(C_arg,ToObject(ldef.At(i-1)),ldef.At((i+1)-1))
                /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
                if ErrorIn(loop_1) {Result = loop_1
                break
                } else {
                i = (i+1)
                }
                /* try?:false, v2:"v_while7" loop will be:tuple("Result", EID) */
                } 
              }
              } 
            } 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          Language.To_Let(ldef.At(n-1)).Arg = x
          /*any->any*/ClEnv.Verbose = vb
          /*integer->integer*/
          Result = F_Generate_g_statement_Let(Language.To_Let(ldef.At(1-1)),
            s,
            v,
            err,
            loop)
          }}
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: unfold_use @ list (throw: true) 
func E_Generate_unfold_use_list (ldef EID,x EID,s EID,v EID,err EID,loop EID) EID { 
    return F_Generate_unfold_use_list(ToList(OBJ(ldef)),
      ANY(x),
      ToClass(OBJ(s)),
      ToString(OBJ(v)),
      ToBoolean(OBJ(err)),
      ANY(loop) )} 
  
//*************************************************************************
//*          Part 2: Error Management and EID Unfolding                   *
//*************************************************************************
// this is the error catching pattern: evaluate(self) and check if error then place it in vglobal,
// if no error we want the value in v with expected gotype e (a true gotype = class)
// if v is an EID variable, do not create an extra variable (we use it temporarily)
// in a loop we generate a break to exit to loop
// v is the variable that must receive self
// note : g_try produces a pattern   <e = code>, if Err(e) {res =e} else { ...
// that must be closed } with a close_try => and nothing after the close_try (nothing must if an error occured)
/* {1} The go function for: g_try(self:any,v:string,e:class,vglobal:string,loop:any) [status=1] */
func F_Generate_g_try_any (self *ClaireAny ,v *ClaireString ,e *ClaireClass ,vglobal *ClaireString ,loop *ClaireAny ) EID { 
    var Result EID 
    { var v2 *ClaireString  
      if (e.Id() == Optimize.C_EID.Id()) { 
        v2 = v
        } else {
        v2 = F_Generate_genvar_string(MakeString("try_"))
        } 
      if (e.Id() != Optimize.C_EID.Id()) { 
        F_Generate_var_declaration_string(v2,Optimize.C_EID,1)
        } 
      /*g_try(v2:"Result",loop:true) */
      PRINC("/*g_try(v2:")
      /*g_try(v2:"Result",loop:true) */
      Result = Core.F_print_any((v2).Id())
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(",loop:")
      /*g_try(v2:"Result",loop:true) */
      Result = Core.F_CALL(C_print,ARGS(loop.ToEID()))
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(") */")
      F_Generate_breakline_void()
      PRINC("")
      Result = EVOID
      }}
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      /*g_try(v2:"Result",loop:true) */
      Result = Core.F_CALL(C_Generate_g_statement,ARGS(self.ToEID(),
        EID{Optimize.C_EID.Id(),0},
        EID{(v2).Id(),0},
        EID{CTRUE.Id(),0},
        loop.ToEID()))
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      /*g_try(v2:"Result",loop:true) */
      if ((self.Isa.IsIn(Language.C_Do) == CTRUE) && 
          (loop.Isa.IsIn(Language.C_Tuple) == CTRUE)) { 
        PRINC("/*A*/{")
        F_Generate_breakline_void()
        PRINC("")
        Result = EVOID
        } else {
        PRINC("/* ERROR PROTECTION INSERTED (")
        F_princ_string(v)
        PRINC("-")
        F_princ_string(vglobal)
        PRINC(") */")
        F_Generate_breakline_void()
        PRINC("")
        if ((v.Value == vglobal.Value) && 
            ((e.Id() == Optimize.C_EID.Id()) && 
              (C_tuple.Id() != loop.Isa.Id()))) { 
          PRINC("if !ErrorIn(")
          F_c_princ_string(v2)
          PRINC(") {")
          F_Generate_breakline_void()
          PRINC("")
          Result = EVOID
          } else {
          PRINC("if ErrorIn(")
          F_c_princ_string(v2)
          PRINC(") {")
          if (v.Value != vglobal.Value) { 
            F_c_princ_string(vglobal)
            PRINC(" = ")
            F_c_princ_string(v2)
            F_Generate_breakline_void()
            PRINC("")
            } 
          /*g_try(v2:"Result",loop:true) */
          if (C_tuple.Id() == loop.Isa.Id()) { 
            /*g_try(v2:"Result",loop:true) */
            var g0142I *ClaireBoolean  
            var try_1 EID 
            /*g_try(v2:"try_1",loop:false) */
            { var arg_2 *ClaireAny  
              _ = arg_2
              var try_3 EID 
              /*g_try(v2:"try_3",loop:false) */
              try_3 = Core.F_CALL(C_nth,ARGS(loop.ToEID(),EID{C__INT,IVAL(1)}))
              /* ERROR PROTECTION INSERTED (arg_2-try_1) */
              if ErrorIn(try_3) {try_1 = try_3
              } else {
              arg_2 = ANY(try_3)
              try_1 = EID{Core.F__I_equal_any(arg_2,(vglobal).Id()).Id(),0}
              }
              } 
            /* ERROR PROTECTION INSERTED (g0142I-Result) */
            if ErrorIn(try_1) {Result = try_1
            } else {
            g0142I = ToBoolean(OBJ(try_1))
            if (g0142I == CTRUE) { 
              /*g_try(v2:"Result",loop:true) */
              { var arg_4 *ClaireAny  
                _ = arg_4
                var try_5 EID 
                /*g_try(v2:"try_5",loop:false) */
                try_5 = Core.F_CALL(C_nth,ARGS(loop.ToEID(),EID{C__INT,IVAL(1)}))
                /* ERROR PROTECTION INSERTED (arg_4-Result) */
                if ErrorIn(try_5) {Result = try_5
                } else {
                arg_4 = ANY(try_5)
                Result = Core.F_CALL(C_c_princ,ARGS(arg_4.ToEID()))
                }
                } 
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              PRINC(" = ")
              F_c_princ_string(v2)
              F_Generate_breakline_void()
              PRINC("")
              Result = EVOID
              }
              } else {
              Result = EID{CFALSE.Id(),0}
              } 
            }
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC("break")
            Result = F_Generate_breakline_void().ToEID()
            }
            } else {
            Result = EID{CFALSE.Id(),0}
            } 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC("} else {")
          F_Generate_breakline_void()
          PRINC("")
          Result = EVOID
          }
          } 
        } 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      if ((e.Id() != C_void.Id()) && 
          (v.Value != v2.Value)) { 
        F_c_princ_string(v)
        PRINC(" = ")
        /*g_try(v2:"Result",loop:true) */
        Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).FromEid(v2,e)
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        F_Generate_breakline_void()
        PRINC("")
        Result = EVOID
        }
        } else {
        Result = EID{CFALSE.Id(),0}
        } 
      }}}
      } 
    return Result} 
  
// The EID go function for: g_try @ any (throw: true) 
func E_Generate_g_try_any (self EID,v EID,e EID,vglobal EID,loop EID) EID { 
    return F_Generate_g_try_any(ANY(self),
      ToString(OBJ(v)),
      ToClass(OBJ(e)),
      ToString(OBJ(vglobal)),
      ANY(loop) )} 
  
// when the error is nested in the expression, the unfold pattern will make sure that we separate the sub_exp that may
// create the error, but assignment is not managed this way, hence this code to avoid double error check
/* {1} The go function for: g_try(self:Assign,v:string,e:class,vglobal:string,loop:any) [status=1] */
func F_Generate_g_try_Assign (self *Language.Assign ,v *ClaireString ,e *ClaireClass ,vglobal *ClaireString ,loop *ClaireAny ) EID { 
    var Result EID 
    { var _Zvar *ClaireAny   = self.ClaireVar
      { var v1 *ClaireString  
        var try_1 EID 
        /*g_try(v2:"try_1",loop:false) */
        try_1 = Core.F_CALL(C_Generate_c_string,ARGS(EID{Optimize.C_PRODUCER.Value,0},_Zvar.ToEID()))
        /* ERROR PROTECTION INSERTED (v1-Result) */
        if ErrorIn(try_1) {Result = try_1
        } else {
        v1 = ToString(OBJ(try_1))
        { var _Zrange *ClaireClass   = ToTypeExpression(OBJ(Core.F_CALL(C_range,ARGS(_Zvar.ToEID())))).Class_I()
          /*g_try(v2:"Result",loop:true) */
          Result = Core.F_CALL(C_Generate_g_try,ARGS(self.Arg.ToEID(),
            EID{(v1).Id(),0},
            EID{_Zrange.Id(),0},
            EID{(vglobal).Id(),0},
            loop.ToEID()))
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          if (e.Id() != C_void.Id()) { 
            F_c_princ_string(v)
            PRINC(" = ")
            /*g_try(v2:"Result",loop:true) */
            Result = F_Generate_use_variable_string(v1,e,_Zrange)
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            F_Generate_breakline_void()
            PRINC("")
            Result = EVOID
            }
            } else {
            Result = EID{CFALSE.Id(),0}
            } 
          }
          } 
        }
        } 
      } 
    return Result} 
  
// The EID go function for: g_try @ Assign (throw: true) 
func E_Generate_g_try_Assign (self EID,v EID,e EID,vglobal EID,loop EID) EID { 
    return F_Generate_g_try_Assign(Language.To_Assign(OBJ(self)),
      ToString(OBJ(v)),
      ToClass(OBJ(e)),
      ToString(OBJ(vglobal)),
      ANY(loop) )} 
  
// each g_try produces a {, which we must balance before returning a new line 
// does NOT change OPT.level !
/* {1} The go function for: close_try(n:integer) [status=0] */
func F_Generate_close_try_integer (n int)  { 
    // procedure body with s = void 
{ var i int  = 1
      _ = i
      { var g0143 int  = n
        _ = g0143
        for (i <= g0143) { 
          /* While stat, v:"Unused" loop:false */
          PRINC("}")
          i = (i+1)
          /* try?:false, v2:"v_while4" loop will be:tuple("Unused", void) */
          } 
        } 
      } 
    if (n > 0) { 
      F_Generate_breakline_void()
      } 
    } 
  
// The EID go function for: close_try @ integer (throw: false) 
func E_Generate_close_try_integer (n EID) EID { 
    F_Generate_close_try_integer(INT(n) )
    return EVOID} 
  
// special case when v is a g_func that can produce an error (s is assumed to be EID)
/* {1} The go function for: error_wrap(self:any,s:class,v:string) [status=1] */
func F_Generate_error_wrap_any (self *ClaireAny ,s *ClaireClass ,v *ClaireString ) EID { 
    var Result EID 
    /*g_try(v2:"Result",loop:true) */
    if (s.Id() != Optimize.C_EID.Id()) { 
      /*g_try(v2:"Result",loop:true) */
      { var arg_1 *ClaireList  
        _ = arg_1
        var try_2 EID 
        /*g_try(v2:"try_2",loop:false) */
        { 
          var v_bag_arg *ClaireAny  
          try_2= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
          ToList(OBJ(try_2)).AddFast(self)
          var try_3 EID 
          /*g_try(v2:"try_3",loop:false) */
          try_3 = Optimize.F_Compile_g_throw_any(self)
          /* ERROR PROTECTION INSERTED (v_bag_arg-try_2) */
          if ErrorIn(try_3) {try_2 = try_3
          } else {
          v_bag_arg = ANY(try_3)
          ToList(OBJ(try_2)).AddFast(v_bag_arg)}
          } 
        /* ERROR PROTECTION INSERTED (arg_1-Result) */
        if ErrorIn(try_2) {Result = try_2
        } else {
        arg_1 = ToList(OBJ(try_2))
        Result = Core.F_tformat_string(MakeString("---- g_throw(~S) = ~S\n"),0,arg_1)
        }
        } 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = ToException(Core.C_general_error.Make(MakeString("design bug for error_wrap with ~S and s = ~S").Id(),MakeConstantList(self,s.Id()).Id())).Close()
      }
      } else {
      Result = EID{CFALSE.Id(),0}
      } 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    F_princ_string(v)
    PRINC(" = ")
    /*g_try(v2:"Result",loop:true) */
    Result = Core.F_CALL(C_Generate_g_expression,ARGS(self.ToEID(),EID{Optimize.C_EID.Id(),0}))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    F_Generate_breakline_void()
    PRINC("")
    Result = EVOID
    }
    }
    return Result} 
  
// The EID go function for: error_wrap @ any (throw: true) 
func E_Generate_error_wrap_any (self EID,s EID,v EID) EID { 
    return F_Generate_error_wrap_any(ANY(self),ToClass(OBJ(s)),ToString(OBJ(v)) )} 
  
// this is a special case when the statement result is not needed (e = void) so we should
// not reuse v as a temporary variable (which we considered) 
// this is called inside a For/While, so loop is a tuple
/* {1} The go function for: g_try_void(self:any,vglobal:string,loop:any) [status=1] */
func F_Generate_g_try_void_any (self *ClaireAny ,vglobal *ClaireString ,loop *ClaireAny ) EID { 
    var Result EID 
    var g0147I *ClaireBoolean  
    if (self.Isa.IsIn(Language.C_Assign) == CTRUE) { 
      { var g0144 *Language.Assign   = Language.To_Assign(self)
        _ = g0144
        g0147I = Equal(ANY(Core.F_CALL(C_range,ARGS(g0144.ClaireVar.ToEID()))),Optimize.C_EID.Id())
        } 
      } else {
      g0147I = CFALSE
      } 
    if (g0147I == CTRUE) { 
      { var _Zvar *ClaireAny   = ANY(Core.F_CALL(Language.C_var,ARGS(self.ToEID())))
        _ = _Zvar
        { var v1 *ClaireString  
          _ = v1
          var try_1 EID 
          /*g_try(v2:"try_1",loop:false) */
          try_1 = Core.F_CALL(C_Generate_c_string,ARGS(EID{Optimize.C_PRODUCER.Value,0},_Zvar.ToEID()))
          /* ERROR PROTECTION INSERTED (v1-Result) */
          if ErrorIn(try_1) {Result = try_1
          } else {
          v1 = ToString(OBJ(try_1))
          Result = Core.F_CALL(C_Generate_g_try,ARGS(Core.F_CALL(C_arg,ARGS(self.ToEID())),
            EID{(v1).Id(),0},
            EID{Optimize.C_EID.Id(),0},
            EID{(vglobal).Id(),0},
            loop.ToEID()))
          }
          } 
        } 
      } else {
      { var v2 *ClaireString   = F_Generate_genvar_string(MakeString("loop_"))
        F_Generate_var_declaration_string(v2,Optimize.C_EID,1)
        PRINC("_ = ")
        F_princ_string(v2)
        F_Generate_breakline_void()
        PRINC("")
        if (self.Isa.IsIn(Language.C_Do) == CTRUE) { 
          { var g0145 *Language.Do   = Language.To_Do(self)
            _ = g0145
            PRINC("{ ")
            F_Generate_breakline_void()
            Result = F_Generate_do_statement_Do(g0145,
              Optimize.C_EID,
              v2,
              CTRUE,
              loop,
              CFALSE)
            } 
          } else {
          Result = Core.F_CALL(C_Generate_g_try,ARGS(self.ToEID(),
            EID{(v2).Id(),0},
            EID{Optimize.C_EID.Id(),0},
            EID{(vglobal).Id(),0},
            loop.ToEID()))
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: g_try_void @ any (throw: true) 
func E_Generate_g_try_void_any (self EID,vglobal EID,loop EID) EID { 
    return F_Generate_g_try_void_any(ANY(self),ToString(OBJ(vglobal)),ANY(loop) )} 
  
// eid_require means that the internal call should better be evaluated in EID mode
// this is really what we need for mClaire/push!(eval(x)) and funcall(f,...) ... but has been be extended to methods
// that do a better job (no allocation) in EID mode
/* {1} The go function for: eid_require?(x:any) [status=0] */
func F_Generate_eid_require_ask_any (x *ClaireAny ) *ClaireBoolean  { 
    // procedure body with s = boolean 
var Result *ClaireBoolean  
    if (x.Isa.IsIn(Language.C_Call) == CTRUE) { 
      { var g0148 *Language.Call   = Language.To_Call(x)
        _ = g0148
        Result = Equal(g0148.Selector.Id(),Core.C_mClaire_push_I.Id())
        } 
      }  else if (x.Isa.IsIn(Language.C_Call_method) == CTRUE) { 
      { var g0149 *Language.CallMethod   = Language.To_CallMethod(x)
        Result = MakeBoolean((g0149.Arg.Selector.Id() == C_funcall.Id()) || (g0149.Arg.Id() == C_Generate__starwrite_value_star.Value) || (g0149.Arg.Id() == C_Generate__starread_property_star.Value) || (g0149.Arg.Selector.Id() == C_write_fast.Id()) || (g0149.Arg.Selector.Id() == Core.C_nth_write.Id()))
        } 
      } else {
      Result = CFALSE
      } 
    return Result} 
  
// The EID go function for: eid_require? @ any (throw: false) 
func E_Generate_eid_require_ask_any (x EID) EID { 
    return EID{F_Generate_eid_require_ask_any(ANY(x) ).Id(),0}} 
  
// eid_provide? says that the call will produce first an EID
/* {1} The go function for: eid_provide?(x:any) [status=0] */
func F_Generate_eid_provide_ask_any (x *ClaireAny ) *ClaireBoolean  { 
    // procedure body with s = boolean 
var Result *ClaireBoolean  
    if (x.Isa.IsIn(Language.C_Call) == CTRUE) { 
      { var g0151 *Language.Call   = Language.To_Call(x)
        _ = g0151
        Result = Equal(g0151.Selector.Id(),Core.C_mClaire_get_stack.Id())
        } 
      }  else if (x.Isa.IsIn(Language.C_Call_method) == CTRUE) { 
      { var g0152 *Language.CallMethod   = Language.To_CallMethod(x)
        Result = MakeBoolean((g0152.Arg.Selector.Id() == Core.C_eval.Id()) || (g0152.Arg.Id() == Optimize.C_Compile_m_unsafe.Value))
        } 
      }  else if (x.Isa.IsIn(C_Variable) == CTRUE) { 
      { var g0153 *ClaireVariable   = To_Variable(x)
        _ = g0153
        Result = g0153.Range.Included(ToType(C_integer.Id()))
        } 
      }  else if (C_integer.Id() == x.Isa.Id()) { 
      Result = CTRUE
      } else {
      Result = CFALSE
      } 
    return Result} 
  
// The EID go function for: eid_provide? @ any (throw: false) 
func E_Generate_eid_provide_ask_any (x EID) EID { 
    return EID{F_Generate_eid_provide_ask_any(ANY(x) ).Id(),0}} 
  
// eid_unfold could use a more general "EID compling mode" (with a list of EID variables passed as context)
// this is a quickfix => we build the EID Let on our own (code borrowed from g_stat@Let)
/* {1} The go function for: unfold_eid(ldef:list,self:any,s:class,v:any,err:boolean,loop:any) [status=1] */
func F_Generate_unfold_eid_list (ldef *ClaireList ,self *ClaireAny ,s *ClaireClass ,v *ClaireAny ,err *ClaireBoolean ,loop *ClaireAny ) EID { 
    var Result EID 
    { var n int  = ldef.Length()
      _ = n
      { var lvar *ClaireList   = ToType(C_Variable.Id()).EmptyList()
        _ = lvar
        { var count_try int  = 0
          _ = count_try
          F_Generate_new_block_string(MakeString("LetEID"))
          /*g_try(v2:"Result",loop:true) */
          { var i int  = 1
            { var g0156 int  = n
              _ = g0156
              Result= EID{CFALSE.Id(),0}
              for (i <= g0156) { 
                /* While stat, v:"Result" loop:true */
                var loop_1 EID 
                _ = loop_1
                { 
                /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
                { var _Zl *Language.Let   = Language.To_Let(ldef.At(i-1))
                  { var v2 *ClaireString  
                    var try_2 EID 
                    /*g_try(v2:"try_2",loop:false) */
                    try_2 = F_Generate_c_string_go_producer1(ToGenerateGoProducer(Optimize.C_PRODUCER.Value),_Zl.ClaireVar)
                    /* ERROR PROTECTION INSERTED (v2-loop_1) */
                    if ErrorIn(try_2) {loop_1 = try_2
                    } else {
                    v2 = ToString(OBJ(try_2))
                    { var x *ClaireAny   = _Zl.Value
                      { var try_ask *ClaireBoolean  
                        _ = try_ask
                        var try_3 EID 
                        /*g_try(v2:"try_3",loop:false) */
                        try_3 = Optimize.F_Compile_g_throw_any(x)
                        /* ERROR PROTECTION INSERTED (try_ask-loop_1) */
                        if ErrorIn(try_3) {loop_1 = try_3
                        } else {
                        try_ask = ToBoolean(OBJ(try_3))
                        F_Generate_var_declaration_string(v2,Optimize.C_EID,0)
                        lvar = lvar.AddFast(_Zl.ClaireVar.Id())/*t=Variable,s=list*/
                        F_Generate_breakline_void()
                        /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
                        if (try_ask == CTRUE) { 
                          count_try = (count_try+1)
                          /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
                          loop_1 = Core.F_CALL(C_Generate_g_try,ARGS(x.ToEID(),
                            EID{(v2).Id(),0},
                            EID{Optimize.C_EID.Id(),0},
                            v.ToEID(),
                            EID{CFALSE.Id(),0}))
                          /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
                          if ErrorIn(loop_1) {Result = loop_1
                          break
                          } else {
                          }
                          } else {
                          loop_1 = F_Generate_statement_any(x,Optimize.C_EID,v2,loop)
                          } 
                        /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
                        if ErrorIn(loop_1) {Result = loop_1
                        break
                        } else {
                        }
                        }
                        } 
                      } 
                    }
                    } 
                  } 
                /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
                if ErrorIn(loop_1) {Result = loop_1
                break
                } else {
                i = (i+1)
                }
                /* try?:false, v2:"v_while7" loop will be:tuple("Result", EID) */
                } 
              }
              } 
            } 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          /*g_try(v2:"Result",loop:true) */
          if (s.Id() != C_void.Id()) { 
            /*g_try(v2:"Result",loop:true) */
            Result = Core.F_CALL(C_princ,ARGS(v.ToEID()))
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC(" = ")
            Result = EVOID
            }
            } else {
            Result = EID{CFALSE.Id(),0}
            } 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          /*g_try(v2:"Result",loop:true) */
          Result = F_Generate_eid_expression_any(self,Optimize.C_EID,lvar)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          F_Generate_close_try_integer(count_try)
          F_Generate_close_block_string(MakeString("LetEID"))
          Result = EVOID
          }}}
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: unfold_eid @ list (throw: true) 
func E_Generate_unfold_eid_list (ldef EID,self EID,s EID,v EID,err EID,loop EID) EID { 
    return F_Generate_unfold_eid_list(ToList(OBJ(ldef)),
      ANY(self),
      ToClass(OBJ(s)),
      ANY(v),
      ToBoolean(OBJ(err)),
      ANY(loop) )} 
  
// eid_expression compiles a call or call_method with one EID variable
// it performs all the compiler optimization (see the eid_fold? pattern in gostat.cl)
/* {1} The go function for: eid_expression(x:any,s:class,lvar:list<Variable>) [status=1] */
func F_Generate_eid_expression_any (x *ClaireAny ,s *ClaireClass ,lvar *ClaireList ) EID { 
    var Result EID 
    if (x.Isa.IsIn(C_Variable) == CTRUE) { 
      { var g0157 *ClaireVariable   = To_Variable(x)
        if ((lvar.Memq(g0157.Id()) == CTRUE) && 
            (s.Id() == Optimize.C_EID.Id())) { 
          { var arg_1 *ClaireString  
            _ = arg_1
            var try_2 EID 
            /*g_try(v2:"try_2",loop:false) */
            try_2 = F_Generate_c_string_go_producer1(ToGenerateGoProducer(Optimize.C_PRODUCER.Value),g0157)
            /* ERROR PROTECTION INSERTED (arg_1-Result) */
            if ErrorIn(try_2) {Result = try_2
            } else {
            arg_1 = ToString(OBJ(try_2))
            F_princ_string(arg_1)
            Result = EVOID
            }
            } 
          } else {
          Result = F_Generate_g_expression_Variable(g0157,s)
          } 
        } 
      }  else if (x.Isa.IsIn(Language.C_Call) == CTRUE) { 
      { var g0158 *Language.Call   = Language.To_Call(x)
        _ = g0158
        PRINC("ClEnv.Push(")
        /*g_try(v2:"Result",loop:true) */
        Result = F_Generate_eid_expression_any(g0158.Args.At(1-1),Optimize.C_EID,lvar)
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC(")")
        Result = EVOID
        }
        } 
      }  else if (x.Isa.IsIn(Language.C_Call_method) == CTRUE) { 
      { var g0159 *Language.CallMethod   = Language.To_CallMethod(x)
        if (g0159.Arg.Selector.Id() == C_funcall.Id()) { 
          PRINC("FASTCALL")
          F_princ_integer((g0159.Args.Length()-1))
          PRINC("(")
          /*g_try(v2:"Result",loop:true) */
          Result = F_Generate_eid_expression_any(g0159.Args.At(1-1),C_method,lvar)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(",")
          /*g_try(v2:"Result",loop:true) */
          Result = F_Generate_eid_expression_any(g0159.Args.At(2-1),Optimize.C_EID,lvar)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          /*g_try(v2:"Result",loop:true) */
          if (g0159.Args.Length() >= 3) { 
            PRINC(",")
            Result = F_Generate_eid_expression_any(g0159.Args.At(3-1),Optimize.C_EID,lvar)
            } else {
            Result = EID{CFALSE.Id(),0}
            } 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          /*g_try(v2:"Result",loop:true) */
          if (g0159.Args.Length() == 4) { 
            PRINC(",")
            Result = F_Generate_eid_expression_any(g0159.Args.At(4-1),Optimize.C_EID,lvar)
            } else {
            Result = EID{CFALSE.Id(),0}
            } 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(")")
          Result = EVOID
          }}}}
          }  else if (g0159.Arg.Id() == C_Generate__starread_property_star.Value) { 
          /*g_try(v2:"Result",loop:true) */
          Result = Core.F_CALL(C_Generate_g_expression,ARGS(g0159.Args.At(1-1).ToEID(),EID{C_property.Id(),0}))
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(".ReadEID(")
          /*g_try(v2:"Result",loop:true) */
          Result = F_Generate_eid_expression_any(g0159.Args.At(2-1),Optimize.C_EID,lvar)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(")")
          Result = EVOID
          }}
          }  else if (g0159.Arg.Selector.Id() == C_write_fast.Id()) { 
          /*g_try(v2:"Result",loop:true) */
          Result = Core.F_CALL(C_Generate_g_expression,ARGS(g0159.Args.At(1-1).ToEID(),EID{C_property.Id(),0}))
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(".WriteEID(")
          /*g_try(v2:"Result",loop:true) */
          Result = Core.F_CALL(C_Generate_g_expression,ARGS(g0159.Args.At(2-1).ToEID(),EID{C_object.Id(),0}))
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(",")
          /*g_try(v2:"Result",loop:true) */
          Result = F_Generate_eid_expression_any(g0159.Args.At(3-1),Optimize.C_EID,lvar)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(")")
          Result = EVOID
          }}}
          }  else if (g0159.Arg.Selector.Id() == Core.C_nth_write.Id()) { 
          /*g_try(v2:"Result",loop:true) */
          Result = Core.F_CALL(C_Generate_g_expression,ARGS(g0159.Args.At(1-1).ToEID(),EID{C_list.Id(),0}))
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(".WriteEID(")
          /*g_try(v2:"Result",loop:true) */
          Result = Core.F_CALL(C_Generate_g_expression,ARGS(g0159.Args.At(2-1).ToEID(),EID{C_integer.Id(),0}))
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(",")
          /*g_try(v2:"Result",loop:true) */
          Result = F_Generate_eid_expression_any(g0159.Args.At(3-1),Optimize.C_EID,lvar)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(")")
          Result = EVOID
          }}}
          } else {
          /*g_try(v2:"Result",loop:true) */
          Result = Core.F_CALL(C_Generate_g_expression,ARGS(g0159.Args.At(1-1).ToEID(),EID{C_Variable.Id(),0}))
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(".WriteEID(")
          /*g_try(v2:"Result",loop:true) */
          Result = F_Generate_eid_expression_any(g0159.Args.At(2-1),Optimize.C_EID,lvar)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(")")
          Result = EVOID
          }}
          } 
        } 
      } else {
      Result = Core.F_CALL(C_Generate_g_expression,ARGS(x.ToEID(),EID{s.Id(),0}))
      } 
    return Result} 
  
// The EID go function for: eid_expression @ any (throw: true) 
func E_Generate_eid_expression_any (x EID,s EID,lvar EID) EID { 
    return F_Generate_eid_expression_any(ANY(x),ToClass(OBJ(s)),ToList(OBJ(lvar)) )} 
  
//**********************************************************************
//*          Part 3: Basic control structures                          *
//**********************************************************************
// The re-entry definition (called within g_statement, not directly)
// if functional, the best compiling is into an expression
// s is the expected go type (as a class) + void + EID
// v is nil or a string (name of the variable)
// note that only 3 additional parameters are used since err is recomputed
/* {1} The go function for: statement(self:any,s:class,v:string,loop:any) [status=1] */
func F_Generate_statement_any (self *ClaireAny ,s *ClaireClass ,v *ClaireString ,loop *ClaireAny ) EID { 
    var Result EID 
    var g0161I *ClaireBoolean  
    var try_1 EID 
    /*g_try(v2:"try_1",loop:false) */
    try_1 = F_Generate_g_clean_any(self)
    /* ERROR PROTECTION INSERTED (g0161I-Result) */
    if ErrorIn(try_1) {Result = try_1
    } else {
    g0161I = ToBoolean(OBJ(try_1))
    if (g0161I == CTRUE) { 
      var g0162I *ClaireBoolean  
      var try_2 EID 
      /*g_try(v2:"try_2",loop:false) */
      { 
        var v_and3 *ClaireBoolean  
        
        v_and3 = Core.F__I_equal_any(s.Id(),C_void.Id())
        if (v_and3 == CFALSE) {try_2 = EID{CFALSE.Id(),0}
        } else { 
          var try_3 EID 
          /*g_try(v2:"try_3",loop:false) */
          { var arg_4 *ClaireClass  
            _ = arg_4
            var try_5 EID 
            /*g_try(v2:"try_5",loop:false) */
            try_5 = Language.F_static_type_any(self)
            /* ERROR PROTECTION INSERTED (arg_4-try_3) */
            if ErrorIn(try_5) {try_3 = try_5
            } else {
            arg_4 = ToClass(OBJ(try_5))
            try_3 = EID{Core.F__I_equal_any(arg_4.Id(),C_void.Id()).Id(),0}
            }
            } 
          /* ERROR PROTECTION INSERTED (v_and3-try_2) */
          if ErrorIn(try_3) {try_2 = try_3
          } else {
          v_and3 = ToBoolean(OBJ(try_3))
          if (v_and3 == CFALSE) {try_2 = EID{CFALSE.Id(),0}
          } else { 
            try_2 = EID{CTRUE.Id(),0}} 
          } 
        }
        } 
      /* ERROR PROTECTION INSERTED (g0162I-Result) */
      if ErrorIn(try_2) {Result = try_2
      } else {
      g0162I = ToBoolean(OBJ(try_2))
      if (g0162I == CTRUE) { 
        F_c_princ_string(v)
        PRINC(" = ")
        /*g_try(v2:"Result",loop:true) */
        Result = Core.F_CALL(C_Generate_g_expression,ARGS(self.ToEID(),EID{s.Id(),0}))
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        F_Generate_breakline_void()
        PRINC("")
        Result = EVOID
        }
        }  else if (self.Isa.IsIn(Language.C_If) == CTRUE) { 
        Result = Core.F_CALL(C_Generate_g_statement,ARGS(self.ToEID(),
          EID{s.Id(),0},
          EID{(v).Id(),0},
          EID{CFALSE.Id(),0},
          loop.ToEID()))
        }  else if (self.Isa.IsIn(Reader.C_delimiter) == CTRUE) { 
        Result = ToException(Core.C_general_error.Make(MakeString("[201] Loose delimiter in program: ~S").Id(),MakeConstantList(self).Id())).Close()
        } else {
        /*g_try(v2:"Result",loop:true) */
        Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).StatExp(self,C_void)
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        if (s.Id() == Optimize.C_EID.Id()) { 
          F_c_princ_string(v)
          PRINC(" = EVOID")
          F_Generate_breakline_void()
          PRINC("")
          Result = EVOID
          } else {
          Result = EID{CFALSE.Id(),0}
          } 
        }
        } 
      }
      } else {
      { var arg_6 *ClaireBoolean  
        _ = arg_6
        var try_7 EID 
        /*g_try(v2:"try_7",loop:false) */
        try_7 = Optimize.F_Compile_g_throw_any(self)
        /* ERROR PROTECTION INSERTED (arg_6-Result) */
        if ErrorIn(try_7) {Result = try_7
        } else {
        arg_6 = ToBoolean(OBJ(try_7))
        Result = Core.F_CALL(C_Generate_g_statement,ARGS(self.ToEID(),
          EID{s.Id(),0},
          EID{(v).Id(),0},
          EID{arg_6.Id(),0},
          loop.ToEID()))
        }
        } 
      } 
    }
    return Result} 
  
// The EID go function for: statement @ any (throw: true) 
func E_Generate_statement_any (self EID,s EID,v EID,loop EID) EID { 
    return F_Generate_statement_any(ANY(self),
      ToClass(OBJ(s)),
      ToString(OBJ(v)),
      ANY(loop) )} 
  
// make a statement from an expression (in C++ we need a ; - with go a breakline is enough)
// we do not want to place the result in a variable (see upper)
// will get simpler once we have a stable compiler without to_C and to_CL
/* {1} The go function for: stat_exp(c:go_producer,self:any,s:class) [status=1] */
func (c *GenerateGoProducer ) StatExp (self *ClaireAny ,s *ClaireClass ) EID { 
    var Result EID 
    var g0163I *ClaireBoolean  
    var try_1 EID 
    /*g_try(v2:"try_1",loop:false) */
    try_1 = Optimize.F_Compile_designated_ask_any(self)
    /* ERROR PROTECTION INSERTED (g0163I-Result) */
    if ErrorIn(try_1) {Result = try_1
    } else {
    g0163I = ToBoolean(OBJ(try_1))
    if (g0163I == CTRUE) { 
      Result = F_Generate_breakline_void().ToEID()
      } else {
      /*g_try(v2:"Result",loop:true) */
      Result = Core.F_CALL(C_Generate_g_expression,ARGS(self.ToEID(),EID{s.Id(),0}))
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      F_Generate_breakline_void()
      PRINC("")
      Result = EVOID
      }
      } 
    }
    return Result} 
  
// The EID go function for: stat_exp @ go_producer (throw: true) 
func E_Generate_stat_exp_go_producer (c EID,self EID,s EID) EID { 
    return ToGenerateGoProducer(OBJ(c)).StatExp(ANY(self),ToClass(OBJ(s)) )} 
  
// a DO is a simple go block if there are not errors, a chain otherwise
// the chain means multiple nestings when an error occurs since the rest of the DO must not be
// this is why the close_try(count) are called at the end, to close the embedded ifs (ErrorIn(e))
// we use a specific method code_statement with an additional parameter %need which is true by default
/* {1} The go function for: g_statement(self:Do,e:class,v:string,err:boolean,loop:any) [status=1] */
func F_Generate_g_statement_Do (self *Language.Do ,e *ClaireClass ,v *ClaireString ,err *ClaireBoolean ,loop *ClaireAny ) EID { 
    var Result EID 
    Result = F_Generate_do_statement_Do(self,
      e,
      v,
      err,
      loop,
      CTRUE)
    return Result} 
  
// The EID go function for: g_statement @ Do (throw: true) 
func E_Generate_g_statement_Do (self EID,e EID,v EID,err EID,loop EID) EID { 
    return F_Generate_g_statement_Do(Language.To_Do(OBJ(self)),
      ToClass(OBJ(e)),
      ToString(OBJ(v)),
      ToBoolean(OBJ(err)),
      ANY(loop) )} 
  
// %need is true = the last argument is needed in v
/* {1} The go function for: do_statement(self:Do,e:class,v:string,err:boolean,loop:any,%need:boolean) [status=1] */
func F_Generate_do_statement_Do (self *Language.Do ,e *ClaireClass ,v *ClaireString ,err *ClaireBoolean ,loop *ClaireAny ,_Zneed *ClaireBoolean ) EID { 
    var Result EID 
    if (self.Args.Length() == 1) { 
      Result = F_Generate_statement_any(self.Args.At(1-1),e,v,loop)
      }  else if (err != CTRUE) { 
      { var l *ClaireList   = self.Args
        { var m int  = l.Length()
          { var n int  = 1
            { var g0164 int  = m
              _ = g0164
              Result= EID{CFALSE.Id(),0}
              for (n <= g0164) { 
                /* While stat, v:"Result" loop:true */
                var loop_1 EID 
                _ = loop_1
                { 
                /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
                { var arg_2 *ClaireClass  
                  _ = arg_2
                  if (n == m) { 
                    arg_2 = e
                    } else {
                    arg_2 = C_void
                    } 
                  loop_1 = F_Generate_statement_any(l.At(n-1),arg_2,v,loop)
                  } 
                /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
                if ErrorIn(loop_1) {Result = loop_1
                break
                } else {
                n = (n+1)
                }
                /* try?:false, v2:"v_while7" loop will be:tuple("Result", EID) */
                } 
              }
              } 
            } 
          } 
        } 
      } else {
      { var l *ClaireList   = self.Args
        { var m int  = l.Length()
          { var count_if int  = 0
            _ = count_if
            /*g_try(v2:"Result",loop:true) */
            { var n int  = 1
              { var g0165 int  = m
                _ = g0165
                Result= EID{CFALSE.Id(),0}
                for (n <= g0165) { 
                  /* While stat, v:"Result" loop:true */
                  var loop_3 EID 
                  _ = loop_3
                  { 
                  /*g_try(v2:"loop_3",loop:tuple("Result", EID)) */
                  { var x *ClaireAny   = l.At(n-1)
                    var g0166I *ClaireBoolean  
                    var try_4 EID 
                    /*g_try(v2:"try_4",loop:false) */
                    try_4 = Optimize.F_Compile_g_throw_any(x)
                    /* ERROR PROTECTION INSERTED (g0166I-loop_3) */
                    if ErrorIn(try_4) {loop_3 = try_4
                    } else {
                    g0166I = ToBoolean(OBJ(try_4))
                    if (g0166I == CTRUE) { 
                      if ((n < m) || 
                          (C_tuple.Id() == loop.Isa.Id())) { 
                        count_if = (count_if+1)
                        /*g_try(v2:"loop_3",loop:tuple("Result", EID)) */
                        loop_3 = Core.F_CALL(C_Generate_g_try,ARGS(x.ToEID(),
                          EID{(v).Id(),0},
                          EID{Optimize.C_EID.Id(),0},
                          EID{(v).Id(),0},
                          loop.ToEID()))
                        /* ERROR PROTECTION INSERTED (loop_3-loop_3) */
                        if ErrorIn(loop_3) {Result = loop_3
                        break
                        } else {
                        }
                        } else {
                        loop_3 = Core.F_CALL(C_Generate_g_statement,ARGS(x.ToEID(),
                          EID{Optimize.C_EID.Id(),0},
                          EID{(v).Id(),0},
                          EID{CTRUE.Id(),0},
                          loop.ToEID()))
                        } 
                      } else {
                      { var arg_5 *ClaireClass  
                        _ = arg_5
                        if ((n == m) && 
                            (_Zneed == CTRUE)) { 
                          arg_5 = e
                          } else {
                          arg_5 = C_void
                          } 
                        loop_3 = F_Generate_statement_any(x,arg_5,v,loop)
                        } 
                      } 
                    }
                    } 
                  /* ERROR PROTECTION INSERTED (loop_3-loop_3) */
                  if ErrorIn(loop_3) {Result = loop_3
                  break
                  } else {
                  n = (n+1)
                  }
                  /* try?:false, v2:"v_while8" loop will be:tuple("Result", EID) */
                  } 
                }
                } 
              } 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            F_Generate_close_try_integer(count_if)
            Result = EVOID
            }
            } 
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: do_statement @ Do (throw: true) 
func E_Generate_do_statement_Do (self EID,e EID,v EID,err EID,loop EID,_Zneed EID) EID { 
    return F_Generate_do_statement_Do(Language.To_Do(OBJ(self)),
      ToClass(OBJ(e)),
      ToString(OBJ(v)),
      ToBoolean(OBJ(err)),
      ANY(loop),
      ToBoolean(OBJ(_Zneed)) )} 
  
// balance 
// a Let is a local variable declaration 
// in CLAIRE 4, a block is anything that fits between {} hence inner/outer is not necessary
// AXIOM if err is true, we require that e = any
/* {1} The go function for: g_statement(self:Let,e:class,v:string,err:boolean,loop:any) [status=1] */
func F_Generate_g_statement_Let (self *Language.Let ,e *ClaireClass ,v *ClaireString ,err *ClaireBoolean ,loop *ClaireAny ) EID { 
    var Result EID 
    var g0167I *ClaireBoolean  
    var try_1 EID 
    /*g_try(v2:"try_1",loop:false) */
    try_1 = F_Generate_let_eid_ask_Let(self)
    /* ERROR PROTECTION INSERTED (g0167I-Result) */
    if ErrorIn(try_1) {Result = try_1
    } else {
    g0167I = ToBoolean(OBJ(try_1))
    if (g0167I == CTRUE) { 
      Result = F_Generate_g_eid_stat_Let(self,
        e,
        v,
        err,
        loop)
      } else {
      { var ns *ClaireString  
        var try_2 EID 
        /*g_try(v2:"try_2",loop:false) */
        try_2 = F_Generate_c_string_go_producer2(ToGenerateGoProducer(Optimize.C_PRODUCER.Value),self.ClaireVar.Pname)
        /* ERROR PROTECTION INSERTED (ns-Result) */
        if ErrorIn(try_2) {Result = try_2
        } else {
        ns = ToString(OBJ(try_2))
        if ((ns.At(1) == 'C') && 
            (ns.At(2) == '%')) { 
          self.ClaireVar.Pname = Core.F_gensym_void()
          /*symbol->symbol*/} 
        { var v2 *ClaireString  
          var try_3 EID 
          /*g_try(v2:"try_3",loop:false) */
          try_3 = F_Generate_c_string_go_producer1(ToGenerateGoProducer(Optimize.C_PRODUCER.Value),self.ClaireVar)
          /* ERROR PROTECTION INSERTED (v2-Result) */
          if ErrorIn(try_3) {Result = try_3
          } else {
          v2 = ToString(OBJ(try_3))
          { var x *ClaireAny   = self.Value
            { var f *ClaireBoolean  
              var try_4 EID 
              /*g_try(v2:"try_4",loop:false) */
              try_4 = F_Generate_g_clean_any(x)
              /* ERROR PROTECTION INSERTED (f-Result) */
              if ErrorIn(try_4) {Result = try_4
              } else {
              f = ToBoolean(OBJ(try_4))
              { var try_ask *ClaireBoolean  
                var try_5 EID 
                /*g_try(v2:"try_5",loop:false) */
                try_5 = Optimize.F_Compile_g_throw_any(x)
                /* ERROR PROTECTION INSERTED (try_ask-Result) */
                if ErrorIn(try_5) {Result = try_5
                } else {
                try_ask = ToBoolean(OBJ(try_5))
                { var ev *ClaireClass   = self.ClaireVar.Range.Class_I()
                  F_Generate_let_block_void()
                  F_Generate_var_declaration_string(v2,ev,0)
                  /*g_try(v2:"Result",loop:true) */
                  if (f == CTRUE) { 
                    PRINC(" = ")
                    /*g_try(v2:"Result",loop:true) */
                    Result = Core.F_CALL(C_Generate_g_expression,ARGS(x.ToEID(),EID{ev.Id(),0}))
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    PRINC("")
                    Result = EVOID
                    }
                    } else {
                    Result = EID{CFALSE.Id(),0}
                    } 
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  F_Generate_breakline_void()
                  if (Language.F_Language_occurexact_any(self.Arg,self.ClaireVar) <= 1) { 
                    
                    PRINC("_ = ")
                    F_princ_string(v2)
                    F_Generate_breakline_void()
                    PRINC("")
                    } 
                  /*g_try(v2:"Result",loop:true) */
                  if (try_ask == CTRUE) { 
                    Result = Core.F_CALL(C_Generate_g_try,ARGS(x.ToEID(),
                      EID{(v2).Id(),0},
                      EID{ev.Id(),0},
                      EID{(v).Id(),0},
                      EID{CFALSE.Id(),0}))
                    }  else if (f != CTRUE) { 
                    Result = F_Generate_statement_any(x,ev,v2,loop)
                    } else {
                    Result = EID{CFALSE.Id(),0}
                    } 
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  /*g_try(v2:"Result",loop:true) */
                  Result = F_Generate_statement_any(self.Arg,e,v,loop)
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  if (try_ask == CTRUE) { 
                    F_Generate_close_try_integer(1)
                    } 
                  F_Generate_close_block_string(MakeString("Let"))
                  Result = EVOID
                  }}}
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
  
// The EID go function for: g_statement @ Let (throw: true) 
func E_Generate_g_statement_Let (self EID,e EID,v EID,err EID,loop EID) EID { 
    return F_Generate_g_statement_Let(Language.To_Let(OBJ(self)),
      ToClass(OBJ(e)),
      ToString(OBJ(v)),
      ToBoolean(OBJ(err)),
      ANY(loop) )} 
  
// then we must close the chain
// detect a pattern (that could be expanded) where the variable should be compiled as EID because
// the value may trigger an error and the body uses v once at the end (this could be : the body is EID-friendly)
// this current specific pattern is introduced to optimize eval_message
/* {1} The go function for: let_eid?(self:Let) [status=1] */
func F_Generate_let_eid_ask_Let (self *Language.Let ) EID { 
    var Result EID 
    { var v *ClaireVariable   = self.ClaireVar
      { var y *ClaireAny   = self.Arg
        { 
          /* Or stat: v="Result", loop=true */
          var v_or4 *ClaireBoolean  
          
          /* Or stat: try (((Compile/g_throw @ any(value @ Let(self))) | (eid_provide? @ any(value @ Let(self)))) & (= @ any(Language/occurrence @ any(arg @ Let(self),var @ Instruction_with_var(self)),1)) & (if (inherit? @ class(owner @ any(y),Do)) let g0168:Do := (<y:Do>) in ((not @ any(Compile/g_throw @ any(g0168))) & (= @ any(last @ list(args @ Do(g0168)),v))) else false)) with try:true, v="Result", loop=true */
          var try_1 EID 
          /*g_try(v2:"try_1",loop:true) */
          { 
            var v_and5 *ClaireBoolean  
            
            var try_2 EID 
            /*g_try(v2:"try_2",loop:false) */
            { 
              /* Or stat: v="try_2", loop=false */
              var v_or6 *ClaireBoolean  
              
              /* Or stat: try Compile/g_throw @ any(value @ Let(self)) with try:true, v="try_2", loop=false */
              var try_3 EID 
              /*g_try(v2:"try_3",loop:false) */
              try_3 = Optimize.F_Compile_g_throw_any(self.Value)
              /* ERROR PROTECTION INSERTED (v_or6-try_2) */
              if ErrorIn(try_3) {try_2 = try_3
              } else {
              v_or6 = ToBoolean(OBJ(try_3))
              if (v_or6 == CTRUE) {try_2 = EID{CTRUE.Id(),0}
              } else { 
                /* Or stat: try eid_provide? @ any(value @ Let(self)) with try:false, v="try_2", loop=false */
                v_or6 = F_Generate_eid_provide_ask_any(self.Value)
                if (v_or6 == CTRUE) {try_2 = EID{CTRUE.Id(),0}
                } else { 
                  try_2 = EID{CFALSE.Id(),0}} 
                } 
              }
              } 
            /* ERROR PROTECTION INSERTED (v_and5-try_1) */
            if ErrorIn(try_2) {try_1 = try_2
            } else {
            v_and5 = ToBoolean(OBJ(try_2))
            if (v_and5 == CFALSE) {try_1 = EID{CFALSE.Id(),0}
            } else { 
              v_and5 = Equal(MakeInteger(Language.F_occurrence_any(self.Arg,self.ClaireVar)).Id(),MakeInteger(1).Id())
              if (v_and5 == CFALSE) {try_1 = EID{CFALSE.Id(),0}
              } else { 
                var try_4 EID 
                /*g_try(v2:"try_4",loop:false) */
                if (y.Isa.IsIn(Language.C_Do) == CTRUE) { 
                  { var g0168 *Language.Do   = Language.To_Do(y)
                    { 
                      var v_and10 *ClaireBoolean  
                      
                      var try_5 EID 
                      /*g_try(v2:"try_5",loop:false) */
                      { var arg_6 *ClaireBoolean  
                        _ = arg_6
                        var try_7 EID 
                        /*g_try(v2:"try_7",loop:false) */
                        try_7 = Optimize.F_Compile_g_throw_any(g0168.Id())
                        /* ERROR PROTECTION INSERTED (arg_6-try_5) */
                        if ErrorIn(try_7) {try_5 = try_7
                        } else {
                        arg_6 = ToBoolean(OBJ(try_7))
                        try_5 = EID{arg_6.Not.Id(),0}
                        }
                        } 
                      /* ERROR PROTECTION INSERTED (v_and10-try_4) */
                      if ErrorIn(try_5) {try_4 = try_5
                      } else {
                      v_and10 = ToBoolean(OBJ(try_5))
                      if (v_and10 == CFALSE) {try_4 = EID{CFALSE.Id(),0}
                      } else { 
                        var try_8 EID 
                        /*g_try(v2:"try_8",loop:false) */
                        { var arg_9 *ClaireAny  
                          _ = arg_9
                          var try_10 EID 
                          /*g_try(v2:"try_10",loop:false) */
                          try_10 = Core.F_last_list(g0168.Args)
                          /* ERROR PROTECTION INSERTED (arg_9-try_8) */
                          if ErrorIn(try_10) {try_8 = try_10
                          } else {
                          arg_9 = ANY(try_10)
                          try_8 = EID{Equal(arg_9,v.Id()).Id(),0}
                          }
                          } 
                        /* ERROR PROTECTION INSERTED (v_and10-try_4) */
                        if ErrorIn(try_8) {try_4 = try_8
                        } else {
                        v_and10 = ToBoolean(OBJ(try_8))
                        if (v_and10 == CFALSE) {try_4 = EID{CFALSE.Id(),0}
                        } else { 
                          try_4 = EID{CTRUE.Id(),0}} 
                        } 
                      }}
                      } 
                    } 
                  } else {
                  try_4 = EID{CFALSE.Id(),0}
                  } 
                /* ERROR PROTECTION INSERTED (v_and5-try_1) */
                if ErrorIn(try_4) {try_1 = try_4
                } else {
                v_and5 = ToBoolean(OBJ(try_4))
                if (v_and5 == CFALSE) {try_1 = EID{CFALSE.Id(),0}
                } else { 
                  try_1 = EID{CTRUE.Id(),0}} 
                } 
              } 
            }}
            } 
          /* ERROR PROTECTION INSERTED (v_or4-Result) */
          if ErrorIn(try_1) {Result = try_1
          } else {
          v_or4 = ToBoolean(OBJ(try_1))
          if (v_or4 == CTRUE) {Result = EID{CTRUE.Id(),0}
          } else { 
            /* Or stat: try = @ any(range @ Variable(v),EID) with try:false, v="Result", loop=true */
            v_or4 = Equal(v.Range.Id(),Optimize.C_EID.Id())
            if (v_or4 == CTRUE) {Result = EID{CTRUE.Id(),0}
            } else { 
              Result = EID{CFALSE.Id(),0}} 
            } 
          }
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: let_eid? @ Let (throw: true) 
func E_Generate_let_eid_ask_Let (self EID) EID { 
    return F_Generate_let_eid_ask_Let(Language.To_Let(OBJ(self)) )} 
  
// force EID compiling : back door :)
// the corresponding compiling (embeds the Do)
/* {1} The go function for: g_eid_stat(self:Let,e:class,v:string,err:boolean,loop:any) [status=1] */
func F_Generate_g_eid_stat_Let (self *Language.Let ,e *ClaireClass ,v *ClaireString ,err *ClaireBoolean ,loop *ClaireAny ) EID { 
    var Result EID 
    { var v2 *ClaireString  
      var try_1 EID 
      /*g_try(v2:"try_1",loop:false) */
      try_1 = F_Generate_c_string_go_producer1(ToGenerateGoProducer(Optimize.C_PRODUCER.Value),self.ClaireVar)
      /* ERROR PROTECTION INSERTED (v2-Result) */
      if ErrorIn(try_1) {Result = try_1
      } else {
      v2 = ToString(OBJ(try_1))
      { var x *ClaireAny   = self.Value
        { var try_ask *ClaireBoolean  
          var try_2 EID 
          /*g_try(v2:"try_2",loop:false) */
          try_2 = Optimize.F_Compile_g_throw_any(x)
          /* ERROR PROTECTION INSERTED (try_ask-Result) */
          if ErrorIn(try_2) {Result = try_2
          } else {
          try_ask = ToBoolean(OBJ(try_2))
          F_Generate_new_block_string(MakeString("LetE"))
          F_Generate_var_declaration_string(v2,Optimize.C_EID,0)
          F_Generate_breakline_void()
          /*g_try(v2:"Result",loop:true) */
          if (try_ask == CTRUE) { 
            Result = Core.F_CALL(C_Generate_g_try,ARGS(x.ToEID(),
              EID{(v2).Id(),0},
              EID{Optimize.C_EID.Id(),0},
              EID{(v).Id(),0},
              EID{CFALSE.Id(),0}))
            } else {
            F_princ_string(v2)
            PRINC(" = ")
            /*g_try(v2:"Result",loop:true) */
            Result = Core.F_CALL(C_Generate_g_expression,ARGS(x.ToEID(),EID{Optimize.C_EID.Id(),0}))
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            F_Generate_breakline_void()
            PRINC("")
            Result = EVOID
            }
            } 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          /*g_try(v2:"Result",loop:true) */
          { var y *Language.Do   = Language.To_Do(self.Arg)
            { var n int  = y.Args.Length()
              _ = n
              /*g_try(v2:"Result",loop:true) */
              { var i int  = 1
                { var g0169 int  = (n-1)
                  _ = g0169
                  Result= EID{CFALSE.Id(),0}
                  for (i <= g0169) { 
                    /* While stat, v:"Result" loop:true */
                    var loop_3 EID 
                    _ = loop_3
                    { 
                    /*g_try(v2:"loop_3",loop:tuple("Result", EID)) */
                    loop_3 = F_Generate_statement_any(y.Args.At(i-1),C_void,v,loop)
                    /* ERROR PROTECTION INSERTED (loop_3-loop_3) */
                    if ErrorIn(loop_3) {Result = loop_3
                    break
                    } else {
                    i = (i+1)
                    }
                    /* try?:false, v2:"v_while9" loop will be:tuple("Result", EID) */
                    } 
                  }
                  } 
                } 
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              F_princ_string(v)
              PRINC(" = ")
              F_princ_string(v2)
              PRINC("")
              Result = EVOID
              }
              } 
            } 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          if (try_ask == CTRUE) { 
            F_Generate_close_try_integer(1)
            } 
          F_Generate_close_block_string(MakeString("LetE"))
          Result = EVOID
          }}
          }
          } 
        } 
      }
      } 
    return Result} 
  
// The EID go function for: g_eid_stat @ Let (throw: true) 
func E_Generate_g_eid_stat_Let (self EID,e EID,v EID,err EID,loop EID) EID { 
    return F_Generate_g_eid_stat_Let(Language.To_Let(OBJ(self)),
      ToClass(OBJ(e)),
      ToString(OBJ(v)),
      ToBoolean(OBJ(err)),
      ANY(loop) )} 
  
// then we must close the chain
// makes a bag from a list of  statements.
// The value cannot be ignored: it is considered as an error (a do should have been used)
// there are two patterns depending if self.of is known : MakeEmptyX(t) or MakeEmptyX(any)
/* {1} The go function for: g_statement(self:Construct,s:class,v:string,err:boolean,loop:any) [status=1] */
func F_Generate_g_statement_Construct (self *Language.Construct ,s *ClaireClass ,v *ClaireString ,err *ClaireBoolean ,loop *ClaireAny ) EID { 
    var Result EID 
    /*g_try(v2:"Result",loop:true) */
    if (F_boolean_I_any((v).Id()).Id() != CTRUE.Id()) { 
      Result = ToException(Core.C_general_error.Make(MakeString("[202] A do should have been used for ~S").Id(),MakeConstantList(self.Id()).Id())).Close()
      } else {
      Result = EID{CFALSE.Id(),0}
      } 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    { var v2 *ClaireString   = MakeString("v_bag_arg")
      { var kind *ClaireClass  
        var try_1 EID 
        /*g_try(v2:"try_1",loop:false) */
        if (self.Isa.IsIn(Language.C_List) == CTRUE) { 
          try_1 = EID{C_list.Id(),0}
          }  else if (self.Isa.IsIn(Language.C_Set) == CTRUE) { 
          try_1 = EID{C_set.Id(),0}
          }  else if (self.Isa.IsIn(Language.C_Tuple) == CTRUE) { 
          try_1 = EID{C_tuple.Id(),0}
          } else {
          try_1 = ToException(Core.C_general_error.Make(MakeString("CONSTRUCT BUG: ~S").Id(),MakeConstantList(self.Id()).Id())).Close()
          } 
        /* ERROR PROTECTION INSERTED (kind-Result) */
        if ErrorIn(try_1) {Result = try_1
        } else {
        kind = ToClass(OBJ(try_1))
        { var count_try int  = 0
          _ = count_try
          { var t *ClaireType  
            _ = t
            if (Core.F_get_property(C_of,ToObject(self.Id())) != CNULL) { 
              t = ToType(OBJ(Core.F_CALL(C_of,ARGS(EID{self.Id(),0}))))
              } else {
              t = ToType(C_void.Id())
              } 
            F_Generate_new_block_string(MakeString("Construct"))
            /*g_try(v2:"Result",loop:true) */
            var g0174I *ClaireBoolean  
            var try_2 EID 
            /*g_try(v2:"try_2",loop:false) */
            { var arg_3 *ClaireAny  
              _ = arg_3
              var try_4 EID 
              /*g_try(v2:"try_4",loop:false) */
              { 
                var x *ClaireAny  
                _ = x
                try_4= EID{CFALSE.Id(),0}
                var x_support *ClaireList  
                x_support = self.Args
                x_len := x_support.Length()
                for i_it := 0; i_it < x_len; i_it++ { 
                  x = x_support.At(i_it)
                  var loop_5 EID 
                  _ = loop_5
                  /*g_try(v2:"loop_5",loop:tuple("try_4", EID)) */
                  var g0175I *ClaireBoolean  
                  var try_6 EID 
                  /*g_try(v2:"try_6",loop:false) */
                  { var arg_7 *ClaireBoolean  
                    _ = arg_7
                    var try_8 EID 
                    /*g_try(v2:"try_8",loop:false) */
                    try_8 = F_Generate_g_clean_any(x)
                    /* ERROR PROTECTION INSERTED (arg_7-try_6) */
                    if ErrorIn(try_8) {try_6 = try_8
                    } else {
                    arg_7 = ToBoolean(OBJ(try_8))
                    try_6 = EID{arg_7.Not.Id(),0}
                    }
                    } 
                  /* ERROR PROTECTION INSERTED (g0175I-loop_5) */
                  if ErrorIn(try_6) {loop_5 = try_6
                  } else {
                  g0175I = ToBoolean(OBJ(try_6))
                  if (g0175I == CTRUE) { 
                    try_4 = EID{CTRUE.Id(),0}
                    break
                    } else {
                    loop_5 = EID{CFALSE.Id(),0}
                    } 
                  }
                  /* ERROR PROTECTION INSERTED (loop_5-try_4) */
                  if ErrorIn(loop_5) {try_4 = loop_5
                  break
                  } else {
                  }
                  } 
                } 
              /* ERROR PROTECTION INSERTED (arg_3-try_2) */
              if ErrorIn(try_4) {try_2 = try_4
              } else {
              arg_3 = ANY(try_4)
              try_2 = EID{F_boolean_I_any(arg_3).Id(),0}
              }
              } 
            /* ERROR PROTECTION INSERTED (g0174I-Result) */
            if ErrorIn(try_2) {Result = try_2
            } else {
            g0174I = ToBoolean(OBJ(try_2))
            if (g0174I == CTRUE) { 
              F_Generate_var_declaration_string(v2,C_any,1)
              Result = EVOID
              } else {
              Result = EID{CFALSE.Id(),0}
              } 
            }
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            /*g_try(v2:"Result",loop:true) */
            F_c_princ_string(v)
            PRINC("= ")
            /*g_try(v2:"Result",loop:true) */
            Result = F_Generate_cast_prefix_class(kind,s)
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            /*g_try(v2:"Result",loop:true) */
            if (kind.Id() == C_tuple.Id()) { 
              PRINC("Make")
              Result = EVOID
              } else {
              /*g_try(v2:"Result",loop:true) */
              if (Core.F_get_property(C_of,ToObject(self.Id())) != CNULL) { 
                { var arg_9 *ClaireAny  
                  _ = arg_9
                  var try_10 EID 
                  /*g_try(v2:"try_10",loop:false) */
                  try_10 = Core.F_CALL(Optimize.C_c_code,ARGS(EID{t.Id(),0},EID{C_object.Id(),0}))
                  /* ERROR PROTECTION INSERTED (arg_9-Result) */
                  if ErrorIn(try_10) {Result = try_10
                  } else {
                  arg_9 = ANY(try_10)
                  Result = Core.F_CALL(C_Generate_g_expression,ARGS(arg_9.ToEID(),EID{C_type.Id(),0}))
                  }
                  } 
                } else {
                PRINC("ToType(CEMPTY.Id())")
                Result = EVOID
                } 
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              PRINC(".")
              Result = EVOID
              }
              } 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC("Empty")
            F_Generate_cap_short_symbol(kind.Name)
            PRINC("()")
            F_Generate_cast_post_class(kind,s)
            PRINC("")
            Result = EVOID
            }}
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            /*g_try(v2:"Result",loop:true) */
            { 
              var x *ClaireAny  
              _ = x
              Result= EID{CFALSE.Id(),0}
              var x_support *ClaireList  
              x_support = self.Args
              x_len := x_support.Length()
              for i_it := 0; i_it < x_len; i_it++ { 
                x = x_support.At(i_it)
                var loop_11 EID 
                _ = loop_11
                /*g_try(v2:"loop_11",loop:tuple("Result", EID)) */
                { var f *ClaireBoolean  
                  var try_12 EID 
                  /*g_try(v2:"try_12",loop:false) */
                  try_12 = F_Generate_g_clean_any(x)
                  /* ERROR PROTECTION INSERTED (f-loop_11) */
                  if ErrorIn(try_12) {loop_11 = try_12
                  } else {
                  f = ToBoolean(OBJ(try_12))
                  { var try_ask *ClaireBoolean  
                    var try_13 EID 
                    /*g_try(v2:"try_13",loop:false) */
                    try_13 = Optimize.F_Compile_g_throw_any(x)
                    /* ERROR PROTECTION INSERTED (try_ask-loop_11) */
                    if ErrorIn(try_13) {loop_11 = try_13
                    } else {
                    try_ask = ToBoolean(OBJ(try_13))
                    F_Generate_breakline_void()
                    if (try_ask == CTRUE) { 
                      count_try = (count_try+1)
                      } 
                    /*g_try(v2:"loop_11",loop:tuple("Result", EID)) */
                    if (f != CTRUE) { 
                      if (try_ask == CTRUE) { 
                        loop_11 = Core.F_CALL(C_Generate_g_try,ARGS(x.ToEID(),
                          EID{(v2).Id(),0},
                          EID{C_any.Id(),0},
                          EID{(v).Id(),0},
                          EID{CFALSE.Id(),0}))
                        } else {
                        loop_11 = F_Generate_statement_any(x,C_any,v2,loop)
                        } 
                      } else {
                      loop_11 = EID{CFALSE.Id(),0}
                      } 
                    /* ERROR PROTECTION INSERTED (loop_11-loop_11) */
                    if ErrorIn(loop_11) {Result = loop_11
                    break
                    } else {
                    /*g_try(v2:"loop_11",loop:tuple("Result", EID)) */
                    /*g_try(v2:"loop_11",loop:tuple("Result", EID)) */
                    loop_11 = F_Generate_cast_prefix_class(s,kind)
                    /* ERROR PROTECTION INSERTED (loop_11-loop_11) */
                    if ErrorIn(loop_11) {Result = loop_11
                    break
                    } else {
                    F_c_princ_string(v)
                    F_Generate_cast_post_class(s,kind)
                    PRINC(".AddFast(")
                    /*g_try(v2:"loop_11",loop:tuple("Result", EID)) */
                    if (f == CTRUE) { 
                      loop_11 = Core.F_CALL(C_Generate_g_expression,ARGS(x.ToEID(),EID{C_any.Id(),0}))
                      } else {
                      F_c_princ_string(v2)
                      loop_11 = EVOID
                      } 
                    /* ERROR PROTECTION INSERTED (loop_11-loop_11) */
                    if ErrorIn(loop_11) {Result = loop_11
                    break
                    } else {
                    PRINC(")")
                    loop_11 = EVOID
                    }}
                    /* ERROR PROTECTION INSERTED (loop_11-loop_11) */
                    if ErrorIn(loop_11) {Result = loop_11
                    break
                    } else {
                    }}
                    }
                    } 
                  }
                  } 
                /* ERROR PROTECTION INSERTED (loop_11-Result) */
                if ErrorIn(loop_11) {Result = loop_11
                break
                } else {
                }
                } 
              } 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            F_Generate_close_try_integer(count_try)
            F_Generate_close_block_string(MakeString("Construct"))
            Result = EVOID
            }}}
            } 
          } 
        }
        } 
      } 
    }
    return Result} 
  
// The EID go function for: g_statement @ Construct (throw: true) 
func E_Generate_g_statement_Construct (self EID,s EID,v EID,err EID,loop EID) EID { 
    return F_Generate_g_statement_Construct(Language.To_Construct(OBJ(self)),
      ToClass(OBJ(s)),
      ToString(OBJ(v)),
      ToBoolean(OBJ(err)),
      ANY(loop) )} 
  
// A if is easy to compile. We check if the logical compiler can be used
// we now assume that the test retuns a boolean !
// note that in GO the "} else " pattern is tricky
/* {1} The go function for: g_statement(self:If,s:class,v:string,err:boolean,loop:any) [status=1] */
func F_Generate_g_statement_If (self *Language.If ,s *ClaireClass ,v *ClaireString ,err *ClaireBoolean ,loop *ClaireAny ) EID { 
    var Result EID 
    { var try_ask *ClaireBoolean  
      var try_1 EID 
      /*g_try(v2:"try_1",loop:false) */
      try_1 = Optimize.F_Compile_g_throw_any(self.Test)
      /* ERROR PROTECTION INSERTED (try_ask-Result) */
      if ErrorIn(try_1) {Result = try_1
      } else {
      try_ask = ToBoolean(OBJ(try_1))
      /*g_try(v2:"Result",loop:true) */
      var g0176I *ClaireBoolean  
      var try_2 EID 
      /*g_try(v2:"try_2",loop:false) */
      try_2 = F_Generate_g_clean_any(self.Test)
      /* ERROR PROTECTION INSERTED (g0176I-Result) */
      if ErrorIn(try_2) {Result = try_2
      } else {
      g0176I = ToBoolean(OBJ(try_2))
      if (g0176I == CTRUE) { 
        /*g_try(v2:"Result",loop:true) */
        PRINC("if ")
        /*g_try(v2:"Result",loop:true) */
        Result = Core.F_CALL(Optimize.C_Compile_bool_exp,ARGS(self.Test.ToEID(),EID{CTRUE.Id(),0}))
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC(" ")
        Result = EVOID
        }
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        F_Generate_new_block_string(MakeString("If"))
        Result = EVOID
        }
        } else {
        { var v2 *ClaireString  
          var try_3 EID 
          /*g_try(v2:"try_3",loop:false) */
          try_3 = F_Generate_c_string_go_producer2(ToGenerateGoProducer(Optimize.C_PRODUCER.Value),F_append_symbol(Core.F_gensym_void(),MakeString("I").Id()))
          /* ERROR PROTECTION INSERTED (v2-Result) */
          if ErrorIn(try_3) {Result = try_3
          } else {
          v2 = ToString(OBJ(try_3))
          F_Generate_var_declaration_string(v2,C_boolean,1)
          /*g_try(v2:"Result",loop:true) */
          if (try_ask == CTRUE) { 
            Result = Core.F_CALL(C_Generate_g_try,ARGS(self.Test.ToEID(),
              EID{(v2).Id(),0},
              EID{C_boolean.Id(),0},
              EID{(v).Id(),0},
              EID{CFALSE.Id(),0}))
            } else {
            Result = F_Generate_statement_any(self.Test,C_boolean,v2,loop)
            } 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC("if (")
          F_princ_string(v2)
          PRINC(" == CTRUE) ")
          F_Generate_new_block_string(MakeString("If"))
          PRINC("")
          Result = EVOID
          }
          }
          } 
        } 
      }
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      /*g_try(v2:"Result",loop:true) */
      Result = F_Generate_statement_any(self.Arg,s,v,loop)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      /*g_try(v2:"Result",loop:true) */
      if ((Equal(self.Other,CNIL.Id()) == CTRUE) || 
          ((self.Other == CFALSE.Id()) && 
              (s.Id() == C_void.Id()))) { 
        F_Generate_close_block_string(MakeString("If"))
        Result = EVOID
        } else {
        var g0177I *ClaireBoolean  
        var try_4 EID 
        /*g_try(v2:"try_4",loop:false) */
        { 
          var v_and4 *ClaireBoolean  
          
          v_and4 = self.Other.Isa.IsIn(Language.C_If)
          if (v_and4 == CFALSE) {try_4 = EID{CFALSE.Id(),0}
          } else { 
            var try_5 EID 
            /*g_try(v2:"try_5",loop:false) */
            try_5 = F_Generate_g_func_any(ANY(Core.F_CALL(Language.C_iClaire_test,ARGS(self.Other.ToEID()))))
            /* ERROR PROTECTION INSERTED (v_and4-try_4) */
            if ErrorIn(try_5) {try_4 = try_5
            } else {
            v_and4 = ToBoolean(OBJ(try_5))
            if (v_and4 == CFALSE) {try_4 = EID{CFALSE.Id(),0}
            } else { 
              var try_6 EID 
              /*g_try(v2:"try_6",loop:false) */
              { var arg_7 *ClaireBoolean  
                _ = arg_7
                var try_8 EID 
                /*g_try(v2:"try_8",loop:false) */
                try_8 = Optimize.F_Compile_g_throw_any(ANY(Core.F_CALL(Language.C_iClaire_test,ARGS(self.Other.ToEID()))))
                /* ERROR PROTECTION INSERTED (arg_7-try_6) */
                if ErrorIn(try_8) {try_6 = try_8
                } else {
                arg_7 = ToBoolean(OBJ(try_8))
                try_6 = EID{arg_7.Not.Id(),0}
                }
                } 
              /* ERROR PROTECTION INSERTED (v_and4-try_4) */
              if ErrorIn(try_6) {try_4 = try_6
              } else {
              v_and4 = ToBoolean(OBJ(try_6))
              if (v_and4 == CFALSE) {try_4 = EID{CFALSE.Id(),0}
              } else { 
                try_4 = EID{CTRUE.Id(),0}} 
              } 
            } 
          }}
          } 
        /* ERROR PROTECTION INSERTED (g0177I-Result) */
        if ErrorIn(try_4) {Result = try_4
        } else {
        g0177I = ToBoolean(OBJ(try_4))
        if (g0177I == CTRUE) { 
          F_Generate_finish_block_string(MakeString("If"))
          PRINC(" else ")
          /*g_try(v2:"Result",loop:true) */
          Result = Core.F_CALL(C_Generate_g_statement,ARGS(self.Other.ToEID(),
            EID{s.Id(),0},
            EID{(v).Id(),0},
            EID{CFALSE.Id(),0},
            loop.ToEID()))
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC("")
          Result = EVOID
          }
          } else {
          PRINC("} else {")
          F_Generate_breakline_void()
          /*g_try(v2:"Result",loop:true) */
          Result = F_Generate_statement_any(self.Other,s,v,loop)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          F_Generate_close_block_string(MakeString("If"))
          PRINC("")
          Result = EVOID
          }
          } 
        }
        } 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      if (try_ask == CTRUE) { 
        F_Generate_close_try_integer(1)
        Result = EVOID
        } else {
        Result = EID{CFALSE.Id(),0}
        } 
      }}}
      }
      } 
    return Result} 
  
// The EID go function for: g_statement @ If (throw: true) 
func E_Generate_g_statement_If (self EID,s EID,v EID,err EID,loop EID) EID { 
    return F_Generate_g_statement_If(Language.To_If(OBJ(self)),
      ToClass(OBJ(s)),
      ToString(OBJ(v)),
      ToBoolean(OBJ(err)),
      ANY(loop) )} 
  
// --------------- logical combinations and/or -------------------------------
// note: we cannot use unfolding because the order of evaluation is important !
// AND is compiled with IF: as soon as an argument is false, the result is false.
/* {1} The go function for: g_statement(self:And,s:class,v:string,err:boolean,loop:any) [status=1] */
func F_Generate_g_statement_And (self *Language.And ,s *ClaireClass ,v *ClaireString ,err *ClaireBoolean ,loop *ClaireAny ) EID { 
    var Result EID 
    { var v2 *ClaireString   = F_Generate_check_var_string(MakeString("v_and"))
      { var count_try int  = 0
        _ = count_try
        F_Generate_new_block_string(MakeString("and"))
        F_Generate_var_declaration_string(v2,C_boolean,1)
        F_Generate_breakline_void()
        /*g_try(v2:"Result",loop:true) */
        { 
          var x *ClaireAny  
          _ = x
          Result= EID{CFALSE.Id(),0}
          var x_support *ClaireList  
          x_support = self.Args
          x_len := x_support.Length()
          for i_it := 0; i_it < x_len; i_it++ { 
            x = x_support.At(i_it)
            var loop_1 EID 
            _ = loop_1
            /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
            { var try_ask *ClaireBoolean  
              _ = try_ask
              var try_2 EID 
              /*g_try(v2:"try_2",loop:false) */
              try_2 = Optimize.F_Compile_g_throw_any(x)
              /* ERROR PROTECTION INSERTED (try_ask-loop_1) */
              if ErrorIn(try_2) {loop_1 = try_2
              } else {
              try_ask = ToBoolean(OBJ(try_2))
              /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
              if (try_ask == CTRUE) { 
                /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
                loop_1 = Core.F_CALL(C_Generate_g_try,ARGS(x.ToEID(),
                  EID{(v2).Id(),0},
                  EID{C_boolean.Id(),0},
                  EID{(v).Id(),0},
                  EID{CFALSE.Id(),0}))
                /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
                if ErrorIn(loop_1) {Result = loop_1
                break
                } else {
                count_try = (count_try+1)
                loop_1 = EID{C__INT,IVAL(count_try)}
                }
                } else {
                loop_1 = F_Generate_statement_any(x,C_boolean,v2,loop)
                } 
              /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
              if ErrorIn(loop_1) {Result = loop_1
              break
              } else {
              /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
              PRINC("if (")
              F_c_princ_string(v2)
              PRINC(" == CFALSE) {")
              /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
              if (s.Id() != C_void.Id()) { 
                F_c_princ_string(v)
                PRINC(" = ")
                /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
                loop_1 = F_Generate_cast_prefix_class(C_boolean,s)
                /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
                if ErrorIn(loop_1) {Result = loop_1
                break
                } else {
                PRINC("CFALSE")
                F_Generate_cast_post_class(C_boolean,s)
                PRINC("")
                loop_1 = EVOID
                }
                } else {
                loop_1 = EID{CFALSE.Id(),0}
                } 
              /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
              if ErrorIn(loop_1) {Result = loop_1
              break
              } else {
              F_Generate_breakline_void()
              PRINC("} else ")
              F_Generate_new_block_string(MakeString("arg"))
              PRINC("")
              loop_1 = EVOID
              }
              /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
              if ErrorIn(loop_1) {Result = loop_1
              break
              } else {
              }}
              }
              } 
            /* ERROR PROTECTION INSERTED (loop_1-Result) */
            if ErrorIn(loop_1) {Result = loop_1
            break
            } else {
            }
            } 
          } 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        /*g_try(v2:"Result",loop:true) */
        if (s.Id() != C_void.Id()) { 
          F_c_princ_string(v)
          PRINC(" = ")
          /*g_try(v2:"Result",loop:true) */
          Result = F_Generate_cast_prefix_class(C_boolean,s)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC("CTRUE")
          F_Generate_cast_post_class(C_boolean,s)
          PRINC("")
          Result = EVOID
          }
          } else {
          Result = EID{CFALSE.Id(),0}
          } 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        { 
          var x *ClaireAny  
          _ = x
          var x_support *ClaireList  
          x_support = self.Args
          x_len := x_support.Length()
          for i_it := 0; i_it < x_len; i_it++ { 
            x = x_support.At(i_it)
            F_Generate_close_block_string(MakeString("arg"))
            } 
          } 
        F_Generate_close_try_integer(count_try)
        F_Generate_close_block_string(MakeString("and"))
        Result = EVOID
        }}
        } 
      } 
    return Result} 
  
// The EID go function for: g_statement @ And (throw: true) 
func E_Generate_g_statement_And (self EID,s EID,v EID,err EID,loop EID) EID { 
    return F_Generate_g_statement_And(Language.To_And(OBJ(self)),
      ToClass(OBJ(s)),
      ToString(OBJ(v)),
      ToBoolean(OBJ(err)),
      ANY(loop) )} 
  
// same thing for OR
/* {1} The go function for: g_statement(self:Or,s:class,v:string,err:boolean,loop:any) [status=1] */
func F_Generate_g_statement_Or (self *Language.Or ,s *ClaireClass ,v *ClaireString ,err *ClaireBoolean ,loop *ClaireAny ) EID { 
    var Result EID 
    { var v2 *ClaireString   = F_Generate_check_var_string(MakeString("v_or"))
      { var count_try int  = 0
        _ = count_try
        F_Generate_new_block_string(MakeString("or"))
        /*g_try(v2:"Result",loop:true) */
        PRINC("/* Or stat: v=")
        /*g_try(v2:"Result",loop:true) */
        Result = Core.F_print_any((v).Id())
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC(", loop=")
        /*g_try(v2:"Result",loop:true) */
        Result = Core.F_CALL(C_print,ARGS(loop.ToEID()))
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC(" */")
        F_Generate_breakline_void()
        PRINC("")
        Result = EVOID
        }}
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        F_Generate_var_declaration_string(v2,C_boolean,1)
        F_Generate_breakline_void()
        /*g_try(v2:"Result",loop:true) */
        { 
          var x *ClaireAny  
          _ = x
          Result= EID{CFALSE.Id(),0}
          var x_support *ClaireList  
          x_support = self.Args
          x_len := x_support.Length()
          for i_it := 0; i_it < x_len; i_it++ { 
            x = x_support.At(i_it)
            var loop_1 EID 
            _ = loop_1
            /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
            { var try_ask *ClaireBoolean  
              var try_2 EID 
              /*g_try(v2:"try_2",loop:false) */
              try_2 = Optimize.F_Compile_g_throw_any(x)
              /* ERROR PROTECTION INSERTED (try_ask-loop_1) */
              if ErrorIn(try_2) {loop_1 = try_2
              } else {
              try_ask = ToBoolean(OBJ(try_2))
              /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
              PRINC("/* Or stat: try ")
              /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
              loop_1 = Core.F_CALL(C_print,ARGS(x.ToEID()))
              /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
              if ErrorIn(loop_1) {Result = loop_1
              break
              } else {
              PRINC(" with try:")
              /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
              loop_1 = Core.F_print_any(try_ask.Id())
              /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
              if ErrorIn(loop_1) {Result = loop_1
              break
              } else {
              PRINC(", v=")
              /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
              loop_1 = Core.F_print_any((v).Id())
              /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
              if ErrorIn(loop_1) {Result = loop_1
              break
              } else {
              PRINC(", loop=")
              /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
              loop_1 = Core.F_CALL(C_print,ARGS(loop.ToEID()))
              /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
              if ErrorIn(loop_1) {Result = loop_1
              break
              } else {
              PRINC(" */")
              F_Generate_breakline_void()
              PRINC("")
              loop_1 = EVOID
              }}}}
              /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
              if ErrorIn(loop_1) {Result = loop_1
              break
              } else {
              /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
              if (try_ask == CTRUE) { 
                /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
                loop_1 = Core.F_CALL(C_Generate_g_try,ARGS(x.ToEID(),
                  EID{(v2).Id(),0},
                  EID{C_boolean.Id(),0},
                  EID{(v).Id(),0},
                  loop.ToEID()))
                /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
                if ErrorIn(loop_1) {Result = loop_1
                break
                } else {
                count_try = (count_try+1)
                loop_1 = EID{C__INT,IVAL(count_try)}
                }
                } else {
                loop_1 = F_Generate_statement_any(x,C_boolean,v2,loop)
                } 
              /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
              if ErrorIn(loop_1) {Result = loop_1
              break
              } else {
              /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
              PRINC("if (")
              F_c_princ_string(v2)
              PRINC(" == CTRUE) {")
              /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
              if (s.Id() != C_void.Id()) { 
                F_c_princ_string(v)
                PRINC(" = ")
                /*g_try(v2:"loop_1",loop:tuple("Result", EID)) */
                loop_1 = F_Generate_cast_prefix_class(C_boolean,s)
                /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
                if ErrorIn(loop_1) {Result = loop_1
                break
                } else {
                PRINC("CTRUE")
                F_Generate_cast_post_class(C_boolean,s)
                PRINC("")
                loop_1 = EVOID
                }
                } else {
                loop_1 = EID{CFALSE.Id(),0}
                } 
              /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
              if ErrorIn(loop_1) {Result = loop_1
              break
              } else {
              F_Generate_breakline_void()
              PRINC("} else ")
              F_Generate_new_block_string(MakeString("or"))
              PRINC("")
              loop_1 = EVOID
              }
              /* ERROR PROTECTION INSERTED (loop_1-loop_1) */
              if ErrorIn(loop_1) {Result = loop_1
              break
              } else {
              }}}
              }
              } 
            /* ERROR PROTECTION INSERTED (loop_1-Result) */
            if ErrorIn(loop_1) {Result = loop_1
            break
            } else {
            }
            } 
          } 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        /*g_try(v2:"Result",loop:true) */
        if (s.Id() != C_void.Id()) { 
          F_c_princ_string(v)
          PRINC(" = ")
          /*g_try(v2:"Result",loop:true) */
          Result = F_Generate_cast_prefix_class(C_boolean,s)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC("CFALSE")
          F_Generate_cast_post_class(C_boolean,s)
          PRINC("")
          Result = EVOID
          }
          } else {
          Result = EID{CFALSE.Id(),0}
          } 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        { 
          var x *ClaireAny  
          _ = x
          var x_support *ClaireList  
          x_support = self.Args
          x_len := x_support.Length()
          for i_it := 0; i_it < x_len; i_it++ { 
            x = x_support.At(i_it)
            F_Generate_close_block_string(MakeString("org"))
            } 
          } 
        F_Generate_close_try_integer(count_try)
        F_Generate_close_block_string(MakeString("or"))
        Result = EVOID
        }}}
        } 
      } 
    return Result} 
  
// The EID go function for: g_statement @ Or (throw: true) 
func E_Generate_g_statement_Or (self EID,s EID,v EID,err EID,loop EID) EID { 
    return F_Generate_g_statement_Or(Language.To_Or(OBJ(self)),
      ToClass(OBJ(s)),
      ToString(OBJ(v)),
      ToBoolean(OBJ(err)),
      ANY(loop) )} 
  
// Here this is the simple assignment, with a true variable
// note that the last line (assigning the value to result is only OK if no error)
/* {1} The go function for: g_statement(self:Assign,s:class,v:string,err:boolean,loop:any) [status=1] */
func F_Generate_g_statement_Assign (self *Language.Assign ,s *ClaireClass ,v *ClaireString ,err *ClaireBoolean ,loop *ClaireAny ) EID { 
    var Result EID 
    { var _Zvar *ClaireAny   = self.ClaireVar
      { var x *ClaireAny   = self.Arg
        { var v2 *ClaireString  
          var try_1 EID 
          /*g_try(v2:"try_1",loop:false) */
          try_1 = Core.F_CALL(C_Generate_c_string,ARGS(EID{Optimize.C_PRODUCER.Value,0},_Zvar.ToEID()))
          /* ERROR PROTECTION INSERTED (v2-Result) */
          if ErrorIn(try_1) {Result = try_1
          } else {
          v2 = ToString(OBJ(try_1))
          { var _Zrange *ClaireClass   = ToTypeExpression(OBJ(Core.F_CALL(C_range,ARGS(_Zvar.ToEID())))).Class_I()
            { var try_ask *ClaireBoolean  
              _ = try_ask
              var try_2 EID 
              /*g_try(v2:"try_2",loop:false) */
              try_2 = Optimize.F_Compile_g_throw_any(x)
              /* ERROR PROTECTION INSERTED (try_ask-Result) */
              if ErrorIn(try_2) {Result = try_2
              } else {
              try_ask = ToBoolean(OBJ(try_2))
              if (try_ask == CTRUE) { 
                /*g_try(v2:"Result",loop:true) */
                Result = Core.F_CALL(C_Generate_g_try,ARGS(x.ToEID(),
                  EID{(v2).Id(),0},
                  EID{_Zrange.Id(),0},
                  EID{(v).Id(),0},
                  EID{CFALSE.Id(),0}))
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                /*g_try(v2:"Result",loop:true) */
                if (s.Id() != C_void.Id()) { 
                  F_c_princ_string(v)
                  PRINC(" = ")
                  /*g_try(v2:"Result",loop:true) */
                  Result = F_Generate_use_variable_string(v2,s,_Zrange)
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  F_Generate_breakline_void()
                  PRINC("")
                  Result = EVOID
                  }
                  } else {
                  Result = EID{CFALSE.Id(),0}
                  } 
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                F_Generate_close_try_integer(1)
                Result = EVOID
                }}
                } else {
                /*g_try(v2:"Result",loop:true) */
                Result = F_Generate_statement_any(x,_Zrange,v2,loop)
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                if (s.Id() != C_void.Id()) { 
                  F_c_princ_string(v)
                  PRINC(" = ")
                  /*g_try(v2:"Result",loop:true) */
                  Result = F_Generate_use_variable_string(v2,s,_Zrange)
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  F_Generate_breakline_void()
                  PRINC("")
                  Result = EVOID
                  }
                  } else {
                  Result = EID{CFALSE.Id(),0}
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
  
// The EID go function for: g_statement @ Assign (throw: true) 
func E_Generate_g_statement_Assign (self EID,s EID,v EID,err EID,loop EID) EID { 
    return F_Generate_g_statement_Assign(Language.To_Assign(OBJ(self)),
      ToClass(OBJ(s)),
      ToString(OBJ(v)),
      ToBoolean(OBJ(err)),
      ANY(loop) )} 
  
// This is the global variable assignment - global variables exist in go so this is pretty simple
// note that the tricky part is the store management
/* {1} The go function for: g_statement(self:Gassign,s:class,v:string,err:boolean,loop:any) [status=1] */
func F_Generate_g_statement_Gassign (self *Language.Gassign ,s *ClaireClass ,v *ClaireString ,err *ClaireBoolean ,loop *ClaireAny ) EID { 
    var Result EID 
    { var _Zvar *Core.GlobalVariable   = self.ClaireVar
      { var x *ClaireAny   = self.Arg
        var g0178I *ClaireBoolean  
        var try_1 EID 
        /*g_try(v2:"try_1",loop:false) */
        { 
          var v_and4 *ClaireBoolean  
          
          var try_2 EID 
          /*g_try(v2:"try_2",loop:false) */
          try_2 = F_Generate_g_func_any(x)
          /* ERROR PROTECTION INSERTED (v_and4-try_1) */
          if ErrorIn(try_2) {try_1 = try_2
          } else {
          v_and4 = ToBoolean(OBJ(try_2))
          if (v_and4 == CFALSE) {try_1 = EID{CFALSE.Id(),0}
          } else { 
            v_and4 = Equal(s.Id(),C_void.Id())
            if (v_and4 == CFALSE) {try_1 = EID{CFALSE.Id(),0}
            } else { 
              v_and4 = _Zvar.Store_ask.Not
              if (v_and4 == CFALSE) {try_1 = EID{CFALSE.Id(),0}
              } else { 
                try_1 = EID{CTRUE.Id(),0}} 
              } 
            } 
          }
          } 
        /* ERROR PROTECTION INSERTED (g0178I-Result) */
        if ErrorIn(try_1) {Result = try_1
        } else {
        g0178I = ToBoolean(OBJ(try_1))
        if (g0178I == CTRUE) { 
          ToGenerateGoProducer(Optimize.C_PRODUCER.Value).GlobalVar(_Zvar)
          PRINC(" = ")
          /*g_try(v2:"Result",loop:true) */
          Result = Core.F_CALL(C_Generate_g_expression,ARGS(x.ToEID(),EID{C_any.Id(),0}))
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          F_Generate_breakline_void()
          PRINC("")
          Result = EVOID
          }
          } else {
          { var v2 *ClaireString   = F_Generate_genvar_string(MakeString("v_gassign"))
            { var try_ask *ClaireBoolean  
              var try_3 EID 
              /*g_try(v2:"try_3",loop:false) */
              try_3 = Optimize.F_Compile_g_throw_any(x)
              /* ERROR PROTECTION INSERTED (try_ask-Result) */
              if ErrorIn(try_3) {Result = try_3
              } else {
              try_ask = ToBoolean(OBJ(try_3))
              if ((try_ask != CTRUE) && 
                  (s.Id() == C_any.Id())) { 
                v2 = v
                } else {
                F_Generate_var_declaration_string(v2,C_any,1)
                } 
              /*g_try(v2:"Result",loop:true) */
              if (try_ask == CTRUE) { 
                Result = Core.F_CALL(C_Generate_g_try,ARGS(x.ToEID(),
                  EID{(v2).Id(),0},
                  EID{C_any.Id(),0},
                  EID{(v).Id(),0},
                  EID{CFALSE.Id(),0}))
                } else {
                Result = F_Generate_statement_any(x,C_any,v2,loop)
                } 
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              if (self.ClaireVar.Store_ask == CTRUE) { 
                F_Generate_thing_ident_thing(ToThing(_Zvar.Id()))
                PRINC(".StoreObj(3,")
                F_c_princ_string(v2)
                PRINC(",CTRUE)")
                F_Generate_breakline_void()
                PRINC("")
                } else {
                ToGenerateGoProducer(Optimize.C_PRODUCER.Value).GlobalVar(_Zvar)
                PRINC(" = ")
                F_c_princ_string(v2)
                F_Generate_breakline_void()
                PRINC("")
                } 
              /*g_try(v2:"Result",loop:true) */
              if ((s.Id() != C_void.Id()) && 
                  (v.Value != v2.Value)) { 
                F_c_princ_string(v)
                PRINC(" = ")
                /*g_try(v2:"Result",loop:true) */
                Result = F_Generate_use_variable_string(v2,s,C_any)
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                F_Generate_breakline_void()
                PRINC("")
                Result = EVOID
                }
                } else {
                Result = EID{CFALSE.Id(),0}
                } 
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              if (try_ask == CTRUE) { 
                F_Generate_close_block_void()
                Result = EVOID
                } else {
                Result = EID{CFALSE.Id(),0}
                } 
              }}
              }
              } 
            } 
          } 
        }
        } 
      } 
    return Result} 
  
// The EID go function for: g_statement @ Gassign (throw: true) 
func E_Generate_g_statement_Gassign (self EID,s EID,v EID,err EID,loop EID) EID { 
    return F_Generate_g_statement_Gassign(Language.To_Gassign(OBJ(self)),
      ToClass(OBJ(s)),
      ToString(OBJ(v)),
      ToBoolean(OBJ(err)),
      ANY(loop) )} 
  
//**********************************************************************
//*          Part 3: iteration                                         *
//**********************************************************************
// we know to iterate sets or lists in Go
// the optimizer should give use something that is properly typed
/* {1} The go function for: bag_class(self:any) [status=1] */
func F_Generate_bag_class_any (self *ClaireAny ) EID { 
    var Result EID 
    { var s *ClaireType  
      var try_1 EID 
      /*g_try(v2:"try_1",loop:false) */
      try_1 = Core.F_CALL(Optimize.C_c_type,ARGS(self.ToEID()))
      /* ERROR PROTECTION INSERTED (s-Result) */
      if ErrorIn(try_1) {Result = try_1
      } else {
      s = ToType(OBJ(try_1))
      if ((s.Included(ToType(C_list.Id())) == CTRUE) || 
          ((s.Included(ToType(C_tuple.Id())) == CTRUE) || 
            (s.Included(ToType(C_array.Id())) == CTRUE))) { 
        Result = EID{C_list.Id(),0}
        }  else if (s.Included(ToType(C_set.Id())) == CTRUE) { 
        Result = EID{C_set.Id(),0}
        } else {
        Result = ToException(Core.C_general_error.Make(MakeString("bag_class(~S) returns ~S: cannot use in for").Id(),MakeConstantList(self,s.Id()).Id())).Close()
        } 
      }
      } 
    return Result} 
  
// The EID go function for: bag_class @ any (throw: true) 
func E_Generate_bag_class_any (self EID) EID { 
    return F_Generate_bag_class_any(ANY(self) )} 
  
// generates the iteration code for a "for x in S ..." expression , once
// all optimization based on code substitution have been performed.
// very nice in go, except that we have to handle error
// if g_member(%set) is native (anything but any) we use the native go form
/* {1} The go function for: g_statement(self:For,s:class,v:string,err:boolean,loop:any) [status=1] */
func F_Generate_g_statement_For (self *Language.For ,s *ClaireClass ,v *ClaireString ,err *ClaireBoolean ,loop *ClaireAny ) EID { 
    var Result EID 
    { var v2 *ClaireString  
      var try_1 EID 
      /*g_try(v2:"try_1",loop:false) */
      try_1 = F_Generate_c_string_go_producer1(ToGenerateGoProducer(Optimize.C_PRODUCER.Value),self.ClaireVar)
      /* ERROR PROTECTION INSERTED (v2-Result) */
      if ErrorIn(try_1) {Result = try_1
      } else {
      v2 = ToString(OBJ(try_1))
      { var count_try int  = 0
        _ = count_try
        { var v2_range *ClaireClass   = self.ClaireVar.Range.Class_I()
          { var v3 *ClaireString   = F_append_string(v2,MakeString("_support"))
            _ = v3
            { var v4 *ClaireString   = F_append_string(v2,MakeString("_iter"))
              { var _Zset *ClaireAny   = self.SetArg
                { var sbag *ClaireClass  
                  _ = sbag
                  var try_2 EID 
                  /*g_try(v2:"try_2",loop:false) */
                  try_2 = F_Generate_bag_class_any(_Zset)
                  /* ERROR PROTECTION INSERTED (sbag-Result) */
                  if ErrorIn(try_2) {Result = try_2
                  } else {
                  sbag = ToClass(OBJ(try_2))
                  { var smember *ClaireClass  
                    var try_3 EID 
                    /*g_try(v2:"try_3",loop:false) */
                    try_3 = F_Generate_g_member_any(_Zset)
                    /* ERROR PROTECTION INSERTED (smember-Result) */
                    if ErrorIn(try_3) {Result = try_3
                    } else {
                    smember = ToClass(OBJ(try_3))
                    { var _Zdirect *ClaireBoolean   = MakeBoolean((v2_range.Id() == C_any.Id()) || (smember.Id() == C_integer.Id()) || (smember.Id() == C_float.Id()))
                      F_Generate_new_block_string(MakeString("For"))
                      F_Generate_var_declaration_string(v2,v2_range,2)
                      if (_Zdirect != CTRUE) { 
                        F_Generate_var_declaration_string(v4,C_any,1)
                        } else {
                        v4 = v2
                        } 
                      /*g_try(v2:"Result",loop:true) */
                      if (s.Id() != C_void.Id()) { 
                        F_c_princ_string(v)
                        PRINC("= ")
                        /*g_try(v2:"Result",loop:true) */
                        Result = F_Generate_cast_prefix_class(C_boolean,s)
                        /* ERROR PROTECTION INSERTED (Result-Result) */
                        if !ErrorIn(Result) {
                        PRINC("CFALSE")
                        F_Generate_cast_post_class(C_boolean,s)
                        F_Generate_breakline_void()
                        PRINC("")
                        Result = EVOID
                        }
                        } else {
                        Result = EID{CFALSE.Id(),0}
                        } 
                      /* ERROR PROTECTION INSERTED (Result-Result) */
                      if !ErrorIn(Result) {
                      var try_4 EID 
                      /*g_try(v2:"try_4",loop:true) */
                      { var arg_5 int 
                        _ = arg_5
                        var try_6 EID 
                        /*g_try(v2:"try_6",loop:false) */
                        try_6 = F_Generate_iteration_statement_For(self,
                          _Zset,
                          sbag,
                          smember,
                          v,
                          v3,
                          v4)
                        /* ERROR PROTECTION INSERTED (arg_5-try_4) */
                        if ErrorIn(try_6) {try_4 = try_6
                        } else {
                        arg_5 = INT(try_6)
                        try_4 = EID{C__INT,IVAL((count_try+arg_5))}
                        }
                        } 
                      /* ERROR PROTECTION INSERTED (count_try-Result) */
                      if ErrorIn(try_4) {Result = try_4
                      } else {
                      count_try = INT(try_4)
                      Result = EID{C__INT,IVAL(count_try)}
                      /*g_try(v2:"Result",loop:true) */
                      if (_Zdirect != CTRUE) { 
                        F_princ_string(v2)
                        PRINC(" = ")
                        /*g_try(v2:"Result",loop:true) */
                        Result = F_Generate_cast_prefix_class(C_any,v2_range)
                        /* ERROR PROTECTION INSERTED (Result-Result) */
                        if !ErrorIn(Result) {
                        F_princ_string(v4)
                        F_Generate_cast_post_class(C_any,v2_range)
                        F_Generate_breakline_void()
                        PRINC("")
                        Result = EVOID
                        }
                        } else {
                        Result = EID{CFALSE.Id(),0}
                        } 
                      /* ERROR PROTECTION INSERTED (Result-Result) */
                      if !ErrorIn(Result) {
                      /*g_try(v2:"Result",loop:true) */
                      var g0179I *ClaireBoolean  
                      var try_7 EID 
                      /*g_try(v2:"try_7",loop:false) */
                      try_7 = Optimize.F_Compile_g_throw_any(self.Arg)
                      /* ERROR PROTECTION INSERTED (g0179I-Result) */
                      if ErrorIn(try_7) {Result = try_7
                      } else {
                      g0179I = ToBoolean(OBJ(try_7))
                      if (g0179I == CTRUE) { 
                        /*g_try(v2:"Result",loop:true) */
                        Result = F_Generate_g_try_void_any(self.Arg,v,MakeTuple((v).Id(),s.Id()).Id())
                        /* ERROR PROTECTION INSERTED (Result-Result) */
                        if !ErrorIn(Result) {
                        count_try = (count_try+1)
                        Result = EID{C__INT,IVAL(count_try)}
                        }
                        } else {
                        Result = F_Generate_statement_any(self.Arg,C_void,v,MakeTuple((v).Id(),s.Id()).Id())
                        } 
                      }
                      /* ERROR PROTECTION INSERTED (Result-Result) */
                      if !ErrorIn(Result) {
                      F_Generate_close_try_integer(count_try)
                      F_Generate_close_block_string(MakeString("loop"))
                      F_Generate_close_block_string(MakeString("For"))
                      Result = EVOID
                      }}}}
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
  
// The EID go function for: g_statement @ For (throw: true) 
func E_Generate_g_statement_For (self EID,s EID,v EID,err EID,loop EID) EID { 
    return F_Generate_g_statement_For(Language.To_For(OBJ(self)),
      ToClass(OBJ(s)),
      ToString(OBJ(v)),
      ToBoolean(OBJ(err)),
      ANY(loop) )} 
  
// for statement
// iteration_statement produces the bulk of the iteration code
// returns 1 if we use a try/pattern for error protection
/* {1} The go function for: iteration_statement(self:For,%set:any,sbag:class,smember:class,v:string,v3:string,v4:string) [status=1] */
func F_Generate_iteration_statement_For (self *Language.For ,_Zset *ClaireAny ,sbag *ClaireClass ,smember *ClaireClass ,v *ClaireString ,v3 *ClaireString ,v4 *ClaireString ) EID { 
    var Result EID 
    var g0180I *ClaireBoolean  
    var try_1 EID 
    /*g_try(v2:"try_1",loop:false) */
    { 
      var v_and2 *ClaireBoolean  
      
      var try_2 EID 
      /*g_try(v2:"try_2",loop:false) */
      try_2 = F_Generate_g_clean_any(_Zset)
      /* ERROR PROTECTION INSERTED (v_and2-try_1) */
      if ErrorIn(try_2) {try_1 = try_2
      } else {
      v_and2 = ToBoolean(OBJ(try_2))
      if (v_and2 == CFALSE) {try_1 = EID{CFALSE.Id(),0}
      } else { 
        var try_3 EID 
        /*g_try(v2:"try_3",loop:false) */
        try_3 = Optimize.F_Compile_designated_ask_any(_Zset)
        /* ERROR PROTECTION INSERTED (v_and2-try_1) */
        if ErrorIn(try_3) {try_1 = try_3
        } else {
        v_and2 = ToBoolean(OBJ(try_3))
        if (v_and2 == CFALSE) {try_1 = EID{CFALSE.Id(),0}
        } else { 
          v_and2 = MakeBoolean((smember.Id() != C_any.Id()) && (sbag.Id() == C_list.Id()))
          if (v_and2 == CFALSE) {try_1 = EID{CFALSE.Id(),0}
          } else { 
            try_1 = EID{CTRUE.Id(),0}} 
          } 
        } 
      }}
      } 
    /* ERROR PROTECTION INSERTED (g0180I-Result) */
    if ErrorIn(try_1) {Result = try_1
    } else {
    g0180I = ToBoolean(OBJ(try_1))
    if (g0180I == CTRUE) { 
      /*g_try(v2:"Result",loop:true) */
      PRINC("for _,")
      F_c_princ_string(v4)
      PRINC(" = range(")
      /*g_try(v2:"Result",loop:true) */
      Result = Core.F_CALL(C_Generate_g_expression,ARGS(_Zset.ToEID(),EID{sbag.Id(),0}))
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      F_Generate_cast_Values_class(sbag,smember)
      PRINC(")")
      F_Generate_new_block_string(MakeString("loop"))
      PRINC("")
      Result = EVOID
      }
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = EID{C__INT,IVAL(0)}
      }
      } else {
      { var try_ask *ClaireBoolean  
        var try_4 EID 
        /*g_try(v2:"try_4",loop:false) */
        try_4 = Optimize.F_Compile_g_throw_any(_Zset)
        /* ERROR PROTECTION INSERTED (try_ask-Result) */
        if ErrorIn(try_4) {Result = try_4
        } else {
        try_ask = ToBoolean(OBJ(try_4))
        F_Generate_var_declaration_string(v3,sbag,1)
        /*g_try(v2:"Result",loop:true) */
        if (try_ask == CTRUE) { 
          Result = Core.F_CALL(C_Generate_g_try,ARGS(_Zset.ToEID(),
            EID{(v3).Id(),0},
            EID{sbag.Id(),0},
            EID{(v).Id(),0},
            EID{CFALSE.Id(),0}))
          } else {
          Result = F_Generate_statement_any(_Zset,sbag,v3,CFALSE.Id())
          } 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        /*g_try(v2:"Result",loop:true) */
        if (sbag.Id() == C_set.Id()) { 
          PRINC("for i_it := 0; i_it < ")
          F_c_princ_string(v3)
          PRINC(".Count; i_it++ ")
          F_Generate_new_block_void()
          if ((smember.Id() == C_integer.Id()) || 
              (smember.Id() == C_float.Id())) { 
            F_c_princ_string(v4)
            PRINC(" = ")
            F_c_princ_string(v3)
            F_Generate_cast_Values_class(sbag,smember)
            PRINC("[i_it]")
            F_Generate_breakline_void()
            PRINC("")
            } else {
            F_c_princ_string(v4)
            PRINC(" = ")
            F_c_princ_string(v3)
            PRINC(".At(i_it)")
            F_Generate_breakline_void()
            PRINC("")
            } 
          PRINC("")
          Result = EVOID
          } else {
          var g0181I *ClaireBoolean  
          var try_5 EID 
          /*g_try(v2:"try_5",loop:false) */
          { var arg_6 *ClaireClass  
            _ = arg_6
            var try_7 EID 
            /*g_try(v2:"try_7",loop:false) */
            try_7 = F_Generate_g_member_any(_Zset)
            /* ERROR PROTECTION INSERTED (arg_6-try_5) */
            if ErrorIn(try_7) {try_5 = try_7
            } else {
            arg_6 = ToClass(OBJ(try_7))
            try_5 = EID{Core.F__I_equal_any(arg_6.Id(),C_any.Id()).Id(),0}
            }
            } 
          /* ERROR PROTECTION INSERTED (g0181I-Result) */
          if ErrorIn(try_5) {Result = try_5
          } else {
          g0181I = ToBoolean(OBJ(try_5))
          if (g0181I == CTRUE) { 
            PRINC("for _,")
            F_c_princ_string(v4)
            PRINC(" = range(")
            F_c_princ_string(v3)
            F_Generate_cast_Values_class(sbag,smember)
            PRINC(")")
            F_Generate_new_block_string(MakeString("loop2"))
            PRINC("")
            Result = EVOID
            } else {
            { var v5 *ClaireString  
              var try_8 EID 
              /*g_try(v2:"try_8",loop:false) */
              { var arg_9 *ClaireString  
                _ = arg_9
                var try_10 EID 
                /*g_try(v2:"try_10",loop:false) */
                try_10 = F_Generate_c_string_go_producer1(ToGenerateGoProducer(Optimize.C_PRODUCER.Value),self.ClaireVar)
                /* ERROR PROTECTION INSERTED (arg_9-try_8) */
                if ErrorIn(try_10) {try_8 = try_10
                } else {
                arg_9 = ToString(OBJ(try_10))
                try_8 = EID{F_append_string(arg_9,MakeString("_len")).Id(),0}
                }
                } 
              /* ERROR PROTECTION INSERTED (v5-Result) */
              if ErrorIn(try_8) {Result = try_8
              } else {
              v5 = ToString(OBJ(try_8))
              F_c_princ_string(v5)
              PRINC(" := ")
              F_c_princ_string(v3)
              PRINC(".Length()")
              F_Generate_breakline_void()
              PRINC("")
              PRINC("for i_it := 0; i_it < ")
              F_c_princ_string(v5)
              PRINC("; i_it++ ")
              F_Generate_new_block_void()
              F_c_princ_string(v4)
              PRINC(" = ")
              F_c_princ_string(v3)
              PRINC(".At(i_it)")
              F_Generate_breakline_void()
              PRINC("")
              PRINC("")
              Result = EVOID
              }
              } 
            } 
          }
          } 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        if (try_ask == CTRUE) { 
          Result = EID{C__INT,IVAL(1)}
          } else {
          Result = EID{C__INT,IVAL(0)}
          } 
        }}
        }
        } 
      } 
    }
    return Result} 
  
// The EID go function for: iteration_statement @ For (throw: true) 
func E_Generate_iteration_statement_For (self EID,_Zset EID,sbag EID,smember EID,v EID,v3 EID,v4 EID) EID { 
    return F_Generate_iteration_statement_For(Language.To_For(OBJ(self)),
      ANY(_Zset),
      ToClass(OBJ(sbag)),
      ToClass(OBJ(smember)),
      ToString(OBJ(v)),
      ToString(OBJ(v3)),
      ToString(OBJ(v4)) )} 
  
// here the value is expected to be important, otherwise an error is raised.
// THIS IS ONLY APPLIED TO COLLECT(f(x) | s in S) on lists => Image is delt with
// we currently do not use the native form => use At and Put to work on generic lists
/* {1} The go function for: g_statement(self:Iteration,s:class,v:string,err:boolean,loop:any) [status=1] */
func F_Generate_g_statement_Iteration (self *Language.Iteration ,s *ClaireClass ,v *ClaireString ,err *ClaireBoolean ,loop *ClaireAny ) EID { 
    var Result EID 
    /*g_try(v2:"Result",loop:true) */
    if (s.Id() == C_void.Id()) { 
      Result = ToException(Core.C_general_error.Make(MakeString("[203] you should have used a FOR ere:~S").Id(),MakeConstantList(self.Id()).Id())).Close()
      } else {
      Result = EID{CFALSE.Id(),0}
      } 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    { var v2 *ClaireString  
      var try_1 EID 
      /*g_try(v2:"try_1",loop:false) */
      try_1 = F_Generate_c_string_go_producer1(ToGenerateGoProducer(Optimize.C_PRODUCER.Value),self.ClaireVar)
      /* ERROR PROTECTION INSERTED (v2-Result) */
      if ErrorIn(try_1) {Result = try_1
      } else {
      v2 = ToString(OBJ(try_1))
      { var v2_range *ClaireClass   = self.ClaireVar.Range.Class_I()
        { var vlist *ClaireString   = F_Generate_check_var_string(MakeString("v_list"))
          { var vlocal *ClaireString   = F_Generate_check_var_string(MakeString("v_local"))
            { var bag_type *ClaireClass  
              var try_2 EID 
              /*g_try(v2:"try_2",loop:false) */
              var g0182I *ClaireBoolean  
              var try_3 EID 
              /*g_try(v2:"try_3",loop:false) */
              { var arg_4 *ClaireType  
                _ = arg_4
                var try_5 EID 
                /*g_try(v2:"try_5",loop:false) */
                try_5 = Core.F_CALL(Optimize.C_c_type,ARGS(self.SetArg.ToEID()))
                /* ERROR PROTECTION INSERTED (arg_4-try_3) */
                if ErrorIn(try_5) {try_3 = try_5
                } else {
                arg_4 = ToType(OBJ(try_5))
                try_3 = EID{arg_4.Included(ToType(C_set.Id())).Id(),0}
                }
                } 
              /* ERROR PROTECTION INSERTED (g0182I-try_2) */
              if ErrorIn(try_3) {try_2 = try_3
              } else {
              g0182I = ToBoolean(OBJ(try_3))
              if (g0182I == CTRUE) { 
                try_2 = EID{C_set.Id(),0}
                } else {
                try_2 = EID{C_list.Id(),0}
                } 
              }
              /* ERROR PROTECTION INSERTED (bag_type-Result) */
              if ErrorIn(try_2) {Result = try_2
              } else {
              bag_type = ToClass(OBJ(try_2))
              { var try_count int  = 0
                _ = try_count
                
                F_Generate_new_block_string(MakeString("Iteration"))
                F_Generate_var_declaration_string(vlist,bag_type,1)
                F_Generate_var_declaration_string(v2,v2_range,1)
                F_Generate_var_declaration_string(vlocal,C_any,1)
                /*g_try(v2:"Result",loop:true) */
                var g0183I *ClaireBoolean  
                var try_6 EID 
                /*g_try(v2:"try_6",loop:false) */
                try_6 = Optimize.F_Compile_g_throw_any(self.SetArg)
                /* ERROR PROTECTION INSERTED (g0183I-Result) */
                if ErrorIn(try_6) {Result = try_6
                } else {
                g0183I = ToBoolean(OBJ(try_6))
                if (g0183I == CTRUE) { 
                  try_count = (try_count+1)
                  Result = Core.F_CALL(C_Generate_g_try,ARGS(self.SetArg.ToEID(),
                    EID{(vlist).Id(),0},
                    EID{bag_type.Id(),0},
                    EID{(v).Id(),0},
                    loop.ToEID()))
                  } else {
                  Result = F_Generate_statement_any(self.SetArg,bag_type,vlist,loop)
                  } 
                }
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                /*g_try(v2:"Result",loop:true) */
                F_princ_string(v)
                PRINC(" = ")
                /*g_try(v2:"Result",loop:true) */
                Result = F_Generate_cast_prefix_class(C_list,s)
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                PRINC("CreateList(")
                /*g_try(v2:"Result",loop:true) */
                if (Core.F_get_property(C_of,ToObject(self.Id())) != CNULL) { 
                  { var arg_7 *ClaireAny  
                    _ = arg_7
                    var try_8 EID 
                    /*g_try(v2:"try_8",loop:false) */
                    try_8 = Core.F_CALL(Optimize.C_c_code,ARGS(Core.F_CALL(C_of,ARGS(EID{self.Id(),0})),EID{C_type.Id(),0}))
                    /* ERROR PROTECTION INSERTED (arg_7-Result) */
                    if ErrorIn(try_8) {Result = try_8
                    } else {
                    arg_7 = ANY(try_8)
                    Result = Core.F_CALL(C_Generate_g_expression,ARGS(arg_7.ToEID(),EID{C_type.Id(),0}))
                    }
                    } 
                  } else {
                  PRINC("ToType(CEMPTY.Id())")
                  Result = EVOID
                  } 
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                PRINC(",")
                F_princ_string(vlist)
                PRINC(".Length())")
                F_Generate_cast_post_class(C_list,s)
                F_Generate_breakline_void()
                PRINC("")
                Result = EVOID
                }}
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                PRINC("for CLcount := 0; CLcount < ")
                F_princ_string(vlist)
                PRINC(".")
                F_princ_string(ToString(IfThenElse((bag_type.Id() == C_set.Id()),
                  MakeString("Count").Id(),
                  MakeString("Length()").Id())))
                PRINC("; CLcount++")
                F_Generate_new_block_void()
                PRINC("")
                /*g_try(v2:"Result",loop:true) */
                F_princ_string(v2)
                PRINC(" = ")
                /*g_try(v2:"Result",loop:true) */
                Result = F_Generate_cast_prefix_class(C_any,v2_range)
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                F_princ_string(vlist)
                PRINC(".At(CLcount)")
                F_Generate_cast_post_class(C_any,v2_range)
                F_Generate_breakline_void()
                PRINC("")
                Result = EVOID
                }
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                /*g_try(v2:"Result",loop:true) */
                var g0184I *ClaireBoolean  
                var try_9 EID 
                /*g_try(v2:"try_9",loop:false) */
                try_9 = Optimize.F_Compile_g_throw_any(self.Arg)
                /* ERROR PROTECTION INSERTED (g0184I-Result) */
                if ErrorIn(try_9) {Result = try_9
                } else {
                g0184I = ToBoolean(OBJ(try_9))
                if (g0184I == CTRUE) { 
                  try_count = (try_count+1)
                  Result = Core.F_CALL(C_Generate_g_try,ARGS(self.Arg.ToEID(),
                    EID{(vlocal).Id(),0},
                    EID{C_any.Id(),0},
                    EID{(v).Id(),0},
                    EID{MakeTuple((v).Id(),s.Id()).Id(),0}))
                  } else {
                  Result = F_Generate_statement_any(self.Arg,C_any,vlocal,MakeTuple((v).Id(),s.Id()).Id())
                  } 
                }
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                /*g_try(v2:"Result",loop:true) */
                /*g_try(v2:"Result",loop:true) */
                Result = F_Generate_cast_prefix_class(s,C_list)
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                F_princ_string(v)
                F_Generate_cast_post_class(s,C_list)
                PRINC(".PutAt(CLcount,")
                F_princ_string(vlocal)
                PRINC(")")
                F_Generate_breakline_void()
                PRINC("")
                Result = EVOID
                }
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                F_Generate_close_block_void()
                F_Generate_close_try_integer(try_count)
                F_Generate_close_block_string(MakeString("Iteration"))
                Result = EVOID
                }}}}}
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
  
// The EID go function for: g_statement @ Iteration (throw: true) 
func E_Generate_g_statement_Iteration (self EID,s EID,v EID,err EID,loop EID) EID { 
    return F_Generate_g_statement_Iteration(Language.To_Iteration(OBJ(self)),
      ToClass(OBJ(s)),
      ToString(OBJ(v)),
      ToBoolean(OBJ(err)),
      ANY(loop) )} 
  
// --------------- WHILE   ------------------------------------------
// if it is possible the logical compiler is used to produce a better code
// self.other = true => until(....) was used 
// error is more tricky => we produce a chain with 3 more blocks
/* {1} The go function for: g_statement(self:While,s:class,v:string,err:boolean,loop:any) [status=1] */
func F_Generate_g_statement_While (self *Language.While ,s *ClaireClass ,v *ClaireString ,err *ClaireBoolean ,loop *ClaireAny ) EID { 
    var Result EID 
    { var f_ask *ClaireBoolean  
      var try_1 EID 
      /*g_try(v2:"try_1",loop:false) */
      { 
        var v_and3 *ClaireBoolean  
        
        var try_2 EID 
        /*g_try(v2:"try_2",loop:false) */
        try_2 = F_Generate_g_clean_any(self.Test)
        /* ERROR PROTECTION INSERTED (v_and3-try_1) */
        if ErrorIn(try_2) {try_1 = try_2
        } else {
        v_and3 = ToBoolean(OBJ(try_2))
        if (v_and3 == CFALSE) {try_1 = EID{CFALSE.Id(),0}
        } else { 
          v_and3 = self.Other.Not
          if (v_and3 == CFALSE) {try_1 = EID{CFALSE.Id(),0}
          } else { 
            try_1 = EID{CTRUE.Id(),0}} 
          } 
        }
        } 
      /* ERROR PROTECTION INSERTED (f_ask-Result) */
      if ErrorIn(try_1) {Result = try_1
      } else {
      f_ask = ToBoolean(OBJ(try_1))
      { var try_ask *ClaireBoolean  
        var try_3 EID 
        /*g_try(v2:"try_3",loop:false) */
        try_3 = Optimize.F_Compile_g_throw_any(self.Test)
        /* ERROR PROTECTION INSERTED (try_ask-Result) */
        if ErrorIn(try_3) {Result = try_3
        } else {
        try_ask = ToBoolean(OBJ(try_3))
        { var try2_ask *ClaireBoolean  
          var try_4 EID 
          /*g_try(v2:"try_4",loop:false) */
          try_4 = Optimize.F_Compile_g_throw_any(self.Arg)
          /* ERROR PROTECTION INSERTED (try2_ask-Result) */
          if ErrorIn(try_4) {Result = try_4
          } else {
          try2_ask = ToBoolean(OBJ(try_4))
          { var v2 *ClaireString   = F_Generate_check_var_string(MakeString("v_while"))
            if (f_ask != CTRUE) { 
              F_Generate_var_declaration_string(v2,C_boolean,1)
              } 
            /*g_try(v2:"Result",loop:true) */
            if (s.Id() != C_void.Id()) { 
              F_c_princ_string(v)
              PRINC("= ")
              /*g_try(v2:"Result",loop:true) */
              Result = F_Generate_cast_prefix_class(C_boolean,s)
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              PRINC("CFALSE")
              F_Generate_cast_post_class(C_boolean,s)
              F_Generate_breakline_void()
              PRINC("")
              Result = EVOID
              }
              } else {
              Result = EID{CFALSE.Id(),0}
              } 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            /*g_try(v2:"Result",loop:true) */
            if (f_ask == CTRUE) { 
              PRINC("for ")
              /*g_try(v2:"Result",loop:true) */
              Result = Core.F_CALL(Optimize.C_Compile_bool_exp,ARGS(self.Test.ToEID(),EID{CTRUE.Id(),0}))
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              PRINC(" ")
              Result = EVOID
              }
              } else {
              /*g_try(v2:"Result",loop:true) */
              if (try_ask == CTRUE) { 
                Result = Core.F_CALL(C_Generate_g_try,ARGS(self.Test.ToEID(),
                  EID{(v2).Id(),0},
                  EID{C_boolean.Id(),0},
                  EID{(v).Id(),0},
                  EID{CFALSE.Id(),0}))
                } else {
                { var arg_5 *ClaireAny  
                  _ = arg_5
                  if (self.Other == CTRUE) { 
                    arg_5 = CFALSE.Id()
                    } else {
                    arg_5 = self.Test
                    } 
                  Result = F_Generate_statement_any(arg_5,C_boolean,v2,CFALSE.Id())
                  } 
                } 
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              F_Generate_breakline_void()
              PRINC("for ")
              F_princ_string(v2)
              PRINC(" ")
              if (self.Other == CTRUE) { 
                PRINC("!=")
                } else {
                PRINC("==")
                } 
              PRINC(" CTRUE ")
              Result = EVOID
              }
              } 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            F_Generate_new_block_string(MakeString("while"))
            /*g_try(v2:"Result",loop:true) */
            PRINC("/* While stat, v:")
            /*g_try(v2:"Result",loop:true) */
            Result = Core.F_print_any((v).Id())
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC(" loop:")
            /*g_try(v2:"Result",loop:true) */
            Result = Core.F_CALL(C_print,ARGS(loop.ToEID()))
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC(" */")
            F_Generate_breakline_void()
            PRINC("")
            Result = EVOID
            }}
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            /*g_try(v2:"Result",loop:true) */
            if (try2_ask == CTRUE) { 
              Result = F_Generate_g_try_void_any(self.Arg,v,MakeTuple((v).Id(),s.Id()).Id())
              } else {
              Result = F_Generate_statement_any(self.Arg,C_void,v,MakeTuple((v).Id(),s.Id()).Id())
              } 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            /*g_try(v2:"Result",loop:true) */
            PRINC("/* try?:")
            /*g_try(v2:"Result",loop:true) */
            Result = Core.F_print_any(try_ask.Id())
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC(", v2:")
            /*g_try(v2:"Result",loop:true) */
            Result = Core.F_print_any((v2).Id())
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC(" loop will be:")
            /*g_try(v2:"Result",loop:true) */
            Result = Core.F_print_any(MakeTuple((v).Id(),s.Id()).Id())
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC(" */")
            F_Generate_breakline_void()
            PRINC("")
            Result = EVOID
            }}}
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            /*g_try(v2:"Result",loop:true) */
            if (try_ask == CTRUE) { 
              Result = Core.F_CALL(C_Generate_g_try,ARGS(self.Test.ToEID(),
                EID{(v2).Id(),0},
                EID{C_boolean.Id(),0},
                EID{(v).Id(),0},
                EID{MakeTuple((v).Id(),s.Id()).Id(),0}))
              }  else if (f_ask != CTRUE) { 
              Result = F_Generate_statement_any(self.Test,C_boolean,v2,CFALSE.Id())
              } else {
              Result = EID{CFALSE.Id(),0}
              } 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            F_Generate_close_block_string(MakeString("while"))
            if (try_ask == CTRUE) { 
              F_Generate_close_try_integer(2)
              } 
            if (try2_ask == CTRUE) { 
              F_Generate_close_try_integer(1)
              Result = EVOID
              } else {
              Result = EID{CFALSE.Id(),0}
              } 
            }}}}}}
            } 
          }
          } 
        }
        } 
      }
      } 
    return Result} 
  
// The EID go function for: g_statement @ While (throw: true) 
func E_Generate_g_statement_While (self EID,s EID,v EID,err EID,loop EID) EID { 
    return F_Generate_g_statement_While(Language.To_While(OBJ(self)),
      ToClass(OBJ(s)),
      ToString(OBJ(v)),
      ToBoolean(OBJ(err)),
      ANY(loop) )} 
  
//------------- compiling a return -------------------------------------
// a return inside a loop is compiled with a break, the go variable is provided
// in the loop argument
/* {1} The go function for: g_statement(self:Return,s:class,v:string,err:boolean,loop:any) [status=1] */
func F_Generate_g_statement_Return (self *Language.Return ,s *ClaireClass ,v *ClaireString ,err *ClaireBoolean ,loop *ClaireAny ) EID { 
    var Result EID 
    /*g_try(v2:"Result",loop:true) */
    if (C_tuple.Id() == loop.Isa.Id()) { 
      { var g0185 *ClaireTuple   = ToTuple(loop)
        { var vreturn *ClaireString   = ToString(ToList(g0185.Id()).At(1-1))
          { var sreturn *ClaireClass   = ToClass(ToList(g0185.Id()).At(2-1))
            /*g_try(v2:"Result",loop:true) */
            if (ToGenerateGoProducer(Optimize.C_PRODUCER.Value).Debug_ask == CTRUE) { 
              PRINC(" /*v = ")
              F_princ_string(vreturn)
              PRINC(", s =")
              /*g_try(v2:"Result",loop:true) */
              Result = Core.F_print_any(sreturn.Id())
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              PRINC("*/")
              F_Generate_breakline_void()
              PRINC("")
              Result = EVOID
              }
              } else {
              Result = EID{CFALSE.Id(),0}
              } 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            Result = F_Generate_statement_any(self.Arg,sreturn,vreturn,CFALSE.Id())
            }
            } 
          } 
        } 
      } else {
      Result = EID{CFALSE.Id(),0}
      } 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("break")
    Result = F_Generate_breakline_void().ToEID()
    }
    return Result} 
  
// The EID go function for: g_statement @ Return (throw: true) 
func E_Generate_g_statement_Return (self EID,s EID,v EID,err EID,loop EID) EID { 
    return F_Generate_g_statement_Return(Language.To_Return(OBJ(self)),
      ToClass(OBJ(s)),
      ToString(OBJ(v)),
      ToBoolean(OBJ(err)),
      ANY(loop) )} 
  
//**********************************************************************
//*          Part 4: CLAIRE-specific structures                        *
//**********************************************************************
// ------------- Messages and optimized instructions ------------------------
// this is one example on how to unfold: a Call
// note that if the error is returned it should be passed away
// we also add inline_stat in v4 for special cases
/* {1} The go function for: g_statement(self:Call,s:class,v:string,err:boolean,loop:any) [status=1] */
func F_Generate_g_statement_Call (self *Language.Call ,s *ClaireClass ,v *ClaireString ,err *ClaireBoolean ,loop *ClaireAny ) EID { 
    var Result EID 
    var g0186I *ClaireBoolean  
    var try_1 EID 
    /*g_try(v2:"try_1",loop:false) */
    try_1 = F_Generate_g_clean_any(self.Args.Id())
    /* ERROR PROTECTION INSERTED (g0186I-Result) */
    if ErrorIn(try_1) {Result = try_1
    } else {
    g0186I = ToBoolean(OBJ(try_1))
    if (g0186I == CTRUE) { 
      Result = F_Generate_inline_stat_Call(self,s,v)
      } else {
      { var l *ClaireList   = self.Args
        { var ld *ClaireList  
          var try_2 EID 
          /*g_try(v2:"try_2",loop:false) */
          try_2 = F_Generate_unfold_args_list(l)
          /* ERROR PROTECTION INSERTED (ld-Result) */
          if ErrorIn(try_2) {Result = try_2
          } else {
          ld = ToList(OBJ(try_2))
          if (Equal(ld.Id(),CNIL.Id()) == CTRUE) { 
            Result = F_Generate_error_wrap_any(self.Id(),s,v)
            } else {
            { var arg_3 *Language.Call  
              _ = arg_3
              var try_4 EID 
              /*g_try(v2:"try_4",loop:false) */
              { var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                _CL_obj.Selector = self.Selector
                /*property->property*//*g_try(v2:"try_4",loop:false) */
                { 
                  var va_arg1 *Language.Call  
                  var va_arg2 *ClaireList  
                  va_arg1 = _CL_obj
                  var try_5 EID 
                  /*g_try(v2:"try_5",loop:false) */
                  { 
                    var v_list9 *ClaireList  
                    var z *ClaireAny  
                    var v_local9 *ClaireAny  
                    v_list9 = l
                    try_5 = EID{CreateList(ToType(CEMPTY.Id()),v_list9.Length()).Id(),0}
                    for CLcount := 0; CLcount < v_list9.Length(); CLcount++{ 
                      z = v_list9.At(CLcount)
                      var try_6 EID 
                      /*g_try(v2:"try_6",loop:tuple("try_5", EID)) */
                      try_6 = F_Generate_unfold_arg_list(l,ld,z)
                      /* ERROR PROTECTION INSERTED (v_local9-try_5) */
                      if ErrorIn(try_6) {try_5 = try_6
                      break
                      } else {
                      v_local9 = ANY(try_6)
                      ToList(OBJ(try_5)).PutAt(CLcount,v_local9)
                      } 
                    }
                    } 
                  /* ERROR PROTECTION INSERTED (va_arg2-try_4) */
                  if ErrorIn(try_5) {try_4 = try_5
                  } else {
                  va_arg2 = ToList(OBJ(try_5))
                  va_arg1.Args = va_arg2
                  /*list->list*/try_4 = EID{va_arg2.Id(),0}
                  }
                  } 
                /* ERROR PROTECTION INSERTED (try_4-try_4) */
                if !ErrorIn(try_4) {
                try_4 = EID{_CL_obj.Id(),0}
                }
                } 
              /* ERROR PROTECTION INSERTED (arg_3-Result) */
              if ErrorIn(try_4) {Result = try_4
              } else {
              arg_3 = Language.To_Call(OBJ(try_4))
              Result = F_Generate_unfold_use_list(ld,
                arg_3.Id(),
                s,
                v,
                err,
                loop)
              }
              } 
            } 
          }
          } 
        } 
      } 
    }
    return Result} 
  
// The EID go function for: g_statement @ Call (throw: true) 
func E_Generate_g_statement_Call (self EID,s EID,v EID,err EID,loop EID) EID { 
    return F_Generate_g_statement_Call(Language.To_Call(OBJ(self)),
      ToClass(OBJ(s)),
      ToString(OBJ(v)),
      ToBoolean(OBJ(err)),
      ANY(loop) )} 
  
// this is our special inling that requires an assignment (not allowed as an expression in go)
/* {1} The go function for: inline_stat(self:Call,s:class,v:string) [status=1] */
func F_Generate_inline_stat_Call (self *Language.Call ,s *ClaireClass ,v *ClaireString ) EID { 
    var Result EID 
    if (self.Selector.Id() == Optimize.C_Compile_object_I.Id()) { 
      { var a1 *ClaireAny   = self.Args.At(1-1)
        { var a2 *ClaireAny   = self.Args.At(2-1)
          /*g_try(v2:"Result",loop:true) */
          if ((a2 == C_property.Id()) && 
              (Core.F_owner_any(ANY(Core.F_CALL(C_value,ARGS(a1.ToEID())))).IsIn(C_property) == CTRUE)) { 
            F_Generate_symbol_ident_symbol(ToSymbol(a1))
            PRINC(" = ")
            /*g_try(v2:"Result",loop:true) */
            Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).Declare(ToProperty(OBJ(Core.F_CALL(C_value,ARGS(a1.ToEID())))))
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            F_Generate_breakline_void()
            PRINC("")
            Result = EVOID
            }
            } else {
            F_Generate_symbol_ident_symbol(ToSymbol(a1))
            PRINC(" = ")
            F_Generate_object_prefix_class(C_any,ToClass(a2))
            PRINC("new(")
            F_Generate_go_class_class(ToClass(a2))
            PRINC(").IsNamed(")
            F_Generate_class_ident_class(ToClass(a2))
            PRINC(",MakeSymbol(")
            /*g_try(v2:"Result",loop:true) */
            { var arg_1 *ClaireAny  
              _ = arg_1
              var try_2 EID 
              /*g_try(v2:"try_2",loop:false) */
              try_2 = Core.F_CALL(C_string_I,ARGS(a1.ToEID()))
              /* ERROR PROTECTION INSERTED (arg_1-Result) */
              if ErrorIn(try_2) {Result = try_2
              } else {
              arg_1 = ANY(try_2)
              Result = Core.F_print_any(arg_1)
              }
              } 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC(",")
            /*g_try(v2:"Result",loop:true) */
            Result = F_Generate_g_expression_module(ToModule(OBJ(Core.F_CALL(C_module_I,ARGS(a1.ToEID())))),C_module)
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC("))")
            F_Generate_object_post_class(C_any,ToClass(a2))
            F_Generate_breakline_void()
            PRINC("")
            Result = EVOID
            }}
            } 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          if (s.Id() != C_void.Id()) { 
            F_Generate_breakline_void()
            F_c_princ_string(v)
            PRINC(" = ")
            F_Generate_symbol_ident_symbol(ToSymbol(a1))
            F_Generate_breakline_void()
            PRINC("")
            Result = EVOID
            } else {
            Result = EID{CFALSE.Id(),0}
            } 
          }
          } 
        } 
      }  else if (s.Id() == Optimize.C_EID.Id()) { 
      F_c_princ_string(v)
      PRINC(" = ")
      /*g_try(v2:"Result",loop:true) */
      Result = F_Generate_g_expression_Call(self,s)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      F_Generate_breakline_void()
      PRINC("")
      Result = EVOID
      }
      } else {
      Result = ToException(Core.C_general_error.Make(MakeString("desing error : inline_stat for ~S").Id(),MakeConstantList(self.Id()).Id())).Close()
      } 
    return Result} 
  
// The EID go function for: inline_stat @ Call (throw: true) 
func E_Generate_inline_stat_Call (self EID,s EID,v EID) EID { 
    return F_Generate_inline_stat_Call(Language.To_Call(OBJ(self)),ToClass(OBJ(s)),ToString(OBJ(v)) )} 
  
// A call method is now simpler with unfolding ! very similar structucture
/* {1} The go function for: g_statement(self:Call_method,s:class,v:string,err:boolean,loop:any) [status=1] */
func F_Generate_g_statement_Call_method (self *Language.CallMethod ,s *ClaireClass ,v *ClaireString ,err *ClaireBoolean ,loop *ClaireAny ) EID { 
    var Result EID 
    { var l *ClaireList   = self.Args
      { var ld *ClaireList  
        var try_1 EID 
        /*g_try(v2:"try_1",loop:false) */
        try_1 = F_Generate_unfold_args_list(l)
        /* ERROR PROTECTION INSERTED (ld-Result) */
        if ErrorIn(try_1) {Result = try_1
        } else {
        ld = ToList(OBJ(try_1))
        if (Equal(ld.Id(),CNIL.Id()) == CTRUE) { 
          Result = F_Generate_error_wrap_any(self.Id(),s,v)
          } else {
          { var arg_2 *Language.CallMethod  
            _ = arg_2
            var try_3 EID 
            /*g_try(v2:"try_3",loop:false) */
            { var _CL_obj *Language.CallMethod   = Language.To_CallMethod(new(Language.CallMethod).Is(Language.C_Call_method))
              _CL_obj.Arg = self.Arg
              /*method->method*//*g_try(v2:"try_3",loop:false) */
              { 
                var va_arg1 *Language.CallMethod  
                var va_arg2 *ClaireList  
                va_arg1 = _CL_obj
                var try_4 EID 
                /*g_try(v2:"try_4",loop:false) */
                { 
                  var v_list8 *ClaireList  
                  var z *ClaireAny  
                  var v_local8 *ClaireAny  
                  v_list8 = l
                  try_4 = EID{CreateList(ToType(CEMPTY.Id()),v_list8.Length()).Id(),0}
                  for CLcount := 0; CLcount < v_list8.Length(); CLcount++{ 
                    z = v_list8.At(CLcount)
                    var try_5 EID 
                    /*g_try(v2:"try_5",loop:tuple("try_4", EID)) */
                    try_5 = F_Generate_unfold_arg_list(l,ld,z)
                    /* ERROR PROTECTION INSERTED (v_local8-try_4) */
                    if ErrorIn(try_5) {try_4 = try_5
                    break
                    } else {
                    v_local8 = ANY(try_5)
                    ToList(OBJ(try_4)).PutAt(CLcount,v_local8)
                    } 
                  }
                  } 
                /* ERROR PROTECTION INSERTED (va_arg2-try_3) */
                if ErrorIn(try_4) {try_3 = try_4
                } else {
                va_arg2 = ToList(OBJ(try_4))
                va_arg1.Args = va_arg2
                /*list->list*/try_3 = EID{va_arg2.Id(),0}
                }
                } 
              /* ERROR PROTECTION INSERTED (try_3-try_3) */
              if !ErrorIn(try_3) {
              try_3 = EID{_CL_obj.Id(),0}
              }
              } 
            /* ERROR PROTECTION INSERTED (arg_2-Result) */
            if ErrorIn(try_3) {Result = try_3
            } else {
            arg_2 = Language.To_CallMethod(OBJ(try_3))
            Result = F_Generate_unfold_use_list(ld,
              arg_2.Id(),
              s,
              v,
              err,
              loop)
            }
            } 
          } 
        }
        } 
      } 
    return Result} 
  
// The EID go function for: g_statement @ Call_method (throw: true) 
func E_Generate_g_statement_Call_method (self EID,s EID,v EID,err EID,loop EID) EID { 
    return F_Generate_g_statement_Call_method(Language.To_CallMethod(OBJ(self)),
      ToClass(OBJ(s)),
      ToString(OBJ(v)),
      ToBoolean(OBJ(err)),
      ANY(loop) )} 
  
/* {1} The go function for: g_statement(self:Call_method1,s:class,v:string,err:boolean,loop:any) [status=1] */
func F_Generate_g_statement_Call_method1 (self *Language.CallMethod1 ,s *ClaireClass ,v *ClaireString ,err *ClaireBoolean ,loop *ClaireAny ) EID { 
    var Result EID 
    { var l *ClaireList   = self.Args
      { var ld *ClaireList  
        var try_1 EID 
        /*g_try(v2:"try_1",loop:false) */
        try_1 = F_Generate_unfold_args_list(l)
        /* ERROR PROTECTION INSERTED (ld-Result) */
        if ErrorIn(try_1) {Result = try_1
        } else {
        ld = ToList(OBJ(try_1))
        
        if (Equal(ld.Id(),CNIL.Id()) == CTRUE) { 
          Result = F_Generate_error_wrap_any(self.Id(),s,v)
          } else {
          { var arg_2 *Language.CallMethod1  
            _ = arg_2
            var try_3 EID 
            /*g_try(v2:"try_3",loop:false) */
            { var _CL_obj *Language.CallMethod1   = Language.To_CallMethod1(new(Language.CallMethod1).Is(Language.C_Call_method1))
              _CL_obj.Arg = self.Arg
              /*method->method*//*g_try(v2:"try_3",loop:false) */
              { 
                var va_arg1 *Language.CallMethod  
                var va_arg2 *ClaireList  
                va_arg1 = Language.To_CallMethod(_CL_obj.Id())
                var try_4 EID 
                /*g_try(v2:"try_4",loop:false) */
                { 
                  var v_list8 *ClaireList  
                  var z *ClaireAny  
                  var v_local8 *ClaireAny  
                  v_list8 = l
                  try_4 = EID{CreateList(ToType(CEMPTY.Id()),v_list8.Length()).Id(),0}
                  for CLcount := 0; CLcount < v_list8.Length(); CLcount++{ 
                    z = v_list8.At(CLcount)
                    var try_5 EID 
                    /*g_try(v2:"try_5",loop:tuple("try_4", EID)) */
                    try_5 = F_Generate_unfold_arg_list(l,ld,z)
                    /* ERROR PROTECTION INSERTED (v_local8-try_4) */
                    if ErrorIn(try_5) {try_4 = try_5
                    break
                    } else {
                    v_local8 = ANY(try_5)
                    ToList(OBJ(try_4)).PutAt(CLcount,v_local8)
                    } 
                  }
                  } 
                /* ERROR PROTECTION INSERTED (va_arg2-try_3) */
                if ErrorIn(try_4) {try_3 = try_4
                } else {
                va_arg2 = ToList(OBJ(try_4))
                va_arg1.Args = va_arg2
                /*list->list*/try_3 = EID{va_arg2.Id(),0}
                }
                } 
              /* ERROR PROTECTION INSERTED (try_3-try_3) */
              if !ErrorIn(try_3) {
              try_3 = EID{_CL_obj.Id(),0}
              }
              } 
            /* ERROR PROTECTION INSERTED (arg_2-Result) */
            if ErrorIn(try_3) {Result = try_3
            } else {
            arg_2 = Language.To_CallMethod1(OBJ(try_3))
            Result = F_Generate_unfold_use_list(ld,
              arg_2.Id(),
              s,
              v,
              err,
              loop)
            }
            } 
          } 
        }
        } 
      } 
    return Result} 
  
// The EID go function for: g_statement @ Call_method1 (throw: true) 
func E_Generate_g_statement_Call_method1 (self EID,s EID,v EID,err EID,loop EID) EID { 
    return F_Generate_g_statement_Call_method1(Language.To_CallMethod1(OBJ(self)),
      ToClass(OBJ(s)),
      ToString(OBJ(v)),
      ToBoolean(OBJ(err)),
      ANY(loop) )} 
  
/* {1} The go function for: g_statement(self:Call_method2,s:class,v:string,err:boolean,loop:any) [status=1] */
func F_Generate_g_statement_Call_method2 (self *Language.CallMethod2 ,s *ClaireClass ,v *ClaireString ,err *ClaireBoolean ,loop *ClaireAny ) EID { 
    var Result EID 
    { var l *ClaireList   = self.Args
      { var ld *ClaireList  
        var try_1 EID 
        /*g_try(v2:"try_1",loop:false) */
        try_1 = F_Generate_unfold_args_list(l)
        /* ERROR PROTECTION INSERTED (ld-Result) */
        if ErrorIn(try_1) {Result = try_1
        } else {
        ld = ToList(OBJ(try_1))
        if (Equal(ld.Id(),CNIL.Id()) == CTRUE) { 
          Result = F_Generate_error_wrap_any(self.Id(),s,v)
          } else {
          { var arg_2 *Language.CallMethod2  
            _ = arg_2
            var try_3 EID 
            /*g_try(v2:"try_3",loop:false) */
            { var _CL_obj *Language.CallMethod2   = Language.To_CallMethod2(new(Language.CallMethod2).Is(Language.C_Call_method2))
              _CL_obj.Arg = self.Arg
              /*method->method*//*g_try(v2:"try_3",loop:false) */
              { 
                var va_arg1 *Language.CallMethod  
                var va_arg2 *ClaireList  
                va_arg1 = Language.To_CallMethod(_CL_obj.Id())
                var try_4 EID 
                /*g_try(v2:"try_4",loop:false) */
                { 
                  var v_list8 *ClaireList  
                  var z *ClaireAny  
                  var v_local8 *ClaireAny  
                  v_list8 = l
                  try_4 = EID{CreateList(ToType(CEMPTY.Id()),v_list8.Length()).Id(),0}
                  for CLcount := 0; CLcount < v_list8.Length(); CLcount++{ 
                    z = v_list8.At(CLcount)
                    var try_5 EID 
                    /*g_try(v2:"try_5",loop:tuple("try_4", EID)) */
                    try_5 = F_Generate_unfold_arg_list(l,ld,z)
                    /* ERROR PROTECTION INSERTED (v_local8-try_4) */
                    if ErrorIn(try_5) {try_4 = try_5
                    break
                    } else {
                    v_local8 = ANY(try_5)
                    ToList(OBJ(try_4)).PutAt(CLcount,v_local8)
                    } 
                  }
                  } 
                /* ERROR PROTECTION INSERTED (va_arg2-try_3) */
                if ErrorIn(try_4) {try_3 = try_4
                } else {
                va_arg2 = ToList(OBJ(try_4))
                va_arg1.Args = va_arg2
                /*list->list*/try_3 = EID{va_arg2.Id(),0}
                }
                } 
              /* ERROR PROTECTION INSERTED (try_3-try_3) */
              if !ErrorIn(try_3) {
              try_3 = EID{_CL_obj.Id(),0}
              }
              } 
            /* ERROR PROTECTION INSERTED (arg_2-Result) */
            if ErrorIn(try_3) {Result = try_3
            } else {
            arg_2 = Language.To_CallMethod2(OBJ(try_3))
            Result = F_Generate_unfold_use_list(ld,
              arg_2.Id(),
              s,
              v,
              err,
              loop)
            }
            } 
          } 
        }
        } 
      } 
    return Result} 
  
// The EID go function for: g_statement @ Call_method2 (throw: true) 
func E_Generate_g_statement_Call_method2 (self EID,s EID,v EID,err EID,loop EID) EID { 
    return F_Generate_g_statement_Call_method2(Language.To_CallMethod2(OBJ(self)),
      ToClass(OBJ(s)),
      ToString(OBJ(v)),
      ToBoolean(OBJ(err)),
      ANY(loop) )} 
  
// in Claire 4, Super is handled as a Call
/* {1} The go function for: g_statement(self:Super,s:class,v:string,err:boolean,loop:any) [status=1] */
func F_Generate_g_statement_Super (self *Language.Super ,s *ClaireClass ,v *ClaireString ,err *ClaireBoolean ,loop *ClaireAny ) EID { 
    var Result EID 
    { var l *ClaireList   = self.Args
      { var ld *ClaireList  
        var try_1 EID 
        /*g_try(v2:"try_1",loop:false) */
        try_1 = F_Generate_unfold_args_list(l)
        /* ERROR PROTECTION INSERTED (ld-Result) */
        if ErrorIn(try_1) {Result = try_1
        } else {
        ld = ToList(OBJ(try_1))
        if (Equal(ld.Id(),CNIL.Id()) == CTRUE) { 
          Result = F_Generate_error_wrap_any(self.Id(),s,v)
          } else {
          { var arg_2 *Language.Super  
            _ = arg_2
            var try_3 EID 
            /*g_try(v2:"try_3",loop:false) */
            { var _CL_obj *Language.Super   = Language.To_Super(new(Language.Super).Is(Language.C_Super))
              _CL_obj.Selector = self.Selector
              /*property->property*/_CL_obj.CastTo = self.CastTo
              /*type->type*//*g_try(v2:"try_3",loop:false) */
              { 
                var va_arg1 *Language.Super  
                var va_arg2 *ClaireList  
                va_arg1 = _CL_obj
                var try_4 EID 
                /*g_try(v2:"try_4",loop:false) */
                { 
                  var v_list8 *ClaireList  
                  var z *ClaireAny  
                  var v_local8 *ClaireAny  
                  v_list8 = l
                  try_4 = EID{CreateList(ToType(CEMPTY.Id()),v_list8.Length()).Id(),0}
                  for CLcount := 0; CLcount < v_list8.Length(); CLcount++{ 
                    z = v_list8.At(CLcount)
                    var try_5 EID 
                    /*g_try(v2:"try_5",loop:tuple("try_4", EID)) */
                    try_5 = F_Generate_unfold_arg_list(l,ld,z)
                    /* ERROR PROTECTION INSERTED (v_local8-try_4) */
                    if ErrorIn(try_5) {try_4 = try_5
                    break
                    } else {
                    v_local8 = ANY(try_5)
                    ToList(OBJ(try_4)).PutAt(CLcount,v_local8)
                    } 
                  }
                  } 
                /* ERROR PROTECTION INSERTED (va_arg2-try_3) */
                if ErrorIn(try_4) {try_3 = try_4
                } else {
                va_arg2 = ToList(OBJ(try_4))
                va_arg1.Args = va_arg2
                /*list->list*/try_3 = EID{va_arg2.Id(),0}
                }
                } 
              /* ERROR PROTECTION INSERTED (try_3-try_3) */
              if !ErrorIn(try_3) {
              try_3 = EID{_CL_obj.Id(),0}
              }
              } 
            /* ERROR PROTECTION INSERTED (arg_2-Result) */
            if ErrorIn(try_3) {Result = try_3
            } else {
            arg_2 = Language.To_Super(OBJ(try_3))
            Result = F_Generate_unfold_use_list(ld,
              arg_2.Id(),
              s,
              v,
              err,
              loop)
            }
            } 
          } 
        }
        } 
      } 
    return Result} 
  
// The EID go function for: g_statement @ Super (throw: true) 
func E_Generate_g_statement_Super (self EID,s EID,v EID,err EID,loop EID) EID { 
    return F_Generate_g_statement_Super(Language.To_Super(OBJ(self)),
      ToClass(OBJ(s)),
      ToString(OBJ(v)),
      ToBoolean(OBJ(err)),
      ANY(loop) )} 
  
// trivial 
/* {1} The go function for: g_statement(self:Cast,s:class,v:string,err:boolean,loop:any) [status=1] */
func F_Generate_g_statement_Cast (self *Language.Cast ,s *ClaireClass ,v *ClaireString ,err *ClaireBoolean ,loop *ClaireAny ) EID { 
    var Result EID 
    Result = Core.F_CALL(C_Generate_statement,ARGS(self.Arg.ToEID(),
      EID{s.Id(),0},
      EID{s.Id(),0},
      EID{(v).Id(),0},
      loop.ToEID()))
    return Result} 
  
// The EID go function for: g_statement @ Cast (throw: true) 
func E_Generate_g_statement_Cast (self EID,s EID,v EID,err EID,loop EID) EID { 
    return F_Generate_g_statement_Cast(Language.To_Cast(OBJ(self)),
      ToClass(OBJ(s)),
      ToString(OBJ(v)),
      ToBoolean(OBJ(err)),
      ANY(loop) )} 
  
//-------------- compiling a handle -------------------------------------
// In most cases, s = EID (err = true) and v is an EID variable => reuse v
// in some cases (s != EID => test = any) .. we need a special variable (v2)
// we see if the catch applied (bool : e % S) 
// in CLAIRE4, we know that self.test is a class
/* {1} The go function for: g_statement(self:Handle,s:class,v:string,err:boolean,loop:any) [status=1] */
func F_Generate_g_statement_Handle (self *Language.ClaireHandle ,s *ClaireClass ,v *ClaireString ,err *ClaireBoolean ,loop *ClaireAny ) EID { 
    var Result EID 
    { var v2 *ClaireString  
      if (s.Id() == Optimize.C_EID.Id()) { 
        v2 = v
        } else {
        v2 = F_append_string(v,MakeString("_H"))
        } 
      F_Generate_new_block_string(MakeString("handle"))
      if (s.Id() != Optimize.C_EID.Id()) { 
        F_Generate_var_declaration_string(v2,Optimize.C_EID,1)
        } 
      PRINC("h_index := ClEnv.Index")
      F_Generate_breakline_void()
      PRINC("")
      PRINC("h_base := ClEnv.Base")
      F_Generate_breakline_void()
      PRINC("")
      /*g_try(v2:"Result",loop:true) */
      Result = F_Generate_statement_any(self.Arg,Optimize.C_EID,v2,CFALSE.Id())
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      /*g_try(v2:"Result",loop:true) */
      if (self.Test == C_any.Id()) { 
        PRINC("if ErrorIn(")
        F_princ_string(v2)
        PRINC(")")
        Result = EVOID
        } else {
        PRINC("if ErrorIn(")
        F_princ_string(v2)
        PRINC(") && ")
        /*g_try(v2:"Result",loop:true) */
        Result = Core.F_CALL(C_Generate_g_expression,ARGS(self.Test.ToEID(),EID{C_type.Id(),0}))
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC(".Contains(ANY(")
        F_princ_string(v2)
        PRINC(")) == CTRUE ")
        Result = EVOID
        }
        } 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      F_Generate_new_block_void()
      /*g_try(v2:"Result",loop:true) */
      if (ToGenerateGoProducer(Optimize.C_PRODUCER.Value).Debug_ask == CTRUE) { 
        PRINC("/* s=")
        /*g_try(v2:"Result",loop:true) */
        Result = Core.F_print_any(s.Id())
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC(" */")
        Result = EVOID
        }
        } else {
        Result = EID{CFALSE.Id(),0}
        } 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC("ClEnv.Index = h_index")
      F_Generate_breakline_void()
      PRINC("")
      PRINC("ClEnv.Base = h_base")
      F_Generate_breakline_void()
      PRINC("")
      /*g_try(v2:"Result",loop:true) */
      Result = F_Generate_statement_any(self.Other,s,v,loop)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      /*g_try(v2:"Result",loop:true) */
      if ((s.Id() == Optimize.C_EID.Id()) || 
          (s.Id() == C_void.Id())) { 
        F_Generate_close_block_void()
        Result = EVOID
        } else {
        PRINC("} else {")
        F_Generate_breakline_void()
        F_c_princ_string(v)
        PRINC(" = ")
        /*g_try(v2:"Result",loop:true) */
        Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).FromEid(v2,s)
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        F_Generate_breakline_void()
        F_Generate_close_block_void()
        PRINC("")
        Result = EVOID
        }
        } 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      F_Generate_close_block_string(MakeString("handle"))
      Result = EVOID
      }}}}}
      } 
    return Result} 
  
// The EID go function for: g_statement @ Handle (throw: true) 
func E_Generate_g_statement_Handle (self EID,s EID,v EID,err EID,loop EID) EID { 
    return F_Generate_g_statement_Handle(Language.To_ClaireHandle(OBJ(self)),
      ToClass(OBJ(s)),
      ToString(OBJ(v)),
      ToBoolean(OBJ(err)),
      ANY(loop) )} 
  
// to_CL and to_C are presently ignored in CLAIRE 4
// [g_statement(self:Generate/to_CL,s:class,v:string,err:boolean,loop:any) : void
// -> g_statement(self.arg, s, v, err, loop) ]
// [g_statement(self:Generate/to_C,s:class,v:string,err:boolean,loop:any) : void
// -> g_statement(self.arg, s, v, err, loop) ]
// same for a cast
// v3.2.06: the case where self.arg is of type any is painful => it is forbiden in osystem.cl
/* {1} The go function for: g_statement(self:Compile/C_cast,s:class,v:string,err:boolean,loop:any) [status=1] */
func F_Generate_g_statement_C_cast (self *Optimize.Compile_CCast ,s *ClaireClass ,v *ClaireString ,err *ClaireBoolean ,loop *ClaireAny ) EID { 
    var Result EID 
    Result = Core.F_CALL(C_Generate_g_statement,ARGS(self.Arg.ToEID(),
      EID{s.Id(),0},
      EID{(v).Id(),0},
      EID{err.Id(),0},
      loop.ToEID()))
    return Result} 
  
// The EID go function for: g_statement @ Compile/C_cast (throw: true) 
func E_Generate_g_statement_C_cast (self EID,s EID,v EID,err EID,loop EID) EID { 
    return F_Generate_g_statement_C_cast(Optimize.To_Compile_CCast(OBJ(self)),
      ToClass(OBJ(s)),
      ToString(OBJ(v)),
      ToBoolean(OBJ(err)),
      ANY(loop) )} 
  
//------------- compiling slot read/write -------------------------------
// new in CLAIRE 4 : there are two kinds => err (EID required) or not (self.arg is just too complex)
// we will follow a pattern similar to unfold => create the let then call g_statement on it
// reads a slot.
// there are two reasons for requiring a statement : complex arg or possible error when reading ! hence we check before using unfold ...
/* {1} The go function for: g_statement(self:Call_slot,s:class,v:string,err:boolean,loop:any) [status=1] */
func F_Generate_g_statement_Call_slot (self *Language.CallSlot ,s *ClaireClass ,v *ClaireString ,err *ClaireBoolean ,loop *ClaireAny ) EID { 
    var Result EID 
    var g0187I *ClaireBoolean  
    var try_1 EID 
    /*g_try(v2:"try_1",loop:false) */
    try_1 = F_Generate_g_clean_any(self.Arg)
    /* ERROR PROTECTION INSERTED (g0187I-Result) */
    if ErrorIn(try_1) {Result = try_1
    } else {
    g0187I = ToBoolean(OBJ(try_1))
    if (g0187I == CTRUE) { 
      F_c_princ_string(v)
      PRINC(" = ")
      /*g_try(v2:"Result",loop:true) */
      Result = F_Generate_g_expression_Call_slot(self,s)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      F_Generate_breakline_void()
      PRINC("")
      Result = EVOID
      }
      } else {
      { var varg *ClaireVariable  
        var try_2 EID 
        /*g_try(v2:"try_2",loop:false) */
        { var arg_3 *ClaireType  
          _ = arg_3
          var try_4 EID 
          /*g_try(v2:"try_4",loop:false) */
          try_4 = Core.F_CALL(Optimize.C_c_type,ARGS(self.Arg.ToEID()))
          /* ERROR PROTECTION INSERTED (arg_3-try_2) */
          if ErrorIn(try_4) {try_2 = try_4
          } else {
          arg_3 = ToType(OBJ(try_4))
          try_2 = EID{F_Generate_build_Variable_string(MakeString("v_slot"),arg_3.Id()).Id(),0}
          }
          } 
        /* ERROR PROTECTION INSERTED (varg-Result) */
        if ErrorIn(try_2) {Result = try_2
        } else {
        varg = To_Variable(OBJ(try_2))
        { var unfold *Language.Let  
          _ = unfold
          { var _CL_obj *Language.Let   = Language.To_Let(new(Language.Let).Is(Language.C_Let))
            _CL_obj.ClaireVar = varg
            /*Variable->Variable*/_CL_obj.Value = self.Arg
            /*any->any*/{ 
              var va_arg1 *Language.Let  
              var va_arg2 *ClaireAny  
              va_arg1 = _CL_obj
              { var _CL_obj *Language.CallSlot   = Language.To_CallSlot(new(Language.CallSlot).Is(Language.C_Call_slot))
                _CL_obj.Selector = self.Selector
                /*slot->slot*/_CL_obj.Arg = varg.Id()
                /*any->any*/va_arg2 = _CL_obj.Id()
                } 
              va_arg1.Arg = va_arg2
              /*any->any*/} 
            unfold = _CL_obj
            } 
          Result = F_Generate_g_statement_Let(unfold,
            s,
            v,
            err,
            loop)
          } 
        }
        } 
      } 
    }
    return Result} 
  
// The EID go function for: g_statement @ Call_slot (throw: true) 
func E_Generate_g_statement_Call_slot (self EID,s EID,v EID,err EID,loop EID) EID { 
    return F_Generate_g_statement_Call_slot(Language.To_CallSlot(OBJ(self)),
      ToClass(OBJ(s)),
      ToString(OBJ(v)),
      ToBoolean(OBJ(err)),
      ANY(loop) )} 
  
// reads an table.
// there are two reasons for requiring a statement : complex arg or possible error when reading !
/* {1} The go function for: g_statement(self:Call_table,s:class,v:string,err:boolean,loop:any) [status=1] */
func F_Generate_g_statement_Call_table (self *Language.CallTable ,s *ClaireClass ,v *ClaireString ,err *ClaireBoolean ,loop *ClaireAny ) EID { 
    var Result EID 
    var g0188I *ClaireBoolean  
    var try_1 EID 
    /*g_try(v2:"try_1",loop:false) */
    try_1 = F_Generate_g_clean_any(self.Arg)
    /* ERROR PROTECTION INSERTED (g0188I-Result) */
    if ErrorIn(try_1) {Result = try_1
    } else {
    g0188I = ToBoolean(OBJ(try_1))
    if (g0188I == CTRUE) { 
      F_c_princ_string(v)
      PRINC(" = ")
      /*g_try(v2:"Result",loop:true) */
      Result = F_Generate_g_expression_Call_table(self,s)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      F_Generate_breakline_void()
      PRINC("")
      Result = EVOID
      }
      } else {
      { var varg *ClaireVariable  
        var try_2 EID 
        /*g_try(v2:"try_2",loop:false) */
        { var arg_3 *ClaireType  
          _ = arg_3
          var try_4 EID 
          /*g_try(v2:"try_4",loop:false) */
          try_4 = Core.F_CALL(Optimize.C_c_type,ARGS(self.Arg.ToEID()))
          /* ERROR PROTECTION INSERTED (arg_3-try_2) */
          if ErrorIn(try_4) {try_2 = try_4
          } else {
          arg_3 = ToType(OBJ(try_4))
          try_2 = EID{F_Generate_build_Variable_string(MakeString("v_table"),arg_3.Id()).Id(),0}
          }
          } 
        /* ERROR PROTECTION INSERTED (varg-Result) */
        if ErrorIn(try_2) {Result = try_2
        } else {
        varg = To_Variable(OBJ(try_2))
        { var unfold *Language.Let  
          _ = unfold
          { var _CL_obj *Language.Let   = Language.To_Let(new(Language.Let).Is(Language.C_Let))
            _CL_obj.ClaireVar = varg
            /*Variable->Variable*/_CL_obj.Value = self.Arg
            /*any->any*/{ 
              var va_arg1 *Language.Let  
              var va_arg2 *ClaireAny  
              va_arg1 = _CL_obj
              { var _CL_obj *Language.CallTable   = Language.To_CallTable(new(Language.CallTable).Is(Language.C_Call_table))
                _CL_obj.Selector = self.Selector
                /*table->table*/_CL_obj.Arg = varg.Id()
                /*any->any*/va_arg2 = _CL_obj.Id()
                } 
              va_arg1.Arg = va_arg2
              /*any->any*/} 
            unfold = _CL_obj
            } 
          Result = F_Generate_g_statement_Let(unfold,
            s,
            v,
            err,
            loop)
          } 
        }
        } 
      } 
    }
    return Result} 
  
// The EID go function for: g_statement @ Call_table (throw: true) 
func E_Generate_g_statement_Call_table (self EID,s EID,v EID,err EID,loop EID) EID { 
    return F_Generate_g_statement_Call_table(Language.To_CallTable(OBJ(self)),
      ToClass(OBJ(s)),
      ToString(OBJ(v)),
      ToBoolean(OBJ(err)),
      ANY(loop) )} 
  
// reads an array.
/* {1} The go function for: g_statement(self:Call_array,s:class,v:string,err:boolean,loop:any) [status=1] */
func F_Generate_g_statement_Call_array (self *Language.CallArray ,s *ClaireClass ,v *ClaireString ,err *ClaireBoolean ,loop *ClaireAny ) EID { 
    var Result EID 
    { var varg1 *ClaireVariable   = F_Generate_build_Variable_string(MakeString("va_arg1"),C_array.Id())
      { var varg2 *ClaireVariable   = F_Generate_build_Variable_string(MakeString("va_arg2"),C_integer.Id())
        { var unfold *Language.Let  
          _ = unfold
          { var _CL_obj *Language.Let   = Language.To_Let(new(Language.Let).Is(Language.C_Let))
            _CL_obj.ClaireVar = varg1
            /*Variable->Variable*/_CL_obj.Value = self.Selector
            /*any->any*/{ 
              var va_arg1 *Language.Let  
              var va_arg2 *ClaireAny  
              va_arg1 = _CL_obj
              { var _CL_obj *Language.Let   = Language.To_Let(new(Language.Let).Is(Language.C_Let))
                _CL_obj.ClaireVar = varg2
                /*Variable->Variable*/_CL_obj.Value = self.Arg
                /*any->any*/_CL_obj.Arg = Language.C_Call_array.Make(varg1.Id(),varg2.Id(),self.Test)
                /*any->any*/va_arg2 = _CL_obj.Id()
                } 
              va_arg1.Arg = va_arg2
              /*any->any*/} 
            unfold = _CL_obj
            } 
          Result = F_Generate_g_statement_Let(unfold,
            s,
            v,
            err,
            loop)
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: g_statement @ Call_array (throw: true) 
func E_Generate_g_statement_Call_array (self EID,s EID,v EID,err EID,loop EID) EID { 
    return F_Generate_g_statement_Call_array(Language.To_CallArray(OBJ(self)),
      ToClass(OBJ(s)),
      ToString(OBJ(v)),
      ToBoolean(OBJ(err)),
      ANY(loop) )} 
  
// places a value in a slot with similar conventions ------------------------------------------------------------------------
// Update = [R(x) := y] where R(x) is a Call_slot, a call_array or a call_table 
// THIS USE OF self.arg IS MEGA UGLY AND SHOULD BE SIMPLIFIED IN THE OPTIMIZER LATER ON ... THERE SHOULD AT LEAST EXIST SOME COMMENTS !
// self.arg is a meta parameter /  it is a property (add or put ...) unless a demon if_write is used 
// self.value is Y and self.var is R(x)  => look in goexp.cl 
/* {1} The go function for: g_statement(self:Update,s:class,v:string,err:boolean,loop:any) [status=1] */
func F_Generate_g_statement_Update (self *Language.Update ,s *ClaireClass ,v *ClaireString ,err *ClaireBoolean ,loop *ClaireAny ) EID { 
    var Result EID 
    { var X *ClaireAny   = self.ClaireVar
      { var p *ClaireAny   = self.Selector
        { var sr *ClaireType  
          if (X.Isa.IsIn(Language.C_Call_slot) == CTRUE) { 
            { var g0189 *Language.CallSlot   = Language.To_CallSlot(X)
              if (self.Arg == C_add.Id()) { 
                sr = Core.F_member_type(g0189.Selector.Range)
                } else {
                sr = g0189.Selector.Range
                } 
              } 
            }  else if (X.Isa.IsIn(Language.C_Call_array) == CTRUE) { 
            { var g0190 *Language.CallArray   = Language.To_CallArray(X)
              _ = g0190
              if (ToType(g0190.Test).Included(ToType(C_float.Id())) == CTRUE) { 
                sr = ToType(C_float.Id())
                } else {
                sr = ToType(C_any.Id())
                } 
              } 
            } else {
            { var arg_1 *ClaireAny  
              _ = arg_1
              if (self.Arg == C_add.Id()) { 
                arg_1 = Core.F_member_type(ToType(OBJ(Core.F_CALL(C_range,ARGS(p.ToEID()))))).Id()
                } else {
                arg_1 = ANY(Core.F_CALL(C_range,ARGS(p.ToEID())))
                } 
              sr = Core.F_U_type(ToType(C_any.Id()),ToType(arg_1))
              } 
            } 
          
          var g0195I *ClaireBoolean  
          var try_2 EID 
          /*g_try(v2:"try_2",loop:false) */
          { 
            var v_and5 *ClaireBoolean  
            
            v_and5 = err.Not
            if (v_and5 == CFALSE) {try_2 = EID{CFALSE.Id(),0}
            } else { 
              var try_3 EID 
              /*g_try(v2:"try_3",loop:false) */
              try_3 = F_Generate_g_func_any(ANY(Core.F_CALL(C_arg,ARGS(self.ClaireVar.ToEID()))))
              /* ERROR PROTECTION INSERTED (v_and5-try_2) */
              if ErrorIn(try_3) {try_2 = try_3
              } else {
              v_and5 = ToBoolean(OBJ(try_3))
              if (v_and5 == CFALSE) {try_2 = EID{CFALSE.Id(),0}
              } else { 
                var try_4 EID 
                /*g_try(v2:"try_4",loop:false) */
                try_4 = F_Generate_g_func_any(self.Value)
                /* ERROR PROTECTION INSERTED (v_and5-try_2) */
                if ErrorIn(try_4) {try_2 = try_4
                } else {
                v_and5 = ToBoolean(OBJ(try_4))
                if (v_and5 == CFALSE) {try_2 = EID{CFALSE.Id(),0}
                } else { 
                  v_and5 = Equal(s.Id(),C_void.Id())
                  if (v_and5 == CFALSE) {try_2 = EID{CFALSE.Id(),0}
                  } else { 
                    try_2 = EID{CTRUE.Id(),0}} 
                  } 
                } 
              } 
            }}
            } 
          /* ERROR PROTECTION INSERTED (g0195I-Result) */
          if ErrorIn(try_2) {Result = try_2
          } else {
          g0195I = ToBoolean(OBJ(try_2))
          if (g0195I == CTRUE) { 
            Result = F_Generate_update_statement_Update(self,sr.Class_I())
            } else {
            var g0196I *ClaireBoolean  
            var try_5 EID 
            /*g_try(v2:"try_5",loop:false) */
            { 
              var v_and6 *ClaireBoolean  
              
              var try_6 EID 
              /*g_try(v2:"try_6",loop:false) */
              try_6 = F_Generate_g_clean_any(ANY(Core.F_CALL(C_arg,ARGS(self.ClaireVar.ToEID()))))
              /* ERROR PROTECTION INSERTED (v_and6-try_5) */
              if ErrorIn(try_6) {try_5 = try_6
              } else {
              v_and6 = ToBoolean(OBJ(try_6))
              if (v_and6 == CFALSE) {try_5 = EID{CFALSE.Id(),0}
              } else { 
                var try_7 EID 
                /*g_try(v2:"try_7",loop:false) */
                try_7 = F_Generate_g_clean_any(self.Value)
                /* ERROR PROTECTION INSERTED (v_and6-try_5) */
                if ErrorIn(try_7) {try_5 = try_7
                } else {
                v_and6 = ToBoolean(OBJ(try_7))
                if (v_and6 == CFALSE) {try_5 = EID{CFALSE.Id(),0}
                } else { 
                  v_and6 = Optimize.F_Compile_update_write_ask_Update(self)
                  if (v_and6 == CFALSE) {try_5 = EID{CFALSE.Id(),0}
                  } else { 
                    try_5 = EID{CTRUE.Id(),0}} 
                  } 
                } 
              }}
              } 
            /* ERROR PROTECTION INSERTED (g0196I-Result) */
            if ErrorIn(try_5) {Result = try_5
            } else {
            g0196I = ToBoolean(OBJ(try_5))
            if (g0196I == CTRUE) { 
              /*g_try(v2:"Result",loop:true) */
              PRINC("/* update_write ")
              /*g_try(v2:"Result",loop:true) */
              Result = Core.F_print_any(self.Id())
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              PRINC(" : ")
              /*g_try(v2:"Result",loop:true) */
              Result = Core.F_print_any(ANY(Core.F_CALL(C_domain,ARGS(self.Selector.ToEID()))))
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              PRINC(" x ")
              /*g_try(v2:"Result",loop:true) */
              Result = Core.F_print_any(ANY(Core.F_CALL(C_range,ARGS(self.Selector.ToEID()))))
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              PRINC(" */")
              F_Generate_breakline_void()
              PRINC("")
              Result = EVOID
              }}}
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              F_princ_string(v)
              PRINC(" = ")
              Result = F_Generate_update_statement_Update(self,sr.Class_I())
              }
              } else {
              { var try_count int  = 0
                _ = try_count
                { var varg1 *ClaireVariable  
                  { var arg_8 *ClaireType  
                    _ = arg_8
                    if (X.Isa.IsIn(Language.C_Call_slot) == CTRUE) { 
                      { var g0192 *Language.CallSlot   = Language.To_CallSlot(X)
                        _ = g0192
                        arg_8 = ToType(Core.F_domain_I_restriction(ToRestriction(g0192.Selector.Id())).Id())
                        } 
                      }  else if (X.Isa.IsIn(Language.C_Call_array) == CTRUE) { 
                      arg_8 = ToType(C_integer.Id())
                      } else {
                      arg_8 = Core.F_U_type(ToType(C_any.Id()),ToType(OBJ(Core.F_CALL(C_domain,ARGS(p.ToEID())))))
                      } 
                    varg1 = F_Generate_build_Variable_string(MakeString("va_arg1"),arg_8.Id())
                    } 
                  { var varg2 *ClaireVariable   = F_Generate_build_Variable_string(MakeString("va_arg2"),sr.Id())
                    { var _Zcall *ClaireAny  
                      _ = _Zcall
                      var try_9 EID 
                      /*g_try(v2:"try_9",loop:false) */
                      { var xx *ClaireAny   = ANY(Core.F_CALL(C_copy,ARGS(X.ToEID())))
                        /*g_try(v2:"try_9",loop:false) */
                        try_9 = Core.F_put_property2(C_arg,ToObject(xx),varg1.Id())
                        /* ERROR PROTECTION INSERTED (try_9-try_9) */
                        if !ErrorIn(try_9) {
                        try_9 = xx.ToEID()
                        }
                        } 
                      /* ERROR PROTECTION INSERTED (_Zcall-Result) */
                      if ErrorIn(try_9) {Result = try_9
                      } else {
                      _Zcall = ANY(try_9)
                      { var _Zunfold *Language.Update  
                        { var _CL_obj *Language.Update   = Language.To_Update(new(Language.Update).Is(Language.C_Update))
                          _CL_obj.Selector = self.Selector
                          /*any->any*/_CL_obj.Value = varg2.Id()
                          /*any->any*/_CL_obj.Arg = self.Arg
                          /*any->any*/_CL_obj.ClaireVar = _Zcall
                          /*any->any*/_Zunfold = _CL_obj
                          } 
                        F_Generate_new_block_string(MakeString("update"))
                        /*g_try(v2:"Result",loop:true) */
                        { var arg_10 *ClaireString  
                          _ = arg_10
                          var try_11 EID 
                          /*g_try(v2:"try_11",loop:false) */
                          try_11 = F_Generate_c_string_go_producer1(ToGenerateGoProducer(Optimize.C_PRODUCER.Value),varg1)
                          /* ERROR PROTECTION INSERTED (arg_10-Result) */
                          if ErrorIn(try_11) {Result = try_11
                          } else {
                          arg_10 = ToString(OBJ(try_11))
                          F_Generate_var_declaration_string(arg_10,F_Generate_go_range_Variable(varg1),1)
                          Result = EVOID
                          }
                          } 
                        /* ERROR PROTECTION INSERTED (Result-Result) */
                        if !ErrorIn(Result) {
                        /*g_try(v2:"Result",loop:true) */
                        { var arg_12 *ClaireString  
                          _ = arg_12
                          var try_13 EID 
                          /*g_try(v2:"try_13",loop:false) */
                          try_13 = F_Generate_c_string_go_producer1(ToGenerateGoProducer(Optimize.C_PRODUCER.Value),varg2)
                          /* ERROR PROTECTION INSERTED (arg_12-Result) */
                          if ErrorIn(try_13) {Result = try_13
                          } else {
                          arg_12 = ToString(OBJ(try_13))
                          F_Generate_var_declaration_string(arg_12,F_Generate_go_range_Variable(varg2),1)
                          Result = EVOID
                          }
                          } 
                        /* ERROR PROTECTION INSERTED (Result-Result) */
                        if !ErrorIn(Result) {
                        /*g_try(v2:"Result",loop:true) */
                        var g0197I *ClaireBoolean  
                        var try_14 EID 
                        /*g_try(v2:"try_14",loop:false) */
                        try_14 = Optimize.F_Compile_g_throw_any(ANY(Core.F_CALL(C_arg,ARGS(X.ToEID()))))
                        /* ERROR PROTECTION INSERTED (g0197I-Result) */
                        if ErrorIn(try_14) {Result = try_14
                        } else {
                        g0197I = ToBoolean(OBJ(try_14))
                        if (g0197I == CTRUE) { 
                          try_count = (try_count+1)
                          Result = Core.F_CALL(C_Generate_g_try,ARGS(Core.F_CALL(C_arg,ARGS(X.ToEID())),
                            EID{MakeString("va_arg1").Id(),0},
                            EID{F_Generate_go_range_Variable(varg1).Id(),0},
                            EID{(v).Id(),0},
                            EID{CFALSE.Id(),0}))
                          } else {
                          Result = F_Generate_statement_any(ANY(Core.F_CALL(C_arg,ARGS(X.ToEID()))),F_Generate_go_range_Variable(varg1),MakeString("va_arg1"),loop)
                          } 
                        }
                        /* ERROR PROTECTION INSERTED (Result-Result) */
                        if !ErrorIn(Result) {
                        /*g_try(v2:"Result",loop:true) */
                        var g0198I *ClaireBoolean  
                        var try_15 EID 
                        /*g_try(v2:"try_15",loop:false) */
                        try_15 = Optimize.F_Compile_g_throw_any(self.Value)
                        /* ERROR PROTECTION INSERTED (g0198I-Result) */
                        if ErrorIn(try_15) {Result = try_15
                        } else {
                        g0198I = ToBoolean(OBJ(try_15))
                        if (g0198I == CTRUE) { 
                          try_count = (try_count+1)
                          Result = Core.F_CALL(C_Generate_g_try,ARGS(self.Value.ToEID(),
                            EID{MakeString("va_arg2").Id(),0},
                            EID{F_Generate_go_range_Variable(varg2).Id(),0},
                            EID{(v).Id(),0},
                            EID{CFALSE.Id(),0}))
                          } else {
                          Result = F_Generate_statement_any(self.Value,F_Generate_go_range_Variable(varg2),MakeString("va_arg2"),loop)
                          } 
                        }
                        /* ERROR PROTECTION INSERTED (Result-Result) */
                        if !ErrorIn(Result) {
                        /*g_try(v2:"Result",loop:true) */
                        if (ToGenerateGoProducer(Optimize.C_PRODUCER.Value).Debug_ask == CTRUE) { 
                          PRINC("/* ---------- now we compile update ")
                          /*g_try(v2:"Result",loop:true) */
                          Result = Core.F_print_any(_Zunfold.Id())
                          /* ERROR PROTECTION INSERTED (Result-Result) */
                          if !ErrorIn(Result) {
                          PRINC(" ------- */")
                          F_Generate_breakline_void()
                          PRINC("")
                          Result = EVOID
                          }
                          } else {
                          Result = EID{CFALSE.Id(),0}
                          } 
                        /* ERROR PROTECTION INSERTED (Result-Result) */
                        if !ErrorIn(Result) {
                        if (Optimize.F_Compile_update_write_ask_Update(self) == CTRUE) { 
                          F_princ_string(v)
                          PRINC(" = ")
                          } 
                        /*g_try(v2:"Result",loop:true) */
                        Result = F_Generate_update_statement_Update(_Zunfold,sr.Class_I())
                        /* ERROR PROTECTION INSERTED (Result-Result) */
                        if !ErrorIn(Result) {
                        /*g_try(v2:"Result",loop:true) */
                        if ((s.Id() != C_void.Id()) && 
                            (Optimize.F_Compile_update_write_ask_Update(self) != CTRUE)) { 
                          F_princ_string(v)
                          PRINC(" = ")
                          /*g_try(v2:"Result",loop:true) */
                          Result = F_Generate_cast_prefix_class(F_Generate_go_range_Variable(varg2),s)
                          /* ERROR PROTECTION INSERTED (Result-Result) */
                          if !ErrorIn(Result) {
                          PRINC("va_arg2")
                          F_Generate_cast_post_class(F_Generate_go_range_Variable(varg2),s)
                          F_Generate_breakline_void()
                          PRINC("")
                          Result = EVOID
                          }
                          } else {
                          Result = EID{CFALSE.Id(),0}
                          } 
                        /* ERROR PROTECTION INSERTED (Result-Result) */
                        if !ErrorIn(Result) {
                        F_Generate_close_try_integer(try_count)
                        F_Generate_close_block_string(MakeString("update"))
                        Result = EVOID
                        }}}}}}}
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
  
// The EID go function for: g_statement @ Update (throw: true) 
func E_Generate_g_statement_Update (self EID,s EID,v EID,err EID,loop EID) EID { 
    return F_Generate_g_statement_Update(Language.To_Update(OBJ(self)),
      ToClass(OBJ(s)),
      ToString(OBJ(v)),
      ToBoolean(OBJ(err)),
      ANY(loop) )} 
  
// this produce the code for an update assuming that self is error-free and functional
// this methiod handles
//    if_write demons (that perform the update)  p_write(x:any,y:any)
//    defeasible updates   o.StoreX(n,v,CTRUE)
// if we cannot find n (type too generic) => revert to a generic Update method
// NOTE: Update is used for many things:
//   Update(p:property, arg: ss | c_code(x,any),  var:call_slot, value:y) 
//   Update(p:property, arg: add,  var:call_slot, value:y) multivalued
//   Update(p:table, arg: put | c_code(x,any), var: call_table, value:y)   // only with a list-based table!
//   Update(p:exp<array>, arg:put, var:call_array(p,x), value:y)
/* {1} The go function for: update_statement(self:Update,s:class) [status=1] */
func F_Generate_update_statement_Update (self *Language.Update ,s *ClaireClass ) EID { 
    var Result EID 
    { var p *ClaireAny   = self.Selector
      { var a *ClaireAny   = self.Arg
        _ = a
        { var v *ClaireAny   = self.Value
          { var x *ClaireAny   = self.ClaireVar
            { var s2 *ClaireClass  
              var try_1 EID 
              /*g_try(v2:"try_1",loop:false) */
              { var arg_2 *ClaireType  
                _ = arg_2
                var try_3 EID 
                /*g_try(v2:"try_3",loop:false) */
                try_3 = Core.F_CALL(Optimize.C_c_type,ARGS(self.ClaireVar.ToEID()))
                /* ERROR PROTECTION INSERTED (arg_2-try_1) */
                if ErrorIn(try_3) {try_1 = try_3
                } else {
                arg_2 = ToType(OBJ(try_3))
                try_1 = EID{arg_2.Class_I().Id(),0}
                }
                } 
              /* ERROR PROTECTION INSERTED (s2-Result) */
              if ErrorIn(try_1) {Result = try_1
              } else {
              s2 = ToClass(OBJ(try_1))
              if (Optimize.F_Compile_update_write_ask_Update(self) == CTRUE) { 
                Optimize.F_Compile_Tighten_I_relation(ToRelation(p))
                PRINC("F_")
                /*g_try(v2:"Result",loop:true) */
                { var arg_4 *ClaireAny  
                  _ = arg_4
                  var try_5 EID 
                  /*g_try(v2:"try_5",loop:false) */
                  try_5 = Core.F_CALL(C_string_I,ARGS(Core.F_CALL(C_name,ARGS(p.ToEID()))))
                  /* ERROR PROTECTION INSERTED (arg_4-Result) */
                  if ErrorIn(try_5) {Result = try_5
                  } else {
                  arg_4 = ANY(try_5)
                  F_c_princ_string(ToString(arg_4))
                  Result = EVOID
                  }
                  } 
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                PRINC("_write(")
                /*g_try(v2:"Result",loop:true) */
                Result = Core.F_CALL(C_Generate_g_expression,ARGS(Core.F_CALL(C_arg,ARGS(x.ToEID())),EID{ToTypeExpression(OBJ(Core.F_CALL(C_domain,ARGS(p.ToEID())))).Class_I().Id(),0}))
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                PRINC(",")
                /*g_try(v2:"Result",loop:true) */
                Result = Core.F_CALL(C_Generate_g_expression,ARGS(v.ToEID(),EID{ToTypeExpression(OBJ(Core.F_CALL(C_range,ARGS(p.ToEID())))).Class_I().Id(),0}))
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                PRINC(")")
                F_Generate_breakline_void()
                PRINC("")
                Result = EVOID
                }}}
                } else {
                var g0204I *ClaireBoolean  
                if (p.Isa.IsIn(C_relation) == CTRUE) { 
                  { var g0199 *ClaireRelation   = ToRelation(p)
                    _ = g0199
                    g0204I = MakeBoolean((g0199.Store_ask == CTRUE) || (a == Core.C_put_store.Id()))
                    } 
                  } else {
                  g0204I = CFALSE
                  } 
                if (g0204I == CTRUE) { 
                  if (x.Isa.IsIn(Language.C_Call_table) == CTRUE) { 
                    PRINC("F_store_list(ToList(")
                    /*g_try(v2:"Result",loop:true) */
                    Result = Core.F_CALL(C_Generate_g_expression,ARGS(p.ToEID(),EID{C_table.Id(),0}))
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    PRINC(".Graph),")
                    /*g_try(v2:"Result",loop:true) */
                    Result = F_Generate_g_table_index_table(ToTable(p),ANY(Core.F_CALL(C_arg,ARGS(x.ToEID()))))
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    PRINC(",")
                    /*g_try(v2:"Result",loop:true) */
                    Result = Core.F_CALL(C_Generate_g_expression,ARGS(v.ToEID(),EID{C_any.Id(),0}))
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    PRINC(",CTRUE)")
                    F_Generate_breakline_void()
                    PRINC("")
                    Result = EVOID
                    }}}
                    } else {
                    { var s2 *ClaireClass  
                      var try_6 EID 
                      /*g_try(v2:"try_6",loop:false) */
                      { var arg_7 *ClaireType  
                        _ = arg_7
                        var try_8 EID 
                        /*g_try(v2:"try_8",loop:false) */
                        try_8 = Core.F_CALL(Optimize.C_c_type,ARGS(Language.To_CallSlot(x).Arg.ToEID()))
                        /* ERROR PROTECTION INSERTED (arg_7-try_6) */
                        if ErrorIn(try_8) {try_6 = try_8
                        } else {
                        arg_7 = ToType(OBJ(try_8))
                        try_6 = EID{arg_7.Class_I().Id(),0}
                        }
                        } 
                      /* ERROR PROTECTION INSERTED (s2-Result) */
                      if ErrorIn(try_6) {Result = try_6
                      } else {
                      s2 = ToClass(OBJ(try_6))
                      { var n int  = INT(Core.F_CALL(C_mClaire_index,ARGS(EID{Core.F__at_property1(ToProperty(p),s2).Id(),0})))
                        _ = n
                        /*g_try(v2:"Result",loop:true) */
                        Result = Core.F_CALL(C_Generate_g_expression,ARGS(Core.F_CALL(C_arg,ARGS(x.ToEID())),EID{C_object.Id(),0}))
                        /* ERROR PROTECTION INSERTED (Result-Result) */
                        if !ErrorIn(Result) {
                        PRINC(".Store")
                        { var arg_9 *ClaireString  
                          _ = arg_9
                          if (s.Id() == C_integer.Id()) { 
                            arg_9 = MakeString("Int")
                            }  else if (s.Id() == C_float.Id()) { 
                            arg_9 = MakeString("Float")
                            } else {
                            arg_9 = MakeString("Obj")
                            } 
                          F_princ_string(arg_9)
                          } 
                        PRINC("(")
                        F_princ_integer(n)
                        PRINC(",")
                        /*g_try(v2:"Result",loop:true) */
                        { var arg_10 *ClaireClass  
                          _ = arg_10
                          if ((s.Id() == C_integer.Id()) || 
                              (s.Id() == C_float.Id())) { 
                            arg_10 = s
                            } else {
                            arg_10 = C_any
                            } 
                          Result = Core.F_CALL(C_Generate_g_expression,ARGS(v.ToEID(),EID{arg_10.Id(),0}))
                          } 
                        /* ERROR PROTECTION INSERTED (Result-Result) */
                        if !ErrorIn(Result) {
                        PRINC(",CTRUE)/*c:")
                        /*g_try(v2:"Result",loop:true) */
                        Result = Core.F_print_any(s2.Id())
                        /* ERROR PROTECTION INSERTED (Result-Result) */
                        if !ErrorIn(Result) {
                        PRINC("*/")
                        F_Generate_breakline_void()
                        PRINC("")
                        Result = EVOID
                        }}}
                        } 
                      }
                      } 
                    } 
                  } else {
                  var g0205I *ClaireBoolean  
                  var try_11 EID 
                  /*g_try(v2:"try_11",loop:false) */
                  if (x.Isa.IsIn(Language.C_Call_array) == CTRUE) { 
                    { var arg_12 *ClaireAny  
                      _ = arg_12
                      var try_13 EID 
                      /*g_try(v2:"try_13",loop:false) */
                      { var arg_14 *ClaireType  
                        _ = arg_14
                        var try_15 EID 
                        /*g_try(v2:"try_15",loop:false) */
                        { var arg_16 *ClaireType  
                          _ = arg_16
                          var try_17 EID 
                          /*g_try(v2:"try_17",loop:false) */
                          try_17 = Core.F_CALL(Optimize.C_c_type,ARGS(p.ToEID()))
                          /* ERROR PROTECTION INSERTED (arg_16-try_15) */
                          if ErrorIn(try_17) {try_15 = try_17
                          } else {
                          arg_16 = ToType(OBJ(try_17))
                          try_15 = EID{Core.F_member_type(arg_16).Id(),0}
                          }
                          } 
                        /* ERROR PROTECTION INSERTED (arg_14-try_13) */
                        if ErrorIn(try_15) {try_13 = try_15
                        } else {
                        arg_14 = ToType(OBJ(try_15))
                        try_13 = Core.F_CALL(C_sort_I,ARGS(EID{arg_14.Id(),0}))
                        }
                        } 
                      /* ERROR PROTECTION INSERTED (arg_12-try_11) */
                      if ErrorIn(try_13) {try_11 = try_13
                      } else {
                      arg_12 = ANY(try_13)
                      try_11 = EID{Equal(arg_12,C_any.Id()).Id(),0}
                      }
                      } 
                    } else {
                    try_11 = EID{CFALSE.Id(),0}
                    } 
                  /* ERROR PROTECTION INSERTED (g0205I-Result) */
                  if ErrorIn(try_11) {Result = try_11
                  } else {
                  g0205I = ToBoolean(OBJ(try_11))
                  if (g0205I == CTRUE) { 
                    
                    /*g_try(v2:"Result",loop:true) */
                    { var arg_18 *ClaireAny  
                      _ = arg_18
                      var try_19 EID 
                      /*g_try(v2:"try_19",loop:false) */
                      try_19 = Core.F_CALL(Optimize.C_c_code,ARGS(p.ToEID(),EID{C_array.Id(),0}))
                      /* ERROR PROTECTION INSERTED (arg_18-Result) */
                      if ErrorIn(try_19) {Result = try_19
                      } else {
                      arg_18 = ANY(try_19)
                      Result = Core.F_CALL(C_Generate_g_expression,ARGS(arg_18.ToEID(),EID{C_list.Id(),0}))
                      }
                      } 
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    PRINC(".PutAt(")
                    /*g_try(v2:"Result",loop:true) */
                    Result = Core.F_CALL(C_Generate_g_expression,ARGS(Core.F_CALL(C_arg,ARGS(x.ToEID())),EID{C_integer.Id(),0}))
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    PRINC(",")
                    /*g_try(v2:"Result",loop:true) */
                    Result = Core.F_CALL(C_Generate_g_expression,ARGS(v.ToEID(),EID{C_any.Id(),0}))
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    PRINC(")/*unsorted-x:")
                    /*g_try(v2:"Result",loop:true) */
                    Result = Core.F_print_any(x.Isa.Id())
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    PRINC(",sm:")
                    /*g_try(v2:"Result",loop:true) */
                    { var arg_20 *ClaireAny  
                      _ = arg_20
                      var try_21 EID 
                      /*g_try(v2:"try_21",loop:false) */
                      { var arg_22 *ClaireType  
                        _ = arg_22
                        var try_23 EID 
                        /*g_try(v2:"try_23",loop:false) */
                        { var arg_24 *ClaireType  
                          _ = arg_24
                          var try_25 EID 
                          /*g_try(v2:"try_25",loop:false) */
                          try_25 = Core.F_CALL(Optimize.C_c_type,ARGS(p.ToEID()))
                          /* ERROR PROTECTION INSERTED (arg_24-try_23) */
                          if ErrorIn(try_25) {try_23 = try_25
                          } else {
                          arg_24 = ToType(OBJ(try_25))
                          try_23 = EID{Core.F_member_type(arg_24).Id(),0}
                          }
                          } 
                        /* ERROR PROTECTION INSERTED (arg_22-try_21) */
                        if ErrorIn(try_23) {try_21 = try_23
                        } else {
                        arg_22 = ToType(OBJ(try_23))
                        try_21 = Core.F_CALL(C_sort_I,ARGS(EID{arg_22.Id(),0}))
                        }
                        } 
                      /* ERROR PROTECTION INSERTED (arg_20-Result) */
                      if ErrorIn(try_21) {Result = try_21
                      } else {
                      arg_20 = ANY(try_21)
                      Result = Core.F_print_any(arg_20)
                      }
                      } 
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    PRINC("*/")
                    F_Generate_breakline_void()
                    PRINC("")
                    Result = EVOID
                    }}}}}
                    } else {
                    var g0206I *ClaireBoolean  
                    var try_26 EID 
                    /*g_try(v2:"try_26",loop:false) */
                    if (x.Isa.IsIn(Language.C_Call_table) == CTRUE) { 
                      { 
                        /* Or stat: v="try_26", loop=false */
                        var v_or11 *ClaireBoolean  
                        
                        /* Or stat: try = @ any(sort!((member @ type(c_type(p)))),integer) with try:true, v="try_26", loop=false */
                        var try_27 EID 
                        /*g_try(v2:"try_27",loop:false) */
                        { var arg_28 *ClaireAny  
                          _ = arg_28
                          var try_29 EID 
                          /*g_try(v2:"try_29",loop:false) */
                          { var arg_30 *ClaireType  
                            _ = arg_30
                            var try_31 EID 
                            /*g_try(v2:"try_31",loop:false) */
                            { var arg_32 *ClaireType  
                              _ = arg_32
                              var try_33 EID 
                              /*g_try(v2:"try_33",loop:false) */
                              try_33 = Core.F_CALL(Optimize.C_c_type,ARGS(p.ToEID()))
                              /* ERROR PROTECTION INSERTED (arg_32-try_31) */
                              if ErrorIn(try_33) {try_31 = try_33
                              } else {
                              arg_32 = ToType(OBJ(try_33))
                              try_31 = EID{Core.F_member_type(arg_32).Id(),0}
                              }
                              } 
                            /* ERROR PROTECTION INSERTED (arg_30-try_29) */
                            if ErrorIn(try_31) {try_29 = try_31
                            } else {
                            arg_30 = ToType(OBJ(try_31))
                            try_29 = Core.F_CALL(C_sort_I,ARGS(EID{arg_30.Id(),0}))
                            }
                            } 
                          /* ERROR PROTECTION INSERTED (arg_28-try_27) */
                          if ErrorIn(try_29) {try_27 = try_29
                          } else {
                          arg_28 = ANY(try_29)
                          try_27 = EID{Equal(arg_28,C_integer.Id()).Id(),0}
                          }
                          } 
                        /* ERROR PROTECTION INSERTED (v_or11-try_26) */
                        if ErrorIn(try_27) {try_26 = try_27
                        } else {
                        v_or11 = ToBoolean(OBJ(try_27))
                        if (v_or11 == CTRUE) {try_26 = EID{CTRUE.Id(),0}
                        } else { 
                          /* Or stat: try = @ any(sort!((member @ type(c_type(p)))),float) with try:true, v="try_26", loop=false */
                          var try_34 EID 
                          /*g_try(v2:"try_34",loop:false) */
                          { var arg_35 *ClaireAny  
                            _ = arg_35
                            var try_36 EID 
                            /*g_try(v2:"try_36",loop:false) */
                            { var arg_37 *ClaireType  
                              _ = arg_37
                              var try_38 EID 
                              /*g_try(v2:"try_38",loop:false) */
                              { var arg_39 *ClaireType  
                                _ = arg_39
                                var try_40 EID 
                                /*g_try(v2:"try_40",loop:false) */
                                try_40 = Core.F_CALL(Optimize.C_c_type,ARGS(p.ToEID()))
                                /* ERROR PROTECTION INSERTED (arg_39-try_38) */
                                if ErrorIn(try_40) {try_38 = try_40
                                } else {
                                arg_39 = ToType(OBJ(try_40))
                                try_38 = EID{Core.F_member_type(arg_39).Id(),0}
                                }
                                } 
                              /* ERROR PROTECTION INSERTED (arg_37-try_36) */
                              if ErrorIn(try_38) {try_36 = try_38
                              } else {
                              arg_37 = ToType(OBJ(try_38))
                              try_36 = Core.F_CALL(C_sort_I,ARGS(EID{arg_37.Id(),0}))
                              }
                              } 
                            /* ERROR PROTECTION INSERTED (arg_35-try_34) */
                            if ErrorIn(try_36) {try_34 = try_36
                            } else {
                            arg_35 = ANY(try_36)
                            try_34 = EID{Equal(arg_35,C_float.Id()).Id(),0}
                            }
                            } 
                          /* ERROR PROTECTION INSERTED (v_or11-try_26) */
                          if ErrorIn(try_34) {try_26 = try_34
                          } else {
                          v_or11 = ToBoolean(OBJ(try_34))
                          if (v_or11 == CTRUE) {try_26 = EID{CTRUE.Id(),0}
                          } else { 
                            try_26 = EID{CFALSE.Id(),0}} 
                          } 
                        }}
                        } 
                      } else {
                      try_26 = EID{CFALSE.Id(),0}
                      } 
                    /* ERROR PROTECTION INSERTED (g0206I-Result) */
                    if ErrorIn(try_26) {Result = try_26
                    } else {
                    g0206I = ToBoolean(OBJ(try_26))
                    if (g0206I == CTRUE) { 
                      /*g_try(v2:"Result",loop:true) */
                      Result = Core.F_CALL(C_Generate_g_expression,ARGS(p.ToEID(),EID{C_table.Id(),0}))
                      /* ERROR PROTECTION INSERTED (Result-Result) */
                      if !ErrorIn(Result) {
                      PRINC(".PutAt(")
                      /*g_try(v2:"Result",loop:true) */
                      Result = F_Generate_g_table_index_table(ToTable(p),ANY(Core.F_CALL(C_arg,ARGS(x.ToEID()))))
                      /* ERROR PROTECTION INSERTED (Result-Result) */
                      if !ErrorIn(Result) {
                      PRINC(",")
                      /*g_try(v2:"Result",loop:true) */
                      Result = Core.F_CALL(C_Generate_g_expression,ARGS(v.ToEID(),EID{C_any.Id(),0}))
                      /* ERROR PROTECTION INSERTED (Result-Result) */
                      if !ErrorIn(Result) {
                      PRINC("-1)")
                      F_Generate_breakline_void()
                      PRINC("")
                      Result = EVOID
                      }}}
                      } else {
                      /*g_try(v2:"Result",loop:true) */
                      if (x.Isa.IsIn(Language.C_Call_array) == CTRUE) { 
                        var try_41 EID 
                        /*g_try(v2:"try_41",loop:true) */
                        { var arg_42 *ClaireType  
                          _ = arg_42
                          var try_43 EID 
                          /*g_try(v2:"try_43",loop:false) */
                          { var arg_44 *ClaireType  
                            _ = arg_44
                            var try_45 EID 
                            /*g_try(v2:"try_45",loop:false) */
                            try_45 = Core.F_CALL(Optimize.C_c_type,ARGS(p.ToEID()))
                            /* ERROR PROTECTION INSERTED (arg_44-try_43) */
                            if ErrorIn(try_45) {try_43 = try_45
                            } else {
                            arg_44 = ToType(OBJ(try_45))
                            try_43 = EID{Core.F_member_type(arg_44).Id(),0}
                            }
                            } 
                          /* ERROR PROTECTION INSERTED (arg_42-try_41) */
                          if ErrorIn(try_43) {try_41 = try_43
                          } else {
                          arg_42 = ToType(OBJ(try_43))
                          try_41 = Core.F_CALL(C_sort_I,ARGS(EID{arg_42.Id(),0}))
                          }
                          } 
                        /* ERROR PROTECTION INSERTED (s2-Result) */
                        if ErrorIn(try_41) {Result = try_41
                        } else {
                        s2 = ToClass(OBJ(try_41))
                        Result = EID{s2.Id(),0}
                        if (s2.Id() == C_object.Id()) { 
                          s2 = C_any
                          Result = EID{s2.Id(),0}
                          } else {
                          Result = EID{CFALSE.Id(),0}
                          } 
                        }
                        }  else if (x.Isa.IsIn(Language.C_Call_slot) == CTRUE) { 
                        s2 = F_Generate_rootSlot_slot(ToSlot(OBJ(Core.F_CALL(C_selector,ARGS(x.ToEID()))))).Range.Class_I()
                        Result = EID{s2.Id(),0}
                        } else {
                        Result = EID{CFALSE.Id(),0}
                        } 
                      /* ERROR PROTECTION INSERTED (Result-Result) */
                      if !ErrorIn(Result) {
                      /*g_try(v2:"Result",loop:true) */
                      Result = Core.F_CALL(C_Generate_g_expression,ARGS(x.ToEID(),EID{s2.Id(),0}))
                      /* ERROR PROTECTION INSERTED (Result-Result) */
                      if !ErrorIn(Result) {
                      PRINC(" = ")
                      /*g_try(v2:"Result",loop:true) */
                      Result = Core.F_CALL(C_Generate_g_expression,ARGS(v.ToEID(),EID{s2.Id(),0}))
                      /* ERROR PROTECTION INSERTED (Result-Result) */
                      if !ErrorIn(Result) {
                      F_Generate_breakline_void()
                      PRINC("/*")
                      /*g_try(v2:"Result",loop:true) */
                      Result = Core.F_print_any(s.Id())
                      /* ERROR PROTECTION INSERTED (Result-Result) */
                      if !ErrorIn(Result) {
                      PRINC("->")
                      /*g_try(v2:"Result",loop:true) */
                      Result = Core.F_print_any(s2.Id())
                      /* ERROR PROTECTION INSERTED (Result-Result) */
                      if !ErrorIn(Result) {
                      PRINC("*/")
                      Result = EVOID
                      }}}}
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
  
// The EID go function for: update_statement @ Update (throw: true) 
func E_Generate_update_statement_Update (self EID,s EID) EID { 
    return F_Generate_update_statement_Update(Language.To_Update(OBJ(self)),ToClass(OBJ(s)) )} 
  
// in the expansion of Defarray, we generate x.graph := make_list(29,unknonw) that we need to trap
/* {1} The go function for: need_shortcut(v:any) [status=0] */
func F_Generate_need_shortcut_any (v *ClaireAny ) *ClaireBoolean  { 
    // procedure body with s = boolean 
var Result *ClaireBoolean  
    if (v.Isa.IsIn(Language.C_Call_method) == CTRUE) { 
      { var g0207 *Language.CallMethod   = Language.To_CallMethod(v)
        _ = g0207
        Result = Equal(g0207.Arg.Selector.Id(),C_make_list.Id())
        } 
      } else {
      Result = CFALSE
      } 
    return Result} 
  
// The EID go function for: need_shortcut @ any (throw: false) 
func E_Generate_need_shortcut_any (v EID) EID { 
    return EID{F_Generate_need_shortcut_any(ANY(v) ).Id(),0}} 
  