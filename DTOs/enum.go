package dtos

import "fmt"

// unmarshalEnum decodes text into a string-kind enum field. An empty value is
// accepted as an unset no-op (callers treat "" as "not provided"); any
// non-empty value that fails valid is rejected so the binder returns an error.
func unmarshalEnum[T ~string](dst *T, text []byte, valid func(T) bool, name string) error {
	v := T(text)
	if v != "" && !valid(v) {
		return fmt.Errorf("invalid %s %q", name, v)
	}
	*dst = v
	return nil
}
