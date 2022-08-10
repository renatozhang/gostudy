package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var (
	length  int
	charset string
)

const (
	NumStr  = "0123456789"
	CharStr = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	SpecStr = "+=-@#~,.[]()!%^*$"
)

func parseArgs() {
	flag.IntVar(&length, "l", 16, "-l 生成密码的长度")
	flag.StringVar(&charset, "t", "num",
		`-t 指定密码生成的字符集,num:只是用数字[0-9], 
		char:只是用英文字母[a-zA-Z],
		mix:使用数字和字母，
		advance:使用数字、字母及特殊字符`)
	flag.Parse()
}

func test1() {
	for i := 0; i < len(CharStr); i++ {
		if CharStr[i] != ' ' {
			fmt.Printf("%c", CharStr[i])
		}
	}
}

func generatePassword() string {
	var password []byte = make([]byte, length, length)
	var sourceStr string
	if charset == "num" {
		sourceStr = NumStr
	} else if charset == "char" {
		sourceStr = CharStr
	} else if charset == "mix" {
		sourceStr = fmt.Sprintf("%s%s", NumStr, CharStr)
	} else if charset == "advance" {
		sourceStr = fmt.Sprintf("%s%s%s", NumStr, CharStr, SpecStr)
	} else {
		sourceStr = NumStr
	}
	for i := 0; i < length; i++ {
		index := rand.Intn(len(sourceStr))
		password[i] = sourceStr[index]
	}
	// fmt.Println("souece:", sourceStr)
	return string(password)
}
func main() {
	rand.Seed(time.Now().Unix())
	parseArgs()
	// fmt.Printf("length:%d charset:%s\n", length, charset)
	// test1()
	password := generatePassword()
	fmt.Println(password)
}
