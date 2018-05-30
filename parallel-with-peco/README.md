# 特集2 Goによる並行処理 〜複雑な処理をスイスイ書こう！〜

## 第2章 Goでの並行処理

### goroutineの特徴

> goroutineが実行される裏ではGoのランタイムが起動したいくつかのスレッドがあり、空いているスレッドに順にgoroutineを乗せて処理を進める、タイムシェア方式とでも呼べる実装になっています。

### スレッドとの違い

#### channelによるデータの送受信

> channelを受け取ったgoroutineが現在の「持ち主」であり、そのgoroutineだけがchannelに対して操作を行える
> この実装の場合、明示的な排他制御は必要なくなります。

#### IDを持たない

> goroutineはIDを持ちません。固有のgoroutineを、スレッドやプロセスのようにそのIDだけでは判別できません。

> IDが指定できないため、シグナルのようなしくみで外部からgoroutineを指定して、その動作に強制的に影響を及ぼすことはできません。そのためgoroutineには、ほぼ必ず明示的に停止するためのしくみを用意する必要があります。

## channelによるデータ受け渡し

### channelの基本

[code](https://github.com/momotaro98/go-codes-for-learning/blob/master/parallel-with-peco/channel-basic.go)

### ブロックせずにchannel処理を行う

> 次の例では、3つのchannelの中から読み込み可能なchannelがあればそれを読み込みます。すべてのchannelがデータ到着待ちでブロックするのであれば、defaultのケースが実行されます。

```
for {
    select {
    case value := <-ch1:
        ...
    case value := <-ch2:
        ...
    case value := <-ch3:
        ...
    default:
        fmt.Println("nothing to do!")
    }
}
```

> なお、nilなchannelに対して読み込みを行うと必ずブロックします。selectと合わせて、特定のselectのケースを無効にしたい場合などに便利です。

```
for {
    select {
    case <-ch1:
      ...
    case <-ch2:
      // ここでch1をnilにすると、上記のcaseはブロックされ実行されない
      ch1 = nil
      ...
    }
}
```

### 閉じたchannelの挙動

> channelに対する**書き込み**が必要なくなった時点で、close関数を使ってchannelを閉じることができます。閉じられたchannelへの書き込みを行うと例外が起こりますので注意してください。

> channelを閉じると、読み込み待ちでブロックしていたすべてのgoroutineのブロック状態が終わります。この性質を使えば、goroutineが終了したことを伝えたり、複数のgoroutineを一斉に通知したりできます。

```
ch := make(chan struct{})
for i := 0; i < 10; i++ {
    go func() {
        <-ch // waiting
        ...
    }()
}

time.AfterFunc(5*time.Second, func() {
    close(ch) // above blocking(<-ch) will be released and process goes next
})
```

> また、閉じられたchannelから読み込みを行うと、それまですでに書き込まれた値が通常どおり返ってきます。ほかに読み込む値がない状態でさらに読み込みを行うと、channelが返すべき型のゼロ値が即座に返ってきます。

[code](https://github.com/momotaro98/go-codes-for-learning/blob/master/parallel-with-peco/closed-channel-behavior.go)

### Condで複数の相手に状態の変更を通知する

> 状態が変わったことを通知するコンディション変数は、Goではsync.Condを使って実装できます。

[code](https://github.com/momotaro98/go-codes-for-learning/blob/master/parallel-with-peco/cond-example.go)

Description of `Cond` struct type

```
type Cond struct {

        // L is held while observing or changing the condition
        L Locker

        // Has unexported fields.
}
    Cond implements a condition variable, a rendezvous point for goroutines
    waiting for or announcing the occurrence of an event.

    Each Cond has an associated Locker L (often a *Mutex or *RWMutex), which
    must be held when changing the condition and when calling the Wait method.

    A Cond must not be copied after first use.


func NewCond(l Locker) *Cond
func (c *Cond) Broadcast()
func (c *Cond) Signal()
func (c *Cond) Wait()
```

## 第3章 並行処理の実装パターン

### セマフォで同時実行数の制御

> Mutexは上限数が1のセマフォと考えることもできます。
> Goにはセマフォとう名前のデータ型は存在しませんが、セマフォに相当するものとして、バッファ付きchannelを利用します。
> 次の例では、バッファ付きchannelの特性を使って同時にhttp.Getを実行可能なgoroutineの数を5個に制限しています。

```
func FetchURL(sem chan struct {}, url string) {
    sem<-struct{}{} // Block when over 10 requests come
    defer func() { <-sem }() // Release sem when http.Get request process done

    res, err := http.Get(url)
    ...
}

func ExampleSemaphore() {
    sem := make(chan struct{}, 10)
    urls := []string{ ... } // URL list
    var wg sync.WaitGroup
    for _, u := range urls {
        wg.Add(1)
        go func() {
            defer wg.Done()
            FetchURL(sem, u) // request to URL with new goroutine
        }()
    }
    wg.Wait()
}
```

### ワーカにタスクをfan-outさせる


```
func ExampleFanout() {
    ch := make(chan FanoutTask) // 0 buffer channel

    for i := 0; i < 10; i++ {
        go FanoutWorker(ch) // channel Reciever
    }

    FanoutDispatcher(ch) // channel Sender
}

// channel Reciever
func FanoutWorker(in chan FanoutTask) {
    for {
        task, ok := <-in
        if !ok {
            return
        }
        ... // process task
    }
}

// channel Sender
func FanoutDispatcher(out chan FanoutTask) {
    defer close(out)
    // Get task from Data Storage like DB
    for {
        task, err := FanoutGetNextTask()
        if err != nil {
            return
        }
        // send to channel
        out<-task
    }
}
```

### ジェネレータで連番の生成

> プログラム内で利用するIDに連番を生成するケースは頻繁にあります。他の言語では、同時に複数の呼び出し元がある場合に連番が正しく生成されていることを保証するには明示的な排他制御が必要ですが、Goであれば安全に連番を生成できます。

[code](https://github.com/momotaro98/go-codes-for-learning/blob/master/parallel-with-peco/generate-series.go)

> この例のポイントは、channelへの書き込み・読み込みは複数goroutine間で安全なため、一切の明示的な排他制御をせずに正しく連番を扱えることです。

### time.Timerでタイムアウト処理

> N秒後に何らかの処理を行いたい場合は、time.Timerオブジェクトを作成します。

[code](https://github.com/momotaro98/go-codes-for-learning/blob/master/parallel-with-peco/timeTimer.go)

### time.Tickerで定期的な処理

> N秒ごとに定期的に処理を行いたい場合もあります。

[code](https://github.com/momotaro98/go-codes-for-learning/blob/master/parallel-with-peco/timeTicker.go)

### context.Contextでキャンセル処理

> 複数のgoroutineが絡んだ処理を実装すると、何らかの条件で関連しているgoroutineにキャンセル通知をしたい場面が出てきます。またその際、一部のgoroutine郡はキャンセルし残りはそのまま処理を続けたいという場合もあります。Go1.7から標準パッケージに同梱されるようになったcontext.Contextを使うと、この処理を簡潔に記述できます。

#### context.Contextの基本

```
func WorkWithContext(ctx context.Context) {
    defer close(done)
    for {
        select {
        case <-ctx.Done():
            return
        default:
            // Do normal process if 'Done' isn't returned
            ...
        }
    }
}

func ExampleContext() {
    // Create cancelable `Context`
    ctx, cancel := context.WithCancel(context.Background())
    // This function is sure to cancel
    defer cancel()

    // goroutine which does process
    go WorkWithContext(ctx)
    ...
}
```

> このようにすると、明示的にcancelを呼ぶか、ExampleContext()関数が終了した時点で暗黙的に呼ばれるcancelによって、WorkingWithContext関数内で待ち受けている<-ctx.Done()に通知が送られ、正しくgoroutineを終了できます。

#### キャンセル効果の範囲

> 親をキャンセルするとすべての子もキャンセルされます。子だけをキャンセルした場合は親には影響がありません。つまり、アプリケーション全体のキャンセルも、局所的なキャンセルも可能になります。

> 次のコードでは、ctx1(親)とctx2(子)のcontext.Contextを作成しています。それぞれのキャンセルを待つgoroutineを作成したのち、5秒後と1秒後にそれぞれのキャンセル処理を実行します。

[code](https://github.com/momotaro98/go-codes-for-learning/blob/master/parallel-with-peco/context-cancel-range.go)

> この特性を使うと、context.Contextを実行される関数に渡していき、必要なところで子context.Contextを作ることによって、「シグナルでアプリケーション全体をキャンセルする」という動作や「特定の処理に紐付いているgoroutineのみをキャンセルする」などの処理を簡単に実装できます。