/*
 1. Value Receiver
    A copy of the struct is passed to the method.
    Changes inside the method do not affect the original struct.
 2. Pointer Receiver
    A reference to the struct is passed to the method.
    Changes inside the method affect the original struct
*/
package main

import "fmt"

// Struct definition
type Car struct {
	Name  string
	Type  string
	Brand string
	Years int
}

// Method with value receiver
func (c Car) DisplayDetails() {
	fmt.Printf("Car: %s, Brand: %s, Type: %s\n", c.Name, c.Brand, c.Type)
}

// Method with pointer receiver
func (c *Car) UpdateBrand(newBrand string) {
	c.Brand = newBrand
}

// Method to check operational status
func (c Car) OperationalStatus() {
	if c.Years > 15 {
		fmt.Println(c.Name, "is no longer operational.")
	} else {
		fmt.Printf("%s has %d years of operation left.\n", c.Name, 15-c.Years)
	}
}

func main() {
	// Initialize struct
	car := Car{
		Name:  "Model S",
		Type:  "Electric",
		Brand: "Tesla",
		Years: 3,
	}

	// Access fields
	fmt.Println("Car Name:", car.Name)
	fmt.Println("Brand:", car.Brand)

	// Call methods
	car.DisplayDetails()
	car.OperationalStatus()

	// Update brand using pointer receiver
	car.UpdateBrand("SpaceX")
	fmt.Println("Updated Brand:", car.Brand)
}
