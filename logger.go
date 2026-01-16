package paystack

import (
	"context"
	"log/slog"
	"os"
)

// LogLevel represents the logging level
type LogLevel string

const (
	LogLevelDebug LogLevel = "debug"
	LogLevelInfo  LogLevel = "info"
	LogLevelWarn  LogLevel = "warn"
	LogLevelError LogLevel = "error"
)

// LogConfig holds logging configuration
type LogConfig struct {
	Level      LogLevel
	JSONOutput bool
	Output     *os.File
}

// Logger interface defines logging behavior
type Logger interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
	WithContext(ctx context.Context) Logger
}

// paystackLogger implements the Logger interface
type paystackLogger struct {
	logger *slog.Logger
	ctx    context.Context
}

// newLogger creates a new logger instance
func newLogger(config *LogConfig) Logger {
	if config == nil {
		config = &LogConfig{
			Level:      LogLevelInfo,
			JSONOutput: true,
			Output:     os.Stderr,
		}
	}

	var handler slog.Handler
	opts := &slog.HandlerOptions{
		Level: getLogLevel(config.Level),
	}

	if config.JSONOutput {
		handler = slog.NewJSONHandler(os.Stderr, opts)
	} else {
		handler = slog.NewTextHandler(config.Output, opts)
	}

	return &paystackLogger{
		logger: slog.New(handler),
		ctx:    context.Background(),
	}
}

func getLogLevel(level LogLevel) slog.Level {
	switch level {
	case LogLevelDebug:
		return slog.LevelDebug
	case LogLevelWarn:
		return slog.LevelWarn
	case LogLevelError:
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

func (l *paystackLogger) Debug(msg string, args ...any) {
	l.logger.DebugContext(l.ctx, msg, args...)
}

func (l *paystackLogger) Info(msg string, args ...any) {
	l.logger.InfoContext(l.ctx, msg, args...)
}

func (l *paystackLogger) Warn(msg string, args ...any) {
	l.logger.WarnContext(l.ctx, msg, args...)
}

func (l *paystackLogger) Error(msg string, args ...any) {
	l.logger.ErrorContext(l.ctx, msg, args...)
}

func (l *paystackLogger) WithContext(ctx context.Context) Logger {
	return &paystackLogger{
		logger: l.logger,
		ctx:    ctx,
	}
}
