package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/582cb0224e56e068d800003c

func Litres(time float64) int {
	return int(time * 0.5)
}

func TestLitres(t *testing.T) {
	assert.Equal(t, 1, Litres(2))
	assert.Equal(t, 0, Litres(1.4))
	assert.Equal(t, 6, Litres(12.3))
	assert.Equal(t, 0, Litres(0.82))
	assert.Equal(t, 0, Litres(0))
}
