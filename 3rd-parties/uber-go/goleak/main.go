package main

import (
	"context"
)

func LeakFunc(ctx context.Context) {
	go func() {
		<-ctx.Done()
	}()
	return
}
