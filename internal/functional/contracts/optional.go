// Package contracts defines shared interfaces and contracts that are used across
// different functional components of the library.
package contracts

// OptionalLike is an interface that describes the common behavior of optional-like
// types. It is used to create generic matchers and other utilities that can
// operate on any type that satisfies this contract.
type OptionalLike[T any] interface {
	// IsPresent returns true if the optional contains a value.
	IsPresent() bool
	// Get returns the value if present, and true. Otherwise, it returns the
	// zero value of the type and false.
	Get() (T, bool)
}
