package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("Received", msg1)
		case msg2 := <-c2:
			fmt.Println("Received", msg2)
		case <-time.After(500 * time.Millisecond):
			fmt.Println("timeout")
			// select with default is non-blocking.
			// default: //will always execute if no other case is ready
			// 	fmt.Println("No activity")
		}
	}
}

/*
Output:
timeout
Received one

Why?
Now select is blocking — it waits for one of the channels or time.After() to be ready.

Loop iteration 1:
Nothing is ready yet.

After 500ms, time.After fires → "timeout" printed.

Loop iteration 2:
time.Sleep(1s) from go func() { ... c1 <- "one" } has completed.

c1 <- "one" is ready → "Received one" printed.

| Time   | c1 status       | c2 status       | select#1             | select#2         |
| ------ | --------------- | --------------- | -------------------- | ---------------- |
| 0ms    | not ready       | not ready       | **default runs**     | **default runs** |
| 500ms  | still not ready | still not ready | timeout (no default) | --               |
| 1000ms | ready (c1)      | not ready       | --                   | receives from c1 |
| 2000ms | --              | ready (c2)      | --                   | --               |
*/
