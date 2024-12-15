package env

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Loads environment variables from a file
func Load(envFile string) error {
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
		log.Printf("%s is not set, using %s as fallback.", envVar, fallback)
		return fallback
	}

	return value
}

// Retrieves an environment variable, returns a given fallback integer if not set
func GetInt(envVar string, fallback int) int {
	value, isSet := os.LookupEnv(envVar)
	parsed, err := strconv.Atoi(value)

	if !isSet || err != nil {
		log.Printf("%s is not set or invalid, using %d as fallback.", envVar, fallback)
		return fallback
	}

	return parsed
}
