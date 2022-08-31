package main

import "fmt"

func test(a interface{}) {
	s, ok := a.(int)
	if ok {
		fmt.Println(s)
		return
	}
	str, ok := a.(string)
	if ok {
		fmt.Println(str)
		return
	}

	f, ok := a.(float32)
	if ok {
		fmt.Println(f)
		return
	}
	fmt.Println("catn not define the type of a")

}

func testInterface1() {
	var a int = 100
	test(a)

	var b string = "hello"
	test(b)

	var c float32 = 12.34
	test(c)

	var d bool
	test(d)
}

func testSwitch(a interface{}) {
	switch a.(type) {
	case string:
		fmt.Printf("a is string, value:%v\n", a.(string))
	case int:
		fmt.Printf("a is int, value:%v\n", a.(int))
	case int32:
		fmt.Printf("a is int32, value:%v\n", a.(int32))
	default:
		fmt.Println("not support type")
	}
}

func testSwitch2(a interface{}) {
	switch v := a.(type) {
	case string:
		fmt.Printf("a is string, value:%v\n", v)
	case int:
		fmt.Printf("a is int, value:%v\n", v)
	case int32:
		fmt.Printf("a is int32, value:%v\n", v)
	default:
		fmt.Println("not support type")
	}
}

func testInterface2() {
	var a int = 100
	testSwitch(a)
}

func testInterface3() {
	var a int = 100
	testSwitch2(a)
}

func main() {
	// testInterface1()
	// testInterface2()
	testInterface3()
}
