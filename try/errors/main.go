package main

import (
	"errors"
	"fmt"
	"runtime"
)

func util() string {
	pc := make([]uintptr, 1)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	return fmt.Sprintf("%s:%d %s\n", frame.File, frame.Line, frame.Function)
}

func repo() error {
	err := errors.New("cause of error")
	return fmt.Errorf("%w %s", err, util())
}

type serviceError struct {
	error
	code int
}

func service() error {
	err := repo()
	if err != nil {
		return serviceError{
			error: fmt.Errorf("%w %s", err, util()),
			code:  10,
		}
	}

	return nil
}

func main() {
	err := service()
	if err != nil {
		fmt.Println(err)
	}
}
