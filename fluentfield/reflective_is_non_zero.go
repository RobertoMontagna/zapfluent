package fluentfield

import "reflect"

// ReflectiveIsNotNil checks if a value is non-nil using reflection.
//
// This function is useful as an `isNonZero` implementation for object fields
// where the zero-value check is simply a nil check. It correctly handles
// interfaces, pointers, channels, functions, maps, and slices. For other
// types, it returns true.
func ReflectiveIsNotNil[T any](v T) bool {
	if any(v) == nil {
		return false
	}

	val := reflect.ValueOf(v)
	switch val.Kind() {
	case reflect.Chan,
		reflect.Func,
		reflect.Map,
		reflect.Pointer,
		reflect.Slice,
		reflect.Interface:
		return !val.IsNil()
	default:
		return true
	}
}
