package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/5933a1f8552bc2750a0000ed

func NthEven(n int) int {
	return n*2 - 2
}

func TestNthEven(t *testing.T) {
	assert.Equal(t, 0, NthEven(1))
	assert.Equal(t, 2, NthEven(2))
	assert.Equal(t, 4, NthEven(3))
	assert.Equal(t, 198, NthEven(100))
	assert.Equal(t, 2597466, NthEven(1298734))
}
