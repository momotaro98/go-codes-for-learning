## Shift operation シフト演算

```go
package main

import (
	"fmt"
)

// Souce link: https://golangtokyo.connpass.com/event/150891/

const (
	a, b = iota, iota << iota
	c, _ // (1 << 1 --> b'10 --> 2)
	_, d // (2 << 2 --> b'1000 --> 8)
	e, f // (3 << 3 --> b'11000 --> 24)
)

func main() {
	fmt.Println("a", a)
	fmt.Println("b", b)
	fmt.Println("c", c)
	fmt.Println("d", d)
	fmt.Println("e", e)
	fmt.Println("f", f)
	/*
		a 0
		b 0
		c 1
		d 8
		e 3
		f 24
	*/
}
```

# 制御構造

## For statement

```go
package main
import "fmt"
// aを逆に並び替える
// ① i++, j-- というのは"文"なので複数変数をfor文で扱うときはこのようにする
func main() {
	a := [5]int{1, 2, 3, 4, 5}
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 { // ①
		a[i], a[j] = a[j], a[i]
	}
	fmt.Println(a) // [5 4 3 2 1]
}
```

## Switch statement

```go
package main
import "fmt"
func retInterface() interface{} {
	v := 100
	return &v
}
func main() {
	interfaceValue := retInterface()
	switch t := interfaceValue.(type) {
	default:
		fmt.Printf("unexpected type %T", t) // %T は型を出力する
	case bool:
		fmt.Printf("boolean %t\n", t)
	case int:
		fmt.Printf("integer %d\n", t)
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *t)
	case *int:
		fmt.Printf("pointer to integer %d\n", *t) // pointer to integer 100
	}
}
```

# 関数

## defer statement

```go
package main

import (
	"fmt"
)

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer un(trace("a"))
	fmt.Println("in a")
}

func b() {
	defer un(trace("b")) // trace("b")はdefer文の時点で評価されて実行される
	fmt.Println("in b")
	a()
}

func main() {
	b()
}

/*
entering: b
in b
entering: a
in a
leaving: a
leaving: b
*/
```

# データ

## new function

> Go言語風に言い換えると、new(T)が返す値は、新しく割り当てられた型Tのゼロ値のポインタです。

```go
var p *[]int = new([]int)       // スライス構造の割り当て(*p == nil)。あまり使わない。
var v  []int = make([]int, 100) // スライスvは100個のintを持つ配列への参照
```

## make function

> 覚えておいていただきたいことは、makeが適用可能なのはマップ、スライス、チャネルだけであり、返される値はポインタではないことです。ポインタが必要であればnewで割り当ててください。

## 配列 Array

```go
array := [...]float64{7.0, 8.5, 9.1} // Arrayの初期化方法の1つ
```

* 配列は値である(いわゆるプリミティブ型で常にスタック領域)。関数に渡せばそのコピーが作られる。
* 配列のサイズは型の一部。`[10]int` and `[20]int` are different type.

## Slice

```go
package main

import "fmt"

func main() {
	a := make([]int, 5)
	fmt.Printf("Pointer &a is %p, len(a) is %d, cap(a) is %d\n", &a, len(a), cap(a))
	b := a[:3]
	fmt.Printf("Pointer &b is %p, len(b) is %d, cap(b) is %d\n", &b, len(b), cap(b))
	c := a[3:]
	fmt.Printf("Pointer &c is %p, len(c) is %d, cap(c) is %d\n", &c, len(c), cap(c))
}

/*
Pointer &a is 0x40a0e0, len(a) is 5, cap(a) is 5
Pointer &b is 0x40a0f0, len(b) is 3, cap(b) is 5
Pointer &c is 0x40a100, len(c) is 2, cap(c) is 2
*/
```

Slice has information of first pointer address, length, and capacity.

## Map

```go
package main

import "fmt"

var timeZone = map[string]int{
	"UTC": 0 * 60 * 60,
	"EST": -5 * 60 * 60,
	"CST": -6 * 60 * 60,
	"MST": -7 * 60 * 60,
	"PST": -8 * 60 * 60,
}

func offset(tz string) int {
	seconds, ok := timeZone[tz] // [Point] Check if there's the key
	if ok {
		fmt.Println("Found the tz, seconds is", seconds)
		return seconds
	}
	fmt.Println("zero value seconds", seconds)
	return 0
}

func main() {
	_ = offset("MST") // Found the tz, seconds is -25200
	_ = offset("JST") // zero value seconds 0
}
```
