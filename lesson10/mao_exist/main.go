package main

import "fmt"

func main() {
	var a map[string]int
	a = make(map[string]int)
	a["stu01"] = 1000
	a["stu02"] = 2000
	a["stu03"] = 3000
	fmt.Printf("a=%#v\n", a)

	var result int
	var ok bool
	var key string = "stu03"
	result, ok = a[key]
	if !ok {
		fmt.Printf("key %s is not exist\n", key)
	} else {
		fmt.Printf("key %s is %d\n", key, result)
	}
}
