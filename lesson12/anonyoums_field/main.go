package main

import "fmt"

type User struct {
	Username string
	Sex      string
	int
	string
}

func main() {
	var user User
	user.Username = "user01"
	user.Sex = "man"
	user.int = 100
	user.string = "hello"

	fmt.Printf("user=%#v\n", user)
}
