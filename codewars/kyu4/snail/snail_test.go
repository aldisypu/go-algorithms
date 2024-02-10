package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/521c2db8ddc89b9b7a0000c1

func Snail(snaipMap [][]int) []int {
	result := []int{}

	if len(snaipMap) == 0 || len(snaipMap[0]) == 0 {
		return result
	}

	for len(snaipMap) > 0 {
		result = append(result, snaipMap[0]...)

		snaipMap = snaipMap[1:]

		snaipMap = transposeAndReverse(snaipMap)
	}

	return result
}

func transposeAndReverse(matrix [][]int) [][]int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return matrix
	}

	transposed := make([][]int, len(matrix[0]))
	for i := range transposed {
		transposed[i] = make([]int, len(matrix))
		for j := range transposed[i] {
			transposed[i][j] = matrix[j][len(matrix[0])-1-i]
		}
	}

	return transposed
}

func TestSnail(t *testing.T) {
	assert.Equal(t, []int{}, Snail([][]int{}))
	assert.Equal(t, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}, Snail([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}
