package main

import "fmt"

/*
Function closures
Go functions may be closures. A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.
For example, the adder function returns a closure. Each closure is bound to its own sum variable.
*/
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

/*
Function values
Functions are values too. They can be passed around just like other values.
Function values may be used as function arguments and return values.
*/
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

// fibonacci is a function that returns
// a function ( a closure) that returns an int.
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		value := a
		a, b = b, a+b
		return value
	}
}

func funcDemo() {
	pos, neg := adder(), adder()
	f := fibonacci()
	for i = 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
		fmt.Println(f())
	}
}
