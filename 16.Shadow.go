package main

import "fmt"

var a = "Global Variable" 

func main() {
	a := "Local Variable"
	age := 20 
	
	if age >= 18 {
		a := "Block Variable"
		fmt.Println("Inside if block:", a) // Prints "Block Variable"
	}
	
	fmt.Println("Inside main function:", a) // Prints "Local Variable"
	
	printGlobal() 	// To access the global variable, we need a separate function

}


// Function to demonstrate access to global variable
func printGlobal() {
	fmt.Println("Global variable:", a) // Prints "Global Variable"
}