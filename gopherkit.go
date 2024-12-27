package gopherkit

import (
	"fmt"
	"strconv"
)

// Parse integer values with a default value
func ParseInt(value string, defaultValue int) int {
	parsed, err := strconv.Atoi(value)

	if err != nil {
		return defaultValue
	}

	return parsed
}

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type UnsignedInt interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Float interface {
	~float32 | ~float64
}

type Complex interface {
	~complex64 | ~complex128
}

type Number interface {
	Integer | UnsignedInt | Float | Complex
}

// Calculate the sum of variadic or slice of numbers
func Sum[T Number](values ...any) T {
	var sum T

	for _, vv := range values {
		switch v := vv.(type) {
		case []T: // If the input is a slice
			for _, item := range v {
				sum += item
			}
		case T: // If the input is a single numeric value
			sum += v
		default:
			panic(fmt.Sprintf("Unsupported type: %T", v))
		}
	}

	return sum
}
