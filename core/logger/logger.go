package logger

import (
	"fmt"
	"log"
	"runtime"

	"github.com/pkg/errors"
)

func I(info interface{}) {
	_, file, line, ok := runtime.Caller(1)

	if ok {
		caller := fmt.Sprintf("%s:%d", file, line)
		log.Printf("[LOGGER-info] %s -> %+v", caller, info)
		return
	}

	log.Printf("[LOGGER-info] %+v", info)
}

func E(err error) {
	log.Printf("[LOGGER-error] %+v", errors.Wrap(err, ""))
}
