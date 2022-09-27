package main

import (
	"time"
)

var i int

func calc() {
	for {
		i++
	}
}

func main() {
	/*
		cpu := runtime.NumCPU()
		fmt.Println("cpu:", cpu)

		runtime.GOMAXPROCS(4)
	*/
	for i := 0; i < 10; i++ {
		go calc()
	}

	time.Sleep(time.Minute)
}
