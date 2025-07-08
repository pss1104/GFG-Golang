// type inference example
package main

import "fmt"

func main() {
	// Type inference in Go
	var a = 10 // inferred as int
	// var b = 3.14       // inferred as float64
	b := 3.14
	var c = "Hello"      // inferred as string
	var d = true         // inferred as bool
	var e = 'a'          // inferred as rune (int32)
	var f = "Hello, Go!" // inferred as string

	fmt.Printf("Type of a: %T, Value: %v\n", a, a)
	fmt.Printf("Type of b: %T, Value: %v\n", b, b)
	fmt.Printf("Type of c: %T, Value: %v\n", c, c)
	fmt.Printf("Type of d: %T, Value: %v\n", d, d)
	fmt.Printf("Type of e: %T, Value: %v\n", e, e) //prints int32, rune is just an alias for int32
	fmt.Printf("Type of f: %T, Value: %v\n", f, f)

	// Type inference with composite types
	slice := []int{1, 2, 3} // inferred as []int
	fmt.Printf("Type of slice: %T, Value: %v\n", slice, slice)

	// := operator for short variable declaration
	x, y := 5, 10 // x and y are inferred as int
	fmt.Printf("Type of x: %T, Value: %v\n", x, x)
	fmt.Printf("Type of y: %T, Value: %v\n", y, y)

	// var x,y= 5, 10 // This also works
}

/*
Note: := operator is used for short variable declaration and can only be used inside functions.
It cannot be used at the package level.
Type inference allows the Go compiler to automatically determine the type of a variable based on its initial value
*/
