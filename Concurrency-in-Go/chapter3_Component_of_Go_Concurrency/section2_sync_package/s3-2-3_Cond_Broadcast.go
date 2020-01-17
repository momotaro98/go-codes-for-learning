package main

import (
	"fmt"
	"sync"
)

// sync.Cond の Broadcastのユースケース
// [ボタンがあるGUIアプリ]
// ボタンがクリックされたときに実行される関数を
// 任意の数だけ登録したい。

// [Output]
// Mouse clicked.
// Maximizing window.
// Displaying dialog box.

func main() {
	type Button struct {
		Clicked *sync.Cond
	}

	button := Button{Clicked: sync.NewCond(&sync.Mutex{})}

	subscribe := func(c *sync.Cond, fn func()) {
		// [My note 2020-01-07] I thought it's unecessary to have
		// goroutineRunning as sync.WaitGroup then I tried
		// to remove but I got the dead lock runtime error as following
		// > fatal error: all goroutines are asleep - deadlock!
		// I'm not sure what and why happens
		var goroutineRunning sync.WaitGroup
		goroutineRunning.Add(1)
		go func() {
			goroutineRunning.Done()
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait()
			fn()
		}()
		goroutineRunning.Wait()
	}

	var clickRegistered sync.WaitGroup
	clickRegistered.Add(3)
	// Display
	subscribe(button.Clicked, func() {
		fmt.Println("Maximizing window.") // task
		clickRegistered.Done()
	})
	// Dialog box
	subscribe(button.Clicked, func() {
		fmt.Println("Displaying dialog box.") // task
		clickRegistered.Done()
	})
	// Mouse
	subscribe(button.Clicked, func() {
		fmt.Println("Mouse clicked.") // task
		clickRegistered.Done()
	})

	// Broadcast to subscribed items
	button.Clicked.Broadcast()

	clickRegistered.Wait()
}
