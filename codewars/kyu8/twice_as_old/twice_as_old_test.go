package kata

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/5b853229cfde412a470000d0

func TwiceAsOld(dadYearsOld, sonYearsOld int) int {
	return int(math.Abs(float64(dadYearsOld - 2*sonYearsOld)))
}

func TestTwiceAsOld(t *testing.T) {
	assert.Equal(t, 22, TwiceAsOld(36, 7))
	assert.Equal(t, 5, TwiceAsOld(55, 30))
	assert.Equal(t, 0, TwiceAsOld(42, 21))
	assert.Equal(t, 20, TwiceAsOld(22, 1))
	assert.Equal(t, 29, TwiceAsOld(29, 0))
}
