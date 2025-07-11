// When working with multiple potential types, using if statements for type assertions can be cumbersome.
// Instead, Type Switches provide a cleaner way to handle type-specific logic

package main

import "fmt"

func main() {
	var interfaceValue interface{}

	// Test with an integer
	interfaceValue = 42
	processValue(interfaceValue)

	// Test with a string
	interfaceValue = "Hello, Go!"
	processValue(interfaceValue)

	// Test with a float
	interfaceValue = 3.14
	processValue(interfaceValue)

	// Test with a boolean
	interfaceValue = true
	processValue(interfaceValue)
}

func processValue(value interface{}) {
	switch v := value.(type) {
	case int:
		fmt.Println("Type is int:", v)
	case string:
		fmt.Println("Type is string:", v)
	default:
		fmt.Println("Unsupported type:", v)
	}
}
