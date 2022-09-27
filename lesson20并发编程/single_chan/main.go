package main

import "fmt"

func sendData(sendch chan<- int) {
	sendch <- 10
	// <- sendch
}

func readData(sendch <-chan int) {
	// sendch <- 10
	data := <-sendch
	fmt.Println(data)
}

func main() {
	chn1 := make(chan int)
	go sendData(chn1)
	readData(chn1)
}
