/*

Internal Memory => Code Segment, Data Segment, Stack, Heap, GC


go run 19.Code_Segment_Data_Segment_Stack_Heap_GC.go


Ram
 |--- Code Segment => All Functions => main(), init(), user-defined functions (I mean all code)
 |--- Data Segment/ Global Memory => Global variable, Static variable 
 |--- Stack => Function Call Stack
 |--- Heap => Dynamic Memory Allocation
	  |--- Garbage Collector (GC)


Stack Frame => allocated memory for each function call
 |--- Function parameters
 |--- Return address
 |--- Local variables


 Stack is faster and Global memory is slower 




 Go Code run by 2 phases
 1. Compile Phase => Code Segment, Data Segment
 2. Execution Phase => Stack, Heap
*/




package main
import "fmt"

const a = 10
var p = 100

func call(){
	add := func(num1 int, num2 int){
		res := num1 + num2
		fmt.Println("Addition is:", res)
	}

	add(15, 6)
	add(p+a)
}


func main(){
	call()

	fmt.Print(a)
}


func init(
	fmt.Println("Hello! I am Init function")
)




/*
    1. Compile Phase => Code Segment, Data Segment
		1.1 Code Segment => All Functions => main(), init(), user-defined functions (I mean all code), constants
		1.2 Data Segment/ Global Memory => Global variable, Static variable
 	2. Execution Phase => Stack, Heap
		2.1 Stack => Function Call Stack
		2.2 Heap => Dynamic Memory Allocation
				2.2.1 Garbage Collector (GC)

	Stack Frame => allocated memory for each function call. Its store reference. Return address, Local variables

	Code Segment store in Binary file
	Binary file store in RAM
	Read code first to last
	Global variable store in Data Segment without any function call and constant
	Function call store in Stack Frame => first init function then main function, then user-defined function
	Anonymous function store in Stack Frame when its called


	code Segment contants are read only, so we can't change it
	Data Segment/ Global Memory variable are read and write both, so we can change it
	
*/