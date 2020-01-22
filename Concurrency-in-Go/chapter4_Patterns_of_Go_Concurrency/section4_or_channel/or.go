package main

import (
	"fmt"
	"time"
)

/// まとめているチャネルのうちどれか1つのチャネルが閉じられたら、
/// まとめたチャネルも閉じられるようにしたいと思うことがあるでしょう。
/// このorパターンはシステム内で複数のモジュールを組み合わせる際の繋ぎ目として利用すると便利です。
/// 同様のことを行う他の方法を4.12 contextパッケージの節で紹介します。
/// また、このパターンの変形を使ってより複雑なパターンを構成する方法を5.4複製されたリクエストで紹介します。

// $ go run or.go
// done after 1.003693475s

// or is a function that binds all channels
func or(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	orDone := make(chan interface{})
	go func() {
		defer close(orDone)

		switch len(channels) {
		case 2: // channels[2]でIndex out error になるのを避けるため
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			case <-or(append(channels[3:], orDone)...): // 再帰
				// case <-or(channels[3:]...): // こっちでも動くがorDoneを渡すことで
				// `defer close(Done)`が効いて最後に呼ばれた関数を含め
				// すべてのゴルーチンが終了するようになる。
			}
		}
	}()
	return orDone
}

// main is a consumer of or function
func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("done after %v", time.Since(start))
}
