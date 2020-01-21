package main

import (
	"fmt"
	"time"
)

func main() {
	doWork := func(done <-chan interface{}, strings <-chan string) <-chan interface{} {
		terminated := make(chan interface{})
		go func() {
			defer fmt.Println("doWork exited.")
			defer close(terminated)
			for {
				select {
				case s := <-strings:
					// Something interesting process
					fmt.Println(s)
				case <-done: // 子はdoneがされれば自身の処理を終えるようにする
					return
				}
			}
		}()
		return terminated
	}

	done := make(chan interface{})
	terminated := doWork(done, nil) // 親(Mainゴルーチン)から子に対してdoneチャネルを渡す

	go func() {
		time.Sleep(1)
		fmt.Println("Done")
		close(done) // 親の責任としてdoneをクローズにする
	}()

	<-terminated
	fmt.Println("End")
}
