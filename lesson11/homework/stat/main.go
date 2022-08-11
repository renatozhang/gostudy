package main

import (
	"fmt"
	"strings"
)

func statWords(str string) map[string]int {
	var result map[string]int = make(map[string]int, 128)
	words := strings.Split(str, " ")
	for _, val := range words {
		_, ok := result[val]
		if !ok {
			result[val] = 1
		} else {
			result[val] += 1
		}
	}
	return result
}

func main() {
	var str = "how do you do ? do you like me ?"
	result := statWords(str)
	fmt.Printf("result:%#v\n", result)
}
