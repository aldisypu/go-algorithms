package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/57f781872e3d8ca2a000007e

func Maps(x []int) []int {
	result := make([]int, len(x))
	for i, v := range x {
		result[i] = v * 2
	}
	return result
}

func TestMaps(t *testing.T) {
	assert.Equal(t, []int{2, 4, 6}, Maps([]int{1, 2, 3}))
	assert.Equal(t, []int{8, 2, 2, 2, 8}, Maps([]int{4, 1, 1, 1, 4}))
	assert.Equal(t, []int{4, 4, 4, 4, 4, 4}, Maps([]int{2, 2, 2, 2, 2, 2}))
}
