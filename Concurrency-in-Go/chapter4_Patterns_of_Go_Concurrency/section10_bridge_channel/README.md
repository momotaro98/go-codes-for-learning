# 4.10 bridge channel

状況によっては、チャネルのシーケンス(チャネルがチャネルに入る)から値を消費したいと思うことがあるでしょう。

``
<-chan <-chan interface{}
```

```go
bridge := func(done <-chan interface{}, chanStream <-chan <-chan interface{}) <-chan interface{} {
    valCh := make(chan interface{})
    go func() {
        defer close(valCh)
        for {
            var stream <-chan interface{}
            select {
            case maybeStream, ok := <-chanStream
                if ok == false {
                    return
                }
                stream = maybeStream
            case <-done:
                return
            }
            for val := range orDone(done, stream) {
                select {
                case valCh <- val:
                case <-done:
                }
            }
        }
    }()
    return valCh
}
```

これでbridgeを使って、チャネルのチャネルを単一のチャネルのように見せかけられます。次の例では10個のチャネルの列を作って、それぞれのチャネルに要素を1つだけ書き込み、それらのチャネルをbridge関数に渡します。

```go
genVals := func() <-chan <-chan interface{} {
    chanStream := make(chan (<-chan interface{}))
    go func() {
        defer close(chanStream)
        for i := 0; i < 10; i++ {
            stream := make(chan interface{}, 1)
            stream <- i
            close(stream)
            chanStream <- stream
        }
    }()
    return chanStream
}

for v := range bridge(nil, genVals()) {
    fmt.Printf("%v", v)
}
```

このコードを実行すると次のようになります。

0 1 2 3 4 5 6 7 8 9

bridgeのおかげで、チャネルのチャネルを1つのrangeを使った繰り返しで扱え、繰り返し処理の中のロジックに集中できます。チャネルのチャネルを崩すことで、その値を扱う問題のみをコードにすれば良くなります。
