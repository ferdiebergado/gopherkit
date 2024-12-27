package json

import (
	"bytes"
	"encoding/json"
)

// Decode json data to a struct
func Decode[T any](data []byte, target *T) error {
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	return decoder.Decode(target)
}

// Encode data into json
func Encode(v any) ([]byte, error) {
	return json.Marshal(v)
}
