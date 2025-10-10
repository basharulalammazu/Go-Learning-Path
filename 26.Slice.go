/*

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
	slices := arr[1:5]
	fmt.Println("Slices:", slices)

	sort.Strings(slices)
	fmt.Println("Sorted Slices:", slices)
}