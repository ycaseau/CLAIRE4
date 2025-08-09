/***** CLAIRE Compilation of file /Users/ycaseau/Dropbox/src/clairev4.12/src/meta/read.cl 
         [version 4.1.6 / safety 5] Friday 08-08-2025 07:31:27 *****/

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
/* The go function for: keyword?(x:any) [status=0] */
func F_keyword_ask_any (x *ClaireAny) *ClaireBoolean { 
  return  x.Isa.IsIn(C_reserved_keyword)
  } 

// The EID go function for: keyword? @ any (throw: false) 
func E_keyword_ask_any (x EID) EID { 
  return EID{F_keyword_ask_any(ANY(x) ).Id(),0}} 

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
/* The go function for: next(r:meta_reader) [status=0] */
func (r *MetaReader) Next ()  { 
  r.Fromp.GetNext()
  } 

// The EID go function for: next @ meta_reader (throw: false) 
func E_next_meta_reader (r EID) EID { 
  ToMetaReader(OBJ(r)).Next( )
  return EVOID} 

/* The go function for: firstc(r:meta_reader) [status=0] */
func (r *MetaReader) Firstc () int { 
  return  r.Fromp.CharInt()
  } 

// The EID go function for: firstc @ meta_reader (throw: false) 
func E_firstc_meta_reader (r EID) EID { 
  return EID{C__INT,IVAL(ToMetaReader(OBJ(r)).Firstc( ))}} 

// when to stop
/* The go function for: stop?(n:integer) [status=0] */
func F_stop_ask_integer (n int) *ClaireAny { 
  return  MakeBoolean((n == 44) || (n == 41) || (n == 93) || (n == 125)).Id()
  } 

// The EID go function for: stop? @ integer (throw: false) 
func E_stop_ask_integer (n EID) EID { 
  return F_stop_ask_integer(INT(n) ).ToEID()} 

// read the next unit (definition, block or expression)
//
/* The go function for: nextunit(r:meta_reader) [status=1] */
func (r *MetaReader) Nextunit () EID { 
  var Result EID
  { var n int = r.Skipc()
    if (n == r.Eof) { 
      r.Next()
      Result = EID{C_Reader_eof.Id(),0}
      }  else if (n == 91) { 
      { var z *ClaireAny
        var try_1 EID
        try_1 = r.Cnext().Nexte()
        if ErrorIn(try_1) {Result = try_1
        } else {
        z = ANY(try_1)
        { var arg_2 *ClaireAny
          var try_3 EID
          try_3 = r.Nexte()
          if ErrorIn(try_3) {Result = try_3
          } else {
          arg_2 = ANY(try_3)
          Result = r.Nextdefinition(z,arg_2,CTRUE)
          }
          } 
        }
        } 
      }  else if (n == 40) { 
      if (r.Toplevel == CTRUE) { 
        Result = r.Nexts(C_none)
        } else {
        { var arg_4 *ClaireAny
          var try_5 EID
          try_5 = r.Cnext().Nexte()
          if ErrorIn(try_5) {Result = try_5
          } else {
          arg_4 = ANY(try_5)
          Result = r.ReadList(arg_4)
          }
          } 
        } 
      }  else if (n == 96) { 
      { var arg_6 *ClaireAny
        var try_7 EID
        try_7 = r.Cnext().Nextunit()
        if ErrorIn(try_7) {Result = try_7
        } else {
        arg_6 = ANY(try_7)
        Result = Language.C_Quote.Make(arg_6).ToEID()
        }
        } 
      }  else if (n == 59) { 
      for ((r.Firstc() != r.Eof) && 
          (r.Firstc() != 10)) { 
        r.Next()
        } 
      if (r.Firstc() == r.Eof) { 
        Result = EID{C_Reader_eof.Id(),0}
        } else {
        ClEnv.NLine = (ClEnv.NLine+1)
        r.Next()
        Result = r.Nextunit()
        } 
      } else {
      { var x *ClaireAny
        var try_8 EID
        if (r.Toplevel == CTRUE) { 
          try_8 = r.Nexts(C_none)
          } else {
          try_8 = r.Nextexp(CTRUE)
          } 
        if ErrorIn(try_8) {Result = try_8
        } else {
        x = ANY(try_8)
        var g0004I *ClaireBoolean
        { 
          var v_and4 *ClaireBoolean
          
          v_and4 = r.Toplevel
          if (v_and4 == CFALSE) {g0004I = CFALSE
          } else { 
            if (x.Isa.IsIn(Language.C_Assign) == CTRUE) { 
              { var g0002 *Language.Assign = Language.To_Assign(x)
                v_and4 = g0002.ClaireVar.Isa.IsIn(Language.C_Vardef)
                } 
              } else {
              v_and4 = CFALSE
              } 
            if (v_and4 == CFALSE) {g0004I = CFALSE
            } else { 
              g0004I = CTRUE} 
            } 
          } 
        if (g0004I == CTRUE) { 
          { var _CL_obj *Language.Defobj = Language.To_Defobj(new(Language.Defobj).Is(Language.C_Defobj))
            _CL_obj.Ident = ToSymbol(OBJ(Core.F_CALL(C_mClaire_pname,ARGS(Core.F_CALL(Language.C_var,ARGS(x.ToEID()))))))
            _CL_obj.Arg = Core.C_global_variable
            { 
              var va_arg1 *Language.Definition
              var va_arg2 *ClaireList
              va_arg1 = Language.To_Definition(_CL_obj.Id())
              var try_9 EID
              { 
                var v_bag_arg *ClaireAny
                try_9= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                var try_10 EID
                { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                  _CL_obj.Selector = ToProperty(C__equal.Id())
                  { 
                    var va_arg1 *Language.Call
                    var va_arg2 *ClaireList
                    va_arg1 = _CL_obj
                    var try_11 EID
                    { 
                      var v_bag_arg *ClaireAny
                      try_11= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
                      ToList(OBJ(try_11)).AddFast(C_range.Id())
                      var try_12 EID
                      try_12 = Language.F_extract_type_any(ANY(Core.F_CALL(C_range,ARGS(Core.F_CALL(Language.C_var,ARGS(x.ToEID()))))))
                      if ErrorIn(try_12) {try_11 = try_12
                      } else {
                      v_bag_arg = ANY(try_12)
                      ToList(OBJ(try_11)).AddFast(v_bag_arg)}
                      } 
                    if ErrorIn(try_11) {try_10 = try_11
                    } else {
                    va_arg2 = ToList(OBJ(try_11))
                    va_arg1.Args = va_arg2
                    try_10 = EID{va_arg2.Id(),0}
                    }
                    } 
                  if !ErrorIn(try_10) {
                  try_10 = EID{_CL_obj.Id(),0}
                  }
                  } 
                if ErrorIn(try_10) {try_9 = try_10
                } else {
                v_bag_arg = ANY(try_10)
                ToList(OBJ(try_9)).AddFast(v_bag_arg)
                { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                  _CL_obj.Selector = ToProperty(C__equal.Id())
                  _CL_obj.Args = MakeConstantList(C_value.Id(),ANY(Core.F_CALL(C_arg,ARGS(x.ToEID()))))
                  v_bag_arg = _CL_obj.Id()
                  } 
                ToList(OBJ(try_9)).AddFast(v_bag_arg)}
                } 
              if ErrorIn(try_9) {Result = try_9
              } else {
              va_arg2 = ToList(OBJ(try_9))
              va_arg1.Args = va_arg2
              Result = EID{va_arg2.Id(),0}
              }
              } 
            if !ErrorIn(Result) {
            Result = EID{_CL_obj.Id(),0}
            }
            } 
          }  else if (C_string.Id() == x.Isa.Id()) { 
          Result = x.ToEID()
          } else {
          var g0005I *ClaireBoolean
          if (x.Isa.IsIn(Language.C_Call) == CTRUE) { 
            { var g0003 *Language.Call = Language.To_Call(x)
              { 
                var v_and7 *ClaireBoolean
                
                v_and7 = r.SProperties.Contain_ask(g0003.Selector.Id())
                if (v_and7 == CFALSE) {g0005I = CFALSE
                } else { 
                  { var arg_13 *ClaireAny
                    { 
                      var y *ClaireAny
                      _ = y
                      arg_13= CFALSE.Id()
                      var y_support *ClaireList
                      y_support = g0003.Args
                      y_len := y_support.Length()
                      for i_it := 0; i_it < y_len; i_it++ { 
                        y = y_support.At(i_it)
                        if (y.Isa.IsIn(Language.C_Vardef) == CTRUE) { 
                          arg_13 = CTRUE.Id()
                          break
                          } 
                        } 
                      } 
                    v_and7 = Core.F_not_any(arg_13)
                    } 
                  if (v_and7 == CFALSE) {g0005I = CFALSE
                  } else { 
                    g0005I = CTRUE} 
                  } 
                } 
              } 
            } else {
            g0005I = CFALSE
            } 
          if (g0005I == CTRUE) { 
            { var z *Language.Call = Language.To_Call(x)
              { var a *ClaireAny = z.Args.At(0)
                if ((z.Selector.Id() == C_begin.Id()) && 
                    (a.Isa.IsIn(C_unbound_symbol) == CTRUE)) { 
                  { var arg_14 *ClaireString
                    var try_15 EID
                    { var arg_16 *ClaireSymbol
                      var try_17 EID
                      try_17 = Language.F_extract_symbol_any(a)
                      if ErrorIn(try_17) {try_15 = try_17
                      } else {
                      arg_16 = ToSymbol(OBJ(try_17))
                      try_15 = EID{arg_16.String_I().Id(),0}
                      }
                      } 
                    if ErrorIn(try_15) {Result = try_15
                    } else {
                    arg_14 = ToString(OBJ(try_15))
                    Result = ToArray(z.Args.Id()).NthPut(1,(arg_14).Id()).ToEID()
                    }
                    } 
                  } else {
                  Result = EID{CFALSE.Id(),0}
                  } 
                if !ErrorIn(Result) {
                if ((z.Selector.Id() == C_end.Id()) && 
                    (a.Isa.IsIn(C_module) == CTRUE)) { 
                  ToArray(z.Args.Id()).NthPut(1,C_claire.Id())
                  } 
                Result = x.ToEID()
                }
                } 
              } 
            }  else if ((r.Toplevel != CTRUE) && 
              (x.Isa.IsIn(Language.C_Assert) != CTRUE)) { 
            { var arg_18 *ClaireAny
              var try_19 EID
              try_19 = r.Nexte()
              if ErrorIn(try_19) {Result = try_19
              } else {
              arg_18 = ANY(try_19)
              Result = r.Nextdefinition(x,arg_18,CFALSE)
              }
              } 
            } else {
            Result = x.ToEID()
            } 
          } 
        }
        } 
      } 
    } 
  return Result} 

// The EID go function for: nextunit @ meta_reader (throw: true) 
func E_nextunit_meta_reader (r EID) EID { 
  return ToMetaReader(OBJ(r)).Nextunit( )} 

// read the next statement & stops at the keyword e or at a delimiter
// the keyword has been read but not the delimiter, so we know which case
// by testing stop?(first(r))
// Note: it actually reads a fragment
//
/* The go function for: nexts(r:meta_reader,e:keyword) [status=1] */
func (r *MetaReader) Nexts (e *ClaireKeyword) EID { 
  var Result EID
  { var n int = r.Skipc()
    if (n == r.Eof) { 
      r.Next()
      Result = EID{C_Reader_eof.Id(),0}
      }  else if (n == 91) { 
      { var z *ClaireAny
        var try_1 EID
        try_1 = r.Cnext().Nexte()
        if ErrorIn(try_1) {Result = try_1
        } else {
        z = ANY(try_1)
        { var arg_2 *ClaireAny
          var try_3 EID
          try_3 = r.Nexte()
          if ErrorIn(try_3) {Result = try_3
          } else {
          arg_2 = ANY(try_3)
          Result = r.Nextdefinition(z,arg_2,CTRUE)
          }
          } 
        }
        } 
      }  else if (e.Id() == C_None.Id()) { 
      Result = r.Nexte()
      } else {
      { var x *ClaireAny
        var try_4 EID
        try_4 = r.Nexte()
        if ErrorIn(try_4) {Result = try_4
        } else {
        x = ANY(try_4)
        if (F_keyword_ask_any(x) == CTRUE) { 
          Result = r.Nextstruct(ToKeyword(x),e)
          } else {
          var g0007I *ClaireBoolean
          { 
            var v_and5 *ClaireBoolean
            
            v_and5 = Equal(e.Id(),C_none.Id())
            if (v_and5 == CFALSE) {g0007I = CFALSE
            } else { 
              if (x.Isa.IsIn(Language.C_Call) == CTRUE) { 
                { var g0006 *Language.Call = Language.To_Call(x)
                  v_and5 = r.SProperties.Contain_ask(g0006.Selector.Id())
                  } 
                } else {
                v_and5 = CFALSE
                } 
              if (v_and5 == CFALSE) {g0007I = CFALSE
              } else { 
                g0007I = CTRUE} 
              } 
            } 
          if (g0007I == CTRUE) { 
            Result = x.ToEID()
            } else {
            Result = r.Loopexp(x,e,CFALSE)
            } 
          } 
        }
        } 
      } 
    } 
  return Result} 

// The EID go function for: nexts @ meta_reader (throw: true) 
func E_nexts_meta_reader (r EID,e EID) EID { 
  return ToMetaReader(OBJ(r)).Nexts(ToKeyword(OBJ(e)) )} 

// loops until the right expression is built (ends with e ',', '}' or ')')
// x is the first expression that was read
/* The go function for: loopexp(r:meta_reader,x:any,e:keyword,loop:boolean) [status=1] */
func (r *MetaReader) Loopexp (x *ClaireAny,e *ClaireKeyword,loop *ClaireBoolean) EID { 
  var Result EID
  { var c int = r.Firstc()
    if ((r.Toplevel == CTRUE) && 
        ((e.Id() == C_none.Id()) && 
          (r.Findeol() == CTRUE))) { 
      Result = x.ToEID()
      }  else if (x == C__ask.Id()) { 
      { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
        _CL_obj.Selector = C_inspect
        { 
          var va_arg1 *Language.Call
          var va_arg2 *ClaireList
          va_arg1 = _CL_obj
          var try_1 EID
          { 
            var v_bag_arg *ClaireAny
            try_1= EID{ToType(CEMPTY.Id()).EmptyList().Id(),0}
            var try_2 EID
            try_2 = r.Nexte()
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
      }  else if (r.Skipc() == 58) { 
      { var y *ClaireAny
        var try_3 EID
        try_3 = r.Cnext().Nexte()
        if ErrorIn(try_3) {Result = try_3
        } else {
        y = ANY(try_3)
        if (y == C__equal.Id()) { 
          { var arg_4 *ClaireAny
            var try_5 EID
            { var arg_6 *ClaireAny
              var try_7 EID
              try_7 = r.Nexte()
              if ErrorIn(try_7) {try_5 = try_7
              } else {
              arg_6 = ANY(try_7)
              try_5 = F_combine_any(x,C_L__equal.Id(),arg_6)
              }
              } 
            if ErrorIn(try_5) {Result = try_5
            } else {
            arg_4 = ANY(try_5)
            Result = r.Loopexp(arg_4,e,CTRUE)
            }
            } 
          }  else if (y == C_L_.Id()) { 
          Result = r.Nextinst(x)
          }  else if (F_operation_ask_any(y) == CTRUE) { 
          { var arg_8 *ClaireAny
            var try_9 EID
            { var arg_10 *ClaireAny
              var try_11 EID
              try_11 = r.Nexte()
              if ErrorIn(try_11) {try_9 = try_11
              } else {
              arg_10 = ANY(try_11)
              try_9 = r.Loopexp(arg_10,e,CFALSE)
              }
              } 
            if ErrorIn(try_9) {Result = try_9
            } else {
            arg_8 = ANY(try_9)
            Result = F_extended_operator_property(ToProperty(y),x,arg_8)
            }
            } 
          }  else if ((x.Isa.IsIn(Language.C_Call) == CTRUE) && 
            (c == 32)) { 
          
          { var w *ClaireAny
            var try_12 EID
            try_12 = r.Nexte()
            if ErrorIn(try_12) {Result = try_12
            } else {
            w = ANY(try_12)
            if (w == C__equal_sup.Id()) { 
              { 
                var va_arg1 *MetaReader
                var va_arg2 *ClaireBoolean
                va_arg1 = r
                va_arg2 = CTRUE
                va_arg1.LastArrow = va_arg2
                Result = EID{va_arg2.Id(),0}
                } 
              }  else if ((Equal(w,C_arrow.Value) != CTRUE) && 
                (w != C_L__equal.Id())) { 
              Result = F_Serror_string(MakeString("[149] wrong keyword (~S) after ~S"),MakeConstantList(w,y))
              } else {
              Result = EID{CFALSE.Id(),0}
              } 
            if !ErrorIn(Result) {
            Result = r.Nextmethod(x,
              y,
              Equal(w,C_L__equal.Id()),
              CFALSE,
              Equal(w,C__equal_sup.Id()))
            }
            }
            } 
          } else {
          
          Result = C_pair.Make(x,y).ToEID()
          } 
        }
        } 
      } else {
      { var y *ClaireAny
        var try_13 EID
        try_13 = r.Nexte()
        if ErrorIn(try_13) {Result = try_13
        } else {
        y = ANY(try_13)
        if ((y == e.Id()) || 
            ((y == C__equal_sup.Id()) && 
                (e.Id() == C_arrow.Value))) { 
          if (y != e.Id()) { 
            r.LastArrow = CTRUE
            } 
          if (F_boolean_I_any(F_stop_ask_integer(r.Firstc())) == CTRUE) { 
            Result = F_Serror_string(MakeString("[151] ~S not allowed after ~S\n"),MakeConstantList(MakeChar(F_char_I_integer(r.Firstc())).Id(),e.Id()))
            } else {
            Result = x.ToEID()
            } 
          }  else if ((Equal(y,C_triangle.Value) == CTRUE) || 
            ((Equal(y,C_arrow.Value) == CTRUE) || 
              ((y == C_L_.Id()) || 
                ((y == C_L_L_.Id()) || 
                  (y == C__equal_sup.Id()))))) { 
          Result = r.Nextdefinition(x,y,CFALSE)
          }  else if ((y.Isa.IsIn(C_delimiter) == CTRUE) && 
            (F_boolean_I_any(F_stop_ask_integer(r.Firstc())) == CTRUE)) { 
          Result = x.ToEID()
          }  else if (F_operation_ask_any(y) == CTRUE) { 
          if (loop == CTRUE) { 
            { var arg_14 *ClaireAny
              var try_15 EID
              { var arg_16 *ClaireAny
                var try_17 EID
                try_17 = r.Nexte()
                if ErrorIn(try_17) {try_15 = try_17
                } else {
                arg_16 = ANY(try_17)
                try_15 = F_combine_any(x,y,arg_16)
                }
                } 
              if ErrorIn(try_15) {Result = try_15
              } else {
              arg_14 = ANY(try_15)
              Result = r.Loopexp(arg_14,e,CTRUE)
              }
              } 
            } else {
            { var arg_18 *ClaireAny
              var try_19 EID
              { var arg_20 *ClaireAny
                var try_21 EID
                try_21 = r.Nexte()
                if ErrorIn(try_21) {try_19 = try_21
                } else {
                arg_20 = ANY(try_21)
                try_19 = F_combine_I_any(x,y,arg_20)
                }
                } 
              if ErrorIn(try_19) {Result = try_19
              } else {
              arg_18 = ANY(try_19)
              Result = r.Loopexp(arg_18,e,CTRUE)
              }
              } 
            } 
          } else {
          Result = F_Serror_string(MakeString("[152] Separation missing between ~S \nand ~S [~S?]"),MakeConstantList(x,y,e.Id()))
          } 
        }
        } 
      } 
    } 
  return Result} 

// The EID go function for: loopexp @ meta_reader (throw: true) 
func E_loopexp_meta_reader (r EID,x EID,e EID,loop EID) EID { 
  return ToMetaReader(OBJ(r)).Loopexp(ANY(x),
    ToKeyword(OBJ(e)),
    ToBoolean(OBJ(loop)) )} 

// this is the special form for x :op y - new in v3.3.32
/* The go function for: extended_operator(p:property,x:any,y:any) [status=1] */
func F_extended_operator_property (p *ClaireProperty,x *ClaireAny,y *ClaireAny) EID { 
  var Result EID
  if (x.Isa.IsIn(Language.C_Call) == CTRUE) { 
    { var g0008 *Language.Call = Language.To_Call(x)
      { var r *ClaireAny
        if (g0008.Selector.Id() == C_nth.Id()) { 
          r = g0008.Args.At(1)
          } else {
          r = g0008.Args.At(0)
          } 
        { var v *ClaireVariable
          { var _CL_obj *ClaireVariable = To_Variable(new(ClaireVariable).Is(C_Variable))
            _CL_obj.Pname = Core.F_gensym_void()
            v = _CL_obj
            } 
          { var x2 *Language.Call
            if (g0008.Selector.Id() == C_nth.Id()) { 
              { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                _CL_obj.Selector = C_nth
                _CL_obj.Args = MakeConstantList(g0008.Args.At(0),v.Id())
                x2 = _CL_obj
                } 
              } else {
              { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                _CL_obj.Selector = g0008.Selector
                _CL_obj.Args = MakeConstantList(v.Id())
                x2 = _CL_obj
                } 
              } 
            if (r.Isa.IsIn(Language.C_Call) == CTRUE) { 
              { var _CL_obj *Language.Let = Language.To_Let(new(Language.Let).Is(Language.C_Let))
                _CL_obj.ClaireVar = v
                _CL_obj.Value = r
                { 
                  var va_arg1 *Language.Let
                  var va_arg2 *ClaireAny
                  va_arg1 = _CL_obj
                  var try_1 EID
                  { var arg_2 *ClaireAny
                    var try_3 EID
                    try_3 = F_combine_any(x2.Id(),p.Id(),y)
                    if ErrorIn(try_3) {try_1 = try_3
                    } else {
                    arg_2 = ANY(try_3)
                    try_1 = F_combine_any(x2.Id(),C_L__equal.Id(),arg_2)
                    }
                    } 
                  if ErrorIn(try_1) {Result = try_1
                  } else {
                  va_arg2 = ANY(try_1)
                  va_arg1.Arg = va_arg2
                  Result = va_arg2.ToEID()
                  }
                  } 
                if !ErrorIn(Result) {
                Result = EID{_CL_obj.Id(),0}
                }
                } 
              } else {
              { var arg_4 *ClaireAny
                var try_5 EID
                try_5 = F_combine_any(g0008.Id(),p.Id(),y)
                if ErrorIn(try_5) {Result = try_5
                } else {
                arg_4 = ANY(try_5)
                Result = F_combine_any(g0008.Id(),C_L__equal.Id(),arg_4)
                }
                } 
              } 
            } 
          } 
        } 
      } 
    } else {
    { var arg_6 *ClaireAny
      var try_7 EID
      try_7 = F_combine_any(x,p.Id(),y)
      if ErrorIn(try_7) {Result = try_7
      } else {
      arg_6 = ANY(try_7)
      Result = F_combine_any(x,C_L__equal.Id(),arg_6)
      }
      } 
    } 
  return Result} 

// The EID go function for: extended_operator @ property (throw: true) 
func E_extended_operator_property (p EID,x EID,y EID) EID { 
  return F_extended_operator_property(ToProperty(OBJ(p)),ANY(x),ANY(y) )} 

// **********************************************************************
// *   Part 3: reading expressions                                      *
// **********************************************************************
// reading the next compact expression - comments are ignored but they can
// be attached to the last read expression
/* The go function for: nexte(r:meta_reader) [status=1] */
func (r *MetaReader) Nexte () EID { 
  var Result EID
  { var x *ClaireAny
    var try_1 EID
    try_1 = r.Nextexp(CFALSE)
    if ErrorIn(try_1) {Result = try_1
    } else {
    x = ANY(try_1)
    if (x.Isa.IsIn(C_Instruction) == CTRUE) { 
      r.LastForm = x
      } 
    Result = x.ToEID()
    }
    } 
  return Result} 

// The EID go function for: nexte @ meta_reader (throw: true) 
func E_nexte_meta_reader (r EID) EID { 
  return ToMetaReader(OBJ(r)).Nexte( )} 

// v3.3
// reading the next compact expression/ same
// str = true <=> a comment is not ignored and returned as a string
// nextexp handles the specific cases based on first char, nexti handles what starts with an identifier
/* The go function for: nextexp(r:meta_reader,str:boolean) [status=1] */
func (r *MetaReader) Nextexp (str *ClaireBoolean) EID { 
  var Result EID
  { var n int = r.Skipc()
    if (n == 41) { 
      Result = r.Paren.ToEID()
      }  else if (n == 125) { 
      Result = r.Curly.ToEID()
      }  else if (n == 93) { 
      Result = r.Bracket.ToEID()
      }  else if (n == 124) { 
      r.Next()
      Result = EID{C_OR.Value,0}
      }  else if (n == 44) { 
      Result = r.Comma.ToEID()
      }  else if (n == r.Eof) { 
      Result = F_Serror_string(MakeString("[153] eof inside an expression"),CNIL)
      }  else if (n == 59) { 
      for ((r.Firstc() != r.Eof) && 
          (r.Firstc() != 10)) { 
        r.Next()
        } 
      if (r.Firstc() == r.Eof) { 
        Result = EID{C_Reader_eof.Id(),0}
        } else {
        ClEnv.NLine = (ClEnv.NLine+1)
        r.Next()
        Result = r.Nexte()
        } 
      }  else if (n == 35) { 
      Result = r.ReadEscape()
      }  else if (n == 96) { 
      { var arg_1 *ClaireAny
        var try_2 EID
        try_2 = r.Cnext().Nexte()
        if ErrorIn(try_2) {Result = try_2
        } else {
        arg_1 = ANY(try_2)
        Result = Language.C_Quote.Make(arg_1).ToEID()
        }
        } 
      } else {
      { var y *ClaireAny = CNULL
        { var x *ClaireAny
          var try_3 EID
          if (n == 34) { 
            try_3 = EID{r.Cnext().Fromp.ReadString().Id(),0}
            }  else if (n == 40) { 
            { var arg_4 *ClaireAny
              var try_5 EID
              try_5 = r.Cnext().Nexte()
              if ErrorIn(try_5) {try_3 = try_5
              } else {
              arg_4 = ANY(try_5)
              try_3 = r.ReadList(arg_4)
              }
              } 
            }  else if ((n >= 48) && 
              (n <= 57)) { 
            try_3 = r.Fromp.ReadNumber().ToEID()
            }  else if (n == 123) { 
            { var arg_6 *ClaireAny
              var try_7 EID
              try_7 = r.Cnext().Nexte()
              if ErrorIn(try_7) {try_3 = try_7
              } else {
              arg_6 = ANY(try_7)
              try_3 = r.Readset(arg_6)
              }
              } 
            } else {
            var try_8 EID
            try_8 = r.Fromp.ReadIdent()
            if ErrorIn(try_8) {try_3 = try_8
            } else {
            y = ANY(try_8)
            try_3 = y.ToEID()
            if (C_string.Id() == y.Isa.Id()) { 
              try_3 = y.ToEID()
              } else {
              try_3 = r.Nexti(y)
              } 
            }
            } 
          if ErrorIn(try_3) {Result = try_3
          } else {
          x = ANY(try_3)
          if (C_string.Id() == y.Isa.Id()) { 
            if (r.ExtendedComment_ask(ToString(y)) == CTRUE) { 
              Result = r.ExtendedComment_I(ToString(y))
              }  else if (r.ConditionalComment_ask(ToString(y)) == CTRUE) { 
              Result = r.ConditionalComment_I(ToString(y))
              }  else if (str == CTRUE) { 
              Result = y.ToEID()
              } else {
              Result = r.Nexte()
              } 
            } else {
            Result= EID{CFALSE.Id(),0}
            for ((r.Firstc() == 91) || 
                ((r.Firstc() == 46) || 
                  (r.Firstc() == 60))) { 
              var loop_9 EID
              _ = loop_9
              if (r.Firstc() == 60) { 
                if (x == Language.C_map.Id()) { 
                  var try_10 EID
                  try_10 = r.Readmap()
                  if ErrorIn(try_10) {loop_9 = try_10
                  } else {
                  x = ANY(try_10)
                  loop_9 = x.ToEID()
                  }
                  } else {
                  { var y *ClaireAny
                    var try_11 EID
                    try_11 = r.Cnext().Nexte()
                    if ErrorIn(try_11) {loop_9 = try_11
                    } else {
                    y = ANY(try_11)
                    if ((C_class.Id() == x.Isa.Id()) && 
                        (r.Firstc() == 62)) { 
                      r.Cnext()
                      var try_12 EID
                      { var arg_13 *ClaireList
                        { 
                          var v_bag_arg *ClaireAny
                          arg_13= ToType(CEMPTY.Id()).EmptyList()
                          { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                            _CL_obj.Selector = ToProperty(C__equal.Id())
                            _CL_obj.Args = MakeConstantList(C_of.Id(),y)
                            v_bag_arg = _CL_obj.Id()
                            } 
                          arg_13.AddFast(v_bag_arg)} 
                        try_12 = Language.F_extract_class_call_class(ToClass(x),arg_13)
                        } 
                      if ErrorIn(try_12) {loop_9 = try_12
                      Result = try_12
                      break
                      } else {
                      x = ANY(try_12)
                      loop_9 = x.ToEID()
                      var try_14 EID
                      try_14 = r.Nexti(x)
                      if ErrorIn(try_14) {loop_9 = try_14
                      Result = try_14
                      break
                      } else {
                      x = ANY(try_14)
                      loop_9 = x.ToEID()
                      }}
                      } else {
                      loop_9 = F_Serror_string(MakeString("[154] ~S<~S not allowed"),MakeConstantList(x,y))
                      } 
                    }
                    } 
                  } 
                }  else if (r.Firstc() == 91) { 
                var try_15 EID
                try_15 = r.Readbracket(x)
                if ErrorIn(try_15) {loop_9 = try_15
                } else {
                x = ANY(try_15)
                loop_9 = x.ToEID()
                }
                } else {
                { var y *ClaireAny
                  var try_16 EID
                  try_16 = r.Cnext().Fromp.ReadIdent()
                  if ErrorIn(try_16) {loop_9 = try_16
                  } else {
                  y = ANY(try_16)
                  { var p *ClaireProperty
                    var try_17 EID
                    try_17 = Language.F_make_a_property_any(y)
                    if ErrorIn(try_17) {loop_9 = try_17
                    } else {
                    p = ToProperty(OBJ(try_17))
                    { var _CL_obj *Language.Call_plus = Language.To_Call_plus(new(Language.Call_plus).Is(Language.C_Call_plus))
                      _CL_obj.Selector = p
                      _CL_obj.Args = MakeConstantList(x)
                      x = _CL_obj.Id()
                      } 
                    if (p.Reified.Id() == CTRUE.Id()) { 
                      { var _CL_obj *Language.Call = Language.To_Call(new(Language.Call).Is(Language.C_Call))
                        _CL_obj.Selector = C_read
                        _CL_obj.Args = MakeConstantList(x)
                        x = _CL_obj.Id()
                        } 
                      loop_9 = x.ToEID()
                      } else {
                      loop_9 = EID{CFALSE.Id(),0}
                      } 
                    }
                    } 
                  }
                  } 
                } 
              if ErrorIn(loop_9) {Result = loop_9
              break
              } else {
              } 
            }
            if !ErrorIn(Result) {
            Result = x.ToEID()
            }
            } 
          }
          } 
        } 
      } 
    } 
  return Result} 

// The EID go function for: nextexp @ meta_reader (throw: true) 
func E_nextexp_meta_reader (r EID,str EID) EID { 
  return ToMetaReader(OBJ(r)).Nextexp(ToBoolean(OBJ(str)) )} 

// extended in CLAIRE4: reads the x[y] patterns
/* The go function for: readbracket(r:meta_reader,x:any) [status=1] */
func (r *MetaReader) Readbracket (x *ClaireAny) EID { 
  var Result EID
  { var l *ClaireAny
    var try_1 EID
    try_1 = r.Cnext().Nextseq(93)
    if ErrorIn(try_1) {Result = try_1
    } else {
    l = ANY(try_1)
    if ((C_class.Id() == x.Isa.Id()) && 
        ((x != C_type.Id()) && 
          (F_boolean_I_any(l) == CTRUE))) { 
      Result = Language.F_extract_class_call_class(ToClass(x),ToList(l))
      } else {
      Result = F_Call_I_property(C_nth,F_cons_any(x,ToList(l)))
      } 
    }
    } 
  return Result} 

// The EID go function for: readbracket @ meta_reader (throw: true) 
func E_Reader_readbracket_meta_reader (r EID,x EID) EID { 
  return ToMetaReader(OBJ(r)).Readbracket(ANY(x) )} 

// new in CLAIRE4: reads map<t1,t2>(pairs*)
/* The go function for: readmap(r:meta_reader) [status=1] */
func (r *MetaReader) Readmap () EID { 
  var Result EID
  
  { var l1 *ClaireAny
    var try_1 EID
    try_1 = r.Cnext().Nextseq(62)
    if ErrorIn(try_1) {Result = try_1
    } else {
    l1 = ANY(try_1)
    
    if !ANY(Core.F_CALL(C_length,ARGS(l1.ToEID()))).IsInt(2) { 
      Result = F_Serror_string(MakeString("[XXX] map<~A requires two types"),MakeConstantList(l1))
      } else {
      Result = EID{CFALSE.Id(),0}
      } 
    if !ErrorIn(Result) {
    { var l2 *ClaireAny
      var try_2 EID
      try_2 = r.Cnext().Nextseq(41)
      if ErrorIn(try_2) {Result = try_2
      } else {
      l2 = ANY(try_2)
      { var m *Language.Map
        var try_3 EID
        { var _CL_obj *Language.Map = Language.To_Map(new(Language.Map).Is(Language.C_Map))
          { 
            var va_arg1 *Language.Map
            var va_arg2 *ClaireType
            va_arg1 = _CL_obj
            var try_4 EID
            { var arg_5 *ClaireAny
              var try_6 EID
              try_6 = Core.F_CALL(C_nth,ARGS(l1.ToEID(),EID{C__INT,IVAL(1)}))
              if ErrorIn(try_6) {try_4 = try_6
              } else {
              arg_5 = ANY(try_6)
              try_4 = Language.F_extract_type_any(arg_5)
              }
              } 
            if ErrorIn(try_4) {try_3 = try_4
            } else {
            va_arg2 = ToType(OBJ(try_4))
            va_arg1.Domain = va_arg2
            try_3 = EID{va_arg2.Id(),0}
            }
            } 
          if !ErrorIn(try_3) {
          { 
            var va_arg1 *Language.Map
            var va_arg2 *ClaireType
            va_arg1 = _CL_obj
            var try_7 EID
            try_7 = Core.F_CALL(C_nth,ARGS(l1.ToEID(),EID{C__INT,IVAL(2)}))
            if ErrorIn(try_7) {try_3 = try_7
            } else {
            va_arg2 = ToType(OBJ(try_7))
            va_arg1.Of = va_arg2
            try_3 = EID{va_arg2.Id(),0}
            }
            } 
          if !ErrorIn(try_3) {
          try_3 = EID{_CL_obj.Id(),0}
          }}
          } 
        if ErrorIn(try_3) {Result = try_3
        } else {
        m = Language.To_Map(OBJ(try_3))
        { 
          var x *ClaireAny
          _ = x
          Result= EID{CFALSE.Id(),0}
          var x_support *ClaireList
          var try_8 EID
          try_8 = Core.F_enumerate_any(l2)
          if ErrorIn(try_8) {Result = try_8
          } else {
          x_support = ToList(OBJ(try_8))
          x_len := x_support.Length()
          for i_it := 0; i_it < x_len; i_it++ { 
            x = x_support.At(i_it)
            var loop_9 EID
            _ = loop_9
            if (x.Isa.IsIn(C_pair) == CTRUE) { 
              { var g0013 *ClairePair = ToPair(x)
                { 
                  var va_arg1 *Language.Construct
                  var va_arg2 *ClaireList
                  va_arg1 = Language.To_Construct(m.Id())
                  va_arg2 = m.Args.AddFast(g0013.Id())
                  va_arg1.Args = va_arg2
                  loop_9 = EID{va_arg2.Id(),0}
                  } 
                } 
              }  else if (x.Isa.IsIn(Language.C_Vardef) == CTRUE) { 
              { var g0014 *Language.Vardef = Language.To_Vardef(x)
                { 
                  var va_arg1 *Language.Construct
                  var va_arg2 *ClaireList
                  va_arg1 = Language.To_Construct(m.Id())
                  va_arg2 = m.Args.AddFast(C_pair.Make(F_Reader_revVar_Vardef(g0014),g0014.Range.Id()))
                  va_arg1.Args = va_arg2
                  loop_9 = EID{va_arg2.Id(),0}
                  } 
                } 
              } else {
              loop_9 = F_Serror_string(MakeString("~S in map<~A>(... is not a pair "),MakeConstantList(x,l1))
              } 
            if ErrorIn(loop_9) {Result = loop_9
            break
            } else {
            }}
            } 
          } 
        if !ErrorIn(Result) {
        Result = EID{m.Id(),0}
        }
        }
        } 
      }
      } 
    }
    }
    } 
  return Result} 

// The EID go function for: readmap @ meta_reader (throw: true) 
func E_Reader_readmap_meta_reader (r EID) EID { 
  return ToMetaReader(OBJ(r)).Readmap( )} 

// returns to the original form from which the Vardef was created (name -> symbol)
/* The go function for: revVar(x:Vardef) [status=0] */
func F_Reader_revVar_Vardef (x *Language.Vardef) *ClaireAny { 
  var Result *ClaireAny
  { var s *ClaireSymbol = x.Pname
    { var v *ClaireAny = s.Value()
      if (v == CNULL) { 
        { var _CL_obj *ClaireUnboundSymbol = ToUnboundSymbol(new(ClaireUnboundSymbol).Is(C_unbound_symbol))
          _CL_obj.Name = s
          Result = _CL_obj.Id()
          } 
        } else {
        Result = v
        } 
      } 
    } 
  return Result} 

// The EID go function for: revVar @ Vardef (throw: false) 
func E_Reader_revVar_Vardef (x EID) EID { 
  return F_Reader_revVar_Vardef(Language.To_Vardef(OBJ(x)) ).ToEID()} 

// reads a compact expression that starts with an ident
/* The go function for: nexti(r:meta_reader,val:any) [status=1] */
func (r *MetaReader) Nexti (val *ClaireAny) EID { 
  var Result EID
  if (r.Firstc() == 40) { 
    if ((val == C_forall.Id()) || 
        ((val == C_exists.Id()) || 
          (val == C_some.Id()))) { 
      { var v *ClaireVariable
        var try_1 EID
        { var arg_2 *ClaireAny
          var try_3 EID
          try_3 = r.Cnext().Nexte()
          if ErrorIn(try_3) {try_1 = try_3
          } else {
          arg_2 = ANY(try_3)
          try_1 = F_extract_variable_any(arg_2)
          }
          } 
        if ErrorIn(try_1) {Result = try_1
        } else {
        v = To_Variable(OBJ(try_1))
        { var _Za2 *ClaireAny
          var try_4 EID
          try_4 = r.Nexte()
          if ErrorIn(try_4) {Result = try_4
          } else {
          _Za2 = ANY(try_4)
          { var _Za3 *ClaireAny = C_any.Id()
            if (_Za2 == C_in.Id()) { 
              var try_5 EID
              try_5 = r.Nexte()
              if ErrorIn(try_5) {Result = try_5
              } else {
              _Za3 = ANY(try_5)
              Result = _Za3.ToEID()
              var g0029I *ClaireBoolean
              var try_6 EID
              { var arg_7 *ClaireAny
                var try_8 EID
                try_8 = r.Nexte()
                if ErrorIn(try_8) {try_6 = try_8
                } else {
                arg_7 = ANY(try_8)
                try_6 = EID{Core.F__I_equal_any(arg_7,C_OR.Value).Id(),0}
                }
                } 
              if ErrorIn(try_6) {Result = try_6
              } else {
              g0029I = ToBoolean(OBJ(try_6))
              if (g0029I == CTRUE) { 
                Result = F_Serror_string(MakeString("[155] missing | in exists / forall"),CNIL)
                } else {
                Result = EID{CFALSE.Id(),0}
                } 
              }
              }
              }  else if (Equal(_Za2,r.Comma) == CTRUE) { 
              Result = EID{r.Cnext().Id(),0}
              } else {
              Result = F_Serror_string(MakeString("[156] wrong use of exists(~S ~S ..."),MakeConstantList(v.Id(),_Za2))
              } 
            if !ErrorIn(Result) {
            { var _CL_obj *Language.Exists = Language.To_Exists(new(Language.Exists).Is(Language.C_Exists))
              _CL_obj.ClaireVar = v
              _CL_obj.SetArg = _Za3
              { 
                var va_arg1 *Language.Iteration
                var va_arg2 *ClaireAny
                va_arg1 = Language.To_Iteration(_CL_obj.Id())
                var try_9 EID
                { var _Zbind *ClaireList = r.Bind_I(v)
                  { 
                    var x EID
                    x = F_nexts_I_meta_reader2(r,41)
                    if ErrorIn(x) {try_9 = x
                    } else {
                    r.Unbind_I(_Zbind)
                    try_9 = x}
                    } 
                  } 
                if ErrorIn(try_9) {Result = try_9
                } else {
                va_arg2 = ANY(try_9)
                va_arg1.Arg = va_arg2
                Result = va_arg2.ToEID()
                }
                } 
              if !ErrorIn(Result) {
              { 
                var va_arg1 *Language.Exists
                var va_arg2 *ClaireAny
                va_arg1 = _CL_obj
                if (val == C_forall.Id()) { 
                  va_arg2 = CTRUE.Id()
                  }  else if (val == C_exists.Id()) { 
                  va_arg2 = CFALSE.Id()
                  } else {
                  va_arg2 = CNULL
                  } 
                va_arg1.Other = va_arg2
                } 
              Result = EID{_CL_obj.Id(),0}
              }
              } 
            }
            } 
          }
          } 
        }
        } 
      }  else if (val == C_rule.Id()) { 
      r.Cnext()
      Result = val.ToEID()
      } else {
      Result = r.Readcall(val,CNULL)
      } 
    }  else if ((val == C_list.Id()) && 
      (r.Firstc() == 123)) { 
    { var s *ClaireAny
      var try_10 EID
      { var arg_11 *ClaireAny
        var try_12 EID
        try_12 = r.Cnext().Nexte()
        if ErrorIn(try_12) {try_10 = try_12
        } else {
        arg_11 = ANY(try_12)
        try_10 = r.Readset(arg_11)
        }
        } 
      if ErrorIn(try_10) {Result = try_10
      } else {
      s = ANY(try_10)
      if (s.Isa.IsIn(Language.C_Image) == CTRUE) { 
        { var g0016 *Language.Image = Language.To_Image(s)
          g0016.Isa = Language.C_Collect
          Result = EID{g0016.Id(),0}
          } 
        }  else if (s.Isa.IsIn(Language.C_Select) == CTRUE) { 
        { var g0017 *Language.Select = Language.To_Select(s)
          g0017.Isa = Language.C_Lselect
          Result = EID{g0017.Id(),0}
          } 
        } else {
        Result = F_Serror_string(MakeString("[157] ~S cannot follow list{"),MakeConstantList(s))
        } 
      }
      } 
    } else {
    var g0030I *ClaireBoolean
    { 
      var v_and2 *ClaireBoolean
      
      if (val.Isa.IsIn(Language.C_Call) == CTRUE) { 
        { var g0019 *Language.Call = Language.To_Call(val)
          v_and2 = MakeBoolean((g0019.Selector.Id() == C_nth.Id()) && (g0019.Args.At(0) == C_list.Id()))
          } 
        } else {
        v_and2 = CFALSE
        } 
      if (v_and2 == CFALSE) {g0030I = CFALSE
      } else { 
        v_and2 = Equal(MakeInteger(r.Firstc()).Id(),MakeInteger(123).Id())
        if (v_and2 == CFALSE) {g0030I = CFALSE
        } else { 
          g0030I = CTRUE} 
        } 
      } 
    if (g0030I == CTRUE) { 
      { var s *ClaireAny
        var try_13 EID
        { var arg_14 *ClaireAny
          var try_15 EID
          try_15 = r.Cnext().Nexte()
          if ErrorIn(try_15) {try_13 = try_15
          } else {
          arg_14 = ANY(try_15)
          try_13 = r.Readset(arg_14)
          }
          } 
        if ErrorIn(try_13) {Result = try_13
        } else {
        s = ANY(try_13)
        { var x *ClaireType
          var try_16 EID
          try_16 = F_extract_of_type_Call(Language.To_Call(val))
          if ErrorIn(try_16) {Result = try_16
          } else {
          x = ToType(OBJ(try_16))
          if (s.Isa.IsIn(Language.C_Image) == CTRUE) { 
            { var g0021 *Language.Image = Language.To_Image(s)
              g0021.Isa = Language.C_Collect
              g0021.Of = x
              Result = EID{g0021.Id(),0}
              } 
            }  else if (s.Isa.IsIn(Language.C_Select) == CTRUE) { 
            { var g0022 *Language.Select = Language.To_Select(s)
              g0022.Isa = Language.C_Lselect
              g0022.Of = x
              Result = EID{g0022.Id(),0}
              } 
            } else {
            Result = F_Serror_string(MakeString("[157] ~S cannot follow list{"),MakeConstantList(s))
            } 
          }
          } 
        }
        } 
      } else {
      var g0031I *ClaireBoolean
      { 
        var v_and3 *ClaireBoolean
        
        if (val.Isa.IsIn(Language.C_Call) == CTRUE) { 
          { var g0024 *Language.Call = Language.To_Call(val)
            v_and3 = MakeBoolean((g0024.Selector.Id() == C_nth.Id()) && (g0024.Args.At(0) == C_set.Id()))
            } 
          } else {
          v_and3 = CFALSE
          } 
        if (v_and3 == CFALSE) {g0031I = CFALSE
        } else { 
          v_and3 = Equal(MakeInteger(r.Firstc()).Id(),MakeInteger(123).Id())
          if (v_and3 == CFALSE) {g0031I = CFALSE
          } else { 
            g0031I = CTRUE} 
          } 
        } 
      if (g0031I == CTRUE) { 
        { var s *ClaireAny
          var try_17 EID
          { var arg_18 *ClaireAny
            var try_19 EID
            try_19 = r.Cnext().Nexte()
            if ErrorIn(try_19) {try_17 = try_19
            } else {
            arg_18 = ANY(try_19)
            try_17 = r.Readset(arg_18)
            }
            } 
          if ErrorIn(try_17) {Result = try_17
          } else {
          s = ANY(try_17)
          { var x *ClaireType
            var try_20 EID
            try_20 = F_extract_of_type_Call(Language.To_Call(val))
            if ErrorIn(try_20) {Result = try_20
            } else {
            x = ToType(OBJ(try_20))
            if (s.Isa.IsIn(Language.C_Image) == CTRUE) { 
              { var g0026 *Language.Image = Language.To_Image(s)
                g0026.Of = x
                Result = EID{g0026.Id(),0}
                } 
              }  else if (s.Isa.IsIn(Language.C_Select) == CTRUE) { 
              { var g0027 *Language.Select = Language.To_Select(s)
                g0027.Of = x
                Result = EID{g0027.Id(),0}
                } 
              } else {
              Result = F_Serror_string(MakeString("[157] ~S cannot follow list{"),MakeConstantList(s))
              } 
            }
            } 
          }
          } 
        }  else if (r.Firstc() == 58) { 
        Result = r.Nextvariable(val)
        }  else if (r.Firstc() == 64) { 
        { var _Za1 *ClaireAny
          var try_21 EID
          try_21 = r.Cnext().Fromp.ReadIdent()
          if ErrorIn(try_21) {Result = try_21
          } else {
          _Za1 = ANY(try_21)
          if (C_class.Id() != _Za1.Isa.Id()) { 
            Result = F_Serror_string(MakeString("[158] wrong type in call ~S@~S"),MakeConstantList(val,_Za1))
            } else {
            Result = EID{CFALSE.Id(),0}
            } 
          if !ErrorIn(Result) {
          if (r.Firstc() == 40) { 
            Result = r.Readcall(val,_Za1)
            } else {
            Result = F_Serror_string(MakeString("[159] missing ( after ~S@~S"),MakeConstantList(val,_Za1))
            } 
          }
          }
          } 
        } else {
        Result = val.ToEID()
        } 
      } 
    } 
  return Result} 

// The EID go function for: nexti @ meta_reader (throw: true) 
func E_nexti_meta_reader (r EID,val EID) EID { 
  return ToMetaReader(OBJ(r)).Nexti(ANY(val) )} 

// we have read the escape character #
//
/* The go function for: read_escape(r:meta_reader) [status=1] */
func (r *MetaReader) ReadEscape () EID { 
  var Result EID
  if (r.Cnext().Firstc() == 47) { 
    { var val int = r.Cnext().Firstc()
      r.Next()
      Result = EID{C__INT,IVAL(val)}
      } 
    }  else if (r.Firstc() == 39) { 
    { var arg_1 *ClaireString
      var try_2 EID
      { var arg_3 *ClaireSymbol
        var try_4 EID
        { var arg_5 *ClaireAny
          var try_6 EID
          try_6 = r.Cnext().Fromp.ReadIdent()
          if ErrorIn(try_6) {try_4 = try_6
          } else {
          arg_5 = ANY(try_6)
          try_4 = Language.F_extract_symbol_any(arg_5)
          }
          } 
        if ErrorIn(try_4) {try_2 = try_4
        } else {
        arg_3 = ToSymbol(OBJ(try_4))
        try_2 = EID{arg_3.String_I().Id(),0}
        }
        } 
      if ErrorIn(try_2) {Result = try_2
      } else {
      arg_1 = ToString(OBJ(try_2))
      Result = F_make_function_string(arg_1).ToEID()
      }
      } 
    }  else if ((r.Firstc() == 105) && 
      (r.Cnext().Firstc() == 102)) { 
    r.Next()
    Result = EID{C_Zif.Id(),0}
    } else {
    Result = F_Serror_string(MakeString("[160] wrong use of special char #"),CNIL)
    } 
  return Result} 

// The EID go function for: read_escape @ meta_reader (throw: true) 
func E_read_escape_meta_reader (r EID) EID { 
  return ToMetaReader(OBJ(r)).ReadEscape( )} 

// **********************************************************************
// *   Part 4: miscellaneous                                            *
// **********************************************************************
// in CLAIRE4: we record the use of classes from other modules, to trigger the do_import pragma
// note : this is not enough, hence the pragma may need to be setup manually
// we could add explicit casts ...
/* The go function for: nextvariable(r:meta_reader,val:any) [status=1] */
func (r *MetaReader) Nextvariable (val *ClaireAny) EID { 
  var Result EID
  if (val == C__inf.Id()) { 
    r.Skipc()
    Result = EID{C_triangle.Value,0}
    } else {
    { var x *ClaireAny
      var try_1 EID
      try_1 = r.Cnext().Nexte()
      if ErrorIn(try_1) {Result = try_1
      } else {
      x = ANY(try_1)
      { var _CL_obj *Language.Vardef = Language.To_Vardef(new(Language.Vardef).Is(Language.C_Vardef))
        { 
          var va_arg1 *ClaireVariable
          var va_arg2 *ClaireSymbol
          va_arg1 = To_Variable(_CL_obj.Id())
          var try_2 EID
          try_2 = Language.F_extract_symbol_any(val)
          if ErrorIn(try_2) {Result = try_2
          } else {
          va_arg2 = ToSymbol(OBJ(try_2))
          va_arg1.Pname = va_arg2
          Result = EID{va_arg2.Id(),0}
          }
          } 
        if !ErrorIn(Result) {
        _CL_obj.Range = ToType(x)
        Result = EID{_CL_obj.Id(),0}
        }
        } 
      }
      } 
    } 
  return Result} 

// The EID go function for: nextvariable @ meta_reader (throw: true) 
func E_nextvariable_meta_reader (r EID,val EID) EID { 
  return ToMetaReader(OBJ(r)).Nextvariable(ANY(val) )} 

// reads an expression, then the exact keyword e
//
/* The go function for: nexts!(r:meta_reader,e:keyword) [status=1] */
func F_nexts_I_meta_reader1 (r *MetaReader,e *ClaireKeyword) EID { 
  var Result EID
  { var x *ClaireAny
    var try_1 EID
    try_1 = r.Nexts(e)
    if ErrorIn(try_1) {Result = try_1
    } else {
    x = ANY(try_1)
    if (F_boolean_I_any(F_stop_ask_integer(r.Firstc())).Id() != CTRUE.Id()) { 
      Result = x.ToEID()
      } else {
      Result = F_Serror_string(MakeString("[161] Missing keyword ~S after ~S"),MakeConstantList(e.Id(),x))
      } 
    }
    } 
  return Result} 

// The EID go function for: nexts! @ list<type_expression>(meta_reader, keyword) (throw: true) 
func E_nexts_I_meta_reader1 (r EID,e EID) EID { 
  return F_nexts_I_meta_reader1(ToMetaReader(OBJ(r)),ToKeyword(OBJ(e)) )} 

// reads an expression, then the exact keyword e
//
/* The go function for: nexte!(r:meta_reader,e:keyword) [status=1] */
func (r *MetaReader) Nexte_I (e *ClaireKeyword) EID { 
  var Result EID
  { var x *ClaireAny
    var try_1 EID
    try_1 = r.Nexte()
    if ErrorIn(try_1) {Result = try_1
    } else {
    x = ANY(try_1)
    var g0032I *ClaireBoolean
    var try_2 EID
    { var arg_3 *ClaireAny
      var try_4 EID
      try_4 = r.Nexte()
      if ErrorIn(try_4) {try_2 = try_4
      } else {
      arg_3 = ANY(try_4)
      try_2 = EID{Equal(arg_3,e.Id()).Id(),0}
      }
      } 
    if ErrorIn(try_2) {Result = try_2
    } else {
    g0032I = ToBoolean(OBJ(try_2))
    if (g0032I == CTRUE) { 
      Result = x.ToEID()
      } else {
      Result = F_Serror_string(MakeString("[161] Missing keyword ~S after ~S"),MakeConstantList(e.Id(),x))
      } 
    }
    }
    } 
  return Result} 

// The EID go function for: nexte! @ meta_reader (throw: true) 
func E_nexte_I_meta_reader (r EID,e EID) EID { 
  return ToMetaReader(OBJ(r)).Nexte_I(ToKeyword(OBJ(e)) )} 

// ... exact separator
/* The go function for: nexts!(r:meta_reader,e:integer) [status=1] */
func F_nexts_I_meta_reader2 (r *MetaReader,e int) EID { 
  var Result EID
  { var x *ClaireAny
    var try_1 EID
    try_1 = r.Nexts(C_none)
    if ErrorIn(try_1) {Result = try_1
    } else {
    x = ANY(try_1)
    if (r.Firstc() == e) { 
      r.Cnext()
      Result = x.ToEID()
      } else {
      Result = F_Serror_string(MakeString("[162] Missing separator ~S after ~S"),MakeConstantList(MakeChar(F_char_I_integer(e)).Id(),x))
      } 
    }
    } 
  return Result} 

// The EID go function for: nexts! @ list<type_expression>(meta_reader, integer) (throw: true) 
func E_nexts_I_meta_reader2 (r EID,e EID) EID { 
  return F_nexts_I_meta_reader2(ToMetaReader(OBJ(r)),INT(e) )} 

// ... keyword e or separator n. DOES NOT SKIP the last character
//
/* The go function for: nexts!(r:meta_reader,e:keyword,n:integer) [status=1] */
func F_nexts_I_meta_reader3 (r *MetaReader,e *ClaireKeyword,n int) EID { 
  var Result EID
  { var x *ClaireAny
    var try_1 EID
    try_1 = r.Nexts(e)
    if ErrorIn(try_1) {Result = try_1
    } else {
    x = ANY(try_1)
    if ((r.Firstc() == n) || 
        (F_boolean_I_any(F_stop_ask_integer(r.Firstc())).Id() != CTRUE.Id())) { 
      Result = x.ToEID()
      } else {
      Result = F_Serror_string(MakeString("[163] wrong separator ~S after ~S"),MakeConstantList(MakeChar(F_char_I_integer(r.Firstc())).Id(),x))
      } 
    }
    } 
  return Result} 

// The EID go function for: nexts! @ list<type_expression>(meta_reader, keyword, integer) (throw: true) 
func E_nexts_I_meta_reader3 (r EID,e EID,n EID) EID { 
  return F_nexts_I_meta_reader3(ToMetaReader(OBJ(r)),ToKeyword(OBJ(e)),INT(n) )} 

// checks if s is an extended comment
//
/* The go function for: extended_comment?(r:meta_reader,s:string) [status=0] */
func (r *MetaReader) ExtendedComment_ask (s *ClaireString) *ClaireBoolean { 
  var Result *ClaireBoolean
  { var n int = F_get_string(s,']')
    if ((s.At(1) == ToChar(Language.C_EOS.Value).Value) || 
        ((s.At(1) != '[') || 
          (n == 0))) { 
      Result = CFALSE
      } else {
      { var arg_1 *ClaireAny
        { var i int = 2
          { var g0033 int = n
            arg_1= CFALSE.Id()
            for (i <= g0033) { 
              if (s.At(n) == '[') { 
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
  return Result} 

// The EID go function for: extended_comment? @ meta_reader (throw: false) 
func E_extended_comment_ask_meta_reader (r EID,s EID) EID { 
  return EID{ToMetaReader(OBJ(r)).ExtendedComment_ask(ToString(OBJ(s)) ).Id(),0}} 

// produce the equivalent extended comment
//
/* The go function for: extended_comment!(r:meta_reader,s:string) [status=1] */
func (r *MetaReader) ExtendedComment_I (s *ClaireString) EID { 
  var Result EID
  { var i int = F_get_string(s,']')
    { var k int = F_included_string(s,MakeString("//"),CTRUE)
      { var m int = F_length_string(s)
        { var cx int = r.Firstc()
          Core.F_print_in_string_void()
          for (F_useless_c_integer(F_integer_I_char(s.At(m))) == CTRUE) { 
            m = (m-1)
            } 
          if (s.At(m) == ',') { 
            cx = 44
            m = (m-1)
            } 
          if (k == 0) { 
            k = (m+1)
            } 
          if ((i == 3) && 
              (s.At(2) == '?')) { 
            PRINC("assert(")
            { var j int = (i+2)
              { var g0035 int = m
                for (j <= g0035) { 
                  F_princ_char(s.At(j))
                  j = (j+1)
                  } 
                } 
              } 
            PRINC(")")
            } else {
            PRINC("trace(")
            { var j int = 2
              { var g0036 int = (i-1)
                for (j <= g0036) { 
                  F_princ_char(s.At(j))
                  j = (j+1)
                  } 
                } 
              } 
            PRINC(",\"")
            { var j int = (i+2)
              { var g0037 int = (k-1)
                for (j <= g0037) { 
                  F_princ_char(s.At(j))
                  j = (j+1)
                  } 
                } 
              } 
            PRINC("\\n\"")
            if ((k+3) <= m) { 
              PRINC(",")
              { var j int = (k+3)
                { var g0038 int = m
                  for (j <= g0038) { 
                    F_princ_char(s.At(j))
                    j = (j+1)
                    } 
                  } 
                } 
              } 
            PRINC(")")
            } 
          { 
            var s2 EID
            { var arg_1 *ClaireString
              var try_2 EID
              try_2 = Core.F_end_of_string_void()
              if ErrorIn(try_2) {s2 = try_2
              } else {
              arg_1 = ToString(OBJ(try_2))
              s2 = F_read_string(arg_1)
              }
              } 
            if ErrorIn(s2) {Result = s2
            } else {
            ClEnv.NLine = (ClEnv.NLine+1)
            F_pushback_port(C_reader.Fromp,cx)
            Result = s2}
            } 
          } 
        } 
      } 
    } 
  return Result} 

// The EID go function for: extended_comment! @ meta_reader (throw: true) 
func E_extended_comment_I_meta_reader (r EID,s EID) EID { 
  return ToMetaReader(OBJ(r)).ExtendedComment_I(ToString(OBJ(s)) )} 

// v4.12 checks if s is a conditional comment //(s) expression
// note : checks the balanced parenthesis
/* The go function for: conditional_comment?(r:meta_reader,s:string) [status=0] */
func (r *MetaReader) ConditionalComment_ask (s *ClaireString) *ClaireBoolean { 
  var Result *ClaireBoolean
  if ((s.At(1) == ToChar(Language.C_EOS.Value).Value) || 
      (s.At(1) != '(')) { 
    Result = CFALSE
    } else {
    { var n int = F_length_string(s)
      { var ct int = 1
        { var i int = 2
          { var g0039 int = n
            Result= CFALSE
            for (i <= g0039) { 
              if (s.At(i) == '(') { 
                ct = (ct+1)
                }  else if (s.At(i) == ')') { 
                ct = (ct-1)
                } 
              if (ct == 0) { 
                Result = CTRUE
                break
                } 
              i = (i+1)
              } 
            } 
          } 
        } 
      } 
    } 
  return Result} 

// The EID go function for: conditional_comment? @ meta_reader (throw: false) 
func E_Reader_conditional_comment_ask_meta_reader2 (r EID,s EID) EID { 
  return EID{ToMetaReader(OBJ(r)).ConditionalComment_ask(ToString(OBJ(s)) ).Id(),0}} 

// v 4.12 produce the equivalent conditional comment
// produces (if (s) expression else assert)  // assert as a marker :)
/* The go function for: conditional_comment!(r:meta_reader,s:string) [status=1] */
func (r *MetaReader) ConditionalComment_I (s *ClaireString) EID { 
  var Result EID
  { var n int = F_length_string(s)
    { var ct int = 1
      { var k int = 0
        { var cx int = r.Firstc()
          { var i int = 2
            { var g0041 int = n
              for (i <= g0041) { 
                if (s.At(i) == '(') { 
                  ct = (ct+1)
                  }  else if (s.At(i) == ')') { 
                  ct = (ct-1)
                  } 
                if (ct == 0) { 
                  k = i
                  
                  break
                  } 
                i = (i+1)
                } 
              } 
            } 
          Core.F_print_in_string_void()
          for (F_useless_c_integer(F_integer_I_char(s.At(n))) == CTRUE) { 
            n = (n-1)
            } 
          if (s.At(n) == ',') { 
            cx = 44
            n = (n-1)
            } 
          PRINC("(if (")
          { var j int = 2
            { var g0042 int = (k-1)
              for (j <= g0042) { 
                F_princ_char(s.At(j))
                j = (j+1)
                } 
              } 
            } 
          PRINC(") ")
          { var j int = (k+1)
            { var g0043 int = n
              for (j <= g0043) { 
                F_princ_char(s.At(j))
                j = (j+1)
                } 
              } 
            } 
          PRINC(" else assert)")
          { 
            var s2 EID
            { var arg_1 *ClaireString
              var try_2 EID
              try_2 = Core.F_end_of_string_void()
              if ErrorIn(try_2) {s2 = try_2
              } else {
              arg_1 = ToString(OBJ(try_2))
              s2 = F_read_string(arg_1)
              }
              } 
            if ErrorIn(s2) {Result = s2
            } else {
            ClEnv.NLine = (ClEnv.NLine+1)
            F_pushback_port(C_reader.Fromp,cx)
            Result = s2}
            } 
          } 
        } 
      } 
    } 
  return Result} 

// The EID go function for: conditional_comment! @ meta_reader (throw: true) 
func E_Reader_conditional_comment_I_meta_reader2 (r EID,s EID) EID { 
  return ToMetaReader(OBJ(r)).ConditionalComment_I(ToString(OBJ(s)) )} 
