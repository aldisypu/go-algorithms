package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/52597aa56021e91c93000cb0

func MoveZeros(arr []int) []int {
	var i, j int

	for i < len(arr) {
		if arr[i] != 0 {
			arr[j], arr[i] = arr[i], arr[j]
			j++
		}
		i++
	}

	return arr
}

func TestMoveZeros(t *testing.T) {
	assert.Equal(t, []int{1, 2, 1, 1, 3, 1, 0, 0, 0, 0}, MoveZeros([]int{1, 2, 0, 1, 0, 1, 0, 3, 0, 1}))
}
