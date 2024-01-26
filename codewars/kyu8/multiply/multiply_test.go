package multiply

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Multiply(a, b int) int {
	return a * b
}

func TestMultiply(t *testing.T) {
	assert.Equal(t, 9, Multiply(1, 9))
}
