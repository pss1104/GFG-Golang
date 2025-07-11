package main

import "fmt"

type Rectangle struct {
	Width  int
	Length int
}

// Struct receiver method to calculate area
// This method uses a struct receiver, so it operates on a copy of the Rectangle.
// Any changes made within the method will not affect the original Rectangle.
func (r Rectangle) Area() int {
	return r.Width * r.Length
}

// Pointer receiver method to scale dimensions
// This method uses a pointer receiver, so it operates on the original struct.
// Changes made within the method will affect the original Rectangle.
func (r *Rectangle) Scale(factor int) {
	r.Width *= factor
	r.Length *= factor
}

func main() {
	// Initialize a Rectangle
	rect := Rectangle{Width: 5, Length: 3}

	// Call the Area method (struct receiver)
	fmt.Println("Initial Area:", rect.Area())

	// Call the Scale method (pointer receiver)
	rect.Scale(10)

	// Check updated area and dimensions
	fmt.Println("Scaled Area:", rect.Area())
	fmt.Println("Updated Width:", rect.Width)
	fmt.Println("Updated Length:", rect.Length)
}
