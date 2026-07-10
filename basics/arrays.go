package main

import (
	"fmt"
)

/*Arrays
The type [n]T is an array of n values of type T.

The expression

var a [10]int
declares a variable a as an array of ten integers.

An array's length is part of its type, so arrays cannot be resized. This seems limiting, but don't worry; Go provides a convenient way of working with arrays. */

func arrayDemo() {
	var aa [2]string
	aa[0] = "Hello"
	aa[1] = "World!"
	fmt.Println(aa)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	/*Slices
	An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays.
	The type []T is a slice with elements of type T.
	A slice is formed by specifying two indices, a low and high bound, separated by a colon*/
	s := primes[1:4]
	fmt.Println(s)

	/*Slices are like references to arrays
	A slice does not store any data, it just describes a section of an underlying array.
	Changing the elements of a slice modifies the corresponding elements of its underlying array.
	Other slices that share the same underlying array will see those changes. */

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)

	/*Slice defaults
	When slicing, you may omit the high or low bounds to use their defaults instead.
	The default is zero for the low bound and the length of the underlying slice or array for the high bound.
	For the array
	var a [10]int
	these slice expressions are equivalent:
	a[0:10]
	a[:10]
	a[0:]
	a[:] */

	/*Slice length and capacity
	A slice has both a length and a capacity.
	The length of a slice is the number of elements it contains.
	The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
	The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).
	You can extend a slice's length by re-slicing it, provided it has sufficient capacity. */

	sl := []int{2, 3, 5, 7, 11, 13}
	printSlice(sl)

	// Slice the slice to give it zero length.
	sl = sl[:0]
	printSlice(sl)

	// Extend its length.
	sl = sl[:4]
	printSlice(sl)

	// Drop its first two values.
	sl = sl[2:]
	printSlice(sl)

	// The zero value of slice is nil. has a length and capacity 0 and has no underlying array

	/*Creating a slice with make
	Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.
	The make function allocates a zeroed array and returns a slice that refers to that array:
	a := make([]int, 5)  // len(a)=5 */

	/** Appending to a slice
	It is common to append new elements to a slice, and so Go provides a built-in append function. The documentation of the built-in package describes append.
	func append(s []T, vs ...T) []T
	The first parameter s of append is a slice of type T, and the rest are T values to append to the slice.
	The resulting value of append is a slice containing all the elements of the original slice plus the provided values.
	If the backing array of s is too small to fit all the given values a bigger array will be allocated. The returned slice will point to the newly allocated array.*/

	sl = append(sl, 14, 15, 16)
	printSlice(sl)

	/*Range
	The range form of the for loop iterates over a slice or map.
	When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index. */

	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	powOfTwo := make([]int, 10)
	for i := range powOfTwo {
		powOfTwo[i] = 1 << uint(i)
	}
	for _, value := range powOfTwo {
		fmt.Printf("%d\n", value)
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

/*Implement Pic. It should return a slice of length dy, each element of which is a slice of dx 8-bit unsigned integers. When you run the program, it will display your picture, interpreting the integers as grayscale (well, bluescale) values.

The choice of image is up to you. Interesting functions include (x+y)/2, x*y, and x^y.

(You need to use a loop to allocate each []uint8 inside the [][]uint8.)

(Use uint8(intValue) to convert between types.) */

func Pic(dx, dy int) [][]uint8 {
	img := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		row := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			row[x] = uint8((x + y) / 2)
			//row[x] = uint8(x * y)
			//row[x] = uint8(math.Pow(float64(x), float64(y)))
		}
		img[y] = row
	}
	return img
}
