package main

import (
	"sync"
	"testing"
)

/*
$ go test -bench=BenchmarkContextSwitch -cpu=1 section3-1_contextswitch_test.go
goos: darwin
goarch: amd64
BenchmarkContextSwitch   7945143               146 ns/op
PASS
ok      command-line-arguments  1.360s
*/

// > OSのコンテキストスイッチよりも10倍くらい速い
// > Goroutineを生成すればするほど、プログラムはマルチコアCPUでスケールするでしょう。
// → Goroutineは生成コストもコンテキストスイッチコストも低いので
// Goroutineをどんどん使っていこうぜ

// [補足] 書籍では225 ns/op でありGo1.13.1バージョンでさらにパフォーマンスが改善されていると考えられる。

func BenchmarkContextSwitch(b *testing.B) {
	var wg sync.WaitGroup
	begin := make(chan struct{})
	c := make(chan struct{}) // struct{} はメモリを消費しない
	// これによって、メッセージを送出する時間だけを計測できる。

	var token struct{}
	sender := func() {
		defer wg.Done()
		<-begin
		for i := 0; i < b.N; i++ {
			c <- token
		}
	}
	receiver := func() {
		defer wg.Done()
		<-begin
		for i := 0; i < b.N; i++ {
			<-c
		}
	}

	wg.Add(2)
	go sender()
	go receiver()
	b.StartTimer() // benchmarkで用意されているので使おう
	close(begin)   // `<-begin`のブロックを解く
	wg.Wait()
}
