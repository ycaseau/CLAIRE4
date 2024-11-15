/***********************************************************************/
/**   metaCLAIRE                                        Yves Caseau    */
/**   readme.cl for meta (CLAIRE in CLAIRE)                            */
/**  Copyright (C) 1998-2013 Yves Caseau. All Rights Reserved.         */
/***********************************************************************/   
                 
This directory contains the first part of the meta-description of the CLAIRE system
It is composed of 3 modules:

/*************/
/**  Core   **/
/*************/

This is the claire-part of the microClaire library. It is defined by the following 
files:

- method:       first set of key methods, for evaluating methods and using slots
- object:       second set of key system methods, mostly for instantiation
- function:     the functions from the
- type:         the reflective definition of the type system

Here are the table of contents for each files + a short description

// *********************************************************************
// *  Cntents of method.cl                                             *
// *      Part 1: Lambda & Methods Evaluation                          *
// *      Part 2: Update methods                                       *
// *      Part 3: Management of definition(p)                          *
// *      Part 4: Matching Methods                                     *
// *********************************************************************
// *********************************************************************
// *  Contents of function.cl                                          *
// *   Part 1: Basics of pretty printing                               *
// *   Part 2: Methods for CLAIRE objects                              *
// *   Part 3: System Methods                                          *
// *   Part 4: Methods for Native entities                             *
// *********************************************************************
// *********************************************************************
// *  Contents of object.cl                                            *
// *   Part 1: Ask, debug & trace                                      *
// *   Part 2: Tables                                                  *
// *   Part 3: Demons & relations for the logic modules                *
// *   Part 4: Basics of Exceptions                                    *
// *********************************************************************
// *********************************************************************
// *  Contents of type.cl                                              *
// *    Part 1: Common Set Methods                                     *
// *    Part 2: definition of the type operators                       *
// *    Part 3: Interface methods                                      *
// *    Part 4: Lattice methods                                        *
// *    Part 5: Type methods                                           *
// *********************************************************************

/*****************/
/**  Language   **/
/*****************/

This module contains the "self-description" of the CLAIRE language, that is:
  - the classes for each syntactic construct
  - the self-print method : how to print each type of instruction
  - self-eval, how to evaluate = the CLAIRE definition of the interpreter

It is defined by the following four files:

- pretty:   this file contains the top of the "instruction" class hierarchy.
- call:     this file contains functional calls (ex-messages)
- control:  this file contains all major control structures
- define:   this file contains all definition & instantiation instructions

// *********************************************************************
// * Contents of pretty.cl:                                            *
// *      Part 1: unbound_symbol and variables                         *
// *      Part 2: lambdas                                              *
// *      Part 3: close methods for lattice_set instantiation          *
// *      Part 4: Pretty printing                                      *
// *********************************************************************
// *********************************************************************
// * Contents of call.cl                                               *
// *      Part 1: the basic object messages                            *
// *      Part 2: Basic structures                                     *
// *      Part 3: Specialized structures                               *
// *      Part 4: Functions on instructions                            *
// *********************************************************************
// *********************************************************************
// *  contents of control.cl                                           *
// *     Part 1: If, Do, Let                                           *
// *     Part 2: set control structures                                *
// *     Part 3: other control structures                              *
// *     Part 4: the constructs                                        *
// *********************************************************************
// **************************************************************************
// * Contents of define.cl:                                                 *
// *     Part 1: Definition instructions (Defobj, Defclass, Defmethod ...)  *
// *     Part 2: the instantiation macros                                   *
// *     Part 3: the useful stuff                                           *
// *     Part 4: the other macros                                           *
// **************************************************************************

/**************/
/**  Reader  **/
/**************/

This module contains the I/O library for CLAIRE: how to handle ports/files and
to read stuff in them. CLAIRE implements a syntactic reader very similar to the
one of LISP.
It is organized around 4 files:

- read:         this file contains the reader object and the top-level read functions
- syntax:       this file contains specialized reading methods
- file:         this file contains all that is related to files + top-level
- inspect:      this file contains the CLAIRE run-time tools: inspect, trace & debug

// **********************************************************************
// *  Content of read.cl:                                               *
// *   Part 1: The reader object                                        *
// *   Part 2: reading blocks                                           *
// *   Part 3: reading expressions                                      *
// *   Part 4: miscellaneous                                            *
// **********************************************************************
// **********************************************************************
// *  Content of syntax.cl                                              *
// *   Part 1: read operation expressions (<exp> <op> <exp>)            *
// *   Part 2: read control structures                                  *
// *   Part 3: read functional calls                                    *
// *   Part 4: read definitions                                         *
// **********************************************************************
// **********************************************************************
// * Contents of file.cl:                                               *
// *  Part 1: Utilities                                                 *
// *  Part 2: Loading                                                   *
// *  Part 3: Reading in a file/string & top-level                      *
// *  Part 4: The show & kill methods + macro-methods                   *
// **********************************************************************
// *********************************************************************
// * Contents of inspect.cl                                            *
// *      Part 1: Inspection                                           *
// *      Part 2: Trace                                                *
// *      Part 3: Debugger                                             *
// *      Part 4: Stepper                                              *
// *      Part 5: Profiler                                             *
// *********************************************************************