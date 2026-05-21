package logger

import log4go "github.com/jeanphorn/log4go"

func InitLogger() {
	log4go.LoadConfiguration("logger/logger.json")
}

func Info(format string, args ...interface{}) {
	log4go.Info(format, args...)
}

func Warn(format string, args ...interface{}) {
	log4go.Warn(format, args...)
}

func Error(format string, args ...interface{}) {
	log4go.Error(format, args...)
}