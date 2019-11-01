package main

import (
	"fmt"
	"sort"
)

func miniMaxSum(arr []int32) {
	// calc sum
	var sum int64
	for _, n := range arr {
		sum += int64(n)
	}
	a := make([]int, len(arr))
	for i, v := range arr {
		a[i] = int(v)
	}
	// Sort and calc min and max sum
	sort.Sort(sort.IntSlice(a)) // Need to make int to IntSlice type for sort
	min := sum - int64(a[len(a)-1])
	max := sum - int64(a[0])
	fmt.Printf("%d %d", min, max)
}
