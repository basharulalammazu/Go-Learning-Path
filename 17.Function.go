/*
	Function Type:
	1. Standard Function: func functionName(parameters) returnType { // function body }
	2. Init Function: func init() { // initialization code }
	3. Anonymous Function: func(parameters) returnType { // function body }()
	4. Function Expression: variableName := func(parameters) returnType { // function body }

*/

package main

import "fmt"

var num = 10  // Global variable

// Standard Function: A function with a name that can be called from other parts of the code
func StandardFunction()  {
	fmt.Println("This is a standard function.")
}


// Init Function: Automatically called when the package is initialized
// At first call global variables, then init function and finally main function
func init() {
	fmt.Println("This is an init function. \n Global variable num is:", num)
	num = 20  // Modify the global variable (not create a new local one)
}


func main() {

	fmt.Println("Hello from main function. \n Global variable num is:", num) // Right now num is 20 because init function changed it
	StandardFunction() // Calling the standard function
}