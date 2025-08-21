package core

import (
	"reflect"

	"go.uber.org/zap/zapcore"

	"go.robertomontagna.dev/zapfluent/internal/functional/lazyoptional"
)

// Field is the interface that all concrete field types must implement. It
// represents a single key-value pair to be encoded.
type Field interface {
	// Name returns the key for the field.
	Name() string
	// Encode writes the field to the underlying zapcore.ObjectEncoder.
	// It returns an error if the encoding fails.
	Encode(zapcore.ObjectEncoder) error
}

// TypedField is a generic interface that extends the basic Field interface with
// methods for filtering and transformation. This allows for creating expressive,
// chainable APIs for constructing fields.
type TypedField[T any] interface {
	Field
	// Filter returns a new field that will only be encoded if the provided
	// condition returns true for its value.
	Filter(condition func(T) bool) TypedField[T]
	// NonZero is a convenience method that filters the field, ensuring it is
	// only encoded if its value is not the type's zero value.
	NonZero() TypedField[T]
	// Format returns a new string-based field by applying a formatting
	// function to the original field's value.
	Format(formatter func(T) string) TypedField[string]
}

type encodeFunc[T any] func(zapcore.ObjectEncoder, string, T) error

type typeFieldFunctions[T any] struct {
	encodeFunc encodeFunc[T]
	isNonZero  func(T) bool
}

type lazyTypedField[T any] struct {
	functions typeFieldFunctions[T]
	value     lazyoptional.LazyOptional[T]
	name      string
}

func newTypedField[T any](
	functions typeFieldFunctions[T],
	name string,
	initialValue T,
) TypedField[T] {
	return &lazyTypedField[T]{
		functions: functions,
		name:      name,
		value:     lazyoptional.Some(initialValue),
	}
}

func (f *lazyTypedField[T]) Name() string {
	return f.name
}

func (f *lazyTypedField[T]) Encode(encoder zapcore.ObjectEncoder) error {
	val, ok := f.value.Get()
	if !ok {
		return nil
	}
	return f.functions.encodeFunc(encoder, f.name, val)
}

func (f *lazyTypedField[T]) Filter(condition func(T) bool) TypedField[T] {
	return &lazyTypedField[T]{
		functions: f.functions,
		name:      f.name,
		value:     f.value.Filter(condition),
	}
}

func (f *lazyTypedField[T]) NonZero() TypedField[T] {
	return f.Filter(f.functions.isNonZero)
}

func (f *lazyTypedField[T]) Format(formatter func(T) string) TypedField[string] {
	return &lazyTypedField[string]{
		name:      f.name,
		functions: stringTypeFns(),
		value:     lazyoptional.Map(f.value, formatter),
	}
}

// ReflectiveIsNotNil checks if a value is non-nil using reflection.
//
// This function is useful as an `isNonZero` implementation for object fields
// where the zero-value check is simply a nil check. It correctly handles
// interfaces, pointers, channels, functions, maps, and slices. For other
// types, it returns true.
func ReflectiveIsNotNil[T any](v T) bool {
	if any(v) == nil {
		return false
	}

	val := reflect.ValueOf(v)
	switch val.Kind() {
	case reflect.Chan,
		reflect.Func,
		reflect.Map,
		reflect.Pointer,
		reflect.Slice,
		reflect.Interface:
		return !val.IsNil()
	default:
		return true
	}
}
