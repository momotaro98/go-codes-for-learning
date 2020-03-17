package main

import "time"

// DoWork はテスト対象のメソッド
// intのチャネルを返す仕事をする。
// ハートビートを持っている。
func DoWork(
	done <-chan interface{},
	nums ...int,
) (<-chan interface{}, <-chan int) {
	heartbeat := make(chan interface{}, 1)
	intStream := make(chan int)
	go func() {
		defer close(heartbeat)
		defer close(intStream)

		time.Sleep(2 * time.Second) // 何らかの遅延をシミュレートしている。
		// 実際にはこの遅延はあらゆる理由で発生しうるもので非決定的。
		// CPU負荷、ディスクの競合、ネットワーク遅延、などが想定される。

		for _, n := range nums {
			select {
			case heartbeat <- struct{}{}: // 仕事を1つ開始する度にハートビートを返すパターン
			default:
			}

			select {
			case <-done:
				return
			case intStream <- n:
			}

		}
	}()

	return heartbeat, intStream
}
