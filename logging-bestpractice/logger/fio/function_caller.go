package fio

import (
	"errors"
	"strings"
	"sync"
)

var (
	closeFunctionCaller  = &functionCaller{}
	rotateFunctionCaller = &functionCaller{}
)

type callerFunc func() error

type functionCaller struct {
	mtx     sync.RWMutex
	callers []callerFunc
}

func (c *functionCaller) append(f callerFunc) {
	c.mtx.Lock()
	defer c.mtx.Unlock()
	c.callers = append(c.callers, f)
}

func (c *functionCaller) call() error {
	var messages []string
	for _, callerFunc := range c.callers {
		if err := callerFunc(); err != nil {
			messages = append(messages, err.Error())
		}
	}
	if len(messages) > 0 {
		return errors.New(strings.Join(messages, ","))
	}
	return nil
}

// Close is ...
func Close() error {
	return closeFunctionCaller.call()
}

// Rotate is ...
func Rotate() error {
	return rotateFunctionCaller.call()
}
