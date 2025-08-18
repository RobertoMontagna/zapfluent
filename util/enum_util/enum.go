package enum_util

import "fmt"

type EnumValue interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type UtilEnum[T EnumValue] struct {
	values  map[T]string
	unknown T
	min     T
	max     T
}

func NewUtilEnum[T EnumValue](values map[T]string, unknown T, min T, max T) UtilEnum[T] {
	return UtilEnum[T]{
		values:  values,
		unknown: unknown,
		min:     min,
		max:     max,
	}
}

func (e UtilEnum[T]) String(v T) string {
	if s, ok := e.values[v]; ok {
		return s
	}
	return fmt.Sprintf("Unknown(%d)", v)
}

func (e UtilEnum[T]) FromInt(i int) T {
	v := T(i)
	if v >= e.min && v <= e.max {
		return v
	}
	return e.unknown
}
