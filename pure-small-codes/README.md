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
