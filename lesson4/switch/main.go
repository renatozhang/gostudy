package main

import "fmt"

func testIf() {
	var a int = 2
	if a == 1 {
		fmt.Printf("a=1\n")
	} else if a == 2 {
		fmt.Printf("a=2\n")
	} else if a == 3 {
		fmt.Printf("a=3\n")
	} else if a == 4 {
		fmt.Printf("a=4\n")
	} else {
		fmt.Printf("a=5")
	}
}

func testSwitch() {
	a := 3
	switch a {
	case 1:
		fmt.Println("a=1")
	case 2:
		fmt.Printf("a=2\n")
	case 3:
		fmt.Printf("a=3\n")
	case 4:
		fmt.Printf("a=4\n")
	case 5:
		fmt.Printf("a=5\n")
	}

}

func getValue() int {
	return 6
}
func testSwitchV2() {
	switch a := getValue(); a {
	case 1:
		fmt.Println("a=1")
	case 2:
		fmt.Printf("a=2\n")
	case 3:
		fmt.Printf("a=3\n")
	case 4:
		fmt.Printf("a=4\n")
	case 5:
		fmt.Printf("a=5\n")
	default:
		fmt.Printf("invalid a=%d\n", a)
	}
}

func testSwitchV3() {

	switch a := getValue(); a {
	case 1, 2, 3, 4, 5:
		fmt.Println("a>=1 and a<=5")
	case 6, 7, 8, 9, 10:
		fmt.Printf("a>=6 and a<=10\n")
	default:
		fmt.Printf("a>10\n")
	}
}

func testSwitchV4() {
	var num = 110
	switch {
	case num >= 0 && num <= 25:
		fmt.Println("num>=0 and num<=25")
	case num > 25 && num <= 50:
		fmt.Printf("num>25 and num<=50\n")
	case num > 50 && num <= 75:
		fmt.Printf("num>50 and num<=75\n")
	case num > 75 && num <= 100:
		fmt.Printf("num>75 and num<=100\n")
	default:
		fmt.Printf("invalid num=%d\n", num)
	}
}

func testSwitchV5() {
	var num = 60
	switch {
	case num >= 0 && num <= 25:
		fmt.Println("num>=0 and num<=25")
	case num > 25 && num <= 50:
		fmt.Printf("num>25 and num<=50\n")
	case num > 50 && num <= 75:
		fmt.Printf("num>50 and num<=75\n")
		fallthrough
	case num > 75 && num <= 100:
		fmt.Printf("num>75 and num<=100\n")
	default:
		fmt.Printf("invalid num=%d\n", num)
	}
}

func testMulti() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, j*i)
		}
		fmt.Println()
	}
}

func main() {
	// testIf()
	// testSwitch()
	// testSwitchV2()
	// testtestSwitchV3()

	// testSwitchV4()
	// testSwitchV5()

	testMulti()

}
