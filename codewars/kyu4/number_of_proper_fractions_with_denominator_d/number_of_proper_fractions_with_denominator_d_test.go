package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func ProperFractions(n int) int {
	if n == 1 {
		return 0
	}

	count := n
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			for n%i == 0 {
				n /= i
			}
			count -= count / i
		}
	}

	if n > 1 {
		count -= count / n
	}

	return count
}

func TestProperFractions(t *testing.T) {
	assert.Equal(t, 0, ProperFractions(1))
	assert.Equal(t, 1, ProperFractions(2))
	assert.Equal(t, 4, ProperFractions(5))
	assert.Equal(t, 8, ProperFractions(15))
	assert.Equal(t, 20, ProperFractions(25))
}
