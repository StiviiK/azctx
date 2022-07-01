package log

import (
	"io"
	"log"
	"os"
)

var loggerI *log.Logger
var loggerW *log.Logger
var loggerE *log.Logger

// InitLogger initializes a new logger
// Initialization must be done, before logging funcs can be called
func InitLogger(info, warn io.Writer, err io.Writer) {
	loggerI = log.New(info, "I: ", 0)
	loggerW = log.New(warn, "W: ", 0)
	loggerE = log.New(err, "E: ", 0)
}

// Info prints the supplied format string using the Info logger
func Info(format string, v ...interface{}) {
	loggerI.Printf(format, v...)
}

// Warn prints the supplied format string using the Warn logger
func Warn(format string, v ...interface{}) {
	loggerW.Printf(format, v...)
}

// Error prints the supplied format string using the Error logger
func Error(format string, v ...interface{}) {
	loggerE.Printf(format, v...)
}

func init() {
	InitLogger(os.Stderr, os.Stderr, os.Stderr)
}
