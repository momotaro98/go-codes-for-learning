package main

type Flyer interface {
	Fly()
}

type Runner interface {
	Run()
}

type Bird struct {
}
