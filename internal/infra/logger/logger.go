package logger

import (
	"log/slog"
	"os"
)

var globalLogger *slog.Logger

func NewLogger() *slog.Logger {
	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}
	globalLogger = slog.New(slog.NewTextHandler(os.Stdout, opts))
	return globalLogger
}

func Info(msg string, args ...any) {
	globalLogger.Info(msg, args...)
}

func Warn(msg string, args ...any) {
	globalLogger.Warn(msg, args...)
}

func Error(msg string, args ...any) {
	globalLogger.Error(msg, args...)
}

func Debug(msg string, args ...any) {
	globalLogger.Debug(msg, args...)
}
