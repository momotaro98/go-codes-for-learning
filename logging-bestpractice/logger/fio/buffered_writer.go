package fio

import (
	"bufio"
	"io"
	"os"
)

const (
	defaultBufferSize = 1028
)

// NewBufferedWriter is ...
func NewBufferedWriter() io.Writer {
	return NewBufferedWriterWithSize(os.Stderr, defaultBufferSize)
}

// NewBufferedWriterWithSize is ...
func NewBufferedWriterWithSize(writer io.Writer, size int) io.Writer {
	out := bufio.NewWriterSize(writer, size)
	// append close caller function
	closeFunctionCaller.append(func() error {
		return out.Flush()
	})
	return out
}
