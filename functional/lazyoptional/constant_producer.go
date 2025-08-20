package lazyoptional

// NewConstantProducer returns a producer function that, when called, always
// returns the same two values that were provided at creation time.
// This is useful for creating the start of a lazy evaluation chain, such as
// in `Some` or `Empty`.
func NewConstantProducer[T1 any, T2 any](v1 T1, v2 T2) func() (T1, T2) {
	return func() (T1, T2) {
		return v1, v2
	}
}
