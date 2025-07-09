// conditional statements
package main

import (
	"fmt"
)

func main() {
	// if-else statement
	x := 10
	if x > 0 {
		fmt.Println("x is positive")
	} else if x < 0 {
		fmt.Println("x is negative")
	} else {
		fmt.Println("x is zero")
	}

	// if-else with short variable declaration
	if y := x * 2; y > 10 {
		fmt.Println("y is greater than 10:", y)
	} else {
		fmt.Println("y is not greater than 10:", y)
	}

	// switch statement
	switch x {
	case 1:
		fmt.Println("x is 1")
	case 2:
		fmt.Println("x is 2")
	default:
		fmt.Println("x is neither 1 nor 2")
	}

	// nested if-else statement
	if x > 0 {
		if x%2 == 0 {
			fmt.Println("x is positive and even")
		} else {
			fmt.Println("x is positive and odd")
		}
	} else {
		fmt.Println("x is not positive")
	}

	// switch with multiple conditions
	switch {
	case x < 0:
		fmt.Println("x is negative")
	case x == 0:
		fmt.Println("x is zero")
	case x > 0 && x < 10:
		fmt.Println("x is positive and less than 10")
	case x >= 10:
		fmt.Println("x is positive and greater than or equal to 10")
	default:
		fmt.Println("x is something else")
	}

}
