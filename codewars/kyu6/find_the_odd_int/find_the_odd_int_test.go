package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/54da5a58ea159efa38000836

func FindOdd(seq []int) int {
	result := 0

	for _, num := range seq {
		result ^= num
	}

	return result
}

func TestFindOdd(t *testing.T) {
	assert.Equal(t, 5, FindOdd([]int{20, 1, -1, 2, -2, 3, 3, 5, 5, 1, 2, 4, 20, 4, -1, -2, 5}))
	assert.Equal(t, -1, FindOdd([]int{1, 1, 2, -2, 5, 2, 4, 4, -1, -2, 5}))
	assert.Equal(t, 5, FindOdd([]int{20, 1, 1, 2, 2, 3, 3, 5, 5, 4, 20, 4, 5}))
	assert.Equal(t, 10, FindOdd([]int{10}))
	assert.Equal(t, 10, FindOdd([]int{1, 1, 1, 1, 1, 1, 10, 1, 1, 1, 1}))
	assert.Equal(t, 1, FindOdd([]int{5, 4, 3, 2, 1, 5, 4, 3, 2, 10, 10}))
}
