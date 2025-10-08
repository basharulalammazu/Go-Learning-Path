package main

import "fmt"


const a = 10 // Constant variable
var p = 100


func outer() (func()) {
	money := 100;
	age := 30;

	fmt.Printf("Outer function: money = %d, age = %d\n", money, age)

	show := func() { // Inner function (closure)[]
		fmt.Printf("Inner function: money = %d, age = %d\n", money, age)
		money += money + a + p
		fmt.Printf("Inner function after modification: money = %d, age = %d\n", money, age)
	}

	return show
}


func call(){
	closureFunc1 := outer() // Call the outer function and get the inner function
	closureFunc1()
	closureFunc1()

	closureFunc2 := outer() // Call the outer function again to get a new inner function
	closureFunc2()
	closureFunc2()


} 

func init(){
	fmt.Println("This is init function.closure.go file")
}

func main(){
	call()
}



/*
2 Phases
  1. Compilation Phase (Compile time
     a. Syntax checking
	 b. Memory allocation
  2. Execution Phase (Run time)
	 a. Variable initialization
	 b. Function execution


command: go build 20.closure.go

====== Code Phase ======
*** Code Segments **
outer = func() (func()) { ... }
outerAnonymous = func() { ... }
init = func() { ... }
main = func() { ... }


command: /.20.closure.go



*/
