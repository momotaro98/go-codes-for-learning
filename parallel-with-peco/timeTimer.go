package main

import (
	"fmt"
	"time"
)

// 実行すると特に2秒後に終了せず無限ループになってしまう

func ExampleTimer() {
	t := time.NewTimer(2 * time.Second)
	defer t.Stop()

	for {
		select {
		case <-t.C:
			break
		default:
			fmt.Println("process...")
		}
	}
}

func main() {
	ExampleTimer()
}
