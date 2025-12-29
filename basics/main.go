package main

import (
	"fmt"
	"math"
)

func main() {
	//functions.go
	fmt.Println(add(42, 13))

	//functions-continued
	fmt.Println(add(42, 13))

	//multiple-results
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	//exported-names
	fmt.Println(math.Pi)
}
