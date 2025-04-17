package response_test

import (
	"bytes"
	"errors"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ferdiebergado/gopherkit/http/response"
)

func TestServerError(t *testing.T) {
	// Create a mock ResponseWriter
	rr := httptest.NewRecorder()

	// Create a mock error
	testErr := errors.New("test server error")

	// Capture the log output
	var buf bytes.Buffer
	h := slog.NewTextHandler(&buf, nil)
	oldHandler := slog.Default()
	slog.SetDefault(slog.New(h))
	defer slog.SetDefault(oldHandler) // Restore the original logger

	// Call the ServerError function
	response.ServerError(rr, testErr)

	// Check the status code
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusInternalServerError)
	}

	// Check the response body
	expected := "An error occurred.\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %q want %q",
			rr.Body.String(), expected)
	}

	// Check the log output
	logOutput := buf.String()
	if !strings.Contains(logOutput, "level=ERROR") {
		t.Errorf("log output should contain level=ERROR, got: %q", logOutput)
	}
	if !strings.Contains(logOutput, "msg=\"server error\"") {
		t.Errorf("log output should contain msg=\"server error\", got: %q", logOutput)
	}
	if !strings.Contains(logOutput, "reason=\"test server error\"") {
		t.Errorf("log output should contain reason=\"test server error\", got: %q", logOutput)
	}
	if !strings.Contains(logOutput, "stack_trace=") {
		t.Errorf("log output should contain stack_trace=, got: %q", logOutput)
	}
	if !strings.Contains(logOutput, "response.ServerError") {
		t.Errorf("log output should contain the function name, got: %q", logOutput)
	}
}
