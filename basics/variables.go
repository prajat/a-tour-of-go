/*
 Variables
 The var statement declares a list of variables; as in function argument lists, the type is last.

 A var statement can be at package or function level. We see both in this example.
*/

package main

//var c, python, java bool

/*
Variables with initializers
A var declaration can include initializers, one per variable.

If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
*/

var i, j int = 1, 2
var c, python, java = true, false, "no!"

/*
Short variable declarations
Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.

Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.
*/
