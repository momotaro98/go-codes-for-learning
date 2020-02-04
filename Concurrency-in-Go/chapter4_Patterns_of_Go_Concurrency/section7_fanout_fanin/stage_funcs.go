package main

import "math"

func repeatFn(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		defer close(ch)
		for {
			select {
			case <-done:
				return
			case ch <- fn():
			}
		}
	}()
	return ch
}

func toInt(done, valueCh <-chan interface{}) <-chan int {
	intCh := make(chan int)
	go func() {
		defer close(intCh)
		for v := range valueCh {
			select {
			case <-done:
				return
			case intCh <- v.(int):
			}
		}
	}()
	return intCh
}

func primeClassifier(done <-chan interface{}, intCh <-chan int) <-chan int {
	primeCh := make(chan int)

	isPrime := func(num int) bool {
		if num < 2 {
			return false
		} else if num == 2 {
			return true
		} else if num%2 == 0 {
			return false
		}
		var sqrt float64 = math.Sqrt(float64(num))
		for i := 3; float64(i) < sqrt; i += 2 {
			if num%i == 0 {
				return false
			}
		}
		return true
	}

	go func() {
		defer close(primeCh)
		for {
			select {
			case <-done:
				return
			case n := <-intCh:
				if isPrime(n) {
					primeCh <- n
				}
			}
		}
	}()
	return primeCh
}

var take = func(done <-chan interface{}, intCh <-chan int, num int) <-chan interface{} {
	takeCh := make(chan interface{})
	go func() {
		defer close(takeCh)
		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case takeCh <- <-intCh:
			}
		}
	}()
	return takeCh
}
