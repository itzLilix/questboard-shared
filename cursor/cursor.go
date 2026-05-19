package cursor

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

// encodeCursor marshals any cursor value to base64-url JSON.
func EncodeCursor[T any](c T) (string, error) {
	b, err := json.Marshal(c)
	if err != nil {
		return "", fmt.Errorf("encode cursor: %w", err)
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

// decodeCursor unmarshals a base64-url JSON cursor into T. Empty input
// returns (nil, nil) — caller treats nil as "no cursor". Malformed input
// returns ErrInvalidCursor so the caller can bubble it as 400.
func DecodeCursor[T any](s string) (*T, error) {
	if s == "" {
		return nil, nil
	}
	b, err := base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		return nil, ErrInvalidCursor
	}
	var c T
	if err := json.Unmarshal(b, &c); err != nil {
		return nil, ErrInvalidCursor
	}
	return &c, nil
}