package logger

import (
	"cloud.google.com/go/logging"
	"context"
	"log"
	"os"
)

type Severity string

const (
	INFO               Severity = "[INFO]:"
	WARN               Severity = "[WARN]:"
	ERROR              Severity = "[ERROR]:"
	FATAL              Severity = "[FATAL]:"
	DEBUG              Severity = "[DEBUG]:"
	googleProjectIdKey          = "GOOGLE_CLOUD_PROJECT"
)

var handlers []Handler

func SetupLogger(ctx context.Context) {
	if value, present := os.LookupEnv(googleProjectIdKey); present {
		client, err := logging.NewClient(ctx, value)
		if err != nil {
			log.Fatalf("Failed to create client: %v", err)
		}
		cloudLogger := NewCloudLoggingHandler(client)
		AddHandler(cloudLogger)
	} else {
		stdLogger := NewStdHandler()
		AddHandler(stdLogger)
	}
}

// SetHandlers set a new array of logger handlers
func SetHandlers(h []Handler) {
	handlers = h
}

// AddHandler adds a new handler to the array
func AddHandler(h Handler) {
	handlers = append(handlers, h)
}

func Info(args ...interface{}) {
	for _, handler := range handlers {
		handler.Info(args...)
	}
}

func Infof(format string, args ...interface{}) {
	for _, handler := range handlers {
		handler.Infof(format, args...)
	}
}

func Warn(args ...interface{}) {
	for _, handler := range handlers {
		handler.Warn(args...)
	}
}

func Warnf(format string, args ...interface{}) {
	for _, handler := range handlers {
		handler.Warnf(format, args...)
	}
}

func Error(args ...interface{}) {
	for _, handler := range handlers {
		handler.Error(args...)
	}
}

func Errorf(format string, args ...interface{}) {
	for _, handler := range handlers {
		handler.Errorf(format, args...)
	}
}

func Debug(args ...interface{}) {
	for _, handler := range handlers {
		handler.Debug(args...)
	}
}

func Debugf(format string, args ...interface{}) {
	for _, handler := range handlers {
		handler.Debugf(format, args...)
	}
}

func Fatal(args ...interface{}) {
	for _, handler := range handlers {
		handler.Fatal(args...)
	}
}

func Fatalf(format string, args ...interface{}) {
	for _, handler := range handlers {
		handler.Fatalf(format, args...)
	}
}
