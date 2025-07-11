// runes are alias for int 32
// represents a single unicode code point
// runes as slice - sequence of unicode characters
/*
| Feature  | String                    | Rune                                   |
| -------- | ------------------------- | -------------------------------------- |
| Type     | `[]byte`                  | `int32`                                |
| Encoding | UTF-8 encoded, byte-based | Represents a single Unicode code point |
| Length   | Measured in bytes         | Measured in characters                 |
| Use Case | Simple text processing    | Multibyte character handling           |
*/
package main

import (
	"fmt"
	"strconv"
	"unicode"
	"unicode/utf8"
)

func main() {
	var ch rune = 'A'
	fmt.Println(ch)         // Output: 65 (Unicode code point of 'A')
	fmt.Println(string(ch)) // Output: A

	s := "Go✓9"

	fmt.Println("Length in bytes:", len(s))
	fmt.Println("Length in runes:", len([]rune(s)))

	// 1. Count runes (characters) in string
	fmt.Println("Rune count:", utf8.RuneCountInString(s)) // Output: 3

	// 2. Convert string to rune slice
	runes := []rune(s)
	fmt.Println("Rune slice:", runes) // Output: [71 111 10003]

	// 3. Print each rune with its properties
	for _, r := range runes {
		fmt.Printf("\nRune: %q\n", r) // Print rune as quoted char

		// Check character type
		fmt.Println("  IsLetter:", unicode.IsLetter(r))
		fmt.Println("  IsDigit:", unicode.IsDigit(r))

		// Change case
		fmt.Println("  ToUpper:", string(unicode.ToUpper(r)))
		fmt.Println("  ToLower:", string(unicode.ToLower(r)))

		// Quoted rune representation
		fmt.Println("  Quoted:", strconv.QuoteRune(r))
	}

	// 4. Convert rune slice back to string
	rebuilt := string(runes)
	fmt.Println("\nRebuilt string:", rebuilt) // Output: Go✓9
}
