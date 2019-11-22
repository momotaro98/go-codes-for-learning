package main

import (
	"math/rand"
	"sort"
	"testing"
)

// How to run benchmark
// go test -bench . -benchmem

func prepare(length int) sort.IntSlice {
	s := make([]int, length)
	for i := 0; i < length; i++ {
		s[i] = rand.Int()
	}
	return sort.IntSlice(s)
}

func BenchmarkSort(b *testing.B) {
	iSlice := prepare(b.N)
	b.ResetTimer()
	sort.Sort(iSlice)
}

func BenchmarkStable(b *testing.B) {
	iSlice := prepare(b.N)
	b.ResetTimer()
	sort.Stable(iSlice)
}
