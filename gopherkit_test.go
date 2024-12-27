package gopherkit

import (
	"testing"

	"github.com/ferdiebergado/gopherkit/assert"
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

type SumTestCase[T Number] struct {
	name     string
	input    []T
	expected T
}

func TestSum(t *testing.T) {
	tests := []SumTestCase[int]{
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
			assert.Equal(t, tt.expected, result)
		})
	}

	float64test := []SumTestCase[float64]{
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
			assert.Equal(t, tt.expected, result)
		})
	}

	t.Run("Sum of variadic ints", func(t *testing.T) {
		t.Parallel()

		result := Sum[int](1, -2, 3)
		expected := 2
		assert.Equal(t, expected, result)
	})

	t.Run("Sum of variadic float64", func(t *testing.T) {
		t.Parallel()

		result := Sum[float64](1.25, 2.5, 3.01)
		expected := 6.76
		assert.Equal(t, expected, result)
	})

	t.Run("Sum of variadic slices", func(t *testing.T) {
		t.Parallel()

		result := Sum[int]([]int{1, 1, 1}, []int{2, 2, 2}, []int{3, 3, 3})
		expected := 18
		assert.Equal(t, expected, result)
	})
}
