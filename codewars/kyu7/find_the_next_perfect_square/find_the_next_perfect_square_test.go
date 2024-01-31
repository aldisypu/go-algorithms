package kata

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/56269eb78ad2e4ced1000013

func FindNextSquare(sq int64) int64 {
	root := int64(math.Sqrt(float64(sq)))

	if root*root != sq {
		return -1
	}

	nextRoot := root + 1
	return nextRoot * nextRoot
}

func TestFindNextSquare(t *testing.T) {
	assert.Equal(t, int64(144), FindNextSquare(121))
	assert.Equal(t, int64(676), FindNextSquare(625))
	assert.Equal(t, int64(320356), FindNextSquare(319225))
	assert.Equal(t, int64(15241630849), FindNextSquare(15241383936))
	assert.Equal(t, int64(-1), FindNextSquare(155))
}
