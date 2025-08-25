package lazy

import "sync"

// Lazy is a monad that represents a value that is computed lazily.
// The value is computed only when it is needed, and the result is memoized.
type Lazy[T any] struct {
	value func() T
	once  sync.Once
	memo  T
}

// New creates a new Lazy instance from a function.
// The function will be called only once, the first time the value is accessed.
func New[T any](f func() T) Lazy[T] {
	return Lazy[T]{
		value: f,
	}
}

// Of creates a new Lazy instance from an already computed value.
func Of[T any](v T) Lazy[T] {
	return Lazy[T]{
		memo: v,
	}
}

// Get returns the value of the Lazy instance.
// The underlying function is called only the first time.
func (l *Lazy[T]) Get() T {
	l.once.Do(func() {
		if l.value != nil {
			l.memo = l.value()
		}
	})
	return l.memo
}

// Map applies a function to the value of the Lazy instance.
// It takes a pointer to ensure that memoization state is correctly handled
// on the original Lazy instance. The result is a new Lazy instance.
func Map[A, B any](l *Lazy[A], f func(A) B) Lazy[B] {
	return New(func() B {
		return f(l.Get())
	})
}

// FlatMap applies a function that returns a Lazy instance to the value of the Lazy instance.
// It takes a pointer to ensure that memoization state is correctly handled
// on the original Lazy instance. The result is the new Lazy instance.
func FlatMap[A, B any](l *Lazy[A], f func(A) Lazy[B]) Lazy[B] {
	return New(func() B {
		lazyB := f(l.Get())
		return lazyB.Get()
	})
}
