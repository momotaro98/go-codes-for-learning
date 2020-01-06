package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	c := 1

	go func() {
		mu.Lock()
		defer mu.Unlock()
		c *= 2
		time.Sleep(time.Second * 3)
		c *= 2
	}()

	time.Sleep(time.Second * 1)
	mu.Lock() // ロックが解除されるまで待つ
	c *= 3
	mu.Unlock()

	fmt.Println(c) // What is printed? => 12
	// mu.Lock(); c *= 3; mu.Unlock がないと特にロックされないので6が出力される。
}
