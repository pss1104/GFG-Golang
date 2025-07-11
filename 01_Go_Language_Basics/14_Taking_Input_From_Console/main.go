package main

import (
	"fmt"
)

func main() {
	var name string
	var age int
	var height float64
	var city string

	// 1. Using fmt.Scan
	fmt.Print("Enter your name and age (e.g., Alice 21): ")
	fmt.Scan(&name, &age) // reads space-separated input
	fmt.Println("Using Scan →", name, age)

	// 2. Using fmt.Scanln
	fmt.Print("Enter your city: ")
	fmt.Scanln(&city) // reads up to newline (1 word)
	fmt.Println("Using Scanln →", city)

	// 3. Using fmt.Scanf
	fmt.Print("Enter your height in meters (e.g., 1.75): ")
	fmt.Scanf("%f", &height) // formatted input for float
	fmt.Println("Using Scanf →", height)
}
