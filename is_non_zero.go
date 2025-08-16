package zapfluent

import "reflect"

func IsNotNil[T any](v T) bool {
	if any(v) == nil {
		return false
	}
	val := reflect.ValueOf(v)
	switch val.Kind() {
	case reflect.Chan, reflect.Func, reflect.Map, reflect.Pointer, reflect.Slice, reflect.Interface:
		return !val.IsNil()
	}
	return true
}
