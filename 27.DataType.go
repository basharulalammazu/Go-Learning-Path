// 27.DataType.go

/*
Data Types in Go

1. Basic Data Types
	(Use signed integers for both positive and negative values)
   - int -> 32 or 64 bit depending on the system
   - int8 -> 8 bit -> -128 to 127
   - int16 -> 16 bit -> -32,768 to 32,767
   - int32 -> 32 bit -> -2,147,483,648 to 2,147,483,647
   - int64 -> 64 bit -> -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807

    (Use unsigned integers for non-negative values)
   - uint -> 32 or 64 bit depending on the system
   - uint8 -> 8 bit -> 0 to 255
   - uint16 -> 16 bit -> 0 to 65,535
   - uint32 -> 32 bit -> 0 to 4,294,967,295
   - uint64 -> 64 bit -> 0 to 18,446,744,073,709,551,615
   - uintptr -> an unsigned integer large enough to store the uninterpreted bits of a pointer value

   (Use floating-point numbers for fractional values)
   - float32 -> 32 bit -> 1.5 x 10^-45 to 3.4 x 10^38
   - float64 -> 64 bit -> 5.0 x 10^-324 to 1.7 x 10^308

   (Use complex numbers for values with both real and imaginary parts)
   - complex64 -> 64 bits (32 bits real + 32 bits imaginary)
   - complex128 -> 128 bits (64 bits real + 64 bits imaginary)

   (Other basic types)
   - byte -> 8 bits -> alias for uint8
   - rune -> 32 bits -> alias for int32 (represents a Unicode code point)

   - bool -> 8 bits - false or true
   - string -> variable bit
*/

package main

import "fmt"

func main() {
	var a int8 = -128
	var b uint8 = 255
	var c float32 = 3.14
	var d complex64 = 1 + 2i
	var e byte = 'A'
	var f rune = '💖'
	var flag bool = true
	var s string = "Hello, Go!"

	fmt.Println(f) // Print rune as character
	fmt.Printf("%c", f) // Print rune as character using format specifier

	// Displaying values using different methods
	/*
	   Format Specifiers:
	   %d -> integer
	   %f -> floating-point
	   %c -> character
	   %t -> boolean
	   %s -> string
	*/	
	fmt.Printf("%d\n", a)
	fmt.Printf("%d\n", b)
	fmt.Printf("%2f\n", c)
	fmt.Printf("%f\n", d)
	fmt.Printf("%c\n", e)
	fmt.Printf("%c\n", f)
	fmt.Printf("%t\n", flag)
	fmt.Printf("%s\n", s)

	fmt.Println()
	fmt.Println(a, b, c, d, e, f, flag, s)

	fmt.Printf("Type of a: %T\n", a)
	fmt.Printf("Type of b: %T\n", b)
	fmt.Printf("Type of c: %T\n", c)
	fmt.Printf("Type of d: %T\n", d)
	fmt.Printf("Type of e: %T\n", e)
	fmt.Printf("Type of f: %T\n", f)
	fmt.Printf("Type of flag: %T\n", flag)
	fmt.Printf("Type of s: %T\n", s)
}