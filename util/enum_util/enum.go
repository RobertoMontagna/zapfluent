package enum_util

import "fmt"

type EnumValue interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type UtilEnum[E EnumValue] struct {
	values  map[E]string
	unknown E
}

func NewUtilEnum[E EnumValue](values map[E]string, unknown E) UtilEnum[E] {
	return UtilEnum[E]{
		values:  values,
		unknown: unknown,
	}
}

func (e UtilEnum[E]) String(v E) string {
	if s, ok := e.values[v]; ok {
		return s
	}
	return fmt.Sprintf("Unknown(%d)", v)
}

func (e UtilEnum[E]) FromInt(i int) E {
	k := E(i)
	if _, ok := e.values[k]; ok {
		return k
	}
	return e.unknown
}
