package utils

import (
	"reflect"
)

func IsSlice(data interface{}) bool {
	return reflect.ValueOf(data).Kind() == reflect.Slice
}
