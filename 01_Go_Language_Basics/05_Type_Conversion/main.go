// type conversion with examples

// Type conversion in Go is only explicit
// No implicit type conversion
// You need to use the type conversion syntax to convert from one type to another
// Syntax: x := TypeName(value)

package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Example 1: Converting int to string
	var num int = 42
	strNum := strconv.Itoa(num) // Convert int to string
	fmt.Printf("Type of strNum: %T, Value: %s\n", strNum, strNum)

	// Example 2: Converting string to int
	str := "123"
	intNum, err := strconv.Atoi(str) // Convert string to int
	if err != nil {
		fmt.Println("Error converting string to int:", err)
	} else {
		fmt.Printf("Type of intNum: %T, Value: %d\n", intNum, intNum)
	}

	// Example 3: Converting float64 to int
	var floatNum float64 = 3.14
	intFromFloat := int(floatNum) // Convert float64 to int (truncates the decimal part)
	fmt.Printf("Type of intFromFloat: %T, Value: %d\n", intFromFloat, intFromFloat)

	// Example 4: Converting byte slice to string
	byteSlice := []byte{72, 101, 108, 108, 111} // Represents "Hello" in ASCII
	strFromBytes := string(byteSlice)           // Convert byte slice to string
	fmt.Printf("Type of strFromBytes: %T, Value: %s\n", strFromBytes, strFromBytes)
}
