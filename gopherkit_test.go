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
			result := ParseInt(tt.input, tt.defaultValue)
			if result != tt.expected {
				t.Errorf("ParseInt(%q, %d) = %d; expected %d", tt.input, tt.defaultValue, result, tt.expected)
			}
		})
	}
}
