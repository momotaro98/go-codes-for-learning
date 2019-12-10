package main

import (
	"testing"
)

func BenchmarkVar(b *testing.B) {
	var list []myStruct
	for i := 0; i < b.N; i++ {
		v := NewMyStructVar()
		list = append(list, v)
	}
}

func BenchmarkPointer(b *testing.B) {
	var list []*myStruct
	for i := 0; i < b.N; i++ {
		v := NewMyStructPtr()
		list = append(list, v)
	}
}
