package main
import "fmt"


func add(num1 int, num2 int) int{
	sum := num1 + num2
	return sum
}


func getNumbers(num1 int, num2 int) (int, int){
	return (num1 + num2), (num1 * num2)
}


// 
func sayHello(){
	fmt.Println("Hello")
}

func sayGoodbye(name string){
	fmt.Println("Goodbye", name)
}

func main() {
	a := 10
	b := 20

	result := add(a, b)
	fmt.Println("The sum is:", result)

	fmt.Println("The sum is:", add(100, 200))

	sum, product := getNumbers(a, b)
	fmt.Println("The sum is:", sum)
	fmt.Println("The product is:", product)


	sayHello()
	sayGoodbye("Mazu")
}