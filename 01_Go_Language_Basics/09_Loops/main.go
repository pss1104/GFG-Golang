// loops in go
package main

import "fmt"

func main() {
	// For Loop
	fmt.Println("For Loop:")
	for i := 0; i < 5; i++ {
		fmt.Println("Iteration:", i)
	}

	// While Loop (using for)
	fmt.Println("\nWhile Loop:")
	j := 0
	for j < 5 {
		fmt.Println("Iteration:", j)
		j++
	}

	// Infinite Loop (use with caution)
	fmt.Println("\nInfinite Loop:")
	k := 0
	for {
		if k >= 3 {
			break // break out of the loop when k is 3 or more
		}
		fmt.Println("Iteration:", k)
		k++
	}

	// Range Loop (for iterating over slices, arrays, maps, etc.)
	// range can also be used without slice or array
	fmt.Println("\nRange Loop 1 to 5:")
	for i := range 5 { //same as for i := 0; i < 5; i++
		fmt.Println(i)
	}

	fmt.Println("\nRange Loop:")
	slice := []string{"apple", "banana", "cherry"}
	for index, value := range slice {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}

	//other use cases of for loop
	fmt.Println("\nOther Use Cases of For Loop:")
	// Looping through a map
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	for key, value := range m {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
	// Looping through a string
	str := "hello"
	for index, char := range str {
		fmt.Printf("Index: %d, Character: %c\n", index, char)
	}
	// Looping through a channel
	ch := make(chan string, 3)
	ch <- "Hello"
	ch <- "World"
	ch <- "Go"
	close(ch) // close the channel to signal no more values will be sent
	for msg := range ch {
		fmt.Println("Channel message:", msg)
	}
	// Looping through a slice with a condition
	fmt.Println("\nLooping through a slice with a condition:")
	for _, fruit := range slice {
		if fruit == "banana" {
			fmt.Println("Found banana, skipping...")
			continue // skip the rest of the loop for this iteration
		}
		fmt.Println("Fruit:", fruit)
	}
	// Nested loops
	// loops can be labeled in Go, allowing you to break or continue specific loops
	fmt.Println("\nNested Loops:")
	// nested loop example with label
OuterLoop:
	for i := range 3 {
		for j := range 3 {
			if i == 1 && j == 1 {
				fmt.Println("Breaking out of outer loop at i=1, j=1")
				break OuterLoop // break out of the outer loop
			}
			fmt.Printf("i: %d, j: %d\n", i, j)
		}
	}

}

// Note: Go does not have a traditional while loop or do-while loop. Instead, you can use the for loop to achieve similar functionality.
