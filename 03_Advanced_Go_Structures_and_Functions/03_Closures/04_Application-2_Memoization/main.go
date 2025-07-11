package main

import "fmt"

func memoize() func(int) int {
	cache := make(map[int]int)
	return func(n int) int {
		if val, exists := cache[n]; exists {
			return val
		}
		result := n * n // Example computation
		cache[n] = result
		return result
	}
}

func main() {
	square := memoize()
	fmt.Println(square(4)) // Output: 16 (computed)
	fmt.Println(square(4)) // Output: 16 (cached)
}
