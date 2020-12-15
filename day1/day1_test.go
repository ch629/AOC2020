package day1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolvePair(t *testing.T) {
	solution := SolvePair()

	assert.Equal(t, 960075, solution)
}

func TestSolveTriple(t *testing.T) {
	solution := SolveTriple()

	assert.Equal(t, 212900130, solution)
}
