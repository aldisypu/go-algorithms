package factorial

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func factorial(value int) int {
	if value <= 0 {
		return 1
	}

	result := 1

	for i := value; i >= 1; i-- {
		result *= i
	}
	return result
}

func factorialRecursive(value int) int {
	if value <= 0 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func factorialTailRecursive(total, value int) int {
	if value <= 0 {
		return total
	} else {
		return factorialTailRecursive(total*value, value-1)
	}
}

func TestFactorial(t *testing.T) {
	assert.Equal(t, 1, factorial(0))
	assert.Equal(t, 1, factorial(1))
	assert.Equal(t, 2, factorial(2))
	assert.Equal(t, 6, factorial(3))
	assert.Equal(t, 24, factorial(4))
	assert.Equal(t, 120, factorial(5))
}

func TestFactorialRecursive(t *testing.T) {
	assert.Equal(t, 1, factorialRecursive(0))
	assert.Equal(t, 1, factorialRecursive(1))
	assert.Equal(t, 2, factorialRecursive(2))
	assert.Equal(t, 6, factorialRecursive(3))
	assert.Equal(t, 24, factorialRecursive(4))
	assert.Equal(t, 120, factorialRecursive(5))
}

func TestFactorialTailRecursive(t *testing.T) {
	assert.Equal(t, 1, factorialTailRecursive(1, 0))
	assert.Equal(t, 1, factorialTailRecursive(1, 1))
	assert.Equal(t, 2, factorialTailRecursive(1, 2))
	assert.Equal(t, 6, factorialTailRecursive(1, 3))
	assert.Equal(t, 24, factorialTailRecursive(1, 4))
	assert.Equal(t, 120, factorialTailRecursive(1, 5))
}
