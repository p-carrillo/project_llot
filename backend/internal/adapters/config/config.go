package config

import (
	"log/slog"
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	HTTPAddr       string
	LogLevel       slog.Level
	MaxBodyBytes   int64
	SessionTimeout time.Duration
	DataFilePath   string
}

func FromEnv() Config {
	return Config{
		HTTPAddr:       getEnv("AGENT_HTTP_ADDR", ":8080"),
		LogLevel:       parseLogLevel(getEnv("AGENT_LOG_LEVEL", "info")),
		MaxBodyBytes:   parseInt64(getEnv("AGENT_MAX_BODY_BYTES", "1048576"), 1048576),
		SessionTimeout: parseDuration(getEnv("AGENT_SESSION_TIMEOUT", "30m"), 30*time.Minute),
		DataFilePath:   getEnv("AGENT_DATA_FILE", "./data/events.jsonl"),
	}
}

func parseLogLevel(value string) slog.Level {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "debug":
		return slog.LevelDebug
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && strings.TrimSpace(value) != "" {
		return value
	}
	return fallback
}

func parseInt64(value string, fallback int64) int64 {
	parsed, err := strconv.ParseInt(strings.TrimSpace(value), 10, 64)
	if err != nil || parsed <= 0 {
		return fallback
	}
	return parsed
}

func parseDuration(value string, fallback time.Duration) time.Duration {
	parsed, err := time.ParseDuration(strings.TrimSpace(value))
	if err != nil || parsed <= 0 {
		return fallback
	}
	return parsed
}
