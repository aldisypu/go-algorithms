package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/55d24f55d7dd296eb9000030

func Summation(n int) int {
	return n * (n + 1) / 2
}

func TestSummation(t *testing.T) {
	assert.Equal(t, 1, Summation(1))
	assert.Equal(t, 36, Summation(8))
	assert.Equal(t, 253, Summation(22))
	assert.Equal(t, 5050, Summation(100))
	assert.Equal(t, 22791, Summation(213))
}
