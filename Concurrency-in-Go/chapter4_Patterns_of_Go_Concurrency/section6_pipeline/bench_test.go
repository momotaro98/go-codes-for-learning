package main

import (
	"testing"
)

// $ go test -bench .
// goos: darwin
// goarch: amd64
// BenchmarkGeneric-4        869703              1280 ns/op
// BenchmarkTyped-4         1479144               839 ns/op

// > interface{}をStringにする処理をした方が1.5倍ほど処理がかかるが単位で言えばごく小さな差です。

func BenchmarkGeneric(b *testing.B) {
	done := make(chan interface{})
	defer close(done)

	b.ResetTimer()
	for range toString(done, take(done, repeat(done, "a"), b.N)) {
	}
}

func BenchmarkTyped(b *testing.B) {
	repeat := func(done <-chan interface{}, values ...string) <-chan string {
		valueCh := make(chan string)
		go func() {
			defer close(valueCh)
			for {
				for _, v := range values {
					select {
					case <-done:
						return
					case valueCh <- v:
					}
				}
			}
		}()
		return valueCh
	}
	take := func(done <-chan interface{}, valueCh <-chan string, num int) <-chan string {
		takeCh := make(chan string)
		go func() {
			defer close(takeCh)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case takeCh <- <-valueCh:
				}
			}
		}()
		return takeCh
	}

	done := make(chan interface{})
	defer close(done)

	b.ResetTimer()
	for range take(done, repeat(done, "a"), b.N) {
	}
}
