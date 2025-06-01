package logutil

import (
	"log/slog"
	"os"
)

// SetupLogger initializes the global logger with the given verbosity setting.
// If verbose is true, the logger outputs debug-level logs and above.
// Otherwise, it logs only info-level and higher messages.
// Logs are written in JSON format to stderr to avoid mixing with program output,
// which is typically written to stdout (e.g., merged code).
func SetupLogger(verbose bool) {
	var level slog.Level
	if verbose {
		level = slog.LevelDebug
	} else {
		level = slog.LevelInfo
	}
	logger := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: level}))
	slog.SetDefault(logger)
}
