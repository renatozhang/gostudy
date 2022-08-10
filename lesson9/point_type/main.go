package main

import "fmt"

func testPoint1() {
	var a int32
	a = 100
	fmt.Printf("the add of a :%p, a:%d\n", &a, a)

	var b *int32
	fmt.Printf("the addr of b: %p, b:%v\n", &b, b)
	// *b = 100
	if b == nil {
		fmt.Println("b is nil addr\n")
	}
	b = &a
	fmt.Printf("the addr of b:%p, b:%v\n", &b, b)
}

func testPoint2() {
	var a int = 200
	var b *int = &a

	fmt.Printf("b 指向地址存储的值为为:%d \n", *b)
	*b = 1000
	fmt.Printf("b指向地址存储的值为:%d\n", *b)
	fmt.Printf("a = %d\n", a)
}

func change(val *int) {
	*val = 55
}

func testPoint3() {
	a := 58
	fmt.Println("value of a befor function call is", a)
	p := &a
	change(p)
	fmt.Println("value of after function call is", a)
}

func modify_arr(a *[3]int) {
	(*a)[0] = 100
}

func testPoint4() {
	var b [3]int = [3]int{1, 2, 3}
	modify_arr(&b)
	fmt.Printf("b:%v\n", b)
}

func modify_slice(a []int) {
	a[0] = 90
}

func testPoint5() {
	a := [3]int{89, 90, 91}
	modify_slice(a[:])
	fmt.Println(a)
}

func testPoint6() {
	var a *int = new(int)
	*a = 100
	fmt.Printf("*a=%d\n", *a)

	var b *[]int = new([]int)
	fmt.Printf("*b = %v\n", *b)
	(*b) = make([]int, 5, 100)
	(*b)[0] = 100
	(*b)[1] = 200
	fmt.Printf("*b = %v\n", *b)
}

func modifyInt(a *int) {
	*a = 1000
}

func testPoint7() {
	var b int = 10
	modifyInt(&b)
	fmt.Printf("b=%d\n", b)
}

func testPoint8() {
	var a int = 100
	var b *int = &a
	var c *int = b
	*c = 200
	fmt.Printf("a:%d, *b:%d,*c:%d", a, *b, *c)
}

func main() {
	// testPoint1()
	// testPoint2()
	// testPoint3()
	// testPoint4()
	// testPoint5()
	// testPoint6()
	testPoint7()
	// testPoint8()
}
