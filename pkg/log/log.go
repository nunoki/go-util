package log

import (
	stdLog "log"
	"os"
)

var stdout = stdLog.New(os.Stdout, "", stdLog.LstdFlags)
var stderr = stdLog.New(os.Stderr, "[ERROR] ", stdLog.LstdFlags)
var stdwarn = stdLog.New(os.Stderr, "[WARN] ", stdLog.LstdFlags)

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
