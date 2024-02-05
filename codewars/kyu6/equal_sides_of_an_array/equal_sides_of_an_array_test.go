package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func FindEvenIndex(arr []int) int {
	total := 0
	leftSum := 0

	for _, v := range arr {
		total += v
	}

	for i, v := range arr {
		total -= v
		if leftSum == total {
			return i
		}
		leftSum += v
	}

	return -1
}

func TestFindEvenIndex(t *testing.T) {
	assert.Equal(t, 3, FindEvenIndex([]int{1, 2, 3, 4, 3, 2, 1}))
	assert.Equal(t, 1, FindEvenIndex([]int{1, 100, 50, -51, 1, 1}))
	assert.Equal(t, -1, FindEvenIndex([]int{1, 2, 3, 4, 5, 6}))
	assert.Equal(t, 3, FindEvenIndex([]int{20, 10, 30, 10, 10, 15, 35}))
	assert.Equal(t, 0, FindEvenIndex([]int{20, 10, -80, 10, 10, 15, 35}))
}
