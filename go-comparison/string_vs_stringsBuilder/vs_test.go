package main

import (
	"strings"
	"testing"
)

func Benchmark_string(b *testing.B) {
	var s string
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		s += "a"
	}
	b.StopTimer()
}

func Benchmark_stringsBuilder(b *testing.B) {
	var sb strings.Builder
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sb.WriteString("a")
	}
	_ = sb.String() // output string
	b.StopTimer()
}
