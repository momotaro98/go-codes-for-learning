## チャネルから繰り返しの変数を送出する

```go
for _, s := range []string{"a", "b", "c"} }
    select {
    case <-done:
        return
    case stringCh <- s:
    }
}
```

## 停止シグナルを待つ無限ループ

```go
for {
    select {
    case <-done:
        return
    default:
    }

    // 割り込みできない処理をする
}
```
