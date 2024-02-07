package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/52f787eb172a8b4ae1000a34

func Zeros(n int) int {
	trailingZeros := 0

	for n >= 5 {
		n /= 5
		trailingZeros += n
	}

	return trailingZeros
}

func TestZeros(t *testing.T) {
	assert.Equal(t, 0, Zeros(0))
	assert.Equal(t, 1, Zeros(6))
	assert.Equal(t, 7, Zeros(30))
}
