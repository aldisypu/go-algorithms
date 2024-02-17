package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/55cbd4ba903825f7970000f5

func GetGrade(a, b, c int) rune {
	average := (a + b + c) / 3

	switch {
	case average >= 90:
		return 'A'
	case average >= 80:
		return 'B'
	case average >= 70:
		return 'C'
	case average >= 60:
		return 'D'
	default:
		return 'F'
	}
}

func TestGetGrade(t *testing.T) {
	assert.Equal(t, 'A', GetGrade(95, 90, 93))
	assert.Equal(t, 'B', GetGrade(70, 70, 100))
	assert.Equal(t, 'C', GetGrade(70, 70, 70))
	assert.Equal(t, 'D', GetGrade(65, 70, 59))
	assert.Equal(t, 'F', GetGrade(44, 55, 52))
}
