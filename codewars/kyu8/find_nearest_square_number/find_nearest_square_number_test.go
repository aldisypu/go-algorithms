package kata

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/5a805d8cafa10f8b930005ba

func NearestSq(n int) int {
	nearestRoot := int(math.Round(math.Sqrt(float64(n))))
	return nearestRoot * nearestRoot
}

func TestNearestSq(t *testing.T) {
	assert.Equal(t, 1, NearestSq(1))
	assert.Equal(t, 1, NearestSq(2))
	assert.Equal(t, 9, NearestSq(10))
	assert.Equal(t, 121, NearestSq(111))
	assert.Equal(t, 10000, NearestSq(9999))
}
