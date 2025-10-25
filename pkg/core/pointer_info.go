package core

import (
	"fmt"

	"go.uber.org/zap/zapcore"
)

// PointerInfo represents a structure holding a pointer to a value of
// generic type T and associated utility functions.
type PointerInfo[T any] struct {
	PtrValue  *T
	functions typeFieldFunctions[T]
}

// MarshalLogObject implements the zapcore.ObjectMarshaler interface for
// PointerInfo. It encodes the pointer's value and its memory address.
func (p PointerInfo[T]) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	if p.PtrValue == nil {
		return AsFluent(enc).
			Add(String("address", "0x0")).
			Add(String("value", "<nil>")).
			Done()
	}

	return AsFluent(enc).
		Add(String("address", fmt.Sprintf("%p", p.PtrValue))).
		Add(p.functions.toField("value", *p.PtrValue)).
		Done()
}

func (p PointerInfo[T]) isNonZero() bool {
	return p.PtrValue != nil && p.functions.isNonZero(*p.PtrValue)
}
