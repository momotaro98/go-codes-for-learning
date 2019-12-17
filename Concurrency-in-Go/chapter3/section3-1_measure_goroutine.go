package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Purpose: Let's measure how many kilo bytes a goroutine is

func MeasureGoroutine() {
	memConsume := func() uint64 {
		runtime.GC()           // 意図的にGCを走らすことができる
		var s runtime.MemStats // メモリ状況を得ることができる
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	var c <-chan interface{} // <-chan T は受信専用チャネル型
	var wg sync.WaitGroup
	noop := func() { wg.Done(); <-c } // noop func never finish

	const numGoroutines = 1e5 // Goのリテラルで 10^5 を表せる
	wg.Add(numGoroutines)
	before := memConsume() // Before generating goroutines
	for i := 0; i < numGoroutines; i++ {
		go noop()
	}
	wg.Wait()
	after := memConsume() // After generating goroutines

	fmt.Printf("%.3fkb", float64(after-before)/numGoroutines/1000)
	// 2.113kb が出力された
	// 環境 go version go1.13.1 darwin/amd64
	// 書籍では2.817kb とあるのでGoのバージョンアップによりさらにGoroutineが軽量になったと思われる。
}
