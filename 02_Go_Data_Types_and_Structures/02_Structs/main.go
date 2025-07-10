package main

import "fmt"

// Define the Engine struct
type Engine struct {
	Horsepower int
	Type       string
}

// Define the Car struct with embedded Engine
type Car struct {
	Name           string
	Brand          string
	Type           string
	YearsInService int
	Engine         Engine
}

func main() {
	// Named field initialization
	car1 := Car{
		Name:           "Model S",
		Brand:          "Tesla",
		Type:           "Electric",
		YearsInService: 3,
		Engine: Engine{
			Horsepower: 670,
			Type:       "Dual Motor",
		},
	}
	fmt.Println("Car Name:", car1.Name)
	fmt.Println("Engine Horsepower:", car1.Engine.Horsepower)

	// Positional field initialization
	car2 := Car{"Model 3", "Tesla", "Electric", 2, Engine{450, "Single Motor"}}
	fmt.Println("Car 2 Engine Type:", car2.Engine.Type)

	// Anonymous struct
	person := struct {
		Name string
		Age  int
	}{
		Name: "Alice",
		Age:  30,
	}
	fmt.Println("Person's Name:", person.Name)

	// Comparing structs
	if car1 == car2 {
		fmt.Println("The cars are identical.")
	} else {
		fmt.Println("The cars are different.")
	}
	car3 := car1
	if car1 == car3 {
		fmt.Println("Car1 and Car3 are identical.")
	} else {
		fmt.Println("Car1 and Car3 are different.")
	}
}
