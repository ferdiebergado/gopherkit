package env

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// LoadEnv loads environment variables from a file
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

// Halts the program if an environment variable is unset
func MustGet(envVar string) string {
	value, isSet := os.LookupEnv(envVar)

	if !isSet {
		log.Fatalf("%s environment variable is not set!\n", envVar)
	}

	return value
}

// Retrieves an environment variable and uses a given fallback if unset
func Get(envVar string, fallback string) string {
	value, isSet := os.LookupEnv(envVar)

	if !isSet {
		log.Println(envVar, " is not set, using fallback of", fallback, ".")
		return fallback
	}

	return value
}
