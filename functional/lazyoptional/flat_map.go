package lazyoptional

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
