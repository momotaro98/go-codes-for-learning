package main

import (
	"fmt"
)

// [note]
// これ↓は "|" で区切ったunionsを持つインターフェースであり、
// Go1.21の段階では型パラメータの制約にしか使うことができない。
type Stream interface {
	[]byte | string
}

// [note]
// var z Stream // ← Stream型の値を作ろうとするとコンパイルエラーになる
// [error text] cannot use type Stream outside a type constraint: interface contains type constraints

func main() {
	x, _ := atoi("24") // [note] "24"というstring型リテラルによって atoi[string]("24") と書かなくて良い (型推論)
	fmt.Println(x)
	y, _ := atoi([]byte("24"))
	fmt.Println(y)
}

// [note]
// 型パラメータ文法における [T C] の関係は、Tは型パラメータTであり、対してCはTに対するインターフェースで定義された制約である。
func atoi[bytes Stream](s bytes) (x int, err error) {
	neg := false
	if len(s) > 0 && (s[0] == '-' || s[0] == '+') {
		neg = s[0] == '-'
		s = s[1:]
	}
	q, rem, err := leadingInt(s)
	x = int(q)
	if err != nil || len(rem) > 0 {
		return 0, err
	}
	if neg {
		x = -x
	}
	return x, nil
}

func leadingInt[bytes Stream](s bytes) (x uint64, rem bytes, err error) {
	i := 0
	for ; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			break
		}
		if x > 1<<63/10 {
			// overflow
			return 0, rem, err
		}
		x = x*10 + uint64(c) - '0'
		if x > 1<<63 {
			// overflow
			return 0, rem, err
		}
	}
	return x, s[i:], nil
}

// Go1.21時点で構造体のメソッドには型パラメータを持たせることができない、の確認のための構造体
type Builder struct {
	addr *Builder // of receiver, to detect copies by value
	buf  []byte
}

// [note]
// Go1.21時点で構造体のメソッドには型パラメータを持たせることができない。
// [error text]: syntax error: method must have no type parameters
// func (b *Builder) Write[bytes []byte | string](p bytes) (int, error) {
// 	b.copyCheck()
// 	b.buf = append(b.buf, p...)
// 	return len(p), nil
// }
