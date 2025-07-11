//For better error tracing, Go provides fmt.Errorf to wrap errors with additional context.
//Wrapping errors can provide us with necessary information such as the error propagation stack
// To extract the original error from a wrapped error, use errors.Unwrap.

package main

import (
	"errors"
	"fmt"
)

// Base error for division by zero
var ErrDivideByZero = errors.New("cannot divide by zero")

// Lowest-level function
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("divide failed: %w", ErrDivideByZero)
	}
	return a / b, nil
}

// Higher-level function that calls divide and wraps error
func compute() error {
	_, err := divide(10, 0)
	if err != nil {
		return fmt.Errorf("compute error: %w", err)
	}
	return nil
}

func main() {
	err := compute()
	if err != nil {
		fmt.Println("Error:", err)

		// Check if the root cause was divide by zero
		if errors.Is(err, ErrDivideByZero) {
			fmt.Println("Cause: division by zero")
		}

		// Show immediate wrapped error
		fmt.Println("Unwrapped once:", errors.Unwrap(err))

		fmt.Println("Unwrapped twice:", errors.Unwrap(errors.Unwrap(err)))

		fmt.Println("Unwrapped thrice:", errors.Unwrap(errors.Unwrap(errors.Unwrap(err))))
	}
}
