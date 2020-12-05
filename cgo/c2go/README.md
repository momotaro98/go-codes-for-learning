
```
//extern goHello
func goHello() {
.
.
```

 ↓ go build すると

 ↓ GCO付きのGoコンパイラが中間生成物を作成

|----------------
| 中間生成物
| `_cgo_export.h` `_cgo_export.c`
|----------------

 ↓ 実行する

 hello.c が中間生成物を通してGoのコードを参照できる

 ↓

 自動生成されたコードの中では、crosscall2を経てGoに処理が返り、runtime.cgocallback関数を経て目的のGo関数が呼ばれる。
