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

	var i In = Ty{"hello"}
	i.M()
}
