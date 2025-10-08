package main

import "fmt"

func main() {
	a := 15

	address := &a // a => ampersand operator => address of operator
	fmt.Println("Address of a:", address)
	fmt.Println("Value of a using pointer:", *address) // \* => value at address


	*address = 20 // Modify the value at the address
	fmt.Println("New value of a:", a)

	x := 10
	y := 25
	swap(&x, &y)
	fmt.Println("Swapped values:", x, y)
}


func swap(x, y *int) {
	temp := *x
	*x = *y
	*y = temp
}