package main

import (
	"fmt"
	"sync"
)

type PingPongPayload struct {
	Counter int
}

func ExamplePingPong() {
	var p PingPongPayload
	chA := make(chan *PingPongPayload)
	chB := make(chan *PingPongPayload)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			p, ok := <-chA // receive from chA
			if !ok {
				break
			}
			fmt.Printf("chA: p.Counter = %d\n", p.Counter)
			p.Counter++ // rewrite data of PingPongPayload instance
			if p.Counter > 6 {
				break
			}

			chB <- p // send to chB
		}
		close(chB) // Close chB after break because this goroutine writes to chB
	}()

	go func() {
		defer wg.Done()
		for {
			p, ok := <-chB
			if !ok {
				break
			}
			fmt.Printf("chB: p.Counter = %d\n", p.Counter)
			p.Counter++ // rewrite data of PingPongPayload instance
			if p.Counter > 6 {
				break
			}

			chA <- p // send to chA
		}
		close(chA)
	}()

	chA <- &p // 1st wirte to chA
	wg.Wait() // wating for end of process
}

func main() {
	ExamplePingPong()
}
