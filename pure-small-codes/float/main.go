package main

import (
	"fmt"
)

func main() {
	arr := [6]int32{-4, 3, -9, 0, 4, 1}
	plusMinus(arr[:])
}

const format = "%.6f\n" // 小数点第6位で四捨五入(round)

func plusMinus(arr []int32) {
	var posNum, negNum, zeroNum float64
	for i := range arr {
		if arr[i] == 0 {
			zeroNum++
		} else if arr[i] < 0 {
			negNum++
		} else {
			posNum++
		}
	}
	l := float64(len(arr))
	fmt.Printf(format, posNum/l) // Goではint同士での割り算(例 1/6) ではFloatにはならない。必ずFloat型にする。
	fmt.Printf(format, negNum/l)
	fmt.Printf(format, zeroNum/l)
}

// リテラル的にどの型にも当てはまる場合は明確な方が優先されるっぽい
// 参考 → https://stackoverflow.com/questions/32815400/how-to-perform-division-in-go
