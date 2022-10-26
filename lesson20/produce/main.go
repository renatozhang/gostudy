package main

import "fmt"

func produce(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}

func main() {
	ch := make(chan int)
	go produce(ch)
	for {
		v, ok := <-ch
		if ok == false {
			fmt.Println("chan is closed")
			break
		}
		fmt.Println("received ", v, ok)
	}
}
