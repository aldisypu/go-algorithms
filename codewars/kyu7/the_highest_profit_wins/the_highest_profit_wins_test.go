package kata

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/559590633066759614000063

func MinMax(arr []int) [2]int {
	sort.Ints(arr)
	return [2]int{arr[0], arr[len(arr)-1]}
}

func TestMinMax(t *testing.T) {
	assert.Equal(t, [2]int{1, 5}, MinMax([]int{1, 2, 3, 4, 5}))
	assert.Equal(t, [2]int{5, 2334454}, MinMax([]int{2334454, 5}))
}
