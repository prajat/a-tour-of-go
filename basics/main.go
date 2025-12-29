package main

import (
	"fmt"
	"math"
)

func main() {
	//functions
	fmt.Println(add(42, 13))

	//multiple-results
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	//exported-names
	fmt.Println(math.Pi)

	//naked-return
	fmt.Println(split(54))

	//variables
	var i int
	fmt.Println(i, c, python, java)
}
