// Package enum provides helpers for creating and working with enum-like types in
// Go.
package enum

import "fmt"

// Value is a constraint that permits any signed integer type to be used as the
// underlying type for an enum.
type Value interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Enum provides a generic way to handle enum-like types, mapping enum values
// to their string representations.
type Enum[E Value] struct {
	values  map[E]string
	unknown E
}

// New creates a new enum helper.
//
// It takes a map of enum values to their string names, and a designated
// "unknown" value to be used as a fallback.
func New[E Value](values map[E]string, unknown E) Enum[E] {
	return Enum[E]{
		values:  values,
		unknown: unknown,
	}
}

// String returns the string representation of an enum value.
//
// If the value is not found in the map, it returns a formatted "Unknown" string.
func (e Enum[E]) String(v E) string {
	if s, ok := e.values[v]; ok {
		return s
	}
	return fmt.Sprintf("Unknown(%d)", v)
}

// FromInt converts an integer into an enum value.
//
// If the integer does not correspond to a known enum value, it returns the
// configured "unknown" value.
func (e Enum[E]) FromInt(i int) E {
	k := E(i)
	if _, ok := e.values[k]; ok {
		return k
	}
	return e.unknown
}
