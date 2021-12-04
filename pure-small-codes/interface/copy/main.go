package main

import (
	"fmt"
)

func f() int {
	var i int = 1
	var a interface{}
	a = &i // 値コピーされない
	i++
	return *a.(*int) // return 2
}

func g() int {
	var i int = 1
	var a interface{}
	a = i // 値コピーされる
	i++
	return a.(int) // return 1
}

func main() {
	fmt.Printf("a is %+v\n", f())
	fmt.Printf("a is %+v\n", g())
}
