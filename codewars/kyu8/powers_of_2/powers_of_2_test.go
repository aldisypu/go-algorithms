package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/57a083a57cb1f31db7000028

func PowersOfTwo(n int) []uint64 {
	result := make([]uint64, n+1)
	for i := 0; i <= n; i++ {
		result[i] = 1 << uint64(i)
	}
	return result
}

func TestPowersOfTwo(t *testing.T) {
	assert.Equal(t, []uint64{1}, PowersOfTwo(0))
	assert.Equal(t, []uint64{1, 2}, PowersOfTwo(1))
	assert.Equal(t, []uint64{1, 2, 4, 8, 16}, PowersOfTwo(4))
}
