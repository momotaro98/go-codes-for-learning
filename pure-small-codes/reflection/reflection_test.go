package main

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {

	// 無名StructでテストケースをTableとして作る
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct{ Name string }{"Yamada"},
			[]string{"Yamada"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Yamada", "Saitama"},
			[]string{"Yamada", "Saitama"},
		},
		{
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"Yamada", 24},
			[]string{"Yamada"},
		},
		{
			"Nested fields",
			Person{
				Name:    "John",
				Profile: Profile{Age: 33, City: "London"},
			},
			[]string{"John", "London"},
		},
		{
			"Pointers to things",
			&Person{
				Name:    "John",
				Profile: Profile{Age: 33, City: "London"},
			},
			[]string{"John", "London"},
		},
		{
			"Slices",
			[]Profile{
				{Age: 33, City: "London"},
				{Age: 17, City: "Paris"},
			},
			[]string{"London", "Paris"},
		},
		{
			"Arrays",
			[2]Profile{
				{Age: 33, City: "London"},
				{Age: 17, City: "Paris"},
			},
			[]string{"London", "Paris"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) { // 各テストケースをサブテストとして扱う
			// Act
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})
			// Assert
			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}
