package main

const size = 10

type myStruct struct {
	arr [size]int
}

func NewMyStructVar() myStruct {
	var ms myStruct
	for i := 0; i < size; i++ {
	}
	return ms
}

func NewMyStructPtr() *myStruct {
	var ms myStruct
	for i := 0; i < size; i++ {
	}
	return &ms
}
