package env

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

// Loads environment variables from a file
func Load(envFile string) error {
	slog.Info("Loading environment file", "file", envFile)
	file, err := os.Open(envFile)

	if err != nil {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		// Skip empty lines and comments
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		}

		// Split the line into key and value
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Set the environment variable
		if err := os.Setenv(key, value); err != nil {
			return fmt.Errorf("os setenv: %v", err)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scanner: %v", err)
	}

	return nil
}

// Stops program execution when an environment variable is not set
func MustGet(envVar string) string {
	value, isSet := os.LookupEnv(envVar)

	if !isSet {
		panic(envVar + " environment variable is not set!\n")
	}

	return value
}

// Retrieves an environment variable, returns a given fallback if not set
func Get(envVar string, fallback string) string {
	value, isSet := os.LookupEnv(envVar)

	if !isSet {
		slog.Warn("Environment variable is not set, using fallback.", "variable", envVar, "fallback", fallback)
		return fallback
	}

	slog.Info("Environment variable is set", "variable", envVar, "value", value)

	return value
}

// Retrieves an environment variable as an int, returns a given fallback if not set
func GetInt(envVar string, fallback int) int {
	value, isSet := os.LookupEnv(envVar)
	parsed, err := strconv.Atoi(value)

	if !isSet || err != nil {
		slog.Warn("Environment variable is not set or invalid, using fallback.", "variable", envVar, slog.Int("fallback", fallback))
		return fallback
	}

	slog.Info("Environment variable is set", "variable", envVar, slog.Int("value", parsed))

	return parsed
}

// Retrieves an environment variable as a bool, returns a given fallback if not set
func GetBool(envVar string, fallback bool) bool {
	value, isSet := os.LookupEnv(envVar)
	parsed, err := strconv.ParseBool(value)

	if !isSet || err != nil {
		slog.Warn("Environment variable is not set or invalid, using fallback.", "variable", envVar, slog.Bool("fallback", fallback))
		return fallback
	}

	slog.Info("Environment variable is set", "variable", envVar, slog.Bool("value", parsed))

	return parsed
}
