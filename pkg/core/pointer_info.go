package core

import (
	"fmt"

	"go.uber.org/zap/zapcore"
)

// PointerInfo represents a structure holding a pointer to a value of
// generic type T and associated utility functions.
type PointerInfo[T any] struct {
	PtrValue  *T
	functions typePointerFieldFunctions[T]
}

// MarshalLogObject implements the zapcore.ObjectMarshaler interface for
// PointerInfo. It encodes the pointer's value and its memory address.
func (p PointerInfo[T]) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	if p.PtrValue == nil {
		return AsFluent(enc).
			Add(String("address", NilPtrAddress)).
			Add(String("value", NilSentinel)).
			Done()
	}

	return AsFluent(enc).
		Add(String("address", fmt.Sprintf("%p", p.PtrValue))).
		Add(p.functions.toField("value", *p.PtrValue)).
		Done()
}

func (p PointerInfo[T]) isNonZero() bool {
	// isNonZero reports whether the pointer is non-nil.
	// Note: underlying zero values are considered non-zero if the pointer itself is set.
	return p.PtrValue != nil
}
