package main

import "fmt"

func main() {
	chanOwner := func() <-chan int {
		results := make(chan int, 5) // This channel is initialized as a lexical(static) scope in chanOwner func.
		go func() {
			defer close(results)
			for i := 0; i < 5; i++ {
				results <- i
			}
		}()
		return results
	}

	consumer := func(results <-chan int) { // Take readonly channel. This is one of "binding"
		for result := range results {
			fmt.Printf("Received: %d\n", result)
		}
		fmt.Println("Done receiving!")
	}

	results := chanOwner() // Return readonly channel
	consumer(results)
}
