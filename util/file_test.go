package util

import (
	"github.com/stretchr/testify/assert"
	"io"
	"strings"
	"testing"
)

type StringCloser struct {
	io.Reader
}

func (str StringCloser) Close() error {
	return nil
}

func TestReadNumbers(t *testing.T) {
	ints := ReadNumbers(StringCloser{
		strings.NewReader("123\n456"),
	})

	assert.Equal(t, []int{123, 456}, ints)
}
