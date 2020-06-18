package main

import (
	"fmt"

	"errors"
)

func repo() error {
	cause := errors.New("core cause of error like DB command execution error")
	return errors.WithStack(cause)
}

func service() error {
	return repo()
}

func main() {
	fmt.Printf("%+v", service())
}
