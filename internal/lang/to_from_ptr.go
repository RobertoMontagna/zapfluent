package lang

func ToPtr[T any](value T) *T {
	return &value
}

func MustFromPtr[T any](ptr *T) T {
	return *ptr
}
