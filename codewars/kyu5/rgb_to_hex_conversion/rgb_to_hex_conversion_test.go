package kata

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/513e08acc600c94f01000001

func RGB(r, g, b int) string {
	r = clamp(r)
	g = clamp(g)
	b = clamp(b)

	return fmt.Sprintf("%02X%02X%02X", r, g, b)
}

func clamp(value int) int {
	if value < 0 {
		return 0
	} else if value > 255 {
		return 255
	}
	return value
}

func TestRGB(t *testing.T) {
	assert.Equal(t, "000000", RGB(0, 0, 0))
	assert.Equal(t, "010203", RGB(1, 2, 3))
	assert.Equal(t, "FFFFFF", RGB(255, 255, 255))
	assert.Equal(t, "FEFDFC", RGB(254, 253, 252))
	assert.Equal(t, "00FF7D", RGB(-20, 275, 125))
}
