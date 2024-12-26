package gopherkit

import (
	"testing"
)

func TestParseInt(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		defaultValue int
		expected     int
	}{
		{
			name:         "Valid Integer",
			input:        "42",
			defaultValue: 0,
			expected:     42,
		},
		{
			name:         "Invalid Integer - Non-numeric",
			input:        "abc",
			defaultValue: 10,
			expected:     10,
		},
		{
			name:         "Invalid Integer - Empty String",
			input:        "",
			defaultValue: 5,
			expected:     5,
		},
		{
			name:         "Valid Integer - Negative Value",
			input:        "-7",
			defaultValue: 0,
			expected:     -7,
		},
		{
			name:         "Valid Integer - Zero",
			input:        "0",
			defaultValue: 100,
			expected:     0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := ParseInt(tt.input, tt.defaultValue)
			if result != tt.expected {
				t.Errorf("ParseInt(%q, %d) = %d; expected %d", tt.input, tt.defaultValue, result, tt.expected)
			}
		})
	}
}

type SumTestData[T Number] struct {
	name     string
	input    []T
	expected T
}

func TestSum(t *testing.T) {
	tests := []SumTestData[int]{
		{
			name:     "Sum of empty slice",
			input:    []int{},
			expected: 0,
		},
		{
			name:     "Sum of slice of ints",
			input:    []int{1, 2, 3},
			expected: 6,
		},
		{
			name:     "Sum of slices with negative ints",
			input:    []int{5, -2, 5},
			expected: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := Sum[int](tt.input)
			if result != tt.expected {
				t.Errorf("Sum(%v)= %d; expected %d", tt.input, result, tt.expected)
			}
		})
	}

	float64test := []SumTestData[float64]{
		{
			name:     "Sum of empty float64 slice",
			input:    []float64{},
			expected: 0,
		},
		{
			name:     "Sum of slice of float64",
			input:    []float64{1.25, 2.5, 3.01},
			expected: 6.76,
		},
	}

	for _, tt := range float64test {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result := Sum[float64](tt.input)
			if result != tt.expected {
				t.Errorf("Sum(%v)= %f; expected %f", tt.input, result, tt.expected)
			}
		})
	}

	t.Run("Sum of variadic ints", func(t *testing.T) {
		t.Parallel()

		result := Sum[int](1, 2, 3)
		expected := 6
		if result != expected {
			t.Errorf("Sum(%v)= %d; expected %d", "1,2,3", result, expected)
		}
	})

	t.Run("Sum of variadic float64", func(t *testing.T) {
		t.Parallel()

		result := Sum[float64](1.25, 2.5, 3.01)
		expected := 6.76
		if result != expected {
			t.Errorf("Sum(%v)= %f; expected %f", "1.25, 2.5, 3.01", result, expected)
		}
	})
}
