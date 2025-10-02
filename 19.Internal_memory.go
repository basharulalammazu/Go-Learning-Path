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
*/