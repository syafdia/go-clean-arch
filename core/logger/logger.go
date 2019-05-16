package logger

import (
	"fmt"
	"log"
	"runtime"
)

func Debug(data interface{}) {
	printLog("debug", data)
}

func Info(data interface{}) {
	printLog("info", data)
}

func Error(data interface{}) {
	printLog("error", data)
}

func printLog(level string, data interface{}) {
	_, file, line, ok := runtime.Caller(1)

	if ok {
		caller := fmt.Sprintf("%s:%d", file, line)
		log.Printf("[LOGGER-%s] %s -> %+v", level, caller, data)
		return
	}

	log.Printf("[LOGGER-%s] %+v", level, data)
}
