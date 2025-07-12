// A mutex is a synchronization primitive used to manage concurrent access to shared data.
// It ensures that only one goroutine can access a critical section of code at a time,
// preventing race conditions and maintaining data consistency.

// Short for Mutual Exclusion, a mutex is used to lock and unlock access to shared resources.

// Protects a critical section by ensuring only one goroutine can access it at a time.
// Methods:
// Lock(): Acquires the lock.
// Unlock(): Releases the lock.
package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	mutex   sync.Mutex
)

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()
	counter++
	mutex.Unlock()
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	wg.Wait()
	fmt.Printf("Final Counter Value: %d\n", counter)
}
