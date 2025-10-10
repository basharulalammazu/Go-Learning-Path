/*
	slice => A slice is a dynamically-sized, flexible view into the elements of an array.
	It is a lightweight data structure that provides access to a contiguous segment of an array.
	Slices are more commonly used in Go than arrays because they offer greater flexibility and convenience.

	Here are some key points about slices in Go:
	1. Dynamic Size: Unlike arrays, slices can grow and shrink in size as needed. You can append elements to a slice using the built-in append function.
	2. Underlying Array: A slice is backed by an underlying array. When you create a slice, it references a portion of an array, and changes made to the slice will affect the underlying array.
	3. Slice Literal: You can create a slice using a slice literal, which is similar to an array literal but without specifying the size. For example: mySlice := []int{1, 2, 3}.
	4. Length and Capacity: Slices have both a length (the number of elements in the slice) and a capacity (the maximum number of elements the slice can hold before needing to allocate more memory). You can use the built-in len and cap functions to get these values.
	5. Slicing Syntax: You can create a new slice from an existing slice or array using slicing syntax. For example: newSlice := mySlice[1:4] creates a new slice containing elements from index 1 to index 3 of mySlice.
*/

package main

import (
	"fmt"
	"sort"
)

func main(){
	arr := [6]string {"This", "is", "a", "GO", "Interview", "Question"}
	fmt.Println(arr)

	// Create a slice from the array
	slices := arr[1:4] // Pointer 1, length 3, capacity 5
	fmt.Println("Length of slices:", len(slices))
	// Length = ending_index - starting_index = 4 - 1 = 3 => "is", "a", "GO"
	fmt.Println("Capacity of slices:", cap(slices)) // Capacity is calculated from the starting index to the end of the array
	// Capacity = len(arr) - starting_index = 6 - 1 = 5 => "is", "a", "GO", "Interview", "Question"
	fmt.Println("Slices:", slices)

	sort.Strings(slices)
	fmt.Println("Sorted Slices:", slices)





	// Create a new slice from the existing slice 
	slices2 := slices[1:3] // Pointer 1, length 2, capacity 4
	fmt.Println("Length of slices2:", len(slices2))	
	fmt.Println("Capacity of slices2:", cap(slices2))
	fmt.Println("Slices2:", slices2)





	// slices_literal => a shorthand syntax used to create a slice by specifying its elements directly
	// Create a slice using make function
	slices_literal := []int {10, 20, 30, 40} // Pointer 0, length 4, capacity 4
	fmt.Println("Slices Literal:", slices_literal)
	fmt.Println("Length of slices_literal:", len(slices_literal))
	fmt.Println("Capacity of slices_literal:", cap(slices_literal))



	// Create a slice using make function with length
	s1 := make([]int, 3) // Pointer 0, length 3, capacity 3
	fmt.Println("Slice s1:", s1)
	fmt.Println("Length of s1:", len(s1))
	fmt.Println("Capacity of s1:", cap(s1))

	// Create a slice using make function with specified capacity and length
	s2 := make([]int, 3, 5) // Pointer 0, length 3, capacity 5
	fmt.Println("Slice s2:", s2)
	fmt.Println("Length of s2:", len(s2))
	fmt.Println("Capacity of s2:", cap(s2))


	// Empty Slice
	var s3 []int
	fmt.Println("Slice s3:", s3)
	fmt.Println("Length of s3:", len(s3))
	fmt.Println("Capacity of s3:", cap(s3))
	 
	// Append elements to a slice
	s3 = append(s3, 1)
	fmt.Println("Slice s3 after appending 1:", s3)
	fmt.Println("Length of s3 after appending 1:", len(s3))
	fmt.Println("Capacity of s3 after appending 1:", cap(s3))

}