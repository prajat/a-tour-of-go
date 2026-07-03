package main

import "fmt"

/*
 * A struct is a collection of fields.
 */

type Vertex struct {
	X int
	Y int
}

func vertex() {
	//fmt.Println(Vertex{1, 2})
	v := Vertex{1, 2}
	v.X = 4

	//Struct fields are accessed using a dot.
	fmt.Println(v.X, v.Y)

	/* Pointers to structs
	Struct fields can be accessed through a struct pointer.
	To access the field X of a struct when we have the struct pointer p we could write (*p).X. However, that notation is cumbersome, so the language permits us instead to write just p.X, without the explicit dereference.
	*/
	p := &v
	p.X = 1e9
	fmt.Println(v)

	/*Struct Literals
	A struct literal denotes a newly allocated struct value by listing the values of its fields.
	You can list just a subset of fields by using the Name: syntax. (And the order of named fields is irrelevant.)
	The special prefix & returns a pointer to the struct value. */
	v1 := Vertex{1, 2}
	v2 := Vertex{X: 1}
	v3 := Vertex{}
	v4 := &Vertex{1, 2}
	fmt.Println(v1, v2, v3, v4)
}
