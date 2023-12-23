package prime

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func isPrime(n int) bool {
	// 0 and 1 are not prime
	if n < 2 {
		return false
	}

	// 2 and 3 are prime
	if n < 4 {
		return true
	}

	// all numbers divisible by 2 except 2 are not prime
	if n%2 == 0 {
		return false
	}

	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func TestPrime(t *testing.T) {
	assert.True(t, isPrime(2))
	assert.True(t, isPrime(3))
	assert.True(t, isPrime(5))
	assert.True(t, isPrime(7))
	assert.True(t, isPrime(11))

	assert.False(t, isPrime(1))
	assert.False(t, isPrime(4))
	assert.False(t, isPrime(6))
	assert.False(t, isPrime(8))
	assert.False(t, isPrime(9))
	assert.False(t, isPrime(10))
}
