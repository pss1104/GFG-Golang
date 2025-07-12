// By default, channels in Go remain open until explicitly closed.
// Closing a channel signals to its receivers that no more data will be sent.

package main

import "fmt"

// Sender function sends values to the channel and then closes it
func sender(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i // Send values to the channel
	}
	close(ch) // Close the channel after sending all values
}

func main() {
	ch := make(chan int) // Create a channel for communication

	go sender(ch) // Start the sender goroutine

	// Attempt to receive 8 values from the channel
	// The channel sends only 5 values, so the remaining receives will return the zero value (0)
	for i := 0; i < 8; i++ {
		fmt.Println(<-ch) // Receive values or zero value if channel is empty
	}

	fmt.Println("Channel is closed.") // Indicate that processing is complete
}
