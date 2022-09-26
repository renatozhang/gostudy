package main

import (
	"fmt"
	"time"
)

func main() {
	var i int
	for {
		var curTime time.Time

		time.Sleep(time.Second * 5)
		i++
		curTime = time.Now()
		fmt.Printf("run %d count, cur time:%v\n", i, curTime)
	}
}
