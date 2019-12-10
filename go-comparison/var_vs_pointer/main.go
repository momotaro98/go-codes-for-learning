package main

func main() {
	var list []*myStruct
	for i := 0; i < 1000; i++ {
		v := NewMyStructPtr()
		list = append(list, v)
	}
}
