package kata

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/55983863da40caa2c900004e

func NextBigger(n int) int {
	digits := toDigits(n)

	i := len(digits) - 2
	for i >= 0 && digits[i] >= digits[i+1] {
		i--
	}

	if i == -1 {
		return -1
	}

	j := len(digits) - 1
	for digits[j] <= digits[i] {
		j--
	}

	digits[i], digits[j] = digits[j], digits[i]
	sort.Ints(digits[i+1:])
	result := toNumber(digits)

	return result
}

func toDigits(n int) []int {
	digits := make([]int, 0)
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}
	reverse(digits)
	return digits
}

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func toNumber(digits []int) int {
	result := 0
	for _, digit := range digits {
		result = result*10 + digit
	}
	return result
}

func TestNextBigger(t *testing.T) {
	assert.Equal(t, 21, NextBigger(12))
	assert.Equal(t, 531, NextBigger(513))
	assert.Equal(t, 2071, NextBigger(2017))
	assert.Equal(t, 441, NextBigger(414))
	assert.Equal(t, 414, NextBigger(144))
}
