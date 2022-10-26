package log

import "fmt"

type Filelog struct {
}

func NewFileLog(file string) Filelog {
	return Filelog{}
}

func (f Filelog) LogDebug(msg string) {
	fmt.Println("file", msg)
}

func (f Filelog) LogWarn(msg string) {
	fmt.Println("file", msg)
}
