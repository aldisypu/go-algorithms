package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/5583090cbe83f4fd8c000051

func Digitize(n int) []int {
	var result []int

	if n == 0 {
		return []int{0}
	}

	for n > 0 {
		digit := n % 10
		result = append(result, digit)
		n /= 10
	}

	return result
}

func TestDigitize(t *testing.T) {
	assert.Equal(t, []int{1, 3, 2, 5, 3}, Digitize(35231))
	assert.Equal(t, []int{0}, Digitize(0))
}
