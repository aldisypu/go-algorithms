package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/55ecd718f46fba02e5000029

func Between(a, b int) []int {
	var result []int
	for i := a; i <= b; i++ {
		result = append(result, i)
	}
	return result
}

func TestBetween(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 4}, Between(1, 4))
	assert.Equal(t, []int{-2, -1, 0, 1, 2}, Between(-2, 2))
}
