package optional

import "reflect"

// isNil checks if a value is nil using reflection.
// It correctly handles interfaces, pointers, channels, functions, maps, and slices.
func isNil[T any](v T) bool {
	if any(v) == nil {
		return true
	}

	val := reflect.ValueOf(v)
	switch val.Kind() {
	case reflect.Chan,
		reflect.Func,
		reflect.Map,
		reflect.Pointer,
		reflect.Slice,
		reflect.Interface:
		return val.IsNil()
	default:
		return false
	}
}
