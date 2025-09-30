package main

import "fmt"

func main() {
	age := 18

	if age > 18 {
		fmt.Println("You are eligible for marriage")
	} else if age < 18 {
		fmt.Println("You are not eligible for marriage")
	} else {
		fmt.Println("You are 18 years old")
	}
}