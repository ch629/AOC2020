package day3

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 211, SolvePart1())
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 3584591857, SolvePart2())
}
