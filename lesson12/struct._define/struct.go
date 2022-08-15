package main

import "fmt"

type User struct {
	Username  string
	Sex       string
	Age       int
	AvatarUrl string
}

func main() {
	var user User
	user.Age = 18
	user.AvatarUrl = "http://www.baidu.com/image/xxx.jpg"
	user.Sex = "男"
	user.Username = "user01"

	fmt.Printf("user.username=%s age=%d sex=%s avatar=%s\n", user.Username, user.Age, user.Sex, user.AvatarUrl)

	var user2 User = User{
		Username: "user02",
		// Age:      18,
		Sex: "女",
		// AvatarUrl: "http://www.baidu.com/image/s.jpg",
	}
	fmt.Printf("user2=%#v\n", user2)

	user3 := User{
		Username: "user03",
		Age:      19,
		Sex:      "男",
		// AvatarUrl: "http://www.baidu.com/image/s.jpg",
	}
	fmt.Printf("user3=%#v\n", user3)

	var user4 User
	fmt.Printf("user4=%#v\n", user4)
}
