package main

import "fmt"

func modifyInt(a *int) {
	*a = 1000
}

func main() {
	var a int = 100
	fmt.Printf("befor modifyInt:%d addr:%p\n", a, &a)
	modifyInt(&a)
	fmt.Printf("befor modifyInt:%d addr:%p\n", a, &a)
}
