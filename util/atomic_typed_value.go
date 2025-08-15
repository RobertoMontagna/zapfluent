package util

import "sync/atomic"

type AtomicTypedValue[T any] struct {
	value atomic.Value
}

func NewAtomicTypedValue[T any](initialValue T) *AtomicTypedValue[T] {
	var v AtomicTypedValue[T]
	v.Store(initialValue)
	return &v
}

func (v *AtomicTypedValue[T]) Load() T {
	return v.value.Load().(T)
}

func (v *AtomicTypedValue[T]) Store(val T) {
	v.value.Store(val)
}

func (v *AtomicTypedValue[T]) Swap(new T) (old T) {
	return v.value.Swap(new).(T)
}

func (v *AtomicTypedValue[T]) CompareAndSwap(old, new T) (swapped bool) {
	return v.value.CompareAndSwap(old, new)
}
