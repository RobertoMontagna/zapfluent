package optional

// LiftToOptional converts a function that returns a pointer into a function
// that returns an Optional. If the pointer is nil, the optional is empty.
func LiftToOptional[T any](f func() *T) func() Optional[T] {
	return func() Optional[T] {
		return OfPtr(f())
	}
}

// LiftToOptional1 converts a function with one argument that returns a pointer
// into a function that returns an Optional.
func LiftToOptional1[T any, P1 any](f func(P1) *T) func(P1) Optional[T] {
	return func(p1 P1) Optional[T] {
		return OfPtr(f(p1))
	}
}

// LiftToOptional2 converts a function with two arguments that returns a pointer
// into a function that returns an Optional.
func LiftToOptional2[T any, P1 any, P2 any](f func(P1, P2) *T) func(P1, P2) Optional[T] {
	return func(p1 P1, p2 P2) Optional[T] {
		return OfPtr(f(p1, p2))
	}
}

// LiftErrorToOptional converts a function that returns an error into a function
// that returns an Optional[error]. If the error is nil, the optional is empty.
func LiftErrorToOptional(f func() error) func() Optional[error] {
	return func() Optional[error] {
		return OfError(f())
	}
}

// LiftErrorToOptional1 converts a function with one argument that returns an
// error into a function that returns an Optional[error].
func LiftErrorToOptional1[P1 any](f func(P1) error) func(P1) Optional[error] {
	return func(p1 P1) Optional[error] {
		return OfError(f(p1))
	}
}

// LiftErrorToOptional2 converts a function with two arguments that returns an
// error into a function that returns an Optional[error].
func LiftErrorToOptional2[P1 any, P2 any](f func(P1, P2) error) func(P1, P2) Optional[error] {
	return func(p1 P1, p2 P2) Optional[error] {
		return OfError(f(p1, p2))
	}
}
