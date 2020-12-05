package main

/*
#cgo LDFLAGS: -L./lib -lrustaceanize
#include <stdlib.h>
#include "./lib/rustaceanize.h"
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func main() {
	s := "I'm a Gopher"

	input := C.CString(s)
	defer C.free(unsafe.Pointer(input))

	// 以下の場合はinputのメモリはGoの管理化である。
	// このときGo側でGCが働くのでこのプログラムではランタイムエラーが発生する!
	// data := (*reflect.StringHeader)(unsafe.Pointer(&s)).Data
	// input := (*C.char)(unsafe.Pointer(data))

	o := C.rustaceanize(input)

	output := C.GoString(o)
	fmt.Printf("%s\n", output)
}
