package main

import (
	"testing"
)

var (
	sl = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ln = len(sl)
	fn = func(i int) int { return i * 2 }
)

func AppendNakedMap(in []int, length int, f func(i int) int) []int {
	var ret []int
	for _, c := range in {
		ret = append(ret, f(c))
	}
	return ret
}

func BenchmarkAppendNakedMap(b *testing.B) {
	s, l, f := sl, ln, fn
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = AppendNakedMap(s, l, f)
	}
}

func AssignMap(in []int, length int, f func(i int) int) []int {
	ret := make([]int, length)
	for i, c := range in {
		ret[i] = f(c)
	}
	return ret
}

func BenchmarkAssignMap(b *testing.B) {
	s, l, f := sl, ln, fn
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = AssignMap(s, l, f)
	}
}

func AppendCapMap(in []int, length int, f func(i int) int) []int {
	ret := make([]int, 0, length)
	for _, c := range in {
		ret = append(ret, f(c))
	}
	return ret
}

func BenchmarkAppendCapMap(b *testing.B) {
	s, l, f := sl, ln, fn
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = AppendCapMap(s, l, f)
	}
}
