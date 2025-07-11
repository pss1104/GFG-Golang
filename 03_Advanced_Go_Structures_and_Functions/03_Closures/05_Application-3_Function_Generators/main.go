package main

import "fmt"

func multiplier(factor int) func(int) int {
	return func(value int) int {
		return factor * value
	}
}

func main() {
	double := multiplier(2)
	triple := multiplier(3)

	fmt.Println(double(4)) // Output: 8
	fmt.Println(triple(4)) // Output: 12
}
