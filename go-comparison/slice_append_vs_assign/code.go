package main

func AppendNakedMap(in []int, length int, f func(i int) int) []int {
	var ret []int
	for _, c := range in {
		ret = append(ret, f(c))
	}
	return ret
}

func AssignMap(in []int, length int, f func(i int) int) []int {
	ret := make([]int, length)
	for i, c := range in {
		ret[i] = f(c)
	}
	return ret
}

func AppendCapMap(in []int, length int, f func(i int) int) []int {
	ret := make([]int, 0, length)
	for _, c := range in {
		ret = append(ret, f(c))
	}
	return ret
}
