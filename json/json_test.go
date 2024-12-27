package json

import "testing"

// Define a struct for testing
type TestStruct struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestDecode(t *testing.T) {
	tests := []struct {
		name      string
		data      []byte
		target    *TestStruct
		expectErr bool
	}{
		{
			name:      "Valid JSON",
			data:      []byte(`{"name": "Alice", "age": 30}`),
			target:    &TestStruct{},
			expectErr: false,
		},
		{
			name:      "Unknown Field",
			data:      []byte(`{"name": "Alice", "age": 30, "unknown": "value"}`),
			target:    &TestStruct{},
			expectErr: true,
		},
		{
			name:      "Invalid JSON",
			data:      []byte(`{"name": "Alice", "age":`), // Invalid JSON
			target:    &TestStruct{},
			expectErr: true,
		},
		{
			name:      "Nil Target",
			data:      []byte(`{"name": "Alice", "age": 30}`),
			target:    nil,
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Decode(tt.data, tt.target)

			if (err != nil) != tt.expectErr {
				t.Errorf("expected error: %v, got: %v", tt.expectErr, err)
				return
			}

			if !tt.expectErr && (tt.target.Name != "Alice" || tt.target.Age != 30) {
				t.Errorf("expected target to be populated with Name: Alice and Age: 30, got Name: %s and Age: %d", tt.target.Name, tt.target.Age)
			}
		})
	}
}

func TestEncode(t *testing.T) {
	tests := []struct {
		name      string
		input     any
		expectErr bool
		expected  string
	}{
		{
			name: "Valid Struct",
			input: struct {
				Name string
				Age  int
			}{"Alice", 30},
			expectErr: false,
			expected:  `{"Name":"Alice","Age":30}`,
		},
		{
			name:      "Empty Struct",
			input:     struct{}{},
			expectErr: false,
			expected:  `{}`,
		},
		{
			name:      "Slice of Structs",
			input:     []struct{ Name string }{{"Alice"}, {"Bob"}},
			expectErr: false,
			expected:  `[{"Name":"Alice"},{"Name":"Bob"}]`,
		},
		{
			name:      "Unsupported Type",
			input:     make(chan int), // Channels cannot be marshaled to JSON
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Encode(tt.input)

			if (err != nil) != tt.expectErr {
				t.Errorf("expected error: %v, got: %v", tt.expectErr, err)
				return
			}

			if !tt.expectErr && string(result) != tt.expected {
				t.Errorf("expected result: %s, got: %s", tt.expected, result)
			}
		})
	}
}
