package main

const size = 100

type myStruct struct {
	arr [size]int
}

func NewMyStructVar() myStruct {
	var ms myStruct
	for i := 0; i < 1; i++ {
	}
	return ms
}

func NewMyStructPtr() *myStruct {
	var ms myStruct
	for i := 0; i < 1; i++ {
	}
	return &ms
}
