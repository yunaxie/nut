package util

import (
	"reflect"
)

func IsNil(obj interface{}) bool {
	switch reflect.ValueOf(obj).Kind() {
	case reflect.String:
		return obj == ""
	case reflect.Struct:
		return obj == struct {}{}
	default:
		return obj == nil
	}
}

func IsNotNil(obj interface{}) bool {
	return !IsNil(obj)
}
