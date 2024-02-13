package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/52f677797c461daaf7000740

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func arrayGCD(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	result := arr[0]
	for i := 1; i < len(arr); i++ {
		result = gcd(result, arr[i])
	}
	return result
}

func Solution(arr []int) int {
	gcdValue := arrayGCD(arr)
	return len(arr) * gcdValue
}

func TestSolution(t *testing.T) {
	assert.Equal(t, 9, Solution([]int{9}))
	assert.Equal(t, 9, Solution([]int{6, 9, 21}))
	assert.Equal(t, 3, Solution([]int{1, 21, 55}))
}
