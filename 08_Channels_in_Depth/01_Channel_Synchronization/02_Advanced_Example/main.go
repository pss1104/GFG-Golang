package main

import (
	"fmt"
)

func worker(id int, ch chan int) {
	for j := 0; j < 3; j++ {
		fmt.Printf("Worker %d: Sending data %d...\n", id, j)
		ch <- (id * 10) + j // Send data to the channel
	}
}

func main() {

	ch := make(chan int)

	// Start three worker goroutines
	for i := 1; i <= 3; i++ {
		go worker(i, ch)
	}

	// Receive data from the channel nine times (3 workers * 3 pieces of data)
	for i := 1; i <= 9; i++ {
		fmt.Printf("Main: Received data %d from channel\n", <-ch)
	}
}
