package kata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://www.codewars.com/kata/57a5c31ce298a7e6b7000334

func BinToDec(bin string) int {
	result := 0

	for _, digit := range bin {
		result *= 2
		if digit == '1' {
			result += 1
		}
	}

	return result
}

func TestBinToDec(t *testing.T) {
	assert.Equal(t, 0, BinToDec("0"))
	assert.Equal(t, 1, BinToDec("1"))
	assert.Equal(t, 2, BinToDec("10"))
	assert.Equal(t, 42, BinToDec("101010"))
	assert.Equal(t, 73, BinToDec("1001001"))
}
