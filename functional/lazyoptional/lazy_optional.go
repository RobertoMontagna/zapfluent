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

// NewConstantProducer returns a producer function that, when called, always
// returns the same two values that were provided at creation time.
// This is useful for creating the start of a lazy evaluation chain, such as
// in `Some` or `Empty`.
func NewConstantProducer[T1 any, T2 any](v1 T1, v2 T2) func() (T1, T2) {
	return func() (T1, T2) {
		return v1, v2
	}
}

// FlatMap transforms the value inside a LazyOptional by applying a function
// that itself returns a LazyOptional.
//
// If the input optional is empty, the result is an empty optional. Otherwise,
// the function `f` is applied to the value, and the resulting optional is
// returned. This is a fundamental operation for chaining lazy computations.
func FlatMap[T any, U any](o LazyOptional[T], f func(T) LazyOptional[U]) LazyOptional[U] {
	return LazyOptional[U]{
		producer: func() (U, bool) {
			val, ok := o.Get()
			if !ok {
				var zero U
				return zero, false
			}
			return f(val).Get()
		},
	}
}

// Map transforms the value inside a LazyOptional by applying a function to it.
//
// If the input optional is empty, the result is an empty optional. Otherwise,
// the `mapper` function is applied to the value, and a new optional containing
// the result is returned.
func Map[T any, U any](o LazyOptional[T], mapper func(T) U) LazyOptional[U] {
	return LazyOptional[U]{
		producer: func() (U, bool) {
			val, ok := o.Get()
			if !ok {
				var zero U
				return zero, false
			}
			return mapper(val), true
		},
	}
}
