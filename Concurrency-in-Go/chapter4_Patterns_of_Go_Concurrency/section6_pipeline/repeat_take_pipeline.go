package main

import "fmt"

// $ go run XXX.go
// 1 2 1 2 1 2 1 2 1 2

// repeatには無限に生成する能力があるがtakeが受け取るまでブロックするので有限しか生成しない。
// doneがCloseになるのでrepeatは終了することができる。

func main() {
	repeat := func(done <-chan interface{}, values ...interface{}) <-chan interface{} {
		valueCh := make(chan interface{})
		go func() {
			defer close(valueCh)
			for {
				for _, v := range values {
					select {
					case <-done:
						return
					case valueCh <- v: // select では送信もcaseに入る
					}
				}
			}
		}()
		return valueCh
	}

	take := func(done <-chan interface{}, valueCh <-chan interface{}, num int) <-chan interface{} {
		takeCh := make(chan interface{})
		go func() {
			defer close(takeCh)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case takeCh <- <-valueCh: // 受信したものを送信するのでこの書き方
				}
			}
		}()
		return takeCh
	}

	done := make(chan interface{})
	defer close(done)

	for num := range take(done, repeat(done, 1, 2), 10) {
		fmt.Printf("%v ", num)
	}
}
