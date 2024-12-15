package assert

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

// AssertEqual asserts that two values are equal.
func Equal(t *testing.T, expected, actual interface{}, msg ...string) {
	if !reflect.DeepEqual(expected, actual) {
		message := formatMessage("expected", expected, actual, msg...)
		t.Error(message)
	}
}

// AssertNotEqual asserts that two values are not equal.
func NotEqual(t *testing.T, expected, actual interface{}, msg ...string) {
	if reflect.DeepEqual(expected, actual) {
		message := formatMessage("not expected", expected, actual, msg...)
		t.Error(message)
	}
}

// AssertNoError asserts that an error is nil.
func NoError(t *testing.T, err error, msg ...string) {
	if err != nil {
		message := formatMessage("no error", nil, err, msg...)
		t.Error(message)
	}
}

// AssertError asserts that an error is not nil.
func Error(t *testing.T, err error, msg ...string) {
	if err == nil {
		message := formatMessage("error", "non-nil error", err, msg...)
		t.Error(message)
	}
}

// AssertContains asserts that a string contains a substring.
func Contains(t *testing.T, s, substr string, msg ...string) {
	if !contains(s, substr) {
		message := formatMessage(fmt.Sprintf("'%s' to contain", substr), substr, s, msg...)
		t.Error(message)
	}
}

// AssertLen asserts that a collection has the expected length.
func Len(t *testing.T, collection interface{}, length int, msg ...string) {
	actualLen := reflect.ValueOf(collection).Len()
	if actualLen != length {
		message := formatMessage("length", length, actualLen, msg...)
		t.Error(message)
	}
}

// Helper function to check if a string contains a substring.
func contains(s, substr string) bool {
	return len(substr) == 0 || strings.Contains(s, substr)
}

// Helper function to format error messages.
func formatMessage(expectationType string, expected, actual interface{}, msg ...string) string {
	if len(msg) > 0 {
		return msg[0]
	}
	return fmt.Sprintf("Expected %v but got %v for %s", expected, actual, expectationType)
}
