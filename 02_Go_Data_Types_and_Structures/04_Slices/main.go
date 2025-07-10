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
	sl := make([]int, 3, 4)
	fmt.Println("sl: ", sl)
	sl = nil
	fmt.Println("emptied sl: ", sl)
	sl = append(sl, 1, 2)
	fmt.Println("sl: ", sl)
	fmt.Println("len: ", len(sl), ", cap: ", cap(sl))
	sl = append(sl, 3, 4, 5)
	fmt.Println("sl: ", sl)
	fmt.Println("len: ", len(sl), ", cap: ", cap(sl))
	sl = append(sl, 6, 7, 8, 9, 10)
	fmt.Println("sl: ", sl)
	fmt.Println("len: ", len(sl), ", cap: ", cap(sl))

	//arrays to slice
	a := [3]int{1, 2, 3}
	as := a[:]
	fmt.Println("slice from array: ", as)
}
