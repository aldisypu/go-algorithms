package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/526d84b98f428f14a60008da

func Hammer(n int) uint {
	if n <= 0 {
		return 0
	}

	hammingNumbers := make([]uint, n)
	hammingNumbers[0] = 1

	i2, i3, i5 := 0, 0, 0

	for i := 1; i < n; i++ {
		hammingNumbers[i] = minUint(hammingNumbers[i2]*2, hammingNumbers[i3]*3, hammingNumbers[i5]*5)

		if hammingNumbers[i] == hammingNumbers[i2]*2 {
			i2++
		}
		if hammingNumbers[i] == hammingNumbers[i3]*3 {
			i3++
		}
		if hammingNumbers[i] == hammingNumbers[i5]*5 {
			i5++
		}
	}

	return hammingNumbers[n-1]
}

func minUint(a, b, c uint) uint {
	min := a
	if b < min {
		min = b
	}
	if c < min {
		min = c
	}
	return min
}

func TestHammer(t *testing.T) {
	assert.Equal(t, uint(6), Hammer(6))
	assert.Equal(t, uint(8), Hammer(7))
	assert.Equal(t, uint(9), Hammer(8))
	assert.Equal(t, uint(10), Hammer(9))
	assert.Equal(t, uint(12), Hammer(10))
}
