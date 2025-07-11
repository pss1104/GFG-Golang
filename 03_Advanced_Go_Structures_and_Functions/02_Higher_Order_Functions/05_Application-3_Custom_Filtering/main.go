package main

import "fmt"

// Filter function
func filter(nums []int, predicate func(int) bool) []int {
	result := []int{}
	for _, v := range nums {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// Predicate functions
func isEven(n int) bool {
	return n%2 == 0
}

func isGreaterThanThree(n int) bool {
	return n > 3
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	evens := filter(numbers, isEven)
	greaterThanThree := filter(numbers, isGreaterThanThree)

	fmt.Println("Even Numbers:", evens)           // Output: [2 4 6]
	fmt.Println("Numbers > 3:", greaterThanThree) // Output: [4 5 6]
}
