package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var x int32
var wg sync.WaitGroup

var mutex sync.Mutex

func addMutex() {
	for i := 0; i < 50000; i++ {
		mutex.Lock()
		x += 1
		mutex.Unlock()
	}
	wg.Done()
}

func add() {
	for i := 0; i < 500000; i++ {
		// mutex.Lock()
		// x += 1
		// mutex.Unlock()
		atomic.AddInt32(&x, 1)
	}
	wg.Done()
}

func main() {
	start := time.Now().UnixNano()
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go add()
		// go addMutex()
	}

	wg.Wait()
	end := time.Now().UnixNano()
	cost := (end - start) / 1000 / 1000
	fmt.Println("x:", x, "cost:", cost, "ms")
}
