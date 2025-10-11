/*

🧩 Go Slice — A Powerful, Dynamic Data Structure

    ➤ Definition:
      A slice is a dynamically-sized, flexible view into the elements of an array.
      It provides access to a contiguous segment of an array and grows automatically when needed.

    ➤ Why Slices Are Preferred:
      Slices are more commonly used than arrays because they are lightweight,
      dynamic, and offer more flexibility.

    🔹 Key Features of Slices in Go:
      1️⃣ Dynamic Size:
          - Slices can grow and shrink using the built-in append() function.
      2️⃣ Underlying Array:
          - Each slice references an underlying array.
          - Modifying the slice affects the underlying array (if sharing the same base array).
      3️⃣ Slice Literal:
          - You can directly create slices using: mySlice := []int{1, 2, 3}.
      4️⃣ Length vs Capacity:
          - len() → current number of elements.
          - cap() → total space available before needing a new allocation.
      5️⃣ Slicing Syntax:
          - newSlice := mySlice[start:end]
            (includes elements from start to end-1)

    🧠 Internals:
      Every slice holds:
        • Pointer → to the first accessible element
        • Length  → number of accessible elements
        • Capacity → number of elements that can fit before reallocation
*/

package main

import (
	"fmt"
	"sort"
)

func main() {

	// --- [1] Create an array and make a slice from it ---
	arr := [6]string{"This", "is", "a", "GO", "Interview", "Question"}
	fmt.Println("Original Array:", arr)

	slices := arr[1:4] // points to elements ["is", "a", "GO"]
	fmt.Println("\n--- Slice 1 (from array) ---")
	fmt.Println("Slice Elements:", slices)
	fmt.Println("Length:", len(slices))   // 4 - 1 = 3
	fmt.Println("Capacity:", cap(slices)) // 6 - 1 = 5

	// Sort the slice (also affects the original array portion)
	sort.Strings(slices)
	fmt.Println("Sorted Slice:", slices)
	fmt.Println("Array after sorting slice:", arr)



	// --- [2] Create a new slice from an existing slice ---
	slices2 := slices[1:3] // picks elements from index 1 to 2 of 'slices'
	fmt.Println("\n--- Slice 2 (from slice 1) ---")
	fmt.Println("Slice Elements:", slices2)
	fmt.Println("Length:", len(slices2))   // 3 - 1 = 2
	fmt.Println("Capacity:", cap(slices2)) // inherited from base array



	// --- [3] Slice Literal ---
	sliceLiteral := []int{10, 20, 30, 40}
	fmt.Println("\n--- Slice Literal ---")
	fmt.Println("Slice:", sliceLiteral)
	fmt.Println("Length:", len(sliceLiteral))
	fmt.Println("Capacity:", cap(sliceLiteral))



	// --- [4] make() function examples ---
	s1 := make([]int, 3) // len = 3, cap = 3
	fmt.Println("\n--- Slice s1 (make with only length) ---")
	fmt.Println("Slice:", s1)
	fmt.Println("Length:", len(s1))
	fmt.Println("Capacity:", cap(s1))

	s2 := make([]int, 3, 5) // len = 3, cap = 5
	fmt.Println("\n--- Slice s2 (make with length & capacity) ---")
	fmt.Println("Slice:", s2)
	fmt.Println("Length:", len(s2))
	fmt.Println("Capacity:", cap(s2))



	// --- [5] Empty slice + append demonstration ---
	var s3 []int
	fmt.Println("\n--- Slice s3 (empty + append) ---")
	fmt.Printf("Initial: len=%d cap=%d slice=%v\n", len(s3), cap(s3), s3)

	for i := 1; i <= 12; i++ {
		s3 = append(s3, i)
		fmt.Printf("After appending %2d → len=%2d cap=%2d slice=%v\n", i, len(s3), cap(s3), s3)
	}


/*
Go Slice Internal Properties

Every slice in Go maintains three core properties:
	Pointer — Points to the first element in the underlying array (that is accessible through the slice).
	Length (len) — The number of elements currently in the slice.
	Capacity (cap) — The maximum number of elements that can be held in the slice before a new underlying array must be allocated.


Behavior When Appending
	When you use the append() function, Go checks if the slice has enough capacity to add the new elements:

	Case 1: len < cap
		The new element is added to the existing underlying array.
		len increases by 1 (or by the number of appended elements).
		cap remains unchanged.
		The Pointer still points to the same array.
	Case 2: len == cap
		Go must allocate a new underlying array.	
		All existing elements are copied into this new array.
		The new element(s) are added after copying.
		The Pointer now points to the new array.

Capacity Growth Rule in Go
When the capacity is exhausted:
	If the current capacity ≤ 1024
		→ New capacity = 2 × old capacity
	If the current capacity > 1024
		→ New capacity = old capacity + (old capacity / 4)
		→ (i.e., increase by 25%)

*/