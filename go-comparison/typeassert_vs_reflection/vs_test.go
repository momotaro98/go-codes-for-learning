package main

import (
	"reflect"
	"strconv"
	"testing"
)

var input []interface{}

func init() {
	for i := 1; i <= 10000; i++ {
		input = append(input, strconv.Itoa(i))
	}
}

// リフレクションを使う関数 は`interface{}`型で受け取りリフレクションをするので
// 柔軟に引数の値をさばくことができる。
func リフレクションを使う関数(target interface{}) []string {
	if target == nil {
		return nil
	}

	v := reflect.ValueOf(target)

	switch v.Kind() {
	case reflect.Slice:
		ret := make([]string, v.Len())
		for i := 0; i < v.Len(); i++ {
			x := v.Index(i)
			k := x.Kind()
			if k == reflect.Interface || k == reflect.String {
				s, ok := x.Interface().(string)
				if ok {
					ret[i] = s
				}
			}
		}
		return ret
	case reflect.String: // stringが引数の場合はそれをスライスにして返す。
		return []string{v.String()}
	default:
		return nil
	}
}

// 型アサーションを使う関数 はスライスであることを決め打ちしたうえで`[]string`
// に変換して返す関数である。
func 型アサーションを使う関数(target []interface{}) []string {
	ret := make([]string, len(target))
	for i, r := range target {
		s, ok := r.(string)
		if ok {
			ret[i] = s
		}
	}
	return ret
}

func Benchmark_リフレクション(b *testing.B) {
	var target interface{}
	target = input
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		リフレクションを使う関数(target)
	}
}

func Benchmark_型アサーション(b *testing.B) {
	var target interface{}
	target = input
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		型アサーションを使う関数(target.([]interface{}))
	}
}
