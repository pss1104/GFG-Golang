package main

import "fmt"

func processValue(value interface{}) {
	switch v := value.(type) {
	case int:
		fmt.Println("Processing integer:", v)
	case string:
		fmt.Println("Processing string:", v)
	case []string:
		fmt.Println("Processing string slice:", v)
	default:
		fmt.Println("Unsupported type:", v)
	}
}

func main() {
	processValue(42)
	processValue("Golang")
	processValue([]string{"Go", "is", "great"})
}
