package lazyoptional

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
