package main

import "sync"

// ファンイン(fan in)というのはマルチプレキシング、つまり複数のデータのストリームを、
// 単一のストリームに統合することを意味します。

func fanIn(done <-chan interface{}, channels ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	multiplexedCh := make(chan int)

	multiplex := func(c <-chan int) {
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case multiplexedCh <- i:
			}
		}
	}

	// select from all channels
	wg.Add(len(channels))
	for _, c := range channels {
		go multiplex(c)
	}

	// Wait for all the reads to complete
	go func() {
		wg.Wait()
		close(multiplexedCh)
	}()

	return multiplexedCh
}
