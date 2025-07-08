//Operators in go

package main

import "fmt"

func main() {
	// Arithmetic Operators
	a, b := 10, 20
	fmt.Println("Addition:", a+b)       // 30
	fmt.Println("Subtraction:", a-b)    // -10
	fmt.Println("Multiplication:", a*b) // 200
	fmt.Println("Division:", b/a)       // 2
	fmt.Println("Modulus:", b%a)        // 0

	// Comparison Operators
	fmt.Println("Equal:", a == b)                    // false
	fmt.Println("Not Equal:", a != b)                // true
	fmt.Println("Greater Than:", a > b)              // false
	fmt.Println("Less Than:", a < b)                 // true
	fmt.Println("Greater Than or Equal To:", a >= b) // false
	fmt.Println("Less Than or Equal To:", a <= b)    // true

	// Logical Operators
	x, y := true, false
	fmt.Println("Logical AND:", x && y) // false
	fmt.Println("Logical OR:", x || y)  // true
	fmt.Println("Logical NOT:", !x)     // false

	// Bitwise Operators
	c, d := 5, 3                      // 0101 and 0011 in binary
	fmt.Println("Bitwise AND:", c&d)  // 1 (0001)
	fmt.Println("Bitwise OR:", c|d)   // 7 (0111)
	fmt.Println("Bitwise XOR:", c^d)  // 6 (0110)
	fmt.Println("Left Shift:", c<<1)  // 10 (1010)
	fmt.Println("Right Shift:", c>>1) // 2 (0010)
	//&^ is a bitwise AND NOT operator
	fmt.Println("Bitwise AND NOT:", c&^d) // 4 (0100

	// Assignment Operators
	var e int = 10
	e += 5                       // e = e + 5
	fmt.Println("After += :", e) // 15

	e -= 3                       // e = e - 3
	fmt.Println("After -= :", e) // 12

	e *= 2                       // e = e * 2
	fmt.Println("After *= :", e) // 24

	e /= 4                       // e = e / 4
	fmt.Println("After /= :", e) // 6

	e %= 5                       // e = e % 5
	fmt.Println("After %= :", e) // 1

	//others: &=, |=, ^=, <<=, >>= are also assignment operators

	// Unary Operators (can only be used for numeric types)
	f := -e                           // negation operator
	g := +e                           // positive operator (no effect)
	h := ^c                           // bitwise NOT operator
	fmt.Println("Unary Negation:", f) // -1
	fmt.Println("Unary Positive:", g) // 1
	fmt.Println("Bitwise NOT:", h)    // -6 (inverts bits of c)

	// Increment and Decrement Operators
	i := 10
	i++                                // increment operator
	fmt.Println("After Increment:", i) // 11
	i--                                // decrement operator
	fmt.Println("After Decrement:", i) // 10

	// Miscellaneous Operators - address of a variable and dereference operator
	var j int = 42
	var ptr *int = &j                  // address of operator
	fmt.Println("Address of j:", ptr)  // prints the address of j
	fmt.Println("Value at ptr:", *ptr) // dereference operator, prints the value of j (42)

	// Miscellaneous Operators - comma operator
	xc, yc := 1, 2                              // multiple assignment
	fmt.Println("Multiple Assignment:", xc, yc) // 1 2

	// Miscellaneous Operators - type assertion operator
	var iFace interface{} = "Hello, Go!"
	str, ok := iFace.(string) // type assertion
	if ok {
		fmt.Println("Type Assertion Successful:", str) // Hello, Go!
	} else {
		fmt.Println("Type Assertion Failed")
	}

}
