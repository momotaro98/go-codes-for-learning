package main

import (
	"fmt"

	"github.com/bxcodec/faker/v3"
)

// SomeStruct ...
type SomeStruct struct {
	Word string `faker:"word,unique"`
}

func main() {
	for i := 0; i < 5; i++ { // Generate 5 structs having a unique word
		a := SomeStruct{}
		err := faker.FakeData(&a)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%+v", a)
	}
	faker.ResetUnique() // Forget all generated unique values. Allows to start generating another unrelated dataset.
}
