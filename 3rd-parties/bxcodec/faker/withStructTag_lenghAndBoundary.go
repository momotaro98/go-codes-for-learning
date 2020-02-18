package main

import (
	"fmt"

	"github.com/bxcodec/faker/v3"
)

// SomeStruct ...
type SomeStruct struct {
	Inta  int   `faker:"boundary_start=5, boundary_end=10"`
	Int8  int8  `faker:"boundary_start=100, boundary_end=1000"`
	Int16 int16 `faker:"boundary_start=123, boundary_end=1123"`
	Int32 int32 `faker:"boundary_start=-10, boundary_end=8123"`
	Int64 int64 `faker:"boundary_start=31, boundary_end=88"`

	UInta  uint   `faker:"boundary_start=35, boundary_end=152"`
	UInt8  uint8  `faker:"boundary_start=5, boundary_end=1425"`
	UInt16 uint16 `faker:"boundary_start=245, boundary_end=2125"`
	UInt32 uint32 `faker:"boundary_start=0, boundary_end=40"`
	UInt64 uint64 `faker:"boundary_start=14, boundary_end=50"`

	ASString []string          `faker:"len=50"`
	SString  string            `faker:"len=25"`
	MSString map[string]string `faker:"len=30"`
	MIint    map[int]int       `faker:"boundary_start=5, boundary_end=10"`
}

func main() {
	faker.SetRandomMapAndSliceSize(20) //Random generated map or array size wont exceed 20...
	a := SomeStruct{}
	err := faker.FakeData(&a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", a)
}
