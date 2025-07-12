package main

import (
	"fmt"
	"sync"
	"time"
)

// Worker function that processes tasks from a task channel and sends results on a result channel
func worker(id int, tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
	for task := range tasks {
		result := processTask(task)
		results <- result
		fmt.Printf("Worker %d processed task %d\n", id, task)
	}
	wg.Done()
}

// Simulates task processing
func processTask(task int) int {
	time.Sleep(time.Second) // Simulate time-consuming task
	return task * 2
}

func main() {
	const numWorkers = 3
	const numTasks = 5

	tasks := make(chan int, numTasks)   // Task channel (receive-only for workers)
	results := make(chan int, numTasks) // Result channel (send-only for workers)
	var wg sync.WaitGroup

	// Start worker goroutines
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, tasks, results, &wg)
	}

	// Send tasks to the worker pool
	for j := 1; j <= numTasks; j++ {
		tasks <- j
	}
	close(tasks) // Close tasks channel when no more tasks are left to send

	// Wait for all workers to finish
	wg.Wait()
	close(results) // Close results channel after workers are done

	// Print results
	for result := range results {
		fmt.Println("Result:", result)
	}
}
