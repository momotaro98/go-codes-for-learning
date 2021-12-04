package main

import (
	"fmt"
	"unsafe"
)

func run() {
	// 0-255の範囲では専用テーブルがある
	var (
		out int = 255
	)
	v255V1 := 255
	v256V1 := 256
	v255V2 := 255
	v256V2 := 256
	fmt.Println(&v255V1)
	fmt.Println(&v256V1)
	fmt.Println(&v255V2)
	fmt.Println(&v256V2)
	doThing(out, func(out interface{}) {
		val := 255
		fmt.Println(&val)
		fmt.Printf("val is 255, right? %d\n", val)    // Prints 999999999
		fmt.Printf("val is 255, right? %d\n", v255V1) // Prints 999999999
		fmt.Printf("val is 255, right? %d\n", v255V2) // Prints 999999999
	})
}

//go:noinline
func doThing(out interface{}, f func(out interface{})) {
	p := (*eface)(unsafe.Pointer(&out)).data
	*(*int)(p) = 999999999
	f(out)
}

type eface struct {
	rtype unsafe.Pointer
	data  unsafe.Pointer
}

func main() {
	run()
}
