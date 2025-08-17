package lazy

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
