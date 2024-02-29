package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/5545f109004975ea66000086

func IsDivisible(n, x, y int) bool {
	return n%x == 0 && n%y == 0
}

func TestIsDivisible(t *testing.T) {
	assert.True(t, IsDivisible(12, 3, 4))
	assert.True(t, IsDivisible(48, 3, 4))
	assert.True(t, IsDivisible(100, 5, 10))
	assert.False(t, IsDivisible(3, 3, 4))
	assert.False(t, IsDivisible(8, 3, 4))
}
