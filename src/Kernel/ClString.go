// ==================================================================
// microCLAIRE                                              CLAIRE 4
// golang microClaire Kernel - started on June 21st, 2020
//
//  Copyright (C) 2020-2022 Yves Caseau. All Rights Reserved.
//  Redistribution and use in source and binary forms are permitted
//  provided that source distribution retains this entire copyright
//  notice and comments.
//  THIS SOFTWARE IS PROVIDED AS IS AND WITHOUT ANY WARRANTY,
//  INCLUDING, WITHOUT LIMITATION, THE IMPLIED WARRANTIES OF
//  MERCHANTABILTY AND FITNESS FOR A PARTICULAR PURPOSE
//
// clString.go
// ==================================================================

// TODO :
// - ensure that all EID{x,y} use x = object and not another subclass

// this file contains the string code for microClaire
// strings and chars (rune) are native imports : this file contains string and rune
// functions
// symbols and modules are claire objects
// ClairePorts are the most specific piece of this file : buffered for better performance

package Kernel

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
	"math"
)

// +---------------------------------------------------------------------------+
// |  Table of contents                                                        |
// |  Part 1: Char objects (Runes)                                             |
// |  Part 2: Strings                                                          |
// |  Part 3: Symbols                                                          |
// |  Part 4: Modules                                                          |
// |  Part 5: Ports (i/o + string buffers)                                     |
// +---------------------------------------------------------------------------+

// exported vars
var C_char_I *ClaireProperty


// +---------------------------------------------------------------------------+
// |  Part 1: Char objects (Runes)                                             |
// +---------------------------------------------------------------------------+

// a Claire Char is a Rune

// constructor for Claire Char
func createChar(x rune) *ClaireChar {
	var o *ClaireChar = new(ClaireChar)
	o.Isa = C_char
	o.Value = x
	return o
}

// constructor
// Notice that we cash the ascii char to avoid useless allocation
func MakeChar(r rune) *ClaireChar {
	if r > 0 && r < 128 {
		return ClRes.ascii[r]
	} else {
		return createChar(r)
	}
}

// initialization: create all chars (this is called once ClaireChar is properly defined)
func InitChar() {
	ClRes.ascii = make([]*ClaireChar, 128)
	for i := 0; i < 128; i++ {
		ClRes.ascii[i] = createChar(rune(i))
	}
}

// check that a char is not a special char
func alpha_char(c rune) *ClaireBoolean {
	if c == CEOF || c == '\n' ||
		c == '}' || c == ')' || c == ']' ||
		c == '{' || c == '(' || c == '[' ||
		c == 9 || c == ';' || c == '|' ||
		c == ',' || c == '/' || c == ':' ||
		c == '@' || c == '.' || c == '\r' ||
		c == '<' || c == '>' || c == ' ' {
		return CTRUE
	} else {
		return CFALSE
	}
}

// ---  API functions ---------------------------------------------------

// translate a integer into a char
func F_char_I_integer(n int) rune { return rune(n) }
func E_char_I_integer(n EID) EID  { return EID{C__CHAR, IVAL(INT(n))} }

// reciprocate function
func F_integer_I_char(r rune) int { return (int)(r) }
func E_integer_I_char(c EID) EID  { return EID{C__INT, IVAL((int)(CHAR(c)))} }

// princ a char / int / float
func F_princ_char(r rune) { ClEnv.Cout.Putc(r) }
func E_princ_char(c EID) EID {
	F_princ_char(CHAR(c))
	return EVOID
}

// special consversion (language dependent)
func F_c_princ_char(r rune) {
	switch r {
	case '.':
		PRINC("_dot")
	case '/':
		PRINC("_7")
	case '\\':
		PRINC("_backslash")
	case '&':
		PRINC("_and")
	case '-':
		PRINC("_dash")
	case '+':
		PRINC("_plus")
	case '€':
		PRINC("_Z")
	case '%':
		PRINC("_Z")
	case '*':
		PRINC("_star")
	case '?':
		PRINC("_ask")
	case '!':
		PRINC("_I")
	case '<':
		PRINC("_inf")
	case '>':
		PRINC("_sup")
	case '=':
		PRINC("_equal")
	case ',':
		PRINC("_comma")
	case '^':
		PRINC("_exp")
	case '@':
		PRINC("_at")
	case '~':
		PRINC("_tilda")
	case ']':
		PRINC("_brack")
	case ':':
		PRINC("L_")
	case '\'':
		PRINC("_prime")
	case '$':
		PRINC("_dollar") // v3.2.14
	case '≤':
		PRINC("_two")
	default:
		ClEnv.Cout.Putc(r)
	}
}

func E_c_princ_char(c EID) EID {
	F_c_princ_char(CHAR(c))
	return EVOID
}

// +---------------------------------------------------------------------------+
// |  Part 2: Strings                                                          |
// +---------------------------------------------------------------------------+

// constructor
// note : we could decide to cash the number of chars (later)
func MakeString(s string) *ClaireString {
	o := new(ClaireString)
	o.Isa = C_string
	o.Value = s
	return o
}

// --- there are all API functions since string are imported char* -----

// length of a string is actually slow and complex
func F_length_string(s *ClaireString) int { return utf8.RuneCountInString(s.Value) }
func E_length_string(s EID) EID {
	return EID{C__INT, IVAL(F_length_string(ToString(OBJ(s))))}
}

// make a local copy of a string
func F_copy_string(s *ClaireString) *ClaireString {
	s2 := make([]byte, len(s.Value))
	copy(s2, s.Value)
	return MakeString(string(s2))
}

func E_copy_string(s EID) EID {
	return EID{F_copy_string(ToString(OBJ(s))).Id(), 0}
}

// internal form of princ_string for the compiler : princ to any port
func F_princ_string(s *ClaireString) {
	for _, r := range s.Value {
		ClEnv.put(r)
	}
}

// even more internal  => optimization for princ(constant string)
func PRINC(s string) {
	for _, r := range s {
		ClEnv.put(r)
	}
}

func E_princ_string(s EID) EID {
	F_princ_string(ToString(OBJ(s)))
	return EVOID
}

// print a string with the "" (what you see is what you read)
// TODO: see what is needed for "special char"
func F_self_print_string(s *ClaireString) {
	ClEnv.put('"')
	for _, r := range s.Value {
		if r == '"' || r == '\\' {
			ClEnv.put('\\')
		}
		if r == '\n' {
			ClEnv.put('\\')
			ClEnv.put('n')
		} else if r == '\t' {
			ClEnv.put('\\')
			ClEnv.put('t')
		} else {
			ClEnv.put(r)
		}
	}
	ClEnv.put('"')
}

func E_self_print_string(s EID) EID {
	F_self_print_string(ToString(OBJ(s)))
	return EVOID
}

// concatenate two strings
func F_append_string(s1 *ClaireString, s2 *ClaireString) *ClaireString {
	return MakeString(s1.Value + s2.Value)
}

func E_append_string(s1 EID, s2 EID) EID {
	return EID{F_append_string(ToString(OBJ(s1)), ToString(OBJ(s2))).Id(), 0}
}

// finds the integer value
// not an integer, does not produce an error (TODO : fix ?)
func F_integer_I_string(s *ClaireString) int {
	i, err := strconv.Atoi(s.Value)
	if err == nil {
		return i
	} else {
		return 0
	}
}

func E_integer_I_string(s EID) EID {
	return EID{C__INT, IVAL(F_integer_I_string(ToString(OBJ(s))))}
}

// create a substring from a string
// this cannot be based on nth_string which is too slow
func F_substring_string(s *ClaireString, n int, m int) *ClaireString {
	ClEnv.bufferStart()
	i := 1 // Char counter
	for _, r := range s.Value {
		if i >= n && i <= m {
			ClEnv.pushChar(r)
		}
		i++
	}
	return MakeString(ClEnv.bufferCopy())
}

func E_substring_string(s EID, n EID, m EID) EID {
	return EID{F_substring_string(ToString(OBJ(s)), INT(n), INT(m)).Id(), 0}
}

func E_slice_string(s EID, n EID, m EID) EID {
	return EID{F_substring_string(ToString(OBJ(s)), INT(n), INT(m)).Id(), 0}
}

// look for th eposition of the CHAR c in s
func F_get_string(s *ClaireString, c rune) int {
	i := 1
	// fmt.Printf("--- call get on %s and c=%c\n", s.Value, c)
	for _, r := range s.Value {
		if r == c {
			return i
		} else {
			i++
		}
	}
	return 0
}

func E_get_string(s EID, c EID) EID {
	return EID{C__INT, IVAL(F_get_string(ToString(OBJ(s)), CHAR(c)))}
}

// compare two strings
func F_less_string(s1 *ClaireString, s2 *ClaireString) *ClaireBoolean {
	if s1.Value < s2.Value {
		return CTRUE
	} else {
		return CFALSE
	}
}

func E_less_string(s1 EID, s2 EID) EID {
	return EID{F_less_string(ToString(OBJ(s1)), ToString(OBJ(s2))).Id(), 0}
}

// test is a string is included into another
// return the match position (i > 0) if true, 0 otherwise
// p tells if lower /uper case match is acceptable
// TODO: write a faster version without nth_string ?
func F_included_string(s1 *ClaireString, s2 *ClaireString, p *ClaireBoolean) int {
	if p == CTRUE {
		s1 = MakeString(strings.ToLower(s1.Value))
		s2 = MakeString(strings.ToLower(s2.Value))
	}
	n := F_length_string(s1)
	m := F_length_string(s2)
	for i := 1; i+m <= n; i++ {
		j := 0
		for _, w := range s2.Value {
			if w == s1.At(i+j) {
				j++
			} else {
				break
			}
		}
		if j == m {
			return i
		}
	}
	return 0
}

func E_included_string(s1 EID, s2 EID, p EID) EID {
	return EID{C__INT, IVAL(F_included_string(ToString(OBJ(s1)),
		ToString(OBJ(s2)), ToBoolean(OBJ(p))))}
}

// get the CHAR at the i-th place in s - this is slow in golang (by construction)
// conversion to list<char> should be prefered
// this function may raise an error ... TODO: consider the use of character CEOF (:))
func F_nth_string(s *ClaireString, n int) EID {
	i := 1
	for _, r := range s.Value {
		if i == n {
			return EID{C__CHAR, CVAL(r)}
		} else {
			i++
		}
	}
	return Cerror(11, AnyInteger(n), s.Id())
}

func E_nth_string(s EID, n EID) EID { return F_nth_string(ToString(OBJ(s)), INT(n)) }

// faster version for compiler use (when we know that the bounds are ok)
func (s *ClaireString) At(n int) rune {
	i := 1
	for _, r := range s.Value {
		if i == n {
			return r
		} else {
			i++
		}
	}
	return CEOF
}

// set the char at the i_th place in s
// this is very unefficient in go - does not work on strings - hence the different form
// note: this is not the regular pattern for naming => special case for the compiler
// NOTE: nth_set @ string is DEPRECATED
func F_nth_set_string(s *ClaireString, n int, c rune) EID {
	ClEnv.bufferStart()
	i := 1
	for _, r := range s.Value {
		if i == n {
			ClEnv.pushChar(c)
		} else {
			ClEnv.pushChar(r)
		}
		i++
	}
	s2 := ClEnv.bufferCopy()
	if n >= i {
		return Cerror(11, AnyInteger(n), s.Id())
	} else {
		s.Value = s2
		return EID{s.Id(), 0}
	}
}

func E_nth_set_string(s EID, n EID, c EID) EID {
	return F_nth_set_string(ToString(OBJ(s)), INT(n), CHAR(c))
}

// shrink @ string is no longer supported  in CLAIRE 4

// create a new string from an integer - use golang
func F_string_I_integer(n int) *ClaireString { return MakeString(fmt.Sprintf("%d", n)) }

func E_string_I_integer(n EID) EID {
	return EID{F_string_I_integer(INT(n)).Id(), 0}
}

// allocate a list with n member equal to m - deprecated
func F_make_string_integer(n int, c rune) *ClaireString {
	if n < 0 {
		n = 0
	} // simpler with CLAIRE 4
	ClEnv.bufferStart()
	for i := 0; i < n; i++ {
		ClEnv.pushChar(c)
	}
	return MakeString(ClEnv.bufferCopy())
}
func E_make_string_integer(n EID, c EID) EID {
	return EID{F_make_string_integer(INT(n), CHAR(c)).Id(), 0}
}

// create a string from a list (v3.0.44) - necessary with CLAIRE 4 since list of chars
// is a proper way to customize strings
// TODO : find the best way to write a rune at a position in a byte[]
func (l *ClaireList) MakeString() EID {
	if l.of != ToType(C_char.Id()) {
		return Cerror(22, l.Id(), C_char.Id())
	} else {
		ClEnv.bufferStart()
		for _, r := range l.ValuesO() {
			ClEnv.pushChar(ToChar(r).Value)
		}
		return EID{MakeString(ClEnv.bufferCopy()).Id(), 0}
	}
}

func E_make_string_list(l EID) EID { return ToList(OBJ(l)).MakeString() }

// return the list of chars from a string
// note : this must be compiled through a specialized iteration
func F_list_I_string(s *ClaireString) *ClaireList {
	n := F_length_string(s)
	var ls []*ClaireAny = make([]*ClaireAny, n)
	i := 0
	for _, r := range s.Value {
		ls[i] = MakeChar(r).ToAny()
		i++
	}
	o := new(ClaireListObject)
	o.Isa = C_list
	o.Srange = C_object
	o.of = ToType(C_char.Id())
	return ToList(o.Id())
}

func E_list_I_string(s EID) EID { return EID{F_list_I_string(ToString(OBJ(s))).Id(), 0} }

// print the name of an object as a C identifier
func F_c_princ_string(s *ClaireString) {
	for _, r := range s.Value {
		F_c_princ_char(r)
	}
}

func E_c_princ_string(s EID) EID {
	F_c_princ_string(ToString(OBJ(s)))
	return EVOID
}

// +---------------------------------------------------------------------------+
// |  Part 3: Symbols                                                          |
// +---------------------------------------------------------------------------+

// -------------- member functions --------------------------------------------

// generates a symbol in a module - this is the basic function that creates the object
// the better function to use is createSymbol =>  includes a lookup
func (m *ClaireModule) produceSymbol(name string) *ClaireSymbol {
	o := MakeSymbol(name,m)
	m.table[name] = o
	return o
}


// read the value bound to a given symbol s. We create an unbound symbol object if necessary
// warning : in CLAIRE 3.5, Symbol::getValue = make, get_symbol = read
func (s *ClaireSymbol) makeValue() *ClaireAny {
	if (s.value == CNULL && s != unknownName) || s.value == nil {
		return MakeUnboundSymbol(s).Id()
	} else {
		return s.value
	}
}

// --------------- API functions ---------------------------------------------

// returns the string
func (s *ClaireSymbol) String_I() *ClaireString {
	return MakeString(s.key)
}

func E_string_I_symbol(s EID) EID {
	return EID{ToSymbol(OBJ(s)).String_I().Id(), 0}
}

// returns the module
func (s *ClaireSymbol) Module_I() *ClaireModule {
	return s.module_I
}

func E_module_I_symbol(s EID) EID {
	return EID{ToSymbol(OBJ(s)).module_I.Id(), 0}
}

// returns the value
func (s *ClaireSymbol) Value() *ClaireAny {
	if s.value == nil {
		return CNULL
	} else {
		return s.value
	}
}

func E_value_symbol(s EID) EID {
	return EID{ToSymbol(OBJ(s)).Value(), 0}
}

// create a symbol in the current module
func F_symbol_I_string(s *ClaireString, m *ClaireModule) *ClaireSymbol {
	return m.createSymbol(s.Value)
}

func E_symbol_I_string(s EID, m EID) EID {
	return EID{ToModule(OBJ(m)).createSymbol(ToString(OBJ(s)).Value).Id(), 0}
}

// writes the value of a symbol
func (s *ClaireSymbol) Put(x *ClaireAny) *ClaireAny {
	s.value = x
	return x
}
func E_put_symbol(s EID, x EID) EID { return EID{ToSymbol(OBJ(s)).Put(ANY(x)).Id(), 0} }

/*
// return the value : unknown if unknown ?  (used to be called get_symbol)
// deprecated => use value @ symbol ! kept only because it helps to move from 3.5 to 4.0
func (s *ClaireSymbol) Get() *ClaireAny {
	if s.value == nil {
		return CNULL
	} else {
		return s.value
	}
}

func E_get_symbol(s EID) EID { return EID{ToSymbol(OBJ(s)).Get().Id(), 0} } 
*/

// concatenate two symbols, or a symbol and a string or a symbol and an integer
// the result is a symbol in the module of the first symbol
func F_append_symbol(s1 *ClaireSymbol,s2 *ClaireAny) *ClaireSymbol {
	ClEnv.bufferStart()
	ClEnv.pushString(s1.key)
	if s2.Isa == C_integer {
		ClEnv.pushInteger(ToInteger(s2).Value)
	} else if s2.Isa == C_symbol {
		ClEnv.pushString(ToSymbol(s2).key)
	} else if s2.Isa == C_string {
		ClEnv.pushString(ToString(s2).Value)
	}
	return s1.module_I.createSymbol(ClEnv.bufferCopy())
}

func E_append_symbol(s1 EID, s2 EID) EID {
	return EID{F_append_symbol(ToSymbol(OBJ(s1)),OBJ(s2)).Id(), 0}
}

// print a symbol with its application name
func (s *ClaireSymbol) Princ() {
	if s.module_I != C_claire && s.module_I != ClEnv.Module_I {
		PRINC(s.module_I.Name.key)
		ClEnv.put('/')
	}
	PRINC(s.key)
}

func E_princ_symbol(s EID) EID {
	ToSymbol(OBJ(s)).Princ()
	return EVOID
}

// find the module where the object is defined
func (s *ClaireSymbol) Defined() *ClaireModule {
	if s.definition != nil {
		return s.definition
	} else {
		return s.module_I
	}
}
func E_defined_symbol(s EID) EID { return EID{ToSymbol(OBJ(s)).Defined().Id(), 0} }

// create a new name
func F_gensym_string(s *ClaireString) *ClaireSymbol {
	ClEnv.bufferStart()
	for _, r := range s.Value {
		ClEnv.pushChar(r)
	}
	ClEnv.pushChar('0' + (rune)((ClEnv.gensym%10000)/1000))
	ClEnv.pushChar('0' + (rune)((ClEnv.gensym%1000)/100))
	ClEnv.pushChar('0' + (rune)((ClEnv.gensym%100)/10))
	ClEnv.pushChar('0' + (rune)(ClEnv.gensym%10))
	ClEnv.gensym = ClEnv.gensym + 1
	return C_claire.createSymbol(ClEnv.bufferCopy())
}

func E_gensym_string(s EID) EID {
	return EID{F_gensym_string(ToString(OBJ(s))).Id(), 0}
}

// print a symbol (the name of an object) as a C identifier */
func (s *ClaireSymbol) CPrinc() {
	if s.module_I != C_claire {
		s.module_I.Name.CPrinc()
		ClEnv.put('_')
	}
	F_c_princ_string(MakeString(s.key))
}

func E_c_princ_symbol(s EID) EID {
	ToSymbol(OBJ(s)).CPrinc()
	return EVOID
}

// +---------------------------------------------------------------------------+
// |  Part 4: Modules                                                          |
// +---------------------------------------------------------------------------+

// ------------- member functions for modules ---------------------------------------

// lookup: check if a string in a given module is represented by a symbol
// => returns nil if no symbol is found - does NOT create a new symbol
// this method embodies the strategy for looking in upper modules (namespace inheritance)
func (m *ClaireModule) Lookup(name string) *ClaireSymbol {
    s := m.table[name]
	if ClEnv.Verbose > 10   {fmt.Printf("lookup(%s) in module %s -> %p\n", name, m.Name.key, s)}
	if s != nil || m == C_claire {
		return s
	} else {
		return m.PartOf.Lookup(name)
	}
}

// hard debug for string - understand what a string that looks like s is not s
func string_explode(name string) {
	i := len(name)
	s2 := make([]byte, i)
	copy(s2, name)
	fmt.Printf("%s:%d : [",name,i)
	for k := 0; k < i ; k++ {
		fmt.Printf("%d ",s2[k])
	}
	fmt.Printf("]\n")
}

// Get a symbol (even if none is there => create it) in the module with the given name,
// this is a combination of lookup + make(Symbol)
// notice that we do not inherit junk (undefined) but rather create a new symbol
// warning : replace module::makeSymbol  (conflict name vs CLAIRE 3.5)
func (m *ClaireModule) createSymbol(name string) *ClaireSymbol {
	cur := m.Lookup(name)
	if cur != nil && (cur.value != CNULL || cur.module_I == m || cur == unknownName) {
		return cur
	} else {
		if ClEnv.Verbose == 12 {fmt.Printf("--- Create a new symbol in %s for %s \n",m.Prt(),name)}
		return m.produceSymbol(name)
	}
}

// similar but also fills the key slots for the module (compiler method)
// new status (0:default, 1:start, 2 compiled, 3:c+loaded, 4:c+l+trace, 5:c+delayed)
// in C++ there was a difference between namespaces and module
func InitModule(name string, father *ClaireModule, usage *ClaireList, dir string, files *ClaireList) *ClaireModule {
	it := MakeModule(name, father)
	it.Uses = usage             // other modules that are used
	it.Source = MakeString(dir) // directory where the sources can be found
	it.MadeOf = files
	it.Comment = MakeString(name)
	it.Status = 3
	father.Parts.AddFast(it.Id())
	return it
}

// --- API functions for modules ---------------------------------------------------

// create the namespace
func (x *ClaireModule) Namespace() {
	x.table = make(map[string]*ClaireSymbol)
}

func E_namespace_module (m EID) EID {
	ToModule(OBJ(m)).Namespace()
	return EVOID
}

// open a module x with module identifier index
func (m *ClaireModule) Begin() {
	// fmt.Printf(">> Begin module %s\n",m.Prt())
	ClEnv.moduleStack.AddFast(ClEnv.Module_I.ToAny())
	ClEnv.Module_I = m
}

func E_begin_module(m EID) EID {
	ToModule(OBJ(m)).Begin()
	return EVOID
}

// close a module
func (x *ClaireModule) End() {
	// fmt.Printf("<< End module %s\n",x.Prt())
	n := len(ClEnv.moduleStack.Values)
	if n == 0 {
		ClEnv.Module_I = C_claire
	} else {
		ClEnv.Module_I = ToModule(ClEnv.moduleStack.Values[n-1])
		ClEnv.moduleStack.Nth_dash(n)
	}
}

func E_end_module(m EID) EID {
	ToModule(OBJ(m)).End()
	return EVOID
}

// the old internal function used for get_value @ string
func F_value_string(s *ClaireString) *ClaireAny {
	return F_value_module(ClEnv.Module_I, s)
}

func E_value_string(s EID) EID { return EID{F_value_string(ToString(OBJ(s))), 0} }

// the new internal function
func F_value_module(m *ClaireModule, name *ClaireString) *ClaireAny {
	s := m.Lookup(name.Value)
	if s == nil {
		return CNULL
	} else {
		return s.value
	}
}

func E_value_module(m EID, s EID) EID {
	return EID{F_value_module(ToModule(OBJ(m)), ToString(OBJ(s))), 0}
}

// access to the symbol - if it exists, CNULL otherwise
func F_get_symbol_module(m *ClaireModule, name *ClaireString) *ClaireAny {
	s := m.Lookup(name.Value)
	// fmt.Printf("@@@@  lookup s:%s -> %x\n", name.Value, s)
	if s == nil {
		return CNULL
	} else {
		return s.Id()
	}
}

func E_get_symbol_module(m EID, s EID) EID {
	return EID{F_get_symbol_module(ToModule(OBJ(m)), ToString(OBJ(s))), 0}
}


// access to environment variable
func F_getenv_string (s *ClaireString) *ClaireString {
	return MakeString(os.Getenv(s.Value))
}

func E_getenv_string (s EID) EID {
	return EID{F_getenv_string(ToString(OBJ(s))).Id(),0}}



// +---------------------------------------------------------------------------+
// |  Part 5: Ports (i/o + string buffers)                                     |
// +---------------------------------------------------------------------------+

// note : the port section is moved to ClString because of runes in golang

// what matters is the status
//     0: file buffer for reading (with a buffer size BUFFSIZE)
//     1: file buffer for writing
//     2: port is a string buffer (write port, use buffer size BUFPORT)
//     3: port is a string buffer (read port)
//     4: input file for reading (no buffer : slow, necessary for stdin)

// Claire Port constructors
// we will need to split read/write because of buffered
// buffer copy size
var BUFSIZE int = 200   // test impact on speed (200 is the good value)
var BUFPORT int = 20000 // max size of string port

func makeFile(f *os.File) *ClairePort {
	o := new(ClairePort)
	o.Isa = C_port
	o.file = f
	return o
}

// create a port for reading from a file (buffered) - this will not work with stdin
func MakeInBufferPort(f *os.File) *ClairePort {
	o := makeFile(f)
	o.buffer = make([]byte, BUFSIZE)
	o.nChar = BUFSIZE
	n, _ := f.Seek(0, 2)
	f.Seek(0, 0)
	o.nEof = BUFSIZE * 10 // never
	o.size = (int)(n)
	o.status = 0
	// fmt.Println("A call to seek give size = ",n)
	return o
}

// create a port for reading from a file (buffered) - this will not work with stdin
func MakeInPort(f *os.File) *ClairePort {
	o := new(ClaireGoPort)
	o.Isa = C_port
	o.file = f
	o.reader = bufio.NewReader(f)
	o.status = 4
	o.firstc = 32
	// fmt.Printf("=== Create a non buffered input port %p\n",o)
	return ToPort(o.Id())
}

// create a port for writing
func MakeOutPort(f *os.File) *ClairePort {
	o := makeFile(f)
	o.nChar = 0
	o.status = 1 // marker that we write !
	o.buffer = make([]byte, BUFSIZE)
	o.nChar = 0
	return o
}

// make a read port from an input string
func F_port_I_string(s *ClaireString) *ClairePort {
	o := new(ClairePort)
	o.Isa = C_port
	o.buffer = ([]byte)(s.Value)
	o.nEof = len(s.Value)
	o.nChar = 0
	o.firstc = ' '
	o.status = 3 // tells that this is an input port (read)
	return o
}

func E_port_I_string(s EID) EID { return EID{F_port_I_string(ToString(OBJ(s))).Id(), 0} }

// make a write port to a string
func F_port_I_void() *ClairePort {
	o := new(ClairePort)
	o.Isa = C_port
	o.buffer = make([]byte, BUFPORT)
	o.nEof = BUFPORT
	o.nChar = 0
	o.firstc = ' '
	o.status = 2 // status = 2 write port
	// fmt.Printf(">>>> create a write buffered port\n")
	return o
}

func E_port_I_void(void EID) EID { return EID{F_port_I_void().Id(), 0} }

// these two functions are imported in read.cl through direct import pattern
// read one char at at time (local to this module ??)
func (p *ClairePort) GetNext() {
	p.firstc = p.Getc()
}

func (p *ClairePort) CharInt() int {
	return (int)(p.firstc)
}

// old int version of character push back (to be read later)
func F_pushback_port(p *ClairePort, n int) {
	p.firstc = rune(n)
}

func E_pushback_port(p EID, n EID) EID {
	F_pushback_port(ToPort(OBJ(p)), INT(n))
	return EVOID
}

// ------- API Functions -------------------------------------------------------------------------

// declare a port to be used as an output
func (p *ClairePort) UseAsOutput() *ClairePort {
	if p != ClEnv.Cout {
		x := ClEnv.Cout
		ClEnv.Cout = p
		return x
	} else {
		return p
	}
}

func E_use_as_output(p EID) EID { return EID{ToPort(OBJ(p)).UseAsOutput().Id(), 0} }

// close a file  : to write later
func (p *ClairePort) Fclose() {
	// if this port was actually in used by ClEnv, we revert to standartd
	if ClEnv.Cout == p {ClEnv.Cout = claireStdout}
	if ClEnv.Cin == p {ClEnv.Cin = claireStdin}        // new in CLAIRE 4
	p.Close()
}

func E_fclose_port(p EID) EID {
	ToPort(OBJ(p)).Fclose()
	return EVOID
}

// buffered version
// uses the size to make a right size Read !
func (p *ClairePort) Getc() rune {
	if p.status == 3 {
	    return p.GetStringChar()
	} else if p.status == 4 {
		char, _, err := ToGoPort(p.Id()).reader.ReadRune()
	    // fmt.Printf("[%d]", (int)(char))  // debug reader : show each char
		if err == nil {
			return char
		} else {
			return CEOF
		}
	} else {
		n := p.nChar
		if p.nChar >= p.nEof {
			return CEOF
		} else if n == BUFSIZE { // get a new piece
			if p.total == p.size {
				return CEOF
			}
			p.getSlice()
			return p.getRune(p.buffer[0])
		} else {
			p.nChar = n + 1
			return p.getRune(p.buffer[n])
		}
	}
}

// this function is exported to the interpreter
func E_getc_port(p EID) EID {
	// c := ToPort(OBJ(p)).Getc()
	// fmt.Printf("readChar(%p) -> %c\n",p,c)
	// return EID{C__CHAR, CVAL(c)}
	return EID{C__CHAR, CVAL(ToPort(OBJ(p)).Getc())}
}

// when we hit the end of the buffer p.getSlice gets the next slice (into the buffer)
// we look at the size of the file to read either a regular slice or the exact end
func (p *ClairePort) getSlice() {
	nextTotal := p.total + BUFSIZE
	if nextTotal > p.size {
		// fmt.Println("!!!! read last slice, s =",p.size - p.total)
		p.nEof = p.size - p.total
		p.buffer = p.buffer[0 : p.size-p.total]
	}
	p.file.Read(p.buffer)
	p.total = nextTotal
	//  fmt.Println("read slice size:",n,": ",string(p.buffer))
	p.nChar = 1
}

// extract a rune that starts with byte b ... can be complex if we slice at the wrong place
// this is the complex function that makes getc work :)
func (p *ClairePort) getRune(b byte) rune {
	if b < 127 {
		return rune(b)
	} else { // try to read with the next 4 bytes)
		// fmt.Println(">>>>>>>>>>> attempt to read Rune <<<<<<<<<<<<<<<<<<<<<")
		b1 := make([]byte, 4) // using local and short buffer b1
		b1[0] = b
		rSize := minInt(4, BUFSIZE+1-p.nChar) // read 4 chars up to end
		fmt.Println("-- read up to rSize = ", rSize)
		for i := 0; i < rSize-1; i++ {
			b1[i+1] = p.buffer[p.nChar+i]
		}
		// fmt.Println(">>>>> buffer =", b1)
		r, size := utf8.DecodeRune(b1)
		if r != utf8.RuneError {
			fmt.Println("rune:", string(r), "code:", (int)(r), " size:", size)
			p.nChar = p.nChar + (size - 1) // don't read the rune char twice
			return r
		} else { // we need to read more chars; rSize was too small
			remain := 4 - rSize
			if remain == 0 {
				fmt.Println("b1 =", string(b1))
				panic("shit reading runes")
			}
			p.getSlice()
			for i := 0; i < remain; i++ {
				b1[rSize+i] = p.buffer[i]
			}
			// fmt.Println("<<<<<<< buffer =",b1)
			r, size := utf8.DecodeRune(b1)
			if r != utf8.RuneError { // found rune with size ...
				// fmt.Println("second rune:",string(r),"code:",(int)(r)," size:",size)
				p.nChar = (size - rSize) // don't read the rune char twice
				return r
			} else {
				panic("shit reading hard runes")
			}
		}
	}
}

// read from the string (similar code)
func (p *ClairePort) GetStringChar() rune {
	n := p.nChar
	if n >= p.nEof {
		return CEOF
	} else {
		b := p.buffer[n]
		if b < 128 {
			p.nChar = n + 1
			return rune(b)
		} else {
			r, size := utf8.DecodeRune(p.buffer[n:minInt(n+3, p.nEof-1)])
			if r != utf8.RuneError { // found rune with size ...
				p.nChar = n + size // don't read the rune char twice
				return r
			} else {
				panic("shit reading hard runes")
			}
		}
	}
}

// =================================== WRITE ========================================

// because there are no exception in go, we avoid write errors by dynamic sizing of buffer string

// buffered put methods
func (p *ClairePort) Putc(c rune) {
	if p.status == 2 { // write to string
		p.PutStringChar(c)
	} else if p.file == os.Stdout {
		p.file.Write([]byte(string(c))) // no buffering for stdout
	} else {                            // p.status = 1
		n := p.nChar
		// fmt.Printf("putc(%c) on %p (n = %d)\n",c,p,n)
		if n == BUFSIZE {
			p.file.Write(p.buffer)
			n = 0
		}
		if c < 127 { // direct case
			p.buffer[n] = (byte)((int)(c))
			p.nChar = n + 1
		} else {
			p.putRune(c)
		}
	}
}


// this function is exported to the interpreter (possible name conflict -> Putc)
func E_putc_char (c EID, p EID) EID {
	ToPort(OBJ(p)).Putc(CHAR(c))
	return EVOID
}


// this is the special case for a Rune (complex char)
func (p *ClairePort) putRune(c rune) {
	b1 := make([]byte, 4)
	n := utf8.EncodeRune(b1, c)
	remain := BUFSIZE - p.nChar
	for i := 0; i < minInt(n, remain); i++ {
		p.buffer[p.nChar+i] = b1[i]
	}
	if n <= remain {
		p.nChar = p.nChar + n
	} else {
		p.file.Write(p.buffer)
		for i := remain; i < n; i++ {
			p.buffer[i-remain] = b1[i]
		}
		p.nChar = n - remain
	}
}

// write a char into the string buffer
func (p *ClairePort) PutStringChar(c rune) {
	n := p.nChar
	if n+3 > p.nEof { // resize
		m := p.nEof
		b2 := make([]byte, 2*m)
		copy(b2[0:n-1], p.buffer)
		p.buffer = b2
		p.nEof = 2 * m
	}
	if c < 127 { // direct case
		p.buffer[n] = (byte)((int)(c))
		p.nChar = n + 1
	} else {
		size := utf8.EncodeRune(p.buffer[n:n+3], c)
		p.nChar = n + size
	}
}

// ¨utString uses PutChar, hence it works for all ports
func (p *ClairePort) PutString(s string) {
	for _, c := range s {
		p.Putc(c)
	}
}

// solution with Fprintf
func (p *ClairePort) PutInteger(n int) { p.PutString(strconv.Itoa(n)) }

// used by princ @ float
func (p *ClairePort) PutFloat(x float64) {
	p.PutString(strconv.FormatFloat(x, 'g', -1, 64))
}

// used by print @ float
// pseudo-integers are printed with .0 to differentiate from int
func (p *ClairePort) PrettyFloat(x float64) {
	// fmt.Printf("--- call pretty float ----- \n")
	if float64(int(math.Round(x))) == x {
		p.PutString(fmt.Sprintf("%.1f", x))
    } else {p.PutString(strconv.FormatFloat(x, 'g', -1, 64))}
}

// used when we ask for a specific number of digit
func (p *ClairePort) PutFormat(x float64, i int) {
	format := "%." + fmt.Sprintf("%df", i)
	p.PutString(fmt.Sprintf(format,x))
}

// flush (for write ports)
func F_flush_port(p *ClairePort) {
	// fmt.Println("call flush with chars",p.nChar)
	if p.nChar > 0 {
		p.buffer = p.buffer[0:p.nChar]
		p.file.Write(p.buffer)
	}
}

func E_flush_port(p EID) EID {
	F_flush_port(ToPort(OBJ(p)))
	return EVOID
}

// close => flush
func (p *ClairePort) Close() {
	if p.status == 1 {
		F_flush_port(p)
	}
	p.file.Close()
}


// create an input/output port à la UNIX ("r","w","a")
// notice that "os.O_APPEND | os.O_CREATE | os.O_WRONLY" is absolutely necessary (stack overflow, don't see why)
func F_fopen_string(name *ClaireString, mode *ClaireString) EID {
	if mode.Value == "r" { // read file
		f, err := os.Open(name.Value)
		if err != nil {return Cerror(36,name.Id(),MakeInteger(0).Id())}
		return EID{MakeInPort(f).Id(),0}
	} else if mode.Value == "w" {
		f, err := os.Create(name.Value)
		// os.OpenFile(name,os.O_CREATE, 0644)
		if err != nil {return Cerror(36,name.Id(),MakeInteger(0).Id())}
		return EID{MakeOutPort(f).Id(),0}
	} else if mode.Value == "a" {
		f, err := os.OpenFile(name.Value, os.O_APPEND | os.O_CREATE | os.O_WRONLY , 0600)
		if err != nil { fmt.Printf("=== open file %s is rejected in append mode \n",name.Value)
			            return Cerror(39,name.Id(),MakeInteger(0).Id())}
		return EID{MakeOutPort(f).Id(),0}
	} else {
		panic("file open mode unknown: " + name.Value)
	}
}

func E_fopen_string(name EID, mode EID) EID {
	return F_fopen_string(ToString(OBJ(name)), ToString(OBJ(mode)))
}

// other API functions
// returns the string associated to the port
func (p *ClairePort) String_I() *ClaireString {
	if p.status == 2 {
		n := p.nChar
		s := string(p.buffer[0:n])
		// fmt.Printf(">>>>  buffered string port has %d chars\n", s)
		return MakeString(s)
	} else {
		return MakeString("")
	}
}

func E_string_I_port(p EID) EID {
	return EID{ToPort(OBJ(p)).String_I().Id(), 0}
}

// this function returns the number of chars in the buffer
func (p *ClairePort) Length() int {
	if p.status == 2 {
		return p.nChar
	} else {
		return 0
	}
}

func E_length_port(p EID) EID {
	return EID{C__INT, IVAL(ToPort(OBJ(p)).Length())}
}

// sets the buffer length to a certain number
// this is crucial to reuse a string port for multiple use !
func (p *ClairePort) SetLength(m int) {
	if p.status == 2 && m >= 0 && m <= p.nChar {
		p.nChar = m
	}
}

func E_set_length_port(p EID, m EID) EID {
	ToPort(OBJ(p)).SetLength(INT(m))
	return EVOID
}
