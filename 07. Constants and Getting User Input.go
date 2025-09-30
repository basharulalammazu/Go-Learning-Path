package main
import "fmt"
  
func main(){
  var name string
  const COUNTRY = "Bangladesh"
  var age, num1, num2 int
  var gpa float32

	fmt.Printf("Enter your name: ")
	fmt.Scanf("%v", &name)

	fmt.Printf("Enter your age: ")
	fmt.Scan(&age)

	fmt.Printf("Enter your SSC gpa: ")
	fmt.Scanln(&gpa)

	fmt.Printf("Enter 2 numbers: ")
	fmt.Scan(&num1, &num2)

  fmt.Printf("%v is a student\n", name)
  fmt.Printf("%v is %v years old\n", name, age)
  fmt.Printf("%v has got GPA %v/5 in SSC\n", name, gpa)
  fmt.Printf("%v originally from %v\n", name, COUNTRY);
  fmt.Printf("num1 = %v, num2 = %v\n", num1, num2);
}
