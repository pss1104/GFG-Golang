/*
Functions program

Syntax of function in Go with multiple parameters and return types:
func functionName(parameter1 type, parameter2 type, ...) (returnType1, returnType2, ...) {
	// function body
	return value1, value2 // if multiple return types are specified
}
*/

package main

import (
	"fmt"
)

// Simple function
func greet() {
	fmt.Println("Hello, Go!")
}

// Function with parameters
func add(a int, b int) int {
	return a + b
}

// Variadic function - can accept a variable number of arguments
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

// Function closure
// A closure is a special type of anonymous function that captures and remembers variables from its surrounding scope, even after the enclosing function has returned.
func createCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// Recursive function example
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// Function with multiple return types
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

// Function with named return values - swapping two integers
func swap(x, y int) (a int, b int) {
	a = y
	b = x
	return
}

func main() {
	// Calling simple function
	greet()

	// Calling function with parameters
	result := add(5, 10)
	fmt.Println("Sum of 5 and 10 is:", result)

	// Calling variadic function
	total := sum(1, 2, 3, 4, 5)
	fmt.Println("Sum of numbers is:", total)

	// Anonymous function -  unnamed functions that are typically defined inline
	af := func() {
		fmt.Println("This is an anonymous function!")
	}
	af() // Calling the anonymous function

	// Function closure example
	counter := createCounter()
	fmt.Println("Counter:", counter())
	fmt.Println("Counter:", counter())
	fmt.Println("Counter:", counter())

	// Calling the recursive function
	num := 5
	fmt.Printf("Factorial of %d is: %d\n", num, factorial(num))

	// Calling function with multiple return types
	quotient, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 divided by 2 is:", quotient)
	}

	// Calling function with named return values - swapping two integers
	x, y := 10, 20
	a, b := swap(x, y)
	fmt.Printf("After swapping: x = %d, y = %d\n", a, b)

	af() // Calling the anonymous function again to demonstrate it can be called multiple times
}
