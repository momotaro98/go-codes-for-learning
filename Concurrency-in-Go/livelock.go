package main

// ライブロックとは並行操作を行っているけれど、
// その操作はプログラムの状態をまったく進めて
// いないプログラムを指す。

/* Output
Barbara is trying to scoot: left right left right left right left right left right
Barbara tosses her hands up in exasperation!
Alice is trying to scoot: left right left right left right left right left right
Alice tosses her hands up in exasperation!
*/

import (
	"bytes"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func LiveLockDemo() {
	cadence := sync.NewCond(&sync.Mutex{}) // cadence 歩調 リズム
	// cadence に対する処理は人間が同じ歩調を歩くのをシュミレートするためだけのもの。
	go func() {
		for range time.Tick(1*time.Millisecond) {
			cadence.Broadcast()
		}
	}()
	takeStep := func() {
		cadence.L.Lock()
		cadence.Wait()
		cadence.L.Unlock()
	}

	tryDir := func(dirName string, dir *int32, out *bytes.Buffer) bool {
		fmt.Fprintf(out, " %v", dirName)
		atomic.AddInt32(dir, 1)
		takeStep()
		if atomic.LoadInt32(dir) == 1 {
			fmt.Fprint(out, ". 成功!")
			return true
		}
		takeStep()
		atomic.AddInt32(dir, -1)
		return false
	}

	var left, right int32
	tryLeft := func(out *bytes.Buffer) bool { return tryDir("left", &left, out) }
	tryRight := func(out *bytes.Buffer) bool { return tryDir("right", &right, out) }

	walk := func(walking *sync.WaitGroup, name string) {
		var out bytes.Buffer
		defer func() { fmt.Println(out.String()) }()
		defer walking.Done()
		fmt.Fprintf(&out, "%v is trying to scoot:", name) // scoot [自動詞]急いで行く
		for i := 0; i < 5; i++ {
			if tryLeft(&out) || tryRight(&out) {
				return
			}
		}
		fmt.Fprintf(&out, "\n%v tosses her hands up in exasperation!", name) // %vさんは腹を立ててお手上げだ
	}

	var peopleInHallway sync.WaitGroup
	peopleInHallway.Add(2)
	go walk(&peopleInHallway, "Alice")
	go walk(&peopleInHallway, "Barbara")
	peopleInHallway.Wait()
}
