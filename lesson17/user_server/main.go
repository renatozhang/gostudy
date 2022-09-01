package main

import (
	"github.com/renatozhang/gostudy/logger"
)

func initLogger(name, logPath, logName string, level int) (err error) {
	m := make(map[string]string)
	m["log_path"] = logPath
	m["log_name"] = logName
	m["log_level"] = logName
	m["log_split_type"] = "size"
	err = logger.InitLogger(name, m)
	if err != nil {
		return
	}
	// log = logger.NewConsoleLogger(level)
	logger.Debug("init logger sucess")
	return
}

func run() {
	for {
		logger.Debug("user server is running /app/zhangzeng/goproject/src/github.com/renatozhang/gostudy/lesson17")
		// time.Sleep(time.Second)
	}

}

func main() {
	/*
		file := log.NewFileLog("c:/a.log")
		file.LogDebug("this is a debug log")
		file.LogDebug("this is a warn log")
	*/
	/*
		console := log.NewConsoleLog("xxxx")
		console.LogConsoleDebug("this is a console log")
		console.LogConsoleWarn("this is a warn log")
	*/
	// log := log.NewFileLog("c:/a.log")
	// log := log.NewConsoleLog("c:/a.log")
	// log.LogDebug("this is debug file")
	// log.LogWarn("this is a warn log")

	initLogger("file", "/home/zhangzeng/log", "user_server", logger.LogLevelDebug)
	run()
}
