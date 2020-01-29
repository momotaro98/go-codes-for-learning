package main

var toString = func(done <-chan interface{}, valueCh <-chan interface{}) <-chan string {
	stringCh := make(chan string)
	go func() {
		defer close(stringCh)
		for v := range valueCh {
			select {
			case <-done:
				return
			case stringCh <- v.(string):
			}
		}
	}()
	return stringCh
}
