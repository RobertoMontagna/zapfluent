package lazy

type LazyOptional[T any] struct {
	producer func() (T, bool)
}

func Some[T any](value T) LazyOptional[T] {
	return LazyOptional[T]{
		producer: ConstantFunction(value),
	}
}

func Empty[T any]() LazyOptional[T] {
	return LazyOptional[T]{
		producer: func() (T, bool) {
			var zero T
			return zero, false
		},
	}
}

func (o LazyOptional[T]) Get() (T, bool) {
	return o.producer()
}

func (o LazyOptional[T]) Filter(condition func(T) bool) LazyOptional[T] {
	return LazyOptional[T]{
		producer: func() (T, bool) {
			val, ok := o.producer()
			if !ok {
				var zero T
				return zero, false
			}
			if condition(val) {
				return val, true
			}
			var zero T
			return zero, false
		},
	}
}
