// Strings in Go are readonly sequence of bytes
// Strings are immutable
// Strings are UTF-8 encoded - each char takes 1-4 bytes
// strings as slice: sequence of bytes
package main

import (
	"fmt"
	"strings"
	// "unicode"
)

func main() {
	var s string
	fmt.Println("empty string: ", s)
	s = "Hello World"
	fmt.Println(s)
	s = s + "!" //concatenation
	fmt.Println(s)

	fmt.Println(strings.ToLower(s))
	fmt.Println(strings.ToUpper(s))
	fmt.Println("count of o: ", strings.Count(s, "o")) //counts non overlapping instances
	fmt.Println("contains lo: ", strings.Contains(s, "lo"))
	fmt.Println("has prefix H: ", strings.HasPrefix(s, "H"))
	fmt.Println("has suffix !: ", strings.HasPrefix(s, "!"))
	fmt.Println("replace ! by *: ", strings.Replace(s, "!", "*", 2))  //replaces first 2 non overlapping instances
	fmt.Println("replace o by O: ", strings.Replace(s, "o", "O", -1)) //replaces all
	fmt.Println("replace l by L: ", strings.ReplaceAll(s, "l", "L"))  //replaces all
	fmt.Println("split at space: ", strings.Split(s, " "))
	sl := strings.Split(s, "")
	fmt.Println("split all chars as strings: ", sl) //slice of split strings
	//join slice of strings
	fmt.Println("joined slice of strings by - : ", strings.Join(sl, "-"))

	//character access
	fmt.Println("2nd char: ", s[1]) //gives ascii num of the character
	fmt.Println("2nd char as string: ", string(s[1]))
	//substring from ind 2 to 5
	fmt.Println("substring: ", s[2:6])

	//byte slice conversion
	b := []byte(s) //same as using unit8
	fmt.Println("byte slice: ", b)

	//number of bytes of string is found using len()
	fmt.Println("number of bytes of string: ", len(s))
	//Some characters (like A, b, 1) are 1 byte.
	// Others (like ✓, 文, 🚀) are 2–4 bytes.
	//So len(s) counts bytes, not characters (runes).
	// normal strings with a-z, A-Z, 1-9 -> len() gives number of characters
	fmt.Println("len of Hello: ", len("Hello"))
	fmt.Println("len of A✓文🚀: ", len("A✓文🚀"))

	//strings are generally declared using " "
	// ` ` back ticks can be used to declare raw strings, i.e., escape sequences are treated literally
	s1 := `old line \n new line`
	fmt.Println("string with ``: ", s1)
}
