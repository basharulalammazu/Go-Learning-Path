package main

import "fmt"

type User struct {
	Name  string // Member veriable or field or property
	Age int
}


func printUserInfo(user User){
	fmt.Println("User Name:", user.Name)
	fmt.Println("User Age:", user.Age)
}

func main() {
	var user1 User
	user1 = User{
		Name: "Basharul",
		Age: 22,
	}
	printUserInfo(user1)

	user2 := User{ // Instance or object of User type
		Name: "Mazu",
		Age: 30,
	}
	printUserInfo(user2)
}
/*
Struct: A collection of fields
1. Define a struct type
2. Create struct instances (objects)
3. Access and modify struct fields

4. Use struct in functions
5. Anonymous struct
6. Nested struct
7. Struct with methods
8. Pointer to struct
9. Struct tags
10. JSON and struct


go run 21.struct.go
*/