// you can define custom error types by implementing the error interface.
package main

import (
	"fmt"
)

// CustomError defines a custom error type with a code and message.
type CustomError struct {
	Code    int
	Message string
}

// Error implements the error interface for CustomError.
func (e *CustomError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

// generateError creates and returns a new CustomError.
func generateError() error {
	return &CustomError{Code: 404, Message: "Resource not found"}
}

// check function returns a custom error for an invalid code.
func check(code int) (string, error) {
	switch code {
	case 400:
		return "", &CustomError{Code: 400, Message: "Invalid Request"}
	case 404:
		return "", &CustomError{Code: 404, Message: "Resource Not Found"}
	default:
		return "Success", nil
	}
}
func main() {
	e := generateError()
	fmt.Println(e) // Output: Error 404: Resource not found

	_, er := check(400)
	if er != nil {
		fmt.Println(er)
	}

	_, err := check(404)
	if err != nil {
		fmt.Println(err)
	}
}
