/*
composite data types are: arrays, slices, maps, structs, and channels.
*/

package main

import "fmt"

func main() {
	// Array
	var arr [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array:", arr)

	// Slice
	slice := []string{"apple", "banana", "cherry"}
	fmt.Println("Slice:", slice)
	// Slices can be created using the make function
	slice1 := make([]string, 3) // explaining in detail
	// make function creates a slice with a length of 3 and capacity of 3
	// can we exceed the capacity of a slice? yes, we can append elements to a slice
	// but we cannot exceed the length of a slice
	slice1[0] = "orange"
	slice1[1] = "grape"
	slice1[2] = "kiwi"
	// if we try to access an index that is out of bounds, it will panic
	// slice1[3] = "melon" // this will panic: index out of range
	// we can also append elements to a slice
	slice1 = append(slice1, "melon") // appending to the slice
	// now the length of the slice is 4 and capacity is 6
	// we can also check the length and capacity of a slice
	fmt.Println("Length of slice1:", len(slice1))
	fmt.Println("Capacity of slice1:", cap(slice1))

	fmt.Println("Slice created with make:", slice1)

	// Map
	m := map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println("Map:", m)

	var m1 map[string]int = make(map[string]int)
	m1["age"] = 30
	fmt.Println(m["age"])

	// Struct
	type Person struct {
		Name string
		Age  int
	}
	p := Person{Name: "Alice", Age: 30}
	fmt.Println("Struct:", p)
	//struct with make is not possible

	// Channel
	ch := make(chan string)
	go func() {
		ch <- "Hello, Channel!"
	}()
	fmt.Println("Channel:", <-ch)

	//make keyword is used to create slices, maps, and channels
	// It allocates and initializes the data structure, returning a reference to it.
	//creating array with make is not possible
}
