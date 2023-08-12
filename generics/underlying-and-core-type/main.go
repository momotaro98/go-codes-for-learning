package main

import (
	"fmt"
)

// [note]
// ~ でunderlying typeをサポートしていない unions の場合、
// NeoNeoIntが対象ではないので t := Max(nni1, nni2) で エラーになってしまう。
// [error text] NeoNeoInt does not satisfy Number (possibly missing ~ for int in Number)
// type Number interface {
// 	int | int32 | int64 | float32 | float64
// }

// [note]
// ~ でunderlying typeをサポートした場合、NeoNeoIntのunderlying typeはintになりコンパイルが通るようになる。
type Number interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64
}

func Max[T Number](x, y T) T {
	if x >= y {
		return x
	}
	return y
}

type NeoInt int // underlying typeはint

type NeoNeoInt NeoInt // NeoNeoIntのunderlying typeもint

// [note]
// intのunderlying typeはint
// []intのunderlying typeは[]int

func main() {
	nni1 := NeoNeoInt(1)
	nni2 := NeoNeoInt(2)
	t := Max(nni1, nni2)
	fmt.Println(t)
}

// [note]
// Q. C1はcore typeは持つか？
// → A. Yes. C1を実装する全ての型のunderlying typeは[]intなのでC1はcore typeをもち、core typeは[]intです。
type C1 interface {
	~[]int
}

// [note]
// Q. C2はcore typeは持つか？
// → A. No. C2を実装する型はint, stringなのでunderlying typeは同一でなくC2はcore typeをもちません。
type C2 interface {
	int | string
}
