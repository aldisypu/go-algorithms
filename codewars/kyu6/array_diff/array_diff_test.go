package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/523f5d21c841566fde000009

func ArrayDiff(a, b []int) []int {
	bMap := make(map[int]bool)
	result := make([]int, 0)

	for _, num := range b {
		bMap[num] = true
	}

	for _, num := range a {
		if !bMap[num] {
			result = append(result, num)
		}
	}

	return result
}

func TestArrayDiff(t *testing.T) {
	assert.Equal(t, []int{2}, ArrayDiff([]int{1, 2}, []int{1}))
	assert.Equal(t, []int{2, 2}, ArrayDiff([]int{1, 2, 2}, []int{1}))
	assert.Equal(t, []int{1}, ArrayDiff([]int{1, 2, 2}, []int{2}))
	assert.Equal(t, []int{1, 2, 2}, ArrayDiff([]int{1, 2, 2}, []int{}))
	assert.Empty(t, ArrayDiff([]int{}, []int{1, 2}))
	assert.Equal(t, []int{3}, ArrayDiff([]int{1, 2, 3}, []int{1, 2}))
}
