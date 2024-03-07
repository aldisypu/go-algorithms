package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/57eae65a4321032ce000002d

func FakeBin(x string) string {
	result := ""
	for _, digit := range x {
		if digit < '5' {
			result += "0"
		} else {
			result += "1"
		}
	}
	return result
}

func TestFakeBin(t *testing.T) {
	assert.Equal(t, "01011110001100111", FakeBin("45385593107843568"))
	assert.Equal(t, "101000111101101", FakeBin("509321967506747"))
	assert.Equal(t, "011011110000101010000011011", FakeBin("366058562030849490134388085"))
	assert.Equal(t, "01111100", FakeBin("15889923"))
	assert.Equal(t, "100111001111", FakeBin("800857237867"))
}
