package main

import (
	"fmt"
	"math"
)

/*Interfaces
An interface type is defined as a set of method signatures.
A value of interface type can hold any value that implements those methods */

type Abser interface {
	//methods signatures...
	Abs() float64
}

/*
Interfaces are implemented implicitly
A type implements an interface by implementing its methods. There is no explicit declaration of intent, no "implements" keyword.
Implicit interfaces decouple the definition of an interface from its implementation, which could then appear in any package without prearrangement.
*/

type In interface {
	M()
}

type Ty struct {
	S string
}

func (t Ty) M() {
	fmt.Println(t.S)
}

/*Interface values
Under the hood, interface values can be thought of as a tuple of a value and a concrete type:
(value, type)
An interface value holds a value of a specific underlying concrete type.
Calling a method on an interface value executes the method of the same name on its underlying type. */

/*type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
*/

/*Interface values with nil underlying values
If the concrete value inside the interface itself is nil, the method will be called with a nil receiver.
In some languages this would trigger a null pointer exception, but in Go it is common to write methods that gracefully handle being called with a nil receiver (as with the method M in this example.)
Note that an interface value that holds a nil concrete value is itself non-nil. */

/*
	type I interface {
		M()
	}

	type T struct {
		S string
	}

	func (t *T) M() {
		if t == nil {
			fmt.Println("<nil>")
			return
		}
		fmt.Println(t.S)
	}

	func main() {
		var i I

		var t *T
		i = t
		describe(i)
		i.M()

		i = &T{"hello"}
		describe(i)
		i.M()
	}

	func describe(i I) {
		fmt.Printf("(%v, %T)\n", i, i)
	}
*/

/*
Nil interface values
A nil interface value holds neither value nor concrete type.

Calling a method on a nil interface is a run-time error because there is no type inside the interface tuple to indicate which concrete method to call.
*/

/*
	type I interface {
		M()
	}

	func main() {
		var i I
		describe(i)
		i.M()
	}

	func describe(i I) {
		fmt.Printf("(%v, %T)\n", i, i)
	}
*/

/*
The empty interface
The interface type that specifies zero methods is known as the empty interface:
interface{}
An empty interface may hold values of any type. (Every type implements at least zero methods.)
Empty interfaces are used by code that handles values of unknown type. For example, fmt.Print takes any number of arguments of type interface{}.
*/

/*
	func main() {
		var i interface{}
		describe(i)

		i = 42
		describe(i)

		i = "hello"
		describe(i)
	}
*/

/*Type assertions
A type assertion provides access to an interface value's underlying concrete value.

t := i.(T)
This statement asserts that the interface value i holds the concrete type T and assigns the underlying T value to the variable t.

If i does not hold a T, the statement will trigger a panic.

To test whether an interface value holds a specific type, a type assertion can return two values: the underlying value and a boolean value that reports whether the assertion succeeded.

t, ok := i.(T)
If i holds a T, then t will be the underlying value and ok will be true.

If not, ok will be false and t will be the zero value of type T, and no panic occurs.

Note the similarity between this syntax and that of reading from a map. */

func interfacesDemo() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex2{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	//a = v

	fmt.Println(a.Abs())

	//var i In = Ty{"hello"}
	//i.M()

	//Type Assertions
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	ff, ok := i.(float64)
	fmt.Println(ff, ok)

	//ff = i.(float64) // panic here
	//fmt.Println(ff)
}
