package logger

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"
)

const logTimeLayout = "02-01-2006_03:04:05"

const outDir = "out/log"

var errOpen = errors.New("couldn't open log file")

// InitLogger initiate auth service logger
func InitLogger() *log.Logger {
	logFileName := fmt.Sprintf("%s/auth_%s.log", outDir, time.Now().Format(logTimeLayout))

	out, err := os.OpenFile(logFileName, os.O_CREATE|os.O_RDWR, 0600) //nolint: gosec
	if err != nil {
		log.Fatalf("%v - %v: %v", errOpen, outDir, err)
	}

	return log.New(out, "", log.Lmsgprefix|log.Ldate|log.Ltime|log.Lshortfile)
}
