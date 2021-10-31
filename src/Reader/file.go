/***** CLAIRE Compilation of file /Users/ycaseau/Dropbox/src/clairev4.0/src/meta/file.cl 
         [version 4.0.02 / safety 5] Saturday 10-30-2021 *****/

package Reader
import (_ "fmt"
	. "Kernel"
	"Core"
	"Language"
)

//-------- dumb function to prevent import errors --------
func import_g0418() { 
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
/* {1} OPT.The go function for: self_eval(self:delimiter) [] */
func (self *Delimiter ) SelfEval () EID { 
    var Result EID 
    C_reader.Next()
    Result = ToException(Core.C_general_error.Make(MakeString("[117] loose delimiter ~S in program [line ~A ?]").Id(),MakeConstantList(self.Id(),MakeInteger(ClEnv.NLine).Id()).Id())).Close()
    return Result} 
  
// The EID go function for: self_eval @ delimiter (throw: true) 
func E_self_eval_delimiter (self EID) EID { 
    return /*(sm for self_eval @ delimiter= EID)*/ ToDelimiter(OBJ(self)).SelfEval( )} 
  
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
/* {1} OPT.The go function for: useless_c(r:integer) [] */
func F_useless_c_integer (r int) *ClaireBoolean  { 
    // use function body compiling 
if (r == 10) /* If:2 */{ 
      ClEnv.NLine = (ClEnv.NLine+1)
      /* If-2 */} 
    return  MakeBoolean((r == C_reader.Space) || 
    (r == 10) || 
    (r == 13) || 
    (r == 32) || 
    (r == 160) || 
    (r == C_reader.Tab))
    } 
  
// The EID go function for: useless_c @ integer (throw: false) 
func E_useless_c_integer (r EID) EID { 
    return EID{/*(sm for useless_c @ integer= boolean)*/ F_useless_c_integer(INT(r) ).Id(),0}} 
  
// take care of PC format (10 + 13)
/* {1} OPT.The go function for: skipc(self:meta_reader) [] */
func (self *MetaReader ) Skipc () int { 
    // use function body compiling 
for (F_useless_c_integer(self.Firstc()) == CTRUE) /* while:2 */{ 
      /* Let:3 */{ 
        var b *ClaireBoolean   = Equal(MakeInteger(self.Firstc()).Id(),MakeInteger(10).Id())
        /* noccur = 1 */
        self.Next()
        if ((b == CTRUE) && 
            (self.Firstc() == 13)) /* If:4 */{ 
          self.Next()
          /* If-4 */} 
        /* Let-3 */} 
      /* while-2 */} 
    return  self.Firstc()
    } 
  
// The EID go function for: skipc @ meta_reader (throw: false) 
func E_skipc_meta_reader (self EID) EID { 
    return EID{C__INT,IVAL(/*(sm for skipc @ meta_reader= integer)*/ ToMetaReader(OBJ(self)).Skipc( ))}} 
  
// look for a meaningful termination char such as ) or ]
/* {1} OPT.The go function for: skipc!(r:meta_reader) [] */
func (r *MetaReader ) Skipc_I () EID { 
    var Result EID 
    /* Let:2 */{ 
      var c int  = r.Skipc()
      /* noccur = 3 */
      if (c == 59) /* If:3 */{ 
        for ((r.Firstc() != r.Eof) && 
            (r.Firstc() != 10)) /* while:4 */{ 
          r.Next()
          /* while-4 */} 
        if (r.Firstc() == r.Eof) /* If:4 */{ 
          Result = EID{Language.C_EOF.Value,0}
          } else {
          ClEnv.NLine = (ClEnv.NLine+1)
          Result = r.Skipc_I()
          /* If-4 */} 
        /* If!3 */}  else if (c == 47) /* If:3 */{ 
        /* Let:4 */{ 
          var x *ClaireAny  
          /* noccur = 1 */
          var x_try04215 EID 
          x_try04215 = r.Fromp.ReadIdent()
          /* ERROR PROTECTION INSERTED (x-Result) */
          if ErrorIn(x_try04215) {Result = x_try04215
          } else {
          x = ANY(x_try04215)
          if (C_string.Id() == x.Isa.Id()) /* If:5 */{ 
            Result = r.Skipc_I()
            } else {
            Result = EID{C__INT,IVAL(47)}
            /* If-5 */} 
          }
          /* Let-4 */} 
        } else {
        Result = EID{C__INT,IVAL(c)}
        /* If-3 */} 
      /* Let-2 */} 
    return RangeCheck(ToType(C_integer.Id()),Result)} 
  
// The EID go function for: skipc! @ meta_reader (throw: true) 
func E_skipc_I_meta_reader (r EID) EID { 
    return /*(sm for skipc! @ meta_reader= EID)*/ ToMetaReader(OBJ(r)).Skipc_I( )} 
  
/* {1} OPT.The go function for: cnext(self:meta_reader) [] */
func (self *MetaReader ) Cnext () *MetaReader  { 
    // use function body compiling 
self.Next()
    return  self
    } 
  
// The EID go function for: cnext @ meta_reader (throw: false) 
func E_cnext_meta_reader (self EID) EID { 
    return EID{/*(sm for cnext @ meta_reader= meta_reader)*/ ToMetaReader(OBJ(self)).Cnext( ).Id(),0}} 
  
/* {1} OPT.The go function for: findeol(self:meta_reader) [] */
func (self *MetaReader ) Findeol () *ClaireBoolean  { 
    // procedure body with s =  
var Result *ClaireBoolean  
    /* or:2 */{ 
      var v_or2 *ClaireBoolean  
      
      v_or2= CFALSE
      for (F_useless_c_integer(self.Firstc()) == CTRUE) /* while:3 */{ 
        if (self.Firstc() == 10) /* If:4 */{ 
           /*v = v_or2, s =boolean*/
v_or2 = CTRUE
          break
          /* If-4 */} 
        self.Next()
        /* while-3 */} 
      if (v_or2 == CTRUE) {Result = CTRUE
      } else /* or:3 */{ 
        v_or2 = Equal(MakeInteger(self.Firstc()).Id(),MakeInteger(self.Eof).Id())
        if (v_or2 == CTRUE) {Result = CTRUE
        } else /* or:4 */{ 
          Result = CFALSE/* org-4 */} 
        /* org-3 */} 
      /* or-2 */} 
    return Result} 
  
// The EID go function for: findeol @ meta_reader (throw: false) 
func E_findeol_meta_reader (self EID) EID { 
    return EID{/*(sm for findeol @ meta_reader= boolean)*/ ToMetaReader(OBJ(self)).Findeol( ).Id(),0}} 
  
// safety checking
//
/* {1} OPT.The go function for: checkno(r:meta_reader,n:integer,y:any) [] */
func (r *MetaReader ) Checkno (n int,y *ClaireAny ) EID { 
    var Result EID 
    if (r.Firstc() != n) /* If:2 */{ 
      Result = EID{r.Id(),0}
      } else {
      Result = F_Serror_string(MakeString("[118] read wrong char ~S after ~S"),MakeConstantList(MakeChar(F_char_I_integer(n)).Id(),y))
      /* If-2 */} 
    return Result} 
  
// The EID go function for: checkno @ meta_reader (throw: true) 
func E_checkno_meta_reader (r EID,n EID,y EID) EID { 
    return /*(sm for checkno @ meta_reader= EID)*/ ToMetaReader(OBJ(r)).Checkno(INT(n),ANY(y) )} 
  
// reads a keyword inside a control structure (used in Reader + OFTO)
//
/* {1} OPT.The go function for: verify(t:any,x:any,y:any) [] */
func F_verify_any (t *ClaireAny ,x *ClaireAny ,y *ClaireAny ) EID { 
    var Result EID 
    var g0422I *ClaireBoolean  
    var g0422I_try04232 EID 
    g0422I_try04232 = Core.F_BELONG(x,t)
    /* ERROR PROTECTION INSERTED (g0422I-Result) */
    if ErrorIn(g0422I_try04232) {Result = g0422I_try04232
    } else {
    g0422I = ToBoolean(OBJ(g0422I_try04232))
    if (g0422I == CTRUE) /* If:2 */{ 
      Result = x.ToEID()
      } else {
      Result = F_Serror_string(MakeString("[119] read ~S instead of a ~S in a ~S"),MakeConstantList(x,t,y))
      /* If-2 */} 
    }
    return Result} 
  
// The EID go function for: verify @ any (throw: true) 
func E_verify_any (t EID,x EID,y EID) EID { 
    return /*(sm for verify @ any= EID)*/ F_verify_any(ANY(t),ANY(x),ANY(y) )} 
  
// prints a syntax error
//
/* {1} OPT.The go function for: Serror(s:string,la:list) [] */
func F_Serror_string (s *ClaireString ,la *ClaireList ) EID { 
    var Result EID 
    PRINC("---- Syntax Error[line: ")
    F_princ_integer(ClEnv.NLine)
    PRINC("]:\n")
    F_flush_port(C_reader.Fromp)
    Result = ToException(Core.C_general_error.Make((s).Id(),la.Id())).Close()
    return Result} 
  
// The EID go function for: Serror @ string (throw: true) 
func E_Serror_string (s EID,la EID) EID { 
    return /*(sm for Serror @ string= EID)*/ F_Serror_string(ToString(OBJ(s)),ToList(OBJ(la)) )} 
  
// the reader-------------------------------------------------------------
//
// variable handling -------------------------------------------------
// reads a variable
//
/* {1} OPT.The go function for: extract_variable(self:any) [] */
func F_extract_variable_any (self *ClaireAny ) EID { 
    var Result EID 
    var g0425I *ClaireBoolean  
    if (self.Isa.IsIn(C_Variable) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0424 *ClaireVariable   = To_Variable(self)
        /* noccur = 2 */
        g0425I = Core.F__I_equal_any(g0424.Pname.Get(),g0424.Id())
        /* Let-3 */} 
      } else {
      g0425I = CFALSE
      /* If-2 */} 
    if (g0425I == CTRUE) /* If:2 */{ 
      /* update:3 */{ 
        var va_arg1 *ClaireVariable  
        var va_arg2 *ClaireType  
        va_arg1 = To_Variable(self)
        var va_arg2_try04264 EID 
        va_arg2_try04264 = Language.F_extract_type_any(ANY(Core.F_CALL(C_range,ARGS(self.ToEID()))))
        /* ERROR PROTECTION INSERTED (va_arg2-Result) */
        if ErrorIn(va_arg2_try04264) {Result = va_arg2_try04264
        } else {
        va_arg2 = ToType(OBJ(va_arg2_try04264))
        /* ---------- now we compile update range(va_arg1) := va_arg2 ------- */
        va_arg1.Range = va_arg2
        Result = EID{va_arg2.Id(),0}
        }
        /* update-3 */} 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = self.ToEID()
      }
      } else {
      /* Let:3 */{ 
        var v *ClaireVariable  
        /* noccur = 2 */
        var v_try04274 EID 
        /* Let:4 */{ 
          var _CL_obj *ClaireVariable   = To_Variable(new(ClaireVariable).Is(C_Variable))
          /* noccur = 3 */
          /* update:5 */{ 
            var va_arg1 *ClaireVariable  
            var va_arg2 *ClaireSymbol  
            va_arg1 = _CL_obj
            var va_arg2_try04286 EID 
            va_arg2_try04286 = Language.F_extract_symbol_any(self)
            /* ERROR PROTECTION INSERTED (va_arg2-v_try04274) */
            if ErrorIn(va_arg2_try04286) {v_try04274 = va_arg2_try04286
            } else {
            va_arg2 = ToSymbol(OBJ(va_arg2_try04286))
            /* ---------- now we compile update mClaire/pname(va_arg1) := va_arg2 ------- */
            va_arg1.Pname = va_arg2
            v_try04274 = EID{va_arg2.Id(),0}
            }
            /* update-5 */} 
          /* ERROR PROTECTION INSERTED (v_try04274-v_try04274) */
          if !ErrorIn(v_try04274) {
          v_try04274 = EID{_CL_obj.Id(),0}
          }
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (v-Result) */
        if ErrorIn(v_try04274) {Result = v_try04274
        } else {
        v = To_Variable(OBJ(v_try04274))
        C_reader.LastForm = v.Id()
        Result = EID{v.Id(),0}
        }
        /* Let-3 */} 
      /* If-2 */} 
    return Result} 
  
// The EID go function for: extract_variable @ any (throw: true) 
func E_extract_variable_any (self EID) EID { 
    return /*(sm for extract_variable @ any= EID)*/ F_extract_variable_any(ANY(self) )} 
  
// create a variable and add it to the lexical environment
/* {1} OPT.The go function for: bind!(self:meta_reader,%v:Variable) [] */
func (self *MetaReader ) Bind_I (_Zv *ClaireVariable ) *ClaireList  { 
    // procedure body with s =  
var Result *ClaireList  
    _Zv.Index = self.Index
    /* Let:2 */{ 
      var value *ClaireAny   = _Zv.Pname.Get()
      /* noccur = 1 */
      self.Index = (self.Index+1)
      if (self.Index > self.Maxstack) /* If:3 */{ 
        self.Maxstack = self.Index
        /* If-3 */} 
      _Zv.Pname.Put(_Zv.Id())
      Result = MakeConstantList(_Zv.Id(),value)
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: bind! @ meta_reader (throw: false) 
func E_bind_I_meta_reader (self EID,_Zv EID) EID { 
    return EID{/*(sm for bind! @ meta_reader= list)*/ ToMetaReader(OBJ(self)).Bind_I(To_Variable(OBJ(_Zv)) ).Id(),0}} 
  
// remove a variable from the lexical environment
//
/* {1} OPT.The go function for: unbind!(self:meta_reader,%first:list) [] */
func (self *MetaReader ) Unbind_I (_Zfirst *ClaireList ) *ClaireAny  { 
    // procedure body with s =  
var Result *ClaireAny  
    /* Let:2 */{ 
      var ClaireVar *ClaireAny   = _Zfirst.At(1-1)
      /* noccur = 1 */
      self.Index = (self.Index-1)
      Result = ToSymbol(OBJ(Core.F_CALL(C_mClaire_pname,ARGS(ClaireVar.ToEID())))).Put(_Zfirst.At(2-1))
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: unbind! @ meta_reader (throw: false) 
func E_unbind_I_meta_reader (self EID,_Zfirst EID) EID { 
    return /*(sm for unbind! @ meta_reader= any)*/ ToMetaReader(OBJ(self)).Unbind_I(ToList(OBJ(_Zfirst)) ).ToEID()} 
  
// declaration of the CLAIRE standard ports ----------------------------
// we create global variables - however they exists as properties of system => redundant ?
//
/* {1} OPT.The go function for: /(s:string,s2:string) [] */
func F__7_string (s *ClaireString ,s2 *ClaireString ) *ClaireString  { 
    // use function body compiling 
return  F_append_string(F_append_string(s,ToString(C__starfs_star.Value)),s2)
    } 
  
// The EID go function for: / @ string (throw: false) 
func E__7_string (s EID,s2 EID) EID { 
    return EID{/*(sm for / @ string= string)*/ F__7_string(ToString(OBJ(s)),ToString(OBJ(s2)) ).Id(),0}} 
  
// basic methods defined in creader.c -----------------------------------
// TODO move!
// flush(self:port) : any -> function!(flush_port)
// this function is called by the main and restores the reader in a good shape. Also
// closes the input port to free the associated file ! <yc>
/* {1} OPT.The go function for: mClaire/restore_state(self:meta_reader) [] */
func (self *MetaReader ) RestoreState ()  { 
    // procedure body with s =  
if (Equal(self.Fromp.Id(),C_stdin.Value) != CTRUE) /* If:2 */{ 
      self.Fromp.Fclose()
      /* If-2 */} 
    self.Fromp = ToPort(C_stdin.Value)
    self.Index = 1
    F_pushback_port(ToPort(C_stdin.Value),32)
    F_restore_state_void()
    } 
  
// The EID go function for: mClaire/restore_state @ meta_reader (throw: false) 
func E_restore_state_meta_reader (self EID) EID { 
    /*(sm for mClaire/restore_state @ meta_reader= void)*/ ToMetaReader(OBJ(self)).RestoreState( )
    return EVOID} 
  
// *********************************************************************
// *   Part 2: Loading                                                 *
// *********************************************************************
// sload is the interactive version (when b = true).
//
/* {1} OPT.The go function for: load_file(self:string,b:boolean) [] */
func F_load_file_string (self *ClaireString ,b *ClaireBoolean ) EID { 
    var Result EID 
    C_reader.Index = 0
    C_reader.Maxstack = 0
    ClEnv.NLine = 1
    C_reader.External = self
    Core.F_tformat_string(MakeString("---- [load CLAIRE file: ~A]\n"),2,MakeConstantList((self).Id()))
    /* Let:2 */{ 
      var s2 *ClaireString   = F_append_string(self,MakeString(".cl"))
      /* noccur = 1 */
      /* Let:3 */{ 
        var p1 *ClairePort  
        /* noccur = 4 */
        var p1_try04314 EID 
        h_index := ClEnv.Index /* Handle */
        h_base := ClEnv.Base
        p1_try04314 = F_fopen_string(s2,MakeString("r"))
        if ErrorIn(p1_try04314){ 
          /* s=EID */ClEnv.Index = h_index
          ClEnv.Base = h_base
          h_index := ClEnv.Index /* Handle */
          h_base := ClEnv.Base
          p1_try04314 = F_fopen_string(self,MakeString("r"))
          if ErrorIn(p1_try04314){ 
            /* s=EID */ClEnv.Index = h_index
            ClEnv.Base = h_base
            p1_try04314 = ToException(Core.C_general_error.Make(MakeString("[120] the file ~A cannot be opened").Id(),MakeConstantList((self).Id()).Id())).Close()
            } 
          } 
        /* ERROR PROTECTION INSERTED (p1-Result) */
        if ErrorIn(p1_try04314) {Result = p1_try04314
        } else {
        p1 = ToPort(OBJ(p1_try04314))
        /* Let:4 */{ 
          var start int  = ClEnv.Base
          /* noccur = 1 */
          /* Let:5 */{ 
            var top int  = ClEnv.Index
            /* noccur = 3 */
            /* Let:6 */{ 
              var p2 *ClairePort   = C_reader.Fromp
              /* noccur = 1 */
              /* Let:7 */{ 
                var b2 *ClaireBoolean   = C_reader.Toplevel
                /* noccur = 1 */
                /* Let:8 */{ 
                  var _staritem_star *ClaireAny   = CNULL
                  /* noccur = 10 */
                  ClEnv.Base= top
                  C_reader.Fromp = p1
                  C_reader.Toplevel = CFALSE
                  var _staritem_star_try04329 EID 
                  _staritem_star_try04329 = F_readblock_port(p1)
                  /* ERROR PROTECTION INSERTED (_staritem_star-Result) */
                  if ErrorIn(_staritem_star_try04329) {Result = _staritem_star_try04329
                  } else {
                  _staritem_star = ANY(_staritem_star_try04329)
                  Result = _staritem_star.ToEID()
                  Result= EID{CFALSE.Id(),0}
                  for (_staritem_star != C_Reader_eof.Id()) /* while:9 */{ 
                    var void_try10 EID 
                    _ = void_try10
                    { 
                    if (b == CTRUE) /* If:10 */{ 
                      PRINC("<")
                      F_princ_integer(ClEnv.NLine)
                      PRINC(">:")
                      void_try10 = Core.F_CALL(C_print,ARGS(_staritem_star.ToEID()))
                      /* ERROR PROTECTION INSERTED (void_try10-void_try10) */
                      if ErrorIn(void_try10) {Result = void_try10
                      break
                      } else {
                      PRINC("\n")
                      void_try10 = EVOID
                      }
                      } else {
                      void_try10 = EID{CFALSE.Id(),0}
                      /* If-10 */} 
                    /* ERROR PROTECTION INSERTED (void_try10-void_try10) */
                    if ErrorIn(void_try10) {Result = void_try10
                    break
                    } else {
                    ClEnv.Index= (top+(C_reader.Maxstack+1))
                    if (C_string.Id() == _staritem_star.Isa.Id()) /* If:10 */{ 
                      /* Let:11 */{ 
                        var g0429 *ClaireString   = ToString(_staritem_star)
                        /* noccur = 2 */
                        if (ToBoolean(Language.C_NeedComment.Value) == CTRUE) /* If:12 */{ 
                          if (Language.C_iClaire_LastComment.Value != CNULL) /* If:13 */{ 
                            var v_gassign14 *ClaireAny  
                            v_gassign14 = ANY(Core.F_CALL(ToProperty(C__7_plus.Id()),ARGS(EID{Language.C_iClaire_LastComment.Value,0},EID{F_append_string(MakeString("\n-- "),g0429).Id(),0})))
                            Language.C_iClaire_LastComment.Value = v_gassign14
                            void_try10 = v_gassign14.ToEID()
                            } else {
                            var v_gassign14 *ClaireAny  
                            v_gassign14 = (F_append_string(F_append_string(F_append_string(F_append_string(F_append_string(MakeString("["),C_reader.External),MakeString("(")),F_string_I_integer(ClEnv.NLine)),MakeString(")]\n-- ")),g0429)).Id()
                            Language.C_iClaire_LastComment.Value = v_gassign14
                            void_try10 = v_gassign14.ToEID()
                            /* If-13 */} 
                          } else {
                          void_try10 = EID{CFALSE.Id(),0}
                          /* If-12 */} 
                        /* Let-11 */} 
                      } else {
                      var _staritem_star_try043311 EID 
                      _staritem_star_try043311 = EVAL(_staritem_star)
                      /* ERROR PROTECTION INSERTED (_staritem_star-void_try10) */
                      if ErrorIn(_staritem_star_try043311) {void_try10 = _staritem_star_try043311
                      Result = _staritem_star_try043311
                      break
                      } else {
                      _staritem_star = ANY(_staritem_star_try043311)
                      void_try10 = _staritem_star.ToEID()
                      var v_gassign11 *ClaireAny  
                      v_gassign11 = CNULL
                      Language.C_iClaire_LastComment.Value = v_gassign11
                      void_try10 = v_gassign11.ToEID()
                      }
                      /* If-10 */} 
                    /* ERROR PROTECTION INSERTED (void_try10-void_try10) */
                    if ErrorIn(void_try10) {Result = void_try10
                    break
                    } else {
                    if ((b == CTRUE) && 
                        (C_string.Id() != _staritem_star.Isa.Id())) /* If:10 */{ 
                      PRINC("=> ")
                      void_try10 = Core.F_CALL(C_print,ARGS(_staritem_star.ToEID()))
                      /* ERROR PROTECTION INSERTED (void_try10-void_try10) */
                      if ErrorIn(void_try10) {Result = void_try10
                      break
                      } else {
                      PRINC(" \n\n")
                      void_try10 = EVOID
                      }
                      } else {
                      void_try10 = EID{CFALSE.Id(),0}
                      /* If-10 */} 
                    /* ERROR PROTECTION INSERTED (void_try10-void_try10) */
                    if ErrorIn(void_try10) {Result = void_try10
                    break
                    } else {
                    var _staritem_star_try043410 EID 
                    _staritem_star_try043410 = F_readblock_port(p1)
                    /* ERROR PROTECTION INSERTED (_staritem_star-void_try10) */
                    if ErrorIn(_staritem_star_try043410) {void_try10 = _staritem_star_try043410
                    Result = _staritem_star_try043410
                    break
                    } else {
                    _staritem_star = ANY(_staritem_star_try043410)
                    void_try10 = _staritem_star.ToEID()
                    }}}}
                    /* while-9 */} 
                  }
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  ClEnv.Base= start
                  ClEnv.Index= top
                  C_reader.Toplevel = b2
                  C_reader.Fromp = p2
                  C_reader.External = MakeString("toplevel")
                  p1.Fclose()
                  Result = EVOID
                  }}
                  /* Let-8 */} 
                /* Let-7 */} 
              /* Let-6 */} 
            /* Let-5 */} 
          /* Let-4 */} 
        }
        /* Let-3 */} 
      /* Let-2 */} 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    Result = EID{CTRUE.Id(),0}
    }
    return Result} 
  
// The EID go function for: load_file @ string (throw: true) 
func E_load_file_string (self EID,b EID) EID { 
    return /*(sm for load_file @ string= EID)*/ F_load_file_string(ToString(OBJ(self)),ToBoolean(OBJ(b)) )} 
  
// the simple load
//
/* {1} OPT.The go function for: load(self:string) [] */
func F_load_string (self *ClaireString ) EID { 
    var Result EID 
    Result = F_load_file_string(self,CFALSE)
    return Result} 
  
// The EID go function for: load @ string (throw: true) 
func E_load_string (self EID) EID { 
    return /*(sm for load @ string= EID)*/ F_load_string(ToString(OBJ(self)) )} 
  
/* {1} OPT.The go function for: sload(self:string) [] */
func F_sload_string (self *ClaireString ) EID { 
    var Result EID 
    Result = F_load_file_string(self,CTRUE)
    return Result} 
  
// The EID go function for: sload @ string (throw: true) 
func E_sload_string (self EID) EID { 
    return /*(sm for sload @ string= EID)*/ F_sload_string(ToString(OBJ(self)) )} 
  
// loading a module into the system.
// The correct package is open and each file is loaded.
/* {1} OPT.The go function for: load_file(self:module,b:boolean) [] */
func F_load_file_module (self *ClaireModule ,b *ClaireBoolean ) EID { 
    var Result EID 
    if (self.Status == 2) /* If:2 */{ 
      Result = Core.F_CALL(C_funcall,ARGS(Core.F_CALL(C_mClaire_evaluate,ARGS(EID{self.Id(),0})),EID{C_any.Id(),0}))
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      /* update:3 */{ 
        var va_arg1 *ClaireModule  
        var va_arg2 int 
        va_arg1 = self
        va_arg2 = 3
        /* ---------- now we compile update mClaire/status(va_arg1) := va_arg2 ------- */
        va_arg1.Status = va_arg2
        Result = EID{C__INT,IVAL(va_arg2)}
        /* update-3 */} 
      }
      /* If!2 */}  else if ((self.Status == 0) && 
        ((self.Source).Id() != CNULL)) /* If:2 */{ 
      Core.F_tformat_string(MakeString("---- Loading the module ~S.\n"),1,MakeConstantList(self.Id()))
      self.Begin()
      /* Let:3 */{ 
        var s *ClaireString   = F_append_string(self.Source,ToString(C__starfs_star.Value))
        /* noccur = 1 */
        /* For:4 */{ 
          var x *ClaireAny  
          _ = x
          Result= EID{CFALSE.Id(),0}
          var x_support *ClaireList  
          x_support = self.MadeOf
          x_len := x_support.Length()
          for i_it := 0; i_it < x_len; i_it++ { 
            x = x_support.At(i_it)
            var void_try6 EID 
            _ = void_try6
            void_try6 = F_load_file_string(F_append_string(F_append_string(s,ToString(x)),MakeString(".cl")),b)
            /* ERROR PROTECTION INSERTED (void_try6-Result) */
            if ErrorIn(void_try6) {Result = void_try6
            Result = void_try6
            break
            } else {
            }
            /* loop-5 */} 
          /* For-4 */} 
        /* Let-3 */} 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      /* update:3 */{ 
        var va_arg1 *ClaireModule  
        var va_arg2 int 
        va_arg1 = self
        va_arg2 = 1
        /* ---------- now we compile update mClaire/status(va_arg1) := va_arg2 ------- */
        va_arg1.Status = va_arg2
        Result = EID{C__INT,IVAL(va_arg2)}
        /* update-3 */} 
      }
      } else {
      Result = EID{CFALSE.Id(),0}
      /* If-2 */} 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    self.End()
    Result = EVOID
    }
    return Result} 
  
// The EID go function for: load_file @ module (throw: true) 
func E_load_file_module (self EID,b EID) EID { 
    return /*(sm for load_file @ module= EID)*/ F_load_file_module(ToModule(OBJ(self)),ToBoolean(OBJ(b)) )} 
  
// the simple load
//
/* {1} OPT.The go function for: load(self:module) [] */
func F_load_module (self *ClaireModule ) EID { 
    var Result EID 
    /* For:2 */{ 
      var x *ClaireAny  
      _ = x
      Result= EID{CFALSE.Id(),0}
      var x_support *ClaireList  
      x_support = F_add_modules_list(MakeConstantList(self.Id()))
      x_len := x_support.Length()
      for i_it := 0; i_it < x_len; i_it++ { 
        x = x_support.At(i_it)
        var void_try4 EID 
        _ = void_try4
        void_try4 = Core.F_CALL(C_Reader_load_file,ARGS(x.ToEID(),EID{CFALSE.Id(),0}))
        /* ERROR PROTECTION INSERTED (void_try4-Result) */
        if ErrorIn(void_try4) {Result = void_try4
        Result = void_try4
        break
        } else {
        }
        /* loop-3 */} 
      /* For-2 */} 
    return Result} 
  
// The EID go function for: load @ module (throw: true) 
func E_load_module (self EID) EID { 
    return /*(sm for load @ module= EID)*/ F_load_module(ToModule(OBJ(self)) )} 
  
/* {1} OPT.The go function for: sload(self:module) [] */
func F_sload_module (self *ClaireModule ) EID { 
    var Result EID 
    /* For:2 */{ 
      var x *ClaireAny  
      _ = x
      Result= EID{CFALSE.Id(),0}
      var x_support *ClaireList  
      x_support = F_add_modules_list(MakeConstantList(self.Id()))
      x_len := x_support.Length()
      for i_it := 0; i_it < x_len; i_it++ { 
        x = x_support.At(i_it)
        var void_try4 EID 
        _ = void_try4
        void_try4 = Core.F_CALL(C_Reader_load_file,ARGS(x.ToEID(),EID{CTRUE.Id(),0}))
        /* ERROR PROTECTION INSERTED (void_try4-Result) */
        if ErrorIn(void_try4) {Result = void_try4
        Result = void_try4
        break
        } else {
        }
        /* loop-3 */} 
      /* For-2 */} 
    return Result} 
  
// The EID go function for: sload @ module (throw: true) 
func E_sload_module (self EID) EID { 
    return /*(sm for sload @ module= EID)*/ F_sload_module(ToModule(OBJ(self)) )} 
  
// This is a very important method which adds the right order the
// modules that must be loaded to load oself. the list l represents the
// list of modules that we know will be in the result. result represent
// the current list of ordered modules
//
/* {1} OPT.The go function for: add_modules(self:module,l:set,result:list) [] */
func F_add_modules_module (self *ClaireModule ,l *ClaireSet ,result *ClaireList ) *ClaireList  { 
    // use function body compiling 
if (result.Memq(self.Id()) == CTRUE) /* body If:2 */{ 
      return  result
      }  else if (l.Contain_ask(self.Id()) == CTRUE) /* body If:2 */{ 
      return  result.AddFast(self.Id())
      } else {
      l = l.AddFast(self.Id())
      /* For:3 */{ 
        var x *ClaireAny  
        _ = x
        var x_support *ClaireList  
        x_support = self.Uses
        x_len := x_support.Length()
        for i_it := 0; i_it < x_len; i_it++ { 
          x = x_support.At(i_it)
          if (x.Isa.IsIn(C_module) == CTRUE) /* If:5 */{ 
            /* Let:6 */{ 
              var g0435 *ClaireModule   = ToModule(x)
              /* noccur = 1 */
              result = F_add_modules_module(g0435,l,result)
              /* Let-6 */} 
            /* If-5 */} 
          /* loop-4 */} 
        /* For-3 */} 
      if (result.Memq(self.Id()) != CTRUE) /* If:3 */{ 
        result = result.AddFast(self.Id())
        /* If-3 */} 
      /* For:3 */{ 
        var x *ClaireAny  
        _ = x
        var x_support *ClaireList  
        x_support = self.Parts
        x_len := x_support.Length()
        for i_it := 0; i_it < x_len; i_it++ { 
          x = x_support.At(i_it)
          result = F_add_modules_module(ToModule(x),l,result)
          /* loop-4 */} 
        /* For-3 */} 
      return  result
      /* body If-2 */} 
    } 
  
// The EID go function for: add_modules @ module (throw: false) 
func E_add_modules_module (self EID,l EID,result EID) EID { 
    return EID{/*(sm for add_modules @ module= list)*/ F_add_modules_module(ToModule(OBJ(self)),ToSet(OBJ(l)),ToList(OBJ(result)) ).Id(),0}} 
  
// this methods takes a list of modules that must be loaded and returns
// a list of modules that are necessary for the definition
//
/* {1} OPT.The go function for: add_modules(self:list) [] */
func F_add_modules_list (self *ClaireList ) *ClaireList  { 
    // procedure body with s =  
var Result *ClaireList  
    /* Let:2 */{ 
      var l *ClaireList   = ToType(C_module.Id()).EmptyList()
      /* noccur = 4 */
      /* For:3 */{ 
        var x *ClaireAny  
        _ = x
        var x_support *ClaireList  
        x_support = self
        x_len := x_support.Length()
        for i_it := 0; i_it < x_len; i_it++ { 
          x = x_support.At(i_it)
          l = F_add_modules_module(ToModule(x),l.Set_I(),l)
          /* loop-4 */} 
        /* For-3 */} 
      Result = l
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: add_modules @ list (throw: false) 
func E_add_modules_list (self EID) EID { 
    return EID{/*(sm for add_modules @ list= list)*/ F_add_modules_list(ToList(OBJ(self)) ).Id(),0}} 
  
// load a file of expressions (quite useful)
/* {1} OPT.The go function for: eload(self:string) [] */
func F_eload_string (self *ClaireString ) EID { 
    var Result EID 
    C_reader.Index = 0
    C_reader.Maxstack = 0
    ClEnv.NLine = 1
    C_reader.External = self
    Core.F_tformat_string(MakeString("---- [eload CLAIRE file: ~A]\n"),2,MakeConstantList((self).Id()))
    /* Let:2 */{ 
      var s2 *ClaireString   = F_append_string(self,MakeString(".cl"))
      /* noccur = 1 */
      /* Let:3 */{ 
        var p0 *ClairePort   = C_reader.Fromp
        /* noccur = 1 */
        /* Let:4 */{ 
          var p1 *ClairePort  
          /* noccur = 4 */
          var p1_try04365 EID 
          h_index := ClEnv.Index /* Handle */
          h_base := ClEnv.Base
          p1_try04365 = F_fopen_string(s2,MakeString("r"))
          if ErrorIn(p1_try04365){ 
            /* s=EID */ClEnv.Index = h_index
            ClEnv.Base = h_base
            h_index := ClEnv.Index /* Handle */
            h_base := ClEnv.Base
            p1_try04365 = F_fopen_string(self,MakeString("r"))
            if ErrorIn(p1_try04365){ 
              /* s=EID */ClEnv.Index = h_index
              ClEnv.Base = h_base
              p1_try04365 = ToException(Core.C_general_error.Make(MakeString("[120] the file ~A cannot be opened").Id(),MakeConstantList((self).Id()).Id())).Close()
              } 
            } 
          /* ERROR PROTECTION INSERTED (p1-Result) */
          if ErrorIn(p1_try04365) {Result = p1_try04365
          } else {
          p1 = ToPort(OBJ(p1_try04365))
          /* Let:5 */{ 
            var start int  = ClEnv.Base
            /* noccur = 1 */
            /* Let:6 */{ 
              var top int  = ClEnv.Index
              /* noccur = 3 */
              /* Let:7 */{ 
                var b2 *ClaireBoolean   = C_reader.Toplevel
                /* noccur = 1 */
                /* Let:8 */{ 
                  var _staritem_star *ClaireAny   = CNULL
                  /* noccur = 5 */
                  ClEnv.Base= top
                  C_reader.Toplevel = CFALSE
                  C_reader.Fromp = p1
                  var _staritem_star_try04379 EID 
                  _staritem_star_try04379 = F_read_port(p1)
                  /* ERROR PROTECTION INSERTED (_staritem_star-Result) */
                  if ErrorIn(_staritem_star_try04379) {Result = _staritem_star_try04379
                  } else {
                  _staritem_star = ANY(_staritem_star_try04379)
                  Result = _staritem_star.ToEID()
                  Result= EID{CFALSE.Id(),0}
                  for (_staritem_star != C_Reader_eof.Id()) /* while:9 */{ 
                    var void_try10 EID 
                    _ = void_try10
                    { 
                    ClEnv.Index= (top+(C_reader.Maxstack+1))
                    var _staritem_star_try043810 EID 
                    _staritem_star_try043810 = EVAL(_staritem_star)
                    /* ERROR PROTECTION INSERTED (_staritem_star-void_try10) */
                    if ErrorIn(_staritem_star_try043810) {void_try10 = _staritem_star_try043810
                    Result = _staritem_star_try043810
                    break
                    } else {
                    _staritem_star = ANY(_staritem_star_try043810)
                    void_try10 = _staritem_star.ToEID()
                    var _staritem_star_try043910 EID 
                    _staritem_star_try043910 = F_read_port(p1)
                    /* ERROR PROTECTION INSERTED (_staritem_star-void_try10) */
                    if ErrorIn(_staritem_star_try043910) {void_try10 = _staritem_star_try043910
                    Result = _staritem_star_try043910
                    break
                    } else {
                    _staritem_star = ANY(_staritem_star_try043910)
                    void_try10 = _staritem_star.ToEID()
                    }}
                    /* while-9 */} 
                  }
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  ClEnv.Base= start
                  ClEnv.Index= top
                  C_reader.Fromp = p0
                  C_reader.Toplevel = b2
                  C_reader.External = MakeString("toplevel")
                  p1.Fclose()
                  Result = EVOID
                  }}
                  /* Let-8 */} 
                /* Let-7 */} 
              /* Let-6 */} 
            /* Let-5 */} 
          }
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    Result = EID{CTRUE.Id(),0}
    }
    return Result} 
  
// The EID go function for: eload @ string (throw: true) 
func E_eload_string (self EID) EID { 
    return /*(sm for eload @ string= EID)*/ F_eload_string(ToString(OBJ(self)) )} 
  
// *********************************************************************
// *   Part 3: Read & Top-level                                        *
// *********************************************************************
// The standard read function.
// This method reads from a CLAIRE port (self).
// We first check if self is the current reading port.
// the last character read (and not used) is in last(reader)
/* {1} OPT.The go function for: readblock(p:port) [] */
func F_readblock_port (p *ClairePort ) EID { 
    var Result EID 
    if (Equal(C_reader.Fromp.Id(),p.Id()) == CTRUE) /* If:2 */{ 
      Result = C_reader.Nextunit()
      } else {
      /* Let:3 */{ 
        var p2 *ClairePort   = C_reader.Fromp
        /* noccur = 1 */
        C_reader.Fromp = p
        /* Let:4 */{ 
          var val *ClaireAny  
          /* noccur = 6 */
          var val_try04405 EID 
          val_try04405 = C_reader.Nextunit()
          /* ERROR PROTECTION INSERTED (val-Result) */
          if ErrorIn(val_try04405) {Result = val_try04405
          } else {
          val = ANY(val_try04405)
          C_reader.Fromp = p2
          if ((Equal(val,C_reader.Paren) == CTRUE) || 
              ((Equal(val,C_reader.Curly) == CTRUE) || 
                ((Equal(val,C_reader.Comma) == CTRUE) || 
                  (Equal(val,C_reader.Bracket) == CTRUE)))) /* If:5 */{ 
            Result = F_Serror_string(MakeString("[117] Loose ~S in file"),MakeConstantList(val))
            } else {
            Result = EID{CFALSE.Id(),0}
            /* If-5 */} 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          Result = val.ToEID()
          }
          }
          /* Let-4 */} 
        /* Let-3 */} 
      /* If-2 */} 
    return Result} 
  
// The EID go function for: readblock @ port (throw: true) 
func E_readblock_port (p EID) EID { 
    return /*(sm for readblock @ port= EID)*/ F_readblock_port(ToPort(OBJ(p)) )} 
  
// read reads a closed expression
/* {1} OPT.The go function for: read(p:port) [] */
func F_read_port (p *ClairePort ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var p2 *ClairePort   = C_reader.Fromp
      /* noccur = 3 */
      if (Equal(p.Id(),p2.Id()) != CTRUE) /* If:3 */{ 
        C_reader.Fromp = p
        /* If-3 */} 
      /* Let:3 */{ 
        var val *ClaireAny  
        /* noccur = 6 */
        var val_try04414 EID 
        if (C_reader.Skipc() == C_reader.Eof) /* If:4 */{ 
          val_try04414 = EID{C_Reader_eof.Id(),0}
          } else {
          val_try04414 = C_reader.Nexte()
          /* If-4 */} 
        /* ERROR PROTECTION INSERTED (val-Result) */
        if ErrorIn(val_try04414) {Result = val_try04414
        } else {
        val = ANY(val_try04414)
        if (Equal(p.Id(),p2.Id()) != CTRUE) /* If:4 */{ 
          C_reader.Fromp = p2
          /* If-4 */} 
        if ((Equal(val,C_reader.Paren) == CTRUE) || 
            ((Equal(val,C_reader.Curly) == CTRUE) || 
              ((Equal(val,C_reader.Comma) == CTRUE) || 
                (Equal(val,C_reader.Bracket) == CTRUE)))) /* If:4 */{ 
          Result = F_Serror_string(MakeString("[117] Loose ~S in file"),MakeConstantList(val))
          } else {
          Result = EID{CFALSE.Id(),0}
          /* If-4 */} 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        Result = val.ToEID()
        }
        }
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: read @ port (throw: true) 
func E_read_port (p EID) EID { 
    return /*(sm for read @ port= EID)*/ F_read_port(ToPort(OBJ(p)) )} 
  
// read into a string
/* {1} OPT.The go function for: read(self:string) [] */
func F_read_string (self *ClaireString ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var b *ClaireBoolean   = C_reader.Toplevel
      /* noccur = 1 */
      /* Let:3 */{ 
        var p *ClairePort   = C_reader.Fromp
        /* noccur = 2 */
        /* Let:4 */{ 
          var x *ClaireAny   = CNULL
          /* noccur = 2 */
          C_reader.Toplevel = CTRUE
          C_reader.Fromp = F_port_I_string(self)
          if (ClEnv.Verbose == 2) /* If:5 */{ 
            PRINC("-- start read in string with nextunit ---\n")
            /* If-5 */} 
          h_index := ClEnv.Index /* Handle */
          h_base := ClEnv.Base
          var x_try04425 EID 
          x_try04425 = C_reader.Nextunit()
          /* ERROR PROTECTION INSERTED (x-Result) */
          if ErrorIn(x_try04425) {Result = x_try04425
          } else {
          x = ANY(x_try04425)
          Result = x.ToEID()
          /* update:5 */{ 
            var va_arg1 *MetaReader  
            var va_arg2 *ClairePort  
            va_arg1 = C_reader
            va_arg2 = p
            /* ---------- now we compile update fromp(va_arg1) := va_arg2 ------- */
            va_arg1.Fromp = va_arg2
            Result = va_arg2.ToEID()
            /* update-5 */} 
          }
          if ErrorIn(Result){ 
            /* s=EID */ClEnv.Index = h_index
            ClEnv.Base = h_base
            C_reader.Fromp = p
            Result = ClEnv.Exception_I.Close()
            } 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          C_reader.Toplevel = b
          Result = x.ToEID()
          }
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: read @ string (throw: true) 
func E_read_string (self EID) EID { 
    return /*(sm for read @ string= EID)*/ F_read_string(ToString(OBJ(self)) )} 
  
// used by the top level
// calls the debugger
/* {1} OPT.The go function for: debug_if_possible(_CL_obj:void) [] */
func F_debug_if_possible_void () EID { 
    var Result EID 
    if (ClEnv.Debug_I >= 0) /* If:2 */{ 
      Result = F_funcall1(ToMethod(C_call_debug.Restrictions.ValuesO()[1-1]).Functional,ClEnv.Id())
      } else {
      Result = F_print_exception_void().ToEID()
      /* If-2 */} 
    return Result} 
  
// The EID go function for: debug_if_possible @ void (throw: true) 
func E_debug_if_possible_void (_CL_obj EID) EID { 
    return /*(sm for debug_if_possible @ void= EID)*/ F_debug_if_possible_void( )} 
  
// a method for calling the printer without issuing a message (that would
// modify the stack and make debugging impossible).
// here we assume that self_print is always defined and is always a compiled
// function
/* {1} OPT.The go function for: print_exception(_CL_obj:void) [] */
func F_print_exception_void () *ClaireAny  { 
    // procedure body with s =  
var Result *ClaireAny  
    /* Let:2 */{ 
      var p *ClairePort   = ToPort(C_stdout.Value).UseAsOutput()
      /* noccur = 1 */
      /* Let:3 */{ 
        var _Zerr *ClaireException   = ClEnv.Exception_I
        /* noccur = 3 */
        /* Let:4 */{ 
          var _Zprop *ClaireMethod   = ToMethod(Core.F__at_property1(C_self_print,_Zerr.Id().Isa).Id())
          /* noccur = 3 */
          var Result_try5 EID 
          h_index := ClEnv.Index /* Handle */
          h_base := ClEnv.Base
          if (_Zprop.Functional.Id() != CNULL) /* If:5 */{ 
            Result_try5 = F_funcall1(_Zprop.Functional,_Zerr.Id())
            } else {
            Result_try5 = Core.F_funcall_method1(_Zprop,_Zerr.Id())
            /* If-5 */} 
          if ErrorIn(Result_try5){ 
            /* s=void */ClEnv.Index = h_index
            ClEnv.Base = h_base
            PRINC("****** ERROR[121]: unprintable error has occurred.\n")
            } 
          Result = p.UseAsOutput().Id()
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: print_exception @ void (throw: false) 
func E_print_exception_void (_CL_obj EID) EID { 
    return /*(sm for print_exception @ void= any)*/ F_print_exception_void( ).ToEID()} 
  
// **********************************************************************
// *  Part 4: The show & kill methods + macro-methods                   *
// **********************************************************************
//----------------- printing an object -------------------------
// %show is an open restriction which allow to show the value of a
// binary relation
//
// this method is the basic method called for show(..)
//
/* {1} OPT.The go function for: show(self:any) [] */
func F_show_any (self *ClaireAny ) EID { 
    var Result EID 
    if (Core.F__Z_any1(self,C_object) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0443 *ClaireObject   = ToObject(self)
        /* noccur = 2 */
        /* For:4 */{ 
          var rel *ClaireAny  
          _ = rel
          Result= EID{CFALSE.Id(),0}
          var rel_support *ClaireList  
          rel_support = g0443.Id().Isa.Slots
          for _,rel = range(rel_support.ValuesO())/* loop2:5 */{ 
            var void_try6 EID 
            _ = void_try6
            { 
            void_try6 = Core.F_print_any(ToRestriction(rel).Selector.Id())
            /* ERROR PROTECTION INSERTED (void_try6-void_try6) */
            if ErrorIn(void_try6) {Result = void_try6
            break
            } else {
            PRINC(": ")
            void_try6 = Core.F_CALL(C_print,ARGS(Core.F_get_slot(ToSlot(rel),g0443).ToEID()))
            /* ERROR PROTECTION INSERTED (void_try6-void_try6) */
            if ErrorIn(void_try6) {Result = void_try6
            break
            } else {
            PRINC("\n")
            }}
            }
            /* loop-5 */} 
          /* For-4 */} 
        /* Let-3 */} 
      } else {
      Result = Core.F_CALL(C_print,ARGS(self.ToEID()))
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(" is a ")
      Result = Core.F_print_any(self.Isa.Id())
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC("\n")
      Result = EVOID
      }}
      /* If-2 */} 
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    Result = EID{CTRUE.Id(),0}
    }
    return Result} 
  
// The EID go function for: show @ any (throw: true) 
func E_show_any (self EID) EID { 
    return /*(sm for show @ any= EID)*/ F_show_any(ANY(self) )} 
  
// This is the good version of kill, the nasty one is dangerous ....
// these restrictions of kill explain the dependencies among objects
//
/* {1} OPT.The go function for: kill(self:object) [] */
func F_kill_object (self *ClaireObject ) *ClaireAny  { 
    // use function body compiling 
if (self.Isa.IsIn(C_thing) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0445 *ClaireThing   = ToThing(self.Id())
        /* noccur = 1 */
        g0445.Name.Put(CNULL)
        /* Let-3 */} 
      /* If-2 */} 
    self.Isa.Instances = self.Isa.Instances.Delete(self.Id())
    return  CEMPTY.Id()
    } 
  
// The EID go function for: kill @ object (throw: false) 
func E_kill_object (self EID) EID { 
    return /*(sm for kill @ object= any)*/ F_kill_object(ToObject(OBJ(self)) ).ToEID()} 
  
/* {1} OPT.The go function for: kill(self:class) [] */
func F_kill_class (self *ClaireClass ) *ClaireAny  { 
    // use function body compiling 
for (self.Instances.Length() != 0) /* while:2 */{ 
      Core.F_CALL(C_kill,ARGS(self.Instances.At(1-1).ToEID()))
      /* while-2 */} 
    /* For:2 */{ 
      var x *ClaireAny  
      _ = x
      for _,x = range(self.Descendents.Values)/* loop:3 */{ 
        if (ToClass(x).Superclass.Id() == self.Id()) /* If:4 */{ 
          F_kill_class(ToClass(x))
          /* If-4 */} 
        /* loop-3 */} 
      /* For-2 */} 
    return  F_kill_object(ToObject(self.Id()))
    } 
  
// The EID go function for: kill @ class (throw: false) 
func E_kill_class (self EID) EID { 
    return /*(sm for kill @ class= any)*/ F_kill_class(ToClass(OBJ(self)) ).ToEID()} 
  
// our two very special inline methods
/* {1} OPT.The go function for: min(x:integer,y:integer) [] */
func F_min_integer (x int,y int) int { 
    // use function body compiling 
if (x <= y) /* body If:2 */{ 
      return  x
      } else {
      return  y
      /* body If-2 */} 
    } 
  
// The EID go function for: min @ integer (throw: false) 
func E_min_integer (x EID,y EID) EID { 
    return EID{C__INT,IVAL(/*(sm for min @ integer= integer)*/ F_min_integer(INT(x),INT(y) ))}} 
  
/* {1} OPT.The go function for: max(x:integer,y:integer) [] */
func F_max_integer (x int,y int) int { 
    // use function body compiling 
if (x <= y) /* body If:2 */{ 
      return  y
      } else {
      return  x
      /* body If-2 */} 
    } 
  
// The EID go function for: max @ integer (throw: false) 
func E_max_integer (x EID,y EID) EID { 
    return EID{C__INT,IVAL(/*(sm for max @ integer= integer)*/ F_max_integer(INT(x),INT(y) ))}} 
  
/* {1} OPT.The go function for: min(x:float,y:float) [] */
func F_min_float (x float64,y float64) float64 { 
    // use function body compiling 
if (x <= y) /* body If:2 */{ 
      return  x
      } else {
      return  y
      /* body If-2 */} 
    } 
  
// The EID go function for: min @ float (throw: false) 
func E_min_float (x EID,y EID) EID { 
    return EID{C__FLOAT,FVAL(/*(sm for min @ float= float)*/ F_min_float(FLOAT(x),FLOAT(y) ))}} 
  
/* {1} OPT.The go function for: max(x:float,y:float) [] */
func F_max_float (x float64,y float64) float64 { 
    // use function body compiling 
if (x <= y) /* body If:2 */{ 
      return  y
      } else {
      return  x
      /* body If-2 */} 
    } 
  
// The EID go function for: max @ float (throw: false) 
func E_max_float (x EID,y EID) EID { 
    return EID{C__FLOAT,FVAL(/*(sm for max @ float= float)*/ F_max_float(FLOAT(x),FLOAT(y) ))}} 
  
/* {1} OPT.The go function for: min(x:any,y:any) [] */
func F_min_any (x *ClaireAny ,y *ClaireAny ) *ClaireAny  { 
    // use function body compiling 
if (ToBoolean(OBJ(Core.F_CALL(ToProperty(C__inf_equal.Id()),ARGS(x.ToEID(),y.ToEID())))) == CTRUE) /* body If:2 */{ 
      return  x
      } else {
      return  y
      /* body If-2 */} 
    } 
  
// The EID go function for: min @ any (throw: false) 
func E_min_any (x EID,y EID) EID { 
    return /*(sm for min @ any= any)*/ F_min_any(ANY(x),ANY(y) ).ToEID()} 
  
/* {1} OPT.The go function for: max(x:any,y:any) [] */
func F_max_any (x *ClaireAny ,y *ClaireAny ) *ClaireAny  { 
    // use function body compiling 
if (ToBoolean(OBJ(Core.F_CALL(ToProperty(C__inf_equal.Id()),ARGS(x.ToEID(),y.ToEID())))) == CTRUE) /* body If:2 */{ 
      return  y
      } else {
      return  x
      /* body If-2 */} 
    } 
  
// The EID go function for: max @ any (throw: false) 
func E_max_any (x EID,y EID) EID { 
    return /*(sm for max @ any= any)*/ F_max_any(ANY(x),ANY(y) ).ToEID()} 
  
// check if the value if known?
/* {1} OPT.The go function for: known?(a:table,x:any) [] */
func F_known_ask_table (a *ClaireTable ,x *ClaireAny ) *ClaireBoolean  { 
    // use function body compiling 
return  Core.F__I_equal_any(Core.F_get_table(a,x),CNULL)
    } 
  
// The EID go function for: known? @ table (throw: false) 
func E_known_ask_table (a EID,x EID) EID { 
    return EID{/*(sm for known? @ table= boolean)*/ F_known_ask_table(ToTable(OBJ(a)),ANY(x) ).Id(),0}} 
  
/* {1} OPT.The go function for: unknown?(a:table,x:any) [] */
func F_unknown_ask_table (a *ClaireTable ,x *ClaireAny ) *ClaireBoolean  { 
    // use function body compiling 
return  Equal(Core.F_get_table(a,x),CNULL)
    } 
  
// The EID go function for: unknown? @ table (throw: false) 
func E_unknown_ask_table (a EID,x EID) EID { 
    return EID{/*(sm for unknown? @ table= boolean)*/ F_unknown_ask_table(ToTable(OBJ(a)),ANY(x) ).Id(),0}} 
  
/* {1} OPT.The go function for: float!(self:string) [] */
func F_float_I_string (self *ClaireString ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var x *ClaireAny  
      /* noccur = 4 */
      var x_try04493 EID 
      x_try04493 = F_read_string(self)
      /* ERROR PROTECTION INSERTED (x-Result) */
      if ErrorIn(x_try04493) {Result = x_try04493
      } else {
      x = ANY(x_try04493)
      if (C_float.Id() == x.Isa.Id()) /* If:3 */{ 
        /* Let:4 */{ 
          var g0446 float64  = ToFloat(x).Value
          /* noccur = 1 */
          Result = EID{C__FLOAT,FVAL(g0446)}
          /* Let-4 */} 
        /* If!3 */}  else if (C_integer.Id() == x.Isa.Id()) /* If:3 */{ 
        /* Let:4 */{ 
          var g0447 int  = ToInteger(x).Value
          /* noccur = 1 */
          Result = EID{C__FLOAT,FVAL(F_to_float(g0447))}
          /* Let-4 */} 
        } else {
        Result = ToException(Core.C_general_error.Make(MakeString("[??] ~A is not a float").Id(),MakeConstantList((self).Id()).Id())).Close()
        /* If-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: float! @ string (throw: true) 
func E_float_I_string (self EID) EID { 
    return /*(sm for float! @ string= EID)*/ F_float_I_string(ToString(OBJ(self)) )} 
  
// v3.00.46 a new macro
/* {1} OPT.The go function for: >=(self:any,x:any) [] */
func F__sup_equal_any (self *ClaireAny ,x *ClaireAny ) *ClaireBoolean  { 
    // use function body compiling 
return  ToBoolean(OBJ(Core.F_CALL(ToProperty(C__inf_equal.Id()),ARGS(x.ToEID(),self.ToEID()))))
    } 
  
// The EID go function for: >= @ any (throw: false) 
func E__sup_equal_any (self EID,x EID) EID { 
    return EID{/*(sm for >= @ any= boolean)*/ F__sup_equal_any(ANY(self),ANY(x) ).Id(),0}} 
  
// v3.3.42 add macros to use float & integers easily
/* {1} OPT.The go function for: +(x:integer,y:float) [] */
func F__plus_integer2 (x int,y float64) float64 { 
    // use function body compiling 
return  (F_to_float(x)+y)
    } 
  
// The EID go function for: + @ list<type_expression>(integer, float) (throw: false) 
func E__plus_integer2 (x EID,y EID) EID { 
    return EID{C__FLOAT,FVAL(/*(sm for + @ list<type_expression>(integer, float)= float)*/ F__plus_integer2(INT(x),FLOAT(y) ))}} 
  
/* {1} OPT.The go function for: *(x:integer,y:float) [] */
func F__star_integer2 (x int,y float64) float64 { 
    // use function body compiling 
return  (F_to_float(x)*y)
    } 
  
// The EID go function for: * @ list<type_expression>(integer, float) (throw: false) 
func E__star_integer2 (x EID,y EID) EID { 
    return EID{C__FLOAT,FVAL(/*(sm for * @ list<type_expression>(integer, float)= float)*/ F__star_integer2(INT(x),FLOAT(y) ))}} 
  
/* {1} OPT.The go function for: /(x:integer,y:float) [] */
func F__7_integer2 (x int,y float64) float64 { 
    // use function body compiling 
return  (F_to_float(x)/y)
    } 
  
// The EID go function for: / @ list<type_expression>(integer, float) (throw: false) 
func E__7_integer2 (x EID,y EID) EID { 
    return EID{C__FLOAT,FVAL(/*(sm for / @ list<type_expression>(integer, float)= float)*/ F__7_integer2(INT(x),FLOAT(y) ))}} 
  
/* {1} OPT.The go function for: -(x:integer,y:float) [] */
func F__dash_integer3 (x int,y float64) float64 { 
    // use function body compiling 
return  (F_to_float(x)-y)
    } 
  
// The EID go function for: - @ list<type_expression>(integer, float) (throw: false) 
func E__dash_integer3 (x EID,y EID) EID { 
    return EID{C__FLOAT,FVAL(/*(sm for - @ list<type_expression>(integer, float)= float)*/ F__dash_integer3(INT(x),FLOAT(y) ))}} 
  
/* {1} OPT.The go function for: +(x:float,y:integer) [] */
func F__plus_float2 (x float64,y int) float64 { 
    // use function body compiling 
return  (x+F_to_float(y))
    } 
  
// The EID go function for: + @ list<type_expression>(float, integer) (throw: false) 
func E__plus_float2 (x EID,y EID) EID { 
    return EID{C__FLOAT,FVAL(/*(sm for + @ list<type_expression>(float, integer)= float)*/ F__plus_float2(FLOAT(x),INT(y) ))}} 
  
/* {1} OPT.The go function for: *(x:float,y:integer) [] */
func F__star_float2 (x float64,y int) float64 { 
    // use function body compiling 
return  (x*F_to_float(y))
    } 
  
// The EID go function for: * @ list<type_expression>(float, integer) (throw: false) 
func E__star_float2 (x EID,y EID) EID { 
    return EID{C__FLOAT,FVAL(/*(sm for * @ list<type_expression>(float, integer)= float)*/ F__star_float2(FLOAT(x),INT(y) ))}} 
  
/* {1} OPT.The go function for: /(x:float,y:integer) [] */
func F__7_float2 (x float64,y int) float64 { 
    // use function body compiling 
return  (x/F_to_float(y))
    } 
  
// The EID go function for: / @ list<type_expression>(float, integer) (throw: false) 
func E__7_float2 (x EID,y EID) EID { 
    return EID{C__FLOAT,FVAL(/*(sm for / @ list<type_expression>(float, integer)= float)*/ F__7_float2(FLOAT(x),INT(y) ))}} 
  
/* {1} OPT.The go function for: -(x:float,y:integer) [] */
func F__dash_float3 (x float64,y int) float64 { 
    // use function body compiling 
return  (x-F_to_float(y))
    } 
  
// The EID go function for: - @ list<type_expression>(float, integer) (throw: false) 
func E__dash_float3 (x EID,y EID) EID { 
    return EID{C__FLOAT,FVAL(/*(sm for - @ list<type_expression>(float, integer)= float)*/ F__dash_float3(FLOAT(x),INT(y) ))}} 
  
// v3.4 a useful macro
/* {1} OPT.The go function for: sqr(x:integer) [] */
func F_sqr_integer (x int) int { 
    // use function body compiling 
return  (x*x)
    } 
  
// The EID go function for: sqr @ integer (throw: false) 
func E_sqr_integer (x EID) EID { 
    return EID{C__INT,IVAL(/*(sm for sqr @ integer= integer)*/ F_sqr_integer(INT(x) ))}} 
  
/* {1} OPT.The go function for: sqr(x:float) [] */
func F_sqr_float (x float64) float64 { 
    // use function body compiling 
return  (x*x)
    } 
  
// The EID go function for: sqr @ float (throw: false) 
func E_sqr_float (x EID) EID { 
    return EID{C__FLOAT,FVAL(/*(sm for sqr @ float= float)*/ F_sqr_float(FLOAT(x) ))}} 
  