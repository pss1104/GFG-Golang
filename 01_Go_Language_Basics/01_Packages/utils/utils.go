// custom package utils
package utils

import "fmt"

func init() {
	fmt.Println("utils package initialized")
}

// PrintMessage prints a message to the console
func PrintMessage(message string) {
	fmt.Println(message)
}

var X int = 10 // Exported variable ( capital letter ), accessible outside the package
var Y int = 20

// Z:=25 //:= Short variable declaration, only works inside functions

// var z int = 30 // Unexported variable (small letter) , not accessible outside the package

// Add is a function that adds two integers and returns the result
func Add(a int, b int) int {
	return a + b
}
