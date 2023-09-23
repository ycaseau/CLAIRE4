// +------------------------------------------------------------+
// | bug0.cl                                                    |
// | last update: Jan 2023 - Y. Caseau                          |
// +------------------------------------------------------------+


// ---------------------------------------------------------------
// this file contains code that will crash claire ... but with a nice error message
// ---------------------------------------------------------------
// by consruction each use case should be commented out when not tested


// first case: create an object class with too many names
myClass <: thing(
    a1:integer = 1,
    a2:integer = 2,
    a3:integer = 3,
    a4:integer = 4,
    a5:integer = 5,
    a6:integer = 6,
    a7:integer = 7,
    a8:integer = 8,
    a9:integer = 9,
    a10:integer = 10,
    a11:integer = 11,
    a12:integer = 12,
    a13:integer = 13,
    a14:integer = 14,
    a15:integer = 15,
    a16:integer = 16,
    a17:integer = 17,
    a18:integer = 18,
    a19:integer = 19,
    a20:integer = 20,
    a21:integer = 21,
    a22:integer = 22,
    a23:integer = 23,
    a24:integer = 24,
    a25:integer = 25,
    a26:integer = 26,
    a27:integer = 27,
    a28:integer = 28,
    a29:integer = 29)

(printf("myClass has been created\n"))

// big class with 50 attributes
myBigClass <: myClass(
    a30:integer = 30,
    a31:integer = 31,
    a32:integer = 32,
    a33:integer = 33,
    a34:integer = 34,
    a35:integer = 35,
    a36:integer = 36,
    a37:integer = 37,
    a38:integer = 38,
    a39:integer = 39,
    a40:integer = 40,
    a41:integer = 41,
    a42:integer = 42,
    a43:integer = 43,
    a44:integer = 44,
    a45:integer = 45,
    a46:integer = 46,
    a47:integer = 47,
    a48:integer = 48,
    a49:integer = 49,
    a50:integer = 50)


