package utils

import "reflect"

func ShallowMergeStructs[T any](dest *T, src *T) *T {
	// Iterate over the fields of the source struct
	// and copy their values to the destination struct
	res := reflect.New(reflect.TypeOf(*dest)).Elem()
	res.Set(reflect.ValueOf(dest).Elem())

	srcValue := reflect.ValueOf(src).Elem()

	for i := 0; i < srcValue.NumField(); i++ {
		resField := res.Field(i)
		srcField := srcValue.Field(i)

		// Check if the source field is non-zero, then update the destination field
		if !srcField.IsZero() {
			resField.Set(srcField)
		}
	}

	resData := res.Interface().(T)

	return &resData
}
