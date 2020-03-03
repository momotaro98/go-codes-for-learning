package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

// 送信先(例:ディスク)よりも速いもの(例:メモリ)からの入力をバッファするステージです。
// これはもちろん、Goのbufioパッケージとしての目的そのものです。
// この例では、バッファありとなしのキューへの書き込みを比較しています。

// $ go test -bench=.
// BenchmarkUnbufferedWrite-4        133764              8914 ns/op
// BenchmarkBufferedWrite-4         1385637               868 ns/op

func BenchmarkUnbufferedWrite(b *testing.B) {
	performWrite(b, tmpFileOrFatal())
}

func BenchmarkBufferedWrite(b *testing.B) {
	bufferredFile := bufio.NewWriter(tmpFileOrFatal())
	performWrite(b, bufio.NewWriter(bufferredFile))
}

func tmpFileOrFatal() *os.File {
	file, err := ioutil.TempFile("", "tmp")
	if err != nil {
		fmt.Println(err)
	}
	return file
}

func performWrite(b *testing.B, writer io.Writer) {
	done := make(chan interface{})
	defer close(done)

	b.ResetTimer()
	for bt := range take(done, repeat(done, byte(0)), b.N) {
		writer.Write([]byte{bt.(byte)})
	}
}
