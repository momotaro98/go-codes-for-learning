package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	done := make(chan interface{})
	defer close(done)

	start := time.Now()

	rand := func() interface{} { return rand.Intn(50000000) }

	randIntCh := toInt(done, repeatFn(done, rand))

	// fan_out (ファンアウト)で処理に時間がかかる primeClassifier を並行にさばく
	// 素数を発見することは順不同なのでファンアウト、ファンインが使える。

	// Fan out
	numFinders := runtime.NumCPU()
	fmt.Printf("Spinning up %d prime finders.\n", numFinders)
	finders := make([]<-chan int, numFinders)
	fmt.Println("Primes:")
	for i := 0; i < numFinders; i++ {
		finders[i] = primeClassifier(done, randIntCh)
	}

	// Fan in
	for prime := range take(done, fanIn(done, finders...), 10) {
		fmt.Printf("\t%d\n", prime)
	}

	fmt.Printf("Search took: %v", time.Since(start))
}
