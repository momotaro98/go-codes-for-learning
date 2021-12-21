package main

import (
	"context"
	"testing"

	"go.uber.org/goleak"
)

func TestA(t *testing.T) {
	defer goleak.VerifyNone(t)

	LeakFunc(context.Background())
}

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}
