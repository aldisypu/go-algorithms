package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func EvenOrOdd(number int) string {
	if number%2 != 0 {
		return "Odd"
	}
	return "Even"
}

func TestEvenOrOdd(t *testing.T) {
	assert.Equal(t, "Odd", EvenOrOdd(1))
	assert.Equal(t, "Odd", EvenOrOdd(3))
	assert.Equal(t, "Odd", EvenOrOdd(5))
	assert.Equal(t, "Odd", EvenOrOdd(7))
	assert.Equal(t, "Odd", EvenOrOdd(9))

	assert.Equal(t, "Even", EvenOrOdd(2))
	assert.Equal(t, "Even", EvenOrOdd(4))
	assert.Equal(t, "Even", EvenOrOdd(6))
	assert.Equal(t, "Even", EvenOrOdd(8))
	assert.Equal(t, "Even", EvenOrOdd(10))
}
