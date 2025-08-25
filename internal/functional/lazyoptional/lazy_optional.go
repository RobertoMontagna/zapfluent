package lazyoptional

import (
	"go.robertomontagna.dev/zapfluent/internal/functional/lazy"
	"go.robertomontagna.dev/zapfluent/internal/functional/optional"
)

// LazyOptional is a container object which may or may not contain a non-nil
// value. It is "lazy" because it is built upon the Lazy monad, deferring
// the retrieval or computation of a value until it is explicitly requested
// by a terminal operation (like `Get`).
//
// This implementation composes `lazy.Lazy` and `optional.Optional` to achieve
// its behavior.
type LazyOptional[T any] struct {
	value lazy.Lazy[optional.Optional[T]]
}

// Some returns a LazyOptional that contains the given non-nil value.
func Some[T any](value T) LazyOptional[T] {
	return LazyOptional[T]{
		value: lazy.Of(optional.Some(value)),
	}
}

// Empty returns a LazyOptional that does not contain a value.
func Empty[T any]() LazyOptional[T] {
	return LazyOptional[T]{
		value: lazy.Of(optional.Empty[T]()),
	}
}

// Get retrieves the value from the LazyOptional.
//
// It returns the value and `true` if a value is present, or the zero value of
// the type and `false` if the optional is empty. This is a terminal operation
// that triggers the evaluation of the underlying Lazy value.
func (o LazyOptional[T]) Get() (T, bool) {
	return o.value.Get().Get()
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

// FlatMap transforms the value inside a LazyOptional by applying a function
// that itself returns a LazyOptional.
func FlatMap[T any, U any](o LazyOptional[T], f func(T) LazyOptional[U]) LazyOptional[U] {
	return LazyOptional[U]{
		value: lazy.FlatMap(&o.value, func(opt optional.Optional[T]) lazy.Lazy[optional.Optional[U]] {
			val, ok := opt.Get()
			if !ok {
				return lazy.Of(optional.Empty[U]())
			}
			return f(val).value
		}),
	}
}

// Map transforms the value inside a LazyOptional by applying a function to it.
func Map[T any, U any](o LazyOptional[T], mapper func(T) U) LazyOptional[U] {
	return LazyOptional[U]{
		value: lazy.Map(&o.value, func(opt optional.Optional[T]) optional.Optional[U] {
			return optional.Map(opt, mapper)
		}),
	}
}
