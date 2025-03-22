package response_test

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ferdiebergado/gopherkit/assert"
	"github.com/ferdiebergado/gopherkit/http/response"
)

func TestServerError(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()

	response.ServerError(rr, req, fmt.Errorf("some function call: %w", errors.New("failed")))

	res := rr.Result()
	defer res.Body.Close()

	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
	assert.Contains(t, rr.Body.String(), "An error occurred.")
}
