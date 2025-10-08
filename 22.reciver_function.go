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
