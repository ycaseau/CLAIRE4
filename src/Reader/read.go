/***** CLAIRE Compilation of file /Users/ycaseau/Dropbox/src/clairev4.0/src/meta/read.cl 
         [version 4.0.02 / safety 5] Saturday 10-30-2021 *****/

package Reader
import (_ "fmt"
	. "Kernel"
	"Core"
	"Language"
)

//-------- dumb function to prevent import errors --------
func import_g0000() { 
    _ = Core.It
    _ = Language.It
    } 
  
  
//+-------------------------------------------------------------+
//| CLAIRE                                                      |
//| reader.cl                                                   |
//| Copyright (C) 1994 - 2021 Yves Caseau. All Rights Reserved  |
//| cf. copyright info in file object.cl: about()               |
//+-------------------------------------------------------------+
// ---------------------------------------------------------------------
// this file contains the reader object and the top-level read functions
// the second part of the reader is in syntax.cl
// ---------------------------------------------------------------------
// **********************************************************************
// *  Content:                                                          *
// *   Part 1: The reader object                                        *
// *   Part 2: reading blocks                                           *
// *   Part 3: reading expressions                                      *
// *   Part 4: miscellaneous                                            *
// **********************************************************************
// **********************************************************************
// *   Part 1: The reader object                                        *
// **********************************************************************
// global definitions
// *arrow*:boolean :: false
// here we define the basic keywords
/* {1} OPT.The go function for: keyword?(x:any) [] */
func F_keyword_ask_any (x *ClaireAny ) *ClaireBoolean  { 
    // use function body compiling 
return  x.Isa.IsIn(C_reserved_keyword)
    } 
  
// The EID go function for: keyword? @ any (throw: false) 
func E_keyword_ask_any (x EID) EID { 
    return EID{/*(sm for keyword? @ any= boolean)*/ F_keyword_ask_any(ANY(x) ).Id(),0}} 
  
// the meta_class of the reader --------------------------------------
// The key values are placed in indexed so that they can be changed (eof ...).
// The slot *internal* is used to give addresses to lexical variables.
// The function next reads a character and places it in the slot first.
//
// **********************************************************************
// *   Part 2: reading blocks                                           *
// **********************************************************************
// these are the two low level functions found in the Kernel - with a direct import pattern :)
// for old historical reasons, the reader code uses integers and the macro pattern #/a
// we could re-write everything using char now that char are natives (rune)
/* {1} OPT.The go function for: next(r:meta_reader) [] */
func (r *MetaReader ) Next ()  { 
    // procedure body with s =  
r.Fromp.GetNext()
    } 
  
// The EID go function for: next @ meta_reader (throw: false) 
func E_next_meta_reader (r EID) EID { 
    /*(sm for next @ meta_reader= void)*/ ToMetaReader(OBJ(r)).Next( )
    return EVOID} 
  
/* {1} OPT.The go function for: firstc(r:meta_reader) [] */
func (r *MetaReader ) Firstc () int { 
    // use function body compiling 
return  r.Fromp.CharInt()
    } 
  
// The EID go function for: firstc @ meta_reader (throw: false) 
func E_firstc_meta_reader (r EID) EID { 
    return EID{C__INT,IVAL(/*(sm for firstc @ meta_reader= integer)*/ ToMetaReader(OBJ(r)).Firstc( ))}} 
  
// when to stop
/* {1} OPT.The go function for: stop?(n:integer) [] */
func F_stop_ask_integer (n int) *ClaireAny  { 
    // use function body compiling 
return  MakeBoolean((n == 44) || (n == 41) || (n == 93) || (n == 125)).Id()
    } 
  
// The EID go function for: stop? @ integer (throw: false) 
func E_stop_ask_integer (n EID) EID { 
    return /*(sm for stop? @ integer= any)*/ F_stop_ask_integer(INT(n) ).ToEID()} 
  
// read the next unit (definition, block or expression)
//
/* {1} OPT.The go function for: nextunit(r:meta_reader) [] */
func (r *MetaReader ) Nextunit () EID { 
    var Result EID 
    /* Let:2 */{ 
      var n int  = r.Skipc()
      /* noccur = 5 */
      if (n == r.Eof) /* If:3 */{ 
        r.Next()
        Result = EID{C_Reader_eof.Id(),0}
        /* If!3 */}  else if (n == 91) /* If:3 */{ 
        /* Let:4 */{ 
          var z *ClaireAny  
          /* noccur = 1 */
          var z_try00045 EID 
          z_try00045 = r.Cnext().Nexte()
          /* ERROR PROTECTION INSERTED (z-Result) */
          if ErrorIn(z_try00045) {Result = z_try00045
          } else {
          z = ANY(z_try00045)
          /* Let:5 */{ 
            var g0005UU *ClaireAny  
            /* noccur = 1 */
            var g0005UU_try00066 EID 
            g0005UU_try00066 = r.Nexte()
            /* ERROR PROTECTION INSERTED (g0005UU-Result) */
            if ErrorIn(g0005UU_try00066) {Result = g0005UU_try00066
            } else {
            g0005UU = ANY(g0005UU_try00066)
            Result = r.Nextdefinition(z,g0005UU,CTRUE)
            }
            /* Let-5 */} 
          }
          /* Let-4 */} 
        /* If!3 */}  else if (n == 40) /* If:3 */{ 
        if (r.Toplevel == CTRUE) /* If:4 */{ 
          Result = r.Nexts(C_none)
          } else {
          /* Let:5 */{ 
            var g0007UU *ClaireAny  
            /* noccur = 1 */
            var g0007UU_try00086 EID 
            g0007UU_try00086 = r.Cnext().Nexte()
            /* ERROR PROTECTION INSERTED (g0007UU-Result) */
            if ErrorIn(g0007UU_try00086) {Result = g0007UU_try00086
            } else {
            g0007UU = ANY(g0007UU_try00086)
            Result = r.Readblock(g0007UU,41)
            }
            /* Let-5 */} 
          /* If-4 */} 
        /* If!3 */}  else if (n == 96) /* If:3 */{ 
        /* Let:4 */{ 
          var g0009UU *ClaireAny  
          /* noccur = 1 */
          var g0009UU_try00105 EID 
          g0009UU_try00105 = r.Cnext().Nextunit()
          /* ERROR PROTECTION INSERTED (g0009UU-Result) */
          if ErrorIn(g0009UU_try00105) {Result = g0009UU_try00105
          } else {
          g0009UU = ANY(g0009UU_try00105)
          Result = Language.C_Quote.Make(g0009UU).ToEID()
          }
          /* Let-4 */} 
        /* If!3 */}  else if (n == 59) /* If:3 */{ 
        for ((r.Firstc() != r.Eof) && 
            (r.Firstc() != 10)) /* while:4 */{ 
          r.Next()
          /* while-4 */} 
        if (r.Firstc() == r.Eof) /* If:4 */{ 
          Result = EID{C_Reader_eof.Id(),0}
          } else {
          ClEnv.NLine = (ClEnv.NLine+1)
          r.Next()
          Result = r.Nextunit()
          /* If-4 */} 
        } else {
        /* Let:4 */{ 
          var x *ClaireAny  
          /* noccur = 14 */
          var x_try00115 EID 
          if (r.Toplevel == CTRUE) /* If:5 */{ 
            x_try00115 = r.Nexts(C_none)
            } else {
            x_try00115 = r.Nextexp(CTRUE)
            /* If-5 */} 
          /* ERROR PROTECTION INSERTED (x-Result) */
          if ErrorIn(x_try00115) {Result = x_try00115
          } else {
          x = ANY(x_try00115)
          var g0012I *ClaireBoolean  
          /* and:5 */{ 
            var v_and5 *ClaireBoolean  
            
            v_and5 = r.Toplevel
            if (v_and5 == CFALSE) {g0012I = CFALSE
            } else /* arg:6 */{ 
              if (x.Isa.IsIn(Language.C_Assign) == CTRUE) /* If:7 */{ 
                /* Let:8 */{ 
                  var g0002 *Language.Assign   = Language.To_Assign(x)
                  /* noccur = 1 */
                  v_and5 = g0002.ClaireVar.Isa.IsIn(Language.C_Vardef)
                  /* Let-8 */} 
                } else {
                v_and5 = CFALSE
                /* If-7 */} 
              if (v_and5 == CFALSE) {g0012I = CFALSE
              } else /* arg:7 */{ 
                g0012I = CTRUE/* arg-7 */} 
              /* arg-6 */} 
            /* and-5 */} 
          if (g0012I == CTRUE) /* If:5 */{ 
            /* Let:6 */{ 
              var _CL_obj *Language.Defobj   = Language.To_Defobj(new(Language.Defobj).Is(Language.C_Defobj))
              /* noccur = 19 */
              _CL_obj.Ident = ToSymbol(OBJ(Core.F_CALL(C_mClaire_pname,ARGS(Core.F_CALL(Language.C_var,ARGS(x.ToEID()))))))
              _CL_obj.Arg = Core.C_global_variable
              /* update:7 */{ 
                var va_arg1 *Language.Definition  
                var va_arg2 *ClaireList  
                va_arg1 = Language.To_Definition(_CL_obj.Id())
                var va_arg2_try00138 EID 
                /* Construct:8 */{ 
                  var v_bag_arg *ClaireAny  
                  va_arg2_try00138= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                  var v_bag_arg_try00149 EID 
                  /* Let:9 */{ 
                    var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                    /* noccur = 5 */
                    _CL_obj.Selector = ToProperty(C__equal.Id())
                    /* update:10 */{ 
                      var va_arg1 *Language.Call  
                      var va_arg2 *ClaireList  
                      va_arg1 = _CL_obj
                      var va_arg2_try001511 EID 
                      /* Construct:11 */{ 
                        var v_bag_arg *ClaireAny  
                        va_arg2_try001511= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                        ToList(OBJ(va_arg2_try001511)).AddFast(C_range.Id())
                        var v_bag_arg_try001612 EID 
                        v_bag_arg_try001612 = Language.F_extract_type_any(ANY(Core.F_CALL(C_range,ARGS(Core.F_CALL(Language.C_var,ARGS(x.ToEID()))))))
                        /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try001511) */
                        if ErrorIn(v_bag_arg_try001612) {va_arg2_try001511 = v_bag_arg_try001612
                        } else {
                        v_bag_arg = ANY(v_bag_arg_try001612)
                        ToList(OBJ(va_arg2_try001511)).AddFast(v_bag_arg)}
                        /* Construct-11 */} 
                      /* ERROR PROTECTION INSERTED (va_arg2-v_bag_arg_try00149) */
                      if ErrorIn(va_arg2_try001511) {v_bag_arg_try00149 = va_arg2_try001511
                      } else {
                      va_arg2 = ToList(OBJ(va_arg2_try001511))
                      /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
                      va_arg1.Args = va_arg2
                      v_bag_arg_try00149 = EID{va_arg2.Id(),0}
                      }
                      /* update-10 */} 
                    /* ERROR PROTECTION INSERTED (v_bag_arg_try00149-v_bag_arg_try00149) */
                    if !ErrorIn(v_bag_arg_try00149) {
                    v_bag_arg_try00149 = EID{_CL_obj.Id(),0}
                    }
                    /* Let-9 */} 
                  /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try00138) */
                  if ErrorIn(v_bag_arg_try00149) {va_arg2_try00138 = v_bag_arg_try00149
                  } else {
                  v_bag_arg = ANY(v_bag_arg_try00149)
                  ToList(OBJ(va_arg2_try00138)).AddFast(v_bag_arg)
                  /* Let:9 */{ 
                    var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                    /* noccur = 5 */
                    _CL_obj.Selector = ToProperty(C__equal.Id())
                    _CL_obj.Args = MakeConstantList(C_value.Id(),ANY(Core.F_CALL(C_arg,ARGS(x.ToEID()))))
                    v_bag_arg = _CL_obj.Id()
                    /* Let-9 */} 
                  ToList(OBJ(va_arg2_try00138)).AddFast(v_bag_arg)}
                  /* Construct-8 */} 
                /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                if ErrorIn(va_arg2_try00138) {Result = va_arg2_try00138
                } else {
                va_arg2 = ToList(OBJ(va_arg2_try00138))
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
            /* If!5 */}  else if (C_string.Id() == x.Isa.Id()) /* If:5 */{ 
            Result = x.ToEID()
            } else {
            var g0017I *ClaireBoolean  
            if (x.Isa.IsIn(Language.C_Call) == CTRUE) /* If:6 */{ 
              /* Let:7 */{ 
                var g0003 *Language.Call   = Language.To_Call(x)
                /* noccur = 2 */
                /* and:8 */{ 
                  var v_and8 *ClaireBoolean  
                  
                  v_and8 = r.SProperties.Contain_ask(g0003.Selector.Id())
                  if (v_and8 == CFALSE) {g0017I = CFALSE
                  } else /* arg:9 */{ 
                    /* Let:10 */{ 
                      var g0018UU *ClaireAny  
                      /* noccur = 1 */
                      /* For:11 */{ 
                        var y *ClaireAny  
                        _ = y
                        g0018UU= CFALSE.Id()
                        var y_support *ClaireList  
                        y_support = g0003.Args
                        y_len := y_support.Length()
                        for i_it := 0; i_it < y_len; i_it++ { 
                          y = y_support.At(i_it)
                          if (y.Isa.IsIn(Language.C_Vardef) == CTRUE) /* If:13 */{ 
                             /*v = g0018UU, s =any*/
g0018UU = CTRUE.Id()
                            break
                            /* If-13 */} 
                          /* loop-12 */} 
                        /* For-11 */} 
                      v_and8 = Core.F_not_any(g0018UU)
                      /* Let-10 */} 
                    if (v_and8 == CFALSE) {g0017I = CFALSE
                    } else /* arg:10 */{ 
                      g0017I = CTRUE/* arg-10 */} 
                    /* arg-9 */} 
                  /* and-8 */} 
                /* Let-7 */} 
              } else {
              g0017I = CFALSE
              /* If-6 */} 
            if (g0017I == CTRUE) /* If:6 */{ 
              /* Let:7 */{ 
                var z *Language.Call   = Language.To_Call(x)
                /* noccur = 5 */
                /* Let:8 */{ 
                  var a *ClaireAny   = z.Args.At(1-1)
                  /* noccur = 3 */
                  if ((z.Selector.Id() == C_begin.Id()) && 
                      (a.Isa.IsIn(C_unbound_symbol) == CTRUE)) /* If:9 */{ 
                    /* Let:10 */{ 
                      var g0019UU *ClaireString  
                      /* noccur = 1 */
                      var g0019UU_try002011 EID 
                      /* Let:11 */{ 
                        var g0021UU *ClaireSymbol  
                        /* noccur = 1 */
                        var g0021UU_try002212 EID 
                        g0021UU_try002212 = Language.F_extract_symbol_any(a)
                        /* ERROR PROTECTION INSERTED (g0021UU-g0019UU_try002011) */
                        if ErrorIn(g0021UU_try002212) {g0019UU_try002011 = g0021UU_try002212
                        } else {
                        g0021UU = ToSymbol(OBJ(g0021UU_try002212))
                        g0019UU_try002011 = EID{g0021UU.String_I().Id(),0}
                        }
                        /* Let-11 */} 
                      /* ERROR PROTECTION INSERTED (g0019UU-Result) */
                      if ErrorIn(g0019UU_try002011) {Result = g0019UU_try002011
                      } else {
                      g0019UU = ToString(OBJ(g0019UU_try002011))
                      Result = ToArray(z.Args.Id()).NthPut(1,(g0019UU).Id()).ToEID()
                      }
                      /* Let-10 */} 
                    } else {
                    Result = EID{CFALSE.Id(),0}
                    /* If-9 */} 
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  if ((z.Selector.Id() == C_end.Id()) && 
                      (a.Isa.IsIn(C_module) == CTRUE)) /* If:9 */{ 
                    ToArray(z.Args.Id()).NthPut(1,C_claire.Id())
                    /* If-9 */} 
                  Result = x.ToEID()
                  }
                  /* Let-8 */} 
                /* Let-7 */} 
              /* If!6 */}  else if ((r.Toplevel != CTRUE) && 
                (x.Isa.IsIn(Language.C_Assert) != CTRUE)) /* If:6 */{ 
              /* Let:7 */{ 
                var g0023UU *ClaireAny  
                /* noccur = 1 */
                var g0023UU_try00248 EID 
                g0023UU_try00248 = r.Nexte()
                /* ERROR PROTECTION INSERTED (g0023UU-Result) */
                if ErrorIn(g0023UU_try00248) {Result = g0023UU_try00248
                } else {
                g0023UU = ANY(g0023UU_try00248)
                Result = r.Nextdefinition(x,g0023UU,CFALSE)
                }
                /* Let-7 */} 
              } else {
              Result = x.ToEID()
              /* If-6 */} 
            /* If-5 */} 
          }
          /* Let-4 */} 
        /* If-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: nextunit @ meta_reader (throw: true) 
func E_nextunit_meta_reader (r EID) EID { 
    return /*(sm for nextunit @ meta_reader= EID)*/ ToMetaReader(OBJ(r)).Nextunit( )} 
  
// read the next statement & stops at the keyword e or at a delimiter
// the keyword has been read but not the delimiter, so we know which case
// by testing stop?(first(r))
// Note: it actually reads a fragment
//
/* {1} OPT.The go function for: nexts(r:meta_reader,e:keyword) [] */
func (r *MetaReader ) Nexts (e *ClaireKeyword ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var n int  = r.Skipc()
      /* noccur = 2 */
      if (n == r.Eof) /* If:3 */{ 
        r.Next()
        Result = EID{C_Reader_eof.Id(),0}
        /* If!3 */}  else if (n == 91) /* If:3 */{ 
        /* Let:4 */{ 
          var z *ClaireAny  
          /* noccur = 1 */
          var z_try00265 EID 
          z_try00265 = r.Cnext().Nexte()
          /* ERROR PROTECTION INSERTED (z-Result) */
          if ErrorIn(z_try00265) {Result = z_try00265
          } else {
          z = ANY(z_try00265)
          /* Let:5 */{ 
            var g0027UU *ClaireAny  
            /* noccur = 1 */
            var g0027UU_try00286 EID 
            g0027UU_try00286 = r.Nexte()
            /* ERROR PROTECTION INSERTED (g0027UU-Result) */
            if ErrorIn(g0027UU_try00286) {Result = g0027UU_try00286
            } else {
            g0027UU = ANY(g0027UU_try00286)
            Result = r.Nextdefinition(z,g0027UU,CTRUE)
            }
            /* Let-5 */} 
          }
          /* Let-4 */} 
        /* If!3 */}  else if (e.Id() == C_None.Id()) /* If:3 */{ 
        Result = r.Nexte()
        } else {
        /* Let:4 */{ 
          var x *ClaireAny  
          /* noccur = 6 */
          var x_try00295 EID 
          x_try00295 = r.Nexte()
          /* ERROR PROTECTION INSERTED (x-Result) */
          if ErrorIn(x_try00295) {Result = x_try00295
          } else {
          x = ANY(x_try00295)
          if (F_keyword_ask_any(x) == CTRUE) /* If:5 */{ 
            Result = r.Nextstruct(ToKeyword(x),e)
            } else {
            var g0030I *ClaireBoolean  
            /* and:6 */{ 
              var v_and6 *ClaireBoolean  
              
              v_and6 = Equal(e.Id(),C_none.Id())
              if (v_and6 == CFALSE) {g0030I = CFALSE
              } else /* arg:7 */{ 
                if (x.Isa.IsIn(Language.C_Call) == CTRUE) /* If:8 */{ 
                  /* Let:9 */{ 
                    var g0025 *Language.Call   = Language.To_Call(x)
                    /* noccur = 1 */
                    v_and6 = r.SProperties.Contain_ask(g0025.Selector.Id())
                    /* Let-9 */} 
                  } else {
                  v_and6 = CFALSE
                  /* If-8 */} 
                if (v_and6 == CFALSE) {g0030I = CFALSE
                } else /* arg:8 */{ 
                  g0030I = CTRUE/* arg-8 */} 
                /* arg-7 */} 
              /* and-6 */} 
            if (g0030I == CTRUE) /* If:6 */{ 
              Result = x.ToEID()
              } else {
              Result = r.Loopexp(x,e,CFALSE)
              /* If-6 */} 
            /* If-5 */} 
          }
          /* Let-4 */} 
        /* If-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: nexts @ meta_reader (throw: true) 
func E_nexts_meta_reader (r EID,e EID) EID { 
    return /*(sm for nexts @ meta_reader= EID)*/ ToMetaReader(OBJ(r)).Nexts(ToKeyword(OBJ(e)) )} 
  
// loops until the right expression is built (ends with e ',', '}' or ')')
// x is the first expression that was read
/* {1} OPT.The go function for: loopexp(r:meta_reader,x:any,e:keyword,loop:boolean) [] */
func (r *MetaReader ) Loopexp (x *ClaireAny ,e *ClaireKeyword ,loop *ClaireBoolean ) EID { 
    var Result EID 
    if ((r.Toplevel == CTRUE) && 
        ((e.Id() == C_none.Id()) && 
          (r.Findeol() == CTRUE))) /* If:2 */{ 
      Result = x.ToEID()
      /* If!2 */}  else if (x == C__ask.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
        /* noccur = 5 */
        _CL_obj.Selector = C_inspect
        /* update:4 */{ 
          var va_arg1 *Language.Call  
          var va_arg2 *ClaireList  
          va_arg1 = _CL_obj
          var va_arg2_try00315 EID 
          /* Construct:5 */{ 
            var v_bag_arg *ClaireAny  
            va_arg2_try00315= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
            var v_bag_arg_try00326 EID 
            v_bag_arg_try00326 = r.Nexte()
            /* ERROR PROTECTION INSERTED (v_bag_arg-va_arg2_try00315) */
            if ErrorIn(v_bag_arg_try00326) {va_arg2_try00315 = v_bag_arg_try00326
            } else {
            v_bag_arg = ANY(v_bag_arg_try00326)
            ToList(OBJ(va_arg2_try00315)).AddFast(v_bag_arg)}
            /* Construct-5 */} 
          /* ERROR PROTECTION INSERTED (va_arg2-Result) */
          if ErrorIn(va_arg2_try00315) {Result = va_arg2_try00315
          } else {
          va_arg2 = ToList(OBJ(va_arg2_try00315))
          /* ---------- now we compile update args(va_arg1) := va_arg2 ------- */
          va_arg1.Args = va_arg2
          Result = EID{va_arg2.Id(),0}
          }
          /* update-4 */} 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        Result = EID{_CL_obj.Id(),0}
        }
        /* Let-3 */} 
      /* If!2 */}  else if (r.Skipc() == 58) /* If:2 */{ 
      /* Let:3 */{ 
        var y *ClaireAny  
        /* noccur = 7 */
        var y_try00334 EID 
        y_try00334 = r.Cnext().Nexte()
        /* ERROR PROTECTION INSERTED (y-Result) */
        if ErrorIn(y_try00334) {Result = y_try00334
        } else {
        y = ANY(y_try00334)
        if (y == C__equal.Id()) /* If:4 */{ 
          /* Let:5 */{ 
            var g0034UU *ClaireAny  
            /* noccur = 1 */
            var g0034UU_try00356 EID 
            /* Let:6 */{ 
              var g0036UU *ClaireAny  
              /* noccur = 1 */
              var g0036UU_try00377 EID 
              g0036UU_try00377 = r.Nexte()
              /* ERROR PROTECTION INSERTED (g0036UU-g0034UU_try00356) */
              if ErrorIn(g0036UU_try00377) {g0034UU_try00356 = g0036UU_try00377
              } else {
              g0036UU = ANY(g0036UU_try00377)
              g0034UU_try00356 = F_combine_any(x,C_L__equal.Id(),g0036UU)
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (g0034UU-Result) */
            if ErrorIn(g0034UU_try00356) {Result = g0034UU_try00356
            } else {
            g0034UU = ANY(g0034UU_try00356)
            Result = r.Loopexp(g0034UU,e,CTRUE)
            }
            /* Let-5 */} 
          /* If!4 */}  else if (y == C_L_.Id()) /* If:4 */{ 
          Result = r.Nextinst(x)
          /* If!4 */}  else if (F_operation_ask_any(y) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var g0038UU *ClaireAny  
            /* noccur = 1 */
            var g0038UU_try00396 EID 
            /* Let:6 */{ 
              var g0040UU *ClaireAny  
              /* noccur = 1 */
              var g0040UU_try00417 EID 
              g0040UU_try00417 = r.Nexte()
              /* ERROR PROTECTION INSERTED (g0040UU-g0038UU_try00396) */
              if ErrorIn(g0040UU_try00417) {g0038UU_try00396 = g0040UU_try00417
              } else {
              g0040UU = ANY(g0040UU_try00417)
              g0038UU_try00396 = r.Loopexp(g0040UU,e,CFALSE)
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (g0038UU-Result) */
            if ErrorIn(g0038UU_try00396) {Result = g0038UU_try00396
            } else {
            g0038UU = ANY(g0038UU_try00396)
            Result = F_extended_operator_property(ToProperty(y),x,g0038UU)
            }
            /* Let-5 */} 
          /* If!4 */}  else if (x.Isa.IsIn(Language.C_Call) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var w *ClaireAny  
            /* noccur = 6 */
            var w_try00426 EID 
            w_try00426 = r.Nexte()
            /* ERROR PROTECTION INSERTED (w-Result) */
            if ErrorIn(w_try00426) {Result = w_try00426
            } else {
            w = ANY(w_try00426)
            if (w == C__equal_sup.Id()) /* If:6 */{ 
              /* update:7 */{ 
                var va_arg1 *MetaReader  
                var va_arg2 *ClaireBoolean  
                va_arg1 = r
                va_arg2 = CTRUE
                /* ---------- now we compile update last_arrow(va_arg1) := va_arg2 ------- */
                va_arg1.LastArrow = va_arg2
                Result = EID{va_arg2.Id(),0}
                /* update-7 */} 
              /* If!6 */}  else if ((Equal(w,C_arrow.Value) != CTRUE) && 
                (w != C_L__equal.Id())) /* If:6 */{ 
              Result = F_Serror_string(MakeString("[149] wrong keyword (~S) after ~S"),MakeConstantList(w,y))
              } else {
              Result = EID{CFALSE.Id(),0}
              /* If-6 */} 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            Result = r.Nextmethod(x,
              y,
              Equal(w,C_L__equal.Id()),
              CFALSE,
              Equal(w,C__equal_sup.Id()))
            }
            }
            /* Let-5 */} 
          } else {
          Result = F_Serror_string(MakeString("[150] Illegal use of :~S after ~S"),MakeConstantList(y,x))
          /* If-4 */} 
        }
        /* Let-3 */} 
      } else {
      /* Let:3 */{ 
        var y *ClaireAny  
        /* noccur = 14 */
        var y_try00434 EID 
        y_try00434 = r.Nexte()
        /* ERROR PROTECTION INSERTED (y-Result) */
        if ErrorIn(y_try00434) {Result = y_try00434
        } else {
        y = ANY(y_try00434)
        if ((y == e.Id()) || 
            ((y == C__equal_sup.Id()) && 
                (e.Id() == C_arrow.Value))) /* If:4 */{ 
          if (y != e.Id()) /* If:5 */{ 
            r.LastArrow = CTRUE
            /* If-5 */} 
          if (F_boolean_I_any(F_stop_ask_integer(r.Firstc())) == CTRUE) /* If:5 */{ 
            Result = F_Serror_string(MakeString("[151] ~S not allowed after ~S\n"),MakeConstantList(MakeChar(F_char_I_integer(r.Firstc())).Id(),e.Id()))
            } else {
            Result = x.ToEID()
            /* If-5 */} 
          /* If!4 */}  else if ((Equal(y,C_triangle.Value) == CTRUE) || 
            ((Equal(y,C_arrow.Value) == CTRUE) || 
              ((y == C_L_.Id()) || 
                ((y == C_L_L_.Id()) || 
                  (y == C__equal_sup.Id()))))) /* If:4 */{ 
          Result = r.Nextdefinition(x,y,CFALSE)
          /* If!4 */}  else if ((y.Isa.IsIn(C_delimiter) == CTRUE) && 
            (F_boolean_I_any(F_stop_ask_integer(r.Firstc())) == CTRUE)) /* If:4 */{ 
          Result = x.ToEID()
          /* If!4 */}  else if (F_operation_ask_any(y) == CTRUE) /* If:4 */{ 
          if (loop == CTRUE) /* If:5 */{ 
            /* Let:6 */{ 
              var g0044UU *ClaireAny  
              /* noccur = 1 */
              var g0044UU_try00457 EID 
              /* Let:7 */{ 
                var g0046UU *ClaireAny  
                /* noccur = 1 */
                var g0046UU_try00478 EID 
                g0046UU_try00478 = r.Nexte()
                /* ERROR PROTECTION INSERTED (g0046UU-g0044UU_try00457) */
                if ErrorIn(g0046UU_try00478) {g0044UU_try00457 = g0046UU_try00478
                } else {
                g0046UU = ANY(g0046UU_try00478)
                g0044UU_try00457 = F_combine_any(x,y,g0046UU)
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (g0044UU-Result) */
              if ErrorIn(g0044UU_try00457) {Result = g0044UU_try00457
              } else {
              g0044UU = ANY(g0044UU_try00457)
              Result = r.Loopexp(g0044UU,e,CTRUE)
              }
              /* Let-6 */} 
            } else {
            /* Let:6 */{ 
              var g0048UU *ClaireAny  
              /* noccur = 1 */
              var g0048UU_try00497 EID 
              /* Let:7 */{ 
                var g0050UU *ClaireAny  
                /* noccur = 1 */
                var g0050UU_try00518 EID 
                g0050UU_try00518 = r.Nexte()
                /* ERROR PROTECTION INSERTED (g0050UU-g0048UU_try00497) */
                if ErrorIn(g0050UU_try00518) {g0048UU_try00497 = g0050UU_try00518
                } else {
                g0050UU = ANY(g0050UU_try00518)
                g0048UU_try00497 = F_combine_I_any(x,y,g0050UU)
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (g0048UU-Result) */
              if ErrorIn(g0048UU_try00497) {Result = g0048UU_try00497
              } else {
              g0048UU = ANY(g0048UU_try00497)
              Result = r.Loopexp(g0048UU,e,CTRUE)
              }
              /* Let-6 */} 
            /* If-5 */} 
          } else {
          Result = F_Serror_string(MakeString("[152] Separation missing between ~S \nand ~S [~S?]"),MakeConstantList(x,y,e.Id()))
          /* If-4 */} 
        }
        /* Let-3 */} 
      /* If-2 */} 
    return Result} 
  
// The EID go function for: loopexp @ meta_reader (throw: true) 
func E_loopexp_meta_reader (r EID,x EID,e EID,loop EID) EID { 
    return /*(sm for loopexp @ meta_reader= EID)*/ ToMetaReader(OBJ(r)).Loopexp(ANY(x),
      ToKeyword(OBJ(e)),
      ToBoolean(OBJ(loop)) )} 
  
// this is the special form for x :op y - new in v3.3.32
/* {1} OPT.The go function for: extended_operator(p:property,x:any,y:any) [] */
func F_extended_operator_property (p *ClaireProperty ,x *ClaireAny ,y *ClaireAny ) EID { 
    var Result EID 
    if (x.Isa.IsIn(Language.C_Call) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0052 *Language.Call   = Language.To_Call(x)
        /* noccur = 8 */
        /* Let:4 */{ 
          var r *ClaireAny  
          /* noccur = 2 */
          if (g0052.Selector.Id() == C_nth.Id()) /* If:5 */{ 
            r = g0052.Args.At(2-1)
            } else {
            r = g0052.Args.At(1-1)
            /* If-5 */} 
          /* Let:5 */{ 
            var v *ClaireVariable  
            /* noccur = 3 */
            /* Let:6 */{ 
              var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
              /* noccur = 3 */
              _CL_obj.Pname = Core.F_gensym_void()
              v = _CL_obj
              /* Let-6 */} 
            /* Let:6 */{ 
              var x2 *Language.Call  
              /* noccur = 2 */
              if (g0052.Selector.Id() == C_nth.Id()) /* If:7 */{ 
                /* Let:8 */{ 
                  var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                  /* noccur = 5 */
                  _CL_obj.Selector = C_nth
                  _CL_obj.Args = MakeConstantList(g0052.Args.At(1-1),v.Id())
                  x2 = _CL_obj
                  /* Let-8 */} 
                } else {
                /* Let:8 */{ 
                  var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                  /* noccur = 5 */
                  _CL_obj.Selector = g0052.Selector
                  _CL_obj.Args = MakeConstantList(v.Id())
                  x2 = _CL_obj
                  /* Let-8 */} 
                /* If-7 */} 
              if (r.Isa.IsIn(Language.C_Call) == CTRUE) /* If:7 */{ 
                /* Let:8 */{ 
                  var _CL_obj *Language.Let   = Language.To_Let(new(Language.Let).Is(Language.C_Let))
                  /* noccur = 6 */
                  _CL_obj.ClaireVar = v
                  _CL_obj.Value = r
                  /* update:9 */{ 
                    var va_arg1 *Language.Let  
                    var va_arg2 *ClaireAny  
                    va_arg1 = _CL_obj
                    var va_arg2_try005410 EID 
                    /* Let:10 */{ 
                      var g0055UU *ClaireAny  
                      /* noccur = 1 */
                      var g0055UU_try005611 EID 
                      g0055UU_try005611 = F_combine_any(x2.Id(),p.Id(),y)
                      /* ERROR PROTECTION INSERTED (g0055UU-va_arg2_try005410) */
                      if ErrorIn(g0055UU_try005611) {va_arg2_try005410 = g0055UU_try005611
                      } else {
                      g0055UU = ANY(g0055UU_try005611)
                      va_arg2_try005410 = F_combine_any(x2.Id(),C_L__equal.Id(),g0055UU)
                      }
                      /* Let-10 */} 
                    /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                    if ErrorIn(va_arg2_try005410) {Result = va_arg2_try005410
                    } else {
                    va_arg2 = ANY(va_arg2_try005410)
                    /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                    va_arg1.Arg = va_arg2
                    Result = va_arg2.ToEID()
                    }
                    /* update-9 */} 
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  Result = EID{_CL_obj.Id(),0}
                  }
                  /* Let-8 */} 
                } else {
                /* Let:8 */{ 
                  var g0057UU *ClaireAny  
                  /* noccur = 1 */
                  var g0057UU_try00589 EID 
                  g0057UU_try00589 = F_combine_any(g0052.Id(),p.Id(),y)
                  /* ERROR PROTECTION INSERTED (g0057UU-Result) */
                  if ErrorIn(g0057UU_try00589) {Result = g0057UU_try00589
                  } else {
                  g0057UU = ANY(g0057UU_try00589)
                  Result = F_combine_any(g0052.Id(),C_L__equal.Id(),g0057UU)
                  }
                  /* Let-8 */} 
                /* If-7 */} 
              /* Let-6 */} 
            /* Let-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      } else {
      /* Let:3 */{ 
        var g0059UU *ClaireAny  
        /* noccur = 1 */
        var g0059UU_try00604 EID 
        g0059UU_try00604 = F_combine_any(x,p.Id(),y)
        /* ERROR PROTECTION INSERTED (g0059UU-Result) */
        if ErrorIn(g0059UU_try00604) {Result = g0059UU_try00604
        } else {
        g0059UU = ANY(g0059UU_try00604)
        Result = F_combine_any(x,C_L__equal.Id(),g0059UU)
        }
        /* Let-3 */} 
      /* If-2 */} 
    return Result} 
  
// The EID go function for: extended_operator @ property (throw: true) 
func E_extended_operator_property (p EID,x EID,y EID) EID { 
    return /*(sm for extended_operator @ property= EID)*/ F_extended_operator_property(ToProperty(OBJ(p)),ANY(x),ANY(y) )} 
  
// **********************************************************************
// *   Part 3: reading expressions                                      *
// **********************************************************************
// reading the next compact expression - comments are ignored but they can
// be attached to the last read expression
/* {1} OPT.The go function for: nexte(r:meta_reader) [] */
func (r *MetaReader ) Nexte () EID { 
    var Result EID 
    /* Let:2 */{ 
      var x *ClaireAny  
      /* noccur = 3 */
      var x_try00613 EID 
      x_try00613 = r.Nextexp(CFALSE)
      /* ERROR PROTECTION INSERTED (x-Result) */
      if ErrorIn(x_try00613) {Result = x_try00613
      } else {
      x = ANY(x_try00613)
      if (x.Isa.IsIn(C_Instruction) == CTRUE) /* If:3 */{ 
        r.LastForm = x
        /* If-3 */} 
      Result = x.ToEID()
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: nexte @ meta_reader (throw: true) 
func E_nexte_meta_reader (r EID) EID { 
    return /*(sm for nexte @ meta_reader= EID)*/ ToMetaReader(OBJ(r)).Nexte( )} 
  
// v3.3
// reading the next compact expression/ same
//
/* {1} OPT.The go function for: nextexp(r:meta_reader,str:boolean) [] */
func (r *MetaReader ) Nextexp (str *ClaireBoolean ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var n int  = r.Skipc()
      /* noccur = 14 */
      if (n == 41) /* If:3 */{ 
        Result = r.Paren.ToEID()
        /* If!3 */}  else if (n == 125) /* If:3 */{ 
        Result = r.Curly.ToEID()
        /* If!3 */}  else if (n == 93) /* If:3 */{ 
        Result = r.Bracket.ToEID()
        /* If!3 */}  else if (n == 124) /* If:3 */{ 
        r.Next()
        Result = EID{C_OR.Value,0}
        /* If!3 */}  else if (n == 44) /* If:3 */{ 
        Result = r.Comma.ToEID()
        /* If!3 */}  else if (n == r.Eof) /* If:3 */{ 
        Result = F_Serror_string(MakeString("[153] eof inside an expression"),CNIL)
        /* If!3 */}  else if (n == 59) /* If:3 */{ 
        for ((r.Firstc() != r.Eof) && 
            (r.Firstc() != 10)) /* while:4 */{ 
          r.Next()
          /* while-4 */} 
        if (r.Firstc() == r.Eof) /* If:4 */{ 
          Result = EID{C_Reader_eof.Id(),0}
          } else {
          ClEnv.NLine = (ClEnv.NLine+1)
          r.Next()
          Result = r.Nexte()
          /* If-4 */} 
        /* If!3 */}  else if (n == 35) /* If:3 */{ 
        Result = r.ReadEscape()
        /* If!3 */}  else if (n == 96) /* If:3 */{ 
        /* Let:4 */{ 
          var g0063UU *ClaireAny  
          /* noccur = 1 */
          var g0063UU_try00645 EID 
          g0063UU_try00645 = r.Cnext().Nexte()
          /* ERROR PROTECTION INSERTED (g0063UU-Result) */
          if ErrorIn(g0063UU_try00645) {Result = g0063UU_try00645
          } else {
          g0063UU = ANY(g0063UU_try00645)
          Result = Language.C_Quote.Make(g0063UU).ToEID()
          }
          /* Let-4 */} 
        } else {
        /* Let:4 */{ 
          var y *ClaireAny   = CNULL
          /* noccur = 8 */
          /* Let:5 */{ 
            var x *ClaireAny  
            /* noccur = 16 */
            var x_try00656 EID 
            if (n == 34) /* If:6 */{ 
              x_try00656 = EID{r.Cnext().Fromp.ReadString().Id(),0}
              /* If!6 */}  else if (n == 40) /* If:6 */{ 
              /* Let:7 */{ 
                var g0066UU *ClaireAny  
                /* noccur = 1 */
                var g0066UU_try00678 EID 
                g0066UU_try00678 = r.Cnext().Nexte()
                /* ERROR PROTECTION INSERTED (g0066UU-x_try00656) */
                if ErrorIn(g0066UU_try00678) {x_try00656 = g0066UU_try00678
                } else {
                g0066UU = ANY(g0066UU_try00678)
                x_try00656 = r.Readblock(g0066UU,41)
                }
                /* Let-7 */} 
              /* If!6 */}  else if ((n >= 48) && 
                (n <= 57)) /* If:6 */{ 
              x_try00656 = r.Fromp.ReadNumber().ToEID()
              /* If!6 */}  else if (n == 123) /* If:6 */{ 
              /* Let:7 */{ 
                var g0068UU *ClaireAny  
                /* noccur = 1 */
                var g0068UU_try00698 EID 
                g0068UU_try00698 = r.Cnext().Nexte()
                /* ERROR PROTECTION INSERTED (g0068UU-x_try00656) */
                if ErrorIn(g0068UU_try00698) {x_try00656 = g0068UU_try00698
                } else {
                g0068UU = ANY(g0068UU_try00698)
                x_try00656 = r.Readset(g0068UU)
                }
                /* Let-7 */} 
              } else {
              var y_try00707 EID 
              y_try00707 = r.Fromp.ReadIdent()
              /* ERROR PROTECTION INSERTED (y-x_try00656) */
              if ErrorIn(y_try00707) {x_try00656 = y_try00707
              } else {
              y = ANY(y_try00707)
              x_try00656 = y.ToEID()
              if (C_string.Id() == y.Isa.Id()) /* If:7 */{ 
                x_try00656 = y.ToEID()
                } else {
                x_try00656 = r.Nexti(y)
                /* If-7 */} 
              }
              /* If-6 */} 
            /* ERROR PROTECTION INSERTED (x-Result) */
            if ErrorIn(x_try00656) {Result = x_try00656
            } else {
            x = ANY(x_try00656)
            if (C_string.Id() == y.Isa.Id()) /* If:6 */{ 
              if (r.ExtendedComment_ask(ToString(y)) == CTRUE) /* If:7 */{ 
                Result = r.ExtendedComment_I(ToString(y))
                /* If!7 */}  else if (str == CTRUE) /* If:7 */{ 
                Result = y.ToEID()
                } else {
                Result = r.Nexte()
                /* If-7 */} 
              } else {
              Result= EID{CFALSE.Id(),0}
              for ((r.Firstc() == 91) || 
                  ((r.Firstc() == 46) || 
                    (r.Firstc() == 60))) /* while:7 */{ 
                var void_try8 EID 
                _ = void_try8
                if (r.Firstc() == 60) /* If:8 */{ 
                  /* Let:9 */{ 
                    var y *ClaireAny  
                    /* noccur = 2 */
                    var y_try007110 EID 
                    y_try007110 = r.Cnext().Nexte()
                    /* ERROR PROTECTION INSERTED (y-void_try8) */
                    if ErrorIn(y_try007110) {void_try8 = y_try007110
                    } else {
                    y = ANY(y_try007110)
                    if ((C_class.Id() == x.Isa.Id()) && 
                        (r.Firstc() == 62)) /* If:10 */{ 
                      r.Cnext()
                      var x_try007211 EID 
                      /* Let:11 */{ 
                        var g0073UU *ClaireList  
                        /* noccur = 1 */
                        /* Construct:12 */{ 
                          var v_bag_arg *ClaireAny  
                          g0073UU= ToType(CEMPTY.Id()).EmptyList()
                          /* Let:13 */{ 
                            var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                            /* noccur = 5 */
                            _CL_obj.Selector = ToProperty(C__equal.Id())
                            _CL_obj.Args = MakeConstantList(C_of.Id(),y)
                            v_bag_arg = _CL_obj.Id()
                            /* Let-13 */} 
                          g0073UU.AddFast(v_bag_arg)/* Construct-12 */} 
                        x_try007211 = Language.F_extract_class_call_class(ToClass(x),g0073UU)
                        /* Let-11 */} 
                      /* ERROR PROTECTION INSERTED (x-void_try8) */
                      if ErrorIn(x_try007211) {void_try8 = x_try007211
                      Result = x_try007211
                      break
                      } else {
                      x = ANY(x_try007211)
                      void_try8 = x.ToEID()
                      var x_try007411 EID 
                      x_try007411 = r.Nexti(x)
                      /* ERROR PROTECTION INSERTED (x-void_try8) */
                      if ErrorIn(x_try007411) {void_try8 = x_try007411
                      Result = x_try007411
                      break
                      } else {
                      x = ANY(x_try007411)
                      void_try8 = x.ToEID()
                      }}
                      } else {
                      void_try8 = F_Serror_string(MakeString("[154] ~S<~S not allowed"),MakeConstantList(x,y))
                      /* If-10 */} 
                    }
                    /* Let-9 */} 
                  /* If!8 */}  else if (r.Firstc() == 91) /* If:8 */{ 
                  /* Let:9 */{ 
                    var l *ClaireAny  
                    /* noccur = 3 */
                    var l_try007510 EID 
                    l_try007510 = r.Cnext().Nextseq(93)
                    /* ERROR PROTECTION INSERTED (l-void_try8) */
                    if ErrorIn(l_try007510) {void_try8 = l_try007510
                    } else {
                    l = ANY(l_try007510)
                    var x_try007610 EID 
                    if ((C_class.Id() == x.Isa.Id()) && 
                        ((x != C_type.Id()) && 
                          (F_boolean_I_any(l) == CTRUE))) /* If:10 */{ 
                      x_try007610 = Language.F_extract_class_call_class(ToClass(x),ToList(l))
                      } else {
                      x_try007610 = F_Call_I_property(C_nth,F_cons_any(x,ToList(l)))
                      /* If-10 */} 
                    /* ERROR PROTECTION INSERTED (x-void_try8) */
                    if ErrorIn(x_try007610) {void_try8 = x_try007610
                    } else {
                    x = ANY(x_try007610)
                    void_try8 = x.ToEID()
                    }
                    }
                    /* Let-9 */} 
                  } else {
                  /* Let:9 */{ 
                    var y *ClaireAny  
                    /* noccur = 1 */
                    var y_try007710 EID 
                    y_try007710 = r.Cnext().Fromp.ReadIdent()
                    /* ERROR PROTECTION INSERTED (y-void_try8) */
                    if ErrorIn(y_try007710) {void_try8 = y_try007710
                    } else {
                    y = ANY(y_try007710)
                    /* Let:10 */{ 
                      var p *ClaireProperty  
                      /* noccur = 2 */
                      var p_try007811 EID 
                      p_try007811 = Language.F_make_a_property_any(y)
                      /* ERROR PROTECTION INSERTED (p-void_try8) */
                      if ErrorIn(p_try007811) {void_try8 = p_try007811
                      } else {
                      p = ToProperty(OBJ(p_try007811))
                      /* Let:11 */{ 
                        var _CL_obj *Language.Call_plus   = Language.To_Call_plus(new(Language.Call_plus).Is(Language.C_Call_plus))
                        /* noccur = 5 */
                        _CL_obj.Selector = p
                        _CL_obj.Args = MakeConstantList(x)
                        x = _CL_obj.Id()
                        /* Let-11 */} 
                      if (p.Reified.Id() == CTRUE.Id()) /* If:11 */{ 
                        /* Let:12 */{ 
                          var _CL_obj *Language.Call   = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                          /* noccur = 5 */
                          _CL_obj.Selector = C_read
                          _CL_obj.Args = MakeConstantList(x)
                          x = _CL_obj.Id()
                          /* Let-12 */} 
                        void_try8 = x.ToEID()
                        } else {
                        void_try8 = EID{CFALSE.Id(),0}
                        /* If-11 */} 
                      }
                      /* Let-10 */} 
                    }
                    /* Let-9 */} 
                  /* If-8 */} 
                /* ERROR PROTECTION INSERTED (void_try8-Result) */
                if ErrorIn(void_try8) {Result = void_try8
                Result = void_try8
                break
                } else {
                /* while-7 */} 
              }
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              Result = x.ToEID()
              }
              /* If-6 */} 
            }
            /* Let-5 */} 
          /* Let-4 */} 
        /* If-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: nextexp @ meta_reader (throw: true) 
func E_nextexp_meta_reader (r EID,str EID) EID { 
    return /*(sm for nextexp @ meta_reader= EID)*/ ToMetaReader(OBJ(r)).Nextexp(ToBoolean(OBJ(str)) )} 
  
// reads a compact expression that starts with an ident
//
/* {1} OPT.The go function for: nexti(r:meta_reader,val:any) [] */
func (r *MetaReader ) Nexti (val *ClaireAny ) EID { 
    var Result EID 
    if (r.Firstc() == 40) /* If:2 */{ 
      if ((val == C_exists.Id()) || 
          ((val == C_forall.Id()) || 
            (val == C_some.Id()))) /* If:3 */{ 
        /* Let:4 */{ 
          var v *ClaireVariable  
          /* noccur = 3 */
          var v_try00925 EID 
          /* Let:5 */{ 
            var g0093UU *ClaireAny  
            /* noccur = 1 */
            var g0093UU_try00946 EID 
            g0093UU_try00946 = r.Cnext().Nexte()
            /* ERROR PROTECTION INSERTED (g0093UU-v_try00925) */
            if ErrorIn(g0093UU_try00946) {v_try00925 = g0093UU_try00946
            } else {
            g0093UU = ANY(g0093UU_try00946)
            v_try00925 = F_extract_variable_any(g0093UU)
            }
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (v-Result) */
          if ErrorIn(v_try00925) {Result = v_try00925
          } else {
          v = To_Variable(OBJ(v_try00925))
          /* Let:5 */{ 
            var _Za2 *ClaireAny  
            /* noccur = 3 */
            var _Za2_try00956 EID 
            _Za2_try00956 = r.Nexte()
            /* ERROR PROTECTION INSERTED (_Za2-Result) */
            if ErrorIn(_Za2_try00956) {Result = _Za2_try00956
            } else {
            _Za2 = ANY(_Za2_try00956)
            /* Let:6 */{ 
              var _Za3 *ClaireAny   = C_any.Id()
              /* noccur = 2 */
              if (_Za2 == C_in.Id()) /* If:7 */{ 
                var _Za3_try00968 EID 
                _Za3_try00968 = r.Nexte()
                /* ERROR PROTECTION INSERTED (_Za3-Result) */
                if ErrorIn(_Za3_try00968) {Result = _Za3_try00968
                } else {
                _Za3 = ANY(_Za3_try00968)
                Result = _Za3.ToEID()
                var g0097I *ClaireBoolean  
                var g0097I_try00988 EID 
                /* Let:8 */{ 
                  var g0099UU *ClaireAny  
                  /* noccur = 1 */
                  var g0099UU_try01009 EID 
                  g0099UU_try01009 = r.Nexte()
                  /* ERROR PROTECTION INSERTED (g0099UU-g0097I_try00988) */
                  if ErrorIn(g0099UU_try01009) {g0097I_try00988 = g0099UU_try01009
                  } else {
                  g0099UU = ANY(g0099UU_try01009)
                  g0097I_try00988 = EID{Core.F__I_equal_any(g0099UU,C_OR.Value).Id(),0}
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (g0097I-Result) */
                if ErrorIn(g0097I_try00988) {Result = g0097I_try00988
                } else {
                g0097I = ToBoolean(OBJ(g0097I_try00988))
                if (g0097I == CTRUE) /* If:8 */{ 
                  Result = F_Serror_string(MakeString("[155] missing | in exists / forall"),CNIL)
                  } else {
                  Result = EID{CFALSE.Id(),0}
                  /* If-8 */} 
                }
                }
                /* If!7 */}  else if (Equal(_Za2,r.Comma) == CTRUE) /* If:7 */{ 
                Result = EID{r.Cnext().Id(),0}
                } else {
                Result = F_Serror_string(MakeString("[156] wrong use of exists(~S ~S ..."),MakeConstantList(v.Id(),_Za2))
                /* If-7 */} 
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              /* Let:7 */{ 
                var _CL_obj *Language.Exists   = Language.To_Exists(new(Language.Exists).Is(Language.C_Exists))
                /* noccur = 7 */
                _CL_obj.ClaireVar = v
                _CL_obj.SetArg = _Za3
                /* update:8 */{ 
                  var va_arg1 *Language.Iteration  
                  var va_arg2 *ClaireAny  
                  va_arg1 = Language.To_Iteration(_CL_obj.Id())
                  var va_arg2_try01019 EID 
                  /* Let:9 */{ 
                    var _Zbind *ClaireList   = r.Bind_I(v)
                    /* noccur = 1 */
                    /* LetE:10 */{ 
                      var x EID 
                      x = F_nexts_I_meta_reader2(r,41)
                      /* ERROR PROTECTION INSERTED (x-va_arg2_try01019) */
                      if ErrorIn(x) {va_arg2_try01019 = x
                      } else {
                      r.Unbind_I(_Zbind)
                      va_arg2_try01019 = x}
                      /* LetE-10 */} 
                    /* Let-9 */} 
                  /* ERROR PROTECTION INSERTED (va_arg2-Result) */
                  if ErrorIn(va_arg2_try01019) {Result = va_arg2_try01019
                  } else {
                  va_arg2 = ANY(va_arg2_try01019)
                  /* ---------- now we compile update arg(va_arg1) := va_arg2 ------- */
                  va_arg1.Arg = va_arg2
                  Result = va_arg2.ToEID()
                  }
                  /* update-8 */} 
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                /* update:8 */{ 
                  var va_arg1 *Language.Exists  
                  var va_arg2 *ClaireAny  
                  va_arg1 = _CL_obj
                  if (val == C_forall.Id()) /* If:9 */{ 
                    va_arg2 = CTRUE.Id()
                    /* If!9 */}  else if (val == C_exists.Id()) /* If:9 */{ 
                    va_arg2 = CFALSE.Id()
                    } else {
                    va_arg2 = CNULL
                    /* If-9 */} 
                  /* ---------- now we compile update iClaire/other(va_arg1) := va_arg2 ------- */
                  va_arg1.Other = va_arg2
                  /* update-8 */} 
                Result = EID{_CL_obj.Id(),0}
                }
                /* Let-7 */} 
              }
              /* Let-6 */} 
            }
            /* Let-5 */} 
          }
          /* Let-4 */} 
        /* If!3 */}  else if (val == C_rule.Id()) /* If:3 */{ 
        r.Cnext()
        Result = val.ToEID()
        } else {
        Result = r.Readcall(val,CNULL)
        /* If-3 */} 
      /* If!2 */}  else if ((val == C_list.Id()) && 
        (r.Firstc() == 123)) /* If:2 */{ 
      /* Let:3 */{ 
        var s *ClaireAny  
        /* noccur = 5 */
        var s_try01024 EID 
        /* Let:4 */{ 
          var g0103UU *ClaireAny  
          /* noccur = 1 */
          var g0103UU_try01045 EID 
          g0103UU_try01045 = r.Cnext().Nexte()
          /* ERROR PROTECTION INSERTED (g0103UU-s_try01024) */
          if ErrorIn(g0103UU_try01045) {s_try01024 = g0103UU_try01045
          } else {
          g0103UU = ANY(g0103UU_try01045)
          s_try01024 = r.Readset(g0103UU)
          }
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (s-Result) */
        if ErrorIn(s_try01024) {Result = s_try01024
        } else {
        s = ANY(s_try01024)
        if (s.Isa.IsIn(Language.C_Image) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var g0079 *Language.Image   = Language.To_Image(s)
            /* noccur = 2 */
            g0079.Isa = Language.C_Collect
            Result = EID{g0079.Id(),0}
            /* Let-5 */} 
          /* If!4 */}  else if (s.Isa.IsIn(Language.C_Select) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var g0080 *Language.Select   = Language.To_Select(s)
            /* noccur = 2 */
            g0080.Isa = Language.C_Lselect
            Result = EID{g0080.Id(),0}
            /* Let-5 */} 
          } else {
          Result = F_Serror_string(MakeString("[157] ~S cannot follow list{"),MakeConstantList(s))
          /* If-4 */} 
        }
        /* Let-3 */} 
      } else {
      var g0105I *ClaireBoolean  
      /* and:3 */{ 
        var v_and3 *ClaireBoolean  
        
        if (val.Isa.IsIn(Language.C_Call) == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var g0082 *Language.Call   = Language.To_Call(val)
            /* noccur = 2 */
            v_and3 = MakeBoolean((g0082.Selector.Id() == C_nth.Id()) && (g0082.Args.At(1-1) == C_list.Id()))
            /* Let-5 */} 
          } else {
          v_and3 = CFALSE
          /* If-4 */} 
        if (v_and3 == CFALSE) {g0105I = CFALSE
        } else /* arg:4 */{ 
          v_and3 = Equal(MakeInteger(r.Firstc()).Id(),MakeInteger(123).Id())
          if (v_and3 == CFALSE) {g0105I = CFALSE
          } else /* arg:5 */{ 
            g0105I = CTRUE/* arg-5 */} 
          /* arg-4 */} 
        /* and-3 */} 
      if (g0105I == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var s *ClaireAny  
          /* noccur = 5 */
          var s_try01065 EID 
          /* Let:5 */{ 
            var g0107UU *ClaireAny  
            /* noccur = 1 */
            var g0107UU_try01086 EID 
            g0107UU_try01086 = r.Cnext().Nexte()
            /* ERROR PROTECTION INSERTED (g0107UU-s_try01065) */
            if ErrorIn(g0107UU_try01086) {s_try01065 = g0107UU_try01086
            } else {
            g0107UU = ANY(g0107UU_try01086)
            s_try01065 = r.Readset(g0107UU)
            }
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (s-Result) */
          if ErrorIn(s_try01065) {Result = s_try01065
          } else {
          s = ANY(s_try01065)
          /* Let:5 */{ 
            var x *ClaireType  
            /* noccur = 2 */
            var x_try01096 EID 
            x_try01096 = F_extract_of_type_Call(Language.To_Call(val))
            /* ERROR PROTECTION INSERTED (x-Result) */
            if ErrorIn(x_try01096) {Result = x_try01096
            } else {
            x = ToType(OBJ(x_try01096))
            if (s.Isa.IsIn(Language.C_Image) == CTRUE) /* If:6 */{ 
              /* Let:7 */{ 
                var g0084 *Language.Image   = Language.To_Image(s)
                /* noccur = 3 */
                g0084.Isa = Language.C_Collect
                g0084.Of = x
                Result = EID{g0084.Id(),0}
                /* Let-7 */} 
              /* If!6 */}  else if (s.Isa.IsIn(Language.C_Select) == CTRUE) /* If:6 */{ 
              /* Let:7 */{ 
                var g0085 *Language.Select   = Language.To_Select(s)
                /* noccur = 3 */
                g0085.Isa = Language.C_Lselect
                g0085.Of = x
                Result = EID{g0085.Id(),0}
                /* Let-7 */} 
              } else {
              Result = F_Serror_string(MakeString("[157] ~S cannot follow list{"),MakeConstantList(s))
              /* If-6 */} 
            }
            /* Let-5 */} 
          }
          /* Let-4 */} 
        } else {
        var g0110I *ClaireBoolean  
        /* and:4 */{ 
          var v_and4 *ClaireBoolean  
          
          if (val.Isa.IsIn(Language.C_Call) == CTRUE) /* If:5 */{ 
            /* Let:6 */{ 
              var g0087 *Language.Call   = Language.To_Call(val)
              /* noccur = 2 */
              v_and4 = MakeBoolean((g0087.Selector.Id() == C_nth.Id()) && (g0087.Args.At(1-1) == C_set.Id()))
              /* Let-6 */} 
            } else {
            v_and4 = CFALSE
            /* If-5 */} 
          if (v_and4 == CFALSE) {g0110I = CFALSE
          } else /* arg:5 */{ 
            v_and4 = Equal(MakeInteger(r.Firstc()).Id(),MakeInteger(123).Id())
            if (v_and4 == CFALSE) {g0110I = CFALSE
            } else /* arg:6 */{ 
              g0110I = CTRUE/* arg-6 */} 
            /* arg-5 */} 
          /* and-4 */} 
        if (g0110I == CTRUE) /* If:4 */{ 
          /* Let:5 */{ 
            var s *ClaireAny  
            /* noccur = 5 */
            var s_try01116 EID 
            /* Let:6 */{ 
              var g0112UU *ClaireAny  
              /* noccur = 1 */
              var g0112UU_try01137 EID 
              g0112UU_try01137 = r.Cnext().Nexte()
              /* ERROR PROTECTION INSERTED (g0112UU-s_try01116) */
              if ErrorIn(g0112UU_try01137) {s_try01116 = g0112UU_try01137
              } else {
              g0112UU = ANY(g0112UU_try01137)
              s_try01116 = r.Readset(g0112UU)
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (s-Result) */
            if ErrorIn(s_try01116) {Result = s_try01116
            } else {
            s = ANY(s_try01116)
            /* Let:6 */{ 
              var x *ClaireType  
              /* noccur = 2 */
              var x_try01147 EID 
              x_try01147 = F_extract_of_type_Call(Language.To_Call(val))
              /* ERROR PROTECTION INSERTED (x-Result) */
              if ErrorIn(x_try01147) {Result = x_try01147
              } else {
              x = ToType(OBJ(x_try01147))
              if (s.Isa.IsIn(Language.C_Image) == CTRUE) /* If:7 */{ 
                /* Let:8 */{ 
                  var g0089 *Language.Image   = Language.To_Image(s)
                  /* noccur = 2 */
                  g0089.Of = x
                  Result = EID{g0089.Id(),0}
                  /* Let-8 */} 
                /* If!7 */}  else if (s.Isa.IsIn(Language.C_Select) == CTRUE) /* If:7 */{ 
                /* Let:8 */{ 
                  var g0090 *Language.Select   = Language.To_Select(s)
                  /* noccur = 2 */
                  g0090.Of = x
                  Result = EID{g0090.Id(),0}
                  /* Let-8 */} 
                } else {
                Result = F_Serror_string(MakeString("[157] ~S cannot follow list{"),MakeConstantList(s))
                /* If-7 */} 
              }
              /* Let-6 */} 
            }
            /* Let-5 */} 
          /* If!4 */}  else if (r.Firstc() == 58) /* If:4 */{ 
          Result = r.Nextvariable(val)
          /* If!4 */}  else if (r.Firstc() == 64) /* If:4 */{ 
          /* Let:5 */{ 
            var _Za1 *ClaireAny  
            /* noccur = 4 */
            var _Za1_try01156 EID 
            _Za1_try01156 = r.Cnext().Fromp.ReadIdent()
            /* ERROR PROTECTION INSERTED (_Za1-Result) */
            if ErrorIn(_Za1_try01156) {Result = _Za1_try01156
            } else {
            _Za1 = ANY(_Za1_try01156)
            if (C_class.Id() != _Za1.Isa.Id()) /* If:6 */{ 
              Result = F_Serror_string(MakeString("[158] wrong type in call ~S@~S"),MakeConstantList(val,_Za1))
              } else {
              Result = EID{CFALSE.Id(),0}
              /* If-6 */} 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            if (r.Firstc() == 40) /* If:6 */{ 
              Result = r.Readcall(val,_Za1)
              } else {
              Result = F_Serror_string(MakeString("[159] missing ( after ~S@~S"),MakeConstantList(val,_Za1))
              /* If-6 */} 
            }
            }
            /* Let-5 */} 
          } else {
          Result = val.ToEID()
          /* If-4 */} 
        /* If-3 */} 
      /* If-2 */} 
    return Result} 
  
// The EID go function for: nexti @ meta_reader (throw: true) 
func E_nexti_meta_reader (r EID,val EID) EID { 
    return /*(sm for nexti @ meta_reader= EID)*/ ToMetaReader(OBJ(r)).Nexti(ANY(val) )} 
  
// we have read the escape character #
//
/* {1} OPT.The go function for: read_escape(r:meta_reader) [] */
func (r *MetaReader ) ReadEscape () EID { 
    var Result EID 
    if (r.Cnext().Firstc() == 47) /* If:2 */{ 
      /* Let:3 */{ 
        var val int  = r.Cnext().Firstc()
        /* noccur = 1 */
        r.Next()
        Result = EID{C__INT,IVAL(val)}
        /* Let-3 */} 
      /* If!2 */}  else if (r.Firstc() == 39) /* If:2 */{ 
      /* Let:3 */{ 
        var g0116UU *ClaireString  
        /* noccur = 1 */
        var g0116UU_try01174 EID 
        /* Let:4 */{ 
          var g0118UU *ClaireSymbol  
          /* noccur = 1 */
          var g0118UU_try01195 EID 
          /* Let:5 */{ 
            var g0120UU *ClaireAny  
            /* noccur = 1 */
            var g0120UU_try01216 EID 
            g0120UU_try01216 = r.Cnext().Fromp.ReadIdent()
            /* ERROR PROTECTION INSERTED (g0120UU-g0118UU_try01195) */
            if ErrorIn(g0120UU_try01216) {g0118UU_try01195 = g0120UU_try01216
            } else {
            g0120UU = ANY(g0120UU_try01216)
            g0118UU_try01195 = Language.F_extract_symbol_any(g0120UU)
            }
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (g0118UU-g0116UU_try01174) */
          if ErrorIn(g0118UU_try01195) {g0116UU_try01174 = g0118UU_try01195
          } else {
          g0118UU = ToSymbol(OBJ(g0118UU_try01195))
          g0116UU_try01174 = EID{g0118UU.String_I().Id(),0}
          }
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (g0116UU-Result) */
        if ErrorIn(g0116UU_try01174) {Result = g0116UU_try01174
        } else {
        g0116UU = ToString(OBJ(g0116UU_try01174))
        Result = F_make_function_string(g0116UU).ToEID()
        }
        /* Let-3 */} 
      /* If!2 */}  else if ((r.Firstc() == 105) && 
        (r.Cnext().Firstc() == 102)) /* If:2 */{ 
      r.Next()
      Result = EID{C_Zif.Id(),0}
      } else {
      Result = F_Serror_string(MakeString("[160] wrong use of special char #"),CNIL)
      /* If-2 */} 
    return Result} 
  
// The EID go function for: read_escape @ meta_reader (throw: true) 
func E_read_escape_meta_reader (r EID) EID { 
    return /*(sm for read_escape @ meta_reader= EID)*/ ToMetaReader(OBJ(r)).ReadEscape( )} 
  
// **********************************************************************
// *   Part 4: miscellaneous                                            *
// **********************************************************************
// in CLAIRE4: we record the use of classes from other modules, to trigger the do_import pragma
// note : this is not enough, hence the pragma may need to be setup manually
// we could add explicit casts ...
/* {1} OPT.The go function for: nextvariable(r:meta_reader,val:any) [] */
func (r *MetaReader ) Nextvariable (val *ClaireAny ) EID { 
    var Result EID 
    if (val == C__inf.Id()) /* If:2 */{ 
      r.Skipc()
      Result = EID{C_triangle.Value,0}
      } else {
      /* Let:3 */{ 
        var x *ClaireAny  
        /* noccur = 1 */
        var x_try01224 EID 
        x_try01224 = r.Cnext().Nexte()
        /* ERROR PROTECTION INSERTED (x-Result) */
        if ErrorIn(x_try01224) {Result = x_try01224
        } else {
        x = ANY(x_try01224)
        /* Let:4 */{ 
          var _CL_obj *Language.Vardef   = Language.To_Vardef(new(Language.Vardef).Is(Language.C_Vardef))
          /* noccur = 5 */
          /* update:5 */{ 
            var va_arg1 *ClaireVariable  
            var va_arg2 *ClaireSymbol  
            va_arg1 = To_Variable(_CL_obj.Id())
            var va_arg2_try01236 EID 
            va_arg2_try01236 = Language.F_extract_symbol_any(val)
            /* ERROR PROTECTION INSERTED (va_arg2-Result) */
            if ErrorIn(va_arg2_try01236) {Result = va_arg2_try01236
            } else {
            va_arg2 = ToSymbol(OBJ(va_arg2_try01236))
            /* ---------- now we compile update mClaire/pname(va_arg1) := va_arg2 ------- */
            va_arg1.Pname = va_arg2
            Result = EID{va_arg2.Id(),0}
            }
            /* update-5 */} 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          _CL_obj.Range = ToType(x)
          Result = EID{_CL_obj.Id(),0}
          }
          /* Let-4 */} 
        }
        /* Let-3 */} 
      /* If-2 */} 
    return Result} 
  
// The EID go function for: nextvariable @ meta_reader (throw: true) 
func E_nextvariable_meta_reader (r EID,val EID) EID { 
    return /*(sm for nextvariable @ meta_reader= EID)*/ ToMetaReader(OBJ(r)).Nextvariable(ANY(val) )} 
  
// reads an expression, then the exact keyword e
//
/* {1} OPT.The go function for: nexts!(r:meta_reader,e:keyword) [] */
func F_nexts_I_meta_reader1 (r *MetaReader ,e *ClaireKeyword ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var x *ClaireAny  
      /* noccur = 2 */
      var x_try01243 EID 
      x_try01243 = r.Nexts(e)
      /* ERROR PROTECTION INSERTED (x-Result) */
      if ErrorIn(x_try01243) {Result = x_try01243
      } else {
      x = ANY(x_try01243)
      if (F_boolean_I_any(F_stop_ask_integer(r.Firstc())).Id() != CTRUE.Id()) /* If:3 */{ 
        Result = x.ToEID()
        } else {
        Result = F_Serror_string(MakeString("[161] Missing keyword ~S after ~S"),MakeConstantList(e.Id(),x))
        /* If-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: nexts! @ list<type_expression>(meta_reader, keyword) (throw: true) 
func E_nexts_I_meta_reader1 (r EID,e EID) EID { 
    return /*(sm for nexts! @ list<type_expression>(meta_reader, keyword)= EID)*/ F_nexts_I_meta_reader1(ToMetaReader(OBJ(r)),ToKeyword(OBJ(e)) )} 
  
// reads an expression, then the exact keyword e
//
/* {1} OPT.The go function for: nexte!(r:meta_reader,e:keyword) [] */
func (r *MetaReader ) Nexte_I (e *ClaireKeyword ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var x *ClaireAny  
      /* noccur = 2 */
      var x_try01253 EID 
      x_try01253 = r.Nexte()
      /* ERROR PROTECTION INSERTED (x-Result) */
      if ErrorIn(x_try01253) {Result = x_try01253
      } else {
      x = ANY(x_try01253)
      var g0126I *ClaireBoolean  
      var g0126I_try01273 EID 
      /* Let:3 */{ 
        var g0128UU *ClaireAny  
        /* noccur = 1 */
        var g0128UU_try01294 EID 
        g0128UU_try01294 = r.Nexte()
        /* ERROR PROTECTION INSERTED (g0128UU-g0126I_try01273) */
        if ErrorIn(g0128UU_try01294) {g0126I_try01273 = g0128UU_try01294
        } else {
        g0128UU = ANY(g0128UU_try01294)
        g0126I_try01273 = EID{Equal(g0128UU,e.Id()).Id(),0}
        }
        /* Let-3 */} 
      /* ERROR PROTECTION INSERTED (g0126I-Result) */
      if ErrorIn(g0126I_try01273) {Result = g0126I_try01273
      } else {
      g0126I = ToBoolean(OBJ(g0126I_try01273))
      if (g0126I == CTRUE) /* If:3 */{ 
        Result = x.ToEID()
        } else {
        Result = F_Serror_string(MakeString("[161] Missing keyword ~S after ~S"),MakeConstantList(e.Id(),x))
        /* If-3 */} 
      }
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: nexte! @ meta_reader (throw: true) 
func E_nexte_I_meta_reader (r EID,e EID) EID { 
    return /*(sm for nexte! @ meta_reader= EID)*/ ToMetaReader(OBJ(r)).Nexte_I(ToKeyword(OBJ(e)) )} 
  
// ... exact separator
/* {1} OPT.The go function for: nexts!(r:meta_reader,e:integer) [] */
func F_nexts_I_meta_reader2 (r *MetaReader ,e int) EID { 
    var Result EID 
    /* Let:2 */{ 
      var x *ClaireAny  
      /* noccur = 2 */
      var x_try01303 EID 
      x_try01303 = r.Nexts(C_none)
      /* ERROR PROTECTION INSERTED (x-Result) */
      if ErrorIn(x_try01303) {Result = x_try01303
      } else {
      x = ANY(x_try01303)
      if (r.Firstc() == e) /* If:3 */{ 
        r.Cnext()
        Result = x.ToEID()
        } else {
        Result = F_Serror_string(MakeString("[162] Missing separator ~S after ~S"),MakeConstantList(MakeChar(F_char_I_integer(e)).Id(),x))
        /* If-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: nexts! @ list<type_expression>(meta_reader, integer) (throw: true) 
func E_nexts_I_meta_reader2 (r EID,e EID) EID { 
    return /*(sm for nexts! @ list<type_expression>(meta_reader, integer)= EID)*/ F_nexts_I_meta_reader2(ToMetaReader(OBJ(r)),INT(e) )} 
  
// ... keyword e or separator n. DOES NOT SKIP the last character
//
/* {1} OPT.The go function for: nexts!(r:meta_reader,e:keyword,n:integer) [] */
func F_nexts_I_meta_reader3 (r *MetaReader ,e *ClaireKeyword ,n int) EID { 
    var Result EID 
    /* Let:2 */{ 
      var x *ClaireAny  
      /* noccur = 2 */
      var x_try01313 EID 
      x_try01313 = r.Nexts(e)
      /* ERROR PROTECTION INSERTED (x-Result) */
      if ErrorIn(x_try01313) {Result = x_try01313
      } else {
      x = ANY(x_try01313)
      if ((r.Firstc() == n) || 
          (F_boolean_I_any(F_stop_ask_integer(r.Firstc())).Id() != CTRUE.Id())) /* If:3 */{ 
        Result = x.ToEID()
        } else {
        Result = F_Serror_string(MakeString("[163] wrong separator ~S after ~S"),MakeConstantList(MakeChar(F_char_I_integer(r.Firstc())).Id(),x))
        /* If-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: nexts! @ list<type_expression>(meta_reader, keyword, integer) (throw: true) 
func E_nexts_I_meta_reader3 (r EID,e EID,n EID) EID { 
    return /*(sm for nexts! @ list<type_expression>(meta_reader, keyword, integer)= EID)*/ F_nexts_I_meta_reader3(ToMetaReader(OBJ(r)),ToKeyword(OBJ(e)),INT(n) )} 
  
// checks if s is an extended comment
//
/* {1} OPT.The go function for: extended_comment?(r:meta_reader,s:string) [] */
func (r *MetaReader ) ExtendedComment_ask (s *ClaireString ) *ClaireBoolean  { 
    // procedure body with s =  
var Result *ClaireBoolean  
    /* Let:2 */{ 
      var n int  = F_get_string(s,']')
      /* noccur = 3 */
      if ((s.At(1) == ToChar(Language.C_EOS.Value).Value) || 
          ((s.At(1) != '[') || 
            (n == 0))) /* If:3 */{ 
        Result = CFALSE
        } else {
        /* Let:4 */{ 
          var g0133UU *ClaireAny  
          /* noccur = 1 */
          /* Let:5 */{ 
            var i int  = 2
            /* noccur = 3 */
            /* Let:6 */{ 
              var g0132 int  = n
              /* noccur = 1 */
              g0133UU= CFALSE.Id()
              for (i <= g0132) /* while:7 */{ 
                if (s.At(n) == '[') /* If:8 */{ 
                   /*v = g0133UU, s =any*/
g0133UU = CTRUE.Id()
                  break
                  /* If-8 */} 
                i = (i+1)
                /* while-7 */} 
              /* Let-6 */} 
            /* Let-5 */} 
          Result = Core.F_not_any(g0133UU)
          /* Let-4 */} 
        /* If-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: extended_comment? @ meta_reader (throw: false) 
func E_extended_comment_ask_meta_reader (r EID,s EID) EID { 
    return EID{/*(sm for extended_comment? @ meta_reader= boolean)*/ ToMetaReader(OBJ(r)).ExtendedComment_ask(ToString(OBJ(s)) ).Id(),0}} 
  
// produce the equivalent extended comment
//
/* {1} OPT.The go function for: extended_comment!(r:meta_reader,s:string) [] */
func (r *MetaReader ) ExtendedComment_I (s *ClaireString ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var i int  = F_get_string(s,']')
      /* noccur = 5 */
      /* Let:3 */{ 
        var k int  = F_included_string(s,MakeString("//"),CTRUE)
        /* noccur = 5 */
        /* Let:4 */{ 
          var m int  = F_length_string(s)
          /* noccur = 10 */
          /* Let:5 */{ 
            var cx int  = r.Firstc()
            /* noccur = 2 */
            Core.F_print_in_string_void()
            for (F_useless_c_integer(F_integer_I_char(s.At(m))) == CTRUE) /* while:6 */{ 
              m = (m-1)
              /* while-6 */} 
            if (s.At(m) == ',') /* If:6 */{ 
              cx = 44
              m = (m-1)
              /* If-6 */} 
            if (k == 0) /* If:6 */{ 
              k = (m+1)
              /* If-6 */} 
            if ((i == 3) && 
                (s.At(i) == '?')) /* If:6 */{ 
              PRINC("assert(")
              /* Let:7 */{ 
                var j int  = (i+2)
                /* noccur = 4 */
                /* Let:8 */{ 
                  var g0135 int  = m
                  /* noccur = 1 */
                  for (j <= g0135) /* while:9 */{ 
                    F_princ_char(s.At(j))
                    j = (j+1)
                    /* while-9 */} 
                  /* Let-8 */} 
                /* Let-7 */} 
              PRINC(")")
              } else {
              PRINC("trace(")
              /* Let:7 */{ 
                var j int  = 2
                /* noccur = 4 */
                /* Let:8 */{ 
                  var g0136 int  = (i-1)
                  /* noccur = 1 */
                  for (j <= g0136) /* while:9 */{ 
                    F_princ_char(s.At(j))
                    j = (j+1)
                    /* while-9 */} 
                  /* Let-8 */} 
                /* Let-7 */} 
              PRINC(",\"")
              /* Let:7 */{ 
                var j int  = (i+2)
                /* noccur = 4 */
                /* Let:8 */{ 
                  var g0137 int  = (k-1)
                  /* noccur = 1 */
                  for (j <= g0137) /* while:9 */{ 
                    F_princ_char(s.At(j))
                    j = (j+1)
                    /* while-9 */} 
                  /* Let-8 */} 
                /* Let-7 */} 
              PRINC("\\n\"")
              if ((k+3) <= m) /* If:7 */{ 
                PRINC(",")
                /* Let:8 */{ 
                  var j int  = (k+3)
                  /* noccur = 4 */
                  /* Let:9 */{ 
                    var g0138 int  = m
                    /* noccur = 1 */
                    for (j <= g0138) /* while:10 */{ 
                      F_princ_char(s.At(j))
                      j = (j+1)
                      /* while-10 */} 
                    /* Let-9 */} 
                  /* Let-8 */} 
                /* If-7 */} 
              PRINC(")")
              /* If-6 */} 
            /* LetE:6 */{ 
              var s2 EID 
              /* Let:7 */{ 
                var g0139UU *ClaireString  
                /* noccur = 1 */
                var g0139UU_try01408 EID 
                g0139UU_try01408 = Core.F_end_of_string_void()
                /* ERROR PROTECTION INSERTED (g0139UU-s2) */
                if ErrorIn(g0139UU_try01408) {s2 = g0139UU_try01408
                } else {
                g0139UU = ToString(OBJ(g0139UU_try01408))
                s2 = F_read_string(g0139UU)
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (s2-Result) */
              if ErrorIn(s2) {Result = s2
              } else {
              ClEnv.NLine = (ClEnv.NLine+1)
              F_pushback_port(C_reader.Fromp,cx)
              Result = s2}
              /* LetE-6 */} 
            /* Let-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: extended_comment! @ meta_reader (throw: true) 
func E_extended_comment_I_meta_reader (r EID,s EID) EID { 
    return /*(sm for extended_comment! @ meta_reader= EID)*/ ToMetaReader(OBJ(r)).ExtendedComment_I(ToString(OBJ(s)) )} 
  