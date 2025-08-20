package lazyoptional

// LazyOptional is a container object which may or may not contain a non-nil
// value. It is "lazy" because it operates on a producer function (`func() (T, bool)`),
// deferring the retrieval or computation of a value until it is explicitly
// requested by a terminal operation (like `Get`).
//
// This allows for chaining multiple transformations (e.g., `Map`, `Filter`)
// without intermediate allocations or computations. The chain is only evaluated
// when a result is needed.
type LazyOptional[T any] struct {
	producer func() (T, bool)
}

// Some returns a LazyOptional that contains the given non-nil value.
func Some[T any](value T) LazyOptional[T] {
	return LazyOptional[T]{
		producer: NewConstantProducer(value, true),
	}
}

// Empty returns a LazyOptional that does not contain a value.
func Empty[T any]() LazyOptional[T] {
	var zero T
	return LazyOptional[T]{
		producer: NewConstantProducer(zero, false),
	}
}

// Get retrieves the value from the LazyOptional.
//
// It returns the value and `true` if a value is present, or the zero value of
// the type and `false` if the optional is empty. This is a terminal operation
// that triggers the evaluation of the producer chain.
func (o LazyOptional[T]) Get() (T, bool) {
	return o.producer()
}

// Filter returns a new LazyOptional that will be empty if the original
// optional was empty or if the value does not satisfy the given condition.
func (o LazyOptional[T]) Filter(condition func(T) bool) LazyOptional[T] {
	return FlatMap(o, func(v T) LazyOptional[T] {
		if condition(v) {
			return Some(v)
		}
		return Empty[T]()
	})
}
