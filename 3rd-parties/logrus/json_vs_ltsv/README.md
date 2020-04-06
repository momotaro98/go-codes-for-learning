# logrus での LTSV と JSON 形式 の出力結果のサイズの差を比較

```
$ touch ltsv.log; touch json.log
$ go run main.go
$ ls -sla | grep log
122000  4  6 13:57 json.log
104000  4  6 13:57 ltsv.log
```

LTSV の方が 約 15% 分サイズを小さくできることがわかった。

# logrus での LTSV と JSON 形式 の出力時のパフォーマンスの差を比較

```
$ go test -bench . -benchmem
```

```
BenchmarkLTSV-4   	  405763	      2841 ns/op	    1368 B/op	      21 allocs/op
BenchmarkJSON-4   	  323056	      3748 ns/op	    1672 B/op	      28 allocs/op
```

LTSV の方が速度とメモリ使用においても優位であるとわかった。

1.3 倍速く、 20% メモリ消費が少ない
