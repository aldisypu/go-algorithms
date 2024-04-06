package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/50654ddff44f800200000007

func Solution(a, b string) string {
	if len(a) < len(b) {
		return a + b + a
	}
	return b + a + b
}

func TestSolution(t *testing.T) {
	assert.Equal(t, "1451", Solution("45", "1"))
	assert.Equal(t, "1320013", Solution("13", "200"))
	assert.Equal(t, "MeSoonMe", Solution("Soon", "Me"))
	assert.Equal(t, "UFalseU", Solution("U", "False"))
}
