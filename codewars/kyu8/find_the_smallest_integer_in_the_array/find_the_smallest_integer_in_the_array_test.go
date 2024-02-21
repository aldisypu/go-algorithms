package kata

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/55a2d7ebe362935a210000b2

func SmallestIntegerFinder(numbers []int) int {
	sort.Ints(numbers)
	return numbers[0]
}

func TestSmallestIntegerFinder(t *testing.T) {
	assert.Equal(t, 2, SmallestIntegerFinder([]int{34, 15, 88, 2}))
	assert.Equal(t, -345, SmallestIntegerFinder([]int{34, -345, -1, 100}))
}
