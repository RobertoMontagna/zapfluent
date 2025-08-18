package lazyoptional

func NewConstantProducer[T1 any, T2 any](v1 T1, v2 T2) func() (T1, T2) {
	return func() (T1, T2) {
		return v1, v2
	}
}
