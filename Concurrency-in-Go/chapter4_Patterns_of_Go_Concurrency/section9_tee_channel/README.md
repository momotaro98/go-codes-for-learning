# tee channel

```go
tee := func(done <-chan interface{}, in <-chan interface{}) (_, _ <-chan interface{}) {
	out1 := make(chan interface{})
	out2 := make(chan interface{})
	go func() {
		defer close(out1)
		defer close(out2)
		for val := range orDone(done, in) {
			var out1, out2 = out1, out2 // なぜこうするか不明
			for i := 0; i < 2; i++ {
				select {
				case out1 <- val:
					out1 = nil
				case out2 <- val:
					out2 = nil
				}
			}
		}
	}()
	return out1, out2
}
```

`in`に対する繰り返しの読み込みはout1とout2の書き込みが終わらない限り進みません。

このパターンを使うことで、チャネルをシステムの接続点として使い続けることが容易になります。
