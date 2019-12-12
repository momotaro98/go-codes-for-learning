package main

import (
	"testing"
)

func BenchmarkVar(b *testing.B) {
	var sum int // short_life
	//var list []myStruct // long_life
	for i := 0; i < b.N; i++ {
		v := NewMyStructVar()
		sum += v.arr[0] // short_life
		// list = append(list, v) // long_life
	}
}

func BenchmarkPointer(b *testing.B) {
	var sum int // short_life
	//var list []myStruct // long_life
	for i := 0; i < b.N; i++ {
		v := NewMyStructPtr()
		sum += v.arr[0] // short_life
		// list = append(list, v) // long_life
	}
}
