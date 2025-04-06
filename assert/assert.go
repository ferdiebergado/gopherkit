package assert

import (
	"reflect"
	"strings"
	"testing"
)

// AssertEqual asserts that two values are equal.
func Equal(t *testing.T, expected, actual any) {
	t.Helper()
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but got %v for equal", expected, actual)
	}
}

// AssertNotEqual asserts that two values are not equal.
func NotEqual(t *testing.T, expected, actual any) {
	t.Helper()
	if reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but got %v for not equal", expected, actual)
	}
}

// AssertNoError asserts that an error is nil.
func NoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("Expected no error but got %v", err)
	}
}

// AssertError asserts that an error is not nil.
func Error(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Errorf("Expected error but got nil")
	}
}

// AssertContains asserts that a string contains a substring.
func Contains(t *testing.T, s, substr string) {
	t.Helper()
	if len(substr) > 0 && !strings.Contains(s, substr) {
		t.Errorf("Expected %s to contain %s", s, substr)
	}
}

// AssertLen asserts that a collection has the expected length.
func Len(t *testing.T, collection any, length int) {
	t.Helper()
	actualLen := reflect.ValueOf(collection).Len()
	if actualLen != length {
		t.Errorf("Expected length %d but got %d", length, actualLen)
	}
}
