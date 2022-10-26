package main

import (
	"fmt"
	"time"
)

func hello(c chan bool) {
	fmt.Println("hello goroutine")
	time.Sleep(5 * time.Second)
	c <- true
}

func main() {
	var exitchan chan bool
	exitchan = make(chan bool)
	go hello(exitchan)
	fmt.Println("main thread terminate")
	<-exitchan
}
