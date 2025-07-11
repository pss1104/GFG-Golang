// Polymorphism allows you to use a single interface to handle multiple types.
// This is achieved in GoLang by defining behaviors through interfaces and dynamically assigning them to various types.

package main

import (
	"fmt"
)

// Define the Car interface
type Car interface {
	StartEngine() string
	Drive() string
}

// Define the Sedan struct
type Sedan struct {
	Brand  string
	Length int // Length in feet
}

// Implement the StartEngine method for Sedan
func (s Sedan) StartEngine() string {
	return fmt.Sprintf("Brand %s has started the engine silently.", s.Brand)
}

// Implement the Drive method for Sedan
func (s Sedan) Drive() string {
	return fmt.Sprintf("Brand %s can drive smoothly with a length of %d feet.", s.Brand, s.Length)
}

// Define the SUV struct
type SUV struct {
	Brand  string
	Height int // Height in meters
}

// Implement the StartEngine method for SUV
func (s SUV) StartEngine() string {
	return fmt.Sprintf("Brand %s has started the engine with a roar.", s.Brand)
}

// Implement the Drive method for SUV
func (s SUV) Drive() string {
	return fmt.Sprintf("Brand %s can drive on any terrain with a height of %d meters.", s.Brand, s.Height)
}

// Main function
func main() {
	var car Car

	car = Sedan{Brand: "Honda", Length: 14}
	fmt.Println(car.StartEngine())
	fmt.Println(car.Drive())

	car = SUV{Brand: "Ford", Height: 3}
	fmt.Println(car.StartEngine())
	fmt.Println(car.Drive())
}
