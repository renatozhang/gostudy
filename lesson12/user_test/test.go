package main

import (
	"fmt"

	"github.com/renatozhang/gostudy/lesson12/user"
)

func main() {
	// var u user.User
	// u.Age = 10
	// fmt.Printf("user=%#v\n", u)

	u := user.NewUser("user01", "女", 18, "xxx.jpg")
	// u.age = 100   小写私有不能在包外被引用
	fmt.Printf("user=%#v\n", u)
}
