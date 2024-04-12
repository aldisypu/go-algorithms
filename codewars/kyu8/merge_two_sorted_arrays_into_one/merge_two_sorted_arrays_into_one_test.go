package kata

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/5899642f6e1b25935d000161

func MergeArrays(arr1, arr2 []int) []int {
	mergedMap := make(map[int]bool)

	for _, num := range arr1 {
		mergedMap[num] = true
	}

	for _, num := range arr2 {
		mergedMap[num] = true
	}

	merged := make([]int, 0, len(mergedMap))
	for num := range mergedMap {
		merged = append(merged, num)
	}

	sort.Ints(merged)
	return merged
}

func TestMergeArrays(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8}, MergeArrays([]int{1, 2, 3, 4}, []int{5, 6, 7, 8}))
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, MergeArrays([]int{1, 3, 5, 7, 9}, []int{10, 8, 6, 4, 2}))
	assert.Equal(t, []int{1, 2, 3, 4, 5, 7, 9, 10, 11, 12}, MergeArrays([]int{1, 3, 5, 7, 9, 11, 12}, []int{1, 2, 3, 4, 5, 10, 12}))
}
