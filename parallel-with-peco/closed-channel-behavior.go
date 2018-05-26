package main

import (
	"fmt"
)

func ExampleReadClosedChannel() {
	ch := make(chan int, 10)
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch) // close channel before receiving

	for i := 0; i < 10; i++ {
		v := <-ch // receive channel from 1 to 10
		fmt.Println("READ", v)
	}
	v, ok := <-ch // i=0. (zero value of int type is 0. And, ok is false)
	fmt.Println("READ", v, ":ok", ok)
}

func main() {
	ExampleReadClosedChannel()
}
