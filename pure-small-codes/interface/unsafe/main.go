package main

import (
	"fmt"
	"unsafe"
)

func run() {
	doThing(func(out interface{}) {
		val := 255
		if val == 255 {
			fmt.Printf("val is 255, right? %d\n", val) // Prints 666
		}
	})
}

func doThing(f func(out interface{})) {
	var i int = 255
	var out interface{}
	out = i
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
