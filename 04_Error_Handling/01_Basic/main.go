/*
In Go, errors are treated as values instead of exceptions
Errors are values that implement the error interface.

Error Interface:

	type error interface {
	    Error() string
	}

Go provides the errors package for creating simple errors.
errors.new("...")

Return nil for no error
*/
package main

import (
	"errors"
	"fmt"
)

func divby0(a, b int) (c float32, e error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	c = float32(a) / float32(b)
	return c, nil
}

func main() {
	c, e := divby0(2, 3)

	if e == nil {
		fmt.Println("Division result: ", c)
	} else {
		fmt.Println("Error occured: ", e)
	}

	c, e = divby0(3, 0)

	if e == nil {
		fmt.Println("Division result: ", c)
	} else {
		fmt.Println("Error occured: ", e)
	}
}
