package main

import (
	"fmt"

	custom_package "github.com/basharulalammazu/Go-Learning-Path/15.custom_package"
)

func main() {
	a := 10
	b := 20
	result := custom_package.Add(a, b)
	fmt.Println("The sum is:", result)
}

