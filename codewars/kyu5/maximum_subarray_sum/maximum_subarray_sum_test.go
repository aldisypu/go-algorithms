package kata

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/54521e9ec8e60bc4de000d6c

func MaximumSubarraySum(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	var maxSum, currentSum int

	for _, num := range numbers {
		currentSum = int(math.Max(float64(num), float64(currentSum+num)))
		maxSum = int(math.Max(float64(maxSum), float64(currentSum)))
	}

	return maxSum
}

func TestMaximumSubarraySum(t *testing.T) {
	assert.Equal(t, 0, MaximumSubarraySum([]int{}))
	assert.Equal(t, 6, MaximumSubarraySum([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	assert.Equal(t, 0, MaximumSubarraySum([]int{-2, -1, -3, -4, -1, -2, -1, -5, -4}))
}
