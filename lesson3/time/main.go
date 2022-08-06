package main

import (
	"fmt"
	"time"
)

func testTime() {
	now := time.Now()
	fmt.Printf("current time:%v\n", now)

	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	send := now.Second()
	fmt.Printf("%02d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, send)

	timestamp := now.Unix()
	fmt.Printf("timestamp is %d\n", timestamp)
}

func testTimestamp(timestamp int) {
	timeObj := time.Unix(int64(timestamp), 0)
	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minute := timeObj.Minute()
	send := timeObj.Second()
	fmt.Printf("%02d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, send)
}

func testTicker() {
	ticker := time.Tick(5 * time.Second)
	for i := range ticker {
		fmt.Printf("%v\n", i)
		processTask()
	}
}

func processTask() {
	fmt.Printf("do task\n")

}

func timeConst() {
	fmt.Printf("nano second:%d\n", time.Nanosecond)
	fmt.Printf("micro second:%d\n", time.Microsecond)
	fmt.Printf("milli second:%d\n", time.Millisecond)
	fmt.Printf("second:%d\n", time.Second)

}

func timeFormat() {
	now := time.Now()
	timeStr := now.Format("2006-01-02 15:04:05")
	fmt.Printf("time:%s\n", timeStr)
}

func testTimeFormat() {
	now := time.Now()
	timeStr := now.Format("2006/01/02 15:06:06")
	fmt.Printf("time:%s\n", timeStr)
}

func testCost() {
	start := time.Now().UnixNano()
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond)
	}
	end := time.Now().UnixNano()
	cost := (end - start) / 1000
	fmt.Printf("code cost:%d us", cost)
}
func main() {
	// testTime()
	// timestamp := time.Now().Unix()
	// testTimestamp(int(timestamp))
	// testTicker()
	// timeConst()

	// timeFormat()

	testCost()
}
