package log

import (
	"log/slog"
	"os"
)

// Creates a structured logger
func CreateLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))
}

// Writes the error message to the stderr and halts the program.
func Fatal(reason string, err error) {
	slog.Error(
		"Fatal error occurred",
		"reason", reason,
		"details", err,
		"severity", "FATAL",
	)

	panic(reason)
}
