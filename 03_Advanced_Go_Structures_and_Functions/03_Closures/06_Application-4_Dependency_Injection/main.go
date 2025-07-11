package main

import "fmt"

func logger(prefix string) func(string) {
	return func(message string) {
		fmt.Println(prefix, message)
	}
}

func main() {
	infoLogger := logger("[INFO]")
	errorLogger := logger("[ERROR]")

	infoLogger("This is an info message.")   // Output: [INFO] This is an info message.
	errorLogger("This is an error message.") // Output: [ERROR] This is an error message.
}
