//channels and pointers are also composite data types
// initializing channels and pointers
// channels are used for communication between goroutines
// pointers are used to store the address of a variable

package main

import (
	"fmt"
)

func main() {
	// Channel
	ch := make(chan string)
	go func() {
		ch <- "Hello, Channel!"
	}()
	fmt.Println("Channel:", <-ch)
	// printing the type of channel
	fmt.Printf("Type of ch: %T\n", ch)

	// Pointer
	var x int = 42
	var p *int = &x                        // p is a pointer to x
	fmt.Println("Pointer:", *p)            // dereferencing the pointer to get the value of x
	fmt.Println("Address of x:", &x)       // printing the address of x
	fmt.Println("Address stored in p:", p) // printing the address stored in p
	//printing the type of pointer
	fmt.Printf("Type of p: %T\n", p)

	// Pointers can be used to modify the value of a variable
	*p = 100
	fmt.Println("Modified value of x:", x)
}
