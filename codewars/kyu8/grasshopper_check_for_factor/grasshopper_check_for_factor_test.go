package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/55cbc3586671f6aa070000fb

func CheckForFactor(base int, factor int) bool {
	return base%factor == 0
}

func TestCheckForFactor(t *testing.T) {
	assert.True(t, CheckForFactor(10, 2))
	assert.True(t, CheckForFactor(63, 7))
	assert.True(t, CheckForFactor(2450, 5))
	assert.False(t, CheckForFactor(9, 2))
	assert.False(t, CheckForFactor(653, 7))
}
