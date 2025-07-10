// pointers in go
package main

import (
	"fmt"
)

func main() {
	// Pointer declaration and initialization
	var ptr *int // ptr is a pointer to an int, initially nil
	fmt.Println("Initial value of ptr:", ptr)

	// Pointer to a variable
	num := 42
	ptr = &num // ptr now points to num
	fmt.Println("Value of num:", num)
	fmt.Println("Address of num:", &num)
	fmt.Println("Value pointed by ptr:", *ptr)

	// Modifying the value using pointer
	*ptr = 100 // changes num to 100
	fmt.Println("Modified value of num:", num)

	// Call by value and reference examples
	x := 10
	fmt.Println("Before callByValue, x:", x)
	callByValue(x) // This will not change the original value
	fmt.Println("After callByValue, x:", x)
	fmt.Println("Before callByReference, x:", x)
	callByReference(&x) // This will change the original value
	fmt.Println("After callByReference, x:", x)

	pp := &ptr
	fmt.Println("Pointer to pointer value:", pp)
}

// call by value and reference examples
func callByValue(x int) {
	fmt.Println("Inside callByValue, original x:", x)
	x = 100 // This will not change the original value
	fmt.Println("Inside callByValue, modified x:", x)
}
func callByReference(x *int) {
	*x = 100 // This will change the original value
	fmt.Println("Inside callByReference, x:", *x)
}
