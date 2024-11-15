/***********************************************************************/
/**   compile CLAIRE4 directory                         Yves Caseau    */
/**   readme.cl                                                        */
/**   Copyright (C) 1998-2021 Yves Caseau. All Rights Reserved.        */
/***********************************************************************/


This is the CLAIRE 4 compiler, made of Optimize and Generate
Optimize is the "target-code-independent" part of the compiler : from CLAIRE instruction to optimized-claire-instruction.
Generate is the target code compiler, geared towards go - THIS IS THE NEW PART IN CLAIRE 4 

We try to keep the code generation generic so that a Java or Swift code producer could be derived in the future 
However, the GC management stuff will disappear forever

/****************/
/**  Optimize  **/
/****************/

CLAIRE instruction-to-instruction optimization is defined by a handful of methods:

  c_type(x)  is the CLAIRE type of x
  c_code(x)  is an optimized instruction
  & c_code(x,s) is an optimized expression of sort s
  c_gc?(x)   [boolean] true if x must be protected from garbage collection
  c_sort(x)  the sort of the expression x (precise sort)
  c_status(x,l) [bitvector] abstract interpretation of x according to a set
             of criterions (allocation, update, etc...)

The sort is a class sub-hierarchy which has an homogeneous way to be translated into the
target language. Here th sorts are integer, char, x <= imported, object, entity

osystem.cl      This file contains the global parameter objects and the key methods: c_type, c_code, c_gc?, c_sort, c_status
otool.cl        This file contains the auxiliary methods for the source optimizer
ocall.cl        this is the heart of the CLAIRE optimizer : message to function calls
ocontrol.cl     contains the optimizer for the control structures - most of them are replaced with simpler structures
                that may be translated directly into the target language
odefine.cl      optimization of the definition instruction -> sort-of macroexpansion

// ******************************************************************
// *   contents of osystem.cl                                       *
// *    Part 1: General Global Variables and Properties             *
// *    Part 2: The defaults for c_type, c_code, c_gc and c_sort    *
// *    Part 3: c_throw and status(m:method)                        *
// *    Part 4: Names & identifiers management                      *
// ******************************************************************

// ******************************************************************
// *  contents of otool.cl                                          *
// *    Part 1: New Instructions & associated stuff                 *
// *    Part 2: Optimizer Warnings                                  *
// *    Part 3: Type Handling                                       *
// *    Part 4: Garbage Collection functions                        *
// *    Part 5: Miscellaneous                                       *
// ******************************************************************

// ******************************************************************
// *  Contents of ocall.cl                                          *
// *    Part 1: Restruction Binding                                 *
// *    Part 2: Generic c_type & c_code                             *
// *    Part 3: specialized c_code                                  *
// *    Part 4: Method optimizing                                   *
// *    Part 5: inline methods                                      *
// ******************************************************************
// ******************************************************************
// * Contents of ocontrol.cl                                        *
// *     Part 1: Basic Instructions                                 *
// *     Part 2: other control structures                           *
// *     Part 3: If, Case, Do, Let                                  *
// *     Part 4: Loops                                              *
// *     Part 5: Iterate                                            *
// ******************************************************************
// ******************************************************************
// *  Contents of define.cl                                         *
// *     Part 1: Set, List and Tuple creation                       *
// *     Part 2: Object definition                                  *
// *     Part 3: Method instantiation                               *
// *     Part 4: Inverse Management                                 *
// ******************************************************************

/****************/
/**  goGen     **/
/****************/

gosystem.cl


gogen.cl


goexp.cl


gostat.cl
// statement is implemented as a general method that calls a restriction
//        g_statement(self:any,e:class,v:string,err:boolean,loop:any)
// (1) e is the goType that the variable v must receive (HENCE goCast must be inserted)
//     a proper goType is a class, or EID, or void
// (2) The argument v is the named of the C variable in which the
//     result of the evaluation must be placed.
// (3) err tells if an error is possible, which forces to create a chain an not a block (see Do for example)
//     Note : if err = true, s is expected to be EID to (a) force a chain (b) place the error value in v
// (4) loop is either false (not within a loop) or a tuple(v,s) inside the compiling of While/For
//     This tuple describes the vreturn Variable in case a break(v) is encountered

//**********************************************************************
//*  Table of contents of gostat.cl:                                   *
//*          Part 1: Management of exception (NEW in claire4)          *
//*          Part 2: Unfolding of complex expressions                  *
//*          Part 2: Basic control structures                          *
//*          Part 3: iteration                                         *
//*          Part 4: CLAIRE-specific structures                        *
//**********************************************************************





