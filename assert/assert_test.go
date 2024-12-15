package assert

import (
	"errors"
	"testing"
)

func TestEqual(t *testing.T) {
	tests := []struct {
		name     string
		expected interface{}
		actual   interface{}
		msg      []string
		wantErr  bool
	}{
		{
			name:     "Equal integers",
			expected: 42,
			actual:   42,
			wantErr:  false,
		},
		{
			name:     "Unequal integers",
			expected: 42,
			actual:   43,
			wantErr:  true,
		},
		{
			name:     "Equal strings",
			expected: "hello",
			actual:   "hello",
			wantErr:  false,
		},
		{
			name:     "Unequal strings",
			expected: "hello",
			actual:   "world",
			wantErr:  true,
		},
		{
			name:     "Equal slices",
			expected: []int{1, 2, 3},
			actual:   []int{1, 2, 3},
			wantErr:  false,
		},
		{
			name:     "Unequal slices",
			expected: []int{1, 2, 3},
			actual:   []int{3, 2, 1},
			wantErr:  true,
		},
		{
			name:     "Custom message",
			expected: 100,
			actual:   200,
			msg:      []string{"values do not match"},
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockT := &testing.T{}
			Equal(mockT, tt.expected, tt.actual, tt.msg...)
			if (mockT.Failed()) != tt.wantErr {
				t.Errorf("Equal() test failed for %s. WantErr = %v", tt.name, tt.wantErr)
			}
		})
	}
}

func TestNotEqual(t *testing.T) {
	tests := []struct {
		name     string
		expected interface{}
		actual   interface{}
		msg      []string
		wantErr  bool
	}{
		{"NotEqual integers", 42, 43, nil, false},
		{"Equal integers", 42, 42, nil, true},
		{"NotEqual strings", "hello", "world", nil, false},
		{"Equal strings", "hello", "hello", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockT := &testing.T{}
			NotEqual(mockT, tt.expected, tt.actual, tt.msg...)
			if (mockT.Failed()) != tt.wantErr {
				t.Errorf("NotEqual() test failed for %s. WantErr = %v", tt.name, tt.wantErr)
			}
		})
	}
}

func TestNoError(t *testing.T) {
	tests := []struct {
		name    string
		err     error
		msg     []string
		wantErr bool
	}{
		{"No error", nil, nil, false},
		{"Has error", errors.New("some error"), nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockT := &testing.T{}
			NoError(mockT, tt.err, tt.msg...)
			if (mockT.Failed()) != tt.wantErr {
				t.Errorf("NoError() test failed for %s. WantErr = %v", tt.name, tt.wantErr)
			}
		})
	}
}

func TestError(t *testing.T) {
	tests := []struct {
		name    string
		err     error
		msg     []string
		wantErr bool
	}{
		{"Has error", errors.New("some error"), nil, false},
		{"No error", nil, nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockT := &testing.T{}
			Error(mockT, tt.err, tt.msg...)
			if (mockT.Failed()) != tt.wantErr {
				t.Errorf("Error() test failed for %s. WantErr = %v", tt.name, tt.wantErr)
			}
		})
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		substr  string
		msg     []string
		wantErr bool
	}{
		{"Contains substring", "hello world", "world", nil, false},
		{"Does not contain substring", "hello world", "foo", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockT := &testing.T{}
			Contains(mockT, tt.s, tt.substr, tt.msg...)
			if (mockT.Failed()) != tt.wantErr {
				t.Errorf("Contains() test failed for %s. WantErr = %v", tt.name, tt.wantErr)
			}
		})
	}
}

func TestLen(t *testing.T) {
	tests := []struct {
		name       string
		collection interface{}
		length     int
		msg        []string
		wantErr    bool
	}{
		{"Correct length slice", []int{1, 2, 3}, 3, nil, false},
		{"Incorrect length slice", []int{1, 2, 3}, 5, nil, true},
		{"Correct length map", map[string]int{"a": 1, "b": 2}, 2, nil, false},
		{"Incorrect length map", map[string]int{"a": 1}, 2, nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockT := &testing.T{}
			Len(mockT, tt.collection, tt.length, tt.msg...)
			if (mockT.Failed()) != tt.wantErr {
				t.Errorf("Len() test failed for %s. WantErr = %v", tt.name, tt.wantErr)
			}
		})
	}
}
