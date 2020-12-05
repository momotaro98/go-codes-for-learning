package main

import "C"

//export Gophernize
func Gophernize(name string) *C.char {
	str := name + " ʕ ◔ϖ◔ʔ"
	return C.CString(str)
}

func main() {}
