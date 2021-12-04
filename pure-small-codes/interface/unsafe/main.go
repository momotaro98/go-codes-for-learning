package main

import (
	"fmt"
	"unsafe"
)

func run() {
	var out int = 255
	doThing(out, func(out interface{}) {
		val := 255
		fmt.Printf("val is 255, right? %d\n", val) // Prints 666
	})
}

//go:noinline
func doThing(out interface{}, f func(out interface{})) {
	p := (*eface)(unsafe.Pointer(&out)).data
	*(*int)(p) = 666
	f(out)
}

type eface struct {
	rtype unsafe.Pointer
	data  unsafe.Pointer
}

func main() {
	run()
}
