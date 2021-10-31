/***** CLAIRE Compilation of file /Users/ycaseau/Dropbox/src/clairev4.0/src/compile/goexp.cl 
         [version 4.0.02 / safety 5] Saturday 10-30-2021 *****/

package Generate
import (_ "fmt"
	. "Kernel"
	"Core"
	"Language"
	"Reader"
	"Optimize"
)

//-------- dumb function to prevent import errors --------
func import_g0223() { 
    _ = Core.It
    _ = Language.It
    _ = Reader.It
    _ = Optimize.It
    } 
  
  
//+-------------------------------------------------------------+
//| CLAIRE                                                      |
//| goexp.cl                                                    |
//| Copyright (C) 2020 - 2021 Yves Caseau. All Rights Reserved  |
//| cf. copyright info in file object.cl: about()               |
//+-------------------------------------------------------------+
// ---------------------------------------------------------------------
// Compiling is based upon three methods:
//  - g_func? tests if the CLAIRE form can be represented by a C/ expression.
//    In this case,
//  - g_expression transforms it into an equivalent go expression.
//    otherwise,
//  - gstatement takes also a variable as an argument, and transforms a CLAIRE
//    expression into a C statement that assigns the value of the expression
//    into the variable;
//
// A special case occurs when the expression represent a boolean value and is
// functional, we can use bool_exp that returns a C boolean
// ---------------------------------------------------------------------
// *********************************************************************
// *  Contents                                                         *
// *  Part 1: g_func & expression for objects                          *
// *  Part 2: expression for messages                                  *
// *  Part 3: the inline coding of function calls                      *
// *  Part 4: expression for structures                                *
// *  Part 5: boolean optimization                                     *
// *********************************************************************
// g_expression(x:any,s:class) produces a go expression based on expected go type
//     s = EID                            => produce an EID
//     s = any, object, c                 => produces a *ClaireAny  representation (default case)
//     s = integer, char, float, string   => produced a native representation
//**********************************************************************
//*          Part 1: g_func & expression for objects                   *
//**********************************************************************
// this methods tells if a CLAIRE instruction can be compiled as an expression,as opposed to a statement.
// CHANGE in CLAIRE 4 : everything that may throw an exception needs a statement (because of go limitation)
// HOWEVER : if a call produces the possible error, it should simply be compiled in EID mode
/* {1} OPT.The go function for: g_func(self:any) [] */
func F_Generate_g_func_any (self *ClaireAny ) EID { 
    var Result EID 
    if (self.Isa.IsIn(C_bag) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0224 *ClaireBag   = ToBag(self)
        /* noccur = 1 */
        /* Let:4 */{ 
          var g0238UU *ClaireAny  
          /* noccur = 1 */
          var g0238UU_try02395 EID 
          /* For:5 */{ 
            var x *ClaireAny  
            _ = x
            g0238UU_try02395= EID{CFALSE.Id(),0}
            var x_support *ClaireList  
            var x_support_try02406 EID 
            x_support_try02406 = Core.F_enumerate_any(g0224.Id())
            /* ERROR PROTECTION INSERTED (x_support-g0238UU_try02395) */
            if ErrorIn(x_support_try02406) {g0238UU_try02395 = x_support_try02406
            } else {
            x_support = ToList(OBJ(x_support_try02406))
            x_len := x_support.Length()
            for i_it := 0; i_it < x_len; i_it++ { 
              x = x_support.At(i_it)
              var void_try7 EID 
              _ = void_try7
              var g0241I *ClaireBoolean  
              var g0241I_try02427 EID 
              /* Let:7 */{ 
                var g0243UU *ClaireBoolean  
                /* noccur = 1 */
                var g0243UU_try02448 EID 
                g0243UU_try02448 = F_Generate_g_func_any(x)
                /* ERROR PROTECTION INSERTED (g0243UU-g0241I_try02427) */
                if ErrorIn(g0243UU_try02448) {g0241I_try02427 = g0243UU_try02448
                } else {
                g0243UU = ToBoolean(OBJ(g0243UU_try02448))
                g0241I_try02427 = EID{g0243UU.Not.Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (g0241I-void_try7) */
              if ErrorIn(g0241I_try02427) {void_try7 = g0241I_try02427
              } else {
              g0241I = ToBoolean(OBJ(g0241I_try02427))
              if (g0241I == CTRUE) /* If:7 */{ 
                 /*v = g0238UU_try02395, s =EID*/
g0238UU_try02395 = EID{CTRUE.Id(),0}
                break
                } else {
                void_try7 = EID{CFALSE.Id(),0}
                /* If-7 */} 
              }
              /* ERROR PROTECTION INSERTED (void_try7-g0238UU_try02395) */
              if ErrorIn(void_try7) {g0238UU_try02395 = void_try7
              g0238UU_try02395 = void_try7
              break
              } else {
              }}
              /* loop-6 */} 
            /* For-5 */} 
          /* ERROR PROTECTION INSERTED (g0238UU-Result) */
          if ErrorIn(g0238UU_try02395) {Result = g0238UU_try02395
          } else {
          g0238UU = ANY(g0238UU_try02395)
          Result = EID{Core.F_not_any(g0238UU).Id(),0}
          }
          /* Let-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(Language.C_Construct) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0225 *Language.Construct   = Language.To_Construct(self)
        /* noccur = 5 */
        if (((g0225.Isa.IsIn(Language.C_Set) == CTRUE) || 
              (g0225.Isa.IsIn(Language.C_List) == CTRUE)) || 
            (g0225.Isa.IsIn(Language.C_Tuple) == CTRUE)) /* If:4 */{ 
          /* and:5 */{ 
            var v_and5 *ClaireBoolean  
            
            v_and5 = Core.F__inf_integer(g0225.Args.Length(),15)
            if (v_and5 == CFALSE) {Result = EID{CFALSE.Id(),0}
            } else /* arg:6 */{ 
              var v_and5_try02457 EID 
              /* Let:7 */{ 
                var g0246UU *ClaireAny  
                /* noccur = 1 */
                var g0246UU_try02478 EID 
                /* For:8 */{ 
                  var x *ClaireAny  
                  _ = x
                  g0246UU_try02478= EID{CFALSE.Id(),0}
                  var x_support *ClaireList  
                  x_support = g0225.Args
                  x_len := x_support.Length()
                  for i_it := 0; i_it < x_len; i_it++ { 
                    x = x_support.At(i_it)
                    var void_try10 EID 
                    _ = void_try10
                    var g0248I *ClaireBoolean  
                    var g0248I_try024910 EID 
                    /* Let:10 */{ 
                      var g0250UU *ClaireBoolean  
                      /* noccur = 1 */
                      var g0250UU_try025111 EID 
                      g0250UU_try025111 = F_Generate_g_func_any(x)
                      /* ERROR PROTECTION INSERTED (g0250UU-g0248I_try024910) */
                      if ErrorIn(g0250UU_try025111) {g0248I_try024910 = g0250UU_try025111
                      } else {
                      g0250UU = ToBoolean(OBJ(g0250UU_try025111))
                      g0248I_try024910 = EID{g0250UU.Not.Id(),0}
                      }
                      /* Let-10 */} 
                    /* ERROR PROTECTION INSERTED (g0248I-void_try10) */
                    if ErrorIn(g0248I_try024910) {void_try10 = g0248I_try024910
                    } else {
                    g0248I = ToBoolean(OBJ(g0248I_try024910))
                    if (g0248I == CTRUE) /* If:10 */{ 
                       /*v = g0246UU_try02478, s =EID*/
g0246UU_try02478 = EID{CTRUE.Id(),0}
                      break
                      } else {
                      void_try10 = EID{CFALSE.Id(),0}
                      /* If-10 */} 
                    }
                    /* ERROR PROTECTION INSERTED (void_try10-g0246UU_try02478) */
                    if ErrorIn(void_try10) {g0246UU_try02478 = void_try10
                    g0246UU_try02478 = void_try10
                    break
                    } else {
                    }
                    /* loop-9 */} 
                  /* For-8 */} 
                /* ERROR PROTECTION INSERTED (g0246UU-v_and5_try02457) */
                if ErrorIn(g0246UU_try02478) {v_and5_try02457 = g0246UU_try02478
                } else {
                g0246UU = ANY(g0246UU_try02478)
                v_and5_try02457 = EID{Core.F_not_any(g0246UU).Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (v_and5-Result) */
              if ErrorIn(v_and5_try02457) {Result = v_and5_try02457
              } else {
              v_and5 = ToBoolean(OBJ(v_and5_try02457))
              if (v_and5 == CFALSE) {Result = EID{CFALSE.Id(),0}
              } else /* arg:7 */{ 
                Result = EID{CTRUE.Id(),0}/* arg-7 */} 
              /* arg-6 */} 
            }
            /* and-5 */} 
          } else {
          Result = EID{CFALSE.Id(),0}
          /* If-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(Language.C_If) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0226 *Language.If   = Language.To_If(self)
        /* noccur = 3 */
        /* and:4 */{ 
          var v_and4 *ClaireBoolean  
          
          var v_and4_try02525 EID 
          v_and4_try02525 = F_Generate_g_func_any(g0226.Test)
          /* ERROR PROTECTION INSERTED (v_and4-Result) */
          if ErrorIn(v_and4_try02525) {Result = v_and4_try02525
          } else {
          v_and4 = ToBoolean(OBJ(v_and4_try02525))
          if (v_and4 == CFALSE) {Result = EID{CFALSE.Id(),0}
          } else /* arg:5 */{ 
            v_and4 = F_Generate_constant_ask_any(g0226.Arg)
            if (v_and4 == CFALSE) {Result = EID{CFALSE.Id(),0}
            } else /* arg:6 */{ 
              v_and4 = F_Generate_constant_ask_any(g0226.Other)
              if (v_and4 == CFALSE) {Result = EID{CFALSE.Id(),0}
              } else /* arg:7 */{ 
                Result = EID{CTRUE.Id(),0}/* arg-7 */} 
              /* arg-6 */} 
            /* arg-5 */} 
          }
          /* and-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(Language.C_And) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0227 *Language.And   = Language.To_And(self)
        /* noccur = 1 */
        Result = F_Generate_g_func_any(g0227.Args.Id())
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(Language.C_Or) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0228 *Language.Or   = Language.To_Or(self)
        /* noccur = 1 */
        Result = F_Generate_g_func_any(g0228.Args.Id())
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(Language.C_Call) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0229 *Language.Call   = Language.To_Call(self)
        /* noccur = 3 */
        /* and:4 */{ 
          var v_and4 *ClaireBoolean  
          
          var v_and4_try02535 EID 
          v_and4_try02535 = F_Generate_g_func_any(g0229.Args.Id())
          /* ERROR PROTECTION INSERTED (v_and4-Result) */
          if ErrorIn(v_and4_try02535) {Result = v_and4_try02535
          } else {
          v_and4 = ToBoolean(OBJ(v_and4_try02535))
          if (v_and4 == CFALSE) {Result = EID{CFALSE.Id(),0}
          } else /* arg:5 */{ 
            v_and4 = Core.F__I_equal_any(g0229.Selector.Id(),Optimize.C_Compile_object_I.Id())
            if (v_and4 == CFALSE) {Result = EID{CFALSE.Id(),0}
            } else /* arg:6 */{ 
              var v_and4_try02547 EID 
              /* Let:7 */{ 
                var g0255UU *ClaireAny  
                /* noccur = 1 */
                var g0255UU_try02568 EID 
                /* For:8 */{ 
                  var x *ClaireAny  
                  _ = x
                  g0255UU_try02568= EID{CFALSE.Id(),0}
                  var x_support *ClaireList  
                  x_support = g0229.Args
                  x_len := x_support.Length()
                  for i_it := 0; i_it < x_len; i_it++ { 
                    x = x_support.At(i_it)
                    var void_try10 EID 
                    _ = void_try10
                    var g0257I *ClaireBoolean  
                    var g0257I_try025810 EID 
                    /* Let:10 */{ 
                      var g0259UU *ClaireBoolean  
                      /* noccur = 1 */
                      var g0259UU_try026011 EID 
                      /* Let:11 */{ 
                        var g0261UU *ClaireBoolean  
                        /* noccur = 1 */
                        var g0261UU_try026212 EID 
                        g0261UU_try026212 = Optimize.F_Compile_g_throw_any(x)
                        /* ERROR PROTECTION INSERTED (g0261UU-g0259UU_try026011) */
                        if ErrorIn(g0261UU_try026212) {g0259UU_try026011 = g0261UU_try026212
                        } else {
                        g0261UU = ToBoolean(OBJ(g0261UU_try026212))
                        g0259UU_try026011 = EID{g0261UU.Not.Id(),0}
                        }
                        /* Let-11 */} 
                      /* ERROR PROTECTION INSERTED (g0259UU-g0257I_try025810) */
                      if ErrorIn(g0259UU_try026011) {g0257I_try025810 = g0259UU_try026011
                      } else {
                      g0259UU = ToBoolean(OBJ(g0259UU_try026011))
                      g0257I_try025810 = EID{g0259UU.Not.Id(),0}
                      }
                      /* Let-10 */} 
                    /* ERROR PROTECTION INSERTED (g0257I-void_try10) */
                    if ErrorIn(g0257I_try025810) {void_try10 = g0257I_try025810
                    } else {
                    g0257I = ToBoolean(OBJ(g0257I_try025810))
                    if (g0257I == CTRUE) /* If:10 */{ 
                       /*v = g0255UU_try02568, s =EID*/
g0255UU_try02568 = EID{CTRUE.Id(),0}
                      break
                      } else {
                      void_try10 = EID{CFALSE.Id(),0}
                      /* If-10 */} 
                    }
                    /* ERROR PROTECTION INSERTED (void_try10-g0255UU_try02568) */
                    if ErrorIn(void_try10) {g0255UU_try02568 = void_try10
                    g0255UU_try02568 = void_try10
                    break
                    } else {
                    }
                    /* loop-9 */} 
                  /* For-8 */} 
                /* ERROR PROTECTION INSERTED (g0255UU-v_and4_try02547) */
                if ErrorIn(g0255UU_try02568) {v_and4_try02547 = g0255UU_try02568
                } else {
                g0255UU = ANY(g0255UU_try02568)
                v_and4_try02547 = EID{Core.F_not_any(g0255UU).Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (v_and4-Result) */
              if ErrorIn(v_and4_try02547) {Result = v_and4_try02547
              } else {
              v_and4 = ToBoolean(OBJ(v_and4_try02547))
              if (v_and4 == CFALSE) {Result = EID{CFALSE.Id(),0}
              } else /* arg:7 */{ 
                Result = EID{CTRUE.Id(),0}/* arg-7 */} 
              /* arg-6 */} 
            /* arg-5 */} 
          }}
          /* and-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(Language.C_Super) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0230 *Language.Super   = Language.To_Super(self)
        /* noccur = 1 */
        Result = F_Generate_g_func_any(g0230.Args.Id())
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(Language.C_Call_method) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0231 *Language.CallMethod   = Language.To_CallMethod(self)
        /* noccur = 3 */
        /* and:4 */{ 
          var v_and4 *ClaireBoolean  
          
          var v_and4_try02635 EID 
          v_and4_try02635 = F_Generate_g_func_any(g0231.Args.Id())
          /* ERROR PROTECTION INSERTED (v_and4-Result) */
          if ErrorIn(v_and4_try02635) {Result = v_and4_try02635
          } else {
          v_and4 = ToBoolean(OBJ(v_and4_try02635))
          if (v_and4 == CFALSE) {Result = EID{CFALSE.Id(),0}
          } else /* arg:5 */{ 
            var v_and4_try02646 EID 
            /* or:6 */{ 
              var v_or6 *ClaireBoolean  
              
              v_or6 = Equal(g0231.Arg.Id(),Optimize.C_Compile_m_unsafe.Value)
              if (v_or6 == CTRUE) {v_and4_try02646 = EID{CTRUE.Id(),0}
              } else /* or:7 */{ 
                var v_or6_try02658 EID 
                /* Let:8 */{ 
                  var g0266UU *ClaireAny  
                  /* noccur = 1 */
                  var g0266UU_try02679 EID 
                  /* For:9 */{ 
                    var x *ClaireAny  
                    _ = x
                    g0266UU_try02679= EID{CFALSE.Id(),0}
                    var x_support *ClaireList  
                    x_support = g0231.Args
                    x_len := x_support.Length()
                    for i_it := 0; i_it < x_len; i_it++ { 
                      x = x_support.At(i_it)
                      var void_try11 EID 
                      _ = void_try11
                      var g0268I *ClaireBoolean  
                      var g0268I_try026911 EID 
                      /* Let:11 */{ 
                        var g0270UU *ClaireBoolean  
                        /* noccur = 1 */
                        var g0270UU_try027112 EID 
                        /* Let:12 */{ 
                          var g0272UU *ClaireBoolean  
                          /* noccur = 1 */
                          var g0272UU_try027313 EID 
                          g0272UU_try027313 = Optimize.F_Compile_g_throw_any(x)
                          /* ERROR PROTECTION INSERTED (g0272UU-g0270UU_try027112) */
                          if ErrorIn(g0272UU_try027313) {g0270UU_try027112 = g0272UU_try027313
                          } else {
                          g0272UU = ToBoolean(OBJ(g0272UU_try027313))
                          g0270UU_try027112 = EID{g0272UU.Not.Id(),0}
                          }
                          /* Let-12 */} 
                        /* ERROR PROTECTION INSERTED (g0270UU-g0268I_try026911) */
                        if ErrorIn(g0270UU_try027112) {g0268I_try026911 = g0270UU_try027112
                        } else {
                        g0270UU = ToBoolean(OBJ(g0270UU_try027112))
                        g0268I_try026911 = EID{g0270UU.Not.Id(),0}
                        }
                        /* Let-11 */} 
                      /* ERROR PROTECTION INSERTED (g0268I-void_try11) */
                      if ErrorIn(g0268I_try026911) {void_try11 = g0268I_try026911
                      } else {
                      g0268I = ToBoolean(OBJ(g0268I_try026911))
                      if (g0268I == CTRUE) /* If:11 */{ 
                         /*v = g0266UU_try02679, s =EID*/
g0266UU_try02679 = EID{CTRUE.Id(),0}
                        break
                        } else {
                        void_try11 = EID{CFALSE.Id(),0}
                        /* If-11 */} 
                      }
                      /* ERROR PROTECTION INSERTED (void_try11-g0266UU_try02679) */
                      if ErrorIn(void_try11) {g0266UU_try02679 = void_try11
                      g0266UU_try02679 = void_try11
                      break
                      } else {
                      }
                      /* loop-10 */} 
                    /* For-9 */} 
                  /* ERROR PROTECTION INSERTED (g0266UU-v_or6_try02658) */
                  if ErrorIn(g0266UU_try02679) {v_or6_try02658 = g0266UU_try02679
                  } else {
                  g0266UU = ANY(g0266UU_try02679)
                  v_or6_try02658 = EID{Core.F_not_any(g0266UU).Id(),0}
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (v_or6-v_and4_try02646) */
                if ErrorIn(v_or6_try02658) {v_and4_try02646 = v_or6_try02658
                } else {
                v_or6 = ToBoolean(OBJ(v_or6_try02658))
                if (v_or6 == CTRUE) {v_and4_try02646 = EID{CTRUE.Id(),0}
                } else /* or:8 */{ 
                  v_and4_try02646 = EID{CFALSE.Id(),0}/* org-8 */} 
                /* org-7 */} 
              }
              /* or-6 */} 
            /* ERROR PROTECTION INSERTED (v_and4-Result) */
            if ErrorIn(v_and4_try02646) {Result = v_and4_try02646
            } else {
            v_and4 = ToBoolean(OBJ(v_and4_try02646))
            if (v_and4 == CFALSE) {Result = EID{CFALSE.Id(),0}
            } else /* arg:6 */{ 
              Result = EID{CTRUE.Id(),0}/* arg-6 */} 
            /* arg-5 */} 
          }}
          /* and-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(Language.C_Call_slot) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0232 *Language.CallSlot   = Language.To_CallSlot(self)
        /* noccur = 1 */
        Result = F_Generate_g_func_any(g0232.Arg)
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(Language.C_Call_table) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0233 *Language.CallTable   = Language.To_CallTable(self)
        /* noccur = 1 */
        Result = F_Generate_g_func_any(g0233.Arg)
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(Language.C_Call_array) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0234 *Language.CallArray   = Language.To_CallArray(self)
        /* noccur = 1 */
        Result = F_Generate_g_func_any(g0234.Arg)
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(Language.C_Cast) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0235 *Language.Cast   = Language.To_Cast(self)
        /* noccur = 1 */
        Result = F_Generate_g_func_any(g0235.Arg)
        /* Let-3 */} 
      /* If!2 */}  else if (self.Isa.IsIn(Optimize.C_Compile_C_cast) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0236 *Optimize.CompileCCast   = Optimize.To_CompileCCast(self)
        /* noccur = 1 */
        Result = F_Generate_g_func_any(g0236.Arg)
        /* Let-3 */} 
      } else {
      Result = EID{MakeBoolean((self.Isa.IsIn(C_thing) == CTRUE) || 
      (C_integer.Id() == self.Isa.Id()) || 
      (C_string.Id() == self.Isa.Id()) || 
      (C_char.Id() == self.Isa.Id()) || 
      (C_float.Id() == self.Isa.Id()) || 
      (self.Isa.IsIn(C_Variable) == CTRUE) || 
      (self.Isa.IsIn(Core.C_global_variable) == CTRUE) || 
      (C_function.Id() == self.Isa.Id()) || 
      (self.Isa.IsIn(C_symbol) == CTRUE) || 
      (self == CNULL) || 
      (C_boolean.Id() == self.Isa.Id()) || 
      (C_class.Id() == self.Isa.Id()) || 
      (self.Isa.IsIn(C_environment) == CTRUE)).Id(),0}
      /* If-2 */} 
    return Result} 
  
// The EID go function for: g_func @ any (throw: true) 
func E_Generate_g_func_any (self EID) EID { 
    return /*(sm for g_func @ any= EID)*/ F_Generate_g_func_any(ANY(self) )} 
  
// manages unknown + catch-all 
/* {1} OPT.The go function for: g_expression(self:any,s:class) [] */
func F_Generate_g_expression_any (self *ClaireAny ,s *ClaireClass ) EID { 
    var Result EID 
    if (self != CNULL) /* If:2 */{ 
      Result = ToException(Core.C_general_error.Make(MakeString("/! design error: g_expression(~S: ~S) unknown").Id(),MakeConstantList(self,self.Isa.Id()).Id())).Close()
      /* If!2 */}  else if (s.Id() == Optimize.C_EID.Id()) /* If:2 */{ 
      Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).ToEid(self,C_object)
      /* If!2 */}  else if (s.Id() == C_any.Id()) /* If:2 */{ 
      PRINC("CNULL")
      Result = EVOID
      } else {
      F_Generate_object_prefix_class(C_any,s)
      PRINC("CNULL")
      F_Generate_object_post_class(C_any,s)
      PRINC("")
      Result = EVOID
      /* If-2 */} 
    return Result} 
  
// The EID go function for: g_expression @ any (throw: true) 
func E_Generate_g_expression_any (self EID,s EID) EID { 
    return /*(sm for g_expression @ any= EID)*/ F_Generate_g_expression_any(ANY(self),ToClass(OBJ(s)) )} 
  
// Things are represented by global variables in the associated go package
/* {1} OPT.The go function for: g_expression(self:thing,s:class) [] */
func F_Generate_g_expression_thing (self *ClaireThing ,s *ClaireClass ) EID { 
    var Result EID 
    if (s.Id() == Optimize.C_EID.Id()) /* If:2 */{ 
      Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).ToEid(self.Id(),C_object)
      } else {
      F_Generate_object_prefix_class(self.Id().Isa,s)
      F_Generate_thing_ident_thing(self)
      F_Generate_object_post_class(self.Id().Isa,s)
      PRINC("")
      Result = EVOID
      /* If-2 */} 
    return Result} 
  
// The EID go function for: g_expression @ thing (throw: true) 
func E_Generate_g_expression_thing (self EID,s EID) EID { 
    return /*(sm for g_expression @ thing= EID)*/ F_Generate_g_expression_thing(ToThing(OBJ(self)),ToClass(OBJ(s)) )} 
  
// note that there are two kinds of modules
//    - packages (when m.made_of != nil)  -> defined in their first members (iClaire in Language)
//    - node modules (abstractions) => need to be attached to packages
/* {1} OPT.The go function for: g_expression(self:module,s:class) [] */
func F_Generate_g_expression_module (self *ClaireModule ,s *ClaireClass ) EID { 
    var Result EID 
    if (s.Id() == Optimize.C_EID.Id()) /* If:2 */{ 
      Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).ToEid(self.Id(),C_object)
      } else {
      F_Generate_object_prefix_class(self.Id().Isa,s)
      if (self.Id() == ToGenerateGoProducer(Optimize.C_PRODUCER.Value).Current.Id()) /* If:3 */{ 
        PRINC("It")
        /* If!3 */}  else if (self.Id() == C_Kernel.Id()) /* If:3 */{ 
        PRINC("C_Kernel")
        /* If!3 */}  else if (Equal(self.MadeOf.Id(),CNIL.Id()) == CTRUE) /* If:3 */{ 
        /* Let:4 */{ 
          var m *ClaireModule   = F_Generate_get_made_module(self)
          /* noccur = 3 */
          if ((m.Id() != C_Kernel.Id()) && 
              (m.Id() != ToGenerateGoProducer(Optimize.C_PRODUCER.Value).Current.Id())) /* If:5 */{ 
            F_Generate_cap_short_symbol(m.Name)
            PRINC(".")
            /* If-5 */} 
          F_Generate_go_var_symbol(self.Name)
          /* Let-4 */} 
        } else {
        F_Generate_cap_ident_symbol(self.Name)
        PRINC(".It")
        /* If-3 */} 
      F_Generate_object_post_class(self.Id().Isa,s)
      PRINC("")
      Result = EVOID
      /* If-2 */} 
    return Result} 
  
// The EID go function for: g_expression @ module (throw: true) 
func E_Generate_g_expression_module (self EID,s EID) EID { 
    return /*(sm for g_expression @ module= EID)*/ F_Generate_g_expression_module(ToModule(OBJ(self)),ToClass(OBJ(s)) )} 
  
// A class is similar to a thing
/* {1} OPT.The go function for: g_expression(self:class,s:class) [] */
func F_Generate_g_expression_class (self *ClaireClass ,s *ClaireClass ) EID { 
    var Result EID 
    if (s.Id() == Optimize.C_EID.Id()) /* If:2 */{ 
      Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).ToEid(self.Id(),C_object)
      } else {
      F_Generate_object_prefix_class(C_class,s)
      F_Generate_class_ident_class(self)
      F_Generate_object_post_class(C_class,s)
      PRINC("")
      Result = EVOID
      /* If-2 */} 
    return Result} 
  
// The EID go function for: g_expression @ class (throw: true) 
func E_Generate_g_expression_class (self EID,s EID) EID { 
    return /*(sm for g_expression @ class= EID)*/ F_Generate_g_expression_class(ToClass(OBJ(self)),ToClass(OBJ(s)) )} 
  
// A named object is designed by a C identifier !
/* {1} OPT.The go function for: g_expression(self:boolean,s:class) [] */
func F_Generate_g_expression_boolean (self *ClaireBoolean ,s *ClaireClass ) EID { 
    var Result EID 
    if (s.Id() == Optimize.C_EID.Id()) /* If:2 */{ 
      Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).ToEid(self.Id(),C_object)
      } else {
      F_Generate_object_prefix_class(C_boolean,s)
      F_princ_string(ToString(IfThenElse((self == CTRUE),
        MakeString("CTRUE").Id(),
        MakeString("CFALSE").Id())))
      F_Generate_object_post_class(C_boolean,s)
      PRINC("")
      Result = EVOID
      /* If-2 */} 
    return Result} 
  
// The EID go function for: g_expression @ boolean (throw: true) 
func E_Generate_g_expression_boolean (self EID,s EID) EID { 
    return /*(sm for g_expression @ boolean= EID)*/ F_Generate_g_expression_boolean(ToBoolean(OBJ(self)),ToClass(OBJ(s)) )} 
  
// Primitive types rely on the producer to generate code that uses their specific implementation
// this is done on purpose: supports the customization through another producer
/* {1} OPT.The go function for: g_expression(self:integer,s:class) [] */
func F_Generate_g_expression_integer (self int,s *ClaireClass ) EID { 
    var Result EID 
    if (s.Id() == Optimize.C_EID.Id()) /* If:2 */{ 
      Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).ToEid(MakeInteger(self).Id(),C_integer)
      /* If!2 */}  else if (s.Id() == C_integer.Id()) /* If:2 */{ 
      F_princ_integer(self)
      Result = EVOID
      } else {
      F_Generate_object_prefix_class(C_integer,s)
      Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).ToCl(MakeInteger(self).Id(),C_integer)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      F_Generate_object_post_class(C_integer,s)
      Result = EVOID
      }
      /* If-2 */} 
    return Result} 
  
// The EID go function for: g_expression @ integer (throw: true) 
func E_Generate_g_expression_integer (self EID,s EID) EID { 
    return /*(sm for g_expression @ integer= EID)*/ F_Generate_g_expression_integer(INT(self),ToClass(OBJ(s)) )} 
  
/* {1} OPT.The go function for: g_expression(self:float,s:class) [] */
func F_Generate_g_expression_float (self float64,s *ClaireClass ) EID { 
    var Result EID 
    if (s.Id() == Optimize.C_EID.Id()) /* If:2 */{ 
      Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).ToEid(MakeFloat(self).Id(),C_float)
      /* If!2 */}  else if (s.Id() == C_float.Id()) /* If:2 */{ 
      F_princ_float(self)
      Result = EVOID
      } else {
      F_Generate_object_prefix_class(C_float,s)
      Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).ToCl(MakeFloat(self).Id(),C_float)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      F_Generate_object_post_class(C_float,s)
      Result = EVOID
      }
      /* If-2 */} 
    return Result} 
  
// The EID go function for: g_expression @ float (throw: true) 
func E_Generate_g_expression_float (self EID,s EID) EID { 
    return /*(sm for g_expression @ float= EID)*/ F_Generate_g_expression_float(FLOAT(self),ToClass(OBJ(s)) )} 
  
/* {1} OPT.The go function for: g_expression(self:char,s:class) [] */
func F_Generate_g_expression_char (self rune,s *ClaireClass ) EID { 
    var Result EID 
    if (s.Id() == Optimize.C_EID.Id()) /* If:2 */{ 
      Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).ToEid(MakeChar(self).Id(),C_char)
      /* If!2 */}  else if (s.Id() == C_char.Id()) /* If:2 */{ 
      Result = Core.F_print_any(MakeChar(self).Id())
      } else {
      F_Generate_object_prefix_class(C_char,s)
      Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).ToCl(MakeChar(self).Id(),C_char)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      F_Generate_object_post_class(C_char,s)
      Result = EVOID
      }
      /* If-2 */} 
    return Result} 
  
// The EID go function for: g_expression @ char (throw: true) 
func E_Generate_g_expression_char (self EID,s EID) EID { 
    return /*(sm for g_expression @ char= EID)*/ F_Generate_g_expression_char(CHAR(self),ToClass(OBJ(s)) )} 
  
// strings are primitive objects, same as function
/* {1} OPT.The go function for: g_expression(self:string,s:class) [] */
func F_Generate_g_expression_string (self *ClaireString ,s *ClaireClass ) EID { 
    var Result EID 
    if (s.Id() == Optimize.C_EID.Id()) /* If:2 */{ 
      Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).ToEid((self).Id(),C_string)
      } else {
      F_Generate_object_prefix_class(C_string,s)
      PRINC("MakeString(")
      Result = Core.F_print_any((self).Id())
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(")")
      F_Generate_object_post_class(C_string,s)
      PRINC("")
      Result = EVOID
      }
      /* If-2 */} 
    return Result} 
  
// The EID go function for: g_expression @ string (throw: true) 
func E_Generate_g_expression_string (self EID,s EID) EID { 
    return /*(sm for g_expression @ string= EID)*/ F_Generate_g_expression_string(ToString(OBJ(self)),ToClass(OBJ(s)) )} 
  
// symboles are primitive objects, same as function
/* {1} OPT.The go function for: g_expression(self:symbol,s:class) [] */
func F_Generate_g_expression_symbol (self *ClaireSymbol ,s *ClaireClass ) EID { 
    var Result EID 
    if (s.Id() == Optimize.C_EID.Id()) /* If:2 */{ 
      Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).ToEid(self.Id(),C_object)
      } else {
      F_Generate_object_prefix_class(C_object,s)
      PRINC("MakeSymbol(")
      Result = Core.F_print_any((self.String_I()).Id())
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(",")
      Result = F_Generate_g_expression_module(self.Module_I(),C_module)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(")")
      F_Generate_object_post_class(C_object,s)
      PRINC("")
      Result = EVOID
      }}
      /* If-2 */} 
    return Result} 
  
// The EID go function for: g_expression @ symbol (throw: true) 
func E_Generate_g_expression_symbol (self EID,s EID) EID { 
    return /*(sm for g_expression @ symbol= EID)*/ F_Generate_g_expression_symbol(ToSymbol(OBJ(self)),ToClass(OBJ(s)) )} 
  
/* {1} OPT.The go function for: g_expression(self:environment,s:class) [] */
func F_Generate_g_expression_environment (self *ClaireEnvironment ,s *ClaireClass ) EID { 
    var Result EID 
    if (s.Id() == Optimize.C_EID.Id()) /* If:2 */{ 
      Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).ToEid(self.Id(),C_object)
      } else {
      F_Generate_object_prefix_class(C_environment,s)
      PRINC("ClEnv")
      F_Generate_object_post_class(C_environment,s)
      PRINC("")
      Result = EVOID
      /* If-2 */} 
    return Result} 
  
// The EID go function for: g_expression @ environment (throw: true) 
func E_Generate_g_expression_environment (self EID,s EID) EID { 
    return /*(sm for g_expression @ environment= EID)*/ F_Generate_g_expression_environment(ToEnvironment(OBJ(self)),ToClass(OBJ(s)) )} 
  
/* {1} OPT.The go function for: g_expression(self:function,s:class) [] */
func F_Generate_g_expression_function (self *ClaireFunction ,s *ClaireClass ) EID { 
    var Result EID 
    if (s.Id() == Optimize.C_EID.Id()) /* If:2 */{ 
      Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).ToEid(self.Id(),C_object)
      } else {
      F_Generate_object_prefix_class(C_function,s)
      PRINC("MakeFunction")
      F_princ_integer(F_arity_function(self))
      PRINC("(E_")
      F_c_princ_function(self)
      PRINC(",")
      Result = Core.F_print_any((F_string_I_function(self)).Id())
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(")")
      F_Generate_object_post_class(C_function,s)
      PRINC("")
      Result = EVOID
      }
      /* If-2 */} 
    return Result} 
  
// The EID go function for: g_expression @ function (throw: true) 
func E_Generate_g_expression_function (self EID,s EID) EID { 
    return /*(sm for g_expression @ function= EID)*/ F_Generate_g_expression_function(ToFunction(OBJ(self)),ToClass(OBJ(s)) )} 
  
// lexical variables are represented by C variables
// notice that we may need native to object conversion
/* {1} OPT.The go function for: g_expression(self:Variable,s:class) [] */
func F_Generate_g_expression_Variable (self *ClaireVariable ,s *ClaireClass ) EID { 
    var Result EID 
    
    /* Let:2 */{ 
      var s2 *ClaireClass   = self.Range.Class_I()
      /* noccur = 4 */
      if (s.Id() == Optimize.C_EID.Id()) /* If:3 */{ 
        if (s2.Id() == Optimize.C_EID.Id()) /* If:4 */{ 
          Result = F_iClaire_ident_go_producer1(ToGenerateGoProducer(Optimize.C_PRODUCER.Value),self)
          } else {
          Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).ToEid(self.Id(),s2)
          /* If-4 */} 
        } else {
        Result = F_Generate_cast_prefix_class(s2,s)
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        Result = F_iClaire_ident_go_producer1(ToGenerateGoProducer(Optimize.C_PRODUCER.Value),self)
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        F_Generate_cast_post_class(s2,s)
        Result = EVOID
        }}
        /* If-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: g_expression @ Variable (throw: true) 
func E_Generate_g_expression_Variable (self EID,s EID) EID { 
    return /*(sm for g_expression @ Variable= EID)*/ F_Generate_g_expression_Variable(To_Variable(OBJ(self)),ToClass(OBJ(s)) )} 
  
// global_variables are CLAIRE objects
/* {1} OPT.The go function for: g_expression(self:global_variable,s:class) [] */
func F_Generate_g_expression_global_variable (self *Core.GlobalVariable ,s *ClaireClass ) EID { 
    var Result EID 
    if (s.Id() == Optimize.C_EID.Id()) /* If:2 */{ 
      Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).ToEid(self.Id(),C_object)
      /* If!2 */}  else if ((Equal(self.Range.Id(),CEMPTY.Id()) == CTRUE) && 
        ((C_integer.Id() == self.Value.Isa.Id()) || 
            ((C_float.Id() == self.Value.Isa.Id()) || 
              (Equal(self.Value,CNIL.Id()) == CTRUE)))) /* If:2 */{ 
      Result = Core.F_CALL(C_Generate_g_expression,ARGS(self.Value.ToEID(),EID{s.Id(),0}))
      } else {
      F_Generate_object_prefix_class(C_any,s)
      ToGenerateGoProducer(Optimize.C_PRODUCER.Value).GlobalVar(self)
      F_Generate_object_post_class(C_any,s)
      Result = EVOID
      /* If-2 */} 
    return Result} 
  
// The EID go function for: g_expression @ global_variable (throw: true) 
func E_Generate_g_expression_global_variable (self EID,s EID) EID { 
    return /*(sm for g_expression @ global_variable= EID)*/ F_Generate_g_expression_global_variable(Core.ToGlobalVariable(OBJ(self)),ToClass(OBJ(s)) )} 
  
// builds a set
/* {1} OPT.The go function for: g_expression(self:Set,s:class) [] */
func F_Generate_g_expression_Set (self *Language.Set ,s *ClaireClass ) EID { 
    var Result EID 
    if (s.Id() == Optimize.C_EID.Id()) /* If:2 */{ 
      Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).ToEid(self.Id(),C_object)
      } else {
      Result = F_Generate_cast_prefix_class(C_set,s)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      /* Let:3 */{ 
        var g0274UU *ClaireType  
        /* noccur = 1 */
        if (self.Of.Id() != CNULL) /* If:4 */{ 
          g0274UU = self.Of
          } else {
          g0274UU = ToType(CEMPTY.Id())
          /* If-4 */} 
        Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).BagExpression(C_set,self.Args,g0274UU)
        /* Let-3 */} 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      F_Generate_cast_post_class(C_set,s)
      Result = EVOID
      }}
      /* If-2 */} 
    return Result} 
  
// The EID go function for: g_expression @ Set (throw: true) 
func E_Generate_g_expression_Set (self EID,s EID) EID { 
    return /*(sm for g_expression @ Set= EID)*/ F_Generate_g_expression_Set(Language.To_Set(OBJ(self)),ToClass(OBJ(s)) )} 
  
/* {1} OPT.The go function for: g_expression(self:set,s:class) [] */
func F_Generate_g_expression_set (self *ClaireSet ,s *ClaireClass ) EID { 
    var Result EID 
    if (s.Id() == Optimize.C_EID.Id()) /* If:2 */{ 
      Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).ToEid(self.Id(),C_object)
      /* If!2 */}  else if ((self.Size() == 0) && 
        (Equal(ToList(self.Id()).Of().Id(),CEMPTY.Id()) == CTRUE)) /* If:2 */{ 
      Result = F_Generate_cast_prefix_class(C_set,s)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC("CEMPTY")
      F_Generate_cast_post_class(C_set,s)
      PRINC("")
      Result = EVOID
      }
      } else {
      Result = F_Generate_cast_prefix_class(C_set,s)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).BagExpression(C_set,self.List_I(),ToList(self.Id()).Of())
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      F_Generate_cast_post_class(C_set,s)
      Result = EVOID
      }}
      /* If-2 */} 
    return Result} 
  
// The EID go function for: g_expression @ set (throw: true) 
func E_Generate_g_expression_set (self EID,s EID) EID { 
    return /*(sm for g_expression @ set= EID)*/ F_Generate_g_expression_set(ToSet(OBJ(self)),ToClass(OBJ(s)) )} 
  
// builds a tuple
/* {1} OPT.The go function for: g_expression(self:Tuple,s:class) [] */
func F_Generate_g_expression_Tuple (self *Language.Tuple ,s *ClaireClass ) EID { 
    var Result EID 
    if (s.Id() == Optimize.C_EID.Id()) /* If:2 */{ 
      Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).ToEid(self.Id(),C_object)
      } else {
      Result = F_Generate_cast_prefix_class(C_tuple,s)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).BagExpression(C_tuple,self.Args,ToType(CEMPTY.Id()))
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      F_Generate_cast_post_class(C_tuple,s)
      Result = EVOID
      }}
      /* If-2 */} 
    return Result} 
  
// The EID go function for: g_expression @ Tuple (throw: true) 
func E_Generate_g_expression_Tuple (self EID,s EID) EID { 
    return /*(sm for g_expression @ Tuple= EID)*/ F_Generate_g_expression_Tuple(Language.To_Tuple(OBJ(self)),ToClass(OBJ(s)) )} 
  
/* {1} OPT.The go function for: g_expression(self:tuple,s:class) [] */
func F_Generate_g_expression_tuple (self *ClaireTuple ,s *ClaireClass ) EID { 
    var Result EID 
    if (s.Id() == Optimize.C_EID.Id()) /* If:2 */{ 
      Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).ToEid(self.Id(),C_object)
      } else {
      Result = F_Generate_cast_prefix_class(C_tuple,s)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).BagExpression(C_tuple,self.List_I(),ToType(CEMPTY.Id()))
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      F_Generate_cast_post_class(C_tuple,s)
      Result = EVOID
      }}
      /* If-2 */} 
    return Result} 
  
// The EID go function for: g_expression @ tuple (throw: true) 
func E_Generate_g_expression_tuple (self EID,s EID) EID { 
    return /*(sm for g_expression @ tuple= EID)*/ F_Generate_g_expression_tuple(ToTuple(OBJ(self)),ToClass(OBJ(s)) )} 
  
// builds a list
/* {1} OPT.The go function for: g_expression(self:List,s:class) [] */
func F_Generate_g_expression_List (self *Language.List ,s *ClaireClass ) EID { 
    var Result EID 
    if (s.Id() == Optimize.C_EID.Id()) /* If:2 */{ 
      Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).ToEid(self.Id(),C_object)
      } else {
      Result = F_Generate_cast_prefix_class(C_list,s)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      /* Let:3 */{ 
        var g0275UU *ClaireType  
        /* noccur = 1 */
        if (self.Of.Id() != CNULL) /* If:4 */{ 
          g0275UU = self.Of
          } else {
          g0275UU = ToType(CEMPTY.Id())
          /* If-4 */} 
        Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).BagExpression(C_list,self.Args,g0275UU)
        /* Let-3 */} 
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      F_Generate_cast_post_class(C_list,s)
      Result = EVOID
      }}
      /* If-2 */} 
    return Result} 
  
// The EID go function for: g_expression @ List (throw: true) 
func E_Generate_g_expression_List (self EID,s EID) EID { 
    return /*(sm for g_expression @ List= EID)*/ F_Generate_g_expression_List(Language.To_List(OBJ(self)),ToClass(OBJ(s)) )} 
  
/* {1} OPT.The go function for: g_expression(self:list,s:class) [] */
func F_Generate_g_expression_list (self *ClaireList ,s *ClaireClass ) EID { 
    var Result EID 
    if (s.Id() == Optimize.C_EID.Id()) /* If:2 */{ 
      Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).ToEid(self.Id(),C_object)
      /* If!2 */}  else if ((self.Length() == 0) && 
        (Equal(self.Of().Id(),CEMPTY.Id()) == CTRUE)) /* If:2 */{ 
      Result = F_Generate_cast_prefix_class(C_list,s)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC("CNIL")
      F_Generate_cast_post_class(C_list,s)
      PRINC("")
      Result = EVOID
      }
      } else {
      Result = F_Generate_cast_prefix_class(C_list,s)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).BagExpression(C_list,self,self.Of())
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      F_Generate_cast_post_class(C_list,s)
      Result = EVOID
      }}
      /* If-2 */} 
    return Result} 
  
// The EID go function for: g_expression @ list (throw: true) 
func E_Generate_g_expression_list (self EID,s EID) EID { 
    return /*(sm for g_expression @ list= EID)*/ F_Generate_g_expression_list(ToList(OBJ(self)),ToClass(OBJ(s)) )} 
  
//**********************************************************************
//*          Part 2: expression for messages                         *
//**********************************************************************
// message compiling is tricky in go : Calls produce EID but for inline, Call_method produce native forms
// calls are expected to produce an EID
/* {1} OPT.The go function for: g_expression(self:Call,s:class) [] */
func F_Generate_g_expression_Call (self *Language.Call ,s *ClaireClass ) EID { 
    var Result EID 
    Result = F_Generate_inline_exp_go_producer1(ToGenerateGoProducer(Optimize.C_PRODUCER.Value),self,s)
    return Result} 
  
// The EID go function for: g_expression @ Call (throw: true) 
func E_Generate_g_expression_Call (self EID,s EID) EID { 
    return /*(sm for g_expression @ Call= EID)*/ F_Generate_g_expression_Call(Language.To_Call(OBJ(self)),ToClass(OBJ(s)) )} 
  
// the other cases will be taken care in the optimization part
/* {1} OPT.The go function for: g_expression(self:Call_method1,s:class) [] */
func F_Generate_g_expression_Call_method1 (self *Language.CallMethod1 ,s *ClaireClass ) EID { 
    var Result EID 
    Result = Core.F_CALL(C_Generate_inline_exp,ARGS(EID{Optimize.C_PRODUCER.Value,0},EID{self.Id(),0},EID{s.Id(),0}))
    return Result} 
  
// The EID go function for: g_expression @ Call_method1 (throw: true) 
func E_Generate_g_expression_Call_method1 (self EID,s EID) EID { 
    return /*(sm for g_expression @ Call_method1= EID)*/ F_Generate_g_expression_Call_method1(Language.To_CallMethod1(OBJ(self)),ToClass(OBJ(s)) )} 
  
/* {1} OPT.The go function for: g_expression(self:Call_method2,s:class) [] */
func F_Generate_g_expression_Call_method2 (self *Language.CallMethod2 ,s *ClaireClass ) EID { 
    var Result EID 
    Result = Core.F_CALL(C_Generate_inline_exp,ARGS(EID{Optimize.C_PRODUCER.Value,0},EID{self.Id(),0},EID{s.Id(),0}))
    return Result} 
  
// The EID go function for: g_expression @ Call_method2 (throw: true) 
func E_Generate_g_expression_Call_method2 (self EID,s EID) EID { 
    return /*(sm for g_expression @ Call_method2= EID)*/ F_Generate_g_expression_Call_method2(Language.To_CallMethod2(OBJ(self)),ToClass(OBJ(s)) )} 
  
/* {1} OPT.The go function for: g_expression(self:Call_method,s:class) [] */
func F_Generate_g_expression_Call_method (self *Language.CallMethod ,s *ClaireClass ) EID { 
    var Result EID 
    Result = Core.F_CALL(C_Generate_inline_exp,ARGS(EID{Optimize.C_PRODUCER.Value,0},EID{self.Id(),0},EID{s.Id(),0}))
    return Result} 
  
// The EID go function for: g_expression @ Call_method (throw: true) 
func E_Generate_g_expression_Call_method (self EID,s EID) EID { 
    return /*(sm for g_expression @ Call_method= EID)*/ F_Generate_g_expression_Call_method(Language.To_CallMethod(OBJ(self)),ToClass(OBJ(s)) )} 
  
// ---------------------------------------- dynamic call -------------------------------------------------------------------
// new in 3.0 : really low level method are virtual and only rely on inline compiling
// note the *_prefix(s) ... *_postfix(s) that add a conversion from * to exprected type s
// WARNING : we can use assignment (x = y) only when s = void (we do not care for the result)
/* {1} OPT.The go function for: inline_exp(c:go_producer,self:Call,s:class) [] */
func F_Generate_inline_exp_go_producer1 (c *GenerateGoProducer ,self *Language.Call ,s *ClaireClass ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var p *ClaireProperty   = self.Selector
      /* noccur = 11 */
      /* Let:3 */{ 
        var a1 *ClaireAny  
        /* noccur = 10 */
        var a1_try02764 EID 
        a1_try02764 = Core.F_car_list(self.Args)
        /* ERROR PROTECTION INSERTED (a1-Result) */
        if ErrorIn(a1_try02764) {Result = a1_try02764
        } else {
        a1 = ANY(a1_try02764)
        /* Let:4 */{ 
          var n int  = self.Args.Length()
          /* noccur = 2 */
          if (p.Id() == Core.C_mClaire_get_stack.Id()) /* If:5 */{ 
            Result = F_Generate_eid_prefix_class(s)
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC("ClEnv.EvalStack[")
            Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_integer.Id(),0}))
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC("]")
            F_Generate_eid_post_class(s)
            PRINC("")
            Result = EVOID
            }}
            /* If!5 */}  else if (p.Id() == Optimize.C_safe.Id()) /* If:5 */{ 
            /* Let:6 */{ 
              var y int  = Optimize.C_compiler.Safety
              /* noccur = 1 */
              Optimize.C_compiler.Safety = 1
              Result = Core.F_CALL(C_Generate_g_expression,ARGS(self.Args.At(1-1).ToEID(),EID{s.Id(),0}))
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              /* update:7 */{ 
                var va_arg1 *Optimize.OptimizeMetaCompiler  
                var va_arg2 int 
                va_arg1 = Optimize.C_compiler
                va_arg2 = y
                /* ---------- now we compile update safety(va_arg1) := va_arg2 ------- */
                va_arg1.Safety = va_arg2
                Result = EID{C__INT,IVAL(va_arg2)}
                /* update-7 */} 
              }
              /* Let-6 */} 
            /* If!5 */}  else if (p.Id() == Core.C_mClaire_base_I.Id()) /* If:5 */{ 
            F_Generate_integer_prefix_class(s)
            PRINC("ClEnv.Base")
            F_Generate_native_post_class(s)
            PRINC("")
            Result = EVOID
            /* If!5 */}  else if (p.Id() == Core.C_Core__inf_equalt.Id()) /* If:5 */{ 
            F_Generate_object_prefix_class(C_boolean,s)
            Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_type.Id(),0}))
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC(".Included(")
            Result = Core.F_CALL(C_Generate_g_expression,ARGS(self.Args.At(2-1).ToEID(),EID{C_type.Id(),0}))
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC(")")
            F_Generate_object_post_class(C_boolean,s)
            PRINC("")
            Result = EVOID
            }}
            /* If!5 */}  else if ((p.Id() == Core.C_mClaire_index_I.Id()) && 
              (n == 1)) /* If:5 */{ 
            F_Generate_integer_prefix_class(s)
            PRINC("ClEnv.Index")
            F_Generate_native_post_class(s)
            PRINC("")
            Result = EVOID
            /* If!5 */}  else if ((p.Id() == Core.C_mClaire_push_I.Id()) && 
              (n == 1)) /* If:5 */{ 
            PRINC("ClEnv.Push(")
            Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{Optimize.C_EID.Id(),0}))
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC(")")
            Result = EVOID
            }
            /* If!5 */}  else if (p.Id() == Core.C_mClaire_put_stack.Id()) /* If:5 */{ 
            if (s.Id() != C_void.Id()) /* If:6 */{ 
              Core.F_tformat_string(MakeString("WARNING: use ~S in non void context\n"),0,MakeConstantList(self.Id()))
              /* If-6 */} 
            PRINC("ClEnv.EvalStack[")
            Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_integer.Id(),0}))
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC("]=")
            Result = Core.F_CALL(C_Generate_g_expression,ARGS(self.Args.At(2-1).ToEID(),EID{Optimize.C_EID.Id(),0}))
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC("")
            Result = EVOID
            }}
            /* If!5 */}  else if ((p.Id() == Core.C_mClaire_set_base.Id()) && 
              (s.Id() == C_void.Id())) /* If:5 */{ 
            PRINC("ClEnv.Base= ")
            Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_integer.Id(),0}))
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC("")
            Result = EVOID
            }
            /* If!5 */}  else if ((p.Id() == Core.C_mClaire_set_index.Id()) && 
              (s.Id() == C_void.Id())) /* If:5 */{ 
            PRINC("ClEnv.Index= ")
            Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_integer.Id(),0}))
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC("")
            Result = EVOID
            }
            /* If!5 */}  else if (p.Id() == Optimize.C_Compile_anyObject_I.Id()) /* If:5 */{ 
            if (a1 == C_Interval.Id()) /* If:6 */{ 
              F_Generate_object_prefix_class(C_any,s)
              F_Generate_class_ident_class(ToClass(a1))
              PRINC(".MakeInts(")
              /* Let:7 */{ 
                var g0277UU *ClaireList  
                /* noccur = 1 */
                var g0277UU_try02788 EID 
                g0277UU_try02788 = self.Args.Cdr()
                /* ERROR PROTECTION INSERTED (g0277UU-Result) */
                if ErrorIn(g0277UU_try02788) {Result = g0277UU_try02788
                } else {
                g0277UU = ToList(OBJ(g0277UU_try02788))
                Result = F_Generate_args_list_list(g0277UU,C_integer)
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              PRINC(")")
              F_Generate_object_post_class(C_any,s)
              PRINC("")
              Result = EVOID
              }
              } else {
              F_Generate_object_prefix_class(C_any,s)
              F_Generate_class_ident_class(ToClass(a1))
              PRINC(".Make(")
              /* Let:7 */{ 
                var g0279UU *ClaireList  
                /* noccur = 1 */
                var g0279UU_try02808 EID 
                g0279UU_try02808 = self.Args.Cdr()
                /* ERROR PROTECTION INSERTED (g0279UU-Result) */
                if ErrorIn(g0279UU_try02808) {Result = g0279UU_try02808
                } else {
                g0279UU = ToList(OBJ(g0279UU_try02808))
                Result = F_Generate_args_list_list(g0279UU,C_any)
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              PRINC(")")
              F_Generate_object_post_class(C_any,s)
              PRINC("")
              Result = EVOID
              }
              /* If-6 */} 
            /* If!5 */}  else if (p.Id() == C_add_slot.Id()) /* If:5 */{ 
            Result = F_Generate_cast_prefix_class(C_slot,s)
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_class.Id(),0}))
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC(".AddSlot(")
            Result = Core.F_CALL(C_Generate_g_expression,ARGS(self.Args.At(2-1).ToEID(),EID{C_property.Id(),0}))
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC(",")
            Result = Core.F_CALL(C_Generate_g_expression,ARGS(self.Args.At(3-1).ToEID(),EID{C_type.Id(),0}))
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC(",")
            Result = Core.F_CALL(C_Generate_g_expression,ARGS(self.Args.At(4-1).ToEID(),EID{C_any.Id(),0}))
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC(")")
            F_Generate_cast_post_class(C_slot,s)
            PRINC("")
            Result = EVOID
            }}}}}
            } else {
            Result = F_Generate_eid_prefix_class(s)
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            F_Generate_preCore_ask_void()
            PRINC("F_CALL(")
            Result = F_Generate_g_expression_thing(ToThing(self.Selector.Id()),C_property)
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC(",ARGS(")
            Result = F_Generate_args_list_list(self.Args,Optimize.C_EID)
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC("))")
            F_Generate_eid_post_class(s)
            PRINC("")
            Result = EVOID
            }}}
            /* If-5 */} 
          /* Let-4 */} 
        }
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: inline_exp @ list<type_expression>(go_producer, Call, class) (throw: true) 
func E_Generate_inline_exp_go_producer1 (c EID,self EID,s EID) EID { 
    return /*(sm for inline_exp @ list<type_expression>(go_producer, Call, class)= EID)*/ F_Generate_inline_exp_go_producer1(ToGenerateGoProducer(OBJ(c)),Language.To_Call(OBJ(self)),ToClass(OBJ(s)) )} 
  
// produces a list of C expressions, separated by commas
/* {1} OPT.The go function for: args_list(self:list,s:class) [] */
func F_Generate_args_list_list (self *ClaireList ,s *ClaireClass ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var _Zfirst *ClaireBoolean   = CTRUE
      /* noccur = 2 */
      /* Let:3 */{ 
        var bk_ask *ClaireBoolean   = Core.F__sup_integer(self.Length(),3)
        /* noccur = 3 */
        if (bk_ask == CTRUE) /* If:4 */{ 
          Optimize.C_OPT.Level = (Optimize.C_OPT.Level+1)
          /* If-4 */} 
        /* For:4 */{ 
          var x *ClaireAny  
          _ = x
          Result= EID{CFALSE.Id(),0}
          var x_support *ClaireList  
          x_support = self
          x_len := x_support.Length()
          for i_it := 0; i_it < x_len; i_it++ { 
            x = x_support.At(i_it)
            var void_try6 EID 
            _ = void_try6
            if (_Zfirst == CTRUE) /* If:6 */{ 
              void_try6 = Core.F_CALL(C_Generate_g_expression,ARGS(x.ToEID(),EID{s.Id(),0}))
              /* ERROR PROTECTION INSERTED (void_try6-void_try6) */
              if ErrorIn(void_try6) {Result = void_try6
              break
              } else {
              _Zfirst = CFALSE
              void_try6 = EID{_Zfirst.Id(),0}
              }
              } else {
              PRINC(",")
              if (bk_ask == CTRUE) /* If:7 */{ 
                F_Generate_breakline_void()
                /* If-7 */} 
              void_try6 = Core.F_CALL(C_Generate_g_expression,ARGS(x.ToEID(),EID{s.Id(),0}))
              /* ERROR PROTECTION INSERTED (void_try6-void_try6) */
              if ErrorIn(void_try6) {Result = void_try6
              break
              } else {
              PRINC("")
              void_try6 = EVOID
              }
              /* If-6 */} 
            /* ERROR PROTECTION INSERTED (void_try6-Result) */
            if ErrorIn(void_try6) {Result = void_try6
            Result = void_try6
            break
            } else {
            }
            /* loop-5 */} 
          /* For-4 */} 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        if (bk_ask == CTRUE) /* If:4 */{ 
          /* update:5 */{ 
            var va_arg1 *Optimize.OptimizeMetaOPT  
            var va_arg2 int 
            va_arg1 = Optimize.C_OPT
            va_arg2 = (Optimize.C_OPT.Level-1)
            /* ---------- now we compile update Compile/level(va_arg1) := va_arg2 ------- */
            va_arg1.Level = va_arg2
            Result = EID{C__INT,IVAL(va_arg2)}
            /* update-5 */} 
          } else {
          Result = EID{CFALSE.Id(),0}
          /* If-4 */} 
        }
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: args_list @ list (throw: true) 
func E_Generate_args_list_list (self EID,s EID) EID { 
    return /*(sm for args_list @ list= EID)*/ F_Generate_args_list_list(ToList(OBJ(self)),ToClass(OBJ(s)) )} 
  
// CLAIRE4 : get rid of fast dispatch (fcall + dispatcher)
// Super is like a call
/* {1} OPT.The go function for: g_expression(self:Super,s:class) [] */
func F_Generate_g_expression_Super (self *Language.Super ,s *ClaireClass ) EID { 
    var Result EID 
    Result = F_Generate_eid_prefix_class(s)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    F_Generate_preCore_ask_void()
    PRINC("F_SUPER(")
    Result = F_Generate_g_expression_thing(ToThing(self.Selector.Id()),C_property)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(", ")
    Result = Core.F_CALL(C_Generate_g_expression,ARGS(EID{self.CastTo.Id(),0},EID{C_class.Id(),0}))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(", ARGS(")
    Result = F_Generate_args_list_list(self.Args,Optimize.C_EID)
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC("))")
    F_Generate_eid_post_class(s)
    PRINC("")
    Result = EVOID
    }}}}
    return Result} 
  
// The EID go function for: g_expression @ Super (throw: true) 
func E_Generate_g_expression_Super (self EID,s EID) EID { 
    return /*(sm for g_expression @ Super= EID)*/ F_Generate_g_expression_Super(Language.To_Super(OBJ(self)),ToClass(OBJ(s)) )} 
  
// *******************************************************************
// *       Part 3: the inline coding of function calls               *
// *******************************************************************
// CLAIRE4 Note : all inline optimization assume than can_throw?(m) = false
// these methods are important since they contain the open-coding optimisations. Some of the method calls are be replaced
// directly by  expressions. We always expect the native form (the sort s is passed as a parameter)
// functions with one argument
// note that we need the *_prefix / *_post 
/* {1} OPT.The go function for: inline_exp(c:go_producer,self:Call_method1,s:class) [] */
func F_Generate_inline_exp_go_producer2 (c *GenerateGoProducer ,self *Language.CallMethod1 ,s *ClaireClass ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var m *ClaireMethod   = self.Arg
      /* noccur = 13 */
      /* Let:3 */{ 
        var p *ClaireProperty   = m.Selector
        /* noccur = 5 */
        /* Let:4 */{ 
          var a1 *ClaireAny  
          /* noccur = 25 */
          var a1_try02835 EID 
          a1_try02835 = Core.F_car_list(self.Args)
          /* ERROR PROTECTION INSERTED (a1-Result) */
          if ErrorIn(a1_try02835) {Result = a1_try02835
          } else {
          a1 = ANY(a1_try02835)
          /* Let:5 */{ 
            var dm *ClaireClass   = Core.F_domain_I_restriction(ToRestriction(m.Id()))
            /* noccur = 4 */
            if ((p.Id() == C__dash.Id()) && 
                ((dm.Id() == C_integer.Id()) || 
                    (dm.Id() == C_float.Id()))) /* If:6 */{ 
              Result = F_Generate_cast_prefix_class(dm,s)
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              PRINC("(-")
              Result = F_Generate_bounded_expression_any(a1,s)
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              PRINC(")")
              F_Generate_cast_post_class(dm,s)
              PRINC("")
              Result = EVOID
              }}
              /* If!6 */}  else if ((p.Id() == Core.C_owner.Id()) && 
                (F_Generate_eid_provide_ask_any(a1) == CTRUE)) /* If:6 */{ 
              F_Generate_object_prefix_class(C_class,s)
              PRINC("OWNER(")
              Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{Optimize.C_EID.Id(),0}))
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              PRINC(")")
              F_Generate_object_post_class(C_class,s)
              PRINC("")
              Result = EVOID
              }
              } else {
              var g0284I *ClaireBoolean  
              var g0284I_try02857 EID 
              /* and:7 */{ 
                var v_and7 *ClaireBoolean  
                
                v_and7 = Equal(p.Id(),Core.C_owner.Id())
                if (v_and7 == CFALSE) {g0284I_try02857 = EID{CFALSE.Id(),0}
                } else /* arg:8 */{ 
                  var v_and7_try02869 EID 
                  v_and7_try02869 = Optimize.F_Compile_designated_ask_any(a1)
                  /* ERROR PROTECTION INSERTED (v_and7-g0284I_try02857) */
                  if ErrorIn(v_and7_try02869) {g0284I_try02857 = v_and7_try02869
                  } else {
                  v_and7 = ToBoolean(OBJ(v_and7_try02869))
                  if (v_and7 == CFALSE) {g0284I_try02857 = EID{CFALSE.Id(),0}
                  } else /* arg:9 */{ 
                    g0284I_try02857 = EID{CTRUE.Id(),0}/* arg-9 */} 
                  /* arg-8 */} 
                }
                /* and-7 */} 
              /* ERROR PROTECTION INSERTED (g0284I-Result) */
              if ErrorIn(g0284I_try02857) {Result = g0284I_try02857
              } else {
              g0284I = ToBoolean(OBJ(g0284I_try02857))
              if (g0284I == CTRUE) /* If:7 */{ 
                F_Generate_object_prefix_class(C_class,s)
                Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_any.Id(),0}))
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                PRINC(".Isa")
                F_Generate_object_post_class(C_class,s)
                PRINC("")
                Result = EVOID
                }
                /* If!7 */}  else if (p.Id() == Core.C_eval.Id()) /* If:7 */{ 
                Result = F_Generate_eid_prefix_class(s)
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                PRINC("EVAL(")
                Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_any.Id(),0}))
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                PRINC(")")
                F_Generate_eid_post_class(s)
                PRINC("")
                Result = EVOID
                }}
                /* If!7 */}  else if (m.Selector.Id() == Core.C_externC.Id()) /* If:7 */{ 
                Result = Core.F_CALL(C_princ,ARGS(a1.ToEID()))
                } else {
                var g0287I *ClaireBoolean  
                var g0287I_try02888 EID 
                /* and:8 */{ 
                  var v_and8 *ClaireBoolean  
                  
                  v_and8 = Equal(m.Id(),C_Generate__starlength_bag_star.Value)
                  if (v_and8 == CFALSE) {g0287I_try02888 = EID{CFALSE.Id(),0}
                  } else /* arg:9 */{ 
                    var v_and8_try028910 EID 
                    v_and8_try028910 = Optimize.F_Compile_designated_ask_any(a1)
                    /* ERROR PROTECTION INSERTED (v_and8-g0287I_try02888) */
                    if ErrorIn(v_and8_try028910) {g0287I_try02888 = v_and8_try028910
                    } else {
                    v_and8 = ToBoolean(OBJ(v_and8_try028910))
                    if (v_and8 == CFALSE) {g0287I_try02888 = EID{CFALSE.Id(),0}
                    } else /* arg:10 */{ 
                      g0287I_try02888 = EID{CTRUE.Id(),0}/* arg-10 */} 
                    /* arg-9 */} 
                  }
                  /* and-8 */} 
                /* ERROR PROTECTION INSERTED (g0287I-Result) */
                if ErrorIn(g0287I_try02888) {Result = g0287I_try02888
                } else {
                g0287I = ToBoolean(OBJ(g0287I_try02888))
                if (g0287I == CTRUE) /* If:8 */{ 
                  F_Generate_integer_prefix_class(s)
                  Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_list.Id(),0}))
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  PRINC(".Length()")
                  F_Generate_native_post_class(s)
                  PRINC("")
                  Result = EVOID
                  }
                  } else {
                  var g0290I *ClaireBoolean  
                  var g0290I_try02919 EID 
                  /* and:9 */{ 
                    var v_and9 *ClaireBoolean  
                    
                    v_and9 = Equal(p.Id(),C_integer_I.Id())
                    if (v_and9 == CFALSE) {g0290I_try02919 = EID{CFALSE.Id(),0}
                    } else /* arg:10 */{ 
                      v_and9 = Equal(Core.F_domain_I_restriction(ToRestriction(m.Id())).Id(),C_char.Id())
                      if (v_and9 == CFALSE) {g0290I_try02919 = EID{CFALSE.Id(),0}
                      } else /* arg:11 */{ 
                        var v_and9_try029212 EID 
                        v_and9_try029212 = Optimize.F_Compile_designated_ask_any(a1)
                        /* ERROR PROTECTION INSERTED (v_and9-g0290I_try02919) */
                        if ErrorIn(v_and9_try029212) {g0290I_try02919 = v_and9_try029212
                        } else {
                        v_and9 = ToBoolean(OBJ(v_and9_try029212))
                        if (v_and9 == CFALSE) {g0290I_try02919 = EID{CFALSE.Id(),0}
                        } else /* arg:12 */{ 
                          g0290I_try02919 = EID{CTRUE.Id(),0}/* arg-12 */} 
                        /* arg-11 */} 
                      /* arg-10 */} 
                    }
                    /* and-9 */} 
                  /* ERROR PROTECTION INSERTED (g0290I-Result) */
                  if ErrorIn(g0290I_try02919) {Result = g0290I_try02919
                  } else {
                  g0290I = ToBoolean(OBJ(g0290I_try02919))
                  if (g0290I == CTRUE) /* If:9 */{ 
                    F_Generate_integer_prefix_class(s)
                    PRINC("int(")
                    Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_char.Id(),0}))
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    PRINC(")")
                    F_Generate_native_post_class(s)
                    PRINC("")
                    Result = EVOID
                    }
                    /* If!9 */}  else if ((m.Id() == C_Generate__starof_bag_star.Value) || 
                      (m.Id() == C_Generate__starof_array_star.Value)) /* If:9 */{ 
                    Result = F_Generate_cast_prefix_class(C_type,s)
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_list.Id(),0}))
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    PRINC(".Of()")
                    F_Generate_cast_post_class(C_type,s)
                    PRINC("")
                    Result = EVOID
                    }}
                    /* If!9 */}  else if (m.Id() == Optimize.C_Compile_m_unsafe.Value) /* If:9 */{ 
                    if (s.Id() == Optimize.C_EID.Id()) /* If:10 */{ 
                      Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{Optimize.C_EID.Id(),0}))
                      } else {
                      var g0293I *ClaireBoolean  
                      var g0293I_try029411 EID 
                      /* Let:11 */{ 
                        var g0295UU *ClaireBoolean  
                        /* noccur = 1 */
                        var g0295UU_try029612 EID 
                        g0295UU_try029612 = Optimize.F_Compile_g_throw_any(a1)
                        /* ERROR PROTECTION INSERTED (g0295UU-g0293I_try029411) */
                        if ErrorIn(g0295UU_try029612) {g0293I_try029411 = g0295UU_try029612
                        } else {
                        g0295UU = ToBoolean(OBJ(g0295UU_try029612))
                        g0293I_try029411 = EID{g0295UU.Not.Id(),0}
                        }
                        /* Let-11 */} 
                      /* ERROR PROTECTION INSERTED (g0293I-Result) */
                      if ErrorIn(g0293I_try029411) {Result = g0293I_try029411
                      } else {
                      g0293I = ToBoolean(OBJ(g0293I_try029411))
                      if (g0293I == CTRUE) /* If:11 */{ 
                        Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{s.Id(),0}))
                        } else {
                        Result = F_Generate_cast_prefix_class(C_any,s)
                        /* ERROR PROTECTION INSERTED (Result-Result) */
                        if !ErrorIn(Result) {
                        PRINC("ANY(")
                        Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{Optimize.C_EID.Id(),0}))
                        /* ERROR PROTECTION INSERTED (Result-Result) */
                        if !ErrorIn(Result) {
                        PRINC(")")
                        F_Generate_cast_post_class(C_any,s)
                        PRINC("")
                        Result = EVOID
                        }}
                        /* If-11 */} 
                      }
                      /* If-10 */} 
                    /* If!9 */}  else if ((m.Id() == C_Generate__starprinc_string_star.Value) && 
                      (C_string.Id() == a1.Isa.Id())) /* If:9 */{ 
                    PRINC("PRINC(")
                    Result = Core.F_CALL(C_print,ARGS(a1.ToEID()))
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    PRINC(")")
                    Result = EVOID
                    }
                    /* If!9 */}  else if (m.Id() == C_Generate__starcopy_list_star.Value) /* If:9 */{ 
                    Result = F_Generate_cast_prefix_class(C_list,s)
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_list.Id(),0}))
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    PRINC(".Copy()")
                    F_Generate_cast_post_class(C_list,s)
                    PRINC("")
                    Result = EVOID
                    }}
                    /* If!9 */}  else if (m.Id() == C_Generate__starlength_array_star.Value) /* If:9 */{ 
                    F_Generate_integer_prefix_class(s)
                    Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_array.Id(),0}))
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    PRINC(".Length()")
                    F_Generate_native_post_class(s)
                    PRINC("")
                    Result = EVOID
                    }
                    } else {
                    var g0297I *ClaireBoolean  
                    var g0297I_try029810 EID 
                    /* and:10 */{ 
                      var v_and10 *ClaireBoolean  
                      
                      v_and10 = Equal(m.Id(),C_Generate__starnot_star.Value)
                      if (v_and10 == CFALSE) {g0297I_try029810 = EID{CFALSE.Id(),0}
                      } else /* arg:11 */{ 
                        var v_and10_try029912 EID 
                        /* Let:12 */{ 
                          var g0300UU *ClaireClass  
                          /* noccur = 1 */
                          var g0300UU_try030113 EID 
                          g0300UU_try030113 = Language.F_static_type_any(a1)
                          /* ERROR PROTECTION INSERTED (g0300UU-v_and10_try029912) */
                          if ErrorIn(g0300UU_try030113) {v_and10_try029912 = g0300UU_try030113
                          } else {
                          g0300UU = ToClass(OBJ(g0300UU_try030113))
                          v_and10_try029912 = EID{ToType(g0300UU.Id()).Included(ToType(C_boolean.Id())).Id(),0}
                          }
                          /* Let-12 */} 
                        /* ERROR PROTECTION INSERTED (v_and10-g0297I_try029810) */
                        if ErrorIn(v_and10_try029912) {g0297I_try029810 = v_and10_try029912
                        } else {
                        v_and10 = ToBoolean(OBJ(v_and10_try029912))
                        if (v_and10 == CFALSE) {g0297I_try029810 = EID{CFALSE.Id(),0}
                        } else /* arg:12 */{ 
                          g0297I_try029810 = EID{CTRUE.Id(),0}/* arg-12 */} 
                        /* arg-11 */} 
                      }
                      /* and-10 */} 
                    /* ERROR PROTECTION INSERTED (g0297I-Result) */
                    if ErrorIn(g0297I_try029810) {Result = g0297I_try029810
                    } else {
                    g0297I = ToBoolean(OBJ(g0297I_try029810))
                    if (g0297I == CTRUE) /* If:10 */{ 
                      F_Generate_object_prefix_class(C_boolean,s)
                      Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_boolean.Id(),0}))
                      /* ERROR PROTECTION INSERTED (Result-Result) */
                      if !ErrorIn(Result) {
                      PRINC(".Not")
                      F_Generate_object_post_class(C_boolean,s)
                      PRINC("")
                      Result = EVOID
                      }
                      /* If!10 */}  else if ((m.Id() == C_Generate__starnew_class1_star.Value) && 
                        (C_class.Id() == a1.Isa.Id())) /* If:10 */{ 
                      F_Generate_object_prefix_class(C_any,s)
                      PRINC("new(")
                      F_Generate_go_class_class(ToClass(a1))
                      PRINC(").Is(")
                      Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_class.Id(),0}))
                      /* ERROR PROTECTION INSERTED (Result-Result) */
                      if !ErrorIn(Result) {
                      PRINC(")")
                      F_Generate_object_post_class(C_any,s)
                      PRINC("")
                      Result = EVOID
                      }
                      } else {
                      Result = c.PrintExternalCall(Language.To_CallMethod(self.Id()),s)
                      /* If-10 */} 
                    }
                    /* If-9 */} 
                  }
                  /* If-8 */} 
                }
                /* If-7 */} 
              }
              /* If-6 */} 
            /* Let-5 */} 
          }
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: inline_exp @ list<type_expression>(go_producer, Call_method1, class) (throw: true) 
func E_Generate_inline_exp_go_producer2 (c EID,self EID,s EID) EID { 
    return /*(sm for inline_exp @ list<type_expression>(go_producer, Call_method1, class)= EID)*/ F_Generate_inline_exp_go_producer2(ToGenerateGoProducer(OBJ(c)),Language.To_CallMethod1(OBJ(self)),ToClass(OBJ(s)) )} 
  
// ===  functions with two arguments ===
/* {1} OPT.The go function for: inline_exp(c:go_producer,self:Call_method2,s:class) [] */
func F_Generate_inline_exp_go_producer3 (c *GenerateGoProducer ,self *Language.CallMethod2 ,s *ClaireClass ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var m *ClaireMethod   = self.Arg
      /* noccur = 31 */
      /* Let:3 */{ 
        var p *ClaireProperty   = m.Selector
        /* noccur = 6 */
        /* Let:4 */{ 
          var a1 *ClaireAny   = self.Args.At(1-1)
          /* noccur = 26 */
          /* Let:5 */{ 
            var a2 *ClaireAny   = self.Args.At(2-1)
            /* noccur = 19 */
            /* Let:6 */{ 
              var s1 *ClaireClass  
              /* noccur = 6 */
              var s1_try03027 EID 
              /* Let:7 */{ 
                var g0303UU *ClaireType  
                /* noccur = 1 */
                var g0303UU_try03048 EID 
                g0303UU_try03048 = Core.F_CALL(Optimize.C_c_type,ARGS(a1.ToEID()))
                /* ERROR PROTECTION INSERTED (g0303UU-s1_try03027) */
                if ErrorIn(g0303UU_try03048) {s1_try03027 = g0303UU_try03048
                } else {
                g0303UU = ToType(OBJ(g0303UU_try03048))
                s1_try03027 = EID{g0303UU.Class_I().Id(),0}
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (s1-Result) */
              if ErrorIn(s1_try03027) {Result = s1_try03027
              } else {
              s1 = ToClass(OBJ(s1_try03027))
              if ((p.Id() == C_class_I.Id()) && 
                  (a1.Isa.IsIn(C_symbol) == CTRUE)) /* If:7 */{ 
                F_Generate_symbol_ident_symbol(ToSymbol(a1))
                PRINC(" = MakeClass(")
                /* Let:8 */{ 
                  var g0305UU *ClaireAny  
                  /* noccur = 1 */
                  var g0305UU_try03069 EID 
                  g0305UU_try03069 = Core.F_CALL(C_string_I,ARGS(a1.ToEID()))
                  /* ERROR PROTECTION INSERTED (g0305UU-Result) */
                  if ErrorIn(g0305UU_try03069) {Result = g0305UU_try03069
                  } else {
                  g0305UU = ANY(g0305UU_try03069)
                  Result = Core.F_print_any(g0305UU)
                  }
                  /* Let-8 */} 
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                PRINC(",")
                Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{C_class.Id(),0}))
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                PRINC(",")
                Result = F_Generate_g_expression_module(ToModule(OBJ(Core.F_CALL(C_module_I,ARGS(a1.ToEID())))),C_module)
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                PRINC(")")
                Result = EVOID
                }}}
                /* If!7 */}  else if ((c.OpenOperators.Memq(p.Id()) == CTRUE) && 
                  ((Equal(m.Domain.ValuesO()[1-1],m.Domain.ValuesO()[2-1]) == CTRUE) && 
                    ((s1.Id() == C_integer.Id()) || 
                        (s1.Id() == C_float.Id())))) /* If:7 */{ 
                Result = F_Generate_cast_prefix_class(s1,s)
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                PRINC("(")
                Result = F_Generate_bounded_expression_any(a1,s1)
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                F_princ_string(p.Name.String_I())
                Result = F_Generate_bounded_expression_any(a2,s1)
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                PRINC(")")
                F_Generate_cast_post_class(s1,s)
                PRINC("")
                Result = EVOID
                }}}
                } else {
                var g0307I *ClaireBoolean  
                var g0307I_try03088 EID 
                /* and:8 */{ 
                  var v_and8 *ClaireBoolean  
                  
                  v_and8 = Equal(m.Id(),C_Generate__starcontain_star.Value)
                  if (v_and8 == CFALSE) {g0307I_try03088 = EID{CFALSE.Id(),0}
                  } else /* arg:9 */{ 
                    var v_and8_try030910 EID 
                    v_and8_try030910 = Optimize.F_Compile_identifiable_ask_any(a2)
                    /* ERROR PROTECTION INSERTED (v_and8-g0307I_try03088) */
                    if ErrorIn(v_and8_try030910) {g0307I_try03088 = v_and8_try030910
                    } else {
                    v_and8 = ToBoolean(OBJ(v_and8_try030910))
                    if (v_and8 == CFALSE) {g0307I_try03088 = EID{CFALSE.Id(),0}
                    } else /* arg:10 */{ 
                      g0307I_try03088 = EID{CTRUE.Id(),0}/* arg-10 */} 
                    /* arg-9 */} 
                  }
                  /* and-8 */} 
                /* ERROR PROTECTION INSERTED (g0307I-Result) */
                if ErrorIn(g0307I_try03088) {Result = g0307I_try03088
                } else {
                g0307I = ToBoolean(OBJ(g0307I_try03088))
                if (g0307I == CTRUE) /* If:8 */{ 
                  F_Generate_object_prefix_class(C_boolean,s)
                  Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_list.Id(),0}))
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  PRINC(".Memq(")
                  Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{C_any.Id(),0}))
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  PRINC(")")
                  F_Generate_object_post_class(C_boolean,s)
                  PRINC("")
                  Result = EVOID
                  }}
                  /* If!8 */}  else if (m.Selector.Id() == Core.C_externC.Id()) /* If:8 */{ 
                  Result = Core.F_CALL(C_princ,ARGS(a1.ToEID()))
                  /* If!8 */}  else if (m.Id() == Optimize.C_Compile_m_member.Value) /* If:8 */{ 
                  Result = F_Generate_belong_exp_any(a1,a2,s)
                  /* If!8 */}  else if ((m.Id() == C_Generate__starwrite_value_star.Value) && 
                    (F_Generate_eid_provide_ask_any(a2) == CTRUE)) /* If:8 */{ 
                  Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_Variable.Id(),0}))
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  PRINC(".WriteEID(")
                  Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{Optimize.C_EID.Id(),0}))
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  PRINC(")")
                  Result = EVOID
                  }}
                  /* If!8 */}  else if (m.Id() == C_Generate__starinherit_star.Value) /* If:8 */{ 
                  F_Generate_object_prefix_class(C_boolean,s)
                  Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_class.Id(),0}))
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  PRINC(".IsIn(")
                  Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{C_class.Id(),0}))
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  PRINC(")")
                  F_Generate_object_post_class(C_boolean,s)
                  PRINC("")
                  Result = EVOID
                  }}
                  /* If!8 */}  else if (m.Id() == C_Generate__starequal_star.Value) /* If:8 */{ 
                  F_Generate_object_prefix_class(C_boolean,s)
                  PRINC("Equal(")
                  Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_any.Id(),0}))
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  PRINC(",")
                  Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{C_any.Id(),0}))
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  PRINC(")")
                  F_Generate_object_post_class(C_boolean,s)
                  PRINC("")
                  Result = EVOID
                  }}
                  /* If!8 */}  else if (m.Id() == C_Generate__starmap_star.Value) /* If:8 */{ 
                  Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_type.Id(),0}))
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  PRINC(".Map_I(")
                  Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{C_type.Id(),0}))
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  PRINC(")")
                  Result = EVOID
                  }}
                  /* If!8 */}  else if (m.Id() == C_Generate__star_Zt_star.Value) /* If:8 */{ 
                  F_Generate_object_prefix_class(C_boolean,s)
                  Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{C_type.Id(),0}))
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  PRINC(".Contains(")
                  Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_any.Id(),0}))
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  PRINC(")")
                  F_Generate_object_post_class(C_boolean,s)
                  PRINC("")
                  Result = EVOID
                  }}
                  /* If!8 */}  else if ((p.Id() == Core.C_Core__inf_equalt.Id()) || 
                    (m.Id() == C_Generate__starincluded_star.Value)) /* If:8 */{ 
                  F_Generate_object_prefix_class(C_boolean,s)
                  Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_type.Id(),0}))
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  PRINC(".Included(")
                  Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{C_type.Id(),0}))
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  PRINC(")")
                  F_Generate_object_post_class(C_boolean,s)
                  PRINC("")
                  Result = EVOID
                  }}
                  } else {
                  var g0310I *ClaireBoolean  
                  var g0310I_try03119 EID 
                  /* or:9 */{ 
                    var v_or9 *ClaireBoolean  
                    
                    var v_or9_try031210 EID 
                    /* and:10 */{ 
                      var v_and10 *ClaireBoolean  
                      
                      v_and10 = MakeBoolean((((m.Id() == C_Generate__starnth_list_star.Value) || 
                            (m.Id() == C_Generate__starnth_tuple_star.Value)) && 
                          (Optimize.C_compiler.Safety >= 2)) || (m.Id() == C_Generate__starnth_1_list_star.Value) || (m.Id() == C_Generate__starnth_1_tuple_star.Value) || (m.Id() == C_Generate__starnth_1_array_star.Value))
                      if (v_and10 == CFALSE) {v_or9_try031210 = EID{CFALSE.Id(),0}
                      } else /* arg:11 */{ 
                        var v_and10_try031312 EID 
                        /* Let:12 */{ 
                          var g0314UU *ClaireClass  
                          /* noccur = 1 */
                          var g0314UU_try031513 EID 
                          g0314UU_try031513 = F_Generate_g_member_any(a1)
                          /* ERROR PROTECTION INSERTED (g0314UU-v_and10_try031312) */
                          if ErrorIn(g0314UU_try031513) {v_and10_try031312 = g0314UU_try031513
                          } else {
                          g0314UU = ToClass(OBJ(g0314UU_try031513))
                          v_and10_try031312 = EID{Core.F__I_equal_any(g0314UU.Id(),C_any.Id()).Id(),0}
                          }
                          /* Let-12 */} 
                        /* ERROR PROTECTION INSERTED (v_and10-v_or9_try031210) */
                        if ErrorIn(v_and10_try031312) {v_or9_try031210 = v_and10_try031312
                        } else {
                        v_and10 = ToBoolean(OBJ(v_and10_try031312))
                        if (v_and10 == CFALSE) {v_or9_try031210 = EID{CFALSE.Id(),0}
                        } else /* arg:12 */{ 
                          v_or9_try031210 = EID{CTRUE.Id(),0}/* arg-12 */} 
                        /* arg-11 */} 
                      }
                      /* and-10 */} 
                    /* ERROR PROTECTION INSERTED (v_or9-g0310I_try03119) */
                    if ErrorIn(v_or9_try031210) {g0310I_try03119 = v_or9_try031210
                    } else {
                    v_or9 = ToBoolean(OBJ(v_or9_try031210))
                    if (v_or9 == CTRUE) {g0310I_try03119 = EID{CTRUE.Id(),0}
                    } else /* or:10 */{ 
                      v_or9 = Equal(m.Selector.Id(),Core.C_mClaire_nth_object.Id())
                      if (v_or9 == CTRUE) {g0310I_try03119 = EID{CTRUE.Id(),0}
                      } else /* or:11 */{ 
                        g0310I_try03119 = EID{CFALSE.Id(),0}/* org-11 */} 
                      /* org-10 */} 
                    }
                    /* or-9 */} 
                  /* ERROR PROTECTION INSERTED (g0310I-Result) */
                  if ErrorIn(g0310I_try03119) {Result = g0310I_try03119
                  } else {
                  g0310I = ToBoolean(OBJ(g0310I_try03119))
                  if (g0310I == CTRUE) /* If:9 */{ 
                    /* Let:10 */{ 
                      var s1 *ClaireClass  
                      /* noccur = 2 */
                      var s1_try031611 EID 
                      if (m.Selector.Id() == Core.C_mClaire_nth_object.Id()) /* If:11 */{ 
                        s1_try031611 = EID{C_object.Id(),0}
                        } else {
                        /* Let:12 */{ 
                          var g0317UU *ClaireClass  
                          /* noccur = 1 */
                          var g0317UU_try031813 EID 
                          g0317UU_try031813 = F_Generate_g_member_any(a1)
                          /* ERROR PROTECTION INSERTED (g0317UU-s1_try031611) */
                          if ErrorIn(g0317UU_try031813) {s1_try031611 = g0317UU_try031813
                          } else {
                          g0317UU = ToClass(OBJ(g0317UU_try031813))
                          s1_try031611 = EID{F_Generate_type_sort_type(ToType(g0317UU.Id())).Id(),0}
                          }
                          /* Let-12 */} 
                        /* If-11 */} 
                      /* ERROR PROTECTION INSERTED (s1-Result) */
                      if ErrorIn(s1_try031611) {Result = s1_try031611
                      } else {
                      s1 = ToClass(OBJ(s1_try031611))
                      Result = F_Generate_cast_prefix_class(s1,s)
                      /* ERROR PROTECTION INSERTED (Result-Result) */
                      if !ErrorIn(Result) {
                      Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_list.Id(),0}))
                      /* ERROR PROTECTION INSERTED (Result-Result) */
                      if !ErrorIn(Result) {
                      PRINC(".")
                      /* Let:11 */{ 
                        var g0319UU *ClaireClass  
                        /* noccur = 1 */
                        var g0319UU_try032012 EID 
                        g0319UU_try032012 = F_Generate_g_member_any(a1)
                        /* ERROR PROTECTION INSERTED (g0319UU-Result) */
                        if ErrorIn(g0319UU_try032012) {Result = g0319UU_try032012
                        } else {
                        g0319UU = ToClass(OBJ(g0319UU_try032012))
                        F_Generate_valuesSlot_class(g0319UU)
                        Result = EVOID
                        }
                        /* Let-11 */} 
                      /* ERROR PROTECTION INSERTED (Result-Result) */
                      if !ErrorIn(Result) {
                      PRINC("[")
                      Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{C_integer.Id(),0}))
                      /* ERROR PROTECTION INSERTED (Result-Result) */
                      if !ErrorIn(Result) {
                      PRINC("-1]")
                      F_Generate_cast_post_class(s1,s)
                      PRINC("")
                      Result = EVOID
                      }}}}
                      }
                      /* Let-10 */} 
                    /* If!9 */}  else if ((m.Id() == C_Generate__starnth_list_star.Value) || 
                      ((m.Id() == C_Generate__starnth_tuple_star.Value) || 
                        ((m.Id() == C_Generate__starnth_1_list_star.Value) || 
                          ((m.Id() == C_Generate__starnth_1_tuple_star.Value) || 
                            (m.Id() == C_Generate__starnth_1_array_star.Value))))) /* If:9 */{ 
                    Result = F_Generate_cast_prefix_class(C_any,s)
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_list.Id(),0}))
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    PRINC(".At(")
                    Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{C_integer.Id(),0}))
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    PRINC("-1)")
                    F_Generate_cast_post_class(C_any,s)
                    PRINC("")
                    Result = EVOID
                    }}}
                    /* If!9 */}  else if ((p.Id() == C_add_I.Id()) && 
                      (ToType(Core.F_domain_I_restriction(ToRestriction(m.Id())).Id()).Included(ToType(C_bag.Id())) == CTRUE)) /* If:9 */{ 
                    /* Let:10 */{ 
                      var sbag *ClaireClass  
                      /* noccur = 2 */
                      if (Core.F_domain_I_restriction(ToRestriction(m.Id())).Id() == C_set.Id()) /* If:11 */{ 
                        sbag = C_set
                        } else {
                        sbag = C_list
                        /* If-11 */} 
                      Result = F_Generate_cast_prefix_class(sbag,s)
                      /* ERROR PROTECTION INSERTED (Result-Result) */
                      if !ErrorIn(Result) {
                      Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{Core.F_domain_I_restriction(ToRestriction(m.Id())).Id(),0}))
                      /* ERROR PROTECTION INSERTED (Result-Result) */
                      if !ErrorIn(Result) {
                      PRINC(".AddFast(")
                      Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{C_any.Id(),0}))
                      /* ERROR PROTECTION INSERTED (Result-Result) */
                      if !ErrorIn(Result) {
                      PRINC(")")
                      F_Generate_cast_post_class(sbag,s)
                      PRINC("")
                      Result = EVOID
                      }}}
                      /* Let-10 */} 
                    /* If!9 */}  else if ((m.Id() == C_Generate__starnth_1_string_star.Value) || 
                      ((m.Id() == C_Generate__starnth_string_star.Value) && 
                          (Optimize.C_compiler.Safety >= 2))) /* If:9 */{ 
                    F_Generate_char_prefix_class(s)
                    Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_string.Id(),0}))
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    PRINC(".At(")
                    Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{C_integer.Id(),0}))
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    PRINC(")")
                    F_Generate_native_post_class(s)
                    PRINC("")
                    Result = EVOID
                    }}
                    /* If!9 */}  else if (m.Selector.Id() == Core.C_identical_ask.Id()) /* If:9 */{ 
                    Result = F_Generate_cast_prefix_class(C_boolean,s)
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    PRINC("MakeBoolean(")
                    Result = Core.F_CALL(Optimize.C_Compile_bool_exp,ARGS(EID{self.Id(),0},EID{CTRUE.Id(),0}))
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    PRINC(")")
                    F_Generate_cast_post_class(C_boolean,s)
                    PRINC("")
                    Result = EVOID
                    }}
                    /* If!9 */}  else if ((p.Id() == Core.C_inlineok_ask.Id()) && 
                      (C_string.Id() == a2.Isa.Id())) /* If:9 */{ 
                    F_Generate_preCore_ask_void()
                    PRINC("F_inlineok_ask_method(")
                    F_Generate_breakline_void()
                    Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_property.Id(),0}))
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    PRINC(",")
                    F_Generate_breakline_void()
                    PRINC("MakeString(")
                    Result = Core.F_CALL(C_print,ARGS(a2.ToEID()))
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    PRINC("))")
                    Result = EVOID
                    }}
                    /* If!9 */}  else if ((m.Id() == C_Generate__starnew_class2_star.Value) && 
                      ((C_class.Id() == a1.Isa.Id()) && 
                        (Optimize.C_compiler.Safety >= 2))) /* If:9 */{ 
                    F_Generate_object_prefix_class(C_any,s)
                    PRINC("new(")
                    F_Generate_go_class_class(ToClass(a1))
                    PRINC(").IsNamed(")
                    Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_class.Id(),0}))
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    PRINC(",")
                    Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{C_symbol.Id(),0}))
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    PRINC(")")
                    F_Generate_object_post_class(C_any,s)
                    PRINC("")
                    Result = EVOID
                    }}
                    } else {
                    Result = c.PrintExternalCall(Language.To_CallMethod(self.Id()),s)
                    /* If-9 */} 
                  }
                  /* If-8 */} 
                }
                /* If-7 */} 
              }
              /* Let-6 */} 
            /* Let-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: inline_exp @ list<type_expression>(go_producer, Call_method2, class) (throw: true) 
func E_Generate_inline_exp_go_producer3 (c EID,self EID,s EID) EID { 
    return /*(sm for inline_exp @ list<type_expression>(go_producer, Call_method2, class)= EID)*/ F_Generate_inline_exp_go_producer3(ToGenerateGoProducer(OBJ(c)),Language.To_CallMethod2(OBJ(self)),ToClass(OBJ(s)) )} 
  
// === functions with three arguments or more
/* {1} OPT.The go function for: inline_exp(c:go_producer,self:Call_method,s:class) [] */
func F_Generate_inline_exp_go_producer4 (c *GenerateGoProducer ,self *Language.CallMethod ,s *ClaireClass ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var m *ClaireMethod   = self.Arg
      /* noccur = 8 */
      /* Let:3 */{ 
        var a1 *ClaireAny   = self.Args.At(1-1)
        /* noccur = 18 */
        /* Let:4 */{ 
          var a2 *ClaireAny   = self.Args.At(2-1)
          /* noccur = 7 */
          /* Let:5 */{ 
            var a3 *ClaireAny   = self.Args.At(3-1)
            /* noccur = 6 */
            var g0321I *ClaireBoolean  
            var g0321I_try03226 EID 
            /* and:6 */{ 
              var v_and6 *ClaireBoolean  
              
              v_and6 = Equal(m.Id(),C_Generate__starnth_equal_list_star.Value)
              if (v_and6 == CFALSE) {g0321I_try03226 = EID{CFALSE.Id(),0}
              } else /* arg:7 */{ 
                v_and6 = F__sup_equal_integer(Optimize.C_compiler.Safety,2)
                if (v_and6 == CFALSE) {g0321I_try03226 = EID{CFALSE.Id(),0}
                } else /* arg:8 */{ 
                  var v_and6_try03239 EID 
                  /* Let:9 */{ 
                    var g0324UU *ClaireClass  
                    /* noccur = 1 */
                    var g0324UU_try032510 EID 
                    g0324UU_try032510 = F_Generate_g_member_any(a1)
                    /* ERROR PROTECTION INSERTED (g0324UU-v_and6_try03239) */
                    if ErrorIn(g0324UU_try032510) {v_and6_try03239 = g0324UU_try032510
                    } else {
                    g0324UU = ToClass(OBJ(g0324UU_try032510))
                    v_and6_try03239 = EID{Core.F__I_equal_any(g0324UU.Id(),C_any.Id()).Id(),0}
                    }
                    /* Let-9 */} 
                  /* ERROR PROTECTION INSERTED (v_and6-g0321I_try03226) */
                  if ErrorIn(v_and6_try03239) {g0321I_try03226 = v_and6_try03239
                  } else {
                  v_and6 = ToBoolean(OBJ(v_and6_try03239))
                  if (v_and6 == CFALSE) {g0321I_try03226 = EID{CFALSE.Id(),0}
                  } else /* arg:9 */{ 
                    v_and6 = Equal(s.Id(),C_void.Id())
                    if (v_and6 == CFALSE) {g0321I_try03226 = EID{CFALSE.Id(),0}
                    } else /* arg:10 */{ 
                      g0321I_try03226 = EID{CTRUE.Id(),0}/* arg-10 */} 
                    /* arg-9 */} 
                  /* arg-8 */} 
                /* arg-7 */} 
              }
              /* and-6 */} 
            /* ERROR PROTECTION INSERTED (g0321I-Result) */
            if ErrorIn(g0321I_try03226) {Result = g0321I_try03226
            } else {
            g0321I = ToBoolean(OBJ(g0321I_try03226))
            if (g0321I == CTRUE) /* If:6 */{ 
              Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_list.Id(),0}))
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              PRINC(".")
              /* Let:7 */{ 
                var g0326UU *ClaireClass  
                /* noccur = 1 */
                var g0326UU_try03278 EID 
                g0326UU_try03278 = F_Generate_g_member_any(a1)
                /* ERROR PROTECTION INSERTED (g0326UU-Result) */
                if ErrorIn(g0326UU_try03278) {Result = g0326UU_try03278
                } else {
                g0326UU = ToClass(OBJ(g0326UU_try03278))
                F_Generate_valuesSlot_class(g0326UU)
                Result = EVOID
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              PRINC("[")
              Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{C_integer.Id(),0}))
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              PRINC("-1]=")
              /* Let:7 */{ 
                var g0328UU *ClaireClass  
                /* noccur = 1 */
                var g0328UU_try03298 EID 
                g0328UU_try03298 = F_Generate_g_member_any(a1)
                /* ERROR PROTECTION INSERTED (g0328UU-Result) */
                if ErrorIn(g0328UU_try03298) {Result = g0328UU_try03298
                } else {
                g0328UU = ToClass(OBJ(g0328UU_try03298))
                Result = Core.F_CALL(C_Generate_g_expression,ARGS(a3.ToEID(),EID{g0328UU.Id(),0}))
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              PRINC("")
              Result = EVOID
              }}}}
              /* If!6 */}  else if ((m.Id() == C_Generate__starnth_put_list_star.Value) || 
                ((m.Id() == C_Generate__starnth_put_array_star.Value) || 
                  ((Optimize.C_compiler.Safety >= 2) && 
                      (m.Id() == C_Generate__starnth_equal_list_star.Value)))) /* If:6 */{ 
              Result = F_Generate_cast_prefix_class(C_any,s)
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_array.Id(),0}))
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              PRINC(".NthPut(")
              Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{C_integer.Id(),0}))
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              PRINC(",")
              Result = Core.F_CALL(C_Generate_g_expression,ARGS(a3.ToEID(),EID{C_any.Id(),0}))
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              PRINC(")")
              F_Generate_cast_post_class(C_any,s)
              PRINC("")
              Result = EVOID
              }}}}
              /* If!6 */}  else if ((m.Id() == C_Generate__starmake_list_star.Value) && 
                (a3 == C_void.Id())) /* If:6 */{ 
              Result = F_Generate_cast_prefix_class(C_list,s)
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              PRINC("CreateList(")
              Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{C_type.Id(),0}))
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              PRINC(",")
              Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_integer.Id(),0}))
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              PRINC(")")
              F_Generate_cast_post_class(C_list,s)
              PRINC("")
              Result = EVOID
              }}}
              } else {
              var g0330I *ClaireBoolean  
              var g0330I_try03317 EID 
              /* and:7 */{ 
                var v_and7 *ClaireBoolean  
                
                v_and7 = Equal(m.Selector.Id(),C_store.Id())
                if (v_and7 == CFALSE) {g0330I_try03317 = EID{CFALSE.Id(),0}
                } else /* arg:8 */{ 
                  var v_and7_try03329 EID 
                  /* or:9 */{ 
                    var v_or9 *ClaireBoolean  
                    
                    var v_or9_try033310 EID 
                    /* Let:10 */{ 
                      var g0334UU *ClaireType  
                      /* noccur = 1 */
                      var g0334UU_try033511 EID 
                      g0334UU_try033511 = Core.F_CALL(Optimize.C_c_type,ARGS(a1.ToEID()))
                      /* ERROR PROTECTION INSERTED (g0334UU-v_or9_try033310) */
                      if ErrorIn(g0334UU_try033511) {v_or9_try033310 = g0334UU_try033511
                      } else {
                      g0334UU = ToType(OBJ(g0334UU_try033511))
                      v_or9_try033310 = EID{g0334UU.Included(ToType(C_list.Id())).Id(),0}
                      }
                      /* Let-10 */} 
                    /* ERROR PROTECTION INSERTED (v_or9-v_and7_try03329) */
                    if ErrorIn(v_or9_try033310) {v_and7_try03329 = v_or9_try033310
                    } else {
                    v_or9 = ToBoolean(OBJ(v_or9_try033310))
                    if (v_or9 == CTRUE) {v_and7_try03329 = EID{CTRUE.Id(),0}
                    } else /* or:10 */{ 
                      var v_or9_try033611 EID 
                      /* Let:11 */{ 
                        var g0337UU *ClaireType  
                        /* noccur = 1 */
                        var g0337UU_try033812 EID 
                        g0337UU_try033812 = Core.F_CALL(Optimize.C_c_type,ARGS(a1.ToEID()))
                        /* ERROR PROTECTION INSERTED (g0337UU-v_or9_try033611) */
                        if ErrorIn(g0337UU_try033812) {v_or9_try033611 = g0337UU_try033812
                        } else {
                        g0337UU = ToType(OBJ(g0337UU_try033812))
                        v_or9_try033611 = EID{g0337UU.Included(ToType(C_array.Id())).Id(),0}
                        }
                        /* Let-11 */} 
                      /* ERROR PROTECTION INSERTED (v_or9-v_and7_try03329) */
                      if ErrorIn(v_or9_try033611) {v_and7_try03329 = v_or9_try033611
                      } else {
                      v_or9 = ToBoolean(OBJ(v_or9_try033611))
                      if (v_or9 == CTRUE) {v_and7_try03329 = EID{CTRUE.Id(),0}
                      } else /* or:11 */{ 
                        v_and7_try03329 = EID{CFALSE.Id(),0}/* org-11 */} 
                      /* org-10 */} 
                    }}
                    /* or-9 */} 
                  /* ERROR PROTECTION INSERTED (v_and7-g0330I_try03317) */
                  if ErrorIn(v_and7_try03329) {g0330I_try03317 = v_and7_try03329
                  } else {
                  v_and7 = ToBoolean(OBJ(v_and7_try03329))
                  if (v_and7 == CFALSE) {g0330I_try03317 = EID{CFALSE.Id(),0}
                  } else /* arg:9 */{ 
                    v_and7 = MakeBoolean(((self.Args.Length() == 4) && 
                        (self.Args.At(4-1) == CTRUE.Id())) || (self.Args.Length() == 3))
                    if (v_and7 == CFALSE) {g0330I_try03317 = EID{CFALSE.Id(),0}
                    } else /* arg:10 */{ 
                      g0330I_try03317 = EID{CTRUE.Id(),0}/* arg-10 */} 
                    /* arg-9 */} 
                  /* arg-8 */} 
                }
                /* and-7 */} 
              /* ERROR PROTECTION INSERTED (g0330I-Result) */
              if ErrorIn(g0330I_try03317) {Result = g0330I_try03317
              } else {
              g0330I = ToBoolean(OBJ(g0330I_try03317))
              if (g0330I == CTRUE) /* If:7 */{ 
                PRINC("F_store_list(")
                Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_list.Id(),0}))
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                PRINC(",")
                Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{C_integer.Id(),0}))
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                PRINC(",")
                Result = Core.F_CALL(C_Generate_g_expression,ARGS(a3.ToEID(),EID{C_any.Id(),0}))
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                PRINC(",CTRUE)")
                Result = EVOID
                }}}
                /* If!7 */}  else if ((m.Selector.Id() == C_add_slot.Id()) && 
                  (C_class.Id() == Core.F_owner_any(F_Generate_getC_any(a1)).Id())) /* If:7 */{ 
                F_Generate_preCore_ask_void()
                PRINC("F_close_slot(")
                Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_class.Id(),0}))
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                PRINC(".AddSlot(")
                Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{C_property.Id(),0}))
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                PRINC(",")
                Result = Core.F_CALL(C_Generate_g_expression,ARGS(a3.ToEID(),EID{C_type.Id(),0}))
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                PRINC(",")
                Result = Core.F_CALL(C_Generate_g_expression,ARGS(self.Args.At(4-1).ToEID(),EID{C_any.Id(),0}))
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                PRINC("))")
                Result = EVOID
                }}}}
                /* If!7 */}  else if (m.Selector.Id() == C_add_method.Id()) /* If:7 */{ 
                if (a1.Isa.IsIn(C_property) == CTRUE) /* If:8 */{ 
                  /* Let:9 */{ 
                    var m *ClaireMethod  
                    /* noccur = 5 */
                    var m_try033910 EID 
                    m_try033910 = F_Generate_retreive_method_any(a1,a2)
                    /* ERROR PROTECTION INSERTED (m-Result) */
                    if ErrorIn(m_try033910) {Result = m_try033910
                    } else {
                    m = ToMethod(OBJ(m_try033910))
                    F_Generate_preCore_ask_void()
                    PRINC("F_attach_method(")
                    Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{a1.Isa.Id(),0}))
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    PRINC(".")
                    F_princ_string(ToString(IfThenElse((a1 == Core.C_self_eval.Id()),
                      MakeString("AddEvalMethod").Id(),
                      MakeString("AddMethod").Id())))
                    PRINC("(")
                    Result = c.Signature_I(F_Generate_full_signature_method(m))
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    PRINC(",")
                    /* Let:10 */{ 
                      var g0340UU int 
                      /* noccur = 1 */
                      var g0340UU_try034111 EID 
                      var g0342I *ClaireBoolean  
                      var g0342I_try034311 EID 
                      g0342I_try034311 = Optimize.F_Compile_can_throw_status_method(m)
                      /* ERROR PROTECTION INSERTED (g0342I-g0340UU_try034111) */
                      if ErrorIn(g0342I_try034311) {g0340UU_try034111 = g0342I_try034311
                      } else {
                      g0342I = ToBoolean(OBJ(g0342I_try034311))
                      if (g0342I == CTRUE) /* If:11 */{ 
                        g0340UU_try034111 = EID{C__INT,IVAL(1)}
                        } else {
                        g0340UU_try034111 = EID{C__INT,IVAL(0)}
                        /* If-11 */} 
                      }
                      /* ERROR PROTECTION INSERTED (g0340UU-Result) */
                      if ErrorIn(g0340UU_try034111) {Result = g0340UU_try034111
                      } else {
                      g0340UU = INT(g0340UU_try034111)
                      F_princ_integer(g0340UU)
                      Result = EVOID
                      }
                      /* Let-10 */} 
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    PRINC(",")
                    Result = F_Generate_goEIDFunction_method(m)
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    if (a1 == Core.C_self_eval.Id()) /* If:10 */{ 
                      F_Generate_goEvalFunction_method(m)
                      /* If-10 */} 
                    PRINC("),MakeString(")
                    /* Let:10 */{ 
                      var g0344UU *ClaireAny  
                      /* noccur = 1 */
                      var g0344UU_try034511 EID 
                      g0344UU_try034511 = Core.F_nth_table1(Optimize.C_Compile_FileOrigin,m.Id())
                      /* ERROR PROTECTION INSERTED (g0344UU-Result) */
                      if ErrorIn(g0344UU_try034511) {Result = g0344UU_try034511
                      } else {
                      g0344UU = ANY(g0344UU_try034511)
                      Result = Core.F_print_any(g0344UU)
                      }
                      /* Let-10 */} 
                    /* ERROR PROTECTION INSERTED (Result-Result) */
                    if !ErrorIn(Result) {
                    PRINC("))")
                    Result = EVOID
                    }}}}}
                    }
                    /* Let-9 */} 
                  } else {
                  PRINC("F_add_method_property(")
                  Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_property.Id(),0}))
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  PRINC(",")
                  Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{C_list.Id(),0}))
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  PRINC(",")
                  Result = Core.F_CALL(C_Generate_g_expression,ARGS(a3.ToEID(),EID{C_type.Id(),0}))
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  PRINC(",")
                  Result = Core.F_CALL(C_Generate_g_expression,ARGS(self.Args.At(4-1).ToEID(),EID{C_integer.Id(),0}))
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  PRINC(",")
                  Result = Core.F_CALL(C_Generate_g_expression,ARGS(self.Args.At(5-1).ToEID(),EID{C_function.Id(),0}))
                  /* ERROR PROTECTION INSERTED (Result-Result) */
                  if !ErrorIn(Result) {
                  PRINC(")")
                  Result = EVOID
                  }}}}}
                  /* If-8 */} 
                } else {
                Result = c.PrintExternalCall(self,s)
                /* If-7 */} 
              }
              /* If-6 */} 
            }
            /* Let-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: inline_exp @ list<type_expression>(go_producer, Call_method, class) (throw: true) 
func E_Generate_inline_exp_go_producer4 (c EID,self EID,s EID) EID { 
    return /*(sm for inline_exp @ list<type_expression>(go_producer, Call_method, class)= EID)*/ F_Generate_inline_exp_go_producer4(ToGenerateGoProducer(OBJ(c)),Language.To_CallMethod(OBJ(self)),ToClass(OBJ(s)) )} 
  
// THIS IS ONE OF THE KEY PATTERNS: calls a method through its compiled function
// the arguments and the result are expected in native format
/* {1} OPT.The go function for: print_external_call(c:go_producer,self:Call_method,s:class) [] */
func (c *GenerateGoProducer ) PrintExternalCall (self *Language.CallMethod ,s *ClaireClass ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var m *ClaireMethod   = self.Arg
      /* noccur = 6 */
      /* Let:3 */{ 
        var l *ClaireList   = self.Args
        /* noccur = 12 */
        /* Let:4 */{ 
          var n int  = 1
          /* noccur = 0 */
          _ = n
          /* Let:5 */{ 
            var _Zsig *ClaireList   = F_Generate_go_signature_method(m)
            /* noccur = 4 */
            /* Let:6 */{ 
              var sm *ClaireAny  
              /* noccur = 3 */
              var sm_try03507 EID 
              sm_try03507 = Core.F_last_list(_Zsig)
              /* ERROR PROTECTION INSERTED (sm-Result) */
              if ErrorIn(sm_try03507) {Result = sm_try03507
              } else {
              sm = ANY(sm_try03507)
              if (l.Length() > 4) /* If:7 */{ 
                Optimize.C_OPT.Level = (Optimize.C_OPT.Level+1)
                /* If-7 */} 
              var g0351I *ClaireBoolean  
              var g0351I_try03527 EID 
              g0351I_try03527 = Optimize.F_Compile_can_throw_status_method(m)
              /* ERROR PROTECTION INSERTED (g0351I-Result) */
              if ErrorIn(g0351I_try03527) {Result = g0351I_try03527
              } else {
              g0351I = ToBoolean(OBJ(g0351I_try03527))
              if (g0351I == CTRUE) /* If:7 */{ 
                sm = Optimize.C_EID.Id()
                Result = sm.ToEID()
                } else {
                Result = EID{CFALSE.Id(),0}
                /* If-7 */} 
              }
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              
              Result = F_Generate_cast_prefix_class(ToClass(sm),s)
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              if (F_Generate_goMethod_ask_any(m.Id()) == CTRUE) /* If:7 */{ 
                Result = F_Generate_external_casted_arg_any(l.At(1-1),ToClass(_Zsig.ValuesO()[1-1]),0,Core.F__sup_integer(l.Length(),4))
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                PRINC(".")
                F_Generate_goMethod_method(m)
                PRINC("(")
                /* Let:8 */{ 
                  var n int  = 2
                  /* noccur = 6 */
                  /* Let:9 */{ 
                    var g0348 int  = l.Length()
                    /* noccur = 1 */
                    Result= EID{CFALSE.Id(),0}
                    for (n <= g0348) /* while:10 */{ 
                      var void_try11 EID 
                      _ = void_try11
                      { 
                      void_try11 = F_Generate_external_casted_arg_any(l.At(n-1),ToClass(_Zsig.ValuesO()[n-1]),(n-1),Core.F__sup_integer(l.Length(),4))
                      /* ERROR PROTECTION INSERTED (void_try11-void_try11) */
                      if ErrorIn(void_try11) {Result = void_try11
                      break
                      } else {
                      n = (n+1)
                      }
                      /* while-10 */} 
                    }
                    /* Let-9 */} 
                  /* Let-8 */} 
                }
                } else {
                Result = F_Generate_goFunction_method(m)
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                PRINC("(")
                Result = EVOID
                }
                {
                if ((l.Length() == 1) && 
                    (Core.F_domain_I_restriction(ToRestriction(m.Id())).Id() == C_void.Id())) /* If:8 */{ 
                  l = CNIL
                  /* If-8 */} 
                /* Let:8 */{ 
                  var n int  = 1
                  /* noccur = 6 */
                  /* Let:9 */{ 
                    var g0349 int  = l.Length()
                    /* noccur = 1 */
                    Result= EID{CFALSE.Id(),0}
                    for (n <= g0349) /* while:10 */{ 
                      var void_try11 EID 
                      _ = void_try11
                      { 
                      void_try11 = F_Generate_external_casted_arg_any(l.At(n-1),ToClass(_Zsig.ValuesO()[n-1]),n,Core.F__sup_integer(l.Length(),4))
                      /* ERROR PROTECTION INSERTED (void_try11-void_try11) */
                      if ErrorIn(void_try11) {Result = void_try11
                      break
                      } else {
                      n = (n+1)
                      }
                      /* while-10 */} 
                    }
                    /* Let-9 */} 
                  /* Let-8 */} 
                }
                /* If-7 */} 
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              PRINC(")")
              if (l.Length() > 4) /* If:7 */{ 
                Optimize.C_OPT.Level = (Optimize.C_OPT.Level-1)
                /* If-7 */} 
              F_Generate_cast_post_class(ToClass(sm),s)
              Result = EVOID
              }}}
              }
              /* Let-6 */} 
            /* Let-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: print_external_call @ go_producer (throw: true) 
func E_Generate_print_external_call_go_producer (c EID,self EID,s EID) EID { 
    return /*(sm for print_external_call @ go_producer= EID)*/ ToGenerateGoProducer(OBJ(c)).PrintExternalCall(Language.To_CallMethod(OBJ(self)),ToClass(OBJ(s)) )} 
  
// prints the n-th arg with a possible cast if necessary since we expect the type t (hence the class class!(t))
// n=0 is a special marker when the arg the receiver x in x.f(....)
// in that case we can do with the static_type because of Go polymorphism
/* {1} OPT.The go function for: external_casted_arg(x:any,s:class,n:integer,nl?:boolean) [] */
func F_Generate_external_casted_arg_any (x *ClaireAny ,s *ClaireClass ,n int,nl_ask *ClaireBoolean ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var st *ClaireClass  
      /* noccur = 2 */
      var st_try03533 EID 
      st_try03533 = Language.F_static_type_any(x)
      /* ERROR PROTECTION INSERTED (st-Result) */
      if ErrorIn(st_try03533) {Result = st_try03533
      } else {
      st = ToClass(OBJ(st_try03533))
      if (n > 1) /* If:3 */{ 
        PRINC(",")
        if (nl_ask == CTRUE) /* If:4 */{ 
          F_Generate_breakline_void()
          /* If-4 */} 
        /* If-3 */} 
      if ((n == 0) && 
          (ToType(st.Id()).Included(ToType(s.Id())) == CTRUE)) /* If:3 */{ 
        Result = Core.F_CALL(C_Generate_g_expression,ARGS(x.ToEID(),EID{st.Id(),0}))
        } else {
        Result = Core.F_CALL(C_Generate_g_expression,ARGS(x.ToEID(),EID{s.Id(),0}))
        /* If-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: external_casted_arg @ any (throw: true) 
func E_Generate_external_casted_arg_any (x EID,s EID,n EID,nl_ask EID) EID { 
    return /*(sm for external_casted_arg @ any= EID)*/ F_Generate_external_casted_arg_any(ANY(x),
      ToClass(OBJ(s)),
      INT(n),
      ToBoolean(OBJ(nl_ask)) )} 
  
//
//**********************************************************************
//*          Part 4: expression for structures                       *
//**********************************************************************
// this is an attempt to get rid of useless parenthesis without creating ambuiguous situations
// bounded_expression(x,loop) adds wrapping ( ) if needed     ==     bounded expression :)
// here we assume that native is needed
/* {1} OPT.The go function for: bounded_expression(self:any,s:class) [] */
func F_Generate_bounded_expression_any (self *ClaireAny ,s *ClaireClass ) EID { 
    var Result EID 
    if (self.Isa.IsIn(Language.C_Assign) == CTRUE) /* If:2 */{ 
      /* Let:3 */{ 
        var g0354 *Language.Assign   = Language.To_Assign(self)
        /* noccur = 1 */
        PRINC("(")
        Result = F_Generate_g_expression_any(g0354.Id(),s)
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC(")")
        Result = EVOID
        }
        /* Let-3 */} 
      /* If!2 */}  else if (C_integer.Id() == self.Isa.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var g0355 int  = ToInteger(self).Value
        /* noccur = 3 */
        if (g0355 < 0) /* If:4 */{ 
          PRINC("(")
          Result = F_Generate_g_expression_integer(g0355,s)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(")")
          Result = EVOID
          }
          } else {
          Result = F_Generate_g_expression_integer(g0355,s)
          /* If-4 */} 
        /* Let-3 */} 
      /* If!2 */}  else if (C_float.Id() == self.Isa.Id()) /* If:2 */{ 
      /* Let:3 */{ 
        var g0356 float64  = ToFloat(self).Value
        /* noccur = 3 */
        if (g0356 < 0) /* If:4 */{ 
          PRINC("(")
          Result = F_Generate_g_expression_float(g0356,s)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(")")
          Result = EVOID
          }
          } else {
          Result = F_Generate_g_expression_float(g0356,s)
          /* If-4 */} 
        /* Let-3 */} 
      } else {
      Result = Core.F_CALL(C_Generate_g_expression,ARGS(self.ToEID(),EID{s.Id(),0}))
      /* If-2 */} 
    return Result} 
  
// The EID go function for: bounded_expression @ any (throw: true) 
func E_Generate_bounded_expression_any (self EID,s EID) EID { 
    return /*(sm for bounded_expression @ any= EID)*/ F_Generate_bounded_expression_any(ANY(self),ToClass(OBJ(s)) )} 
  
// if can be represented by an expression if the two arguments are constants (evaluation does not cost)
/* {1} OPT.The go function for: g_expression(self:If,s:class) [] */
func F_Generate_g_expression_If (self *Language.If ,s *ClaireClass ) EID { 
    var Result EID 
    Core.F_tformat_string(MakeString("g_exp @ IF, s = ~S \n"),0,MakeConstantList(s.Id()))
    F_Generate_object_prefix_class(C_any,s)
    PRINC("IfThenElse(")
    Result = Core.F_CALL(Optimize.C_Compile_bool_exp,ARGS(self.Test.ToEID(),EID{CTRUE.Id(),0}))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(",")
    Result = EVOID
    }
    {
    Optimize.C_OPT.Level = (Optimize.C_OPT.Level+1)
    F_Generate_breakline_void()
    Result = Core.F_CALL(C_Generate_g_expression,ARGS(self.Arg.ToEID(),EID{C_any.Id(),0}))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(",")
    Result = EVOID
    }
    {
    F_Generate_breakline_void()
    Result = Core.F_CALL(C_Generate_g_expression,ARGS(self.Other.ToEID(),EID{C_any.Id(),0}))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(")")
    Result = EVOID
    }
    {
    F_Generate_object_post_class(C_any,s)
    /* update:2 */{ 
      var va_arg1 *Optimize.OptimizeMetaOPT  
      var va_arg2 int 
      va_arg1 = Optimize.C_OPT
      va_arg2 = (Optimize.C_OPT.Level-1)
      /* ---------- now we compile update Compile/level(va_arg1) := va_arg2 ------- */
      va_arg1.Level = va_arg2
      Result = EID{C__INT,IVAL(va_arg2)}
      /* update-2 */} 
    }}}
    return Result} 
  
// The EID go function for: g_expression @ If (throw: true) 
func E_Generate_g_expression_If (self EID,s EID) EID { 
    return /*(sm for g_expression @ If= EID)*/ F_Generate_g_expression_If(Language.To_If(OBJ(self)),ToClass(OBJ(s)) )} 
  
// a conjunction is also a C expression
// note that go requires && before the line break hence the more complex code
/* {1} OPT.The go function for: g_expression(self:And,s:class) [] */
func F_Generate_g_expression_And (self *Language.And ,s *ClaireClass ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var b *ClaireBoolean   = Core.F__sup_integer(self.Args.Length(),5)
      /* noccur = 1 */
      /* Let:3 */{ 
        var n int  = self.Args.Length()
        /* noccur = 2 */
        F_Generate_object_prefix_class(C_boolean,s)
        PRINC("MakeBoolean(")
        /* Let:4 */{ 
          var i int  = 1
          /* noccur = 5 */
          /* Let:5 */{ 
            var g0360 int  = n
            /* noccur = 1 */
            Result= EID{CFALSE.Id(),0}
            for (i <= g0360) /* while:6 */{ 
              var void_try7 EID 
              _ = void_try7
              { 
              /* Let:7 */{ 
                var x *ClaireAny   = self.Args.At(i-1)
                /* noccur = 1 */
                void_try7 = Core.F_CALL(Optimize.C_Compile_bool_exp,ARGS(x.ToEID(),EID{CTRUE.Id(),0}))
                /* ERROR PROTECTION INSERTED (void_try7-void_try7) */
                if ErrorIn(void_try7) {Result = void_try7
                break
                } else {
                if (i < n) /* If:8 */{ 
                  PRINC(" && ")
                  if (b == CTRUE) /* If:9 */{ 
                    F_Generate_breakline_void()
                    /* If-9 */} 
                  PRINC("")
                  void_try7 = EVOID
                  } else {
                  void_try7 = EID{CFALSE.Id(),0}
                  /* If-8 */} 
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (void_try7-void_try7) */
              if ErrorIn(void_try7) {Result = void_try7
              break
              } else {
              i = (i+1)
              }
              /* while-6 */} 
            }
            /* Let-5 */} 
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC(")")
        F_Generate_object_post_class(C_boolean,s)
        Result = EVOID
        }
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: g_expression @ And (throw: true) 
func E_Generate_g_expression_And (self EID,s EID) EID { 
    return /*(sm for g_expression @ And= EID)*/ F_Generate_g_expression_And(Language.To_And(OBJ(self)),ToClass(OBJ(s)) )} 
  
// same thing for a disjunction
/* {1} OPT.The go function for: g_expression(self:Or,s:class) [] */
func F_Generate_g_expression_Or (self *Language.Or ,s *ClaireClass ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var b *ClaireBoolean   = Core.F__sup_integer(self.Args.Length(),5)
      /* noccur = 1 */
      /* Let:3 */{ 
        var n int  = self.Args.Length()
        /* noccur = 2 */
        F_Generate_object_prefix_class(C_boolean,s)
        PRINC("MakeBoolean(")
        /* Let:4 */{ 
          var i int  = 1
          /* noccur = 5 */
          /* Let:5 */{ 
            var g0361 int  = n
            /* noccur = 1 */
            Result= EID{CFALSE.Id(),0}
            for (i <= g0361) /* while:6 */{ 
              var void_try7 EID 
              _ = void_try7
              { 
              /* Let:7 */{ 
                var x *ClaireAny   = self.Args.At(i-1)
                /* noccur = 1 */
                void_try7 = Core.F_CALL(Optimize.C_Compile_bool_exp,ARGS(x.ToEID(),EID{CTRUE.Id(),0}))
                /* ERROR PROTECTION INSERTED (void_try7-void_try7) */
                if ErrorIn(void_try7) {Result = void_try7
                break
                } else {
                if (i < n) /* If:8 */{ 
                  PRINC(" || ")
                  if (b == CTRUE) /* If:9 */{ 
                    F_Generate_breakline_void()
                    /* If-9 */} 
                  PRINC("")
                  void_try7 = EVOID
                  } else {
                  void_try7 = EID{CFALSE.Id(),0}
                  /* If-8 */} 
                }
                /* Let-7 */} 
              /* ERROR PROTECTION INSERTED (void_try7-void_try7) */
              if ErrorIn(void_try7) {Result = void_try7
              break
              } else {
              i = (i+1)
              }
              /* while-6 */} 
            }
            /* Let-5 */} 
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC(")")
        F_Generate_object_post_class(C_boolean,s)
        Result = EVOID
        }
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: g_expression @ Or (throw: true) 
func E_Generate_g_expression_Or (self EID,s EID) EID { 
    return /*(sm for g_expression @ Or= EID)*/ F_Generate_g_expression_Or(Language.To_Or(OBJ(self)),ToClass(OBJ(s)) )} 
  
// to_CL(x) produces a CLAIRE id from an external representation
// [g_expression(self:Generate/to_CL,s:class) : void
//  -> //[5] toCL -> ~S:~S // self.arg, owner(self.arg),
//    g_expression(self.arg, s)]
// to_C(x) produces an external representation from a CLAIRE id
// g_expression(self:Generate/to_C,s:class) : void
// -> g_expression(self.arg, s)
// C_cast(x) produces a cast for go  => unclear if it is still needed
/* {1} OPT.The go function for: g_expression(self:Compile/C_cast,s:class) [] */
func F_Generate_g_expression_C_cast (self *Optimize.CompileCCast ,s *ClaireClass ) EID { 
    var Result EID 
    Result = Core.F_CALL(C_Generate_g_expression,ARGS(self.Arg.ToEID(),EID{s.Id(),0}))
    return Result} 
  
// The EID go function for: g_expression @ Compile/C_cast (throw: true) 
func E_Generate_g_expression_C_cast (self EID,s EID) EID { 
    return /*(sm for g_expression @ Compile/C_cast= EID)*/ F_Generate_g_expression_C_cast(Optimize.To_CompileCCast(OBJ(self)),ToClass(OBJ(s)) )} 
  
// reads a slot : more complex that it looks
// when the test is on, we produce Known(p,x->p).To 
// because slots can be native, we need the generic pre/post to convert to the proper slot
/* {1} OPT.The go function for: g_expression(self:Call_slot,s:class) [] */
func F_Generate_g_expression_Call_slot (self *Language.CallSlot ,s *ClaireClass ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var sc *ClaireClass   = self.Selector.Range.Class_I()
      /* noccur = 2 */
      /* Let:3 */{ 
        var dc *ClaireClass  
        /* noccur = 2 */
        var dc_try03624 EID 
        dc_try03624 = Language.F_static_type_any(self.Arg)
        /* ERROR PROTECTION INSERTED (dc-Result) */
        if ErrorIn(dc_try03624) {Result = dc_try03624
        } else {
        dc = ToClass(OBJ(dc_try03624))
        /* Let:4 */{ 
          var s2 *ClaireClass  
          /* noccur = 1 */
          if (ToType(dc.Id()).Included(ToType(Core.F_domain_I_restriction(ToRestriction(self.Selector.Id())).Id())) == CTRUE) /* If:5 */{ 
            s2 = dc
            } else {
            s2 = Core.F_domain_I_restriction(ToRestriction(self.Selector.Id())).Class_I()
            /* If-5 */} 
          Result = F_Generate_cast_prefix_class(sc,s)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          if ((self.Test.Id() != CNULL) && 
              (self.Test == CTRUE)) /* If:5 */{ 
            F_Generate_cast_class_class(self.Selector.Range.Class_I())
            PRINC("(KNOWN(")
            Result = F_Generate_g_expression_thing(ToThing(self.Selector.Selector.Id()),C_object)
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC(",")
            Result = EVOID
            }
            } else {
            Result = EID{CFALSE.Id(),0}
            /* If-5 */} 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).CMember(self.Arg,s2,self.Selector.Selector)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          if ((self.Test.Id() != CNULL) && 
              (self.Test == CTRUE)) /* If:5 */{ 
            PRINC("))")
            /* If-5 */} 
          F_Generate_cast_post_class(sc,s)
          Result = EVOID
          }}}
          /* Let-4 */} 
        }
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: g_expression @ Call_slot (throw: true) 
func E_Generate_g_expression_Call_slot (self EID,s EID) EID { 
    return /*(sm for g_expression @ Call_slot= EID)*/ F_Generate_g_expression_Call_slot(Language.To_CallSlot(OBJ(self)),ToClass(OBJ(s)) )} 
  
// reads an (integer) table  = WARNING - this will change in the future when tables are implemented with dictionaries
/* {1} OPT.The go function for: g_expression(self:Call_table,s:class) [] */
func F_Generate_g_expression_Call_table (self *Language.CallTable ,s *ClaireClass ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var a *ClaireTable   = self.Selector
      /* noccur = 3 */
      /* Let:3 */{ 
        var p *ClaireAny   = a.Params
        /* noccur = 4 */
        /* Let:4 */{ 
          var l *ClaireAny   = self.Arg
          /* noccur = 3 */
          F_Generate_object_prefix_class(C_any,s)
          if (self.Test == CTRUE) /* If:5 */{ 
            PRINC("KNOWN(")
            Result = F_Generate_g_expression_thing(ToThing(a.Id()),C_object)
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            PRINC(",")
            Result = EVOID
            }
            } else {
            Result = EID{CFALSE.Id(),0}
            /* If-5 */} 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          Result = F_Generate_g_expression_thing(ToThing(a.Id()),C_object)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(".Graph[")
          if (C_integer.Id() == p.Isa.Id()) /* If:5 */{ 
            /* Let:6 */{ 
              var g0363 int  = ToInteger(p).Value
              /* noccur = 1 */
              Result = Core.F_CALL(C_Generate_g_expression,ARGS(l.ToEID(),EID{C_integer.Id(),0}))
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              PRINC(" - ")
              F_princ_integer(g0363)
              PRINC("")
              Result = EVOID
              }
              /* Let-6 */} 
            /* If!5 */}  else if (p.Isa.IsIn(C_list) == CTRUE) /* If:5 */{ 
            /* Let:6 */{ 
              var g0364 *ClaireList   = ToList(p)
              /* noccur = 2 */
              Result = Core.F_CALL(C_Generate_g_expression,ARGS(ToList(OBJ(Core.F_CALL(C_args,ARGS(l.ToEID())))).At(1-1).ToEID(),EID{C_integer.Id(),0}))
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              PRINC(" * ")
              Result = Core.F_CALL(C_princ,ARGS(g0364.At(1-1).ToEID()))
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              PRINC(" + ")
              Result = Core.F_CALL(C_Generate_g_expression,ARGS(ToList(OBJ(Core.F_CALL(C_args,ARGS(l.ToEID())))).At(2-1).ToEID(),EID{C_integer.Id(),0}))
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              PRINC(" - ")
              Result = Core.F_CALL(C_princ,ARGS(g0364.At(2-1).ToEID()))
              /* ERROR PROTECTION INSERTED (Result-Result) */
              if !ErrorIn(Result) {
              PRINC("")
              Result = EVOID
              }}}}
              /* Let-6 */} 
            } else {
            Result = EID{CFALSE.Id(),0}
            /* If-5 */} 
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC("]")
          Result = EVOID
          }}
          {
          if (self.Test == CTRUE) /* If:5 */{ 
            PRINC(")")
            /* If-5 */} 
          F_Generate_object_post_class(C_any,s)
          Result = EVOID
          }}
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: g_expression @ Call_table (throw: true) 
func E_Generate_g_expression_Call_table (self EID,s EID) EID { 
    return /*(sm for g_expression @ Call_table= EID)*/ F_Generate_g_expression_Call_table(Language.To_CallTable(OBJ(self)),ToClass(OBJ(s)) )} 
  
// reads an array - remember that in CLAIRE 4, arrays are nothing but fixed size lists (with 3 sorts)
/* {1} OPT.The go function for: g_expression(self:Call_array,s:class) [] */
func F_Generate_g_expression_Call_array (self *Language.CallArray ,s *ClaireClass ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var sa *ClaireClass  
      /* noccur = 2 */
      var sa_try03653 EID 
      /* Let:3 */{ 
        var g0366UU *ClaireType  
        /* noccur = 1 */
        var g0366UU_try03674 EID 
        /* Let:4 */{ 
          var g0368UU *ClaireType  
          /* noccur = 1 */
          var g0368UU_try03695 EID 
          g0368UU_try03695 = Core.F_CALL(Optimize.C_c_type,ARGS(self.Selector.ToEID()))
          /* ERROR PROTECTION INSERTED (g0368UU-g0366UU_try03674) */
          if ErrorIn(g0368UU_try03695) {g0366UU_try03674 = g0368UU_try03695
          } else {
          g0368UU = ToType(OBJ(g0368UU_try03695))
          g0366UU_try03674 = EID{Core.F_member_type(g0368UU).Id(),0}
          }
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (g0366UU-sa_try03653) */
        if ErrorIn(g0366UU_try03674) {sa_try03653 = g0366UU_try03674
        } else {
        g0366UU = ToType(OBJ(g0366UU_try03674))
        sa_try03653 = EID{F_Generate_type_sort_type(g0366UU).Id(),0}
        }
        /* Let-3 */} 
      /* ERROR PROTECTION INSERTED (sa-Result) */
      if ErrorIn(sa_try03653) {Result = sa_try03653
      } else {
      sa = ToClass(OBJ(sa_try03653))
      /* Let:3 */{ 
        var sm *ClaireClass  
        /* noccur = 2 */
        var sm_try03704 EID 
        sm_try03704 = F_Generate_g_member_any(self.Selector)
        /* ERROR PROTECTION INSERTED (sm-Result) */
        if ErrorIn(sm_try03704) {Result = sm_try03704
        } else {
        sm = ToClass(OBJ(sm_try03704))
        Result = F_Generate_cast_prefix_class(sa,s)
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        if (sm.Id() != C_any.Id()) /* If:4 */{ 
          Result = Core.F_CALL(C_Generate_g_expression,ARGS(self.Selector.ToEID(),EID{C_list.Id(),0}))
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(".")
          F_Generate_valuesSlot_class(sm)
          PRINC("[")
          Result = Core.F_CALL(C_Generate_g_expression,ARGS(self.Arg.ToEID(),EID{C_integer.Id(),0}))
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(" - 1]")
          Result = EVOID
          }}
          } else {
          Result = Core.F_CALL(C_Generate_g_expression,ARGS(self.Selector.ToEID(),EID{C_list.Id(),0}))
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(".At(")
          Result = Core.F_CALL(C_Generate_g_expression,ARGS(self.Arg.ToEID(),EID{C_integer.Id(),0}))
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(" - 1)")
          Result = EVOID
          }}
          /* If-4 */} 
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        F_Generate_cast_post_class(sa,s)
        Result = EVOID
        }}
        }
        /* Let-3 */} 
      }
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: g_expression @ Call_array (throw: true) 
func E_Generate_g_expression_Call_array (self EID,s EID) EID { 
    return /*(sm for g_expression @ Call_array= EID)*/ F_Generate_g_expression_Call_array(Language.To_CallArray(OBJ(self)),ToClass(OBJ(s)) )} 
  
//**********************************************************************
//*          Part 5: the logical expression compilation                *
//**********************************************************************
// bool_exp(x,pos?) returns a native boolean go expression, assumes that g_func(x) !
// bool_expression(x) could be g_expression(x,boolean)
// however, boolean are not native in CLAIRE4 () to avoid conversions
// note : we drop bool_exp? and bool_exp!
// this is the boolean compiler. An automatic computation of negation is
// included. The flag pos? tells if the assertion is positive. When a
// negation occurs, we simply change the flag. At the end of compiling,
// the flag is used to generate == or != according to this method:
// generate the = or /=
/* {1} OPT.The go function for: sign_equal(self:boolean) [] */
func F_Generate_sign_equal_boolean (self *ClaireBoolean )  { 
    // procedure body with s =  
if (self == CTRUE) /* If:2 */{ 
      PRINC("==")
      } else {
      PRINC("!=")
      /* If-2 */} 
    } 
  
// The EID go function for: sign_equal @ boolean (throw: false) 
func E_Generate_sign_equal_boolean (self EID) EID { 
    /*(sm for sign_equal @ boolean= void)*/ F_Generate_sign_equal_boolean(ToBoolean(OBJ(self)) )
    return EVOID} 
  
// generate a conjunction/disjunction
/* {1} OPT.The go function for: sign_or(self:boolean) [] */
func F_Generate_sign_or_boolean (self *ClaireBoolean )  { 
    // procedure body with s =  
if (self == CTRUE) /* If:2 */{ 
      PRINC("||")
      } else {
      PRINC("&&")
      /* If-2 */} 
    } 
  
// The EID go function for: sign_or @ boolean (throw: false) 
func E_Generate_sign_or_boolean (self EID) EID { 
    /*(sm for sign_or @ boolean= void)*/ F_Generate_sign_or_boolean(ToBoolean(OBJ(self)) )
    return EVOID} 
  
// default solution
/* {1} OPT.The go function for: Compile/bool_exp(self:any,pos?:boolean) [] */
func F_Compile_bool_exp_any (self *ClaireAny ,pos_ask *ClaireBoolean ) EID { 
    var Result EID 
    PRINC("(")
    Result = Core.F_CALL(C_Generate_g_expression,ARGS(self.ToEID(),EID{C_boolean.Id(),0}))
    /* ERROR PROTECTION INSERTED (Result-Result) */
    if !ErrorIn(Result) {
    PRINC(" ")
    F_Generate_sign_equal_boolean(pos_ask)
    PRINC(" CTRUE)")
    Result = EVOID
    }
    return Result} 
  
// The EID go function for: Compile/bool_exp @ any (throw: true) 
func E_Compile_bool_exp_any (self EID,pos_ask EID) EID { 
    return /*(sm for Compile/bool_exp @ any= EID)*/ F_Compile_bool_exp_any(ANY(self),ToBoolean(OBJ(pos_ask)) )} 
  
// if we have a CL, we know that the self.arg is of type boolean
// [bool_exp(self:Generate/to_CL,pos?:boolean) : void
//  -> bool_exp(self.arg,pos?) ]
// If is supported with IfThenElse (means that all terms will be evaluated),
/* {1} OPT.The go function for: Compile/bool_exp(self:If,pos?:boolean) [] */
func F_Compile_bool_exp_If (self *Language.If ,pos_ask *ClaireBoolean ) EID { 
    var Result EID 
    if (F_boolean_I_any(self.Other) == CTRUE) /* If:2 */{ 
      PRINC("(")
      Result = Core.F_CALL(Optimize.C_Compile_bool_exp,ARGS(self.Test.ToEID(),EID{CTRUE.Id(),0}))
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(" ? ")
      Result = Core.F_CALL(Optimize.C_Compile_bool_exp,ARGS(self.Arg.ToEID(),EID{pos_ask.Id(),0}))
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(" : ")
      Result = Core.F_CALL(Optimize.C_Compile_bool_exp,ARGS(self.Other.ToEID(),EID{pos_ask.Id(),0}))
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(")")
      Result = EVOID
      }}}
      } else {
      PRINC("(")
      Result = Core.F_CALL(Optimize.C_Compile_bool_exp,ARGS(self.Test.ToEID(),EID{pos_ask.Id(),0}))
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(" ")
      F_Generate_sign_or_boolean(pos_ask.Not)
      PRINC(" ")
      Result = Core.F_CALL(Optimize.C_Compile_bool_exp,ARGS(self.Arg.ToEID(),EID{pos_ask.Id(),0}))
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(")")
      Result = EVOID
      }}
      /* If-2 */} 
    return Result} 
  
// The EID go function for: Compile/bool_exp @ If (throw: true) 
func E_Compile_bool_exp_If (self EID,pos_ask EID) EID { 
    return /*(sm for Compile/bool_exp @ If= EID)*/ F_Compile_bool_exp_If(Language.To_If(OBJ(self)),ToBoolean(OBJ(pos_ask)) )} 
  
// for a AND, we can used the && C operation
/* {1} OPT.The go function for: Compile/bool_exp(self:And,pos?:boolean) [] */
func F_Compile_bool_exp_And (self *Language.And ,pos_ask *ClaireBoolean ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var l *ClaireList   = self.Args
      /* noccur = 2 */
      /* Let:3 */{ 
        var m int  = l.Length()
        /* noccur = 2 */
        /* Let:4 */{ 
          var n int  = 0
          /* noccur = 3 */
          /* Let:5 */{ 
            var _Zl int  = Optimize.C_OPT.Level
            /* noccur = 1 */
            Optimize.C_OPT.Level = (Optimize.C_OPT.Level+1)
            /* For:6 */{ 
              var x *ClaireAny  
              _ = x
              Result= EID{CFALSE.Id(),0}
              var x_support *ClaireList  
              x_support = l
              x_len := x_support.Length()
              for i_it := 0; i_it < x_len; i_it++ { 
                x = x_support.At(i_it)
                var void_try8 EID 
                _ = void_try8
                { 
                n = (n+1)
                if (n == m) /* If:8 */{ 
                  void_try8 = Core.F_CALL(Optimize.C_Compile_bool_exp,ARGS(x.ToEID(),EID{pos_ask.Id(),0}))
                  } else {
                  PRINC("(")
                  void_try8 = Core.F_CALL(Optimize.C_Compile_bool_exp,ARGS(x.ToEID(),EID{pos_ask.Id(),0}))
                  /* ERROR PROTECTION INSERTED (void_try8-void_try8) */
                  if ErrorIn(void_try8) {Result = void_try8
                  break
                  } else {
                  PRINC(" ")
                  F_Generate_sign_or_boolean(pos_ask.Not)
                  PRINC(" ")
                  void_try8 = EVOID
                  }
                  {
                  Optimize.C_OPT.Level = (Optimize.C_OPT.Level+1)
                  void_try8 = F_Generate_breakline_void().ToEID()
                  }
                  /* If-8 */} 
                /* ERROR PROTECTION INSERTED (void_try8-void_try8) */
                if ErrorIn(void_try8) {Result = void_try8
                break
                } else {
                }
                }
                /* loop-7 */} 
              /* For-6 */} 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            /* Let:6 */{ 
              var x int  = 2
              /* noccur = 3 */
              /* Let:7 */{ 
                var g0373 int  = m
                /* noccur = 1 */
                for (x <= g0373) /* while:8 */{ 
                  PRINC(")")
                  x = (x+1)
                  /* while-8 */} 
                /* Let-7 */} 
              /* Let-6 */} 
            /* update:6 */{ 
              var va_arg1 *Optimize.OptimizeMetaOPT  
              var va_arg2 int 
              va_arg1 = Optimize.C_OPT
              va_arg2 = _Zl
              /* ---------- now we compile update Compile/level(va_arg1) := va_arg2 ------- */
              va_arg1.Level = va_arg2
              Result = EID{C__INT,IVAL(va_arg2)}
              /* update-6 */} 
            }
            /* Let-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: Compile/bool_exp @ And (throw: true) 
func E_Compile_bool_exp_And (self EID,pos_ask EID) EID { 
    return /*(sm for Compile/bool_exp @ And= EID)*/ F_Compile_bool_exp_And(Language.To_And(OBJ(self)),ToBoolean(OBJ(pos_ask)) )} 
  
// idem for OR: we use ||
/* {1} OPT.The go function for: Compile/bool_exp(self:Or,pos?:boolean) [] */
func F_Compile_bool_exp_Or (self *Language.Or ,pos_ask *ClaireBoolean ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var l *ClaireList   = self.Args
      /* noccur = 2 */
      /* Let:3 */{ 
        var m int  = l.Length()
        /* noccur = 2 */
        /* Let:4 */{ 
          var n int  = 0
          /* noccur = 3 */
          /* Let:5 */{ 
            var _Zl int  = Optimize.C_OPT.Level
            /* noccur = 1 */
            Optimize.C_OPT.Level = (Optimize.C_OPT.Level+1)
            /* For:6 */{ 
              var x *ClaireAny  
              _ = x
              Result= EID{CFALSE.Id(),0}
              var x_support *ClaireList  
              x_support = l
              x_len := x_support.Length()
              for i_it := 0; i_it < x_len; i_it++ { 
                x = x_support.At(i_it)
                var void_try8 EID 
                _ = void_try8
                { 
                n = (n+1)
                if (n == m) /* If:8 */{ 
                  void_try8 = Core.F_CALL(Optimize.C_Compile_bool_exp,ARGS(x.ToEID(),EID{pos_ask.Id(),0}))
                  } else {
                  PRINC("(")
                  void_try8 = Core.F_CALL(Optimize.C_Compile_bool_exp,ARGS(x.ToEID(),EID{pos_ask.Id(),0}))
                  /* ERROR PROTECTION INSERTED (void_try8-void_try8) */
                  if ErrorIn(void_try8) {Result = void_try8
                  break
                  } else {
                  PRINC(" ")
                  F_Generate_sign_or_boolean(pos_ask)
                  PRINC(" ")
                  void_try8 = EVOID
                  }
                  {
                  Optimize.C_OPT.Level = (Optimize.C_OPT.Level+1)
                  void_try8 = F_Generate_breakline_void().ToEID()
                  }
                  /* If-8 */} 
                /* ERROR PROTECTION INSERTED (void_try8-void_try8) */
                if ErrorIn(void_try8) {Result = void_try8
                break
                } else {
                }
                }
                /* loop-7 */} 
              /* For-6 */} 
            /* ERROR PROTECTION INSERTED (Result-Result) */
            if !ErrorIn(Result) {
            /* Let:6 */{ 
              var x int  = 2
              /* noccur = 3 */
              /* Let:7 */{ 
                var g0376 int  = m
                /* noccur = 1 */
                for (x <= g0376) /* while:8 */{ 
                  PRINC(")")
                  x = (x+1)
                  /* while-8 */} 
                /* Let-7 */} 
              /* Let-6 */} 
            /* update:6 */{ 
              var va_arg1 *Optimize.OptimizeMetaOPT  
              var va_arg2 int 
              va_arg1 = Optimize.C_OPT
              va_arg2 = _Zl
              /* ---------- now we compile update Compile/level(va_arg1) := va_arg2 ------- */
              va_arg1.Level = va_arg2
              Result = EID{C__INT,IVAL(va_arg2)}
              /* update-6 */} 
            }
            /* Let-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: Compile/bool_exp @ Or (throw: true) 
func E_Compile_bool_exp_Or (self EID,pos_ask EID) EID { 
    return /*(sm for Compile/bool_exp @ Or= EID)*/ F_Compile_bool_exp_Or(Language.To_Or(OBJ(self)),ToBoolean(OBJ(pos_ask)) )} 
  
// membership
/* {1} OPT.The go function for: Compile/bool_exp(self:Call,pos?:boolean) [] */
func F_Compile_bool_exp_Call (self *Language.Call ,pos_ask *ClaireBoolean ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var p *ClaireProperty   = self.Selector
      /* noccur = 1 */
      if (p.Id() == C__Z.Id()) /* If:3 */{ 
        PRINC("(")
        Result = F_Generate_belong_exp_any(self.Args.At(1-1),self.Args.At(2-1),C_boolean)
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC(" ")
        F_Generate_sign_equal_boolean(pos_ask)
        PRINC(" CTRUE)")
        Result = EVOID
        }
        } else {
        Result = Core.F_SUPER(Optimize.C_Compile_bool_exp, C_any, ARGS(EID{self.Id(),0},EID{pos_ask.Id(),0}))
        /* If-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: Compile/bool_exp @ Call (throw: true) 
func E_Compile_bool_exp_Call (self EID,pos_ask EID) EID { 
    return /*(sm for Compile/bool_exp @ Call= EID)*/ F_Compile_bool_exp_Call(Language.To_Call(OBJ(self)),ToBoolean(OBJ(pos_ask)) )} 
  
// compile (a % ..), s is always a boolean but for EID mode
// the notOpt() test in gostat.cl ensures that the first three cases are seen as not-throw (not EID)
// however this fragment may be called to return an EID hence the global wrap with prefix/post
/* {1} OPT.The go function for: belong_exp(a1:any,a2:any,s:class) [] */
func F_Generate_belong_exp_any (a1 *ClaireAny ,a2 *ClaireAny ,s *ClaireClass ) EID { 
    var Result EID 
    var g0377I *ClaireBoolean  
    var g0377I_try03782 EID 
    /* Let:2 */{ 
      var g0379UU *ClaireClass  
      /* noccur = 1 */
      var g0379UU_try03803 EID 
      g0379UU_try03803 = Language.F_static_type_any(a2)
      /* ERROR PROTECTION INSERTED (g0379UU-g0377I_try03782) */
      if ErrorIn(g0379UU_try03803) {g0377I_try03782 = g0379UU_try03803
      } else {
      g0379UU = ToClass(OBJ(g0379UU_try03803))
      g0377I_try03782 = EID{ToType(g0379UU.Id()).Included(ToType(C_type.Id())).Id(),0}
      }
      /* Let-2 */} 
    /* ERROR PROTECTION INSERTED (g0377I-Result) */
    if ErrorIn(g0377I_try03782) {Result = g0377I_try03782
    } else {
    g0377I = ToBoolean(OBJ(g0377I_try03782))
    if (g0377I == CTRUE) /* If:2 */{ 
      Result = F_Generate_cast_prefix_class(C_boolean,s)
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{C_type.Id(),0}))
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(".Contains(")
      Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_any.Id(),0}))
      /* ERROR PROTECTION INSERTED (Result-Result) */
      if !ErrorIn(Result) {
      PRINC(")")
      F_Generate_cast_post_class(C_boolean,s)
      PRINC("")
      Result = EVOID
      }}}
      } else {
      var g0381I *ClaireBoolean  
      var g0381I_try03823 EID 
      /* and:3 */{ 
        var v_and3 *ClaireBoolean  
        
        var v_and3_try03834 EID 
        /* Let:4 */{ 
          var g0384UU *ClaireClass  
          /* noccur = 1 */
          var g0384UU_try03855 EID 
          g0384UU_try03855 = Language.F_static_type_any(a2)
          /* ERROR PROTECTION INSERTED (g0384UU-v_and3_try03834) */
          if ErrorIn(g0384UU_try03855) {v_and3_try03834 = g0384UU_try03855
          } else {
          g0384UU = ToClass(OBJ(g0384UU_try03855))
          v_and3_try03834 = EID{ToType(g0384UU.Id()).Included(ToType(C_integer.Id())).Id(),0}
          }
          /* Let-4 */} 
        /* ERROR PROTECTION INSERTED (v_and3-g0381I_try03823) */
        if ErrorIn(v_and3_try03834) {g0381I_try03823 = v_and3_try03834
        } else {
        v_and3 = ToBoolean(OBJ(v_and3_try03834))
        if (v_and3 == CFALSE) {g0381I_try03823 = EID{CFALSE.Id(),0}
        } else /* arg:4 */{ 
          var v_and3_try03865 EID 
          /* Let:5 */{ 
            var g0387UU *ClaireClass  
            /* noccur = 1 */
            var g0387UU_try03886 EID 
            g0387UU_try03886 = Language.F_static_type_any(a1)
            /* ERROR PROTECTION INSERTED (g0387UU-v_and3_try03865) */
            if ErrorIn(g0387UU_try03886) {v_and3_try03865 = g0387UU_try03886
            } else {
            g0387UU = ToClass(OBJ(g0387UU_try03886))
            v_and3_try03865 = EID{ToType(g0387UU.Id()).Included(ToType(C_integer.Id())).Id(),0}
            }
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (v_and3-g0381I_try03823) */
          if ErrorIn(v_and3_try03865) {g0381I_try03823 = v_and3_try03865
          } else {
          v_and3 = ToBoolean(OBJ(v_and3_try03865))
          if (v_and3 == CFALSE) {g0381I_try03823 = EID{CFALSE.Id(),0}
          } else /* arg:5 */{ 
            g0381I_try03823 = EID{CTRUE.Id(),0}/* arg-5 */} 
          /* arg-4 */} 
        }}
        /* and-3 */} 
      /* ERROR PROTECTION INSERTED (g0381I-Result) */
      if ErrorIn(g0381I_try03823) {Result = g0381I_try03823
      } else {
      g0381I = ToBoolean(OBJ(g0381I_try03823))
      if (g0381I == CTRUE) /* If:3 */{ 
        Result = F_Generate_cast_prefix_class(C_boolean,s)
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC("BitVectorContains(")
        Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{C_integer.Id(),0}))
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC(",")
        Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_integer.Id(),0}))
        /* ERROR PROTECTION INSERTED (Result-Result) */
        if !ErrorIn(Result) {
        PRINC(")")
        F_Generate_cast_post_class(C_boolean,s)
        PRINC("")
        Result = EVOID
        }}}
        } else {
        var g0389I *ClaireBoolean  
        var g0389I_try03904 EID 
        /* or:4 */{ 
          var v_or4 *ClaireBoolean  
          
          var v_or4_try03915 EID 
          /* Let:5 */{ 
            var g0392UU *ClaireClass  
            /* noccur = 1 */
            var g0392UU_try03936 EID 
            g0392UU_try03936 = Language.F_static_type_any(a2)
            /* ERROR PROTECTION INSERTED (g0392UU-v_or4_try03915) */
            if ErrorIn(g0392UU_try03936) {v_or4_try03915 = g0392UU_try03936
            } else {
            g0392UU = ToClass(OBJ(g0392UU_try03936))
            v_or4_try03915 = EID{ToType(g0392UU.Id()).Included(ToType(C_list.Id())).Id(),0}
            }
            /* Let-5 */} 
          /* ERROR PROTECTION INSERTED (v_or4-g0389I_try03904) */
          if ErrorIn(v_or4_try03915) {g0389I_try03904 = v_or4_try03915
          } else {
          v_or4 = ToBoolean(OBJ(v_or4_try03915))
          if (v_or4 == CTRUE) {g0389I_try03904 = EID{CTRUE.Id(),0}
          } else /* or:5 */{ 
            var v_or4_try03946 EID 
            /* Let:6 */{ 
              var g0395UU *ClaireClass  
              /* noccur = 1 */
              var g0395UU_try03967 EID 
              g0395UU_try03967 = Language.F_static_type_any(a2)
              /* ERROR PROTECTION INSERTED (g0395UU-v_or4_try03946) */
              if ErrorIn(g0395UU_try03967) {v_or4_try03946 = g0395UU_try03967
              } else {
              g0395UU = ToClass(OBJ(g0395UU_try03967))
              v_or4_try03946 = EID{ToType(g0395UU.Id()).Included(ToType(C_array.Id())).Id(),0}
              }
              /* Let-6 */} 
            /* ERROR PROTECTION INSERTED (v_or4-g0389I_try03904) */
            if ErrorIn(v_or4_try03946) {g0389I_try03904 = v_or4_try03946
            } else {
            v_or4 = ToBoolean(OBJ(v_or4_try03946))
            if (v_or4 == CTRUE) {g0389I_try03904 = EID{CTRUE.Id(),0}
            } else /* or:6 */{ 
              g0389I_try03904 = EID{CFALSE.Id(),0}/* org-6 */} 
            /* org-5 */} 
          }}
          /* or-4 */} 
        /* ERROR PROTECTION INSERTED (g0389I-Result) */
        if ErrorIn(g0389I_try03904) {Result = g0389I_try03904
        } else {
        g0389I = ToBoolean(OBJ(g0389I_try03904))
        if (g0389I == CTRUE) /* If:4 */{ 
          Result = F_Generate_cast_prefix_class(C_boolean,s)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{C_list.Id(),0}))
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(".Contain_ask(")
          Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_any.Id(),0}))
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(")")
          F_Generate_cast_post_class(C_boolean,s)
          PRINC("")
          Result = EVOID
          }}}
          } else {
          Result = F_Generate_cast_prefix_class(Optimize.C_EID,s)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          F_Generate_preCore_ask_void()
          PRINC("F_BELONG(")
          Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_any.Id(),0}))
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(",")
          Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{C_any.Id(),0}))
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(")")
          F_Generate_cast_post_class(Optimize.C_EID,s)
          PRINC("")
          Result = EVOID
          }}}
          /* If-4 */} 
        }
        /* If-3 */} 
      }
      /* If-2 */} 
    }
    return Result} 
  
// The EID go function for: belong_exp @ any (throw: true) 
func E_Generate_belong_exp_any (a1 EID,a2 EID,s EID) EID { 
    return /*(sm for belong_exp @ any= EID)*/ F_Generate_belong_exp_any(ANY(a1),ANY(a2),ToClass(OBJ(s)) )} 
  
// some special functions are open coded when used in a logical test
/* {1} OPT.The go function for: Compile/bool_exp(self:Call_method1,pos?:boolean) [] */
func F_Compile_bool_exp_Call_method1 (self *Language.CallMethod1 ,pos_ask *ClaireBoolean ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var m *ClaireMethod   = self.Arg
      /* noccur = 4 */
      /* Let:3 */{ 
        var a1 *ClaireAny   = self.Args.At(1-1)
        /* noccur = 3 */
        if (m.Id() == C_Generate__starnot_star.Value) /* If:4 */{ 
          Result = Core.F_CALL(Optimize.C_Compile_bool_exp,ARGS(a1.ToEID(),EID{pos_ask.Not.Id(),0}))
          /* If!4 */}  else if (m.Id() == C_Generate__starknown_star.Value) /* If:4 */{ 
          Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).EqualExp(a1,
            pos_ask.Not,
            CNULL,
            CTRUE.Id())
          /* If!4 */}  else if (m.Id() == C_Generate__starunknown_star.Value) /* If:4 */{ 
          Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).EqualExp(a1,
            pos_ask,
            CNULL,
            CTRUE.Id())
          /* If!4 */}  else if (m.Range.Included(ToType(C_boolean.Id())) == CTRUE) /* If:4 */{ 
          PRINC("(")
          Result = F_Generate_g_expression_Call_method1(self,C_boolean)
          /* ERROR PROTECTION INSERTED (Result-Result) */
          if !ErrorIn(Result) {
          PRINC(" ")
          F_Generate_sign_equal_boolean(pos_ask)
          PRINC(" CTRUE)")
          Result = EVOID
          }
          } else {
          Result = Core.F_SUPER(Optimize.C_Compile_bool_exp, C_any, ARGS(EID{self.Id(),0},EID{pos_ask.Id(),0}))
          /* If-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: Compile/bool_exp @ Call_method1 (throw: true) 
func E_Compile_bool_exp_Call_method1 (self EID,pos_ask EID) EID { 
    return /*(sm for Compile/bool_exp @ Call_method1= EID)*/ F_Compile_bool_exp_Call_method1(Language.To_CallMethod1(OBJ(self)),ToBoolean(OBJ(pos_ask)) )} 
  
// same thing for two arguments functions
/* {1} OPT.The go function for: Compile/bool_exp(self:Call_method2,pos?:boolean) [] */
func F_Compile_bool_exp_Call_method2 (self *Language.CallMethod2 ,pos_ask *ClaireBoolean ) EID { 
    var Result EID 
    /* Let:2 */{ 
      var m *ClaireMethod   = self.Arg
      /* noccur = 9 */
      /* Let:3 */{ 
        var p *ClaireProperty   = m.Selector
        /* noccur = 7 */
        /* Let:4 */{ 
          var lop *ClaireList   = ToGenerateCodeProducer(Optimize.C_PRODUCER.Value).OpenComparators
          /* noccur = 3 */
          /* Let:5 */{ 
            var a1 *ClaireAny   = self.Args.At(1-1)
            /* noccur = 7 */
            /* Let:6 */{ 
              var a2 *ClaireAny   = self.Args.At(2-1)
              /* noccur = 7 */
              if (p.Id() == Core.C__I_equal.Id()) /* If:7 */{ 
                Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).EqualExp(a1,
                  pos_ask.Not,
                  a2,
                  CFALSE.Id())
                /* If!7 */}  else if (p.Id() == Core.C_identical_ask.Id()) /* If:7 */{ 
                Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).EqualExp(a1,
                  pos_ask,
                  a2,
                  CTRUE.Id())
                /* If!7 */}  else if (p.Id() == C__equal.Id()) /* If:7 */{ 
                Result = ToGenerateGoProducer(Optimize.C_PRODUCER.Value).EqualExp(a1,
                  pos_ask,
                  a2,
                  CFALSE.Id())
                /* If!7 */}  else if (m.Id() == Optimize.C_Compile_m_member.Value) /* If:7 */{ 
                PRINC("(")
                Result = F_Generate_belong_exp_any(a1,a2,C_boolean)
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                PRINC(" ")
                F_Generate_sign_equal_boolean(pos_ask)
                PRINC(" CTRUE)")
                Result = EVOID
                }
                /* If!7 */}  else if ((lop.Memq(p.Id()) == CTRUE) && 
                  ((Core.F_domain_I_restriction(ToRestriction(m.Id())).Id() == C_float.Id()) || 
                      (Core.F_domain_I_restriction(ToRestriction(m.Id())).Id() == C_integer.Id()))) /* If:7 */{ 
                PRINC("(")
                Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{Core.F_domain_I_restriction(ToRestriction(m.Id())).Id(),0}))
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                PRINC(" ")
                if (pos_ask == CTRUE) /* If:8 */{ 
                  Result = Core.F_print_any(p.Id())
                  } else {
                  /* Let:9 */{ 
                    var g0397UU *ClaireAny  
                    /* noccur = 1 */
                    var g0397UU_try039810 EID 
                    /* Let:10 */{ 
                      var g0399UU int 
                      /* noccur = 1 */
                      var g0399UU_try040011 EID 
                      /* Let:11 */{ 
                        var g0401UU int 
                        /* noccur = 1 */
                        var g0401UU_try040212 EID 
                        g0401UU_try040212 = F_mod_integer((F_index_list(lop,p.Id())+1),4)
                        /* ERROR PROTECTION INSERTED (g0401UU-g0399UU_try040011) */
                        if ErrorIn(g0401UU_try040212) {g0399UU_try040011 = g0401UU_try040212
                        } else {
                        g0401UU = INT(g0401UU_try040212)
                        g0399UU_try040011 = EID{C__INT,IVAL((g0401UU+1))}
                        }
                        /* Let-11 */} 
                      /* ERROR PROTECTION INSERTED (g0399UU-g0397UU_try039810) */
                      if ErrorIn(g0399UU_try040011) {g0397UU_try039810 = g0399UU_try040011
                      } else {
                      g0399UU = INT(g0399UU_try040011)
                      g0397UU_try039810 = lop.At(g0399UU-1).ToEID()
                      }
                      /* Let-10 */} 
                    /* ERROR PROTECTION INSERTED (g0397UU-Result) */
                    if ErrorIn(g0397UU_try039810) {Result = g0397UU_try039810
                    } else {
                    g0397UU = ANY(g0397UU_try039810)
                    Result = Core.F_print_any(g0397UU)
                    }
                    /* Let-9 */} 
                  /* If-8 */} 
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                PRINC(" ")
                Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{Core.F_domain_I_restriction(ToRestriction(m.Id())).Id(),0}))
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                PRINC(")")
                Result = EVOID
                }}}
                /* If!7 */}  else if (m.Id() == C_Generate__starnth_integer_star.Value) /* If:7 */{ 
                PRINC("(BitVectorContains(")
                Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_integer.Id(),0}))
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                PRINC(",")
                Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{C_integer.Id(),0}))
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                PRINC(") ")
                F_Generate_sign_equal_boolean(pos_ask)
                PRINC(" CTRUE)")
                Result = EVOID
                }}
                /* If!7 */}  else if ((p.Id() == Core.C_inherit_ask.Id()) && 
                  (Core.F_domain_I_restriction(ToRestriction(m.Id())).Id() == C_class.Id())) /* If:7 */{ 
                PRINC("(")
                Result = Core.F_CALL(C_Generate_g_expression,ARGS(a1.ToEID(),EID{C_class.Id(),0}))
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                PRINC(".IsIn(")
                Result = Core.F_CALL(C_Generate_g_expression,ARGS(a2.ToEID(),EID{C_class.Id(),0}))
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                PRINC(") ")
                F_Generate_sign_equal_boolean(pos_ask)
                PRINC(" CTRUE)")
                Result = EVOID
                }}
                /* If!7 */}  else if (m.Range.Included(ToType(C_boolean.Id())) == CTRUE) /* If:7 */{ 
                PRINC("(")
                Result = F_Generate_g_expression_Call_method2(self,C_boolean)
                /* ERROR PROTECTION INSERTED (Result-Result) */
                if !ErrorIn(Result) {
                PRINC(" ")
                F_Generate_sign_equal_boolean(pos_ask)
                PRINC(" CTRUE)")
                Result = EVOID
                }
                } else {
                Result = Core.F_SUPER(Optimize.C_Compile_bool_exp, C_any, ARGS(EID{self.Id(),0},EID{pos_ask.Id(),0}))
                /* If-7 */} 
              /* Let-6 */} 
            /* Let-5 */} 
          /* Let-4 */} 
        /* Let-3 */} 
      /* Let-2 */} 
    return Result} 
  
// The EID go function for: Compile/bool_exp @ Call_method2 (throw: true) 
func E_Compile_bool_exp_Call_method2 (self EID,pos_ask EID) EID { 
    return /*(sm for Compile/bool_exp @ Call_method2= EID)*/ F_Compile_bool_exp_Call_method2(Language.To_CallMethod2(OBJ(self)),ToBoolean(OBJ(pos_ask)) )} 
  