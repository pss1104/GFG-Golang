package main

import "fmt"

func adder() func(int) int {
	x := 5
	return func(y int) int {
		return x + y
	}
}

func main() {
	add := adder()
	fmt.Println(add(3)) // Output: 8
	fmt.Println(add(5)) // Output: 10
}
