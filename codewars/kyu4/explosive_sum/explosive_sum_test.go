package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/52ec24228a515e620b0005ef

func ExpSum(n uint64) uint64 {
	partitions := make([]uint64, n+1)
	partitions[0] = 1

	for i := 1; i <= int(n); i++ {
		for j := i; j <= int(n); j++ {
			partitions[j] += partitions[j-i]
		}
	}

	return partitions[n]
}

func TestExpSum(t *testing.T) {
	assert.Equal(t, uint64(1), ExpSum(1))
	assert.Equal(t, uint64(2), ExpSum(2))
	assert.Equal(t, uint64(3), ExpSum(3))
	assert.Equal(t, uint64(5), ExpSum(4))
	assert.Equal(t, uint64(7), ExpSum(5))
	assert.Equal(t, uint64(42), ExpSum(10))
}
