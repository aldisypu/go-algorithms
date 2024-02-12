package kata

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/5659c6d896bc135c4c00021e

func NextSmaller(n int) int {
	digits := toDigits(n)

	i := len(digits) - 2
	for i >= 0 && digits[i] <= digits[i+1] {
		i--
	}

	if i == -1 {
		return -1
	}

	j := len(digits) - 1
	for digits[j] >= digits[i] {
		j--
	}

	if i == 0 && digits[j] == 0 {
		return -1
	}

	digits[i], digits[j] = digits[j], digits[i]
	sort.Sort(sort.Reverse(sort.IntSlice(digits[i+1:])))
	result := toNumber(digits)

	if result < n {
		return result
	}

	return -1
}

func toDigits(n int) []int {
	var digits []int
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

func TestNextSmaller(t *testing.T) {
	assert.Equal(t, 790, NextSmaller(907))
	assert.Equal(t, 513, NextSmaller(531))
	assert.Equal(t, -1, NextSmaller(135))
	assert.Equal(t, 2017, NextSmaller(2071))
	assert.Equal(t, 144, NextSmaller(414))
}
