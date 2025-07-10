// In Go, defer delays the execution of a function until the surrounding function completes.
// Regardless of whether the function ends normally, encounters an error, or panics, the deferred statements are always executed.
// It is often used for cleanup tasks, such as closing files or releasing resources.
// The deferred function calls are executed in LIFO (Last In, First Out) order. (Stack is used to store the deferred function calls.)
// Can also be used to unlock mutexes or to log messages after a function execution.

package main

import (
	"fmt"
	"log"
	"os"
	"sync"
)

func defer_with_for() {
	// Defer statement in a loop
	for i := range 5 {
		defer fmt.Println(i)
	}
	fmt.Println("Loop completed")
}

// use case 1 - defer to close a file
func readFile() {
	file, err := os.Open("temp.txt")
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	// Perform operations on the file
	log.Println("File opened successfully")
}

// use case 2 - defer to unlock a mutex

var mu sync.Mutex

func accessResource() {
	mu.Lock()
	defer mu.Unlock()
	fmt.Println("Accessing shared resource")
	// Safely access shared resource
}

// use case - 3 modifying return values
func modifyReturnValue() (result int) {
	defer func() {
		result += 10 // Modify the return value before returning
	}()
	return 5 // Initial return value
}

// use case 4 - defer to log messages after function (in this case - main function) execution
func main() {
	defer fmt.Println("main 1")
	fmt.Println("main 2")
	defer fmt.Println("main3")

	defer_with_for()
	readFile()
	accessResource()
	modifiedValue := modifyReturnValue()
	fmt.Println("Modified return value:", modifiedValue)
}
