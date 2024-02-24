package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/5a3fe3dde1ce0e8ed6000097

func century(year int) int {
	return (year + 99) / 100
}

func TestCentury(t *testing.T) {
	assert.Equal(t, 20, century(1990))
	assert.Equal(t, 18, century(1705))
	assert.Equal(t, 19, century(1900))
	assert.Equal(t, 17, century(1601))
	assert.Equal(t, 20, century(2000))
}
