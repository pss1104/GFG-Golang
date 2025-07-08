// basic data types example
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// Integer types
	var a int = 10
	var b int64 = 20

	// Floating-point types
	var c float32 = 3.14
	var d float64 = 2.71828

	// Boolean type
	var e bool = true

	// String type
	var f string = "Hello, Go!"

	//printing data types and their values along with their types
	fmt.Printf("Type of a: %T, Type of b: %T\n", a, b)
	fmt.Printf("Type of c: %T, Type of d: %T\n", c, d)
	fmt.Printf("Type of e: %T, Type of f: %T\n", e, f)

	fmt.Println("\nInteger a:", a)
	fmt.Println("Integer b:", b)
	fmt.Println("Float c:", c)
	fmt.Println("Float d:", d)
	fmt.Println("Boolean e:", e)
	fmt.Println("Size of int:", unsafe.Sizeof(a), "bytes")

}
