/// [ポイント1] 埋め込みの関係
package main

type Flyer interface {
	Fly() string
}

type Runner interface {
	Run() string
}

// ①  インターフェースにインターフェースを埋め込む
type FlyingRunner interface {
	Flyer
	Runner
}

// ② 構造体にインターフェースを埋め込む
type ToriJin struct { // 鳥人
	FlyingRunner
}

// ③ 構造体に構造体を埋め込む
type ShinJinrui struct { // 新人類
	*ToriJin
}

type RealToriJin struct{}

func (r RealToriJin) Fly() string { return "Fly!" }

func (r RealToriJin) Run() string { return "Run!" }
