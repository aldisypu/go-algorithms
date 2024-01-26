package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func MakeNegative(x int) int {
	if x < 0 {
		return x
	}
	return -x
}

func TestMakeNegative(t *testing.T) {
	assert.Equal(t, -1, MakeNegative(1))
	assert.Equal(t, -5, MakeNegative(-5))
	assert.Equal(t, 0, MakeNegative(0))
}
