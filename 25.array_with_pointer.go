// 25.array_with_pointer.go

// pass by value: normal passing
// pass by reference:  using pointer
package main

import "fmt"

func print(array *[5]int) {
	fmt.Println("Array inside function:", *array)
	fmt.Println("Array Length inside function:", len(*array))
}

func main() {
	var array = [5]int{1, 2, 3, 4, 5}
	print(&array)
}