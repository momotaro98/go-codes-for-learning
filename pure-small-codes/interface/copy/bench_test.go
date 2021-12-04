package main

import (
	"testing"
)

func BenchmarkF(t *testing.B) {
	for i := 0; i < t.N; i++ {
		f()
	}
}

func BenchmarkG(t *testing.B) {
	for i := 0; i < t.N; i++ {
		g()
	}
}
