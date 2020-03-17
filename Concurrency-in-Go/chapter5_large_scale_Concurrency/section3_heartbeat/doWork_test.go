package main

import (
	"testing"
	"time"
)

// これは悪いテスト
// 悪い理由は決定的だからです。テスト対象の実行時間が非決定の場合、
// テストはパスしたり失敗したりまちまちになります。
func TestDoWork_GeneratesAllNumbers_BAD(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	intSlice := []int{0, 1, 2, 3, 5}
	_, results := DoWork(done, intSlice...)

	for i, expected := range intSlice {
		select {
		case r := <-results:
			if r != expected {
				t.Errorf(
					"index %v: expected %v, but received %v",
					i, expected, r,
				)
			}
		case <-time.After(1 * time.Second): // 壊れたゴルーチンがテストで
			// デッドロックを起こしてしまわないのに十分と思われる時間が
			// 経過したあとでタイムアウトしています。
			t.Fatal("test time out")
		}
	}
}

// これは良いテスト
// ハートビートを待つおかげでタイムアウトを使わずに安全にテストが書けます。
func TestDoWork_GeneratesAllNumbers_GOOD(t *testing.T) {
	done := make(chan interface{})
	defer close(done)

	intSlice := []int{0, 1, 2, 3, 5}
	heartbeat, results := DoWork(done, intSlice...)

	<-heartbeat // ゴルーチンが繰り返しを始めるというシグナルを送るのを待ちます。

	i := 0
	for r := range results {
		if expected := intSlice[i]; r != expected {
			t.Errorf("index %v: expected %v, but received %v",
				i, expected, r)
		}
		i++
	}
}
