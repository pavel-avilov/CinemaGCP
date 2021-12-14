package logger

import (
	"fmt"
	"log"
)

type StdHandler struct {
}

// NewStdHandler creates StdHandler
func NewStdHandler() *StdHandler {
	return &StdHandler{}
}

func (h StdHandler) Info(args ...interface{}) {
	h.logMessage(INFO, args...)
}

func (h StdHandler) Infof(format string, args ...interface{}) {
	h.logMessage(INFO, fmt.Sprintf(format, args...))
}

func (h StdHandler) Warn(args ...interface{}) {
	h.logMessage(WARN, args...)
}

func (h StdHandler) Warnf(format string, args ...interface{}) {
	h.logMessage(WARN, fmt.Sprintf(format, args...))
}

func (h StdHandler) Error(args ...interface{}) {
	h.logMessage(ERROR, args...)
}

func (h StdHandler) Errorf(format string, args ...interface{}) {
	h.logMessage(ERROR, fmt.Sprintf(format, args...))
}

func (h StdHandler) Debug(args ...interface{}) {
	h.logMessage(DEBUG, args...)
}

func (h StdHandler) Debugf(format string, args ...interface{}) {
	h.logMessage(DEBUG, fmt.Sprintf(format, args...))
}

func (h StdHandler) Fatal(args ...interface{}) {
	args = append([]interface{}{FATAL}, args...)
	log.Fatalln(args...)
}

func (h StdHandler) Fatalf(format string, args ...interface{}) {
	args = append([]interface{}{FATAL}, fmt.Sprintf(format, args...))
	log.Fatalln(args...)
}

// logMessage logs a message at level severity.
func (h StdHandler) logMessage(severity Severity, args ...interface{}) {
	args = append([]interface{}{severity}, args...)
	log.Println(args...)
}
