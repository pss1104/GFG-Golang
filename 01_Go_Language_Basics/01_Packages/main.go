// 2 types of packages - Executable and Library
package main

import (
	"fmt"
	"gfg-golang/01_Go_Language_Basics/01_Packages/utils" // Importing a custom package named utils
	//path should start from the root of the module(go.mod)
)

func main() {

	// Using a function from the utils package
	utils.PrintMessage("This is a message from the utils package!")

	// using X & Y from utils package
	// Using the Add function from the utils package
	sum := utils.Add(utils.X, utils.Y)
	fmt.Println("Sum of 5 and 10 is:", sum)

}
