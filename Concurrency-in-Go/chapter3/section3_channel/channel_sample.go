package main

// $ go run this_file.go
// Received: 0
// Received: 1
// Received: 2
// Received: 3
// Received: 4
// Received: 5
// Done receiving!

import "fmt"

func main() {
	// チャネル所有者側
	chanOwner := func() <-chan int {
		resultCh := make(chan int, 5) // バッファは1以上は持たない方が良いがこれは例としてそうしている。
		go func() {
			defer close(resultCh) // 責任範囲のチャネルを閉じることは所有者の義務
			for i := 0; i <= 5; i++ {
				resultCh <- i
			}
		}()
		return resultCh // Read only channel (<-chan型) として返る
	}

	// チャネル利用者側
	// 利用者側はチャネルが閉じられていることを確認すること
	// チャネル受け取りはブロックされる場合があることを把握して実装する必要がある。
	resultCh := chanOwner()
	for result := range resultCh {
		fmt.Printf("Received: %d\n", result)
	}
	fmt.Println("Done receiving!")
}
