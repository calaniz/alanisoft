package handlers

import (
	"fmt"
	"log/syslog"
	"net/http"
)

type Error error

type AppError struct {
	Error
	Code int
	Message string
}

type AppHandler func(http.ResponseWriter, *http.Request) *AppError

func (h AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		logger.Err(err.Message)
		w.Header().Add("x-error", err.Message)
		http.Error(w, err.Message, err.Code)
	}
}

type Logger struct {
	Writer *syslog.Writer
}

func (l *Logger) Err(message string, args ...interface{}) {
	if len(args) > 0 {
		l.Writer.Err(fmt.Sprintf(message, args...))
	} else {
		l.Writer.Err(message)
	}
}

func (l *Logger) Notice(message string, args ...interface{}) {
	if len(args) > 0 {
		l.Writer.Notice(fmt.Sprintf(message, args...))
	} else {
		l.Writer.Notice(message)
	}
}

func (l *Logger) Debug(message string, args ...interface{}) {
	if len(args) > 0 {
		l.Writer.Debug(fmt.Sprintf(message, args...))
	} else {
		l.Writer.Debug(message)
	}
}

func (l *Logger) Info(message string, args ...interface{}) {
	if len(args) > 0 {
		l.Writer.Info(fmt.Sprintf(message, args...))
	} else {
		l.Writer.Info(message)
	}
}

var logger *Logger

func init() {
	writer, _ := syslog.New(syslog.LOG_NOTICE|syslog.LOG_LOCAL1, "alanisoft-handler")
	logger = &Logger{Writer: writer}
}