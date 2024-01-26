package multiply

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/50654ddff44f800200000004

func Multiply(a, b int) int {
	return a * b
}

func TestMultiply(t *testing.T) {
	assert.Equal(t, 9, Multiply(1, 9))
}
