package core

import (
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

// TypedPointerField is a generic interface that represents a field with a
// pointer value. It implements the base Field interface and provides a `NonNil`
// method to safely convert it to a TypedField for chaining.
type TypedPointerField[T any] interface {
	Field
	// NonNil converts the pointer field into a standard TypedField. If the
	// original pointer was nil, the resulting TypedField will be empty,
	// preventing further operations in a chain.
	NonNil() TypedField[T]
	// WithAddress returns a new field that, when encoded, produces an object
	// containing both the pointer's value and its memory address. This is
	// useful for debugging and understanding pointer references in logs.
	WithAddress() TypedField[PointerInfo[T]]
}

type encodeFunc[T any] func(zapcore.ObjectEncoder, string, T) error

type typeFieldFunctions[T any] struct {
	encodeFunc encodeFunc[T]
	isNonZero  func(T) bool
}

type typePointerFieldFunctions[T any] struct {
	typeFieldFunctions[T]
	toField func(name string, value T) TypedField[T]
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
		functions: stringTypeFns,
		value:     lazyoptional.Map(f.value, formatter),
	}
}

type pointerField[T any] struct {
	functions typePointerFieldFunctions[T]
	value     *T
	name      string
}

func newPointerField[T any](
	functions typePointerFieldFunctions[T],
	name string,
	value *T,
) TypedPointerField[T] {
	return &pointerField[T]{
		functions: functions,
		name:      name,
		value:     value,
	}
}

func (p *pointerField[T]) Name() string {
	return p.name
}

func (p *pointerField[T]) Encode(encoder zapcore.ObjectEncoder) error {
	if p.value != nil {
		return p.functions.encodeFunc(encoder, p.name, *p.value)
	}
	encoder.AddString(p.name, NilSentinel)
	return nil
}

func (p *pointerField[T]) NonNil() TypedField[T] {
	return &lazyTypedField[T]{
		functions: p.functions.typeFieldFunctions,
		name:      p.name,
		value:     lazyoptional.OfPtr(p.value),
	}
}

func (p *pointerField[T]) WithAddress() TypedField[PointerInfo[T]] {
	return Object(
		p.name,
		PointerInfo[T]{
			PtrValue:  p.value,
			functions: p.functions,
		},
		PointerInfo[T].isNonZero,
	)
}
