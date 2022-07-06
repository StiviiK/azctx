package log

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

// InitLogger initializes a new logger
// Initialization must be done, before logging funcs can be called
func InitLogger(output io.Writer) {
	logger = logrus.New()
	logger.SetOutput(output)
}

// Info prints the supplied format string using the Info logger
func Info(format string, v ...interface{}) {
	logger.Infof(format, v...)
}

// Warn prints the supplied format string using the Warn logger
func Warn(format string, v ...interface{}) {
	logger.Warnf(format, v...)
}

// Error prints the supplied format string using the Error logger
func Error(format string, v ...interface{}) {
	logger.Errorf(format, v...)
}

func init() {
	InitLogger(os.Stderr)
}
