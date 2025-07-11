package main

import "fmt"

func eventHandler(initialValue int) func() int {
	value := initialValue
	return func() int {
		value++
		return value
	}
}

func main() {
	handler := eventHandler(10)
	fmt.Println(handler()) // Output: 11
	fmt.Println(handler()) // Output: 12
}
