package main

import (
	"fmt"
)

/*
type Address struct {
	Province string
	City     string
}

type User struct {
	Usename string
	Sex     string
	address *Address
}
*/

type Address struct {
	Province   string
	City       string
	CreateTime string
}

type Email struct {
	account    string
	CreateTime string
}

type User struct {
	Username string
	Sex      string
	*Address
}

func test1() {
	var user User
	user.Username = "user01"
	user.Sex = "main"
	// 匿名结构体第一种初始化方式
	user.Address = &Address{
		Province: "bj",
		City:     "bj",
	}
	//第二种方式
	user.Province = "bj01"
	user.City = "bj01"

	fmt.Printf("user=%#v addr=%#v city=%#v\n", user, user.Address, user.City)
}

type User01 struct {
	// City     string
	Username string
	Sex      string
	*Address
	*Email
}

func test2() {
	var user User01
	user.Username = "user01"
	user.Sex = "man"
	user.Address = new(Address)
	user.Email = new(Email)
	user.Address.City = "bj01"
	fmt.Printf("user=%#v city of address:%s\n", user, user.Address.City)

	user.Address.CreateTime = "001"
	user.Email.CreateTime = "002"
	fmt.Printf("user=%#v createtime:%s, %s\n", user, user.Address.CreateTime, user.Email.CreateTime)

}

func main() {
	// test1()
	test2()
}
