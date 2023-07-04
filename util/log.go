package util

import (
	"os"

	"gitea.com/lunny/log"
)

// Logger 应用日志对象
var (
	Logger     *log.Logger
	FileWriter *log.Files
)

func init() {
	//init log
	Logger = log.New(os.Stderr, "", log.Ldefault())
	FileWriter = log.NewFileWriter(log.FileOptions{
		ByType: log.ByDay,
		Dir:    "./logs",
	})
	Logger.SetOutput(FileWriter)
}

// Info ...
func Info(v ...interface{}) {
	Logger.Info(v...)
}

// Infof ...
func Infof(format string, v ...interface{}) {
	Logger.Infof(format, v...)
}

// Debug ...
func Debug(v ...interface{}) {
	Logger.Debug(v...)
}

// Debugf ...
func Debugf(format string, v ...interface{}) {
	Logger.Debugf(format, v...)
}

// Warn ...
func Warn(v ...interface{}) {
	Logger.Warn(v...)
}

// Warnf ...
func Warnf(format string, v ...interface{}) {
	Logger.Warnf(format, v...)
}

// Error ...
func Error(v ...interface{}) {
	Logger.Error(v...)
}

// Errorf ...
func Errorf(format string, v ...interface{}) {
	Logger.Errorf(format, v...)
}

// Fatal ...
func Fatal(v ...interface{}) {
	Logger.Fatal(v...)
}

// Fatalf ...
func Fatalf(format string, v ...interface{}) {
	Logger.Fatalf(format, v...)
}

// Panic ...
func Panic(v ...interface{}) {
	Logger.Panic(v...)
}

// Panicf ...
func Panicf(format string, v ...interface{}) {
	Logger.Panicf(format, v...)
}
