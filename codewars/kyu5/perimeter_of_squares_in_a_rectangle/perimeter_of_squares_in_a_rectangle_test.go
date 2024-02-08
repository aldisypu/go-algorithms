package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/559a28007caad2ac4e000083

func Perimeter(n int) int {
	if n < 0 {
		return 0
	}

	a, b := 1, 1
	sum := 0

	for i := 0; i <= n; i++ {
		sum += a
		a, b = b, a+b
	}

	return 4 * sum
}

func TestPerimeter(t *testing.T) {
	assert.Equal(t, 80, Perimeter(5))
	assert.Equal(t, 216, Perimeter(7))
	assert.Equal(t, 114624, Perimeter(20))
	assert.Equal(t, 14098308, Perimeter(30))
}
