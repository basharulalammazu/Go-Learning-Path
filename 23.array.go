// 23.array.go

package main

import "fmt"

// Using individual var declarations (previous grouped var block triggered a syntax error in this lesson context)
var name = [2]string{"Basharul", "Mazu"}
var age  = [3]int{20, 30, 40}

func main() {
	var array [2]int

	fmt.Printf("Array: %v\n", array)

	array2 := [3]int{10, 20, 30}
	fmt.Printf("Array 2: %v\n", array2)

	fmt.Printf("Array 2 Length: %d\n", len(array2))

	// Global Array
	fmt.Printf("Name Array: %v\n", name)
	fmt.Printf("Name Array Length: %d\n", len(name))
	fmt.Printf("Age Array: %v\n", age)
	fmt.Printf("Age Array Length: %d\n", len(age))
}