/*
goto in Go is used to jump to a labeled statement within the same function.
It can be useful for breaking out of deeply nested loops or for skipping certain iterations
Using goto can make code harder to read, debug, and maintain, especially when multiple goto statements are present.
This leads to "spaghetti code," where the program's flow becomes unpredictable.
*/
package main

import (
	"fmt"
)

func main() {
	// Using goto to jump to a label
	fmt.Println("Using goto:")
	i := 0
LoopStart:
	if i < 5 {
		fmt.Println("Iteration:", i)
		i++
		goto LoopStart // Jump back to LoopStart label
	}
	fmt.Println("Loop completed using goto")
	// Using goto to skip iterations
	fmt.Println("\nUsing goto to skip iterations:")
	for j := 0; j < 10; j++ {
		if j == 3 {
			fmt.Println("Skipping iteration:", j)
			goto SkipIteration // Jump to SkipIteration label
		}
		fmt.Println("Processing:", j)
	}
SkipIteration:
	fmt.Println("Skipped iteration 3")
	// Using goto to break out of a loop
	fmt.Println("\nUsing goto to break out of a loop:")
	for k := 0; k < 10; k++ {
		if k == 5 {
			fmt.Println("Breaking out of the loop at iteration:", k)
			goto EndLoop // Jump to EndLoop label
		}
		fmt.Println("Processing:", k)
	}
EndLoop:
	fmt.Println("Loop ended at iteration 5")

	// Using goto with error handling
	fmt.Println("\nUsing goto with error handling:")
	for l := range 5 {
		if l == 2 {
			fmt.Println("Error encountered at iteration:", l)
			goto ErrorHandler // Jump to ErrorHandler label
		}
		fmt.Println("Processing:", l)
	}
ErrorHandler:
	fmt.Println("Error handled at iteration 2, continuing execution")
	fmt.Println("End of program")

}
