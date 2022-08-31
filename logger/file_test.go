package logger

import "testing"

func TestFileLogger(t *testing.T) {
	logger := NewFileLogger(LogLevelDebug, "/home/zhangzeng/log", "test")
	logger.Debug("user id[%d] is conme from china", 324234)
	logger.Fatal("test warn log")
	logger.Info("test warn log")
	logger.Warn("test warn log")
	logger.Error("test error log")
	logger.Fatal("test Fatal log")
}
