package logger

import (
	"fmt"
	"os"
)

type FileLogger struct {
	level    int
	logPath  string
	logName  string
	file     *os.File
	warnFile *os.File
}

func (f *FileLogger) init() {
	filename := fmt.Sprintf("%s/%s.log", f.logPath, f.logName)
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(fmt.Sprintf("open firl %s failed, err:%v", filename, err))
	}
	f.file = file
	//写错误日志和fatal日志的文件
	filename = fmt.Sprintf("%s/%s.log.wf", f.logPath, f.logName)
	file, err = os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(fmt.Sprintf("open firl %s failed, err:%v", filename, err))
	}
	f.warnFile = file
}

func NewFileLogger(level int, logPath, logName string) LogInterface {
	logger := &FileLogger{
		level:   level,
		logPath: logPath,
		logName: logName,
	}
	logger.init()
	return logger
}

func (f *FileLogger) SetLevel(level int) {
	if level <= LogLevelDebug || level >= LogLevelFatal {
		level = LogLevelDebug
	}
	f.level = level
}

func (f *FileLogger) Debug(format string, args ...interface{}) {
	if f.level > LogLevelDebug {
		return
	}
	fmt.Fprintf(f.file, format, args...)
	fmt.Fprintln(f.file)
}
func (f *FileLogger) Trace(format string, args ...interface{}) {
	if f.level > LogLevelDebug {
		return
	}
	fmt.Fprintf(f.file, format, args...)
	fmt.Fprintln(f.file)
}
func (f *FileLogger) Info(format string, args ...interface{}) {
	if f.level > LogLevelDebug {
		return
	}
	fmt.Fprintf(f.file, format, args...)
	fmt.Fprintln(f.file)
}
func (f *FileLogger) Warn(format string, args ...interface{}) {
	if f.level > LogLevelDebug {
		return
	}
	fmt.Fprintf(f.warnFile, format, args...)
	fmt.Fprintln(f.warnFile)
}
func (f *FileLogger) Error(format string, args ...interface{}) {
	if f.level > LogLevelDebug {
		return
	}
	fmt.Fprintf(f.warnFile, format, args...)
	fmt.Fprintln(f.warnFile)
}
func (f *FileLogger) Fatal(format string, args ...interface{}) {
	if f.level > LogLevelDebug {
		return
	}
	fmt.Fprintf(f.warnFile, format, args...)
	fmt.Fprintln(f.warnFile)
}

func (f *FileLogger) Close() {
	defer f.file.Close()
	defer f.warnFile.Close()
}
