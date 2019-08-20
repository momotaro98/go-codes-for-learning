//// 元ネタ https://qiita.com/tenntenn/items/eac962a49c56b2b15ee8
package main

import (
	"fmt"
	"reflect"
)

func main() {
	/// [Topic1] 関数にインタフェース実装させる
	/// from func_with_interface.go
	// myFuncはHWriterとHRequestを引数に取る関数
	// この時点ではただの関数型の変数
	myFunc := func(w HWriter, r *HRequest) {
		fmt.Fprintf(w, "Hello world!")
	}
	fmt.Println(reflect.ValueOf(myFunc).Type()) // func(main.HWriter, *main.HRequest)
	// ただの関数型のmyFuncをHHandlerインターフェースを
	// 実装しているHHandlerFunc型へキャストすることができる。
	castedFunc := HHandlerFunc(myFunc)
	fmt.Println(reflect.ValueOf(castedFunc).Type()) // main.HHandlerFunc
	// HHandleメソッドにHHandlerインターフェース型として
	// castedFuncを渡すことができる。
	HHandle("/api/v2", castedFunc)

	/// [ポイント1] 埋め込みの関係
	aRealToriJin := &RealToriJin{}
	// ② 構造体ToriJinにFlyingRunnerインターフェースを
	// 実装したRealToriJinの変数を埋め込む
	aToriJin := &ToriJin{
		FlyingRunner: aRealToriJin,
	}
	// ③ 構造体ShinJinruiに構造体Torijinを埋め込む
	// 実装したRealToriJinの変数を埋め込む
	aShinJinrui := &ShinJinrui{
		ToriJin: aToriJin,
	}
	fmt.Println(aShinJinrui.Fly())
	fmt.Println(aShinJinrui.Run())
}
