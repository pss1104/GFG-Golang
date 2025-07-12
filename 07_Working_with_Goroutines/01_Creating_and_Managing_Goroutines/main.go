// Go routine - a lightweight thread managed by the Go runtime
// goroutines are used to perform concurrent tasks
// go function() starts a new goroutine
// main(0) does not wait for goroutines to finish by default

package main

import (
	"fmt"
	"runtime"
	"time"
)

func printMessage(msg string) {
	for i := 0; i < 3; i++ {
		fmt.Println(msg)
		time.Sleep(300 * time.Millisecond)
	}
}
func main() {

	runtime.GOMAXPROCS(4) // Set number of OS threads to 1 for demonstration

	// Starting a goroutine
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("Goroutine:", i)
			time.Sleep(100 * time.Millisecond) // Simulate work
		}
	}()
	// Calling a function in a goroutine
	go printMessage("Hello from goroutine")

	// Main function continues executing - does not wait for goroutines
	for i := 0; i < 5; i++ {
		fmt.Println("Main:", i)
		time.Sleep(200 * time.Millisecond) // Simulate work
	}

	time.Sleep(1 * time.Second) // Wait for goroutine to finish before exiting
}
