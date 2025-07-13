package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done() // Defer to ensure Done() is called when worker finishes

	for j := range jobs {
		fmt.Printf("worker %d started job %d\n", id, j)
		time.Sleep(time.Second) // Simulating work
		fmt.Printf("worker %d finished job %d\n", id, j)
		results <- j * 2 // Send the result back to results channel
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)    // Task channel
	results := make(chan int, numJobs) // Result channel

	var wg sync.WaitGroup

	// Launching 3 workers.
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Sending 5 jobs to the job channel
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // Close job channel to indicate no more jobs will be sent

	// Waiting for all workers to finish.
	wg.Wait()
	close(results) // Close result channel when all workers are done

	// Collecting results.
	for r := range results {
		fmt.Println("result:", r)
	}
}
