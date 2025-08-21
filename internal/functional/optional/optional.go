// Package optional provides a generic, eager implementation of an Optional type.
package optional

// Optional is a container object which may or may not contain a non-nil value.
//
// This is an "eager" optional, meaning that transformations are applied
// immediately. For a "lazy" alternative, see the `lazyoptional` package.
type Optional[T any] struct {
	value    T
	hasValue bool
}

// Some returns an Optional that contains the given non-nil value.
func Some[T any](value T) Optional[T] {
	return Optional[T]{
		value:    value,
		hasValue: true,
	}
}

// Empty returns an empty Optional instance.
func Empty[T any]() Optional[T] {
	return Optional[T]{}
}

// OfPtr creates an Optional from a pointer.
// If the pointer is nil, an empty Optional is returned.
// Otherwise, an Optional containing the dereferenced value is returned.
func OfPtr[T any](ptr *T) Optional[T] {
	if ptr == nil {
		return Empty[T]()
	}
	return Some(*ptr)
}

// OfError creates an Optional from an error.
// If the error is nil, an empty Optional is returned.
// Otherwise, an Optional containing the error is returned.
func OfError(err error) Optional[error] {
	if err == nil {
		return Empty[error]()
	}
	return Some(err)
}

// Get returns the value if present, and `true`. Otherwise, it returns the
// zero value of the type and `false`.
func (o Optional[T]) Get() (T, bool) {
	if !o.hasValue {
		var zero T
		return zero, false
	}
	return o.value, true
}

// IsPresent returns `true` if there is a value present, otherwise `false`.
func (o Optional[T]) IsPresent() bool {
	return o.hasValue
}

// Map applies the given mapping function to the value if it is present.
//
// It returns an Optional describing the result of the mapping function.
// If the original optional is empty, this returns an empty optional.
func Map[T any, R any](o Optional[T], f func(T) R) Optional[R] {
	if o.IsPresent() {
		return Some(f(o.value))
	}
	return Empty[R]()
}

// FlatMap applies the given Optional-bearing mapping function to the value if
// it is present.
//
// This is useful for chaining operations that each return an Optional.
// If the original optional is empty, this returns an empty optional.
func FlatMap[T any, R any](o Optional[T], f func(T) Optional[R]) Optional[R] {
	if o.IsPresent() {
		return f(o.value)
	}
	return Empty[R]()
}
