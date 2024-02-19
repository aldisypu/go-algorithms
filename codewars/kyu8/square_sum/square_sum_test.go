package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/515e271a311df0350d00000f

func SquareSum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num * num
	}
	return sum
}

func TestSquareSum(t *testing.T) {
	assert.Equal(t, 5, SquareSum([]int{1, 2}))
	assert.Equal(t, 50, SquareSum([]int{0, 3, 4, 5}))
	assert.Equal(t, 0, SquareSum([]int{}))
}
