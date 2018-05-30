package main

import (
	"context"
	"log"
	"sync"
	"time"
)

func ExampleContextRelationship() {
	parentCtx, parentCancel := context.WithCancel(context.Background())
	childCtx, childCancle := context.WithCancel(parentCtx)
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		log.Printf("a process waiting for cancel of parent ctx")
		<-parentCtx.Done()
		log.Printf("parent, parentCtx.Done() returns")
	}()

	go func() {
		defer wg.Done()
		log.Printf("a process waiting for cancel of child ctx")
		<-childCtx.Done()
		log.Printf("child, childCtx.Done() returns")
	}()

	// Both parent and child will be canceled
	// time.AfterFunc(time.Second*5, childCancle)
	// time.AfterFunc(time.Second*1, parentCancel)

	// Only child will be canceled
	time.AfterFunc(time.Second*5, parentCancel)
	time.AfterFunc(time.Second*1, childCancle)

	wg.Wait()
}

func main() {
	ExampleContextRelationship()
}
