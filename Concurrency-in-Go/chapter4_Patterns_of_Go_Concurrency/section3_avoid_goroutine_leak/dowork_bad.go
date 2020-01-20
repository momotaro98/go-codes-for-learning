package main

import "fmt"

func main() {
	doWork := func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(completed)
			for s := range strings { // doWork(nil)のように渡されているのでずっと処理が終わらずゴルーチンはずっとメモリ内に残される
				// Do something interesting process
				fmt.Println(s)
			}
		}()
		return completed
	}

	doWork(nil)
	// Do anything here
	fmt.Println("Done.")
}
