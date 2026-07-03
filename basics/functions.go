package main

import (
	"fmt"
	"math"
)

func add(x int, y int) int {
	return x + y
}

/*
 * If with a short statement
 Like for, the if statement can start with a short statement to execute before the condition.

 Variables declared by the statement are only in scope until the end of the if.
*/

func pow(x, n, limit float64) float64 {
	if v := math.Pow(x, n); v < limit {
		return v
	}
	return limit
}

/*Defer
A defer statement defers the execution of a function until the surrounding function returns.

The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
*/

func deferDemo() {
	defer fmt.Println("world.")
	fmt.Println("Hello")
}

/*
 * Stacking defers
 Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.
*/

func deferStackDemo() {
	fmt.Println("Starting Counting...")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("Counting Done.")
}
