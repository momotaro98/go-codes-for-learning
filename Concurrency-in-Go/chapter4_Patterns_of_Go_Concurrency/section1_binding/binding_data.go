package main

import (
	"bytes"
	"fmt"
	"sync"
)

// $ go run this_file.go
// ang
// gol

func main() {
	printData := func(wg *sync.WaitGroup, data []byte) {
		// 引数にすることで`data`の一部にしかアクセスできないように"拘束"(binding)している。
		// > unsafeパッケージを使ってメモリを手で操作できる可能性を無視しました。
		// > unsafeと呼ばれるには理由があるのです！
		defer wg.Done()

		var buff bytes.Buffer
		for _, b := range data {
			fmt.Fprintf(&buff, "%c", b)
		}
		fmt.Println(buff.String())
	}

	var wg sync.WaitGroup
	wg.Add(2)
	data := []byte("golang")
	go printData(&wg, data[:3])
	go printData(&wg, data[3:])

	wg.Wait()
}
