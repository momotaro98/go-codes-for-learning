# Book

https://www.oreilly.co.jp/books/9784873118468/

# Contents

## 1.2.5 並行処理の安全性を見極める

```go
// CalculatePi は円周率のbegin桁目からend桁目の数字を計算します。
func CalculatePi(begin, end int64, pi *Pi)
```

この関数の例では多くの疑問(以下)が湧くことになる

* この関数を使ってどうやってπ計算ができるのか。
* この関数を複数並行起動するところも自分でやらないといけないのか。
* この関数は、自分でアドレスを渡しているPiのインスタンスを直接操作しているように見える。このPiのメモリアクセスは自分で行う必要があるのか、それともPi型が管理してくれるのか。

__コメントを書くことで利用者の疑問を晴らしてあげよう__

```go
// CalculatePi は円周率のbegin桁目からend桁目の数字を計算します。

// 内部的には、CalculatePi は FLOOR((end-begin)/2) 個の並行プロセスを立ち上げて
// 再帰的に CalculatePi を呼び出します。piへの書き込みの同期はPi構造体の内部で処理されます。
func CalculatePi(begin, end int64, pi *Pi)
```

並行処理・同期処理プログラムにおいて重要なこと

* 誰が並行処理を担っているか。
* 問題空間がどのように並行処理のプリミティブに対応しているか。
* 誰が同期処理を担っているか。

Goでは以下のようにするとコメント無しでも上記疑問を解決できる。

```go
func CalculatePi(begin, end int64) <-chan uint
```
