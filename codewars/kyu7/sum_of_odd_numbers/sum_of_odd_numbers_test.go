package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/55fd2d567d94ac3bc9000064

func RowSumOddNumbers(n int) int {
	return n * n * n
}

func TestRowSumOddNumbers(t *testing.T) {
	assert.Equal(t, 1, RowSumOddNumbers(1))
	assert.Equal(t, 8, RowSumOddNumbers(2))
	assert.Equal(t, 2197, RowSumOddNumbers(13))
	assert.Equal(t, 6859, RowSumOddNumbers(19))
	assert.Equal(t, 68921, RowSumOddNumbers(41))
	assert.Equal(t, 74088, RowSumOddNumbers(42))
	assert.Equal(t, 405224, RowSumOddNumbers(74))
	assert.Equal(t, 636056, RowSumOddNumbers(86))
	assert.Equal(t, 804357, RowSumOddNumbers(93))
	assert.Equal(t, 1030301, RowSumOddNumbers(101))
}
