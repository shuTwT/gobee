package logger

import (
	"log"
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
	args = append([]any{"pid", os.Getpid()}, args...)
	globalLogger.Info(msg, args...)
}

func Warn(msg string, args ...any) {
	args = append([]any{"pid", os.Getpid()}, args...)
	globalLogger.Warn(msg, args...)
}

func Error(msg string, args ...any) {
	args = append([]any{"pid", os.Getpid()}, args...)
	globalLogger.Error(msg, args...)
}

func Debug(msg string, args ...any) {
	globalLogger.Debug(msg, args...)
}

func Panic(msg string, args ...any) {
	args = append([]any{"pid", os.Getpid()}, args...)
	log.Panicf(msg, args...)
}
