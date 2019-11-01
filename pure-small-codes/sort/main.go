package main

import (
	"fmt"
	"sort"
)

// Sortなインターフェイスを持つための型
// Len, Swap, Lessのメソッドを実装する必要がある
type myrunes []rune

func (r myrunes) Len() int {
	return len(r)
}

func (r myrunes) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r myrunes) Less(i, j int) bool {
	return r[i] < r[j] // ASC (96, 98, 101...)
}

func twoStrings(s1 string, s2 string) string {
	var runes1 myrunes
	for _, r := range s1 {
		runes1 = append(runes1, r)
	}
	var runes2 myrunes
	for _, r := range s2 {
		runes2 = append(runes2, r)
	}
	sort.Sort(runes1) // myrunes型はインターフェイスを実装しているのでSort関数に食わせることで中身がソートされる。
	sort.Sort(runes2)
	if runes1[len(runes1)-1] < runes2[0] || runes2[len(runes2)-1] < runes1[0] {
		return "NO"
	}
	if len(s1) < len(s2) {
		return search(runes1, runes2)
	}
	return search(runes2, runes1)
}

func search(runes1 myrunes, runes2 myrunes) string {
	for _, r1 := range runes1 {
		i := sort.Search(len(runes2), func(i int) bool { return runes2[i] >= r1 })
		if i < len(runes2) && runes2[i] == r1 {
			return "YES"
		}
	}
	return "NO"
}

func main() {
	fmt.Println(twoStrings("aardvark", "apple"))

	// another file
	arr := [5]int32{7, 69, 2, 221, 8974}
	miniMaxSum(arr[:])
}
