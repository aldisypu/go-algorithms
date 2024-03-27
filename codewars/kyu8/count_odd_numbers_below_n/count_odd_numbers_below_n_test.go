package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/59342039eb450e39970000a6

func OddCount(n int) int {
	return n / 2
}

func TestOddCount(t *testing.T) {
	assert.Equal(t, 7, OddCount(15))
	assert.Equal(t, 7511, OddCount(15023))
}
