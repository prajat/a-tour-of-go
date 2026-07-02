package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	//Functions
	fmt.Println(add(42, 13))

	//Functions Continued
	fmt.Println(add2(10, 90))

	//Multiple Results
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	//Exported Names
	fmt.Println(math.Pi)
	fmt.Println(rand.Intn(10))

	//Naked Return
	fmt.Println(split(54))

	//Variables
	//var i int
	//var c, python, java = true, false, "no!"

	//Short Variables Declaration
	k := 3
	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java)

	//Basic Types
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", Z, Z)

	//Zero Values
	fmt.Printf("%v %v %v %q\n", I, F, B, S)
	//fmt.Printf("%T %T %T %T\n", I, F, B, S)var x, y int = 3, 4

	//Type Conversions
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

	//Type Interface
	//v := 42.6 + 34.89i  // -> complex128
	//v := 42            // -> int
	v := 42.66    // -> float64
	K := int32(v) // -> int32
	fmt.Printf("v is of type %T, %T, value of K -> %v\n", v, K, K)

	//loops
	fmt.Println(forloop())
	fmt.Println(forcontinued())
	fmt.Println(forlikewhile())
}
