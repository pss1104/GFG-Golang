package main

import "fmt"

// Higher-order function
func operate(a, b int, op func(int, int) int) int {
	return op(a, b)
}

// Functions to pass
func add(x, y int) int {
	return x + y
}

func multiply(x, y int) int {
	return x * y
}

func main() {
	// Using the higher-order function
	fmt.Println("Sum:", operate(3, 5, add))          // Output: Sum: 8
	fmt.Println("Product:", operate(3, 5, multiply)) // Output: Product: 15
}
