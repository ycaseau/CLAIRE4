/***** CLAIRE Compilation of file /Users/ycaseau/Dropbox/src/clairev4.07/src/meta/file.cl 
         [version 4.0.08 / safety 5] Sunday 03-12-2023 14:47:34 *****/

package Reader
import (_ "fmt"
	. "Kernel"
	"Core"
	"Language"
)

//-------- dumb function to prevent import errors --------
func import_g0130() { 
    _ = Core.It
    _ = Language.It
    } 
  
  
//+-------------------------------------------------------------+
//| CLAIRE                                                      |
//| file.cl                                                     |
//| Copyright (C) 1994 - 2021 Yves Caseau. All Rights Reserved  |
//| cf. copyright info in file object.cl: about()               |
//+-------------------------------------------------------------+
// --------------------------------------------------------------------
// this file contains all that is related to files + top-level
// --------------------------------------------------------------------
// **********************************************************************
// * Contents:                                                          *
// *  Part 1: Utilities                                                 *
// *  Part 2: Loading                                                   *
// *  Part 3: Reading in a file/string & top-level                      *
// *  Part 4: The show & kill methods + macro-methods                   *
// **********************************************************************
// **********************************************************************
// *   Part 1: Utilities                                                *
// **********************************************************************
// useful gadgets
//
/* The go function for: self_eval(self:delimiter) [status=1] */
func (self *Delimiter) SelfEval () EID { 
    var Result EID
    C_reader.Next()
    Result = ToException(Core.C_general_error.Make(MakeString("[117] loose delimiter ~S in program [line ~A ?]").Id(),MakeConstantList(self.Id(),MakeInteger(ClEnv.NLine).Id()).Id())).Close()
    return Result} 
  
// The EID go function for: self_eval @ delimiter (throw: true) 
func E_self_eval_delimiter (self EID) EID { 
    return ToDelimiter(OBJ(self)).SelfEval( )} 
  
// The EVAL go function for: delimiter 
func EVAL_delimiter (x *ClaireAny) EID { 
     return ToDelimiter(x).SelfEval()} 
  
// a small useful function
// PORTABILITY WARNING: the following assumes newline is ^J (ASCII 10 dec)
// PORTABILITY WARNING: what about ^M (ASCII 13 dec)
//
// a small usefull function
// note that these char cannot be written using #/_, it would be nicer with native chars
// CLAIRE4: added 160 which is an ' ' after a '|' produced by the Mac
/* The go function for: useless_c(r:integer) [status=0] */
func F_useless_c_integer (r int) *ClaireBoolean { 
    if (r == 10) { 
      ClEnv.NLine = (ClEnv.NLine+1)
      } 
    return  MakeBoolean((r == C_reader.Space) || 
    (r == 10) || 
    (r == 13) || 
    (r == 32) || 
    (r == 160) || 
    (r == C_reader.Tab))
    } 
  
// The EID go function for: useless_c @ integer (throw: false) 
func E_useless_c_integer (r EID) EID { 
    return EID{F_useless_c_integer(INT(r) ).Id(),0}} 
  
// take care of PC format (10 + 13)
/* The go function for: skipc(self:meta_reader) [status=0] */
func (self *MetaReader) Skipc () int { 
    for (F_useless_c_integer(self.Firstc()) == CTRUE) { 
      { var b *ClaireBoolean = Equal(MakeInteger(self.Firstc()).Id(),MakeInteger(10).Id())
        self.Next()
        if ((b == CTRUE) && 
            (self.Firstc() == 13)) { 
          self.Next()
          } 
        } 
      } 
    return  self.Firstc()
    } 
  
// The EID go function for: skipc @ meta_reader (throw: false) 
func E_skipc_meta_reader (self EID) EID { 
    return EID{C__INT,IVAL(ToMetaReader(OBJ(self)).Skipc( ))}} 
  
// look for a meaningful termination char such as ) or ]
/* The go function for: skipc!(r:meta_reader) [status=1] */
func (r *MetaReader) Skipc_I () EID { 
    var Result EID
    { var c int = r.Skipc()
      if (c == 59) { 
        for ((r.Firstc() != r.Eof) && 
            (r.Firstc() != 10)) { 
          r.Next()
          } 
        if (r.Firstc() == r.Eof) { 
          Result = EID{Language.C_EOF.Value,0}
          } else {
          ClEnv.NLine = (ClEnv.NLine+1)
          Result = r.Skipc_I()
          } 
        }  else if (c == 47) { 
        { var x *ClaireAny
          var try_1 EID
          try_1 = r.Fromp.ReadIdent()
          if ErrorIn(try_1) {Result = try_1
          } else {
          x = ANY(try_1)
          if (C_string.Id() == x.Isa.Id()) { 
            Result = r.Skipc_I()
            } else {
            Result = EID{C__INT,IVAL(47)}
            } 
          }
          } 
        } else {
        Result = EID{C__INT,IVAL(c)}
        } 
      } 
    return Result} 
  
// The EID go function for: skipc! @ meta_reader (throw: true) 
func E_skipc_I_meta_reader (r EID) EID { 
    return ToMetaReader(OBJ(r)).Skipc_I( )} 
  
/* The go function for: cnext(self:meta_reader) [status=0] */
func (self *MetaReader) Cnext () *MetaReader { 
    self.Next()
    return  self
    } 
  
// The EID go function for: cnext @ meta_reader (throw: false) 
func E_cnext_meta_reader (self EID) EID { 
    return EID{ToMetaReader(OBJ(self)).Cnext( ).Id(),0}} 
  
/* The go function for: findeol(self:meta_reader) [status=0] */
func (self *MetaReader) Findeol () *ClaireBoolean { 
    var Result *ClaireBoolean
    { 
      var v_or2 *ClaireBoolean
      
      v_or2= CFALSE
      for (F_useless_c_integer(self.Firstc()) == CTRUE) { 
        if (self.Firstc() == 10) { 
          v_or2 = CTRUE
          break
          } 
        self.Next()
        } 
      if (v_or2 == CTRUE) {Result = CTRUE
      } else { 
        v_or2 = Equal(MakeInteger(self.Firstc()).Id(),MakeInteger(self.Eof).Id())
        if (v_or2 == CTRUE) {Result = CTRUE
        } else { 
          Result = CFALSE} 
        } 
      } 
    return Result} 
  
// The EID go function for: findeol @ meta_reader (throw: false) 
func E_findeol_meta_reader (self EID) EID { 
    return EID{ToMetaReader(OBJ(self)).Findeol( ).Id(),0}} 
  
// safety checking
//
/* The go function for: checkno(r:meta_reader,n:integer,y:any) [status=1] */
func (r *MetaReader) Checkno (n int,y *ClaireAny) EID { 
    var Result EID
    if (r.Firstc() != n) { 
      Result = EID{r.Id(),0}
      } else {
      Result = F_Serror_string(MakeString("[118] read wrong char ~S after ~S"),MakeConstantList(MakeChar(F_char_I_integer(n)).Id(),y))
      } 
    return Result} 
  
// The EID go function for: checkno @ meta_reader (throw: true) 
func E_checkno_meta_reader (r EID,n EID,y EID) EID { 
    return ToMetaReader(OBJ(r)).Checkno(INT(n),ANY(y) )} 
  
// reads a keyword inside a control structure (used in Reader + OFTO)
//
/* The go function for: verify(t:any,x:any,y:any) [status=1] */
func F_verify_any (t *ClaireAny,x *ClaireAny,y *ClaireAny) EID { 
    var Result EID
    var g0133I *ClaireBoolean
    var try_1 EID
    try_1 = Core.F_BELONG(x,t)
    if ErrorIn(try_1) {Result = try_1
    } else {
    g0133I = ToBoolean(OBJ(try_1))
    if (g0133I == CTRUE) { 
      Result = x.ToEID()
      } else {
      Result = F_Serror_string(MakeString("[119] read ~S instead of a ~S in a ~S"),MakeConstantList(x,t,y))
      } 
    }
    return Result} 
  
// The EID go function for: verify @ any (throw: true) 
func E_verify_any (t EID,x EID,y EID) EID { 
    return F_verify_any(ANY(t),ANY(x),ANY(y) )} 
  
// prints a syntax error
//
/* The go function for: Serror(s:string,la:list) [status=1] */
func F_Serror_string (s *ClaireString,la *ClaireList) EID { 
    var Result EID
    PRINC("---- Syntax Error[line: ")
    F_princ_integer(ClEnv.NLine)
    PRINC("]:\n")
    F_flush_port(C_reader.Fromp)
    Result = ToException(Core.C_general_error.Make((s).Id(),la.Id())).Close()
    return Result} 
  
// The EID go function for: Serror @ string (throw: true) 
func E_Serror_string (s EID,la EID) EID { 
    return F_Serror_string(ToString(OBJ(s)),ToList(OBJ(la)) )} 
  
// the reader-------------------------------------------------------------
//
// variable handling -------------------------------------------------
// reads a variable
//
/* The go function for: extract_variable(self:any) [status=1] */
func F_extract_variable_any (self *ClaireAny) EID { 
    var Result EID
    var g0135I *ClaireBoolean
    if (self.Isa.IsIn(C_Variable) == CTRUE) { 
      { var g0134 *ClaireVariable = To_Variable(self)
        g0135I = Core.F__I_equal_any(g0134.Pname.Value(),g0134.Id())
        } 
      } else {
      g0135I = CFALSE
      } 
    if (g0135I == CTRUE) { 
      { 
        var va_arg1 *ClaireVariable
        var va_arg2 *ClaireType
        va_arg1 = To_Variable(self)
        var try_1 EID
        try_1 = Language.F_extract_type_any(ANY(Core.F_CALL(C_range,ARGS(self.ToEID()))))
        if ErrorIn(try_1) {Result = try_1
        } else {
        va_arg2 = ToType(OBJ(try_1))
        va_arg1.Range = va_arg2
        Result = EID{va_arg2.Id(),0}
        }
        } 
      if !ErrorIn(Result) {
      Result = self.ToEID()
      }
      } else {
      { var v *ClaireVariable
        var try_2 EID
        { var _CL_obj *ClaireVariable = To_Variable(new(ClaireVariable).Is(C_Variable))
          { 
            var va_arg1 *ClaireVariable
            var va_arg2 *ClaireSymbol
            va_arg1 = _CL_obj
            var try_3 EID
            try_3 = Language.F_extract_symbol_any(self)
            if ErrorIn(try_3) {try_2 = try_3
            } else {
            va_arg2 = ToSymbol(OBJ(try_3))
            va_arg1.Pname = va_arg2
            try_2 = EID{va_arg2.Id(),0}
            }
            } 
          if !ErrorIn(try_2) {
          try_2 = EID{_CL_obj.Id(),0}
          }
          } 
        if ErrorIn(try_2) {Result = try_2
        } else {
        v = To_Variable(OBJ(try_2))
        C_reader.LastForm = v.Id()
        Result = EID{v.Id(),0}
        }
        } 
      } 
    return Result} 
  
// The EID go function for: extract_variable @ any (throw: true) 
func E_extract_variable_any (self EID) EID { 
    return F_extract_variable_any(ANY(self) )} 
  
// create a variable and add it to the lexical environment
/* The go function for: bind!(self:meta_reader,%v:Variable) [status=0] */
func (self *MetaReader) Bind_I (_Zv *ClaireVariable) *ClaireList { 
    var Result *ClaireList
    _Zv.Index = self.Index
    { var value *ClaireAny = _Zv.Pname.Value()
      self.Index = (self.Index+1)
      if (self.Index > self.Maxstack) { 
        self.Maxstack = self.Index
        } 
      _Zv.Pname.Put(_Zv.Id())
      Result = MakeConstantList(_Zv.Id(),value)
      } 
    return Result} 
  
// The EID go function for: bind! @ meta_reader (throw: false) 
func E_bind_I_meta_reader (self EID,_Zv EID) EID { 
    return EID{ToMetaReader(OBJ(self)).Bind_I(To_Variable(OBJ(_Zv)) ).Id(),0}} 
  
// remove a variable from the lexical environment
//
/* The go function for: unbind!(self:meta_reader,%first:list) [status=0] */
func (self *MetaReader) Unbind_I (_Zfirst *ClaireList) *ClaireAny { 
    var Result *ClaireAny
    { var ClaireVar *ClaireAny = _Zfirst.At(0)
      self.Index = (self.Index-1)
      Result = ToSymbol(OBJ(Core.F_CALL(C_mClaire_pname,ARGS(ClaireVar.ToEID())))).Put(_Zfirst.At(1))
      } 
    return Result} 
  
// The EID go function for: unbind! @ meta_reader (throw: false) 
func E_unbind_I_meta_reader (self EID,_Zfirst EID) EID { 
    return ToMetaReader(OBJ(self)).Unbind_I(ToList(OBJ(_Zfirst)) ).ToEID()} 
  
// declaration of the CLAIRE standard ports ----------------------------
// we create global variables - however they exists as properties of system => redundant ?
//
/* The go function for: /(s:string,s2:string) [status=0] */
func F__7_string (s *ClaireString,s2 *ClaireString) *ClaireString { 
    return  F_append_string(F_append_string(s,ToString(C__starfs_star.Value)),s2)
    } 
  
// The EID go function for: / @ string (throw: false) 
func E__7_string (s EID,s2 EID) EID { 
    return EID{F__7_string(ToString(OBJ(s)),ToString(OBJ(s2)) ).Id(),0}} 
  
// basic methods defined in creader.c -----------------------------------
// TODO move!
// flush(self:port) : any -> function!(flush_port)
// this function is called by the main and restores the reader in a good shape. Also
// closes the input port to free the associated file ! <yc>
/* The go function for: mClaire/restore_state(self:meta_reader) [status=0] */
func (self *MetaReader) RestoreState ()  { 
    if (Equal(self.Fromp.Id(),C_stdin.Value) != CTRUE) { 
      self.Fromp.Fclose()
      } 
    self.Fromp = ToPort(C_stdin.Value)
    self.Index = 1
    F_pushback_port(ToPort(C_stdin.Value),32)
    F_restore_state_void()
    } 
  
// The EID go function for: mClaire/restore_state @ meta_reader (throw: false) 
func E_restore_state_meta_reader (self EID) EID { 
    ToMetaReader(OBJ(self)).RestoreState( )
    return EVOID} 
  
// *********************************************************************
// *   Part 2: Loading                                                 *
// *********************************************************************
// sload is the interactive version (when b = true).
//
/* The go function for: load_file(self:string,b:boolean) [status=1] */
func F_load_file_string (self *ClaireString,b *ClaireBoolean) EID { 
    var Result EID
    C_reader.Index = 0
    C_reader.Maxstack = 0
    ClEnv.NLine = 1
    C_reader.External = self
    Core.F_tformat_string(MakeString("---- [load CLAIRE file: ~A]\n"),2,MakeConstantList((self).Id()))
    { var s2 *ClaireString = F_append_string(self,MakeString(".cl"))
      { var p1 *ClairePort
        var try_1 EID
        { 
          h_index := ClEnv.Index
          h_base := ClEnv.Base
          try_1 = F_fopen_string(s2,MakeString("r"))
          if ErrorIn(try_1){ 
            ClEnv.Index = h_index
            ClEnv.Base = h_base
            { 
              h_index := ClEnv.Index
              h_base := ClEnv.Base
              try_1 = F_fopen_string(self,MakeString("r"))
              if ErrorIn(try_1){ 
                ClEnv.Index = h_index
                ClEnv.Base = h_base
                try_1 = ToException(Core.C_general_error.Make(MakeString("[120] the file ~A cannot be opened").Id(),MakeConstantList((self).Id()).Id())).Close()
                } 
              } 
            } 
          } 
        if ErrorIn(try_1) {Result = try_1
        } else {
        p1 = ToPort(OBJ(try_1))
        { var start int = ClEnv.Base
          { var top int = ClEnv.Index
            { var p2 *ClairePort = C_reader.Fromp
              { var b2 *ClaireBoolean = C_reader.Toplevel
                { var _staritem_star *ClaireAny = CNULL
                  ClEnv.Base= top
                  C_reader.Fromp = p1
                  C_reader.Toplevel = CFALSE
                  var try_2 EID
                  try_2 = F_readblock_port(p1)
                  if ErrorIn(try_2) {Result = try_2
                  } else {
                  _staritem_star = ANY(try_2)
                  Result = _staritem_star.ToEID()
                  Result= EID{CFALSE.Id(),0}
                  for (_staritem_star != C_Reader_eof.Id()) { 
                    var loop_3 EID
                    _ = loop_3
                    { 
                    if (b == CTRUE) { 
                      PRINC("<")
                      F_princ_integer(ClEnv.NLine)
                      PRINC(">:")
                      loop_3 = Core.F_CALL(C_print,ARGS(_staritem_star.ToEID()))
                      if ErrorIn(loop_3) {Result = loop_3
                      break
                      } else {
                      PRINC("\n")
                      loop_3 = EVOID
                      }
                      } else {
                      loop_3 = EID{CFALSE.Id(),0}
                      } 
                    if ErrorIn(loop_3) {Result = loop_3
                    break
                    } else {
                    ClEnv.Index= (top+(C_reader.Maxstack+1))
                    if (C_string.Id() == _staritem_star.Isa.Id()) { 
                      { var g0136 *ClaireString = ToString(_staritem_star)
                        if (ToBoolean(Language.C_NeedComment.Value) == CTRUE) { 
                          if (Language.C_iClaire_LastComment.Value != CNULL) { 
                            var v_gassign4 *ClaireAny
                            v_gassign4 = ANY(Core.F_CALL(ToProperty(C__7_plus.Id()),ARGS(EID{Language.C_iClaire_LastComment.Value,0},EID{F_append_string(MakeString("\n-- "),g0136).Id(),0})))
                            Language.C_iClaire_LastComment.Value = v_gassign4
                            loop_3 = v_gassign4.ToEID()
                            } else {
                            var v_gassign5 *ClaireAny
                            v_gassign5 = (F_append_string(F_append_string(F_append_string(F_append_string(F_append_string(MakeString("["),C_reader.External),MakeString("(")),F_string_I_integer(ClEnv.NLine)),MakeString(")]\n-- ")),g0136)).Id()
                            Language.C_iClaire_LastComment.Value = v_gassign5
                            loop_3 = v_gassign5.ToEID()
                            } 
                          } else {
                          loop_3 = EID{CFALSE.Id(),0}
                          } 
                        } 
                      } else {
                      var try_6 EID
                      try_6 = EVAL(_staritem_star)
                      if ErrorIn(try_6) {loop_3 = try_6
                      Result = try_6
                      break
                      } else {
                      _staritem_star = ANY(try_6)
                      loop_3 = _staritem_star.ToEID()
                      var v_gassign7 *ClaireAny
                      v_gassign7 = CNULL
                      Language.C_iClaire_LastComment.Value = v_gassign7
                      loop_3 = v_gassign7.ToEID()
                      }
                      } 
                    if ErrorIn(loop_3) {Result = loop_3
                    break
                    } else {
                    if ((b == CTRUE) && 
                        (C_string.Id() != _staritem_star.Isa.Id())) { 
                      PRINC("=> ")
                      loop_3 = Core.F_CALL(C_print,ARGS(_staritem_star.ToEID()))
                      if ErrorIn(loop_3) {Result = loop_3
                      break
                      } else {
                      PRINC(" \n\n")
                      loop_3 = EVOID
                      }
                      } else {
                      loop_3 = EID{CFALSE.Id(),0}
                      } 
                    if ErrorIn(loop_3) {Result = loop_3
                    break
                    } else {
                    var try_8 EID
                    try_8 = F_readblock_port(p1)
                    if ErrorIn(try_8) {loop_3 = try_8
                    Result = try_8
                    break
                    } else {
                    _staritem_star = ANY(try_8)
                    loop_3 = _staritem_star.ToEID()
                    }}}}
                    } 
                  }
                  if !ErrorIn(Result) {
                  ClEnv.Base= start
                  ClEnv.Index= top
                  C_reader.Toplevel = b2
                  C_reader.Fromp = p2
                  C_reader.External = MakeString("toplevel")
                  p1.Fclose()
                  Result = EVOID
                  }}
                  } 
                } 
              } 
            } 
          } 
        }
        } 
      } 
    if !ErrorIn(Result) {
    Result = EID{CTRUE.Id(),0}
    }
    return Result} 
  
// The EID go function for: load_file @ string (throw: true) 
func E_load_file_string (self EID,b EID) EID { 
    return F_load_file_string(ToString(OBJ(self)),ToBoolean(OBJ(b)) )} 
  
// the simple load
//
/* The go function for: load(self:string) [status=1] */
func F_load_string (self *ClaireString) EID { 
    var Result EID
    Result = F_load_file_string(self,CFALSE)
    return Result} 
  
// The EID go function for: load @ string (throw: true) 
func E_load_string (self EID) EID { 
    return F_load_string(ToString(OBJ(self)) )} 
  
/* The go function for: sload(self:string) [status=1] */
func F_sload_string (self *ClaireString) EID { 
    var Result EID
    Result = F_load_file_string(self,CTRUE)
    return Result} 
  
// The EID go function for: sload @ string (throw: true) 
func E_sload_string (self EID) EID { 
    return F_sload_string(ToString(OBJ(self)) )} 
  
// loading a module into the system.
// The correct package is open and each file is loaded.
/* The go function for: load_file(self:module,b:boolean) [status=1] */
func F_load_file_module (self *ClaireModule,b *ClaireBoolean) EID { 
    var Result EID
    if (self.Status == 2) { 
      { 
        var va_arg1 *ClaireModule
        var va_arg2 int
        va_arg1 = self
        va_arg2 = 3
        va_arg1.Status = va_arg2
        Result = EID{C__INT,IVAL(va_arg2)}
        } 
      }  else if ((self.Status == 0) && 
        ((self.Source).Id() != CNULL)) { 
      Core.F_tformat_string(MakeString("---- Loading the module ~S.\n"),1,MakeConstantList(self.Id()))
      self.Begin()
      { var s *ClaireString = F_append_string(self.Source,ToString(C__starfs_star.Value))
        { 
          var x *ClaireAny
          _ = x
          Result= EID{CFALSE.Id(),0}
          var x_support *ClaireList
          x_support = self.MadeOf
          x_len := x_support.Length()
          for i_it := 0; i_it < x_len; i_it++ { 
            x = x_support.At(i_it)
            var loop_1 EID
            _ = loop_1
            loop_1 = Core.F_CALL(C_Reader_load_file,ARGS(Core.F_CALL(ToProperty(C__7_plus.Id()),ARGS(EID{F_append_string(s,ToString(x)).Id(),0},EID{MakeString(".cl").Id(),0})),EID{b.Id(),0}))
            if ErrorIn(loop_1) {Result = loop_1
            break
            } else {
            }
            } 
          } 
        } 
      if !ErrorIn(Result) {
      { 
        var va_arg1 *ClaireModule
        var va_arg2 int
        va_arg1 = self
        va_arg2 = 1
        va_arg1.Status = va_arg2
        Result = EID{C__INT,IVAL(va_arg2)}
        } 
      }
      } else {
      Result = EID{CFALSE.Id(),0}
      } 
    if !ErrorIn(Result) {
    self.End()
    Result = EVOID
    }
    return Result} 
  
// The EID go function for: load_file @ module (throw: true) 
func E_load_file_module (self EID,b EID) EID { 
    return F_load_file_module(ToModule(OBJ(self)),ToBoolean(OBJ(b)) )} 
  
// the simple load
//
/* The go function for: load(self:module) [status=1] */
func F_load_module (self *ClaireModule) EID { 
    var Result EID
    { 
      var x *ClaireAny
      _ = x
      Result= EID{CFALSE.Id(),0}
      var x_support *ClaireList
      x_support = F_add_modules_list(MakeConstantList(self.Id()))
      x_len := x_support.Length()
      for i_it := 0; i_it < x_len; i_it++ { 
        x = x_support.At(i_it)
        var loop_1 EID
        _ = loop_1
        loop_1 = Core.F_CALL(C_Reader_load_file,ARGS(x.ToEID(),EID{CFALSE.Id(),0}))
        if ErrorIn(loop_1) {Result = loop_1
        break
        } else {
        }
        } 
      } 
    return Result} 
  
// The EID go function for: load @ module (throw: true) 
func E_load_module (self EID) EID { 
    return F_load_module(ToModule(OBJ(self)) )} 
  
/* The go function for: sload(self:module) [status=1] */
func F_sload_module (self *ClaireModule) EID { 
    var Result EID
    { 
      var x *ClaireAny
      _ = x
      Result= EID{CFALSE.Id(),0}
      var x_support *ClaireList
      x_support = F_add_modules_list(MakeConstantList(self.Id()))
      x_len := x_support.Length()
      for i_it := 0; i_it < x_len; i_it++ { 
        x = x_support.At(i_it)
        var loop_1 EID
        _ = loop_1
        loop_1 = Core.F_CALL(C_Reader_load_file,ARGS(x.ToEID(),EID{CTRUE.Id(),0}))
        if ErrorIn(loop_1) {Result = loop_1
        break
        } else {
        }
        } 
      } 
    return Result} 
  
// The EID go function for: sload @ module (throw: true) 
func E_sload_module (self EID) EID { 
    return F_sload_module(ToModule(OBJ(self)) )} 
  
// This is a very important method which adds the right order the
// modules that must be loaded to load oself. the list l represents the
// list of modules that we know will be in the result. result represent
// the current list of ordered modules
//
/* The go function for: add_modules(self:module,l:set,result:list) [status=0] */
func F_add_modules_module (self *ClaireModule,l *ClaireSet,result *ClaireList) *ClaireList { 
    if (result.Memq(self.Id()) == CTRUE) { 
      return  result
      }  else if (l.Contain_ask(self.Id()) == CTRUE) { 
      return  result.AddFast(self.Id())
      } else {
      l = l.AddFast(self.Id())
      { 
        var x *ClaireAny
        _ = x
        var x_support *ClaireList
        x_support = self.Uses
        x_len := x_support.Length()
        for i_it := 0; i_it < x_len; i_it++ { 
          x = x_support.At(i_it)
          if (x.Isa.IsIn(C_module) == CTRUE) { 
            { var g0138 *ClaireModule = ToModule(x)
              result = F_add_modules_module(g0138,l,result)
              } 
            } 
          } 
        } 
      if (result.Memq(self.Id()) != CTRUE) { 
        result = result.AddFast(self.Id())
        } 
      { 
        var x *ClaireAny
        _ = x
        var x_support *ClaireList
        x_support = self.Parts
        x_len := x_support.Length()
        for i_it := 0; i_it < x_len; i_it++ { 
          x = x_support.At(i_it)
          result = F_add_modules_module(ToModule(x),l,result)
          } 
        } 
      return  result
      } 
    } 
  
// The EID go function for: add_modules @ module (throw: false) 
func E_add_modules_module (self EID,l EID,result EID) EID { 
    return EID{F_add_modules_module(ToModule(OBJ(self)),ToSet(OBJ(l)),ToList(OBJ(result)) ).Id(),0}} 
  
// this methods takes a list of modules that must be loaded and returns
// a list of modules that are necessary for the definition
//
/* The go function for: add_modules(self:list) [status=0] */
func F_add_modules_list (self *ClaireList) *ClaireList { 
    var Result *ClaireList
    { var l *ClaireList = ToType(C_module.Id()).EmptyList()
      { 
        var x *ClaireAny
        _ = x
        var x_support *ClaireList
        x_support = self
        x_len := x_support.Length()
        for i_it := 0; i_it < x_len; i_it++ { 
          x = x_support.At(i_it)
          l = F_add_modules_module(ToModule(x),l.Set_I(),l)
          } 
        } 
      Result = l
      } 
    return Result} 
  
// The EID go function for: add_modules @ list (throw: false) 
func E_add_modules_list (self EID) EID { 
    return EID{F_add_modules_list(ToList(OBJ(self)) ).Id(),0}} 
  
// load a file of expressions (quite useful)
/* The go function for: eload(self:string) [status=1] */
func F_eload_string (self *ClaireString) EID { 
    var Result EID
    C_reader.Index = 0
    C_reader.Maxstack = 0
    ClEnv.NLine = 1
    C_reader.External = self
    Core.F_tformat_string(MakeString("---- [eload CLAIRE file: ~A]\n"),2,MakeConstantList((self).Id()))
    { var s2 *ClaireString = F_append_string(self,MakeString(".cl"))
      { var p0 *ClairePort = C_reader.Fromp
        { var p1 *ClairePort
          var try_1 EID
          { 
            h_index := ClEnv.Index
            h_base := ClEnv.Base
            try_1 = F_fopen_string(s2,MakeString("r"))
            if ErrorIn(try_1){ 
              ClEnv.Index = h_index
              ClEnv.Base = h_base
              { 
                h_index := ClEnv.Index
                h_base := ClEnv.Base
                try_1 = F_fopen_string(self,MakeString("r"))
                if ErrorIn(try_1){ 
                  ClEnv.Index = h_index
                  ClEnv.Base = h_base
                  try_1 = ToException(Core.C_general_error.Make(MakeString("[120] the file ~A cannot be opened").Id(),MakeConstantList((self).Id()).Id())).Close()
                  } 
                } 
              } 
            } 
          if ErrorIn(try_1) {Result = try_1
          } else {
          p1 = ToPort(OBJ(try_1))
          { var start int = ClEnv.Base
            { var top int = ClEnv.Index
              { var b2 *ClaireBoolean = C_reader.Toplevel
                { var _staritem_star *ClaireAny = CNULL
                  ClEnv.Base= top
                  C_reader.Toplevel = CFALSE
                  C_reader.Fromp = p1
                  var try_2 EID
                  try_2 = F_read_port(p1)
                  if ErrorIn(try_2) {Result = try_2
                  } else {
                  _staritem_star = ANY(try_2)
                  Result = _staritem_star.ToEID()
                  Result= EID{CFALSE.Id(),0}
                  for (_staritem_star != C_Reader_eof.Id()) { 
                    var loop_3 EID
                    _ = loop_3
                    { 
                    ClEnv.Index= (top+(C_reader.Maxstack+1))
                    var try_4 EID
                    try_4 = EVAL(_staritem_star)
                    if ErrorIn(try_4) {loop_3 = try_4
                    Result = try_4
                    break
                    } else {
                    _staritem_star = ANY(try_4)
                    loop_3 = _staritem_star.ToEID()
                    var try_5 EID
                    try_5 = F_read_port(p1)
                    if ErrorIn(try_5) {loop_3 = try_5
                    Result = try_5
                    break
                    } else {
                    _staritem_star = ANY(try_5)
                    loop_3 = _staritem_star.ToEID()
                    }}
                    } 
                  }
                  if !ErrorIn(Result) {
                  ClEnv.Base= start
                  ClEnv.Index= top
                  C_reader.Fromp = p0
                  C_reader.Toplevel = b2
                  C_reader.External = MakeString("toplevel")
                  p1.Fclose()
                  Result = EVOID
                  }}
                  } 
                } 
              } 
            } 
          }
          } 
        } 
      } 
    if !ErrorIn(Result) {
    Result = EID{CTRUE.Id(),0}
    }
    return Result} 
  
// The EID go function for: eload @ string (throw: true) 
func E_eload_string (self EID) EID { 
    return F_eload_string(ToString(OBJ(self)) )} 
  
// *********************************************************************
// *   Part 3: Read & Top-level                                        *
// *********************************************************************
// The standard read function.
// This method reads from a CLAIRE port (self).
// We first check if self is the current reading port.
// the last character read (and not used) is in last(reader)
/* The go function for: readblock(p:port) [status=1] */
func F_readblock_port (p *ClairePort) EID { 
    var Result EID
    if (Equal(C_reader.Fromp.Id(),p.Id()) == CTRUE) { 
      Result = C_reader.Nextunit()
      } else {
      { var p2 *ClairePort = C_reader.Fromp
        C_reader.Fromp = p
        { var val *ClaireAny
          var try_1 EID
          try_1 = C_reader.Nextunit()
          if ErrorIn(try_1) {Result = try_1
          } else {
          val = ANY(try_1)
          C_reader.Fromp = p2
          if ((Equal(val,C_reader.Paren) == CTRUE) || 
              ((Equal(val,C_reader.Curly) == CTRUE) || 
                ((Equal(val,C_reader.Comma) == CTRUE) || 
                  (Equal(val,C_reader.Bracket) == CTRUE)))) { 
            Result = F_Serror_string(MakeString("[117] Loose ~S in file"),MakeConstantList(val))
            } else {
            Result = EID{CFALSE.Id(),0}
            } 
          if !ErrorIn(Result) {
          Result = val.ToEID()
          }
          }
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: readblock @ port (throw: true) 
func E_readblock_port (p EID) EID { 
    return F_readblock_port(ToPort(OBJ(p)) )} 
  
// read reads a closed expression
/* The go function for: read(p:port) [status=1] */
func F_read_port (p *ClairePort) EID { 
    var Result EID
    { var p2 *ClairePort = C_reader.Fromp
      if (Equal(p.Id(),p2.Id()) != CTRUE) { 
        C_reader.Fromp = p
        } 
      { var val *ClaireAny
        var try_1 EID
        if (C_reader.Skipc() == C_reader.Eof) { 
          try_1 = EID{C_Reader_eof.Id(),0}
          } else {
          try_1 = C_reader.Nexte()
          } 
        if ErrorIn(try_1) {Result = try_1
        } else {
        val = ANY(try_1)
        if (Equal(p.Id(),p2.Id()) != CTRUE) { 
          C_reader.Fromp = p2
          } 
        if ((Equal(val,C_reader.Paren) == CTRUE) || 
            ((Equal(val,C_reader.Curly) == CTRUE) || 
              ((Equal(val,C_reader.Comma) == CTRUE) || 
                (Equal(val,C_reader.Bracket) == CTRUE)))) { 
          Result = F_Serror_string(MakeString("[117] Loose ~S in file"),MakeConstantList(val))
          } else {
          Result = EID{CFALSE.Id(),0}
          } 
        if !ErrorIn(Result) {
        Result = val.ToEID()
        }
        }
        } 
      } 
    return Result} 
  
// The EID go function for: read @ port (throw: true) 
func E_read_port (p EID) EID { 
    return F_read_port(ToPort(OBJ(p)) )} 
  
// read into a string
/* The go function for: read(self:string) [status=1] */
func F_read_string (self *ClaireString) EID { 
    var Result EID
    { var b *ClaireBoolean = C_reader.Toplevel
      { var p *ClairePort = C_reader.Fromp
        { var x *ClaireAny = CNULL
          C_reader.Toplevel = CTRUE
          C_reader.Fromp = F_port_I_string(self)
          { 
            h_index := ClEnv.Index
            h_base := ClEnv.Base
            var try_1 EID
            try_1 = C_reader.Nextunit()
            if ErrorIn(try_1) {Result = try_1
            } else {
            x = ANY(try_1)
            Result = x.ToEID()
            { 
              var va_arg1 *MetaReader
              var va_arg2 *ClairePort
              va_arg1 = C_reader
              va_arg2 = p
              va_arg1.Fromp = va_arg2
              Result = va_arg2.ToEID()
              } 
            }
            if ErrorIn(Result){ 
              ClEnv.Index = h_index
              ClEnv.Base = h_base
              C_reader.Fromp = p
              Result = ClEnv.Exception_I.Close()
              } 
            } 
          if !ErrorIn(Result) {
          C_reader.Toplevel = b
          Result = x.ToEID()
          }
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: read @ string (throw: true) 
func E_read_string (self EID) EID { 
    return F_read_string(ToString(OBJ(self)) )} 
  
// used by the top level
// calls the debugger
/* The go function for: debug_if_possible(_CL_obj:void) [status=1] */
func F_debug_if_possible_void () EID { 
    var Result EID
    if (ClEnv.Debug_I >= 0) { 
      Result = F_funcall1(ToMethod(C_call_debug.Restrictions.ValuesO()[0]).Functional,ClEnv.Id())
      } else {
      Result = F_print_exception_void().ToEID()
      } 
    return Result} 
  
// The EID go function for: debug_if_possible @ void (throw: true) 
func E_debug_if_possible_void (_CL_obj EID) EID { 
    return F_debug_if_possible_void( )} 
  
// a method for calling the printer without issuing a message (that would
// modify the stack and make debugging impossible).
// here we assume that self_print is always defined and is always a compiled
// function
/* The go function for: print_exception(_CL_obj:void) [status=0] */
func F_print_exception_void () *ClaireAny { 
    var Result *ClaireAny
    { var p *ClairePort = ToPort(C_stdout.Value).UseAsOutput()
      { var _Zerr *ClaireException = ClEnv.Exception_I
        { var _Zprop *ClaireMethod = ToMethod(Core.F__at_property1(C_self_print,_Zerr.Id().Isa).Id())
          { 
            var Result_H EID
            h_index := ClEnv.Index
            h_base := ClEnv.Base
            if (_Zprop.Functional.Id() != CNULL) { 
              Result_H = F_funcall1(_Zprop.Functional,_Zerr.Id())
              } else {
              Result_H = Core.F_funcall_method1(_Zprop,_Zerr.Id())
              } 
            if ErrorIn(Result_H){ 
              ClEnv.Index = h_index
              ClEnv.Base = h_base
              PRINC("****** ERROR[121]: unprintable error has occurred.\n")
              } 
            } 
          Result = p.UseAsOutput().Id()
          } 
        } 
      } 
    return Result} 
  
// The EID go function for: print_exception @ void (throw: false) 
func E_print_exception_void (_CL_obj EID) EID { 
    return F_print_exception_void( ).ToEID()} 
  
// **********************************************************************
// *  Part 4: The show & kill methods + macro-methods                   *
// **********************************************************************
//----------------- printing an object -------------------------
// %show is an open restriction which allow to show the value of a
// binary relation
//
// this method is the basic method called for show(..)
//
/* The go function for: show(self:any) [status=1] */
func F_show_any (self *ClaireAny) EID { 
    var Result EID
    if (Core.F__Z_any1(self,C_object) == CTRUE) { 
      { var g0139 *ClaireObject = ToObject(self)
        { 
          var rel *ClaireSlot
          _ = rel
          var rel_iter *ClaireAny
          Result= EID{CFALSE.Id(),0}
          var rel_support *ClaireList
          rel_support = g0139.Id().Isa.Slots
          for _,rel_iter = range(rel_support.ValuesO()){ 
            rel = ToSlot(rel_iter)
            var loop_1 EID
            _ = loop_1
            { 
            loop_1 = Core.F_print_any(rel.Selector.Id())
            if ErrorIn(loop_1) {Result = loop_1
            break
            } else {
            PRINC(": ")
            loop_1 = Core.F_CALL(C_print,ARGS(Core.F_get_slot(rel,g0139).ToEID()))
            if ErrorIn(loop_1) {Result = loop_1
            break
            } else {
            PRINC("\n")
            }}
            }
            } 
          } 
        } 
      } else {
      Result = Core.F_CALL(C_print,ARGS(self.ToEID()))
      if !ErrorIn(Result) {
      PRINC(" is a ")
      Result = Core.F_print_any(self.Isa.Id())
      if !ErrorIn(Result) {
      PRINC("\n")
      Result = EVOID
      }}
      } 
    if !ErrorIn(Result) {
    Result = EID{CTRUE.Id(),0}
    }
    return Result} 
  
// The EID go function for: show @ any (throw: true) 
func E_show_any (self EID) EID { 
    return F_show_any(ANY(self) )} 
  
// This is the good version of kill, the nasty one is dangerous ....
// these restrictions of kill explain the dependencies among objects
//
/* The go function for: kill(self:object) [status=0] */
func F_kill_object (self *ClaireObject) *ClaireAny { 
    if (self.Isa.IsIn(C_thing) == CTRUE) { 
      { var g0141 *ClaireThing = ToThing(self.Id())
        g0141.Name.Put(CNULL)
        } 
      } 
    self.Isa.Instances = self.Isa.Instances.Delete(self.Id())
    return  CEMPTY.Id()
    } 
  
// The EID go function for: kill @ object (throw: false) 
func E_kill_object (self EID) EID { 
    return F_kill_object(ToObject(OBJ(self)) ).ToEID()} 
  
/* The go function for: kill(self:class) [status=0] */
func F_kill_class (self *ClaireClass) *ClaireAny { 
    for (self.Instances.Length() != 0) { 
      Core.F_CALL(C_kill,ARGS(self.Instances.At(0).ToEID()))
      } 
    { 
      var x *ClaireClass
      _ = x
      var x_iter *ClaireAny
      var x_support *ClaireSet
      x_support = self.Descendants
      for i_it := 0; i_it < x_support.Count; i_it++ { 
        x_iter = x_support.At(i_it)
        x = ToClass(x_iter)
        if (x.Superclass.Id() == self.Id()) { 
          F_kill_class(x)
          } 
        } 
      } 
    return  F_kill_object(ToObject(self.Id()))
    } 
  
// The EID go function for: kill @ class (throw: false) 
func E_kill_class (self EID) EID { 
    return F_kill_class(ToClass(OBJ(self)) ).ToEID()} 
  
// our two very special inline methods
/* The go function for: min(x:integer,y:integer) [status=0] */
func F_min_integer (x int,y int) int { 
    if (x <= y) { 
      return  x
      } else {
      return  y
      } 
    } 
  
// The EID go function for: min @ integer (throw: false) 
func E_min_integer (x EID,y EID) EID { 
    return EID{C__INT,IVAL(F_min_integer(INT(x),INT(y) ))}} 
  
/* The go function for: max(x:integer,y:integer) [status=0] */
func F_max_integer (x int,y int) int { 
    if (x <= y) { 
      return  y
      } else {
      return  x
      } 
    } 
  
// The EID go function for: max @ integer (throw: false) 
func E_max_integer (x EID,y EID) EID { 
    return EID{C__INT,IVAL(F_max_integer(INT(x),INT(y) ))}} 
  
/* The go function for: min(x:float,y:float) [status=0] */
func F_min_float (x float64,y float64) float64 { 
    if (x <= y) { 
      return  x
      } else {
      return  y
      } 
    } 
  
// The EID go function for: min @ float (throw: false) 
func E_min_float (x EID,y EID) EID { 
    return EID{C__FLOAT,FVAL(F_min_float(FLOAT(x),FLOAT(y) ))}} 
  
/* The go function for: max(x:float,y:float) [status=0] */
func F_max_float (x float64,y float64) float64 { 
    if (x <= y) { 
      return  y
      } else {
      return  x
      } 
    } 
  
// The EID go function for: max @ float (throw: false) 
func E_max_float (x EID,y EID) EID { 
    return EID{C__FLOAT,FVAL(F_max_float(FLOAT(x),FLOAT(y) ))}} 
  
/* The go function for: min(x:any,y:any) [status=1] */
func F_min_any (x *ClaireAny,y *ClaireAny) EID { 
    var Result EID
    Result = IfThenElse((ToBoolean(OBJ(Core.F_CALL(ToProperty(C__inf_equal.Id()),ARGS(x.ToEID(),y.ToEID())))) == CTRUE),
      x,
      y).ToEID()
    return Result} 
  
// The EID go function for: min @ any (throw: true) 
func E_min_any (x EID,y EID) EID { 
    return F_min_any(ANY(x),ANY(y) )} 
  
/* The go function for: max(x:any,y:any) [status=1] */
func F_max_any (x *ClaireAny,y *ClaireAny) EID { 
    var Result EID
    Result = IfThenElse((ToBoolean(OBJ(Core.F_CALL(ToProperty(C__inf_equal.Id()),ARGS(x.ToEID(),y.ToEID())))) == CTRUE),
      y,
      x).ToEID()
    return Result} 
  
// The EID go function for: max @ any (throw: true) 
func E_max_any (x EID,y EID) EID { 
    return F_max_any(ANY(x),ANY(y) )} 
  
// check if the value if known?
/* The go function for: known?(a:table,x:any) [status=0] */
func F_known_ask_table (a *ClaireTable,x *ClaireAny) *ClaireBoolean { 
    return  Core.F__I_equal_any(Core.F_get_table(a,x),CNULL)
    } 
  
// The EID go function for: known? @ table (throw: false) 
func E_known_ask_table (a EID,x EID) EID { 
    return EID{F_known_ask_table(ToTable(OBJ(a)),ANY(x) ).Id(),0}} 
  
/* The go function for: unknown?(a:table,x:any) [status=0] */
func F_unknown_ask_table (a *ClaireTable,x *ClaireAny) *ClaireBoolean { 
    if (Core.F_get_table(a,x) == CNULL) {return CTRUE
    } else {return CFALSE}} 
  
// The EID go function for: unknown? @ table (throw: false) 
func E_unknown_ask_table (a EID,x EID) EID { 
    return EID{F_unknown_ask_table(ToTable(OBJ(a)),ANY(x) ).Id(),0}} 
  
/* The go function for: float!(self:string) [status=1] */
func F_float_I_string (self *ClaireString) EID { 
    var Result EID
    { var x *ClaireAny
      var try_1 EID
      try_1 = F_read_string(self)
      if ErrorIn(try_1) {Result = try_1
      } else {
      x = ANY(try_1)
      if (C_float.Id() == x.Isa.Id()) { 
        { var g0142 float64 = ToFloat(x).Value
          Result = EID{C__FLOAT,FVAL(g0142)}
          } 
        }  else if (C_integer.Id() == x.Isa.Id()) { 
        { var g0143 int = ToInteger(x).Value
          Result = EID{C__FLOAT,FVAL(F_to_float(g0143))}
          } 
        } else {
        Result = ToException(Core.C_general_error.Make(MakeString("[??] ~A is not a float").Id(),MakeConstantList((self).Id()).Id())).Close()
        } 
      }
      } 
    return Result} 
  
// The EID go function for: float! @ string (throw: true) 
func E_float_I_string (self EID) EID { 
    return F_float_I_string(ToString(OBJ(self)) )} 
  
// v3.00.46 a new macro
/* The go function for: >=(self:any,x:any) [status=1] */
func F__sup_equal_any (self *ClaireAny,x *ClaireAny) *ClaireBoolean { 
    return  ToBoolean(OBJ(Core.F_CALL(ToProperty(C__inf_equal.Id()),ARGS(x.ToEID(),self.ToEID()))))
    } 
  
// The EID go function for: >= @ any (throw: false) 
func E__sup_equal_any (self EID,x EID) EID { 
    return EID{F__sup_equal_any(ANY(self),ANY(x) ).Id(),0}} 
  
// v3.3.42 add macros to use float & integers easily
/* The go function for: +(x:integer,y:float) [status=0] */
func F__plus_integer2 (x int,y float64) float64 { 
    return  (F_to_float(x)+y)
    } 
  
// The EID go function for: + @ list<type_expression>(integer, float) (throw: false) 
func E__plus_integer2 (x EID,y EID) EID { 
    return EID{C__FLOAT,FVAL(F__plus_integer2(INT(x),FLOAT(y) ))}} 
  
/* The go function for: *(x:integer,y:float) [status=0] */
func F__star_integer2 (x int,y float64) float64 { 
    return  (F_to_float(x)*y)
    } 
  
// The EID go function for: * @ list<type_expression>(integer, float) (throw: false) 
func E__star_integer2 (x EID,y EID) EID { 
    return EID{C__FLOAT,FVAL(F__star_integer2(INT(x),FLOAT(y) ))}} 
  
/* The go function for: /(x:integer,y:float) [status=0] */
func F__7_integer2 (x int,y float64) float64 { 
    return  (F_to_float(x)/y)
    } 
  
// The EID go function for: / @ list<type_expression>(integer, float) (throw: false) 
func E__7_integer2 (x EID,y EID) EID { 
    return EID{C__FLOAT,FVAL(F__7_integer2(INT(x),FLOAT(y) ))}} 
  
/* The go function for: -(x:integer,y:float) [status=0] */
func F__dash_integer3 (x int,y float64) float64 { 
    return  (F_to_float(x)-y)
    } 
  
// The EID go function for: - @ list<type_expression>(integer, float) (throw: false) 
func E__dash_integer3 (x EID,y EID) EID { 
    return EID{C__FLOAT,FVAL(F__dash_integer3(INT(x),FLOAT(y) ))}} 
  
/* The go function for: +(x:float,y:integer) [status=0] */
func F__plus_float2 (x float64,y int) float64 { 
    return  (x+F_to_float(y))
    } 
  
// The EID go function for: + @ list<type_expression>(float, integer) (throw: false) 
func E__plus_float2 (x EID,y EID) EID { 
    return EID{C__FLOAT,FVAL(F__plus_float2(FLOAT(x),INT(y) ))}} 
  
/* The go function for: *(x:float,y:integer) [status=0] */
func F__star_float2 (x float64,y int) float64 { 
    return  (x*F_to_float(y))
    } 
  
// The EID go function for: * @ list<type_expression>(float, integer) (throw: false) 
func E__star_float2 (x EID,y EID) EID { 
    return EID{C__FLOAT,FVAL(F__star_float2(FLOAT(x),INT(y) ))}} 
  
/* The go function for: /(x:float,y:integer) [status=0] */
func F__7_float2 (x float64,y int) float64 { 
    return  (x/F_to_float(y))
    } 
  
// The EID go function for: / @ list<type_expression>(float, integer) (throw: false) 
func E__7_float2 (x EID,y EID) EID { 
    return EID{C__FLOAT,FVAL(F__7_float2(FLOAT(x),INT(y) ))}} 
  
/* The go function for: -(x:float,y:integer) [status=0] */
func F__dash_float3 (x float64,y int) float64 { 
    return  (x-F_to_float(y))
    } 
  
// The EID go function for: - @ list<type_expression>(float, integer) (throw: false) 
func E__dash_float3 (x EID,y EID) EID { 
    return EID{C__FLOAT,FVAL(F__dash_float3(FLOAT(x),INT(y) ))}} 
  
// v3.4 a useful macro
/* The go function for: sqr(x:integer) [status=0] */
func F_sqr_integer (x int) int { 
    return  (x*x)
    } 
  
// The EID go function for: sqr @ integer (throw: false) 
func E_sqr_integer (x EID) EID { 
    return EID{C__INT,IVAL(F_sqr_integer(INT(x) ))}} 
  
/* The go function for: sqr(x:float) [status=0] */
func F_sqr_float (x float64) float64 { 
    return  (x*x)
    } 
  
// The EID go function for: sqr @ float (throw: false) 
func E_sqr_float (x EID) EID { 
    return EID{C__FLOAT,FVAL(F_sqr_float(FLOAT(x) ))}} 
  