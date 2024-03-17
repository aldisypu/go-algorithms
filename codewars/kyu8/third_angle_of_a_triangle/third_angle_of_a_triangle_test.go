package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/5a023c426975981341000014

func OtherAngle(a int, b int) int {
	return 180 - a - b
}

func TestOtherAngle(t *testing.T) {
	assert.Equal(t, 90, OtherAngle(30, 60))
	assert.Equal(t, 60, OtherAngle(60, 60))
	assert.Equal(t, 59, OtherAngle(43, 78))
	assert.Equal(t, 150, OtherAngle(10, 20))
}
