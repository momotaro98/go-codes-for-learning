package main

import (
	"fmt"
	"sync"
	"time"
)

func ExampleCond() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	c := sync.NewCond(&mu)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			defer mu.Unlock()

			fmt.Printf("waiting %d\n", i)
			mu.Lock() // lock with Cond wait
			c.Wait()  // waiting for Cond.Signal()
			fmt.Printf("go %d\n", i)
		}(i)
	}

	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("signaling!\n")
		// Notify
		c.Signal()
	}

	wg.Wait()
}

func main() {
	ExampleCond()
}
