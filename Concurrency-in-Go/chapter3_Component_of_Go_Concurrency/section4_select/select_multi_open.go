package main

import "fmt"

// $ go run this_file.go
// c1Count: 487
// c2Count: 514

func main() {
	c1 := make(chan interface{})
	close(c1)
	c2 := make(chan interface{})
	close(c2)

	var c1Count, c2Count int
	for i := 1000; i >= 0; i-- {
		select { // c1, c2 共にOpenな場合、"ランダム"でどちらかが選ばれる
		case <-c1:
			c1Count++
		case <-c2:
			c2Count++
		}
	}

	fmt.Printf("c1Count: %d\nc2Count: %d", c1Count, c2Count)
}
