package optional

// Optional is a container object which may or may not contain a non-nil value.
type Optional[T any] struct {
	value    T
	hasValue bool
}

// Some returns an Optional with the specified present non-nil value.
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

// Get returns the value if present, otherwise returns the zero value of the type and false.
func (o Optional[T]) Get() (T, bool) {
	if !o.hasValue {
		var zero T
		return zero, false
	}
	return o.value, true
}

// IsPresent returns true if there is a value present, otherwise false.
func (o Optional[T]) IsPresent() bool {
	return o.hasValue
}

// ForEach performs the given action with the value if a value is present.
func (o Optional[T]) ForEach(f func(T)) {
	if o.IsPresent() {
		f(o.value)
	}
}

// Map applies the given mapping function to the value if a value is present,
// and returns an Optional describing the result.
func Map[T any, R any](o Optional[T], f func(T) R) Optional[R] {
	if o.IsPresent() {
		return Some(f(o.value))
	}
	return Empty[R]()
}

// FlatMap applies the given Optional-bearing mapping function to the value if a value is present,
// and returns the result.
func FlatMap[T any, R any](o Optional[T], f func(T) Optional[R]) Optional[R] {
	if o.IsPresent() {
		return f(o.value)
	}
	return Empty[R]()
}
