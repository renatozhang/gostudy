package main

import "fmt"

func testIf() {
	num := 11
	if num%2 == 0 {
		fmt.Printf("num:%d is even\n", num)
	} else {
		fmt.Printf("num:%d is odd\n", num)
	}
}

func testIf2() {
	num := 10
	if num > 5 && num < 10 {
		fmt.Printf("num:%d is >5 and <10\n", num)
	} else if num >= 10 && num < 20 {
		fmt.Printf("num:%d is >=10 and <20\n", num)
	} else if num >= 20 && num < 30 {
		fmt.Printf("num:%d is >=20 and <30\n", num)
	} else {
		fmt.Printf("num:%d is >=30\n", num)
	}
}

func testIf3() {
	if num := 10; num%2 == 0 {
		fmt.Printf("num:%d is even\n", num)
	} else {
		fmt.Printf("num:%d is odd\n", num)
	}

	// fmt.Printf("num=%d\n", num)  //num只在if语句块内有效
}

func getNum() int {
	return 10
}

func testIf4() {
	if num := getNum(); num%2 == 0 {
		fmt.Printf("num:%d is even\n", num)
	} else {
		fmt.Printf("num:%d is odd\n", num)
	}
}

func main() {
	// testIf()
	// testIf2()
	// testIf3()
	// testIf4()
}
