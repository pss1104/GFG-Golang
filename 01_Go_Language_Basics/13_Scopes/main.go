// package level, function level, block level, loop level, closure or anonymous function scope
package main

import (
	"fmt"
	"gfg-golang/01_Go_Language_Basics/13_Scopes/util"
)

var globalVar = "Global variable within this package" //can be accessed inside the package only

func functionScope() {
	localVar := "I am a local variable"
	fmt.Println("Function Scope:", localVar)
}
func blockScope() {
	if true {
		blockVar := "I am a block variable"
		fmt.Println("Block Scope:", blockVar)
	}
	// fmt.Println(blockVar) // This will cause an error because blockVar is not accessible here
}
func loopScope() {
	for i := 0; i < 3; i++ {
		loopVar := fmt.Sprintf("I am a loop variable %d", i)
		fmt.Println("Loop Scope:", loopVar)
	}
	// fmt.Println(loopVar) // This will cause an error because loopVar is not accessible here
}

// lexical scoping in Go allows functions to access variables from their enclosing scope, which is known as closure.
// A closure is a function that captures the variables from its surrounding context, allowing it to access those variables even after the context has exited.
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {

	fmt.Println("Global Variable:", util.GlobalVar)
	fmt.Println("Package Level Variable:", globalVar)

	functionScope()
	blockScope()
	loopScope()

	// Closure example
	counterFunc := counter()
	fmt.Println("Counter:", counterFunc())
	fmt.Println("Counter:", counterFunc())

}
