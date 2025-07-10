package main

import "fmt"

func main() {
	var a = [3]int{1, 2, 3}
	fmt.Println("a: ", a)

	b := [7]int{1, 2, 3}
	fmt.Println("b ", b)

	a[0] = -1
	fmt.Println("a: ", a)

	b[3] = 4
	fmt.Println("b ", b)

	var c [3]int
	fmt.Println("c: ", c)
	c = [3]int{1}
	fmt.Println("c: ", c)

	var matrix [2][3]int // A 2x3 matrix (2 rows, 3 columns)
	matrix[0][0] = 1
	fmt.Println(matrix[0][0]) // Output: 1

}
