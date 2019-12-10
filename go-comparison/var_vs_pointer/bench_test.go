package main

import (
	"testing"
)

func BenchmarkVar(b *testing.B) {
	var sum int
	for i := 0; i < b.N; i++ {
		v := NewMyStructVar()
		sum += v.arr[0]
	}
}

func BenchmarkPointer(b *testing.B) {
	var sum int
	for i := 0; i < b.N; i++ {
		v := NewMyStructPtr()
		sum += v.arr[0]
	}
}
