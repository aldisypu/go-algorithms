package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/55e7280b40e1c4a06d0000aa

func ChooseBestSum(t, k int, ls []int) int {
	calculateSum := func(combination []int) int {
		sum := 0
		for _, index := range combination {
			sum += ls[index]
		}
		return sum
	}

	var combinations func([]int, int, int, []int, int, *int)
	combinations = func(arr []int, index int, r int, data []int, i int, result *int) {
		if index == r {
			currentSum := calculateSum(data)
			if currentSum <= t && currentSum > *result {
				*result = currentSum
			}
			return
		}

		if i >= len(arr) {
			return
		}

		data[index] = arr[i]
		combinations(arr, index+1, r, data, i+1, result)
		combinations(arr, index, r, data, i+1, result)
	}

	n := len(ls)
	indices := make([]int, n)
	for i := 0; i < n; i++ {
		indices[i] = i
	}

	result := -1
	combinations(indices, 0, k, make([]int, k), 0, &result)

	if result == -1 {
		return -1
	}

	return result
}

func TestChooseBestSum(t *testing.T) {
	assert.Equal(t, 163, ChooseBestSum(163, 3, []int{50, 55, 56, 57, 58}))
	assert.Equal(t, -1, ChooseBestSum(163, 3, []int{50}))
	assert.Equal(t, 228, ChooseBestSum(230, 3, []int{91, 74, 73, 85, 73, 81, 87}))
	assert.Equal(t, 178, ChooseBestSum(331, 2, []int{91, 74, 73, 85, 73, 81, 87}))
	assert.Equal(t, 331, ChooseBestSum(331, 4, []int{91, 74, 73, 85, 73, 81, 87}))
	assert.Equal(t, -1, ChooseBestSum(331, 5, []int{91, 74, 73, 85, 73, 81, 87}))
}
