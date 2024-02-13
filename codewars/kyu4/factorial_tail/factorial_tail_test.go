package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/55c4eb777e07c13528000021

func Zeroes(base, number int) int {
	count := 0

	for prime := 2; prime <= base; prime++ {
		if base%prime == 0 {
			exponent := 0
			for base%prime == 0 {
				exponent++
				base /= prime
			}

			currentCount := 0
			factorialNumber := number
			for factorialNumber >= prime {
				currentCount += factorialNumber / prime
				factorialNumber /= prime
			}
			currentCount /= exponent

			if count == 0 || currentCount < count {
				count = currentCount
			}
		}
	}

	return count
}

func TestZeroes(t *testing.T) {
	assert.Equal(t, 2, Zeroes(10, 10))
	assert.Equal(t, 3, Zeroes(16, 16))
}
