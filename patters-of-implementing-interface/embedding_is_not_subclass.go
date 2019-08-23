/// [ポイント3] 構造体の埋め込みは継承のサブクラス化とは異なる
package main

import "fmt"

// Status は健康状態を意味する
type Status int

const (
	// Good is 良好 status
	Good Status = iota
	// Tired is 疲れている status
	Tired
)

func (s Status) String() string {
	switch s {
	case Good:
		return "Good!"
	case Tired:
		return "Tired..."
	default:
		return ""
	}
}

type poorGrasshopper struct {
	status Status // poorGrasshopperには健康状態がある
}

func (g *poorGrasshopper) HighJump() {
	fmt.Println("High Jump!")
	g.status = Tired // 飛ぶと疲れてしまう
}

type ShinJinrui3 struct { // 新人類3
	status           Status // ShinJinrui3も健康状態がある
	*poorGrasshopper        // 構造体の埋め込み(③)
}
