package main

import "math"

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
