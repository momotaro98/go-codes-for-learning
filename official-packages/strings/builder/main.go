package main

import (
	"fmt"
	"strings"
)

func f(myList []string) (string, error) {
	const (
		Threshold = 1000
	)

	var (
		builderList = []strings.Builder{}
	)
	var (
		bullder = strings.Builder{}
	)
	var (
		count uint64 = 0
	)
	for range myList {
		if count > Threshold {
			builderList = append(builderList, bullder)
			bullder = strings.Builder{}
			count = 0
		}
		bullder.WriteString("123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890")
		count++
	}

	for _, l := range builderList {
		s := l.String()
		fmt.Println(len(s))
	}

	return bullder.String(), nil
}

func main() {
	myList := make([]string, 1000000)
	_, err := f(myList)
	if err != nil {
		panic(err)
	}
	fmt.Println("success")
}
