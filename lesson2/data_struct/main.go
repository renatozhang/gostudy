package main

import (
	"fmt"
	"strings"
)

func testBool() {
	var a bool
	fmt.Println(a)
	a = true
	fmt.Println(a)

	a = !a
	fmt.Println(a)

	var b bool = true

	if a == true && b == true {
		fmt.Println("right")
	} else {
		fmt.Println("not right")
	}

	if a == true || b == true {
		fmt.Println("|| right")
	} else {
		fmt.Println("|| not right")
	}

	fmt.Printf("%t %t\n", a, b)
}

func testInt() {
	var a int8
	a = 18
	fmt.Printf("a=%d\n", a)
	a = -12
	fmt.Printf("a=%d\n", a)

	var b int
	b = 1225456665
	b = int(a)
	fmt.Println("b=", b)

	// var c float32 = 5.6
	var c = 5.6
	fmt.Println("c=", c)

	fmt.Printf("a=%d b=%x c=%f", a, b, c)
}

func testStr() {
	// var a string = "hello"
	var a = "a:hello"
	fmt.Println(a)

	b := a
	fmt.Println(b)

	c := "\nc:hello"
	fmt.Println(c)
	// c := "x"

	fmt.Printf("a-%v b-%v c-%v\n", a, b, c)
	c = `c:\nhello
	abc
	dcd
	elsfs

	sfdsfa
	`
	fmt.Println(c)

	clen := len(c)
	fmt.Printf("len of c = %d\n", clen)

	fmt.Println("===============")
	c = c + c
	fmt.Printf("c=%s\n", c)
	c = fmt.Sprintf("%s%s\n", c, c)
	fmt.Printf("c=%s", c)

	fmt.Println("===============")
	ips := "10.108.34.30,10.108.34.31"

	ipArry := strings.Split(ips, ",") //字符串切割

	fmt.Printf("first ip :%s\n", ipArry[0])
	fmt.Printf("second ip :%s\n", ipArry[1])
	result := strings.Contains(ips, "10.108.34.30")
	fmt.Println(result)

	// 前缀判断
	str := "http://baidubaidu.com"
	if strings.HasPrefix(str, "http") {
		fmt.Println("str is http url")
	} else {
		fmt.Println("str not is http url")
	}

	if strings.HasSuffix(str, "baidu.com") {
		fmt.Println("str is baidu url")
	} else {
		fmt.Println("str not is baidu url")
	}

	index := strings.Index(str, "baidu")

	fmt.Printf("baidu is index:%d\n", index)

	index = strings.LastIndex(str, "baidu")
	fmt.Printf("baidu is index:%d\n", index)

	var strArr []string = []string{"10.108.34.30", "10.108.34.31", "10.108.34.32"}
	resultStr := strings.Join(strArr, ",")
	fmt.Printf("result=%s\n", resultStr)

}

func testOperator() {
	a := 2
	if a != 2 {
		fmt.Printf("is right\n")
	} else {
		fmt.Printf("not is right\n")
	}

	a = a + 100
	fmt.Println(a)

}

func testAccess() {
	fmt.Printf("access.a=%d\n", access.a)
}
func main() {
	//testBool()
	// testInt()
	// testStr()
	// testOperator()

	testAccess()

}
