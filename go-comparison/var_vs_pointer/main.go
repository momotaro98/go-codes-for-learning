package main

func main() {
	var sum int
	for i := 0; i < 1000; i++ {
		v := NewMyStructPtr()
		sum += v.arr[0]
	}
}
