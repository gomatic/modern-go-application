// Package log provides logger configuration and creation for the application.
package log

import (
	"log/slog"
	"os"

	"github.com/urfave/cli/v2"
)

const LoggerMetadataKey = "logger"

// Level represents the logging level as a string.
type Level string

// Format represents the log output format.
type Format string

// Common log format constants.
const (
	TextFormat Format = "text"
	JSONFormat Format = "json"
)

// Config holds the logger configuration
type Config struct {
	Level  Level
	Format Format
}

// GetLoggerFunc is a function type for getting a logger
type GetLoggerFunc func(*cli.Context, Config) *slog.Logger

// GetLogger creates and configures a logger based on the provided configuration
func GetLogger(c *cli.Context, cfg Config) *slog.Logger {
	// Get logger from metadata if it already exists
	if logger, ok := c.App.Metadata[LoggerMetadataKey].(*slog.Logger); ok {
		return logger
	}

	// Parse log level from string
	level := slog.LevelInfo // default
	_ = level.UnmarshalText([]byte(cfg.Level))

	// Create is the handler based on format
	var handler slog.Handler
	opts := &slog.HandlerOptions{
		Level: level,
	}

	switch cfg.Format {
	case JSONFormat:
		handler = slog.NewJSONHandler(os.Stdout, opts)
	case TextFormat:
		fallthrough
	default:
		handler = slog.NewTextHandler(os.Stdout, opts)
	}

	return slog.New(handler)
}
