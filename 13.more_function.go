package main
import "fmt"

// Collect the name from user and print a greeting message
func greetUser(name string){
	fmt.Println("Hello", name, "Welcome to Go programming!")
}


// 
func getNumbers(num1 int, num2 int) (int, int){
	return (num1 + num2), (num1 * num2)
}


// main function
func main(){
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	greetUser(name)


	var a int
	var b int
	fmt.Print("Enter two numbers: ")
	fmt.Scanln(&a, &b)
	sum, product := getNumbers(a, b)
	fmt.Println("The sum is:", sum)
	fmt.Println("The product is:", product)
}