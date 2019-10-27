package main

import "fmt"

// パフォーマンスが悪いが読みやすい
// TODO: ikeda_zipとzipの比較用ベンチマークをさくっと作って検証してみよう。
func ikeda_zip(lists ...[]int) func() []int {
	i := 0
	return func() []int {
		var ret []int
		for _, list := range lists {
			if len(list) <= i {
				return nil
			}
			ret = append(ret, list[i])
		}
		i++
		return ret
	}
}

// パフォーマンスが良いがちょっと読みにくい
// zip関数の環境としてもっているzipというArrayのメモリ領域を上書きして使いまわしていてエコ
// Link: https://stackoverflow.com/questions/26957040/how-to-implement-the-python-zip-function-in-golang/26958771#26958771
func zip(lists ...[]int) func() []int {
	zip := make([]int, len(lists))
	i := 0
	return func() []int {
		for j, _ := range lists {
			if len(lists[j]) <= i {
				return nil
			}
			zip[j] = lists[j][i]
		}
		i++
		return zip
	}
}

func main() {
	a := [3]int{4, 3, 5}
	b := [3]int{2, 1, 7}
	c := [3]int{9, 8, 6}
	iter := zip(a[:], b[:], c[:]) // a[:] is a golang basic tip for array to slice
	for tupple := iter(); tupple != nil; tupple = iter() {
		fmt.Println(tupple)
	}
}
