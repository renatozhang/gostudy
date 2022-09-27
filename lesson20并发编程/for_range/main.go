package main

import (
	"fmt"
	"time"
)

func produce(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
		time.Sleep(time.Second)
	}
	close(chnl)
}

func main() {
	ch := make(chan int)
	go produce(ch)
	for i := range ch {
		fmt.Println("received ", i)
	}
}
