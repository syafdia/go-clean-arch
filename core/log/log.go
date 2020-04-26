package log

import (
	"fmt"
	"log"
	"runtime"
)

func Debug(msg string, args ...interface{}) {
	printLog("debug", msg, args)
}

func Info(msg string, args ...interface{}) {
	printLog("info", msg, args)
}

func Error(msg string, args ...interface{}) {
	printLog("error", msg, args)
}

func printLog(level string, msg string, args ...interface{}) {
	_, file, line, ok := runtime.Caller(1)

	if ok {
		caller := fmt.Sprintf("%s:%d", file, line)
		log.Printf("[LOGGER-%s] %s -> %+v", level, caller, args)
		return
	}

	log.Printf("[LOGGER-%s] %+v", level, args)
}
