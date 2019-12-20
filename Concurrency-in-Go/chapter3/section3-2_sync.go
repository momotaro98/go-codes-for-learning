package main

import (
	"fmt"
	"math"
	"os"
	"sync"
	"text/tabwriter"
	"time"
)

//// From "sync" package
// A Locker represents an object that can be locked and unlocked.
//type Locker interface {
//	Lock()
//	Unlock()
//}

func main() {
	producer := func(wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()
		l.Lock()
		l.Unlock()
		time.Sleep(1)
	}

	observer := func(wg *sync.WaitGroup, l sync.Locker) {
		defer wg.Done()
		l.Lock()
		defer l.Unlock()
	}

	test := func(count int, mutex, rwMutex sync.Locker) time.Duration {
		var wg sync.WaitGroup
		wg.Add(count + 1)
		beginTestTime := time.Now()
		go producer(&wg, mutex)
		for i := count; i > 0; i-- {
			go observer(&wg, rwMutex)
		}
		wg.Wait()
		return time.Since(beginTestTime)
	}

	tw := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ', 0)
	defer tw.Flush()

	//// From "sync" package
	// A RWMutex is a reader/writer mutual exclusion lock.
	// The lock can be held by an arbitrary number of readers or a single writer.
	// The zero value for a RWMutex is an unlocked mutex.
	//
	// A RWMutex must not be copied after first use.
	//
	// If a goroutine holds a RWMutex for reading and another goroutine might
	// call Lock, no goroutine should expect to be able to acquire a read lock
	// until the initial read lock is released. In particular, this prohibits

	//// From 書籍 about RWMutex
	// メモリに対する読み込みのロックを要求した場合、
	// ロックが書き込みで保持されていなければ、アクセスを得ることができる。
	var m sync.RWMutex
	fmt.Fprintf(tw, "Readers\tRWMutex\tMutex\n")
	for i := 0; i < 20; i++ {
		count := int(math.Pow(2, float64(i)))
		fmt.Fprintf(
			tw,
			"%d\t%v\t%v\n",
			count,
			test(count, &m, m.RLocker()), // 目的はRWMutexを使っていこうということらしいがこのコードとの関連が不明
			test(count, &m, &m),
		)
	}
}
