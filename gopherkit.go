package gopherkit

import "strconv"

// Parse integer values with a default value
func ParseInt(value string, defaultValue int) int {
	i, err := strconv.Atoi(value)

	if err != nil {
		return defaultValue
	}

	return i
}
