package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/5547cc7dcad755e480000004

func RemovNb(m uint64) [][2]uint64 {
	var result [][2]uint64

	totalSum := m * (m + 1) / 2

	for a := uint64(1); a <= m; a++ {
		b := (totalSum - a) / (a + 1)
		if a*b == totalSum-a-b && b <= m {
			result = append(result, [2]uint64{a, b})
		}
	}

	return result
}

func TestRemovNb(t *testing.T) {
	assert.Equal(t, [][2]uint64{{15, 21}, {21, 15}}, RemovNb(26))
	assert.Equal(t, [][2]uint64{{55, 91}, {91, 55}}, RemovNb(101))
	assert.Equal(t, [][2]uint64{{70, 73}, {73, 70}}, RemovNb(102))
	assert.Equal(t, [][2]uint64{{70, 85}, {85, 70}}, RemovNb(110))
	assert.Nil(t, RemovNb(100))
}
