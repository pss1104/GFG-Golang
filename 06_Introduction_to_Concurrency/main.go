// buffered and unbuffered channels examples with basic go routines
// bufferd channel - allows sending multiple messages without waiting for a receiver
// unbuffered channel - requires a receiver to be ready to receive the message immediately
// go routines - lightweight threads managed by the Go runtime
// channels - used for communication between goroutines

package main

import (
	"fmt"
	"time"
)

func sendMessage(ch chan string) {
	ch <- "Hello from unbuffered channel"
}

func main() {
	// Unbuffered channel example
	unbufferedChannel := make(chan string)

	go sendMessage(unbufferedChannel)

	fmt.Println(<-unbufferedChannel)

	// Buffered channel example
	bufferedChannel := make(chan string, 2) //channel with a buffer size of 2

	bufferedChannel <- "Hello from buffered channel 1"
	bufferedChannel <- "Hello from buffered channel 2"

	fmt.Println(<-bufferedChannel)
	fmt.Println(<-bufferedChannel)

	time.Sleep(1 * time.Second) // Wait for goroutines to finish
}
