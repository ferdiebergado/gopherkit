package log

import (
	"fmt"
	"log/slog"
	"os"
)

// Creates a structured logger
func CreateLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))
}

// Writes the error message to the stderr and halts the program.
func Fatal(message string, err error) {
	fmt.Fprintf(os.Stderr, "%s: %v\n", message, err)
	os.Exit(1)
}
