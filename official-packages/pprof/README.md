```
go run runtime.pprof.go
```

-----------------------------------------------

```
go run net.http.pprof.go
```

In another terminal, run below and wait 10 secs

```
go tool pprof http://localhost:6070/debug/pprof/profile?seconds=10
```
