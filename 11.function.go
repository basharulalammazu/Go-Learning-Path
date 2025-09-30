package main
import "fmt"


func add(num1 int, num2 int){

	sum := num1 + num2
	fmt.Println(sum)



	// or you can do it like this also
	// without using a variable
	// fmt.Println(num1 + num2)
}

func main() {
	a := 10
	b := 20

	add(a, b)

	add(100, 200)
}