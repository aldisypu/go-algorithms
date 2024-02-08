package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/561e9c843a2ef5a40c0000a4

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func Gap(g, m, n int) []int {
	previousPrime := -1
	for i := m; i <= n; i++ {
		if isPrime(i) {
			if previousPrime != -1 && i-previousPrime == g {
				return []int{previousPrime, i}
			}
			previousPrime = i
		}
	}

	return nil
}

func TestGap(t *testing.T) {
	assert.Equal(t, []int{101, 103}, Gap(2, 100, 110))
	assert.Equal(t, []int{103, 107}, Gap(4, 100, 110))
	assert.Nil(t, Gap(6, 100, 110))
	assert.Equal(t, []int{359, 367}, Gap(8, 300, 400))
	assert.Equal(t, []int{337, 347}, Gap(10, 300, 400))
}
