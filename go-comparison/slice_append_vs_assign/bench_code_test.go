package main

import "testing"

// BenchmarkAppendNakedMap-4        6103308               187 ns/op             248 B/op          5 allocs/op
// BenchmarkAssignMap-4            20554850                52.3 ns/op            80 B/op          1 allocs/op
// BenchmarkAppendCapMap-4         20464335                56.6 ns/op            80 B/op          1 allocs/op

var (
	s = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	l = len(s)
	f = func(i int) int { return i * 2 }
)

func BenchmarkAppendNakedMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = AppendNakedMap(s, l, f)
	}
}

func BenchmarkAssignMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = AssignMap(s, l, f)
	}
}

func BenchmarkAppendCapMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = AppendCapMap(s, l, f)
	}
}
