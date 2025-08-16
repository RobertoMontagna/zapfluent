package lazy

func ConstantFunction[T any](value T) func() (T, bool) {
	return func() (T, bool) {
		return value, true
	}
}
