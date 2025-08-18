package fluentfield

type TypedField[T any] interface {
	Field
	Filter(condition func(T) bool) TypedField[T]
	NonZero() TypedField[T]
	Format(formatter func(T) string) TypedField[string]
}
