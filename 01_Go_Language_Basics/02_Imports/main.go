package main

import (
	"fmt"
	m "math" // Importing the math package with an alias 'm'

	"github.com/fatih/color"
)

func main() {
	color.Red("This is Red text!")
	fmt.Println("Square root of 16 is:", m.Sqrt(16)) // Using the Sqrt function from the math package
}
