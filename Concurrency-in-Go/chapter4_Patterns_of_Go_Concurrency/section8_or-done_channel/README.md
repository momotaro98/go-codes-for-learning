# or-Done channel

本チャネルに対してdoneチャネルもコンシューマ側として扱うとき

```go
for val := range {
    // valに対して何かする
}
```

という読みやすかったコードが以下のように膨れ上がってしまう。

```go
loop:
for {
    select {
    case <-done:
        break loop
    case maybeVal, ok := <-myChan:
        if !ok {
            return
        }
        // valに対して何かする
    }
}
```

この冗長な状況をカプセル化して、他の人が触らない済むようにする。

```go
orDone := func(done, c <-chan interface{}) <-chan interface{} {
    valCh := make(chan interface{})
    go func() {
        defer close(valCh)
        for {
            select {
            case <-done:
                return
            case v, ok := <-c:
                if !ok {
                    return
                }
                select {
                case valCh <- v:
                case <-done:
                    return
                }
            }
        }
    }()
    return valCh
}
```

こうすることで単純なforループで書ける。

```go
for val := range orDone(done, myChan) {
    // valに対して何かする
}
```