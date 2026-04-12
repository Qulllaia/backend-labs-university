package logger

import (
	"log/slog"
	"os"
)

var (
	fileLog    *os.File
	fileLogger *slog.Logger
)

func InitLogger(name string) {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})))

	setFileLogger(name)
}

func setFileLogger(name string) {
	var err error
	fileLog, err = os.Create(name)
	if err != nil {
		slog.Error("Failed to create log file", "error", err)
		return
	}

	fileLogger = slog.New(slog.NewJSONHandler(fileLog, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
}

func GetFileLogger() *slog.Logger {
	return fileLogger
}
