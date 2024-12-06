package request

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Decodes the JSON payload from the request body
func JSON[T any](r *http.Request) (T, error) {
	var v T
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, fmt.Errorf("decode json: %w", err)
	}
	return v, nil
}
