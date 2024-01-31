package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/56606694ec01347ce800001b

func IsTriangle(a, b, c int) bool {
	return a+b > c && b+c > a && c+a > b
}

func TestIsTriangle(t *testing.T) {
	assert.False(t, IsTriangle(5, 1, 2))
	assert.False(t, IsTriangle(1, 2, 5))
	assert.False(t, IsTriangle(2, 5, 1))
	assert.True(t, IsTriangle(4, 2, 3))
	assert.True(t, IsTriangle(5, 1, 5))
	assert.True(t, IsTriangle(2, 2, 2))
	assert.False(t, IsTriangle(-1, 2, 3))
}
