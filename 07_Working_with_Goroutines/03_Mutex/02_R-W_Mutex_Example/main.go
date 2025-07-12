/*
Differentiates between read and write access.
Multiple readers can access shared data concurrently, but writes are exclusive.
Methods:
	RLock(): Acquires a read lock.
	RUnlock(): Releases a read lock.
	Lock(): Acquires a write lock.
	Unlock(): Releases a write lock.
*/

package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	rwMutex sync.RWMutex
	wg      sync.WaitGroup
)

func readCounter() {
	defer wg.Done()
	rwMutex.RLock()
	fmt.Printf("Reading Counter: %d\n", counter)
	rwMutex.RUnlock()
}

func writeCounter() {
	defer wg.Done()
	rwMutex.Lock()
	counter++
	rwMutex.Unlock()
}

func main() {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go writeCounter()
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go readCounter()
	}

	wg.Wait()
}
