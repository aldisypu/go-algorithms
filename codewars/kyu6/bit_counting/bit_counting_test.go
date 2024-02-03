package kata

import (
	"math/bits"
	"testing"

	"github.com/stretchr/testify/assert"
)

func CountBits(n uint) int {
	return bits.OnesCount(n)
}

func TestCountBits(t *testing.T) {
	assert.Equal(t, 0, CountBits(0))
	assert.Equal(t, 1, CountBits(4))
	assert.Equal(t, 3, CountBits(7))
	assert.Equal(t, 2, CountBits(9))
	assert.Equal(t, 2, CountBits(10))
}
