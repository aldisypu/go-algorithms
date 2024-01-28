package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/5513795bd3fafb56c200049e

func CountBy(x, n int) []int {
	if n == 0 {
		return []int{}
	}

	return append(CountBy(x, n-1), n*x)
}

func TestCountBy(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 4, 5}, CountBy(1, 5))
	assert.Equal(t, []int{2, 4, 6, 8, 10}, CountBy(2, 5))
	assert.Equal(t, []int{3, 6, 9, 12, 15}, CountBy(3, 5))
	assert.Equal(t, []int{50, 100, 150, 200, 250}, CountBy(50, 5))
	assert.Equal(t, []int{100, 200, 300, 400, 500}, CountBy(100, 5))
}
