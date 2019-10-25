package main

import (
	"reflect"
)

func walk(x interface{}, fn func(input string)) {
	// Reflectionで値情報を取得
	val := getValue(x)

	switch val.Kind() {
	case reflect.String:
		fn(val.String()) // Edge case (End of 再帰)
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walk(val.Field(i).Interface(), fn)
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walk(val.Index(i).Interface(), fn)
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walk(val.MapIndex(key).Interface(), fn)
		}
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}
