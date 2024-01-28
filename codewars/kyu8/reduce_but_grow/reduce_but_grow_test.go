package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/57f780909f7e8e3183000078

func Grow(arr []int) int {
	result := 1
	for _, num := range arr {
		result *= num
	}
	return result
}

func TestGrow(t *testing.T) {
	assert.Equal(t, 6, Grow([]int{1, 2, 3}))
	assert.Equal(t, 16, Grow([]int{4, 1, 1, 1, 4}))
	assert.Equal(t, 64, Grow([]int{2, 2, 2, 2, 2, 2}))
}
