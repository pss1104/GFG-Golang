// slices in go
package main

import "fmt"

func main() {
	var s []int
	fmt.Println("just declared slice: ", s)
	s = []int{1, 2}
	fmt.Println("initialized slice: ", s)
	s = append(s, 3, 4)
	fmt.Println("appened slice: ", s)

	//slice with make
	sl := make([]int, 3, 4) //make uses type, len, capacity of slice
	//len is current number of elements stored
	//cap is the size of underlying array
	fmt.Println("sl: ", sl)
	sl = nil
	fmt.Println("emptied sl: ", sl)
	sl = append(sl, 1, 2)
	fmt.Println("sl: ", sl)
	fmt.Println("len: ", len(sl), ", cap: ", cap(sl))
	sl = append(sl, 3, 4, 5)
	fmt.Println("sl: ", sl)
	fmt.Println("len: ", len(sl), ", cap: ", cap(sl))
	//if you exceed the capacity, the capacity doubles
	sl = append(sl, 6, 7, 8, 9, 10)
	fmt.Println("sl: ", sl)
	fmt.Println("len: ", len(sl), ", cap: ", cap(sl))

	// 	For slices declared with zero elements, the capacity starts at 1.
	// Each time the slice's capacity is exceeded, Go doubles the capacity until it reaches the required size.

	//arrays to slice
	a := [3]int{1, 2, 3}
	as := a[:]
	fmt.Println("slice from array: ", as)

	//multi-dim
	multiDimSlice := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(multiDimSlice)
}
