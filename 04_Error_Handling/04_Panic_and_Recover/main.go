// A panic is a built-in mechanism in GoLang that stops the normal flow of a program when it encounters a severe error.
// A panic immediately halts the execution of the current function.
// Deferred functions are executed before the program terminates.
// The program prints a stack trace and exits unless the panic is recovered.

package main

import "fmt"

// Function without recover
// func divide(a, b int) int {
//     if b == 0 {
//         panic("Division by zero!")
//     }
//     return a / b
// }

// Function with recover
func divide(a, b int) (result int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
			result = 0 // Set result to 0 or any default value after recovery
		}
	}()
	if b == 0 {
		panic("Division by zero!")
	}
	return a / b
}

func main() {
	fmt.Println(divide(4, 2)) // Executes successfully
	fmt.Println(divide(4, 0)) // Triggers panic //Recovers if second function is used
	fmt.Println(divide(8, 2)) // This line will not execute
}
