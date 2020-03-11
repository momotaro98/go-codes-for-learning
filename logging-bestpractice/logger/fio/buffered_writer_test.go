package fio

import (
	"bufio"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBufferedWriter(t *testing.T) {
	actual := NewBufferedWriter()
	expected := bufio.NewWriterSize(os.Stderr, defaultBufferSize)
	assert.Equal(t, expected, actual)
}
