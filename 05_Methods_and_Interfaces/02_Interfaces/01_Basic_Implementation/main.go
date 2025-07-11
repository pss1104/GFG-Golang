// An interface in GoLang is a collection of method signatures.
// It defines what a type can do without specifying how it does it. Any type that implements all the methods declared in an interface satisfies the interface.
// No Explicit Implementation: Unlike Java or C++, Go doesn’t require explicit keywords like implements or extends.
package main

import "fmt"

type Sedan struct {
	Brand  string
	Length int // Length in feet
}

func (s Sedan) StartEngine() string {
	return fmt.Sprintf("Brand %s has started the engine silently.", s.Brand)
}

func (s Sedan) Drive() string {
	return fmt.Sprintf("Brand %s can drive smoothly with a length of %d feet.", s.Brand, s.Length)
}

type SUV struct {
	Brand  string
	Height int // Height in meters
}

func (s SUV) StartEngine() string {
	return fmt.Sprintf("Brand %s has started the engine with a roar.", s.Brand)
}

func (s SUV) Drive() string {
	return fmt.Sprintf("Brand %s can drive on any terrain with a height of %d meters.", s.Brand, s.Height)
}

func main() {
	var vehicle interface{}

	// Assigning Sedan to the interface
	vehicle = Sedan{Brand: "Toyota", Length: 15}
	fmt.Println(vehicle.(Sedan).StartEngine())
	fmt.Println(vehicle.(Sedan).Drive())

	// Assigning SUV to the interface
	vehicle = SUV{Brand: "Ford", Height: 2}
	fmt.Println(vehicle.(SUV).StartEngine())
	fmt.Println(vehicle.(SUV).Drive())
}
