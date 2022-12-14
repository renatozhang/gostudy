package logger

import "testing"

func TestFileLogger(t *testing.T) {
	logger := NewFileLogger(LogLevelDebug, "/home/zhangzeng/log", "test")
	logger.Debug("user id[%d] is conme from china", 324234)
	logger.Trace("test Trace log")
	logger.Info("test Info log")
	logger.Warn("test warn log")
	logger.Error("test error log")
	logger.Fatal("test Fatal log")
}

func TestConsoleLogger(t *testing.T) {
	logger := NewConsoleLogger(LogLevelDebug)
	logger.Debug("user id[%d] is conme from china", 324234)
	logger.Trace("test Trace log")
	logger.Info("test Info log")
	logger.Warn("test warn log")
	logger.Error("test error log")
	logger.Fatal("test Fatal log")
}
