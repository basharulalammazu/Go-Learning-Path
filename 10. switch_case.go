package main 

import "fmt"

func main(){
	age := 18


	// Basic switch case
	switch age {
		case 18:
			fmt.Println("You are 18 years old")
		case 20:
			fmt.Println("You are 20 years old")
		default:
			fmt.Println("Your age is neither 18 nor 20")
		}


		
	// Also use conditions in switch case
	switch {
		case age > 18:
			fmt.Println("You are eligible for marriage")
		case age < 18:
			fmt.Println("You are not eligible for marriage")
		default:
			fmt.Println("You are 18 years old")
	} 
}