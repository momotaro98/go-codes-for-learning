/// ポイント4: 埋め込み元と埋め込み先のフィールド名重複時の挙動
package main

import "log"

// dolphin はイルカ
// 水中に潜る能力がある
type dolphin struct{}

func (g *dolphin) Dive() string {
	return "Dolpin Dive!"
}

// ShinJinrui4 はdolphinの能力を
// 構造体埋め込み(③)により"そのまま"借りる
type ShinJinrui4 struct {
	*dolphin
}

// ShinJinrui4は自分の能力だけでもDiveできる
func (sj *ShinJinrui4) Dive() string {
	return "ShinJinrui Dive!"
}

// ShinJinrui5 はdolphinの能力を借りるし
// dolphinという名前のフィールドを持つ
type ShinJinrui5 struct {
	*log.Logger
}

/// effective_goの資料には外部からアクセスされなければ重複を許容とあるが
/// 実際にこのメソッドを定義すると
/// type ShinJinrui5 has both field and method named Loggerの
/// コンパイルエラーがビルド時に発生する。
// func (sj *ShinJinrui5) Logger() string {
// 	return "A method name, dolphin of ShinJinru5"
// }
