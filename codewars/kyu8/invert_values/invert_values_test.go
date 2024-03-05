package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/5899dc03bc95b1bf1b0000ad

func Invert(arr []int) []int {
	result := make([]int, len(arr))
	for i, num := range arr {
		result[i] = -num
	}
	return result
}

func TestInvert(t *testing.T) {
	assert.Equal(t, []int{-1, -2, -3, -4, -5}, Invert([]int{1, 2, 3, 4, 5}))
	assert.Equal(t, []int{-1, 2, -3, 4, -5}, Invert([]int{1, -2, 3, -4, 5}))
	assert.Equal(t, []int{0}, Invert([]int{0}))
}
