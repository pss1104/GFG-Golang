package main

import "fmt"

// Higher-order function that returns a function
func multiplier(factor int) func(int) int {
	return func(value int) int {
		return factor * value
	}
}

func main() {
	double := multiplier(2) // Returns a function to double the input
	triple := multiplier(3) // Returns a function to triple the input

	fmt.Println("Double of 5:", double(5)) // Output: Double of 5: 10
	fmt.Println("Triple of 5:", triple(5)) // Output: Triple of 5: 15
}
