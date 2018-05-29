package main

import (
	"fmt"
	"time"
)

func ExampleTicker() {
	t := time.NewTicker(2 * time.Second)
	defer t.Stop()

	for i := 0; i < 5; i++ {
		select {
		case <-t.C: // Invoke every 2 seconds
			fmt.Printf("%d\n", i)
		}
	}
}

func main() {
	ExampleTicker()
}
