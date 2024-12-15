package json

import (
	"bytes"
	"encoding/json"
)

func Decode[T any](data []byte, target *T) error {
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	return decoder.Decode(target)
}

func Encode(v any) ([]byte, error) {
	return json.Marshal(v)
}
