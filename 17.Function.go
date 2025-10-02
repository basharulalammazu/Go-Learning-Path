/*
	Function Type:
	1. Standard/Named Function: func functionName(parameters) returnType { // function body }
	2. Init Function: func init() { // initialization code }
	3. Anonymous Function: func(parameters) returnType { // function body }()
	4. Immediate Invoked Function Expression (IIFE): (func(parameters) returnType { // function body })(arguments)
	5. Function Expression: variableName := func(parameters) returnType { // function body }
	6. Noob Function Expression: variableName := func(parameters) returnType { // function body }


	// Invoke mean to call a function
*/

package main

import "fmt"

var num = 10  // Global variable

// Standard/Named Function: A function with a name that can be called from other parts of the code
func StandardFunction()  {
	fmt.Println("This is a standard function.")
}


// Init Function: Automatically called when the package is initialized
// At first call global variables, then init function and finally main function
func init() {
	fmt.Println("This is an init function. \nGlobal variable num is:", num)
	num = 20  // Modify the global variable (not create a new local one)
}


func main() {

	fmt.Println("Hello from main function. \nGlobal variable num is:", num) // Right now num is 20 because init function changed it
	StandardFunction() // Calling the standard function

	// When a function does not have a name, it is called an anonymous function
	// It can be assigned to a variable or called immediately so here we are calling it immediately Invoked Expression (IIFE)
	func () {
		fmt.Println("This is an anonymous function.")
	}()


	// Noob function expression
	noobFunction := func() {
		fmt.Println("This is a noob function expression assigned to a variable.")
	}
}