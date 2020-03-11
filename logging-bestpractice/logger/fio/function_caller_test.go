package fio

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClose(t *testing.T) {
	closeFunctionCaller.append(func() error {
		return nil
	})
	closeFunctionCaller.append(func() error {
		return errors.New("unexpected error-1")
	})
	closeFunctionCaller.append(func() error {
		return errors.New("unexpected error-2")
	})
	actual := Close()
	expected := errors.New("unexpected error-1,unexpected error-2")
	assert.Equal(t, expected, actual)
}

func TestRotate(t *testing.T) {
	rotateFunctionCaller.append(func() error {
		return nil
	})
	rotateFunctionCaller.append(func() error {
		return errors.New("unexpected error-1")
	})
	rotateFunctionCaller.append(func() error {
		return errors.New("unexpected error-2")
	})
	actual := Rotate()
	expected := errors.New("unexpected error-1,unexpected error-2")
	assert.Equal(t, expected, actual)
}
