package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/5715eaedb436cf5606000381/solutions/go

func PositiveSum(numbers []int) int {
	result := 0
	for _, number := range numbers {
		if number > 0 {
			result += number
		}
	}
	return result
}

func TestPositiveSum(t *testing.T) {
	assert.Equal(t, 15, PositiveSum([]int{1, 2, 3, 4, 5}))
	assert.Equal(t, 13, PositiveSum([]int{1, -2, 3, 4, 5}))
	assert.Equal(t, 0, PositiveSum([]int{}))
	assert.Equal(t, 0, PositiveSum([]int{-1, -2, -3, -4, -5}))
	assert.Equal(t, 9, PositiveSum([]int{-1, 2, 3, 4, -5}))
}
