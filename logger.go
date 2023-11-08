package paystack

import (
	"log/slog"
)

type Logger struct {
	log *slog.Logger
}

func (l Logger) Debug(msg string, args ...any) {
	l.log.Debug(msg, args...)
}


func (l Logger) Info(msg string, args ...any) {
	l.log.Info(msg, args...)
}

func (l Logger) Warn(msg string, args ...any) {
	l.log.Warn(msg, args...)
}

func (l Logger) Error(msg string, args ...any) {
	l.log.Error(msg, args...)
}

func (l Logger) With(args ...any) *Logger {
	return &Logger{log: l.log.With(args...)}
}

func (l Logger) WithGroup(group string) *Logger {
	return &Logger{log: l.log.WithGroup(group)}
}

func NewLogger(l *slog.Logger) *Logger {
	return &Logger{log: l}
}
