package main

import (
	"context"
	"fmt"
)

// $ go run XXX.go
// handling response for jane (auth: abc123)

func main() {
	ProcessRequest( /* userID */ "jane" /* authToken */, "abc123")
}

// 【ポイント】
// i. Contextのキーのルール
// * 使用するキーはGoでの比較可能性を満たさなければならない。つまり、等価演算子の==と！=を使ったときに正しい値を返す必要がある。
// * 返された値は複数のゴルーチンからアクセスされても安全でなければならない。
// 型安全の観点から、Goではパッケージ内で独自のキーの型を定義することを推奨している。型さえ異なればContextのmap[interface{}]内で衝突することはない。
// しかし、定義した型はどのサブパッケージからも参照されるように工夫しないといけない。
//
// ii. ctxUser, ctxAuthToken 2つのContextのキーはプライベートなため、データを取得する関数をエクスポートする必要がある。
// これは良い方法で、このデータを扱う側が静的で型安全な関数を使えるようになる。

type ctxKey int // i.

const (
	ctxUserID ctxKey = iota
	ctxAuthToken
)

func UserID(c context.Context) string { // ii.
	return c.Value(ctxUserID).(string)
}

func AuthToken(c context.Context) string { // ii.
	return c.Value(ctxAuthToken).(string)
}

func ProcessRequest(userID, authToken string) {
	ctx := context.WithValue(context.Background(), ctxUserID, userID)
	ctx = context.WithValue(ctx, ctxAuthToken, authToken)
	HandleResponse(ctx)
}

func HandleResponse(ctx context.Context) {
	fmt.Printf(
		"handling response for %v (auth: %v)",
		UserID(ctx),
		AuthToken(ctx),
	)
}
