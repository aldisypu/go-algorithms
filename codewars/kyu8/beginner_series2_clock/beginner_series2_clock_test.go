package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/55f9bca8ecaa9eac7100004a

func Past(h, m, s int) int {
	return (h*3600 + m*60 + s) * 1000
}

func TestPast(t *testing.T) {
	assert.Equal(t, 61000, Past(0, 1, 1))
	assert.Equal(t, 3661000, Past(1, 1, 1))
	assert.Equal(t, 0, Past(0, 0, 0))
	assert.Equal(t, 3601000, Past(1, 0, 1))
	assert.Equal(t, 3600000, Past(1, 0, 0))
}
