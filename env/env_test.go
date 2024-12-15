package env

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	// Create a temporary .env file
	tempFile, err := os.CreateTemp("", "test.env")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name()) // Clean up the temp file

	// Write sample environment variables
	content := `# Sample .env file
TEST_KEY=test_value
ANOTHER_KEY=another_value

# Commented out
# IGNORE_ME=ignored_value

EMPTY_LINE

INVALID_LINE
`
	_, err = tempFile.WriteString(content)
	if err != nil {
		t.Fatalf("failed to write to temp file: %v", err)
	}
	tempFile.Close()

	// Load environment variables from the file
	err = Load(tempFile.Name())
	if err != nil {
		t.Fatalf("Load() failed: %v", err)
	}

	// Check loaded variables
	if os.Getenv("TEST_KEY") != "test_value" {
		t.Errorf("expected TEST_KEY=test_value, got %s", os.Getenv("TEST_KEY"))
	}
	if os.Getenv("ANOTHER_KEY") != "another_value" {
		t.Errorf("expected ANOTHER_KEY=another_value, got %s", os.Getenv("ANOTHER_KEY"))
	}

	// Ensure invalid and commented lines are ignored
	if os.Getenv("IGNORE_ME") != "" {
		t.Errorf("expected IGNORE_ME to be unset, got %s", os.Getenv("IGNORE_ME"))
	}
}

func TestMustGet(t *testing.T) {
	// Set and unset environment variables for testing
	os.Setenv("MUSTGET_TEST", "mustget_value")
	defer os.Unsetenv("MUSTGET_TEST")

	if val := MustGet("MUSTGET_TEST"); val != "mustget_value" {
		t.Errorf("MustGet() returned %s, expected mustget_value", val)
	}

	// Test missing variable
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("MustGet() did not halt the program for missing variable")
		}
	}()
	MustGet("MISSING_KEY")
}

func TestGet(t *testing.T) {
	os.Setenv("GET_TEST", "get_value")
	defer os.Unsetenv("GET_TEST")

	if val := Get("GET_TEST", "fallback_value"); val != "get_value" {
		t.Errorf("Get() returned %s, expected get_value", val)
	}

	if val := Get("MISSING_GET", "fallback_value"); val != "fallback_value" {
		t.Errorf("Get() returned %s, expected fallback_value", val)
	}
}

func TestGetInt(t *testing.T) {
	os.Setenv("GET_INT_TEST", "42")
	defer os.Unsetenv("GET_INT_TEST")

	if val := GetInt("GET_INT_TEST", 99); val != 42 {
		t.Errorf("GetInt() returned %d, expected 42", val)
	}

	if val := GetInt("MISSING_GET_INT", 99); val != 99 {
		t.Errorf("GetInt() returned %d, expected 99", val)
	}

	// Test invalid integer
	os.Setenv("INVALID_INT", "notanumber")
	defer os.Unsetenv("INVALID_INT")
	if val := GetInt("INVALID_INT", 88); val != 88 {
		t.Errorf("GetInt() returned %d, expected 88", val)
	}
}
