package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/52a382ee44408cea2500074c

func Determinant(matrix [][]int) int {
	rows := len(matrix)

	if rows == 1 {
		return matrix[0][0]
	}
	if rows == 2 {
		return matrix[0][0]*matrix[1][1] - matrix[0][1]*matrix[1][0]
	}

	det := 0

	for j, element := range matrix[0] {
		minor := make([][]int, rows-1)
		for i := 1; i < rows; i++ {
			minor[i-1] = append(minor[i-1], matrix[i][:j]...)
			minor[i-1] = append(minor[i-1], matrix[i][j+1:]...)
		}

		sign := 1
		if j%2 == 1 {
			sign = -1
		}
		det += sign * element * Determinant(minor)
	}

	return det
}

func TestDeterminant(t *testing.T) {
	assert.Equal(t, 1, Determinant([][]int{{1}}))
	assert.Equal(t, -1, Determinant([][]int{{1, 3}, {2, 5}}))
	assert.Equal(t, -20, Determinant([][]int{{2, 5, 3}, {1, -2, -1}, {1, 3, 4}}))
}
