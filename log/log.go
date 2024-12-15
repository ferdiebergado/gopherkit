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
		slog.String("reason", reason),
		slog.String("error", err.Error()),
		slog.String("severity", "FATAL"),
	)

	panic(reason)
}
