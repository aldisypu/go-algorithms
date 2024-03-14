package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/5ce9c1000bab0b001134f5af

func QuarterOf(month int) int {
	return (month + 2) / 3
}

func TestQuarterOf(t *testing.T) {
	assert.Equal(t, 1, QuarterOf(3))
	assert.Equal(t, 3, QuarterOf(8))
	assert.Equal(t, 4, QuarterOf(11))
}
