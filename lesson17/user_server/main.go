package main

import "github.com/renatozhang/gostudy/lesson17/log"

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
	log := log.NewConsoleLog("c:/a.log")
	log.LogDebug("this is debug file")
	log.LogWarn("this is a warn log")
}
