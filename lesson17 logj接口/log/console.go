package log

import "fmt"

type ConsoleLog struct {
}

func NewConsoleLog(file string) ConsoleLog {
	return ConsoleLog{}
}

func (f ConsoleLog) LogDebug(msg string) {
	fmt.Println("file", msg)
}

func (f ConsoleLog) LogWarn(msg string) {
	fmt.Println("file", msg)
}
