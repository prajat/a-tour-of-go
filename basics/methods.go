package main

import (
	"fmt"
	"math"
)

/*
Methods
Go does not have classes. However, you can define methods on types.

A method is a function with a special receiver argument.

The receiver appears in its own argument list between the func keyword and the method name.
*/
type Vertex2 struct {
	X, Y float64
}

func (v Vertex2) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/*
Methods are functions
Remember: a method is just a function with a receiver argument.
Here's Abs written as a regular function with no change in functionality.
*/

func Abs2(v Vertex2) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/*
 Methods continued
 You can declare a method on non-struct types, too.
 In this example we see a numeric type MyFloat with an Abs method.
 You can only declare a method with a receiver whose type is defined in the same package as the method.
 You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int).
*/

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

/*Pointer receivers
You can declare methods with pointer receivers.
This means the receiver type has the literal syntax *T for some type T. (Also, T cannot itself be a pointer such as *int.)
For example, the Scale method here is defined on *Vertex.
Methods with pointer receivers can modify the value to which the receiver points (as Scale does here). Since methods often need to modify their receiver, pointer receivers are more common than value receivers.
Try removing the * from the declaration of the Scale function on line 16 and observe how the program's behavior changes.
With a value receiver, the Scale method operates on a copy of the original Vertex value. (This is the same behavior as for any other function argument.) The Scale method must have a pointer receiver to change the Vertex value declared in the main function. */

func (v *Vertex2) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

/*Choosing a value or pointer receiver
There are two reasons to use a pointer receiver.
The first is so that the method can modify the value that its receiver points to.
The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct, for example.
In this example, both Scale and Abs are methods with receiver type *Vertex, even though the Abs method needn't modify its receiver.
In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both. (We'll see why over the next few pages.) */

func methodsDemo() {
	v := Vertex2{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(Abs2(v))

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	v.Scale(10)
	fmt.Println(v.Abs())
}
