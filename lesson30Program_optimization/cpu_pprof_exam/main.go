package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

func logicCode() {
	var c chan int = make(chan int, 1)
	for {
		select {
		case v := <-c:
			fmt.Printf("read from chan, v:%v\n", v)
		default:
			time.Sleep(time.Second)
		}
	}
}

func main() {
	var isCpuPprof bool
	var isMemPprof bool
	flag.BoolVar(&isCpuPprof, "cpu", false, "turn cpu pprof on")
	flag.BoolVar(&isMemPprof, "mem", false, "trun mem pprof on")
	flag.Parse()

	if isCpuPprof {
		file, err := os.Create("./cpu.pprof")
		if err != nil {
			fmt.Printf("create cpu pprof failed, err:%v\n", err)
			return
		}
		pprof.StartCPUProfile(file)
		defer pprof.StopCPUProfile()
	}

	for i := 0; i < 8; i++ {
		go logicCode()
	}
	time.Sleep(30 * time.Second)
}
