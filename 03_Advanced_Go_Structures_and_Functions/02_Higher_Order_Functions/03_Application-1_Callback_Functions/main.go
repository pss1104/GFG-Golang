package main

import "fmt"

func process(nums []int, fn func(int) int) []int {
	result := make([]int, len(nums))
	for i, v := range nums {
		result[i] = fn(v)
	}
	return result
}

func square(n int) int {
	return n * n
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	squared := process(numbers, square)
	fmt.Println("Squared Numbers:", squared) // Output: [1 4 9 16 25]
}
