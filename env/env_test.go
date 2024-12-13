package env

import (
	"os"
	"testing"
)

type MockLogger struct {
	FatalCalled  bool
	FatalfCalled bool
	FatalArgs    []interface{}
}

func (m *MockLogger) Fatal(args ...interface{}) {
	m.FatalCalled = true
	m.FatalArgs = args
}

func (m *MockLogger) Fatalf(format string, args ...interface{}) {
	m.FatalfCalled = true
	m.FatalArgs = args
}

func (m *MockLogger) Write(p []byte) (n int, err error) {
	// Implement your desired behavior for writing to the mock logger
	// For example, you could store the written bytes in a buffer
	// or log them to a testing log.
	return len(p), nil
}

func TestMustGet(t *testing.T) {
	t.Run("Should return the value if the environment variable is set", func(t *testing.T) {
		const envVar = "ENV"
		const dev = "development"

		if err := os.Setenv(envVar, dev); err != nil {
			t.Errorf("failed to set env vars: %v", err)
		}

		want := dev

		if got := MustGet(envVar); got != want {
			t.Errorf("MustGet() = %v, want %v", got, want)
		}

	})
}

func TestGet(t *testing.T) {
	t.Run("Should return the value if the environment variable is set", func(t *testing.T) {
		const envVar = "ENV"
		const dev = "development"

		if err := os.Setenv(envVar, dev); err != nil {
			t.Errorf("failed to set env vars: %v", err)
		}

		want := dev

		if got := Get(envVar, dev); got != want {
			t.Errorf("MustGet() = %v, want %v", got, want)
		}
	})

	t.Run("Should return the fallback if the environment variable is not set", func(t *testing.T) {
		const envVar = "PORT"
		const fallback = "8000"

		want := fallback

		if got := Get(envVar, fallback); got != want {
			t.Errorf("MustGet() = %v, want %v", got, want)
		}
	})
}
