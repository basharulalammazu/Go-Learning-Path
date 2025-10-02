/*
	1. Parameters vs Arguments
	2. First Order Function
        i. Named or standard function
		ii. Anonymous function
		iii. IIFE
		iv. Function expression
	3. Higher Order Function


	First work with object(people, animal, car etc), property(color, shape etc), and relation
	Ex: Mazu studies cse (mazu is object, studies is relation, cse is object)

	Rule: All customer must pay their pizza bill before leaving the restaurant
	1. First Order Function work with object and property
	2. Second Order Function work with relation
	Here, Mazu is a student (mazu is object, student is property)


	Higher Order Function work with rule
	1. All customer must pay their pizza bill before leaving the restaurant
	2. If a customer is a student, they get a 10% discount
	3. If a customer is a member, they get a 20% discount


*/

package main

import "fmt"

// Higher Order Function: A function that takes another function as a parameter or returns a function
func processOperation(a int, b int, op func(p int, q int))  { // Parameters => a, b, op
	 op(a, b)
}

func callbackFunction() (func(p int, q int)) {
	return add
}

func add(num1 int, num2 int)  {// Parameters => a, b
	fmt.Println("Sum:", num1 + num2)
}

func main(){
	processOperation(5, 10, add) // Arguments => 5, 10

	// Callback function
	sum := callbackFunction()
	sum(10, 20)
}