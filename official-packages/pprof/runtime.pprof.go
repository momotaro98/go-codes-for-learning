package main

import (
	"os"
	"runtime/pprof"
	"time"
)

func run() error {
	f, err := os.Create("cpu.pprof")
	if err != nil {
		return err
	}

	if err := pprof.StartCPUProfile(f); err != nil {
		return err
	}
	defer pprof.StopCPUProfile()

	process := func() {
		time.Sleep(time.Second * 5)
	}
	process()

	return nil
}

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}
