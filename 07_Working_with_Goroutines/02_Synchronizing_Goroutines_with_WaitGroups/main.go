/*
WaitGroup Methods
1. Add(delta int)
	Increases the counter by the specified delta.
	Use this to specify the number of goroutines to wait for.
	Example: wg.Add(1)
2. Done()
	Decrements the counter by 1.
	Typically called using defer at the start of a goroutine to ensure it’s executed even if the goroutine exits prematurely.
3. Wait()
	Blocks the main function until the counter reaches zero.
	Should be called in the main function to prevent premature termination.
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

// Worker function simulating a task
func worker(id int, wg *sync.WaitGroup, delay time.Duration) {
	defer wg.Done() // Decrease counter at the end
	// This ensures that no matter how the function exits, wg.Done() is called.
	// If it were at the end without defer, and a panic or early return occurred, wg.Done() might never be called — leading to a deadlock (main waits forever at wg.Wait()).

	fmt.Printf("Worker %d: Starting\n", id)
	time.Sleep(delay) // Simulate work
	fmt.Printf("Worker %d: Completed\n", id)
}

func main() {
	var wg sync.WaitGroup

	// Number of workers
	numWorkers := 3

	// Add the number of workers to WaitGroup
	wg.Add(numWorkers) // Tells the WaitGroup to expect 3 goroutines to call wg.Done().

	// Start workers
	for i := 1; i <= numWorkers; i++ {
		go worker(i, &wg, time.Duration(i)*time.Second)
	}

	// Wait for all workers to complete
	wg.Wait()

	fmt.Println("All workers completed. Main process exits.")
}
