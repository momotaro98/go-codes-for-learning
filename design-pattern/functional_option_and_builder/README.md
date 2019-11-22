# How to benchmark and investigate allocation stat

Ref link: https://methane.github.io/2015/02/reduce-allocation-in-go-code/

```
$ go clean; go build -o a.out; rm trace.log; GODEBUG=allocfreetrace=1 ./a.out -test.run=none -test.benchtime=1ms 2>trace.log
```

`go test -c` で`bench_test.go`でのバイナリを動かしてGODEBUG=allocfreetrace=1 してもなぜかテストコードの行が出てこない。
