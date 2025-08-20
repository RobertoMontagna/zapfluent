package lazyoptional

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
