package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/56fa3c5ce4d45d2a52001b3c

func Xor(a, b bool) bool {
	return a != b
}

func TestXor(t *testing.T) {
	assert.False(t, Xor(false, false))
	assert.True(t, Xor(true, false))
	assert.True(t, Xor(false, true))
	assert.False(t, Xor(true, true))
}
