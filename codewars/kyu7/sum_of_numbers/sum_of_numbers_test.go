package kata

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/55f2b110f61eb01779000053

func GetSum(a, b int) int {
	a, b = int(math.Min(float64(a), float64(b))), int(math.Max(float64(a), float64(b)))

	return (b - a + 1) * (a + b) / 2
}

func TestGetSum(t *testing.T) {
	assert.Equal(t, 1, GetSum(0, 1))
	assert.Equal(t, 3, GetSum(1, 2))
	assert.Equal(t, 14, GetSum(5, -1))
	assert.Equal(t, 127759, GetSum(505, 4))
	assert.Equal(t, 44178, GetSum(321, 123))
	assert.Equal(t, -1, GetSum(0, -1))
	assert.Equal(t, -1275, GetSum(-50, 0))
	assert.Equal(t, -15, GetSum(-1, -5))
	assert.Equal(t, -5, GetSum(-5, -5))
	assert.Equal(t, -127755, GetSum(-505, 4))
	assert.Equal(t, -44055, GetSum(-321, 123))
	assert.Equal(t, 0, GetSum(0, 0))
	assert.Equal(t, -15, GetSum(-5, -1))
	assert.Equal(t, 15, GetSum(5, 1))
	assert.Equal(t, -17, GetSum(-17, -17))
	assert.Equal(t, 17, GetSum(17, 17))
}
