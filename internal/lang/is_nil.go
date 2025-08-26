package lang

import "reflect"

// ReflectiveIsNil checks if a value is nil using reflection.
// It correctly handles interfaces, pointers, channels, functions, maps, and slices.
func ReflectiveIsNil(v any) bool {
	if v == nil {
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
