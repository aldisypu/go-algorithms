package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/56dec885c54a926dcd001095

func Opposite(value int) int {
	return -value
}

func TestOpposite(t *testing.T) {
	assert.Equal(t, -1, Opposite(1))
	assert.Equal(t, 1, Opposite(-1))
}
