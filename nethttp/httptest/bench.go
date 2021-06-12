package main

import (
	"net/http"
	"time"
)

func Handler(rw http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Millisecond * 50)
	msg := []byte("Hello world.\n")
	rw.Write(msg)
}
