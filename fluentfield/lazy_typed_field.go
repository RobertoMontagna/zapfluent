package fluentfield

import (
	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent/functional/lazyoptional"
)

// EncodeFunc is a generic function type for encoding a field of a specific type.
type EncodeFunc[T any] func(zapcore.ObjectEncoder, string, T) error

// TypeFieldFunctions holds the core functions for handling a specific type in
// a TypedField. It includes how to encode the type and how to check if a value
// of the type is "non-zero" (and thus should be included by default).
type TypeFieldFunctions[T any] struct {
	EncodeFunc EncodeFunc[T]
	IsNonZero  func(T) bool
}

// LazyTypedField is the default implementation of the TypedField interface.
// It uses a LazyOptional to defer transformations and filtering, which avoids
// unnecessary allocations and computations if a field is ultimately filtered out.
type LazyTypedField[T any] struct {
	functions TypeFieldFunctions[T]
	optional  lazyoptional.LazyOptional[T]
	name      string
}

// NewTypedField creates a new TypedField with the given functions, name, and
// initial value. This is the primary constructor for creating concrete field types.
func NewTypedField[T any](
	functions TypeFieldFunctions[T],
	name string,
	value T,
) TypedField[T] {
	return &LazyTypedField[T]{
		functions: functions,
		name:      name,
		optional:  lazyoptional.Some(value),
	}
}

// Name returns the key for the field.
func (f *LazyTypedField[T]) Name() string {
	return f.name
}

// Encode writes the field to the underlying zapcore.ObjectEncoder.
// If the internal LazyOptional is empty (e.g., due to filtering), this
// method does nothing.
func (f *LazyTypedField[T]) Encode(encoder zapcore.ObjectEncoder) error {
	val, ok := f.optional.Get()
	if !ok {
		return nil
	}
	return f.functions.EncodeFunc(encoder, f.name, val)
}

// Filter returns a new field that will only be encoded if the provided
// condition returns true for its value.
func (f *LazyTypedField[T]) Filter(condition func(T) bool) TypedField[T] {
	return &LazyTypedField[T]{
		functions: f.functions,
		name:      f.name,
		optional:  f.optional.Filter(condition),
	}
}

// NonZero is a convenience method that filters the field, ensuring it is only
// encoded if its value is not the type's zero value.
func (f *LazyTypedField[T]) NonZero() TypedField[T] {
	return f.Filter(f.functions.IsNonZero)
}

// Format returns a new string-based field by applying a formatting function to
// the original field's value.
func (f *LazyTypedField[T]) Format(formatter func(T) string) TypedField[string] {
	return &LazyTypedField[string]{
		name:      f.name,
		functions: stringTypeFns(),
		optional:  lazyoptional.Map(f.optional, formatter),
	}
}
