package main

import (
	"fmt"
	"time"
)

func calc() {
	time.Sleep(1 * time.Second)
	for i := 0; i < 10; i++ {
		fmt.Println("run ", i, "times")

	}
	fmt.Println("calc finish")
}

func main() {
	go calc()
	fmt.Println("1 exited")
	time.Sleep(11 * time.Second)
}
