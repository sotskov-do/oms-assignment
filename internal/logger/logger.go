package logger

import (
	"log/slog"
	"os"

	"github.com/sotskov-do/oms-assignment/internal/config"
)

const (
	LevelCritical = slog.Level(12)
)

var LevelNames = map[slog.Leveler]string{
	LevelCritical: "CRITICAL",
}

func New(addSource bool) *slog.Logger {
	logLevel := getLogLevel()

	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: addSource,
		Level:     logLevel,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.LevelKey {
				level := a.Value.Any().(slog.Level)
				levelLabel, exists := LevelNames[level]
				if !exists {
					levelLabel = level.String()
				}
				a.Value = slog.StringValue(levelLabel)
			}
			return a
		},
	}))
}

func getLogLevel() slog.Level {
	level := os.Getenv(config.LogLevel)
	switch level {
	case "DEBUG":
		return slog.LevelDebug
	case "INFO":
		return slog.LevelInfo
	case "WARN":
		return slog.LevelWarn
	case "ERROR":
		return slog.LevelError
	case "CRITICAL":
		return LevelCritical
	default:
		return slog.LevelDebug
	}
}
