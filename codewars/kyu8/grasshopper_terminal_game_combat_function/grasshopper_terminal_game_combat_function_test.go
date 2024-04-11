package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/586c1cf4b98de0399300001d

func combat(health, damage float64) float64 {
	if health < damage {
		return 0
	}
	return health - damage
}

func TestCombat(t *testing.T) {
	assert.Equal(t, 50.0, combat(100.0, 50.0))
	assert.Equal(t, 0.0, combat(1.0, 100))
}
