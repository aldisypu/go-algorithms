package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/563b662a59afc2b5120000c6

func NbYear(p0 int, percent float64, aug int, p int) int {
	years := 0

	for p0 < p {
		p0 += int(float64(p0)*percent/100) + aug
		years++
	}

	return years
}

func TestNbYear(t *testing.T) {
	assert.Equal(t, 15, NbYear(1500, 5, 100, 5000))
	assert.Equal(t, 10, NbYear(1500000, 2.5, 10000, 2000000))
	assert.Equal(t, 94, NbYear(1500000, 0.25, 1000, 2000000))
	assert.Equal(t, 151, NbYear(1500000, 0.25, -1000, 2000000))
}
