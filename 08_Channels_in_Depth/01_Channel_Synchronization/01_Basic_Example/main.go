/*
| Feature               | **Channels**                                        | **Locks (`sync.Mutex`)**                           |
| --------------------- | --------------------------------------------------- | -------------------------------------------------- |
| **Purpose**           | Communication between goroutines                    | Protection of shared data                          |
| **Style**             | Message-passing concurrency                         | Shared-memory concurrency                          |
| **Preferred use**     | When goroutines need to coordinate or exchange data | When goroutines access the same variable or struct |
| **Built-in blocking** | Yes — `chan <- value` blocks if no receiver         | No — you must explicitly `Lock()` and `Unlock()`   |
| **Deadlock risk**     | If not carefully structured (e.g., no receiver)     | If `Unlock()` is forgotten or misused              |
| **Go philosophy**     | "Communicate by sharing memory"                     | "Protect shared memory with mutual exclusion"      |

*/

package main

import (
	"fmt"
	"time"
)

func worker(id int, ch chan int) {
	fmt.Printf("Worker %d: Sending data...\n", id)
	ch <- id // Send data to the channel
	fmt.Printf("Worker %d: Sent data\n", id)
}

func main() {
	ch := make(chan int) // Create an unbuffered channel
	// ch := make(chan int, 3) // Create a buffered channel

	// Start three worker goroutines
	for i := 1; i <= 3; i++ {
		go worker(i, ch)
	}

	// Receive data from the channel three times
	for i := 1; i <= 3; i++ {
		val := <-ch // Receive data from the channel
		fmt.Printf("Main: Received data %d from channel\n", val)
	}

	time.Sleep(2 * time.Second) // Wait for all goroutines to finish
}
