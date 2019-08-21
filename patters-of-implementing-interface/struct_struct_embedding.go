/// [ポイント2] 埋め込みのメリット
package main

type HighJumper interface {
	HighJump() string
}

// grasshopper はバッタのこと
// 高く飛ぶ能力がある
type grasshopper struct{}

func (g *grasshopper) HighJump() string {
	return "High Jump!"
}

// ShinJinrui2 は*grasshopperの能力を
// 構造体埋め込み(③)により"そのまま"借りる
type ShinJinrui2 struct { // 新人類2
	*grasshopper
}

//// 以下のように埋め込まない(そのまま借りない)場合、
//// ShinJinrui2としてのメソッドとしてHighJumpを作って
//// *grasshopperの者を呼ぶ必要がある
// type ShinJinrui2 struct { // 新人類2
// 	ghopper *grasshopper
// }
//// 埋め込みではこのメソッドをわざわざ実装しなくて良いメリットがある。
// func (sj *ShinJinrui2) HighJump() string {
// 	return sj.grasshopper.HighJump()
// }
