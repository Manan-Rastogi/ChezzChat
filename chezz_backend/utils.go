package main

import (
	"gopkg.in/natefinch/lumberjack.v2"
)

func generateGinLoggerFile() *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   "./gin.log",
		MaxSize:    20, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
	}
}
