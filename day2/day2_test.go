package day2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolveCount(t *testing.T) {
	assert.Equal(t, 560, SolveCount())
}

func TestSolvePlaces(t *testing.T) {
	assert.Equal(t, 303, SolvePlaces())
}
