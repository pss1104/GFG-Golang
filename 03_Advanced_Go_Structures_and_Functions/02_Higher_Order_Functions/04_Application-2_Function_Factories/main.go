package main

import "fmt"

// Factory function to create greeting functions
func greeting(language string) func(string) string {
	switch language {
	case "English":
		return func(name string) string {
			return "Hello, " + name
		}
	case "Spanish":
		return func(name string) string {
			return "Hola, " + name
		}
	default:
		return func(name string) string {
			return "Hi, " + name
		}
	}
}

func main() {
	greetInEnglish := greeting("English")
	greetInSpanish := greeting("Spanish")

	fmt.Println(greetInEnglish("Alice")) // Output: Hello, Alice
	fmt.Println(greetInSpanish("Bob"))   // Output: Hola, Bob
}
