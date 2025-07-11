package main

import "fmt"

func counter() (func() int, func() int) {
	x := 0
	increment := func() int {
		x++
		return x
	}
	decrement := func() int {
		x--
		return x
	}
	return increment, decrement
}

func main() {
	inc, dec := counter()

	fmt.Println(inc()) // Output: 1
	fmt.Println(inc()) // Output: 2
	fmt.Println(dec()) // Output: 1
	fmt.Println(dec()) // Output: 0
}
