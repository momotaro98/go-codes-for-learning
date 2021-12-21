# repo

https://github.com/uber-go/goleak

# how to run

```
$ go test
--- FAIL: TestA (0.44s)
    leaks.go:78: found unexpected goroutines:
        [Goroutine 19 in state chan receive (nil chan), with 3rd-parties/uber-go/goleak.LeakFunc.func1 on top of the stack:
        goroutine 19 [chan receive (nil chan)]:
        3rd-parties/uber-go/goleak.LeakFunc.func1()
                /Users/shintaro/workspace/github.com/momotaro98/go-codes-for-learning/3rd-parties/uber-go/goleak/main.go:9 +0x29
        created by 3rd-parties/uber-go/goleak.LeakFunc
                /Users/shintaro/workspace/github.com/momotaro98/go-codes-for-learning/3rd-parties/uber-go/goleak/main.go:8 +0x6f
        ]
FAIL
exit status 1
FAIL    3rd-parties/uber-go/goleak      0.788s
```
