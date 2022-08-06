package main

import "fmt"

const (
	A = iota
	B = iota
	C = iota
	D = 8
	E
	F = iota
	G
)

const (
	A1 = iota
	A2
)

func main() {
	fmt.Println(A) //0
	fmt.Println(B) //1
	fmt.Println(C) //2
	fmt.Println(D) //8
	fmt.Println(E) //8
	fmt.Println(F) //5
	fmt.Println(G) //6

	fmt.Println("A1A2")
	fmt.Println(A1)
	fmt.Println(A2)
}
