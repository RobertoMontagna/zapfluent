package zapfluent

type LazyOptional[T any] struct {
	producer func() (T, bool)
}

func NewLazyOptional[T any](value T) LazyOptional[T] {
	return LazyOptional[T]{
		producer: func() (T, bool) {
			return value, true
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

func (o LazyOptional[T]) MapToString(mapper func(T) string) LazyOptional[string] {
	return LazyOptional[string]{
		producer: func() (string, bool) {
			val, ok := o.producer()
			if !ok {
				var zero string
				return zero, false
			}
			return mapper(val), true
		},
	}
}
