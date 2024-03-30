package kata

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/57f6ad55cca6e045d2000627

func SquareOrSquareRoot(arr []int) []int {
	result := make([]int, len(arr))
	for i, num := range arr {
		sqrt := int(math.Sqrt(float64(num)))
		if sqrt*sqrt == num {
			result[i] = sqrt
		} else {
			result[i] = num * num
		}
	}
	return result
}

func TestSquareOrSquareRoot(t *testing.T) {
	assert.Equal(t, []int{2, 9, 3, 49, 4, 1}, SquareOrSquareRoot([]int{4, 3, 9, 7, 2, 1}))
	assert.Equal(t, []int{10, 10201, 25, 25, 1, 1}, SquareOrSquareRoot([]int{100, 101, 5, 5, 1, 1}))
	assert.Equal(t, []int{1, 4, 9, 2, 25, 36}, SquareOrSquareRoot([]int{1, 2, 3, 4, 5, 6}))
}
