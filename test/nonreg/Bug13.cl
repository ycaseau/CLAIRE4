// +------------------------------------------------------------+
// | bug13.cl                                                   |
// | last update: Jan 2000 - Y. Caseau                          |
// +------------------------------------------------------------+

// ---------------------------------------------------------------
// general test file for handling unknown
// ---------------------------------------------------------------

[A <: thing(x:integer, y:integer = 0, z:(integer U {unknown}) = 3, w:(integer U {unknown}))]

a1 :: A()
a2 :: A(x = 1, y = 1, z = 1, w = 1)

// use an assertion
// CLAIRE4: 0 is the default default value for integer
assert(get(x,a1) = 0 & y(a1) = 0 & z(a1) = 3 & w(a1) = unknown)        
(check("assert 1", (get(x,a1) = 0 & y(a1) = 0 & z(a1) = 3 & w(a1) = unknown)))

assert(x(a2) = 1 & y(a2) = 1 & z(a2) = 1 )

put(z,a2,unknown) 
put(w,a2,unknown)         

assert(get(z,a2) = unknown & get(z,a2) = unknown)

(x(a1) := 2, y(a1) := 2)

assert(x(a1) = 2 & y(a1) = 2)

// same test with store option
[A2 <: thing(X:integer, Y:integer = 0, Z:(integer U {unknown}) = 3, W:(integer U {unknown}) ) ]

store(X,Y,Z)

b1 :: A2()
b2 :: A2(X = 1, Y = 1, Z = 1, W = 1)

assert(get(X,b1) = 0 & Y(b1) = 0 & Z(b1) = 3 & W(b1) = unknown)         

assert(X(b2) = 1 & Y(b2) = 1 & Z(b2) = 1)

put(Z,b2,unknown)         
put(W,b2,unknown)         

(check("Z & W", get(Z,b2) = unknown & get(W,b2) = unknown))

(X(b1) := 2, Y(b1) := 2)

(check("X & Y", X(b1) = 2 & Y(b1) = 2))

// an old bug that is now (2.4.9) fixed.
[b <: thing]

[b <: thing(_next_b:b)]
[a <: thing(_first_b:b)]

[testX()
  -> let a1 := a(), b1 := b()
     in (put(_next_b, b1, get(_first_b, a1)),
	 if (known?(_next_b, b1))
	     (printf("Here's the problem! -> ~S\n",get(_next_b,b1)),
              record("testX has a problem",true))) ]

(testX())

// inverse - in CLAIRE4, this requires a set based slot (unique form of multivalued)
CB <: thing
CA <: thing(sbs:set[CB])
CB <: thing(sa:CA)

(inverse(sa) := sbs)

x1 :: CA()
y1 :: CB(sa = x1)
x2 :: CA()
y2 :: CB(sa = x2)

[testErase()
  -> erase(sa,y1),
     check("erase a",size(x1.sbs) = 0),
     erase(sbs,x2),
     check("erase bs",unknown?(sa,y2)) ]

(check("simple inverse",x1.sbs = set<CB>(y1)))
(testErase())

// a small exception bug from Arnaud

// exceptions - diagnostic is in diag
COMError <: exception(diag:string)
[self_print(e:COMError) : any -> printf("libCOMLAIRE Error : ~A",e.diag)]

// convert error number to text
// hresult is an HRESULT anded with 0x3FFFFFFF to get rid of extra
// high-level bits
// paramcount is a normal order, 1-based param number.
[raiseCOMError(hresult:integer,paramcount:integer) : void -> 
  COMError(diag = "GPF")]

[testCOM() : void
  -> try raiseCOMError(3,3)
     catch COMError nil ]

 (testOK())

