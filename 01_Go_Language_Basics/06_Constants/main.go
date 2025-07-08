// syntax for constants in Go: const name type = value
//2 types of constants in Go: 1. Untyped and 2. Typed

package main

import (
	"fmt"
	"math"
)

func main() {
	// Untyped constant
	const pi = 3.14 // untyped float constant
	fmt.Printf("Type of pi: %T, Value: %v\n", pi, pi)

	// Typed constant
	const e float64 = math.E // typed float constant
	fmt.Printf("Type of e: %T, Value: %v\n", e, e)

	// Constant Expressions - Constants can be used in expressions
	const radius = 5.0
	area := pi * radius * radius // using untyped constant in expression
	fmt.Printf("Area of circle with radius %.2f: %.2f\n", radius, area)

	//Enumerated constants
	const (
		Sunday = iota // iota starts at 0
		Monday        // iota increments by 1 for each constant
		Tuesday
		Wednesday
	)
	fmt.Printf("Days of the week: %d, %d, %d, %d\n", Sunday, Monday, Tuesday, Wednesday)

	// Advanced usage of iota
	const (
		_  = iota             // ignore the first value (0)
		KB = 1 << (10 * iota) // 1 << 10 = 1024
		MB                    // 1 << 20 = 1048576
		GB                    // 1 << 30 = 1073741824
	)
	fmt.Printf("1 KB = %d bytes, 1 MB = %d bytes, 1 GB = %d bytes\n", KB, MB, GB)

	// type conversion with constants
	const x = 10           // Untyped constant
	var y int64 = int64(x) // Convert to int64
	fmt.Println(y)
}
