package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/56f69d9f9400f508fb000ba7

func monkeyCount(n int) []int {
	result := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		result = append(result, i)
	}
	return result
}

func TestMonkeyCount(t *testing.T) {
	assert.Equal(t, []int{1}, monkeyCount(1))
	assert.Equal(t, []int{1, 2}, monkeyCount(2))
	assert.Equal(t, []int{1, 2, 3}, monkeyCount(3))
	assert.Equal(t, []int{1, 2, 3, 4}, monkeyCount(4))
	assert.Equal(t, []int{1, 2, 3, 4, 5}, monkeyCount(5))
}
