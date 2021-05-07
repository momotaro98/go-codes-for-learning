
```
go test -bench . -benchmem
```

result 2021-03-19 MacBook Pro

```
goos: darwin
goarch: amd64
Benchmark_string-4                769185             52775 ns/op          388557 B/op          1 allocs/op
Benchmark_stringsBuilder-4      204743859                5.09 ns/op            5 B/op          0 allocs/op
PASS
```
