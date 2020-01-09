package main

import (
	"fmt"
	"sync"
)

// $ go run s3-2-5_Pool_ex1.go                                                                                                                         master ◼
// 4 calculators created.

func main() {
	var numCalcsCreated int

	calcPool := &sync.Pool{
		New: func() interface{} {
			numCalcsCreated++
			mem := make([]byte, 1024)
			return &mem // [Note] Return pointer type
			// Official doc says
			// > The Pool's New function should generally only
			// > return pointer types, since a pointer can be put
			// > into the return interface value without an allocation
		},
	}

	// Allocate 3kB
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())

	const numWorkers = 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func() {
			defer wg.Done()

			mem := calcPool.Get().(*[]byte) // Check the [Note] above
			defer calcPool.Put(mem)

			// Do something interesting
			// It should be, however, fast process to the memory
		}()
	}

	wg.Wait()
	fmt.Printf("%d calculators created.", numCalcsCreated)
}
