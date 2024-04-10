package log

import (
	stdlibLogger "log"
	"os"
)

var stdout = stdlibLogger.New(os.Stdout, "", 0)
var stderr = stdlibLogger.New(os.Stderr, "[ERROR]", 0)
var stdwarn = stdlibLogger.New(os.Stderr, "[WARN]", 0)

// Info is like log.Println, but outputs to stdout
func Info(v ...any) {
	stdout.Println(v...)
}

// Infof is like log.Printf, but outputs to stdout
func Infof(format string, v ...any) {
	stdout.Printf(format, v...)
}

// Error is like log.Println, but messages are prefixed with "[ERROR]"
func Error(v ...any) {
	stderr.Println(v...)
}

// Errorf is like log.Printf, but messages are prefixed with "[ERROR]"
func Errorf(format string, v ...any) {
	stderr.Printf(format, v...)
}

// Error is like log.Println, but messages are prefixed with "[WARN]"
func Warn(v ...any) {
	stdwarn.Println(v...)
}

// Warnf is like log.Printf, but messages are prefixed with "[WARN]"
func Warnf(format string, v ...any) {
	stdwarn.Printf(format, v...)
}
