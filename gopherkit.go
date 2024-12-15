package gopherkit

import "strconv"

// Parse integer values with a default value
func ParseInt(value string, defaultValue int) int {
	parsed, err := strconv.Atoi(value)

	if err != nil {
		return defaultValue
	}

	return parsed
}
